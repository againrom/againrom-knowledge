# Sprite placement and composition

[Reference](format.md)

## Unit health and mana bars

The original 16-bit CUnit/CAirUnit status-bar path reads signed current HP
H and maximum HP M. For positive M its colour family is red when
H < floor(M/4), yellow when H < floor(M/2), otherwise green. Equality
advances to the next family. These comparisons truncate the maximum
before comparison; M=101 changes colour at H=25 and H=50.
— TERR-224

Four filled rows run from Y-2 through Y+1. Their source RGB components
before surface packing are:

| Family | Row 0 | Row 1 | Row 2 | Row 3 |
|---|---|---|---|---|
| HP red | (128,0,0) | (255,0,0) | (192,0,0) | (128,0,0) |
| HP yellow | (128,128,0) | (255,255,0) | (192,192,0) | (128,128,0) |
| HP green | (0,128,0) | (0,255,0) | (0,192,0) | (0,128,0) |
| Mana blue | (0,0,128) | (0,0,255) | (0,0,192) | (0,0,128) |

Each component c is packed by `(c >> (8-bits)) << shift`; combine
channels with OR. These values are generated RGB, not palette indices.
Mana has no colour threshold in this arm; maximum mana <= 0 skips its
bar, and its Y is health Y+4. — TERR-224, TERR-226

The ordinary unfilled rows use grey RGB intensities `[64,128,96,64]`.
If show-health is on and the unit is unselected, the remainder is untouched.
Only filled pixels blend: for packed destination B and fill C, output is
`((B >> 1) & mask) + ((C >> 1) & mask)`; RGB565 uses mask 0x7bef and
RGB555 uses 0x3def. A selected unit keeps the ordinary grey remainder
and full fill colours in either show-health state. — TERR-225

The map painter admits selected objects in all three grids; its +0x8c
and +0x90 grids also admit unselected objects when show-health is on and
object+0x78 is zero. The earlier all-grids selected-only clause of
TERR-SPR-048 is partially retracted. — TERR-225

Inner width W is `class.Right-class.Left-8`, using class+0x8c/+0x84.
Fill starts at L+4 inside `[L+4,R-4)`. Length is the signed quotient
`trunc(W*current/maximum)` after a 32-bit multiplication. A zero HP
quotient becomes one pixel only when current HP is nonzero; a zero mana
quotient always becomes one pixel, including current mana zero.
There is no clamp to W. The pixel writers separately clip to the viewport.
— TERR-227

These contracts are High for the named 16-bit interior path. Native format
choice, cap pixels, other bar paths and native admission of malformed HP
maxima or overflowing products remain Unknown. — TERR-224, TERR-225,
TERR-226, TERR-227

## Sprite placement — where a unit or object stands (`TERR-SPR-038…TERR-SPR-041`)

Sprite placement applies one vertical terrain lift to the frame; sprite
blitters do not use the sloped-terrain step table. The draw call carries
scalar placement rather than four corner heights. TERR-SPR-038's fifth-
argument label is amended: this argument is the shadow's per-row X slope,
not a brightness level. — TERR-SPR-038, TERR-SPR-066

```
alt(col,row) = ( h(wc,wr) + h(wc+1,wr) + h(wc,wr+1) + h(wc+1,wr+1) ) / 4     trunc toward zero
               four signed-byte reads of the raw type2 grid; wc = col+scrollX, wr = row+scrollY
               built by R1818 pass 3 into CMapView+0xc0        (TERR-SPR-039)

anchorX = (CenterX - Width /2) + frameWidth /2      class +0x20, +0x18   (TERR-SPR-040)
anchorY = (CenterY - Height/2) + frameHeight/2      class +0x24, +0x1c
    frameWidth/frameHeight are the BODY pass's own drawn frame; the SHADOW pass measures
    frame 0 instead, and frames of one .256 need not share a size  (TERR-SPR-043)

dstX = col*32 + 16 - anchorX          ( - T on the SHADOW pass only, see below )
dstY = row*32 + 16 - anchorY - alt

dst is the frame's TOP-LEFT. With frameW=Width and frameH=Height this reduces to:
    frame pixel (CenterX, CenterY)  lands on  (col*32 + 16, row*32 + 16 - alt)
```

**The shadow pass's extra `dstX` term is a pixel count, not the slope.** Two sun-derived values
travel together and they are not interchangeable (`TERR-SHDW-130`, and `TERR-SPR-067`'s naming of
one as the other is amended in `retracted.md`):

```
s = ftol( tan(theta) * 65536.0 )                         the 16.16 per-row X slope; blit arg 5
T = ftol( tan(theta) * (floor(frameH/2) + floor(Height/2) - CenterY) )      a PIXEL count
      the multiplicand is frameH - anchorY for even frameH, one less for odd:
      the two /2 are separate truncations
```

