# Coordination, ticks and cell transitions

[Reference](format.md)

## Seed endpoint and immediate caller

The selected centred zero-radius R0178 static-search tail uses route
count, not the picker's nonzero answer. An admitted substitute equal to the
seed and no substitute both yield zero count from an empty input list. The
caller writes current cell to mover+76, sets+98=1, clears the dynamic list,
then sets+90=1 and returns on current cell==+76. With nonzero static count,
it takes node+8 from actor+164, the resolved end; that branch preserves+98.
The local relation is High; topology replay under declared private services
is Medium. Native actor/order continuation and notification effects remain
Unknown. — MOVE-082

A direct centred zero-radius request equal to current cell exits earlier,
before search and these stores. It preserves prior routes, target/resolved
fields and+90/+98. This local shortcut does not establish native invocation,
later order cadence or final actor behaviour. — MOVE-083

## Coordination between units

The **dynamic block plane is the only channel**, and it carries two things: where units *are* and
where they are *going*.

- **Occupancy.** `R0058` ORs the mover's own bit (`0x40` ground for movementType < 3, `0x80`
  air for 3) over its n×n footprint; `R1332` clears it. The cell is recorded in `mover+0xa6`.
  A dynamic search brackets itself with clear/restore at the unit's own cell so it does not block
  itself.
- **Reservation.** Before stepping, `R0054` reads the next route cell into `mover+0x06` and,
  if it differs from `mover+0x80`, releases the old claim (`R0046`) and marks the new one
  (`R0257`) — on a cell the unit has **not yet entered**. `mover+0x80` is the unit's intended
  cell. The pair is asymmetric: the claim ORs unconditionally, the release skips any cell inside the
  unit's current footprint, so releasing never un-occupies the ground it stands on.
- Every mask includes its own domain's occupancy bit — `0x41` = terrain bit 0 + ground bit 6,
  `0x44` = object bit 2 + bit 6, `0x82` = border bit 1 + air bit 7 — which is why the same predicate
  serves both planes, and why the static search is unit-blind: bits 6/7 are never set on
  `world+0x10000`.
- **Blocked next cell:** a centred ordinary walk chooses a turn-and-return branch or a dynamic
  search through `R0402`; the verdict depends on the requesting actor and the blocker.
  Pursuit has no such verdict branch. An admitted transit continues without a destination-plane
  test, and the occupy can refuse after Position changes. The local branches are High; the
  no-occupant fallback population and later retry cadence remain Medium or Unknown.
  — MOVE-130, MOVE-131, MOVE-132, MOVE-127
