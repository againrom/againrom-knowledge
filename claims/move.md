# Claim registry — MOVE (unit movement and path selection)

Level 2 ledger. Index: [registry.md](registry.md) · spec: [`formats/move/format.md`](../formats/move/format.md). Legend: `✔ promoted` (also in the spec) · `● active` · `✖ retracted`. IDs are permanent.

Not a file format — the **simulation area** that moves a unit from A to B. It sits on top of the
planes `TERR-PASS-049…053` derived, the mover block `TERR-MOVE-054…058` identified, and the
`data/map.reg` reader `RES-CODE-020` names. Opened by
[EXP-0054](../experiments/EXP-0054-unit-movement/), which read the routines at instruction level;
the tick-order half — what fixes the walk order that *is* the contention priority — closed by
[EXP-0055](../experiments/EXP-0055-tick-order/) (`MOVE-TICK-013…017`); the fallback half — which
cell a failed search settles for, and whether anything distributes destinations across a group — by
[EXP-0068](../experiments/EXP-0068-alt-goal/) (`MOVE-ALT-018…022`, `MOVE-ORDER-023`); and the
**movement domain** itself — what the mover's one dispatch byte selects, and what each domain is
permitted to enter — by [EXP-0078](../experiments/EXP-0078-movement-domains/)
(`MOVE-DOM-024…028`); and the **rate** — what makes one unit faster than another, and against which
clock — by [EXP-0093](../experiments/EXP-0093-move-rate/) (`MOVE-RATE-029…034`), which also closes
`TERR-MOVE-058`'s Unknown and corrects a clause of `MOVE-STEP-010`. **When** that group term applies — the two gates
`MOVE-GROUP-030` located and could not identify — by [EXP-0094](../experiments/EXP-0094-group-rate/)
(`MOVE-GATE-035…037`), which also finds that a group move **does** distribute destinations and that nothing clears
the term.

Vocabulary used below is the **engine's own**, taken from its `data/map.reg` key names: a *static*
search reads the block plane without occupancy and cannot see units; a *dynamic* search reads the
plane that carries occupancy. Each produces its own route list on the actor.

| ID | Claim | Confidence | Status | Evidence |
|----|-------|-----------|--------|----------|
| MOVE-SEARCH-001 | The route search is `R0053`, and it is a double-buffered label-correcting wave — not Dijkstra, not A\*. | High | ● active | [EXP-0054](../experiments/EXP-0054-unit-movement/) |
| MOVE-COST-002 | The step costs, read out of the image, and there is no distance estimate at all. | High | ● active | [EXP-0054](../experiments/EXP-0054-unit-movement/) |
| MOVE-TERM-003 | Termination: three exits, one of them a generation budget — and no node budget, no frontier bound, no closed-set limit. | High / Medium | ● active (amended, partially retracted) | [EXP-0054](../experiments/EXP-0054-unit-movement/), [EXP-0068](../experiments/EXP-0068-alt-goal/) |
| MOVE-ROUTE-004 | The route is extracted by walking the label field downhill from the goal, and its tie-break is asymmetric and favours the last straight step. | High | ● active | [EXP-0054](../experiments/EXP-0054-unit-movement/) |
| MOVE-PLANE-005 | Two searches over two planes, two routes on the actor — and the static search is structurally unit-blind. | High / Medium | ● active | [EXP-0054](../experiments/EXP-0054-unit-movement/) |
| MOVE-PARAM-006 | The six `[Path Finding]` scalars — the engine's own names, and the shipped file agrees with the code defaults 7/7. | High | ● active | [EXP-0054](../experiments/EXP-0054-unit-movement/) |
| MOVE-CLAIM-007 | There IS a reservation: a unit marks the cell it intends to enter, before it enters it. | High / Unknown | ● active | [EXP-0054](../experiments/EXP-0054-unit-movement/) |
| MOVE-WAIT-008 | When the next cell is occupied the unit waits, facing it — and it neither pushes, swaps, nor steps aside. At step time there is no collision test at all. | High / Medium | ● active | [EXP-0054](../experiments/EXP-0054-unit-movement/) |
| MOVE-TICK-009 | The per-tick update imposes no priority, so insertion order decides which actor gets a contested cell. | High | ● active (partially retracted) | [EXP-0054](../experiments/EXP-0054-unit-movement/), [EXP-0055](../experiments/EXP-0055-tick-order/), [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |
| MOVE-STEP-010 | Movement is sub-cell, 1/256 of a cell per axis, and the position is a byte pair per axis. | High | ● active (amended, partially retracted) | [EXP-0054](../experiments/EXP-0054-unit-movement/) |
| MOVE-SPEED-011 | Speed does not feed the search. | High | ● active | [EXP-0054](../experiments/EXP-0054-unit-movement/) |
| MOVE-REFRESH-012 | When a route is recomputed — per cell for the dynamic one, per target change for the static one, and nothing is staggered. | High / Medium | ● active | [EXP-0054](../experiments/EXP-0054-unit-movement/) |
| MOVE-TICK-013 | The tick loop's container is a pooled doubly-linked list embedded in a 0x20-byte manager at `[L00240]`, its walk is head→tail, and its order is pure insertion history — the only insert operation that exists for it is AddTail. | High | ● active | [EXP-0055](../experiments/EXP-0055-tick-order/) |
| MOVE-TICK-014 | What fixes the order at map load: the creators' call sequence — campaign party first, then the map's type-6 records in ascending record order — and the list is per-*ticking-thing*, not total. | High / Medium | ● active (contested, partially retracted) | [EXP-0055](../experiments/EXP-0055-tick-order/) |
| MOVE-TICK-015 | How the order changes during play — and death leaves no hole. | High | ● active | [EXP-0055](../experiments/EXP-0055-tick-order/) |
| MOVE-ID-016 | The runtime id (`actor+0x04`, `SAV-ID-015`'s) is a lowest-free-bit bitmap allocation — creation-ordered only until the first death decays, and never the walk order. | High | ● active | [EXP-0055](../experiments/EXP-0055-tick-order/) |
| MOVE-TICK-017 | The walk order is NOT preserved across save/load — the saved stream never carries it, and the loader rebuilds it grouped by player. | High / Medium | ● active | [EXP-0055](../experiments/EXP-0055-tick-order/) |
| MOVE-ALT-018 | The substitution is not the caller's — `R0053` does it itself, in three branches, and `altTarget` is a target *actor*, not a cell. | High | ● active | [EXP-0068](../experiments/EXP-0068-alt-goal/) |
| MOVE-ALT-019 | Picker A (`R0392`): expanding square rings around the requested cell, and the whole ring is scanned before the best in it is taken. | High | ● active | [EXP-0068](../experiments/EXP-0068-alt-goal/) |
| MOVE-ALT-020 | Picker B (`R0435`): the contact ring around the target actor, entered where the line between the two actors crosses it and walked in both directions. | High / Medium | ● active (amended) | [EXP-0068](../experiments/EXP-0068-alt-goal/), [EXP-0494](../experiments/EXP-0494-pursuit-search-cadence/EXP-0494.md) |
| MOVE-ALT-021 | What a candidate is tested against, and what the choice is measured from — and the answer to both is the label plane, so the substitute is the cell cheapest to reach *from the mover*, not the cell nearest the click. | High | ● active | [EXP-0068](../experiments/EXP-0068-alt-goal/) |
| MOVE-ALT-022 | The substitute is never written back — it is consumed by one route extraction and forgotten. | High / Medium | ● active | [EXP-0068](../experiments/EXP-0068-alt-goal/) |
| MOVE-ORDER-023 | There is no multi-unit destination distribution anywhere: no formation, no offset table, no spread. A group move writes the same cell into every member's order block. | Medium | ● active (partially retracted) | [EXP-0068](../experiments/EXP-0068-alt-goal/), **[EXP-0094](../experiments/EXP-0094-group-rate/)** |
| MOVE-DOM-024 | The movement domain is one byte with six consumers, and passability is only one of them. | High / Medium | ● active | [EXP-0078](../experiments/EXP-0078-movement-domains/) |
| MOVE-DOM-025 | The mask is installed once, at spawn, is the mover's only passability state, and survives a save because the mover is serialized whole. | High / Medium | ● active | [EXP-0078](../experiments/EXP-0078-movement-domains/) |
| MOVE-DOM-026 | What each domain is permitted to enter, terrain and non-terrain counted apart — and the axis is not terrain, it is which of the three block bits the mask carries. | High / Medium | ● active | [EXP-0078](../experiments/EXP-0078-movement-domains/) |
| MOVE-DOM-027 | Nothing narrows a mover's verdict after the mask: no height term, no corner rule, no order-time check, no step-time check. | High / Medium | ● active | [EXP-0078](../experiments/EXP-0078-movement-domains/) |
| MOVE-DOM-028 | Which shipped classes can be non-ground — and that a player can never command one. | Medium / Unknown | ● active | [EXP-0078](../experiments/EXP-0078-movement-domains/) |
| MOVE-RATE-029 | The whole rate law, composed, with the instruction for every term — and the composition is five terms, not one. | High | ● active | [EXP-0093](../experiments/EXP-0093-move-rate/) |
| MOVE-GROUP-030 | The speed override `TERR-MOVE-058` left Unknown is a GROUP term, and it is the minimum `Speed` over the group's members. | High / Medium | ● active (amended, partially retracted) | [EXP-0093](../experiments/EXP-0093-move-rate/), **[EXP-0094](../experiments/EXP-0094-group-rate/)** |
| MOVE-TURN-031 | Stepping requires matching facing; route destruction and final default0x10 were false interpretations. | High | ● active (partially retracted) | [EXP-0093](../experiments/EXP-0093-move-rate/), [EXP-0324](../experiments/EXP-0324-turn-continuation/), [EXP-0329](../experiments/EXP-0329-mover-rate-binding/); [retraction](retracted.md) |
| MOVE-CLOCK-032 | The displacement is one step per actor per SUB-TICK, it is counted in ticks and never measured in milliseconds, and the presentation tick that drives animation is issued from the same loop iteration. | High / Medium | ● active | [EXP-0093](../experiments/EXP-0093-move-rate/) |
| MOVE-LIMIT-033 | The customisation limits of the rate (goal G2), by complete enumeration of the law's input domain rather than by sample. | High / Medium | ● active | [EXP-0093](../experiments/EXP-0093-move-rate/) |
| MOVE-DIR-034 | The eight directions, the diagonal constant, and the per-axis step — the last three things a consumer needs to compute the next position. | High | ● active | [EXP-0093](../experiments/EXP-0093-move-rate/) |
| MOVE-GATE-035 | The group rate term applies exactly when the group moves in FORMATION — one local flag decides both, and it is governed by two gates neither of which is data. | High | ● active | [EXP-0094](../experiments/EXP-0094-group-rate/) |
| MOVE-FORM-036 | A group move DOES distribute destinations: there is a formation, there is a per-member offset table, and there is a spread test — `MOVE-ORDER-023`'s headline is refuted, through the blind spot that row's own confidence cell named. | High | ● active | [EXP-0094](../experiments/EXP-0094-group-rate/) |
| MOVE-GROUP-037 | The complete writer set of `grpAI+0x44`, and the finding that NOTHING clears it — so a group keeps a rate it can no longer justify. | High / Medium | ● active | [EXP-0094](../experiments/EXP-0094-group-rate/) |
| MOVE-AREA-038 | The complete movement-side reader set for both block planes, and the one channel through which an area effect reaches it. | High / Medium | ● active | [EXP-0174](../experiments/EXP-0174-area-movement/) |
| MOVE-GATE-039 | (rom.exe) No actor is displaced except from inside one call of the order machine, so an actor the machine refuses cannot move at all. | High / Medium | ● active | [EXP-0179](../experiments/EXP-0179-effect-action/) |
| MOVE-STEP-040 | (rom.exe) The refusal begins on a cell centre, so a step already in flight completes. | High | ● active | [EXP-0179](../experiments/EXP-0179-effect-action/) |
| MOVE-EFFLIST-041 | (rom.exe) No routine in the movement path reads the drawable's effect list, so the list is not a second immobilisation mechanism. | High | ● active | [EXP-0182](../experiments/EXP-0182-effect-at-actor/) |
| MOVE-072 | Movement rule recorded under MOVE-072. | High / Medium | ● active | [EXP-0386](../experiments/EXP-0386-engagement-under-occupancy/EXP-0386.md) |
| MOVE-073 | Movement rule recorded under MOVE-073. | High / Medium / Unknown | ● active | [EXP-0386](../experiments/EXP-0386-engagement-under-occupancy/EXP-0386.md) |

### MOVE-SEARCH-001

**The route search is `R0053`, and it is a double-buffered label-correcting wave — not Dijkstra, not A\*.** `__thiscall(world, actor, srcX, srcY, dstX, dstY, staticFlag, altTarget)`. State: a **u16 label plane at `world+0x30000`** (65 536 cells, 256-stride, cleared to `0xffff` by `REP STOSD` of `0x8000` dwords at `L01912`), `label[src] = 0` at `L01965`; and **two frontier lists** used alternately — each generation relaxes every cell of the current list and appends every cell whose label **improved** to the other list, then swaps. Two list shapes for the same two slots: parallel **byte** arrays x`world+0x50008` / y`+0x51008` and x`+0x52008` / y`+0x53008` (4096 entries each, the `0x1000` stride being the x→y offset) when the mover's footprint side is > 1, and **u16 packed-cell** arrays `+0x545b4` / `+0x565b4` when it is 1; both shapes share the u16 fill counts `+0x54008` / `+0x5400a`. There is **no priority queue, no ordering of any kind, and no closed set** — a cell is re-expanded every time its label improves, and the relaxation is the same nine-cell scan (centre included) in five places: `R1326` (static plane, n×n), `R1327` / `R1328` (dynamic, n×n), `R1329` / `R1330` (dynamic, 1×1, 8 neighbours unrolled), plus two copies inlined in the driver itself — the static 1×1 pair, one half-pass per list (pushes at `L06754…L06755` and `L06756…L06757`, eight neighbours each), and the static n×n half-pass built on the predicate `R1331` (push at `L06758`/`L06759`)

**Confidence.** High (the plane's clear, its stride, the seed, both list shapes, both counters and the improve-then-append form are named instructions, `evidence/rom-move-excerpt.md` §2–3. The live rivals are excluded by *absent* instructions, which is why they are listed: a priority queue needs an extract-min — there is none; A\* needs `h` added to the label — the only distance term computed is compared against the generation counter, never added; a "seeded at the goal" reading is excluded by the callers, which pass the unit's own `pos+0`/`pos+1` as `src` and the ordered cell as `dst`)

**Original status.** ● active

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/)

### MOVE-COST-002

**The step costs, read out of the image, and there is no distance estimate at all.** For a mover whose `vt+0x20()` returns **1** (movementType 1, the ordinary ground domain): a **straight** step adds `cost[dst]`, a **diagonal** step adds `cost[dst] + (cost[dst]>>1)`. For **every other** mover: flat **2** straight and **3** diagonal, the cost plane not read. `cost[]` is the byte plane at `world+0` (`TERR-COST-052`), indexed at the **destination** cell, read **inline** (the byte load at `L06760`). The diagonal is the straight cost `× 3/2` **truncated** (the shift right by 1 at `L06761`), so with the shipped cost bytes (6…16, `TERR-COST-052`) a diagonal is 9…24, and on a cost byte of 1 the two would be equal; in the flat arm the ratio is exactly 3:2. All arithmetic is 16-bit (the 16-bit add) and the improvement test is unsigned strictly-less (the 16-bit compare at `L06762` with the skip on not-below at `L06763`), which is what makes `0xffff` the unlabelled marker. **No heuristic term exists**: the Chebyshev distance `D = max(\|Δx\|,\|Δy\|)` is computed once (`L06764…L06765`) and used **only** to size the iteration budget (`MOVE-TERM-003`). The same four arms appear in seven further places (`MOVE-SEARCH-001`, `MOVE-ROUTE-004`) with identical constants

**Confidence.** High (each of the four arms, the `SHR`, the operand widths and the compare are transcribed instructions; the rivals are excluded by the listing rather than by ratio-fitting — a table lookup or an FPU √2 would need instructions the arm does not contain, and "no straight/diagonal distinction" is refuted by the two separate arms)

**Original status.** ● active

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/)

### MOVE-TERM-003

**Termination: three exits, one of them a generation budget — and no node budget, no frontier bound, no closed-set limit.** Tested once per generation, *before* the generation runs (`L06766`, `L06767`, `L06768`): (a) `label[goal] != 0xffff`; (b) the current frontier is empty; (c) the generation counter (`+2` per iteration, one iteration = both half-passes) reaches the budget. Budget = `scalar + D` for the n×n arms and `max(scalar, D>>2) + D` for the 1×1 arms, `scalar` being `StaticScanAhead` (5) or `DynamicScanAhead` (3) per `MOVE-PARAM-006`; the static 1×1 arm substitutes a flat **1000** when `actor[5]->[0x28] == 0` *and* the goal's own footprint is free on the static plane (`L06769`). *(Amended 2026-07-31 by EXP-0072, which named the term: `actor[5]` is `actor+0x14`, **the owning `Player`**, and `Player+0x28` is **0 exactly when a human participant owns the unit** — 1 or 2 for every scenario-authored owner, the constructor's own default being 1 (`UNIT-OWNER-009`). **So the flat 1000 is the NORMAL budget for a player-ordered move, not a rare case**, and an engine that implements only the two computed forms gives a human player's units `max(scalar, D>>2)` rings of slack where the game gives them a thousand generations — a long land detour round an obstacle fails and the unit does not move. Read at instruction level: `L06770`…`L06771` load the actor's field at `+0x14` (the owning `Player`), then that `Player`'s field at `+0x28`, and jump to `L06772` when that is non-zero — the same block uses the actor's field `+0x154` for the mover and calls the actor's `vt+0x1c`, the footprint getter. The second half is an `n×n` scan of the **goal** footprint against the mover's own domain mask `mover+0x5` on the static plane `world+0x10000` (`L06773`…`L06774`), `n` being `tokenSize`; any masked cell jumps past the override, and `n <= 0` skips straight to it. The store is the 32-bit write of 0x3e8 at `L06775`, into the same slot the two computed forms write at `L06776`/`L02029`.)* Since one generation advances the wave by one ring, the slack over the straight-line distance is only `max(scalar, D>>2)` rings — a detour costing more than that fails, and **the search itself** then substitutes a nearby goal and extracts to it (`R0392` / `R0435`, `MOVE-ALT-018`…`022`). *(Amended 2026-07-31 by EXP-0068: this clause said "the **caller** then substitutes". It does not — both pickers are called from inside `R0053`'s own tail and from nowhere else in the image, and no caller ever sees a substitute; `claims/retracted.md`.)* **The frontier arrays are unguarded**: the push loads the count, stores into the x and y arrays at that index and increments the count, with no comparison (`L06777…L06778`), the count is a u16, and each array is 4096 bytes — the four `0x1000` immediates in the whole driver are all the x→y stride, not a bound. Exit (a) also means the route is taken from the **first** generation in which the goal acquires any label, so with unequal step costs the label is not necessarily minimal when the search stops

**Confidence.** High (the three tests, both budget forms, the 1000 override and the unguarded push are named instructions; "no capacity test" is an immediate scan of the driver and of all five relaxation routines) / Medium (that a >4096-cell generation is therefore reachable on a shipped map: the overflow follows from the code, but no shipped-map census of generation widths was run)

**Original status.** ● active (amended)

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/), [EXP-0068](../experiments/EXP-0068-alt-goal/)

**Amended.** The caller-substitution clause is partially retracted. Search selects its own substitute; MOVE-ALT-018 through MOVE-ALT-022 and retracted.md record the correction. Budget arithmetic stands.

### MOVE-ROUTE-004

**The route is extracted by walking the label field downhill from the goal, and its tie-break is asymmetric and favours the last straight step.** `R0438` (dynamic list) and `R0437` (static list) are the same code: start at the cell the search terminated on, and repeatedly choose the 3×3 neighbour minimising `label[nb] + stepcost(current cell)` — the *same* four cost arms as `MOVE-COST-002` — until the seed cell is reached. The **straight** arm accepts on `<=` (the compare at `L06779` and the above-skip at `L06780`) and the **diagonal** arm only on `<` (`L06781` / the not-below skip at `L06782`), over the scan order Δx = −1,0,+1 (outer) × Δy = −1,0,+1 (inner), **the centre included** — so among equal candidates the *last straight one in scan order* wins, a diagonal never displaces an equal straight, and the walk is not stable under a relabelling that only permutes equal costs. Guards: a candidate must be labelled (a compare with `0xffff`) and inside `8 ≤ x ≤ W+8`, `8 ≤ y ≤ H+8` (`world+0x50000` = W, `+0x50004` = H) — **the only rectangle test in the whole search**; passability is *not* re-tested and there is **no corner rule**, so a diagonal step between two blocked cells is legal. Output: a doubly-linked list of 12-byte pooled nodes (`+0x00` toward the goal, `+0x04` toward the unit, `+0x08` the packed cell `(y<<8)\|x`), built goal-first, with the unit's own cell unlinked before returning; a walk longer than **1000** steps discards the whole list (`L06783`) — reachable, because the centre is a candidate

**Confidence.** High (both compares, the scan order, the bounds, the node layout and the 1000 cap are named instructions)

**Original status.** ● active

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/)

### MOVE-PLANE-005

**Two searches over two planes, two routes on the actor — and the static search is structurally unit-blind.** `staticFlag != 0` reads the **static** block plane `world+0x10000` and writes the static list at `actor+0x160/0x164/0x168/0x16c/0x170` (tail/head/count/free/pool); `staticFlag == 0` reads the **dynamic** plane `world+0x20000` and writes the dynamic list at `actor+0x17c/0x180/0x184/0x188/0x18c`. Both use the same mask byte `mover+0x05`, and the reason the split works is the mask's own bits: `0x41` = terrain bit 0 **+ ground-occupancy bit 6**, `0x44` = static-object bit 2 + bit 6, `0x82` = border bit 1 **+ air-occupancy bit 7** (`TERR-PASS-051`) — the occupancy half of every mask is inert on the static plane, because bits 6/7 are never set there. Writers of those two bits, image-wide: `R0453`, `R0058`, `R1332`, `R0257`, `R0046` — **every one writing `world+0x20000`**

**Confidence.** High (the two plane displacements are in the relaxation routines' `TEST` operands; the mask bit meanings are `TERR-PASS-051`'s, re-derived here from the same three constants) / Medium (the writer list as an *enumeration*: two instruments, each with a blind spot the other covers only partly — `EnumRefs disp:10000`/`disp:20000` (60 hits, 19 owners, 0 orphan) cannot see a write whose address was folded into a register first, which is exactly how `R0257` writes; and `EnumRefs re:` over the four read-modify-write forms (`OR …,0x40/0x80`, `AND …,0xbf/0x7f`) **misses `R0046` entirely**, which clears with a load/AND/store pair. Both routines are in the list because they were read, not because a scan found them)

**Original status.** ● active

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/)

### MOVE-PARAM-006

**The six `[Path Finding]` scalars — the engine's own names, and the shipped file agrees with the code defaults 7/7.** `R0116` reads them through the `.reg` int getter `R0452(section, key, default)` (`RES-CODE-020`) into `world+0x585b4` `StaticScanAhead` (default **5**), `+0x585b8` `DynamicScanAhead` (**3**), `+0x585bc` `StaticRefreshRate` (**16**), `+0x585c0` `DynamicRefreshRate` (**32**), `+0x585c4` `DynamicByStaticLookup` (**3**), `+0x585c8` `StaticIsntNeeded` (**5**); `SpeedMultiplier` → `+0x58db4` (**8**) is the same section and is `TERR-MOVE-056`'s. `world.res:data/map.reg` `[Path Finding]` ships **all seven equal to those defaults, identically in the EN and RU roots** (`evidence/pathfinding-params.md`; `[Scanning] ScanShift = 7` beside them). Names are clamped to 15 characters in the file, matching the 15-char compare `TERR-COST-052` names

**Confidence.** High (the seven immediates and their store addresses are in one function's listing; the file values are the file's own bytes under EXP-0032's framing, and the 7/7 agreement is an independent check on the name→default pairing — a mis-attributed default would disagree with the file. The reading is from `tools/regcorpus`. It was **not** taken from `tools/regdump`, which until 2026-07-30 walked the pre-EXP-0032 framing — an 8-byte-shifted window pairing each key with the *next* record's value — and read this very section as `StaticScanAhead = 3`. That tool now carries the corrected framing and reads 5, agreeing with `regcorpus` and with the code default; the hazard is recorded because the figure was chosen to avoid it, not because it still exists)

