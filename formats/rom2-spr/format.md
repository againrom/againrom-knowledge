<a id="rom2-sprite--palette-containers--identity-survey"></a>

# ROM2 sprite and palette containers

The preserved `.16a`, `.256`, `.16` and `.pal` resources have the container
shapes below. This reference establishes frame geometry and residual regions;
the survey rows do not establish general ROM2 RLE, pixel, color or runtime
draw semantics. The named town-1 ShopFrame reader below is a separate,
bounded native contract.
— R2-ASSET-007, R2-ASSET-008, R2-ASSET-009

<a id="16a--256-result"></a><a id="16-result"></a>

## Sprite envelope

| Family | Prefix | Frame | Final trailer |
|---|---|---|---|
| `.16a` / `.256` | Optional 1024-byte palette selected by trailer bit 31 | u32 width, u32 height, u32 dataSize, dataSize bytes | Low 31 bits = count; bit 31 = palette |
| `.16` | No palette | Same three-u32 frame header and counted data | Plain u32 frame count |

Read the trailer, select the origin, then walk the counted records using
`12+dataSize`. Bound every frame against its containing resource. A zero-byte
archive entry is a stub, not a zero-frame instance of this envelope.
The same region order defines a structural emitter, but original ROM2 pixel
encoding and writer acceptance remain unspecified. — R2-ASSET-007, R2-ASSET-009

## Residual regions

Five `.256` entries contain bytes between their indexed frames and final
trailer. They are identical to the corresponding ROM1 RU resources with
appended secondary sections. `.16` font1 and font2 likewise retain 17716 and
32 residual bytes; their corresponding ROM1 RU resources have in-place
overwrite layers. Font3 has 64 frames and ends at its trailer; font1/font2
have 224 indexed frames. — R2-ASSET-007, R2-ASSET-009

The ROM1 residual interpretations are cross-referenced by SPR256-EXC-017,
SPR256-EXC-020, SPR16A-FONT-014 (partially retracted for ROM1 tail reachability)
and SPR16A-FONT-021. Their native ROM2
consumer behavior is not independently specified here.

<a id="pal-result"></a>

## Palette layouts

| Shape | Layout |
|---|---|
| BMP palette | 8-bpp BMP, `bfOffBits=0x436`; 1024-byte color table at byte 0x36 |
| Raw owner tables | 16 consecutive 1024-byte tables, total 16384 bytes |

The preserved resources use 156 BMP-shaped and three raw palettes. The
shape determines where to read/write the table bytes. The ordinary-unit
consumer below adds a bounded native interpretation; other palette
consumers and preparation modes 0, 1, 3 and 4 remain unspecified.
— R2-ASSET-008, R2-ASSET-226

## Ordinary mission unit palettes

The client draw reads the unit type's DWORD +0x98. Zero selects the shared
palette collection by the owner getter's DWORD +0x0c; one selects type
palette 0; other values select type palette `unit DWORD +0x28 - 1`
(High for the read branches; Medium that the record is the unit's type
and that unit +0x14 is the owner object).
The owner getter returns that field directly. The type-loader local
stores the `Palette` integer getter's return in +0x98. Unknown includes
the default pushed for Palette. The bounded producer paths below do not
close event ordering on START MISSION, Restart or LOAD; mission 42's
runtime values follow below. — R2-ENGINE-955,
R2-ENGINE-956

The read map-player path takes the first file DWORD as row +0x0c. Its
server builder stores row +0x0c + 1 at Player byte +0xa44; message 0x96
sends that byte minus one. A missing EN client owner receives the message
colour at +0x0c. For map owner ordinals 1..6 of mission 42 the conditional
client values are 1,1,2,3,4,5 in both maps. A separate read server local can
replace +0xa44. Complete writer/order coverage remains open (High for
positive local stores, Medium for the composition). — R2-ENGINE-963

The client copier stores server WORD +0x0e as DWORD +0x24 and passes server
BYTE +0x4b to the +0x28 setter. For types below 0x1a that setter removes
bit 0x80; types 0x20..0x3f become client type 1 with separate flags. The
EN map loop's binding of these inputs to successive template parameter slots
16 and 17 is Medium, with a conditional bit-0x80 addition from slot 18.
The retained WORD-reader prefix EN L2.02668 / RU L2.01281 computes index + 1
but ends before its writeback and return; its tail is unread. The parameter
getter and complete constructor/write population remain open. The EN
sender transmits type then variant under mask 0x20. One conditional
existing-unit client arm writes +0x28 directly, bypassing the creation
setter (High for the paired copier/setter stores, observed reader-prefix
instructions and EN transfer stores; Medium for numeric parameter binding,
including the unread WORD-reader tail and getter, and full composition).
— R2-ENGINE-964, R2-ENGINE-965

