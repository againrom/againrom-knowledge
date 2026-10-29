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
