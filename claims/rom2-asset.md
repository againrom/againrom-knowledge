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
| R2-ASSET-019 | The condition that selects ROM2's 28-byte extended type-4 (placed object) record layout over the 20-byte base layout, read directly from `allods2.exe`'s decompiled dispatcher `R2.0002` (case 4), is `kind==0x21 ||…` | High | ✔ promoted (amended) | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-020 | `allods2.exe`'s ALM record dispatcher `R2.0002` enforces a widened, code-verified acceptance ceiling, not merely a corpus that happens to ship higher values. | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-021 | `R2.0002` (`allods2.exe`, entry `R2.0002`) is ROM2's own record-body dispatcher, the counterpart of ROM1's `R0478` (`ALM-META-024`). | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-022 | ROM2-only types 10, 11 and 12's element counts are read from the type-0 (metadata) record's own version-gated fields, never from the type's own record payload. | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-023 | ROM2 extended the dispatcher's existing generic per-element read/allocate mechanism to serve its 3 new types; it did not add a second dispatch mechanism. | High | ✔ promoted | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
| R2-ASSET-024 | `R2.0002`'s case 0, case 6, case 8 and case 9 each gate one or more field reads on `formatVersion` thresholds, read directly from the decompiled body in source order: case 0 has four (`0x47d`/1149, `0x4cd`/1229… | High | ✔ promoted (amended) | [EXP-2003](../experiments/EXP-2003-rom2-alm-records/) |
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

**Amended.** The sentence "This corpus's own type-4 `kind` vocabulary is exactly two values, `0x21` and `0x1000021`" holds for the files that census described and no wider. R2-ASSET-123 reads the same field (low word of dword 0x08) in the 46 campaign maps per root and finds kinds 0x46..0x4b on 13 records per root, in maps 10, 20, 21, 30, 40, 50 and 76. Both walks end at their payload end under the 20/28-byte rule, which is not in question. Whether those maps are in the 83-file population of this claim is not established here; the low-byte-mask remark stands for that population only. See retracted.md.

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

**Amended.** The case 6 thresholds `0x47e`/1150, `0x3db`/987 and `0x456`/1110 are tested as `cmp`, `jb`: the two reads are present at the value and above, and the four-byte form is taken at `0x456` and above, not only above it (`ru-dispatcher.txt` `L2.02479`, `L2.02480`, `L2.02481`, read by R2-ASSET-119). The text "if `formatVersion>t`" for these three values is read as "at or above t". The case 0 values are not re-read here. No installed map (1300, 1600) is affected.


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
| R2-ASSET-073 | ROM2 sfx.res holds 35 town_kaarg WAV keys per locale; Kaarg loaders find all they name except Kin1 and Kin2, and every found file is mono 16-bit 22050 Hz PCM, equal between locales. | High / Unknown | ● active (amended, branch candidate) | [EXP-2034](../experiments/EXP-2034-kaarg-room-schedules/) |

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

**Amended.** R2-ENGINE-687 narrows the absent-key Unknown: the slot holds a
sound object with the key and no data, and each start's lazy load ends
without data when stream open `L2.03656` returns 0. That archive lookup
`L2.03657` returns null for an absent key was not read.

## Character generator resources and templates

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-075 | ROM2 graphics.res holds 114 interface/chrgen keys, equal EN and RU; the detail stats art main.res graphics/chrgen/leftup.bmp is localized; generator buttons reuse 7 Inn keys and 17 sfx chrgen keys are equal. | High | ● active (branch candidate) | [EXP-2035](../experiments/EXP-2035-rom2-chargen/) |
| R2-ASSET-076 | The ROM2 Start_MF, Start_FF, Start_MM and Start_FM Humans rows carry attributes 40/36/25/17, 37/39/21/25, 28/20/41/32, 19/23/30/42, school skill 1 at 20 and fighter armour or mage robe items. | High | ● active (branch candidate) | [EXP-2035](../experiments/EXP-2035-rom2-chargen/) |

### R2-ASSET-075

`graphics.res` `interface/chrgen/` holds 114 keys in both locales with
equal bytes: pre-create `mainarea` 640x480, `mask` 640x480 8-bit,
`cancell` 64x76, `okl` 80x76, `tablol` 312x84, `blind\sprites.16a`,
heroes `h{1..4}on` 64x244 and `h{1..4}sel`, `h1sel2`, `h2sel1`,
`h3sel4`, `h4sel3` 268x340, levels `level{0..2}{on,l,lon}` 48x72, 15
frames each of `torch1` and `torch2` 48x72; stats buttons `{p,m}{lon,loff,
nlon,nloff,disable}` 20x20; `fullstatsl` 160x242, `fullstatsr` 16x242,
`rollstatsl` 160x238, `rollstatsr` 16x238, `buttonsarea` 160x238,
`centerarea` 320x480; per class `column` 320x480, `mask` 320x480 8-bit
and five skills of `on`, `shine_off`, `shine_on`; `cube\sprites.16a` and
`loader\*`. `main.res` `graphics/chrgen/leftup.bmp` 160x238 differs
between EN and RU. The detail buttons use `interface/inn/button{1,2,3}
{on,off}` and `buttonsarea`, 7 keys. `sfx.res` holds 17 `chrgen` keys
(`+_-`, `char`, `fire`, `level1..3`, `ok`, ten `skill\*`), equal in both
locales, one channel, 16 bits, 22050 Hz.

Mask censuses (raw indexes, bounding boxes half-open): pre-create values
20, 40, 60 at the three level rectangles, 80, 100, 120, 140 inside the
hero columns, 160 (16,400)-(80,476), 180 (548,400)-(628,476), 200
(173,404)-(480,478). The fighter mask holds 255, 191, 152, 127, 102 in
five bands at x 83..225; the mage mask holds 255, 191, 152, 127, 102 and
24 stray pixels of 103, 138 and 153.

**Confidence.** High for the populations, dimensions, equality and
censuses.

### R2-ASSET-076

Humans rows 26..29 of the client data.bin (EN and RU `world.res`, equal
bytes) and of RU `world_srv.res` hold the four rows with identical values. Parameter n carries
title n+1; the first title names the entry and the last (`EquipItem`) the
ten strings.

| Row | Body/Reaction/Mind/Spirit | HealthMax | ManaMax | Speed/Rotation | Scan/Defence | typeID/face/gender | serverID | knownSpells |
|---|---|---|---|---|---|---|---|---|
| Start_MF | 40/36/25/17 | 50 | −1 | 19/19 | 6/7 | 3/32/0 | 10210 | −1 |
| Start_FF | 37/39/21/25 | 20 | −1 | 20/20 | 6/8 | 14/11/1 | 10211 | −1 |
| Start_MM | 28/20/41/32 | 30 | 70 | 16/16 | 6/4 | 24/15/0 | 10212 | 150994944 |
| Start_FM | 19/23/30/42 | 30 | 70 | 17/17 | 6/5 | 24/6/1 | 10213 | 150994944 |

All four store Skill.General 0, school skill 1 (Blade or Fire) 20 and −1
for school skills 2..5. Start_FF stores AttackChargeTime 20 and
AttackRelaxTime 5; the mage rows DyingTime 12; other timing fields −1.
Fighter items: Iron Long Sword, Iron Plate Cuirass, Iron Plate Boots, Iron
Plate Bracers, Iron Scale Gauntlets, Hard Leather Mail. Mage items: Wood
Staff with castSpell Fire_Arrow:1, Uncommon Robe, Uncommon Cloak. Loader
`R2.0088` compares the requested name with each Humans row name
(R2-ENGINE-126 names the rows).

**Confidence.** High for the stored rows in the three payloads.

## Tip texts and popup art

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-080 | ROM2 ships tip text only as twelve sections #tips1..#tips12 of main.res text/town.txt, once each in EN and RU; no other archive entry, map, library or executable in either root holds a #tips key or a tips= value. | High | ● active (branch candidate) | [EXP-2037](../experiments/EXP-2037-rom2-tips/) |
| R2-ASSET-081 | ROM2 tip popups draw graphics.res interface/lm.256 frames 9..17 (48x32 and 32x32), interface/radiob.256 frames 4 and 5 (16x16) and the font2 8x10 font of 224 glyphs; the bytes are equal in EN and RU. | High | ● active (branch candidate, amended) | [EXP-2037](../experiments/EXP-2037-rom2-tips/), [EXP-2091](../experiments/EXP-2091-rom2-shop/) |

### R2-ASSET-080

The census reads every entry of every `.res` archive and every `.alm`,
`.dll` and `.exe` file under both roots as raw bytes and searches
`#tips` with optional digits and `tips=`, ignoring case. Mission and quest
texts are uncompressed `main.res` entries (`text/mission<N>.txt`,
`text/quest.txt`) and are inside that population.

| File | Entry | Keys |
|---|---|---|
| EN and RU `main.res` | `text/town.txt` | `#tips1`..`#tips12`, each once |
| EN and RU `allods2.exe`, RU `a2server.exe` | | the strings `#tips` and `tips=` |

Section lengths in bytes, from the key end plus two to the next `#`, after
CharToOemA with an assumed code page (EN 1252 to 437, RU 1251 to 866):

| Key | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 |
|---|---|---|---|---|---|---|---|---|---|---|---|---|
| EN | 313 | 398 | 173 | 239 | 329 | 374 | 281 | 89 | 119 | 62 | 207 | 259 |
| RU | 284 | 335 | 150 | 183 | 389 | 396 | 321 | 94 | 92 | 70 | 205 | 281 |

Every section is one paragraph except RU `#tips12`, which has two.
R2-ENGINE-307 gives the code site of each key.

**Confidence.** High for the population searched. A text loaded from a
file outside the two roots, or a compressed payload, is outside it.

### R2-ASSET-081

`graphics.res` `interface/lm.256` holds 18 frames. The tip overlay
(R2-ENGINE-313) uses frames 9 (48x32 fill), 10, 12, 15 and 17 (32x32
corners), 11 and 16 (48x32 top and bottom edges), and 13 and 14 (32x32
left and right edges). `interface/radiob.256` frames 4 and 5 are 16x16.
`graphics\font2\font2.16` holds 224 frames of 8x10 pixels in 4-bit runs
(decoder EN `L2.01189`: control 0x40 skips rows, 0x80 skips pixels, any
other byte N copies N bytes of two pixels, low nibble first, ending at a
zero high nibble) followed by 32 bytes. `font2.dat` holds 224 dword
advances; entry 0x20 is 6 (labelled space here; corrected under Amended). Frame sizes, opaque counts and SHA-256
hashes are equal in EN and RU.

**Confidence.** High for the frame headers, the decoder rule and the
equality.

**Amended.** R2-ASSET-182 corrects the space label: the draw indexes
`font2.dat` by the mapped code minus 0x20, so entry 0x20 is '@'
(advance 6) and the space is entry 0 (advance 0), High. That the drawn
space width also adds half an image width through an unread vtable
+0x24 is Medium.

## Music archive and map music areas

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-083 | ROM2 EN and RU music.res are identical: 21 PCM keys b00..b16, chrgen, credit, map and menu, all stereo 22050 Hz 16-bit; Music.ini lists 17 titles; no game image holds music.ini or melody in ASCII. | High | ✔ promoted (branch candidate) | [EXP-2040](../experiments/EXP-2040-rom2-music/) |
| R2-ASSET-084 | ROM2 ALM type 12 holds music areas of seven int32 (x, y, radius in tiles, four themes, -1 for none); all 46 campaign maps per root carry a default record and 2..33 areas; themes are 0..16. | High | ✔ promoted (branch candidate) | [EXP-2040](../experiments/EXP-2040-rom2-music/) |

### R2-ASSET-083

EN `music.res` and RU `MUSIC.RES` are 140067844 bytes with equal SHA-256.
Each holds 21 file nodes at the root, in this order:

| Key | Bytes | Seconds | Key | Bytes | Seconds |
|---|---:|---:|---|---:|---:|
| b00.wav | 7688332 | 87.169 | b11.wav | 4619292 | 52.372 |
| b01.wav | 6920696 | 78.465 | b12.wav | 9261044 | 105.000 |
| b02.wav | 8564176 | 97.099 | b13.wav | 12170708 | 137.989 |
| b03.wav | 6669132 | 75.613 | b14.wav | 4467500 | 50.651 |
| b04.wav | 8899164 | 100.897 | b15.wav | 7404536 | 83.951 |
| b05.wav | 6655172 | 75.455 | b16.wav | 6050536 | 68.600 |
| b06.wav | 8874020 | 100.612 | chrgen.wav | 4565508 | 51.763 |
| b07.wav | 7371544 | 83.577 | credit.wav | 7835180 | 88.834 |
| b08.wav | 5924884 | 67.175 | map.wav | 3437504 | 38.973 |
| b09.wav | 4810760 | 54.543 | menu.wav | 3528496 | 40.005 |
| b10.wav | 4348964 | 49.307 | | | |

Every entry is a RIFF WAVE whose RIFF size equals the entry size minus 8,
with a 16-byte fmt chunk (format 1, 2 channels, 22050 Hz, 88200 bytes per
second, block align 4, 16 bits) and one data chunk of the entry size
minus 44 bytes.

Root `Music.ini` (399 bytes, equal in EN and RU) has one section
`[Global]` with `Count=17` and `Melody1`..`Melody17`. Each `MelodyN`
value has the length and SHA-256 of the EN `main.res` `text/tunes.txt`
name of `b(N-1).wav`. `text/tunes.txt` holds 21 `key=name` lines, one per
archive key, except that it names `credits.wav` for the archive key
`credit.wav`; the RU file has the same keys with other names.

A raw ASCII census of every file in both roots and every `.res` entry,
ignoring case, finds `music.ini` only in `ROM2 Map Editor.exe` and the
GOG installer files `goggame-galaxyFileList.ini` and `goglog.ini`, and
the token `melody` only in `Music.ini` and `ROM2 Map Editor.exe`. No game
image (`allods2.exe`, `Scenario.dll`, RU `a2server.exe`) contains either
string. A profile read of the `MelodyN` keys would need the token
`melody`, in full or composed from parts; its absence is the evidence
against a reader that builds the file name at run time. The map editor's
Help describes map music areas with up to four themes and a radius.

**Confidence.** High for the archive, header and file measurements and
for the ASCII census. UTF-16 strings and a key name composed entirely
from other parts are outside the census.

### R2-ASSET-084

Type 12 (R2-ASSET-023 gives its extent) is one 28-byte head record then
`meta[+0x40]` area records of the same shape:

| Offset | Type | Field |
|---:|---|---|
| 0x00 | i32 | x, tiles |
| 0x04 | i32 | y, tiles |
| 0x08 | i32 | radius, tiles |
| 0x0c..0x18 | 4 × i32 | themes 0..3: index into the mission list `B00`..`B16`, -1 for none |

The field meanings come from the loader and area select of
R2-ENGINE-336. The population is 142 maps: 46 `scenario.res` maps and
13 (EN) or 37 (RU) root `.alm` files, versions 1300 and 1600, parsed
with record 0 consuming 660 bytes.

| Population | Head | Areas |
|---|---|---|
| 46 campaign maps per root | (0,0,0) with 1..4 themes at least 0 (EN: 1, 7, 23 and 15 maps) | 470 per root, 2..33 per map; 10 have all themes -1 |
| EN `CROSS2.ALM`, `SHALLOW.ALM` | (0,0,0), themes -1 | 9, 7 |
| RU `CROSS2.ALM`, `Road.alm`, `shallow.alm`, `wsdesert.alm`, `wsenigma.alm`, `wsruined.alm` | (0,0,0), themes -1 | 9, 8, 7, 24, 6, 4 |
| other root maps (11 EN, 31 RU) | (0,0,0), themes -1 | 0 |

The campaign records are equal in EN and RU. Area radii are 4..25 (EN)
and 3..25 (RU). Every theme value 0..16 occurs; no value lies outside
-1..16. Root maps admit no head record, so until the hero first enters
an area their music follows the list order; campaign maps admit the head,
so their mission music is always a theme of an area or of the head
(R2-ENGINE-336).

**Confidence.** High for the record shape over the 142 maps and for the
field meanings, which the loader and area select read.

## Automatic-casting text

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-085 | EN/RU main.res dialogs.txt lines 171..173 name Minimum/Average/Maximum and Минимальный/Средний/Максимальный; RU options request those indices for control 0x8d. The main.txt autohealing labels are a separate population. | High / Medium | ✔ promoted (branch candidate) | [EXP-2042](../experiments/EXP-2042-rom2-autocast/) |

### R2-ASSET-085

