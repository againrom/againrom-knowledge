# Structure placement

[Reference](format.md)

## Structure placement — where a building stands (`TERR-STRUCT-100…TERR-STRUCT-107`)

A `structures.reg` placement is **not** the unit/object shape above. It has no canvas and no
anchor pixel; its anchor is a **cell** and every frame of its sheet is exactly one tile
(`TERR-STRUCT-100`, `SPR256-STR-040`).

```
anchor cell (ac, ar) = (fine.x >> 8, fine.y >> 8)      installed fine = cellByte*256 + 128
                                                        i.e. the footprint's min col / min row

drawn once per footprint cell (c, r), c in [ac, ac+TileWidth), r in [ar, ar+TileHeight):
    COL0   = c - ac                                     local column
    ROW0   = r - ar                                     local row
    rowTop = ROW0 - TileHeight + FullHeight
    limit  = (ROW0 != 0) ? rowTop : 0                   the BACK row also draws the overhang

    for k = rowTop down to limit:
        frame = block(k*TileWidth + COL0)               see SPR256-STR-041
        dstX  = (c - scrollX) * 32
        dstY  = (r - scrollY) * 32 - lift - (k_first - k)*32       ( - drawable+0x10; 0 at selected creation )
```

`lift` is the drawable's shared `+0x68`, sampled at the footprint centre
`(TileWidth<<7)+fineX-128`, `(TileHeight<<7)+fineY-128`. The containing
four signed heights feed three interpolations with truncation toward zero.
The unconditional structure cell-centre/mean consequence of TERR-STRUCT-106
is partially retracted: footprint parity changes the weights, and nested
truncation need not equal a four-corner mean. The named arithmetic stands.
`dst` is the frame's top-left. — TERR-221

So a `TileWidth × TileHeight` building draws `TileWidth × FullHeight` tiles: the bottom
`TileHeight` rows sit on the footprint, the top `FullHeight − TileHeight` rows hang above the
back row.

The mission 90 Castle is 11x5 cells with `FullHeight` 5 at columns
104..114, rows 9..13 of a 144x144 map. The placement clause of TERR-STRUCT-207
stands; its 4 px top-margin projection is partially retracted. The shared
centre `(109.5,11.5)` has heights 28,28,28,44 and lift 32. At the selected
creation offset 0, the first strip is at world Y256, exactly the highest
view's top. Computed margin: 0 px. Native current offset and complete frame
visibility remain Unknown. — TERR-221

**Shadow** (`TERR-STRUCT-103`, `TERR-SHDW-136`): the same loop and the same `dstY`, through the
sprite's `vt+0x3c` with `[L04368]`, and `dstX` displaced **per strip** by
`ftol( tan(sunAngle) · ((FullHeight − k)·32 − ShadowY) )`. That is not the only shear: the call
also passes the **live 16.16 slope** into blit argument 5 (`L10836`/`L10813`), so each strip's
rows are sheared again inside the blit. The two compose into one pivot — `X = col*32` at the screen
row `dstY_k + h` of the strip with `(FullHeight − k)·32 = ShadowY`, `h` being the strip frame's
height. Wooden bridges override the shadow with a `RET`: they cast none.

**`ShadowY` is also a suppression sentinel** (`TERR-SHDW-136`). It is read by one instruction,
`L10842`, a signed 32-bit integer read of the object's `+0x30` field, and no compare in the routine or the loader touches it.
Over the 66 shipped classes it is 28…55 on 51 of them and then `10000` on four and `20000` on
eleven, against a maximum `FullHeight·32` of 192 px. Those fifteen push the strip 660–11 400 px
sideways, past any viewport, so the blitter rejects it: **they cast no shadow, and the mechanism is
displacement rather than a branch.** They are the flat and the hollow — four bridges, four graves,
a cave, three wells, a teleport, a campfire, `magic`. A consumer must not clamp the field.

**Order** (`TERR-STRUCT-104`, amended): the drawable registers its prepared
footprint into `CMapView+0x94`. The early body phase takes `Flat!=0`; the main
cell phase takes `Flat==0` before the selector-2 CUnit calls in that cell.
Both body gates require signed `+0x78<2`. The former main-body `==0` clause is
retracted: that later test guards work after the body call. Structure shadows
have an earlier separate phase (`ANIM-DRAWGATE-087`, `ANIM-CELL-085`).

**Ownership never reaches the sprite** — only the minimap blip `vt+0x34` (`TERR-STRUCT-107`).

## Destroyed structures

The client draws a structure's ruin grid when its published health is not positive and its class is not
Indestructible. Published health is Building word `+0x42`, or 1000 when `+0x44` is zero
(`TERR-STRUCT-210`).

