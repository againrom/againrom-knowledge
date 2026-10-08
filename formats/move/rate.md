# Movement rate and formation gates

[Reference](format.md)

## Movement rate and clock (`MOVE-RATE-029`…`034`)

The rate is computed **once per cell transit**, by `R1088`, from the actor's current facing
and cell, and is then frozen until the next cell.

```
dir  = ((facing + 0x10) >> 5) & 7          facing is a byte, 8 directions x 32 units
dx   = [ 0, +1, +1, +1,  0, -1, -1, -1]    world+0x58eb0, clockwise from north
dy   = [-1, -1,  0, +1, +1, +1,  0, -1]    world+0x58eb8
dst  = src + ((dy[dir] << 8) + dx[dir])    world+0x58ec0, built from the two above

speed = grpAI+0x44 (u8) if nonzero         the group's slowest member's Speed, set by the
        else actor+0x8c (i16)              last FORMATION group move; else the class Speed

domain == 1 (ground):
   d = clamp((i8)(height[src] - height[dst]), -32, +32)     downhill is d > 0
   v = SpeedMultiplier * speed                              map.reg [Path Finding], ships 8
   v = v + ((v * d) >> 6)                                   arithmetic shift; uphill reduces
   c = ((u8)(cost[src] + cost[dst])) >> 1 ; if c == 0 -> 8  byte-wide add, wraps at 256
   v = v / c                                                signed
domain != 1 (the other movement domains):
   v = speed                                                no multiplier, no slope, no cost
v = clamp(v, 1, 63)

stepX,stepY = (dx*dy == 0) ? (v*dx, v*dy)                   straight: an 8-bit IMUL
                           : (trunc(v*dx*K), trunc(v*dy*K)) diagonal: K = 0.707 at L06908
ticks       = ceil(256 / (stepX != 0 ? |stepX| : |stepY|))  -> mover+0xaa
```

Per tick, the position advances by `(stepX, stepY)` in 1/256ths of a cell per axis and `mover+0xac`
counts up; on reaching `mover+0xaa` the fractions snap to the centre and the surplus is discarded.

### When the group term is set, and when it is not (`MOVE-GATE-035`…`037`)

`grpAI+0x44` is written by exactly two routines — the group Move setter `R0108` (group order
4) and the Swarm-2 setter `R0157` (order 5) — and by each of them **only when the order is
issued in formation**. One local flag decides it, and the same flag decides where each member is
sent:

```
flag = 1
mode = [[grp+0x44] + 0x30] + 0x1f          the owning Player's formation mode, default 2
if mode != 2:                              0 = never, anything else = always
    flag = mode
else:
    centre = mean over members of (fineX, fineY) >> 8       R0140, unsigned divide
    for each member:
        if max(|mx - cx|, |my - cy|) > AImanager+0xa824:    Chebyshev, whole cells
            flag = 0                                        threshold = 2, a code constant

for each member:
    if flag:  order to (X + (mx - cx), Y + (my - cy))   offsets kept at ord+0x24 / ord+0x26
              min = min(min, member Speed)               seeded 0xfa
    else:     order to (X, Y)                            every member the same cell
if flag:  grpAI+0x44 = min
grpAI+0x20 = 4 or 5 ; grpAI+0x0a = (Y << 8) | X
```

**Nothing clears it.** The only instruction in the image that writes `grpAI+0x44 = 0` is inside
`R0200`, and that routine is unreachable — no call, no reference, no immediate, and no
occurrence of its address as a dword anywhere in the image (`AI-DEAD-036`). What resets the term in
practice is allocation: every **player** order builds a new group, whose record the constructor
zeroes. A **scenario-authored** group persists, so its rate outlives the order that set it and
survives every later group command, a member dying, a member being ordered away, and a save/load —
the record is (de)serialized raw, `0x50` bytes, by `R1350`.

A consumer that wants one rule: **the group term is live for a group iff the last Move/Swarm-2
order that group received was issued in formation, and it then stays live until another such order
replaces it.**

**Turning advances before stepping.** A next-cell step starts only when current
facing byte+0 already equals desired byte+1. Its mismatch arm calls the turn
routine and returns even if that call reaches the target facing. The full
active DWORD+a0 matters: only zero admits the short-arc (<=32) snap. An active
turn, or a larger fresh turn, advances current by RotationSpeed byte+a along
the shorter arc, clamps at desired and wraps modulo256. An exact 128 tie takes
addition. This leaf has no route/list access; the older route-destruction
interpretation is partially retracted in MOVE-TURN-031.