The selected `main.res:text/dialogs.txt` nodes have label at physical line
170 (`AutoCast mode:` / `Автокастинг:`), followed by lines 171..173:
Minimum / Average / Maximum; Минимальный / Средний / Максимальный.
Lines 167..169 are Own / Allies / Neutral and Своих персонажей /
Союзников / Нейтральных. Indices are zero-based physical source lines.
Input and selected-node hashes are in the experiment evidence.

RU options request native table indices `0xab`, `0xac`, `0xad` for
control `0x8d` and apply indices 0/1/2 to masks 8/24/56. The table-loader
binding of object `L2.01500` to dialogs.txt was not read. A matching label
sequence and the explicit indices support the pairing, at Medium.

Both roots' `text/main.txt` lines 218..220 instead say Autohealing off /
Occasional autohealing on / Frequent autohealing on and Режим автолечения
выкл. / Режим автолечения редко / Режим автолечения часто. No control
binding for those main.txt lines was established. They are not the label
evidence for the three masks.

**Acquisition support.** The [clause table](../experiments/EXP-2042-rom2-autocast/evidence/support-other.tsv)
records first main.txt labels in navigation run 8 (run-end minimum 40),
options indices/masks by run 10 (55), and the selected dialogs.txt nodes in
generation (121). EN stored-source High and RU-view/native-binding Medium
remain; rendered glyphs remain Unknown. Late reproduction supplies no High
rationale.

**Confidence.** High for EN ASCII stored labels and source indices.
Medium for RU labels viewed with CP1251 and for the options-to-source
binding. The short-label population is two selected nodes per root,
not every installed text source.

**Unknown.** Native RU conversion/glyph identity, table-loader binding and
EN options construction. Read the loader and glyph conversion or observe
the original options dialog; a CP1251 viewing transform alone does not
prove native display encoding.

## Automatic-casting options and feedback pools

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-088 | RU options label access uses the dialogs pool, while selected key feedback uses the patch/patch.txt pool at indices 115..117; the actual feedback text was not read. | High / Medium / Unknown | ✔ promoted (branch candidate) | [EXP-2044](../experiments/EXP-2044-rom2-autocast-scope/) |

### R2-ASSET-088

RU selected loader sequence pushes the dialogs path at L2.01547 and
loads receiver L2.01500 through L2.01548. The acquired .data ASCII-string
pass identifies that path as main/text/dialogs.txt. Options construction uses
that receiver with indices 167..174 for target choices, levels and a combo
caption/tooltip argument. R2-ASSET-085 supplies the short EN/RU source labels;
new source row 174 is also measured, without asserting its visible role.

The same selected loader pushes L2.01549 to receiver L2.01550.
The bounded NUL-terminated string read at that address gives patch/patch.txt.
Accessor L2.01551 adds the argument to the pool's base index. The selected key
arm uses receiver L2.01550 with indices 0x73..0x75, not the main.txt pool.
No patch archive node or patch text file was acquired.

**Confidence.** High for local receiver/path operands, the explicit patch
path string and requested indices. Medium for complete source-to-native-label
binding and caption/tooltip role; live pool mutation and encoding are not
closed. Unknown for actual patch feedback text and rendered glyphs.

**Unknown.** Patch source values, runtime text conversion and display.
Acquire the named patch text population or inspect an owner screenshot after
the named cycle input. The unrelated main.txt autohealing labels cannot be
substituted as this feedback evidence.

## Shared inn lower-right literal resources

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-089 | Selected EN/RU shared loader chains bind HumanBackL.bmp and TextBackL.bmp to the two globals used by the druid inn lower-right crop; allocation and delivered pixels remain Unknown. | High / Unknown | ● active (branch candidate) | [EXP-2045](../experiments/EXP-2045-rom2-druid-town-remainder/) |

### R2-ASSET-089

The selected EN loader pushes `graphics\interface\HumanBackL.bmp`
at `L2.01601`, calls constructor `L2.00573` at `L2.01602`, and stores
its result at global `L2.00747` at `L2.01603`. The second chain pushes
`graphics\interface\TextBackL.bmp` at `L2.01604`, calls the same
constructor at `L2.01605`, and stores at `L2.00748` at `L2.01606`.
The allocation-null arms store zero instead.

RU pushes the same two literal keys at `L2.01607` and `L2.01608`, calls
`L2.01609` at `L2.01610` and `L2.01611`, and stores at globals `L2.01612`
and `L2.01613` at `L2.01614` and `L2.01615`. R2-ENGINE-368 maps the first
global to nonzero child+0x70 and the second to zero, with the same crop.

The four strings are in `.data`, at EN `L2.01616` / `L2.01617` and RU
`L2.01618` / `L2.01619`. A bounded 128-byte file window at each address
finds the NUL-terminated ASCII key. The first pair of window SHA256 hashes
agrees between locales, as does the second pair; this is not an archive
payload or whole-image equality claim.

Evidence: `evidence/measured/g2/` call-minus-8/11/16 listings preserve all
failed selection attempts and the constructor/store operands.
`evidence/measured/g3/strings.tsv` holds the literal keys and bounded
window hashes. The G2 selector rejected `.data` and emitted no strings.
G3 uses committed G1 section metadata and prior full-file identities,
opens each native file for exactly one 128-byte seek/read, and checks size.
It does not establish an independent current whole-file identity.

**Confidence.** High for the selected key pushes, result/store chains and
bounded literal contents. Unknown for allocation, later replacement and
delivered pixels. The leading partial instruction in some retained
attempts supports no fact; the named push/call/store sequence is decoded.

**Unknown.** Constructor failure, later global replacement, actual archive
payload selection and live crop appearance. Trace these bounded dependencies
or have the owner toggle the inn panel and observe its lower-right crop.

## Generator presentation

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-090 | ROM2 graphics.res has 24 interface BMP keys per locale matching basename button/buttonsarea/panel; their EN/RU pairs agree, with 160x238 panels and Inn or four-button shop image sets. | High | ● active (branch candidate) | [EXP-2048](../experiments/EXP-2048-rom2-generator-look/) |

### R2-ASSET-090

The full graphics.res registry is traversed in each locale. Reduction selects interface/ entries whose basename contains button, buttonsarea or panel; all 24 hits per locale are BMPs and every selected header was read. EN/RU payload SHA-256 agrees for every pair. The header safety cap48 is not reached.

Panels interface/chrgen/buttonsarea.bmp and interface/inn/buttonsarea.bmp are 160x238. Inn button{1,2,3}{on,off}.bmp are 140x46; lbutton/rbutton{on,off}.bmp are 32x28. The generic interface/shopbutton{1..4}.bmp, interface/shop_druid/shopbutton{1..4}.bmp and interface/shop_kaarg/shopbutton{1..4}.bmp use 120x52 for 1/4 and 140x46 for 2/3. All measured BMPs are 24bpp.

R2-ENGINE-286 and R2-ENGINE-246 bind three generator/selected inn slots; R2-ENGINE-249 binds four selected generic shop slots. R2-ENGINE-387 proves native generator captions are drawn separately. No absence of additional words in art is asserted.

**Confidence.** High for the exact filtered registry population, headers and equality. The RES traversal validates registry ownership and extents; unnamed filtered hits are retained. No function-discovery instrument is involved. The reduction excludes names without the three substrings and resources outside graphics.res:interface/, so it cannot exclude other command panels.

**Unknown.** Other panel names and archive roots, all town/menu consumers, druid/Kaarg command-slot binding and shop caption font/ink. A loader-key census and native caption painters would settle them.

## Global map resources

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-093 | ROM2 global-map art is GMap.bmp, Umoir.bmp and text/globalmap.txt in main.res and 17 global.map/ members in graphics.res; PathMap.bmp (EN, RU) and UmoirPath.bmp (EN) are 640x480 8-bit images with 100 and 81 index-2 pixels. | High / Medium | ● active (amended) | [EXP-2054](../experiments/EXP-2054-rom2-global-map/) |
| R2-ASSET-109 | ROM2 graphics.res `global.map/cross/sprites.16a` is 6690 bytes at offset 25578479, byte-equal in EN and RU, with 13 frames and its own palette. | High | ● active | [EXP-2058](../experiments/EXP-2058-rom2-map-presentation/RESULTS.md) |

### R2-ASSET-093

`main.res` (RU registry, EN offsets from the same reader):
`graphics/global.map/gmap.bmp` 921654 bytes in both roots,
`graphics/global.map/umoir.bmp` 921656 EN and 921654 RU, and
`text/globalmap.txt` 5179 bytes EN (184 CRLF rows, SHA256
`32e720979da07860fdc570436d0900c839de20b98ebe8e654a109586fc6baa21`) and
5403 RU (185 rows, the last empty, SHA256
`5a39b88acd42878875eddb73998ee8c1792a474ac64b8880e55cd8fab5294fdc`).

`graphics.res` lists 17 `global.map/` members at equal offsets and sizes
in both roots: ballmap, flag, hero, pathmap, umoirpath, scroll01..03 and
scrollp1..3 (`.bmp`), and cross, flag, flag1, flagonmap, missionflag and
yourflag (`sprites.16a`). `flag.bmp` has no path in the read string block of
the map (R2-ENGINE-429).

| Member | Size | Header | Index 0..3 pixels | SHA256 |
|---|---|---|---|---|
| `pathmap.bmp` (EN, RU) | 307272 | 640x480, 8 bit | 304548, 2552, 100, 0 | `d6f4f1a3c590fd6e66a11f8d74b32eaa072dd2241f16de2e53e89f298096e988` |
| `umoirpath.bmp` (EN) | 308278 | 640x480, 8 bit | 304042, 3077, 81, 0 | `c2b9357a647c36bcd253b70b7e609d2baca2e6c9e2db2ba5e95660bf789d2b86` |

The node coordinates are `out/code/nodes.tsv`.

**Confidence.** High for the RU registry listings, the path images and
the text files read by the final plan. Medium for the EN registry
listing, a navigation receipt (`evidence/navigation/nav2-registry.tsv`);
the EN members read at its offsets have the sizes it lists. Medium for
the backgrounds: their 640x480 24-bit headers
(EN SHA256 `4d4533eafc1772e4a1f84b54f3ced0ed9fb0342d3d85db8e937a58865f3ebb58`
GMap, `f80aaaa994886ae5fddd27b16b6caa93efdba229840ca17fadf22180b4db7c4e`
Umoir) are a navigation receipt (`evidence/navigation/nav3-members.tsv`).

**Unknown.** RU UmoirPath.bmp and the backgrounds' RU bytes; the .16a frame
counts.

**Amended.** R2-ASSET-109 gives Cross's frame count and its EN/RU byte
equality; the other five .16a frame counts stay Unknown.

**Evidence.** [EXP-2054](../experiments/EXP-2054-rom2-global-map/)

### R2-ASSET-109

The member is read at the registry offset and size R2-ASSET-093 lists, in
both roots. Its SHA256 is
`62d40f76601af4952822627cd776d0d6224050ba2e6a5818dc2542644aa005ab` in EN
and RU. Its last dword holds the frame count in the low 31 bits, 13, and
the own-palette flag in the top bit, set (the trailer R2-ASSET-008
describes and loader `L2.00643` reads, R2-ASSET-066). The map draws Cross with frames 0..12 during a travel and
holds frame 12 (R2-ENGINE-464).

**Confidence.** High for the size, hash and trailer values in both roots.

**Unknown.** The palette's colours and the frames' sizes.

**Evidence.** [EXP-2058](../experiments/EXP-2058-rom2-map-presentation/RESULTS.md)

## Main menu and save page resources

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-097 | ROM2 main.res MenuMask.bmp is a byte-equal 640x480 8-bit image in EN and RU whose values 128..240 in steps of 16 mark eight button regions (screen position Medium); the menu also loads menu_.bmp, button, text and sprite pictures. | High / Medium | ● active | [EXP-2055](../experiments/EXP-2055-rom2-menu-load/) |
| R2-ASSET-098 | ROM2 dialogs.txt holds 176 rows in EN and RU with the menu, load and save page texts at the same rows; patch.txt holds 119 EN and 214 RU rows, with the delete confirmation at rows 50 and 51 and the web list at rows 102 and 103. | High / Medium | ● active | [EXP-2055](../experiments/EXP-2055-rom2-menu-load/) |

### R2-ASSET-097

The menu art body EN `L2.01993` loads
`main\graphics\MainMenu\MenuMask.bmp`, `menu_.bmp`,
`graphics\interface\sprites.16a` and, for eight buttons,
`button%d.bmp` (normal), `button%dp.bmp` (pressed) and `text%d.bmp`.
main.res `graphics/mainmenu/menumask.bmp` is 308280 bytes in both
roots with SHA256
`c0fba12987dfca80f5406fe41129e556ee853cb1851eacf848cb29720af8700e`
(`out/code/members.tsv`). Its value regions (pixel count summary
`bbox8`, top-origin x0,y0-x1,y1):

| value | button | box |
|---:|---:|---|
| 128 | 0 | 203,49-303,136 |
| 144 | 1 | 123,153-231,220 |
| 160 | 2 | 124,250-220,331 |
| 176 | 3 | 207,340-310,427 |
| 192 | 4 | 343,54-427,142 |
| 208 | 5 | 424,150-501,227 |
| 224 | 6 | 411,258-508,339 |
| 240 | 7 | 337,345-417,422 |

Value 0 fills the rest. Values 16..30 form small regions that the
mouse rule maps to no button (R2-ENGINE-436).

**Confidence.** High for the paths, the hash, the size and the value
boxes as top-origin pixel counts. Medium for the box position on screen:
the loader's row order and the base constructor EN `L2.01470` are
unread, and the hit test reads the surface at x + 640*y.

**Unknown.** The button pictures' content and size were not read.

**Evidence.** [EXP-2055](../experiments/EXP-2055-rom2-menu-load/)

### R2-ASSET-098

main.res `text/dialogs.txt` is 3117 bytes EN and 3248 bytes RU, 176
rows each, every row ending in CRLF (`out/code/members.tsv` counts the
empty piece after the last CRLF as a 177th). The screens use rows 0 and 1
(Ok, Cancel), 24..28 (load page texts and the Empty Slot), 29..33
(save page texts), 34, 35, 37, 40 and 77 (in-game menu buttons), 144,
151, 153 (View Cutscenes), 157 and 158 (Delete and its tip). patch.res
`patch.txt` is 2786 bytes EN with 119 rows and 7521 bytes RU with 214
rows, counted the same way. EN rows 50 and 51 are the two halves of the
delete confirmation around the label; EN rows 102 and 103 are the web
list's status text and its address prompt. The row texts are not
quoted.

**Confidence.** High for the row counts, sizes and EN roles. Medium for
the RU roles: they rest on the same indices in RU bodies equal to EN.

**Unknown.** What the RU-only patch.txt rows 119..213 hold.

**Evidence.** [EXP-2055](../experiments/EXP-2055-rom2-menu-load/)

## Spell presentation resources

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-101 | ROM2 graphics.res projectiles.reg is byte-equal in EN and RU and defines 31 effect types by ID, File, Phases and draw flags; the 29 projectile .16a sprites are equal between roots and five of them no type names. | High | ● active | [EXP-2056](../experiments/EXP-2056-rom2-spell-presentation/) |
| R2-ASSET-102 | ROM2 sfx.reg is byte-equal in EN and RU; it names magic sounds at Sfx90 and at 30 slots from Sfx508 to Sfx566, and the 33 magic WAV keys are equal mono 16-bit 22050 Hz PCM in both roots. | High | ● active | [EXP-2056](../experiments/EXP-2056-rom2-spell-presentation/) |
| R2-ASSET-103 | ROM2 main.res spell.txt holds 29 spell names in spell-row order and spells.txt holds 24 hover entries whose spellbook-slot order is Medium (the names are not committed); both have the same line counts in EN and RU and differ in text bytes. | High / Medium | ● active (amended) | [EXP-2056](../experiments/EXP-2056-rom2-spell-presentation/) |

### R2-ASSET-101

`graphics.res` node `projectiles/projectiles.reg` is 9584 bytes with SHA256
`97dc39049a41b3709b0d1e199e3565ac2c6962a6bd1ce8af9316e6a159dd187a` in both
roots. It is an inline REG store (R2-ASSET-006): `Global/Count` = 31 and
sections `Projectile0`..`Projectile30`. The native loader (EN `L2.02132`,
RU in `L2.02133+0x800`) reads `File` and `A16`, then `ID` (default -1),
`Phases` (-1), `RotationPhases` (16), `Width` (64), `Height` (64), `Palette`,
`Homing`, `Flip` and `SFX` (0 each) into a 0x3c-byte type object and adds it
to the type map (EN `L2.02134`, RU `L2.02135`) under `ID`.

