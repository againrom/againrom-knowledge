# Tile resources and selection

[Reference](format.md)

## Location & format

The terrain images live inside **`graphics.res`** (the `RES-MAGIC-001` `&YA1` archive) under
both **`terrain/`** and **`terrain.3d/`**. Each family has the same 53 names and
geometry in EN and RU. They use **uncompressed 8-bpp Windows BMP** layout (magic `BM`,
14-byte `BITMAPFILEHEADER` + 40-byte `BITMAPINFOHEADER`, 256-entry RGBQUAD palette,
`bfOffBits = 1078 = 14 + 40 + 1024`). The 52 tile headers in each family
understate file length by 1024 bytes; use the bounded geometric pixel span,
which is present. Dirt has two trailing bytes. (`world.res` holds terrain
config `data/map.reg`.) — TERR-LOC-001 (route wording amended), TERR-FAMILY-188

| File(s) | Count | BMP W×H | 32×32 sub-cells | Role |
|---------|-------|---------|------|------|
| `dirt.bmp` | 1 | 32×128 | 4 | impassable-tile composite overlay |
| `tile1-00..15` | 16 | 32×448 | 14 | Land group (Grass / Cracked / Sand / Savanna) |
| `tile2-00..15` | 16 | 32×448 | 14 | Stones / Cracked-Stones / Flowers-Savanna / Mountain |
| `tile3-00..15` | 16 | 32×256 | 8 | **Water** (animated) |
| `tile4-00..03` | 4 | 32×448 | 14 | **Road** |

Each BMP is **32 px wide** and is a **vertical strip of 32×32 sub-cells** (top-to-bottom;
BMP rows are bottom-up per the standard, so sub-cell `k` occupies the *k*-th 32×32 block
counting from the strip origin the game reads — pixel byte offset `8 + k·0x400` into the
loaded image buffer). The resource families differ in both palettes and
index data: 647,325 of 651,264 paired indices and 647,586 RGB samples differ
per locale. All 53 pairs change 255 palette RGB entries. Every `terrain.3d`
file shares one palette; the 52 ordinary tiles share another, and ordinary
dirt differs from them in 41 RGB entries (reserved bytes equal).
Corresponding EN/RU files are byte-identical within each family.
— TERR-FAMILY-188

For `terrain.3d minus terrain`, aggregate channel deltas are R -69..113,
G -58..99 and B -58..164. Positive/zero/negative counts respectively are
R 647,384/3,698/182, G 647,523/3,678/63 and B 647,533/3,680/51. These are
decoded source RGB comparisons, not native rendered brightness. Geometry
equality does not make the two sets interchangeable. — TERR-FAMILY-188

## Loader (`rom.exe R0516`)

The loader builds `tileG-VV.bmp` (`%d-%d` /
zero-padded `%d-0%d` for `V<10`) and loads each into a **128-slot pointer array `tiles[]`**
at the global `L10218`:

```
tiles[(G-1)*16 + V]  <-  tileG-VV.bmp      G = 1..8, V = 0..15   (absent file -> null)
[L10220]         <-  dirt.bmp
[L10222]         =   first non-null block leader + 0x14     (object address)
```

On disk only `tile1/2/3` (16 each) + `tile4` (0..3) exist. Source pointers and
precomposed work surfaces have different lifetimes. If selector
`L10964 & 0x02` is nonzero, the loader first loads `terrain.3d` and, for a
nonzero group mask, precomposes its work surfaces (`L10221[]`). Graphics
object `+b84 == 8` selects the palette/index arm; other values select the
other surface arm. Both normal arms delete and zero all 128 temporary
source pointers and dirt, then fall through to the ordinary `terrain/`
loading block at `L10965`. With the selector mask clear, the loader enters
that ordinary block directly. The final source table therefore holds the
ordinary pass's results. Final default-handle search tests indices
0,4,...,124; it stores the first non-null block leader's address plus 0x14,
or zero. — TERR-LOAD-002 (order and handle search amended), TERR-FAMILY-187

This is sequential optional precomposition, not a recovery/fallback choice.
Normal startup writes zero to `L10964` before graphics initialization.
The complete file-backed embedded-address census for this dword and its
overlapping addresses finds 12 references, with one other direct write:
`OR 1`, guarded by mask 0x02 and a nonzero graphics member `+12b0`.
The mask-0x02 arm initializes an object whose call chain creates IDirect3D.
Its `-systemmemory` and `-emulation` options change that object's members;
`-safevideo` changes a separate global. These are not discovered setters
of mask 0x02, nor proof that the branch is hardware-only. — TERR-FAMILY-187

