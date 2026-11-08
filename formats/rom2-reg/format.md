<a id="rom2-inline-registry-ya1-nested--identity-survey"></a>

# ROM2 inline REG store (`&YA1`)

Nested ROM2 registries use the inline REG envelope. It is distinct from the
tail-registry RES envelope despite the shared magic. Resource context and
record geometry select the parser. — R2-ASSET-006

<a id="result"></a>

## Layout

| Position | Width | Data |
|---|---:|---|
| 0 | 24 bytes | Six-dword header; magic `&YA1`, root fields, record count at 0x10 |
| 0x18 | 32×R bytes | Record array |
| 0x18+32×R | 4 bytes | Pool byte length |
| 0x1c+32×R | poolLen bytes | Pool |

`registryEnd = 24 + 32*R + 4 + poolLen`. This framing covers the preserved
nested registries. The installed `data/ai.reg` and `data/map.reg` lengths are
156 and 1020 bytes in both world archives; equal lengths do not prove equal
values. — R2-ASSET-006

## Read and write order

Read the six header dwords, R records, u32 pool length and pool bytes.
The corresponding structural emitter writes those same regions in order.
Raw record and pool values must remain opaque where ROM2 type/lookup semantics
have not been established. [ROM1 REG](../reg/format.md) supplies a layout
cross-reference, not an authority for unverified ROM2 value behavior.
REG-REC-032's ROM1 record layout is retained; its lookup clause is partially
retracted and cannot be imported as an unconditional ROM2 rule.
— R2-ASSET-006, REG-FMT-031, REG-REC-032

<a id="not-yet-surveyed"></a>

## Unknowns

General ROM2 record kinds, value interpretation, key-set differences, native
lookup, sort and application/writer acceptance are unspecified. Recognition by a
nested `&YA1` payload is broader than a `.reg` extension filter.

## Save-tail keys on the mission branch

The int getter returns record +4 for value type 1, where type is
(flags & 0xe) >> 1, and the default for a missing key. The int-array
getter returns +8/4 DWORDs at pool offset +4 for type 3, one element for
type 1 and none for type 0. Lookup branches on flag bit 0x10 and compares
names over 15 characters. Mission setup requests CurrentState/InBattle,
Character/Name, GameOptions Wimpy, ShowHP, FlyingHP, Formation, Speed and
ShowTimeFlow, View X and Y, SpellBook IsOpen and Pressed, Objects Selection
and Group0..9, Inventory IsOpen, Projectiles FreeIndex and IDs with one
root directory Prj%d per ID, then Fog FirstState and Data when InBattle is
nonzero. SpellBook/Shortcuts is stored but not requested there.
— R2-ENGINE-614

In the five mission-like saves the stores hold 28 records (45 in
game0006). InBattle is 1, View 8/29 and Selection [1] in all five.
— R2-SESSION-244

## Selected movie companion consumer

The selected native movie-open path replaces the SMK suffix with `.reg` and
loads an inline store. Common requests startx, starty, nFadings and
nPanaramings, with zero defaults. Fading1..N requests startframe/endframe
and floating startfade/endfade; Panaraming1..N requests startframe/endframe
and stepx/stepy. Preserve the native spelling Panaraming. — R2-ENGINE-077

The first selected teleport companion in both installs has 14 records and
a zero-byte pool, totaling 476 bytes under the envelope above. The native
lower open path then supplies the SMK resource to its named SmackOpen import.
General lookup/type semantics, malformed companion acceptance, full frame-plan
application, rendering and audio behavior remain Unknown. — R2-ENGINE-077

## Selected numeric branches and frame plan

The movie's selected getters mask kind with 0xe and shift right one.
Result 1 loads integer DWORD +4; result 2 loads double QWORD +4/+8.
Missing records return the caller's default. These branches do not establish
a general ROM2 lookup/comparator or acceptance grammar. The movie stores
double fade endpoints as float values in 16-byte segment records.
— R2-ENGINE-097

The measured first companion has Common/startx=0, starty=0 and nFadings=2;
nPanaramings is absent. Fading1 stores start/end 0/5 and doubles 0/1;
Fading2 stores 161/166 and 1/0. There are no Panaraming records in this
companion. Fading triggers at exact descriptor start equality and advances
the plan index; Panaraming sets/clears stored steps at start/end equality
and updates destination coordinates after the next-frame call. Missed
equality, visual interpolation, malformed lookup and timing remain Unknown.
— R2-ENGINE-097

## Effect type and sound registries

