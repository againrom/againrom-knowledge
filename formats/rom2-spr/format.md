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
shape determines where to read/write the table bytes; color-value meanings
and ROM2 table-building modes remain Unknown. — R2-ASSET-008

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

## Blizzard presentation

The Blizzard cell effect, identifier 23 with argument 16, is the type
`Meteor\sprites`, 9 frames of 64x64. The client effect lives 16 updates:
frame 8 falls for 8 updates, then frames 0..7 play and Sfx 523 sounds once.
R2-ENGINE-394 is amended by these facts. The update period in seconds
remains Unknown. — R2-ENGINE-394, R2-ENGINE-446, R2-ENGINE-447
