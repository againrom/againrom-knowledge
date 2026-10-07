<a id="rom2-sprite--palette-containers--identity-survey"></a>

# ROM2 sprite and palette containers

The preserved `.16a`, `.256`, `.16` and `.pal` resources have the container
shapes below. This reference establishes frame geometry and residual regions;
it does not establish ROM2 RLE, pixel, color or runtime draw semantics.
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

<a id="not-yet-surveyed"></a>

## Unknowns

RLE opcodes, pixel/color conversion, native use of residual bytes, arbitrary
malformed resources and other locales are outside the established contract.
Related ROM1 references: [SPR16A](../spr16a/format.md),
[SPR256](../spr256/format.md), [PAL](../pal/format.md).
