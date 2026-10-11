# ROM2 generator presentation

These contracts describe selected native drawing paths. Coordinates are
relative to the 640x480 page unless a rectangle origin is named. Source
RGB values precede device packing. Live timing and complete pixel equality
remain unobserved.

## Pre-create hover help

The getter returns fixed main.txt lines. It does not compose the help
text. Its page admission word must be nonzero. Difficulty help also
requires campaign mode 2. Name help comes from the name child's getter
under the same page admission word. Line numbers are zero-based
(R2-ENGINE-384, Medium).

| Control | Mask | main.txt line |
|---|---|---|
| Difficulty 0, 1, 2 | 20, 40, 60 | 247, 248, 249 |
| Hero 0, 1, 2, 3 | 80, 100, 120, 140 | 250, 252, 253, 251 |
| Cancel | 160 | 255 |
| OK | 180 | 254 |
| Name | child | 256 |

The detail skill getter returns main.txt 171+i for a fighter and 176+i
for a mage, i=0..3. It rejects a missing active page or a failed mask hit.
The fifth skill is outside this getter's successful hit population
(R2-ENGINE-384, R2-ENGINE-288).

## Hover box

The selected EN shared receiver draws font2 with foreground shades
floor((160,120,50)*i/15), i=0..15, and a constant (8,8,8) shadow at
(x+1,y+1). Its fill RGB is (36,44,39). Bright border RGB is
(160,120,50), dark border RGB is (80,60,24). Separate 4x4 corner calls
have an unclosed resource binding (R2-ENGINE-385, Medium).

The help parser splits at byte `#`, retaining a line for each separator
and the final remainder. Consecutive separators account for blank lines.
There is no width budget or word-wrap call in these selected bodies.
With cursor (x,y), maximum measured segment width W and N lines, the
nominal bounds are (x,y-5-14N,x+W+11,y). Text starts at
(left+5,top+4+14i). A right overflow shifts the box left; a top overflow
shifts it down. The selected builder has no left or bottom clamp
(R2-ENGINE-385).

The receiver tries to open help only on the accumulator's crossing from
below 500 to at least 500 and below 25500, under its receiver and
application admission gates. Movement clears the accumulator, restores
the saved background and clears visible help. At 25500 it also restores
the background. The selected timer import was not resolved here;
R2-ENGINE-239 supplies the earlier millisecond interpretation. Button
events, the RU shared receiver and corner endpoint semantics remain
Unknown (R2-ENGINE-385).

## Name entry

The campaign name field is (300,433)-(464,449). The EN painter uses
font4, graphics/font4/font4.16a with its .dat advances and spacing 2.
It draws at (300,449-h), where h is the first glyph's height. The prompt
uses main.txt 365, font4 and gold source palette endpoint (185,159,73),
right-aligned ten pixels left of the field at its top. The typed name
uses source palette endpoint (255,255,116), without the shadow wrapper
(R2-ENGINE-386, Medium).

An admitted field appends `|` while its caret bit is set. Paint toggles
that bit when elapsed ticks exceed 500. Character input resets the bit
and timestamp. At length 10 the character handler appends nothing.
EN admits converted bytes at least 32; RU additionally rejects `~`.
The converted character's code-page mapping and the first glyph's exact
height remain Unknown (R2-ENGINE-386).

## Attribute values and command captions

Attribute row i uses font4 at (84,58+32i), left aligned, with gold source
palette endpoint (185,159,73) and black shadow offset (1,1). The pool
uses the same font, palette and shadow, anchored at (84,194). Its mode
centres by measured text width and half the first glyph width; that
vertical adjustment is the native operand, not an assumed glyph height
(R2-ENGINE-387, Medium).

The generator's three captions are main.txt 238, 239 and 260. They are
drawn separately from the Inn button images, with font4 and the same
centering mode. Normal source palette endpoint is (185,159,73); hovered
is (255,255,116). The pressed caption anchor moves right and down by one
pixel. These direct caption calls do not use the shadow wrapper
(R2-ENGINE-387).

## Statistics sheet

Let L,T,R be the rectangle supplied to the shared EN sheet painter.
Every text uses font2. Mode 0 is left aligned, mode 1 subtracts the full
measured width and mode 2 subtracts half of it. Name and ordinary labels
use shades of (185,159,73); values use (107,154,120); section headings
use (150,90,0), each floor(component*i/15), i=0..15. Every text has the
(8,8,8) shadow at offset (1,1) (R2-ENGINE-388, Medium).

| Block | Label or heading anchor | Value anchor | Rows relative to T |
|---|---|---|---|
| Name, optional second heading | R-88, centred | — | 18, 28 |
| Four attributes | L+6, left | L+80, right | 44, 54, 64, 74 |
| Health, mana | L+115, centred | L+115, centred | labels 44,64; values 54,74 |
| Damage, attack | L+6, left | L+70, right | 88, 98 |
| Absorption, defence | L+74, left | L+140, right | 88, 98 |
| Skills, resistance headings | L+6, L+74, left | — | 112 |
| Five skill, resistance rows | L+6, L+74, left | L+70, L+140, right | 122+10i, i=0..4 |
| Weight, experience | L+16, left | L+127, right | 174, 186 |
| Sight, speed | L+40, left | L+106, right | 200, 210 |

Attribute labels are main.txt 15..18; health/mana 19/20; sight/speed
21/22; damage/absorption/attack/defence 23..26; section headings 27/28;
skill rows 30..34 or 36..40; resistances 41..45; weight 35; experience
46. The painter selects the skill group and visibility grade from its
display record and flags (R2-ENGINE-388).

For the detail child bounds in R2-ENGINE-286, the candidate screen
translation is L=0,T=238,R=160. The specific caller-to-rectangle join and
visibility grade remain Medium. The shared painter and font drawer pass
the complete strings with no local shortening, ellipsis, wrapping or
per-row clip rectangle. Downstream glyph clipping, long-name overlap and
the RU sheet binding remain Unknown (R2-ENGINE-388).

## Guidance tip bounds

With TipsMode enabled, pre-create constructs its guidance child at
(232,48)-(640,184), extent 408x136, for tips8. Detail constructs it at
(0,280)-(312,480), extent 312x200, for tips5 or tips6. These are literal
parent-relative operands in both selected activation bodies. Hero,
difficulty and skill events retext the existing child; these bounds
are independent of cursor and measured text (R2-ENGINE-389, High).