Only `T` enters `dstX`. `s` goes to the blitter, which shears about the **bottom edge of the blit
rectangle** all by itself; subtracting `T` moves that pivot up to the **anchor row**, so that the
composed placement of image row `r` is

```
X(r) = (col*32 + 16 - anchorX) + tan(theta) * (anchorY - r)        residue <= 2 px
```

A structure writes the same rule upward instead of downward and therefore **adds** where a sprite
subtracts — `dstX = col*32 + ftol(tan(theta)*((FullHeight-k)*32 - ShadowY))` — because
`(FullHeight-k)*32` is a height above the image bottom while `anchorY` is a depth below the frame
top. The two signs are both correct and neither may be copied onto the other path
(`TERR-SHDW-131`).

### Which silhouettes a unit's shadow draws, and through which arm (`TERR-191`, `TERR-192`, `TERR-193`, `TERR-194`, `TERR-195`, `TERR-196`)

One routine, `R0553`, is `vt+0x2c` of both `CUnit` and `CAirUnit`. It draws one pair of
sheets, each as a silhouette (`vt+0x3c` sheared, or `vt+0x1c` flat), at one frame index:

```
pair   = hero pair, +0x194 and +0x198        if drawable+0x18c bit 0 is set
         main pair, class sheets +0x04, +0x08   otherwise                  (TERR-192)
key    = the unit's effect list holds kind 0x26 (invisibility)
first  : key absent                          -> shroud cell [L04368]
         key present, owner row bit 3 clear  -> the routine draws nothing
         key present, owner row bit 3 set    -> shroud cell [L04335]
second : only if key absent and the Smoothing option cell [L04367] != 0,
         shroud cell [L04335]                                          (TERR-191)
arm    = flat (vt+0x1c)    for an instance of CAirUnit, both main-pair silhouettes
         sheared (vt+0x3c) for every other unit and for the hero pair      (TERR-193)
```

The levels of the two cells by day band are `TERR-LIGHT-126`'s (amended): the second silhouette and an
invisible owner's shadow are drawn at half the level of an ordinary first silhouette. The owner row
bit marks the local player's own entry, so an invisible unit's shadow appears on its owner's client
only. The hero pair has no flat arm. Both silhouettes recolour the pixels already on screen through two word writes in each of the eight cores,
flat and sheared, with no coverage buffer, so a pixel stamped by both is recoloured twice (`TERR-196`); the shipped
sheared pairs (hero pairs and the main pairs of non-air units) do not overlap at equal frame size, and overlap at
unequal frame size once the shear is non-zero (`TERR-192`). The flat-arm pairs of the two air-unit sheets (`Sonic Bat`,
`Dragon`) translate each sheet by its own size and do not overlap in any of their 273 frames (`TERR-195`). Of the 68
displacement stores through `+0x18c`, 45 reach a drawable (receiver classes assigned by hand, Medium) and 5 of those set bit 0 (`TERR-194`).

**Two altitude models per frame.** The terrain raster uses `+0xb4` — one *corner*, `r*32 − h`
(`TERR-GEOM-031`). Everything standing on it uses `+0xc0` — the *mean of four corners*. A port must
implement both; using the terrain mesh to lift a unit puts it at a corner's height instead of the
quad centre's.

The four-corner mean equals a corner height on flat cells and can also
equal it on sloped cells. Flatness therefore cannot be tested from that
equality. The lift uses the four named samples and truncation toward zero.
Installed contact points can lie up to two pixels below their own terrain
span at the cell center, within pixels drawn by neighboring cells.
— TERR-SPR-039

The occlusion test `R1401(classId, col, row, alt)` recomputes the same `dstY`, converts the
sprite's top and bottom back to terrain rows with the picker `R0377`, and skips the draw iff
every covered cell's shroud level is `0x10` (fully dark).

### Which frame a static object draws (`TERR-SPR-042` (amended, superseded), `TERR-SPR-043`, `TERR-TILE-044` (amended, partially retracted))

The `frame` argument of the object draw call is not simply the class's `Index`. `R0379`'s
type-3 pass writes it on three arms and then overrides it once (`TERR-SPR-042` (amended, superseded)):

```
c    = type3[row*W + col]                     0 -> nothing here
k    = classes[c - 1]                         objects.reg section index   (ALM-CLS-035)
imp  = tile[row*W + col] & 0x2000             bit 13 -- see below
anim = ( (tile[i] | tile[i+1] | tile[i+W] | tile[i+W+1]) & 0xc000 ) == 0xc000

if k.timelineLen != 0 and anim and not (imp and k.DeadObject != -1):
      phase = (animCtr + col*(row+1)) % k.timelineLen        # SIGNED remainder
      frame = k.Index + k.timeline[phase]                    # timeline = class+0x2c
elif imp and k.DeadObject != -1:
      k     = classes[k.DeadObject]           # the whole sprite is swapped
      frame = 0
else: frame = k.Index

if animations_disabled:  frame = 0            # [L05659] == 0, -noanimation / -detail0
sprite = Files[k.File]                        # graphics.res!objects/<path>.256
```