The turn caller writes BYTE+a4=ceil(pre-step shorter arc/RotationSpeed) after
the advancing call, rather than decrementing a countdown. A fresh short snap
writes 1. It resets BYTE+9d when inactive, increments it on each call, sets the
active DWORD and clears that dword if the updated facings match. The division
arm requires nonzero RotationSpeed. The selected move/order caller chain can repeat
this step; complete order scheduling and route/callback effects remain
Unknown. These are local call boundaries, not a whole-action tick count —
MOVE-TURN-044. Serialized local treatment and the post-LOAD frontier are
SAV-TURNLOAD-822.

**The turning byte has several producers.** Unit initialization stores the
180-byte mover allocation at actor+154. Its constructor writes byte+a=16,
then actor defaults write 8 to the same object. Unit table slot 9 and Human
table slot 7 address that byte; their byte helper preserves the incoming value
on -1 and otherwise stores low8. The former final-default interpretation in
MOVE-TURN-031 is partially retracted — MOVE-RATE-052.

Human derive later writes low8(actor+8c) to mover+a after its speed derivation
and modifier fold. It can therefore replace the independent table value.
The exact Unit derive slot lacks this assignment; the measured Humanoid and
Human slots share it — MOVE-RATE-053. Effect selector 18 first adds into the
target's mover+a modulo256, then calls that same target's derive. This local
effect write alone does not establish a surviving Human bonus — MOVE-RATE-054.

The complete turn leaf gets its actor from the stack argument and reloads
actor+154; incomingECX is not its mover receiver. It changes current facing
only. The selected next-cell arm with actor+184=0 calls it and returns without
a Position change. The turn caller also uses byte+a for its estimate, while
the selected positional-rate body reads actor+8c or the formation override.
All three direct leaf callers are selected; eight of eleven direct turn-caller
sites and all unresolved indirect/rebased accesses remain outside this local
proof. Full scheduling and elapsed time remain Unknown — MOVE-RATE-055.

**The turn schedule.** A fresh turn of at most two sixteenths (an arc of 32
or less) sets the facing in the calling sub-tick. A larger turn moves the
facing byte by RotationSpeed once per actor sub-tick and ends after
ceil(arc/rate) sub-ticks: at rate 16 a quarter turn takes 4 and a half turn
8, at rate 20 a half turn takes 7 — MOVE-105. The rate is mover+a for every
actor; the local turn arithmetic reads only that byte. Each derive of a
Humans-table actor, including a hero, a hired Human, a mounted rider and a
mage, copies the low byte of its speed word into it. Whether a Human holds
that derived value at every turn is Medium, and the rate of its first turn
after spawn is Unknown. A Units-table monster's RotationSpeed is set in the
shipped table independently of Speed — MOVE-106. A turning unit does not
step. The non-self unit-target and point-target act gates need the current
facing on the 8-way heading; a self-target cast skips them. A new target
continues from the current byte; the stop reset ends the turn at the current
byte and leaves the active flag set, so the next short fresh turn steps by
rate — MOVE-107.

**The clock.** One `R0047` per actor per **sub-tick** — the counter `server+0x04`, paced by
`R0454` against `timeGetTime` at `campaign+0x3f0 = 1000/R` ms, `R` from the nine-arm ladder
`{8,10,12,14,16,20,24,28,32}` defaulting to index 4 (`SESS-CLOCK-005`). Nothing in the step routine
reads elapsed time: the rate is per tick, so cells/tick is invariant and cells/second moves with the
game-speed setting. The same loop iteration issues the `0x401` presentation tick that drives
animation, immediately after the simulation tick; `campaign+0x3dc & 1` clear stops both
(`SESS-PACE-018`).

Worked example, the shipped defaults: `Speed` 16, `SpeedMultiplier` 8, both cells cost 8, level
ground → `v = 16`, step 16, `ticks = 16` — exactly one full tick, ≈ 992 ms, per cell.

**Customisation limits (G2).** `v` is clamped to `[1,63]`; the transit is a whole number of ticks, so
`v ∈ [1,63]` yields only 27 distinct transit times and the shipped 22 `Speed` values yield 15
distinct times on cost-8 terrain; the slope term is a `>>6`; a diagonal is 0.97..1.15× of `√2 ×` the
straight time rather than exactly `√2`. Only `Speed`, `RotationSpeed` (`data.bin`),
`SpeedMultiplier` and the `Cost*` alphabet (`map.reg`) are carried by a shipped file
(`MOVE-LIMIT-033`). The spread threshold is not carried by a shipped file, being a compile-time `2`;
the formation mode is per-player runtime state whose only authored surface is trigger instant 7
(`Set formation`), which the campaign's one node uses to author 0 rather than the constructor
default (`AI-SPREAD-038`; `AI-FORM-037`, whose corpus clause reading 2 is retracted, and
`TRIG-PARAM-030`).

## Area effects (`MOVE-AREA-038`)