No setter of mask 0x02 was found in that address-form population. Computed
pointers, bulk copies and native activation remain Unknown. Normal startup
zero assumes the nonvolatile EBX convention across intervening calls; no
native startup trace is claimed. The previous assertion that `terrain.3d`
is the shipped route is withdrawn. — TERR-FAMILY-187

**The loader takes one argument, and the map supplies it.** `R0516` is `cdecl` with a
single dword, read at `[esp+0x558]` and used as a 32-bit mask: the body runs exactly 32
iterations over `tiles[]` in steps of four slots, and enters iteration `i` only when
`mask & (1<<i)`. Iteration `i` fills `tiles[4i .. 4i+3]`, which is precisely the block the
render mapping reaches with strip group `g = i` and blend column `b = 0..3`, so one mask bit
is one value of `g`, i.e. one file group `G = (i>>2)+1` with variants `V = (i&3)*4 .. +3`.
The value comes from the `.alm` type-0 record: payload `+0x18` → landscape slot `P+0x28`,
pushed at the end of the map-build. Over 72 shipped maps it takes eight values, all
`≤ 0x1fff` — `0x1fff`×44, `0x0fff`×10, `0x00ff`×4, `0x0fbf`×4, `0x1fdf`×4, `0x0f7f`×2,
`0x1f1f`×2, `0x1fbf`×2 — never a bit above 12, which decodes to groups 1–3 entire plus
group 4 variants 0–3: the same set the disk holds. A clear bit leaves all four images that
tile-word group can select null (`TERR-LOAD-152`).

## Tile-word → graphic (render mapping)

Four render routines — `R0539`, `R1814`, `R1815`, `R1816`
— extract the following fields from the render-grid word. Their branch rules
differ: `R1814` skips non-water groups before its pointer lookup.
— TERR-GFXBOUND-165, TERR-WATERPASS-167

```
g    = (w & 0x1fff) >> 6          strip group      (bits 6..12; installed range 0..12)
b    = (w >> 4) & 3               blend column     (bits 4..5)
sub  = w & 0xf                    sub-cell index    (bits 0..3)

image = tiles[ g*4 + b ]          == tileG-VV.bmp with:
          G = (g >> 2) + 1        file group  (installed 1..4; loader covers 1..8)
          V = (g & 3)*4 + b       file variant (00..15)
src   = image.pixels + 8 + sub*0x400        one 32x32 8-bpp cell (0x400 = 1024 B)
```

The graphic arithmetic admits `g=0..127`, but the 3d loader fills only
128 pointer slots, covering `g=0..31`. Starting from word zero, bit 10
selects slot 64; bits 11 and 12 select slots 128 and 256, outside that
array. The selected full-render kernels contain no intervening bound check.
This is an address boundary, not a new terrain type or a graceful refusal.
The bit interval of TERR-IDX-003 is amended; its arithmetic is retained.
— TERR-GFXBOUND-165

Loading is a separate decision. Mask bit `i` admits slots `4i..4i+3`.
A masked-out preexisting slot is destroyed and set to null. An admitted
slot still needs a usable bitmap from the resource constructor. A null
selected pointer reaches a pixel-buffer-field read at pointer `+0x10` in
the inspected full-render arm; bounded execution faults there. Native
resource failure behavior and pixels for unpopulated groups remain Unknown.
— TERR-SLOTADMIT-166, TERR-GFXBOUND-165

With animation disabled, the water arm normalizes groups 8–11 to group 8.
The `R1814` pass instead skips every non-water group before lookup;
that skip does not establish what the other passes draw.
— TERR-WATERPASS-167

The light parser first clears bit 13. The render-grid RLE updater can later
replace it with `(word & 0xdfff) | runFlag`, where `runFlag` is 0 or
`0x2000`; the other 15 bits survive. The write kernel is established,
while first-message timing and subsequent native ordering remain Unknown.
— TERR-RLEBIT-168

Mapping of strip group → file group → terrain (terrain names via `ALM-TERR-015`):

```
g  0  tile1-00..03  Grass  / Land        g  4  tile2-00..03  Stones          / Land
g  1  tile1-04..07  Cracked/ Land        g  5  tile2-04..07  Cracked         / Stones
g  2  tile1-08..11  Sand   / Land        g  6  tile2-08..11  Flowers         / Savanna
g  3  tile1-12..15  Savanna/ Land        g  7  tile2-12..15  Mountain        / Stones
g 8..11  tile3-*    Water (animated)     g 12  tile4-00..03  Road            / Land
```

