# ANIM — what drives a unit's animation

Not a file format. This ledger owns the **boundary** between the simulation and
what is drawn: how an actor's behaviour becomes a message, how the client turns
that message into a frame per tick, and which of the two sides owns each field.
`claims/spr256.md` owns the sheet's layout (`SPR256-UNIT-024`),
`claims/terrain.md` the switch that turns `(state, phase, facing)` into an index
(`TERR-SPR-047`) and the blit; `claims/hero.md` and `claims/move.md` own the
simulation-side timers this ledger's runs are compared against. Format of this
file: [registry.md](registry.md). IDs are permanent.

## Terms

Vocabulary is the engine's own field layout. The **drawable** is `CUnit`
(vtable `L02468`, the `CRuntimeClass` at `L00620` names it; `CAirUnit`
and `CProjectile` derive from it). The **actor** is the simulation object of
vtable family `L00001`/`L00002`/`L00003`, whose class names in the
engine's own runtime-class table are `Unit` / `Humanoid` / `Human`. The two
hierarchies are disjoint and only the actor is serialized.

## Clock, draw state and phase

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-CLOCK-001 | Animation is presentation state, and it advances on the game tick — outside the simulation, inside the simulation's clock. | High / Medium | ● active | [EXP-0071](../experiments/EXP-0071-animation-driver/) |
| ANIM-STATE-002 | The nine-value draw state is a copy of a one-byte action code, and two of its nine values are set by nothing. | High / Medium | ● active (amended) | [EXP-0071](../experiments/EXP-0071-animation-driver/) |
| ANIM-PHASE-003 | What advances the phase is per action, and the walk cycle is advanced by DISTANCE, not by time. | High | ● active (amended) | [EXP-0071](../experiments/EXP-0071-animation-driver/) |
| ANIM-RUN-004 | A run's length comes from the ART, not from the duration the simulation ships — and the two disagree on 143 of 262 shipped pairs. | High / Medium | ● active | [EXP-0071](../experiments/EXP-0071-animation-driver/) |
| ANIM-MSG-005 | Eleven message opcodes fill the action block, and the mapping is read out of the dispatcher's own two tables. | High / Medium / Unknown | ● active (partially retracted) | [EXP-0071](../experiments/EXP-0071-animation-driver/) |
| ANIM-DIR-006 | The facing is sixteen-way, it is client state, and a turn is the standing frame at a moving facing. | High | ● active | [EXP-0071](../experiments/EXP-0071-animation-driver/) |
| ANIM-DEATH-007 | The whole death animation is driven by one byte the server ships, and `REG-UNITS-050`'s "what advances `unit+0x15a`" is answered: nothing on the client does. | High / Medium | ● active (amended, partially retracted) | [EXP-0071](../experiments/EXP-0071-animation-driver/), **[EXP-0405](../experiments/EXP-0405-hurt-voice-bank/)** |

### ANIM-CLOCK-001

No simulation actor carries a frame, a phase or an animation state: its entire
animation surface is `actor+0x138`, the tick the current run is due to end
(written at `L02469`/`L02470`, read only by `R0547`, 3 callers), plus
the notifications it broadcasts. Every animated quantity lives on the
**drawable** `CUnit` — `+0x70` phase, `+0x74` draw state, `+0x84` action code,
`+0x85` target facing, `+0x88`/`+0x8c` remaining position delta, `+0x94` the
action's clock, `+0xa0` ticks remaining, `+0xbc` the turn accumulator, `+0x15a`
corpse stage — and the drawable is **not** among the 28 `CRuntimeClass` records
at `L02471..L02472` that the save stream carries (`SAV-STREAM-013`), so
no animation quantity is ever written to a save. The advance is `R0548` =
`CUnit vt+0x3c`, and the chain above it is single-caller at every step:
`R0334` calls it for every drawable in the map view's container
(`L02473`, `L02474`) and increments the object/water counter
`CMapView+0xa70` in the same pass (`L02475..L02476`);
`EnumRefs callto:R0334` → **1 hit, 1 owner, 0 orphan**, and that hit is
`L02477`, arm 0 of `R0333`'s `msg − 0x401` jump table — the **`0x401`
paced play-loop tick** of `TERR-ANIM-008`, one per `1000/tps` ms with
`tps ∈ {8,10,12,14,16,20,24,28,32}` from the game-speed index (16 tps, 62 ms, by
default). **Consequence for a consumer:** the frame is not hashed state and need
not be reproduced deterministically, but it must be stepped by the same tick as
the simulation — it slows when the game speed drops and stops when the tick
stops

**Confidence.** **High** (the field census, the caller chain and the tick's
identity are named instructions; the rival "a field of simulation state advanced
once per simulation tick" is excluded by *absent* instructions, not plausibility
— a `disp:70` sweep returns 581 hits over 276 owners and a `disp:74` sweep 465
over 243, and **not one owner of either is in the actor module**
`L02478..L00612`; and a serializable drawable would have to appear in the
28-record table, which is enumerated in full) / **Medium** (that no
`memcpy`-shaped write reaches these fields: a displacement sweep cannot see a
wholesale structure copy)

### ANIM-STATE-002

`TERR-SPR-047`'s `unit+0x74` is assigned at the common tail of every arm of
`R0548`: the sign-extended action-code byte `+0x84` is stored as the
32-bit state `+0x74` (`L02479`/`L02480`), alongside `L02481`, which
decrements the ticks-remaining counter `+0xa0`. The action code `+0x84` selects
the arm through an 8-entry jump table at `L02482` indexed `code − 1`
(`ScanField data:L02482:8`): **1 move · 2 → the bare tail · 3 attack · 4 → the
bare tail · 5 turn · 6 die · 7 shoot · 8 cast**, and `code > 8` also falls to
the tail (`L02483`: an unsigned compare of the biased index with 7 skips the
table). With `+0xa0 == 0` no arm runs at all and
the state is forced to 0 (`L02484`). **Codes 2 and 4 are written by no
immediate anywhere in the image**, which is why `TERR-SPR-047`'s arms 2 and 4
fall to a default that draws the class id: they are unreachable

**Confidence.** **High** (the copy and the table are read instructions, the
table is dumped from static data; the write census is complete under the field's
own width — `+0x84` is a byte, so all **20** byte-wide writes image-wide are
listed by address in the experiment's `evidence/rom-anim-excerpt.md` §C, and
every immediate among them is one of `{0,1,3,5,6,7,8}`; a dword hit at that
displacement would span `+0x85`, the target facing, and there is none in the
drawable module) / **Medium** (the stronger reading *"2 and 4 are never set at
all"*: four of the 20 writes take a register — the base constructor `L02485`
writing the constant 0 loaded three instructions earlier, the copy constructor `L02486`
copying another instance's own byte, and `L02487`/`L02488`, both outside the
class and not shown not to alias a `CUnit`. **Both re-read by [EXP-0084]** and
neither is an origin: `L02487`'s stored byte is the return of the key/value getter
`R0452` called with the field's own current value as the default
(`L02489`), under the keys `"action"`/`"actiondir"`/`"actionphase"`/… — a
**restore** of a value that store was given; `L02488` writes `[arg+0x154]`,
not the argument, beside two stream-fed byte stores to `+0x82`/`+0x83`
(`L02490`, `L02491`). Narrowed, **not lifted**: the first can reproduce
whatever a drawable already carried)

**Amended.** The Medium clause, that codes 2 and 4 are never set at all, is
narrowed and not lifted by EXP-0084's re-read of the register writers `L02487`
and `L02488`; **Confidence.** carries that re-read.

### ANIM-PHASE-003

One call of `R0548` per tick. **move (1)**: the arm divides the remaining
delta `+0x88`/`+0x8c` by the ticks remaining and adds it to the `/256` position
(`L02492`, `L02493`, `L02494`), then adds the step's length —
the root of the sum of the two squared per-axis steps, `L02495..L02496` — to the
clock `+0x94` and sets `phase = clock / 16` (`L02497`, an arithmetic right shift by 4); the draw
then takes it **modulo** `len(MoveAnimTime expansion)` (`L02498`). One walk
frame is therefore 16/256 of a cell of travel, independent of speed, so a slow
unit holds each frame longer and no speed term appears anywhere in the animation
code. **attack (3) / shoot (7) / cast (8) / die (6)**: `phase = +0x94`,
incremented once per tick (`L02499`, `L02500`, `L02501`), indexed with
**no** modulo. **turn (5)**: no run of its own — `+0xbc` interpolates the facing
in 1/16 steps toward `+0x85 << 4` by the shortest arc (`±0x80` wrap at
`L02502`/`L02503`), divided by the ticks remaining, and
`facing = +0xbc >> 4` (`L02504`); the draw takes the **stand** arm at that
facing. **idle (state 0, `IdlePhases != 0`)**: `+0x94` free-runs and
`phase = +0x94 % len(IdleAnimTime expansion)` (`L02505`); with
`IdlePhases == 0` both are pinned to 0 (`L02506`, `L02507`) and the draw
takes the stand/corpse fork. *(amended by [EXP-0111]: the pairing below is a
MERGE across arms and reads as though every arm fired all four. Per arm it is
attack -> `vt+0x64` only, shoot -> `vt+0x58` then `vt+0x64`, cast -> `vt+0x60`
and `vt+0x5c`, die and turn -> none: `ANIM-STATE-023`. And `vt+0x64` is a sound,
while `vt+0x68`, a fifth hook, is fired by the message dispatcher rather than by
this driver: `ANIM-SND-022`.)* The hooks the arms fire are class scalars, on the
exact tick the clock equals them: `ShootDelay` → `vt+0x58` (the projectile
spawn, gated `Projectile != 0`) and `vt+0x5c`; `AttackDelay` → `vt+0x60` and
`vt+0x64`

**Confidence.** **High** (every arm read at instruction level; the FPU sequence
is quoted from the raw listing, not from a decompilation)

**Amended.** The hook pairing at the end of the body is a merge across arms. The
per-arm firing is `ANIM-STATE-023`, and `vt+0x68` is fired by the message
dispatcher, not by this driver (`ANIM-SND-022`); the inline EXP-0111 note in the
body carries the correction.

### ANIM-RUN-004

`+0xa0`, the ticks the action lasts, is set by the message arm and decremented
once per tick at `L02481`. For **attack and shoot** it is `class+0x68`, the
length of the `AttackAnimTime` expansion (`L02508`/`L02509` for opcode
`0x71`, `L02510`/`L02511` for `0x72`) — the duration byte the server puts in
the message is **not read by either arm** (0 hits for the operand text `0xd]`
over the whole `0x71` arm `L02512..L02513`). For **death** it is
`2 × classes[Dying].DyingPhases` (`L02514` doubles the count / `L02515`), which is
exactly what the draw's `phase/2` needs to play the run once. For **move and
turn** it is the message's own byte `msg+0xd` (`L02516`, `L02517`), which
the mover fills with its per-step tick count. Both attack arms additionally
refuse while `+0xa0 != 0` (`L02518`, `L02519`) and when `AttackPhases == 0`
(`L02520`). Because every run is sized from the art and its clock starts at 0,
**no arm can index outside its own run**, and the corpus confirms the three
checks the arms cannot make for themselves: every `AttackAnimFrame` value lands
inside its own direction slot **262/262**, `ShootDelay` and `AttackDelay` both
fall inside the run **262/262** (a hook outside it would never fire), and
`BonePhases ≥ 3` on **261/262**. Measured over the same 262 (class, definition)
pairs: `len(AttackTL) − (attackChargeTime + attackRelaxTime)` =
`-36:1 -22:1 -14:4 -4:15 -2:16 **+0:119** +2:1 +12:5 +13:5 +14:71 +16:23 +26:1`,
and `2·DyingPhases − dyingTime` = `**+0:114** +9:132 +15:16` — so the fall
outlasts the corpse's hold on its cells on 148 pairs

**Confidence.** **High** (each `+0xa0` source is a named instruction, and the
"not read" is a text search over one enumerated arm, which is the shape that can
carry it) / **Medium** (the two difference histograms: a corpus statistic over
`units.reg` × `Data.bin`, 34 classes resolved to 262 definition rows by
`R0495`'s rule, EN root)

### ANIM-MSG-005

`R0509` dispatches on `msg+0x9` biased by 3 (`L02521` subtracts 3,
`L02522` bounds the biased index at 0xbb), through a 188-byte arm-index table at `L02523` into
a 46-entry jump table at `L02524` — both decoded from the image by
`tools/animdrv -mode dispatch`, and every field-writing instruction attributed
to the arm containing it. `0x6b` → action **1 move**, with `+0x85 = msg+0xc`,
`+0xa0 = msg+0xd` and the remaining delta taken from **two per-direction tables
at `L02525` and `L02526`** shifted left 8 (`L02527`, `L02528`).
`0x6d` → **5 turn**. `0x71` → **3 attack**. `0x72` → **7 shoot**, and it alone
also stores `msg+0xe`, the target's runtime id, at `+0x86` (`L02529`). `0x86`
→ **1** or **8 cast**; `0x8a` → **8**; `0x8b`, `0x8c` → **1**. Opcodes `0x6c`,
`0x6e`, `0x6f`, `0x70` share one arm (`L02530`): a **field-masked state
sync** whose mask at `rec+0x14` gates health `+0xfc` (bit 0), mana `+0x13c` (1),
`+0x108` (2), **corpse stage `+0x15a` (3)**, facing `+0x6c` (4) and position
`+0x8` (5), each read from its own decode slot in `L02531..L02532`. On
the producing side, `R0549` — the only caller of both builders, itself
called only from the swing start `R0245` — sends `0x72` when
`actor+0x12c > 1` and `0x71` otherwise, carrying
`attackChargeTime + attackRelaxTime` as the duration; and `R0550`,
reached from arm 1 of the actor tick's 15-entry state table at `L01776`,
sends `0x6d` when the actor's `1/256` position has changed since its caller
sampled it and `0x6b` when it has not and both fine coordinates read `0x80`
(cell-centred), the latter carrying `2 × mover+0xae` as the facing and
`mover+0xaa` as the duration

**Confidence.** **High** (the tables are static data decoded from `rom.exe` by
RVA with no disassembler in the path, and each attribution is an address inside
one arm's own range) / **Medium** (the *reading* of `R0550`'s two arms as
"the position changed" vs "centred and about to step": the tests are
instructions, the situation each serves is inference) / **Unknown** (what
distinguishes `0x86`/`0x8a`/`0x8b`/`0x8c` from `0x6b`/`0x71`)

**Amended.** The Medium reading of `R0550` is refuted by EXP-0498: 0x6d goes when the position is unchanged and the desired facing byte changed, and 0x6b when the position changed from a centred start (`ANIM-134`; see [`retracted.md`](retracted.md)). The opcode-to-action mapping stands.

### ANIM-DIR-006

`unit+0x6c` is written on the client by three routes and by no other: from the
action block's `+0x85` when a move or attack arm runs (`L02533`, `L02534`,
`L02535`), by the turn arm's interpolation (`L02536`), and directly from the
state sync under mask bit 4 (`L02537`). The draw's stand arm indexes it whole,
`(facing − 8) & 0xf` (`L02538`); every other block halves it,
`((facing − 8) >> 1) & 7` (`L02539`/`L02540`); `Flip` mirrors the upper half
(`SPR256-UNIT-024`, `REG-UNITS-051`, unchanged). So a consumer needs **one**
facing field, quantised to 16, and needs no separate "turning" animation: while
`+0x84 == 5` the unit is drawn standing, at a facing that advances by
`(target − current)/ticksRemaining` sixteenths per tick

**Confidence.** **High** (writers enumerated by `disp:6c` and read in place; the
arithmetic is immediates in two functions read end to end)

### ANIM-DEATH-007

The corpse stage is written by exactly one non-zeroing instruction image-wide
(`L02541`, the state sync under mask bit 3 — `REG-UNITS-050`'s own census, now
with its source), and the value is the server's `actor+0x13c`, the decay stage
`HERO-DEATH-026` reads: 1 at death, 2/3/4 as health passes **−10 / −20 / −40**,
5 below −600. The client saves the previous value first (`L02542`) and
dispatches on the **transition**, through a six-arm table at `L02543`: **0 →
1** starts action 6 for `2 × classes[Dying].DyingPhases` ticks with the clock at
0 — the fall, one dying frame held for exactly two ticks; **1 → 1** re-runs
action 6 for **4** ticks with the clock preset to `2·DyingPhases − 4`
(`L02544`, `L02545`), replaying the last two frames; **0 → 0** clears the
action. When the run's counter reaches 0 no arm runs and the draw's state-0
corpse fork takes over on the stage alone (`TERR-SPR-047`): 1 freezes on dying
frame `DyingPhases − 1`, 2/3/4 draw bone frames 0/1/2, and at 5 the actor is
gone (`MOVE-ID-016`). **The thresholds never reach the client as numbers** — it
sees only the stage — and the four `movementType > 1` classes (`Ghost`, `Bee`,
`Bat_Sonic`, `Dragon`) take `actor+0x94 = -1000` the moment `dyingTime` expires
(`L02546`), so they pass −10 without pausing and never occupy stages 2..4:
they play the fall and vanish, leaving no corpse and no bones. Corpus:
`BonePhases ≥ 3` on 261 of 262 pairs; the exception is `Death Star` (ID 72,
`Daemon`), whose `BonePhases` is the loader's **absent −1**, which the corpse
arm's guard (`L02547` compares the 32-bit field `+0x24` of the drawable with zero and skips on equal) does not catch — its
bone frame is `BaseBone − dir + stage − 2`, an index inside the dying block that
walks backwards with the facing

**Confidence.** **High** (every transition, immediate and threshold is a named
instruction; the writer census is `disp:15a`, complete, and its blind spot — a
store through a computed pointer — is `REG-UNITS-050`'s own and is restated
here) / **Medium** (the `Daemon` consequence: the arithmetic is read, whether
the class is placed on a shipped map was not re-measured)

**Amended.** The arm for a current stage of 0 acts only on a non-zero saved
stage: it clears the action and, on a hero-shaped drawable, calls
`R0551`; **0 → 0** does nothing (`claims/retracted.md`). Both stage-1 arms also fire the hurt hook, 3 on
**0 → 1** and 2 on **1 → 1**, and the switch runs on every state sync that
reaches it, whether or not its mask carries bit 3 (`ANIM-095`).

## Objects, idle and the animation tick

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-OBJ-008 | Objects animate off the same clock and by a different mechanism, and there is no second driver. | High | ● active (amended, partially retracted) | [EXP-0071](../experiments/EXP-0071-animation-driver/), **[EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/)** |
| ANIM-IDLE-009 | The fidget is draw state 0 forking on `IdlePhases`, and it runs on the drawable's own clock — one animation tick per timeline step. | High / Medium | ● active | [EXP-0084](../experiments/EXP-0084-idle-animation/) |
| ANIM-IDLE-010 | Five of the 34 `units.reg` classes carry idle art; the other 29 stand on a single frame per facing, and the cadence is data. | High / Medium | ● active | [EXP-0084](../experiments/EXP-0084-idle-animation/) |
| ANIM-TICK-011 | The animation tick has a second job — and it is the writer of tile bits 15..14 that `TERR-TILE-044` looked for and could not find. | High / Medium / Unknown | ● active | [EXP-0084](../experiments/EXP-0084-idle-animation/) |
| ANIM-VT-012 | `CAirUnit` shares the driver and both draw passes verbatim; `CProjectile` shares none of them. | High / Unknown | ● active | [EXP-0084](../experiments/EXP-0084-idle-animation/) |

### ANIM-OBJ-008

`R0334` does both jobs in one pass of the `0x401` tick: it increments
`CMapView+0xa70` (`L02475..L02476`), the counter `TERR-SPR-042`'s animated
object arm and `TERR-SEM-004`'s water phase both read, **and** calls `vt+0x3c`
on every drawable in the container. But a placed object carries no per-instance
animation state at all — no phase, no clock, no action code: its frame is
`Index + T[(animCtr + col·(row+1)) mod class+0x3c]` recomputed at every draw
(`TERR-SPR-042`), where `T` is the built `AnimationTime`/`AnimationFrame`
expansion (`REG-OBJ-046`). So the two systems share a clock and nothing else,
and the object arm's gate is the four-corner `0xc000` test that 0 of 880 704
shipped **cells** satisfy (`TERR-TILE-044`). **⚠ The clause this row printed
next — that the arm is therefore “unreachable anyway” on a shipped map — is
withdrawn by [EXP-0084]**: the census is of map *data*, and `ANIM-TICK-011`
finds the runtime writer, `vt+0x48` ORing `0xc0` into the high byte of exactly
those four corners around every drawable, with `R0374` clearing bit 14
map-wide every 32 ticks. Whether the gate is satisfied in play is now
**unestablished** — the two conditions guarding the stamp are unread — rather
than answered either way. **⚠ RESOLVED by [EXP-0085]: it IS satisfied, and the
arm fires.** The guards are three, all read (`TERR-TILE-079`), and the stamp's
immediate is `0xc0` — both gate bits into one word — so a single stamped cell
passes the four-corner OR. Bits 15..14 are the **fog of war**, and the
animated-object arm therefore runs on **exactly the cells the local player can
currently see**: a shipped map's fires and trees animate inside the field of
view and hold frame 0 outside it. A consumer needs one tick source, one per-unit
state block, and for objects a pure function of `(animCtr, cell)` **plus the fog
state of the cell's own quad**

**Confidence.** **High** (the shared increment and the dispatch are in one
function read end to end; the object side is `TERR-SPR-042`/`REG-OBJ-046`, cited
not re-derived) / ~~**Unknown** (the reachability of the object arm)~~ — **High
by [EXP-0085]**, on named instructions and two enumerations that state their
instruments

**Amended.** The clause that the object arm is unreachable on a shipped map is
withdrawn (`claims/retracted.md`), and EXP-0085 resolves the arm's reachability:
the gate bits are the fog of war, so the arm runs on the cells the local player
can see. The two ⚠ notes in the body carry both changes.

### ANIM-IDLE-009

When the action counter `+0xa0` is 0 no arm of `ANIM-STATE-002`'s 8-entry table
runs (`L02548`, `L02549`), the draw state is forced to 0 (`L02484`), and
the class's `IdlePhases` (`class+0x28`, `REG-UNITS-049`) decides everything
after that: **nonzero** and the clock `+0x94` is incremented **once per tick**
(`L02550`, `L02551`) with `phase = clock mod class+0x80`, the length of the
*built* Idle timeline; **zero** and phase and clock are both pinned to 0
(`L02506`, `L02507`). Both draw passes read the same fork before anything
else and compute the same index —
`idleBase + dir·IdlePhases + idleTL[unit+0x70]`, with **no** modulo — body
`R0552` arm 0 (`L02552`, `L02553`, `L02554`, `L02555`/`L02556`)
and shadow `R0553` arm 0 (`L02557`, `L02558`, `L02559`,
`L02560`/`L02561`), the base being `SPR256-UNIT-024`'s `S + D·(MB+MV+AT+DY)`
— read here for the **body** pass for the first time. Three things a consumer
must carry: **(a) idle is not a state of its own** — a fidgeting bee and a
motionless swordsman are both in state 0 and the difference is one registry key,
so `TERR-SPR-047`'s state list needs no tenth entry; **(b) the corpse fork is
unreachable for a class with idle art**, because this arm forks on `IdlePhases`
before it reads the corpse stage `+0x15a` — which is *why* `SPR256-UNIT-024`
found the idle and bone blocks sharing a base with no class carrying both: the
collision is unreachable, not merely absent; **(c) the idle clock is the action
clock** — nothing on the path from an action ending to this arm stores 0 into
`+0x94`, every reset being at an action's *start* in the message dispatcher
(`L02562`, `L02563`, `L02564`, `L02565`, `L02566`, `L02567` := 0;
`L02568`, `L02569`, `L02570` := −1), so a unit enters its idle loop at
whatever phase its last action left it and two units of one class idle **out of
phase** unless neither has ever acted. The drawable also raises its repaint flag
`+0x10c` on every tick while `class+0x28 > 0` and not otherwise (`L02571`,
`L02572`)

**Confidence.** **High** (the fork, the increment, the modulus and both draw
arms are named instructions in functions read end to end, and the computation
appears **twice independently** — body and shadow, instruction for instruction;
the driver's referrer set is `callto:R0548` = 2 hits, 0 owners, 0 orphan, both
of them RDATA `vt+0x3c` slots, so nothing else in the image calls it) /
**Medium** (that no *other* per-tick route varies an idle unit's frame: the two
per-tick virtuals the driver itself calls, `vt+0x44` at `L02573`/`L02574`
and `vt+0x50` at `L02575`, were read and touch no animation field, but a
`callto:` sweep cannot enumerate an indirect call, so this is a reading of one
call chain rather than an enumeration)

### ANIM-IDLE-010

Re-executing the loader's own run-length expansion over all 34 classes:
`Sonic Bat` (ID 70) 6 frames / 12-step loop, `Bee` (73) 4 / 4, `Dragon` (71) 7 /
14, `Ghost` (69) **3 frames over a 4-step ping-pong** `0 1 2 1` / 12,
`Death Star` (72) 7 / 14 — the loop length is the *expansion's* length, not
`IdlePhases`, which is why the two disagree on `Ghost` exactly as
`REG-UNITS-018`'s withdrawn clause predicted from the other side. One step is
one animation tick, so at the shipped speed index (16 tps, 62.5 ms) the cycles
are 750 / 250 / 875 / 750 / 875 ms and they scale over the whole nine-arm ladder
— `Bee` runs 500 ms at 8 tps and 125 ms at 32. Every `(direction, phase)` index
lands inside its own direction slot **5/5** and inside the sheet **5/5**, and in
all five the highest idle index is the sheet's **last** frame (108/109, 88/89,
163/164, 98/99, 153/154), which closes the block layout at the top end. Corpus
over the 38 shipped maps: **1740 of 8094** type-6 placements (21.5 %) are of a
class that can fidget, on **34 of 38** maps — `Bee` 608, `Sonic Bat` 580,
`Ghost` 416, `Dragon` 135, `Death Star` **1** (`scn:150.alm`, the last campaign
mission). Four of the five are exactly the four `movementType > 1` classes of
`MOVE-DOM-028`; that is a coincidence of the shipped data, **not** a rule
anything in the image states

