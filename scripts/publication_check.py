#!/usr/bin/env python3
# SPDX-License-Identifier: Apache-2.0
"""Report publication-review candidates. This is not a legal clearance tool.

Reads tracked working-tree files only; never modifies them or prints their contents.
Private experiment citations are counted separately, not treated as broken links.
"""
from __future__ import annotations

import argparse
import json
import posixpath
import re
import subprocess
import sys
from collections import Counter
from pathlib import Path
from urllib.parse import unquote, urlsplit

ASSET_EXTENSIONS = frozenset({
    '.exe', '.dll', '.res', '.alm', '.sav', '.ags', '.16', '.16a', '.256',
    '.pal', '.wav', '.mp3', '.ogg', '.smk', '.png', '.jpg', '.jpeg', '.bmp',
    '.gif', '.zip', '.7z', '.rar', '.tar', '.gz', '.bin', '.dat', '.ttf', '.otf',
})
ID = re.compile(r'\b[A-Z][A-Z0-9]*(?:-[A-Z0-9]+)*-[0-9]{3}\b')
ASM = re.compile(
    r'\b(?:MOV(?:SX|ZX)?|LEA|PUSH|POP|CMP|TEST|XOR|AND|OR|ADD|SUB|'
    r'INC|DEC|NEG|NOT|XCHG|ADC|SBB|RO[LR]|SET(?:[GL]E?|[AB]E?|N?[EZ])|'
    r'SH[LR]|SA[LR]|IMUL|IDIV|MUL|DIV|CALL|RET|'
    r'J(?:MP|Z|NZ|E|NE|A|AE|B|BE|G|GE|L|LE)|'
    r'F(?:LD|LDZ|LD1|STP?|ILD|ISTP?|ADD|ADDP|SUB|MUL|DIV|SIN|COS|SQRT|COMP))'
    r'\s+(?:(?:byte|word|dword|qword)\s+ptr\b|\[|0x[0-9A-Fa-f]+|reg\b|'
    r'(?:E(?:AX|BX|CX|DX|SI|DI|BP|SP)|(?:A|B|C|D)[HLX]|SI|DI|BP|SP)\b)'
)
# Lowercase listings use the same mnemonics. Requiring a register, size or
# memory operand keeps ordinary prose such as "or 0x20" out of the match.
ASM_LOWER = re.compile(
    r'\b(?:mov(?:sx|zx)?|lea|push|pop|cmp|test|xor|and|or|add|sub|'
    r'inc|dec|neg|xchg|adc|sbb|ro[lr]|set(?:[gl]e?|[ab]e?|n?[ez])|'
    r'sh[lr]|sa[lr]|imul|idiv|mul|div|call|ret|'
    r'j(?:mp|z|nz|e|ne|a|ae|b|be|g|ge|l|le)|'
    r'f(?:ld|ldz|ld1|stp?|ild|istp?|add|addp|sub|mul|div|sin|cos|sqrt|comp))'
    r'\s+(?:(?:byte|word|dword|qword)\b|\[(?![`\]])(?![^\]\n]*\]\()|'
    r'(?:e(?:ax|bx|cx|dx|si|di|bp|sp)|(?:a|b|c|d)[hlx])\b)'
)
HEX_RUN = re.compile(r'(?<![0-9A-Fa-f])(?:[0-9A-Fa-f]{2}[ \t]+){15}[0-9A-Fa-f]{2}(?![0-9A-Fa-f])')
BASE64_RUN = re.compile(r'(?<![A-Za-z0-9+/])(?:[A-Za-z0-9+/]{4}){40,}={0,2}(?![A-Za-z0-9+/])')
# Deliberately limited inline-link parser: reference-style links and complex
# parenthesized destinations still require a separate human/link-checker pass.
LINK = re.compile(r'\[[^\]\n]*\]\(([^)\n]+)\)')
# A link destination is a path, not table content.
LINK_DESTINATION = re.compile(r'\]\([^)\n]*\)')
CONTENT_TABLE = re.compile(r'^\s*\|(?=[^\n]*\bname\b)(?=[^\n]*\b(?:score|price|damage|health)\b).*\|', re.IGNORECASE)
PRIVATE_ROOTS =frozenset({'experiments', 'pipeline', 'gameversions'})