- **Water (`g 8..11` = `tile3`)** is animated (fully decoded by `TERR-ANIM-006…TERR-ANIM-010`; see "Water
  animation" below). The renderer replaces the group with `8 + phase`,
  `phase = (g + (worldCol+1)*worldRow + (animCtr>>2)) & 3` when enabled (`[L05659] ≠ 0`),
  else `phase = 0`. Water uses only 8 sub-cells (tile3 is 32×256).
- **Impassable tiles (bit 13 set, non-water)** are composited over `dirt.bmp` before blit
  (`dirt` sub-cell `= (col + row*5) & 3`), giving the "blocked ground" look. The composite is a
  **transparent-index (index 0) keyed overlay**, not a blend: copy the terrain sub-cell, then
  paint the dirt sub-cell's non-zero pixels over it (`R1792`; `TERR-DIRT-017`). Drawn shaded.
- The four **corner altitudes** (type 2 Altitudes grid, `ALM-GRID-013`) choose a flat vs
  sloped blit **and displace the pixels** — see [cell geometry](geometry.md).
- Every tile is drawn **relief-shaded** — see [terrain lighting](lighting.md).

## Installed tile domain

Installed maps use strip groups 0..12. Groups 1–3 of the resource filenames
are complete; group 4 supplies variants 00–03. Land strips contain 14
subcells and water strips contain 8. A decoder must use the selected strip's
own subcell count and reject or handle missing resources explicitly; these
installed populations are not a new upper bound on the tile-word field.
— TERR-VER-005

## Tile draw sequence

Source initialization follows the Loader section's selected-group mask and
optional preceding precomposition. — TERR-FAMILY-187, TERR-LOAD-002

```
# one-time: apply the map's selected-group mask
# if selector & 0x02: precompose terrain.3d sources, then release temporary bitmaps
# final sources: terrain/*.bmp -> tiles[(G-1)*16+V]; terrain/dirt.bmp -> dirt
def draw_cell(w, col, row):
    g   = (w & 0x1fff) >> 6
    b   = (w >> 4) & 3
    sub = w & 0xf
    if 8 <= g <= 11:                      # animated water
        g = 8 + (0 if anim_off else phase(g,col,row,anim_ctr))
    img = tiles[g*4 + b]                  # == tileG-VV.bmp, G=(g>>2)+1, V=(g&3)*4+b
    cell = img.subcell(sub)              # 32x32 block at pixel offset 8 + sub*0x400
    if (w & 0x2000) and not water:        # impassable, non-water
        cell = composite(cell, dirt.subcell((col + row*5) & 3))
    blit(cell, x=col*32, y=..., corner_heights=Altitudes[...])
```

## Water animation (`TERR-ANIM-006…TERR-ANIM-010`)

The water group is animated on a fixed-timestep logic clock (all from `rom.exe`):

- **Sequence.** A water cell keeps `b=(w>>4)&3` and sub-cell `w&0xf`; only the group rotates,
  so the drawn image walks `tile3` variant `V = phase*4 + b`, i.e. `tile3-{b, b+4, b+8, b+12}`,
  as `phase` steps `0→1→2→3→0`. The 16 `tile3` files = 4 phases × 4 blend columns. The
  `(worldCol+1)*worldRow` term (scroll-adjusted) offsets neighbours into a diagonal ripple.
- **Counter.** `animCtr = *(mapObj+0xa70)`; its only writes are init-0 (`R0391`) and
  `+1` per logic tick (`R0334`, the handler of window message `0x401`). `phase` uses
  `animCtr>>2` → one variant every **4 ticks**, full cycle every **16 ticks**.
- **Cadence.** The paced loop `R0454` fires one `0x401` per `dtMs = 1000/tps` ms;
  `SetGameSpeed` (`R0559`) maps speed index `0..8` → `tps ∈ {8,10,12,14,16,20,24,28,32}`.
  Map-load defaults to **index 4 → 16 tps → 62 ms/tick → ~250 ms/frame → ~1 s/cycle**; the +/-
  keys and config key 2 change it. Per-frame ms = `4·(1000/tps)`, per-cycle = `16·(1000/tps)`.
- **Enable.** the global at `L05659` (static default `1`); cleared by `-noanimation` / `-detail0` (then
  `phase=0`, water static). Each phase selects a present tile3 resource. — TERR-ANIM-006…TERR-ANIM-010
