# ROM2 character generator

A fresh ROM2 campaign creates its hero on two screens: pre-create
(character, difficulty, name) and detail (attributes and the chosen
skill). The scope is selected native paths in the preserved EN/RU clients
and the stored Humans templates; the original game was not run.

## Route

Main-menu New Game posts 0x425; its handler enters campaign mode 2 and
shows pre-create with `music\chrgen.wav`. OK (or Enter) with a non-empty
name shows detail; Cancel (or Esc) returns to the main menu. Detail
Accept (or Enter) commits the hero, sets scenario slots 776 and 781 and
the session difficulty, and posts 0x42e, which shows town 1. Detail Back
returns to pre-create. Only pre-create handles Esc. Every compared EN/RU
body pair is equal or differs in numeric operands (RU application fields
+0x18), except pre-create OK: its mode-2 arm, the New Game path, is the
same, while the RU arm for other modes also checks the name's length and
leading space. Six bodies on the route have no located RU entry, so their
RU equality is not measured: New Game, the campaign-mode switch, the
application command with the Accept arm, the detail show, the commit and
the cursor loader. The RU SetVar 776/781 pair sits at `L2.01108` and
`L2.01109`, EN at `L2.01106` and `L2.01107`. — R2-ENGINE-283

## Pre-create screen

| Control | Rectangle | Mask | Art |
|---|---|---|---|
| Difficulty 0..2 | (8,0), (296,0), (580,0), 48x72 | 20, 40, 60 | `level{n}on` selected, `level{n}l` hover, `level{n}lon` both |
| Heroes 0..3 | (116,44), (180,44), (392,44), (456,44), 64x244 | 80, 100, 120, 140 | `h{n}on` hover; `h{n}sel` 268x340 selected at (112,44) or (260,44); pair art `h1sel2`, `h2sel1`, `h3sel4`, `h4sel3` |
| Cancel | (16,400), 64x76 | 160 | `cancell` while hovered or lit |
| OK | (548,400), 80x76 | 180 | `okl` while hovered or lit |
| Name | (300,433)-(464,449) | – | label main.txt 365 |

The background is `mainarea` with two 15-frame torches and a `blind`
sprite played over the levels, Cancel and OK. — R2-ENGINE-284

Hero indexes 0..3 are male fighter, female fighter, female mage and male
mage (Start_MF, Start_FF, Start_FM, Start_MM). The defaults are hero 0,
difficulty index 1 and npcnames line 20; a hero click replaces a default
name with npcnames 23, 24, 26 or 25. Difficulty changes no other control.
— R2-ENGINE-285

## Detail screen

| Child | Rectangle | Art |
|---|---|---|
| Stats | (0,0,160,238) | localized `graphics/chrgen/leftup.bmp`; plus and minus 20x20 |
| Stat sheet | (0,238,160,480) | `FullStatsL` and the shared stat-sheet painter |
| Skills | (160,0,480,480) | class `column`; skill `on`, `shine_off`, `shine_on` |
| Buttons | (480,0,640,238) | Inn `ButtonsArea` and `button{1,2,3}{on,off}`: Accept, Reset, Back |
| Inventory | (480,238,640,480) | the inventory panel |

Stats rows Body, Agility, Mind and Spirit have a value at (82,54+32i),
plus at (107,54+32i) and minus at (132,54+32i); the free pool sits at
(46,181). Activation shows the template with pool 0. EN Reset sets all
four to 25 with pool 100; RU Reset restores the template with pool 0.
— R2-ENGINE-286

Cost is F(v) = trunc(0.349·1.15^(v−1) + 0.5). Plus needs pool ≥
F(v+1)−F(v) and v < 45; minus needs v > 15 and refunds F(v)−F(v−1). The
producer accepts 140 − ΣF ≥ 0, else all four become 25. Every Start row
sums to 140. — R2-ENGINE-287

| Class | Selectable index 0..3 | Not selectable |
|---|---|---|
| Fighter | sword, axe, mace, pike | bow |
| Mage | fire, water, air, earth | astral |

The hit test compares four mask bytes only; the fifth skill has no click
and no art. The template main skill 1 (sword or fire) is lit until a
click. — R2-ENGINE-288

## Guidance

Under the tips mode, pre-create shows town.txt `#tips8`, then `#tips9`
after the first hero click and `#tips10` after the next difficulty click.
Detail shows `#tips5` (fighter) or `#tips6` (mage), then `#tips7` after
the first skill click. Two cycles light hero, level and button targets or
the four skills while the player hesitates. Tooltips come from main.txt
155..158, 171..180 and 273. Pre-create uses the select cursor and detail
the default cursor. — R2-ENGINE-289

## Hero at town 1

The producer loads the Start row, sets the chosen school skill to 20,
skill 5 to 10 and the others to 0, equips a weapon for the chosen skill
and recomputes derived values; HP and mana start at their maxima. A
fighter gets an iron weapon or the Uncommon Wood Long Bow. A mage gets a
staff casting the school's spell at 20: Wood Staff for the female mage
and Uncommon Wood Staff for the male mage. Hero word +0xe is 0x21 for MF,
0x22 FF, 0x23 MM and 0x24 FM; 0x22 and 0x24 are female for both the staff
and the attribute caps.

| Template | Experience | HP max | Mana max | Speed | Sight byte |
|---|---|---|---|---|---|
| Start_MF | 7320 | 142 | – | 19 | 6 |
| Start_FF | 7320 | 123 | – | 19 | 6 |
| Start_FM | 7320 | 29 | 157 | 16 | 6 |
| Start_MM | 7320 | 42 | 99 | 16 | 6 |

Values are for template attributes before item modifiers and load.
— R2-ENGINE-290

The four Start rows store the attributes, school skill 1 at 20 and their
items; the three data.bin payloads agree. — R2-ASSET-076. The generator's
art, masks and sounds are equal in EN and RU except the localized stats
art. — R2-ASSET-075

A town-1 save carries slots 776 and 781 at scenario bank offsets 0xC20
and 0xC34. Player gold 1000 at +0x3c is Medium: the server sets 1000 when
the field is 0, but the order of that call before the first save was not
traced. Which bytes hold the hero record remains Unknown. — R2-SESSION-131

## Unknown boundaries

Item modifiers and load on HP, mana and speed, a displayed level, the
stat-sheet and inventory content, the hero record in the SAV and live
timing remain Unknown.
