# World and character-panel appearance

[Reference](format.md)

## World appearance (`HERO-APPEAR-040`…`HERO-APPEAR-046`)

The class discriminator above decides *stats*. What decides the **drawn unit class** is a separate
chain, it lives on the client, and it does not read the stat discriminator at all.

What the server sends is not a drawn class. `R0059` pushes, under field-mask bit `0x4000`,
`R0880(actor)` — nine instructions returning `word[actor+0xe]`, the `typeID` — and then the
face byte `actor+0x4b`. The client parks them in `[L03223]` / `[L04218]` and assigns the
first to `drawable+0x20`, which is the subscript the frame selector `R0553` uses on the
`units.reg` class array `L02113`.

`R0590` then splits on that id:

| id | meaning | what happens to `+0x20` |
|---|---|---|
| `< 0x1a`, `>= 0x40` | a class a map places | left alone — it **is** the `units.reg` `ID` |
| `[0x20, 0x40)` | a player's character | discarded; see below |

The shipped roster occupies `1..27` and `64..80`, so `[0x20,0x40)` holds no class and is free for the
four archetypes `0x21..0x24`. For them the routine computes `typeID - 0x21`, files bit 0 into
`drawable+0x18c` bit 2 and bit 1 into bit 1 — the **same** assignment the chargen record uses, so
bit 2 is the sex axis and bit 1 the mage axis — sets bits 0 and 3, and stores the literal `1`.

`R0551`, the next instruction in the caller, does the real work. It returns at once unless
`+0x18c & 1`. Otherwise it builds a **body name** from a twelve-slot array of visible equipment at
`drawable+0x15c`, which the equipment messages fill from twelve `u16` at `msg+0xc`:

```
name = registryString[(byte[slot0 + 6] & 0x1f) - 1]        slot 0 absent -> index 0
if slot1 != 0            name += "_"
if mage and name == "unarmed"   name = (action == 6 ? "mage_st" : "mage")
```

and then maps the name to the drawn class with a seventeen-arm chain:

| name | id | name | id | name | id |
|---|---|---|---|---|---|
| `unarmed` | 1 | `axeman_` | 8 | `pikeman_` | 13 |
| `unarmed_` | 2 | `axeman2h` | 9 | `archer` | 14 |
| `swordsman` | 3 | `clubman` | 10 | `bowman` | 14 |
| `swordsman_` | 4 | `clubman_` | 11 | `xbowman` | 15 |
| `swordsman2h` | 5 | `pikeman` | 12 | `mage_st` | 24 |
| `axeman` | 7 | | | `mage` | 23 |

Seventeen names, sixteen distinct ids -- `archer` and `bowman` are the only pair that share one. `bowman` ships no sheet on either root and is dead.

**The pixels do not come from `units.reg`'s `File`.** The same routine composes

```
graphics\units\ <dir> \ <name> \sprites.256      -> drawable+0x194
graphics\units\ <dir> \ <name> \spritesb.256     -> drawable+0x198
```

where `<dir>` is `heroes` for a mage, `heroes_l` for a fighter with nothing in slot 7, and otherwise
`material.reg`'s `Material[word[slot7 + 6] >> 12].Path` — sixteen blocks carrying `heroes` for
0–7 and 14 and `heroes_l` for 8–13 and 15. So the class record supplies only the **geometry**:
phases, `Width`/`Height`, `CenterX`/`CenterY`, `Flip` and the three timelines.

`+0x19c` caches the last name and `+0x1ac` the last material index; nothing is reloaded unless one of
them changed.

**It is recomputed while the mission runs**, on the ordinary field-masked state message (opcodes
`0x6c`/`0x6e`/`0x6f`/`0x70`) as well as on the two equipment messages (`0x9c`, `0x76`) and from the
character screen. The server picks between the two equipment opcodes on the same `[0x21,0x40)` test.