| ID | File | Phases | Size | RotationPhases | Homing | Flip |
|---:|---|---:|---|---:|---:|---:|
| 1, 2, 3 | archer, xbowman, orc `arrow` | 1 | 64x64 | 16 | 0 | 0 |
| 4 | goblin `arrow` | 1 | 64x64 | 1 | 0 | 0 |
| 5, 6, 7 | catap1, catap2, catap2 `sprites` | 1 | 64x64 | 16 | 0 | 0 |
| 8 | steam | 9 | 64x96 | 1 | 0 | 0 |
| 10 | firebolt | 4 | 64x64 | 16 | 1 | 1 |
| 12 | fireball | 4 | 128x128 | 16 | 0 | 1 |
| 13 | fireexpl | 11 | 128x128 | 1 | 0 | 0 |
| 15 | firewall | 11 | 128x128 | 1 | 0 | 0 |
| 16, 24, 34, 46 | p_fire, p_water, p_air, p_earth | 6 | 12x12 | 1 | 1 | 0 |
| 18 | icemissile | 6 | 16x16 | 1 | 1 | 0 |
| 20 | poison_d | 8 | 12x12 | 1 | 0 | 0 |
| 21 | poison | 15 | 64x64 | 1 | 0 | 0 |
| 23 | Meteor | 9 | 64x64 | 1 | 0 | 0 |
| 27 | acid | 9 | 64x96 | 1 | 0 | 0 |
| 28 | lightnin | 5 | 16x16 | 1 | 1 | 0 |
| 30 | chain | 35 | 16x16 | 1 | 1 | 0 |
| 40 | meteorit | 4 | 64x64 | 16 | 1 | 1 |
| 43 | wall | 1 | 32x64 | 1 | 0 | 0 |
| 48, 62, 64 | bless, shield, Curse | 5 | 12x12 | 1 | 1 | 0 |
| 54 | teleport | 21 | 96x96 | 1 | 0 | 0 |
| 56 | healing | 7 | 16x16 | 1 | 1 | 0 |
| 60 | Drain | 9 | 12x12 | 1 | 1 | 0 |

Every type from ID 8 up sets `Palette` 1 and `A16` 1; `SFX` is 1 for every
type from ID 8 up except 23 (Meteor). The `File` value is joined under
`graphics\projectiles\` with `.16a` when `A16` is set, else `.256` and
`b.256`.

`graphics.res` has 46 `projectiles/` keys in each root. The 29 `.16a`
sprites read (R2-ASSET-007 envelope) all set the palette bit and are
byte-equal between roots. Types 10, 12 and 40 (RotationPhases 16 with Flip)
have 36 frames, Phases 4 times nine directions. Every other named `.16a`
has Phases frames, except `healing` with 8 frames for Phases 7. Five `.16a`
sprites no type names: `diamonddust` (5 frames, 32x32), `freeze` (15,
64x64), `smallxpl` (11, 64x64), `smoke0` (6, 12x12) and `smoke1` (6, 24x24).

**Confidence.** High for the registry values, defaults, keys, frame counts
and root equality.

**Unknown.** Whether the executable names `diamonddust`, `freeze` or
`smallxpl` outside this registry; no executable string census was made.

**Evidence.** [EXP-2056](../experiments/EXP-2056-rom2-spell-presentation/)

### R2-ASSET-102

`sfx.res` node `sfx.reg` is 9864 bytes with SHA256
`ee0ecdc9bfc118e74ba79d446d3ee63661a0240e8ed9d2405e6ea0aaa8da2b06` in both
roots. `Global/SfxCount` is 660; the `Sfx` section holds 200 `Sfx<N>`
string entries, each a path under `sfx.res` without the `.wav` suffix. The
31 entries under `magic\` are:

| N | Key | N | Key | N | Key |
|---:|---|---:|---|---:|---|
| 90 | firewall | 528 | lightning | 548 | bless |
| 508 | ex-water | 530 | spray | 550 | haste |
| 510 | arrow | 532 | invisib | 552 | spirit |
| 512 | fireball | 534 | plight | 554 | teleport |
| 513 | explosion | 536 | darkness | 556 | heal |
| 516 | pfire | 538 | light | 558 | summon |
| 518 | icemissile | 540 | dust | 560 | drain |
| 520 | poison | 542 | wall | 562 | pfield |
| 523 | blizzard | 544 | stone | 564 | curse |
| 524 | pwater | 546 | pearth | 566 | slow |
| 526 | acid | | | | |

Between 508 and 566 the even slots 514 and 522 and every odd slot except
513 and 523 are absent. `sfx.res` holds 33 `magic/` WAV keys per root; no
`sfx.reg` entry names `freeze` or `sacrifice`. Every one is RIFF WAVE format
tag 1, one channel, 16 bits, 22050 Hz, byte-equal between roots. Durations
in milliseconds (data bytes over byte rate): acid 1892, arrow 505, bless
1949, blizzard 485, curse 1147, darkness 1738, drain 1403, dust 791,
ex-water 1799, explosion 856, fireball 888, firewall 1623, freeze 2559,
haste 946, heal 2093, icemissile 621, invisib 1614, light 1425, lightning
942, pearth 1065, pfield 1413, pfire 1025, plight 1100, poison 2552, pwater
696, sacrifice 1973, slow 1102, spirit 1929, spray 1384, stone 526, summon
516, teleport 1071, wall 881.

**Confidence.** High for the registry entries, the archive population, the
WAVE fields and root equality.

**Evidence.** [EXP-2056](../experiments/EXP-2056-rom2-spell-presentation/)

### R2-ASSET-103

`main.res` `text/spell.txt` splits on CRLF into 30 lines in both roots:
lines 0..28 are nonempty and line 29 is empty. EN line i names the client
`data.bin` Spells row i+1 for 26 of the 29 rows; line 1 reads `Fireball`
where row 2 is `Fire Ball`, line 15 reads `Stone Missile` where row 16 is
`Diamond Dust`, and line 21 reads `Animate Dead` where row 22 is `Control
Spirit`. RU line lengths differ from EN.

`text/spells.txt` splits into 25 lines in both roots: lines 0..23 are
nonempty and line 24 is empty. Each EN line is a spell name with a
parenthesized kind, a `#`, and a description. EN line order: Fire Arrow,
Fireball, Wall of Fire, Protection from Fire, Heal, Bless, Haste, Drain
Life, Protection from Air, Invisibility, Prismatic Spray, Lightning, Ice
Missile, Poison Cloud, Blizzard, Protection from Water, Summon, Animated
Dead, Teleport, Shield, Protection from Earth, Stone Curse, Wall of Earth,
Stone Missile. No line names Acid Stream, Darkness, Light, Curse or Slow.

`text/main.txt` EN lines used by spell text: 90 `of`, 92 `casts`, 91 and
93 empty, 117 `Mana cost`, 118 `Damage`, 123 `Range`, 124 `Duration`, 182
`Speed`, 183 `Resistance`, 184 `Sight`, 185 `Maximum Damage Probability`,
186 `rays`, 187 `Minimum Damage Probability`, 192 `Spells: `. RU lines 91
and 93 are also empty.

**Confidence.** High for line counts, empty lines, lengths, hashes and the
EN `spell.txt` and `main.txt` labels in the committed evidence. Medium for
the `spells.txt` line shape and name order, which were read from a
navigation pass whose text stays outside the repository. R2-ENGINE-444
supplies a slot table whose values fit that order; that fit is
consistency, not an independent read, so the join of hover lines to spell
rows is Medium.

**Unknown.** The 24 `spells.txt` names as a committed list, and the
consumer EN `L2.02136` of the 'A' key handler.

**Amended.** R2-ASSET-177 commits six `spells.txt` lines (slots 3, 5, 8, 15,
19 and 20) and joins them to spell rows by equal names in both roots; the
other 18 lines stay uncommitted.

**Evidence.** [EXP-2056](../experiments/EXP-2056-rom2-spell-presentation/)

## Spell rows 5, 16, 18 and 25 and the Summon units

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-107 | ROM2 client Spells rows 5 and 16 equal row 1 except Sphere (2, 4); row 18's Effects string opens with a quote before absorbtion; Units rows 109..123 hold Squirrel, Foot_Animated and Turtle, bare and .2 to .5. | High | ● active | [EXP-2057](../experiments/EXP-2057-rom2-spell-rules/RESULTS.md) |

### R2-ASSET-107

The client `data/data.bin` inside `world.res` (RU read in full; the EN and RU
world.res and data.bin hashes are equal, R2-ENGINE-395) holds 34 Spells rows
and 241 Units rows through the grammar of R2-ASSET-029. Parameter n carries
title n+1.

| Row | Name | Cost | Sphere | Target | Delivery | Max Range | Effect Speed | Distribution | damage | Defensive | Effects |
|---|---|---|---|---|---|---|---|---|---|---|---|
| 1 | Fire Arrow | 3 | 1 | 1 | 2 | 7 | 200 | 1 | 5..10 | -1 | empty |
| 5 | Ice Missile | 3 | 2 | 1 | 2 | 7 | 200 | 1 | 5..10 | -1 | empty |
| 16 | Diamond Dust | 3 | 4 | 1 | 2 | 7 | 200 | 1 | 5..10 | -1 | empty |
| 18 | Stone Curse | 40 | 4 | 1 | 1 | 5 | 0 | 1 | -1..-1 | -1 | two pieces |
| 25 | Summon | 100 | 5 | 2 | 1 | 0 | 0 | 1 | -1..-1 | -1 | empty |

Rows 1, 5 and 16 agree in all 22 stored values except Sphere. Row 18's Effects
string has two comma pieces: the first opens with byte 0x22 (a quote), key
`"absorbtion`, value `+5:duration 8`; the second is key `defence`, value
`-20:duration 8` followed by a closing quote. Row 18 Spell Duration is 10 and
Apply on Unit Method 2.

Units rows 109..113 are Squirrel, Squirrel.2 .. Squirrel.5; 114..118
Foot_Animated and .2 .. .5; 119..123 Turtle and .2 .. .5. Units titles 49..54
are Spell 1, Probability 1, Spell 2, Probability 2, Spell 3, Probability 3.

**Confidence.** High: the committed derived rows reproduce byte for byte, and
the archive and table grammar leave zero residue.

## Weapons attack fields

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-111 | ROM2 client Weapons rows (27) give attackType 1 for swords and the dagger, 2 axes, 3 clubs, maces, hammers, staves and BareHands, 4 pole arms, 5 bows and the crossbow; range is 4 to 6 for bows, crossbow and staves, -1 for melee. | High / Medium | ● active | [EXP-2060](../experiments/EXP-2060-rom2-physical-combat/) |

### R2-ASSET-111

The client `data/data.bin` inside `world.res` (RU read; the EN and RU
world.res and data.bin hashes are equal, `out/inputs.tsv` of the experiment;
R2-ENGINE-395 corroborates) holds 27 Weapons
rows under the shared Armors/Shields/Weapons titles Item, Shape, Material,
Price, weight, Slot, @.attackType, @.physicalMin, @.physicalMax, @.toHit,
#.deIrnce, #.absorbtion, @.range, @.charge, @.relax, 2 handed, sutableFor and
Other Parameter (titles 0 to 17), read through the grammar of R2-ASSET-029.

- attackType 1: Dagger, Short Sword, Long Sword, Bastard Sword, Two Handed
  Sword, Plasma Sword; 2: Axe, Two Handed Axe; 3: BareHands, Club, Spiked
  Club, Mace, Morning Star, Pick Hammer, War Hammer, Staff, Shaman Staff;
  4: Pike, Halberd, Lance; 5: Short Bow, Long Bow, Crossbow, Sonic Beam,
  Boulder Thrower; 11: Flame Thrower. Row 23 (`rem`) holds -1 in every
  field.
- range: Short Bow 4, Long Bow 5, Crossbow 6, Staff 6, Shaman Staff 5,
  Sonic Beam 4, Flame Thrower 8, Boulder Thrower 20; every other row -1.
- The two staves carry physicalMin and physicalMax 0; BareHands 1 and 2.
- 2 handed is 2 for the Two Handed Sword, War Hammer, Two Handed Axe, both
  staves, the two bows and the crossbow, else -1.

The RU weapon fill `L2.02311` reads parameters 6, 7, 8, 9, 3 and 11 and
equip reads parameter 5 (R2-ENGINE-479); with parameter n carrying title n+1
(R2-ASSET-107) these are physicalMin, physicalMax, toHit, defence, weight,
range and attackType.

**Confidence.** High for the row values: committed derived rows reproduced
from the Git blobs. Medium for the parameter-to-title binding in the
Weapons table, which rests on R2-ASSET-107's rule and the eight parameter
reads agreeing with it; the parameter accessor `L2.02312` was not read.

**Evidence.** [EXP-2060](../experiments/EXP-2060-rom2-physical-combat/)

## Inn dialogue sections

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-115 | ROM2 text/town.txt finds every inn talk key exactly, holds about for hires 1..7 and accept/reject by hire and stage; RU lacks npc520talk7, its #npc2006talk110 is not line-leading, and five talk sections differ in headers. | High | ● active | [EXP-2062](../experiments/EXP-2062-rom2-inn-talk/) |

### R2-ASSET-115

The census reads `main.res:text/town.txt` in each root (EN 103816 bytes,
SHA256 `0b934e34987a372e55667f8aa83bfbd6f7e320b9099a4fb4d36020ca886880e7`;
RU 99082 bytes,
`a24079a569bf3e979681846ec16c1df8eb867cdcbad5c392b9f75a5248e24f56`). It
lists line-leading `#` tags matching `npc<digits><letters><digits>` and,
for each of the 46 talk keys the EnterInn constants form
(R2-ENGINE-494), the first `strstr`-shaped hit of `#<key>`.

| Family | EN sections | RU sections | Keys |
|---|---:|---:|---|
| talk | 47 | 45 | the 46 option keys; EN also `npc520talk7` |
| about | 7 | 7 | `npc1about`..`npc7about` |
| accept | 41 | 41 | hires 1, 2 at stages 30..110; 3, 6 at 60..110; 4, 5 at 70..110; 7 at 110 |
| reject | 41 | 41 | as accept |

In EN all 46 keys hit their own line-leading tag; none hits inside a longer
key. In RU 45 do. The first RU `#npc2006talk110` is at offset 98402 and is
not at a line start; the census assigns that offset to the
`npc2005talk103` span. RU lacks `npc520talk7`, which no EnterInn option
names. Header counts differ in `npc22talk40`, `npc2111talk42`,
`npc2022talk84`, `npc2003talk92` and `npc2005talk103`, the last because the
RU span includes the `npc2006talk110` headers. RU `npc5accept100` has one
`sound=` header. No header in either root's census carries `npcalive=`,
`npcdead=` or `tips=`.

**Confidence.** High for the counts, keys and offsets: a committed derived
census reproduced from Git blobs. The native lookup is not line-anchored
(R2-ENGINE-050); which RU body that lookup cuts for `npc2006talk110` is not
claimed.

**Evidence.** [EXP-2062](../experiments/EXP-2062-rom2-inn-talk/)

## Map records 10 and 11 and the unit record

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-119 | ROM2 ALM type 6: the loader moves each field of the 48-byte unit record to a fixed slot of a 0x54-byte row, equal in EN and RU allods2.exe; offsets, widths and slots are tabled for versions 1300 and 1600. | High | ✔ promoted (branch candidate) | [EXP-2064](../experiments/EXP-2064-rom2-map-records/) |
| R2-ASSET-120 | ROM2 ALM type 10 entries are four i32: group number, a percent scaling the slot-5 countdown reload value, flags 1/2/4/8 and a type-7 node position; the loader copies each entry whole and the group pass reads all four. | High / Medium | ✔ promoted (amended, branch candidate) | [EXP-2064](../experiments/EXP-2064-rom2-map-records/) |
| R2-ASSET-121 | ROM2 ALM type 11 arrays key on a type-4 object id (file word 0x12): 11a holds flags and a word, 11b twenty i32 in four groups of five, 11c a nonzero gate and a type-7 node position. | High / Unknown | ✔ promoted (amended, branch candidate) | [EXP-2064](../experiments/EXP-2064-rom2-map-records/) |
| R2-ASSET-122 | Over the 46 scenario.res campaign maps per root, types 10, 11a and 11b are empty, 11c has 40 entries in 10 maps and type-0 dword 0x84 is 1; EN and RU differ in dword 0x88 on 3 maps, in version on 4 and in unit counts. | High | ✔ promoted (amended, branch candidate) | [EXP-2064](../experiments/EXP-2064-rom2-map-records/) |

### R2-ASSET-119