`graphics.res` `projectiles/projectiles.reg` (9584 bytes) and `sfx.res`
`sfx.reg` (9864 bytes) are byte-equal in EN and RU. The effect registry
holds `Global/Count` = 31 and sections `Projectile0..30`; the native loader
reads `File`, `A16`, `ID` (default -1), `Phases` (-1), `RotationPhases`
(16), `Width` (64), `Height` (64), `Palette`, `Homing`, `Flip` and `SFX`
(0 each) and keys the type by `ID`. — R2-ASSET-101

The sound registry holds `Global/SfxCount` = 660 and 200 `Sfx<N>` string
entries; each names a `sfx.res` path without `.wav`. Spell sounds are
`Sfx(500+ID)` for the effect identifiers of R2-ENGINE-445; the magic
entries are Sfx90 and 30 slots from Sfx508 to Sfx566, and an absent slot
gives no sound at the play sites read; no census of other sound sites was
made (Medium), and no read site places Sfx90. — R2-ASSET-102, R2-ENGINE-446

## Restored projectile continuation

The selected EN mission reader restores each Prj%d into client +0x9ec.
Object fields are x/y/z +8/+0xc/+0x10, picture +0x24, dir +0x70,
phase +0x74, lastaction +0x78, action byte +0x8c, actiondir byte +0x8d,
target word +0x8e, destination +0x90/+0x94/+0x98, actionphase +0x9c,
remaining segments +0xa8 and actionspell +0xac. Its update returns zero
before movement when +0xa8 is zero. The read client loop removes zero-return
objects. game0006 Prj16 stores zero segments, so unchanged entry fields
request removal on its first call. With nonzero segments, action1 can
refresh the destination from a found client target; the default movement
arm divides remaining distance by remaining segments. Picture-specific
presentation arms do not establish server damage or a SpellTransport link.
Actual first update state, visible timing and RU parity remain Unknown.
— R2-ENGINE-661

A unit caster's release gives the cast object of identifiers 16, 24, 34,
46, 48 and 62 zero remaining segments +0xa8 at creation, so that object
ends at its first update; the held image of these spells is a unit overlay,
not a projectile. The client marks +0x130 and the overlay list +0x11c are no
REG keys of the EN restore; the 38 EN sites that add the mark list's offset
were censused, four of them in functions of an unread class. —
R2-ENGINE-691, R2-SESSION-274, R2-ENGINE-701

A unit caster's Teleport (identifier 54) release puts two objects into
client +0x9ec, each with picture 54 and remaining segments +0xa8 = 21: one
at the caster's position shifted up-left by its unit type's Center minus half
Selection size, and a copy at destination +0x90/+0x94 shifted the same way.
Neither moves; phase +0x74 runs 0 to 20 and the object ends at its 22nd
update. — R2-ENGINE-703

A cast picture is not part of the caster's sprite. A caster in state 8
shows the attack frames of its current direction through the draw arm of a
strike (state 3) and state 7: sprite (MovePhases + MoveBeginPhases) × 8 +
16 + direction × AttackPhases + the per-update attack frame, or × 5 + 9 for
a Flip type. Neither the cast packet's caster arm nor the state 8 update
stores the direction in its own body. The state 5 update eases the
direction, the client writer of +0x8d and whether the shown direction is
the server facing are Unknown, and so is whether an item cast or a creature
cast reaches the client through this arm. — R2-ENGINE-929, R2-ENGINE-704

The release of an identifier other than 54 creates one effect object at the
caster's +0x5c/+0x60 plus 8 × (ShootOffset point of the draw direction
minus Center): the sprite point, in pixels from the sprite's Center. The
meaning of the caster's +0x5c/+0x60 is Medium and the screen direction of
the draw direction is Unknown. A type
without ShootOffset uses the Center/Selection anchor of Teleport. The
release sets no height; the draw lifts by +0x6c and +0x10, whose values are
Unknown. — R2-ENGINE-930, R2-ENGINE-703

The lift +0x6c is the bilinear sample, at the object's point, of a signed-byte
grid that the client's dword +0x80 supplies, the ground height (Medium); its
bytes at a given cell are Unknown. The teleport objects' slot +0x40 is the
effect light method, whose default arm stores no light for them; other
light writers are not excluded (Medium). —
R2-ENGINE-993, R2-ENGINE-703

For a type 23 or 24 caster a Fire Arrow (identifier 10) is created at state
8 counter 8, on attack frame 4 (Medium, through Parent), with remaining
segments distance / 200; a distance below about 25 px gives no picture. Each update moves it by the remaining distance over the
remaining segments toward its homing target and shows frame (counter / 2)
mod 4, from frame 0 one step from the start. A first draw in the release
tick and the seconds are Medium. — R2-ENGINE-931

