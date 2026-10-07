<a id="spr256-256-sprite--specification-partial"></a>

# SPR256 sprites (`.256`)

A `.256` contains an optional 256-entry BGR0 palette, counted frame records
and a final u32 trailer. Each frame holds byte-control RLE pixels. Transparency
comes from skip operations; a literal palette index, including zero, is written.
— SPR256-STRUCT-001, SPR256-TRLR-021, SPR256-RLE-022, SPR256-KEY-044

Frame meanings come from the consuming registry and drawable. The file itself
does not identify animation actions or directions.

## Minimap cursor sheets (`SPR256-CURSOR-046`)

The five mode cursors are `cursors/{smove,sattack,sdefend,spatrol,scast}.256`.
Each installed resource contains one palette-bearing 16×16 frame. Registration
uses hotspot `(0,0)` and the large static-cursor period argument. They serve
overview drawing and the physical minimap handler. — SPR256-CURSOR-046

<a id="structure-standard-form--spr256-struct-001-13701384-non-empty"></a>

## Container

```
[ 1024 B palette ]              256 × [B,G,R,0] (SPR256-PAL-003/011)
repeat frames:
   u32  width
   u32  height                                                  SPR256-STRUCT-001
   u32  dataSize
   u8   data[dataSize]          RLE pixels — see below          SPR256-RLE-007
[ 4 B trailer ]  = [31-bit frameCount][bit31 = has-palette]     SPR256-TRLR-016/021

frameCount = number of frame records before the trailer         SPR256-COUNT-002
             AND trailer & 0x7FFFFFFF (the two agree exactly)   SPR256-TRLR-016/021
             (SPR256-COUNT-002's "not stored" clause superseded by that trailer count)
```

- A palette-bearing payload begins with palette entry 0, not a frame count.
- The final u32 contains the 31-bit frameCount and bit 31 has-palette flag.
  `frameCount=trailer&0x7fffffff`; read 1024 palette bytes only when
  `trailer&0x80000000` is nonzero. The set high bit also makes this word the
  frame walk's terminator sentinel. `SPR256-COUNT-002`'s clause that the count
  is not stored is superseded by this stored count; its walk derivation and
  `u32@0` clause stand. — SPR256-TRLR-016, SPR256-TRLR-021
- Installed maxima over the 1370 standard `.256` are width 640, height 480 and
  256 frames; these are not field-width or loader-admission limits.
  — SPR256-CORPUS-006

<a id="frame-pixel-data--rle-spr256-rle-007-23-79123-791-frames"></a>

## Byte RLE

`data[dataSize]` is a run-length stream of single-byte controls, `[2-bit opcode | 6-bit
count]`. Exact-width installed rows advance one row at `width`. The original
does not clamp oversized runs or validate row-control position. — SPR256-062

```
c & 0xC0 == 0x00   (0x00–0x3F)  literal      : emit next (c & 0x3F) bytes as indices
c & 0xC0 == 0x40   (0x40–0x7F)  blank rows   : emit (c & 0x3F) fully-transparent rows
c & 0xC0 == 0x80   (0x80–0xBF)  transparent  : emit (c & 0x3F) transparent pixels
c & 0xC0 == 0xC0   (0xC0–0xFF)  == 0x80 (loader aliases it); unused in data
```

The decoder tests `0x00`, then `0x40`, and sends both remaining quadrants
to transparent skip. Thus `0xc0` is an alias of `0x80`.
— SPR256-RLE-009, SPR256-RLE-020, SPR256-RLE-022

- In a structurally complete frame, every row reaches exactly `width`,
  the stream produces `height` rows, and all `dataSize` bytes are consumed.
  — SPR256-RLE-008
- Installed blank-row opcodes occur at column 0. A row control also executes at
  a partial column, preserving it in the selected original decoders. — SPR256-062
- "Transparent" here = *no pixel emitted*.

## Selected original decoder operations

The own-table receiver `R0794`, vtable `L04339+0x18`, takes
`(x,y,frame,level,mirror)` and selects `this.table + (level << 9)`.
Its last argument chooses `R0796` forward or `R0895` reversed.
The corresponding `.16a` argument is a pointer and its source units are words.
— SPR256-061

After a byte control, literals consume `n` source bytes; skip/row controls
consume none. Literal indices address `sourceTable + 2*index`. Contained runs
write pairs with one dword store, reversing their order when mirrored, followed
by a u16 store for an odd final pixel. These two decoders read no destination
pixel. Zero and 255 both lookup/write. — SPR256-063