**A consumer must therefore not store a drawn class on a player's character.** It must keep the
twelve visible-equipment slots, recompute the name whenever they change, load the sheet by path, and
take the geometry from the resulting `units.reg` record. For a unit a map places it must do none of
this: that actor's drawn class is its `typeID` and nothing derives it.

**Which `typeID` a placed human holds.** The `typeID` of the Humans row the placement resolves to,
not the placement record's own class key (`ANIM-116`). No store reachable from the client unit dispatcher selects it again, because the
derivation above returns unless drawable `+0x18c` bit 0 is set, and creation leaves that bit clear
for every class below `0x1a` (`ANIM-117`, amended). A unit's bit is set only by the hero arm for an id in `[0x20,0x40)`, reached through a nonzero Humans constructor mode; of the 464 placements of the shipped maps 4 have such a mode, and the primary-hero and `AddHero` modes were not traced (`UNIT-140`). The swing sound is `Sound[0]` of the class held at the
call (`ANIM-105`). A placed mage whose record names class 23 (`Unarmed Mage`, silent) and whose row
names class 24 (`Human Mage`) holds 24 and the swing hook is passed slot 510 (`ANIM-118`).

Claims: `HERO-APPEAR-040`…`HERO-APPEAR-046`, `UNIT-APPEAR-030`, `ANIM-116`…`ANIM-118`. The name list's order and the
second sheet's consumer follow the figure-equipment rules below.

## Figure equipment slots (`HERO-APPEAR-047`…`HERO-APPEAR-055`, `HERO-FIGURE-062`)

**Who sends them.** `R0669` sends nothing at all unless `actor->vt+0x30()` — the humanoid
predicate — is true, so a non-humanoid actor's twelve slots stay null for its whole life. With no
named recipient it walks the player list and calls itself once per player.

**Which field becomes which slot.** Both opcodes visit `actor+0x74`, `actor+0x78`, then
`actor+0x198+4i` for `i = 3..12` — `ITEM-EQUIP-006`'s own 1..12 equipment numbering (its wrapper
and recompute clauses are retracted; the slot map stands). Wire slot `k` is equipment slot `k+1`,
and an empty slot is sent as 0 rather than omitted.

**The two opcodes do not carry the same bytes.**

```
0x9c  twelve bare u16 at msg+0xc+2k, each = word[item+0x40]
0x76  twelve serialised item records from msg+0x13, each 7 + len bytes:
        rec+0  u16   word[item+0x40]        -> slot+0x06
        rec+2  u16   item+0x42, the count   -> slot+0x10
        rec+4  u8    bit7 = item+0x14 != 0  -> slot+0x08
                     bit6 = item+0x8  != 0
        rec+5  u8                           -> slot+0x09
        rec+6  u8    a length               -> slot+0x0a
        rec+7  ...   len bytes              -> malloc'd at slot+0x0c
      rec+0 == 0 means the slot is empty.  An empty slot is sent as a
      throwaway Item constructed, serialised and destroyed.
```

**The `u16`.** Every bit is consumed, by `R0886`, which turns it into a seven-digit name:

```
  bits 15..12  A = v >> 12          the material.reg index  (the world-appearance slot-7 axis)
  bits 11..8   B = (v >> 8) & 0xf   the equipment slot, 1..12
  bits  7..5   C = (v >> 5) & 7     unnamed
  bits  4..0   D = v & 0x1f         the item's Data.bin definition row index

  name = sprintf("%02d%02d%1d%02d", A, B, C, D)
  or     sprintf("%02d%02d%03d",    A, B, v & 0xff)   when B == 14
```

`R0667` assembles it from `item+0x46` (A), the item's kind (B), `item+0x45` (C) and
`item+0x0c` (D), each masked to a byte and `OR`ed — so a definition row index above 31 would corrupt
C.

**What the name addresses.** Two trees, and they are different pictures:

```
graphics\equipment\ <figure> \ primary   \ <name>.256    the info window's figure, per slot
graphics\equipment\ <figure> \ secondary \ <name>.256    slots 3, 8, 9, and 7 for a mage
graphics\equipment\ <figure> \ <face>.256                 the head, from drawable+0x24
graphics\inventory\ <name>.16a                            the icon; missing -> "Invalid item weared ",
                                                          posted only under -trace

<figure> is one of mfighter, mmage, ffighter, fmage, chosen by (drawable+0x18c & 6) —
the mage bit and the SEX bit together.
```

`R0745`, virtual slot `+0x80` of the drawable class, loops all twelve slots and builds these.
**It is not the world renderer.** Every one of the 928 shipped equipment sheets holds exactly one
frame of 160×240, against a hero body sheet's 129–216 frames of 24×40 to 40×48, and the routine's
identified callers are the info-window routines. In the world nothing is composited over the body:
a held weapon is visible because the world body name changes.

**The name list is a shipped file.** `main\text\heropicture.txt`, loaded into the object at
`L04227` at startup, 26 entries (the blank line is entry 22), identical on both roots, indexed by `D - 1` (`HERO-APPEAR-052`, `ANIM-108`):

```
 0 unarmed      5 swordsman2h  10 axeman      15 pikeman   20 archer
 1 swordsman    6 clubman      11 axeman2h    16 pikeman   21 xbowman
 2 swordsman    7 clubman      12 mage_st     17 axeman    22 (blank)
 3 swordsman    8 clubman      13 mage_st     18 axeman2h  23 Sonic Beam
 4 swordsman    9 clubman      14 pikeman     19 archer    24 Flame Thrower
                                                              25 swordsman
```

`bowman` never occurs, which is why the seventeenth world-appearance arm is dead. Three entries name no shipped body
directory. The list is 26 long while D is five bits wide.

**The second sheet is drawn.** `R0553` blits `drawable+0x194` and then `drawable+0x198` at
the same frame index behind one `+0x18c & 1` gate, the second centred from its own width and height
and the class record's `+0x30`/`+0x38`. Its six arguments are `dstX`, `dstY`, frame,
`[L04335]`, sun shear, mirror — `TERR-SPR-038`'s push list, so the fourth is
`TERR-SPR-066`'s **shroud level** and the second sheet belongs to the silhouette family.
**`0x26` is not one of them** (figure composition below): `HERO-FIGURE-062` corrects the partially retracted
`HERO-APPEAR-056` by identifying it as the preceding gate's key, not a blit argument.

**A consumer must therefore** keep twelve slots per humanoid actor, fill them from the equipment
slots offset by one, decode each `u16` into the four fields above, use D of slot 0 against
`heropicture.txt` for the body name, A of slot 7 against `material.reg` for the directory, and treat
`graphics\equipment` as a portrait tree that never touches the world sprite.

Claims: `HERO-APPEAR-047`…`HERO-APPEAR-055`, `HERO-FIGURE-062`, `ITEM-APPEAR-023`, `ITEM-APPEAR-024`,
`UNIT-APPEAR-031`, `SPR256-EQUIP-042`.
Not established here: field C beyond the figure-composition clause below; which garment each of slots 2..6 and 8..11 is.

## Figure composition and hit map (`HERO-FIGURE-057`…`HERO-FIGURE-064`)

`R0745` builds primary/secondary equipment sheets and a face sheet. It
fills a colour surface and, when supplied, a separate item-picking surface,
binding them in turn with `vt+0x28`. The two passes have different layer sets:

```
  pass 1   vt+0x18 = R0794   ->  R0895 / R0796
           (x, y, frame, palRow, mode)      palette = layer+0x1c + (palRow << 9)
           the picture

  pass 2   vt+0x40 = R0896   ->  R0897
           (x, y, frame, tag)               tag is the SIXTH argument of the blitter
           a byte-per-pixel stencil: for an opaque run R0897 copies no source
           pixel at all, it stores the tag  (L04343 / L04344)
```

**The tag is the equipment slot number**, `drawable index + 1`, over all 28 draw sites; the head
layer takes 0. So the second surface is a map from screen position to equipment slot, which is what
an info window with clickable equipment needs, and it is a second, independent derivation of the
`k+1` slot map in the wire-equipment rules above.