def split_row(line: str) -> list[str]:
    """Split table cells without consuming Windows paths or quoted pipes."""
    s = line.strip()
    if not s.startswith('|'):
        return []
    cells, current = [], []
    delimiter = 0
    i = 0
    while i < len(s):
        char = s[i]
        if char == '\\' and i + 1 < len(s) and s[i + 1] in '|\\':
            current.extend(s[i:i + 2])
            i += 2
            continue
        if char == '`':
            end = i + 1
            while end < len(s) and s[end] == '`':
                end += 1
            run = end - i
            if delimiter == 0:
                delimiter = run
            elif delimiter == run:
                delimiter = 0
            current.append(s[i:end])
            i = end
            continue
        if char == '|' and delimiter == 0:
            cells.append(''.join(current).strip())
            current = []
        else:
            current.append(char)
        i += 1
    cells.append(''.join(current).strip())
    if cells and not cells[0]:
        cells.pop(0)
    if cells and not cells[-1]:
        cells.pop()
    return cells


def finding(path: str, line: int, code: str, severity: str, claim: str = '') -> dict:
    return {'path': path, 'line': line, 'code': code, 'severity': severity, 'claim': claim}


def analyze(files: dict[str, bytes]) -> dict:
    """Analyze supplied path/byte pairs; pure and usable with synthetic fixtures."""
    findings: list[dict] = []
    text_files: dict[str, str] = {}
    private_citations = 0
    external_links = 0
    definitions: set[str] = set()
    retractions: set[str] = set()
    for path, data in sorted(files.items()):
        if Path(path).suffix.lower() in ASSET_EXTENSIONS:
            findings.append(finding(path, 0, 'ASSET_OR_BINARY_EXTENSION', 'error'))
        if b'\0' in data:
            findings.append(finding(path, 0, 'NUL_BYTES', 'error'))
            continue
        try:
            text = data.decode('utf-8')
        except UnicodeDecodeError:
            findings.append(finding(path, 0, 'NON_UTF8', 'error'))
            continue
        if not path.endswith('.md'):
            continue
        text_files[path] = text
        if path.startswith('claims/') and path != 'claims/registry.md':
            for line in text.splitlines():
                cells = split_row(line)
                if cells:
                    target = retractions if path == 'claims/retracted.md' else definitions
                    target.update(ID.findall(cells[0]))

    available = definitions | retractions
    for path, text in sorted(text_files.items()):
        # Scan research-facing content, not this tool's policy examples or licenses.
        is_knowledge = path.startswith(('claims/', 'formats/'))
        in_fence = False
        for number, line in enumerate(text.splitlines(), 1):
            cells = split_row(line)
            ids = ID.findall(cells[0]) if cells else []
            claim = ','.join(ids)
            fence = re.match(r'^\s*(`{3,}|~{3,})(\w*)', line)
            if fence:
                in_fence = not in_fence
                if is_knowledge and fence.group(2).lower() in {'asm', 'assembly', 'x86', 'nasm'}:
                    findings.append(finding(path, number, 'ASSEMBLY_FENCE_REVIEW', 'review', claim))
            if is_knowledge:
                for pattern, code in ((ASM, 'INSTRUCTION_EXCERPT_REVIEW'),
                                      (ASM_LOWER, 'INSTRUCTION_EXCERPT_REVIEW'),
                                      (HEX_RUN, 'BYTE_RUN_REVIEW'),
                                      (BASE64_RUN, 'ENCODED_PAYLOAD_REVIEW')):
                    if pattern.search(line):
                        findings.append(finding(path, number, code, 'review', claim))
                if CONTENT_TABLE.search(LINK_DESTINATION.sub('](', line)):
                    findings.append(finding(path, number, 'CONTENT_TABLE_REVIEW', 'review', claim))
                if 'pipeline/reviews/' in line:
                    findings.append(finding(path, number, 'INTERNAL_WORKFLOW_REVIEW', 'review', claim))
                if path.startswith('formats/'):
                    for missing in sorted(set(ID.findall(line)) - available):
                        findings.append(finding(path, number, 'UNRESOLVED_CLAIM_REVIEW', 'review', missing))
            if in_fence:
                continue
            for match in LINK.finditer(line):
                destination = match.group(1).strip().split(' "', 1)[0].strip('<>')
                try:
                    parsed = urlsplit(destination)
                except ValueError:
                    findings.append(finding(path, number, 'INVALID_LINK', 'error', claim))
                    continue
                if parsed.scheme or parsed.netloc:
                    external_links += 1
                    continue
                if not parsed.path:
                    continue
                raw = unquote(parsed.path).replace('\\', '/')
                parts = [p for p in raw.split('/') if p not in {'', '.', '..'}]
                if parts and parts[0] in PRIVATE_ROOTS:
                    private_citations += 1
                    continue
                target = posixpath.normpath(posixpath.join(posixpath.dirname(path), raw))
                exists = target in files or target == '.' or any(p.startswith(target.rstrip('/') + '/') for p in files)
                if not exists:
                    findings.append(finding(path, number, 'MISSING_LOCAL_LINK', 'error', claim))

    counts = Counter(f['severity'] for f in findings)
    return {
        'scope': 'tracked working-tree files; not historical Git objects',
        'legal_clearance': 'not_assessed',
        'files': len(files), 'claim_ids': len(definitions),
        'retracted_ids': len(retractions), 'private_citations': private_citations,
        'external_links': external_links,
        'errors': counts['error'], 'review_candidates': counts['review'],
        'findings': findings,
    }