The type-6 case of the map loader (RU `L2.02482`..`L2.02483` in the
dispatcher `R2.0002`; EN dispatcher `L2.02484`) reads each record into a
fresh 0x54-byte row and appends the row to map+0x3a4 (`L2.02485`). File
offsets are for the 48-byte record at versions 1300 and 1600, the two
versions in the installed campaign maps; the version-conditioned rows
below hold at those versions only, and the block read below version 0x44c
uses other fields (R2-ASSET-024).

| File offset | Width | Row slot | Transfer |
|---:|---|---:|---|
| 0x00 | i32 | +0x00 | x; the actor builder shifts it right by 8 into `L2.00313`, unread |
| 0x04 | i32 | +0x04 | y; shifted likewise |
| 0x08 | i16 | +0x08 | sign-extended to a dword |
| 0x0a | u16 | +0x0c | word |
| 0x0c | u32 | +0x48 | flags; bit 0x10 selects the person arm (R2-ENGINE-038) |
| 0x10 | u32 | +0x50 | read at version 0x47e and above, else 0 |
| 0x14 | u32 | +0x10 | definition id (R2-ENGINE-038); read at version 0x3db and above, else 0 |
| 0x18 | u32 | +0x14 | low word only; the high word is dropped |
| 0x1c | u32 | +0x44 | stored minus 1 |
| 0x20 | u32 | +0x1c | dword |
| 0x24 | i16 | +0x20 | word |
| 0x26 | i16 | +0x22 | word |
| 0x28 | u32 | +0x3c | four bytes at version 0x456 and above, else two |
| 0x2c | u32 | +0x40 | group number |

The three thresholds are tested as `cmp`, `jb`: below the value the field
is skipped (0x47e, 0x3db) or the two-byte form is read (0x456), so the
value itself takes the later form (`L2.02479`..`L2.02486`,
`L2.02480`..`L2.02487`, `L2.02481`..`L2.02488`).

After each row the loader sets map+0x90 to the largest row +0x40. When row
+0x44 is not -1 and is below the type-8 count, the type-8 element it
indexes (map+0x2e8, getter `L2.02489`) receives row WORD +0x3c,
sign-extended, at its +0x3c; otherwise row +0x44 becomes -1. The readers of
the row are in R2-ENGINE-513.

**Confidence.** High: the RU dispatcher body was read, and the EN body
matches it instruction for instruction after address masking (1623 of 1623
instructions). Medium that x and y are in 1/256 tile: the shift is read,
its consumer `L2.00313` is not, and R2-ENGINE-336 grades the same scale
Medium. Field names other than the transfers and the cited claims are not
established here.

### R2-ASSET-120

The loader reads each 16-byte type-10 entry in one read into a 16-byte
allocation and tail-appends it to map+0x310; it splits no field. The
type-10 group pass (R2-ENGINE-510) reads the fields of every entry whose
e0 matches a group:

| Offset | Field | Use |
|---:|---|---|
| 0x00 | e0 | compared with the group number (unit record 0x2c, R2-ASSET-119) |
| 0x04 | e1 | object +0xb0 and +0xb4 = e1 * S / 100, S the scale of R2-ENGINE-510; +0xb0 is the reload value of the slot-5 countdown on +0xb4 (Medium) |
| 0x08 | e2 | bit 1: object +0xb8 from e3, else -1; bit 4: object +0xc4 and group +0x48 = 1, else 0; bit 8: object +0xc8 and group +0x4c likewise; bit 2 clear: object +0xbc = 0 |
| 0x0c | e3 | type-7 first-array position: the number of leading nodes scanned; object +0xb8 = count of those with dword +0x40 below 0x10002, passed to `L2.02490` |

Every entry whose e0 matches applies in file order; a later entry
overwrites the fields an earlier one set. A group with no matching entry
keeps the object defaults of R2-ENGINE-510. The pass runs under the gate
of R2-ENGINE-509, whose "only" clause is Medium. The bit masks are the
dwords 1, 2, 4, 8 at RU `L2.02491` and EN `L2.02492`.

**Confidence.** High for the field use: the pass bodies RU `L2.02493` and
EN `L2.02494` were read in both roots. Medium that no other code reads the
entries: the census of map+0x310 is a displacement scan over the RU image
and a call-target scan of the RU element getters `L2.02495` and
`L2.02496`; its hits outside the map ctors, dispatcher, teardown and two
consumer passes were classified after the receipt was read, and rows on
other bases at 0x310..0x31c stay unclassified. Medium for the role of
+0xb0 as a countdown reload value: it rests on the slot-5 body, which
R2-ENGINE-510 grades Medium. The node-iterating helpers `L2.02497` and
`L2.02498` of the e3 count are unread.

**Unknown.** The EN image was not scanned for other readers of map+0x310.

**Amended.** R2-ENGINE-537: group +0x48 and +0x4c, set by e2 bits 4 and 8, select the groups of Inn offer blocks 3 and 5. R2-ENGINE-538 reads `L2.02502`. R2-ASSET-132 gives the root-map values.

### R2-ASSET-121

The loader reads each type-11 entry whole, in three loops, and appends
11a (12 bytes) to map+0x338, 11b (84 bytes) to map+0x324 and 11c
(12 bytes) to map+0x34c. The type-4 object pass (R2-ENGINE-512) matches
entry e0 against the type-4 element id at element +0x10, which the
element ctor (RU `L2.02499`, EN `L2.02500`) takes from type-4 file word
0x12, sign-extended. The same ctor puts file word 0x0c at element +6 and
the low word of file dword 0x0e at element +8.

| Array | Offset | Use |
|---|---:|---|
| 11a | 0x00 | type-4 id |
| 11a | 0x04 | flags: 2 enables the 0x08 word; 0x10, 0x20, 0x40 and 4 each set one object field |
| 11a | 0x08 | low word, used only under flag 2 |
| 11b | 0x00 | type-4 id |
| 11b | 0x04..0x53 | five runs of four i32 (0x04, 0x14, 0x24, 0x34, 0x44), interleaved into four object groups of five |
| 11c | 0x00 | type-4 id |
| 11c | 0x04 | read only as nonzero; a zero entry is skipped |
| 11c | 0x08 | number of leading type-7 first-array nodes scanned |

The type-10 e3 and type-11c 0x08 fields refer to positions in the type-7
first array (map+0x40). The reverse reference was searched by a
displacement scan of map+0x310..+0x35f over the RU image: its classified
hits are the three map ctors, the dispatcher, the teardown, unwind
funclets and the two consumer passes, with no script compiler or action
site among them.

**Confidence.** High for the layout and the id join: the loader, the
element ctor and the type-4 pass were read in both roots. Unknown for what
the receiving objects do with these fields (R2-ENGINE-512). The node-iterating helpers
`L2.02497` and `L2.02498` of the 11c count are unread.

**Unknown.** Whether a script record refers to a type-10 or type-11
entry stays Medium-bounded for RU and unsearched for EN: scan rows on other
bases were left unclassified, and a pointer to an entry could pass through
a getter. A call-target scan of the eight type-10/11 element getters and a
displacement scan in EN would settle it for EN.

**Amended.** R2-ASSET-131 names the 11b fields (mask, minimum and maximum price, draws and quantity bound per category) and R2-ENGINE-536 a reader of the 11a fields. Flag 1 occurs in the root maps and no field receives it (R2-ASSET-132).

### R2-ASSET-122

The census reads the 46 maps of `scenario.res` in each root (EN 5548580
bytes, SHA256
`388c955f49c02ecfed765340ff3a5b3e8f37210e2c85c8d877bc04b949d60d7f`; RU
5549452 bytes,
`c19950690570d94b40c68acffd210cdab94c2a12b5d6ad5718730392c6b6e08b`) by
the ALM walk of R2-ASSET-002 and splits each type-10 and type-11 entry
into i32 fields at every fourth byte. These are the campaign maps; the
root `.alm` maps (R2-ASSET-084) are outside it.

- 46 maps per root; versions EN 1300 x41 and 1600 x5, RU 1300 x45 and
  1600 x1.
- Types 10, 11a and 11b: 0 entries in every map of both roots.
- Type 11c: 40 entries per root in maps 10, 20, 21, 30, 40, 50, 53, 62,
  76 and 86 (3, 11, 7, 3, 9, 2, 1, 1, 2, 1). 0x00 takes 17 values in
  10..170; 0x04 is 0 or 1; 0x08 takes 8 values in 0..48. Per-map type-10
  and type-11 hashes are equal in EN and RU.
- Type-0 payload dword 0x84 is 1 in all 92 maps. Dword 0x88 is 1 except
  EN 110 (4), 52 (2), 74 (3), 95 (3) and RU 95 (3): the roots differ on
  maps 110, 52 and 74; map 95 is equal. The roots also differ in the
  version of four maps and in unit counts.
- Type 6: EN 4212 records, RU 4234. Distinct values per field (EN; RU
  equal ranges): 0x08 53 in 1..107; 0x0a 6 in 1..6; 0x0c {0, 16};
  0x10 {0, 8}; 0x14 205 in 600..10230; 0x18 9 in 1..9; 0x1c 5 in 0..4;
  0x20 6 in 0..14, 10 in 4199 records; 0x24 and 0x26 -1 or 15..4992;
  0x28 EN 3834 and RU 3856 distinct dwords; 0x2c 39 in 0..38.

**Confidence.** High: a committed derived census reproduced byte for byte
from Git blobs. The counts bound which fields vary in the 46 campaign
maps per root; they are not format limits and do not extend to the root
`.alm` maps.

**Unknown.** The 13 EN and 37 RU root `.alm` maps were not read. The
committed census generator (`evidence/maps.py` through `regen.sh`) run
over those files would settle whether any of them has dword 0x84 above 1
or a type-10 or type-11 entry.

**Amended.** R2-ASSET-130 and R2-ASSET-132 census the root `.alm` maps this card's Unknown names: dword 0x84 is above 1 in all of them, and all carry type-10 and 11b entries.

## Map structure kinds and type-11c entries in the campaign

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-123 | The 46 ROM2 campaign maps per root hold 13 Pointer structures and no Shop or Inn; 35 gate-1 11c entries name 12 Pointers, each with its action node at index position present and of operation 2 (High); run-time selection is Medium. | High / Medium | ✔ promoted (amended, branch candidate) | [EXP-2065](../experiments/EXP-2065-rom2-structure-actions/) |
| R2-ASSET-124 | ROM2 Data.bin Buildings (180 stored entries, EN = RU) names entries 0x46..0x4b "Pointer 1".."Pointer 6", 0x22/0x23 "Shop"/"Shop 2" and the other R2-ENGINE-512 Shop and Inn kinds as Multiplayer shops and inns. | High / Medium | ✔ promoted (branch candidate) | [EXP-2065](../experiments/EXP-2065-rom2-structure-actions/) |

### R2-ASSET-123

Population: the 46 `.alm` members of `scenario.res` in each of the EN and
RU ROM2 roots, walked as R2-ASSET-002/003, type-4 records as R2-ASSET-019
(20 bytes, or 28 when dword 0x08 is 0x21 or has bit 0x1000000; kind is the
low word of dword 0x08; the vocabulary sentence of that claim is amended
by this population, see its Amended paragraph), type-7 as 796-byte action and check nodes and
184-byte triggers, type 11 split as R2-ASSET-121.
Every walk ends exactly at its payload end.

Totals per root: 1721 type-4 records, 1334 action nodes, 986 check nodes,
40 type-11c entries (R2-ASSET-122); triggers 715 EN and 714 RU, the one EN/RU difference
being `52.alm` (25 EN, 24 RU). In all 92 maps the type-4 record count
equals type-0 header dword 8.

Of the kinds R2-ENGINE-512 builds as structures, only Pointer kinds occur:
0x46 three times, 0x48 four, 0x49 one, 0x4a four and 0x4b one, 13 records
per root. Kinds 0x22, 0x23, 0x43..0x45, 0x47, 0x5d..0x5f, 0x63..0x65,
0x69..0x6b and 0x6f..0x71 occur in no map of either root.

The 40 type-11c entries (id, gate, position) are in 10 maps. The 35 with
gate 1 name 12 Pointer records by type-4 id (file word 0x12); repeated entries carry the same id, gate
and position. For each, the type-4 record with that id is unique, action
node `position` exists, and its operation is 2 (instant 2, R2-ENGINE-047).
The count column is the number of action nodes before `position` whose
operation is below 0x10002 compared as unsigned. The type-4 pass and the
compiler count with a signed test (`jge`, R2-ENGINE-520), so the column
equals the value the pass stores only if no action node has operation
0x80000000 or above. The tables list the operation of the selected nodes
only; the operation range over the 1334 action nodes per root is not
committed, so that equality is an assumption.

| map | id | kind | x | y | entries | position | node ID | count |
|---|---|---|---|---|---|---|---|---|
| 10 | 42 | 0x4a | 3712 | 10368 | 3 | 25 | 26 | 24 |
| 20 | 87 | 0x4a | 9600 | 14208 | 4 | 23 | 25 | 22 |
| 20 | 88 | 0x48 | 9088 | 14208 | 3 | 24 | 26 | 23 |
| 20 | 89 | 0x48 | 12160 | 6784 | 4 | 25 | 27 | 24 |
| 21 | 30 | 0x4b | 13184 | 12160 | 2 | 23 | 24 | 22 |
| 21 | 31 | 0x48 | 12416 | 11392 | 3 | 21 | 22 | 20 |
| 21 | 32 | 0x46 | 11648 | 12160 | 2 | 22 | 23 | 21 |
| 30 | 74 | 0x46 | 7040 | 9088 | 3 | 48 | 50 | 47 |
| 40 | 168 | 0x4a | 15232 | 19584 | 4 | 48 | 49 | 47 |
| 40 | 169 | 0x49 | 21376 | 16512 | 3 | 48 | 49 | 47 |
| 40 | 170 | 0x48 | 25728 | 15232 | 2 | 48 | 49 | 47 |
| 50 | 46 | 0x46 | 33408 | 7296 | 2 | 35 | 36 | 34 |

The five gate-0 entries are `53.alm` id 29 (kind 0x4c), `62.alm` id 64
(0x4d), `76.alm` id 13 (0x4d), `76.alm` id 28 (0x4a) and `86.alm` id 10
(0x4c). The thirteenth Pointer, `76.alm` id 28, has only the gate-0 entry.
The three Pointers of `40.alm` select the same node. The tables are equal
in EN and RU.