Mission 42's retained numeric map/template join identifies four type-24
person records owned by 4: records 1,2,86,87 at raw positions
14720/17792, 27776/32128, 30336/32128 and 27520/33920 (Medium for
identity). R2-ENGINE-966's composed clause that the Unit24 registry
Palette value 5 selects a type palette for them is partially retracted.
— R2-ENGINE-966, R2-ENGINE-955, R2-ENGINE-956

In the running EN original, after START MISSION, after Restart and after
LOAD of a save written in mission 42, these four client units hold type
24, variant 16 and an owner whose +0x0c is 3. Type record 24 holds
+0x98 = 0 and no type palette, so the selector takes shared collection
entry 3. The party takes entry 1 and Peasants entry 2. Dragon and
Monster types take type palettes: +0x98 = 5 selects index `+0x28 - 1`,
+0x98 = 1 selects palette 0. Wherever read, all 92 units keep these
values across the four reads; the scan after the dialogue read the fields
of all 92 units but two owners and one type record. High for the read values; Medium that the draw call
itself resolved them, since no breakpoint trapped it. In a frame drawn
through the cnc-ddraw wrapper, a head pixel of mage record 1 has word
0xec28, entry 3 row 3 index 109, which only table 3 produces (Medium for
the index and the hood). Not read: why the registry Palette = 5 does not
reach type record 24 (the directory bound to type 24); shared entries 0, 2
and 4..16; the hood's own sprite index; units created later in the
mission; the draw's own palette argument. The RU original was not run.
— R2-ENGINE-975, R2-ENGINE-977, R2-ENGINE-978

`graphics/units/humans/human.pal` has 16 tables of 256 B,G,R,X entries,
1024 bytes per table, and is byte-equal in EN and RU. The EN shared builder
passes that path and a 16384-byte read to unread resource helpers (Medium
that they read this entry) and prepares source tables 0..15 with
16 rows, mode 2 and channel additions enabled. Both builders append, in
the last iteration, a mode-5 palette from table 0 with (16,5,0).
Source-table i mapping to shared collection index i remains Medium
because initial emptiness and
the nested insert helper are unclosed. In the running EN original,
entries 1 and 3 equal prepared tables 1 and 3 in all 4096 row WORDs with
zero channel additions (High for those two entries). EN type palettes are
loaded from BMP colour-table offset 0x36 into 1024-byte source slots;
their full path composition remains Medium. — R2-ASSET-226, R2-ENGINE-956,
R2-ENGINE-957, R2-ENGINE-976

Mode 2 writes rows for L=16 down to 1. Each channel is
`clamp(trunc((source + current channel byte)*2*L/16),0,255)`, then packed
using the device's current channel widths and shifts. The normal byte
sprite wrapper addresses row `palette base + 512*shade`. Its literal
commands overwrite destination WORDs with that row's indexed colour;
skips preserve destination pixels. The read normal EN/RU blitters use
no destination-colour blend. With mark 32 the call goes to slot +0x34
when unread L2.02954 on view +0x9cc accepts the owner's +4, and the draw
leaves without a sprite call otherwise; the slot +0x34 blitters were not
read. The upstream light calculation, channel-byte writers, special
blitters, complete RU producer parity and final screen
colour remain Unknown. — R2-ENGINE-957, R2-ENGINE-958

## Town-1 ShopFrame reader

The EN shop loader builds `interface/ShopFrame.256` as a byte sprite with
one palette row, palette mode 1 and zero colour adjustment. Mode 1 reads
B, G and R from each four-byte palette entry and packs a WORD with the
active output channel widths and shifts. Paint slot +0x18 calls normal
reader `L2.00798` (`R2-ENGINE-270`, High).

Each command is one byte. The low six bits are a count. Top bits 00 copy
that many following palette indexes, 01 skip that many rows keeping the
column, and 10 or 11 skip that many pixels. A row ends when the column
reaches the width. The reader stops after the frame height, not at
dataSize (`R2-ENGINE-270`, High).

| Offset | Bytes | Content |
|---|---|---|
| 0 | 1024 | palette, 256 four-byte entries |
| 1024 | 12 | width 316, height 303, dataSize 11015 |
| 1036 | 11015 | 651 literal and 1743 pixel-skip commands; 8621 pixels, 87127 skipped cells |
| 12051 | 1743 | residual, starting with a copy of the trailer |
| 13794 | 4 | trailer `0x80000001` |

EN and RU resources are the same 13798 bytes. The selected reader ends at
12051 and never reads the residual. Only `graphics.res` holds .256 keys:
1929 per locale, eight empty. Of the 1921 non-empty entries, five
single-frame entries carry residual bytes between the frame and the
trailer, and in each the residual length equals the frame's
skip-command count. Their purpose remains Unknown (`R2-ASSET-070`,
High / Unknown).

<a id="not-yet-surveyed"></a>

## Unknowns