**Confidence.** **Medium** (a corpus measurement over one registry and 38 maps,
which caps at Medium by `METHODOLOGY.md`; the arithmetic it re-executes is
`ANIM-IDLE-009`'s and is High) / **High** (the in-slot and in-sheet closure: it
is an exhaustive re-execution of the draw's own index expression over every
direction and every phase of every class, against each sheet's own frame count)

### ANIM-TICK-011

`R0374`, the routine this ledger listed as located and unread, runs once
every 32 counts of the same `0x401` tick (`L02576`..`L01627`, gated
`[L01661] == 0`): it sets `CMapView+0xdc`/`+0xe0` to 1, **clears bit 14 of
every tile word of the map** — `L01659` masks each word with 0xbfff over
`[mapView+0x84] × [mapView+0x88]` words of `[[mapView+0x80]+0xc]`, the map's own
grid — and then walks the same drawable container the driver walks, calling
`vt+0x48` on each (`L02577`). It writes **no** animation field. `vt+0x48`
(`R0554` → `L02578`, also reached once per tick per drawable through
`vt+0x44`'s position-changed test at `L02579`) is the stamp: gated on
`[[mapView+0x9b4]+0x38][player·2] & 8` and `unit+0x102 != 0`, it walks a 41×41
window centred on the drawable, masked by a 41×41 dword table at
`CMapView+0x17cc` (`L02580`), and for each admitted cell executes **four**
bit-set operations (0xc0 into the high byte of the tile word) at `idx`, `idx+1`, `idx+W+1`, `idx+W`
(`L02581`, `L02582`, `L02583`, `L02584`) — **exactly the four corners
`TERR-TILE-044`'s gate ORs together**. The immediate is `0xc0` on the word's
**high byte**, which is why that row's sweep for the 16-bit immediates
`0xc000`/`0x4000`/`0x8000` returned nothing: the failure was the operand width,
not a missing instruction. The two halves also explain that row's own
three-state finding — bit 15 is set by the stamp and cleared by nothing
(`imm:bfff` = 2 hits image-wide, one of them this clear, the other an unrelated
dword mask), while bit 14 is cleared every 32 ticks and re-set by the next
stamp, so `== 0x8000` and `== 0xc000` distinguish *stamped once* from *stamped
this cycle*

**Confidence.** **High** (the clear, the walk, the four ORs and the two gates
are named instructions in three functions read end to end, and the four
displacements are the gate's four corners — an agreement of shape, not of
plausibility) / **Medium** (that this is *the* writer: it was found by reading
the tick's own call chain, and an `imm:` sweep for byte-wide ORs of `0xc0` is
far too common to enumerate, so a second writer elsewhere is not excluded) /
**Unknown** (what the `CMapView+0x17cc` mask holds, what `unit+0x102` and the
`&8` player flag are, and therefore whether the four-corner gate is actually
satisfied in a running game — what falls is `ANIM-OBJ-008`'s and
`TERR-TILE-044`'s *“the arm is unreachable”*, which was inferred from a census
of map **data**; what replaces it is not “it fires” but “unestablished”, and it
needs its own experiment)

### ANIM-VT-012

The three drawable vtables, dumped whole: `CUnit` `L02468`, `CAirUnit`
`L02585` (the immediate its constructor stores, `L02586`), `CProjectile`
`L02587` (the only referrer of its `GetRuntimeClass` stub `R0555`).
`CAirUnit` differs from `CUnit` in **four** of 33 slots — `+0x00`
`GetRuntimeClass`, `+0x04` the destructor, `+0x10`, and `+0x38` the draw-layer
registration `REG-UNITS-061` already named — and slots `+0x28` body, `+0x2c`
shadow and `+0x3c` driver are the *same functions*, so a flier is animated by
the routine this ledger describes and by no other. `CProjectile` overrides all
three: body `R0556`, shadow `R0557`, driver `R0558`, which
is why `callto:R0548` finds exactly two RDATA slots and no third. A consumer
therefore needs one animation model for units and fliers and a **separate** one
for projectiles

**Confidence.** **High** (three vtable dumps on the repaired function table, 0
of 33 slots without a function in each, and each base established from an
instruction or an RTTI referrer rather than from where a run of pointers happens
to start) / **Unknown** (what `R0558` does — the projectile's own driver
is located and unread)

## Walk cadence and the pacer

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-WALK-013 | The walk phase is advanced by an ODOMETER, and it is the only animation in the engine not advanced by the tick. | High / Medium | ● active | [EXP-0095](../experiments/EXP-0095-walk-cadence/) |
| ANIM-WALK-014 | One straight cell crossing advances the Move timeline by exactly 16 steps — for every unit, at every speed, and the engine fixes it rather than the data. | High / Medium | ● active | [EXP-0095](../experiments/EXP-0095-walk-cadence/) |
| ANIM-WALK-015 | How many walk CYCLES those 16 steps are is data — and the shipped data puts exactly one cycle in one cell on 208 of 262 pairs and leaves 44 unaligned. | High / Medium | ● active | [EXP-0095](../experiments/EXP-0095-walk-cadence/) |
| ANIM-AMBIENT-016 | The world's ambient animation is a THIRD counter on the same pacer, and it coincides with the walk at exactly one movement rate. | High / Medium | ● active | [EXP-0095](../experiments/EXP-0095-walk-cadence/) |
| ANIM-PACE-017 | The game-speed index moves every clock together and moves no ratio — and the pacer is a catch-up loop, not a frame scaler. | High / Medium | ● active | [EXP-0095](../experiments/EXP-0095-walk-cadence/) |
| ANIM-ARM-018 | There is a SECOND walk-advance arm, on a different pair of odometers, and no claim had noticed it. | High / Unknown | ● active | [EXP-0095](../experiments/EXP-0095-walk-cadence/) |

### ANIM-WALK-013

`R0548`'s move arm splits the remaining delta by the ticks remaining
(`L02492`/`L02588` signed divide by the ticks remaining, `unit+0xa0` loaded at `L02548` before
the tail's decrement at `L02481`), applies it to the `/256` position, and then
adds the **Euclidean length of that tick's own displacement** to the clock
`unit+0x94` — squares each per-axis step (`L02589`, `L02590`), sums and
roots them in floating point (`L02495..L02591`), adds the 32-bit clock
`+0x94` (`L02592`), truncates toward zero through the CRT float-to-int
helper R0279 (`L02496`) and stores the result back to `+0x94`
(`L02593`) — then shifts it right by 4 (arithmetic) and stores the result as
the phase at `+0x70` (`L02497`/`L02594`). So `+0x94` is
**distance travelled in 1/256 of a cell**, truncated once per tick, and one
Move-timeline step is **16/256 = 1/16 of a cell**. Every other timeline in the
image is `+1` per tick at the same field: idle `L02550` (an increment), attack
`L02595`, shoot `L02500`, cast/die `L02501` — and a placed object's is
`+1` per tick on a different counter (`ANIM-AMBIENT-016`). The arm also
maintains two **per-axis** odometers, `+0x98 += |sx|` (`L02596`/`L02597`)
and `+0x9c += |sy|` (`L02598`/`L02599`), which the phase does not use
(`ANIM-ARM-018`). **The clock is not reset when a walk starts**: over
`EnumRefs disp:94`, every writer inside the message dispatcher — `L02562`,
`L02600`, `L02563`, `L02564`, `L02565`, `L02568`, `L02566`,
`L02569`, `L02567`, `L02570` — is in the death, turn, attack, shoot or
`0x86`/`0x8a`/`0x8b`/`0x8c` arm, and **none** lies in `[L02516, L02517)`,
the `0x6b` move arm; so consecutive cells continue one odometer. What does reset
it is standing still for one tick with `IdlePhases == 0`, which stores 0 into
both `+0x70` and `+0x94` (`L02506`, `L02507`), so for the 29 classes without
idle art the walk restarts from frame 0 after any pause

**Confidence.** **High** (the arm is transcribed from the raw listing operand by
operand — the rival "a per-tick counter" is excluded by a *present* `FSQRT`, not
by plausibility, and the four tick-driven arms are quoted beside it for
contrast; the reset census is `disp:94`, whose hits are read by address rather
than by function name) / **Medium** (that nothing else writes `+0x94`: a `disp:`
sweep cannot see a wholesale copy of the drawable, and `ANIM-CLOCK-001` already
bounds the same blind spot)

### ANIM-WALK-014

The `0x6b` arm loads the remaining delta from the two per-direction tables at
`L02525`/`L02526` shifted left 8 (`ANIM-MSG-005`), so a straight cell is
`±256` on one axis; the per-tick shares sum to exactly 256 by construction (each
tick takes `rx / ticksRemaining` and subtracts it), and for a straight step
`sqrt(sx² + 0)` is the integer `|sx|`, so **no truncation occurs** and the clock
lands on 256 and the phase on `256>>4 = 16`. Re-executed by `tools/walkcadence`
over all **262** (class, definition) pairs at every shipped `Speed`: **262/262
advance exactly 16**, histogram `16:262`. A **diagonal** cell is `±256` on
*both* axes — `256√2 = 362` in the metric the clock accumulates — but the
per-tick `ftol` truncation removes up to 1 per tick, giving `20:6 21:126 22:130`
over the same pairs. Two consequences a consumer must carry: the frame count per
cell is a property of the **grid**, not of the unit, so a slow unit holds each
frame longer and never skips one; and the walk is the one animation whose rate a
consumer cannot get right by counting ticks

**Confidence.** **High** (an exhaustive re-execution of a transcribed arm over
the whole shipped input domain, not a sample, and the straight case additionally
closes by an arithmetic identity that a rounding error would break — the rival
"the count varies with the unit" scores 0 of 262) / **Medium** (the diagonal
histogram and the 262-pair join: one root, and the `units.reg`↔`Data.bin`
resolution is EXP-0071's rule, cited not re-derived)

### ANIM-WALK-015

The modulus is `class+0x50`, the element count of the
`MoveAnimTime`/`MoveAnimFrame` run-length expansion: `R0395` reads the
two keys at `L02601`/`L02602` (`StrDump fnstr:R0395`), expands them at
`L02603..L02604`, and stores the array's own count at
`class+0x50` (`L02605`); the body draw applies it as
`phase mod class+0x50` (`L02606` Flip arm, `L02607` non-Flip) and indexes
the array at `class+0x40` (`L02608`/`L02609`). Re-expanding the loader's
loop over all 34 `units.reg` classes, `class+0x50` per pair is
`4:6 8:4 10:4 12:8 14:20 16:208 20:8 24:4`. The standard shape is
**`MovePhases = 8` art frames each held two timeline steps**, `L = 16`, one
cycle per cell, on 23 of the 34 classes. The exceptions are exact: `Bee` and
both `Catapult`s run **four** cycles per cell (`L=4`), `Ghost` two (`L=8`), and
the **44 pairs with `L ∈ {10,12,14,20,24}`** — `Squirrel`, `Snake`,
`Ogre (Lord)`, `Turtle`, both heroes, both `Goblin`s, `Sonic Bat` — have
`16 mod L != 0`, so their walk cycle is **not** cell-aligned and its phase at a
cell boundary depends on how far the unit has walked since its last stop.
Nothing in the code requires alignment: it is an authoring convention
`units.reg` mostly keeps

**Confidence.** **High** (the modulus, the array and the builder are named
instructions, and the key binding is a `StrDump` of the loader's own `PUSH`
operands, so the field↔key pairing is read rather than assumed) / **Medium**
(the histogram: a corpus measurement over one registry on one root, which caps
here by rule)

### ANIM-AMBIENT-016

It is neither the walk clock (`unit+0x94`, per drawable) nor the displacement
counter (`mover+0xac`, per mover, simulation side): it is `CMapView+0xa70`,
incremented once per `0x401` tick at `L02475`/`L02610`/`L02476`.
`EnumRefs disp:a70` is **21 hits / 13 owners / 0 orphan**, with exactly one
non-increment writer (`L02611 := 0`, the map view's construction) and no owner
in the drawable module. A placed object reads it **unshifted** —
`frame = Index + T[(animCtr + col + row·col) mod class+0x3c]`
(`L02612..L02613`, `TERR-SPR-042`) — so an object's timeline advances **one
step per tick**, the same cadence as idle, attack and death. The *repaint* is
throttled separately by a latch: the elapsed count is the 32-bit field `+0xa70` less the latch `+0xa74` (`L02614`/`L02615`), `L02616` tests it signed-above 3, and `L02617` writes the latch `+0xa74` — one ambient frame per **4** ticks, and
`disp:a74` is **3 hits / 2 owners / 0 orphan** (the constructor's zero plus
those two), so nothing else in the image touches it. **The arithmetic that
answers the question:** the walk advances `16 / ceil(256/v)` timeline steps per
tick and everything else advances 1, so they are in step **iff
`ceil(256/v) == 16`, i.e. `v ∈ {16,17}`** — **78 of 262** shipped pairs (29.8
%); over the remaining 184 the walk runs between **0.5×** (`Speed 8`, 32 ticks
per cell) and **2.0×** (`Speed 35`, 8 ticks) the universal cadence.
**Falsifiable prediction:** at the map-load speed a `Human Swordsman` at
`Speed 19` on cost-8 ground crosses a cell in **14 ticks ≈ 868 ms** playing
exactly **one** 8-frame walk cycle, so its feet cycle 1.14× faster than a fire
beside it; a `Fat troll` at `Speed 8` takes 32 ticks and cycles at half the
fire's rate. A stopwatch on either refutes this row

**Confidence.** **High** (the increment, the object arm's shift and the latch
are named instructions in three functions read end to end, and both `disp:`
enumerations report 0 orphan; the rival "one shared counter" is excluded by two
disjoint owner sets rather than by argument) / **Medium** (the 78/262 and the
two example durations: a corpus join on one root at `meanCost = 8`, and no
runtime session has been observed)

### ANIM-PACE-017

`R0559` clamps the index to `[0,8]` (`L02618`, `L02619`), selects
`tps ∈ {8,10,12,14,16,20,24,28,32}` through the jump table at `L02620`, and
computes `dtMs` with a **truncating** divide — `L02621`/`L02622` divide 0x3e8 by the selected tps (signed) — storing it at `campaign+0x3f0` (`L02623`) while zeroing
the window index `+0x3e4` and rebasing `+0x3ec`. `EnumRefs disp:3f0` is **7 hits
/ 5 owners / 0 orphan** and `L02623` is its only writer, so nothing else
alters the pace. `R0454` returns unless
`base + (step+1)·dtMs <= timeGetTime()` (`L02624`/`L02625`), and after
running one tick body — the simulation sub-tick `L02626`, then the `0x401`
presentation tick `L02627`/`L02628`, in that order, in one iteration — it
re-tests and **repeats the whole body** while still overdue (`L02629`,
`L02630 JBE`). The step index wraps `& 0xf` (`L02631`) and `base` is re-read
from the wall clock whenever it is 0 (`L02632`), so lateness is discarded
every **16** ticks: a slow machine loses time permanently and a fast one cannot
run ahead. Realised `dtMs` at the nine indices is
`125 100 83 71 62 50 41 35 31`, and only three of the nine divide 1000 — the
map-load default is **62 ms, i.e. 16.13 ticks/s, not 16**. What the index does
**not** change: ticks per cell (`ceil(256/v)`), walk steps per cell (16), object
steps per tick (1) — so the walk:ambient ratio is the same at all nine speeds
and **no speed setting can bring them into step**. **`campaign+0x3e4` animates
nothing**: `disp:3e4` is **20 hits / 9 owners / 0 orphan** and every owner is a
pacer, the speed setter, or the in-game `+`/`-` handler — no renderer reads it

**Confidence.** **High** (the clamp, the ladder, the truncating divide, the
deadline test, the repeat branch, the wrap and the rebase are named instructions
in two functions read end to end, and the three `disp:` enumerations each report
0 orphan; the "no renderer reads `+0x3e4`" is an absence over a complete owner
list, not over a name search) / **Medium** (that `R0454` is the arm a
normal single-player session runs: that is `TERR-ANIM-008`'s selection through
`R0560`, cited not re-derived, and `R0455` is a second pacer of
the same shape that was not read)

### ANIM-ARM-018

The branch at `L02633`/`L02634` to `L02635`, taken when `unit+0x20` is negative — the
class id the driver itself uses as `classes[+0x20]` at `L02636`/`L02637` —
diverts the move arm's tail into roughly 600 bytes that never touch `+0x94`.
They read the **per-axis** odometers instead: `+0x98` divided by 32
(`L02638` shifts right by 5) and by 25 (`L02639` magic `0x51eb851f`,
`L02640` shifts right by 3), `+0x9c` by 26 (`L02641` magic `0x4ec4ec4f`,
`L02642` shifts right by 3), each masked to **8 states**; the state is latched in
`unit+0x30` and re-entered only on change (an equality test against `+0x30` jumps to `L02643`
in every arm), it is keyed on the facing through a 13-entry byte table at
`L02644` into the jump table at `L02645` with facings above 12 falling
to `L02646`, and each state writes a sub-cell draw offset of ±8/256 into
`+0x28`/`+0x2c` (`L02647`..`L02648`, `L02649` forming a base plus doubled index,
`L02650`/`L02651`). This is what the two per-axis odometers of
`ANIM-WALK-013` exist for; the main arm computes them and never reads them.
**Reachability is open and the row does not claim it either way.** Two things
point at unreachable: no instruction image-wide stores a negative immediate into
a `+0x20` displacement (`EnumRefs`, a regular-expression search for immediate stores into that displacement — every
hit on a drawable stores a small non-negative class id, `1..0xb`, `5`, `0x24`),
and **both** draw passes dereference `classes[unit+0x20]` unconditionally before
any sign test (`L02652` loads the class pointer from the class table at the id, then
`L02653` reads the 32-bit field `+0x14` of that class), so a drawable reaching this arm would already
have faulted in its own draw on the same tick

**Confidence.** **High** (the divisors, the magics, the state latch, both tables
and the offsets are transcribed from a full listing of the arm, and the two
odometers' writers are the main arm's own instructions) / **Unknown** (whether
any drawable ever has `unit+0x20 < 0`. Two things would settle it: a classified
`disp:20` write census over the drawable module — the immediate sweep above
cannot see a store from a register, which is how a `Data.bin` absent `−1` would
arrive — or a runtime witness. Until then the arm is located and read, not
attributed)

## Blow display and unit sounds

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-BLOW-019 | When a blow lands the original draws a NUMERAL and plays a sound, and those two are the whole of it — and the engine names the event itself. | High / Medium | ● active (amended, partially retracted) | [EXP-0111](../experiments/EXP-0111-blow-display/), **[EXP-0323](../experiments/EXP-0323-damage-numeral-pacing/)**, **[EXP-0405](../experiments/EXP-0405-hurt-voice-bank/)** |
| ANIM-NUM-020 | The numeral is a 0x24-byte object with a 1000 ms life, drawn in the VICTIM OWNER's colour, and its whole existence is four routines with one caller each. | High / Medium | ● active (amended) | [EXP-0111](../experiments/EXP-0111-blow-display/), **[EXP-0323](../experiments/EXP-0323-damage-numeral-pacing/)** |
| ANIM-NUM-021 | A severity colour is chosen on every landed blow and read by nothing — a live computation with no consumer, which is neither a dead arm nor a dead field. | High / Medium | ● active | [EXP-0111](../experiments/EXP-0111-blow-display/) |
| ANIM-SND-022 | Two hooks, not one — the attacker's swing and the victim's hurt cue — both positional DirectSound, but the hurt hook can switch from the five-element `Sound` array to a filename-composed voice bank. | High / Medium | ● active (amended, partially retracted) | [EXP-0111](../experiments/EXP-0111-blow-display/), **[EXP-0230](../experiments/EXP-0230-non-music-sfx-events/)**, **[EXP-0405](../experiments/EXP-0405-hurt-voice-bank/)**, **[EXP-0474](../experiments/EXP-0474-sfx-spatial-inputs/)** |
| ANIM-STATE-023 | Which arm fires which hook: `ANIM-PHASE-003`'s four hooks are not two arms firing two each, they are three arms firing one or two apiece — and the delay keys are shared, not owned. | High | ● active | [EXP-0111](../experiments/EXP-0111-blow-display/) |
| ANIM-CLOCK-024 | The swing's sound and the swing's damage are timed by two different files, and on the shipped data the sound comes FIRST on 153 of the 157 pairs that answer. | High / Medium | ● active | [EXP-0111](../experiments/EXP-0111-blow-display/) |
| ANIM-094 | The hurt hook's voice-bank field is fixed by the event index alone — k=1 `+0x24` easy.wav, k=2 `+0x28` hard.wav, k=3 `+0x2c` die.wav — whenever `unit+0x18c & 0x11` is non-zero, and k=0 always plays class `Sound[1]`. | High | ✔ promoted | [EXP-0405](../experiments/EXP-0405-hurt-voice-bank/) |
| ANIM-095 | Five `vt+0x68` calls in the client dispatcher emit the hurt event: the `0x73` arm passes 0, 1 or 2 from the drawable's health before the message, and the state-sync stage switch passes 3 on stage 0→1 and 2 on 1→1. | High / Medium / Unknown | ✔ promoted (amended) | [EXP-0405](../experiments/EXP-0405-hurt-voice-bank/) |
| ANIM-096 | For a hero-shaped drawable each hurt branch has one source: the bank follows `+0x18c` bits 0x2 (mage) and 0x4 (sex), the field follows k, and equipment and the drawn class change only the no-damage cue. | High / Medium / Unknown | ✔ promoted (amended) | [EXP-0405](../experiments/EXP-0405-hurt-voice-bank/) |
| ANIM-119 | The five voice readers `vt+0x6c`..`+0x7c` each read one bank slot, except `+0x6c` (`command1..3` or `defend`) and `+0x78` (`select1` or `select2`), which draw once from `rand()`; they write only the stamp `unit+0x190`. | High / Medium / Unknown | ● active | [EXP-0445](../experiments/EXP-0445-command-voices/) |
| ANIM-120 | The default drawable constructor sets the voice stamp `+0x190` to 0, so a first command or selection reply is admitted once `timeGetTime` reaches 3000 ms (2000 ms for the selection reply); the copy constructor carries the source's stamp. | High / Medium / Unknown | ● active | [EXP-0445](../experiments/EXP-0445-command-voices/) |

### ANIM-BLOW-019

The strike `R0246`, read end to end, ends in exactly two calls: the
experience payment `HERO-KILL-027` owns, and `R0561(target, dmg)` at
`L02654`, gated on the target having been alive before the blow or still above
-10. That builder is twelve instructions read whole and sends **opcode `0x73`**
(`L02655` supplies 0x73 as the opcode argument) through `R0562`, and **it never reads the
damage**: `dmg` is pushed at the call site (`L02656`) and its 8-byte callee-popped
frame's second argument is touched by no instruction. What the message carries
is the victim's **new health as a level** — the client stores `msg+0xc` into the
drawable's health at `L02657` — and the **client subtracts** to get the delta,
`L02658` / `L02659` / `L02660`. The client arm is `L02661`, resolved
from the dispatcher's own two tables by RVA with no disassembler in the path
(`evidence/tables.md`), 231 instructions with **0 bytes in range never
disassembled**, and the engine's own name for the opcode sits in that arm's
failure path: `L02662` = `Invalid unit #%d. Command Take damage.` The arm
does five things and stops. It resolves the victim through `view+0x9b8`
(`AI-SELECT-065`), logging and returning when the id is unknown. If `view+0xaa4`
is set **and** the drawable's health exceeds `msg+0xc`, it builds the floating
numeral (`ANIM-NUM-020`). It calls `vt+0x68` with 0, 1 or 2 by remaining-health
band (`ANIM-SND-022`). It stores the new health. And if the victim is the map
view's `+0x138` it posts `0x408` to the two panels at `view+0xe0` / `+0xe4`,
then raises the drawable's repaint flag `+0x10c`. **There is no palette swap, no
overlay frame, no flash and no second sprite** — and the field-masked state sync
of `ANIM-MSG-005` cannot supply one either, being 616 instructions of stores
with no call among them. `EnumRefs callto:R0561` is **5 hits / 4 owners / 0
orphan**: this strike, the structure strike `R0563` (whose target's
health is a `u16` at `+0x42`, not `+0x94`), `R0564` twice and
`R0565`

**Confidence.** **High** for what the strike emits and what the arm does — both
read end to end, every step a named instruction, and the opcode-to-arm binding
decoded from static data / **Medium** for *these two are the only things a blow
produces*: the strike side is an enumeration over one routine read whole, but
the three non-strike producers of `0x73` were not read, so a spell arriving at
the same actor is covered by the arm half and not by the producer half

**Amended.** The fifth owner's own gate is read by `ANIM-074`; the producer
population, the melee/missile/structure synthesis and the
wall_of_fire/poison_cloud cross-reference are given by `ANIM-075`. The band
operand is the health before the message, not the remaining health, and the
state sync is not call-free: its arm holds 22 calls and its stage switch fires
the hurt hook with 3 and 2 (`ANIM-095`, `claims/retracted.md`). Whether any of
those calls draws was not read.

### ANIM-NUM-020

`R0566` constructs it on the arm's own stack (`L02663`) and
`R0567(this, format, dmg)` formats the string at `rec+0x00` — the format
literal is `L02664` = `%d`, so what is drawn is a decimal numeral and
nothing else. `rec+0x04` keeps the number, `rec+0x08` the colour table
`L02665 + (Player+0x8 << 5)` taken from the **victim's** owning `Player`
(`L02666`..`L02667`), `rec+0x0c` / `+0x10` the offset from the victim's own
screen position, `rec+0x14` `timeGetTime()` at birth
(`L02668`, an indirect call through the import slot `L00849`), `rec+0x18` the not-mine flag and
`rec+0x1c` the victim drawable. `R0568` appends it to the map view's list
at `view+0x3f3c` at stride `0x24` and **merges rather than duplicates**: a
record already present with the same `+0x1c` and `+0x20` takes the new damage
**added into its own `+0x04`** (`L02669`..`L02670`). `R0569`, called
once per `0x401` animation tick from `R0334` (`L02671`), steps every
record — `+0x10 -= 2` unconditionally and `+0x0c` by one, **away** from the
local player's own units and toward everyone else's (`L02672`..`L02673`).
`R0570`, called once per frame from the map view's paint `R0379`
(`L02674`), draws and expires: `L02675`..`L02676` take `timeGetTime` less the birth time `rec+0x14` and test the difference against 0x3e8 (unsigned) — past **1000
ms of wall clock** the record returns 0 and is destructed and the array
compacted with a repeated 32-bit string move. Under it, and unless the victim drawable's `+0x78`
is non-zero, it calls `R0571` on the global at `L02677` at
`(rec+0x0c + owner+0x60, rec+0x10 + owner+0x64 - owner+0x68)`, which issues the
same text virtual **twice** at offset positions — a shadow and a face. The
initial offset is the victim's own `vt+0x20()` scaled: 16 horizontally with the
sign taken from ownership, `-0x30` vertically. **The display is optional and
defaults to on**: `view+0xaa4` is set to 1 by the map view's constructor
(`L02678`) and flipped by the arm at `L02679`, which requires the Ctrl
latch `[L00627]` (`AI-KEYMOD-059`) and is reached from vkey `0x4c` alone —
**Ctrl+L** — solved by scanning the handler's own byte table at `L02680` for
every vkey whose arm is that address, 1 of 161 (`evidence/tables.md`). **So the
life is wall-clock and the drift is tick-paced**, and how far a numeral drifts
before it vanishes scales with the tick count at each speed rather than a single
ratio this row names — `ANIM-071` separately narrows when two hits on one victim
merge to sub-tick granularity, and `ANIM-073` gives the per-speed tick counts
and scopes the drift ratio (1:2 against the default speed, 1:4 against the
fastest, 1:3 against 24 tps)

**Confidence.** **High** (all four routines read end to end; `callto:` on each
is **1 hit / 1 owner / 0 orphan**, so the chain has no second entry and no
second consumer; the lifetime, the drift, the stride, the merge and the key are
immediates or static bytes) / **Medium** (that the global at `L02677` is a
glyph blit: it has one writer and 135 reads over 22 owners and was not followed
into the font. What carries *drawn on the map* independently of that is the
position — an offset from the victim drawable's own coordinates — and the
drain's caller being the map view's paint rather than a text control)

**Amended.** `ANIM-071` names `+0x20` and shows the merge window this row's own
`+0x1c`/`+0x20` condition gates is bounded by sub-tick granularity, narrower
than the 1000 ms life stated two sentences later in this same row —
`claims/retracted.md`; the closing ratio is scoped by `ANIM-073` —
`claims/retracted.md`.

### ANIM-NUM-021

Between resolving the victim and constructing the numeral the `0x73` arm picks
one of three word arrays by the victim's **remaining** health:
`msg+0xc >= (unit+0x100)/2` takes `L02681` (`L02682`), else
`>= (unit+0x100)/4` takes `L02683` (`L02684`), else `L02685`
(`L02686`); the choice is stored in the stack local at frame displacement 0xfffffaa0. An
operand-text scan of the containing function's complete listing —
`R0509`, 3 500 instructions, **0 bytes in range never disassembled** —
returns **three writes of that slot and zero reads**, and no `LEA` of it either,
so nothing can alias it. The numeral is coloured instead by the owner's player
colour, which `R0566` takes unconditionally because the arm's second
argument is the literal 0 at `L02687` and that constructor has one caller. The
three targets are not strings: they are **word arrays**, written only by
`R0572` at `L02688`, `L02689` and `L02690`, each the store at the
end of a shift-and-OR of three channel fields, i.e. a colour packed into the
surface's own pixel format. **Consequence for a consumer:** rendering the
numeral in a severity colour adds something the original computes on every blow
and then discards, and the bands it would have used are `>= max/2`, `>= max/4`,
below

**Confidence.** **High** (the three-way test and the three writes are named
instructions; the absence of a read rests on a text scan of one function's
complete listing, an instrument that cannot miss a frame-local displacement, and
whose only blind spot — an address taken of the slot — the same scan excludes) /
**Medium** (that the three arrays are colour ramps: `R0572` was read at
the three stores and at its packing idiom, not decoded, so *colour* is a reading
of a shift-and-OR rather than a decoded table)

### ANIM-SND-022

The `units.reg` class record's `Sound` key builds a `CArray` at `class+0xbc`
whose data pointer is `class+0xc0` and element count `class+0xc4`; the loader
reads it at `L02691` and falls back to the parent section when the **count**
is 0 (`L02692`). **The swing** is `CUnit vt+0x64` = `R0573`, fired by
the driver's attack arm (`L02693`) and shoot arm (`L02694`) on the tick
`unit+0x94 == class+0x100`, `AttackDelay`: it reads `Sound[0]` (`L02695`),
does nothing when that element is 0, converts the drawable's map position into a
pan and an attenuation through `R0574` — clamped to 10000 either way at
`L02696` / `L02697`, computed with an `FSQRT` distance and a `log10`
(the log10 arithmetic label is retracted; VIDEO-SFX-088 establishes exp) — and
hands the pair to `R0386`. **The grunt** is `vt+0x68` = `R0575`,
which the animation driver never calls: the `0x73` arm does, with **0** when the
health did not change (`L02698`), **2** when the new health is below
`unit+0x100 / 2` and **1** otherwise, the latter two only while the new health
is above -10 (`L02699`). **Amended by
[EXP-0230](../experiments/EXP-0230-non-music-sfx-events/):** `k=0` plays
`Sound[1]`. For `k` in 1 or 2 it is **throttled per drawable to one in 1500
ms**; after that gate it plays `Sound[k+1]` only when `unit+0x18c` has neither
bit `0x01` nor bit `0x10`. Either bit instead selects the filename-composed
voice-bank receiver identified by `HERO-FIGURE-061`, and the two sources
converge at terminal call `L02700` (`VIDEO-SFX-018`):
`L02701` takes `timeGetTime` less the field `unit+0x190`; a difference below 0x5dc (unsigned)
returns without playing. `R0386` is the channel manager and it is
**DirectSound**: its first act is
`L02702`, a branch on the 32-bit global `L02703` being non-zero, and `[L02703]` is the
out-parameter `R0576` hands to `DirectSoundCreate` at `L02704` /
`L02705`, zeroed again when that call fails. `CAirUnit` shares both hook slots
with `CUnit`; `CProjectile` shares neither. Shipped, **byte-identical on both
roots**: 34 of 34 classes ship exactly **five** elements, and the element a hook
reads is the silent 0 on `Sound[0]` for **3** classes (both unarmed fighters and
the unarmed mage), `Sound[1]` for **10**, and `Sound[2]` / `Sound[3]` for **2**
(both `Catapult`s)

**Confidence.** **High** (both hooks and the three driver arms that fire them
read end to end; the key-to-field binding is a `StrDump` of the loader's own
`PUSH` operands; the element indices, the throttle and the clamps are
immediates; `callto:` on each hook is **2 RDATA slots / 0 code callers / 0
orphan**, so neither is reachable except through a vtable; and the DirectSound
identification is the *same global* the play routine gates on, read at the
`DirectSoundCreate` call site) / **Medium** (the census: `units.reg` re-executed
by the loader's own rule over 34 classes, which caps at Medium by
`METHODOLOGY.md` however exact)

**Amended.** EXP-0230 withdrew the universal registry-array source for the
victim grunt (`claims/retracted.md`); the correction is the inline EXP-0230 note
in the body. The operand of the `0x73` arm's band tests is the stored health
before the message, not the new health (`claims/retracted.md`, `ANIM-095`); the
state-sync stage switch also calls the grunt, with 3 and 2 (`ANIM-095`); and
`ANIM-094` binds each `k` to its bank field.

**Amended by EXP-0474.** Only the former log10 arithmetic label is refuted.
The reached original body uses exp after distance division by 256 and 8;
VIDEO-SFX-088 states the attenuation and pan separately, including clamp and
conversion placement. Original-instruction controls at explicit CW027f are
synthetic execution, not native precision or playback observations. The
hooks, selectors, throttle, shipped population and their grades above stand.

### ANIM-STATE-023

The action jump table at `L02482`, decoded out of the image
(`evidence/tables.md`): 1 move `L02706` · 2 `L02707` · 3 **attack**
`L02708` · 4 `L02707` · 5 turn `L02709` · 6 die `L02710` · 7 **shoot**
`L02711` · 8 **cast** `L02712`. Each hook call is an instruction inside
exactly one arm's own range. **attack (3)** fires `vt+0x64` and nothing else,
when the clock equals `class+0x100` = `AttackDelay` (`L02713` the compare,
`L02693` the call). **shoot (7)** fires `vt+0x58`, the projectile spawn, at
`class+0xfc` = `ShootDelay` and only while `class+0xd4` = `Projectile` is
non-zero (`L02714`, `L02715`, `L02716`), then `vt+0x64` at `AttackDelay`
(`L02717`, `L02694`). **cast (8)** fires `vt+0x60` at `AttackDelay`
(`L02718`, `L02719`) and `vt+0x5c` at `ShootDelay` — or at clock 0 instead,
when `unit+0xa4` is `0x3c` (`L02720`, `L02721`, `L02722`, `L02723`).
**die (6) and turn (5) fire none.** So a swing is action code **3** when the
attacker's reach is 1 and **7** when it is greater — `R0549` chooses on
`actor+0x12c > 1` (`ANIM-MSG-005`) — it is entered by that message and left when
the run counter `+0xa0` reaches 0 and the state is forced to 0 (`L02484`); and
a consumer needs one hook on the melee and ranged swing and a **different** one
on a cast

**Confidence.** **High** (the table is static data read by RVA with no
disassembler in the path, and every call site is an address inside one arm's own
range, the arm bounds coming from the table itself rather than from a function
name)

### ANIM-CLOCK-024

The sound fires at `class+0x100`, the `units.reg` **`AttackDelay`** key (stored
at `L02724`), counted in ticks because the attack arm's clock is `+1` per tick
from 0 and its run is sized from the art (`ANIM-PHASE-003`, `ANIM-RUN-004`). The
blow is fired by `actor+0x6c`, loaded from `Data.bin`'s **`attackChargeTime`**
at attack sub-phase 0 and reaching the strike on the tick it decrements to
exactly 0 (`HERO-CADENCE-023`). Both are counted from the same instant: the
swing start sends the `0x71` / `0x72` message at `L02725` and its caller loads
the countdown at `L02726`, inside the **same call** of the actor's own tick
`R0037`. Joined over the (class, definition) pairs by EXP-0071's
resolution rule and measured **identically on both shipped roots**: of 262
pairs, **105 are excluded** because their `attackChargeTime` cell is **absent**,
not zero — every one of them a `Humans` row, whose charge comes from the
equipped weapon instead (`HERO-CADENCE-023` amendment) — leaving **157** that
carry a template answer. Over those, `attackChargeTime - AttackDelay` in ticks
is `-2:2 +0:2 +1:39 +2:6 +3:34 +4:27 +6:21 +7:4 +8:17 +12:1 +15:4`: the sound
leads the numeral on **153**, coincides on **2**, and trails on **2**.
`len(AttackTL) - attackChargeTime` is **negative on 4** pairs, so on those the
blow resolves after the attack animation has already ended and the unit is
standing again. **The exclusion is the finding's own weak point and is stated as
one**: counting an absent cell as -1 would have produced 105 phantom pairs, all
of them negative, and reversed the conclusion. **Consequence for a consumer:**
the swing frame, the swing sound and the damage are three events scheduled from
three different numbers, and binding any two of them together reproduces the
original on at most 2 of the 157 pairs that answer. Parameters, carried with the
figures: `units.reg` joined to `Data.bin` `Units` / `Humans`, both roots,
**template** columns — a weapon overrides `+0x134` and `+0x135` on 16 of 56
classes (`HERO-CADENCE-023` amendment), and a ranged swing adds `R0245`'s
own `extra = (d*256+128)/200` to the charge, so these are the unarmed melee
figures

**Confidence.** **High** (both clocks, and their common origin, are named
instructions re-read here) / **Medium** (the histograms: a corpus join over one
registry and one table, which caps at Medium by rule, though it was measured on
both roots and agrees exactly; and the 105-pair exclusion is a judgement about
what an absent cell means, argued from `HERO-CADENCE-023` rather than re-derived
here)

### ANIM-094

`R0575` is `CUnit`/`CAirUnit` `vt+0x68` (slots `L02727`,
`L02728`), a thiscall with one stack argument `k` (the callee pops 4 bytes). Read whole:

- `0 < k < 3`: silent while `timeGetTime` less `unit+0x190` is below `0x5dc`
  (1500 ms, `L02729`); otherwise the time goes to `+0x190` (`L02730`)
  before a source is chosen. `k = 0` and `k = 3` skip this gate.
- `k = 0`: `Sound[1]` of the class `[L02113][unit+0x20]`
(`L02731` reads offset `+0x4` of the array at `class+0xc0`), with no flag test, played at
  `L02732` with the attenuation offset `[L02733]`.
- `k ≠ 0` with `unit+0x18c & 0x11` non-zero (`L02734`): `R0577`
  returns the bank, and a `DEC` chain reads `[bank+0x24]` for `k = 1`
  (`L02735`), `[bank+0x28]` for 2 (`L02736`) and `[bank+0x2c]` for 3
  (`L02737`); any other `k` keeps the source 0, which is silent.
- `k ≠ 0` with the mask clear: `Sound[k+1]` (`L02738`).
- Both `k ≠ 0` sources play at `L02700` with the attenuation offset
  `[L02739]`.

`+0x190` is one voice timestamp per drawable. The five other callers of
`R0577` — `R0578`, `R0579`, `R0580`, `R0581`
and `R0582` — test it against 3000 or 2000 ms and store it too; they read
`+0x04`..`+0x20` and never a field at `+0x24` or above.

`R0583`, the bank constructor, sets each field to a 20-byte object that
`R0584` builds from one path: `+0x04`/`+0x08` `sfx\<bank>\select1/2.wav`,
`+0x0c`..`+0x14` `sfx\<bank>\command1..3.wav`, then `sfx\<bank>` followed by
`\retreat.wav` at `+0x18`, `\defend.wav` at `+0x1c`, `\idle.wav` at `+0x20`,
**`\easy.wav` at `+0x24`** (`L02740`), **`\hard.wav` at `+0x28`**
(`L02741`) and **`\die.wav` at `+0x2c`** (`L02742`), with those three
strings at `L02743`, `L02744` and `L02745`. `R0585` names
the eight banks `mf_hero`/`ff_hero` (`L02746`/`L02747`), `mf_merc`/`ff_merc`
(`L02748`/`L02749`), `m_mage`/`f_mage` (`L02750`/`L02751`) and
`m_peasant`/`f_peasant` (`L02752`/`L02753`) (`HERO-FIGURE-061`). Both
roots' `sfx.res` hold all 88 constructed leaves, eight banks by eleven files,
the 24 hurt leaves among them; each EN payload differs from its RU payload.

**Confidence.** **High**: the receiver, the selector, the constructor, the
initializer and the five other callers are read whole from raw bytes; the three
displacements are distinct immediates chosen by `k` alone, which rules out one
field shared by both health branches and a field chosen by any other input; the
strings are read at their pushed addresses; the corpus is an exact set
comparison on both roots.

**Unknown.** A later store into a bank field through a computed pointer. Of the
92 dwords naming an address in `[L02750, L02754)`, 69 name `L02755`..
`L02756`, the gap between the mage and the peasant banks, and 23 name a bank:
the initializer, four static thunks that run the default constructor and
register the destructor, the cleanup, the selector, two address comparisons in
`R0578` (`L02757`, `L02758`) and one read of `mf_merc`'s `+0x08` at
`L02759`. None is a store at an absolute bank address; the constructor writes
the fields through `this`.

### ANIM-095

**The `0x73` arm** (`L02661`) tests the drawable's stored health
`unit+0xfc`, the level before this message; it stores the new level `msg+0xc`
after the call, at `L02657`:

- stored equals `msg+0xc`: k = 0 (`L02698`);
- stored differs and is at most -10: no call (the signed compare at `L02699`);
- stored differs, is above -10 and below half: k = 2 (`L02760`);
- stored differs, is above -10 and at least half: k = 1 (`L02761`);
- half is `unit+0x100 / 2`, truncated toward 0 (`L02762`).

No instruction of the arm from its entry to the band test stores at
displacement `0xfc`. The calls on that path are the drawable's `vt+0x20`
(`R0586`, a five-instruction getter of `class+0xd0`) and, when a numeral
is built, `R0566`, `R0568` and `R0587` (`ANIM-NUM-020`).
`R0566` takes the drawable as its sixth argument, reads its `+0x14` and
keeps the pointer at numeral `+0x1c`; every store it makes goes to the numeral,
the stack or the exception chain. `R0568` is passed two arguments, the
numeral and the list at the dispatcher's `this + 0x3f3c`, and `R0587` is
one tail jump. Neither `R0566` nor `R0568` stores at displacement
`0xfc`; the routines they call were not read.

**The state-sync arm** (`L02530`, opcodes `0x6c`/`0x6e`/`0x6f`/`0x70`,
`ANIM-MSG-005`) saves the corpse stage `unit+0x15a` at `L02763`, applies the
mask-bit-3 store at `L02541`, and switches on the current stage at `L02764`
through the six-entry table `L02543`:

```
current 1, saved 0      action 6 for 2 x DyingPhases ticks, then k = 3   L02765
current 1, saved 1      action 6 re-run for 4 ticks, then k = 2         L02766
current 1, saved 2..5   nothing
current 0, saved != 0   +0x84 = 0, then R0551 when +0x18c & 1
current 0, saved 0      nothing
current 2..5            cases L02767 and L02768, no vt+0x68 call
```

The arm's 1004 instructions `[L02530, L02769)` hold no indirect jump
and no return. Five direct jumps leave for the dispatcher exit `L02770`,
each right after a call to the logger `R0588`; every other path reaches the
switch, so the switch runs on each state-sync message that takes none of those
five exits, whether or not bit 3 is in its mask.

**Census.** A raw scan of `.text` for `FF /2` and `FF /4` through displacement
`0x68`, 8- and 32-bit, with and without SIB, finds 96 matches: 21 are
instruction starts in a linear decode from the nearest routine start, 75 lie
inside other instructions. Besides the five above, the 21 hold four calls that
pass the receiver on the stack (`L02771`, `L02772`, `L02773`,
`L02774`), eight that push 0, 2 or 3 arguments where `R0575` takes one
(`L02775`, `L02776`, `L02777`, `L02778`, `L02779`, `L02780`,
`L02781`, `L02782`), one whose caller pops two arguments after a call
through the table `R0228` returns (`L02783`), and three
one-argument calls on MFC objects identified by `CRuntimeClass` and
vtable (`L02784`, `L02785`, `L02786`). Ghidra 12.1.2 (fresh project,
`RepairVtables` twice) holds 20 of the 21 in 15 functions and leaves `L02771`
undisassembled; `callto:R0575` finds only the two vtable slots, and no dword
names either slot. Its seven same-function pairs of a displacement-`0x68`
load and a call through the loaded value, and the calls through frame-local or stack slots in
the seven functions that also hold such a load, include no one-argument call of
a `+0x68` slot; indirect calls in orphan or undisassembled bytes lie in routines
with no displacement-`0x68` load.

**Confidence.** **High** for the five sites and the condition of each `k`: named
instructions, and the dispatch is the table `ANIM-MSG-005` decodes / **Medium**
that no other emitter exists: the census covers direct, register-held and
stack-held slot calls, and a slot address computed another way, or code outside
both decodes, is not excluded / **Unknown** how often the stage-1 arms run:
which server events send a state sync while a drawable's stage is 1 was not
measured.

**Unknown.** Alternative receiver bindings and native message ordering.
ANIM-133 resolves the three intervening calls for their native bindings:
`L02787` is the campaign record's scalar getter, `R0589` is scalar
with stack stores only, and `R0544` copies 12 bytes into
drawable offsets `0xe4..0xef`, outside stage `+0x15a`.

**Amended.** `EXP-0463`: the voice call of the stage switch is made when the message is handled, after the action stores and before `R0551`, not after the action's ticks (`ANIM-126`). The senders of the stage-1 syncs are the death arm, the negative-health bleed and the session-entry projection (`ANIM-126`); the zero-change `0x73` of event 0 comes from a zero-damage strike (`ANIM-125`). ANIM-133 resolves the three-helper Unknown for native receiver bindings.

### ANIM-096

**Which drawables.** `+0x18c` bit `0x1` marks the hero-shaped drawable
(`HERO-APPEAR-041`, which reads bit `0x4` as the sex axis and bit `0x2` as the
fighter/mage axis). The three routines read here that set it take it from a
wire class or a record's `+0x74` bits, none from an equipment slot:

- `R0590`, for a wire class `T` in `[0x20, 0x40)`: `0x9`, then `0x4`
  when `(T - 0x21) & 1` and `0x2` when `(T - 0x21) & 2`;
- `R0591`, on the view's `+0x3f54` drawable: the word set to 0, then
  `0x2` on a record's `+0x74` bit `0x40`, `0x4` on its bit `0x80`, then `0x9`;
- `R0592`: `0x29`, or `0x2b` on the record's `+0x74` bit `0x40`, then
  `0x4` on its bit `0x80` (`L02788`..`L02789`).

For `T` below `0x1a`, `R0590` sets `0x18` instead, `0x4` from its
parameter's bit `0x80`, and `0x2` for `T` `0x17` and `0x18`: such a drawable
takes the bank through bit `0x10`, the mage pair when bit `0x2` is set, and
otherwise the slot-0 choice between the merc and the peasant pairs.

**Each branch, bit `0x1` set.** Bit `0x1` is inside the receiver's `0x11`
mask, so every `k ≠ 0` takes a bank (`ANIM-094`), and `R0577` returns on
bit `0x2` or bit `0x1` before it reads slot 0 at `+0x15c`:

```
event (ANIM-095)                   k   source (ANIM-094)             gate
0x73, health unchanged             0   Sound[1] of the drawn class   none
0x73, stored >= half and > -10     1   bank +0x24  easy.wav          1500 ms
0x73, stored <  half and > -10     2   bank +0x28  hard.wav          1500 ms
0x73, stored <= -10                -   nothing
state sync, stage 0 -> 1           3   bank +0x2c  die.wav           none
state sync, stage 1 -> 1           2   bank +0x28  hard.wav          1500 ms

bank   bit 0x2 set:   m_mage, or f_mage when bit 0x4 is set
       bit 0x2 clear: mf_hero, or ff_hero when bit 0x4 is set
```

**The no-damage cue.** `k = 0` reads `Sound[1]` of `unit+0x20`, which
`R0551` derives from slots 0 and 1 and the mage bit (`HERO-APPEAR-042`).
Shipped `Sound[1]`, the same on both roots: `units\soft2` for classes 1, 2, 14,
23 and 24; `units\metal1` for 3, 5, 10 and 12; `units\metal2` for 7 and 9;
`units\metal3` for 4, 8, 11 and 13; `units\soft3` for 15. `Sound[k+1]`, the
mask-clear source, cannot be reached while bit `0x1` is set.

For a hero-shaped drawable, fighter or mage, the sex bit picks the `f` bank of
the pair and the mage bit picks the mage pair over the hero pair. Equipment
slots 0 and 1 and the drawn class `+0x20` enter neither the bank nor the field
of any `k ≠ 0` event; they reach only the `k = 0` cue, through `+0x20`.

**Confidence.** **High** for the table given bit `0x1`: every step is a named
instruction in a routine read whole (`ANIM-094`, `ANIM-095`) / **Medium** for
the `k = 0` sample per class: the sixteen ids are `HERO-APPEAR-042`'s, and the
`units.reg` join is re-executed by the loader's rule, which caps at Medium /
**Unknown** whether another store changes a hero drawable's bits later: Ghidra
counts 68 stores to displacement `0x18c` in 39 functions; 13 are the three
writers above and 6 are the dispatcher's, which set or clear only bits `0x8`,
`0x20` and `0x80`; the other 49, in 35 functions, were not classified by object
type.

**Amended.** `EXP-0463`: for a class below `0x1a` the parameter whose bit `0x80` sets the sex bit is server `actor+0x4b`; its sources are the spawner's secondary-key byte and placement flag bit 2, or the Humans-row constructor value (`ANIM-127`).

### ANIM-119

- **Slots.** `CUnit` (`L02468`) and `CAirUnit` (`L02585`) hold the same five routines: `+0x6c` `R0578`, `+0x70` `R0580`, `+0x74` `R0581`, `+0x78` `R0579`, `+0x7c` `R0582`. Ten dwords name them (`L02790..L02791`, `L02792..L02793`) and no direct call does (`evidence/xref-voice.txt`).
- **Common shape.** Each routine calls `timeGetTime` (`[L00849]`), subtracts `unit+0x190` and leaves when the unsigned difference is below its threshold; otherwise it stores the time to `+0x190` before it chooses a source. The bank comes from `R0577` (`HERO-APPEAR-055`), the source is a bank field, `R0574` turns the drawable's position into pan and volume without a branch on its result, and a null field skips the play. A refused call stores nothing; a call whose field is null stores the stamp and plays nothing.
- **Fixed readers.** `+0x70` reads `[bank+0x1c]` (`defend`), `+0x74` `[bank+0x18]` (`retreat`), `+0x7c` `[bank+0x20]` (`idle`), each at 3000 ms (`0x0bb8`), with no `rand()` call.
- **`+0x6c`, 3000 ms.** `r = rand() >> 13` (`L02794`, `L02795`), 0..3. For `r < 3`, or when the bank is `m_peasant` or `f_peasant` (`L02752`, `L02753`; `L02757`, `L02758`), it reads `[bank + 0xc + (r mod 3) * 4]` (`L02796`): `command1..3`. Otherwise (`r = 3`, any other bank) it reads `[bank+0x1c]`, `defend`. Over the 32768 values of `rand()`, a non-peasant bank gives `command1`, `command2`, `command3` and `defend` 8192 each; a peasant bank gives `command1` 16384, `command2` 8192, `command3` 8192.
- **`+0x78`, 2000 ms.** `r = rand() >> 14` (`L02797`..`L02798`), 0 or 1, reads `[bank + 4 + 4r]` (`L02799`): `select1` or `select2`, 16384 values each.
- **State.** The CRT generator (`R0179`, result masked to `0x7fff` at `L02800`, seed in the per-thread block `+0x14`) is the only other state a reply touches; the chooser and `+0x6c` and `+0x78` draw from it (5 of 88 direct call sites). No bank field, index or cycle position is stored. Executed on the original instructions with a sentinel in each of the eleven fields of all eight banks, 192 bank and draw combinations and 40 clock and stamp combinations, the reply is a function of bank, draw, clock and stamp alone: the writes the five routines make are the stamp and nothing else, and the drawable displacements read are `+0x08`, `+0x0c`, `+0xe0`, `+0x15c`, `+0x18c` and `+0x190`.
- **Boundary.** At stamp 0, `+0x6c`, `+0x70`, `+0x74` and `+0x7c` play at clock 3000 and not at 2999, `+0x78` at 2000 and not at 1999. A stamp later than the clock (difference wraps) plays.

**Confidence.** **High** for the five routines' arithmetic, thresholds, fields and stores: each read whole and executed on both roots (EN = RU, one executable), which excludes a fixed slot per gesture, a cycle and a stored index; the only drawable store in them is the stamp. **Medium** for the rand-share count (a direct-call census) and for the unit-wide state claim: a store through a computed pointer to a drawable or a bank, or a bulk copy of a drawable, is outside the sweep. **Unknown** whether the sample service `R0386`, reached at the end of each routine, refuses or queues the request (channel availability, its global at `L02703`); the executed boundary is the call, not a sound.

### ANIM-120

- **Constructor.** `R0593` stores zero to the field `+0x190` (`L02801`) beside `+0x18c` and the other `0x15c` block fields; executed on a block filled with `0xa5`, it leaves `+0x190 = 0` and `+0x18c = 0`. `CAirUnit`'s constructor `R0594` calls it and then sets its vtable (`L02802`, `L02586`).
- **Copy.** `R0595` copies the field `+0x190` from the source object to the destination object (`L02803`, `L02804`). Its only caller is `R0596` (`L02805`), which no direct call or table dword names.
- **Creation.** Direct calls to the default constructor: `L02806`, `L02807`, `L02808`, `L02809`, `L02810`, and `L02811` for `CAirUnit` (`evidence/xref-voice.txt`). `L02807` is the client dispatcher's unit creation block (`ANIM-117`). The drawable is not among the classes the save stream carries (`SAV-STREAM-013`), so no stamp is loaded from a save.
- **Other stores.** The sweep for displacement `0x190` finds 22 register-based stores: 8 on this object (constructor, copy, the hurt hook and the five voice routines) and 14 in 11 other owner routines (nearest-entry owners `R0597`, `R0598`, `R0599`, `R0600`, `R0601`, `L02812`, `R0602`, `L02813`, `L02814`, `L02815`, `L02816`), classed as other objects by their neighbouring fields, not by a receiver proof (`evidence/store-0x190.txt`).
- **Hurt hook.** The hurt hook `R0575` tests the same stamp against 1500 ms (`L02729`, `ANIM-094`), so a hurt reply is admitted at clock 1500, not 3000.
- **First reply.** Executed: after the constructor, `+0x6c` plays no recording at clock 2999 and plays one at 3000. The comparison is unsigned and the clock is `timeGetTime` (WINMM), so a first reply of the five voice readers is refused only while `timeGetTime` is below the threshold.

**Confidence.** **High** for the constructor's value, the copy, and the first-reply boundary. **Medium** that every drawable starts from the constructor: five direct constructor callers and no vtable or clone path were found, and the load path to them was not traced. **Unknown** the stamp of a drawable made by a bulk copy, and the first-reply result in the first 3000 ms after the operating system's clock starts or after the 32-bit clock wraps.

## Projectile driver and draw

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-PROJ-025 | The projectile's own driver, read — closing `ANIM-VT-012`'s Unknown. A cast's map object lives by a COUNTDOWN the spawner sets, and a homing shot re-aims every tick. | High / Medium | ● active | [EXP-0139](../experiments/EXP-0139-spell-pictures/) |
| ANIM-PROJ-026 | The projectile's draw: the frame index is `Phases * facing + phase`, `Flip` HALVES the sheet exactly as it does for a unit, and two sprites outside the registry are the smoke. | High / Medium | ● active | [EXP-0139](../experiments/EXP-0139-spell-pictures/) |
| ANIM-CAST-027 | (rom.exe) The projectile draw is not uniform: `R0556` switches on the picture id, with six special arms and a smoke trail in the default one. | High / Unknown | ● active | [EXP-0140](../experiments/EXP-0140-cast-art-drawn/) |
| ANIM-PHASECLOCK-028 | (rom.exe) A projectile advances one sheet frame every TWO game ticks, and four picture ids replace that clock with their own. | High | ● active (amended, partially retracted) | [EXP-0167](../experiments/EXP-0167-spell-art/) |
| ANIM-BOLTDRAW-034 | Each 8-byte bolt point is stamped once; all points of one link share its sheet frame, with picture 36 adding five times the link tag and both arms centring by eight pixels. | High | ✔ promoted (partially retracted) | [EXP-0178](../experiments/EXP-0178-bolt-path/); [EXP-0504](../experiments/EXP-0504-bolt-figure/) |
| ANIM-BOLTRAMP-035 | (rom.exe) Pictures 34 and 36 replace the projectile frame clock with a fixed 13-step ramp, and that ramp and the two shipped sheets tile each other exactly. | High | ✔ promoted (amended) | [EXP-0178](../experiments/EXP-0178-bolt-path/); [EXP-0504](../experiments/EXP-0504-bolt-figure/) |

### ANIM-PROJ-025

`R0558` is `CProjectile`'s `vt+0x3c` (`ANIM-VT-012`, vtable
`L02587`). The engine's own names for the fields come from the savegame
loader `R0099`, which reads a `[Prj%d]` section key by key: `x` `+0x08`,
`y` `+0x0c`, `z` `+0x10`, **`picture` `+0x20`**, `dir` `+0x6c`, `phase` `+0x70`,
`lastaction` `+0x74`, `action` `+0x84` (byte), `actiondir` `+0x85` (byte),
`actiontarget` `+0x86` (word), `actionx`/`actiony`/`actionz`
`+0x88`/`+0x8c`/`+0x90`, `actionphase` `+0x94`, `actionsegments` `+0xa0`,
`actionspell` `+0xa4`; the object is `0x14c` bytes and the save also carries
`[Projectiles] Count`, `FreeIndex` and `IDs`. The driver:
`L02817` reads the 32-bit field `actionsegments` at `+0xa0` — when it is 0 it
returns 0 and the object is finished (`L02818` sets the return value to 0), first setting the
view's `+0x74` when `picture` is 13 (`L02819` compares the 32-bit field `+0x20` with 0xd);
otherwise `L02820` and `L02821` fetch `[L02822][picture]`,
`L02823` takes the 16-bit field `+0x86` and walks the world's hash at `+0x9bc` for
`actiontarget`, and `L02824`, `L02825`, `L02826` **overwrite
`actionx`/`actiony`/`actionz` from the target's current position on every tick**
— so tracking is a property of having a target, not of `Homing`, which the
driver never reads; `L02827` calls L02828 to recompute `actiondir` from the
two coordinates; `L02829`..`L02830` is three signed divides by
`actionsegments`, i.e. **`step = (actionx - x) / actionsegments`** in each
axis, a straight line onto wherever the target is now; `L02831` and `L02832`
increment `actionphase`; `L02833` biases `picture` by −13 and `L02834` bounds the
result at 0x33, a `switch` on `picture` biased by 13 and spanning
13..64; and `L02835`..`L02836` set `lastaction = action` and
`actionsegments -= 1` before returning 1. Decompiled rather than cited by
address, the arms are: `0x0d` writes tile bit `0x2000` over a 3x3 block at tick
8; `0x12 0x18 0x1c 0x28 0x2c 0x30 0x34 0x36 0x38 0x3e 0x40` snap to the target
and notify it; `0x14 0x1e` do the same into the target's own slot array at
`+0x128`; `0x22 0x24` take `phase` from a fixed 13-step ramp; `0x33` travels and
raises a quake at tick 8; `0x3c` sets `phase = actionphase` with no modulus; the
default travels with `phase = (actionphase / 2) % Phases`. **So caster-attached,
target-attached, travelling and area are one mechanism**: one class, one array,
one driver, and a `switch` on the picture id. The other spawner of the same
object is `R0603`, a unit's shot, which takes `picture` from the
`units.reg` class record's `+0xd4` and sets `actionsegments` to
`ftol(<distance term>) / 200`. **G2:** the lifetime is a per-object integer with
no clamp in the driver, so a longer-lived cast is a spawner change and touches
no shipped file

**Confidence.** High (the entry, the target re-read, the three `IDIV`s, the
switch bias and the exit are each their own instructions in a dumped listing,
and the field names are the game's own registry keys rather than ours) / Medium
(the per-arm reading, which is decompiled)

### ANIM-PROJ-026

`R0556` is `CProjectile`'s `vt+0x28`. `L02837`, `L02838` and
`L02839` refuse a `picture` past the global at `L02840` and `L02841`, `L02842`,
`L02843` an empty slot — in both cases **nothing at all is drawn**, which is
what makes `MAGIC-PIC-026`'s seven undefined ids silent rather than a crash.
Then `L02844` with `L02845` and `L02846` subtract `Width / 2` and
`L02847` with `L02848` and `L02849` `Height / 2`, so the art is centred on
the object, minus `z` (`L02850`) and the view's own offset (`L02851`);
`L02852` reads `dir` (`+0x6c`), `L02853` subtracts 8 and `L02854` masks the
result with 0xf, which folds `dir` to a 16-way facing biased by 8;
`L02855` reads `Flip` (the 32-bit record field `+0x30`), `L02856` compares the facing with 8 and
`L02857`/`L02858` form 16 minus the facing, which **mirrors facings 9..15 onto
7..1, leaving nine stored facings**; `L02859` reads the record's `Phases` (the 32-bit field `+0x10`),
`L02860` multiplies it by the facing and `L02861` adds the phase, giving
**`frame = Phases * facing + phase`**, overridden to `frame = phase` when
`RotationPhases == 1` (`L02862`, `L02863`, `L02864`); and
`L02865` biases `picture` by −7, `L02866` bounds the result at 0x35 and
`L02867` reads one byte of the table at `L02868`: a second `switch` on `picture`,
**biased by 7 and spanning 7..60 through a `0x36`-byte arm-index table at
`L02868`** — a different bias and a different span from the driver's.
Decompiled: `Palette == 0` blits through the surface's `vt+0x14` with the shared
the global at `L01248` (`projectiles.pal`) and otherwise through `vt+0x18` with the
sheet's own; `class+0x38 == 0` calls `R0604`, so a projectile's sheet is
loaded **on first draw**, not at start-up; and `picture` 10 and 12 additionally
blit `[L02869][0]` and `[1]` — the two
`graphics\projectiles\smoke%d\sprites.16a` sheets that `projectiles.reg` does
not mention — once per point of the object's own trail array (`+0x13c` count,
`+0x140` pointer). **G2:** the frame formula is the binding limit on the art — a
sheet must hold `Phases * 9` frames when `Flip` is set and
`Phases * RotationPhases` otherwise, so adding a facing or a phase changes a
shipped `.16a` or `.256`'s bytes

**Confidence.** High (every quantity in the formula is its own `IMUL`, `SAR`,
`SUB` or `CMP` in a dumped listing, and the two refusals are their own
conditional jumps) / Medium (the palette fork, the lazy load and the smoke
trail, which are decompiled)

### ANIM-CAST-027

The routine is the projectile draw (`MAGIC-PIC-027` fixes `+0x20` as the picture
id and `0x14c` as the object size); it is reached only through the vtable slot
at `L02870`. It looks the record up as `[L02822][picture]`, centres by
the registry's `Width`/`Height` halves, folds the direction (`(dir - 8) & 0xf`,
mirrored above 8 when `Flip` is set), computes `frame = Phases * facing + base`
unless `RotationPhases == 1`, and then forks: **id 7** calls `R0605` and
blits nothing; **id 20** (`healing`) has an arm that falls straight through with
**no blit at all**; **ids 34 and 36** (`lightnin`, `chain`) ignore the
projectile's own position and instead iterate a point list at `+0x114` of length
`+0x118`, stamping the sheet at each `(x-8, y-8)` -- chain adding a per-point
phase of `byte * 5` from the fourth `short` of each 8-byte record, which is what
its 35 frames are for; **id 51** (`Meteor`) forces frame 8 and lifts the sprite
by `0x1c - 4*frame` while its frame is below 8, so the meteor descends; **id
60** (`teleport`) blits the raw `+0x70` frame with no direction fold; and the
**default** arm, after the sprite pass, walks a trail of `+0x140` points at
`+0x13c` and stamps `[L02869][0]` for id 10 and `[1]` for id 12 -- the two
`smoke%d` sheets `R0606` loads and the registry never names. The blit
slot itself is chosen by the record's `Palette` (`REG-PROJ-087`): zero takes
`vt+0x14` with the global at `L01248` and a light level, non-zero takes `vt+0x18` with
the sprite's own table and then `vt+0x38` on the `b` sibling when one exists.
**G2:** every one of these arms is a hard-coded id in `.text`, so re-pointing a
spell's art at a different `projectiles.reg` row moves the art but **not** the
behaviour -- a customisation that gave lightning a new id would lose its
polyline, and one that reused id 20 would get a picture that never draws. That
is an engine limit and lifting it edits `rom.exe`

**Confidence.** High (the switch and every arm read from the raw listing of one
routine, with the vtable offsets confirmed against a side-by-side dump of both
sprite vtables) / Unknown (what the routine's third argument is: it is a light
level on the `.256` arm -- multiplied by `0x200` and added to a table base --
and a raw table **pointer** on the `.16a` arm, and the only value safe for both
is 0. The caller was not read)

### ANIM-PHASECLOCK-028

`ANIM-PROJ-025` reported the default arm as `phase = (actionphase / 2) % Phases`
from a decompilation; the code is a sign-corrected signed halving of the
just-incremented `actionphase` (`L02871` increments it; `L02872`..`L02873`),
a signed divide of that half by the 32-bit field `+0x10` of the record, its `Phases`
(`REG-PROJ-086`, `L02874`), and a store of the remainder as `phase` at `+0x70`
(`L02875`); the same three steps appear at `L02876`..`L02877`
on the sibling arm. When the record pointer is null the arm stores `phase = 0`
instead (`L02878`). The exceptions, from the 52-byte switch table at
`L02879` read out of the PE by section walk: picture **60** takes
`phase = actionphase - 1` with no modulus (`L02880` decrements it), pictures **34**
and **36** take `phase` from a 13-entry jump table at `L02881` bounded by
bounded at 0xc (`L02882`) that yields the constants 4, 3, 2 and 1, and picture
**51** takes `phase = actionphase` raw (`L02883`). The clock is therefore not
the unit clock: `MAGIC-CASTANIM-029` fixes a unit's action frame as the tick
index with no divisor, so **a consumer needs two frame clocks, one per tick for
units and one per two ticks for projectiles.** Corroboration from the shipped
data: the two burst lifetimes that are not the default 16 are 18 and 22
(`MAGIC-BURSTLIFE-034`), against `acid`'s 9 phases and `fireexpl`'s 11 --
`2 * Phases` in both cases, which is one pass under this clock and under no
other divisor. **G2:** the divisor is a `SAR` by 1 in `.text` and the
per-picture exceptions are table entries, so the frame rate of effect art is an
**engine** limit; a sheet's frame count is data and free

**Confidence.** High (the divisor, the modulus and the increment are each their
own instruction in a listing dumped over an address range, the exception table
is read from the PE by section walk, and the competing model "one frame per
tick" is excluded both by the `SAR` and by the two shipped lifetimes agreeing
with `2 * Phases` rather than `Phases`)

**Amended.** Ramp enumeration amended: the 13-entry table yields five distinct
constants, 0 among them, and the full sequence is in `claims/retracted.md`; the
two-tick divisor, the modulus, the increment, the null-record store and the
picture 51 and 60 exceptions stand.

### ANIM-BOLTDRAW-034

The record is written in the generator's rotation loop:
`L02884` and `L02885` store the two 32-bit halves of an 8-byte record slot
(offsets 0 and 4 of slot `index * 8`), assembled from a zero 16-bit word
(`L02886`, local stack offset `0x2c`), the tag byte argument read from `[EBP+0x28]` in the generator's stack frame
(`L02887`/`L02888`, local stack offset `0x2e`) and a zero byte (`L02889`, local stack offset `0x2f`), plus two coordinate words each
produced by the CRT float-to-int helper R0279 on a rotated double. So the layout is
`int16 x; int16 y; int16 zero; uint8 tag; uint8 zero`. Both draw arms iterate
the array at `+0x114` with stride 8 (`L02890` forms the slot address) while
the index is below `+0x118`, read the two coordinate words as sign-extended 16-bit
values (`L02891`, `L02892`), and subtract the immediate 8 from each
(`L02893`, `L02894`) rather than the `Width`/`Height` halves the routine
computed at `L02895` to `L02849`. The **frame** is `+0x70` for picture 34,
the same value for every point of the figure, and `phase + 5 * record[6]` for
picture 36 — `L02896`..`L02897` read the tag byte at offset 6 of the record, multiply
it by 5 and add the phase. The **sheet** is a
constant, not the projectile's own picture:
`L02898` loads the table entry at offset 0x88, `[L02822][34]`, and `L02899` the
entry at offset 0x90, `[L02822][36]`. Neither arm
reads the direction. Both shipped records carry `RotationPhases = 1`, which
already disables `ANIM-PROJ-026`'s `Phases * facing` term, so **no rotation law
reaches these two arms from either side**. **G2:** the 8-pixel centring is an
immediate, so a replacement sheet of any size other than 16 by 16 draws
off-centre; the tag stride 5 is an immediate; the two sheet indices are
immediates, so the art of these two figures cannot be re-pointed by data

**Confidence.** High. Every store, load, immediate and displacement above is its
own instruction in a dumped listing, and all are re-read from the raw image
bytes at their own addresses by `tools/boltpath -mode asserts`, 74 of 74 holding
on three roots. The `RotationPhases` values are a corpus measurement over both
preserved roots (`tools/castflight -mode map`)

**Amended.** The same-frame headline is narrowed from an entire figure to
one link; picture 36 can assign distinct frames to distinct links. The tag
operand is a stack argument, not caller-object+0x28. Those clauses are
partially retracted in claims/retracted.md; MAGIC-275 and MAGIC-280 supply
the exact argument and draw scopes. Record layout and body formulas stand.

### ANIM-BOLTRAMP-035

`ANIM-PHASECLOCK-028` reports that these two picture ids take their phase from a
13-entry jump table at `L02881` bounded at 0xc by `L02882` and
yielding the constants 4, 3, 2 and 1. Read out of the PE by section walk, the
table's thirteen entries in order are the arms `L02900`, `L02901`,
`L02902`, `L02903`, `L02878`, `L02903`, `L02902`, `L02903`,
`L02878`, `L02903`, `L02902`, `L02901`, `L02900`, which write `+0x70` =
**4, 3, 2, 1, 0, 1, 2, 1, 0, 1, 2, 3, 4** for `actionphase` 1 through 13. The
index is `actionphase - 1` and `actionphase` is incremented once per tick at
`L02871` (an increment), so the ramp is one value per tick and consumes the normal
caster route's 13 successful calls (`MAGIC-BOLTSTILL-072`, amended). The fifth constant is 0, which
`ANIM-PHASECLOCK-028` did not list. Two independent quantities agree on the
frame budget: the ramp's range is 0 to 4 and `lightnin` ships **5** phases;
picture 36 adds `5 * (linkIndex mod 7)` to the same ramp, so its range is 0 to
34 and `chain` ships **35** phases, which is `7 * 5`. **G2:** the ramp is a
`.text` jump table, so the flicker pattern of these two figures is an engine
limit; a sheet with a different frame count is a data change that the ramp will
index past or under

**Confidence.** High for the ramp: the table is read out of the PE by section
walk with no disassembler in the path, and the bound, the index and the
increment are each their own instruction. High for the tiling, because the two
quantities come from disjoint instruments — a `.text` jump table and a walked
frame count in a shipped archive, measured on both preserved roots
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration. The agreement between the per-root
data stands as corroboration: the archive frame counts are separate files in each root.

**Amended.** The universal 13-tick lifetime clause is narrowed to normal
caster construction in claims/retracted.md. MAGIC-281 preserves the ramp
but measures route-specific initial phases and countdowns. The thirteen
values and shipped frame-budget agreement stand.

## Retained area effects

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-AREAPHASE-029 | (rom.exe) A retained area-effect cell advances one sheet frame every two game ticks, from the shared object clock plus a per-cell constant offset. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| ANIM-AREAOFFSET-030 | (rom.exe) The per-cell phase offset is the cell's position in the VIEW, not in the world, and the one other animated term in the same feature uses the world position instead. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| ANIM-AREANOSTATE-031 | (rom.exe) A retained area effect carries no per-instance animation state at all, on either side. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| ANIM-AREATRANS-032 | (rom.exe) The blast and the staged modes create ordinary projectiles, so `ANIM-PROJ-025` and `ANIM-PROJ-026` govern them unchanged. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |
| ANIM-WALLFIREFRAME-033 | (rom.exe) `wall_of_fire`'s overlay frame law is not `mod Phases`: it uses five of its sheet's eleven frames, on a 10-tick cycle. | High | ● active | [EXP-0176](../experiments/EXP-0176-area-draw/) |

### ANIM-AREAPHASE-029

`R0607` computes, per arm,
`phase = abs(view+0xa70 / 2 + screenX * screenY) mod Phases`:
`L02904` reads the counter at `view+0xa70`, `L02905` halves it (signed shift
right by 1), `L02906` multiplies it by the 32-bit stack slot at frame offset +0xc,
`L02907` calls the absolute-value helper R0376, `L02908` is a signed
divide by the 32-bit field `+0x10` of the record, its `Phases`
(`REG-PROJ-086`), repeated at `L02909` and `L02910`. `view+0xa70` is
`ANIM-CLOCK-001`'s object and water counter, incremented once per paced
play-loop tick, so the rate is one frame per two ticks, the same rate
`ANIM-PHASECLOCK-028` establishes for a projectile and half the rate
`MAGIC-CASTANIM-029` establishes for a unit action. The offset differs per cell,
so cells of one effect are at different frames in the same video frame. **G2:**
the divisor is a `SAR` by 1 in `.text`; the sheet's frame count is data

**Confidence.** High: every term is its own instruction, re-asserted as bytes on
both roots, and the routine was read whole
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### ANIM-AREAOFFSET-030

The two draw call sites pass `screenX = 32 * viewColumn + 16`
(`L02911` shifts left by 5, `L02912` adds 0x10) and `screenY` = the average of
two per-corner altitude arrays at `view+0xb8` and `view+0xbc` (`L02913`,
`L02914`, `L02915`), both indexed by the cell's row and column **within the
visible window**. `ANIM-AREAPHASE-029`'s offset is their product, so the frame a
given world cell shows is a function of where that cell currently sits on
screen. The `wall_of_fire` light flicker in `R0608` computes its offset
from the world cell instead: `L02916` halves a value (signed shift right by 1) and
`L02917` multiplies it by the 32-bit frame local at offset −0x20 over the cell key's own low byte and
high byte (`L02918`, `L02919`). The two laws are therefore not the same
function of position

**Confidence.** High for both expressions, byte-asserted at their own addresses
in two routines read end to end. That scrolling the map changes a standing
effect's frame is a consequence of the first expression, not a separate
observation; the instrument that would show it is a running original

### ANIM-AREANOSTATE-031

The client stores one dword per covered cell, a mask of active area layers, and
`R0483` writes nothing else into the hash node: the only stores into that
node's payload are `L02920` on the add and `L02921` on the remove, and the
third `+0xc` store in the routine, `L02922`, belongs to `wall_of_fire`'s
separate hash at `view+0xa7c`. The draw recomputes the frame from the shared
counter and the cell's screen position on every pass, so there is nothing to
seed, nothing to reset and nothing to save. Two consequences a consumer needs:
an effect laid later is automatically in phase with an earlier one at the same
screen cell, and the frame does not restart when an effect is re-laid over the
same cell. This is the opposite of the drawable model `ANIM-PROJ-025`
establishes for a projectile, whose phase is its own `actionphase` field

**Confidence.** High: the store is one routine read end to end, and the absence
is structural rather than a search: the hash node has one payload dword and both
writers write the mask into it

### ANIM-AREATRANS-032

The client arm for opcode `0x86` allocates `0x14c` bytes, runs the projectile
constructor `R0609`, writes the message's picture into `proj+0x20`, the
cell centre into `proj+0x8`/`+0xc`, the message's lifetime word `msg+0x0f` into
`proj+0xa0` (`L02923`, `L02924`), which is `actionsegments`, and action code
1 into `proj+0x84`. The three lifetimes the senders write are 16 (`L02925`),
18 for `acid_stream` (`L02926`) and 22 for `fire_ball` (`L02927`). Under
`ANIM-PHASECLOCK-028`'s default clock of one frame per two ticks, 18 is exactly
one pass of `acid`'s 9 phases and 22 exactly one pass of `fireexpl`'s 11. The
default 16 is **not** a whole pass of either sheet it serves: `smallxpl` has 11
phases and `Meteor` 9, and picture 51 takes `ANIM-PROJ-026`'s raw-`actionphase`
arm rather than the default one. A transient burst is therefore not a loop and
is not tied to the effect's duration, but only two of the three lifetimes are
one pass of their own sheet

**Confidence.** High: the allocation, the constructor call and the four field
writes are in one listing read end to end, and the lifetime immediates are
byte-asserted at their own addresses on both roots
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### ANIM-WALLFIREFRAME-033

Arm 1 of `R0607` replaces the record's phase count with the immediate 5
(`L02928`) and biases the remainder by 3 (`L02929`),
so the frame is `abs(counter/2 + screenX*screenY) mod 5 + 3` and only frames 3,
4, 5, 6 and 7 of the 11-frame `firewall` sheet are ever drawn. At one frame per
two ticks the cycle is 10 ticks, against 22 for an 11-phase sheet under the same
clock. The other three arms take the record's own `Phases` (`L02908`,
`L02909`, `L02910`). The `wall_of_fire` light flicker in `R0608` also
divides by 5 (`L02930`) but then takes the low bit of the quotient
(`L02931`), giving 10 ticks bright and 10 ticks dim, so the sprite cycle and
the light cycle have the same period and different waveforms. **G2:** the 5 and
the 3 are `.text` immediates, so the five frames a burning cell can show are an
engine limit that no `projectiles.reg` edit moves

**Confidence.** High: the divisor, the bias and the three unmodified arms are
each their own instruction, byte-asserted on both roots; the eleven-frame sheet
count is `REG-PROJ-086`'s field measured over the shipped `projectiles.reg`
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

## Effect action and spell visuals

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-PARK-039 | (rom.exe) The act state the refusal writes falls outside the actor tick's own state switch, so no state arm runs either — and the value does not identify the spell. | High / Medium | ● active | [EXP-0179](../experiments/EXP-0179-effect-action/) |
| ANIM-044 | (rom.exe) Actor marks rebuild once per paced presentation tick before the unit action branch, then draw around that actor body by signed depth. | High | ✔ promoted | [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/) |
| ANIM-045 | (rom.exe) Shield record composition is deterministic and ordered: every Component A source point and pair precedes every Component B source point and pair. | High | ✔ promoted | [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/) |
| ANIM-046 | (rom.exe) A Meteor is stationary in world space and draws sixteen exact phases: eight falling stamps of frame 8, then impact frames 0–7. | High | ✔ promoted | [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/) |
| ANIM-047 | (rom.exe) Projectiles have a no-op grid-registration slot and are associated with the separate collection between retained-area phases. | Medium | ✔ promoted (partially retracted) | [EXP-0189](../experiments/EXP-0189-spell-visual-algorithms/); [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |

### ANIM-PARK-039

`R0016` writes `actor+0x54 = 0x1a` on both exits of progress arm 4
(`L02932`). Four instructions after the call, the actor tick dispatches on
that field: `L00538` reads the 32-bit field `+0x54`, `L02933` subtracts 1, and
`L02934`/`L02935` branch when the biased index (kept in a frame slot at −0x78) exceeds 0xe, to `L02936` (rel32
`0x806`). `0x1a − 1 = 0x19` exceeds `0xe`, so the branch is taken, and
`L02936` is the routine's epilogue — `L02937` is its return — not a fifteenth arm.
The same value is written by three other sites of the same routine: the
early-out at `L00089`, the `ord+0x09 == 0xff` arm at `L00544`, and the
default progress arm at `L00544`/`L00542`. So `0x1a` is the machine's
"nothing to do" act state and carries no information about why. What the client
draws for an actor in it is a separate question with a separate answer:
`MAGIC-ACTOR-066` reports the unit draw `R0552` holding the animation
frame for the same spell, keyed on the effect record's `+0x0e`
(`2 × id + 8 = 0x30`) on the drawable's own list rather than on `actor+0x144`.
The two hierarchies are disjoint and this ledger's header says only the actor is
serialized, so the animation hold and the action refusal are two consequences of
one spell, independently keyed and independently readable

**Confidence.** **High** (the comparison, its branch displacement, the epilogue
and the four writers of the value are named instructions, byte-asserted at their
own addresses by `tools/effectgate -mode asserts`) / **Medium** for the relation
to the draw: `MAGIC-ACTOR-066` was read and cited, not re-derived, and it grades
its own animation-hold reading Medium

### ANIM-044

The paced `0x401` map tick drives CUnit/CAirUnit; their unit driver invokes
vtable `+0x50` once before its action switch and both vtables bind that slot to
`R0610`. Rebuild finishes before the element countdown decrements. The
compositor centres a mark at `x=unit[+0x60]-Width/2+dx`,
`y=unit[+0x64]-unit[+0x68]-unit[+0x10]-Height/2+dy-depth`; positive depth draws
before that actor body and non-positive depth after it, preserving array order
within each pass. At default 16 ticks/s, one rebuild is one 62-ms paced tick;
game speed scales wall time.

**Confidence.** High. Driver call, two vtable bindings and decrement order are
byte-asserted on all three roots; the anchor/pass split is the independently
reviewed consumer chain.

### ANIM-045

Within one sampled midpoint iteration the source order is
`(+,+),(-,+),(+,-),(-,-)` including duplicates; each source's first record
precedes its second. The shared mark compositor then performs its positive-depth
pass and non-positive-depth pass without reordering inside either subset.
Because `.16a` is alpha-composited, duplicates and this order are visible rather
than set-equivalent geometry.

**Confidence.** High for producer and mark-pass order, reconstructed
independently from asserted loops and stores. Exact host reproduction of x87
transcendental boundary cases remains outside the committed Go hashes.

### ANIM-046

The sender initializes `actionphase=-1`; each paced projectile-driver tick
increments it and picture 51 copies it raw. For phases 0–7 the draw forces sheet
frame 8 and adds `4*phase-28` to y, giving offsets
`-28,-24,-20,-16,-12,-8,-4,0`. For phases 8–15 it uses frame `phase-8` at zero
offset. The object ends on frame 7 and never wraps to frame 8 or the beginning.
Its centred anchor is `x=projectile[+0x50]-Width/2`,
`y=projectile[+0x54]-Height/2-projectile[+0x10]-projectile[+0x68]+yOffset`. At
the default clock the 16 ticks last one second.

**Confidence.** High. Sender initialization, driver increment/store, phase fork
and both frame formulas are byte-asserted on all three roots; all sixteen rows
are enumerated in evidence and independently reviewed.

### ANIM-047

The relevant painter order is: earlier ordinary/alternate CUnit cell
compositions and retained-area selector1; selector3 (+0x90) shadow sweep; the
+0x9d4 collection payload body call; retained-area selector0; selector3 body
sweep; later marker/bar and shroud phases. The late split population is CAirUnit
through its registration override, not every CUnit (ANIM-CATEGORY-084,
ANIM-AIRPASS-086). Selector1 contains the wall_of_fire/wall_of_earth arms;
selector0 contains freezing_cloud/poison_cloud. Earlier ordinary CUnit bodies
are not repainted by the late body sweep. The complete projectile insertion/type
population and native overlap order remain Unknown.

**Confidence.** Medium overall. The original no-registration vtable and painter
order are confirmed, but the projectile-to-collection assignment remains a
type/lifecycle inference. The former all-unit-body population and unconditional
Meteor overlap conclusion are narrowed; no native pixel witness.

**Amended.** The all-unit-body ordering and the unconditional Meteor overlap are
withdrawn (`claims/retracted.md`); the body states the narrowed
selector3/CAirUnit population.

## Numeral merge, pacing and producers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-071 | The numeral's `+0x20` merge key is the raw campaign sub-tick counter. Matching a victim alone or sharing a 1000 ms display window does not merge two records. | High / Medium / Unknown | ● active | [EXP-0323](../experiments/EXP-0323-damage-numeral-pacing/) |
| ANIM-072 | `R0509` uses the caller's low argument byte as a wanted opcode: the shared post-dispatch tail returns 1 when the handled message matches. It is not merely an empty-queue retry flag. | High / Unknown | ● active | [EXP-0323](../experiments/EXP-0323-damage-numeral-pacing/) |
| ANIM-073 | The nominal numeral tick-life table is `floor(1000/dtMs)` through that value plus one; it is not a realised bound on a stalled machine. | High / Unknown | ● active | [EXP-0323](../experiments/EXP-0323-damage-numeral-pacing/) |
| ANIM-074 | The general effect applier `R0565` notifies damage in its Token-8 branch only when computed damage is nonzero. | High / Unknown | ● active (amended) | [EXP-0323](../experiments/EXP-0323-damage-numeral-pacing/) · [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) |
| ANIM-075 | The five numeral-producing calls `ANIM-BLOW-019` counts are the complete population for direct calls; every other health-changing path this ledger's own claims identify is cross-referenced against that closed census, not re-read. | High / Medium | ● active (amended, partially retracted) | [EXP-0323](../experiments/EXP-0323-damage-numeral-pacing/) · [EXP-0332](../experiments/EXP-0332-area-damage-overlap/) |

### ANIM-071

In `R0566`, the campaign-object fetch established by `SESS-OBJ-001` is
followed by `L02938`, which reads the 32-bit field `+0x3e0` of that object, and `L02939`, which stores it at `+0x20` of the new record.
`SESS-TICK-026` identifies that field as the sub-tick counter mirrored by opcode
`0x64`. The intervening `L02940`, a load of the stack argument at `+0x8`, leaves the counter unchanged:
there is no shift, excluding the rival that the stored key is the `>>4`
full-tick derivative and permits a sixteen-sub-tick merge window. A separate
`timeGetTime` call writes birth at `+0x14`; it does not supply `+0x20`. The
found branch of `R0568` compares victim `+0x1c` at `L02941` and
sub-tick `+0x20` at `L02942`, both exact dword comparisons. Only identical
keys merge. `L02669`..`L02670` adds damage into `+0x04`;
`L02943`..`L02944` immediately re-formats the text through `R0567`.
No store in that branch refreshes birth, drift, colour or either key, excluding
the rival that merging restarts lifetime or position. The remaining branch is
captured through `L02945`: it appends a full `0x24`-byte record, with array
growth capped at `0x400`. `ANIM-072` establishes the normal paced ordering:
simulation, pump until a `0x64` is handled, then `0x401`. This is a key equality
rule, not a measured wall-time merge bound under delayed delivery This claim
extends `ANIM-NUM-020`.

**Confidence.** High for key identity, exact comparisons and immediate
formatting: named instructions and the full merge routine exclude the full-tick
and refresh rivals. **Medium** for design intent; **Unknown** for transport
batching and correspondence between a received broadcast and the latest server
tick, since `R0611`/`R0562` are unread

### ANIM-072

Entry `L02422`..`L02522` pops from `L00625` with `R0513` and
dispatches on the popped message's `+0x9` byte. At `L02770`..`L02946`, the
shared tail reloads that byte and compares it with the low byte of the first stack argument (`+0x8`); a match
returns 1 (`L02947`..`L02948`), a mismatch re-polls
(`L02949`..`L02950`). The `0x64` arm writes `campaign+0x3e0` and jumps to
that comparison (`L02951`..`L02952`). This reached argument read excludes
the bare-retry-flag rival. A whole-range frame-displacement sweep also finds the
empty-queue read at `L02953`: that path re-polls while the low argument byte
is nonzero or `campaign+0x3dc & 1` is set, otherwise returns 0. **Not every arm
reaches the match tail:** the conditional branch at `L02954` to `L02955` (taken on zero) drains/discards the
remaining queue and explicitly returns 0 at `L02956`..`L02957`. The ten
direct call sites comprise eight that supply 0x64 and two that supply 0; four
callers test the return, including `L02958`, which tests the return value for zero. The normal paced arm
orders simulation (`L02626`), the call supplying 0x64 and the pump (`L02959`..`L02960`),
then a separate virtual `0x401` post (`L02627`..`L02628`). Normal matching
completion therefore applies a sub-tick broadcast before presentation; the
exceptional return and transport freshness are separate boundaries. In the
paused idle path, the call supplying 0 and the pump at `L02961`..`L02962` precedes a distinct
guarded `0x402` virtual post at `L02963`..`L02964`. `SESS-IDLE-007` was
correct to name both operations; the five idle outcomes are specified by
`SESS-IDLE-019`, with the paused command work qualified by `SESS-PAUSE-020`

**Confidence.** High for the two argument reads, reached `0x64` comparison,
bounded direct-call census and separate window-message posts: reproduced static
disassembly excludes the retry-only rival. **Unknown** for server-broadcast
transport through `R0611`/`R0562`, which are not read here;
neither the caller census nor this tail capture proves every dispatch arm's
behaviour

### ANIM-073

Applying `ANIM-NUM-020`'s 1000 ms wall-clock expiry cutoff to `ANIM-PACE-017`'s
nine `dtMs` values `{125,100,83,71,62,50,41,35,31}` gives nominal counts
`{8..9,10..11,12..13,14..15,16..17,20..21,24..25,28..29,32..33}`. The arithmetic
is reproduced in `q3-tick-life-table.tsv`; the extra tick allows different
birth/tick alignment. Drift stays `-2` per tick at every speed, so slower
nominal cadence gives fewer ticks and less total drift, excluding a
speed-independent tick-count or total-drift reading of the wall-clock lifetime.
`ANIM-PACE-017`'s catch-up loop repeats while overdue and discards lateness
every sixteen ticks: actual ticks in a wall-time window can fall below or rise
above the nominal counts. `ANIM-NUM-020`'s aside "a third as far" names no
reference speed. The nominal lower counts give slowest/default = 1:2,
slowest/fastest = 1:4 and slowest/24-tps = 1:3. The aside is unscoped, not
disproved, and no correction to that row is asserted. The exact merge-key
equality is separately established by `ANIM-071` This claim scopes the
reference-speed comparison; it does not overturn `ANIM-NUM-020`.

**Confidence.** High for the nominal arithmetic from two published constants and
for fixed per-tick drift excluding speed-independent total drift. **Unknown**
for realised timing, catch-up and rendering on a running original; no stopwatch
observation was made

### ANIM-074

The whole routine `R0565`..`L02965` dispatches on Effect Token `+0x0c`: 8
at `L02966`, 12 at `L02967`, 17 at `L02968`, otherwise general apply
through a virtual call of the slot at offset 0x48 with the argument 1 at `L02969`. `ITEM-EFFDISP-075` states
that the equipment grammar makes Token 0, not 8/12/17. This capture did not
identify an item equip/use producer. `MAGIC-POISONINPUT-157` now identifies the
Poison Cloud arm as a spell producer of Token 8; it does not establish a
separate item equip/use producer. Its damage arm starts with `-magnitude` from
`this+0x40` and, for nonzero `target+0xc6`, scales by `(100 - target+0xc6)/100`
and adds `0.5` before truncating conversion (`L02970`..`L02971`). The
captured PE constants are `L02972 = 100.0` and `L02973 = 0.5`.
a zero-result test (`L02974`/`L02975`: the local at frame offset −0x4 against zero, branching to `L02976` when equal) skips both the health write
and notification; `L02976` jumps to `L02977`, which reaches the epilogue. This real
zero-result branch excludes unconditional notification once magnitude is
computed. Otherwise it writes `target+0x94 -= computed`, runs the
Building-repair check, and calls `R0561(target, 0)` with L00522 as the receiver
at `L02978`..`L02979`. The same zero `dmg` argument agrees with
`ANIM-BLOW-019`'s builder reading. `HERO-HEALTH-032` names this writer without
reading its gate; the retained `R0565` corpus search scopes that novelty This
claim closes the `R0565` clause of Open item 10.

**Confidence.** High for Token dispatch, resistance constants, zero gate and
notify arguments: the full routine and PE constant reads exclude the
unconditional-notify rival. **Unknown** for a separate item equip/use producer;
the Poison spell producer is joined by `MAGIC-POISONINPUT-157` and
`MAGIC-POISONREFRESH-158`

**Amended.** Poison producer amended (`claims/retracted.md`):
`MAGIC-POISONINPUT-157` identifies the Poison Cloud arm as a spell producer of
Token 8, and an item equip/use producer remains Unknown.

### ANIM-075

**The complete population of numeral-producing calls is the five `ANIM-BLOW-019`
already counts, full stop for direct calls; every other health-changing path
this ledger's own claims name is cross-referenced against that closed census,
not re-read.** `ANIM-BLOW-019`'s `EnumRefs callto:R0561` (5 hits / 4 owners / 0
orphan) is a complete direct-call census: melee/general strike `R0246`,
structure strike `R0563`, area `R0564` (twice) and, by `ANIM-074`,
the general effect applier's Token-8 branch `R0565` (a general
dispatcher, not an item-specific one; `MAGIC-POISONINPUT-157` now identifies the
Poison spell producer). **Melee and missile are the same producer, delayed, not
two paths**: `ANIM-MSG-005` names `R0549` — called only from swing-start
`R0245` — as the sole chooser between opcode `0x71` (melee) and `0x72`
(missile); `UNIT-STRUCTDELIVERY-065` shows both "remain in action 3's countdown
and reach `R0001` at phase-5 expiry" regardless of which opcode was sent,
and `UNIT-STRUCTDAMAGE-064` shows that same `R0001` branches only on the
TARGET's own vtable — unit to `R0246`, Building to `R0563` — both
numeral producers, so a missile hit produces the identical numeral a melee hit
does, later by the projectile's own travel countdown, and a Building under
physical attack gets one too. **Wall_of_fire and poison_cloud share the same
outer dispatch and differ only in inner class**: `MAGIC-AREATICK-036`'s own
corpus count places both in the six "cloud"-mode `AreaEffect` spells, so their
per-tick call into `R0267`'s target apply is identical. `MAGIC-ARM-014`
shows `wall_of_fire` is one of "the seven pure damage spells" whose kind-parsed
Effect-building arm never executes because it carries positive scratch damage,
and `UNIT-AREADIRECT-072` independently reads that same fork in `R0003`:
"positive scratch damage takes its generic damage arm" — the arm that installs
`Effect_DirectDamage`'s vtable (`L02980`, `+0x3c = R0564`) rather than
base-`Effect`'s (`L02981`, `+0x3c = R0612`). `poison_cloud`'s `Effects`
column ("health=-2") is instead parsed by the kind-lookup arm
(`MAGIC-EFFECT-015`), which builds the base-`Effect` object. The former
inference from `R0612`'s absence in the direct-call census to Poison
changing health with NO numeral is retracted. First attachment and periodic
ticking invoke base Effect virtual `+0x40`, reaching `R0565`'s Token-8
branch; nonzero computed damage writes HP and calls `R0561`, while zero
skips both (`ANIM-074`, `MAGIC-POISONINPUT-157`, `MAGIC-POISONREFRESH-158`,
`MAGIC-POISONPHASE-159`). This establishes the local notification call, not a
drawable numeral or client pacing. **What this claim does not establish**: no
claim, including this one, reads the instruction that writes `wall_of_fire`'s
own scratch-damage columns as positive — `MAGIC-ARM-014`'s classification of it
as one of the seven is itself the evidence; and the same reasoning was not
extended to `fire_arrow`, `lightning` or `prismatic_spray`, three of the seven
that are not among `MAGIC-AREATICK-036`'s ten `AreaEffect` spells and so never
reach `R0267` through this route at all — how their damage is delivered
is untraced here. **The broader census**: `HERO-ZERO-070`'s 36-writer
attribution of `+0x94` names heal, removal from the map, constructor, template,
load, script action and the clock (regen/decay) alongside "a blow"/"an effect"
as causes of a health change; none of those other categories' write sites is
among `ANIM-BLOW-019`'s four owners, so by the same closed census they change
health with no numeral

**Confidence.** High that the five-site, four-owner census is exhaustive over
direct calls (`ANIM-BLOW-019`'s own enumeration, `0` orphan) and that
melee/missile/structure share one producer (three claims' own instructions,
cross-read rather than re-disassembled) / High that wall_of_fire's
positive-damage classification routes it through `UNIT-AREADIRECT-072`'s own
"generic damage arm"/`DirectDamage` fork, agreeing across two independently
authored claims / **Medium** overall: neither claim names wall_of_fire's own
cast-time instruction by address, the three non-`AreaEffect` pure-damage spells
are outside this cross-reference entirely, and an indirect call to
`R0561` outside the direct-call census would not be seen by it

**Amended.** Poison notification inference amended (`claims/retracted.md`): the
intervening virtual `+0x40` application reaches the Token-8 body, which calls
`R0561` when computed damage is nonzero. The inference that periodic
reapplication changes health with no numeral is withdrawn; notification is
shown, a visible numeral is not. The five-site direct-call census stands.

## Drawable registration and cell composition

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-REGISTER-083 | (rom.exe) Registration selectors 0,1,2,3,4 map to view fields +0x98,+0x94,+0x8c,+0x90,+0x9c respectively. | High | ✔ promoted | [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |
| ANIM-CATEGORY-084 | (rom.exe) The category join uses actual runtime-class tables, not the informal word unit. | High / Unknown | ✔ promoted | [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |
| ANIM-CELL-085 | (rom.exe) Ordinary and alternate CUnit populations participate in different per-cell compositions. | High | ✔ promoted | [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |
| ANIM-AIRPASS-086 | (rom.exe) The separated late shadow and body sweeps read selector3 storage (+0x90), reached by CAirUnit registration. | High / Medium | ✔ promoted | [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |
| ANIM-DRAWGATE-087 | (rom.exe) Both structure body phases accept signed drawable+0x78<2. | High | ✔ promoted | [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |
| ANIM-WALKORDER-088 | (rom.exe) The painter composition has nine cell sweeps and one intervening non-cell collection walk. | High | ✔ promoted | [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |

### ANIM-REGISTER-083

The five-entry switch in `R0613` reads `L02982`. Registration first
requires drawable+0x4c nonzero. Selector1 writes the supplied pointer over the
inclusive prepared +0x3c/+0x40..+0x44/+0x48 rectangle; the other four select one
anchor using virtual dimensions and +0x28/+0x2c, bounded to columns
[-3,visibleCols+3) and rows [-3,visibleRows+7). Index is
`(row+3)*(visibleCols+6)+col+3`. `R0614` clips the prepared rectangle. A cell
stores one pointer, so a later write replaces it. The out-of-table arm uses the
selector argument as a pointer, not a validated default plane. Native refresh
order, competing last writer and stale-entry lifetime remain Unknown.

**Confidence.** High for the complete switch, store/clip instructions and
conditional instruction vectors. No native lifetime or malformed-memory safety
claim.

### ANIM-CATEGORY-084

CBackPack calls registration selector0; CStructure and the three bridge tables
call selector1. CUnit calls selector2 only while unsigned byte +0x15a<2 and
(+0x18c&0x80)==0, otherwise selector4. CAirUnit calls selector3 without either
test. Thus a stage1 CUnit can remain in selector2 and a stage0 CUnit can use
selector4. The CUnit/CAirUnit body and shadow slots share targets
R0552/R0553 while their registration slots differ. CProjectile
registration is a bare return. CAirUnit construction also writes +0x10=16; this
is separate from the registry Z value. The broader server-Sack/client-CBackPack
message join and native reachability of synthetic stage/flag combinations are
not established here.

**Confidence.** High for the joined vtable records and complete producer bodies,
with 36 CUnit/CAirUnit branch vectors. Unknown for untraced lifecycle
populations.

### ANIM-CELL-085

In the earlier cell sweep, selector4 shadow/body dispatches at L02983/L02984
precede CBackPack shadow/body at L02985/L02986 within that cell. A later
cell sweep dispatches non-flat structure body at L02987, selector2 shadow/body
at L02988/L02989, an optional auxiliary +0x18 dispatch at L02990, the
map-byte static-object path, then retained-area selector1 at L02991. Static
objects use the map-byte/class-array path rather than a joined drawable
registration selector; their cull at L02992 can skip both shadow and body.
Presence/visibility/shadow gates apply independently. These are call-order
facts, not an unconditional pixel-overlap result.

**Confidence.** High for complete painter control flow, actual loads/dispatches
and the selector4/CBackPack gate vectors. Auxiliary marker meaning, pixel
coverage and unrelated helper side effects are not promoted.

### ANIM-AIRPASS-086

The complete shadow sweep calls L02993; the non-cell +0x9d4 collection walk
calls L02994; retained-area selector0 runs at L02995; the complete body
sweep calls L02996. Retained-area selector1 already ran inside the earlier
main cell sweep. These area selectors belong to R0607 and are a different
domain from registration selectors. Ordinary selector2 and alternate selector4
CUnit body calls occur earlier, so the late body phase cannot be called all unit
bodies. The projectile-to-collection interpretation of amended ANIM-047 remains
Medium: CProjectile has no grid registration, but complete insertion/type
population is unjoined.

**Confidence.** High for the selector/storage/control-flow join; Medium for the
inherited projectile-list type interpretation. No native pixel or
insertion-order witness.

### ANIM-DRAWGATE-087

Their Flat predicates are opposite: nonzero early, zero in the main cell sweep.
The main body is called at L02987 before the later +0x78==0 test guarding
rectangle/mark work; that later test is not a body gate. The complete CStructure
body contains no +0x78 test. Selector2/3/4 and CBackPack body paths require
+0x78==0. Selector2 additionally queries key0x26 through R0615: when found,
the tested indexed 16-bit mask must contain bit8. This does not assign a new
meaning to that key. Shadow calls have their separate global enable gate; it can
skip the whole selector3 shadow sweep without skipping the body sweep.

**Confidence.** High for the inspected local predicates and 36 conditional
painter vectors. No assertion that malformed signed fog values or every tested
lifecycle combination occurs natively.

### ANIM-WALKORDER-088

Retain phase labels1..10 with phase6 denoting the collection, not a cell sweep.
Phases1..5,7..8 iterate rows -4 through visibleRows+7 ascending and columns
visibleCols+3 through -4 descending. The phase9 marker/bar walk is narrower:
rows -3 through visibleRows+6 and columns visibleCols+2 through -3. Phase10
shroud is column-major: columns 0..<visibleCols and rows 0..<visibleRows+4,
ascending. Within a drawable cell sweep, later cells have greater row or equal
row and smaller column; across phases, the earlier phase finishes before the
later one regardless of coordinates. Dispatch guards further reduce the visited
population; visit order alone does not guarantee opaque pixel overpainting.

**Confidence.** High for the complete painter branch/initializer/index operands.
Decoder instruction totals are instrument measurements, not game properties; no
runtime pixel equivalence claim.

## Picture-7 coordinate callback

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-097 | In the inspected picture-7 callback, each admitted pixel is a whole 16-bit word copied in place from the current buffer through a signed coordinate map. | High | ✔ promoted | [EXP-0416](../experiments/EXP-0416-projectile-callback/) |
| ANIM-098 | The visited initializer constructs picture 7's shared coordinate map with 20 and 16, producing 19 by 19 source pairs and write offsets -18..18 around its centre. | High | ✔ promoted | [EXP-0416](../experiments/EXP-0416-projectile-callback/) |
| ANIM-099 | The inspected picture-7 callback clips each copy's destination and source separately against its captured active rectangle; rejected copies preserve the destination. | High | ✔ promoted | [EXP-0416](../experiments/EXP-0416-projectile-callback/) |
| ANIM-100 | The reached picture-7 draw passes a shared map and projected centre coordinates, bypassing the computed sprite frame, facing and registry-centred extents. | High | ✔ promoted | [EXP-0416](../experiments/EXP-0416-projectile-callback/) |

### ANIM-097

Callback R0605 reads the buffer base at L01168 and its stride in words
at L01169. Receiver+0 contains row pointers; each entry is a pair of signed
16-bit source offsets. For indices i and j descending from receiver[+0c]-1
to 0, it copies the four reflected offsets in this order: (+i,+j), (+i,-j),
(-i,+j), (-i,-j), from the corresponding signs of (sx,sy).

Each copy reads and stores two bytes in the same buffer. It preserves the
entire word, with no channel conversion or blending in the complete bounded
body. It has no sprite or palette input. Axes and the centre are copied more
than once. No source snapshot is created by this callback.

**Confidence.** High for the bounded body: 195 raw instructions and four word
stores, confirmed by independent decoder boundaries and original-instruction
replay. The live map fits both snapshot and in-place models. A synthetic
overlap map distinguishes them by two final words; only the in-place model
matches the original instruction trace.

**Unknown.** The actual buffer RGB masks, previous-frame retention and native
compositing remain unobserved. Copying a packed word does not establish them.

### ANIM-098

Initializer R0616 calls constructor L02997 with A=20 and B=16. It stores
the returned 16-byte map object at L02998. Its +4/+8 fields hold A/B; +0c
holds N=trunc(sqrt(A*A-(A-B)*(A-B)))=19.

Builder L02999 defines f[0]=1 and f[k]=k/(20*sin(atan(k/4))) for k=1..18.
For i,j=0..18, k=trunc(sqrt(i*i+j*j)+0.5). If k<19 the source pair is
(trunc(i*f[k]+0.5),trunc(j*f[k]+0.5)); otherwise it is (i,j). The original
R0279 conversion helper temporarily selects truncation toward zero and
restores its incoming x87 control word.

The live callback writes a 37 by 37 square, with identity copies in its outer
corners. A unique-word 75 by 61 synthetic background at centre (37,30) receives
1444 copies at 1369 distinct words; 1084 final words change. A constant word
background has no changed words. Reapplying to the first output can change it
again; the callback is not generally idempotent.

**Confidence.** High for the visited construction and reached 20/16 mapping.
Original x87 replay and an independently expressed model agree on all 361
entries and five smaller valid synthetic parameter pairs with control word
037f. The square-root simplification is a real-arithmetic identity, not a claim
of every transcendental implementation's identical rounding.

**Unknown.** Other precision modes, invalid/custom constructor parameters,
initialization failures and indirect changes to the shared map remain open.

### ANIM-099

The callback snapshots the RECT at L01503 through R0372/CopyRect. Every
quadrant first calls imported PtInRect for the destination; on success it
calls the same API for the source, then copies only if both succeed. It neither
clamps the source nor rejects an entire straddling effect.

The rectangle contract includes left/top and excludes right/bottom. Empty and
inverted rectangles reject the selected point population. A synthetic admitted
destination whose two reflected sources lie outside the rectangle produces no
write. Four edges, one corner, an inset rectangle, an outside centre and an
empty rectangle preserve all guards and match the per-copy model.

**Confidence.** High for the complete local two-test/store gates. Original
callback replay bridges the imported APIs to native host user32 on synthetic
structs; 20 separate rectangle cases check the boundary contract. This is a
static/synthetic result rather than an original game's native rendering witness.

**Unknown.** The actual native rectangle/surface supplied during play and the
original 32-bit OS linkage are not observed.

### ANIM-100

Draw R0556 is projectile vtable+28 at L02870. The picture-7 switch arm
calls R0605 on the object at [L02998] (the `this` argument), x=P[+50], and
y=P[+54]-P[+10]-P[+68], using 32-bit arithmetic. The callback pops 8 bytes
on return. The registry count and nonnull-slot guards still apply.

The arm reloads stored coordinates after the ordinary width/height halves,
direction fold and frame computation. It supplies no phase or draw-entry
argument. The complete callback reads no clock and changes no map entry.
Position, height, the underlying buffer and clipping remain output inputs;
unchanged centre alone does not imply unchanged pixels.

**Confidence.** High for the bounded handoff and body. Six synthetic prefix
cases per locale independently vary spatial fields and phase/facing/registry
extents, and exercise null/count refusals. EN/RU images are byte-identical.

**Unknown.** The complete invocation schedule and indirect map mutations remain
open. Three classified whole-image absolute references to the shared pointer
do not establish that aliases or bulk writes cannot change its object.

## Physical shot flight, direction and the Fire_Ball burst arm

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-101 | The projectile driver's default arm moves a record each call by the truncated quotient (action - position) / actionsegments per axis; stored arrow, bolt and rock records from three original saves reproduce exactly. | High | ● active | [EXP-0430](../experiments/EXP-0430-physical-projectiles/) |
| ANIM-102 | `L02828` turns a vector into a 16-way direction (0 north, 4 east, 8 south, 12 west, y down) from slope bins 1/4, 3/4, 4/3 and 4 in each quadrant; the saved actiondir of three records agrees. | High | ● active | [EXP-0430](../experiments/EXP-0430-physical-projectiles/) |
| ANIM-103 | Picture 13 (`fireexpl`) draws frame (actionphase / 2) mod 11 for 22 calls; its driver arm acts when actionphase is 4 (a 3x3 cell registration) and 8 (an OR of tile-word bit 0x2000 over the 3x3 block). | High / Medium | ● active (amended) | [EXP-0430](../experiments/EXP-0430-physical-projectiles/) |
| ANIM-109 | Projectile records of picture 14 or more come only from the cast spawner (even picture 2*spell+8), client arm 0x86 (odd: 17, 27, 51), arm 0x8b, arm 0x8c (picture 36) and the loader; no shipped unit class has Projectile above 12. | High / Medium | ● active | [EXP-0432](../experiments/EXP-0432-projectile-pictures/) |
| ANIM-110 | The only immediate stores of opcode 0x8b and 0x8c are in `R0617`/`R0618` and `R0619` (spell 14, picture 36), reached via `R0268` from the actor tick `R0037`; two register-form writers stay unresolved. | High / Medium | ● active (amended) | [EXP-0432](../experiments/EXP-0432-projectile-pictures/) |
| ANIM-111 | Every Fire_Ball area effect, from any actor or the siege rider, is built by `R0003`, waits in a SpellTransport of distance/384 ticks and ends in one blast-arm 0x86 (picture 13, 22 calls); shipped data gives no second 0x86. | High / Medium | ● active | [EXP-0432](../experiments/EXP-0432-projectile-pictures/) |
| ANIM-112 | An area effect stores spell id at +0x0c and picture 2*spell+9 at +0x0e; in shipped data only spells 2, 4, 9 and 21 send that picture in an 0x86 (13, 17, 27, 51), the other area-effect spells send 0x87 masks. | High | ● active | [EXP-0432](../experiments/EXP-0432-projectile-pictures/) |
| ANIM-113 | In `R0260` and `R0454` one timer step runs the simulation step, the client dispatcher, then the 0x401 pass, which sweeps the actor map (CUnit driver `R0548`) before the record map (CProjectile driver `R0558`). | High / Medium | ● active | [EXP-0441](../experiments/EXP-0441-shot-timing/) |
| ANIM-114 | A Fire_Ball burst record is built by client arm 0x86 from a message the inner effect's first tick sends; an effect appended in the walk's last visit waits for the next walk, and same-step delivery is not traced. | High / Medium | ● active (amended) | [EXP-0441](../experiments/EXP-0441-shot-timing/) |
| ANIM-115 | A catapult or ballista Fire_Ball rider builds one area effect and one SpellTransport and appends only the transport; the damage message 0x73 follows, and the burst 0x86 comes later from the inner effect. | High / Medium | ● active (amended) | [EXP-0441](../experiments/EXP-0441-shot-timing/) |
| ANIM-121 | On the local peer path the flush `R0611` appends the packets to the client endpoint's receive list in the same call, so the client dispatcher of the same timer step dequeues them: the flush-to-dequeue delay is 0 ticks. | High / Medium | ● active | [EXP-0453](../experiments/EXP-0453-shot-burst/) |
| ANIM-122 | No routine read builds a world record for a projectile picture of 14 or above except the cast spawner, the loader and client arms 0x86, 0x8b and 0x8c; the writers `ANIM-110` left open write no projectile opcode. | Medium / Unknown | ● active | [EXP-0453](../experiments/EXP-0453-shot-burst/) |
| ANIM-123 | The Fire_Ball burst message carries picture 13, the cell, and a segment word of 22, and the client builds a 0x14c-byte record from them with no size field: its footprint registration is a constant 3 x 3. | High / Medium | ● active | [EXP-0453](../experiments/EXP-0453-shot-burst/) |
| ANIM-124 | Client arm 0x73 applies the damage message when it is dispatched: it sets the target's hit points from the message and reads no burst or projectile state, so a pending burst does not delay or queue it. | High / Medium | ● active | [EXP-0453](../experiments/EXP-0453-shot-burst/) |
| ANIM-125 | The server builds damage message 0x73 only for an actor target; of its five senders, only the unit strike sends it for zero damage, so hurt event 0 comes only from a zero-damage unit strike leaving the target above -10. | High / Medium | ● active | [EXP-0463](../experiments/EXP-0463-hurt-voice-event/) |
| ANIM-126 | Three routines were identified as sending state syncs that resolve to hurt event 3 or 2: the death arm, the negative-health bleed and the session-entry projection; the other 48 projector sites were classified by mask only. | High / Medium | ● active (amended) | [EXP-0463](../experiments/EXP-0463-hurt-voice-event/) |
| ANIM-127 | For a person class below 0x1a, the setter's stack value is server actor+0x4b: spawner bytes when the placement's secondary key word is non-zero, otherwise the Humans-row constructor value; its bit 0x80 is the sex bit. | High / Medium | ● active (amended) | [EXP-0463](../experiments/EXP-0463-hurt-voice-event/) |
| ANIM-128 | [L02739] is the speech volume attenuation of the sound configuration object (default -700, saved as SoundSpeechPos); it does not combine with the effects attenuation [L02733], and each is added once to the distance term. | High / Medium | ● active | [EXP-0463](../experiments/EXP-0463-hurt-voice-event/) |

### ANIM-101

- Each call of `R0558` with segments non-zero and action 1 does the following (listing in `evidence/disasm-driver-entry-and-arm-0d.txt`). If ActionTarget is non-zero and found in the unit hash, ActionX, ActionY and ActionZ take the target's `+0x58`, `+0x5c` and `+0x10`, and actiondir takes `L02828`'s result. Three `IDIV` by actionsegments give the per-axis step `(action - position) / segments`, truncated toward zero.
- `dir` takes actiondir, actionphase gains 1 and the default arm adds the three steps, sets `phase = (actionphase / 2) mod Phases` (`ANIM-PHASECLOCK-028`), sets lastaction to action and subtracts 1 from actionsegments.
- The order is: direction from the position before the step, then the step, then the counters.
- Check (`evidence/flight-model.txt`, `evidence/prj-leaf-values.txt`), target stationary in all three saves: starting at the `SAV-1142` positions with the target at (6016, 24448), 6 calls of Prj27, 5 of Prj28 and 3 of Prj30 reproduce the saved x, y, dir, phase and actionsegments exactly: (6016, 24448), (5764, 24666), (5888, 25162).
- Pictures 1, 2 and 5 have Phases 1 in projectiles.reg (`evidence/projectiles-reg-rows.txt`), so phase stays 0, as saved.

**Confidence.** High: the arithmetic is instructions and three records agree on every moving leaf. The target was stationary in all three saves, so the per-call re-read of its position is not exercised by them.

**Unknown.** A shot at a moving target in a save.

### ANIM-102

- With a = |dx| and c = |dy|, where dx and dy are action minus position: bin b is 0 if a >= 4c; else 1 if 3a >= 4c; else 4 if c >= 4a; else 2 if 3c < 4a, 3 otherwise.
- The result is (4 - b) & 15 for dx >= 0 and dy <= 0, (b + 4) & 15 for dx > 0 and dy > 0, (12 - b) & 15 for dx <= 0 and dy > 0, and (b + 12) & 15 for dx < 0 and dy <= 0. A zero vector gives 4.
- The three saved records end on bins 3, 2 and 4 (Prj27, Prj28, Prj30), and the saved actiondir 1, 2 and 0 agree. Prj29 keeps its constructed 0 because ActionTarget 0 skips the call.

**Confidence.** High for the function (a full listing, `evidence/disasm-direction-L02828.txt`) and for the three agreements. Bins 0 and 1 are not exercised by a record.

### ANIM-103

- The driver's picture switch (`L02879`, `L03000`) gives picture 13 alone the arm `L03001` (`evidence/driver-tables.txt`). Every other picture takes another arm or the default.
- Arm 0x0d acts only on the call whose incremented actionphase is 4 or 8. Other calls fall to the shared phase store.
- At 4, a cell of the 3x3 block around the record's cell (`+0x34`, `+0x38`) whose tile word has bit 0x2000 is skipped. Of the others, a cell whose second-plane byte is non-zero and whose terrain-type entry `+0x44` is not -1 is added to the hash at `+0xa7c` of the record's `+0xe0` object, key `(y << 8) | x`; a cell whose tile type field (bits 6 to 12) is 8 to 11 is added likewise.
- At 8, each of the 9 tile words gets bit 0x2000, except where the type field is 8 to 11 and the plane byte is 0.
- Frame: projectiles.reg ID 13 has Phases 11 and RotationPhases 1, so `frame = (actionphase / 2) mod 11` is one pass of 11 frames over 22 calls (`ANIM-PHASECLOCK-028`, `SAV-1144`). `MAGIC-BURSTLIFE-034` gives the sound.
- The saved Prj29 has actionphase 7: the call-4 registration has run and the call-8 write has not. The frame law has one sample (actionphase 7, phase 3), which several rival laws fit; it rests on the code reading and `ANIM-PHASECLOCK-028`, not on the save.
- The bit and the `+0xa7c` hash are the ones `MAGIC-WALLFIRE-058` records for `wall_of_fire`.

**Confidence.** High for the tick numbers, the 3x3 extent, the bit and the frame law, read from the listing. Medium for the cell predicates, whose table at `L02099` and plane meaning were not decoded.

**Unknown.** Whether bit 0x2000 or the hash entries survive into a SAV: the terrain cells around (23, 95) in a later save were not compared. Next question: those cells in a save taken after tick 8 of a burst.

**Amended.** The persistence Unknown is answered in part by `SAV-1150` and `SAV-1151`: the Fog store writes bit 15 only, no other store of the tile plane into a SAV was found, and no SAV store or load routine addresses the `+0xa7c` hash. Whether a load clears the bit is Medium (`SAV-1150`), and a saved-cell comparison does not apply since a SAV carries no tile plane beyond the Fog runs.

### ANIM-109

- Constructor `R0609` has six direct call sites: client arms 0x86, 0x8b, 0x8c (`L03002`, `L03003`, `L03004`), unit shot `R0603` (`L03005`), cast spawner `R0620` (`L03006`) and the loader (`L03007`). The copy constructor `L03008` has one, the spawner's second record for picture 60 (`L03009`). Seven `push 0x14c` allocations match (`evidence/xrefs.txt`, `evidence/imm-14c-record-size.txt`). Neither address appears in a table or as an immediate.
- Unit shot: the picture is the class Projectile. The 266 unit rows of `EXP-0428` `shot-classes.csv` that carry a value hold 0 to 7, 10 and 12, so no shipped unit shot has picture 14 or more.
- Cast spawner: the picture is `actionspell`, which the cast message sets to `2*spell+8` (`MAGIC-CASTSPAWN-033`). Pictures 14 to 64 take the arm `L03010` (0 segments) except 20 and 30 (1), 34 and 36 (13) and 60 (21) (`evidence/spawner-driver-tables.txt`). The record starts at actionphase 0 (`SAV-1147`).
- Client arm 0x86 builds a record only for an odd picture with a loaded registry slot (`L03011`, `L03012`); an even picture sets the caster's action (`L03013` to `L03014`, `evidence/disasm-client-arm-86-even-path.txt`). The odd pictures of 14 or more that a sender writes are 17, 27 and 51 (`ANIM-112`); each has a registry row (`evidence/spell-pictures.txt`).
- Client arm 0x8b builds a record of the picture in `msg+0xc` with no parity guard. Leaves: x and y are `msg+0xa` and `msg+0xb` times 256 plus 128; ActionTarget is the word at `msg+0xd` when the registry slot's `+0x2c` is non-zero, otherwise ActionX and ActionY come from `msg+0xd` and `msg+0xe` the same way; actionsegments is the word at `msg+0xf`; actionphase is -1 (`L03015` to `L03016`). The senders write `2*spell+8` (`ANIM-110`).
- Client arm 0x8c builds picture 36 only (`L03017`): x and y from the two bytes of `msg+0xe` (`L03018` to `L03019`), actionsegments 13 (`L03020`), actionphase -1, ActionTarget the word at `msg+0x10`, and a list of the words from `msg+0x10` on, counted by the dword at `msg+0xa`.
- The loader restores records and is not a picture source.

**Confidence.** High for the site list: a raw `rel32` scan (`evidence/xrefs.txt`) and a linear sweep of `.text` (`evidence/sweep-constructor-calls.txt`) agree, and both constructors are referenced by direct calls only. High for the unit class bound over the 266 rows. Medium for the odd pictures 17, 27 and 51 as the only ones an 0x86 sender writes: they take the grade of `ANIM-110` and `ANIM-112`. The spawner is reached through its vtable slot only (`MAGIC-CASTSPAWN-033`).

**Unknown.** Arm 0x8a has no constructor site and carries no picture; what it does for picture 36 was not read. Pictures above 64, or a registry with other rows (customised data), are outside the shipped population.

### ANIM-110

- Opcode stores: A search for immediate byte stores at offset +0x9 finds 0x8b at `L03021` and `L03022` and 0x8c at `L03023`, each once (`evidence/opcode-stores.txt`, 59 stores in `.text`: 25 immediate, 34 register form). Their routines are `R0617` and `R0618` (0x8b) and `R0619` (0x8c).
- Callers by raw `rel32` scan (`evidence/xrefs.txt`): `R0617` at `L03024` and `R0618` at `L03025`, both inside the spell apply `R0268`; `R0619` at `L03026`, inside `R0269`, whose one call site is `L03027` in `R0268`. Spell id 14 takes that branch (`L03028`). The other spells reach `L03029`, where `R0621` selects one of the two 0x8b senders.
- `R0268` has two call sites, `L01994` and `L01995`, both in the actor tick `R0037`, which is installed in three vtables (`L03030`, `L03031`, `L03032`); those `.rdata` slots are its only references (`evidence/xrefs.txt`).
- `R0617` and `R0618` write opcode 0x86 first and rewrite it to 0x8b when the source's word at `+0xe` is zero. The source's cell goes to `msg+0xa` and `msg+0xb`; `msg+0xc` is `2*spell+8`; `msg+0xd` is the target's id word (`R0617`) or its cell bytes (`R0618`); `msg+0xf` is the segment count. The count is distance/speed when spell parameter 5 is 2, 5 for spell ids 13 and 14, and 0 otherwise; `L03033` returns without a send when `R0416` is non-zero (`L03034` to `L03035`).
- `R0619` writes 0x8a, or 0x8c when the source's word at `+0xe` is zero. It writes no picture byte (`MAGIC-PIC-027`).
- Register-form writers (`evidence/callargs.txt` lists every call site with the visible pushes):
  - `R0622`: the opcode argument is a constant at each call site, 0x71 (`L03036`), 0x6d (`L03037`) and 0 (`L03038`, written as 0x6b). None is 0x86, 0x8b or 0x8c.
  - `L03039`: constants 0xb7, 0xb8, 0xaf, 0x03, 0x83. `R0623`: constants 0x0b, 0xaa, 0x84, 0xb4, 0xb5, 0x74, 0x92, 0x67, 0x97.
  - `R0624` and `R0625` copy spell parameter 6 into `msg+9` (`L03040`, `L03041`); the shipped spell rows hold 0, 1, 3, 5, 6, 7, 8 and 10 there.
  - `L02428` is the receive-side decoder keyed on its opcode argument; `R0626` sets 0x86 once in the shared buffer `L03042` (`evidence/xrefs-message-buffer-L03042.txt`).
  - `R0562` and `R0627` take the opcode as an argument whose callers pass variables (13 and 3 call sites); `R0628` passes 0x89 to the second.

**Confidence.** High for the immediate stores and the call chain: instructions in listings and scans over all of `.text`. Medium for "no other routine sends 0x8b or 0x8c": the register-form writers above are resolved except the last two, and the scan does not see a block copy or an encoding it does not decode. This answers the open item of `MAGIC-PIC-027`.

**Unknown.** Whether any call of `R0562` or `R0627` carries 0x8b or 0x8c. A customised Data.bin could put 0x8b or 0x8c into spell parameter 6 and reach `R0624` or `R0625`. The condition on the source's `+0xe` word was not decoded.

**Amended.** `EXP-0453`: the two register-form writers are read and write no projectile opcode (`ANIM-122`).

### ANIM-111

- `R0003` has three call sites: `L00008` in the actor tick `R0037`, `L03043` in `R0269` (spell 14) and `L03044` in `R0002`. `R0002` has two callers, `L00007` in the actor tick and `L03045` in the rider `R0246`; it skips spell id 14 (`evidence/xrefs.txt`, `evidence/disasm-effect-create-R0002.txt`).
- Default arm `L03046`: when spell parameter 8 is not 1 it builds an AreaEffect (`L03047`) with `+0x0c` the spell id and `+0x0e` `2*id+9` (`L03048` to `L03049`, `L03050` to `L03051`). When parameter 5 is 2 it then builds a SpellTransport around it (`L03052`) and adds the transport to the effect list (`L03053`), not the inner effect. Its tests read spell parameters 5, 8, 11 and the spell id.
- Fire_Ball is spell 2: parameter 5 is 2, parameter 7 is 384, parameter 8 is 3, parameter 11 is 0 (`evidence/databin-spell-rows-en-ru.txt`). The effect's mode word `+8` stays 0, so the tick `R0629` takes the blast arm `R0630` (`R0631`, `R0632`; `MAGIC-AREADRAW-049`).
- Transport: countdown `distance/speed` at `+0x4c` (`R0633`); each tick subtracts 1; at 0 it adds the inner effect to the list and sets its own reap flag (`R0634`).
- Blast arm: one call `R0635(effect, 1)` at `L03054`. For spell id 2 it writes opcode 0x86, picture `+0x0e` (13) and 22 calls (`L03055` to `L02927`). The arm then sets the reap flag (`L03056`).
- The cell loop calls `R0636`, which dispatches on spell id minus 2 (`evidence/layer-erase-table.txt`). Its sends pass an object found in the cell's layer slots, not the blast; such a send is an 0x86 only if that object's spell id is 2. The blast arm registers nothing in the layers, so shipped Fire_Ball sends no second 0x86. The cloud end arm `R0637` is reached only with mode bit 1 set (`L03057`, `L03058`).
- The one other caller of `R0635` for live effects is the player-join routine `R0131` (`L03059`), which walks the effect list.

**Confidence.** High that every Fire_Ball effect takes this creation path and single send: all three call sites enter one creator whose default arm tests spell parameters only. Medium for the shipped-data limit on the second send, which rests on layer registration being cloud-only (`MAGIC-AREADRAW-049`), not on a run. A Data.bin that gives Fire_Ball a duration changes `+8` and the mode.

**Unknown.** The branch tests in the actor tick that choose `L00008` or `L00007`. A Fire_Ball cast by a caster outside the classes that use `R0037`. Whether a SAV restores a live effect (the loader was not read for effects), which would be a creation path outside `R0003`.

### ANIM-112

- `R0003` stores `+0x0c = [spell+8]` (the spell id) and `+0x0e = 2*id + 9` (`L03048` to `L03049`, `L03050` to `L03051`).
- Senders: `R0638` (staged arm `R0639`, one send per accepted cell, `L03060`) writes opcode 0x86 with `msg+0xc = [effect+0xe]`; `R0635` writes 0x86 with `msg+0xc = [effect+0xe]` only when `[effect+0xc]` is 2 and writes 0x87 otherwise.
- The staged mode is `+8 = 2`, set when parameter 8 is 5 (`L03061`); the shipped spells with parameter 8 equal to 5 are 4, 9 and 21. Spell 2 is the one blast. So the odd pictures sent in an 0x86 are 13, 17, 27 and 51, and each has a projectiles.reg row, so the client builds the record (`evidence/spell-pictures.txt`).
- The other area-effect spells (3, 7, 8, 12, 17, 19: pictures 15, 23, 25, 33, 43, 47) send 0x87; pictures 33 and 43 have no registry row.
- Data.bin has 28 spell rows; the EN and RU parameter lists are compared in `evidence/databin-spell-rows-en-ru.txt`.

**Confidence.** High: the formula and both senders are instructions; the spell sets come from the parameter columns of the shipped rows (28 rows).

**Unknown.** Customised Data.bin rows move spells between the modes (G2).

### ANIM-113

- Timer handler `R0260` calls the simulation step `R0147` (`L03062`), the client dispatcher `R0509` (`L03063`) and the window's virtual slot +0x48 (`L03064`), which carries the 0x401 broadcast. `R0147` calls the sub-tick `R0193` at `L01863`. `R0454` makes the same three calls in the same order (`R0147`, `R0509`, slot +0x48). `R0455` calls `R0147` but no `R0509` and no slot +0x48; `R0640` calls `R0611`, `R0509` and slot +0x48 but not `R0147` (`evidence/calls-q1.txt`).
- Sub-tick `R0193`: message drain `R0191`, `R0426`, `R0427`, queue flush `R0611`. `R0426` calls virtual slot +0x18 of every element of the list at `[sim]+0x2c`, then the effect walk `R0641` on the list at `[sim]+4` (`L03065` to `L03066`).
- Slot +0x48 of the class table at `L03067` is `R0333` (`evidence/vtable-L03067.txt`). `R0333` calls `R0390` first (`L03068`; it forwards to the children through `R0642`) and, when that returns zero, switches on the message; the 0x401 case reaches `R0334` at `L02477` (`evidence/disasm-window-handler-R0333.txt`). `R0334` sweeps the map at view `+0x9b8` with virtual slot +0x3c (`L02473`), the same map with slot +0x4c (`L03069`), then the map at view `+0x9d4` with slot +0x3c (`L02474`). Each sweep takes the next bucket entry before it calls.
- Slot +0x3c of the CUnit and CAirUnit tables (`L02468`, `L02585`) is `R0548` (`L03070`, `L03071`); its action-7 arm calls slot +0x58 (`L02716`) when the phase equals ShootDelay. Slot +0x3c of the CProjectile table (`L02587`) is `R0558` (`L03072`).
- The client dispatcher stores records into `+0x9d4` at `L03073`, `L03074` and `L03075`; the unit-shot constructor `R0603` stores into it at `L03076` to `L03077`, chaining the new entry at the head of bucket `(id >> 4) mod count`.

**Confidence.** High: the call order inside `R0260` and `R0454` and the sweep order are instructions read from the EN image (the RU image is byte-identical, `evidence/image-hashes.txt`). Medium for the delivery path from the flush `R0611` to the queue the dispatcher drains: the sender chain was read through `R0217` and `R0643`; the queue's first dequeue in `R0509` was not traced.

**Unknown.** Which timer entry the saved battles ran through: `R0455` and `R0640` lack part of the order, and `R0193` has a second caller (`L01864`), `R0509` two others (`R0644`, `R0645`). Whether `L03078` skips the 0x401 case when `R0390` returns nonzero, and on what input. What the map at `+0x9b8` holds besides actors (`AI-SELECT-065` calls it the unit-id map). The conditions on which `R0509` skips an arm.

### ANIM-114

- The inner effect's first tick `R0629` reaches `R0630` (`L03079`) when neither stage test `R0631` nor `R0632` holds. `R0630` calls `R0635(effect, 1)` (`L03054`), which calls `R0646` (`L03080`); the client builds the record in arm 0x86 (`ANIM-112`).
- The send happens inside the effect walk, which precedes the flush `R0611` in the same sub-tick (`ANIM-113`). If the dispatcher of the same timer step drains that flush, it builds the record and the record sweep (`+0x9d4`) of that step runs after it, so the record's first driver call would be on the tick of the send. That drain was not traced.
- The effect list is a linked list with head at `+4` and tail at `+8`; the append `R0647` passes the tail field as the previous node (`L03081`). The iterator start `R0648` and the step `R0649` both return the current node and move the cursor to its successor before the caller visits it (`R0650`; the walk calls slot +0x18 at `L03082`). A node appended while the last node is visited is therefore not reached in that walk, and a node appended while an earlier node is visited is (`evidence/disasm-effect-walk-R0641.txt`, `evidence/calls-q2.txt`).
- The transport's fire arm `R0634` appends the inner effect to the same list through `R1113` (`L03083`, `L03084`).

**Confidence.** High for the builder, its caller and the pass order (instructions). Medium for same-step delivery and the first driver call on the send tick (the dequeue was not traced). High for the iterator taking the successor before the visit (`R0648`, `R0649`, `R0650` read); Medium for the append rule, because the tail link of `R0651` was not read.

**Unknown.** Whether another effect can follow the transport in the list when it fires (it would make the inner effect run in the same walk).

**Amended.** `EXP-0453`: the flush-to-dequeue path delivers in the sending timer step on the local path (`ANIM-121`).

### ANIM-115

- Rider `R0246`: `R0002` (`L03045`) builds the effect, then the damage message `R0561` (opcode 0x73, `L02655`) is sent when the target's hit points are above -10 (`L03085` to `L02654`).
- `R0003` builds the area effect by `R0652` and then the SpellTransport by `R0633` with the effect as inner object (`L03052`), the caster position and the spell's speed (`ANIM-111`). It sets the transport's countdown to 10 for spell ids 13 and 14 (`L03086`) and appends the transport alone (`L03053`, `R1113`); the inner effect is appended when the transport fires (`ANIM-114`).
- The messages are, in order: 0x73 at the rider tick, then 0x86 from the inner effect's first tick, one transport countdown later (`SAV-1154`). The rider calls `R0002` first, so the transport already exists when 0x73 is sent.
- Catapult and Ballista Units rows: charge 2, relax 38 and 50, token size 2 (`evidence/shot-class-columns.txt`, EN root).

**Confidence.** High: the rider, the constructor arguments and the append are instructions. Medium that the Catapult and Ballista reach this path through Fire_Ball: `EXP-0428` `shot-classes.csv` gives the Catapult's weapon spell; the Ballista's was not read here.

**Unknown.** The Ballista row's weapon spell. What the client does with the 0x73 message while the burst is pending.

**Amended.** `EXP-0453`: the Ballista weapon spell is Fire_Ball (`MAGIC-247`); client arm 0x73 applies at dispatch and holds no burst state (`ANIM-124`).

### ANIM-121

- Sim sub-tick `R0193` ends with the flush `R0611` on manager `L00522`. For a local session the setup `L03087` to `L03088` (when `+0x6bc` is 1 or 2) calls `L03089`, which sets each manager's peer pointer (`+8`) to the other (`L00522` and `L00625`); `L03090` moves the pending endpoints into the active list (`evidence/disasm-session-link-L13133.txt`, `evidence/disasm-peer-link-L03089.txt`, `evidence/disasm-pending-to-active-L03090.txt`).
- The pump `R0643` calls `L02449` per endpoint. For an endpoint whose local flag (`+0x104c`, set by `L03091`) is set, `L02449` finds the peer's local endpoint (`L03092`) and calls `L03093`, which appends the packet to that endpoint's receive list (`+0x1014`) under a critical section, before returning (`evidence/disasm-endpoint-send-L02449.txt`, `evidence/disasm-local-append-L03093.txt`, `evidence/disasm-local-peer-endpoint-L13134.txt`).
- The client dispatcher `R0509` loops `R0513` (dequeue through `L02425`, `L02429`) until the queue is empty, in the same timer step that ran the simulation step (`ANIM-113`; `evidence/disasm-client-dispatcher-R0509.txt`, `evidence/disasm-dequeue-R0513.txt`). A packet flushed in the sub-tick is therefore dequeued in that timer step.
- A TCP mode exists in the same code for multiplayer; it is not the local path and was not read for latency.

**Confidence.** High for the append and the dequeue loop (instructions). Medium that the saved battles ran the local path: the runtime value of `+0x6bc` is not in a save.

**Unknown.** Latency on the multiplayer socket path.

### ANIM-122

- The projectile-picture sources are the cast spawner (even pictures), client arm 0x86 (13, and the odd 17, 27 and 51), arm 0x8b, arm 0x8c (picture 36) and the loader (`ANIM-109`, `ANIM-110`).
- `R0562` stores only opcode 0x73 (when its argument is 0x73), 0x7a and 0x82; it has 13 direct call sites and no data-pointer reference. `R0627` stores 0x88 (argument 0) or its argument byte; its 3 call sites (`L03094`, `L03095`, `L03096`) push 0, 0x89 and 0 as the first-pushed argument and it has no data-pointer reference (`evidence/disasm-opcode-writer-R0562.txt`, `evidence/disasm-opcode-writer-R0627.txt`, `evidence/disasm-callsite-R0627-from-L03094.txt`, `evidence/disasm-callsite-R0627-from-L03096.txt`, `evidence/xrefs-q7.txt`). The 13 call sites of `R0562` are listed in `evidence/xrefs-q7.txt`; their arguments were not read, and the stored opcodes above do not depend on them. This closes the two register-form writers `ANIM-110` left open: neither writes an opcode that arms 0x86, 0x8b or 0x8c handle.
- Arm 0x8a (`L03097`) builds no record. It looks up an existing drawable by the id word in the `+0x9b8` map and sets its `+0xa0` from its class, `+0x84` = 8, `+0x94` = 0, a word list, `+0x86` and `+0xa4` = 0x24 (36): it starts the picture-36 channel on a drawable (`evidence/disasm-client-arm-0x8a-L03097.txt`, `evidence/dispatch-arms.txt`).
- Saves: the rescan of 120 `.sav` files finds 79 with a Projectiles store, 5 non-empty, pictures 1, 2, 5, 10 and 13, none of 14 or above (`evidence/corpus-projectiles.txt`). These are the counts of `SAV-1152` without the one engine-written save, which is excluded as not ROM1 evidence; the engine-written file is not counted.

**Confidence.** Medium: the census is of direct calls and data pointers to the two writers and of the arms the dispatcher tables name. Calls through a register formed elsewhere are not enumerated.

**Unknown.** A projectile record with picture 14 or above built from a message the tables do not route; no save holds one.

### ANIM-123

- Sim side: `R0630` calls `R0635(effect, 1)` (`L03054`, `evidence/disasm-blast-arm-R0630.txt`). For spell id 2 it fills buffer `L03042`: opcode byte `+9` = 0x86, effect id word `+0xa`, picture byte `+0xc` copied from effect `+0xe` (13 for Fire_Ball, `EXP-0432` `evidence/spell-pictures.txt`), cell x and y bytes at `+0xd` and `+0xe`, word `+0xf` = 0x16 (22) (`evidence/disasm-effect-sender-R0635.txt`).
- Client arm 0x86 (`L03098`) builds the record by constructor `R0609`: picture at `+0x20`, x and y = cell * 256 + 128, action phase `+0x94` = -1, action segments `+0xa0` = message word `+0xf`, `+0x84` = 1, `+0x14` = view `+0x9b4`; the id comes from the counter at view `+0xa0c` and the record is stored in the map `+0x9d4` (`evidence/disasm-client-arm-0x86-L03098.txt`).
- `R0614` derives from the record through CProjectile slots `+0x20` and `+0x24` (`L03099`, `L03100`), both returning the constant 1 (`evidence/disasm-record-derived-R0614.txt`, `evidence/disasm-projectile-size-slots-L03099.txt`). Driver arm 0x0d (`L03001`) registers cells -1 to +1 on both axes, a constant 3 x 3 (`evidence/disasm-driver-arm-0x0d-L03001.txt`).
- The record has no size field; the size byte exists only on the simulation side (`MAGIC-246`).

**Confidence.** High for the message layout and the record constructor (instructions); Medium that the word `+0xf` feeds only the segment count: the population searched is client arm 0x86 and `R0614`, and the other readers of the record's `+0xa0` were not enumerated.

**Unknown.** The readers of record `+0xa0` outside the two routines read.

### ANIM-124

- Dispatcher arm 0x73 (`L02661`; `evidence/dispatch-arms.txt`) reads the id word at message `+0xa` and the new hit points at `+0xc`, looks the id up in the `+0x9b8` map and, when absent, logs and ends (`evidence/disasm-client-arm-0x73-L02661.txt`).
- When the stored hit points (`+0xfc`) exceed the message value and view flag `+0xaa4` is set, it builds a `R0566` object and adds it by `R0568` to the list at view `+0x3f3c`; it then calls actor slot `+0x68` with 0 for equal values, else 2 or 1 depending on the new value against half the maximum (`+0x100`), stores the hit points, sends 0x408 to the two windows at `+0xe0` and `+0xe4` when the actor is the selection (view `+0x138`), and sets `+0x10c` = 1.
- The arm touches the `+0x9b8` map and the `+0x3f3c` list only; it reads neither the `+0x9d4` record map nor any burst or projectile field.

**Confidence.** High for the arm's reads and writes (instructions); Medium for the role of the `R0566` object, which is inferred as a damage figure from its arguments.

**Unknown.** The role of the `+0x3f3c` list's object; how the hit reaction of slot `+0x68` looks.

### ANIM-125

- Builder. `R0561(target, dmg)` pushes tag `0x73` and sends through the layout routine `R0562`. That routine builds `0x73` only when the target's `vt+0x2c` returns non-zero (`L03101`..`L03102`, then the tag test at `L03103`); the message carries the target id at `+0xa` and the target's current health `[target+0x94]` at `+0xc`, and `dmg` is not copied. A target whose `vt+0x2c` is zero goes to the class-cast branch at `L03104` (tag `0x7a`) or, for a structure, to the `vt+0x34` branch at `L03105`, which builds tag `0x82` with the `+0x42` health at message `+0x12`; the tag argument is not tested there. Delivery goes to human or owner players, filtered by `L09285` and `R0676`. The builder has five direct call sites (`evidence/callers.txt`) and the image holds no other four-byte occurrence of its entry address.
- Unit strike `R0246` (`L02654`). It subtracts the damage from the target's health, then sends when the target was alive before the blow (`L03106` tests the saved flag at frame offset −8), whatever the damage; for a target that was already dead it sends when the health after the blow is above -10 (`L03107`). It returns without sending when the reach test fails or the attacker or target is null.
- Structure strike `R0563` (`L03108`). The same alive-before gate on `+0x42`, with damage from `R0653`; its target is a structure, so the layout routine builds tag `0x82` and not `0x73`. The structure arm of area direct damage (`L03109`) does the same.
- Area direct damage `R0564` (`L03110`, `L03109`) sends only after `jle` skips a non-positive damage (`L03111`, `L03112`). The Token-8 applier `R0565` (`L02979`) sends only for a non-zero result (`ANIM-074`).
- Zero damage in the unit strike. The resolver `R0265` returns `max(total, 0)` (`HERO-HEALTH-032`); the physical component is clamped at 0 after absorption (`HERO-CLAMP-030`). A unit-strike blow with total 0 leaves health unchanged and still reaches `0x73`.
- Client. Arm `0x73` compares the drawable's stored `+0xfc` with `msg+0xc` as signed 16-bit values and calls the hook with 0 when they are equal (`L02698`), before any band test (`ANIM-095`); the call has no view-flag gate.

**Confidence.** High for the five sites, their gates and the arm's order: instructions in routines read whole, and no pointer form of the builder. Medium that a shipped blow reaches zero damage in play: the strike's zero return is read, the population of shipped attackers and targets that produce it was not measured.

**Unknown.** The client arm that receives tag `0x82`, and whether it calls a hurt hook; whether the stored `+0xfc` can differ from the server's health at the time of a zero-damage message (a regen sync in flight would turn event 0 into a band event).

### ANIM-126

- Stage switch. The state-sync arm saves the old stage `+0x15a` at `L02542`, applies a new stage only under mask bit 3 (`L02541`), then switches on the stage (`L02764`). Stage 1 with old 0 stores action 6, calls `vt+0x68(3)` at `L02765` and then `R0551`; stage 1 with old 1 calls `vt+0x68(2)` at `L02766`. The call follows the action stores in program order: it is made when the message is handled, not after the action's ticks. The switch runs on any state sync that reaches it, whatever the mask.
- Projector `R0059(object, recipient, mask)` has 51 direct call sites and no pointer form. Recipient 0 means every human player and the owner; non-zero means that player only. Mask bit `0x1` health, `0x2` `+0x9a`, `0x4` a dword at `+0x130`, `0x8` the stage byte, `0x20` position, `0x4000` class and `actor+0x4b`.
- Death arm `L12989`, from the actor tick `R0037`: on the tick where the dying branch finds stage 0 it sets the stage to 1 and sends mask `0x409`, recipient 0, once. The client sees 0 to 1 and plays event 3 (`die.wav`, no 1500 ms gate).
- Bleed `R0654` (`L03113`): for health below 0, when the low two bits of the full-tick counter `[[L00285]]` are 0 (every fourth full tick), it subtracts 1 from health and sends mask `0x1` to the owner `[actor+0x14]` only. The stage byte is unchanged, so the client sees 1 to 1 and plays event 2 (`hard.wav`), subject to the 1500 ms gate. By `HERO-DECAY-069` health goes from -h to -10 in `(10 - h) x 4` full ticks, so a corpse that starts the live-list decay at health -1 to -9 produces at most 9 such syncs, 4 full ticks apart. The same routine's positive-health regen sends mask 1 or 2 to recipient 0.
- Session-entry projection `L03114` (callers `L03115`, `L03116`): for the joining player it sends mask -1 for the player's actor and every actor on the world list, then for every dead-list actor with stage below 5. A new drawable starts at stage 0, so a stage-1 corpse yields 0 to 1 and event 3 once per corpse per join or load. A second projection `R0655` sends the same mask for the owner's actors.
- Other senders of the projector (spell casts, effect ticks, hire, command handlers, chargen, experience, decay, teardown) were classified by mask in the listings, not by trigger. A projector call whose mask omits bit 3 preserves client stage only if it emits an ordinary state sync that reaches the switch. ANIM-129 establishes the effective ordinary-mask send gate; mask 0 and mask `0x80` alone do not emit that sync.

**Confidence.** High for the death arm, the bleed and the entry projection: named instructions and gates. Medium for the completeness of the sender list: ANIM-129 reproduces the finite 51 direct sites, and ANIM-129..133 classify their local predicates; computed callers and native stage-1 reachability outside the named controls remain unresolved.

**Unknown.** Native stage-1 reachability and scheduling outside the named paths.
ANIM-129..133 provide the finite per-site trigger table and bounded remaining
aliases/callbacks. Dead-list decay sends stages 2..5 under mask 9; it is a
transition away from stage 1, not a repeated-stage source.

**Amended.** `ANIM-095`: the voice call is immediate; the "then" in its table is an order of stores, not a delay. ANIM-129 narrows the omitted-stage-mask clause to admitted ordinary state syncs; ANIM-130..133 add bounded remaining caller predicates and client helper bindings. The named death, bleed and entry routes and Medium completeness stand.

### ANIM-127

- Setter. `R0590` (sole caller `L03117`, in the state-sync arm) takes the stack value P, the second byte of the mask-`0x4000` field, which is server `actor+0x4b` (`HERO-APPEAR-040`). For a class `T` below `0x1a` it stores `+0x24 = P & 0x7f`, sets `+0x18c = (old & 0x80) | 0x18`, ORs `4` when `P & 0x80` (the sex bit) and `2` for `T` `0x17` and `0x18`. For `T` in `[0x20, 0x40)` the sex bit comes from `(T - 0x21) & 1` instead (`ANIM-096`).
- Humans row constructor `R0656`: P starts as the row's slot 17 (`R0657`; -1 keeps the base constructor value 1). A key with a dot gives a face number (the digits two characters after the dot, applied when above 0) and a female flag (the character after the dot equals `f`); a key without a dot gives female 1. A non-negative row slot 18 replaces the flag (`L03118`..`L03119`). With hero mode 0, `P |= female << 7` (`L03120`..`L03121`); with hero mode non-zero the class becomes `0x21` or `0x23` plus the flag instead.
- Map placement `R0151`. Ordinary Human type-key arm: always `P = secondary key low byte`, then bit 7 set or cleared from placement flag bit 2 (`ALM-FLAGPATH-109`). Definition-id arm: the same stores only when the secondary key word is non-zero; otherwise P stays the constructor value. NPC arm: no store; the constructor gets hero mode from the `Hero` flag.
- Tavern hire. The call sites `L03122` and `L03123` build the key with the format `NPC%02d_%d` and pass hero mode 0; the site `L03124` passes a Humans row by index. None overrides P, so the value is the constructor's.
- Summons. The cast spawner `R0003` builds an actor with the Units constructor `R0501` (`L03125`). The Humans constructor has 19 direct call sites, none in the cast spawner, and no pointer form. No row of the Units collection has a class below `0x1a` (118 rows, both roots), so a summoned unit does not take the below-`0x1a` branch of the setter.
- Data, both roots (`evidence/people-*.txt`). 215 Humans rows each; slot 18 is -1, 0 or 1 in 31, 110 and 74 rows on both roots; no Humans key contains a dot, so the face and the dot flag are not reached. Slot 17 differs in four rows (the `F_Kadagan` rows, 1 on EN and 29 on RU), slot 16 in four (the `NPC10_` rows, 3 on EN and 4 on RU). The 52 `NPC%02d_%d` rows give 28 male and 24 female values on both roots. Human-band placements, over ten hard-coded root map names where the file is present plus the `scenario.res` M7R payloads (464 per root): 1422 on EN and 625 on RU. RU lacks four of the ten root files (`Tomb`, `Beast`, `Cross`, `Kids2`), so the two counts are install populations and are not comparable data. Of EN's, 972 go through the constructor only (396 with the bit set), 433 through the spawner's stores (152 with the bit set), 15 are NPC and 2 type-key. Of RU's, 592 go through the constructor only (241 with the bit set), 17 through the spawner's stores (none with the bit set), 15 NPC and 1 type-key. No placement resolves to a row whose slot 18 is -1, so the default female flag is not reached by shipped placements.

**Confidence.** High for the setter, the constructor and spawner stores and the summon route (instructions; call census with no pointer form). Medium that the loader's `actor+0x4b` store, which goes through a pointer and is not visible by displacement, adds no further source.

**Unknown.** The loader's store. Hero-mode persons (`0x21`/`0x23`) take their sex bit from the class, so this claim does not cover them.

**Amended.** `ANIM-096`: the source of the setter's parameter bit `0x80`.

### ANIM-128

- Object. The sound configuration object at `L03126` is built by `R0658` (via `L03127`): music, effects and speech attenuations at `+0x8`, `+0x10`, `+0x18` are each `0xfffffd44` (-700), with range words `0x1388` at `+0xc`, `+0x14`, `+0x1c`. `[L02733]` is `+0x10` (effects) and `[L02739]` is `+0x18` (speech). The object's only code references are the constructor, the options load and save (`L03128`, `L03129`) and the dialog creation (`L03130`).
- Writers of `[L02739]`. The dialog handler `L03131` stores it at `L03132` for control id 8, from `R0659(position, range)`, `trunc(-(a - b)^2 / b)` (range -max to 0). The registry load `R0660` (single caller `L03133`) reads `SoundSpeechPos` into `+0x18` through a pointer with `RegQueryValueExA`, without clamping; the save `L03134` writes it back. The effects slider (id 7) stores `[L02733]` at `L03135`.
- Readers of `[L02739]`: the slider's test sound (`L03136`), the hurt hook for event 1 to 3 (`L03137`), five other bank voice routines (`L03138`, `L03139`, `L03140`, `L03141`, `L03142`), `L03143`, `L03144`, `L03145`, `L03146` and the town pager voice (`L03147`). `[L02733]` has more than 100 reader sites, among them the swing and the hurt event 0.
- Combination. The hurt hook passes `out1 + setting` to `R0386`, where `out1` is the distance term of `R0574`, and the setting is `[L02739]` for event 1 to 3 and `[L02733]` for event 0 (`ANIM-094`). No instruction combines the two settings. `R0386` clamps only the low side at -10000 before `SetVolume`.

**Confidence.** High for the writers visible by absolute address, the reader list, the default and the absence of a combining instruction. Medium that no pointer-based store changes the object elsewhere: the object's address has four code references, and the load writes through a pointer.

**Unknown.** What a positive value from the registry does at `SetVolume` (inferred to fail); the roles of the `L03143`, `L03144`, `L03145` and `L03146` readers.

## Human class record and swing sound

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-105 | The swing hook `R0573` plays `Sound[0]` of the class record `units.reg[drawable+0x20]` and reads no equipment; it plays nothing for classes 1, 2 and 23 (3 of 34, EN = RU) and when the sound buffer slot is empty. | High | ● active | [EXP-0431](../experiments/EXP-0431-human-class-sound/) |
| ANIM-106 | A map-authored human holds its Humans-row `typeID` as class record from tick zero: `R0657` streams row slot 16 into `actor+0xe`, the client copies it to `drawable+0x20`, equipment does not choose it; npc `Hero` placements excepted. | High / Medium | ● active | [EXP-0431](../experiments/EXP-0431-human-class-sound/) |
| ANIM-107 | A hero's drawn class is re-derived from visible equipment on each equipment message: Weapons row D picks a `heropicture` name, a shield adds `_`, armour is not read; a dagger gives swordsman (class 3, `Sound[0]` 100). | High / Medium | ● active | [EXP-0431](../experiments/EXP-0431-human-class-sound/) |
| ANIM-108 | `heropicture.txt` is 26 entries, not 25: `R0661` makes every line an entry, a blank one included, so Weapons rows 24 and 25 meet their own names and row 27 indexes past the file into the next loaded file. | High | ● active | [EXP-0431](../experiments/EXP-0431-human-class-sound/) |
| ANIM-116 | A placed human built by the definition-id arm holds its Humans row's typeID as class; `R0151` reads the record's class key only for the band test and the type-key and Units arms. 7 of 7 differing saved actors (2 maps) agree. | High / Medium | ● active | [EXP-0442](../experiments/EXP-0442-class-swing/) |
| ANIM-117 | No store reachable from the client unit dispatcher `R0509` sets `+0x18c` bit 0 for a class below 0x1a, so `R0551` returns and equipment does not re-select a placed human's class. | High / Medium | ● active (amended, partially retracted) | [EXP-0442](../experiments/EXP-0442-class-swing/) |
| ANIM-118 | The dragontooth mage (`141.alm` unit 5, Humans row 208) holds class 24 and its swing hook is passed slot 510, `magic/arrow.wav`; its record class key 23 would be silent. 43 placements in 10 maps share that pair. | High / Medium | ● active | [EXP-0442](../experiments/EXP-0442-class-swing/) |

### ANIM-105

- `R0573` (`evidence/disasm-swing-hook-R0573.txt`) loads the 32-bit field `+0x20` of the drawable (`R0573`), indexes the class array `[L02113]` with it (`L03148`, `L03149`), reads the `Sound` data pointer `class+0xc0` (`L03150`) and returns when `Sound[0]` is 0 (`L03151` compares the value with zero, jumping to `L03152` on equal). Otherwise it takes `[L01636][Sound[0] * 4]` as the buffer (`L03153`, `L03154`), returns when that is null (`L03155`) and passes it to the channel manager `R0386` (`L03156`).
- The routine reads drawable fields `+0x08`, `+0x0c`, `+0x20` and `+0xe0` only (pan and distance come from the first, second and last). It reads no equipment slot (`+0x15c...`) and no flag word (`+0x18c`), and it keeps no copy of the class or the sound: each call resolves the class from `+0x20` again.
- `units.reg` `Sound[0]` of the 34 shipped classes (`evidence/class-sound-en.txt`; the RU file is byte-identical, `evidence/roots-comparison.txt`): 0 for class 1 (Unarmed Fighter), 2 (Unarmed Fighter with Shield) and 23 (Unarmed Mage). The others hold 100 (swordsman family, ids 3 to 5, 19, 80), 110 (axe family, ogre), 120 (club family, catapults), 130 (archer, orc archer), 140 (crossbow), 150 (pike family, goblin, lance hero), 510 (Human Mage, id 24) and creature values from 160 to 512.
- Two vtable copies carry the hook (`.rdata` `L03157` and `L03158`, `evidence/dword-scan-class-routines.txt`), and no direct call reaches it (`evidence/callers-class-routines.txt`). `ANIM-SND-022` names the driver arms that fire it.
- A human attack therefore plays no swing in these cases: the drawn class is 1, 2 or 23; the buffer slot of `Sound[0]` is null; or no DirectSound object exists (`ANIM-SND-022`). A drawn class among the 31 sounding ids (3, 4, 5, 7 to 15, 19, 21, 24, 26, 27, 64 to 66, 68 to 76, 79, 80) passes the first test; ids 6, 16 to 18, 20 and 22 are not shipped classes.

**Confidence.** High: a routine of fewer than 50 instructions read whole, both vtable slots and the absence of direct callers enumerated, the `Sound[0]` table read from both roots by the loader's own arithmetic.

**Unknown.** Whether each shipped `Sound[0]` id has a loaded buffer (the population of `[L01636]` was not read). "Plays" here means the pair is handed to `R0386`.

### ANIM-106

- **Server.** The Humans constructor `R0656` (`evidence/disasm-humans-ctor-R0656.txt`) zeroes `actor+0xe` (`L03159`) and calls `R0657` (`L03160`) with the row cursor at `row+8`. That routine reads 17 values in a fixed order through three row readers (`R0662` u16, `R0281` u8, `R0663` u16). The 17th, after 10 single reads and a 6-element loop, is the u16 read through `R0663` at `L03161`, with the row offset advanced by 0xe at `L03162` (`evidence/disasm-humans-row-stream-R0657.txt`). The reader stores only when the cell is not -1. Row slot 16 is titled `typeID` in the shipped Data.bin (`evidence/humans-rows-en.txt` header).
- **Overwrite.** `L03163` compares the constructor-mode argument (frame offset +0xc) with zero and jumps to `L03120` on equal, which skips the player-character stores `L03164` and `L03165` (`gender + 0x23` or `+ 0x21`) when the constructor mode is 0. ALM placements by definition id or explicit typeID pass 0; the npc arm passes the npc's `Hero` flag (`PARTY-M20-031`), so only npc placements carrying that flag take the hero path of `ANIM-107`.
- **Client.** The creation block of `R0509` stores the wire class byte into `drawable+0x20` (`L02143`, `evidence/disasm-client-creation-class-store.txt`; the byte is `word[actor+0xe]`, `HERO-APPEAR-040`). `R0590` (`L03117`) rewrites `+0x20` only for ids in `[0x20, 0x40)` (`L03166`, `L03167`, stores `L03168` and `L03169`); every other id jumps to `L03170`, whose branches write `+0x24` and flag bits and no `+0x20` (`evidence/disasm-drawable-class-split-R0590.txt`). `R0551` tests bit 0 of the flag word `+0x18c` (`L03171`) and leaves to `L03172` when it is clear (`evidence/disasm-hero-derive-gate-R0551.txt`). The hero block sets that bit (`L03173` sets bits 0 and 3); the character-screen routines `R0592` and `R0591` set it on their own drawables.
- **Population.** The EN Humans table has 215 entries: 5 without a `typeID` (slot -1), 3 whose `typeID` (17, 18, 20) names no `units.reg` class, 207 with a class (`evidence/humans-rows-en.txt`). The name chain of `HERO-APPEAR-042`, applied to each row's weapon cell and shield cell, would assign a different class to 60 of the 207 (52 in RU, where 8 sword-and-shield rows carry `typeID` 4 instead of 3), and a different `Sound[0]` to 33 of them. 21 rows carry a silent class (typeID 1 or 23) and none of the 21 names a weapon.
- **Saves.** The original saves, all EN (15 dated folders, 63 `game*.sav`, `evidence/save-corpus-digest.txt`) hold 161 Human records owned by a human Player in 50 saves. 93 carry `word[actor+0xe]` in `0x21..0x24`. The other 68 (rows 50, 54, 58, 200 and 201) carry exactly their EN row's `typeID`: 68 of 68 (`evidence/save-typeword-join.txt`), but on a small population of distinct actors: row 58 is 3 guards of one map (43 records in 15 saves), row 54 is 3 actors in two saves of one dated folder (6 records), and rows 50, 200 and 201 agree under both rules. The join is against EN rows; row 54 separates the rules only in EN (`typeID` 3, chain 4; RU `typeID` 4, chain 4). `savunit -mode party` prints no visible weapon or shield, so the chain value is computed from the authored Humans cells and the constructor equip rule, not from saved equipment. That includes 6 records of row 54 (sword and shield: the name chain gives class 4, the saved word is 3) and 43 of row 58 (mace and shield: chain 11, saved 10).
- **Mission 20.** Sarindar is Humans row 201, `typeID` 23, Unarmed Mage, `Sound[0]` 0, no weapon cell (row 201 has 7 saved records, word `0x17`). His swing plays nothing. The three guards are row 58, class 10 (club family), `Sound[0]` 120 (43 saved records, word `0x0a`).

**Confidence.** High for the server stream, the mode test, the client creation store and the two class gates; the High rests on the constructor and gate listings, and the saved words (68 of 68) corroborate on about 6 actors over 2 maps, on which the equipment-derived alternative gives a different class for rows 54 and 58. Medium for "nothing rewrites the class during play": the instrument is a capstone linear sweep of every code section for stores to `+0x0e` of width 2 (54 hits, `evidence/sweep-word-stores-disp-0e.txt`) and `+0x20` of width 4 (629 hits, `evidence/sweep-dword-stores-disp-20.txt`). The actor-family hits were read (`L03174`, `L03175`, `L03159`, `L03164`, `L03165`, `L03176`); the drawable-class hits read are the virtual init `R0664`, `L02143`, `R0590` and the 17 arms of `R0551`. The rest of the 629 were not each classified. Writes through block string moves or a copied object are outside the sweep.

**Unknown.** The three rows with `typeID` 17, 18 and 20 (ManHorse) name no class; whether any shipped map places them is open. The creation path through `R0664`'s caller was not listed. The client store at `L02143` runs only when creation mask bit `0x4000` is set (`L03177`, the bit of `HERO-APPEAR-040`). The 5 Humans rows with no `typeID` leave `actor+0xe` at the 0 written at `L03159`, a value that names no class.

### ANIM-107

- **Inputs.** Visible-equipment slot 0 is `actor+0x74` and slot 1 is `actor+0x78` (`HERO-APPEAR-047`). Slot 0's low five bits D index the `heropicture` entries (`HERO-APPEAR-042`, `HERO-APPEAR-052`); a non-null slot 1 appends `_`; a mage whose name stays `unarmed` becomes `mage`. D is the Weapons runtime row: `R0665` stores the `R0666` lookup at `item+0x0c` (`L03178`, `L03179`), and `R0667` builds the word as `(material << 12) | (kind << 8) | (shape << 5) | row` (`evidence/disasm-weapon-item-word.txt`, `L03180`, `L03181`).
- **Table** (`evidence/hero-class-table-en.txt`, RU identical). Fighter without shield: rows 2 to 5 (Dagger, Short, Long, Bastard Sword) swordsman, class 3, `Sound[0]` 100; row 6 swordsman2h, 5, 100; rows 7 to 10 clubman, 10, 120; rows 11 and 18 axeman, 7, 110; rows 12 and 19 axeman2h, 9, 110; rows 13 and 14 (staves) mage_st, 24, 510; rows 15 to 17 pikeman, 12, 150; rows 20 and 21 archer, 14, 130; row 22 xbowman, 15, 140. With a shield: rows 2 to 5 give class 4, rows 7 to 10 class 11, rows 11 and 18 class 8, rows 15 to 17 class 13, each with the same `Sound[0]`.
- **Dagger.** Row 2 gives D 2, entry 1 `swordsman`, class 3, `Sound[0]` 100 (class 4 with a shield, also 100). The first swing after the message that carries the dagger reads class 3.
- **Silent heroes.** Empty slot 0 or BareHands (row 1) gives class 1 (class 2 with a shield) and a mage gives 23: all `Sound[0]` 0. Rows 23, 24 and 25 name entries that match no arm (blank, `Sonic Beam`, `Flame Thrower`); a name no arm matches leaves class 1, stored at `L03169`. Row 26 (Boulder Thrower) reads the last entry, `swordsman`, and gives class 3 (`Sound[0]` 100), an audible row. A two-handed weapon, staff or bow with a shield gives a `_` name no arm matches and also leaves class 1. Row 27 reads the entry past the file (`ANIM-108`) and gives class 1.
- **What changes it.** For a living hero only a change to slot 0 or the shield presence of slot 1 changes the class; a mage whose action code `+0x84` is 6 (dying) takes the `mage_st` arm of `HERO-APPEAR-042`. The armour material D-field of slot 7 selects a sheet directory (`HERO-APPEAR-043`) and never the class. The class is rewritten on each processed message that runs `R0551` (`HERO-APPEAR-044`); the hook of `ANIM-105` re-reads `+0x20` at each call, so the swing after such a message uses the new class and the swing before it the old one.

**Confidence.** High for the derivation inputs and the table: every row is the recomputation of `HERO-APPEAR-042`'s chain over the shipped `heropicture.txt`, `Weapons` and `units.reg` on both roots, the row-to-D identity read from the constructor and the builder. Medium for timing: outside the actor serializer only two unauthored trigger arms send equipment (`HERO-FIGURE-064`), and the serializer sends it on mask bit `0x80` (`L03182`, `evidence/disasm-equipment-sender.txt`); which serializer call carries that bit after an equip, drop or swap was not enumerated.

**Unknown.** Whether a hero can hold a two-hander with a shield (`Weapon::Equip` has a two-hand arm, not traced here).

### ANIM-108

- `R0661` (`evidence/disasm-heropicture-loader-lines.txt`) runs one loop per line: store the line start in the shared pointer table (`L03183`), scan to the next CR (`L03184`...`L03185`), zero it (`L03186`), step two bytes and stop at the payload end (`L03187`). A blank line is an entry of its own.
- The shipped `heropicture.txt` therefore has 26 entries in both roots: entries 0 to 21 are the names listed in `HERO-APPEAR-052`, entry 22 is blank, entries 23, 24 and 25 are `Sonic Beam`, `Flame Thrower` and `swordsman`. `HERO-APPEAR-052` numbered the last three 22, 23 and 24.
- Check by names neither side was fitted to: Weapons row 24 is `Sonic Beam` and row 25 `Flame Thrower`, and index D - 1 reaches those very names only with the blank counted. Row 23 (`rem`) reaches the blank entry; row 26 reaches `swordsman`.
- `R0668` adds the file's base `[this+0xc]` to the index with no bound (`HERO-FIGURE-063`), and each file's base is the running table count (stored at `+0xc` of the file object by `L03188`). The files load in the order `L03189` (`heropicture.txt`) then `L03190` (`stats.txt`, string at `L03191`), so D = 27 (Weapons row 27) reads the first entry of `stats.txt`. That text matches no chain arm, so the class stays 1 (`evidence/hero-class-table-en.txt`, row 27). `HERO-FIGURE-063` put the first past-the-end index at D = 26; it is D = 27.

**Confidence.** High for the line loop, the 26 entries and the name alignment (both roots). The `TEXT` ledger independently records the shared-index bases `heropicture` 274 and `stats` 300 (`TEXT-STRTAB-023`), a difference of 26, which matches the 26 entries and puts `stats.txt` immediately after. High for the adjacency of `stats.txt` on those bases and the two call sites; the first-entry text is the shipped `stats.txt` line, and the table is not readable statically (`HERO-APPEAR-052`).

Corrects the entry numbering and the length in `HERO-APPEAR-052` and the first out-of-range index in `HERO-FIGURE-063` (`retracted.md`).

### ANIM-116

- **Arm.** In `R0151` a record with class key below `0x1a`, flag bit 0 clear and a nonzero definition id other than `0xcdcdcdcd` takes the definition-id arm (`L03192`..`L03193`). It calls `R0498(L02110, record+0x10)` (`L03194`..`L03195`), which scans the Humans table from the last row down for the row whose server id (slot `0x18`) equals the definition id, and passes that row to the Human constructor `R0497` with mode 0 and `record+0x48 & 0x80` (`L03196`..`L02315`). `R0657` then streams Humans slot 16, `typeID`, into `actor+0xe` (`ANIM-106`) and the equipment cells are equipped after it.
- **Class key.** `record+0x08` is read at `L03197` for the lookup `R0495(L02110, ...)` and at `L03198` for the band test. The lookup result (frame local −0x38) is read at `L03199` (the type-key arm, where the key is the `typeID` by construction) and `L03200` (the Units arm) only. The definition-id arm reads `record+0x10`, `record+0x48` and, for the face byte, `record+0x0c` (`L03201`..`L03202`). The flag-bit consumers are `ALM-FLAGPATH-109`; the arm order is `MISSION-ARM-006`. The listing is `evidence/disasm-spawner-R0151.txt`.
- **Population.** `tools/campaign -mode swing` over every type-6 record of the 28 campaign maps with a Humans-band placement (`evidence/swing-en.tsv`, `swing-ru.tsv`): 464 placements, 448 definition-id, 15 npc, 1 type-key; all resolve a Humans row whose `typeID` names a `units.reg` class. Four are npc placements with the `Hero` token (constructor mode nonzero), which take the hero path of `ANIM-107`. Of the 449 definition-id and type-key placements, `record+0x08` differs from the row `typeID` in 102 (EN) and 97 (RU), and names a class with a different `Sound[0]` in 77 (both roots) (`evidence/census-en.txt`, `census-ru.txt`).
- **Saves.** Four original-written documents (two of mission 41, two of mission 141) hold 17 distinct map-placed Humans. The serialized type word equals the row `typeID` for 17 of 17 and equals the record key for 10 of 17; for the 7 actors whose key differs (mission 141 unit 5; mission 41 units 34, 35, 40, 64, 65, 66) it never equals the key (`evidence/save-typeword-mages.txt`). Mission 41 units 34, 35, 64, 65, 66 (saved 3, key 4, chain 4) separate the row `typeID` from both the key and the equipment chain; the two mage units (mission 141 unit 5, mission 41 unit 40) separate it from the key only, their chain also being 24. The four documents are 2 independent samples from 2 maps.

**Confidence.** High for the arm, its reads and the row resolution: one routine read from `L03203` to `L03204`, with the lookup result's only consumers enumerated by operand. Medium for the population and the saved words: a join of decoded tables, and 7 differing actors on 2 maps (2 independent samples).

**Unknown.** Whether any later store changes `actor+0xe` of a placed human (the instrument is the word-store sweep of `ANIM-106`, Medium).

### ANIM-117

- **Gate.** `R0551` reads the flag word `+0x18c` at `L03205` and leaves to `L03172` at `L03206` when bit 0 is clear (`evidence/disasm-hero-derive-gate-R0551.txt` of `EXP-0431`). The hero arm of `R0590` sets bits 0 and 3 (`L03173`). For a class below `0x1a` it stores `(old & 0x80) | 0x18`, plus 4 when the face byte's bit 7 is set and plus 2 for classes 23 and 24 (`L03207`..`L03208`), so bit 0 is clear at creation for every map-authored human of the shipped classes.
- **Message arms.** `R0509` holds all six callers of `R0551` on the client's unit path (`L03209`, `L03210`, `L03211`, `L03212`, `L03213`, `L03214`; `HERO-APPEAR-044`). Its stores to `+0x18c` are six, none setting bit 0: `L03215`..`L03216` sets bit 0x20, `L03217`, `L03218` and `L03219` set bit 0x08, `L03220` sets bit 0x80 and `L03221` masks the word with 0x7f (`TERR-194`). `R0592` (stores of `0x29` and `0x2b`) and `R0591` set bit 0 on the character screen's own drawables (`ANIM-106`). `R0339` and `R0240` also call `R0551` and store to `+0x18c`; their receivers were not traced.
- **Class stores.** The dispatcher has nine stores to the class field `+0x20` (`ITEM-136`). Four are the backpack arm. `L02143` is the unit creation block after allocations of `0x1b0` bytes (`L03222` then `R0593`); its class source is the global `[L03223]`, decoded under creation-mask bit `0x4000` (`L03177`). That it runs on creation only rests on that guard and the allocation, not on a traced dominance proof (Medium). `L03224` is the Building arm, allocations of `0x138` bytes then `L03225`. `L03226`, `L03015` and `L03017` are the projectile arms, allocations of `0x14c` bytes then `R0609` (`evidence/dispatcher-class-stores.txt`).
- **Swing.** The hook resolves `units.reg[drawable+0x20]` on every call and keeps no second class field (`ANIM-105`). A hero is the exception: `ANIM-107` derives its class on each equipment message.
- **Effect.** A placed human that picks up, drops or swaps a weapon keeps its class and so its `Sound[0]`. A silent class stays silent whatever it wields: 66 of the 460 mode-0 placements of the shipped maps hold class 1 or 23. The 448 definition-id placements are read here; the 11 npc and 1 type-key placements are the arms of `MISSION-ARM-006` and share the count by that claim, not by a read of their own.

**Confidence.** High for the gate, the creation values and the attribution of the dispatcher's nine class stores and six flag stores: named instructions and allocation sizes. Medium for the creation-only reading of `L02143` and for the `+0x20` population outside the dispatcher, which was not classified.

**Unknown.** The `+0x20` population outside the dispatcher.

**Amended.** The dispatcher's flag stores are six, not four (`TERR-194`); the former wording "Its only stores to `+0x18c` are four" is withdrawn in `retracted.md`. The 68 displacement store sites are classified by receiver by hand (Medium) in `TERR-194`: of the 45 that reach a drawable, five set bit 0, and the unit path's single setter is the hero arm (`UNIT-140`).

### ANIM-118

- **Unit.** `141.alm` (80 by 80) places six Humans-band records (`tools/campaign -mode place`, `-mode swing`). Record 0 is map unit 5 at cell 35,68, owner slot 3 `MageAlchemist`, flags `0x4`, class key 23, secondary key 0, definition id 517. Definition id 517 resolves to Humans row 208 `M141_MageTooth`: `typeID` 24, weapon cell Weapons row 13, no shield; the equipment chain of `HERO-APPEAR-042` also gives 24 (`evidence/census-en.txt`, RU identical).
- **Class.** Class 24 is `Human Mage` and class 23 `Unarmed Mage`. `Sound[0]` is 510 for 24 and 0 for 23, five elements each. The mage's attack fields are `AttackPhases` 6, `AttackDelay` 8, `ShootDelay` 8, `Projectile` 10 (`evidence/class-attack-en.txt`, RU identical).
- **Sound.** Registry slot 510 is `magic\arrow`, archive member `magic/arrow.wav`, 22446 bytes, the same hash on both roots (`evidence/sound-slot-510.txt`, from the `VIDEO-SFX-015` and `VIDEO-SFX-020` tables). The hook of `ANIM-105` is passed slot 510 and, when the buffer is non-null, plays it with a pan and an attenuation. Action 3 or 7 reaches it at tick 8 (`ANIM-STATE-023`).
- **Saves.** The owner's resave of the document, written by the original, and a second original-written save of the mission hold type word 24 for unit 5 at cell 35,68 (`evidence/save-typeword-mages.txt`). The class key would give 23, silent.
- **Same pattern.** 43 definition-id placements in 10 maps (41, 61, 80, 90, 100, 110, 111, 130, 141, 151) carry class key 23 with a Humans row of type 24, all mages; in mission 41 unit 40 is saved as 24 as well.

**Confidence.** High for the unit identity, row, class and registry join, and for the hook being passed slot 510: decoded tables joined by the arm of `ANIM-116`. Medium for the population and the saved words; that a sound plays is Unknown: a join of decoded tables, 2 actors in 2 maps.

**Unknown.** Whether `[L01636][510]` is non-null at run time. Whether the regular attack is action 3 or 7 (the actor reach, `actor+0x12c`; `ANIM-STATE-023`); both fire the hook at tick 8. Which sound a cast plays: it uses `vt+0x60` and `vt+0x5c`, outside this question.

## Open questions

1. **Opcodes `0x86` / `0x8a` / `0x8b` / `0x8c`** — *narrowed by `MAGIC-PIC-027`*: all four are
   cast messages, `0x86`/`0x8b` are one sender's unit- and cell-addressed forms and carry the
   picture id in `msg+0xc`, `0x8a`/`0x8c` are another's and carry none. What remains open is which
   routine calls each sender, and therefore which spells take the deferred path.
2. ~~**`R0558`**, `CProjectile`'s own driver, and `R0556`, its own body draw~~ —
   closed by `ANIM-PROJ-025` and `ANIM-PROJ-026`. A projectile is indeed a different animation
   model: its phase is not advanced by the animation tick but by a per-object countdown the
   spawner sets, and its behaviour is a `switch` on the picture id rather than on an action code.
3. **What `ANIM-TICK-011`'s stamp is for.** The 41×41 dword mask at `CMapView+0x17cc`, the
   `unit+0x102` guard and the `&8` player flag are all unread, so *whether* a shipped map ever
   satisfies the four-corner `0xc000` gate in play is open — and with it `TERR-SPR-042`'s animated
   object arm and `TERR-EDGE-024`'s partial-repaint renderer. This is a TERRAIN-ledger question and
   wants its own experiment.
4. **Action codes 2 and 4** — narrowed by `ANIM-STATE-002`'s amendment, still not shown unwritable.
5. **No runtime session.** The witnesses are now cheap, precise and **six**: a `Bee` standing idle must
   cycle four frames in **250 ms** at the default game speed and **500 ms** at the slowest, and must
   stop dead on pause (EXP-0084's instrument); and a `Human Swordsman` at `Speed 19` must cross a cell
   in **14 ticks ≈ 868 ms** while playing exactly **one** whole walk cycle, i.e. 1.14× the rate of a
   fire beside it, at **every** one of the nine game speeds (`ANIM-AMBIENT-016`). `ANIM-CLOCK-001`
   rests on the caller chain alone; one stopwatch settles all of them. [EXP-0111] adds three of
   the same kind, all on one blow: the numeral over a struck unit must be the colour of **its own
   owner** rather than of the attacker, must vanish **1000 ms** after it appears **at every game
   speed** while its drift slows with the speed, and a second blow inside that second must **raise
   the first numeral's value** rather than add a second numeral. The same blow's two sounds must be
   separable: the swing is heard `attackChargeTime - AttackDelay` ticks before the numeral, which for
   a `Human Swordsman` is **+3 ticks ≈ 190 ms** (`ANIM-CLOCK-024`), and the struck unit's grunt must
   not repeat more than once per **1500 ms** however fast it is being hit (`ANIM-SND-022`).
6. **Whether `unit+0x20` is ever negative** (`ANIM-ARM-018`) — the second walk-advance arm is
   transcribed and its reachability is argued, not shown. A classified `disp:20` write census over
   the drawable module would settle it; the immediate sweep already run cannot see a register store.
7. **Why the walk is distance-clocked at all** while every other timeline is tick-clocked
   (`ANIM-WALK-013`). The consequence is measured (`ANIM-AMBIENT-016`); no instruction states the
   intent, and none is expected to — recorded so the asymmetry is not later read as a defect.
8. **What the three colour ramps at `L02681` / `L02683` / `L02685` were for.**
   `ANIM-NUM-021` establishes that the `0x73` arm chooses one on every landed blow and reads it
   back nowhere; `R0572`, their only writer, was read at the three stores and at its packing
   idiom, not decoded. Whoever else consumes those arrays would say what the discarded choice
   would have coloured.
9. **`drawable+0x78`**, the flag that suppresses the numeral's draw while leaving the record alive
   and ageing (`ANIM-NUM-020`, `L03227`). A `disp:78` sweep puts most of its hits inside the map
   view's paint, on objects this experiment did not separate.
10. ~~**The three non-strike producers of `0x73`** — `R0564` (twice) and `R0565`
    (`ANIM-BLOW-019`). The client arm is the same for all five, so what is open is only whether a
    spell's numeral carries the same value law, not whether it appears.~~ — `R0565`'s own
    gate and value law are closed by `ANIM-074`: it notifies iff its own computed damage is
    nonzero, through the same builder and pool, with the same unread literal-zero `dmg` argument
    the strike uses, so the value law is the same by construction (one shared builder, not five
    separate ones). `R0564`'s two call sites remain addressed only by citation
    (`UNIT-AREADIRECT-072`, `UNIT-AREAHP-073`), not independently re-read — `ANIM-075` names this
    as its own scope boundary rather than closing it.

## State-projector trigger contract

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-129 | The 51 direct projector sites do not all emit state sync: the ordinary send requires the effective mask, masked with 0xff7fff7f, to be nonzero, while 0x80 and 0x800000 have separate sends. | High / Medium | ✔ promoted | [EXP-0475](../experiments/EXP-0475-hurt-projector-triggers/) |
| ANIM-130 | Token 6 sends dirty after restoration: health remaining nonpositive keeps stage, crossing zero adds mask 0x428 and stage 0; token 11's caster instead sends literal mask 1, omitting that stage reset. | High / Medium | ✔ promoted | [EXP-0475](../experiments/EXP-0475-hurt-projector-triggers/) |
| ANIM-131 | Four attachment projector sites send callback-produced dirty masks; a zero effective ordinary mask emits no state sync, and stage-1 native callback reachability remains unresolved. | High / Medium / Unknown | ✔ promoted | [EXP-0475](../experiments/EXP-0475-hurt-projector-triggers/) |
| ANIM-132 | The finite direct-site table separates command, projection, placement, award, stock and shop predicates from client stage; full-mask stage-zero and decay sends move old stage 1 away rather than repeat it. | High / Medium | ✔ promoted | [EXP-0475](../experiments/EXP-0475-hurt-projector-triggers/) |
| ANIM-133 | The three intervening client calls do not clear corpse stage at native bindings: campaign-record scalar getter, scalar price helper and a 12-byte copy into drawable+0xe4..0xef. | High / Medium | ✔ promoted | [EXP-0475](../experiments/EXP-0475-hurt-projector-triggers/) |

### ANIM-129

The current identical EN/RU executable has SHA256
`942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`.
The raw `.text` E8 scan to `R0059` reproduces 51 hits. Each is at an
instruction boundary in a linear Capstone 5.0.7 decode from the original-owned
inventory entry. A raw four-byte scan of the image finds zero entry-address
occurrences. The canonical `state-projector.md` table retains every direct hit,
including recipient recursion `L03228` and construction paths.

Recipient zero selects human players (`L03229..L03230`) or the actor's
owner (`L03231..L03232`) and recurses with the same object and mask.
For a nonowner, kind/class predicates reduce mask using `0x507b`, `0x50fb`
and a possible `0x1000` clear (`L03233..L03234`). Maximum mana zero
clears `0x2` (`L03235..L03236`). Stage bit 8 survives these reductions.

The ordinary send `L03237` requires the effective mask, masked with `0xff7fff7f`,
to be nonzero (`L03238..L03239`). Mask `0x80` calls the separate routine
`R0669` at `L03240`; mask `0x800000` calls `R0670` at
`L03241` only for the owner. Mask 0, and mask `0x80` alone, do not emit an
ordinary state sync. Dynamic dirty and nonowner reductions must be evaluated
before inferring repeated-stage receipt. This narrows ANIM-126's any-mask
omitting-stage wording; its death, bleed and entry facts stand.

**Confidence.** High for the named branches, 51 raw direct hits and finite
argument classification. Medium for completeness: Capstone plus inventory
navigation verifies these literal E8 sites, not computed calls, undiscovered
code or equivalent entry arithmetic. No new displacement/width field absence
claim is made; pointer-table scanning excludes literal pointers only.

**Unknown.** Computed callers, native scheduling, recipient visibility changes,
and callback dirty-mask effects outside the selected closure.

### ANIM-130

Token selector `token+8` dispatches through `L03242` at `L03243`.
Token 6 enters `L03244`. It requires the relationship gate,
signed health above -10 and nonzero `actor+0x98` (`L03245..L03246`).
It clears dirty at `L03247`, calls `R0671` at `L03248` with a
restoration capped to maximum minus current health, then projects the target's
dirty to recipient zero at `L03249`.

`R0671` ORs dirty with 1 (`L03250..L03251`), records initial
health <=0, adds the signed word delta and caps at maximum. A crossing to
positive health calls `R0672` (`L03252..L03253`). That helper ORs
dirty with `0x428` at `L03254` and writes stage 0 at `L03255`, agreeing
with HERO-REVIVE-068. Thus token-6 restoration still <=0 supplies no stage
bit and can preserve a client's stage 1; crossing zero sends stage 0.

Token 11 enters `L03256`: actor target, positive
`min(request, target.health+10)` at `L03257..L03258`. It subtracts the
amount and sends target mask 1 at `L03259`. It restores the caster through
the same helper, then sends literal mask 1 at `L03260`. That literal omits
a stage reset even if the helper produced one. Native caster-stage admission
is not established by this local argument.

Unicorn 2.1.4 executes the original local health and dirty/stage slices.
Health -5, max 100, stage 1, dirty 0 with deltas 3/5/6 produce
`(-2,1,1)/(0,1,1)/(1,0x429,0)` for health/dirty/stage. The position call,
subsequent derive and registration are outside those controls.

**Confidence.** High for local predicates, literal masks and the restoration
split. Medium for native repeated-stage feasibility: the tested local slices
and selected token body do not establish every upstream actor admission,
callback side effect or message order.

**Unknown.** Native occurrence and frequency, upstream target-kind and caster admission at stage 1,
and indirect callback/alias effects outside the selected closure.

### ANIM-131

Four sites send `actor+0x150` to recipient zero. `R0673` accepts an
attachment when its `+0x3d` intersects byte `[L04012]` or `[L04013]` through
`R0674`. At `L03261` the continuous flag `[L04013]` is set and
duration is divisible by 8; dirty is cleared before slot `0x40`.
At `L03262`, duration <=9600 decrements to zero, the continuous flag is
clear, and dirty is cleared before slot `0x44`.

`R0612` clears dirty at entry. Site `L03263` follows removal of an
opposite pair of effect types `0x17/0x1b` and its slot `0x44` callback.
Site `L03264` follows the other application path: same-type refresh or
replacement, construction/application of a new attachment, or a direct
non-timed callback through slot `0x40`. None of these four supplied masks is
a constant stage-free send. ANIM-129's effective-mask gate decides whether
an ordinary state sync is sent at all.

**Confidence.** High for the four sites, local flags, duration gates and mask
provenance. Medium for finite path classification beyond the callbacks.
Unknown for their native repeated-stage reachability; indirect callbacks and
their dirty/stage contribution are not exhaustively enumerated here.

### ANIM-132

The canonical direct-site table retains the original local gates for every
command, projection, placement, skill-award, stock and shop site, including
construction and unresolved paths. It assigns numeric opcodes from byte
table `L00564`, dword table `L03265` and bias 2; command byte `+4` must
be zero. It invents no command name.

`L03266` supplies full mask after a server-stage-zero check
(`L03267..L03268`); old client stage 1 would move to 0. Other string
branches in that handler bypass this check. Site `L03269` supplies mask
`0x1f001304`, omitting stage; its local checks are lookup, index 0..5, cost
and actor slot `0x30`, not client stage. `L03270` ORs local mask bits with
dirty; the local bits omit 8 but helper dirty contributions remain unresolved.

The full-projection wrapper `R0675` requires human recipient, actor
kind and `R0676` true. The last helper returns true exactly when
`actor.word18 & recipient.word2c` is zero. Its direct routes include actor
event `0x73`, player-list message broadcast and player-list re-entry projection;
clearing/re-entry lifetime is not established. Successful placement at
`L03271` and `L03272` sends mask `0x20`; neither local body tests stage.

Decay `L03273/L03274` sends mask 9 only when its stage changes. Its
explicit threshold stores supply 2/3/4 or 5, so old client stage 1 moves
away. Token 25 likewise writes stage 5 before `L03275`, then separately
constructs the actor sent at `L03276`. These are not repeated-stage sends.
Teardown `L03277` supplies mask `0x80`, which is a separate send under
ANIM-129. Construction paths do not prove an existing actor's reachability.

**Confidence.** High for the named local gates, arguments and explicit stage
stores. Medium for finite per-site interpretation and completeness; outer
command, shop, placement and collection admission is not a native play proof.

**Unknown.** Upstream actor admission, opaque derive/load helpers, dirty/stage
aliases, previous drawable/id reuse and message scheduling.

### ANIM-133

The state-sync arm saves old stage at `L02542`, applies stage under mask 8
at `L02541`, and switches at `L02764` (ANIM-095).

The call `L02787` receives the embedded record at campaign screen `+0x548`.
Dispatcher frame local −0x10 comes from the application singleton's slot
`0x7c` (`L03278..L03279`), not the dispatcher's incoming receiver.
The campaign constructor calls record constructor `R0677` for `+0x548`
at `L03280..L03281`. That constructor writes vptr `L10123` at `L03282`.
Slot `L13135` points to `R0678`: returns the 32-bit field at offset 4 of the object, no store or
call. REG-SCN-064 supplies the embedded-record native association.

`R0589` at `L03283` receives two scalar arguments and has no store
outside its stack. `R0544` at `L03284` receives destination
`drawable+0xe4`, source cursor and length 12. Its count-controlled copying
ends at `drawable+0xef`, before corpse stage `+0x15a`. Sixteen original-code
Unicorn controls cover four destination alignments and equal, forward-overlap,
backward-overlap and disjoint pointers: every observed destination write stays
in the 12-byte range and a stage-1 sentinel survives. No callback is called.

These three native bindings do not clear the stage between receipt and switch.
This resolves the three-helper Unknown in ANIM-095 within that receiver scope.

**Confidence.** High for explicit scalar getter/price-helper stores and the
copy's supplied length, destination and tested writes. Medium for the native
receiver closure: replacement singleton objects/vptrs and malformed pointers
are outside the selected original constructor association.

**Unknown.** Alternative receiver bindings, malformed aliases, callback paths
elsewhere in the arm and native message ordering. This is not a global stage
writer or audibility census.

## Turn message and drawn turn

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ANIM-134 | Turn message 0x6d has one found sender, the action-1 arm of the actor tick; it goes when the tick left the position unchanged and changed the desired facing byte, with (desired+8)>>4 and the estimate mover+0xa4. | High / Medium | ✔ promoted | [EXP-0498](../experiments/EXP-0498-hero-turn-rate/EXP-0498.md) |
| ANIM-135 | The client 0x6b and 0x6d arms apply a message while a run is active, replacing it after an Overriding log line; only the 0x72 shoot arm returns without applying it. | High | ✔ promoted | [EXP-0498](../experiments/EXP-0498-hero-turn-rate/EXP-0498.md) |
| ANIM-136 | A turn is drawn standing at a 16-way facing moving from the drawn facing to the message target over the message count: one tick for a snap with server flag +0xa0 clear, ceil(arc/rate) ticks otherwise. | High / Medium | ✔ promoted | [EXP-0498](../experiments/EXP-0498-hero-turn-rate/EXP-0498.md) |

### ANIM-134

- The actor tick samples the two position words (`R0165`/`R0166`: cell byte times 256 plus sub-cell byte), `mover+0` and `mover+1` (`L04563`..`L13307`) before the executor call at `L00086`. Action value 1 selects `L01777` (table `L01775`), which calls `R0550` with those samples; that call at `L04571` is its only direct reference.
- When both words are unchanged, `R2134` compares the sampled desired byte with the current `mover+1`. When they differ it returns `mover+0xa4`, and `L13308`..`L03037` builds 0x6d with facing `(mover+1 + 8) >> 4` (0..16) and that count. The builder `R0622` writes opcode, facing at `msg+0xc`, count at `msg+0xd`, and `actor+0x138` = clock + count.
- When a word changed and both sampled sub-cell bytes were 0x80, the arm builds 0x6b with `2 * mover+0xae` and `mover+0xaa` instead.
- So one 0x6d goes per new desired byte of an action-1 actor, in the sub-tick of the turn call; a continuing turn sends none. A fresh snap sends count 1. A tick that ends with an action other than 1, such as after the stop reset, sends no turn message.
- Search: among four raw `push 0x6d` sites in `.text`, `L13308` is the only one feeding a message; the other three feed `R0668` and `L11280`. No byte-immediate store of 0x6d to `+9` exists. `R0622` has three direct callers, `L03036` (0x71), `L03037` (0x6d) and `L03038` (opcode 0, written as 0x6b).

**Confidence.** High for the arm, both predicates and the fields. Medium for the sender being the only one: the scans cover immediate pushes and byte-immediate stores, not an opcode built in a register.

**Unknown.** Message delivery latency between the server tick and the client's next presentation tick.

### ANIM-135

- 0x6d arm: `L13309` tests `+0xa0`. When it is nonzero the arm formats "Overriding '<action>' by 'Turn'. %d segments lost." by the running action (strings `L13310`..`L13311`), logs it only when `[L04662]` is nonzero, and continues to `L13312`. There it writes `+0xa0` = `msg+0xd`, `+0x84` = 5, `+0x85` = `msg+0xc`, `+0xbc` = `+0x6c << 4`, and zeroes `+0x94`, `+0x9c` and `+0x98`.
- 0x6b arm: the same structure at `L13313`..`L13314`, with "... by 'Move' ..." strings at `L13315`..`L13316`.
- 0x72 arm: `L02519` jumps away when `+0xa0` is nonzero; that path formats "... by 'Shoot' ..." and ends at `L13317` without writing the action block.
- A turn or move message therefore replaces a running attack, shoot, cast, move or turn run on the client at once.

**Confidence.** High: each branch and store is a named instruction and the strings are read from their addresses.

### ANIM-136

- Driver turn arm `L02709`..`L02536`, once per presentation tick with n = `+0xa0`: d = (`+0x85` << 4) − `+0xbc`, folded into (−128, 128]; `+0xbc` += d / n, truncated toward zero; a negative result adds 256; `+0x6c` = `+0xbc >> 4`. The tail decrements `+0xa0` (`L02481`). The last tick, n = 1, lands on the target exactly.
- The draw takes the standing frame at the whole 16-way facing during the turn; move and attack frames halve it to 8 (`ANIM-DIR-006`). The server byte has 256 steps; the client sees only the message target rounded to a sixteenth.
- Composed with `MOVE-105` and `ANIM-134` (`evidence/turnsim.txt`), drawn facings per tick from 0: rate 16, 8/16: 1 2 3 4 5 6 7 8; rates 19 and 20, 8/16: 1 2 3 4 5 6 8; rate 23, 8/16: 1 2 3 5 6 8; rates 16..21, 4/16: 1 2 3 4; any rate, 1/16 or 2/16: one tick.
- A turn of up to two sixteenths, which includes every one-heading change of an 8-way walk, is drawn as a one-tick change of facing when the server active flag `mover+0xa0` was clear at the turn call and the client run is not replaced. With the flag set the call takes the leaf and sends its count: after the stop reset of `MOVE-107`, a 2/16 request at rate 16 sends count 2 and is drawn 1, 2 (`evidence/turnsim.txt`, `A32:2`).

**Confidence.** High for the conditional arithmetic and the drawn sequences given a delivered, uninterrupted message run. Medium that server and client ticks stay aligned through a turn: delivery order in the session loop was not read.

**Unknown.** Rendering of a frame between presentation ticks; observed play.
