# Sprite encoding and pixels

[Reference](format.md)

## Sprite container

```
[ 1024 B leading block ]           per-file 256 × [B,G,R,0] palette table
                                                                   SPR16A-PAL-008
repeat frames:
   u32  width
   u32  height                                                    SPR16A-STRUCT-001
   u32  dataSize
   u8   data[dataSize]             16-bit word-RLE — see below     SPR16A-RLE-002
u32  trailer                       [31-bit frameCount][bit31 = has-palette?]
                                                                   SPR16A-TRLR-012 (partially retracted for .16)
```

- **`frameCount = trailer & 0x7FFFFFFF`, and `trailer & 0x80000000` gates the 1024-byte
  palette read** — both quoted from `R0753`, the constructor `.16a` shares with `.256`
  (SPR16A-TRLR-012, SPR256-TRLR-021). Read the trailer; do not only walk. The `16+16`
  alternative is refuted: the loader reads the trailer with a single 4-byte read and touches
  it only as a dword.
- **For `.16`, distinguish the stored word, allocation request and indexing predicates.**
  The full u32 is stored unchanged. The pointer-table request is
  `(raw << 2) mod 2^32`; entry requires `int32(raw) > 0` and continuation
  compares the 32-bit index to the raw word with signed `JL`. These facts do
  not determine native allocator outcomes. — SPR16A-070, SPR16A-071
- **The frame-indexing block has no frame-header validation.** Its
  `cursor += 12 + dataSize` step has no size threshold, width/height read or
  buffer-end comparison. Zero dataSize advances by 12. File-open and I/O
  guards are outside that block; the old whole-reader no-rejection wording
  is partially retracted. — SPR16A-RDR-017 (partially retracted), SPR16A-071
- `.16` copies from resource offset 0 and stores a null palette pointer.
  Its adapter passes frame dimensions, frame data at header plus 12, and
  the caller's ramp to the byte decoder; it does not pass the trailer.
  — SPR16A-072
- Every frame is compressed: `dataSize != w*h` and `!= w*h*2` (SPR16A-STRUCT-001).
- Installed `.16a` sheets have uniform frame dimensions. The frame grammar
  still stores dimensions on every record; this observation is not a general
  format limit. — SPR16A-BOUND-016

## Word RLE

`data[dataSize]` is a run-length stream of **u16 LE control words**,
`[2-bit op | 14-bit count]`. Exact-width installed rows advance one row at
`width`. The original does not clamp runs or validate row boundaries; the
operations below amend that shorthand. — SPR16A-079

```
cw >> 14 == 0b00   literal   : the next n u16 words are source-table byte offsets
cw >> 14 == 0b01   blank rows: emit n fully-transparent rows (at a row boundary)
cw >> 14 == 0b10   skip      : emit n transparent pixels in the current row
cw >> 14 == 0b11   BLANK ROWS again — unused in ROM1 data, but the dispatch is two
                   sequential bit tests and bit 14 is tested FIRST, so both-bits-set
                   takes the blank-rows arm                      SPR16A-RLE-002
            n = cw & 0x3FFF
```

> **Do not carry `.256`'s alias across.** `.256` and the `.16` byte blitter both fold their
> fourth quadrant into *skip*; the `.16a` u16 blitter folds it into *blank rows*, because the
> two routines test the same two control bits in opposite order. Nothing shipped exercises
> either alias.

- Each row's literal and skip counts sum to width. The stream describes
  height rows and has no remaining byte after the last row. — SPR16A-RLE-003
- Installed blank-row opcodes start at column 0. The original also executes
  a row control at a partial column and preserves that column. — SPR16A-079
- The `0b11` quadrant takes the blank-row arm, because bit 14 is tested
  before bit 15; installed streams do not use this alias. — SPR16A-RLE-002 The count field is 14 bits (maximum 16383); smaller
  installed maxima are not encoding limits.
- A transparent span emits no pixel.

## Literal-word colour — the display model (`SPR16A-PIX-011`)

A literal u16 word is a raw source-table byte offset. Installed words are even
and inside its `[16][256]u16` shape. The original does not mask high bits or
reject odd offsets; the earlier universal wording is narrowed to the installed
population. For that population: — SPR16A-PIX-011, SPR16A-080

```
paletteIndex = (word >> 1) & 0xFF     // bits 1..8   (0..255)
level        = (word >> 9) & 0x0F     // bits 9..12  (0..15)  -- shipped range 1..15
src          = srcLUT[level][paletteIndex]      // palette scaled by (level+1)/16
pixel        = src + destTable[1 + level][old_fb]   // u16 add, then framebuffer write
```