**Confidence.** High for the counts, the join by node index (the node at
index `position` exists and has operation 2) and the EN/RU equality: one
census generator over the full `scenario.res` of each root, reproduced
against its committed output. Medium that a Pointer selects that node at
run time: it rests on the compiled-slot rule of R2-ENGINE-520 (Medium for a
campaign entry's node) and on the unread `L2.02565`. Medium that the count
column is the stored value, for the signed comparison above.

**Unknown.** The event ID each instant-2 node passes and its text: reading
the operation-2 arm of the executor and tabulating the fields of the 12
selected nodes would settle it. The behaviour of the gate-0 entries: the
copy body of R2-ENGINE-512 read for the gate test. The operation range of
the 1334 action nodes per root: one table from the committed generator
would settle the signed comparison. The root `.alm` maps (R2-ASSET-084) are
not censused; the same generator over them would settle which structure
kinds and 11c entries they hold.

**Amended.** R2-ASSET-130: the root `.alm` maps hold Shop and Inn kinds (every 11a entry names an Inn, every 11b entry but five per root a Shop), and all their 11c entries have gate 0.

### R2-ASSET-124

`world.res` is byte-identical in the EN and RU roots (SHA-256
`863b15ce6a304263d9ed97ca14426db42e07825b91165b39d5adf4054d9dd8a9`). Its
`data/data.bin`, walked with the grammar of R2-ASSET-029 to its exact end,
has a Buildings collection with titles Building, sizeX, sizeY, scanRange,
healthMax, Passability, BuildingPresent, Start ID and Tiles and a stored
count of 180. At the R2-ENGINE-512 structure kinds used as indices:

| index | name | parameters |
|---|---|---|
| 0x22, 0x23 | Shop, Shop 2 | 3,3,4,0,447,511 |
| 0x43..0x45 | Multiplayer Inn 1..3 | 3,3,3,0; 191, 511, 447; 511 |
| 0x46..0x4b | Pointer 1..6 | 1,1,0,0,1,1 |
| 0x5d..0x5f | Multiplayer Kaarg Shop 1..3 | 4,4,5,0; 61439, 28671, 65535; 65535 |
| 0x63..0x65 | Multiplayer Kaarg Inn 1..3 | 3,3,3,0; 447, 191, 511; 511 |
| 0x69..0x6b | Multiplayer Druid Shop 1..3 | 4,3,7,0; 3839, 4095, 4095; 4095 |
| 0x6f..0x71 | Multiplayer Druid Inn 1..3 | 4,3,7,0; 1791, 1791, 2047; 4095 |

The parameter array has six values; where they differ across a range, the
fifth value is listed per entry.

**Confidence.** High for the names and parameters at these indices. Medium
that a type-4 kind indexes Buildings: the names match the structure class
of R2-ENGINE-512 at every listed kind, but no read body was followed from
the kind to the collection.

**Unknown.** The meaning of the parameter values beyond the titles. The
code path from kind to Buildings entry: a call-target scan of the
Buildings collection accessor and a read of its callers in the type-4 pass
and the structure constructors would settle it.

## Hire template rows in world.res

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-125 | ROM2 world.res has hire-keyed rows only for hires 1..7, equal in EN and RU: Humans XBowman, Knight and Mage (typeID 15, 5, 23) and Units Druid, Troll, Orc and Catapult (100, 102, 101, 26). | High | ● active | [EXP-2066](../experiments/EXP-2066-rom2-hire-entry/) |

### R2-ASSET-125

A census of `data/data.bin` in each root's `world.res` lists every Humans
row whose parameter 24 and every Units row whose parameter 55 is in
10000..10399. The EN and RU payloads have one SHA-256, so the rows are
equal. Humans ordinals 263..272 are 2MercenaryXBowmanBase and
2MercenaryXBowman30..110 (serverIDs 10010..10019), 273..282 the Knight rows
(10020..10029) and 283..289 2MercenaryMageBase and 2MercenaryMage60..110
(10060, 10064..10069). Units ordinals 27 and 28 are CatapultBase (10070) and
MercenaryCatapult (10079), 221..227 MercenaryDruidBase and
MercenaryDruid60..110 (10030, 10034..10039), 228..233 the Troll rows (10040,
10045..10049) and 234..240 the Orc rows (10050, 10054..10059). The other rows
in the range are the hero rows Start_MF, Start_FF, Start_MM, Start_FM
(10210..10213), Fighter_Hero (10220) and Mage_Hero (10230). No row has a key
in 10080..10209, so hires 8..20 have no base row and no suffix row. RU
`world_srv.res` adds DM_ rows with serverIDs 10300..10308; EN has no
`world_srv.res`.

The typeID parameter (Humans parameter 16, Units parameter 29, each named by
the title at the following index) is 15 for XBowman, 5 for Knight, 23 for
Mage, 100 for Druid, 102 for Troll, 101 for Orc and 26 for both Catapult
rows. The Humans gender parameter is 1 for XBowman and 0 for Knight and
Mage. For comparison, the Start_ hero rows carry typeID 3 (MF), 14 (FF), 24
(MM) and 24 (FM), and Fighter_Hero 5 and Mage_Hero 23.

**Confidence.** High for the row census and the values: derived from the
committed census over both roots and reproduced byte for byte. The title
naming of each parameter follows R2-ENGINE-319.

**Unknown.** Rows of other archives the game might load in place of
`world.res` (the loader's archive list and a census of each archive's
`data.bin` settle it).

**Evidence.** [EXP-2066](../experiments/EXP-2066-rom2-hire-entry/)

## Map record loading and the root-map census

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-129 | The ROM2 ALM loader reads types 10, 11 and 12 by their type-0 counts in signed loops and never by payload size: a count of 0 or below reads nothing; a count that disagrees with the size misframes later records (Medium). | High / Medium | ✔ promoted (branch candidate) | [EXP-2068](../experiments/EXP-2068-rom2-map-records/) |
| R2-ASSET-130 | The 13 EN and 37 RU ROM2 root .alm maps are version 1600 with every counted payload exact; all hold type-10 and 11b entries, type-0 dword 0x84 is 4..16, and each EN file equals its RU namesake byte for byte. | High / Medium | ✔ promoted (branch candidate) | [EXP-2068](../experiments/EXP-2068-rom2-map-records/) |
| R2-ASSET-131 | A ROM2 ALM 11b entry is a shop stock record: for category i (0..3) mask 0x04+4i, minimum price 0x14+4i, maximum price 0x24+4i, draws 0x34+4i, quantity bound 0x44+4i; the Shop holds it in the Scenario.dll record layout. | High | ✔ promoted (branch candidate) | [EXP-2068](../experiments/EXP-2068-rom2-map-records/) |
| R2-ASSET-132 | In the ROM2 root .alm maps type-6 0x10 uses bits 1, 2, 4 and 8 and 0x0a takes 0..7; 11a flag 1 occurs although the type-4 pass stores nothing from it; the 11a word is 3661..3680 whenever flag 2 is set. | High | ✔ promoted (branch candidate) | [EXP-2068](../experiments/EXP-2068-rom2-map-records/) |

### R2-ASSET-129

The record dispatcher (RU `R2.0002`, EN `L2.02484`) reads the 20-byte
file header and stores an error code in map+0x8c: 2 when the magic
differs from the string at RU `L2.00002`, 3 when the record count is below
3, 4 when the version is above 0x640 (1600), and 1 when the archive does
not open. After the record loop it stores 5 when map+0xc is 0 and 6 when
map+0x14 is 0. Map+0xd8 and map+0xdc stay 1 unless a type-0 record sets
them.

The record loop runs while its index is below the record count (unsigned
`jae`). The header reader (RU `R2.0003`, EN `L2.02735`) reads 20 bytes,
except at version 0x3e8 (1000), where it reads none. The type is header
+0x0c: a type above 12 takes the default, 0..12 go through a 13-entry
table (RU `L2.02736`, EN `L2.02737`). The default skips header +0x08
bytes through archive slot +0x30.

Type 0 reads the counts the later cases use, each only from a version on:
the type-10 count at version 0x47e and above, the three type-11 counts at
0x4ce and above, the type-12 count at 0x514 and above (unsigned `jb`
tests). The dispatcher zeroes the five counts before the loop, so a map
below those versions loads no element of those types.

| Case | Loop test | Element read | Destination |
|---|---|---|---|
| 10 | `cmp i, count; jge` (signed) | 0x10 bytes | appended to map+0x310 |
| 11 | three loops, each `jge` | 0x0c, 0x54, 0x0c bytes | map+0x338 (11a), map+0x324 (11b), map+0x34c (11c) |
| 12 | `jge` | a 0x1c head, then count x 0x1c bytes | head at map+0x374, elements to map+0x360 |

Each element is read whole through archive slot +0x3c into a new
allocation. A count of 0 or a negative count reads no element. After each
case control returns to the loop head; no instruction of these cases or
of the loop compares the bytes read with header +0x08 or seeks to the
record end. So a count above the payload reads into the next record
header, and a count below it leaves payload bytes that the loop reads as
the next header. In both map populations of R2-ASSET-122 and R2-ASSET-130
every type-6, 10, 11 and 12 payload has exactly its counted size, so no
installed map exercises a mismatch.

**Confidence.** High for the error codes, the version gates, the signed
count loops and the absence of a size check: the dispatcher (1740
instructions) and the header reader (14) were read in both roots and are
equal after address masking, and both jump tables were read. Medium for
the misframing: that slot +0x3c reads sequentially from the current
position is inferred from its use; the archive class and its slots +0x30
and +0x3c were not read.

**Unknown.** What the archive does on a read past the end of the file, and
whether the loader then refuses the map. Reading the archive slots +0x30
and +0x3c settles both.

**Evidence.** [EXP-2068](../experiments/EXP-2068-rom2-map-records/)

### R2-ASSET-130

Population: every file under each install root whose name ends in `.alm`
in any case and any subdirectory, walked with the record framing of
R2-ASSET-002 and R2-ASSET-003 and checked as in R2-ASSET-122.

| Root | Maps | Type 4 | Type 6 | Type 10 | 11a | 11b | 11c | Type 12 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| EN | 13 | 1549 | 3249 | 674 | 376 | 452 | 28 | 16 |
| RU | 37 | 4360 | 11248 | 1696 | 477 | 555 | 31 | 58 |

- Every map is version 1600. Every type-6, 10, 11 and 12 payload size
  equals the size its type-0 counts give, and every type-4 walk ends at
  its payload end.
- Maps holding entries, EN and RU: type 10 in 13 and 37, 11a in 12 and
  29, 11b in 13 and 37, 11c in 6 and 8, type 12 in 2 and 6.
- Type-0 dword 0x84: EN 4 in 7 maps, 8 in 3, 16 in 3; RU 4 in 15, 6 in 1,
  8 in 9, 16 in 12. Dword 0x88: EN 1 in 6, 2 in 3, 3 in 2, 4 in 2; RU 1
  in 10, 2 in 8, 3 in 7, 4 in 12.
- Every 11a entry names a type-4 record of an Inn kind (0x43..0x45,
  0x63..0x65, 0x6f..0x71). Every 11b entry names a type-4 record of a Shop
  kind (0x22, 0x23, 0x5d..0x5f, 0x69..0x6b), except five per root: entries
  1, 4, 11, 16 and 19 of the 11b array of `SHALLOW.ALM`, all with id 265,
  which no type-4 record of that map carries. Every root 11c entry has gate 0 and position 0, on kinds
  0x46..0x4d.
- Matched by name without case, each of the 13 EN files has the SHA-256
  of its RU namesake; 24 RU files have no EN namesake.

Dword 0x84 is above 1 in every root map. R2-ENGINE-509 runs the type-10
pass and sets Session+0x74 in the map build it reads, so the gate opens if
the root maps are built by that routine; that they are was not read
(Medium).

**Confidence.** High: a committed derived census over both roots,
reproduced byte for byte from Git blobs. The counts are installed values,
not format limits. Medium for the sentence on the gate: it rests on the
root maps passing through the map build of R2-ENGINE-509, which carries
Medium in R2-ENGINE-108.

**Unknown.** Maps inside archives other than `scenario.res`: searched every
`.alm` file under each install root and `scenario.res`; listing the other
archives' directories settles it. The loader of a root map: no read shows
which routine builds it; reading the callers of the build RU `L2.00113`
settles it.

**Evidence.** [EXP-2068](../experiments/EXP-2068-rom2-map-records/)

### R2-ASSET-131

The type-4 pass (R2-ENGINE-512) stores entry dword 0x14+4i at Shop
+0x70+0x14i, 0x24+4i at +0x74+0x14i, 0x34+4i at +0x78+0x14i, 0x44+4i at
+0x7c+0x14i and 0x04+4i at +0x80+0x14i. Shop +0x70..+0xbf is therefore
four 0x14-byte entries ordered minimum price, maximum price, draws,
quantity bound and mask: the category entry of the Scenario.dll
ordinal-15 record (R2-ENGINE-263). The Shop fill reads Shop +0x70 in place
of that record when Session+0x74 is not 0 (R2-ENGINE-539). The field roles
are those R2-ENGINE-263, R2-ENGINE-265 and R2-ENGINE-471 read for the
record.

Values in the root maps of R2-ASSET-130 (EN 452 entries, RU 555), over
the four categories:

| Field | EN | RU |
|---|---|---|
| minimum price | 0..1000000 | 0..5000000 |
| maximum price | 0..10000000 | 0..20000000 |
| draws | 0..100 | 0..100 |
| quantity bound | 0..20 | 0..20 |

Draws 20 is the most frequent value in every category, and the masks take
24 to 43 distinct dwords per category in EN and 56 to 110 in RU.

**Confidence.** High for the field order: the stores of R2-ENGINE-512 and
the record layout of R2-ENGINE-263 were read in both roots, and
R2-ENGINE-539 reads the source switch in both roots. The values are a
reproduced census. How each field acts in a draw is stated by the cited
fill claims, not read here.

**Evidence.** [EXP-2068](../experiments/EXP-2068-rom2-map-records/)

### R2-ASSET-132

Values over the root maps of R2-ASSET-130. Type 6 (EN 3249 records, RU
11248):

| Offset | EN | RU |
|---:|---|---|
| 0x0a | 0..7 | 0..7 |
| 0x0c | 0 x2786, 16 x448, 20 x13, 4 x2 | 0 x10004, 16 x1229, 20 x13, 4 x2 |
| 0x10, records with bit 1 / 2 / 4 / 8 | 316 / 30 / 27 / 218 | 786 / 131 / 57 / 567 |
| 0x18 | 2..15; high word 0 in every record | same |
| 0x1c | 0 in 3191, 1..58 once each | 0 in 10937, 1..58 |
| 0x20 | 0..14; 10 in 3122, 0 in 85 | 0..14; 10 in 6103, 0 in 4900 |
| 0x24, 0x26 | -1..12000, equal distributions | -1..18432, equal distributions |
| 0x2c | 0..360 | 0..555 |

Dword 0x10 never has a bit above 8. The high word of dword 0x28 has the
same number of distinct values, range and twelve most frequent values as
dword 0x00 in both roots; per-record equality was not computed. In the
campaign maps 0x10 is 8 in 8 records per root and 0 in the rest.

Type 10 (EN 674 entries, RU 1696): e1 takes 22 values in 5..7500 in EN
and 72 in 5..30000 in RU, 120 in 494 and 1009 entries. e2 is in {0, 2, 3,
6, 10, 14} in EN, and also 4 and 7 in RU; bit 1 is set in 7 EN and 26 RU
entries, each with e3 at 0 or above. e3 is -1 in 666 EN and 1668 RU
entries, otherwise 0..3 (EN) or 0..9 (RU).

11a (EN 376, RU 477): the flags dword uses only bits 1, 2, 4, 0x10, 0x20
and 0x40; bit 1 is set in 143 EN and 158 RU entries. The type-4 pass
stores no field from bit 1 (R2-ENGINE-512). With flag 2 set (223 EN, 250
RU) the word at 0x08 is 3662..3679 in EN and 3661..3680 in RU; without it
the word is -1, 0 or a value in the same range and is not used.

11c: the 28 EN and 31 RU root entries all have gate 0, which the type-4
pass skips, so no root-map Pointer takes a count from 11c.

**Confidence.** High: a committed derived census reproduced byte for byte
from Git blobs. These are installed values, not format limits.

**Unknown.** What the 11a word values 3661..3680 identify (R2-ENGINE-536
passes the word to an offer constructor): searched the 11a words of the 13
EN and 37 RU root maps and the generator's pass of the word to block 10;
reading the block-10 constructor `L2.02738` settles it.

**Evidence.** [EXP-2068](../experiments/EXP-2068-rom2-map-records/)

## Route text as the EN original shows it

Runtime observation of the EN original `allods2.exe` in a 640x480 client
(EXP-2071 runs B6 and B6b). Steps are catalogue step ids; the
[receipt](../experiments/EXP-2071-rom2-route/evidence/receipt-en.tsv) gives
each step's screenshot name and PNG SHA-256. The tip panels and caption
boxes that carry this text are R2-ENGINE-563 and R2-ENGINE-564.

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-135 | EN route tips show 2 lines on pre-create, 8 on the detail page, 8 on the square, 10 in the inn and 4 in the shop, gold capitals, none clipped; the inn tip names a Hire button that its first visit never showed. | High / Medium | ● active | [EXP-2071](../experiments/EXP-2071-rom2-route/) |
| R2-ASSET-136 | EN route hover captions are capitals, a key hint in angle brackets (GUARD <G>, SWITCH TO STATS <TAB>); command buttons show mixed-case words (Accept, Talk, Undo); item texts give the name, then one line per stat. | High / Medium | ● active | [EXP-2071](../experiments/EXP-2071-rom2-route/) |

### R2-ASSET-135

Steps s2.01, s2.23, s2.90, s3.21, s3.60.

| Surface | Lines shown | Content |
|---|---|---|
| generator pre-create | 2 | choose class and sex by clicking a portrait |
| generator detail (male fighter) | 8 | Accept skips the screen; choosing one of five weapon classes gives that skill 20 |
| town-1 square | 8 | the town after a mission; the shop and the inn (hire mercenaries, receive missions); the game menu by Esc or the big statue |
| inn | 10 | patrons at the bottom; mercenaries on a greenish-gray back with a Hire button; Talk for a story; other patrons on light green give missions |
| shop | 4 | click a shelf, drag an item onto the table, Buy or Sell |

Every line is gold capitals; lines are justified to the panel width except
the last of a paragraph; no text runs past its panel. Key and button names
are in angle brackets (`<ESC>`, `<HIRE>`, `<TALK>`, `<ACCEPT>`, `<BUY>`,
`<SELL>`).

Relation to existing claims. It agrees with R2-ASSET-080 (the tips are
`town.txt` sections) and R2-ENGINE-314 (greedy wrap, justified non-final
lines, no shipped tip clips). Which section each panel shows follows
R2-ENGINE-289 and R2-ENGINE-308 (`#tips8`, `#tips5`, `#tips1`, `#tips2`,
`#tips3`), Medium: the text was not compared with the archive bytes. The
inn tip's Hire button and greenish-gray backs did not show at the first
visit (R2-SESSION-203).

**Confidence.** High for the line counts and look (each panel stayed until
Close). Medium for the section identity.

**Unknown.** The female and mage detail tip (`#tips6`), `#tips4`, `#tips7`,
`#tips9`, `#tips10`; the RU texts (the RU route was not run).

**Evidence.** [EXP-2071](../experiments/EXP-2071-rom2-route/)

### R2-ASSET-136

Steps s1.20, s2.04, s2.12, s2.25, s2.30, s2.35, s3.03, s3.05..s3.08, s3.24,
s3.29, s3.61, s3.64, s3.65, s3.72, s3.74, s5.05, s5.06.

| Screen | Captions in boxes | Words on buttons |
|---|---|---|
| main menu (pictures in the disc, not boxes) | NEW GAME, LOAD GAME, SERVER, HALL OF FAME, EXIT, CREDITS, CUT SCENES, MULTI PLAYER | none |
| generator | MALE FIGHTER, FEMALE FIGHTER, FEMALE MAGE, MALE MAGE, EASY, MEDIUM, HARD, RETURN TO MAIN MENU, CONTINUE, FREE POOL, signed numbers (-12, +10, +0) | Accept, Reset, Back |
| square | INN, SHOP, START MISSION, SCHOOL OF MAGIC AND MARTIAL ARTS, CLICK TO BRING UP THE MAIN MENU | none |
| inn | PREVIOUS HERO, NEXT HERO, SWITCH TO STATS <TAB>, SWITCH TO PICTURE <TAB> | Talk, Exit |
| shop | ARMOR, SHOPKEEPER, YOUR HERO'S INVENTORY | Undo, Buy, Sell, Exit |
| mission | GUARD <G>, DEFEND <D> | none |

Item texts: IRON PLATE CUIRASS / DEFENSE +12 / ABSORPTION +2; MITHRILL
CUIRASS / DEFENSE +26 / ABSORPTION +1 / MAGIC: / DEFENSE +23; LEATHER SMALL
SHIELD / DEFENSE +3. Weapon and attribute descriptions are capitals in
several lines with a blank line between a weapon list and its properties.

Relation to existing claims. It agrees with R2-ENGINE-384 (fixed help
lines for the generator's heroes and controls), R2-ENGINE-267 (item text
from attribute tags, a stat per line) and R2-ENGINE-437 (menu labels are
pictures).

**Confidence.** High for the case and bracket form across the 20 shots.
Medium for each single caption (one shot each).

**Unknown.** The RU wording; captions on controls not hovered (school
click, other mission icons, the minimap).

**Evidence.** [EXP-2071](../experiments/EXP-2071-rom2-route/)

## Game-menu, shop and mission text as the EN original shows it

Runtime observation of the EN original `allods2.exe` in a 640x480 client
(EXP-2072 runs C1, C2 and C3). Steps are catalogue step ids; the
[receipt](../experiments/EXP-2072-rom2-route-rest/evidence/receipt-en.tsv)
gives each step's screenshot name and PNG SHA-256. The boxes that carry
this text are R2-ENGINE-569 to R2-ENGINE-572 and R2-SESSION-215 to
R2-SESSION-217.

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-141 | EN game-menu boxes show mixed-case titles, labels and buttons and capital hover captions; the save box's Cancel caption is spelled CANCLE, the load box's CANCEL. | High / Medium | ● active | [EXP-2072](../experiments/EXP-2072-rom2-route-rest/) |
| R2-ASSET-142 | EN shop and mission captions add SCROLL UP, SCROLL DOWN, WEAPONS, SCROLLS, BOOKS & POTIONS, MAGIC ITEMS, OPEN SPELLBOOK <Q>,<B>, OPEN INVENTORY <~>,<I> and MAIN MENU <ESC>; an open panel's caption begins CLOSE. | High / Medium | ● active | [EXP-2072](../experiments/EXP-2072-rom2-route-rest/) |

### R2-ASSET-141

Steps ra.01, ra.05..ra.08, ra.10..ra.14, ra.16..ra.20.

| Box | Title, labels and buttons | Hover captions |
|---|---|---|
| game menu | Save Game, Load Game, Sound Options, Abort Game, Return to Game | none (the word is drawn gold) |
| save | Save Game, Saved Games, <<< Empty Slot >>>, Ok, Delete, Cancel | CONFIRM SELECTION; DELETE SELECTED / SAVED GAME; CANCLE |
| load | Load Saved Game, Load Game, Restart last mission, Abort mission and return to town, Ok, Delete, Cancel | CONFIRM SELECTION; DELETE SELECTED / SAVED GAME; CANCEL; RESTART LAST MISSION with a date and time |
| sound | Sound Options, Tracks, Plagat, Acknowledgments, Music volume, SFX volume, Speech volume, Play, Stop, Ok | RETURN TO GAME (Ok); PLAY; ACKNOWLEDGMENTS; MUSIC VOLUME |

CANCLE was read from a 4x zoomed crop of the caption box.

Relation to existing claims. It agrees with R2-ASSET-136 (capital
captions, mixed-case button words) and R2-ENGINE-441 (the menu's text
rows). R2-ENGINE-440 gives the save Cancel tip as dialogs row 33 and
R2-ENGINE-438 the load Cancel tip as row 27; CANCLE is then row 33's EN
text (Medium: the archive bytes were not compared). The row labels agree
with R2-ENGINE-271's `patch.txt` lines.

**Confidence.** High for the case pattern across the 15 shots. Medium for
each caption and label (one shot each), CANCLE included.

**Unknown.** The RU wording; the Stop, SFX and Speech captions (not
hovered). A read of EN and RU dialogs rows 26, 27, 32, 33 and 158 would
settle the row identity.

**Evidence.** [EXP-2072](../experiments/EXP-2072-rom2-route-rest/)

### R2-ASSET-142

Steps r2.02..r2.06, r2.25, r1.41, r1.44, r1.47, r1.48, r1.49.

| Screen | Control | Caption |
|---|---|---|
| shop | ornament above the grid | SCROLL UP |
| shop | gem below the grid | SCROLL DOWN |
| shop | room objects | ARMOR, WEAPONS, SCROLLS, BOOKS & POTIONS, MAGIC ITEMS |
| shop | top-left frame icon, open | CLOSE SPELLBOOK <D>,<B> (key hint an eye read) |
| mission | top-left frame icon | OPEN SPELLBOOK <Q>,<B> |
| mission | eye | SWITCH TO STATS <TAB> |
| mission | bottom-left frame icon, closed and open | OPEN INVENTORY <~>,<I>, CLOSE INVENTORY <~>,<I> |
| mission | bottom-right frame icon | MAIN MENU <ESC> |

Relation to existing claims. It agrees with R2-ASSET-136 (capitals, key
hints in angle brackets) and narrows its Unknown captions on other mission
icons.

**Confidence.** High for the capitals and the bracketed key hints across
the 11 shots. Medium for each caption (one shot each) and for the shop's
<D>.

**Unknown.** The RU wording; the shop key hint against the mission's <Q>.
A zoomed read of the shop caption would settle the letter.

**Evidence.** [EXP-2072](../experiments/EXP-2072-rom2-route-rest/)

## Inn card text and dialogue face sprites

Derived reads of the EN and RU `main.res` and `graphics.res` archives
(EXP-2073): CRLF line tables under R2-ASSET-010 and archive key censuses
(key, payload size, SHA-256; no payload byte). The routines that use these
lines and sprites are R2-ENGINE-577 to R2-ENGINE-583.

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-145 | ROM2 main.txt 83/84/190/191 are Enemy/Вражеский, Mercenary/Наемный, spellcaster/использует магию, Armor piercing/игнорирует броню; unitname.txt 1/15/24 Peasant, Crossbowman, Mage (RU Крестьянин, Арбалетчик, Маг). | High | ● active | [EXP-2073](../experiments/EXP-2073-rom2-town-persons/) |
| R2-ASSET-146 | The ROM2 graphics.res holds 116 two-level equipment/<class>/<file> keys and 185 infowindow/ keys, equal in name, size and SHA-256 in EN and RU; equipment/fmage/17.256 (11517 bytes) is byte-equal to equipment/mfighter/14.256. | High | ● active | [EXP-2073](../experiments/EXP-2073-rom2-town-persons/) |

### R2-ASSET-145

Zero-based CRLF line indexes of `main.res` `text/main.txt` and
`text/unitname.txt`, read in each root (RU decoded as cp1251):

| Table, line | EN | RU |
|---|---|---|
| main.txt 15..20 | Body, Agility, Mind, Spirit, Health, Mana | Сила, Ловкость, Разум, Дух, Жизнь, Мана |
| main.txt 83 | Enemy | Вражеский |
| main.txt 84 | Mercenary | Наемный |
| main.txt 190 | spellcaster | использует магию |
| main.txt 191 | Armor piercing | игнорирует броню |
| unitname.txt 1 | Peasant | Крестьянин |
| unitname.txt 15 | Crossbowman | Арбалетчик |
| unitname.txt 24 | Mage | Маг |

The card draws these in capital letters (R2-SESSION-203); the case
mapping is the font's, not the text's.

**Confidence.** High: both roots' tables read whole by the line reader.

**Evidence.** [EXP-2073](../experiments/EXP-2073-rom2-town-persons/)

### R2-ASSET-146

The census listed the keys of each root's `graphics.res` that fully match
`equipment/[^/]+/[^/]+` (two path parts below `equipment/`) and
`infowindow/.*`: 116 and 185 keys, equal in key, size and hash between EN
and RU. Deeper `equipment/` keys (three path parts, the item sprites under
`primary/` and `secondary/`) were not listed, and no EN/RU equality is
claimed for them. Five of the 116 equipment payload hashes occur more than
once; one pair is `equipment/fmage/17.256` and
`equipment/mfighter/14.256`, 11517 bytes, SHA-256
`bc1130885eb0b1a6616df37ce92c096ad2510e07b473a45c7a73eed5d2d8b04e`.

**Confidence.** High for the two listed key sets of the two installed
archives.

**Unknown.** The deeper equipment keys; the picture the sprite shows (no
sprite was decoded).

**Evidence.** [EXP-2073](../experiments/EXP-2073-rom2-town-persons/)

## Mission text at 1024x768 as the EN original shows it

Runtime observation of the EN original `allods2.exe` launched with
`-1024` (EXP-2074 run K1). Steps are catalogue step ids; the
[receipt](../experiments/EXP-2074-rom2-mission-1024/evidence/receipt-en.tsv)
gives each step's screenshot name and PNG SHA-256. Every text is an eye
read of the shot (crops at 2x or 3x).

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-148 | At 1024x768 the EN first mission's stats sheet shows the name line and 16 label lines from BODY to SPEED (17 lines in all); the mission game menu lists seven buttons from Save Game to Return to Game. | High / Medium | ● active | [EXP-2074](../experiments/EXP-2074-rom2-mission-1024/) |

### R2-ASSET-148

Steps k2.01, k2.04, k3.11.

The stats sheet below the portrait (R2-ENGINE-588), lines as drawn for
the new hero Aldor:

| Line | Left | Right |
|---|---|---|
| 1 | ALDOR | |
| 2 | BODY 40 | HEALTH |
| 3 | AGILITY 36 | 142/142 |
| 4 | MIND 25 | MANA |
| 5 | SPIRIT 17 | 0/0 |
| 6 | DMG 12-19 | ABSORB 4 |
| 7 | ATTACK 78 | DEFENSE 43 |
| 8 | SKILLS | RESISTANCE |
| 9 | BLADE 20 | FIRE 8 |
| 10 | AXE 0 | WATER 8 |
| 11 | BLUDGEON 0 | AIR 8 |
| 12 | PIKE 0 | EARTH 8 |
| 13 | SHOOTING 10 | ASTRAL 8 |
| 14 | WEIGHT 23.6 | |
| 15 | XP 7320 | |
| 16 | SIGHT 6.4 | |
| 17 | SPEED 19 | |

SKILLS and RESISTANCE are drawn in orange, the other labels in light
green.

The mission game menu box (MAIN MENU icon, R2-SESSION-221) lists Save
Game, Load Game, Game Options, Sound Options, Quest Objectives, End Quest
and Return to Game, top to bottom; the town-1 statue's box of
R2-ENGINE-569 lists five (no Game Options, Quest Objectives or End
Quest; Abort Game in their place).

Relation to existing claims. The body, agility, mind, spirit and blade
values and the health maximum equal the generator's opening values of
R2-SESSION-202.

**Confidence.** High for the presence of the sheet, the order of its
17 lines (the name line and 16 label lines from BODY to SPEED) and the
seven buttons (states that waited for input). Medium for
each number and word (eye reads of the small font).

**Unknown.** Which game state makes the mission box differ from the
town's (the screen size or the mission): a 640x480 mission run that
clicks MAIN MENU would settle it.

**Evidence.** [EXP-2074](../experiments/EXP-2074-rom2-mission-1024/)

## Mission-menu and item texts at 1024x768 as the EN original shows it

Runtime observation of the EN original `allods2.exe` launched with
`-1024` (EXP-2075 run L1). Texts are eye reads of the shots named by step;
the [receipt](../experiments/EXP-2075-rom2-screens/evidence/receipt-en.tsv)
gives each step's screenshot and PNG SHA-256.

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-150 | EN texts at 1024x768: the delete question, Game Options and Sound Options labels, the first quest objective, End Quest and Abort Game buttons, item captions and the first signpost's line. | Medium | ● active | [EXP-2075](../experiments/EXP-2075-rom2-screens/) |

### R2-ASSET-150

| Step | Where | Text as drawn |
|---|---|---|
| m1.02 | Save Game name field hover | ENTER A NAME FOR THE SAVED GAME |
| m1.08 | delete question | Are you sure you want to delete game 'exp75' ? · Yes · No |
| m1.15 | Game Options | Game Options · Game Speed · Day/Night changes · Smoothing · Shadows · Dynamic lighting · Object animations · Show Health · Display damage points · Tips · Alt. message colors · AutoCast for units: · Own · Allies · Neutral · Formation Mode: Off, Auto, On · Retreat Mode: Never, Low Health, Medium Health · AutoCast mode: Minimum, Average, Maximum · Ok · Cancel |
| m1.17 | Sound Options tracks in the mission | In Restless Waiting · Astral Threat · War of Druids · Tragedy of Battle |
| m1.19 | Quest Objectives | Quest Objectives and Hints: · Great Mage Skrakan · You must get to Rakhl-Uimoir tower and talk to Great Mage Skrakan. · Ok |
| m1.21 | End Quest box | Victory! · Exit to Main Menu · Exit to Windows · Return to Game |
| m1.32 | town Abort Game box | Exit to Main Menu · Exit to Windows · Return to Game |
| m2.07, m2.11 | buckler hover | BRONZE BUCKLER · DEFENSE +3 |
| m3.14, m2.26 | scroll hover | SCROLL OF HEAL · DAMAGE: 21-42 · RANGE: 7 |
| m2.13 | mission gold hover | MONEY |
| m2.01, m2.14 | empty or shop gold cell hover | YOUR HERO'S INVENTORY |
| m3.02, m3.04 | backpack end gems | SCROLL LEFT · SCROLL RIGHT |
| m2.28 | no hero selected | NO UNITS SELECTED (portrait) · NO HERO SELECTED (backpack row) |
| m2.06 | shop table hover | PUT ITEMS HERE TO BUY/SELL |
| m4.05 | first signpost box | To Great Mage Skrakan's Tower · Ok |

Menu labels draw one underlined letter each.

**Confidence.** Medium: eye reads of the small font, one shot each.

**Evidence.** [EXP-2075](../experiments/EXP-2075-rom2-screens/)

## First-mission signpost record and its event text

Derived measurements of the `10.alm` member of `scenario.res` and the
`text/mission10.txt` member of `main.res` in the owned EN and RU roots; no
member bytes.

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-152 | The first-mission signpost (10.alm Pointer id 42) selects event 19; `#event19` occurs once in each mission10.txt; the EN body's one text line equals R2-ASSET-150's drawn line, and the RU line is 23 bytes. | High / Medium | ✔ promoted (branch candidate) | [EXP-2077](../experiments/EXP-2077-rom2-signpost/EXP-2077.md) |

### R2-ASSET-152

`10.alm` is equal in both roots: 74298 bytes, SHA-256
`23d22542745ab0c101a60b2da5ea80b858eeb4e1840515bd8398ceb2ad9efeca`. Its
only Pointer (R2-ASSET-123) is the type-4 record id 42, kind 0x4a, x 3712,
y 10368: cell (14, 40), fraction (128, 128). No other type-4 record lies
within 3 cells. Its three 11c entries are equal: gate 1, node position 25;
24 of the 25 nodes before position 25 have signed dword +0x40 (operation)
below 0x10002 (node 0 has 0x10002), so Pointer +0x6c = 24
(R2-ENGINE-517). Node 25 (ID 26) has operation 2, type 1 and parameter 0 =
19: event 19 (R2-ENGINE-047), formatted as `event19` (R2-ENGINE-050).

| Root | mission10.txt | `#event19` at | Body | Body SHA-256 | Text line |
|---|---:|---:|---|---|---|
| EN | 9775 B | 9723 | 9733..9774, 41 B | `996f769ffa17beb25628a0a4073d3952da0bd2c23442d9a00f5bbf26df6d254f` | 29 B, `ec1b8c2b84796e4a8c595e730ba2de33716a482268e63db51d5094d9f4d8f9cc` |
| RU | 9390 B | 9342 | 9352..9387, 35 B | `003498ab2709b44c52a5aa93621934a78ad7703e18f30cbd8f332695b114bfff` | 23 B, `45572f371d4f992c17e4f02aa1a3303c76b41e78d983e5a1bc630e44edb8fdaf` |

Each body is one header line that begins with `<` and one text line, both
ended by CR LF. The EN text line is the 29 bytes of "To Great Mage
Skrakan's Tower", the line R2-ASSET-150 saw drawn in the box.

**Confidence.** High for the record, the node, the key counts, offsets,
lengths and hashes, and the EN equality. Medium that node 25 is the node the
Pointer runs (the node-to-slot step of R2-ENGINE-520) and that operation 2
parameter 0 is the event ID (R2-ENGINE-047).

**Unknown.** The RU line as the RU original draws it: no RU run was made.

## Spell row 29 names

`main.res` of the EN and RU installs, read through the repository archive
reader; RU bytes decoded as code page 1251.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-154 | ROM2 spell.txt line 28 names spell row 29: EN "Slow", RU "Замедление"; none of the 24 spells.txt hover lines in either root names it. | High | ● active | [EXP-2078](../experiments/EXP-2078-rom2-mission-load/) |

### R2-ASSET-154

`text/spell.txt` of `main.res` splits on CRLF into 30 lines in each of the
EN and RU roots (EN SHA-256
`33ab26bfa0d5ee632dac051cafca9d277650167fe423ca17bf2f429eb7102990`, RU
`8eb0fc59835ef1fbc575854506d6be98d250d3c3623d18019ba6b2b1bdb3d999`); line
28 is row 29 under the row r to line r-1 rule of R2-ENGINE-444.
`text/spells.txt` splits into 25 lines (EN
`8e2e1a3e9da40f74d44af7f84c6e2241ae966986098fc920c7e32d1b2a4b19c1`, RU
`9ba346dcd71347089ddb1562d284b42df60b631d724cb09ad033c13c65b31a8a`); for
each line the text before " (" was compared with every `spell.txt` line,
and none equals line 28 in either root.

**Confidence.** High: whole-file byte comparisons in both roots.

## Identity 21..30 rows and map literals

`world.res` of both roots, read through EXP-2047's `data.bin` walker, and
`scenario.res` of both roots, read through `tools/r2script`; numbers, row
names and SHA-256 only.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-160 | Both ROM2 world.res (one SHA-256) give identity 22 only to Humans Fighter_Hero (raw 10220) and hold no 10221..10229 row; the 46-map scripts of each root hold no 10001..10010 or 10220..10229 parameter and one 22 (map 96). | High | ● active (branch candidate) | [EXP-2080](../experiments/EXP-2080-rom2-iglez-join/) |

### R2-ASSET-160

Both `world.res` files have SHA-256
`863b15ce6a304263d9ed97ca14426db42e07825b91165b39d5adf4054d9dd8a9`. The
measure of a row:

- Humans: parameter 24, mapped to (w/10)%1000 when w is above 10000 (the
  template load's rule, R2-ENGINE-525);
- Units: the low word of parameter 55.

| Table | Rows | With the parameter | Identity in 21..30 |
|---|---:|---:|---:|
| Humans | 289 | 284 | 9 |
| Units | 241 | 180 | 0 |

| Humans row | Raw parameter 24 | Identity |
|---|---:|---:|
| `ManHorse_LanceShield` | 21 | 21 |
| `ManMage_Unarmed` | 23 | 23 |
| `ManMage_Staff` | 24 | 24 |
| `Start_MF`, `Start_FF`, `Start_MM`, `Start_FM` | 10210..10213 | 21 |
| `Fighter_Hero` | 10220 | 22 |
| `Mage_Hero` | 10230 | 23 |

No row holds raw 22 or a raw value in 10221..10229.

The r2script tables of the 46 maps per root were searched for parameters
holding 10001..10010 (R2-ENGINE-107's controlled literals for identities
21..30), 10220..10229 or 22. One parameter per root matches: map 96, check
list, operation 26, parameter slot 0, value 22, type 4. The EN and RU
results are equal.

**Confidence.** High for the counts, values and names in both roots. The
search covers parameter values only. A reference built at run time, or
through a unit record's template, is outside it.

## First mission map planes and fog fields

`scenario.res:10.alm` of the EN and RU installs, read through the
repository resource and ALM readers.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-162 | ROM2 scenario.res:10.alm is byte-identical in EN and RU: version 1300, 80 x 80; byte plane 2 spans 0..127, plane 3 0..81; no type 1 word has 0x8000 or 0x4000; no type 5 entry has 0x8. | High | ✔ promoted (branch candidate) | [EXP-2081](../experiments/EXP-2081-rom2-fog-stamp/) |

### R2-ASSET-162

`scenario.res` differs between the roots (EN 5548580 bytes, SHA-256
`388c955f49c02ecfed765340ff3a5b3e8f37210e2c85c8d877bc04b949d60d7f`; RU
5549452 bytes,
`c19950690570d94b40c68acffd210cdab94c2a12b5d6ad5718730392c6b6e08b`); its
`10.alm` entry is 74298 bytes in both, SHA-256
`23d22542745ab0c101a60b2da5ea80b858eeb4e1840515bd8398ceb2ad9efeca`.

| Field | Value |
|---|---|
| version, W, H | 1300, 80, 80 |
| record order by typeId | 0, 1, 2, 3, 5, 11, 4, 9, 8, 6, 7, 10, 12 |
| type 1 words (6400) with 0x8000, 0x4000, 0xc000 | 0, 0, 0 |
| type 2 bytes, min and max | 0, 127 |
| type 3 bytes, min and max | 0, 81 |
| type 2 SHA-256 | `65788c11786050a685e789a7a0652455d460af5af1f2e0fb6214da4f9a3b078a` |
| type 3 SHA-256 | `e0238972bf618c0c64fcae451113489f8ebeb330081fdc0a9fcb9c39937e6aca` |
| type 2 cells within 20 of cell 11/41 unequal to that cell | 1206 |
| type 6 units | 29 |
| type 6 units within 20 cells of 11/41 | 6, all with owner byte 3, distances 16..20 |

Type 5 holds 8 records. u32 +0 is 1..8; +4/+8 are 0/0 in the first record
and 1/5000 in the others. Of the 16 u16 entries at +0x2c..+0x4a, entry k of
record k (k = 0..7) is 0x12 in records 0..6 and 0x10 in record 7; every
other entry is 0, 1 or 2. No entry has the value bit 0x8.

**Confidence.** High: whole-entry hashes and counts in both roots. The
names W, H, typeId and owner byte follow the ALM page's record layout.

## Units special-spell slots

The client `data/data.bin` inside `world.res` (RU read in full; both roots'
`world.res` have one SHA-256, R2-ASSET-160), read through the grammar of
R2-ASSET-029. Parameter n carries title n+1 (R2-ASSET-107): p48..p53 are
Spell 1, Probability 1 .. Spell 3, Probability 3, p54 Spell Power, p56
knownSpells.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-164 | In the one ROM2 data.bin, 38 Units rows, all .4 or .5, hold a Spell 1..3 above 0: 54 entries of 17 ids, Probability 10 (Spider 20) in slot 1, 1 and 3 in slots 2 and 3; 64 rows hold a knownSpells mask, none both. | High | ● active (branch candidate) | [EXP-2082](../experiments/EXP-2082-rom2-creature-aim/) |

### R2-ASSET-164

Of 241 Units rows, 180 have 62 values and 61 have none. A slot value of -1
counts as no spell (the loader tests above 0, R2-ENGINE-652): 142 rows have
Spell 1 = -1, 172 have Spell 2 = -1 and 172 have Spell 3 = -1. Each of these
base names has a slot spell at levels .4 and .5, and at no other level:

| Units base | Slot 1 (Probability) | Slot 2 | Slot 3 |
|---|---|---|---|
| Goblin_Pike | 29 Slow (10) | | |
| Goblin_Sling | 27 Shield (10) | | |
| Orc_Sword | 23 Teleport (10) | | |
| Orc_Bow | 20 Bless (10) | | |
| Ogre | 21 Haste (10) | | |
| Troll | 18 Stone Curse (10) | | |
| Bat_Sonic | 14 Darkness (10) | | |
| Ghost | 26 Drain Life (10) | | |
| Bee | 12 Invisibility (10) | | |
| Squirrel | 1 Fire Arrow (10) | | |
| Foot_Animated | 28 Curse (10) | | |
| Turtle | 9 Acid Stream (10) | | |
| F_Zombie | 6 Poison Cloud (10) | 18 (1) | 27 (3) |
| A_Zombie | 7 Blizzard (10) | 18 (1) | 27 (3) |
| F_Skeleton | 11 Prismatic Spray (10) | 26 (1) | 12 (3) |
| A_Skeleton | 14 Darkness (10) | 26 (1) | 12 (3) |
| Wolf | 16 Diamond Dust (10) | | |
| Spider | 6 Poison Cloud (20) | | |
| Dino | 5 Ice Missile (10) | | |

The 17 ids are 1, 5, 6, 7, 9, 11, 12, 14, 16, 18, 20, 21, 23, 26, 27, 28 and
29; no row holds 22 or an id above 29 in a slot. No row has Spell 1 at or
below 0 with Spell 2 or 3 above 0. 64 rows hold a knownSpells value other
than 0 and -1, and none of them holds a slot spell above 0. The Spells table
has 34 rows; rows 30..34 have no name and no value.

**Confidence.** High: committed derived rows that reproduce byte for byte;
EN equal by the `world.res` and `data.bin` hashes.

**Evidence.** [EXP-2082](../experiments/EXP-2082-rom2-creature-aim/RESULTS.md)

## Mission 40 persons and their sprite keys

The type-6 records of `scenario.res` `40.alm` through the field split of
R2-ASSET-119 (`w08` typeID, `d0c` flags, `d14` definition ID, `d18`
player); the Humans rows of `world.res` `data/data.bin` through the grammar
of R2-ASSET-029 (parameter 16 typeID, 17 face, 18 gender, 24 serverID, the
trailing item strings); the inline REG store `units/units.reg` and the keys
of `graphics.res`; `main.res` `text/globalmap.txt` and
`text/heropicture.txt`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-172 | ROM2 `#title40` (EN The Astrologer, RU Звездочет.) matches 40.alm by number; of 108 type-6 records 35 carry flag 0x10 and Humans serverIDs 2104..2112 or 10230 (Mage_Hero, NPC 23), 73 flag 0 (units.reg types, no Units row read); EN = RU. | High / Medium | ● active (branch candidate) | [EXP-2086](../experiments/EXP-2086-rom2-villagers/) |
| R2-ASSET-173 | ROM2 units.reg gives person types 1, 8, 9, 10, 14, 23, 24 no Z key and the sprite files unarmed, axeman_sh, axeman_2hd, clubman, archer, mage, mage_st; heropicture.txt line 0 is unarmed; the selected keys exist; EN = RU. | High | ● active (branch candidate) | [EXP-2086](../experiments/EXP-2086-rom2-villagers/) |

### R2-ASSET-172

`#title40` is the only key with the EN label The Astrologer; its RU label is
Звездочет. `#title40` and 40.alm correspond by number: R2-ENGINE-434 formats
`#title%d` from a global-map record ID, and nothing read here shows that
record ID 40 loads the member `40.alm`. The shared NPC words support the
join: the map's placed NPC words 23, 694 and 2112 are the speakers of
`mission40.txt` (R2-ENGINE-679). The map census and the template census are
equal in both roots row for row.

| Humans row | serverID | Records | Player | typeID | face | gender | Item strings |
|---|---:|---|---:|---:|---:|---:|---|
| 2U_PeasantMale | 2108 | 25..29, 35 | 3 | 1 | 46 | 0 | none |
| 2U_PeasantFemale | 2109 | 30..34 | 3 | 1 | 13 | 1 | none |
| 2U_PeasantLeader | 2112 | 9 | 3 | 1 | 44 | 0 | Uncommon Silver Ring |
| 2F_PeasantGuard | 2110 | 10..17 | 3 | 10 | 45 | 0 | Wood Club |
| 2A_PeasantGuard | 2111 | 18..24 | 3 | 14 | 16 | 1 | Wood Short Bow |
| Mage_Hero | 10230 | 36 | 3 | 23 | 20 | 1 | Uncommon None Shoes, None Dress |
| 2F_Brigand2 | 2104 | 0, 39, 40 | 5 | 8 | 51 | 0 | Iron Axe, Leather Boots |
| 2A_Brigand2 | 2105 | 3, 4 | 5 | 14 | 17 | 1 | Wood Long Bow, Leather Boots |
| 2M_Brigand2 | 2106 | 2 | 5 | 24 | 15 | 1 | Uncommon Wood Shaman Staff |
| 2F_BrigandLeader2 | 2107 | 5 | 6 | 9 | 52 | 0 | Iron Two Handed Axe, Leather Boots |

Each record's `w08` equals its row's typeID. The NPC word of serverID 10230
is (10230 / 10) % 1000 = 23 (R2-ENGINE-072); the other rows keep their
serverID. Mage_Hero has ManaMax 70. Row names are internal and equal in both
roots. The other 73 records have `d0c` 0: types 66, 68, 69, 70, 73, 74, 75, 76,
99 and 103, players 2, 4 and 7; records 6..8 are type 99 with serverIDs 694
and 695. No Units row of `data.bin` was read for these 73 records; their
types are named from `units.reg` `DescText` (type 99 is a Necromancer), and
the flag-0 arm is the flag test of R2-ENGINE-038, read on RU only.

**Confidence.** High for the record and row values: committed derived rows
reproduced byte for byte. Medium for the record-to-row join: it is the
serverID match of R2-ENGINE-038, read on RU only. Medium for the
title-to-map join: a matching number, corroborated by the shared NPC words.

**Unknown.** The name the game shows for each person; `text/unitname.txt`
was not read.

**Evidence.** [EXP-2086](../experiments/EXP-2086-rom2-villagers/RESULTS.md), question 1.

### R2-ASSET-173

| units.reg | ID | Parent | File | `Files/File<n>` | InfoPicture |
|---|---:|---:|---:|---|---|
| Unit21 Unarmed Fighter | 1 | 3 | 21 | `humans\unarmed\sprites` | Bee |
| Unit4 Human AxeMan with shield | 8 | 3 | 4 | `humans\axeman_sh\axeman_sh` | AxeManSh |
| Unit5 Human AxeMan with 2-handed axe | 9 | 3 | 5 | `humans\axeman_2hd\axeman_2hd` | 2HAxeMan |
| Unit6 Human ClubMan | 10 | 3 | 6 | `humans\clubman\clubman` | ClubMan |
| Unit8 Human Archer | 14 | 3 | 8 | `humans\archer\archer` | Archer |
| Unit20 Unarmed Mage | 23 | 3 | 20 | `humans\mage\sprites` | Mage |
| Unit16 Human Mage | 24 | 3 | 16 | `humans\mage_st\mage_st` | Mage |

Unit0 (ID 3, Human Swordsman) has no Parent. Key `Z` occurs in three units
only: 70, 71 and 105, value 96. `Global/FileCount` is 60 and
`Global/UnitCount` 61.

`text/heropicture.txt` lines 0..25: unarmed, swordsman (1..4), swordsman2h,
clubman (6..9), axeman, axeman2h, mage_st (12, 13), pikeman (14..16),
axeman, axeman2h, archer (19, 20), xbowman, empty, Sonic Beam, Flame
Thrower, swordsman.

`graphics.res` holds `units/humans/<dir>/<name>.256` and `<name>b.256` for
the six human files above, `units/heroes/mage/` and
`units/heroes/mage_st/` `sprites.256` and `spritesb.256`, and the paperdoll
bodies `equipment/mfighter/44, 45, 46, 51, 52`, `equipment/ffighter/13, 16,
17` and `equipment/fmage/15, 20` (`.256`).

**Confidence.** High: committed derived rows reproduced byte for byte, equal
in both roots.

**Evidence.** [EXP-2086](../experiments/EXP-2086-rom2-villagers/RESULTS.md), questions 3 and 4.

## Support spell names and «Удача в бою»

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-176 | No node of ROM2 EN or RU main.res, patch.res or world.res and no htm, html or txt file of either install holds «удача в бою»; by meaning it names Bless, row 20, slot 5, RU «Благословение». | High / Medium | ● active (branch candidate) | [EXP-2088](../experiments/EXP-2088-rom2-held-marks/) |
| R2-ASSET-177 | ROM2 spell rows 4, 8, 13, 19, 20 and 27 are Protection from Fire, Water, Air, Earth, Bless and Shield in EN and RU spell.txt, at spells.txt slots 3, 15, 8, 20, 5 and 19 by equal names. | High | ● active (branch candidate) | [EXP-2088](../experiments/EXP-2088-rom2-held-marks/) |

### R2-ASSET-176

Needles, case-folded over cp1251 text: «удача в бою», «удача», «удач», «в
бою», «благослов», «магический щит», «защита от». Population per root: every
node of `main.res`, `patch.res` and `world.res`, and every `*.htm`, `*.html`
and `*.txt` file under the install (6 archives and 41 files in all). The
phrase has 0 hits. The «удача» and «удач» hits are mission and town
dialogue, the attack-accuracy help line of `text/main.txt`, two RU Manual
pages and the RU `readme.txt`; none names a spell.

Of the six spells, only Bless's hover line speaks of success in a blow: EN "Bless
(Homing)#Increases warrior's chances to do maximum damage with a hand-held
weapon", RU «Благословение (действует на цель)#Повышает воину шанс нанести
максимальный урон». The RU `readme.txt` describes «благословление» the same
way.

**Confidence.** High for the search over the named population. Medium for
the identity, which rests on meaning, not on a text that joins the phrase
to a row; the other 18 hover lines were not compared in RU. Editions, patches and fan translations outside the two installs
were not searched.

**Evidence.** [EXP-2088](../experiments/EXP-2088-rom2-held-marks/RESULTS.md), question 1.

### R2-ASSET-177

| Row | EN | RU | Slot | Kind, EN / RU |
|---:|---|---|---:|---|
| 4 | Protection from Fire | Защита от магии огня | 3 | Homing / действует на цель |
| 8 | Protection from Water | Защита от магии воды | 15 | Homing / действует на цель |
| 13 | Protection from Air | Защита от магии воздуха | 8 | Homing / действует на цель |
| 19 | Protection from Earth | Защита от магии земли | 20 | Homing / действует на цель |
| 20 | Bless | Благословение | 5 | Homing / действует на цель |
| 27 | Shield | Магический щит | 19 | Caster Only / действует на себя |

Row r is line r-1 of `text/spell.txt`; slot n is line n of
`text/spells.txt`, whose text before the parenthesis equals the row's name
in the same root. The six rows' identifiers are 16, 24, 34, 46, 48 and 62
(R2-ENGINE-445).

**Confidence.** High for the committed lines in both roots and the name
equality.

**Evidence.** [EXP-2088](../experiments/EXP-2088-rom2-held-marks/RESULTS.md), question 1.

## Spells records of data.bin and names of the other held-mark rows

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-178 | ROM2 world.res data/data.bin holds 29 named Spells records then 5 empty entries, equal in EN and RU; 14 carry an Effects string, and record k matches spell.txt row k by name except rows 2, 16 and 22. | High / Medium | ● active (branch candidate) | [EXP-2089](../experiments/EXP-2089-rom2-other-marks/) |
| R2-ASSET-179 | ROM2 spell rows 6, 12, 14, 15, 18, 21, 22, 24, 26, 28 and 29 are Poison Cloud, Invisibility, Darkness, Light, Stone Curse, Haste, Animate Dead, Heal, Drain Life, Curse and Slow in EN spell.txt, with the RU names of the card. | High | ● active (branch candidate) | [EXP-2089](../experiments/EXP-2089-rom2-other-marks/) |

### R2-ASSET-178

Record shape: CString name, u16 parameter count 22, 22 u32 values, CString
Effects (formats/rom2-databin). Both roots' `world.res` have SHA-256
`863b15ce6a304263d9ed97ca14426db42e07825b91165b39d5adf4054d9dd8a9` and its
`data/data.bin` node
`aa9952354f1cbc4329528dff49ed4a6d9c875877b3ce22fe2a1c48b12d3b84d6`; the 29
named records start at offsets `0x23cb3`..`0x2494a`. The last, Slow, ends at
`0x249bf`; the 20 bytes to the end of the node are zeros, which read as 5
empty entries of 4 bytes (empty name, parameter count 0, empty Effects), so
34 entries follow the stored count 35 (formats/rom2-databin).

| Record | Name | Effects |
|---:|---|---|
| 4, 8, 13, 19 | Protection from Fire, Water, Air, Earth | `protectionFire=+30: duration 30` and the like |
| 6 | Poison Cloud | `health=-2:continuous 8` |
| 12 | Invisibility | `Invisibility = ON` |
| 14 | Darkness | `scanRange=-2:duration 1` |
| 15 | Light | `scanRange=+2:duration 1` |
| 18 | Stone Curse | `"absorbtion=+5:duration 8,defence=-20:duration 8"` (the string is stored inside double quotes) |
| 20 | Bless | `toHit=+1:duration 30` |
| 21 | Haste | `speed=+2: duration 30` |
| 27 | Shield | `absorbtion=+5:duration 30` |
| 28 | Curse | `toHit=-1:duration 30` |
| 29 | Slow | `speed=-2: duration 30` |

The other 15 named records have an empty Effects string. Records 2, 16 and 22 are
named Fire Ball, Diamond Dust and Control Spirit where EN spell.txt rows 2,
16 and 22 read Fireball, Stone Missile and Animate Dead.

**Confidence.** High for the records and strings in both roots. Medium for
the join of records 2, 16 and 22 to those rows, which rests on order only.

**Evidence.** [EXP-2089](../experiments/EXP-2089-rom2-other-marks/RESULTS.md), question 1.

### R2-ASSET-179

| Row | EN | RU |
|---:|---|---|
| 6 | Poison Cloud | Ядовитый туман |
| 12 | Invisibility | Невидимость |
| 14 | Darkness | Тьма |
| 15 | Light | Свечение |
| 18 | Stone Curse | Каменное проклятие |
| 21 | Haste | Ускорение |
| 22 | Animate Dead | Перерождение |
| 24 | Heal | Исцеление |
| 26 | Drain Life | Вампиризм |
| 28 | Curse | Проклятие |
| 29 | Slow | Замедление |

Row r is line r-1, counted from 0, of `text/spell.txt` in each root's
`main.res` (EN node
`33ab26bfa0d5ee632dac051cafca9d277650167fe423ca17bf2f429eb7102990`, RU node
`8eb0fc59835ef1fbc575854506d6be98d250d3c3623d18019ba6b2b1bdb3d999`), as in
R2-ASSET-177.

**Confidence.** High for the committed lines in both roots.

**Evidence.** [EXP-2089](../experiments/EXP-2089-rom2-other-marks/RESULTS.md), question 1.

## Unit types' attack timing and selection boxes

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-180 | ROM2 units.reg is one file in EN and RU; mage types 23 and 24 have AttackPhases 6, AttackAnimTime 2 x 6, AttackDelay 8, ShootDelay 8, Center 64/78 and Selection 48, 48, 80, 90; 61 types are tabulated. | High / Medium | ● active (branch candidate) | [EXP-2090](../experiments/EXP-2090-rom2-teleport/) |

### R2-ASSET-180

`units/units.reg` in EN and RU `graphics.res` has the same SHA256,
`335809b0672fdb796e8ae3399e176c7920167be8c7271b8abd7bbfaf76a0ac8a`
(`data/units-reg-inputs.tsv`). The reader takes the REG inline store of
R2-ASSET-006 and R2-ENGINE-614 (count at 0x10, 32-byte records from 0x18, type (flags & 0xe)
>> 1, type 3 an int array) and writes every section's ID, Parent, DescText,
File, Index, AttackPhases, Width, Height, CenterX, CenterY, SelectionX1, Y1,
X2, Y2, Palette, ShootDelay, AttackDelay, AttackAnimTime, AttackAnimFrame,
Homing and Projectile (`data/units-reg.tsv`). `data/teleport-unit-values.tsv`
resolves the scalar keys through Parent for the 61 types with an ID:

| Type | AttackPhases | AttackAnimTime sum | AttackDelay | ShootDelay | Center | Selection |
|---|---|---|---|---|---|---|
| 1 Unarmed Fighter | 7 | 14 | 0 | absent | 64, 78 | 48, 48, 80, 90 |
| 23 Unarmed Mage | 6 | 12 | 8 | 8 | 64, 78 | 48, 48, 80, 90 |
| 24 Human Mage | 6 | 12 | 8 | 8 | 64, 78 | 48, 48, 80, 90 |
| 71 Dragon | 7 | 14 | 8 | 8 | 80, 80 | 32, 32, 128, 128 |
| 107 WarSheep | 0 | none | 0 | absent | 64, 64 | 48, 48, 80, 90 |

The client loader EN `L2.03721` (RU `L2.03722`) reads each scalar through
`R2.0113` (unread) with a default: the parent's loaded field (type map
`L2.02969`) when Parent is not -1, else -1 for AttackPhases, Width, Height,
CenterX, CenterY and the four Selection keys, and 0 for ShootDelay and
AttackDelay (`L2.03824`, `L2.03825`). The table writes -1 where neither the
type nor an ancestor carries a key; for ShootDelay ("absent" above) the
loader's value is 0. It stores AttackPhases at type +0x1c, CenterX/Y at
+0x34/+0x38, Selection X1, Y1, X2, Y2 at +0x84..+0x90 (through `L2.01491`
and `L2.03826`), ShootDelay at +0x11c and AttackDelay at +0x120; +0x68 is the
length of the attack frame list built from AttackAnimTime and
AttackAnimFrame (R2-ENGINE-704). When a type has a Parent and its own
AttackAnimTime list is empty, the loader reads AttackAnimTime and
AttackAnimFrame again from the Parent's `Unit%d` section (`L2.03827`..
`L2.03828`; the section name is built from the Parent's type object at
`L2.03829`..`L2.03830`). Of the 61 types, 3 have no list of their own: types
4 and 2, whose Parent is type 3 with a list (sum 11), and type 107, which has
no Parent. The sum cells of those three are blank in
`data/teleport-unit-values.tsv`; the blank is the table's, which sums the
type's own list. AttackDelay is 0 for types 1, 2, 69, 74, 76, 103, 104, 106
and 107.

