# Actor action phases

[Reference](format.md)

## State and identity

Animation fields belong to the client drawable CUnit. The simulation actor
holds its action-end tick at+0x138; drawable frames/phases are not serialized.
The paced0x401 game tick drives presentation at `1000/tps` milliseconds,
with `tps={8,10,12,14,16,20,24,28,32}` and default 16. Truncating division
gives periods `{125,100,83,71,62,50,41,35,31}` milliseconds. Stopping the
tick stops animation. Idle, attack, shoot, cast, death and placed-object
timelines advance by ticks; walking advances by distance, one step per 1/16
cell. — ANIM-CLOCK-001, ANIM-WALK-013, ANIM-AMBIENT-016

## Drawable state block

```
+0x6c  facing, 16-way                       +0x94  the running action's clock
+0x70  phase, consumed by the frame switch  +0xa0  ticks the action still has to run
+0x74  draw state = a copy of +0x84         +0xbc  turn accumulator, 1/16 of a facing step
+0x84  action code, one byte                +0x15a corpse stage (the server's, shipped)
+0x85  the facing the action ends on        +0x88/+0x8c  remaining position delta, 1/256 cell
```

## Tick driver

```
R0333 (map view message handler), arm msg == 0x401
  -> R0334
       CMapView+0xa70 += 1                        the object / water animation counter
       for each drawable in the view's container:
            vt+0x3c()                             = R0548 for a CUnit
            returns false -> drop it from the walk

R0548:
  ticksRemaining (+0xa0) == 0  ->  state := 0
        IdlePhases != 0 : clock++ ; phase := clock % len(IdleAnimTime expansion)
        IdlePhases == 0 : phase := 0 ; clock := 0     (the stand / corpse fork draws)
  otherwise                     ->  switch (actionCode - 1), 8 arms
        ... the arm updates phase / position / facing ...
        state := actionCode ; ticksRemaining -= 1
```

## Action states

| code | run | what advances the phase | run length `+0xa0` |
|---|---|---|---|
| 1 | move | `clock += \|step\|`, `phase = clock/16`; the draw takes it **mod** `len(MoveAnimTime expansion)` | the message's duration byte (the mover's per-step ticks) |
| 2 | — | nothing; the arm only assigns the state | — |
| 3 | attack | `phase = clock`, `clock++`; **no modulo** | `len(AttackAnimTime expansion)` |
| 4 | — | nothing | — |
| 5 | turn | facing interpolates by the shortest arc; the draw takes the **stand** arm | the message's duration byte |
| 6 | die | `phase = clock`, `clock++`; the draw halves it | `2 × classes[Dying].DyingPhases` |
| 7 | shoot | as attack, plus the projectile spawn on the tick `clock == ShootDelay` | `len(AttackAnimTime expansion)` |
| 8 | cast | as attack, over `ShootDelay` / `AttackDelay` | `len(AttackAnimTime expansion)` |

Codes **2 and 4 are set by no immediate anywhere in the image**, which is why the frame switch's
arms 2 and 4 fall to a default that draws the class id.

Attack animation length comes from the class timeline. The client does
not read the attackChargeTime+attackRelaxTime duration in the server's
attack message; animation and simulation action durations can differ.

**A walk cycle is advanced by distance, not by time.** The clock `+0x94` accumulates the
**Euclidean** length of each tick's own displacement — `FSQRT` of `sx² + sy²`, truncated per tick —
so it is an odometer in 1/256 of a cell, and `phase = clock >> 4` is one timeline step per **1/16
of a cell**. No speed term appears anywhere in the animation code: a slow unit takes more ticks to
cover the same distance and holds each frame longer; it never skips one.

- **A straight cell is exactly 16 timeline steps, for every unit at every speed**.
  A **diagonal** is `256√2 = 362` in that metric, so 20…22 steps after the per-tick truncation.
- **How many *cycles* those 16 steps are is data**: `class+0x50 = len(MoveAnimTime expansion)`, and
  `16/L` is the cycles per cell. The shipped convention is `MovePhases = 8` frames held two steps
  each, `L = 16`, **one cycle per cell**. `Bee` and the two `Catapult`s run four
  cycles per cell, `Ghost` two, Other installed pairs have `16 mod L != 0` and are not cell-aligned at all.
- **The clock is not reset when a walk starts**, so consecutive cells continue one odometer. It
  *is* zeroed by any turn, attack, shoot or death message, and by standing still for a single tick
  when `IdlePhases == 0` (`ANIM-WALK-013`…`015`).

## Idle and fidget

State 0 selects idle animation or a standing/corpse frame according to the
class IdlePhases value. — ANIM-IDLE-009, ANIM-IDLE-010

```
state 0, IdlePhases != 0     clock++ once per tick; phase = clock % len(IdleAnimTime expansion)
                             frame = idleBase + dir*IdlePhases + idleTL[phase]   (NO modulo)
                             idleBase = S + D*(MoveBegin+Move+Attack+Dying)      = the BONE base
state 0, IdlePhases == 0     phase = 0, clock = 0
                             corpse stage >= 1 -> the dying/bone fork; else the standing frame
```