| Draw when HP is not positive | Classes |
|---|---:|
| ruin grid | 33, including the two Switches, whose ruin grid is their toggled state |
| intact art again: Cave (14), Inn 3 (45) | 2 |
| no change: maximum HP zero | 26 |
| no change: Indestructible with positive maximum HP (31 classes are Indestructible in all) | 3 |
| separate bridge selectors | 2 |

A placement word of zero in a type-4 record starts the Building at HP 0; its draw of the ruin grid at mission
start is High / Unknown, because mission-start publication is not read. 39 of 950 shipped Building placements, on 7 maps, carry it, all on ruin-grid classes
(`TERR-STRUCT-211`). In 339 distinct files, one structure other than the two Switches, a Bee House of
mission 131, reached HP 0 in play; the 994 dwelling and 100 Cave records are all at maximum HP
(`TERR-STRUCT-212`). The saved record and its footprint persist (`SAV-1163`, `SAV-1164`).

## Structure light (`TERR-LIGHT-285`…`TERR-LIGHT-290`)

Five shipped classes are light sources: Tower 1 and Tower 2 (IDs 12, 13), Well 2 (15), Well 3 (16)
and the Campfire (53), all with `LightRadius` 1 (`REG-125`). The structure method `vt+0x40`
stamps their light into the vertex stamp grid `CMapView+0xa8` each time the per-frame light grids
are rebuilt, for every structure in the view's object map, on screen or not. Unit vtables have an
empty `vt+0x40`; no unit is a light source this way (`TERR-LIGHT-290`).

```
gate:   Lighting option on  AND  LightRadius != 0  AND  health word (signed) > 0
rect:   x1 = anchorCol - scrollX, x2 = x1 + TileWidth - 1      (rows alike)
        clamped to x1 >= -3, x2 <= visCols+2, y1 >= -3, y2 <= visRows+6
level:  P = LightPulse, t = view tick counter (+1 per game tick, 0 at view construction)
        P == 0            -> 32
        t mod 2P <  P     -> (t mod P) - P/2 + 32
        otherwise         -> P/2 - (t mod P) + 32
stamp:  for col in x1+R-1 .. x2+R, row in y1+R-1 .. y2+R  (R = LightRadius)
            if -3 <= col < visCols+3 and -3 <= row < visRows+3:   # stamp grid, not the visible window
                a8[(visCols+7)*(row+3) + col+3] = level        # overwrite
        writeCount += 1; dirty |= ((x1-2)*32, (y1-2)*32, (x2+3)*32, (y2+3)*32)
```

With R = 1 the stamp covers the footprint's `(TileWidth+1) × (TileHeight+1)` corner vertices: 4
for the Campfire, 9 for the Towers and Well 2, 16 for Well 3 (`TERR-LIGHT-286`). A larger R moves
the block by R-1 vertices right and down without widening it. The Animation option, the ambient,
fog and explored state are not read. A structure whose health word is not positive stamps
nothing (`TERR-LIGHT-285`). That a destroyed Tower reaches that word, and so also draws its ruin
grid, is Medium (`TERR-STRUCT-210`); later writers of the word are not censused.

| Class | stamp range | period (ticks) | period at 16 ticks/s |
|---|---|---:|---:|
| Tower 1, Tower 2 | 28 … 36 | 16 | 1.0 s |
| Campfire | 29 … 35 | 12 | 0.75 s |
| Well 3 | 30 … 34 | 8 | 0.5 s |
| Well 2 | 32 | — | — |

Lower is brighter. Every structure reads the same counter, so lights of one class pulse in phase;
there is no random draw (`TERR-LIGHT-287`).

A rebuild runs the structure walk, then the projectile walk, then a third point stamp (its
source is held Unknown by `MAGIC-269`), then the merge routine; every stamp is a plain byte store,
so the last write to a vertex stands (High, `TERR-LIGHT-288`). That every projectile stamp follows
every structure stamp is Medium: the classes entered into the structure walk are not censused.
The merge routine's own stamps come last (`MAGIC-UNITLIGHT-057`, partially retracted for its
square-footprint and stamp-gate clauses, which this page does not use; `MAGIC-270`). The ground
takes `min(stamp, terrain byte)` at a stamped corner (`MAGIC-273`, partially retracted for its
hardware-mode label only). By the sun model flat daytime
ground is level 46, gain 1.5625 (`TERR-LIGHT-013`, `TERR-LIGHT-020`), so a stamp of 28…36 gives
gain 2.125…1.875 there (Medium: other times of day give other flat levels). By the model of the
merge arithmetic (Medium, `TERR-LIGHT-288`), a sprite cell with four stamped corners gets level 0
(1 at stamp 36) against 3 by day and 8 at night unlit.

The sprite frame steps a separate per-object counter once per tick. The Tower and Campfire frame
periods equal their light periods; Well 3's 12-tick timeline runs against an 8-tick light, and
the relative phase depends on when the structure was constructed (`TERR-LIGHT-289`).