def tracked_files(root: Path) -> dict[str, bytes]:
    result = subprocess.run(['git', 'ls-files', '-z'], cwd=root, check=True, capture_output=True)
    files = {}
    for encoded in result.stdout.split(b'\0'):
        if not encoded:
            continue
        path = encoded.decode('utf-8')
        resolved = root / path
        # Do not follow a tracked symlink outside the checkout or execute anything.
        if resolved.is_symlink():
            raise ValueError(f'tracked symlink requires manual review: {path}')
        if not resolved.is_file():
            raise ValueError(f'tracked file missing or not regular: {path}')
        files[path] = resolved.read_bytes()
    if not files:
        raise ValueError('no tracked files; stage the intended candidate first')
    return files


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--root', type=Path, default=Path(__file__).resolve().parents[1])
    parser.add_argument('--json', action='store_true', help='emit a machine-readable report without source excerpts')
    parser.add_argument('--strict', action='store_true', help='also fail on heuristic review candidates')
    args = parser.parse_args()
    try:
        report = analyze(tracked_files(args.root.resolve()))
    except (OSError, ValueError, subprocess.CalledProcessError) as exc:
        print(f'publication-check: {exc}', file=sys.stderr)
        return 2
    if args.json:
        print(json.dumps(report, ensure_ascii=True, indent=2))
    else:
        print(f"{report['files']} files; {report['errors']} mechanical errors; "
              f"{report['review_candidates']} review candidates; "
              f"{report['private_citations']} private citations.")
        for item in report['findings']:
            print(f"{item['severity']}: {item['path']}:{item['line']} "
                  f"{item['code']} {item['claim']}")
        print('Not legal clearance. No matches is not proof of independent origin or absence of game data.')
    return int(report['errors'] > 0 or (args.strict and report['review_candidates'] > 0))


if __name__ == '__main__':
    raise SystemExit(main())
