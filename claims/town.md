# Claim registry — TOWN (the between-missions surfaces: the town, the world map, and the rooms reached from them)

Level 2 ledger. Index: [registry.md](registry.md). Legend: `✔ promoted` (also in the spec) · `● active` · `✖ retracted`. IDs are permanent.

This ledger covers what a player sees and can click **between** campaign missions: the town
surface itself, the world map, and the rooms a click on either one opens. Two experiments
published into it from disjoint id blocks and disjoint draw paths.

`TOWN-001`..`TOWN-026` are `EXP-0183`, over the town surface and the two rooms it was required to
cover, the tavern where mercenaries are hired and the school where skills are trained. The shop is
covered structurally by `claims/shop.md`'s `SHOP-SCREEN-*` and `SHOP-TRAY-*` rows and is
cross-referenced here rather than re-published.

`TOWN-036`..`TOWN-045` are `EXP-0184`, over the world-map surface, amended where stated by
`EXP-0188`. `TOWN-116`..`TOWN-124` close its route, input, paint, and completion path.
`TOWN-141`..`TOWN-145` are `EXP-0194`, which re-derived the route criterion `TOWN-119` states and
amended that row's wording. The rows cite
a routine in the same code region and never the town's own class range. The two experiments share this
ledger because they share the campaign-state consumers `REG-SCN-059`, `REG-SCN-062`, `REG-SCN-065`
and `REG-GMAP-066` in `claims/reg.md`.

Ids here are written plain in the first cell, without a topic segment. Both check scripts read that
shape as of 2026-08-16; before that date neither could see a row in this file. **The check scripts
were not the only readers, and scoping this note to them was itself the defect.** `tools/claim`,
which both `AGENTS.md` files route every single-claim lookup through, carried the same three-segment
assumption and answered `TOWN-001 is not in any ledger` for every row here until `ac7d03a` the same
day — a present, correct row reported as a missing claim, which `AGENTS.md` defines as a defect in
the brief that asked for it. `check-claim-ids.sh` now resolves one id from every ledger through that
tool, so the two cannot diverge again without a gate failing.

## Id allocation notes

`TOWN-027`..`TOWN-035` were allocated to `EXP-0183` and `TOWN-046`..`TOWN-060` to `EXP-0184`. All 24
are returned unused and none is ever reissued. `EXP-0185` was allocated `TOWN-061`..`TOWN-085` and
spent `TOWN-061`..`TOWN-068`; `TOWN-069`..`TOWN-085` (17 ids) are returned unused, none ever
reissued. `EXP-0186` was allocated `TOWN-086`..`TOWN-105` and spent `TOWN-086`..`TOWN-096`, 11 ids;
`TOWN-097`..`TOWN-105` (9 ids) are **returned unused**, none ever reissued. `EXP-0187` spent
`TOWN-GENERAL-106`..`TOWN-GENERAL-108` and permanently returned
`TOWN-GENERAL-109`..`TOWN-GENERAL-115`. `EXP-0188` spent `TOWN-116`..`TOWN-124` and permanently
returned `TOWN-125`..`TOWN-135`. `EXP-0193` was allocated `TOWN-136`..`TOWN-140` and spent
all five; `EXP-0194` was allocated `TOWN-141`..`TOWN-145` and spent all five. The two ran
concurrently and their ranges were allocated by the orchestrating seat rather than derived by
either lane, because two lanes starting from the same master arrive at the same next free id.
`EXP-0195` was allocated `TOWN-146`..`TOWN-155` and spent all ten. `EXP-0196` was allocated
`TOWN-156`..`TOWN-180` and spent `TOWN-156`..`TOWN-165`, 10 ids; `TOWN-166`..`TOWN-180` (15 ids)
are **returned unused**, none ever reissued. `EXP-0197` was allocated `TOWN-181`..`TOWN-205` and
spent `TOWN-181`..`TOWN-188`, 8 ids; `TOWN-189`..`TOWN-205` (17 ids) are **returned unused**, none
ever reissued. **Both experiments opened on 2026-08-19 were first allocated ranges inside that
already-returned range, and neither used one**: `EXP-0198` was given `TOWN-189`..`TOWN-196` and
`EXP-0199` was given `TOWN-197`..`TOWN-204`. The seat issued corrected, disjoint ranges.
`EXP-0198` was allocated `TOWN-206`..`TOWN-213` and spent `TOWN-206`..`TOWN-211`, 6 ids;
`TOWN-212`..`TOWN-213` (2 ids) are **returned unused**. `EXP-0199` was allocated
`TOWN-214`..`TOWN-221` and spent `TOWN-214`..`TOWN-217`, 4 ids; `TOWN-218`..`TOWN-221` (4 ids) are
**returned unused**. `EXP-0200` was allocated `TOWN-222`..`TOWN-231` and spent `TOWN-222`..`TOWN-224`,
3 ids; `TOWN-225`..`TOWN-231` (7 ids) are **returned unused**. `EXP-0201` was allocated
`TOWN-232`..`TOWN-241` and spent `TOWN-232`..`TOWN-236`, 5 ids; `TOWN-237`..`TOWN-241` (5 ids) are
**returned unused**. None is ever reissued. The next free `town.md` id is therefore `TOWN-242`.

Two statements in this paragraph were corrected at the merge, 2026-08-19. `EXP-0199`'s branch wrote
its own first range as `TOWN-197`..`TOWN-205` here while its own `claims/registry.md` paragraph
wrote `TOWN-197`..`TOWN-204`; the register is the authority and the endpoint is `TOWN-204`, and no
id in either range is spent, so nothing downstream moves. `EXP-0198`'s branch declared `TOWN-214`
free, which was true on its own branch and false after `EXP-0199` landed; the merge carries
`TOWN-222`.

`EXP-0202` was allocated `TOWN-242`..`TOWN-251` and spent `TOWN-242`..`TOWN-246`, 5 ids;
`TOWN-247`..`TOWN-251` (5 ids) are **returned unused**.

`EXP-0203` was allocated `TOWN-252`..`TOWN-263` and spent `TOWN-252`..`TOWN-261`, 10 ids;
`TOWN-262`..`TOWN-263` (2 ids) are **returned unused**.

`EXP-0204` was allocated `TOWN-264`..`TOWN-279` and spent `TOWN-264`..`TOWN-266`, 3 ids;
`TOWN-267`..`TOWN-279` (13 ids) are **returned unused**. The next free `town.md` id was therefore
`TOWN-280`, the top of that range plus one. It read `TOWN-267` until 2026-08-20, which is the
first id the same sentence returns; `claims/registry.md` carries the correction note.

`EXP-0205` was allocated `TOWN-280`..`TOWN-295` and spent `TOWN-280`..`TOWN-284`, 5 ids;
`TOWN-285`..`TOWN-295` (11 ids) are **returned unused**. `TOWN-296`..`TOWN-311` (16 ids, reserved
to `EXP-0206`) are also **returned unused**, none reissued. `TOWN-312`..`TOWN-314` (3 ids) are
spent below (`EXP-0207`); `TOWN-315`..`TOWN-327` (13 ids, reserved to the same experiment) are
**returned unused**. `EXP-0208` was allocated `TOWN-328`..`TOWN-343` and spent
`TOWN-328`..`TOWN-333`, 6 ids; `TOWN-334`..`TOWN-343` (10 ids) are **returned unused**, none ever
reissued. `EXP-0210` was allocated `TOWN-344`..`TOWN-355` and spent all 12; none is returned.
`EXP-0212` was allocated `TOWN-356`..`TOWN-371` (16 ids) and spent `TOWN-356`..`TOWN-360`, 5 ids;
`TOWN-361`..`TOWN-371` (11 ids) are **returned unused**, never reissued.
The next free `town.md` id is therefore `TOWN-372`.

No returned id is reissued. `claims/registry.md`'s allocation section is the authority; this
paragraph is not.

`EXP-0214` was allocated `TOWN-372`..`TOWN-378` (7 ids) and spent `TOWN-372`..`TOWN-373`,
2 ids. **`TOWN-374`..`TOWN-378` (5 ids) are returned unused**, none ever reissued. The next
free `town.md` id is therefore `TOWN-379`.

## Tavern roster geometry, halves and number grouping

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-467 | The tavern roster places list position `i` in rect `i` of an 18-rect table filled bottom row first: column `i%6`, row `⌊i/6⌋` counted up from the bottom. | High | ● active | [EXP-0400](../experiments/EXP-0400-tavern-roster-presentation/EXP-0400.md), `evidence/listing.txt`, `evidence/roster-grid.csv`; supersedes TOWN-065 (retracted), restores TOWN-010's formula |
| TOWN-468 | Mercenary cells come first and talk-only cells after them; the two halves differ in background, text and source. | High | ● active | [EXP-0400](../experiments/EXP-0400-tavern-roster-presentation/EXP-0400.md), `evidence/listing.txt`; corrects TOWN-010, REG-118, REG-SCN-064 |
| TOWN-469 | Nineteen call sites in nine routines group a decimal string with `R0807`, and the rule keeps a leading sign in front with no comma after it. | High / Medium | ● active | [EXP-0400](../experiments/EXP-0400-tavern-roster-presentation/EXP-0400.md), `evidence/grouping-callers.txt`, `evidence/grouping.txt`, `evidence/listing.txt`; extends SHOP-052; answers FAME-DISPLAY-012's downstream-transformation Unknown for the score (`L11474`) |
| TOWN-470 | `graphics.res` ships 28 inn sprite sheets on each root, and the `Unit<n>` sheets above 15 are talk-candidate sheets keyed by npc id. | High / Medium | ● active | [EXP-0400](../experiments/EXP-0400-tavern-roster-presentation/EXP-0400.md), `evidence/inn-sheets.csv`, `evidence/talk-candidates.csv`; narrows TOWN-012 |

### TOWN-467

`R1903`, called at the end of the roster child's constructor, fills 18 `RECT`s from `this+0x60` in 16-byte steps: rows at `top = this.bottom − 0x40·(r+1)` (`L11475` loads the constant 0xffffffc0, `L11476` subtracts 0x40 from it and `L11477` compares the result with 0xffffff00), six 0x30x0x40 cells from `this.left + 0x10` (`L11478`, `L11479`, `L11480`). Both paint loops and the hit test address the rect with a multiplication by the constant 0x2aaaaaab, a sign fix and a division by 6, an address formed from a base plus twice `3·(q+1)`, and a left shift of that base by 4 (`L11481`..`L11482`, `L11483`..`L11484`, `L11485`..`L11486`). `0x2AAAAAAB` is ⌈2^32/6⌉, so the quotient is ⌊i/6⌋ (for `i = 3` the product is `0x80000001`, high dword 0) and the offset is `(i%6 + (⌊i/6⌋+1)·6)·0x10 = 0x60 + 0x10·i` — `TOWN-010`'s formula, not `TOWN-065`'s.

With the child at `(160,0)-(480,480)` (`TOWN-282`) cell `i` is 48x64 at `x = 176 + 48·(i%6)`, `y = 480 − 64·(⌊i/6⌋+1)` in view coordinates. The shipped campaign offers at most 16 entries (13 types with a non-zero `MercenaryCount`, at most 3 `InnNPC`), inside the table

**Confidence.** High. Rival excluded: `TOWN-065`'s division by 3, which needs the multiplier `0x55555556`, absent from all three sites, and which would address rect 25 of 18 from position 9 on; the owner's original screenshot for a 14-entry roster shows two full bottom rows and a top row of two cells, which only the ÷6 mapping produces

### TOWN-468

`R1412` fills `view+0xc0` (data `+0xc4`, count `+0xc8`) from `R0746`'s `CUnit` list (`L11487`..`L10174`) and `view+0xe8` (data `+0xec`, count `+0xf0`) with `R0741(InnNPC[j])` for each `InnNPC` element (`L08032`..`L09157`). Paint loop 1 (`i < view+0xc8`, `L11488`) draws `manback.bmp`, the `Unit<+0x15b>` sheet, the price string after `R0807` at the cell's top (`L11489`) and `"%d/%d"` working/pristine pool at its bottom (`L11490`), plus a glyph when the type's hire flag is set (`L11491`). Paint loop 2 (`j < view+0xf0`, `L11492`) draws `ManBackTalk.bmp` and the talk object's sheet at list position `view+0xc8 + j` (`L11493`), with no text.

Both loops skip a cell's whole draw while the selection `view+0xb8` is −1 (`L11494`/`L11495`, `L11496`/`L11497`); activation stores 0 there for a non-zero mercenary count and −1 for zero (`L08095`..`L11498`), so a roster with no mercenary draws no talk cell until a selection is stored. The hit test covers `view+0xc8 + view+0xf0` positions and stores the hit in `view+0xb8` (`R1807`, `L11499`). So `view+0xc8` is the mercenary count: the bound `REG-118` names `InnNPC.Size()` is this count, and `view+0xc4` is the mercenary `CUnit*` array, not `InnNPC.m_pData`

**Confidence.** High. Rival excluded: `TOWN-010`'s talk-first order — the first loop reads the `CUnit` type byte and the pool-and-price strings and is bounded by the list `R0746` built, and in the owner's screenshot the 13 priced cells occupy positions 0..12 and the green cell position 13

### TOWN-469

A raw scan of every PE section finds 19 `E8` calls and no absolute dword naming `R0807`: character generator final stage remaining points `+0x1f0` (`L11500`, `R0833`, `HERO-BUY-003`); character generator stat text `"%+d"`, the negated raise cost (`R0828`, the negation at `L11501`) or the lowering refund (`R0829`) unchanged, so a raise reads negative and a lowering positive (`L11502`, `R1211`; the buy handler subtracts the cost at `L11503`, the sell handler adds the refund at `L11504`); hall of fame record score `+4` (`L11474`, `R0806`, `FAME-DISPLAY-012`); tavern button panel numbers `+0x124`/`+0x12c` (`L11505`, `L11506`, `R1904`, `TOWN-392`); tavern mercenary price (`L11507`); the money cell of the item-grid class at vptr `L06494`, player `+0xc` (`L11508`, `R1772`, screen not identified); shop stock-grid quantity and price (`L11509`, `L11510`, `R0994`, `SHOP-SCREEN-039`); shop button panel values `+0x150`..`+0x15c`, two arms each (`L11511`..`L11512`, `R1027`, `SHOP-SCREEN-035`); school widget values (`L11513`, `L11514`, `R1905`).

The routine prefixes `','` + the last three characters while more than three remain, and stops at four when the first character is `-` or `+` (`L11210`..`L11515`, `L11516`, `L11216`): `-1000` → `-1,000`, `-999` → `-999`, `+1000` → `+1,000`. Drawn without it on these screens: the tavern `"%d/%d"` count, the hall of fame `"%d."` rank and the generator's `"%s = %d"` text. No call lies in the character/unit panel `R0877`

**Confidence.** High for the census and the rule. Rival excluded: an indirect or table caller, by the absolute-dword half of the scan (0 hits). **Medium** that no second grouping routine exists: the image has no grouping format literal, and a hand-written inserter elsewhere was not searched

### TOWN-470

The nodes under `interface\inn\` ending in `sprites.16a` are `HeroFighter`, `HeroMage`, `Unit1`..`Unit15`, `Unit29`, `Unit30`, `Unit32`, `Unit41`..`Unit44`, `Unit52`, `Unit61`, `Unit62`, `Unit64`, identical on EN and RU. The talk-cell loader draws `Unit<+0x15b>` for a non-Hero object and the dialogue synthesiser stores the npc id there (`TAVERN-TALKPIC-016`); the ids 30, 32, 41, 52, 62 and 64 are `InnNPC` values of the shipped campaign, and 43 and 44 name RU-only inn texts (`REG-INNCLOSE-116`)

**Confidence.** High for the census (every node of both archives listed). **Medium** that 29, 42, 43, 44 and 61 are talk sheets too: no shipped `InnNPC` names them

## Exterior horse, baba and dervish lifecycles

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-439 | The town loader binds one selected position from each complete horse, baba and dervish family to separate view fields. | High | ✔ promoted | [EXP-0304](../experiments/EXP-0304-horse-baba-dervish/), `assets.tsv`, `tables-{en,ru}.tsv`, loader anchors and bounded raw trace |
| TOWN-440 | Baba is selected on entry but independently reselected by elapsed time in every eligible town paint. | High | ✔ promoted | [EXP-0304](../experiments/EXP-0304-horse-baba-dervish/), `anchors.tsv`, `state-transitions.tsv`, `model-vectors.tsv` |
| TOWN-441 | Baba advances forward through the selected sheet and clears at that sheet's own terminal, then waits for another delay. | High | ✔ promoted | [EXP-0304](../experiments/EXP-0304-horse-baba-dervish/), hub/helper bodies, `anchors.tsv`, `state-transitions.tsv`, terminal vectors |
| TOWN-442 | Horse has its own delayed paint-time selector rather than sharing baba state. | High / Medium | ✔ promoted | [EXP-0304](../experiments/EXP-0304-horse-baba-dervish/), entry/reroll/painter bodies, `anchors.tsv`, `state-transitions.tsv` |
| TOWN-443 | Horse is a single forward 15-frame episode with terminal clear and retained frame0, not a loop or ping-pong. | High | ✔ promoted | [EXP-0304](../experiments/EXP-0304-horse-baba-dervish/), `assets.tsv`, helper disassembly, terminal vectors |
| TOWN-444 | Dervish is entry-armed and continuously wraps, unlike both delayed families. | High | ✔ promoted | [EXP-0304](../experiments/EXP-0304-horse-baba-dervish/), dervish entry/hub/helper bodies, `assets.tsv`, `model-vectors.tsv` |
| TOWN-445 | Horse sound requests are variant/frame gates, and loaded `Horse3.wav` has no direct request receiver in the bounded town closure. | High / Medium | ✔ promoted | [EXP-0304](../experiments/EXP-0304-horse-baba-dervish/), `assets.tsv`, `references.tsv`, sound/painter anchors and vectors; request contract `VIDEO-SFX-013` |
| TOWN-446 | Leave, destruction and re-entry release and rebuild all three object-local family episodes. | High / Medium | ✔ promoted | [EXP-0304](../experiments/EXP-0304-horse-baba-dervish/), lifecycle bodies, `references.tsv`, `anchors.tsv`, state table |
| TOWN-447 | The separate-delay/entry-arm model discriminates the complete measured direct population; the town mask selector does not arm these high bits. | High / Medium / Unknown | ✔ promoted | [EXP-0304](../experiments/EXP-0304-horse-baba-dervish/), `analysis-population.tsv`, `references.tsv`, `raw-manifest.tsv`, `model-vectors.tsv` |

### TOWN-439

`R1906` is the sole exact literal owner of `TownBirds/HORSE%d/A%d`, `BABA%d/A%d` and `DERVISH%d` formats. It selects horse position1..5, baba position1..4 and a dervish position1..4 different from the baba choice, then loads three horse A sheets into array `+128`, two baba sheets into `+fc`, and one dervish sheet into `+150`; position fields are `+138/+13c`, `+10c/+110` and `+154/+158`. The complete installed population per root is 15 horse sheets of 15 frames, eight baba sheets with A1=31/A2=32 frames, and four dervish sheets of 30 frames. All selected payload hashes agree across roots. This refines `TOWN-004`'s position selection with the loaded family/field bindings; a loaded sheet is not by itself a draw or cadence result.

**Confidence.** High for the unique literals, byte-checked bindings/positions and complete two-root population; computed names and load failure remain Unknown

### TOWN-440

Entry selects array0/A1 into `+f4`, writes current `+114=-1`, clock `+118=now`, a 2000..3999-ms delay at `+11c`, and leaves bit200 clear. Painter helper `R1907` tests strict unsigned `now-clock > delay`; expiry writes current0, rerolls delay2000..6999, chooses A1/A2 into `+f4` and sets bit200. The own painter always draws the selected pointer and clamps current -1 to frame0. Thus entry-only selection and paint-time reselection are both real stages, while the idle/terminal picture is the selected variant's frame0. This refines `TOWN-159`'s `+0xf4` reroll mechanism with the field's family identity, entry state and draw-clamp relation; `TOWN-159`'s undetermined roster identity for `+0xf4` is baba.

**Confidence.** High for local entry, strict arm, random-domain, selected-pointer and draw relations; actual paint delivery and visible result are Unknown

### TOWN-441

The admitted `R1908` hub tests bit200 and calls `R1909`, which increments current once and refreshes the baba clock. Reaching selected count31 or 32 writes current -1 and clears bit200 without releasing or replacing `+f4`; painter clamping therefore holds frame0 until a later strict delay rerolls/rearms. There is no wrap, reverse, immediate terminal rearm or catch-up loop in this path. This refines `TOWN-159`'s bit `0x200`/`+0xf4` dispatch with the per-sheet terminal count and clear/retention writes.

**Confidence.** High for hub bit/helper, frame-count comparison and clear/retention instructions; delivered cadence and scheduling outside the local hub are Unknown

### TOWN-442

Entry selects A1 into `+120`, writes current `+140=-1`, clock `+148=now`, delay `+14c=2000..3999` and leaves bit100 clear. On its separate strict elapsed test, `R1907` writes current0, rerolls 2000..6999 ms, chooses selector `+144=0..2` and corresponding A1/A2/A3 pointer, then sets bit100. Painter always draws the selected pointer and clamps -1 to frame0. Entry does not explicitly reset selector `+144`, but it is dormant while current is -1 and every arm overwrites it before the frame-specific sound tests can match. This refines `TOWN-159`'s `+0x120` reroll and its `+0x144`/`+0x140` sound-dispatch keys with the field's family identity and selector-dormancy behaviour; `TOWN-159`'s undetermined roster identity for `+0x120` is horse.

**Confidence.** High for separate local fields, strict arming and pointer/selector draw relation; Medium for the bounded dormant-retention interpretation; runtime paint and load failure remain Unknown

### TOWN-443

The admitted hub tests bit100 and calls `R1910`, which increments current once and refreshes the horse clock. Every installed horse sheet has 15 frames; reaching count15 writes current -1 and clears bit100 while retaining `+120`, so the next and subsequent idle paints use frame0 until another independent delay. There is no same-call rearm or catch-up loop. This refines `TOWN-159`'s bit `0x100`/`+0x120` dispatch with the all-sheet frame-count terminal and clear/retention writes.

**Confidence.** High for the bounded helper, all-sheet counts and terminal writes; actual frame delivery and visible repetition are Unknown

### TOWN-444

Entry writes current `+15c=0` and bit400; the admitted hub tests bit400 and calls `R1911`. With non-null selected pointer `+150`, each call writes `(current+1)%frameCount`; all four installed sheets have 30 frames, so 29 wraps to 0. No delay, random reselection, terminal hold or bit400 clear is present. A null pointer writes current -1 but leaves bit400 set. The painter draws the pointer/current unconditionally when present. This refines `TOWN-159`'s third, unrandomised `+0x150`/`[+0x15c]` field with its family identity, entry arm and modulo-wrap progression; `TOWN-159`'s undetermined roster identity for `+0x150` is dervish.

**Confidence.** High for entry arm, modulo helper, absent local clear and complete sheet counts; runtime hub delivery and failed-load behavior are Unknown

### TOWN-445

Sound loader `L11517` binds Horse2 to `+a0`, Horse3 to `+a4` and Horse1 to `+a8`. The painter conditionally requests Horse1 at A3 frame1; Horse2 at A1 frame14, A2 frames8/14 and A3 frame14. Each requires a non-null object and status0, then calls the accepted request helper with repeat0 and priority128. The complete town-range direct displacement population for `+a4` contains only construction zero, load and cleanup; no baba/dervish-specific sound literal or direct request appears. The complete case-insensitive stem population also contains two unrelated `units/horse/*` sounds, which are classified rather than mistaken for town events. Computed aliases, indirect/out-of-range requests and audibility remain outside this negative.

**Confidence.** High for literal-to-field bindings and named conditional request gates; Medium for the bounded Horse3/baba/dervish negative; runtime acceptance, mixing and audibility Unknown

### TOWN-446

Leave `R1912` clears the active-view gate and calls asset cleanup `L11518` plus sound cleanup `R1913`; destructor `L11519` calls the same cleanups. Asset cleanup releases arrays and zeros selected pointers `+f4/+120/+150`; sound cleanup detaches/releases/zeros Horse1/2/3 objects. Re-entry first cleans/reloads, selects positions/assets, restores baba and horse A1/current -1/new delays, and arms dervish at current0. The families share view flag word `+208` and process-static admitted-hub clock `L11520`, but no family-specific process-static state was reached. Object reuse and alias/bulk writers remain open.

**Confidence.** High for named call paths, releases, local resets and shared direct references; Medium for the bounded no-other-static interpretation

### TOWN-447

Exhaustive byte inputs80h..c0h to `R1914` yield only -1,1,2,4,8,16, so its default pointer OR cannot directly set bits100/200/400. Direct exact literal/call sweeps and town-range displacement sweeps connect horse/baba paint-delay arms and dervish entry arm to separate helpers and find no input arm, reverse helper or second direct consumer. The measured range has 39 functions, 12,915 function bytes, zero orphan bytes and 462 undisassembled bytes; 20 distinct bodies and 14 complete disassembly representations were inspected. This bounds, but cannot exclude, computed/indirect targets, overlap writers, the undisassembled remainder or out-of-range consumers.

**Confidence.** High for enumerated direct edges, exhaustive selector outputs and discriminating local state; Medium for exclusivity; Unknown outside the stated population

## School training presentation contracts

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-427 | All four `movies/training/{mage,fighter}/{m,tr}` families are loaded into the school-room object. | High | ✔ promoted | [EXP-0303](../experiments/EXP-0303-school-training-consumer/), `anchors.tsv`, `assets.tsv`, `asset-members.tsv`, `relations.tsv`, bounded raw trace |
| TOWN-428 | The four loaded families are school-room presentation, not an independent movie surface. | High | ✔ promoted | [EXP-0303](../experiments/EXP-0303-school-training-consumer/), school painter body, `anchors.tsv`, `relations.tsv`; helper contract `TOWN-407` |
| TOWN-429 | Each `tr` family is a bounded class-transition cycle coupled to the rotating column. | High | ✔ promoted | [EXP-0303](../experiments/EXP-0303-school-training-consumer/), entry/picker/painter/updater bodies, `anchors.tsv`, `model-vectors.tsv`; class/column contracts `TOWN-147`, `TOWN-149` |
| TOWN-430 | The two `m` episodes are independently idle-armed and share the `tr` paint gate, rather than advancing on input. | High | ✔ promoted | [EXP-0303](../experiments/EXP-0303-school-training-consumer/), armer/painter bodies, `anchors.tsv`, `model-vectors.tsv`; accepted PRNG bound from `TOWN-409` |
| TOWN-431 | Mage `m` advances forward, holds its terminal cached picture, then deliberately resumes reverse below the top. | High | ✔ promoted | [EXP-0303](../experiments/EXP-0303-school-training-consumer/), `L11521`, accepted forward/reverse helpers, scaler-span `anchors.tsv`, `model-vectors.tsv` |
| TOWN-432 | Fighter `m` uses the same forward/hold/reverse shape without the mage index skip. | High | ✔ promoted | [EXP-0303](../experiments/EXP-0303-school-training-consumer/), `L11522`, accepted forward/reverse helpers, scaler-span `anchors.tsv`, `model-vectors.tsv` |
| TOWN-433 | School entry and leave reset/release object-local training state, while the idle clocks are separate static state. | High / Medium | ✔ promoted | [EXP-0303](../experiments/EXP-0303-school-training-consumer/), entry/leave/destructor/cleanup bodies, `relations.tsv`, bounded reference trace |
| TOWN-434 | The separate-movie, data-only and mixed-family ownership models are refuted inside the measured literal/direct surface; computed aliases remain open. | High / Medium / Unknown | ✔ promoted | [EXP-0303](../experiments/EXP-0303-school-training-consumer/), string xrefs, call census, 41-body manifest, `anchors.tsv`, `relations.tsv`; audio boundary `VIDEO-SFX-013` |

### TOWN-427

`R1486`, the established school loader, is the sole direct caller of family loaders `L07743` and `L07744`; neither target has a parsed-`.rdata` code pointer. The mage loader requests `tr0000..tr0022` into `room+e4/+e8` and `m0001..m0011` into sequence `+164`; the fighter loader requests `tr0000..tr0018` into `+198/+19c` and `m0001..m0009` into sequence `+218`. Both `movies.res` roots contain exactly those four gapless groups: 62 BMPs per root, 124 measured instances, with identical hashes across roots. This resolves the four-family part of `TOWN-022` and retracts `TOWN-215`'s callee-blind source negative; the three `npc*` key formats in `TOWN-022` remain open.

**Confidence.** High for the literal-reference, direct-call, field-binding and complete four-folder corpus; computed names/targets remain outside the negative boundary

### TOWN-428

School own-painter `R1487` places the mage side at room-relative `(0,200)` and fighter side at `(320,200)`. Mage flag bit2 selects cached sequence `+164`; otherwise current `tr` surface `+f4` is blitted at the same anchor. Fighter bit8 similarly selects sequence `+218`; otherwise `+1a8` is blitted there. Both sequence draws call accepted helper `L11523`, while the `tr` arms invoke the picture's draw slot directly. Thus `m` replaces `tr` on its side while active; it is not an overlay or separate cutscene frame. Loaded-pointer success remains a precondition.

**Confidence.** High for direct fields, mutually exclusive branches, draw primitives and coordinates in the hash-identical executable; runtime visibility remains Unknown

### TOWN-429

Entry resets both counters to -1, primes each current pointer at index0, and arms the selected class's `tr` bit. A changed class sets transition-pending state; after the local delay the matching bit is armed and school input remains blocked until that `tr` counter returns to 0 and the column reaches the matching endpoint. The shared unsigned strict `>83`-ms paint gate admits at most one local step: mage advances modulo23 and clears bit1 on wrap0; fighter advances modulo19 and clears bit4 on wrap0. Mage index at least5 can start column step+1 when the mage class is selected; fighter index at least6 can start step-1 for fighter. Those gated starts pass `SFX\Town\School\Rotate.wav` to the school request helper.

**Confidence.** High for local state, modulo bounds, wait release, column and request conditions; delivered paint/input cadence and audible output are Unknown

### TOWN-430

Mage armer rejects side flags1|2 and fighter armer rejects4|8; while either side is busy it refreshes that side's timestamp. When idle, unsigned elapsed must be strictly greater than `3000 + rand()/10`; with the accepted 15-bit PRNG range this is 3000..6276 ms. Arming sets direction1, sequence index -1 and side bit2 or 8. The same strict `>83`-ms block advances at most one `m` or `tr` step per side and has no catch-up loop. The draw follows the update, but additional rejected paints can repeat the cached frame.

**Confidence.** High for local masks, arithmetic, strict comparisons and single-step order; actual frame rate and distribution of visible episode spacing are Unknown

### TOWN-431

With all eleven loaded pointers, entry/arm first presents sequence index0 (`m0001`), ascent reaches index10 (`m0011`), and the next admitted terminal call changes direction to -1 and samples hold threshold `((rand()*20)/0x7fff)%20 + 20`, range20..39. This scaler is not simplified to `20+rand()%20`, and no uniform distribution is asserted. The updater then writes sequence index8 without changing cached index10. During hold, `m0011` remains selected; threshold admission calls reverse from 8 to 7, so return presents indices7..0 (`m0008` down to `m0001`) and never presents indices9 or 8 on that return. Reverse terminal at 0 returns zero and paint clears bit2 before the `m` draw branch. Active `tr` forces the hold counter high and an ascending episode into reverse, but does not clear it instantly.

**Confidence.** High for conditional local scaler/index/cache/direction writes and derived all-pointers vector; runtime episode duration, missing-load behavior, distribution and visible interruption remain Unknown

### TOWN-432

With nine loaded pointers it presents indices0..8 (`m0001..m0009`), samples the same `((rand()*20)/0x7fff)%20 + 20` threshold, range20..39 with no uniform-distribution claim, at terminal index8, then reverses 7..0. It does not rewrite the sequence index at terminal. Reverse completion clears bit8 before the `m` draw branch. Active fighter `tr` similarly forces an ascending episode toward reverse. This distinguishes the two `m` families despite their shared helper and clocks.

**Confidence.** High for conditional local scaler/state and the all-pointers vector; actual cadence, distribution, load failure and visible interruption remain Unknown

### TOWN-433

Entry invokes both family loaders, zeros local flags/states, resets the `tr` counters through -1 to index0 and rebuilds both `m` sequences with index/cache0. Leave `R1915` and destructor `L11524` call both family cleanups; those release the `tr` arrays and sequence storage. Re-entry therefore reloads object fields and begins the selected-class transition anew. The armer-initialization bytes, timestamps, random extras and hold counters are globals and are not explicitly cleared by the read entry/leave/destructor bodies. Busy-side arming refreshes timestamps, so retained static values do not imply an immediate idle episode. Alias/overlap writers and object allocation identity remain open.

**Confidence.** High for named entry, cleanup and local resets; Medium for the bounded static-state persistence interpretation; object identity and cross-visit timing remain Unknown

### TOWN-434

Each of four format literals has one referencing family loader, each loader has one direct caller (`R1486`), no parsed-`.rdata` pointer to either loader exists, and all four field chains converge in school painter `R1487`. None of the `m` arm/step/draw bodies directly requests a sound; the reached animation-specific request is `Rotate.wav` in the two `tr` updaters under the column gates. Other school dialogue, music and SFX paths are separate and were not exhaustively attributed. This is not a whole-image proof against a computed filename/target or a guarantee of silence/audibility.

**Confidence.** High for the enumerated literal/direct-call/draw population; Medium for ownership exclusivity; Unknown for computed aliases, original visibility and audio result

## Exterior bird, statue-star and crowd lifecycles

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-415 | The exterior bird armer selects a prefix of one three-sprite group and ties its sound choice to selected count. | High / Unknown | ✔ promoted | [EXP-0302](../experiments/EXP-0302-town-bird-star-crowd/), `anchors.tsv`, `assets.tsv`, path bindings, armer/painter bodies and vectors |
| TOWN-416 | An active bird paint composes each nonterminal selected sprite before keyed `Town_add.bmp`, and clears bit80h only after every selected sprite is terminal. | High / Unknown | ✔ promoted | [EXP-0302](../experiments/EXP-0302-town-bird-star-crowd/), painter/hub ranges, anchors, 57-frame asset rows and terminal vectors |
| TOWN-417 | Town entry resets the effective per-view bird episode but not the three exact-address process statics on any direct measured path. | High / Medium | ✔ promoted | [EXP-0302](../experiments/EXP-0302-town-bird-star-crowd/), `references.tsv`, `reference-sites.tsv`, entry/armer/painter bodies and re-entry vector |
| TOWN-418 | The statue-star is a selector16-armed nine-picture episode with a separately stored paint pointer. | High / Medium | ✔ promoted | [EXP-0302](../experiments/EXP-0302-town-bird-star-crowd/), star anchors/ranges, mask/path tables, assets and vectors |
| TOWN-419 | The statue-star does not autonomously wrap: terminal calls retain a process-static every-ten rearm counter. | High / Medium | ✔ promoted | [EXP-0302](../experiments/EXP-0302-town-bird-star-crowd/), `anchors.tsv`, exact references, star body and terminal/re-entry vectors |
| TOWN-420 | Crowd is a town-entry repeating sound request on the measured route, independent of bird/star frame state; no crowd paint/progression owner was found in the bounded town range. | High / Medium | ✔ promoted | [EXP-0302](../experiments/EXP-0302-town-bird-star-crowd/), direct-call census, analysis population, sound anchors/path/assets and cleanup body |

### TOWN-415

Loader `R1906` binds `TownBirds/Birds1..9/sprites.16a` into view array `+c8`; all nine sheets measure57 frames on both roots. `R1916` derives group `+bc=0..2` and count `+c0=1..3`, so painter `R1917` indexes `array[group*3+i]` for `i=0..count-1`, never crossing a three-entry group under armer-produced state. Count1 selects sound object `+80` (`Birds1.wav`); count2/3 selects `+84` (`Birds2.wav`). A non-null object whose status query returns 0 reaches `R0386` with repeat0. Exact PRNG scaler operands are measured; uniform frequency is not asserted. This refines, and does not retract, `TOWN-157`'s arming result.

**Confidence.** High for loader/selector/index and conditional request relations plus both-root asset counts; Unknown for runtime distribution and audibility

### TOWN-416

`R1917` is gated by view flag `+208&80h`; for each selected entry it compares that slot's `+d8/+dc/+e0` progress with the sprite's own count and draws nonterminal entries through picture slot+18. It then draws `view+70` (`Town_add.bmp`) through keyed slot+38 even on the final active paint when zero birds remain drawable. Only after that composition does terminal-count equality clear bit80h. The admitted `R1908` hub increments all three progress words once while bit80h is active, even when count is 1/2; there is no catch-up loop. This makes the event one layered episode rather than an indefinitely looping sheet and preserves `TOWN-156`/`TOWN-158` outside their already-recorded count correction.

**Confidence.** High for local order, progress and terminal branches; Unknown for actual paint delivery and visible key result

### TOWN-417

Entry `R1383` clears all view flags at `+208` and writes current time to bird clock `+b8`; it does not directly clear progress `+d8/+dc/+e0`, but inactive progress is ignored and the next armer zeroes all three before use. Painter one-shot latch `L11525` initializes hub clock `L11520` and bird delay `L11526`; the armer later replaces delay while preserving that process latch. Image-wide exact-address enumeration finds 4/3/3 references respectively, all owned by `R1489`, and direct entry/leave bodies contain no static reset. Thus re-entry establishes a fresh effective episode/clock while retaining the last process delay absent another writer.

**Confidence.** High for direct reset/reference ownership; Medium for retained-static lifecycle because alias, overlap, bulk and undisassembled writes remain possible

### TOWN-418

Loader `R1906` binds nine distinct 64×44 `Town/stars/S00..S08.bmp` payloads at `+1ac`, writes current `+1c0=-1`, then calls `R1918`, which advances to 0 and selects S00 without sound. Town paint draws selected pointer `+1bc`, when non-null, through slot+18 at view-relative `(340,288)` without testing bit10h. `TOWN-399`'s `b0h→16` mask result reaches pointer default OR at `L11527`; the admitted hub tests bit10h and calls the star step. A step beginning at current0 conditionally requests `Stars.wav` with repeat0, then current1..8 select S01..S08. A request still requires non-null sound and status0 and is not proof of playback.

**Confidence.** High for the local loader/arm/step/paint/request relation and both-root assets; Medium for physical pointer delivery; runtime visibility/audibility Unknown

### TOWN-419

When star step increments current to 9 or above, it increments `L11528`, clears selected pointer `+1bc` and bit10h, and returns. If that static reaches 10, the same terminal arm resets current and the static to 0 but still ends hidden/inactive; only a subsequent selector16 arm starts at 0, can request `Stars.wav`, and selects S01. Consequently the first post-entry arm advances S01..S08 then hides; later arms with retained current≥9 terminate immediately until the tenth terminal call resets the pair. Entry independently repeats loader `-1→0`, restoring S00 without clearing `L11528`. The static has exactly3 image-wide references, all in `R1918`; alias/bulk reset remains outside the census.

**Confidence.** High for local terminal/rearm/reset instructions; Medium for universal process-cross-entry retention; visible repetition Unknown

### TOWN-420

Entry calls sound loader `L11517`, which first calls cleanup then binds `SFX/Town/Crowd.wav` to `+74`; its tail calls `L11529`, where non-null/status0 reaches `R0386` with repeat1. Image-wide direct-call enumeration finds request and loader called only from entry; cleanup `R1913` is called from loader, leave and destructor, and detaches/releases/zeros the crowd object. The `L11530..L11531` census covers39 functions, 12,711 function bytes, zero orphan bytes and 462 undisassembled bytes; within it no crowd draw, frame counter or hub arm is reached. Bird and star share only the admitted hub after separate activation; crowd bypasses it.

**Confidence.** High for entry/request/cleanup calls; Medium for the bounded negative and family separation; audio and any indirect/out-of-range visual crowd remain Unknown

## Tavern interior animation contracts

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-407 | Four loaded tavern interior series reach the central child's own painter, not the button panel. | High / Unknown | ✔ promoted | [EXP-0300](../experiments/EXP-0300-tavern-interior-animation/), `bindings.tsv`, `anchors.tsv`, `assets.tsv`, `bodies.tsv` |
| TOWN-408 | Candle and cauldron share one strict >100-ms paint gate, and their cyclic updater excludes the last loaded entry. | High / Unknown | ✔ promoted | [EXP-0300](../experiments/EXP-0300-tavern-interior-animation/), cyclic helper, strict-gate anchors, `assets.tsv`, `model-vectors.tsv` |
| TOWN-409 | Tender episodes are time-armed by a separate static state machine, not by a local pointer or party-selection test. | High / Unknown | ✔ promoted | [EXP-0300](../experiments/EXP-0300-tavern-interior-animation/), delay/PRNG/mode anchors, `references.tsv`, `model-vectors.tsv` |
| TOWN-410 | Tender breath completes forward, then clears its index without refreshing its cached picture. | High / Unknown | ✔ promoted | [EXP-0300](../experiments/EXP-0300-tavern-interior-animation/), `L07745/L11532/L11523/R1902`, anchors and synthetic stale-cache vector |
| TOWN-411 | Tender drink uses a bounded forward/reverse sequence, not modulo looping. | High / Unknown | ✔ promoted | [EXP-0300](../experiments/EXP-0300-tavern-interior-animation/), forward/reverse/turnaround anchors, `model-vectors.tsv` |
| TOWN-412 | Tavern entry resets each sequence's index and cached pointer; the paint clocks, mode and direction are separate absolute-address state. | High / Medium / Unknown | ✔ promoted | [EXP-0300](../experiments/EXP-0300-tavern-interior-animation/), loader/leave bodies, `references.tsv`, `calls.tsv`, scope and frontier report |
| TOWN-413 | The reached interior sound calls use separate conditions, not one sound per animation frame or guaranteed playback. | High / Unknown | ✔ promoted | [EXP-0300](../experiments/EXP-0300-tavern-interior-animation/), `bindings.tsv`, `assets.tsv`, sound loader/status/request/cleanup bodies and `calls.tsv` |
| TOWN-414 | Local time-driven interior progress is established, but the central child's complete runtime paint schedule and complete interaction effects remain open. | High / Medium | ✔ promoted | [EXP-0300](../experiments/EXP-0300-tavern-interior-animation/), 35-address export manifest, `calls.tsv`, `bindings.tsv`, explicit alternatives and reachable falsifying predictions |

### TOWN-407

Parent entry `R1412` binds sequences at central-child offsets `+1dc/+20c/+23c/+26c` to candle files 0..9, cauldron 0..20, tender breath 1..24 and drink 1..40. `R1902`, installed at `L10203+2c`, draws candle and cauldron unconditionally and the two tender sequences under modes 2 and 1. Draw helper `L11523` reads sequence cached pointer `+14` and invokes picture slot `+18`; it does not select by the index at draw time. Coordinates relative to parent origin are candle `(160,48)`, cauldron `(420,160)`, tender `(240,152)`, from child rectangle `(160,0)-(480,480)` and painter addends. `CenterArea.bmp` is drawn before them. Entry success, non-null loaded pointers and invocation of this painter are preconditions; the selected-character preview is outside this contract.

**Confidence.** High for the local loader-to-draw bindings, primitive slot and coordinates in both hash-identical executables; Unknown for runtime visibility and complete invocation route

### TOWN-408

Painter `R1902` draws current pictures before testing unsigned `now-[L11533] >100`; it then calls `L11534` once on each sequence and stores now. The helper computes signed `(index+1) % (count-1)`, then writes both index and cached pointer. From entry index0 and non-null frames, candle visits indices 0..8 and cauldron 0..19: 9/20 selected entries despite 10/21 loaded. No elapsed-time catch-up loop exists. Both roots contain all loaded files; each last BMP payload hash differs from every earlier payload in its series. This is not a measured 10-fps rate or a claim about other writers.

**Confidence.** High for the conditional instruction relation and excluded final entries; Unknown for actual paint cadence

### TOWN-409

Painter initializes delay `[L11535]` once as `3000 + rand()/16`; the measured PRNG helper returns 0..32767, giving delay 3000..5047 ms. When unsigned `now-[L11536]` exceeds that delay, odd delay sets mode `[L11537]=1` and direction `[L11538]=1`; even delay sets mode2. This arming test does not require idle mode, so a long paint gap can rearm an active drink ascent. Selected mode is drawn before a separate strict >83-ms single-step gate using the same tender timestamp; admitted steps reset that timestamp. Completion chooses a new delay. Mode0 makes no tender-series draw call. Delays and frequencies are not asserted to be uniformly distributed.

**Confidence.** High for local mode writers, timing comparisons and the active-mode rearm counterexample; Unknown for delivered paint frequency

### TOWN-410

Mode2 uses sequence `child+23c`; forward helper `L11532` returns zero at `index=count-1`, otherwise increments index, caches that entry and returns its pointer. With all 24 pointers non-null, completion at index23 clears mode, replaces delay and writes only `child+254=0`. Cached picture `child+250` still names entry23. A later breath episode with no intervening reload therefore first draws cached `br0024.bmp` while index is 0, then advances to index1 (`br0002.bmp`). The first episode after entry starts with `br0001.bmp`; missing pointers can end an episode earlier.

**Confidence.** High for the separate index/cache writes and this conditional counterexample to index-equals-displayed-frame; Unknown for visibility of any particular paint

### TOWN-411

Mode1 draws sequence `child+26c`, then with direction1 calls forward `L11532`. At last index39 that call returns zero; the same admitted step sets direction -1 and calls reverse `L11539`, selecting index38. Reverse decrements and caches while index is nonzero; a later call beginning at index0 returns zero, clears mode/direction and chooses a new delay. With non-null entries and no long-gap rearm, step-admitted draws project indices `0..39,38..0` (79 draws), not two copies of index39. Additional paints can repeat a frame. The arming-before-step counterexample in TOWN-409 can reverse a descending drink back to ascent after a sufficiently long gap.

**Confidence.** High for conditional local progression and turnaround; Unknown for runtime episode duration, interruptions and load failures

### TOWN-412

Each of four entry calls reaches `L07745`, whose tail writes sequence `+18=0` and `+14=frames[0]`. Painter once-only guard byte `L11540` initializes four timestamps and the random delay, not mode/direction. Reference-manager searches for exact addresses `L11536/33c/340/344/348/34c/350/354` yield 40 read/write references, all inside `R1902`. Parent leave `R0786` calls unexported `L11541` on the four sequence objects and `L11542` on sounds; the deleting wrapper calls unexported `L11543`. No direct reset of the named globals occurs in the read entry/leave/picker/input bodies. This does not exclude alias, overlap, bulk, undisassembled or untraced callee writes; cross-visit retention and actual frame release remain Unknown.

**Confidence.** High for entry's four resets and bounded reference counts; Medium for lifecycle interpretation; Unknown for complete exit/destruction and cross-visit state

### TOWN-413

`L11544` binds parent `+84/+88/+8c/+90/+94/+a4/+ac` to `Town/Inn/drink`, `glotok`, `steam`, `water`, `chair`, `enter`, and `Town/Shop/Breath.wav`. After non-null/status-zero checks, central paint requests steam at strict >10000 ms (independent timestamp), chair and Shop/Breath when arming mode2, drink on every paint whose drink index is 30, and glotok on reverse completion. The index30 test is outside the >83-ms step gate and is not direction-gated. Parent own-paint requests water with repeat argument1; entry requests enter with repeat0. Other named requests use repeat0. `R0386` can stop before buffer play because of audio availability, missing buffers, status or allocation; `R1919` reaches indirect buffer slot+30 with repeat flag0/1. Sound teardown is separate from animation globals.

**Confidence.** High for conditional request sites, paths and arguments; Unknown for actual audible output and device timing

### TOWN-414

Parent message `402h` calls its own slot+34 only when campaign state equals 4; dispatcher `R0366` conditionally calls the receiver's own +2c and +38. The read +38/+40 and parent paint helper do not directly invoke `R1902`. Central pointer slot `L11545` returns 0; click/secondary-input bodies `R1806/R1437` delegate selection/actions without directly writing the named animation state. Party pickers `R1920/R1921` directly change `parent+bc`, not interior frames. Their callees and dynamic aliases were not exhausted, so whole-event preservation, actual cadence and no-interaction-ever claims are not promoted. This narrows TOWN-016's display-selection Unknown without closing TOWN-063's invocation Unknown.

**Confidence.** Medium for event/lifetime coverage; High only for named local paths and bounded direct-write observations; runtime Unknown

## Exterior entrance reaction contracts

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-399 | Town entrance selection is written on each delivered pointer-handler invocation, not only on an entrance transition. | High / Medium | ✔ promoted | [EXP-0299](../experiments/EXP-0299-town-entrance-reactions/), `anchors.tsv`, `tables-{en,ru}.tsv`, complete pointer and paint bodies |
| TOWN-400 | The exterior shop sprite is probabilistically armed per delivered selector-1 update and completes a forward cycle even after a delivered leave-to-blank update. | High | ✔ promoted | [EXP-0299](../experiments/EXP-0299-town-entrance-reactions/), `anchors.tsv`, `art.tsv`, path bindings and synthetic boundary vectors |
| TOWN-401 | Every delivered selector-2 update arms the exterior tavern sprite without resetting its current frame. | High / Medium | ✔ promoted | [EXP-0299](../experiments/EXP-0299-town-entrance-reactions/), `anchors.tsv`, `art.tsv`, tables and complete `R1922/L11546` |
| TOWN-402 | The school's two exterior figures have separate retained directions but share one enable bit, and leave does not command a reversal. | High | ✔ promoted | [EXP-0299](../experiments/EXP-0299-town-entrance-reactions/), `anchors.tsv`, `art.tsv`, `model-vectors.tsv`, complete two helpers and hub |
| TOWN-403 | The gate door is polled and reversible; its guard is a separate direction-controlled sprite. | High / Medium | ✔ promoted | [EXP-0299](../experiments/EXP-0299-town-entrance-reactions/), `anchors.tsv`, `art.tsv`, `sounds.tsv`, gate/guard complete bodies and vectors |
| TOWN-404 | The town sign and fluger sequences are independently random-armed ambience, not shop entrance animations. | High | ✔ promoted | [EXP-0299](../experiments/EXP-0299-town-entrance-reactions/), independent PE anchors, tables, loader/paint/hub/helper reads |
| TOWN-405 | Partial correction to TOWN-158: the hub tests nine non-bird flag bits, not eight. | High | ✔ promoted | [EXP-0299](../experiments/EXP-0299-town-entrance-reactions/), `tables-{en,ru}.tsv`, hub range hash |
| TOWN-406 | Partial correction to TOWN-211: mouse and keyboard handler fallback are not identical. | High | ✔ promoted | [EXP-0299](../experiments/EXP-0299-town-entrance-reactions/), mouse/keyboard anchors and complete `R0390` |

### TOWN-399

Constructor `L11547` installs vtable `L11548`; `+4c` is `L11549`, whose call `L11550` reaches `L11546`. Generic message `200h` reaches this slot at `L11551/L11552` if preceding routing returns zero. The handler calls mask sampler `R1914` at `L11553`, writes its result to `town+b4` at `L11554`, and starts guard direction `+ec=1`. Sampled mask `80h/90h/a0h/b0h/c0h` returns `2/1/8/16/4`; other bytes return -1. Paint tests `b4=1/4/2` and draws `Shop_l/Trener_l/Tavern_l` respectively (`L07768/523f/526b`); there is no gate label among these three. The blank-mask arm clears sound latches `+ac/+b0` without resetting entrance frames or their enable bits.

Physical residence, capture, focus loss and delivery through untaken indirect handlers remain Unknown. Sound request sites are conditional calls, not guarantees of audio output.

**Confidence.** High for the local selector/writer/draw contract and table discrimination; Medium for end-to-end physical-pointer behavior

### TOWN-400

`L11555..7135` obtains `rand()%100` and sets flag bit1 only if remainder >95 and bit1 is clear; it does not reset frame `+1e0`. Hub helper `R1923` increments that frame, then resets to 0 and clears bit1 only on equality with `[town+1d8]+4`. The loader binds `+1d8` to `townbirds/shopie/sprites.16a`, whose measured count is 30 on both roots. Paint draws it whenever non-null. If sound latch `+ac` is 0, the pointer arm queries the school sound and can reach cancellation slots `+48/+34`, then conditionally requests `SFX/Town/Shop/Enter.wav` through `R0386` at `L11556`; it sets `+ac=1` and clears `+b0`. Further updates can retry random arming even without first leaving, but do not restart an active frame sequence.

**Confidence.** High for the reached conditional instruction contract; no probability distribution or runtime repetition frequency asserted

### TOWN-401

Selector2 reaches default OR at `L11527` and sets bit2. `R1922` increments `town+168`, resetting to 0 and clearing bit2 on equality with `[town+160]+4`; loader binding is `townbirds/tavern/sprites.16a`, measured10 frames per root. At pre-increment frame0 only, the helper queries shop and school sounds and can call cancellation slots `+48/+34`, then conditionally requests `SFX/Town/Point.wav` at `L11557`. Leaving does not locally clear bit2; after a completed cycle, another delivered update can arm a new cycle. A stationary physical pointer is not proof of absent delivered updates.

**Confidence.** High for conditional local arming, cycle and sound-call sites; Medium for physical residence

### TOWN-402

Selector4 ORs bit4 at `L11558..71c3`, without resetting frames `+1cc/+1d4` or directions `L11559/L11560`. `R1924/R1925` use `rand()%100 >95` at frame0/direction0 to start +1, and at frame10 to choose -1 instead of 0; intermediate frames add direction. At frame0/direction-1 the helper zeroes direction and clears bit4. The hub tests bit4 once and calls both helpers in order, so one clear can leave the other at a nonterminal displayed frame after that tick. Fighter/mage files each contain11 frames. Sound latch `+b0=0` permits school sound `SFX/Town/School/Point.wav` at `L11561` and shop cancellation calls, then sets `+b0=1` and clears `+ac`.

Cross-visit static-direction initialization and actual frequencies remain Unknown. This claim covers exterior figures only.

**Confidence.** High for local transitions and conditional shared-bit consequence; no runtime witness

### TOWN-403

Unconditional hub helper `R1926` first calls availability helper `R1416`; return -1 selects frame8 of `door/T00..T08.bmp` and returns without changing gate latch `+1a4`. Otherwise the current global pointer mask is sampled at `L11562`: selector8 decrements `+1a0` toward0, every other result increments toward8, selecting table `+18c` into displayed `+19c`. Latch transitions conditionally request `GateUp.wav`/`GateDn.wav`; endpoint clearing of bit8 does not gate the helper. Pointer selector8 sets guard step `+ec=-1` only for unavailable return -1; other recovered arms set+1. Unconditional `R1927` adds it to guard frame `+e8`, clamps and sets step0 on overshoot of the eight-frame sprite `+e4`, using direction latch `+f0` to conditionally request `Guard1.wav`/`Guard2.wav`.

Overshoot releases guard audio, potentially in the request's same hub. Re-entry reverses door progress rather than resetting it. Availability changes, indirect callbacks and audible playback remain Unknown.

**Confidence.** High for conditional door/guard paths and frame bindings; Medium for complete runtime behavior

### TOWN-404

Paint `R1489` admits one hub only when unsigned elapsed time exceeds67ms (`L11563/5131`), then stores fresh time without a catch-up loop. Its random checks independently set flags40h/20h (`L11564..516f`, `L11565..51ad`), which dispatch `R1928/R1929`: they advance sign `V00..V09` and fluger `F00..F07`, wrap at 10/8 and clear their bits, conditionally requesting `Flag.wav`/`Flugel.wav` at frame0. Birds retain their separate random delay and flag80h; horse/baba table picks retain separate timers and dervish flag400h advances independently. Stars flag10h is armed by selector16, outside the four entrance selectors. Town-entry crowd request uses repeat argument1. This is separation of named reached mechanisms, not a census of every possible town effect.

**Confidence.** High for the local cause/clock discrimination; no guaranteed wall-clock cadence or global sound absence

### TOWN-405

`R1908..L11566` tests `1,2,4,10h,40h,20h,400h,200h,100h` in addition to bird bit80h. Those nine tests call ten gated helpers because bit4 invokes two. Gate `R1926` and guard `R1927` are unconditional. The original named targets, bird increments and elapsed-time gate stand; only the eight-count wording is retracted.

**Confidence.** High; complete hub instructions and independently extracted nine test immediates discriminate the count

### TOWN-406

In `R0390`, the mouse/400h arm tests `control+34` at `L07746`; null branches to child broadcast `L07722`. Non-null calls its handler at `L03626`, saves its result and unconditionally jumps from `L11567` to `L03641`, bypassing the child broadcast even for return0. Zero still permits the final per-message vtable slot. The keyboard `+38` arm instead tests the returned value at `L11568` and reaches child broadcast `L11569` on zero. The popup-specific null-field result and numeric final-slot mappings of TOWN-211 stand.

**Confidence.** High; explicit unconditional bypass and independently extracted branch operand/direct targets exclude the identical-fallback model

## Town surface, tavern and school: first survey

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-001 | The town surface is composited from many independently loaded assets at view-enter time, not drawn from one picture. | High | ● active | EXP-0183 |
| TOWN-002 | The town's own base picture is `graphics/interface/town/townmain.bmp` (view `+0x68`), and a second, always-loaded picture `graphics/interface/Town/Town_add.bmp` (view `+0x70`) is composited alongside it on every enter, unconditionally. | High | ● active | EXP-0183 |
| TOWN-003 | The mask surface `R1914` (EXP-0060) samples for hit-testing is the exact field `R1906` populates from `graphics/interface/town/townmask.bmp`: view `+0x6c`. | High | ● active | EXP-0183, EXP-0060 |
| TOWN-004 | The town surface's decorative composition is not authored per chapter, mission or campaign progress: every variable element is re-rolled by the engine's PRNG at each view-enter, from small, compiled position tables. | High | ● active (amended) | EXP-0183 |
| TOWN-005 | Four class-keyed decorative sprite sheets — `townbirds/tavern`, `townbirds/fighter`, `townbirds/mage`, `townbirds/shopie` — are loaded unconditionally on every town-enter, into four distinct fields (`+0x160/+0x1c8/+0x1d0/+0x1d8`). | High / Unknown | ● active | EXP-0183 |
| TOWN-006 | Every numbered frame series `R1906` loads has a fixed count that is a compiled loop bound, and the shipped archives on both preserved roots carry exactly that many frames, no more, no fewer. | High | ● active | EXP-0183 |
| TOWN-007 | The three door-label bitmaps (`Shop_l.bmp`, `Trener_l.bmp`, `Tavern_l.bmp`) carry no baked-in localized text: their bytes are identical on the EN and RU roots. | High | ● active | EXP-0183 |
| TOWN-008 | Every literal path this round located in `rom.exe` for the town surface and its two required rooms resolves to a shipped archive node, on both preserved roots, at matching dimensions. | High | ● active | EXP-0183 |
| TOWN-009 | The tavern's left-side portrait/stat panel is two single bitmaps with no loop and no branch: `LeftStats.bmp` → `+0x78`, `LeftPicture.bmp` → `+0x7c`, both under `graphics/Interface/Inn/`. | High | ● active | EXP-0183 |
| TOWN-010 | The tavern draws the roster of people the player can talk to or hire in one shared 6-column cell grid, and the mercenary-hire cells continue the same index numbering right after the talk-only NPC cells, not in a separate grid. | High / Unknown | ● active (partially retracted, superseded) | EXP-0183, EXP-0062 |
| TOWN-011 | `R1902`, the tavern's roster-grid paint routine, is reached exclusively through virtual dispatch, never by a direct call. | High | ● active | EXP-0183 |
| TOWN-012 | The mercenary/NPC portrait sprite sheet `graphics/interface/inn/Unit%d/sprites.16a` ships for indices 1–15 and 29–30 (17 sheets), with indices 0, 16–28 and 31+ never authored. | High / Medium | ● active (partially retracted) | EXP-0183, EXP-0062 |
| TOWN-013 | The tavern's own top-level click handler (`vt+0x54` = `R0790`) does not hit-test locally: it locates the clicked child object and forwards the click to that child's own `vt+0x54`. | High / Unknown | ● active | EXP-0183 |
| TOWN-014 | The shop's and the tavern's top-level click handlers are the same function, byte for byte, not merely the same pattern: both views' `vt+0x54` slot is `R0790`. | High | ● active | EXP-0183 |
| TOWN-015 | The tavern reads its own tip text, `main/text/tips/inn.txt`, through the identical mechanism the town uses for `town.txt`: the same enable flag, the same reader, the same popup constructor, differing only in the popup's own rect. | High | ● active (amended, partially retracted) | EXP-0183, [EXP-0202](../experiments/EXP-0202-chargen-page-geometry/) |
| TOWN-016 | The tavern loads four ambiance decoration frame series unconditionally on enter, with fixed frame counts that are compiled loop bounds: `candle/t%.4d.bmp` = 10, `cauldron/t%.4d.bmp` = 21, `tender/breath/br%.4d.bmp` = 24… | High / Unknown | ● active | EXP-0183 |
| TOWN-017 | The school implements its own top-level click dispatch directly (unlike the shop/tavern's forwarding model, `TOWN-013`/`TOWN-014`) using the town's own raster-mask mechanism, applied independently to two panels. | High | ● active | EXP-0183 |
| TOWN-018 | The school's two panels correspond to 5 mage skill columns (astral, air, earth, water, fire) and 5 fighter skill columns (bow, club, pike, axe, sword), each with 3 art states (`on`/`shine`/`shine_on`)… | High / Medium | ● active (amended) | EXP-0183 |
| TOWN-019 | Withdrawn: the school's ten per-slot fields were read as enabled flags gating a purchase; they are sound sample objects, and the arms they guard play the slot's sample. | — | ✖ retracted | EXP-0183, retracted by [EXP-0485](../experiments/EXP-0485-ledger-corrections/EXP-0485.md) |
| TOWN-020 | The school's skill selection maps five visual indices to slots `1,2,4,3,5`, opcode `0x40` obtains the price and the enabled Train widget sends opcode `0x3d`; the hit test `R1477` sends no command. | High / Medium | ● active (amended, partially retracted) | EXP-0183, EXP-0135, [EXP-0187](../experiments/EXP-0187-general-skill/) |
| TOWN-021 | The school's enter routine has the identical tip-popup call structure as the town's and the tavern's (`TOWN-015`), consistent with reading `main/text/tips/training.txt`… | Medium | ● active | EXP-0183 |
| TOWN-022 | The school ships `movies/training/{mage,fighter}/{m%.4d,tr%.4d}.bmp` animation-frame paths and three key-format strings (`training/npc34m%d`, `training/npc33s%dl%d`, `training/npc34s%dl%d`) whose reader was not located this round. | Unknown | ● active (partially retracted) | EXP-0183; [retraction note](retracted.md#school-training-consumer-corrections) |
| TOWN-023 | A whole-image sweep for the writer that clears the campaign screen's `+0x3dc` state bits (the field EXP-0060 reads as "town is state 0") is inconclusive: the displacement is shared by too many unrelated classes to sweep generically. | Unknown | ● active | EXP-0183 |
| TOWN-024 | The town/school and the shop/tavern do not share one hit-test architecture: two distinct models coexist among the four surfaces reachable from the town… | High | ● active | EXP-0183, EXP-0060, EXP-0164 |
| TOWN-025 | The tip-popup mechanism is uniform across all three views read this round (town, tavern, school): one enable flag, one text reader, one popup constructor, and a per-view rect. | High | ● active | EXP-0183 |
| TOWN-026 | Across the full 252-entry (path, root) census of the town surface and its two required rooms, the only bytes that differ between the EN and RU shipped releases are the three tip text files; every graphic, mask and sprite sheet is identical. | High | ● active | EXP-0183 |

### TOWN-001

`R1906` is the town view's sole asset loader (its only caller is `R1383`, town `vt+0x80`, at instruction `L11570`). Over its full body it issues 19 distinct load call sites against 4 single bitmaps, 3 door-label bitmaps, 4 class-keyed sprite sheets, 4 numbered frame series (10/9/9/8 frames) and 3 PRNG-selected wildlife sprite groups (horse/baba/dervish) — 24 distinct archive paths in total, each a separate `.res` node, each stored in its own field of the view object (`evidence/rom-town-refs.md` §4)

**Confidence.** High (single routine, single caller, every load site and its literal operand read directly)

### TOWN-002

Both load through `R1176` with no branch between them and the loader's entry (`evidence/rom-town-refs.md` §4). Which one paints first, or whether `Town_add.bmp` is a full-frame or partial overlay, was not traced past the load site — this claim is about *what is loaded*, not paint order

**Confidence.** High (literal operands, direct calls)

### TOWN-003

`R1906` loads this path via `R1152` into `+0x6c` and never blits it in that routine (`evidence/rom-town-refs.md` §3–§4); EXP-0060 reads `R1914` sampling `+0x6c` under the cursor. This ties the load path to the known consumer; EXP-0060 read the consumer only, not the loader

**Confidence.** High (both offsets read directly; EXP-0060's own read of the consumer is cited, not re-derived)

### TOWN-004

`R1906` reads no chapter/mission/save field before choosing which wildlife sprite or frame to load. The horse's screen position is one of 5 table entries (the global at `L11571`, `.rdata` literal), rolled via `R0179() % 5`; the baba's is one of 4 (the global at `L08048`, `% 4`); the dervish's is one of 4 (the global at `L11572`), re-rolled until it differs from the baba's pick (`do {...} while (iVar1 == iVar9)`). No branch reads a mission index, a `.reg` key or a save field in this selection path (`evidence/rom-town-refs.md` §4)

**Confidence.** High (the full selection code path was read; the position tables are `.rdata` literals, not `.reg`/save reads)

**Amended.** The roll formulas only: the horse index is `(r*5/0x7fff) mod 5` and the baba and dervish indices are `(r*4/0x7fff) and 3`, not `rand() % n` (`TOWN-505`). The no-chapter, no-mission, no-save finding stands. [`retracted.md`](retracted.md) holds the entry.

### TOWN-005

All four load regardless of the player's chosen class; which one (if any single one, as opposed to a layered set) is drawn was not traced past `R1906` — no paint-dispatch routine for the town's own surface was read this round

**Confidence.** High (load unconditional, all four literal operands read) for the loading fact / Unknown for the selection-for-display fact

### TOWN-006

`sign/V%.2d.bmp` = 10 (`iVar9<10`), `door/T%.2d.bmp` = 9, `stars/S%.2d.bmp` = 9 (same loop as the door series), `fluger/F%.2d.bmp` = 8, `TownBirds/Birds%d/sprites.16a` = 9. `tools/townart` (`-root gameversions/en` and `-root gameversions/ru`) finds exactly 10/9/9/8/9 matching archive nodes on **both** roots (`experiments/EXP-0183-town-screen/evidence/town-screen-art-{en,ru}.csv`)

**Confidence.** High (loop-bound immediates read directly; dual-root archive enumeration confirms the count, not merely one sample)

### TOWN-007

`tools/townart` computes a truncated SHA-256 for every matched node on both roots; of 252 matched (path, root) pairs across the whole town/tavern/school census, exactly 3 differ between EN and RU, and all 3 are `main/text/tips/{town,inn,training}.txt` — the three door-label bitmaps and every other art file in the census are byte-for-byte identical on both roots (`experiments/EXP-0183-town-screen/evidence/town-screen-art-{en,ru}.csv`, diffed by path)

**Confidence.** High (whole-population hash comparison, both roots, all 252 matched entries)

### TOWN-008

`tools/townart` matched 118 of 131 literal/expanded names against 252 total (path, root) rows (118 distinct paths × 2 roots + the 13 numbered-frame expansions each root, plus text files); the 13 unmatched names are gaps this round's own glob over-guessed (mercenary `Unit16..28`, `Unit31`+ absent — see `TOWN-013`), not resolution failures. 0 path-set or dimension mismatches between EN and RU over the matched population (`experiments/EXP-0183-town-screen/evidence/town-screen-art-{en,ru}.csv`)

**Confidence.** High (direct archive enumeration, both roots)

### TOWN-009

`R1930`, 41-line decompile, two `R1176` calls (`evidence/rom-town-refs.md` §5)

**Confidence.** High (literal operands, direct calls, no branch)

### TOWN-010

`R1902`'s first loop (`0 < *(record+200)`) draws one portrait per entry of the tavern's own inner record, gating a decorative icon draw on the mercenary-type-id byte `record+0x15b` (the field published as `MERC-TYPE-001`, EXP-0062); its second loop (`0 < *(record+0xf0)`) indexes from `*(record+200) + i`, continuing the first loop's numbering into the same `(i%6 + (i/6+1)*6)*0x10` cell formula. The two loops' own fields were not identified against EXP-0062's named arrays (`Mercenaries`, `MercenaryCount`); only the shared-grid structure is established, not which named field backs the second loop (`evidence/rom-town-refs.md` §5)

**Confidence.** High (the grid formula and both loop bounds are read directly) for the shared-grid structure / Unknown for the second loop's field identity

**Amended.** half order, see `TOWN-468`; formula restored by `TOWN-467`. [`retracted.md`](retracted.md) holds 2 entries for this claim (partially retracted, superseded).

### TOWN-011

`EnumRefs xr:R1902` over the repaired call table finds exactly one reference, from address `L11573` — a vtable slot, not a function body — consistent with the tavern's own generic child-forwarding click model (`TOWN-014`) also governing paint: the roster grid is a separate child object, not code inline in the tavern view (`evidence/rom-town-refs.md` §5)

**Confidence.** High (single xref, address is inside a known vtable region)

### TOWN-012

`tools/townart`'s dual-root census (`experiments/EXP-0183-town-screen/evidence/town-screen-art-{en,ru}.csv`) finds the identical 17-index set on both roots. Indices 1–15 match EXP-0062's 15-slot mercenary pool (`MercenaryCount`); indices 29–30 exceed that pool and are presumed talk-only NPCs (e.g. the innkeeper) sharing the same sprite-sheet numbering and the same portrait grid (`TOWN-010`) — not confirmed against a named field this round

**Confidence.** High (the index census itself, both roots) / Medium (that 29–30 are talk-only NPCs rather than an unused reservation — inferred from the count mismatch against `MercenaryCount`, not read from a field)

**Amended.** census range, see `TOWN-470`. [`retracted.md`](retracted.md) holds an entry for this claim (partially retracted).

### TOWN-013

`R0790`'s full body: when the container's field at index `0x10` equals its field at index `0xd`, it looks up the child at the point with `R0791`; when a child is found and it is not the container itself, it calls that child's virtual method at vtable offset `0x54` (`evidence/rom-town-refs.md` §2). The forwarded-to child's own hit-test method — and so the tavern's own interactive region shape (rect, mask, or something else) — was not read this round

**Confidence.** High (the forwarder's own bytes, read in full) for the forwarding fact / Unknown for the tavern's own region shape

### TOWN-014

Read directly off the repaired vtables `L03649` (shop) and `L09234` (tavern), `EnumRefs vt:<addr>:144` (`evidence/rom-town-refs.md` §1–§2). This is the generic child-container forwarder of `TOWN-013`

**Confidence.** High (identical address in two independently enumerated vtables)

### TOWN-015

`R1412` (tavern `vt+0x80`) gates on the global at `L03631`, calls `R1293(s_main_text_tips_inn_txt, ...)`, then `R1261(0x467, 0, 0, 0x138, 200, ...)` — popup 312×200 at (0,0). `R1383` (town `vt+0x80`) gates on the same global, calls the same reader with `s_main_text_tips_town_txt`, then `R1261(0x467, 0x148, 0, 0x280, 200, ...)` — popup 640×200 at (328,0). First constructor argument (`0x467`, a window class/style id) is literal-identical in both (`evidence/rom-town-refs.md` §7, §10). **Amended by `EXP-0202` (`TOWN-246`): the town-popup width above is wrong.** The operands are LTRB (`TOWN-184`), so width `=RIGHT-LEFT=0x280-0x148=312`, not the raw `RIGHT` operand `0x280=640`; the correct reading is popup 312×200 at (328,0), equivalently right-edge-anchored at `x=640`. The tavern's own phrase above (312×200 at (0,0)) is unaffected and was always correct

**Confidence.** High (both call sites read directly, all operands literal); the town-popup derived width, 640, carried High while believed and is retracted by `TOWN-246`

**Amended.** [`retracted.md`](retracted.md) holds an entry for this claim (refuted).

### TOWN-016

The tavern loads four ambiance decoration frame series unconditionally on enter, with fixed frame counts that are compiled loop bounds: `candle/t%.4d.bmp` = 10, `cauldron/t%.4d.bmp` = 21, `tender/breath/br%.4d.bmp` = 24, `tender/drink/dr%.4d.bmp` = 40. All four loops run inside `R1412` (`evidence/rom-town-refs.md` §7); which frame is shown at a given moment (timer-driven animation vs. a single static frame) was not traced

**Confidence.** High (loop bounds are immediates, read directly) for the counts / Unknown for the display-selection mechanism

### TOWN-017

School `vt+0x54` (`R1931`) forwards, when a state field `+0x328` is 0, to `R1477`, which after a `PtInRect` on the view's own client rect reads a state field `+0x31c` (values 0 or `0xf`) to select one of two independent `{rect, mask-surface}` pairs — `+0xfc`/`+0xf8` for one panel, `+0x1b0`/`+0x1ac` for the other — each read with the identical "cursor position → row/col → mask-byte lookup → 5-arm switch" structure already established for the town by `R1914` (`TOWN-003`, EXP-0060) (`evidence/rom-town-refs.md` §9)

**Confidence.** High (the dispatch structure is read directly, both panels)

### TOWN-018

The school's two panels correspond to 5 mage skill columns (astral, air, earth, water, fire) and 5 fighter skill columns (bow, club, pike, axe, sword), each with 3 art states (`on`/`shine`/`shine_on`), and one raster mask image per panel (`column/mage/mask.bmp` 100×120 px 8bpp, `column/fighter/mask.bmp`) — all confirmed present at matching dimensions on both preserved roots. `R1477` reads exactly 5 sample-object offsets per panel (`TOWN-019`, retracted: the fields are samples, not enabled flags); `tools/townart`'s dual-root census matches all 5×3 art files per class plus both mask files, 0 dimension mismatches (`experiments/EXP-0183-town-screen/evidence/town-screen-art-{en,ru}.csv`). Which panel (`+0x31c==0` vs `==0xf`) is mage and which is fighter was not tied to a specific mask file by a direct read — inferred only from the 5-slot count matching both classes equally

**Confidence.** High (art inventory and dimensions, dual-root) / Medium (the panel-to-class assignment, inferred from slot count, not read from `+0x31c`'s own writer)

**Amended.** The `enabled-flag` wording for the per-panel offsets, corrected in the text above (`TOWN-019`, [`retracted.md`](retracted.md)). The art inventory, dimensions and panel-to-class inference stand.

### TOWN-019

- What stands is the read. Each of `R1477`'s five arms per panel tests the left-click argument and a field, offsets `+0x84/+0x88/+0x8c/+0x90/+0x94` for one panel and `+0x98/+0x9c/+0xa0/+0xa4/+0xa8` for the other, 5 distinct offsets each, none shared between panels (`evidence/rom-town-refs.md` §9).
- What fell is the meaning. The field is loaded and tested for non-zero (`L07696`..`L11574` for `+0xa8`), then is the `this` of `R1476` (`L11575`) and of the sound call `R0386` (`L11576`..`L07697`). Each is a sample object, listed in the SFX request census as an owner sample field (`school+0x94..+0xa8`, `ChrGen/Skill/*.wav`). No purchase or command is sent from these arms; the arms play the slot's sample when `R1476` returns 0 (`TOWN-478`, `VIDEO-SFX-059`).

**Confidence.** The High was earned for the offsets and the guard shape; it stands only for them. See [`retracted.md`](retracted.md).

**Amended.** Retracted as a whole by EXP-0485 ([`retracted.md`](retracted.md)). The withdrawn headline read: "Each of the school's 10 skill slots (5 per panel) is gated by its own enabled-flag field, checked before any purchase action fires."

### TOWN-020

`R1477`'s click arms call `R1476` with a sample object as `this` and play that sample through `R0386` when it returns 0 (`L11575`..`L07697`); `R1476` is a lookup in the sample channel table `L07698`, not a session or party-list lookup, and the sound is the slot's sample, not an error sound. Command opcodes come from the later widgets: selection maps five visual indices to slots `1,2,4,3,5`, opcode `0x40` obtains the price, and the enabled Train widget sends opcode `0x3d` through `R1932`. `R1476` remains a lookup rather than a packet builder. General is accepted by the server arm but cannot be selected by this UI

**Confidence.** High for the completed school chain (`TOWN-GENERAL-106`..`108`); the older session or party-list lookup and error-sound reading is refuted (Amended)

**Amended.** The command tail by EXP-0187, and the refutation of the lookup and error-sound reading by EXP-0485, are stated in the claim text above. [`retracted.md`](retracted.md) holds the entry.

### TOWN-021

The school's enter routine has the identical tip-popup call structure as the town's and the tavern's (`TOWN-015`), consistent with reading `main/text/tips/training.txt`, but the string operand at the school's own call site was not independently confirmed. `R0706` gates on the same the global at `L03631`, calls the same reader `R1293()`, then `R1261(0x467,0,0,0x1c8,...)` — popup 456×200. The decompiler does not resolve an explicit string argument at this call site (`out/ghidra-export-town/dec-school4/R0706_R0706.c` line 60); a direct `FindStrXref` check on the literal failed with a tool/environment error unrelated to the string's presence (`evidence/rom-town-refs.md` §8)

**Confidence.** Medium (structural pattern match across three views; the operand itself not directly observed at the school's call site)

### TOWN-022

Established only by string presence in the `rom.exe` data section (`str-data.txt` lines 1219–1275, `evidence/rom-town-refs.md` §8); no xref, no loader, no consumer was traced for any of the four

**Confidence.** Unknown (existence only; no reader traced)

**Amended.** the four animation-family consumer Unknowns are resolved by `TOWN-427`..`TOWN-434`; the two `training\npc33s%dl%d` and `training\npc34s%dl%d` key formats are resolved by `TOWN-502`, and the `training\npc34m%d` reader is the training-hall enter routine `R0706` (call site `L03612`, `DLG-RECT-037`). [`retracted.md`](retracted.md) holds entries for this claim (resolved).

### TOWN-023

`EnumRefs disp:3dc` returns 242 references from 80 distinct owner functions spanning code-region — far outside the campaign-screen class's own code range — matching the documented common-base-class-displacement hazard (`docs/INSTRUMENT.md`). A sweep scoped to the campaign screen's own range (approx. `L11577`–`L11578`) was not run. The routine (if any) that explicitly clears these bits on returning to town is not identified

**Confidence.** Unknown (negative finding: this sweep does not discriminate; a scoped sweep was not attempted)

### TOWN-024

The town/school and the shop/tavern do not share one hit-test architecture: two distinct models coexist among the four surfaces reachable from the town, refuting this round's own initial assumption that the town's mask model would generalise. Town (`R1914`/`R0704`, EXP-0060) and school (`R1477`, `TOWN-017`) each implement direct raster-mask hit testing at their own `vt+0x54`. Shop and tavern share one generic child-forwarding function, `R0790` byte-identical in both vtables (`TOWN-014`), which locates a child object and defers to *that* object's own `vt+0x54` — for the shop, a precomputed `CRect` array (EXP-0164, `SHOP-SCREEN-031`); for the tavern, not traced past the forward (`TOWN-013`).

This experiment's own first hypothesis, that `R0385` (called at every enter routine, town/tavern/school alike) implemented view-exclusivity/switch semantics, was refuted by decompiling it: it is a generic `Add(item); item->owner=this;` list utility with no exclusivity logic (`evidence/rom-town-refs.md` §10)

**Confidence.** High (both architectures read at instruction level; the refuted hypothesis is recorded, not merely dropped)

### TOWN-025

All three call `R1293` gated on the global at `L03631`, then `R1261` with a literal-identical first argument `0x467` and a view-specific `(x, y, w, h)` — town `(0x148,0,0x280,200)`, tavern `(0,0,0x138,200)`, school `(0,0,0x1c8,200)` (`evidence/rom-town-refs.md` §7, §9, §10). The school's own text-file operand is not directly confirmed (`TOWN-021`)

**Confidence.** High (town and tavern operands read directly; school inferred by structure only, per `TOWN-021`)

### TOWN-026

`tools/townart` computes a SHA-256 per matched node on both roots and diffs them by path: 249 of 252 match exactly, and the 3 mismatches are `main/text/tips/{town,inn,training}.txt` (sizes 443/444/505 bytes EN vs 421/493/368 bytes RU) (`experiments/EXP-0183-town-screen/evidence/town-screen-art-{en,ru}.csv`). This is the one place this round found the two releases diverge at the byte level; it is a DATA VALUE (localized text), not an engine constant, and lifting it changes only the three `.txt` payloads, no code and no other shipped file

**Confidence.** High (whole-population hash comparison, both roots)

## World-map view

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-036 | The world-map background is one stored picture, loaded once by one unconditional call, and never reloaded while campaign progress changes. | High | ● active (amended) | [EXP-0184](../experiments/EXP-0184-world-map/), [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-037 | `R1933`'s other 11 loads are the same pattern — one literal, one reference, no branch — and split by archive on the `main\` prefix alone. | High | ● active | [EXP-0184](../experiments/EXP-0184-world-map/) |
| TOWN-038 | A dedicated routine builds two parallel per-MapObject arrays from the registry — one absolute rectangle for hit-testing, one raw point for anchoring and route endpoints — and it is the sole reader of both `MapRect` and `MapPoint`. | High | ● active (amended) | [EXP-0184](../experiments/EXP-0184-world-map/), [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-039 | A region is a rectangle; hit-testing is a linear `PtInRect` scan of the rect array `TOWN-038` built, and the winning index feeds a table lookup, not a direct dispatch. | High / Medium | ● active | [EXP-0184](../experiments/EXP-0184-world-map/) |
| TOWN-040 | The map-rectangle return gate has two fully traced sources: a static `Picture == "nothing"` bit and the selected-mission marker cache. It is not an availability test over two offer lists. | High | ● active (amended, partially retracted, superseded) | [EXP-0184](../experiments/EXP-0184-world-map/), [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-041 | A per-mission marker icon is built on demand, memoized by mission number, and only for a MapObject whose `Picture` key is not `"nothing"` — closing the picture-source half of a hit region's marker. | High | ● active | [EXP-0184](../experiments/EXP-0184-world-map/) |
| TOWN-042 | The `PictureOffset` registry key is read by the marker-icon builder but is never present in either shipped registry — the code path it feeds is dead on shipped data, not absent from the engine. | High / Unknown | ● active | [EXP-0184](../experiments/EXP-0184-world-map/) |
| TOWN-043 | `PathMap.bmp` loads through a different path than the fixed pictures, into a temporary route-graph builder; it is consumed as indexed topology and never blitted. | High | ● active (amended, partially retracted, superseded) | [EXP-0184](../experiments/EXP-0184-world-map/), [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-044 | Mission title and briefing text for the world map's offered-mission display are read from `main\text\battle\m<n>\{title,briefmap}.txt`, keyed by mission number… | High | ● active | [EXP-0184](../experiments/EXP-0184-world-map/) |
| TOWN-045 | The world-map view draws no aggregate campaign-progress indicator: no chapter number, completion count, or equivalent. | High | ● active (amended) | [EXP-0184](../experiments/EXP-0184-world-map/), [EXP-0188](../experiments/EXP-0188-pathmap/) |

### TOWN-036

`R1933`, the world-map screen's asset-load routine, is a straight-line function of 12 sequential resource loads. The first, `main\graphics\Global.Map\GMap.bmp`, has one literal reference image-wide and is the only load carrying the `main\` archive-selector prefix. `EXP-0188` read paint `R1530`: this picture is the first, full background blit. “One picture” describes the background source, not the whole composed frame; route dots, markers, flags, cross, and scrolls are later overlays (`TOWN-120`)

**Confidence.** High (sole literal reference and straight-line load, plus direct paint consumer)

**Amended.** clarified by EXP-0188.

### TOWN-037

An `imm:` sweep of all 11 remaining literal addresses (Flag1/Flag/Cross `sprites.16a`, BallMap/Hero/Scroll01-03/ScrollP1-3 `.bmp`) returns exactly 1 hit each, in `R1933`, 0 in orphan code — the same result `TOWN-036` established for `GMap.bmp`. None of the 11 carries the `main\` prefix. Archive census (`evidence/gmap-assets.txt`, both `gameversions/en` and `gameversions/ru`): the one prefixed file (`GMap.bmp`) resolves inside `main.res`, stored there as `graphics/global.map/gmap.bmp` (the archive keeps the `graphics\` segment); the 11 un-prefixed files resolve inside `graphics.res`, each stored with its leading `graphics\` segment already removed (`global.map/ballmap.bmp`, `global.map/flag/sprites.16a`, …).

So the selector observed in code (a `main\` prefix on the load string) and the selector observed in the archive (which container actually holds the path) agree: `main\` routes a load to `main.res`, its absence to `graphics.res`, and `graphics.res`'s own internal paths omit the segment that names it

**Confidence.** High (both the code-side selector and the archive-side resolution are read directly — an exhaustive `imm:` sweep for the code half, a full-archive listing for the data half, reproduced on both preserved roots)

### TOWN-038

`R1529` opens `Scenario\GlobalMap.reg`, reads `[General] ObjectCount`, and loops the 1-based `MapObject%d` sections `1..ObjectCount`. Per object it reads `MapPoint` (2 ints, x/y) and `MapRect` (4 ints, stored x/y/width/height), converts the rect to absolute `{x, y, x+w, y+h}` at load time (`local_54 = x + w; local_50 = y + h`, R1529:53-56) and appends it to a grow-by-doubling array at `this+0xb0` (data) / `this+0xb4` (count); the raw point is appended, unconverted, to a second grow-by-doubling array at `this+0xc4` (data) / `this+0xcc` (capacity) / `this+200` (count). `EXP-0188` follows that second array into scroll destinations, current-position flags, markers, and route endpoints (`TOWN-118`..`TOWN-120`). `MapRect` is not the mission-selection control. Sole builder caller: `R1659` (`callto:R1529` = 1 hit, 0 orphan)

**Confidence.** High (the read, conversion arithmetic, both appends, and later MapPoint consumers are read at instruction level; sole-builder-caller and sole-registry-reader results are exhaustive sweeps)

**Amended.** extended by EXP-0188.

### TOWN-039

`R1863` (entered only while `this+0x188 != 0` and four other busy-flags are all 0) computes the cursor position relative to the view origin (`this+8`/`+0xc`) and scans `i = 0..this+0xb4-1` over the rect array at `this+0xb0` (stride `0x10`), calling the Win32 `PtInRect` on each entry until one contains the point; `i` is the hit region's index, in `MapObject` order. The winning index is then passed to `R0668(i)`, which returns `*([L04369] + (base + i) * 4)` — an indexed table lookup, not a call. What the returned value does (a command id posted to a message dispatcher, a callback pointer, or something else) was not traced past this point

**Confidence.** High for the shape and the scan (region = axis-aligned rectangle, hit test = linear `PtInRect`, both read at instruction level) / Medium for what a hit *does* (the lookup that follows is read, but its consumer is not, so "leads to a mission" is not yet shown end to end from this function alone — `TOWN-041` closes part of that gap from the mission-select side)

### TOWN-040

`R1529` appends one short at view `+0x98`: 1 for each MapObject whose `Picture` resolves to `"nothing"`, 0 for the five picture-bearing objects; `+0x9c` is that array's data pointer. On a zero, `R1863` scans campaign `+0x6a4/+0x6a8`. That is campaign record `+0x548` plus marker-cache `+0x15c/+0x160`, whose writer is `R1325`: selecting a mission appends one 0x14-byte record only when its mapped object has a non-`"nothing"` picture, and world-map paint consumes the same cache. Therefore 30 non-picture objects return their indexed table value unconditionally; a picture-bearing object returns it only after its marker exists. The generic consumer of that returned value remains the `TOWN-039` Unknown. The earlier “currently-offered missions” and writer-Unknown clauses are retracted

**Confidence.** High (both writers, both reads, the `Picture` comparison, the shared cache identity, and the paint consumer are direct instructions; the correction discriminates against the offer-list model by object-relative field arithmetic `0x548+0x15c=0x6a4`)

**Amended.** amended; two clauses retracted. [`retracted.md`](retracted.md) holds an entry for this claim (superseded).

### TOWN-041

`R1325(this, missionNumber)` checks a cache array at `this+0x15c`/`+0x160` (stride 20 bytes, first field = mission number); on a miss it maps the mission number to an object index via `R1324`, reads that `MapObject<k+1>`'s `Picture` key (default `"nothing"`), and only when the value is not `"nothing"` builds the icon path `main\graphics\Global.Map\<Picture>.256` and appends a new cache entry. Sole caller: `R0785`, itself the last of a four-step sequence — `R1413`; `R1407` (mission-select router, `REG-SCN-062`); `R1410` (announce-flag latch, `REG-SCN-065`); `R1325` — i.e. the marker icon is (re)built as the final step of *selecting* a mission, not on a fixed schedule or a draw-loop poll.

Registry census, both roots (`evidence/globalmap-reg-catalog-{en,ru}.txt`): exactly 5 of 35 `MapObject` sections carry a `Picture` value, identical set on both roots (`onmap03`, `onmap13`, `onmap14`, `onmap18`, `onmap19`); archive census (`evidence/gmap-assets.txt`) finds the corresponding 5 `.256` files present in `main.res` on both roots, byte-different between EN and RU on all 5, and differing in file *size* on 2 of the 5 (`onmap18`: 4245 EN / 4118 RU; `onmap19`: 1841 EN / 1799 RU)

**Confidence.** High (the memoization, the mission-select call chain and the `Picture`-key gate are read at instruction level; the 5/35 and the EN/RU size difference are exhaustive corpus censuses on both preserved roots)

### TOWN-042

`R1325` reads `PictureOffset` (`R1430`, address L11579) inside the same block that builds a marker's `.256` path; an `imm:`/`refto:` sweep of that literal's address returns exactly 1 hit each, both at `L11580`, inside `R1325`. The full key catalog of `globalmap.reg` on both roots (`evidence/globalmap-reg-catalog-{en,ru}.txt`, 141 records each) lists `MapPoint` x35, `MapRect` x35, `Picture` x5, `MissionObjects` x1, `ObjectCount` x1, and **no `PictureOffset` entry on either root**. What the key would do if present — the read is immediately followed by two locals set to 0 unconditionally, and whether that overwrites or is independent of the read was not resolved — is Unknown

**Confidence.** High for the absence (an exhaustive key-catalog census of the whole registry file, both roots, and an exhaustive reference sweep of the literal) / Unknown for the key's intended effect (the consuming instructions past the read were not resolved)

### TOWN-043

Its literal has one reference image-wide, `R1934` at `L11581`, reached through `R1531` from world-map enter `R1409`. Graph construction is unconditional and precedes the current-position test that may build the mission-scroll list; the former placement “conditionally, inside the mission-offer build” is retracted. The builder scans literal 640x480 dimensions, records index-2 nodes, walks index-1 corridors, precomputes node-to-node point lists, transfers the matrix to view `+0x148`, and destroys the loaded bitmap before view activation (`TOWN-116`, `TOWN-117`). The prior “consumer was not located” clause is also retracted

**Confidence.** High (sole literal reference and enter control flow, full builder/transfer/paint read in decompiler and raw instructions; three-root pixel census)

**Amended.** amended; two clauses retracted. [`retracted.md`](retracted.md) holds an entry for this claim (superseded).

### TOWN-044

Mission title and briefing text for the world map's offered-mission display are read from `main\text\battle\m<n>\{title,briefmap}.txt`, keyed by mission number — the same file-per-mission family `MISSION-TEXT-005` already documents for in-mission dialogue, extended with two new leaf names. `R1409` (the offer-list builder `REG-SCN-062`/`REG-SCN-065` already locate) opens `main\text\battle\m%d\title.txt` and `main\text\battle\m%d\briefmap.txt` per offered mission, plus a `"%s: %d c"` payment-amount format. Corpus census, both roots (`evidence/mission-text-files-{en,ru}.txt`): exactly 28 `title.txt` and 28 `briefmap.txt` files under `main\text\battle\m*\`, matching `REG-SCN-059`'s 28 `[MissionObjects] Mission<n>` keys 1:1 on both EN and RU.

RU ships 321 total `battle\m*\*.txt` files against EN's 318 — 3 more, unidentified this round, outside `title.txt`/`briefmap.txt` which are equal on both roots

**Confidence.** High (the format strings and the three open calls are `MISSION-TEXT-005`'s own instruction-level reading, re-confirmed for these two leaf names; the 28/28 correspondence is an exhaustive census against `REG-SCN-059`'s own denominator, on both preserved roots)

### TOWN-045

`EXP-0188` closes the prior scope gap by reading the view's complete own-paint routine `R1530`: its consumers are background, cached mission markers, route dots, destination cross, hovered-destination flag, current-position flag, and the scroll list. The scroll renderer draws per-mission title, briefing, and optional payment. Neither routine reads or formats an aggregate progress value. This does not deny the route's own spatial progress animation (`TOWN-120`)

**Confidence.** High for this view's own paint and scroll composition (full routines and all direct draw calls read); it does not cover a generic parent/overlay independently composed outside the view

**Amended.** closed by EXP-0188.

## School paint, tavern children and the right-hand column

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-061 | The school's own-surface paint routine is `R1487` (school `vt+0x2c`), and its draw order is: (1) main background/portrait `+0x70`; (2) mage-side lower area `+0xf4` or a fill; (3) fighter-side lower area `+0x1a8` or a fill… | High | ● active (amended, partially retracted, superseded) | EXP-0185, [EXP-0195](../experiments/EXP-0195-school-column/) |
| TOWN-062 | Tavern's own `vt+0x2c` (`R1935`) is not a paint routine: it has no blit call, only two independent gated checks — a conditional forward to the shared epilogue `R1311` on flag `+0x13c`… | High | ● active | EXP-0185 |
| TOWN-063 | A per-view chain shared byte-for-byte across tavern/shop/town/school's own vtables gates then calls the view's own `+0x2c` followed by its own `+0x38`; neither iterates a child list. | High / Unknown | ● active | EXP-0185 |
| TOWN-064 | The tavern constructs exactly three fixed-rect children, in this order: left panel (id `0x44d`, rect (0,0)-(160,480)), button/decoration panel (id `0x44e`, rect (480,0)-(640,238)), center/roster panel (id `0x450`… | High | ● active | EXP-0185 |
| TOWN-065 | Supersedes `TOWN-010`'s cell-index formula only. | High | ✖ retracted | EXP-0185 |
| TOWN-066 | The roster/center child's own hit-test region is a Win32 `RECT` stored inline in the child object's own memory, addressed by `TOWN-065`'s corrected slot formula — not a raster mask and not live cell-size arithmetic. | High / Unknown | ● active (partially retracted) | EXP-0185 |
| TOWN-067 | Closes `TOWN-018`'s open panel-to-class clause: `+0x31c==0xf` is the mage panel, `+0x31c==0` is the fighter panel. | High | ● active | EXP-0185 |
| TOWN-068 | The school's fighter panel has 5 skill-icon sub-rects at object offset `+0x1c0`..`+0x20c` (five consecutive 16-byte rects): slot0 (200,196,280,228), slot1 (200,216,280,252), slot2 (200,272,280,288), slot3 (200,248,280,276)… | High | ● active (partially retracted) | EXP-0185, [EXP-0195](../experiments/EXP-0195-school-column/) |
| TOWN-086 | The right-hand 160-pixel column (container `campaign+0xd4` plus its four children, ids 5/6/7/8 at `campaign+0xd8/dc/e0/e4`) is built at exactly one construction site, `R0315`, and its own rect exactly complements the map view's… | High | ● active | [EXP-0186](../experiments/EXP-0186-right-column/) |
| TOWN-087 | The tavern borrows `campaign+0xe0` (the id-7 panel `SHOP-FIGURE-041` names) into its own uncovered rect by the identical `RemoveChild`/`OffsetRect(±(640-width))`/`SetRect`/`AddChild` sequence the shop uses… | High | ● active | [EXP-0186](../experiments/EXP-0186-right-column/) |
| TOWN-088 | The school borrows the same panel by the same mechanism as `TOWN-087`'s tavern — not previously published. | High | ● active | [EXP-0186](../experiments/EXP-0186-right-column/) |
| TOWN-089 | The town surface does not borrow `campaign+0xe0`: its enter/leave routines never reference that field, the container, or the literal `0x280`, and its own construct-children step builds exactly one child, with an all-zero rect. | High | ● active | [EXP-0186](../experiments/EXP-0186-right-column/) |
| TOWN-090 | Town's own paint routine covers the column's screen-space region with continuous full-width background art plus one class-keyed decorative sprite whose anchor lands inside the column's x-range… | High | ● active | [EXP-0186](../experiments/EXP-0186-right-column/) |
| TOWN-091 | Widget 5 (`campaign+0xd4`'s first child, container-local rect (0,0,160,158)): constructor `R1936`, vtable `L01176`, own `vt+0x2c` `R0310`. | High / Medium | ● active (amended, partially retracted) | [EXP-0186](../experiments/EXP-0186-right-column/), [EXP-0211](../experiments/EXP-0211-column-controls/) |
| TOWN-092 | Widget 6 (container-local rect (0,158,160,238)): constructor `R1150`, vtable `L01177`, own `vt+0x2c` `R0316`, an 8-cell selectable grid. | High / Unknown | ● active | [EXP-0186](../experiments/EXP-0186-right-column/) |
| TOWN-093 | Widget 8 (container-local rect (0,480,160,screenH), zero height at the shipped 640x480): constructor `R1937`, vtable `L06483`, own `vt+0x2c` `R0349`… | High / Unknown | ● active (amended, partially retracted) | [EXP-0186](../experiments/EXP-0186-right-column/), **[EXP-0210](../experiments/EXP-0210-figure-panel-controls/)** |
| TOWN-094 | Each surface's upper-right widget is a distinct, non-shared routine and rect — the tavern's `(480,0,640,238)` is not shared by the other two, and no button/routine is common to all three. | High | ● active | [EXP-0186](../experiments/EXP-0186-right-column/) |
| TOWN-095 | Refines `TOWN-062`: the shop, like the tavern, has no direct-paint content of its own — its own `vt+0x2c` is the shared epilogue `R1311`, the same function town/school/tavern call at the end of their own `+0x2c`. | High | ● active | [EXP-0186](../experiments/EXP-0186-right-column/) |
| TOWN-096 | Every rect, offset and count found in the right column this round is a `PUSH` immediate in `rom.exe`; none traces to a `.reg` key or a shipped file's own content — each is a G2 engine-class limit. | High | ● active | [EXP-0186](../experiments/EXP-0186-right-column/) |

### TOWN-061

The school's own-surface paint routine is `R1487` (school `vt+0x2c`), and its draw order is: (1) main background/portrait `+0x70`; (2) mage-side lower area `+0xf4` or a fill; (3) fighter-side lower area `+0x1a8` or a fill; (4) a state-indexed decorative sprite chosen from `+0x30c` by the active-panel selector `+0x31c`; (5) the active panel's 5 skill icons (fighter array `+0x264` when `+0x31c==0`, mage array `+0x2b4` when `+0x31c==0xf`, corrected by `TOWN-149`); (6) the current frame `+0x25c` of the 9-frame `diamond` animation (`TOWN-154`, correcting this row's original "selected-skill description panel"); (7) a scrollbar-thumb element gated on `+0x68`; (8) a scrollbar-track element `+0x74`.

Gated on `+0x338 != 0`; ends with a call to the shared epilogue `R1311`. Full decompile (`evidence/rom-refs.md` §6, `EXP-0185`). **Amended by `EXP-0195`**: the step-5 class labels were the wrong way round and step 6 was misidentified; the eight-step order, the fields, the gate and the epilogue stand, and `TOWN-146`, `TOWN-149`, `TOWN-150` and `TOWN-154` give each step's primitive and destination

**Confidence.** High (single routine, full decompile, every blit call and its gate read directly) for the draw order, the fields and the gate / the two corrected clauses carried High while believed

**Amended.** [`retracted.md`](retracted.md) holds 2 entries for this claim (refuted, superseded).

### TOWN-062

Tavern's own `vt+0x2c` (`R1935`) is not a paint routine: it has no blit call, only two independent gated checks — a conditional forward to the shared epilogue `R1311` on flag `+0x13c`, and a resolve/error-tone check (`R1476`/`R0386`) on flag `+0x90`. Tavern's own `vt+0x30` (`R1938`) is an empty stub. By contrast town's own `+0x2c` (`R1489`) and school's own `+0x2c` (`TOWN-061`) are both long direct-blit routines. `+0x2c` is therefore each top-level view's own-content paint slot for town and school, and is repurposed at the tavern, which has no own-surface content to draw — its whole surface is its three children (`TOWN-064`) (`evidence/rom-refs.md` §6, §8)

**Confidence.** High (both routines fully decompiled; the four-view comparison is a direct read at every vtable, not an inference from one)

### TOWN-063

`R0366` (shared `+0x34`): `if ((*this+0x20)(0x20)==0) { (*this+0x2c)(); (*this+0x38)(); }`. `R0354` (shared `+0x38`): builds a `CRect` from the view's own client-rect fields and forwards it to `R0351` (not decompiled). A separate shared slot, `+0x40` (`R1696`), does walk the generic child list (`R1868`/`R1869`, the same accessor `TOWN-013`'s forwarder uses) calling each child's own `+0x40` — but nothing read this round calls `+0x40` from the `+0x2c`/`+0x30`/`+0x34`/`+0x38` chain, so what invokes the tavern's three children's own paint overrides (`TOWN-064`) was not located

**Confidence.** High (all four vtables' slots `+0x00..+0x4c` read directly and compared; every routine quoted decompiled in full) for the shared-slot facts / Unknown for what calls a child's own `+0x2c`

### TOWN-064

The tavern constructs exactly three fixed-rect children, in this order: left panel (id `0x44d`, rect (0,0)-(160,480)), button/decoration panel (id `0x44e`, rect (480,0)-(640,238)), center/roster panel (id `0x450`, rect (160,0)-(480,480)) — and each has its own `+0x2c` doing real drawing. `R1939` (tavern `vt+0x78`), full decompile. Left panel's own `+0x2c` (`R1804`) reads the shared hover index `+0xb8` and the two arrays `+0xc4`/`+0xec` split at `+200` (`TOWN-010`'s own fields) and hands the resolved pointer to `R0892`, consistent with displaying the hovered roster entry's portrait/stats (`LeftStats.bmp`/`LeftPicture.bmp`, `TOWN-009`).

Button panel's own `+0x2c` (`R1904`) draws up to 3 action buttons plus a numeric readout, gated by a talk/hire toggle state. Center panel's own `+0x2c` is `R1902`, the already-published roster grid (`TOWN-011`) (`evidence/rom-refs.md` §1-§3)

**Confidence.** High (the constructor is fully decompiled; each child's own vtable base is corrected from a `FindVtables` merge artifact and re-verified by a clean `EnumRefs` at each corrected base; left/button panel `+0x2c` bodies are read directly)

### TOWN-065

The tavern roster grid's cell-index formula, read from raw disassembly rather than decompiler C, is `slot(i) = (i mod 6) + (floor(i/3) + 1) * 6`, byte offset `slot(i) * 0x10` — not `(i%6) + (i/6+1)*6` as published. The two agree only for `i < 3` and diverge from `i == 3` on (published: slot 9; corrected: slot 15). Confirmed in three independent instruction sequences: the roster paint routine's two loops (`R1902`, talk-only and mercenary, the latter on the global index `talkCount+j`) and the hit-test routine (`R1528`) — all three compute the same magic-multiply-by-3-then-`IDIV`-by-6 sequence, byte-identical in shape. `tools/roomgeom` tabulates `slot(i)` for `i=0..23` mechanically, both roots (`evidence/rom-refs.md` §4-§5, `experiments/EXP-0185-room-composition/evidence/roster-grid-slots-{en,ru}.csv`)

**Confidence.** High (raw disassembly of three independent code sites, byte-identical pattern, discriminates against the published formula from `i=3` on — not corpus agreement, an instruction-level correction)

**Amended.** see `TOWN-467`. [`retracted.md`](retracted.md) holds an entry for this claim (retracted).

### TOWN-066

`R1806` (center child `vt+0x54`) calls `R1528`, which computes `slot(i)` per `TOWN-065` and calls `PtInRect` against `this + slot(i)*0x10` for `i = 0..talkCount+mercCount-1`, then `R1807(i)` applies the resolved index. The left panel's own `+0x54` (`R1805`) is a trivial `return 0` — not independently hit-testable. The button panel's own `+0x54` (`R1811`) resolves a button index through `R1810` (not decompiled this round); its own region shape is not established (`evidence/rom-refs.md` §3-§4)

**Confidence.** High (the center/roster child's shape: raw disassembly, cross-verified against the independently-derived paint-routine formula) / Unknown (the button panel's own shape)

**Amended.** formula citation, see `TOWN-467`. [`retracted.md`](retracted.md) holds an entry for this claim (partially retracted).

### TOWN-067

`R1940`'s raw disassembly (one basic block, no branch) gives rect `+0xfc`=(188,188,288,308), 100x120, tested when `+0x31c==0xf` (`TOWN-017`); rect `+0x1b0`=(192,192,284,312), 92x120, tested when `+0x31c==0`. `tools/roomgeom` reads the shipped mask files on both roots: `column/mage/mask.bmp` = 100x120 bpp=8, `column/fighter/mask.bmp` = 92x120 bpp=8. The two panel rects match the two class masks exactly and exclusively (`evidence/rom-refs.md` §7, `experiments/EXP-0185-room-composition/evidence/school-panel-rects-{en,ru}.csv`, `school-mask-dims-{en,ru}.csv`)

**Confidence.** High (exact, exclusive dimension match, dual-root, superseding an inference-only Medium clause)

### TOWN-068

The school's fighter panel has 5 skill-icon sub-rects at object offset `+0x1c0`..`+0x20c` (five consecutive 16-byte rects): slot0 (200,196,280,228), slot1 (200,216,280,252), slot2 (200,272,280,288), slot3 (200,248,280,276), slot4 (200,288,280,308) — all 80px wide (`left=200,right=280`). **The shared-array clause is refuted by `TOWN-148`**: the paint loop reads this run only on the `+0x31c == 0` branch, and reads a second run at `+0x10c`..`+0x14c` on the `+0x31c == 0xf` branch. The icon pointer arrays are `+0x264` fighter and `+0x2b4` mage (`TOWN-149`), and the per-slot enable flags (`TOWN-019`) differ by class. Read from `R1940`'s raw disassembly, one basic block, no branch (`evidence/rom-refs.md` §7).

Why the array's index order and its visual y-order diverge (slot2's y-range sits below slot3's despite the lower array index) was **answered by `TOWN-152`**: the array is indexed by the stored slot, which is the loader's file order, while the per-slot enabled flags are indexed by the visual top-to-bottom order, and the two differ by the permutation `0,1,3,2,4`

**Confidence.** High (the rect values: raw disassembly, no control-flow ambiguity) / the shared-array reuse carried High while believed and is refuted / the index-to-row visual mapping was Unknown and is now answered

**Amended.** [`retracted.md`](retracted.md) holds an entry for this claim (refuted).

### TOWN-086

The right-hand 160-pixel column (container `campaign+0xd4` plus its four children, ids 5/6/7/8 at `campaign+0xd8/dc/e0/e4`) is built at exactly one construction site, `R0315`, and its own rect exactly complements the map view's — the two partition the screen width at every resolution, not only 640x480. `R0315` calls `R1941(4,screenW-0xa0,0,screenW,screenH)` for the container immediately after `R0391(0,0,screenW-0xa0,screenH)` for the map view (`campaign+0xd0`, `SESS-VIEW-028`'s object), then `R1936(5,...)`, `R1150(6,...)`, `R1763(7,...)` (`SHOP-FIGURE-041`), `R1937(8,...)` in that order, all inside the same routine.

`EnumRefs callto:` on `R0315` itself and on the id-4/5/6/8 constructors each returns exactly 1 hit / 1 owner image-wide (`evidence/callsites.txt`); id 7's uniqueness was already established by `SHOP-FIGURE-041`. `screenW`/`screenH` are the global at `L00618`/the global at `L01259`, the same globals `SESS-VIEW-028` reads at this identical allocation for the map view's own rect argument

**Confidence.** High (five constructors, one routine, `EnumRefs` cross-check on all five image-wide)

### TOWN-087

The tavern borrows `campaign+0xe0` (the id-7 panel `SHOP-FIGURE-041` names) into its own uncovered rect by the identical `RemoveChild`/`OffsetRect(±(640-width))`/`SetRect`/`AddChild` sequence the shop uses, reading and writing the same fields `campaign+0xd0`/`+0xd4`/`+0xe0`. `R1412` (tavern enter): `RemoveChild` at `L11582` (target `R0384`), the 0x280 offset added to the running value at `L11583`, `OffsetRect` at `L11584` (`[L09783]`), `SetRect` at `L11585` (`L11586`), `AddChild` at `L11587` (`R0385`). `R0786` (leave) is the exact inverse. Falsifies the alternative that the tavern owns a separate, similarly-coded panel: all three surfaces' routines dereference the identical `campaign+0xe0` field, not a field of their own object, and resolve `RemoveChild`/`AddChild` through the one container `TOWN-086` names (`evidence/rom-column-refs.md` §2-§3)

**Confidence.** High (both routines read at instruction level, both roots agree, discriminated against the separate-panel alternative)
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### TOWN-088

`R0706` (school enter): `RemoveChild` at `L11588`, the 0x280 offset added to the running value at `L11589`, `OffsetRect` at `L11590`, `SetRect` at `L11591`, `AddChild` at `L11592`. `R1915` (leave) is the inverse: the 0x280 offset subtracted from the running value at `L11593`, `OffsetRect` at `L11594`, `RemoveChild` at `L11595`, `AddChild` at `L11596`. Byte-for-byte the same call-target sequence as the tavern's, with the `+0xd0`/`+0xe0` field-load order swapped (`evidence/rom-column-refs.md` §3)

**Confidence.** High (both routines read at instruction level, both roots agree)
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### TOWN-089

`R1383`/`R1912` (town enter/leave), full decompile, contain neither the field nor the literal (town's own tip widget's `(328,0,640,200)` rect shares the value `0x280` only as a right-edge coordinate, not as an offset). `R1942` (town construct-children): `R0700(4,0,0,0,0,...)`, all-zero rect (`evidence/rom-column-refs.md` §4)

**Confidence.** High (both routines full decompile, negative result scoped to the fields and literal actually searched)

### TOWN-090

Town's own paint routine covers the column's screen-space region with continuous full-width background art plus one class-keyed decorative sprite whose anchor lands inside the column's x-range — closing `EXP-0183`'s open item on which routine reads the class-keyed sprite fields for paint. `R1489` (town's own `vt+0x2c`, gated on `+0x204 != 0`) blits `+0x68` (`townmain.bmp`, BMP 640x480 24bpp, 921656 bytes both roots) at local offset `(0,0)`, covering the full 640px width in one blit; and blits `+0x1c8` (`townbirds/<class>/sprites.16a`, class-keyed per `EXP-0183`'s `rom-town-refs.md`) at local offset `(0x204,0x158)` = `(516,344)`, inside `x in [480,640)`. No other offset in the routine's paint list falls in that x-range. Sprite geometry, both roots: fighter 36x60/11 frames, mage 36x60/11 frames, tavern 64x96/10 frames, shopie 28x44/30 frames (`evidence/rom-column-refs.md` §5, `evidence/rightcol-scan.txt`)

**Confidence.** High (single routine, full decompile, every blit offset read directly; both roots agree on asset geometry)

### TOWN-091

The paint routine, read to line ~210 of an estimated 300+ (no string operands per `StrDump`), performs a zoom-shift-scaled icon size computation, a 16-bit color-channel composition from global width/shift constants, and a masked 2x2-block pixel copy/blend loop over a raw buffer the global at `L01168` (stride the global at `L01169`) — a downsample filter over live pixels, not a discrete-icon blit through the `vt+0x18`/`vt+0x34` primitives every other widget in this experiment's paint routine calls. Consistent with a rendered preview/thumbnail rather than a symbolic minimap; the source buffer's identity and on-screen role were not established this round (`evidence/rom-column-refs.md` §6) — *corrected by [EXP-0211]: the "not through the `vt+0x18`/`vt+0x34` primitives" clause is false.

Within the same first ~210 lines this row already read, the routine calls `[L01161]->vt+0x18` once, confirmed, at `L01162`; a second `vt+0x18` call follows at `L01163`, on a distinct object reached through `this+0x60`/`this+0x64` whose identity was not established; and, past line 210, `[L01161]->vt+0x20`/`vt+0x24` (`L01166`, `L01167`), confirmed; the masked pixel loop this row describes is additional to those calls, not a replacement for them (`AI-MINIMAP-156`). The routine is read to its own terminator, `RET` at `L01160`, 623 lines, and its remaining structure — an unidentified icon-array walk and two symmetric closing arms calling the shared rectangle-outline primitive — is also `AI-MINIMAP-156`.

Widget 5's functional identity, left open by this row, is closed by `AI-MINIMAP-157`: its own vtable at `L01176` also carries `AI-MINIMAP-062`'s cited left-down handler at `vt+0x54`, so widget 5 is the minimap `AI-MINIMAP-062`/`AI-MINIMAP-124` documents by its interactive behaviour. The masked-pixel-loop description is narrowed by `MISSION-064` (the unit is a fog block with three cases, not a fixed 2x2 block); the buffer-identity finding is not contested.*

**Confidence.** Medium (the routine's shape and its point of divergence from every sibling widget's blit-call pattern are read directly; the buffer's content and the widget's gameplay meaning are not established, so the "which widget" question is closed and the "what does it show" question is not) — *the "point of divergence" clause is retracted by [EXP-0211]'s correction above; the routine's overall shape and the buffer/pixel-loop portion remain High-supported by direct reading*

**Amended.** corrected. [`retracted.md`](retracted.md) holds an entry for this claim (refuted). `MISSION-063` and `MISSION-064` (EXP-0455) narrow this row's account of the paint loop: it is a three-case fog block loop, and the object walk calls every node's slot `+0x34`.

### TOWN-092

Gated on `campaign+0x3dc == 1`. 4 columns x 2 rows, pitch `0x22` (34px), cell origin `((i&3)*0x22+8,(i>>2)*0x22+7)`; each of 8 bits at `+0x64` skips its cell; a 9th field `+0x60` (signed, -1 = none) highlights one selected cell through a separate image pointer. Which asset backs each cell and what a slot represents in play was not traced past this routine (`evidence/rom-column-refs.md` §6)

**Confidence.** High (single routine, full decompile, gate and geometry read directly) for the widget's shape / Unknown for what each cell represents

### TOWN-093

Widget 8 (container-local rect (0,480,160,screenH), zero height at the shipped 640x480): constructor `R1937`, vtable `L06483`, own `vt+0x2c` `R0349`, a selected-map-object info readout gated on a session field and on its own rect being at least as tall as the borrowed panel's. Reads `campaign+0xd0` (map view) and `campaign+0xe0` (panel, `TOWN-086`/`SHOP-FIGURE-041`); first tests `[sess+0x3dc] == 1` (`L01386`, branching to `L01385` when not equal), a field-level match with `TOWN-092`'s widget-6 gate on the same offset. **Corrected by [EXP-0210] (second correction): a resolution-keyed background bitmap draws at both 800x600 and 1024x768; only the name/owner/ratio readout text is limited to 1024x768.** The routine is decoded whole, to its own terminator at `L01381`.

The height-comparison branch at `L08006` (a branch to `L08007` when less) and the fall-through from the readout path both land on one shared block at `L08007`, which reads `screenH` (`[L01259]`) and blits nothing at 640x480 (`screenH<=480`, `L11597`), `graphics\interface\extra800r.bmp` (global `[L11598]`, set at `L11599`) at 800x600 (`480<screenH<=600`), and `graphics\interface\extra1024r.bmp` (global `[L11600]`, set at `L11601`) at 1024x768 (`screenH>600`) — through the same `vt+0x18` blit primitive `TOWN-091` names for this widget family (`evidence/listing-R0349-widget8.txt`, `dwordscan-widget8-bitmaps.txt`, `listing-L11602-extrabmp-loaders.txt`, `bytes-widget8-bitmap-strings.txt`).

That blit is gated only by `[sess+0x3dc] == 1` (`L01386`), independent of the height comparison below it. Separately, the routine skips the name/owner/ratio readout text when its own client rect is shorter than the panel's (`L08005`..`L08006`): own height is `[rect+0x14]-[rect+0xc]`, panel height is `[panel+0x14]-[panel+0xc]`, and the panel's own construction (`SHOP-FIGURE-041`, `ctor(id 7, 0, 238, 160, 480)`) fixes that height at `480-238 = 242`. Own height tracks `screenH-480` — `0` at 640x480, `120` at 800x600, `288` at 1024x768, the resolution set `SESS-VIEW-028` measures — so the readout text draws only at 1024x768 (`288 >= 242`); at 800x600 the background bitmap draws but the readout text does not (`120 < 242`).

When the readout text draws, it hash-looks-up a map object by `*(mapview+0x990)`, RTTI-checks it against the `CStructure` name string (`R0214`), and on match draws two text lines plus a `"%d/%d"` pair from the object's own `+0xfc`/`+0x100` — a name/owner label and a numeric ratio; the fields' exact meaning was not traced past the format string (`evidence/rom-column-refs.md` §6)

**Confidence.** High (single routine decoded to its own terminator, both blit gates, both bitmap globals and the panel-height arithmetic read directly) for the shape and the visibility/blit conditions / Unknown for the `+0xfc`/`+0x100` field semantics

**Amended.** [`retracted.md`](retracted.md) holds an entry for this claim (refuted). `MENU-070`, `MENU-071` and `MENU-072` (EXP-0455) read the structure arm: the two lines are the `building.txt` name and the word `Health`, with no owner line in that arm, and the pair is current health over a maximum read from a class table (Medium). The non-structure arm `R0877` is unread, so an owner line there is not excluded.

### TOWN-094

Tavern: id `0x44e`, `R1939` (re-verified independently this round, matching `TOWN-064` exactly), own paint `R1904`, rect `(480,0,640,238)`. Shop: id `1006`, `R1733` (`SHOP-SCREEN-030`), own paint `R1027` (`SHOP-SCREEN-035`, not re-decompiled this round), rect `(464,0,640,238)`. School: id `0x3fd`, `R1940`, own paint `R1905`, rect `(464,0,640,238)`. Town: no discrete widget (`TOWN-089`). School's routine, read in full, shares tavern's two-part shape (one decoration blit, then an N-slot loop of sprite+digit pairs through the shared primitive `[L06186]+0x14`) but is a different routine with a different loop count (school: 2 iterations at `+0x94`..`+0xa8`; tavern: 3 at `+0x124`..`+0x130`) and different field offsets — not byte-identical.

Shop's routine (per `SHOP-SCREEN-035`) is structurally unlike either: four static rects, four numbers, a command jump table, no sprite loop (`evidence/rom-column-refs.md` §7-§8)

**Confidence.** High (all three routines' rects and construction sites read directly; the tavern/school shape comparison and shop/school rect equality are direct instruction-level reads)

### TOWN-095

`EnumRefs vt:L03649:36` (shop's vtable): `+0x2c` = `R1311` directly, no blit call of its own. Shop's own `+0x30` is `R1769` (`SHOP-FIGURE-041`'s fill routine). Of the four surfaces, shop and tavern draw entirely through constructed children; town and school also draw directly from their own `+0x2c` (`R1489`, `TOWN-061`) (`evidence/rom-column-refs.md` §9)

**Confidence.** High (vtable slot read directly and compared against the three already-published `+0x2c` slots)

### TOWN-096

The borrow offset `0x280` (640): 6 sites (tavern/school enter+leave here, shop enter+leave, `SHOP-FIGURE-041`). Column width `0xa0` (160) and children heights `0x9e`/`0xee`/`(screenH-0x1e0)` (158/80/242/variable): `R0315`. Widget 6's grid: pitch `0x22` (34px), 4x2 = 8 cells: `R0316` — a 9th cell needs a code change, not a data change. Upper-right block rects `(480,0,640,238)`/`(464,0,640,238)`: `R1939`/`R1733`/`R1940`. The one data-class element found: town's paint anchors `(0,0)` and `(0x204,0x158)` in `R1489` are engine-class, but the pictures drawn there (`townmain.bmp`, four `townbirds/<class>/sprites.16a`) are ordinary shipped files, so their *content* is data-class even though their placement is not.

Negative result scoped to the routines `EXP-0186` §1-§9 read; no whole-registry key search against these specific fields was run

**Confidence.** High (every geometry value cited is a literal read directly in the routine that uses it) for the engine-class immediates found / this is not an exhaustive search of the whole `.reg` catalog for a counterexample

## World-map route graph and travel

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-116 | `graphics\Global.Map\PathMap.bmp` is a 640x480 indexed route-topology mask, not visible map art. Index 1 is corridor, index 2 is graph node, and every other sampled value is impassable. | High | ● active | [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-117 | World-map enter unconditionally precomputes an all-node route-segment matrix from `PathMap.bmp`, transfers it into the view, and destroys the bitmap-bearing builder before activation; the bitmap is never blitted. | High | ● active | [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-118 | A mission is selected by clicking its scroll entry, not by clicking a map region. | High | ● active | [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-119 | The chosen travel route is the chain of precomputed segments whose accumulated segment point count is least over the PathMap node graph, expanded back into a pixel-coordinate list. | High | ● active (amended, superseded) | [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-120 | World-map paint composes travel in this order: `GMap.bmp`; cached mission markers; `BallMap.bmp` stamps at every eighth revealed route coordinate; animated destination `Cross`; hovered-scroll `Flag1` while at town… | High | ● active | [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-121 | Travel completes before the destination opens, and it is skippable after it starts. | High / Unknown | ● active | [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-122 | Each offered mission is a three-part scroll card with state-selected normal/pressed art, a title, wrapped map briefing, and an optional payment line. | High | ● active | [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-123 | The map-rectangle gate's campaign list is the same persisted 0x14-byte marker cache world-map paint draws, not the sub-mission offer array. | High | ● active | [EXP-0188](../experiments/EXP-0188-pathmap/) |
| TOWN-124 | World-map route customization is partly data-class and partly engine-class, and its save boundary separates campaign selection art from transient route presentation. | High / Unknown | ● active | [EXP-0188](../experiments/EXP-0188-pathmap/) |

### TOWN-116

`R1531` scans the whole loaded pixel plane and records every byte equal to 2; `R1532` recursively traverses eight-neighbour bytes equal to 1 and stops a branch when it reaches a different 2. Corpus, live/EN/RU byte-identical: 308278 bytes, SHA-256 `c2b9357a647c36bcd253b70b7e609d2baca2e6c9e2db2ba5e95660bf789d2b86`, 304042 index-0, 3077 index-1, 81 index-2 pixels. All 35/35 `globalmap.reg` MapPoints land on an index-2 pixel; the other 46 graph nodes have no MapObject, and their individual semantic roles are Unknown

**Confidence.** High for the operational index meanings (full consuming control flow and raw instructions) and dimensions/counts (committed tool, source-run and built-tool outputs identical); EN/RU agreement alone is not the grade

### TOWN-117

In `R1409`, allocation and `R1531` construction occur before the current-position branch that may build mission scrolls. `R1943` starts a walk from every discovered node and stores both directions' point lists, `R1533` transfers matrix and node coordinates to view `+0x148/+0x150/+0x154`, then `R1944` and delete tear the builder down before `R0775` activates the view. Full paint `R1530` has no read of the bitmap object and no PathMap draw

**Confidence.** High (enter control flow, allocation, scan, transfer, destruction, activation order, and full paint read directly; the sole PathMap literal reference is enumerated)

### TOWN-118

`R1409` enumerates announced campaign missions into 0x4c-byte scroll entries, mission number at `+0x28` and zero-based MapObject index at `+0x2c`. Click override `R1945` hit-tests the entry rectangles through `R1946`, calls `R0785(mission)`, resolves the destination MapPoint, stores it at view `+0x100/+0x104`, and starts the route. The separate MapRect override `R1863` only linearly `PtInRect`-tests a region and returns an indexed text-table value; it calls neither mission selection nor route construction

**Confidence.** High (both paths fully read; raw `L11603`..`L11604` fixes the selection call and destination stores, while the negative is scoped to the complete MapRect routine)

### TOWN-119

`R1947` maps current and destination coordinates to node indices; `R1493` recursively visits nodes not already in the candidate chain, retains a candidate only strictly below the best accumulated point total, and records the best node-index chain; `R1947` then concatenates the corresponding matrix point lists into view `+0xd8`, count `+0xdc`. The route is not a straight interpolation between endpoints and not a path recomputed over visible background pixels on every frame

**Confidence.** High (solver, bound, recorded chain, and concatenation read end to end; the competing straight-line model is excluded by the matrix reads)

**Amended.** amended — `EXP-0194` corrected the criterion's wording: the accumulated quantity is the segment's own point count and never the number of segments; `TOWN-141`..`TOWN-143`. [`retracted.md`](retracted.md) holds an entry for this claim (superseded).

### TOWN-120

World-map paint composes travel in this order: `GMap.bmp`; cached mission markers; `BallMap.bmp` stamps at every eighth revealed route coordinate; animated destination `Cross`; hovered-scroll `Flag1` while at town; animated current-position `Flag`; then scroll cards. `R1530` is world-map vtable `L11605+0x2c` and contains every listed consumer in that order. `PathMap.bmp` and `Hero.bmp` are not read by this own-paint routine. Route progress increases by 8 coordinates per paint, while Cross/Flag counters advance one frame per paint. This is spatial/presentation progress, not `TOWN-045`'s absent campaign aggregate

**Confidence.** High (full paint and corrected vtable slot read directly; draw-order claim is instruction order, not inferred from asset names)

### TOWN-121

Only after route progress and the Cross animation have both passed their ends does paint copy destination to current, clear the route, resolve the current MapPoint, and post message `0x467` with its MapObject index. `R0701` routes index 0 to `R1320` (town, which starts `music\Town.wav`) and non-zero to `R0099(0)`. The latter reads selected mission at screen `+0x660` (campaign record `+0x548+0x118`), formats `%d.alm`, passes it to the map load path through `R0473`/`R0512`, and starts the loaded session through `R0545`. A click that misses a scroll calls `R1948`; if route progress is non-zero it sets both completion counters past their ends.

The known paint driver is world-map `+0x48` `R1949`: on message `0x402`, once more than 99 ms have elapsed, it calls shared `+0x34` `R0366`, which gates and calls own-paint `+0x2c`. Thus a successful timer-message arm advances 8 route coordinates and cannot do so more often than that gate; the arrival cadence of `0x402` and whether other paths paint remain Unknown

**Confidence.** High for ordering, message payload, destination arms, selected mission to map-load chain, skip, and timer-to-paint dispatch (raw instructions, correct vtable, and both handlers) / Unknown for exclusive wall-clock paint cadence

### TOWN-122

`R1409` reads `main\text\battle\m<n>\title.txt` and `briefmap.txt`, formats non-zero payment, and lays out one entry; `R1950` selects `Scroll01/02/03` or `ScrollP1/P2/P3` from the entry state and draws the stored text. The list exists only while current position equals the first MapPoint, and `R1946` uses its same stored rectangles for hover and click

**Confidence.** High (builder, renderer, and hit-test read directly; `TOWN-044` independently censuses all shipped title/briefing files)

### TOWN-123

The campaign object is addressed at `record+0x548`; `R1325` writes marker cache `+0x15c/+0x160`, hence absolute campaign `+0x6a4/+0x6a8`, and `R1530` iterates those exact absolute fields. Mission selection memoizes by mission number and appends only for a non-`"nothing"` MapObject `Picture`. Campaign serializer `R1294` writes the `+0x160` count and each record through `R1951`; loader `R0434` restores the count and records through `R1952`; `R1490` rehydrates each marker picture through `R1953`. This identity and persistence path correct `TOWN-040`'s former “currently-offered missions” reading

**Confidence.** High (writer, object-relative arithmetic, paint reader, serializer, loader, rehydration loop, and element stride agree; an offer-list model predicts `record+0x48`, which these readers never access)

### TOWN-124

Corridor/node topology and marker coordinates are shipped bytes: changing index-1/index-2 pixels changes `PathMap.bmp`, and every destination MapPoint must coincide with an index-2 node. Node count is discovered from data (81 shipped), not a fixed maximum. Bitmap extent and coordinate basis are engine-class: `R1531` scans literal `0x280 × 0x1e0` and `R1532` indexes rows with literal width `0x280`, so a different map extent needs code changes as well as new files. The campaign marker cache and selected-mission scalar are serialized, loaded, and marker art is rehydrated (`TOWN-123`); the expanded route, destination, progress, and animation counters are world-map-view fields, and no serializer for those view fields was identified. Whether the game exposes saving during travel, and whether such a route would resume or rebuild, remains Unknown

**Confidence.** High for the data/engine split and campaign marker persistence (literal loops, data census, and save/load consumers) / Unknown for user-visible mid-travel save/load, which was not dynamically reachable or observed

## School skill messages and view message tables

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-GENERAL-106 | The school has five selectable skills per class and no General selection. | High | ● active | [EXP-0187](../experiments/EXP-0187-general-skill/) |
| TOWN-GENERAL-107 | Selecting a school skill sends opcode `0x40` and displays the server-computed price. | High | ● active (amended) | [EXP-0187](../experiments/EXP-0187-general-skill/) |
| TOWN-GENERAL-108 | Train sends opcode `0x3d` only for the selected one of those five slots. | High | ● active (amended) | [EXP-0187](../experiments/EXP-0187-general-skill/) |
| TOWN-136 | The tavern view's message table is the compiled `switch` in `R1808`, and it carries non-default entries for both party-picker ids. | High | ● active | [EXP-0193](../experiments/EXP-0193-picker-message-tables/) |
| TOWN-137 | The school view's message table is the compiled `switch` in `R0735`, and it carries non-default entries for both party-picker ids. | High | ● active | [EXP-0193](../experiments/EXP-0193-picker-message-tables/) |
| TOWN-138 | The picker step is implemented three times, once per view, over three different rosters; no routine is shared between the shop, the tavern and the school. | High | ● active | [EXP-0193](../experiments/EXP-0193-picker-message-tables/) |
| TOWN-139 | A view's own message table does not consume the message: every arm, matched or default, forwards it down the same three-routine chain, and an id no arm matches is offered to the container's children and then dropped. | High | ● active | [EXP-0193](../experiments/EXP-0193-picker-message-tables/) |
| TOWN-140 | Each picker id is pushed at exactly one site in `.text`, both to the root container, and the panel's gate admits four screens: the shop, the tavern, the school, and a fourth view whose table carries the same two arms. | High / Unknown | ● active | [EXP-0193](../experiments/EXP-0193-picker-message-tables/) |

### TOWN-GENERAL-106

`R1862` maps visual indices 0..4 to stored slots `1,2,4,3,5` on both class branches. `R1954` stores that result at school `+0x34c`; no valid result is 0

**Confidence.** High (both mapping branches and every valid input are read at instruction level)

### TOWN-GENERAL-107

`R1955` writes actor and slot to packet `+0x0a/+0x0e`; the server accepts slots 0..5 and computes `ftol(1.1^word[actor+0x116+2*slot] * 200)`. For selectable slots 1..5 this is their maintained base; for abstract slot 0 it is a serialized shadow that every base-maintenance loop skips (`HERO-GENERAL-086`). Reply `0x84` reaches `R1956`, which writes school `+0x354`. `R1957` writes purse and `purse-price` to `+0x35c/+0x358`

**Confidence.** High (continuous producer, dispatcher, reply and UI-state chain, with both raw price reads distinguished from the base loops)

**Amended.** corrected after independent falsification.

### TOWN-GENERAL-108

`R1548` refuses when `purse-price < 0` or price < 1, then calls the sole opcode builder `R1932`. The server operation accepts 0..5, but the shipped school producer can supply only `1,2,4,3,5`. General is abstractly server-purchasable but not trainable through the shipped school UI; if another producer supplies 0, both query and purchase price it from the decoupled `+0x116` word rather than live General

**Confidence.** High (the only direct callers of both school builders, the mapping, both packet layouts and both server arms were read; pointer-table reachability is not used for the uniqueness clause)

**Amended.** corrected after independent falsification.

### TOWN-136

The class is identified by its vtable, `L09234`, whose `+0x80` and `+0x84` slots are `R1412` and `R0786` — the tavern enter and leave routines `TOWN-087` reads — and whose `+0x48` slot is `R1808`. The dispatch at `L11606` is an indirect jump through the dword table at `L11607` indexed by the byte read from the index array (4 bytes per entry), over the byte index array at `L11608`, rebased by subtracting 0x402 from the message id at `L11609` and bounded by a compare of the rebased index with 0x58 at `L11610`: 89 entries, ids `0x402..0x45a`, default arm `L11611` carrying 85 of them. Four ids reach another arm: `0x402` to `L11612`, `0x414` to `L11613`, `0x415` to `L11614`, `0x45a` to `L11615`. The arms for `0x414` and `0x415` call `R1921` and `R1920`.

Table decoded from raw PE bytes by `tools/townmsg -mode table` with no disassembler in the path, identical on both roots (`evidence/tables-decoded.md`, `evidence/vtable-slots.md`, `evidence/rom-dispatch-excerpt.md` §3)

**Confidence.** High (table decoded from bytes and handler read at instruction level; the class identification discriminates, since a different routine pair at `+0x80`/`+0x84` would have named a different view. Both roots share one `.text` payload SHA-256, so their agreement is not counted)

### TOWN-137

The class is identified by its vtable, `L03653`, whose `+0x80` and `+0x84` slots are `R0706` and `R1915` — the school enter and leave routines `TOWN-088` reads — and whose `+0x48` slot is `R0735`. The dispatch at `L11616` is an indirect jump through the dword table at `L11617` indexed by the byte read from the index array (4 bytes per entry), over the byte index array at `L11618`, rebased by subtracting 0x402 from the message id at `L11619` and bounded by a compare of the rebased index with 0x58 at `L11620`: 89 entries, ids `0x402..0x45a`, default arm `L11621` carrying 84 of them. Five ids reach another arm: `0x402` to `L11622`, `0x414` to `L11623`, `0x415` to `L11624`, `0x43f` to `L11625`, `0x45a` to `L11626`. The arms for `0x414` and `0x415` call `R1958` and `R1959`.

Same instrument and same cross-root result as `TOWN-136` (`evidence/tables-decoded.md`, `evidence/vtable-slots.md`, `evidence/rom-dispatch-excerpt.md` §4)

**Confidence.** High (same grounds as `TOWN-136`)

### TOWN-138

The six arm addresses are pairwise distinct: shop `L11627`/`L11628`, whose `CALL rel32` targets decode from raw bytes to `R1766` and `R1765`, reproducing `SHOP-PICKER-043`'s mapping, tavern `L11613`/`L11614` (`R1921`/`R1920`), school `L11623`/`L11624` (`R1958`/`R1959`). The tavern steps a 32-bit index `view+0xbc` over the pointer array `view+0xd8` with count `view+0xdc`; the school steps a 16-bit index `view+0x33c` over the global pointer array `[L11629]` with count `[L11630]`; the shop steps a 16-bit index `view+0x130` over the `CArray` at `view+0x108` (`SHOP-PICKER-043`). All six share four steps — post `0x405` to the figure widget, move the index with wrap at both ends, call the new element's `vt+0x14(1)`, call `R0082` on the figure widget — and differ in what else they reset.

The tavern resets nothing further. The school writes `-1` into the five dwords at `view+0x2a4`, the five at `view+0x2f4` and `view+0x34c`, writes `0` into `view+0x354`, stores the result of the import at `[L00849]` in `view+0x340`, sets bit 3 of the new element's `+0x18c`, and sets `view+0x348` to 1 when bit 1 of `+0x18c` differs between the old and the new element. The shop rebinds the backpack grid (`SHOP-PICKER-043`). `view+0x34c` and `view+0x354` are the selected-slot and price fields `TOWN-GENERAL-107` and `TOWN-GENERAL-108` read, so a step on the school screen discards a pending skill selection and its quoted price (`evidence/rom-dispatch-excerpt.md` §6, §8)

**Confidence.** High (all four tavern and school routines read end to end at instruction level, the six arm addresses come from the decoded tables rather than from any name, and the shop's two call targets are decoded from the arm bytes rather than cited)

### TOWN-139

In `R1808` the arms for `0x414` and `0x415` are two-instruction stubs that call one routine and then jump to `L11611`, which is also the guard's `JA` target, that is the default arm; it pushes the three original arguments and calls `R0716`. `R0735` and `R1960` give each arm its own copy of the same tail. `R0716` handles only `0x445` and `0x446`, and otherwise calls `R0390` at `L11631`. `R0390` handles only `0x100..0x102`, `0x200..0x206` and `0x400`, and otherwise calls `R0642` at `L11632`, which walks the container's child collection and calls each child's own `vt+0x48`; if that returns 0 the routine enters a second stage at `L11633` which tests the same id set again and falls to the return at `L11634` with the result still 0.

There is no further table and no parent-class table below this chain. The negative is scoped to these three routines, each read in full (`evidence/rom-dispatch-excerpt.md` §7)

**Confidence.** High (all three routines read end to end, both stages of `R0390` included; the "no further table" clause is scoped to the routines actually read)

### TOWN-140

A byte-level scan of `.text` for the 4-byte values, with no disassembler in the path, finds `0x414` at 25 offsets and `0x415` at 4; exactly one of each is a push of a 32-bit immediate (the encoded push opcode followed by the immediate, at `L09801` and `L09802`), both in `R0332` and each followed by an indirect call through slot +0x48 of the object's table, on `[campaign+0xcc]`. Every other hit is a struct displacement `[reg+0x414]`, an add or subtract of 0x414 on the stack pointer, a `rel32` displacement, or the interior of an encoded move of the 32-bit immediate 0x4 to a stack local. `[campaign+0xcc]` is a `0x5c`-byte object built by `R0773` (vtable `L03640`, stored at `L11635`), whose `+0x48` is `R0390` — the container handler of `TOWN-139`, which broadcasts both ids unchanged.

The gate at `L03383` tests the 32-bit field `campaign+0x3dc` against the mask 0x226; the four bits are set by `R1315` (`0x2`, view `campaign+0xf0`), `R1318` (`0x4`, `+0xfc`), `R1319` (`0x20`, `+0x104`) and `R1321` (`0x200`, `+0x358`), each of which adds its view to `[campaign+0xcc]` and then calls that view's `vt+0x80`. The fourth view's class vtable is `R1312` (constructor `R1314`, called from the campaign screen constructor `R0315` at `L06735`); its `+0x48` is `R1960`, an 89-entry table over `0x402..0x45a` routing `0x414` to `L11636` and `R1674`, and `0x415` to `L11637` and `R1673`, both of which step `view+0xfc` modulo one of four count fields and return without acting when `[R0347()->vt+0x7c() + 0x6bc] == 2`.

Which screen that fourth view presents is Unknown: its class carries no `CRuntimeClass`, so the image holds no name for it (`evidence/immediate-census.md`, `evidence/dispatch-census.md`, `evidence/rom-dispatch-excerpt.md` §1, §2, §5)

**Confidence.** High for the two posting sites, the receiving object, the gate mask and the fourth table (byte-level census whose false-match classes are printed and classified, plus instruction-level reads of every routine named) / Unknown for which screen the fourth view is

## World-map route criterion

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-141 | The world-map chain solver accumulates the number of points in each precomputed segment, not the number of segments. | High | ● active | [EXP-0194](../experiments/EXP-0194-mappoint-chain-criterion/) |
| TOWN-142 | One running best at world-map view `+0x12c` governs both discards, and both comparisons are strictly-less. | High | ● active | [EXP-0194](../experiments/EXP-0194-mappoint-chain-criterion/) |
| TOWN-143 | The number of segments in a chain is computed but never compared, and the winning accumulated value is stored where nothing reads it. | High | ● active | [EXP-0194](../experiments/EXP-0194-mappoint-chain-criterion/) |
| TOWN-144 | Equal accumulated values are resolved by exploration order, and that tie is not reached by any shipped MapPoint pair. | High / Medium | ● active | [EXP-0194](../experiments/EXP-0194-mappoint-chain-criterion/) |
| TOWN-145 | Reading the criterion as a segment count rather than a point count changes the chosen route for 217 of the 1190 ordered shipped MapPoint pairs. | Medium | ● active | [EXP-0194](../experiments/EXP-0194-mappoint-chain-criterion/) |

### TOWN-141

Per extension `R1493` loads `[matrix[from][b] + 0x8]` and adds it to the candidate chain's own field `+0x14` at `L11638`..`L11639`; the same sum is formed for the pruning test at `L11640`..`L11641`. `[segment + 0x8]` is the segment's element count: `R1532` raises it by one per appended `(x,y)` point at `L11642`..`L11643`, `R1947` copies exactly that many 8-byte elements out of `[segment + 0x4]` into the route at `L11644`..`L11645`, and the class's own `SetSize` `R1961` allocates `count*8` bytes with the count at `+0x8`. The accumulated total is therefore the point count of the pixel route the caller builds.

It is neither a hop count nor a geometric distance: `R1493` contains no coordinate arithmetic, no multiplication and no FPU instruction. A segment's point list carries both endpoint nodes, so the total equals the corridor pixel count plus twice the number of hops

**Confidence.** High (the accumulating instructions, the object layout taken from the routines that allocate and fill it, and the caller's own use of the same field, all read in raw listings; the hop-count and geometric models are each refuted by a named instruction)

### TOWN-142

`R1947` initialises it to `0x77359400` at `L11646`. Inside the neighbour loop `R1493` forms `chain[+0x14] + segment[+0x8]` and takes `JGE` to the next neighbour at `L11647`..`L11648`, so a candidate is pruned unless strictly below. At the terminal test the completed chain's `+0x14` is compared with the same field and `JGE` returns at `L11649`..`L11650`; otherwise `+0x12c` is overwritten at `L11651` and the chain's element array is copied to `+0x130`. The per-hop cost is positive, so the pruning bound is admissible and the retained chain is a true minimum

**Confidence.** High (both comparisons and the initialisation read in raw listings; the field's accesses enumerated with `EnumRefs disp:12c` and `imm:12c`, 73 and 40 whole-image hits, of which four fall in the world-map class and all four are in these two routines; blind only to a wholesale copy of the containing object)

### TOWN-143

The candidate chain's element count `[chain + 0x8]` is read once, at `L07793`, as the index at which `R1493` appends the next node. The retained chain's element count at view `+0x138` is read twice, at `L07794` and `L07795`, both times in `R1947` as the bound of the concatenation loop. Neither is compared against `+0x12c` or against any other running value. The winning accumulated value is written to view `+0x144` by three instructions — `L11652` in the view constructor, `L11653` in the route builder and `L11654` in the solver — and no instruction reads it

**Confidence.** High (both routines read whole in raw listings; enumerated with `EnumRefs disp:138`, `disp:144`, `disp:130`, `imm:138` and `imm:144` together, because a displacement sweep at `0x144` cannot see the two writes that reach the field as `[base + 0x14]` off `+0x130`; blind only to a wholesale copy of the containing object)

### TOWN-144

Both comparisons are strictly-less, so the first chain to reach a given accumulated value is retained and every later chain of the same value is discarded. `R1493` tries neighbours by ascending node index — the neighbour index zeroed at `L11655` and incremented at `L11656` — and recurses depth-first, so the retained chain among equals is the first in that order. No secondary comparison exists in either routine. Over the 81-node graph reproduced from the shipped `PathMap.bmp`, none of the 1190 ordered pairs of the 35 `globalmap.reg` MapPoints has more than one minimum-cost chain, on the live install and both preserved roots

**Confidence.** High for the resolution rule (the two comparisons and the neighbour loop read in raw listings); Medium for the shipped-corpus uniqueness, which is measured on `tools/roadchain`'s reproduction of `R1531`, `R1943` and `R1532` rather than on output taken from the original

### TOWN-145

Same reproduced graph and same solver, run with the executable's per-hop cost at `[segment + 0x8]` and with a constant cost of 1 per hop: the chains differ for 217 ordered pairs, which is 110 of the 595 unordered pairs, 107 of them in both directions and 3 in one. The segment-count reading never produces a shorter route — its routes are longer by 6 to 384 points, median 119, 35046 in total. It also makes the tie-break load-bearing, since 230 ordered pairs then have more than one minimum, and for 18 ordered pairs it selects one of the two shipped segments whose point list contains a coordinate discontinuity, which the executable's own cost never selects. An independent Dijkstra agrees with the reproduced solver's minimum on all 1190 pairs under both costs

**Confidence.** Medium (measured over the whole shipped corpus on the live install and both preserved roots with byte-identical results, and internally checked — 3076 of 3077 index-1 pixels covered, no pixel claimed twice, 35 of 35 MapPoints on nodes — but the graph is a reproduction of three named routines rather than output taken from the original)

## School column pictures, masks and skill icons

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-146 | The school's step-4 picture is one frame of a 16-frame column, held in the pointer array `view+0x30c` with element count `+0x310` and capacity `+0x314`, and indexed directly by `view+0x31c` with no bounds test. | High | ● active | [EXP-0195](../experiments/EXP-0195-school-column/) |
| TOWN-147 | `view+0x31c` is one field with two roles: the column frame index `0..15` and the panel selector, and it takes every value in that range. | High | ● active | [EXP-0195](../experiments/EXP-0195-school-column/) |
| TOWN-148 | The school has two skill-icon rectangle arrays, one per class, not one shared array. The mage array is `+0x10c`..`+0x14c` and the fighter array is `+0x1c0`..`+0x200`, each preceded by its own panel rectangle at `+0xfc` and `+0x1b0`. | High | ● active | [EXP-0195](../experiments/EXP-0195-school-column/) |
| TOWN-149 | `view+0x31c == 0` is the fighter panel and `== 0xf` is the mage panel, established four independent ways, one of which uses no filename. | High | ● active | [EXP-0195](../experiments/EXP-0195-school-column/) |
| TOWN-150 | The column picture and the skill icons use different blit primitives: the column is copied opaque and each icon is keyed on source value 0. | High | ● active | [EXP-0195](../experiments/EXP-0195-school-column/) |
| TOWN-151 | Each class has fifteen icon pointers, three art states per slot, and a five-dword state array whose value selects the state. | High / Medium | ● active | [EXP-0195](../experiments/EXP-0195-school-column/) |
| TOWN-152 | The school indexes its per-slot data in two different orders, and they differ by the permutation `0,1,3,2,4`. | High | ● active (amended) | [EXP-0195](../experiments/EXP-0195-school-column/) |
| TOWN-153 | The two column masks are read and never drawn, and the 8-bit loader reverses their rows, so the hit test's buffer row 0 is the panel's top screen row. | High | ● active | [EXP-0195](../experiments/EXP-0195-school-column/) |
| TOWN-154 | All 63 archive entries under `interface/training/` are reached from the executable; none is unreferenced. | High | ● active | [EXP-0195](../experiments/EXP-0195-school-column/) |
| TOWN-155 | The number of column frames cannot be changed by shipped data alone: exactly one site derives it from the loaded count and the rest are `.text` literals. | High | ● active | [EXP-0195](../experiments/EXP-0195-school-column/) |

### TOWN-146

`R1486` formats `graphics\interface\training\column\rt%.4d.bmp`, loads each result through the 24bpp surface loader `R1176` and stores it at 4 bytes times the loop index in the array whose pointer is at offset 0x30c of the object (`L11657`); the loop bound is the literal compare of the loop index with 0x10 at `L11658`, and a load failure stores 0. Both roots ship `rt0000.bmp`..`rt0015.bmp`, 16 entries, each 148x208 24bpp. The paint reads `[[+0x30c] + [+0x31c]*4]` at `L11659`..`L11660`, skips a null, and draws the whole surface at the view-relative point `(x+0xa8, y+0xb0)` = `(168,176)`; both displacements are `.text` immediates, so the destination is a code constant and not a value from a shipped file (`evidence/disasm-R1487-paint-R1940-ctor.txt`, `evidence/disasm-R1486-loader.txt`, `evidence/training-art-en.csv`, `evidence/training-art-ru.csv`)

**Confidence.** High (loader loop, array fields, paint indexing and the destination immediates all read at instruction level; the 16 entries counted in both shipped archives)

### TOWN-147

All 19 instructions in the image that reference the displacement lie inside the school's address range. Four are writes: `L11661` writes `0xf` and `L11662` writes `0` in `R0706`, chosen by a test of the byte at offset 0x18c against 0x2 on the selected roster element; `L11663` writes `7` in the loader `R1486` (the value `7` is set in a callee-saved register at `L11664`, preserved across the intervening calls, and not written again before the store); `L11665` adds the step `+0x320`, which holds `-1`, `0` or `+1`. The step runs at most once per paint call, gated on `timeGetTime()` — the import at `[L00849]` — minus the global at `L11666` exceeding `0x53`, that is 83 ms, and stops itself by writing `0` into `+0x320` at value `0xf` with step `+1` and at value `0` with step `-1`.

`R1962` starts the `+1` rotation and `R1963` the `-1`; each requires `+0x344` to agree with its direction, plays `SFX\Town\School\Rotate.wav` and writes `-1` into `+0x34c` and `0` into `+0x354` (`evidence/disasm-R0706-R1962-R1963-R1971.txt`)

**Confidence.** High (every reference to the displacement enumerated image-wide, every write read at instruction level, and the timer named from the import directory rather than inferred)

### TOWN-148

`R1940` writes 80 object-relative dwords in one basic block; grouped into 16-byte strides without assuming a base, they form exactly two contiguous runs of six rectangles: `+0x0fc` (188,188,288,308), `+0x10c` (264,232,284,260), `+0x11c` (192,240,216,260), `+0x12c` (228,272,256,298), `+0x13c` (224,200,252,224), `+0x14c` (224,236,256,262), and `+0x1b0` (192,192,284,312), `+0x1c0` (200,196,280,228), `+0x1d0` (200,216,280,252), `+0x1e0` (200,272,280,288), `+0x1f0` (200,248,280,276), `+0x200` (200,288,280,308). The paint's icon loop starts its walking pointer at the view's offset `+0x1cc` and steps by `0x10`; the `+0x31c == 0` branch at `L07788` reads the dwords at the pointer, pointer-4, pointer-8 and pointer-0xc and the `+0x31c == 0xf` branch at `L07789` reads the dwords at pointer-0xb4, pointer-0xb8, pointer-0xbc and pointer-0xc0, which are the same four fields of the rectangle 0xb4 bytes lower.

Each branch also reads its own icon pointer array. This refutes `TOWN-068`'s shared-array clause; its twelve values are correct and are the fighter panel's alone (`evidence/rect-runs.txt`, `evidence/disasm-R1487-paint-R1940-ctor.txt`)

**Confidence.** High (the store enumeration assumes no base and reports the runs it finds; both branches read at instruction level)

### TOWN-149

(1) The loader's own filename literals paired with their store offsets: `column/fighter/mask.bmp` into `+0x1ac`, the mask field used with panel rectangle `+0x1b0`; `column/mage/mask.bmp` into `+0xf8`, used with `+0xfc`; the fifteen `column/fighter/*` icons into `+0x268`..`+0x2a0` and the fifteen `column/mage/*` into `+0x2b8`..`+0x2f0`. (2) Rectangle sizes against shipped picture sizes: the `+0x1c0` run is 80x32, 80x36, 80x16, 80x28, 80x20 and the fighter art is `sword` 80x32, `axe` 80x36, `pike` 80x16, `club` 80x28, `bow` 80x20 in load order; the `+0x10c` run is 20x28, 24x20, 28x26, 28x24, 32x26 and the mage art is `fire` 20x28, `water` 24x20, `earth` 28x26, `air` 28x24, `astral` 32x26 in load order.

Every fighter rectangle is 80 wide and no mage picture exceeds 32, so no cross-class assignment fits. (3) `R0706` writes `0xf` when bit 1 of the selected hero's `+0x18c` is set and `0` when clear, and `R1964` selects state array `+0x2f4` or `+0x2a4` by the same bit. (4) The mask geometry of `TOWN-153`. This confirms `TOWN-067` and corrects `TOWN-061`, whose step-5 clause labels `+0x264` the mage array and `+0x2b4` the fighter array; the two offsets are attached to the correct `+0x31c` values and the class names are the wrong way round (`evidence/loader-slots.csv`, `evidence/training-art-en.csv`, `evidence/training-art-ru.csv`)

**Confidence.** High (four readings that do not share a premise, three of them at instruction level and one from the shipped pictures alone)

### TOWN-150

The bitmap vtable at `L11667` has four blit entries. `+0x18` (`R1965`) and `+0x34` (`R1425`) call `R0797`, which copies every row with repeated dword moves and repeated byte moves and writes every source pixel. `+0x38` (`R1426`) calls `R0798`, whose inner loop loads the 16-bit source word, tests it, skips the pixel when it is zero, and otherwise stores that word to the destination: source value `0x0000` is not written, the key is the source value itself, and no separate mask surface is read. `+0x3c` (`R1966`) calls the additive `R1967` and the school does not use it. Each wrapper reads the 16-bit surface at `[this+0x10]` (width at `[buf]`, height at `[buf+4]`, pixels at `buf+8`) and forwards nine arguments: `destX, destY, srcLeft, srcTop, srcRight, srcBottom, pixels, width, height`.

The school paints `+0x70`, `+0xf4`, `+0x1a8`, the column picture and `+0x25c` through `+0x18`, and all ten skill icons and steps 7 and 8 through `+0x38`. Every icon blit passes the destination rectangle's own top-left as the destination point and `(0, 0, right-left, bottom-top)` as the source rectangle (`evidence/disasm-blit-wrappers.txt`, `evidence/disasm-blit-primitives.txt`)

**Confidence.** High (both wrappers read to their `CALL`, both primitives' inner loops read, and the argument construction read at the call site)

### TOWN-151

The fighter pointers occupy `+0x268`..`+0x2a0` and the mage pointers `+0x2b8`..`+0x2f0`, written by the loader in the order `on`, `shine`, `shine_on` per slot. The paint indexes them as `[view + (state + 3*slot)*4 + base]` with base `0x264` for the fighter and `0x2b4` for the mage — four bytes below the first pointer, the compiler having folded the state value's bias of one into the displacement. The state values are the five dwords at `+0x2a4` (fighter) and `+0x2f4` (mage); `-1` draws nothing and is what the roster picker writes (`TOWN-138`). `R1968` writes `[this + slot*4 + 0x2a4]` or `[this + slot*4 + 0x2f4]` by the same `+0x31c` test, setting bit 1 on the slot it is given and clearing bit 0 on the other four; `R1964` tests bit 1 across all five (`evidence/disasm-R1487-paint-R1940-ctor.txt`, `evidence/disasm-R1477-hittest.txt`, `evidence/loader-slots.csv`)

**Confidence.** High for the addressing, the array bases and the writers (all read at instruction level) / Medium for state `1` meaning `on`, `2` meaning `shine` and `3` meaning `shine_on`, which follows from the loader's file order and from which bit each writer sets, not from a site that names a file

### TOWN-152

The rectangle, icon pointer and state arrays are indexed by the stored slot, which is the loader's file order: `sword, axe, pike, club, bow` and `fire, water, earth, air, astral`. The per-slot sample fields at `+0x84`..`+0x94` and `+0x98`..`+0xa8` are indexed by the other order. Both hit-test dispatches are byte tables indexed by `maskCode - 0x37` into five-way jump tables at `L11668` and `L11669`; each arm passes a slot immediate to `R1968`, which writes the state array at that index, and tests one sample field. Decoded from the image: mage codes `0x87, 0x37, 0xd2, 0xff, 0x9e` give slots `0, 1, 2, 3, 4` against flag indices `0, 1, 3, 2, 4`; fighter codes `0xff, 0x9e, 0xd2, 0x87, 0x37` give slots `0, 1, 3, 2, 4` against flag indices `0, 1, 2, 3, 4`.

The fighter rectangles sorted by top edge are slots `0, 1, 3, 2, 4` (y196, y216, y248, y272, y288), so for that panel the flag order is the top-to-bottom visual order and the array order is the load order. This is the divergence `TOWN-068` left open. The permutation is the same one `TOWN-GENERAL-106` records for `R1862`, established here from the hit test's own tables. The mage rectangles are scattered across x 192..284 and have no single visual order (`evidence/hittest-arms.txt`, `evidence/rect-runs.txt`)

**Confidence.** High (both jump tables decoded from the image, and each arm's slot immediate and flag displacement read from its own bytes rather than from a listing)

**Amended.** The `enabled flag` wording for the per-slot fields, corrected in the text above (`TOWN-019`, [`retracted.md`](retracted.md)). The two index orders and the permutation stand.

### TOWN-153

`+0xf8` and `+0x1ac` hold surfaces loaded by `R1152` from `column/mage/mask.bmp` (100x120 8bpp) and `column/fighter/mask.bmp` (92x120 8bpp). An image-wide displacement scan finds five references to each inside the school's address range — the constructor writing null, the loader writing the surface, `R1969` reading and rewriting, and `R1477` reading — and no blit wrapper ever receives either field. `R1152` reads `width*height` bytes and then calls `R1970`, which loops `height/2` times and exchanges row `i` with row `height-1-i` through the scratch buffer at `L11670`. Measured against that row order, each mask's five mapped code regions have their largest overlap with the rectangle their own switch arm selects, 5 of 5 in both panels, mean IoU 0.840 for the mage panel and 0.768 for the fighter panel, against 0.097 and 0.107 under the other panel's arms and rectangles.

Under the unreversed row order the same measurement gives 3 of 5 and 0.437, and 0 of 5 and 0.025. Both roots give identical numbers (`evidence/disasm-R1970-mask-rowflip.txt`, `evidence/mask-census.txt`)

**Confidence.** High for the code path — the loader's call, the reversal loop and every reference to both fields read at instruction level; the geometry is a corpus measurement over the two shipped mask pictures and agrees with it

### TOWN-154

38 are reached by a filename literal the code pushes and 25 by one of two `%.4d` format strings, counted on both roots against every image string carrying the prefix. The inventory: `trnhall.bmp` 480x480 into `+0x70`, drawn opaque at the view origin; `column/mage/mask.bmp` 100x120 8bpp into `+0xf8` and `column/fighter/mask.bmp` 92x120 8bpp into `+0x1ac`, never drawn; 15 `column/fighter/{sword,axe,pike,club,bow}/{on,shine,shine_on}.bmp` into `+0x268`..`+0x2a0` and 15 `column/mage/{fire,water,earth,air,astral}/{on,shine,shine_on}.bmp` into `+0x2b8`..`+0x2f0`, drawn keyed on their own class's rectangles; 16 `column/rt0000.bmp`..`rt0015.bmp` 148x208 into the `+0x30c` array, drawn opaque at (168,176); 9 `diamond/on0000.bmp`..`on0008.bmp` 80x76 into the `+0x24c` array, whose two further fields are `+0x250` and `+0x254`, whose animator `R1492` steps `+0x260` between 0 and 8 by `+0x264` and publishes the current frame in `+0x25c`, drawn opaque at (200,60) when both `+0x25c` and `+0x264` are non-zero; `buttonsarea.bmp` 160x238 and four `buttons/b{1,2}{on,off}.bmp` 140x46, whose filenames are pushed at `L11671`..`L11672` in a routine not read here.

This corrects `TOWN-061`'s step-6 clause, which calls `+0x25c` a selected-skill description panel (`evidence/art-census.txt`, `evidence/training-art-en.csv`, `evidence/training-art-ru.csv`, `evidence/disasm-R1486-loader.txt`)

**Confidence.** High (dual-root census with every entry matched to a push site or a format string, and the animator read at instruction level)
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration. The agreement between the per-root
data stands as corroboration: the training-art archive entries are separate files in each root.

### TOWN-155

`R1971` gates its idle branch on `+0x31c == 0 || +0x31c == [+0x310]-1`, reading the element count field. Every other site uses a literal: the loader's own loop bound, a compare of the loop index with 0x10 at `L11658`, the paint's mage icon branch comparing the 32-bit field at offset 0x31c of the view with 0xf at `L07789`, the paint's rotation stop comparing the frame index with 0xf at `L11673`, the hit test's panel test comparing the index with 0xf at `L11674`, `R1968`'s comparison of the index with 0xf at `L11675`, and `R1962`'s start guard comparing the 32-bit field at offset 0x31c of the view with 0xf at `L11676`. Adding a seventeenth `rt00NN.bmp` would not be loaded at all, and lowering the count below 16 would move only `R1971`'s idle test while every panel test still waited for 15.

The two icon rectangle runs are likewise `.text` immediates written by the constructor, so panel geometry is a code constant as well (`evidence/disasm-R1486-loader.txt`, `evidence/disasm-R0706-R1962-R1963-R1971.txt`, `evidence/rect-runs.txt`)

**Confidence.** High (every cited comparison read at instruction level, and the one data-driven site read in full)

## Town paint, door labels and the click table

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-156 | `townmain.bmp` draws first through the opaque blit and `Town_add.bmp` draws second through the keyed blit, both at the same destination point, and `Town_add.bmp` is gated on a runtime flag, not painted every frame. | High | ● active | [EXP-0196](../experiments/EXP-0196-town-paint/) |
| TOWN-157 | `Town_add.bmp` and its accompanying 1-3 bird sprites are one periodic event, re-armed on a randomised 1000-2999 ms delay recomputed at every firing, not a fixed frame-advance rate. | High | ● active | [EXP-0196](../experiments/EXP-0196-town-paint/) |
| TOWN-158 | The bird-progress array advances by one count each time a shared dispatch hub runs, and that hub runs at most once per roughly 67 ms of elapsed real time — a different constant from the school's 83 ms (`TOWN-147`). | High / Unknown | ● active (partially retracted) | [EXP-0196](../experiments/EXP-0196-town-paint/) |
| TOWN-159 | Two further object fields, `+0xf4` and `+0x120`, each reroll on their own randomised 2000-6999 ms delay and pick a random element from their own per-instance table, in a routine that runs unconditionally on every paint call… | High / Unknown | ● active (partially retracted, superseded) | [EXP-0196](../experiments/EXP-0196-town-paint/) |
| TOWN-160 | the global at `L11571` is a contiguous run of thirteen `{x,y}` dword pairs (screen-coordinate-range values, `0x68`..`0x250`), immediately followed at `L11677` by the ASCII string `main\text\tips\town.txt`… | High / Unknown | ● active | [EXP-0196](../experiments/EXP-0196-town-paint/) |
| TOWN-161 | The three door labels are gated on one shared field, `+0xb4`, tested for exact equality against `1` (shop), `4` (school) and `2` (tavern), each blitted opaque at its own fixed destination; `+0xb4` is initialised to `-1` at enter… | High / Unknown | ● active (partially retracted) | [EXP-0196](../experiments/EXP-0196-town-paint/) |
| TOWN-162 | All four of `TOWN-005`'s "class-keyed" sprite fields draw every paint call, gated only by their own null-check, never by a class-selector field, alongside at least five further sprite fields of the same shape… | High | ● active | [EXP-0196](../experiments/EXP-0196-town-paint/) |
| TOWN-163 | Exactly five sampled mask-byte values produce a click action; every other sampled value, and every value outside `0x80..0xc0`, produces no message at all, and the complete byte value → posted-message table is: `0x80`→`0x42b`… | High | ● active | [EXP-0196](../experiments/EXP-0196-town-paint/) |
| TOWN-164 | Of the five cases, only the `0x42d` case (mask `0xa0`) is gated, on a mission-validity check that redirects to an error dialog instead of posting; four of the five cases, all but `0x41f` (mask `0xb0`)… | High | ● active | [EXP-0196](../experiments/EXP-0196-town-paint/) |
| TOWN-165 | Town's own tip-popup call is read directly, not by structural analogy: filename operand `main\text\tips\town.txt` at `L11677`, rect `(328,0,640,200)`, popup-class id `0x467`, gated on `[L03631] != 0` at enter… | High | ● active | [EXP-0196](../experiments/EXP-0196-town-paint/) |

### TOWN-156

`R1489` (town paint) blits `+0x68` (`townmain.bmp`) through vtable slot `+0x18` at `L11678`, destination `([+0xc],[+0x8])`. The very next call, at `L11679`, is `R1917`, whose first instruction tests the byte at `[+0x208]` against 0x80 and branches to its own return when zero (`L11680`..`L11681`), skipping everything including the `Town_add.bmp` blit at `L11682` (through vtable slot `+0x38`, the same destination fields `[+0x8]`,`[+0xc]`). `TOWN-002` established both pictures load unconditionally and scoped itself to loading, leaving paint order open; the load is unconditional and the paint of `Town_add.bmp` is not (`evidence/disasm-townpaint.txt`, `evidence/disasm-townadd.txt`)

**Confidence.** High (both blit sites and the gating test read at instruction level; the vtable slot semantics — `+0x18` opaque, `+0x38` keyed — established by `TOWN-150`)

### TOWN-157

`R1916` sets `+0x208` bit `0x80` (the gate `TOWN-156` reads), rolls `R0179` twice to set `+0xbc` (`0..2`) and `+0xc0` (`1..3`, the bird count), conditionally plays one of two sound effects through `R1476`/`R0386` keyed on whether the roll gave exactly one bird, zeroes `+0xd8`/`+0xdc`/`+0xe0` (the per-bird progress array `R1917` reads), and returns a delay computed as `(rand() scaled) mod 2000 + 1000`. The paint routine calls `R1916` only when `timeGetTime() - [+0xb8]` exceeds a delay held in the static `[L11526]` and `+0x208` bit `0x80` is clear (`L11683`..`L11684`); the call's own return value is stored back into `[L11526]` for the next check, and `[L11526]` is seeded once, on the same formula, behind a one-shot latch at `[L11525]` bit `0x2` (`L11685`..`L11686`).

`R1917` clears bit `0x80` (ending the event) once every bird's progress value reaches its own object's own frame-count field, `L11687`..`L11688` (`evidence/disasm-townpaint.txt`, `evidence/disasm-misc1.txt`, `evidence/disasm-townadd.txt`)

**Confidence.** High (every cited field write and comparison read at instruction level)

### TOWN-158

The town paint routine gates a block on `timeGetTime() - [L11520] > 0x43` (67 decimal) at `L11563`/`L11689`, storing a fresh `timeGetTime()` into `[L11520]` only when the block runs (`L11690`..`L11691`). Inside that block, `R1908` is called unconditionally (`L11692`) and: increments `+0xd8`,`+0xdc`,`+0xe0` by one each, only if `+0x208` bit `0x80` is set (`L11693`..`L11694`); dispatches eight further routines by testing eight other bits of `+0x208` (`R1923` bit `0x1`, `R1922` bit `0x2`, `R1924` then `R1925` bit `0x4`, `R1918` bit `0x10`, `R1928` bit `0x40`, `R1929` bit `0x20`, `R1911` bit `0x400`, `R1909` bit `0x200`, `R1910` bit `0x100`); and calls `R1926` and `R1927` unconditionally every time it runs.

The same 67 ms-gated block also rolls two independent PRNG checks (thresholds `0x5e`/100 and `0x61`/100) that set `+0x208` bits `0x40` and `0x20` (`L11695`..`L11696`). Which named series (sign/door/stars/fluger) each of the eight bit-gated routines animates was not traced past its own entry point in this experiment; only the birds' own routine (`TOWN-157`) was read to completion (`evidence/disasm-townpaint.txt`, `evidence/disasm-misc2.txt`)

**Confidence.** High for the dispatch structure and the 67 ms gate (every branch and the timer read at instruction level); Unknown for which named series each of the eight callees animates

**Amended.** see TOWN-405 and retracted.md. [`retracted.md`](retracted.md) holds an entry for this claim (partially retracted).

### TOWN-159

Two further object fields, `+0xf4` and `+0x120`, each reroll on their own randomised 2000-6999 ms delay and pick a random element from their own per-instance table, in a routine that runs unconditionally on every paint call, not on the 67 ms hub of `TOWN-158`. `R1907` runs unconditionally every paint call, called at `L11697` immediately before the 67 ms-gated block. First field: when `timeGetTime() - [+0x118] > [+0x11c]`, rerolls `[+0x11c]` to `(rand() scaled) mod 5000 + 2000` and picks `[[+0xfc] + (rand() mod [+0x100])*4]` into `+0xf4` (`L11698`..`L11699`). Second field is the same shape at `+0x148`/`+0x14c`/`+0x128`/`+0x12c`/`+0x120`, and additionally sets `+0x208` bit `0x100` (`L11700`..`L11701`), where the first sets bit `0x200`.

The paint routine draws `+0xf4` opaque, frame `[+0x114]` (clamped to 0 if `-1`), at `([+0x110]+[+0xc], [+0x10c]+[+0x8])` (`L11702`..`L11703`); draws `+0x120` opaque, frame `[+0x140]`, at `([+0x138]+[+0xc], [+0x13c]+[+0x8])`, preceded by a sound-effect dispatch keyed on `+0x144` (`0`, `1`, or other) and on `+0x140` against the literals `0`, `8` and `0xe`, using the same `R1476`/`R0386` pair as `TOWN-157` (`L11704`..`L07764`); and separately draws a third, unrandomised field `+0x150`, own field `[+0x15c]` as its third argument, at `([+0x158]+[+0xc], [+0x154]+[+0x8])` (`L11705`..`L11706`). No filename or roster identity for `+0xf4`, `+0x120` or `+0x150`, and no mapping to a named series (sign/stars/fluger), was established in this experiment (`evidence/disasm-townpaint.txt`, `evidence/disasm-misc2.txt`)

**Confidence.** High for the mechanism (every field, comparison and call site read at instruction level); Unknown for which named series either reroll corresponds to and for the identity of any of the three drawn objects

**Amended.** see TOWN-445 and retracted.md. [`retracted.md`](retracted.md) holds 4 entries for this claim (partially retracted, resolved, superseded, refuted). The destination-pair clause of the `+0xf4`, `+0x120` and `+0x150` draws printed above is refuted: the pairs list the pushes in push order, and the blit's first argument is the first table field plus `[view+8]` (`TOWN-503`, `retracted.md`).

### TOWN-160

the global at `L11571` is a contiguous run of thirteen `{x,y}` dword pairs (screen-coordinate-range values, `0x68`..`0x250`), immediately followed at `L11677` by the ASCII string `main\text\tips\town.txt`, and the two further addresses the loader cites, the global at `L08048` and the global at `L11572`, land exactly on pair boundaries five and nine into that same run — consistent with either one 13-entry table or three back-to-back tables of 5, 4 and 4 entries. Raw bytes read directly at all three addresses on both preserved roots, byte-identical (`evidence/dat-tables-hex.txt`). An image-wide 32-bit-immediate scan for all three addresses (`L11571`, `L08048`, `L11572`) finds zero instructions in `.text` that reference any of them directly; the same scan for two further offsets inside the run, `L11707` and `L11708`, is also zero.

No consumer of this table was found in this experiment, and the `{x,y}` reading is a plausible interpretation of the value ranges, not a confirmed one (`evidence/scan-tipstring-and-nearby.md`)

**Confidence.** High for the raw bytes and the byte-offset structure; Unknown for which grouping (one table or three) the code intends, for the semantic type, and for the table's reader

### TOWN-161

The three door labels are gated on one shared field, `+0xb4`, tested for exact equality against `1` (shop), `4` (school) and `2` (tavern), each blitted opaque at its own fixed destination; `+0xb4` is initialised to `-1` at enter, and no writer of `1`, `2` or `4` was found in this experiment. In paint order: a compare of `[+0xb4]` with 1, branching when not equal, skips a blit of `+0x1dc` at destination `([+0xc]+0x108,[+0x8]+0x108)` (`L07768`..`L07769`); a compare of `[+0xb4]` with 4, branching when not equal, skips `+0x1c4` at `([+0xc]+0x12c,[+0x8]+0x1b4)` (`L07770`..`L11709`); a compare of `[+0xb4]` with 2, branching when not equal, skips `+0x164` at `([+0xc]+0x14c,[+0x8]+0x90)` (`L07771`..`L11710`). All three blit through vtable slot `+0x18`.

`R1383` (town enter) sets `+0xb4` to `-1` at `L11711`, alongside two other fields reset to the same sentinel (`+0x114`, `+0x140`). No instruction that writes `1`, `2` or `4` into `+0xb4` was found while reading the paint, enter and click-dispatch routines this experiment covers; the writer, presumably a mouse-move hover handler, was not located (`evidence/disasm-townpaint.txt`)

**Confidence.** High for the discriminator, the three comparisons and the three draw sites (all read at instruction level); Unknown for the runtime writer

**Amended.** see TOWN-453 and retracted.md. [`retracted.md`](retracted.md) holds an entry for this claim (partially retracted).

### TOWN-162

All four of `TOWN-005`'s "class-keyed" sprite fields draw every paint call, gated only by their own null-check, never by a class-selector field, alongside at least five further sprite fields of the same shape; the fields do not select among alternatives, they all draw together. Each of the following is null-tested (the field value tested for zero, branching when it is zero) past its own blit and nothing else gates it, in paint order: `+0x160` (frame `+0x168`, dest offsets `+0xc`+`0x138`, `+0x8`+`0x7c`); `+0x180` (no frame field, dest offsets `+0xc`+`0xe8`, `+0x8`+`0x168`); `+0x19c` (no frame field, dest offsets `+0xc`+`0x94`, `+0x8`+`0xb4`); `+0x1bc` (no frame field, dest offsets `+0xc`+`0x120`, `+0x8`+`0x154`); `+0x1c8` (frame `+0x1cc`, dest offsets `+0xc`+`0x158`, `+0x8`+`0x204`); `+0x1d0` (frame `+0x1d4`, dest offsets `+0xc`+`0x148`, `+0x8`+`0x1c4`); `+0x1d8` (frame `+0x1e0`, dest offsets `+0xc`+`0x128`, `+0x8`+`0x114`); `+0x1f8` (no frame field, dest offsets `+0xc`+`0x40`, `+0x8`+`0x134`); `+0xe4` (own field `[+0xe8]` as third argument, dest offsets `+0xc`+`0x9e`, `+0x8`+`0xb8`) (`L11712`..`L11713`).

This directly answers `TOWN-005`'s open "which one (if any single one)" question for its four named fields (`+0x160`,`+0x1c8`,`+0x1d0`,`+0x1d8`): none is selected, all draw whenever loaded. Mapping each field to its own filename (`townbirds`/tavern/fighter/mage/shopie) was not attempted in this experiment (`evidence/disasm-townpaint.txt`)

**Confidence.** High (every null-check and blit call read at instruction level)

### TOWN-163

Exactly five sampled mask-byte values produce a click action; every other sampled value, and every value outside `0x80..0xc0`, produces no message at all, and the complete byte value → posted-message table is: `0x80`→`0x42b`, `0x90`→`0x42a`, `0xa0`→`0x42d`, `0xb0`→`0x41f`, `0xc0`→`0x42c`. `R1914` samples one byte from the mask surface, subtracts `0x80`, and rejects (returns `-1`) any value where the result is unsigned-greater than `0x40` (`L11714`..`L11715`). The remaining 65 values index a byte table at `L11716`: only offsets `0`, `0x10`, `0x20`, `0x30`, `0x40` (raw bytes `0x80,0x90,0xa0,0xb0,0xc0`) hold `0,1,2,3,4`; the other 60 hold `5`.

A jump table at `L11717` sends case `5` to a "return -1" tail (`L11718`) and cases `0..4` to five `RET` arms returning `2,1,8,16,4` respectively. The caller, `R0704`, decrements that return value and rejects (posts nothing, `L11719`) anything outside `0..15` (`L11720`..`L11721`); the survivors index a second byte table at `L11722` — only offsets `0,1,3,7,15` (values `1,2,4,8,16`) hold `0,1,2,3,4`, the other eleven hold `5` — and a second jump table at `L11723` sends case `5` to the same "nothing" tail and cases `0..4` to five case bodies, which push `0x42b` (`L11724`), `0x42a` (`L11725`), `0x42c` (`L11726`), `0x42d` (`L11727`) and `0x41f` (`L11728`) and call `PostMessageA` at import `[L03535]`.

Both lookup tables and both jump tables were read raw byte-for-byte, on both preserved roots, byte-identical (`evidence/hexdump-tables.txt`)

**Confidence.** High (every table byte and every case body read directly, on both roots)

### TOWN-164

Of the five cases, only the `0x42d` case (mask `0xa0`) is gated, on a mission-validity check that redirects to an error dialog instead of posting; four of the five cases, all but `0x41f` (mask `0xb0`), first send an internal message `0x445` through the target object's own `vt+0x48` slot before posting the external message, into the same forwarding chain `TOWN-139` already traces to a named terminal handler. The `0x42d` case calls `R1416`; if it returns `-1`, the case shows a fixed dialog string at `L11729` through `R0695` and returns without posting (`L11730`..`L11731`); otherwise it posts `0x42d` exactly like the other three ungated cases.

`R1972` — which loads the object's table pointer from its first field, pushes 0, 0 and 0x445, calls through slot +0x48 of that table and returns — is called at `L11732` (`0x42b` path), `L11733` (`0x42a` path), `L11734` (`0x42c` path) and `L11735` (`0x42d` path, after the validity check passes), each before the object lookup that leads to the `PostMessageA` call; the `0x41f` case (`L11728`..`L11736`) contains no such call. `TOWN-139` already establishes that a `vt+0x48` call carrying an unmatched id reaches `R0716` through the same three-routine forwarding chain, and that `R0716` "handles only `0x445` and `0x446`" rather than forwarding them further; this experiment did not read what that handling does.

An image-wide 32-bit-immediate scan for `0x445`, `0x41f`, `0x42a`, `0x42b`, `0x42c` and `0x42d` finds, besides the cited sites and matches inside unrelated `CALL`/`JMP rel32` displacement encodings: the same idiom of pushing 0, 0 and the id and then calling through slot +0x48 of a table, recurring at roughly forty further addresses across `.text` for `0x445` alone, none in town's own address range, consistent with the same internal-message convention used by other views; one further genuine `PostMessageA`-shaped push each for `0x41f` (`L11737`) and `0x42a` (`L11738`), in code outside town, consistent with a different view class reusing the same numeric id under MFC's per-class message numbering rather than a shared consumer; and, for `0x42c`, sixteen further hits, most in the pattern of a 32-bit store of the constant 0x42c to a memory location immediately followed by a second dword of `0xffffffff`, which is not the `PUSH`/`CALL PostMessageA` shape and was not resolved.

No Win32-level message-map table (a static `.rdata` array pairing a message id with a handler address, of the kind an MFC-style message-dispatch macro expands to) was read in this experiment; this repository has no existing tool to decode one, and neither of the already-decoded tavern/school 89-entry internal message tables (`TOWN-136`/`TOWN-137`) carries an arm for any of the five posted ids (`evidence/scanimm-msgids.txt`)

**Confidence.** High for the gate, the internal send, and the connection to `TOWN-139`'s forwarding chain (all read at instruction level); explicit negative finding for what `R0716` does with `0x445`/`0x446` and for a Win32 message-map consumer of the five posted ids, scoped to a whole-image immediate-operand scan, the two already-decoded internal message tables, and the absence of message-map decoding tooling in this repository

### TOWN-165

Town's own tip-popup call is read directly, not by structural analogy: filename operand `main\text\tips\town.txt` at `L11677`, rect `(328,0,640,200)`, popup-class id `0x467`, gated on `[L03631] != 0` at enter — matching `TOWN-015`/`TOWN-025` by a second, independent instrument. `R1383` pushes `L11677` at `L11739` (the only instruction in the image that references this address, confirmed by an image-wide immediate scan) into `R1293`, then pushes `0x148, 0, 0x280, 0xc8, 0x467` (left, top, right, bottom, class id) at `L11740`..`L11741` into `R1261`, reached only when `[L03631] != 0` (`L11742`..`L11743`).

`TOWN-015`/`TOWN-025` established the same facts by decompiler string resolution; this reads the same call site as raw bytes plus a direct address scan, and the two instruments agree exactly (`evidence/disasm-townpaint.txt`, `evidence/scan-tipstring-and-nearby.md`)

**Confidence.** High (raw byte read and independent address scan agree with the prior decompiler reading)

## Room buttons and tip popups

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-181 | The school's button-bitmap loader (`TOWN-154`'s unread push range `L11671`..`L11672`, `R1973`) maps each of its five pushes to a named field of the object at `[school_view+0x78]`… | High | ● active | [EXP-0197](../experiments/EXP-0197-room-buttons-and-tips/) |
| TOWN-182 | The two training buttons' hit-rects are standard Win32 `RECT`s at `+0x88` and `+0x98` (16-byte stride), tested with the genuine `PtInRect` import against the cursor position relative to an origin at `[+0x5c]->+8/+0xc`… | High | ● active | [EXP-0197](../experiments/EXP-0197-room-buttons-and-tips/) |
| TOWN-183 | `R1905` is `TOWN-094`'s already-published school upper-right-widget paint routine (id `0x3fd`, constructed by `R1940`, rect `(464,0,640,238)`), and its full body — read for the first time in this experiment… | High | ● active (amended, partially retracted) | [EXP-0197](../experiments/EXP-0197-room-buttons-and-tips/), **[EXP-0205](../experiments/EXP-0205-town-page-composition/)** |
| TOWN-184 | `R1261`'s six operands are, in order: `operand1 = control/dialog-item id` (stored at `object+4`, not a window class or vtable selector), `operand2 = xLeft`, `operand3 = yTop`, `operand4 = xRight`, `operand5 = yBottom`… | High | ● active (amended, partially retracted) | [EXP-0197](../experiments/EXP-0197-room-buttons-and-tips/), [EXP-0202](../experiments/EXP-0202-chargen-page-geometry/) |
| TOWN-185 | `R1974` (called from `R1261` immediately after the rect is set) constructs three child controls on the popup, ids `0xd`/`0xe`/`0xf`, each through a different constructor function; control `0xf` is the "show tips" checkbox… | High | ● active (amended) | [EXP-0197](../experiments/EXP-0197-room-buttons-and-tips/), **[EXP-0198](../experiments/EXP-0198-tip-popup-presentation/)** |
| TOWN-186 | the global at `L03631` has two direct writers, both read in full, and it persists across process runs in the Windows Registry under the value name "TipsMode"… | High | ● active (amended) | [EXP-0197](../experiments/EXP-0197-room-buttons-and-tips/), [EXP-0412](../experiments/EXP-0412-town-square/EXP-0412.md); TOWN-480 |
| TOWN-187 | The character generator's two own tip-popup call sites are read directly: a class-conditional (not gender-conditional) first popup at the lower-left of the screen… | High | ● active | [EXP-0197](../experiments/EXP-0197-room-buttons-and-tips/) |
| TOWN-188 | the global at `L03631` is referenced by one fixed absolute address from every tip call site read across this repository, in at least six structurally unrelated screens… | High / Medium | ● active (partially retracted) | [EXP-0197](../experiments/EXP-0197-room-buttons-and-tips/) |
| TOWN-206 | Both caption ids `TOWN-185` left unread name plain lines of `text/main.txt`, and both are now read on both preserved roots: id `0x7f` is "Close"/"Закрыть", id `0x80` is "Show tips next time"/"Показывать далее". | High / Medium | ● active | [EXP-0198](../experiments/EXP-0198-tip-popup-presentation/) |
| TOWN-207 | Control `0xe`, captioned "Close" (`TOWN-206`), is the popup's own close control, established by an executable binding and not by caption text alone… | High / Medium | ● active | [EXP-0198](../experiments/EXP-0198-tip-popup-presentation/) |
| TOWN-208 | The popup's own `vt+0x2c` (`TOWN-062`'s address for the tavern's shared "no own content" epilogue) is a generic three-step composite — compute a rect, dispatch through the instance's own `vt+0x30`, then walk children — and for the popup… | High | ● active | [EXP-0198](../experiments/EXP-0198-tip-popup-presentation/) |
| TOWN-209 | The popup's own `vt+0x30` (`R1975`) draws a tiled panel/border through a single fixed global object, using integer tile ids `9`..`0x11` and no archive path literal of its own. | High / Medium | ● active (amended, superseded) | [EXP-0198](../experiments/EXP-0198-tip-popup-presentation/), [EXP-0217](../experiments/EXP-0217-cursor-composition-order/) |
| TOWN-210 | The global object the popup's own border draw reads, `[L03604]` (`TOWN-209`), has exactly one write site system-wide and is constructed from the archive node `graphics\interface\lm.256`. | High | ● active | [EXP-0198](../experiments/EXP-0198-tip-popup-presentation/) |
| TOWN-211 | The base Control class's generic message handler (`TOWN-139`'s already-named `vt+0x48` default, `R0390`) checks three message-id ranges numerically matching the real Win32 keyboard (`0x100`-`0x102`)… | High / Medium | ● active (partially retracted) | [EXP-0198](../experiments/EXP-0198-tip-popup-presentation/) |

### TOWN-181

The school's button-bitmap loader (`TOWN-154`'s unread push range `L11671`..`L11672`, `R1973`) maps each of its five pushes to a named field of the object at `[school_view+0x78]`, a sub-object distinct from the school view's own `+0x74` field (`TOWN-061`'s scrollbar-track, on a different object). `L11671`→`+0x74`=`buttons\b1on.bmp`, `L11744`→`+0x7c`=`buttons\b1off.bmp`, `L11745`→`+0x78`=`buttons\b2on.bmp`, `L11746`→`+0x80`=`buttons\b2off.bmp`, `L11672`→`+0x84`=`ButtonsArea.bmp` (all under `graphics\interface\training\`), each store a 32-bit write of the returned bitmap pointer to the field at that offset of the object, read directly after its own `R1176` call (`TOWN-150`'s picture loader; `R1176` itself sets the object's first field to L11667, tying every one of these five bitmaps to `TOWN-150`'s already-decoded vtable).

`R0706` (school enter) calls this loader on `[this+0x78]`, alongside a release routine `R1976` (frees the same five fields on re-entry, two loops of two plus the fifth field alone) and an initializer `R1977` (sets two further fields, `+0xac` and `+0xb0`, to `-1`) (`evidence/disasm-button-loader-release-R1973-R1976-R1977.txt`, `evidence/strdump-button-bitmap-filenames.txt`)

**Confidence.** High (every push, its store and its filename string read directly; the loader/vtable tie to `TOWN-150` read at instruction level)

### TOWN-182

The two training buttons' hit-rects are standard Win32 `RECT`s at `+0x88` and `+0x98` (16-byte stride), tested with the genuine `PtInRect` import against the cursor position relative to an origin at `[+0x5c]->+8/+0xc`, and the result feeds one of the two fields that gate the ON bitmap; the writer of the other gating field was not found. `R1978` (`this` = panel object) computes `(cursor.x−origin.x, cursor.y−origin.y)` and calls `PtInRect` (import `[L01178]`) against `[this+0x88]`, then `[this+0x98]`, returning the matching index (`0`/`1`) or `-1` (`R1978`..`L11747`). `R1977` initializes two further fields, `+0xac` and `+0xb0`, to `-1` at panel construction (sentinel: no button selected).

`R1979` writes `+0xb0` from `R1978`'s return and sets a repaint-dirty bit at `[+0x5c]->+0x350` when the value changes (`evidence/disasm-button-hittest-R1979.txt`). No writer of `+0xac` was found; every instruction that touches it in the six routines this experiment reads only compares it (see `TOWN-183`). A `.text`-wide displacement scan for a store through `[reg+0xac]` returned 461 hits across unrelated object classes, too many to isolate within this experiment's scope, and was abandoned (`evidence/disasm-button-latch-hittest-R2042-R1978.txt`)

**Confidence.** High for the hit test and the `+0xb0` writer (both read at instruction level, `PtInRect` confirmed by import symbol resolution); explicit negative finding for `+0xac`'s writer, scoped to the six routines this experiment reads plus one whole-`.text` displacement scan

### TOWN-183

`R1905` is `TOWN-094`'s already-published school upper-right-widget paint routine (id `0x3fd`, constructed by `R1940`, rect `(464,0,640,238)`), and its full body — read for the first time in this experiment — draws the two training buttons with an ON/OFF choice at exactly their own hit-rect's corner, not the plain sprite/digit display `TOWN-094`'s structural comparison implied. The routine first blits `+0x84` (`ButtonsArea.bmp`, 160×238 per `TOWN-154`) opaque through vtable slot `+0x18` (`TOWN-150`'s primitive) at `(x=[+0x8]+origin.x+0x10, y=[+0xc]+origin.y)`, `origin` read from `[+0x5c]->+8/+0xc` (`L11748`..`L11749`). It then loops `i=0,1` over a stride-4 range (the offset runs `0x360..0x368`): for each `i` it reads `+0xac` and `+0xb0` (`TOWN-182`) and blits the ON bitmap (`+0x74` for `i=0`, `+0x78` for `i=1`) through the same vtable `+0x18` slot only when `+0xb0==i` **and** `+0xac==i` (`L11750`..`L11751`), else the OFF bitmap (`+0x7c`/`+0x80`, `L11752`..`L11753`).

Both branches draw at `origin + (rect[i].left, rect[i].top)`, where `rect[i]` is the **same** `+0x88`/`+0x98` `RECT` `TOWN-182`'s hit test reads: the walking pointer starts at `+0x94` (`i=0`) and `+0xa4` (`i=1`) — each rect's own `bottom` field — and the draw reads pointer-0xc/pointer-8, which land on that same rect's `left`/`top` (`+0x94-0xc=+0x88`, `+0x94-8=+0x8c`; `+0xa4-0xc=+0x98`, `+0xa4-8=+0x9c`). **Amended by [EXP-0205]: each iteration draws TWO text/number labels, not one.** Both branches (raw disassembly, `L07784`/`L07785` in the ON path, `L07786`/`L07787` in the OFF path) call the same undecoded vtable `+0x14` slot twice — first with a string pointer read from an array at `widget+0x64` indexed by button, then, after a call to `R0807`, with a value read from a per-button scratch dword at `room+0x360`/`room+0x364` (`room` being the widget's own `+0x5c` parent, `TOWN-261`) — on an object selected from `[L09922]` or `[L09928]` by the same `+0xb0==i` test.

`TOWN-280` resolves both calls' X and the second call's Y for both buttons. The routine references all 5 fields `R1973` loads (`TOWN-181`) but issues exactly 3 opaque blits per call through `vt+0x18` (background once, and exactly one of ON/OFF per button, never both) — that blit count is unaffected by the label-count correction (`evidence/disasm-button-panel-paint-R1905.txt`)

**Confidence.** High (every branch, comparison, blit call and the hit-rect/paint-offset identity read at instruction level; the routine's own invocation is established by the already-published `TOWN-094`, not re-derived here)

**Amended.** [`retracted.md`](retracted.md) holds an entry for this claim (refuted).

### TOWN-184

`R1261`'s six operands are, in order: `operand1 = control/dialog-item id` (stored at `object+4`, not a window class or vtable selector), `operand2 = xLeft`, `operand3 = yTop`, `operand4 = xRight`, `operand5 = yBottom`, `operand6 = text pointer` (stored at `object+0x64`) — settled by reading the constructor chain to its own leaf, not by which reading fits a 640×480 screen. Two published rows characterize `operand1` as a class: `TOWN-015` calls it "a window class/style id" and `TOWN-165` calls it "popup-class id"; both phrases are superseded by this row — `operand1` is a control/dialog-item id, and the vtable that actually determines the popup's own class (`L06362`) is set unconditionally by the constructor, never read from `operand1`.

`TOWN-025` and `SHOP-TIP-045` do not characterize `operand1` as a class (`SHOP-TIP-045` already calls it "id 1011") and are unaffected on this point. Separately, this row also refutes `TOWN-015`'s derived "640×200 at (328,0)"; the raw operands are unaffected, only that narrative phrase is wrong (correct reading: 312×200, right-edge-anchored at x=640). `SHOP-TIP-045`'s explicit left/top/right/bottom phrasing is confirmed correct; `TOWN-165`'s raw-tuple phrasing carried no derived width/height and is unaffected on that second point. Amended by `EXP-0202` (`TOWN-246`): this sentence named the wrong room.** The refuted phrase, "640×200 at (328,0)," is `TOWN-015`'s phrase for the TOWN's tip rectangle (`R1383`), not the tavern's (`R1412`); the tavern's own phrase in `TOWN-015`, "312×200 at (0,0)," was always correct and needed no refutation. The numeric correction here (312×200, right-edge-anchored at x=640) is right and describes the town's own popup, not the tavern's.** `R1261` sets its own vtable to `L06362` unconditionally (`L11754`..`L11755`, no branch guards it, so `operand1` cannot select a class) and calls `R0759(id, xLeft, yTop, xRight, yBottom, text)` unchanged, which sets vtable `L06368`, stores the text pointer at `+0x64` (`L11756`..`L11757`), and calls `R0773(id, xLeft, yTop, xRight, yBottom)`, which stores `id` at `+4` and forwards the four rect operands unchanged to `R0388`, a thin forwarder to `R1980`.

`R1980` is Ghidra's own auto-analysis label "`SetRect`", a **local wrapper**, distinct from the genuine Win32 import: an image-wide symbol scan for `SetRect` finds two separate symbols, `SYMBOL SetRect @ EXTERNAL:0000014e` (the import, resolving `PTR_SetRect_L11758`, called from 24 unrelated sites elsewhere in the image) and `SYMBOL SetRect @ R1980` (this local wrapper, called only from `R0388` and one unrelated routine `R1219`). `R1980` pushes its own four rect parameters unchanged and calls through the dword at `[L11758]` — the genuine `SetRect` import, whose documented parameter order is `(LPRECT, xLeft, yTop, xRight, yBottom)` — which fixes the LTRB reading for every operand this chain forwards, back to `R1261`'s own `operand2..5` (`evidence/disasm-popup-ctor-R1261.txt`, `evidence/disasm-popup-base-ctor-controls-R0759-R1974.txt`, `evidence/disasm-control-base-ctor-R0773-rectaccessors.txt`, `evidence/disasm-setrect-forwarder-R0388.txt`, `evidence/disasm-setrect-local-wrapper-R1980.txt`, `evidence/scan-setrect-symbol-resolution.md`)

**Confidence.** High (the full constructor chain read at instruction level to the genuine Win32 `SetRect` import, identified by symbol resolution rather than by name similarity or by which reading looks plausible on screen); the operand-order and `operand1` findings are unaffected by the amendment above — only the room name in the "Separately" aside carried High while wrong, retracted by `TOWN-246`

**Amended.** [`retracted.md`](retracted.md) holds an entry for this claim (refuted).

### TOWN-185

`R1974` (called from `R1261` immediately after the rect is set) constructs three child controls on the popup, ids `0xd`/`0xe`/`0xf`, each through a different constructor function; control `0xf` is the "show tips" checkbox — it is default-checked at construction and its id is the exact sub-id the flag's own message handler matches. Control `0xd` (`0x98` bytes, `R0699(id=0xd, 0x14, 0x18, width-0x1c, height-0x24)`, `width`/`height` from the popup's own client rect via `TOWN-184`'s accessors) carries no caption. Control `0xe` (`0x78` bytes, `R0700(id=0xe, width-0x78, height-0x28, width-0x28, height-0x16, textPtr)`) is constructed with a caption already resolved: string-table id `0x7f`, loaded via `R1773(table=L09841, langSelector=[L02677], id=0x7f)` before the constructor call.

**Amended by `TOWN-206`: both parameter lists above are correct but incomplete** — `R0699` takes nine stack arguments, not five (a return cleaning 0x24 bytes, confirmed from the callee's own epilogue), the four unlisted trailing ones being `outerField[+0x18]` (`R1974`'s own local at that call site, not traced further), `[L02677]`, `L03616` and `0x0`; `R0700` takes eleven, not six (a return cleaning 0x2c bytes), the five unlisted trailing ones being `[L02677]`, `L03616`, `0x45a`, `0x0` and `0x0`. `R1981` (control `0xf`, a return cleaning 0x20 bytes = eight arguments) was already complete as published. The shared trailing triple `[L02677]`, `L03616`, `0x0` (plus, for `0xe` only, an extra `0x45a` before the two trailing zeros) is passed to all three sibling constructors regardless of whether the control being built has a caption, so it is not exclusively a caption-lookup parameter (`evidence/disasm-control-ctors-R0699-R0700-R1981-L06358.txt`).

Control `0xf` (`0x8c` bytes, `R1981(id=0xf, 0x28, height-0x28, width-0x7c, height-0x18, [L02677], L03616, 0)`) has its caption set **after** construction — string-table id `0x80`, same loader, applied through a `SetText`-shaped call (`L11759`, a call to `L06358`) — and is then set checked by a virtual call on its **own** vtable slot `+0x44` with a local flag set to `1` (`L11760`..`L11761`). Control `0xf`'s id, `0xf`, is exactly the sub-id `R1478` matches for message `0x46d` (`TOWN-186`), the checkbox-toggle notification that writes straight into the global at `L03631`; and its default-checked state matches the global at `L03631`'s own hardcoded default of `1` (`TOWN-186`).

Which of controls `0xd`/`0xe` is the "close" control is settled by `TOWN-207`: it is `0xe`, not by caption alone but by an Enter-key-gated dispatch override neither the caption nor the id predicts on its own. The two string-table captions' text is read by `TOWN-206`: id `0x7f` = "Close"/"Закрыть", id `0x80` = "Show tips next time"/"Показывать далее". There is no Win32 `STRINGTABLE` resource involved: `TEXT-STRTAB-030` reads `R1773` to its own leaf and finds a four-byte array index into the same global line-pointer array `TEXT-STRTAB-023` already documents, not a resource API call — the earlier sentence naming a separate, unbuilt `STRINGTABLE` instrument described a mechanism this row's own captions do not use (`evidence/disasm-popup-base-ctor-controls-R0759-R1974.txt`)

**Confidence.** High for the three controls' ids, sizes, constructors and argument roles (now complete, `TOWN-206`), for control `0xf`'s checkbox identity (two independent structural ties: default-checked state matching the flag's default, and construction id matching the toggle message's sub-id, both read at instruction level), for both captions' text (`TOWN-206`), and for `0xe` being the close control (`TOWN-207`)

**Amended.** The amendment is stated in the claim text above.

### TOWN-186

the global at `L03631` has two direct writers, both read in full, and it persists across process runs in the Windows Registry under the value name "TipsMode", read/written by the same paired save/load routines that persist twelve other named engine options. `R1845` hardcodes the global at `L03631` (with two neighboring globals `L04367`, `L06260`) to `1` unconditionally — an initializer, with no read of any prior state. `R1478` is a message handler: `if (msg==0x46d && wParam==0xf) [L03631] = lParam` (`L07705`..`L07706`); it returns 0 for `0x100`/`0x445`/`0x446` and forwards every other id to `TOWN-139`'s already-documented `R0716` chain (`TOWN-480`).

`wParam==0xf` is exactly `TOWN-185`'s checkbox control id. The value survives a process restart: `R1982` (save) pushes `"TipsMode"` (`L11762`) then the address `L03631` (`L11763`..`L11764`) into a call chain terminating in a call through the dword at `[L10640]`, resolved by symbol scan to `PTR_RegSetValueExA_L10640`; `R1983` (load) does the same with `"TipsMode"` at `L11765` (`L11766`..`L11767`) into a call through the dword at `[L10643]`, resolved to `PTR_RegQueryValueExA_L10643`. Both routines carry the same 13-entry named-option table (`GameSpeed`, `FormationMode`, `WimpyMode`, `ShowAllHitPoints`, `Smoothing`, `ShowFlyingHP`, `ShowTimeFlow`, `TipsMode`, `AutoCasting`, `Acknowledgement`, `Shadows`, `Lighting`, `Animation`), each string read directly.

The registry key handle (`hKey`) is a parameter to both routines, passed by address (`&local_8`/`&local_c`) from each routine's own single caller: `R1984` (save side) and `R1985` (load side) each call `RegOpenKeyExA((HKEY)0x80000002, "SOFTWARE\1C\Allods", 0, 0x20019, &local)` before handing the local `HKEY` variable's address down through the chain that reaches `R1982`/`R1983`, and each calls `RegCloseKey` after. `0x80000002` is the Win32 predefined key constant `HKEY_LOCAL_MACHINE`. The full path for the "TipsMode" value is therefore `HKEY_LOCAL_MACHINE\SOFTWARE\1C\Allods` (`evidence/decompile-registry-save-caller-R1984.txt`, `evidence/decompile-registry-load-caller-R1985.txt`, `evidence/strdump-registry-key-path.txt`) (`evidence/disasm-tipsmode-flag-writers-R1845-R1478.txt`, `evidence/disasm-tipsmode-registry-save-R1982.txt`, `evidence/disasm-tipsmode-registry-load-R1983.txt`, `evidence/strdump-options-table-names.txt`, `evidence/scan-regsetvalueexa-symbol-resolution.md`, `evidence/scan-regqueryvalueexa-symbol-resolution.md`)

**Confidence.** High (both writers, both registry calls, and both routines' own single caller read at instruction level; both value-persistence imports and the key-opening call read directly, the key path read from the literal string operand, not inferred from the value name)

**Amended.** TOWN-480 corrects the message-forwarding clause from the complete popup slot: `L07701`..`L07702` returns 0 for `0x100`, `0x445` and `0x446`; `L07703`..`L07704` forwards every other id to `R0716`. The flag-write and registry-persistence clauses stand. The former wording is retained as CORRECTED in `retracted.md`.

### TOWN-187

The character generator's two own tip-popup call sites are read directly: a class-conditional (not gender-conditional) first popup at the lower-left of the screen, and a second popup that replaces the first popup's text in place rather than constructing a new one, on a once-only latch — the same mechanism `SHOP-TIP-045` already documents for the shop. `R1870` (chargen's own enter routine), gated on `[L03631] != 0`, tests the byte at offset 0x18c of the object against 0x2 (`L11768`) — the identical mask `TOWN-149` reads as the mage/non-mage panel split (`R0706` writes `view+0x31c=0xf` when this bit is set, `0` when clear) and `HERO-FIGURE-059` reads as the mage/non-mage draw-order split (`L04346`, a test of the byte at offset 0x18c against 0x2); this is a **different** field mask from the sex bit `HERO-APPEAR-055`/`HERO-APPEAR-045` establish at `+0x18c & 4`.

The branch pushes `main\text\tips\chrgen1m.txt` (`L11769`) when the bit is set, else `main\text\tips\chrgen1f.txt` (`L11770`) when clear — the mage branch and the non-mage branch. The shipped English text of both files, read directly on both preserved roots for this correction, matches the class reading and not a sex reading: `chrgen1f.txt` (416 bytes, EN) reads "...you may change your character's starting specialization by selecting one of the five **weapon skills**...", `chrgen1m.txt` (415 bytes, EN) reads "...selecting one of the five **magic spheres**..." — a fighter text and a mage text; the RU root's two files differ from EN in raw bytes (per `TOWN-007`'s established EN/RU tip-text divergence) and were read but not translated in this experiment (`evidence/dump-chrgen-tip-texts.txt`).

The `1f`/`1m` filename letters do **not** encode sex here, despite the unrelated, coincidentally identical-looking convention in `graphics.res`'s `equipment/` tree (`ffighter`/`fmage`/`mfighter`/`mmage`, where the same two letters *are* sex): read as "fighter"/"mage" initials instead of "female"/"male", `1f`/`1m` match the branch exactly. The popup itself: `R1261(id=0x467, left=0, top=0x118, right=0x138, bottom=0x1e0, text)` — a 312×200 popup at `(0,280)`, the lower-left quadrant, unlike every other room's own tip rect (all in the upper area of the screen). `R1986`, gated on the same `[L03631]` **and** a once-only latch at `+0x100` (zero = not yet fired) and an existing-popup pointer at `+0x80`, reads `main\text\tips\chrgen2.txt` and calls `R1987(existingPopup, newText)` to replace the first popup's text, never constructing a second popup object (`evidence/disasm-chargen-tips-R1870-R1986.txt`, `evidence/dump-chrgen-tip-texts.txt`)

**Confidence.** High (both call sites and the latch mechanism read at instruction level; the class-vs-sex mask identity confirmed against `TOWN-149`'s and `HERO-FIGURE-059`'s own instruction-level reads of the same mask, contrasted against `HERO-APPEAR-055`'s and `HERO-APPEAR-045`'s own reads of the different sex-bit mask; independently corroborated by the two branch files' own shipped English text, read directly on both preserved roots)

### TOWN-188

the global at `L03631` is referenced by one fixed absolute address from every tip call site read across this repository, in at least six structurally unrelated screens, which discharges `TOWN-021`'s "not independently confirmed" caveat for the school and extends `TOWN-025`'s 3-view corpus reading. Confirmed call sites, each read at instruction level, gate on the identical literal `L03631`: tavern (`TOWN-015`), school — this experiment reads `R0706`'s own gate directly at `L11771` and its filename push of `L11772` at `L11773`, resolving to `main\text\tips\training.txt` (`evidence/strdump-school-tip-filename-L11772.txt`), discharging `TOWN-021` — shop (`SHOP-TIP-045`), town (`TOWN-165`), and the character generator's two call sites (`TOWN-187`).

An image-wide scan of `R1293` (the tip-text reader every one of these calls into) finds 16 callers across 14 functions; the remaining callers not read in this experiment are two character-select screens and a per-mission battle-tip call, each a further structurally distinct screen referencing the same address (`evidence/scan-tipreader-all-callers.md`, `evidence/strdump-tipreader-caller-filenames.txt`). A single memory cell can hold only one value, so six or more independent call sites gating on the same literal address is direct, not corpus-agreement, evidence that the setting is one process-wide flag rather than a per-view field; this experiment does not by itself establish that every one of the further, unread call sites shares the exact same semantics as the six read here

**Confidence.** High for "one fixed address, referenced by every read call site" (read directly, no ambiguity in what a literal operand cites); Medium for the broader claim that the two remaining unread call sites (character select, battle tip) carry the identical semantics, since their own gating code was not read in this experiment

**Amended.** The clause naming the remaining reader callers is partially retracted: of the 16 call sites, 5 read non-tip nodes (`TEXT-112`). The flag-address conclusion stands, and the character-select and battle-tip gates are now read (`TOWN-518`, `TRIG-TIPS-087`). [`retracted.md`](retracted.md) holds the entry.

### TOWN-206

`TEXT-STRTAB-030` reads `R1773` to its own leaf: it is a four-byte array index, not a Win32 `STRINGTABLE` call, and its table `L09841` is `TEXT-STRTAB-023`'s own shared global line-pointer array; `main.txt` is that array's base-0 file. `tools/strtabprobe` reads `text/main.txt` directly out of each preserved root's `MAIN.RES`, splits it on `CR` matching `R0661`'s own scan (confirmed at instruction level by `TEXT-STRTAB-030`'s evidence), and prints the line at each requested index. Both roots hold 274 lines for this file, matching `TEXT-STRTAB-023`'s own count exactly. EN: id `0x7f` (line 127) = "Close", id `0x80` (line 128) = "Show tips next time".

RU: same two line numbers; the raw bytes decode to legible Russian words under CP866 ("Закрыть", "Показывать далее") and do not under CP1251, tried side by side on the same bytes — the evidence file carries the raw hex for both lines on both roots, not only the decode (`evidence/disasm-strtab-loader-R1773.txt`, `tools/strtabprobe/main.go`). **Falsifiable prediction, since no reader here renders a font glyph:** the RU install's own tip popup, viewed on screen or in a screenshot, shows a close button reading "Закрыть" and a checkbox label reading "Показывать далее"; refuted if either renders differently, or if the game applies a code page other than CP866 to this text at draw time

**Confidence.** High for the EN root's two lines and for both roots' raw bytes (direct read of the shipped file, corroborated by an independent line-count match against `TEXT-STRTAB-023`); Medium for the CP866 rendering as the words the game itself displays — a display-only decode not confirmed against a running install, scoped as the falsifiable prediction above

### TOWN-207

Control `0xe`, captioned "Close" (`TOWN-206`), is the popup's own close control, established by an executable binding and not by caption text alone: its own keyboard-message override recognizes `VK_RETURN` and dispatches through an owner/target pointer to a close-shaped call; control `0xd` carries no comparable code anywhere in its own vtable. `TOWN-211` establishes the base Control class's generic keyboard dispatch (`msg` `0x100`..`0x102`) ends, when unhandled by an instance field and by every child, in a fixed per-message vtable slot called with the message's own `wParam` as its only argument; for `msg==0x100` that slot is `+0x6c`. Control `0xe`'s own vtable (`L03644`) overrides `+0x6c` with `R0713`, a real function distinct from the shared no-op stub (`R1248`) every unmodified class carries at the same offset — the popup's own `+0x6c` (`R1988`) is one such stub, confirmed by direct read (returning zero with 4 argument bytes cleaned).

`R0713(this, key)`: if `this+0x30 != 0` and a virtual call through `vt+0x20` with `this` is nonzero (an enabled/visible-shaped gate), it compares `key` against `0xd`; on a match it pushes `0, 0, [this+0x70]` and calls `R1989` then `R1990` on the result — the shape of a dispatch through a stored owner/target pointer — and returns `1` (handled); otherwise it returns `0`. `0xd` (13) is the real Win32 `VK_RETURN` code, not either control's own construction id (which happen to also be `0xd` and `0xe`, a numeric coincidence this row keeps separate from the key comparison): `TOWN-211`'s own message-range reading already places `msg==0x100` in the keyboard family, and `claims/text.md`'s `TEXT-NAMEIN-024` (a different class, from `EXP-0143`) independently documents a structurally identical "hotkey matcher" pattern that explicitly names `0x0d` as one of two accepted key values.

Control `0xe`'s own vtable also overrides `+0x74` (the `msg==0x102`, `WM_CHAR`-family slot in the same dispatch) with `R0780` — the exact address `TEXT-NAMEIN-024` already documents, from its own reading, as "converts a keystroke only to compare it with `this+0x74` and stores nothing," corroborating that control `0xe` carries a real accelerator-character match (its own `~X~`-parsed caption accelerator, `TOWN-185`), not only the Enter-key path. Control `0xd`'s own vtable (`L06234`) overrides `+0x6c` instead with `R0731`, a `WM_KEYDOWN`-shaped handler keyed on the virtual-key-code range `0x21`..`0x28` (`VK_PRIOR`..`VK_DOWN`: PageUp, PageDown, End, Home and the four arrow keys), relaying a `msg=0x46d` notification to a target field on specific keys — unrelated to closing; its own `+0x74` is **not** overridden (shared stub `R1203`), so `0xd` carries no accelerator-character match at all, consistent with `TOWN-185`'s finding that it has no caption (`evidence/disasm-ctrl0d-vt6c-R1186-ctrl0e-vt6c-R0713.txt`, `evidence/disasm-ctrl0d-R0731-and-targets.txt`, `evidence/scan-ctrl-vtables-wide/SCAN_SUMMARY.md`, `evidence/scan-popup-vtable-wide/SCAN_SUMMARY.md`).

**Not established**: the exact mouse-click message/binding for `0xe` (only its keyboard path was traced to a close-shaped dispatch), and the concrete effect of `R1990`'s own indirect call through `[L03535]` (destroy vs. hide vs. post-message) beyond its own four-argument shape

**Confidence.** High that `0xe`, not `0xd`, is the close control (three convergent, independently-read instruments: caption text, the accelerator-character override, and the Enter-key-gated owner-dispatch override, none present on `0xd`); Medium for the exact closing mechanism (the dispatch shape is read, its final effect is not) and for a mouse click specifically (not traced; only the keyboard path was read)

### TOWN-208

The popup's own `vt+0x2c` (`TOWN-062`'s address for the tavern's shared "no own content" epilogue) is a generic three-step composite — compute a rect, dispatch through the instance's own `vt+0x30`, then walk children — and for the popup, `vt+0x30` draws a real tiled panel (`TOWN-209`); `TOWN-062`'s "no own content" reading is refined, not contradicted. `R1311(this)`, read in full: builds a local rect via `R1865`, calls `R1224(this, &rect, this+8)`, makes a **virtual** call through `vt+0x30` with `(this, &rect)`, then calls `R1991(this)` directly (not virtual). `R1991`, also read in full, is the generic child-paint walk: a virtual call through `vt+0x20` with `(this, 0x20)` gates the whole routine (nonzero skips all children — a hidden/disabled-shaped test); otherwise it iterates `i = 0..count-1`, `count` from `R1868(this+0x1c)`, fetches each child via `R1869(this+0x1c, i)`, and for each one calls the **child's own** `childVtable+0x2c` with `(child)` and no further arguments.

This is the structural mirror of `TOWN-139`'s already-published `R0642`, which walks children calling each child's own `vt+0x48` for message dispatch; `R1991` is the same shape for paint. `TOWN-062` read this same shared address for the tavern and found no drawing of its own; this row's own read shows `R1311` always also dispatches through the instance's own `vt+0x30` first — a real, non-trivial background/border draw for the popup (`TOWN-209`), and evidently a no-op for the tavern, whose own `vt+0x30` was not read in this experiment (`evidence/disasm-popup-paint-R1311.txt`, `evidence/disasm-popup-fillborder-476ae-childwalk-R1991.txt`)

**Confidence.** High (both functions read in full, every branch and call target resolved to an address; the parallel to `R0642` is a structural comparison against that function's own already-published, independently-read shape)

### TOWN-209

Read in full: it builds a local rect via `R1865` from the popup's own fields, calls `R1224(this, &localRect, this+8)`, shrinks that rect by `8` on its own right and bottom fields, builds a second, fresh rect, and issues two calls shaped like a clip-save/clip-set pair (`R0372` on the fresh rect, then `R0373` on the caller-supplied paint rect), followed by `R0346()` (no arguments). It then issues a sequence of virtual calls through one fixed global object at `[L03604]` (`TOWN-210`), through that object's own `vt+0x18` and `vt+0x1c` slots, each call passing rect-relative x/y coordinates and one of the small integer tile ids `9`, `0xa`, `0xb`, `0xc`, `0xd`, `0xe`, `0xf`, `0x10`, `0x11` (9 through 17 decimal); three loops compute repeat counts along the rect's width and height by integer division against the pitch constants `0x30` (48) and `0x20` (32).

It closes with `R0312()` (no arguments, ~~clip-restore-shaped~~ *the back surface's DirectDraw Unlock, `DLG-DIM-013` and `AI-CURSOR-225`; `R0346` above it is the matching Lock*) and a second `R0373` on the fresh rect. No `PUSH` of a string or path literal appears anywhere in this function's own instructions (`evidence/disasm-popup-fillborder-476ae-childwalk-R1991.txt`)

**Confidence.** High for the mechanical shape (every instruction read); Medium for the interpretation "tiled 9-slice panel/border draw" — inferred from the loop, pitch and tile-id structure, not confirmed by rendering the bitmap itself

**Amended.** [`retracted.md`](retracted.md) holds an entry for this claim (superseded).

### TOWN-210

`ScanField all:L03604` finds 57 read references across exactly 4 functions (`R1057` itself, `R1058`, `R0760`, and the popup's own `R1975`) and exactly 1 write reference, at `L06204` inside `R1057`: `[L03604] = R0753(<path>)`, guarded by a preceding `0x24`-byte allocation. The pushed string operand at static address `L06205` reads, byte for byte off a raw dword dump (not off Ghidra's auto-generated symbol name), `graphics\interface\lm.256` (25 bytes plus a NUL terminator). `R1057` is a large routine that constructs a whole series of shared UI-graphics globals the same way, each from its own path literal — the assignment immediately preceding this one builds a neighboring global from `graphics\interface\inv1024r.bmp` — so the global at `L03604` is one entry in a startup sequence, not something the tip popup or its own class constructs (`evidence/scan-canvas-L03604/SCAN_SUMMARY.md`, `evidence/scan-canvas-string/SCAN_SUMMARY.md`)

**Confidence.** High (the write-site and read-site counts are from a whole-image reference scan; the path is read directly off its own byte dump)

### TOWN-211

The base Control class's generic message handler (`TOWN-139`'s already-named `vt+0x48` default, `R0390`) checks three message-id ranges numerically matching the real Win32 keyboard (`0x100`-`0x102`), mouse (`0x200`-`0x206`) and `WM_USER`-based (`0x400`+) families, falling back through a per-instance handler field, then a children broadcast, before a final per-message vtable slot; for the tip popup both per-instance handler fields are zero, so every such message reaches the children broadcast. `R0390(this, msg, wParam, lParam)`, read in full: for `msg` in `[0x100,0x102]`, if `this+0x38` is non-null it calls `[this+0x38`'s own vtable `+0x48]`; if that field is null, or its call returns `0`, it calls `TOWN-139`'s `R0642(this, msg, wParam, lParam)` (the children broadcast); if that also returns `0`, it dispatches by `msg` to a fixed per-message vtable slot with the original `wParam` as the slot's only argument — `0x100`→`+0x6c`, `0x101`→`+0x70`, `0x102`→`+0x74`.

`msg` in `[0x200,0x206]` or `msg==0x400` follows the identical shape gated on `this+0x34` instead, with final slots `0x200`→`+0x4c`, `0x400`→`+0x50`, `0x201`..`0x206`→`+0x54`/`+0x58`/`+0x5c`/`+0x60`/`+0x64`/`+0x68` in order. Any other `msg` calls `R0642` unconditionally, with no final-stage fallback. The base Control constructor `R0773` zeroes both `this+0x34` and `this+0x38` unconditionally (`TOWN-185`'s own evidence file); no instruction in the popup's own constructor chain (`R0759`, `R1974`) writes either field, so for the tip popup both remain zero and every keyboard-, mouse- or `WM_USER`-range message reaches the children broadcast before any per-instance or class-level final-stage slot is consulted (`evidence/disasm-basedispatch-R0390.txt`)

**Confidence.** High for the dispatch shape and the popup's own zeroed fields (both read at instruction level); Medium for the Win32-constant-family reading of the message-id ranges — a numeric coincidence, not confirmed against how these values are produced or delivered

**Amended.** see TOWN-406 and retracted.md. [`retracted.md`](retracted.md) holds an entry for this claim (partially retracted).

## Room composition and character generator stages

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-214 | The tavern's button/decoration panel (`R1904`, vtable `L08008+0x2c`, id `0x44e`, matching `TOWN-064`) blits seven archive nodes, all under `graphics\interface\Inn\`… | High / Medium | ● active (partially retracted) | [EXP-0199](../experiments/EXP-0199-room-composition/) |
| TOWN-215 | Refines `TOWN-061` steps 2-3: the mage-side/fighter-side lower-area fields `+0xf4`/`+0x1a8` ("or a fill") are populated from frame-indexed pointer arrays at `+0xe4`/`+0x198`, not a single stored field… | High | ● active (partially retracted) | [EXP-0199](../experiments/EXP-0199-room-composition/); [retraction note](retracted.md#school-training-consumer-corrections) |
| TOWN-216 | The character generator's precreate stage loads 26 archive nodes through one routine, `R1672`, which is a loader and not the stage's paint routine. | High | ● active (amended) | [EXP-0199](../experiments/EXP-0199-room-composition/) |
| TOWN-217 | The character generator's final/detailed stage loads 48 archive nodes across two loaders; one of the four routines prior prose named for this stage, `R0833`, is confirmed a real paint routine… | High / Unknown | ● active (amended) | EXP-0199, [EXP-0201](../experiments/EXP-0201-chargen-destinations/) |
| TOWN-222 | `TOWN-214`'s paint-time origin addend for the tavern button panel is the panel's parent pointer's own stored LEFT/TOP (`this+8`/`+0xc`)… | High / Medium | ● active | [EXP-0200](../experiments/EXP-0200-paint-destinations/) |
| TOWN-223 | The character generator's precreate stage is painted by `R1474` (vtable `L07691` slot `+0x2c`); the hit-test/click handler `R1472` first read as that slot belongs to the name field class… | High / Medium / Unknown | ● active (amended, partially retracted) | [EXP-0200](../experiments/EXP-0200-paint-destinations/) |
| TOWN-224 | `TOWN-217`'s unresolved background coordinate call, `R1224`, is a general ancestor-chain rectangle accumulator, not the plain rectangle-copy this row originally read — corrected by `TOWN-232`. | High / Unknown | ● active (amended, partially retracted) | EXP-0200, [EXP-0201](../experiments/EXP-0201-chargen-destinations/) |
| TOWN-232 | `R1224`/`R1212` is a general ancestor-chain rectangle accumulator, not a plain copy as `TOWN-224` read it, and for all four of the final/detailed stage's own top-level children the chain is empty at construction… | High / Medium | ● active (amended, partially retracted) | [EXP-0201](../experiments/EXP-0201-chargen-destinations/), [EXP-0202](../experiments/EXP-0202-chargen-page-geometry/) |
| TOWN-233 | The final stage's `+0x74` and `+0x78` children (`TOWN-232`) paint through `R1992` and `R1993`, confirmed at vtable slots `L11774` and `L11775` (each child's own vtable `+0x2c`)… | High / Medium / Unknown | ● active (amended) | [EXP-0201](../experiments/EXP-0201-chargen-destinations/), [EXP-0202](../experiments/EXP-0202-chargen-page-geometry/) |
| TOWN-234 | The final stage's `+0x7c` child (`TOWN-232`, absolute rect `(160,0)-(480,480)`) paints through `R1871`, confirmed at vtable slot `L11776`… | High / Medium / Unknown | ● active | [EXP-0201](../experiments/EXP-0201-chargen-destinations/) |
| TOWN-235 | `TOWN-224`'s unresolved icon-loop virtual dispatch in `R0833` (the final stage's `+0x70` stats-panel child) is `vt+0x24`/`vt+0x20` called on the icon object itself — `GetHeight()`/`GetWidth()`… | High / Unknown | ● active | [EXP-0201](../experiments/EXP-0201-chargen-destinations/) |
| TOWN-236 | `TOWN-223`'s precreate-stage Levels (3-slot, `+0x124`) and Heroes (4-slot, `+0x110`) loops read POINTERS to an external rectangle table, not embedded rectangle arrays: `+0x124`/`+0x110` hold addresses… | High / Unknown | ● active (partially retracted) | [EXP-0201](../experiments/EXP-0201-chargen-destinations/) |

### TOWN-214

The tavern's button/decoration panel (`R1904`, vtable `L08008+0x2c`, id `0x44e`, matching `TOWN-064`) blits seven archive nodes, all under `graphics\interface\Inn\`, all through the opaque primitive `vt+0x18` (`TOWN-150`'s first primitive) — no call in this routine reaches `vt+0x38`. The panel's constructor (`R1534`, one of two non-destructor writers of vtable `L08008`; `R1535` is the parameterized sibling, corrected by `TOWN-391`) calls `R1809`, which writes twelve numeric fields and zeroes seven pointer fields (`+0x74`,`+0x78`,`+0x7c`,`+0x80`,`+0x84`,`+0x88`,`+0x8c`). A separate loader, `R1994`, fills those seven fields by the same three-call idiom (allocate, construct-from-path through `R1176`, register through `R0370`) seen elsewhere in this ledger (`TOWN-181`): `+0x74`=`button1on.bmp`, `+0x78`=`button2on.bmp`, `+0x7c`=`button3on.bmp`, `+0x80`=`button1off.bmp`, `+0x84`=`button2off.bmp`, `+0x88`=`button3off.bmp`, `+0x8c`=`ButtonsArea.bmp`.

The paint routine, gated on `*(int*)(*(int*)(this+0x5c)+0x13c)!=0`, blits `ButtonsArea.bmp` once, then the middle button (`+0x78`/`+0x84`, chosen by `*(this+0xc4)==1 && *(this+0xc0)==1`) at the rectangle `(484,91)-(624,137)`, then loops the other two buttons (`+0x74`/`+0x80` at `(484,44)-(624,90)`, `+0x7c`/`+0x88` at `(484,138)-(624,184)`, chosen by `*(this+0xc4)==slot`), each with a centred caption draw through the shared text primitive `[L06186]+0x14`; the upper and lower buttons additionally draw separate numeric strings when their caption is nonempty (`TOWN-392` corrects the former digit identification). The three rectangles are LTRB, not XYWH: the left/right fields (`0x1e4`=484, `0x270`=624) are identical across all three, and only the top/bottom fields vary, in three non-overlapping increasing bands — a shape only an LTRB reading produces.

The panel's own construction call (`R1535(0x44e,0x1e0,0,0x280,0xee,parent)`, `EXP-0186`'s evidence, `(id 0x44e, left 480, top 0, right 640, bottom 238)`) gives the panel's own absolute left/right edges as 480/640, four and sixteen pixels outside the button rectangles' shared left/right — consistent with the button rectangles already being absolute screen coordinates under the game's default resolution, where `TOWN-222` traces the paint routine's own runtime origin addend (`*(this+0x5c)+8`/`+0xc`, added at every blit call site) to the tavern's own stored LEFT/TOP and finds it zero under the default 640x480 configuration, nonzero (`(80,60)` or `(192,144)`) when an explicit alternate-resolution flag is matched at startup (`evidence/decomp-R1904-tavern-button-panel-paint.c`, `evidence/decomp-R1534-tavern-button-panel-ctor.c`, `evidence/decomp-R1809-tavern-button-panel-init.c`, `evidence/vt-L08008-tavern-button-panel-vtable.md`, `evidence/disasm-R1994-tavern-button-loader.txt`, `evidence/strings-tavern-button-icons.txt`, `evidence/enumrefs-callto-R1904.txt`, `evidence/enumrefs-refto-L08008.txt`)

**Confidence.** High for the seven paths, the primitive (all four button-object blit call sites in the routine read directly, none targets `vt+0x38`), and the stored rectangle values (constructor read to raw disassembly); High for the rectangles being the final absolute screen position under the game's default resolution, and Medium for which resolution is active during actual play, per `TOWN-222`'s trace of the paint-time origin addend

**Amended.** constructor exclusivity and digit identification only; `TOWN-391`, `TOWN-392`. [`retracted.md`](retracted.md) holds an entry for this claim (refuted).

### TOWN-215

Refines `TOWN-061` steps 2-3: the mage-side/fighter-side lower-area fields `+0xf4`/`+0x1a8` ("or a fill") are populated from frame-indexed pointer arrays at `+0xe4`/`+0x198`, not a single stored field; the arrays' own archive-path source was searched in three places and found in none. `R1962` (step 2) and `R1963` (step 3), called from `R0706`, each index a pointer array and store the result: loading the array pointer from offset 0xe4 of the view, loading the element at 4 bytes times the frame counter and storing it at offset 0xf4 of the view, and the same shape at `+0x198`→`+0x1a8`, with the index a frame counter advanced in the same two routines, gated by `+0x31c`/`+0x320`/`+0x344` in the same shape as the diamond-animation gate `TOWN-154` already reads at `+0x25c`/`+0x264`.

**Negative finding, three instruments.** Searched: (1) the paint routine `R1487`, full decompile; (2) the constructor `R1940`, full decompile (both already committed, `EXP-0195`'s `evidence/disasm-R1487-paint-R1940-ctor.txt`); (3) a fresh full 943-line raw disassembly of the bitmap loader `R1486`, every instruction in the function (`evidence/disasm-R1486-school-loader-full.txt`, this experiment). Method: text search for the literal displacements `0xe4]` and `0x198]` as a store target in each. Result: zero occurrences in all three. Conclusion: the pointer arrays feeding steps 2-3 are not populated by the school's own paint routine, its own constructor, or its own bitmap loader; where they are populated is not established

**Confidence.** High for the indexing mechanism (both counter routines read at instruction level); source conclusion partially retracted by `TOWN-427`

**Amended.** callee writers resolve the archive source; see `TOWN-427` and the retraction note. [`retracted.md`](retracted.md) holds an entry for this claim (partially retracted).

### TOWN-216

`R1672` (307 decompiled lines) issues, for each of 26 nodes, the same allocate/construct-from-path/register idiom `TOWN-214` reads in the tavern; a text search of its full decompiled body for the blit-primitive call shapes (`+ 0x18))`, `+ 0x34))`, `+ 0x38))`) finds none. The 26 nodes, read with a new helper (`PrintStrings.java`, this experiment): `PreCreate\{MainArea,Mask,Amulet,ButtonOk}.bmp`, `PreCreate\Blind\sprites.16a` (a sprite archive, not a flat bitmap), `PreCreate\Heroes\{mf,mm,ff,fm}{on,l,lon}.bmp` (4 sex/class combinations x 3 states), `PreCreate\Levels\level{0,1,2}{on,l,lon}.bmp` (3 levels x 3 states). The stage's paint routine is now found and its full blit list read: `TOWN-223`. The stage's own vtable is `L07691`, 34 slots (the entry referencing `R1475` is its `+0x80`, one `EnumRefs callto:R1475` hit): `R1473` stores it at `L07692` (`TEXT-073`), its `+0x2c` slot is the paint routine `R1474`, and the stub `R0531` the earlier row named sits at `+0x0c`. `R1472`, a hit-test/click handler (`PtInRect` then `R1885`), is a slot of the neighbouring name field class's table (`+0x54` of `L07729`, `TEXT-075`) (`evidence/decomp-R1672-chargen-precreate-loader.c`, `evidence/strings-chargen-precreate.txt`)

**Confidence.** High for the node list (loader read to full decompile, absence of blit calls confirmed by text search over the whole body)

**Amended.** The table identity, `+0x2c` slot and stub-offset clauses, corrected in the text above; the node list and the loader-is-not-paint finding stand. [`retracted.md`](retracted.md) holds the entry.

### TOWN-217

The character generator's final/detailed stage loads 48 archive nodes across two loaders; one of the four routines prior prose named for this stage, `R0833`, is confirmed a real paint routine, blitting its own background through `vt+0x18` and four looped icons through `vt+0x38` — the same two primitives `TOWN-150` names, extended here to a new screen. `R1995` loads `main\graphics\chrgen\leftup.bmp` plus ten `graphics\interface\chrgen\buttons\{m,p}{disable,nloff,nlon,loff,lon}.bmp` nodes (two buttons, minus/plus, five states each). `R0133` (13,505 decompiled bytes) loads 37 more: `FullStatsR.bmp`/`RollStatsR.bmp`, `chrgen\fighter\{Bow,Pike,Mace,axe,sword}\{shine_on,shine_off,on}.bmp` (5 weapon classes x 3 states) plus `fighter\column.bmp`/`fighter\mask.bmp`, `chrgen\mag\{air,astral,earth,fire,water}\{shine_on,shine_off,on}.bmp` (5 magic schools x 3 states) plus `mag\column.bmp`/`mag\mask.bmp` — mirroring the school's own fighter/mage column-and-mask convention (`TOWN-149`) closely enough to be the same authoring pattern, though the class names differ (`Mace` here, `club` in `TOWN-149`) and this experiment did not verify whether the file sets are identical — and one further node, `graphics\interface\inn\RUOver.bmp`.

That last path was read off `R0133`'s own instruction operand, the chargen final-stage loader; it sits in the tavern's own asset folder by path, but the tavern's own paint routine was not searched for a reference to this file in this experiment, so this is reported as where the path was found, not as a blit the tavern also issues. `R0833`'s background blit reads `this+0x60` (`leftup.bmp`) through `vt+0x18`, dest `(x,y)` from an unresolved call `R1224(&x,&y,this+8)`; its 4-iteration icon loop blits through `vt+0x38`. `R1870` loads no bitmaps — four identifiers (`Multiplayer`/`FacesMM`/`FacesMF`/`FacesFM`/`FacesFF`) and two tip-text paths — and is not a paint node.

`R1871` and `R1993` (the nav buttons and one skill-choice routine named by prior prose) contain no string literal, consistent with consuming pointers `R0133` already loaded rather than loading their own; neither was read past that check — this finding is about the loaded-path question only and does not extend to whether either routine draws text: `TOWN-233` reads `R1993`'s full body and finds a per-icon text-label draw whose string comes from a runtime table, not a literal. The background coordinate call is resolved by `TOWN-224`. Destinations for the other 47 nodes remain unresolved (`evidence/decomp-R0833-chargen-final-statspanel-paint.c`, `evidence/decomp-R1870-chargen-final-facepool.c`, `evidence/strings-chargen-final.txt`)

**Confidence.** High for the node list and the two primitives (both loaders read to full decompile, both blit call sites in `R0833` read directly); explicitly Unknown for the 47 remaining destinations, and for whether `RUOver.bmp` is also blitted by the tavern

**Amended.** clarified by EXP-0201.

### TOWN-222

`TOWN-214`'s paint-time origin addend for the tavern button panel is the panel's parent pointer's own stored LEFT/TOP (`this+8`/`+0xc`), read once at the top of `R1904` and added to every blit and text-draw coordinate in that routine; the parent is the tavern object itself, and the tavern's own stored LEFT/TOP is zero under the game's default 640x480 configuration, making `TOWN-214`'s published rectangles the final absolute screen destination as written — nonzero only when an explicit alternate-resolution flag is matched at startup. The panel's constructor (`R1535`) writes its `+0x5c` field directly from its sixth argument: `param_1[0x17] = param_7`.

The tavern's own child-construction routine, `R1939`, calls `R1535(0x44e,0x1e0,0,0x280,0xee,param_1)` where `param_1` is `R1939`'s own `this` — the tavern object — so the panel's `this+0x5c` is the tavern itself. `R1904` (the panel's paint routine, already read by `TOWN-214`) reads `iVar1=*(this+0x5c)+8`, `iVar2=*(this+0x5c)+0xc` once at entry and adds them into every blit and text-draw coordinate that follows, including the panel's own background blit (`iVar5=*(this+8)+iVar1`) and every button rectangle's stored corner. The tavern object is constructed by `R1996` (`EnumRefs refto:L08008`, 3 hits: `R1997`, a variant with zero callers per `EnumRefs callto:R1997`; `R1996`; `R1998`, the destructor), called exactly once, from `R0315` at `L11777`, with `id=0x44c` and left/top/right/bottom pushed from the globals the global at `L09808`/`L09809`/`L09810`/`L09811`.

Those four globals are written exactly once each in the whole image (`ScanField wr:`, one hit apiece), inside `R0341`: `[L09808]=([L00618]-0x280)/2`, `[L09809]=([L01259]-0x1e0)/2`, where the global at `L00618`/`L01259` (screen width/height) default to `0x280`/`0x1e0` (640/480) unless a string compare matches a literal resembling a command-line switch — `"-800"` (`L13168`) sets 800x600, `"-1024"` (`L13169`) sets 1024x768; the source of the compared string was not traced past the routine that supplies it (`R1877`), so whether it is argv or another source is unresolved. `R0341` is called exactly once (`EnumRefs callto:R0341`), from `R0326` at `L09309`.

Under the default path (no matching flag), `[L09808]=[L09809]=0`; under `"-800"`, `(80,60)`; under `"-1024"`, `(192,144)`. **Absolute button destinations.** Adding this addend to `TOWN-214`'s three stored rectangles: under the default resolution, unchanged — `(484,44)-(624,90)`, `(484,91)-(624,137)`, `(484,138)-(624,184)`; under `"-800"`, shifted by `(80,60)` to `(564,104)-(704,150)`, `(564,151)-(704,197)`, `(564,198)-(704,244)`; under `"-1024"`, shifted by `(192,144)` to `(676,188)-(816,234)`, `(676,235)-(816,281)`, `(676,282)-(816,328)`. Rectangles are LTRB, the reading `TOWN-184`/`TOWN-214` already fix by the non-overlapping-band argument (`evidence/R1535_R1535.c`, `evidence/R0773_R0773.c`, `evidence/R1939_R1939.c`, `evidence/enumrefs-refto-L09234.txt`, `evidence/R1997_R1997.c`, `evidence/enumrefs-callto-R1997.txt`, `evidence/R1996_R1996.c`, `evidence/R1998_R1998.c`, `evidence/enumrefs-callto-tavern-ctors.txt`, `evidence/disasm-R0315-tavern-construct-call.txt`, `evidence/scan-tavern-rect-globals/`, `evidence/strings-resolution-detect.txt`, `evidence/enumrefs-callto-R0341.txt`, `evidence/R0388_R0388.c`, `evidence/disasm-R0388-setrect-forwarder.txt`, `evidence/R0759_R0759.c`, `../EXP-0199-room-composition/evidence/decomp-R1904-tavern-button-panel-paint.c`)

**Confidence.** High for the addend mechanism and its two operand fields (every blit/text-draw call site in `R1904` read directly, the parent-pointer write read at the constructor's own instruction, the tavern's single construction call read to raw disassembly, the four globals' single write site each confirmed by a whole-image scan); High for the default-path value being zero (the constants and the subtraction that produce it are read directly); Medium for which resolution is active during actual play, since the source of the compared flag string was not traced past `R1877`

### TOWN-223

The character generator's precreate stage is painted by `R1474` (vtable `L07691` slot `+0x2c`). `R1472`, the hit-test/click handler (`PtInRect` then `R1885`) this claim first read as that table's `+0x2c`, is the name field class's slot `+0x54` (`TEXT-075`). `R1474` was found by decompiling every remaining slot of the vtable that carries the stage's enter-handler `R1475` at `+0x80` (the same table `TOWN-216` cites) and searching each for indirect-call shapes; it is the only slot whose body issues blit-primitive calls (`+0x18`, `+0x38`) against the loader's own fields. Ruled out along the way: `R1672` (the loader, `TOWN-216`); `R0598` (a release routine, called from the loader's own top, frees `+0x70`/`+0x68`/`+0x6c`/`+0x160`/`+0x15c` and the two group arrays before reload); `R1472`, read in the scan window as the table's own `+0x2c` slot (`PtInRect((RECT*)(this+8),pt)` then `R1885(this,1)` on a hit — a click handler, not paint; the slot belongs to the name field class's table, `TEXT-075`); `R1879` (a region-code hit-test dispatcher, sets `+0x184`/`+0x188` to copies of `+0x15c`/`+0x160` on hovering regions `0xa0`/`0xb4`); `R1999` (a `timeGetTime()`-gated hover-highlight blink state machine, one `vt+0x38` call on a sub-rectangle of `+0x15c`/`+0x160`); and two further large slots with no blit-shaped indirect call in their decompiled body (`R1878`, 18033 bytes; `R0597`, 12038 bytes, 0 indirect calls).

**Full ordered blit list.** (1) Background: `+0x68` (`MainArea.bmp`) via `vt+0x18` at `(this+8,this+0xc)`, unconditional. (2) Hover highlight, Amulet: `+0x184` (a copy of `+0x15c`=`Amulet.bmp`, set by `R1879` while hovering region `0xa0`) via `vt+0x38`, sub-rect `(+0x164,+0x168)-(+0x16c,+0x170)` offset by `(this+8,this+0xc)`, gated on `+0x184 != 0`. (3) Hover highlight, ButtonOk: `+0x188` (copy of `+0x160`=`ButtonOk.bmp` on region `0xb4`) via `vt+0x38`, sub-rect `(+0x174,+0x178)-(+0x17c,+0x180)`, gated on `+0x188 != 0`. (4) Levels loop, 3 iterations: each slot's state at `+0x14c[slot]` (1/2/3) selects `+0xd4`(1)/`+0xe8`(2)/`+0xfc`(3) — `+0xd4`=`level{0,1,2}on.bmp`, `+0xe8`=`level{0,1,2}l.bmp`, `+0xfc`=`level{0,1,2}lon.bmp` — via `vt+0x38`, dest `(this+8+rect.left,this+0xc+rect.top)`, size `(rect.right-rect.left,rect.bottom-rect.top)`, rect read from the 4-dword array `+0x124[slot]`; literal rect values not traced.

(5) Heroes loop, 4 iterations (array slots 0,1,2,3 = files `mf`,`ff`,`fm`,`mm`): per-slot state at `+0x138[slot]` selects `+0x98`(1)/`+0xac`(2)/`+0xc0`(3) — `+0x98`=`{mf,ff,fm,mm}on.bmp`, `+0xac`=`{mf,ff,fm,mm}l.bmp`, `+0xc0`=`{mf,ff,fm,mm}lon.bmp` — same dest formula, rect array `+0x110[slot]`; literal rect values not traced. (6) Steps 2-3 repeated identically, immediately after the two loops — a second draw of the same two hover highlights, not explained further here. (7) A tip-popup call gated on `[L03631]!=0 && +0x1c0!=0`, delegating into the blink routine already ruled out above; no new blit read at this site. (8) `+0x70` (`Blind\sprites.16a`) via `vt+0x18`, dest `(+0x78+this+8,+0x7c+this+0xc)`, frame index `+0x74`, gated on `+0x74!=0` — an animated, intermittently-hidden decoration whose position and 63-tick frame advance are timer-driven, not a fixed node.

(9) A text/digit draw through the shared primitive `[L06186]+0x14`, dest tied to object `+0x1bc`, offset `(+8,+0xc-5)` from that object plus `(this+8,this+0xc)`. `R1474` ends with `R1311()` (`TOWN-208`'s shared epilogue: `vt+0x30` then a generic child-walk), so whether a child object or `vt+0x30` paints anything further is not established here. This routine reads `this+8`/`this+0xc` directly with no further parent-pointer indirection, unlike the tavern's button panel (`TOWN-222`), which is a child widget; this class's own rect was not traced back to its own construction call, so whether `this+8`/`this+0xc` here is itself `(0,0)` is open.

**Negative finding on one node.** `+0x6c` (`Mask.bmp`, loaded through a different constructor call, `R1152`, than the other 25 nodes' `R1176`) is never read by this routine. Searched: `R1474`'s full decompiled body and `R1999`'s full body, both for the literal displacement `+ 0x6c`. Zero occurrences in either. Conclusion: `Mask.bmp` is not painted by the stage's own paint routine or its hover-blink helper; its consumer was not located (`evidence/R1672_R1672.c`, `evidence/findvt-L07693-L13170.txt`, `evidence/enumrefs-callto-R1672.txt`, `evidence/R1475_R1475.c`, `evidence/disasm-R1475-precreate-caller-start.txt`, `evidence/R1472_R1472.c`, `evidence/R1474_R1474.c`, `evidence/R1879_R1879.c`, `evidence/R0598_R0598.c`, `evidence/R1999_R1999.c`, `evidence/R0599_R0599.c`, `evidence/R2000_R2000.c`, `evidence/R1878_R1878.c`, `evidence/R0597_R0597.c`, `evidence/strings-precreate-full.txt`, `evidence/enumrefs-disp-precreate-fields.txt`, `evidence/enumrefs-disp-precreate-15c-160.txt`)

**Confidence.** High for the routine identity, the +0x2c correction, and the ordered blit list, primitives and gating conditions (every blit call site in `R1474` read directly, the vtable slots read from one direct scan); Medium for the archive-node-to-file mapping's internal consistency (the `on`/`l`/`lon` suffix pattern repeats identically across the Heroes and Levels groups — an internal cross-check, not an external confirmation); explicitly Unknown for every rectangle's literal numeric value, for whether `vt+0x30`/a child paints anything further, and for why the two hover highlights are drawn twice

**Amended.** The vtable base, slot index and enter-slot clauses. The scan read a 59-slot window from `L07693`, which starts inside the name field class's table (`L07729`, 30 slots, ending before `L07691`) and runs 5 slots into the table at `L11778`. The stage's own table starts at `L07691` and has 34 slots; `R1473` stores it at `L07692` (`TEXT-073`). Paint is `+0x2c` and enter is `+0x80`. The routine identity, blit list, gates, confidence grades and Unknowns stand. [`retracted.md`](retracted.md) holds the entry. The role given to `R1999` and item (7) of the blit list are partially retracted: the routine is the guided-step cycle of `TOWN-519`, and item (7) is its call, not a popup call. `TOWN-520` answers the Mask.bmp and rectangle Unknowns. [`retracted.md`](retracted.md) holds the entry.

### TOWN-224

`TOWN-217`'s unresolved background coordinate call, `R1224`, is a general ancestor-chain rectangle accumulator, not the plain rectangle-copy this row originally read — corrected by `TOWN-232`. For this call site the ancestor chain is empty at construction, so the final/detailed stage's background blit destination is still exactly the caller's own stored rect's LEFT/TOP, `(this+8,this+0xc)`, taken with no addend — because the chain has no ancestors here, not because the routine is incapable of one. Raw disassembly of the call site in `R0833` (`L11779`-`L11780`) shows the address `this+8` and the address of a local stack buffer each formed and pushed, `this` passed as the object, and `R1224` called — three real arguments (`this` as the object, the local buffer, `this+8`), where `TOWN-217`'s own evidence names only two (`&x,&y,this+8`), an artifact of the decompiler under-reporting the callee's two-stack-argument (8 argument bytes cleaned on return) signature.

Inside `R1224` (raw disassembly, this experiment), the two stack arguments are passed through `R1866` (`return param_1`, identity) and `R1867` (`return param_1+8`), then into `R1212(this0,dst,src)`, which performs `*dst=*src; dst[1]=src[1]` — a 2-dword copy — called twice, once on the identity pair and once on the `+8`-offset pair, together copying 4 dwords (16 bytes) from `src=this+8` into `dst=&localbuf`. The local buffer is `TOWN-217`'s own `(x,y)`: its first two dwords are read immediately afterward as the arguments to the background's `vt+0x18` blit. `TOWN-232` reads the rest of `R1212`'s body past the 2-dword copy this row stopped at: it also walks an ancestor-chain field (`this+0x30`) and adds each ancestor's own stored LEFT/TOP via `R1213`, so the routine performs arithmetic whenever the chain is non-empty.

`TOWN-232` finds this call site's own chain empty at construction (the shared `R0773` unconditionally zeroes `this+0x30` for this child), so the walk executes zero iterations here and the background destination is unaffected: exactly `(this+8,this+0xc)` — the object's own stored rect's LEFT/TOP, LTRB order fixed by the same convention `TOWN-184`/`TOWN-214`/`TOWN-222` already establish for `this+8..+0x14`. `R0833` sits in orphan/undisassembled territory (Ghidra's own marker on the containing range), so its own construction call was not traced, and whether this class's `this+8`/`this+0xc` is an absolute screen position or itself needs a parent addend of the kind `TOWN-222` finds for the tavern's button panel is not established here.

**The 4-iteration icon loop.** Each iteration reads a per-slot record at `this+0xb4` (stride 0x30), draws a text label via `R0571` at an origin built from register values the decompiler names `unaff_EBP`/`unaff_EBX` (consistent with, but not confirmed as, `this+8`/`this+0xc` cached across the earlier call, given the orphan-code caveat above), then this row read it as resolving an icon object through two virtual-call indirections (`(**(iVar7+0x24))(0)` then `(**(piVar1+0x20))(0,uVar2)`) before blitting it through `vt+0x38`, with the object reached by that indirection not traced past this point. `TOWN-235` resolves this same call site and corrects the shape: each row draws one text label plus TWO icon blits (`OBJ1` at `this+0x174+8k`, `OBJ2` at `this+0x178+8k`), not the one icon this row read, and the two virtual calls it names are `GetHeight()`/`GetWidth()` called on the icon object itself, not an indirection reaching a second, untraced object.

`TOWN-235` gives every literal coordinate in the loop — the label and both icon blits, across all 4 rows. **Negative finding on the remaining 43 nodes.** Searched: `R0833` in full (this experiment and `TOWN-217`'s prior read) — one background blit and one 4-iteration loop, at most 9 blit call sites, none reaching the ten button-state nodes (`R1995`) or the 37 weapon/magic/column/mask nodes (`R0133`). No further paint routine for those was located: this experiment did not repeat, for the final stage's own vtable, the systematic per-slot decompilation that found `TOWN-223`'s precreate paint routine. Whether `R0833`'s icon loop covers the button-state nodes, the weapon/magic icons, or neither is not established (`evidence/R1224_R1224.c`, `evidence/disasm-R1224-coord-conv.txt`, `evidence/R1866_R1866.c`, `evidence/R1867_R1867.c`, `evidence/R1212_R1212.c`, `evidence/disasm-R0833-range.txt`, `../EXP-0199-room-composition/evidence/decomp-R0833-chargen-final-statspanel-paint.c`)

**Confidence.** High for the background destination formula at this call site, unaffected by the correction (`TOWN-232`: the chain is empty here) / the general-accumulator-vs-plain-copy clause carried High while believed and is refuted by `TOWN-232` / the icon-loop dispatch-and-one-icon-per-row clause carried Unknown while believed and is resolved by `TOWN-235`; Unknown stands for all 43 remaining nodes' destinations (not searched beyond `R0833` in this experiment)

**Amended.** [`retracted.md`](retracted.md) holds 2 entries for this claim (refuted).

### TOWN-232

`R1224`/`R1212` is a general ancestor-chain rectangle accumulator, not a plain copy as `TOWN-224` read it, and for all four of the final/detailed stage's own top-level children the chain is empty at construction, so each child's absolute screen rectangle equals the literal LTRB values passed to its own constructor. `R1212(this,dst,src)` performs `dst[0]=src[0]; dst[1]=src[1]` (the 2-dword copy `TOWN-224` reads), then walks `local_8=*(int*)(this+0x30)` and, for every non-null link, calls `R1213(dst, ancestor[0xc], ancestor[8])` before advancing `local_8=*(int*)(local_8+0x30)`; raw disassembly of the call site confirms the exact push order (the value at offset 0xc of the ancestor pushed, then the value at offset 0x8, `dst` passed as `this`, and `R1213` called).

`R1213(int*p,int dx,int dy)` is `*p+=dx; p[1]+=dy;`. `this+0x30` is a distinct field from the `this+0x5c` parent pointer `TOWN-222`/`TOWN-223`/`TOWN-224` already read for the paint-visibility gate (`*(this+0x5c)+0x104 != 0`); `R1224` calls this walk twice, once on `(this+8,this+0xc)`→`dst[0..1]` and once on `(this+0x10,this+0x14)`→`dst[2..3]`, so both LTRB corners accumulate the same ancestor chain. **Amended by `EXP-0202` (`TOWN-242`): `R1313` is never invoked; the top-level chargen-final screen's constructor genuinely reached at runtime is `R1314`.** Both share vtable `R1312` and both call `R1215`, so the four children's own construction rectangles, vtable identities and empty ancestor chain below are unaffected.

**The final stage's own four children.** `R1313` (originally read here as the top-level chargen-final screen constructor, vtable `R1312`) calls `R1215`, which constructs four children in order: `R1214(0x457,0,0,0xa0,0xee,this)` at `this+0x70` (vtable `L06316`), `R2001(0x458,0,0xee,0xa0,0xf2,this)` at `+0x74` (`L11781`), `R2002(0x459,0x1e0,0,0x280,0xee,this)` at `+0x78` (`L11782`), `R2003(0x45a,0xa0,0,0x1e0,0x1e0,this)` at `+0x7c` (`L11783`) — vtable identities confirmed by matching each constructor's own `*param_1=&PTR_LAB_xxxxxx` write. Each constructor forwards its five geometry arguments unchanged into `R0773(id,left,top,right,bottom,0)`, which stores them via `R0388(left,top,right,bottom)`→`SetRect(rect,left,top,right,bottom)` (LTRB, `evidence/R0388_R0388.c` from `EXP-0200`), so the literal construction rectangles are (LTRB): `+0x70`=`(0,0)-(160,238)`, `+0x74`=`(0,238)-(160,242)`, `+0x78`=`(480,0)-(640,238)`, `+0x7c`=`(160,0)-(480,480)`.

All four constructors also call `R0773(...,0)` with a literal `0` sixth argument (not the parent), and `R0773` itself unconditionally sets `param_1[0xc]=0` (`=*(this+0x30)`, the ancestor-chain field the accumulator walks) before the `if(param_7!=0)` block that would otherwise register a real parent — this is read identically in the decompiled body of `R1214`, `R2001`, `R2002`, `R2003`. So each child's own ancestor chain is empty at construction, and `R1224`'s ancestor-walk loop executes zero iterations for each: the accumulated LTRB equals the literal construction rectangle unmodified, for the three children whose paint routines call `R1224` at all (`TOWN-233`, `TOWN-235`; the fourth, `+0x78`, uses a different mechanism, `TOWN-233`).

Whether anything writes `this+0x30` for these objects later at runtime, outside this construction path, was not searched — a whole-image `disp:30` sweep returns hits in dozens of unrelated classes at this common offset and was not narrowed further in this experiment (`evidence/R1224_R1224.c`, `evidence/disasm-R1224.txt`, `evidence/R1866_R1866.c`, `evidence/R1867_R1867.c`, `evidence/R1212_R1212.c`, `evidence/disasm-R1212.txt`, `evidence/R1213_R1213.c`, `evidence/R1313_R1313.c`, `evidence/R1215_R1215.c`, `evidence/R1214_R1214.c`, `evidence/R2001_R2001.c`, `evidence/R2002_R2002.c`, `evidence/R2003_R2003.c`, `evidence/R0773_R0773.c`, `evidence/enumrefs-callto-R1215.txt`, `evidence/enumrefs-callto-R1214-bc30.txt`)

**Confidence.** High for the accumulator mechanism and the ancestor-chain field identity (every instruction in `R1212`/`R1213` and the call site read directly); High for the four children's construction rectangles and vtable identities (each constructor's own literal arguments and `PTR_LAB` write read directly); High for the ancestor chain being empty at construction for all four (the shared `R0773`'s unconditional zero-write read directly, repeated identically in all four constructors' decompiled bodies); Medium for whether `this+0x30` is written later at runtime by code outside this construction path, narrowed from Unknown by `TOWN-243`; the original identification of `R1313` as the invoked top-level constructor carried High while believed and is retracted by `TOWN-242`

**Amended.** [`retracted.md`](retracted.md) holds an entry for this claim (refuted).

### TOWN-233

The final stage's `+0x74` and `+0x78` children (`TOWN-232`) paint through `R1992` and `R1993`, confirmed at vtable slots `L11774` and `L11775` (each child's own vtable `+0x2c`); `R1993` draws a centered text label per icon, correcting the scope of `TOWN-217`'s "no string literal" observation about this routine. `R1992` (`+0x74`, absolute rect `(0,238)-(160,242)` per `TOWN-232`): one unconditional background blit, `vt+0x18(0,238,0,0,0)` via `this+0x60`, followed by an undecompiled call `R0877(&stack)` passing three literal small values (`0,0,0xc`) built after the blit's own coordinates are already consumed — not resolved further; no visibility gate on this routine (unlike its siblings).

`R1993` (`+0x78`, absolute rect `(480,0)-(640,238)`): gated on `*(this+0x5c)+0x104 != 0`; one background blit, `vt+0x18(*(this+8)+iVar1, *(this+0xc)+iVar2, 0,0)` via `this+0x8c`, where `iVar1=*(iVar6+8)`/`iVar2=*(iVar6+0xc)` and `iVar6=*(this+0x5c)` — a single-hop parent addend using the top-level screen's own stored LEFT/TOP, structurally the mechanism `TOWN-222` finds for the tavern's button panel, not the `TOWN-232` ancestor-chain accumulator (`R1993` never calls `R1224`); the top-level chargen-final screen's own `(this+8,this+0xc)` was not traced to its own construction call in this experiment, so the absolute value of this addend is open (by the tavern's precedent, `TOWN-222`, it is plausibly `(0,0)` under default resolution, but this is not confirmed for this screen).

**Amended by `EXP-0202` (`TOWN-242`): the top-level screen's own construction call is `R1314`, and under the default resolution the same globals `TOWN-222` reads for the tavern give this addend as `(0,0)`, confirmed rather than merely plausible.** Then a 3-iteration loop over `this+0x80[i]` (icon-state selector) and `this+0x94[i]*4` (per-slot rect record): each iteration blits one icon (normal state `puVar5[0]` or highlighted state `puVar5[-3]`, selected by `*(unaff_ESI+0xc4)` against the loop index) via that icon's own `vt+0x18(iVar1+piVar4[-1], iVar2+*piVar4, 0,0,0)`, then draws a centered text label via the shared primitive `*[L06186]+0x14` at `(iVar1+piVar4[-1]+(piVar4[1]-piVar4[-1])/2, iVar2+*piVar4+(piVar4[2]-*piVar4)/2 [+1 if highlighted])`, text source `*(int*)(*(int*)(iVar3+100)+(int)puVar5)` — the text pointer's own source table was not traced further.

This directly extends `TOWN-217`'s reading: `TOWN-217` found no string literal in `R1993`'s decompiled body and stopped there, correctly noting it had not read past that check; this experiment reads the full body and finds a per-icon text-label draw whose string comes from a runtime table, not a literal, so `TOWN-217`'s narrow finding (no embedded string constant) stands, but its inference (consistent with drawing no text) does not extend past the point `TOWN-217` itself flagged as unread (`evidence/R1992_R1992.c`, `evidence/R1993_R1993.c`, `evidence/enumrefs-callto-R1992.txt`, `evidence/enumrefs-callto-R1993-R1871.txt`)

**Confidence.** High for both routines' vtable-slot identity and blit/text-draw call structure (every call site in each routine's decompiled body read directly, cross-referenced against each child's own constructor `PTR_LAB` write); High that `R1993` draws text (the shared primitive and its five-argument shape match the convention `TOWN-223` already establishes); Medium for the icon loop's exact per-slot semantics (the state-selector and highlight-offset logic is read but not exercised at runtime); High for the top-level screen's own absolute position, resolved by `TOWN-242`; Unknown for `R0877`'s purpose and for the text-source table at `iVar3+100`

**Amended.** The amendment is stated in the claim text above.

### TOWN-234

The final stage's `+0x7c` child (`TOWN-232`, absolute rect `(160,0)-(480,480)`) paints through `R1871`, confirmed at vtable slot `L11776`: one background blit and four 16px-wide border-column blits along its own left and right edges, followed by a 5-slot, 3-state icon loop — raw-disassembly-verified because the decompiler renders the border coordinates as `unaff_EBP` and small negative literal offsets. `R1871` calls `R1224(&local,this+8)` (`TOWN-232`: ancestor chain empty, so `local`= this child's own literal LTRB `(160,0)-(480,480)`, i.e. `L=160,T=0,R=480,B=480`), gated on `*(this+0x5c)+0x104 != 0`. Raw disassembly of `L11784`-`L11785`, tracking each stack-offset read against the four-dword local buffer's fixed stack offsets, gives: background `vt+0x18(L,T,0,0,0)` via `this+0x60`; border `this+0x68`: `vt+0x38(L,T,0,0,0x10,0xee)`; border `this+0x6c`: `vt+0x38(L,T+0xee,0,0,0x10,0xf2)`; border `this+0x70`: `vt+0x38(R-0x10,T,0,0,0x10,0xee)`; border `this+0x74`: `vt+0x38(R-0x10,T+0xee,0,0,0x10,0xf2)` — numerically, with `L=160,T=0,R=480`: `(160,0,0,0,16,238)`, `(160,238,0,0,16,242)`, `(464,0,0,0,16,238)`, `(464,238,0,0,16,242)` — two 16px columns (left edge at `x=160`, right edge at `x=464`), each split into a 238px and a 242px vertical strip (summing to the child's own 480px height), at the panel's own left and right screen edges.

Then a 5-iteration loop (`this+0xdc`/`this+0xa0` arrays, stride 8/4): each iteration selects one of three icon-object pointers per slot (`puVar3[-10]`/`puVar3[-5]`/`*puVar3`, chosen by `puVar3[0x1b]`∈{1,2,3}) and blits it via that icon's own `vt+0x38(puVar2[-10],puVar2[-9],0,0,*puVar2,puVar2[1])` — a per-slot stored rect, not derived from the border/background coordinates. Followed by an undecompiled `R2004()` call, not read in this experiment. `TOWN-217` named this routine only to note it has no string literal and was not read past that check; this reading is the first full trace (`evidence/R1871_R1871.c`, `evidence/disasm-R1871.txt`, `evidence/enumrefs-callto-R1993-R1871.txt`)

**Confidence.** High for the vtable-slot identity, the background and four border blits' exact literal coordinates (full raw-disassembly stack-offset trace from function entry to each indirect call through slot +0x38 of the object's table, cross-checked against the fixed local-buffer layout `TOWN-232` establishes); Medium for the 5-slot icon loop's per-slot rect source and 3-state selection (the loop structure and call shape are read directly, but the per-slot data's own writer was not traced); Unknown for `R2004`'s purpose

### TOWN-235

`TOWN-224`'s unresolved icon-loop virtual dispatch in `R0833` (the final stage's `+0x70` stats-panel child) is `vt+0x24`/`vt+0x20` called on the icon object itself — `GetHeight()`/`GetWidth()`, the same pair `TOWN-208`'s bitmap-wrapper class already establishes — not a dispatch chain reaching some other object; the loop draws one text label plus two icon blits per row, across 4 rows, all at literal coordinates now resolved. `R0833` (absolute rect `(0,0)-(160,238)` per `TOWN-232`) opens with `R1224` (chain empty, so `L=0,T=0`), background `vt+0x18(0,0,0,0,0)` via `this+0x60`, then a 4-iteration loop (`this+0x1d0[k]`, `this+0xb4+0x30k`, `this+0x174+8k`/`this+0x178+8k`, stride confirmed by the loop's own `ADD` increments at `L11786`/`L11787`/`L11788`).

Per row `k` (raw-disassembly stack-slot trace of `L11789`-`L11790`): (1) a value `*(this+0x1d0+4k)` is formatted via `R0567(L11791, L02664, value)` into the shared buffer at `L11791` (the format/semantics of this call were not read further), then drawn as a text label via `R0571(X,Y,textPtr,0,uiObjField,1)` at `X=L+*(this+0xb4+0x30k)+2`, `Y=T+*(this+0xb4+0x30k+4)+4`; (2) icon `OBJ1=*(this+0x174+8k)` blitted via `OBJ1->vt@0x38` called with `(L+*(this+0xb4+0x30k+0x10),T+*(this+0xb4+0x30k+4),0,0,w,h)` where `w,h` come from `OBJ1->vt@0x20` and `OBJ1->vt@0x24`, each called with `(0)` (`GetWidth`/`GetHeight`, called in that raw-disasm order, `L11792`/`L11793`); (3) icon `OBJ2=*(this+0x178+8k)` blitted the same way at `X=L+*(this+0xb4+0x30k+0x20)`.

The three X-anchors and the shared row Y come from a 12-dword table `R1210` (this class's own field-init routine, called from all three `L06316` constructors — `R2005`, `R1214`, `R2006`) writes to `this+0xc4+0x30k` (`=` `this+0xb4+0x30k+0x10`): three sub-rectangles per row, left edges `82`,`107`,`132` (right edges `102`,`127`,`152`, unused by this routine except the first, which anchors the label's `+2`), row top/bottom `(iVar1_k, iVar1_k+0x14)` for `iVar1_k=0x36,0x56,0x76,0x96` (`54,86,118,150`) — row spacing `0x20` (`32`). With `L=T=0`: label positions `(84,58)`,`(84,90)`,`(84,122)`,`(84,154)`; `OBJ1` blit positions `(107,54)`,`(107,86)`,`(107,118)`,`(107,150)`; `OBJ2` blit positions `(132,54)`,`(132,86)`,`(132,118)`,`(132,150)`.

`this+0xb4`/`this+0xb8` (the row's own first two dwords, read as the label's X/Y offsets) are written by `R1210` as part of the same 12-dword row block. `R1210` also writes four literal fields read elsewhere in `R0833`'s tail (`this+0xa4=0x2e`, `this+0xa8=0xb5`, `this+0xac=0x7b`, `this+0xb0=0xcb`), not traced further in this experiment (`evidence/R0833_R0833.c`, `evidence/disasm-R0833-full.txt`, `evidence/R1210_R1210.c`, `evidence/R2007_R2007.c`, `evidence/R2008_R2008.c`, `evidence/R1176_R1176.c`)

**Confidence.** High for the two virtual calls' identity as `GetWidth`/`GetHeight` and for all twelve resolved coordinates (full raw-disassembly stack-offset trace of the whole loop body, cross-checked against `R1210`'s own literal writes); Unknown for what the sprintf-like call at `L11791` formats and for the four tail fields' consumer

### TOWN-236

`TOWN-223`'s precreate-stage Levels (3-slot, `+0x124`) and Heroes (4-slot, `+0x110`) loops read POINTERS to an external rectangle table, not embedded rectangle arrays: `+0x124`/`+0x110` hold addresses, dereferenced and offset by `0x10`/slot inside `R1474`; the table's writer was not read here. Raw disassembly of `R1474`'s two loops (`disasm-R1474-full.txt`) shows, for each, the table pointer `*(this+0x124)` (or `+0x110`) loaded as a value, then `rollingOffset` added to it and a dereference of the result at `+0`,`+4`,`+8`,`+0xc` for the rect's four corners, with 0x10 added to the slot pointer advancing it between slots — a pointer-to-table read, refining `TOWN-223`'s "4-dword array" phrasing for these two fields.

A whole-image scan for the literal displacement `0x124`/`0x110` (`ScanField disp:`) returns 71 and 50 hits across dozens of unrelated functions (both offsets are common small struct fields reused throughout the binary); restricted to reads inside `R1474` itself, this confirms the loop's own access pattern but does not identify the table's writer, since no other function in either scan's hit list shares precreate's own vtable family. **Withdrawn negative on precreate's construction call.** Searched: `ScanField imm:L07693`, `all:L07693`, `sym:L07693` — zero hits in all three modes. `L07693` is not the class's vtable base (`L07691`, stored three times, `TEXT-073`), so the zero hits say nothing about the class's construction; the final stage's `L06316` vtable, found at three constructor sites by the same method (`TOWN-232`/`TOWN-233`), is a separate class.

One candidate, `R1473` (also read this experiment: writes fields at `+0x110`/`+0x124` via 32-bit stores into the fields at offsets 0x110 and 0x124 of the object at `L11794`/`L11795`), is the pre-create screen's own constructor, not a sibling class: it stores the screen's vtable `L07691` at `L07692`, and its one caller (`EnumRefs callto:R1473`) is `R0315` at `L07694`, which builds `main+0x360` (`TEXT-073`); `L11778` is the vtable of the embedded objects it constructs at `+0x80`, `+0x10c` and `+0x120`. Conclusion: the rectangle table's writer was not located in this experiment; the literal rectangle values for the Levels and Heroes loops remain unresolved (`evidence/disasm-R1474-full.txt`, `evidence/scanfield-disp-124-110.md`, `evidence/scanfield-precreate-ctor-negative.md`, `evidence/scanfield-final-child-ctor.md`, `evidence/R1473_R1473.c`, `evidence/disasm-R1473-ruledout.txt`, `evidence/enumrefs-callto-R1473-ruledout.txt`)

**Confidence.** High for the pointer-indirection reading of `+0x124`/`+0x110` (the dereference and stride are read directly from raw disassembly); the negative findings on `R1473` and on precreate's construction call are withdrawn (Amended); explicitly Unknown for the table's writer and its literal values

**Amended.** The base `L07693` and two negatives are withdrawn: `L07693` is not a vtable base, and `R1473`, ruled out here as a sibling class, is the pre-create screen's constructor (`L07691` at `L07692`, built at `L07694`; `TEXT-073`). The pointer-indirection reading of `+0x124`/`+0x110` in `R1474` stands as read from the loop; `TEXT-073` reads both fields as null at construction inside embedded objects, so their writer is a later store. [`retracted.md`](retracted.md) holds the entry.

## Character generator pages, widget painters and interface census

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-242 | The final/detailed chargen screen's constructor genuinely reached at runtime is `R1314`, not `R1313` (`TOWN-232`); its construction rectangle under the default resolution is `(0,0)-(640,480)`… | High / Medium | ● active | [EXP-0202](../experiments/EXP-0202-chargen-page-geometry/) |
| TOWN-243 | Across the full 65-function method surface of the final/detailed stage's five class vtables, only two functions access the ancestor-chain field `this+0x30` (`TOWN-232`) with register-relative addressing, and both only read it… | Medium | ● active | [EXP-0202](../experiments/EXP-0202-chargen-page-geometry/) |
| TOWN-244 | A raw byte-level dword scan of every PE section finds zero occurrences of `L07693`, which is not a vtable base: the pre-create class's table starts at `L07691`… | High / Unknown | ● active (partially retracted) | [EXP-0202](../experiments/EXP-0202-chargen-page-geometry/) |
| TOWN-245 | The final/detailed stage's top-level screen draws nothing of its own — its own paint routine delegates entirely to the four children's paint walk (`TOWN-232`/`233`/`234`/`235`)… | High / Unknown | ● active | [EXP-0202](../experiments/EXP-0202-chargen-page-geometry/) |
| TOWN-246 | `TOWN-015`'s own derived width for the town's tip-popup rectangle is wrong (should read 312, not 640); `TOWN-184`'s own correction of that error names the wrong room (the tavern, not the town). | High | ● active | [EXP-0202](../experiments/EXP-0202-chargen-page-geometry/) |
| TOWN-252 | Each of the final/detailed chargen stage's four graphics-loading sub-calls in `R1870` runs on a different CHILD object as `this`, never on the top-level screen's own `this` — the premise `TOWN-245`'s own negative was scoped against… | High | ● active | [EXP-0203](../experiments/EXP-0203-node-painters/) |
| TOWN-253 | `R0133`, run on the final stage's `+0x7c` child, writes the exact fields `TOWN-234` left open at Medium confidence ("the per-slot data's own writer was not traced")… | High | ● active | [EXP-0203](../experiments/EXP-0203-node-painters/) |
| TOWN-254 | `R1995` (default init) and `R0835` (runtime hover/press selector), both run on the final stage's `+0x70` stats-panel child, together write `this+0x174`/`this+0x178`… | High | ● active | [EXP-0203](../experiments/EXP-0203-node-painters/) |
| TOWN-255 | `R0877` (10,911 decompiled bytes) draws no archive node: every one of its calls goes through the text/number primitive `R0571(x,y,value,kind,fontTable,count)`, none through a blit vtable slot… | High | ● active | [EXP-0203](../experiments/EXP-0203-node-painters/) |
| TOWN-256 | No reader of the final stage's `+0x7c` child's mask-bitmap field (`this+0x64`, `fighter\mask.bmp`/`mag\mask.bmp` per `TOWN-253`) was found among that child's own known methods. | Medium | ● active | [EXP-0203](../experiments/EXP-0203-node-painters/) |
| TOWN-257 | `TOWN-183`'s `+0x10` school background-blit addend is a literal immediate in the instruction stream, present only in the school's own paint routine… | High | ● active | [EXP-0203](../experiments/EXP-0203-node-painters/) |
| TOWN-258 | The school's two training-button hit-rects `TOWN-182` names at `+0x88`/`+0x98` (16-byte stride) have literal values `(484,71)-(624,117)` and `(484,117)-(624,163)` (LTRB), each 140×46, both lying entirely inside `x:[480,640)`… | High | ● active | [EXP-0203](../experiments/EXP-0203-node-painters/) |
| TOWN-259 | Every node `R1905` (`TOWN-183`) draws into the school's widget rect `(464,0)-(640,238)` is accounted for; none reaches the rect's own left 16 pixels, `x:[464,480)`, and no node covering that strip was found. | High | ● active (amended, partially retracted) | [EXP-0203](../experiments/EXP-0203-node-painters/), **[EXP-0205](../experiments/EXP-0205-town-page-composition/)** |
| TOWN-260 | The shop's own widget-rect background is drawn through the KEYED primitive (`vt+0x38`), unlike school's and tavern's opaque (`vt+0x18`) backgrounds… | High / Medium | ● active | [EXP-0203](../experiments/EXP-0203-node-painters/) |
| TOWN-261 | The object `TOWN-181` names only as "`[school_view+0x78]`" and the object `TOWN-182`/`TOWN-183`'s own `this` are the same object — the school room's own `+0x78` child… | High | ● active | [EXP-0203](../experiments/EXP-0203-node-painters/) |
| TOWN-264 | Every `graphics.res` `interface/` archive entry the executable references by path is counted on both preserved roots, and the corrected instrument reproduces an independently published subtree count exactly. | High | ● active | [EXP-0204](../experiments/EXP-0204-interface-census-and-panel-text/) |
| TOWN-265 | The 77 EN-root / 75 RU-root unreferenced `interface/` entries (`TOWN-264`) fall into seven groups, five of them a loose per-frame `.bmp` file beside a packed sprite container of a related name the code does reference by path… | High / Unknown | ● active | [EXP-0204](../experiments/EXP-0204-interface-census-and-panel-text/) |
| TOWN-280 | The school training widget's origin is `(0,0)` at the game's default 640x480 resolution, which resolves both of `TOWN-183`'s undecoded `vt+0x14` label draws — now known to be four call sites, two per button, not two… | High / Unknown | ● active | [EXP-0205](../experiments/EXP-0205-town-page-composition/) |
| TOWN-281 | The school's 640x480 page in draw order, assembled from previously published claims plus this experiment's own resolution of the label draws (`TOWN-280`): a 480-wide main content area painted by the room's own routine… | High | ● active | [EXP-0205](../experiments/EXP-0205-town-page-composition/) |
| TOWN-282 | The tavern's 640x480 page, assembled from previously published claims: three fixed-rect siblings tiling `x:[0,480)`, plus the same borrowed right-column lower panel the school and shop use, in the SAME 160-pixel right column. | High | ● active (partially retracted) | [EXP-0205](../experiments/EXP-0205-town-page-composition/) |
| TOWN-283 | The character generator's two 640x480 pages, assembled from previously published claims: the precreate stage (`R1474`, 26 archive nodes) and the final/detailed stage (top-level screen `R1314`… | High / Unknown | ● active (amended) | [EXP-0205](../experiments/EXP-0205-town-page-composition/) |
| TOWN-284 | 165 `graphics.res` `interface/` entries are classed `literal` (`TOWN-264`/`EXP-0204`) and sit directly under `interface/` or under `interface/inn/`/`interface/chrgen/`: 67 direct, 17 under `inn/`, 81 under `chrgen/`. | High | ● active | [EXP-0205](../experiments/EXP-0205-town-page-composition/) |
| TOWN-266 | The same three shared routines `TOWN-063` already reads for the four top-level room views are also what the tavern's own left-panel child's vtable resolves to at `+0x34`, `+0x38` and `+0x40`… | High | ● active | [EXP-0204](../experiments/EXP-0204-interface-census-and-panel-text/) |

### TOWN-242

The final/detailed chargen screen's constructor genuinely reached at runtime is `R1314`, not `R1313` (`TOWN-232`); its construction rectangle under the default resolution is `(0,0)-(640,480)`, resolving `TOWN-233`'s open fourth-child addend to `(0,0)`, confirmed. `R1313` and `R1314` both write vtable `R1312` and both call `R1215` (`TOWN-232`'s four-children constructor), but differ in shape: `R1313` takes no explicit id/rect arguments and calls `R2009()` (a no-argument base-constructor variant); `R1314` takes five explicit arguments (`id,left,top,right,bottom`) forwarded unchanged into `R0759`, the same LTRB-setting chain `TOWN-184`/`TOWN-232` already trace.

Three independent instruments against `R1313` — `EnumRefs callto:`, `refto:`, `imm:` — each return zero hits in the whole image; `R1313` also does not appear in any of the 240 combined slots of the final stage's five relevant vtables (`evidence/vt-final-stage-classes.txt`). `EnumRefs callto:` against `R1314` returns exactly one hit: `L06735`, inside `R0315`, matching `TOWN-140`'s own independent citation of "the fourth view['s]... constructor `R1314`, called from the campaign screen constructor `R0315` at `L06735`" exactly. Decompiling `R0315` directly (not read by `TOWN-232`) shows the call itself: `R1314(0x456, [L09808], [L09809], [L09810], [L09811])`, result stored at `campaign+0x358` — the same offset and id `TOWN-140` names for the fourth view.

`TOWN-232`'s own evidence list cites `evidence/R1313_R1313.c` but neither `R1314` nor `TOWN-140`; `R1313` is an uncalled sibling overload, and `TOWN-232`'s identification of it as "the" invoked constructor is corrected. `TOWN-232`'s downstream findings (the four children's own construction rectangles, vtable identities, and empty ancestor chain) are unaffected, since both candidate functions call the identical `R1215`. Under the default (640×480, no alternate-resolution flag matched) resolution path `TOWN-222` already establishes for the same four globals (re-verified directly against `R0341`'s own decompiled body, `evidence/R0341_R0341.c` from `EXP-0200`): `[L09808]=[L09809]=0`, `[L09810]=640`, `[L09811]=480`.

So the top-level screen's own construction rectangle (LTRB) is `(0,0)-(640,480)`: `this+8=0` (LEFT), `this+0xc=0` (TOP). This resolves `TOWN-233`'s open item: the `+0x78` child's single-hop parent addend, read via `*(this+0x5c)+8`/`+0xc`, is `(0,0)` under default resolution, by the same mechanism and the same four globals `TOWN-222` already reads at High confidence for the tavern — confirmed, not merely plausible by precedent. Which real-world screen this class presents remains supported only by content correlation (the weapon/magic/level/hero iconography `TOWN-217`/`TOWN-223`/`TOWN-224`/`TOWN-233`/`TOWN-234`/`TOWN-235` already read); no string resource, transition call from precreate, or other independent tie was found in this experiment, so this identification is not strengthened beyond what those prior experiments already establish (`evidence/R1313_R1313.c`, `evidence/R1314_R1314.c`, `evidence/R0315_R0315.c`, `evidence/enumrefs-callto-R1313-R1314.txt`, `evidence/enumrefs-refto-R1313.txt`, `evidence/enumrefs-imm-R1313.txt`, `evidence/vt-final-stage-classes.txt`)

**Confidence.** High for the constructor identity and `R1313` being uncalled (three named instruments individually resolved, cross-validated against `TOWN-140`'s independently-derived citation of the same address and call site); High for the top-level construction rectangle under default resolution (identical evidence class `TOWN-222` already carries at High); Medium for which real-world screen the class presents (content correlation only, unchanged from prior experiments)

### TOWN-243

Across the full 65-function method surface of the final/detailed stage's five class vtables, only two functions access the ancestor-chain field `this+0x30` (`TOWN-232`) with register-relative addressing, and both only read it — narrowing `TOWN-232`'s open question of whether the field is repopulated at runtime to a scoped negative. The five vtables `TOWN-232`/`TOWN-233`/`TOWN-234`/`TOWN-235` read (`L06316`, `L11781`, `L11782`, `L11783`, `R1312`; 48 slots each) list 65 distinct function addresses once duplicates (slots inherited unchanged across vtables) are collapsed. Filtering the whole-image `EnumRefs disp:0x30` sweep (1206 hits, 456 distinct owner functions) against exactly this list of 65 names finds matches only for `R0776` (vtable slots `+0x24`/`+0x9c`/`+0xac` across the five vtables) and `R0777` (`+0x28`/`+0xa0`/`+0xb0`) — 13 hits total, all reads: a comparison of the 32-bit field at `[reg+0x30]` with zero or a load of a register from that field.

No 32-bit store to `[reg+0x30]` (a write) appears among them; no other of the 65 functions accesses `+0x30` with a register-relative operand (the sweep's remaining hits against these 65 names are all stack-relative locals at +0x30 (stack-pointer or frame-pointer relative), excluded as unrelated). Raw disassembly of both functions in full (`R0776`-`L11796`, `R0777`-`L11797`) confirms every `+0x30` access is a read; both instead write to `+0x34`/`+0x40` (`R0776`) or `+0x38`/`+0x44` (`R0777`) — fields distinct from the ancestor-chain pointer. Combined with `TOWN-232`'s own reading (`R0773` unconditionally zeroes `this+0x30` at construction, the only write found anywhere), within the reachable virtual-method surface of these five classes `this+0x30` is written exactly once, at construction, to zero, and is never rewritten.

Scope: this instrument covers the direct vtable-slot method surface only. It does not rule out a non-virtual member function of these classes absent from every vtable slot, an external class holding a raw pointer to one of these objects and writing `+0x30` directly, or an instruction `EnumRefs disp:` cannot see because Ghidra's own analysis never disassembled it (`docs/INSTRUMENT.md`'s blind-spot note on `disp:` modes: exposed to orphan code, blind to code never disassembled at all) (`evidence/vt-final-stage-classes.txt`, `evidence/enumrefs-disp-30.txt`, `evidence/disasm-R0776-R0777.txt`)

**Confidence.** Medium for "`this+0x30` is not repopulated at runtime within the reachable vtable surface" — a reproducible whole-corpus filter against a named, complete function list, but scoped to virtual-method reachability and to code Ghidra's own analysis has disassembled, not a proof against every code path

### TOWN-244

A raw byte-level dword scan of every PE section, immune to Ghidra's own disassembly coverage, finds zero occurrences of `L07693`, which `TOWN-223` named as precreate's vtable base and which is not one (the table starts at `L07691`, stored three times, `TEXT-073`); independently re-checking `TOWN-236`'s `+0x124`/`+0x110` write-hit owners adds no new candidate for the Levels/Heroes rectangle table's writer. A raw little-endian dword scan of every PE section (`tools/chargenorigin`, no disassembler in the path) for `L07693` finds 0 hits across 5 sections. Three control addresses confirm the scan's own correctness: `R1312` (the final stage's top-level vtable) returns 3 hits, `L06316` returns 3 hits, and `R1472` (a slot of the name field class's table, `+0x54` of `L07729`, `TEXT-075`) returns exactly 1 hit, at `L07695`. The same scan for `L07691` returns 3 hits, at `L11798`, `L11799` and `L11800`.

This is a differently-blind-spotted negative than `TOWN-236`'s `ScanField imm:`/`all:`/`sym:` (bound by Ghidra's own disassembly and reference-manager coverage, `docs/INSTRUMENT.md`): the raw scan would find a vtable base literal even inside bytes Ghidra never disassembled, and still finds none for `L07693`. `EnumRefs disp:124`/`disp:110` (whole image) return 71/50 hits; filtering for register-relative write instructions gives 23 distinct owner functions (`evidence/write-hits-124-110.txt`). Two owners write both offsets in the same function: `R1473` (the pre-create constructor, `TEXT-073`; `TOWN-236` excluded it in error) and `R2010` (re-examined this experiment) — the latter's own decompiled body constructs an object whose vtable is `L11778`/`L11027`, a 9-slot embedded-sub-object pattern unrelated to precreate; its `+0x110`/`+0x124` writes are two of nine identical embedded-object zero-initializations, not a rectangle table.

One owner writes a non-register literal address, `R1652` (a 32-bit store of the constant `R2011` to the field at offset 0x110 of the object): decompiled and found to be an MFC thread-context class (calls `AfxGetThread()`, installs its own embedded-object vtables at `+0x2a`/`+0x32`/`+0x44`), unrelated to precreate. Neither newly-examined candidate belongs to precreate's own class family. Raw disassembly of `R1474`'s Levels/Heroes loops (`TOWN-236`) confirms both pointer dereferences execute unconditionally with no null check and no lazy-initialization branch: the table must already be populated by the time paint runs, and there is no fallback-to-default pattern to trace instead.

The remaining 21 of 23 write-hit owners were not individually decompiled in this experiment. Conclusion: the Levels/Heroes rectangle table's literal values remain unresolved, and the construction-call negative is withdrawn (Amended); the scan rules out two further specific candidates (`evidence/rawscan-vtable-bases.txt`, `evidence/write-hits-124-110.txt`, `evidence/R2010_R2010.c`, `evidence/R1652_R1652.c`, `experiments/EXP-0201-chargen-destinations/evidence/disasm-R1474-full.txt`, `tools/chargenorigin/main.go`)

**Confidence.** High for the raw-scan count of zero for `L07693`, which is not the vtable base (byte-level, immune to disassembly coverage, correctness verified against three known control addresses); High for the two additionally-ruled-out candidates (each decompiled directly, vtable and class shape read); explicitly Unknown for the table's literal values

**Amended.** The headline's vtable base and the negative on precreate's construction call are withdrawn: `L07693` is not a vtable base, and the class's table `L07691` has three stores, two constructors (`R1473` among them) and the destructor. The control-address reading of `R1472` is corrected above. The zero count, the two ruled-out candidates `R2010` and `R1652` and the access-pattern reading stand. [`retracted.md`](retracted.md) holds the entry.

### TOWN-245

The final/detailed stage's top-level screen draws nothing of its own — its own paint routine delegates entirely to the four children's paint walk (`TOWN-232`/`233`/`234`/`235`) — and a third graphics loader was found that `TOWN-217`'s node inventory did not count, raising the stage's own total from 48 to at least 55; the painter for the 54 archive nodes still unaccounted for (all but `leftup.bmp`, already resolved by `TOWN-217`) remains unresolved. `R1310` (`vt+0x2c` of the top-level screen, vtable `R1312`) matches the shared "no own content" epilogue pattern `TOWN-208` documents: it conditionally calls `R0320` (a state-check call, not a blit; the same call also appears at the tail of `R1870`, below), then unconditionally calls `R1311`, whose own-content dispatch is the instance's own `vt+0x30` — here `R2012`, an empty stub (`return;`, no body).

The top-level screen issues no blit of its own; every pixel of the final stage is drawn through the generic child-paint walk over the four children already read in full. `R2004` (the call `TOWN-234` leaves undecompiled at the tail of the `+0x7c` child's paint, `R1871`) is a real-time-gated (`timeGetTime`, ≥500ms) animation tick: when the top-level screen's own `+0x100==0` and `+0x80!=0` (the tip-popup pointer `R1870` sets, below), it selects one of 5 icon-state slots on the `+0x7c` child's own `+0x8c[]`/`+0xa0[]` arrays (the same arrays `TOWN-234`'s main loop already reads) and blits it via that icon's own `vt+0x38`. This resolves `TOWN-234`'s open item but recycles the `+0x7c` child's own already-inventoried icon slots; it draws none of the nodes below.

The top-level screen's own `vt+0x80` (`R1870`, `TOWN-217`'s "loads no bitmaps... not a paint node," and the routine `docs/DOORS-OPEN.md` names as the measurement nobody had made) is a one-time initialization routine, not a paint routine: `TOWN-217`'s characterization is accurate about its own direct instructions (no blit, no bitmap-load call issued directly) but incomplete about its role — it is the single dispatcher for six loader sub-calls: `R1995` (10 button-state nodes, `TOWN-217`), `R2013` (1 node, `FullStatsR.bmp`, into the top-level's own `+0x60`), **`R2014` (7 nodes: `graphics/interface/Inn/button{1,2,3}on.bmp`, `button{1,2,3}of.bmp`, `ButtonsArrows.bmp`, into `+0x74`..`+0x8c` — not in `TOWN-217`'s inventory)**, and `R0133` (37 nodes, `TOWN-217`).

The remaining two sub-calls, `R2015` and `R2016`, load sound effects (`.wav`), not graphics. The final stage's total archive-node count is therefore at least `48 + 7 = 55`, not 48. `R1870` also constructs the stage's own tip popup directly (`R1261(0x467,0,0x118,0x138,0x1e0,...)`, gated on the global at `L03631`, stored at the top-level's own `+0x80`) — the same construction primitive `TOWN-015`/`TOWN-184`/`TOWN-246` fix the operand order for — and loads four face-pool identifiers via `R1430` into `+0xa0`/`+0xb4`/`+0xc8`/`+0xdc`, matching `TOWN-217`'s "four identifiers" note. Negative, scoped: searched the full decompiled bodies of all four children's paint routines (`TOWN-224`/`233`/`234`/`235`, prior experiments) and the top-level's own paint dispatch and the `R2004` timer callback (this experiment) for any reference to the offsets the three graphics loaders write their loaded pointers into on the top-level object.

None found. Not conclusive: `R0133`'s own body, which stores the 37 nodes' own pointers and so names the offsets to search for, was not decompiled in this experiment (13,505 decompiled bytes per `TOWN-217`); and `R0877` (10,911 decompiled bytes, called from both `R1992`, the `+0x74` child, and `R2017`, itself a slot of the `+0x7c` child's own vtable at `+0xb4`) was decompiled but not read in this experiment, and could plausibly be a shared consumer of some of these nodes (`evidence/R1310_R1310.c`, `evidence/R2012_R2012.c`, `evidence/R2004_R2004.c`, `experiments/EXP-0199-room-composition/evidence/decomp-R1870-chargen-final-facepool.c`, `evidence/R2013_R2013.c`, `evidence/R2014_R2014.c`, `evidence/R2015_R2015.c`, `evidence/R2016_R2016.c`, `evidence/enumrefs-callto-loaders-and-undecomp.txt`, `evidence/R0877_R0877.c`)

**Confidence.** High for the top-level's own paint being empty and delegating entirely to the child walk (both routines fully decompiled, the epilogue chain matches `TOWN-208`'s already-documented shared pattern exactly); High for `R2004`'s role (fully decompiled, cross-checked against the fields `TOWN-234` already reads); High for the three-loader/two-sound-loader dispatch and the additional 7-node count (`R2014` decompiled in full, every load call and destination field read directly); explicitly Unknown for the painter of the 54 unresolved nodes

### TOWN-246

Re-deriving both call sites directly: the tavern's popup (`R1412`) calls `R1261(0x467, 0, 0, 0x138, 200, ...)`; by the LTRB order `TOWN-184` fixes (traced to the genuine Win32 `SetRect` import), this is `LEFT=0, TOP=0, RIGHT=0x138=312, BOTTOM=200` — width `=RIGHT-LEFT=312`, height `=BOTTOM-TOP=200`, position `(LEFT,TOP)=(0,0)`. `TOWN-015`'s own phrase for the tavern, "popup 312×200 at (0,0)," is correct. The town's popup (`R1383`) calls `R1261(0x467, 0x148, 0, 0x280, 200, ...)`, independently confirmed at raw-byte level by `TOWN-165`: `LEFT=0x148=328, TOP=0, RIGHT=0x280=640, BOTTOM=200` — width `=RIGHT-LEFT=640-328=312`, height `=200`, position `(328,0)`.

`TOWN-015`'s own phrase for the town, "popup 640×200 at (328,0)," states the wrong width: 640 is the RIGHT operand, not `RIGHT-LEFT`. The correct phrase is "popup 312×200 at (328,0)" (equivalently, 312×200 right-edge-anchored at `x=640`, since `RIGHT=640`= the screen width). This error is invisible for the tavern (`LEFT=0`, so `RIGHT-LEFT=RIGHT` trivially) and visible only for the town (`LEFT=328≠0`). `TOWN-184`'s own body already identifies this exact numeric error ("refutes `TOWN-015`'s derived '640×200 at (328,0)'... correct reading: 312×200, right-edge-anchored at x=640") but attributes the phrase to "the tavern's tip rect." Per `TOWN-015`'s own text quoted above, the "640×200 at (328,0)" phrase belongs to the TOWN's rectangle, not the tavern's; the tavern's own phrase ("312×200 at (0,0)") was never wrong and needed no correction.

A reader trusting `TOWN-184`'s room label as written would conclude the tavern's popup is anchored at `x=640` (position `(328,0)`) — the tavern's actual popup position is `(0,0)` — while never learning under the town's own name that the town's popup needed the width correction. `TOWN-165`'s own text ("rect (328,0,640,200)") is a raw tuple carrying no derived width/height phrase, and is unaffected; this matches `TOWN-184`'s own note on that point. Amended in place: `TOWN-015`'s town-popup clause, `TOWN-184`'s room attribution (`evidence` — no new Ghidra reads; re-derived from the operands and LTRB order `TOWN-165`/`TOWN-184` already publish)

**Confidence.** High (both call sites' raw operands and the LTRB order are already published at High by `TOWN-165` and `TOWN-184` respectively; this experiment's own contribution is arithmetic, width `=right-left`, checked against those already-High operands, and the room attribution is read directly from each row's own text)

### TOWN-252

Each of the final/detailed chargen stage's four graphics-loading sub-calls in `R1870` runs with a different CHILD object as `this`, never on the top-level screen's own `this` — the premise `TOWN-245`'s own negative was scoped against, which is why that search could not have found a reader regardless of whether one exists. Raw disassembly of `R1870` (`L11801`-`L11802`) shows four consecutive loads of the child pointer (the `this` of the next call) from the dword at offset 0x70/0x74/0x78/0x7c of the screen object, each immediately followed by a call to `R1995`/`R2013`/`R2014`/`R0133`: `R1995` runs with `this` = the `+0x70` child (`TOWN-235`'s stats-panel object, `R0833`'s own `this`); `R0133` runs with `this` = the `+0x7c` child (`TOWN-234`'s object, `R1871`'s own `this`).

`TOWN-245`'s own negative searched "the offsets the three graphics loaders write their loaded pointers into on the top-level object" and found no reference in the top-level's own paint dispatch or in `R2004`; since neither loader writes the top-level object at all, that search's target offsets never existed on the object it searched, independent of whether a reader exists on the object the loaders actually write. `TOWN-253` and `TOWN-254` trace the two loaders' actual writes to the two children's own already-published paint routines (`evidence/disasm-R1870.txt`, already on file from `EXP-0201`/`EXP-0202`, re-read this experiment)

**Confidence.** High (four consecutive pairs of a load of the child pointer from the field at that offset and a call read directly in raw disassembly, each offset matching `TOWN-232`'s own attribution of the same four children)

### TOWN-253

`R0133`, run on the final stage's `+0x7c` child, writes the exact fields `TOWN-234` left open at Medium confidence ("the per-slot data's own writer was not traced") — the icon-pointer arrays at `this+0x78`/`this+0x8c`/`this+0xa0` and the rect-quad array at `this+0xb4`/`this+0xb8`/`this+0xdc`/`this+0xe0` — resolving that item to High and naming the 34 branch-specific archive paths (of the routine's 37 total) those fields carry. `R0133` loads 3 nodes unconditionally (`RollStatsR.bmp`→`+0x68`, `FullStatsR.bmp`→`+0x6c`, `Inn\RUOver.bmp`→`+0x70`), then branches on `param_2`: fighter (`param_2==0`) writes `fighter\mask.bmp`→`+0x64`, `fighter\column.bmp`→`+0x60`, and 5 weapons (sword/axe/Mace/Pike/Bow) × 3 states — `on`→`+0x78,0x7c,0x80,0x84,0x88`; `shine_off`→`+0x8c,0x90,0x94,0x98,0x9c`; `shine_on`→`+0xa0,0xa4,0xa8,0xac,0xb0` — plus 5 destination-rect quads (X/Y/W/H, stride 8: `+0xb4/+0xb8/+0xdc/+0xe0`, `+0xbc/+0xc0/+0xe4/+0xe8`, `+0xc4/+0xc8/+0xec/+0xf0`, `+0xcc/+0xd0/+0xf4/+0xf8`, `+0xd4/+0xd8/+0xfc/+0x100`); the mage branch (`else`) writes the identical field layout with `mag\mask.bmp`/`mag\column.bmp` and fire/water/air/earth/astral in place of the fighter set (`evidence/strings-R0133.txt`, `evidence/R0133_R0133.c`).

`TOWN-234`'s own pointer arithmetic matches these offsets exactly: its icon-selector `puVar3[-10]/puVar3[-5]/*puVar3`, relative to `this+0xa0` (stride 4), lands on `this+0x78`/`this+0x8c`/`this+0xa0` — the "on"/"shine_off"/"shine_on" state arrays this row's loader writes; its rect-selector `puVar2[-10]/puVar2[-9]/*puVar2/puVar2[1]`, relative to `this+0xdc` (stride 8), lands on `this+0xb4`/`this+0xb8`/`this+0xdc`/`this+0xe0` — the X/Y/W/H fields of the first rect quad this loader writes. The 37-node count reconciles exactly: 3 shared + 17 per-branch (mask + column + 5×3) × 2 branches = 37, matching `TOWN-217` (`evidence/R0133_R0133.c`, `evidence/strings-R0133.txt`, `evidence/disasm-R1870.txt`)

**Confidence.** High (every field offset read directly from the decompiled store and cross-checked against `TOWN-234`'s own published pointer arithmetic, which independently derives the same four offsets from the paint side)

### TOWN-254

`R1995` (default init) and `R0835` (runtime hover/press selector), both run on the final stage's `+0x70` stats-panel child, together write `this+0x174`/`this+0x178` — the icon-pointer fields `TOWN-235`'s own paint loop already reads without tracing their writer — establishing that all ten button-state nodes `TOWN-217` lists as loaded (`{m,p}{disable,nloff,nlon,loff,lon}.bmp`, two buttons × five states) are reachable by `R0833`'s already-published blit. `R1995` loads the ten icons into `this+0x194`..`+0x1b8` (`evidence/strings-R1995.txt`: `plon/ploff/pnlon/pnloff/pdisable/mlon/mloff/mnlon/mnloff/mdisable.bmp`, all under `graphics\interface\chrgen\buttons\`), then copies the `nloff` pair (`this+0x1a4`/`this+0x1b4`) as a default into `this+0x174+8k`/`this+0x178+8k` for k=0..3 (`evidence/R1995_R1995.c`).

`R0835` (5 callers per `evidence/enumrefs-callto-R0835.txt`: three input-handling vtable slots of the same `+0x70` child, plus the one-time dispatcher `R1870` and `R0834`) reads a hover index and a pressed-state value, and for the matching slot overwrites `this+0x174+param_3*8`/`this+0x178+param_3*8` with `plon`/`ploff`/`mlon`/`mloff` in place of the `nloff`/`disable` defaults (`evidence/R0835_R0835.c`). `TOWN-217` names these ten nodes as loaded but does not establish a painter; `TOWN-235` documents `R0833`'s own paint reading `this+0x174+8k`/`this+0x178+8k` without tracing their writer. Not amending either row: both stand as published, this row closes the gap between them (`evidence/R1995_R1995.c`, `evidence/R0835_R0835.c`, `evidence/R2018_R2018.c`, `evidence/strings-R1995.txt`, `evidence/enumrefs-callto-R0835.txt`)

**Confidence.** High (both routines fully decompiled, every store offset read directly, and `R1995`'s destination fields confirmed as the same `+0x70` child `TOWN-235`'s own paint routine belongs to via the shared dispatcher's load of the `+0x70` child pointer, `TOWN-252`)

### TOWN-255

`R0877` (10,911 decompiled bytes) draws no archive node: every one of its calls goes through the text/number primitive `R0571(x,y,value,kind,fontTable,count)`, none through a blit vtable slot, and it references none of `this+0xa0`, `this+0xdc`, `this+0x174`, `this+0x178` or `this+0x64` — closing `TOWN-245`'s own hedge that it "could plausibly be a shared consumer" of the unaccounted nodes. Read in full: the routine gates a detail level (0-7) from several checks on `param_1` (a character-data record, not a chargen-screen widget — its own fields include `+0x18c` flags, `+0xfc`/`+0x100`/`+0x102`/`+0x104`/`+0x106`/`+0x108` stat values, `+0x138`..`+0x14b` skill-level bytes) and issues a bounded sequence of `R0571` calls drawing formatted numbers and short strings at coordinates derived from `param_2[0]`/`param_2[1]` (an (x,y) origin) plus small literal offsets.

No call in the body reaches a bitmap object's own vtable. `EnumRefs callto:R0877` (`evidence/enumrefs-callto-R0877.txt`) finds 5 callers in 5 distinct routines (`L11803`, `L11804`, `L10185`, `L11805`, `L06402`), spanning the chargen area and at least one other screen — a shared, multi-screen character-stat renderer, not a routine local to the final stage's icon fields (`evidence/R0877_R0877.c`, `evidence/enumrefs-callto-R0877.txt`)

**Confidence.** High (the full body read; every call site's own target and argument shape confirmed; the caller enumeration is a whole-image `EnumRefs`)

### TOWN-256

`ScanField disp:64` returns 539 hits across the whole image (`evidence/scan-mask-offset/SCAN_SUMMARY.md`) — `0x64` is a common small displacement reused by unrelated classes throughout the image, so the raw count is not itself informative. Filtered to the `+0x7c` child's own class code range (the addresses `R1871` and its siblings occupy, per `TOWN-234`'s vtable identification), no reader of `+0x64` was found; the field is written (`TOWN-253`) but its consumer, if any, was not located within this scope. Not conclusive: this instrument is bound by Ghidra's own disassembly and reference-manager coverage (`docs/INSTRUMENT.md`) and by the class-range filter's own boundary; a reader reached through a virtual dispatch this experiment did not enumerate, or code outside the filtered range, would not appear (`evidence/scan-mask-offset/SCAN_SUMMARY.md`)

**Confidence.** Medium. Explicit negative, scoped to the `+0x7c` child's own known method-address range and to `ScanField disp:` coverage; does not conclude the field is unread

### TOWN-257

`TOWN-183`'s `+0x10` school background-blit addend is a literal immediate in the instruction stream, present only in the school's own paint routine; it right-aligns the school's 160-wide background bitmap within a 176-wide widget rect the school shares with the shop, and the shop's own paint routine — decompiled here for the first time — carries no equivalent addend because its own background bitmap already exactly fills the same-sized rect. Raw disassembly of `R1905` (`L11806`-`L11749`) reads the 32-bit field at offset 0x8 (the widget's own LEFT, 464 per `TOWN-094`) then forms the sum of it, a second running value and 0x10 — the `0x10` is a literal 8-bit displacement in the address computation itself, not a field load or a computed value.

School's widget rect is `(464,0)-(640,238)` (`TOWN-094`), identical to the shop's own rect; school's background bitmap (`buttonsarea.bmp`, measured 160×238 via a new probe, `tools/bmpdims`, on both preserved roots) is 16px narrower than the 176px-wide rect, and the `+0x10` addend moves its destination X from 464 to 480, landing its right edge flush with the rect's own right edge at 640. `R1027` (the shop's paint routine, `SHOP-SCREEN-035`, not previously decompiled) issues its own background blit as `(local_20+*(this+8), local_1c+*(this+0xc), 0,0,uVar2)` — no additive literal — and the shop's own background bitmap (`ShopMenu.bmp`, measured 176×238) exactly fills its 176px-wide rect starting at its own LEFT (464), reaching the same right edge (640) with no gap.

The tavern's own rect, `(480,0)-(640,238)` (`TOWN-094`/`TOWN-214`), is 160px wide and its own bitmap (`Inn\ButtonsArea.bmp`, measured 160×238) exactly fills it, also with no addend (`TOWN-214`). The `+0x10` therefore reconciles school's narrower bitmap against a rect it shares with the shop, landing its visible left edge at the same absolute x=480 the tavern's differently-sized rect starts at; why school's own rect is 176 wide while its own bitmap is 160 is a content-authoring fact this experiment did not trace to a cause (`evidence/disasm-R1905.txt`, `evidence/R1027_R1027.c`, `evidence/strings-R2122.txt`, `tools/bmpdims/main.go`)

**Confidence.** High (the literal byte read directly from the instruction encoding; the shop's blit call site read directly with no addend present; all four bitmap dimensions measured directly on both preserved roots)

### TOWN-258

The school's two training-button hit-rects `TOWN-182` names at `+0x88`/`+0x98` (16-byte stride) have literal values `(484,71)-(624,117)` and `(484,117)-(624,163)` (LTRB), each 140×46, both lying entirely inside `x:[480,640)` — the span the background blit (`TOWN-183`, `TOWN-257`) covers — so neither button reaches the widget rect's own unpainted left strip, `x:[464,480)`. `R2019` (the training-panel widget's own constructor, called directly from `R1940` with `id=0x3fd, left=464, top=0, right=640, bottom=238`, `evidence/disasm-R1940.txt`) writes these values as literal immediates: the first rect at `this+0x88` receives left 0x1e4 (484), top 0x47 (71), right 0x270 (624) and bottom 0x75 (117), each from a literal loaded into a working value; then the second rect at `this+0x98` receives left 484 (reused), top 117 (reused), right 624 (reused) and bottom 0xa3 (163) (`evidence/disasm-R2019.txt`, `004bb0f`-`004bb64`). Field order (left,top,right,bottom) is fixed by `TOWN-182`'s own established `PtInRect` argument contract; this row supplies the literal contents `TOWN-182` left unstated. Both rects are 140×46, matching the measured button bitmaps (`b1on.bmp`, `button1on.bmp`, both 140×46 via `tools/bmpdims`) (`evidence/disasm-R2019.txt`, `evidence/R2019/R2019_R2019.c`)

**Confidence.** High (every literal value read directly from the instruction stream, both stores confirmed byte-for-byte)

### TOWN-259

Enumeration, all through `R1905`'s own body (`TOWN-183`, re-confirmed this experiment): background (opaque, `vt+0x18`, `buttonsarea.bmp` 160×238) at `origin+(480,0)`, covering `x:[480,640), y:[0,238)` (`TOWN-183`, `TOWN-257`); button 0 icon (opaque, `vt+0x18`, ON=`+0x74` or OFF=`+0x7c` per `TOWN-181`, 140×46) at `origin+(484,71)`, covering `x:[484,624), y:[71,117)` (`TOWN-183`, `TOWN-258`); button 1 icon (opaque, `vt+0x18`, ON=`+0x78`/OFF=`+0x80`, 140×46) at `origin+(484,117)`, covering `x:[484,624), y:[117,163)`. **Amended by [EXP-0205]: the label count and the open Unknown are both closed.** `TOWN-183` undercounted the vt+0x14 draws as one per iteration; the routine issues two per iteration (four call sites in the routine, two of which execute per button per paint), and `TOWN-280` resolves every one of them to X `= origin.x + 554` (both buttons share the same rect center), inside `x:[484,624)`. Negative: none of the four label draws' destinations reach `x<480`, closing this row's own explicit Unknown, in addition to the three previously-resolved blits; `EnumRefs callto:R1905` finds exactly one reference in the whole image, a virtual-dispatch-table entry (`TOWN-261`), so no other routine paints into this rect through a direct call. Whatever the room's own background (`trnhall.bmp`, `TOWN-154`, drawn under every widget) shows through the strip is expected background, not a defect (`evidence/disasm-R1905.txt`, `evidence/enumrefs-callto-R1905.txt`)

**Confidence.** High for all resolved blits' and labels' destinations and the negative among them

**Amended.** [`retracted.md`](retracted.md) holds an entry for this claim (refuted).

### TOWN-260

The shop's own widget-rect background is drawn through the KEYED primitive (`vt+0x38`), unlike school's and tavern's opaque (`vt+0x18`) backgrounds; its 176-wide bitmap exactly fills the 176-wide widget rect with no gap and no `+0x10`-style addend, and a second, conditional KEYED blit draws a press-highlight bitmap at one of `SHOP-SCREEN-035`'s own four rects. `R1027` (`SHOP-SCREEN-035`'s paint routine, decompiled in full for the first time this experiment) issues exactly two blit calls, both through `(**(**(this+0x74)+0x38))(...)` — the keyed primitive (`TOWN-150`; source pixel value `0x0000` is the transparency key, not written). First: the background (`this+0x74`, `ShopMenu.bmp` per `evidence/strings-R2122.txt`, measured 176×238 via `tools/bmpdims`) at `(local_20+*(this+8), local_1c+*(this+0xc), 0,0,uVar2)` — no additive literal; `176 = 640-464`, the widget's own rect width (`TOWN-094`), so the bitmap exactly fills the rect from its own LEFT (464) to 640, unlike school's 16px-short bitmap in the identically-sized rect (`TOWN-257`).

Second: a conditional highlight, one of four bitmaps at `this+0x64`/`0x68`/`0x6c`/`0x70` (`ShopButton1..4.bmp`, sizes 120×52/140×46/140×46/120×52, each matching its own rect's own size exactly, `evidence/strings-R2122.txt`, `tools/bmpdims`), blitted at that rect's own left/top (`this+0x78+idx*0x10`/`this+0x7c+idx*0x10`, the same rect array `SHOP-SCREEN-035` already publishes), gated on `-1<this+0xf8 && -1<this+0xfc && this+0xf8==this+0xfc` (pressed index equals hovered index). `SHOP-SCREEN-035`'s own eight numeral draws are unaffected (unconditional, through the shared text primitive, no blit). This clarifies rather than contradicts `SHOP-SCREEN-035`'s own Medium-confidence pairing of the four bitmaps to the four rects: no unconditional per-frame button-face blit exists in this routine; the pairing is a press-only highlight drawn at the paired rect's own position (`evidence/R1027_R1027.c`, `evidence/strings-R2122.txt`, `tools/bmpdims/main.go`)

**Confidence.** High for the two blit calls, their primitive, and the background's dimension match (read directly at instruction level; dimensions measured on both preserved roots); the highlight/rect pairing inherits `SHOP-SCREEN-035`'s own Medium (store-order basis, corroborated but not independently established by the size match alone, since two of the four rect pairs already share sizes)

### TOWN-261

The object `TOWN-181` names only as "`[school_view+0x78]`" and the object `TOWN-182`/`TOWN-183`'s own `this` are the same object — the school room's own `+0x78` child — and that object is also the sole one whose vtable (`R2020`) resolves `R1905` (`TOWN-183`'s paint routine) at slot `+0x2c`, naming the view that owns the `L11671`..`L11672` loader `TOWN-154`/`TOWN-181` leave attributed only by offset. `EnumRefs callto:R1905` returns exactly one hit in the whole image: a virtual-dispatch-table entry at `L11807` — no direct `CALL` reaches it anywhere in `.text` (`evidence/enumrefs-callto-R1905.txt`). `L11807` is vtable `R2020`'s own `+0x2c` slot (the same paint-slot offset `TOWN-214` documents for the tavern's own vtable, `L08008+0x2c`); `R2020` is the vtable `R2019` installs unconditionally (`*param_1=&PTR_R2020`) on the object it constructs, and `R2019` is called directly from `R1940` (the school widget's own constructor, `TOWN-094`, id `0x3fd`, rect `(464,0)-(640,238)`).

So `R1905`'s own `this` is that widget, reached only through this one virtual slot. Separately, raw disassembly of `R0706` ("school enter," `TOWN-181`'s own citation) shows three consecutive calls — to `R1973` (the button-bitmap loader `TOWN-181` reads), `L11808`, and `R1977` (the `+0xac`/`+0xb0` initializer `TOWN-182` reads) — each immediately preceded by a load of the dword at offset 0x78 of the screen object as `this` (`L11809`-`L11810`, `evidence/disasm-R0706.txt`), where the screen object is `R0706`'s own `this` (the school room object). Since `R1977` is `TOWN-182`'s own cited routine and it is called here with `this=[room+0x78]`, the object `TOWN-181` calls "`[school_view+0x78]`" and `TOWN-182`/`TOWN-183`'s own `this` are the same object (`evidence/enumrefs-callto-R1905.txt`, `evidence/enumrefs-callto-R1940.txt`, `evidence/disasm-R0706.txt`, `evidence/disasm-R1940.txt`, `evidence/R2019/R2019_R2019.c`)

**Confidence.** High (the vtable-slot dispatch resolved by whole-image `EnumRefs`, the vtable installation read directly in the decompiled constructor, and the shared-`this` proof read directly from three consecutive call sites in raw disassembly)

### TOWN-264

`tools/intfcensus` opens the `interface/` subtree, extracts every printable-ASCII run of length ≥4 from `rom.exe`'s raw bytes, keeps runs containing the substring `"interface"`, trims each run to start at its own first `"interface"` occurrence (a run otherwise still carries the archive-name prefix, e.g. `graphics\interface\...`, which defeats a path-anchored match), and classifies each archive entry as `literal` (the trimmed run contains the entry's own backslash path as a substring) or `fuzzy` (the trimmed run, with each printf conversion `%\.?[0-9]*[sdux]` rewritten to `[^\/]*`, matches the entry's path as a regex) or `unreferenced`. EN root: 586 `interface/` entries, 232 literal + 277 fuzzy = 509 referenced, 77 unreferenced.

RU root: 584 entries, the same 232 + 277 = 509 referenced, 75 unreferenced. Calibration: restricted to `interface/training/`, the tool reports 63 entries / 38 literal / 25 fuzzy / 0 unreferenced — an exact match to `TOWN-154`'s independently published count over that same subtree. **Instrument bound:** a reference built from two string fragments concatenated at runtime (a stored prefix joined to a separately stored suffix) would not appear as one printable run and is not found by this scan; otherwise the scan is exhaustive over ASCII runs ≥4 bytes in the executable's own bytes and does not depend on disassembly coverage (`tools/intfcensus/main.go`, `evidence/en/intfcensus.csv`, `evidence/en/summary.txt`, `evidence/ru/intfcensus.csv`, `evidence/ru/summary.txt`)

**Confidence.** High (the counts are read from an exhaustive raw-byte scan to its own stated bound, and the scan reproduces an independently published subtree count exactly); the two-fragment-concatenation blind spot is named, not closed
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration. The agreement between the per-root
data stands as corroboration: the `interface/` archive entries are separate files in each root.

### TOWN-265

The 77 EN-root / 75 RU-root unreferenced `interface/` entries (`TOWN-264`) fall into seven groups, five of them a loose per-frame `.bmp` file beside a packed sprite container of a related name the code does reference by path, and the fragment names most likely to hide a concatenated reference are independently absent from the executable at any string length. Groups, EN root, 77 total: chrgen cube loose frames beside a referenced `chrgen/cube/sprites.16a`, 32 (`a_cb0000`..`0015.bmp`, `cb0000`..`0015.bmp`); chrgen precreate/blind loose frames beside a referenced `chrgen/precreate/blind/sprites.16a`, 30 (`a_bl0000`..`0014.bmp`, `bl0000`..`0014.bmp`); a chrgen-root duplicate of a file referenced under `chrgen/loader/` (`centerarea.bmp`) and an unpaired `rollstatsl.bmp` beside a referenced `rollstatsr.bmp`, 2; the tavern's (`interface/inn/`) alternate two-button naming — `lbuttonon/off.bmp`, `rbuttonon/off.bmp` — beside the three-button `button1..3{on,off}.bmp` scheme the code does reference, 4; a single stray `tav_09.bmp`, with no `tav_01`..`tav_08` present in the archive at all, 1; seven shop-related loose `.bmp` files each beside a referenced `.256`/`.16a` container of a related name (`myitem.bmp` and `shopanim/myitem.bmp` beside a referenced `myitem.256`; `shopanim/shopitem.bmp` and `shopitem.bmp` beside a referenced `shopitem.256`; `shopframe.bmp` beside a referenced `ShopFrame.256`; `shopie.bmp` beside a referenced `townbirds/shopie/sprites.16a`; `t_border.bmp` beside a referenced `t_border.256`), 7; `batlmen7.bmp`, EN-only, no `batlmen1`..`6` present in the archive, 1 (32+30+2+4+1+7+1 = 77).

Cross-check: an unfiltered raw-ASCII-run scan (length ≥3, no `"interface"` filter) over the whole executable finds zero occurrences, at any string length, of `lbutton`, `rbutton`, `tav_0`, `batlmen`, `a_cb`, `a_bl` or `rollstatsl` — ruling out a two-fragment concatenation for these seven names, the specific blind spot `TOWN-264`'s own instrument names. For `myitem`, `shopie`, `shopitem`, `shopframe`, `t_border`, the same unfiltered scan finds exactly one occurrence each, and each one is the already-identified referenced sibling's own path (e.g. `graphics\interface\shopframe.256` for `shopframe`), not a second reference to the unreferenced `.bmp` name.

The two EN-only entries, `batlmen7.bmp` and `shopie.bmp`, account for the entire 586-vs-584 EN/RU entry-count difference: both are absent from the RU archive outright, not merely unreferenced on one root (`evidence/en/summary.txt`, `evidence/ru/summary.txt`, `evidence/en/intfcensus.csv`)

**Confidence.** High for the grouping, the sibling identification and the fragment-absence cross-check (each read directly from the census's own classification or from an independent unfiltered scan); Unknown for why the loose frames and alternate schemes remain in the archive — retired art, an abandoned two-button layout, and a build leftover are equally consistent, and no instrument in this repository distinguishes intent

### TOWN-280

The school training widget's origin is `(0,0)` at the game's default 640x480 resolution, which resolves both of `TOWN-183`'s undecoded `vt+0x14` label draws — now known to be four call sites, two per button, not two — to an absolute X of 554 for every one, inside `x:[484,624)` and clear of the widget's own unpainted `x:[464,480)` strip; and the widget's own rect literals `(464,0,640,238)` are a hardcoded immediate independent of the background bitmap's own dimensions, so `TOWN-257`'s 16px gap is two unrelated authored numbers, not a computed relationship. Origin trace: `R1905` (`TOWN-183`) reads origin from `[widget+0x5c]->+8/+0xc`. `R2019`, the widget's own constructor, stores `param_1[0x17] = param_7` (`widget+0x5c = param_7`, `evidence/R2019_R2019.c`), called as `R2019(0x3fd,0x1d0,0,0x280,0xee,param_1)` from `R1940` (`evidence/R1940_R1940.c`) — six literal/forwarded arguments, the last being `R1940`'s own `this`. Raw disassembly of `R2021` (the school room's own constructor) shows the incoming `this` kept at entry and passed again as `this` immediately before the call to `R1940` (`L11811`/`L11812`, `evidence/disasm-R2021-R1940.txt`), so `R1940`'s own `this`, and therefore `widget+0x5c`, is the school room object. The room's own `+8`/`+0xc` are set by the same shared base-constructor chain `TOWN-184` already traces for a different concrete class: `R2021` → `R0759` → `R0773` → `R0388(left,top,right,bottom)` → `R1980` → the genuine `SetRect` import, whose `LPRECT` is `this+8` (`this+8` formed at `L11813`, `experiments/EXP-0200-paint-destinations/evidence/disasm-R0388-setrect-forwarder.txt`, re-read this experiment). The room is constructed as `R2021(0x3fc,[L09808],[L09809],[L09810],[L09811])` at `campaign+0x104` (`experiments/EXP-0202-chargen-page-geometry/evidence/R0315_R0315.c:153`), the same four globals `TOWN-222` already reads as `(0,0,640,480)` at the shipped default resolution with a single writer, `R0341`. So `room+8 = room+0xc = 0` under the default resolution, and `widget+0x5c->+8/+0xc = (0,0)`. **Labels:** raw disassembly of `R1905` shows two calls into the shared `vt+0x14` primitive per branch — `L07784`/`L07785` in the ON branch (fall-through at `L11814`), `L07786`/`L07787` in the OFF branch (`L11752`) — not one as `TOWN-183` originally stated (amended above). Both calls in a branch share the same X argument, `(RIGHT-LEFT)/2 + origin.x + LEFT`; with `origin.x=0` and both buttons' rects `LEFT=484/RIGHT=624` (`TOWN-258`), X `= 554` for all four. Y differs: the first call is `TOP + floor(H/4) + iVar4 + {8 OFF, 9 ON}` where `H=BOTTOM-TOP=46` gives `floor(46/4)=11`, and `iVar4` is a metric read from one of two globals (`[L09922]`/`[L09928]`, selected by the same `+0xb0==i` test) not traced further in this experiment; the second call is `TOP + floor(3H/4) + iVar2 + {0 OFF, 1 ON}` where `floor(3*46/4)=34` and `iVar2 = origin.y + widget-local-TOP = 0`, giving a fully numeric Y of `TOP+34` (OFF) / `TOP+35` (ON) — button 0 (`TOP=71`): 105/106; button 1 (`TOP=117`): 151/152. The first call's string pointer is read from a 2-entry array at `widget+0x64`, indexed by button; the second reads a per-button scratch dword at `room+0x360`/`room+0x364`, written immediately before by `R0807`. Neither array's contents (the caption text, or what `R0807` formats) were read — out of this claim's scope, which is destination only. **Q2's second half**: nothing draws into the widget's own `x:[464,480)` strip — `TOWN-259`'s three resolved blits and all four labels here stay inside `x:[480,640)`/`x:[484,624)`; the room's own background bitmap `trnhall.bmp` (480x480 at `(0,0)`, `TOWN-154`), drawn under every widget, shows through unobstructed (`evidence/R1940_R1940.c`, `evidence/R2019_R2019.c`, `evidence/disasm-R2021-R1940.txt`, `evidence/R1905_R1905.c`, `experiments/EXP-0203-node-painters/evidence/disasm-R1905.txt`)

**Confidence.** High for the origin trace, the X destinations of all four labels and the negative for `x:[464,480)` (every step read at instruction level or inherited from an already-published High-confidence claim tracing the identical shared routines); High for the second call's Y (fully numeric); explicitly Unknown for `iVar4`'s value (bounds the first call's Y to a formula, not a number) and for both arrays' contents

### TOWN-281

The school's 640x480 page in draw order, assembled from previously published claims plus this experiment's own resolution of the label draws (`TOWN-280`): a 480-wide main content area painted by the room's own routine, and a 160-wide right column shared with the shop and tavern surfaces. Room paint `R1487` (`TOWN-061`, amended): (1) `trnhall.bmp`, 480x480, opaque, at `(0,0)`, covering `x:[0,480), y:[0,480)` (`TOWN-154`); (2) mage-side lower-area element `+0xf4` or a fill — archive entry and destination not published by any claim this experiment cites, a gap; (3) fighter-side lower-area element `+0x1a8` or a fill — same gap; (4) one of 16 `column/rt0000..0015.bmp` frames (148x208), opaque, at `(168,176)`, selected by `+0x31c` (`TOWN-146`, `TOWN-154`); (5) the active class panel's 5 skill icons, keyed-blit (source value 0 transparent) from one of 15 per-class bitmaps chosen by a 3-state (`on`/`shine`/`shine_on`) per-slot array — mage panel rect `(188,188)-(288,308)` with icon rects `+0x10c`..`+0x14c`, fighter panel rect `(192,192)-(284,312)` with icon rects `+0x1c0`..`+0x200`, selected by `+0x31c==0xf`/`==0` (`TOWN-148`, `TOWN-149`, `TOWN-150`, `TOWN-151`, `TOWN-152` for the two field-vs-file index orders); (6) one of 9 `diamond/on0000..0008.bmp` frames (80x76), opaque, at `(200,60)` (`TOWN-154`); (7) scrollbar-thumb element `+0x68`, gated — archive entry and destination not published, a gap; (8) scrollbar-track element `+0x74`, gated — same gap. Right column, borrowed from the same 160-pixel strip the shop and tavern share: upper widget (id `0x3fd`, rect `(464,0)-(640,238)`, `TOWN-094`) blits `buttonsarea.bmp` (160x238) opaque at `(480,0)` (`TOWN-183`, `TOWN-257`), then for each of its two buttons — rect `(484,71)-(624,117)` and `(484,117)-(624,163)` (`TOWN-258`) — one ON or OFF icon (140x46, opaque) at the rect's own top-left, then two `vt+0x14` label draws, both at absolute X `554`, Y `TOP+11+iVar4+{8,9}` (caption, `iVar4` untraced) and Y `TOP+34`/`+35` = `105`/`106` (button 0) or `151`/`152` (button 1) (number, fully resolved) (`TOWN-183` amended, `TOWN-280`); lower panel (id 7, rect `(480,238)-(640,480)`, the same `campaign+0xe0` object `SHOP-FIGURE-041` names, transferred into the school by the identical `RemoveChild`/`OffsetRect`/`SetRect`/`AddChild` sequence `TOWN-088` reads at instruction level, byte-for-byte the tavern's own `TOWN-087`) — the composed equipped-hero figure through the same compositor `SHOP-FIGURE-041`/`DLG-FIGURE-020` name for the shop and dialogue surfaces; this experiment did not independently re-verify that the school's own borrow paints the identical content, only that it is the identical object reached by the identical transfer mechanism

**Confidence.** High for every step this row cites from an already-published claim (unchanged from that claim's own grade) and for the label destinations (`TOWN-280`); explicitly no destination for the two lower-area fills or the two scrollbar elements — not found in any claim this experiment cites, and not independently traced here

### TOWN-282

Siblings, all constructed by `R1939` (`TOWN-064`); the paint order **among** the three siblings is not independently established by this experiment or any claim it cites — `TOWN-063`/`TOWN-266` name the shared child-walker `R1696` but leave its own caller Unknown, so the order given here is construction order, not a confirmed paint order. **Left panel** (id `0x44d`, rect `(0,0)-(160,480)`): the loader `R1930` (`TOWN-009`) fills `+0x78`=`LeftStats.bmp` (160x238) and `+0x7c`=`LeftPicture.bmp` (160x242, both measured this experiment, `tools/bmpdims`, identical on the install and both preserved roots). The paint `R0892`, read for the first time this experiment, is substantially richer than `TOWN-064`'s "portrait/stats" characterization: both bitmaps are blitted unconditionally near the top of the routine through `vt+0x18` (opaque), at coordinates this experiment did not resolve to a pixel destination (an origin-fetch helper `R1224` and a hidden-struct return convention were not traced); when a roster entry is hovered, the routine additionally branches on bits of a per-entry flags byte into at least three further paths — a default two-line text draw, a `graphics\infowindow_%s.bmp`-derived custom picture, and a `GetTempPathA`-based `allods_d_%d`-named external bitmap load outside `graphics.res` entirely — none of which is decoded past its own branch condition; a fourth call blits a fixed `0xb`x`0xa0`x`0xf0` region through `vt+0x38` (keyed) at `(dy+0xf0)` at the routine's own end.

This is a scoped negative on top of `TOWN-009`/`TOWN-064`, not a correction: neither prior row asserted these paths do not exist, both are silent on them. **Button panel** (id `0x44e`, rect `(480,0)-(640,238)`, already fully enumerated by `TOWN-214`): `ButtonsArea.bmp` (160x238) opaque at `(480,0)`, then the middle button at `(484,91)-(624,137)` and the other two at `(484,44)-(624,90)`/`(484,138)-(624,184)`, each an ON/OFF icon plus one centred digit draw through the shared text primitive — `TOWN-214` names no caption text on any of the three, which this experiment did not check against the shop's own parallel finding (`SHOP-050`, this same experiment's ledger-integrity fix) for whether the same digit-plus-caption pattern applies here; left open.

**Center/roster panel** (id `0x450`, rect `(160,0)-(480,480)`): the shared 6-column cell grid at `slot(i)*0x10` byte offset, `slot(i) = (i mod 6) + (floor(i/3)+1)*6` (`TOWN-065`, superseding `TOWN-010`'s formula clause), talk-only NPCs then mercenaries continuing the same index (`TOWN-010`'s structure clause, unaffected); each cell's portrait is one of 17 shipped `interface/inn/Unit%d/sprites.16a` sheets, indices 1-15 and 29-30 (`TOWN-012`); reached only through virtual dispatch (`TOWN-011`). No claim this experiment cites gives the per-cell destination rectangle in absolute screen coordinates, only the byte-offset formula — a gap. **Right column, lower panel** (id 7, rect `(480,238)-(640,480)`, the same object `SHOP-FIGURE-041` names): transferred into the tavern by the identical `RemoveChild`/`OffsetRect`/`SetRect`/`AddChild` sequence `TOWN-087` reads at instruction level — the tavern is `TOWN-087`'s own subject, so this is not new to this experiment, cited here only to complete the page (`evidence/R1804_R1804.c`, `evidence/R0892_R0892.c`)

**Confidence.** High for every step cited from an already-published claim (unchanged from that claim's own grade) and for the two measured bitmap dimensions; explicitly no destination for the left panel's own two unconditional blits or for the roster grid's absolute per-cell rectangle; the left panel's hover-state branches are read but not decoded, and the button panel's own caption question is stated but not checked

**Amended.** roster clause, see `TOWN-467`, `TOWN-468`. [`retracted.md`](retracted.md) holds an entry for this claim (partially retracted).

### TOWN-283

The character generator's two 640x480 pages, assembled from previously published claims: the precreate stage (`R1474`, 26 archive nodes) and the final/detailed stage (top-level screen `R1314`, which paints nothing of its own and delegates to at least 4 children covering at least 55 archive nodes) -- three of the final stage's four children have a known destination rect, the fourth does not, and the region those three leave uncovered is consistent with an unread fourth child, not confirmed. **Precreate** (loader `R1672`, `TOWN-216`; paint `R1474`, vtable `L07691+0x2c`, `TOWN-223`): this stage's own top-level origin (`this+8`/`this+0xc`) was not traced to its own construction call by `TOWN-223` or by any claim this experiment cites, unlike the final stage below.

Ordered blits, in `R1474`'s own order: (1) background `+0x68`=`MainArea.bmp`, `vt+0x18` at `(this+8,this+0xc)`, unconditional; (2)-(3) two hover highlights (`+0x184` from an `Amulet.bmp` copy, `+0x188` from a `ButtonOk.bmp` copy) via `vt+0x38`, each gated on its own state field; (4) Levels loop, 3 iterations, `level{0,1,2}{on,l,lon}.bmp` via `vt+0x38`, dest read through a pointer at `+0x124[slot]` (`TOWN-236`); (5) Heroes loop, 4 iterations, `{mf,ff,fm,mm}{on,l,lon}.bmp` via `vt+0x38`, dest through a pointer at `+0x110[slot]` (`TOWN-236`); (6) steps 2-3 repeated, unexplained; (7) a tip-popup call, gated; (8) `+0x70`=`Blind\sprites.16a`, animated, `vt+0x18`, gated; (9) one text/digit draw through the shared `[L06186]+0x14` primitive.

Two gaps carried forward, not re-attempted this experiment: the Levels/Heroes loops' own rectangle table was searched for by two instruments keyed on the wrong vtable base (`TOWN-236`, `TOWN-244`) and has not been read, so its literal values remain Unknown; `Mask.bmp` (`+0x6c`, loaded through a different constructor than the other 25 nodes) is read by neither the paint routine nor its hover-blink helper (`TOWN-223`'s own scoped negative), and its consumer, if any, was not located by any claim this experiment cites. **Final/detailed stage** (top-level screen, vtable `R1312`, constructed `R1314(0x456,[L09808],[L09809],[L09810],[L09811])` from the campaign screen constructor, `TOWN-242`): under the default resolution the same four globals give this screen's own construction rect as `(0,0)-(640,480)`, confirmed rather than assumed (`TOWN-242`).

The top-level screen's own paint routine issues no blit of its own; every pixel is drawn through the generic child-paint walk (`TOWN-245`, matching the shared no-own-content epilogue `TOWN-208` documents elsewhere). `TOWN-245` names four children's own paint routines as `TOWN-224`/`TOWN-233`/`TOWN-234`/`TOWN-235`; this experiment fetched `TOWN-233` and `TOWN-234` only, covering three of the four (`TOWN-233` covers two: `+0x74` and `+0x78`; `TOWN-234` covers one: `+0x7c`). Their own destination rects: `+0x74` (`R1992`) `(0,238)-(160,242)`; `+0x78` (`R1993`) `(480,0)-(640,238)`, gated, addend confirmed `(0,0)` (`TOWN-242`); `+0x7c` (`R1871`) `(160,0)-(480,480)`.

**Gap**: the fourth child's own identity, rect and content, `TOWN-224` or `TOWN-235`, not distinguished here, was not fetched by this experiment. The three known rects leave `x:[0,160)` above `y=238` and below `y=242`, and `x:[480,640)` below `y=238`, uncovered; this is consistent with the unfetched fourth child painting that remainder, but not confirmed by anything this experiment read. Per-child content, from the two fetched claims: `+0x74` one background blit (`this+0x60`) plus an undecompiled call (`R0877`); `+0x78` one gated background blit (`this+0x8c`) plus a 3-iteration icon loop, each iteration drawing one icon and one centred text label through the shared primitive (this corrects `TOWN-217`'s scope, not its finding: `TOWN-217` read no string literal in this routine's own bytes and did not read further; the label's text comes from a runtime table, not a literal); `+0x7c` one background blit, four 16px border-column blits at the panel's own left (`x=160`) and right (`x=464`) edges, a 5-slot 3-state icon loop, and the previously-undecompiled `R2004` (resolved by `TOWN-245`: a 500ms-gated animation tick that reuses the same child's own icon-slot arrays, not a new node). Node count: at least 55 archive nodes reach this stage (`TOWN-245`, correcting `TOWN-217`'s 48 by a 7-node loader `TOWN-217` did not enumerate); the painter for 54 of the 55 (all but `leftup.bmp`) is not established by any claim this experiment cites, including the icon loops' own text-source table and the fourth child's own content

**Confidence.** High for every step cited from an already-published High-confidence claim, unchanged from that claim's own grade; explicitly Unknown for precreate's own top-level origin, the Levels/Heroes loops' literal rects, `Mask.bmp`'s consumer, the fourth final-stage child's own identity and rect, and the painter of 54 of the final stage's 55 archive nodes

**Amended.** The precreate vtable base and slot (`L07691`, `+0x2c`) and the table-search sentence, corrected in the text above; the ordered blit list, node counts, final-stage rectangles and Unknowns stand. [`retracted.md`](retracted.md) holds the entry.

### TOWN-284

165 `graphics.res` `interface/` entries are classed `literal` (`TOWN-264`/`EXP-0204`) and sit directly under `interface/` or under `interface/inn/`/`interface/chrgen/`: 67 direct, 17 under `inn/`, 81 under `chrgen/`. Of the 81 `chrgen/` entries, 59 have a located reference site (26 resolved to a destination, 30 to a surface and mechanism only, 3 to a scoped negative consumer); of the 17 `inn/` entries, 9 have one (7 resolved, 2 to a surface with an unresolved destination); the 67 direct entries were not traced in this experiment beyond 4 already-published destinations. Population re-derived independently this experiment from `EXP-0204`'s own `evidence/en/intfcensus.csv` (586 rows): filtered to `status==literal`, then to `path` starting with `interface/` and either containing no further `/` (67 rows) or starting with `interface/inn/` (17 rows) or `interface/chrgen/` (81 rows, any depth) -- 165 total, reproduced by a fresh script run against the same CSV rather than carried over from any prior count of this same population.

**`chrgen/` (81).** `interface/chrgen/precreate/*` is exactly the 26 nodes `TOWN-216` enumerates and `TOWN-283` locates: `mainarea.bmp`, `amulet.bmp`, `buttonok.bmp`, `blind/sprites.16a` each to one named blit site; `heroes/{mf,ff,fm,mm}{on,l,lon}.bmp` (12) and `levels/level{0,1,2}{on,l,lon}.bmp` (9) each to a `vt+0x38` call gated on a per-slot state field, destination known only as a table read through an unresolved pointer (`TOWN-236`, `TOWN-244`); `mask.bmp` (1) to a scoped negative -- referenced by the loader, read by neither the paint routine nor its hover-blink helper (`TOWN-223`). `fighter/{axe,bow,mace,pike,sword}` and `mag/{air,astral,earth,fire,water}`, each `{on,shine_off,shine_on}.bmp` (30 total): surface and mechanism located, the final stage's `+0x7c` child's 5-slot 3-state icon loop (`TOWN-234`), destination not resolved past "a per-slot stored rect" -- the rect's own literal values and writer were not traced by `TOWN-234` or by this experiment.

`fighter/mask.bmp`, `mag/mask.bmp` (2): a scoped negative already published, `TOWN-256` -- written to the same child's `+0x64`, no reader found within that child's own known methods. `chrgen/fullstatsr.bmp` (1): loaded by `R2013` into the top-level screen's own `+0x60` (`TOWN-245`); `TOWN-245`'s own search of all four children's paint bodies, the top-level's own paint dispatch and its animation timer for any reference to `+0x60` found none -- a scoped negative, and `TOWN-245` itself flags it as not conclusive, since two further routines that could plausibly consume it were decompiled but not read. `chrgen/loader/leftup.bmp` (1): named by `TOWN-245` as resolved by `TOWN-217`, a claim this experiment did not fetch; cited, not independently verified here.

**Not reached in this experiment (21 of 81):** `fighter/column.bmp`, `mag/column.bmp` (2); `buttons/{m,p}{disable,loff,lon,nloff,nlon}.bmp` and `buttonsarea.bmp` (11); `cube/sprites.16a` (1); `fullstatsl.bmp` (1, no loader for this specific file located, unlike its `r` sibling above); `loader/{centerarea.bmp,down/shine.bmp,down/shine_on.bmp,up/shine.bmp,up/shine_on.bmp}` (5, plus `leftup.bmp` above, suggesting a third chargen sub-stage, a "loader" screen, distinct from precreate and final -- not otherwise touched by this experiment or `TOWN-283`); `rollstatsr.bmp` (1). No claim this experiment cites gives a reference site for any of these 21.

**`inn/` (17).** `buttonsarea.bmp` and `button{1,2,3}{on,off}.bmp` (7): resolved, the tavern's button panel (`TOWN-214`, restated with rects in `TOWN-282`) -- `buttonsarea.bmp` opaque at `(480,0)`, each button's on/off icon at its own rect, `(484,44)-(624,90)`, `(484,91)-(624,137)`, `(484,138)-(624,184)`. `leftstats.bmp`, `leftpicture.bmp` (2): loaded (`TOWN-009`) and blitted unconditionally, confirmed by this experiment's own new decompile of the left panel's paint routine (`R0892`, `TOWN-282`), but the blit's own pixel destination was not resolved -- an origin-fetch helper (`R1224`) and a hidden-struct return convention were not traced.

**Not reached in this experiment (8 of 17):** `herofighter/sprites.16a`, `heromage/sprites.16a` -- no claim this experiment cites names these two files; `TOWN-012`'s 17 roster-portrait sheets are named `Unit%d/sprites.16a`, a different path shape, so they are not assumed to be the same nodes. `ldover.bmp`, `luover.bmp`, `ruover.bmp`, `manback.bmp`, `manbacktalk.bmp`, `centerarea.bmp` (6): no reference site located by any claim this experiment cites. **Direct under `interface/` (67).** `shopbutton1..4.bmp` (4): resolved, the shop's own button panel (`SHOP-SCREEN-035`) -- rects `(494,15,614,67)`, `(483,67,623,113)`, `(483,114,623,160)`, `(494,160,614,212)`, pairing at Medium confidence (`SHOP-SCREEN-035`'s own grade, unchanged).

**Not reached in this experiment (63 of 67):** the remaining 63 direct entries were not traced to a reference site in this experiment. Roughly a third by name (`shoparrow1..4.bmp`, `shopframe.256`, `shopinv.bmp`, `shopitem.256`, `shopmenu.bmp`, `shoptable.bmp`, 9 entries) plausibly belong to the shop surface the brief's own model claims cover, but this was not checked against those claims' own text in this experiment. The remainder (`ar1..4.bmp`, `backinv*.bmp`, `ball.bmp`, `book{closed,opened}.bmp`, `commandbar{l,r}.bmp`, `commanddnr.bmp`, `commandempr.bmp`, `crystal{l,r}.bmp`, `diskette.bmp`, `extra{800,1024}{l,r}.bmp`, `heads{l,r}.bmp`, `humanback{l,r}.bmp`, `humanmode.bmp`, `inv1024{l,r}.bmp`, `invarrow1..4.bmp`, `invframe.bmp`, `lm.256`, `minimapdata.bmp`, `myitem.256`, `radiob.256`, `scrlbars.256`, `server.bmp`, `spb{800,1024}{l,r}.bmp`, `spellback.bmp`, `spellbook.bmp`, `t_back.bmp`, `t_border.256`, `testiva.bmp`, `textback{l,r}.bmp`, `textmode.bmp`) names, by content, an inventory screen, a command bar and a minimap -- surfaces outside the school, the tavern and the character generator this experiment traces, and outside the shop this experiment's own brief names as a model. None of these 54 were searched for a reference site in this experiment.

**Confidence.** High for every count and every destination cited from an already-published claim, unchanged from that claim's own grade; High for the 165-entry population figure and its three-way split (a fresh filter over `EXP-0204`'s own CSV, reproduced); explicitly no reference site located, in this experiment, for the 21 `chrgen/` entries, the 8 `inn/` entries and the 63 direct entries named above as not reached

### TOWN-266

The same three shared routines `TOWN-063` already reads for the four top-level room views are also what the tavern's own left-panel child's vtable resolves to at `+0x34`, `+0x38` and `+0x40` — so a child-specific override at these three slots is not `TOWN-063`'s open invoker either. The left-panel child's paint routine, `R1804` (`TOWN-064`), is reached through exactly one virtual slot in the whole image (`EnumRefs refto:R1804`, one hit, a vtable entry at `L11815`; no direct `CALL` anywhere in `.text`), which is vtable base `L11816`'s own `+0x2c`. `FindVtables` over `L11817`..`L11818` reads that same vtable's `+0x34`, `+0x38` and `+0x40` slots as `R0366`, `R0354`, `R1696` — the identical three addresses `TOWN-063` already names and decompiles as the chain "shared byte-for-byte across tavern/shop/town/school's own vtables," there checked only on the four top-level views.

This experiment's own reading of the three bodies matches `TOWN-063`'s: `R0366` (`+0x34`) gates on its own `+0x20` slot returning zero, then calls its own `+0x2c` and its own `+0x38`; `R0354` (`+0x38`) is three opaque calls (`R1865`, `R1224`, `R0351`) with no visible reference to a child list or to any vtable slot; `R1696` (`+0x40`) walks the child list (`R1868`/`R1869`) and calls each child's own `+0x40`, never `+0x2c` or `+0x34`. What this experiment adds is the child-level reading: the sharing `TOWN-063` found among the four top-level views extends one level deeper, to at least this one child, which carries no override of its own at these three slots.

`TOWN-063`'s open question — what invokes the tavern's three children's own paint overrides — is not resolved by this reading; it is narrowed by ruling out a child-specific override of these three slots as the mechanism, for this one child (`evidence/vt-tavern-children.txt`, `evidence/refs-tavern-children.txt`, `evidence/refs-tavern-children2.txt`, `evidence/R0366_R0366.c`, `evidence/R0354_R0354.c`, `evidence/R1696_R1696.c`)

**Confidence.** High for the vtable-slot identity (read directly by `FindVtables`/`EnumRefs`) and for the three routines' own bodies (independently decompiled this experiment, matching `TOWN-063`'s own published reading). The invoker itself remains Unknown, consistent with `TOWN-063`; the other two tavern children's own vtables were not checked

## Character generator final stage and tip popup rectangles

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-312 | The final/detailed stage's `+0x7c` child's four border-strip source fields (`this+0x68`/`+0x6c`/`+0x70`/`+0x74`, the blit sources `TOWN-234` reads) are populated by `R0133`… | High / Unknown | ● active | [EXP-0207](../experiments/EXP-0207-chargen-final-stage/) |
| TOWN-313 | The final/detailed stage's fourth top-level child (`TOWN-283`'s open item, "`TOWN-224` or `TOWN-235`, not distinguished") is the `+0x70` object — both already name the same routine, `R0833`, not two candidates… | High / Unknown | ● active | [EXP-0207](../experiments/EXP-0207-chargen-final-stage/) |
| TOWN-314 | The character generator's tip popup (`TOWN-187`) is attached to the `+0x7c` child at construction, and the popup's own paint routine resolves its absolute rect through the same ancestor accumulator `TOWN-232` establishes… | High / Medium / Unknown | ● active | [EXP-0207](../experiments/EXP-0207-chargen-final-stage/) |
| TOWN-328 | Town's and school's own tip popups resolve to exactly their published literal rects: `TOWN-314`'s ancestor walk applies zero further offset, because each popup's parent is the room view itself and that view's own ancestor chain is empty. | High / Medium | ● active | [EXP-0208](../experiments/EXP-0208-room-tip-screen-rects/) |
| TOWN-329 | Tavern's own tip popup does not resolve to its published literal rect: a two-level ancestor walk applies a `+160,0` offset, giving absolute `(160,0)-(472,200)`, not `(0,0)-(312,200)`. | High / Medium | ● active | [EXP-0208](../experiments/EXP-0208-room-tip-screen-rects/) |
| TOWN-330 | `R0701`'s unmatched `R1261` call site (`L11819`, `TOWN-314`) belongs to the campaign screen's own message handler, and its popup's absolute rect equals its published literal `(0xa,0x14,0x172,0xbc)`: the parent… | High / Medium | ● active | [EXP-0208](../experiments/EXP-0208-room-tip-screen-rects/) |
| TOWN-331 | `R1475`'s unmatched `R1261` call site (`L11820`, `TOWN-314`) reads the tip file `main\text\tips\chrsel1.txt` and sits among character-generator "loader" strings, identifying it as a character-generator sub-page… | High / Medium / Unknown | ● active | [EXP-0208](../experiments/EXP-0208-room-tip-screen-rects/) |
| TOWN-332 | Two distinct, non-interchangeable mechanisms compute a widget's own screen-paint origin in this engine: an ancestor-chain accumulator (`this+0x30`, walked by `R1224`/`R1212`, `TOWN-232`)… | High / Medium | ● active | [EXP-0208](../experiments/EXP-0208-room-tip-screen-rects/) |
| TOWN-333 | 33 of `claims/town.md`'s pre-existing 149 rows (`TOWN-001`..`TOWN-314`) give a screen rectangle traced to a literal operand of a construction or paint routine; of those… | High / Medium | ● active | [EXP-0208](../experiments/EXP-0208-room-tip-screen-rects/) |

### TOWN-312

The final/detailed stage's `+0x7c` child's four border-strip source fields (`this+0x68`/`+0x6c`/`+0x70`/`+0x74`, the blit sources `TOWN-234` reads) are populated by `R0133`, called with `this` set to the `+0x7c` child itself (not the top-level screen), from three fresh `graphics.res` loads and one cached global. `R1870` (`TOWN-245`) calls `R0133` at `L11263` with `this` loaded at `L11821` (a load of the dword at offset 0x7c of the top-level `this`); inside `R0133` (`this` taken at entry), before any class branch: `this+0x68` written at `L11822` from the call to `R1176` at `L11823` on string `L11824`=`graphics\interface\chrgen\RollStatsR.bmp`; `this+0x6c` at `L11825` from the call to `R1176` at `L11826` on `L11827`=`graphics\interface\chrgen\FullStatsR.bmp`; `this+0x70` at `L11828` from the call to `R1176` at `L11829` on `L11830`=`graphics\interface\inn\RUOver.bmp`; `this+0x74` at `L11831` (a store into the field at offset 0x74 of `this`) from a load of `[L11832]` at `L11833` — a direct global copy, no call.

The `this`-identity is cross-checked against the adjacent field `this+0x64` (written six instructions later, `L11834`), which matches `TOWN-256`'s already-published reading of the same object's mask field. `chrgen/fullstatsr.bmp` is loaded twice in the image, once here and once by `R2013` into the top-level screen's own `+0x60` (`TOWN-245`) — the same archive path, two destinations. `EnumRefs refto:L11832` (9 hits, 7 owners) finds one writer, `R1057` at `L11835`: `[L11832] = R1176("graphics\interface\HumanBackL.bmp")`, one of roughly 30 loads in a routine whose surrounding filenames (Crystal, Heads, CommandBar, Book, BackPack, SpellBook) are not chargen-specific; `EnumRefs disp:L11832` returns 0 hits (no register-relative access anywhere).

This resolves the reference sites `TOWN-284` lists as not reached for `chrgen/rollstatsr.bmp` and `inn/ruover.bmp` (`experiments/EXP-0207-chargen-final-stage/evidence/disasm-R0133-head.txt`, `R1057_R1057.c`, `refs-L11832.txt`, `str-humanbackl.txt`)

**Confidence.** High for the call's `this`-identity, the three fresh loads' exact paths and `this+0x74`'s mechanism (every instruction and literal operand read directly, cross-checked against `TOWN-256`); Unknown for whether the global at `L11832` is populated before the final stage runs (writer's own call chain traced two hops, not to completion)

### TOWN-313

The final/detailed stage's fourth top-level child (`TOWN-283`'s open item, "`TOWN-224` or `TOWN-235`, not distinguished") is the `+0x70` object — both already name the same routine, `R0833`, not two candidates — and identifying it refutes `TOWN-283`'s own coverage prediction: 76,800 of 307,200 screen pixels (25%) remain with no established painter. `TOWN-232`'s own body gives the `+0x70` child's construction record (`R1214(0x457,0,0,0xa0,0xee,this)`, vtable `L06316`, literal rect `(0,0)-(160,238)`) but decompiles no paint routine for it. `TOWN-235`'s own subject line names `R0833` "the final stage's `+0x70` stats-panel child", absolute rect `(0,0)-(160,238)` "per `TOWN-232`"; `TOWN-224`'s current (amended) row cites the same call site and states "`TOWN-235` resolves this same call site".

`TOWN-245`'s own headline cites `TOWN-232` for this child's paint routine, which is imprecise (`TOWN-232` establishes construction, not paint); `TOWN-283`'s own paraphrase, `TOWN-224`, is the technically correct citation for it. `TOWN-283`'s own rect union for its known three children is `160×4 + 160×238 + 320×480 = 192,320` px, against 307,200; its own stated uncovered region is `x:[0,160)` above `y=238` and below `y=242` (2×160×238=76,160px) plus `x:[480,640)` below `y=238` (160×242=38,720px) = 114,880px. Adding the `+0x70` child's own rect (160×238=38,080px) leaves 230,400px covered, 76,800px uncovered, in two disjoint regions: `x:[0,160)` `y:[242,480)` (38,080px) and `x:[480,640)` `y:[238,480)` (38,720px). The `+0x70` child's own single rectangle exactly fills only the `y:[0,238)` portion of the left-column gap and none of the right-column gap; `TOWN-283`'s own "consistent with the unfetched fourth child painting that remainder" does not hold once the child's own rect is known (`experiments/EXP-0207-chargen-final-stage/EXP-0207.md`, Q2)

**Confidence.** High for the fourth child's identity, rect and content (re-reading already-published High-confidence claims `TOWN-232`/`TOWN-224`/`TOWN-235`, no re-derivation); High for the rect-union arithmetic (direct subtraction over already-published rects); Unknown for what, if anything, paints the two still-uncovered regions

### TOWN-314

The character generator's tip popup (`TOWN-187`) is attached to the `+0x7c` child at construction, and the popup's own paint routine resolves its absolute rect through the same ancestor accumulator `TOWN-232` establishes: `(160,280)-(472,480)`, not `TOWN-187`'s own `(0,280)`. `R1870` constructs the popup at `L11836`, passes `this=[screen+0x7c]` (the `+0x7c` child, absolute rect `(160,0)-(480,480)`, `TOWN-232`) and calls `R0385` at `L11837` — an instruction already present in `TOWN-187`'s own cited evidence file (`experiments/EXP-0197-room-buttons-and-tips/evidence/disasm-chargen-tips-R1870-R1986.txt`, lines 157/164) but not read there as a parent-attach.

`R0385(this,child)` is `Add(child); *(child+0x30)=this;`. The popup class's own paint routine (vtable `L06362+0x30`=`R1975`) opens with `R1224(&local_14, param_1+8)`, the same accumulator `TOWN-232` reads. With the `+0x7c` child's own chain empty at construction, the walk applies one step: `(0,0x118,0x138,0x1e0)` + `(160,0)` = `(160,280)-(472,480)`, lower-center (`x:[160,472)`, straddling the screen's own midpoint at 320), not lower-left. **`R1986`'s once-only text replacement** (`this` = the dword at offset 0x80 of the screen object at `L11838`, the call to `R1987` at `L11839`) chains into `R0725`→`R0726(popup+8,text)`→`R0727`, which reads `popup+8`/`+0xc` only to compute `right−left` as a line-wrap width and writes through that pointer nowhere; no function in the chain writes `popup+8`/`+0xc`/`+0x10`/`+0x14` or `popup+0x30`; `R0725`'s only per-instance write is `popup+0x8c`, a derived line count. **All 7 `R1261` call sites** (`EnumRefs callto:R1261`, exhaustive, 0 orphan) attach a non-null parent immediately via `R0385`; 5 of 7 match already-published rooms at instruction level: chargen `(0,0x118,0x138,0x1e0)` parent the dword at offset 0x7c of the screen object (`TOWN-187`, above); shop `(0,0xa2,0x138,0x12a)` parent the dword at offset 0x74 of its owner object (`SHOP-TIP-045`); town `(0x148,0,0x280,0xc8)` parent the town view's own `this` ( `TOWN-165`/`015`/`025`); tavern `(0,0,0x138,0xc8)` parent the dword at offset 0x7c of the view (`TOWN-015`/`025`); school `(0,0,0x1c8,0xc8)` parent the school view's own `this` (`TOWN-021`/`025`/`188`). Two sites are unidentified: `L11819` in `R0701`, rect `(0xa,0x14,0x172,0xbc)`, parent the dword at offset 0xd0 of the campaign object; `L11820` in `R1475`, rect `(0xe8,0,0x280,0x88)`, parent the caller's own `this`. The uniform rule: construction never leaves the chain empty, so the literal LTRB is never itself the mechanism's absolute rect; whether it happens to equal the absolute rect depends on the parent's own accumulated origin, resolved here as non-zero for chargen and (already, `SHOP-TIP-045`) for shop. For town and school the parent is the room view's own `this` directly; for tavern it is the dword at offset 0x7c of the view, a child of the view, structurally like chargen's own site. Whether these three rooms' own already-published literal rects equal their absolute positions was not checked by `TOWN-165`/`015`/`025`/`021`/`188` or by this experiment — none of those rows states an ancestor walk was performed (`experiments/EXP-0207-chargen-final-stage/evidence/R0385_R0385.c`, `R1975_R1975.c`, `vt-L06362.txt`, `disasm-R1986.txt`, `R1986_R1986.c`, `R1987_R1987.c`, `R0725_R0725.c`, `R0726_R0726.c`, `R0727_R0727.c`, `callto-R1261.txt`, `callsites-R1261.txt`, `callsite-shop-R1261.txt`)

**Confidence.** High for the mechanism identity, the chargen site's own absolute-rect resolution, the 7-site enumeration's completeness, and each site's own literal operands and parent argument (all instruction-level, `R1986`'s full 5-function call chain read directly with no write to the rect or ancestor fields found); Medium for "construction never leaves the chain empty" as a general rule across `R1261`'s own 7 call sites (7 of 7 agree; the population is this constructor's own call sites, not every popup-like object in the image); Unknown for the two unidentified sites' own rooms, for whether town/tavern/school's own published literal rects equal their absolute positions, and for whether anything outside `R1870` rewrites the `+0x7c` child's own `this+0x30` (`TOWN-243` covers the five final-stage classes only, not the popup class)

### TOWN-328

Town: `R1383` attaches the popup via `R0385` at `L11840` with the view's own `this` as the parent (loaded at `L11841`) — the parent is the town view's own `this`, not a child object. School: `R0706` attaches at `L11842` with the view's own `this` as the parent (loaded at `L11843`) — the same pattern. Both views are built by `R0315` (campaign's own construction factory) as `R0602(0x3fc,[L09808],[L09809],[L09810],[L09811])` (town) and `R2021(0x3fc,...)` (school), the same four globals `TOWN-222` reads as `(0,0,640,480)` at the shipped default resolution; each constructor's own tail calls `R0759(left,top,right,bottom,0)` with a literal `0` sixth argument, forwarding to `R0773`, which unconditionally zeroes `this+0x30` (re-verified directly this experiment, matching `TOWN-232`'s own reading of the same routine) before the dead `if(parent!=0)` branch.

Neither `R0315`'s own body nor either room's own construction routine calls `R0385` on the view itself. Applying `R1224`'s walk: town's own literal `(0x148,0,0x280,0xc8)` (`TOWN-165`) plus the view's own `(L,T)=(0,0)` is unchanged, `(328,0)-(640,200)`; school's own literal `(0,0,0x1c8,0xc8)` (`TOWN-025`/`TOWN-021`) plus `(0,0)` is unchanged, `(0,0)-(456,200)`. This closes the question `TOWN-314` left open for these two rooms (`experiments/EXP-0208-room-tip-screen-rects/evidence/disasm-all-5-callsites-R1261.txt`, `R0315_R0315.c`, `R0602_R0602.c`, `R2021_R2021.c`, `R0759_R0759.c`, `R0773_R0773.c`)

**Confidence.** High for the attach instructions, the room-view construction calls, and the unconditional zero-write (all read at instruction level); Medium for "no code anywhere re-attaches the view", since only `R0315`'s own body and the two rooms' own construction routines were checked, not an exhaustive vtable-method-surface sweep of the kind `TOWN-243` ran for the chargen classes

### TOWN-329

`R1412` attaches the popup via `R0385` at `L11844` with the dword at offset 0x7c of the view as the parent (loaded at `L11845`) — the parent is the tavern's own center/roster panel (`TOWN-064`'s id `0x450`), not the tavern view. That panel is itself constructed by `R0601(0x450,0xa0,0,0x1e0,0x1e0,param_1)` inside `R1939` (literal rect `(160,0)-(480,480)`) and is separately, explicitly re-attached to the tavern view by its own `R0385` call at `L11846` (the view's own `this` as the object and the dword at offset 0x7c of the view pushed as the argument), so the panel's own `this+0x30` is the tavern view. The tavern view is constructed the same way as town/school (`R1996(0x44c,[L09808],[L09809],[L09810],[L09811])`, `R0759(...,0)`, ancestor empty, no re-attach found in `R0315`).

Walk: popup's own literal `(0,0,0x138,0xc8)` (`TOWN-015`) plus the panel's own `(L,T)=(160,0)` plus the view's own `(0,0)` gives `(160,0)-(472,200)`. The published literal `(0,0)-(312,200)` is the popup's own rect relative to the center/roster panel, not its screen position (`experiments/EXP-0208-room-tip-screen-rects/evidence/disasm-tavern-tip-popup-construct-attach.txt`, `disasm-tavern-children-construct-attach.txt`, `R1939_R1939.c`, `R0315_R0315.c`, `R1996_R1996.c`)

**Confidence.** High for every attach instruction and every construction call cited (instruction-level); Medium for "no code re-attaches the tavern view or the panel outside the two call sites read", the same scope caveat as `TOWN-328`

### TOWN-330

`R0701`'s unmatched `R1261` call site (`L11819`, `TOWN-314`) belongs to the campaign screen's own message handler, and its popup's absolute rect equals its published literal `(0xa,0x14,0x172,0xbc)`: the parent, the dword at offset 0xd0 of the campaign object, is campaign's own map-view widget, whose ancestor chain is empty and adds nothing. `R0701` is reached only through a 4-slot virtual-dispatch table at `L11847` (`FindVtables`); its own field offsets (`+0xcc`,`+0xd0`,`+0xd4`, read as fields of `this` at `L11848`) match campaign's own layout in `R0315`, and the `this` pointer kept from entry (`L11849`) confirms the base is `this`. `campaign+0xd0` is written exactly once, at `R0315`'s own `R0391(0,0,[L00618]-0xa0,[L01259])` call, the map view `TOWN-086` already names (id `1`); `R0391` calls `R0773(1,left,top,right,bottom,0)` with a literal `0` parent — ancestor empty, the same unconditional-zero mechanism as town/school/tavern's own views — and `R0315`'s own body has no `R0385` call re-attaching it.

Walk: popup's literal `(0xa,0x14,0x172,0xbc)` plus the map view's own `(L,T)=(0,0)` is unchanged. No instruction assigning vtable `L11847` to an object (`*this=&vtable`) was found by an image-wide `imm:`/`refto:` scan of that address; the class identity rests on the field-offset match against campaign's own known layout plus the vtable-slot's own location, not a direct construction-site read (`experiments/EXP-0208-room-tip-screen-rects/evidence/R0701_R0701.c`, `disasm-R0701-entry.txt`, `vt-L13171-R0701.txt`, `R0315_R0315.c`, `R0391_R0391.c`)

**Confidence.** High for the map view's own construction and ancestor facts and for "if the parent is campaign's own map view, its absolute rect equals its literal" (instruction-level); Medium for "this popup site belongs to the campaign screen" (field-offset match only, no direct vtable-write instruction found); Medium for "no code re-attaches the map view", same scope caveat as `TOWN-328`/`TOWN-329`

### TOWN-331

`R1475`'s unmatched `R1261` call site (`L11820`, `TOWN-314`) reads the tip file `main\text\tips\chrsel1.txt` and sits among character-generator "loader" strings, identifying it as a character-generator sub-page; its class's own top-level constructor was not found, so its popup's absolute rect is Unknown. `R1475` pushes `s_main_text_tips_chrsel1_txt_L11850` (`main\text\tips\chrsel1.txt`) at `L11851`..`L11852` before the same the global at `L03631`-gated `R1261(0x467,0xe8,0,0x280,0x88,uStack_14)` call `TOWN-314` names; the surrounding string table (`L11853`-`L11854`) is entirely chargen-related (`graphics\interface\chrgen\loader\...`, `graphics\interface\chrgen\PreCreate\Levels\...`, and the placeholder names `Master Oberic`/`Unnamed`), and the filename `chrsel1.txt` reads "character select 1" — consistent with, though not matched by address to, one of the "two character-select screens" `TOWN-188` names among the tip-reader's unread callers.

`R1475`'s own class was searched for: `EnumRefs callto:R1475` and `refto:R1475` return 0 hits image-wide (virtual-dispatch only). `FindVtables` locates it at `+0x20` of a 15-slot candidate vtable (`L11855`); an image-wide `imm:L11855` scan (the vtable's own address as a literal operand, the shape a `*this=&vtable` write would take) returns 0 hits, and `refto:L11855` also returns 0 hits. The sibling candidate vtable `L06361` has 2 callers (`R2022`, `R2023`), both chaining through `R1253` to `R2024`, whose own allocation is 368 bytes — too small for this class, which reads fields to byte offset `0x1e0` (a 480-byte minimum object).

No live constructor for either candidate vtable was found. The parent argument (`TOWN-314`) is this class's own `this`, cached at `this+0x1c0`; without the class's own construction site, neither this widget's own literal `(L,T)` nor its own ancestor chain can be walked (`experiments/EXP-0208-room-tip-screen-rects/evidence/R1475_R1475.c`, `disasm-R1475-entry.txt`, `strdump-chrsel-region.txt`, `vt-L11855-L06361-candidates.txt`, `refto-L11855-L06361.txt`, `callto-L06361-chain1.txt`, `callto-L06361-chain2.txt`, `callto-L13172-L13173-dead.txt`, `imm-L11855-zero-hits.txt`)

**Confidence.** High for the string identity (`chrsel1.txt` and the surrounding chargen strings, read directly); Medium for "this is one of `TOWN-188`'s two unread character-select sites" (plausible from the filename and context, not matched by address); Unknown for the class's own construction site and therefore for the popup's absolute rect — searched by `callto:`, `refto:`, and `imm:` against both candidate vtables and by object-size elimination against the one live caller found, none identifies a constructor

### TOWN-332

Two distinct, non-interchangeable mechanisms compute a widget's own screen-paint origin in this engine: an ancestor-chain accumulator (`this+0x30`, walked by `R1224`/`R1212`, `TOWN-232`), used by the chargen final-stage children and by every `R1261` tip popup; and a single-hop parent-pointer addend (`this+0x5c`, read once and added directly to blit/text-draw coordinates inside the specific paint routine, no chain walk), used by the tavern's button/decoration panel (`R1904`, `TOWN-214`/`TOWN-222`) and the school's training-button widget (`R1905`, `TOWN-183`/`TOWN-280`). The tip popup's own use of the chain mechanism is confirmed directly, not by analogy: the popup's own paint routine (`R1975`, vtable `L06362+0x30`) opens with `R1224(&local_14,param_1+8)`, the exact call `TOWN-232` reads for the chargen children.

The two fields are written by different code: `this+0x30` only by `R0385`'s `*(child+0x30)=this`, unconditionally zeroed at construction by `R0773` and left empty unless a later `R0385` call attaches it; `this+0x5c` written directly by a constructor's own body from its own sixth argument (`TOWN-222`: `R1535` writes `param_1[0x17]=param_7`), with no chain walk at paint time — the routine reads `[this+0x5c]->+8/+0xc` once and stops. Applying the chain-walk mechanism to the tavern's button panel would find no `this+0x30` write to walk, since that panel's own origin field is `this+0x5c`; applying the single-hop addend to a chargen child or a tip popup would stop after one level where the mechanism itself performs a multi-level walk whenever the chain is non-empty, as `TOWN-329` shows for the tavern's own tip popup.

No widget using both fields, and no third mechanism, was found, but this experiment did not exhaustively enumerate every paint routine in the image — only the routines this experiment's own two questions and `TOWN-333`'s sweep of `claims/town.md` touch (`experiments/EXP-0208-room-tip-screen-rects/evidence/R1975_R1975.c`)

**Confidence.** High for both mechanisms' own existence and their distinct write/read sites (every instruction cited is read directly by this experiment for the chain family, and by the cited already-published rows, re-read here, for the single-hop family); Medium for "no widget uses both" and "no third mechanism exists", since the search was not exhaustive over every paint routine in the image

### TOWN-333

33 of `claims/town.md`'s pre-existing 149 rows (`TOWN-001`..`TOWN-314`) give a screen rectangle traced to a literal operand of a construction or paint routine; of those, 8 establish the coordinate space as absolute by the row's own evidence or a row it directly cites, 6 name the space as parent/widget-relative or origin-unresolved without resolving it, 1 is a retracted row whose original coordinate-mechanism reading was wrong, and 18 are silent on the space (the literal operand is given, but the row neither states nor computes whether it equals the screen position). Selection: every table row of `claims/town.md` matched by `grep` for `rect (`, `LTRB`, `container-local`, `own rect`, `absolute rect`, `construction rect`, `client rect`, `hit-test region`, `at (`, `own (L,T`, `destination`, `origin add`, or `paint-time`, then read individually and kept only if its own text gives numeric geometry sourced from a literal construction/paint-routine operand — excluding rows that only name a field's own storage offset without literal contents (`TOWN-182`), cite another row's rect without new geometry (`TOWN-089`,`TOWN-096`,`TOWN-242`,`TOWN-244`,`TOWN-260`,`TOWN-261`,`TOWN-284`), give a bitmap source path rather than a destination (`TOWN-312`), or draw geometry from a registry-built data table rather than a code literal (`TOWN-038`,`TOWN-039`,`TOWN-040`,`TOWN-160`). The full 33-row table, with each row's own space classification and citation, is in `experiments/EXP-0208-room-tip-screen-rects/EXP-0208.md` ("Q3 enumeration"). Absolute, established: `TOWN-214`,`TOWN-222`,`TOWN-232`,`TOWN-234`,`TOWN-235`,`TOWN-280`,`TOWN-281`,`TOWN-313`. Named but unresolved by the row itself: `TOWN-091`,`TOWN-092`,`TOWN-093` (explicitly "container-local"), `TOWN-183`,`TOWN-257`,`TOWN-259` (explicitly "origin+" notation or the addend mechanism named without a resolved value at the time the row was written; `TOWN-183`/`TOWN-257`/`TOWN-259` are later resolved by `TOWN-280`, a separate row). Retracted, wrong mechanism: `TOWN-224` (its original reading of `R1224` as a plain 2-dword copy, not a chain, was overturned by `TOWN-232`). Silent on space: `TOWN-015`,`TOWN-021`,`TOWN-025`,`TOWN-064`,`TOWN-068`,`TOWN-086`,`TOWN-094`,`TOWN-165`,`TOWN-185`,`TOWN-187`,`TOWN-223`,`TOWN-233`,`TOWN-236`,`TOWN-246`,`TOWN-258`,`TOWN-282`,`TOWN-283`,`TOWN-314` — two of these (`TOWN-282`,`TOWN-283`) are aggregator rows that explicitly name their own remaining gap ("no claim...gives the per-cell destination rectangle in absolute screen coordinates, only the byte-offset formula") rather than silently omitting it, and one (`TOWN-314`) is this same experiment's own foundation row, five of whose seven sites this experiment resolves in `TOWN-328`..`TOWN-330`. This experiment's own new rows (`TOWN-328`..`TOWN-332`) are excluded from the swept population, which is fixed at the 149 rows that existed when the sweep was run, to avoid the ledger enumerating its own concurrent edits; `TOWN-328`,`TOWN-329`,`TOWN-330` each independently add an absolute-established row by the same standard applied above. This selection is a keyword sweep of one ledger only, read to the end of the table at 149 rows; a row phrasing geometry without any swept keyword (for example, by giving only a bitmap's pixel dimensions, or a formula with no literal operand shown) would not be found by it, and `claims/shop.md`, `claims/session.md` and every other format ledger were not swept

**Confidence.** High for the classification of every row named above as absolute, relative, retracted, or silent (each row's own text was read via `go run ./tools/claim` and quoted directly); Medium for "33 rows total", bounded by the keyword sweep's own blind spot stated above, not by a full-text scan for numeric-pair patterns of every possible phrasing

## Character figure widget

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-344 | The character figure widget (`campaign+0xe0`, id 7, vtable `L00722`) owns eight interactive behaviours and no child controls, and the set is fixed in code rather than built from data. | High | ● active | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/) |
| TOWN-345 | Every rectangle this widget tests is expressed relative to an ancestor-accumulated origin, not in absolute screen coordinates, and the widget's own stored rect is parent-local. | High | ● active | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/) |
| TOWN-346 | The six rectangles of `R0332` and their gates, panel-relative to the accumulated rect `(L,T,R,B)`. | High | ● active | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/) |
| TOWN-347 | What each of the six rectangles does: five post an engine message through `vt+0x48`, one posts a Win32 message through `PostMessageA`, and one also writes a map-view field directly. | High / Medium / Unknown | ● active (amended, partially retracted) | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/), **[EXP-0230](../experiments/EXP-0230-non-music-sfx-events/)** |
| TOWN-348 | A held item pre-empts all six rectangles: `R0332` inspects `[sess+0x3cc]+0x18` before any point-in-rect call and returns early on six of its values, in two arms. | High | ● active | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/) |
| TOWN-349 | The widget's principal control surface is not a rectangle: it is a per-pixel slot-id map in the second of its two `0xa0` x `0xf0` surfaces, sampled by double-click and by left-button drag. | High | ● active | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/) |
| TOWN-350 | `vt+0x78` (`R0339`) is the slot pick-up and it names every field it moves; `vt+0x7c` (`R0240`) is its inverse. | High / Medium | ● active | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/) |
| TOWN-351 | `this+0x5c` is the map view, not the parent: it is written only by message `0x403`, posted with `wParam = campaign+0xd0` to `campaign+0xd4` from three call sites of one shape. | High | ● active (amended, partially retracted) | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/), **[EXP-0230](../experiments/EXP-0230-non-music-sfx-events/)** |
| TOWN-352 | `campaign+0x3dc` is a bitmask, not an enumeration, and bit 1 is the shop; a shop entered from a mission leaves at least bits 0 and 1 set. | Medium | ● active | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/) |
| TOWN-353 | The shop screen carries the same regions because it carries the same object: there is one instance of this widget, re-parented, and every screen difference is a runtime branch on `campaign+0x3dc`. | High | ● active | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/) |
| TOWN-354 | Each of the six rectangles has its own art, and the art for the two picker rectangles is drawn under the same mask that gates their hit test, with one further exclusion the hit test does not share (below)… | High | ● active | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/) |
| TOWN-355 | Three gate mismatches leave a live rectangle undrawn, and two rectangle pairs overlap with no early return, so one click can post two different messages. | High / Medium | ● active | [EXP-0210](../experiments/EXP-0210-figure-panel-controls/) |
| TOWN-356 | `TOWN-354`'s eleven blit-source globals are loaded by one shared resource loader in address-shifted blocks: each block's own store is the PRECEDING block's object, not the object its own string builds. | High | ● active | [EXP-0212](../experiments/EXP-0212-figure-sources/) |
| TOWN-357 | The paint routine reaches `this+0x74` and `this+0x78` through one virtual call, `vt+0x80`, and the callee independently reads `member`'s equipment array at the offset `TOWN-350` already publishes. | High | ● active | [EXP-0212](../experiments/EXP-0212-figure-sources/) |
| TOWN-358 | The per-pixel writer of `this+0x78`'s slot-id map: a two-level virtual dispatch from the compositor to a literal byte store at the destination pointer minus one at `L04344`. | High / Medium | ● active | [EXP-0212](../experiments/EXP-0212-figure-sources/) |
| TOWN-359 | `member`'s uniform 12-slot equipment array is not expressible as `ITEM-EQUIP-006`'s `actor` layout under any single pointer offset, refuting a this-adjusted-alias model between the two objects under either live slot-numbering reading. | High / Unknown | ● active | [EXP-0212](../experiments/EXP-0212-figure-sources/) |
| TOWN-360 | A second image-wide indexed writer of `+0x15c`, in the same convention `TOWN-350` documents, at an untraced call site. | High / Unknown | ● active | [EXP-0212](../experiments/EXP-0212-figure-sources/) |

### TOWN-344

The class overrides fifteen vtable slots: `+0x04` `R2025` (destructor), `+0x14` `R0734`, and the thirteen slots below `+0x2c`..`+0x80`. Slot `+0x84` is `00000000`, terminating this class's own table; `+0x88` begins widget 8's own vtable (`TOWN-093`), a second table read by a naive code-region-address scan as two more slots of this class (`+0x8c`, `+0x9c`) when they are not, and by the same error `+0x04` was left off the count entirely. The ones that take input are `+0x4c` `R2026` (`WM_MOUSEMOVE`), `+0x54` `R2027` (`WM_LBUTTONDOWN`), `+0x58` `R0332` (`WM_LBUTTONUP`), `+0x5c` `R2028` (`WM_LBUTTONDBLCLK`), `+0x60` `R2029` (`WM_RBUTTONDOWN`), `+0x64` `R2030` (`WM_RBUTTONUP`), `+0x68` `R2031` (`WM_RBUTTONDBLCLK`) and `+0x6c` `R1249` (`WM_KEYDOWN`).

`R2027`, `R2029` and `R2031` each return the constant 1 and clean 0xc argument bytes, and do nothing. The behaviours are: six rectangles tested in `R0332`, one per-pixel region sampled by `R2028` and `R2026`, and one coordinate-free `WM_RBUTTONUP` that posts `0x405` to `this+0x5c` when `sess+0x3dc == 1` (`L11856`). Six rectangles, one per-pixel region and one coordinate-free control is eight, not seven. `WM_KEYDOWN` adds a keyboard route to rect C's message, the third of the six (`TOWN-346`), not the sixth: `R1249` posts `0x412` (`L11857`), and rect C is the poster of `0x412` among the six (`TOWN-347`). **No child controls exist:** every rel32 call site of `R0385` AddChild (235 sites), `R0384` RemoveChild (60) and `R1252` RemoveAllChildren (8) was enumerated and classified by containing address, and none lies in code-region; the base Control constructor `R0773` has 33 sites, of which 3 are in code-region and are the sibling widget constructors `L11858`, `L11859` and `L11860`.

A whole-image raw dword scan for each of the four addresses returns 0 occurrences, so no vtable or pointer table reaches them and the call-site enumerations are complete. In all 235 AddChild contexts this widget appears only as the pushed child, at `L11861` and `L11862`, both with `campaign+0xd4` as parent (`evidence/vtable-L00722.txt`, `callctx-R0385-addchild.txt`, `callto-child-api.txt`, `dwordscan-child-api.txt`, `listing-trivial-slots.txt`)

**Confidence.** High (complete rel32 enumerations with a zero-result pointer-table cross-check, and every cited handler tiled end to end)

### TOWN-345

`R0332` (`L11863`), `R2028` (`L11864`), `R2026` and `R0389` (`L11865`) each begin with `R1224(&local, this+0x8)`; none reads `this+0x8` directly. `R1224` calls `R1212` once for the top-left pair and once for the bottom-right pair; `R1212` copies the pair then walks `this+0x30`, adding each ancestor's `(+0x8,+0xc)` through `R1213(p,dx,dy){p[0]+=dx;p[1]+=dy;}`. `R0385` sets `child+0x30 = parent` and `R0384` clears it, so the chain is non-empty whenever the widget is parented. This is `TOWN-232`'s case, established here from the accumulation primitive rather than by analogy.

The constructor's literal rect is `(0,238,160,480)` and `campaign+0xd4`'s origin is `(screenW-0xa0, 0)` per `TOWN-086`, so on the mission screen at 640x480 the accumulated rect is `(480,238,640,480)`. `B = T + 0xf2` at every resolution, because the constructed vertical pair is literal and neither the shop transfer of `SHOP-FIGURE-041` nor any resolution arm changes the height (`evidence/listing-R1224-accumulate.txt`, `listing-R0332-lbuttonup.txt`)

**Confidence.** High (the accumulator and both callers read at instruction level; the absolute value follows from `TOWN-086`'s already-published container origin)

### TOWN-346

In the order tested, with the point-in-rect import `[L01178]`: **A** `(L, B-0x28, L+0x1c, B)`, 28x40, gate `sess+0x3dc & 1`, tested at `L11866`; **B** `(L, T, L+0x1c, T+0x24)`, 28x36, gate `sess+0x3dc & 3`, `L11867`; **C** `(L+0x80, T, R, T+0x24)`, 32x36, gate `!(sess+0x3dc & 0x600)` and `flag == 0`, `L11868`; **D** `(L+1, T+0xcd, L+0x21, T+0xed)`, 32x32, gate `sess+0x3dc & 0x226`, `L11869`; **E** `(L+0x77, T+0xcd, L+0x97, T+0xed)`, 32x32, same gate, `L11870`; **F** `(L+0x7e, T+0xce, L+0x9e, T+0xee)`, 32x32, gate `sess+0x3dc & 1`, `L11871`. Every one of the twenty-four bounds is either an accumulated value itself or a constant displacement from one, formed by a copy, an address computation or an addition between `L11872` and `L11873`; two of the twenty-four are formed before `L11874`, rect A's `B` at `L11875` (the accumulated `B` restored unchanged) and rect A's `T` at `L11876` (a base formed as another running value minus 0x28 at `L11872`).

No table and no data operand participates anywhere in the window, so the population cannot vary with content. `flag` is 1 when `screenH - B > B - T` and `sess+0x3dc == 1`, computed at `L11877`..`L11878` with `screenH = [L01259]`; `B - T` is always `0xf2`, so with `B = 480` on the mission screen the condition is `screenH > 722`, which of the three resolutions `SESS-VIEW-028` enumerates only 1024x768 satisfies. That last step assumes the containers above `campaign+0xd4` contribute no vertical origin, which was not read. **No arm returns between the six tests:** A jumps to `L11879` which is B's block, B falls to `L11880`, C falls to `L03383`, D falls into E's test at `L11881`, E jumps to `L11882` which is F's gate, and the routine returns 1 at `L11883`.

`SHOP-PICKER-043` names three of these six, D, E and F, correctly and as the three 32x32 rectangles; A, B and C are 28x40, 28x36 and 32x36 and were not previously published (`evidence/listing-R0332-lbuttonup.txt`)

**Confidence.** High (every bound, gate and branch target read from a listing that tiled exactly; D and E reproduce `SHOP-PICKER-043`'s published panel-relative pair by a different tool)

### TOWN-347

`wParam` and `lParam` are 0 at every post site. **A** posts `0x40e` to `this` (`L11884`) and to `this+0x5c` (`L11885`). **B** posts `0x40f` to `this` (`L11886`), to `this+0x5c` (`L11887`), and, only when `sess+0x3dc & 2`, to `sess+0xf0` (`L11888`). **C** posts `0x412` to `this` (`L11889`) and then executes a 32-bit store of 1 to offset 0xe0 of the object at `[this+0x5c]`, a pointer cached at `L11890` and unchanged through `L01512`. **D** posts `0x414` to `sess+0xcc` (`L11891`). **E** posts `0x415` to `sess+0xcc` (`L11892`). **F** calls `PostMessageA([sess+0x1c], 0x416, 0, 0)` through `[L03535]` (`L11893`). **Amended by [EXP-0230](../experiments/EXP-0230-non-music-sfx-events/): D, E and F each select registry slot 1 (`click00`) through `[L01636][1]`, pass that sample as `this` to `R0386`, and push `0xdc` as the terminal's priority/category byte; `0xdc` is not the sound or slot (`VIDEO-SFX-013`, `VIDEO-SFX-016`).** `SHOP-PICKER-043` establishes the receivers of `0x414` and `0x415` at the shop view's own message table, not at `sess+0xcc` (`sess+0xcc` is only the post's target address, the shop view itself).

`0x40e` and `0x40f` reach the map view, whose art predicate the widget itself reads: rect A's art switches on `R0375(mapview, 2) != 0` and rect B's on `R0375(host, 3) != 0`, where `host` is `sess+0xf0` when `sess+0x3dc & 2` and the map view otherwise (`R0380`, `R0317`). `0x416`'s consumer was not located: an image-wide dword census finds two `PUSH imm32` sites, `L11894` here and `L11895` under `sess+0x3dc == 1`, three `.data` dwords belonging to the CRT locale-name to LCID table near `L11896`, and one call-operand coincidence; no 32-bit compare against `0x416` exists anywhere in the image, so the receiver dispatches by subtraction and table, and the window procedure behind `[sess+0x1c]` was not identified (`evidence/listing-R0332-lbuttonup.txt`, `imm-messages.txt`, `ctx-other-posters.txt`, `listing-R0380-R0317-predicates.txt`)

**Confidence.** High for every post site, target expression, message number and the direct field write; **Medium** for `0x40e` and `0x40f` toggling the presence of child id 2 and child id 3, which is inferred from the widget's own paint predicate and not read from the map view's handler; **Unknown** for `0x416`'s consumer, searched as stated

**Amended.** [`retracted.md`](retracted.md) holds an entry for this claim (refuted).

### TOWN-348

At `L11897` the routine tests `sess+0x3cc`, the cursor-held item. When non-zero it reads `[sess+0x3cc]+0x18`: on 1 or 2 it plays `[L01211]` through `R0320`, computes `slot = ((*(u16 *)([sess+0x3cc]+6) >> 8) & 0xf) - 1`, calls `this->vt+0x7c(slot)` (`L11898`) and returns 1; on a value in the inclusive range 5 to 8 it calls `R1781` with `this = [sess+0xf0]` (`L11899`) and returns 1; on 0, 3, 4 or above 8 it falls through to the rectangle tests at `L03375`. The drop slot is therefore taken from the held item's own bytes, not from the cursor position (`evidence/listing-R0332-lbuttonup.txt`)

**Confidence.** High (every branch and operand read at instruction level)

### TOWN-349

The constructor allocates `+0x74` through `R0744` (`L09791`) and `+0x78` through `R2032` (`L09792`), both `0xa0` x `0xf0`. `R2028` (`WM_LBUTTONDBLCLK`) and `R2026` (`WM_MOUSEMOVE`) test no rectangle. Both accumulate the rect, then read one byte at `[[this+0x78]+0x10] + (y-T)*160 + (x-L) - 0x138` (`L11900`, `L11901`). The paint routine treats `[surface+0x10]+8` as the first pixel (`L11902`), and `-0x138 = 8 - 2*160`, so the byte is element `(y-T-2)*160 + (x-L)` of a 160-wide map. A zero byte ends the handler; a non-zero byte `p` calls `this->vt+0x78(p-1)`. `R2026` requires `wParam & 1` first (`L11903`, a test of the low byte against 0x1), so it is the drag path.

Both require `[this+0x5c]+0x140 == 1` and `[[this+0x5c]+0x138]+0x18c & 1`; `R2028` (`WM_LBUTTONDBLCLK`) has no reference to `sess+0x3cc` anywhere in its tiled listing, and only `R2026` (`WM_MOUSEMOVE`) additionally tests `sess+0x3cc == 0` (`L11904`). `R2028`, when `vt+0x78` returns non-zero, additionally calls `[sess+0xe8]->vt+0xa4(*(void **)([sess+0xe8]+0x90))` (`L11905`). A second model for the byte, an alpha or colour-key channel, is refuted by the consumer: the value is decremented and used as the subscript in `member + slot*4 + 0x15c` (`evidence/listing-R2028-dblclick.txt`, `listing-R2026-mousemove.txt`, `listing-R1763-ctor.txt`)

**Confidence.** High (the address arithmetic, both gates and the consumer's use of the value are instruction-level; the decomposition into `(x-L, y-T-2)` is arithmetically identical to the folded constant)

### TOWN-350

`R0339(slot)` returns 0 unless all four hold: `sess+0x3cc == 0` (`L11906`); `sess+0x3dc & 2` or `sess+0x3dc == 1` (`L11907`, `L11908`); `[this+0x5c]+0x140 == 1` (`L11909`); `[[this+0x5c]+0x138]+0x14 == [[this+0x5c]+0x9b4]` (`L11910`, `L11911`). With `member = [[this+0x5c]+0x138]`, a slot below 2 is refused when `member+0x74` is 3, 7 or 8 (`L11912`..`L11913`). Otherwise it sets bit 3 of `member+0x18c` (`L11914`), posts `0x408` to itself, moves `[member + slot*4 + 0x15c]` into `sess+0x3cc` (`L09862`), zeroes that slot (`L11915`), writes `sess+0x3d0 = slot` (`L09865`) and `sess+0x3d4 = this->vt+0x80()`, which is the constant 1 (`R1782` returns the constant 1), calls `R0551(member)` and posts `0x46d` to `sess+0xcc` with `wParam = [this+0x4]` (`L11916`), the widget's own id field, 7.

`R0240` places `sess+0x3cc` into `member+0x15c + slot*4`, with arms for stacking and splitting (`R2033`, `R0747`), for two-handed displacement between `member+0x15c` and `member+0x160`, for `sess+0xe8`'s `vt+0xa4`, `vt+0x78` and `vt+0x94`, and for `R0238` on `sess+0xd0`; it ends at `R0340`, which zeroes `sess+0x3cc` and `sess+0x3d8` and sets `sess+0x3d0` and `sess+0x3d4` to -1. The equipment array is therefore `member+0x15c`, indexed by the slot id the pixel map supplies (`evidence/listing-R0339-pickup.txt`, `listing-R0240-drop.txt`)

**Confidence.** High for `R0339`'s four preconditions, its refusal rule and all five of its writes, all read at instruction level; **Medium** for the enumeration of `R0240`'s arms, which were identified by their calls and their destination fields rather than read to the end of each arm

### TOWN-351

`R1205` (`vt+0x48`) calls the base dispatcher `R0390` first and dispatches its own arms only when the base returned 0. Its dispatch is `msg - 0x402`, rejected above `0x10`, through a 17-byte index table at `L11917` (offsets 0, 1, 6, 12, 13, 14 and 16 select arms 0 to 6 in that order; every other offset selects arm 7) into an 8-entry arm table at `L11918` (`L11919, L11920, L06300, L11921, L11922, L06300, L11923, L11924`), both read as raw bytes. Resolved: `0x402` calls `vt+0x34` when `sess+0x3dc` is 1 or 3 and `this+0x60 != 0`, then `R0338` when `sess+0x3dc & 3`; `0x403` sets `this+0x5c = wParam` and `this+0x60 = 1`; `0x408` and `0x410` set `this+0x60 = 1` (index-table offset 11, message `0x40d`, resolves to the default arm, not this one); `0x40e` inverts `this+0x68`; `0x40f` inverts `this+0x6c`; **amended by [EXP-0230], `0x412` selects registry slot 1 (`click00`) and passes it as `this` while `0xdc` is the priority/category byte**, then inverts `this+0x70` and sets `this+0x60 = 1`, all three skipped when `flag != 0` (the sample pointer at `L11923`, computed at `L11925`..`L11926`); every other message takes the default arm.

A whole-image raw-dword census of `0x403` finds six `PUSH imm32` sites; three post the shape `wParam = [sess+0xd0]` to `[sess+0xd4]`: `R1660` at `L11927`, and two further sites inside `R1321` (`L11928`) and `R0909` (`L11929`). A fourth poster, `L11930`, targets `[[ebp-0x14]+0xd4]` with a different `wParam` and is not this class's session. The remaining two, `L11931` and `L11932`, post `0x403` through a two-argument call to an imported function pointer at `[L11933]`, not through the four-argument shape the other four share; this row does not account for them, and their owning object was not traced. `campaign+0xd0` is `SESS-VIEW-028`'s map view.

A second model, that `+0x5c` is the parent pointer, is refuted: the engine's parent field is `+0x30`, written only by `R0385`. `this+0x68` and `this+0x6c` are written by those two arms and read by no routine of this class, including the tiled paint routine. `this+0x70` is read four times by the paint routine (`L11934`, `L11935`, `L11936`, `L11937`) and is forced to 1 in three places outside the paint routine: the constructor (`L11938`), and two sites inside the shop-open call chain, `L11939` (in `R1321`) and `L11940` (a 32-bit store of 1 at `+0x70` of the object held in the field at `+0xe0` of the receiver, inside a routine starting `R1316` that is called from the shop-open path but is not `R1315` itself; `L11941` fourteen bytes later is `call R0385`, AddChild, not a store) (`evidence/listing-R1205-msg.txt`, `bytes-L11918-msgtables.txt`, `listing-R1149-mainframe.txt`, `listing-R2025-dtor.txt`, `listing-R0734-slot14.txt`)

**Confidence.** High for the tables, every arm's effect, and `this+0x5c` receiving `campaign+0xd0` (three call sites of one shape agree, not one as first stated); High for `this+0x68` and `this+0x6c` having no reader inside this class, checked by tiling every one of the fifteen override slots this class actually has (`+0x04`, `+0x14`, `+0x2c`, `+0x48`, `+0x4c`, `+0x54`, `+0x58`, `+0x5c`, `+0x60`, `+0x64`, `+0x68`, `+0x6c`, `+0x78`, `+0x7c`, `+0x80`), not the sixteen the report first listed

**Amended.** [`retracted.md`](retracted.md) holds an entry for this claim (refuted).

### TOWN-352

Fifteen writers read at instruction level (the sentence below lists fifteen sites, not fourteen): `= 1` at `L06481`; OR with `2` at `L11942`, `4` at `L11943`, `8` at `L03329`, `0x10` at `L11944`, `0x20` at `L11945`, `0x40` at `L11946`, `0x80` at `L11947`, `0x100` at `L11948`, `0x400` at `L11949`, `0x1000` at `L11950`, `0x4000` at `L11951`; AND NOT `0x2000` at `L09240`; `= 0` at `L11952` and `L11953`. **This is not every writer of the displacement.** A whole-image scan of every instruction encoding disp32 `0x3dc` (`tools/figurecontrols -mode field -val 0x3dc`) finds 242 instructions total and 58 that write to `[reg+0x3dc]` in an executable section, of which the fifteen above are a subset.

Concrete misses on the widget's own object, read at instruction level: `L11954` (a mask of the low byte with 0xfe at `L09322`, a bit-0 clear inside `R1660`, the same routine the row's prose already names for clearing bit 0, but the site itself is absent from the enumeration); `L11955`; `L11956` (the value combined with 8 by bitwise or) and `L03317` (a mask of the low byte with 0xf7, a bit-3 clear); `L03360` (a mask of the low byte with 0xf7, a second bit-3 clear); `L11957`. A further site, `L11958` (the value combined with 0x2 by bitwise or at `L11959`, inside `R2034`), is reached by a call at `L11960` from inside the very `R1315` this row cites, immediately after the direct bitwise-or with 2 at `L11942`; whether its target is `campaign+0x3dc` itself or a different object's own field at the same displacement was not traced past the call.

Consumers test the field as a mask: `0x627` gates the whole paint routine (`L03384`), `0x226` gates two rectangles and their art, `0x600` gates a third. An enumeration model is refuted because an OR of 2 applied to the value 1 produces 3, which no enumeration would define; the population of 58 writers, mixing `OR` and `AND`-clear forms, is consistent with a bitmask and does not revive an enumeration model. Bit 1 is the shop: `R1315` calls `AddChild(sess+0xcc, sess+0xf0)`, ORs `2` into `sess+0x3dc`, calls `[sess+0xf0]->vt+0x80()`, and later builds the string at `L11961`; it does not clear bit 0, so `sess+0x3dc` carries at least bits 0 and 1 while a shop opened from a mission is on screen.

`R1149` sets the value to 1 when it builds the mission frame; `R1660` clears bit 0 (`evidence/listing-R1315-shopopen.txt`, `listing-R1149-mainframe.txt`, `listing-R0389-paint.txt`)

**Confidence.** Medium (fifteen of 58 image-wide writers to this displacement were read at instruction level; the fifteen and the mask tests are correct as far as they reach, but "one bit per surface" is not established as exhaustive, and whether every remaining writer targets this same object was not checked site by site)

### TOWN-353

`SHOP-FIGURE-041` establishes that shop activation moves `campaign+0xe0` into `view+0x7c`, offsets its stored rect and adds it to the shop view, and that deactivation is the exact inverse. The AddChild census of `TOWN-344` finds no other site that parents this object and no child of its own. This class has three constructor overloads, all writing vtable `L00722`: `R1763` (already `SHOP-FIGURE-041`'s published constructor for this widget), whose own `call R0773` at offset `0x3c` (`L11859`) reaches the base constructor (33 call sites image-wide, 3 in code-region, all three sibling widgets); `L11962` calling `R1219`; and `L11963` calling `R1218`.

The latter two have 0 rel32 call sites and 0 raw dword occurrences each, image-wide: dead code, never instantiated. `R1763` is therefore this class's only live construction site. The six rectangles and the pixel map are therefore identical on both screens; what differs is which gates pass. On the mission screen alone `sess+0x3dc == 1`, so rects A, B and F are live and rects D and E are gated off, since `1 & 0x226 == 0`; in a shop opened from a mission `sess+0x3dc == 3` and all six are live, since `3 & 0x226 == 2`. Rect C is live in both, its gate being `!(sess+0x3dc & 0x600)` and `flag == 0`, and `flag` requires `sess+0x3dc == 1` together with a screen taller than 722 pixels.

Because the rects are built from the accumulated origin rather than from `this+0x8`, the rect re-set that `SHOP-FIGURE-041` describes changes the stored rect without changing any panel-relative offset (`evidence/callctx-R0385-addchild.txt`, `listing-R0332-lbuttonup.txt`)

**Confidence.** High (the single-instance conclusion rests on the child-API census with a zero-result pointer-table cross-check, and on all three constructor overloads this class has, not the one the report first enumerated; the gate values follow from `TOWN-352`'s confirmed bit 0/bit 1 semantics)

### TOWN-354

Each of the six rectangles has its own art, and the art for the two picker rectangles is drawn under the same mask that gates their hit test, with one further exclusion the hit test does not share (below), which supplies the evidence `SHOP-PICKER-043` records as missing. `R0389` recomputes the same accumulated rect and the same `flag`, takes the object at `this+0x5c`, and blits through `obj->vt+0x38(x, y, 0, 0, w, h)`. Bindings: rect B takes `[L11964]` at `(L,T)` 28x38 or `[L11965]` at `(L,T+4)` 28x37, under `sess+0x3dc & 3`, chosen by `R0317`; rect D takes `[L11966]` at `(L+1,T+0xcd)` 32x32 and rect E takes `[L11967]` at `(L+0x77,T+0xcd)` 32x32, under `sess+0x3dc & 0x226`, except that both are skipped when `sess+0x3dc & 0x200` and `sess+0x6bc == 2` (`L11968`..`L11969`; `TOWN-355`'s third gate mismatch is this same exclusion); rect A takes `[L11970]` at `(L,T+0xd0)` 32x31 or `[L11971]` at `(L+1,T+0xc9)` 28x30 and rect F takes `[L11972]` at `(L+0x7e,T+0xce)` 32x32, under `!(sess+0x3dc & 0x226)` and `!(sess+0x3dc & 0x400)`, chosen by `R0380`; rect C takes `[L11973]` or `[L11974]` at `(L+0x80,T+4)` 28x32 under `!(sess+0x3dc & 0x600)` and `flag == 0`, chosen by `this+0x70`.

The panel background is `[L11975]` or `[L11976]` at `(L,T)`, chosen by `this+0x70`. The rectangle D and E blits are pixel-exact with the hit rectangles of `TOWN-346`, which are `SHOP-PICKER-043`'s `(1,205,33,237)` and `(119,205,151,237)`; that row's **Medium** clause, that no art was tied to either rect, is answered here. `R0380(v)` is `R0375(v,2) != 0`; `R0317(v)` returns `R0375(sess+0xf0,3) != 0` when `sess+0x3dc & 2` and `R0375(v,3) != 0` otherwise; `R0375(this,id)` walks the child vector at `this+0x1c` and returns the child whose `+0x4` equals `id` (`evidence/listing-R0389-paint.txt`, `listing-R0380-R0317-predicates.txt`, `listing-R0375-findchildbyid.txt`)

**Confidence.** High (every blit's source global, destination and size is an operand of a listing that tiled exactly, and the two picker blits match the hit rectangles exactly)

### TOWN-355

Rect A's and rect F's hit tests are gated on `sess+0x3dc & 1` while their art is gated on `!(sess+0x3dc & 0x226)` and `!(sess+0x3dc & 0x400)`. With `sess+0x3dc == 3`, the value `TOWN-352` establishes for a shop opened from a mission, both rects are live and neither is drawn. Rects D and E keep their hit test under `sess+0x3dc & 0x226` but lose their art when `(sess+0x3dc & 0x200)` and `sess+0x6bc == 2` (`L11968`..`L11969`). Geometrically, rect A `(L, T+0xca, L+0x1c, T+0xf2)` overlaps rect D `(L+1, T+0xcd, L+0x21, T+0xed)` over 27x32 pixels, and rect E `(L+0x77, T+0xcd, L+0x97, T+0xed)` overlaps rect F `(L+0x7e, T+0xce, L+0x9e, T+0xee)` over 25x31 pixels.

`TOWN-346` establishes that no arm of `R0332` returns between the six tests, so a left-button-up inside the first overlap posts `0x40e` twice and `0x414` once, and inside the second posts `0x415` and `0x416`. Rect C is not the only rectangle whose art gate and hit gate are the same expression: rect B's hit gate is a test of the field at offset 0x3dc of the session object against 0x3 (`L11977`) and its art gate is the same test in the paint routine (`L11978`), both `sess+0x3dc & 3`, chosen between by the same `R0317` call the paint routine uses. Rect C additionally has a second route: `R1249` posts `0x412` when its `wParam` is 9 and `!(sess+0x3dc & 0x600)`, and `wParam` of `WM_KEYDOWN` is a virtual-key code, so the key is Tab.

**Falsifiable prediction, on a running original at 640x480:** opening the shop during a mission and clicking `(485,450)` produces both rect A's effect and a change of shown party member in one click, refuted if only one occurs; clicking `(610,450)`, inside the 25x31 overlap that spans x 606 to 630 and y 444 to 474, produces both rect E's member change and rect F's `0x416`, refuted if only the member changes; on the mission screen with nothing else open the same two points produce only A's and F's effects, refuted if the shown member changes; pressing Tab produces the same change as clicking `(608,238)` to `(640,274)`, refuted if the two differ (`evidence/listing-R0389-paint.txt`, `listing-R0332-lbuttonup.txt`, `listing-R1249-keydown.txt`, `listing-R0390-dispatch.txt`)

**Confidence.** High for the gate expressions, the overlap arithmetic and the absence of an early return, all instruction-level; **Medium** for the overlap being reachable in shipped play, which depends on a screen state and a cursor position no static evidence decides, and which the prediction above names the refuter for

### TOWN-356

The loader (at least `L04549`..`L11979`) repeats one block shape: pushing the size 0x24, saving the working pointer in a stack local, storing the allocation result to `[prevGlobal]` and calling `R0747` (`operator new`); on success pushing `<path>`, passing the new object as `this` to `R1176` (bitmap constructor); on failure using a null pointer. The string a block pushes is the constructor argument for the object stored at the top of the NEXT block, not the global this block itself stores into — reading a block's store as paired with its own same-block string reverses every one of the eleven bindings. Corrected and verified against each string's own bytes: `L11975` = `graphics\interface\HumanBackR.bmp` (string `L11980`, pushed `00046a193`, stored `L11981`); `L11982` = `TextBackL.bmp` (`L11983`, excluded — not one of `TOWN-354`'s eleven); `L11976` = `TextBackR.bmp` (`L11984`); `L11964` = `BookOpened.bmp` (`L11985`); `L11965` = `BookClosed.bmp` (`L11986`); `L11970` = `BackPackOp.bmp` (`L11987`); `L11971` = `BackPackCl.bmp` (`L11988`); `L11973` = `HumanMode.bmp` (`L11989`); `L11974` = `TextMode.bmp` (`L11990`); `L11972` = `diskette.bmp` (`L11991`); `L11966` = `ar1.bmp` (`L11992`); `L11967` = `ar2.bmp` (`L11993`); every path carries the `graphics\interface\` prefix.

`R0747` is the standard MSVC `operator new`/`new_handler` retry loop (allocate through `R0722`, retry on the global handler at `L11994`). `R1176` sets the object's vtable to `L11667`, builds a diagnostic string (the string at L11995, `"FATAL ERROR: can..."`) concatenated with the passed path before calls to `R0517` and `R0754`, then allocates a pixel buffer sized `width*height` (a multiplication of the two dimension values) through the same `R0722` allocator. Three otherwise-identical blocks, feeding `L11832`, `L11971` and `L11996` and spaced `0x1c` bytes apart, additionally pass L01181 as `this` to a call to `R0370` before their own allocation; that call's effect was not traced.

The corrected bindings now match every art/gate correspondence `TOWN-347` and `TOWN-354` already document: the `this+0x70`-gated background pair is Human-mode/Text-mode art; rect B's pair is the open/closed journal icon; rect A's pair is the open/closed backpack icon (`TOWN-347`'s child-id-2/3 test); rect C's pair is the Human/Text mode toggle icon; rects D and E, the party-picker rectangles `SHOP-PICKER-043` already names, are `ar1.bmp`/`ar2.bmp`. **The same loader continues past this widget's own resources and reaches `TOWN-093`'s two globals**, `L11600` and `L11598` (`L11601`, `L11599`): under the corrected, address-shifted reading their strings are `L11997` (`extra1024r.bmp`) and `L11998` (`extra800r.bmp`) respectively, which is exactly `TOWN-093`'s own published attribution (`L11600` = `extra1024r.bmp`, `L11598` = `extra800r.bmp`), reached there from the consumer side, not the loader.

This is an independent cross-check of the corrected reading: the naive same-block reading would instead pair `L11600` with `L11999` (`inv1024l.bmp`) and `L11598` with `L06205` (`lm.256`), neither a match (`evidence/listing-L04549-loader.txt`, `bytes-string-literals.txt`, `listing-R0747-opnew.txt`, `listing-R1176-bitmapctor.txt`)

**Confidence.** High (every store address, every pushed string address and every string's own bytes were read at instruction level, and the eleven-way semantic match to `TOWN-347`, `TOWN-354` and `SHOP-PICKER-043` corroborates the corrected mapping independently of the disassembly that produced it)

### TOWN-357

`R0389` at `L12000`..`L09794` loads the table pointer of the drawn actor (`TOWN-349`'s `member`) and calls through slot +0x80 of that table with `this = member` and three stack arguments pushed in order `this+0x78` (`L12001`), `this+0x74` (`L12002`) and a local rect (`L12003`). Both drawable vtables already cited by `HERO-DOLL-078` and `UNIT-PICT-035`, `L02468` and `L02585`, carry `+0x80 = R0745` and terminate at `+0x84`. `R0745` takes exactly these three parameters (returning with 0xc argument bytes cleaned): the rect (`param1`), `this+0x74` (`param2`, loaded `L12004`) and `this+0x78` (`param3`, loaded `L12005`); both surfaces are cleared through a matching pair of block fills (a four-byte-unit fill and a one-byte-unit fill) before anything else, `this+0x74` unconditionally and `this+0x78` only when non-null (`L12006`..`L12007`).

It then runs a 12-iteration loop, index `0..11`, with the loop-back branch to `L04310` at `L04308`, reading `member[index*4+0x15c]` at `L03443` — the same array `TOWN-350` documents from the pickup routine's mouse handler, now reached from a third, unrelated code path, the render compositor. A non-null slot's item pointer is passed twice to `R0886` (`L12008`, `L12009`), the sprite-path resolver `SPR256-DOLL-045` already names. Loop indices 3, 7, 8 and 9 (0-based) each receive one extra helper object (`L04291`..`L12010`); whether these correspond to the same digits `TOWN-350` compares `member+0x74` against was not established, and nothing in this experiment ties a loop index to that unrelated field test (`evidence/dwords-vtables-L02468-L02585.txt`, `listing-R0745-compositor.txt`)

**Confidence.** High for the call site, the parameter mapping, the vtable resolution and the array read, every step a named instruction that tiled; the loop-index/field-value digit overlap is reported as an unresolved coincidence, not a finding

### TOWN-358

Inside `R0745`, after the main 12-slot loop, a scan of the cited evidence file for indirect calls through the dword at `[reg + 0x40]` finds 28 such calls, not four, in two structurally identical clusters: thirteen at `L12011`..`L12012`, and a second run of fifteen at `L12013`..`L12014`. Each call is null-guarded and pushes three zero-filled arguments plus a first argument that is a literal constant, except one call in the first cluster whose four arguments are all zero and two calls in the second cluster whose first argument is `0x1`/`0x2` followed by three copies of a register value rather than zero. The row's original four sites are the first four of the first cluster: first arguments `0x8` (`L12015`), all-zero (`L12016`, from the pair of stack locals at stack offsets +0x20 and +0x24), `0xc` (`L12017`) and `0xa` (`L12018`).

The distinct literal constants observed across all 28 calls: `0x1`, `0x2`, `0x4`, `0x5`, `0x6`, `0x7`, `0x8`, `0x9`, `0xa`, `0xb`, `0xc` — every value from `0x1` to `0xc` except `0x3` — plus the one all-zero call. Immediately before the first cluster, `this+0x78` (reloaded at `L12019` from stack offset +0x8a4) is tested non-null and its own parameterless `vt+0x28` is called (`L12020`); whether the two clusters execute on the same call or are alternative branches of the routine was not traced. The helper objects' class vtable is `L04339`; `+0x28 = R1689`, `+0x40 = R0896`. `this+0x78`'s own class vtable is `L09601` (set by its constructor `R2032`, already cited by `TOWN-349`); its `+0x28` is the SAME `R1689`.

`R1689` computes `this+0x10 + 8` (`TOWN-349`'s own "first pixel" convention) and the object's width/stride through `vt+0x20`/`vt+0x24`/`vt+0x30`, storing them into four fixed globals: `L01168` (buffer base), `L01169`, `L09597`, `L01511` (stride). `R0896` is a thin argument-marshalling thunk: it reads `this+0xc[0]`, extracts three of that struct's fields, and forwards them together with the caller's own constant into a 6-argument call to `R0897`; the constant reaches `R0897`'s stack argument at frame offset +0x1c unchanged, traced through `R0896`'s own push order back to the exact value `R0745` pushed. `R0897` reads the buffer base and stride from `L01168` and `L01511` — the globals `R1689` just filled for `this+0x78` — decodes a run-length opacity byte stream (the stack argument at frame offset +0x18, a count/skip byte tested against `0x40`), and for each byte of an opaque run stores the caller-supplied fill byte, the stack argument at frame offset +0x1c, at `L04344`.

This is the store `TOWN-349` establishes the reader of: a value written by this chain, decremented by the consumer, indexes `member+slot*4+0x15c`. What the constants denote (a fixed body-region id, an engine slot id, or something else) was not established, tested against the full 28-call population rather than the three constants first read: the observed literal set spans `0x1` to `0xc` missing only `0x3`, which is not a subset either candidate hypothesis names, and none of the 28 constants is derived from the main loop's own index (`evidence/listing-R0745-compositor.txt`, `dwords-L04339-layervtable.txt`, `dwords-L09601-surfacevtable.txt`, `listing-R1689-selecttarget.txt`, `listing-R0896-stampthunk.txt`, `listing-R0897-rlestamp.txt`)

**Confidence.** High for the call chain and the literal store instruction, read at instruction level to the byte store; **Medium** for what the fill-byte constants denote, tested against the full 28-call population's constant set (the main loop's own special-cased indices, `TOWN-350`'s field comparison) and not resolved by either

### TOWN-359

`TOWN-350` gives `member`'s array as `member + w*4 + 0x15c`, `w` the routine's own 0-based `slot` parameter; `TOWN-349` gives `w = p-1` for the pixel-map's raw byte `p`. `ITEM-EQUIP-006` gives `actor`'s layout as `s == 1 -> actor+0x74`, `s == 2 -> actor+0x78`, `3 <= s <= 12 -> actor + s*4 + 0x198`, `s` the engine's own 1-based slot number. Two live readings of how `w` relates to `s` were tested, both terms solved the same way, `actor - member` isolated from the equality each field pair requires: `actor + A = member + M` gives `actor - member = M - A`. Reading (a), `w = s` (no decrement): the array term requires `actor - member = -0x3c` (from `s = 3`: `(member+0xc+0x15c) - (actor+0xc+0x198) = 0x168 - 0x1a4`), the scalar term requires `actor - member = 0xec` (from `s = 1`: `(member+0x160) - (actor+0x74) = 0x160 - 0x74`); `-0x3c != 0xec`.

Reading (b), `w = s-1` (`TOWN-349`'s own decrement, so `p` itself is engine-style 1-based): the array term requires `actor - member = -0x40` (from `s = 3`: `(member+0x8+0x15c) - (actor+0xc+0x198) = 0x164 - 0x1a4`), the scalar term requires `actor - member = 0xe8` (from `s = 1`: `(member+0x15c) - (actor+0x74) = 0x15c - 0x74`); `-0x40 != 0xe8`. Reading (b) is the better-supported of the two: it makes `member`'s own `w = 0, 1` (`member+0x15c`, `member+0x160`) the pair `TOWN-350` already names for two-handed weapon displacement, matching `actor`'s own two live fields `+0x74`/`+0x78` for the same two-handed pair. Under both readings the array term requires a negative `actor - member` and the scalar term requires a positive one: no single constant address adjustment supplies both at once.

This refutes "the two arrays are the same memory reached through two constant-offset views of one object" without deciding whether `member` and `actor` are otherwise related (a copy, a lookup, or unrelated objects); that question was not addressed by this experiment. No new disassembly was produced for this row; it is arithmetic over `TOWN-349`, `TOWN-350` and `ITEM-EQUIP-006`'s own already-published field offsets

**Confidence.** High for the arithmetic and for both readings being tested (the conclusion holds regardless of which is intended); Unknown for whether `member` and `actor` name the same underlying object through any relationship other than a constant pointer offset, which this experiment did not investigate

### TOWN-360

A whole-image census of every instruction encoding `disp32 == 0x15c` finds 87 hits; exactly two are an indexed write of the shape of a store of a register to `[reg+reg2*4+0x15c]`: `L04848`, `TOWN-350`'s own drop routine, and `L04320`. `HERO-APPEAR-054` (`claims/hero.md:78`, `EXP-0115`) already publishes `L04320` as one of the same array's four message-store writers; this experiment's addition at that address is the surrounding context given below, not the address's existence as a writer. At `L04320`, the cached held-item pointer (the stack local at frame offset -0x390) is stored at 4 bytes times a slot index plus 0x15c from a base pointer (the stack local at frame offset -0x380), the slot index being the stack local at frame offset -0x384; two instructions earlier the same held item's `+0x18` field is set to `1` (`L09888`).

Immediately after, the routine sets bit 3 in the field at offset 0x18c of the object (`L03218`..`L12021`) and calls `R0551(object)` (`L03213`) — the same bit and the same call `TOWN-350` documents at the end of `R0339`'s own pickup arm (`L11914`). This is the same equip-write convention `TOWN-350` already names, reached from a second call site; what object the frame local at -0x380 holds (`member` itself, an alias of it, or a distinct object) was not traced past this window. The surrounding routine also reads `[[frame local -0x10]+0x3cc]` (the held-item field `TOWN-348`/`TOWN-350` name `sess+0x3cc`) and, on held-item type 1 or 2, plays `[L01211]` through `R0320` (`L12022`) — the same sound and call `TOWN-348` documents for the same held-item types.

This experiment did not trace the frame local at -0x380's own resolution to its origin, nor the containing function's own boundaries or caller (`evidence/field-15c-census.txt`, `ctx-L04320-second-writer.txt`)

**Confidence.** High for the census completeness (a whole-image `disp32` scan, both write sites read at instruction level) and for the shared bit/call convention; Unknown for whether the frame local at -0x380 is `member`, an alias of `member`, or a distinct object, left open for a follow-up experiment

## Room transitions and campaign state bits

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-372 | Seventeen routines in `R1315`..`R2035` are surface transitions. Twelve of them set a cursor, and ten of those set `wait` on entry and `default` on exit. | High / Medium | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |
| TOWN-373 | Fourteen routines set `campaign+0x3dc` bits on entry, one bit each except one that sets two, and the town-screen routine sets none. | High / Medium / Unknown | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |

### TOWN-372

Each loads a registry slot and calls the set-cursor adapter four bytes later. The pattern, from the complete 69-caller population: `R1315` (`wait` `L12023`, `default` `L12024`), `R1316` (`L12025`, `L12026`), `L06738` (`L12027`, `L12028`), `R1165` (`L12029`, `L12030`), `R1317` (`L12031`, `L12032`), `R0816` (`L06215`, then **`select`** at `L06218`, not `default` -- `MENU-CURSOR-046`), `R1318` (`L12033`, `L12034`), `R1319` (`L12035`, `L12036`), `R1320` (`L12037`, `L12038`), `R1321` (`L12039`, `L12040`), `R0909` (`L12041`, `L12042`). `R0361` sets `default` only, at `L12043`.

Five further routines set a `+0x3dc` bit and no cursor at all: `L08084`, `L12044`, `R2036`, `R1658`, `R2035`. Eight of the seventeen identify their surface by a music path the routine itself pushes: `R1315` `music\shop.wav` at `L12045`, `R1317` `music\map.wav` at `L12046`, `R0816` `music\menu.wav` at `L06213`, `R1318` `music\inn.wav` at `L12047`, `R1319` `music\schoolm.wav`/`music\schoolw.wav` at `L12048`/`L12049`, `R1320` `music\Town.wav` at `L12050`, `L08084` `music\inn_ssi.wav` at `L12051`, `R0909` `music\chrgen.wav` at `L12052`, the character-generation screen. A ninth, `R2036`, pushes a resource literal that is not a music path: `graphics\interface\logo\allods.bmp` at `L12053`.

Seven of these push sites carry an identical ten-byte instruction lead-in and all seven push a music path; the four `school` sites and the `allods.bmp` site do not carry it. **`R0909` was missing from the first draft of this row's four counts**, because `evidence/regen.sh`'s per-routine literal scan used a range that stopped at `L12054` while the routine runs to `R2036`, and the whole-image string scan used a block that does not contain `music\chrgen.wav`. A seventeenth-family candidate, `R0099`, sets `wait` (`L12055`) and a second cursor (`L12056`) with no `+0x3dc` bit and no identifying literal; it has the shape of `R1320` and is **not** counted here, because the family's membership criterion is not stated precisely enough to settle it

**Confidence.** High for the cursor pattern: every load and every adapter call is a named instruction, each load exactly four bytes before its call, and the 69 callers are a complete population (`AI-CURSOR-175`) -- which is the instrument that caught the missing routine. High for the nine push sites, which were derived from the opposite direction -- a whole-image dword scan giving each string's own push site -- rather than from the `PUSH imm32` literal heuristic, which produces false positives. The music paths are not contiguous and no single range holds them: eight lie in `L11961`..`L12057`, a cluster that also holds five strings that are not music paths, and `music\chrgen.wav` lies at `L12058` with `allods.bmp` at `L12059`. Both clusters are scanned in `evidence/surface-listings.txt`; the per-routine literal scan over each routine's own byte range is the primary instrument and the cluster scans are the cross-check. **Medium** that a routine pushing a surface's music path is that surface's transition: the alternative that it is a shared audio helper is not excluded by the literal alone, though the one-cursor-pair-per-routine shape argues against it

### TOWN-373

The fourteen are enumerated in full below; an earlier draft listed thirteen and omitted `R0909`. `SESS-SCREEN-003` establishes `campaign+0x3dc` as a bitmask and names bit 0, marking bits 1, 2, 9, 10 and 13 Unknown individually. `TOWN-023` records that a generic `disp:3dc` sweep is inconclusive and that a sweep scoped to the campaign screen's own range was not run. This is that sweep, run whole-image with the base register retained: 57 stores to `[base+0x3dc]`, of which fourteen are a single bitwise or at the head of a surface-transition routine (`TOWN-372`). Bit and site: `R1315` sets mask 0x2 at `L12060` (bit 1, shop); `R1316` sets mask 0x4 in the second byte at `L12061` (bit 10); `L06738` sets mask 0x10 in the second byte at `L12062` (bit 12); `R1165` sets mask 0x40 in the second byte at `L12063` (bit 14); `R0361` sets mask 0x80 in the second byte and mask 0x8 in the low byte at `L03330` (bits 15 and 3); `R1317` sets mask 0x10 at `L12064` (bit 4, map screen); `R0816` sets mask 0x80 in the low byte at `L06214` (bit 7, main menu); `R1318` sets mask 0x4 at `L12065` (bit 2, inn); `R1319` sets mask 0x20 at `L12066` (bit 5, school); `L08084` sets mask 0x40 at `L12067` (bit 6, second inn); `L12044` sets mask 0x1 in the second byte at `L12068` (bit 8); `R1321` sets mask 0x2 in the second byte at `L12069` (bit 9); `R0909` sets mask 0x8 in the second byte at `L12070` (bit 11, character generation -- `music\chrgen.wav`, pushed at `L12052` in the same ten-byte call idiom as the shop, map, menu, inn, Town and second-inn music loads); `R2036` sets mask 0x2000 at `L12071` (bit 13).

`R1658` and `R2035` each set and clear bit 13 within themselves. **`R1320`, the routine that loads the town screen's own music, contains no store to `+0x3dc` at all**, which agrees with `SHOP-TOWN-023` reading value 0 as the campaign's home state

**Confidence.** High for the sweep and the per-routine table: the read-modify-write pairs are named instructions with the actual bitwise operation and the register width recorded, and the absence of any `+0x3dc` store in `R1320` is a scan of that routine's whole byte range for the displacement, not an inference. **Medium** for reading a bit as *the named surface is displayed*: each naming rests on a resource literal in the same routine, and the live alternative that the bit means *this resource set is loaded* is excluded only for **bit 1**; bit 11's naming rests on the same literal evidence as the rest. Where `TOWN-350` gates the character figure widget's slot pick-up on `sess+0x3dc & 2` or `sess+0x3dc == 1` -- a resource-loaded flag does not gate an inventory interaction. Bits 3, 8, 10, 12, 14 and 15 are set by routines with no identifying literal and are **Unknown**

## School diamond lifecycle

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-379 | The measured Train action arms the school diamond before constructing its purchase request, after local price/purse guards. | High / Unknown | ● active | [EXP-0265](../experiments/EXP-0265-school-diamond-lifecycle/), `raw-anchors.tsv`, `ghidra-ranges.txt`, `raw-call-pointer-census.tsv` |
| TOWN-380 | The recovered diamond updater advances once per school own-surface paint invocation while its step is nonzero; the nearby column clock does not gate it. | High / Unknown | ● active | [EXP-0265](../experiments/EXP-0265-school-diamond-lifecycle/), `raw-anchors.tsv`, `ghidra-ranges.txt`, `raw-call-pointer-census.tsv` |
| TOWN-381 | The diamond's retrigger preserves phase, forces step `+1`, and completion caches frame zero while step zero hides it. | High / Unknown | ● active | [EXP-0265](../experiments/EXP-0265-school-diamond-lifecycle/), `raw-anchors.tsv`, `ghidra-ranges.txt`, `paint-sequence.tsv`, `retrigger-model.tsv` |
| TOWN-382 | School entry resets diamond state; picker bodies do not directly reset or rearm it, while leave releases frames without an explicit scalar-control clear. | High / Medium / Unknown | ● active | [EXP-0265](../experiments/EXP-0265-school-diamond-lifecycle/), `event-range-census.tsv`, `ghidra-field-occurrences.tsv`, `ghidra-ranges.txt`, `raw-call-pointer-census.tsv` |

### TOWN-379

In `R1548`, `L10204` branches past the arm when room `+0x358` (purse minus price) is negative; `L10205` does so when `+0x354` (price) is nonpositive. `L12072` loads 1 and `L12073` stores it to room step `+0x264`. Only afterwards does `L12074` call `R1932`, whose `L12075` writes opcode `0x3d`. Acceptance of this new request is therefore not a precondition for this arm. Local refusal bypasses the arm, rather than cancelling an already active animation. This narrows “attempted Train” to a locally admitted action; no server-declined transaction or later response effect was witnessed. Whole-`.text` E8 and `.rdata` pointer scans reproduce the builder call and Train handler slot `L12076`; literal field scans do not establish that this is every possible dynamic starter.

**Confidence.** High for the named instruction order and refusal branches, independently decoded from raw PE and fresh repaired Ghidra; Unknown for later reply effects and a universal no-other-starter claim

### TOWN-380

`R1487` calls `R1492` at `L12077`; all branches of the exported paint prefix converge on that call. The complete updater reads no clock and calls no helper. The whole-`.text` raw E8 census finds exactly this one updater candidate, decoded as a call; the whole-`.rdata` dword census finds no updater pointer. The saved-clock subtraction `L12078` and comparison with `0x53` at `L12079` occur after the diamond call on the separate column route. This establishes an invocation rule, not an 83-ms diamond period or a measured duration. It does not exclude state writes through unmeasured aliases, overlapping-width stores, bulk copies or dynamic indirect targets.

**Confidence.** High for the complete updater, paint-prefix control flow and exact searched-section counts; Unknown for actual paint cadence, visible duration and mutations outside those populations

### TOWN-381

`R1492` returns unchanged for step zero; otherwise it adds step to phase `+0x260`. A signed result at least 8 is clamped to 8 with step `-1`; a result exactly zero sets step zero and temporarily clears current `+0x25c`. The common tail always stores `frames[phase]` from array `+0x24c` into current at `L07792`, including after that clear. Paint requires both current and step nonzero. From initialized idle after an arm, the tested static relation gives phases `1,2,3,4,5,6,7,8,7,6,5,4,3,2,1,0`: 16 updater transitions, 15 draw-eligible states if frames are nonnull. Train writes only step `+1`, with no activity check or phase/current reset in the arm: ascent continues, interior descent reverses, and a retrigger at phase 8 produces one extra top state before descent. These are state projections, not observed displayed frames.

**Confidence.** High for the raw/Ghidra instruction relation and the phase `0..8` retrigger enumeration; Unknown for malformed states, load failures and which eligible states reach the display

### TOWN-382

Entry `R0706` calls loader `R1486` at `L12080`; the loader clears current at `L12081`, then entry writes phase and step zero at `L12082`/`L12083`. Re-entry through this route resets regardless of object reuse. Complete picker bodies `R1959` and `R1958` contain no diamond-field literal or updater call; the price-reply body `R1956` writes only `+0x354`. Leave calls release `R1969` at `L12084`, and the destructor calls it at `L12085`; release destroys frame entries and reduces the array count, without a direct literal write to current, phase or step. Constructor/initializer ranges initialize the frame-array header but contain no control-field literal.

These body-level facts do not prove complete event preservation through callees: picker-triggered painting may still advance phase, and dynamic aliases/partial-width writes were not exhausted.

**Confidence.** High for named entry clears, release calls and bounded body scans with raw/Ghidra agreement; Medium for whole-picker-event preservation; Unknown for actual object reuse, indirect event effects and unmeasured writer encodings

## School and tavern captions

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-383 | The two school-widget caption sources are fixed global `main.txt` indices 231 and 232: element 0 is `Train` / `Учиться`, element 1 is `Exit` / `Выход`, for the measured EN / RU roots. | High | ● active | [EXP-0267](../experiments/EXP-0267-school-captions/), `source-loads.tsv`, `source-bindings.tsv`, `main-table.tsv`, `school-listing.txt` |
| TOWN-384 | The school captions use owned string copies rather than borrowing the global loader's character pointers. | High / Unknown | ● active | [EXP-0267](../experiments/EXP-0267-school-captions/), `school-listing.txt`, `storage-helpers.txt`, `school-fields.tsv`, `program-digest.tsv` |
| TOWN-385 | The school's pressed/unpressed paint branches select presentation, not different caption sources, and no caption refresh was found on the directly inspected widget event and entry/leave-helper routes. | High / Medium / Unknown | ● active | [EXP-0267](../experiments/EXP-0267-school-captions/), `school-listing.txt`, `room-routes.txt`, `school-fields.tsv`, `constructor-refs.txt` |
| TOWN-391 | The tavern button panel constructs three resource-derived captions, then a selection writer replaces the first. | High / Unknown | ● active | [EXP-0268](../experiments/EXP-0268-tavern-captions/), `source-loads.tsv`, `caption-lines.tsv`, `panel-range.txt`, `ghidra-functions.txt` |
| TOWN-392 | Tavern captions and numeric strings are separate paint arguments; pressed appearance does not select another caption. | High / Medium | ● active | [EXP-0268](../experiments/EXP-0268-tavern-captions/), `panel-range.txt`, `ghidra-functions.txt`, `ghidra-refs.txt`, `raw-references.tsv` |
| TOWN-393 | Tavern caption assignments copy character bytes into owned string storage; the panel destructor releases that storage. | High / Unknown | ● active | [EXP-0268](../experiments/EXP-0268-tavern-captions/), `panel-range.txt`, `ghidra-functions.txt` |

### TOWN-383

Both constructors `R2019` and `L12086` read `[L04369]+0x39c` and `+0x3a0` and pass those lines to the two successive four-byte string entries addressed by `[widget+0x64]`. The independent PE probe decodes all four source-load displacements separately in both identical executable inputs; it passes those measured indices to the existing text probe. TEXT-STRTAB-023 supplies the established main.txt base of zero. No actor, class, skill, price or session selector participates in either constructor's caption lookup. This resolves TOWN-280's caption-source Unknown, not its separate per-paint value lines.

**Confidence.** High for both constructor bindings and separately measured EN/RU line data; no running-original glyph or caption observation

### TOWN-384

At `widget+0x60`, the constructor initializes an array header and sizes it to two four-byte string entries via `L12087` and `R0728`; its backing pointer is `widget+0x64`. Each `R1774` assignment reaches `R2037`, which prepares storage through `R2038`, copies source bytes through `R0544`, writes length and adds NUL. The allocation path `R2039` stores a newly allocated buffer pointer; the shared empty-string path is separate. Destructor `L12088` invokes `R2040` on this embedded array at `L12089`; it releases the counted strings through `R2041`/`R1687`, then frees the array. This establishes constructor-time copies and the destruction route, not a runtime destruction time or proof of immutable text.

**Confidence.** High for the complete named allocation/copy/release instruction paths in the digest-matched executable; Unknown for allocation failure, runtime object reuse and later writes outside the searched routes

### TOWN-385

`R1905` reads `[widget+0x64]+(offset-0x360)` in both branches with offset `0x360,0x364`; caption draws are `L07784` and `L07786`. The pressed branch requires nonnegative `+0xb0`, equality with `+0xac`, and equality with the loop element. Entry `R0706` calls widget `R1973`, `L11808`, `R1977`; leave `R1915` calls `R1976`, `R2042`. Those complete helper bodies and widget pointer-event bodies do not assign the caption array. A decoded-instruction census over `[L12090,R1218)` finds six widget `+0x64` accesses (four constructor receivers, two painter reads) and three `+0x60` address formations (two constructors, destructor); its 5,870 decoded instructions leave 1,077 bytes unmatched. The census and decoded call/aligned-rdata reference searches do not exclude alias-based, bulk, overlapping-width, undisassembled or external writes, indirect calls, or effects in the unexhausted event call graph.

**Confidence.** High for the named direct branches and bounded search counts; Medium for preservation across complete selection/room events; Unknown for runtime object reuse, hot reload and actual visible text

### TOWN-391

`R1809`, called from both `R1534` and parameterized `R1535`, sizes the string array at panel `+0x60` to three. Its backing element pointer is `+0x64`. Assignments at `L12091`, `L12092`, `L12093` copy global slots 243, 242, 232 into elements 0, 1, 2. The parent directly constructs the parameterized panel at `L12094`; neither direct-call nor raw `.rdata` pointer census finds the default constructor. `R0600` replaces element 0: selected index `[parent+0xb8]` greater than signed `[parent+0xc8]-1` passes zero-initialized address `L12095`; otherwise it reads selected object type byte `+0x15b` and the working hire vector `[[application+0x5d0]+4*(type-1)]`, choosing slot 259 if nonzero, 258 if zero.

It has no negative-index guard. All five slots are in `main.res::text/main.txt` under `TEXT-STRTAB-023`; raw PE load-pair extraction agrees with Ghidra, and fresh EN/RU archive reads find 274 lines per root, five differing nonempty source lines per locale, no percent byte in those ten lines. Identical executables use identical slots, not locale-specific code branches.

**Confidence.** High for the two writer bodies, five source operands, destination elements and resource measurements; the type/hire meaning uses `MERC-TYPE-001` and `MERC-HIRE-003`. Unknown for malformed selection, later arbitrary mutation of the initially empty source, and untraced computed alias writers

### TOWN-392

Painter `R1904` returns when `[parent+0x13c]` is zero. It draws caption element 1 at the middle button, refreshes the parent's numeric strings via `R1556`, then loops elements 0 and 2 for the upper and lower buttons. Each caption reaches font global `L06186`, virtual slot `+0x14`, alignment `0xa`. Only a nonzero caption length permits the corresponding parent numeric string at `+0x124` or `+0x12c` to draw below it; the middle button has no numeric draw in this routine. Both pressed/unpressed arms read the same element; panel `+0xc4`/`+0xc0` comparisons affect appearance and one-pixel vertical displacement, not the source index.

The activation body `R1412` calls the first-caption selector at `L08097` before setting the active paint gate at `L12096`. Seven direct selector calls in six functions are reproduced by raw `E8` and Ghidra censuses, with no raw `.rdata` pointer to the selector. The constructor's combined first-caption source is therefore replaced on that normal activation route.

**Confidence.** High for these body-level arguments, branch conditions and activation order; Medium for event-wide completeness because computed targets, indirect alias writes and actual displayed frames were not observed

### TOWN-393

Both panel constructors initialize the string-array object at `+0x60`. The common initializer and first-caption selector pass source pointers to `R1774`, which takes `lstrlenA` and calls `R2037`. That helper ensures writable capacity through `R2038`, copies through `R0544`, writes the string length at buffer `-8` and a trailing NUL. Shared or insufficient buffers are replaced by the capacity path through `R2039`; unchanged capacity may reuse a buffer. Thus a source pointer is not retained as a borrowed caption buffer, and a cached character pointer need not survive a later assignment. Panel destructor `R2043` calls `R2040` on `+0x60`; it releases every string through `R2041`/`R1687` and frees the element array. String release decrements the reference count and frees a non-sentinel buffer when it reaches zero.

**Confidence.** High for the assignment/copy and destructor mechanisms; Unknown for the actual room-exit destruction schedule, allocation failure, all external aliases, locale hot reload and live pointer lifetimes

## Town door label gate correction

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-453 | Correction to `TOWN-161`: the shop arm's comparison, `L07768`, is a register compare against `EBX`, not the immediate `1`. | High / Medium | ● active | [EXP-0306](../experiments/EXP-0306-town161-operand/), `static-measurements.tsv`, `image-range-digests.tsv` |

### TOWN-453

The bytes `399eb4000000` decode as a comparison of the 32-bit field at offset 0xb4 of the object with a held value (opcode `0x39`, ModRM `0x9e`), not an immediate form. The held value is established as `1` once, at `L07763` (loaded as the constant 1), inside the same routine (`R1489`), and is not rewritten between that write and the comparison. The routine's complete body — all 438 instructions from entry to its single return at `L07766` — was read: after the shop comparison, the held value is rewritten three more times (`L07765`, `L12097`, `L12098`, each a load of the dword at offset 0x8 of the object), feeding the `+0xf4`/`+0x120`/`+0x150` door blits `TOWN-159` already documents, none of which can affect a comparison that already executed.

This experiment reads 16 functions across 11 disassembled address ranges in the direct-call closure from routine entry to `L07768` — levels one and two of the experiment's own any-depth reduction rule, not its full extent: the Town_add gate (`R1917`), the reroll (`R1907`), one 67 ms-hub callee (`R1927`) and the shared blit primitive (`R0797`, `TOWN-150`'s vt`+0x18` target) each bracket every `EBX` use with a matching entry-save/exit-restore; the PRNG core and its nested call, the bird-arm routine and its two sound helpers, the composite-lock/present family and the blit wrapper reference `EBX` nowhere in their own bodies. Deeper levels of the closure were not enumerated or read.

Five call sites across three targets in that closure are not resolved here: three calls against the imported `timeGetTime` (`L12099`, `L12100`, `L11690`; the function pointer loaded at `L12101` from IAT slot `L00849`, `SESS-TIMER-022`), an indirect call inside the composite-lock routine, which `AI-CURSOR-225` (already published) identifies as the DirectDraw back-surface Lock — an external system-DLL method trusted by the standard callee-saved convention, not independently disassembled; and a conditional vtable dispatch on the child object at offset 0x68 of the object whose own vtable identity was not traced to a constructor. School (`L07770`, `0x4`) and tavern (`L07771`, `0x2`) remain genuine immediates, unchanged.

`TOWN-161`'s stated door-label values — shop`1`, school`4`, tavern`2` — are unaffected; only the shop arm's mechanism wording is corrected. Fresh Ghidra disassembly and an independent `tools/pebytes` dual-root digest pin agree with each other and with the digests already committed in `EXP-0299`'s own `anchors.tsv` (`en==ru` in both).

**Confidence.** High for the register-vs-immediate byte decode; High for the establishing write and its persistence between `L07763` and `L07768`, over the routine's now-completely-read 438-instruction body; Medium for the population read, bounded to 16 functions across the closure's first two levels, not its full any-depth extent; Medium for "no other value is ever possible", bounded by five call sites across three named unresolved targets

## Option controls and consumers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-OPTIONS-457 | Game Options message0445 exports its controls before forwarding the close. | High / Medium | ✔ promoted | [EXP-0366](../experiments/EXP-0366-option-consumers/), derived instructions, controlled measurements and two branch mutations |
| TOWN-AUTOHEAL-458 | The recovered autohealing command maps modes0/1/2 to player mana-floor percentages100/50/0. | High / Medium | ✔ promoted | [EXP-0366](../experiments/EXP-0366-option-consumers/), derived instructions, controlled measurements and two branch mutations |
| TOWN-GRAPHICS-459 | In the measured unit dispatch, Shadows0 suppresses the shadow pass while preserving the body pass. | High / Medium | ✔ promoted | [EXP-0366](../experiments/EXP-0366-option-consumers/), derived instructions, controlled measurements and two branch mutations |
| TOWN-SMOOTH-460 | Smoothing gates the additional boundary-sprite draw after the base draw in the recovered backpack sprite painterR1795. | High / Medium | ✔ promoted | [EXP-0366](../experiments/EXP-0366-option-consumers/), derived instructions, controlled measurements and two branch mutations |

### TOWN-OPTIONS-457

Smoothing, Shadows, Lighting and Animation reach their separate stored flags; child014 exports the global AutoCasting mode and emits the autohealing command. Message0446 alone leaves the tested values unchanged. On OK, Animation0 forces Lighting0.

**Confidence.** High for the bounded original message arms, exports and flag coupling; Medium for dialog membership interpretation. Native registry lifetime and physical UI input remain Unknown.

### TOWN-AUTOHEAL-458

Direct values3..100 are percentages; negative or greater than 100 retain the preceding value. The actor fold truncates signed maximum mana times percentage divided by 100 into its signed-word floor. The heal AI requires nonzero book/mana and current mana strictly greater than this floor before looking up spell6. Equality does not pass. This joins the global option to SAV-PLAYER-028 and HERO-MP-006; selected-spell autocast is a separate surface.

**Confidence.** High for the command, arithmetic and strict lookup boundary, with a distinguishing mutation; Medium for the joined UI interpretation. Complete target selection, battle eligibility, successful casting and native scheduler delivery remain Unknown.

### TOWN-GRAPHICS-459

In the named scenery paint, Animation0 selects frame0 instead of the computed frame. In the named dynamic-light method, Lighting0 exits before the light body; its flag differs from DayNight. Game Options couples Animation0 to Lighting0 on commit. These results cover the named unit/scenery/light paths, not every drawable or animation family.

**Confidence.** High for the discriminated original branches and animation mutation; Medium for family interpretation. Water, town, UI and complete moving-actor option populations and native cadence remain Unknown.

### TOWN-SMOOTH-460

LoaderR1057 assigns graphics/backpack/sprites.256 and spritesb.256 to its respective globalsL10109/L12102 throughR0753; the boundary method reaches the half-colour painter. For an opaque indexed boundary pixel that painter computes ((destination16>>1)&mask)+((palette16[index]>>1)&mask), with mask0x7bef for RGB565 or0x3def for RGB555; tested left-clipped pixels remain untouched. This joins the Smoothing option to the overlay contract in SPR256-OVL-014.

**Confidence.** High for the option gate and synthetic original pixel execution; Medium for the static loader/dispatch join. Complete family coverage, full native screen equivalence and temporal filtering elsewhere remain Unknown.

## Town figure clicks and guard state

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-473 | In the town view slot `+0x54` (`R0704`) decides a left click by the picture `townmask.bmp` alone and tests no figure; while the tip popup is shown (`TOWN-185`), the popup and its controls are offered the click first. | High / Medium | ● active | [EXP-0409](../experiments/EXP-0409-town-figures/EXP-0409.md), `evidence/d-town-click.txt`, `evidence/d-capstone.txt`, `evidence/mask-histogram.tsv`, `evidence/mask-regions.tsv`, `evidence/sampler-table.tsv`; extends TOWN-163, TOWN-211, TOWN-406 |
| TOWN-474 | Ten painted town figures lie on one click region each: tavern `0x80`, shop `0x90`, gate and guards `0xa0`, statue star `0xb0`, school `0xc0`; horses 2 and 4 reach two regions with 7 and at most 167 pixels; the other figures reach none. | High | ● active | [EXP-0409](../experiments/EXP-0409-town-figures/EXP-0409.md), `evidence/overlap-summary.tsv`, `evidence/overlap.tsv`, `evidence/figures.tsv`, `evidence/exe-anchors.tsv` |
| TOWN-475 | Each town click region posts one message that opens a screen: tavern `0x42b`, shop `0x42a`, school `0x42c`, gate `0x42d`, statue `0x41f`; only the gate depends on quest state, and answers with the `npc35` line when nothing is on offer. | High / Medium | ● active | [EXP-0409](../experiments/EXP-0409-town-figures/EXP-0409.md), `evidence/click-arms.tsv`, `evidence/campaign-arms.tsv`, `evidence/r-campaign.txt`, `evidence/node-families.tsv`, `evidence/nodes.tsv`, `evidence/d-capstone.txt`; extends TOWN-163, TOWN-164 |
| TOWN-476 | A delivered pointer message starts the guards' halberd sweep and the paint hub steps it one frame per admitted hub: the gate with nothing on offer sweeps frames 7 to 0, other messages sweep back; no clock or return trigger was found. | High / Medium | ● active | [EXP-0409](../experiments/EXP-0409-town-figures/EXP-0409.md), `evidence/d-town-guards.txt`, `evidence/model-vectors.tsv`, `evidence/gate-step-domain.tsv`, `evidence/refs-reduction.tsv`, `evidence/refs-cluster-hits.tsv`, `evidence/refs-outside-writers.tsv`; extends TOWN-399, TOWN-403, TOWN-404 |
| TOWN-477 | A room click leaves the town view through message `0x445`, which frees the guard sheet and sounds but writes no frame, step or latch; the return runs enter `R1383`, which sets frame 7, step 0, latch 0 and reloads the sheet. | High / Medium | ● active | [EXP-0409](../experiments/EXP-0409-town-figures/EXP-0409.md), `evidence/d-town-guards.txt`, `evidence/d-campaign.txt`, `evidence/model-vectors.tsv`, `evidence/refs-cluster-hits.tsv`; extends TOWN-164, TOWN-403, TOWN-412 |
| TOWN-478 | In the shop and school panels read, the painted merchant, trainers, diamond and column frame have no click test; the shop panel tests four shelf and two prompt rectangles, the school two class masks in two states. | High / Medium | ● active | [EXP-0409](../experiments/EXP-0409-town-figures/EXP-0409.md), `evidence/d-shop.txt`, `evidence/d-school.txt`, `evidence/d-capstone.txt`, `evidence/rooms.tsv`, `evidence/school-masks.tsv`, `evidence/refs.txt`; extends SHOP-SHELF-047, TOWN-154, TOWN-428 |
| TOWN-479 | Without the tip popup, merchant down is inert only with prompts disabled or missed, and up acts with a held item; school figure down acts only in class panels, and 608 fighter pixels consume up without a button action. | High / Medium | ● active | [EXP-0412](../experiments/EXP-0412-town-square/EXP-0412.md), `evidence/figure-presses.tsv`, `evidence/windows.tsv`, `evidence/d-shop.txt`, `evidence/d-school.txt`, `evidence/d-dispatch.txt`, `evidence/slot-table.tsv`, `evidence/exe-anchors.tsv`; extends TOWN-478, TOWN-473, TOWN-348 |
| TOWN-480 | The four room tip popups pass list/body presses below, consume check box and Close presses, write the tips flag on check box down, and route Close release through `0x45a`; delivery to a room is conditional. | High / Medium | ● active (amended) | [EXP-0412](../experiments/EXP-0412-town-square/EXP-0412.md), `evidence/popup-controls.tsv`, `evidence/popup-presses.tsv`, `evidence/windows.tsv`, `evidence/d-popup.txt`, `evidence/d-dispatch.txt`, `evidence/d-town.txt`, `evidence/d-tavern.txt`, `evidence/d-shop.txt`, `evidence/d-school.txt`, `evidence/exe-anchors.tsv`; extends TOWN-185, TOWN-186, TOWN-207, SHOP-TIP-045 |
| TOWN-481 | Town enter calls a synchronous paint before queued pointer dispatch; an unblocked paint can run the hub when its timer is due. Static room-close routes do not establish which pointer messages the OS generates next. | High / Medium | ● active | [EXP-0412](../experiments/EXP-0412-town-square/EXP-0412.md), `evidence/d-town.txt`, `evidence/d-campaign.txt`, `evidence/r-campaign.txt`, `evidence/guard-steps.tsv`, `evidence/api-summary.tsv`, `evidence/refs-api.txt`, `evidence/code-anchors.tsv`; extends TOWN-476, TOWN-477 |
| TOWN-482 | The 96 residual stores have 33 owners and no direct-call reach from the measured town closure; indirect attribution remains open. Guard frames 0 and 1 are crossed, and sound overlap with the closing endpoint is timing-dependent. | High / Medium | ● active | [EXP-0412](../experiments/EXP-0412-town-square/EXP-0412.md), `evidence/guard-stores.tsv`, `evidence/guard-store-owners.tsv`, `evidence/guard-store-summary.tsv`, `evidence/guard-cones.tsv`, `evidence/guard-frames.tsv`, `evidence/guard-steps.tsv`, `evidence/guard-timing.tsv`, `evidence/wave-headers.tsv`; extends TOWN-476, TOWN-477 |
| TOWN-483 | Gate helper R1416 returns -1 when neither the main record nor a child is latched. Hearing alone does not latch; accepted children can survive a main win, and the last-mission arm posts 0x428 before town entry. | High / Medium | ● active | [EXP-0412](../experiments/EXP-0412-town-square/EXP-0412.md), `evidence/r-record.txt`, `evidence/r-campaign.txt`, `evidence/d-campaign.txt`, `evidence/code-anchors.tsv`; extends REG-SCN-062, REG-SCN-063, REG-SCN-064, REG-SCN-065, SAV-CAMPAIGN-076, SAV-CAMPAIGN-080, SAV-CAMPAIGN-081, SAV-CAMPAIGN-082 |
| TOWN-484 | The npc35 fallback builds a non-modal dialogue whose pager owns voice start/stop. Only its button blocks a left press in the measured screen; a missing text read takes the exception route rather than building a substitute panel. | High / Medium | ● active | [EXP-0412](../experiments/EXP-0412-town-square/EXP-0412.md), `evidence/d-dialogue.txt`, `evidence/dialogue-presses.tsv`, `evidence/dialogue-summary.tsv`, `evidence/d-town.txt`, `evidence/d-exception.txt`, `evidence/eh-frames.tsv`, `evidence/wave-headers.tsv`, `evidence/code-anchors.tsv`; extends DLG-KEYS-040, DLG-ENTRY-016, DLG-DRAW-015 |

### TOWN-473

Town view vtable `L11548` slot `+0x54` is `R0704`. Slots `+0x58`, `+0x5c`, `+0x60` and `+0x64` are the shared stubs that return zero with 0xc argument bytes cleaned `R1812`, `L12103`, `L12104` and `R1889`, and `+0x68` is the framed stub `R1890`, which also returns 0 (`evidence/exe-anchors.tsv`, `evidence/d-capstone.txt`). Slot `+0x4c` is `L11549`, which forwards the position to the pointer handler `L11546` and returns 0 (`TOWN-399`). The base dispatcher `R0390` maps messages `0x201`..`0x206` to slots `+0x54`..`+0x68` (`TOWN-211`, `TOWN-406`), so `R0704` is the view's only left-click routine and the right button reaches stubs. It calls `R1914` with the point (`L12105`) and dispatches on the result through the two tables `TOWN-163` reads. Every path returns 1, a click on a pixel with no action included (`L11721` to `L11719`), and the routine holds no other rectangle, sprite bound or per-figure test.

`R1914` tests the point against the view rectangle at `+8` with `PtInRect` (`L12106`), then reads one byte at `[[view+0x6c]+0x10] + 8 + 640*(y-top) + (x-left)` (`L12107`..`L11714`), subtracts `0x80` and returns -1 when the result is above `0x40` (`L12108`..`L11715`). The `+8` is a header. The loader `R1152` built the object at `view+0x6c` from the literal `graphics\interface\town\townmask.bmp` (`L12109`..`L12110`). It allocates `w*h+0x408` bytes, writes the width and height dwords at the start, reads the `w*h` pixel bytes to `+8` and puts the palette after them (`L12111`, `L12112`..`L12113`), and `R1970` then exchanges rows `i` and `h-1-i` (`TOWN-153`). Buffer row 0 is the picture's top row, so the byte read is picture pixel (x-left, y-top) with no shift.

The node is one file of 308280 bytes, 640x480 at 8 bits, sha256 `fc7aea72cf08ad269f0bf8115715f37afddcfc231abcf2740881439d2e2ba859`, identical on both roots (`evidence/input-nodes.tsv`). Its pixels are 266003 zero, 39771 with one of the five action values and 1426 with another nonzero value, which selects nothing (`evidence/mask-histogram.tsv`, identical on both roots). The byte table at `L11716` and the dword table at `L11717` return selector 2 for `0x80`, 1 for `0x90`, 8 for `0xa0`, 16 for `0xb0` and 4 for `0xc0`, and -1 for the other 251 values (`evidence/sampler-table.tsv`). Regions, as inclusive boxes (`evidence/mask-regions.tsv`, identical on both roots): `0x80` 5764 pixels in 15 components, the largest 5750 pixels at (127,325)-(213,412); `0x90` 5771 pixels in 5, largest 5767 at (254,261)-(334,347); `0xa0` 7911 pixels in 17, largest 7895 at (156,92)-(245,220); `0xb0` 5592 pixels in one component at (342,289)-(406,419); `0xc0` 14733 pixels in one component at (418,299)-(583,425). Every other component is a single pixel, 14, 4 and 16 of them for `0x80`, `0x90` and `0xa0`, and they stretch the all-pixel boxes to (127,93)-(399,412), (254,261)-(396,394) and (156,92)-(569,423).

The base dispatcher offers a mouse message to the children before it calls the slot: `R0642` runs `PtInRect` on each child's rectangle for messages `0x200`..`0x206` and `0x400`, calls the first child that contains the point through `vt+0x48` and stops after it, whatever it returns (`evidence/d-capstone.txt`, `TOWN-139`). The town view's children read in its constructor `R1942` and its enter `R1383` are the button child with an all-zero rectangle (`TOWN-089`), which contains no point, and the tip popup (id `0x467`, rectangle (328,0)-(640,200), created at `L11740`..`L12114` while the global `L03631` is nonzero and deleted by leave `R1912`). The popup's rectangle holds 0 action pixels on both roots (`evidence/mask-regions.tsv`, column `tip_rect_pixels`). Its three controls (`TOWN-185`) are offered a click inside them before the slot runs; what they do with it was not read here (`TOWN-207`). The hover text for the five regions is bound by `R0739` (`TEXT-HOVERROOM-051`).

**Confidence.** High for the slot words, the stub bodies, the complete body of `R0704`, the sampler, the loader layout and the mask node with its pixel census. Each was read at instruction level or measured over the shipped picture on both roots, and the EN and RU `rom.exe` are one program (sha256 `942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`, `evidence/rom-program-digest.tsv`), so their agreement adds nothing for a code fact. Medium that no other routine acts on a town click: the popup's controls were not read for mouse messages, children were enumerated from the constructor and enter routines only, the per-instance handler at `this+0x34`, which `TOWN-406` says runs before the children, was not enumerated for this view, and a child added by code outside those routines was not searched.

**Unknown.** Whether physical pointer messages reach `R0704` in every state (capture, focus loss), what the tip popup's three controls do with a left click (`TOWN-207`), the writers of the town view's `+0x34`, and children added outside the town's own routines.

### TOWN-474

Each figure's paint anchor is the immediate of the `ADD` pair in the town paint `R1489` (`evidence/exe-anchors.tsv`); the birds and the `town_add` overlay start at the view origin (`R1917`). The join counts the pixels each figure can cover on each mask code. A `.16a` figure counts the union of the opaque pixels of all its frames, and a `.bmp` figure counts its whole rectangle, so its count is an upper bound on the visible pixels (`evidence/overlap-summary.tsv`, `evidence/overlap.tsv`, identical on both roots).

Overlapping one region: the tavern door `townbirds/tavern` (10 frames of 64x96 at (124,312), union 986 pixels) has 980 on `0x80`, and the label `tavern_l.bmp` (28x64 at (144,332)) 1783 of 1792. The shopkeeper `townbirds/shopie` (30 frames of 28x44 at (276,296), union 435) has all 435 on `0x90`, and the label `shop_l.bmp` (52x76 at (264,264)) 3883 of 3952. The gate door `door/t00.bmp` (9 frames of 36x48 at (180,148)) has all 1728 rectangle pixels on `0xa0`, and the guards `townbirds/guards` (8 frames of 48x48 at (184,158), union 696) all 696. The statue star `stars/s00.bmp` (9 frames of 64x44 at (340,288), rectangle 2816) has 1977 on `0xb0` and 3, 1 and 4 on `0x80`, `0x90` and `0xa0`. The school fighter `townbirds/fighter` (11 frames of 36x60 at (516,344), union 670) and mage `townbirds/mage` (11 frames of 36x60 at (452,328), union 819) lie wholly on `0xc0`, and the label `trener_l.bmp` (140x116 at (436,300)) has 13210 of 16240 on `0xc0` and 7 on `0xa0`.

Overlapping none: the sign `sign/v00.bmp` (10 frames of 40x32 at (360,232)), the weathervane `fluger/f00.bmp` (8 frames of 64x64 at (308,64)), `birds1`..`birds9` (57 frames of 552x92, unions 1835..2704 pixels), the `town_add.bmp` overlay, horses 1, 3 and 5, and every baba and dervish variant have 0 action pixels. Horse 2 has 7 opaque pixels on `0x80` in each of its variants, and horse 4 has 163..167 on `0xc0`. A region is larger than its figures: the largest `0xa0` component has 7895 pixels against at most 2424 for the gate door and guards together (their two counts added), so the regions are not the outlines of the animated sprites. Outside the tip popup's rectangle, what a click does at a pixel depends on the mask pixel only (`TOWN-473`), so there a figure is clickable where it lies on a region and inert elsewhere.

**Confidence.** High for the measurement: the anchors are immediates read at instruction level and the counts are taken over the shipped pictures on both roots. Eight figures have at least 98 percent of their pixels or rectangle on one region, which agrees with the view-relative reading of the anchors. No frame was rendered. The counts for the `.bmp` figures (three labels, the gate door and the statue star) are rectangle overlaps and so upper bounds on the visible pixels; their transparent key was not applied.

### TOWN-475

After the pre-send `TOWN-477` describes, `R0704` posts one message to the window `[object+0x1c]` through `PostMessageA` (`L03535`): `0x42b` from arm `L11724` for `0x80`, `0x42a` from `L11725` for `0x90`, `0x42c` from `L11726` for `0xc0`, `0x42d` from `L11727` for `0xa0`, and `0x41f` from `L11728` for `0xb0`, which has no pre-send (`evidence/click-arms.tsv`, `TOWN-163`, `TOWN-164`). The campaign dispatcher `R0701` maps them (`evidence/campaign-arms.tsv`, `evidence/r-campaign.txt`): `0x42b` at `L12115` to the tavern open `R1318` and `0x42c` at `L12116` to the school open `R1319`, each only while `campaign+0x3dc` is 0; `0x42a` at `L12117` to the shop open `R1315` while `+0x3dc` is 0, or is 1 with `+0x6bc` at most 1 as a signed value; `0x42d` at `L12118` to the map open `R1317`, which starts `music\map.wav` and sets `+0x3dc` bit `0x10`; `0x41f` at `L12119` to the town menu constructor `R0363` and `R0361`. The last two have no condition. `TOWN-372` and `TOWN-373` give the `+0x3dc` bits and `DLG-ENTRY-016` the `+0x6bc` phase word. The campaign arms `0x428`, `0x429` and `0x42f` are not posted from the town click table.

The gate arm calls `R1416` on the record at `[campaign+0x548]` first (`L11727`..`L11730`). The helper returns -1 when the record's latch `+0x18` is 0 and no entry of its child array, 76 bytes each with the id at `+4`, holds a latched `+0x18` (`evidence/d-town-click.txt`, `REG-SCN-065`). The arm then opens the dialogue `inn\mercenary\npc35` (literal `L11729`, pushed at `L12120`, host `[campaign+0xd0]`) through `R0695`, posts nothing and sends no `0x445`, so the town view stays active. Otherwise it sends the `0x445` pre-send and posts `0x42d`. The literal names `main.res` node `text/inn/mercenary/npc35.txt` (EN 223 bytes, sha256 `660b51c0ff6c4e79b58724a67d424cf5a21451c4dca45a22a907a37a270a5284`; RU 233 bytes, sha256 `f7ca39b3ceb910b262bbad00854e77342988ea6379252dc657a9db14e7e3c06a`) and `speech.res` node `inn/mercenary/npc35p1.wav` (EN 630792 bytes, RU 634548 bytes) (`evidence/nodes.tsv`).

Two opened screens speak on entry from the campaign's offer lists (`REG-SCN-064`). The shop's enter `R0705` continues only when `L03650` finds `ShopMission` non-empty (`L12121`..`L12122`), builds `shop\npc31m%d` (literal `L12123`, `L09773`) from element 0 and opens it through `R0695`, then accepts the element and removes it. The school's enter `R0706` does the same when the `TCMission` count is above 0 (`L12124`..`L12125`) with `training\npc34m%d` (literal `L12126`, `L12127`) (`evidence/d-capstone.txt`). The text nodes are `text/shop/npc31m<N>.txt` for N = 31, 91, 110 and 130 (EN 4 nodes, 1540 bytes; RU 4 nodes, 1360 bytes) and `text/training/npc34m<N>.txt` for N = 61, 121 and 131 (EN 3 nodes, 1824 bytes; RU 3 nodes, 1674 bytes), with 8 and 10 voice nodes under `speech.res` `shop/` and `training/` (EN 3773720 and 4628740 bytes; RU 4847984 and 5453874 bytes) (`evidence/node-families.tsv`).

The tavern click posts no text of its own; its dialogues belong to the tavern screen (`TAVERN-CLICK-019`, `TAVERN-BUTTON-020`), and the statue opens the town menu (`MENU-ITEM-012`). No separate mercenary hall is reached from the town: `DLG-WIN-001` names the tavern button panel (`R0703`) and this gate arm as the owners of `text/inn/mercenary/` (14 nodes).

**Confidence.** High for the arms, posted messages, conditions, the gate test and the literal sources: each was read at instruction level, and both roots run one program. Medium that these five are the only town click outcomes: the immediates `0x41f`, `0x428`..`0x42f` and `0x445` were enumerated over the whole image and reduced to the town cluster `L12192`..`L12090` (`evidence/refs-reduction.tsv`), and a post built from a computed id would not appear in that census. Medium that `-1` means nothing on offer: the helper was read here, its meaning comes from `REG-SCN-065`.

**Unknown.** Whether the tavern's enter routine speaks on entry (not read), whether a state exists in which a room arm's condition fails while the town view receives the click (the town view has already been left by then), and physical pointer delivery.

### TOWN-476

Guard state on the town view is sheet `+e4` (`interface/townbirds/guards/sprites.16a`, 8 frames of 48x48), frame `+e8`, step `+ec`, direction latch `+f0` and sound object `+7c`. The pointer handler `L11546` runs on each delivered pointer message (slot `+0x4c`, `TOWN-399`). It stores step +1 at `L12128` and the selector at `+b4`. For selector 8 it calls `R1416` and stores at `L12129` step -1 when the result is -1 and +1 for every other value (`evidence/gate-step-domain.tsv`). `TOWN-403` states this contract; the census below lists the writers.

The helper `R1927` adds the step to the frame (`L12130`). With step +1 and latch 0 it requests `SFX\Town\Guard2.wav` and sets the latch (`L12131`); with step -1 and a nonzero latch it requests `SFX\Town\Guard1.wav` and clears it (`L12132`). A frame below 0 sets step 0 and frame 0 (`L12133`, `L12134`), a frame at or above the sheet's frame count sets step 0 and frame count-1 (`L12135`, `L12136`), and each clamp releases the sound object. Its only caller is the paint hub `R1908`, which calls it unconditionally at `L12137`, and the hub's only caller is the paint `R1489` (`L11692`), which runs it only when `timeGetTime` minus the process-wide tick `[L11520]` is above 67 (`L11563`) and resets the tick after it (`L11690`) (`evidence/refs.txt`, `TOWN-404`).

The frame therefore moves one step per admitted hub, at most one hub per 68 ms, and stops at a clamp. The model over these instructions (`evidence/model-vectors.tsv`, derived from the disassembly and not observed at runtime) gives, after enter, frame 7, step 0, latch 0 and no motion. With the pointer over the gate and nothing on offer it gives frames 6, 5, 4, 3, 2, 1, 0 on seven admitted hubs and a clamp on the eighth, with Guard1 requested only when the latch is 1; a further pointer message inside the sweep leaves it running. Any other delivered message, or the gate with an offer, sets step +1. From frame 7 the next hub then clamps at once, requesting and releasing Guard2 with the latch set to 1. From a lower frame the frame rises by one per hub to 7, and Guard2 is requested on the first hub when the latch is 0. A message during motion reverses it from the current frame. The guard step reads the offer only in the selector 8 arm of `L11546`, so a change in the offer moves nothing until the next delivered pointer message; the door helper `R1926` reads it on every hub (`TOWN-403`).

Writers in the town cluster `L12192`..`L12090` (`disp:` and `imm:` sweeps over the repaired project) are those named. `+ec`: `L12138`, `L12133`, `L12135`, `L12128`, `L12129`. `+e8`: `L12139`, `L12130`, `L12134`, `L12136`. `+f0`: `L12140`, `L12131`, `L12132`. `+e4`: `L12141`, `L12142`, `L12143`. The sweeps also hold 20, 25, 22 and 29 stores at `+ec`, `+e8`, `+f0` and `+e4` in routines outside the town cluster and the merchant cluster, and 5 stores of the immediates `0xf0` (four) and `0xf020` (one) to other locations, which are not guard fields. The 96 stores at the four displacements lie in routines outside both clusters and were not attributed to a class (`evidence/refs-reduction.tsv`, `evidence/refs-cluster-hits.tsv`, `evidence/refs-outside-writers.tsv`).

**Confidence.** High for the instruction-level contract (`L11546`, `R1927`, `R1908`, the pacing compare, the unconditional call and the sole callers) and for the model traces as consequences of it. Medium that no other routine writes the guard fields: the census reads the town cluster only, and the outside stores are unattributed. Medium that a 68 ms admission is an elapsed-time cadence in play: neither the resolution of `timeGetTime` nor the paint rate was measured, and no run of the original was observed. The alternatives excluded within the cluster are an independent clock (no store of a timer value into the guard fields), a trigger on enter or return (`TOWN-477`) and a trigger on a change of the offer.

**Unknown.** Which frame shows the crossed halberds, which needs the sheet's pixels; audible playback of the Guard requests; and pointer messages the system generates after a room closes.

### TOWN-477

The four room arms send `0x445` through the town view's own `vt+0x48` (`R1972`, `TOWN-164`). The town handler `L12144` passes `0x445` to the base handler `R0716`, which with the view active (`+0x5c` nonzero) calls `vt+0x84`, the town's leave `R1912`, and posts `0x44c` through the `PostMessageA` wrapper `R1990` (`L12145`..`L12146`, `evidence/d-campaign.txt`, `evidence/d-capstone.txt`). Leave deletes the tip popup, releases the sheets through `L11518` (which stores 0 at `+e4`, `L12143`) and the sounds through `R1913` (`+0x74`..`+0x94`, including the guard sound `+7c`), and runs the base leave `R1228`. It writes no `+e8`, `+ec` or `+f0` (`evidence/d-town-guards.txt`). The gate arm with nothing on offer and the statue arm send no `0x445`, so the town view stays active there.

The return from a room is message `0x42e`. The immediate census finds six routines that push it (`evidence/refs.txt`): the tavern routine `L10201` and the school routine `L12147`, each of which sends `0x445` first (`L12148`, `L12149`), the shop routine `R1757`, and `R0709`, `R1284` and `R1732`, which were not read. Its campaign arm runs the town open `R1320`, which calls the town view's `vt+0x80` at `L12150`. Enter `R1383` loads the sheets through `R1906` (guards stored at `L12142`) and clears the hub flags `+208` (`L12151`), the gate latch `+1a4` (`L12152`) and the guard latch `+f0` (`L12140`). It stores step 0 (`L12138`) and frame count-1 (`L12139`), which is 7 for the eight-frame sheet, clears `+15c` (`L12153`), reseeds timers from `timeGetTime` (`L12154`, `L12155`, `L12156`) and sets the selector `+b4` to -1 (`L11711`). The guards therefore start every visit at frame 7 with step 0 and latch 0, whatever frame, step or latch they held when the town view was left, and in the town routines read nothing moves until a pointer message is delivered (`TOWN-476`, `evidence/model-vectors.tsv`, scenario `close-then-reenter`).

The process-wide statics persist across a visit: the pacing tick `[L11520]` and the once-latch bits `[L11525]` (`L12157`..`L12158`), so the first paint after enter admits its hub at once when more than 67 ms have passed. The tavern interior's own sequence state is separate (`TOWN-412`). Enter clears the gate latch `+1a4` only; the door progress `+1a0` and the other figures' fields were not censused here (`TOWN-403`, `TOWN-439`..`TOWN-441`).

**Confidence.** High for the leave and enter bodies, the `0x445` chain and the model trace, read at instruction level with both roots one program. Medium that nothing writes the guard fields between leave and enter: the census blind spot of `TOWN-476` applies. Medium for the first paint after a return: the model assumes the paint reaches its hub before any pointer message.

**Unknown.** Whether the original runs a hub before the first pointer message after a return, and the state of the door progress across a visit.

### TOWN-478

The shop's merchant panel has vtable `L09658`. Its `+0x54`, `R1526`, with `+0x240 & 0x80` set tests two prompt rectangles built from the fields at `+0xf0`, `+0x100` and `+0x110` and clears the flag on a hit. Otherwise it loops four shelf rectangles at `panel+0x60+0x10*i` (`SHOP-SHELF-047`: (354,110)-(459,295), (169,110)-(274,295), (314,5)-(454,105), (172,5)-(314,105)) and, for a hit where `R1523(i)` is nonzero, sets `+0x240` bit `0x20` and refreshes (`L12159`..`L12160`). It returns 0 on every path (`L12160`). `+0x58`, `L12161`, tests no position: `campaign+0x3cc` decides. `+0x5c`, `+0x60` and `+0x64` are default stubs. The merchant figure (277,112)-(353,288) is placed by the anchors at `L09829` and `L09830`, and its rectangle copy at `+0xe0` (`L12162`) has no displacement read inside the merchant cluster `L13174`..`L13175` (`evidence/refs-reduction.tsv`, `evidence/d-shop.txt`). It overlaps none of the four shelf rectangles (0 pixels each, `evidence/rooms.tsv`). The shop routine `R1757` runs one of four arms when the button under the point equals the pressed one at `+0xf8` (`L12163`..`L12164`), and one arm posts `0x42e` when `campaign+0x3dc` is 2 (`L12165`..`L09758`, `evidence/d-shop.txt`). The hero figure belongs to the borrowed character panel (`SHOP-FIGURE-041`, `TOWN-344`..`TOWN-350`).

The school view has vtable `L03653`. Its `+0x54`, `R1931`, calls `R1477(1, x, y)` when `[+0x328]` is 0 and returns 1, and the pointer slot `+0x4c`, `R2044`, calls it with 0. `R1477` tests the view rectangle with `PtInRect` (`L12166`) and acts only in state `+0x31c` equal to `0xf`, the mage panel `+0xfc` (188,188)-(288,308) with the mask `[+0xf8]`, or equal to 0, the fighter panel `+0x1b0` (192,192)-(284,312). It reads the class mask at `[[mask]+0x10] + 8 + width*(y-top) + (x-left)` (`L12167`..`L12168`, the layout of `TOWN-473`), subtracts `0x37`, takes one of five arms through jump tables and calls `R1968(i, flag)`, which updates the icon's state word in the array `+0x2f4` (mage) or `+0x2a4` (fighter). For a click it then plays the icon's sound (fighter `+0x84`..`+0x94`, mage `+0x98`..`+0xa8`, `R1476`, `R0386`) when idle, and calls `R1954(0)`, which stores at `+0x34c` the code of the first icon whose state is 1 or 3, when it differs from the stored one, and passes the selected hero and that code to `R1955` (`evidence/d-capstone.txt`). Other states do nothing. Both class masks equal their panel in size, 100x120 and 92x120 (`evidence/school-masks.tsv`). The trainers (0,200)-(172,424) and (320,200)-(480,424) and the diamond (200,60)-(280,136) overlap neither panel (0 pixels), and the column picture (168,176)-(316,384) overlaps both (12000 and 11040 pixels) as the ground of the panels' own art (`evidence/rooms.tsv`, `TOWN-154`, `TOWN-428`).

**Confidence.** High for the slot bodies, prompt and shelf tests, class-panel tests and rectangle overlap counts, which are computed from committed anchors and archive sizes. Medium that the merchant, trainers, diamond and column frame are click-inert: the other children of the two views (shop goods grids, backpack, table and buttons; school buttons; the tip popups) and their base handlers were not read, and each figure's rectangle comes from painter anchors and archive sizes, not from a test. The school's tip popup, built at `L12169` over (0,0)-(456,200) while the tips option is on, contains the diamond whole and overlaps the column picture, and the shop's (`SHOP-TIP-045`) overlaps the merchant rectangle; `evidence/rooms.tsv` has no popup row.

**Unknown.** What `R1955` does with the selected skill code and what the school then shows for it, and the click behaviour of the children not read here, the two tip popups among them.

### TOWN-479

The base figure routes below assume the tip popup is absent and no capture child is set. The campaign window procedure `R0701` sends messages `0x201`..`0x206` to the root view `[campaign+0xcc]` through `vt+0x48`. The root's children broadcast `R0642` gives the message to the room view, whose message slot (`R1764` shop, `R0735` school) forwards to the base handler `R0716` and so to the base dispatcher `R0390`. That dispatcher runs the capture child `[+0x34]` when one is set and otherwise the children broadcast, and it calls the view's own slot only when the result is 0 (`evidence/d-dispatch.txt`, `TOWN-211`, `TOWN-406`). The shop view's left-down slot is the base redelivery `R0790`. It offers the press once more to the child `R0791` finds at the point when `[+0x40]` equals `[+0x34]`, and otherwise returns 0 through the stub `L03642`.

Shop merchant. The figure (277,112)-(353,288), 13376 px (`TOWN-478`), lies inside the merchant panel (id `0x3ed`, vtable `L09658`, (164,0)-(480,303)) and inside no other child of the shop view (`evidence/windows.tsv`). Left-down: `R1526` returns 0 for all 13376 px. The shelf-only press model records no action on those pixels because no shelf rectangle overlaps the figure (`TOWN-478`); inertness requires `+0x240 & 0x80` clear or a proved miss of both enabled prompt rectangles. The model contains no prompt-state input. The enabled arm builds its two rectangles from the live fields `+0xf0`, `+0x100` and `+0x110`, which were not evaluated. A hit at `L12170` or `L12171` clears bit `0x80` and calls `R1781` and `R0340`, still returning 0 (`evidence/d-shop.txt`). The base redelivery can call the merchant again; the press ends unconsumed even if the prompt arm acts. Left-up: `L12161` returns 1 for all 13376 px. It reads the campaign object through `R1989`. With `campaign+0x3cc`, the cursor-held item (`TOWN-348`), equal to 0 it does nothing more. With an item held it clears `[[this+0x5c]+0x144]`, sets the default cursor `[L01211]` (cursor slot 0, `SPR16A-CURSOR-067`) through `L09894`, calls `R1781` on `[this+0x5c]` and then `R0340`, which clears the held item (`TOWN-350`). `R1781` sets the same cursor and switches on the origin code `campaign+0x3d4` (1 to 8) to hand the item, with the origin slot `campaign+0x3d0`, to one of the shop's panels; `TOWN-348` names it as the routine a held item released outside the grids reaches. With the prompt arm disabled or a proved prompt miss and no held item, a press and release on the figure posts no text, opens no dialogue, plays no sound and writes no state. With the tip popup shown, its check box (204,258)-(352,274) intersects the merchant figure and consumes down, writing the tips flag `[L03631]`; this is a known overlay action (`TOWN-480`).

School figures. Left-down on the mage trainer (0,200)-(172,424; 38528 px), the fighter trainer (320,200)-(480,424; 35840 px), the diamond (200,60)-(280,136; 6080 px) and the column picture (168,176)-(316,384; 30784 px) reaches the school view's slot `R1931` for every pixel except 608. It calls `R1477(1,x,y)` when `[+0x328]` is 0 and returns 1, so the view consumes the press (`evidence/figure-presses.tsv`). `R1477` acts only in the states `TOWN-478` lists, at the mage class panel (188,188)-(288,308) or the fighter class panel (192,192)-(284,312). The trainers and the diamond overlap neither panel (0 px). The column picture overlaps the mage panel in 12000 px and the fighter panel in 11040, of which 368 lie outside the mage panel. Elsewhere the press returns with no state write, sound or text. The 608 px of the fighter trainer at x 464..479 and y 200..237 lie under the school button panel `0x3fd` (464,0)-(640,238). It takes the press through `L12172` and the release through `R1548`, and neither acts because no button rectangle holds the point (`R1978` returns -1). Left-up returns 0 through the school view's slot `+0x58`, stub `R1812`, on all mage, diamond and column pixels and the fighter's other 35232 px. The fighter's 608-pixel button-panel overlap instead returns 1 through `R1548` without a button action (`L12173`, `evidence/d-school.txt`).

**Confidence.** High for the slot bodies (`R1526`, `L12161`, `R1931`, `L12172`, `R1548`, `R0790`), the routing order and the pixel counts, which are computed from committed anchors and archive sizes on both roots; the executables are one program (`evidence/rom-program-digest.tsv`). In the popup-absent case, alternative (a), a named window acting, holds for an enabled merchant prompt hit, a held item on merchant release or a class panel in the school. Alternative (b), a handler that does nothing, holds for the trainers and diamond, merchant down with the prompt arm disabled or a proved prompt miss, and merchant up without a held item. Alternative (c), another route, is the 608 px button panel, which consumes down and up without a button action. The shown popup adds the known check-box action above; the base no-action result does not cover it. Medium that nothing else acts: the shop's goods grids, backpack, table and buttons and the school's other buttons were read only as far as the routes above, a figure's rectangle is a painter rectangle and no opaque-pixel mask was applied, and the model assumes no capture child at the press.

**Unknown.** The merchant's live prompt state and prompt rectangles from the live fields, the arms of `R1781` beyond its switch, the school view's `+0x328`, and what `R1955` does with the selected skill code (`TOWN-478`).

### TOWN-480

Each room builds one popup of class `L06362` (constructor `R1261`, controls `R1974`) at its enter while the tips flag `[L03631]` is nonzero (`TOWN-185`, `TOWN-186`). The tavern builds it at `L10197` as a child of the roster (id `0x467`), the shop at `L12174` as a child of the merchant panel (id `0x3f3`), the school at `L12169` as a child of the school view (`0x467`) and the town at `L12114` as a child of the town view (`0x467`). Its children are the list `0xd` (vtable `L06234`), the Close button `0xe` (`L03644`) and the check box `0xf` (`L12175`), added in that order. Absolute rectangles (`evidence/windows.tsv`):

| room | popup | list | check box | Close |
|---|---|---|---|---|
| tavern | (160,0)-(472,200) | (180,24)-(444,164) | (200,160)-(348,176) | (352,160)-(432,178) |
| shop | (164,162)-(476,298) | (184,186)-(448,262) | (204,258)-(352,274) | (356,258)-(436,276) |
| school | (0,0)-(456,200) | (20,24)-(428,164) | (40,160)-(332,176) | (336,160)-(416,178) |
| town | (328,0)-(640,200) | (348,24)-(612,164) | (368,160)-(516,176) | (520,160)-(600,178) |

Controls. List: left-down `L12176` calls the stub `L03642` and returns 0; left-up `L12177` sends `0x472` (wParam its id, lParam `[+0x88]`) to the popup and returns 1; move `L12178` is a stub. Check box: left-down `L06426` toggles the bit that `vt+0x78` selects in `[+0x80]`, stores it at `+0x88`, repaints, sends `0x46d` (wParam `0xf`, lParam the mask) to the popup and returns 1; left-up is the stub `R1812` and returns 0. The popup's message slot `R1478` stores that lParam in `[L03631]` (`L07705`..`L07706`), so the flag changes on the press; no routine read deletes the popup when the flag is cleared. Close: left-down `R0778` sets the press state `[+0x6c]` through `L03645(1)` and returns 1 when none is registered; left-up `R0779` clears the state and, when the release point lies inside the button, posts the command stored at `[+0x70]`, `0x45a` (`TOWN-185`, `TOWN-206`), to the campaign window through `R1989` and `R1990`, and returns 1; with no registered press it returns 0. The Close button becomes the popup's focus and capture child while the pointer is inside it and releases capture on a move outside while not pressed (`R0771`, `R0776`, `R0777`); the check box's move slot was read only in outline. The same slot `R1478` returns 0 for `0x100`, `0x445` and `0x446` and forwards every other message to `R0716` (`L07701`..`L07704`). This reverses TOWN-186's forwarding sentence; its flag-write clause stands. The Close slots answer TOWN-207's mouse-binding question.

Presses (`evidence/popup-presses.tsv`; every pixel of each popup, identical on both roots). Left-down passes below on the list and the body, and the list's own rectangle overlaps the check box and Close button in 592 and 320 px, where the redelivery `R0790` gives the press to the control:

| room | popup px | passes below | check box consumes | Close consumes | left-up: list consumes | Close consumes | passes below |
|---|---|---|---|---|---|---|---|
| tavern | 62400 | 58592 | 2368 | 1440 | 36960 | 1120 | 24320 |
| shop | 42432 | 38624 | 2368 | 1440 | 20064 | 1120 | 21248 |
| school | 91200 | 85088 | 4672 | 1440 | 57120 | 1120 | 32960 |
| town | 62400 | 58592 | 2368 | 1440 | 36960 | 1120 | 24320 |

What lies below. Tavern: the roster's cell test, with 0 px of cells under the popup, returns 0 (`TAVERN-FIGURE-021`). Shop: the merchant's `R1526` returns 0 on all 37712 px it receives. The shelf-only model assumes prompt bit `+0x240 & 0x80` clear: 25370 of these pixels lie on shelf rectangles 0 and 1 and set bit `0x20` and refresh when `R1523(i)` is nonzero (`TOWN-478`). An enabled prompt arm can instead clear bit `0x80` and call `R1781` and `R0340` on a prompt hit; its live geometry/state were not evaluated (`TOWN-479`). A separate 912 px reach the shop button panel, which consumes them through `L12179`. The 21248 px whose left-up passes below the popup reach the merchant's `L12161`, which returns 1 and acts only while an item is held. School: the view's `R1931` consumes all 85088 px and acts only inside a class panel, 1200 px of the popup (`TOWN-479`). Town: `R0704` consumes all 58592 px, and the popup's rectangle holds no action pixel (`TOWN-473`), so nothing opens. A left-up on the list is consumed by the list, on the check box and body it passes below, and on the Close button (1120 px) it is consumed.

Closing. The campaign window procedure runs the arm `L12180` for `0x45a`. It looks for a child with id `0x10` in `[campaign+0xd0]` (`R0375`) and, when there is one, removes and deletes it without forwarding. Otherwise it forwards `0x45a` to the root view `[campaign+0xcc]` through `vt+0x48`. Each room's message slot then removes the popup from its parent (`R0384`), deletes it (`vt+4` with 1) and passes the message to `R0716`: tavern `L11615` (popup `[+0x80]`, parent `[+0x7c]`), shop `L12181` (`[+0x88]`, `[+0x74]`), school `L11626` (`[+0x7c]`) and town `L12182`..`L12183` (`[+0x200]`). The room leaves delete the popup too (`TOWN-473`; tavern `L12184`..`L12185`).

**Confidence.** High for the control slots, the pixel table and the closing arms in the four room slots: read at instruction level or computed over the popup rectangles on both roots, which run one program. Alternative (a), a named window acting on the press, holds for the check box and Close button; alternative (b), a handler that does nothing, is the result on the list and body of the tavern; alternative (c), another route, is the shop button panel. Medium for the delivery of `0x45a` to the room: the campaign arm's child `0x10` of `[campaign+0xd0]` was not identified, so a popup closed while it exists is not shown to reach the room arm. Medium for a press after the pointer has moved over a control: the routing above assumes no capture child, and the check box's move slot `L12186` was read only in outline.

**Unknown.** What child `0x10` of `[campaign+0xd0]` is and whether it exists while a popup is shown, the consumers of `0x472`, and whether a game with tips off shows the popup after a later toggle without a re-enter.

**Amended.** Child `0x10` of `[campaign+0xd0]` is the mission tip popup (`TRIG-TIPS-087`). A later toggle of the flag without a re-enter shows no popup, because only the enter constructs one (`TOWN-516`).

### TOWN-481

Town enter `R1383` calls base enter, stores 1 at `this+0x204`, and invokes the repaint slot `vt+0x34`, `R0366`. That slot reaches town paint `R1489` when bit `0x20` of `this+0x18` is clear. The call is synchronous inside enter, before that call returns to queued pointer dispatch (`evidence/d-town.txt`, `evidence/d-campaign.txt`). This is a possible hub step before the first queued pointer message, not an unconditional guard movement.

Town paint runs hub `R1908` only when `timeGetTime` is more than `0x43` (67) ms past `[L11520]`. The first paint of the process stores the current time and runs no hub. The three references to that timestamp in the preserved town listing are `L12187`, `L12188` and `L11691`. The guard routine `R1927` changes no frame when its step is 0; the enter-idle model in `evidence/guard-steps.tsv` holds frame 7, step 0 and latch 0 through four hub ticks. A due hub can therefore run during enter without moving the guards (`TOWN-476`, `TOWN-477`).

The campaign procedure sends `0x200`..`0x206` and `0x400` to the root view `[campaign+0xcc]` through `vt+0x48`. Town message `0x402` skips its repaint when the whole dword `campaign+0x3dc` is nonzero (`L12189`..`L12190`), not only when one bit is set. The application cursor pump `R0370`, called from idle `R0560`, posts `0x400`. Its `[L12191]` flag is set after dispatching `0x201` and cleared after dispatching `0x202`. The campaign procedure's `SetCursorPos` call is conditional on a nonzero root result for `0x200`; the town's corresponding route returns 0.

The room-close and enter bodies preserved in `evidence/d-town.txt`, `evidence/d-shop.txt`, `evidence/d-school.txt` and `evidence/d-tavern.txt` contain no window, capture, cursor or clip API call that establishes a synthetic pointer message on room closure. The import/reference census in `evidence/api-summary.tsv` and `evidence/refs-api.txt` locates those APIs elsewhere. This excludes an explicit call in the read bodies; it does not exclude a message generated by the operating system or an unread indirect route.

**Confidence.** High for enter's call order, the timer comparison, the message routing and the local repaint predicate: these are instructions in the preserved listings, anchored by `evidence/code-anchors.tsv`. Medium that a re-enter paint runs the hub before a queued pointer message: bit `0x20`, elapsed time and the timer's resolution were not observed at runtime. Medium for the absence of a room-close API mechanism outside the read bodies: the exported direct/import references do not cover every computed code pointer. EN/RU instruction agreement is not independent confirmation; the executables have one digest (`evidence/rom-program-digest.tsv`).

**Unknown.** OS-side pointer messages after a room closes, whether a message arrives before physical mouse movement, bit `0x20` at re-entry, timer resolution and actual hub intervals, other repaint sources, door progress across a visit, and town state after room messages ignored following the pre-send leave.

### TOWN-482

The residual store population at displacements `0xe4`, `0xe8`, `0xec` and `0xf0`, outside the guard clusters already attributed by `TOWN-476` and `TOWN-477`, is 96 stores in 33 owners: 29, 25, 20 and 22 stores respectively (`evidence/guard-stores.tsv`, `evidence/guard-store-summary.tsv`). The owner reduction and constructor/caller evidence are preserved in `evidence/guard-store-owners.tsv`. No owner has evidence for town-view class `L11548`. Ten stores in six owners remain unattributed by that reduction.

The town closure begins with 62 roots: 28 functions in the `L12192`..`L12090` cluster and 34 vtable slot targets. Its direct-call closure has 567 functions and reaches none of the 33 residual-store owners (`evidence/guard-cones.tsv`). Two owners have constructor links to class `L00575`, whose vtable the closure stores. Those links leave a live indirect-call alternative; zero direct reach does not establish that none can run between a room click and town re-entry. One owner also reaches the instrument's cap of 12 ancestor-evidence entries. The census concerns these explicit stores, not every bulk copy or differently based access that could touch a containing object.

The guard sheet has eight 48x48 frames on each root. The shape classifier in `evidence/guard-frames.tsv` marks frames 0 and 1 as crossed halberds; the preserved visual inspection corroborated that identification. Frames 2 through 7 are not classified as crossed. A closing sweep with step -1 ends at frame 0, not frame 7 (`evidence/guard-steps.tsv`).

The decoded PCM headers give Guard1 418.14 ms and Guard2 356.96 ms (`evidence/wave-headers.tsv`: 18440 and 15742 data bytes at 44100 bytes/s). In a sweep that requests a sample at tick 1, the model reaches its endpoint at tick 7 and clamps at tick 8. Six hub intervals separate request from the first endpoint paint. Under the integer timer gate's minimum 68 ms interval, that is at least 408 ms: Guard1 can have 10.14 ms left, while Guard2 has already ended. Guard1 overlaps the first closing-endpoint paint only when those six intervals total less than 418.14 ms, equivalently a mean below 69.69 ms (`evidence/guard-timing.tsv`). With the fresh-enter latch 0, the closing sweep requests no Guard1 at all. These are request/duration bounds, not an observation of audible playback.

**Confidence.** High for the recorded store counts, closure counts, decoded frame sizes, PCM durations and step/latch transitions; each names its input and measurement. Medium for residual-store attribution and non-reachability beyond the direct-call closure: class-linked virtual calls, computed pointers, bulk copies, neighbouring-width stores and other object bases are not excluded by this reduced population. Medium for crossed-frame identification, which combines a shape heuristic with visual inspection. Medium for sound at an endpoint: the model and timer predicate constrain it, but actual intervals and audio start latency were not measured. EN/RU code is one program (`evidence/rom-program-digest.tsv`).

**Unknown.** The six unattributed owners' runtime objects, indirect dispatch through class `L00575`, other stores outside the reduced population, actual hub intervals and timer resolution, audio start/output latency, and whether the halberd sound is audible on a native closing endpoint.

### TOWN-483

Gate helper `R1416` reads the main record's latch at `+0x18`. A nonzero main latch returns the main mission; otherwise walker `R1408` searches the children and returns a latched child's mission, or -1 when none is latched. Thus a heard offer and a playable mission are different states: an offered but unaccepted mission does not by itself make this helper return a mission. The latch writer is `R1415`, reached through `R1410`; acceptance is `R0785`. Hearing in the tavern does not write this latch (`SAV-CAMPAIGN-081`).

The preserved acceptance callers are map-marker click `L12193`, tavern leave over the heard list `L10211`, shop opener `L03632` and school opener `L03633`. The caller at `L12194` is in a branch excluded by the preserved `[L10126] = 10` value. `R1411` reads AutoGetMission at record `+0x110`, whose default is -1 and whose shipped Mission10 value is 20 (`REG-SCN-062`, `REG-SCN-064`, `REG-SCN-065`). This identifies a separate automatic-acceptance input; hearing is not a substitute for acceptance.

Loader `R1297` clears the main latch at `L12195`, increments each retained child's age at `+0x48`, and deletes a child at age 2. Completion is `R1414`. After a non-last main win the new main latch is 0; an older latched child still returns its mission while it remains below the deletion age. With no such child, the helper returns -1. An accepted mission not yet played returns its mission through its latch. With no accepted main or child, both no offer and a merely heard offer give -1. These results are conditional on the current latches, not on the caption or whether an offer was heard during an earlier visit.

The last-mission test reads `campaign+0x664` at `L12196`, the loaded record's `+0x11c` LastMission value (`REG-SCN-063`; shipped Mission150 sets 1). When that value is nonzero and the played mission at `campaign+0x660` is divisible by 10, the campaign posts `0x428` and does not take the town-entry route. This does not establish a reachable town state after that final win. A loaded save carries the main and child latches and child ages through the record serializers already described by `SAV-CAMPAIGN-076`, `SAV-CAMPAIGN-080`, `SAV-CAMPAIGN-081` and `SAV-CAMPAIGN-082`; the helper reads the restored state by the same rule.

**Confidence.** High for the local helper/walker rule, latch setter, loader clear/age operations and last-mission branch: their preserved instruction ranges and routine digests are in `evidence/r-record.txt`, `evidence/r-campaign.txt` and `evidence/code-anchors.tsv`. Medium for complete acceptance reachability: the listed callers and registry facts do not exclude callers through computed code pointers or establish every campaign path from runtime. The competing model that hearing itself supplies a gate latch is excluded by the separated hearing and acceptance stores. EN/RU instruction identity is one program, not a second oracle (`evidence/rom-program-digest.tsv`).

**Unknown.** Which non-decade missions use AutoGetMission beyond the stated shipped case, code-pointer callers outside the inspected tables, a save made inside the tavern before its leave path accepts the heard list, whether a last-mission state can enter town by another route, effects of `L03688` in that arm, and the consumer and player-visible result of `0x428`.

### TOWN-484

Builder `R0695`, used for the gate fallback resource `inn\mercenary\npc35`, ignores its host argument. Constructor `R0696` builds panel id `0x9` at (76,124)-(564,356), portrait id `0xc` at (106,178)-(194,292), text id `0xa` at (204,160)-(504,295), and button id `0xb` at (276,296)-(356,322), command `0x46f`. The preserved node selects NPC 35 and initial part 1; construction of its drawable speaker from that selection is an inference, not a separately observed runtime portrait.

Show `R0710` calls pager `R0711`. On each call the pager stops and deletes the prior voice held at `this+0x70`, then starts the next part's voice through `R0584` and `R0386` at volume `[L02739]`. The initial voice therefore starts on show. The next pager call stops it, including the call made by pressing the button on the last page. Leave `L12197` and the destructor do not touch `+0x70`; leaving by another route is not proved to stop the voice. The preserved PCM headers give the part-1 sample lengths as 14302.68 ms EN and 14387.85 ms RU (`evidence/wave-headers.tsv`).

Message slot `R0715` runs the pager for `0x46f`. When the pager answers zero it posts `0x45b` if `this+0x80` is set, then posts `0x445` to itself. The window is non-modal in the measured dispatch model: of the native 640x480 screen's 307200 points, 305120 left-down points reach the town click slot and 2080 are consumed by the button. The blocked points lie over shop 1497, statue 240, tavern 2 and no selected town target 341. Both tips-on and tips-off cases have the same totals; pointer movement reaches the town pointer slot at all 307200 points (`evidence/dialogue-presses.tsv`, `evidence/dialogue-summary.tsv`). The model does not establish the capture child held at a native press.

Town `0x402` skips repaint while the whole dword `campaign+0x3dc` is nonzero (`L12189`..`L12190`). That condition gates the paint containing the guards and crowd. It is not evidence that every clock, pointer handler or other repaint source stops merely because the dialogue is visible. The press and move routes above remain available in the measured dispatch model.

A missing node read throws through `R0533` and reaches the runtime window-procedure catch at `L12198`, then `L12199`, `L12200` and `L12201`. The builder, panel constructor and text loader have cleanup frames but no catch covering this read before the runtime handler (`evidence/d-exception.txt`, `evidence/eh-frames.tsv`). This route builds no substitute dialogue panel and reaches the application's exception report. The exception's displayed text was not recorded. It answers the missing-node and outside-button questions left open by `DLG-KEYS-040` within this route.

**Confidence.** High for the constructor geometry, pager voice operations, message-slot actions, local town repaint predicate and exception unwind/catch chain: they are preserved instruction bodies and exception metadata. High for the stated pixel counts within the supplied dispatcher/window model; Medium that it exhausts native input behavior, since capture state and indirect children remain runtime inputs. Medium for the speaker drawable inference and for guard/crowd behavior beyond the named repaint. EN/RU executable equality is one instruction source; the two PCM lengths are separate data measurements (`evidence/rom-program-digest.tsv`).

**Unknown.** Runtime portrait construction, capture state at a press, other repaint sources, campaign handling of `0x41f` while the dialogue is open, identity and presence of child `0x10` of `[campaign+0xd0]` while a popup is shown, voice lifetime after exits that bypass the pager, and exception message text.

## World-map Return input

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-485 | The world-map key-down and character slots call one progress helper without testing the key value and return handled; its key-up slot returns zero. | High | ● active (branch candidate) | [EXP-0423](../experiments/EXP-0423-world-map-enter/) |
| TOWN-486 | The world-map input helper changes only progress and Cross counters when progress is nonzero; zero progress does nothing, and selection, route and position are untouched. | High | ● active (branch candidate) | [EXP-0423](../experiments/EXP-0423-world-map-enter/) |
| TOWN-487 | Campaign Return key-down and character forwarding discard the third message parameter; keyboard focus and ordered child return values condition whether the world-map handler is reached. | High / Medium | ● active (branch candidate) | [EXP-0423](../experiments/EXP-0423-world-map-enter/) |

### TOWN-485

World-map table `L11605` maps key-down slot `+0x6c` to `L12202`,
key-up `+0x70` to `R1891` and character `+0x74` to `L12203`.
The key-down and character bodies each contain three instructions: call
`R1948`, return value 1, and return with one stack argument consumed.
Neither reads that argument. The key-up body returns 0 without touching
the world-map object. Container dispatcher `R0390` selects these slots
for messages `0x100`, `0x101` and `0x102` respectively.

The original-instruction controls supply keys 0 through 255 to each of
key-down and character at progress 0 and 1, 1024 vectors. Each 256-key
population has one result (`evidence/key-controls.tsv`). This local identity
does not mean the campaign forwards every physical key; TOWN-487 binds Return.

**Confidence.** High for the three table dwords and complete local bodies.
The direct PE/recursive-descent instrument binds instruction digests and
independently computes relative branch operands. A raw file-offset scan
retains the two input-handler pointers at `L12204` and `L12205`, and
all pointer occurrences of the shared key-up body. This does not depend on
named function ownership or assert image-wide key-handler absence.

**Unknown.** Physical delivery, system-key messages, native child lifetime,
other world-map classes and whether a key-down is followed by a character.

### TOWN-486

Helper `R1948` has eleven instructions. It tests the whole dword at
world-map `+0x120`. Zero branches directly to return. Nonzero reads route
count `+0xdc`, adds one in a register and stores the result at `+0x120`. It reads
the Cross object through `+0x7c`, reads its frame count at `Cross+4`, adds one
in a register, and stores the result at world-map `+0x128`. The Cross object
is not written. Both additions have dword semantics. The complete body has
no other store, call, selection read, destination read or current-position read.

The 33 supplied-state/message vectors preserve the selected mission at
campaign `+0x660`, current/destination points at view `+0xf8/+0xfc/+0x100/+0x104`,
route count and synthetic route-array digest. Progress-zero idle and
pre-first-reveal cases write nothing. Nonzero outward, homeward and
fully-revealed/Cross-active cases write the two counter assignments. Independent
route/frame counts 1/1 and 31/13 distinguish dynamic counts from constants.
An already-larger Cross counter is assigned frame-count-plus-one too
(`evidence/state-controls.tsv`). These labels describe supplied fields,
not observed native state admission.

The helper does not copy destination to current or post arrival. The later
paint completion remains TOWN-121's route; this experiment rereads its
counter checks and destination-copy/post instructions but does not replay
paint or establish its cadence.

**Confidence.** High for this complete helper's reads, stores and local state
relation. Its sole conditional transfer is bound to an internal instruction
start. Its operands use one unchanged object base; no whole-image field or
writer enumeration is claimed. Controls execute these original instructions,
and no supplied service replaces the helper or its counts.

**Unknown.** Native admission of selected-idle or already-ready counter states,
malformed counts, arithmetic overflow, invalid Cross pointers, paint scheduling,
visible final animation and the physical mission/town transition.

### TOWN-487

The campaign message-map records at `L06189` and `L12206` bind `0x100`
to `R0231` and `0x102` to `R0774`. For Return value `0x0d`, the first
handler subtracts 8, reads byte-table index 5 as arm 1 and selects
`L03620`. That arm forwards `0x100`, the key value and a literal zero to
root `campaign+0xcc`, slot `+0x48`. The character handler forwards
`0x102`, the character and zero to the same root. The complete handler
listings show that neither reads its second or third argument. Those
arguments therefore cannot distinguish main/keypad Return in these handlers;
their selected forwarding passes a literal zero as the third root parameter.

Root table `L03640` maps `+0x48` to container dispatcher `R0390`.
Keyboard messages `0x100..0x102` take focus pointer `+0x38`, not mouse
capture pointer `+0x34`. A nonzero focused result stops dispatch. A zero
result calls ordered child pass `R0642`, then the root's own slot if no
child consumes the message. That keyboard child pass does not require a
rectangle hit. The world-map message slot `R1949` passes these messages
through `R0716` to the same container dispatcher and its own input slots.
World-map enter `R1409` calls base show `R0775`; that show calls the
capture and focus setters. This binds setters, not persistent live ownership.

Twenty-four private dispatcher vectors vary focus, capture, child order and
an explicit alternate-recipient result. World-map focus reaches the helper;
alternate focus or an earlier child returning 1 excludes it. Zero permits
the child walk. Changing mouse capture alone leaves the supplied keyboard
routes unchanged. Four campaign vectors vary key-down/character and an unread
third argument, supplied as `0x001c0001` or `0x011c0001`. Both values reach
the same local state result; they do not model the native extended flag
(`evidence/dispatch-controls.tsv`, `evidence/campaign-controls.tsv`).

**Confidence.** High for the message-map records, Return table target,
forwarded arguments and named dispatch predicates. Medium for the composed
campaign route: private focus, child identity/order, state, clock and alternate
recipient returns are supplied. The world map's own focus pointer `+0x38`
is supplied null and its child array at `+0x1c` has zero supplied children.
The clock, rectangle and label-string services are declared in
`evidence/services.tsv`; none supplies the world-map helper's
result. Replays stop before unrelated campaign cleanup. These controls do not
establish a native focus or capture observation.

**Unknown.** Physical main/keypad Return translation, system keys, actual
key-down/character ordering, native focus/capture changes, other indirect
recipients, object teardown, campaign key-up entry, and the paint cadence or
presentation after these local assignments. The world-map vtable `+4` routine
`L12207` and the native world-map child population were not read. Native
world-map focus was not observed. Its own null focus and empty child array are
supplied replay conditions; a native focusable child could receive the key before its
own input slot. An authorized native message/focus trace is the next
discriminator.

## Exterior wildlife presentation

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-489 | The town painter draws horse, baba and dervish last and in that order, above the main picture, the bird layer, the door labels and nine static sprite fields; its direct blit census in the town range is 18 sites. | High | ● active (branch candidate) | [EXP-0436](../experiments/EXP-0436-town-families/) |
| TOWN-490 | Horse, baba and dervish are drawn at a view-relative table position with no per-frame offset: one frame size per sheet, all 27 canvases inside 640x480 from the view origin, 14 ending at 480; the view origin was not read. | High / Medium | ● active (branch candidate) | [EXP-0436](../experiments/EXP-0436-town-families/) |
| TOWN-491 | A horse or baba episode starts when time since the family clock strictly exceeds its delay, and every step refreshes that clock; delays are rolled 2000..3999 ms at entry and 2000..6999 ms at each arm, and the wait exceeds the delay. | High | ● active (branch candidate) | [EXP-0436](../experiments/EXP-0436-town-families/) |
| TOWN-492 | Horse, baba and dervish change frame at most once per admitted hub, at least 68 ms apart: horse frames 1..14 span at least 14 hub intervals, baba 30 or 31, one dervish revolution 30; delivered cadence is Unknown. | High / Unknown | ● active (branch candidate) | [EXP-0436](../experiments/EXP-0436-town-families/) |
| TOWN-493 | Horse1, Horse2, Horse3 and Crowd are 22050 Hz mono 16-bit samples of 538.8, 44.7, 29.2 and 3664.6 ms on both roots; horse requests pass repeat 0 and priority 128, the crowd request repeat 1. | High / Medium | ● active (branch candidate) | [EXP-0436](../experiments/EXP-0436-town-families/) |
| TOWN-494 | Horse sound requests are re-tested on every eligible paint while the frame gate holds and pass only when no buffer of the sample plays, so Horse2 (44.7 ms) can start again inside one frame dwell of at least 68 ms. | High / Medium | ● active (branch candidate) | [EXP-0436](../experiments/EXP-0436-town-families/) |
| TOWN-495 | No town code stops a horse sample before its end; the looping crowd buffer and the horse buffers are stopped, rewound, released and zeroed only by sound cleanup, called from entry reload, leave and destruction. | High / Medium | ● active (branch candidate) | [EXP-0436](../experiments/EXP-0436-town-families/) |
| TOWN-496 | In the 12 EN and 11 RU archives `crowd` names one node per root, sfx.res town/crowd.wav; no graphics node path contains `crowd`, and the TownBirds directory holds 41 nodes in nine groups on both roots. | High | ● active (branch candidate) | [EXP-0436](../experiments/EXP-0436-town-families/) |
| TOWN-497 | Every direct call through slots `0x18` and `0x38` in the town range is one of 18 sites; no other painter-route routine in range calls them; the hook's child walk paints the tip popup child (`TOWN-208`), not examined for a crowd. | High / Medium | ● active (branch candidate) | [EXP-0436](../experiments/EXP-0436-town-families/) |

### TOWN-489

Painter `R1489` is reached only through town class slot `+0x2c` (table `L11548`, the single reference `L12208`). In address order it calls the begin-paint helper `R0346`, blits `view+68` (the main picture) through slot `0x18`, calls the bird and `Town_add` routine `R1917`, blits the three door labels (`TOWN-161`), blits nine static sprite fields in the order `+160`, `+180`, `+19c`, `+1bc`, `+1c8`, `+1d0`, `+1d8`, `+1f8`, `+e4` (`TOWN-162`), then draws horse `+120` (`L07764`), baba `+f4` (`L11703`) and dervish `+150` (`L11706`), each through slot `0x18`. The end-paint helper `R0312` and the generic hook `R1311` follow. The horse sound gates sit between the static fields and the horse draw (`TOWN-445`).

The town range `L11530..L12209` holds 18 direct calls through slot `0x18` or `0x38`: 16 in the painter (one main picture, three labels, nine static fields, three wildlife) and two in `R1917` (bird sprites through `0x18`, `Town_add` through `0x38`). The hook `R1311` calls virtual slot `+0x30` of the view, whose town-class target `L12210` is a single `RET imm16`. The helpers `R0346` and `R0312` bracket a surface-lock counter and contain no blit.

**Confidence.** High: the order is the address order of read call sites, the census enumerates the instruction population, and the slot table was read from the EN executable; the RU executable is byte-identical.

**Unknown.** Which pixels of overlapping layers survive the keyed or opaque semantics of each slot, and any draw outside slots `0x18` and `0x38`.

### TOWN-490

The painter's wildlife blits pass `table x + view+8` and `table y + view+0xc` as first and second arguments, with the table values bound by `TOWN-439`. All 27 sheets per root have one width and height for every frame: horse position 1 is 132x76, position 2 136x76, position 3 96x80, position 4 136x80, position 5 148x80; baba positions 1..4 are 44x44, 48x52, 56x56 and 48x48; dervish positions 1..4 are 28x48, 36x52, 32x56 and 28x48. A `.16a` frame header holds width, height and data size only (`SPR16A-STRUCT-001`), so a sheet occupies one rectangle at every frame and its motion lies inside that canvas.

Position tables give view-relative (x, y) top-left rectangles. Horse 1 (104,404), 2 (104,404), 3 (256,344), 4 (448,400), 5 (140,400); baba 1 (216,364), 2 (308,424), 3 (384,424), 4 (580,384); dervish 1 (224,364), 2 (324,424), 3 (392,420), 4 (592,388). All 27 canvases end inside 640x480 measured from the view origin, and 14 end at y=480: horse positions 1, 2, 4 and 5 and baba position 3. Read as (y, x), baba 4 and dervish 4 would start at y=580 and 592.

Bounding boxes of horse overlap baba in 5 of 20 position pairs (80..352 px²) and dervish in 3 of 20 (96..336 px²). Baba and dervish overlap only at equal position index (1232..1664 px²), which the dervish re-roll excludes (`TOWN-004`), so their relative order has no bounding-box effect.

The rectangles are relative to the town view: each blit adds `view+8` and `view+0xc`, and this experiment did not read those fields for the town view. `TOWN-222` shows the origin addend is zero only under the default 640x480 configuration, for the tavern.

**Confidence.** High for the table values, frame sizes and rectangle arithmetic, identical on both roots. Medium for the x-then-y, top-left reading of the blit arguments: it rests on the fit above and on `TOWN-474`'s figure overlaps, not on a rendered frame.

**Unknown.** The town view's `+8/+0xc` origin and so the absolute position under any display configuration; opaque pixel extents inside the canvases; rendered composition.

**Amended.** The origin and the argument-order clauses: `+8` and `+0xc` are read (left and top, zero on the default 640x480 configuration) and the first table pair is x then y by the blit routine's address arithmetic, High (`TOWN-503`). [`retracted.md`](retracted.md) holds the entry.

### TOWN-491

Entry (`R1383`) writes the baba clock `+118` and horse clock `+148` from `timeGetTime`, rolls the baba delay at `+11c` and the horse delay at `+14c` as `(r*2000)/0x7fff mod 2000 + 2000`, sets current `-1`, selects element 0 of the sheet array, and arms the dervish with current 0 and bit `400h`. Helper `R1907` runs on every eligible paint after the hub block: it arms a family when `now - clock` is unsigned-greater than its delay. An arm writes clock = now, current = 0, a new delay `(r*5000)/0x7fff mod 5000 + 2000` and a sheet index `(r*n)/0x7fff mod n` with n = 2 for baba and 3 for horse, then sets the bit (`L12211..L11701`).

The step helpers `R1909` and `R1910` write the family clock from `timeGetTime` at every step (`L12212`, `L12213`). After the terminal step the clock holds that step's time, and the delay rolled at the arm applies from there. The idle gap from the last step of one episode to the start of the next is therefore the delay plus the interval to the next eligible paint, and the wait from entry to the first episode is the entry delay plus that interval. The same test runs while a bit is active: a pause between two steps above the delay restarts that family at frame 0 on a freshly drawn sheet. The three horse sheets and two baba sheets are drawn independently of the previous sheet. The delay bounds assume the generator returns 0..0x7fff (`AI-RAND-058`).

No chapter, mission or save input reaches these paths (`TOWN-004`). Eligibility is the view-active word `+204`, stored 1 near the end of entry (`L12214`) and cleared on leave (`TOWN-446`).

**Confidence.** High: every write, comparison and formula operand was read at instruction level (`evidence/anchors.tsv`).

**Unknown.** The random generator's output distribution and the longest realised pause between paints.

### TOWN-492

The hub block runs when unsigned elapsed time since `[L11520]` is above `0x43` (`L11563`, `L11689`). Elapsed time is in `timeGetTime` milliseconds, so consecutive hub runs are at least 68 ms apart. All three families step in the same hub run, in the order dervish, baba, horse (`R1908`). The image-wide call sites of `R1908`, `R1909`, `R1910` and `R1911` and the three direct references of `L11520` are those of `TOWN-446` and `TOWN-447`, which bound the once-per-hub reach. Each step advances one frame, so one frame lasts at least 68 ms.

Horse and baba episodes show frame 0 from the arm to the first step, frames 1..count-1 at one step each, then return to frame 0 at the terminal step. Horse: 15 frames, frames 1..14, at least 14 intervals (952 ms) from the first step to the terminal step. Baba A1: 31 frames, at least 30 intervals (2040 ms); A2: 32 frames, at least 31 intervals (2108 ms). Dervish: 30 steps per revolution, at least 2040 ms. The first step follows the arm by at least one later paint, because the hub block precedes the arm inside one paint.

The executable imports `timeGetTime` and none of `timeBeginPeriod`, `timeEndPeriod`, `timeSetEvent` or `SetTimer` (`evidence/timer-imports.tsv`, both roots). Clock granularity is the operating system's current timer resolution unless another module changes it; this experiment did not observe it.

**Confidence.** High for the gate, the step order and the interval arithmetic. Unknown for the delivered paint cadence: `TOWN-481` shows entry paints synchronously and the hub waits for a paint, but no source of repeated paints was traced.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Unknown.** Actual step intervals, stalls and repeated frames.

### TOWN-493

Header fields read from `sfx.res` on both roots: `town/horse1.wav` 23762 data bytes (538.8 ms), `town/horse2.wav` 1972 (44.7 ms), `town/horse3.wav` 1288 (29.2 ms), `town/crowd.wav` 161608 (3664.6 ms); format tag 1, one channel, 22050 Hz, 16 bits, 44100 bytes per second (`evidence/wav-headers.tsv`). Duration is data bytes over bytes per second.

Horse requests pass volume `[L02733]`, pan 0, repeat 0, priority `0x80` (`L12215`). The crowd request passes repeat 1 and priority `0x80` (`L12216`, `L12217`). The request routine `R0386` hands repeat to `R1919`, which calls buffer slot `0x30` with flag 0 for repeat 0 and flag 1 otherwise. Read against the published `IDirectSoundBuffer` method order, slot `0x30` is `Play` and flag 1 is looping, slot `0x24` is `GetStatus`, `0x3c` `SetVolume`, `0x40` `SetPan`, `0x48` `Stop`.

**Confidence.** High for header values and request arguments. Medium for the looping reading, which rests on the interface layout and not on an audio observation.

**Unknown.** Resampling, mixing and audibility.

### TOWN-494

The painter reaches the horse sound block on every eligible paint while `+120` is non-null. The block compares the current frame with its gate (`TOWN-445`), tests that the sample object exists, then calls status helper `R1476` and requests only on a zero result. The helper returns zero when the sound subsystem word `L02703` is zero, when the object's loaded flag at `+0xc` is zero (`+0x10` is the buffer table), or when no buffer reports playing through slot `0x24` bit 0. It returns nonzero only when a buffer reports playing and a channel-table entry whose buffer field equals that buffer exists; a playing buffer with no table entry reads as not playing. The request routine selects the first non-playing buffer, returns silently when every buffer plays or the subsystem word is zero, and otherwise applies volume and pan and plays (`VIDEO-SFX-060`).

The frame changes only at a hub step, so a gate frame is current for every paint inside one dwell of at least 68 ms. Horse1 (538.8 ms) outlasts that minimum. Horse2 (44.7 ms) ends inside it: a paint that arrives after the sample has ended and before the next hub step requests it again.

**Confidence.** High for the gate, status helper and request selection. Medium for the repeat inside one dwell: it needs a paint in the window between sample end and the next step, and paint cadence is Unknown (`TOWN-492`).

**Unknown.** Whether a repeated Horse2 start occurs in a play session, and its audibility.

### TOWN-495

Sound cleanup `R1913` handles the twelve town sound fields. For `+74` (crowd) and `+78` it calls status helper `R1476` and, when a buffer plays, calls buffer slot `0x48` and then slot `0x34` with position 0 (`L12218`, `L12219`), releases the object and zeroes the field. Fields `+98`..`+a8`, which include the three horse samples, go through `R2045`, which does the same status, stop, rewind, release and zero sequence. Fields `+7c`..`+94` use `R2046`.

The cleanup is called from the sound loader, town leave and the destructor (`TOWN-420`). The only crowd request site is the entry tail, so a crowd buffer that stopped by another route is not restarted until the next entry. Horse samples play once (`TOWN-493`). No town-range code other than sound cleanup stops these buffers. The shared channel allocator stops a playing sound of lower priority than a new request when all 16 channels are busy (`VIDEO-SFX-060`), so a request of higher priority than 128, if one exists, can end a horse or crowd buffer.

**Confidence.** High for the cleanup bodies and callers. Medium that no other stop route exists: the search is the call and reference populations of `TOWN-420` and `TOWN-445`.

**Unknown.** Device-side stops and buffer loss.

### TOWN-496

A case-insensitive substring sweep over every node path of all 12 EN archives (`MUSIC.RES`, `VIDEO4.RES`, `VIDEO8.RES`, `KIDS.LM`, `graphics.res`, `main.res`, `movies.res`, `patch.res`, `scenario.res`, `sfx.res`, `speech.res`, `world.res`) and all 11 RU archives (the same without `KIDS.LM`) finds the stem `crowd` in one node per root, `sfx.res` `town/crowd.wav`; no other node path contains it. It finds `baba` in 8 and `dervish` in 4 `graphics.res` nodes per root and no sound node for either. `horse` appears in 16 `graphics.res` nodes (15 `TownBirds` sheets and `infowindow/horse.bmp`) and 5 `sfx.res` nodes (`TOWN-445`).

The `graphics.res` directory `interface/TownBirds` holds 41 nodes per root: baba 8, birds 9, dervish 4, fighter 1, guards 1, horse 15, mage 1, shopie 1, tavern 1 (`evidence/townbirds-population.tsv`). Both roots agree.

**Confidence.** High for the enumerated population. The sweep covers names, not picture content.

**Unknown.** A crowd under another name (people, citizens, townsfolk, npc); a crowd drawn inside a named picture; a crowd loaded by a computed path; the map `.alm` files, which the sweep does not cover. `KIDS.LM` is 24 bytes and has no nodes.

### TOWN-497

The census of `TOWN-489` accounts for every direct call through slots `0x18` and `0x38` in the town range: 16 painter sites (`evidence/painter-sequence.tsv`, `evidence/blit-census.tsv`) and 2 bird-layer sites. The painter also calls `R1916` (in range) and, outside the range, the begin and end helpers and hook `R1311`. The hook's virtual slot `+0x30` is empty in the town class (`evidence/town-class-slots.tsv`); it also calls `R1865`, `R1224` and `R1991`, the child-paint walk (`TOWN-208`), which paints the tip popup child that draws a tiled panel (`TOWN-209`). Of the 34 slots from `L11548` to the next table, 11 lie in the town range and 23 are shared routines outside it. Entry creates one child, the tip popup at `+200` (`TOWN-165`), only when the tips flag is set.

No sprite sheet, picture field or draw site is named for a crowd, and the crowd sound is requested once at entry (`TOWN-420`). Together with `TOWN-496` no crowd owner was found through slots `0x18` and `0x38` on the painter route.

**Confidence.** High for the census and slot table. Medium for the negative: it covers the painter route through two slots, not other windows or the 462 undisassembled bytes of the town range.

**Unknown.** Draw calls through other object slots; the body of `R1916`; the hook helper bodies `R1865`, `R1224` and `R1991`; the tip popup child's content; crowd imagery inside `townmain.bmp`, `Town_add.bmp` or another named picture; owners reached by computed or indirect targets outside the town class.

**Amended.** The 462 undisassembled bytes hold no code, and the three hook helpers, the bird routine and the tip popup's draw route were read and carry no sprite draw (`TOWN-507`). [`retracted.md`](retracted.md) holds the entry.

## Town square remainder

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-503 | The town view's `+8` and `+0xc` are its left and top, set once by its constructor; the default 640x480 configuration gives a zero origin, and the first pair of each wildlife position table is x, then y. | High / Medium | ● active (branch candidate) | [EXP-0462](../experiments/EXP-0462-town-square-remainder/) |
| TOWN-504 | The horse sample volume `[L02733]` is the effects volume, written by the initializer, the registry load and the options slider; a priority-128 request takes an idle channel or stops a lower-priority sound, else is dropped. | High / Medium | ● active (branch candidate) | [EXP-0462](../experiments/EXP-0462-town-square-remainder/) |
| TOWN-505 | No `srand` call lies in the town range: three direct sites seed from the clock in sound, scenario and AI manager initialisation, and entry draws at least five times in fixed order. | High / Medium / Unknown | ● active (branch candidate) | [EXP-0462](../experiments/EXP-0462-town-square-remainder/) |
| TOWN-506 | A horse, baba or dervish sheet that fails to open is not skipped or stubbed: the loader reports "FATAL ERROR: can't load" plus the path through the application object and aborts the process. | High / Medium | ● active (branch candidate) | [EXP-0462](../experiments/EXP-0462-town-square-remainder/) |
| TOWN-507 | The 462 undisassembled bytes of the town range hold no code, and no crowd draw owner is found in the range, its child widgets or the archive node names; crowd imagery inside the named pictures was not examined. | High / Medium / Unknown | ● active (branch candidate) | [EXP-0462](../experiments/EXP-0462-town-square-remainder/) |

### TOWN-503

The town view is constructed once, at `L12220`, by `R0602` with the class id and the rectangle (left, top, right, bottom) read from the globals `L09808`, `L09809`, `L09810` and `L09811`. The constructor chain `R0602`, `R0759`, `R0773`, `R0388` passes the rectangle to `CRect::SetRect` (`R1980`) on `this+8`, so `+8` is left and `+0xc` is top. Left and top are stored once each, at `L09812` and `L09813`, as half of the screen width minus 640 and half of the screen height minus 480. The screen size is 640x480 by default and 800x600 or 1024x768 when the command line holds `-800` or `-1024`. A command-line token is tested first, then the `RESOLUTION` registry value under `HKLM\SOFTWARE\1C\Allods`; an absent value selects `-640`. The origin is therefore zero on the default configuration and half the size excess otherwise (`evidence/anchors.tsv`, `town-ctor-*`, `screen-*`, `resolution-*`).

The painter pushes the position-table second field plus `[view+0xc]` first and the first field plus `[view+8]` last. Arguments are pushed right to left, so the first argument of blit slot `0x18` is table field one plus left and the second is table field two plus top (`L12221..L07764`). Slot `0x18` forwards to `R1518`, which computes the destination as base plus second argument times pitch plus first argument times two and clips the first argument against the left and right edges. The first argument is x and the second is y. The loader stores table dword one at `+0x138` (horse), `+0x10c` (baba) and `+0x154` (dervish), with dword two at the next field, so the table's first pair is x then y.

No store to `+8` or `+0xc` through a register was found in the 38 functions of the town class table `L11548`, and no `SetRect`, `OffsetRect` or `SetRectEmpty` call site lies in the town range (`evidence/raw-manifest.tsv`). The view is not moved after construction in that population. The older rows that print a destination as (second field plus `+0xc`, first field plus `+8`) list the pushes in push order, not argument order (`TOWN-159`).

**Confidence.** High for the constructor chain, the single stores of left and top, the formula, the three screen sizes and the x-then-y argument order; the RU executable is byte-identical to the EN one (`evidence/exe.tsv`). Medium that the origin is constant after construction: the callees of the 38 functions were not scanned for stores. The shipped default configuration gives zero because `RESOLUTION` is absent from a clean install; the value on a running machine is Unknown.

**Unknown.** The `RESOLUTION` value of the machine that runs the game; stores to the view rectangle by routines outside the town class table.

### TOWN-504

`[L02733]` is the effects-volume field at `+0x10` of the sound configuration at `L03126`. Three code paths write it. The static initializer `L03127`, through `R0658`, stores the default, -700. The registry load `R0660`, called from `L03133` in `R1985`, which `L12222` calls, reads `SoundSfxPos` into the field and ignores a missing value. The options handler `L03131` for control 7 stores it at `L03135` as the mapping of the slider position through `R0659`, the quadratic mapping of `VIDEO-SFX-060`. The save routine `R1984` only reads it. `L03135` is the only direct write reference to the address in the image (`evidence/anchors.tsv`, `volume-*`). The horse request at `L12223` loads it and passes it with pan 0, repeat 0 and priority `0x80` (`TOWN-493`).

The request routine `R0386` clamps the volume at the lower bound and applies it through buffer slot `0x3c`. Its channel pick is `R2047`. It first selects the first non-playing buffer of the sample, the 16 buffers per sample being counted in `L12224` and set by `L12225` and `R0576`; when all play, the request returns silently. It then takes the first channel entry that is empty or whose buffer is not playing. With all 16 entries busy, the victim is the entry of lowest priority strictly below the request priority, the first one on a tie. The victim is stopped through `R2046` (slot `0x48`, then slot `0x34` with position 0) and cleared. When no entry is below the request priority, the request is dropped silently (`evidence/anchors.tsv`, `alloc-*`).

A census of the 112 direct call sites of `R0386` shows priority immediates `0x80` at 79 sites, `0xdc` at 15 and `0x64` at 3, all three in `R0509`; 13 take the priority from a register and 2 were not parsed (`evidence/sfx-request-census.tsv`). A priority-128 horse request therefore displaces only a priority-100 sound or a register-priority sound, never another 128 or a 220, and is dropped when no entry qualifies. This restates the allocator of `VIDEO-SFX-060` for the town and answers the conditional of `TOWN-495`: requests at priority `0xdc` exist at 15 sites and can stop a playing horse or crowd buffer when all channels are busy.

**Confidence.** High for the three writers, the clamp and the victim rule, read as instructions. Medium that no other writer exists: the sound configuration is also reached through its base pointer, and that population was not enumerated. The register-priority sites were not resolved to values.

**Unknown.** The volume value on a running machine; which register-priority requests exist and their values.

### TOWN-505

`R0291` is `srand` and `R0179` is `rand()` (`AI-RAND-058`). A thread block starts with seed 1 (`L12226`). The image holds three direct `srand` call sites and no table entry pointing at it: `L12227` in sound initialisation `R0576` with `timeGetTime`, reached from `R2048` (callers `R2049` and `R0701`); `L12228` in `R1716` with `timeGetTime`, called from the scenario constructors `R0967` and `R0512`; and `L00376` in the AI manager constructor `R0124` with `time()`, which has 6 callers (`evidence/srand-sites.tsv`). No direct call to `srand` lies in the town range `L11530..L12209`, and no data reference to `srand` exists. The direct-call closures of town entry `R1383`, the painter `R1489` and the click handler `L11546` reach none of the three sites (`evidence/srand-closure.tsv`). No constant seed was found.

Town entry draws in this order (`evidence/draw-order.tsv`). The loader `R1906` draws the horse index, the baba index and the dervish index, the last repeated while it equals the baba index. The horse index is `(r*5/0x7fff) mod 5`; baba and dervish use `(r*4/0x7fff) and 3`. The two scaled quotients are near-uniform and not `r mod 5` or `r mod 4` (this corrects `TOWN-004`). The loader tail then calls the sign stepper `R1924` and the vane stepper `R1925`, each drawing once only while its static (`L11559`, `L11560`) is zero (draw sites `L12229` and `L12230`; the later draws `L12231` and `L12232` are not entry draws, because `+0x1cc` and `+0x1d4` are reset first). Back in entry, the baba delay (`L12233`) and the horse delay (`L12234`) are drawn, each `(r*2000/0x7fff) mod 2000 + 2000`; entry sets the dervish bit `0x400` and the active word `+0x204` and calls town slot `0x34`, which reaches the painter synchronously.

The painter draws on the first paint of the process only, when the hub clock `[L11520]` is unset, one bird interval at `L11685`. Every hub run draws twice, at `L12235` and `L12236`, each `(r*100/0x7fff) mod 100`, then runs the steppers. Every eligible paint runs the arm helper `R1907` with four draw sites. The bird routine `R1916` has three draw sites and the click handler `L11546` one. A re-entry in the same process finds the hub clock stale, so its entry paint probably runs the hub; this is an inference from the read stores.

The static direct-call closures of town entry, the painter and the click handler (481, 223 and 205 functions toward `srand`, 480, 218 and 200 toward `rand`; 246, 128 and 100 computed call sites not followed, `evidence/srand-closure.tsv`) reach only draw sites in the town range, and do not reach the music routines `R1230` and `R2050`. The image holds about 88 direct `rand` call sites in 44 owner functions (87 references by `EnumRefs`, 88 call encodings by a byte count).

**Confidence.** High for the three direct `srand` sites, the absence of a data reference to `srand`, the absence of an `srand` call in the town range, the entry draw order and the formulas, read as instructions. Medium that the town does not reseed and that no other consumer draws between entry and the first family step: the closures leave 246, 128 and 100 computed call sites unfollowed. Unknown for a reseed or draw through those sites, by other threads or by idle-time callers.

**Unknown.** The seed value of a running session; a reseed or draw behind the unfollowed computed calls or on other threads.

### TOWN-506

The loader builds the horse, baba and dervish sheets with `R1771(path)`, which calls `R0753`. It opens the archive node through `R0535`. On failure it builds the text of the literal at `L11995` (`FATAL ERROR: can't load `) followed by the path, calls virtual slot `0x94` of the application object with the text and two zero arguments (`L12237`), and then calls `R0754`, which unregisters the window classes, removes the hooks (`L12238`) and calls `_abort`, which does not return (`L03529`; `evidence/anchors.tsv`, `sheet-*`). The loader does not test the returned object. The outcome is an error message and process termination: no skip, no null stub and no replacement sheet.

Both roots carry every sheet (`TOWN-439`), so the path does not run on a lawful install.

**Confidence.** High for the failure branch, its text and the abort. Medium that slot `0x94` displays the text as a message box: the slot was read as a call with a text argument and not observed to draw. An allocation failure (`new` returning 0) stores a null in the sheet array; its effect was not traced (Unknown).

**Unknown.** The display form of the message; the effect of a null allocation.

### TOWN-507

The town range `L11530..L12209` holds 39 functions covering 12915 bytes and 462 undisassembled bytes in 38 runs (`evidence/undisassembled-runs.tsv`). Thirty-four runs are `0x90` padding. Four are switch tables, each a `0x90` byte, dword code pointers and a byte index table, two of them (`L12239`, `L12240`) with an `8b ff` no-op after the `0x90`, at `L12239`, `L12241`, `L12240` and `L12242`; each is referenced from one indirect jump through a table indexed by `reg*4` at `L12243`, `L12244`, `L12245` and `L12246`. No run holds code.

The computed calls and import calls in the range are: slot `0x18` 17 times (16 painter blits and one bird layer), slot `0x38` once, slots `0x20` and `0x24` as the bird layer size getters, `+4` 34 times (release), `+0x34` 10, `+0x48` 9, `+0x7c` 11, 4 jump tables, 15 import-slot calls (`PostMessageA` 9, `timeGetTime`, `PtInRect`) and 6 register-indirect `timeGetTime` calls. No GDI blit import is called in the range. The hook `R1311` calls town slot `0x30`, an empty return in the town class, and the child walk `R1991`. The children are the text button of class table `L03644` (id `0x445`, created at `R1942`), whose paint `R0768` uses text and rectangle primitives, and the tip popup `R1975`, which tiles a nine-slice panel. The helpers `R1865`, `R1224` and `R1991` and the bird routine `R1916` carry no sprite draw.

A case-insensitive substring sweep of every node path in 12 EN and 11 RU archives over 18 stems (`crowd`, `people`, `peopl`, `citizen`, `townsfolk`, `folk`, `npc`, `mob`, `pedestrian`, `passer`, `populace`, `villager`, `human`, `audience`, `spectator`, `horse`, `baba`, `dervish`) finds `crowd` only in `sfx.res` `town/crowd.wav` and finds no node for `people`, `peopl`, `citizen`, `townsfolk`, `folk`, `mob`, `pedestrian`, `passer`, `populace`, `villager`, `audience` or `spectator`. `npc` matches inn text, speech and `npc.reg` nodes and `human` matches interface bitmaps, a unit palette and unit nodes (`evidence/archive-stems.tsv`). The town pictures under `interface/town*` other than `townbirds` are 43 nodes per root with identical hashes on both roots; `townmain.bmp` is 640x480 and `town_add.bmp` 552x92 (`evidence/town-pictures.tsv`).

**Confidence.** High for the enumerated populations. Medium for the absence of a crowd draw owner: it covers the town range, its child widgets and the two painter slots, not other windows.

**Unknown.** Crowd imagery inside the pictures, which were not examined; draws through indirect targets outside the town class; the `.alm` map files, which the sweep does not cover.

## School clock, state dwords, speech and pressed command

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-499 | School global `L11666` is written and read only in school paint `R1487`, which gates the column step at more than 83 ms; paint frequency is Unknown, with one idle-handler repaint request path found. | High / Unknown | ● active (branch candidate) | [EXP-0456](../experiments/EXP-0456-shop-school/) |
| TOWN-500 | The ten school state dwords start at -1 and a hit test turns 2 into -1 and 3 into 1; the main loop skips -1 and `R1964` paints the idle icon; the readers of the marker, timestamp and flag bit 3 are listed. | High / Medium | ● active (branch candidate) | [EXP-0456](../experiments/EXP-0456-shop-school/) |
| TOWN-501 | A pressed-and-hovered school button draws its on picture and captions one pixel lower in the hover colour object; every other state draws the off picture in the object chosen by hover. | High / Medium | ● active (branch candidate) | [EXP-0456](../experiments/EXP-0456-shop-school/) |
| TOWN-502 | The school mouse-up handler `R1548` requests the teacher speech `training\npc33s%dl%d` or `npc34s%dl%d` after a paid training step, once per latch, then posts `teach.wav`. | High / Medium | ● active (branch candidate) | [EXP-0456](../experiments/EXP-0456-shop-school/) |

### TOWN-499

An image-wide scan for displacement and immediate `L11666` finds three hits, all in `R1487`, the school's own-surface paint (vtable `L03653` slot `+0x2c`): the first-paint write of now minus `0x64` (`L12247`, flag byte `L12248` bit 0), the read that subtracts the global `[L11666]` from the loaded value (`L12078`) with a comparison of the difference against 0x53 and a branch when below or equal (`L12079`), and the write of now (`L12249`). The gated block (`L12079`..`L12250`) steps the column pictures through `L11521`, `R1962`, `L11522` and `R1963` and adds `[+0x320]` to `[+0x31c]`; it runs once more than 83 ms have passed since the last step. Other school timers are separate globals: `L12251`, `L12252`, `L12253` (with `L12254`, `L12255`, flag `L12256`) in `R1964`, and `L12257` (flag `L12258`) in `R1971`. The message arm `0x402` of `R0735` repaints through `vt+0x34` when `[session+0x3dc]` bit 3 is clear (`TOWN-063`), and the idle handler `R0560` posts `0x402` once per call on its repaint-only path to `[campaign+0xcc]`. The image imports no timer API for this (`SESS-TIMER-022`). One repaint request path is therefore identified: the idle handler's repaint-only arm, taken when `[session+0x3dc]` bit 0 is clear and the receiver's field at `+0xbc` is non-zero, posts one `0x402` per call, and the school arm repaints unless `[session+0x3dc]` bit 3 is set. That the school paints once per idle iteration is an inference from this path alone (Medium); the idle handler's arm with bit 0 set (`R0260`, `R0454`, `R0640`) was not read, and the delivery from `[campaign+0xcc]` to the school view was not traced.

**Confidence.** High for the three-hit population and the gate. Unknown for paint frequency; Medium for the inference of one paint per idle iteration on the identified path (`TOWN-481` shows the town view receives `0x402` directly).

**Unknown.** The paint frequency; the idle handler's bit-0-set arm; the root-to-room delivery of `0x402`; the idle-handler call rate.

### TOWN-500

The ten state dwords are the fighter slots at `view+0x2a4`..`0x2b4` and the mage slots at `view+0x2f4`..`0x304`. The constructor `R1940` and the enter routine `R0706` (loop `L12259`..`L12260`) store -1 in all ten. The hit test `R1477` visits every dword not equal to -1, clears bit 1, and stores -1 when the value is then 0: 2 becomes -1 and 3 becomes 1; `R1968(slot, flag)` sets the target's bit 1. The main paint loop (`L12261`..`L12262`) skips a -1 dword and picks the state picture by `[view+(state+3*slot)*4+0x264]` (fighter) or `+0x2b4` (mage). `R1964` (called at `L12263`) paints the idle shine icon of the cycling slot: it advances `[L12253]` modulo 5 every 500 ms or more while no slot has bit 1, and uses the shine picture (`+0x2bc` or `+0x26c`) for -1 or bit 0 clear and the shine-on picture (`+0x2c0`, `+0x270`) when bit 0 is set. Timestamp `view+0x340` is read only at `L12264` (against `0xc8`) in the school range and written by paint, enter and the picker steps `R1959` and `R1958`; the class-change marker `view+0x348` is read only at `L12265` and cleared at `L12266`. The picker steps write -1 to all ten dwords, set `+0x340` to the time, and OR bit 3 into the new element's flag `+0x18c`; the enter routine ORs bit 3 into the hero element at `L12267`. Direct tests of element flag bit 3 are at `L12268`, `L12269`, `L09798` and `L12270` (info-window builders). Other readers of the states are `R1954` (training quote) and `L11143` (hover cursor).

**Confidence.** High for the dword values, the hit test and the painter routines. Medium for the reader lists: the direct bit-3 tests were found at four sites in the list of 232 `+0x18c` accesses of all classes (`evidence/scan-element-flags-18c.txt`); indirect consumers (register copies, whole-dword copies and pushes such as `L12271`, `L12272`, `L12273`, `L12274`, `L12275`) were not classified.

**Unknown.** The mage branch of `R1964` calls `R1862` and then indexes by the cycle counter directly; its purpose was not resolved. The consumers of copied flag words.

### TOWN-501

The school button painter `R1905` loops over rectangles `0x360`..`0x368` of the view. For rectangle i it picks the colour object `[L09922]` when the hovered index `[this+0xb0]` equals i, else `[L09928]` (the pair of `SHOP-107`). When `[this+0xb0]` is non-negative and `[this+0xac]`, `[this+0xb0]` and i are equal (pressed and hovered) it blits the on picture `[this+ebx-0x2ec]` (`L12276`) and draws the caption at y offset `+9` and the value at `+1` (`L12277`); otherwise it blits the off picture `[this+ebx-0x2e4]` (`L12278`) with caption offset `+8` and value `+0` (`L12279`..`L12280`, no constant added). Both texts sit one pixel lower when pressed. Both texts go through `[L06186]` `vt+0x14` with flag `0xa`; the value is the grouped quote `view+0x360` or the difference `view+0x364` (`L12281`). The loader `R1973` names the button pictures.

**Confidence.** High for the branch condition, the pictures and the offsets. Medium for the on and off naming: it follows the picture slots read, not the loader's file names.

**Unknown.** The RGB of the two colour objects (`SHOP-107`).

### TOWN-502

The mouse-up handler `R1548` (slot `L12076` of the button child) reads the pressed index `[this+0xac]` (0..2), confirms release by `R1978`, and for button 0 runs the training step `R1957`. If the purse and price fields (`+0x358`, `+0x354`) allow it, it sets `+0x264`, issues the command `R1932`, and reads the hero element `[L11629][view+0x33c]`, the skill byte `[hero+slot+0x14a]` and flag `+0x18c` bit 1 (mage). The tier is 0 up to 20, 1 up to 50, else 2. A latch dword at `view+0x36c+4*(class + 6*slot0 + 2*tier)` must be non-zero, where class is 0 or 2; the latches are set to 1 by `L12282` and zeroed after the request (`L12283`). The key is `training\npc33s%dl%d` (fighter, `L12284`, `L12285`) or `training\npc34s%dl%d` (mage, `L12286`, `L12287`), with slot code 1..5 and tier plus 1, prefixed `Speech\` and suffixed `.wav`; the sound object is built by `R0584` and requested through `R0386` (`L12288`). Then `SFX\Town\School\teach.wav` is posted (`L12289`). The three key literals each have one immediate reference, all in `R1548`.

**Confidence.** High for the routine, the keys and the request. Medium for the latch index: the arithmetic gives the mage tier t and the fighter tier t plus 1 of one slot the same dword, and the next slot's first fighter tier follows the last mage tier; no run was made. The `training\npc34m%d` movie key belongs to the enter routine (`DLG-RECT-037`).

**Unknown.** Who unlatches between visits (`R2021` calls `L12282` twice); the role of `[L02739]`.

## World-map mission marker lifecycle

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-MARKER-508 | The marker routine `R0785` is entered from five sites: the world-map click (selected mission), world-map view setup, shop entry, training-hall entry and one unreached arm; in play nothing removes a marker. | High / Medium | ● active | [EXP-0478](../experiments/EXP-0478-mission-scripts/) |

### TOWN-MARKER-508

`R0785` runs `R1413`, the router `R1407`, the announce latch `R1410` and `R1325`. The last reads `MapObject<n>` for the mission (`MISSION-MAP-073`), skips a mission whose entry is already present (`*piVar4 == param_2`) and a `Picture` of `nothing`, and appends a 20-byte record to the cache at campaign `+0x15c/+0x160`; the cache is serialized by `R1294`, loaded by `R0434` and rehydrated by `R1490`. The background `GMap.bmp` and the town labels never change. For the owner's example, mission 70 selects `MapObject3`, whose overlay is `onmap03` (data); that `onmap03` depicts the monastery rests on the owner's example, and no picture was read.

Callers of `R0785` (`callto:R0785`, 5 hits, 5 owners), each decompiled:

1. `L12193`, `R1945`, world-map click: sets the clicked node's flag `+0x48` to 1, then passes the node's mission number (`+0x28` of the 0x4c-byte node). This is the selection case.
2. `L12194`, `R0701`: one call, directly after `R1414` at `L10125`, with argument `L10126`, in the `else` of `if (L10126 < 0xb)`. `refto:L10126` finds two reads and no write, and the image holds 10 there, so the arm is not reached with that value.
3. `L10211`, `R0786`, world-map view setup: loops over an int array (`this+0x114`, count `this+0x118`) and passes every entry. The array was not identified.
4. `L03632`, `R0705`, shop entry: passes the first element of an array; the same value keys the `shop\npc31m%d` movie.
5. `L03633`, `R0706`, training-hall entry: passes the first element of an array (`+0x648`, count `+0x64c`); the same value keys the `training\npc34m%d` movie.

A marker can therefore appear when the world map opens or when the player enters a shop or a training hall, without a click on the mission; the click is the only caller whose argument is the selected mission. Completion is not a caller.

The marker array's size routine `R2051` has five call sites in four owners (`callto:R2051`): destructor `R1803` (`L12290`), campaign reset `R0757` (`L12291`, size 0), loader `R0434` (`L12292`, `L12293`) and appender `R1325` (`L12294`). `R0757` opens `Scenario.reg`, resets the mercenary arrays and counters and empties the marker array, so campaign reset clears the cache and load `R0434` replaces it. Outer displacements `0x6a4` and `0x6a8` show two readers (`R1530`, `R1863`) and no writer. No call outside reset, load and destructor shrinks the array.

**Confidence.** **High** for the builder, the dedup by mission number, the `Picture` gate, the data and the five callers with their arguments as read. **Medium** for when a mission first gets its marker (three callers pass array contents not traced to their source) and for "no routine removes a marker during play" (the instrument is the callers of the size routine and the displacement sweep of `+0x15c/+0x160`; a removal through another routine type was not searched).

**Unknown.** The array that world-map setup, shop entry and training-hall entry hand to the routine; the picture's draw order against the scroll list.

## Tips in the rooms and the character generator

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TOWN-516 | The town, tavern, school and shop build their tip popup at every room enter while `TipsMode` is set, with no latch; with the flag clear the enter deletes a leftover popup. | High | ✔ promoted (branch candidate) | [EXP-0506](../experiments/EXP-0506-tips-state-machine/) |
| TOWN-517 | The shop's second tip `shop2` replaces `shop1` in the open popup at most once per shop activation, on the shop's idle message, when a tray getter returns nonzero; it does not read `TipsMode`. | High / Medium | ✔ promoted (branch candidate) | [EXP-0506](../experiments/EXP-0506-tips-state-machine/) |
| TOWN-518 | Pre-create step `+0x1dc` restarts at 0 (`chrsel1`) on every enter; a portrait click at step 0 sets 1 (`chrsel2`) and a level click at step 1 sets 2 (`chrsel3`); nothing else moves it. | High | ✔ promoted (branch candidate) | [EXP-0506](../experiments/EXP-0506-tips-state-machine/) |
| TOWN-519 | `R1999` is the pre-create guided cycle: per step it highlights the 4 portraits, the 3 levels, or the amulet and OK, waits 500 ms after hover and steps after more than 300 ms. | High | ✔ promoted (branch candidate) | [EXP-0506](../experiments/EXP-0506-tips-state-machine/) |
| TOWN-520 | Pre-create targets come from `Mask.bmp`: `R1880` returns its pixel as a region code (levels 20/40/60, portraits 80..140, amulet 160, OK 180); the draw rectangles are literals in the loader. | High | ✔ promoted (branch candidate) | [EXP-0506](../experiments/EXP-0506-tips-state-machine/) |
| TOWN-521 | Portrait and level state bit 0 means chosen and bit 1 hovered; state 1 draws `on`, 2 draws `l`, 3 draws `lon`, 0 draws no overlay; a click chooses, the pointer hovers. | High | ✔ promoted (branch candidate) | [EXP-0506](../experiments/EXP-0506-tips-state-machine/) |
| TOWN-522 | The detailed page shows `chrgen1m` (mage) or `chrgen1f` at enter and `chrgen2` after the first skill click while `TipsMode` is set; its skill cycle runs until that click or until the popup closes, without reading `TipsMode`. | High | ✔ promoted (branch candidate) | [EXP-0506](../experiments/EXP-0506-tips-state-machine/) |
| TOWN-523 | Pre-create shows cursor slot 5 `select` at enter and at every paint; the detailed page shows slot 0 `default` at enter and at paint unless the current cursor is `dice`. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0506](../experiments/EXP-0506-tips-state-machine/) |

### TOWN-516

Each room enter tests `[L03631]` and, when it is nonzero, reads its node and constructs the popup (`MENU-137`). When it is zero, the enter removes and deletes any popup still stored. The enter routines, gates and stores:

| room | enter | gate | node | popup field | parent | absolute rectangle |
|---|---|---|---|---|---|---|
| town | `R1383` | `L11742` | `town.txt` | view `+0x200` | the view | (328,0)-(640,200) |
| tavern | `R1412` | `L13535` | `inn.txt` | `+0x80` | roster `[+0x7c]` | (160,0)-(472,200) |
| school | `R0706` | `L11771` | `training.txt` | `+0x7c` | the view | (0,0)-(456,200) |
| shop | `R0705` | `L09820` | `shop1.txt` | `+0x88` | merchant panel `[+0x74]` | (164,162)-(476,298) |

The rectangles are TOWN-480's. No field outside the room's own popup field records that a tip was shown, so the popup returns at the next enter. A popup closed with Close does not return until the room is entered again (`MENU-137`).

**Confidence.** High: each gate, read, construction and deletion is read at instruction level, and the raw scan of `L03631` classifies all 16 references (`MENU-135`).

### TOWN-517

The shop's message `0x402` arm (`L13536`) runs when `[campaign+0x3dc] & 8` is clear and calls `R1760` at `L09823`. `0x402` is the periodic idle post that `L02963` and `L13537` send to the root view. `R1760` acts only when all three hold:

1. Popup `+0x88` exists.
2. The getter `L13538` (`[this+8]`) on `[[view+0x70]+0x84]` returns nonzero.
3. Latch `+0x8c` is 0.

It then sets `+0x8c` to 1, reads `shop2.txt` and retexts the popup through `R1770`, which calls `R1987`. The shop activation stores 0 in `+0x8c` at `L13539`. The arm reads no `TipsMode`. Because the popup exists only when the flag was set at activation, a flag cleared later still lets `shop2` replace the text.

**Confidence.** High for the call chain, the latch and its reset. Medium that the test means "the tray holds an item": `view+0x70` is the tray (`SHOP-TRAY-024`), and the `+0x84` object's `+8` was not traced to its writer.

### TOWN-518

Page vtable `L07691`; enter `R1475`. While `TipsMode` is set, the enter builds the popup with `chrsel1.txt` into `+0x1c0`. With the flag clear it deletes any popup there. In both cases it stores 0 in step `+0x1dc`.

The step routine `R2167(region)` has one caller, `L13540`, in the left-down slot `+0x54` (`R1878`). That slot first runs the hit-test `R1879(x,y,1)`. The step routine returns at once when `TipsMode` or `+0x1c0` is 0. Otherwise:

- portrait regions `0x50`, `0x64`, `0x78` and `0x8c` at step 0 set step 1 and retext `chrsel2.txt`;
- level regions `0x14`, `0x28` and `0x3c` at step 1 set step 2 and retext `chrsel3.txt`.

Other routes leave the step unchanged:

- The move slot `+0x4c` (`R1882`) passes `wParam & 1` to the hit-test, so a drag with the button held selects without moving the step.
- The key slot `+0x6c` (`R2168`) sends Enter to `R2116` and Escape to `R2117`.
- The name field is a separate child.

The leave slot `R2169` deletes the popup. Message `0x45a` (Close) deletes it in the message slot `R2162` and stores 0 in `+0x1c0`.

**Confidence.** High: the single caller and the jump table are in the committed listing and `switches.tsv`.

### TOWN-519

The paint `R1474` calls `R1999` once, at `L13541`, only when `[L03631]` and `+0x1c0` are both nonzero. Its argument is the region under the pointer, from `R1880` on `[L01257]`,`[L01258]`. The timers `T_hover` `L13542` and `T_step` `L13543`, and the index `L13544`, are set on first use and never reset by page enter. The length `L13545` starts at 1 and the last hovered region `L13546` at −1.

Each call:

1. If less than 500 ms (`0x1f4`) has passed since `T_hover`, set `T_step` to now and draw nothing.
2. When a last hovered region is recorded, set the index to the target after it. The region-to-index table at `L13547` maps 20→1, 40→2, 60→3, 80→1, 100→2, 120→3, 140→4, 160→1 and 180→2.
3. Draw by the step `+0x1dc`:
   - Step 0: if a portrait region is hovered, store now in both timers, record the region and return. Otherwise the length is 4 and the index wraps mod 4. It draws `+0xc0` (`lon`) when that portrait's state is 1 and `+0xac` (`l`) otherwise, at rectangle `+0x110[index]`.
   - Step 1: the same for the levels, with length 3, art `+0xfc` or `+0xe8` and rectangles `+0x124`.
   - Step 2: hovering region 160 or 180 freezes the cycle the same way. Otherwise the length is 2: index 0 draws `Amulet.bmp` (`+0x15c`) through its sub-rectangle `+0x164..+0x170`, index 1 draws `ButtonOk.bmp` (`+0x160`) through `+0x174..+0x180`.
   - Any other step: length −1, no draw.
4. Clear the last hovered region. When more than 300 ms (`0x12c`) has passed since `T_step` and the length is not −1, advance the index mod the length and store now in `T_step`.

The cycle runs at the page's paint rate, which `0x402` drives (`vt+0x34` in `R2162`). It stops when `TipsMode` clears or the popup is deleted.

**Confidence.** High for the gates, constants, targets, art fields and order, read at instruction level. The wall-clock period also depends on the repaint rate, which this experiment did not measure.

### TOWN-520

`R1880(x,y)` tests the point against the page rectangle (`PtInRect`). It then returns the byte of `Mask.bmp` (`+0x6c`) at that pixel, with row stride 640. Mask.bmp is TOWN-223's unlocated node. The hit-test `R1879(x,y,click)` then works in this order:

1. Clear bit 1 of every portrait and level state.
2. On a click, region 20/40/60 chooses level 0/1/2 (`+0x1d0`), clearing the other levels' bit 0. Region 80/100/120/140 chooses portrait 0/1/2/3 (`+0x1cc`).
3. Set bit 1 on the hovered item.
4. Region 160 stores `Amulet.bmp` in `+0x184` and region 180 stores `ButtonOk.bmp` in `+0x188`.

The loader `R0597` writes the draw rectangles as literals, page-relative (left,top)-(right,bottom):

- levels 0..2: (60,196)-(120,270), (0,110)-(76,222), (48,65)-(148,217);
- portraits 0..3: (16,273)-(164,424), (124,166)-(260,302), (288,130)-(408,259), (416,190)-(548,342);
- amulet: (528,140)-(640,344);
- OK: (468,373)-(568,429).

The rectangles place art. Only the mask decides what a click hits. This answers TOWN-223's Unknowns for Mask.bmp's consumer and the rectangles' literal values.

**Confidence.** High: the jump tables `L13548` and `L13549` and the loader stores are in the committed evidence.

### TOWN-521

Pre-create paint draws each level and portrait by its state, with fields as in TOWN-223. Portrait files: `+0x98` `{mf,ff,fm,mm}on`, `+0xac` `…l`, `+0xc0` `…lon`. Level files: `+0xd4`, `+0xe8`, `+0xfc` = `level{0,1,2}{on,l,lon}`. The state values mean:

| state | meaning | art |
|---|---|---|
| 0 | unchosen, not hovered | none; `MainArea.bmp` shows |
| 1 | chosen | `on` |
| 2 | hovered | `l` |
| 3 | chosen and hovered | `lon` |

Bit 0 is set by a click and bit 1 by the hit-test under the pointer (`TOWN-520`). The page enter chooses portrait 0 and the level in `+0x1d0`, both state 1. The guided cycle draws over this: `lon` on a chosen target and `l` on an unchosen one (`TOWN-519`).

The sparkle `Blind\sprites.16a` (`+0x70`) plays at a rectangle drawn at random from the 3 levels, the amulet and OK (`+0x84`, filled in `R0597`). Its frame step is 63 ms and its interval 500 + rand/65 ms. It reads neither `TipsMode` nor the step.

On the detailed page the skill art `+0x78` `on` draws for state 1, `+0x8c` `shine_off` for state 2 and `+0xa0` `shine_on` for state 3. The hover handler `R2170` sets bit 1. Fighter skills are sword, axe, mace, pike and bow; mage skills are fire, water, air, earth and astral.

**Confidence.** High: the paint selectors, hit-test and enter stores are read at instruction level.

### TOWN-522

Page vtable `R1312`; enter `R1870`. While `TipsMode` is set, the enter reads `chrgen1m.txt` when `[hero+0x18c] & 2` (mage) and `chrgen1f.txt` otherwise. It builds the popup into `+0x80`, as a child of the skill panel `+0x7c`, at absolute (160,280)-(472,480). With the flag clear it deletes any popup there. It stores 0 in step `+0x100`.

A skill click (slot `+0x54` `R1552`, hit-test `R1216` returns 0..4 or −1) chooses the skill, plays a sound and calls `R1986` at `L13550`. `R1986` acts only when `TipsMode` is set, `+0x100` is 0 and popup `+0x80` exists. It then retexts `chrgen2.txt` and increments `+0x100`. The message slot `R1960` deletes the popup on `0x45a` and stores 0 in `+0x80`. The leave `L13551` deletes it.

The skill cycle `R2004` runs from the skill panel paint `R1871` (at `L13552`). It acts only when `+0x100` is 0 and `+0x80` is nonzero, and it reads no `TipsMode`. Its timing is TOWN-519's: 500 ms after hover, then a step after more than 300 ms. The statics are `L13553`, `L13554` and `L13555`, with last hovered `L13556` = −1 and length `L13557` = 5. Hovering a skill freezes it, and it resumes after that skill. It draws `+0xa0` (`shine_on`) when the skill's state is 1 and `+0x8c` (`shine_off`) otherwise, at the skill's rectangle (`+0xb4`,`+0xb8`,`+0xdc`,`+0xe0` per index).

With `TipsMode` cleared through the popup checkbox, the popup stays and the cycle keeps running. A later skill click no longer advances the step.

**Confidence.** High: gates, constants and callers are in the committed listing and scans.

### TOWN-523

The pre-create enter calls `R0320` with `[L06217]` at `L13558`. That is cursor slot 5, `graphics\cursors\select\sprites.16a`. The paint calls it again at `L13559` on every pass, so `select` shows over every control on that page. The detailed enter sets `[L01211]`, slot 0 `default`, at `L13560`. Its paint `R1310` sets `default` again unless the current cursor is the `dice` slot (`MISSION-066`).

**Confidence.** High for the calls and slots read.

**Unknown.** Which routine sets `dice` on the detailed page. Searched: the committed detailed-page listings, where the `dice` slot `[L06731]` is read only by the paint compare at `L06734`; no image-wide census of `L06731` was run. Settled by that census.
