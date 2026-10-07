# Publication policy

This is a maintainer review policy, not a legal opinion or an additional condition on
reuse of work licensed under CC BY 4.0 or Apache-2.0.

## Preserve useful knowledge

Prefer independently written descriptions of file layouts, API/wire contracts,
algorithms, state transitions, exact arithmetic and observed behaviour. Keep the claim
ID, evidence identity, confidence, scope, edge cases and explicit Unknowns. A finite
corpus observation must not become a universal rule during shortening.

Field offsets, wire offsets and constants stay: they are the interoperability contract.
Executable locators do not: routine names and code, data and library addresses are
replaced at export by opaque aliases (`R0001`, `L00001`, `D00001`; `R2.`, `L2.`, `D2.`
for the second game), and the mapping stays in the private research source. Review
literal instruction sequences separately from independently expressed functional rules.
Do not use a global regular expression to delete instructions: some legacy rows express
an essential condition only in the cited operation.

## Material needing a separate decision

Do not export original game/decoder binaries, assets, complete resource payloads or
reconstructable content tables merely because they are text rather than binary files.
Review long code excerpts, disassembly/decompilation, quoted game text and encoded data
before export. Keep unnecessary original expression in private evidence. Any retained
excerpt needs a recorded purpose, proportionate scope, source and applicable permission
or other legal basis; a technical confidence grade is not that basis.

A paraphrase alone does not decide whether obtaining or publicly disclosing the
underlying information is permitted. Likewise "old game", "noncommercial", "clean
room" and "interoperability" are not blanket clearances. Apply the appropriate rules
to the actual method, purpose, component and jurisdiction rather than making an
all-repository legal guarantee.

## Evidence and authoring

Maintain two linked records: the unabridged private research snapshot and the public
functional edition. Record material omissions and editorial transformations without
fabricating experiments or changing what an old confidence level meant. Preserve
retractions and their clause scope; do not erase a former error to improve appearances.

Factual corrections belong in the research process. Publication corrections can be
reviewed here, but must feed the upstream export transformation before the next release.
Do not maintain two independently edited factual ledgers.

## Before approving an export

1. Record input identity, acquisition/terms review, method and purpose for the affected
   game or component. Separately assess the need and basis for public disclosure. Do
   not infer these from a SHA or a general README statement.
2. Inspect content as well as filenames, including reconstructable tables, quotations,
   code excerpts and embedded encodings. Run the supplied check as an aid, not a verdict.
3. Confirm that all public normative claims have available authorities or are clearly
   marked incomplete; internal experiment citations are allowed but not publicly
   reproducible evidence. Check license notices and public claims against real contents.
4. Verify claim IDs, confidence levels and retraction scope against the source; inspect
   the semantic diff. Review branches, tags and history separately when withdrawal of
   earlier material is required.

## Tool limits and current backlog

`python3 scripts/publication_check.py --strict` examines tracked working-tree content,
not previous commits, Git LFS storage, releases, attachments or caches. It reports only
locations/codes, not the potentially sensitive source excerpts. Inline-link parsing and
instruction detection are heuristic; reference-style links, unusual encodings and
paraphrased copied expression can be missed. Small legitimate technical fragments may
be flagged. No automatic deletion or approved-all baseline is provided.

### Completed in `publication/legal-hygiene-k1`

- `claims/registry.md` was replaced by a public index; private experiment/review narrative
  was removed from that surface.
- `claims/fame.md` and `formats/fame/format.md` no longer reproduce the shipped default
  hall-of-fame table or unnecessary instruction excerpts.
- `formats/res/format.md` now states the RES wire grammar and resolver semantics without
  executable-address/disassembly evidence or the complete shipped archive inventory.
- `formats/text/format.md` keeps byte conversion, indexing and table semantics while
  omitting localized prose and reconstructable shipped string/content tables.
- `claims/video.md` and `formats/video/format.md` now expose the ROM1 game-side media
  contract without bundled-Smacker structure offsets, instruction evidence or complete
  media-name/content inventories. Decoder-internal evidence remains private pending a
  separate third-party review.
- README/NOTICE/license scope, source mapping, third-party boundary, contribution rules,
  review tooling, synthetic tests and read-only CI were added or corrected in the first
  pass.

### Completed ROM2 functional edition

The 28-row asset ledger and its dependent survey text have a bounded semantic
review. Functional wording preserves complete confidence/evidence cells, earlier
withdrawals and Unknowns. One false maxima attribution receives an explicit
clause-scoped retraction; contradictory arithmetic/census prose follows the
already accepted tables. No original payload, title CSV or instruction listing
is added. The export checks source parity, claim dependencies and the reviewed
removals. Supporting experiment artifacts remain private.

### Completed functional edition k2

Every claim ledger and format page was reviewed for literal instruction text and
reconstructable content tables. 1,690 flagged lines in 77 files now state the same
operation functionally: the condition, arithmetic, loop or data movement, with field
offsets, constants, widths, branch sense, call targets and code addresses kept as
locators. Literal instruction sequences, register-by-register transliteration and
machine-code byte runs were removed. Retracted former wording is unchanged unless it
carried instruction text, in which case it is labelled as functional wording. A
fresh semantic review checked token loss, meaning, retraction scope and identity;
its findings were corrected once. Claim IDs, Status and Evidence cells and
Confidence grades are unchanged.

The strict check now also detects lowercase listings and more mnemonics, and ignores
link destinations when testing for content tables.

### Locator aliases and method index

From k194 the export replaces every executable locator, in each spelling the ledgers
use, with an opaque alias from an append-only private map, and fails when one remains.
An alias only links statements about the same routine or location. METHODS.md labels
each claim as observation, static analysis or mixed by a conservative rule.

### Remaining review

- Four prose rows in `claims/retracted.md` (SAV-608, ITEM-SCALE-017, R2-ENGINE-009,
  TOWN-093) still match the content-table or lowercase-listing heuristics. They quote
  no table or listing; the former wording is kept verbatim.
- Single register names and mnemonics used as locators in prose ("the count in ECX",
  "the return at X") remain where they name a fact rather than quote a sequence.
- Component-specific publication bases and acquisition/terms records.

Neither this checklist nor a green synthetic-test job is legal clearance.

A content cleanup commit does not remove its parent tree. Changes to history, `k1`,
visibility or published copies require a distinct, explicit decision; preserve private
provenance rather than destroying evidence.
