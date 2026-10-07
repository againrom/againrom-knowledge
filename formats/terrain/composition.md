# Sprite composition and shadows

[Reference](format.md)

## The order one frame is painted in, and everything that bounds a drawn silhouette

`TERR-SPR-137…TERR-SPR-141` (the 137–139 count, population and gate clauses are amended),
`TERR-FOG-142`, `TERR-LIGHT-143`, `TERR-SPR-144`.

`R0379` has nine cell sweeps and one non-cell collection walk. The retained
phase labels 1..10 count the collection as phase 6. The prefix's terrain/rect work
precedes these loops; the lowest direct back-edge target is `L10857`. The old
ten-cell-sweeps count is retracted in `TERR-SPR-137` and `ANIM-WALKORDER-088`.

Registration selectors are joined to their actual destinations
(`ANIM-REGISTER-083`, `ANIM-CATEGORY-084`):

| Registration selector | View plane | Conditional population |
|---|---|---|
| 0 | `+0x98` | CBackPack |
| 1 | `+0x94` | CStructure and bridge subclasses |
| 2 | `+0x8c` | CUnit with unsigned `+0x15a<2` and `(+0x18c&0x80)==0` |
| 3 | `+0x90` | CAirUnit, without those two tests |
| 4 | `+0x9c` | The other CUnit branch |

Selector 1 stores the prepared, clipped footprint rectangle; the other selectors
store one computed anchor. Each plane/cell holds one pointer. Native refresh
order, the final writer after collisions and stale-entry lifetime remain Unknown.
The index is `(row+3)*(visCols+6)+col+3`; the registration gate is `+0x4c!=0`.
The out-of-table selector arm uses its argument as a pointer, not a validated
default plane (`ANIM-REGISTER-083`).

| Phase | Dispatch composition, subject to gates |
|---|---|
| 1 | Structure shadows |
| 2 | Flat structure bodies |
| 3 | Per cell: selector 4 shadow/body, then CBackPack shadow/body |
| 4 | Per cell: non-flat structure body, selector 2 shadow/body, auxiliary `vt+0x18` dispatch, static-object path, then retained-area selector 1 |
| 5 | Complete selector 3 shadow sweep |
| 6 | Non-cell collection payload body calls |
| 7 | Retained-area selector 0 |
| 8 | Complete selector 3 body sweep |
| 9 | Structure/selector 2/selector 3 marker and bar calls |
| 10 | Shroud |

This is the amended form of the former universal unit wording in `TERR-STRUCT-104`,
`TERR-SPR-065` and `TERR-SPR-139`. CUnit and CAirUnit share drawing methods but
not these registration/phase populations (`ANIM-CELL-085`, `ANIM-AIRPASS-086`).
The retained-area selector passed to the overlay routine is a separate domain
from a drawable registration selector.

Both structure body phases accept signed `+0x78<2`; Flat chooses the phase.
Selector 2/3/4 and CBackPack body paths require `+0x78==0`. Selector 2 has the
additional key `0x26` lookup and indexed-mask gate. Shadow calls have their
separate global enable gate. These are corrected local dispatch rules, not
proof that a sprite body paints opaque pixels (`ANIM-DRAWGATE-087`).

Phases 1..5,7..8 walk rows -4..visRows+7 ascending and columns visCols+3..-4
descending. Phase 9 instead walks rows -3..visRows+6 and columns visCols+2..-3.
Phase 10 is column-major, ascending columns below visCols and rows below
visRows+4. The phase-9 common-bound wording in `TERR-SPR-138` is retracted.
Within a drawable cell sweep, later cells have greater row or equal row and
smaller column. Across phases, the earlier phase finishes first regardless of
cell coordinates (`ANIM-WALKORDER-088`). The late shadow/body phases have one
collection walk and one cell sweep between them; identical entry tests do not
establish identical painted pixels (`TERR-SPR-139`, amended).

**What bounds a silhouette.** One global device-space clip rectangle,
`L01503` / `0x440c` / `0x4410` / `0x4414` = left/top/right/bottom, set from `*(CMapView+0xf4)`
at each of the four block boundaries by `R0373`, which **replaces** — there is no
`IntersectRect`, no clamp against the surface, and no per-cell destination window anywhere on the
draw path. `R1832`, the sheared silhouette blit
`(dstX, dstY, w, h, src, level, shear)`, tests the box `[dstX - span, dstX + w + span] ×
[dstY, dstY + h]` with `span = |(shear*h) >> 16|`; entirely inside takes an unclipped loop,
otherwise it tries a whole-sprite reject on the same four edges and then clips **per row** (against
top/bottom) and **per pixel** (the sheared X against left/right). That is the opposite of the
terrain blitters, which drop a cell whole when its 32-px X range is not entirely inside
(`TERR-SPR-140`).

