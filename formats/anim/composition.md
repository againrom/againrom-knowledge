# Drawable registration and composition

[Reference](format.md)

## Drawable registration and cell composition

`ANIM-REGISTER-083` joins registration selectors 0/1/2/3/4 to map-view planes
`+0x98/+0x94/+0x8c/+0x90/+0x9c`. Selector 1 stores the prepared footprint
rectangle; the others store one computed anchor. Registration requires
`+0x4c!=0`. A plane/cell retains one pointer, so later writes replace earlier
ones. Native refresh, collision-winner and stale-entry lifetime remain Unknown.

CBackPack uses selector 0; structures and bridge subclasses use selector 1.
CUnit uses selector 2 only when unsigned `+0x15a<2` and bit `0x80` of `+0x18c`
is clear; otherwise it uses selector 4. CAirUnit uses selector 3 without those
tests. Stage 1 therefore need not mean the earlier plane, and the alternate
flag can change the route at stage 0. — ANIM-CATEGORY-084

The earlier cell sweep dispatches selector 4 shadow/body before CBackPack
shadow/body in that cell. The main cell sweep dispatches non-flat structure,
selector 2 shadow/body, auxiliary `vt+0x18` dispatch, static objects and then area
selector 1. Static objects use their map-byte/class-array path. The late split
shadow/body sweeps belong to selector 3 (`ANIM-CELL-085`, `ANIM-AIRPASS-086`).
Area selectors and drawable registration selectors are separate domains.

Both structure body phases admit signed `+0x78<2`; their Flat predicates choose
the phase. CUnit/CAirUnit and CBackPack body paths require `+0x78==0`, with an
additional key/indexed-mask gate on selector 2. Shadow dispatches have an
independent global enable gate (`ANIM-DRAWGATE-087`).

Nine cell sweeps and one non-cell collection walk form the ten named phases.
Drawable cell sweeps walk rows ascending and columns descending; the marker/bar
window is narrower and the shroud is column-major. A later phase follows the
completed earlier phase regardless of cell coordinates. Within a drawable
phase, later cells have greater row or equal row and smaller column
(`ANIM-WALKORDER-088`). Pixel coverage, opaque overlap and native lifecycle
equivalence remain outside these static dispatch results.