An area cell picture (odd identifier, such as 23) is created by the packet
handler itself at the cell centre, height 0, with no caster animation; its
first update shows frame 0, in the next client tick (Medium). —
R2-ENGINE-932

The render draws per cell, rows top-down and columns right to left, the
structure and then the unit of grid +0x8c; then a sweep of grid +0x90 by a
call whose drawing is Unknown; then every effect in one pass; then the units
of grid +0x90. The order is High for these listed loops and Medium for the
absence of another pass in the unlisted render code. An effect covers every
grid +0x8c unit in the units' +0x28 draws;
whether a missile covers its caster depends on the caster's grid, which is
Unknown. The order among effects is Medium. — R2-ENGINE-933, R2-ENGINE-703

## Mission View and Selection consumers

The EN mission entry stores View/X and View/Y to the client origin
inputs. Constructor, resize, panel layout and draw store to the same
fields and can change them whenever they run; their order against the
REG stores is unread. The last writer at presentation and any followed
unit remain Unknown.
— R2-ENGINE-667

Objects/Selection word lookups call slot +0x14(1) of the object found
at client +0x9d0, whose class is not read; two acquired EN unit classes
bind that slot to L2.03528, which calls L2.03529 to store +0x80=1 and
+0x114=1. The entry opens or closes the book and inventory from the REG
SpellBook and Inventory IsOpen values through controls that are separate
from the Selection restore. Delivered client identity, membership and first-frame
state remain Medium/Unknown. — R2-ENGINE-670, R2-ENGINE-614

## Unit type attack timing keys

`graphics.res` `units/units.reg` is one file in EN and RU. Each unit type
section carries `ID` (61 of the 63 sections of the installed file; `Global`
and `Files` carry none); 16 of the 61 carry `Parent`. The loader reads `Parent`
with a default of -1 and takes a missing scalar from the parent's loaded type
only when `Parent` is not -1, else -1, except `ShootDelay` and `AttackDelay`,
which default to 0. `AttackPhases` (type +0x1c) of 0 means
the type shows no cast; the attack frame list built from `AttackAnimTime`
and `AttackAnimFrame` sets the cast animation length +0x68; `AttackDelay`
(+0x120) is the counter of the cast sound and `ShootDelay` (+0x11c) the
release counter for every cast identifier except Teleport's 54, which
releases at counter 0. `CenterX`/`CenterY` (+0x34/+0x38) and
`SelectionX1`, `SelectionY1`, `SelectionX2`, `SelectionY2` (+0x84..+0x90)
place a Teleport image. Mage types 23 and 24 carry AttackPhases 6,
AttackAnimTime 2 x 6, AttackDelay 8, ShootDelay 8, Center 64/78 and
Selection 48, 48, 80, 90. A type with a `Parent` and no `AttackAnimTime` of its own takes the list
and `AttackAnimFrame` from the parent's section; the case of a type with
neither a list nor a `Parent` is Unknown. — R2-ASSET-180, R2-ENGINE-704

`ShootOffset` (type +0x108) is 8 (x, y) sprite points, one per draw
direction, in 25 of the 61 types; types 23 and 24 carry their own, within
28 px across and 45 px up of Center (64, 78; Center through Parent).
`MovePhases` (+0x14),
`MoveBeginPhases` (+0x18) and `Flip` (+0x128) place the attack frames in the
unit sheet. — R2-ASSET-222, R2-ENGINE-929

## Dialogue picture keys and bounded crop search

The read EN/RU `scenario.res:npc.reg` contains one Multiplayer directory
and the four arrays FacesMF, FacesMM, FacesFF and FacesFM; there is no
individual speaker section in that payload. The read `units/units.reg`
contains 63 directories (61 Unit sections, Global and Files) and 60
InfoPicture leaves. No key name in those two payloads contains rect, crop
or portrait, case-insensitively. This does not exclude another field name
or reader route supplying portrait coordinates: that conclusion is Medium.
All 185 read infowindow BMP headers per root measure 160 by 240, 24 bpp.
— R2-ENGINE-820

The selected dialogue builder uses the type's InfoPicture string and the
client unit's variant to form its bitmap path. The source-box operands in
the selected route come from constants and a program table indexed by type;
full registry lookup, variant conversion and final rendering remain open.
— R2-ENGINE-818, R2-ENGINE-819
