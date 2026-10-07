<a id="terrain-tile-graphics-terrain3dbmp--specification"></a>

# Terrain graphics and cell rules

Terrain draws the ALM Tiles and Altitudes planes through bitmap strips.
The loader's final source table is filled by `terrain/`; selector mask
`L10964 & 0x02` optionally enables an earlier `terrain.3d/` precomposition
pass whose temporary source objects are then deleted. The two families have
equal paired geometry but different palettes and indexed pixels.
— TERR-FAMILY-187, TERR-FAMILY-188

Tile words select files and subcells; the terrain state also carries
block planes, cell records, lighting and fog. — TERR-LOC-001 (route wording
amended), TERR-LOAD-002 (order and handle search amended),
TERR-IDX-003 (amended bit interval; arithmetic retained), ALM-GRID-012, ALM-GRID-013

This reference describes decoded terrain paths, including 16-bpp lighting
and cell geometry. Full 8-bpp behavior and native activation of the optional
Direct3D branch remain incomplete. Startup writes a clear selector; no
mask-0x02 setter was found in the stated absolute-address census. Computed
writers and other native activation remain Unknown. — TERR-FAMILY-187
ALM grid payloads begin after the full 20-byte record header.
— ALM-FRAME-031 (amended; framing retained), ALM-GRID-032, TERR-GRID-027

## Input planes and draw order

| Input | Representation and consumer |
|---|---|
| ALM Tiles | W×H u16 LE words selecting tile group/subcell and flags |
| ALM Altitudes | W×H bytes; four neighboring vertices determine cell geometry |
| terrain and terrain.3d strips | Ordinary source pixels and optional precomposition inputs; equal paired geometry, unequal RGB |
| Sun, visibility and cell state | Lighting tables, fog and passability |

Load the planes and selected bitmap groups, build the lighting tables, then
select each tile's animated subcell. The four corner heights choose flat or
sloped drawing and its edge walk. Sprite/structure passes have separate
placement and palette rules. Movement consumes the block planes, not the
drawn pixel colors. — TERR-LOAD-002, TERR-IDX-003 (amended bit interval;
arithmetic retained), ALM-GRID-012,
ALM-GRID-013, TERR-GRID-027

## Reference map

| Reference | Contents |
|---|---|
| <a id="invariants-hold-across-all-38-maps-880-704-cells-no-exclusions"></a><a id="reading-algorithm"></a><a id="location--format"></a><a id="loader-romexe-R0516"></a><a id="tile-word--graphic-render-mapping"></a><a id="installed-tile-domain"></a><a id="tile-draw-sequence"></a><a id="water-animation-terr-anim-006terr-anim-010"></a> [Tile resources and selection](tiles.md) | Tile-word lookup, strip loading, dirt and water animation |
| <a id="tile-word--movement-terr-pass-049terr-pass-051-terr-cost-052-terr-pass-053"></a><a id="who-is-asking--the-mover-terr-move-054terr-move-057"></a><a id="movement-speed-terr-move-056"></a> [Passability and cell records](movement.md) | Block/cost planes, movement domains and cell records |
| <a id="cell-geometry--where-the-pixels-land-terr-geom-031terr-geom-036"></a><a id="far-edge-cells-terr-edge-024terr-edge-026"></a> [Cell geometry and far edges](geometry.md) | Corner sampling, edge tables, clipping and last-row behavior |
| <a id="terrain-lighting-terr-light-011terr-light-013-terr-light-015"></a> [Terrain lighting and fog](lighting.md) | Gradient, sun cycle, relight, fog and bounded field consumers |
| <a id="sprite-placement--where-a-unit-or-object-stands-terr-spr-038terr-spr-041"></a><a id="which-frame-a-static-object-draws-terr-spr-042-terr-spr-043-terr-tile-044"></a><a id="sprite-lighting--how-bright-a-unit-or-object-is-drawn-terr-light-059terr-light-064"></a><a id="terrain-output-range-terr-light-063"></a> [Sprite placement and composition](sprites.md) | Object/unit palette, placement, shadows and composition |
| <a id="structure-placement--where-a-building-stands-terr-struct-100terr-struct-107"></a> [Structure placement](structures.md) | Structure sprite grid and footprint distinctions |
| <a id="coverage-boundaries-and-open-questions"></a><a id="unknowns"></a> [Unknowns](limits.md) | Unsupported paths and unresolved runtime state |
| [Fog and visibility](fog.md) | Tile bits 14/15, stamp algorithm, shroud and saved exploration |
| <a id="the-order-one-frame-is-painted-in-and-everything-that-bounds-a-drawn-silhouette"></a><a id="a-unit-places-itself--and-its-shadow-places-itself-differently-terr-spr-047-terr-spr-065terr-spr-067-late-caller-population-amended"></a> [Sprite composition and shadows](composition.md) | Frame pass order, clip bounds, unit lift and shadow placement |
| <a id="structures-on-the-block-plane-terr-struct-068072-terr-pass-073"></a> [Cell records and Building attachment](cells.md) | Cell payload, Building masks, attach/detach and save/load state |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.

Metadata+0x0c storage at landscape P+0x2c shares the map reader's resolved
stream wrapper. Complete lower reads transfer the four bytes unchanged;
short reads can preserve part of the previous destination. Later P/M aliases,
first use and semantic role remain Unknown. — TERR-STREAM-157

The selected loader caller allocates 0x30 bytes for P; its constructor
assigns vtable L10958.
Publication and binder R0515 retain P at view+0x80; the binder copies
P+4/+8 into view dimensions and leaves P+0x2c untouched. On replacement,
the original slot+4 binds L10962→L10963, which frees arrays at
P+0x0c/+0x10/+0x14/+0x18 before the optional P free. Twelve bounded
original-x86 cases preserve the scalar through binding and observe no scalar
read in the executed destructor bodies. Allocator retention and the
post-publication message L10931 are unclosed; its receiver is selected
through session+0xd4, slot+0x48. Native LOAD, message delivery and first use
remain Unknown. This P allocation is distinct from M and the editor map E.
— TERR-METALIFE-175
