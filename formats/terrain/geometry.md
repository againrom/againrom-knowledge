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
frame for a several-cell over-scan margin (`TERR-EDGE-026`, whose terrain-admission
clause is partially retracted); the *derived* `+0xc0` grid (mean of
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

## Far-edge cells (`TERR-EDGE-024…TERR-EDGE-026`, the latter partially retracted)

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
- The terrain-dispatch and guard-implies-raster clauses of TERR-EDGE-026 are
  partially retracted. Projection and drawable overscan are separate from
  the four terrain routines' admitted cells. TERR-217 supplies their bounds.
- The per-vertex brightness outer ring is not computed. Raw corner reads
  at a stored outermost cell therefore have no defined shading contract;
  those cells are outside the terrain-dispatch population of a nonempty
  eight-cell camera band. — TERR-EDGE-025, TERR-217

## Camera stops and terrain edges

For map dimensions W,H and nominal cell span C,R, the cell origin is bounded
by `8..W-8-C` and `8..H-8-R`. Thus the top/left nominal first cell is 8;
the bottom/right nominal last cell is H-9/W-9. The playable rectangle is
inclusive `(8,8,W-9,H-9)`. It is a movement rectangle, not a sprite clip.
— SESS-VIEW-030, TERR-SIGHT-116

The absolute target routine is `R1677`; axis deltas use
`R1678/R1679`. The inspected original clamp bodies read dimensions,
spans and origins, without grid content. The recompute prefix
`R1280` applies the upper bound only. Crossed bands are possible on
small supplied maps; no measured installed map exercises one. — TERR-216

| Resolution | View pixels without child panel | Cell span | Maximum origin X,Y |
|---|---|---|---|
| 640x480 | 480x480 | 15x15 | W-23,H-23 |
| 800x600 | 640x576 | 20x18 | W-28,H-26 |
| 1024x768 | 864x768 | 27x24 | W-35,H-32 |

The view excludes the 160-pixel side panel and snaps its bottom to whole
32-pixel rows. Open child panels can reduce its vertical span.
— SESS-VIEW-028

All four terrain routines admit local rows `0..R+3` and columns `0..C-1`.
Rows ascend and columns descend. At the top stop, the first terrain row is
8; the top, left and right movement-border strips are not terrain-dispatched.
At the bottom stop, ordinary terrain includes border rows H-8..H-5. Their
height projection can move pixels into the view. No stored outermost cell
is terrain-dispatched when the camera band is nonempty. The full software
redraw arm enters ordinary terrain without a preceding black-clear call;
conditional color-zero grid lines are a separate overlay. Native pixel
coverage, retained surface contents and alternate rendering remain Unknown.
— TERR-217

## Art at an edge

The software painter resolves the absolute widget view, intersects its
repaint rectangle with it, and resets that rectangle to the view after a
camera change. Repaint work and child overlays can reduce the active area.
The resulting global clip is screen space; it is not the playable rectangle
or a cell window. — TERR-218

Ordinary and mirrored indexed body blitters accept art across a cell line
and clip rows/pixels to the supplied global rectangle. In 18 supplied
original-x86 cases, only the sprite/clip intersection changes and all
outside or rejected pixels remain intact. Native registration, alternate
body overlays, fog and frame selection are outside that finite population.
The sheared shadow uses its separate global clip contract.
— TERR-219, TERR-SPR-140

Terrain finishes before the drawable phases. Non-flat structures, selector-2
units, auxiliary calls, static objects and retained-area overlays are ordered
within the main cell phase; other categories use separate phases. Within
the relevant sweeps, rows ascend and columns descend. Admission and dispatch
do not by themselves guarantee an opaque pixel.
— ANIM-CELL-085, ANIM-AIRPASS-086, ANIM-DRAWGATE-087, ANIM-WALKORDER-088

Mission 90 Castle's shared lift uses footprint-centre corners 28,28,28,44,
giving 32. At the selected create-time vertical offset 0, its first strip
and the highest view both start at world Y256: computed margin 0 px for all
three resolutions. Native current offset and the complete captured frame
remain Unknown. — TERR-221

## First-row populations

Stored row 0 and first playable row 8 are distinct. Over columns 8..W-9,
the measured 38 EN/34 RU maps have three all-terrain-blocked row-8 maps in
each root and some terrain-open cells in the other 35/31. The 140.alm row
has 240 blocked/0 open cells; 90.alm has 60 blocked/68 open, including five
water cells. These counts exclude objects, stamped simulation borders,
occupants and runtime flags. Every measured altitude is signed 0..127;
all 216 map/resolution camera bands are nonempty. Native blocked-versus-open
frame differences remain Unknown. — TERR-220