Consumers outside the named byte reader, the residual purpose, malformed
resources and device output are outside the established contract. A
residual consumer or the producing tool's contract would settle the tail.
RU reader bodies were not decoded separately (`R2-ENGINE-270`,
`R2-ASSET-070`).
Related ROM1 references: [SPR16A](../spr16a/format.md),
[SPR256](../spr256/format.md), [PAL](../pal/format.md).

## Spell effect sprites

`graphics.res` `projectiles/projectiles.reg` names 31 effect types by `ID`,
`File`, `Phases`, `RotationPhases`, `Width`, `Height`, `Palette`, `Homing`,
`Flip`, `A16` and `SFX`. An `A16` type loads `<File>.16a`; others load
`<File>.256` and `<File>b.256`. A type with RotationPhases 16 and Flip
stores Phases frames for each of nine directions; directions above 8 are
drawn mirrored. A type with RotationPhases 1 stores Phases frames. The
installed `healing` sprite has one frame more than its Phases; `diamonddust`,
`freeze`, `smallxpl`, `smoke0` and `smoke1` have no type. EN and RU sprites
and registry are byte-equal. — R2-ASSET-101, R2-ENGINE-447, R2-ENGINE-449

A spell row selects type 2*row+8 when cast; Prismatic chains use type 30.
An area cell uses 2*row+9: that value is read for row 7 only, and for the
other rows, including the bindings of identifiers 13, 15, 21, 27 and 43 to
their rows, it is the formula (Medium). The client draws the type with the
identifier in its cell and non-unit arms; the type drawn by the flying
object of a unit caster's even identifier is unread. An identifier without
a type draws nothing. Frame choice is per identifier: (update/2) mod Phases
by default, the lightning and chain swing 4..0..4, and the teleport
counter. — R2-ENGINE-445, R2-ENGINE-447

The lightning and chain objects (28, 30) show frames 4, 3, 2, 1, 0, 1, 2,
1, 0, 1, 2, 3, 4 at update counters 1..13 and keep their last frame at any
other counter; a unit caster's object lives 13 updates, so no later update
runs (Medium). Heal (56) and Drain Life (60) objects take the target unit's
position and frame = counter only when the target is a client unit; the
effect draw method draws nothing for 56, and draws `Drain` for 60 by the
default rule. Both update arms also append a unit mark that feeds the unit
overlay generators of R2-ENGINE-692, whose result was not read. A unit
caster's one-update object, found as a client unit, shows frame 1
(Medium). — R2-ENGINE-992, R2-ENGINE-994, R2-ENGINE-692

Effect pictures drawn through the common draw head are centred on the
object's point, raised by its height and by a bilinear sample of a
signed-byte grid of the client at that point, which is the ground height
(Medium); the Blizzard object (23) adds its fall lift. The Blizzard cell
object (23) and the teleport object (54) never move in the update arms
read; the other area cell identifiers (13, 21, 27, 43) were not traced for
movement. — R2-ENGINE-993

## Blizzard presentation

The Blizzard cell effect, identifier 23 with argument 16, is the type
`Meteor\sprites`, 9 frames of 64x64. The client effect lives 16 updates:
frame 8 falls for 8 updates, then frames 0..7 play and Sfx 523 sounds once.
R2-ENGINE-394 is amended by these facts. The update period in seconds
remains Unknown. — R2-ENGINE-394, R2-ENGINE-446, R2-ENGINE-447

Opcode 0x86 reaches the client's cell arm through the dispatcher's tables in
both roots; RU skips the arm when the client's dword +0x80 is 0. The piece
stays at the cell centre and falls from 28 px to 0 px above it, 4 px per
update, then shatters in place. The effect light method's default arm,
which this identifier reaches, stores no light; the cell routine, the list
writer and the lightning block are other writers and no census of the light
writers exists (Medium). In the cell arm window read, the packet argument
is stored only as the lifetime. — R2-ENGINE-991

## Effect light

The client keeps one light byte per corner of each view cell. Each render
fills the first +0x70 / 4 bytes of the map with 0xff, unlit; that this
length is the whole map was not read (Medium). Then units, cells and
effects store values; a store replaces the byte, so among the stores read
(units, the cell routine, effects, the list) the last store wins (the order
among objects is Medium, and the census of map writers is Unknown). An effect stores one value at the corners within
radius R of its cell, the corners with i² + j² < R(R + 1): 4, 12, 32 or 52
corners for R 0, 1, 2 or 3. The values per identifier and the lightning and
chain blocks of frame × 10 at each path point are R2-ENGINE-448. A writer
flag can switch the radius stores off; who sets it is Unknown. —
R2-ENGINE-448, R2-ENGINE-995

The terrain draw skips a cell whose four corners are unlit. Otherwise each
corner takes the lower of its light byte and the terrain's own corner shade,
and the tile is shaded bilinearly between its corners. A lower value selects
a brighter shade row: row 0 triples each channel, row 64 keeps the colour and
row 95 is the darkest (Medium). Effect light values 0..46 therefore brighten
the ground (Medium). — R2-ENGINE-996
