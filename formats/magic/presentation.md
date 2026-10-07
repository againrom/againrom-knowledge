# Icons, cast presentation and map objects

[Reference](format.md)

## Spell icons

Both are settled by `MAGIC-ICON-024`, `MAGIC-ICON-025`, `MAGIC-PIC-026`, `MAGIC-PIC-027` (amended and superseded in part in the ledger).
They share a spell id and nothing else.

**The icon is a position, not an index.** `graphics\interface\SpellBook.bmp` is 480x85, 24bpp,
and the panel blits it **whole and once**. Nothing in the engine cuts it. What is per-spell is a
mask: 24 cells at

```
x = xBase + 6 + 38 * (slot % 12)
y = yBase + 6 + 38 * (slot / 12)      slot = 0 .. 23, cell 36 x 36
```

and every cell whose known-spell bit is **clear** has `SpellBack.bmp` (36x36) pasted over it. The
selected slot gets a 36x36 outline; four hot-key slots get a numeral. The
book's slot-to-spell mapping is a fixed 24-entry table:

```
slot   0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23
id     1  2  3  4  5 23 24 16 15 14 13 12  6  7  8  9 10 25 26 22 21 20 19 18
```

**Ids 11, 17, 27 and 28 — `drain_life`, `darkness`, `curse`, `slow` — are absent**, so four spells
have no cell in this book. `main.res::text/spell.txt` is 28 lines keyed by id;
`main.res::text/spells.txt` is 24 lines keyed by slot; each line is the first line of that
cell's hover text, to which the hover getter adds the live values of the selected actors
(TEXT-080, TEXT-081).

The current spell and quick binding are zero-based cells. A normal book
command carries `cell+1`, not the intrinsic ID. DispatcherL00198-based
lookup translates that byte; statistics use the same table through base+4.
Thus cell5 travels as command 6 and resolves spell 23. The mapped ID indexes
the actor's sparse book. — AI-SPELLIDENT-286

Book availability is the OR of+18 over selected objects with nonnull+7c;
actual book-command members additionally need the chosen cell bit.
The shared Cast bit uses a narrower per-object `CUnit`, type 17h/18h and
nonzero-mask predicate under the session gate. It is not primary-only and
does not require all members to cast. Upstream snapshot freshness and full
mixed-selection runtime reachability remain Unknown.
— AI-SPELLPOP-287, AI-SPELLCAP-288

Mouse selection checks availability before storing the cell; a nonnegative
shortcut stores it before later arming checks. The book's target-production
flag uses the cell-indexed L00200 table, not intrinsic IDs. Item context
selects a different target table and opcodes 25h/26h: command+10 is then a
word-sized inventory position, and the item effect supplies the constructed
Spell's ID. This is a separate identity route, not another book-table lookup.
— AI-SPELLGUARD-289, AI-SPELLITEM-290

**The map picture is arithmetic.** No `Spells` column names it. A cast's picture id is

```
picture = 2 * spellId + 8        the homing / attached form
picture = 2 * spellId + 9        the burst / area form
```

and that value indexes `projectiles.reg` by its **`ID`** key. The parity is load-bearing: an even
picture creates no map object at all — the caster is looked up and given the cast action — while an
odd one creates a projectile. Two further message opcodes create one unconditionally. A picture id
with no `projectiles.reg` row is dropped silently at both the spawn and the draw, which is why
`light`, `invisibility`, `darkness`, `stone_curse`, `haste`, `control_spirit` and `slow` put nothing
on the map, and `fire_ball` and `poison_cloud` own two pictures each.

