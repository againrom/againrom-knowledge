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

Settled by `MAGIC-CASTSPAWN-033` (amended and partially retracted in the ledger), `MAGIC-BURSTLIFE-034` (amended and partially retracted in the ledger), `MAGIC-DELIVER-035`. Section 11
describes the message; this section describes the object. The two are built by different code at
different times.

**The cast object is built by the caster, not by the message.** The even picture puts the caster
into the [cast action](casting.md). On the tick `actionphase == ShootDelay` the caster's own
`vt+0x5c` runs and allocates one `0x14c`-byte projectile:

```
picture         = 2*spellId + 8          taken from the caster's actionspell field
position        = cached caster centre plus the class offset defined below
actionx/y/z     = the target's current position, or the caster's own aim point
action          = 1
actionphase     = 0
actionsegments  = a hard-coded switch on the picture id (below)
```

`picture == 60` copy-constructs a second projectile. Its raw `+08/+0c`
is copied target `+88/+8c` plus the Selection fallback. Its construction
cache `+28/+2c` still holds the first object's source-derived point; the
driver later copies raw coordinates into it. The caster-centre placement
clause of MAGIC-CASTSPAWN-033 is partially retracted. Native first draw and
external geometry/draw ordering remain Unknown. — MAGIC-265

### Cast origin and human equipment

The normal CUnit/CAirUnit cast producer uses class ID `caster+20` and
facing `caster+6c`. Let `A=(caster+58,caster+5c)`, the cached footprint
centre, and `i=(facing-8)&14`. Its exact integer formula is:

```text
nonempty ShootOffset and picture != 60:
    x = A.x + 8*(ShootOffset[i]   - CenterX)
    y = A.y + 8*(ShootOffset[i+1] - CenterY)
otherwise:
    x = A.x + trunc((SelectionX2-SelectionX1)/2) - CenterX
    y = A.y + trunc((SelectionY2-SelectionY1)/2) - CenterY
```

The fallback is unscaled and does not add SelectionX1/Y1. Both raw
`+08/+0c` and cached `+28/+2c` receive the first object's result.
A is `P28/P2c + 128*(TileSize-1)`, not necessarily raw caster `+08/+0c`.
The complete origin window reads no animation frame or equipment pointer;
the action clock determines when it runs. — MAGIC-261

In `R0551`, human state 8 uses the idle weapon/shield name selection.
The selector's branches and name-to-ID stores are High. Its class-store path
maps an unshielded mage with empty hands to `mage` (23) and staff categories
to `mage_st` (24). At ShootDelay, the held class matches this selection when
the selector last wrote it for the current equipment and no other `caster+20`
writer ran since. The writer population is unenumerated: hero creation writes
`+20`, and the selector can return without a write. Equipment freshness and changes
during wind-up remain Unknown. Sword, axe, club and pike classes have empty
arrays; bow class 14 and crossbow class 15 have arrays. State 6's weapon
bypass is the death control. Body art and registry geometry are separate.
— MAGIC-262

All sixteen shipped human classes have Center `(64,78)`, Selection
`(48,48,80,90)` and TileSize 1. Normal pictures other than 60 produce
these deltas from A, in fine integer coordinates, with positive y south:

| Direction | Facing | mage / xbowman (23/15) | mage_st (24) | archer (14) | Empty array |
|---|---:|---|---|---|---|
| N | 0 | (32,-288) | (88,-352) | (8,-384) | (-48,-57) |
| NE | 2 | (128,-224) | (216,-256) | (168,-304) | (-48,-57) |
| E | 4 | (136,-128) | (216,-104) | (256,-144) | (-48,-57) |
| SE | 6 | (72,-40) | (88,8) | (184,24) | (-48,-57) |
| S | 8 | (-56,-24) | (-96,8) | (-24,96) | (-48,-57) |
| SW | 10 | (-152,-96) | (-224,-112) | (-224,16) | (-48,-57) |
| W | 12 | (-160,-200) | (-216,-256) | (-288,-160) | (-48,-57) |
| NW | 14 | (-96,-280) | (-80,-360) | (-192,-312) | (-48,-57) |

Odd facing values share the preceding even pair. Teleport uses `(-48,-57)`
for all human classes. Integer values and direction labels are High;
the eight-fine-units-per-map-pixel interpretation retains Medium confidence.
— MAGIC-263

Within the twelve shipped Units rows with positive spell slots,
Goblin_Sling.4 (79), Orc_Bow.4 (65) and Bat_Sonic.4 (70) have arrays.
The bat's pairs equal Center and give zero non-Teleport delta; the other
nine rows use their class Selection fallback. This stored population does
not establish visible-cast reachability for every row. — MAGIC-264

The searched delivery population is all 28 Spells rows, the normal producer
and three client cell-effect arms. Normal active pictures 10,12,20,30,34,36,60
share the formula; Teleport forces the fallback. The odd-picture 0x86 arm
uses message cell `+0d/+0e`; source-cell 0x8b uses `+0a/+0b`; source-cell
0x8c picture 36 uses packed source word `+0e`. Each writes cell*256+128.
Global producer completeness and native first-visible-frame position remain
Unknown. — MAGIC-265

Visible SAVE records current `Prj<ID>/x` and `/y`, without a separate
immutable launch-point leaf. Simulation transport has a separate caster
Position. Native SAVE/LOAD continuation remains unobserved.
— MAGIC-266, MAGIC-267

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
