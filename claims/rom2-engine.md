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
| R2-ENGINE-050 | In both ROM2 clients, UI message `0x433` selects `event%d` for ordinary IDs or defers the ID under UI bit eight; IDs 250, 253, 254 and 255 take separate branches. | Medium | ✔ promoted (branch candidate) | [EXP-2009](../experiments/EXP-2009-rom2-text/) |
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
| R2-ENGINE-147 | The add helpers scan the catalog for matching type and ID, append the record pointer to the available list only when no node already holds that pointer, and the remove and clear helpers unlink or empty that doubly linked list. | High | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-148 | The departure output DWORD is -1 by default, 1 for ID 10, 2 for ID 30, 3 for IDs 70 and 80 only under opposite bank777/bank778 gates (the gate skips only the store), 5 or 4 for ID 110 by bank779, unchanged otherwise and for type 2. | High | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-149 | A nonzero bank775 makes ordinary departure clear the available list and append 42 fixed record pointers plus one of four bank776/bank781-selected records before the normal path; the types and IDs behind those pointers are not resolved here. | Medium | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-150 | In the 23-routine EN/RU population no direct callee is unresolved except two unread excluded callees of the node shells, and no indirect edge exists in the controller or list helpers. | Medium | ● active | [EXP-2018](../experiments/EXP-2018-rom2-campaign-departure/) |
| R2-ENGINE-159 | Catalog record kind 2 / ID 2 is built by one constructor, D2.00016, calling the record initializer D2.00017 on object D2.00018; it is one of 52 constant-argument initializer call sites in EN and RU scenario.dll. | Medium | ● active | [EXP-2020](../experiments/EXP-2020-rom2-second-town/) |
| R2-ENGINE-160 | The client body that selects per-ID town data has a separate arm for ID 2 beside arms for IDs 1 and 3; the ID 2 arm pushes the resource key music\b16.wav in EN (arm L2.00264) and RU (arm L2.00265). | Medium | ● active | [EXP-2020](../experiments/EXP-2020-rom2-second-town/) |
| R2-ENGINE-161 | The EnterInn export has 10 stage case bodies (8 distinct) and a default; stage 20 and other unlisted stages select the shared continuation D2.00019 (295 instructions), which holds 15 packed-entry stores behind bank-slot compares. | Medium | ● active | [EXP-2020](../experiments/EXP-2020-rom2-second-town/) |

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
| R2-ENGINE-191 | The selected RU Player list calls Group inline; Group invokes two unresolved embedded programmes before its member path and three final scalar calls. | Medium | ✔ promoted | [EXP-2024](../experiments/EXP-2024-rom2-group-serialization/) |
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