**Draw order is program order and depends on the mage bit** (`L04346`):

```
  mage       p7 head p11 p9 s9 p3 s3 p6 p4 p8 p0 p5 s7
  non-mage   p11 p10 p6 p3 p4 p8 p9 p0 p7 s3 p5 s8 s9   then  p1 or p0
```

Three asymmetries a consumer must reproduce rather than smooth over: index **2** (equipment slot 3)
is drawn by neither half and ships no sheet; index **10** (slot 11) is drawn only by the non-mage
half and ships no sheet; index **8** (slot 9) is **tagged but not colour-blitted for a mage**.

**Which held layer is on top.** `R0898` takes slot 0's field D, indexes `heropicture.txt`,
and returns 1 for `bowman`, `archer`, `xbowman`, `axeman2h`, `swordsman2h`, `mage_st` — every
two-handed and ranged body. Predicate 1 → index **0** last; predicate 0 → index **1** last. Field D
is the only field of the `u16` this routine reads.

### Non-equipment hero background

The hero fighter also draws `graphics\interface\heroback\backm.256` or
`backf.256`. Loader `L04549` creates these `.256` objects in globals
`L04550` and `L04552`. The compositor admits them when
`(drawable+0x18c & 3) == 1`: hero bit0 set, mage bit1 clear. Sex bit2 selects
backf when set and backm when clear. Ordinary-human bit4 does not admit this
layer, even with the same sex, face and worn items. This gate is not a check
for the primary protagonist. — HERO-FIGURE-144, HERO-APPEAR-041, UNIT-PICT-035

Draw order is optional `graphics\infowindow\horse.bmp`, hero background,
fighter face, then equipment. The background draws at `(0,0)`, frame0,
palette row 0, without mirroring. Setup uses its own embedded BGR0 palette
with one packed-16-bit row and zero global colour additions; framebuffer
channel packing remains an input. It does not use an owner-shade table.
The selected EN/RU backm/backf payloads are identical: one 160×240 frame,
3547 literal pixels and inclusive bounds `(31,123)..(115,199)`.
— HERO-FIGURE-144

This background contributes colour only. It is absent from the separate
item-mask pass, so pixels covered only by it retain tag 0; overlapping
equipment can supply its own tag. Neither background is drawn after
equipment in this compositor. Native captured colours, untraced alias
writes and decoration outside the compositor remain Unknown.
— HERO-FIGURE-144

**The list accessor has no bound.** `R0668` is `[[L04369] + ([list+0xc] + i)*4]`, six
instructions, no test — and this caller runs only when slot 0 is **occupied**, so `D = 0` indexes at
**−1**. A consumer must clamp at both ends; `D ≥ 27` runs off the other, into the first line of `stats.txt` (`HERO-FIGURE-063`, `ANIM-108`).

**And the same drawable field feeds sound.** `R0577` picks one of eight 48-byte voice banks
on the mage bit, the composed-figure bit, **whether slot 0 is occupied**, and the sex bit:

```
  mage bit set                          -> L02750   m_mage    / f_mage
  composed-figure bit set, mage clear    -> L02746   mf_hero   / ff_hero
  neither, slot 0 occupied               -> L02748   mf_merc   / ff_merc
  neither, slot 0 empty                  -> L02752   m_peasant / f_peasant
```

`R0585` fills them at start-up from `sfx\<bank>\` + `select1/2`, `command1/2/3`
(`HERO-FIGURE-061`, amended: eleven leaves per bank, `ANIM-094`), `retreat`, `defend`, `idle`,
`easy`, `hard`, `die`. The slot-0 condition is reached only when the mage bit and the
composed-figure bit `0x1` are both clear, so it can select different voice banks for armed and
unarmed humanoids without either bit, never for a drawable with bit `0x1` (`ANIM-096`).

**The hurt cue reads three bank fields.** The hurt hook `vt+0x68` takes an event index `k`. While
`+0x18c & 0x11` is non-zero, `k` alone picks the bank field; otherwise `k` reads the drawn class's
`Sound[k+1]` (`ANIM-094`):

```
field   file       k   event (ANIM-095)
+0x24   easy.wav   1   0x73 blow, health before it at least half and above -10
+0x28   hard.wav   2   0x73 blow, health before it below half and above -10;
                       state sync while the corpse stage stays 1
