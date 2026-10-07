# Actor effect marks

[Reference](format.md)

## Actor effect marks

The mark layouts and initial builders are specified by MAGIC-MARK-059,
MAGIC-MARK-060, MAGIC-MARK-061 and MAGIC-ACTOR-066. The single-figure account of MAGIC-SHIELD-063 is superseded: Shield
concatenates two separately rasterized components, A then B. Its constants
and 90-step cycle are retained; use MAGIC-091/MAGIC-092 for coordinates.

Presentation follows `MAGIC-089` through `MAGIC-093`, `ANIM-044` through
`ANIM-047` and `SPR16A-031`.
The all-unit late-pass wording of `ANIM-047` is partially retracted: its late
shadow/body population is CAirUnit; ordinary and alternate CUnit routes occur
earlier (`ANIM-CATEGORY-084`, `ANIM-AIRPASS-086`). The mark-builder and actor-local
before/after-body order below is retained; no native overlap result is added.

### The two states an actor's marks are held in

The presentation side of an actor holds two arrays.

```
effect list     unit+0x124 array object, unit+0x128 data, unit+0x12c count
                one dword per element:  bits 31..16 kind, bits 15..0 countdown
mark array      unit+0x110 array object, unit+0x114 data, unit+0x118 count
                8 bytes per record:
                  +0x00 i16 dx      horizontal offset
                  +0x02 i16 dy      vertical offset
                  +0x04 i16 depth   vertical offset, and selects the draw pass
                  +0x06 u8  record  index into the projectile record array
                  +0x07 u8  phase   frame index for the blit
```

A `kind` is `2*spellId + 8`, the cast half of the [picture pair](presentation.md).

### Lifecycle

The simulation opens and closes each element with a dedicated message. On attach, the effect's spell
id sets a bit in `actor+0x144` and, when that spell id is non-zero, an attach message is sent
carrying the target's runtime id and the effect's picture field. On expiry the bit is cleared and a
detach message is sent. A recast of the same spell evicts the old effect first, so it sends detach
then attach.

The client creates the element with countdown `0xffff` on attach, replacing any element of the same
kind, and removes it on detach. The countdown falls by one on every rebuild and the element is
removed when it reaches zero, but at `0xffff` that does not happen in play: the countdown is the
phase clock, and the element's lifetime is the effect's.

A third writer exists. The projectile driver dispatches an arriving projectile's picture over the
closed range 13..64, and two of its arms write the same list. Pictures 20 `healing` and 30 `Drain`
replace or append an element with countdown **32**, which does expire on its own. Pictures 18, 24,
28, 40, 44, 48, 52, 54, 56, 62 and 64 append one with countdown `0xffff` and no by-kind check, but
the [flight-length switch](presentation.md) gives all eleven a length of 0, so on the cast path the driver
returns before reaching that arm.

`heal` and `drain_life` therefore run from a 32-rebuild element rather than from an attach message.
Their builders use countdown only as the unsigned fresh-spawn gate; their motion magnitude comes
from the separate constant argument `32T` (`MAGIC-089`, `MAGIC-090`).

### Rebuild

Every rebuild discards the mark array and re-derives it. For each element the kind selects an arm;
kinds outside `0x12..0x3e` and 35 of the 45 kinds inside it select an arm that appends nothing. Ten
kinds reach seven builders:

| kind | spell | records appended | phase input |
|---|---|---|---|
| 0x12 0x1c 0x28 0x34 | the four Protections | 1 each | countdown mod 6 |
| 0x14 | heal | carry own phase below 7; append 3..5 while countdown >= 8 | `32T`, countdown gate |
| 0x18 | poison_cloud | 1 | countdown mod 6 |
| 0x1e | drain_life | carry own phase below 7; append 3..5 while countdown >= 8 | `32T`, countdown gate |
| 0x2c | shield | two independently rasterised components | countdown mod 90 |
| 0x36 | bless | 20 | countdown mod 5 |
| 0x3e | curse | 20 | countdown mod 5 |

