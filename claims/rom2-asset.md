# Claim registry — ROM2 asset identity

Level 2 ledger. Index: [registry.md](registry.md). Spec pages:
[`formats/rom2/`](../formats/rom2/README.md). Legend: `✔ promoted` (also cited by a
`formats/rom2-*/format.md` page) · `● active` · `✖ retracted`. IDs are permanent.

This ledger holds `R2-` claims about **Rage of Mages II** from one preserved Russian
installation. A ROM1 claim is a comparison reference, never the evidence for a ROM2
finding; these R2-ASSET claims likewise do not establish ROM1 behaviour.

The initial layout survey replays published ROM1 grammars on the ROM2 population:
the research archive reader (`internal/rom.OpenArchive`) is reused unmodified, and
record grammars are transcribed from promoted claims into `tools/r2asset`. Later
rows add the accepted ALM structural walks and static reads of dispatchers,
constructors, cross-references and compiled-function comparisons. Each row names
its actual instrument, population, controls and limits. A static code read is not
a runtime observation; a complete layout survey is not a complete semantic decoder.

The functional text retains the permanent IDs, confidence, amendments and private
evidence identities. Evidence-column experiment links locate the unabridged private
research record; named evidence files belong to that record, not to the public
snapshot. Public snapshots identify their exact source commit in `SOURCE.md`.
Unknown field meanings, consumers and codepage assignments remain Unknown.