`k.timeline` is not a registry key: the loader run-length expands `AnimationFrame[i]`
`AnimationTime[i]` times into `class+0x2c` and stores the length at `class+0x3c`
(`REG-OBJ-046`). `animCtr` is `CMapView+0xa70`, the same counter the water phase uses,
unshifted here.

**The two flag bits of the tile word** (`TERR-TILE-044` (amended, partially retracted)). Bit 13 is the one `TERR-DIRT-017`
already uses for the dirt composite; the object pass reads the same bit for the `DeadObject`
swap, and `R0483` reads it to decide whether a cell still holds a standing destructible.
It is **rewritten at runtime** by an RLE decoder (`R1656`) over the interior from `(8,8)`,
so the `.alm` field supplies only its initial state.

See [fog and visibility](fog.md) for tile bits 14/15, and
[sprite composition](composition.md) for frame passes and unit shadows.

### Sprite lighting — how bright a unit or object is drawn (`TERR-LIGHT-059`…`TERR-LIGHT-064` (amended, partially retracted))

A sprite **is** lit, per sprite rather than per pixel, through the same builder and the same kind of
level-indexed LUT the terrain uses. The `.256` class dispatches on vtable `L04339` (stored by its
own payload constructor `R0753`) and exposes four blit entry points, which split in two:

| slot | routine | pixel loops | reads | role |
|---|---|---|---|---|
| `vt+0x14` | `R1791` | `R0796`, `R0895` | the **source** index | body, plain |
| `vt+0x34` | `R1547` | `R0789`, `R1793` | source + destination + `[L07813]` | body, `spritesb` overlay |
| `vt+0x1c` | `R0795` | `L03664/e240/eef0/f140` | the **destination** only | silhouette (shadow) |
| `vt+0x3c` | `R1831` | `R1832/e750/f380/f690` | the **destination** only | silhouette (shadow), sheared |

The lit pair takes `(dstX, dstY, frame, level, shadeObj, mirror)` and forms

```
row = shadeObj[+8] + (level << 9)          # 512-B row stride, 256 u16 entries
dst = row[srcIndex]                        # L10374 reads the source byte; L10099 reads the 16-bit entry at 2 * srcIndex of the row
```

The silhouette pair takes no shading object: it advances the source without reading it, reads the
destination pixel, `SHR 3`, and looks *that* up in the shroud table `[L03341]` at row
`arg4 × [L03340]` — so those two passes recolour what is already on screen under the sprite's
outline. `vt+0x3c`'s fifth argument is the shadow's **16.16 per-row X slope** from the sun angle, not
a brightness (`TERR-SPR-066`, correcting `TERR-SPR-038`); `vt+0x1c`'s caller instead offsets `dstX`
by `shear/2000`.

**Which pass is which.** An object cell issues `vt+0x3c` twice (shadow) then `vt+0x14` + `vt+0x34`
(body). A unit is dispatched twice over the drawable grid `CMapView+0x90`, both times at
`idx = (row+3)*(visCols+6) + (col+3)` under `drawable+0x78 == 0`: `vt+0x2c` = `R0553` gets the
`+0xc0` altitude and draws the **shadow**, `vt+0x28` = `R0552` gets the `+0xb0` light level and
draws the **body** (`TERR-SPR-065`, correcting `TERR-SPR-048`, partially retracted).

**The level.** `CMapView+0xb0` is a per-frame byte grid of `(visCols+6) × (visRows+10)`, filled by
`R0608`:

```
memset(grid, [L05650] >> 2, size)                 # ambient 0x0e by day -> 3, everywhere
for each entry of the lit-object list this+0x9f0:
    flags & 0x1000  -> grid[cell] = 0                 # brightest
    flags & 0x20000 -> grid[cell] = 0xc               # x0.5
    flags & 0x8     -> flicker 0 / 0xc, + a radius-1 stamp into +0xa8
for every cell:                                       # +0xa8 = light-source stamps, 0xff = none
    q[k] = max(0, a8corner[k] - 0x20)  (0xdf -> ambient)
    if all four corners were 0xff: keep the value above
    else grid[cell] = min( (q0+q1+q2+q3) >> 4 , ambient >> 2 )
```

Lower level means brighter. The merge caps its result at the ambient level.
The sole stamp-writer and lit-object-only clauses of TERR-LIGHT-061 are
partially retracted: Lightning and Prismatic Spray directly overwrite four
vertices per admitted drawn-path cell with `u8(10 * phase)`. Fire Arrow,
Fire Ball flight and its explosion also deposit point stamps through the
helper; its footprint follows a clipped squared-distance test rather than
a radius square. — TERR-LIGHT-061, MAGIC-270, MAGIC-271

