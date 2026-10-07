# Cell geometry and far edges

[Reference](format.md)

## Cell geometry — where the pixels land (`TERR-GEOM-031` (amended)…`TERR-GEOM-036` (amended, superseded), of which `TERR-GEOM-034` is amended)

The terrain raster is **not** flat. Each grid **vertex** is projected to a destination row and a
cell is drawn as the quad between its four projected corners.

```
vertex screen Y   V(c,r) = r*32 - (int8)Altitudes[worldCol + W*worldRow]
                            ^ screen row, not world row      ^ SIGNED (sign-extended byte), raw type2 byte
corners           yTL=V(c,r)  yTR=V(c+1,r)  yBL=V(c,r+1)  yBR=V(c+1,r+1)
destination X     col*32 + i, i = 0..31 — always 32 columns, NO horizontal term
```

**A larger altitude subtracts from the destination row index** (`R1818` @`L10245`), i.e.
it moves the pixel toward the clip rectangle's `top` field — up the screen in the Win32 device
coordinates the clip rect is expressed in. The mesh lives at `CMapView+0xb4` and is rebuilt every
frame for a several-cell over-scan margin (`TERR-EDGE-026`); the *derived* `+0xc0` grid (mean of
four heights) belongs to a different consumer and must not be used here.

The flat blitter is selected exactly when
`yTL==yTR && yBL==yBR && yTL+32==yBL`: all four corner altitudes are equal.

```
FLAT   (R1506): the axis-aligned rectangle [x,x+32) x [yTL,yTL+32);  src(i,j) -> (x+i, yTL+j)

SLOPED (R1507): per destination column i = 0..31
    top(i)    = yTL walked toward yTR by the step table below
    bottom(i) = yBL walked toward yBR by the same table
    H         = bottom(i) - top(i)
    if H <= 0:  the column is SKIPPED — no pixel, no clamp, no fallback
    else rows top(i) .. bottom(i)-1          (top INCLUSIVE, bottom EXCLUSIVE)
         srcStep  = (32<<16) / H                       integer divide, 16.16
         srcRow(j)= (j * srcStep) >> 16                floor; j = 0..H-1, v starts at 0
         pixel    = LUT16[ level(j) + src[i + 32*srcRow(j)]*2 ]
         level(j) = (levelTop + 0x100 + j*((levelBot-levelTop)/H)) & 0xfffffe00
              levelTop = (c0<<9) + i*((c1-c0)<<4)   levelBot = (c2<<9) + i*((c3-c2)<<4)
```

Each source column's32 rows are resampled into H destination rows.
Installed H values span 1..159, not a format limit. Flat cells use H 32 and
srcStep 0x10000. The shading ramp runs corner-to-corner over the destination
span.

**The edge walk.** Both edges are rasterised from a startup-built step table at `L10254`
(`.bss`; built by `R0788`, 128 rows × 128 bytes, row = `|Δy|`):

```
for d = 1..127:  s = (32<<16) / (d+1)                 # unsigned
                 acc = 0x8000
                 for k = 0..d:  acc += s;  T[d][k] = acc >> 16     # T[d][d] == 32 = sentinel

offset(i) = #{ k < d : T[d][k]    <= i }   # forward:  top edge if yTL<yTR, bottom if yBL>=yBR
offset(i) = #{ k < d : 32-T[d][k] <= i }   # mirrored: the other two cases
                                            # each edge latches once its index reaches d
closed form for d in [1,127]:
offset(i) = min(d, floor((2i+1)(d+1)/64))   # exact except d = 63 and d = 127
```

The convention is **inclusive at both ends** (`d+1` accumulator steps across 32 columns, sampled
at pixel centres) — *not* `d` steps. For `d >= 63` the drawn edge does not reach the far vertex by
column 31.

**Tiling, seams and degeneracies** (`TERR-GEOM-036` (amended, superseded)):

- A cell's bottom edge and the next row's top edge use opposite
  quantizations. They agree except at `|Δy|=63` or 127, where the upper cell
  paints one extra row. Ascending row order lets the lower cell repaint it.
- Columns with `H<=0` are dropped because their top and bottom edges cross.
  Neighboring cells share those edges and paint the crossed range; the
  dropped column's source pixels are lost.
- The step table has 128 rows. Adjacent signed heights could differ by 255;
  behavior for `|Δh|>=128` remains Unknown. Installed differences are<=127.
- Horizontal clipping drops a cell unless its complete 32-pixel width is
  inside the clip rectangle. Vertical clipping operates per row.
- Hit testing uses `y0+((y1-y0)*(x&31))/32`, which can differ from the drawn
  table walk by up to 3 rows. Rendering uses the table walk.

## Far-edge cells (`TERR-EDGE-024…TERR-EDGE-026`)

Render grids are unpadded W×H arrays: type 1 uses 2WH bytes; the byte planes
use WH. A cell samples four vertex corners, so the last row and column have
no in-grid +1 vertex. The renderer does not clamp, duplicate, wrap or use a
side table for those reads:

- **Corner sourcing = raw flat row-major addressing.** The four renderers
  (`R0539/R1814/R1815/R1816`) read `grid[idx], grid[idx+1], grid[idx+W], grid[idx+W+1]`
  (`idx = col + row·W`) and pass the values to the blitters. So a **last-column** cell's `+1` corner
  reads `grid[(row+1)·W + 0]` — the **next row's column 0** (a row-major wrap, in-bounds except at the
  last row); a **last-row** cell's `+W` corner reads `grid[H·W + col]` — **past the `W×H` allocation**.
  The same flat scheme drives heights (projection `R1818`) and tile words (object passes).
- **The outer ring is drawn, not skipped.** No renderer special-cases it; the mesh is built (plus a
  several-cell over-scan margin beyond the grid) and blits are issued, gated only by screen culling.
  The one bounds check in the whole draw is a single `if (worldRow < H)` guard on the smoothed-height
  pass — its presence proves the draw reaches/exceeds the grid edge.
- **But the outer band's shading is degenerate.** The per-vertex brightness outer ring is never
  computed (interior `{1..W-2}×{1..H-2}` only) and the allocator does not zero it, so the far corners
  of the outermost cells are uninitialised / wrapped / out-of-bounds. The engine tolerates this
  because gameplay is bounded by the **derived 8-cell sim border** (`R0470`, `ALM-TERR-016`) and
  the camera keeps that band at the extreme edge.
- **Reimplementation guidance:** the far edge has **no defined shading** to reproduce — a faithful
  port should **clamp-to-edge** (safe, defined) and treat the outer band as the non-gameplay border it
  is. See `claims/terrain.md` `TERR-EDGE-024…026`.