| ID | Claim | Confidence | Status | Evidence |
|----|-------|-----------|--------|----------|
| R2-ASSET-001 | ROM2's `.res` container is byte-identical to ROM1's `&YA1` container (`RES-MAGIC-001`, `RES-HDR-002`, `RES-ACCEPT-031`, cross-reference only). | High | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-002 | ROM2's `.alm` file header uses ROM1's exact field layout (`ALM-FRAME-031`, `ALM-HDR-001`, cross-reference only), not byte-identical VALUES. | High | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-003 | Past record 0, ROM2's `.alm` record chain uses ROM1's exact ALM-FRAME-031 record framing (the 20-byte `[tag][hdrLen][payloadSize][typeId][f32]` header, `payloadSize` used unmodified to advance to the next record) — **all 83… | High | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-004 | `world.res:data/data.bin` and `world_srv.res:data/data.bin` (149 971 and 152 527 bytes) both parse under DAT-GRAM-003's wire grammar (cross-reference only) identically to ROM1 EN through group A's title array: the leading… | High | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-005 | The ROM2 root's `templates.bin` (61 802 bytes) is a distinct file from `world.res`'s own `data/data.bin` — different size, different content (`R2-ASSET-004`) — but whether it shares DAT-GRAM-003's own wire grammar at all is… | High / Unknown | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-006 | ROM2's nested inline registries are byte-identical to ROM1's REG-FMT-031 grammar (cross-reference only), found the same way RES-SCOPE-015 found ROM1's — by `&YA1` magic inside a container's file payloads, not by a `.reg` name. | High | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-007 | ROM2's `.16a` and `.256` sprite containers are byte-identical in framing to ROM1's (`SPR16A-STRUCT-001`, `SPR256-STRUCT-001`, cross-reference only), measured over the FULL population found by the extension census, not a… | High | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-008 | ROM2's `.pal` palettes match one of ROM1's two known shapes (BMP colour table at a fixed seek, or a flat 16×1024-byte block — `formats/pal/format.md`, cross-reference only) on **159/159** (100%) of the full population found… | High | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-009 | ROM2's `.16` bitmap font containers (the SPR16A-TRLR-012 no-palette variant, cross-reference only) are a complete population of 3 files, and 2 of the 3 do not tile exactly to the trailer despite every individual frame validating. | High | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-010 | ROM2's `main.res` still carries a `text/` directory of CRLF-delimited string-table files (TEXT-STRTAB-023, cross-reference only): all 15 of TEXT-STRTAB-023's named files (`main.txt` through `credits.txt`) are present under… | High | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-011 | ROM2's text-shaped nodes (`.txt`/`.ini`/`.lst`/extensionless payloads in any of the 11 `.res` containers — 69 nodes, 513 002 bytes: 68 `.txt` files corpus-wide plus 1 extensionless node, ROM2 having 0 `.ini`/`.lst` files… | High | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-012 | `world.res` and `world_srv.res` hold the identical 5 top-level `data/` entry names (`ai.reg`, `data.bin`, `itemname.bin`, `itemname.pkt`, `map.reg` — 0 client-only, 0 server-only entries), diffed with the unmodified ROM1… | High | ✔ promoted | [EXP-2000](../experiments/EXP-2000-asset-identity/) |
| R2-ASSET-017 | ROM2's per-map `typeId` set and record order strictly contain ROM1's own (Q1; `R2-ASSET-003` already publishes ROM2's own flat order and set — this row adds the cross-game RELATION and one population exception). | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-018 | Nine of the ten shared `typeId`s parse ROM1's own unmodified per-type formula/walk to exactly 0 residue on all 83 ROM2 maps and both `formatVersion` groups (1300, 1600) uniformly, or fail on all 83 maps and both groups… | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-019 | The condition that selects ROM2's 28-byte extended type-4 (placed object) record layout over the 20-byte base layout, read directly from `allods2.exe`'s decompiled dispatcher `R2.0002` (case 4), is `kind==0x21 ||…` | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-020 | `allods2.exe`'s ALM record dispatcher `R2.0002` enforces a widened, code-verified acceptance ceiling, not merely a corpus that happens to ship higher values. | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-021 | `R2.0002` (`allods2.exe`, entry `R2.0002`) is ROM2's own record-body dispatcher, the counterpart of ROM1's `R0478` (`ALM-META-024`). | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-022 | ROM2-only types 10, 11 and 12's element counts are read from the type-0 (metadata) record's own version-gated fields, never from the type's own record payload. | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-023 | ROM2 extended the dispatcher's existing generic per-element read/allocate mechanism to serve its 3 new types; it did not add a second dispatch mechanism. | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-024 | `R2.0002`'s case 0, case 6, case 8 and case 9 each gate one or more field reads on `formatVersion` thresholds, read directly from the decompiled body in source order: case 0 has four (`0x47d`/1149, `0x4cd`/1229… | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-025 | No literal reference to `templates.bin` (or two plausible variant spellings) exists anywhere in `allods2.exe`'s image, by direct string search — confirmed, and its bound corrected, below. | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-026 | `allods2.exe` contains its own copy of the exact Data.bin loading shape `DAT-LOC-001` already publishes for ROM1's `rom.exe`: a guard flag at a fixed offset (`+0xf0`) on a long-lived object makes the load run at most once… | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-027 | ROM2's own corpus does not exceed ROM1's own corpus maximum in `W`, `H`, or `#objects`/`#units` (record-0 metadata, `ALM-META-024`/`ALM-META-025`), and exceeds it in `#players` — but within an already-published structural… | High | ✔ promoted (partially retracted) | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-029 | ROM1's DAT-GRAM-003 grammar, corrected only in group C's raw byte block (widened from ROM1's published 10 bytes to 14), replayed corpus-wide reaches EXACT 0 residue on both ROM2 `data.bin` files (client 149 971 B, server 152… | High | ✔ promoted | [EXP-2004](../experiments/EXP-2004-rom2-databin-schema/) |
| R2-ASSET-030 | The full 2556-byte client/server `data.bin` size delta (`R2-ASSET-012`) is located exactly, not merely bounded. | High | ✔ promoted | [EXP-2004](../experiments/EXP-2004-rom2-databin-schema/) |
| R2-ASSET-031 | ROM2's group title arrays, read via the corpus-wide-clean grammar (`R2-ASSET-029`) and diffed column-for-column against ROM1's own published `DAT-SCHEMA-007` list (`databin-slots.csv`), are IDENTICAL on 6 of 8 groups (A… | High | ✔ promoted | [EXP-2004](../experiments/EXP-2004-rom2-databin-schema/) |
| R2-ASSET-032 | Replaying the corpus-wide-clean grammar (`R2-ASSET-029`) against `templates.bin` reaches the same divergence point `R2-ASSET-005` already reported under the unmodified grammar — group A, `Shapes` collection, entry 112… | Medium | ✔ promoted | [EXP-2004](../experiments/EXP-2004-rom2-databin-schema/) |
| R2-ASSET-033 | Classifying `allods2.exe`'s Data.bin-reading population against `a2server.exe`'s own function census (`EXP-2001`'s already-published `allods2.tsv`/`a2server.tsv`, `tools/enginematch`'s strict… | High | ✔ promoted | [EXP-2004](../experiments/EXP-2004-rom2-databin-schema/) |

### R2-ASSET-001

ROM2's `.res` container is byte-identical to ROM1's `&YA1` container (`RES-MAGIC-001`, `RES-HDR-002`, `RES-ACCEPT-031`, cross-reference only). All 11 top-level `.res`/`.RES` files (`MUSIC.RES`, `graphics.res`, `main.res`, `movies.res`, `patch.res`, `scenario.res`, `sfx.res`, `speech.res`, `video.res`, `world.res`, `world_srv.res` — the population is **11**, not the 10 this experiment's own brief named; `MUSIC.RES` was the omission) open with `internal/rom.OpenArchive`, UNMODIFIED for this experiment, and every one carries magic `26 59 41 31` ("&YA1"). Applying RES-ACCEPT-031's stricter, non-engine-required corpus invariant (every file node's payload range inside `[24,regOff)`, exact contiguous tiling, printable-ASCII names) finds **0 violations across all 11 containers** — every file payload range tiles `[24,regOff)` exactly. The entry-extension vocabulary inside these containers is **not** a subset of ROM1's own: 12 of ROM2's 14 non-empty extension categories match ROM1's exactly (`.16`, `.16a`, `.256`, `.alm`, `.bin`, `.bmp`, `.dat`, `.pal`, `.reg`, `.txt`, `.wav`, plus the extensionless entry both roots have), but `.pkt` (2 files — `data/itemname.pkt` in `world.res` and in `world_srv.res`, `R2-ASSET-012`) and `.smk` (8 files, all in `video.res` — Smacker video) appear in **neither** ROM1 EN's nor ROM1 RU's own container population, measured by running this same extension census, unmodified, rooted at each ROM1 install as a positive control (`res-rom1-en/res-extensions-total.tsv`, `res-rom1-ru/res-extensions-total.tsv`: 12 categories each, identical set, no `.pkt`/`.smk` row in either)

**Confidence.** High (the discriminating fact is that the reader is the exact ROM1 package, unmodified, and the STRICT corpus check — which the engine itself does not require and which has no per-file tolerance — also passes with zero exceptions on all 11 files; a variant `.res` layout that merely LOOKS like `&YA1` at the header would not also satisfy exact byte-for-byte payload tiling on every one of 11 differently-sized archives). The extension-novelty finding is a direct census count run identically on both games, ruling out "ROM2's container reader/extension-splitting logic itself invents these two names" as an explanation

### R2-ASSET-002

ROM2's `.alm` file header uses ROM1's exact field layout (`ALM-FRAME-031`, `ALM-HDR-001`, cross-reference only), not byte-identical VALUES. Over the full population — 37 root `.alm` files + 46 `M7R\0`-magic payloads nested in `scenario.res` (named `<n>.alm`) = 83 maps, not the ~30 this experiment's own brief assumed — every one carries magic `4D 37 52 00` ("M7R\0") and header field `hdrLen=20` (83/83, both fields uniform with 0 exceptions), and every one's first record reads as type 0 with a payload `>= 0x28` bytes, from which `W`/`H` are read at `+0x00`/`+0x04` per ALM-HDR-001. Replaying ALM-HDR-001's own corpus arithmetic identity `dataSize = 4*W*H + 72` (an independent cross-check spanning 3 separately-read fields, unaffected by the correction below since `W`/`H` are read before it applies) holds on **83/83**, 0 exceptions. Record 0's own declared `payloadSize` is a corpus-wide constant **644** (ROM1's own constant: **632**, `ALM-SEC-004`) — a genuine 12-byte content difference — and separately, record 0's TRUE span (its header's end to record 1's header's start) is **660** bytes, 16 more than even the declared 644: a `type0PayloadOverhang` this survey had to add to its own cursor arithmetic to reach record 1 correctly. Replaying that same +16 correction against ROM1's own two preserved roots, as a falsification check, does NOT also produce a clean walk (`alm-control.txt`: ROM1 EN 0/38 clean, ROM1 RU 0/34 clean, both with `firstDivergent` at record 1 uniformly) — ruling out "this is a universal reader forgiveness that would clean up any file, ROM1 included." Before this correction, this survey's own tool carried the uncorrected 644-byte advance forward into every later record read, desyncing the cursor; the far-out `typeId` values that produced (e.g. `343807103`, `4292870144`) were an artifact of that desync, not a ROM2 grammar fact (`R2-ASSET-003`, Attempts to falsify)

**Confidence.** High (three independently-positioned fields — `dataSize` from the file header, `W` and `H` from record 0's payload — satisfying one fixed arithmetic identity on 100% of an 83-file corpus is not plausible if any of the three were being read from the wrong offset; this rules out "the offsets merely look right but are shifted". The 644-vs-632 and +16-overhang facts are direct, uniform, corpus-wide reads with 0 variance; the overhang specifically is confirmed by a targeted falsification control that rules out "any additive fudge factor would work here," not merely asserted by trial against ROM2 alone)

### R2-ASSET-003

Past record 0, ROM2's `.alm` record chain uses ROM1's exact ALM-FRAME-031 record framing (the 20-byte `[tag][hdrLen][payloadSize][typeId][f32]` header, `payloadSize` used unmodified to advance to the next record) — **all 83 maps reach exact EOF with 0 residue**, once record 0's own 16-byte overhang (`R2-ASSET-002`) is applied to the cursor. The originally-reported divergence within the first four records, and the far-out `typeId` values read there, were this survey's own probe carrying record 0's undercounted declared size forward — a cursor desync, not a ROM2 grammar change (Attempts to falsify). Two corpus-wide CONTENT differences remain, distinct from framing: `recordCount` (file header `+0x0C`) is **13** on all 83 maps (ROM1: constant 10, ALM-HDR-001 — 3 more record types recorded, not 3 more framing fields); `formatVersion` (`+0x10`) is **1300** on 45 of the 46 `scenario.res`-nested maps and **1600** on the 37 root maps plus that one exception (ROM1: constant 990). All 13 `typeId` values `0..12` occur exactly once per map, corpus-wide, in one canonical order — `[0,1,2,3,5,11,4,9,8,6,7,10,12]`, 83/83 uniform. Replaying ROM1's own per-type payload-size formulas (`ALM-SEC-004`, `ALM-GRID-032`, `ALM-GRP-020`, `ALM-UNIT-018`, transcribed unmodified) against each record's declared `payloadSize`: types 1 (`2*W*H`), 2/3 (`W*H`), 5 (`76*nPlayers`) match exactly on **83/83**, 0 exceptions. Type 6 (ROM1's own formula, ALM-UNIT-018: `70*nUnits`) matches on **0/83** — but on all 83, the declared value instead equals exactly `48*nUnits` for the SAME `nUnits` read from record 0's own metadata (`alm-records.tsv`: `predicted = 70*nUnits` and `got = 48*nUnits` hold simultaneously and exactly on every one of the 83 rows checked) — a clean, corpus-uniform alternate constant, not scatter around 70. What accounts for 48 bytes/unit rather than ROM1's 70 was not decoded by this initial survey; that byte-width question is subsequently resolved by R2-ASSET-024, without assigning new field meanings. Types 4, 7, 8, 9, 10, 12 have no published ROM1 per-type formula to replay (581 of 1079 total records read corpus-wide are outside `{0,1,2,3,5,6}`)

**Confidence.** High for framing-past-record-0 being identical: reaching exact 0-residue EOF on 100% of an 83-file corpus, for all 13 record types, using ONLY ROM1's unmodified field layout plus one documented, falsifiably-scoped correction, is not plausible if the record BODY grammar had actually changed shape (an extra field or different stride would desync the chain somewhere across 83 varied files) — this rules out "the record framing itself changed," the alternative the original, uncorrected version of this row could not exclude. High also for `recordCount`/`formatVersion`/`typeId`-population/ordering (direct, uniform, 0-variance reads) and for the type-6 formula-mismatch FACT itself (an exact, corpus-uniform arithmetic identity, not a scatter plot); no confidence is asserted for the type-6 mismatch's own CAUSE, which this survey does not decode

### R2-ASSET-004

`world.res:data/data.bin` and `world_srv.res:data/data.bin` (149 971 and 152 527 bytes) both parse under DAT-GRAM-003's wire grammar (cross-reference only) identically to ROM1 EN through group A's title array: the leading **102 bytes** are byte-identical to ROM1 EN's own `world.res:data/data.bin`, confirmed two independent ways — a raw byte-for-byte common-prefix comparison from offset 0, and group A's own parsed title-array span landing on the same boundary, `[0,102)`, in both files independently. (An earlier review pass cited 32 bytes; that number does not reproduce — 102 is this survey's own, independently re-derived figure, matching the title-array span exactly.) Past that: `Materials` is fully consumed at count 16 (ROM1: 16, exact match) and `Shapes` at count 7 (ROM1: 5 — 2 additional rows, the one difference inside group A); group B's title array is 30 strings (ROM1: 30, exact match) with `Magic` fully consumed at count 50 (ROM1: 50, exact match); group C's title array is 18 strings (ROM1: 18, exact match) and the declared `Armors` count reads 31 (ROM1: 31, exact match) — three independent exact count matches beyond the shared prefix. Only 3 of 31 `Armors` entries are then consumed before a length-prefixed field read demands more bytes than either file has left, at an offset exactly equal to that file's own total size for BOTH files (149 971 and 152 527) — a fact about running out of bytes (any bounds-checked reader that overruns necessarily reports its failure at the file's own remaining-byte count), not independent evidence of anything past that. Per-entry byte spans for the 3 consumed `Armors` entries (`databin-walk.txt`): entry 0 spans 87 bytes (comparable to ROM1's own Armors-collection average of ~91.5 B/entry, `(11231−8486)/30`); entry 1 spans 15 bytes (already short of any plausible 18-field armor record); entry 2 spans 5135 bytes (about 56x the ROM1 average). The visible break in plausible per-entry sizing is at entry 1, not entry 0 or "after 3" — entry 0's own SPAN is not itself anomalous, though this survey reads no entry CONTENT, so whether entry 0 is already invalid by content cannot be determined from a span measurement alone. This initial replay's stopping point is superseded on reach by R2-ASSET-029, R2-ASSET-030 and R2-ASSET-031; the original measured spans and their confidence qualifications remain unchanged

**Confidence.** High for the schema/prefix match: rules out "this is a coincidentally similar but structurally unrelated container" — three independent exact count matches (16, 50, 31) plus a byte-identical 102-byte prefix (which, per the redacted title-length list, has the same 11 title lengths as ROM1's own visible title strings) would not survive an unrelated format by chance. The divergence-lands-at-EOF fact is reported but is NOT part of this rationale — it is true of any overrun read by construction and does not discriminate between hypotheses. Medium for the entry-1-onward characterization: the byte-span progression (87, 15, 5135) is directly observed and reproducible, but content is never read, so entry 0's own validity is Unknown, not asserted either way

### R2-ASSET-005

The ROM2 root's `templates.bin` (61 802 bytes) is a distinct file from `world.res`'s own `data/data.bin` — different size, different content (`R2-ASSET-004`) — but whether it shares DAT-GRAM-003's own wire grammar at all is **Unknown**, not "Data.bin-shaped": this survey's replay of that exact grammar (`u16` title count, then `u8`/`0xff+u16`-length-prefixed strings) reads a declared count of 54 and consumes to byte 1698 with 0 read errors, then its `Shapes` collection (declared count 2560) runs out of bytes at entry 112, offset 35 004 of 61 802. But the title-array READING itself is not trusted as templates.bin's real grammar: the first two of the 54 "title" lengths read as 0 (`alm`-style formula cross-checks find no such anomaly elsewhere in this survey; a real title string being 0 bytes twice in a row is implausible), and an independent, content-free alternate reading of the SAME head bytes fits at least as well and needs no 0-length titles: a `u32` count (54 — matching the `u16` reading's own value, since a small `u32`'s upper 16 bits are zero either way) followed by a `u32` length field (11) followed by 11 bytes that ARE printable ASCII (`databin-walk.txt`'s "alternate raw head reading"). This survey does not decode which reading, if either, is templates.bin's real grammar — so whether its title/collection LAYOUT is DAT-GRAM-003-shaped at all is left open. This row records only: the file is distinct from `world.res`'s data.bin; the DAT-GRAM-003 replay's own byte offset of divergence (35 004); and the two candidate head readings, both directly observed, neither asserted as correct

**Confidence.** **Unknown** for templates.bin's own wire grammar — two candidate head readings (`u16`-count/`u8`-length-prefix vs `u32`-count/`u32`-length-prefix) both fit the observed head bytes with 0 read errors, and this survey has no independent field cross-check (no `W`/`H`-style arithmetic identity, unlike `R2-ASSET-002`) to prefer one over the other. High only for the directly observed, reading-independent facts themselves: the file's distinctness from `world.res`'s data.bin, the DAT-GRAM-003 replay's own divergence offset, and the literal output of both head readings

### R2-ASSET-006

ROM2's nested inline registries are byte-identical to ROM1's REG-FMT-031 grammar (cross-reference only), found the same way RES-SCOPE-015 found ROM1's — by `&YA1` magic inside a container's file payloads, not by a `.reg` name. Across all 11 `.res` containers, **19** such payloads exist (`graphics.res` 5, `scenario.res` 1, `sfx.res` 1, `video.res` 8, `world.res` 2, `world_srv.res` 2) and **19/19 parse cleanly**: header fields read, `recordCount` 32-byte records read, and the payload tiles `0x18 + recordCount*32 + 4 + poolLen` exactly to the payload's own end, 0 exceptions. `world.res` and `world_srv.res` each nest the identical pair `data/ai.reg` (156 B) and `data/map.reg` (1020 B), matching `split.go`'s independent byte-for-byte size comparison (`R2-ASSET-012`)

**Confidence.** High (100% of the full population — an exhaustive by-magic search, not a sample — satisfies an exact-tiling grammar with a 4-field header, a record array, and a length-prefixed pool all agreeing to the byte; this is the same discriminating strength RES-ACCEPT-031 and REG-FMT-031 already carry on ROM1, replayed here without modification)

### R2-ASSET-007

ROM2's `.16a` and `.256` sprite containers are byte-identical in framing to ROM1's (`SPR16A-STRUCT-001`, `SPR256-STRUCT-001`, cross-reference only), measured over the FULL population found by the extension census, not a sample: `.16a` **623/623** (100%) walk with exact tiling from the trailer-derived frame count through every frame's `[w,h,dataSize,data]` header to the trailer boundary, 0 exceptions (ROM1 EN and RU: **542/542** each, same code, positive control). `.256` **1916/1929** (99.3%) walk exactly clean; of the remainder, 8 are 0-byte stub entries (not evaluable — e.g. `cursors/cast.256`, `cursors/defend.256`) and **5 show residue** (every frame internally consistent, but bytes remain between the last frame and the trailer). All 5 residue `.256` files are **sha256-identical** to the identically-named node in ROM1 RU (`spr-residue-identity.txt`, full population, not sampled: `interface/myitem.256`, `interface/shopframe.256`, `interface/shopitem.256`, `projectiles/goblin/arrow.256`, `projectiles/goblin/arrowb.256`) — the already-published `SPR256-EXC-017` Bucket-B pattern (a bracketed secondary section appended after the frame data, `0x80000001…0x80000001`, resolved by `SPR256-EXC-020` as loader-inert) reproduced byte for byte, not a ROM2-specific extension

**Confidence.** High for the framing match itself (full-population, not sample-limited, with the trailer/palette-flag/frame-tiling logic exercised on nearly 2500 files, a clean ROM1 positive control on `.16a`, and only the two named, separately-characterized exceptions on `.256`). High also for the residue files' identity: a full-population sha256 comparison against ROM1 RU, 0 exceptions on 5/5, directly rules out "this is new ROM2 content that happens to look similar" — the only alternative a shape/size match alone could not exclude

### R2-ASSET-008

ROM2's `.pal` palettes match one of ROM1's two known shapes (BMP colour table at a fixed seek, or a flat 16×1024-byte block — `formats/pal/format.md`, cross-reference only) on **159/159** (100%) of the full population found by the extension census — 0 files matching neither shape. Of the 156 matching the BMP-colour-table shape, all **156/156** also pass the strict structural test `bfOffBits == 0x436 && biBitCount == 8` (0 partial matches — no file has the outer BMP shape without also having the exact expected header field values); the other 3 match the flat 16×1024-byte shape. ROM1 EN and RU each show the identical breakdown proportionally on their own smaller population, run through the same code as a positive control: **82/82** total, **79/79** shape-A-and-strict, 3/3 shape B, 0 matching neither, on BOTH roots (`spr-rom1-en/spr-summary.txt`, `spr-rom1-ru/spr-summary.txt`)

**Confidence.** High (full population, binary shape test, 0 exceptions — the only way this could mislead is if a THIRD shape happened to alias one of the two structural tests, which a 100%-clean population of 159 varied-size files makes implausible without at least one visible outlier). The strict field-value test strengthens this further and is itself controlled: it also passes 100% on both ROM1 roots' own smaller population under the identical code, ruling out "the strict test is loose enough to pass almost any BMP-shaped blob"

### R2-ASSET-009

ROM2's `.16` bitmap font containers (the SPR16A-TRLR-012 no-palette variant, cross-reference only) are a complete population of 3 files, and 2 of the 3 do not tile exactly to the trailer despite every individual frame validating. All 3 have `hasPalette=false` (matching ROM1's own "a bare `.16` never carries a palette" rule) and the trailer-declared frame count is walked in full for every file with no single frame's `[w,h,dataSize]` header ever exceeding file bounds: `font3.16` (64 frames) consumes exactly to the trailer, 0 residue; `font1.16` (224 frames) leaves 17 716 bytes between the last frame and the trailer; `font2.16` (224 frames) leaves 32 bytes. Both residue files, `font1.16` and `font2.16`, are **sha256-identical** to the identically-named node in ROM1 RU (`spr-residue-identity.txt`, full population). This is the already-published `SPR16A-FONT-014`/`SPR16A-FONT-021` finding — older in-place-overwrite build layers of the same font, each boundary run the byte-suffix of a record cut by the next build's own end — reproduced byte for byte in ROM2, not a ROM2-specific phenomenon. What the residue bytes hold is not re-decoded here — that decoding is `SPR16A-FONT-014`/`-021`'s own, cross-referenced, not repeated

**Confidence.** High — a full-population sha256 identity against ROM1 RU (not a shape or size match) is a byte-for-byte proof this is the SAME file, not merely a similarly-shaped one, which resolves what an earlier pass left as an n=2 correlation ("residue correlates with 224 frames") into a direct identity check: these bytes are not new ROM2 content at all

### R2-ASSET-010

ROM2's `main.res` still carries a `text/` directory of CRLF-delimited string-table files (TEXT-STRTAB-023, cross-reference only): all 15 of TEXT-STRTAB-023's named files (`main.txt` through `credits.txt`) are present under `main.res:text/`, and `patch.res:patch.txt` is present, matching ROM1's own load-order population exactly by name. The directory is also substantially larger: **67** files live under `main.res:text/` (**52** beyond the 15 ROM1 names found there, not 54 beyond 14 — recount from `text-strtab.tsv`'s own per-container rows), plus the 1 file under `patch.res` (68 corpus-wide, matching `res-extensions-total.tsv`'s `.txt` row). Of the 52 beyond ROM1: **46** are `missionNN.txt` campaign files, `mission10.txt` through `mission110.txt` (not every number in between is present, 46 of a possible 101), and **6** are new non-mission names: `globalmap.txt`, `help.txt`, `itemserv.txt`, `quest.txt`, `town.txt`, `docs/1.txt`

**Confidence.** High (a presence/name check over a fully enumerated directory listing, not a parse-dependent measurement — every one of the 15 names either is or is not present, and all 15 are, alongside a directly counted, exhaustively-enumerated 52 additional files)

### R2-ASSET-011

ROM2's text-shaped nodes (`.txt`/`.ini`/`.lst`/extensionless payloads in any of the 11 `.res` containers — 69 nodes, 513 002 bytes: 68 `.txt` files corpus-wide plus 1 extensionless node, ROM2 having 0 `.ini`/`.lst` files anywhere; a broader, container-wide population than `R2-ASSET-010`'s text/-path-scoped 68, not the same count by coincidence of definition) do not stay inside TEXT-CONV-001's two source blocks `{0x80..0xAF, 0xE0..0xEF}` (cross-reference only) the way ROM1's own text does. Measured fresh with this survey's own tool and its own population definition, as a same-tool positive control (not cited from TEXT-FIT2-013's differently-scoped 87293/87293 figure, which also folds in `.reg` and `data.bin` embedded strings this survey counts separately): ROM1 EN scores 1/1 high bytes in-domain (100.0000%, 425 nodes) — with only 1 high byte total, this control carries no discriminating information and is reported for completeness only, not as support. ROM1 RU scores 87446/87446 (100.0000%, 436 nodes, corroborating TEXT-FIT2-013's conclusion independently): of those, 62933 fall in `0x80..0xAF` and 24513 in `0xE0..0xEF` — RU's in-domain share spans BOTH source blocks. ROM2 scores only 210405/317058 (**66.3617%**) in-domain, and its in-domain share does NOT span both blocks the way RU's does: **0 (zero)** of ROM2's 317058 high bytes fall in `0x80..0xAF` at all (against RU's 62933); ROM2's entire in-domain total of 210405 is inside the single shared `0xE0..0xEF` window. The remaining 106653 out-of-domain bytes (33.6%) populate 47 of the 64 values in the two blocks TEXT-CONV-001's converter never touches, `0xB0..0xDF` and `0xF0..0xFF`, and the populated values are not spread evenly across that range: all 16 of `0xF0..0xFF` are used, 31 of 32 of `0xC0..0xDF` are used (only `0xda` is absent), and **none** of the 16 values `0xB0..0xBF` appear at all — see `evidence/text/text-domain-census.txt` for the exact per-value histogram

**Confidence.** High, resting on the RU control alone (n=87446, not the uninformative EN n=1): a same-tool, same-population-definition comparison where ROM1 RU measures a clean, exact 100.0000% under this survey's own code rules out "the population or the domain test itself is biased toward a fit" as the explanation for ROM2's 66.3617% — the identical instrument finds a real, large, structured gap only on ROM2. The 0-of-317058-in-`0x80..0xAF` finding is a direct count, not an inference

### R2-ASSET-012

`world.res` and `world_srv.res` hold the identical 5 top-level `data/` entry names (`ai.reg`, `data.bin`, `itemname.bin`, `itemname.pkt`, `map.reg` — 0 client-only, 0 server-only entries), diffed with the unmodified ROM1 archive reader on both containers. 4 of the 5 are byte-identical in size (`ai.reg` 156, `itemname.bin` 982, `itemname.pkt` 7763, `map.reg` 1020); `data/data.bin` differs — client 149 971 bytes, server 152 527 bytes, the server copy 2556 bytes larger — consistent with `R2-ASSET-004`'s finding that both copies share groups A/B exactly and diverge only from group C onward, where the extra server-side bytes could live

**Confidence.** High (a direct entry-name-and-size diff over a fully enumerated 5-entry population on both sides; no parsing beyond what `R2-ASSET-001` already establishes for the container reader itself)

### R2-ASSET-017

ROM2's per-map `typeId` set and record order strictly contain ROM1's own (Q1; `R2-ASSET-003` already publishes ROM2's own flat order and set — this row adds the cross-game RELATION and one population exception). ROM1's own corpus (this survey's own instrument, run fresh against both preserved roots for the first time under one tool): all 38 EN maps and 33 of 34 RU maps carry `typeId` set `{0..9}` in one canonical order, `[0,1,2,3,5,4,9,8,6,7]`, uniform with 0 exceptions in that population. The one RU exception is `Horror.alm`: its own file header declares `recordCount=4` (types `{0,1,2,3}` only — no players/objects/units/triggers/loot/type-9 records at all), confirmed by a direct raw header read (`magic M7R\0, hdrLen 20, dataSize 262216, recordCount 4, formatVersion 990`, file size 262876 bytes), against the EN root's own same-named `Horror.alm` carrying the full 10-record set at 409386 bytes. **Correction, finding 7:** a byte comparison over the shared 262876-byte prefix (`horror-diff.txt`) finds exactly 2 differing offsets — file header `+0x0C` (`recordCount`, EN 10 → RU 4, the exception already published above) and payload offset 152 (record 0's own `+0x70`, value not reported per the install-byte boundary) — with every other byte identical, including `dataSize` and `formatVersion`. RU is 146510 bytes shorter than EN with no other divergence in what the two files share: the RU file is EN's own file with `recordCount` rewritten and the trailing bytes for the now-uncounted records cut off, not an independently authored 4-record map. This is still a ROM1-corpus population fact, not a ROM2 one, but it is one truncated copy of an EN map, not a second authored map — "38/38 EN and 33/34 RU" counts files that pass the typeId-set/order test, not 34 independently authored RU maps

**Confidence.** High for the order-containment relation (a full, 0-exception population comparison on both sides, not a sample; the alternative this rules out is H1-disjoint-in-part — a shared `typeId` number reused for unrelated ROM2 content would not also preserve ROM1's own relative record ORDER around it, and it does, on every file). High for the truncation reading: a 2-byte total divergence over a 262876-byte shared prefix, with the file's own declared `recordCount` explaining the size difference exactly, is not plausible as an independent re-encode producing the same byte stream by coincidence

### R2-ASSET-018

Nine of the ten shared `typeId`s parse ROM1's own unmodified per-type formula/walk to exactly 0 residue on all 83 ROM2 maps and both `formatVersion` groups (1300, 1600) uniformly, or fail on all 83 maps and both groups uniformly (Q2's H2-generation vs H2-format test). Types 1, 2, 3, 5, 7, 8, 9: pass 45/45 at `formatVersion` 1300 and 38/38 at 1600 (0 exceptions, both groups) — H2-identical. Types 0 and 6: fail 0/45 and 0/38 (both groups) against ROM1's own unmodified formula — already known to instead match a corpus-uniform alternate constant (`R2-ASSET-002`'s +16 record-0 overhang, `R2-ASSET-003`'s `48*nUnits` for type 6) — H2-extends, and the 0/45-and-0/38 split confirms this is NOT formatVersion-gated either. Type 4 is the tenth, and does not fit the same two-way split. **Correction, findings 5 and 6:** replaying three distinct rules independently, cross-tabbed by `formatVersion` (`type4-rule-crosstab.txt`) — ROM1's own unmodified literal `kind==0x21` test (`type4Walk`), a low-byte mask `(kind&0xff)==0x21`, and the dispatcher's own ground-truth condition `kind==0x21||(kind&0x1000000)!=0` (`R2-ASSET-019`) — only the FIRST of these three shows a `formatVersion` split: 45/45 at 1300, 19/38 at 1600 (this row's earlier draft misattributed these two numbers to "a first, empirically-derived low-byte-mask rule"; they are ROM1's own literal rule's numbers). The low-byte mask and the ground-truth condition both reach 83/83 at both groups — neither one splits by `formatVersion` on this corpus. The 19 files where ROM1's literal rule fails are exactly the 19 files carrying `kind` value `0x1000021` (`0x21` with bit 24 set), and all 19 are at `formatVersion` 1600; 0 of the 45 `formatVersion`-1300 files carry that value or any other kind sharing the `0x21` low byte (`type4-kind-partition.txt`). So type 4's own correct structural rule (the dispatcher's ground truth) is itself uniform across both `formatVersion` groups — H2-generation is excluded for the RULE, matching the other nine types — but the corpus CONTENT that exercises the rule's extension arm is completely `formatVersion`-partitioned: `0x1000021` is a `formatVersion`-1600-only kind value in this corpus, full stop, which the row's earlier flat "no shared `typeId` shows a split that tracks `formatVersion`" wording denied without stating

**Confidence.** High for the H2-generation-excluded-per-correct-rule finding: a real formatVersion-gated shape change in a type's own correct rule would show as a pass/fail split under that rule, and none of the ten types shows one under its own correct rule — verified for type 4 specifically by testing three distinct candidate rules independently rather than asserting the chosen one is correct. High, as a directly counted fact, for the `0x1000021`/`formatVersion`-1600 partition (19 of 19 exception files carry it, 0 of 45 non-1600 files do) — this is a population statement, not an inference, and is reported separately from the rule-uniformity finding, not folded into it

### R2-ASSET-019

The condition that selects ROM2's 28-byte extended type-4 (placed object) record layout over the 20-byte base layout, read directly from `allods2.exe`'s decompiled dispatcher `R2.0002` (case 4), is `kind==0x21 || (kind&0x1000000)!=0` — `kind` at the same record `+0x08` offset ROM1's own `ALM-CLS-036`/`ALM-OBJ-034` already document. `ALM-CLS-036` already publishes `kind==0x21` as ROM1's own extension trigger; ROM2's dispatcher extends that exact-match test with an independent bit-24 flag that also triggers the same 28-byte layout regardless of the low byte. Replaying this exact condition reaches 0 residue on 83/83 ROM2 records (`structure-summary.txt`). **Correction, finding 5:** the stated discriminator for this reading was wrong and is replaced, not merely softened. This corpus's own type-4 `kind` vocabulary is exactly two values, `0x21` and `0x1000021` (`type4-kind-census.tsv`), so a low-byte mask `(kind&0xff)==0x21` classifies every one of the 83 files identically to this ground-truth condition — both reach 83/83 at both `formatVersion` groups (`type4-rule-crosstab.txt`); the corpus cannot discriminate the two. What the corpus DOES separate this condition from is ROM1's own unmodified literal `kind==0x21` exact-match test, which reaches only 64/83 (`R2-ASSET-018`) — that is the one rule whose pass rate differs by `formatVersion`, not the low-byte mask

**Confidence.** High for the condition itself — read directly from the decompiled dispatcher body (cross-checked against the function's own raw, byte-free disassembly, not the decompiled C alone), and its own replay reaches 0 residue on the full 83-map/all-type-4-records population, which a materially wrong condition could not achieve by chance. Explicitly not claiming more: this experiment's own corpus cannot separate this exact condition from a simpler low-byte mask, since no observed record carries a kind that would make the two disagree (one sharing the `0x21` low byte with bit 24 clear, or vice versa) — the code read, not a corpus discrimination, is this row's basis

### R2-ASSET-020

`allods2.exe`'s ALM record dispatcher `R2.0002` enforces a widened, code-verified acceptance ceiling, not merely a corpus that happens to ship higher values. Before its per-record loop it runs `strcmp(<header>, [L2.00002])` where the global at `L2.00002` holds `"M7R"` (the printable prefix of ROM1's own `M7R\0` magic, `ALM-HDR-001`) — mismatch fails with error code 2; `recordCount<3` fails with error code 3, the identical threshold ROM1's own `ALM-META-024` gate uses (`>=3` accepted); `formatVersion > 0x640` (1600 decimal) fails with error code 4 — ROM1's own `ALM-META-024` gate on `rom.exe`'s equivalent function is `<=1001` (`0x3e9`). This reads the ENGINE'S OWN compiled boundary, not the corpus's observed values (`R2-ASSET-003` already reports those: 1300, 1600) — the ceiling sits exactly at the higher of the two shipped values, with no headroom measured above 1600

**Confidence.** High for the three gate facts themselves (magic substring, `recordCount` threshold, `formatVersion` ceiling), each read directly from one decompiled acceptance-gate block, not inferred from what ships; this discriminates "the engine's own limit widened" from "the shipped corpus just happens to use higher numbers than ROM1's own corpus, but an unwidened engine would still reject them" — the latter is excluded because the ceiling is read from the CODE'S own comparison, not backed out from what files exist

### R2-ASSET-021

`R2.0002` (`allods2.exe`, entry `R2.0002`) is ROM2's own record-body dispatcher, the counterpart of ROM1's `R0478` (`ALM-META-024`). Located by full-image cross-reference from EXP-2001's own already-`identical`-classified callee pair (`classification.tsv`: ROM1 `R0467`/`R0466` ↔ ROM2 `R2.0003`/`L2.00003`, ROM1's own header/record readers): a full xref scan (`callto:R2.0003`, `callto:L2.00003`) finds exactly 1 caller each, both `R2.0002` (call sites `L2.00004`, `L2.00005`). Its body is an outer acceptance gate (`R2-ASSET-020`), then a loop bounded by the header's own record count, each iteration reading one record header then a `switch` on its type field with exactly 13 cases — `0,1,2,3,4,5,6,7,8,9,10,0xb,0xc` — plus one `default` case that calls a single generic vtable slot (`+0x30`, distinct from every named case's own `+0x3c`) with no type-specific logic: the dispatcher does not hard-reject an out-of-range type id, it simply has no dedicated case for one. Independently, `tools/r2alm`'s full corpus enumeration finds every one of the 83 files uses exactly `typeId` `0..12` and no others (`population-table.tsv`) — the code's 13 dedicated cases and the corpus's 13 observed types match exactly

**Confidence.** High — a sole-caller convergence from TWO independently-already-verified-identical functions onto exactly one address rules out "one of several plausible dispatcher candidates was picked"; the case-count/corpus-population match (13 dedicated cases, 13 observed types, both independently derived) is a second, unrelated-instrument cross-check of the same identification

### R2-ASSET-022

ROM2-only types 10, 11 and 12's element counts are read from the type-0 (metadata) record's own version-gated fields, never from the type's own record payload. Read directly from `R2.0002`'s decompiled body: case 10's loop bound is the `u32` field case 0 reads once `formatVersion>=0x47e` (1150); case `0xb` (11)'s three loop bounds are the three `u32` fields case 0 reads once `formatVersion>=0x4ce` (1230); case `0xc` (12)'s loop bound is the `u32` field case 0 reads once `formatVersion>=0x514` (1300) — all three gated fields corpus-uniformly present, since every file's `formatVersion` (1300 or 1600, `R2-ASSET-003`) is at or above all three thresholds (`R2-ASSET-024`). **Correction, finding 4:** this row's own earlier contrast — "unlike every ROM1-documented type, whose own record supplies its own count" — is false and is withdrawn, not narrowed: case 0's own complete field-read table (`dispatcher-findings.md`) shows cases 4, 5, 6 and 8 also take their own loop bounds from type-0 fields read at payload `+0x20` (case 4's `nObjects`), `+0x1c` (case 5's `nPlayers`), `+0x24` (case 6's `nUnits`) and `+0x2c` (case 8's `nType8`) — all four unconditional, not gated. Types 1, 2 and 3 are likewise sized from type-0's own `W`/`H`. Record-0-supplies-the-count is the dispatcher's general mechanism, used by at least seven of the ten shared+new types read here (0 as source; 1, 2, 3, 4, 5, 6, 8 as consumers) plus 10, 11 and 12 — not a ROM2-only-type exception to an otherwise self-counted rule. `tools/r2alm/structure.go`'s own dispatch (`f.nObjects` to `type4Walk`, `f.nPlayers` to type 5, `f.nUnits` to type 6, `f.nType8` to `type8Walk`) and the already-promoted `ALM-META-025` (`+0x1c`=#type5, `+0x20`=#type4, `+0x24`=#type6, `+0x2c`=#type8, all ROM1) already said as much; the previous draft of this row did not check its own contrast against either. What IS distinctive about 10, 11 and 12: they are the only cases whose OWN count field is itself `formatVersion`-gated — zero below the threshold (case 0's own destinations are explicitly zero-initialized before the gates run), so an older-`formatVersion` file still runs the case but iterates zero times, where every other type-0-sourced count (4, 5, 6, 8, and case 0's own twelve unconditional fields) is populated on every accepted file regardless of `formatVersion`

**Confidence.** High — the count-field source (type-0's own payload, not the consuming type's own record) is read directly from the decompiled control flow linking case 0's field reads to case 10/11/12's own loop-bound locals, not inferred from any corpus size relationship. The corrected contrast (gated vs unconditional count fields, not type-0-sourced vs self-sourced) is read from the same case-0 field table, cross-checked against a second, independent source already in this repository (`ALM-META-025` for ROM1, `structure.go`'s own dispatch for this experiment)

### R2-ASSET-023

ROM2 extended the dispatcher's existing generic per-element read/allocate mechanism to serve its 3 new types; it did not add a second dispatch mechanism. Every per-element LOOP (shared type or ROM2-only), read directly from `R2.0002`'s decompiled body, follows the identical pattern: one call to an allocate/append accessor on the record's destination collection, then one or more reads through the same archive object's own vtable slot `+0x3c` into the newly allocated element — **correction, N5:** narrower than "every case": case 0 allocates nothing and reads into the record object's own fields, and cases 7, 9 and 12 each read one leading field (a count word for 7/9, a fixed 28-byte head for 12) before their first allocate call. Every case's per-element reads do route through the same `+0x3c` vtable slot with no differently-shaped reader anywhere in the switch. Type 11 (case `0xb`)'s body runs three separate loops, bounded respectively by the three counts case 0 reads at its own `0x4ce` gate (`R2-ASSET-022`), in the same order they are read. **Correction, N5 (related):** calling this a "three-counted-array record" — as an earlier draft of this row did, echoing ROM1's own type-7 name — inverts the analogy: type 7's three counts are read from its OWN payload (`ALM-TRIG-021`); type 11's three counts are in record 0, not in type 11's own payload at all, so type 11 is three flat arrays, not three *counted* arrays in type 7's sense. The two share a per-element shape (three back-to-back fixed-width loops), not the count-provenance property the retired name implied. **Correction, finding 2:** field widths, count-field offsets and destination collections for all three ROM2-only types ARE decompiler-legible and are published: type 10 is a flat array of 16-byte elements (its allocation call takes the literal element size 16) at `map+0x310`, count `meta+0x30`; type 11 is three back-to-back flat arrays of 12/84/12 bytes at `map+0x338`/`map+0x324`/`map+0x34c`, counts `meta+0x34`/`+0x38`/`+0x3c`; type 12 is one fixed 28-byte head at `map+0x374` (unconditional, one per record) then a flat array of 28-byte elements at `map+0x360`, count `meta+0x40` (`dispatcher-findings.md`, disassembly `L2.00006`-`L2.00007`). Replayed corpus-wide (`type0ext-identity.tsv`): **83/83 for type 10, 83/83 for type 11, 83/83 for type 12** — these are the first per-type size formulas published for these three types. Field-level SEMANTICS — what each byte of a type-10/11/12 element MEANS — are not decoded here and remain Unknown; the previous draft of this row conflated field width (now decoded) with field meaning (still Unknown) under one Unknown clause

**Confidence.** High for the mechanism-reuse structural fact itself (every per-element loop, shared and new, read as following the same two-step allocate-then-vtable-read pattern, with no branch to a differently-shaped reader anywhere in the switch), narrowed by N5 to the loops rather than every case's first instruction. High for the three per-type size formulas — each is a single pushed immediate at a fixed call site, replayed and matched 83/83 against every declared `payloadSize` in the corpus, not merely asserted from the disassembly read alone. Unknown, not guessed, for field content

### R2-ASSET-024

`R2.0002`'s case 0, case 6, case 8 and case 9 each gate one or more field reads on `formatVersion` thresholds, read directly from the decompiled body in source order: case 0 has four (`0x47d`/1149, `0x4cd`/1229, `0x513`/1299, `0x487`/1159, all "if `formatVersion>t` read one more field", no else) — one `u32` (4 B), three `u32` (12 B), one `u32` (4 B) and two `u32` (8 B) respectively, summing to exactly 28 bytes, which is exactly record 0's own TRUE span growth over ROM1 (632→660, `R2-ASSET-002`'s +16-over-declared overhang plus the +12 declared difference, `4+12+4+8=28=660-632`) — a directly-computable arithmetic check against already-published span numbers, not a new probe; this does not determine which specific bytes of that 28 the file's own declared `payloadSize` (644) counts and which land in the undeclared overhang, since that split is a property of the SAVE path this experiment's Ghidra pass (load path only, `R2.0002`) did not read; case 6 has six (`0x47e`/1150 and `0x3db`/987, both "if `formatVersion>t` read, else zero", same direction; `0x456`/1110, selects 2-byte vs 4-byte field width, wider above; and `0x44c`/1100, an outer branch taken only BELOW the threshold, with no else arm at all — the block it wraps, including two nested sub-thresholds `0x3b6`/950 and `0x3d8`/984, is emitted only when `formatVersion<0x44c` and is skipped entirely, not zero-filled, at or above 1100); case 8 and case 9 each gate one field on the same `0x3dd`/989 threshold (read only above it) — the identical numeric threshold `ALM-SACK-065` already publishes for ROM1's own type-8 20-byte-head gate, consistent with the same version-gated-growth lineage `ALM-META-024`/`ALM-SACK-065` document for ROM1, not a new mechanism (case 6's own `0x3db`/987 is close to but not identical to `ALM-META-024`'s listed `0x3da`/986; this experiment does not resolve whether that is a one-version difference or two coincident fields, and asserts neither). Every file in this corpus has `formatVersion` 1300 or 1600 (`R2-ASSET-003`) — above every threshold in this section except case 6's own `0x44c` (and its two nested sub-thresholds), so this corpus takes the newer/larger branch uniformly at every one of those gates, and takes the OTHER (absent) branch, uniformly, at `0x44c` and its nested sub-thresholds: that whole block is provably absent from every one of the 83 files' own type-6 records, not merely absent from some. **Correction, finding 1: the no-confidence clause below is closed.** Reading case 0's and case 6's raw disassembly straight through (not only their threshold gates) finds every field either case reads, in order (`dispatcher-findings.md`, "Case 0's complete field layout" and "Case 6's complete field layout" sections), and summing them closes both reconciliations this row previously declined to attempt. Case 0: twelve unconditional 4-byte fields (48) + the first three gated groups (4+12+4) + three further unconditional fields (64+4+4=72) + the fourth gated group at `0x488`/1160 (8) + one closing unconditional 512-byte field = **660 bytes at any `formatVersion` clearing all four gates** (every file in this corpus) and **632 bytes when none fire** (ROM1's own 990) — replayed per file against every declared type-0 `payloadSize` (`case0-width-check.tsv`): **155/155** (83 ROM2 + 38 ROM1 EN + 34 ROM1 RU, including RU `Horror.alm`'s own 4-record file; every file carries a type-0 record). 660 matches ROM2's own declared 644 plus exactly the already-published `rom2Type0Overhang` (16, `R2-ASSET-002`), now reproduced from the code rather than only fitted to the corpus; 632 matches ROM1's own declared constant (`ALM-SEC-004`) exactly, 0 slack. Case 6: eleven unconditional fields (36) + the `0x47e` and `0x3db` gated fields (4 each, fire uniformly on this corpus) + the `0x44c` block (0, corpus-uniformly absent, as already published) + the `0x456` width selector (4, wide branch) = **48 bytes/unit at this corpus's own `formatVersion`s**; at ROM1's 990, the `0x44c` block fires (28) and the `0x456` selector takes its narrow 2-byte form, giving `36+0+4+28+2=` **70 bytes/unit**: 990 clears the `0x3db`/987 gate, while the `0x47e`/1150 field is absent. Replayed per file against every declared type-6 `payloadSize`, both games (files with no type-6 record, i.e. RU `Horror.alm`, excluded from the denominator, not scored as failing) (`case6-field-check.tsv`): **154/154** (83 ROM2 at 48, 71 ROM1 at 70). The public-edition maintenance corrects three explanatory transcription errors above: the former ROM1 equation omitted the present `0x3db` field; the former case-0 prose counted its fourth gate twice as a fifth; and the former 155-file parenthetical mislabelled the population. The existing experiment summary, complete field table, canonical ALM equations and 155/154 result summaries already give these totals; no measurement or confidence changes. This closes `R2-ASSET-003`'s initial byte-width Unknown and `formats/rom2-alm/format.md`'s **Not yet surveyed** entry for the same question. What remains genuinely open is not the byte count but the semantics: why the writer's own declared `payloadSize` (644) omits exactly 16 of the 660 bytes the loader itself reads, and what (if anything) occupies that gap — this experiment reads the load path only and does not answer it; the declared/overhang split stays Unknown. **N6, checked, not narrowed:** case 0's own four `>`-form thresholds (decompiler-synthesized) each run exactly one below their own raw `>=`-form `JC` constant, confirmed on all four gates with no exception — a real candidate explanation for a one-apart pair, but case 6's OTHER two nested thresholds (`0x3b6`/950, `0x3d8`/984) are already published as EXACT matches to `ALM-META-024`'s own listed values, on gates in the same function and the same nested block as the disputed `0x3db`/`0x3da` pair — so the pattern that would explain the third does not hold for the two adjacent ones this pass could check directly. Whether case 6's `0x3db` and `ALM-META-024`'s `0x3da` are the same field stays Unknown

**Confidence.** High for every threshold value, its direction, and the corpus-uniform branch each one takes (all read directly from decompiled control flow, cross-checked against the corpus's own two `formatVersion` values by direct arithmetic). High for the case-0 and case-6 field-by-field reconciliation: both sums are read from the same raw disassembly already cited for the threshold facts, not fitted to the target numbers, and each is validated per file against the full declared-`payloadSize` population at 155/155 and 154/154 with 0 exceptions — a wrong field list could not reproduce both games' own constants simultaneously by chance. Unknown, unchanged, for whether case 6's `0x3db` and `ALM-META-024`'s `0x3da` are the same field, and for which specific bytes of case 0's 28-byte growth the writer's own declared `payloadSize` counts versus which land in the 16-byte undeclared overhang

### R2-ASSET-025

No literal reference to `templates.bin` (or two plausible variant spellings) exists anywhere in `allods2.exe`'s image, by direct string search — confirmed, and its bound corrected, below. Six raw-ASCII needles were originally searched against the whole `allods2.exe` image: `"templates.bin"`, `"templates"`, `"template.bin"`, `"data.bin"` (lowercase), `"data/data.bin"`, `"WorldDataData.bin"` — all six reported 0 hits, and a same-search case-sensitivity control showed the instrument distinguishing case correctly (`"Data.bin"`, capital D, found at four addresses; lowercase `"data.bin"` 0 hits). **Correction, finding 8:** that search covered one binary and exact-case ASCII only. Repeating the `templates.bin`-family and `Data.bin`-family needles across all 12 top-level `.exe`/`.dll` files under the ROM2 root and four encodings (ASCII exact, ASCII case-insensitive, UTF-16LE exact, UTF-16LE case-insensitive — `string-search.tsv`) confirms the `allods2.exe` absence is real under every encoding tested (0 hits, `templates.bin`/`templates`/`template.bin`, all four), so the "case-insensitive coverage was not run" caveat is dropped. But `templates.bin` is not absent from the ROM2 root as a whole: `ROM2 Map Editor.exe` carries the literal string `"templates.bin"` at file offset `0x7afdc`, the only hit anywhere in the 12 binaries. This row's own previous bound list named `engine32.dll`, `scenario.dll` and `a2server.exe` as unsearched — all three are now searched and also report 0 hits — but did not name `ROM2 Map Editor.exe`, which is the one binary in the same preserved root that does carry the string; the previous bound was drawn in the wrong place

**Confidence.** High for the absence AS SEARCHED, now over 12 root binaries and 4 encodings rather than 1 binary and exact-case ASCII only: `allods2.exe` carries no `templates.bin`-family string under any encoding tested. High, as a directly observed positive, that `ROM2 Map Editor.exe` — not `allods2.exe` — carries the string, at one fixed offset. Explicitly bounded to the population searched: 12 root binaries, 4 encodings, the 5 needles tested — not every binary, not every possible spelling, and not a claim about what a network-facing consumer might do with a value it never itself stores as this literal string

### R2-ASSET-026

`allods2.exe` contains its own copy of the exact Data.bin loading shape `DAT-LOC-001` already publishes for ROM1's `rom.exe`: a guard flag at a fixed offset (`+0xf0`) on a long-lived object makes the load run at most once; the fast path (`R2.0004`) opens a loose file at a fixed relative path ending `Data.bin`, falling back to the same relative path as an entry inside the `World.res` archive; on failure of both, it parses the same fixed, named set of 11 source tables `DAT-LOC-001` lists for ROM1 (`Spells`, `Armors`, `Materials`, `Shapes`, `Magic`, `Weapons`, `Shields`, `Magic Items`, `Units`, `Humans`, `Buildings`, here `.txt` not ROM1's `.csv`), then rewrites `Data.bin` (`R2.0005`) — both called only from the guarded once-only entry point `R2.0006`. **Correction, finding 9:** the eleven table names were previously stated as "read directly from three decompiled functions"; they were not — they came from a `.rdata` string dump of `[L2.00008,L2.00009)`, adjacent to but not inside any decompiled body. The function that actually reads them, `R2.0007` (called from `R2.0006` between its source-table parsing and binary-write diagnostics — the same position `DAT-LOC-001` gives ROM1's own `R0682`), is now decompiled directly: it calls the same table-reader helper, `R2.0008`, exactly eleven times with the successive table-name arguments in a nested early-exit chain (each table opened only if the previous one succeeded), in the same order the `.rdata` dump already showed — `Spells, Armors, Materials, Shapes, Magic, Weapons, Shields, Magic Items, Units, Humans, Buildings`. The conclusion is unchanged; what changes is the basis — a decompile of the actual reader, not adjacency in a 336-byte data window. `R2.0006`'s only two callers in the whole binary are `R2.0009` and `R2.0010`; the smaller reads a `"BaseDir"` value under an `"Allods Server"` profile section and takes an argument-count-shaped parameter, consistent with a one-time process/session-startup routine, not a per-map call site (a structural signal, not a full call-graph trace to a named entry point). This experiment did not locate ROM2's own counterpart of ROM1's `R0486` (`ALM-CLS-036`'s type-id-to-definition resolver) — the function that would actually consume an `.alm` record's raw type id against this Data.bin-built table is Unknown; the record dispatcher (`R2-ASSET-021`) stores raw ids into its record objects, and resolution against a type table, if it happens the same way ROM1's does, happens in code this experiment did not trace

**Confidence.** High that a Data.bin-shaped loading subsystem exists and runs in `allods2.exe` — now resting on a decompile of the table-name reader itself (`R2.0007`), not only on 11 names' adjacency to two already-decompiled `Data.bin` literals in `.rdata`, plus the matching guard-offset/fallback structure, which together make an unrelated, coincidentally similar subsystem implausible. Unknown, not H5-databin confirmed, for whether THIS is what an `.alm` record's type id actually resolves against — Q5's own consumer link was not traced, so alongside `R2-ASSET-025`'s templates.bin string-absence this makes H5-databin the structurally more plausible of the two named candidates without confirming it

### R2-ASSET-027

ROM2's own corpus does not exceed ROM1's own corpus maximum in `W`, `H`, or `#objects`/`#units` (record-0 metadata, `ALM-META-024`/`ALM-META-025`), and exceeds it in `#players` — but within an already-published structural capacity, not past it. Measured maxima, both games' own preserved-root corpora, full population: `W` ROM2 256 (`CROSS2.ALM`) vs ROM1 256 (EN `Beast.ALM`, RU `FORESTER.ALM`) — equal, not exceeded; `H` identical pattern; `#objects` ROM2 324 (`CROSS2.ALM`) vs ROM1 478 (EN `Beast.ALM`); `#units` ROM2 785 (`CROSS2.ALM`) vs ROM1 1815 — ROM2 under ROM1's own maximum on both. **Correction, finding 7 (smaller consequence):** the `#units` maximum was attributed to "`Horror.alm`, both roots"; RU's own `extents.tsv` row for `Horror.alm` reports the identical declared `#units`/`#objects` values (1815/415) only because RU `Horror.alm` is a truncation of the EN file with two differing prefix bytes and the same declared unit/object counts (`R2-ASSET-017`'s own correction) — RU's own copy carries no type-6 (units) or type-4 (objects) record at all, having been cut after record 3. The maxima themselves are unaffected: the EN full `Horror.alm` supplies the unit maximum, while EN `Beast.ALM` supplies the 478-object maximum. The former claim that Horror dominates both figures is partially retracted for that attribution only (see `retracted.md`). The unit-maximum attribution remains "`Horror.alm`, EN root — RU's own copy of the file declares the identical metadata value with no backing record, being a truncation of the EN file, not an independent measurement." `#players` ROM2 15 (`CROSS2.ALM`) vs ROM1's own separately-OBSERVED corpus maximum of 9 (`scenario.res:90.alm`, both roots) — exceeds. But ROM1's own type-5 per-player record already documents a fixed 16-slot diplomacy-row field (`ALM-GRP-041`: one `u16` for each of sixteen editor player slots; `AI-DIPLO-005`'s 38-map/191-player/3056-cell corpus replay found every shipped roster covered by that row) — a structural capacity ROM1's own shipped corpus never filled past 9. ROM2's own type-5 record size still matches ROM1's exactly, `76*nPlayers` on 83/83 (`R2-ASSET-003`). **Correction, finding 3: this Medium is closed.** ROM2's own type-5 record CONTENT, not only its overall size, is now directly read: case 5 of `R2.0002` reads four leading fields (widths 4, 4, 4, 32), then a loop whose own bound is the literal constant `0x10` (16) at `L2.00010`, ending iteration at that fixed slot count rather than at a file-supplied count, each iteration a 2-byte field at a destination reached through `R2.0011(player+0x38,i)` — ROM2's own counterpart of the `R0130(player+0x34,i)` `ALM-GRP-041` names for ROM1. Total: `4+4+4+32+16×2=76` bytes, the same arithmetic `ALM-GRP-020`/`ALM-GRP-041` publish for ROM1's own type-5 record, matched field for field (same four leading widths, same fixed sixteen-slot loop, same 2-byte slot width), not only by total size. ROM2's observed 15-player maximum is within that now-directly-read 16-slot capacity. Separately, and not scored as a hypothesis (a direct value comparison, per `HYPOTHESES.md`): 0 of 83 ROM2 files pass ROM1's own documented header gate (`magic` + `recordCount>=3` + `formatVersion<=1001`) — every file fails on `formatVersion` alone, since ROM2's own corpus-wide values (1300, 1600, `R2-ASSET-003`) are both above 1001; whether ROM1's loader enforces any hard reject specifically on `#players` above 9 (as opposed to the type-5 record's own 16-slot field capacity) is a separate question this experiment does not resolve, consistent with `HYPOTHESES.md`'s own scoping

**Confidence.** High for the five measured maxima and the 0/83 header-gate pass rate (direct, full-population value comparisons on both sides, both games read fresh by the same instrument). High, corrected from Medium, for the 16-slot-capacity framing of the `#players` finding: ROM2's own type-5 record content is now directly read (a literal `0x10` loop bound and the same field widths ROM1's own claims publish, not only a total-byte-count match), field for field against the two cited ROM1 claims, which check out for object and width

**Amended.** ✔ promoted (partially retracted: maxima attribution only; see retracted.md)

### R2-ASSET-029

ROM1's DAT-GRAM-003 grammar, corrected only in group C's raw byte block (widened from ROM1's published 10 bytes to 14), replayed corpus-wide reaches EXACT 0 residue on both ROM2 `data.bin` files (client 149 971 B, server 152 527 B) — resolving `R2-ASSET-004`'s own stopping point (3 of 31 `Armors` entries). All 7 other groups (A, B, D, E, F, G, H) use ROM1's unchanged entry shapes without modification. Negative control: the same 14-byte-widened grammar, replayed against ROM1's own EN/RU `data.bin` (88 327 B each), does NOT walk cleanly — diverges at offset 88 325/88 327 (needs 4 bytes, 2 available). Under the UNCHANGED (10-byte) grammar, by contrast, ROM1's own file walks to exact 0 residue at 88 327 — `DAT-GRAM-003`'s own already-published result, reconfirmed here by an independent walker — so 88 325/88 327 is a new divergence point this 14-byte grammar alone produces, not a reproduction of any offset `R2-ASSET-004` reports (that claim's own two divergence offsets are 149 971 and 152 527, ROM2's own file sizes, explicitly discounted there as non-discriminating; it says nothing about ROM1's file). Independently swept over every group-C raw width from 0 to 80 bytes: 14 is the ONLY width that reaches exact 0 residue on either ROM2 file (client 149 971 B, server 152 527 B), and 10 is the ONLY width that reaches exact 0 residue on either ROM1 file (EN and RU, 88 327 B each) — ruling out "this widened grammar is a universal read-any-Data.bin fix" by the width's own uniqueness in both directions, not merely a single before/after pair (`groupc-width-sweep.txt`). Instruction-level correlate in `allods2.exe`, independently triple-confirmed and tied together by a directly-read pointer, not by address order: the class whose real constructor (`R2.0012`) runs a 7-iteration `u16` zero-init loop (14 bytes) at offset `0x1c` also stamps vtable `L2.00011` at its own separate instruction, forty bytes from the loop's own store, in the same function; that vtable's own slot+8 entry is `R2.0013`; and `R2.0013`'s two raw-block helper calls each receive the literal length 14, supplied four instructions before the call in each direction. This class matches ROM1's own already-published `kArmor` shape (dwordArray, raw block, dwordArray — `tools/r2asset/databin.go`) and is tied to the letter C specifically by the dispatcher's own program-order call position (this row's own call-graph chain, below) — the same kind of argument this row already makes for group A — corroborated independently, on the wire rather than in Ghidra, by the group's already-published title array content (`Item \| Shape \| Material \| ...`). Group A is confirmed the same way: its own real constructor (`R2.0014`) has no zero-init loop; its Serialize body (`R2.0015`) calls one raw-block helper with a literal length of `0x48` (72), matching `kMatShape`'s raw-only shape exactly, no dwordArray or cstring call

**Confidence.** High for the wire-level fact (group C widened 10→14, all other 7 groups unchanged): a corpus-wide clean walk to exact 0 residue across two files totalling over 300 KB, from a grammar that changes exactly one number in one group, is not plausible by chance if any other group's shape had also silently changed, and an 81-value width sweep (0..80) shows 14 is the unique closing width for both ROM2 files while 10 is the unique closing width for both ROM1 files — ruling out both "some other nearby width would also close" and "any widened count would coincidentally also clean up ROM1's own file." High also for group C's and group A's own Ghidra-level identity: three independently-obtained facts (ctor loop trip count, vtable-stamp linkage, Serialize's own disassembled literal) agreeing on the same number, tied to the same compiled class via a directly-read vtable pointer, rules out "the Ghidra find and the wire-level find are two different, coincidentally-matching classes." High also, now, for the remaining six groups' (B, D, E, F, G, H) own Ghidra-level identity: one directly-read call-graph chain — the dispatcher's own program order names each `SetSize` address's letter; each `SetSize` body calls exactly one ctor-loop; each ctor-loop calls exactly one real constructor and passes it the class's own element byte size; each real constructor stamps exactly one vtable literal; each vtable's own slot+8 entry is that class's Serialize body — pins all six the same way A and C were already pinned, independent of address-list position (`serialize-addrs.tsv`). B and G are individually distinguished, not left as an unordered pair, by which construction chain reaches which vtable: B is `L2.00012` (Serialize calls the shared base body `L2.00013` directly), G is `L2.00014` (Serialize is a one-line forward to the same `L2.00013`) — told apart by the chain, not by shape, since both compile to the identical `kBase` behaviour. A ninth candidate vtable in the same slot-shape range, `L2.00015`, shares `L2.00013`'s own Serialize address but is reached by none of the eight construction chains and is excluded from the family on that ground, ruling out "`L2.00015` is a third, ambiguous B/G candidate." Every pinned body's own shape also matches ROM1's already-published kind table (`kMagicItem` for D, `kNStr2` for E, `kNStr10` for F, `kSpell` for H), corroborating the chain-based pin a second way

### R2-ASSET-030

The full 2556-byte client/server `data.bin` size delta (`R2-ASSET-012`) is located exactly, not merely bounded. Using the corpus-wide-clean grammar (`R2-ASSET-029`), the client and server files agree on every group's declared count and consumed span except two: `Units` — identical declared/stored count (242/241, both files) but a +10-byte span difference (50 594 client vs 50 604 server bytes for the collection); `Humans` — declared count 290 (client) vs 299 (server), 9 additional rows, and a +2546-byte span difference (68 070 vs 70 616). Every other group (A, B, C, D, E's title array, G, H) spans byte-identically between the two files. Sum of the two located deltas: 10 + 2546 = 2556, exactly matching `R2-ASSET-012`'s own independently-measured whole-file size difference, with no unaccounted remainder

**Confidence.** High — an exact arithmetic identity (10 + 2546 = 2556) between two independently-derived numbers (a per-group span accounting on one hand, a whole-archive size difference measured by a completely different method — a directory-entry size diff, `R2-ASSET-012`, on the other) is not plausible if any group besides Units/Humans also carried an unreported difference; this rules out "the delta is spread across several groups and only appears to concentrate in two by coincidence"

### R2-ASSET-031

ROM2's group title arrays, read via the corpus-wide-clean grammar (`R2-ASSET-029`) and diffed column-for-column against ROM1's own published `DAT-SCHEMA-007` list (`databin-slots.csv`), are IDENTICAL on 6 of 8 groups (A Shapes, B Magic, C Armors, D MagicItems, F Humans, G Buildings — exact text, exact column order, exact count). Group H (Spells) differs on exactly 1 of 23 titles by two literal double-quote characters only (ROM1: `Radius, Length/2`; ROM2: `"Radius, Length/2"` — confirmed as bytes genuinely present in ROM2's own title string, not a CSV-escaping artifact of either side's own extraction code, since ROM2's side comes directly from this survey's own non-CSV tool output). Group E (Units) grew from 56 to 63 slot titles (excluding the shared name-title): a shared 44-title prefix and a shared 1-title suffix (`EquipItem`) bracket a middle span where ROM1's 11 titles (`treasure.3 Magic`, `treasureMin.3`, `treasureMax.3`, `Power`, `Spell 1..3`+`Probability 1..3`, `Spell Power`) are replaced by ROM2's 18 (`treasureMask.2`, 2 unused-labelled slots, the same `Power`/`Spell 1..3`/`Probability 1..3`/`Spell Power` 8 titles ROM1 already has, plus 7 wholly new: `serverID`, `knownSpells`, `skillFire`, `skillWater`, `skillAir`, `skillEarth`, `skillAstral`) — net a smaller "treasure tier 3" replacement plus 7 new server/skill columns inserted before the pre-existing tail. The full ROM2 per-group title list is retained in the private evidence as `databin-slots-rom2.csv` (180 title rows across 8 groups; name-titles are excluded from slot numbering exactly as `databin-slots.csv` excludes them, DAT-SCHEMA-004's own convention, so the two files' slot numbers for the same column line up directly), title text only — no dest/width/site (that streamer cross-reference is not run by this survey). Separately, an independent walk of each collection's own per-entry PARAM ARRAY (the numeric `dwordArray` most entry kinds carry, distinct from title text) finds every collection's width unchanged between the two games except Units: Magic 28, Armors/Shields/Weapons 17, MagicItems 2, Humans 26, Buildings 6, Spells 22 identical on both sides; Units widens from 55 to 62 (`param-widths.txt`) — tracking this row's own 56→63 title-count growth net of the trailing `EquipItem` string column, which is not itself a numeric param on either side

**Confidence.** High for the direct comparison itself: rules out "the 2 differing groups are a parsing or transcription artifact of this survey's own tooling, not a genuine ROM2 content difference" — the underlying title arrays are independently confirmed to be read cleanly (`R2-ASSET-029`'s 0-residue corpus walk), and group H's own quote characters are traced specifically to this survey's non-CSV Go tool output, ruling out a CSV-escaping artifact on either side. The param-array width census adds a second, independent signal for group E specifically: the per-entry numeric array itself widens by the same count (net of the one non-numeric trailing column) that the title array grows by, ruling out "the extra titles are label-only, with no corresponding wire-level column" as a reading of the title diff. No confidence is asserted for WHY group E grew or WHY group H's one title gained quote characters — this survey reports the schema difference, not its authorial cause

### R2-ASSET-032

Replaying the corpus-wide-clean grammar (`R2-ASSET-029`) against `templates.bin` reaches the same divergence point `R2-ASSET-005` already reported under the unmodified grammar — group A, `Shapes` collection, entry 112, offset 35 004 of 61 802 — because group A's own shape is unaffected by the group-C-only correction; this independently reconfirms, but does not advance past, `R2-ASSET-005`'s own stopping point. Separately, an exhaustive ASCII string search (case-sensitive, both `templates.bin`/`Templates.bin`/`TEMPLATES.BIN` and the bare word `templates`) over `allods2.exe`'s and `a2server.exe`'s own loaded memory images (Ghidra headless, full default auto-analysis; `FindStrXref.java` scans every loaded memory block by a raw byte search, not a defined-data-only search — both halves of this row's own population description previously named the wrong scope) finds ZERO occurrences of any of the four spellings in EITHER binary. A same-search positive control, run in the identical pass, DOES find `Data.bin` (3 distinct string locations in `allods2.exe`, 4 in `a2server.exe`, each with at least one direct code cross-reference) and `World.res` (1–3 locations each) in both binaries — proving the search methodology locates a filename that IS present, so the `templates.bin` absence is not a search-methodology artifact. Independently widened to the whole preserved ROM2 root (all 12 `.exe`/`.dll` files, exact-case ASCII, case-insensitive ASCII, and UTF-16LE, not only `allods2.exe`/`a2server.exe` and not only exact-case ASCII): `templates.bin`/`Templates.bin`/`TEMPLATES.BIN`/bare `templates` remain at ZERO occurrences in both game binaries under every encoding tried, reproducing the same `Data.bin` (4 each) and `World.res` (1 in `allods2.exe`, 1–2 in `a2server.exe`) positive-control counts. The literal string `templates.bin` DOES occur, once, at file offset `0x7afdc` in `ROM2 Map Editor.exe` — a standalone NUL-terminated string among resource/configuration filenames, and in no other binary under the root — locating where the literal name is referenced textually; this pass does not check whether that editor binary's own code reads the file by that name (no cross-reference query was run there), so it is a location, not a confirmed reader (`templates-root-scan.txt`)

**Confidence.** Medium, properly bounded. The string search rules out "a literal filename string 'templates.bin' (in these four spellings, under exact-case, case-insensitive or UTF-16LE encoding) is referenced anywhere in either binary's own loaded memory image" — but does not rule out a dynamically composed path (a directory prefix concatenated with a stored extension, or a name built from a resource ID) reaching the same file by a route this search cannot see, nor does it identify which of `R2-ASSET-005`'s two candidate head readings is correct. Finding the literal string in `ROM2 Map Editor.exe` and nowhere else under the root narrows which binary COULD be its reader without confirming one: that binary's own code was not itself checked for a cross-reference to the string, so this is a location, not evidence the editor actually opens the file. This bounds the no-reader-found-in-the-game-binaries outcome as the best-supported of this experiment's four preregistered alternatives without confirming it outright, and is scoped exactly to the two game binaries, four spellings, and (after independent widening) three encodings and the twelve-binary root this pass searched — not to any other locale build or code path

### R2-ASSET-033

Classifying `allods2.exe`'s Data.bin-reading population against `a2server.exe`'s own function census (`EXP-2001`'s already-published `allods2.tsv`/`a2server.tsv`, `tools/enginematch`'s strict `mnemHash`+`instrCount`+`byteLen` tier) finds: the top-level Serialize dispatcher (`R2.0016`, 752 instructions) — 0 strict matches; the shared `CArchive::IsStoring()` helper (`R2.0017`) — exactly 1 strict match (`R2.0018`); 7 of the 8 group-level Serialize call sites (`R2.0015`, `L2.00013`, `R2.0013`, `L2.00016`, `L2.00017`, `L2.00018`, `L2.00019`) — 0 strict matches each; the 8th (`R2.0019`, an 11-instruction wrapper) — 81 simultaneous strict matches, which this survey does not count as positive evidence of identity (81 simultaneous exact matches for one specific compiled routine already rules out genuine multi-way identity; this is same-shape hash-collision noise against a short, generic instruction sequence, not a discriminating result); the 8 SetSize call sites — 3 or 4 strict matches each (two template-instantiation-shape buckets), the same collision-noise caveat applying (a generic 183-instruction MFC `CArray::SetSize` template shape, content-free of any Data.bin-specific logic). A same-population self-collision control (the identical strict triple, joined against `allods2.tsv` itself rather than `a2server.tsv`) measures the noise reading directly instead of arguing it: the dispatcher, the shared `IsStoring()` helper, and each of the 7 substantive group-level Serialize call sites match EXACTLY 1 function in `allods2.exe`'s own census — themselves — while the 11-instruction wrapper matches 102 and the two SetSize hash-buckets match 11 and 5; every address this row counts as discriminating is unique inside its own binary, every address it sets aside as noise is not (`q5-self-collision.txt`). A second, method-independent check corroborates the strict tier from the bytes rather than the census: a relocation- and `rel32`-tolerant raw search of `a2server.exe`'s whole file image (Ghidra-derived operand masking feeding a byte-level slide match, sharing no code with the `mnemHash` census) finds 0 matches for the dispatcher and all 7 substantive Serialize bodies, agreeing with the strict tier, and ALSO finds 0 matches for all 8 SetSize addresses individually — stronger than the strict tier's own 3–4-per-bucket count there, because the byte method is sensitive to the literal per-instantiation element-size operand `mnemHash` discards entirely. The wrapper and SetSize shapes remain non-discriminating under the byte method too: the wrapper self-matches 102 times (matching the census exactly) and cross-matches `a2server.exe` 82 times (close to, not identical to, the strict tier's own 81 — the two methods compare by different units, a raw byte offset versus a census-recognised function boundary); and the strict tier's two SetSize mnemHash buckets (A, B, D, E, F, G share `a318e0db7a995c24`, 11 self/4 cross; C, H share `00ed006553916876`, 5 self/3 cross) resolve further under the byte method into two true byte-identical pairs (`L2.00020`~`L2.00021`, i.e. B~G; `L2.00022`~`L2.00023`, i.e. E~F) plus four byte-unique singles (`L2.00024`=A, `R2.0020`=D, `L2.00025`=C, `L2.00026`=H) — sharper than either hash bucket alone, though this pass tests each single only against its own seven named siblings and `a2server.exe`, not against the two buckets' own unnamed members inside `allods2.exe` itself (5 further addresses in the six-member bucket, 3 further in the two-member bucket) (`mask-match.txt`)

**Confidence.** High for the dispatcher and the 7 substantive group-level Serialize call sites (all but the 11-instruction wrapper): two independent methods sharing no code — a function census keyed on mnemonic-sequence hash, instruction count plus byte length, and a raw relocation-tolerant byte search of the target's whole file image — both return 0 matches in `a2server.exe`, while the self-collision census shows every one of these addresses is unique even inside `allods2.exe`'s own function population. The byte-level method rules out the alternative the strict tier alone could not exclude: that a genuinely shared, byte-identical function is merely relocated differently between the two binaries and so hashes differently — a relocation-tolerant search would find that function anyway, and does not. Medium remains for the classification's treatment of the wrapper and the SetSize shapes specifically: both are now measured, not merely argued, as generic/multiply-instantiated within `allods2.exe` itself (self-collision 102, 11, 5; the byte method further resolves the two mnemHash buckets into 2 true byte-identical pairs (B~G, E~F) plus 4 byte-unique singles (A, C, D, H), all eight individually 0 in `a2server.exe`), so they are excluded from the positive count for a demonstrated reason rather than a structural assumption alone — but no non-Data.bin negative-control population was run, and the SetSize byte-pairing is not traced past this experiment's own 18 named addresses, so this pass does not independently establish which other MFC template instantiations, if any, they collide with beyond what is named here

## Town and room art populations

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-049 | The declared ROM2 town/room art subtrees contain 1101 keys per locale; 1099 payloads match, and both differences belong to druid lizard BMPs. | High | ● active (branch candidate) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |
| R2-ASSET-050 | Both preserved Kaarg masks are 640x480 indexed BMPs with eight populated action values; native selector byte 80 has no pixel in either mask. | High | ● active (branch candidate) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |
| R2-ASSET-051 | The preserved Kaarg square has 218 numbered 24-bpp animation BMPs in seven series, plus a background, mask and two building highlights. | High | ● active (branch candidate) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |
| R2-ASSET-052 | Kaarg inn art contains a 320x480 background and five taverner series; shop art contains a 288x288 background, six fire series and two movie BMP keeper series. | High | ● active (amended, branch candidate) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |

### R2-ASSET-049

The complete selected population is ten graphics.res interface subtrees
and four movies.res subtrees. Each locale contains 1101 keys. The graphics
subtrees town, town_druid, town_kaarg, inn, inn_druid, inn_kaarg,
shop_druid, shop_kaarg, shopanim and townbirds contain respectively
42, 204, 222, 154, 82, 80, 19, 54, 47 and 41 entries. The movies
subtrees inn_kaarg, shop_kaarg, shop_druid and shopanim contain 0, 41,
61 and 54. The archive walker enumerates 4049 graphics file nodes and
156 movies file nodes on each root. All 1099 equal payloads are compared
by SHA256. The two differing keys are
interface/town_druid/lizard/a_a10013.bmp and a_a10042.bmp.

Evidence is measured/art/{archives,entries,populations,locales}.tsv.
Presence does not establish native selection. R2-ENGINE-231 supplies that
separate contract.

**Confidence.** High for the complete declared archive population and
per-key byte comparison, using the recorded RES walker and source hashes.
It does not depend on function recovery or sprite pixel interpretation.
The zero inn_kaarg movie count is bounded by both 156-node movie listings.

**Unknown.** Resources outside these subtrees and native uses outside
the selected town and room methods.

### R2-ASSET-050

The mask has 307200 pixels on each locale: 217863 zero and 89337
with one of eight nonzero bytes. Counts for bytes 32, 64, 96, 128,
144, 160, 176 and 192 are 10285, 4751, 3180, 18424, 3792, 4415,
40667 and 3823. The payloads match. Both are uncompressed 8-bpp BMPs.
The measurement reads the original indexed pixels, accounts for row
padding and reverses positive-height stored rows before computing boxes.
Exclusive rectangles are recorded in measured/art/masks.tsv.

The native byte-80 arm exists, but neither preserved Kaarg mask contains
that value. This absence is not asserted for another mask or another path.
R2-ENGINE-233 maps the whole byte domain to native selectors.

**Confidence.** High for every pixel in these two hashed resources.
The instrument enumerates all 307200 cells per resource rather than
inferring regions from visible color or animation bounds.

### R2-ASSET-051

The seven numbered series under interface/town_kaarg are:
dervish/d0000..d0029 (30, 112x148); guard/g0000..g0054 (55,
72x68 or 76x68); girl1/g1000..g1030 (31) and g2001..g2030 (30),
all 44x108; girl2 has the same two index ranges at 68x104; and
maingates/m0000..m0010 (11, 60x84). The 218 numbered entries and
four single pictures total 222. townmain and townmask are 640x480;
hili_shop is 40x76 and hili_tavern is 140x128. Every visible square
BMP is uncompressed 24-bpp. The indexed mask is separate.

Evidence is measured/art/{entries,sequences}.tsv. Native loader bounds
agree with these ranges in measured/square/bounds.tsv. Native frame
cursors index vectors: cursor 0 of a second girl subset selects file
g2001, not a missing g2000. R2-ENGINE-235 supplies advancement.

**Confidence.** High for header geometry, every filename and contiguous
range in the measured population. Both Kaarg art payloads agree.

**Unknown.** Runtime device color conversion and owner-observed pixels.
Stored RGB does not establish a fixed RGB565 display.

### R2-ASSET-052

interface/inn_kaarg/TavernMain.bmp is 320x480. Its taverner series
a1 has 15 frames, 0001..0015, at 248x296; a2 has 25, 0001..0025,
at 184x188; a3 has 3, 0001..0003, at 44x20; a4 has 7,
0000..0006, at 64x64; a5 has 29, 0000..0028, at 84x116.

interface/shop_kaarg/ShopMain.bmp is 288x288. ShopFrame.256 has
one indexed sprite envelope at 316x303; native pixel semantics are not
established by this envelope. Each of the fire/dark and fire/select
lite, burn and cicle BMP families has six frames, 1..6, at 80x88.
movies.res shop_kaarg/a1 has 21 BMPs, 0000..0020, and a2 has 20,
0001..0020, all 72x84. Every measured Kaarg room payload agrees
between locales.

Evidence is measured/art/{entries,sequences}.tsv. These are art
populations; R2-ENGINE-237 separately identifies native selection.

**Confidence.** High for the named resources, header measurements and
full sequence populations. No claim imports ROM1 room drawing behavior.

**Unknown.** Full room composition, animation destinations and scheduling,
and the original pixel decoder of the shop frame.

**Amended.** R2-ENGINE-277, R2-ENGINE-278 and R2-ENGINE-279 answer the inn and shop animation destinations, triggers and the 100 ms step rule; delivered paint cadence stays Unknown.

## First-town art

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|

| R2-ASSET-057 | The selected first-town ROM2 art contains 339 equal payloads per locale; both 640x480 town masks contain 150 index values. | High | ● active (branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ASSET-058 | The generic ROM2 square has 36 numbered animation BMPs and 41 own-palette sprite resources containing 1180 frames per locale. | High | ● active (branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ASSET-059 | Selected ROM2 .16a frames fit native WORD runs and BMP copy paths differ; the source-RGB composition remains Medium. | High / Medium | ● active (branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ASSET-060 | Generic ROM2 inn art has 154 keys; shop art has 47 graphics and 54 movie BMPs plus a ShopFrame.256 whose pixel stream remains unread. | High / Unknown | ● active (amended, branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |

### R2-ASSET-057

The declared population is graphics.res interface/town, townbirds, inn,
shopanim and ShopFrame.256, plus movies.res shopanim. Counts are
42, 41, 154, 47, 1 and 54 per locale: 339 keys. Every selected payload agrees
by SHA256 between the installed EN/RU resources. This is resource equality,
not a runtime comparison.

Each townmask.bmp contains 307200 pixels and 150 original index values.
Zero occupies 266003 pixels. Counts for native recognized bytes
32, 64, 80, 96, 128, 144, 160, 176, 192 are 17, 18, 4, 6, 5764, 5771, 7911, 5592, 14733.
Their total is 39816. The other 141 present values contain 267384 cells,
including zero. Nonzero default values contain 1381 cells. The native
selector mapping is a separate engine claim.

**Confidence.** High for the complete declared subtree population and
all mask pixels. The generator walks every file node, reads the indexed
pixels at the BMP offset, includes row padding and reverses positive-height
rows. It records every populated value and its exclusive box. This count
does not infer regions from visible color and makes no claim about the
intent behind minor index values.

**Unknown.** Resources outside the declared population and runtime use
outside the selected loaders. Evidence is measured/art/{archives,entries,
locales,populations,masks}.tsv.

### R2-ASSET-058

The 42 generic town BMPs comprise townmain/townmask at 640x480,
Tavern_l 28x64, Shop_l 52x76, Trener_l 140x116, Town_add 552x92 and 36
numbered pictures. Sign V00..09 has 10 frames at 40x32; door T00..08 has 9
at 36x48; stars S00..08 has 9 at 64x44; fluger F00..07 has 8 at 64x64.
The visible BMPs are uncompressed 24bpp; the mask is uncompressed 8bpp.

The 41 townbirds .16a resources contain 1180 frames. Birds1..9 each have 57 frames;
Guards has 8, Tavern has 10, Shopie has 30, Fighter has 11 and Mage has 11. Each Horse1..5/A1..3
has 15. Each Baba1..4/A1 has 31 and /A2 has 32. Each Dervish1..4 has 30.
Per-frame dimensions are recorded rather than inferred from file names.
Resource presence does not establish a paint or advance call.

**Confidence.** High for each header, sequence range and frame count
in both hashed populations. Every .16a palette flag and frame envelope
fits exactly through the final frame-count trailer. Evidence is
measured/art/{entries,frames,sequences}.tsv.

**Unknown.** Device conversion and live frame choices. R2-ASSET-059
separates native pixel-command semantics from owner RGB composition.

### R2-ASSET-059

The ROM2 EN bitmap table L2.00574 slot+18 calls L2.00575 then L2.00576
for opaque converted-WORD copying. Slot+38 calls L2.00704 then L2.00705
and skips source WORD zero. Town_add uses the second path. The loader
L2.00573 and converter L2.00577 use live display channel bit counts/shifts;
source RGB does not establish a fixed device pixel format.

The .16a constructor L2.00639 installs table L2.00640. Slot+18 calls
L2.00641 then L2.00642. Commands test 0x4000 first (skip rows), then 0x8000
(skip pixels); their low 14 bits are counts. Otherwise the WORD counts a
literal run followed by WORD palette addresses. All 1385 selected .16a
frames per locale, from 67 resources, fit these commands and their measured
extents. Literal addresses are even and confined to 16*256 WORD entries:
index = (word>>1)&255 and level = (word>>9)&15.

Generic square setup L2.00644 receives (16, 4, 0). Type 4 of the six-entry
palette dispatcher L2.00645 selects L2.00647. Normal mode source channels
use floor(channel*(level+1)/16); destination table L2.00649 uses
floor(quantized-channel*(15-level)/16). The reduced mode uses /18 for the
source and a smaller destination table. The owner composition uses normal
RGB factors before device conversion and explicit static frame selections.

**Confidence.** High for these selected native paths and all recorded
command-fit checks. Direct receiver tables, the finite type 4 dispatcher
and complete selected bodies are retained in measured/art/native-*.
The .16a decoder follows these ROM2 instructions. Medium for the resulting
RGB composition because device masks, reduced-mode state and live frame
choices are unobserved.

**Unknown.** Exact runtime framebuffer pixels, display masks and selected
palette mode. A lawful observed framebuffer paired with the native device
state would settle them; no original game was executed.

### R2-ASSET-060

Generic inn art contains 128 BMPs and 26 .16a resources with 205 frames.
CenterArea is 320x480. LeftStats and ButtonsArea are 160x238; LeftPicture
160x242; ManBack/ManBackTalk 48x64; LU/RUOver 16x238; LDOver/Tav_09 16x242;
the six main button pictures 140x46. Candle T0000..0009 is 80x120;
Cauldron T0000..0020 is 60x252; Tender drink DR0001..0040 and breath
BR0001..0024 are 180x212.

Generic shop graphics contain ShopMain 288x288, MyItem/ShopItem 80x80
and four subtrees 01..04 with 11 numbered BMPs each (1..11), at 112x88,
132x88, 80x112, 80x112. The 54 shopanim movie BMPs are 76x176: No1..12,
Pose2-3/1..29 and Yes1..13. Every named BMP is uncompressed 24bpp.

ShopFrame.256 has one palette-flagged frame header 316x303 and data-size
field 11015. Header-described bytes total 12051 and leave 1743 bytes before
the frame-count trailer. Their role and the palette 8 pixel stream remain
Unknown; this measurement does not claim an exact rendered shop frame.

**Confidence.** High for selected headers, sequence populations and hashes,
using measured/art/{entries,frames,sequences}.tsv. Unknown for ShopFrame.256
pixel semantics and its unclassified tail. The native palette 8 reader
would settle that bounded gap. Room composition is a separate engine claim.

**Amended.** R2-ASSET-070 and R2-ENGINE-270 decode the ShopFrame.256 pixel stream with the native byte reader and palette mode 1; the 1743-byte tail stays unread by that reader and its purpose remains Unknown.

## Druid archive population

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-063 | The four declared druid art subtrees contain 366 keys per locale; 364 payloads match, while two lizard source BMPs differ. | High | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |
| R2-ASSET-064 | Both preserved druid town masks have 307200 identical indexed cells, with five populated nonzero action bytes: 80, 96, 128, 144 and 160. | High | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |
| R2-ASSET-065 | The preserved druid square contains 113 person BMPs and 183/42-frame bug/lizard sprite envelopes; its inn and shop art contain the measured keeper and waterdrop series. | High | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |
| R2-ASSET-066 | Selected ROM2 sprite code establishes the druid 16a envelopes and word-run geometry; owner compositions approximate its full-table memory-mode blend in source RGB. | High / Medium | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |

### R2-ASSET-063

The complete selected archive population is graphics.res
interface/town_druid, interface/inn_druid and interface/shop_druid,
plus movies.res shop_druid. Each locale contains respectively 204, 82,
19 and 61 entries. Town has 202 BMPs and two 16a members. Inn has
82 BMPs. Shop has 18 BMPs and one 256 member. The movie subtree has
61 BMPs. Both archive registry traversals account for every node.

All 366 keys agree. SHA256 payload comparison identifies 364 equal
members. The two differences are graphics.res keys
interface/town_druid/lizard/a_a10013.bmp and a_a10042.bmp.
Presence does not establish native selection. These source BMPs are not
selected by the measured druid square loader, which loads sprites.16a.

Evidence: evidence/measured/assets/{archives,entries,populations,locales}.tsv.

**Confidence.** High within the four complete declared subtrees of the
two hashed resource installations. The RES walker reads registry nodes
and payload extents without depending on native function recovery or
sprite pixel decoding. This narrows and agrees with R2-ASSET-049.

**Unknown.** Art outside these prefixes and resource uses outside the
selected native square and room paths.

### R2-ASSET-064

Both interface/town_druid/townmask.bmp members are uncompressed 8-bpp
BMPs at 640x480. Original indexed bytes give 262188 cells of byte0,
5421 of byte80, 2144 of byte96, 3397 of byte128, 24098 of byte144
and 9952 of byte160. The sum is 307200. EN/RU payloads and every
indexed cell agree. There is no differing pixel bounding rectangle.

Exclusive nonzero rectangles are: byte80 (338,253)-(390,368);
byte96 (166,202)-(201,270); byte128 (120,189)-(169,274);
byte144 (375,223)-(570,395); byte160 (0,293)-(174,374).
No other byte occurs in either selected resource. This includes byte176,
although a native menu selector exists outside this populated mask.

Evidence: evidence/measured/assets/{masks,mask-locales}.tsv. Native action semantics
come from the separately measured druid sampler and click methods.

**Confidence.** High for all cells in these two hashed BMP members.
The probe accounts for BMP row padding and positive-height row reversal,
then counts raw indices before any palette or RGB conversion.

**Unknown.** Runtime pointer coordinates and any other mask resource.

### R2-ASSET-065

The town_druid background and mask are 640x480. hili_shop is
152x164; hili_tavern is 152x96. Man BMP series a1/a2/a3 contain
19/19/20 files, numbered from 1; a1/a2 are 48x72 and a3 is 60x88.
Woman a1/a2/a3 contain 15/20/20 files, numbered from 1; a1/a2 are
84x128 and a3 is 80x148. All six series have no internal missing
number. Their combined count is 113.

Bug sprites.16a contains 183 frames, all 640x268. Lizard sprites.16a
contains 42 frames, all 100x120. The lizard subtree also contains two
42-file BMP source series a1 and a_a1, all 100x120, and palette.bmp
at 102x122. These extra source BMPs are a resource census, not additional
layers inferred into the native square composition.

Inn TavernMain is 320x480. Inn taverner a1 contains 40 files at 128x196;
a2 contains 30 at 128x196; a30001 is 32x24. Waterdrop d0001..d0010
contains 10 files at 116x96.

ShopMain is 288x288; ShopInv is 164x303; ShopMenu is 176x238; ShopTable is 472x87.
ShopFrame.bmp and the one-frame ShopFrame.256 envelope are 316x303.
ShopArrow1..4 are 72x32. ShopButton1..4 use 120x52 or 140x46.
Elven and hili_elven are 96x104; hili_armor is 88x112;
hili_magic is 96x60; hili_potion is 72x112. The movies.res keeper a1/a2/a3
groups contain 30/30/1 BMPs, all 72x164, numbered from 1.

Evidence: evidence/measured/assets/{entries,frames,sequences}.tsv. Archive frame
counts and dimensions do not independently establish native schedules.

**Confidence.** High for the complete selected filename ranges, BMP
headers and sprite envelope geometry. EN/RU members agree except the
two source BMPs named by R2-ASSET-063. Native loader bounds and
cursor semantics are separate engine claims.

**Unknown.** Runtime device colors, the native indexed ShopFrame.256
pixel decoder, and resource consumers outside the measured classes.

### R2-ASSET-066

EN sprite constructorL2.00639 installs vtableL2.00640. Slot18 targets
L2.00641. The no-flag painter branch forwards the selected frame header
width/height, data at header+12 and palette table to L2.00642.
LoaderL2.00643 reads the final DWORD count, clears its high own-palette
bit, reads 1024 leading palette bytes when set, and walks12-byte
width/height/payload-length headers without remaining envelope bytes.

The selectedL2.00642 opcode order tests4000 for row skip, then 8000
for pixel skip, otherwise consumes that word's literal count. Skip
counts retain the low14bits. A completed width advances one row.
All450 bug/lizard frame streams across the two installations fit the
selected word-run geometry and own-palette offset bounds.

Druid configurationL2.00644(16,4,0) reaches palette builderL2.00645.
The six-entry tableL2.00646 positively selects L2.00647 for mode4.
For memory-table flag L2.00648=0, mode4 builds palette rows1..16 with
each source component multiplied by row and truncated after division16,
then packs the runtime color masks. The normal branch of destination
LUT builderL2.00649 and blitterL2.00642 gives foreground weight
(level+1)/16 and destination weight(15-level)/16 for literal bits9..12.

The owner render keeps source RGB and approximates those factors using
RGBA before native packing. It uses the separately measured square
destinations and frame selectors. It omits child overlays and device
presentation. No screenshot or original game execution is a witness.

The table-mode initializer calls the positively identified KERNEL32
GlobalMemoryStatus import at cellL2.00650. It compares result structure+8
with 24000000 and sets flag L2.00648 to 1 when that unsigned word is below
the threshold, otherwise 0. This selector is a memory-status-dependent
table mode, not a display-format selector. The result word's interpretation
and actual runtime value are not required by the conditional arithmetic.

Evidence: evidence/measured/assets/{entries,frames}.tsv and
evidence/measured/sprite/{bodies,controls,inputs}.json plus the nine native listings.

**Confidence.** High for selected EN loader, word-run, table and
conditional full-table memory-mode arithmetic, and for complete frame fits in
the hashed EN/RU resource population. Medium for source-RGB owner
composition: packing, truncation and the live memory-table flag are not
reproduced. EN/RU resource equality is one shared-content dependency,
not a second native execution witness.

**Unknown.** Live memory-table flag, original device packing/presentation,
compact-table pixel output, child drawing and the indexed shop
frame pixel decoder. Authorized device observation or a complete
measured presentation configuration would settle pixel identity.

## First-town shop keeper and frame

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-069 | The ROM2 generic shop keeper uses 29 Pose2-3, 13 Yes and 12 No 76x176 movie BMPs at center (113,112); the native loaders skip Yes 1, Yes 13 and No 1. | High | ● active (branch candidate) | [EXP-2032](../experiments/EXP-2032-rom2-first-town-shop/) |
| R2-ASSET-070 | ROM2 ShopFrame.256 is a 1024-byte palette and one 316x303 byte-RLE frame; 1743 residual bytes, equal to its skip-command count, precede the trailer unread by the shop reader. | High / Unknown | ● active (branch candidate) | [EXP-2032](../experiments/EXP-2032-rom2-first-town-shop/) |

### R2-ASSET-069

`movies.res` `shopanim/Pose2-3` holds files 1..29, `Yes` 1..13 and `No`
1..12 in both locales; each is an uncompressed 24-bit BMP, 76x176 with a
positive height. `graphics.res` `interface/shopanim/ShopMain.bmp` is
288x288, 24-bit. The center draws ShopMain at (5,8) and every keeper
image at (113,112), page (169,8) and (277,112) inside the center at
(164,0).

Loader slot +0x78 loads resting file Pose2-3 1. Slot +0x80 fills the Yes
array with Pose2-3 1 at index 0 and Yes 2..12 at 1..11; slot +0x84 fills
the No array the same way from No 2..12. Slot +0x88 loads Pose2-3 file
counter+1. No selected loader names Yes 1, Yes 13 or No 1. EN and RU
populations, dimensions and loader formats are equal.

**Confidence.** High for these two archive populations and the selected
loader formats.

### R2-ASSET-070

EN and RU `graphics.res` `interface/shopframe.256` are the same 13798 bytes.
Bytes 0..1023 are a 256-entry four-byte palette. At 1024 one frame header
gives width 316, height 303 and dataSize 11015; commands occupy
1036..12050. R2-ENGINE-270 reads them to exactly 12051 after 303 rows: 651
literal commands carrying 8621 palette-index pixels and 1743 horizontal
skip commands covering 87127 cells, with no row-skip command. The final
DWORD at 13794 is `0x80000001`.

The 1743 bytes 12051..13793 start with another `0x80000001`. The loader
creates one frame pointer and the selected reader stops before them, so
they supply no frame and no pixel to the shop paint. Among the top-level
.res archives of each install (ten EN, eleven RU), only `graphics.res`
holds .256 keys: 1929 per locale, eight of them empty
(`cursors/{cast,defend,move,patrol}.256` and
`equipment/{f,m}fighter/primary/0{0,1}05002.256`). A walk of the 1921
non-empty entries by the envelope of R2-ASSET-008 (trailer bit 31 selects
the palette, low bits count frames) decodes every frame to its dataSize.
1916 entries end exactly at the trailer. Five single-frame entries carry residual bytes:
`interface/{myitem,shopitem,shopframe}.256` and
`projectiles/goblin/{arrow,arrowb}.256`. In each the residual length
equals the frame's skip-command count (20, 20, 1743, 8, 11) and begins
with a copy of the trailer.

**Confidence.** High for the layout, the decode stop and the census
counts. Unknown for the residual purpose; the equal counts are an
observation, not a decoded structure.

**Unknown.** What produced or reads the residual bytes; a consumer outside
the selected paint or the producing tool's contract would settle it.

## Kaarg room and square sounds

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-073 | ROM2 sfx.res holds 35 town_kaarg WAV keys per locale; Kaarg loaders find all they name except Kin1 and Kin2, and every found file is mono 16-bit 22050 Hz PCM, equal between locales. | High / Unknown | ● active (branch candidate) | [EXP-2034](../experiments/EXP-2034-kaarg-room-schedules/) |

### R2-ASSET-073

EN `sfx.res` has 446 entries and RU 556; in each, 35 keys contain
`kaarg`: 19 directly under `town_kaarg/`, 10 under `town_kaarg/inn/` and 6 under
`town_kaarg/shop/`. The selected Kaarg square, inn and shop sound loaders
push 35 distinct `Town_kaarg` keys (R2-ENGINE-277, R2-ENGINE-279,
R2-ENGINE-280). 33 are found. `SFX\Town_kaarg\Inn\Kin1.wav` (inn slot
+0xa4) and `SFX\Town_kaarg\Shop\Kin2.wav` (shop slot +0xac) are absent
from both archives. The two archive keys the loaders do not name are
`Kdoor1` and `Kdoor2`, which the gate arms request (R2-ENGINE-236).

Every found key is RIFF WAVE format tag 1, one channel, 16 bits,
22050 Hz. Durations are data-chunk bytes over the byte rate, rounded to milliseconds:

| Keys | Milliseconds |
|---|---|
| `Kvox1`, `Kvox2`, `Kvox3`, `Kvox4` | 3665, 1556, 1469, 1570 |
| `Kbird1`..`Kbird4` | 625, 657, 1016, 1599 |
| `Kman1`, `Kenter1`, `Kenter2` | 4365, 811, 549 |
| `Ksteps1`, `Ksteps11`, `Ksteps2`, `Ksteps21`, `Ksteps3`, `Ksteps31` | 50, 335, 52, 333, 610, 908 |
| inn `Kdish1`..`Kdish4` | 2159, 701, 1712, 1962 |
| inn `Kman2`, `Kman3`, `Kvox5`..`Kvox8` | 849, 615, 3665, 1469, 1570, 1016 |
| shop `Kman4`, `Kotdel`, `Ktools1`..`Ktools4` | 780, 893, 1518, 376, 1402, 2229 |

`Kvox1`, inn `Kvox5` and `SFX\Town\Crowd.wav` have the same bytes. Among
the found keys pushed by the selected bodies, no other pair is byte-equal.
EN and RU payloads are equal for every pushed key.

**Confidence.** High for the archive populations, absences and WAV fields.
Unknown for the effect of loading an absent key.

**Unknown.** What the slot loader stores when the key is absent; a read of
the sound object constructor on a missing entry would settle it.
