# Target acquisition and sight

[Reference](format.md)

## Group members and nonadjacent damage

Engagement can use weapon reach, a class spell or a weapon spell when those
inputs exist. Point and area payloads and attached poison can damage after the
source separates; their admission, delivery and damage predicates are separate.
The identified paths are a bounded inventory: computed writes, other-width
stores and indirect callbacks remain unclassified. — AI-745

Both shipped mission-20 maps place ten `A_Brigand1` and one
`F_BrigandLeader1` in owner-3 group 5. Their initial weapons give reach 4
(Short Bow) and 1 (Two Handed Axe). Both definitions have no mana, known spells
or weapon `castSpell` property. Bow damage is therefore available without
adjacency; no initial spell path is supplied by those definitions. Later
equipment, script, save and external-effect changes are outside this statement.
— AI-744

## Target-selection sequence

1. **Population — sight.** Clear `world+0x82ef0` (`0x10000` bytes, one per cell). Stamp the
   deciding actor's line of sight into it (below). If the actor is AI-owned and remembers being
   attacked, mark that cell too — for 20 ticks. Then walk the global on-map actor list and keep
   every actor standing on a marked cell.
2. **Relation.** Drop every candidate `c` for which `matrix[me.Player+0x04][c.Player+0x04] & 1`
   is 0. Drop an invisible candidate (`c+0x144 & 0x8000`) unless some member of the decider's
   group is within that member's own `order+0x71` of it.
3. **The dead.** Move survivors with `health < 1` to a second list. If nothing living is left
   and that list is not empty, move it back: a unit with only corpses in view targets a corpse.
4. **Selection.** Minimise the footprint-aware edge distance `d` (1 = touching), with `d + 1`
   for a flying candidate when the decider does not fly; break ties on the 16-way turn cost from
   the mover's current facing. Seed the best distance at `reach + 1`, and **discard the winner
   if its distance exceeds `actor+0x12c`**. On a full tie — identical distance and identical turn
   cost — the **later**-scanned candidate replaces the earlier one, the opposite convention from
   the group-level scorer's own first-wins tie rule (`AI-340`).