**Original status.** ● active

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/)

### MOVE-CLAIM-007

**There IS a reservation: a unit marks the cell it intends to enter, before it enters it.** `R0054` reads the next route cell from the dynamic list's tail into `mover+0x06`, and when that differs from `mover+0x80` it calls `R0046(mover+0x80)` to drop the previous claim and **`R0257(mover+0x06)` to OR the mover's own occupancy bit over the n×n footprint of a cell the unit has not yet reached**, then stores it in `mover+0x80`. Every other unit's *dynamic* search therefore treats that cell as blocked (`MOVE-PLANE-005`), and the static search does not. The claim is released together with the vacated cell when the unit lands centred on the new one (`R0039`: `mover+0xaa/0xac/0xa8 = 0`, `R0046(mover+0x70)`, `mover+0x80 = 0`, `mover+0xa6 = 0`). The set/clear pair is **asymmetric**: `R0257` ORs unconditionally, while `R0046` skips any cell inside the unit's *current* footprint — so releasing a claim never un-occupies the ground the unit is standing on. So the answer to "is there any structure recording a unit's intended position" is **yes, `mover+0x80`**, and the shared dynamic plane is the channel; `mover+0xa6` records the cell whose footprint currently carries the bits

**Confidence.** High (each call, its argument, the guard forms and the two mover fields are named instructions, `evidence/rom-move-excerpt.md` §6) / **Unknown** (whether a unit's own outstanding claim can block its own next dynamic search: the search lifts only its *current*-cell occupancy — `R1332` at `L06784`, restored at `L01945` — and never touches `mover+0x80`, but no reachability argument was built for the case where the claimed cell is not on the new route)

**Original status.** ● active

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/)

### MOVE-WAIT-008

**When the next cell is occupied the unit waits, facing it — and it neither pushes, swaps, nor steps aside. At step time there is no collision test at all.** In `R0178`, on a tick where a dynamic route is needed (list empty, or the per-tick counter `mover+0x78` over `DynamicRefreshRate`) *and* the next static waypoint is **exactly Chebyshev-1 away** *and* its n×n footprint is blocked on the dynamic plane, `R0402` decides: nonzero ⇒ the unit turns toward the cell (`R0089` → `R0056`), sets `actor[0x56]+8 = 10`, and **returns — no step, no re-search**; zero ⇒ fall through to the dynamic re-search `R0055`. When the waypoint is farther than 1 the test degenerates to cell 0 (a border cell, always blocked) and `R0402` answers for that cell instead, so the wait is reachable only in the adjacent case. Nothing anywhere in the movement path displaces another unit. And `R0054` — the routine that actually steps — **tests no plane**: it claims the next cell and moves regardless of what is in it, so two units whose dynamic routes were both computed before either claimed can walk into the same cell

**Confidence.** High (the caller's control flow: the Chebyshev-1 gate, the footprint test's plane, the turn-and-return, and the absence of any plane test in `R0054`) / Medium (**which** blocker yields wait vs re-search: `R0402`'s arms were decompiled — the cell-record lookup `R0035`, the actor list at the cell, the match on the other mover's `+0x70`, the `1`/`4` state byte — but its two helpers `R0040` and `R0019` were not read, so the verdict table is not pinned)

**Original status.** ● active

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/)

### MOVE-TICK-009

**The per-tick update imposes no priority, so insertion order decides which actor gets a contested cell.** Simulation actors share `vt+0x18 = R0037`, which reaches the order state machine. `R0427` walks one pooled linked list, invokes that virtual on each active element, and contains no sort, priority key or reordering. Its `[0x21,0x3f]` typeID test lies only in the `+0x54==0x10` teardown arm; it is not what proves the update elements' class, and it does not contain every Human because zero-mode map Humans can retain lower table typeIDs (`PARTY-M20-031`). A movement claim is written immediately, so the first actor visited claims the cell and later searches see it. The list is AddTail-only: map load inserts party then type-6 record order, death unlinks, spawns append, and save/load regroups by Player (`MOVE-TICK-013`..`017`)

**Confidence.** High for the virtual, whole loop and insertion-order writers; the corrected typeID scope is independently witnessed by original saves

**Original status.** ● active (corrected)

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/), [EXP-0055](../experiments/EXP-0055-tick-order/), [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/)

**Amended.** The teardown typeID range as proof of every list member being Human is partially retracted. Update order and list mechanics stand; PARTY-M20-031 and retracted.md record the correction. EXP-0515 also refutes "map load inserts party then type-6 record order" for the members the join walk places on the new-mission path: the type-6 records are inserted first and those members after the binder (High, `MOVE-114`, [`retracted.md`](retracted.md)). A companion inserted by AddHero between missions stays open (Medium).

### MOVE-STEP-010

**Movement is sub-cell, 1/256 of a cell per axis, and the position is a byte pair per axis.** The unit's position object is `actor[4]` = `*(actor+0x10)`: `+0x00` x cell, `+0x01` y cell, `+0x02` the packed cell `(y<<8)\|x`, `+0x04` x fraction, `+0x05` y fraction, **`0x80` = centred**. `R0047` treats `(cell<<8)\|fraction` as one 16-bit value per axis and adds the signed per-axis step `mover+0xb0` / `mover+0xb1` (`TERR-MOVE-056` derives their magnitudes) to it, increments `mover+0xac`, and when the *cell* byte changes calls `R0058` on the ~~new~~ **old** cell (`L06785` loads `word[pos+2]` before `L02039` rewrites it; **corrected by [EXP-0093]**, `claims/retracted.md`) plus `R0050` after the write; when `mover+0xac >= mover+0xaa` it snaps both fractions back to `0x80`. So a cell transit takes `mover+0xaa = ceil(256 / max\|step\|)` ticks and ends exactly at the cell centre, and a unit is *between* cells for most of its life — the packed cell is the near cell, not a rounded one. `R0178` refuses to do anything but continue the transit while either fraction is not `0x80`. The steps are **signed** bytes (`L06786`/`L06787`, sign-extending loads), and one direction-keyed special case (`L06788`/`L06789`, comparisons of `mover+0xae` with 0x1 / 0x5) subtracts 1 from the combined x at `L06790` when both fractions land on 0

**Confidence.** High (the 16-bit add, the snap condition, the boundary-crossing test and the field offsets are named instructions; the argument of `R0058` was a name, not a read, and fell)

**Original status.** ● active (amended)

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/)

**Amended.** The boundary hook receives the old cell, not the new cell. That argument clause is partially retracted; MOVE-CLOCK-032 and retracted.md record the correction. The other step relations stand.

### MOVE-SPEED-011