**Confidence.** High for the values read from the file and its equality in
both roots, and for the key-to-field stores in the EN loader. Medium for the
default rule: `R2.0113` returning its default for a missing key is inferred
from the call shape. High for the control flow of the Parent fallback, equal in both roots
(R2-ENGINE-708). Medium that types 2 and 4 therefore play type 3's 11
updates: the read of a key into a list is inferred from the call shape of
`L2.03128` (unread), and the Parent's section is its own section, so a
grandparent's list is not consulted by this read.

**Unknown.** The frame list of a type with neither a list of its own nor a
Parent, such as type 107 (AttackPhases 0, so it shows no cast).

**Format surface.** units.reg keys; `formats/rom2-reg/format.md`.

**Evidence.** [EXP-2090](../experiments/EXP-2090-rom2-teleport/RESULTS.md), question 2.

## Shop cell plates, backs and font2 advances

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-182 | ROM2 graphics.res costs1..7 and costm1..7 are 60x10 24-bit with first drawn columns 45, 39, 33, 24, 18, 12, 3; backinvg/b/s are 80x80; font2.dat advances digits 4 ('1' 2), comma 1, space entry 0 (drawn width Medium); EN and RU are equal. | High | ✔ promoted (branch candidate) | [EXP-2091](../experiments/EXP-2091-rom2-shop/) |