`haste`, `slow`, `stone_curse`, `invisibility`, `freezing_cloud`, `light`, `darkness`, `lightning`,
`prismatic_spray`, `acid_stream`, `wall_of_earth`, `meteor_storm`, `control_spirit`, `teleport`,
`fire_arrow`, `fire_ball`, `wall_of_fire` and `fire_sacrifice` append nothing. Each element is
dispatched independently, so an actor carrying several effects shows the union of their records.

Every builder scales its lengths by `T`, the actor class's `units.reg` `TileSize`, default 1.

### Drawing

The actor draw walks the mark array twice: once before the actor's own sprite, drawing only records
whose `depth` is positive, and once after, drawing only records whose `depth` is not positive. Both
compute

```
x = anchorX - Width/2  + dx
y = anchorY - Height/2 + dy - depth
```

where `Width` and `Height` are the projectile record's [centering halves](presentation.md) and the anchor
is the actor's own draw position. Both blit the sheet at `record` with frame `phase`.

The `depth` field is therefore both a vertical offset and a z-order selector, which is how a mark
set can enclose the actor: records with a positive vertical term pass behind the sprite and records
with a non-positive one draw in front of it.

### The four Protections

One builder serves all four and appends exactly one record with `depth = 0`, `record = kind`, and

```
protection_from_fire   dx = -6  dy = -(T*32)
protection_from_water  dx = +6  dy = -(T*32)
protection_from_air    dx =  0  dy = -(T*32) - 6
protection_from_earth  dx =  0  dy = -(T*32) + 6
```

Four simultaneous Protections therefore stand at the four corners of a diamond of half-width 6
pixels, centred `T*32` pixels above the actor anchor. Nothing in the builder recognises that case;
it is the union of four independent placements.

### shield, bless and curse

`shield` concatenates two independently rasterised kind-44 components, A then B. The common midpoint
routine starts `x=0,y=r,err=2-2r`, emits sampled symmetry points in order
`(+x,+y),(-x,+y),(+x,-y),(-x,-y)`, including axis duplicates, and emits nothing at radius zero.

For `p=countdown%90`, Component A rasterises radius `16T` at interval 4 with `theta=4p` degrees:

```
q = trunc(y*28T/16T)
f = abs(4-trunc(abs(y)*5/16T))
A1 = (trunc(x*cos(theta)), -q, 11T+trunc(x*sin(theta)), phase f)
A2 = (-trunc(x*sin(theta)), -q, 11T+trunc(x*cos(theta)), phase f)
```

Component B computes

```
c45 = binary64(bits 0x3f96c16c16c16c17)
x1 = round53(p*c45)
x2 = round53(x1-1)
x3 = abs(x2)
x4 = round53(x3*float32(28T))
E = trunc(x4)
F = 4-floor(abs(p-45)/9)
R = trunc(sin(acos(E/(28T)))*16T)
```

and rasterises `R` at interval `2F`. Each `(u,v)` emits
`(u,-(E+11T),v,phase F)` then `(u,E-11T,v,phase F)`.
At post-prologue stack base `P`, `[P+0x30]` contains `16T`; `[P+0x34]`
contains `E`. After the interval push, the multiplication reads `16T` from stack offset +0x34 and
`E` is at stack offset +0x38.

The envelope operations are ordered: load integer phase; multiply by the
stored qword nearest `1/45`; subtract 1; take absolute value; multiply by
float32 `28T`; truncate through x87. Normal execution uses PC53,
round-to-nearest (control word `0x027f`) after each arithmetic instruction.
`__ftol` temporarily changes rounding-control bits and restores the saved
word. The stored qword is exactly `6405119470038039/2^58`; preserving this
rational across operation boundaries or replacing multiplication by division
by 45 changes outputs. For `T=3`, `E=56` at `p=15`, `27` at 30 and `28` at 60.

At `p=0`, `F=-1` and `R=0`, so no invalid frame is emitted. It is the sole
zero-output phase for installed `T={1,2,3}`. At `p=45`, `E=0` and `R=16T`.
Every emitted B record uses frame 0..4. The builder reads neither RNG nor
previous marks. — MAGIC-091, MAGIC-092 (amended and partially retracted in the ledger)

