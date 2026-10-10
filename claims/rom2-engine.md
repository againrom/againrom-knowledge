# Claim registry — ROM2 ENGINE (compiled-code relationship to ROM1)

Level 2 ledger. Index: [registry.md](registry.md). Legend: `✔ promoted` (also in a spec) ·
`● active` · `✖ retracted`. IDs are permanent.

Not a file format — no claim here describes a stored byte layout; see
`experiments/EXP-2001-engine-sharing/format-surfaces.tsv`. This ledger is a static, structural
comparison of `gameversions/en/rom.exe` (ROM1) against the four ROM2 binaries named in
`pipeline/ROM2-MULTIPLAYER.md` (`allods2.exe`, `engine32.dll`, `scenario.dll`, `a2server.exe`),
never executed. `WININET.DLL`/`SHLWAPI.DLL` under the same ROM2 install root appear only as an
independent null-population control (`R2-ENGINE-003`), never as a ROM1 or ROM2 subject.
`claims/rom2-asset.md` and `formats/rom2*` belong to the concurrent `EXP-2000` and are not cited
here. Several rows below were corrected by a confidence-review pass; the overturned wording is in
`claims/retracted.md` under "ROM2 engine-sharing floor and null-population corrections".

`R2-ENGINE-017` to `R2-ENGINE-032` are a second kind of claim: a static read of the magic routines of
`allods2.exe` (spell cast, power, cost, apply arms, resolver and regeneration), with ROM1's `MAGIC-*`
and `HERO-*` claims cited for each difference. They describe behaviour in code and the numeric spell
table of `data.bin`, not a stored byte layout.

`R2-ENGINE-033` to `R2-ENGINE-040` are a third kind: a static read of the single-mission load of
`allods2.exe` and censuses of the installed ROM2 maps, scripts, text and sprite files. They describe
what the load reads and which ROM1 grammars still parse, not a new stored byte layout.

## Engine sharing between ROM1 and the ROM2 binaries

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-001 | `engine32.dll` is a registry/configuration helper with two exports and no graphics or audio imports; only `allods2.exe` imports it. | High | ● active | [EXP-2001](../experiments/EXP-2001-engine-sharing/) |
| R2-ENGINE-002 | `scenario.dll` is named in both `allods2.exe` and `a2server.exe` but in neither import table, so both most likely load it dynamically. | High / Medium | ● active | [EXP-2001](../experiments/EXP-2001-engine-sharing/) |
| R2-ENGINE-003 | Same-binary self-collision and cross-binary chance-match rate are different measurements; the corrected cross-binary `normHash` floor pools to 1.29% at 51 or more instructions. | High | ● active (amended, partially retracted) | [EXP-2001](../experiments/EXP-2001-engine-sharing/) |
| R2-ENGINE-004 | 110 `rom.exe` library names also name a ROM2 function, but 107 have several candidates, so the 99.1% `normHash` agreement is an any-candidate figure. | High | ● active | [EXP-2001](../experiments/EXP-2001-engine-sharing/) |
| R2-ENGINE-005 | Against the corrected floor five of six instruction-count buckets clear it for identical matches; only the 1-5 instruction bucket fails. | High / Medium | ● active (amended, partially retracted) | [EXP-2001](../experiments/EXP-2001-engine-sharing/) |
| R2-ENGINE-006 | The 109 of 110 synthetic names in the large identical tier discriminate nothing: the base rate over the whole 51+ population is 99.69% synthetic. | High / Unknown | ● active (amended, partially retracted) | [EXP-2001](../experiments/EXP-2001-engine-sharing/) |
| R2-ENGINE-007 | `allods2.exe` and `a2server.exe` share 77.35% of `normHash` values with each other, so no per-binary floor may be squared; 75.4% of large matched functions match both. | High / Medium | ● active (amended, partially retracted) | [EXP-2001](../experiments/EXP-2001-engine-sharing/) |
| R2-ENGINE-008 | ROM1 record anchors survive unevenly: session `Serialize` is a mnem-match in both ROM2 binaries, the campaign writer and reader stay absent. | High / Medium / Low | ● active (amended, partially retracted) | [EXP-2001](../experiments/EXP-2001-engine-sharing/) |
| R2-ENGINE-009 | New ROM2 code is smaller than first measured: `a2server.exe` 39.0%, `allods2.exe` 31.9%, `scenario.dll` 12.3%, `engine32.dll` 80.3% of non-thunk functions. | High / Low | ● active (amended) | [EXP-2001](../experiments/EXP-2001-engine-sharing/) |
| R2-ENGINE-010 | The identical-plus-changed rate per ledger family does not separate along an engine-versus-content line. | Medium | ● active | [EXP-2001](../experiments/EXP-2001-engine-sharing/) |
| R2-ENGINE-011 | The "changed" tier is a shared-string co-reference detector with a same-binary floor of 4.0% at 51 or more instructions, not a modified-copy signal. | High / Medium | ● active | [EXP-2001](../experiments/EXP-2001-engine-sharing/) |

### R2-ENGINE-001

**`engine32.dll` is a small registry/configuration helper, not a rendering or sound engine — the name is the only evidence for that reading, and the import table refutes it.** Its own PE import table lists only `ADVAPI32.DLL`, `KERNEL32.DLL`, `USER32.DLL` — no `DDRAW`, `DSOUND`, `D3D`, or any other graphics/audio API. It exports exactly two named routines, `InitEngine` and `DeinitEngine`, and four of its 118 functions reference the registry path `Software\1C\Allods 2` together with the string `RESOLUTION`. Among the four ROM2 binaries, only `allods2.exe` (the client) statically imports it; `a2server.exe` and `scenario.dll` do not reference it at all. `a2server.exe`'s own import table is not a stripped, headless-server set: it imports `DDRAW.DLL`, `DSOUND.DLL`, `SMACKW32.DLL`, `MSACM32.DLL`, `WINMM.DLL`, `WSOCK32.DLL`, `OLE32.DLL`, `GDI32.DLL` and the MFC-typical `COMCTL32`/`COMDLG32`/`SHELL32`/`WININET`/`WINSPOOL.DRV` — the same kind of subsystem set `rom.exe` itself imports, not a graphics/sound-free build

**Confidence.** High (complete, direct PE import/export table reads and a full string-reference scan over `engine32.dll`'s own 118 functions; no decompilation or behavioural inference)

### R2-ENGINE-002

**`scenario.dll` is the one ROM2 module shared between the client and the dedicated server by construction, and it is not statically linked by either.** The literal string `scenario.dll` occurs twice in `allods2.exe`'s own raw bytes and twice in `a2server.exe`'s, while neither binary's PE import table contains an entry for it (0 of `allods2.exe`'s 504 imports and 0 of `a2server.exe`'s 494 name it). The most parsimonious reading is that both binaries load it dynamically (e.g. `LoadLibrary`) rather than linking it; no `LoadLibrary` call site referencing that exact string was traced, so the loading mechanism itself is inference, not a read instruction

**Confidence.** **High** for the byte/table facts (exact string counts, exact import-table absence) / **Medium** for "loaded dynamically" as the mechanism, which is the simplest reading and not a traced call site

### R2-ENGINE-003

**Same-binary self-collision (`rom.exe`'s own function population matching itself) and cross-binary false-match rate (an unrelated binary matching `rom.exe` by chance) are different measurements. A correction pass retracted this row's own original use of the same-binary rate as the cross-binary floor and replaced it with an independently measured one, which is far lower at every bucket but the shortest.** Over `rom.exe`'s own 6,830 non-thunk functions, the share sharing a `normHash` with a *different* address in the *same* file (unchanged by the correction) is 75.8% (1–5 instr.), 71.0% (6–10), 57.4% (11–20), 29.3% (21–50), 7.5% (51–100), 6.6% (101+); `mnemHash` collides more, 46.2% overall against `normHash`'s 36.8%. This is an exact count of intra-program code duplication, not a cross-binary chance-collision bound. `WININET.DLL` (1,614 functions, linker 6.0, 1998-04-29) and `SHLWAPI.DLL` (589 functions, linker 6.0, 1998-03-11) — 2,200 non-thunk functions together, under the same ROM2 install root, Microsoft-authored Win32 system libraries with no Allods game code — supply the independent null, censused through the identical unmodified `FuncCensus.java`. Their linker major version (6.0) matches neither `rom.exe` (5.0) nor `allods2.exe`/`scenario.dll` (5.2) nor `a2server.exe` (5.0); "unrelated" rests on authorship and installation path, not toolchain match. Matched against `rom.exe` and scaled from the 2,200-function null to the 18,639-function non-thunk ROM2 union (`allods2.exe`+`engine32.dll`+`scenario.dll`+`a2server.exe`, 8.47× the null's size) by `1-(1-p)^scale` — the correct exposure model for "at least one collision," not linear scaling, which produces impossible values above 100% at short buckets — the estimated cross-binary `normHash` floor is 71.45% (1–5), 1.00% (6–10), 30.01% (11–20), 1.21% (21–50), 2.60% (51–100, rule-of-three estimate on 0 hits in 968 trials), 2.53% (101+, 0/993), pooled 1.29% (0/1,961) for instrCount ≥ 51. A parallel null measurement for the strict mnemHash+instrCount+byteLen tier (`R2-ENGINE-008`) gives the same pooled ≥51 floor, 1.29% (0/1,961)

**Confidence.** High for both floors as exact/complete counts and a stated, tested formula, not a point guess at the zero-hit buckets (`self-collision-by-size.tsv`, `null-control.tsv`, `TestWriteNullControlScalingIsExponentialNotLinear`) / the null is unrelated-*vendor* code, not same-toolchain code — a same-era MFC/CRT-linked, non-Allods binary would be a stronger control and remains unavailable inside this repository's lawful sources; if two MFC-linked programs share more incidental runtime boilerplate with each other than either shares with a Microsoft system DLL, the true chance floor for an MFC-to-MFC comparison could sit higher than measured here, narrowing every downstream margin that reads against it

**Amended.** A correction pass changed this claim; the overturned wording is in [`retracted.md`](retracted.md) under "ROM2 engine-sharing floor and null-population corrections".

### R2-ENGINE-004

**Ghidra's own Function ID analyzer supplies a true-positive control of 110 real (non-`FUN_XXXXXXXX`) `rom.exe` names that also name a function in at least one ROM2 binary, but 107 of those 110 (97.3%) have more than one same-named ROM2 candidate, so "agrees on `normHash`" is almost always an any-candidate claim — at least one of several same-named candidates matches — not a one-to-one comparison.** A name's set of ROM2 candidates does not depend on which `rom1_entry` address carries it (an overloaded/templated name such as `Add`, 16 `rom.exe` addresses, always the same 66-candidate ROM2 set), so 110 distinct names is the right denominator, not the 219 (name, address) pairs those names occupy. Read on the any-candidate basis, the original headline stands: 109 of 110 names (99.1%) have at least one ROM2 candidate matching by `normHash`; the sole disagreement is `FID_conflict:operator-` (2 candidates, neither matches). Read on the strict one-to-one basis — exactly one ROM2 candidate sharing the name, no aggregation — only 3 of 110 names qualify (`FillInToolInfo`, `OnFinalRelease`, `_IsEqualGUID`); all 3 of 3 (100%) agree on `normHash`. Every one of the 110 shared names is a statically-linked MFC/CRT library routine (`CWnd`, `CArray<>`, `AfxGetThread`, `CCmdUI`, `CRect`, `AddTail`, and similar) — none is Allods-specific

**Confidence.** High for the counts (219 rows, 110 distinct names, exact; `fid-control.tsv`'s own summary line) / the one-to-one true-positive rate is honestly small-n (3 names) and is reported as such, not inflated to the 110-name any-candidate figure

### R2-ENGINE-005

**A correction pass retracted this row's own six-bucket floor-clearance verdict, which was read against the invalid same-binary floor `R2-ENGINE-003` has since withdrawn; read against the corrected, independently measured cross-binary floor, five of the six instruction-count buckets clear it, not three, and the specific buckets differ.** Bucketed by the cited ROM1 function's own instruction count (claim-cited population, 1,708 addresses: 432 identical / 91 mnem-match / 112 changed / 937 absent / 136 not resolved to a function in this Ghidra pass / 0 thunks — 91 of the 1,708 move into the new mnem-match tier introduced by this correction, 88 previously absent and 3 previously changed), the identical rate against `R2-ENGINE-003`'s corrected cross-binary floor is: 27.27% vs. 71.45% floor at 1–5 instr. (n=22, does not clear), 62.75% vs. 1.00% at 6–10 (n=102, clears 62.7×), 62.30% vs. 30.01% at 11–20 (n=183, clears 2.08×), 34.07% vs. 1.21% at 21–50 (n=405, clears 28.2×), 20.94% vs. 2.60% at 51–100 (n=320, clears 8.05×), 7.96% vs. 2.53% at 101+ (n=540, clears 3.15×). Only the shortest bucket fails to clear; every other bucket clears by at least a factor of 2. The new mnem-match tier is a separately weaker signal against its own floor: 13.64% vs. 97.74% (1–5, does not clear), 0.98% vs. 2.97% (6–10, does not clear), 2.19% vs. 29.64% (11–20, does not clear), 10.12% vs. 1.21% (21–50, clears 8.36×), 7.81% vs. 2.60% (51–100, clears 3.00×), 3.15% vs. 2.53% (101+, clears 1.24×, weak). Restricted to instrCount ≥ 51 (860 of the 1,708 addresses, matching `R2-ENGINE-006`/`008`'s population): identical is 110/860 (12.79%) against the pooled 1.29% floor (9.9× clears); mnem-match is 42/860 (4.88%) against the same 1.29% floor (3.8× clears); "changed" is 107/860 (12.44%) — 5/405+22/320+85/540 across the same buckets — against `R2-ENGINE-011`'s own, unrelated, same-binary string+size floor of 4.0% (3.1× clears; no cross-binary null exists for that signal, see `R2-ENGINE-011`)

**Confidence.** High for the arithmetic — exact counts in both populations, no sampling — / Medium overall for what the clearance implies about genuine sharing, because `R2-ENGINE-003`'s null is unrelated-vendor, not same-toolchain, code, and that gap is not closed here

**Amended.** A correction pass changed this claim; the overturned wording is in [`retracted.md`](retracted.md) under "ROM2 engine-sharing floor and null-population corrections".

### R2-ENGINE-006

**A correction pass retracted this row's own conclusion that the large-function identical tier's synthetic-name share "rules out" a library-code explanation: the comparison never had a measured base rate, and with one measured, it discriminates nothing.** The raw count is unchanged and reproduces exactly: 109 of 110 (99.1%) of the ≥51-instruction identical tier (`R2-ENGINE-005`) carry a synthetic `FUN_XXXXXXXX` name in `rom.exe`, not a Ghidra Function-ID-recognized one. But the base rate of real names across the *whole* instrCount≥51 `rom.exe` population is already 99.69% synthetic — only 6 of 1,961 functions carry any real name at all, regardless of whether they are library or game code, because Ghidra's Function ID recovery in this pass named a small fraction of the whole population. A subset drawn entirely from `R2-ENGINE-004`'s own MFC/CRT library control would be expected to show a similarly near-total synthetic share under this same base rate. Whether the large-function identical tier is mostly shared runtime library code or mostly Allods-specific code remains open; no comparison this experiment ran separates the two

**Confidence.** The 109/110 count stays High (exact, complete, `large-identical-name-check.tsv`); the discriminating conclusion is retracted to Unknown — no instrument here separates library-code origin from game-code origin for this tier

**Amended.** A correction pass changed this claim; the overturned wording is in [`retracted.md`](retracted.md) under "ROM2 engine-sharing floor and null-population corrections".

### R2-ENGINE-007

**A correction pass retracted this row's own "floor squared" argument that allods2.exe/a2server.exe co-occurrence is statistically surprising: the two binaries are not independent code populations, so no per-binary floor may be squared to estimate their joint rate.** `allods2.exe` and `a2server.exe` share 4,312 of 5,575 distinct `normHash` values with *each other* directly (77.35%, all instruction counts; 1,417/2,174=65.2% and 1,417/2,335=60.7% restricted to instrCount≥51) — independent of any relationship to ROM1 at all. The co-occurrence fact this row originally reported stands and reproduces under the corrected classification: restricted to instrCount ≥ 51 (259 claim-cited matched addresses under identical+mnem-match+changed, up from 220 under the original identical+changed), 173 (66.8%) include `a2server.exe`, 86 (33.2%) are `allods2.exe`-only; over the full instrCount≥51 `rom.exe` population against the whole ROM2 union (not just claim-cited addresses), 676 of 1,961 functions match at least one ROM2 binary and 510 of those 676 (75.4%) match both. Most matched ROM1-derived code is present in the dedicated server, not client-only — because the client and server binaries already share most of their own compiled code, a direct, non-statistical explanation. A shared static library or common source tree is the parsimonious reading; this experiment did not trace which

**Confidence.** High for the co-occurrence and binary-overlap counts (complete, exact; `binary-overlap.tsv`, independently cross-checked with raw `awk`/`sort`/`comm` outside the Go tool) / Medium for "shared library or common source tree," not traced to a specific mechanism

**Amended.** A correction pass changed this claim; the overturned wording is in [`retracted.md`](retracted.md) under "ROM2 engine-sharing floor and null-population corrections".

### R2-ENGINE-008

**Named ROM1 "record" anchor functions do not survive into ROM2 as uniformly as one another. A correction pass added a strict mnemHash+instrCount+byteLen matching tier and recovered one of the three anchors this row originally classified absent.** Of five cited by standing claims: session `Serialize` `R0270` (132 instr., 462 bytes) is a mnem-match in both `allods2.exe` (`R2.0001`) and `a2server.exe` (`L2.00001`) — identical mnemHash, instruction count and the byte length, with its twelve small-immediate operands changed at most positions (as an order-independent set, 6 of 12 ROM1 values — `8, 400, 1000, 48, 2632, 2633` — recur somewhere in the ROM2 candidate's own list; 6 — `48436, 48836, 43048, 2508, 43452, 2636` — do not, and the ROM2 candidate's own 6 non-recurring values are `4000, 51076, 59076, 42792, 4908, 43196`; which operand corresponds to which field is a decode task outside this experiment's structural-comparison scope). Campaign writer `R1294` (279 instr., 779 bytes) and campaign reader `R0434` (700 instr., 2,130 bytes) remain absent under the same strict instrument — no ROM2 function shares their mnemHash+instrCount+byteLen either — so the anchor list narrows from three absent to two. The small Human constructor `R0497` (34 instr., 114 bytes) classifies identical in both `allods2.exe` and `a2server.exe`, its three small immediates (12, 16, 8) unchanged; the cell-record serializer `R1521` was again not resolved to a function entry in this Ghidra pass (one of the 136 `rom1-not-found` addresses, a tooling/version limitation, not a ROM2 fact). Separately, a raw literal search of the five component values these claims name (488, 52, 400, 1,000, 48, 2,508) found each already common enough in `rom.exe`'s own population, except 488 (9 functions, 0.13%) and 2,508 (1 function), to make bare presence in ROM2 weak evidence on its own; 2,508 in particular is asymmetric (1 ROM1 function, 74 `allods2.exe` and 71 `a2server.exe` functions), not decoded further (out of scope). The literal 2 was never searchable: the census instrument's own capture floor is `[4, 0x100000)`, so its "0 occurrences" row is a tool artifact, not a finding

**Confidence.** Medium for the two remaining absent classifications and for `R1521`'s non-resolution being a tooling limitation (both unchanged from the original grade) / the session `Serialize` mnem-match carries `R2-ENGINE-003`'s own confidence split (High for the exact skeleton match, Medium for what it implies given the unrelated-vendor null) / Low for the Human-constructor identical match on its own (34 instr. sits in the 21–50 bucket, where the corrected cross-binary floor is 1.21%, so this one match is not itself distinguishing) / the literal-search table is not scored — it is reported as inconclusive

**Amended.** A correction pass changed this claim; the overturned wording is in [`retracted.md`](retracted.md) under "ROM2 engine-sharing floor and null-population corrections".

### R2-ENGINE-009

**New ROM2 code (no `rom.exe` counterpart under either the `normHash` or the strict mnemHash+instrCount+byteLen signal) is smaller than this row originally measured. A correction pass's strict instrument recovers a same-skeleton `rom.exe` match for part of what the `normHash`/string-only pass counted as new.** Recovered share: 388 of `a2server.exe`'s originally-new 3,657 (10.6%), 494 of `allods2.exe`'s 3,328 (14.8%), 13 of `scenario.dll`'s 170 (7.6%), 2 of `engine32.dll`'s 96 (2.1%). Corrected new-code counts: `a2server.exe` 3,269/8,374 non-thunk functions (39.0%, 1,269,077 bytes, down from 43.7%/1,319,223 bytes), `allods2.exe` 2,834/8,875 (31.9%, 965,167 bytes, down from 37.5%/1,031,216 bytes), `scenario.dll` 157/1,273 (12.3%, 21,989 bytes, down from 13.4%/23,631 bytes), `engine32.dll` 94/117 (80.3%, 15,098 bytes, down from 82.1%/15,323 bytes) — every binary's share is smaller than originally published, the per-binary ranking is unchanged, and this remains consistent with `R2-ENGINE-001`: `engine32.dll` still holds little beyond its registry-helper role. The recovery rate carries `R2-ENGINE-003`'s same null-population limitation: the true new-code share could be smaller still against a same-toolchain (MFC/CRT) control this experiment does not have. Clustering the corrected new population by first-string-reference keyword: 95.8% (`a2server.exe`, 3,131/3,269) and 96.0% (`allods2.exe`, 2,721/2,834) still reference no string at all (`NOSTR`) and are not clustered by this method; the labeled remainder is unchanged in shape (largest named clusters: `GRAPHICS`, 24 functions each in `allods2.exe`/`a2server.exe`; `MAIN`, 5 each)

**Confidence.** High for the per-binary new-function counts and the byte totals (complete population, exact) / Low for the string-keyword clusters, which label under 5% of the new population

**Amended.** A correction pass changed this claim; the overturned wording is in [`retracted.md`](retracted.md) under "ROM2 engine-sharing floor and null-population corrections".

### R2-ENGINE-010

**The per-ledger-family identical+changed rate does not separate cleanly along an engine-versus-content line, refuting H3 as precisely stated; a correction pass's new mnem-match tier does not change this, and the 8 families below reproduce their originally published rate exactly before adding it.** Restricted to instrCount ≥ 21 (`family-breakdown-min21instr.tsv`): infrastructure-flavoured families score both high and low (`res` 46.2% identical+changed, n=39, +0.0pp with mnem-match added since `res` has none in this tier; `move` 21.6%, n=111, +2.7pp to 24.3%; `terrain` 21.0%, n=186, +1.6pp to 22.6%; `ai` 16.8%, n=244, +2.9pp to 19.7%), and so do content-flavoured ones (`menu` 46.7%, n=45, +2.2pp to 48.9%; `dialogue` 31.6%, n=57, +3.5pp to 35.1%; `town` 28.5%, n=239, +2.1pp to 30.5%; `tavern` 12.9%, n=31, +3.2pp to 16.1%). The two highest-scoring families (`res`, `menu`) and the lowest (`tavern`) fall on opposite sides of an engine/content split, not the same side each, both before and after adding mnem-match. `R2-ENGINE-005`'s instruction-count floor comparison remains the one dimension this experiment found that does separate cleanly

**Confidence.** Medium. Several families are small enough (`fame` n=3, `spr256` n=10, `inv` n=1) that their own rate is not a reliable estimate at any n; the comparison is restricted to families with n ≥ 30

### R2-ENGINE-011

**The "changed" classifier detects a shared referenced string (≥6 characters) together with a similarly-sized function body. It is a string co-reference detector, not a code-similarity or "modified copy" signal, and its own same-binary chance-collision floor must not be conflated with `R2-ENGINE-003`'s `normHash` floor.** With the correction pass's new mnem-match tier now sitting between identical and changed in classification priority, "changed" is what remains after both exact code-shape (identical) and same-skeleton-different-immediates (mnem-match) are removed: a "changed" pair shares no code-shape signal at all, only a referenced string and a coarse size bracket. Over `rom.exe`'s own 6,830-function non-thunk population and the same instruction-count buckets as `R2-ENGINE-003` (unchanged by the correction pass), the same-binary string+size collision rate is 0.0% (1–5, n=269), 0.24% (6–10, n=844), 0.12% (11–20, n=1,673), 0.67% (21–50, n=2,083), 2.07% (51–100, n=968), 5.94% (101+, n=993); pooled over instrCount ≥ 51 (1,961 functions), 4.0% (79/1,961). This floor is same-binary only; no independent-null equivalent was built for the string+size signal, an unresolved gap this correction pass did not close. Among the 112 claim-cited "changed" rows, the referenced strings are not highly repetitive — 96 distinct string hashes account for 112 matches, the most frequent recurring in only 4 rows — but the instrument works from a hashed string token (length + a truncated hash, no verbatim text, a repository hygiene requirement), so genericity (a common resource label vs. a distinctive one) cannot be judged from the evidence alone; a "changed" classification is read as "shares a reference," never as "is a modified copy of"

**Confidence.** High for the same-binary floor counts (complete population, `string-size-self-collision.tsv`) / Medium for the string-diversity read (indirect, from hash+length only, an acknowledged instrument limit)

## Casting, cost and the spell table

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-017 | `allods2.exe` casts through `L2.00027(caster, target-or-null, x, y)`; the spellbook is an array at actor+0x140 indexed by spell id, and a spell is a 0x14-byte object over a table row. | High | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |
| R2-ENGINE-018 | A cast costs the table Mana Cost column once; only a mage casting without an item pays, and no skill or stat term enters the cost. | High / Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |
| R2-ENGINE-019 | The spell table holds 29 spells with 22 numeric columns and an Effects string; client and server tables are identical, and four spells are new against ROM1. | High | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |

### R2-ENGINE-017

The cast routine is `L2.00027`. The actor tick `L2.00028` calls it at `L2.00029` and `L2.00030`; a direct-call census of `L2.00027` was not widened past the actor tick. Delivery column 2 (projectile) delays the apply by `distance / Effect Speed` ticks (`L2.00031` is the distance); ids 10 and 11 force a delay of 5. Prismatic Spray (id 11) takes a separate path: refill `L2.00032`, then the victim fan `L2.00033` (`R2-ENGINE-027`). Every other spell reaches the apply routine `L2.00034`, directly or through `L2.00035`.

The spell object is 0x14 bytes (constructors `L2.00036` and `L2.00037`, copy `L2.00038`): +4 row pointer, +8 id byte (the row index, found by name), +9 range byte, +0xa defensive flag, +0xc cost word, +0xe damage base byte, +0xf damage spread byte, +0x10 duration word. The spellbook is a `CObArray` at actor+0x140 indexed by spell id from 1 (helpers `L2.00039` SetAt and `L2.00040` GetAt). The derive refills every book spell through `L2.00041`, which calls `L2.00032`. The timed-effect bitmask is at actor+0x144 (`L2.00042` tests it, `L2.00043` finds an effect).

An item cast (actor+0x68 nonzero) bypasses the book: its power comes from the item's castSpell effect (`R2-ENGINE-020`).

ROM1 has the same cast shape (`MAGIC-CAST-003`, `R0268`), the same 0x14-byte object and the same sparse book subscripted by id (`MAGIC-SPELL-001`, `MAGIC-BOOK-002`).

**Confidence.** High: the routine, object fields and book array were read from constructors, helpers and the actor tick, not inferred from names.

### R2-ENGINE-018

The cost word at spell+0xc is table column 1 (Mana Cost), copied unchanged by the fill; it never reads P. In the cast, a mage (mage bit `L2.00044`, actor+0x4c & 4) with no item (actor+0x68 zero) whose cost exceeds current mana (actor+0x9a) gets a return of 0 and no cast; otherwise mana falls by the cost. Non-mages and item casts are not charged. Costs run from 3 (Fire Arrow, Ice Missile, Diamond Dust) to 100 (Blizzard, Acid Stream, Invisibility, Summon); the full list is in `evidence/catalogue.tsv`.

ROM1 states the same flat single subtraction (`MAGIC-CAST-003`). The class gate was read here from the ROM2 cast only; the ROM1 claim was not re-read for it.

The Scroll Cost and Book Cost columns (20 and 21) are not read by the cast or the apply routine.

**Confidence.** High for the arithmetic of the charge. Medium that no other deduction exists: the population read is the cast, the apply routine and the Prismatic path, not a census of every writer of actor+0x9a.

**Unknown.** The consumer of columns 20 and 21 (names suggest shop prices; unread). Whether casting trains a school as `MAGIC-TRAIN-018` records for ROM1.

### R2-ENGINE-019

`data.bin` group H (read by `tools/r2databin`) lists 29 spells, ids 1 to 29, each with a name, 22 numeric columns and a trailing Effects string. Column order: Complication, Mana Cost, Sphere (1 Fire, 2 Water, 3 Air, 4 Earth, 5 Astral), Item, Spell Target, Delivery, Max Range, Effect Speed, Distribution, Radius, AreaAffect, AreaDuration, AreaFreq, ApplyMethod, Duration, Freq, damageMin, damageMax, Defensive, SkillOffset, scrollCost, bookCost. The client table (`world.res`) and the server table (`world_srv.res`) agree row for row.

New against ROM1 (`MAGIC-SPELL-001`, 28 spells): Ice Missile (5), Blizzard (7), Diamond Dust (16), Summon (25). Present in ROM1 and absent here: Fire Sacrifice, Freezing Cloud, Meteor Storm. Id 29 is Slow, id 28 Curse. Effects strings are length-prefixed; the stored keys are `protectionFire`, `protectionWater`, `protectionAir`, `protectionEarth`, `health`, `Invisibility`, `scanRange`, `toHit`, `speed`, `absorbtion`, and a Stone Curse string whose first key carries a literal quote (`R2-ENGINE-025`).

**Confidence.** High: a mechanical walk of the table, with the client and server bodies compared.

## Power and school skills

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-020 | Spell power is `P = max(skill[Sphere] + Mind - 30, 0)`, clamped to 255 in the apply and refill paths; the Prismatic cap reads a low byte; an item cast takes P unclamped. ROM1 clamps at 100. | High | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |
| R2-ENGINE-021 | In a cast, Mind enters only the power term (it also enters speed, experience, the AI cast choice and two other readers); Spirit enters maximum mana and the resistance words; stat caps depend on the mage bit and unit type, not one cap of 50. | High / Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |
| R2-ENGINE-022 | Fill rules `L2.00045` turn P into range, damage base and spread, and a duration word; every spell uses one of four duration rules. | High / Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |
| R2-ENGINE-023 | A school skill buys power and nothing else: it adds to damage, range, duration and magnitude through P, and no cast reads it for cost or a success roll. | High / Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |

### R2-ENGINE-020

`L2.00046` returns the power byte for a spell cast: the school skill word at caster+0xa8+2*Sphere, plus the Mind word at +0x88, minus 30, floored at 0. `L2.00046` returns the low byte of the sum, so it wraps at 256 and has no upper clamp; its one reader here is the Prismatic victim cap (`R2-ENGINE-027`, call at `L2.00047`). The wrapper `L2.00048` works on byte arguments and also keeps a low byte. Only the refill `L2.00032` and the apply routine (`L2.00049` to `L2.00050`) clamp to 255. The wrap case needs `skill + Mind - 30` above 255, which is a school skill bonus above 85 over the P 170 ceiling, so it is unreachable without such a bonus.

An item cast reads the castSpell effect of the item (found through `L2.00051`) and takes the signed word at effect +0x42 as P without a clamp.

ROM1 clamps to 0..100 (`MAGIC-POWER-004`) and reads the one unclamped consumer elsewhere (`MAGIC-CEIL-013`). With the stat caps of `R2-ENGINE-021`, a mage at the table skill ceiling of 100 and the largest Mind cap of 52 reaches P = 122 without effects (`evidence/stat-caps.tsv`); a Mind of 100 from effects gives 170. P above 170 needs a school skill bonus from effects (`R2-ENGINE-021`).

One point of P near its ceiling is worth, per `evidence/point-value.tsv`: damage base and spread at most +1 each (Blizzard +1 base at 254 to 255); range +0 or +1 byte per 30 points; effect duration about +2.5 percent of its value (the base 1.025 power law of `R2-ENGINE-022`); protection magnitude +0.5 (P/2, integer).

**Confidence.** High for the apply path and the item path. The power byte of `L2.00046` wraps at 256; the three sites do not agree on the clamp, as in ROM1 (`MAGIC-CEIL-013`).

### R2-ENGINE-021

Stat words in the actor: +0x84 Body, +0x86 Reaction, +0x88 Mind, +0x8a Spirit. The derive `L2.00052` sets each stat to the smaller of itself and a cap plus the effect modifier byte (+0xd4 to +0xd7). The cap depends on the mage bit (`L2.00044`) and the unit type test `L2.00053` (type word +0xe equal to 0x22 or 0x24):

| Case | Body | Reaction | Mind | Spirit |
|---|---|---|---|---|
| non-mage, other type | 52 | 50 | 48 | 46 |
| non-mage, type 0x22 or 0x24 | 50 | 52 | 46 | 48 |
| mage, other type | 48 | 46 | 52 | 50 |
| mage, type 0x22 or 0x24 | 46 | 48 | 50 | 52 |

ROM1 has one cap of 50 for all four (`HERO-CAP-015`). The effect arms (`L2.00054` kinds 2 to 5) clamp the live stat to 100. The skill loop clamps each of the five school slots to 0..100 and then adds the effect bonus at +0xe8+2i without a second clamp; ROM1 clamps twice (`HERO-SKILL-009`).

Mind readers (`EnumRefs disp:88`: 383 hits, word-wide 37 hits in 21 owners): the readers read here are the power term, the derive (also speed, `((Mind + Reaction) / 25 + 4) * 256`), the experience award `L2.00055` (scale `Mind/30 + 0.25`, unit types 0x21 to 0x3f), the AI spell choice `L2.00056` (Mind above 59 adds a skip roll, as `MAGIC-AI-012` records for ROM1), a value routine `L2.00057` and a requirement check `L2.00058`. Spirit (`disp:8a`: 43 hits, 16 word-wide owners) enters maximum mana (`R2-ENGINE-031`) and the five resistance words (`R2-ENGINE-029`). Neither stat is read by the cost or by the fill rules.

**Confidence.** High for the cap table, the clamp placement and the readers listed, which were read instruction by instruction. Medium for the census completeness: the word-wide Mind owners beyond those named were not read.

**Unknown.** The base-stat snapshot readers `L2.00059` and `L2.00060`; the other word-wide owners.

### R2-ENGINE-022

The fill `L2.00045` takes the argument byte P and writes the cached fields of `R2-ENGINE-017`. With `f = P/30.0 + 1.0`:

- Range byte = Max Range + (id 23 Teleport: P/3; otherwise P/30 when Max Range is nonzero, else 0), byte-wrapped.
- Damage base = `ftol(f * damageMin) & 0xff` when damageMin is positive, else 0.
- Damage spread = `ftol(f * damageMax - base) & 0xff` when damageMax is positive (floating subtraction, then truncation).
- Duration word, first match: Duration column positive: id 12 gives `P << 4`; id 18 gives `(P*15*16*2)/100`; any other id gives `ftol(Duration * 1.025^P * 16)`. Otherwise AreaDuration positive: `AreaDuration*16 + (P*16)/10`. Otherwise 0.

The duration base is 1.025: the immediate pushed before `pow` is `0x3ff0666666666666` (bytes `66 66 f0 3f 66 66 66 66`) at `L2.00061`, `L2.00062`, `L2.00063`, `L2.00064` and `L2.00065`; a scan of the code section for the 1.05 immediate finds none. This is ROM1's law (`MAGIC-CEIL-013`). Invisibility differs: ROM2 gives `P << 4` where ROM1 uses `1.05^P` capped at 65000. Constants read: 30.0 at `L2.00066`, 1.0 at `L2.00067`, 16.0 at `L2.00068`. The damage factor, the range term, the Teleport range and the Prismatic cap equal ROM1's (`MAGIC-CEIL-013`).

The common apply tail adds the area duration for area objects (Distribution not 1): `(AreaDuration << 4)` plus `(P << 4)/10` when that is nonzero; Distribution 5 forces mode 2 and a duration of 0.

`evidence/samples.tsv` lists every spell at P of 0, 30, 60, 100, 122, 170 and 255.

**Confidence.** High for the integer branches. Medium for the exact floating results: the tool uses 53-bit doubles, the binary uses the x87 path whose precision control was not read, and `evidence/power-domain.tsv` flags the P values where a result lies within 1e-9 of an integer (Wall of Fire, Blizzard, Prismatic Spray and Drain Life spread). Those rows may differ by 1 under 64-bit intermediates.

### R2-ENGINE-023

Reads of the school skill: only the power term (`R2-ENGINE-020`), the wrapper and the refill. Through P it scales damage base and spread (`R2-ENGINE-022`), range (+1 byte per 30 points), duration (1.025 power), and each arm's magnitude (`R2-ENGINE-024`). A spell's cost is the table column (`R2-ENGINE-018`). The cast routine and the apply routine contain no random draw against the school skill, so there is no cast failure or success roll in the path read. The SkillOffset column (19) is 0 in all 29 rows.

ROM1 states the same one-term use of the skill (`MAGIC-POWER-004`), with its ceiling of 100.

**Confidence.** High for the readers listed. Medium for "nothing else": the claim covers the cast, fill and apply routines and the book refill, not every reader of the skill array; the AI and training readers were not read.

**Unknown.** Whether casting raises the school skill in ROM2 (`MAGIC-TRAIN-018` is ROM1's rule); the SkillOffset consumer.

## Apply arms and shared rules

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-024 | Fifteen apply arms give the effect magnitudes: protections P/2, Shield P/10+3, Haste P/15+1, Bless and Curse 4P/5+20, Light and Darkness P/30+1, Poison Cloud a health change of -2*(P/45.0+1). | High / Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |
| R2-ENGINE-025 | Stone Curse draws a random duration up to `4.8*P` units for a spell cast (`2.4*P` for an item cast), cut by Earth resistance; its Effects string begins with a quote that probably stops the kind parsing. | Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |
| R2-ENGINE-026 | Heal, Drain Life, Summon, Control Spirit and Teleport each have a unique arm; Summon and Teleport depend on P, Control Spirit does not read it. | High / Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |
| R2-ENGINE-027 | Prismatic Spray hits min(P/20 + 2, 7) victims, where P for a book cast comes from the power byte and for an item cast from the castSpell effect. | High | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |

### R2-ENGINE-024

The apply routine `L2.00034` dispatches by spell id through the table at `L2.00069`. Ids and arms: 6 `L2.00070`; 4, 8, 13, 19 `L2.00071`; 12 `L2.00072`; 14 `L2.00073`; 15 `L2.00074`; 17 `L2.00075`; 18 `L2.00076`; 20 and 28 `L2.00077`; 21 and 29 `L2.00078`; 22 `L2.00079`; 23 `L2.00080`; 24 `L2.00081`; 25 `L2.00082`; 26 `L2.00083`; 27 `L2.00084`; all other ids reach the blank-effect default `L2.00085`. Effect objects come from the Effects string through `L2.00086`/`L2.00087` (key before the equals sign compared exactly against a 50-name table at `L2.00088`) or are blank (`L2.00089`).

Magnitude at effect+0x40 (P is the power byte; integer division) and duration at +0x42:

| Rule | Spells | Magnitude | Duration |
|---|---|---|---|
| protection | 4, 8, 13, 19 | P/2 | `ftol(15 * 1.025^P * 16)` |
| shield | 27 | P/10 + 3 | Duration value 20 in the 1.025 law |
| speed | 21 Haste, 29 Slow | P/15 + 1, Slow negated | Duration value 10 in the 1.025 law |
| to-hit | 20 Bless, 28 Curse | 4P/5 + 20 | Duration value 10 in the 1.025 law |
| light | 15 Light | P/30 + 1 | area, 15*16 + (16P)/10 |
| darkness | 14 Darkness | -1 - P/30 | area, same |
| poison cloud | 6 | `ftol(-2 * (P/45.0 + 1.0))` (float division) on the health kind | Duration value 5 in the 1.025 law; area 10*16 + (16P)/10 |
| invisibility | 12 | blank effect | `P << 4`; 0 deletes the effect |

Wall of Earth (17) and the damage spells use a blank effect; damage spells carry their numbers in the damage block of `R2-ENGINE-028`. The Bless and Curse arm writes the same magnitude for both and does not apply a sign; the sign appears in the resolver (`R2-ENGINE-028`). Duration words are 16 bits (`R2-ENGINE-032`).

ROM1 has the same arm-per-rule structure, 28 spells in 17 rules (`MAGIC-ARM-014`, `MAGIC-EFFECT-015`). The protection, Shield, Haste, Slow, Bless and Curse magnitudes equal ROM1's (`MAGIC-CEIL-013`). The duration law is ROM1's 1.025 (`R2-ENGINE-022`). ROM2 differs in Invisibility and in Poison Cloud, which ROM1 computes as `ftol(effect+0x40 * f)` with `f = P/30 + 1`.

**Confidence.** High for the arm addresses and the integer expressions. Medium for the Poison Cloud scale (the template value -2 was read, its consumer was not) and for the unsigned Bless and Curse arm (the sign handling is inferred from the resolver).

### R2-ENGINE-025

Stone Curse (id 18, arm `L2.00076`): the duration is a uniform draw `0..D` (`R2.0021`), made only inside the target-test branch (`vtable+0x2c`, `L2.00090`; without it the duration is the fixed `D`), with `D = ((P*240*m)/100) * (100 - EarthResist)/100`, where EarthResist is the word at +0xca and `m` is 2 for a spell cast and 1 for an item cast (so D is `4.8*P` for a spell cast before resistance); a draw of 0 deletes the effect. The fill's own duration for id 18 is overwritten. The stored Effects string begins with a literal quote character before `absorbtion`. The kind parser compares the key exactly (`R2-ENGINE-024`), so the key is not `absorbtion`; the resulting kind is 0 and the stated magnitudes (absorption +5, defence -20) are probably not applied. The comparison routine `L2.00091` was not read.

ROM1 gives Stone Curse a `pow(1.025, P)` duration (`MAGIC-CEIL-013`).

**Confidence.** Medium: the random draw and the clamp factor were read; the kind-0 outcome is inferred from the unread comparator.

**Unknown.** The runtime consumer of the effect; whether kind 0 is a silent no-op or a default.

### R2-ENGINE-026

- Heal (24, `L2.00081`): amount `base + U[1, spread]` (`L2.00092` draws 1..n; a spread of 0 still gives 1), capped at maximum minus current health; refused when target health is -10 or lower or +0x98 is zero.
- Drain Life (26, `L2.00083`): `v = (base + U[1, spread]) * (100 - target word +0xcc)/100`, capped at target health plus 10; target health falls by `v` and the caster's rises by `v`. 
- Summon (25, `L2.00082`): level = `clamp(Astral skill word +0xb2 / 25 + 1, 1, 4)`; the creature kind is a roll 1..3 (`L2.00092`); the creature is built from name strings, placed at the target cell and flagged.
- Control Spirit (22, `L2.00079`): finds a creature at the target cell with kind byte +0x13c in 2..4 and sprite word +0xe in 0x52..0x62; kills it (health set to -10001, kind 5) and builds a replacement from a template with halved Body and Reaction per kind, owned by the caster's owner. The arm does not read P.
- Teleport (23, `L2.00080`): moves the caster to the target cell through `L2.00093`; range is Max Range + P/3 (`R2-ENGINE-022`).

ROM1's unique-spell arms are listed in `MAGIC-SING-019`; they were not compared spell by spell here. Teleport's range term equals ROM1's (`MAGIC-CEIL-013`).

**Confidence.** High for Heal, Drain Life, Summon level and Teleport range (full arms read). Medium for Control Spirit: the template construction and the owner transfer were read once, not cross-checked, and the kind and sprite filters were not mapped to named creatures.

**Unknown.** The Control Spirit template contents; the Summon creature names and statistics.

### R2-ENGINE-027

Prismatic Spray (id 11) bypasses the normal arm: refill `L2.00032`, then `L2.00033`. The victim count is `min(P/20 + 2, 7)`; P is the power byte of `R2-ENGINE-020` for a book cast and `castSpell effect +0x42 / 20 + 2` for an item cast. The routine passes that cap and the spell's range byte (+9) to the list builder `L2.00094` and applies `L2.00034` to each victim. The cap reaches 7 at P = 100. It equals ROM1's (`MAGIC-CEIL-013`).

ROM1: the capped argument and ranked victim selector are in `MAGIC-SPRAY-134`; the ROM2 selector and its ranking were not read.

**Confidence.** High for the cap formula and the call structure. The victim ranking inside `L2.00094` is not read here.

## Resolution of a magic hit

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-028 | A magic hit resolves as `base + U[0, spread]`, cut by the matching elemental resistance as `ftol(v*(100-p)/100 + 0.75)`; spells never roll to hit and absorption does not touch the magic part. | High / Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |
| R2-ENGINE-029 | Each elemental resistance word is Spirit/2 plus effect modifiers, clamped to 0..min(Spirit/2 + 70, 100), in the same derive order as ROM1. | High / Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |

### R2-ENGINE-028

The apply routine builds a damage block (0x60 bytes) for every spell with `base + spread > 0` except Heal and Drain Life: +0x5b base byte, +0x5c spread byte, +0x5d kind. The kind equals the Sphere column through the setters `L2.00095`, `L2.00096`, `L2.00097`, `L2.00098` and `L2.00099` (kinds 1 to 5), the icon id is `2*id + 9`, and the sphere switch is the table at `L2.00100`.

The resolver `L2.00101` (block fed through vtable +0x4c) computes the magic component as block +0x13 base plus `U[0, +0x14]`. Spells always apply: the component is admitted when the physical swing hit or the block has no physical damage. The component is cut by the resistance word of kind +0x15: kind 1 word +0xc4, 2 +0xc6, 3 +0xc8, 4 +0xca, 5 +0xcc, giving `ftol(v*(100-p)/100 + 0.75)` clamped at 0. Absorption (+0xc0) applies to the physical part only; armour is chosen by damage type from the byte array at +0xce. Bless and Curse (effect ids 20 and 28) are tested on the attacker's physical damage: with probability `magnitude/101` (a draw 0..100 below the magnitude) Bless uses the swing's base plus spread and Curse uses its base only.

Worked example, Fire Ball at P = 70 (`evidence/hit-example.tsv`): base 23, spread 20, damage 23..43; against Fire resistance 0, 25, 50, 75 and 100 the resolved damage is 23..43, 18..33, 12..22, 6..11 and 0..0.

ROM1 resolves a spell the same way, as read here against `MAGIC-RESIST-006` (same component order and the same +0.75 term; straight percentage, elemental component only, no hit roll); Protection spells add to the resistance words through the effect dispatch `L2.00054` (protection words +0x102 to +0x10a); the fold of those into +0xc4 to +0xcc passes through `L2.00102`, which was not read.

**Confidence.** High for the damage block, the kind map, the resistance arithmetic and the absorption split. Medium for Bless and Curse in the resolver: the meaning was read in this lane and the ROM1 counterpart was not re-read.

**Unknown.** The fold `L2.00102`; how a protection spell's magnitude (`R2-ENGINE-024`) reaches the resistance word in the resolver path.

### R2-ENGINE-029

The derive writes `prot[i] = Spirit/2` for the five words +0xc4 to +0xcc, folds the effect protections through `L2.00102`, and clamps with `clamp(min(Spirit/2 + 70, prot), 0, 100)`. The order (set, fold, clamp) matches ROM1 (`HERO-RESIST-012`, `MAGIC-SPIRIT-011`).

For Spirit up to 52 the effective ceiling is 96.

**Confidence.** High for the order, the Spirit/2 base and the clamp instructions (read after the fold call). The effect of the fold on the value is Unknown.

**Unknown.** The fold `L2.00102`.

## Mana regeneration

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-030 | Mana regenerates in the actor tick every tick while below maximum: `max*(modifier+100)*K/50` hundredths per tick, K 3 when more than 80 ticks have passed since actor+0x138 (writer unread) and 1 otherwise; the block has no cast condition. | High / Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |
| R2-ENGINE-031 | Maximum mana is `2*Spirit` plus an experience term, then scaled by `pow(1.1, Spirit)/100 + 1`; health regenerates every fourth tick per actor phase. | Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |
| R2-ENGINE-032 | Duration words are 16 bits and the arithmetic wraps from P 228 (Protection), 216 (Shield) and 244 (Bless, Curse, Haste, Slow); Poison Cloud never does, and no damage byte wraps. | High / Medium | ● active | [EXP-2006](../experiments/EXP-2006-rom2-magic/) |

### R2-ENGINE-030

The block `R2.0022` in the actor tick `L2.00028` runs for an actor with positive health whose state word +0x54 is not 0x10 (`L2.00103`, `L2.00104`), when mana is below maximum, every tick. The accumulator is `mana*100 + remainder(+0xa3) + max*(modifier + 100)*K / pace`, where `modifier` is the effect byte +0xe2, `pace` is the word +0x9e (50: set in the constructor `L2.00105`, copied at `L2.00106`) and `K` is 3 when `(session tick + 4) - actor+0x138` exceeds 0x50, else 1. The remainder byte is stored first, then the quotient word, capped at maximum. With modifier 0 and K 1 the gain is `2*max` hundredths of a point per tick, so an empty pool fills in 50 ticks (17 ticks at K 3; `evidence/regen.tsv`).

Health uses the same shape with `2*max*(+0xde + 100)*K / +0x98`, every fourth tick on a per-actor phase `(tick & 3) == ((actor+4) & 3)`; ROM1 gates health on `tick % 4 == 0`. Dying actors lose 1 health per 4 ticks. The ROM1 shape is `HERO-REGEN-021` and `MAGIC-SPIRIT-011`.

Spirit scales the pool only through the maximum (`R2-ENGINE-031`); the block reads no Mind word.

**Confidence.** High for the formula and the gate (no cast or combat condition on the mana block, so a cast does not stop the refill). Medium for K: the timestamp +0x138 was not traced to a writer.

**Unknown.** What writes actor+0x138, and so whether casting resets the K timestamp ("quiet" is ROM1's meaning in `HERO-REGEN-021`, not ROM2 evidence); the tick-to-second scale (`SESS-TICK-006` defines the sub-tick).

### R2-ENGINE-031

Maximum mana is `m = 2*Spirit`; then `m = ftol(m + log1.1(XP/5000 + 1) * (non-mage ? 1 : 2))` (constant 5000.0 at `L2.00107`, `L2.00108` is log base 1.1); then `m = ftol(m * (pow(1.1, Spirit)/100 + 1))` (`L2.00109` is pow(1.1, x)). The multiplier on the experience term is 2 for a mage and 1 otherwise, the same arrangement as ROM1's inverted fighter multiplier note in `HERO-MP-006`. The order of truncations was read as listed; the second truncation was not confirmed against the stored word.

The health regeneration phase is stated in `R2-ENGINE-030`.

**Confidence.** Medium: the formula reproduces `HERO-MP-006` in shape and constants but the final store was not traced.

### R2-ENGINE-032

Duration words are 16-bit. The raw value `ftol(Duration * 1.025^P * 16)` first exceeds 65535 at P 228 for the four Protections (Duration value 15), at P 216 for Shield (20) and at P 244 for Bless, Curse, Haste and Slow (10); Poison Cloud (5) reaches 43418 at P 255 and does not wrap (`evidence/power-domain.tsv`). Past those thresholds the stored word is the low 16 bits (Protection at P 255 stores 64719 from a raw 130255). Inside the normal domain the ceiling is P 122 without effects and 170 with Mind 100 (`R2-ENGINE-020`): the wrap needs skill plus Mind of 246 (Shield), 258 (Protection) or 274 (to-hit and speed), that is a school skill bonus of at least 46, 58 or 74 over the P 170 ceiling. Damage bytes cannot wrap in P of 0..255: the largest base is 95 and spread 190 (Blizzard at 255), below 256. The tool flags rows where the floating result lies within 1e-9 of an integer.

ROM1 caps P at 100 with the same base 1.025 (`MAGIC-POWER-004`, `MAGIC-CEIL-013`), so there is no wrap there; ROM2's wider P range is what makes a wrap possible at all.

**Confidence.** High for the 16-bit arithmetic on the P range `R2-ENGINE-020` admits. Medium that the word is stored without a later saturating clamp: the stores read are word writes, and the effect timer consumer was not read.

**Unknown.** The effect timer's own handling of a wrapped word.

## Single-mission load

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-033 | RU build: a campaign mission is chosen by the number `scenario.dll` returns; `allods2.exe` formats a numbered map name, prefixes `Scenario\` and parses the records with dispatcher `R2.0002`; other modes pick a loose map from a listing. | High / Medium | ● active | [EXP-2007](../experiments/EXP-2007-rom2-mission/) |
| R2-ENGINE-034 | RU build of `allods2.exe`: opens four secondary archives, loads 13 text tables and data.bin, then loads `scenario.dll` and binds ordinals 1 to 3 and 5 to 19; the DLL is required. No `templates` string in this build. | High | ● active | [EXP-2007](../experiments/EXP-2007-rom2-mission/) |
| R2-ENGINE-035 | RU build: mission logic runs from the ALM records and `allods2.exe`; `scenario.dll` supplies the campaign location, a numbered variable store and town, inn and shop flows; script arms for scenario variables and subobjectives call it. | Medium | ● active | [EXP-2007](../experiments/EXP-2007-rom2-mission/) |
| R2-ENGINE-036 | ROM1's type 7 grammar parses all 46 campaign maps; the vocabulary adds instants 35 to 39 and checks 23 to 27, used by 35 maps; every map also carries one undeclared opcode word (65538). | High / Medium | ● active | [EXP-2007](../experiments/EXP-2007-rom2-mission/) |
| R2-ENGINE-037 | Each campaign map has one `main.res` file `text/mission<N>.txt`, loaded in campaign mode; briefing, subobjective, failure and event are section families across the corpus; ROM1's `text/battle/m<N>/` family is absent. | High / Medium | ● active (amended) | [EXP-2007](../experiments/EXP-2007-rom2-mission/), [EXP-2009](../experiments/EXP-2009-rom2-text/) |
| R2-ENGINE-038 | RU build: unit placement resolves the record's definition id against the Units column `serverID` (slot 55) or, for the person arm, the Humans column `serverID` (slot 24); ROM1 keyed Units on other columns. | Medium | ● active | [EXP-2007](../experiments/EXP-2007-rom2-mission/) |
| R2-ENGINE-039 | ROM1's byte-run (`.256`) and the word-run (`.16a`) frame grammars parse every frame of every ROM2 sprite file: 40,570 and 2,865 frames, no failure; pixel meaning is not tested. | High / Medium | ● active | [EXP-2007](../experiments/EXP-2007-rom2-mission/) |
| R2-ENGINE-040 | Among the 46 campaign maps the axes disagree: map 83 has the fewest triggers and declared operations and no ROM2-only operation; map 32 has the fewest units and objects; the ten-map Pareto set is the same with or without opcode 65538. | High | ● active | [EXP-2007](../experiments/EXP-2007-rom2-mission/) |

### R2-ENGINE-033

Observation, RU build of `allods2.exe` (the EN build is a different image and was not decompiled): the mission-enter routine `R2.0023` takes the mission number from the `ScenarioGetCurrentLocation` pointer (`L2.00110`, read in `L2.00111`), formats the map name `<number>.alm` and calls the map load `L2.00112`, which calls `L2.00113`. That routine prefixes `Scenario\` when `this+0x20` is nonzero and calls the record dispatcher `R2.0002` (`R2-ASSET-021`); its refusals are file not found, not a map file, wrong block number, version too new, Tiles block not found and Altitudes block not found. The stages after parsing are players (`R2.0024`), world allocation, `L2.00114`, unit placement (`L2.00115`), script compile (`L2.00116`) and loot and casters (`L2.00117`); `L2.00118` runs only with the multiplayer flag set. Campaign mode (`L2.00119`) sets mode 2; other modes choose from a `*.alm` listing (`L2.00120`).

Both roots hold the 46 campaign maps as `scenario.res` nodes named by number, plus 37 (RU) or 13 (EN) loose root maps. The 13 EN root maps are byte-identical to RU's, and 42 of the 46 `scenario.res` maps are identical (maps 110, 42, 52 and 74 differ). `scenario.res` holds no `Scenario.reg` or `GlobalMap.reg`, which ROM1's holds.

**Confidence.** High for the call chain and the name format on the RU build; the EN build is not covered. Medium that `Scenario\` resolves into `scenario.res` and not a directory: the prefix routine was read, the archive lookup was not.

**Unknown.** The EN executable was not decompiled; its mission entry is assumed to match from its path strings (one differs) and is not shown. The body of `ScenarioGetCurrentLocation` in `scenario.dll`.

### R2-ENGINE-034

On the RU build the startup routine `L2.00121` opens `sfx.res`, `movies.res`, `scenario.res` and `speech.res` (`L2.00122`), reads `update.lst`, loads 13 `main\text\*.txt` tables and `patch\patch.txt`, reads `famehall.dat`, runs the data.bin loader `R2.0006` (path `World\Data\`), then calls `LoadLibraryA("scenario.dll")` and stops with "Can't find scenario.dll" on failure. It binds by ordinal 1, 2, 3 and 5 to 19 to globals: GetVar, SetVar, TalkTo, EnterLocation, LeaveLocation, EnterShop, LeaveShop, EnterInn, LeaveInn, NewGame, Save, Load, GetAvailableLocations, GetShopAssortment, IsTownAvailable, IsMissionAvailable, GetCurrentLocation, GetAllLocations. IsMissionAvailable has no reader in `allods2.exe`.

The path strings of `allods2.exe` that resolve against the root's own archives are 481 of 595 (RU) and 481 of 596 (EN), against 405 of 548 for ROM1's `rom.exe`. 201 ROM2 path strings have no ROM1 counterpart, 76 of them under `graphics/interface` (`evidence/exe-paths-diff.tsv`). The string `templates` appears in no string of `allods2.exe`; the string is in the ROM2 Map Editor, whose `templates.bin` reference `R2-ASSET-025` records, so the absence is bounded to `allods2.exe`.

**Confidence.** High. The RU and EN executables are different builds (2,208,768 bytes with a sixth section `.st1073`, against 2,157,568 bytes with five); only RU was read.

**Unknown.** Whether the EN build binds the same ordinals; ordinal 4 is not bound in the RU build.

### R2-ENGINE-035

Alternatives D1 (records only), D2 (the DLL drives the mission) and D3 (both, the DLL limited to campaign state) were posed. The reads support D3, on the RU build only. The 18 exports of `scenario.dll` carry no game strings; its imports include profile-string, file and resource functions. In `allods2.exe` the instant arms `L2.00123` (Set Scenario Variable) and `L2.00124` (Set Subobjective, a state machine over variable `n+0x2f0`) and the check arms `L2.00125` and `L2.00126` call GetVar and SetVar; results go to a session register array at `+0xc784` (ROM1: `+0xbd34`). Town, inn and shop screens also call GetVar (`druidinnkeeper%d` and `druidshopkeeper%d` from variable `0x300`). Win and lose (`L2.00127`) keep the ROM1 shape.

**Confidence.** Medium: the DLL body was not decompiled, so what it does inside each export is inferred from call sites and the absence of content strings.

**Unknown.** The DLL's variable store layout and its save record (`R2-SESSION-017`).

### R2-ENGINE-036

Input: 46 campaign maps per root, type 7 payloads (`tools/r2mission`, `evidence/map-census.tsv`). The grammar of `ALM-TRIG-044` (counted action nodes of 796 bytes, check nodes of 796 bytes, triggers of 184 bytes, opcode at node `+0x40`) consumes every payload exactly. Against the ROM1 EN vocabulary files the ROM2 files add instants 35 (Set Scenario Variable), 36 (Set Subobjective), 37 (Set music node), 38 (Remove item from everybody's inventory), 39 (Suspend group) and checks 23 (Get Scenario Variable), 24 (Get Subobjective), 25 (Spell on Tile), 26 (Spell on Unit), 27 (Centered on Tile): 49 against 44 instant groups and 27 against 22 check groups (`evidence/vocab.tsv`). 35 of the 46 maps use at least one of these (RU and EN alike), 32 of them instant 36. Every campaign map has a win action and 34 have a lose action; the root maps have neither.

Undeclared opcode: the value 65538 (0x10002) is in no vocabulary file (instants stop at 49, checks at 27). It occurs once as an action on every one of the 46 maps and on 164 check nodes spread over all 46 (`undeclaredActionNodes` and `undeclaredCheckNodes` in `evidence/map-census.tsv`). The census counts it as a distinct operation; it is neither ROM1-only nor ROM2-only in the classification. ROM1's 38 maps carry none (`ALM-TRIG-045`). The exact-consumption walk shows the counted-array layout holds; it does not show what these nodes do.

**Confidence.** High for the census, the vocabulary difference and the parse. Medium for the nodes with opcode 65538, whose meaning is unread. The meaning of the signatures comes from the vocabulary files; the arms of instants 35 and 36 and checks 23 and 24 were read in code on the RU build (`R2-ENGINE-035`), the others not.

**Unknown.** What opcode 65538 does; the code arms of instants 37 to 39 and checks 25 to 27; every arm on the EN build.

### R2-ENGINE-037

`main.res` holds `text/mission<N>.txt` for 46 of 46 campaign map numbers in both roots, with no orphan file (`evidence/mission-text-summary.tsv`). Sections are tagged `#briefing`, `#subobjectiveN`, `#failureN` and `#eventN` (`evidence/mission-text-tags.tsv`). In the mission-enter routine of the RU build, mode 2 loads `main\text\mission%d.txt` and parses `#briefing` (`R2.0025`, `L2.00128`); other modes load `main\text\quest.txt`. ROM1 instead keeps `text/battle/m<N>/` files (briefing, briefmap, event, tips, title; `evidence/rom1-text-battle-families.tsv`).

**Confidence.** High for the file census in both roots; Medium for the load order (RU build only), because the event, failure and subobjective consumers were not read.

**Unknown.** Consumers outside the campaign and dialogue paths bounded by R2-ENGINE-049 through R2-ENGINE-052.

**Amended.** R2-ENGINE-049 through R2-ENGINE-052 read campaign caches, UI event selection, dialogue pages and font conversion in both clients. The section families are a corpus vocabulary, not a claim that every file has each family: mission 10 has no subobjective section. The headline's per-file section assertion is narrowed in the [correction record](retracted.md#rom2-mission-section-headline-narrowing). The plain briefing reader is `R2.0025`; `L2.00128` is a separate numbered briefing/description helper.

### R2-ENGINE-038

In the RU build `L2.00115`, a bit `0x10` of the placement record's in-memory flags selects an arm. The ordinary arm scans the Units collection from the top for the entry whose parameter `0x37` equals the record's `+0x10` definition id and builds a 0x208-byte unit (`L2.00129`); the person arm scans Humans for parameter `0x18` and builds a 0x254-byte unit (`L2.00130`). In `databin-slots-rom2.csv` both slots are the column `serverID`. Error text: "Can't resolve player %d for unit %d", "Invalid unit %d during loading", "Can't place unit %d during loading". The building loader `L2.00131` treats the kind byte `+0x40` as a 1-based index bounded by the collection count; the table behind it is not named in the executable.

**Confidence.** Medium, RU build only: the scan and the slot numbers were read; the identity of the Units and Humans collections with the data.bin groups of `R2-ASSET-009` rests on slot numbering, not on a load-time link.

**Unknown.** Whether every `serverID` used in the maps is present in the table (no join was run); the consumers of the type 10 to 12 record arrays; whether `L2.00132` (hero required) is reachable in single-player.

### R2-ENGINE-039

`tools/r2mission` replays the frame container of ROM1's sprite formats and both pixel grammars over every frame. ROM2 RU and EN each give 1,921 `.256` files (40,570 frames) and 623 `.16a` files (2,865 frames) with no frame-walk failure, no `.256` control with both high bits set, and no `.16a` literal word outside the mask. ROM1 gives 1,384 files (23,895 frames) and 542 files (2,727 frames) the same way. The files come from `graphics.res`, `main.res`, `patch.res`, `scenario.res`, `world.res` and `cs.res` of each root; 8 `.256` payloads are empty in every root.

**Confidence.** High that the container and run structure parse; Medium that pixel meaning, palettes and fonts hold, since no frame was rendered and no palette or font container was compared.

**Unknown.** The palette and font containers.

### R2-ENGINE-040

Input: the 46 `scenario.res` maps of each root (`evidence/map-census.tsv`, `evidence/pareto-rom2-ru.tsv`). The maps not dominated on triggers, distinct script operations, ROM2-only operations, units, objects, unit class values and cells are 94, 96, 85, 87, 83, 81, 82, 84, 43 and 32 in both roots. Map 83 has 2 triggers, 3 declared operations (5 with opcode 65538) and no ROM2-only operation, on 144 by 144 cells with 2 players, 21 objects and 92 units; it has a win action and no lose action. Map 32 has the fewest units (20) and objects (2) on 80 by 80 cells with 3 players, 6 triggers and 12 distinct operations. Recomputed with opcode 65538 left out of the operation count (`evidence/pareto-declared-rom2-ru.tsv`, EN likewise), the non-dominated set is the same ten maps and map 83 stays the unique minimum on triggers and on operations. That the one action and one check of map 83 carrying 65538 need nothing outside ROM1's vocabulary is not shown. No single map is smallest on every axis (alternative G2).

**Confidence.** High for the census. The unit class value axis reads the type 6 field `+0x08` under a ROM1 reading that was not checked in ROM2.

**Unknown.** Which map plays shortest; the census measures records and script features, not play time.

## ROM2 campaign script contracts

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-041 | Both preserved ROM2 campaign corpora have 46 exactly parsed type 7 payloads; new instants/checks retain separate full-word operation and ID namespaces. | High | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-042 | In the verified RU/EN compilers, action word 65538 appends a packed low-byte parameter pair to a compiler collection instead of creating an executable action. | High | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-043 | In the verified RU/EN compilers, check word 65538 initializes its map result-register slot from parameter 0 instead of creating an executable check. | High | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-044 | ROM2 literal checks pack ordinary scalar operands densely and resolve unit/group/player references separately; unresolved literal references can omit the compiled node. | High / Medium | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-045 | Instants 35 and checks 23/24 access the scenario.dll DWORD bank; its exports have no explicit index guard, and NewGame initially clears a1024-DWORD span. | High | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-046 | In both clients, instant 36 mutates752+n with exact mode-dependent stores and 253/254 notifications; objective rows read the resulting bits. | High | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-047 | ROM2 instant 2 emits packet 0xb6; actual RU/EN receiver tables forward its event ID to UI 433, which separates ordinary events from 250/253/254/255. | High | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-048 | The corresponding scenario.dll code clears slots 752..767 on EnterLocation; the verified LeaveLocation mission10 arm marks completion and adds location20. | High | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-053 | ROM2 check 25 tests a terrain cell's indexed effect-layer pointer after a low-u16 X+256Y lookup; the arm has no explicit layer bound. | High / Medium | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-054 | ROM2 check 26 tests the resolved unit's active-effect DWORD at+0x144 using scalar 0 as an x86 shift count; registration/expiry set and clear those bits. | High | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-055 | ROM2 check 27 compares unit tile X/Y with compiled scalars 1/2 and requires fractional bytes 128/128; literal catalog X/Y instead pack into scalars 0/1. | High | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-056 | ROM2 instant 38 makes one matching inventory extraction call per unit in the terrain-provided list, destroys a returned item and refreshes that unit. | High / Medium | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-057 | ROM2 instant 39 clears a group's forced-activity DWORD; actor-proximity marking and ordinary group/unit ticks determine its later activity. | High / Medium | ● active (amended) | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-058 | In both verified ROM2 executors, declared instant 37 takes Bad instant default, optionally sends diagnostic packet 0x91 and returns normally; no campaign map contains it. | High | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-059 | Both ROM2 trigger runners compare map result registers with signed comparisons, short-circuit their conjunction, then set a pass latch and run actions in stored order. | High | ● active (amended) | [EXP-2008](../experiments/EXP-2008-rom2-script/) |
| R2-ENGINE-060 | ROM2 result actions increment the session win counter or store a failure reason; the verified RU result tick converts those fields into0xb5/0xb4 packets. | High | ● active | [EXP-2008](../experiments/EXP-2008-rom2-script/) |

### R2-ENGINE-041

Input is all 46 ALM members of scenario.res independently read from each preserved root. `evidence/inputs.tsv` records map/input hashes; `nodes.tsv`, `parameters.tsv`, `triggers.tsv` and `counts.tsv` retain numerical measurements. The record walk consumes the type 0 actual660-byte payload despite declared644, then each declared record extent. Type 7 has counted actions 796 bytes, checks 796 bytes and triggers184 bytes, with exact payload EOF. Per root, there are 1334 actions and 986 checks. Node word 0x40 is the full DWORD operation;0x44 is its ID,0x48 is a separate stored word,0x4c is the ten-DWORD parameter array and 0x74 the ten-DWORD type array. Action and check IDs never share the scanner's namespace.

`opcounts.tsv` measures instants 35/36/37/38/39 as 4/140/0/7/6 nodes on 3/32/0/4/1 maps; checks 23/24/25/26/27 as 5/3/24/3/0 nodes on 4/2/3/1/0 maps. Word 65538 occurs in 46 action and 164 check nodes. `declarations.tsv` reads catalog keys case-insensitively and uses ParN_Value, including mixed-case unit declaration keys for 26/27. Catalog types specify operands, not executable semantics. The exact identity/geometry of each independently parsed image is in `pe.tsv`.

**Confidence.** High for this complete two-root corpus and grammar. The scanner sorts the exact46 members and rejects overflowing counts or residual bytes. The negative37/27 clauses count full words on every node of that population, independently of Ghidra function recovery.

**Unknown.** Noncampaign maps, other installations, malformed payload acceptance and arbitrary operation domains are outside this census.

### R2-ENGINE-042

Compiler entries are RU L2.00116 and EN L2.00201. The action branch compares the entire node operation with 0x10002. Equality takes a separate collection path; it does not dispatch the low16 value 2. It combines low8(parameter 0) with low8(parameter 1)<<8, appending the u16 to compiler-this+0x28. Selected raw sites are RU L2.00439 and EN L2.00440 in `instructions.tsv`. RU node loader R2.0002 moves stored parameters0x4c/types0x74 to runtime0x48/0x70. The EN compiler independently uses the latter layout; its loader permutation was not separately located.

The mission10 action ID 1 has parameter 11/type 5 and 41/type 6, yielding 0x290b. Both complete maps have the same hash recorded by 041. The ordinary executable action population excludes this full-word special action.

**Confidence.** High for the full-word branch, exact byte packing and append in each client, using raw compiler sites. This excludes ordinary instant 2 execution and a scenario-variable write as the special action's direct effect.

**Unknown.** The downstream consumer and semantic role of the collection, and EN's node-loader permutation, remain unverified. A packed coordinate-looking pair does not identify a spawn or camera effect.

### R2-ENGINE-043

RU compiler L2.00116 and EN L2.00201 assign original check IDs to compiled result slots. The check 65538 branch copies runtime parameter 0 into Session+0xc784+4*slot, advances the slot, and creates no ordinary executable check. Selected sites RU L2.00441 and EN L2.00442 retain the full-word comparison and store. Ordinary compiled check IR has result slot at+0, operation at+4 and dense scalar base+8. Runtime+0x318 retains the stored node word 0x48 in compiled metadata; its full meaning is not supplied by that copy.

Mission10's special check IDs 2/4/6/7/8 initialize0/4/1/2/3; ID 21 initializes1229. These are map result registers consumed by the trigger runner, not scenario.dll variables. They coexist with executable check results in the same register namespace.

**Confidence.** High for both actual compiler branches and the exactly parsed mission10 values. Raw stores distinguish constants from executing ordinary check 2 and from calling DLL GetVar/SetVar.

**Unknown.** General result-register allocation bounds, the stored metadata word's full role and malformed/high operation words beyond the observed equality remain unverified.

### R2-ENGINE-044

The literal check compiler loops ten stored parameters. Nonreference values append densely to compiled scalars at IR+8+4*denseIndex, incrementing denseIndex independently of the source slot. Types2/3/4 resolve group/player/unit references into separate fields: first group IR+0x60, first player+0x64, first unit+0x5c and a second reference+0x68. Type 8 stores low16(value+0xe18) at+0x6c. Selected dense-copy sites are RU L2.00443/EN L2.00444. Unused scalar zeros are still walked.

The added arms' literal packing is:35/36 scalar 0,scalar 1;37 scalar 0..7;38 item u16+0x6c;39 group pointer+0x60;23/24 scalar 0;25 scalar 0..2;26 unit+0x5c and scalar 0;27 unit+0x5c and scalar 0/1. All ordinary scalars are DWORDs. Successful compiler group/unit references can set the referenced group's record+0x48 forced-activity flag, with the observed check 1 exception. Unresolved group/unit/player/building references set a compile-success flag false; the failure branch clears the ID mapping and omits the executable IR while reporting the failed node.

References in 8000..9999 take a different path: original parameter-slot positions, mode bytes and quotient/remainder buckets are retained for later dynamic resolution. These are not evidence that every literal scalar is a map-register reference.

**Confidence.** High for literal scalar/reference packing and the positive omission branch, independently read in both compilers. Medium for the complete dynamic-reference/error contract because some indirect lookup and dynamic execution edges remain frontier.

**Unknown.** Arbitrary invalid references, callbacks, dynamic source domains and malformed scalar admission remain unverified. Check 27's consequence of dense packing is stated separately in 055.

### R2-ENGINE-045

RU 35 at L2.00123 and EN 35 at L2.00445 call SetVar(scalar 0,scalar 1). Check 23 at L2.00125/L2.00446 calls GetVar(scalar 0) and writes its DWORD to the IR result slot in Session+0xc784. Check 24 at L2.00126/L2.00447 instead calls GetVar(752+scalar 0). Client pointers are RU L2.00448/L2.00449 and EN L2.00450/L2.00451; startup binds SetVar ordinal2 and GetVar ordinal1.

The DLL exports at D2.00001/D2.00002 address DWORD baseD2.00003+4*index. GetVar returns the loaded DWORD with RET4; SetVar stores its DWORD value with RET8. Neither export contains an explicit index guard. NewGameD2.00004 begins by zeroing0x1000 bytes at D2.00003, a contiguous1024-DWORD initialized span. It callsD2.00005 and later writes slot 768=10. This proves the initial clear and explicit later value, not that every slot's final value is zero. The bank is static data, not a heap allocation. The two DLL .text/.data hashes agree; separate PE replay corroborates correspondence without supplying two independent implementations.

**Confidence.** High for literal operands, export widths/calling convention, unguarded instruction bodies, initialized span and explicit slot 768 store. Selected sites and exports are regenerated from both DLLs and clients.

**Unknown.** Valid external indices, allocation/domain boundaries beyond the initialized span and all final constructor/NewGame defaults remain Unknown. The missing guard does not establish safe unbounded indices or default-zero arbitrary variables.

### R2-ENGINE-046

RU L2.00124 and EN L2.00452 read scalar 0=n and scalar 1=mode and use DLL slot 752+n. Mode 1 stores 1 only if old value is 0. Mode 2 does nothing if old value is exactly4; otherwise it emits event253 only when(old&2)==0 and stores 3. Mode 4 emits event254 and stores 5 without the old 4 guard. Every other mode directly stores mode. The producer arguments are null player, event ID, extra0. Mission10 action IDs 24/25/27/28 use(n,mode)=(1,1)/(1,2)/(2,1)/(2,2), targeting753/754.

Objective consumers RU L2.00453 and EN L2.00454 read GetVar(752+row) for cached row indices. Zero omits the row. The native rendering tests bits 2,4,1 in that priority and chooses sprite frames 11,12,10 respectively. Notification253/254 takes separate UI 433 banner branches. The examined banner setup and objective-render routes do not establish a SetVar acknowledgement transition. The decoded SetVar-reference population has 9 sites/5 owners per client: binding,35,36, options writes 776/781 and a unit packet path at 531+unit.u16(+0x1dc). This navigational population is not an alias/indirect-writer absence proof.

Known exact slot writers are generic35,36 and the DLL reset of 752..767 on EnterLocation. None of the four installed35 nodes targets752+n: their slots are 779/779/772/780. State 4 is not proved to mean acknowledged/completed merely because mode 2 guards it.

**Confidence.** High for exact mode branches, stores, event arguments and objective bit consumers in both clients, using selected instructions rather than numeric-state names.

**Unknown.** Any indirect acknowledgement writer, semantic names for states4/5, cache/text population beyond the independently traced objective reader and notification presentation timing remain unverified. The actual3 store suppresses a repeated253 while bit 2 remains set; mode 4 itself has no repeat suppression.

### R2-ENGINE-047

Ordinary instant 2 RU L2.00455/EN L2.00456 passes(0,scalar 0,0) to event producer RU L2.00457/EN L2.00458. The packet has player u16+7 (zero for the null/broadcast argument), opcode byte+9=0xb6, event DWORD+0xa and extra DWORD+0xe. RU packet bus L2.00459 and repaired-vtable target L2.00460 copy the payload to a player's queue; decode and actual client reception were followed through the packet object's virtual methods.

Actual RU receiver L2.00461 uses selector L2.00462 and target table L2.00463 after subtracting3 from opcode. Packet 0xb6 maps to L2.00464, which calls L2.00465(0x433,eventID,0). Independently, EN receiver L2.00466 uses selector L2.00467/table L2.00468 and maps 0xb6 to L2.00469, calling L2.00470 with the same UI arguments. `dispatch.tsv` measures those actual entries. The containing RU fragment L2.00471 supplied misleading decompiler case labels and is not the opcode authority.

UI 433 stores its selector at RU+0x454/EN+0x43c. Values250/253/254/255 take separate branches. Ordinary values with UI flag bit 8 clear format event%d and call RU R2.0037/EN R2.0038; otherwise their IDs enter the deferred queue. Mission10 contains ordinary instant 2 event IDs 1..10, including a second event5. The native packet interface, not catalog naming, connects those actions to the event dialog entry.

**Confidence.** High for the producer, repaired packet queue edge, raw receiver dispatch and UI selection in the stated images. Independent raw tables rule out the rival map-variable-only interpretation of ordinary event emission.

**Unknown.** Text pages, speaker selection, acknowledgement/button parsing and speech are outside this interface claim. Other packet routes and complete networking are not enumerated.

### R2-ENGINE-048

Corresponding DLL EnterLocation export5 at D2.00006 stores the location-record pointer at D2.00007 and clears16 DWORDs at D2.00008: bank slots 752..767. NewGame enters the first location record and explicitly stores slot 768=10. GetCurrentLocation export18 returns the record pointer; the client's location-number wrapper reads the record ID, so the export itself is not a numeric mission ID.

LeaveLocation export6 at D2.00009 initializes its output to-1, requires a nonnull current record, reads its type/ID through D2.00010/D2.00011, and selects the ordinary continuationD2.00012 unless type 2 takesD2.00013. Ordinary continuation resets20 transition-working slots, removes the current record from the available list, clears its pointer and writes 1 to completed-bank slot 896+locationID. Its actual case 10 callsD2.00014(20), which adds the matching location record to the available list, and stores output 1. That case does not test753/754; this bounded absence follows the entire selected case, not a corpus inference.

**Confidence.** High for export/reset boundaries and the exact mission10 branch in the corresponding DLL code. Selected instructions and the DLL section hashes identify the dependency independently of the client script result registers.

**Unknown.** Other continuation arms, registry/catalog persistence, full UI scheduling from a victory acknowledgement, failures, SAV and arbitrary location records remain frontier. A static continuation branch is not a completed native campaign-playthrough witness.

### R2-ENGINE-053

RU check L2.00472 and EN L2.00473 form X+256Y from compiled scalars 0/1, then pass its low u16 to terrain helper L2.00474/L2.00475. The helper requires bit 0x20 in terrain+0x10000+cell and an existing cell entry in the map atterrain+0x54084. It copies the 60-byte cell record to scratchterrain+0x5402c and returns scratch+0x14. The arm returns 1 iff DWORD[returned+4*scalar 2] is nonzero; otherwise 0. It has no explicit layer-index guard.

The RU removal consumer L2.00476 maps an effect object's u16 ID+0xc through L2.00477, clears a matching layer pointer atterrain+0x54040+4*layer, counts nonnull pointers over six slots and stores that count atscratch+2 before updating the cell entry. The observed ID-to-layer mapping is 3->0,6->2,17->3,15->4,14->5; other IDs take a diagnostic/default path. This identifies cell effect-layer occupancy, not a direct comparison of check 25 scalar 2 with a spell ID.

**Confidence.** High for the actual check/lookup reads in both clients and the literal six-pointer RU removal consumer. Medium for full effect-category identification: the mapping's default path and all producer domains were not completed on both builds.

**Unknown.** Layer 1's producer identity, other60-byte record fields, every effect ID, bounds/malformed coordinates and the default helper's complete error semantics remain unverified.

### R2-ENGINE-054

RU L2.00478/EN L2.00479 use the resolved unit pointer at IR+0x5c and dense scalar 0 at IR+8. They return1 iff(unit.DWORD(+0x144)&(1<<scalar 0)) is nonzero. Native x86 SHL uses the low5 count bits for this32-bit operand. The three installed nodes use effectID 12 for unit IDs 20/21/22 in map96.

Registration RU L2.00480/EN L2.00481 appends a timed effect to unit+0x20, reads its u16 ID at effect+0xc and ORs the corresponding bit into unit+0x144. Expiry RU L2.00482/EN L2.00483 decrements the duration word+0x42 and, when expiry removes that effect, clears the bit. Selected registration/expiry instruction sites show the same field and mask dataflow, not only a shared displacement.

**Confidence.** High for the actual unit effect bit test and positive registration/expiry links independently read in both clients. This excludes interpreting+0x144 as a campaign variable or GUI selection flag.

**Unknown.** Whole effect timing, duplicate-effect ownership, arbitrary IDs, every removal path and other mechanics represented by the bitset remain frontier.

### R2-ENGINE-055

RU L2.00484/EN L2.00485 use the resolved unit pointer atIR+0x5c and its position pointer atunit+0x10. The X/Y getters L2.00486/L2.00487 and L2.00488/L2.00489 return position bytes 0/1. The check compares X withIR+0x0c (scalar index 1), Y withIR+0x10 (scalar index 2), then requires position fractional bytes 4/5 equal128/128 via L2.00490/L2.00491. Coarse activity-position consumers use the same tile-byte getters, supplying a positive position identity link.

Both catalogs declare Target_Unit,X,Y. The literal compiler resolves the unit separately, putting X/Y densely atIR+8/+0xc (scalar indices0/1). It still walks unused trailing zeros. Thus the executor reads scalar 1/2, not the intuitive declared coordinate pair0/1. Both 46-map corpora contain zero check 27 nodes; no authored occurrence resolves this mismatch.

**Confidence.** High for the exact native predicate, scalar-index mismatch and bounded corpus absence. Raw comparisons and dense-copy instructions distinguish the rival intuitive X/Y interpretation.

**Unknown.** Intended authored coordinate convention, dynamic-reference cases, invalid position admission and a native runtime witness for an authored27 remain Unknown. The catalog label alone does not repair the observed IR reads.

### R2-ENGINE-056

RU L2.00492/EN L2.00493 iterate the unit list provided atterrain+0xa456c through Session+0xa50. For each iterated unit they call inventory extraction L2.00494/L2.00495 through unit+0x7c, passing the u16 atIR+0x6c. Literal Target_Item compilation adds 3608 and retains the low16 bits. There is one extraction call per iteration, not a repeated drain-until-empty loop.

The extraction finds one matching inventory item. Quantity word+0x42<=1 unlinks that item; larger quantity invokes its virtual+0x40 extraction callback. It subtracts returned quantity*weight from inventory+0x20. On a nonnull result the arm calls the returned object's virtual destructor+4 with 1, then refreshes the iterated unit through L2.00496/L2.00497. Missing items do not take that destructor/refresh branch.

**Confidence.** High for the literal operand, per-iteration call, quantity branch, accounting, destructor and refresh in both actual clients. Medium for the entire list population and stack callback quantity because those producers/callees remain unverified.

**Unknown.** The list's complete membership/filter, equipped-item scope, the split callback's exact returned quantity, arbitrary item references and all inventory side effects remain frontier. This is not authority for deleting every copy from every party inventory.

### R2-ENGINE-057

RU L2.00498/EN L2.00499 take the resolved group pointerIR+0x60, follow group+0x3c and clear its DWORD+0x48. The record constructor L2.00500/L2.00501 zeroes its 0x50-byte record and sets active byte+0x45=1, proving a native initial override0 independently of corpus zeros. Successful literal compiler references may set override+0x48=1.

The ordinary mission tick route passes the embedded activity object at Session.DWORD(+0xa50)+0x92ecc, where Session.DWORD(+0xa50) is the terrain pointer. The activity reset L2.00159/L2.00158 clears activity-object+0x400 for 0x1210 bytes, hence terrain+0x932cc on this route, and clears group active+0x45 only in rosters whose owner DWORD+0x2c is nonzero. Producer R2.0095/R2.0094 has the literal predicate actor.owner(+0x14)->DWORD(+0x2c)==0. For such actors it reads position tile X/Y, shifts each right3, and ORs owner.u16(+0x32) over the 5x5 coarse-cell neighborhood centered there; the native padded bitmap has byte stride0x88. The alternative owner branch can directly mark the actor's group active when actor.i16(+0x94)<actor.i16(+0x96). No GUI fog equality follows from this producer.

Recount R2.0098/R2.0097 resets activity-object counters+0x1614/+0x1624 and invokes the producers/markers. These counters are terrain+0x944e0/+0x944f0 on the observed tick route. Marker L2.00155/L2.00154 sets group active+0x45 and increments activity-object+0x1614 when the unit's cell mask is nonzero, copying its low word to unit+0x1a4. A zero cell mask with override+0x48 nonzero also marks active and increments activity-object+0x1614 and forced-count activity-object+0x1624. Zero mask and zero override do not mark it active.

Ordinary mission ticks L2.00502/L2.00503 recount when Session+0xbbe8==0. Group tick L2.00504/L2.00505 then skips ordinary group AI for inactive groups under the same gate. Unit tick L2.00506/L2.00507 additionally tests explicit-command objectunit+0x1c4 byte+9: an inactive group with command byte 0 sets action kind 27 and skips ordinary AI, while a nonzero explicit command continues processing. A nonzero Session gate bypasses this particular inactivity rule. Clearing override therefore permits groups outside active cells to become inactive on recount; groups inside active cells remain active, and explicit commands can continue. It is not an unconditional immediate suspension.

**Confidence.** High for the exact clear, constructor, mask producer, counters and positive tick branches independently read in RU/EN. Medium for player classification and complete producer population; field identities do not yet support a human/AI or GUI-fog equivalence.

**Unknown.** Semantic classification of owner+0x2c/u16+0x32, the compared actor+0x94/+0x96 fields, all activity producer exceptions, other Session gate modes and whole AI behavior remain frontier.

**Amended.** The bitmap and counter offsets are local to the embedded activity object, not to terrain. The [correction record](retracted.md#rom2-script-base-and-trigger-corrections) preserves the former terrain+0x400 statement. The observed tick route establishes terrain+0x92ecc as the object base; the clear, producer predicates, tick effects and remaining confidence boundaries stand.

### R2-ENGINE-058

The actual action tables atRU L2.00508/EN L2.00509 map operation 37 (entry 36 after subtract1) directly to default L2.00510/L2.00511. The arm formats Script: Bad instant with the operation number. It tests Session+0x170->DWORD; when nonzero it emits diagnostic packet 0x91 through L2.00512/L2.00513, then joins ordinary compiled-IR cleanup and RET4. It has no abort instruction or direct result-register/campaign-variable store in this selected path. The diagnostic producer places its report string in the ordinary packet, with zero player selector for the passed null argument. Actual receiver table entries0x91 reach L2.00514/L2.00515; the zero message selector takes the ordinary text/message storage call.

The catalog has 37 with X,Y and six integer operands. All 46 campaign maps per root have zero full-word 37 action nodes. This is the actual executor/default path on the verified builds, not absence of every music feature in the image.

**Confidence.** High for raw table equality, diagnostic/control flow and complete bounded corpus absence. The entry is measured directly, independent of decompiler switch coverage. Normal return excludes the rival aborting default within this arm; emitted diagnostics remain an observable side effect when the native guard is enabled.

**Unknown.** Other music paths, arbitrary37 runtime authorship, the guard's full configuration source and full downstream diagnostic presentation remain unverified.

### R2-ENGINE-059

RU L2.00516 and EN L2.00517 read the compiled check-result slots in Session+0xc784. Stored comparison codes0..5 implement signed DWORD ==,!=,>,<,>=,<=, short-circuiting the trigger conjunction. Both stored operands are check IDs mapped to result slots; the right operand is not an immediate scalar. A nonzero compiled trigger flag skips an already-set byte latch at Session+0xe6c4+triggerIndex. Otherwise the runner first clears that byte, evaluates conditions, and on a pass sets it to 1 unconditionally before executing action IDs in stored order. The flag controls the next scan's skip, not whether a passing trigger stores the latch.

Mission10 `triggers.tsv` retains all 16 trigger rows; every stored DWORD+0xb4 is 1. Trigger0 tests result(check 2)==result(check 2) and executes action IDs 2,27; trigger7 tests result(check 12)<=result(check 4) and executes 9,23,12,28; trigger10 tests result(check 18)<=result(check 8) and result(check 17)==result(check 6), then executes 16,15,20,24. Some passing triggers have no actions. These numeric dependencies are local check results/IDs, not scenario-variable indices.

**Confidence.** High for the independently read runners and exact stored sequence measurements. Raw signed branch instructions distinguish signed register comparisons from unsigned or truthiness-only evaluation.

**Unknown.** General trigger validity, every malformed comparison/action ID and scheduler timing outside the selected ordinary runner remain frontier.

**Amended.** The mission10 trigger7 left operand is original check ID 12 in both reproduced rows. The [correction record](retracted.md#rom2-script-base-and-trigger-corrections) preserves the former check10 statement. The signed comparison, latch and action-order contracts stand.

### R2-ENGINE-060

Instant 4 RU L2.00518/EN L2.00519 increments Session+0xbe0c. Instant 5 RU L2.00520/EN L2.00521 stores scalar 0 into Session+0xbe14. Instant 8 RU L2.00522/EN L2.00523 increments the map result register selected by scalar 0. These fields are separate from the scenario.dll bank.

The RU result consumer L2.00127 requires player byte+0x41 nonzero and DWORD+0x2c zero. In the live-player branch it prioritizes failure reason Session+0xbe14>=1, emits 0xb4 with that reason and updates player result byte+0x40; otherwise nonzero win counter+0xbe0c emits 0xb5 and sets player result1. The actual receiver maps 0xb4 toUI433 with selector255 and the reason, and 0xb5 toUI430. Mission10 contains the ordinary win action and failure reason 19, measured independently from their native arms.

**Confidence.** High for both action bodies and the stated RU result-consumer/receiver boundary. Failure priority and counter increment follow raw instructions, not a boolean-win convention inferred from labels.

**Unknown.** EN's ordinary mission-result tick was not independently followed. Full victory-window acknowledgement scheduling into LeaveLocation and other failure sources remain frontier;048 supplies the independently established DLL continuation only.

## Mission text consumers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-049 | Both ROM2 campaign loaders select mission text by scenario location; substring spans feed briefing, failure and contiguous subobjective caches, and mission 10 has no subobjective section. | Medium | ✔ promoted (branch candidate) | [EXP-2009](../experiments/EXP-2009-rom2-text/) |
| R2-ENGINE-050 | In both ROM2 clients, UI message `0x433` selects `event%d` for ordinary IDs or defers the ID under UI bit eight; IDs 250, 253, 254 and 255 take separate branches. | Medium | ✔ promoted (branch candidate, amended) | [EXP-2009](../experiments/EXP-2009-rom2-text/) |
| R2-ENGINE-051 | Both ROM2 dialogue consumers select ordered part/header alternatives, present text and optional portrait/speech, and advance on `0x470`; mission 10 selects no speech name through this page route. | Medium | ✔ promoted (branch candidate) | [EXP-2009](../experiments/EXP-2009-rom2-text/) |
| R2-ENGINE-052 | Both ROM2 mission loaders call `CharToOemA`; font1 uses selector-gated byte remapping before glyph/advance indexing, with Windows conversion results and glyph labels still Unknown. | Medium | ✔ promoted (branch candidate) | [EXP-2009](../experiments/EXP-2009-rom2-text/) |

### R2-ENGINE-049

Input: the pinned EN and RU clients and `main.res` resources in EXP-2009.
RU `R2.0023`, controller `+0x63c == 2`, and EN `R2.0026`,
controller `+0x5d8 == 2`, use scenario current location to format
`main\\text\\mission%d.txt`. Other mode uses `main\\text\\quest.txt`.
R2-ENGINE-052 describes the conversion before CString storage.

Plain briefing readers RU `R2.0025` / EN `R2.0027` find the first
`#briefing` substring and start at match plus eleven. Failure readers RU
`R2.0028` / EN `R2.0029` and subobjective readers RU
`R2.0030` / EN `R2.0031` find their formatted numbered keys and
start at match plus key length plus two. They cut at the next `#` or EOF;
missing keys yield empty strings. Searches are not line-anchored and do not
validate a numeric token boundary.

Failure caching starts at two. Missing two through four use text-table slot
`0x118 + number`; the first empty number above four stops caching.
Subobjective caching starts at zero and stops at the first empty result.
Mission 10 in each root has one briefing, two failure and eleven event tags,
and no subobjective tag. Its `failure2` and `failure5` yield four cached
failure entries including fallback entries three and four. Script objective
flags do not manufacture a missing source label.

At UI message `0x431`, the campaign branch gates the failure dialog on the
not-yet-shown field and campaign mode, and supplies cached entry `reason - 2`
to RU `R2.0032` / EN `R2.0033`. The briefing control constructors
RU `R2.0034` / EN `R2.0035` receive the cached plain briefing.
Their UI entry is `0x434` with UI state one; RU has an additional alternate
control branch. Initial-load `0x442` is bookkeeping, not itself a briefing
dialog request.

**Confidence.** Medium for the positive native paths and corpus observations.
Independent resource metadata and instruction ranges reproduce from both
hashed installs. The tag probe differs intentionally from the native parser.
No whole-image reader absence or malformed-file acceptance is asserted.

**Unknown.** Initial-game dispatch to `0x434`, failure-reason producers,
malformed delimiter behavior, and other text consumers. The negative
subobjective result covers the two installed mission 10 files only.

### R2-ENGINE-050

This claim starts at UI message `0x433`; its script/packet producer is outside
the published authority here. RU handler `R2.0036` stores the ID at
controller `+0x454`; the EN branch at `L2.00133` stores it at `+0x43c`.
ID 255 posts `0x431` carrying the other argument as failure reason. IDs 254,
253 and 250 select separate banner/audio branches. Code 250 also performs a
separate string lookup. None of these four branches calls the event dialog
builder.

For other IDs, RU tests controller `+0x41c & 8` and EN tests `+0x404 & 8`.
When clear, each formats `event%d` and calls RU `R2.0037` / EN
`R2.0038`. When set, each appends the ID to a deferred collection.
The builders' constructors RU `R2.0039` / EN `R2.0040` pass the
key to section loaders RU `R2.0041` / EN `R2.0042`.
Those lowercase the formatted `#<key>` key, find its first substring in
the loaded mission source, skip key length plus two and cut at next `#`.
Missing keys create an empty NUL body. A temporary lowercase copy checks
the substring `npc`; the stored body retains its loaded bytes.

**Confidence.** Medium. The bounded UI branches, key construction, direct
builder call and copied body are raw instruction witnesses on both images.
UI field offsets and control IDs differ between EN and RU; they are not a
shared-image assumption.

**Unknown.** Packet reachability until its separate evidence is published,
deferred-event replay ordering and all bit-eight writers, behavior under
overlapping dialogs, other callers of the key-based builder, and dynamic
screen presentation.

**Amended.** A correction pass changed this claim; the overturned wording is in [`retracted.md`](retracted.md) under "ROM2 dialogue lookup source file".

### R2-ENGINE-051

RU `R2.0043` / EN `R2.0044`, vtable offset `0x80` in their actual
constructor-selected tables, call next-part methods RU `R2.0045` / EN
`R2.0046`. The counter starts at zero and increments before parsing.
Parsers RU `R2.0047` / EN `R2.0048` scan `<...>` headers in source
order, lowercase a temporary header and seek substring `part=<number>`.
The first matching admissible alternative supplies a body beginning after LF,
ending at NUL or before the next header after backing up to CR.

Hero `+0x1b8` bit four satisfies `iamfemale`, its absence `iammale`; bit two
satisfies `iammage`, its absence `iamfighter`. `npc=` parses five characters
as a decimal ID. `npcalive=` requires a matching NPC lookup result and
`npcdead=` requires no result; RU `R2.0049` / EN `R2.0050`
search the controlled unit collection by word `+0x1dc`. These tokens do not
prove a combat-health predicate. `sound=` reads through quote or semicolon;
absent sound clears the name. Absent tune is minus one. `tips=` updates a
value retained until another update or dialogue destruction.

Text goes to RU `R2.0051` / EN `R2.0052` with its end temporarily
NUL, then through wrapping and CString storage. NPC headers can select a
portrait. RU child IDs are text eleven/button twelve/portrait thirteen;
EN uses ten/eleven/twelve. The button supplies `0x470`. Handlers RU
`R2.0053` / EN `R2.0054`, actual vtable offset `0x48`, advance a
part; exhaustion optionally posts `0x45b` with tips and closes via `0x445`.

Nonempty speech names form `speech\\<name>.wav`. Empty names have default
branches for `talk`, `about`, `accept`, `reject`, `keeper`, `man` and `guard`
keys. The two mission 10 files contain no explicit sound header and their
event keys match none of these default families, so this page route selects
no speech name. EN event7 has parts one through four; RU has one through
three. Shared section names therefore do not establish equal page populations.

**Confidence.** Medium. Both positive consumers and the bounded mission 10
header census reproduce. The no-speech conclusion follows from the complete
next-part consumer and these keys, not a sound-file or whole-image absence
search. Audio-device/buffer gates were read but playback was not observed.

**Unknown.** Rendered output, audible playback, other script sound actions,
malformed headers, LF-only files, numeric-prefix collisions and untraced
lookup semantics outside the controlled-unit predicate. No other mission's
voicing is characterized.

### R2-ENGINE-052

RU `R2.0055` / EN `R2.0056` append NUL, call the imported
`CharToOemA(buffer, buffer)` and then store the converted CString.
The actual import slots are recorded beside the instructions. Static
evidence does not establish a named ANSI/OEM code page or Windows API result.

RU `R2.0057` / EN `R2.0058` read `main\\id` and store last
non-NUL byte minus ASCII zero at RU `L2.00134` / EN `L2.00135`.
The installed selectors are one in RU and zero in EN. This is one observed
writer, not a writer-completeness claim.
RU `R2.0059` / EN `R2.0060` pass bytes through unless the selector
equals one. Under one, `0x80..0xaf` add `0x30`, `0xe0..0xef` add `0x10`;
other bytes are unchanged. RU `R2.0061` / EN `R2.0062` subtract
`0x20` and truncate to a byte for sprite/advance indexing. Width measurement
uses the same mapper. Space and tilde branches remain separate.

The dialogue selects font1. RU `R2.0063` / EN `R2.0064` create it
with `graphics\\font1\\font1` and spacing two. RU `R2.0065` / EN
`R2.0066` append `.16` for sprites and `.dat` for four-byte advances.
Both font1 node hashes match across roots; each advance node has 224 entries.
The sprite header and pixels were not assigned Unicode glyph labels.

**Confidence.** Medium for the bounded load, selector, byte arithmetic and
font route. Raw instructions and import identity replace a guessed source
encoding. The selector's input bytes and font hashes are corpus measurements.
No ROM1 display claim supplies ROM2 semantics.

**Unknown.** Windows conversion configuration/results, named source encoding,
Unicode glyph identities, all selector writers, and independent writer
acceptance. Preserve the distinction between installed bytes, API output,
remapped bytes and a rendered glyph.

## Native field producers and collection transitions

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-069 | In the selected EN/RU client paths, unit u16+0x1dc receives zero, a builder argument, a copied unit key, server-unit u16+0x14c or the packet's optional u16 key. | High | ✔ promoted | [EXP-2010](../experiments/EXP-2010-rom2-fields/) |
| R2-ENGINE-070 | EN/RU selected packet and reset paths change the controlled map at receiver+0x9d0 by an entity u16 key; that membership key is separate from the NPC lookup word. | High | ✔ promoted | [EXP-2010](../experiments/EXP-2010-rom2-fields/) |
| R2-ENGINE-071 | EN/RU selected owner construction, registration and archive paths write activity owner DWORD+0x2c and u16+0x32; actor DWORD+0x14 supplies the owner receiver. | High | ✔ promoted | [EXP-2010](../experiments/EXP-2010-rom2-fields/) |
| R2-ENGINE-072 | Selected EN/RU native name parsers transfer slots 55/24 into the NPC key chain; Data.bin schema identity is Medium, while an authored NPC registry join remains Unknown. | High / Medium / Unknown | ✔ promoted | [EXP-2010](../experiments/EXP-2010-rom2-fields/) |

### R2-ENGINE-069

Client unit constructors R2.0067 (EN) and R2.0068 (RU) clear
u16+0x1dc. Builders R2.0069 / R2.0070 allocate 0x1e4, call the
constructor and set that word to the low 16 bits of their argument.
Wrappers R2.0071 / R2.0072 use the builder after a missing lookup.
The selected builder body contains no direct controlled-map insertion.
The lookup and dialogue-token contract remains R2-ENGINE-051.

Copy constructors R2.0073 / R2.0074 copy source u16+0x1dc.
Update routines R2.0075 / R2.0076 transfer server-unit
u16+0x14c into it at L2.00136 / L2.00137. These are WORD loads and stores.

Packet opcodes 108, 110, 111, 112 reach the same arm at L2.00138 / L2.00139.
Under packet.DWORD(+0x14)&0x80, it consumes a WORD from the optional
payload cursor and advances that cursor by two. Stores L2.00140 / L2.00141
transfer the temporary word to unit+0x1dc. Packet u16+0x12 is a separate
entity key, not this optional word.

**Confidence.** High for these bounded positive instruction paths in the
separately hashed clients. Complete ranges, field widths and dispatch-table
slots are replayed; the headline does not enumerate every possible writer.

**Unknown.** All writers, packet sender semantics, key uniqueness, an authored
NPC registry mapping and the reachable key population. Bulk copies,
address aliases and indirect callers remain unexamined.

### R2-ENGINE-070

The lookup iterates the associative collection embedded at receiver+0x9d0
and compares each selected unit's u16+0x1dc. In the selected packet arm,
lookup uses packet.u16(+0x12). If no entry exists and flags&0xa0 is zero,
the arm exits. Otherwise the allocation branch constructs a unit and calls
map SetAt L2.00142 / L2.00143 with that entity key and unit pointer. It also
stores the entity key in unit.u16(+4). An existing entry avoids this insertion.

The same arm classifies unit.i16(+0x104): below -600 sets state byte+0x184
at 5; otherwise threshold branches set 4, 3, 2, 1 or 0. Its six-entry state table
selects the removal branch only for 5. That branch calls map removal
L2.00144 / L2.00145 and deletes the unit when unit.u16(+4)<0x6000.
Other states and keys at least 0x6000 bypass this removal. Neither threshold
nor state is given a combat-death meaning here.

Reset R2.0077 / R2.0078 first snapshots keys from receiver+0x9d0,
then looks up, deletes and removes those entries in a loop. This reset path
establishes another membership-removal producer; its caller semantics are
untraced.

**Confidence.** High for the selected insertion, literal removal conditions,
reset loop and different key coordinates. The byte 5 route follows the native
pointer table, and the methods use an embedded map receiver rather than the
outer object's field offsets.

**Unknown.** All collection mutators and callers, unit-state meanings, control
transfer transitions, reset scheduling and whether any actor death always
causes lookup absence. Other mutators, aliases and bulk transfers remain open.

### R2-ENGINE-071

Activity reads owner DWORD+0x2c and owner u16+0x32 through actor DWORD+0x14
(R2-ENGINE-057). There is no new actor+0x32 field attribution.

Owner constructors R2.0079 / R2.0080 set DWORD+0x2c=1 and
u16+0x30=0. Selected record builders R2.0081 / R2.0024 walk
argument+0x28 records. They construct/register an owner or reuse one on a
bounded first-record arm, then set owner+0x2c=1. They set it to 0 when
Session.DWORD(+0x74)==0 and the record ordinal is 1. This is a runtime
condition, not a direct copy of an authored type DWORD.

Registration R2.0082 / R2.0083 chooses owner.u16(+4) from an
unused positive slot below 32, normally starting at 1; a conditional branch
starts at 16. A zero-key fallback uses the collection's advancing counter.
When owner+0x2c==0, it stores at+0x30 the low word of a DWORD one shifted
by signed owner.i16(+4)%16, with x86's five-bit shift-count masking, then
copies+0x30 to+0x32. Nonzero+0x2c clears+0x30 but this branch contains no
+0x32 assignment. Arbitrary negative keys do not have an absolute-value mask.

Archive routines R2.0084 / R2.0085 load DWORD+0x2c and
u16+0x30, then copy loaded+0x30 to+0x32. Their post-load roster loop sets
each iterated actor DWORD+0x14 to the owner. No archive byte layout or
whole-save acceptance is claimed.

**Confidence.** High for the bounded typed stores, signed shift arithmetic,
record-ordinal condition and owner receiver established in both clients.
This is a positive producer set, not a complete owner-field writer census.

**Unknown.** Other owner/type/mask writers, the final+0x32 value on nonzero
registration, negative-key reachability, original record-to-runtime type
classification, human/AI meaning, diplomacy/presentation meaning and GUI fog.

### R2-ENGINE-072

Name parser R2.0086 / R2.0087 scans native table L2.00146 /
L2.00147 and selects by its entry name. It obtains parameter array entry 55
and transfers its low WORD to server-unit+0x14c. The getter uses a zero-based
DWORD array index: element address equals array base plus 4*index.
The other name parser
R2.0088 / R2.0089 scans table L2.00148 / L2.00149, selects by
name and transfers parameter 24's low WORD. If that WORD is greater than 10000,
it replaces it by (word/10)%1000. The client update in R2-ENGINE-069 can
then place this WORD in the controlled-NPC lookup coordinate.

The two slot positions match the existing Units/Humans `serverID` schema
and the RU placement path in R2-ENGINE-038. That prior mapping rests on slot
numbering, not a closed native load-time table join. A name-matched native
parameter record, its complete stored serverID and the transformed NPC WORD
are therefore distinct identities. No equality with NPC registry section IDs
or ALM definition IDs is established for arbitrary authored inputs.

The selected owner builder obtains records at argument+0x28, copies a name
and other fields, and separately computes the activity+0x2c condition.
Registration computes+0x32 from the runtime owner key; it does not copy an
observed authored mask field. Existing ALM player-record grammar does not
supply these semantic defaults.

**Confidence.** High for the native name/slot/WORD transfer and >10000
transformation in both clients. Medium for the Units/Humans Data.bin schema
association inherited from R2-ENGINE-038. Unknown for the missing authored
NPC registry, ALM definition/owner type and control-record identity joins.

**Unknown.** Native table load-time identity, authored key constraints,
collisions after WORD truncation/transformation, packet-source joins and
original-compatible arbitrary-value authoring. No new authored corpus join
or complete constructor census was performed.


## First single-player campaign route

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-073 | The selected EN/RU fresh campaign starts with DLL type 2/ID 1; packed inn TALK unlocks type 1/ID 10, whose record ID supplies the first mission path. | Medium | ● active | [EXP-2011](../experiments/EXP-2011-rom2-campaign/) |
| R2-ENGINE-074 | In selected EN/RU campaign result paths, victory acknowledgement calls DLL LeaveLocation; the first failure report instead offers menu or save-slot selection. | Medium | ● active | [EXP-2011](../experiments/EXP-2011-rom2-campaign/) |
| R2-ENGINE-075 | The enabled EN/RU movie dispatcher indexes cutpaths by DLL output; output 1 selects teleport files 01..99, skips missing files and stops on player return zero. | High | ● active | [EXP-2011](../experiments/EXP-2011-rom2-campaign/) |
| R2-ENGINE-076 | Selected EN/RU additional-unit entry filters two bank ranges and uses stage-sensitive native templates; those paths do not establish the complete entering party. | Medium | ● active | [EXP-2011](../experiments/EXP-2011-rom2-campaign/) |
| R2-ENGINE-077 | The selected EN/RU movie-open path reads Common, Fading and Panaraming keys from a same-stem inline REG before opening the SMK through its native import. | Medium | ● active | [EXP-2011](../experiments/EXP-2011-rom2-campaign/) |

### R2-ENGINE-073

The matching DLL text/data sections contain an initial constructor supplying type 2 and ID 1. NewGame places that record in availability and current. These are separate raw fields. The client's type-2 route invokes town presentation; the RU consumer reads current ID 1 and dispatches controller scene +0x110 through its actual presentation virtual slot.

Stage 10 EnterInn supplies packed 0x300a0205 among three entries. Both client TALK bodies format npc517talk10, open its dialogue and immediately call DLL TalkTo with that same value. Both town sources contain that header once. TalkTo extracts low-word ID 517, topic 10 from bits 16..27 and kind 3 from bits 28..30. Kind 3 adds the matching type-1 ID-10 record to availability. The page acknowledgement is not the gate in this selected body.

Type-2 ID-1 Leave removes the initial record without selecting a movie. Both selected world-map consumers choose an available record by rectangle or first entry, pass its pointer to EnterLocation and post UI 0x468. The type-1 branch reads current record ID to form the ALM path; R2-ENGINE-033 supplies the earlier RU loader boundary. Initial record ID 1 and mission ID 10 must not be conflated.

**Confidence.** Medium for the composed first route. Raw constructors, packed decoding, client ordinal bindings, pointer-table targets and selected calls identify the sources without a mission-number assumption. The entire input/event schedule was not executed. DLL section equality is one code dependency.

**Unknown.** Automatic initial briefing dispatch, every town interaction, arbitrary availability mutations, and the complete character-selection schedule. The selected campaign path is not a native playthrough witness.

### R2-ENGINE-074

EN result tick L2.00524 and RU L2.00127 require player byte +0x41 nonzero and DWORD +0x2c zero. In their live-player branch, failure register Session+0xbe14>=1 takes priority; otherwise a nonzero win counter +0xbe0c can emit B5. These fields are distinct from the DLL bank. R2-ENGINE-060 supplies the action producers and the mission-10 failure reason 19.

The actual packet tables map B5 to UI 0x430 and B4 to UI 0x433 with selector 255 and reason. The victory body sets the controller win flag (RU +0x3fc, EN +0x3e4) and builds a general-table report using slot 0x8c. Its button builder contains 0x41d. The automation branch posts 0x41d directly. In campaign mode with the win flag set, that handler calls LeaveLocation, conditionally dispatches its movie output and consumes the available list. With the win flag clear, the same message instead cleans up and posts menu 0x421.

The first failure report reads cached mission text at reason minus two and constructs buttons 0x445/0x446. Its recorded button travels through close notification 0x44c. Selected close branches RU L2.00525 and EN L2.00526 post 0x41e for 0x445, otherwise 0x418. The former cleans up and posts 0x421; the latter constructs the save-slot selector. No direct call through the bound LeaveLocation pointer appears in these selected branch bodies. Their explicit choices differ from the first victory's successor/movie route.

**Confidence.** Medium for the composed report/choice routes. Instructions, actual UI/packet jump tables and report-object vtable targets distinguish victory continuation from the selected failure choice. The local EN tick closes the earlier RU-only result-consumer boundary. No whole-image no-call claim is made.

**Unknown.** Runtime presentation/timing, later loading from the failure selector, other failure selectors/sources and the complete dispatch population. Initial automatic briefing scheduling remains Unknown.

### R2-ENGINE-075

Both startup bodies bind main\text\cutpaths.txt to the string-table object used by the movie dispatcher. Its CRLF loader assigns row order; the getter uses base index plus the supplied output index. The installed sources have six rows and index 1 is teleport.

When enabled, both movie dispatchers mark a seen bit, format video\<stem>\%02d.smk for integers 1 through 99 and check resource existence. Missing nodes are skipped. Existing nodes reach the native player; player return zero stops the loop. The player may return zero for an input interruption, so this does not mean every matching file is displayed. The independently established mission-10 DLL case returns output 1 (R2-ENGINE-048).

The selected first teleport SMK node exists in both video archives and differs in hash between locales. The companion REG is identical, 476 bytes. Exact identities are retained in the experiment rather than inferring audiovisual equivalence.

**Confidence.** High for this conditional dispatcher/index/resource contract. Complete selected dispatcher spans, native CRLF/index helpers, startup bindings, PE mappings and archive-node measurements establish table-selected stems and numeric segment iteration, excluding a hard-coded teleport stem within this dispatcher. Capstone 5.0.7 and pefile 2024.8.26 decode/map the selected ranges; no global consumer or absence enumeration is claimed.

**Unknown.** Movie-enable setting provenance, decoded media, full playback/audio behavior, other output values and later campaign movie ordering.

### R2-ENGINE-076

In UI 0x468's type-1 branch, both clients select ID i+1 for i=0..19 only when DLL bank 512+i and 532+i are both nonzero. Nonempty selection reaches native additional-unit creation with bank slot 768 as stage.

Both selected native gateways try template 10000+10*ID. If an additional variant exists, a stage-specific lookup 9998+10*ID+stage/10 can replace that template. They write ID to server-unit WORD +0x14c, allocate an entity key, assign owner and register actor/group before emitting its update. This distinguishes the template/NPC identity from the runtime entity key; R2-ENGINE-070/072 supply the independently established identity boundaries.

R2-SESSION-023 gives the selected DLL transition algebra: ordinary Leave clears 512+i and normalizes retained 532+i. An immediate zero selection does not establish an empty final party. Main controlled hero selection, map population and possible intervening bank writes are separate inputs.

**Confidence.** Medium for entering-party inference. The filter, lookup arguments and selected registration stores are direct code measurements in both images. No fixed roster or preservation rule follows from these additional-unit bodies.

**Unknown.** Complete final roster, all bank writers, template contents, mission-load mutations, carryover serialization and main-hero selection/construction inputs.

### R2-ENGINE-077

Selected native movie-open bodies RU L2.00527 and EN L2.00528 derive a .reg companion from the SMK name and load an inline registry. They request Common/startx, starty, nFadings and nPanaramings with zero defaults. Fading1..N rows request startframe/endframe and floating startfade/endfade; Panaraming1..N rows request startframe/endframe and stepx/stepy. The native spelling is Panaraming.

The selected lower open bodies RU R2.0110 and EN R2.0108 pass a resource-backed source to the image's named SmackOpen import. The first teleport companion has 14 records, a zero-byte pool and exact framed size 476 in both locales, consistent with R2-ASSET-006. This is a native first-route consumer, not a general REG type specification.

**Confidence.** Medium for the consumer contract. Original key strings, instruction call arguments, FPU result stores and resource geometry establish the selected requests. Native record lookup/type handling and arbitrary writer acceptance were not closed.

**Unknown.** General REG record kinds, malformed/missing companion behavior, full frame-plan application, rendering, audio synchronization and third-party codec internals.

## Native group lookup and activity ordering

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-085 | EN/RU literal group compilation overwrites a DWORD-keyed index in current owner-list and group-list order; duplicate keys select the last visited group. The selected registration path appends owners. | High | ✔ promoted | [EXP-2012](../experiments/EXP-2012-rom2-activity/) |
| R2-ENGINE-086 | EN/RU activity's nonzero-owner branch compares signed actor words +0x94/+0x96, also used as current and maximum refill values; a lower current value sets that actor's group active. | High / Medium | ✔ promoted | [EXP-2012](../experiments/EXP-2012-rom2-activity/) |
| R2-ENGINE-087 | EN/RU recount resets activity, visits its actor collection and zero-class owners' +0x38 actors as producers, then marks collection actors; later phases update additional receivers' masks. | High | ✔ promoted | [EXP-2012](../experiments/EXP-2012-rom2-activity/) |

### R2-ENGINE-085

Compiler spans EN L2.00150..L2.00151 and RU L2.00152..L2.00153
walk the global owner list, then each owner's DWORD+0x28 group list.
Each visit passes group+0x1c to the index and stores the current group
pointer into the returned value slot. Equal keys therefore overwrite that
slot. Literal-reference wrappers R2.0090 / R2.0091 read it back.
The index's equality routine compares the complete key DWORD; its hash is
the unsigned key shifted right four before bucket division.

Both iterators start at list+4 and follow each node's DWORD+0, returning
node.DWORD(+8). This is current list order. Registration R2.0082 /
R2.0083 reaches append wrappers R2.0092 / R2.0093, whose
native helpers link the new owner after the previous tail. This bounded
producer does not sort by owner key.

**Confidence.** High for the selected compiler, lookup, equality, iterator
and append paths. Raw receiver flow excludes first-match and owner-qualified
lookup on this path. The repeated unconditional slot store excludes retaining
the first duplicate. List links, rather than a key comparison, decide traversal.
The evidence is 34 complete selected EN/RU spans in ranges.tsv.

**Unknown.** Every list producer, archive reconstruction order, removal or
reordering, dynamic references and malformed-key reachability. Last current
list occurrence does not establish a universal greatest-owner rule or equal
ordering after every campaign/load route.

### R2-ENGINE-086

The full producer R2.0094..L2.00154 / R2.0095..L2.00155
reads actor.DWORD(+0x14) as owner. Its nonzero owner.DWORD(+0x2c)
branch compares signed actor.i16(+0x94) with actor.i16(+0x96).
Only strict less-than sets active byte +0x45 through actor+0x70,
then group+0x3c. Equality and greater-than skip this store.

Refill functions R2.0096..L2.00156 / R2.0022..L2.00157
use the same actor-relative words: +0x94 is the current accumulator's
whole value, and +0x96 supplies the maximum and the upper clamp.
They write only the current word. The positive-current branch has a
four-phase gate; the negative-current branch decrements that word.
The activity producer is therefore comparing these refill pool values,
not the actor's position words. R2-ENGINE-030 identifies the health
regeneration role; no new full regeneration formula is asserted here.

**Confidence.** High for signed widths, receiver identity, strict predicate,
active destination and current/maximum roles in these separately hashed
clients. Medium for the player-facing health label, inherited from the
earlier regeneration interpretation without a new GUI/packet observation.

**Unknown.** All damage and refill producers, health presentation, dead-actor
collection membership and how every actor reaches the selected producer.
The classification of owner+0x2c remains the R2-ENGINE-071 frontier.

### R2-ENGINE-087

Recount R2.0097..L2.00158 / R2.0098..L2.00159 first
calls reset and clears local counters +0x1614/+0x1624. It visits the
collection at local DWORD+0x1610, calling the full activity producer
for every returned actor. It then walks the global owner list and calls
the same producer for a nonnull owner.DWORD(+0x38) only when
owner.DWORD(+0x2c)==0. This extra producer precedes all mark calls.

It visits the +0x1610 collection again, calling the mark routine on each
actor. Mark uses current coverage first; nonzero coverage sets group active,
increments +0x1614 and copies its low word to actor+0x1a4. Only zero
coverage reaches the forced-DWORD(+0x48) branch, which can set active and
increment both counters. A previously stored actor+0x1a4 does not create
coverage in this routine.

Later recount loops copy coverage to receivers in Session.DWORD(+0x7c)
collection+0xc and combine four sampled cells for receivers in that object's
collection+4. The final +0x1618 counter is the +0x1610 list count minus
+0x1614. Reset clears the 0x1210-byte bitmap, clears group active bytes
only for nonzero-class owners, and clears actor+0x1a4 in the separate
global actor list. R2-ENGINE-057 supplies the embedded object coordinates.

**Confidence.** High for this complete reset/producer/mark ordering and
typed receiver flow in both clients. The second producer loop rules out a
recount limited to the +0x1610 collection. Coverage is rebuilt before the
mark pass; a stale per-actor word is not its input.

**Unknown.** Collection membership and every producer of owner+0x38,
owner-class meaning, GUI fog, final masks on every registration path,
other recount callers, session-gate meaning and command-byte producers.
The selected unit gate reads actor.DWORD(+0x1c4).BYTE(+9), but its
nonzero branch only bypasses this activity check; it proves neither every
explicit order's producer nor successful execution after all later checks.

## Completion presentation consumers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-093 | Selected EN/RU victory reports use general UI text and dialogs labels; the startup source join identifies main row 140 and dialogs rows 43/154, rather than mission text. | Medium | ● active | [EXP-2013](../experiments/EXP-2013-rom2-completion/) |
| R2-ENGINE-094 | The selected EN/RU outer movie player returns zero for six input/message IDs, but returns one for open failure or a terminating tick; both paths clean up the player. | High | ● active | [EXP-2013](../experiments/EXP-2013-rom2-completion/) |
| R2-ENGINE-095 | Selected EN/RU movie ticks terminate on a library-descriptor comparison or presentation failure; that client boundary does not establish physical EOF or successful decoding. | Medium | ● active | [EXP-2013](../experiments/EXP-2013-rom2-completion/) |
| R2-ENGINE-096 | Selected EN/RU movie opens seek a RES member and pass its backing handle to SmackOpen; the member reader's clamp does not prove a codec EOF bound. | Medium | ● active | [EXP-2013](../experiments/EXP-2013-rom2-completion/) |
| R2-ENGINE-097 | Selected EN/RU movie companions feed typed numeric getters and equality-triggered frame plans; the measured first companion contains two Fading segments and no Panaraming records. | Medium | ● active | [EXP-2013](../experiments/EXP-2013-rom2-completion/) |

### R2-ENGINE-093

The published selected victory route sets the win flag and supplies flat general-table slot 0x8c (R2-ENGINE-074). Native constructors EN R2.0099 and RU R2.0100 store that pointer at report+0x7c and install their actual report vtables. Slot +0x78 targets initialization EN R2.0101 / RU R2.0102. That method passes the stored pointer to the text-control constructor. It builds button ID 2 with dialogs index 0x2b and command 0x41d; button ID 3 uses index 0x9a and command 0x446. Passed binding codes are 0x56 and 0x43; complete input interpretation is not claimed.

The selected startup excerpts load main\\text\\main.txt as the first shared string table and bind dialogs.txt to the table object used by these labels. The complete CRLF/index helpers append row pointers, record the table base and use base plus row index; CharToOemA is imported at the conversion call. Both installed main sources have 369 CRLF rows, and both dialogs sources have 176. Main row 140 is EN `Mission Completed`; dialogs rows 43 and 154 are `~Victory!` and `Continue`. Selected RU row hashes/lengths are recorded without asserting a named encoding. These report inputs are separate from the cached mission-text failure input established by R2-ENGINE-049/R2-ENGINE-074.

**Confidence.** Medium for the composed startup/resource source join. Both image call operands, constructors, report vtable targets, complete initialization/index helpers and selected source measurements agree. A zero-filled PE global was not treated as a live bank value. The entire startup and input schedule was not executed, so this is a conditional selected route.

**Unknown.** Earlier or later live table mutation, report fonts/background resources, generic input bindings, RU displayed glyphs and other report producers.

### R2-ENGINE-094

Complete outer players EN R2.0103..L2.00529 and RU R2.0104..L2.00530 use PeekMessageA with removal enabled. Message IDs 0x12, 0x10, 0x201, 0x204, 0x100 and 0x104 take the cleanup-and-zero return. IDs 0x12/0x10 are translated and dispatched first; the four key/mouse IDs take that return directly. Other queued messages are translated/dispatched and the loop continues.

With no queued message the player calls its tick with argument 1. A zero tick result takes cleanup and returns one. An unsuccessful presentation-open takes that same cleanup-and-one return before entering the message loop. Cleanup EN R2.0105 / RU R2.0106 invokes lower close and clears owned plan storage; the destructor invokes cleanup again. Lower close conditionally invokes the named SmackBufferClose, SmackClose and SmackBlitClose imports and clears their stored handles.

The complete selected dispatcher (R2-ENGINE-075) stops remaining numbered segments on outer return zero, while return one continues the bounded search. Its temporary controller presentation bit is cleared on batch cleanup; its seen bit set before the loop is retained by this method. This is a batch-control result, not a decoded-success result.

**Confidence.** High for this conditional native return/cleanup contract. Complete paired player, dispatcher, lower close and cleanup spans expose every explicit ordinary return in those methods. The same result cannot be interpreted as both a distinct successful EOF signal and the observed open-failure path. No wider input or consumer enumeration is claimed.

**Unknown.** Exceptions, reentrant effects of dispatched messages, timing, the provenance of the movie-enable setting, other callers and whether original media is rendered successfully.

### R2-ENGINE-095

Complete native ticks EN R2.0107..R2.0108 and RU R2.0109..R2.0110 call a blocking SmackWait loop when argument bit 1 is set. An unfocused SmackBufferFocused result returns one without advancing this tick. The active path applies Fading/Panaraming callbacks, calls SmackDoFrame, and invokes presentation receiver slots +0x64 and +0x80. Either receiver failure returns zero.

After frame processing, the client compares library descriptor DWORD +0x374 against DWORD +0xc minus one. Equality returns zero; otherwise it calls SmackNextFrame, applies the pan step and returns one. R2-ENGINE-094 maps zero to the outer player's one return. The first measured SMK header DWORD +0xc is 167, but the client-to-library descriptor/header binding is unverified.

**Confidence.** Medium for the composed completion boundary. Complete paired ticks, named IAT entries and exact descriptor operands establish the client predicate. They do not establish the library's interpretation, decode outcome or the physical stream endpoint. The presentation receiver's dynamic implementation was not closed.

**Unknown.** Descriptor-to-header identity, physical EOF versus truncated/error termination, codec internals, focused/wait timing, actual receiver behavior, audio synchronization and frame fidelity.

### R2-ENGINE-096

Complete lower opens EN R2.0108..L2.00531 and RU R2.0110..L2.00532 open the selected resource-backed source, call its read/position method with count zero, then pass source+4 to the named SmackOpen import with flags 0xff000 and final argument -1. Complete source-open methods retain the backing stream at +0x10, member base/length at +0x14/+0x18 and cursor at +0x1c. Source+4 is copied from backing-stream+4.

Complete read/position methods EN R2.0111..L2.00533 and RU R2.0112..L2.00534 clamp requested bytes to member length minus cursor and call backing-stream virtual +0x30 with base plus cursor before checking the zero count. A zero count therefore positions the stream without a member-byte read. Nonzero reads advance the cursor and call stream virtual +0x3c. Passing the backing handle to the import does not show that subsequent codec reads pass through this member clamp.

R2-ENGINE-075 supplies the indexed first stem. In each measured video.res, exactly one of its 99 selected numeric SMK keys exists: teleport/01.smk, 6753596 bytes. Its EN/RU payload hashes differ. Both teleport/01.reg nodes exist, 476 bytes and equal hashes. These are archive-node measurements, not a census of loose overrides or aliases.

**Confidence.** Medium for the composed resource-to-import boundary. Complete paired source/open/read spans, import operands and checked archive geometry expose the backing handle and positioned member. The virtual seek/read implementation and codec internals are unresolved.

**Unknown.** Codec read extent, physical archive/member EOF, loose-file or alias precedence, truncation behavior and decoded equivalence between locales.

### R2-ENGINE-097

R2-ENGINE-077 supplies the same-stem key requests. Selected integer getters EN R2.0113 / RU R2.0114 mask record kind with 0xe, shift by one, and for result 1 load DWORD +4. Double getters EN R2.0115 / RU R2.0116 select result 2 and load the QWORD at +4/+8. Missing records return the supplied defaults. These are conditional getter branches, not a general key lookup/comparator specification.

Both measured first companions have 14 records, a zero-byte pool and exact inline length 476 (R2-ASSET-006). Named records store Common/startx=0, starty=0 and nFadings=2; nPanaramings is absent. Fading1 stores start/end frames 0/5 and kind-4 doubles 0/1; Fading2 stores 161/166 and 1/0. There are no Panaraming records in this selected companion. The native open stores the double getter results as float endpoints in 16-byte segment arrays.

Complete Fading callbacks EN R2.0117 / RU R2.0118 compare the next segment start for equality with descriptor+0x374, pass end minus start and endpoints to the native fade method, then advance the segment index. Complete Panaraming callbacks compare start first, set stored x/y steps, and on end equality clear steps and advance their index. A post-SmackNextFrame step adds those stored values to the destination coordinates. The import name is SmackNextFrame; no decoded frame effect is inferred from the method names.

**Confidence.** Medium for the typed companion/application join. Checked framing, record values, paired getter masks and complete array/callback instructions support the conditional numeric plan. Live lookup success and visual interpretation were not observed. Equal companion bytes are one data witness, not two independent confirmations.

**Unknown.** Full lookup/sort/type acceptance, malformed or missing companion behavior, effects of missed frame equality, visual fade interpolation/clamping, pan coordinate appearance and codec/hardware timing.

## Selected campaign actor construction

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-105 | Selected EN/RU character producers reuse owner +0x38 or construct and register an actor; their Session +0x74==0 branch writes actor WORD+0x14c=21, with the EN opcode-0x48 submission route traced conditionally. | Medium | ● active | [EXP-2014](../experiments/EXP-2014-rom2-party/) |
| R2-ENGINE-106 | The selected EN location-leave arm forwards nonzero scenario slot 773 to an EN/RU addition producer whose IDs 21..30 branch avoids new creation when owner actors already contain the incoming WORD+0x14c identity. | Medium | ● active | [EXP-2014](../experiments/EXP-2014-rom2-party/) |
| R2-ENGINE-107 | EN/RU literal Unit references 10001..11000 use reference-10001 and compare actor WORD+0x14c-21 in the selected owner's list, gated by Session +0x74==0 and nonnull owner +0x38; complete controlled membership remains unproved. | High / Medium | ● active | [EXP-2014](../experiments/EXP-2014-rom2-party/) |
| R2-ENGINE-108 | Selected EN/RU map construction conditionally reuses its first owner, matches actor rows through owner DWORD+8 and forms groups by actor-row DWORD+0x40; fresh owners clear +0x38 and complete transition carryover is unproved. | Medium | ● active | [EXP-2014](../experiments/EXP-2014-rom2-party/) |

### R2-ENGINE-105

EN producer `L2.00160..L2.00161` and RU `L2.00162..L2.00163` first test
owner DWORD(+0x38). A nonzero actor pointer is returned through a reuse
branch. A selected actor-state branch appends it to owner +0x24, refills two
current fields from their maxima and creates a group only if actor +0x70 is
null. It does not allocate a second actor on that path.

The zero-pointer branch allocates a human using flag-selected native names.
Bit 0x40/0x80 selects one of four measured `Start_*` prefixes. The native
suffix is `.f5` for zero flags; other flags select `.f` or `.m` plus the
decimal low six bits. These are constructor inputs, not measured template
contents or a complete authored roster.

That branch obtains a separate entity key, assigns actor +0x14=owner and
registers owner actor/group lists. It writes owner +0x38=actor at EN
`L2.00164` / RU `L2.00165`. Session DWORD(+0x74)==0 selects the WORD(+0x14c)
literal 21 at EN `L2.00166` / RU `L2.00167`. RU has an additional native call
before the owner-field store; its complete effect is unexamined.

EN controller NewGame mode 2 reaches session initialization with +0x74=0.
Later map build can overwrite that field. EN submit `L2.00168` conditionally
calls `L2.00169`, which fills native opcode BYTE(+9)=0x48, owner WORD(+5),
zero WORD(+7) and character bytes +0x0a..+0x10 before calling its sender.
The mode-1/2 submit path also calls packet drain `L2.00170`. For packet
BYTE(+4)<1, dispatcher `L2.00171` selects opcode 0x48 through table selector
22 to `L2.00172`; that arm resolves the owner and calls the producer with
bytes +0x0a..+0x0f. This buffer is not a complete proved wire grammar.

**Confidence.** Medium. Complete producer intervals in each locale support
the conditional allocation/reuse and identity writes. The selected EN
submission and dispatcher have direct calls and measured table slots.
Indirect sender delivery, native templates and later field changes prevent
an end-to-end initial controlled-group assertion.

**Unknown.** Complete initial roster/size, installed template contents,
runtime packet acceptance, all owner-field writers, virtual constructor
effects and whether every campaign route reaches this producer.

### R2-ENGINE-106

In the selected EN UI 0x41d arm, ScenarioLeaveLocation precedes a read of
scenario slot 773 through `L2.00173`. A nonzero value enters a temporary ID
list. Stage slot 768 and that list are passed to `L2.00174`.

The complete selected EN `L2.00174..L2.00175` and RU
`L2.00176..L2.00177` producers use the incoming low WORD ID and stage for
native template lookup. IDs 21..30 take a separate branch. It visits the
selected owner's +0x24 list and compares actor WORD(+0x14c) to the incoming
ID. A match bypasses new actor construction. The no-match branch constructs
a human, stores WORD(+0x14c)=ID, obtains an entity key, assigns an owner
pointer and registers owner actor/group collections. Other incoming IDs use
different native creation branches. R2-ENGINE-076 covers the separately
measured type-1 entry filter.

**Confidence.** Medium. Both producer bodies and the selected EN caller
support this conditional addition seam. The slot's source and indirect
runtime delivery remain open; this does not prove a particular authored join.

**Unknown.** Slot-773 producer, authored mission/stage schedule, template
data, actor positioning/acceptance, complete transfer or carryover, and all
other addition routes. No party size or joining actor is predicted.

### R2-ENGINE-107

EN `L2.00178..L2.00179` and RU `L2.00180..L2.00181` route references
10001..11000 inclusive to helpers `L2.00182..L2.00183` /
`L2.00184..L2.00185`, passing reference-10001. The helpers return null when
Session DWORD(+0x74)!=0 or selected owner DWORD(+0x38) is null. Otherwise
they compare zero-extended actor WORD(+0x14c)-21 with the offset in owner
+0x24 and return the first encountered match. Literal 10001 therefore
compares WORD(+0x14c)=21, and 10002 compares it to 22. This is conditional
numeric identity, separate from actor DWORD(+4) entity keys and roster slots.

An alternate global list requires nonempty actor string +0x80 and compares
the same arithmetic. Its mismatch branch advances an iterator before the
common advance, so complete alternate coverage is not claimed. References
below 10001 use an ordinary lookup with the unmodified argument; references
above 11000 take a separate string-table path.

**Confidence.** High for the bounded EN/RU numeric interval, gating and
owner-list comparison: both complete bodies include their local branches.
Medium for interpreting this as the complete campaign-controlled identity
route: collection producers and the alternate branch leave live alternatives.

**Unknown.** Actor presence, uniqueness, full alternate-list coverage,
unexamined collection helper behavior, references above 11000 and dynamic
references. This does not guarantee that any literal successfully resolves.

### R2-ENGINE-108

Selected EN/RU map builders `L2.00186..L2.00187` /
`L2.00113..L2.00188` call owner builders `R2.0081..L2.00189` /
`R2.0024..L2.00114`. When Session DWORD(+0x74)==0, row index==1 and the
selected count wrapper returns 1, they reuse the existing selected owner.
Otherwise they allocate and construct an owner. Constructors
`R2.0079..L2.00190` / `R2.0080..L2.00191` clear owner +0x38.

The actor builders `L2.00192..L2.00178` / `L2.00115..L2.00180` match
signed actor-row WORD(+0x14) against owner DWORD(+8), which owner construction
copies from owner-row DWORD(+4). Generated owner WORD(+4) is a separate
coordinate. A successful actor gets actor +0x14=owner and owner-list
registration. Groups use actor-row DWORD(+0x40); group append sets group
+0x44=actor +0x14. Missing owner, state, positioning and actor-validation
branches can skip a row.

**Confidence.** Medium. Independently measured EN/RU calls and field writes
support the selected reuse, owner matching and grouping boundary. Input rows,
nested loaders, templates and positioning are not measured. The positive
reuse branches here and in R2-ENGINE-105 do not prove complete carryover.

**Unknown.** Initial authored owner/actor rows, reused-owner contents,
count-wrapper callee semantics, every postactor branch, complete entering
membership, transition serialization and SAV restoration.

## Selected construction and literal-list dependencies

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-123 | Selected EN/RU list helpers append at the tail and iterate from the head; the literal Unit helper compares identities in the selected owner's nonnull payload prefix. | High | ✔ promoted | [EXP-2016](../experiments/EXP-2016-rom2-controlled-construction/) |
| R2-ENGINE-124 | Selected EN/RU character construction calls a parameter-record loader before registration and the conditional WORD+0x14c=21 store; object class and complete ordinary-start reachability remain open. | Medium | ✔ promoted | [EXP-2016](../experiments/EXP-2016-rom2-controlled-construction/) |
| R2-ENGINE-125 | Selected EN/RU Session+0x74 stores use signed initialization mode<2 and accepted map-object DWORD+0xd8>1; the map reuse comparison reads a distinct owner-collection DWORD+0x0c. | High / Medium | ✔ promoted | [EXP-2016](../experiments/EXP-2016-rom2-controlled-construction/) |
| R2-ENGINE-126 | Four Start_* names match stored Humans rows with parameter24=10210..10213 in two identical client payloads and the RU server payload; native table load-time identity remains unclosed. | High / Medium | ✔ promoted | [EXP-2016](../experiments/EXP-2016-rom2-controlled-construction/) |
| R2-ENGINE-127 | Initial object classes, complete selected-owner membership and successful literal Unit resolution remain Unknown in the 54 measured native construction/list intervals. | Unknown | ✔ promoted | [EXP-2016](../experiments/EXP-2016-rom2-controlled-construction/) |
| R2-ENGINE-135 | Twenty-four EN/RU routine pairs (48 complete bodies; 16,475 EN and 16,746 RU instructions) that reach the EXP-2016 seeds were decoded; 15 pairs match in normalized shape and 9 differ. | High / Medium | ● active | [EXP-2017](../experiments/EXP-2017-rom2-initial-collection/) |
| R2-ENGINE-136 | Static class records name the 0xa80-byte owner class Player (create slot EN L2.00193 / RU L2.00194) and a 0x254-byte chain Human > Humanoid > Unit > Token > CObject. | High / Medium | ● active | [EXP-2017](../experiments/EXP-2017-rom2-initial-collection/) |
| R2-ENGINE-137 | Heap Player construction appears in EN L2.00195 / RU L2.00196, reached through EN L2.00197 / RU L2.00198 and the dispatcher EN L2.00171 / RU R2.0119; EN L2.00199 / RU L2.00200 builds a stack-frame Player and allocates no heap object. | High | ● active | [EXP-2017](../experiments/EXP-2017-rom2-initial-collection/) |
| R2-ENGINE-138 | In the measured map builder the owner-build, actor-build and literal-reference-caller calls occur in that order with no exit between the first and last; the caller EN L2.00201 / RU L2.00116 has two direct callers. | High / Medium | ● active | [EXP-2017](../experiments/EXP-2017-rom2-initial-collection/) |
| R2-ENGINE-139 | Ten of 13 actor-list append sites in the new population load their receiver from reg+0x24 in each locale; the base object is not identified at any site. | Medium | ● active | [EXP-2017](../experiments/EXP-2017-rom2-initial-collection/) |
| R2-ENGINE-140 | The 48 bodies hold 109 EN / 113 RU indirect calls and 14 indirect jumps with 14 decoded jump tables per locale; outside .text, selected entries have raw pointers only at one .data class-record slot and five .rdata cells per locale. | High | ● active | [EXP-2017](../experiments/EXP-2017-rom2-initial-collection/) |
| R2-ENGINE-141 | Positive ordinary-start membership of the owner actor collection, its initial contents and order, the startup predicate and literal-consumer success remain Unknown at the measured direct/indirect boundary. | Unknown | ● active | [EXP-2017](../experiments/EXP-2017-rom2-initial-collection/) |

### R2-ENGINE-123

Owner construction creates a collection at owner+0x24. Its constructor
initializes the embedded header at collection+4, with head+4, tail+8 and
count+0x0c cleared. EN append chain L2.00202 -> L2.00203 -> L2.00204 and RU
L2.00205 -> L2.00206 -> L2.00207 allocate a node with next=0 and previous=old
tail, place the supplied pointer at node+8, link old tail.next or empty head,
and set the new tail. The node allocator increments header+0x0c. No pointer
comparison or deduplication occurs in these append bodies.

Iterator begin L2.00208 / L2.00209 reads the header's head. Step
L2.00210 / L2.00211 returns node+8 and advances its node coordinate to next.
Next L2.00212 / L2.00213 invokes one such step. The selected owner's list is
therefore traversed in tail-append order while payload pointers are nonnull.
The literal helper in R2-ENGINE-107 returns its first encountered matching
WORD identity. A null payload terminates that primary loop even if a later
node exists. The alternate global-list branch remains separate.

The character producer passes its constructed or selected reused pointer to
this append chain. The accepted map-row branch also appends through the same
chain after its validation and positioning calls. These call sites establish
conditional pointer delivery, not a complete entering member population.

**Confidence.** High for the bounded helper contract on a well-formed list
and successful allocation. Complete matched EN/RU bodies include every local
branch and return, the node link stores and the iterator receiver join.
They exclude front insertion, reversal and deduplication inside this chain.
No global list-mutation or writer absence is asserted.

**Unknown.** Allocation-failure recovery, null/duplicate payload reachability,
other append/remove/load producers, initial content, complete membership,
postcallback changes and runtime literal success.

### R2-ENGINE-124

The complete character producers in R2-ENGINE-105 test owner+0x38 first.
The nonzero-pointer branch returns the existing object and conditionally
appends it; that branch contains no second actor allocation. The zero-pointer
branch requests 0x254 (596) bytes and passes one of four native Start_*
input prefixes, a flag-selected suffix and two additional scalar arguments
to EN constructor L2.00214..L2.00215 / RU L2.00130..L2.00216.

That constructor calls a base constructor, stamps a measured vtable pointer
(EN L2.00217 / RU L2.00218), then calls EN R2.0088..L2.00219 /
RU R2.0089..L2.00220 on the same object. R2-ENGINE-072 supplies that loader's
name-matched parameter-24 low-WORD store and conditional transformation.
R2-ENGINE-126 supplies separately measured stored names and values, with the
load-time table join still open. A constructor call and vtable store alone
do not identify the full native class or inheritance.

After the constructor returns, the producer modifies fields, calls a virtual
slot, obtains a separate entity key and registers the pointer with owner+0x24.
It then writes owner+0x38 and, only under Session+0x74==0, overwrites
WORD+0x14c with literal 21. Later external calls remain unexpanded. The reuse
branch does not pass through this conditional literal store.

**Confidence.** Medium for the construction/identity origin chain as a whole.
The direct same-object constructor, loader and append receiver joins are
measured in both complete locale bodies. Base effects, virtual slots and
later external calls leave live alternatives for type and final state.

**Unknown.** Full object class, base initialization of WORD+0x14c, final
contents after callbacks, native table load-time identity, indirect submission
acceptance, positive initial membership and complete ordinary-start reachability.

### R2-ENGINE-125

EN initialization L2.00221..L2.00222 / RU R2.0010..L2.00112 set
Session DWORD+0x74 to the signed predicate argument<2. They separately set
+0x70 from argument>0. These are distinct stores, not a classification inferred
from offset positions. Initialization mode 2 therefore selects gate zero at
that instruction; complete caller delivery remains a separate boundary.

EN map builder L2.00186..L2.00187 / RU L2.00113..L2.00188 branch on the
returned map object's DWORD+0x14. The failure branch returns before the
selected gate overwrite. On the other branch, they set Session+0x74 to the
signed predicate map-object DWORD+0xd8>1 before calling the owner builder.
The authored origin and meaning of map-object+0xd8 were not followed through
its nested constructor/loader.

The owner builder's reuse predicate calls EN L2.00223 -> L2.00224 /
RU L2.00225 -> L2.00226. The complete callee returns receiver DWORD+0x0c.
The receiver is the same global owner collection used to select its first
owner, not that owner's +0x24 actor collection. The comparison to literal 1
therefore supplies no actor roster size. Complete mutation semantics of this
owner-collection field remain outside the measured dependency population.

**Confidence.** High for the local signed comparisons, distinct stores and
receiver/field joins in the complete EN/RU bodies. Medium for their ordinary
mission-entry state interpretation: map-object field origin, intervening
external effects and full runtime reachability remain open.

**Unknown.** Authored map-field meaning, all other writers, postcall gate
persistence, initial selected owner/list content and complete startup/entry.

### R2-ENGINE-126

A complete public-grammar replay reaches zero residue on the preserved EN/RU
client world.res:data/data.bin payloads and RU world_srv.res:data/data.bin.
The client payloads are byte-identical; they are one stored-content dependency.
The server payload is distinct. Selected row hashes match across all three.

The four character-producer prefix stems Start_MF, Start_FF, Start_MM and
Start_FM each have exactly one case-insensitive name match in stored group F
Humans rows 26..29. Their numeric parameter 24, labelled serverID by title 25,
is 10210, 10211, 10212 and 10213 respectively. These are stored values, not
runtime object identities. Applying R2-ENGINE-072's low-WORD and conditional
(word/10)%1000 arithmetic to each gives 21.

The constructor passes a stem plus a native suffix; the selected loader has
suffix and separate Hero/Man string branches. The native table's load-time
identity with these stored group-F rows remains unclosed. Matching a stem,
position and value cannot replace that producer join. The conditional
arithmetic result does not guarantee that the object receives it.

**Confidence.** High for the three actual payload parses, exact selected name
matches, row/parameter measurements and stated arithmetic. Medium for native
constructor use of those installed rows because load-time identity and all
name-normalization branch dependencies are unresolved.

**Unknown.** Native table load-time join, complete suffix/name branch behavior,
base/default WORD contents, runtime template acceptance, later overrides and
arbitrary authored inputs. No EN server payload was measured.

### R2-ENGINE-127

The selected 54 complete code intervals comprise six preregistered routine
pairs and 21 additional dependency pairs. Positive pointer reuse, constructor
calls, list-link stores and local field writes are measured. The population
does not include complete sender delivery, every virtual target, map metadata
loaders or all collection producers. Global direct-edge/raw-pointer scans
supply navigation only and assert no writer/caller absence.

**Confidence.** Unknown for the initial class/member population and literal
Unit success. The unexpanded dependencies remain live alternatives. The
measured conditional paths do not prove a successful ordinary single-player
start or a complete mission transition.

**Unknown.** Native class/inheritance, positive initial construction coverage,
complete selected-owner membership/order across other mutators, field changes
after callbacks, authored owner/map origin, runtime acceptance, uniqueness,
roster size, join schedule and complete transition carryover.

### R2-ENGINE-135

A fresh generator selects 24 new routine pairs by direct-call edges from the
EXP-2016 seeds, in the preregistered order: the literal-reference caller, then
callers of the character producer, owner constructor and actor-list append,
then their callers in ascending EN address. Each EN and RU body is decoded
from its entry through every local branch and jump-table entry. The 48 bodies
hold 16,475 EN and 16,746 RU instructions with no gap bytes. Offsets are
section-derived and each interval carries a SHA256. Normalized shape (mnemonics,
relative local jumps, absolute addresses masked) is equal in 15 pairs and
differs in 9; 5 pairs score below 0.9 under sequence similarity.

**Confidence.** High for decoded bodies, hashes and counts. Medium for the
nine differing pairs: the RU counterpart was located by callee/caller
correspondence and shape similarity, and a shape difference is not explained.
Direct-edge and raw-pointer scans are navigation only.

**Unknown.** Callers that no direct edge reveals, direct-edge functions not
selected under the cap of 24, level-1 seeds outside the priority list, and
indirect callers.

### R2-ENGINE-136

Class records in both images hold name pointer, object size, schema, create
slot and base pointer. The Player record (EN L2.00227 / RU L2.00228) has size
0xa80, base CObject and create slot EN L2.00193 / RU L2.00194. That complete
factory pushes 0xa80, allocates, and calls the owner constructor (EN R2.0079 /
RU R2.0080). The Human record (EN L2.00229 / RU L2.00230) has size 0x254 and
the chain Human > Humanoid > Unit > Token > CObject. Its create slot (EN
L2.00231 / RU L2.00232) pushes 0x254, allocates and calls EN L2.00233 / RU
L2.00234, not the constructor EN L2.00214 / RU L2.00130 that the character
producer calls. The create-slot window is navigation only.

**Confidence.** High for the parsed records and the Player factory join.
Medium that the producer's 0x254-byte object is class Human: equal size and a
different constructor do not identify the class.

**Unknown.** The complete Human create-slot body, the relation of EN L2.00233
to EN L2.00214, whether archive or name lookup instantiates either class, and
the class of the owner collection's elements.

### R2-ENGINE-137

EN L2.00195 / RU L2.00196 pushes 0xa80, allocates and calls the owner
constructor. Its one direct caller is EN L2.00197 / RU L2.00198, whose one
direct caller is the dispatcher EN L2.00171 / RU R2.0119; that dispatcher has
one direct caller, EN L2.00235 / RU L2.00236. EN L2.00199 / RU L2.00200
reserves a 0xa84-byte frame, calls the owner constructor on a frame object and
stores WORD 0x270f in it; it calls no allocator. It has 12 EN direct callers
(11 in EN L2.00197, 1 in the dispatcher) and 10 RU direct callers (all in RU
L2.00198). The factory EN L2.00193 / RU L2.00194 has no direct caller. The
measured owner builder reaches the owner constructor from the map builder.

**Confidence.** High for the complete bodies and direct-edge counts. The EN/RU
difference of two stack-frame callers is measured and unexplained.

**Unknown.** Callers of the factory through the class record or tables,
virtual construction, archive restore, which route builds the ordinary-start
owner, and the stack-frame object's purpose.

### R2-ENGINE-138

In the measured map builder (EN L2.00186 / RU L2.00113) the owner-build call
(EN L2.00237 / RU L2.00238) precedes the actor-build call (EN L2.00239 / RU
L2.00240), which precedes the call to the literal-reference caller (EN
L2.00241 / RU L2.00242). The decoded flow records no exit between the first
and last call. The caller (EN L2.00201 / RU L2.00116) reaches the literal
router at four sites. It has two direct callers: the map builder and EN
L2.00243 / RU R2.0120. No raw pointer to it was found.

**Confidence.** High for the ordered local flow. Medium for consumer
reachability: a callee that never returns, jump-table dispatches inside the
caller (two per locale) and the map builder's own callers are not excluded.

**Unknown.** The map builder's caller and its conditions, router inputs,
whether the owner-build call yields a populated collection, and the second
caller's entry path.

### R2-ENGINE-139

Each locale holds 13 actor-list append call sites (EN L2.00202 / RU L2.00205)
inside the new population. Ten per locale are preceded by a load of a DWORD
from reg+0x24 into the receiver register: EN L2.00244, L2.00245, L2.00246,
L2.00247, L2.00248, L2.00249, L2.00250, L2.00251, L2.00252 and L2.00253, with
RU counterparts. Three are not: EN L2.00254, L2.00255 and L2.00256 (RU
L2.00257, L2.00258 and L2.00259). The displacement is generic; the base
register origin was not traced to an owner at any site.

**Confidence.** Medium for the site population and displacement. Unknown for
ownership of any receiver.

**Unknown.** The base object at each site, appends outside the new population,
and removal or bulk-copy writers.

### R2-ENGINE-140

The 48 bodies hold 109 EN and 113 RU indirect calls, 14 indirect jumps and 14
decoded jump tables per locale. Raw-pointer scans of both images find, outside
.text, the selected entries only at the .data Player create slot (EN L2.00260 /
RU L2.00261), one .rdata cell for EN R2.0084 / RU R2.0085, one for EN L2.00262 / RU
R2.0036 and three for EN L2.00263 / RU L2.00028. No raw pointer to the other
selected entries was found. The scan reads every byte offset of each image and
excludes cells inside .text.

**Confidence.** High for the counted sites, decoded tables and pointer cells
in the two images.

**Unknown.** All indirect-call targets, the owners of the .rdata cells,
callbacks formed by computed addresses, and pointer references inside .text
operands other than direct edges.

### R2-ENGINE-141

The measured sources do not establish the members the ordinary-start owner
collection receives, their order, the startup predicate's value, or successful
literal resolution. The measured frontier names three routes: the map
builder's owner, actor and consumer sequence; a dispatcher route to heap Player
construction; and class-record or table cells that reach Player construction.
No route was shown to be the ordinary start.

**Confidence.** Unknown. High would need the missing caller conditions,
indirect targets and a complete owner-collection mutator population.

**Unknown.** Initial contents and order, member uniqueness and identity, the
startup predicate, the dispatcher's entry conditions, archive or table
construction, callback effects and runtime reachability of the first literal
consumer.

## Campaign departure controller

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-143 | The EN/RU ScenarioLeaveLocation export writes -1 to its output DWORD, returns 0 without further effect when current is null, routes type 2 to a type-2 routine and every other type to an ordinary routine. | High | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-144 | The type-2 departure routine removes the current record available node only for ID 1 (without a null check on the found node) and clears current for every ID; it writes no output and no bank slot. | High | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-145 | The ordinary departure routine is one 2383-byte body with a 12-case ID switch (IDs 10..110) and a 9-case stage switch on bank slot 768 (values 30..110); every other selector takes the join without a case body. | High | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-146 | Ordinary departure adds: 10 mission 20; 20 town 2, plus mission 21 if bank772; 31 mission 32; 40 missions 50, 60; 50 town 3, plus fixed record D2.00015 (type and ID unresolved) if bank780; 60 mission 80; other IDs none. | High | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-147 | The add helpers scan the catalog for matching type and ID, append the record pointer to the available list only when no node already holds that pointer, and the remove and clear helpers unlink or empty that doubly linked list. | High | ● active (amended) | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-148 | The departure output DWORD is -1 by default, 1 for ID 10, 2 for ID 30, 3 for IDs 70 and 80 only under opposite bank777/bank778 gates (the gate skips only the store), 5 or 4 for ID 110 by bank779, unchanged otherwise and for type 2. | High | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-149 | A nonzero bank775 makes ordinary departure clear the available list and append 42 fixed record pointers plus one of four bank776/bank781-selected records before the normal path; the types and IDs behind those pointers are not resolved here. | Medium | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-150 | In the 23-routine EN/RU population no direct callee is unresolved except two unread excluded callees of the node shells, and no indirect edge exists in the controller or list helpers. | Medium | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-159 | Catalog record kind 2 / ID 2 is built by one constructor, D2.00016, calling the record initializer D2.00017 on object D2.00018; it is one of 52 constant-argument initializer call sites in EN and RU scenario.dll. | Medium | ● active | [EXP-2020](../experiments/EXP-2020-rom2-second-town/) |
| R2-ENGINE-160 | The client body that selects per-ID town data has a separate arm for ID 2 beside arms for IDs 1 and 3; the ID 2 arm pushes the resource key music\b16.wav in EN (arm L2.00264) and RU (arm L2.00265). | Medium | ● active | [EXP-2020](../experiments/EXP-2020-rom2-second-town/) |
| R2-ENGINE-161 | The EnterInn export has 10 stage case bodies (8 distinct) and a default; stage 20 and other unlisted stages select the shared continuation D2.00019 (295 instructions), which holds 15 packed-entry stores behind bank-slot compares. | Medium | ● active (amended) | [EXP-2020](../experiments/EXP-2020-rom2-second-town/) |

### R2-ENGINE-143

Export ordinal 6 (ScenarioLeaveLocation, D2.00009..D2.00020, 27 instructions, hash dc27fbe1d668) takes one stack argument, the output pointer, and returns with ret 4. It stores 0xffffffff through the pointer first. If DWORD D2.00007 (current) is zero it returns eax=0. Otherwise it calls the record-type getter (D2.00010, DWORD +0) and compares with 2: equal calls the type-2 routine D2.00013, otherwise the ordinary routine D2.00012, each with the output pointer and the record-ID getter result (D2.00011, DWORD +4). The callee result is returned; the only return-local write in the ordinary routine is its zero initialization, and the type-2 routine returns zero.

A second call after a completed departure finds current zero and has no effect except the -1 store. Types 1, 3 and any other value use the ordinary routine; only type 2 is distinguished.

**Confidence.** High: complete body decoded from the export entry, both callees read to their return. EN and RU bodies are byte-identical (.text and .data hashes equal), so this is one dependency.

**Unknown.** Callers use of the return value, which types exist in the catalog, and whole-client behavior around the call.

### R2-ENGINE-144

D2.00013..D2.00012 (21 instructions, hash f91aaa85a8) compares the ID argument with 1. For ID 1 it finds the node holding current in the available list (D2.00021) and calls the remove routine (D2.00022) on the result without testing it for null; the ordinary routine does test it. In both arms it then stores zero to current. It does not touch the output, the bank or any auxiliary word. For a type-2 ID other than 1 the available list is unchanged.

**Confidence.** High for the control flow. R2-SESSION-023 gives the first-transition consequence for ID 1.

**Unknown.** Whether current can be absent from the available list at a type-2 ID-1 departure; the unguarded remove would then dereference null.

### R2-ENGINE-145

D2.00012..D2.00023 (497 instructions, hash 6dc5fef76f2f) was decoded by recursive descent including both switch tables; a linear sweep leaves zero uncovered bytes and every transfer lands on an instruction boundary. The first dispatch (D2.00024) indexes byte table D2.00025 by ID minus 10 (limit 100) into dword table D2.00023 with 13 targets: cases for IDs 10, 20, 30, 31, 40, 50, 60, 70, 80, 90, 100, 110 plus the join D2.00026 for the other IDs. The second dispatch (D2.00027) indexes byte table D2.00028 by bank slot 768 minus 30 (limit 80) into dword table D2.00029 with 10 targets: cases for 30, 40, 50, 60, 70, 80, 90, 100, 110 plus the return D2.00030. The tables are contiguous D2.00023..D2.00009, ending at the export entry.

**Confidence.** High native decode. Earlier snippets (R2-SESSION-023) are parts of this body.

**Unknown.** Which catalog IDs the campaign data presents; per-case semantics are in R2-ENGINE-146 and R2-ENGINE-148.

### R2-ENGINE-146

Case bodies in branch-matrix.tsv: ID 10 AddMission(20); ID 20 AddTown(2), then AddMission(21) when bank772 is nonzero; ID 31 AddMission(32); ID 40 AddMission(50) then AddMission(60); ID 50 AddTown(3), and when bank780 is nonzero appends the fixed record D2.00015; ID 60 AddMission(80). IDs 30, 70, 80, 90, 100 and 110 contain no add call. The unlisted selectors in 11..109 and any ID outside 10..110 reach the join with no case body. Mission and town departures are not equivalent: the type-2 routine adds nothing (R2-ENGINE-144).

**Confidence.** High for the native conditional additions. The arguments are record IDs; whether a catalog record of that type and ID exists is decided by the helper (R2-ENGINE-147).

**Unknown.** Catalog contents at runtime; a missing catalog record adds nothing.

### R2-ENGINE-147

add-mission (D2.00014) and add-town (D2.00031) are identical except the type compare constant (1 versus 2) and local branch targets. Each walks the catalog list at D2.00032 with the begin and next helpers, and for every record whose type and ID match calls find (D2.00021) with a null start: find walks from the list head and compares each node payload pointer with the record pointer. When none matches, append (D2.00033) adds the record at the tail. The loop continues past a match, so every matching catalog record is considered. Remove (D2.00022) unlinks a node by head, tail and neighbor pointers and returns it to a free chain, calling clear (D2.00034) when the count reaches zero.

**Confidence.** High for the helper bodies. The node-shell routines are read; their callees D2.00035 and D2.00036 are excluded and unread; the role label in exclusions.tsv comes from call context and is not measured.

**Unknown.** The behavior of the two excluded callees and any node-pool aliasing.

**Amended.** R2-ENGINE-218 resolves the two excluded wrappers as size-taking allocation and pointer release paths. Node-pool aliasing remains Unknown.

### R2-ENGINE-148

The output DWORD is stored only by case bodies. ID 10 stores 1. ID 30 stores 2. ID 70 stores 3 only when bank777 is zero and bank778 nonzero; on both paths of that gate it then sets bank777 and bank770. ID 80 stores 3 only when bank778 is zero and bank777 nonzero; on both paths it then sets bank778. The gate decides only the output store. ID 110 stores 5 when bank779 is nonzero, otherwise 4. Every other ID, including every type-2 departure, leaves the -1 written by the export.

**Confidence.** High for the DLL output contract. The client consumer of this DWORD is the one in R2-ENGINE-074 and is not re-read.

**Unknown.** Meaning of values 1..5 beyond that consumer, and how a gated case that stores no value is handled by the caller.

### R2-ENGINE-149

At D2.00037, when bank775 (D2.00038) is nonzero, the routine zeroes it, clears the available list and appends in order the records at D2.00018, D2.00039, D2.00040, D2.00041, D2.00042, D2.00043, D2.00044, D2.00045, D2.00046, D2.00047, D2.00048, D2.00049, D2.00015, D2.00050, D2.00051, D2.00052, D2.00053, D2.00054, D2.00055, D2.00056, D2.00057, D2.00058, D2.00059, D2.00060, D2.00061, D2.00062, D2.00063, D2.00064, D2.00065, D2.00066, D2.00067, D2.00068, D2.00069, D2.00070, D2.00071, D2.00072, D2.00073. It then appends one arm record: D2.00074 (bank776 and bank781 nonzero), D2.00075 (bank776 nonzero, bank781 zero), D2.00076 (bank776 zero, bank781 nonzero) or D2.00077 (both zero). Every arm then appends the five common records D2.00078, D2.00079, D2.00080, D2.00081, D2.00082 at D2.00083 and continues into the normal path, which then removes current. The region holds 46 append calls: 37 fixed, 4 arm records (one per arm) and 5 common, so 42 are unconditional and 43 run per arm.

**Confidence.** Medium: order, conditions and pointers are native. The type and ID of each static record come from separate initializer routines that were not read, because they would exceed the routine cap.

**Unknown.** The type and ID of every appended record and therefore the restored membership by mission or town.

### R2-ENGINE-150

closure.tsv lists every direct call of the 23 routines in each locale: all targets are read routines except D2.00035 and D2.00036, which are excluded and unread. No indirect jump or call remains in the controller, add helpers or list helpers. The population is the EN and RU scenario.dll, which share .text and .data hashes. No claim covers other types, the client consumer, owner saves, runtime execution or any other DLL.

**Confidence.** Medium: the closure is bounded to the named population and is not a global absence statement.

**Unknown.** The behavior of the two excluded callees, other callers of the helpers, and any path that changes current outside this export.

### R2-ENGINE-159

Selector S1 scans scenario.dll .text for calls of the record initializer D2.00017 with constant kind and ID arguments (pushed right to left, object in ecx). It finds 52 call sites in EN and 52 in RU, with identical constructor, kind, ID and object columns. Exactly one carries kind 2 / ID 2: constructor D2.00016 (12 instructions, hash c997266266f1), call site D2.00084, object D2.00018. The other two kind-2 sites are ID 1 (object D2.00085) and ID 3 (object D2.00039). The rebuild routine D2.00005 carries one request for object D2.00018 at catalog position 1. The record initializer (25 instructions, hash a980ca05dde8) calls the record-writer D2.00086 and the helper D2.00087 (8 instructions, hash befca3570939); the writer is a reused span.

Selector controls recover kind 2 / ID 1 and kind 1 / ID 10 and 20 in both locales, and a synthetic interior address D2.00088 is rejected as an instruction start.

**Confidence.** Medium: the call-site scan is complete for direct calls in .text. A constructor reached only through static-initializer registration, a record built at run time with a non-constant ID, or a save-loaded record is not excluded. EN and RU .text and .data are identical, so this is one dependency. Method: the experiment exceeded its preregistered inputs cap (20 reported and 24 conservative against 16), so this claim is bounded by an incomplete method record.

**Unknown.** Constructors outside the D2.00017 call-site population that store the immediates 2 and 2 directly (the preregistered selector population) were not searched. How the constructor is registered and invoked, the body of helper D2.00087 beyond its 8 instructions, and the x and y arguments of the record.

### R2-ENGINE-160

EN body L2.00266 and RU body L2.00267 (the RU decode extends to L2.00268, past the end of the earlier town-consumer range at L2.00269, so it counts as new) enter three arms after a compare chain on the record kind/ID member at record+4: EN L2.00266 calls import cell L2.00270, then getter L2.00271 (returns [ecx+4]) and compares the result with 1, 2 and 3 at L2.00272, L2.00273 and L2.00274; RU L2.00267 uses getter L2.00275 (returns [ecx+4]) and compares at L2.00276, L2.00277 and L2.00278. The arms are EN arms L2.00279, L2.00264 and L2.00280; RU arms L2.00281, L2.00265 and L2.00282. The ID 2 arm pushes its key and does not touch TALK. Each arm pushes one string address and calls a resolver. The arm for ID 2 pushes the key music\b16.wav in both locales. The string cells of arm 1 and arm 3 were not read, so the ID 1 arm is not compared with the ID 2 arm. The client function is shared by IDs 1..3 and diverges at arm level.

**Confidence.** Medium: the arm chain is read in 1 EN and 1 RU body found by the consumer scan. Other consumers of the ID are not excluded, and the compare operands were matched by pattern, not read by hand for every arm. Register-based indirect calls remain unresolved in 14 of the 20 distinct new client bodies, so no decoded client body is a closed control-flow graph. Method: the experiment exceeded its preregistered inputs cap (20 reported and 24 conservative against 16), so this claim is bounded by an incomplete method record.

**Unknown.** The keys of arm 1 and arm 3, the resolver callees, the consumer of member +0x118, and whether any other client body branches on ID 2.

### R2-ENGINE-161

The export D2.00089 (698 instructions, hash a359200ba501) indexes a 101-entry byte table (top 100) and jumps through a dword table; the dispatch is at D2.00090 and the default is D2.00019. Stage labels are table index plus 10, following the earlier stage-10 numbering. Case bodies exist for stages 10, 30, 40, 50, 60, 70, 80, 90, 100 and 110; stages 60, 70 and 80 share one body. Stage 10 holds packed entries kind 0 topic 9 NPC 207, kind 0 topic 8 NPC 2108 and kind 3 topic 10 NPC 517. Stage 30 holds kind 3 topic 30 NPC 22, a compare of bank slot 927 with 0, kind 3 topic 31 NPC 2108 and kind 0 topic 39 NPC 2110. Stages 40 and 50 hold packed entries interleaved with compares of bank slots 937..939 and 533, 947, 949. The rows are in enterinn-cases.tsv; the EN and RU bodies are identical.

No dedicated case body exists for stage 20. The byte table maps stage 20 to the dword-table default, the shared continuation D2.00019. Seven case bodies end with a jump to D2.00019 (D2.00091, D2.00092, D2.00093, D2.00094, D2.00095, D2.00096, D2.00097) and the stage-110 body falls into it. The continuation is 295 instructions to D2.00098 and holds 15 packed-entry stores, each after compares of bank slots against immediates (37 compare rows): kind 0 topic 79 NPC 5 (D2.00099), kind 0 topic 78 NPC 675 (D2.00100), kind 3 NPC 2022 topics 74..77, 84..87 and 93..96, and kind 3 topic 62 NPC 22 (D2.00101). The addresses, entries, compare slots and immediates are in the continuation table of the evidence. Whether a stage-20 call satisfies any of those compares is not measured.

**Confidence.** Medium: the case bodies are decoded by recursive descent with the jump table resolved. The continuation entries and compares are tabulated mechanically; their predicates are not evaluated, and the callees D2.00035, D2.00036 and D2.00011 are not read, so the statement covers packed entries, bank compares and calls, not every effect. Method: the experiment exceeded its preregistered inputs cap (20 reported and 24 conservative against 16), so this claim is bounded by an incomplete method record.

**Unknown.** Which stage value the client passes for catalog ID 2, which continuation entries a stage-20 call reaches at run time, whether the bank-slot compares are gates, the effects of the continuation beyond its stores and compares, and the meaning of the stage labels beyond the earlier numbering.

**Amended.** R2-ENGINE-215/216/217/218 resolve the selector, output-buffer contract, current-ID predicates and named helpers. R2-ENGINE-221 separates the dynamic tail; R2-SESSION-109 bounds the early writer inference. The earlier method grade is unchanged.

## Startup caller and concrete collection receiver

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-151 | The selected population has 21 frontier complete pairs plus one bounded attribution miss, 17 recaptured prior pairs and 14 counted cells including stopping sentinels. | High / Medium | ✔ promoted | [EXP-2019](../experiments/EXP-2019-rom2-startup-collection/) |
| R2-ENGINE-152 | EN L2.00222 / RU L2.00112 calls the map builder only under Session+0x1ac==0; its alternate-load arm is separate, and its local source/alternate/map failure returns are 3, 4 and 5. | High | ✔ promoted | [EXP-2019](../experiments/EXP-2019-rom2-startup-collection/) |
| R2-ENGINE-153 | EN L2.00283 / RU L2.00284 calls the map caller only after initialization returns 0; the other measured caller requires its first argument to equal zero and its receiver EN+0x5d4 / RU+0x638 field to be nonzero, then clears Session+0x1ac. | High / Medium | ✔ promoted | [EXP-2019](../experiments/EXP-2019-rom2-startup-collection/) |
| R2-ENGINE-154 | The nonnull allocation arm constructs an empty global owner collection; new owners are conditionally tail-appended, and fresh owner+0x24 actor collections have their own allocation. | High / Medium | ✔ promoted | [EXP-2019](../experiments/EXP-2019-rom2-startup-collection/) |
| R2-ENGINE-155 | Map actor lookup compares signed row WORD+0x14 to owner DWORD+8 copied from the owner record DWORD+4; a global-head match returns the same first owner selected by the literal helper. | High / Medium | ✔ promoted | [EXP-2019](../experiments/EXP-2019-rom2-startup-collection/) |
| R2-ENGINE-156 | After nonzero positioning return, selected map actor code registers globally, assigns actor+0x14 to the looked-up owner and appends to that owner's+0x24; the global search result does not gate its append. | High / Medium | ✔ promoted | [EXP-2019](../experiments/EXP-2019-rom2-startup-collection/) |
| R2-ENGINE-157 | The alternate literal caller iterates owner+0x24 and appends through the global actor receiver; these are different storage coordinates, with runtime pointer aliasing unproved. | High / Medium | ✔ promoted | [EXP-2019](../experiments/EXP-2019-rom2-startup-collection/) |
| R2-ENGINE-158 | Ordinary-start members, first accepted row, total order, uniqueness, final numeric identity and literal success remain Unknown at the measured caller, lookup-next, world-result and callback boundaries. | Unknown | ✔ promoted | [EXP-2019](../experiments/EXP-2019-rom2-startup-collection/) |

### R2-ENGINE-151

The fixed population has 22 new paired decoded intervals and 17 reused pairs.
Each locale has 3,535 new and 2,291 reused decoded instructions with zero interval
gap bytes. Twenty of the 21 frontier pairs and the miss match normalized shape. The differing outward-low pair
uses EN receiver+0x5d4 and RU receiver+0x638; containing-object equivalence is
not established. The 21 frontier pairs have complete direct branch/return
coverage. The extra attribution-miss pair is a bounded interval whose compact
byte selector is not enumerated. The stopping DWORD sentinel reads its first
four bytes, without resolving the dispatch population.

Nearest-container navigation misassigned call L2.00285 / L2.00286 to a preceding
body ending before that site. Complete L2.00283 / L2.00284 contains the call.
The miss is preserved, counted and excluded from caller conclusions. Its five
jump targets and stopping sentinel per locale plus two conditional Human
vptr+0x58 cells yield 14 conservative new cell reads. The 12 automatic miss reads
deviate from the preregistered receiver-guided explicit cell selection; no
population was expanded. All 39 pairs hold 25 indirect calls, two indirect
jumps and two jump-table notes per locale. Direct-edge scans cover raw-size .text in each
image, not virtual/archive/table/computed callers.

**Confidence.** High for measured bytes, counts, site containment and exact
local instructions. Medium for differing locale objects and inferred body
attribution. A decoded compact-table interval does not close its selector or
prove a caller. Generator, Capstone 5.0.7, pefile 2024.8.26, Python 3.12.10 and
source/input/interval hashes are recorded in the evidence tables.

**Unknown.** The miss's selector population, caller paths outside the selected
bodies, indirect effects and runtime reachability. No completeness or global
absence follows from the linear section scan.

### R2-ENGINE-152

Complete map caller EN L2.00222..L2.00287 / RU L2.00112..L2.00288 tests
Session+0x1ac. Zero selects map build L2.00186 / L2.00113; nonzero selects
L2.00289 / R2.0121. The map call receives ECX=Session+0x98, after an
unexpanded helper prepares an argument using Session+0x90. In the zero Session+0xd4 branch, source lookup returning -1 exits
with 3 before map build. A nonzero alternate-load return exits with 4. A nonzero
map-builder return exits with 5; the accepted local continuation returns 0.
These are local return codes, not inferred UI messages or ordinary-start modes.

**Confidence.** High for all decoded direct branches, returns and caller
receiver coordinates. Both complete bodies have 332 instructions and identical
normalized shape. Each local alternative is present in the measured body.

**Unknown.** Entry predicates' runtime values, source lookup/content, alternate
load effects, callees returning and whether this is the owner's ordinary start.

### R2-ENGINE-153

Verified wrapper EN L2.00283..L2.00221 / RU L2.00284..R2.0010 passes argument 2
to initialization L2.00221 / R2.0010. Nonzero initialization returns without
calling the map caller. The zero arm prepares an argument through unexpanded
L2.00290 / L2.00291 using the address of its first stack argument, reloads the
saved receiver and calls L2.00222 / L2.00112. It returns that call result after
cleanup helpers return. The other complete outward caller
R2.0026 / R2.0023 requires argument 1==0 and a nonzero field in its receiver
(EN+0x5d4 / RU+0x638), clears the global Session+0x1ac and calls the map caller
with the global Session receiver after separate helper-mediated argument
preparation. This caller does not forward its zero flag as the map argument.
Its 1002 instructions per locale differ in some field coordinates.

**Confidence.** High for the complete local gates, same-receiver wrapper and
clear/call instructions. Medium for EN/RU containing-object correspondence;
matching call roles do not prove equivalent layouts or player-visible mode.

**Unknown.** Both outer entries, object classes, field meanings and actual
argument/field values, argument-preparation and cleanup helper effects. No
composed ordinary-start call is asserted.

### R2-ENGINE-154

The reused initializer requests 0x24 bytes. Its nonnull allocation arm constructs
and stores the global owner collection at EN L2.00292 / RU L2.00293; its null
allocation arm stores zero at that coordinate. New complete chain L2.00294 -> L2.00295 -> L2.00296
and L2.00297 -> L2.00298 -> L2.00299 clears head+4, tail+8 and count+0x0c; the
outer constructor sets key counter+0x20=1. Fresh Player construction separately
requests its actor collection and stores the constructor result or zero at
owner+0x24. It clears owner+0x38; successful actor-list construction uses the
complete reused list initialization in R2-ENGINE-123.

Owner registration R2.0082 / R2.0083 calls R2.0092 -> L2.00300 -> L2.00301
and R2.0093 -> L2.00302 -> L2.00303. Its complete node-zero helpers L2.00304 /
L2.00305 initialize the payload; node allocation assigns next/previous and
increments count. Append writes node+8, links old tail.next or empty head,
sets tail and returns the node. R2-ENGINE-125's reuse arm remains separate.

**Confidence.** High for local field clears, distinct collection coordinates,
complete link/return stores and conditional tail order on valid allocated lists.
Medium for preserving that state through unexpanded external helpers and later
code. Neither allocation nor local empty initialization is a positive member.

**Unknown.** Allocation-failure recovery, global owner members before this route,
intervening callbacks, all other mutators and ordinary-start entering state.

### R2-ENGINE-155

The reused owner builder copies map owner record DWORD+4 to owner DWORD+8 at
EN L2.00306 / RU L2.00307 on its new-owner branch before registration. Actor
build reads actor row WORD+0x14 and calls L2.00308 / L2.00309 on the same global
owner collection. The complete lookup starts with L2.00310 / L2.00311, the
same head helper that selects the literal consumer's owner. The reused
selected-owner getters test collection DWORD+0x0c: zero returns null; otherwise
head node+8 supplies the payload (R2-ENGINE-123, R2-ENGINE-125). It compares a
sign-extended input WORD to current owner DWORD+8 and returns a matching owner.
The runtime owner WORD+4 is not this comparison. A first-head match proves the
returned pointer equals the consumer's locally selected first owner.

**Confidence.** High for the same global receiver, copied field and complete
lookup's compare/head-match return. Medium for head and record persistence
through intervening calls and composition over later owners. The complete
lookup-next wrapper has unexpanded search/step/payload callees.

**Unknown.** Actual row values, duplicate owner DWORD+8 values, loaded/reused
owner field origin, complete mismatch traversal effects and runtime head match.

### R2-ENGINE-156

In reused actor build, the candidate insertion requires a nonnull world, row
and looked-up owner, `(Session DWORD+0x174==0 || owner DWORD+0x2c==0)`,
actor WORD+0x0e!=0 and
positioning return!=0. The complete positioning L2.00312 / L2.00313 returns 0
when its acceptance flag remains 0; its accepted branch returns unexpanded
L2.00314 / L2.00315. A virtual+0x58 call precedes it. Constructor-guided cells
EN L2.00316 / RU L2.00317 point into .text at L2.00318 / L2.00052; complete target
bodies have external callees, and actual vptr persistence is unproved.

After the nonzero position return, actor DWORD+8 receives signed row WORD+0x3c.
Complete global registration L2.00319 / L2.00320 searches for the supplied
pointer, ignores that return, tail-appends and writes actor DWORD+4 as+8+0x6000.
Search L2.00321 / L2.00322 walks nodes; complete L2.00323 / L2.00324 compares
payload pointers. No result-dependent branch suppresses that append. The
caller then stores actor+0x14=the looked-up owner and calls the complete reused
append L2.00202 / L2.00205 with that owner's+0x24. R2-ENGINE-123 supplies its
full append/node chain and tail-order contract without deduplication.

**Confidence.** High for the local receiver/payload joins, complete conditional
branch to insertion, search return use and bounded append stores/returns.
Medium for same-object preservation through constructors, virtual/external
calls and acceptance of any actual row. A conditional delivery is not proof
of the ordinary first insertion or complete uniqueness/order.

**Unknown.** Actual first accepted row, world-result helper, actor vptr persistence,
external allocation/node helpers, other mutations, global/owner-list runtime aliasing, class and final WORD+0x14c.

### R2-ENGINE-157

Complete alternate caller L2.00243 / R2.0120 loads each owner's+0x24 for an
iterator and appends a payload only when `(actor BYTE+0x4c & 0x08)==0`. Its
append receiver is the global actor collection EN L2.00325 / RU L2.00326 at
L2.00256 / L2.00259. Source iteration instead loads owner+0x24. These distinct
receiver storage coordinates do not prove equality or distinctness of their
runtime pointer values through archive/virtual loaders. The later
literal-reference caller remains conditional on the archive/load branch and
prior indirect calls. Matching helper addresses do not prove receiver aliasing.

**Confidence.** High for local source/receiver coordinates, bit test and selected
append. Medium for construction and survival through its indirect loaders;
those effects and caller inputs remain unresolved.

**Unknown.** Initial source members, archive/table creation, callbacks, mutation
of either receiver, runtime pointer aliasing and reachability of the literal caller.

### R2-ENGINE-158

The selected evidence adds concrete map caller predicates, verified outer
wrapper conditions, independent owner/actor collection allocation, a first-head
owner key match, conditional owner+0x24 pointer delivery and the alternate's
global receiver storage coordinate. It still does not establish an ordinary-start native
state or a successful first literal consumer. The compact-selector miss,
outer runtime fields, lookup-next callee effects, positioning-world result and
virtual/constructor callback persistence are explicit boundaries.

**Confidence.** Unknown for positive ordinary-start members and runtime consumer
success. R2-ENGINE-152..157's local contracts do not exclude every intervening
mutation or prove any live caller/input condition.

**Unknown.** First accepted row, complete contents/order, unique pointers and
numeric identities, native class, final WORD+0x14c, owner+0x38 population, loaded
owner/actor construction, receiver aliasing, callback effects and actual literal
resolution.

## Selected initialization receiver and field boundary

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-167 | Selected initialization's explicit indexed Session stores cover DWORD starts +0xfc..+0x16c; its 197-instruction body has no memory displacement +0x174. | High | ● active | [EXP-2021](../experiments/EXP-2021-rom2-startup-gate/) |
| R2-ENGINE-168 | First selected initialization callee receives unchanged entry ECX but never explicitly uses it; it passes a fixed global buffer, zero and 0xc00 to an unread callee, then stores global DWORD 1. | High | ● active | [EXP-2021](../experiments/EXP-2021-rom2-startup-gate/) |

### R2-ENGINE-167

Reused initialization EN L2.00221..L2.00222 / RU R2.0010..L2.00112 has 959 bytes and 197 reachable instructions in each locale. The entry receiver is saved at frame offset -0x150. Its explicit indexed receiver store uses base +0xf8, scale four and DWORD width. A local counter initialized to one advances by one and exits above 29. The DWORD starts are therefore +0xfc..+0x16c and the last stored byte is +0x16f. An explicit DWORD store `[receiver+0xf8]=0` through the reloaded saved receiver precedes the loop at EN L2.00327 / RU L2.00328, so the covered bytes are +0xf8..+0x16f. No decoded memory operand in either selected body has displacement +0x174.

**Confidence.** High for the complete-body displacement observation and explicit indexed-store interval. Published reuse hashes match. The two normalized bodies differ at one instruction, a locale-specific global array displacement in the word store EN L2.00329 / RU L2.00330; the displacement observation is decoded in each locale separately. This is a bounded encoding/coordinate result, not an absence of all possible +0x174 mutations.

**Unknown.** Other writes can use loaded pointers or opaque calls. In particular the store through [Session+0x7c]+0x0c has no proved runtime alias relation to Session+0x174. The saved receiver is also stored to a fixed global DWORD (EN L2.00331 / RU L2.00332); that global's relation to the Session is unproved. Receiver persistence and actual field values remain Unknown.

### R2-ENGINE-168

The saved-entry receiver is unchanged in ECX at initialization call EN L2.00333 / RU L2.00334. The selected targets EN L2.00335..L2.00336 / RU L2.00337..L2.00338 are complete 35-byte, 10-instruction bodies with equal normalized shape and no explicit ECX operand. They push 0xc00, zero and fixed global address EN L2.00339 / RU L2.00340, call unread EN L2.00341 / RU L2.00342, then store one to that global DWORD and return. Each locale reserved a body before its frozen 8192-byte window read; counterpart proof charges one new pair.

**Confidence.** High for the exact entry register, explicit arguments, global store and complete local return. The experiment uses eight native input keys including whole-file dependency hashes, PE headers, reused bodies and new body windows. No speculative callee or cell population was read.

**Unknown.** Unchanged ECX at a call does not prove that the callee consumes it as a receiver. The opaque callee's calling convention, ECX use, effects and return are unread. Global/Session aliasing and any composed +0x174 value remain Unknown. The author stopped the queue after this target and left two eligible sites per locale unexpanded: this target's own call EN L2.00343 / RU L2.00344 with unchanged ECX, and the initialization call EN L2.00345 -> L2.00346 / RU L2.00347 -> L2.00348 with receiver Session+0x78 (symbolic saved entry receiver plus 0x78). Their callees are unread.

## Selected RU Player-to-Group programme

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-191 | The selected RU Player list calls Group inline; two embedded programmes precede its member path and three final scalar calls, with selected binding/extents narrowed by R2-ENGINE-199 and R2-ENGINE-200. | Medium | ✔ promoted (amended) | [EXP-2024](../experiments/EXP-2024-rom2-group-serialization/) |
| R2-ENGINE-192 | The selected RU Group member writer calls the published u32 helper for its count, then repeats a call to an unread member helper; encoding and membership remain unproved. | Medium | ✔ promoted | [EXP-2024](../experiments/EXP-2024-rom2-group-serialization/) |

### R2-ENGINE-191

The RU Player body R2.0085 calls list R2.0122, which calls Group body
R2.0123 directly in both directions. The list reader allocates 80 runtime
bytes and calls constructor L2.00349 before that direct Group call. This is
inline dispatch at the Group layer, not a newly established archive wrapper.
The separate constructor body read goes beyond the prereg's literal
serialization/list/count/reference admission. It is retained as a charged
construction-read deviation, not serializer provenance or support for this
claim. Its embedded +0x20 child constructor L2.00350 remains unread.

Group first calls virtual slot +8 on embedded receiver Group+0x20, then
L2.00351 on the pointer at Group+0x3c. That helper requests 80 bytes through
raw-transfer helpers and invokes virtual slot +8 on its own pointer at +0x4c.
Group then invokes the member path and the same scalar helpers used for
published Player u32 fields, for runtime +0x1c, +0x40 and +0x44. The two
embedded virtual targets and raw-transfer helper bodies were not read.

**Confidence.** Medium for this partial programme. Complete directed CFG reads
and two fresh replays establish the local call order and selected operands.
Published RU scope supplies locator identity. Unread embedded programmes
prevent a complete Group grammar or extent claim. The 80-byte value is a
native call request, not new proof of the raw helper's transfer semantics.
No image-wide enumeration, locale pairing or absence claim is made.

**Unknown.** Embedded classes, extents, field meanings, reference repair,
constructor internals beyond the selected body, and the EN counterparts.

**Amended.** R2-ENGINE-199 and R2-ENGINE-200 extend the earlier partial
population: selected construction paths bind both embedded +8 calls and
establish conditional counted-programme extents. The original
constructor-read deviation and then-unread clauses above remain historical
facts of that population. Other runtime classes/vptr changes, field
meanings, reference/alias effects through unread callees, complete Group
grammar and EN counterparts remain Unknown. Grade and inline dispatch facts
are unchanged.

### R2-ENGINE-192

Group writer site L2.00352 calls L2.00353. That complete body calls L2.00354
at L2.00355 for its count, then loops through cursor operations and calls
L2.00356 at L2.00357 with the archive and the selected item pointer.
L2.00354 is the u32 helper already used by R2-SESSION-020. The Group load path
reads its own count at L2.00358 and calls L2.00359 at L2.00360 in its inline
loop, followed by unread append helper L2.00361. It does not call L2.00353.
The latter body's own reader branch calls L2.00359 at L2.00362 and uses
L2.00207; that branch is not reached by this Group load path. The reference,
append and body helpers remain unread. No actor body was read.

**Confidence.** Medium for the bounded member path. Native operands and
complete local CFGs establish the calls and reuse of the published scalar
helper. Unread cursor and reference helpers leave encoding, valid class,
alias behavior and complete list membership open.

**Unknown.** Null, class reuse, object alias, body and member validation arms;
actor grammar; current-player, hero and party identity; complete LOAD.

## Selected RU embedded Group programmes

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-199 | The selected RU construction paths bind both embedded Group +8 calls to one counted two-byte-element programme; other runtime class targets remain unproved. | Medium | ✔ promoted | [EXP-2025](../experiments/EXP-2025-rom2-group-embedded-programmes/) |
| R2-ENGINE-200 | The selected RU embedded programme stores an escaped count and two bytes per visited node; load uses decoded count, while count/chain equality and complete transfer remain unproved. | Medium | ✔ promoted | [EXP-2025](../experiments/EXP-2025-rom2-group-embedded-programmes/) |

### R2-ENGINE-199

The Group constructor L2.00349 calls L2.00350 on Group+0x20 at L2.00535.
It constructs the Group+0x3c receiver through L2.00500 at L2.00536.
The latter constructs its +0x4c pointer through L2.00350 at L2.00537;
field-helper load reconstructs that pointer through the same constructor
at L2.00538. Constructor L2.00350 writes vtable L2.00539 at L2.00540.
Only its selected +8 cell L2.00541 was read; it points to L2.00542.
The Group+0x20 and field+0x3c.pointer+0x4c +8 dispatches use this cell
under these selected construction paths. No class descriptor or other cell
was selected. Constructor reads establish binding, not byte widths.

**Confidence.** Medium for constructor-selected binding. Direct caller
operands, the final constructor store and exact selected cell reproduce in
both fresh runs. No image-wide class enumeration, runtime-vptr invariant
or complete ordinary-save reachability was established.

**Unknown.** Other classes or vptr mutations, allocation failure, old-object
virtual +4 destruction, and complete lifecycle/runtime reachability.

### R2-ENGINE-200

Programme L2.00542 calls empty local base serializer L2.00543. Its count
writer L2.00544 emits u16 n when n<65535; otherwise it emits u16 65535
then u32 n. Reader L2.00545 uses the same escape. Selected scalar helpers
store/load 2 or 4 bytes and advance their buffer cursors by that width.
Element helper L2.00546 requests twice its element argument; both selected
programme branches pass 1, giving two requested bytes per operation.

Let c(n) be 2 for n<65535 and 6 otherwise. Load loops decoded count n,
so its extent is c(n)+2*n after complete successful transfers and normal
helper returns. Store writes receiver+0x0c count n, then follows the linked
chain until null: its extent is c(n)+2*L for visited node count L. This body
checks neither n==L nor complete chain consistency.

Between the two embedded programmes, field helper L2.00351 requests 80 raw
bytes. Writer L2.00547 partitions that request; reader L2.00548 can return
less than requested. The field helper ignores that return. Buffer spill,
refill and file-transport virtual targets remain unread. The boundary model
requires complete successful transfer; it proves no native acceptance.

The reached local L2.00542 CFG shows no explicit reference/alias transfer.
Its unread append call L2.00549 receives no archive operand at this caller;
that does not exclude archive/global effects in its body. Unread lifecycle,
transport and refill paths remain alternatives. The unread L2.00550 call
between field-helper return and member-count read has the same limitation.

**Confidence.** Medium for the selected programme and conditional extents.
Complete local CFGs, primitive operand widths and cursor increments exclude
a fixed embedded extent on the decoded-count load path. Constructor/class,
count/chain and transport alternatives remain open. Synthetic and reporting
controls are instrument sanity checks and add no native evidence.

**Unknown.** Element meanings/namespaces, n versus L consistency, short-read
and invalid-input behavior, transport/refill semantics and original LOAD.


## Selected RU first Group member programme

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-207 | The selected RU archive object reference is a u16 tag with u32 escape: null, earlier object, earlier class or new class with schema and name; class and object share one index counter before the payload call. | High / Medium | ✔ promoted | [EXP-2026](../experiments/EXP-2026-rom2-member-programme/) |
| R2-ENGINE-208 | The RU Human class record chains to Humanoid and Unit; Human adds no stream bytes, and Humanoid writes the Unit programme, 24 raw bytes and twelve Item references. | High | ✔ promoted | [EXP-2026](../experiments/EXP-2026-rom2-member-programme/) |
| R2-ENGINE-209 | The RU Unit serializer writes a 38-byte base, counted references, three counted word lists, six raw blocks of 498 bytes, scalars, three references, a name and two gated nested payloads in one fixed order. | High / Medium | ✔ promoted | [EXP-2026](../experiments/EXP-2026-rom2-member-programme/) |
| R2-ENGINE-210 | The RU Item serializer writes the shared 38-byte base, counted references and eight scalars; Weapon adds 24 and 22 raw bytes, a byte and a reference, and Armor adds 22 raw bytes and a byte. | High | ✔ promoted | [EXP-2026](../experiments/EXP-2026-rom2-member-programme/) |

### R2-ENGINE-207

Object load L2.01148 calls tag reader L2.01149; store L2.01150 calls
L2.01151. The tag is a u16 w. When w is 0x7fff, a u32 follows and is the
tag; otherwise the tag is (w & 0x7fff) | ((w & 0x8000) << 16).

- High bit clear: an object index. 0 is null. Load checks the index against
  the map size and checks the stored object against the expected class.
  No payload follows.
- w = 0xffff: a new class. New-class reader L2.01152 reads u16 schema,
  u16 name length and the name bytes; length 64 or more is refused. It finds
  the registered class record by exact name through imported lstrcmpA.
  The writer measures the name through imported lstrlenA.
- Other high-bit tags: an earlier class index (index | 0x80000000).

A new class enters the shared map at counter archive+0x30, then the new
object enters at the next counter value, then the object's virtual +8
serializer runs. An earlier class creates a new object the same way.
Load checks the actual class against the caller's expected class record
through L2.01153. A saved schema different from the class record's schema
is accepted only when the record's schema word has its top bit set;
otherwise the reader reports error 7.

A class record is 24 bytes: name pointer, object size, schema, factory,
base record and one more pointer, which is 0 in every record read.

**Confidence.** High for the stream grammar: the complete load CFGs of
L2.01148 and L2.01149 were read, and the four frozen saves exercise the
null, new-class and earlier-class forms at exact predicted offsets
(R2-SESSION-099, R2-SESSION-100). Medium for runtime acceptance: error
reporter L2.01154, the schema-mapping store and the registry walk inside
L2.01152 beyond the name compare were not followed. The earlier-object and
u32-escape forms were read but not exercised by any save.

**Unknown.** Registry insertion order and duplicate names; the effect of
an error report; the 0x7fff escape in real saves.

### R2-ENGINE-208

Class record L2.00230 names Human: object size 0x254, schema 1, factory
L2.00232, base record L2.01155. Record L2.01155 names Humanoid: size 0x254,
schema 1, base L2.01156, the Unit record. The factory allocates 0x254 bytes
and calls constructor L2.00234, which first calls L2.01157 (unread) and
then stores vtable L2.00218. Slot +8 is L2.01158.

The Human serializer L2.01158 calls Humanoid serializer L2.01159 and adds
no stream bytes. On load it stores in +0x3c the result of the unread
method L2.01160 on object L2.00149, passed u16 +0x0c when +0x0e is below
0x21 and 5 otherwise. That call receives no archive pointer.

Humanoid L2.01159 calls Unit serializer L2.01161, transfers 24 raw bytes
at +0x23c, then references at +0x208+4*i for i = 1..12. Slot 0 is not
written. Load reads each reference through L2.01162, whose expected class
record is L2.01163 (Item).

**Confidence.** High. The record bytes, factory, constructor store,
vtable cell and both complete serializer CFGs were read. The Human name
search over the RU .data section found one string, and the pointer search
over the same section found one record pointer. Pointers in other sections,
in instruction operands or in records built at runtime were not searched;
the byte fit rules out a different programme for this name.
The four saves place the following Unit and Humanoid fields with exact
class-name alignment (R2-SESSION-100).

**Unknown.** The meaning of the 24 raw bytes and of the twelve slots; the
result of L2.01160; the effects of L2.01157 and the rest of the Human
constructor.

### R2-ENGINE-209

Unit serializer L2.01161 transfers, on load and store in the same order:

1. Base L2.01164: raw 12 bytes at pointer +0x10, then u32 +4, u16 +0xc,
   u16 +0xe, u32 +8, u16 +0x18, u32 +0x1c, u32 saved object address,
   u32 +0x14: 38 bytes.
2. +0x20: u32 count, then that many references (expected record L2.01165).
3. Two counted word lists at +0x1c8 and +0x1e4: count u16, or 0xffff and
   u32, then 2 bytes per element (R2-ENGINE-200 programme).
4. Raw 24 (+0xa6), 22 (+0xbe), 24 (+0x114), 64 (+0xd4), 180 (pointer
   +0x1c0), 184 (pointer +0x1c4) bytes, then a counted word list rebuilt on
   load for the +0x1c4 object.
5. u8 +0x49, +0x4a, +0x4b, +0x4c; raw 4 at +0x50, +0x54, +0x58; u8 +0x60,
   +0x61, +0x6c.
6. References +0x74 and +0x78 (expected Item L2.01163).
7. A counted string at +0x80.
8. Fourteen u16 at +0x84..+0x9e; u8 +0xa2, +0xa3; u16 +0xa0, +0xa4.
9. u8 +0x12c; u32 +0x130; u8 +0x134, +0x135, +0x136; u32 +0x138; u8 +0x13c.
10. u32 packing u16 +0x14c, flag +0x204 (bits 16..23) and flag +0x1a0
    (bits 24..31); u32 +0x144.
11. Reference +0x68 (expected Item).
12. u8 gate; when nonzero, the +0x7c object: u32 count, that many
    references, u32 +0x1c, u32 +0x20.
13. u8 gate; when nonzero, load allocates a 0x1c-byte object and reads
    u32 +0x18, u32 size n, then references for indices 1..n-1.
14. u32 +0x5c, +0x64, +0x44, +0x40; u8 +0x48.

Store writes each gate as 1 when the pointer is nonnull, else 0.

**Confidence.** High for order and widths on the load branch: complete
CFGs of L2.01161, its base, raw, count and reference helpers, and both
gated helpers L2.01166/L2.01167. The unread callees in these bodies,
L2.00338, L2.01168 and L2.01169 in the base, the Unit virtual +0x30 call
and the +0x3c setter L2.01170 on L2.00147, receive no archive pointer, so
they cannot read or write the stream. The four saves fit this programme
with exact class names at predicted offsets (R2-SESSION-099). Medium for
step 13 with n > 0: the resize call L2.01171 and element accessor
L2.01172 are unread and no save exercises it.

**Unknown.** Field meanings; nonempty lists and gate-13 contents in real
saves; transport short reads and refill.

### R2-ENGINE-210

Class record L2.01163 names Item: size 0x58, schema 1, factory L2.01173,
base L2.01174. Weapon record L2.01175 (size 0x8c) and Armor record
L2.01176 (size 0x70) have schema 1 and base L2.01163.

Item serializer L2.01177: first the 38-byte base programme L2.01164 that
Unit and Item both call first, then +0x20 counted
references, then u16 +0x40, u16 +0x42, u8 +0x44, +0x45, +0x46, u16 +0x48,
u16 +0x4a, u8 +0x47: 54 bytes when the count is 0.

Weapon serializer L2.01178 (vtable L2.01179 slot +8): Item programme, raw
24 at +0x5a, raw 22 at +0x72, u8 +0x58, reference +0x88 (load expects
record L2.01180). Armor serializer L2.01181 (vtable L2.01182 slot +8):
Item programme, raw 22 at +0x5a, u8 +0x58. On load, Item stores in +0x3c
the result of unread L2.01183 on object L2.01184 when unread L2.01185
returns more than u16 +0x0c, otherwise 0, and Weapon and Armor the
result of unread L2.01186 on L2.01187 and L2.01188, each passed u16 +0x0c;
none of these calls receives an archive pointer.

**Confidence.** High. Names come from exact searches over the RU .data
section: Weapon has four string hits and Armor five; for each, exactly one
hit has a record pointer in the .data section. Pointers in other sections,
in instruction operands or in records built at runtime were not searched;
the byte fit rules out a different programme for these names. Factories,
constructors, vtable cells and
serializer CFGs were read completely. The saves fit 103-byte Weapon and
77-byte Armor payloads.

**Unknown.** Item field meanings; other Item subclasses; the L2.01180
class.


## Inn option contract and town 2 TALK

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-215 | EnterInn selects bank slot 768 for every catalog record; all four measured EN/RU client call sites pass an option-array pointer and count pointer, without a scalar stage argument. | High | ✔ promoted | [EXP-2027](../experiments/EXP-2027-rom2-inn-entries/) |
| R2-ENGINE-216 | The stage-30 EnterInn body returns NPC22/topic30 kind3, NPC2108/topic31 kind3 only when slot927=0, and NPC2110/topic39 kind0; its stores fill the caller's option buffer, not the catalog. | High | ✔ promoted | [EXP-2027](../experiments/EXP-2027-rom2-inn-entries/) |
| R2-ENGINE-217 | The 15 fixed EnterInn continuation stores depend on current-record ID and exact bank predicates; three NPC2022 families have no stage bound and use slots776/781 to select one topic each. | High | ✔ promoted | [EXP-2027](../experiments/EXP-2027-rom2-inn-entries/) |
| R2-ENGINE-218 | D2.00011 reads record ID; D2.00035/00036 allocate/release list-node blocks through selected heap paths and do not decode packed options; EnterInn calls only the ID getter. | High | ✔ promoted | [EXP-2027](../experiments/EXP-2027-rom2-inn-entries/) |
| R2-ENGINE-219 | Selected town2 stage30 kind3 TALK admits type1/ID-topic to availability; NPC22/topic30 also sets slots533=1,553=2, while NPC2108/topic31 has no special bank store and NPC2110/topic39 kind0 makes no state change. | High | ✔ promoted | [EXP-2027](../experiments/EXP-2027-rom2-inn-entries/) |
| R2-ENGINE-220 | The selected EN/RU inn builder lists option kinds0/3 as talk actors; its TALK action finds the first matching low16 actor key, formats npc%dtalk%d, dispatches text, then passes the same packed word to TalkTo. | High / Medium | ✔ promoted | [EXP-2027](../experiments/EXP-2027-rom2-inn-entries/) |
| R2-ENGINE-221 | The EnterInn dynamic tail emits kind1/2, low ID i+1 when slot552+i=currentID and slot532+i is 1/2, with bit31 iff signed slot512+i>0; an admitted currentID2 EnterInn call after Leave20 without an intervening writer emits kind1/ID1. | High | ✔ promoted | [EXP-2027](../experiments/EXP-2027-rom2-inn-entries/) |

### R2-ENGINE-215

EnterInn D2.00089 uses parameter+8 as an array pointer and parameter+0xc
as a count pointer, zeroes the count at D2.00122, reads bank768 at
D2.00123 and returns with RET8 at D2.00098. The selector is unrelated
to the current catalog ID. EnterLocation stores current and clears
752..767 without a stage store.

Ordinal9 binds the DLL export to EN cellL2.00409 and RU cellL2.00551.
The four EN calls areL2.00399,L2.00400,L2.00401,L2.00402; their RU
counterparts areL2.00410,L2.00411,L2.00412,L2.00413. Each pushes two
stack addresses. The first array/count pair is frame-0x94/frame-0x14;
the other three use frame-0xb8/frame-0x38. Each local call gate compares
EN member+0x5d8 or RU member+0x63c with 2 and bypasses the setup on inequality.

**Confidence.** High for the positive selector, ABI, bindings and eight
local call setups. PE section translation, complete DLL decode and
linear/recursive instruction boundaries exclude a scalar stage argument
in this contract. EN/RU DLL equality is one dependency. This is not a
whole-image call-site enumeration or an all-path client dispatch proof.

**Unknown.** The mode member's complete producer set, the first caller's
unresolved indirect switch, intervening aliases and runtime admission.

### R2-ENGINE-216

The stage30 body has three stores: kind3/topic30/NPC22 atD2.00124;
kind3/topic31/NPC2108 atD2.00125; kind0/topic39/NPC2110 atD2.00126.
The compare atD2.00127 reads slot927 against0. JNED2.00128 bypasses the
middle store; the first and third stores have no additional bank gate.
Every selected store writes array[count] and increments the pointed count.
R2-ENGINE-219 identifies the subsequent TALK effects.

**Confidence.** High for this entire case and its gate polarity. The
native walk exercises slot927=0,1,-1 in both DLLs. The complete EnterInn
body calls only the read-only ID getter and writes its output/stack locals;
the measured client arguments are stack buffers. Arbitrary aliased caller
arguments are outside this contract.

### R2-ENGINE-217

The continuation atD2.00019 contains fifteen constant packed stores before
its dynamic loop. Topic79/NPC5 kind0 requires currentID3, signed768>60
and 774=1. Topic78/NPC675 kind0 requires currentID3, signed768>60 and 770=1.

All twelve NPC2022 options have kind3 and require currentID2. Topics74..77
require959!=0 and every970..973=0. Topics84..87 require any970..973!=0
and every980..983=0. Topics93..96 require any980..983!=0 and every989..992=0.
The families have no stage compare. For each family, slots776/781 select
topics as follows: zero/zero gives74,84,93; zero/nonzero gives75,85,94;
nonzero/nonzero gives76,86,95; nonzero/zero gives77,87,96.

Topic62/NPC22 kind3 requires currentID2, signed768>=60,958=0 and 949!=0.
R2-ENGINE-216 states the output-buffer contract. Kind3 subsequently requests
available type1/ID-topic; kind0 adds no catalog record. The exact individual
stores and bank-compare branches are regenerable measurements.

**Confidence.** High for this bounded constant-store population. Native
branch polarity is established by direct decoding. The 2,046 static cases
check the other gates, currentID, each family blocker, flag arms, signed
stage boundaries and the separate stage30 gate against an independent
predicate expression. This control set fixes slots774/770 at 1, leaving both
gates' false branches uncovered. One focused test covers the false branch
of the slot770 gate; no test covers the false branch of the slot774 gate.
The whole
EnterInn instruction set has complete linear agreement. No global absence
or runtime state claim follows.

**Unknown.** Actual bank predicates for every campaign visit. The native
stage20/currentID2/959=1 counterexample admits topic74.

### R2-ENGINE-218

D2.00011 atD2.00011 returns DWORD[this+4], the record ID. The complete
EnterInn body has fifteen direct calls, all to this getter, and no
unresolved indirect edge.

D2.00035 atD2.00035 takes a byte size, callsD2.00129 and can invoke an
allocation-failure callback before retrying. The selected allocation path
reachesD2.00130/D2.00131 and the imported HeapAlloc cellD2.00132.
D2.00036 atD2.00036 forwards its pointer toD2.00133; the selected release
path reaches imported HeapFree cellD2.00134 or small-block helpers.

Node-block allocationD2.00135 requests 4+nodeCount*nodeSize bytes; the
measured node-acquire caller supplies nodeSize12. A four-byte header links
blocks. ReleaseD2.00136 follows that header chain and passes each block
pointer to D2.00036. List append stores the catalog pointer in node+8.
These wrappers process sizes and storage pointers, not packed-option fields.

**Confidence.** High for the selected local contracts and positive heap
paths. Complete wrapper/caller ranges and import names reproduce. This
does not identify every allocator caller or close the CRT implementation.

**Unknown.** Failure callbacks, small-block internals, allocation success,
malformed block chains and node-pool aliasing.

### R2-ENGINE-219

TalkToD2.00137 extracts low16 NPC, bits16..27 topic and bits28..30 kind.
Kind3 calls AddMission(topic) atD2.00138 without a bank compare. That helper
admits matching type1/ID-topic catalog pointers to availability only when
each pointer is absent; missing catalog matches add nothing, and the
catalog itself is unchanged (R2-ENGINE-147).

NPC22/topic30 then passes two explicit identity compares and stores
bank533=1 atD2.00139 and bank553=2 atD2.00140. NPC2108/topic31 passes
neither special identity branch. NPC2110/topic39 kind0 matches neither
kind0 effect arm and reaches the return without a state store.
The offer predicates are R2-ENGINE-216. AddMission's topic is a mission
record ID; the NPC low16 is not the added mission ID.

**Confidence.** High for the complete selected TalkTo arms and call
arguments in the identical DLL code. No rendering or allocation-success
witness is claimed.

### R2-ENGINE-220

Builder ENL2.00552/RUL2.00553 copies the returned packed words, in order,
to its vector+0xfc. It extracts kind and low16; kinds1/2 resolve actors
into vector+0xc0, while other kinds resolve actors into vector+0xe8.
The measured vector count/index/append helpers establish count at+8,
backing pointer at+4 and four-byte indexing/appending.

Action ENL2.00554/RUL2.00555 selects a+0xe8 actor after subtracting the
first vector's count. Helper ENL2.00556/RUL2.00557 finds the first packed
word whose low16 matches actor u16+0x1dc. The action formats npc%dtalk%d
using NPC/topic, calls the text dispatcher, then calls TalkTo with the
same full word. Topic30/NPC22 therefore selects npc22talk30; topic31/NPC2108
selects npc2108talk31. R2-ENGINE-051 and R2-ENGINE-069 supply the published
dialogue/actor boundaries.

**Confidence.** High for native word copying, vector routing, first-match
selection, format arguments and call order. Medium for the composed
player-visible route: complete selected ranges agree across independent
linear/recursive decoding, but virtual rendering calls and resizing
callees are not closed and no original run occurred.

**Unknown.** Rendered rows, actor lookup/allocation success, dialogue
alternatives, multiple topics for one actor key and input scheduling.

### R2-ENGINE-221

The loopD2.00141..D2.00098 visits i=0..19. It compares bank552+i with
currentID and accepts bank532+i only when1 or 2. The emitted low ID is i+1,
kind is the retained value, and bit31 is set iff signed bank512+i>0.
These stores are not among R2-ENGINE-161's fifteen constant entries.

R2-SESSION-047/048 give immediate Leave20 values512=0,532=1,552=2.
An admitted currentID2 call in that state emits kind1/ID1 with bit31 clear.
R2-ENGINE-219's subsequent topic30 TALK adds 533=1,553=2, enabling kind1/ID2
under the same currentID. Intervening writers can change either inference.

**Confidence.** High for the complete native loop and conditional named
state. Static controls separately exercise the high bit. This is not a
claim of a complete runtime roster or an empty stage30 continuation.

## Remaining inn stage bodies

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-223 | Stage40 emits kind3 topics40/41/42/43 and kind0 topic48; only 41/42/43 require zero slots937/938/939, with no current-record gate in that body. | High | ✔ promoted | [EXP-2028](../experiments/EXP-2028-rom2-inn-stages/) |
| R2-ENGINE-224 | Stage50 emits kind0 topic49, kind3 topic51 only at 533!=0 and 947=0, and kind3 topic53 only at 949=0; the body has no current-record gate. | High | ✔ promoted | [EXP-2028](../experiments/EXP-2028-rom2-inn-stages/) |
| R2-ENGINE-225 | The shared stage60/70/80 body has nine kind3 stores, selected by currentID2/3, exact stage and bank gates; topic71 also requires966=0 because that compare skips both 70 and 71. | High | ✔ promoted | [EXP-2028](../experiments/EXP-2028-rom2-inn-stages/) |
| R2-ENGINE-226 | Stage90 emits kind3 topic90 for currentID2, or topics91/92 for currentID3 only when slots987/988 respectively equal zero. | High | ✔ promoted | [EXP-2028](../experiments/EXP-2028-rom2-inn-stages/) |
| R2-ENGINE-227 | Stage100 emits kind3 topics100/102/103 for currentID2, with 102 gated by 998=0 and 103 by 987!=0 and 999=0; currentID3 gets101 only at 997=0. | High | ✔ promoted | [EXP-2028](../experiments/EXP-2028-rom2-inn-stages/) |
| R2-ENGINE-228 | Stage110 emits its single kind3 topic110/NPC2006 store only for currentID2 before entering the shared continuation. | High | ✔ promoted | [EXP-2028](../experiments/EXP-2028-rom2-inn-stages/) |
| R2-ENGINE-229 | All 23 remaining kind3 options pass topic to TalkTo's AddMission call as type1/ID-topic; with initial recordID10 they add no special bank store, and TalkTo has no direct stage store. | High | ✔ promoted | [EXP-2028](../experiments/EXP-2028-rom2-inn-stages/) |
| R2-ENGINE-230 | Remaining kind0 options NPC2108/topic48 and NPC23/topic49 have no catalog or bank effect in TalkTo; the published client route selects npc2108talk48 and npc23talk49 text keys. | High / Medium | ✔ promoted | [EXP-2028](../experiments/EXP-2028-rom2-inn-stages/) |

### R2-ENGINE-223

Stage40 bodyD2.00146 has five stores in this order: kind3/topic40/NPC22
at D2.00147; kind0/topic48/NPC2108 at D2.00148; kind3/topic41/NPC2015
at D2.00149; kind3/topic42/NPC2111 at D2.00150; kind3/topic43/NPC2004
at D2.00151. The first two have no additional gate. CMP slot937,0
at D2.00152 guards topic41, CMP938,0 at D2.00153 guards topic42 and
CMP939,0 at D2.00154 guards topic43. Each JNE skips that store;
equality falls through. No current-record condition precedes these stores.
R2-ENGINE-229/230 give their subsequent TALK effects. Stores fill the
caller buffer; mission admission occurs on TALK, not on EnterInn.

**Confidence.** High for this complete local body. Capstone 5.0.7 recursive
and linear boundaries agree for complete EnterInn, with its switch tables
resolved and no unresolved indirect edge. Three-value zero/positive/negative
bank controls and currentIDs0,1,2,3,4,-1 agree with independent predicates.
Identical EN/RU code is one dependency. This is not runtime admission.

**Unknown.** Actual bank values, rendering, actor resolution and call scheduling.

### R2-ENGINE-224

Stage50 bodyD2.00155 stores kind0/topic49/NPC23 at D2.00156,
kind3/topic51/NPC2 at D2.00157 and kind3/topic53/NPC2019 at D2.00158.
Topic49 has no additional gate. Topic51 requires CMP533,0 at D2.00159
and CMP947,0 at D2.00160: nonzero 533 bypasses JE, then zero947 bypasses
JNE. Topic53 requires CMP949,0 at D2.00161, with equality bypassing JNE.
All admitting outcomes fall through. No current-record condition is tested
in this body. Kind3 selects type1/ID-topic through R2-ENGINE-229;
kind0's consumer is R2-ENGINE-230.

**Confidence.** High for all three stores and complete predicates. The
complete EnterInn decode and static three-value controls exclude inverted
zero gates and omission of the outer533 predicate. These are local contracts,
not a complete live bank population.

### R2-ENGINE-225

Stages60,70,80 select bodyD2.00162. All nine stores have kind3 and
subsequently request type1/ID-topic. Their order and predicates are:

| Store | Topic / NPC | Required predicate |
|---|---|---|
| D2.00163 | 70 / 680 | currentID3;771!=0;966=0 |
| D2.00164 | 71 / 681 | currentID3;771!=0;966=0;967=0 |
| D2.00165 | 83 / 677 | currentID3;768=80;979=0;536!=0 |
| D2.00166 | 61 / 2006 | currentID2;768=60;957=0 |
| D2.00167 | 63 / 2109 | currentID2;768=60;959=0 |
| D2.00168 | 73 / 2004 | currentID2;768=70;969=0 |
| D2.00169 | 72 / 2108 | currentID2;768=70;968=0 |
| D2.00170 | 81 / 2010 | currentID2;768=80;776!=0;977=0 |
| D2.00171 | 82 / 2009 | currentID2;768=80;776=0;978=0 |

Every bank predicate compares against zero except the explicit stage
comparisons against 60/70/80. CurrentID compares against 2 or 3 after getter
D2.00011. All admitting compare branches fall through except776=0,
which takes JED2.00172. CMP966,0 at D2.00173 followed by JNED2.00174
skips both 70 and 71 on nonzero; CMP967,0 at D2.00175 guards71 separately.
The per-store compare addresses, mnemonics and targets are in the
experiment's inn-gates.tsv, with every selected row independently replayed.

**Confidence.** High for the complete shared body and local gate conjunctions.
The full dispatch maps all three labels to one target. Complete instruction
boundaries, native static walks and independent predicates exclude treating
these as three unrestricted mission lists or ignoring966 for topic71.
No Ghidra function census or runtime visit is used.

**Unknown.** Live bank771/536/776 values, current-record scheduling and
presentation. R2-SESSION-113 separates the named pre-50 town paths.

### R2-ENGINE-226

Stage90 bodyD2.00176 emits kind3/topic90/NPC2004 at D2.00177 only for
currentID2. CurrentID3 instead admits kind3/topic91/NPC681 at D2.00178
only at 987=0 and kind3/topic92/NPC2003 at D2.00179 only at 988=0.
Current-ID comparisons are D2.00180 and D2.00181 against 2 and 3.
The bank compares are D2.00182/D2.00183 against 0. Each admitting
outcome falls through a JNE. All other current IDs bypass these stores.

**Confidence.** High for this complete body and the local current-ID split.
Resolved dispatch, complete boundaries and positive/negative bank controls
agree with the predicates. Each kind3 store has R2-ENGINE-229's consumer.

### R2-ENGINE-227

Stage100 bodyD2.00184 emits four kind3 options. CurrentID2 admits
100/NPC2006 at D2.00185 unconditionally within that arm, 102/NPC2109
at D2.00186 only at 998=0, and 103/NPC2005 at D2.00187 only at 987!=0
and 999=0. CurrentID3 admits101/NPC681 at D2.00188 only at 997=0.
Current-ID compares are D2.00189 and D2.00190. Bank compares against 0
are D2.00191,D2.00192,D2.00193 and D2.00194. Their admitting outcomes
fall through;987's JE skips103 on equality. There is no bank998
gate on 103 and no bank987 gate on 101.

**Confidence.** High for the complete four-store population and exact
predicate separation. Native boundaries and independently specified
three-value controls exclude transferring one option's gate to another.
The subsequent mission contract is R2-ENGINE-229.

### R2-ENGINE-228

Stage110 bodyD2.00195 compares currentID with 2 at D2.00196. JNE skips
kind3/topic110/NPC2006 at D2.00197; equality falls through. No additional
bank compare guards this store. Both outcomes then enter continuation
D2.00019. Its fixed and dynamic outputs remain R2-ENGINE-217/221.

**Confidence.** High for the single local store, equality gate and join.
Complete dispatch/boundary coverage and current-ID controls reproduce.
This does not assert that the entire EnterInn output is a single option.

### R2-ENGINE-229

The 25 remaining stage stores contain23 kind3 words. TalkToD2.00137
extracts kind from bits28..30 and topic from bits16..27, then calls
AddMissionD2.00014 at D2.00138 with topic. The requested catalog type
is 1 and ID is that topic; low16 NPC is not the mission ID.
R2-ENGINE-147 supplies absent-pointer admission and missing-key behaviour.
The caller-buffer stores themselves add no catalog record.

A generic store bank769=1 at D2.00121 requires topic equal to the ID
of objectD2.00143. R2-SESSION-110 identifies its initial ID10; none
of the 23 remaining topics is 10. None is NPC22/topic30, the other
special kind3 identity branch. With that initial identity the measured
remaining paths write no bank slot. The full 61-instruction TalkTo body
has no direct stage store; its only direct callees are AddMission and
ID getterD2.00011. The condition on object identity is preserved: a
synthetic matching ID enables bank769=1.

The separate kind1 TalkTo arm at D2.00198..D2.00199 toggles
bank[511+low16]. R2-ENGINE-221 bounds inn-emitted kind1 low IDs to 1..20,
so their indexed stores address slots512..531, not stage slot768.

**Confidence.** High for the complete local kind3 contract, conditional
bank result and no-direct-stage-store clause. PE translation, complete
recursive/linear agreement and 124 word/initial-ID static controls reproduce.
The negative covers these selected arms with initial ID10, not arbitrary
catalog mutation, allocator callbacks or every possible bank writer.

**Unknown.** Catalog matching/allocation success and mutations outside
that initial identity; no native click or live availability was observed.

### R2-ENGINE-230

Kind0/topic48/NPC2108 at D2.00148 and kind0/topic49/NPC23 at D2.00156
match neither state-changing kind0 identity arm in complete TalkTo.
They return without a catalog addition or bank store. Stage is unchanged
by those selected paths. R2-ENGINE-220's published client consumer routes
kind0 to talk actors, finds the first matching low16 key, formats
npc%dtalk%d and dispatches text before TalkTo. These pairs select keys
npc2108talk48 and npc23talk49; kind3 keys follow their own NPC/topic pair.

**Confidence.** High for the complete selected native no-state-effect
paths and key arithmetic; Medium for composed player presentation.
Native controls execute both kind0 words under two initial-record-ID
states. The reused client route leaves virtual rendering, actor lookup,
dialogue alternatives and input scheduling open.

**Unknown.** Displayed text, topic availability and successful interaction;
an authorized original visit and click would settle that composed route.

## Selected town presentation

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-231 | Selected ROM2 town IDs 1/2/3 bind generic/Kaarg/druid views and music b14/b16/b15; their measured loaders select different square art. | High | ● active (branch candidate) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |
| R2-ENGINE-232 | The active Kaarg square paints background, shop/inn highlight, girl1, girl2, guard, dervish and gates in that order, then delegates overlays and children. | High | ● active (branch candidate) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |
| R2-ENGINE-233 | The selected town sampler maps nine mask bytes to selectors; eight occur in the Kaarg mask, and hit testing uses the view-relative pixel rather than figure bounds. | High | ● active (branch candidate, partially retracted) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |
| R2-ENGINE-234 | Kaarg mask clicks open shop, inn, mission navigation or the main menu; its four person regions request stage-keyed dialogue, with no school click arm. | High | ● active (branch candidate) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |
| R2-ENGINE-235 | Kaarg person animation starts use separate elapsed-time intervals; frame advancement admits one step only after more than 100 ms, while scheduling runs on every active paint. | High | ● active (amended, branch candidate) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |
| R2-ENGINE-236 | Kaarg gates advance toward frame 10 while the pointer samples selector 8 and toward frame 0 otherwise; they use their frame series rather than a separate highlight bitmap. | High | ● active (branch candidate) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |
| R2-ENGINE-237 | Campaign town ID 2 selects Kaarg inn and shop pages with shared native bases; the inn mixes generic overlays with Kaarg art, while the shop selects Kaarg frame, fire and keeper art. | High / Medium | ● active (amended, branch candidate) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |

### R2-ENGINE-231

EN dispatcher L2.00266 and RU L2.00267 read the current record ID
through getters L2.00271 and L2.00275. The getters return record+4.
EN arms L2.00279, L2.00264 and L2.00280 select controller members
+0x110, +0x118 and +0x114, then call the chosen view's virtual +0x80.
They push music\b14.wav, music\b16.wav and music\b15.wav.
RU arms L2.00281, L2.00265 and L2.00282 have the same choices.
Playback is conditional on the dispatcher music-enable word; key
selection alone is not evidence of audible output.

EN construction L2.00558 stores generic constructor L2.00559 at
+0x110, Kaarg L2.00560 at +0x118 and druid L2.00561 at +0x114.
Their installed tables are L2.00562, L2.00563 and L2.00564.
The selected virtual +0xa0 loaders are EN L2.00565 / RU L2.00566
(generic), EN L2.00567 / RU L2.00568 (Kaarg), and EN L2.00569 /
RU L2.00570 (druid).

Generic loads interface/town/{townmask,townmain,Tavern_l,Trener_l,
Shop_l,Town_add}.bmp; sign/V%.2d.bmp, door/T%.2d.bmp, stars/S%.2d.bmp,
fluger/F%.2d.bmp; and interface/townbirds/{tavern,fighter,mage,shopie,
Guards}/sprites.16a, Birds%d/sprites.16a, HORSE%d/A%d/sprites.16a,
BABA%d/A%d/sprites.16a and DERVISH%d/sprites.16a.
Druid loads interface/town_druid/{townmask,townmain,hili_tavern,
hili_shop}.bmp, woman/a%d%04d.bmp, man/a%d%04d.bmp,
bug/sprites.16a and Lizard/sprites.16a.
Kaarg loads interface/town_kaarg/{townmask,townmain,hili_tavern,
hili_shop}.bmp, dervish/d%04d.bmp, guard/g%04d.bmp,
girl1/g%d%03d.bmp, girl2/g%d%03d.bmp and maingates/m%04d.bmp.
Each native graphics key has prefix graphics\. Literal casing and
individual push/call addresses are retained in the measured key tables.

Evidence is measured/selection/{instructions,resource-pushes,class-slots,
music}.tsv and measured/square/{anchors,resource-pushes}.tsv.
The three music nodes exist in both hashed music archives and agree.
The Kaarg loader and painter contain no TownBirds key or sprite draw;
that is a bounded local result, not executable-wide absence.

**Confidence.** High for the positive local branch, constructor, table
and pushed-key contracts. The selection decoder resolves each selected
body's direct CFG; square table and instruction controls retain the
Kaarg dispatch targets. This closes R2-ENGINE-160's unread arm strings
and identifies its +0x118 view without claiming a global consumer census.

**Unknown.** Other ID consumers, loose-resource overrides, runtime
admission of each view and audible music delivery.

### R2-ENGINE-232

EN painter L2.00571 / RU L2.00572 returns when view+0x204 is zero.
For an active view it paints the following destinations relative to
view left/top: townmain (0,0); shop highlight (328,256) for hover 1
or inn highlight (480,196) for hover 2; girl1 (216,284); girl2
(260,284); guard (140,152) for active cursor 13..37, otherwise
(184,156); dervish (416,328); and maingates (152,256).
Inactive person flags select the first subset's first frame. Active
flags select the retained subset/frame. Gate frame is always selected
from its current cursor.

Every square raster call uses virtual +0x18. EN BMP constructor
L2.00573 installs table L2.00574, slot +0x18=L2.00575, which forwards
x,y to L2.00576. That bounded receiver copies opaque 16-bit rows,
decrementing the source row from the BMP bottom-up buffer. There is no
color-key test in that receiver. BGR24 conversion L2.00577 uses device
format globals; a fixed RGB565 pixel format is not established.
Generic overlay L2.00578 calls the view's +0x30 and child paint
L2.00579. The Kaarg +0x30 target L2.00580 is empty.

Evidence is measured/square/{painter-en,painter-ru,bitmap-load-en,
bitmap-paint-en,bitmap-blit-en,bgr24-convert-en,overlay-dispatch-en,
child-paint-en,own-overlay-en}.txt and geometry.tsv.

**Confidence.** High for the selected painter order, operands and
identified EN bitmap receiver. All selected painter branch targets
are decoded boundaries. No claim enumerates other child constructors
or substitutes a source-RGB PNG for observed original pixels.

**Unknown.** The current device format, the wider shell's overlays,
capture/cursor effects and original runtime presentation.

### R2-ENGINE-233

EN sampler L2.00581 / RU L2.00582 rejects an absent mask or an
out-of-view point with -1. It subtracts the view origin, samples the
indexed buffer at x+640*y and uses a byte-index/dword jump table.
Bytes 32,64,80,96,128,144,160,176,192 map to selectors
0x400,0x800,0x1000,0x200,2,1,8,16,4 respectively; the other
247 byte values map to -1. R2-ASSET-050 bounds the actual population:
byte 80 is absent in both Kaarg masks.

EN mask loader L2.00583 and row transform L2.00584 reverse stored
BMP rows before sampling. Screen pixel (x-left,y-top) therefore uses
the top-origin picture index with no sprite extent test.
Hover L2.00585 / L2.00586 stores a changed selector at +0xb4.
Selectors 1 and 2 take latched sound arms: shop starts Kenter2 at
+0x22c under latch +0xac and stops +0x230; inn starts Kenter1 at
+0x230 under latch +0x250 and stops +0x22c. Selector 8 stops both
sounds under latch +0x254; -1 calls reset L2.00587 and clears
the latches. EN sound loader L2.00588 binds the two Kenter keys.
These arms do not write +0x208. Selectors 4,0x200,0x400,0x800
take a switch no-op arm. Among the named selectors, only
0,3,16,0x1000 reach the OR into +0x208; only 16 and 0x1000
from that group are sampler outputs. The default OR arm also admits
other values outside the selected sampler's output domain.
Person hovering does not start their timer-driven episodes.
The inherited EN slot +0x4c at L2.00589 calls hover +0x98.
Hover is therefore also reachable outside the painter's 100 ms arm.

Evidence is measured/square/{mask-map,mask-index-en,mask-dword-en,
hover-dword-en,cfg-controls,branches}.tsv and mask/hover loader
instruction excerpts. The selected Kaarg click body calls this sampler.
Supplemental sound-loader and inherited-hover excerpts and slot
bindings are in measured/correction.

**Confidence.** High for the complete 256-byte selector domain and
the selected sampler, hover and click paths. The decoder preserves
every finite table target on an instruction boundary. Archive pixel
counts are measured independently of function recovery.

**Unknown.** Physical pointer delivery, mouse capture and added child
widgets outside the selected construction/entry paths.

**Amended.** The remaining-selectors flag-write clause is partially
retracted in claims/retracted.md. The decoded sound, reset and OR arms
replace it; the sampler and person-hover clauses stand.

### R2-ENGINE-234

EN click L2.00590 / RU L2.00591 dispatches decoded selector 1 to
message 0x42a, 2 to 0x42b, 8 to 0x442(1,0) followed by 0x42d,
and 16 to 0x41f. Shop, inn and gates call leave preparation first;
the main-menu arm does not. Person selectors 0x800,0x200,4,0x400
request kaargguard%d, kaargwoman%d, kaargwoman%d and kaargman%d,
using Scenario variable 0x300 as the suffix. They call dialogue
entry R2.0038 rather than open a separate room.

The window dispatcher EN L2.00262 / RU R2.0036 maps 0x42a and
0x42b to the room selectors R2-ENGINE-237 identifies. Message 0x41f
creates popup L2.00592 / L2.00593, mode 1, at (100,100), size
440x340. Its dialogs.txt binding labels save, load, sound options,
quest end, game return and exit actions: it is the main menu.
Gate posting is measured locally; the full mission-navigation
acceptance and movie/route sequence are not established here.

Tip getter L2.00594 / L2.00595 returns main-table indices
233,236,237,235 for shop, inn, gates and menu; 359,360,361,362
for girl1, girl2, guard and dervish. Selected EN text identifies
the last as sleeping man, while selected RU text calls him beggar.
These are labels for one native region, not different sampled masks.

Evidence is measured/square/{click,tip-getter}-*.txt,
measured/rooms/{room-dispatch,menu-text}.tsv and measured/text/.
The menu identity is a native constructor/table/text join.
No school region is inferred from ROM1's mask16 semantics.

**Confidence.** High for the selected click arms, posted messages,
room/menu receivers and text-index bindings. The selected byte-domain
and finite target tables are complete. This is not an enumeration of
all physical input handlers or popup controls.

**Unknown.** Full gate receiver flow, person dialogue contents and
runtime input/admission. Following those native receiver paths or an
authorized original observation would settle them.

### R2-ENGINE-235

EN active painter L2.00571 imports timeGetTime. When elapsed time
from its process baseline exceeds 100 ms it polls the pointer, calls
advance L2.00596 and resets that baseline to current time. It never
catches up several frames. Scheduler L2.00597 then runs on every
active paint, including paints that did not admit a frame step.
RU counterparts are L2.00572, L2.00598 and L2.00599.

Enter initializes girl1/girl2/guard waits to 2000+R(2000) ms and
dervish to 4000+R(500). At a later strictly elapsed wait, scheduler
sets the family cursor to 0, saves the current family clock and
replaces its interval: girls 3500+R(5000), guard 7500+R(5000),
dervish 3700+R(500). Girls choose subset 0 or 1 with R(2).
The elapsed-wait arms do not test the family flag. A slow redraw
cadence can therefore restart an episode that is still running.
R(n)=floor(rand()*n/32767)%n on the read integer path;
the native source masks its output to 0x7fff. This establishes
ranges, not a uniform or independent distribution.

Thus initial waits are 2000..3999 ms for girls/guard and
4000..4499 for dervish; subsequent waits are 3500..8499,
7500..12499 and 3700..4199 respectively. The start flags are
girl1 0x200, girl2 4, guard 0x800 and dervish 0x400 in +0x208.
Each admitted family advance increments once; equality with the
selected vector count clears the flag and resets its cursor.
Girl subset cursors correspond to the filename ranges in R2-ASSET-051.
Idle drawing selects the first frame separately from the episode.

Evidence is measured/square/{clocks,bounds,timer-import,
random-source,cfg-controls,branches}.tsv and painter, scheduler,
advance, person-advance and random helper excerpts.

**Confidence.** High for this native predicate, update order and
bounded interval arithmetic. Timer imports, random source mask,
17 selected method decodes, 193 direct branch targets and 34 finite
table targets per locale are explicit controls. The methods are
examined without global writer or PRNG independence assumptions.

**Unknown.** Delivered redraw cadence, the process baseline before
re-entry and actual random sequence. An authorized timed original
session would settle appearance and elapsed-time delivery.
Audible background ambience remains Unknown. The selected scheduler
also requests sound slots +0x20c/+0x210/+0x214 after 2000+R(2000)
ms with R(3), +0x218..+0x224 after a separate 2000+R(2000) ms
with R(4), +0x234 after more than 45000 ms, and +0x238..+0x24c
at guard cursor states. These are bounded request clocks, not
observed sound playback or a complete ambience contract.

**Amended.** R2-ENGINE-280 names the keys behind the scheduler's sound slots, their waits and the reachable guard cursors; audible output stays Unknown.

### R2-ENGINE-236

EN gate advance L2.00600 / RU L2.00601 samples the current
pointer through the native mask helper on each admitted advance.
Selector 8 increments cursor +0x1a0 and clamps it to 10; another
selector decrements and clamps it to 0. It records direction state
at +0x1a4 and clears flag 8 at the endpoint. Direction changes
request SFX\Town_kaarg\Kdoor1.wav or Kdoor2.wav once in these
local arms. The painter always draws the retained frame at (152,256).

Evidence is measured/square/{gate-en,gate-ru,advance-en,
advance-ru,painter-en,painter-ru}.txt. R2-ASSET-051 provides
the complete eleven-frame archive range.

**Confidence.** High for the bounded pointer predicate, direction
state, integer clamps and draw choice. The whole selected method
and its direct branch targets are retained. Gate completion is
independent of whether the +0x208 gate flag is currently set.

**Unknown.** Physical hover delivery, audible sound and the complete
mission-start receiver. The click's posted messages are R2-ENGINE-234.

### R2-ENGINE-237

In campaign mode EN app+0x5d8==2 / RU +0x63c==2, shop
selectors L2.00602 / L2.00603 and inn selectors L2.00425 /
L2.00427 read the Scenario current ID. ID2 chooses app+0x100
(shop) and +0x10c (inn). EN construction L2.00558 binds
those members to page constructors L2.00604 and L2.00605.

Kaarg inn page L2.00605 / L2.00606 calls common inn base
L2.00607 / L2.00608. Its center L2.00609 calls common inn
center L2.00610, then installs a derived table. Art loader
L2.00611 / L2.00612 requests generic interface/inn/manback.bmp,
ManBackTalk.bmp, LUOver.bmp, LDOver.bmp and RUOver.bmp,
plus interface/inn_kaarg/TavernMain.bmp. Campaign entry
L2.00613 / L2.00614 copies taverner a1..a5 vectors of
15,25,3,7,29 frames into center fields +0x3cc,+0x3fc,+0x42c,
+0x45c,+0x48c. R2-ASSET-052 records their filenames and geometry.

Kaarg shop page L2.00604 constructs child L2.00615 /
L2.00616, which calls common shop base L2.00617 / L2.00618
then installs its derived table. Loader L2.00619 / L2.00620
requests interface/shop_kaarg/ShopFrame.256, ShopMain.bmp,
hili_armor.bmp, hili_magic.bmp, hili_potion.bmp and hili_weapon.bmp;
six fire families dark/select crossed with lite/burn/cicle;
and movies\shop_kaarg\a10000.bmp. Updater L2.00621 /
L2.00622 increments its counter modulo 20. Flag 0x10 chooses
movies\shop_kaarg\a1%04d.bmp and 0x20 chooses a2%04d.bmp,
with counter+1; otherwise it uses a10000.bmp. Counter zero
clears the episode flags outside +0x240's low nibble.

Evidence is measured/rooms/{native-resource-loads,inn-frame-loads,
room-dispatch,ranges}.tsv, selected constructor/art/update excerpts
and vtable cells, plus measured/selection/instructions.tsv.

**Confidence.** High for the positive native selection, keys, vector
ranges and common-base calls. Medium for reuse of room layout:
shared base widgets establish shared structure, but derived
constructors and paint methods can use different rectangles.
Only-art differences and complete layout equality are not proven.

**Unknown.** Full room layout, animation destinations, caller cadence
and trigger conditions. These need the later derived room painter
and scheduling probe; native play would settle live presentation.

**Amended.** R2-ENGINE-277 answers the inn taverner destinations and triggers; R2-ENGINE-278 and R2-ENGINE-279 answer the shop fire and keeper destinations, triggers and the 100 ms tick rule; delivered paint and tick cadence stays Unknown.

## Kaarg view rectangle and shell boundary

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-238 | The selected EN initializer centers the 640x480 town view in three screen sizes; Kaarg entry conditionally adds a child, while final shell text and button drawing remain unidentified. | High / Unknown | ● active (branch candidate, partially retracted) | [EXP-2029](../experiments/EXP-2029-rom2-town-screens/) |

### R2-ENGINE-238

EN initializer L2.00623 first checks the word at L2.00624;
a nonzero value forces 640x480. Otherwise it tests the command
line at app+0x70 for -800, -1024 and -640 in that order.
Without a matching switch it tests buffer L2.00625 for -800
and -1024, then falls back to 640x480. Registry initialization
queries RESOLUTION into that buffer at L2.00626 and calls the
fallback copy of -640 at L2.00627 if the query fails. Thus 640x480
is the default only without a selecting command-line or registry value, or
when the force word is nonzero. The other sizes are 800x600
and 1024x768. It stores width/height at
L2.00628/L2.00629 and computes left=(width-640)/2 and
top=(height-480)/2, with signed division truncated toward zero.
Right=width-left and bottom=height-top. The measured origins
are therefore (0,0), (80,60) and (192,144).
Construction at L2.00630 passes this rectangle to Kaarg
constructor L2.00560 with ID 0x3fc.

Kaarg entry EN L2.00631 / RU L2.00632 checks a global word.
The admitted EN arm constructs L2.00633, stores it at view+0x200,
and adds child ID 0x467 with left=328, top=0, right=640,
bottom=200 and resource argument 12. The rectangle is 312x200.
The EN chain L2.00633 -> L2.00634 -> L2.00635 -> L2.00636 ->
L2.00637 forwards the bounds to SetRect at import cell L2.00638.
Whether these bounds are view-relative or screen-relative is Unknown.
This is a positive child path, not
a census of application overlays. R2-ENGINE-232 identifies the
empty own-overlay slot and following child paint.

Evidence is measured/selection/instructions.tsv, selected bodies
campaign-screen-initializer and campaign-child-construction,
and measured/square/{enter-en,enter-ru,ctor-en,ctor-ru}.txt.
Supplemental registry and rectangle receiver excerpts, decode
controls and the SetRect import binding are in measured/correction.
R2-ENGINE-234 identifies the pointer text keys. The selected
square painter, entry and tip bodies do not identify the final
application text receiver or all shell widgets.

**Confidence.** High for the EN screen-selection arithmetic,
constructor arguments and positive EN/RU child path. Unknown
for the child coordinate reference, purpose, live visibility, final
tip/status/button destinations and their draw order. No absence of status lines,
buttons or text is claimed. This confidence does not infer an
unmeasured RU screen initializer from EN agreement elsewhere.

**Unknown.** Wider shell composition and device presentation.
Trace the active Kaarg virtual +0x14 consumer into its final
text renderer and child paint receivers to settle destinations.
An authorized original capture would settle live appearance.

**Amended.** The child width/height interpretation and unconditional
default-resolution clause are partially retracted in claims/retracted.md.
Rectangle bounds and command-line/registry precedence replace them.

## First-town square

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-239 | ROM2 town ID 1 binds the generic native square and a centered 640x480 rectangle; the selected shared receiver draws delayed pointer tips, while other shell visibility remains unobserved. | High | ● active (amended, branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ENGINE-240 | The selected ROM2 town-1 loader has 19 graphics-interface key literals; its painter orders bitmap, sprite and child layers, with three loaded graphics absent from the complete painter and advance bodies. | High | ● active (branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ENGINE-241 | ROM2 town 1 has nine positive mask selectors across 150 raw pixel indices; finite native tables separate hover and click effects, and the school selector has a tip but opens no room. | High | ● active (amended, branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ENGINE-242 | The selected ROM2 town-1 painter admits one animation step after elapsed time exceeds 67 ms; its tavern, sign, stars, shopie and weather vane use distinct finite episode rules. | High | ● active (branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ENGINE-243 | ROM2 town 1 draws nine gate bitmaps and eight guard sprite frames; native flag 0x301 and pointer selection control direction, endpoint clamps and sound requests. | High | ● active (amended, branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ENGINE-244 | ROM2 town 1 schedules 57-frame bird groups and BABA/HORSE episodes, chooses actor art and positions at load, and advances a 30-frame DERVISH continuously under its entry flag. | High | ● active (branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |

### R2-ENGINE-239

Town ID 1 selects the generic view through controller member +0x110 (`R2-ENGINE-231`). EN construction at `L2.00706` calls `L2.00559` with window ID `0x3fc` and rectangle globals `L2.00707..L2.00708`. The constructor installs vtable `L2.00562` and calls initializer `L2.00709`. RU uses vtable `L2.00700` and corresponding generic methods at EN+`0x8170`. The bindings are in [anchors](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/anchors.tsv).

The generic enter, painter and loader are EN `L2.00710`, `L2.00711`, `L2.00565`. The view is 640x480. Shared screen initializer EN `L2.00623..L2.00712` centers it in 640x480, 800x600 or 1024x768, giving origins (0,0), (80,60), (192,144). A nonzero force word selects 640x480. Otherwise command-line `-800`, `-1024`, `-640` are tested in order, then the retained resolution buffer for `-800` and `-1024`, then the 640x480 fallback (`R2-ENGINE-238`). These are the same native rectangle globals used for the Kaarg construction.

The generic initializer zeros resource and state pointers and attaches child ID `0x445`. Enter conditionally creates child `0x467` when global `L2.00713` is nonzero, with native rectangle bounds (left=328, top=0, right=640, bottom=200), extent 312x200. The same receiver and bounds occur in the selected Kaarg entry (`R2-ENGINE-238`). The generic own-overlay method `L2.00580` is empty. The painter delegates children through `L2.00578` after its square layers.

The EN shared pointer receiver `L2.00714..L2.00715` finds the child at the pointer through `L2.00716` and calls its tip slot +0x14. A new tip requires elapsed accumulator +0x58 to cross from a previous unsigned value below 500 to a signed value at least 500 and below 25500, plus the receiver/controller admission fields. Nonempty text passes to `L2.00717`, `L2.00718`, `L2.00719` and font call `L2.00720`. For pointer (x,y), maximum measured line width w and line count n, the initial box is (x,y-5-14n,x+w+11,y). The receiver shifts a right overflow to screen right and a top overflow to the maintained screen top. Text begins at (left+5,top+4+14i) for line i. Exact bodies are in [shell controls](../experiments/EXP-2030-rom2-first-town/evidence/measured/shell/controls.tsv).

**Confidence.** High for the selected generic bindings, rectangle arithmetic, conditional child operands and EN tip receiver. The PE/Capstone instrument records input hashes and versions, verifies complete selected instruction ranges and direct branch boundaries, and checks vtable cells. This excludes an art-only view substitution, a full-screen-scaled town rectangle and an untraced final tip destination within the selected receiver.

**Unknown.** Live visibility and purpose of child `0x467`, its view/screen coordinate reference, and other status text or buttons remain unobserved. The seven EN shell bodies are not a complete shell census. Native execution with a recorded screen state would settle these clauses.

**Amended.** R2-ENGINE-274 identifies child `0x467` as the TipsMode tips panel with parent-relative bounds, and R2-ENGINE-273 names the writers of `L2.00713`. Live visibility remains unobserved.

### R2-ENGINE-240

The generic square loader has 19 graphics-interface literal keys per locale. [Resource pushes](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/resource-pushes.tsv) identify each literal and instruction. Paths below are relative to `graphics.res`; native literals prefix `graphics\`.

| Resource | Native loaded population |
|---|---|
| `interface/town/{townmain,townmask,Tavern_l,Trener_l,Shop_l,Town_add}.bmp` | six fixed files, including the mask |
| `interface/town/sign/V%.2d.bmp` | V00..V09, ten bitmaps |
| `interface/town/door/T%.2d.bmp` | T00..T08, nine bitmaps |
| `interface/town/stars/S%.2d.bmp` | S00..S08, nine bitmaps |
| `interface/town/fluger/F%.2d.bmp` | F00..F07, eight bitmaps |
| `interface/townbirds/{tavern,fighter,mage,shopie,Guards}/sprites.16a` | five fixed sprite files |
| `interface/townbirds/Birds%d/sprites.16a` | Birds1..Birds9 |
| `interface/townbirds/HORSE%d/A%d/sprites.16a` | one random variant 1..5, its A1..A3 |
| `interface/townbirds/BABA%d/A%d/sprites.16a` | one random variant 1..4, its A1..A2 |
| `interface/townbirds/DERVISH%d/sprites.16a` | one variant 1..4 unequal to the selected BABA variant |

Town-1 music is `music\b14.wav`, the `b14.wav` key in `music.res`, selected outside the graphics loader under the music-enable word (`R2-ENGINE-231`).

The active generic painter uses this order. Coordinates add to view origin. [Geometry](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/geometry.tsv) and the complete [EN painter](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/painter-en.txt) preserve load-bearing immediates and conditions.

| Order | Layer | Destination | Condition |
|---|---|---|---|
| 1 | townmain | (0,0) | loaded pointer |
| 2 | selected birds | (0,0), in selection order | active bit 0x80 and frame below count |
| 3 | Town_add | (0,0) | bird-paint branch; follows the bird sprites |
| 4 | Shop_l | (264,264) | selector 1 |
| 5 | Tavern_l | (144,332) | selector 2 |
| 6 | tavern sprite | (124,312) | loaded pointer, retained frame |
| 7 | sign/V | (360,232) | retained bitmap |
| 8 | gate/T | (180,148) | retained bitmap |
| 9 | stars/S | (340,288) | retained bitmap, may be null |
| 10 | shopie sprite | (276,296) | loaded pointer, retained frame |
| 11 | weather vane/F | (308,64) | retained bitmap |
| 12 | guards sprite | (184,158) | loaded pointer, retained frame |
| 13 | HORSE sprite | selected position | active frame; otherwise frame 0 |
| 14 | BABA sprite | selected position | active frame; otherwise frame 0 |
| 15 | DERVISH sprite | selected position | retained frame |
| 16 | shared child dispatch | shared receiver | after square surface end |

The resource destination fields are background+0x68, mask+0x6c, addition+0x70, inn highlight+0x164, shop highlight+0x1dc and school highlight+0x1c4. The complete generic painter `L2.00711..L2.00565` and advance `L2.00721..L2.00722` contain no use of loaded school highlight+0x1c4, fighter+0x1c8 or mage+0x1d0: exactly three graphics in this bounded negative. The complete loader does not attach these graphics as children. The Kaarg painter instead uses its own numbered BMP figures and positions (`R2-ENGINE-232`); these generic sprites are not substitutions for those Kaarg figures.

Town_add uses bitmap slot+0x38, whose selected native receiver skips a zero converted 16-bit pixel. Other selected BMP layers use slot+0x18, the ordinary copy receiver. The [native art receivers](../experiments/EXP-2030-rom2-first-town/evidence/measured/art/native-bodies.tsv) preserve their distinct targets. Neither path establishes a fixed RGB565 display.

**Confidence.** High for the complete selected loader key population, painter order and destinations, and the three-resource negative bounded to the named painter/advance bodies. Complete instruction ranges, local branch targets and resource-result stores rule out archive-order composition and preceding-store misassociation. The instrument does not use function-name enumeration for absence.

**Unknown.** Runtime frame/hover state, other whole-image uses of the three loaded graphics, delivered pixels and device conversion remain unobserved. A complete pointer-use sweep would settle global use; an authorized runtime image would settle live composition.

### R2-ENGINE-241

Both preserved generic masks are 640x480, with 307200 cells and 150 distinct raw pixel indices. The selected sampler `L2.00581` rejects a missing mask or an out-of-view pointer, subtracts view origin, then reads byte[x+640*y]. It uses a 161-byte finite index table and ten DWORD arm pointers. The complete [256-byte map](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/mask-map-en.tsv), [tables](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/tables.tsv) and [mask measurements](../experiments/EXP-2030-rom2-first-town/evidence/measured/art/masks.tsv) establish every positive selector present.

| Pixel byte | Selector | Cells | Tip index | Native click |
|---|---|---|---|---|
| 32 | 0x400 | 17 | none | default, returns 1 |
| 64 | 0x800 | 18 | none | default, returns 1 |
| 80 | 0x1000 | 4 | none | default, returns 1 |
| 96 | 0x200 | 6 | none | default, returns 1 |
| 128 | 2 | 5764 | 236, inn | leave-and-remove request 0x445, then 0x42b |
| 144 | 1 | 5771 | 233, shop | leave-and-remove request 0x445, then 0x42a |
| 160 | 8 | 7911 | 237, gate | flag 0x301 gate; allowed leave-and-remove request 0x445, 0x442(1,0) then 0x42d, otherwise `plagatguard` text |
| 176 | 0x10 | 5592 | 235, menu | message 0x41f |
| 192 | 4 | 14733 | 234, school | default, returns 1; opens no room |
| all other indices | -1 | 267384 | none | default, returns 1 |

The all-other cell includes zero and the 140 remaining present indices, not only background zero. The nine mapped bytes total 39816 cells. Tip slot+0x14 returns null when the square active word+0x204 is zero. The finite getter maps only selectors 1,2,4,8,16 to the five text indices above. `R2-ENGINE-239` supplies the final delayed-tip receiver.

Hover `L2.00723` writes selector+0xb4 and guard direction+0xec=1 before dispatch. Selector 1 draws Shop_l through the painter; when shop latch+0xac is zero it stops School/Point.wav, starts Shop/enter.wav and sets the latch. If `rand()%100>95` and bit 1 is inactive, it sets bit 1 for shopie. Selector 4 has an explicit no-op arm after the common writes: no school highlight, sound start or animation OR. Selector 8 sets guard direction-1 when flag 0x301 is false and 1 when true. Selector-1 forwards the retained global descriptor through `L2.00587`. Gate, selector-1 and ordinary OR arms clear the shop latch and +0xb0.

Other selectors OR their value into animation flags+0x208: inn 2 starts tavern bit 2, menu 0x10 activates stars, small mask 0x200 can set BABA bit 0x200, and 0x400 sets DERVISH bit 0x400. Bits 0x800/0x1000 have no receiver in the complete selected advance body. The school pointer still takes the common guard-direction write; no absence of all hover state change is asserted.

The complete click `L2.00724..L2.00725` uses a finite 16-case byte table. Shop and inn send leave-and-remove request 0x445 before posting. Gate flag false invokes `plagatguard` through the controller text helper; flag true sends the same request then posts the two messages. R2-ENGINE-250 traces square release and removal. School and the four small figure selectors take the default arm. The menu opens the existing main-menu receiver (`R2-ENGINE-234`), not another town room.

The EN [input-slot census](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/input-slots-en.tsv) covers the remaining vtable `L2.00562` input slots. Slot +0x48 `L2.00726` handles messages without a direct selector+0xb4 read; +0x4c `L2.00589` calls the measured hover and returns 0; +0x6c `L2.00727` accepts Enter/Escape and forwards other keys to the base key handler. Slots +0x50, +0x58..+0x68 and +0x70 return 0. None of these named bodies directly reads selector+0xb4 to open a room. This census supports the bounded school and figure click result; it is not a whole-image indirect-input census.

**Confidence.** High for the complete mask population, finite mapping, positive hover/click branches and school no-op. The archive walker measured every mask pixel; the native generator covered all 256 possible bytes and every switch target in both selected locales. This excludes treating mask indices as direct selectors, inferring school behavior from its loaded art and equating hover flags with click destinations.

**Unknown.** Live flag 0x301 value, the retained descriptor's actual pointer graphic, accepted gate navigation and the effects of unobserved shell admission remain open. Selected posting is not proof that the destination accepts navigation. Recorded native execution would settle these live clauses; a bounded descriptor initialization trace would settle its graphic.

**Amended.** The former "departure prepare 0x445" label is corrected to a leave-and-remove request. The selector map and school no-op stand. See the correction entry in [retracted.md](retracted.md). R2-ENGINE-271 and R2-ENGINE-272 give the receivers of the posted gate messages 0x442 and 0x42d; R2-SESSION-128 gives the flag value.

### R2-ENGINE-242

Generic painter EN `L2.00711..L2.00565` reads imported `timeGetTime`. Unsigned elapsed time strictly greater than 67 ms admits hover, random triggers and one call to advance slot+0xa8, then replaces the global baseline with current time. There is no accumulated multi-frame catch-up loop in this body. The independent actor scheduler still runs on every active paint. [Timer identity](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/timer.tsv), [painter](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/painter-en.txt) and [advance](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/advance-en.txt) preserve this branch.

Define R(n) = floor(rand()*n/32767)%n. Helper `L2.00728` and source `L2.00729` return 0..n-1 from the measured 15-bit source. Uniformity and independence are unestablished. The generic step requests sign animation when R(100)>94 and weather-vane animation when a second R(100)>97. These OR bits 0x40/0x20 before the admitted advance.

| Family | Frames / destination | Admission | Advance and end |
|---|---|---|---|
| tavern | sprite 0..9 / (124,312) | bit 2 from inn hover | increment 1; equality with count 10 restores 0 and clears bit 2 |
| sign | V00..V09 / (360,232) | bit 0x40 from random request | increment 1; at 10 restore 0 and clear bit 0x40 |
| stars | S00..S08 / (340,288) | bit 0x10 from menu hover | increment 1; below 9 selectS; at/above 9 increment persistent blank counter, set current bitmap null and clear bit 0x10; on its tenth such call restore index 0/counter 0, still null that call |
| shopie | sprite 0..29 / (276,296) | bit 1 from special shop hover | increment 1; equality with count 30 restores 0 and clears bit 1 |
| weather vane | F00..F07 / (308,64) | bit 0x20 from random request | increment 1; at 8 restore 0 and clear bit 0x20 |

The tavern helper plays Town/Point.wav when its old frame is 0, after stopping Shop/enter.wav and School/Point.wav. Sign, stars and weather vane request Flag.wav, Stars.wav and Flugel.wav respectively when their old index is 0. The loop-start helper invoked at entry requests Crowd.wav. The one-shot helper checks whether its retained sound is already playing before starting it. These are native requests, not observed audible playback.

Loader initialization calls set tavern 0, V00, S00, shopie 0 and F00. The stars global blank counter and painter lazy clock variables are not reset by generic enter; exact re-entry history can change the retained phase. Gate and guard advancement are `R2-ENGINE-243`; the wildlife/actor schedules are `R2-ENGINE-244`. The loaded fighter/mage sprites are outside the selected advance population (`R2-ENGINE-240`).

**Confidence.** High for the bounded timing branch, helper ranges, request inequalities and episode rules. The complete finite advance bodies, direct boundary controls, timer import identity and inspected random source exclude a nominal fixed-frame-rate guarantee, an unbounded random range and a simple nine-frame stars loop.

**Unknown.** Redraw cadence, random outputs, retained global history, active sound handles and audible playback were not observed. A recorded native run with state/timer capture would settle them. No precise live duration or random probability is claimed.

### R2-ENGINE-243

Town-1 gates use `interface/town/door/T00..T08.bmp` at view-relative (180,148). Native gate helper EN `L2.00730..L2.00731` runs on each admitted generic step. The pointer selects gate 8 through the shared mask sampler. The helper queries Scenario flag 0x301. A false flag sets cursor 8 and selects T08 immediately. A true flag decrements the cursor toward 0 while the pointer is gate 8, and increments it toward 8 otherwise. Endpoints clamp to 0/8. A change to hover direction requests GateUp.wav; leaving that direction requests GateDn.wav. The retained direction field and current bitmap are updated in the selected arms.

This direction convention differs from the Kaarg gate cursor, which increases under hover (`R2-ENGINE-236`). File index alone therefore cannot be shared as an opening-direction rule.

The guard sprite is `interface/townbirds/Guards/sprites.16a`, eight frames 0..7 at (184,158). Entry sets frame 7, direction 0 and sound-direction latch 0. Every hover call first sets direction 1. Only gate 8 with a false flag changes it to-1. Guard helper `L2.00732..L2.00733` adds this direction once per admitted step and clamps to 0..7, setting direction 0 at either endpoint. Transition to direction 1 with latch 0 loads/requests Guard2.wav and sets the latch; transition to direction-1 with nonzero latch loads/requests Guard1.wav and clears it. Endpoint handling releases the retained guard sound handle.

Gate art, guard art and click acceptance are distinct paths: a false flag selects T08, reverses the guard and shows `plagatguard` on click; a true flag admits the gate click posting in `R2-ENGINE-241`. Naming T08 "closed" is unmeasured; neither file index nor the flag-false branch establishes that visual label.

Evidence is the complete [gate helper](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/gate-advance-en.txt), [guard helper](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/guard-advance-en.txt), [hover](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/hover-en.txt), [entry](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/enter-en.txt) and [art frames](../experiments/EXP-2030-rom2-first-town/evidence/measured/art/frames.tsv).

**Confidence.** High for these selected finite paths and coordinates. Complete instruction ranges and branch controls exclude an ungated hover-only mechanism, interchangeable gate/guard cursors and Kaarg's index direction. The grade does not include live flag admission or a rendered open state.

**Unknown.** The visual "closed" label for T08, actual flag 0x301 value, campaign reason for its value, sound playback and acceptance of posted navigation remain unobserved. A measured art-state join would settle the label. The campaign flag writer graph and an authorized runtime trace would settle the live clauses without changing the bounded cursor rules.

**Amended.** The unmeasured "closed" label for T08 is withdrawn from the observed rule. The flag-false T08 selection and cursor mechanics stand. See the narrowing entry in [retracted.md](retracted.md). R2-SESSION-128 settles the campaign value of flag 0x301: 0 at the first town-1 entry, 1 after the stage-10 topic-10 talk.

### R2-ENGINE-244

The selected loader chooses HORSE variant 1..5 with R(5), BABA variant 1..4 with R(4), then retries DERVISH R(4) until its variant differs from BABA. Native coordinates and table references are in [positions](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/positions.tsv).

| Variant | HORSE position | BABA position | DERVISH position |
|---|---|---|---|
| 1 | (104,404) | (216,364) | (224,364) |
| 2 | (104,404) | (308,424) | (324,424) |
| 3 | (256,344) | (384,424) | (392,420) |
| 4 | (448,400) | (580,384) | (592,388) |
| 5 | (140,400) | absent variant | absent variant |

Each HORSE variant loads A1..A3, all 15 frames. Each BABA variant loads A1=31 frames and A2=32 frames. Each DERVISH variant has 30 frames. Entry retains HORSE A1 and BABA A1 with frame -1, which the painter renders as frame 0. Entry sets DERVISH frame 0 and bit 0x400; admitted advancement uses (frame+1)%30.

Independent BABA/HORSE scheduler EN `L2.00734..L2.00723` runs on every paint. Both initial waits are 2000+R(2000), range 2000..3999 ms. Later waits are 2000+R(5000), range 2000..6999 ms. A start uses strict unsigned elapsed>wait, sets frame 0, randomly chooses one loaded action and sets BABA bit 0x200 or HORSE bit 0x100. The elapsed arms do not test the active flag. A slow redraw can therefore restart an episode whose clock has exceeded its wait. Every active frame advance updates that actor's baseline; equality/overflow at sprite count restores frame -1 and clears its flag. A tiny mask selector 0x200 can set BABA's flag, but the frame -1 guard prevents that OR alone from restarting a completed episode. Exact operands are in [clocks](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/clocks.tsv).

The HORSE painter uses action index 0/1/2 for A1/A2/A3. It requests Horse2.wav at A1 frame 14 and A2 frames 8/14. It requests Horse1.wav at A3 frame 1 and Horse2.wav at A3 frame 14. Horse3.wav is loaded but has no request in the complete selected painter. The selected one-shot helper prevents a new start while the handle is already playing; playback is unobserved.

Birds1..Birds9 each contain 57 frames. An inactive bird branch starts when unsigned time since retained baseline exceeds 1000+R(2000), range 1000..2999 ms for the initial and later waits. It chooses group R(3), amount 1+R(3), and indices group*3 through group*3+amount-1, corresponding to a prefix of Birds1..3, Birds4..6 or Birds7..9. Each starts frame 0. Bit 0x80 admits one increment per generic step. The painter draws selected sprites at view origin while their frame is below 57, then draws Town_add at the same origin. It clears bit 0x80 only when all selected sprites have completed; even that paint draws Town_add. The bird-paint branch updates its baseline on every active paint. A one-bird group requests Birds1.wav; two/three birds request Birds2.wav.

Frame counts are from the complete selected [art frame envelopes](../experiments/EXP-2030-rom2-first-town/evidence/measured/art/frames.tsv). Schedule, reset and paint conditions are in the [scheduler](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/scheduler-en.txt), [bird clock](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/bird-clock-en.txt), [bird painter](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/bird-paint-en.txt) and per-actor advance listings.

**Confidence.** High for the finite loaded populations, position tables, measured frame counts and selected schedule/advance rules. The complete loader loops, inspected table contents, strict elapsed comparisons, bit receivers and envelope enumeration exclude art-only fixed positions, a single fixed actor action, identical BABA/DERVISH variants and birds painted above their occluder. EN/RU are not separate runtime witnesses.

**Unknown.** The actual selected variants/actions, random sequence, redraw cadence, re-entry state, active sound handles and whole-image Horse3 requests remain unobserved. Runtime capture would settle the live state; a complete sound-pointer-use sweep would settle global Horse3 use. The three loaded-only graphics stay bounded by `R2-ENGINE-240`.

## First-town rooms


| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-245 | Campaign town ID 1 selects the generic inn and shop in preserved ROM2 EN/RU; selected EN generic and Kaarg constructors share native page or child bases. | High | ● active (branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ENGINE-246 | The selected ROM2 EN generic inn fixes three panel rectangles, primary art layers and three button targets; complete portrait/stat drawing and the lower-right resource binding remain Unknown. | High / Unknown | ● active (branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ENGINE-247 | The selected ROM2 EN inn consumes packed DLL options, seats actors and dispatches talk by NPC/topic; the stage-10 roster is conditional, and its complete live actor-art join is Unknown. | High / Medium / Unknown | ● active (branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ENGINE-248 | Selected ROM2 EN inn branches load candle, cauldron, breath and drink lists and advance them on native clocks; container-count interpretation and the complete visible schedule remain bounded. | High / Medium / Unknown | ● active (branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ENGINE-249 | The selected ROM2 EN generic shop fixes five child rectangles, primary art and Undo/Buy/Sell/Exit targets; complete stock, item-price text and episode state remain Unknown. | High / Medium / Unknown | ● active (amended, branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |
| R2-ENGINE-250 | Selected ROM2 EN inn/shop departure releases and removes the square; conditional Exit/Escape posts re-enter it and rebuild local art/state, while complete queue admission and visible return are unobserved. | High / Medium | ● active (partially retracted, branch candidate) | [EXP-2030](../experiments/EXP-2030-rom2-first-town/) |


### R2-ENGINE-245

R2-ENGINE-237 supplies the positive campaign room selectors in the preserved EN/RU clients. Town ID 1 selects the generic shop and inn; town ID 2 selects their Kaarg variants. EN application members are +0xf8 and +0x104 for the generic pages, and +0x100 and +0x10c for Kaarg. EN campaign mode is application+0x5d8=2.

The selected EN generic inn constructor is `L2.00607`; its center constructor is `L2.00610`. The generic shop page constructor is `L2.00735`; its art child is `L2.00617`. The Kaarg page constructors call the generic page bases; its derived center/child constructors call those generic center/child bases. Both inn builders instantiate left panel `L2.00736`, right panel `L2.00737` and seat builder `L2.00738`.

These positive calls establish shared native receivers. Per-town center/art methods still provide different backgrounds, overlays and animation episodes. The following room layout contracts measure selected EN bodies. They do not establish complete EN/RU room geometry equality or differences limited to art.

The [inn page entry slots](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-page-entry-slots.tsv) separate roster entry from shared controls. Generic table `L2.00739` uses enter +0x80 `L2.00552`; third table `L2.00740` uses `L2.00684`; Kaarg table `L2.00741` uses `L2.00613`. Kaarg's own [enter](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/kaarg-inn-enter-0049e451.txt) calls actor resolution. All three tables use Escape +0x6c `L2.00742`. The generic, third and Kaarg builders call right-panel constructor `L2.00737`; right-panel action `L2.00743` calls talk `L2.00554`. Shared buttons, talk and exit do not establish a shared roster enter body.

Selected evidence: [class pointer cells](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/class-slots.tsv), [generic inn constructor](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-page-ctor-0049a428.txt), [generic inn builder](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-page-slot78-0049a65e.txt), [generic shop constructor](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/shop-page-ctor-004b6bb2.txt), [generic shop builder](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/shop-page-slot78-004b85e2.txt). R2-ENGINE-237 carries the selected Kaarg constructor calls and locale selectors.

**Confidence.** High for positive campaign selection and the named shared-base constructor calls. Direct-CFG boundaries agree with independent linear decoding; this is a selected native contract, not a runtime observation.

**Unknown.** Complete RU layout/event-body equality and all state-dependent room presentation. The selected EN bodies and positive EN/RU selectors do not close these alternatives. RU room painters and original observation would settle them.

### R2-ENGINE-246

Generic inn builder `L2.00744` adds the left, right and center children in that order. Coordinates below are its passed corner operands for the 640x480 page:

| Child | Constructor | ID | Corners |
|---|---|---|---|
| Left | `L2.00736` | 0x44d | (0,0,160,480) |
| Right buttons | `L2.00737` | 0x44e | (480,0,640,238) |
| Center | `L2.00610` | 0x450 | (160,0,480,480) |

Enter `L2.00552` attaches the existing campaign hero panel at `(640-panelWidth,0)`. Its contents depend on campaign state.

The selected loaders name these archive keys under `graphics.res:interface/inn/`; native literals prefix `graphics/`: `LeftStats.bmp`, `LeftPicture.bmp`, `ButtonsArea.bmp`, `CenterArea.bmp`, `manback.bmp`, `ManBackTalk.bmp`, `LUOver.bmp`, `LDOver.bmp`, `RUOver.bmp`, `button1on.bmp` through `button3on.bmp` and `button1off.bmp` through `button3off.bmp`. The quest-art loop is noncampaign-only.

Left painter `L2.00745` draws LeftStats at (0,0) and LeftPicture at (0,238), then selected actor material. It formats `graphics/infowindow/%s.bmp`, but its complete live portrait/stat join is unclosed.

Center painter `L2.00746` draws CenterArea at (160,0), candle at (160,48), cauldron at (420,160), conditional tender art at (240,152), then hire actors followed by talk actors, each back before its actor sprite. It then draws LUOver at (160,0), LDOver at (160,238), RUOver at (464,0), and a conditional shared lower-right image at (464,238), cropped to 16x242. This last branch selects global `L2.00747` or `L2.00748` using the shared hero panel's +0x6c state. The key-to-global binding is not established. Shared child painting follows.

Right painter `L2.00749` returns when page+0x110 is zero; otherwise it draws ButtonsArea at (480,0). Button hit rectangles are:

| Index | Corners | Caption/action |
|---|---|---|
| 0 | (484,44,624,90) | Selected hire: Hire or Fire; selected talk: blank |
| 1 | (484,91,624,137) | Talk |
| 2 | (484,138,624,184) | Exit |

Refresh `L2.00750` uses main-text keys 258/259 for Hire/Fire according to option bit 31 and key 242 for Talk. The initial generic caption uses key 243 Hire/Fire; Exit is key 232. Release `L2.00743` requires the same pressed/released index. Index 1 sends a selected talk entry through `L2.00554`, or formats `npc%dabout` for a hire entry. Index 2 calls the exit receiver `L2.00751`; index 0 takes hire/fire branches.

Selected evidence: [builder](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-page-slot78-0049a65e.txt), [art pushes](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/resource-pushes.tsv), [center painter](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-center-slot2c-00497adf.txt), [left painter](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-left-paint-selected-00494832.txt), [right builder](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-right-build-00496172.txt), [caption refresh](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-buttons-refresh-00497223.txt), [button action](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-button-action-00496bba.txt), [text keys](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/main-text.tsv).

**Confidence.** High for the selected EN corner operands, primary keys, painter order, caption keys and local button targets. Unknown for complete portrait/stat drawing and lower-right resource identity; no complete visible-room equality is asserted.

**Unknown.** The selected left painter does not close the fresh actor portrait/stat state; the center painter exposes conditional globals without a closed resource initializer. Follow the corresponding actor state and global initialization, or obtain original observation. No complete RU geometry claim follows from shared art payloads.

### R2-ENGINE-247

In EN campaign mode, generic enter `L2.00552`, bound by `L2.00739` slot +0x80, copies DLL `EnterInn` options into page+0xfc in DLL order. Option kinds 1 and 2 resolve into hire vector +0xc0; other kinds resolve into talk vector +0xe8. Actor resolution calls `R2.0071`. The initial selected index is 0 for a nonempty combined roster, otherwise -1. This option-copying body is not attributed to Kaarg; its enter override is R2-ENGINE-245.

The stage-10 options in R2-ENGINE-161 and R2-ENGINE-073 are NPC 207/topic 9/kind 0, NPC 2108/topic 8/kind 0 and NPC 517/topic 10/kind 3. They are talk entries in that order. This is a conditional roster: shared continuation can append state-gated options, and this experiment did not execute a full fresh-bank admission sequence.

Seat builder `L2.00738` creates three rows of six 48x64 rectangles. For row `r=0..2` and column `c=0..5`, corners are `(176+48*c,480-64*(r+1),224+48*c,480-64*r)`. Hire entries precede talk entries. With only those three stage-10 options, their seats begin at (176,416), (224,416), (272,416).

Hit receiver `L2.00752` subtracts the page origin, checks populated seats in order and returns the first matching index or -1. Center slot +0x54 (`L2.00753`) calls that hit and left selection `L2.00754`; selection writes page+0xb8, refreshes buttons and requests the Helper sound at page+0xa8. Pointer update `L2.00755` calls its selection slot and sends selected talk entries to `L2.00554`. Complete input gesture classification remains bounded by unclosed virtual dispatch.

Talk `L2.00554` indexes the talk vector with `selected-hireCount`. It first-matches the actor's u16 at +0x1dc against copied option low 16 bits through `L2.00556`, extracts topic bits 16..27, formats `npc%dtalk%d`, dispatches text through `R2.0038`, then calls DLL `TalkTo` with the packed option. R2-ENGINE-220 supplies the broader positive talk chain.

Actor-art builder `L2.00756` formats `graphics/interface/inn/Unit%d/sprites.16a` from actor+0x24. Talk actor flag bit 0 instead selects HeroMage when bit 1 is set, otherwise HeroFighter. Newly allocated frame-index cells initialize to zero through `L2.00757`. Missing-lookup fallback `R2.0069` initializes flags+0x1b8=0x56, actor+0x24=1, actor+0x28=NPC-ID-80 and u16+0x1dc=NPC-ID. A missing-map composition therefore chooses Unit1, but this does not prove the fresh live map misses the three actors.

The installed EN `scenario.res:npc.reg` is 448 bytes and has five records: Multiplayer, FacesMF, FacesMM, FacesFF, FacesFM. It has no row for those three NPC IDs; that bounded registry result does not establish global absence of campaign actors.

Selected evidence: [enter](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-page-slot80-0049abfe.txt), [seats](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-seats-build-0049783c.txt), [hit](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-pointer-hit-00499367.txt), [selection](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-actor-select-004946b6.txt), [talk](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-talk-0049ba28.txt), [art selection](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-center-slot78-00499783.txt), [fallback](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/npc-fallback-actor-0041d8bb.txt), [frame initialization](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-animation-index-init-0059fb96.txt), [registry records](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/npc-registry-records.tsv).

**Confidence.** High for the selected EN consumption, seating, selection and key/call chain. Medium for composing the named stage-10 roster with the fresh campaign route. Unknown for its complete live actor-to-art join and all continuation admissions.

**Unknown.** The fresh controlled actor map was not reconstructed, and the selected registry lacks the necessary per-NPC data. Readonly closure of fresh scenario/map construction through `R2.0050`/`R2.0071`, or original observation, would settle the art. Full fresh-bank inputs and continuation predicates would settle roster completeness. No runtime click or dialogue presentation was observed.

### R2-ENGINE-248

Enter `L2.00552` constructs four filename lists and copies them into the generic inn center in campaign mode 2. Archive keys are under `graphics.res:interface/inn/`; native literals prefix `graphics/`:

| List | Center offset | Files | Loaded list count | Destination |
|---|---|---|---|---|
| Candle | +0x280 | `candle/t0000.bmp`..`t0009.bmp` | 10 | (160,48) |
| Cauldron | +0x2b0 | `cauldron/t0000.bmp`..`t0020.bmp` | 21 | (420,160) |
| Breath | +0x2e0 | `tender/breath/br0001.bmp`..`br0024.bmp` | 24 | (240,152) |
| Drink | +0x310 | `tender/drink/dr0001.bmp`..`dr0040.bmp` | 40 | (240,152) |

Painter `L2.00746` draws current candle/cauldron art before their updates. Unsigned elapsed time greater than 100 ms advances each once and resets the baseline to the current clock; there is no catch-up loop. Helper `L2.00685` computes `(cursor+1) % L`, where `L` is getter `L2.00758`, exactly `container field+0x8-1`. This measured expression is not identified with the filename-list count.

Tender state 0 is idle, 1 drink and 2 breath. Elapsed time greater than the retained delay selects drink for an odd delay, with forward direction, or breath for an even delay, requesting chair and shop Breath sounds. Delay resets use signed truncation of the return from `L2.00729` divided by 16, plus 3000. R2-ENGINE-242 bounds that source to 0..32767, so the reset delay is 3000..5047 ms. Uniformity and episode probabilities are unestablished.

An active episode advances once after elapsed time exceeds 83 ms. Forward helper `L2.00686` increments until cursor equals `L`, then returns zero; reverse helper `L2.00759` decrements until zero, then returns zero. Drink changes to reverse after forward failure and ends after reverse failure, requesting glotok. Painted drink cursor 30 requests drink.wav. Breath failure returns to idle and resets its cursor to zero. Steam requests occur after elapsed time exceeds 10000 ms.

The painter advances only the selected seated actor after elapsed time exceeds 125 ms; other seats retain their frame index. The selected index wraps using the sprite's reported frame count. Sound loader `L2.00760` names the corresponding inn/shop sound keys; requests do not prove audible playback.

Selected evidence: [filename loops](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-page-slot80-0049abfe.txt), [painter clocks](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-center-slot2c-00497adf.txt), [wrap](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/animation-next-wrap-00401658.txt), [container getter](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/animation-count-004019c0.txt), [forward](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/animation-next-stop-004015c5.txt), [reverse](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/animation-previous-stop-00401612.txt), [sound loader](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-page-slot88-0049bb47.txt).

**Confidence.** High for filename loops, fixed destinations, native timer comparisons and helper expressions. Medium for a composed visible schedule because container loading and retained live state are unclosed. Unknown for the relation of `L` to loaded list counts. The RNG range is bounded by R2-ENGINE-242; its distribution remains unobserved.

**Unknown.** Loader `L2.00761` was not followed; the raw field-minus-one getter does not settle whether a final loaded picture is skipped. Its allocation/append contract would settle that relation. Selected helper/caller listings do not establish redraw cadence, pre-entry baselines, RNG distribution or audible playback. Native runtime observation would settle their presentation.

### R2-ENGINE-249

Generic shop builder `L2.00762` adds these children using passed corner operands:

| Child | Constructor | ID | Corners |
|---|---|---|---|
| Shop inventory | `L2.00763` | 0x3eb | (0,303,480,390) |
| Hero inventory | `L2.00764` | 0x3e9 | (0,390,480,480) |
| Item information | `L2.00765` | 0x3ea | (0,0,164,303) |
| Main shop art | `L2.00617` | 0x3ed | (164,0,480,303) |
| Buttons | `L2.00766` | 0x3ee | (464,0,640,238) |

Enter `L2.00767` attaches the existing campaign hero panel at `(640-panelWidth,0)` and joins current character/item state. Those fixed child rectangles do not imply a fixed initial stock list.

Native paths prefix archive keys with `graphics/` or `movies/`, for graphics.res or movies.res respectively. Main-art loader `L2.00768` names `graphics/interface/ShopFrame.256`, `graphics/interface/shopanim/ShopMain.bmp` and `movies/shopanim/Pose2-3/1.bmp`. Painter `L2.00769` draws the frame at (164,0), ShopMain at (169,8), enabled category layers 0..3 at their stored rectangle origins, then the chosen keeper frame at (277,112), then shared children. The constructor's layer origins make category 0 (353,108), 1 (197,108), 2 (313,20) and 3 (201,20) in page coordinates. Shelf loader `L2.00770` formats `graphics/interface/shopanim/%.2d/%d.bmp` from category/episode fields. Keeper methods name Pose2-3, Yes and No movie series.

Support loader `L2.00771` names `graphics/interface/{myitem.256,shopitem.256,backinvg.bmp,backinvb.bmp,backinvs.bmp}`, `costs1.bmp` through `costs7.bmp` and `costm1.bmp` through `costm7.bmp`. Button loader `L2.00772` appends page slot +0x94's prefix and `ShopButton1.bmp` through `ShopButton4.bmp` under `graphics/interface/`. Generic prefix method `L2.00773` constructs the default native string through `L2.00774`; an empty prefix is an inference because its underlying storage initializer was not closed. Kaarg supplies a per-town prefix.

| Index | Hit corners | Text key | Local action target |
|---|---|---|---|
| 0 | (494,15,614,67) | 72 Undo | `L2.00775` |
| 1 | (483,67,623,113) | 70 Buy | `L2.00776` |
| 2 | (483,114,623,160) | 71 Sell | `L2.00777` |
| 3 | (494,160,614,212) | 73 Exit | `L2.00778` |

The four-entry native click table is `L2.00779`. Cases clear the pressed index to -1 and set page+0x14c bit 0x20. Hit, art and text rectangles differ. These local calls do not establish every transaction side effect.

Selected item traversal `L2.00780` enumerates state-supplied item objects. Inventory, hero and information-panel painters are included, with unresolved switches and virtual edges preserved. Complete initial stock, all item sprite/name/price destinations, applicable buttons and episode state are not established. R2-ASSET-060 bounds the unresolved ShopFrame pixel decoding.

Selected evidence: [builder](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/shop-page-slot78-004b85e2.txt), [art-child constructor](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/shop-child-ctor-004b9cf5.txt), [art loader](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/shop-child-slot78-004baa98.txt), [painter](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/shop-child-slot2c-004bb3d7.txt), [support art](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/shop-support-art-004b796f.txt), [button constructor](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/shop-buttons-ctor-004bc2b0.txt), [button tables](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/fixed-switches.tsv), [item traversal](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/shop-item-iteration-004b5e2f.txt), [unclosed edges](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/bodies.tsv).

**Confidence.** High for selected EN constructor geometry, primary art operands, text keys and finite local button targets. Medium for complete composition and the generic empty-prefix inference. Unknown for the full initial stock, item-price text layout and complete live episode state.

**Unknown.** The selected constructors, panels and traversal do not close the town-1 stock source or every item-panel virtual edge. Follow them on a known readonly initial campaign state, or obtain original observation. Complete item sprite/name/price drawing requires those receivers; no catalog-wide absence or stock identity is claimed.

**Amended.** R2-ENGINE-263 through R2-ENGINE-267 answer the town-1 stock, price and item-text Unknowns; R2-ENGINE-268 and R2-ENGINE-269 answer the keeper episode state; R2-ENGINE-270 decodes the ShopFrame pixels.

### R2-ENGINE-250

Shop and inn click arms call square slot +0xac `L2.00781`, which sends 0x445 to its own slot +0x48 `L2.00726`. The handler forwards 0x445 to `L2.00782`. Square enter calls base enter `L2.00783`, which sets +0x5c=1. The shared message receiver admits virtual leave +0x84 `L2.00784`: it clears active+0x204, removes child+0x200, calls art release +0xa4 and container release +0x8c, then base leave `L2.00785`. The receiver posts 0x44c with the square pointer. Application arm `L2.00786` calls removal `L2.00787`; `L2.00788` removes the square from application+0xcc through `L2.00789`. Town removal branch `L2.00790..L2.00791` adds no return dispatch. The square is released and removed while the room is open.

Inn Escape `L2.00742` tests key 27 and calls Exit `L2.00751`. Shop Escape `L2.00792` and Exit click `L2.00793` call `L2.00778`. These send current-room removal 0x445; the shared receiver posts 0x44c and removal drops the room. A separate conditional post returns to town: shop Exit and Escape post 0x42e only when application+0x404=2; inn Exit, also reached by Escape, posts it only when +0x404=4. With another value these bodies send no 0x42e. This is not an unconditional Exit return guarantee.

The selected [application switch](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/application-message-index-0048539d-004853d0.txt) subtracts 0x416, bounds the index by 0x73, reads byte table `L2.00794` and jumps through DWORD table `L2.00795`. Its [measured cells](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/application-message-arms.tsv) map 0x44c to `L2.00786` and 0x42e to `L2.00796`. The 0x42e arm calls town dispatcher `L2.00266`, which re-adds member+0x110 through `L2.00797` and calls enter +0x80 `L2.00710`. Enter unconditionally calls loader +0xa0 and container setup +0x88.

Re-entry reloads square art and reruns HORSE R(5), BABA R(4) and DERVISH R(4), retrying DERVISH until its variant differs from BABA. It selects their positions again; a new draw may still choose the previous variant. Per-window flags+0x208 clear, then DERVISH bit 0x400 is set with frame 0. Tavern, sign, stars, shopie and vane initialize their frame selections; BABA/HORSE return to idle frame -1 with first-action pointers, new clocks and 2000..3999 ms initial waits. Gate initialization yields cursor 8 and T08, then clears gate direction latch+0x1a4. The guard returns to its last sprite frame, 7, with step+0xec=0 and sound-direction latch+0xf0=0. Selector+0xb4 resets to -1. The painter's global baseline and lazy clock state, bird delay and stars blank counter are not cleared by the selected loader/enter/release. Re-entry does not reset every animation variable.

Selected evidence: [square leave request](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/leave-prep-en.txt), [square message](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/message-en.txt), [square leave](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/leave-en.txt), [square enter](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/enter-en.txt), [square loader](../experiments/EXP-2030-rom2-first-town/evidence/measured/square/loader-en.txt), [base enter](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/base-enter-004dc232.txt), [shared message receiver](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/room-stack-message-004dc2b0.txt), [common removal](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/room-remove-common-004891d8-004892ab.txt), [town removal](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/room-remove-town-0048957b-004895b3.txt), [shop Exit](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/shop-button-click-case3-004bdb2d.txt), [shop Escape](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/shop-page-slot6c-004b70a2.txt), [inn Exit](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/inn-choose-0049b69b.txt), [0x42e arm](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/room-message42e-dispatch-0048637d-0048638d.txt), [town dispatcher](../experiments/EXP-2030-rom2-first-town/evidence/measured/rooms/town-dispatcher-0048bbc0.txt).

**Confidence.** High for the selected EN departure, release/removal, conditional Exit/Escape posts and local re-entry/reset instructions. Medium for their composed input-to-visible campaign return because complete queue and virtual input admission were not executed or closed.

**Unknown.** Application+0x404 writers and campaign values, subsequent behavior when these bodies send no 0x42e, original input delivery, complete queue acceptance and visible return timing remain unmeasured. A bounded +0x404 writer and no-post continuation trace would settle those alternatives. Complete virtual dispatch closure or original observation would settle admission and visible return.

**Amended.** The retained-square clause is refuted and partially retracted in [retracted.md](retracted.md). Departure releases and removes the square; the measured conditional 0x42e route re-enters it. The local room-removal clause stands.

## ROM2 town classes and druid square

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-251 | The three selected ROM2 square classes have 44 virtual slots; druid and Kaarg constructors call the generic town constructor, which calls the shared native page base. | High | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |
| R2-ENGINE-252 | Selected ROM2 town overrides mix resource data with distinct native behavior: town 1 has a campaign gate, Kaarg reversible gate frames, and druid hover-started people and sprite selector tables. | High / Medium | ● active (amended, branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |
| R2-ENGINE-253 | The active druid square draws background, shop highlight, woman, inn highlight, man, lizard and bug before inherited overlay and child dispatch; person subset changes also change destinations. | High | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |
| R2-ENGINE-254 | The druid view inherits the shared ROM2 mask sampler; changed shop/inn hover can immediately start person subset 2, while its two dialogue selectors do not start an episode. | High | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |
| R2-ENGINE-255 | Selected druid click arms post shop, inn, navigation and menu messages or stage-keyed keeper dialogue; its complete selected loader, painter and advance contain no gate frame operation. | High | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |
| R2-ENGINE-256 | The druid square advances once after more than 100 ms and schedules each paint; people require idle state, bug uses three 61-frame routes, and lizard uses four sentinel-terminated selector vectors. | High | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |

### R2-ENGINE-251

EN town-ID construction binds generic/druid/Kaarg constructors
L2.00559/L2.00561/L2.00560 to tables L2.00562/L2.00564/L2.00563.
Each table has 44 slots, +0x00 through +0xac. The two derived constructors
call generic L2.00559. Generic calls native page constructor L2.00634,
which assigns table L2.00651 with 34 slots through +0x84. Slot identity
is recorded separately from a method's interpreted purpose.

Evidence is the complete slot matrix, constructor assignments and selected
method excerpts in evidence/measured/classes. The base-to-derived comparison
uses observed table cells, not a source-language class name or imported symbols.
Shared native methods include pointer forwarding +0x4c, leave +0x84,
mask sampling +0x94 and departure preparation +0xac. A matching pointer
proves method reuse; it does not prove identical data, state or presentation.

**Confidence.** High for the finite table population and positive constructor
chain. Every selected table cell is retained with its address. The native
function instrument controls selected decode boundaries. No global class census
or complete inherited window semantics is asserted.

**Unknown.** Source class names and unmeasured shell/window callback purposes.
Native observation or additional receiver tracing would settle delivered behavior.

### R2-ENGINE-252

The complete matrix in evidence/measured/classes distinguishes reused pointers,
data overrides, logic overrides and callbacks whose purpose remains unresolved.
Druid differs from the generic table at twelve slots:
04,2c,54,80,88,8c,90,98,9c,a0,a4,a8. Kaarg also changes tip slot 14.
The lifecycle and sound methods are kept separate from paint and input.
Sound-entry slot 90 uses the same helper but a different druid sound field.
Resource-release counts and keys are data. Kaarg sound loading omits the
pre-load virtual release used by generic and druid; this changes control flow.
Art release L2.00652 (druid) and L2.00653 (Kaarg) call a subset of generic
L2.00654 cleanup primitives. Generic also calls L2.00655, L2.00656 and
L2.00657. The data label does not assert identical teardown programs.

Generic's selected gate click tests Scenario variable 0x301 before navigation.
Kaarg has reversible gate frames and four timed person families
(R2-ENGINE-235, R2-ENGINE-236). Druid starts shop/inn person subset 2 from
hover and uses separate bug and lizard sprite rules (R2-ENGINE-254,
R2-ENGINE-256). These named native branches prevent a claim that every
town difference is a resource key or constant.

**Confidence.** High for positive table differences and named instructions.
Medium for the data/logic labels: a changed address alone does not establish
a different algorithm, and shared pointers still consume per-instance state.
Classification is bounded by the selected methods, not whole-client equivalence.

**Unknown.** Complete semantic equivalence of all inherited callbacks, wider
shell behavior and live delivery. The slot matrix names unresolved purposes.

**Amended.** R2-ENGINE-281 bounds the effect of Kaarg's omitted pre-load release: each slot load releases its own slot first, so no different sound state results.

### R2-ENGINE-253

EN painter L2.00658 / RU L2.00659 returns if +0x204 is zero. Destinations
add the view origin. Background is (0,0), shop highlight (420,224), inn
highlight (0,184), lizard (0,300), bug (0,212). Woman subset s=0..2 uses
(336+4q,244-20q); man uses (164-12q,200-16q), where q=trunc(s/2).
Idle hover 1 selects woman subset 2/file 1 at (340,224); otherwise subset
0/file 1 at (336,244). Idle hover 2 selects man subset 2/file 1 at
(152,184); otherwise subset 0/file 1 at (164,200).

The shop highlight is drawn when hover=1 or woman subset=2. The inn
highlight is drawn when hover=2 or man subset=2. Thus an episode can
retain a building highlight after the pointer leaves. Lizard draws its
selector value minus one when route and cursor are nonnegative, otherwise
frame 0. Bug draws 61*route+cursor only when both values are nonnegative.
The inherited overlay dispatcher then invokes empty own slot 30 and children.

Evidence is evidence/measured/square/painter-*.txt, geometry.tsv and
lizard-sequences.tsv. The full selected painter and its direct branches
are controlled. Native sprite binding is separately bounded by R2-ASSET-066.

**Confidence.** High for local draw order, arithmetic and conditions.
**Unknown.** Device conversion, wider children, shell text/cursor destinations
and live frames. Source-RGB reconstruction does not establish original pixels.

### R2-ENGINE-254

Druid slot 94 points to the same sampler as generic and Kaarg. It samples
the top-origin indexed byte at (x-left)+640*(y-top), rejecting absent mask
or an out-of-view point with -1. Its full 256-byte map is retained.
The actual druid population is R2-ASSET-064.

EN hover L2.00660 / RU L2.00661 stores the selector. For shop 1 or inn 2,
changed admission and a current subset other than 2 stop the family sound,
reset its clock, select subset 2, store cursor -1, set its flag and call
advance immediately. This produces cursor 0.
Ddruid1 at +248 is stopped and replayed by hover selector 1 (L2.00662);
Ddruid2 at +24c by selector 2 (L2.00663). Scheduler person start arms
stop/replay the same keys at L2.00664/L2.00665 and L2.00666/L2.00667.
Shop/inn use latched Denter2/Denter1 and stop the other entry and Dout sound.
Selector 8
requests Dout under its latch. -1 calls reset and clears three latches.
0x200 and 0x1000 are no-op arms. The inherited slot 4c calls virtual 98
outside the painter too; hover is not limited to the 100 ms paint path.

Evidence is evidence/measured/square/{mask-map,branches,anchors}.tsv and
mask-sampler, hover, inherited-hover, person-advance, scheduler and sound-load
excerpts.

**Confidence.** High for the complete byte domain and selected finite hover
table, plus local person state writes. No global input-handler census is implied.
**Unknown.** Physical pointer capture/delivery and audible sound.

### R2-ENGINE-255

EN click L2.00668 / RU L2.00669 dispatches selectors 1 and 2 to messages
0x42a (shop) and 0x42b (inn). Selector 8 posts 0x442(1,0), then 0x42d.
These three arms call departure preparation first. Selector 16 posts
0x41f (main menu) without it. This menu arm has no mask pixel in the
preserved druid masks (R2-ASSET-064).

Selectors 0x200 and 0x1000 format druidinnkeeper%d and druidshopkeeper%d
with Scenario variable 0x300 and call dialogue entry R2.0038. They do
not construct new room classes. The inherited tip returns indices
233/236/237/235 for shop/inn/navigation/menu, and no text for the two
keeper selectors. The selected complete loader, painter and advance have
no gate frame family or gate frame operation. Navigation posting and the
hover sound remain present; this is not a global absence of gates.

Evidence is evidence/measured/square/{click,loader,painter,advance,tip}-*.txt,
tip-indices.tsv and the retained room dispatch. R2-ENGINE-257 identifies
the per-town inn/shop selectors.

**Confidence.** High for local click arms, strings, message values, tip table
and the bounded gate-frame negative. Whole selected methods and their
direct/table targets are retained.
**Unknown.** Acceptance after navigation posting, dialogue contents and final
tip rendering. Trace the native receivers or observe an authorized session.

### R2-ENGINE-256

EN painter L2.00658 admits one virtual advance only when unsigned elapsed
time is greater than 100 ms. It resets its process baseline to current time,
without a catch-up loop. Scheduler L2.00670 runs on every active paint.
Define R(n)=floor(rand()*n/32767)%n on the shared 15-bit helper already
measured by R2-ENGINE-235. This establishes bounds, not uniformity.

Woman and man enter with 2000+R(2000) ms waits. Later strictly elapsed waits
replace the baseline and wait with 3500+R(5000). They start only if their
subset is -1: hover 1/2 selects subset 2, otherwise R(2), cursor 0,
flag 1/2. Each admitted active advance increments once; equality with
the selected vector count resets cursor to 0, subset to -1 and clears the flag.
The native loader attempts files 1..20 for each of three subsets; failed
lookup ends that subset early. R2-ASSET-065 bounds the installed population.

Bug starts after 7000+R(5000) ms, then 10000+R(10000). The start arm
selects R(3), sets flag 0x80 and updates the baseline. It does not test
active state or reset the cursor. Entry cursor -1 becomes 0 on the first
advance; equality with 61 resets cursor 0, route -1 and clears the flag.
Later starts therefore use the retained cursor. A slow redraw cadence can
reschedule a running route without rewinding it.

Idle lizard starts each paint with R(15)+1: values 1/2/3 choose routes
1/2/3; values 4..15 choose route 0. Cursor starts 0, flag 0x100.
Four explicit vectors have lengths 8/11/25/47 followed by -1. Their
one-based values, including repeated holds, are reduced by one for drawing.
One advance increments cursor; reaching a -1 cell ends the episode.

The sound loader selects 21 town_druid keys. Bird/tree waits are separately
2000+R(2000); wolf requests after more than 60000 ms. Bug/lizard starts
request route sounds. These are requests, not observed playback.
Evidence is evidence/measured/square/{constants,lizard-sequences,controls,
branches,resource-pushes}.tsv and complete entry, painter, hover, scheduler,
advance, person/bug/lizard advance and sound excerpts.

**Confidence.** High for the selected native predicates, order, integers,
finite tables and frame state transitions. All selected methods decode
completely and branch/table targets are retained. EN/RU is one code dependency.
**Unknown.** Delivered cadence, PRNG sequence, process baseline before re-entry,
sound delivery and live episode appearance. Authorized observation would settle
the delivered timing; no game execution occurred in this experiment.

## ROM2 room classes and druid behavior

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-257 | Selected ROM2 inn/shop pages and centers reuse native room bases; per-town overrides mix data and control flow, while five changed shop cleanup methods retain equivalent native algorithms. | High / Medium | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |
| R2-ENGINE-258 | Selected druid inn methods add water and taverner state transitions over shared widgets; ten water names feed a nine-cursor loop, and taverner cursor reset preserves its cached bitmap. | High / Medium | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |
| R2-ENGINE-259 | Selected druid shop methods gate the fourth category through campaign query 0x302 and use a modulo-30 keeper counter; native paint, click and selection differ from the shared shop base. | High / Medium | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |
| R2-ENGINE-260 | Selected inn/shop pages reuse shared native child classes; EN construction binds app+0xe0 to a panel with 33 nonzero targets, while changed callback purposes and its null boundary remain Unknown. | High / Unknown | ● active (branch candidate) | [EXP-2031](../experiments/EXP-2031-rom2-town-views/) |

### R2-ENGINE-257

Campaign selectors choose inn pages at application +0x104/+0x10c/+0x108
and shop pages at +0xf8/+0x100/+0xfc for IDs 1/2/3. Original EN/RU
constructors bind fifteen inn page/center and shop page/center/header tables.
The room matrix has 1868 slot rows across both locales, including shared
side widgets, inventories and actions. Constructor and boundary controls
retain every callable target and its next cell. Native window table
L2.00671 ends at +0x74 before the next table head at L2.00672.

Inn pages 2/3 change 04,78,80,84,88,8c; centers change 04,2c,80,84.
Shop pages change 04,78,80,88,8c,90,94; leave 84 is inherited.
Shop centers change 04,14,2c,54,78,7c,80,84,88,8c,90,94,98,9c,a0,a4.
Headers change 04 and b4. Five center cleanup overrides at 90,94,98,9c,a0
are identical after normalizing local branch targets in both locales.
A different code address alone is not a different native algorithm.

For these ten derived classes, each locale has 396 slots: 326 inherited,
37 data overrides, 23 logic overrides and 10 equivalent overrides.
Both locales have 652 inherited, 74 data, 46 logic and 20 equivalent rows.
Inn pages each have 4 data/2 logic; inn centers 3 data/1 logic; shop pages
5 data/2 logic; shop center 2 has 4 data/7 logic/5 equivalent; shop center 3
has 5 data/6 logic/5 equivalent; headers each have 2 data/0 logic.

All ten derived +04 wrappers have their base wrapper's 17 instructions,
changing only the destructor call target. This is a data override of the
wrapper. Derived destructor bodies were read only for the headers:
EN druid L2.00673 forwards to base L2.00674; the other derived room
destructor bodies remain unclassified. Shared widget and shell +04
listings have the same deleting-destructor wrapper shape; window +04
semantics remain Unknown. Shop-center +14 EN L2.00675 and L2.00676
change only the text-index base from base L2.00677's 0x3e to 0x116 and
0x112. RU L2.00678 and L2.00679 have the same single difference.
The four-category rectangle scan is the same loop as the base.

Evidence is evidence/measured/rooms/{class-slots,constructor-bindings,
boundary-controls,equivalent-controls,range-controls}.tsv and their selected
native listings. The experiment room table separates resource data, logic,
equivalent overrides and unidentified callback purposes. Room data includes
keys, frame vectors, constructor choices, sound fields and keeper moduli.
Paint, entry, schedule and category selection carry per-town branches.

**Confidence.** High for positive bindings, pointer reuse, boundary controls
and normalized native cleanup equivalence. Medium for the whole semantic
data/logic taxonomy and whole room equivalence: shared widget message paths
and every callback purpose were not closed. This is a selected class census,
not an executable-wide enumeration of all possible rooms.

**Unknown.** Other derived room destructor bodies, unlabelled callbacks,
complete Kaarg state schedules, town 1 room measurements assigned separately
and live presentation. Read the
named destructor bodies and native callbacks or observe an authorized
original route.

### R2-ENGINE-258

Druid page L2.00680 and center L2.00681 call the shared inn bases and use
shared sides plus the campaign panel. Center painter L2.00682 draws
TavernMain, water at center+(168,144), a1/a2 at +(40,128), static a30001
at +(104,152), shared portraits/quest presentation, edges and children.
The center is page+(160,0). Art loader L2.00683 loads druid background
and static taverner with generic manback/ManBackTalk/LUOver/LDOver/RUOver.
Campaign entry L2.00684 binds ten water files, forty a1 and thirty a2.
R2-ASSET-065 separately measures their installed headers and keys.

Water admits advancement after more than 100 ms. Helper L2.00685 uses
(cursor+1) mod (loaded-count-1), so its steady cursors are 0..8 for ten
loaded names. A1/a2 use one-shot helper L2.00686 through cursors 39/29.
Reset L2.00687 changes the cursor without changing the cached bitmap.
An episode can therefore first paint its previous terminal image before
advancing to file 2. The idle wait rand()/16+3200 is 3200..5247 ms on
the measured 15-bit source. Expiry chooses (wait&3)+1 with 4 mapped to 3.
State 3 stores a new wait and changes to 4; the next admitted >100-ms
state-4 tick returns to idle. Its stored wait is not a simple static-frame
duration. Hover, tooltip and click targets are inherited.

Evidence is evidence/measured/rooms selected entry/art/painter listings,
frame-binding/count/loop/once/reset helpers, resource pushes and finite
table controls. Matching RU methods and bindings are retained.

**Confidence.** High for bounded native keys, destinations, state arithmetic
and frame rules. Medium for the complete room composition because global
lower-right edge art and every shared child presentation were not identified.
**Unknown.** Lower-right global bitmap identities, physical hover, audible
delivery and live timing. Trace their loaders or observe authorized output.

### R2-ENGINE-259

Druid center L2.00688 and header L2.00689 reuse shop/widget bases.
Loader L2.00690 selects ShopFrame.256, ShopMain, elven and four highlights
under interface/shop_druid, with movies/shop_druid/a10001 as default.
Header loader L2.00691 selects four arrows and ShopInv. At center origin
page+(164,0), painter L2.00692 draws frame (0,0), background (5,8),
admitted elven (93,172), selected armor (5,112), magic (5,52) or potion
(121,36), keeper (197,92), then children.

Predicate L2.00693 / RU L2.00694 admits outside campaign. In campaign it
queries 0x302 through the EN dynamic Scenario pointer L2.00451 and requires nonzero.
Original startup binds ordinal 1, ScenarioGetVar at DLL D2.00001, which
reads bank[argument]. Thus 0x302 is bank770 at D2.00215. Selected ordinary
Leave70 reaches case D2.00216 and stores bank770=1 at D2.00217. It enables
this gate absent an intervening writer; no complete producer census follows.
The predicate gates elven paint, category-3 click and category-3 selection.
Tooltip L2.00676 scans all four category rectangles without that predicate.

Keeper ticks require at least 100 ms. Each tick samples a fresh
5000+1000*(rand()%5) idle threshold, rather than retaining one deadline.
An admitted idle start selects flag 0x10 or 0x20. Updater L2.00695 deletes
the current image, increments counter modulo 30, clears episode flags on
zero and formats counter+1 for a1/a2 or a10001 otherwise. Shared constructor
initializes counter/flags to zero. An admitted episode first updates to file
2, reaches file 30 and returns to default file 1. Kaarg updater L2.00621
uses modulus 20 and default a10000, with separate fire selection transitions.
Five cleanup overrides are equivalent native algorithms (R2-ENGINE-257).

Evidence is evidence/measured/rooms selected gate/art/paint/click/tooltip/
selection/frame/header listings, native resource pushes and controls.
R2-ASSET-065 gives the installed art population, not a clock witness.
Evidence/measured/query retains the startup/export/getter binding, Leave70
selector cells and decoded writer excerpts.

**Confidence.** High for selected native gate, counter, update order and
draw rules. Medium for whole shop composition and data/logic equivalence:
all shared shell/button resources and sibling message paths were not closed.
**Unknown.** Complete bank770 campaign progression and restoration, pending dialogue, shared
shell art, audible delivery and live pixels. Additional scenario producers
and authorized native observation would settle those boundaries.

### R2-ENGINE-260

The selected inn pages share left/right classes and shop pages share player
inventory, sell list and actions, each with positively bound native tables.
Their differing targets against window/inventory bases are shared room
methods, not per-town variants. Full addresses and inherited relations are
in evidence/measured/rooms/class-slots.tsv.

EN application construction allocates 0x17c, calls L2.00696 at L2.00697
and stores its result at application+0xe0 at L2.00698. Constructor L2.00696
calls window L2.00635 and binds table L2.00699. This panel prefix has 33
nonzero targets at 00..80: 18 same-offset inherited targets, twelve changed
targets and three extensions against the 30-target window prefix. Cell 84
is zero, followed by a new table head at 88. Evidence is the EN constructor
excerpts, slots.tsv, boundaries.tsv and bindings.tsv in measured/panel.

**Confidence.** High for positive EN panel binding, finite pointer values,
shared room references and same-offset target equality. Unknown for the
changed panel method purposes and zero cell's role. It may be reserved null
or alignment; this probe does not label it a callable slot. No per-town
character-panel class is established by this construction.

**Unknown.** Panel callback algorithms, null boundary semantics and final
shared-widget presentation. Read the identified targets and consumers or
observe an authorized original session. The panel probe is EN only.

## First-town shop

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-263 | ROM2 town-1 shop stock is parameterized by the Scenario.dll ordinal-15 record: four categories with prices 0..1500, draws 100/100/20/20, quantity bounds 2/2/1/1 and fixed masks. | High / Unknown | ● active (amended, branch candidate) | [EXP-2032](../experiments/EXP-2032-rom2-first-town-shop/) |
| R2-ENGINE-264 | Town-1 category masks admit 70 armor/shield rows, 44 weapon rows, 4 enchantable weapon rows and 5 books plus scroll rows per fill from data.bin tables, by price, class and material. | High | ● active (branch candidate) | [EXP-2032](../experiments/EXP-2032-rom2-first-town-shop/) |
| R2-ENGINE-265 | A town-1 shop fill draws random admitted candidates with stackable quantities, preloads six potions and the books, and enchants category 2 into 160 possible spell-level variants. | High | ● active (branch candidate) | [EXP-2032](../experiments/EXP-2032-rom2-first-town-shop/) |
| R2-ENGINE-266 | ROM2 shop buy moves affordable pending items to the hero, sale and undo return items into categories, and refills replace all lists only while no deal is open; finite stock is implied. | High / Medium / Unknown | ● active (branch candidate, partially retracted) | [EXP-2032](../experiments/EXP-2032-rom2-first-town-shop/) |
| R2-ENGINE-267 | ROM2 shop text and prices come from text tables, item attributes and data.bin factors; cells show unit P or (P+1)/2, pending totals P*q and ceil(P/2)*q, sale credit floor(P*q/2+0.5). | High / Medium | ● active (branch candidate) | [EXP-2032](../experiments/EXP-2032-rom2-first-town-shop/) |
| R2-ENGINE-268 | The ROM2 generic shop keeper updates after 100 ms, resamples a 5000+1000*(rand()%5) ms idle threshold, plays Pose2-3 through counter 28 and responses through counter 12. | High / Unknown | ● active (branch candidate) | [EXP-2032](../experiments/EXP-2032-rom2-first-town-shop/) |
| R2-ENGINE-269 | Generic shop responses start the keeper's Yes on category change, affordable buy or credited sale and No on an unaffordable buy; first category choice and exit clear all episode bits, not the counter. | High / Medium / Unknown | ● active (branch candidate) | [EXP-2032](../experiments/EXP-2032-rom2-first-town-shop/) |
| R2-ENGINE-270 | ROM2 EN draws the town-1 ShopFrame with byte reader L2.00798 and palette mode 1: low six bits count, top bits select palette literals, row skip or pixel skip. | High / Unknown | ● active (branch candidate) | [EXP-2032](../experiments/EXP-2032-rom2-first-town-shop/) |

### R2-ENGINE-263

Scenario.dll export ordinal 15, `ScenarioGetShopAssortment` (`D2.00220`),
returns `D2.00104 + location*0x50`, with the location ID read through
`D2.00011` at location+4. The record has four 0x14-byte category entries:
minimum price, maximum price, draws, quantity bound and mask.
`ScenarioNewGame` (`D2.00004`) writes town 1 at `D2.00221`:

| Category | Min | Max | Draws | Quantity bound | Mask |
|---|---|---|---|---|---|
| 0 | 0 | 1500 | 100 | 2 | 0x1381cc03 |
| 1 | 0 | 1500 | 100 | 2 | 0x10418103 |
| 2 | 0 | 1500 | 20 | 1 | 0x28018100 |
| 3 | 0 | 1500 | 20 | 1 | 0x04000000 |

`ScenarioSave` (`D2.00102`) and `ScenarioLoad` (`D2.00103`) copy every
location's entries verbatim. A linear census of every .text operand in
`D2.00104..D2.00007` in both DLLs finds no other town-1 writer; the other
writers index the location-2 and location-3 records by category times 0x14.
EN and RU DLL exports, bodies and records are equal.

The client binds ordinal 15 once at startup: EN `L2.00799` stores the pointer
in `L2.00800`, RU `L2.00801` in `L2.00802`. Fill `L2.00803` and sold-item placement
`L2.00804` call it only while `[L2.00805]+0x74` is zero; otherwise both read
the deal source at +0x94+0x70.

**Confidence.** High for the record, the town-1 writes, the copy bodies,
the census population and the client binding. Unknown for when the deal
source replaces the record.

**Unknown.** When `[L2.00805]+0x74` is nonzero in a single-player campaign; a census of the writers of that field would settle it.

**Amended.** R2-ENGINE-334 reads the writers of `+0x74`: the server
constructor writes 0 and init writes (mode < 2) (High); campaign start
passing mode 2, which leaves it 0, and network paths setting it to 1 are
Medium. R2-ENGINE-329 and R2-ENGINE-330 give the other location records.

### R2-ENGINE-264

Kind builder `L2.00806` decodes the mask: bits 0..14 select materials (bit 13
is material 14, bit 14 material 15), bits 15..21 classes 0..6, bit 29 an
enchanted category and bit 28 an extra 50 percent enchant gate with bit 29.
Kind 0x400000 admits weapon table `L2.00807` in mode 2 (column15 bit 0 set);
0x1000000 armor `L2.00808` in mode 1; 0x800000 shield `L2.00809` in mode 7;
0x8000000 with bit 29 weapon mode 8 (column15 bit 0 clear); 0x4000000 the
magic admission `L2.00810`. Data.bin tables load from owner `L2.00811`; armor,
shield, weapon, magic and spell rows start at index 1.

Table admission `L2.00812` requires the row's class mask word to carry the
material bit and price trunc(column2 x material factor x class factor) in
[min, max]; an enchanted category skips the range and keeps rows with
min^0.4 at most the material x class level product; with min 0 that is
every row, taking the C library's pow(0, 0.4) as 0. Armor mode excludes
(class, material, row) triples (0,0,2), (0,1,2) and (6,4,6).

| Category | Admitted from both installed data.bin files |
|---|---|
| 0 | 57 armor and 13 shield rows; 8 rows above 1500; 2 excluded triples |
| 1 | 44 weapon rows, classes 0..1, materials 0, 1 and 8 |
| 2 | 4 weapon rows: classes 0 and 1, material 8, rows 13 and 14 |
| 3 | books for spells 1, 5, 10, 16 and 26 at 1000; per fill one of MagicItems rows i+5 or i+34 for each i in 1..29, admitted at column0 in range |

Books cover spell IDs 1..29 except 9, 14, 15, 24, 28 and 29, priced by
spell column 21. Of the 58 phase-two rows, 23 are admissible and all are
type 4 (Scroll or SuperScroll prefix). EN and RU populations are equal.

**Confidence.** High. The pow(0, 0.4) reading is an inference about the
CRT helper. Exact rational admission products equal host double
products for every candidate and none lies within 1e-9 of an integer.

### R2-ENGINE-265

Fill `L2.00803` clears four lists and runs draw `L2.00813` per category. The
magic category first clones each admitted book at quantity 2 and adds six
named potions (MagicItems rows 69..74) at quantity 51..100 regardless of
price. Each draw picks index (rand()*n)>>15 over the n candidates, retrying
up to 1000 times on a book. A category without bit 29 accepts the pick.
Quantity is 1+((rand()*(bound+1))>>15) for a stackable item (type 3 or 4, or
no effects) and 1 otherwise; merge `L2.00814` adds quantities of equal
stackable items; sort `L2.00815` orders the list.

Category 2 enchants each pick (`L2.00816`, `L2.00817`). Budget is
min(2*max - P, 100*P), with P the weapon reader price (base product plus
0.5, truncated): 167, 83, 333 and 167. The spell is uniform over table
`L2.00818` {1, 10, 11, 18, 26, 5, 16}, redrawn up to 100 times while
`L2.00819` returns -1. The level cap is trunc((1.2^log2(budget/(10*c20))
- 1)*30), capped by weapon level field +0x48 (20 or 40) and 100; spells 11
and 18 (column20 5000 and 10000) give -1. Effect slot +0x54 (`L2.00820`)
draws the level uniformly in 1..cap. Effect slot +0x4c (`L2.00821`) prices
the tag-41 effect at trunc(10*c20*2^(ln(1+L/30)/ln 1.2)); the weapon
reader adds it. A draw outside [min, max] fails; failures stop the loop
after 10*draws. In both locales 160 of 687 (base, spell, level) rows
survive: levels 1..8, 1..9, 1..7 and 1..8 for each of five spells, final
prices 649..1442.

**Confidence.** High for the rules and the replayed population. Host
doubles replace the x87 ln/pow helpers; every kept or rejected price lies
more than 1e-9 from an integer, except level 6, where both logarithms take
the same double argument and the ratio is exactly 1.

### R2-ENGINE-266

Shop object `L2.00822+0x6c` holds an open-deal count at +4, a pending-refill
counter at +8 and four category lists. Fill runs only when +4 <= 0.
Enter `L2.00823` (message 0x32) fills when the category-0 list is empty.
Message 0x3f (`L2.00824`, only when `+0x74` is zero) refills through
`L2.00825`; the client sends it from campaign start `L2.00826` and its 0x468
handler `L2.00827`. With a deal source, tick `L2.00828` adds a pending refill
every 180 ticks, applied by `L2.00829` when no deal is open, including on
leave `L2.00830`.

Buy (0x33, `L2.00831`) walks the pending deal list and stops at the first
shop item whose q*P exceeds hero gold; each bought item debits q*P, takes
the hero as owner and enters the hero inventory. Sell (0x34, `L2.00832`)
credits trunc(q*P*0.5+0.5) for each hero item with nonzero P, clears its
owner and places it through `L2.00804` by the rule of R2-ENGINE-333;
`L2.00833` merges into an equal stackable
entry or a quantity-0 placeholder or appends. Server cases 0x33 and 0x34
reach buy and sell through `L2.00834` and `L2.00835`. Undo (0x35) runs
`L2.00836`, `L2.00837`, which sets the deal's customer through `L2.00838`, then
`L2.00839`: pending items with an owner go to the hero inventory, items
without one go through `L2.00804`, and the pending list is cleared. Leave
removes quantity-0 entries whose +0x4c byte is zero.

Finite stock is implied, not read as one body: a deal references shop
items and increments their +0x4c byte, bought items leave with the hero,
and leave drops unreferenced quantity-0 entries. The producer that moves
a selection into the pending list is the Unknown below.

**Confidence.** High for these bodies, including undo. Medium for finite
stock as an inference from them, and for the campaign meaning of the 0x468
senders and location-type gate. Unknown for how a selection moves
part of a shop stack into the pending list, and whether the client lists
survive SAV.

**Unknown.** The selection producer of the pending list and the stock's
SAV round trip; reading the client transfer message and the save writer
would settle them.

**Amended.** The sold-item fallback clause is partially retracted in
claims/retracted.md; R2-ENGINE-333 gives the placement rule. Buy, sell
credit, undo and refill clauses stand.

### R2-ENGINE-267

Name: `main.res` `text/itemname.txt` lines join in order with `world.res`
`data/itemname.bin` IDs (491 IDs, templates and lines in each locale) into
name map `L2.00840`. Description `L2.00841` walks the item attribute
sequence: tag 1 is the price and is not described; default tags use
`stats.txt` key equal to the tag with `#%s %d`; tags 13 and 44..48 form
damage ranges; tags 41 and 42 use `spell.txt` key value-1 with `main.txt`
keys 5a..5d.

Price: armor, shield and weapon slot +0x4c (`L2.00842`, `L2.00843`, `L2.00844`)
give trunc(column2 x material factor x class factor + 0.5); a tag-1 effect
replaces the price; tag-41 effects add their own price (R2-ENGINE-265);
other effects add trunc(50*S*(1+(S/70)^1.5)); the total caps at 19999999.
Magic items use MagicItems column 0; books use spell column 21. Shop paint
`L2.00845` prints one unit price per item cell: (P+1)/2 for an item with
+0x18 == 2, else P, with P from wrapper `L2.00846`; it does not multiply by
quantity. Pending totals come from `L2.00847` (RU `L2.00848`), which buy
and sell responses call first: page+0x150 takes hero gold, page+0x158 sums
((P+1)/2)*q for +0x18 == 2 items, page+0x154 subtracts P*q for the others,
and page+0x15c holds the sum of the three. Server sale credit
(R2-ENGINE-266) is floor(q*P/2+0.5); for odd P it is floor(q/2) below the
client's ((P+1)/2)*q. The measured quote, preview and settlement bodies
apply no town or campaign multiplier.

**Confidence.** High for text bindings, keys and price arithmetic in the
measured consumers. Medium for the whole rendered description, whose
helper-derived spell clauses were not executed.

**Unknown.** A price factor outside the measured consumers; a census of the callers of the price slots and the price wrapper would settle it.

### R2-ENGINE-268

Painter `L2.00769` (RU `L2.00849`) runs only when page+0x148 is nonzero.
Two static baselines start at now-100 ms and now on first use and are
shared by all centers. A paint call updates once when
unsigned(now - update baseline) >= 100, then sets the baseline to now.
Each update samples 5000+1000*(rand()%5) ms and sets the random-episode bit
0x10 when unsigned(now - episode baseline) reaches it and bits 0x10, 0x20
and 0x40 are clear.

Bit 0x10: slot +0x88 sets counter = (counter+1) % 30 and loads Pose2-3
file counter+1; at counter 28 the bit clears, the counter resets and the
baseline restarts, and the resting file 1 is drawn. Bits 0x20 and 0x40
increment the counter and end at 12, freeing the Yes or No array. Arrays
hold Pose2-3 file 1 at index 0 and Yes or No files 2..12 at 1..11. Update
priority is 0x10, 0x20, 0x40.

**Confidence.** High for the static update and draw rules. Unknown for live
timing and overlapping-event appearance.

**Unknown.** Native presentation; original observation would settle it.

### R2-ENGINE-269

Category click slot +0x54 calls slot +0xa4, which returns 1 for a changed
valid category; the click then sets bit 0x20 and loads the Yes array. Buy
response `L2.00776` needs nonzero pending cost page+0x154; when
page+0x150 + page+0x154 is negative it sets bit 0x40 and plays the refusal,
else bit 0x20 and sends the purchase. Sell response `L2.00777` sets bit
0x20 when pending credit page+0x158 is nonzero. The page sound clock never
sets episode bits.

Two writers clear the whole flag word +0x240, including bits 0x10, 0x20
and 0x40. Page slot +0x80 sets the category word page+0x132 to 0x64 on
entry; when slot +0xa4 finds 0x64 (EN `L2.00850`, RU `L2.00851`) it zeroes
+0x240 and +0x244..+0x250 and calls slot +0xa0, so the first category
choice ends a running episode before setting 0x20. Page slot +0x84 (EN
`L2.00852`, RU `L2.00853`) zeroes center +0x240 on exit. Neither writes
counter +0x254; in the committed listings only the constructor, the
increments, the modulo-30 advance and the three episode ends write it. A Yes episode started by a response therefore
starts from the counter value it inherits.

**Confidence.** High for these paths and the two flag-word writers.
Medium for completeness: no executable-wide census of center flag and
counter writers was made. Unknown for a Yes or No episode started at
counter 12 or more: the 0x20 and 0x40 arms end only at counter 12 and draw
array index counter.

**Unknown.** The outcome of an episode started at counter 12 or more; a
read of the array bounds and allocation, or original observation, would
settle it.

### R2-ENGINE-270

Shop loader `L2.00768` builds the byte sprite through `L2.00643` and sets
one palette row, mode 1 and zero colour adjustment through `L2.00644`.
Mode 1 (`L2.00854`) reads B, G and R from each four-byte entry and packs a
WORD with the active channel widths and shifts. Sprite table `L2.00855`
slot +0x18 is `L2.00856`; normal reader `L2.00798` reads one-byte
commands: low six bits are the count; top bits 00 copy that many palette
indexes, 01 skip rows keeping the column, 10 and 11 skip pixels. A row ends
when the column reaches the width. The reader stops by height, not by
dataSize.

**Confidence.** High for the selected EN bodies and their exact stop at the
installed frame's data end in both locales. Unknown for RU reader bodies
and device output.

## Town-1 gate navigation and tips panel

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-271 | ROM2 message 0x442 sets a patch.txt label, names game(9999-wParam).sav and calls the save driver: mode word 2 saves and appends the label, other modes send the name in a type-7 record; the gate gives game9998.sav. | High / Medium | ● active (branch candidate) | [EXP-2033](../experiments/EXP-2033-rom2-town-gate/) |
| R2-ENGINE-272 | ROM2 message 0x42d leaves a current town through LeaveLocation and opens the global map; leaving town ID 1 sets a map mode that routes to and enters the first available location. | High / Medium | ● active (branch candidate) | [EXP-2033](../experiments/EXP-2033-rom2-town-gate/) |
| R2-ENGINE-273 | ROM2 global L2.00713 is the TipsMode option: the options constructor sets 1, and registry load, the options Tips checkbox and the tips-panel checkbox write it; the client has no other writer. | High | ● active (branch candidate) | [EXP-2033](../experiments/EXP-2033-rom2-town-gate/) |
| R2-ENGINE-274 | ROM2 town child 0x467 is a 312x200 tips panel framed from lm.256 with town.txt section #tips1, a Close button and a checked Show-tips checkbox at fixed panel-relative rectangles. | High / Medium | ● active (amended, branch candidate) | [EXP-2033](../experiments/EXP-2033-rom2-town-gate/) |
| R2-ENGINE-275 | The ROM2 tips panel takes mouse input: Close posts 0x45a, which makes the town square delete the panel; the checkbox state is stored into TipsMode and decides the panel at later town entries. | High / Medium | ● active (branch candidate) | [EXP-2033](../experiments/EXP-2033-rom2-town-gate/) |

### R2-ENGINE-271

The application message dispatcher EN `L2.00262` / RU `R2.0036` subtracts
0x416 and indexes a byte table and a dword table (EN limit 0x73, RU 0x78).
Message 0x442 resolves to EN `L2.00857` / RU `L2.00858`; both arms have 40
instructions with equal mnemonics.

1. wParam 0 copies `patch.txt` zero-based line index 55 into app+0x148; any
   other wParam copies index 95. Pool `L2.00859` holds `patch.txt`; the pool
   load appends lines from the first line on, and the getter adds index n to
   the pool base. The EN lines are "Restart last mission" and "Abort mission
   and return to town".
2. `game%d.sav` is formatted with 9999-wParam into app+0x248.
3. The arm calls save driver EN `L2.00860` / RU `R2.0166`, the
   single-player save driver of R2-SESSION-021. Its first test compares
   app+0x5d8 (RU +0x63c) with 2.
4. timeGetTime goes to app+0x424; app+0x40c and app+0x428 become 0 (RU
   +0x43c, +0x424, +0x440).

The two branches of the driver:

- **Mode word 2.** The driver calls `L2.00861` on global `L2.00805` with
  the app+0x248 name. It then opens that file without the create flag,
  moves to its end through SetFilePointer and writes the 256 bytes at
  app+0x148 through WriteFile (`L2.00862`, `L2.00863`, `L2.00864`). The
  label is appended at the end of the named file. The body ends with a
  jump past the other branch.
- **Any other value.** The `jne` goes to EN `L2.00865` / RU `L2.00866`. That
  branch passes the app+0x248 name to EN `L2.00867` / RU `L2.00868` on
  app+0xd0 and returns. The callee fills global record EN `L2.00869` / RU
  `L2.00870`: byte +9 = 7, word +5 from the word at app+0xd0's +0x9cc
  object +4, word +7 = 0, and the name at +0xf. It passes the record to EN
  `L2.00871` on object `L2.00872`. With word +7 zero, that body walks the
  list at +0x18b8 and calls the record's slot +8 for each element that
  passes its admission tests. The RU dispatch body `L2.00459` is paired by
  call operand; its instruction sequence differs from EN.

The town-1 gate posts 0x442 with wParam 1 (R2-ENGINE-241). In mode word 2
the driver therefore works on `game9998.sav` and appends the index-95
label to it.

**Confidence.** High for the switch cells, the arm, the branch structure
of the driver, the label append in mode 2 and the type-7 record fields in
both locales. Medium that the labels describe restore points for
"restart" and "abort to town", and that mode 2 is single-player campaign
play: the readers of the appended label and the writers of app+0x5d8 were
not traced.

**Unknown.** What the type-7 record does at its recipients, and which mode
values reach it. Which code loads `game9998.sav` or `game9999.sav`, and the
meaning of app+0x404 bits beyond 0x10. Reading the recipients' slot +8
bodies, the writers of app+0x5d8 and the load path of `game%d.sav` would
settle them.

### R2-ENGINE-272

Message 0x42d resolves to EN `L2.00873` / RU `L2.00874`; both arms have
37 instructions with equal mnemonics and operands. The map view is
app+0xf0.

1. GetCurrentLocation (ordinal 18) gives the current record. Map+0x1b0
   becomes 0.
2. If the record type is 2, map position +0x104/+0x108 is set from the
   record coordinates (town 1: 215,234) and LeaveLocation (ordinal 6) runs.
   An output of at least 0 plays `video\%s\%02d.smk` through EN `L2.00875`;
   leaving town 1 gives -1 (R2-SESSION-023). If the record ID is 1, map+0x1b0
   becomes 1.
3. EN `L2.00876` / RU `L2.00877` opens the map: it adds app+0xf0 to room
   stack app+0xcc, calls its enter slot +0x80 and sets bit 0x10 in app+0x404
   (RU +0x41c). It requests `music\map.wav` only when EN global `L2.00878` is
   nonzero (`L2.00879`). The map enter EN `L2.00880`
   loads `main\text\globalmap.txt`.

With map+0x1b0 nonzero, map paint EN `L2.00881` / RU `L2.00882` sets a
route from the current position to the coordinates of the first record of
GetAvailableLocations and sets +0x12c to 8. When the route ends it posts
0x468 and calls EnterLocation (ordinal 5) on that first record; with
map+0x1b0 zero it takes the selection path instead. The 0x468 arm EN
`L2.00827` calls `L2.00466` and `R2.0026` for a type-1 current record and
town dispatcher `L2.00266` (R2-ENGINE-231) otherwise.

Neither arm refuses navigation. A false flag 0x301 stops the gate click
before posting (R2-ENGINE-241). A current record that is not type 2 skips
the map position set, LeaveLocation, the cutscene and the town-ID test; a
town other than ID 1 opens the map without the route mode.
After the stage-10 inn talk the availability list holds mission 10
(R2-SESSION-110, R2-SESSION-023).

**Confidence.** High for both arms, the map open and the map-paint branch
in both locales. Medium that the town-1 gate leads to mission 10: the list
content is a static derivation, not an observed map. Medium for labelling
`R2.0026` the mission path.

**Unknown.** The map's appearance in route mode and the live route
duration. An authorized capture would settle them.

### R2-ENGINE-273

`L2.00713` (RU `L2.00883`). A raw scan of every byte start in every
section of both clients finds 21 references per locale: four writers and
seventeen readers.

| Writer | EN | RU | Value |
|---|---|---|---|
| options constructor | `L2.00884` | `L2.00885` | 1 |
| registry load | `L2.00886` | `L2.00887` | registry value `TipsMode` |
| options dialog on 0x445 | `L2.00888` | `L2.00889` | control 0xd state through checkbox get +0x3c |
| tips panel on 0x46e | `L2.00890` | `L2.00891` | lParam from control 0xf (RU 0x10) |

The options dialog binds control 0xd to `dialogs.txt` index 156 ("Tips")
only in mode 2; otherwise it binds `L2.00892` with index 175 ("Names &
Clans"). Registry store `L2.00893` writes the global back. Readers include
the town-1 enter `L2.00894` and the Kaarg enter `L2.00895`.

**Confidence.** High. The raw scan covers every byte start; each hit
decodes to the named instruction, and EN and RU bodies agree.

### R2-ENGINE-274

Town-1 enter EN `L2.00710` tests TipsMode at `L2.00894`. When it is
nonzero it reads `town.txt` section `#tips1` through `L2.00896`
(`#tips%d` with argument 1) and constructs class EN `L2.00633` / RU
`L2.00897` with ID 0x467 and rectangle (328,0,640,200). The Kaarg enter
builds the same class under the same TipsMode test (R2-ENGINE-238,
R2-ENGINE-273).

Panel init EN `L2.00898` / RU `L2.00899` builds three children from the
panel extent W=312, H=200:

| Control | ID (RU) | Class | Rectangle | Content |
|---|---|---|---|---|
| text | 0xd (0xe) | `L2.00900` | (20,24)-(W-28,H-36) = (20,24)-(284,164) | the section text |
| button | 0xe (0xf) | `L2.00901` | (W-120,H-40)-(W-40,H-22) = (192,160)-(272,178) | `main.txt` index 127, EN "Close"; message 0x45a |
| checkbox | 0xf (0x10) | `L2.00902` | (40,H-40)-(W-124,H-24) = (40,160)-(188,176) | `main.txt` index 128, EN "Show tips next time"; state 1 |

Text indices are zero-based line indices of the text pool (R2-ENGINE-271).
The EN `#tips1` section is 313 bytes and RU 284 bytes. Panel overlay EN
`L2.00903` draws frame pieces of `graphics\interface\lm.256` (loaded at
`L2.00904`): a shadow pass through slot +0x1c with pieces 0xc, 0xf, 0x11,
0x10 and 0xe, then corners 0xa, 0xc, 0xf and 0x11, edges 0xb and 0x10 every
48 pixels, edges 0xd and 0xe every 32 pixels and fill piece 9. Children
paint through `L2.00578`.

A child rectangle is parent-relative: `L2.00905` calls `L2.00906`, which
adds each parent's left and top along +0x30. Room stack app+0xcc is built
at `L2.00907` with rectangle (0,0,width,height) and no parent.

**Confidence.** High for the construction, rectangles, keys and frame
pieces; EN and RU bodies agree. Medium that the panel appears at view
origin plus (328,0): the live parent chain of the square is inferred from
these construction paths.

**Unknown.** Live pixels and text wrapping in the text control. An
authorized capture would settle them.

**Amended.** R2-ENGINE-314 answers the text wrapping from the text
control bodies; live pixels stay Unknown. R2-ENGINE-312 and R2-ENGINE-313
extend the panel class to every tip surface.

### R2-ENGINE-275

Panel message slot EN `L2.00908` returns 0 for 0x100, 0x445 and 0x446. For
0x202 it first calls slot +0x58; 0x202 and every other message then go to
the shared child dispatcher `L2.00782`.

- The Close button posts its message 0x45a to the application on release
  `L2.00909` or hotkey `L2.00910`. App arm 0x45a EN `L2.00911` removes and
  deletes mission-view child 0x10 (RU 0x11) when it exists; otherwise it
  calls slot +0x48 of room stack app+0xcc. The stack is built by window
  constructor EN `L2.00635`, whose vtable EN `L2.00671` / RU `L2.00912`
  has slot +0x48 EN `L2.00913` / RU `L2.00914`. That slot passes a
  non-mouse message to EN `L2.00915`, which offers it to each child in
  list order, without a rectangle test, until one returns nonzero. Town
  square slot +0x48 `L2.00726` receives it as a stack child and removes and
  deletes the child at view+0x200 and clears that field. Delivery to the
  square rather than an earlier stack child is an inference from that
  loop.
- Checkbox toggle `L2.00916` flips its state and sends 0x46e with its ID and
  state to the parent's slot +0x48. The panel stores the state in TipsMode
  (R2-ENGINE-273). A cleared box therefore skips the panel at the next town
  entry.

**Confidence.** High for the panel, button, checkbox, application, window
and square bodies in both locales. Medium that no stack child before the
square consumes 0x45a.

## Kaarg room schedules and square sounds

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-277 | The ROM2 Kaarg inn taverner picks a1..a5 by ((rand*21)>>15)+1 with no idle wait, runs each series once in >100 ms steps at fixed center positions and keeps its state in a process static. | High / Medium / Unknown | ● active (branch candidate) | [EXP-2034](../experiments/EXP-2034-kaarg-room-schedules/) |
| R2-ENGINE-278 | The ROM2 Kaarg shop fire steps once per 100 ms tick through 18 frames: idle loops lite or cicle, choosing weapons runs it up through burn to cicle, another category runs it back to lite. | High / Unknown | ● active (branch candidate) | [EXP-2034](../experiments/EXP-2034-kaarg-room-schedules/) |
| R2-ENGINE-279 | The ROM2 Kaarg shop keeper plays a1 or a2 files 2..20 under one modulus-20 counter; an idle start picks either, buy and sale start a2, and refusals and category clicks start nothing. | High / Medium / Unknown | ● active (branch candidate) | [EXP-2034](../experiments/EXP-2034-kaarg-room-schedules/) |
| R2-ENGINE-280 | ROM2 Kaarg square slots +0x20c..+0x24c hold 17 town_kaarg keys: three voice and four bird clocks, Kman1 after 45 s, two hover keys, the Kvox1 entry loop and six guard steps at nine guard cursors. | High / Unknown | ● active (branch candidate) | [EXP-2034](../experiments/EXP-2034-kaarg-room-schedules/) |
| R2-ENGINE-281 | Kaarg's omitted pre-load release leaves no different sound state: each slot load releases its slot first and leave releases the same 17 slots; only clock baselines and waits carry across visits. | High / Unknown | ● active (branch candidate) | [EXP-2034](../experiments/EXP-2034-kaarg-room-schedules/) |

### R2-ENGINE-277

Kaarg inn center painter slot +0x2c, EN `L2.00935` (RU `L2.00936`),
draws the taverner only in campaign mode (application +0x5d8, RU +0x63c,
equal to 2). Process static `L2.00937` (RU `L2.00938`) holds the state;
its image value is -1, idle. Positions are center-local.

| State | Series, vector | Files | Destination | Steps per episode | Source values of 32768 | Start sound |
|---|---|---|---|---|---|---|
| 1 | a1, +0x3cc | 0001..0015 | (72,88) | 15 | 1561 | +0x154 `Kman3.wav` |
| 2 | a2, +0x3fc | 0001..0025 | (72,88) | 25 | 1560 | +0x150 `Kman2.wav` |
| 3 | a3, +0x42c | 0001..0003 | (112,112) | 3 | 4682 | none |
| 4 | a4, +0x45c | 0000..0006 | (200,136) | 7 | 23405 | none |
| 5 | a5, +0x48c | 0000..0028 | (88,84) | 95 | 1560 | none |

Selection. When the state is -1 the same paint call zeroes center +0x4bc,
draws r = ((rand()*21)>>15)+1 through `L2.00939` and maps it through jump
table `L2.00940`: 1, 2 and 5 keep their value, 3, 4 and 6 give state 3,
and 7..21 give state 4. States 1 and 2 then request their start sound
through play-once `L2.00941`. No wait precedes the choice: the paint after
an episode ends starts the next one.

Draw and advance. States 1..4 draw the cached bitmap of their vector
(`L2.00942`); state 5 draws vector element `L2.00943`[+0x4bc]
(`L2.00944`). Static `L2.00945` takes the first paint's time. A paint
advances when unsigned(now - `L2.00945`) > 100 and then stores now. States
1..4 call one-shot `L2.00686`: below the last cursor it increments the
cursor and caches that element; at the last cursor it returns 0, the
state becomes -1 and reset `L2.00687` zeroes the cursor without changing
the cached bitmap. State 5 increments +0x4bc and ends when the table value
is -1. The table holds 95 element indexes before the -1, with repeated
holds; `taverner-a5-steps.tsv` lists them. A paint draws before it
advances, so each terminal image stays on screen for one step. A restarted
series first draws its cached previous terminal image.

Entry. Campaign entry `L2.00613` binds the five vectors through
`L2.00761`, which zeroes each cursor and caches element 0. It writes
neither `L2.00937` nor +0x4bc. The raw-dword census finds `L2.00937` only
in this painter (14 operands), so a running episode continues after a
re-entry from element 0 of its series, or for a5 from its stored step.

Relation to R2-ENGINE-258. Kaarg uses the same one-shot and reset helpers
and the same >100 ms step as the druid taverner, but has no idle wait: the
druid rand()/16+3200 wait and (wait&3)+1 choice have no Kaarg counterpart.
The painter stores rand()/65 in `L2.00946` and times in `L2.00947` and
`L2.00948` once; each has one raw-dword operand, the write, so none is
read through a direct operand; indirect access is not excluded. State 3 is a one-shot a3 episode, not a static pose.

Inn ambience. The same painter, under two further campaign-mode tests after
the non-campaign call `L2.00949`, requests
R(3)+1 of +0x15c, +0x160, +0x164 (`Kvox6`, `Kvox7`, `Kvox8.wav`) when
now - +0x4c8 > +0x4cc, and R(4)+1 of +0x140..+0x14c (`Kdish1`..`Kdish4.wav`)
when now - +0x4c0 > +0x4c4; each request stores 2000+R(2000) and now.
Constructor `L2.00609` initializes all four fields the same way. R(n) is
the R2-ENGINE-235 helper. Every selected EN body has an RU body equal to
it after normalizing addresses and application-field displacements.

**Confidence.** High for the selection arithmetic, state counts,
destinations, step and end rules, start and ambience requests. Medium for
re-entry continuation and for the static being written only here: the
raw-dword census does not exclude indirect access. Unknown for live
timing and for what the screen shows during a restarted series.

**Unknown.** Delivered paint cadence and audible output; an authorized
timed original session would settle both.

### R2-ENGINE-278

Kaarg shop center painter slot +0x2c, EN `L2.00950` (RU `L2.00951`),
paints only while page+0x148 is nonzero. Static `L2.00952` (RU
`L2.00953`) starts at now-100. A paint with unsigned(now - `L2.00952`) >= 100
runs one tick and stores now. Each tick runs the page sound clock and the
keeper (R2-ENGINE-279), then moves fire cursor +0x2fc by direction +0x300:

- direction nonzero: cursor += direction; at cursor <= 0 or >= 17 the
  direction becomes 0;
- direction zero: cursor + 1; 6 wraps to 0, and 18 or more becomes 12.

Art loader slot +0x78 `L2.00619` fills two 18-cell arrays from
`interface/shop_kaarg/fire/dark/` at +0x26c and `fire/select/` at +0x2b4:
cells 0..5 `lite1..6.bmp`, 6..11 `burn1..6.bmp`, 12..17 `cicle1..6.bmp`.
It zeroes cursor and direction. Shop entry `L2.00954` calls it on every
entry, so the fire starts at `lite1` in its lite loop.

Category select slot +0xa4 `L2.00955` returns 0 for an unchanged word.
Otherwise, when the previous category word page+0x132 is not 0x64,
choosing 3 (weapons) with cursor < 12 sets direction +1 and choosing
another category with cursor > 0 sets direction -1. Entry sets the word
to 0x64 and calls +0xa4(0), so the first choice sets no direction.
Direction +1 passes the burn cells and stops at 17, then the cicle loop
12..17 runs. Direction -1 runs down to 0, then the lite loop runs. A later
choice can reverse a transit. Burn cells appear only in transit.

The paint draws, center-local: ShopFrame (0,0), ShopMain (5,8), the
highlight of category word 0, 1, 2 or 3 at (5,8), (49,72), (49,8) or
(153,8), the keeper at (125,116), then the fire at (193,204): the select
cell when the word is 3, otherwise the dark cell. Word 0x64 draws no
highlight. EN and RU bodies are equal after normalizing addresses.

**Confidence.** High for these selected bodies, the arrays and the draw
order. Unknown for live timing and pixels.

**Unknown.** Delivered tick cadence; authorized original observation
would settle it.

### R2-ENGINE-279

In each tick of R2-ENGINE-278 the painter first calls page slot +0x88
`L2.00956`, then draws a new threshold 5000+1000*(rand()%5) ms at
`L2.00957`, one rand() per tick, also while an episode runs. When
unsigned(now - `L2.00958`) reaches that tick's threshold (RU `L2.00959`;
the static takes the first paint's time) and center flags +0x240 bits
0x10 and 0x20 are clear, it sets
flags |= 0x10<<R(2): 0x10 for 16385 and 0x20 for 16383 of 32768 source
values. Bit 0x40 is not tested. When flags & 0x30 the updater runs; a zero
counter after it restarts `L2.00958`. An idle start therefore comes at the
first tick whose elapsed time reaches its own threshold: rand()%5 takes
0..4 for 6554, 6554, 6554, 6553 and 6553 of 32768 values, so a tick admits
the start with about 0.2 of source values from 5000 ms, 0.4 from 6000 ms,
0.6 from 7000 ms, 0.8 from 8000 ms and always from 9000 ms. With one tick
per 100 ms, about 0.20 of starts fall on the 5000 ms tick, 0.936 by
6000 ms and 0.9994 by 6900 ms; delivered tick cadence is Unknown.

Updater slot +0x88 `L2.00621` frees image +0x12c through slot +0x9c,
sets counter +0x254 = (counter+1) % 20, keeps only flags & 0xf at 0, and
loads `movies\shop_kaarg\a1%04d.bmp` for 0x10 or `a2%04d.bmp` for 0x20
with counter+1, else `a10000.bmp`. 0x10 has priority. The painter draws
+0x12c at center (125,116) while flags & 0x30, else resting +0x128
(`a10000` from the art loader). An episode from counter 0 shows files
0002..0020 for 19 ticks; the 20th tick wraps, ends and restores the rest
image. No episode selects a10001 or a20001. The page sound clock requests
+0xb4 or +0xb8, both `Shop\Kman4.wav`, when bit 0x10 or 0x20 is set and the
counter is 1.

Triggers. The Kaarg action callback `L2.00960` (RU `L2.00961`) calls the
shared responses of R2-ENGINE-269: buy `L2.00776` sets 0x20 for an
affordable buy and 0x40 for a refusal; sell `L2.00777` sets 0x20 for a
credited sale. Their center calls +0x80 and +0x84 reach `L2.00962` and
`L2.00963` through slots +0x94 and +0x98; these only free twelve-cell
arrays +0x1e0 and +0x210 and load no art. Updater and painter test only
0x30, so 0x40 draws and advances nothing and stays set until a counter
wrap or a whole-word clear. Kaarg category click `L2.00964` sets no
episode bit in its own body; its callees `L2.00965` and page+0x68 slot
+0x34 were not read. A buy during an a1 episode adds 0x20, but a1 keeps priority
and both bits clear at its wrap.

Shared with R2-ENGINE-269: the buy and sell bodies and conditions; the
first-choice clear (+0xa4 with word 0x64 zeroes +0x240 and +0x244..+0x250);
inherited page exit `L2.00966` zeroing +0x240. Neither writes the counter.
An episode started after such a clear at counter c shows files c+2..0020.
The modulus bounds the counter to 0..19, so R2-ENGINE-269's counter-12
case has no Kaarg analogue. Different: one modulus-20 counter for both
series with an a10000 rest instead of a modulus-30 pose series and
Yes/No series ending at 12; an idle start choosing 0x10 or 0x20 instead of
0x10; no 0x40 test at idle start; no category-click episode; no No series;
statics separate from the generic painter.

**Confidence.** High for these bodies, slots and draw rules. Medium for
writer completeness: no executable-wide census of +0x240 and +0x254
writers was made. Unknown for live timing and appearance.

**Unknown.** Writers of the flag word or counter outside the selected
bodies; a census of center-field writers would settle them.

### R2-ENGINE-280

Kaarg square sound loader slot +0x88, EN `L2.00588` (RU `L2.00967`), binds:

| Slot | Key under `sfx\town_kaarg\` | Request |
|---|---|---|
| +0x20c, +0x210, +0x214 | `Kvox2`, `Kvox3`, `Kvox4.wav` | clock A, R(3) = 0, 1, 2 |
| +0x218..+0x224 | `Kbird1`..`Kbird4.wav` | clock B, R(4) = 0..3 |
| +0x228 | `Kvox1.wav` | square entry, looping |
| +0x22c | `Kenter2.wav` | hover selector 1, shop (R2-ENGINE-233) |
| +0x230 | `Kenter1.wav` | hover selector 2, inn (R2-ENGINE-233) |
| +0x234 | `Kman1.wav` | now - +0x260 > 45000 ms, then +0x260 = now |
| +0x238 / +0x23c | `Ksteps2` / `Ksteps21.wav` | guard cursors 5, 13, 21, 47 |
| +0x240 / +0x244 | `Ksteps1` / `Ksteps11.wav` | guard cursors 9, 17, 23, 51 |
| +0x248 / +0x24c | `Ksteps3` / `Ksteps31.wav` | guard cursor 31 |

Scheduler slot +0x9c `L2.00597` runs on every active paint
(R2-ENGINE-235). Clock A requests when unsigned(now - +0xb8) exceeds
static `L2.00968`; clock B when now - +0x25c exceeds static `L2.00969`.
Each static takes 2000+R(2000) on its first use, flagged in `L2.00970`, and
a new 2000+R(2000) after each request, when the baseline takes now. The
request choice is stored in +0xc0 for clock A.

Guard steps run while flag 0x800 is set. Cursor +0x1cc minus 5 indexes
65-byte table `L2.00971` into jump table `L2.00972`. An arm cursor
requests its first key when R(4) is nonzero and its second otherwise,
then sets latch +0x258; any other cursor clears the latch, so each arm
cursor requests once per pass. Table cells at cursors 55, 59, 63 and 69
also name arms, but guard advance `L2.00973` zeroes the cursor and clears
0x800 when it reaches the vector count, 55 frames (R2-ENGINE-235), so they
never run. The square paint calls advance at most once per paint and the
scheduler on every active paint after it, so the scheduler sees every
cursor and a guard episode requests nine steps.

Square entry `L2.00631` calls +0x88, zeroes flags +0x208 and finally
requests +0x228 through `L2.00974`, which passes Play flag 1 (looping) to
`L2.00975`. Play-once `L2.00941` passes 0. Both skip a slot whose buffer is
playing (`L2.00976`). Kaarg slot +0x90 `L2.00977` loops +0x74, which no
selected Kaarg body loads and view init `L2.00709` zeroes; of 32 EN
`call [reg+0x90]` sites, four lie in selected bodies (generic entry,
druid entry, shop entry and shop sound load) and none in a Kaarg square
body.

**Confidence.** High for the keys, conditions, waits and the reachable
guard cursors in the selected EN/RU bodies. Unknown for audible output:
mixing, volume and stops are not read.

**Unknown.** Physical playback; authorized original observation would
settle it.

### R2-ENGINE-281

Generic sound loader `L2.00978` and druid `L2.00979` call this +0x8c
first, then load 12 and 21 slots. Kaarg `L2.00588` loads its 17 slots
directly through `L2.00980` and zeroes hover latches +0x250, +0xac and
+0x254. `L2.00980` first calls release `L2.00981` on the same slot. The
release stops a playing channel (`L2.00976`, `L2.00982`), deletes the
object and zeroes the slot; the load then stores a new object for the key.
Kaarg release +0x8c `L2.00983` calls `L2.00981` on the same 17 slots,
+0x20c..+0x24c, and on nothing else. The omitted call would therefore
release exactly the slots that each load releases. Both orders end with
each of the 17 slots holding the object that constructor `L2.00984`
returns for its key; that constructor was not read. No selected body
starts playback before entry's `Kvox1` loop. Only the order differs:
Kaarg stops each old sound just before its replacement.

Inherited leave `L2.00784` calls art release +0xa4 and sound release
+0x8c; entry `L2.00631` calls +0xa0, then +0x88, then loops `Kvox1`. Every
visit therefore starts from released slots.

What does differ on re-entry is clock state. Statics `L2.00968` and
`L2.00969` (RU `L2.00985`, `L2.00986`) and fields +0xb8, +0x25c and
+0x260 are written only by the scheduler among the selected Kaarg view
methods and constructor chain (`L2.00559`, `L2.00560`, `L2.00709`). Entry,
loader, release and leave leave them unchanged, so after time away the
first scheduler pass can request a voice, a bird and `Kman1` together.

**Confidence.** High for slot equivalence and for the field census within
the selected bodies. Unknown for the first-visit values of +0xb8, +0x25c
and +0x260 and for what constructor `L2.00984` stores.

**Unknown.** Their first values; an executable-wide writer census or a
read of the view allocation would settle them.

## Character generator

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-283 | ROM2 New Game (0x425) shows pre-create, then on OK the detail screen; detail Accept (0x445) commits the hero, sets slots 776/781 and difficulty, then posts 0x42e to the town; Cancel and Back (0x446) step back. | High / Unknown | ● active (branch candidate) | [EXP-2035](../experiments/EXP-2035-rom2-chargen/) |
| R2-ENGINE-284 | The ROM2 pre-create screen has three difficulty, four hero, Cancel and OK rectangles under mask values 20..180 and one campaign name field; selected and hover states draw level, h-on, h-sel and pair art. | High | ● active (branch candidate) | [EXP-2035](../experiments/EXP-2035-rom2-chargen/) |
| R2-ENGINE-285 | ROM2 hero indexes 0..3 select Start_MF, Start_FF, Start_FM and Start_MM through code table {0,2,3,1}; defaults are hero 0, difficulty index 1 and npcnames 20; a hero click names npcnames 23/24/26/25. | High | ● active (branch candidate) | [EXP-2035](../experiments/EXP-2035-rom2-chargen/) |
| R2-ENGINE-286 | The ROM2 detail screen has stats, stat-sheet, skill, button and inventory children; EN Reset sets four 25s and pool 100, while RU Reset rebuilds the template with pool 0. | High / Unknown | ● active (branch candidate) | [EXP-2035](../experiments/EXP-2035-rom2-chargen/) |
| R2-ENGINE-287 | ROM2 attributes cost F(v)=trunc(0.349*1.15^(v-1)+0.5); plus needs the pool and v<45, minus v>15; the producer accepts 140-sum(F)>=0 else sets 25s; each Start row sums exactly 140. | High | ● active (branch candidate) | [EXP-2035](../experiments/EXP-2035-rom2-chargen/) |
| R2-ENGINE-288 | The ROM2 skill click tests four mask bytes: fighters sword/axe/mace/pike, mages fire/water/air/earth; bow and astral are never tested or drawn; the template main skill 1 is lit until a click. | High | ● active (branch candidate) | [EXP-2035](../experiments/EXP-2035-rom2-chargen/) |
| R2-ENGINE-289 | ROM2 generator tips are town.txt #tips8/#tips9/#tips10 on pre-create and #tips5 or #tips6 and #tips7 on detail under the tips mode; two cycles light controls; pre-create sets the select cursor, detail the default. | High / Medium | ● active (branch candidate) | [EXP-2035](../experiments/EXP-2035-rom2-chargen/) |
| R2-ENGINE-290 | The ROM2 producer sets the chosen school skill 20, skill 5 to 10, others 0, a skill weapon (the mage staff by sex) and HP/mana at max; recompute gives 142, 123, 29/157, 42/99 and speed 19, 19, 16, 16 for the templates before items. | High / Medium / Unknown | ● active (amended, branch candidate) | [EXP-2035](../experiments/EXP-2035-rom2-chargen/) |

### R2-ENGINE-283

Route, EN entries (RU in the experiment's alignment table). Main-menu
release `L2.01097` maps button index 0 to message 0x425. Handler
`L2.01098` loads town tips, calls the scenario DLL NewGame, enters
campaign mode 2 through `L2.01084` (application +0x5d8, RU +0x63c), sets
difficulty control +0x598 index 1, creates the client unit and shows the
pre-create view (table `L2.01099`) through `L2.01100`, which starts
`music\chrgen.wav`.

| Screen | Leaves by | Result |
|---|---|---|
| Pre-create | OK click or Enter with name length > 0 | 0x445: name, difficulty and code<<6 (0x40 mage, 0x80 female) stored; preview hero built from the template; detail view (table `L2.01101`) |
| Pre-create | Cancel click or Esc | 0x446: main menu (0x421) |
| Detail | Accept click or Enter | 0x445: application +0x480 gains bit value 4; in mode 2 commit `L2.00168`, SetVar(776, mage) and SetVar(781, female) through `[L2.00450]` (RU `[L2.00448]`), session difficulty = index+1, post 0x42e |
| Detail | Back click | 0x446: pre-create again |

Message 0x42e shows the town the current location selects (R2-ENGINE-231);
the fresh campaign enters town ID 1 (R2-SESSION-077). Town 1 is a square
screen (R2-ENGINE-239), so the generator assigns no town-1 map position.
Every selected EN body with an RU entry except pre-create OK is equal
after normalizing addresses or differs only in numeric operands
(application fields +0x18, mode 0x5d8→0x63c, allocation sizes +4).

Pre-create OK `L2.01102` (RU `L2.01103`) has the same mode-2 arm, the
New Game path, in both: the name length goes through `setg`, and only a
length above 0 plays the +0x1d4 sound and posts 0x445. The RU arm for
other modes also requires length 1, or 3 when application +0x3f4 and
+0x3bc are both nonzero, and answers a name that begins with a space or a
short name with a message window. Esc `L2.01104` (RU `L2.01105`) plays
the +0x1d8 sound and posts 0x446. The detail screen handles Enter only.
A byte search of each `.text` section finds `push 0x308` and
`push 0x30d` followed by `call [cell]` once each: EN `L2.01106` and
`L2.01107` through `[L2.00450]`, RU `L2.01108` and `L2.01109` through
`[L2.00448]`. Commit `L2.00168` returns 0 on a timeout; the Accept arm
ignores the result and posts 0x42e.

**Confidence.** High for the static route and messages. Unknown for live
timing, including the commit's wait for the server reply.

**Unknown.** The RU entries of the application command, new-game,
new-game mode, detail show, detail commit and cursor-loader bodies were
not located, so their RU equality is not measured; only the RU SetVar
pair is placed, by byte pattern. An RU call-site walk from the aligned
dispatcher would settle them.

### R2-ENGINE-284

Pre-create init `L2.01110`, loader `L2.01111`, hover `L2.01112`, click
`L2.01113` and paint `L2.01114`. Positions are screen coordinates.

| Control | Rectangle | Mask | Art by state |
|---|---|---|---|
| Difficulty 0..2 | (8,0), (296,0), (580,0), 48x72 | 20, 40, 60 | state bit 0 selected, bit 1 hover: `level{n}on`, `level{n}l`, `level{n}lon`; drawn only in mode 2 |
| Heroes 0..3 | (116,44), (180,44), (392,44), (456,44), 64x244 | 80, 100, 120, 140 | hover `h{n}on`; selected `h{n}sel` 268x340 at (112,44) for 0/1, (260,44) for 2/3; `h1sel2`, `h2sel1`, `h3sel4`, `h4sel3` while the pair neighbour is hovered |
| Cancel | (16,400), 64x76 | 160 | `cancell` while hovered or lit |
| OK | (548,400), 80x76 | 180 | `okl` while hovered or lit |
| Name | (300,433)-(464,449) | – | label main.txt 365 ten pixels left of the field |

Index n of the art is the hero index plus 1. Outside mode 2 a name field
at (300,422) and a clan field at (300,444), labelled main.txt 366, replace
the campaign field. Mask value 200 covers (173,404)-(480,478) and has no
arm in the selected bodies. The paint runs after more than 67 ms and draws
`mainarea`, `torch1` at (4,200) and `torch2` at (588,200) at frames
counter mod 15 and (counter+8) mod 15, and the `blind` sprite at a random
point of a random rectangle of the levels, Cancel and OK, stepping after
more than 63 ms and replaying after 500+rand()/65 ms. `tablol` is loaded
and not drawn. Sounds: `Char.wav` on a hero click, `Level1..3.wav` on a
difficulty click, `Ok.wav` on OK and Cancel.

**Confidence.** High: rectangles, mask values and art fields come from
the selected init, loader and paint bodies and the mask census.

### R2-ENGINE-285

Code table EN `L2.01115` / RU `L2.01116` holds 0, 2, 3, 1. Hero index i
stores code<<6 in application +0x484 (RU +0x49c); the template is
Start_MF (0), Start_MM (0x40), Start_FF (0x80) or Start_FM (0xc0).
Indexes 0..3 are therefore male fighter, female fighter, female mage and
male mage. Activation `L2.01117` selects index 0, takes the difficulty
from control +0x598 (set to 1 by New Game) and, when the name is a default
(`L2.01118`), loads npcnames line 20. A hero click replaces a default name
(npcnames 23..26 or "Unnamed") with line 23, 24, 26 or 25. The name is
typed in the field; OK refuses an empty name. Difficulty has three levels
and changes no other control.

**Confidence.** High for the table, defaults and name lines.

### R2-ENGINE-286

| Child | Rectangle | Art |
|---|---|---|
| Stats 0x457 | (0,0,160,238) | `main\graphics\chrgen\leftup.bmp`, localized with labels in the art |
| Stat sheet 0x458 | (0,238,160,480) | `FullStatsL`, then shared painter `L2.01119` |
| Skills 0x45a | (160,0,480,480) | `fighter\column` or `mag\column` |
| Buttons 0x459 | (480,0,640,238) | `interface\inn\ButtonsArea`, `button{1,2,3}{on,off}` |
| Inventory | (480,238,640,480) | the application inventory panel, reparented |

Stats row i (Body, Agility, Mind, Spirit; main.txt 15..18) has a value
cell at (82,54+32i), plus at (107,54+32i), minus at (132,54+32i), all
20x20, and a label rectangle at (16,57+33i); the pool cell is (46,181)
77x22. Plus art is `plon` hovered and held, `ploff` hovered, `pnloff`
normal, `pdisable` disabled; `pnlon` is loaded and not drawn; minus uses
the `m*` set. Buttons (484,44), (484,91), (484,138), 140x46: Accept
main.txt 238, Reset 239, Back 260; each fires on release over the pressed
button. Activation shows the template values with pool 0. EN Reset
`L2.01120` sets all four values to 25 and the pool to 100. RU Reset
`L2.01121` sets the pool to 0, calls `L2.01122`, which zeroes the
attribute and skill overrides and rebuilds the template hero, then
reloads values and relights the template skill. The detail paint draws
children only while detail +0x104 is set by activation.

**Confidence.** High for layout, art and both Reset bodies. Unknown for
the stat-sheet and inventory content.

**Unknown.** `RollStatsR` and `FullStatsR` are loaded into skill-child
+0x68/+0x6c with no drawing site in the selected bodies.

### R2-ENGINE-287

Cost helper `L2.01123`: F(v) = trunc(0.349·1.15^(v−1) + 0.5), constants
at EN `L2.01124`/`L2.01125`. Plus `L2.01126` requires pool ≥ F(v+1)−F(v)
and v < 45 and takes that cost; minus `L2.01127` requires v > 15 and adds
F(v)−F(v−1). Each change plays `+_-.wav` and passes the four values and
the skill to `L2.01128`, which rebuilds the preview hero. Producer
`L2.00160` receives the values in packet bytes +0xa..+0xd and accepts
them when 140 − ΣF ≥ 0; otherwise all four become 25. Start_MF 40/36/25/17,
Start_FF 37/39/21/25, Start_FM 19/23/30/42 and Start_MM 28/20/41/32 each
give ΣF = 140, equal to 4·F(25) + 100. The pool cell shows the free
points; tooltips give "%+d" of the plus cost and of the minus refund.

**Confidence.** High for the formula, limits, budget and template sums.

### R2-ENGINE-288

Loader `L2.01129` chooses fighter or mage art by hero +0x1b8 bit 1 and
stores five mask bytes at skill-child +0x104..+0x108: fighter 255, 191,
152, 127, 102; mage 127, 102, 255, 152, 191. Hit test `L2.01130` reads
the mask at (x−160, y) and compares +0x104..+0x107 only, so the fifth
skill (bow, astral) cannot be clicked; the paint draws indexes 0..3 only.

| Class | Index 0..3 | Positions | Fifth |
|---|---|---|---|
| Fighter | sword, axe, mace, pike | (248,93), (252,126), (248,182), (244,225) | bow (248,250), column art only |
| Mage | fire, water, air, earth | (360,150), (232,165), (292,98), (300,228) | astral (296,158), column art only |

Art: state 1 `on`, 2 `shine_off`, 3 `shine_on` (bit 0 selected, bit 1
hover). A click selects the skill, plays its `SFX\ChrGen\Skill` sound,
sends skill index+1 and raises `#tips7` once. Activation lights index
(application skill − 1); with no click that skill is the template main
skill, the argmax of skills 1..5, which is 1 for all four templates:
sword or fire.

**Confidence.** High for the tested bytes, art and default skill.

### R2-ENGINE-289

| Event | Text |
|---|---|
| Pre-create activation | `#tips8` in panel 0x467 (232,48)-(640,184) |
| Hero click at stage 0 (mask 80..140) | `#tips9`, stage 1 |
| Difficulty click at stage 1 (mask 20..60) | `#tips10`, stage 2 |
| Detail activation | `#tips5` fighter or `#tips6` mage, panel 0x467 (0,280)-(312,480) |
| First skill click | `#tips7` |

The tips need the tips mode (`L2.00713`). Pre-create cycle `L2.01131`
lights stage 0 heroes, stage 1 levels or stage 2 Cancel/OK in turn: it
starts 500 ms after the pointer leaves that stage's targets and steps
after more than 300 ms. Skill cycle `L2.01132` lights the four skills the
same way until the first click. Tooltips: stats labels main.txt 155..158,
pool 273, value "%s = %d"; skills main.txt 171+i fighter, 176+i mage. The
pre-create paint sets cursor `graphics\cursors\select`; the detail paint
sets `graphics\cursors\default` unless the cursor is default or dice.

**Confidence.** High for the keys, events and cycles. Medium for the
cursor: other bodies can set it between paints.

**Unknown.** Which cursor shows on the detail screen between paints. A
census of the direct callers of the cursor setter `L2.01133`, which the
detail paint `L2.01134` calls, and of their screens would settle it.

### R2-ENGINE-290

Producer `L2.00160` builds `Start_XX` from the flags, loads it through
`R2.0088` and `L2.00219` (−1 keeps a field), applies chosen-skill
`L2.01135(skill, 20)`, runs recompute `L2.00318` and sets HP and mana to
their maxima. `L2.01135` deletes the hand item, zeroes school skills
1..5, sets the chosen skill to 20 and skill 5 to 10, computes
E(s) = trunc((1.1^s − 1)·1000) per school skill (s capped at 149) and
equips Iron Long Sword, Iron Axe, Iron Mace, Iron Pike or Uncommon Wood
Long Bow for a fighter, or for a mage a staff with castSpell Fire_Arrow,
Ice_Missile, Lightning, Diamond_Dust or Drain_Life at 20. Experience is
E(20)+E(10) = 7320.

Class and sex come from two predicates. `L2.01136` returns hero byte
+0x4c & 4 (mage). `L2.01137` returns 1 when hero word +0xe is 0x22 or
0x24 (female). Template load `R2.0088` reads Humans parameter 0x12
(`gender ( is female? )`, 0 or 1) and writes +0xe = gender+0x23 for a
mage (`L2.01138`/`L2.01139`) and gender+0x21 for a fighter
(`L2.01140`/`L2.01141`): MF 0x21, FF 0x22, MM 0x23, FM 0x24. The mage
arm of `L2.01135` calls `L2.01137` at `L2.01142`: the female mage gets
`Wood Staff {castSpell=<spell>:20}` and the male mage
`Uncommon Wood Staff {castSpell=<spell>:20}`. Recompute calls the same
two predicates (`L2.01143`, `L2.01144`) to choose the caps below.

The producer appends ".f5" to the template name when the code byte is 0
(Start_MF) and otherwise the string at `L2.01145` or `L2.01146` plus
code & 0x3f, which is 0 for the generator's codes. Template load writes a
positive parsed number to hero byte +0x4b (`L2.01147`), so only Start_MF
receives +0x4b = 5.

Recompute caps Body/Agility/Mind/Spirit at 52/50/48/46 for a male
fighter, 50/52/46/48 female fighter, 48/46/52/50 male mage and
46/48/50/52 female mage. HPmax = trunc(trunc(Body·k + log1.1(exp/5000+1)·k)
·(1.1^Body/100+1)), k = 2 fighter, 1 mage; mana uses Spirit·2 with k = 1
fighter, 2 mage, only when the mana field is nonzero. Speed is Agility (title Reaction)
when below 12, else Agility/5+12, +10 when +0xe is 0x13 or 0x15 (none of
the four heroes), minus the
load term, minimum 6. Sight word +0xa4 = trunc(((Mind+Agility)/25+4)·256);
+0xbe = Agility/3; resistances Spirit/2; item modifiers `L2.01058` follow.
After restoring the school skills recompute adds the +0xe8 skill
modifiers.
The template values give HP 142, 123, 29, 42, mana –, –, 157, 99, speed
19, 19, 16, 16 and sight byte 6 for MF, FF, FM, MM before modifiers.

**Confidence.** High for the arithmetic, the predicates and the weapon
table. Medium for the listed values at town 1: item modifiers and load are
not applied, and a fighter's mana field is assumed zero. Medium for the
name suffix: the string parse inside `R2.0088` was not traced beyond the
+0x4b write. Unknown for item effects.

**Unknown.** Item modifier totals, including +0xe8 skill modifiers, and
the load term; the item-effect tables or a live save would settle them.

**Amended.** R2-ENGINE-319 reads the RU suffix store (`L2.01270`), which
writes 5 for the name `Start_MF.f5` after the parameter routine stores the
face. The Start_MF hero of the four frozen RU saves holds +0x4b = 32, the
row's face parameter (R2-SESSION-144); which writer or construction path
leaves 32 in the generated hero is Unknown. R2-SESSION-144 gives that hero's values after its
attribute choice and items.

## Cheat and debug commands

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-295 | ROM2 parses chat lines that start with # in the client's in-process server (EN L2.00987, RU L2.00988, equal): 24 distinct texts in 27 string instances, case-sensitive prefix match in fixed order; no chat relay. | High | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |
| R2-ENGINE-296 | The ROM2 server offers cheats only when created in mode 2, which only the campaign new and load paths pass; other modes offer #kick and #locate to the host connection and the two latency commands to others. | High / Medium | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |
| R2-ENGINE-297 | In ROM2 campaign mode a locked Player's chat line whose first 7 bytes are ##Cowar sets its flag byte +0xa78 to 0xff and broadcasts reply 5; other lines are dropped and the unlock line runs no command. | High | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |
| R2-ENGINE-298 | ROM2 command replies are message 0x92 to every client with code 5, 6 or 7 and a Player ID; the client shows that Player's name with main.txt lines 221-222, 223-224 or 225-226 for 5 seconds. | High | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |
| R2-ENGINE-299 | ROM2 #create [N] Gold adds N to Player +0x3c; #create [N] name puts the named item with quantity N in the hero's inventory; both refuse with reply 6 when hero kind byte +0x13c is nonzero. | High / Medium | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |
| R2-ENGINE-300 | ROM2 #modify self or army +god sets six words +0x102.. and six bytes +0x10e.. to 100 and re-derives; +spell N and +spells fill the hero's book; +knowledge only sends the Player its 2,560-byte block. | High / Unknown | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |
| R2-ENGINE-301 | ROM2 #summon [N] [hero] name builds N creatures or humans by name near the hero in a new group owned by the hero's owner; #pickup all gives every world sack's money and items to the hero. | High / Medium | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |
| R2-ENGINE-302 | ROM2 #killall and #kill all set word +0x94 to -50 on every unit of each Player whose diplomacy byte toward the sender has bit 0; #kill cheaters also clears their flags; #kill name hits one Player. | High / Medium | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |
| R2-ENGINE-303 | ROM2 #show map and #hide map send message 0xaa to the Player, setting client global L2.00989 and on show marking its grid; #victory posts the normal completion message 0x430; #event N sends 0xb6. | High / Unknown | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |
| R2-ENGINE-304 | ROM2 Alt+B..Y in the mission view send message 0x46 sub 0x80; the server acts only for an unlocked Player: D turn tracing, H help, I last turn, Q safe mode, T script tracing, U experience. | High / Medium | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |
| R2-ENGINE-305 | ROM2 allods2.exe has no cheat switch: the only 0xff store to the flag is the chat unlock, and no command-line test in the substring census (EN 32, RU 31 sites) feeds a flag or mode store. | High | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |
| R2-ENGINE-306 | RU a2server.exe holds the client's 24 command texts plus #ready, the same password object and the flag at +0xa98; a Player identity pair at +0x10 reaches the cheat path in any mode. | High / Unknown | ● active (branch candidate) | [EXP-2036](../experiments/EXP-2036-rom2-cheats/) |

### R2-ENGINE-295

Chat entry: Enter in the game view (key switch `L2.00990`) opens child
+0x138 (`L2.00991`); its result 0x445 is read at `L2.00992`. A `=` prefix
sends the rest as type 3, a `-` prefix sends type 1 or 2, other text type 0;
`L2.00993` builds message 0x91 (text at +0xf) and sends it to the remote
server or, without one, to the local connection layer `L2.00994`, which
passes it to the server object held in global `L2.00805` when that global
is nonzero. Server switch `L2.00995` routes 0x91 to `L2.00996`, which resolves the sender from +5,
logs the line and calls the parser `L2.00987` when the first byte is `#`.

The 24 distinct texts, in test order: `#kick `, `#locate `, `#set latency `,
`#show latency`, `#create ` (argument `Gold`), `#modify ` (target `self`
or `army`; modifier `+god`, `+spell `, `+spells`, `+knowledge`),
`#summon ` (option `hero`), `#killall`, `#kill all`, `#kill cheaters`,
`#kill `, `#pickup all`, `#show map`, `#hide map`, `#victory`, `#event `.
Each is found by `L2.00997` at offset 0, removed (`L2.00998`) and the rest
left-trimmed (`L2.00999`). `Gold` is compared case-insensitively with the
whole argument (`L2.01000`); Player names exactly (`L2.01001`). Count
parser `L2.01002`: a positive `atoi` token before a space is the count,
otherwise 1. Each image stores 27 string instances of these texts (`self`,
`army` and `+spell ` twice each); every instance has exactly one code
reference, inside the parser, in EN and RU. The 12 texts tested after the
unlock gate are the cheat commands (`#create` through `#event `); the
four before it are the host and latency commands.

**Confidence.** High: the EN and RU parser bodies are equal at level 0
(1,293 instructions) and the string census covers every non-code section.

### R2-ENGINE-296

Server init `L2.00221(mode)` stores +0x74 = (mode < 2). Creation
`L2.01003` passes application mode +0x5d8 at `L2.01004` and `L2.01005`,
each right after that mode is set to 2 (the new-campaign path reached by
application message 0x425, and a load path); 0 at `L2.01006` and
`L2.01007`; the mode at `L2.01008` after `L2.01009` stores 1. Mode 3
(`L2.01010`, `L2.01011`) creates with 0. The only +0x74 stores in the
server-code span (EN `L2.01012..L2.01013`) are constructor `L2.01014` (0),
init `L2.01015`, and `L2.01016` on another object.

RU agrees through its own census: application mode +0x63c is stored 2 at
`L2.01017` and `L2.01018`, each followed by creation `L2.01019` with that
mode; 0 at `L2.01020`, `L2.01021`, `L2.01022`, `L2.01023`; 3 at `L2.01024`,
`L2.01025`; 1 at `L2.01026`. RU init `R2.0010` and creation `L2.01019`
equal the EN bodies (levels 0 and 1); +0x74 stores in RU
`L2.01027..L2.01028` are `L2.01029` (0), `L2.01030` and `L2.01031`.

With +0x74 set the parser tests the sender's connection +0x29c
(`L2.01032`), which the one setter call `L2.01033` sets to 1 on the
server's own local connection. That connection may use `#kick name`
(message 0x93 through `L2.01034` when the target's connection is not the
host's) and `#locate name` (a `"%s (%d,%d)"` line when the target's +0x2c is
0). Any other connection may use `#set latency N` (N = 0 or 50..10000, else
reply 6; stored through `L2.01035`, `L2.01036`) and `#show latency`. The
parser then returns.

**Confidence.** High for the code paths in EN and RU. Medium that mode 2
means the single-player campaign: it rests on the two creation sites and
the campaign reading of application mode 2 in R2-ENGINE-271.

### R2-ENGINE-297

Flag getter `L2.01037` returns byte +0xa78 > 50; setter `L2.01038`. The
constructor `L2.01039` stores 0; the unlock stores 0xff; `#kill cheaters`
stores 0. No other EN or RU instruction names +0xa78.

With +0x74 clear and the flag not above 50, the line goes to check
`L2.01040` on the object at server +8 (constructor `L2.01041`, 0x68 bytes
zeroed). Entry 0 is `39 20 59 7c 02 11 5d 40 46 00`; entry 1 is
`##Coward  ` minus entry 0, byte 7 then set to 0. The check accepts when
entry1[j] + entry0[j] equals input[j] for every j before the first zero of
entry 1 and j > 2: the first 7 input bytes must be `##Cowar`. On a match
the parser logs the unlock (printed only in application mode 3 by
`L2.01042`), sets 0xff and sends reply 5 (R2-ENGINE-298). On no match it
sends nothing. Both cases return.

With the flag above 50 the parser goes straight to the command tests.
`#create`, `#summon`, the three kill forms, `#pickup all`, `#show map`,
`#hide map` and `#victory` re-test the flag; that failure branch cannot be
reached after the gate.

The check returns the first matching entry index and tests entry 1
first. The parser accepts only a result of exactly 1 (`cmp eax, 1` at EN
`L2.01043`, RU `L2.01044`, server `L2.01045`). A line that fails entry 1
can match only at i >= 10, which returns a value other than 1 and is
rejected; entries 2..9 are zero, and entry 10 is the check's own result
field, which ends before j > 2. The loop also reads 10-byte entries past
the 0x68-byte object for i >= 11; no such match can be accepted.

**Confidence.** High that only a line whose first 7 bytes are `##Cowar`
unlocks, in EN, RU and `a2server.exe` (byte-identical constructor and
check, the same result test).

### R2-ENGINE-298

`L2.01046(code, id)` calls `L2.01047(0x92, code, id, 0)`; target 0 sends to
every client. Client switch `L2.01048` (code - 1, bound 0x7f) sends code 5
to `L2.01049`, 6 to `L2.01050` and 7 to `L2.01051`. Each formats the
Player's name between two `main.txt` lines (221/222, 223/224, 225/226) and
shows it for 0x1388 ms. Code 5 is the unlock, 6 a refusal, 7 a success.
EN and RU hold all six lines.

**Confidence.** High.

### R2-ENGINE-299

`#create` requires the hero's kind byte +0x13c to be 0, else reply 6.
`Gold`: `L2.01052(N)` adds N to Player +0x3c and sends message 0x67 with
the new value; reply 7. Other text: `L2.01053` builds an item by name;
validity `L2.01054` fails: the item is deleted, reply 6. Else item +0x42 =
N, the item goes into hero +0x7c through `L2.01055`, unit update
`L2.01056`, reply 7.

**Confidence.** High for the stores and messages. Medium for the names
money and inventory, read from field use and R2-SESSION-020.

### R2-ENGINE-300

`#modify` takes `self` (the hero) or `army` (every unit in Player +0x24);
other text returns. `+god` calls `L2.01057` on each target: words
+0x102..+0x10c and bytes +0x10e..+0x113 = 100, all inside the 0x40-byte
block at +0xd4, then vtable +0x58 (derive `L2.00318`, which folds the block
through `L2.01058` and clamps resistances as R2-ENGINE-029). `+spell N`
(self, book +0x140 present, 0 < N < spell count) builds spell N
(`L2.01059`) and stores it at index N (`L2.01060`), replacing what was
there. `+spells` does the same for 1..29. `+knowledge` calls
`L2.01061(0, Player)`, which sends message 0xba with the Player's
2,560-byte +0x44 block; no store. All reply 7 except `+knowledge`.

**Confidence.** High for the stores and calls. Unknown for which unit values
the six bytes reach.

**Unknown.** `L2.01058`; reading it would settle the bytes' effect.

### R2-ENGINE-301

`#summon` requires Player +0x38 (hero). After the count, a `hero` token
sets the hero option. `L2.01062` tries a creature by name (`L2.01063`,
0x208 bytes), then a human (`L2.00214`, 0x254 bytes); a name neither knows
is deleted. A built unit is placed within 3 cells of the hero
(`L2.00312`) when server +0x94 is set, owned by the hero's owner, put in a
new group and added to the World. No reply. `#pickup all` (hero required)
removes every sack listed at server +0x7c from the world and calls
`L2.01064`: sack money to Player +0x3c through `L2.01052`, items to the
inventory; reply 7.

**Confidence.** High for the calls. Medium for the creature and human
naming, read from object sizes and constructors.

### R2-ENGINE-302

`L2.01065` stores -50 in word +0x94 of every unit in the Player's +0x24
list; +0x94 is the current value bounded by +0x96 (R2-ENGINE-086).
`#killall`/`#kill all` apply it to each Player P whose byte at World
`L2.01066` + 0xa8c4 + P.id*0x46 + sender.id has bit 0 set; reply 7.
`#kill cheaters` applies it to every other Player whose flag is above 50,
after storing flag 0. `#kill name` applies it to the Player with that
exact name; reply 7 with that Player's ID.

**Confidence.** High for the stores and selection. Medium that bit 0 marks
hostility and that -50 kills, both inferred.

### R2-ENGINE-303

`#show map` and `#hide map` send message 0xaa to the Player with argument
1 or 0; reply 7. Client arm `L2.01067`: 0 stores global `L2.00989` = 0; 1
stores 1 and ORs 0xc000 into every word of the grid at client +0x80. Its
readers skip the periodic `L2.01068` (`L2.01069`) and force visibility
level 7 (`L2.01070`, `L2.01071`). `#victory` sends argument 2: the client
posts application message 0x430, which normal completion message 0xb5
also posts. Arm `L2.01072` posts 0x41d when Scenario slot 0x300 is at
least 120, else opens the dialog with `main.txt` line 140. No server
completion state is written. `#event N` sends message 0xb6 with N to the
Player through `L2.00458`, read as application message 0x433
(R2-ENGINE-047).

**Confidence.** High for the messages and stores. Unknown for the grid's
meaning.

**Unknown.** What client +0x80 holds. Searched: the 0xaa arm only; the
grid's other readers and writers were not censused. A census of client
+0x80 users and the visibility test at `L2.01070` would settle it.

### R2-ENGINE-304

Key handler `L2.01073`: Alt (lParam bit 0x2000), VK 0x42..0x59 and
application +0x404 bit 0 (`L2.01074` stores 1 when the mission view is
built) call `L2.01075`, which sends message 0x46 sub 0x80 with VK - 0x41.
Server switch `L2.01076` calls World method `L2.01077(Player, index)`,
which returns unless flag > 50, then switches on index - 3:

| Key | Index | Store or output |
|---|---|---|
| D | 3 | server +0x170 word +0 toggled |
| H | 7 | six help lines |
| I | 8 | `L2.01078` |
| Q | 0x10 | World +0xbbe8 toggled |
| T | 0x13 | server +0x170 word +4 toggled |
| U | 0x14 | `L2.01079` |

Output goes to chat through `L2.00513`.

**Confidence.** High for the route and stores. Medium for the names, read
from the help lines.

### R2-ENGINE-305

Census: every call of the substring search (EN `L2.01080`, RU `L2.01081`)
whose 8-instruction window reads application +0x70. Switches found:
`-saveonserver`, `-internetserver`, `-latency`, `-timeout`, `-nomusic`,
`-trace`, `-safevideo`, `-cfg"`, `-window`, `-startserver`, `-cfg`, `.asl`,
`-800`, `-1024`, `-640`, `-protocol`, `-map"`, `-protocol0..4`, `-female`,
`-mage`, `-name`, `-waitforever`, `-ip"`. `-trace` (`L2.01082`) stores
`L2.01083` = 1, which enables diagnostic logging. `-female`, `-mage` and
`-name` are read in the new-campaign path `L2.01084`. The flag stores, the
+0x74 init and the +0x5d8 stores take no value from these tests.

The claim that no switch enables cheats rests on the flag-writer census,
not on the switch census alone: every EN and RU instruction naming +0xa78
is the constructor, the getter or the setter (R2-ENGINE-297), and the
setter's two calls push 0xff (the chat unlock) and 0 (`#kill cheaters`).

**Confidence.** High. The switch census alone is bounded: a switch read
through a routine other than the substring search is outside it.

### R2-ENGINE-306

Parser `L2.01085` (1,849 instructions) references the 24 client texts (28
string instances with `#ready`) and
`#ready` (`L2.01086`), each once. Password constructor `L2.01087` and check
`L2.01088` equal the client's byte for byte. The flag is Player +0xa98.
After the +0x74 block, a Player whose DWORDs at +0x10 are `0xf6d04773` and
4 continues to the unlock and commands instead of returning; the same
pair suppresses replies 5 and 6. The pair is compared at 22 sites.
`#ready` tests global `L2.01089` == 2 and Player +0xa6c.

**Confidence.** High for the inventory and gate code. Unknown for which
Player carries the pair, the modes the server reaches and `#ready`'s effect.

**Unknown.** Searched: the server's compares of the pair (22 sites) and the
parser body. Not searched: writers of Player +0x10 in the server, the
server's creation calls and `#ready`'s callees. Reading the server writers
of Player +0x10 would identify the pair's holder; a census of the server's
creation calls with their modes would settle the modes; reading the
`#ready` arm would settle its effect.

## Tip popups

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-307 | ROM2 has 17 tip formatter calls per locale: 16 push a constant tip number 1..12 and one passes the 0x45b wParam; 12 direct panel-constructor calls create tip popups and 4 sites retext them. | High | ● active (branch candidate) | [EXP-2037](../experiments/EXP-2037-rom2-tips/) |
| R2-ENGINE-308 | ROM2 room tips open at every room entry while TipsMode is set: squares #tips1, #tips12 (Kaarg), #tips11 (druid); inns #tips2 and shops #tips3 also need mode 2; #tips4 then replaces #tips3 once per shop entry. | High / Medium | ● active (branch candidate) | [EXP-2037](../experiments/EXP-2037-rom2-tips/) |
| R2-ENGINE-309 | ROM2 message 0x45b shows #tips<wParam> of the current text pool in mission-view child 0x10 (RU 0x11) at (10,20,370,188); the only constant 0x45b in either client is a dialogue's post of its nonzero tips= value, which no shipped text has. | High | ● active (branch candidate) | [EXP-2037](../experiments/EXP-2037-rom2-tips/) |
| R2-ENGINE-310 | ROM2 tip popups have no show-once latch: Close (0x45a) and the room leave delete the panel, and an entry that does not build one deletes a leftover; the shop and generator latches reset at each entry. | High | ● active (branch candidate) | [EXP-2037](../experiments/EXP-2037-rom2-tips/) |
| R2-ENGINE-311 | ROM2 TipsMode is loaded from HKLM once at startup and stored at exit and after a cutscene through a key opened with KEY_READ; clearing it closes no open tip, stops later builds and the pre-create cycle, not the detail cycle. | High / Medium | ● active (branch candidate) | [EXP-2037](../experiments/EXP-2037-rom2-tips/) |
| R2-ENGINE-312 | ROM2 tip panels have a fixed rectangle per creation site and never size to their text: square (328,0)-(640,200), inn (160,0)-(472,200), shop (164,162)-(476,298), mission view (10,20)-(370,188); no gate tip. | High / Medium | ● active (branch candidate) | [EXP-2037](../experiments/EXP-2037-rom2-tips/) |
| R2-ENGINE-313 | Every ROM2 tip popup is one panel class: an lm.256 frame tiled to the panel, font2 gold text at (20,24), a Close button and a checked Show-tips box; no room or mission tip lights another control. | High | ● active (branch candidate) | [EXP-2037](../experiments/EXP-2037-rom2-tips/) |
| R2-ENGINE-314 | ROM2 tip text splits at CR LF, wraps greedily at spaces to the panel width minus 48 in font2, indents each paragraph 6 pixels, justifies non-final lines and clips past (height-60)/12 lines; no shipped tip clips. | High / Medium | ● active (branch candidate) | [EXP-2037](../experiments/EXP-2037-rom2-tips/) |

### R2-ENGINE-307

The tip-section formatter EN `L2.00896` / RU `L2.01190` formats
`#tips%d`, finds its first occurrence in the text pool `L2.01191` and
returns the bytes from the key end plus two to the next `#`, or an empty
string when the key is absent. Panel constructor EN `L2.00633` / RU
`L2.00897` takes (ID, left, top, right, bottom, text). Retext EN
`L2.01192` / RU `L2.01193` replaces the text child's lines; wrapper EN
`L2.01194` / RU `L2.01195` calls it.

A byte scan of every `.text` offset for direct calls finds, per locale,
17 formatter calls, 12 constructor calls, 3 calls to the retext (one is the
wrapper) and 2 calls to the wrapper.

| Tip | EN function | RU function | Action | Panel ID | Rectangle in parent | Panel field |
|---|---|---|---|---|---|---|
| 1 | `L2.00710` | `L2.01196` | create | 0x467 | (328,0,640,200) | square +0x200 |
| 12 | `L2.00631` | `L2.00632` | create | 0x467 | (328,0,640,200) | square +0x200 |
| 11 | `L2.01197` | `L2.01198` | create | 0x467 | (328,0,640,200) | square +0x200 |
| 2 | `L2.00552`, `L2.00684`, `L2.00613` | `L2.00553`, `L2.01199`, `L2.00614` | create | 0x467 | (0,0,312,200) | inn page +0x80 |
| 3 | `L2.00767`, `L2.01200`, `L2.00954` | `L2.01201`, `L2.01202`, `L2.01203` | create | 0x3f3 | (0,162,312,298) | shop page +0x88 |
| 4 | `L2.01204` | `L2.01205` | retext | | | shop page +0x88 |
| wParam | arm `L2.01206` | arm `L2.01207` | create | 0x10 (RU 0x11) | (10,20,370,188) | none |
| 5, 6 | `L2.01208` | `L2.01209` | create | 0x467 | (0,280,312,480) | detail +0x80 |
| 7 | `L2.01210` | `L2.01211` | retext | | | detail +0x80 |
| 8 | `L2.01117` | `L2.01212` | create | 0x467 | (232,48,640,184) | pre-create +0x1f8 |
| 9, 10 | `L2.01213` | `L2.01214` | retext | | | pre-create +0x1f8 |

R2-ENGINE-289 states the generator events for tips 5 to 10. Room table
slots (enter +0x80, leave +0x84, message +0x48) bind the square, inn and
shop bodies: EN `L2.00562`, `L2.00563`, `L2.00564`, `L2.00739`,
`L2.00741`, `L2.00740`, `L2.01215`, `L2.01216`, `L2.01217`.

**Confidence.** High. The call census covers every byte offset of `.text`
in both clients, every site decodes to the listed constant or wParam, and
the 76 named EN bodies equal their RU partners in mnemonic sequence except
the wrap body (R2-ENGINE-314) and the application initializer, which is
not a tip site.

### R2-ENGINE-308

| Tip | Surface | Enter test | Shown |
|---|---|---|---|
| 1 | town 1 square | TipsMode nonzero | every entry |
| 12 | Kaarg (town 2) square | TipsMode nonzero | every entry |
| 11 | druid (town 3) square | TipsMode nonzero | every entry |
| 2 | generic, druid and Kaarg inn | TipsMode nonzero and application mode word +0x5d8 (RU +0x63c) equal to 2 | every entry |
| 3 | generic, druid and Kaarg shop | the same two tests | every entry |
| 4 | the shop popup | no TipsMode test | once per shop entry |

Town IDs follow R2-ENGINE-231. The mode word is the one the save driver
tests (R2-ENGINE-271).

Tip 4: the shop message slot EN `L2.01218` / RU `L2.01219` decodes message
0x402 to EN `L2.01220`; when application +0x404 has bit 8 clear it calls
EN `L2.01204` / RU `L2.01205`. That body needs the popup field +0x88
nonzero, a nonzero count (`L2.01221`, the dword at +8) of the item list at
shop-inventory child (+0x70, ID 0x3eb, R2-ENGINE-249) +0x84, and latch
+0x8c zero. It then sets +0x8c to 1, formats `#tips4` and retexts the open
popup. The application idle body EN `L2.01222` / RU `L2.01223` sends 0x402
to the room stack while application +0xbc is nonzero. The child's list
pointer is set by its table slot +0x90 (EN `L2.01224`).

No tip site reads a script, pickup, kill or timer state. Missions raise
tips only through message 0x45b (R2-ENGINE-309).

**Confidence.** High for the tests, fields and message routes in both
clients. Medium that 0x402 reaches the shop as a room-stack child; the
stack offers a message to its children in list order (R2-ENGINE-275).
Medium that the item list becomes nonempty when a shelf is opened; the
callers of slot +0x90 were not read.

**Unknown.** Which shop actions fill the shop-inventory list and what
application +0x404 bit 8 means. Reading the callers of slot +0x90 and the
writers of that bit would settle them.

### R2-ENGINE-309

Arm 0x45b of the application dispatcher (EN `L2.00262` byte table
`L2.00794`, dword table `L2.00795`; RU `R2.0036`, `L2.01225`,
`L2.01226`) runs when TipsMode is nonzero. It formats `#tips<wParam>` from
the current pool, removes and deletes mission-view child 0x10 (RU 0x11) if
present, constructs a panel with that ID at (10,20,370,188) and adds it to
application +0xd0. That field holds the window built by EN `L2.01227` with
(0,0,width-160,height) at EN `L2.01228` / RU `L2.01229`.

The only `.text` instruction per locale whose immediate is 0x45b is the
push in dialogue message body EN `R2.0054` / RU `R2.0053`. On exhaustion
it posts 0x45b with dialogue +0x80 as wParam when that field is nonzero,
then closes with 0x445 (R2-ENGINE-051). Constructor EN `R2.0040` zeroes
+0x80; parser EN `R2.0048` / RU `R2.0047` reads five characters after
`tips=` with `%d` into it.

The pool holds `town.txt` in towns and the generator, `globalmap.txt` on
the map, and the mission or quest text in a mission. A census of every
`.res` entry, `.alm`, `.dll` and `.exe` file in both roots finds `#tips`
sections only in `town.txt` and `tips=` only as a string in the executables
(R2-ASSET-080). No shipped dialogue posts 0x45b through that push, and a mission pool has no
`#tips` section, so a `tips=` added to mission text would open an empty
panel.

**Confidence.** High. The arm, the poster census and the resource census
cover both clients and both roots. The poster census is bounded to a
constant 0x45b; a message number computed at run time is not excluded.

### R2-ENGINE-310

| Family | Panel field | Close (0x45a) deletes in | Leave deletes in | Entry without a build deletes in | Latch and reset |
|---|---|---|---|---|---|
| squares | +0x200 | square message `L2.00726` | `L2.00784` | each square enter | none |
| inns | +0x80 | inn message `L2.01230` | `L2.01231`, `L2.01232`, `L2.01233` | each inn enter | none |
| shops | +0x88 | shop message `L2.01218` | `L2.00966` | each shop enter | +0x8c, set by tip 4, cleared by every shop enter |
| mission | child 0x10 | arm 0x45a `L2.00911` | | the next 0x45b | none |
| detail | +0x80 | detail message `L2.01234` | `L2.01235` | activation | +0x100, cleared by activation |
| pre-create | +0x1f8 | pre-create message `L2.01236` | `L2.01237` | activation | stage +0x218, cleared by activation |

Addresses are EN; `room-tables.tsv`, `field-writes.tsv` and
`message-switches.tsv` give the RU partners. Close is the panel button's
0x45a; the application arm deletes mission child 0x10 if present and
otherwise offers 0x45a to the room stack (R2-ENGINE-275). An entry builds
no panel when TipsMode is zero, or for inns and shops when the mode word is
not 2; it then deletes a panel left in the field. Square, inn and shop
popups therefore show at each entry while TipsMode is set, and closing one
holds only until the room is entered again. Every +0x8c store in the
bodies of the three EN shop tables is the enter's zero; the only other
store is tip 4's one. R2-SESSION-140 covers SAVE and LOAD.

**Confidence.** High. Every field store in the named bodies of both
clients is listed, and the table slots bind those bodies to the rooms.

### R2-ENGINE-311

Load EN `L2.01238` / RU `L2.01239` and store EN `L2.01240` / RU
`L2.01241` open `HKEY_LOCAL_MACHINE` key EN `SOFTWARE\Rage of Mages 2` /
RU `SOFTWARE\1C\Allods 2` with access mask 0x20019 (KEY_READ). The load
reads value `TipsMode` into the global through `RegQueryValueExA` and runs
once, from the application initializer (EN call `L2.01242`). The store
ignores the open result and writes `TipsMode` with `RegSetValueExA` from
EN `L2.00893`. It runs from the application exit (EN `L2.01243`) and from
the cutscene player EN `L2.00875` after it marks a video as seen (EN
`L2.01244`). R2-ENGINE-273 gives the default 1 and the four writers; no
writer deletes a panel.

| Tip use | Tests TipsMode |
|---|---|
| every panel creation (12 sites) | yes, before creating |
| retexts 7, 9 and 10 | yes |
| retext 4 | no |
| pre-create highlight cycle (paint `L2.01114`) | yes, each paint |
| detail highlight cycle `L2.01132` | no; it runs while the detail panel exists and the skill latch is clear |

Clearing TipsMode through the panel checkbox or the options dialog leaves
an open popup until Close, leave or the next entry. A shop popup can
still change to tip 4. The pre-create cycle stops at the next paint; the
detail cycle continues until the first skill click or deactivation.

**Confidence.** High for the key, mask, call sites and gate table in both
clients. Medium that a change persists only where the store's write is
admitted: on Windows NT `RegSetValueExA` needs KEY_SET_VALUE on the handle,
which KEY_READ does not grant, so the stored value would stay as installed;
this was not observed.

**Unknown.** The registry contents of the preserved installs and the
store's result on the owner's system. Reading the key and observing one
exit would settle them.

### R2-ENGINE-312

| Surface | Panel in parent | Parent and its page rectangle | Page rectangle of the panel |
|---|---|---|---|
| town squares | (328,0,640,200) | square view (R2-ENGINE-274) | (328,0)-(640,200) |
| inns | (0,0,312,200) | inn center +0x7c, ID 0x450, (160,0,480,480) (R2-ENGINE-246) | (160,0)-(472,200) |
| shops | (0,162,312,298) | main shop art +0x74, ID 0x3ed, (164,0,480,303) (R2-ENGINE-249) | (164,162)-(476,298) |
| mission view | (10,20,370,188) | window at application +0xd0, (0,0,width-160,height) | (10,20)-(370,188) |
| generator detail, pre-create | R2-ENGINE-289 | | (0,280)-(312,480), (232,48)-(640,184) |

All three inn builders (EN `L2.00744`, `L2.01245`, `L2.01246`) store the
center at +0x7c, and all three shop builders (EN `L2.00762`, `L2.01247`,
`L2.01248`) store the main art at +0x74. Each rectangle is a constant
pushed at its site; no site measures the text. Child rectangles are
parent-relative (R2-ENGINE-274). None of the 17 formatter calls or 12
constructor calls lies in a gate body (R2-ENGINE-307), so the gate shows
no tip.

**Confidence.** High for the constants, parents and the absence at the
gate in both clients. Medium for page positions: the page origin and the
chain to the screen are inferred from construction paths
(R2-ENGINE-238, R2-ENGINE-274).

### R2-ENGINE-313

Panel init EN `L2.00898` / RU `L2.00899` builds the text child, the Close
button and the checkbox (R2-ENGINE-274). Text and button use font2
(global `L2.01249`, created from `graphics\font2\font2` with spacing 2 by
EN `R2.0064`) and colour table `L2.01250`, entry i = (0xb9,0x9f,0x49)·i/15
from `L2.01251`; hover table `L2.01252` is (0x96,0x5a,0)·i/15.

Overlay EN `L2.00903` / RU `L2.01253` uses `lm.256` pieces (R2-ASSET-081)
on the panel less 8 pixels at the right and bottom; with L, T, R'=R-8,
B'=B-8, nx=(R'-L-64)/48 and ny=(B'-T-64)/32:

| Pass | Pieces and positions |
|---|---|
| shadow, slot +0x1c level 6 | 0xc (R'-24,T+8), 0xf (L+8,B'-24), 0x11 (R'-24,B'-24), 0x10 at (L+40+48i,B'-24), 0xe at (R'-24,T+40+32j) |
| corners | 0xa (L,T), 0xc (R'-32,T), 0xf (L,B'-32), 0x11 (R'-32,B'-32) |
| edges | 0xb top and 0x10 bottom at x=L+32+48i; 0xd left and 0xe right at y=T+32+32j |
| fill | 9 at (L+32+48i, T+32+32j) |

Every panel size in R2-ENGINE-312 gives an exact tiling: 312 and 360 and
408 wide, 200, 168 and 136 high.

Button paint EN `L2.01254`: label centred at the rectangle centre plus one
pixel right, shadow offset 2, colour table `L2.01250` (`L2.01252` while
+0x68 is set); bevel (0x29,0x45,0x3f) top and left, (7,0xc,9) bottom and
right. Checkbox paint EN `L2.01255` (slot +0x2c of the checkbox class
`L2.00902`, table `L2.01256`): `radiob.256` frame 5 when set, 4 when
clear, at (left+1,top) with its shadow at (left+5,top+4) level 4; label at
(left+16+6, top+3).

The tip sites, the 0x45b arm and the room message bodies set no highlight
on another control. The generator's two cycles are R2-ENGINE-289.

**Confidence.** High for the bodies in both clients. Pixel colours are
16-bit packings of these values; their exact device colours were not
observed.

### R2-ENGINE-314

Layout EN `L2.01257` / RU `L2.01258` splits the text at CR LF
(`L2.01259`): each paragraph keeps its CR, and left trimming of the rest
removes blank lines. Wrap EN `L2.01260` / RU `L2.01261` trims a paragraph
and, while it does not fit, extends a candidate to the next space,
including that space, until the candidate's width reaches the text width;
it cuts at the last space before that. A paragraph's last line gains a CR.
A first word at least as wide as the line is kept whole; when that word
is exactly as wide, EN (`jle` at `L2.01262`) cuts nothing and RU (`jl` at
`L2.01263`) keeps the word.

Width EN `L2.01264` adds advance+2 per byte, plus 5 for a space; bytes
below 0x20 add nothing, and `~` adds nothing unless doubled. The byte
mapper EN `R2.0060` applies R2-ENGINE-052's selector remap.

Paint EN `L2.01265` / RU `L2.01266` draws at most (text height)/12 lines:
12 is the font height 10 plus 2. A paragraph's first line is indented by
`font2.dat` entry 0x20 (6 pixels). A line that does not end its paragraph,
or a paragraph's first line whose width with one more space exceeds the
text width, is justified (`L2.01267`): words are spread with an equal
floating gap and their positions truncated. Text has a shadow of
(8,8,8) one pixel down and right; `~` followed by another byte underlines
that byte in table entry 15. Click scrolling (text +0x90) is zero for tip
text, so extra lines are clipped.

Text width is panel width minus 48; visible lines are (panel height
minus 60)/12: 264 and 11 for 312x200, 264 and 6 for 312x136, 360 and 6
for 408x136, 312 and 9 for 360x168.

| Key | EN lines | RU lines | Panel |
|---|---|---|---|
| #tips1, #tips11, #tips12 | 8, 6, 7 | 8, 6, 8 | 312x200 |
| #tips2 | 10 | 9 | 312x200 |
| #tips3, #tips4 | 4, 6 | 4, 5 | 312x136 |
| #tips5, #tips6, #tips7 | 8, 9, 7 | 11, 11, 9 | 312x200 |
| #tips8, #tips9, #tips10 | 2, 2, 1 | 2, 2, 2 | 408x136 |

**Confidence.** High for the rules read from the bodies of both clients.
Medium for the line counts: they come from re-implementing these rules,
and the RU counts also assume CharToOemA maps code page 1251 to 866
(R2-ENGINE-052). The EN tip sections hold no byte above 0x7f.

**Unknown.** Live pixels. An authorized capture would settle them. The
CharToOemA code page behind the RU line counts: the conversion result of
one RU section on the owner's system, or a capture, would settle it.

## Hero record writers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-319 | RU ROM2 template load R2.0089 and its parameter routine L2.00220 copy the Humans row into the hero record by a sequential cursor; the producer L2.00162 then writes attributes, name, key, owner and +0x14c. | High / Medium | ✔ promoted (branch candidate) | [EXP-2038](../experiments/EXP-2038-rom2-hero-fields/) |
| R2-ENGINE-320 | RU ROM2 keeps school skills at hero +0xa8+2i, their base at +0x116+2i, the main skill at +0x23c, per-skill experience E(s) at +0x23c+4i (i=1..5), their sum at +0x130 and trunc(0.01 x sum) at +0x1c. | High | ✔ promoted (branch candidate) | [EXP-2038](../experiments/EXP-2038-rom2-hero-fields/) |
| R2-ENGINE-321 | RU ROM2 recompute L2.00052 derives speed, HP and mana maxima, sight, defence and resistances, then adds the item modifier totals kept at hero +0xd8..+0x113 through L2.00102. | High / Medium | ✔ promoted (branch candidate) | [EXP-2038](../experiments/EXP-2038-rom2-hero-fields/) |
| R2-ENGINE-322 | RU ROM2 equip puts a Weapon at hero +0x74 and its modifiers into +0xd4 block fields, an Armor at slot +0x208+4·(armor +0x58), and a refused item into the bag at +0x7c; the spellbook +0x140 exists only with ManaMax > 0. | High / Medium | ✔ promoted (branch candidate) | [EXP-2038](../experiments/EXP-2038-rom2-hero-fields/) |
| R2-ENGINE-323 | The RU ROM2 Player body stores +0x3c and +0xa48 XORed with 0x5c073f4d (L2.01271); Player and Unit store their own address, which load registers in one saved-address map and uses to resolve Player +0x38 and Unit +0x14. | High / Medium | ✔ promoted (branch candidate) | [EXP-2038](../experiments/EXP-2038-rom2-hero-fields/) |
| R2-ENGINE-324 | After its members the RU ROM2 Group body writes u32 +0x1c, +0x40 and +0x44 and ends; the Player body then writes 36 raw bytes of [+0x34] through R2.0175 and ends. Load resolves +0x40 and +0x44 as saved addresses. | High | ✔ promoted (branch candidate) | [EXP-2038](../experiments/EXP-2038-rom2-hero-fields/) |

### R2-ENGINE-319

RU producer `L2.00162` calls template constructor `L2.00130` (at
`L2.01272`). It calls the Humanoid constructor `L2.01157`, which runs the
Unit constructor `L2.01273` and the Humanoid initializer `L2.01274`, sets
vtable `L2.00218` and calls template load `R2.0089`. Template load writes
+0xe = 0, the Humans row index to +0xc (`L2.01275`) and calls the parameter
routine `L2.00220` (`L2.01276`).

The parameter routine sets a cursor {row, index 0} (`L2.01277`). Its
readers `L2.01278` (word), `L2.01279` (byte), `L2.01280` (dword) and
`L2.01281` (word) store a value only when it is not −1 and advance the
index in every case, so the n-th read takes parameter n. Read in order,
under the R2-ASSET-076 titles:

| Parameter | Hero field |
|---|---|
| Body, Reaction, Mind, Spirit | +0x84, +0x86, +0x88, +0x8a |
| HealthMax, ManaMax | +0x96 (copied to +0x94), +0x9c (copied to +0x9a) |
| Speed, RotationSpeed, ScanRange | +0x8c, [+0x1c0]+0xa, +0xa5 |
| Defence | +0xbe |
| Skill.General and school skills 1..5 | +0xa8+2i; school skills also +0x116+2i |
| (computed) | +0xa6 = 0; +0x23c = index of the largest school skill, only when the template name contains `_Hero` or `Start_` |
| typeID, face | +0xe, +0x4b |
| gender | local; +0xe = 0x21 + gender for a fighter, 0x23 + gender for a mage (`L2.01282`, `L2.01283`) |
| AttackChargeTime, AttackRelaxTime | +0x134, +0x135 |
| TokenSize, MovementType | +0x49, +0x4a |

Template load also writes +0x14c from serverID ((v/10) mod 1000 when v
exceeds 10000), equips the row's items, sets +0x4c |= 6 when ManaMax > 0,
builds the spellbook +0x140 from knownSpells, writes the name-suffix
number to +0x4b when it is positive (`L2.01270`; the substring after the
dot and one more character, parsed by `L2.01284`), then calls skill
experience `L2.01285`, recompute (vtable +0x58) and vtable +0x5c, and sets
HP and mana to their maxima.

The producer then writes the four attributes (or 25 each, R2-ENGINE-287)
to +0x84..+0x8a, calls chosen skill `L2.01286` (R2-ENGINE-320), recompute,
HP and mana to their maxima, +0x80 = Player +0x18 name, +0x4 = the low
word of the `L2.01287` result, +0x14 = owner, +0x14c = 21 when session
+0x74 is 0, +0x1a4 = owner +0x30 and owner +0x38 = hero.

The +0x23c store (`L2.01288`) runs only when `L2.01289` finds `_Hero`
(string `L2.01290`) or `Start_` (`L2.01291`) in the template name at
[+0x3c]+4 (`L2.01292`..`L2.01293`); otherwise +0x23c keeps the Humanoid
initializer's 0. Every Start_ row qualifies.

**Confidence.** High for the call chain and every listed store, including
the +0x23c condition. Medium for the binding of each read to a title: the
row element accessor `L2.01294` was not read; the binding assumes it
returns parameter n. That `L2.01289` is a substring search follows from
its call shape; it was not read.

**Unknown.** Which writer leaves +0x4b = 32 in a Start_MF hero
(R2-SESSION-144) when template load writes the suffix number 5; a hero
built under a breakpoint, or the load-time path, would settle it. The
R2-ENGINE-209 positions with no direct displacement store in the 19
selected writer bodies: the raw 12 at [+0x10], +0x8, +0x18, the +0x20
count, +0xbb..+0xbd, +0x114..+0x117, +0x122..+0x12b, +0xd4..+0xd7, the raw
180 at [+0x1c0] except +0xa, the raw 184 at [+0x1c4], +0x50, +0x54, +0x58,
+0x5c, +0x60, +0x61, +0x64, +0x68, +0x6c, +0x78, +0x8e, +0x98, +0x9e,
+0xa2, +0xa3, +0x136, +0x138, +0x13c and +0x144. This bounds direct stores
only; a callee that receives the hero or an embedded object can still
write them. The Unit constructor `L2.01273` passes the hero to the base
constructor `L2.01295` (`L2.01296`), hero+0xa6 and hero+0x114 to `L2.01297`
(`L2.01298`, `L2.01299`), hero+0xbe to `L2.01300` (`L2.01301`) and
hero+0xd4 to `L2.01302` (`L2.01303`). None of these was read. They are
the candidate writers of the base fields +0x8, +0x18 and [+0x10], of
+0xbb..+0xbd, +0x114..+0x117 and +0x122..+0x12b, and of +0xd4..+0xd7.

### R2-ENGINE-320

Chosen skill `L2.01286(skill, 20)` removes the hand item through vtable
+0x48, zeroes school skills 1..5 (+0xaa..+0xb2), writes the chosen skill
20, writes +0xb2 = 20/2 = 10, writes +0x23c = the chosen skill and calls
skill experience `L2.01285`. That routine copies school skill i to
+0x116+2i, writes E(s) = trunc((1.1^s − 1)·1000) to +0x23c+4i for
i = 1..5 (R2-ENGINE-290) and their sum to +0x130. Vtable +0x5c `L2.01304`
writes +0x1c = trunc(0.01 x the sum of +0x240..+0x250). Of the selected bodies
only template load calls it (`L2.01305`), before the producer calls chosen
skill. Chosen skill then adds the skill weapon through vtable +0x44
(R2-ENGINE-322). Recompute restores +0xa8+2i from +0x116+2i and adds the +0xe8
skill modifiers.

**Confidence.** High: every store is read in the four bodies.

### R2-ENGINE-321

Recompute `L2.00052` (Human vtable +0x58) caps the four attributes by
class and sex plus the signed bytes +0xd4..+0xd7, writes +0x96 and +0x9c
by the R2-ENGINE-290 formulas, +0x8c speed, +0x92 from Body, +0x90 from
+0x8e and bag +0x20, +0xa4 sight, +0xa6, +0xb4 and +0xb5, restores the
skills (R2-ENGINE-320), zeroes +0xb7..+0xba, clears 22 bytes at +0xbe
(`L2.01306`), sets +0xbe = Reaction/3 and +0xc2+2i = Spirit/2 (i = 1..5),
then applies `L2.00102` on the record at +0xd4: +0xd8 into +0x8c, +0xda
into +0x92, +0xdc into +0x96, +0xe0 into +0x9c, +0xe4 into +0xa4, the
+0xe6 block into +0xa6 and the +0xfe block into +0xbe. It bounds HP and
mana by their maxima, writes +0xa0 from +0x9c and Player +0xa5c, and the
low byte of speed into [+0x1c0]+0xa.

**Confidence.** High for the stores and their sources. The RU arithmetic
uses the constants at `L2.01307`, `L2.00067`, `L2.00107` and `L2.01308`
and the helpers `L2.00109`, `L2.00108` and `L2.01309`, none read here; the
RU evidence for the R2-ENGINE-290 formulas is the matching instruction
shape and one hero whose HP 149 and sight 1617 they reproduce exactly
(R2-SESSION-144). Medium for the names of +0xbe (defence) and +0xc2+2i
(resistances), which follow R2-ENGINE-290.

**Unknown.** The writers of +0x8e and +0xd4..+0xd7, and the meaning of
+0x90, +0x92, +0xa0, +0xa6, +0xb4, +0xb5 and +0xb7..+0xba.

### R2-ENGINE-322

Add item (vtable +0x44, `L2.01310`) calls the equip test (vtable +0x40,
`L2.01311`). It tests the item class against Armor `L2.01176`, Weapon
`L2.01175` and Shield `L2.01312`, item parameter 0xf and the mage
predicate `L2.00044`, then calls item vtable +0x38; an item the test
returns goes into the bag [+0x7c] through `L2.01313`.

Weapon equip `L2.01314` writes +0x74 = weapon; for weapon parameter 0xe
equal to 2 it moves +0x78 to the bag. It writes the modifier fields +0xe6,
+0xf4, +0xf5, +0xf9..+0xfb and +0xfe, +0xb6 = weapon parameter 5 when below
10, else 0, calls recompute, takes +0x134 and +0x135 from weapon
parameters 0xc and 0xd, and adds weapon +0x58 − 1 to +0x12c. Armor equip
`L2.01315` writes hero +0x208 + 4·(armor +0x58) = armor and adds the
armor's 22-byte block +0x5a into hero +0xfe and +0xbe (`L2.01316`).
Template load allocates +0x140 only when ManaMax > 0.

**Confidence.** High for the stores. Medium for parameter 0xe = 2 as a
two-handed weapon and +0x78 as the off hand: no weapon row was read.

**Unknown.** The writer of +0x78 and the bag object's fields.

### R2-ENGINE-323

Player body `R2.0085` (R2-SESSION-020) writes the name, then the 51-byte
prefix: u16 +0x4, u32 +0x8, raw 8 +0x10, u8 +0xa44, u32 +0x2c, u16 +0x30,
u32 +0x3c, u8 +0x40, u8 +0x41, u32 +0xa48, u32 +0xa50, u16 +0xa54, u16
+0xa4c, u32 +0xa5c, u32 +0x38 and u32 = the Player's own address. Store and
load pass +0x3c and +0xa48 through `L2.01271`, an XOR with 0x5c073f4d. Load
calls `L2.01168`([`L2.01317`]+0xdc, saved own address, Player)
(`L2.01318`), reads the 2560 raw bytes, the group list (`R2.0122`) and the
36 raw bytes, then passes +0x38 to `L2.01319` and stores the Player
directly into each member's +0x14 (`L2.01320`). The Unit base `L2.01164`
writes the Unit's own address (`L2.01321`), calls `L2.01168` with it on
load (`L2.01322`) and passes +0x14 to `L2.01169`. `L2.01169` and
`L2.01319` are identical instruction for instruction, address aside: each
calls `L2.01323` on the same map with the saved value and an out slot and
stores the out value when the call returns nonzero, else 0. Read as a
map, `L2.01168` inserts saved address to object and `L2.01323` looks it
up.

**Confidence.** High for the field order, the XOR and the call sequence
of both arms. Medium that `L2.01168` registers and `L2.01323` resolves
saved addresses: neither callee was read.

### R2-ENGINE-324

Group body `R2.0123` (R2-ENGINE-191, R2-ENGINE-192) writes, after the
member list, u32 +0x1c, u32 +0x40 and u32 +0x44 (`L2.01324`, `L2.01325`,
`L2.01326`) and returns. Load reads the three, resolves +0x40 through
`L2.01319` (`L2.01327`) and +0x44 through `L2.01169` (`L2.01328`). Player
body `R2.0085` calls the group list `R2.0122` and then `R2.0175` on
[+0x34], which transfers 36 raw bytes, and returns.

**Confidence.** High: both arms of `R2.0123` and the Player body's tail
are read.

**Unknown.** The meaning of Group +0x1c and +0x40, and of the 36 bytes.

## Shop pool in every location

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-327 | ROM2 armour admission excludes (class, material, row) (0,0,2), (0,1,2) and (6,4,6) by program immediates in both locales; the other admission modes test data.bin fields or other program constants. | High | ✔ promoted (branch candidate) | [EXP-2039](../experiments/EXP-2039-rom2-shop-pool/) |
| R2-ENGINE-328 | No ROM2 shop fill admits a beard: the two beard-named armour keys 0x0502 and 0x1502 are removed only by the armour exclusion, from any record; a sold beard can still be shelved. | High / Unknown | ✔ promoted (branch candidate) | [EXP-2039](../experiments/EXP-2039-rom2-shop-pool/) |
| R2-ENGINE-329 | The ROM2 Scenario.dll shop table holds four zero-filled location records; 62 operand sites per DLL reach it, from NewGame, departure, TalkTo, Save, Load and the lookup. No direct store writes record 0. | High / Unknown | ✔ promoted (branch candidate) | [EXP-2039](../experiments/EXP-2039-rom2-shop-pool/) |
| R2-ENGINE-330 | ROM2 NewGame writes shop records 1 and 2; departures rewrite record-2 masks after missions 30, 40, 60, 80 and 90, write record 3 after mission 50 and raise record 2 and 3 price bounds at stages 30..110. | High / Medium | ✔ promoted (branch candidate) | [EXP-2039](../experiments/EXP-2039-rom2-shop-pool/) |
| R2-ENGINE-331 | ROM2 shop record 3 category 3 has no class bit until TalkTo kind 0, NPC 0x2a3, topic 0x4e ORs in class 4; EnterInn offers that word at location 3 when stage > 60 and bank 770 is 1. | High / Medium | ✔ promoted (branch candidate) | [EXP-2039](../experiments/EXP-2039-rom2-shop-pool/) |
| R2-ENGINE-332 | Under the record values the Scenario.dll direct stores write, ROM2 shops admit 458 data.bin keys (190 armour, 35 shield, 146 weapon, 64 MagicItems, 23 books); none selects class 5 or 6. | High / Unknown | ✔ promoted (branch candidate) | [EXP-2039](../experiments/EXP-2039-rom2-shop-pool/) |
| R2-ENGINE-333 | ROM2 sold-item placement L2.00804 tests the item tag, type and stackability and the category kind bits and bit 29, never class, material or row; a stackable item prefers a category without bit 29. | High | ✔ promoted (branch candidate) | [EXP-2039](../experiments/EXP-2039-rom2-shop-pool/) |
| R2-ENGINE-334 | The ROM2 server object's init sets +0x74 to (mode < 2); campaign start passes mode 2 (Medium), so shop fills read the Scenario.dll record; network paths pass 0; other +0x74 stores are unclassified. | High / Medium | ✔ promoted (branch candidate) | [EXP-2039](../experiments/EXP-2039-rom2-shop-pool/) |

### R2-ENGINE-327

Table admission `L2.00812` (RU `L2.01331`) dispatches on mode through the
jump table at `L2.01332`. The mode-1 armour arm `L2.01333` (RU `L2.01334`)
compares its class, material and row arguments with immediates in three
groups:

| Triple | EN compares | RU compares | Item ID | itemname line |
|---|---|---|---|---|
| (0, 0, 2) | `L2.01333`, `L2.01335`, `L2.01336` | `L2.01334`, `L2.01337`, `L2.01338` | `0x0502` | 17 |
| (0, 1, 2) | `L2.01339`, `L2.01340`, `L2.01341` | `L2.01342`, `L2.01343`, `L2.01344` | `0x1502` | 3 |
| (6, 4, 6) | `L2.01345`, `L2.01346`, `L2.01347` | `L2.01348`, `L2.01349`, `L2.01350` | `0x46c6` | 64 |

The armour constructor `L2.01351` stores item ID material<<12 | part<<8 |
class<<5 | row at +0x40, part being armour column 4 (5, 5 and 6 here). The
three IDs name Magic Beard / Магическая Борода, Beard / Борода and Gold
Crown / Золотая Корона (R2-ENGINE-267 name map).

Other arms: weapon mode 2 (`L2.01352`) keeps weapon column 15 bit 0 set;
weapon mode 8 (`L2.01353`) keeps it clear; shield mode 7 (`L2.01354`) has
no extra test; modes 3..6 (`L2.01355`) build nothing. Magic admission
`L2.00810` skips book spells {9, 14, 15, 24, 28, 29}, picks MagicItems row
i+5 or i+34 for i in 1..29, adds potions by name (rows 69..74) and never
builds rows 64..68 or 75..96. A bit-29 category removes a row when
pow(min, 0.4) exceeds trunc(material column 8 x class column 8).

**Confidence.** High. All admission and mode bodies are complete linear
decodes with in-body branch targets in both locales.

### R2-ENGINE-328

The only two beard names among the 491 itemname lines of each locale
(casefold `beard`, `борода`) are IDs `0x0502` and `0x1502`: armour row 2,
class 0, materials 0 and 1, part 5. Their class mask words admit both
materials. Only record 1 category 0 (both keys, range 0..1500) and
record 2 category 0 under its NewGame mask `1383c305` (material 0, range
0..5000 or 0..10000) select them; their admission prices, 100 and 25, lie
inside those ranges, so the armour exclusion (R2-ENGINE-327) is their only
removing test. Every later record-2 and record-3 category-0 mask clears
material bit 0 and selects neither, and their minima (499, 12000, 40000)
would exclude both prices. The exclusion sits in table admission, which every fill uses
for armour, including fills from the deal source (R2-ENGINE-334). No other
admission mode builds armour, and the draw takes only admitted candidates
(R2-ENGINE-265).

Sold-item placement (R2-ENGINE-333) tests no row, and server sale
`L2.00832` places every pending hero item with an owner and a nonzero
price.

**Confidence.** High that no fill admits a beard. Unknown whether a
player sale can place one: the client's admission of a beard to the
pending sale list was not read.

**Unknown.** The client sale-selection producer.

### R2-ENGINE-329

Ordinal 15 (R2-ENGINE-263) returns `D2.00104 + ID*0x50` with no bound
check. Records 0..3 span `D2.00104..D2.00232`; `D2.00007` holds the
current-location pointer. Both DLLs' .data has raw data to `D2.00233` and
virtual end `D2.00234`, so the records start zero. A record with zero
masks builds no candidate.

An operand census of every executable section for displacements or
immediates in `D2.00235..D2.00236` finds 62 sites per DLL, each matched by
a raw dword scan of every section at every byte: 16 in ScenarioNewGame,
41 in departure `D2.00012`, a read and a write in TalkTo `D2.00137`, and
the record address in ScenarioSave `D2.00237`, ScenarioLoad `D2.00238`
and ScenarioGetShopAssortment `D2.00239`. EN and RU sites are equal. No
direct store writes record 0; ScenarioLoad's copy at `D2.00238` restores
every record, record 0 included, from the bytes Save wrote
(R2-ENGINE-263). The census counts direct stores only; bulk copies and
writes through a computed pointer are not stores it can classify.

**Confidence.** High for the DLL census population.

**Unknown.** Writes through the record pointer the lookup returns, in the
client and inside the DLL; neither was searched.

### R2-ENGINE-330

| Writer | Record | Values |
|---|---|---|
| NewGame | 1 | R2-ENGINE-263 |
| NewGame | 2 | masks `1383c305`, `2bc3c304`, `04000000`, `10438305`; 0..5000; draws 100/20/20/100; bounds 2/1/1/2 |
| leave 30 | 2 | masks cat 0/1/3 `1383c22c`, `2bc3c324`, `10438324` |
| leave 40 | 2 | `1387c26c`, `2bc7c324`, `10478264` |
| leave 50 | 3 | min 499/499/0/499; draws 100/20/20/100; bounds 2/1/1/2; masks `13c7d318`, `2bc7d318`, `04000000`, `2bc042f0` |
| leave 60 | 2 | `1387c268`, `2bc7c268`, `10478260` |
| leave 80 | 2 | `138752e0`, `2bc752e0`, `104702e0` |
| leave 90 | 2 | `138772c0`, `2bc772c0`, `104722c0` |
| stage 30, 40 | 2 | max 10000, 22000 |
| stage 50..80 | 2, 3 | max 60000, 150000, 400000, 800000 |
| stage 90 | 2, 3 | max 1500000; record 2 min 12000 (cats 0, 1, 3) |
| stage 100 | 2, 3 | max 5000000; record 2 min 40000 (cats 0, 1, 3) |
| stage 110 | 2, 3 | max 10000000 |

LeaveLocation `D2.00009` calls departure `D2.00012` when the current
location's type is not 2. Departure switches on the location ID (byte
table `D2.00025`), adds 10 to stage bank 768 (`D2.00118`) when ID%10 is 0
and switches on the new stage (`D2.00028`). Town IDs 1, 2 and 3
(R2-ENGINE-231) index records 1, 2 and 3.

Admitted keys per category, record 2, in departure order 10..110 from
stage 10: 56/75/39/33 at NewGame, 60/80/51/36 after 20, 107/130/71/47
after 50, 42/120/86/33 after 90. Record 3: 43/69/71/0 after 50,
50/71/87/0 after 100.

**Confidence.** High for the writer values and switch conditions. Medium
for the per-state counts as a campaign order: they assume departures in
ID order through the non-type-2 path.

**Unknown.** Whether town 3 is entered after a bank-775 restoration
(R2-SESSION-123) with record 3 still zero.

### R2-ENGINE-331

TalkTo `D2.00137` splits its word into NPC (bits 0..15), topic (bits
16..27) and kind (bits 28..30). Kind 0 with NPC 0x2a3 and topic 0x4e sets
bank 770 to 2 and ORs `0x80000` (class 4) into record 3's category-3 mask
at `D2.00240` (`D2.00241..D2.00242`). Mask `2bc042f0` has no class bit, so
category 3 admits nothing before that; `2bc842f0` admits 38 keys (23
armour, 15 weapon). EnterInn `D2.00089` offers word `L2.01356` at
`D2.00100` when stage bank 768 > 0x3c, bank 770 is 1 (`D2.00215`) and the
current location ID is 3. The mission-70 departure arm sets bank 770 to 1.
Bank 770 is the druid category-3 gate query 0x302 (R2-ENGINE-259).

**Confidence.** High for the TalkTo and EnterInn bodies. Medium that
EnterInn is the only offer of the word: dword scans of the EN DLL and both
clients found one site.

### R2-ENGINE-332

Over every legal key of both installed data.bin tables (row class mask
word carrying the material bit), every MagicItems row and spells 1..29,
across every record state the direct DLL stores of R2-ENGINE-330 and
R2-ENGINE-331 produce, and the widest-bound control:

| Table | Legal | Admitted | Never admitted |
|---|---|---|---|
| armour | 204 | 190 | 2 beard keys by constant; 11 keys of class 5 or 6; row 28 class 2 material 11 by material selection |
| shield | 37 | 35 | row 5 class 5 by class; row 8 class 1 material 1 by price 1600 > 1500 |
| weapon | 154 | 146 | 8 keys of class 5 or 6; rows 1, 23..27 have no legal key |
| MagicItems | 96 | 64 | rows 1..5 built by the book path; 64..68 and 75..96 never built |
| books | 29 | 23 | spells 9, 14, 15, 24, 28, 29 |

The OR of all directly stored masks is `0x3fcfffff`, classes 0..4. Union per
record and category: 70/44/4/34, 151/233/87/95, 50/71/87/38. EN and RU
data.bin are byte-equal and verdicts are equal.

Inputs: record values are DLL immediates; legality is the data.bin row
mask word; price is row column 2 x material column 2 x class column 2;
the armour exclusion, book skips, phase-two row arithmetic and potion
names are program constants; weapon modes read column 15 bit 0; the
enchant level reads material and class column 8.

**Confidence.** High for admission under those record values. Unknown
which bit-29 keys survive the enchanted draw outside record 1: enchant
`L2.01357` is unread.

**Unknown.** Record values outside the direct stores: writes through the
lookup's pointer in the client or the DLL (R2-ENGINE-329), a loaded SAV's
bytes, and deal-source records (R2-ENGINE-334). The `L2.01357` enchant and
its final-price failures.

### R2-ENGINE-333

Placement `L2.00804` (RU `L2.01358`) uses a nonzero tag +0x4d as category
tag-1. Otherwise it sets kind `0x4000000` for types 3, 4 and 5,
`0x400000` for type 2 and `0x1000000` for type 1. A non-stackable
(vtable +0x50) non-book goes to the first category with bit 29, else
category 3. A stackable item or a book goes to the first category with its
kind bit and no bit 29, else the first with its kind bit, else category 3.
Insert `L2.00833` merges into an equal stackable entry or appends to the
shop category list ([deal+0x9c]+0xc+k*0x1c); an appended item also enters
the deal list +4+k*0x1c. The body reads no class, material or row.

**Confidence.** High: complete body with its insert.

### R2-ENGINE-334

Server global `L2.00805` (RU `L2.01317`): constructor `L2.01359` (RU
`L2.01360`) writes +0x74 = 0; init `L2.00221(mode)` (RU `R2.0010`)
writes +0x74 = (mode < 2) and +0x70 = (mode > 0); creator `L2.01003` (RU
`L2.01019`) stores the global and calls init. Campaign start callers EN
`L2.01004`, `L2.01005` (RU `L2.01361`, `L2.01362`) store app+0x3ec = 1 and
app+0x5d8 = 2 (RU +0x404, +0x63c) and pass mode 2. EN `L2.01006`,
`L2.01007` (RU `L2.01363`, `L2.01364`) pass 0. EN `L2.01008` (RU
`L2.01365`) passes app+0x5d8, whose stores are 0, 2, 2, 0, 0, 3, 3 and 1.
Of 14 methods called through the global, only init writes +0x74. A .text
linear sweep finds 119 (EN) and 121 (RU) other dword [reg+0x74] stores.

**Confidence.** High for the committed bodies: constructor `L2.01359`
writes 0, init `L2.00221` writes (mode < 2) and (mode > 0), creator
`L2.01003` stores the global and passes its argument to init. Medium for
the caller set, the caller arguments, the 14-method census and that a
single-player campaign keeps +0x74 at 0: they come from the client linear
sweeps and the global raw scan, which ran beyond the preregistered scan
budget, and the load that puts app+0x5d8 into each pushed register is not
in a committed listing.

**Unknown.** The owners of the unclassified +0x74 stores and the mode
`L2.01008` passes in a campaign.

## Music selection and playback

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| R2-ENGINE-335 | ROM2 screens request fixed music keys: menu 0x421 menu, New Game 0x425 chrgen, town 0x42e/0x468 b14/b16/b15, map 0x42d/0x41d map, credits 0x428 credit; menu, chrgen and town skip a list already holding the key. | High | ✔ promoted (branch candidate) | [EXP-2040](../experiments/EXP-2040-rom2-music/) |
| R2-ENGINE-336 | A ROM2 mission sets the list B00..B16; every 16 ticks the type-12 music area holding the hero, or the map's default record, picks one of its four themes at random, and the next track end plays that list index. | High / Medium | ✔ promoted (branch candidate) | [EXP-2040](../experiments/EXP-2040-rom2-music/) |
| R2-ENGINE-337 | The ROM2 music player streams one list entry through a looping DirectSound buffer; at a track end it plays the area pick, else the held index, else the next list entry; stop pauses and start resumes. | High / Medium | ✔ promoted (branch candidate) | [EXP-2040](../experiments/EXP-2040-rom2-music/) |
| R2-ENGINE-338 | While no dialogue tune is held, ROM2 music stops for logos, every cutscene, mission entry, mission end and message 0x451 and resumes only on a later screen request; a stream read failure posts 0x486. | High / Medium | ✔ promoted (branch candidate) | [EXP-2040](../experiments/EXP-2040-rom2-music/) |
| R2-ENGINE-339 | ROM2 stores SoundRandom, SoundMusPos, SoundSfxPos, SoundSpeechPos and MusicEnabled in HKLM; volume defaults to -700, MusicEnabled to 1; -nomusic or a music.res mount failure disables screen and mission music requests. | High | ✔ promoted (branch candidate) | [EXP-2040](../experiments/EXP-2040-rom2-music/) |
| R2-ENGINE-340 | The ROM2 sound panel turns music on (resume or restart the chosen melody) and off (stop or fade), stores the chosen melody, the random flag and three volumes, and names melodies from main.res text/tunes.txt. | High / Medium | ✔ promoted (branch candidate) | [EXP-2040](../experiments/EXP-2040-rom2-music/) |
| R2-ENGINE-341 | A ROM2 dialogue tag tune=N cross-fades music to list index N, which repeats only while no area pick exists; no data file or archive entry in either root contains tune=, so shipped dialogue never changes music. | High | ✔ promoted (branch candidate) | [EXP-2040](../experiments/EXP-2040-rom2-music/) |
| R2-ENGINE-343 | On a -1 find the ROM2 dialogue loader leaves a one-byte empty body, a clear flag and no error (EN High, RU Medium); that absent means -1, and the TALK body ignoring the result (TalkTo), are Medium. | High / Medium | ● active (branch candidate) | [EXP-2041](../experiments/EXP-2041-rom2-inn-talk-missing/) |
| R2-ENGINE-344 | The `#<key>` lookup is reached via the dialogue builder (21 direct call sites, 10 functions per image; register calls unseen); only TALK formats `npc%dtalk%d`; no save-load read found. | Medium | ● active (branch candidate) | [EXP-2041](../experiments/EXP-2041-rom2-inn-talk-missing/) |

### R2-ENGINE-335

Each body tests the music-available word (EN `L2.00878`, RU `L2.01366`)
and does nothing to music when it is zero. The application dispatcher
(EN jump tables `L2.00794`/`L2.00795`, base 0x416) reaches the bodies.

| Screen | Message | EN / RU body | Key | Same-key test |
|---|---|---|---|---|
| Main menu | 0x421 | `L2.01096` / `L2.01367` | `music\menu.wav` | yes |
| Pre-create (New Game) | 0x425 | `L2.01100` / `L2.01368` | `music\chrgen.wav` | yes |
| Town | 0x42e, 0x468 | `L2.00266` / `L2.00267` | `music\b14.wav`, `b16`, `b15` (R2-ENGINE-231) | yes |
| World map | 0x42d, 0x41d | `L2.00876` / `L2.00877` | `music\map.wav` | no |
| Credits | 0x428 | `L2.01369` / `L2.00268` | `music\credit.wav` | no |

A body with the same-key test calls `L2.01370` (RU `L2.01371`), which
compares the requested key with the player's current list entry. When they
match it skips the list set and only calls start, so the track resumes from
its paused position (R2-ENGINE-337). Otherwise, and always for the map and
credits, it sets a one-entry list and starts it from its first byte. The
pre-create body is also called twice from the view-close body EN
`L2.00787` (`L2.01372`, `L2.01373`). The six list-set callers and eight
start callers of the EN rel32 call census are these five bodies, the
mission body (R2-ENGINE-336) and the sound panel (R2-ENGINE-340).

`music.res` holds the keys without the `music\` prefix and in lower case;
the mission keys are written `B00`..`B16`. Each `allods2.exe` holds 24
`music\` literals; the five screen bodies and the mission body push all
24, and the raw reference scan of the sampled 10 EN and 8 RU literals
finds one reference each.

**Confidence.** High: every body is decoded completely in EN; the RU
bodies have equal normalized mnemonics except the main menu (208 EN, 210
RU instructions), whose key and same-key test were read in RU.

**Unknown.** A screen reached only through an indirect call; the RU
fourth pre-create caller at RU `L2.01374`; the resolver's case rule for
archive keys (inferred case-insensitive, since `B00` must reach `b00.wav`).

### R2-ENGINE-336

Mission body EN `L2.01375` / RU `L2.01376` builds the 17-entry list
`music\B00.wav`..`music\B16.wav` and calls list set EN `L2.01377`, which
stops the player, copies the list, opens entry rand() mod 17, and sets
the area pick EN `L2.01378` / RU `L2.01379` to -1. If the mission view
has a hero object (view+0x3f6c), area select EN `L2.01380` / RU
`L2.01381` runs on the hero's +8/+0xc position and, for a pick of at least
0, select opens that list index. Start follows. It has no same-key test.
The random source is `L2.00729`, the step R2-ENGINE-242 bounds.

Callers: mission entry EN `R2.0026` at `L2.01382` (skipped when
application +0x5d8 is 3) and view-close EN `L2.00787` at `L2.01383` and
`L2.01384`, when the closed view is application +0xf8, +0xfc or +0x100
(or +0x104, +0x108, +0x10c) and +0x5d8 is not 2. Closing such a view
sets the list again, as at mission entry.

Mission load EN `L2.01385` / RU `L2.01386` fills global array EN
`L2.01387` (RU `L2.01388`) of 28-byte records from the map's type-12 data
(R2-ASSET-084): the head record at map+0x374 first, only when its theme 0
is at least 0, then the area array at map+0x360. Area select, per record:

1. A record at (0,0) is the default. It is taken, with distance 1e10,
   when the best distance is still above 1e15.
2. Another record with all four themes -1 is skipped.
3. Otherwise distance is the Euclidean distance from the position to
   (x·256, y·256). The record is a candidate when distance < radius·256
   and is taken when distance is below the best so far (initially 1e20).

When a record is taken the pick is theme[R(3)], where R(n) is
(rand()·(n+1))>>15 (EN `L2.00939`), drawn again while the theme is -1.
Without a taken record the pick keeps its value. Mission tick EN
`L2.01389` / RU `L2.01390` calls area select when its counter +0xa88 has
low four bits 0, the music-available word is set and the hero object
exists. The player reads the pick only at a track end (R2-ENGINE-337), so
a new area changes the music after the current track.

Only set list writes -1 to the pick. On a map whose head record is
admitted (every campaign map, R2-ASSET-084) the default record is taken
whenever no area holds the hero, so every area select leaves a pick of at
least 0:

- When the hero object exists at mission entry, the mission body's own
  area select replaces the random entry before start; the first entry is
  a theme of the area holding the hero or of the head.
- Every later track end opens a theme drawn at most 16 ticks earlier; the
  next-entry branch (index+1) mod n is not reached.
- The random start entry plays only when no hero object exists at mission
  entry, and only until the first track end after an area select.

On a map without an admitted head (the root maps) the list plays in order
from its random start until the hero first enters an area; after that the
last pick stays in force, since nothing else in a mission resets it.

**Confidence.** High for the list, the record walk and the theme draw in
both locales (equal normalized mnemonics). Medium for the array
population and the consequences above that depend on it: the head append
`L2.01391`, the area append `L2.01392` and the count and index helpers
`L2.01393` and `L2.01394` are not decoded, and eight references to
`L2.01387` (`L2.01395`..`L2.01396`, `L2.01397`..`L2.01398`, `L2.01399`,
`L2.01400`) lie in unread bodies. Whether the head stays in the array and
whether the array is cleared between missions is therefore Medium. Medium
that the object at view+0x3f6c is the player's hero and that its
+8/+0xc are map units of 1/256 tile; the scale is inferred from the shift
by 8.

**Unknown.** The identities of the closed views and of mode 3; the
semantics of the four helpers and the eight unread `L2.01387` references;
the pick after a mission ends while a later screen leaves it at a stale
value.

### R2-ENGINE-337

Player object application +0xc8, created by EN `L2.01401` / RU
`L2.01402` with a streaming buffer of 0x56000 or 0xac000 bytes (chosen by
EN `L2.00648`) and initial volume SoundMusPos. Methods (EN / RU):

| Method | EN / RU | Behaviour |
|---|---|---|
| set list | `L2.01377` / `L2.01403` | stop, copy list, select rand() mod n, pick = -1 |
| select | `L2.01404` / `L2.01405` | open entry n from its start |
| start | `L2.01406` / `L2.01407` | needs MusicEnabled and a buffer; re-arms the timer after a stop; Play with the looping flag |
| stop | `L2.01408` / `L2.01409` | not holding: kill the timer and Stop the buffer; holding: select the held index and restore SoundMusPos |
| fade | `L2.01410` / `L2.01411` | volume ramp, called with (2000, 8000) |
| cross-fade | `L2.01412` / `L2.01413` | hold flag +0x18 = 1, held index +0x1c = n, fade |
| set random | `L2.01414` / `L2.01415` | +0x20 = value, hold flag = 0 |

Refill EN `L2.01416` streams the open entry into the looping buffer. At an
entry's end it opens the next one: the area pick when at least 0, else
the current index when the hold flag is set, else (index+1) mod n. A
one-entry list therefore repeats its key. The mission list advances in
list order only while the area pick is -1 (R2-ENGINE-336). The timer EN `L2.01417` drives refill and
the fade; a fade that reaches the end calls stop. DirectSound Stop keeps
the play position, so start after stop resumes the paused entry.

**Confidence.** High for the methods and the next-entry rule; every body
is decoded in both locales with equal normalized mnemonics. Medium for
audible resume: it rests on documented DirectSound Stop/Play semantics,
not on an observation.

**Unknown.** The reader of +0x20 (none in the player class range
EN `L2.01418`..`L2.01419`); the effect of stop called by list set while
the hold flag is set.

### R2-ENGINE-338

The EN rel32 call census finds ten stop callers: the sound panel (two),
the player's own timer, list set and destructor, and the five below.

| Caller | EN site | Context |
|---|---|---|
| Logos `L2.01420` | `L2.01421` | startup logos |
| Cutscene `L2.00875` | `L2.01422` | every movie: intro, 0x41d, 0x436, 0x42d, view-close |
| Mission entry `R2.0026` | `L2.01423` | before the mission list (R2-ENGINE-336) |
| Mission end `L2.01424` | `L2.01425` | 0x41d, 0x45c, 0x45d, exit |
| Message 0x451 arm `L2.01426` | `L2.01427` | posted from view-close |

The cutscene and logos bodies contain no start call. Music resumes only
when the next screen requests it; a same-key screen resumes the paused
entry (R2-ENGINE-335). These are stops only while the hold flag is clear.
With the flag set by a dialogue tune (R2-ENGINE-341), each stop call
instead selects the held index from its start and restores SoundMusPos,
and the flag stays set (R2-ENGINE-337). On a refill read failure the player clears +0x10
and posts 0x486; the application arm for 0x486 posts WM_CLOSE (0x10).
Startup order in the application initializer EN `L2.01428`: registry load
`L2.01242`, player creation `L2.01429`, logos `L2.01430`, intro cutscene
`L2.01431`; the main menu then requests `menu.wav`.

**Confidence.** High for the stop sites, the hold-flag branch and the
absence of a start call in the cutscene and logos bodies. Medium for the
0x486 consequence, including whether a read at an exact track end returns
0, and the context labels of the 0x41d, 0x45c and 0x45d arms.

**Unknown.** Stops reached through indirect calls; the sender and meaning
of 0x451.

### R2-ENGINE-339

Settings object EN `L2.01432` / RU `L2.01433` (no file bytes; constructor
EN `L2.01434` / RU `L2.01435`, static initializer EN `L2.01436`):

| Offset | Registry value | Meaning | Constructor value |
|---|---|---|---|
| +0x00 | SoundRandom | random flag | none (static 0) |
| +0x08 | SoundMusPos | music volume, DirectSound hundredths of a dB | -700 |
| +0x0c | | music scale | 5000 |
| +0x10 | SoundSfxPos | effects volume | -700 |
| +0x14 | | effects scale | 5000 |
| +0x18 | SoundSpeechPos | speech volume | -700 |
| +0x1c | | speech scale | 5000 |
| +0x20 | | music available (EN `L2.00878`) | 1 |
| +0x24 | MusicEnabled | music on (EN `L2.01437`) | 1 |

Load EN `L2.01438` / RU `L2.01439` reads the five values with
RegQueryValueExA into those offsets; store EN `L2.01440` / RU `L2.01441`
writes them. They are called from the registry load and store of
R2-ENGINE-311, so the key, the once-at-startup load and the store at exit
and after a cutscene are those of TipsMode. A missing value keeps the
constructor value.

Music available is not stored. The initializer clears it when the
`music.res` mount throws (EN `L2.01442`) or when the command line contains
`-nomusic` (EN `L2.01443`); a failed sound initialization also clears
it. Every screen request, the mission tick and the mission list test it.
The sound panel's music-on arm does not: the EN census of `L2.00878`
has no site in the handler `L2.01444`..`L2.01445`, and player creation
always builds the player. MusicEnabled gates start only (R2-ENGINE-337).

**Confidence.** High: the constructor, load and store are decoded in both
locales with equal normalized mnemonics.

**Unknown.** Values outside the panel's range written to the registry by
another program; what the panel's music-on arm plays while music is
unavailable.

### R2-ENGINE-340

Panel handler EN `L2.01444` / RU `L2.01446` (table slot +0x48 of EN
`L2.01447`), messages 0x467..0x478 through EN tables `L2.01448`/`L2.01445`:

| Message | Action |
|---|---|
| 0x477 music on | MusicEnabled = 1; when the chosen melody is the current index: set volume, start; else stop, select it, set volume, start |
| 0x478 music off | player state 2: stop; state 1: fade (2000, 8000); then MusicEnabled = 0 |
| 0x46e, wParam 2 | SoundRandom = lParam; player set random (clears the hold flag) |
| 0x46e, wParam 3 | chosen melody = lParam (no player call) |
| 0x46e, wParam 6, 7, 8 | music, effects, speech volume from the slider through EN `L2.01449`, clamped to -10000..0 |
| 0x474 | effects or speech test sound |
| 0x476 | close |
| 0x467, other | no action |

Panel build EN `L2.01450` lists the current list's entries. Each name is
the entry key without its first six characters (`music\`) looked up in
the map EN `L2.01451`, which tune-name fill EN `L2.01452` builds from
`main\text\tunes.txt` (loaded at EN `L2.01453`) by splitting each line at
`=`. That file names `credits.wav` while the archive key and the literal
are `credit.wav` (R2-ASSET-083), so the credits entry has no name.

**Confidence.** High for the EN arms and the RU handler alignment (equal
normalized mnemonics). Medium for the RU tune-name path: the RU build and
fill bodies are not decoded. Medium for the slider curve, read from
`L2.01449` only in EN.

**Unknown.** The on-screen layout of the panel; the player state values
the off arm tests beyond 1 and 2.

### R2-ENGINE-341

Dialogue parser EN `R2.0048` / RU `R2.0047` reads `tune=` (literal EN
`L2.01454`) with `%d`; page EN `R2.0046` / RU `R2.0045` calls
cross-fade EN `L2.01412` with the value when it is at least 0. The
cross-fade is its only caller. After the fade, stop selects the held
index. The refill tests the area pick before the hold flag, so the held
index repeats at a track end only while the pick is -1; in a campaign
mission (R2-ENGINE-336) it plays once and the area theme follows. The
flag stays set until set random clears it (R2-ENGINE-337,
R2-ENGINE-340).

The census reads every file of both roots and every entry of every `.res`
archive as raw bytes, ignoring case: `tune=` occurs only in EN and RU
`allods2.exe` and RU `a2server.exe`. Declared instant 37 is not a music
path (R2-ENGINE-058).

**Confidence.** High for the parser and page path in both locales and for
the absence within the census population. A compressed payload or a file
outside the two roots is outside it.

### R2-ENGINE-343

The section loader EN `R2.0042` / RU `R2.0041` (R2-ENGINE-050) formats
`#%s` (EN `L2.01455`) from the dialogue key at object +0x6c, lowercases it
and calls `strstr`-shaped find EN `L2.00997` on the shared text buffer (EN
`L2.01191`, RU `L2.01456`). The EN `L2.00997` listing (in the lane notes,
not in `evidence/`) tests its callee `L2.01457` result and returns -1 for
NULL and the offset otherwise; `L2.01457` was not read, so "returns -1 when
the section is absent" rests on an unread callee. The RU find `L2.01289`
was not read. On -1 the loader
stores +0x78 = 1, allocates one byte, stores the pointer at +0x74, writes
NUL to it, clears +0x7c and returns. No error call, no exit and no second
lookup follow. The found arm sets +0x7c to 1 only when the cut body holds
`npc` (R2-ENGINE-050). The EN dialogue builder `R2.0038` / RU `R2.0037`
tests +0x7c right after the constructor: when clear it takes a different
construction arm (EN `L2.01458`, one 0x98-byte child instead of a 0x64-byte
and a 0x98-byte child). The arm contents, the constructor `R2.0040` and
its loader call `L2.01459` were read in the lane notes and are not in
`evidence/`. The parser EN `R2.0048` scans the body for `<` and
jumps to `L2.01460` on a NUL; that target was not read.

The TALK body EN `L2.00554` / RU `L2.00555` formats `npc%dtalk%d` (EN
`L2.01461`, RU `L2.01462`), calls the builder, overwrites `eax` at once
and then calls the import slot EN `L2.01463` / RU `L2.01464` that
R2-ENGINE-073 identifies as DLL TalkTo, with the same packed value. No
test follows the builder call. The conversation therefore reaches TalkTo
with an empty body when `npc517talk10` is missing, and the kind-3 unlock
of R2-ENGINE-073 does not depend on the section.

EN and RU agree: the 124-instruction loader, the three-instruction
builder branch and the 31-instruction TALK tail are identical after
address normalisation (0 mismatches of 158). The comparison is over
normalised instruction text; callees are not compared, so it states shape,
not the callee each call reaches.

**Confidence.** EN High for the loader's not-found stores (`L2.01465` to
`L2.01466`) and the builder branch (`L2.01467`): both were read in hand
windows 1 and 4 within the preregistered budget of 12. Medium for: the find
returning -1 when the section is absent (unread `L2.01457`); the builder
arm, the constructor and the loader call (listings not committed); the
TALK tail (`L2.01468` to `L2.01469`), whose span was first read as a
generator window past the budget (the EXP-2011 listing also holds those
instructions); the EN/RU shape identity (the RU spans were read in hand
windows 9 and 10, the TALK tail only by the generator); and the whole RU
half, because for RU -1 means absent only if `L2.01289` returns it. Medium
for the displayed result: the page code after `L2.01460` (the target of
the parser's NUL exit, not read), the constructor base and the page
handlers were not read, so an empty box versus an immediate close is not
decided.

**Unknown.** Each with its settling step: what the page draws or does with
an empty body (read the page code after `L2.01460` and the constructor
base `L2.01470`); whether a missing `town.txt` and an empty buffer reach
the same arm (read `L2.01457` and `L2.01289` on an empty string); sections
whose header exists but whose body is empty (read the cut and parser on a
zero-length body); the RU find result for absent (read `L2.01289`).

### R2-ENGINE-344

Every `#<key>` lookup the census saw runs in the one loader of
R2-ENGINE-343, called once from the dialogue constructor EN `R2.0040`,
which the builder EN `R2.0038` / RU `R2.0037` calls once. That constructor
and loader chain is EN only and from the lane notes (not in `evidence/`).
A raw census of `call rel32` and of every dword offset in every section
(`callers.tsv`, `refs.tsv`) finds 21 builder call sites in 10 functions in
each image and no dword reference to the builder or TALK; constructor and
loader callers and dword scans were run for EN only, in the notes. The
formats pushed near the sites are
`event%d`, `npc%dabout`, `quest%d`, five `treasure*` keys, `npc%daccept%d`,
`npc%dreject%d`, `npc%dtalk%d`, `shop\npc31m%d`, `plagatguard`,
`druidinnkeeper%d`, `druidshopkeeper%d`, `kaargguard%d`, `kaargwoman%d`
and `kaargman%d`. `npc%dtalk%d` has one code reference per image, in TALK,
which has two direct call sites per image. The bytes `npc517` and `talk10`
occur in neither `allods2.exe` nor `Scenario.dll` of either root
(`literals.tsv`).

The buffer is a global string filled by the text loader EN `R2.0056` /
RU `R2.0055` (R2-ENGINE-052). The EN code passes the buffer address to it
for `globalmap.txt` (`L2.01471`), `town.txt` (three sites: `L2.01472`,
`L2.01473`, `L2.01474`), `quest.txt` (`L2.01475`) and `mission%d.txt`
(`L2.01476`). `help.txt` goes to another destination (`L2.01477`,
`refs.tsv`); the destination of `Docs\%d.txt` is Unknown (its pushed
operand is blank in `refs.tsv`). The lookup of a key therefore searches whichever of those files was loaded
last. The RU census finds the same path strings with the pushed buffer
`L2.01456` for `globalmap.txt`, `town.txt` x3 and `quest.txt`; its
`mission%d.txt` site was not decoded.

No site of the census belongs to a save-load routine by its format string,
and the first-town section name is built only by TALK. Loading a first-town
save therefore reads `npc517talk10` only if the load path reaches TALK;
no such route was found.

**Confidence.** Medium. The census is a raw `call rel32` and dword-offset
scan of `.text` and all sections; it misses a call through a register
holding a computed address. The save-load routines themselves were not
read, and the town screen's reload of `town.txt` on every entry was not
established (the three loader sites are inside screen message handlers).

**Unknown.** Each with its settling step: which routine reaches the two
TALK call sites (find the callers of `L2.00743` and `L2.00755`, including
register calls); whether any load path posts the town screen messages that
call TALK (read the save-load routines); the RU `mission%d.txt`
destination (decode that RU site); the 14 dword references to the EN
buffer between `L2.01478` and `L2.01479` and the two in the text loader
(`L2.01480`, `L2.01481`), which are not classified; the `Docs\%d.txt`
destination; which file the buffer holds when TALK runs.
