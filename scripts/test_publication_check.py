# SPDX-License-Identifier: Apache-2.0
"""Synthetic fixtures only; no original game data or code is used."""
import unittest
from publication_check import analyze, split_row


class PublicationCheckTests(unittest.TestCase):
    def codes(self, files):
        return {f['code'] for f in analyze(files)['findings']}

    def test_synthetic_instruction_is_review_not_infringement(self):
        result = analyze({'claims/test.md': b'| TEST-ONE-001 | `MOV EAX,0x1234` | High | active | local |'})
        self.assertIn('INSTRUCTION_EXCERPT_REVIEW', {f['code'] for f in result['findings']})
        self.assertEqual(result['errors'], 0)
        self.assertEqual(result['legal_clearance'], 'not_assessed')

    def test_content_table_is_reviewed_even_when_text(self):
        self.assertIn('CONTENT_TABLE_REVIEW', self.codes({'formats/test.md': b'| # | name | score |\n| 0 | Example | 123 |'}))

    def test_lowercase_instruction_is_review(self):
        self.assertIn('INSTRUCTION_EXCERPT_REVIEW', self.codes({'claims/test.md': b'`cmp dword [ecx+0x10],0x0`'}))

    def test_increment_instruction_is_review(self):
        self.assertIn('INSTRUCTION_EXCERPT_REVIEW', self.codes({'claims/test.md': b'`00401000 INC [EBP + 0x10]`'}))

    def test_prose_conjunction_is_not_instruction(self):
        self.assertNotIn('INSTRUCTION_EXCERPT_REVIEW', self.codes({'claims/test.md': b'bit 0x4 or 0x20 and the flag'}))
        self.assertNotIn('INSTRUCTION_EXCERPT_REVIEW', self.codes({'claims/test.md': b'this is not [`inv.md`](inv.md) and [page](x.md)'}))

    def test_link_destination_is_not_table_content(self):
        row = b'| TEST-ONE-001 | the bare item name | High | active | [E](../experiments/EXP-9999-price-line/) |'
        self.assertNotIn('CONTENT_TABLE_REVIEW', self.codes({'claims/test.md': row}))

    def test_formula_and_offsets_not_flagged(self):
        self.assertFalse(analyze({'formats/test.md': b'offset 0x04; value = (x << 2) | (x >> 4)\n'})['findings'])

    def test_digest_is_not_a_byte_dump(self):
        self.assertNotIn('BYTE_RUN_REVIEW', self.codes({'claims/test.md': ('`' + 'a' * 64 + '`').encode()}))

    def test_spaced_byte_dump_is_review_candidate(self):
        self.assertIn('BYTE_RUN_REVIEW', self.codes({'claims/test.md': ('12 ' * 16).encode()}))

    def test_private_experiment_citation_is_counted(self):
        result = analyze({'claims/test.md': b'[EXP](../experiments/EXP-9999-test/)\n'})
        self.assertEqual(result['private_citations'], 1)
        self.assertFalse(result['findings'])

    def test_missing_public_ledger_is_error(self):
        self.assertIn('MISSING_LOCAL_LINK', self.codes({'formats/test.md': b'[ledger](../claims/missing.md)'}))

    def test_matching_local_link(self):
        self.assertNotIn('MISSING_LOCAL_LINK', self.codes({'README.md': b'[page](formats/test.md)', 'formats/test.md': b'test'}))

    def test_extension_is_case_insensitive(self):
        self.assertIn('ASSET_OR_BINARY_EXTENSION', self.codes({'sample.EXE': b'synthetic'}))

    def test_binary_content_not_hidden_by_markdown_extension(self):
        self.assertIn('NUL_BYTES', self.codes({'claims/test.md': b'a\0b'}))
        self.assertIn('NON_UTF8', self.codes({'claims/test.md': b'\xff'}))

    def test_link_inside_fence_not_followed(self):
        self.assertNotIn('MISSING_LOCAL_LINK', self.codes({'formats/test.md': b'```text\n[x](missing)\n```\n'}))

    def test_no_false_legal_pass(self):
        result = analyze({'README.md': b'plain text'})
        self.assertEqual(result['errors'], 0)
        self.assertEqual(result['legal_clearance'], 'not_assessed')

    def test_claim_definition_and_missing_reference(self):
        result = analyze({'claims/test.md': b'| TEST-001 | a | High | active | source |',
                          'formats/test.md': b'TEST-001 TEST-002'})
        missing = [f['claim'] for f in result['findings'] if f['code'] == 'UNRESOLVED_CLAIM_REVIEW']
        self.assertEqual(missing, ['TEST-002'])

    def test_retraction_ids_are_available(self):
        result = analyze({'claims/retracted.md': b'| `TEST-001`, `TEST-002` | former | Medium | source | correction |',
                          'formats/test.md': b'TEST-002'})
        self.assertEqual(result['retracted_ids'], 2)
        self.assertFalse(result['findings'])

    def test_escaped_and_code_pipes_preserve_cells(self):
        cells = split_row(r'| TEST-001 | `a|b` and c\|d and C:\data | High |')
        self.assertEqual(len(cells), 3)
        self.assertEqual(cells[2], 'High')
        self.assertIn(r'C:\data', cells[1])

    def test_reference_style_link_limitation_is_not_faked(self):
        result = analyze({'README.md': b'[label][reference]\n[reference]: missing.md\n'})
        self.assertEqual(result['legal_clearance'], 'not_assessed')
        self.assertEqual(result['external_links'], 0)


if __name__ == '__main__':
    unittest.main()