Three consequences, each of which a consumer gets wrong by default:

- **The corpse fork is unreachable for a class with idle art.** The draw reads `IdlePhases` before
  it reads the corpse stage, and the idle and bone blocks share one base — so a class either
  fidgets or leaves bones, never both, and the sheet layout is consistent because it has to be.
- **The idle clock is the action clock.** `+0x94` is not reset when an action ends, only when one
  starts, so a unit resumes its idle loop at `(whatever the last action left) % len` — two units of
  one class idle out of phase. Reproducing this needs no extra field; forgetting it makes a group of
  monsters flap in lockstep, which the engine never does.
- **The loop length is the expansion's length, not `IdlePhases`.** `Ghost` carries 3 frames and a
  4-step timeline `0 1 2 1` — a ping-pong. `IdlePhases` sizes the *block*; `class+0x80` sizes the
  *loop*.

Shipped data, all five classes, one step = one animation tick (62.5 ms at the default speed index):

| class | `IdlePhases` | timeline (expanded) | loop | cycle @16 tps | placements / 8094 |
|---|---:|---|---:|---:|---:|
| Bee (73) | 4 | `0 1 2 3` | 4 | 250 ms | 608 |
| Ghost (69) | 3 | `0 0 0 1 1 1 2 2 2 1 1 1` | 12 | 750 ms | 416 |
| Sonic Bat (70) | 6 | `0 0 1 1 2 2 3 3 4 4 5 5` | 12 | 750 ms | 580 |
| Dragon (71) | 7 | `0 0 1 1 … 6 6` | 14 | 875 ms | 135 |
| Death Star (72) | 7 | `0 0 1 1 … 6 6` | 14 | 875 ms | 1 |

The other 29 classes carry no `IdlePhases`, and their whole idle presentation is the standing
block: one frame per facing, changing only when the facing does.

## Directions

One 16-way facing field. The standing block indexes it whole, `(facing − 8) & 0xf`; every other
block halves it, `((facing − 8) >> 1) & 7`; `Flip` mirrors the upper half (`SPR256-UNIT-024`,
`REG-UNITS-051`). There is no separate turning run: while the turn action is active the unit is
drawn standing at a facing that advances by `(target − current)/ticksRemaining` sixteenths per tick,
wrapping by the shortest arc.

## Action messages

`R0509` dispatches on `msg+0x9` biased by 3, through a byte arm-index table at `L02523`
into a jump table at `L02524`.

```
0x6b  action 1 move    facing, duration, and the per-direction delta from L02525 / L02526
0x6d  action 5 turn    facing, duration
0x71  action 3 attack  facing            (refused while an action is running, or AttackPhases == 0)
0x72  action 7 shoot   facing, and the target's runtime id at +0x86
0x86  action 1 or 8 · 0x8a action 8 · 0x8b action 1 · 0x8c action 1
0x6c 0x6e 0x6f 0x70    the field-masked state sync — one bit per field:
      0 health  1 mana  2 +0x108  3 corpse stage  4 facing  5 position
```

## Death and corpse phases

```
server   actor+0x13c = 1 at death; R0867 raises it to 2 / 3 / 4 as health passes
         -10 / -20 / -40, and to 5 below -600; every change is broadcast
client   the stage arrives in the state sync under mask bit 3; the client keeps the previous
         value and, on every state sync it applies, dispatches on the TRANSITION:
             0 -> 1   action 6, 2 x DyingPhases ticks, clock 0        the fall
                      and the hurt hook with k = 3, called as the message is handled
             1 -> 1   action 6, 4 ticks, clock 2*DyingPhases - 4      the last two frames again
                      and the hurt hook with k = 2, called as the message is handled
             N -> 0   N != 0: the action is cleared
         when the run ends, no action runs and the frame switch's state-0 corpse fork draws
         on the stage alone:
             1        dying frame DyingPhases - 1, frozen
             2 3 4    bone frames 0 1 2
             5        the actor is gone
```

The hurt hook reads a voice bank's `die` file for `k = 3` and its `hard` file for `k = 2` when
`+0x18c & 0x11` is non-zero, and the class's `Sound[4]` and `Sound[3]` otherwise; `k = 2` stays
silent within 1500 ms of the drawable's voice timestamp `+0x190`, which `k = 3` neither tests nor
sets (`ANIM-094`, `ANIM-095`).

The named death, bleed and session-entry routes stand. A projector call is not
necessarily an ordinary state-sync send: mask 0 and mask `0x80` alone do not
emit one; effective-mask admission is ANIM-129. The finite remaining caller
predicates are in [State-projector callers](state-projector.md), with dynamic
attachment masks and native reachability Unknowns (ANIM-131, ANIM-132).
Restoration that crosses zero sends server stage 0 when dirty is used;
restoration remaining nonpositive can leave client stage 1 (ANIM-130).
The three intervening client helpers do not clear stage at native bindings
(ANIM-133). The death arm sends mask
`0x409` to every player once, on the tick it sets stage 1 (`k = 3`). The negative-health bleed
sends mask `0x1` to the owner only, every fourth full tick, at most 9 times for a corpse that starts
at health -1 to -9 on the live list, with the stage unchanged (`k = 2`). The session-entry projection sends every actor of the world list and
every dead-list actor below stage 5 to the joining player once, so each stage-1 corpse plays
`k = 3` once per join or load (`ANIM-126`). Event 0 of the hook comes from a `0x73` message of a
zero-damage unit strike that leaves the target above -10; area damage and the Token-8 applier send
only for non-zero damage, and a structure target gets tag `0x82` (`ANIM-125`).