- **No priority.** The tick loop (`R0427` → the actor's `vt+0x18`) applies no sort and no
  priority key, and claims land in the shared plane immediately — so whichever unit the loop reaches
  first that tick takes the cell.
- **Formation, conditionally** (`MOVE-FORM-036`; `MOVE-ORDER-023`'s "no formation" is retracted).
  A group Move / Swarm-2 order runs one of two arms. **In formation**, each member is ordered to
  `target + (memberCell − groupCentroidCell)` — the offsets are kept at `ord+0x24`/`ord+0x26` — so
  the destinations differ by construction. **Out of formation**, and for `R0025`'s group
  order 2, every member gets the same loop-invariant cell and spreads only because each one's own
  search fails at the crowded cell and substitutes independently against a plane that already
  carries the earlier movers' claims. Which arm runs is the same gate that decides the group rate
  term — see [group speed gates](rate.md).


## Tick ordering (`MOVE-TICK-013…017`, `MOVE-ID-016`)

The loop walks a pooled doubly-linked list embedded at `+4` of the manager at `[L00240]`
(12-byte nodes: next/prev/element; CPlex blocks of 10), head→tail. **The order is insertion
history and nothing else** — the family's only insert is AddTail; no sort, no AddHead, no
mid-list insert exists.

```
fresh map:    party first, then the map's type-6 records in record order
              (Humans, Units, Sacks tick here; no Building insert was found)
during play:  death       -> unlink (no hole), actor moves to the dead list *(world+0xc)
              spawn/summon-> AddTail (one sack creator is the actor tick itself)
              garrison    -> unlink;  return to map -> AddTail (loses its old position)
              owner change-> tick position unchanged (only player/group lists move)
save/load:    the tick list is NOT serialized. The stream carries player -> group -> actors
              (each list head->tail); the loader rebuilds the tick list as
                for each player (manager order):
                  for its list (groups in creation order, actors in group order):
                    AddTail, skipping off-map actors (actor+0x4c bit 3)
              => within-group relative order survives; the cross-player interleave does NOT.
                 A save/load cycle can change which of two contending units moves first.
```

The **runtime id** (`actor+0x04`, the SAV head id) is *not* the order: it is the lowest free bit
of the bitmap `L06793`, assigned at insert, freed (and the field zeroed) when a corpse reaches
decay stage 5, reused by the next spawn, and restored exactly across save/load (read back and
re-marked by the head serializer `R0950`). A consumer reproducing contention must keep the
**list**, not sort by id.


## Motion fields and dispatch

Position is `*(actor+0x10)`, Mover is `*(actor+0x154)` and the order block is
`*(actor+0x158)`. Centred means both Position fractions are `0x80`; it does
not mean the turn flag is clear. Motion state and order progress are separate.
— MOVE-125, MOVE-126, AI-417

| Owner and field | Width | Meaning |
|---|---|---|
| Position `+00`, `+01`, `+02` | byte, byte, word | Cell X, cell Y, packed `(Y<<8)\|X` |
| Position `+04`, `+05` | byte each | X and Y fractions |
| Mover `+00`, `+01`, `+0a` | byte each | Current facing, desired facing, rotation rate |
| Mover `+a0`, `+9d`, `+a4` | dword, byte, byte | Active turn, admitted turn-call count, recomputed pre-step estimate |
| Mover `+06`, `+70` | word each | Next route cell, source cell of the selected transit |
| Mover `+80`, `+a6` | word each | Target claim, footprint most recently marked by a claim helper |
| Mover `+a8`, `+aa`, `+ac`, `+ae` | word each | Translation rate, transit duration, elapsed calls, direction |
| Mover `+b0`, `+b1` | signed byte each | X and Y strides in 1/256-cell units |
| Order `+08`, `+09` | byte each | Pending order, action progress; progress 3 continues a transit |

The selected stepper `R0054` writes next cell, source cell and the
target reservation before comparing facings. `R0056` either snaps an
inactive arc of at most 32 or applies one rotation-rate step along the shorter
arc, addition on the 128 tie. It recomputes `+a4`, increments `+9d` and clears
`+a0` on equality. For an unchanged heading and nonzero rate, completion takes
one admitted call for the inactive short snap, otherwise `ceil(arc/rate)`.
The exact local law is High. Actor/order gates, new headings and rate producers
prevent an unconditional elapsed-time promise. — MOVE-125

| Entry condition | Local action | Next condition |
|---|---|---|
| Centred, route nonempty, current != desired | Turn and return without a Position step | A later admitted call can start the transit |
| Centred, route nonempty, current == desired | Derive the frozen stride, reset elapsed, take step 1 | Off-centre Position enters progress 3 |
| Dynamic route empty | Call the turn leaf only | Desired, active flag, estimate and Position are unchanged by this path |
| Progress 3 | `R0039` takes one live step | Centred arrival clears progress and skips the pending arm in that invocation |

The empty-route leaf does not maintain the driver's active flag or estimate.
A changed desired heading continues from the current byte while active stays
nonzero. The stop reset copies current into desired without clearing active;
the next short turn can therefore advance by rate instead of snapping.
Whole command-arrival scheduling remains conditional. — MOVE-125, MOVE-107,
MOVE-090, AI-417

## Fine position and boundary calls

`R0047` increments word `+ac`, adds the signed strides to the low 16-bit
fine axes, and compares the resulting cell bytes with the stored cells. For
directions 1 and 5, resulting fractions `0/0` subtract one from fine X modulo
65536. At unsigned `+ac >= +aa`, both fractions become `0x80`; the final snap
does not run another boundary transaction. These local stores and the derived
grid are High; they are not observed runtime timing. — MOVE-126

For a centred straight transit of magnitude `s`, duration is `ceil(256/s)`
calls. Cell-byte change occurs on `ceil(128/s)` for a positive axis and
`floor(128/s)+1` for a negative axis, counting the start call as 1. With equal
diagonal magnitudes, direction 1 uses the negative-axis crossing call and
direction 5 the positive-axis call; the zero-fraction correction aligns the
two axes when `s` divides 128. Translation rate is frozen at the start of the
transit. — MOVE-126, MOVE-RATE-029, MOVE-DIR-034

**Footprint position** (`MOVE-087`). The stored cell is the top-left cell of the `n x n` footprint, with
`n` = actor byte `+0x49`. The fine value `P = (cell << 8) + sub-cell` is that corner. The centre that the
range test, the edge gap and the bearing read is `P + (n-1) * 128` per axis, 16 bits. A step moves `P` by the
step bytes whatever `n` is. A resting mover has both sub-cell bytes `0x80`, so its centre is the geometric
centre of its block. The range test returns 1 when the centre distance minus `((n1+n2) << 7) - 0x100` is at
most `0x180`, else `(v + 0x40) >> 8`; the edge gap subtracts `(n1+n2) << 7`, clamps at 0 and returns
`(v >> 8) + 1`. Which actors have `n` above 1 is not measured.

### What the cell-boundary calls do (`TERR-CELLREC-146`, `TERR-FOOTPRINT-147` (amended, partially retracted))

`R0050(map, actor)` and `R0057(map, actor, x, y)` are the enter and the leave of the
map's **cell record**, the per-cell structure [TERRAIN](../terrain/format.md) specifies. Both read the
actor's movement domain through `vt+0x20` and pick a slot from it: domain 1 or 2 uses the record's
`payload+0x04`, domain 3 uses `payload+0x08`. The per-cell entry refuses other
domain values; the detach statement here is scoped to domains 1/2/3. `R0050` reads the
footprint side once (`L05505`) and writes the actor into **every** one of the `n x n` cells it
covers, through `R0458`, one call per cell; each call fails when that cell's slot is already
taken, and one failure stops further iteration without local rollback of earlier
cells or the mover+72/+82..85 caches. Only successful cell entries reach `R0453`,
which recomputes that cell's cost byte and both block-plane bytes from the record — which is where
dynamic bits 6 and 7 come from, and where an area effect reaches the search at all
(`MOVE-AREA-038`).

The former all-or-nothing interpretation is withdrawn. A2x2 footprint can retain
its completed row-major prefix on refusal. The step caller stops its old-cell
detach loop on false but continues the position rewrite, and it does not test
entry's return before the center test. At center, the known dynamic-route cleanup
and progress 3 completion can occur despite a refused destination entry. Local
selected serializers subsequently emit that state; no rollback or next-SAVE
reconciliation is established. The local failure branches are High; later
callback and other-actor effects remain Unknown. — MOVE-127, SAV-CELLFAIL-583,
SAV-CROSSNEXT-585

Entry samples Position cell-X/Y through getters `R0299/R0300` into
caller locals before `R1345`; the domain-1 cost read uses those argument
low bytes. After normal return, the caller explicitly stores mover +72,
+82/+83 from the local cell bytes, then +84/+85 from full-X/Y getters
(`R0165/R0166`, hence Position+04/+05). Full-X/Y are sampled after the
call. Word+72 is that call's AX, not a proved numeric default; the synthetic
`0x1357` cut does not establish a state default. The native local CFG places
the call before every path to these five explicit mover stores.
Successful detach instead directly loads Position+00/+01/+04/+05 into
+86/+87/+88/+89, in that order, after recompute/optional removal. It reloads
actor+10 for each source. Neither sequence establishes callback purity or an
atomic entry snapshot. — SAV-CELLFAIL-583, SAV-CELLLEAVE-584, SAV-1161

**Crossing order and removal** (`MOVE-088`, `MOVE-089`). The release loop, the claim, the position rewrite and
the occupy are one call of the step, so the per-cell recompute runs in the tick of the crossing. A release whose
slot is empty and an occupy whose slot is taken skip that cell's recompute and end their loop. A release that finds no
cell record returns 0, and a domain outside 1..3 recomputes without a slot test. The domain 1 and 2 occupy
builds the cell's trigger caster before its slot test; the taken branch then calls a routine that is one `RET 4`.
No deferred recompute exists in these bodies. The release clears a
non-zero slot whatever actor holds it. Removal (`R0865`) runs the same release loop at the stored
position at once, then clears the claim bits when every release succeeded and the claim cell differs from the
current cell; an empty slot ends it early with 0 and the claim state is left. Whether a caller tests that
result is Unknown.

Two consequences a consumer must reproduce. A unit of footprint side `n` is present in `n²` cell
records while it stands, so anything walking cells finds it `n²` times — that is what makes
`fire_ball`'s footprint-squared divide a normalisation (`MAGIC-FIREDIV-047` (amended and partially retracted in the ledger)). And `R0453`
assigns the dynamic byte from the static byte before rebuilding bits 6 and 7 from the record, so
occupancy written straight onto the dynamic plane for a cell that holds a record does not survive
the next recompute of that cell.

### Occupancy phases

For successful footprint release and entry, with no intervening unrelated
recompute, the source and destination have these distinct slot/plane states.
— MOVE-127

| Phase | Cell-record actor slots | Dynamic plane |
|---|---|---|
| Centred before transit | Current footprint holds actor | Current occupied; next footprint reserved before turning or stepping |
| Before cell-byte change | Source footprint holds actor | Source occupied and destination reserved |
| After cell-byte change | Destination footprint holds actor | Source claim and destination occupancy |
| Centred completion | Destination footprint holds actor | Source-only cells cleared; current footprint retained |

## Centred arrival

`R0039` takes one step and, after a cell-byte change, updates word
`mover+72` from the new cell cost and actor speed. It does not replace the frozen
strides. When both fractions are `0x80`, it zeroes words `+a8/+aa/+ac`, clears
the source claim at word `+70`, zeroes words `+80/+a6`, and frees the whole
dynamic route. The source, next-cell, direction and stride fields retain their
values in this local cleanup. Claim clear excludes the current footprint.
The instruction relation is High; callback purity and reconciliation of a
refused footprint are not established. — MOVE-128, SAV-CROSSNEXT-585

After cleanup, a current cell record with payload byte `+2c == 0x1a` can
relocate the actor. Payload `+30/+31` supplies the destination. Both dynamic
and static footprint checks must pass. The helper releases the stored
footprint, writes the destination cell with fractions `0x80/0x80`, occupies it,
and calls the order reset and notification. A failed release or occupy does
not locally roll back that Position write. — MOVE-129

The reached reset copies current facing into desired, zeroes actor action
`+54`, writes actor command state `+50 = 0x0c`, stores the current cell in
order word `+00`, clears pending byte `+08` and cached target `mover+7c`, and
writes order byte `+14` from actor reach `+12c` and order dwords `+50/+60 = 0`.
Its own body leaves turn active `+a0` and order progress `+09` untouched.
The relocation/reset clauses are local; their whole outcome is Medium because
the final notification and callback effects remain Unknown. — MOVE-129

The arrival block of the step calls that relocation after it looks up the cell record in the
hash table, copies the record, and finds type `0x1a` and two passing guards; the call site
is at `L00117`. — MOVE-149

The step and its stepper touch these mover words: reads `+70`, `+80`, `+a6`, `+aa`, `+ac`, `+ae`,
`+b0`, `+b1`; writes `+72`, `+76`, `+80`, `+a6`, `+a8`, `+aa`, `+ac`; neither body has an operand at the
order record. — MOVE-148

## Blocked routes and retries

An off-centre ordinary walk or pursuit continues `R0039` and returns
before its search branches. A blocker that appears during an admitted transit
therefore does not cause a local destination check or return to the source.
The boundary occupy can still refuse. — MOVE-130, MOVE-131, MOVE-127

For the centred ordinary walk, a needed dynamic route tests the adjacent
static waypoint against the dynamic plane. If the waypoint is not exactly one
Chebyshev cell away, this test uses cell 0. A blocked footprint calls
`R0402`; nonzero turns and returns without a search or step, while zero
runs the near search. Unequal heading stores pending `0x0a`; an already matched
heading stores pending 0. Retry cadence after those pending changes is not
proved by the verdict alone. — MOVE-130, MOVE-132

Let `A` be the requesting actor, `C` the tested cell and `B` its ground-slot
actor returned by `R0035`. The verdict applies these guards in order.
Return values 1 and 2 both take the current caller's turn-and-return branch.
The table gives High local conditions; the no-occupant lookup population,
unread getter dependencies and semantic name of the `A+14` guard remain
Medium or Unknown. — MOVE-132

| Condition | Verdict |
|---|---|
| A footprint side != 1 | 0 |
| Dword `[[A+14]+28] != 0` | 0 |
| A Mover word `+80 == C` | 0 |
| B exists and signed HP word `+94 < 1` | 0 |
| B is centred, pending byte is neither 1 nor 4 | 0 |
| B is centred, pending byte is 1 or 4 | 2 when A order word `+0a != C`, otherwise 0 |
| B is off-centre | 1 when A Mover word `+06 != C`, otherwise 0 |
| B absent | 2 if the fallback list contains an off-centre actor whose Mover source word `+70 == C`, otherwise 0 |

The last row is conditional on the fallback list supplied by `R0020`;
its full cell population is unread. It is not a global absence or wait-policy
claim. The centred and off-centre rows read A's fields, not B's destination.
— MOVE-132

| Driver | Counter and trigger | Effect |
|---|---|---|
| Ordinary walk | Target differs from word `+74`, or byte `+09 > StaticRefreshRate` | Full static search, reset `+09`, discard dynamic route |
| Ordinary walk | Dynamic route empty, or byte `+78 > DynamicRefreshRate` | Blocked verdict or near search; near search increments `+09` and clears `+78` |
| Pursuit | Byte `+09 > static count/3 + 1` | Full search when cached word `+8a > StaticIsntNeeded`, otherwise single-node rebuild |
| Pursuit | Dynamic route empty, or byte `+78 > DynamicRefreshRate`, and facings match | Near search, increment `+09`, clear `+78` |

Both counters are bytes and wrap. They count admitted passes/searches, not
seconds. Ordinary walk increments `+78` only on a centred out-of-stop pass;
pursuit does likewise, with no ordinary blocked-verdict branch. A changed
pursuit target resets its cached static route, while a route-end match
increments `+09`, clears cached target `+7c` and returns. Shipped thresholds
are 16, 32 and 5. No give-up counter exists in these read driver bodies;
empty-search failure and its executor-tail consumption are separate.
— MOVE-130, MOVE-131, MOVE-PARAM-006, AI-384, AI-ROUTE-045


## Restored actor registration and the reached turn

LOAD rebuilds Player+20 from group membership, then the global actor list from
those Player lists, skipping actor+4c mask0x08. The restored registration helper
appends the exact actor pointer; the creator helper's ID assignment is a separate
path. The later manager+4 callback invokes actor+24, not actor+50. Unit uses the base hook;
Humanoid and Human delegate to it. Stage BYTE+13c=0 admits
mover reference repair at+7c. These local paths do not by themselves establish
preservation through every callback. — SAV-LOADREG-878, SAV-LOADHOOK-879

The three measured actor classes share the+18 tick. It processes attached effects
before signed HP and order admission; HP>0 and actor+3c!=0 are necessary for the
selected order call. Within an admitted order prefix, progress+9=0 without the
status hold lets pending byte+8=10 choose the explicit turn arm. Unequal current
and desired facing call the turn with the same actor; equality clears the pending
byte. The inactive short-arc turn arm snaps without reading mover+a, whereas the
other arm reads that actor's allocated byte. — MOVE-EVENT-060

A selected continuous Effect callback can call the same actor's derive first:
its original+38 -> +40 -> generic+48 route passes the unchanged actor to+50.
Unit's derive differs from Humanoid/Human's, whose reached store replaces mover+a
with low8(actor+8c). Empty/ineligible effects and special identities can skip that
route. This conditional order does not identify the first restored effect or the
first post-LOAD mover access. — MOVE-EVENT-061

The selected frontend resume can bootstrap a server call, but earlier world,
frontend, phase-dependent full-tick, command and world-object callbacks remain
between LOAD and the selected actor consumer. Absolute first producer/read order,
the byte at that read and native resume timing remain Unknown. — SAV-FIRSTMOVE-880

## Mover state at mission start

The actor base constructor builds every spawned actor's Mover: zero bytes, mask `+0x05` 0x41,
`+0x08` 5, `+0x09` 0xff, rate `+0x0a` 0x10 then 8, and facing `+0x00` equal to wanted facing
`+0x01` = 0x40 plus one `rand()` draw over 0..0x80. Of the image's four Mover constructor calls
only this one stores a facing. — MOVE-119

A type-6 placement then takes its mask from its movement domain, its rate from the Units table or,
for a Human, from the derive the spawner runs before placing it, and five fields from the occupy:
word `+0x72` from the speed word and the cell's cost byte, `+0x82`/`+0x83` the cell and
`+0x84`/`+0x85` the sub-cell. The spawner places at radius 0, so the placement's own attempts
draw no random value, and the occupy enters the footprint's cell-record slots. — MOVE-120,
MOVE-121

The start constructs the type-6 actors in record order, each drawing its facing before the next
construction, and constructs no carried member. A footprint cell that carries a cell-entry tail
makes the occupy construct a temporary caster, one more facing draw; the type-9 tails are written
after the spawner, so on the fresh-map path such a draw falls between the carried members'
placements (Medium for tails before the spawner). — MOVE-141

The join walk places a carried member the same way, then replaces its Mover and order block
with constructor-state ones: facing 0, mask 0x41, no cell fields, and rate 0x10. The walk leaves
the two embedded route lists as they were. — MOVE-122

After the join, the start's settings command `0x46` subtype 3 runs `vt+0x50` on every actor of
the sending player's list, so the Human derive stores the low byte of the speed word at `+0x0a`
of each carried member's new Mover before the restart save (Medium: delivery of the command is
unread). All 22 carried members in the
original restart-slot saves hold that derived rate (15..20, 16 of 22 differ from 0x10). That this
command is the first write after the walk is Medium as well. — MOVE-142, MOVE-122

## Other step routine and death

The second release, claim and occupy sequence is in a run of code at `L07128`..`L07129` with no direct relative transfer or literal dword reference found (Medium).
It releases, claims at the old cell, rewrites the position and occupies in the live step's order and does not test
the occupy return (`MOVE-093`).

After attached effects, signed HP at or below 0 excludes the order executor.
The selected death path does not advance the admitted step or turn; the stored
cell and fractions survive the death countdown. Thus a pre-boundary death
keeps the source cell and a post-boundary death the destination cell, with
their current fractions until teardown. Motion exclusion is High; composed
body/occupancy outcomes and callback effects are bounded separately.
— MOVE-133, MOVE-096

Teardown releases the footprint at that stored cell and calls the recentre
helper: the fractions become `0x80/0x80`, without changing the stored cell or
returning to the transit source. Claim clear runs only after every footprint
release succeeds and the nonzero `+a6` cell differs from the current cell.
An early release refusal retains the remaining slots and claim fields.
These local branches are High; later repairs and other-actor consequences
remain Unknown. The route-list freeing clause of MOVE-134 is withdrawn;
route-list cleanup across unexpanded callees remains Unknown, and complete
death aftermath remains Medium. — MOVE-134, MOVE-095, MOVE-089

## Death and map departure fields

Footprint release caches Position before recenter. A release refusal skips the
reservation tail; success clears words `+a6/+80/+76` only when the nonzero
reservation differs from the retained current cell. Occupancy clearing selects
a nonempty slot without comparing its actor pointer. — MOVE-135

Death teardown selects action/policy `0x10` and centers the retained cell.
Departure instead sets off-map bit8 and unlinks the normal actor without direct
stores to its queued order, current strike target, phase or countdown. Return
bodies place, append and clear bit8. Complete callback cleanup and all tick
callers remain Unknown. — MOVE-136

Inside `R0193`, `R0426` precedes normal-actor manager
`R0427` (High). Inside `R0426`, the manager `+2c` list precedes
global-effect manager `R0641` (High). Joining these local orders gives
manager `+2c`, global effects, then normal actors (Medium). Terminal actors
are reaped immediately after their own slot. The every-actor-before-every-area-effect
clause of MOVE-084 is retracted; its transit-cost and crossing facts stand.
Full-frame and pixel-paint order remain Unknown. — AI-474 (amended)

The measured Unit destructor passes Mover `+154` and Orders `+158` to separate
cleanup callees and zeros both pointers. Internal static/dynamic route-node
disposal and destruction timing remain Unknown. — MOVE-137
