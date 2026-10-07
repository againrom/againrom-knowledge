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
- **Blocked next cell:** the unit turns to face it and waits. It never pushes, swaps or steps aside.
  The step routine itself tests no plane, so two units whose routes predate each other's claims can
  enter the same cell.
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


## Movement step

Position is `actor[4]` = `*(actor+0x10)`: `+0x00` x cell, `+0x01` y cell, `+0x02` packed cell,
`+0x04` x fraction, `+0x05` y fraction, `0x80` = centred. `R0047` treats `(cell<<8)|fraction`
as one 16-bit value per axis and adds the signed per-axis step `mover+0xb0` / `mover+0xb1`; when the
cell byte changes it hands `R0058` the **old** packed cell, then writes the new position, then
calls `R0050`; when `mover+0xac >= mover+0xaa` it snaps both fractions to `0x80`. A transit
takes `mover+0xaa = ceil(256 / |step|)` ticks. Speed never enters a label — two units of different
speed pick the same route.

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
reconciliation is established. — SAV-CELLFAIL-583, SAV-CROSSNEXT-585

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

## Walk step at order recovery

A fresh call of the walk routine writes a sub-cell step in that call only when the facing byte already equals
the direction to the first path node; otherwise it turns and the step follows at the next call (`MOVE-090`).

## Step routines, transit ticks and teardown

The second release, claim and occupy sequence is in a run of code at `L07128`..`L07129` with no direct relative transfer or literal dword reference found (Medium).
It releases, claims at the old cell, rewrites the position and occupies in the live step's order and does not test
the occupy return (`MOVE-093`). In the live step a transit from the cell centre with an axis step s crosses on
tick ceil(128/s) on a positive axis and floor(128/s)+1 on a negative one, the start tick counting as 1; release and
occupy are called by the step on that tick only (`MOVE-094`). A removal whose release ends early on an empty slot returns before
the claim clear, so the claim bits the transit set stay in the plane unless a store outside the teardown body and `R0866` rewrites them
(`MOVE-095`, Medium). A unit at or below 0 health is not stepped; its slot and claim stay frozen through the death
countdown, and the teardown releases at the stored cell and clears the claim when `+0xa6` is not the current packed cell (`MOVE-096`).