Literal and skip counts accumulate without a width clamp. Signed column
`>=width` resets the counter and adds `stride-direction*2*width` to the current
destination, retaining any excess. Row controls subtract count from remaining
height and return on signed `<=0` before moving the destination; otherwise add
`count*stride`, preserving the partial column. Neither dataSize nor a source-end
pointer is an argument. — SPR256-062

A contained zero-count literal is a no-op. On a vertically visible clipped row,
the per-byte loop instead executes before decrementing its count, including for
zero. The finite probe stops at its source guard; this is not an original
rejection rule. — SPR256-063

Clipped literal runs read every horizontally excluded source byte but omit its
lookup and write. Vertically excluded runs advance source by `n` without reading
the literals. Fully disjoint rectangles return before reading a control.
Reversed clipping checks each pixel; the tested split-run result is invariant,
unlike the selected reversed `.16a` path. — SPR256-064

The comparison population is 88 synthetic cases and frames 0 and 1 of
`backpack/sprites.256` and `backpack/spritesb.256` from both preserved locales.
All selected forward/reversed and contained/clipped access traces agree with
separate pseudocode on supplied nonuniform memory. The `spritesb` stream is
tested through the plain receiver here; its native overlay route, native
framebuffer packing and captured pixels remain Unknown. — SPR256-065

## Palette (`SPR256-PAL-011`)

The palette holds 256 four-byte `[B,G,R,reserved]` entries. Normal palettes
use a zero reserved byte. `cursors/attack.256` and `cursors/pickup.256` have
anomalous leading regions and are not covered by the ordinary palette rule.
— SPR256-PAL-011, SPR256-PAL-018

- Index 0 is unused by the literal runs of the 1370 standard `.256`
  (`SPR256-PAL-012`, its population amended from 1376), but is not a color
  key; that claim's key clause is superseded. The blitters write
  every literal value, including zero; transparency comes from skip and
  blank-row opcodes. — SPR256-KEY-044

The zero-low-palette cases are overlays and the two cursor exceptions below.
— SPR256-PAL-013

## Overlay layers — `spritesb.256` (`SPR256-OVL-014`, draw-path clauses amended)

A `spritesb.256` overlay has one frame per base `sprites.256` frame and
uses the base sibling's palette: its literal indices are zero in its own
palette and coloured in the base's, at 99 % or more of literals in 179 of 180
pairs (median 100 %). A zeroed low palette is carried by 179 `…b.256` files;
at least one of the 180 paired `…b.256` files lacks one, and by name and frame
count it is, in the standard form, `projectiles/firebolt/spritesb.256` (`SPR256-PAL-013`, its
every-b-variant reading amended).
`SPR256-OVL-014`'s no-path-read clause is superseded: the pair is loaded and
drawn together. Its six-argument `vt+0x14` pairing is narrowed to the object
and unit body passes; the structure draw and the backpack painter pair the
overlay's `vt+0x34` with the five-argument `vt+0x18` (`SPR256-STR-040`,
`SPR256-077`).
The recovered plain half-colour painter computes
`((destination16>>1)&mask)+((palette16[index]>>1)&mask)` for opaque boundary
pixels; masks are 0x7bef (RGB565) and 0x3def (RGB555). The backpack painter
gates this overlay with Smoothing. Native full composition, mirrored paths
and coverage of other drawable families remain Unknown. — SPR256-OVL-014,
TOWN-SMOOTH-460

- **Open anomaly** (SPR256-OVL-015): `cursors/attack.256` + `cursors/pickup.256` are the
  only zero-low-palette sprites with **no base sibling**; they carry their own partial
  palette (low 16 zero) and dense content, and decode to non-image bands. Cause open.

## The equipment tree — portraits, not world sprites (`SPR256-EQUIP-042`)

`graphics\equipment` is addressed by `rom.exe`'s own literals and is **not** part of the world
sprite path. Four figure directories — `ffighter`, `fmage`, `mfighter`, `mmage` — two layer
directories under each — `primary`, `secondary` — and 987 file nodes, identical between the EN and
RU roots.

```
equipment/<figure>/primary/<7 digits>.256     784 sheets   1 frame each, 160x240
equipment/<figure>/secondary/<7 digits>.256   144 sheets   1 frame each, 160x240
equipment/<figure>/<1..31>.256                 59 sheets   the head, indexed by the face byte
```

Four of the 784 `primary` sheets yield 0 frames. Every seven-digit leaf name is produced by exactly
one u16 through the [item picture formatter](../item/appearance.md).