+0x2c   die.wav    3   state sync that moves the corpse stage from 0 to 1
half = maximum health unit+0x100 / 2, truncated toward 0
```

`k = 0`, a `0x73` message that leaves health unchanged, always plays the drawn class's `Sound[1]`;
for a drawable with bit `0x1` that class follows equipment slots 0 and 1 and the mage bit. A blow
that finds health at -10 or below plays nothing. `k = 1` and `k = 2` stay silent within 1500 ms of
the drawable's voice timestamp `+0x190`, which the select, command, retreat, defend and idle voices
also set; `k = 3` neither tests nor sets it. With bit `0x1` set, the sex bit and the mage bit
choose the bank, and equipment and the drawn class choose neither the bank nor the field of `k` 1
to 3 (`ANIM-096`). A `0x73` that leaves health unchanged comes only from a zero-damage unit strike on an actor
(`ANIM-125`); a structure target gets tag `0x82`.

**The sex bit of a person below class `0x1a`.** The setter's parameter is server `actor+0x4b`;
its bit `0x80` becomes drawable bit `0x4`. A map-placed person takes the byte from the spawner
(secondary key low byte, bit 7 from placement flag bit 2) when the placement's secondary key
word is non-zero, and from the Humans-row constructor otherwise (row slot 17, with bit 7 from row
slot 18). A tavern hire takes the constructor value of the row named `NPC%02d_%d`. A summon
builds a Units actor and never takes this branch (`ANIM-127`).

**Five readers play the other bank fields.** `CUnit` and `CAirUnit` `vt+0x6c`, `+0x70`, `+0x74`,
`+0x78` and `+0x7c` read the fields `+0x04`..`+0x20`, each behind the unit's voice stamp `+0x190`
(`ANIM-119`):

```
vt     ms    field read                                      reached by (VIDEO-067)
+0x6c  3000  command1, command2, command3 or defend, by       move, attack, swarm, patrol, town
             r = rand() >> 13 (defend only when r = 3 and
             the bank is not a peasant bank)
+0x70  3000  defend    +0x1c                                  guard, stand ground, defend (VIDEO-070)
+0x74  3000  retreat   +0x18                                  retreat (VIDEO-070)
+0x78  2000  select1 or select2, by r = rand() >> 14          selection (VIDEO-069)
+0x7c  3000  idle      +0x20                                  pickup (VIDEO-070)
```

A reply stores the stamp before it picks a field, so a null field still starts the cooldown. No
other state is stored: the pick is a function of bank, draw, clock and stamp (`ANIM-119`). The
stamp starts at 0, so the first reply is admitted once the clock reaches the threshold (`ANIM-120`).

The speaker is one member of the selection, chosen by `R2089` before the reader runs: a
hero-shaped drawable if any is selected (`+0x18c` bit 0x1), else an armed human (bit 0x10 and slot 0
occupied), else an unarmed human, at random inside that tier, among members with `+0x7c` set and
health above 0. The chooser reads no owner, position or visibility field and picks one member
only, so a speaker on cooldown leaves the reply silent (`VIDEO-068`). The cast orders call no
reader (`VIDEO-067`).

Claims: `HERO-FIGURE-057`…`HERO-FIGURE-064`, `UNIT-FIGURE-032`, `ITEM-APPEAR-025`, `ANIM-119`, `ANIM-120`,
`VIDEO-067`…`VIDEO-070`.
Not established: what the effect-list keys `0x26`/`0x30` name; which garment
each of indices 2..6 and 8..11 is.