5. **Outcome.** A winner sets `order+8 = 6`, `order+0xc = target`, `order+0x14 = reach`. No
   winner: a human participant's unit tries to heal; an AI-owned unit goes to the guard state,
   which walks it back to its post whenever it is farther from the post than `actor+0xa5`.
   The values left at `order+8` are 6 with a winner; with none, `0xb` for an AI-owned unit (a
   mage then 8 when its cast arm finds a target) and 0 or 8 from the heal arm for a human
   participant's unit (`AI-715`). The mage's cast arm targets itself when damaged, else the first
   kept unit below maximum health within Heal's range, else one drawn unit of its own Player for a
   buff (`AI-785`, `AI-786`); the heal arm targets the first listed damaged unit (`AI-789`). Callers,
   the null-target search and the per-spell routes: `AI-784`, `AI-787`, `AI-788`. No count of units on a spell's cells enters the choice (`AI-796`, Medium; each of the five `R0020` calls in the six chain windows centres on the actor's own cell, `AI-862`; the engage routine `R0009` and the heal routine `R0113` that use those gathers store `ord+8` = 5 with the victim and `ord+0x60` = 1 and, in their own instructions, no route field; Medium for the composition, `AI-886`); the session-list cast routine that picks a unit on a cell is `AI-797` (the list entries are not decoded; that it is not a computer unit's cast is Medium).

**Sight decides the population and reach decides the pick.** An engine that acquires on sight
alone starts fights the game does not; one that scans only within reach acquires through walls,
and acquires things it cannot see. Neither radius is a substitute for the other.

### Prismatic Spray is a distinct consumer of the group population (`AI-SPRAY-266`, `AI-SPRAY-267`)

`R0109` is the fifth direct caller of the group builder. It consumes shared list A head to
tail and then list B head to tail; it does not run a radius search. The ordinary secondary gates are
therefore the builder's group-wide visibility, first-member diplomacy and See-invisible rule, then
its living/corpse split. The primary spell target is separate: the selector alarms it and performs
the hostility flip before the builder call, appends it first, and skips an equal secondary later.
It can therefore survive conditions that would remove it as a secondary. The selector has one rel32
caller, Prismatic Spray's `R0269` (census of rel32 transfers in `.text` and aligned dwords; immediates and computed calls not searched; `MAGIC-474`).

Prismatic rank differs from ordinary acquisition. List A uses
`((edgeDistance<<8)+turnCost)&0xffff`; list B uses
`((((edgeDistance<<8)+turnCost)&0xffff)<<8)`. Repeated strict-minimum scans preserve source order on
equal stored scores. Selection first requires a score below 65530 and append later requires the
saved threshold below 65000. A B score below that bound needs an edge distance that is a multiple of 256; with
footprint sizes of at least 1 the distance is at most 255, so no B corpse is selected
beside living A (`AI-863`, Medium: the footprint getter is unread). If there are no living entries, the
builder has already moved B back into A and the only-corpses population uses the A formula. The
spell's capped byte limits the final output list after ranking; it is neither actor reach nor order
`+0x14`, and the primary consumes one position on the ranked path. When the builder returns empty A, the selector instead
appends the primary and returns before reading the cap. Its stack pointer region holds 100 entries
and its winner/threshold regions ten; neither population overflow nor a custom cap above ten is
checked. Candidate scores are scratch at `+0xd74` of the object at `[L00004]`, allocated with 0xc320 bytes, so at most 11627 dwords; the capacity itself is Unknown (`AI-871`).

### The sight stamp, complete (`AI-GROUPSEE-068`, `AI-LOS-087`…`AI-LOS-091`, `AI-SIGHT-092`…`AI-SIGHT-094`; the writer count of `AI-SIGHT-092` is partially retracted)

The sight object is embedded at `world+0x58ee8`. Its init runs in each of the four world
constructors, **before the map is loaded**, and builds three tables that depend on
`k = [Scanning] ScanShift` of `World\Data\map.reg` (ships **7**) and on nothing else — no terrain,
no actor, no tick. Two of them serve the stamp: `step`, `+0x22000`, and `cost`, `+0x28000`, both
`u16`/`i16` over a 41×41 window inside a 64×64 grid, addressed as `base + (a<<7) + 2b` with `a`
the column offset and `b` the row offset, both `0..40` around a centre of `(20,20)`.

```
step[a][b]  = the cell one Bresenham step toward the centre, as (i8 dx, i8 dy) in {-1,0,+1}
              zone by slope:  j <  i>>1 -> column step | j > 2i -> row step | else diagonal
              the builder's four quadrant mirrors collide on the j==0 axis, so it repairs
              cells (+1,0) and (-1,0) with four literal stores at the end
cost[a][b]  = ftol( (1<<k) * sqrt(i*i + j*j) / max(i,j) )   -- the mean length of one step
              along that ray, in 1/(1<<k) cell.  k=7: 128 on the axes, 181 on the diagonals

stamp(actor at cell C, scanRange = actor+0xa5):
  acc[*][*] = 0 ; acc[20][20] = (1 << (k-1)) + (scanRange << k)      -- a budget in 1/128 cell
  vis[C] = 1 ; alt = (i8) height[C]                                  -- height = world+0x9451c
  for r = 1 .. 19:                            -- Chebyshev rings; ring 20 is unreachable
     allBlocked = true
     for each of the ring's four edges, 2r+1 cells each, in edge order:
        skip the cell unless 8 <= absCol <= W-9 and 8 <= absRow <= H-9   (world+0x58ee0..3)
        p = step[a][b]
        v = acc[a+p.dx][b+p.dy] - cost[a][b] - (i8) height[cell] + alt
        acc[a][b] = v                         -- stored EVEN when it blocks
        if v <= 0: continue                   -- blocked: NOT pruned, the ray is not cut
        vis[cell] = 1 ; allBlocked = false
     if allBlocked: stop
```

The stamp uses a height-sensitive budget:
`acc(n) = seed − Σcost − Σh(cell) + n·h(observer)`.
The observer's height is added at every step; each traversed cell's height is
charged once. On flat ground the region is a disc of radius `scanRange`
(145 cells at 6). A single high cell is charged once and does not cast a
continuing shadow. A negative-budget cell is still written and the march
continues. The minimum step cost is `1<<k = 128`; installed altitude bytes
`0..127` cannot restore a negative budget, while an authored byte `>=0x80`
can. All three stamp sites center the window on the actor's own cell.

Drawn fog uses the same predecessor/cost tables and an algebraically equal
seed for whole-cell sight. The shared algorithm has two parameters:
— AI-SIGHT-093

- **the seed's width.** The fog reads the whole `u16` at `actor+0xa4` — sight in 1/256 cell — and
  shifts it right by `8 − k`; the sight stamp reads only its **high byte** `actor+0xa5`, the
  whole-cell radius. A monster's sight is whole cells and the two coincide; a hero's is
  `ftol(((mind + reaction)/25 + 4) × 256)` and is not, so **the AI grants a hero less sight than
  his own fog shows him** — 105 cells against 145 at `actor+0xa4 = 1535` (`AI-SIGHT-094`).
- **the playable rectangle.** `8 .. W−9` here against the view's `7 .. W−8`, and the byte-wide compares
  here against 32-bit ones there (`TERR-FOG-118`).

A guard group's base radius is the largest member distance to the centroid plus that
member's `actor+0xa5`, then the guard setter's floor or script argument (`AI-793`); the one
shipped Guard node passes 13 (`AI-902`).

`actor+0xa4` has at least twelve writers: the six `AI-SIGHT-092` names, whose count of six is
partially retracted, the four Light and Darkness apply and remove stores (`MAGIC-443`),
`L17538` in `R0840` (`HERO-MOD-016`) and `L17539` in `R0843`. A
`disp:a5` sweep sees two of the six; the seven `+0xa5` byte reads it lists include the stamp
seed (`MAGIC-444`). A Humanoid derive keeps Darkness's reduction through `+0xe4`
(`MAGIC-445`). The Darkness tail that stores the caster has no comparison of caster and unit (`AI-829`). The three area loops pass their own area object, which carries the caster in `+0x3c` when the caster's own `+0x3c` is nonzero at the cast (`AI-867`; Medium for later stages).

**Light and Darkness shrink the stamp.** Light lowers a reached plain Unit's sight, and Darkness a
reached Humanoid's or Human's, by p/30 + 1 whole cells, so the stamp's seed and the candidate set
acquisition draws from shrink with it. The reach cap is unchanged: a reached unit whose lowered
sight is below its reach acquires only inside that sight (`MAGIC-446`).

**A unit only closes distance when it is given a pursuit order** — `order+8 = 5`,
`order+0xc = target`, `order+0x14 = reach` as the stop distance. Acquisition never produces one;
the guard and engage arms do.

**Two clauses suppress engagement entirely**, both requiring all three of `actor+0x4c & 4`,
`actor+0x12c < 2` and a human-participant owner (`Player+0x28 == 0`). Mask bit 4 is set at
construction on an actor with mana, and reach is the worn weapon's `@.range`, 1 for a melee weapon
or none (`AI-714`), so the clauses cover the human participant's mages and mana units holding a
melee weapon or nothing; no install-save record is such a unit (`AI-717`). The state-`0x17` arm
sets reach 5 for class word `0x18`, a Humans class carried by 57 mana rows; no writer of state
`0x17` was found (`AI-STATE-043`, Medium; `AI-716`).

## After the target is set

Once `order+0xc` holds a target, engage, the cast setter, the strike gate, the executor and the
cast admission call no sight-map routine directly or through two further levels of direct calls,
outside acquisition and reacquisition: engage, the cast setter and the executor reach those two,
and they were not entered. The invisibility filter `R0021`, called by engage and the cast
setter, applies a group detection range to invisible victims. A weapon carrier stops at its reach
and a book caster at the spell's stop distance with no sight-map test on these routines. The
cast setter given no target takes the first hostile of its mover-radius square, filtered for
relation and invisibility, with no sight map (`AI-718`, Medium). The executor body holds two timer-import calls and two table jumps, and the admission body and three direct callees of the admission hold one `vt+0x1c` size-getter call; engage, the cast setter and the gate hold none (`AI-868`, Medium: the getter's targets, the jump targets, the executor's direct callees and the callees of `R0269` are unread).

No target path filters at −10 health. A body below 1 health stays a fallback target until its
dying arm, the countdown spent, finds health at or below −10, tears it down and the world step
unlinks it in the same pass (`AI-719`).

## What a refused route does to a pursuit order

The route search sets one flag on failure, `mover+0x98`, with exactly three setters and one
consumer — the per-actor order machine's own epilogue, which clears it unconditionally the tick it
reads it (`AI-ROUTE-045`, `MOVE-GATE-039`). It is a plain one-shot boolean, not a counter: nothing
accumulates a refusal count and nothing times one out. For any actor state outside the three the
epilogue special-cases (teardown, patrol advance, one hard override), the response is the same on
every refusal: `order+8 = 0` and a fresh reacquisition scan, `R0004`.

That scan's candidate filter is a raw Chebyshev cell distance (`AI-327`) — the same kind of metric
the group scorers reach through their own routine (`AI-COST-071`), not the footprint-subtracted
edge distance this page's selection sequence above uses. So an actor's own footprint size, and the
footprint of whoever is blocking its route, play no part in *which* new victim the reacquisition
scan can pick. Four things narrow the list: the reach threshold `actor+0x12c`, a running-nearest
gate that drops any candidate farther than the nearest one seen so far, the living/dead partition
that holds corpses back unless nothing living survives, and the diplomacy filter. The previous
victim is not excluded by identity — only the pursuer's own self is. Footprint size re-enters
exactly once, after a winner is already chosen: the same in-position test ordinary pursuit uses
(above) decides whether the newly-acquired target is struck immediately or walked toward first.

The running gate precedes both later filters. A closer nonhostile or dead
candidate can lower it and prevent later living hostiles from entering the
list. Previously admitted farther candidates remain. Permuting identical
actors can therefore change the survivors, select a corpse through fallback,
or leave no winner (`AI-360`). The final score is only circular byte turn cost:
`min(abs(facing-direction), 256-abs(facing-direction))`. The selected direction
helper returns eight headings, multiples of 32. Equal scores replace the
winner, so the last surviving tied candidate wins; a farther admitted actor
can beat a nearer one on turn cost (`AI-361`).
The same claim retains all 81 replayed heading outputs for equal size-1
footprints at centred cells with target dx/dy from -4 through 4. Its
coincident-centre result is 224. The helper's law for every fraction, footprint
and 16-bit wrap, with no floating-point step, is `AI-444`. The same centre sum, less one fine unit,
places a spell cast at a target wider than one cell (`MAGIC-449`); `MAGIC-450` lists the spells that
read that point and `MAGIC-451` the units wider than one cell.

This scan follows the bound manager from head to tail. Selected creation and
return paths append, while unlink preserves survivor order (`AI-362`). The
selected load rebuild instead appends in player, group and member order,
skipping off-map members. It need not preserve the pre-load global sequence
(`AI-363`). Full-game creation chronology and arbitrary native post-load
reordering remain Unknown; local pointer relations do not establish them.

On total refusal — no candidate within reach survives the filters — the order kind still changes
(`order+8 = 0` for a human-owned actor, `0xb` for an AI-owned one), but the stale target pointer at
`order+0xc` is left exactly as it was: the scan's only write to that field sits inside the
candidate-found branch, and the no-candidate exit never touches it (`AI-328`). The same happens on
the found path when the engagement-suppression clause above fires: the target is written, then the
order kind is put back to `0` (`AI-GUARD-007`). The engage routine's refusal for that clause is a tail call of the acquire routine `R0022`, which acquires; it is not a no-op (`AI-GUARD-007`, `AI-886`). Inside the scan itself, the order kind and the flag
that triggered the re-scan are the only fields reset; the AI-owner exit then calls `R0205`, an idle turn that stores `actor+0x54`, `ord+0x54` and
the second facing byte and calls the turn routine `R0056` (`AI-864`). `R0056` and its one callee store the wanted facing, a turn counter, a turning flag, the current facing and `+0x9d` (meaning unread), and in the window `R0056..L18254` no order, route or position field (`AI-884`, High for the window, Medium for the call; `AI-395` records two stores past it); the divisor `[this]` of the idle turn's random step stays Medium.

## What the group-level scorers call

Group orders 1, 2, 3 and 5 re-score every member's target on every AI tick rather than holding one
sticky (`AI-REISSUE-077`). Both scoring routines behind that re-score (`AI-COST-071`,
`AI-REACH-072`) use a plain Chebyshev cell distance, and their complete outgoing call inventory is
four direct targets and six virtual dispatches (`AI-329`). Neither body directly calls the route
search, the occupancy-block builder, or the dynamic-plane claim setters that record a moving unit's
reserved cell. That negative is an exhaustive read of both bodies, not an absence of search, but it
covers the direct calls only: the six dispatches go through `vt+0x1c` and `vt+0x20`, the two slots
`TERR-MOVE-055` reads out of the vtables as field getters, and no dispatch implementation was read
for `AI-329`. On the evidence available, the group layer's periodic re-evaluation has no direct way
to see whether a route to its chosen candidate exists or whether that candidate's cell is claimed
by another mover; that information exists in the per-actor layer's `mover+0x98` flag above, which
the group arms never consult.

## Heal-list membership and a missed blow

The block builder the computer's Heal routine uses appends occupants centre cell first with no identity test (High), so its candidate list can hold the caster (Medium); at full health the caster is never the pick. — AI-800

A blow whose roll fails still records the attacker in the victim's `+0x40` when the attacker is set (High for the resolver body). A census of 246 `disp:0x40` stores finds no other store in the direct-call closure of the resolver to depth 8 or in the three callers above it; indirect calls and copies are unread, so the absence is Medium (AI-869). — AI-801

A staged area spell's first stage runs one world step after the applying step, then every 3 steps (Medium). The cast arm registers the area object through a list append with no drive call, and the DirectDamage `vt+0x3c` body subtracts health in the stage call (AI-870). — MAGIC-463

A refused flee walk under explicit Retreat hands the unit to the reacquisition above, which strikes a hostile at reach when the in-position gate passes (Medium: composed from listings; the gate tests facing and reach, AI-865; the stand-down of Move state 1 belongs to the Wimpy case, AI-929). — AI-798

The in-position gate passes when the direction helper's result equals the actor's facing byte and the edge distance is at most the reach byte. — AI-865

The block builder's cell lookup returns slot `+4` of the cell record, which names the actor of movement domain 1 or 2, so a domain-3 caster is not in the Heal list; a slot `+4` caster is, when registered in its cell's record (Medium). Whether domain 3 is the flier class is Medium (MOVE-DOM-024). — AI-866

## Defender ranking, distance, group ties and retreat detection (first game)

Cover engages the nearest flier when one remains, else the nearest unit, and removes fliers when the protected unit's reach is below 2 (AI-812). The block builder and the engage routine load none of the six health and state offsets searched, so a unit that is still in a cell record is a candidate whatever its health (AI-872; Medium: the other engage callees and the cell record of a dying unit are unread). The block builder lists slot `+4` only, so the flier branch is reached only by a slot `+4` unit whose `vt+0x20` is 3, which is unread (AI-866). Distance there and in the invisibility test is edge to edge in sub-cell steps (AI-813). A group scorer keeps the first of equal costs (AI-814). An invisible enemy stays a candidate of the retreat check when any groupmate's See-invisible byte reaches that distance (AI-815). The filter `R0021` stores list-node words through its list argument (read as candidate links, Medium), has no order-word store in its own instructions, and calls `R0036` and `R1355` (AI-849).

The group candidate list that the assignment loop scores is built by `R0110` from the shared sight map; entries with word `+0x94` below 1 come back only when the list would be empty (AI-913).

The group arms assign each member the lowest-scoring candidate of that list by the loop `R0177` with scorer `R0225`, or `R0252` under order 3 (AI-909); the Guard arm runs it after the clip loop (AI-905). The dispatcher maps each group order to its arm routine (AI-908); the Guard arm reads `actor+0xa5` of every member in its geometry routine (AI-906) and the other store of that field is at `L00261` (AI-903). The arms 2, 4, 5 and 0x11 are read (AI-910); the follow, target-pick and heal helpers store `ord+0x08` = 4, 6, 0xb, 0 or, on the heal path, 8 (AI-911, AI-349); the one immediate store of 0x17 to a displacement-0x50 operand in `L06931..L07036` is in `L18326`, which has no direct caller (AI-912). The engagement setter `R0009` has 14 direct call sites (AI-907). The AI driver has two direct call sites, and the Guard counter `grpAI+0x34` counts Guard-arm passes since the last latch flip, not session ticks (AI-904).

Spell-choice and cast-entry reads: AI-922, AI-923, AI-924, AI-925, AI-926, AI-927.

Withdrawal and retreat reads: the callee `R1362` (body unread) has no direct caller outside the occupy `R0458` (AI-942, Medium; a computed call is not excluded); within the routines read (AI-947, AI-950, Medium) 38 decoded routines that load the scratch record store neither occupant slot through the scratch base (AI-947) and the release `R0057` clears one (AI-950); the capped tails store nothing there (AI-949); the retreat tail's direct callees load no spell id in their own bodies (AI-948, Medium; the closure below the callees and the ordinary cast pick stay as AI-816 states, and the direct callee `R0009` calls the cast table); the occupy calls the recompute that sets plane bit 5 (AI-943, AI-944); a corpse that is not torn down and is alone in the retreat square gives a zero positive-HP count, and whether this occurs in play is Unknown (AI-945), and the list count byte does not wrap for a radius byte of 7 or less (AI-946).