The one-frame, 160×240 shape is what separates this tree from a world sheet: the sixteen hero body
sheets under `units/heroes*` carry 129 to 216 frames of 24×40 to 40×48, one per direction and phase.
An equipment sheet cannot be composited into a world frame, and the routine that loads them is a
virtual method of the drawable whose identified callers are the info-window routines
([hero appearance](../hero/appearance.md)).

## Unit sheet layout (`SPR256-UNIT-024`)

A `.256` is a flat frame array; nothing inside the file says what any frame is. For the 34
`units.reg` classes the sheet is **block-addressed**, and the block bases are arithmetic on the
class's own phase counts, computed fresh at every draw. With `MB, MV, AT, DY, BN, ID` the
class's `MoveBeginPhases, MovePhases, AttackPhases, DyingPhases, BonePhases, IdlePhases`:

```
Flip == 0  ->  S = 16, D = 8            Flip != 0  ->  S = 9, D = 5

  0                       stand    S frames        index = the facing itself (16-way)
  S                       move     D*(MB + MV)     dir*(MB+MV) + MB + moveTL[phase mod len]
  S + D*(MB+MV)           attack   D*AT            dir*AT      + attackTL[phase]
  S + D*(MB+MV+AT)        dying    D*DY            dir*DY      + phase/2
  S + D*(MB+MV+AT+DY)     bone     D*BN            dir*BN      + (stage - 2)
  S + D*(MB+MV+AT+DY)     idle     D*ID            dir*ID      + idleTL[phase]
```

`dir` is 16-way for the standing block (`(facing - 8) & 0xf`) and 8-way everywhere else
(`((facing - 8) >> 1) & 7`). On a `Flip` sheet only directions `0..8` / `0..4` are stored and
the rest are drawn mirrored. **Bone and idle share one base** — no shipped class carries both.
The dying and bone blocks are read out of `classes[Dying]`'s sheet with *its* counts and *its*
`Flip` (`REG-UNITS-050`).

Sheet length is therefore `S + D*(MB+MV+AT+DY+max(BN,ID))`, which reproduces the frame count of
**33 of the 34 shipped unit sheets exactly** and overshoots none. (The 34th, `Unit2`
*Unarmed Fighter with Shield*, spells a `Flip = 1` sheet but inherits `Flip = 0`; it is a defect
in the registry, and four shipped map placements reach it.)

**Every one of the 34 unit sheets mixes frame sizes**, so the draw anchor must be built from the
frame actually being drawn (`SPR256-FRAME-023`, `TERR-SPR-043`).

## Structure sheet layout (`SPR256-STR-040`, `SPR256-STR-041`)

A `structures.reg` class opens **two** sheets, `graphics\structures\<File>.256` and
`<File>b.256` — the `spritesb` overlay class of `SPR256-OVL-014`, drawn with the same frame
index through `vt+0x34` and gated on `[L04367]`. The base sheet is drawn by the
five-argument own-table `vt+0x18`, not `vt+0x14`, so `SPR256-OVL-014`'s six-argument pairing is
narrowed to the object and unit body passes (`SPR256-STR-040`, `SPR256-061`).

Neither sheet is block-addressed like a unit's and neither is `Index`-addressed like an
object's. It is a **row-major grid of tile-sized frames**, `TileWidth` columns wide and
**`FullHeight`** rows tall — `FullHeight`, not `TileHeight`, so the grid is taller than the
footprint by the structure's overhang. **All 130 structure sheets in the install carry
uniformly `32×32` frames**, which is why no anchor pixel is needed: a frame *is* a map cell.

```
cellIndex(k, c) = k * TileWidth + c          k = image row 0..FullHeight-1
                                             c = image column 0..TileWidth-1

  [ 0                          , TW*FH )         BASE    frame = cellIndex
  [ TW*FH                       , TW*FH + P*L )  ANIM    frame = TW*FH + (p-1)*L + rank(c,k)
  [ frames - TW*FH              , frames )       RUIN    frame = frames - TW*FH + cellIndex

  TW = TileWidth   FH = FullHeight   L = AnimMask's non-'-' count   P = max timeline value
  p    = the class's current animation value (0 = not animating this frame)
  rank = how many non-'-' AnimMask cells precede this one
```

The ruin block is addressed **from the end of the file**; nothing stores its base. Which
block a cell takes is decided by `TERR-STRUCT-102`. Shipped identity:
`frames = k·TW·FH + P·L`, `k ∈ {1,2}`, on **64 of 66** classes — `k = 2` on 39 (a ruin grid
ships), `k = 1` on 25 (none does, and `frames − TW·FH` is then 0, so a destroyed structure
re-draws its intact art). The two exceptions are `bridge1v` and `bridge2`, whose drawable is
a C++ subclass with a 9- and a 14-frame nine-patch selector (`TERR-STRUCT-105`).