### R2-ASSET-182

`graphics.res` members, each SHA-256 equal in EN and RU:

| Member | Size | Bits | First and last drawn column |
|---|---|---|---|
| `interface/costs1..7.bmp` | 60x10 | 24 | 45, 39, 33, 24, 18, 12, 3; last 59 |
| `interface/costm1..7.bmp` | 60x10 | 24 | the same columns |
| `interface/backinvg.bmp`, `backinvb.bmp`, `backinvs.bmp` | 80x80 | 24 | no pixel converts to 0 |

A drawn column holds a pixel that is not 0 after the 24-to-16-bit
conversion, by the 555 and by the 565 rule alike (R2-ENGINE-714). The
columns are those of the plate bitmaps; they say nothing about where a glyph's
ink lies.

`font2/font2.dat` holds 224 dword advances indexed by the mapped code
minus 0x20; map `R2.0060` returns codes below 0x80 unchanged, so the index is code
minus 0x20: '0' and '2'..'9' 4, '1' 2, ',' 1, space (entry 0) 0, '@'
(entry 32) 6. With spacing 2 a digit takes 6 px, '1' 4 px and a comma 3 px.
Space is the table entry; the width routine adds half of an image width for
entry 0 through the unread vtable +0x24 (`L2.03899`), so the drawn width of
a space is Medium (see the amendment of R2-ASSET-081).