**The normal-memory level is the pixel's alpha** (SPR16A-ALPHA-025, partially
retracted for the low-memory generalization, 15-usable-steps limit and
native-unreachability assertion;
SPR16A-080). The destination
table `R0788` builds has **17** rows and its row `k` is every representable pixel
scaled by `(16 - k)/16`; the blitter's base skips exactly one row, so the row it uses for
level `L` scales the old pixel by `(15 - L)/16` while the source row scales the art by
`(L + 1)/16`. The two weights sum to 16, so the write is an exact 16-step linear blend:

```
out = palette[index] * (level+1)/16  +  destination * (15-level)/16
```

In the normal table, level 15 is the only opaque value (its destination row is
all zeros). Installed level 0 absence does not exclude it from the decoder:
literal word 0 executes the same lookup and write. Three installed projectile sheets use a single
palette index and encode their variation through alpha (`SPR16A-PROJ-026`). Straight-alpha RGBA reproduces
the normal-memory weighting; native channel packing and captured pixel agreement
remain Unknown. A binary mask cannot express the weighting. — SPR16A-080, SPR16A-083

The `/18` source variant and a destination table quantized to 8 192 entries a row are **not
a display mode**: `R0788` sets the flag that selects them from `GlobalMemoryStatus`,
under 24 MB of physical RAM (PAL-MODE4-010). Its native-unreachability assertion
is partially retracted: this predicate does not establish native feasibility,
which remains Unknown.

The low-memory decoder indexes destination row `L`, using `old >> 3`, without
the normal branch's added row. Its source table uses the separate `/18` builder
branch. The normal `1+L` complementarity formula does not apply. — SPR16A-080

## Selected original decoder operations

The own-table receiver `R1785` takes `(x,y,frame,tableOverride,mirror)`.
Zero override selects `this+0x1c`; other values are raw table pointers.
The last argument selects `R1518` or `R1784`. The `.256` class instead
interprets that fourth argument as a table level. — SPR16A-078

For positive finite geometry and nonwrapping arithmetic, start the destination
at `surface + y*stride + 2*x`, plus `2*(width-1)` when reversed. Read each
control as u16 and advance the source two bytes. A pixel skip advances the
destination by `direction*2*n` and adds `n` to the consumed-column counter.
A literal run does the same per pixel, with the lookup operation below.
— SPR16A-078, SPR16A-079

After an operation, signed consumed-column `>=width` resets that counter and
adds `stride-direction*2*width` to the current destination. It advances only
once; excess from an oversized run remains in the pointer. Row controls
subtract `n` from remaining height and return on signed `<=0` before moving
the destination. Otherwise they add `n*stride` and preserve the partial column.
No dataSize or source-end argument is passed to these decoders. — SPR16A-079

An executed literal reads old destination u16, source word `W`, destination
contribution u16, and source contribution u16, then adds modulo 65536 and writes
u16. These are byte offsets: — SPR16A-080

```
if lowMemorySelector == 1:
    destOffset = 2*(old >> 3) + ((W << 5) & 0x3c000)
else:
    destOffset = 2*entriesPerRow + 2*old + ((W << 8) & 0x1e0000)
destination = read16(destinationTable + destOffset)
source = read16(sourceTable + W)
write16(d, source + destination)
```

Clipping advances over excluded source words without reading them. The reversed
clipped path uses the coordinate immediately beyond the run to trim its left
tail, then uses the shortened count in the right test. A four-pixel run at
`x=0` and clip `[0,3)` writes `3,2,1`; splitting it into `2+2` writes `2,1`.
For clip `[1,4)`, the full run writes `3,2`. This path is not the assumed mirror
of the forward clip algorithm. These are local instruction observations.
— SPR16A-081

A contained zero literal enters the decrementing loop. A clipped forward zero
literal can enter after the row/horizontal tests; a clipped reversed nonpositive
trimmed run skips the loop. Finite source guards stop the probes; the original
has no corresponding source-bound rejection. — SPR16A-082

The comparison population is 111 synthetic cases and frames 0 and 1 of
`cursors/attack/sprites.16a` and `cursors/cast/sprites.16a` from both preserved
locales. Supplied nonuniform memory agrees with separate pseudocode for the
selected forward/reversed and contained/clipped paths. This does not supply
native framebuffer packing or captured pixel agreement. — SPR16A-083