Every plane read on the movement path is a byte-wide bit test of the cell's plane byte against the mover's mask at `mover+0x5`,
reloaded per cell, in `R1326` and `R1331` on the static plane, `R1327`,
`R1328`, `R1329`, `R1330` and `R0144` on the dynamic plane, plus the
copies inlined in the driver. No movement routine reads an area-effect layer slot, a cell-record
occupant slot, or a plane bit by immediate, and the only mask values that exist are `0x41`, `0x44`
and `0x82`. An area effect therefore reaches the search through exactly two writes, both made by
`R0453` when it recomputes a cell from its record:

- the **Wall of Earth** layer slot `payload+0x20` sets bits 0 and 2 on both planes, so masks `0x41`
  and `0x44` are blocked and `0x82` is not — movement domains 1 and 2 stop, domain 3 crosses;
- **any** occupied layer slot multiplies the cell's cost byte by 4 (an 8-bit `SHL`), which the
  `movementType == 1` cost arm reads inline at the destination cell. Two layers on a cost-16 cell
  truncate the byte to 0.

The step-duration routine `R1088` does not read the cost plane directly: it goes through
`R1087`, which for a cell with any layer returns `cost >> 2` and **stores that value back**
into the plane.

`R0144` is a standalone `n x n` footprint query on the dynamic plane, instruction-for-
instruction `R1331` apart from the displacement. It is not part of the search: 5 call sites
in 2 owners, `R0209` and `R0142`.

## Refresh policy

| trigger | effect |
|---|---|
| a cell transit completes | the **whole dynamic route** is freed (`R0039`) |
| `mover+0x78` (per tick) > `DynamicRefreshRate` | dynamic re-search; counter reset |
| ordered target != `mover+0x74` | static re-search; the dynamic route is freed |
| `mover+0x09` (dynamic searches since the last static one) > `StaticRefreshRate` | static re-search |

The dynamic search does not aim at the final goal: it takes the static route's tail waypoint when it
is more than `DynamicByStaticLookup` cells away, else the node `DynamicByStaticLookup + 1` further
along, and the final goal only when the static route holds `StaticIsntNeeded` nodes or fewer. A
waypoint is popped once the unit is within `DynamicByStaticLookup` of it. Every counter is per-unit
and reset on use — **nothing is staggered**, and neither terrain change nor target movement
invalidates a stored route by itself.

## Pre-search gates (`MOVE-GATE-039`, `MOVE-STEP-040`)

Movement is not gated inside the search or the step. It is gated by the per-actor order machine,
one switch above the walk order.

The routine that writes an actor's position is `R0047`. Closing the call graph above it
(`EnumRefs callto:`, 0 orphan at every step):

```
R0047  <- R0039, R0054
R0039  <- R0016 @L06938, R0178, R0043
R0054  <- R0178, R0043
R0178  <- R0016 @L01932, R0087
R0043  <- R0016 @L06939, R0042
R0087  <- R0016 only        R0042 <- R0016 only
R0016  <- R0037 (actor vtable slot 6), R0038 (unreachable)
```

`R0038` is a loop that calls the machine over a list; `callto:` returns 0 hits for it, and a
raw scan of every section for the stored dword `R0038` also returns 0, so nothing reaches it.
So **an actor is displaced only from inside one call of `R0016`**, and any state of that
machine which does not reach the walk arm or progress arm 3 leaves the position untouched. The
actor is not slowed and does not drift.

The states that do this:

| state | set by | effect |
|---|---|---|
| `ord+0x09 = 4` | `L00546`, when `actor+0x144 & 0x100000` (spell 20) and the byte is 0 | no order arm runs until the bit clears (`MAGIC-ACTGATE-079`) |
| `ord+0x09 = 0xff` | `L00545`, when the queued command is `actor+0x50 == 0x17` | no order runs; cleared only by `R0146` |
| `actor+0x54 == 0x10` | elsewhere | `R0037` returns at `L05756` before the machine runs |

A refusal takes effect only from `ord+0x09 == 0`, which `AI-ORDER-039` (whose arm-`0xb` clause is
retracted; this clause stands) fixes as the tick the actor
stands on a cell centre (`R0040` testing `pos+0x4 == pos+0x5 == 0x80`). An actor in transit
is in progress state 3, whose arm still calls `R0039` and clears the byte on arrival. **A
stopped unit is therefore always aligned to the grid, never caught between cells.**

The attack tail has a separate gate. The machine's tail at `L00010` is not
gated, and it can install an attack (`MAGIC-ACTGATE-079`); it runs only when `mover+0x98` is
non-zero, and that flag's three setters — `L01720`, `L00161`, `L00162` — are all inside
`R0178`, `R0043` and `R0055`, which are reached through the gated order arms.
A refused actor can therefore never raise the flag again, and the tail clears it at `L05750`.