**Speed does not feed the search.** The search reads exactly two things about the mover — the footprint side `vt+0x1c()` and the mask `mover+0x05` — plus the cost plane; no speed term enters any label. Speed enters only *after* a route exists: `R1088` computes the per-step duration (`TERR-MOVE-056`: `SpeedMultiplier × unitSpeed`, height-tilted, divided by the two cells' mean cost), and `R0039` recomputes `mover+0x72 = (actor+0x8c << 3) / cost[cell]` on each cell entry for a movementType-1 mover (0 for types outside 1..3, the raw class speed for 2 and 3). So two units of different speed choose the **same** route and traverse it at different rates, and the per-class speed source is `Data.bin` (`TERR-MOVE-057`/`058`) — **except under a group order **issued in formation**, where the rate is the group's slowest member's `Speed` and the two units traverse it at the SAME rate** ([EXP-0093], `MOVE-GROUP-030`; the *when* is `MOVE-GATE-035`, and because nothing clears the term the exception outlives the order that set it — `MOVE-GROUP-037`). The full composition is `MOVE-RATE-029`

**Confidence.** High (the search's whole mover interface is two virtual calls and one byte; the speed sites are named instructions)

**Original status.** ● active

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/)

### MOVE-REFRESH-012

**When a route is recomputed — per cell for the dynamic one, per target change for the static one, and nothing is staggered.** (a) The **dynamic** route is destroyed on **every completed cell transit** — `R0039` frees the whole list the moment the unit lands centred — so a dynamic search runs at least once per cell stepped. (b) It is also recomputed when `mover+0x78`, incremented once per `R0178` tick and zeroed on re-search, exceeds `DynamicRefreshRate` (32) — the path that matters while a unit is waiting or turning rather than stepping. (c) The **static** route is recomputed only when the ordered target differs from `mover+0x74`, or when `mover+0x09` — the number of dynamic re-searches since the last static one — exceeds `StaticRefreshRate` (16); a static recompute frees the dynamic list. (d) The dynamic search does **not** aim at the final goal: `R0055` takes the static list's tail waypoint when it is more than `DynamicByStaticLookup` (3) cells away, else the node `DynamicByStaticLookup+1` further along, and uses the final goal `mover+0x76` only when the static list holds `StaticIsntNeeded` (5) nodes or fewer; a static waypoint is popped once the unit is within `DynamicByStaticLookup` of it. Every counter is per-unit and reset on use, so **there is no stagger and no shared phase** — two units ordered in the same tick re-search in the same ticks. Nothing invalidates a stored route on terrain change or on target *movement*: only a new target cell (c), a completed step (a), the tick counter (b), and a blocked adjacent waypoint via `MOVE-WAIT-008`

**Confidence.** High (each counter, its threshold, its increment and its reset are named instructions, and the target-selection arms are `R0055`'s own three branches) / Medium (calling `mover+0x78` a *stuck* counter: it increments on every tick of an unfinished move order, stepping or not, so it is a re-search period that a stalled unit reaches sooner — no separate give-up counter was found, and `mover+0x98`/`mover+0x90`, set on "no static route" and "arrived", are read but not traced to an effect -- *`mover+0x98` traced by [EXP-0099]: it is set on an empty route by `R0043` at `L00161` and consumed only by `R0016`'s epilogue, which cancels the pending order and re-acquires within reach; see `AI-ROUTE-045`. `mover+0x90` traced by [EXP-0386]: it is a plain boolean, "actor already stands on the route's own resolved endpoint", not a route-list-emptiness test; see `MOVE-072`*)

**Original status.** ● active

**Evidence.** [EXP-0054](../experiments/EXP-0054-unit-movement/)

### MOVE-TICK-013

**The tick loop's container is a pooled doubly-linked list embedded in a 0x20-byte manager at `[L00240]`, its walk is head→tail, and its order is pure insertion history — the only insert operation that exists for it is AddTail.** Layout (ctors `R0120` → `R1333`): vptr `L06791` @+0, then an MFC-shaped CObList @+4 — head @+8, tail @+0xc, count @+0x10, node free-list @+0x14, CPlex block chain @+0x18, blockSize 10 @+0x1c. Nodes are 12 bytes (`next`@+0, `prev`@+4, `element`@+8), pooled in CPlex blocks of 10 (`R0302` → `R0540(…,10,0xc)`), freed nodes pushed LIFO on the free list. The iterator (`R0028`/`R1334`) advances **before** the element is processed, which is what makes the teardown's self-removal safe. The same class serves three more populations: the **dead list** at `*(world+0xc)`, a **per-player list** at `*(player+0x20)`, and the 0x48-byte **group** objects on `player+0x24`. Mutation surface, complete: insert = `R0412` (AddTail; NewNode has exactly two callers, the other an inlined tail-append on a caller-local list in `R0109`); remove = unlink (`R0413`/`R1335`) via the by-value wrappers; RemoveAll; the serialize pair. **No AddHead, no InsertBefore/After, no sort and no compare-driven reorder exists in the family**

**Confidence.** High (every field offset, both ctors, the node functions and the iterator are named instructions, `evidence/rom-tick-excerpt.md` §2–3; the array rival is killed by the pointer walk and the unlink, the AddHead rival by `callto:R0302` = 2 callers, the id-keyed rival by the iterator never reading `actor+4`. Enumeration instrument: `EnumRefs callto:` on every family function — 0 hits in orphan/undisassembled code)

**Original status.** ● active

**Evidence.** [EXP-0055](../experiments/EXP-0055-tick-order/)

### MOVE-TICK-014

**What fixes the order at map load: the creators' call sequence — campaign party first, then the map's type-6 records in ascending record order — and the list is per-*ticking-thing*, not total.** Every creator inserts through one routine, `R0411([L00240], actor)` (AddTail + id, six callers, each loading the global immediately before the call): the party placement `R0065`, the map type-6 walk `R0151` (index ascends 1..count over the map's record array; the record's own unit id is stored to `actor+8` at `L06792`), a trigger-parameter spawner `R0062`, a roster walk `R0432`, spawn-by-name `R0066`, and sack creation `R0003` — one of whose callers is `R0037`, the actor tick itself, so a mid-tick death-drop appends to the very list being walked. Humans, Units and Sacks are members; **no Building insert into this list was found**, though buildings draw ids from the same allocator (direct `R0936` callers with no list insert: `R0941`, `R1336`, `R1337`) — which is how `SAV-ID-015`'s ids interleave buildings 2..19 between the hero and the units while the tick walk carries no buildings

**Confidence.** High (the six-caller enumeration: `callto:R0411`, 0 orphan, each site's `ECX=[L00240]` shown by `refto:L00240` adjacency; the unit walk's ascending index and the `actor+8` store are named instructions) / Medium (that the record array the walk indexes = the `.alm` type-6 **file** order: the accessor `R1338` was not chased into the map loader — `SAV-ID-015`'s 35/35 file-order ids corroborate on one map) / Medium ("buildings never enter this list" — an absence over the enumerated insert surface, blind to an insert through a wrapper this experiment did not identify; whether buildings tick from another container — `world+0x2c`, `*(world+0)/(+4)/(+8)` — is open)

**Original status.** ● active (contested)

**Evidence.** [EXP-0055](../experiments/EXP-0055-tick-order/)

**Amended.** Building membership remains contested: no Building insert was found over the named insert surface, but untraced wrappers and other tick containers remain open. The record-array to file-order link also remains Medium. The party-first clause is contested by EXP-0514 (`TRIG-MAPORD-105`): the session start `R0512` loads the map, spawns the type-6 records and runs the binder without calling the placement walk `R0065`, whose session-path caller is the join arm of the command executor. Two readings stand: the party is inserted first (this card), or after the map's records, when the join command executes. Neither experiment read when the join command executes relative to the session start, and the walk's caller `L13863` is unclassified. EXP-0515 refutes the party-first clause for the members the join walk places on the new-mission path, High ([`retracted.md`](retracted.md)): the join command is sent after the session start has run the spawner and the binder, and `L13863` is the watchdog arm (`MOVE-114`). A companion that an AddHero command inserted between missions, through the drain caller `R0192` whose callers are not enumerated, would still precede the type-6 records; that case stays contested at Medium. Building membership stays contested.

### MOVE-TICK-015

**How the order changes during play — and death leaves no hole.** (a) The teardown arm (`actor+0x54 == 0x10`) removes the actor's node by **unlink** (`L01926`) — nothing is compacted, swapped from the tail, or replaced; every survivor keeps its relative position — then AddTails the actor to the **dead list** `*(world+0xc)` (`L01927`), where `R0860` runs corpse decay per tick. (b) New spawns/summons append at the **tail**. (c) **Leaving the map** (into a building/transport, `R0121`) unlinks from the tick list and sets `actor+0x4c` bit 3; **returning** (`R0122`/`R0123`, "Unit can't return to map — no free place") re-AddTails — the returning unit moves to the **back** of the walk, losing its old contention position. (d) An **owner change** (`R0064`) moves the actor between the players' own lists and groups but never touches the tick list — its walk position survives a conversion. So mid-game, the walk order = spawn order, minus the dead and the garrisoned, plus returners and newcomers at the tail

**Confidence.** High (each arm is a raw listing with the `ECX` chain shown, `evidence/rom-tick-excerpt.md` §3, §5, §6; the replace-in-place and periodic-resort rivals are excluded by the complete mutation surface of `MOVE-TICK-013`)

**Original status.** ● active

**Evidence.** [EXP-0055](../experiments/EXP-0055-tick-order/)

### MOVE-ID-016

**The runtime id (`actor+0x04`, `SAV-ID-015`'s) is a lowest-free-bit bitmap allocation — creation-ordered only until the first death decays, and never the walk order.** `R0936` scans the bitmap `L06793` for the first clear bit (id 0 pre-marked by `R1339`, which runs at world construction and again at save-load start), marks it, and `R0411` stores it at insert time. `R0868` clears a bit; its actor-side caller is the **corpse decay**: at stage 5 (`actor+0x94` reaching −600) `R0867` frees the id and zeroes `actor+4` (`L06794`/`L04100`) — which is why long-dead units serialize with id 0, and why the **next** spawn after a decay reuses the lowest freed id while still appending at the tail. The id space is shared with non-ticking placeables (three direct allocator callers with no tick-list insert). On load the id is read back from each actor's head and **re-marked** (`R0950` at `SetAt`… listing §7), so ids round-trip exactly

**Confidence.** High (allocator, set, clear, reset, the insert-time store, the decay-time free and the load-time re-mark are named instructions; `refto:L06793` = exactly 4 touching functions)

**Original status.** ● active

**Evidence.** [EXP-0055](../experiments/EXP-0055-tick-order/)

### MOVE-TICK-017

**The walk order is NOT preserved across save/load — the saved stream never carries it, and the loader rebuilds it grouped by player.** The world Serialize (`R0414`) stores: the players (each Player storing its groups inline, each group WriteObject-ing **its own actor list head→tail** — `R0415` → `R1340` → `R1143` → `R1341`), then the dead list; the tick list appears nowhere in the store arm's exhaustively-walked call sequence, and EXP-0048's byte-level tiling of four saves corroborates (393/393 objects under players/groups/dead, ≤ 70 bytes unattributed — no room for a 400-word reference list). On load, Player::Serialize rebuilds `*(player+0x20)` from its groups (groups in list order, actors in group order, `L01794`), and the world's load arm then rebuilds the tick list: **for each player in players-manager order, append `*(player+0x20)` head→tail, skipping off-map (`+0x4c` bit 3) actors** (`L06795…L01793`). Consequence, joint with `MOVE-TICK-009`: a fresh map walks global creation order (players interleaved as the records interleave); a loaded game walks player-by-player, group-by-group — relative order within one group survives, the cross-player interleave does not — so **which of two contending units of different players moves first can change across one save/load cycle**

**Confidence.** High (both rebuild loops and the store sequence are raw listings, §7–§8; the round-trip rival is killed by the absence of the list from the store arm *and* by the rebuild's existence) / Medium (the regrouping's *observable* consequence on a real contention: derived from the loops, not yet demonstrated in a live session — the named follow-up)

**Original status.** ● active

**Evidence.** [EXP-0055](../experiments/EXP-0055-tick-order/)

### MOVE-ALT-018

**The substitution is not the caller's — `R0053` does it itself, in three branches, and `altTarget` is a target *actor*, not a cell.** The two pickers have **two and one call sites, all three inside `R0053`** (`L01913`, `L01915`, `L01914`) and are reached from nowhere else in the image. They run in the tail entered when the goal is still unlabelled (`L06796` compares the 16-bit label with `0xffff`, `L06797` branches on equal). Which one runs is decided by the `staticFlag` argument at `L06798` and the `altTarget` argument at `L06799`: **static** → picker A around the requested cell with bound `(D>>2) + 4` (a shift right by 2 at `L01963` and an add of 4 at `L01964`), `altTarget` **ignored**; **dynamic** → `altTarget != 0` ? picker B(`mover`, `altTarget`) : picker A with bound **8** (the literal 8 pushed at `L06800`). The answer is fed straight to the ordinary route extraction with the substitute as `(dstX, dstY)` — `R0437` at `L01941`, `R0438` at `L01943` (`MOVE-ROUTE-004`) — and a zero answer means no route (`L06801…L01944` frees the static list through `R0439(actor+0x15c)`; the dynamic branch simply returns). **`altTarget` is an object**: picker B reads `arg2+0x10` as the position record (`MOVE-STEP-010`) and calls `arg2->vt+0x1c()` for its footprint side. The only caller that passes a non-zero value is `R0043`, the go-to-actor order, via `R0055(mover, targetActor)` at `L01901`; `R0178` passes 0 at both `L06802` and `L01717`, and `R0043`'s own **static** call at `L01894` passes the actor into an argument the branch above discards

**Confidence.** High (the two `EnumRefs callto:` sweeps — 2 hits / 1 owner and 1 hit / 1 owner, 0 in orphan or undisassembled code — bound the surface; the three branch tests, both bounds and the two extraction calls are named instructions, `evidence/rom-altgoal-excerpt.md` §3. The rival "a caller substitutes" is excluded by the absent call sites, not by plausibility; the rival "`altTarget` is a cell" by the two dereferences)

**Original status.** ● active

**Evidence.** [EXP-0068](../experiments/EXP-0068-alt-goal/)

### MOVE-ALT-019

**Picker A (`R0392`): expanding square rings around the requested cell, and the whole ring is scanned before the best in it is taken.** `__thiscall(world, mover, unused, cell, limit)` — the second argument is `0` at both call sites and is never read. Rings `r = 1, 2, 3, …`; for each `i = -r..r` it probes the four cells `(x+i, y+r)`, `(x+i, y-r)`, `(x+r, y+i)`, `(x-r, y+i)` — the four `LEA`/`ADD` forms at `L06803`, `L06804`, `L06805`, `L06806` over the packed cell `(y<<8)\|x`, with `i` incremented at `L06807` and compared with `r+1` at `L06808`, and `i`'s start decremented once per ring at `L06809`. Each probe is **one instruction**, a 16-bit load from the label plane at offset 0x30000, and the keep test is an unsigned compare against the running minimum that keeps only a strictly smaller value — the **strict** minimum. The ring is scanned to the end; if it produced any labelled cell at all the loop exits (`L06810` compares with `0xffff`, `L06811` branches on carry), otherwise `r` grows while `r+1 < limit` (the signed compare at `L06812`, branch below at `L06813`), and `limit <= 1` skips everything (`L06814`). The **centre is never probed** (`r` starts at 1), and correctly so — the branch is only entered when the goal is unlabelled. Return: the packed cell, or **0** if no ring had one (`L06815`: a branch-free select of the packed cell unless it is the `0xffff` sentinel). **The mover's footprint does not change the scan**: `L06816`, a compare of `mover->vt+0x1c()` with 0x2, selects between two copies of the same code whose 199-byte inner bodies differ in exactly **4 bytes**, the operands of a swapped register-restore pair, with no jump displacement among them (`evidence/image.txt`)

**Confidence.** High (every instruction above is transcribed from a full listing of a 595-byte routine, not from an enumeration, and the routine's bytes are hashed in `evidence/image.txt`. Check **D1** (`evidence/ring.txt`) re-executes the four side formulas and confirms they cover the Chebyshev ring of radius `r` exactly for `r = 1..8`, corners twice and nothing else duplicated — a sign error would break it. The rivals die on absent instructions: a spiral or a raster needs a different index recurrence than `±r`/`±(r<<8)`; a distance-sorted list needs a sort, and there is none; "stop at the first labelled cell" is excluded because the exit test sits **after** the inner loop, not inside it)

**Original status.** ● active

**Evidence.** [EXP-0068](../experiments/EXP-0068-alt-goal/)

### MOVE-ALT-020

**Picker B (`R0435`): the contact ring around the target actor, entered where the line between the two actors crosses it and walked in both directions.** `__thiscall(world, mover, target)`. The box is `x ∈ [tx - nM, tx + nT]`, `y ∈ [ty - nM, ty + nT]`, built from four `vt+0x1c` calls — two on the **target** (`L06817`, `L06818`, giving `nT`) and two on the **mover** (`L06819`, `L06820`, giving `nM`) — and it grows by one in each direction per ring, for **8 rings** (`L01916`, the ring count compared with 0x8 with a greater-or-equal exit, expansion at `L06821…L06822`), the loop continuing only while nothing has been found (`L06823`, a compare with `0xffff` and a branch on equal). The entry cell is where the straight line between the two actors' **fine footprint centres** meets that box: `R0436` returns a 16-way direction code, folded to a quadrant by `(dir + 2) >> 2 & 3` (`L06824…L06825`), which selects one of four edges through the inline jump table at `L06826`; the crossing is `slope × fine_edge + intercept` → `__ftol` → `SAR 8`, with `slope` = `dy/dx` or `dx/dy` per quadrant (table `L06827`, two distinct arms) and a divide-by-zero guard that turns `dx == 0` into `1` (`L06828 FCOM float ptr [L06829]` = `0.0` / `L06830 FSUB float ptr [L06831]` = `-1.0`). Two walkers then run the perimeter from that cell in opposite directions, stepping by the 12-entry `(dx,dy)` byte table at `world+0x54186` — `(+1,0) (0,+1) (-1,0) (0,-1)` **three times**, written only by the world constructor (`EnumRefs disp:54186/54187/54188` → 2 / 2 / 1 hits, 0 orphan; the loop register is zeroed from `L06832`, the only write to it (full, 16-bit or 8-bit) before its epilogue) — turning at a corner through the two 12-entry jump tables at `L06833` / `L06834`, each of which is four handlers repeated three times. Each probe is again **one** label-plane read (`L06835`, `L06836`) kept by a signed compare against the running minimum with a greater-or-equal skip, the strict minimum. Return: the packed cell `(y<<8)\|x` (`L06837`), or **0** (`L06838`, zeroing the result)

**Confidence.** High (the box, the 8-ring bound, the two walkers, the direction table, the four jump tables and both label probes are named instructions over a full listing of a 1559-byte routine whose bytes are hashed; the jump tables are read out of the image, not off the decompiler, in `evidence/image.txt`. Check **D2** (`evidence/box.txt`): at ring 0 the box perimeter is exactly the set of mover origins whose `nM × nM` footprint touches the target's `nT × nT` footprint without overlapping it, for every `(nM,nT)` in `1..4` — so "contact ring" is measured, not asserted) / Medium (the entry-cell arithmetic: the four FPU arms were read and their *shape* — line, slope, intercept, truncate — is transcribed, but the quadrant→edge mapping was not checked cell by cell against a worked example, and no consumer of the entry cell's exact value was tested. What would lift it: re-execute the four arms over a grid of `(dx,dy)` and confirm the entry cell always lies on the box)

**Original status.** ● active (amended)

**Amended.** The entry clause is narrowed: for a mover footprint of 2 to 4 with bearing codes 0, 1, 10 or 11 the entry cell can lie outside the box and one walker then leaves the ring (`MOVE-098`); the edge selection and the entry arithmetic are measured in `MOVE-097` and `MOVE-098`. `retracted.md` records the narrowing.

**Evidence.** [EXP-0068](../experiments/EXP-0068-alt-goal/)

### MOVE-ALT-021

**What a candidate is tested against, and what the choice is measured from — and the answer to both is the label plane, so the substitute is the cell cheapest to reach *from the mover*, not the cell nearest the click.** Across **both** pickers read in full, the only world displacements are `+0x30000` (the u16 label plane, `MOVE-SEARCH-001`) — eight sites in `R0392`, two in `R0435` — and `+0x54186/7` (picker B's direction table). Neither routine contains a passability predicate, a cost read or an occupancy test, and neither calls one: the complete callee set is the footprint virtual `vt+0x1c`, `R0862`/`R0863` (which read `pos+0/1` and `pos+4/5` and add `(side-1)<<7`, reaching only `R0165`/`R0166`), `R0436` (which calls only those two) and `__ftol`. So **the test is "this cell carries a label"** — the verdict of the wave that has just failed, which already folded in the plane it ran on, the mover's footprint and its mask (`MOVE-PLANE-005`, `TERR-PASS-051`) — and **the metric is the label itself**, the accumulated step cost from the mover's own start cell under `MOVE-COST-002`. Two consequences a consumer must implement. (a) The rival "the nearest free cell to the clicked cell, by Chebyshev, scanning outward" is **half right and half wrong**: the ring order *is* Chebyshev outward from the requested cell, but within a ring the whole ring is scanned and the cheapest-from-the-mover cell wins, and "free" is *not* passability — a perfectly passable cell that the wave's generation budget never reached carries no label and is invisible to the picker. (b) Substitution is **per-mover**, on that mover's own failed wave: on the dynamic branch the labels were laid over `world+0x20000`, which carries other movers' claims (`MOVE-CLAIM-007`), so a claimed cell is never labelled and can never be returned; on the static branch `world+0x10000` carries no occupancy, so a static substitute is chosen as if the mover were alone on the map

**Confidence.** High (an **exhaustive** read of two hashed byte ranges, which is a stronger instrument than any enumeration and has no orphan-code blind spot: the claim is not "no other reader was found" but "these are all the instructions". The three rivals are excluded by absent instructions — a distance-to-click metric needs a subtraction of the two cells and there is none; a distance-to-mover metric needs the mover's position, which picker A never loads; a fresh passability test needs the block plane, whose displacement appears nowhere in either routine)

**Original status.** ● active

**Evidence.** [EXP-0068](../experiments/EXP-0068-alt-goal/)

### MOVE-ALT-022

**The substitute is never written back — it is consumed by one route extraction and forgotten.** Over the whole tail `L06839…L06840` the only stores reaching the actor are the route list the extraction builds (`actor+0x160…` / `actor+0x17c…`, `MOVE-PLANE-005`) and `R0439(actor+0x15c)` freeing the static list when no substitute was found. `mover+0x74` (the last ordered static target) and `mover+0x76` (the final goal) are **not** touched, the order block `actor+0x158` is not touched, and the picker's return value never leaves `R0053` — it is pushed to the extraction call and to nothing else. The static branch does *measure* the substitute against the request — `R0167(substitute, requested)`, the Chebyshev distance of two packed cells — with a tolerance of **1**, raised to **2** when the requested cell has static-plane **bit 5** set (`L06841`, a byte test of the static plane at offset 0x10000 against 0x20, one of `TERR-PASS-051`'s 13 bit-5 sites) *and* the cell-record lookup `R1342(world+0x540b4, cell)` returns a record whose `+0x0c` is non-zero; exceeding the tolerance queues a UI message (`R1343` on the object `L00522`) and **does not cancel the move** (`L01952`, a zero test of the substitute, gates the extraction on the substitute existing, not on the tolerance)

**Confidence.** High (the tail is listed in full and hashed, so "no store" is an exhaustive read rather than an enumeration; the tolerance test, both constants and the message call are named instructions) / Medium (the run-time consequence — that the order still points at the original cell, so the unit re-runs the same failing search whenever `MOVE-REFRESH-012`'s counters fire, and substitutes again from wherever it now stands: this follows from the two rows but was not observed. `R1342`'s record and its `+0x0c` are named, not read)

**Original status.** ● active

**Evidence.** [EXP-0068](../experiments/EXP-0068-alt-goal/)

### MOVE-ORDER-023

**There is no multi-unit destination distribution anywhere: no formation, no offset table, no spread. A group move writes the same cell into every member's order block.** Every site that sets the move opcode is a byte store of the immediate 1 at offset 0x8 of an order block loaded from `actor+0x158` — **16 sites over 12 owners, 0 in orphan or undisassembled code**. Nine owners write one actor's order per call. The three that walk a container of actors are the group orders `R0024`, `R0025` and `R0154`, and only `R0025` takes a cell: `__thiscall(issuer, group, cell)`, reached from the group order machine `R0023` case 2 with the group order's own `+0x0a`. Its cell is loaded **once**, `L06842` loads it from the frame at offset 0x18 into a register, *before* the loop head at `L06843`; that is the **only** write to that register in the routine; the loop body contains no counter, no index and no arithmetic on it; and the store is the 16-bit write of that register into offset 0xa of the member's order block at `L00448` for every member the list walk reaches. The other two re-issue each actor's *own* already-stored cell (`order+0x00` at `L06844`, `order+0x0a` at `L06845`). So the spread a player sees when several units are sent to one tile is produced entirely by `MOVE-ALT-019`/`021` — each mover substituting independently against a plane that already carries the earlier movers' claims — and not by anything at order time

**Confidence.** Medium (the enumeration is over **one instruction form**: a move order set with the opcode in a register, through a store wide enough to span `+0x8` and its neighbour, or by copying a whole order block, is invisible to it — `EnumRefs re:` sees orphan code, so that half of the blind spot is closed, but the encoding half is not. What is High inside the row is `R0025` itself: the single `BP` load, its position before the loop head, and the absence of any other `BP` write are a full listing of a 226-byte routine, hashed in `evidence/image.txt`. What would lift the row: a `disp:158` sweep classified owner by owner, or locating the order block's class and enumerating its writers through the vtable)

**Original status.** ● partially retracted — **the headline is refuted by [EXP-0094]**, through this cell's own named blind spot: the per-member issue routine `R0301` sets the move opcode from a **register** (the byte store at `L01930` into offset 0x8 of the order block, the value 1 loaded into it at `L00604`), and on the formation arm of `R0108`/`R0157` it is called once per member with `target + (memberCell − centroidCell)`. Formation, offset table and spread test all exist (`MOVE-FORM-036`, `MOVE-GATE-035`, retracted.md). What survives untouched: the three routines this row read are still right about themselves

**Evidence.** [EXP-0068](../experiments/EXP-0068-alt-goal/), **[EXP-0094](../experiments/EXP-0094-group-rate/)**

**Amended.** The no-formation headline is partially retracted. MOVE-FORM-036 and MOVE-GATE-035 establish the counterexample; retracted.md records its scope. The three selected routines retain their bounded local conclusions.

### MOVE-DOM-024

**The movement domain is one byte with six consumers, and passability is only one of them.** `actor+0x4a`, read through slot **`+0x20`** of all three simulation-actor vtables (`L00001`/`L00002`/`L00003`, whose `+0x20` all hold `R1344` = a byte load of `+0x4a` and return; `TERR-MOVE-055`). What the code branches on it for: (a) the **block mask** — `R0471`, `MOVE-DOM-025`; (b) **which occupancy bit** the mover sets and tests on the dynamic plane — `R0058`/`R1332`/`R0257`/`R0046`, `< 3 → 0x40`, `== 3 → 0x80`; (c) the **step-cost arm** — `== 1` reads the cost plane, every other value takes flat 2/3 (`MOVE-COST-002`), at nine sites across the driver, the three relaxation routines and the two route extractions; (d) the **speed** — `R1088` skips the height tilt and the cost divide for `!= 1`, and `R1345` (read whole) returns `(class speed << 3)/cost[cell]` for 1, the **raw class speed** for 2 and 3, and **0** for 0 and for anything above 3; (e) the **AI's target choice** — a candidate whose domain is 3 counts one cell farther away when the decider's is not (`L00225…L00226`, `AI-ACQUIRE-002`), which is the image's own word for *flier* and the only attestation of that name that does not come from the mask; (f) the **corpse** — the actor tick reads the byte directly at `L04425` and, for `> 1`, stores `0xfc18` (−1000) into the decay counter `actor+0x94` (`L04398`), which the next test at `L01925` turns into `actor+0x54 = 0x10`, teardown — so a non-ground mover leaves no corpse, which is the other half of `ANIM-DEATH-007`'s “the four `movementType > 1` classes leave no corpse at all”

**Confidence.** High (the getter and its three vtable slots, the mask arms, the two occupancy immediates, `R1345`'s four arms and the corpse slam are named instructions, `evidence/rom-domain-excerpt.md` §1, §2, §6; the getter's identity comes from a `.rdata` slot read, the instrument that *can* see a virtual-only target) / Medium (that the consumer list is **complete**: two instruments, `EnumRefs disp:4a` (32 hits / 0 orphan — the direct reads) and an instruction-pattern census of indirect calls through offset +0x20 (229 hits / 128 owners / 0 orphan, of which the actor-family receivers are listed in §6). Neither sees a call through a computed pointer, and six of the listed owners — `R0460`, `R0306`, `R0225`, `R0252`, `R0181`, `R0267` — were located and **not read**)

**Original status.** ● active

**Evidence.** [EXP-0078](../experiments/EXP-0078-movement-domains/)

### MOVE-DOM-025

**The mask is installed once, at spawn, is the mover's only passability state, and survives a save because the mover is serialized whole.** `R0471` is `__thiscall(mover, actor)`: it calls `actor->vt+0x20()`, then decrements the result and tests for zero three times — `1 →` a byte store of 0x41 at `mover+0x5` (`L06846`), `2 → 0x44` (`L06847`), `3 → 0x82` (`L06848`) — and a value outside `{1,2,3}` falls through to `L06849` and **stores nothing**, leaving whatever was there. It has **2 call sites** (`EnumRefs callto:R0471`, 2 owners, 0 orphan): `R0184` @`L06850`, the Units-table spawn constructor's tail, and `R0876` @`L06851`, the Humans-table defaults routine — both inside object construction, neither reachable afterwards. The byte's **complete direct-write population** is four instructions: the mover constructor `R0206` @`L06852` (`0x41`, after a `REP STOSD` of `0x2d` dwords = the mover's `0xb4` bytes) and `R0471`'s three (`EnumRefs re:` on the store form — 33 hits / 25 owners / 0 orphan, every other hit on a different object, most of them `MOVE-STEP-010`'s `pos+0x04`/`+0x05` fraction pair written as a `[+0x4]`/`[+0x5]` couple). **So the mask is a cache with no refresh** — and the rival that follows from that, *a deserialized actor keeps the constructor's `0x41`*, is real all the way to its last step and then false: `Unit::CreateObject` (`R1346`) runs the **default** ctor `R0901`, which pushes the literal `"null"` (`L06853`, `StrDump`) into the `Data.bin` name search; no shipped Units row is named `null`; and the search's failure arm reports through `R1347` and the jump to `L06854` at `L06855` — **past** the mask install. It dies because `Unit::Serialize` `R0210` calls `R1348` at `L06856`, *before* its own `CArchive::IsStoring` test at `L06857`, and `R1348` is eleven instructions that push `0xb4` and the mover and call `CArchive::Read` (`R0271`) or `Write` (`R0272`) — the mask is byte 5 of that block and round-trips verbatim

**Confidence.** High (every instruction above is transcribed from full listings of five routines, `evidence/rom-domain-excerpt.md` §2 and §7; the `"null"` literal is `StrDump`'s and the three `CRuntimeClass` records are read out of the PE section table with no disassembler in the loop; the save rival is refuted by a **present** instruction, not by an absent one) / Medium (the write enumeration's blind spot, which is exactly what the rival tripped on: a `re:` sweep over one store form cannot see a **wholesale copy** of the mover, and `R1348` is such a copy)

**Original status.** ● active

**Evidence.** [EXP-0078](../experiments/EXP-0078-movement-domains/)

### MOVE-DOM-026

**What each domain is permitted to enter, terrain and non-terrain counted apart — and the axis is not terrain, it is which of the three block bits the mask carries.** Over the 38 shipped maps' static block plane, rebuilt as the ingest (`TERR-PASS-049`/`050`) and the structure pass (`TERR-STRUCT-068`/`070`/`071`) build it and classified by the arm that set each byte (`tools/domainstop -mode stops`, `evidence/stops.txt`): domain **1** (`0x41`) is stopped by **199 361** terrain cells (Mountain 117 572, water 77 845, tile-word bit 13 3 944) **and** 242 179 non-terrain (type-3 object 69 676, structure 13 527, border 158 976) = **441 540** of 880 704 (50.13 %); domain **2** (`0x44`) by **0** terrain and the same 242 179 non-terrain (27.50 %); domain **3** (`0x82`) by **0** terrain and **158 976** — the border alone — (18.05 %). So **domain 2 passes water and mountain and is stopped by every object, building and ground occupant, while domain 3 is stopped by nothing on shipped terrain except the 8-cell border and another domain-3 occupant**: they are not two grades of one thing, and they disagree on **83 203** cells. Runtime occupancy (bits 6/7) is by construction invisible to a static census and is stated separately: the same masks make bit 6 block domains 1 and 2 and bit 7 block domain 3 (`MOVE-DOM-024`(b), `TERR-PASS-051`). Consequence a route consumer must reproduce: the 8-connected free set — the search steps 8 neighbours and applies **no** corner rule (`MOVE-ROUTE-004`) — falls into **more than one piece on 23 of 38 maps for domain 1 and on 0 of 38 for domain 3**, whose largest piece is the whole interior on every map (`evidence/cross.txt`)

**Confidence.** High (the rule is `TERR-PASS-051`'s named instructions and the arm attribution is the ingest's own four assignments plus the structure pass's OR with 5 and AND with 0xfa; the discriminator is that **bit 1 is set on 158 976 cells and 0 of them by any arm but the border stamp** — an exhaustive re-execution, not a sample — so domain 3's stop set is measured rather than inferred, and the `==`-comparison rival scores **0 / 0 / 0** blocked cells) / Medium (the figures are one corpus on one root; and the attribution reports the arm that wrote a byte **last**, which the ingest's four assignments make well defined but which cannot separate two arms that agree on a cell)

**Original status.** ● active

**Evidence.** [EXP-0078](../experiments/EXP-0078-movement-domains/)

### MOVE-DOM-027

**Nothing narrows a mover's verdict after the mask: no height term, no corner rule, no order-time check, no step-time check.** (a) Both passability predicates were read **end to end** — `R1331` (static plane) and `R0144` (dynamic) are instruction-for-instruction identical apart from the plane displacement: `n = actor->vt+0x1c()`, `n <= 0` returns free, and the `n × n` walk's only memory operand is a byte bit test of the cell's plane byte against a mask **re-loaded from `mover[0x154]+5` on every cell**. No cost read, no height read, no edge or adjacency term, no reachability seed. (b) The **height plane has six references image-wide** (`EnumRefs disp:9451c`, 5 owners, 0 orphan): the line-of-sight pair `R0136`/`R0134` (`AI-SIGHT-006`), the constructor's memset, the ingest's write, and the two reads in `R1088` (`TERR-MOVE-056`) — **none in the search**, so height gates sight and speed and never passability. (c) The two block planes have **no reader below `R1349`** over `EnumRefs disp:10000` (78 hits / 33 owners) + `disp:20000` (60 / 19), 0 orphan — every other owner lies in `R0144…R0227` — so no interface, input, campaign or view routine can consult passability, and there is **no click-time or order-time refusal of a destination anywhere in the image**; a group move writes one loop-invariant cell into every member's order block with no test at all (`MOVE-ORDER-023`). (d) The step routine `R0054` tests no plane (`MOVE-WAIT-008`), and the mask's only other consumer at step time is the claim, which *writes* (`MOVE-CLAIM-007`)

**Confidence.** High for (a) and (d) — full listings, and “no other operand” is an exhaustive read rather than an enumeration / Medium for (b) and (c) as **enumerations**: a `disp:` sweep cannot see an access through a pointer the routine `LEA`'d first (the height plane's one `LEA` is the memset, whose destination never leaves the constructor; both block planes' `LEA` hits are followed to their stores in `TERR-PASS-073`) nor a wholesale copy of the containing structure

**Original status.** ● active

**Evidence.** [EXP-0078](../experiments/EXP-0078-movement-domains/)

### MOVE-DOM-028

**Which shipped classes can be non-ground — and that a player can never command one.** `Data.bin` **Humans**: 210 parameterised rows, `movementType` (slot `0x16`) is **−1 on all 210**, so the base constructor's `1` stands and every human, hero included, is a ground mover. `Data.bin` **Units**: 57 named rows, `movementType` (slot `0x20`) is 1 on 40 and non-1 on **16** — the four difficulty variants each of **Ghost** and **Bee** (2) and **Bat_Sonic** and **Dragon** (3); no row anywhere carries a value outside `{−1,1,2,3}`. The two domain-3 classes are exactly the two with `units.reg` `Z != 0`, i.e. exactly the two allocated as `CAirUnit` and drawn in layer 3 (`REG-UNITS-061`) — an independent second witness for the name *air*, the third being `AI-ACQUIRE-002`'s distance penalty (`MOVE-DOM-024`(e)). The player's own roster contains none of them: the hero is a Human, and the tavern's fifteen hireable types are `Catapult` and `Ballista` (Units, `movementType` 1) plus thirteen `NPC%02d_%d` **Humans** (`MERC-TYPE-001`). Over the 38 maps, all **1739** non-ground type-6 placements (1024 domain-2, 715 domain-3, resolved by `ALM-CLS-038`'s key rule; 4 of 8094 records unresolved) belong to scenario-authored groups — `Monsters`, `Nocturnal`, `Beasts`, `Enemy`, `Ghosts`, `Wild`, `Evil-Creatures`, … — whose `Player+0x28` is 1 or 2, never a human participant's 0 (`UNIT-OWNER-009`). **The corroboration that would fail if the domain assignment were wrong:** scored against its own mask, **0 of 1739** non-ground placements stands on a cell that blocks it, while **258** of them (213 domain-2, 45 domain-3) stand on cells the **ground** mask blocks — against 4 of 6351 for ground placements on either scoring (`evidence/who.txt`)

**Confidence.** Medium (the value tables and the placement scoring are a corpus census on one root, which caps here by rule even though the null model fails hard — a wrong domain assignment cannot make 258 in-terrain placements come out 0/1739 — and the hireable roster is `MERC-TYPE-001`'s reading rather than this experiment's) / **Unknown** (whether a *trigger* can hand a non-ground actor to a human participant mid-mission: the spawners `R0062` and `R0066` and the owner-change `R0064` exist and were not walked over the shipped trigger set, and `Control Spirit` spawns a `Ghost` owned by the caster — `MAGIC-SING-019` — which is a domain-**2** mover a player *can* end up owning)

**Original status.** ● active

**Evidence.** [EXP-0078](../experiments/EXP-0078-movement-domains/)

### MOVE-RATE-029

**The whole rate law, composed, with the instruction for every term — and the composition is five terms, not one.** `R1088(world; actor, u16 srcCell, u8 facing)` is called from **one site image-wide** (`EnumRefs callto:R1088` → 1 hit / 1 owner / 0 orphan, `R0054` @`L06858`), at the **start of a cell transit**, so the rate is frozen for the whole transit and re-derived at the next cell. `dir = ((facing + 0x10) >> 5) & 7` (`L06859`…`L06860`); `dst = src + (i16)world[0x58ec0 + 4·dir]` (`L06861`). Then `domain = actor->vt+0x20()` (`L06862`) forks. **Ground (`domain == 1`)**: `speed` = `[[actor+0x70]+0x3c]+0x44` **zero-extended** when that byte is nonzero (`L06863`…`L06864`, the group term, `MOVE-GROUP-030`), else `(i16)actor+0x8c` (`L06865`, the class `Speed` column); `d = clamp((i8)(height[src] − height[dst]), −32, +32)` over `world+0x9451c` (`L06866`…`L06867`); `v = SpeedMultiplier · speed` with `SpeedMultiplier = world+0x58db4` (`L06868`/`L06869`, `MOVE-PARAM-006`, ships 8 in **both** roots); `v += (v·d) >> 6` **arithmetic** shift, uphill (`d < 0`) reducing (`L06870`…`L06871`); `c = ((u8)(cost[src] + cost[dst])) >> 1`, a **byte** add that wraps at 256, and `if c == 0 then c = 8` (`L06872`…`L05504`); `v = v / c` **signed** (`L06873`); `v = clamp(v, 1, 63)` (`L06874`…`L06875`). **Non-ground (`domain != 1`)**: `v = (speed·8)/8` — the identity — from the **same two sources in the same order** (`L06876`…`L06877`), with **no** multiplier, **no** slope, **no** cost lookup, then the same clamp. Consequence a consumer must carry: the two arms **agree exactly** when `SpeedMultiplier == 8` and mean cost is 8, which is what ships, so no corpus can separate them; move either and only the ground movers change. Then `mover+0xa8 = v`, `mover+0xae = dir` (`L06878`/`L06879`), the per-axis steps by `MOVE-DIR-034`, and `mover+0xaa = ceil(256 / step)` (`L06880`…`L06881`). `cost(cell)` is `R1087`, reading the plane at **`world+0x00000`** and **not a pure read** — with `block[cell] & 0x20` and a nonzero byte at the cell's `world+0x540b4` record `+0xe` it shifts the stored byte right by 2 and writes it back (`L05498`/`L05499`)

**Confidence.** High (every term, its address and its arithmetic form — `SAR` against `SHR`, the byte-width add, the signed `IDIV`, the zero- against sign-extension of the two speed sources — is transcribed from the raw listing in `evidence/listings.md` §1–2, and the single call site is an `EnumRefs callto:` with 0 orphan)

**Original status.** ● active

**Evidence.** [EXP-0093](../experiments/EXP-0093-move-rate/)

### MOVE-GROUP-030

**The speed override `TERR-MOVE-058` left Unknown is a GROUP term, and it is the minimum `Speed` over the group's members.** `actor+0x70` is the AI **group** (`AI-GROUP-009`) and `grp+0x3c` its `0x50`-byte AI record, so `R1088`'s `[[actor+0x70]+0x3c]+0x44` is `grpAI+0x44` — the byte `AI-MOVE-023` already reads as a flag at `L06882`. Both group-order setters compute it the same way: a local initialised to **`0xfa`** (`L06883`, `L06884`), then a per-member loop that compares the member's 16-bit speed at `+0x8c` with the running minimum and, unless greater-or-equal, copies the low byte of that speed into the local (`L06885`…`L06886`, and `L06887`…`L06888`), then `grpAI+0x44 = min` (`L06889`, `L06890`) beside `grpAI+0x20 = 4`, the Move order. It is cleared to **0** together with the order code when the group's per-member pass ends (`R0200` @`L06891`/`L06892`). So **a unit under a group order moves at its slowest companion's rate and reverts to its own class `Speed` when the order clears**, and because the store masks to 8 bits while the class fallback is a sign-extending load, the group term is **unsigned** where the class term is signed. Instrument for the writer set: `EnumRefs disp:44`, 74 byte-wide hits image-wide, of which exactly four in the AI module code-region touch `[reg+0x44]` as a byte — two writes, one clear, one compare; blind, as every `disp:` sweep is, to a wholesale structure copy

**Confidence.** High (the initialisation, the comparison, the store, the clear and the two call sites are named instructions, and the `min` shape is the loop's own arithmetic) / **Medium** (that a *running* game reaches this store on any particular order: the gate at `[ESP+0x18]` is derived from `[[grp+0x44]+0x30]+0x1f == 2` and a spread test against `world+0xa824`, neither identified — see the write-up's open questions) / ~~**Unknown** (whether any path writes `grpAI+0x44` while `grpAI+0x20` is not a move order)~~ — **both close, and one clause of this row falls, by [EXP-0094]**. The gate is a single formation flag: the two-indirection comparison is the owning `Player`'s formation mode `[[grp+0x44]+0x30]+0x1f`, default **2** (`AI-FORM-037`), and the spread test is a **Chebyshev distance in whole cells to the group centroid** against `AImanager+0xa824` — a compile-time **2** on the **AI manager**, not `world+0xa824`, which is the wrong object (`AI-SPREAD-038`, retracted.md). The store fires **iff** the group moves in formation (`MOVE-GATE-035`), and no path writes the byte outside a group Move/Swarm-2 order except the record's own constructor and its raw `CArchive` (de)serialization (`MOVE-GROUP-037`). **And the clear is retracted**: `R0200` has no caller by any of four instruments (`AI-DEAD-036`), so a unit does **not** revert to its own `Speed` when the order ends

**Original status.** ● active (amended)

**Evidence.** [EXP-0093](../experiments/EXP-0093-move-rate/), **[EXP-0094](../experiments/EXP-0094-group-rate/)**

**Amended.** The group-rate clear and world-field ownership clauses are partially retracted. MOVE-GROUP-037 and the correction in retracted.md state the retained field and AI-manager owner.

### MOVE-TURN-031

**Stepping requires matching facing; route destruction and final default0x10 were false interpretations.** `R0054` derives next-cell facing in 32-unit steps and tests current against desired before its rate/step branch. `L06893 CALL R0249` advances current facing without touching a route; the short-arc snap requires the full mover+a0 dword to be zero. MOVE-TURN-044 gives the local turn update. The table destination binding stands: Units slot9 atL06894 and Humans slot7 atL06895 address mover+a. The old clause that67 unparameterized rows keep0x10 as the resulting actor rate is partially retracted. The mover constructor writes 16, then Unit defaults write8 to the same allocated object; Human derive later copies low8(actor+8c). MOVE-RATE-052 and MOVE-RATE-053 distinguish these producers. Earlier raw table statistics were not remeasured by these corrections and do not establish final live rates.

**Confidence.** **High** for the named local instructions and table destinations. Whole-action elapsed time, final corpus-wide actor rates and first post-LOAD scheduling are not established.

**Original status.** ● partially retracted

**Evidence.** [EXP-0093](../experiments/EXP-0093-move-rate/), [EXP-0324](../experiments/EXP-0324-turn-continuation/), [EXP-0329](../experiments/EXP-0329-mover-rate-binding/); [retraction](retracted.md)

**Amended.** Route destruction, unconditional snap, elapsed-time and final-default-rate clauses are partially retracted. MOVE-TURN-044 and MOVE-RATE-052/MOVE-RATE-053 state the narrower relations; retracted.md records their scope.

### MOVE-CLOCK-032

**The displacement is one step per actor per SUB-TICK, it is counted in ticks and never measured in milliseconds, and the presentation tick that drives animation is issued from the same loop iteration.** `R0047` has **2 call sites image-wide** (`EnumRefs callto:R0047`, 0 orphan): `R0039` @`L02036` and `R0054` @`L06896`. Every path through the actor update reaches exactly one of them — `R0178` calls `R0039` and **returns immediately** when either fraction is not `0x80` (`L06897`…`L06898`), and otherwise reaches `R0054` once at its tail (`L06899`); `R0016`'s arms are a jump table. Above that the chain is single-caller at every step: `R0427` ← `R0193` (1 hit) ← `R0147` (`L01863`, unconditional), and `R0193` increments `server+0x04`, the sub-tick counter. **So one actor advances by `mover+0xb0`/`+0xb1` exactly once per sub-tick**, and the arrival test is `mover+0xac >= mover+0xaa` — a tick count against a tick count (`L06900`…`L06901`), with **no elapsed-time term anywhere in the routine that writes the position**. Three consequences, all of which a consumer gets wrong by modelling seconds: cells **per tick** is invariant; cells **per second** is `1000 / (T · campaign+0x3f0)` and therefore moves with the `[GameOptions] Speed` index; frame rate cannot enter, because `R0454` is a **deadline loop that repeats** rather than a scaler (`L02630 JBE`); and `campaign+0x3dc & 1` clear stops **both** the simulation and the `0x401` presentation tick (`L03387`/`L06902`), so a paused unit does not drift and its animation does not either. On the normal play arm the simulation sub-tick (`L02626`) and the `0x401` tick (`L02627`) are two calls of **one** iteration in a fixed order, simulation first — they are not two clocks but two consumers of one pacer (`SESS-PACE-018` bounds where that stops being true)

**Confidence.** High (the two call sites, the single-caller chain, the tick-against-tick arrival test and the two calls' adjacency in one loop body are named instructions and `callto:` enumerations with 0 orphan) / **Medium** (that no *other* arm of the engine advances an actor twice in one pacer iteration: `R0193`'s second caller `R0075` does run it sixteen times in a loop, and is a profiler reached from `R0261`, not from the idle handler — read, but its own caller set was not enumerated)

**Original status.** ● active

**Evidence.** [EXP-0093](../experiments/EXP-0093-move-rate/)

### MOVE-LIMIT-033

**The customisation limits of the rate (goal G2), by complete enumeration of the law's input domain rather than by sample.** (a) **The clamp `[1,63]`** (`L06903`, `L06875`) is a hard-coded pair of immediates; at `SpeedMultiplier` 8 and mean cost 8 the ceiling first bites at `Speed = 63`, against a shipped maximum of 35, so it is a guard on shipped data and a wall on authored data. (b) **The transit quantisation is the severe one.** `mover+0xaa` is a whole number of sub-ticks and the fractions are snapped to the cell centre when it is reached, so the surplus travel is discarded and every `v` in a class crosses a cell in the same time: `v ∈ [1,63]` yields only **27** distinct transit times and the largest class is **12** values wide (`v` 52..63 all take 5 ticks). Against the shipped `Speed` column that is **22 distinct values → 15 distinct transit times at cost 8, 13 at cost 6, 11 at cost 16**. (c) **The slope term is a `>>6`**, so a height delta under 4 changes nothing at `Speed` 16. (d) **A diagonal is not `√2`**: the `0.707` truncation and the round-up together leave the diagonal transit between **0.97×** and **1.15×** of `√2 ×` the straight one over the shipped speeds. Which of these move a shipped file's bytes if lifted: `Speed` and `RotationSpeed` **yes** (`world.res:data/data.bin`), `SpeedMultiplier` and the `Cost*` alphabet **yes** (`world.res:data/map.reg`); the clamp, the `>>6`, the `256` grid, the `0.707` and the nine-arm tick ladder **no** — none of them is carried by any shipped file. Shipped alphabets, identical in the EN and RU roots (0 differing rows over 333): `Speed` 22 distinct values in 8..35 (mode 19), `RotationSpeed` 16 distinct in 8..23 (mode 19), 67 rows unparameterised in both, `movementType` `−1:277 1:40 2:8 3:8`

**Confidence.** High (the collapse classes are a **complete enumeration** of the law's domain by a re-implementation whose every step cites its instruction, `tools/movespeed -mode limits`; the byte-moves-or-not column is a statement about which file carries each value, and each is named) / **Medium** (the shipped alphabets: two roots, one version of each, under EXP-0049's `data.bin` grammar and EXP-0032's `.reg` framing)

**Original status.** ● active

**Evidence.** [EXP-0093](../experiments/EXP-0093-move-rate/)

### MOVE-DIR-034

**The eight directions, the diagonal constant, and the per-axis step — the last three things a consumer needs to compute the next position.** `world+0x58eb0` and `world+0x58eb8` are two 8-byte tables written once by the world init `R0116` (`L00472`…`L00473`; `EnumRefs disp:58eb0` 4 hits / 1 write, `disp:58eb8` 5 hits / 1 write, 0 orphan): `dx = [0,+1,+1,+1,0,−1,−1,−1]`, `dy = [−1,−1,0,+1,+1,+1,0,−1]` — clockwise from north, all in `{−1,0,+1}`. The packed-cell delta table at `world+0x58ec0` is built from those two in an 8-iteration loop as `(dy << 8) + dx` (`L06904`…`L06905`), matching `MOVE-STEP-010`'s `(y<<8)\|x` packing. The step stores: **straight** (`dx·dy == 0`) is an 8-bit `IMUL` of `v` by the component (`L06906`/`L06907`), so the step is `±v`; **diagonal** is `trunc(v · component · K)` through `FILD` / `FMUL double ptr [L06908]` / the CRT `ftol` at `R0279` (`L06909`…`L06910`). **`K` is `0.707` exactly, not `√2/2`** — the eight bytes at `L06908`, read out of `rom.exe` through the PE section table without a disassembler, are `39 b4 c8 76 be 9f e6 3f` = `0.70699999999999996` (`evidence/diag-constant.txt`)

**Confidence.** High (the sixteen immediate stores, the delta loop, both step forms and the `FMUL` operand are named instructions; the constant's value is the shipped file's own bytes)

**Original status.** ● active

**Evidence.** [EXP-0093](../experiments/EXP-0093-move-rate/)

### MOVE-GATE-035

**The group rate term applies exactly when the group moves in FORMATION — one local flag decides both, and it is governed by two gates neither of which is data.** `R0108`/`R0157` seed `[stack+0x18] = 1` (`L06911`, `L06912`) and then run **gate 2**, the per-player formation mode `[[grp+0x44]+0x30]+0x1f` (`AI-FORM-037`): `== 2` takes the conditional arm, `!= 2` sets the flag **to the mode byte itself** and skips the spread test entirely (`L06913`…`L06914`, `L06915`…`L06916`), so `0` disables formation and any other nonzero value forces it. On the conditional arm the routine computes the group centroid with `R0140` — the members' 1/256-cell positions summed, divided **unsigned** by `grp+0x0c`, high byte of each taken — and per member runs **gate 1**: `R0167(memberCell, centreCell)` is `max(\|dx\|,\|dy\|)`, Chebyshev in whole cells, and the unsigned compare of the distance with the byte at `[table + 0xa824]` (the below-or-equal branch keeps the flag) clears the flag for the whole group the moment one member exceeds the threshold (`L01033`/`L06917`, `L01034`/`L06918`; the threshold is `AI-SPREAD-038`). The **same** flag is then tested twice: at `L06919` it forks the per-member loop (`MOVE-FORM-036`) and at `L06920` it gates the rate store `L06889` — and the running minimum is updated **only inside the formation arm**, because the plain arm jumps over it (`L06921` → `L06922` → `L06923`). Consequence for a consumer: `MOVE-RATE-029`'s group speed source is live iff the last group Move/Swarm-2 order this group received was issued in formation

**Confidence.** High (both gates, the fork and the store gate are cited instructions in two routines read end to end; the mode byte's identity rests on an `EnumRefs "re:.*\+ 0x1f\].*"` sweep of **8 hits / 8 owners / 0 orphan** image-wide, four of them stack, and the threshold on a `disp:a824` sweep of **3 hits / 3 owners / 0 orphan**)

**Original status.** ● active

**Evidence.** [EXP-0094](../experiments/EXP-0094-group-rate/)

### MOVE-FORM-036

**A group move DOES distribute destinations: there is a formation, there is a per-member offset table, and there is a spread test — `MOVE-ORDER-023`'s headline is refuted, through the blind spot that row's own confidence cell named.** On the formation arm each member's order block records `ord+0x24 = memberCellX − centroidCellX` (`L01031`, `L01032`) and `ord+0x26 = memberCellY − centroidCellY` (`L06924`, `L06925`) as `i16`, then the routine reads them back **as bytes** and issues `R0301(member, X + dx, Y + dy)` (`L06926`…`L06927`, `L06928`…`L06929`); on the plain arm every member gets the unmodified `(X, Y)` (`L06921`, `L06930`). `R0301` clamps the destination into the playable rectangle `world+0x58ee0..+0x58ee3`, writes `ord+0x0a = (y<<8)\|x`, `actor+0x50 = 1` and `ord+0x08 = 1` — **and that last store is the byte store of a register value at `L01930`, the register loaded `1` at `L00604`**, i.e. the move opcode arrives in a register, which is exactly the encoding `MOVE-ORDER-023`'s immediate-store enumeration states it cannot see. The offsets are scratch: `EnumRefs disp:26` is 11 hits image-wide and exactly **4** lie in `L06931..L06932`, the two stores and the two byte reads inside these two routines, so nothing re-forms a group later

**Confidence.** High (both arms and the clamp are full listings of routines read end to end; the refutation is an instruction the earlier row's instrument is documented as excluding, not a disagreement about what a listing says)

**Original status.** ● active

**Evidence.** [EXP-0094](../experiments/EXP-0094-group-rate/)

### MOVE-GROUP-037

**The complete writer set of `grpAI+0x44`, and the finding that NOTHING clears it — so a group keeps a rate it can no longer justify.** Instruction-level writers, from `EnumRefs disp:44` over the whole image (**630 hits / 270 owners / 1 orphan**, of which 48 are byte-wide on a register base and exactly **6** have a `[grp+0x3c]` base): `L06889` and `L06890`, the two gated stores (`MOVE-GATE-035`); `L06891`, the only clear — **and its routine `R0200` is unreachable** (`AI-DEAD-036`). The remaining two hits are reads: `L06882` in the group order-4 arm and `L06933`/`L06863` in the rate routine. Two further writers exist that **no displacement sweep can see**, and they are named rather than left to the blind spot: the record's constructor `R0150` zeroes the whole `0x50` bytes (`L06934`…`L06935`, `REP STOSD` × `0x14`), and `R1350`, the record's `Serialize`, is a **raw `CArchive::Write`/`::Read` of `this` for `0x50` bytes** (`L06936`/`L06937`) — so the group rate is in every save carrying a group, and a load writes it wholesale. Lifecycle, each hop a cited routine: a **player** order always allocates a new group (`R0196` → `new(0x48)` → `R0149` → `R0150`), so it starts at 0; a **scenario** group is the authored one and persists, and guard / aggressive / stand-ground / patrol / roam write only `grpAI+0x20`; a member dying or being re-ordered goes through `R1140`, which unlinks it and zeroes `actor+0x70` and touches nothing else, **so the group keeps the departed member's `Speed`**. The order ending does not clear it either — `R0154`'s only use of the byte is as a boolean guarding `ord+0x6c = 1`, and over the image `disp:6c` (496 hits / 248 owners) puts **three** hits in the simulation modules, all of them stores

**Confidence.** High (the writer and reader sets are printed enumerations with their instrument and both of that instrument's blind spots discharged by name; the unreachability is `AI-DEAD-036`) / **Medium** (that a *running* game ever leaves a stale value in play: the persistence argument is a reading of which routines write the byte, not an observation of a session)

**Original status.** ● active

**Evidence.** [EXP-0094](../experiments/EXP-0094-group-rate/)

### MOVE-AREA-038

**The complete movement-side reader set for both block planes, and the one channel through which an area effect reaches it.** Instrument: `EnumRefs disp:10000` (78 hits / 33 owners) and `disp:20000` (60 / 19), 0 hits in orphan or undisassembled code, re-run here and agreeing with `MOVE-DOM-027`'s counts. Every plane read in the movement path is a byte bit test of the cell's plane byte against the mask `mover+0x5` reloaded per cell, in six routines: `R1326` and `R1331` on the static plane, `R1327`, `R1328`, `R1329`, `R1330` and `R0144` on the dynamic plane, plus the copies inlined in the driver `R0053` (`MOVE-SEARCH-001`). **No movement routine reads an area-effect layer slot, a cell-record occupant slot, or a plane bit by immediate.** The mask population is `R0471`'s three stores `0x41`, `0x44`, `0x82` (`MOVE-DOM-025`), so the only plane bits that can ever block are 0, 1, 2, 6 and 7. An area effect therefore reaches the search through exactly two channels, both of them writes performed by `R0453` when it recomputes a cell from its record (`MAGIC-WALLBLOCK-045`, `MAGIC-AREACOST-046`): the `wall_of_earth` layer slot `payload+0x20` sets bits 0 and 2 on **both** planes, and any occupied layer slot multiplies the cell's cost byte by 4, which `MOVE-COST-002`'s `movementType == 1` arm reads inline at the destination cell. The area module's own two plane writes reach neither: bit 5 (`L05450`) and bit 4 (`L05456`, `L05457`) are in no mask, and bit 4's only consuming reader is a serializer (`TERR-PASS-148`). **`R0144` is a standalone footprint query on the dynamic plane** — instruction-for-instruction `R1331` apart from the displacement, per `MOVE-DOM-027` — with 5 call sites in 2 owners, `R0209` and `R0142`; it is not part of the search

**Confidence.** High for the reader set as an exhaustive read of six listings, and for the two channels, each a named instruction in `R0453` / Medium for the enumeration's completeness, whose blind spot is the one `TERR-PASS-073` names: a `disp:` sweep cannot see an access through a pointer the routine `LEA`'d first, and the two address-of-cell getters `R1349` and `R1351` had their callers enumerated (1 and 2, neither a mover) rather than being covered by the sweep

**Original status.** ● active

**Evidence.** [EXP-0174](../experiments/EXP-0174-area-movement/)

### MOVE-GATE-039

**(rom.exe) No actor is displaced except from inside one call of the order machine, so an actor the machine refuses cannot move at all.** `MOVE-CLOCK-032` names `R0047` as the routine that writes the position. Closing the call graph above it with `EnumRefs callto:`, whole image, 0 orphan at every step: `callto:R0047` → 2 hits, `R0039` @`L02036` and `R0054` @`L06896`; `callto:R0039` → 3, `R0016` @`L06938`, `R0178` @`L02035`, `R0043` @`L06114`; `callto:R0054` → 2, `R0178` @`L06899`, `R0043` @`L01907`; `callto:R0178` → 2, `R0016` @`L01932`, `R0087` @`L06117`; `callto:R0043` → 2, `R0042` @`L06113`, `R0016` @`L06939`; `callto:R0087` → 1 and `callto:R0042` → 2, all of them inside `R0016`. `callto:R0016` → 2, `R0037` @`L00086` and `R0038` @`L00087`. `R0038` is a loop that calls the machine for every element of a list and its own `callto:` is **0 hits**, which `INSTRUMENT.md` rule 7 says is a tell rather than an answer — so every section of the image was scanned for the stored little-endian dword (`tools/effectgate -mode refs`): `R0038` appears **0** times, so nothing reaches it; `R0016` appears 0 times, consistent with one direct `CALL`; and `R0037` appears **3** times, at `L03030`, `L03031` and `L03032`, offset `0x18` in the three actor vtables, which reproduces `MOVE-TICK-009`'s slot 6 from a different instrument. Inside the machine, only one arm above the order switch reaches a displacement — progress arm 3 at `L00096`, which calls `R0039`. Progress arm 4, the refusal `MAGIC-ACTGATE-079` describes, calls nothing. **So a refused actor is not slowed and does not drift: no instruction that writes its position can execute.** The same closure bounds the refusal's one escape (`MAGIC-ACTGATE-079`): `mover+0x98`, the flag that admits the machine's un-gated tail, has exactly three setters — `L01720` in `R0178`, `L00161` in `R0043` and `L00162` in `R0055`, whose two callers are those same two routines (`EnumRefs disp:98`, 215 hits / 140 owners / 4 orphan; `disp:99`, 0 hits) — and all three sit inside executors the closure puts behind an order arm. So while an actor is refused the flag can never be raised again, and the tail's own `L05750` clears it

**Confidence.** **High** (six `callto:` enumerations with 0 orphan hits at every step, and the two results that a `callto:` cannot settle — a zero and an all-`RDATA-SLOT` set — settled by a raw scan of every section for the stored dword, which is the instrument that can see a pointer table) / **Medium** that no *other* routine writes an actor's position without going through `R0047`: that identification is `MOVE-CLOCK-032`'s and was cited rather than re-derived

**Original status.** ● active

**Evidence.** [EXP-0179](../experiments/EXP-0179-effect-action/)

### MOVE-STEP-040

**(rom.exe) The refusal begins on a cell centre, so a step already in flight completes.** `R0016` writes the parking value only when the order object's progress byte is already 0: `L05744` loads the byte at `+0x9`, `L06940` tests it and `L05745` jumps to `L06941` when it is non-zero (rel8 `0x03`). `AI-ORDER-039` establishes that progress 0 is the tick the actor stands on a cell centre, `R0040` testing `pos+0x4 == pos+0x5 == 0x80`. An actor in transit is in progress state 3, whose arm at `L00096` calls `R0039` — the one arm above the order switch that still advances a step — sets `actor+0x54 = 1`, and clears the progress byte when `R0040` reports arrival (`L06942`…`L04569`). The state is entered by the walk arm itself at `L06943` and by `L06944`, `L06945` and `L00107`, each immediately after the same `R0040` test fails. So the sequence is: the effect attaches, the actor finishes the cell it is entering, the progress byte returns to 0, and the gate fires on the following tick. **The observable is a stop aligned to the grid, never a stop between cells**

**Confidence.** **High** (the order of the two tests inside one routine, the arm's own call and its clearing condition are named instructions, byte-asserted at their own addresses; the cell-centre meaning of progress 0 is `AI-ORDER-039`'s, cited)

**Original status.** ● active

**Evidence.** [EXP-0179](../experiments/EXP-0179-effect-action/)

### MOVE-EFFLIST-041

**(rom.exe) No routine in the movement path reads the drawable's effect list, so the list is not a second immobilisation mechanism.** `MOVE-GATE-039` and `MAGIC-ACTGATE-079` place the only refusal in the order machine, on `actor+0x144`. The effect list at `drawable+0x124`/`+0x128`/`+0x12c` was checked separately. It has one reader, `R0615`, whose complete call-site set is fourteen addresses over six routines, every one of them between `L05789` and `L05801` in presentation code (`MAGIC-EFFLOOK-083`). A displacement sweep of the list's own three fields over the whole image, `EnumRefs disp:124 disp:128 disp:12c imm:124`, returns 71, 78, 130 and 10 hits. Every hit in the code-region and code-region family -- the order machine, the movers, the AI -- is a **byte** access at `+0x12c` on the simulation actor, which is the reach field `R0004` compares a distance against at `L06946`, while `R0615` reads `drawable+0x12c` as a **dword**. The `imm:124` hits are three stack frames, two address computations at `+0x1244` and `+0x124c`, and five additions of the constant 0x124 to a base, of which two are `R0509`'s own message arms. No simulation-family access to the list was found.

**Confidence.** **High**, scoped to what was searched: two instruments with different blind spots for the reader's call sites (a reference index on the repaired function table, and a raw `E8 rel32` scan of every executable section) agreeing address for address, and a displacement sweep over the whole image reduced by access width and by base class rather than by address family alone. Blind spot: a read of the list through a pointer held in a register with no displacement, or through a computed call, is invisible to both

**Original status.** ● active

**Evidence.** [EXP-0182](../experiments/EXP-0182-effect-at-actor/)

### MOVE-072

**`mover+0x90` is not "route list non-empty" — it is a boolean "actor already stands on the route's own resolved endpoint" flag. Image-wide, `1` is written only at `L00457`, inside `R0178`'s own walk-arm tail; `0` is written at eight further sites across seven routines, `R0178` itself included — not "only in `R0178`," as this row first had it.** A whole-image write-site census (`evidence/mover90-write-census.json`, this round) finds every store to `[reg+0x90]` with no index register across the whole `.text` section by linear-sweep capstone decode — sequential decode with no control-flow following, whose own blind spot is a byte sequence that is really embedded data, or a misaligned instruction tail, decoding as a spurious hit — then keeps the sites whose base register was loaded from `[x+0x154]` (the mover-pointer fetch) within the 25 preceding decoded instructions. It finds 150 raw `disp+0x90` stores and nine through a mover pointer: `L00457` (value 1) and `L06947` (value 0), both inside `R0178`, read whole this experiment; `L00458` and `L01715` (both value 0), inside `R0154`, the Move arm's own per-member routine, also read whole this experiment; and five sites outside this experiment's nine-routine scope, each storing 0 — `L06948` (`R0301`), `L06949` (`R1352`), `L06950` (`R0010`), `L06951` (`R1353`), `L06952` (`R1354`). Attribution for those five follows `Image.body()` recursive descent seeded from those five entries, which independently confirms each site falls inside that routine's own reachable body; this experiment did not read any of the five whole, so nothing beyond "this routine writes 0 here" is claimed for them. `R0178`'s own two writers work as before: the 32-bit store of zero to `+0x90` at `L06947` clears it unconditionally right after the static search `R0053` returns (`L06802`), regardless of that search's own result; the 32-bit store of 1 to `+0x90` at `L00457` sets it, later in the same tail, only when the actor's own current cell (`[actor+0x10]+2`) equals `mover+0x76` (`L06953` compares the 16-bit cell with `+0x76`, `L06954` branches on unequal) — and only on that branch does the routine return immediately (`L06955`..`L06956`), skipping the rest of the tail, including the call into the stepper `R0054`. `mover+0x76` itself, set earlier in the same tail (`L06957`/`L01721`), takes the cell word from the search result node at `[actor+0x164]+8` when the static search's own count field `[actor+0x168]` is nonzero, or collapses to the actor's own current cell when it is zero — the same branch that sets `mover+0x98` (`AI-ROUTE-045`'s route-search-failure flag) to 1 (`L01720`). So `mover+0x90 = 1` fires in exactly two circumstances, both inside `R0178`: a genuine "already there," or the degenerate case where a total search failure collapsed `mover+0x76` onto the actor's own cell. The gloss `AI-MOVE-023` gives the same flag, at its own read of `R0154`'s `L06958` branch comparing `+0x90` with zero, is corrected the same way — see `claims/retracted.md`, whose row is extended this round to also cover that routine's own arrival-branch gloss, "route list `pth+0x90` cleared": the two 0-writes this census finds inside `R0154` (`L00458`, `L01715`) clear `mover+0x90`, not a route list, either

**Confidence.** High for the two `R0178` writers and their conditions (cited instructions inside one routine read whole this round) and for the census's own site-and-value findings (a deterministic scan, independently reproduced from `regen.sh`) / Medium for the routine attribution of the five sites outside this experiment's own nine-routine scope, which rests on recursive descent from entries the census script did not itself discover, not on a fresh whole-body read of those five routines

**Original status.** ● active

**Evidence.** [EXP-0386](../experiments/EXP-0386-engagement-under-occupancy/EXP-0386.md)

### MOVE-073

**`R0043`, the pursuit order's own walk/search routine, has no `MOVE-WAIT-008`-style wait-and-turn branch anywhere in its 247 instructions — but it is not a universal "every path reaches the stepper" either: three named exits return without calling it, and one of the routines it calls can raise `mover+0x98` against an occupied approach through a mechanism this row did not previously read.** The three exits: not centred (`L06959`/`L06960` false, `L06961` calls `R0039` and returns at `L01891`); already within stop distance (`L00583 ja L06116` taken, turns via `R0051`/`R0056` and returns at `L06962`); and the routine's own reused-route match, where the current cell already equals both `mover+0x76` and `mover+0x8c` (`L06963`/`L06964`, both `jne` not taken), which increments the progress counter `mover+0x09`, clears `mover+0x7c`, and returns at `L06965` without calling the stepper. Every other path — a refreshed dynamic route, or the static search running at all — does reach the stepper `R0054` unconditionally, at `L01907`. The static search this routine runs is at `L01894`, not `L06802` as this row first had it (`L06802` is `R0178`'s own call to the same search); `MOVE-PLANE-005`'s own unit-blind static plane means a destination cell occupied only by another actor does not fail that search — it still returns a route ending at the literal requested cell — and this routine's own direct setter (`L00161`) fires only on a **total** static-search failure (`[esi+0x168] == 0` after the call at `L01894`). But this routine also calls `R0055` at `L01901`, passing the target actor as `altTarget` — Picker B's own contact-ring search (`MOVE-ALT-018`/`MOVE-ALT-020`), read whole this round for `AI-335`'s own correction — and that routine's own `mover+0x98 = 1` store at `L00162` (`MOVE-GATE-039`'s third setter) is reachable from here too, under the same two-part gate `AI-335` names: the inner dynamic-plane search against the goal comes back empty, and the branch that picked that goal aimed straight at the final destination. **So "never on mere occupancy" is unsupported**, and what an every-approach-occupied pursuit actually produces is Unknown — narrower than either "always waits like an ordinary walk" or a flat "fails" — because no reading traces Picker B's own contact-ring search for the specific case where every ring cell is itself occupied. `R0042`, the out-of-position wrapper `AI-PURSUE-040` already reads, confirms the wait-and-turn absence at its own call site: it calls `R0043` once (`L06113`), tests only whether the actor ended up centered (`R0040`, `L06966`), and on failure writes `ord+0x09 = 3` at `L00107` — entry into progress state 3, the in-transit step-along-the-path arm `MOVE-STEP-040`/`AI-ORDER-039` already name, not a retry marker — before unconditionally setting `actor+0x54 = 1` on every path

**Confidence.** High for the structural exit enumeration (both routines read whole this round — 247 and 21 instructions, 0 unresolved — every exit from `R0043`'s tail traced by hand) and for the corrected static-call address and progress-state gloss / Medium for what an occupied-but-passable approach ordinarily produces through the unconditional stepper call, which composes `MOVE-PLANE-005`/`TERR-PASS-051`'s already-published occupancy-blindness with this round's caller rather than re-deriving the passability test itself / Unknown, explicitly, for the every-approach-occupied scenario specifically: it depends on Picker B's own contact-ring search, which this experiment did not trace for that case

**Original status.** ● active

**Evidence.** [EXP-0386](../experiments/EXP-0386-engagement-under-occupancy/EXP-0386.md)


## Centered turn continuation

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-TURN-044 | The turn call advances orientation; its byte estimate is recomputed, not consumed as a countdown. | High | ✔ promoted | [EXP-0324](../experiments/EXP-0324-turn-continuation/), complete instruction listings and labelled derived transition examples |
| MOVE-RATE-052 | Mover construction, actor defaults and table binding are separate producers. | High | ✔ promoted | [EXP-0329](../experiments/EXP-0329-mover-rate-binding/), instruction, binding and vector measurements |
| MOVE-RATE-053 | Human derive replaces the mover byte with the low byte of the current actor speed word. | High | ✔ promoted | [EXP-0329](../experiments/EXP-0329-mover-rate-binding/), complete derive listing and discriminating transfer vectors |
| MOVE-RATE-054 | Effect selector18 adjusts the target's mover byte before invoking that same target's derive. | High | ✔ promoted | [EXP-0329](../experiments/EXP-0329-mover-rate-binding/), original table values and effect-to-call vectors |
| MOVE-RATE-055 | The reached byte consumer changes facing and an estimate; its actor comes from an argument, not incoming ECX. | High | ✔ promoted | [EXP-0329](../experiments/EXP-0329-mover-rate-binding/), all-file-backed-section references and bounded instruction vectors |

### MOVE-TURN-044

**The turn call advances orientation; its byte estimate is recomputed, not consumed as a countdown.** `R0056..L06967` writes desired facing to mover+1 and computes the pre-step shorter arc. When the DWORD at+a0 is zero it clears BYTE+9d; only that inactive arm snaps an arc<=32, writing current=desired and BYTE+a4=1. Otherwise it calls the complete leaf `R0249..L06968`, which reads BYTE current+0, desired+1 and RotationSpeed+a. The leaf snaps if the shorter arc is strictly less than the rate; otherwise it adds or subtracts the rate modulo256 in the direction reaching desired, taking addition for the exact128 tie. It contains no call or route/list access. The caller writes BYTE+a4=ceil(pre-step shorter arc/rate), sets DWORD+a0=1, increments BYTE+9d modulo256, then clears DWORD+a0 if the updated facings match. Nonzero rate is required on the division arm; zero-rate gameplay validity is Unknown. `R0054..L06969` tests facing equality before the rate/position step; its mismatch arm callsR0056 and returns, even when that turn finishes. The ordinary move caller reachesR0054 atL06899; the selected order arm calls that move routine atL01932. This establishes local repeatable progression and a later-call step boundary, not an unconditional next-server-tick promise for every order.

**Confidence.** **High** for the complete local turn/update and next-cell branch; **Medium** for the bounded ordinary-caller interpretation. Full route replanning, callback effects, all order scheduling, malformed/zero-rate lifecycle and post-LOAD first dispatch remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0324](../experiments/EXP-0324-turn-continuation/), complete instruction listings and labelled derived transition examples

### MOVE-RATE-052

**Mover construction, actor defaults and table binding are separate producers.** Unit initializeR0185 requests 180 bytes atL00963, passes the nonnull allocation toR0206, then stores that constructor's returned pointer in actor+154 atL06970. The mover constructor clears180 bytes and sets BYTE+a=16. The same initializer callsR0160 atL06971; itsL06972 loads actor+154 andL06973 writes BYTE+a=8. Unit tableR0180 passes that allocated byte address toR0281 atL06894 after nine sequential slot advances; Human tableR0657 does so atL06895 after seven. HelperR0281 reads the selected dword, skips a -1 assignment, otherwise stores low8, and always advances the cursor. A missing table value therefore preserves the incoming byte, not an independently chosen16.

**Confidence.** **High** for these complete helpers and selected caller prefixes: both images and 12 table vectors distinguish sentinel, zero, truncation and receiver aliases. Allocation-vector return is synthetic; allocation failure, all construction paths and final spawned rates remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0329](../experiments/EXP-0329-mover-rate-binding/), instruction, binding and vector measurements

### MOVE-RATE-053

**Human derive replaces the mover byte with the low byte of the current actor speed word.** Derive entryR0280 stores its actor receiver atEBP-14. AtL06974..L06975 it loads ECX=[actor+154], AL=BYTE[actor+8c], then writes [ECX+a]=AL. This follows its Reaction-based speed stores, type19/21 addition, carried-load penalty andR0840 fold of the speed modifier at actor+d8 into WORD+8c. The fold's signed-negative branch clears the modifier, not the already stored negative speed word. A complete derive listing and bounded original-instruction slices distinguish low-byte copy from an independent retained RotationSpeed, saturation or pointer alias. Unit vt+50=R0836 has no such store; Humanoid/Human vt+50=R0280.

**Confidence.** **High** for the stated local sources, widths and selected vtable slots. The eight composed speed-slice vectors do not execute intervening derive code; arbitrary inputs, final spellbook callback effects, all event reachability and first post-LOAD recomputation remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0329](../experiments/EXP-0329-mover-rate-binding/), complete derive listing and discriminating transfer vectors

### MOVE-RATE-054

**Effect selector18 adjusts the target's mover byte before invoking that same target's derive.** Original dispatch tableL04018 slot18 selectsL06976. Target is argument1 atEBP+8; the arm loads [target+154], adds low8 of the computed effect value to BYTE+a, and stores modulo256. The value is magnitude times signed16(effect+40) when mode+3d meets the original masks1/2/4, otherwise magnitude times DWORD(effect+40), with wrapped32-bit product. AtL06977 ECX is the same target and the call is [target.vtable+50]. For the measured Unit slot this selectsR0836; Humanoid/Human selectR0280, whose later store is MOVE-RATE-053. The direct effect adjustment is therefore an intermediate store, not proof of a surviving Human rate bonus.

**Confidence.** **High** for selector, arithmetic, exact target and outgoing virtual call; five original-instruction vectors per image end immediately before that call. Item application scheduling, duration/removal, intervening callbacks and full post-effect values remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0329](../experiments/EXP-0329-mover-rate-binding/), original table values and effect-to-call vectors

### MOVE-RATE-055

**The reached byte consumer changes facing and an estimate; its actor comes from an argument, not incoming ECX.** LeafR0249 reads actor fromESP+4 and loads its+154 intoESI before reading BYTE+a atL06978. It updates only current facing BYTE+0. All three raw E8 callers of this leaf decode in selected windows:L06979 when actor+184 is zero,L06980 after a requested-facing comparison, andL06893 in the turn caller. The first path returns directly after the leaf. The turn caller also reads the same allocated byte atL06981/L06982 as quotient/remainder divisor for BYTE+a4. Selected translation-rate bodyR1088 reads actor+8c or the formation override and does not directly read mover+a. Separate rate values7/19 change facing with Position unchanged in the selected no-route-count path.

**Confidence.** **High** for exact receiver and local consumer actions. Thirty-four selected instruction/data windows do not close the whole mover graph: the turn caller has 11 raw direct callers, eight outside the selection; rebased/indirect accesses, complete scheduling and elapsed gameplay time remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0329](../experiments/EXP-0329-mover-rate-binding/), all-file-backed-section references and bounded instruction vectors


## Restored actor event ordering

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-EVENT-060 | The common actor tick admits the selected turn consumer after effects and explicit actor/order gates. | High | ✔ promoted | [EXP-0330](../experiments/EXP-0330-first-mover-event/), selected complete bodies, original class/order tables and relative-transfer checks |
| MOVE-EVENT-061 | A selected attached-effect route can invoke derive on the very actor whose order consumer follows; the derive target is class-specific. | High | ✔ promoted | [EXP-0330](../experiments/EXP-0330-first-mover-event/), Effect/actor slots, constructor assignments and complete selected tick/apply/derive listings |

### MOVE-EVENT-060

**The common actor tick admits the selected turn consumer after effects and explicit actor/order gates.** Original Unit, Humanoid and Human tables all holdR0037 at+18. ManagerR0427 passes its current actor as ECX atL01868 only when actor+54!=16. That actor's tick first dispatches its attached entries through+38 with the same actor as argument1; only afterward does signed WORD(actor+94)>0 reach the order path. PredicateR0416 returns whether actor+3c is zero, so only a nonzero+3c admitsL00086 ->R0016(actor). With the order-machine prefix admitted, order+9=0 and the status hold not installed, its selected table maps order+8=10 toL06983. Unequal mover+0/+1 passes this same actor toR0056 atL00752; equality clears the pending byte instead. The turn's inactive short-arc arm snaps without reading+a; its other arm callsR0249 and reads the same actor's allocated mover+a for the estimate.

**Confidence.** **High** for the named original slots, receiver aliases, branch table and local control order. The prefix's imported critical-section call, position helpers, earlier effect/dispatcher calls and mutable receiver state remain explicit boundaries. This is conditional reachability, not an assertion that order10 or either turn arm is first after a real LOAD.

**Original status.** ✔ promoted

**Evidence.** [EXP-0330](../experiments/EXP-0330-first-mover-event/), selected complete bodies, original class/order tables and relative-transfer checks

### MOVE-EVENT-061

**A selected attached-effect route can invoke derive on the very actor whose order consumer follows; the derive target is class-specific.** The two measured Effect tablesL02981/L02980 holdR0673 at+38,R0565 at+40 andR0843 at+48. If an attached receiver has one of these tables, mode+3d includes mask0x02 and its incoming unsigned WORD+42 is a multiple of 8, the tick calls+40 before decrementing duration. For identities other than 8,12,17,R0565 forwards the unchanged actor argument and multiplier1 through+48; genericR0843 calls that actor's+50 atL06977. Selector0's measured table arm jumps directly to this tail, demonstrating that a direct magnitude store is not required to reach derive. Original archive constructors install Unit/Humanoid/Human tablesL00001/L00002/L00003; their+50 entries areR0836/R0280/R0280. Thus, if both this generic call and the later order consumer are reached in one actor tick, the same actor's derive call precedes that consumer. The Human/Humanoid derive's reachedL03932 stores low8(actor+8c) into mover+a; the selected Unit derive body has no such store.

**Confidence.** **High** for the conditional selected-table chain, exact actor and class distinction. An empty effect list, an ineligible pulse, a special identity or another virtual receiver need not take that chain. Actual restored effect population, intervening derive/notification callbacks, normal completion and the first global mover access remain **Unknown**; no native event was observed.

**Original status.** ✔ promoted

**Evidence.** [EXP-0330](../experiments/EXP-0330-first-mover-event/), Effect/actor slots, constructor assignments and complete selected tick/apply/derive listings


`EXP-0386` was allocated ids `72`..`73` of `claims/move.md` (2 ids) and spent both,
`MOVE-072`..`MOVE-073`. **No id is returned unused.** The next free `move.md` id is
therefore `MOVE-074`.

## Seed-cell route boundary

Member offsets, addresses and packed cells below are hexadecimal. Counts,
pool sizes and flag values are decimal.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-080 | Both selected original extractors leave an initially empty list empty when endpoint equals seed; one node is allocated and removed, with count 0 to 1 to 0. | High | ● active | [EXP-0422](../experiments/EXP-0422-seed-route/) |
| MOVE-081 | R0053 clears the static list before its seed shortcut; the dynamic entry and both extractors require an explicit input-list premise. | High / Medium | ● active | [EXP-0422](../experiments/EXP-0422-seed-route/) |
| MOVE-082 | In the selected zero-radius R0178 static-search tail, substitute=seed and no substitute both take zero count: resolved=current cell, mover+90=1 and mover+98=1. | High / Medium / Unknown | ● active | [EXP-0422](../experiments/EXP-0422-seed-route/) |
| MOVE-083 | A centred zero-radius R0178 request equal to the current cell returns before search and preserves route lists and prior mover flags. | High / Unknown | ● active | [EXP-0422](../experiments/EXP-0422-seed-route/) |

### MOVE-080

+4/+8/+c of the static embedded list at actor+15c are first/last/count;
the dynamic list starts at actor+178. Traversal follows node+0 toward the
resolved endpoint. Older movement prose calls these first/last ends tail/head.
The offset and direction relations stand; this card states its naming.

Static R0437 calls real NewNode at L06984, stores the endpoint word
at L06985, and links it at L06986..L06987. Its seed test at L06988 and
L06989 reaches L01955, then L06990. That tail removes the first node,
updates both end pointers, decrements count at L06991 and stores it at
L06992. Zero count reaches full pool teardown at L01958..L06993.
Dynamic R0438 has the corresponding operations shifted by 310 hex.
NewNode R0052 increments the original embedded count at L06994/L06995;
the pool helpers R0540 and R1355 execute, with only heap malloc/free cut.

Private replay crosses both extractors, empty/old input lists, zero through
three transitions and two memory placements: 32 rows. Initially empty seed
extraction yields first=last=0, count=0, and no stored cells. The ordinary
controls retain exactly the one, two or three non-seed cells. A pre-existing
one-node list retains that old node after seed extraction; extraction does
not clear its input. Allocations are A5-filled, not silently zeroed nodes.

**Confidence.** High for the conditional instruction/list relation. Complete
byte binding, real count/link operations, independent memory walks and
retained-seed/omitted-decrement/backlink loss controls exclude a sentinel
node, a retained seed and count-only normalization. The input list and heap
services are explicit.

**Unknown.** Native allocation failure, corrupt list inputs, and actor
reachability are unmeasured. No universal route result is claimed without
the input-list premise.

### MOVE-081

R0053's static arm clears actor+160/+164/+168/+16c/+170 at
L06996..L01947, before coordinate equality at L06997..L01950. Equality
returns at L01951 without calling either extractor. The dynamic arm skips
that static teardown. Its requested-seed shortcut preserves an old dynamic
list. Picker A's nonzero seed answer instead reaches the relevant extractor
through L01941 static orL01943 dynamic. A picker miss clears the static
list through L01944 but returns directly on the dynamic arm.

The actual driver/picker/search replay crosses both paths, empty/old lists,
four topologies and two placements: 32 rows. Requested seed, substitute=seed
and no substitute all yield zero nodes from empty input. An ordinary admitted
endpoint 1013 retains 1011,1012,1013 from seed 1010. With an old dynamic cell
1414, those three boundary cases retain 1414; the ordinary route has that
old suffix. Static search clears the old cell in all four cases.

**Confidence.** High for the named local teardown/shortcut/extraction relations.
Medium for reaching the selected topologies under footprint 1, domain 2,
synthetic owner/geometry and declared allocation/notification services. The
code images are identical, one population. Actual helper execution and the
old-list control discriminate an unconditional dynamic-clear model.

**Unknown.** Native caller input-list invariants, notification side effects,
other footprint/domain branches and picker B are unmeasured.

### MOVE-082

R0178 calls actual static search at L06802. It clears mover+90 at
L06947, writes the requested cell to+74 at L06998, and tests count at
L06999..L07000. Zero count loads the current Position cell into+76 at
L06957 and sets+98=1 at L01720. Nonzero count takes the resolved cell
from actor+164's node+8 at L07001..L01721; it does not clear+98.
The following local teardown clears the dynamic list. The comparison at
L06953 then sets+90=1 at L00457 and returns when+76 equals current cell.

Eight full-entry caller replay rows cross four topologies and two placements.
Both substitute=seed for blocked goal 1011 and no substitute for blocked
goal 3030 leave empty static/dynamic lists, requested goal in+74, seed 1010
in+76, and+90/+98 both 1. Ordinary goal 1013 leaves three static cells,
+76=1013,+90=0 and unchanged prior+98. Its next dynamic-refresh call is
the stopping boundary, not a witnessed downstream action.

**Confidence.** High for local branch use given the count and Position inputs.
Medium for topology reach under the declared private services. The
three-node control and wrong-end loss exclude selection from the first node.
Unknown for native actor/order continuation.

**Unknown.** A zero-count admitted endpoint and true search failure fit the
same caller outputs. The flags alone do not identify semantic search success.
Native reachability, notification effects and the subsequent order result
remain open. AI-335's nonempty-substitute clause is partially retracted.

### MOVE-083

R0178 samples current Position cell and the requested cell, then tests
Position fractions for 128/128 at L06897..L07002. With equal cells the
computed distance is 0. Range argument 0 reaches L07003's branch directly
toL07004, before the search call and all selected route/flag writes.

The requested-seed full-entry controls preserve supplied static/dynamic old
lists, mover+74=5555,+76=6666,+90=7 and+98=0. The deliberate late-shortcut
loss instead enters search and produces+90/+98 both 1. This distinguishes
a caller-entry shortcut from extraction of a seed endpoint.

**Confidence.** High for the centred zero-radius local entry branch and its
skipped stores. Unknown for native invocation of that input and later order
behaviour. No nonzero-range or mid-transit arm is generalized.

**Unknown.** Native reachability, actor/order cadence, callback effects and
final actor behaviour remain unobserved. A reachable original breakpoint
trace of the same centred zero-radius entry that reaches search or rewrites
these fields refutes this local prediction.

`EXP-0434` was allocated ids `84`..`86` of `claims/move.md` (3 ids) and spent all three,
`MOVE-084`..`MOVE-086`. Its `MAGIC-229`..`MAGIC-234` ids are returned unused because
`claims/magic.md` is still in the table format. The next free `move.md` id is `MOVE-087`.

## Area layers and the cost byte

Addresses, offsets and cell values below are hexadecimal unless a count or a cost is named. Every
quoted instruction is asserted byte for byte on both editions' `rom.exe`, which are identical.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-084 | A transit start reads the cell costs with R1088 before any R0453 recompute, and every actor tick runs before every area-effect tick; the recompute comes later, at the cell-crossing step. | High / Medium / Unknown | ● active | [EXP-0434](../experiments/EXP-0434-area-cost-order/) |
| MOVE-085 | R1087 divides a layered cell's cost byte by four once per read, whatever the layer count, Wall of Earth included; two reads with no recompute between them divide twice, and no call-graph edge forbids that. | High / Medium / Unknown | ● active | [EXP-0434](../experiments/EXP-0434-area-cost-order/) |
| MOVE-086 | The cost byte is restored from payload+0 or the CostCracked constant and no cost plane is saved, so a loaded layered cell keeps its ingest value until a recompute; two inline readers divide by it with no zero guard, and decay can zero it. | High / Medium / Unknown | ● active | [EXP-0434](../experiments/EXP-0434-area-cost-order/) |

### MOVE-084

R0054 (transit start) runs in this order. When the word at movement-record `+0x80` differs from
the target-cell word at `+6`, it calls R0046 at `L07005` (claim clear, skipped when `+0x80`
is 0, `L07006`) and R0257 at `L07007` (claim set), then R1356 at `L07008` (heading
for the target cell, stored at record `+1`). When record byte `+0` equals `+1` (`L07009`) it calls
R1088 at `L06858` and then R0047 (the step) at `L06896`. Otherwise it calls
R0056 at `L07010` and reads no cost.

The claim routines write the dynamic plane, not the cost plane: set bit 0x80 (`L07011`)
and bit 0x40 (`L07012`); clear bit 0x80 (`L07013`, mask 0x7f) and bit 0x40 (`L07014`, mask 0xbf). No
direct-call chain from R0046, R0257 or R1356 reaches R0453, R1087,
R1088, R1075 or R1078 (`callgraph.txt`). R0046 and R0257 each make
two indirect calls on the actor; vtable `+0x1c` and `+0x20` of the three actor vtables are
R0256 and R1344, which return actor bytes `+0x49` and `+0x4a` (`L07015`, `L07016`).

R1088 divides by the mean of two R1087 reads (`L05500` source cell, `L05501`
destination cell). The mean is `(a+b)>>1` in 8 bits and a result of 0 becomes 8 (`L05502`..`L05504`),
so this divide cannot see a zero. The speed term is clamped to 63 (`L07017`) and the per-tick step
is that term times a table value, times the double `0.707` at `L06908` when both table values are
non-zero (`TERR-MOVE-056`). R0047 adds the step
bytes (record `+0xb0`, `+0xb1`) to the position and compares the cell bytes (`L07018`). Equal cells
jump past every recompute (`L07019`); a changed cell runs R0057 (`L07020`, release),
R0058 (`L02037`) and R0050 (`L02042`, occupy); the release and the occupy reach
R0453 (`L07021`, `L07022`, `L07023`). R0054 is called only by R0178
(`L06899`) and R0043 (`L01907`), and R0043 sends any position whose sub-cell bytes
`+4` and `+5` are not both `0x80` to R0039 instead (`L01890`..`L06960`); R0178's
equal test is `MOVE-083`'s. A transit start therefore begins at sub-cell `0x80/0x80`. A step is
the speed term, at most 63, times a direction factor, so it is at most 63 units on an axis (44 on a
diagonal with the `0.707` factor) and cannot change the cell from `0x80`. The crossing recompute follows the cost
read by at least one tick.

R0193 calls R0426 once per tick (`L01866`). That routine calls vtable `+0x18` of
each element of the list at `this+0x2c` (`L07024`), tests byte `+0x136` after each call
(`L07025`), and only after the list is exhausted calls R0641 (`L03066`), which calls
vtable `+0x18` of each area effect (`L03082`). The direct-caller closure of R1088,
R1087 and R0047 has one root, R0037, in `+0x18` of the vtables at `L00001`,
`L00002` and `L00003` (`L03030`, `L03031`, `L03032`). The closure of R1075 and
R1078, the layer add and removal, has one root, R0629, in `+0x18` of the vtable at
`L05352` (`L05390`). Within it the layers are laid when effect byte `+0x48` is 0 (`L05625`,
`L05598 CALL R1086`; `+0x48` is set to 1 at `L05626` and `L05627`) and removed by
R0637 when the word at `+0x4c` reaches 0 (`L05599`, `L03058`).

**Confidence.** High for the call order inside R0054, the dynamic-plane stores of the claim
routines, the getter bodies and the tick order, each a quoted instruction. Medium for the absence of
a direct-call chain from the claim routines to R0453, since each makes two indirect calls
that are not followed, and for the closures
and the cell-crossing bound: the instrument is direct `CALL rel32` plus aligned data dwords, so
indirect calls, other `vtable+0x18` call sites, the identity of the `this+0x2c` list's elements with
the actor vtables' objects and the direction tables at `world+0x58eb0` and `world+0x58eb8`, which `TERR-MOVE-056` records as
immediates written by the world constructor, are not re-read here. The sub-cell bound assumes their
entries have magnitude at most 1.

**Unknown.** Observed ticks. An effect added to the area list during the actor walk of the same
tick. On which tick of a transit the crossing falls, which depends on the speed.

### MOVE-085

R1087 returns the plane byte unchanged unless static bit 5 (`L05496`, a bit test of 0x20) is set.
For a cell with a record it zero-fills 13 dwords at `ESP+0xc` (`L07026`), copies the 13 dwords of the
record there (`L07027`), reads the byte at `ESP+0xe` (`L05497`) and, when it is non-zero, runs
a shift right by 2 and stores the result to the cost plane (`L05498`). Scratch `+2` is payload `+2` (the
copy is the record itself). The divide runs once per read and does not use the count's value.

Payload `+2` is written at six sites (`payload2.txt`): zeroed then incremented once per non-null
slot of the six in R1075 (`L07028`, `L07029`, second arm `L07030`, `L07031`) and
R1078 (`L07032`, `L07033`); a seventh store, `L07034`, rewrites the byte read at
`L07035`. The one-byte displacement-2 operands over `L07036..L07037` also read it at `L07038`,
`L07039`, `L07040`, `L07041`, `L07042` and `L07043`.

Wall of Earth is spell 19, layer index 3 at payload `+0x20` (`MAGIC-MAPLAYER-040`, `TERR-CELLREC-146`).
R0453 walks the six slots from `+0x14` and runs the shift left by 2 of the byte at `L05495` once per
non-null one (`L05493`..`L07044`), so the Wall of Earth slot takes the multiply like any other,
and R1087 reads only the count, so it takes the divide the same way. A separate test of that
slot at `L05481` ORs bits into the static and dynamic planes and does not touch the cost byte. The recompute first resets
the byte from payload `+0` (`L05491`), or from the `CostCracked` constant at `world+0x5417b` on the
footprint-clearing arm (`L07045`, `L07046`; `TERR-STRUCT-071`), so a recompute gives that source
`<< 2` per layer, in 8 bits, whatever was read before. The shipped `CostCracked` value is 6, inside
the range of the table below.

`costbyte.tsv` applies both routines to the shipped range 6..16 (`TERR-COST-052`). One layer: read 1
returns the baseline, read 2 returns `c >> 2` (1..4 for 6..16), read 3 returns 0 except for 16 (1).
Two layers: read 1 returns `c << 2` in 8 bits, so 16 gives 0. Three layers give 0 for 8, 12 and 16.

R1088 calls R1087 for the two cells only on its domain-1 arm (`L07047`). The
recompute call sites are 16 direct calls in the map module (`xref-recompute.txt`): the crossing
release and occupy (R0057, R0458), the area module (R1075, R1078), the
building routines (R1357, R1358), the sack routines (R0938, R0447,
R0227) and R1359. R0057 skips the recompute when the actor's slot is empty (`L07048`,
`L07049`); R0458 skips it when the slot is taken (`L02040`, `L02041`). From a transit
start the call graph reaches a recompute of its two cells only through R0047, at the crossing.
No edge orders another actor's transit start between the read and that recompute, and none forbids
it.

**Confidence.** High for the scratch identity, the writers within the instrument, the one divide
per read, the slot-blind divide and multiply, and the table, all quoted instructions and arithmetic.
Medium for the absence of a forbidding edge: its population is the bodies of R0054,
R1088, R0047 and the closures of `MOVE-084`; the guards in R0178 and
R0043 on a destination another actor has claimed were not read. The writer sweep misses word
plus dword stores that overlap the byte and repeated-move record copies.

**Unknown.** Whether two reads of one cell without a recompute happen in play. A runtime trace of
R1087 and R0453 over one cloud confirms or refutes the decay.

### MOVE-086

The cost byte is written from payload `+0` by R0453 (`L05491`), from the `CostCracked`
constant at `world+0x5417b` by the same routine's footprint-clearing arm (`L07045` loads the byte at `world+0x5417b`, `L07046` stores it to the cell; `TERR-STRUCT-071`), and by `L05851` in
R1078 (from the byte load at `L07050`), and the sweep of byte operands of the
form `[reg+reg]` over `L07036..L07037` (`byteidx-map-module.txt`) shows further stores at
`L07051`, `L07052`, `L07053`, `L07054`, `L07055`, `L07056`, `L07057`, `L07058`,
`L07059`, `L07060` and `L07061`, left unclassified here. The sweep does not match the write of
the ingest R0469 (`TERR-COST-052`), which uses another operand form. Baselines are read from
the plane when a record is made (`L07062`, `L07063`, `L07064`, `L07065`). A cell whose static
bit 5 is clear has no record and no divide (`L05496`).

No cost plane is saved. The sweep finds no displacement-0 byte operand between `L07066` and `L07062`, a span that holds
R1360, and `SAV-BLOCK-011` and `SAV-CELLREC-017` describe only the two flag planes and the
52-byte records with baseline, count and slots. LOAD in R0414 builds the terrain first
(`L07067 CALL R0118`, ingest), then runs R1360 (`L07068`). After it the direct-call
chains to R0453 are `L07069` R0270 > R0428 > R0262 > R0227
and `L01841` R0067 > R0865 > R0057 (`callgraph-load.txt`): the sack cells
and an actor release. A tick lays layers only for an effect whose `+0x48` is 0 (`MOVE-084`). A saved
layered cell therefore keeps the ingest byte, with a non-zero count, until a recompute; its first
R1087 read returns `c >> 2` and stores it.

Two inline readers divide by the byte with no zero test. R0039 runs after R0047 and
reads the cost of the cell word (`L07070` loads the cost byte, `L07071` divides by it as a signed division; the arm is
selected by `L07072` and `L07073`). R1345, called at `L07074` before the occupy
recompute, does the same (`L07075`, `L07076`). `IDIV` by 0 raises a divide error. The byte is 0
after a recompute for two layers at cost 16 and three layers at 8, 12 and 16, and by decay with a
non-zero count after read 3 of a one-layer cell at cost 6..15 and after read 4 of a two-layer cell
at cost 6..15 (`costbyte.tsv`), so a layer count of one or two reaches the zero divisor if
two or three reads of the cell occur with no recompute (`MOVE-085`).

**Confidence.** High for the two divide sites and their missing guards and for the load sequence.
Medium for the restore sources, the absence of a cost plane in the serializer and the population of
readers and writers: the instrument is the `[reg+reg]` sweep over the map module, classified by
hand, with no original run, no corpus read and whole-image writes outside `L07036..L07037`
unclassified. The `[reg]` displacement-0 store form (about 85 byte stores and read-modify-writes in
`L07036..L07037`, among them `L05492`, `L07046`, `L05499` and `L05852`) was not swept or
classified; only the sites quoted here are named. Other index forms are not seen.

**Unknown.** Whether LOAD restores the area effects with `+0x48` set, so that no layer is laid again; the effect save routine was not read. What the process does on a zero divisor, since no exception handler was read.
Whether a decayed byte, 0 after read 3 or 4, reaches an inline reader; that needs the double read of
`MOVE-085` and a mover starting or crossing in that cell. Whether a
shipped or generated save holds a layered cell. Whether two or three layers meet on a cost-8 or
cost-16 cell in play; `MAGIC-MAPLAYER-040`'s conflict rules do not forbid it.

## Footprint position and cell transit

Addresses are hexadecimal. Every quoted instruction is asserted byte for byte on both editions' `rom.exe`,
which are identical (162 rows, 0 mismatches). The instrument is capstone disassembly of address ranges plus
direct `CALL rel32` cross-references and a linear byte-store sweep; indirect calls are not followed.

`EXP-0443` was allocated ids `87`..`89` of `claims/move.md` (3 ids) and spent all three, `MOVE-087`..`MOVE-089`.
The next free `move.md` id is `MOVE-090`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-087 | A mover of footprint side n stores the footprint's top-left cell and a sub-cell offset; the fine point P = cell*256 + sub is that corner, and the centre read for range, edge gap and bearing is P + (n-1)*128 per axis. | High / Medium / Unknown | ● active | [EXP-0443](../experiments/EXP-0443-footprint-slot/EXP-0443.md) |
| MOVE-088 | A cell crossing releases the old footprint, claims, rewrites the position and occupies the new one in one R0047 call; an empty release slot or a taken occupy slot skips that cell's recompute, and no deferral exists in those bodies. | High / Medium / Unknown | ● active (amended) | [EXP-0443](../experiments/EXP-0443-footprint-slot/EXP-0443.md) |
| MOVE-089 | Actor removal through R0865 releases the footprint at the stored position, recomputing each held cell, and clears the claim bits only after every release succeeds; release is blind to which actor holds the slot. | High / Medium / Unknown | ● active | [EXP-0443](../experiments/EXP-0443-footprint-slot/EXP-0443.md) |

### MOVE-087

The position record at `actor+0x10` holds byte `+0` cell x, `+1` cell y, word `+2` the packed cell, plus bytes `+4`
and `+5` the sub-cell x and y. `R0165` returns `(cell x << 8) + sub x` and `R0166` the same for y
(`L07077`..`L07078`). Call that 16-bit value P. The footprint side n is actor byte `+0x49`, read through
vtable `+0x1c` (`MOVE-084`).

The stored cell is the origin of the occupied cells. The occupy routine `R0050` reads n once
(`L05505`), then calls `R0458` for every offset pair 0..n-1 added to the cell bytes from
`R0299` and `R0300` (`L07079` and `L07080` add the frame-local offsets). The step
release loop in `R0047` adds the same offsets to the stored cell bytes (`L07081`, `L07082`), and the
claim loops of `R0058` run 0..n-1 from the packed cell (`L07083`..`L07084`). The footprint is
therefore the n by n block whose top-left cell is the stored cell. A centre-cell anchor and a stored footprint
centre are both excluded in these bodies: the step, occupy, release and claim loops read here add offsets 0..n-1 to
the stored cell and none subtracts one.

`R0862` returns the x centre and `R0863` the y centre: P, plus `(n-1) << 7`
(the virtual call through vtable slot `0x1c` at `L07085`, the subtraction of 1 at `L07086`, the shift left by 7 at `L06157`, and the add at `L06158`), in 16 bits.
Three consumers read it:

- `R0247`, the range test, takes the absolute difference of the two x centres and of the two y centres
  (`R0376` is `NEG` on a negative argument), keeps the larger, subtracts `((n1+n2) << 7) - 0x100`
  (`L04080`, `L04081`), and returns 1 when the result is at most `0x180`, else `(v + 0x40) >> 8`
  (`L04082`, `L07087`, `L04083`).
- `R0036`, the edge gap, builds each centre as `(2*cell + n + 0x1ff) << 7` plus the sub-cell, masked to
  16 bits (`L07088`, `L07089`, `L07090`, `L07091`). Since `0x1ff*128 = 0x10000 - 128`, that value is
  P + (n-1)*128 modulo 65536, the same centre. It subtracts the axes, takes the absolute value, subtracts
  `(n1+n2) << 7` (`L07092`), clamps at 0, keeps the larger axis, and returns `(v >> 8) + 1` (`L02064`,
  `L02065`).
- `R0051`, the bearing, subtracts the centres from the same two routines (`L06155`..`L06156`).

The direct-call readers of these routines are `R0247` at `L07093`, `L02044`, `L07094`;
`R0036` in `R0021`, `R0022`, `R0306`, `R0126`, `R0109`, `R0041`
and `R0043`; and `R0051` in 11 owners (`xref-position.txt`).

A step adds the signed step bytes to P itself (the adds at `L07095` and `L07096`) and stores the
resulting cell bytes, sub-cell bytes and packed word back (`L07097`..`L07098` when the cell is unchanged,
`L02038`..`L02039` at a crossing). The same offset moves the corner and the centre whatever n is. At arrival
both sub-cell bytes return to `0x80` (`L07099`..`L07100`), as they do in the constructor and in the placement
and portal writers (`L00968`, `L00973`, `L00974`). A resting mover therefore has centre
cell*256 + 0x80 + (n-1)*128: for n = 2 that is the grid line shared by its two cell columns, and for n = 3 the
middle of its middle cell, which is the geometric centre of the block in both cases.

Fourteen routines store the four position bytes within 24 instructions of one another (`posstores.txt`):
`L07101`, `R0287`, `R0288`, `R0289`, `L07102`, `R0255`, `L06144`, `R0290`, `R1010`,
`L05872`, `R0469`, `R0047`, `R0039` and `R1055`.

**Confidence.** High for the centre formulas, the footprint origin and the step arithmetic, each an asserted
instruction sequence read whole, and for excluding a centre-cell anchor and a stored centre in the step, occupy, release and claim loops read. Medium for the
writer list and the reader lists: the instrument is a linear sweep with a 24-instruction window and direct
calls only, and it misses stores split over longer windows, changed base registers, `REP MOVS` copies of the
record, computed indices and indirect callers.

**Unknown.** Which shipped or authored actors have n above 1; no census was run. Observed positions of a size
above 1 in play. The full bodies of the 13 writers other than `R0047`, which were located and only their quoted stores asserted.

### MOVE-088

`R0047` (the step) runs, when the cell bytes differ after the step is added (`L07018`, a byte compare):

1. a release loop over the n by n footprint at the old position, one `R0057` call per cell
   (`L07020`), which stops at the first call that returns 0 (`L07103`, a zero test, then the equal branch to `L07104`);
2. `R0058` with the old packed cell (`L02037`), which stores the claim cell at `mover+0xa6`
   (`L07105`) and ORs the dynamic bits over the n by n footprint (`MOVE-084`);
3. the new cell bytes, sub-cell bytes and packed word (`L02038`..`L02039`);
4. `R0050`, the occupy (`L02042`). Its return value is not tested before the arrival recentre at
   `L02043`..`L07100`.

All four are in one call of the step, and the loop never yields. The two release outcomes decide the recompute
of one cell. `R0057` selects the slot from the actor's movement domain (the virtual call through vtable slot `0x20` at `L07106`):
domain 1 or 2 uses record `+0x04`, domain 3 uses `+0x08`. A missing cell record returns 0 at `L07107`, before any
slot test. A domain of 0 or above 3 takes neither arm (`L07108`, `L07109`): it skips the slot test and runs the
recompute at `L07021` unconditionally. A slot that is empty returns 0 at `L07110` or `L07111`, before the
clear and before the recompute. A slot that holds any
non-zero value is cleared (`L07112`, `L07113`) and `R0453` runs at `L07021`; the routine then returns
1. The slot test is a non-zero test and does not compare the held pointer with the actor, so the release clears
a slot another actor holds.

The occupy `R0050` calls `R0458` for each cell in row-major order and continues only on a
return of 1 (`L07114`, a zero test, then the unequal branch to `L07115` at `L07116`); the first return of 0 ends it with 0 and the
later cells are not entered (`MOVE-087` gives the order). In the domain 1 and 2 arm of `R0458`, a cell record whose byte `+0x2c` is non-zero and not `0x1a`
(`L07117`, `L07118`) first builds a temporary caster through `R1137` or `R1138`
(`L06049`, `L06050`; `MAGIC-235`); the builder appends a new object to the list at `this+0x2c` (`L06032`).
That happens before the slot test, whether or not the slot is taken. The domain 3 arm (`L07119`..`L07120`)
contains no such call. A slot that is already taken (`L02040` compares the slot at `+4` with zero, `L02041` the slot at `+8`) calls
`R0459` and returns 0 with no slot store and no recompute. `R0459` and `R1361` are each one return popping 4 argument bytes (`R1361`, `R0459`). An empty slot is
written with the actor (`L07121`, `L07122`) and `R0453` runs at `L07022` or `L07023`.

A deferred recompute needs a stored mark or a queue entry that a later routine reads. The release, the occupy,
the per-cell routine, the stubs and the step contain no store of that kind. Their stores are the slot, the
write-back of the record copy through `R1362`, the position, the claim cell, the cached bytes at
`mover+0x86`..`+0x89` (`L07123`..`L07124`) and, in the per-cell occupy, the trigger caster's construction and
list append. None is read here as a recompute mark; the consumers of that list were not read. The recompute therefore runs in the tick of the crossing, in the same call, for every
cell whose release found a held slot and every cell whose occupy found an empty one. A cell whose release found
an empty slot, or whose occupy found a taken one, gets no recompute from that call, and the first such
release or occupy skips the cells after it in that footprint. The mover's `+0x76` word and `R0469`, which
has its own release, claim and occupy call sites (`L07125`, `L07126`, `L07127`), were not read.

**Confidence.** High for the order inside the step, the two slot branches of the release and of the per-cell
occupy, the no-op stubs, the loop exits and the untested occupy return, each an asserted instruction. Medium
for the absence of deferral: the population is the bodies named above and the two stubs; stores outside
them, the consumers of the trigger caster list, indirect callers, and any routine that reads the bytes at `mover+0x86`..`+0x89` were not enumerated.

**Unknown.** An observed tick. Which tick of a transit holds the crossing, which depends on speed
(`MOVE-084`). Whether `R0469` follows the same order.

**Amended.** The unread `R0469` clause is closed by `MOVE-093`: Ghidra's `R0469` is an unrelated cost helper, and the second release, claim and occupy call sites (`L07125`, `L07126`, `L07127`) belong to an unreferenced run at `L07128`..`L07129` that follows the same order. `MOVE-094` gives the crossing tick for the Unknown above. The rest of the claim stands.

### MOVE-089

`R0865` is the removal of an actor from the cell records. It reads n through vtable `+0x1c`
(`L07130`) and runs the same row-major release loop as the step, at the stored position, one
`R0057` call per cell (`L07131`). A return of 0 ends the routine with 0 (`L07132`, the equal branch to `L07133`), so
the claim state below is skipped. After every release has returned 1 it reads the word at `mover+0xa6`
(`L07134`). A zero word, or a word equal to the packed cell at position `+2` (`L07135`), ends the routine
with 1. Any other word is passed to `R0046` (`L07136`), which clears the dynamic bit (`0x80` for
domain 3, `0x40` for domain 1 or 2) over the n by n cells at that claim cell except those inside the current
footprint (`L07013`, `L07014`, `MOVE-084`); `+0xa6`, `+0x80` and `+0x76` are then zeroed
(`L07137`, `L07138`, `L07139`).

The four direct callers are `L07140` in `R0067`, `L07141` in `R0121`, `L07142` in the
teardown `R0208`, and `L07143` in `R1132`. The teardown then calls `R0866` at `L07144`,
which sets both sub-cell bytes to `0x80`. When every footprint cell's release finds a held slot, each is cleared with its recompute, so a removed mover
leaves no occupancy in those cells and no cost contribution: the alternative that the mover's contribution stays
in the plane after removal is rejected for that all-cells-held case, and a mover removed after a completed
crossing is fully released by this call. A release that finds an empty slot ends the routine early and leaves the
later cells' slots and the claim bits `0x40` or `0x80` as they were. A crossing is not a state that can be cut short between ticks, since release, position rewrite and
occupy are one call (`MOVE-088`).

A footprint that was refused part-way keeps its row-major prefix (`SAV-CELLFAIL-583`). The release at removal
walks the whole footprint at the new corner. Where a refused cell is held by another actor, the non-zero slot
is cleared and recomputed, since the release tests no identity; the walk then reaches a never-entered cell,
finds it empty, and returns 0, leaving the claim state in place. That consequence is read from the bytes and
not run.

**Confidence.** High for the loop, the exit on an empty slot, the claim clear and its conditions and the
teardown recentre, each an asserted instruction. Medium for the consequences on a refused footprint, which
combine these bytes with `SAV-CELLFAIL-583` and `MOVE-088` without a run, and for the four callers, a direct
`CALL rel32` list that misses indirect callers.

**Unknown.** Whether any caller tests the return of 0. Whether each caller removes the actor in the tick in which
it leaves the world. Whether two actors reach one cell in play and one of
them is then removed. Native behaviour of the claim bits after the early return.

## Walk step at order recovery

Evidence is a static read of `rom.exe` (one image on both lawful installs) in Ghidra 12.1.2 headless; no process
was run. `EXP-0451` was allocated ids `90`..`92` of `claims/move.md` and spent `MOVE-090`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-090 | A fresh walk call of R0178 writes a sub-cell step in that call only when the facing byte already equals the direction to the first path node; otherwise it turns and the step comes at the next call. | High / Medium / Unknown | ● active | [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md) |

### MOVE-090

- Entry: executor row 1 at `L07145` calls `R0178(actor, ord+0x0a, 0)` at `L01932`. `R0040` is the centred test (both sub-offsets 0x80); a mover that is not centred stores progress 3.
- A fresh walk forces a full search; the dynamic list is empty, so the near search `R0055` runs and then the stepper `R0054`.
- The stepper computes the facing to the head node with `R1356` into `mover+1`. When `mover+0` equals `mover+1` it calls `R1088` (step vector `mover+0xb0` and `mover+0xb1`, count `mover+0xaa`) and `R0047`, which writes the sub-cell position at `+4` and `+5` in that call. Otherwise `R0056` turns only; it snaps the facing at once when the difference is under 0x21, and the step is written by the next call. An empty dynamic list calls `R0249`, a turn only.

**Confidence.** High for the control flow, each branch read from `rom-walk.txt`. Medium for the composed outcome at the recovery tick, since the facing at that tick is a run-time value.

**Unknown.** The facing byte's usual value when a recovered actor was last stopped; whether the turn path has a second step in the same call under any condition not shown in the listed routines.

**Evidence.** [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md), `evidence/listings/rom-walk.txt`, `evidence/listings/rom-row-callees.txt`

## Second step routine, transit ticks and teardown claim state

Evidence is a static read of `rom.exe` (one image on both lawful installs) with a capstone sweep; no process was run. `EXP-0452` was allocated ids `93`..`96` of `claims/move.md` and spent all four.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-093 | The step routine at L07128..L07129, with no direct reference found, releases the old footprint, sets the claim at the old cell, rewrites the position, then occupies, with no arrival compare and no test of the occupy return. | High / Medium / Unknown | ● active | [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md) |
| MOVE-094 | In the live step a transit from the centre crosses on tick ceil(128/s) on a positive axis and floor(128/s)+1 on a negative one, start tick 1; the step calls release and occupy on that tick only, whatever the slots hold. | High / Medium | ● active | [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md) |
| MOVE-095 | A removal whose release ends early on an empty slot returns before the claim clear: the transit's claim bits and the later cells' slots stay as they were; no store of the claim bits follows in the teardown body or R0866. | Medium / Unknown | ● active | [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md) |
| MOVE-096 | A unit with health at or below 0 is not stepped; its slot, claim bits and +0xa6 stay frozen through the death countdown, and the teardown releases at the stored cell and clears the claim when +0xa6 is not the current packed cell. | High / Medium / Unknown | ● active | [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md) |

### MOVE-093

`MOVE-088` left `R0469` unread. Ghidra's `R0469` ends at `L07146` and is an unrelated cost helper; the second step routine is the code after it, `L07128`..`L07129`, which has no function entry. Its order, from `evidence/listings/rom-step.txt`:

1. the move computation, reading record `+0x72` (`L07147`, `L07148`); a crossing is a flip of bit 0x80 of the moving axis's sub-cell byte together with a cell change (`L07149`..`L07150`);
2. on a crossing, a release loop over the n by n footprint, one `R0057` call per cell (`L07125`, row-major), stopping at the first return of 0 (`L07151`, `L07152`);
3. `R0058` with the old packed cell (`L07126`);
4. the position rewrite: cell bytes, sub-cell bytes, packed word (`L07153`..`L07154`);
5. `R0050`, the occupy (`L07127`), then a return popping 4 argument bytes. The return is not tested.

The stride `+0x72` is recomputed at the crossing (`L07155`..`L07156`). A branch with no crossing writes the position only (`L07157`); a same-cell half-flip from an off-centre start sets both sub-cell bytes to 0x80 (`L07158`..`L07159`). The routine has no `+0xaa` arrival compare.

The order matches the live step of `MOVE-088` (release `L07020`, claim `L02037`, position, occupy `L02042`). References: `evidence/listings/orphanrefs.txt` finds no `E8`, `E9` or `0F 8x` rel32 into the range from outside it and one dword in the whole file with a value in the range (file offset `0x174d83`, inside a `CALL` operand). `xref.txt` and the displacement scan `scan.txt` (value 0x72) find no other reader of record `+0x72` in `L06029..L07160`; stores to it are at `L07161`, `L07156`, `L07162`, `L07163`, `L07164`.

**Confidence.** High for the order and the untested occupy return, each asserted instruction on both roots (`asserts-en.txt`, `asserts-ru.txt`). Medium for "unreferenced": the population is direct relative transfers and literal dwords over the file; a computed jump into the range was not excluded.

**Unknown.** Whether the range is ever executed. Whether a pre-release build called it.

**Evidence.** [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md), `evidence/listings/rom-step.txt`, `evidence/listings/orphanrefs.txt`, `evidence/listings/xref.txt`

### MOVE-094

The live step `R0047` runs once per tick: `R0054` on the start tick (`+0xac` reset to 0, then the step) and `R0178` then `R0039` on later ticks while the sub-cell bytes are not both 0x80 (`L02035`, `L02036`). It calls release, claim and occupy only on the call whose cell bytes differ after the step is added (`L07018`); otherwise it writes the position and jumps to `L02043`.

Starting from sub 0x80 with signed axis step s (`R1088`, clamp 1..63, `TERR-MOVE-056` gives the shipped range 4..32), the crossing tick counted with the start tick as 1 is ceil(128/s) for a positive axis and floor(128/s)+1 for a negative one; the transit length is N = ceil(256/s), where the `+0xaa` compare recentres to 0x80 (`evidence/listings/crossing.txt`, v = 1..63). At v = 16: N = 16, positive tick 8, negative tick 9. Anti-diagonal directions 1 and 5 add 0xffff to the fine X when both sub-cell bytes are 0 (`L07165`..`L07166`); diagonal steps are `R1088`'s truncated v*0.707, so the table is exact for the four axis directions only.

The slot outcomes are those of `MOVE-088`: an empty release slot returns 0 and skips that cell's recompute; a taken occupy slot returns 0 with no slot store and no recompute. Neither defers: the step calls them on no other tick (`MOVE-088`). The scope is the step `R0047`, not every caller of the occupy or the claim clear.

**Confidence.** High for the tick rule given the start state and the step bytes: the compare and the add are asserted instructions and the table is arithmetic over them. Medium for the diagonal and tie-break paths, read but not tabulated.

**Unknown.** The arrival path of `R0039`, which the walk calls on later ticks: it calls the claim clear `R0046` at `L07167` and `L07168` and the occupy `R0050` at `L07169` (the occupy when the cell byte is `0x1a`, `L00160`, `L00973`). These sites are in `rom-transit.txt` and `xref.txt`, were read but not traced, and are not reached through a crossing.

**Evidence.** [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md), `evidence/listings/rom-step.txt`, `evidence/listings/rom-transit.txt`, `evidence/listings/crossing.txt`

### MOVE-095

`R0865` (the removal) runs the release loop over the footprint and returns 0 at the first release that returns 0 (`L07132`, the equal branch to `L07133`), before the claim clear. The clear (`R0046`, `L07136`) is reached only after the whole loop and only when `+0xa6` is non-zero and not the current packed cell. `R0866`, which follows the removal in the teardown, sets both sub-cell bytes to 0x80 and writes no plane byte.

The plane byte at `world+0x20000 + packed cell` carries the claim bits 0x40 and 0x80 (`MOVE-084`). It is rewritten whole by `R0453` (callers `L07022`, `L07170`, `L07023`, `L07171`, `L07021`, `L07172`, `L07173`, `L07174`, `L07175`, `L07176`) or cleared by `R0046` and `R1332`. None of these is called in the teardown body or in `R0866` after the early return. The claim bits at the `+0xa6` footprint therefore read as set. Slots of the cells after the empty one keep what they held: zero, or a pointer to another unit, or a stale pointer to the removed unit.

Plane readers found by displacement sweeps (`scan.txt`, values 0x20000 and 0x200): `R0144`, `R1327`, `R1328`, `R1329`, `R1330`, `R0142`, `R0178` (`L07177`), `R0048`, `R0055`, `R1287`, `R1080`, `R1081`, `R1055`, `R0409`, `R0410`.

**Confidence.** Medium: the stores and the early return are asserted; the claim bits reading composes them without a run, and no reader of the stale bits was traced to an effect.

**Unknown.** Four teardown callees after `L07142` were not read: the two virtual calls through `[edx+0x40]` (`L07178`, `L07179`), `R0929` (called twice) and `R0476`. What each plane reader does with a stale claim bit. Slot-pointer dereferencers were not enumerated. Observed play.

**Evidence.** [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md), `evidence/listings/rom-removal.txt`, `evidence/listings/rom-claim.txt`, `evidence/listings/rom-slots.txt`, `evidence/listings/scan.txt`

### MOVE-096

The actor tick `R0037` does nothing when action `+0x54` is 0x10. When health `+0x94` is at or below 0 it takes the dying branch, which does not call the executor `R0016` (its direct rel32 plus dword references are `L00086`, in the health above 0 branch, and `L00087` in `R0401`; `xref.txt`), so no walk and no step runs. The first dying tick sets `+0x13c` to 1, calls `R0014`, halves `+0xbe` and sets `+0x6c`; later ticks count `+0x6c` down. At zero, if `+0x4a` is above 1 health is set to 0xfc18; when health is at or below -10 the actor sets `+0x54` to 0x10 and calls the teardown `R0208`.

Through the countdown the occupancy slot, the claim bits, `+0xa6`, `+0x80` and `+0xac` are unchanged. The teardown calls `R0865` (`L07142`, return not tested): it releases at the stored position (the new cell if the transit had crossed), then clears the claim at `+0xa6` (the target cell before the crossing, the old cell after it) unless it equals the current packed cell (`L07135`, a compare with the packed cell at `+2`; for a footprint wider than 1 a `+0xa6` cell inside the footprint but not the packed cell is cleared), zeroes `+0xa6`, `+0x80` and `+0x76`, and `R0866` recentres. When the release ends early, `MOVE-095` applies. Other callers of `R0865` (`L07140`, `L07141`, `L07143`) were not covered.

**Confidence.** High for the dying branch's exclusion of the executor and the teardown order, asserted instructions. Medium for the stored position at death: it composes the step's store order with the freeze and no run was observed.

**Unknown.** Other writers of `+0x54`. Observed play. Whether a unit dies in the tick it crosses.

**Evidence.** [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md), `evidence/listings/rom-death.txt`, `evidence/listings/rom-removal.txt`

## Contact-ring picker geometry and dynamic labels

Evidence is a static read of `rom.exe` (one image on both lawful installs) and CPU emulation (unicorn) of the image's own bytes for `R0436`, `R0435` and their leaves `R0862`, `R0863` and `R0279`, over fake actors whose footprint side and movement class come from stubbed vtable slots `+0x1c` and `+0x20`; no process of the game was run. `EXP-0494` was allocated ids `97`..`104` of `claims/move.md` and spent `97`..`100`. Fine coordinates are `cell * 256 + sub-cell + (side - 1) * 128` per axis, the footprint centre as `R0862` and `R0863` return it; `dx`, `dy` are mover minus target.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-097 | `R0436` returns the bearing of the mover's fine centre from the target's as a 16-way code, 0 just east of north, clockwise, with sector edges at the axes, the diagonals and the 2:1 slopes. | High | ● active | [EXP-0494](../experiments/EXP-0494-pursuit-search-cadence/EXP-0494.md) |
| MOVE-098 | Picker B enters each ring on the edge facing the mover's bearing quadrant where the centre-to-centre line meets that edge's row or column; for a mover footprint of 2 to 4 near NNE or WSW the entry falls outside the ring. | High / Medium | ● active | [EXP-0494](../experiments/EXP-0494-pursuit-search-cadence/EXP-0494.md) |
| MOVE-099 | Picker B keeps the first strictly lowest label over its executed probes, walker 1 before walker 2 per step; the first ring iteration whose probes find a label ends the scan. An off-ring entry can skip ring cells and return an outer cell. | High | ● active | [EXP-0494](../experiments/EXP-0494-pursuit-search-cadence/EXP-0494.md) |
| MOVE-100 | In the dynamic search a non-seed cell whose occupancy bits meet the mover mask is not labelled by relaxation, so picker B skips it; the start cell keeps its seed label 0 and can be picked. Own bits are cleared for the wave, then reset. | High / Medium | ● active | [EXP-0494](../experiments/EXP-0494-pursuit-search-cadence/EXP-0494.md) |

### MOVE-097

- With `a = |dx|` and `b = |dy|`: for `dx > 0, dy <= 0` the code is 0 when `b > 2a`, 1 when `a <= b <= 2a`, 2 when `b < a <= 2b`, 3 when `a > 2b`. For `dx > 0, dy > 0`: 4 when `a > 2b`, 5 when `b < a <= 2b`, 6 when `a <= b <= 2a`, 7 when `b > 2a`. For `dx <= 0, dy > 0`: 8 when `b > 2a`, 9 when `a <= b <= 2a`, 10 when `b < a <= 2b`, 11 when `a > 2b`. For `dx <= 0, dy <= 0`: 12 when `a > 2b`, 13 when `b < a <= 2b`, 14 when `a <= b <= 2a`, 15 when `b > 2a`.
- The routine is integer only (`R0436`..`L13318`): two calls to each centre leaf, absolute values, and compares of one magnitude with the other and with its double.
- Emulation of the routine bytes against this rule: 107201 cases (every `|dx|, |dy| <= 40`; the lines `|dx| = 2|dy|` and `|dx| = |dy|` with offsets -2..2 up to 8000; 60000 random up to 20000), 0 mismatches. The byte range hash is in `q4-r0436-verify.txt`.

**Confidence.** High: closed form and emulation agree on every case, including each sector edge.

### MOVE-098

- Picker B folds the code to an edge `q = ((code + 2) >> 2 & 3) + 4` (`L06824`..`L06825`): codes 14, 15, 0, 1 give the north edge (row `ty - nM - r + 1` on ring r), 2..5 the east edge (column `tx + nT + r - 1`), 6..9 the south edge, 10..13 the west edge (`MOVE-ALT-020` gives the box).
- North and south: `slope = (Xt - Xm) / (Yt - Ym)` in single precision, `c = trunc(Xt - Yt * slope)`, entry x `= trunc((row * 256 + 0x80) * slope + c) >> 8`. East and west: `slope = (Yt - Ym) / (Xt - Xm)` with a zero `Xt - Xm` replaced by 1, `c = trunc(Yt - Xt * slope)`, entry y `= trunc((column * 256 + 0x80) * slope + c) >> 8`. `trunc` is `R0279` (toward zero); `>> 8` is an arithmetic shift.
- Emulation at FPU control word 0x27f: the first probe cell equals this rule in all 29992 grid runs (mover offsets -12..12 per axis, every `(nM, nT)` in 1..4, three target sub-cell choices; `q4-entry-grid.tsv` lists the 9992 centred-target runs) and in 20000 random runs (`q4-fpu-precision.txt`; also at 0x37f). At 0x07f (24-bit mantissa) 7 of the 20000 random runs enter one cell away (`q4-fpu-precision.txt`).
- Off the ring: with `nM = 1` every entry of the centred-target grid (2498 runs, offsets -12..12) lies on the ring. With `nM` of 2 to 4, codes 0 and 1 can place the north entry up to `nM - 1` columns east of the east edge, and codes 10 and 11 the west entry up to `nM - 1` rows south of the south edge (`q4-entry-grid.tsv`). Then one walker never meets its corner and runs straight outward. For `nM = 2, nT = 1` and a mover at offset (9, -12), ring 1 probes 7 cells east of the box on its north row and leaves 6 of its 12 ring cells unprobed (`q4-offbox.txt`); for `(4, 4)` 18 cells off the ring and 17 ring cells unprobed.

**Confidence.** High for the edge selection, the entry rule at 0x27f and 0x37f and the off-ring cases, all from emulation of the bytes. Medium for the entry in play: the FPU precision the original runs at was not observed, and at 24-bit precision boundary cases move.

**Unknown.** The FPU control word in a running session. Which units have a footprint of 2 or more and pursue.

### MOVE-099

- Each step probes walker 1's cell (`L06835`) and then walker 2's (`L06836`); each replaces the kept cell only when its label is strictly lower (`L13319`, `L13320`, `jge`). Walker 1 starts in the edge's direction (`world+0x54186` table index q: north edge eastward, east edge southward) and turns at each box corner with an increasing index; walker 2 runs the opposite way. The table bytes are stored by the world constructor (`L13321`..`L13322`, `ebx` zeroed at `L13323`).
- A ring iteration runs `half its cell count + 1` steps (both walkers probe the meeting cell), capped at 0x64 (`L13324`). The ring loop continues only while the kept label is 0xffff (`L06823`), up to 8 rings; no label returns 0. The minimum and the stop are over the executed probes, not the ring's cells. Walker 1 runs clockwise on screen only when the entry lies on the perimeter. With an off-ring entry (`MOVE-098`, the `MOVE-ALT-020` narrowing) one walker leaves the ring: for target (100,100) side 1 and a side-2 mover at (109,88), a lone label at (108,98) outside ring 1 is returned in the first ring iteration, and a lone label at ring-1 cell (101,100) is never probed and the routine returns 0 after 336 probes.
- Emulation (`q4-tie.txt`, target side 1, mover side 1 north or east): equal labels on the whole ring return the entry cell; equal labels at the two cells one step from the entry return walker 1's; a later cell with a lower label beats an earlier higher one; a ring-1 label 9 beats a ring-2 label 1; with no label the routine probes 304 cells and returns 0.

**Confidence.** High for the strict comparison, walker-1 priority and the stop over executed probes: the compares are cited instructions and every case was run on the bytes.

### MOVE-100

- The flag-0 search calls `R1332(world, actor, cell)` at `L06784` before the wave and `R0058` at `L01945` after it, before the picker tail. Both read the movement class through vtable `+0x20`: class 1 or 2 clears or sets bit 0x40, class 3 bit 0x80, class 0 neither, over the side-by-side footprint (`L13325`..`L13326`, `L13327`, `L13328`).
- The four dynamic relaxers `R1327`, `R1328`, `R1329`, `R1330` test the occupancy byte `world+0x20000` against the mask before any label store (`L13329`, `L13330`, `L13331` and its seven siblings, `L13332` and its seven siblings) and skip the cell when a bit meets it. The label plane is reset to 0xffff (`L01912`) in every search that passes the start-equals-goal return (`L01950`); that return leaves earlier labels in place. The start cell is seeded with label 0 (`L01965`) before the occupancy clear at `L06784`.
- Picker B reads only the label plane (`L06835`, `L06836`), so an occupied or claimed cell other than the seed is not returned. The mover's own bits are restored at `L01945` before the picker, but the seed keeps label 0: for target (100,100) and a side-1 class-1 mover at (101,100), the flag-0 search reaches picker B with the start cell occupied (0x40) and labelled 0, picker B returns the start cell, and the dynamic count is 0. The goal cell of a near search at an occupied victim cell is unlabelled, and the search goes to picker B (`AI-373`).
- Writers of bits 0x40 and 0x80: `R0058` and `R0257` set them by class, `R1332` and `R0046` clear them, and `R0453` rebuilds a cell's byte from its cell record and sets 0x40 or 0x80 when the record's `+0x04` or `+0x08` is nonzero (`L10529`, `L10530`). `MOVE-PLANE-005` names the same five and the claim protocol is the reservation `MOVE-PLANE-005` cites.

**Confidence.** High for the seed, the clear-wave-set order, the bit selection and the relaxer tests, cited instructions; the start-cell pick was run on the bytes. Medium for the mask meaning (bit 0x40 ground, 0x80 air) taken from `MOVE-PLANE-005`.

**Unknown.** What `R0453`'s record fields `+0x04` and `+0x08` hold in play.

## Turn schedule by unit kind

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-105 | A fresh turn of at most two sixteenths sets the facing byte in the calling sub-tick; a larger turn advances it by the mover rate byte per actor sub-tick and ends after ceil(arc/rate) sub-ticks. | High / Medium | ✔ promoted | [EXP-0498](../experiments/EXP-0498-hero-turn-rate/EXP-0498.md) |
| MOVE-106 | The turn rate is mover byte +0x0a; each Human derive copies the low byte of the speed word into it, and Units-table RotationSpeed is 8..23; whether a Human holds the derived value at every turn is Medium, its first turn Unknown. | High / Medium / Unknown | ✔ promoted (amended) | [EXP-0498](../experiments/EXP-0498-hero-turn-rate/EXP-0498.md) |
| MOVE-107 | A turning unit does not step; the non-self unit and point act gates need the current facing on the heading; a self-target cast skips them; a new target continues from the current byte; a stop reset leaves the active flag. | High / Medium | ✔ promoted | [EXP-0498](../experiments/EXP-0498-hero-turn-rate/EXP-0498.md) |

### MOVE-105

- The facing byte `mover+0` holds 256 steps per circle; one sixteenth is 16. `R0056` writes the desired byte `mover+1`. When the active dword `mover+0xa0` is 0 and the shortest arc is at most 32 (`L00749`, `L00750`), it writes current := desired and `mover+0xa4` = 1 in that call. Otherwise it calls the leaf `R0249`, which snaps when the arc is below the rate byte `mover+0x0a` and else moves the byte by the rate along the shorter arc, addition on the 128 tie (`MOVE-TURN-044`). The call then writes `mover+0xa4` = ceil(arc/rate) from the arc before the step and clears `mover+0xa0` when the facings match.
- Counted from the first call, a fresh turn reaches the desired byte on call ceil(arc/rate), or on call 1 for an arc up to 32. Ported arithmetic over rates 1..40 (`evidence/turnsim.txt`) gives:

  | rate | 1/16 | 4/16 | 8/16 |
  |---|---|---|---|
  | 8 | 1 | 8 | 16 |
  | 12 | 1 | 6 | 11 |
  | 16 | 1 | 4 | 8 |
  | 19, 20, 21 | 1 | 4 | 7 |
  | 23 | 1 | 3 | 6 |

  A 2/16 arc also ends on call 1 at every rate.
- One call per actor sub-tick: the actor tick `R0037` runs once per sub-tick (`HERO-DYETICK-067`) and calls the executor `R0016` at `L00086` when `actor+0x3c` is nonzero (`R0416`). The executor's turning paths call `R0056` at most once per invocation: the walk stepper's mismatch arm (`MOVE-090`), the approach `R0043` through `R0042`, pending order 0xa (`L06983`), the face helper `L13333`, the idle turn `R0205` and the face-and-mark routines at `L13334` and `R0250` (the latter from pending arm 6 at `L13335`). A sub-tick is 1/16 of a full tick: 62 ms at the default speed index (`MOVE-CLOCK-032`). At rate 16 a half turn takes 8 sub-ticks, about 0.5 s; a quarter turn 4 sub-ticks.

**Confidence.** High for the routine arithmetic, read end to end, and the ported table. Medium that every order path issues exactly one call in each sub-tick of a turn: the eleven direct call sites are named (`evidence/callers.txt`), but the scheduling of every executor arm across sub-ticks was read per arm, not enumerated.

**Unknown.** Sub-ticks in which `actor+0x3c` is 0 and the executor is skipped; what `actor+0x3c` holds. Observed timing in a run.

### MOVE-106

- Producers of `mover+0x0a`: the mover constructor, actor defaults and the table binding at spawn (`MOVE-RATE-052`); for Humanoid and Human actors, every derive `R0280` copies the low byte of the speed word `actor+0x8c` (`MOVE-RATE-053`), computed from Reaction, the type-word +10 arm, carried load and the speed modifier `+0xd8` (`SAV-1116`, `HERO-SPEED-008`). Effect selector 18 adds into the byte before calling that target's derive (`MOVE-RATE-054`), so a Human ends at the derived value. The Unit derive `R0836` does not write the byte. The turn routines only read it (`L06981`, `L06982`, `L06978`); no per-sub-tick producer exists in them. Serialized movers carry the byte (`SAV-1116`).
- Shipped `world.res:data/data.bin` (`tools/movespeed -mode classes`, `evidence/rates/class-rates.csv`; EN and RU, 0 differing rows over 333): the Humans table has Speed equal to RotationSpeed on 210 of 210 populated rows. Its named classes include `Man_*`, the mounted `ManHorse_*`, `ManMage_*` and the heroes; `PC_Reniesta` is 16, `PC_Reniesta_3` 17. The Units table has 56 populated rows with RotationSpeed 8..23, equal to Speed on 15; `Goblin_Pike` is Speed 24 RotationSpeed 16, `Dragon` 24 and 12, `Troll` 8 and 8.
- The local turn arithmetic of `MOVE-105` is one routine for every actor with a mover, and it reads only this byte for the rate. Scheduling across unit kinds is not established. A mounted rider is a Humans-table class, so its derive is the Human derive. A monster's table rate is independent of its walking speed, and an effect-18 change stays in its byte until another producer writes it.

**Confidence.** High for the derive's store and the shipped table values. Medium that a Human holds its derived value when it turns: table equality does not show when derive ran, and neither the derive trigger set nor post-LOAD recomputation is closed (`SAV-1116`).

**Unknown.** Whether derive runs at Human spawn before the first turn, so whether a fresh Human turns once at its table value. The derived value of a given hero in play, which depends on that hero's Reaction, load and modifiers.

**Amended.** `MOVE-120` narrows the spawn-derive question for a type-6 Human placement only: the spawner calls vtable `+0x50` before placing the actor, so such a Human holds its derived rate from the spawn on, Medium over the unread Human constructor. For the carried party the derive time is Unknown (`SAV-1213`, `MOVE-122`), and the first-turn question is unread.

### MOVE-107

- The stepper's mismatch arm turns and returns without a position change; the step comes at the next call after the facings match (`MOVE-TURN-044`, `MOVE-090`).
- The attack and unit-target cast gate `R0041` requires `mover+0` equal to the 8-way heading of `R0051` and edge distance within reach (`AI-405`); cast arm 8 tests it at `L00742` and otherwise calls the approach `R0042`. The point-target gate `R0086` requires `mover+0` equal to the 8-way direction of `R0089` and the larger cell-axis distance within reach; pending arms 9 (`L06108`) and 0xf (`L13336`) call it. Both gates compare the current facing `mover+0` with the requested heading (`L00577`, `L06109`) and test reach; neither reads `mover+1` or `mover+0xa0`, so a current facing off the heading blocks them, and that test alone does not show that every turn has ended. A self-target cast skips the gate: `L06104` compares actor and target and `L05041` jumps to `L06105`, which sets cast action 0xd at `L13337`.
- A later call with another desired byte replaces `mover+1`. While `mover+0xa0` is 1 that call goes to the leaf from the current byte, with no short-arc snap.
- The stop reset `R0007` (64 direct call sites) writes `mover+1` := `mover+0` when they differ and `actor+0x54` = 0 (`L13338`, `L00017`). It does not write `mover+0xa0`. The turn stops at the current byte, and the next fresh short turn steps by rate: a 2/16 arc takes 2 calls at rates 16..20 instead of 1 (`evidence/turnsim.txt`).
- Pending order 0xa turns until `mover+0` equals `mover+1` and then clears the pending byte (`L06983`..`L13339`); the only byte-immediate store of 0xa to `ord+8` in `.text` is in the walk at `L13340`, after a turn toward an occupied cell.

**Confidence.** High for each branch and store named. Medium for the composed sequences: other writers of `mover+0xa0` in register form were not enumerated, and no run shows whether an act and a turn overlap.

**Unknown.** How often a reset lands during a turn in play.

## Registry at a campaign binder

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-113 | The mission-end server cull empties the actor registry unconditionally and the stepper does not run before the next binder, so that binder sees the new map's type-6 placements and no previous-map actor, bar out-of-tick inserts. | High / Medium | ● active | [EXP-0515](../experiments/EXP-0515-map-load-registry/) |
| MOVE-114 | On the new-mission path the members the join walk places enter the registry after the map's type-6 records; a companion an AddHero command inserted between missions would precede them, which stays open. | High / Medium | ● active | [EXP-0515](../experiments/EXP-0515-map-load-registry/) |

### MOVE-113

- `R0826` (called at `L07384` on the campaign mission-end arm, `PARTY-M100-033`) empties the tick list `[L00240]+4` (`PARTY-ENDCULL-026`). In `w09` it has no early exit: after its player loop, which calls `R0123` at `L13876`, it deletes `[L00004]` and the map grid `[L04624]`, then calls RemoveAll (`R0033` to `R0251`) on `[L00240]+4` (`L13877`..`L13878`) with no test between.
- It also clears the run flag `server+0x2c` (`MISSION-STOP-016`). The stepper `R0147` returns while the flag is 0. The thread loop `R0075`, the only other caller of the sub-tick and of the revive chain, is created by a launcher that tests the same flag and has no found caller or stored address (`AI-385`, `SESS-DEFEAT-065`). `R0512` sets the flag at `L06208`, after the map load and the binder.
- All 40 references to `[L00240]` (linear and raw scans agree) are classified. Inserts: through `R0411` at `L13879` (walk), `L13880` (`R0066`), `L13881` (type-6 spawner), `L13882` (trigger spawner), `L13883` (`R0003`), `L13884` (`R0432`); through `R0032` at `L12539` (SAV restore), `L13885` (walk), `L12785`, `L13886`. Removals: `L13887`, `L12779`, `L13888` through `R1133`, and the RemoveAll. `L00241` and `L13889` pass the registry to the grid constructor `R0118`, which keeps a copy.
- Of the classified inserts, the map load reaches the spawner's; the map-load routines were not read whole for another. `R0066` (AddHero `0x49`, summon) is called only from the command executor `R0061`, which the drain `R0191` calls; the drain runs from the sub-tick (`L09146`) and from `R0192` (`L09418`), whose callers were not enumerated.

**Confidence.** **High** that the cull empties the registry (`PARTY-ENDCULL-026`; the `w09` read agrees) and that the stepper does not run before the next binder. **Medium** for the cull's order of deletions and the absence of a test before the RemoveAll as read here: `w09` lies past the fourteenth range of the preregistered window budget. **Medium** that nothing else inserts between the cull and the binder: the thread loop's launcher is Unknown (`AI-385`); the callers of `R0192`, of the routines holding `L12785` and `L13886`, and of `R0432` were not enumerated; the grid's stored copy of the registry pointer is not traced.

**Unknown.** Whether a command drained through `R0192`, such as AddHero `0x49`, can execute between the cull and the binder. An actor inserted then would also sit in its owner's flat list. Enumerating the callers of `R0192`, or a save written in town after a hire that shows the hired actor in the registry, settles it.

### MOVE-114

- `R0099`, on the new-mission arm, writes `server+0x154` = 0, calls the session start `R0512` at `L06688`, then the opcode-`0x04` sender `R0545` at `L03843` (its only caller), then the stepper at `L03780`. `R0512` has one other caller, `L09338` in `R1662`, which has no caller.
- The placement walk `R0065` has three direct callers. `L06675` in `R0131` is called from the executor's opcode-`0x04` arm (`L06672`); the executor's only caller is the drain `R0191`, whose callers are the sub-tick `R0193` (`L09146`) and `R0192` (`L09418`, callers not enumerated). `L06590` in the revive is reached from `R1576`, which the stepper calls on subclock 15 (`L13890`) and the thread loop at `L13891`. `L13863` in the watchdog `R1290` is taken only when `server+0x148` is nonzero (`L13892`); the constructor writes 0 and a restore copies the saved word (`SESS-DEFEAT-065`, `SAV-WHEADWATCH-524`). A raw dword scan of every section finds none of these routine addresses, so no pointer table reaches them.
- The join command of this mission is sent at `L03843`, after `R0512` has run the spawner and the binder, so the walk it triggers runs after the binder, whichever drain executes it. Every registry insert is an AddTail (`MOVE-TICK-013`), so the members the walk places (the primary and the carried members, whose registry nodes the cull removed, `MOVE-113`) follow the type-6 records.
- As read in `w08`, the walk appends each member, seated or not, with `R0032` at `L13885`, and skips a member the new grid already holds at its position (`L13893`..`L13894`).
- The open alternative: AddHero (`R0066`, insert at `L13880`) is reached only from the executor. If a hire executes through `R0192` or the thread loop in town, between the cull and the next map's spawner, the hired companion is in the registry before the type-6 records.

**Confidence.** **High** that the members the join walk places enter after the type-6 records on the new-mission path: the session start, the sender's position and the walk's callers are read at instruction level, the raw scans exclude table dispatch, and the order does not depend on whether `R0545` queues the command or runs it. **Medium** for the whole party's order, open to the AddHero alternative above. **Medium** for the insert site and the grid-held skip of the walk: `w08` lies past the fourteenth range of the preregistered window budget.

**Unknown.** The callers of `R0192`, and so whether a companion hired in town enters the registry before the next map's type-6 records; enumerating them, or a save written in town after a hire, settles it. The order on the SAV restore path, which inserts through `L12539` (`MOVE-TICK-017`).

## Step heading and turn sources

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-115 | Before a step the stepper turns to `R1356(actor, next cell)`: eight bytes from the signs alone of the cell centre minus the actor's fine point; the zero vector returns `(facing << 5) & 0xff`. | High / Medium | ✔ promoted | [EXP-0522](../experiments/EXP-0522-target-heading/) |
| MOVE-116 | Each of the 11 direct calls of the turn `R0056` takes its desired byte from `R0051`, `R0089`, `R1356`, the stored `mover+1`, a caller's argument or the idle draw. | High / Medium | ✔ promoted | [EXP-0522](../experiments/EXP-0522-target-heading/) |

### MOVE-115

- The stepper `R0054` copies the route head node's cell word to `mover+6` (`L14205`..`L00957`), calls `R1356(actor, mover+6)` at `L07008` and stores the low byte at `mover+1` (`L14206`). When `mover+0` equals it, it takes the rate and steps (`L06858`, `L06896`); otherwise it calls the turn at `L07010` (`MOVE-090`, `MOVE-084`).
- `R1356` (`w05`, 70 instructions, no call): `dx = (cell low byte << 8) - ((pos[0] << 8) + pos[4]) + 0x80` (`L14207`..`L14208`), `dy` likewise from the cell's bits 8..15, `pos[1]` and `pos[5]` (`L14209`..`L14210`). The result:

  | | `dy < 0` | `dy = 0` | `dy > 0` |
  |---|---|---|---|
  | `dx > 0` | 32 | 64 | 96 |
  | `dx = 0` | 0 | see below | 128 |
  | `dx < 0` | 224 | 192 | 160 |

- There is no magnitude test: any vector strictly inside a quadrant gives the diagonal, so a cell two right and one down gives 96 where `AI-444` gives 64.
- Zero vector: the routine loads the current facing `mover+0` into AL at `L14211`; the arm `L14212`..`L14213` keeps it and shifts left by 5, returning `(facing & 7) * 32`, so 0 for every facing that is a multiple of 8 (`s4300-zero-vector.tsv`, all 256 facings).
- The actor's anchor fine point is used, with no footprint term.
- A transit starts at sub-cell 0x80/0x80 (`MOVE-084`). From there an adjacent cell gives `dx`, `dy` in {-256, 0, 256}, and the result equals `AI-444`'s for the same vector.
- Replay: 24,760 inputs equal the table: cell offsets in [-2,2]^2 with own sub bytes {0, 1, 0x7f, 0x80, 0x81, 0xff}^2 and five facings (4,500), the zero vector at every facing (256), 20,000 seeded draws and 4 map-corner extremes (the committed `summary.tsv` set label says 8). EN and RU `rom.exe` have one SHA-256.

**Confidence.** **High** for the law (complete body and instruction replay) and for the stepper's call site and store. **Medium** that every step turns through this site: `R1356` has 11 direct callers (`s01`), and the ten besides `L07008` (`L14214`, `L14215`, `L14216`, `L14217`, `L14218`, `L14219`, `L14220`, `L14221`, `L14222`, `L14223`) were not read.

**Unknown.** Whether a route head node can be the centred actor's own cell, which would reach the zero-vector arm.

### MOVE-116

The 11 direct calls (`s01`; no dword in any section holds `R0056`) and the source of the desired byte at each:

| call | routine | desired byte |
|---|---|---|
| `L00752` | pending order 0xa | `mover+1`, the last stored desired byte; the wait arm sets pending 0xa at `L13340` after its turn |
| `L13266` | `L13333` | its stack argument; no direct caller and no dword reference (`s02`) |
| `L00751` | idle turn `R0205` | `mover+1` when it differs from `mover+0`, else a rand-based byte (`L14224`..`L00997`) |
| `L13267` | face-and-mark arm | `R0051(actor, ord+0x0c)` at `L14175` |
| `L00753` | `R0250` | its argument; callers `L13335` (`R0051` at `L14168`) and `L14174` |
| `L01933` | walk stop arm | `R0089(actor, cell)` at `L14194` |
| `L01934` | walk wait arm | `R0089(actor, next cell)` at `L14202` |
| `L00754` | approach `R0043` | `R0051(actor, target)` at `L06115` |
| `L07010` | stepper | `R1356(actor, mover+6)` at `L07008` |
| `L13268` | wrapper `R2199` | `R0051`; no direct caller and no dword reference |
| `L13269` | wrapper `R2200` | `R0089`; no direct caller and no dword reference |

So a turn before an act or a step aims at one of three heading routines (`AI-444`, `AI-446`, `MOVE-115`), a byte one of them stored earlier, or a byte a caller supplies.

**Confidence.** **High** for the enumeration of direct calls and for the source at each site read (`w07`..`w16`). **Medium** for `L14174`, placed by the adjacent `R0051` call at `L14173` without a read of its body, and for the argument of `L13333`, whose caller is unknown. A computed call of `R0056` is outside the instrument.

## Mover state at mission start

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MOVE-119 | The actor base constructor sets Mover facing `+0x00` and wanted facing `+0x01` to 0x40 plus one rand() draw over 0..0x80; of the image's four Mover constructor calls only this one stores a facing. | High | ✔ promoted | [EXP-0523](../experiments/EXP-0523-start-motion/) |
| MOVE-120 | The type-6 spawner has no Mover access of its own; it runs the class derive at `L04446`, so a placed Human holds its derived rate, then places the actor through `R1145` at radius 0, which draws no random value. | High / Medium | ✔ promoted | [EXP-0523](../experiments/EXP-0523-start-motion/) |
| MOVE-121 | The occupy `R0050` writes Mover `+0x72` (a step-cost word), `+0x82`/`+0x83` (cell) and `+0x84`/`+0x85` (sub-cell) before it enters the footprint's cell-record slots; its own body writes no other Mover byte. | High / Medium | ✔ promoted | [EXP-0523](../experiments/EXP-0523-start-motion/) |
| MOVE-122 | The join walk places each carried member, then deletes its Mover and order block and installs constructor-state ones (facing 0, mask 0x41, rate 0x10); all 22 saved carried members hold the derived rate instead, its writer Unknown. | High / Medium / Unknown | ✔ promoted | [EXP-0523](../experiments/EXP-0523-start-motion/) |

### MOVE-119

- `R0185`, the actor base constructor (`w02`): `new(0xb4)` at `L00963`, `R0206` at `L00964`, the store into `actor+0x154` at `L06970`. It then calls `R0861(0x80)` (`L14121`, `L14122`), adds 0x40 (`L14123`), stores the low byte at Mover `+0x00` (`L14124`) and copies that byte to `+0x01` (`L14125`, `L14126`). It calls the stat defaults `R0160` at `L06971`, which writes Mover `+0x0a` = 8 (`L06973`), and the Units-table spawn `R0184` at `L00493` when `L04697` returns a positive value.
- `R0206` (`w01`, whole): `REP STOSD` of 0x2d zero dwords (`L14127`), then `+0x05` = 0x41, `+0x0a` = 0x10, `+0x09` = 0xff, `+0x08` = 5 (`L06852`..`L00587`).
- `R0861(n)` (`w12`, whole): 0 without a call when n is 0, otherwise `(rand() * (n + 1)) >> 15` with the signed rounding at `L14128`..`L06614` (`MAGIC-284`). With n = 0x80 the draw is 0..0x80, so the facing is 0x40..0xC0, and 0xC0 needs a `rand()` of at least 32514. Each actor construction consumes one value of the shared `rand()` stream here.
- Callers of `R0206` (`s01`, a capstone operand sweep of `.text` plus a raw dword scan of every section): four direct calls and no stored address. `L00964` is this constructor. `L14119` is the join walk (`MOVE-122`). `L14120` is the carry importer (`PARTY-LOSS-006`). `L13291` is in `R0825` (`w09`), which deletes the actor's Mover and installs a constructor-state one (`L14129`..`L13292`) and a new order block (`L14130`..`L14131`); `PARTY-037` names it as the revive repair arm's call. The other three install the constructor's facing 0 and rate 0x10 and write no facing.
- `R0471`, the mask install (`MOVE-DOM-025`), has two direct calls, `L06850` in `R0184` and `L06851` in `R0876`, and no stored address (`s01`).
- This enumeration answers the open item of `AI-413`, which named only the call at `L13291`.
- EN and RU `rom.exe` are one image (SHA-256 `942e9b72…7d03`, `image-hashes.txt`).

**Confidence.** High. The stores are named instructions in whole listings. The caller set rests on two instruments: the operand sweep misses a call from bytes the linear decode does not reach, and the raw dword scan covers every table, callback array and stored address; an indirect call needs the address as data, which the scan finds nowhere.

### MOVE-120

- `R0151`, the type-6 spawner (`w04`, read whole to its return at `L14132`), contains no access to `actor+0x154`. Per record it resolves the owner and class, constructs the actor (Human constructor `R0497` at `L02316`, `L02314`, `L02315`; Unit `new(0x198)` and `R0501` at `L02317`), applies the Units-arm difficulty changes, destroys the actor when `actor+0x0e` is 0 (`L12640`..`L14133`), writes the authored overrides, calls vtable `+0x50` (`L04446`), then calls `R1145(rec+0x00 >> 8, rec+0x04 >> 8, 0)` (`L12752`..`L11042`). A zero return logs and destroys the actor through vtable `+0x04` (`L13866`..`L14134`); otherwise the actor enters the registry (`L01792`), its owner and its group.
- Vtable `+0x50` is `R0836` for Unit and `R0280` for Humanoid and Human (`t01`). `R0280` stores the low byte of `actor+0x8c` into Mover `+0x0a` (`MOVE-RATE-053`); `R0836` has no such store. A placed Human therefore holds its derived rate from the spawn on, and a placed Unit holds the Units-table byte or the default 8 (`MOVE-RATE-052`). This settles `MOVE-106`'s open question for type-6 Human placements: derive runs at spawn, before the first turn.
- `R1145(x, y, r)` (`w05`, read to `L14135`): the position store `R0289(x, y, grid)` (`L14136`), then attempts at `x - r/2 + R0861(r)`, `y - r/2 + R0861(r)` through `R0290` and the test `R1287` (`L14137`..`L06650`), `(r*r)/2 + 2` attempts. With r = 0 `R0861(0)` returns without a draw, so both attempts test the placement cell, and a refusal returns 0 (`L14138`..`L14139`, `L14140`). An admitted cell calls the occupy `R0050` (`L06653`, `MOVE-121`). With r above 0 a refusal goes on to a scan of the square (`L14141`..`L14142`).
- `R1287` (`w06`, whole) only reads: for each of the n by n footprint cells it tests the bytes at `grid + 0x20000` and `grid + 0x10000` against Mover byte `+0x05` and returns 0 on any common bit.

**Confidence.** High for the spawner's lack of Mover access, its call order and the radius-0 path of `R1145`, each read in whole listings. Medium that `R1145` returns the occupy's result: the epilogue after `L14135` was not read. Medium for the Units-table rate source (table byte or default 8): `MOVE-RATE-052` leaves final spawned rates Unknown, and the census (`SAV-1213`) holds 312 Unit rates, none equal to 8, and compares none to the table.

**Unknown.** The Human constructor `R0497`, the Unit constructor `R0501`, the Humanoid defaults `R0876` (caller of the second `R0471` call) and the vtable `+0x50` bodies other than the rate store were not read. The first `R0471` call at `L06850` follows a test of `L04697`'s result (`w02`, `L14143`..`L00493`), also unread. `MOVE-119` bounds who builds the Mover (no other Mover constructor call), not who stores into it afterwards.

### MOVE-121

- `R0050` (`w07`, whole) reads the footprint side through vtable `+0x1c` and the domain through `R1851`, takes the cell bytes from the Position object (`R0299`, `R0300`), and stores, in order: Mover word `+0x72` = `R1345(actor, x, y)` (`L07074`, `L07161`); `+0x82` and `+0x83` = cell x and y (`L02490`, `L02491`); `+0x84` and `+0x85` = the low bytes of `R0165` and `R0166`, the sub-cell (`L02488`, `L09282`). When `R1287` refuses the cell and Mover word `+0x80` differs from the packed cell it calls the two one-return stubs `R1361` and `R0459` (`MOVE-088`). It then calls `R0458` for each footprint cell in row-major order and stops at the first refusal.
- `R0458` (`w07`, whole) contains no Mover access. It writes the actor into cell-record slot `+0x04` (domains 1 and 2, `L07121`, `L14144`) or `+0x08` (domain 3, `L07122`, `L14145`), creating a missing record through `R1351`, and runs the recompute `R0453` (`TERR-CELLREC-146`, `MOVE-088`).
- `R1345` (`w08`, whole): for domain 1, `(signed word actor+0x8c * 8)` divided by the byte at `grid[y * 256 + x]` (`L14146`..`L07076`); for domains 2 and 3, the word `actor+0x8c`; otherwise 0.
- A placement through `R1145` therefore leaves the Mover with these five fields set from the placement cell, the speed word and the cost byte, and the cell-record slots of its footprint holding the actor.

**Confidence.** High for the stores and their sources, read whole. Medium that the occupy writes no other Mover byte: its callees `R0453`, `R1342`, `R1362`, `R1351` and the trigger builders `R1137`/`R1138` were not read here. The census in `SAV-1213` finds no other nonzero Mover byte in 421 placed records.

### MOVE-122

- `R0065`, the join walk (`w10`; the member loop `L14147`..`L14148` read whole): for each actor of `player+0x20` it skips the actor when the grid's slot accessor `R1826` already returns it at its cell (`L14149`..`L14150`); otherwise it stores the start cell through `R0289` (`L14151`), places `player+0x34` first with `R1145(x, y, 0)` (`L06656`) and any member still unplaced with `R1145(x, y, r)` (`L06655`), r computed from the member count (`L14152`..`L06586`); a failure only logs (`L14153`..`L14154`).
- After the registry AddTail (`L14155`) it deletes the Mover (`R0525` at `L14156`), constructs a new one (`new(0xb4)`, `R0206` at `L14119`) and stores it in `actor+0x154` (`L14157`); it releases the order block (`R1379(1)` at `L14158`) and installs a new one (`R0286` at `L14159`, `actor+0x158` at `L14160`); it calls `R0063` on `actor+0x70` (`L14161`).
- The occupy's Mover stores (`MOVE-121`) go into the deleted Mover. The new Mover holds the constructor's values only: `+0x00`/`+0x01` 0, `+0x05` 0x41 whatever the member's domain, `+0x08` 5, `+0x09` 0xff, `+0x0a` 0x10, every other byte 0. The walk contains no facing store and no `R0471` call.
- The walk body has no access to `actor+0x15c` or `+0x178`, the two embedded route lists (`SAV-630`).
- The start cell is `player+0x60` when `server+0x0c` and that word are nonzero, else a random entry of the list at `this+0x6c`, else `0x1e + R0861(0x46)` per axis with a log line (`L14162`..`L14163`). A placement with r above 0 draws two `rand()` values per attempt.
- The walk runs from the join command the mission start sends at `L03843` (`MOVE-114`); in the census the carried members' Movers already hold the walk's values at the restart save (`SAV-1213`).

**Confidence.** High for the walk's own stores and calls, read whole, and for the walk's installation of rate 0x10. Medium that the route lists survive the walk (the walk body has no access to them): the callees `R0289`, `R0063` and `R1379` were not read; one census record agrees (`SAV-1214`). A member the grid already holds at its cell keeps its old Mover (`L14150`); whether that occurs at a mission start is Unknown. The rate clause is Medium as a statement about the saved carried member: the census contradicts rate 0x10 on 16 of 22 carried records, which hold the derived value (15..20, `SAV-1213`), so the walk's rate is not the value a carried member holds at the save.

**Unknown.** The first writer of Mover `+0x0a` after `L14157`. The census shows the derived value before the restart save (`SAV-1213`); the call that runs Human derive on a carried member between the walk and that save was not identified. Read without finding it: the join entry `L14164`..`L14165` (it calls the walk at `L06675` and goes on to `L14166`, `R1626`, `R1625`, `R0081`, `R0127`, `R0299`, `R0300`, `R0217`, `R0059`) and `R1626`..`L14167`; both hold no derive call and no Mover store, and their callees are unread. The settling read is outside that range (the actor tick `R0037`, the order executor `R0016`, or a watchpoint on Mover `+0x0a`).