**Confidence.** High for the headers, the columns, the advances and the
equality. Medium for the drawn width of a space.

**Evidence.** [EXP-2091](../experiments/EXP-2091-rom2-shop/RESULTS.md), question 5.

## Unit voice files

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ASSET-185 | ROM2 sfx.res holds unit voices as characters/<voice>/<gesture>/<n>.wav: EN 88 keys (gestures of six named voices, easy/hard/dead of twelve), RU 198 (gestures of twelve, easy/hard/dead of all eighteen voices). | High | ✔ promoted (branch candidate) | [EXP-2092](../experiments/EXP-2092-rom2-unit-voices/) |

### R2-ASSET-185

Census of the `characters/` keys of the installed `sfx.res` (EN 18,219,970
bytes, RU 25,041,450) against the names the voice init of R2-ENGINE-718
builds. EN / RU counts:

| Voice | select | attack | move | swarm | retreat, defend | pickup | easy, hard, dead |
|---|---|---|---|---|---|---|---|
| `m_`/`f_` + `fighter`, `mage`, `peasant` | 0/0 | 0/0 | 0/0 | 0/0 | 0/0 | 0/0 | 0/1 each |
| aldor | 3/4 | 2/2 | 2/3 | 2/2 | 1/1 | 1/1 | 1/1 |
| sagita | 2/3 | 1/1 | 1/3 | 1/3 | 1/1 | 1/1 | 1/1 |
| hildarius | 2/3 | 1/1 | 1/3 | 1/3 | 1/1 | 1/1 | 1/1 |
| galinel | 2/4 | 1/1 | 1/2 | 1/3 | 1/1 | 1/1 | 1/1 |
| igles | 1/3 | 1/1 | 1/3 | 1/3 | 1/1 | 1/1 | 1/1 |
| dina | 1/4 | 1/1 | 1/3 | 1/2 | 1/1 | 1/1 | 1/1 |
| xbow | 0/2 | 0/1 | 0/3 | 0/2 | 0/1 | 0/0 | 1/1 |
| knight | 0/3 | 0/1 | 0/3 | 0/2 | 0/1 | 0/0 | 1/1 |
| druid | 0/2 | 0/1 | 0/2 | 0/3 | 0/1 | 0/0 | 1/1 |
| troll | 0/3 | 0/1 | 0/1 | 0/3 | 0/1 | 0/0 | 1/1 |
| orc | 0/3 | 0/1 | 0/2 | 0/3 | 0/1 | 0/0 | 1/1 |
| mage | 0/3 | 0/2 | 0/3 | 0/3 | 0/1 | 0/0 | 1/1 |

Every present numbered set starts at 1 without a gap and has at most 4
files. EN also holds `characters/aldor/defend/2.wav` and
`characters/aldor/pickup/2.wav`, which no built name reaches; every RU key is
a built name. EN files are 22050 Hz mono, 35 PCM 16-bit and 53 with format
tag 2 at 4 bits; RU files are all PCM 16-bit, 22050 Hz mono. No `speech.res`
key in either root lies under a voice-name directory.

**Confidence.** High: the complete key lists of both archives were read.

**Evidence.** [EXP-2092](../experiments/EXP-2092-rom2-unit-voices/RESULTS.md), question 2.
**Format surface.** `sfx.res` keys; `formats/rom2-res/format.md`.