**Fog does not clip a shadow.** The only per-draw shroud test in the image is `R1401`, and
it has **exactly one caller** — sweep 4's inlined object draw, where a true verdict skips the
object's shadow *and* body on one boolean. No unit shadow is subject to it. A silhouette crossing
fog is therefore drawn whole and then overpainted by sweep 10, so its visible edge is the shroud's
**per-vertex gradient**, never a cell boundary; on a fully visible cell the shroud draws nothing
and the shadow stands as drawn (`TERR-SPR-141`, `TERR-FOG-142`, `TERR-FOG-083`).

**Overlapping shadows compound.** The blit writes `table[level][dst >> 3]` back over the
destination with no guard against a second application and no per-pixel state that could hold one,
so where two silhouettes cross the ground is `((16-L)/16)²` rather than `(16-L)/16`
(`TERR-LIGHT-143`, on `TERR-FOG-084`'s law). A consumer that unions shadows into a coverage mask
and darkens once draws something the original never draws.

**Customisation (G2).** `[L06416]` is the shadow switch: seven reads image-wide, all inside
`R0379`, each gating exactly one shadow draw and nothing else, so clearing it removes every
silhouette and changes no body, bar, terrain cell or shroud cell. It is a runtime global — no
shipped file's bytes move (`TERR-SPR-144`).

**Open:** the runtime value of `CMapView+0xf4`, the complete producer/lifecycle
join of `drawable+0x78`, native registration/collision order and actual pixel overlap. The collection's complete insertion/type
population remains unjoined; its projectile interpretation remains Medium.

## A unit places itself — and its shadow places itself differently (`TERR-SPR-047`, `TERR-SPR-065…TERR-SPR-067`, late-caller population amended)

The shared CUnit/CAirUnit drawing methods place their own image. The late
`CMapView+0x90` callers below are CAirUnit's selector-3 path; ordinary and
alternate CUnit reach the same methods in earlier cell compositions
(`TERR-SPR-065`, amended; `ANIM-CATEGORY-084`, `ANIM-AIRPASS-086`). Both late
paths require `drawable+0x78==0`, and both methods ignore the `(col,row,…)`
destination arguments:

```
vt+0x2c  R0553   the SHADOW   third argument = the CMapView+0xc0 four-corner altitude
vt+0x28  R0552   the BODY     third argument = the CMapView+0xb0 light level

anchorX = frameW/2 + (CenterX - Width /2)      class +0x34, +0x2c   (units.reg record, REG-UNITS-049)
anchorY = frameH/2 + (CenterY - Height/2)      class +0x38, +0x30
    frameW/frameH via vt+0x20 / vt+0x24 at the DRAWN frame, in BOTH passes

BODY     dstX = unit+0x60 - anchorX                                    (L10428..L10429)
         dstY = unit+0x64 - anchorY - unit+0x10 - unit+0x68
SHADOW   dstX = unit+0x60 - anchorX - T                                (L07949..L07950)
         dstY = unit+0x64 - anchorY            - unit+0x68
             T = ftol(tan(theta) * (floor(frameH/2)+floor(Height/2)-CenterY)),  a PIXEL
             count loaded at L07951 -- NOT the 16.16 slope, which lives one slot away
             at frame base+0x18 and is pushed as blit argument 5 (TERR-SHDW-130)
```

**The two differ by exactly `(sunShear, −unit+0x10)`.** One rule will not serve both: applied to the
body, the shadow's rule puts every unit `unit+0x10` pixels too low; applied to the shadow, the body's
rule stands the shadow on the unit's feet instead of on the ground and drops its lean.
`unit+0x68` is subtracted by all four CUnit drawing routines (body, shadow, `vt+0x38`, the HP/mana
bars) and `unit+0x10` by every one except the shadow, so `+0x68` belongs to the plane the shadow lies
in and `+0x10` lifts what is drawn *at* the unit off it (`TERR-SPR-067`; the interpretation is
Medium, the reader split is not). Neither field has a writer inside `L05122..L05123`.

Frame selection is **duplicated**: each routine has its own nine-arm switch on `unit+0x74` with its
own jump table (body `L10320`, shadow `L10316`), same arm shape, and the same arithmetic arm
for arm — so a shadow is a silhouette of the same pose. The one divergence is the states-2-and-4
default arm, which draws the class id in the shadow and the light level in the body: garbage in both,
which is the standing reason to think those two states never reach a draw.