Dynamic lighting gates point stamps. It does not gate the direct bolt
stores or the named unit-grid merge. With `L10964 & 2` clear, Dynamic
lighting gates the terrain-light pass. With that bit clear and Dynamic
lighting off, bolt light reaches units but not the ground. The
`L10964 & 2` set branch passes the stamp through the inspected terrain
dispatch/call paths without a local Lighting test. It uses the stamp
directly, falling back to the terrain byte for an unstamped corner. The
local branch and fallback are High. Native activation and the visible
result of the bit-set branch are Unknown: TERR-FAMILY-187 found no
enabling writer in its bounded address-form search, and startup writes 0.
Ordinary consumers share `+0xb0`; special unit passes can force zero or
interpret the argument
differently. The same-grid equality clause of TERR-LIGHT-061 is therefore
narrowed to its named reads. Native pixels and outer option/frame gates
remain Unknown.
— MAGIC-273 (hardware-mode label partially retracted), TERR-LIGHT-061,
TERR-FAMILY-187

**Mode 2 — the table.** `R1107`'s arm at `L10378` (the dispatch table sits at
`L05822`), built as `(nLevels = 0x10, mode = 2, useTint = 1)`:

```
out_chan = clamp( ((palette_chan + skyTint_chan) * (nLevels - level) * 2) / nLevels, 0, 255 )
```

so **row `L` has gain `2(nLevels − L)/nLevels`**: row 0 → ×2.0, row `nLevels/2` → ×1.0 (identical to
no table at all), row 15 → ×0.125.

**The two ladders are one ladder.** `2(16 − S)/16 = (96 − (4S + 32))/32` exactly, and on the same
palette the rows are bit-identical (0 of 256 entries differ at `L = 32, 36 … 64`). Shipped daytime:
ground `L = 46` → ×1.5625, sprite `S = 3` → ×1.6250 — half a step apart because `ambient/4 = 3.5`
truncates to 3. Both indices descend from the same byte the global at `L05650`, whose only two readers in the
image are `R0608` and `R0468`.

**Which table each drawable gets.** Terrain: one, from the tile palette. Each object sheet: its own
`(0x10, 2, 1)` table from its own BMP palette (`R2112`), passed as `spriteObj + 0x14`. A unit:
selected on the `units.reg` `Palette` key at `class+0x98`: `0` selects one of
16 shared owner tables from `human.pal`; `1` selects `class+0x9c[0]`; `> 1`
selects `class+0x9c[face-1]`, where face is the actor's Data.bin tier. The last
arm is not owner-indexed (`PAL-KEY-002`, `PAL-OWN-007`, correcting that clause
of `TERR-LIGHT-064` (amended, partially retracted)). The owner-index producer and first-message lifetime are
specified by `PAL-RULE-021` and `PAL-FIRST-022`; see [palette selection](../pal/format.md).
`MAGIC-STONEDRAW-084` identifies both greyscale overrides as the stone-curse
draw arm; its stated confidence and the `drawable+0x15a` Unknown remain.
The terrain.3d tiles share one palette; sprite ramps are built from each
sprite's own palette.

**Reimplementation guidance.** Light every sprite through its own palette's 16-row mode-2 table at the
cell's `+0xb0` level. Drawing sprites unlit is not a small error: the shipped swordsman frame means
104.9 luma at the engine's row 3 and 65.2 with no table, against 106.1–118.3 for the terrain in the
same crop (`tools/terrgfx -mode sprlightart`).

### Terrain output range (`TERR-LIGHT-063`)

For the shared terrain.3d palette and installed per-vertex levels 30..70 at
θ=0.78539815, mode 3 gives the following output ranges:

| level | gain | entries with a clipped channel | entries packing to `0xffff` | pixels clipped | pixels pure white | mean luma |
|---:|---:|---:|---:|---:|---:|---:|
| 30 (brightest shipped) | ×2.0625 | 89 | 33 | 16.99 % | 5.09 % | 153.25 |
| 46 (flat, daytime) | ×1.5625 | 47 | 11 | 10.59 % | 0.10 % | 119.49 |
| 64 (unattenuated) | ×1.0000 | 8 | 1 | 0.30 % | 0.00 % | 79.48 |
| 70 (darkest shipped) | ×0.8125 | 0 | 0 | 0 % | 0 % | 64.14 |

**The mapping saturates**: bright shipped terrain reaching pure white is what the engine does, on
about a twentieth of all tile pixels at the brightest level a shipped map produces. Pixel weights are
over 651 266 `terrain.3d` pixels; the level range carries its θ, framing base and window because the
ceiling moves with all three (`TERR-LIGHT-029`).