**What a consumer must not do:** treat the icon as indexed art (it is a strip and a mask); treat
`Delivery System` or `Distribution system` as the picture selector (they are not); or give a spell
an id-independent sprite (the formula pins the two together, and re-pointing one spell's art means
editing `projectiles.reg`'s `ID`).


## Cast presentation and application tick

Settled by `MAGIC-CASTANIM-029`, `MAGIC-CASTTICK-030`, `MAGIC-BURST-031`, `MAGIC-SENDER-032`. (Delivery/timing clause narrowed by MAGIC-CASTCLOCK-171.)
Section 10's parity rule is unchanged; this section says what each parity is *for*.

**Every cast animates the caster.** The cast routine sends the picture message at the **start** of
the wind-up, carrying `2*spellId + 8`, which is even for every id. The client's response to an even
picture is not "draw nothing": it finds the **caster** and sets

```
unit.actionCode   = 8        the cast action
unit.actionClock  = 0
unit.ticksLeft    = len(expanded Attack timeline)     units.reg, the same run a melee attack uses
```

skipping the whole branch if the unit is already animating or if its class has `AttackPhases == 0`.
The action then advances one frame per game tick with **no modulus** — frame index equals the tick
index — and ends when `ticksLeft` reaches 0.

**The two clocks.** They are separate numbers and the original does not reconcile them.

```
simulation   windUp   = 8 ticks   (actor field; an equipped item may override it)
             recovery = 4 ticks   (same)
             the effect payload is prepared on tick windUp, starting tick counted as0;
             a melee strike lands on that same tick, in the other branch of one if

client       the visible projectile is spawned on tick ShootDelay   (units.reg)
             a melee hit reaction fires on tick AttackDelay         (units.reg)
```

A consumer that wants the swing and the application to coincide must choose one. The simulation's
tick is the one that changes hashed state.

**Which spells put a picture on the map.** The apply builds one of two effect classes, chosen by
the `Distribution system` column from [effect application](application.md):

| Column value | Class | Sends a map picture |
|---|---|---|
| 1 | `PointEffect` | no |
| 3, 4, 5 | `AreaEffect` | yes — one `2*spellId + 9` message per cell, per ring |

Shipped `Data.bin`: **18 spells are `PointEffect` and 10 are `AreaEffect`** — `fire_ball`,
`wall_of_fire`, `fire_sacrifice`, `freezing_cloud`, `poison_cloud`, `acid_stream`, `light`,
`darkness`, `wall_of_earth`, `meteor_storm`. Of those ten, `light` and `darkness` compute an id
`projectiles.reg` does not define, so eight spells draw a spreading picture. The `AreaEffect` walks
a cell pattern with its own arms for `fire_sacrifice`, `acid_stream` and `meteor_storm` and a
default for the rest.

**What a consumer must not do:** read `MAGIC-PIC-027` (amended and superseded in part in the ledger)'s even-parity rule as licence to draw nothing
on a cast. Every cast of every spell animates its caster; only the map object is parity-gated.


## SpellEffect map objects

Settled by `MAGIC-CASTSPAWN-033`, `MAGIC-BURSTLIFE-034` (amended and partially retracted in the ledger), `MAGIC-DELIVER-035`. Section 11
describes the message; this section describes the object. The two are built by different code at
different times.

**The cast object is built by the caster, not by the message.** The even picture puts the caster
into the [cast action](casting.md). On the tick `actionphase == ShootDelay` the caster's own
`vt+0x5c` runs and allocates one `0x14c`-byte projectile:

```
picture         = 2*spellId + 8          taken from the caster's actionspell field
position        = the class's muzzle point for facing (dir - 8) & 0xe,
                  or the class bounding-box centre when the class has no muzzle table
                  or the picture is 60 (teleport)
actionx/y/z     = the target's current position, or the caster's own aim point
action          = 1
actionphase     = 0
actionsegments  = a hard-coded switch on the picture id (below)
```

`picture == 60` allocates a **second** projectile, copy-constructed, at the caster's own centre —
teleport draws two sprites.

**The flight length is an engine table, not data.** 51 index bytes over pictures 10..60 through an
8-entry jump table. Seven ids are non-zero:

| picture | sheet | spell | actionsegments |
|---|---|---|---:|
| 10 | `firebolt` | Fire Arrow | `distance / 200` |
| 12 | `fireball` | Fire Ball | `distance / 384` |
| 20 | `healing` | Heal | 1 |
| 30 | `Drain` | Drain Life | 1 |
| 34 | `lightnin` | Lightning | 13 |
| 36 | `chain` | Prismatic Spray | 13 |
| 60 | `teleport` | Teleport | 21 |
| any other | — | — | 0 |

`actionsegments = 0` makes the driver return finished on its first tick, before its own picture
switch. So 21 of the 28 spells put an object on the map that never executes a driver arm.

**The burst object does not move.** The odd-picture client arm builds it at a cell and sets
`actionx/actiony` to its own `x/y` and `actiontarget` to 0, so every per-tick step divides zero. Its
lifetime is `msg+0xf`, chosen by the sender:

```
AreaEffect staged walker 16 ticks, raised to 18 on the acid_stream arm
                         a stage runs immediately, then every 3 ticks
the fire_ball sender     22 ticks
```

Two of the eight burst sheets are given `2 * Phases`, which is exactly one pass under the projectile
frame clock ([ANIM](../anim/format.md)): `acid` has 9 phases and gets 18, `fireexpl` has 11 and gets 22. The
other six run 16 ticks against the 22 or 30 a full pass would need.

**A burst also plays a sound**, id `500 + picture`, at volume `(10000 - screenDistance)/100`,
skipped for picture 51 (`Meteor`).

**Two flight-length rules exist and only one is normally used.** The simulation computes its own
value in the cast routine — 0, or `distance / Data.bin parameter 7` when `Delivery System` is 2, or
5 for spell ids 13 and 14 — and puts it in `msg+0xf`. On a normal cast that value is discarded,
because the even picture routes the message to the caster-animation branch and the projectile is
built later by the table above. The simulation's value is used only when the source has no
client-side runtime id, which rewrites the opcode to `0x8b`.

`Delivery System == 2` holds for exactly `fire_arrow`, `fire_ball`, `lightning` and
`prismatic_spray`, which are exactly the four spells the client's own table gives a travelling or
ramped arm, and exactly the four picture ids with a special draw arm. A consumer may treat that
column as "this spell throws something", but the number it produces is not the number the original
draws with.

## Simulation transport and presentation clocks

A visible CProjectile does not deliver the simulation payload. Delivery2 uses a
SpellTransport with a signed countdown: expiry enqueues its nested effect and
retires the transport. It does not move Position or poll a sprite collision.
The animation packet's five-tick value for IDs13/14 is separate from the
simulation transport's counter10. Prismatic victim preparation runs during
admission; ordinary phase5 preparation follows charge. Native first-frame
ordering and visible impact alignment remain unmeasured. — MAGIC-DELIVERY-170,
MAGIC-CASTCLOCK-171