`bless` and `curse` each append 20 records: five trail steps of four records, on a circle of radius
20.0 centred `T*32` above the anchor. Angles are in degrees. `bless` steps 89, 71, 53, 35, 17 and
`curse` steps 0, 18, 36, 54, 72, so the two trails sweep opposite ways as the countdown falls. The
frame index of a step is `4 - step/18`, so five different frames are visible at once. A sine term
goes into `depth` rather than into `dy`.

### poison_cloud, heal and drain_life

`poison_cloud` appends one record at `dx = 0`, `dy = -(T*32)`, `depth = 0`, `record = 0x18`.

`heal` and `drain_life` first walk the previous mark array and carry their own records whose phase is
below 7. They increment the phase, then Heal subtracts and Drain adds
`s=trunc_toward_zero(1+32T/7)` to `dy`; shipped `T={1,2,3}` gives `s={5,10,14}`. The step is
constant throughout the effect. A phase-6 record is copied as visible phase 7, and phase 7 is then
dropped, so a cohort has eight outputs.

While unsigned countdown is at least 8, each builder appends `rand()%3+3` fresh particles. Every
particle consumes one further `rand()%360` angle `a`:

```
heal   dx=trunc(cos(a)*16T), dy=0,    depth=trunc(sin(a)*8T), record=0x14, phase=0
drain  dx=trunc(cos(a)*16T), dy=-32T, depth=trunc(sin(a)*8T), record=0x1e, phase=0
```

Countdown 32 through 8 produces 25 cohorts and 75–125 particles in total. Eight cohorts overlap in
steady state, 24–40 particles. One angle supplies both ellipse coordinates; there are not two
independent position samples (`MAGIC-089`, `MAGIC-090`).

### Two spells that change the actor instead

`stone_curse` and `invisibility` reach no mark arm. The actor draw looks their kinds up directly.

`0x30` gates two separate arms in `R0552`, each additionally conditional on
`byte [drawable+0x15a] <= 2`. The first, at `L05737`, discards the frame index the per-state arm
computed and substitutes `(drawable+0x6c - 8) & 0xf`, the facing at full 16-way resolution, mirrored
above 8. The frame therefore stops following the animation clock and changes only when the unit
turns. The second, at `L05815`, replaces the shade table the blit takes: element 16 of the array
at `[L05816]` for a class whose `Palette` is 0, otherwise a table built per draw from the
class's own palette with `R1107(palette, 0x10, mode 5, 0)`, mode 5 being the greyscale ramp
`(r+g+b) * level * 2 / 3 / levels`. Both forms draw the unit grey. `R0553` carries the frame
arm a second time at `L05813` (`MAGIC-STONEDRAW-084`).

`0x26` gates the actor's whole sprite on a per-player bit. That bit is bit 3 of the 32-entry
`CWordArray` row at `CPlayer+0x34` on the local player, indexed by the drawn unit's owner index, and
it is set on exactly one entry — the local player's own (`UNIT-VISBIT-044`). The test is therefore
an ownership test: an invisible unit is drawn to its owner through a different blit slot and is not
drawn at all to any other client (`MAGIC-INVISOWN-085`). The map draw applies the same test before
drawing a unit at all.

The by-kind lookup `R0615` has fourteen call sites over six routines and no consumer outside
the presentation layer; three of the six are consecutive vtable slots of the client unit class,
`R0552`, `R0553` and `R1106` (`MAGIC-EFFLOOK-083`).

### Customisation limits

The kind range, the dispatch table, the seven builders, the 6-pixel Protection radius, the four-way
Protection assignment, the `T*32`, `T*28`, `T*16`, `T*11` multipliers, the 20-mark trail, the
18-degree step, the 20.0 radius, the 90-step and 6-step and 5-step moduli, the eight-output particle
life and the 8-byte record are all engine code or `.text` immediates. Changing any of them changes
`ROM.EXE`. `TileSize`, the sheets behind each record index and their `Phases`, `Width` and `Height`
are data and can be replaced without touching the executable. A sheet with more phases than its arm
indexes gains nothing: `poison_d` declares 8 phases against 6 the engine can index, and `Drain`
declares 9 against 8.
