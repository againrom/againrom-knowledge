# Publication hygiene: first editorial pass

Base: the unedited research-derived export at private research source
`80670a957c38a6b281e8144e6ad76dc81cd1abb9`. That raw tree is not part of this
repository's published history; the first published tree is the edition described here.

## Changes in this branch

- README/NOTICE now distinguish the actual corpus and source policy from an absolute
  no-excerpts/no-exposure or legal-clearance assurance. Scope includes selected ROM2
  research and static, runtime and corpus evidence.
- SOURCE identifies the editorial edition separately from k1, records the omitted
  R2-ASSET ledger, and preserves original manifest counts as reported metadata.
- THIRD_PARTY distinguishes an external source from a bundled component being analyzed,
  including Smacker, and limits license statements to rights the project can grant.
- FAME's full ten-entry shipped name/score table and redundant name list are omitted
  from the public descriptions. The wire layout, count/size measurements, field-width
  evidence and uncertainty remain. The old table plus the published layout and zero
  tails reconstructs 228 bytes with SHA-256
  `1845f7728c6ee8b27c6a7bf086f0e09481412468705df3386bd922a8db1cbe9d`;
  retaining the digest does not redistribute that payload.
- FAME's literal record-stride and name-length instruction excerpts are replaced with
  descriptions of the same operations. No new experiment is asserted; its eight claim
  IDs, status values and confidence qualifications are preserved.
- The reader receives a minimal Go module file so the documented package invocation
  does not depend on a parent research checkout. No `require` dependency is added.
- A non-mutating publication check and synthetic tests report candidates and broken
  references without treating a regex result as evidence of infringement or clearance.

## Not completed or changed

This is not a whole-corpus sanitization. Other claim ledgers, retractions and the long
registry still require a semantic review. The R2-ASSET public evidence gap is recorded,
not repaired by copying an unreviewed private file or inventing stub claims. Component-
specific legal bases, purchase/terms verification and complete historical-object scans
are not certified by this pass.

No existing license text, factual research source, original-game rule or retraction
record is changed. The private research record is unchanged and complete. Publishing an
edited first tree is not a withdrawal of anything the private record holds, and it does
not retroactively establish a clean room.

## Semantic corrections after the first pass

A review of the first pass against the raw export found semantic drift that the
publication rule does not permit, and this edition corrects it. `claims/video.md`
regains one status cell (`VIDEO-045`, which clause the amendment refuted), the
`Medium` grade for EN/RU sample parity (`VIDEO-048`), the decoder-internal Unknown
(`VIDEO-034`) and the residual-scope statements of `VIDEO-SFX-014`, `VIDEO-SFX-016`,
`VIDEO-MUSIC-011`, `VIDEO-MUSIC-012`, `VIDEO-049`, `VIDEO-050` and `VIDEO-051`; an
Unknown the raw export did not carry was removed from `VIDEO-045`. Exact arithmetic
returns as functional knowledge: the Russian lowercase helper's two byte ranges in
`formats/text/format.md`, and the member count, archive length, archive digest and
duration arithmetic of `VIDEO-MUSIC-001` and `VIDEO-MUSIC-003`. Canonical format pages
again cite every claim id literally instead of abbreviating a family as a range, and
`formats/res/format.md` regains the citations it dropped. `RES-HDR-001` stays withheld
and [SOURCE.md](SOURCE.md) says so.

None of this reinstates an instruction listing, a shipped-content table or quoted game
text. Confidence levels, statuses and Unknowns are restored to what the research record
states, never raised.

## ROM2 functional edition

Snapshot k15 derives from accepted research
`845a36f4276b7e0c23db2710895c3e2f7857729c`. Its 28 R2-ASSET claim rows now accompany
the seven survey pages and both indices. All full Confidence/Evidence cells and
27 Status cells are preserved. R2-ASSET-027 alone carries an explicit partial
retraction of the false joint maxima attribution; the measured maxima remain.
R2-ASSET-024's contradictory arithmetic/census expressions are repaired from the
existing accepted field tables and results, with no new measurement.

Unnecessary literal operations and diagnostic excerpts are independently
expressed as functional rules. Earlier inline withdrawals, scope limits and
Unknowns remain. No raw evidence, complete content table, title CSV, game payload
or original executable material is exported. Private evidence identities remain
traceable through SOURCE and the claim Evidence cells.

The mandatory ROM2 export check covers exact source text in 12 pages, all 49 R2
claim identities and their references, including the 28 asset rows and formal
retraction scope. Omission of a row or ledger, metadata drift and reintroduced
reviewed long lines fail independently of the unrelated editorial backlog.
This finite check supplements semantic review; it does not certify arbitrary
future paraphrases or the rest of the publication history. The k1 statements
above describe that earlier edition and are retained as history.

## Functional edition k2

Snapshot k193 derives from accepted research
`611f926fdd252a14e9ea724aa5fb9b24b977f4d9`. Every claim ledger and format page
flagged for literal instruction text, machine-code bytes or content-table
patterns was rewritten upstream: 1,690 lines in 77 files. Each rewrite states
the operation the code performs on data and keeps offsets, constants, widths,
branch sense, call targets and code addresses as locators. All 2,945 claim IDs,
every Status and Evidence cell and every Confidence grade are unchanged; 20
Confidence cells and 76 Confidence paragraphs changed in argument prose only.
Retracted former wording is unchanged unless it carried instruction text, and
then it is labelled as functional wording.

A semantic review compared every lost address, offset and function token,
sampled the rewrites per area, and checked retraction scope; the nine lost
facts and two wording drifts it found were restored once. The literal text
remains in the private research history. `scripts/publication_check.py` now
detects lowercase listings and the increment, decrement, negation, exchange,
rotate and set-on-condition mnemonics, and ignores link destinations when it
tests for content tables.
Three links to private research files (`METHODOLOGY.md`, `docs/REGISTRY-LOG.md`)
are stated as plain text.

## Locator aliases and method index

Snapshot k194 derives from research `cb47bc9466d57c7779349392cb88b0e4de54f695` and
starts this repository's history. The export replaces routine names and code, data
and library addresses of the original executables with opaque aliases: written as
`FUN_` names, `0x` values, bare eight- and six-digit tokens, little-endian byte
strings and evidence file names. Partial address patterns become neutral wording. The
mapping is an append-only table in the private research source; the same address
keeps the same alias in later snapshots. 39,021 occurrences were replaced; the export
check finds none left. Field offsets, constants and claim text are otherwise
unchanged. Twelve short instruction-byte runs beside an alias remain; they carry no
address.

METHODS.md labels all 2,948 claim IDs: 95 observation, 2,429 static analysis and 424
mixed. A claim citing any executable location counts as static analysis.