## Backpack sheet frame selection (`SPR256-077`)

As with every `.256` sheet (see above), frame meanings come from the consuming registry and
drawable, not the file; no claim in this repository reads `backpack/sprites.256`'s own bytes,
so its frame count and any addressing scheme inside the file are unestablished here. The
recovered painter (`TOWN-SMOOTH-460`, `R1795`, CBackPack's own vtable `+0x28` slot)
reads the frame index from `CBackPack+0x20` at both its draw calls — the base draw and the
Smoothing-gated boundary/overlay draw — pushed in the same argument position, third of five
and third of six respectively, that `TERR-LIGHT-059` and `SPR256-061` name as the `frame`
argument for this blitter family, through a loaded sprite object sharing their named vtable
identity (`SPR256-077`). The sibling `+0x2c` slot is a separate shadow draw and was not shown
to read this field.

What writes `CBackPack+0x20`, and from what [Sack field](../item/sacks.md#ground-sack-drawn-frame),
is item-domain and is not repeated here. Because the sheet's own frame count is unestablished,
whether the frame-index clamp at 5 (`ITEM-136`) saturates at the sheet's last frame or
addresses past populated frames is not resolved by any claim here.

## Known variants / exceptions

- **No-palette variant** (6 projectile arrows): the 1024-byte palette is absent;
  frames begin at offset 0. This is the **same format with the trailer's has-palette bit
  clear** (SPR256-TRLR-021) — the loader simply skips the palette read; bit 31=0 so the
  trailer is a plain `u32` frame count (16). The external palette source is still TBD
  (SPR256-VAR-004).
- **Frame + trailing section** (8 UI / global-map / goblin-arrow sprites): a normal
  palette + one frame + its count-trailer (`0x80000001`), then an **appended secondary
  section bracketed `0x80000001 … 0x80000001`** (8/8). The inner block is opaque — not a
  frame, not a same-width RLE image; 4–16 B (small UI/projectile) or 303–2521 B dense
  (world-map tiles + shop frame, using `0xC0-0xFF` bytes). **The shipped loader ignores
  it**: the `CSprite256` ctor indexes exactly `frameCount = 1` frames (the final
  `0x80000001`), so the inner is loaded but never indexed/decoded/blitted — runtime-inert
  (SPR256-EXC-020). Authoring origin still open but non-runtime. (SPR256-EXC-005/017)

<a id="open-own-experiments"></a><a id="closed"></a>

## Unknowns

- The `spritesb` overlay uses the base palette; plain-loop half-colour arithmetic
  is measured, while native full composition and unmeasured paths remain Unknown.
  `SPR256-OVL-014`'s clause that nothing loads the pair together is superseded;
  whether the structure and backpack `vt+0x18` pairings shade the overlay through
  the base's table row is Unknown. The pairing of the one paired `…b.256` with
  a nonzero low palette rests on name and frame count, since no pair list is
  committed (`SPR256-PAL-013`). — SPR256-OVL-014, TOWN-SMOOTH-460
- The attack/pickup resources have a nonzero reserved byte in the nominal
  palette region. That region is not an ordinary palette; the cause and
  corresponding decoder behavior remain Unknown. — SPR256-OVL-015,
  SPR256-PAL-018
- Object sheets use their class's Index, Index+timeline[phase], or 0; they
  have no unit-style block layout. — TERR-SPR-042, REG-OBJ-046
- External palette selection is specified for the named projectile paths by
  `PAL-PROJ-011`; other no-palette variants remain Unknown.
  — SPR256-VAR-004

## Read and write sequence

1. Read the final u32 trailer: count is `trailer & 0x7fffffff`, and bit 31
   selects whether the first 1024 bytes are a palette.
2. Start the frame cursor after that optional palette. Read exactly count
   records, each `[u32 width][u32 height][u32 dataSize][dataSize bytes]`.
3. Decode each frame with byte controls. Bound literal input and destination
   rows; blank-row controls begin at column zero.
4. Keep runtime-inert residual bytes distinct from indexed frames. The
   described appended-section variant does not add live frames.

To emit the ordinary form, write the optional BGR0 palette, the selected frame
records and a trailer whose count equals the number emitted. Encode literal
runs and skips with six-bit counts; start each blank-row run at a row boundary.
Use each frame's own dimensions. Installed maxima are not format limits.
The authoring origin of appended residual sections and the two anomalous
cursor files is not a general writer contract. — SPR256-STRUCT-001,
SPR256-TRLR-021, SPR256-RLE-020, SPR256-RLE-022, SPR256-EXC-020