Installed literal words leave bit 0 and bits 13–15 clear. Bit0 being zero
makes the LUT offset even. This installed pattern is not an arbitrary-input
validation rule. The optional 1024-byte palette has 256 entries of `[B,G,R,0]`.
— SPR16A-PIX-011, SPR16A-PAL-008

The direct-RGB models (`SPR16A-PIX-004`, `SPR16A-PIX-005`), unused-palette
interpretation (`SPR16A-PAL-006`), high-byte/gamma model (`SPR16A-PIX-009`)
and 5+7 or 8×32 bank models (`SPR16A-PIX-010`) are retracted. The established
address split is 4 level bits and 8 palette-index bits.

## Structural bounds

| Quantity | Encoding |
|---|---|
| Frame width, height, dataSize | u32 each |
| `.16a` frame count | Low 31 bits of final trailer |
| `.16` stored count word | Complete final u32; signed indexing and wrapping allocation are separate operations |
| `.16a` RLE count | 14 bits, maximum 16383 |
| `.16` RLE count | 6 bits, maximum 63 |

An indexed record occupies `12+dataSize` bytes. Check those extents against
the containing input and the decoded cursor against the frame bounds.
`.16a` data is a u16 stream and therefore has even byte length.
The original loader does not impose these defensive extent checks.
Installed maxima, uniform sheet dimensions and unused literal bit patterns
are not arbitrary-input admission limits. — SPR16A-RDR-017 (partially retracted), SPR16A-BOUND-016

## Unknowns

- **The framebuffer's own channel widths and shifts** — six runtime globals the table
  builders and the blend all read, so the *absolute* colour of a drawn pixel is derived only
  up to the assumption that they are 5-6-5. The same residual applies to the `.16` glyph
  ramps. What is **no longer** open: the arithmetic of `src + dest` itself, and the pair of
  variants that used to be recorded here as display modes — they are a low-memory fallback
  (SPR16A-ALPHA-025, partially retracted for its low-memory generalization and
  native-unreachability assertion; PAL-MODE4-010, partially retracted only for
  native unreachability; SPR16A-080). The predicates and local operations stand;
  native low-memory reachability remains Unknown.
- **Which of the thirteen ramps a given UI string gets** — traced to the blit argument, not
  through the UI code that chooses it.
- **The `^` escape** in the `.16` `DrawText` — a two-character sequence that diverts to a
  separate draw call; located, not decoded.
- **Animation frame-block roles** — as with SPR256, per-class block sums vs registry
  phase counts, once cross-referenced.
- `font5.16a` has no identified loader; its runtime use remains Unknown.
- The named `.16` and `.16a` entry identities are verified, but their
  retained reference enumeration is limited to its repaired, disassembled
  project. Bytes never disassembled and other entry paths are outside that
  result; global loader exclusivity remains Unknown. — SPR16A-RDR-017

## Read and write sequence

1. Select `.16a` or `.16` from the resource contract; do not infer a shared
   codec from the frame-header shape.
2. Read the final trailer. For `.16a`, extract the low-31-bit count and optional
   palette flag. For `.16`, retain the complete word and no palette.
3. For ordinary positive-count resources, walk `12+dataSize` records from
   the selected origin. The original `.16` allocation request and signed
   loop predicates remain distinct; their arithmetic does not prove a
   native traversal completes. — SPR16A-070, SPR16A-071
4. Decode `.16a` with word controls and its shade-table blend, or `.16` with
   byte controls, low-nibble-first literals and the caller's text ramp.
5. For text, load the paired `.dat` advance table and use its space rule.

The named positive-count `.16` fonts' trailing sections add no pointers to
the selected frame table. Full-buffer reading and copying are separate:
the loader reads the full resource, and the selected copy constructor
copies the full stored byte length if invoked. Other runtime uses of those
tails remain Unknown. — SPR16A-FONT-014 (partially retracted), SPR16A-072

An ordinary `.16a` emitter writes the optional palette, frame headers/RLE data
and count/flag trailer. A `.16` emitter writes its glyph records and plain
count trailer; an odd literal run uses the final high nibble as zero padding.
Widths, encoded lengths and counts must describe the emitted records. No
writer for arbitrary retained overwrite tails is established.
— SPR16A-TRLR-012 (partially retracted), SPR16A-RLE-002, SPR16A-FONT-013,
SPR16A-FONT-014 (partially retracted), SPR16A-FONT-018

No original standalone `.16` producer was identified in the bounded game
plus the editor paths searched. Dynamic filenames and other entry paths remain
outside that result. — SPR16A-073