The thresholds never reach the client as numbers — it sees only the stage. The four
`movementType > 1` classes (`Ghost`, `Bee`, `Bat_Sonic`, `Dragon`) take `health = -1000` the moment
`dyingTime` expires, so they pass −10 without pausing, never occupy stages 2..4, and leave neither
corpse nor bones.

`Death Star` (ID72, `Daemon`) has absent `BonePhases`, loaded as `-1`.
The corpse arm checks only `==0`; its bone index is therefore
`BaseBone − dir + stage − 2`, inside the dying block and decreasing with
facing. Other installed unit/face pairs have `BonePhases >=3`.

## Objects

Same clock, different mechanism, no second driver. The routine that ticks every drawable also
increments `CMapView+0xa70`, and a placed object's frame is a pure function of that counter and its
cell — `Index + T[(animCtr + col·(row+1)) mod class+0x3c]`, recomputed at every draw, with no
per-instance state at all (`TERR-SPR-042`, `REG-OBJ-046`). The animated arm's gate is the four-corner `0xc000` test that 0 of
880 704 cells of a shipped **map file** satisfy (`TERR-TILE-044`) — but those two bits are
**runtime** state: `vt+0x48` ORs `0xc0` into the tile word's high byte at exactly those four corners
around every drawable, and the 32-tick sweep `R0374` clears bit 14 map-wide (`ANIM-TICK-011`).
Those two bits are the **fog of war**, and the gate holds on exactly the cells the local player can
currently see, so the arm **does** fire: a shipped map's fires and trees animate inside the field of
view and hold frame 0 outside it (`TERR-TILE-079`; `ANIM-OBJ-008`, whose unreachability clause is
withdrawn).

**Cadence, which is the thing the walk does not share.** `CMapView+0xa70` advances an object's
timeline **one step per tick**, unshifted. The ambient *repaint* is throttled by a second field:
`CMapView+0xa74` latches the counter, and a redraw is admitted only once it has advanced by more
than 3 — **one ambient frame per 4 ticks**, i.e. 248 ms at the default speed. Neither field is the
walk clock and neither is the mover's tick count: they are a third counter on the same pacer, and
no game-speed setting changes the ratio between them (`ANIM-AMBIENT-016`, `ANIM-PACE-017`).

## State requirements

One tick source at the game-speed rate; per unit, the eight-field block above; the eight arms
**and the state-0 idle fork**; the frame switch of [TERRAIN](../terrain/format.md); the sheet layout of
[SPR256](../spr256/format.md). For objects, one
counter and a pure function of `(animCtr, cell)`. Nothing here belongs in a hash of simulation
state, and nothing here belongs in a save.

**And two cadences, not one.** Everything above runs at one timeline step per tick; the walk runs
at one step per 1/16 cell. A consumer that implements a single "animation frame counter" produces a
walk that is in step with the world for units of `Speed` 16–17 and out of step for every other, and
no amount of tuning the tick rate fixes it — the ratio is the same at all nine game speeds
(`ANIM-WALK-013`, `ANIM-AMBIENT-016`, `ANIM-PACE-017`).

## Refused-actor state (`ANIM-PARK-039`)

`R0016`, the per-actor order machine, runs immediately before the actor tick's own
act-state switch, and it can leave the actor in a state that switch has no arm for.

The actor tick's switch reads `actor+0x54` at `L00538`, subtracts 1, and compares the result as unsigned against 14 at `L02934`; an out-of-range value branches at `L02935` to `L02936`, the epilogue, not a fifteenth arm.

The machine writes `actor+0x54 = 0x1a` from four sites: its early-out at `L00089`, the refusing
progress arm at `L02932`, the `ord+0x09 == 0xff` arm at `L00544`, and the default progress arm.
`0x1a − 1 = 0x19` exceeds the switch's bound of `0xe`, so none of the fifteen arms runs and the tick
returns. The value means "the machine had nothing to run" and carries no information about why, so
it cannot be used to identify a spell, a death or a halt.

What the client draws for such an actor is decided on the drawable, not here. For `stone_curse` the
unit draw replaces the frame index with the unit's facing, so the frame no longer follows the
animation clock, and separately replaces the shade table with a greyscale one; both are keyed on the
effect record's `+0x0e` rather than on `actor+0x144` (`MAGIC-STONEDRAW-084`). The two are separate
mechanisms in the two hierarchies this ledger's header keeps apart: the draw substitution is
presentation and is not serialized, the action refusal is simulation and is.
