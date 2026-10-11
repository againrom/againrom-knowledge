# Claim registry — TERRAIN (terrain.3d tile graphics)

Level 2 ledger. Index: [registry.md](registry.md) · spec: [`formats/terrain/format.md`](../formats/terrain/format.md). Legend: `✔ promoted` (also in the spec) · `● active` · `✖ retracted`. IDs are permanent.

> **Confidence re-graded by the standing sweep (2026-07-25).** Every row was re-judged against
> the evidence it already cites, under the discriminating-evidence rule. No new bytes were read.
> This ledger came through well — it is overwhelmingly instruction-level — and the three rows
> `AGENTS.md` predicted would fail (`TERR-VER-005`, `TERR-ANIM-010`, `TERR-LIGHT-016`) are
> exactly the three that did. **`TERR-LIGHT-016`'s figures are withdrawn**: its height census is
> `TERR-LIGHT-018`'s failure a second time, and worse — see the row.

| ID | Claim | Confidence | Status | Evidence |
|----|-------|-----------|--------|----------|
| TERR-LOC-001 | Terrain graphics in `graphics.res` include 53 uncompressed 8-bpp BMPs under `terrain.3d/`: dirt 32×128, 36 land/road strips 32×448 and 16 water strips 32×256, with a 40-byte header, 256-colour palette and pixel offset 1078. | High | ● active (partially retracted) | [EXP-0021](../experiments/EXP-0021-terrain-graphics/), [EXP-0370](../experiments/EXP-0370-terrain-family/) |
| TERR-LOAD-002 | Loader `R0516` uses 128 source pointers at `L10218..L10219`, indexed by `(G-1)*16+V`, and dirt at `L10220`. | High | ● active (amended) | [EXP-0021](../experiments/EXP-0021-terrain-graphics/), [EXP-0370](../experiments/EXP-0370-terrain-family/) |
| TERR-IDX-003 | tile-word → graphic render mapping | High | ● active (partially retracted) | [EXP-0021](../experiments/EXP-0021-terrain-graphics/), [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |
| TERR-SEM-004 | The four file groups are the EXP-0020 terrain strips: **`tile1`** = Land-based (g0 Grass, g1 Cracked, g2 Sand, g3 Savanna), **`tile2`** = g4 Stones, g5 Cracked/Stones, g6 Flowers/Savanna, g7 Mountain/Stones, **`tile3`** = **Water** ... | High / Medium | ● active | [EXP-0021](../experiments/EXP-0021-terrain-graphics/) |
| TERR-VER-005 | Corpus verification (38 maps, 880 552 type1 cells, overlay corner excluded) with two discriminating falsification checks passing at **0 violations**: (a) strip group `g` occurs only in **0..12** so `G ∈ {1,2,3,4}`, and **0 cells** ... | High / Medium | ● active | [EXP-0021](../experiments/EXP-0021-terrain-graphics/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| TERR-ANIM-006 | Water animation — sequence. | High | ● active | [EXP-0022](../experiments/EXP-0022-water-animation/) |
| TERR-ANIM-007 | Water animation — counter. | High | ● active | [EXP-0022](../experiments/EXP-0022-water-animation/) |
| TERR-ANIM-008 | Water animation — cadence. | High / Medium | ● active (amended, superseded) | [EXP-0022](../experiments/EXP-0022-water-animation/), **[EXP-0095](../experiments/EXP-0095-walk-cadence/)** |
| TERR-ANIM-009 | Water animation — enable flag. | High | ● active | [EXP-0022](../experiments/EXP-0022-water-animation/) |
| TERR-ANIM-010 | Water animation — corpus falsification. | Medium | ● active | [EXP-0022](../experiments/EXP-0022-water-animation/) |
| TERR-LIGHT-011 | Terrain is relief-shaded, not flat-copied. | High | ● active (amended, superseded) | [EXP-0023](../experiments/EXP-0023-terrain-lighting/), [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/) |
| TERR-LIGHT-012 | The brightness grid is runtime-computed, not stored. | High | ● active | [EXP-0023](../experiments/EXP-0023-terrain-lighting/) |
| TERR-LIGHT-013 | Per-vertex relief model (`R0468`, writes `P+0x18`) — exact byte. | High | ● active (amended) | [EXP-0023](../experiments/EXP-0023-terrain-lighting/), [EXP-0029](../experiments/EXP-0029-terrain-edge-cells/), [EXP-0031](../experiments/EXP-0031-terrain-gradient-span/) |
| TERR-LIGHT-014 | Day/night sun model (`R1813`). | High / Medium | ● active (amended, partially retracted) | [EXP-0023](../experiments/EXP-0023-terrain-lighting/), [EXP-0025](../experiments/EXP-0025-shading-table/), [EXP-0116](../experiments/EXP-0116-daylight-cycle/) |
| TERR-LIGHT-015 | Solar-angle field identity (resolves `ALM-META-009` R-2). | High / Unknown | ● active (amended) | [EXP-0023](../experiments/EXP-0023-terrain-lighting/), [EXP-0025](../experiments/EXP-0025-shading-table/), [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |
| TERR-LIGHT-016 | Corpus: relief is meaningful everywhere | Medium | ● active (amended) | [EXP-0023](../experiments/EXP-0023-terrain-lighting/) |
| TERR-DIRT-017 | Impassable dirt composite = transparent-keyed overlay, not a blend. | High | ● active | [EXP-0024](../experiments/EXP-0024-dirt-composite/) |
| TERR-LIGHT-018 | Shading-table builder + dimensions. | High | ● active (amended) | [EXP-0025](../experiments/EXP-0025-shading-table/), [EXP-0031](../experiments/EXP-0031-terrain-gradient-span/) |
| TERR-LIGHT-019 | Exact per-entry transform (mode 3). | High / Medium | ● active | [EXP-0025](../experiments/EXP-0025-shading-table/) |
| TERR-LIGHT-020 | Level 64 is unattenuated; the byte is an attenuation index. | High / Medium | ● active | [EXP-0025](../experiments/EXP-0025-shading-table/) |
| TERR-LIGHT-021 | Sky tint enters the table; the intensity bytes do not. | High | ● active | [EXP-0025](../experiments/EXP-0025-shading-table/) |
| TERR-LIGHT-022 | One table serves all terrain; relight trigger. | High | ● active | [EXP-0025](../experiments/EXP-0025-shading-table/) |
| TERR-LIGHT-023 | Type-0 light-field offsets (resolves the `TERR-LIGHT-015` prose/evidence conflict). | High / Unknown | ● active (amended) | [EXP-0025](../experiments/EXP-0025-shading-table/), [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |
| TERR-EDGE-024 | Far-edge cell corner sourcing = raw flat row-major addressing (no clamp/dup/wrap/table). | High | ● active | [EXP-0029](../experiments/EXP-0029-terrain-edge-cells/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| TERR-EDGE-025 | The per-vertex brightness OUTER RING is never computed (amends `TERR-LIGHT-013`). | High / Unknown | ● active | [EXP-0029](../experiments/EXP-0029-terrain-edge-cells/) |
| TERR-EDGE-026 | The projection and drawable paths overscan the view; the inference that this admits the stored outermost terrain cells is partially retracted by TERR-217. | High / Medium | ● active (partially retracted) | [EXP-0029](../experiments/EXP-0029-terrain-edge-cells/) |
| TERR-GRID-027 | The tile grid every `TERR-*` render claim consumes starts 8 bytes later in the `.alm` than `ALM-GRID-011` said — a render indexing the section payload from byte 0 is four cells off in X. | High | ● active | [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| TERR-LIGHT-028 | The height-gradient span is the ADJACENT ONE-CELL difference — twice, along the same axis (amends `TERR-LIGHT-013`; resolves R-3). | High | ● active | [EXP-0031](../experiments/EXP-0031-terrain-gradient-span/) |
| TERR-LIGHT-029 | Corpus level census, corrected — and it is a function of θ (corrects `TERR-LIGHT-018`'s `30…74`). | High / Medium | ● active | [EXP-0031](../experiments/EXP-0031-terrain-gradient-span/) |
| TERR-LIGHT-030 | The sun model's cycle-off default θ is a literal, and it is not π/4. | High | ● active (amended, partially retracted) | [EXP-0031](../experiments/EXP-0031-terrain-gradient-span/), [EXP-0116](../experiments/EXP-0116-daylight-cycle/) |
| TERR-GEOM-031 | The altitude becomes a destination row, by subtraction — and the heights are signed at the point of use. | High | ● active (amended) | [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/) |
| TERR-GEOM-032 | Flat vs sloped is chosen from the four projected corner Y, and is exactly "all four corner altitudes equal". | High | ● active | [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/) |
| TERR-GEOM-033 | What the two blitters actually draw. | High | ● active | [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/) |
| TERR-GEOM-034 | The edge interpolation is a table walk, and the table is a startup-built Bresenham step table. | High | ● active (amended) | [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/) |
| TERR-GEOM-035 | The displacement is vertical only, and the horizontal clip is whole-cell. | High | ● active | [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/) |
| TERR-GEOM-036 | Ownership, seams, and the two places the model degenerates. | High / Medium | ● active (amended, superseded) | [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/) |
| TERR-FOG-037 | The three "further blitters" are not sprite paths — they are the shroud/fog pass over the *same* terrain quad. | High | ● active (amended, superseded) | [EXP-0036](../experiments/EXP-0036-sprite-geometry/) |
| TERR-SPR-038 | A sprite is lifted by the terrain, never warped by it. | High | ● active (amended) | [EXP-0036](../experiments/EXP-0036-sprite-geometry/) |
| TERR-SPR-039 | The altitude a sprite is lifted by is a *different* grid from the one the terrain is drawn on. | High / Medium | ● active (amended) | [EXP-0036](../experiments/EXP-0036-sprite-geometry/) |
| TERR-SPR-040 | The anchor: the class canvas is centred on the cell centre, and `(CenterX, CenterY)` is the pixel that touches the ground. | High | ● active (amended) | [EXP-0036](../experiments/EXP-0036-sprite-geometry/), [EXP-0038](../experiments/EXP-0038-object-frame/) |
| TERR-SPR-041 | What is settled for units, and what is not. | High / Medium | ● active (amended) | [EXP-0036](../experiments/EXP-0036-sprite-geometry/), [EXP-0039](../experiments/EXP-0039-unit-frame/) |
| TERR-SPR-042 | What reaches the sprite draw call's `frame` argument — three arms and two gates, none of which is "`Index`, always". | High / Medium / Unknown | ● active (amended, superseded) | [EXP-0038](../experiments/EXP-0038-object-frame/) |
| TERR-SPR-043 | The anchor's frame index is not always 0 — the body pass uses the drawn frame, the shadow pass uses frame 0 (amends `TERR-SPR-040`). | High / Medium | ● active | [EXP-0038](../experiments/EXP-0038-object-frame/) |
| TERR-TILE-044 | The tile word's top three bits: bit 13 is mutable at runtime and shared by three consumers; bits 15..14 gate a whole second draw path and are 0 on every shipped cell. | High / Medium | ● active (amended, partially retracted) | [EXP-0038](../experiments/EXP-0038-object-frame/), **[EXP-0040](../experiments/EXP-0040-enumeration-audit/)**, **[EXP-0084](../experiments/EXP-0084-idle-animation/)**, **[EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/)** |
| TERR-SPR-047 | What reaches a unit's `frame` argument: nine animation states, six arms, and a default that draws the class id. | High / Medium / Unknown | ● active (amended) | [EXP-0039](../experiments/EXP-0039-unit-frame/) |
| TERR-SPR-048 | Where a unit is drawn, and by which dispatch — the unit sprite pass is `vt+0x2c`, not `vt+0x30`. | High | ● active (partially retracted) | [EXP-0039](../experiments/EXP-0039-unit-frame/), [EXP-0501](../experiments/EXP-0501-status-bar-colour/) |
| TERR-PASS-049 | The map→sim ingest: three 256-stride byte planes, and the five arms that set the block byte. | High / Medium | ● active | [EXP-0041](../experiments/EXP-0041-terrain-passability/) |
| TERR-PASS-050 | The terrain classifier `R0469`: which tile-word bits reach movement, and the 5-level cost blend. | High / Medium | ● active (partially retracted) | [EXP-0041](../experiments/EXP-0041-terrain-passability/), [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |
| TERR-PASS-051 | The rule: a cell blocks a mover iff `block[cell] & mover.mask` over the mover's `n×n` footprint — the block byte is a BITMASK, not the enum `ALM-TERR-016` published. | High / Medium | ● active (amended, superseded) | [EXP-0041](../experiments/EXP-0041-terrain-passability/), **[EXP-0044](../experiments/EXP-0044-mover-domain/)**, **[EXP-0078](../experiments/EXP-0078-movement-domains/)** |
| TERR-COST-052 | What the restored `Cost` scalars are consumed by — and it is never the block decision. | High / Unknown | ● active (amended) | [EXP-0041](../experiments/EXP-0041-terrain-passability/), **[EXP-0054](../experiments/EXP-0054-unit-movement/)** |
| TERR-PASS-053 | The block planes reach persistent state; the cost plane does not. | High / Medium | ● active | [EXP-0041](../experiments/EXP-0041-terrain-passability/) |
| TERR-MOVE-054 | Who the mover is: not `CUnit`. | High | ● active | [EXP-0044](../experiments/EXP-0044-mover-domain/) |
| TERR-MOVE-055 | What the two virtuals return, and that nothing in the image ever varies them. | High | ● active (amended) | [EXP-0044](../experiments/EXP-0044-mover-domain/), **[EXP-0049](../experiments/EXP-0049-placeable-db/)** |
| TERR-MOVE-056 | The per-step movement speed, `R1088`, end to end. | High / Medium | ● active | [EXP-0044](../experiments/EXP-0044-mover-domain/) |
| TERR-MOVE-057 | What fills the simulation actor's `+0x49`/`+0x4a` (EXP-0044's open item): the placeable-definition database, at spawn — and block-mask codes 2 and 3 ARE produced. | High | ● active (amended, superseded) | [EXP-0049](../experiments/EXP-0049-placeable-db/) |
| TERR-MOVE-058 | The speed override `actor+0x70 → +0x3c → +0x44` (`TERR-MOVE-056`) is NOT the definition database; the DB supplies the fallback. | High | ● active (amended) | [EXP-0049](../experiments/EXP-0049-placeable-db/) |
| TERR-LIGHT-059 | A sprite IS lit — and the sprite blit family splits in two, decided by which pointer its pixel loop reads. | High | ● active | [EXP-0053](../experiments/EXP-0053-sprite-lighting/) |
| TERR-LIGHT-060 | Mode 2 — the exact per-entry transform every sprite table is built with. | High | ● active | [EXP-0053](../experiments/EXP-0053-sprite-lighting/) |
| TERR-LIGHT-061 | What feeds a sprite's level: `CMapView+0xb0`, a per-frame byte grid whose base value is a single global. | High / Medium | ● active (amended, partially retracted) | [EXP-0053](../experiments/EXP-0053-sprite-lighting/), [EXP-0503](../experiments/EXP-0503-spell-light/EXP-0503.md) |
| TERR-LIGHT-062 | The sprite ladder and the terrain ladder are the SAME ladder, and in the shipped daytime band the two are half a step apart. | High / Medium | ● active | [EXP-0053](../experiments/EXP-0053-sprite-lighting/) |
| TERR-LIGHT-063 | The terrain mapping's OUTPUT range, and it saturates to white on shipped data. | High | ● active | [EXP-0053](../experiments/EXP-0053-sprite-lighting/) |
| TERR-LIGHT-064 | Which drawable gets which table, and the palettes are disjoint. | High / Unknown | ● active (amended, partially retracted) | [EXP-0053](../experiments/EXP-0053-sprite-lighting/), [EXP-0089](../experiments/EXP-0089-tier-hue/) |
| TERR-SPR-065 | The unit BODY pass is `vt+0x28`, not `vt+0x2c` — `vt+0x2c` is the unit's shadow (amends `TERR-SPR-048`). | High | ✔ promoted (partially retracted) | [EXP-0053](../experiments/EXP-0053-sprite-lighting/); [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |
| TERR-SPR-066 | The silhouette blit's fifth argument is the SUN SHEAR, which reconciles the contradiction `TERR-SPR-048` left open (amends `TERR-SPR-038`). | High | ● active | [EXP-0053](../experiments/EXP-0053-sprite-lighting/) |
| TERR-SPR-067 | The unit BODY places itself one term away from its own shadow — the rule `TERR-SPR-048` published is the shadow's. | High / Medium | ● active (amended) | [EXP-0053](../experiments/EXP-0053-sprite-lighting/) |
| TERR-STRUCT-068 | How a placed structure makes its footprint impassable — and it is BOTH planes, from the object's own constructor, not the map ingest. | High | ● active (amended, superseded) | [EXP-0070](../experiments/EXP-0070-structure-passability/), **[EXP-0081](../experiments/EXP-0081-bridge-crossing/)** |
| TERR-STRUCT-069 | The cell record: the plane byte is a cache, the record is the storage. | High / Unknown | ● active | [EXP-0070](../experiments/EXP-0070-structure-passability/) |
| TERR-STRUCT-070 | A structure's footprint is a rectangle plus TWO 32-bit masks, and it is not the mover's `n×n`. | High / Medium | ● active (amended, partially retracted) | [EXP-0070](../experiments/EXP-0070-structure-passability/), **[EXP-0081](../experiments/EXP-0081-bridge-crossing/)** |
| TERR-STRUCT-071 | Something SUBTRACTS from the block plane: a footprint cell whose `Passability` bit is clear has bits 0 and 2 erased on both planes — this is how a bridge crosses water. | High | ● active (amended) | [EXP-0070](../experiments/EXP-0070-structure-passability/) |
| TERR-STRUCT-072 | When it happens, and that it is undone. | High / Medium | ● active | [EXP-0070](../experiments/EXP-0070-structure-passability/) |
| TERR-PASS-073 | The two block planes' complete writer list, the instrument, and where the bit semantics stop. | High / Unknown | ● active (amended, partially retracted) | [EXP-0070](../experiments/EXP-0070-structure-passability/), [EXP-0174](../experiments/EXP-0174-area-movement/) |
| TERR-STRUCT-074 | A bridge deck is free to a ground mover, and the crossing is load-bearing — the model's own prediction, measured. | High / Medium | ● active (amended) | [EXP-0081](../experiments/EXP-0081-bridge-crossing/), **[EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/)** |
| TERR-STRUCT-075 | `obj+0x10` is a POINTER to a 12-byte position object — not two coordinate bytes — and with it the footprint anchor is read from the image instead of fitted to the corpus. | High | ● active (superseded, contested) | [EXP-0081](../experiments/EXP-0081-bridge-crossing/) |
| TERR-STRUCT-076 | The attach chain reaches the recompute on ONE of `R1357`'s two paths, not unconditionally. | High / Unknown | ● active | [EXP-0081](../experiments/EXP-0081-bridge-crossing/) |
| TERR-STRUCT-077 | The `.alm` extension kind is a bridge class, and it opens its whole rectangle. | High / Medium | ● active (amended, superseded) | [EXP-0081](../experiments/EXP-0081-bridge-crossing/), [EXP-0091](../experiments/EXP-0091-placement-axes/) |
| TERR-STRUCT-078 | The recompute's polarity, decided by one branch displacement — a SET `Passability` bit BLOCKS, and the shipped column title is the inverse of what the code does with it. | High / Medium | ● active | [EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/) |
| TERR-TILE-079 | Tile bits 15..14 are the FOG OF WAR, and the gate they feed is satisfied exactly where the local view currently has line of sight — so the animated arm and the partial-repaint renderer are live, not dead. | High / Medium | ● active (amended, superseded) | [EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/) |
| TERR-FOG-080 | The stamp's 41×41 mask is not a shape — it is a per-drawable line-of-sight field, rebuilt from scratch on every stamp. | High / Medium | ● active (superseded) | [EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/), **[EXP-0120](../experiments/EXP-0120-sight-writers/)** |
| TERR-FOG-081 | The sight value that drives the fog has TWO widths, and which one you get depends on whether the actor is a hero. | High / Medium | ● active | [EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/) |
| TERR-FOG-082 | How a consumer BRANCHES on the three states, and why the fourth is unreachable. | High | ● active (partially retracted) | [EXP-0088](../experiments/EXP-0088-fog-of-war/), [EXP-0349](../experiments/EXP-0349-tile-word-consumers/) |
| TERR-FOG-083 | What the shroud pass consumes is not the tile word but a per-frame projection of it: `CMapView+0xa0`, one dword per lattice vertex, memset and rebuilt every frame. | High | ● active | [EXP-0088](../experiments/EXP-0088-fog-of-war/) |
| TERR-FOG-084 | The shroud table `[L03341]` is 17 rows and its law is one line: `out_channel = (in_channel × (16 − L)) >> 4`. This CLOSES `TERR-FOG-037`'s Medium — level 16 is exactly black and level 8 exactly half. | High | ● active | [EXP-0088](../experiments/EXP-0088-fog-of-war/) |
| TERR-FOG-085 | The clock: the map-wide clear runs on the PRESENTATION tick with period 32 — neither simulation clock touches the fog — and one global can freeze it. | High / Unknown | ● active | [EXP-0088](../experiments/EXP-0088-fog-of-war/) |
| TERR-FOG-086 | The fog gates two things — the shroud pixels and whether a drawable is drawn at all — and reaches nothing in the simulation. | High / Medium | ● active (partially retracted) | [EXP-0088](../experiments/EXP-0088-fog-of-war/), [EXP-0349](../experiments/EXP-0349-tile-word-consumers/) |
| TERR-FOG-087 | No shipped map authors either fog bit | High | ● active (partially retracted) | [EXP-0088](../experiments/EXP-0088-fog-of-war/), [EXP-0150](../experiments/EXP-0150-save-fog-record/) |
| TERR-FOG-088 | The engine implements the same line-of-sight algorithm TWICE, on two objects, and only one of them is the player's fog — so the AI's vision and the fog are two mechanisms that a consumer must not merge. | High | ● active (amended) | [EXP-0088](../experiments/EXP-0088-fog-of-war/), **[EXP-0120](../experiments/EXP-0120-sight-writers/)** |
| TERR-FOG-089 | The reveal-permission array has THREE writers, not two, and the one `TERR-TILE-079` could not see is a wholesale `memcpy` — so the two message opcodes are now named. | High / Medium | ● active | [EXP-0088](../experiments/EXP-0088-fog-of-war/) |
| TERR-STRUCT-090 | The footprint-override arm is selected by the SUM of the resolver's two byte arguments, not by the extension kind — and it has exactly ONE reachable caller in the image. | High | ● active | [EXP-0091](../experiments/EXP-0091-placement-axes/) |

### TERR-LOC-001

Terrain graphics in `graphics.res` include 53 uncompressed 8-bpp BMPs under `terrain.3d/`: dirt 32×128, 36 land/road strips 32×448 and 16 water strips 32×256, with a 40-byte header, 256-colour palette and pixel offset 1078. The vertical 32×32 sub-cell interpretation stands. **Amended:** the parallel `terrain/` family is not a fallback selected instead of these sources. `TERR-FAMILY-187` establishes optional 3d precomposition followed by ordinary source loading and withdraws the unsupported shipped-route promise. `TERR-FAMILY-188` specifies both families and their header-length anomalies.

**Confidence.** High for the format/strip identification. The old selector-to-shipped-route inference is withdrawn; the new producer search remains bounded.

**Original status.** ● partially retracted

**Evidence.** [EXP-0021](../experiments/EXP-0021-terrain-graphics/), [EXP-0370](../experiments/EXP-0370-terrain-family/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### TERR-LOAD-002

Loader `R0516` uses 128 source pointers at `L10218..L10219`, indexed by `(G-1)*16+V`, and dirt at `L10220`. **Amended:** its six path strings comprise three per family. Selector mask 0x02 enables an earlier `terrain.3d` load/precomposition; temporary source objects are then deleted before the ordinary `terrain/` pass fills the final source table. Separate work surfaces at `L10221` survive that source cleanup. The final default-handle search checks block leaders at indices 0,4,...,124 and stores the first non-null object address plus 0x14, else zero, at `L10222`. Group-mask selection is `TERR-LOAD-152`; full order and native-reachability bounds are `TERR-FAMILY-187`.

**Confidence.** High for the bounded loader operations and table arithmetic. Complete native graphics success and activation remain Unknown.

**Original status.** ● active (amended)

**Evidence.** [EXP-0021](../experiments/EXP-0021-terrain-graphics/), [EXP-0370](../experiments/EXP-0370-terrain-family/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-IDX-003

**tile-word → graphic render mapping** (four agreeing `rom.exe` renderers `R0539`/`R1814`/`R1815`/`R1816`): for cell word `w`, `g = (w & 0x1fff) >> 6` (strip group, bits 6–12); the image = `tiles[g*4 + ((w>>4)&3)]`; the drawn 32×32 sub-cell = `w & 0xf` (bits 0–3), source byte offset `pixels + 8 + (w&0xf)*0x400` (`0x400` = 32×32 @ 8bpp). In filename terms: **file group `G = (g>>2)+1`**, **file variant `V = (g&3)*4 + ((w>>4)&3)`** ⇒ `tileG-VV.bmp`, **sub-cell row `= w & 0xf`**

**Confidence.** High (the masks, the shift and the `0x400` stride are immediates in four independently-compiled renderers that agree; a different bit split would have to appear in all four)

**Original status.** ● active, partially retracted: bit-interval parenthetical; arithmetic retained

**Evidence.** [EXP-0021](../experiments/EXP-0021-terrain-graphics/), [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

**Amended.** The table ledger carried the status "● active, partially retracted: bit-interval parenthetical; arithmetic retained". `retracted.md` records a correction against this claim.

### TERR-SEM-004

The four file groups are the EXP-0020 terrain strips: **`tile1`** = Land-based (g0 Grass, g1 Cracked, g2 Sand, g3 Savanna), **`tile2`** = g4 Stones, g5 Cracked/Stones, g6 Flowers/Savanna, g7 Mountain/Stones, **`tile3`** = **Water** (the animated group `g 8..11` — the renderer replaces the group with `8+phase`, `phase = f(pos, animCtr)&3` when the global at `L05659` set, else 0), **`tile4`** = **Road** (g12). Each land group's 16 variants = its 4 strip subtypes (`g&3`) × 4 blend columns (bits 4–5). A tile with **bit 13** (impassable, non-water) is composited over `dirt.bmp` (`dirt` sub-cell `= (col + row*5)&3`) before blit

**Confidence.** High (grouping — the group→file mapping and the bit-13 dirt path are instruction-level, `TERR-IDX-003`/`TERR-DIRT-017`) / Medium (human terrain labels, inherited from `ALM-TERR-015` — the *names* there survive the 2026-07-25 sweep, only its `Cost`/`Pass` figures were withdrawn; water-phase constants transcribed)

**Original status.** ● active

**Evidence.** [EXP-0021](../experiments/EXP-0021-terrain-graphics/)

### TERR-VER-005

Corpus verification (38 maps, 880 552 type1 cells, overlay corner excluded) with two discriminating falsification checks passing at **0 violations**: (a) strip group `g` occurs only in **0..12** so `G ∈ {1,2,3,4}`, and **0 cells** reference an absent tile file — `tile4` is referenced only at `g=12` → variants 00–03, exactly the 4 files present; (b) **land** sub-cell max = **13** (14-cell BMPs) and **water** sub-cell max = **7** (tile3 is only 8 cells tall) — the shorter water BMP is respected exactly, which pins `g∈8..11 ⇒ tile3`. A wrong group/variant/sub split fails both. **Strengthened by [EXP-0030]: the "overlay corner excluded" carve-out is retired** — with the corrected grid base there is no overlay, and re-running the census over **all 880 704 cells** (0 excluded) still gives 0 absent-file references, land sub-max 13 and water sub-max 7 at 0 out-of-range. The 152 previously-excluded words were never cells

**Confidence.** High (as a **measurement**, and one of the few whose base move was handled: 38 maps, **all 880 704** type1 cells, 0 exclusions, EXP-0030 grid base, 0 violations — confirmed regenerated by EXP-0030's `-corner 0` re-run) / **Medium** (as a *pinning* of the group/variant/sub split: it is a consistency test against one enumerated family of wrong splits — those that would reference an absent file or overrun `tile3`'s 8 sub-cells — not against every second model. What actually pins the split is `TERR-IDX-003`'s four renderers; this row is corroboration, and EXP-0030 is the standing reminder that a census can pass under two framings at once)

**Original status.** ● active (strengthened)

**Evidence.** [EXP-0021](../experiments/EXP-0021-terrain-graphics/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

### TERR-ANIM-006

**Water animation — sequence.** A water cell (`g∈8..11`=`tile3`) keeps its blend column `b=(w>>4)&3` and sub-cell `w&0xf`; the four agreeing renderers (`rom.exe` `R0539`/`R1814`/`R1815`/`R1816`) override only the strip group with `8+phase`, `phase = (g + (worldCol+1)·worldRow + (animCtr>>2)) & 3`. So the drawn image walks `tile3` variant `V = phase*4 + b`, i.e. `tile3-{b, b+4, b+8, b+12}`, as `phase` steps `0→1→2→3→0`. The `(worldCol+1)·worldRow` term (scroll-adjusted world coords) offsets neighbours into a diagonal ripple, stable under scroll. The 16 `tile3` files = exactly 4 phases × 4 blend columns

**Confidence.** High (the group override and the `phase` expression are the same instruction sequence in four renderers; the gate on the global at `L05659` is in all four)

**Original status.** ● active

**Evidence.** [EXP-0022](../experiments/EXP-0022-water-animation/)

### TERR-ANIM-007

**Water animation — counter.** `animCtr = *(mapObj + 0xa70)`. A whole-binary scan shows exactly two writes: init-to-0 (`R0391` @L02611) and `+1` (`R0334` @L02476); every other access is a read. `R0334` is the handler of window message **`0x401`** (map window proc `R0333` `case 0x401`; `0x402` = render). So water advances **once per logic tick**, and `phase` uses `animCtr>>2` → one variant every **4 ticks**, full 4-variant cycle every **16 ticks**. No other code advances water

**Confidence.** High (a whole-binary write scan on `mapObj+0xa70` — two writes, both located, every other access a read; "once per logic tick" then follows from the `0x401` case, not from observation)

**Original status.** ● active

**Evidence.** [EXP-0022](../experiments/EXP-0022-water-animation/)

### TERR-ANIM-008

**Water animation — cadence.** The paced play loop `R0454` (selected by `R0560` for real-time single-player, field `+0x40c==0`) fires one `0x401` tick per `dtMs = *(mapObj+0x3f0)` ms via a `timeGetTime` accumulator; its step counter `+0x3e4` wraps `& 0xf` (16 steps = one water cycle). `dtMs = 1000/tps` is written by `SetGameSpeed` `R0559` (the constant `0x3e8` divided by the ticks per second, a signed division), which maps a clamped speed index `0..8` → `tps ∈ {8,10,12,14,16,20,24,28,32}`. **Map-load `R0099` @L10223 pushes index 4** (→16 tps→62 ms/tick→~248 ms/frame→~992 ms/cycle; ideal 62.5/250/1000) unless game-mode `+0x6bc==2` re-applies the persisted `+0x3f4`; the in-game +/- keys step `idx±1` (`R0231`), and message `0x445` loads it from config key 2. Per-frame ms = `4·(1000/tps)`, per-cycle = `16·(1000/tps)`; range speed0 500 ms/frame … speed8 124 ms/frame. **⚠ The MECHANISM clause is withdrawn by [EXP-0095]; the figures are not.** `campaign+0x3e4` is the pacer's own **16-tick resync window index**, not a water-cycle counter: `EnumRefs disp:3e4` is **20 hits / 9 owners / 0 orphan** and every owner is a pacer, `R0559`, or the in-game `+`/`-` handler — **no renderer reads it**. Its `& 0xf` wrap exists so that `campaign+0x3ec` can be rebased to `timeGetTime` whenever it is 0 (`L02632`), discarding accumulated lateness every 16 ticks. The real ambient cadence is a **latch on a different counter pair**: `L02614`/`L02615` compare `CMapView+0xa70` against `CMapView+0xa74` and `L02616` compares the difference with `0x3` and sets a flag when it is greater, which admits a repaint once the counter has advanced by more than 3, `L02617` re-latching it (`disp:a74` = **3 hits / 2 owners / 0 orphan**). Four phases × four ticks is also 16, which is exactly why the wrong mechanism survived — the two readings agree on every figure this row publishes and disagree only on which field carries them (`ANIM-AMBIENT-016`, `ANIM-PACE-017`). Also corrected: `dtMs = 1000/tps` is a **truncating** signed division (`L02622`), so the realised ms/tick are `125 100 83 71 62 50 41 35 31` and only three of the nine indices divide 1000 exactly

**Confidence.** High (per-index cadence + map-load default) / Medium (which index a saved/MP session resumes at) / ~~the `+0x3e4` attribution~~ **withdrawn by [EXP-0095]**

**Original status.** ● active (amended)

**Evidence.** [EXP-0022](../experiments/EXP-0022-water-animation/), **[EXP-0095](../experiments/EXP-0095-walk-cadence/)**

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-ANIM-009

**Water animation — enable flag.** the global at `L05659` gates animation (`if ([L05659]==0) phase=0` in the renderers). **[EXP-0038] widens its reach beyond water: the same global also forces a *static object*'s frame to 0** (`L10224`/`L10225`), so under `-noanimation`/`-detail0` every placed object draws frame 0 rather than `Frames[Index]` — nine `stones/sprites.256` classes collapse to one boulder (`TERR-SPR-042`). It is a third reader too: `R1814`, the partial-repaint terrain renderer, is called only when it is set (`TERR-TILE-044`). Static image default = **`0x00000001`** (on); the only writes (in `R0326`) set it to `0` under command-line `-noanimation` / `-detail0`; message `0x445` also loads it from config key `0x21`. So water animates by default and is disabled by those switches/detail setting. (The flag is shared with the fog-of-war reveal radius; siblings the global at `L06417`/`L06416` are other detail flags, not water.)

**Confidence.** High (default read from the static image; a whole-binary write scan finds only the two `R0326` stores and the `0x445` config load)

**Original status.** ● active

**Evidence.** [EXP-0022](../experiments/EXP-0022-water-animation/)

### TERR-ANIM-010

**Water animation — corpus falsification.** Across 38 maps / **87 580 water cells**, expanding all 4 animation phases (`V=phase*4+b`, 350 320 (cell,phase) pairs): **0** reference an absent `tile3` file and **0** exceed `tile3`'s 8 sub-cells; animation reaches exactly `tile3-00…15`, all present. Authored `b` distribution `{0:14639,1:50878,2:13328,3:8735}`. A wrong `b`/`V` split (e.g. `b` spanning 0..7) would push `V>15` into an absent file — it does not

**Confidence.** **Medium** (corpus statistics alone: this is a consistency test of the model `TERR-ANIM-006` established from four `rom.exe` renderers, against one named family of wrong splits — it does not discriminate a second model, and the row is corroboration, not the proof) / **⚠ stale base** (the counts — 87 580 water cells, 350 320 (cell,phase) pairs, and the `b` distribution — were measured in EXP-0022 on the **pre-EXP-0030 grid base**, the same census family as `ALM-GRID-012`, and were **not** regenerated. What *was* regenerated is the invariant that matters: EXP-0030's `-corner 0` re-run over all 880 704 cells still gives water sub-cell max 7 and 0 absent-file references)

**Original status.** ● active

**Evidence.** [EXP-0022](../experiments/EXP-0022-water-animation/)

### TERR-LIGHT-011

**Terrain is relief-shaded, not flat-copied.** The two terrain blitters `rom.exe` `R1506` (flat) / `R1507` (sloped) draw each pixel as `LUT16[(level<<9 & 0xfffffe00) + srcIndex*2]`, where `level` is **bilinearly interpolated from four per-vertex corner brightness bytes** (a Gouraud level ramp). A no-shading flat blit would not interpolate four distinct corners nor index a level dimension; two independent blitters do. **Amended by [EXP-0035] — this row was INCOMPLETE, not wrong, and the omission mattered: it described the two blitters' shading and said nothing about the fact that the sloped one also puts the pixels somewhere else (`TERR-GEOM-031…036`). Two corrections to the shading detail itself: (a) the per-row level step is `(levelBottom − levelTop)/32` (`SAR 5`) only in the FLAT blitter; the sloped one divides by the column's actual span height (`IDIV` @L10226), so the ramp runs corner-to-corner over the stretched span, not over 32 rows; (b) the accumulator is seeded `+0x100` (half a `0x200` LUT row) in both. The corner brightness bytes are read `MOVZX` (unsigned) — unlike the heights, which are `MOVSX`.**

**Confidence.** High (the LUT index arithmetic and the interpolation, instruction by instruction, in two independent blitters) — **and see `retracted.md`: this row was High and *incomplete* for twelve experiments, because reading a routine is not the same as reading everything it does**

**Original status.** ● active (amended)

**Evidence.** [EXP-0023](../experiments/EXP-0023-terrain-lighting/), [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-LIGHT-012

**The brightness grid is runtime-computed, not stored.** The CMapView landscape builder `R0480` allocates three `W·H` byte grids in the type2 case — `P+0x10` heights (read from file), `P+0x14` type3 (read in case 3), and **`P+0x18` allocated but never read from the file**. `P+0x18` is the grid the blitters sample; it is written by the lighting routine, confirming "the ALM stores no lightmap" — the shading is computed

**Confidence.** High (three allocations and their readers in one function; `P+0x18` has no file read anywhere on the path — an absence established over the whole builder, not assumed)

**Original status.** ● active

**Evidence.** [EXP-0023](../experiments/EXP-0023-terrain-lighting/)

### TERR-LIGHT-013

**Per-vertex relief model (`R0468`, writes `P+0x18`) — exact byte.** For each interior vertex, `stepH = 32.0/cos θ`; per axis `i` the slope `sᵢ = atan2(Δhᵢ, stepH)` from neighbour height deltas (the sun azimuth `θ` selecting the sampled neighbour via `tan θ`), then `axisᵢ = clamp(L − R·sin(π/6 − sᵢ), 0, 95)`; the **level byte** `= ftol(0.5·(axis₀ + axis₁))` (x87: a floating-point add, a multiply by 0.5, then a call to `__ftol`). Intensity: `R = [L02293]` (day 0x20=32), `L = (R>>1) + [L05650] + 0x20` (day = 62). Flat day vertex → 46. Written region: **only the strict interior `{1..W-2}×{1..H-2}`** — `R0468(1,1,0,0)` loops `col∈{1..W-2}, row∈{1..H-2}`, using a `±1` neighbour whose sign follows the sun azimuth (starting at 1 / ending at W-2 keeps both in bounds). **Amended by [EXP-0029] (`TERR-EDGE-025`): the outer ring of vertices (row 0, row H-1, col 0, col W-1) is NOT computed at all — there is no one-sided fallback; the outer ring is left at the non-zeroing allocator's value. The original "first row/col one-sided" wording read the decompiler's `if (iVar3 + -1 == 0)` as a border special-case; instruction-level it is not — both arms read row `row-1` and differ only in the column direction, and the flags the branch tests come from the decrement at `L10227`, which overwrites the zero flag set by the preceding test of bit `0x1` of the azimuth compare result. Nothing is clamped and the written region is unchanged. (What that column direction implies for the axis-B Δh term is noted in EXP-0029 but NOT re-derived; the Δh terms above stand as EXP-0023 published them.)** **Amended by [EXP-0031] (`TERR-LIGHT-028`): the Δh terms are now derived — they are two ADJACENT ONE-CELL differences along the SAME (row) axis, `Δh₁ = H[x∓tan\|θ\|, y+1] − H[x,y]` and `Δh₂ = H[x,y] − H[x∓tan\|θ\|, y−1]`, not two perpendicular axes and not a two-cell central difference. "The sun azimuth selecting the sampled neighbour via `tan θ`" is right but is a *lateral, same-row* interpolation, not the choice of gradient axis.** `θ = *(P+0x20)`. Constants (IEEE-754, read): `32.0` (`_L10228`), `π/6≈0.5236` (`_L10229`), `95.0` clamp (`_L05651`), `0.5` scale (`_L10230`)

**Confidence.** High (the transform `clamp(L − R·sin(π/6 − s), 0, 95)`, the average, the `__ftol`, the written region, and every constant — each an immediate or a named `_DAT_` read off the listing) / High (the Δh span — **but only since `TERR-LIGHT-028`**: the wording "formula + all constants read from the x87 disassembly" sat on this row for two experiments *while the span was wrong*, because the decompiler drops the FPU stack and the span was inferred, then confirmed by a fit that two different wrong rules both satisfy. See `retracted.md`; this is why "read from the disassembly" is not by itself a grade)

**Original status.** ● active (amended)

**Evidence.** [EXP-0023](../experiments/EXP-0023-terrain-lighting/), [EXP-0029](../experiments/EXP-0029-terrain-edge-cells/), [EXP-0031](../experiments/EXP-0031-terrain-gradient-span/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-LIGHT-014

**Day/night sun model (`R1813`).** Input `(clock>>4)+0x168`; `hour = (t/60)%24`; per time-of-day band sets a sky-light RGB tint (`[L10231]/491/492`; 0,0,0 in the 6–17 daytime band, graded at dawn/dusk), intensity bytes (`0x494=0x0e`, `0x498=0x20`), and the sun angle `_L07948` (radians). **⚠ The sweep in this row is WRONG and is corrected by `TERR-LIGHT-109`: `_L07947` is +0.78539815 and `_L07946` is −0.78539815, both a truncated-π QUARTER — the sweep is ±π/4, not ±π/2, and the "fixed ±π/4 at some bands" is the same number, not a different one.** **⚠ The two intensity figures in this row are the CYCLE-OFF and DAYTIME values only and are retracted as a description of the fields: `0x494` reaches 32 and `0x498` falls to 8 at night, both ramping through four twilight formulas - `TERR-LIGHT-119`, `retracted.md`.** The full six-arm colour schedule is `TERR-LIGHT-119`, the arm structure is `TERR-LIGHT-110`, the input's unit `SESS-TICK-026`, the cadence `TERR-LIGHT-114`, and whether the cycle runs at all `TERR-LIGHT-108`. `R1368` rebuilds the shading LUTs from the tile palettes on each light change (the global at `L10222` = the LUT handle the blitters pass). **Amended by [EXP-0025]:** this claim originally said "8-level" — the terrain table has **96** levels (`TERR-LIGHT-018`), and `R1368` is a relight *driver*, not the builder

**Confidence.** High (the LUT rebuild) / **the angle SWEEP is retracted** — off by a factor of two, corrected by `TERR-LIGHT-109` from the PE bytes on both roots / Medium (per-hour colour schedule, still not re-derived)

**Original status.** ✖ partially retracted (sweep corrected by `TERR-LIGHT-109`)

**Evidence.** [EXP-0023](../experiments/EXP-0023-terrain-lighting/), [EXP-0025](../experiments/EXP-0025-shading-table/), [EXP-0116](../experiments/EXP-0116-daylight-cycle/)

**Amended.** The table ledger carried the status "✖ partially retracted (sweep corrected by `TERR-LIGHT-109`)". `retracted.md` records a correction against this claim.

### TERR-LIGHT-015

**Solar-angle field identity (resolves `ALM-META-009` R-2).** The type-0 `+0x10` float is the map's **stored sun/light angle**. The landscape builder `R0480` case 0 reads it (forming the address `base+0x20`) into `P+0x20` — the **exact double `R0468` reads as θ** (a double load from `+0x20`); the adjacent type-0 scalars (**`+0x18`/`+0x1C`**, EXP-0018's "stored scalars") load into `P+0x1c/0x1d`, the light-**intensity** bytes (ambient/range). Units = radians, the same the sun model produces (θ ∈ [−π/2, +π/2]). **Corrected/amended by [EXP-0025]** (`TERR-LIGHT-023`): the intensity scalars are file `+0x18/+0x1C`, **not** `+0x14/+0x18` as this claim's prose first said (`+0x14` → `P+0x2c`, unread by lighting); and the stored values are **not** demonstrably the map's initial lighting — `R0468` overwrites `P+0x20/0x24` and `P+0x1c/0x1d` from the sun globals in its first ten instructions, before reading them. **The Unknown is answered by [EXP-0298] (`TERR-LIGHT-149`): no path reads them.** `θ`, ambient and range are overwritten before their first read on every path the direct call graph reaches, the load path forces that overwrite behind one gate (`TERR-LIGHT-151`), and their map-object copies are never read at all (`ALM-META-092`)

**Confidence.** High (shared slot + file offsets, from disasm) / **Unknown resolved to Medium** (no path reads the stored values, bounded by the searches `TERR-LIGHT-149` names)

**Original status.** ● active (amended)

**Evidence.** [EXP-0023](../experiments/EXP-0023-terrain-lighting/), [EXP-0025](../experiments/EXP-0025-shading-table/), [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

**Amended.** The table ledger carried the status "● active (amended)".

### TERR-LIGHT-016

**Corpus: relief is meaningful everywhere** — *claim reduced, figures withdrawn.* All **38/38** maps carry a type2 height grid: that part stands (`ALM-SEC-003`, `ALM-SEC-004`). **⚠ The figures "height σ 7.0…41.1; range 0…246" are WITHDRAWN by the 2026-07-25 sweep, and so is "0 are flat" as *this* row's evidence establishes it.** `evidence/height-census.txt` was measured on the **pre-EXP-0030 grid base**, and its every-map maximum is a record-header byte that base injected into the height grid: each map's `hMax` is exactly the third byte of that map's own `selectorA` `f32` (`ALM-HDR-029`) — Forester `0xBFF62B6D` → `0xF6` = **246**, Beast `0xBFF52D9D` → `0xF5` = **245**, Islands/Kids `0xBFC02B6D` → `0xC0` = **192** — and every one of the 38 rows reads 192, 245 or 246. **Not one map's true maximum height appears in that table.** It is decisively contradicted from inside this ledger: `TERR-LIGHT-028`(d) finds **0 of 880 704** shipped height bytes ≥ `0x80` on the corrected base, so no true height can exceed 127. This is `TERR-LIGHT-018`'s failure repeated — the same 8 injected bytes, the same un-regenerated evidence file — and it reaches further, because **`hStdev` is contaminated too**: on the 80×80 maps 8 bytes of value ~192 alone would produce σ ≈ 6.8, the same order as the smallest reported σ (Kids 7.04), so this evidence does **not** establish that those maps have any relief at all. Restoring the row is one command: re-run the same probe on the EXP-0030 base

**Confidence.** **Medium** (that every map has a height grid — from the record sizes, not from this census) / **⚠ withdrawn** (σ, range, and "0 are flat")

**Original status.** ● active (figures withdrawn)

**Evidence.** [EXP-0023](../experiments/EXP-0023-terrain-lighting/)

**Amended.** The table ledger carried the status "● active (figures withdrawn)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-DIRT-017

**Impassable dirt composite = transparent-keyed overlay, not a blend.** For a bit-13 (impassable) non-water tile, the renderer copies the terrain 32×32 sub-cell (`R0544`) then overlays a `dirt.bmp` sub-cell (`R1792`): per byte, a **non-zero** dirt palette index replaces the terrain pixel, a **zero** (index 0 = transparent, `SPR256-PAL-012` — this row cited `SPR256-PAL-013`, which is about zeroed low palettes and does not support the point; corrected by the 2026-07-25 sweep) leaves terrain showing — no arithmetic/blend. Dirt sub-cell `= (col + row*5)&3` (position hash over `dirt.bmp`'s 4 cells). The composite is drawn through the normal shaded blit, so the overlay is relief-lit with the ground

**Confidence.** High (the per-byte test and branch in `R1792`; "no arithmetic" is the absence of any blend instruction in that loop, read, not inferred)

**Original status.** ● active

**Evidence.** [EXP-0024](../experiments/EXP-0024-dirt-composite/)

### TERR-LIGHT-018

**Shading-table builder + dimensions.** `R1368` is a relight **driver**, not the builder: it walks the lit object lists and calls `R0919`, which frees the old table (`R1817`) and calls the real builder **`R1107`** with `(ECX = obj+0x14, palette = *(obj+0x20), nLevels, mode, useTint)`. Terrain is rebuilt as **`(0x60, 3, 1)`**. The builder does `*(this+4) = nLevels; *(this+8) = malloc(nLevels<<9)` — so the terrain table is **96 rows × 256 entries × 2 B = 49 152 B, row stride 512 = `1<<9`**, exactly the stride the blitters add (`TERR-LIGHT-011`). 96 is not arbitrary: it is the `[0,95]` clamp of the per-vertex byte (`TERR-LIGHT-013`, `_L05651 = 95.0`). Object layout (ctor `R1152`, `new(0x24)`): `+0x14` shading sub-object (= the global at `L10222`), `+0x18` nLevels, `+0x1c` table ptr, `+0x20` the BMP's own palette; renderers pass `*([L10222]+8)` = the table. **Supersedes `TERR-LIGHT-014`'s "8-level" wording.** Corpus check: `TERR-LIGHT-013` applied to all 38 height grids yields levels 30…74 — 0/880 704 vertices outside a 96-row table, 100 % outside an 8-row one. **Census figure corrected by [EXP-0031] (`TERR-LIGHT-029`): `30…74` is STALE — it was measured at EXP-0025 on the pre-[EXP-0030] grid base by a probe whose Δh rule was an approximation (½·two-cell central difference on two perpendicular axes), and its maximum lands on one of the 8 record-header bytes the superseded split injected into the height grid. On the corrected base with the exact `TERR-LIGHT-028` span the corpus yields `30…70` (θ = 0.78539815, interior window). The conclusion is unaffected and in fact unbreakable: the per-axis clamp bounds the level to `[30, 89]` for ANY height field, so 0 vertices can ever fall outside the 96 rows and 100 % must fall outside 8.**

**Confidence.** High (single-path instruction reads) / High (table bound) — census figure amended

**Original status.** ● active (amended)

**Evidence.** [EXP-0025](../experiments/EXP-0025-shading-table/), [EXP-0031](../experiments/EXP-0031-terrain-gradient-span/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-LIGHT-019

**Exact per-entry transform (mode 3).** Per level row and palette entry, per channel, in integers: `out = clamp(((palette_chan + skyTint_chan) × m) / 32, 0, 255)` with `m` counting **down** from `nLevels` to 1 while the destination advances (a multiply; trunc-toward-zero `/32` by sign-extending, adding a bias masked with `0x1f` and shifting right by 5; two 16-bit clamps at `L10232`/`L10233`). Packed as `(R>>(8−rBits))<<rShift \| (G>>(8−gBits))<<gShift \| (B>>(8−bBits))<<bShift`, the bits/shifts read from `[L06281]/L10234` (R), `L10235/L10236` (G), `L06282/L06471` (B) — which `R1242` derives at runtime from the DirectDraw surface's pixel-format masks (`lowBit`/`highBit`); the static image ships **RGB565** (11/5, 5/6, 0/5). Palette source order is the BMP RGBQUAD `[B,G,R,x]` (`SPR256-PAL-011`). Linear per channel — no gamma, no cross-channel term

**Confidence.** High (every operation named to an address; "linear, no gamma" is the absence of any further term in a loop read end to end) / Medium (that a **shipped** run is RGB565 — the bits come from the DirectDraw surface at runtime; RGB565 is what the static image's defaults give, not an observed session)

**Original status.** ● active

**Evidence.** [EXP-0025](../experiments/EXP-0025-shading-table/)

### TERR-LIGHT-020

**Level 64 is unattenuated; the byte is an attenuation index.** Row `L` carries multiplier `(96 − L)/32`: `L=0` → ×3.0, **`L=64` → ×1.0**, `L=95` → ×1/32. Higher level = darker. With the shipped daytime intensities the flat-ground level is 46 (`TERR-LIGHT-013`) → **×1.5625**, i.e. *flat daytime terrain is deliberately overdriven*. The shipped art is authored for it: pixel-weighted over the 651 264 `terrain.3d` pixels, mean luminance is **78.5/255 at ×1.0** (visibly dark) and **119.2 at ×1.5625** (a natural mid-tone), with 0 % of pixels clipping at level 64 and 10.6 % (the brightest highlights) at level 46. Consequence: rendering terrain "unshaded" means level **46**, not ×1.0 — ×1.0 is ~35 % too dark

**Confidence.** High (arithmetic) / Medium (the "authored for it" reading, inferred from luminance not observed in-game)

**Original status.** ● active

**Evidence.** [EXP-0025](../experiments/EXP-0025-shading-table/)

### TERR-LIGHT-021

**Sky tint enters the table; the intensity bytes do not.** The builder's only global reads are `[L10231]/491/492` (sky RGB), added **per channel to the palette entry before the level multiply**, and zeroed when the caller's 4th arg is 0 (terrain passes 1). `[L05650]/498` appear nowhere in the builder; their only readers are in `R0468`, as `R = P[+0x1d]` and `L = (R>>1) + P[+0x1c] + 0x20` — the per-vertex level's swing and base (confirms `TERR-LIGHT-013`'s L/R). Daytime band (`R1813` @L10237…): tint `0,0,0`, ambient `0x0e`, range `0x20`. So a tinted sky warms shadows more than highlights, which are already clipping

**Confidence.** High (the builder's global reads enumerated over the whole function — the claim is an **absence** (`[L05650]/498` appear nowhere in it) established by a whole-binary reader scan, which is the form of evidence that can carry an absence)

**Original status.** ● active

**Evidence.** [EXP-0025](../experiments/EXP-0025-shading-table/)

### TERR-LIGHT-022

**One table serves all terrain; relight trigger.** The driver's terrain loop scans `tiles[]` (the global at `L10218`) every 4 slots, rebuilds the **first non-null** object, sets `[L10222] = thatObj + 0x14` and `break`s — a single table for every terrain tile. Sound because **all 53 shipped `terrain.3d` BMPs carry a byte-identical 1 024-B palette** (max per-channel deviation from `tile1-00` = 0, measured). Non-terrain classes are rebuilt as `(0x10, 2, 1)` — a 16-level ramp whose neutral row is `nLevels/2` — including the 16 sprite palettes at `[L10238] + k*0x400`. Relight fires from `R0532` (the driver's **only** caller) when `(clock & 0xf) == 0 && ((clock>>4) + 0x168) % 0x14 == 0`, or when forced: sun model → per-vertex grid → all tables → redraw (`0x404`)

**Confidence.** High (the loop, the `break` and the relight condition are instruction-level, and `R0532` is the driver's only caller — a scan result, not a sample) / High (the palette identity — a **measurement over all 53 shipped files**, max per-channel deviation 0, with no model to be wrong about)

**Original status.** ● active

**Evidence.** [EXP-0025](../experiments/EXP-0025-shading-table/)

### TERR-LIGHT-023

**Type-0 light-field offsets (resolves the `TERR-LIGHT-015` prose/evidence conflict).** `R0480` case 0 issues seven `read(dest,4)` calls; anchoring them on the `W`/`H` slots (pinned by the function's own `malloc(*(P+8) * *(P+4))`) against `ALM-META-008`'s `W@+0x08, H@+0x0C, f32@+0x10` gives file → slot: `+0x08→P+0x04`, `+0x0C→P+0x08`, `+0x10→P+0x20` (θ), `+0x14→P+0x2c`, **`+0x18→P+0x1c`** (ambient), **`+0x1C→P+0x1d`** (range), `+0x20→P+0x28`. So the stored intensity scalars are `+0x18/+0x1C`, **not** `+0x14/+0x18`. **Amended by [EXP-0030]: every *file* offset in this row is −8** — it was written against `ALM-META-008`'s superseded anchor, and `typeId`/`f32` are record-header words, not payload. Re-anchored: `W@+0x00`, `H@+0x04`, `θ@+0x08`, `+0x0c→P+0x2c`, **ambient `@+0x10→P+0x1c`**, **range `@+0x14→P+0x1d`**, `+0x18→P+0x28`. The `P+…` runtime slot displacements are unchanged, and so are the file bytes named — this is the same read sequence re-expressed on the corrected payload base, no new fact and no confidence change. Further: the map-stored light fields are **not shown to be read** — `R0468` overwrites `P+0x20/0x24` (from `[L07948]/4ac`) and `P+0x1c/0x1d` (from `[L05650]/498`) in its first ten instructions, unconditionally, before any read; and the loader is structurally incapable of supplying θ anyway (the slot is a `double`, case 0 writes only its low 4 B and never `P+0x24`, and the following `read(P+0x1d,4)` spills over `P+0x20`'s byte 0). **Amended by [EXP-0298]: the Unknown is closed and the seventh read is not.** No path reads the stored `θ`, ambient or range values (`TERR-LIGHT-149`, `TERR-LIGHT-150`, `TERR-LIGHT-151`); the seventh destination, `P+0x28` from payload `+0x18`, is the one field of the seven with a consumer — the terrain tile-group mask `R0516` reads (`TERR-LOAD-152`)

**Confidence.** High (offsets + the overwrite) / **Unknown resolved to Medium** (no other path reads the stored values, bounded by the searches `TERR-LIGHT-149` names)

**Original status.** ● active (amended)

**Evidence.** [EXP-0025](../experiments/EXP-0025-shading-table/), [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

**Amended.** The table ledger carried the status "● active (amended)".

### TERR-EDGE-024

**Far-edge cell corner sourcing = raw flat row-major addressing (no clamp/dup/wrap/table).** All render grids are **unpadded `W×H`**: `R0480` allocates heights/type3/brightness each as `malloc(W·H)` and type1 as `malloc(W·H·2)`, back-to-back, with no `(W+1)/(H+1)` padding and no guard row/col (corpus: type1 `==2·W·H`, type2 `==W·H`, **38/38**). The four terrain renderers (`R0539/R1814/R1815/R1816`, dispatched from the per-frame draw `R0379`, msg 0x402) read each cell's four corner brightness bytes from `P+0x18` at **flat offsets** `grid[idx], grid[idx+1], grid[idx+W], grid[idx+W+1]` (`idx = col + row·W`, `W = CMapView+0x84`) and pass the values to the blitters `R1506`/`R1507`; none contains a `col==W-1`/`row==H-1` test, a `min(col+1,W-1)` clamp, a duplicate slot, or a modular wrap. Consequence: a **last-column** cell's `+1` corner reads `grid[(row+1)·W + 0]` (the **next row's column 0** — a row-major wrap, in-bounds except at the last row); a **last-row** cell's `+W` corner reads `grid[H·W + col]`, **past the `W×H` allocation**. The same flat scheme drives heights in the projection builder `R1818` and tile words in the object passes. Falsified H2 (clamp — no `min`), H3 (dup — exact `malloc(W·H)`, corpus size), torus-wrap (reads next row, not same-row opposite edge). **Note [EXP-0030]: EXP-0029's *stored-grid content* statistics were computed at the 8-byte-early grid base and have been recomputed (`evidence/edge-census.txt`); the mechanism above is unaffected (it comes from the disassembly and from the section sizes), but the recomputed edge figures differ — see the correction box in `EXP-0029.md`**

**Confidence.** High (the addressing, the allocations and the four falsified alternatives — H2 clamp, H3 duplicate, torus wrap — each excluded by a *named absence* in code read end to end, plus the corpus size identities; and the one statistic whose base moved was recomputed rather than trusted, which is the handling this ledger's other censuses did not get)

**Original status.** ● active

**Evidence.** [EXP-0029](../experiments/EXP-0029-terrain-edge-cells/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

### TERR-EDGE-025

**The per-vertex brightness OUTER RING is never computed (amends `TERR-LIGHT-013`).** The relight driver `R0532` calls `R0468(1,1,0,0)`; the x87 loop covers exactly the strict interior `{1..W-2}×{1..H-2}` (starts at col/row 1, ends at W-2/H-2), using a `±1` neighbour whose sign follows the sun azimuth so both signs stay in bounds. The outer ring of vertices (row 0, row H-1, col 0, col W-1) is **not written and has no one-sided fallback** — `TERR-LIGHT-013`'s "first row/col one-sided" wording read the decompiler's `if (iVar3 + -1 == 0)` as a border case, but both arms of that branch read row `row-1` and differ only in column direction, and its flags come from the decrement at `L10227` overwriting the zero flag of the preceding azimuth test of bit `0x1` — nothing is clamped or substituted. Nothing on the relight path fills the ring: the driver's next call `R1819` is a **per-object** brightness stamp (footprint corners → level 0 / 0x50 by object flags 0x1000/0x20000, driven by the lit-object list `+0x9f4`/`+0x9f8`), not an edge fill; and the allocator chain `R0722→R0748→R1567` **never clears the block** — it returns either a recycled small-block free-list entry (`R1568`) or `HeapAlloc(heap, dwFlags=0)`, i.e. without `HEAP_ZERO_MEMORY`. So the outer ring — and thus even the in-bounds far-corner reads of the second-to-last cells — samples memory the engine never initialised

**Confidence.** High (loop bounds + no-fill + no-zeroing, single-path instruction reads) / **Unknown** (what those bytes actually are at runtime: not zeroed by the engine, but a first-touch heap page may still arrive demand-zeroed from the OS — not observed running)

**Original status.** ● active

**Evidence.** [EXP-0029](../experiments/EXP-0029-terrain-edge-cells/)

### TERR-EDGE-026

The affected former wording below is partially retracted. See the
**Amended.** paragraph for the retained facts and correcting claim.

**The outer ring is drawn, not padding-never-drawn; its degenerate shading is tolerated by the border+camera.** No renderer skips or special-cases the outer ring: the screen-Y mesh is built for it (and for a several-cell over-scan margin beyond the grid — the draw loops run **viewport-relative** `row ∈ [-4…-3, visRows+7…+9]`, `col ∈ [-4…-3, visCols+3…+5]` in `R1818`/`R0379`, offset by the scroll origin `+0x5c/+0x60` and reading the grids flat, never tested against `W`/`H`), and blits are issued, gated only by normal screen-rect culling. The **only** bounds check in the entire draw is a single `if (worldRow < H)` guard on the smoothed-height `+0xc0` pass in `R1818` — its presence proves the draw expects `worldRow` to reach/exceed `H`. The engine tolerates the degenerate outer band (uncomputed brightness ring + wrap/OOB far corners) because gameplay is bounded by the **derived 8-cell sim border** (`R0470` marks `0x1f` 8 cells deep on the 256×256 sim grid, `ALM-TERR-016`) and the camera keeps that band at/beyond the extreme edge. **For a reimplementation the far edge has no defined shading to reproduce — clamp-to-edge is the safe, intent-matching choice.**

**Confidence.** High (not skipped; over-scan; the guard) / **Medium clause ANSWERED (EXP-0118, `SESS-VIEW-030`)**: the CMapView scroll clamp is `8 <= origin <= dim − 8 − span` on each axis, four enforcement sites, each one named instruction — so the origin never enters the outer 8 cells and the outermost ring is reachable only through this row's own 3–4 cell over-scan. The scroll origin this row locates at `+0x5c/+0x60` is confirmed and given a unit (cells) and a value at map start by `SESS-VIEW-029` / `MISSION-VIEW-019`; `visCols`/`visRows` are `+0x64`/`+0x68` and are resolution-dependent (`SESS-VIEW-028`)

**Original status.** ● active

**Evidence.** [EXP-0029](../experiments/EXP-0029-terrain-edge-cells/)

**Amended.** TERR-217 partially retracts the outermost-terrain-dispatch and guard-implies-raster clauses. Projection/drawable overscan is not the four terrain routines' admitted cell population. The raw corner arithmetic and uncomputed brightness-ring facts stand.

### TERR-GRID-027

**The tile grid every `TERR-*` render claim consumes starts 8 bytes later in the `.alm` than `ALM-GRID-011` said — a render indexing the section payload from byte 0 is four cells off in X.** `rom.exe`'s `.alm` reader allocates `W·H·2` and reads exactly `W·H·2` bytes at the stream position after a **20-byte record header whose last two words are the `typeId` and the per-map `f32`** (`ALM-FRAME-031`/`ALM-GRID-032`), so the buffer the renderers index as `grid[row·W+col]` (`TERR-IDX-003`, `TERR-EDGE-024`) begins at EXP-0007's *payload + 8* = **word `row·W+col+4`** of that payload. Nothing in the render, ingest or lighting path compensates — the offset is entirely in how we split the file. **Consequence for the port: any terrain frame built on the old base is displaced 4 cells in X** (and a heights/overlay layer read the old way is displaced 8, since those are u8 — so an old-base render is also internally 4 cells out between tiles and relief). Nothing else in `TERR-001…026` changes: the tile-word→graphic mapping, the water phase, the dirt composite, the shading table and the far-edge behaviour are all statements about a *word* or about the buffer, not about the file offset. Two corollaries: `TERR-VER-005`'s "overlay corner excluded" carve-out is no longer needed (there is no overlay corner — those 4 words are header fields), and the far-edge ring analysis of `TERR-EDGE-024…026` applies to the corrected grid unchanged

**Confidence.** High (inherited whole from `ALM-FRAME-031`/`ALM-GRID-032`, which are instruction-level plus a refutation from bytes alone)

**Original status.** ● active

**Evidence.** [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

### TERR-LIGHT-028

**The height-gradient span is the ADJACENT ONE-CELL difference — twice, along the same axis (amends `TERR-LIGHT-013`; resolves R-3).** Reading the raw x87 of `R0468` with the FPU stack tracked by hand (the decompiler drops it, which is why EXP-0023 left this Medium), the two Δh terms are `Δh₁ = H[x∓tan\|θ\|, y+1] − H[x, y]` (forward, **one** row) and `Δh₂ = H[x, y] − H[x∓tan\|θ\|, y−1]` (backward, **one** row). Consequences, each read off the listing: (a) **no two-cell central difference exists** — no instruction forms `H[…,y+1] − H[…,y−1]`, and the only `0.5` (`L10230`) multiplies the sum of the two already-**clamped levels**, never a height difference; (b) **there is no perpendicular axis** — EXP-0023's "two axes" are the forward and backward differences of the *same* row axis, and the only `x±1` reads are the lateral operands of the `tan\|θ\|` shear *within* rows `y+1`/`y−1`; (c) the azimuth enters twice, as that lateral interpolation and as `stepH = 32/cos\|θ\|`; (d) heights are read sign-extended (**signed**), unobservable on the shipped corpus (0/880 704 bytes ≥ `0x80`); (e) the two one-cell slopes are clamped **independently** before averaging, so the routine is not algebraically equal to any single-difference form. The lateral column selector of axis 2 is chosen by `y == 1`, not by the θ sign it computes and discards (the decrement at `L10227` overwrites the zero flag of the bit-`0x1` test at `L10239` — `TERR-EDGE-025`); measurable effect nil (identical corpus range either way). Falsified against synthetic fixtures: flat grids → exactly 46 at any height level; monotone sign-correct ramps; a 0/127 checkerboard → 46 everywhere, which only holds if the `tan\|θ\|=1` diagonal shear is transcribed correctly; deleting the shear moves the corpus ceiling 70→66, so no term is decorative

**Confidence.** High (instruction-level, single writer, falsified)

**Original status.** ● active

**Evidence.** [EXP-0031](../experiments/EXP-0031-terrain-gradient-span/)

### TERR-LIGHT-029

**Corpus level census, corrected — and it is a function of θ (corrects `TERR-LIGHT-018`'s `30…74`).** With the exact `TERR-LIGHT-028` span, the corrected [EXP-0030] grid base, the interior window the engine writes, and the sun model's own cycle-off configuration (θ = 0.78539815, `L = 62`, `R = 32`): levels **`[30..70]`** over the 38-map corpus (859 768 vertices) and **`[30..65]`** over the 10 root maps. The previously published **`30…74` is stale**: reproduced here bit-exactly — range *and* the 880 704 vertex count — **only** on the superseded EXP-0007 framing, with its maximum at `scn:61.alm` vertex (3,0), i.e. inside the 8 record-header bytes that split injected into the height grid; the unmodified EXP-0025 probe on the corrected base prints `[30..69]` today. Two further facts that make any bare "ceiling" figure unsafe: (i) the ceiling moves with θ — 70 at the default, 72 at midday, up to 81 near the sweep ends — and `R1813` sweeps θ over `[−π/2, +π/2]` every game day; (ii) the floor **30** is not an agreement between models but a **saturation** — `clamp(62 − 32·sin(π/6 − s), 0, 95)` over `s ∈ (−π/2, π/2)` is minimised at `s = −π/3` → exactly 30 — so any span reaches it. Analytic corollary: the level is confined to `[30, 89]` for **any** height field, so the `[0,95]` clamps never bind and no span choice could put a vertex outside the 96-row table (0/859 768 observed)

**Confidence.** High (arithmetic over the full corpus, 0 invariant violations) / Medium (that a live session runs at this θ — static observation only)

**Original status.** ● active

**Evidence.** [EXP-0031](../experiments/EXP-0031-terrain-gradient-span/)

### TERR-LIGHT-030

**The sun model's cycle-off default θ is a literal, and it is not π/4.** `R1813`'s `[L06260] == 0` arm stores `0x4d12d84a` to `[L07948]` and `0x3fe921fb` to `[L10240]` → the double `0x3fe921fb4d12d84a` = **0.78539815**, i.e. `3.1415926/4` from a truncated π literal, **not** `π/4` (`0x3fe921fb54442d18`); the same arm sets ambient `0x0e` and range `0x20`, so this is exactly the `L=62, R=32` daytime configuration. With the cycle enabled, the 6…17 band computes θ instead: the integer minutes converted to a double, multiplied by `[L07945]` and reverse-subtracted from `[L07946]` = `−π/2 + m·0.00218`, and `0x2d0 = 720` minutes × `0.00218 ≈ π`, so θ runs across the twelve daylight hours. **⚠ The span in this row is WRONG and is corrected by `TERR-LIGHT-109` (see retracted.md): `[L07946]` is −0.78539815, NOT −π/2, and `[L07945]` is −0.0021816615277777777, so `720 × step = π/2` and θ runs −π/4 → +π/4 — half the arc this row claims.** The refinement of `TERR-LIGHT-014`'s "fixed ±π/4 at some bands" stands and is completed by `TERR-LIGHT-110`: two of the five arms compute, three store a literal ±0.78539815. Because `stepH = 32/cos\|θ\|` and the lateral shear `tan\|θ\|` both depend on it, **no per-vertex level statistic is meaningful without its θ**

**Confidence.** High for the `[L06260] == 0` literal and the shape of the daytime computation (immediates read from the listing) / **the SPAN is retracted** — it was read off the listing's symbol names rather than the bytes, and `EXP-0116` dumped both doubles from the PE on both roots

**Original status.** ✖ partially retracted (span corrected by `TERR-LIGHT-109`)

**Evidence.** [EXP-0031](../experiments/EXP-0031-terrain-gradient-span/), [EXP-0116](../experiments/EXP-0116-daylight-cycle/)

**Amended.** The table ledger carried the status "✖ partially retracted (span corrected by `TERR-LIGHT-109`)". `retracted.md` records a correction against this claim.

### TERR-GEOM-031

**The altitude becomes a destination row, by subtraction — and the heights are signed at the point of use.** The per-frame mesh builder `R1818` (pass 1, `L10241…L10242`) writes `CMapView+0xb4` as **`mesh[r][c] = (r<<5) − (int8)H[worldCol + W·worldRow]`**: `L10243` shifts the screen row left by `0x5` (× 32), `L10244` loads the byte sign-extended (**`0f be`**, from the raw type2 grid at `P+0x10`, the same bytes `R0480` reads from the file), `L10245` subtracts the loaded byte from the shifted row. There is no second write to the mesh and no negation before it, so **a larger altitude yields a strictly smaller destination row index**. That index is what gets multiplied by the surface pitch (a multiplication by `[L01511]`) and added to the surface base (`[L01168]`) to form the write address, and it is bounded from below by field `+0x04` and from above by field `+0x0c` of the struct at `L01503` — which `R0372` passes to `CopyRect` as a `RECT*`, i.e. those fields are `top` and `bottom`. **Under the Win32 device-coordinate convention (`top < bottom`, y increasing downward) a larger altitude therefore moves the pixel UP the screen; that last link is a documented OS convention, not a ROM1 instruction, and everything before it is.** The *derived* smoothed grid `+0xc0` (mean of four sign-extended heights, shifted right by 2) exists but feeds a different consumer — the blit never reads it. **AMENDED by [EXP-0036]: that consumer is the *sprite* path — `+0xc0` is what lifts a unit or object off the ground, and it is the mean of the cell's four corners rather than one of them (`TERR-SPR-039`). So the frame carries two altitude models, and this row's `r*32 − h` governs only the terrain raster.** Signedness is decidable **only** from the instruction: 0/880 704 shipped height bytes are ≥ `0x80` (`TERR-LIGHT-028` (d)), so no corpus can discriminate a sign-extended from a zero-extended read

**Confidence.** High (mesh formula, sign, signedness, grid identity — single instruction each, single writer) / High (screen direction, via the `CopyRect` typing + the Win32 RECT convention — link explicitly named)

**Original status.** ● active (amended)

**Evidence.** [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-GEOM-032

**Flat vs sloped is chosen from the four projected corner Y, and is exactly "all four corner altitudes equal".** All four renderers (`R0539` @L10246, `R1814` @L10247, `R1815` @L10248, `R1816` @L10249 — byte-identical sequences) test `yTL==yTR && yBL==yBR && yTL+32==yBL` on the `+0xb4` mesh values and call `R1506` on all three, `R1507` otherwise. Because `mesh` is affine in the height with coefficient −1, the three comparisons reduce exactly to `h(c,r)==h(c+1,r)==h(c,r+1)==h(c+1,r+1)`: **not** a threshold, **not** any-difference, **not** a test on the raw grid. Reduction verified computationally, **0 disagreements over 870 198 interior cells** of the 38-map corpus. Share on that corpus: **10.28 % flat / 89.72 % sloped** — i.e. a renderer that only implements the flat path is wrong on ~9 cells in 10

**Confidence.** High (the three `CMP/JNZ` pairs, four agreeing renderers, reduction falsified against the corpus)

**Original status.** ● active

**Evidence.** [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/)

### TERR-GEOM-033

**What the two blitters actually draw.** `R1506` (8 args `x, y, c0..c3, src, LUT`) writes the axis-aligned rectangle `[x, x+32) × [y, y+32)`, source pixel `(i,j)` → destination `(x+i, y+j)` — **no altitude reaches it**. `R1507` (12 args `xL, xR, yTL, yTR, yBL, yBR, c0..c3, src, LUT`) is **column-major with a per-column vertical span**: for each of the 32 columns it walks the top edge from `yTL` toward `yTR` and the bottom edge from `yBL` toward `yBR` (`TERR-GEOM-034`), then fills `top(i) .. bottom(i)−1` — **top inclusive, bottom exclusive** (`count = bottom − top`, `L10250`) — resampling the 32 source rows with a **16.16 DDA**: `srcStep = (32<<16)/H` (`L10251…L10252`, a signed division by the span height, so it is a **stretch, not a rigid shift**), `v` starts at **0**, source row `= v>>16` (floor, no rounding), source byte `= src[i + 32·(v>>16)]` via a shift right by `0xb` and a mask `0x1fffe0`. Shading is interpolated over the same span (`levelStep = (levelBot−levelTop)/H`, a signed division, trunc toward zero; accumulator seeded `+0x100`). **`H ≤ 0` (a collapsed or inverted span) skips the column entirely** — the branch at `L10253`, taken when not positive — no clamp, no minimum, no fallback pixel. Flat is the degenerate case of the same convention (`yBL = yTL+32` ⇒ `H = 32`, `srcStep = 0x10000`)

**Confidence.** High (every step a named instruction; the stretch-vs-shift and the exclusive bound each pinned by one instruction)

**Original status.** ● active

**Evidence.** [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/)

### TERR-GEOM-034

**The edge interpolation is a table walk, and the table is a startup-built Bresenham step table.** `L10254` is `.bss` (zero in the static image) filled once by `R0788` @`L10255…L10256`: for `d = 1..127`, `s = 0x200000/(d+1)` (unsigned division), `acc = 0x8000`, and for `k = 0..d`: `acc += s; T[d][k] = acc>>16`. Row stride `0x80`, 128 rows (`L10254…L03341`); row 0 never written. The sloped blitter indexes row `d = \|Δy\|` (a shift left by `0x7` and an add of `L10254`) and compares entries against the **0-based column counter**: **forward** walk (`entry ≤ col` → `idx++`, y steps toward the far vertex) when `yTL<yTR` on the top edge or `yBL≥yBR` on the bottom, **mirrored** walk (`32−entry ≤ col`, `idx--`) in the other two cases; each edge **latches** when its index reaches `d` (or drops below 0) and is left alone for the remaining columns. `T[d][d] == 32` on all 127 rows (a sentinel the counter, max 31, can never reach) and every row is monotone. Closed form, verified exhaustively over `d∈[1,127] × col∈[0,31]`: `offset = min(d, floor((2·col+1)(d+1)/64))`, **exact except at `d = 63` and `d = 127`** where the truncation in `s` bites — so the edge is the vertex-to-vertex line sampled at **pixel centres with an inclusive-endpoint (`d+1` step) convention**, not `d` steps over 32 columns. For `d ≥ 63` the drawn edge does not reach the far vertex by column 31 (65 of the 127 rows). **AMENDED by [EXP-0036]: the three further blitters this row left unexamined (`R1508`, `R1509`, `R1510`) are the *shroud* pass over the same quad, not sprite paths — `TERR-FOG-037`. Enumerated over the whole binary, the table has exactly four readers and none of them is a sprite blitter (`TERR-SPR-038`).**

**Confidence.** High (14-instruction builder reproduced exactly; invariants hold on all 127 rows; cross-checked against a second routine, see `TERR-GEOM-036`). **Re-checked by [EXP-0040]** on a repaired function table (the vtable blind spot closed, 607 functions created): the four-reader enumeration reproduces exactly, 0 hits in orphan code — **confirmed**

**Original status.** ● active (amended)

**Evidence.** [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-GEOM-035

**The displacement is vertical only, and the horizontal clip is whole-cell.** Destination X in the sloped blitter is the address `base + 2 × column`, with `base = surfaceBase + xLeft*2 − 2` and `column` the plain column counter running `1..32` — **exactly 32 columns at `xLeft+0…31`, with no altitude-derived term anywhere in the `*2` path**; the `xRight` argument (`+0x0c`) is consumed **only** by the clip test. Both blitters reject a cell outright if `x < clip.left` or `x+32 > clip.right` (`L10257…L10258`, `L10259…L10260`) — **horizontal clipping is all-or-nothing per cell**, while vertical clipping is per-row (rows above `clip.top` are skipped with the source/level accumulators advanced in lockstep; the row count is truncated at `clip.bottom`). So there is no isometric X shear: the terrain frame is a plain 32-px column grid whose *rows* bend

**Confidence.** High (destination-address arithmetic and both clip tests read instruction by instruction)

**Original status.** ● active

**Evidence.** [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/)

### TERR-GEOM-036

**Ownership, seams, and the two places the model degenerates.** (a) **Vertical seams:** a cell's bottom edge and the cell below's top edge are the same line with the same endpoints, but the engine walks one forward and the other mirrored. Exhaustive comparison of all 127 table rows: the two agree everywhere **except `d = 63` and `d = 127`**, where the mirrored walk sits exactly 1 row lower across all 32 columns. Every such column is an **overdraw** (the upper cell paints one row the lower cell repaints, and the lower cell is drawn later — rows are drawn in ascending order), never a gap: **9280 boundary columns on the corpus, overdraw 9280 / holes 0**, classified per column rather than inferred. (b) **Collapsed spans drop a column, but are not a visible hole:** `H ≤ 0` drops the column — **3576 of 24 983 232 drawn columns (0.0143 %), in 353 cells** of the 38-map corpus. **AMENDED (render, `tools/terrgfx -mode geomart`, 2026-07-26): the dropped column leaves no unwritten destination pixel.** A whole-frame render of Kids/Islands/Tomb — 375 collapsed columns between them — has **0 unwritten interior destination pixels** (measured by painting the surface magenta first and counting what survived, inside a window inset by `maxH+64` rows so the map's own top/bottom boundary cannot contribute). Reason, structural: `H ≤ 0` means `top(i) ≥ bottom(i)`, i.e. the cell's own edges have crossed; the cell above shares the crossed cell's top edge and covers to `top(i)−1`, the cell below shares its bottom edge and covers from `bottom(i)` — so the crossed range `[bottom(i), top(i))` is painted by the neighbour, not left blank. This corrects the word "hole" only; the column really is skipped and the *source* pixels it would have carried are lost. Not yet measured on the other 35 maps. (c) **Table bound:** `\|Δh\|` between adjacent `int8` heights could reach 255 and index past the 128-row table; corpus max is exactly **127** (the last row *is* exercised) with **0** edges needing a row ≥ 128 — the engine is not robust here, the shipped data merely stays inside it. (d) **The engine contains a second, disagreeing model of the same quad:** the picker `R0377` bounds a cell by an *exact* lerp `y0 + ((y1−y0)·(x&31))/32` (trunc toward zero) of the same mesh; against the blitter's table walk it disagrees on **3665 of 4064 `(d,col)` pairs, by up to 3 rows** (worst `d=75, col=29`: raster 70, picker 67). A port must copy the **table walk** — that is what is drawn; the lerp is what the engine hit-tests against. Span heights over the corpus run **1…159** destination rows for the same 32 source rows

**Confidence.** High (a) — exhaustive over the table and the corpus, with the sign of each seam classified, not assumed / High (b) the drop itself (one `JLE`) and the 0-unwritten-pixel measurement on the three rendered maps; Medium that no shipped map anywhere has a visible hole (3 of 38 rendered) / High (c) — exhaustive over the table and the corpus / High (d) — both routines read at instruction level

**Original status.** ● active (amended)

**Evidence.** [EXP-0035](../experiments/EXP-0035-terrain-cell-geometry/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-FOG-037

**The three "further blitters" are not sprite paths — they are the shroud/fog pass over the *same* terrain quad.** `R1509(xL,xR, yTL,yTR,yBL,yBR, c0..c3)` has the sloped terrain blitter's argument list **minus the source pointer and the LUT**: it walks the same two step-table edges (`L10261`/`L10262`, adding `L10254`), fills the same per-column `[top,bottom)` span with the same signed-division-by-`H` level ramp (`L10263`) over exactly 32 columns (`L10264`, a comparison of the column counter with `0x20`), and per pixel does **`dst = LUT[level][dst]`** — `L10265` reads the 16-bit framebuffer word it is about to overwrite (`L10266`/`L10267`), through the pointer at `[L03341]`, with two row strides selected by `[L03346]` (mask `0xfffe0000`, 16-bit index; or mask `0xffffc000` with a shift right by 3, 13-bit index). `R1508` is the same quad at fixed level **`0x10`** (`L10268`, four pushes of `0x10`) and `R1510` at fixed level **`8`** (`L10269`); each adds a path for a *self-intersecting* quad (`max(yTL,yTR) ≥ min(yBL,yBR)`) that fills the four wedges directly — `R1508` with colour **0** (a zeroed value stored by a repeated dword store), `R1510` with **`(px>>1) & mask`** (`L10270` shifts right by 1, then masks with the doubled `[L07813]` channel mask). All three are called from **one** place in the binary — `R0379` @`L10271`/`L10272`/`L10273` — with `(col<<5, (col+1)*32, yTL,yTR,yBL,yBR)`, the terrain cell quad. `0x10` is the same constant the occlusion test culls on (`TERR-SPR-041`), so it is the fully-dark row. **Supersedes EXP-0035's guess "presumed other pixel formats/blend modes": same 16-bpp format, different operation, same geometry**

**Confidence.** High (argument lists, the framebuffer read, the fixed levels and the call sites all read instruction by instruction; "no source bitmap" is a property of the argument list, not an observation that one was not seen) / ~~Medium (that level `0x10` means fully dark — the two constants agree but the remap table behind `[L03341]` was not read)~~ — **the table is now read and the clause rises to High: [EXP-0088] reproduces it as 17 rows of `out = (in × (16 − L)) >> 4` and checks that law against the engine's own two table-free fast paths over the whole 16-bit pixel space, 65536/65536 exact (`TERR-FOG-084`). Also amended: the level argument is not a shroud *setting* but the per-vertex fog state, and the "same quad" is drawn only for the corner combinations `TERR-FOG-083` enumerates — for an all-visible cell these blitters are not called at all**

**Original status.** ● active (amended)

**Evidence.** [EXP-0036](../experiments/EXP-0036-sprite-geometry/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-SPR-038

**A sprite is lifted by the terrain, never warped by it.** Enumeration over the **whole 1 977 344-byte image**: the geometry step table `L10254` is touched by exactly four routines — `R1507` (terrain, sloped) and the three shroud blitters of `TERR-FOG-037` — plus `R0788`'s loop bound. **No sprite blitter indexes it.** Independently, the sprite draw call passes **two scalars**: `L10274`, a call through the vtable slot at `+0x3c`, with pushes (last-arg-first) `0, level, [L04368], frame, dstY, dstX` — there is no argument that could carry a second or third corner Y. **AMENDED by [EXP-0039]: that leading `0` is not a constant of the call, it is the not-mirrored case of a horizontal-mirror flag** — the unit path pushes the flag it builds from `Flip` into the same slot (`L07461`, `L07462`; `REG-UNITS-051`, `TERR-SPR-048`). Nothing about the object path changes; a consumer implementing `vt+0x3c` from this row alone would hard-code the parameter away. **AMENDED AGAIN by [EXP-0053]: the slot this push list names `level` is the SUN SHEAR** — a 16.16 per-row X slope the blitter accumulates, not a brightness; the shroud level is the **fourth** argument, not the fifth (`TERR-SPR-066`, retracted.md). Read as written, the list feeds a brightness into a geometry parameter and the shadow does not lean. The destination itself is `L10275…L10276`: **`dstX = col*32 + 16 − anchorX`**, **`dstY = row*32 + 16 − anchorY − alt`** (a shift left by 5, an add of `0x10` and two subtractions), i.e. the cell **centre**, with the altitude **subtracted** — the same sign as the terrain mesh (`TERR-GEOM-031`), so higher ground moves a unit up the screen by the same rule. On the shadow pass only, `dstX` loses one further term, a sun-derived shear (`L10277`, from `R1820`→`R0812`→`__ftol`); the body pass subtracts nothing there. Every shift left by 5 followed by an add of `0x10` in the image (byte scan, 19 hits) lies in `R0379` or `R1401`, except one pathfinding grid index in `R1532` — **the cell-centre convention has exactly one definition**. **Rendered 2026-07-26** (`tools/terrgfx -mode sprart`, render): plotting `(col*32+16, row*32+16 − alt)` against the terrain raster over three maps puts 98.81 % of contact points inside the drawn span of the very cell they stand on and **none** in empty sky — see `TERR-SPR-039` for what that measurement does and does not discriminate

**Confidence.** High (the step-table absence is an enumeration over every instruction in the binary, which is the only evidence that can carry an absence; the destination arithmetic and the push order are named instructions). **Re-checked by [EXP-0040]** on a repaired function table (the vtable blind spot closed, 607 functions created): the four readers are unchanged and **0** of them lie in orphan code, and the cell-centre idiom is still exactly 19 hits (`R0379` ×14, `R1401` ×4, `R1532` ×1) — **confirmed, no change**

**Original status.** ● active (amended)

**Evidence.** [EXP-0036](../experiments/EXP-0036-sprite-geometry/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-SPR-039

**The altitude a sprite is lifted by is a *different* grid from the one the terrain is drawn on.** It is `CMapView+0xc0`, built by `R1818` **pass 3** @`L10278…L10279` as the mean of the cell's four corner heights: four sign-extended byte reads of the raw type2 grid at `idx`, `idx+1`, `idx+W`, `idx+W+1` where `idx = (c+scrollX−1) + W·(r+scrollY−1)`, summed, then a sign-extension bias (masked with `0x3`, added) and a shift right by `0x2` — **divide by 4, truncated toward zero** — stored at `(c+3) + (r+3)·(visibleCols+8)`. The draw path reads it at `(col+4) + (row+4)·(visibleCols+8)` (`L10280…L10281`), which is the **same cell** the sprite stands on: the `+1` in the index difference cancels the `−1` in the world-coordinate base. So the engine carries **two altitude models per frame** — `+0xb4` (one corner, `r*32 − h`, drawn terrain) and `+0xc0` (the mean of four corners, sprite lift). **AMENDED 2026-07-26 by the render: this row said they "disagree by construction on any sloped cell", and that is false as a universal.** Slope says the four corners are unequal; agreement says their mean equals the cell's own corner, and neither implies the other. Over the 136 291 interior cells of Kids/Islands/Tomb (flat 1162 + sloped 135 129) the two grids **agree exactly on 12 513 sloped cells — 9.26 % of them** — and **6227** of those have four *pairwise-distinct* corners, so it is not a near-flat artefact. The converse half does hold: **0** flat cells disagree, as four equal corners must. Where they differ, `mean − corner` runs **−73 … +55** rows, with `\|Δ\| > 2` on 55.52 % and `\|Δ\| > 4` on 30.55 % of cells. Readers of each, enumerated over the binary: `+0xb4` = the four terrain renderers, the builder, the shroud pass, the picker; `+0xc0` = the builder and `R0379` (8 sites). **No sprite path reads `+0xb4`.** This identifies the consumer `TERR-GEOM-031` said existed and did not name

**Confidence.** High (the build is a single 30-instruction block with every read `MOVSX` and the divisor an immediate; the index algebra is reproduced on both the write and the read side; **re-checked by [EXP-0040]: `disp:b4` has 0 hits in orphan code and, decisively, **no hit anywhere in the CUnit block `L05122..L05123`** — the very population that hid the unit `Draw` contains no reader of `+0xb4`, so "no sprite path reads it" survives a sweep that can see orphan code. The four `+0xc0` hits in that block are `class+0xc0` on a `units.reg` record reached via `[L02113][unit+0x20]`, a different structure sharing the displacement, not `CMapView+0xc0`** — both reader sets are enumerations over every instruction) / High (the agree/disagree census — a universal is refuted by counterexamples, and there are 12 513) / **Medium** (that the resulting lift stands a sprite on the ground the *other* grid draws: **134 666 / 136 291** contact points land inside the destination-row span their own cell's raster fills at screen column `col*32+16`, **0** land where nothing is drawn, none above their own span, ≤ **2 px** below it, and the point sits at that span's vertical centre to +0.85 px mean. That is a corpus cross-check between two separately transcribed code paths, so Medium is its ceiling, and it cannot catch an error the two share. It **discriminates**: the corner grid (10 388 outside, up to 48 px), a one-cell index-base slip either way (≈16 000 outside, up to 61 px), and no lift at all (668 contacts in empty sky). It **cannot discriminate** the rounding rule — `round4` scores *better*, 135 673 — nor the trunc-toward-zero fixup, since `SAR 2` without it is **bit-identical** on all 136 291 cells because no shipped height byte is ≥ `0x80` (`TERR-LIGHT-028`(d)); both are settled by the listing alone)

**Original status.** ● active (amended)

**Evidence.** [EXP-0036](../experiments/EXP-0036-sprite-geometry/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-SPR-040

**The anchor: the class canvas is centred on the cell centre, and `(CenterX, CenterY)` is the pixel that touches the ground.** `L10282…L10283` and `L10284…L10285`: `anchorX = (class[+0x20] − class[+0x18]/2) + frameWidth/2`, `anchorY = (class[+0x24] − class[+0x1c]/2) + frameHeight/2`, all three divisions by 2 done as a sign-extension, a subtraction and a shift right by 1 (truncate toward zero). `class[+0x18]/[+0x1c]/[+0x20]/[+0x24]` are the class record's `Width`/`Height`/`CenterX`/`CenterY` — the `objects.reg` keys of `REG-OBJ-020`/`REG-OBJ-039`, read from the loaded class table `[L02099]`, **not** re-read from the registry. `frameWidth`/`frameHeight` come from the sprite object itself, calls through the vtable slots at `+0x20` / `+0x24` with frame index 0. **AMENDED by [EXP-0038]: "with frame index 0" is true of the block this row read and false of the pass that draws the object.** The cell issues four draws and builds this anchor twice — the two **shadow** draws push `0x0` (`L07845`/`L07846`, `L07847`/`L07848`), the two **body** draws push the frame index actually being drawn (`L07849`/`L07850`, `L07851`/`L07852`). The formula is unchanged; which frame's size enters it is not, and it matters because frames of one sheet need not be the same size — `TERR-SPR-043`, `SPR256-FRAME-023`. Since `dst` is the frame's **top-left**, substituting `frameW = Width`, `frameH = Height` gives: frame pixel `(CenterX, CenterY)` lands exactly on `(col*32+16, row*32+16 − alt)`. A frame of a different size is centred inside the `Width × Height` canvas first. Ruled out by the same instructions: the anchor is not the frame's top-left (both `Center*` and `*/2` appear) and not the canvas centre (`Center*` is not `Width/2`). **Rendered 2026-07-26** in the render: `units.reg` `Unit0` (swordsman, canvas 128×128, `Center (64,78)`) resolved through `Files[class.File]` (`REG-VAL-029`), frame 0 decoded and blitted at `dst − anchor` on real placed-unit cells — its `(CenterX,CenterY)` pixel lands on the ground-contact point, at the figure's feet. **That render is a *display*, not a check**: every term in it is this row transcribed, and nothing in it could have come out otherwise. What it settles is only that the deferral which said this needed a rendered unit on sloped ground has been paid

**Confidence.** High (six named instructions per axis; the alternatives are excluded by which fields appear, not by a fit) — the render adds nothing to that and is marked a display; the frame index the two `CALL`s take is corrected by [EXP-0038] and was never the discriminating part

**Original status.** ● active (amended)

**Evidence.** [EXP-0036](../experiments/EXP-0036-sprite-geometry/), [EXP-0038](../experiments/EXP-0038-object-frame/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-SPR-041

**What is settled for units, and what is not.** The per-cell drawable dispatch `L07855…L07856` loads the **same** smoothed altitude at the unit's own cell (`[this+0xc0]`, index `(col+4)+(row+4)·(cols+8)`) and pushes it as the third argument of `drawable->Draw(col, row, alt)` — a call through the vtable slot at `+0x30`. There is **no other altitude in that call**, and the object at that cell takes its class from the **unit** class table `[L02113]`. Also confirmed on the object path: the occlusion test `R1401(classId, col, row, alt)` recomputes `dstY` identically and feeds `pick(x, dstY)` / `pick(x, dstY + frameHeight)` to `R0377` as the low/high ends of a row range, culling the draw iff every covered cell's shroud level equals `0x10`. ~~**Unknown:** what a unit's `vt+0x30` implementation does with those three arguments — it was not located, and since the cell-centre idiom has exactly one definition in the binary (`TERR-SPR-038`) it is not inside such a method.~~ **AMENDED by [EXP-0039]. The Unknown is closed, and this row's identification of the dispatch is corrected: `vt+0x30` is not the sprite draw.** `R0379` issues *two* per-cell dispatches with the identical `(col, row, alt)` list; the one that draws the sprite is `L02993`, a call through the vtable slot at `+0x2c`, over `CMapView+0x90`, and the `vt+0x30` pass this row read runs later over all three grids and, for a unit, is `R1106` — the HP/mana bars. The unit `Draw` body is `R0553`, and it **ignores** all three arguments; the "second placement model" below is not a rival, it is the real one — `TERR-SPR-047`, `TERR-SPR-048`. It was never located because Ghidra had created no function over `L05122..L05123`: the CUnit methods are reached only through the vtable at `L02468`. A **second placement model** exists and is read but not resolved: `L10286…L10287` builds a unit's screen rectangle as `OffsetRect(class[+0x84…+0x90], obj[+0x60] − class[+0x34], obj[+0x64] − class[+0x38] − obj[+0x68] − obj[+0x10])` — a different origin (the unit's own pixel position), a different anchor field, and two further subtracted terms whose meaning (`obj[+0x10]` plausibly a cached altitude, `obj[+0x68]` plausibly a fly/jump height) is **not established** — **and also [EXP-0053] finds those two guesses the wrong way round**: `+0x68` is subtracted by all four CUnit drawing routines including the ground-lying shadow, while `+0x10` is subtracted by three of the four and by the shadow never, so `+0x68` is the ground-plane term and `+0x10` the lift. That also makes this `OffsetRect` rectangle the **body's** bounds, not the shadow's, which is what a picking rectangle should be (`TERR-SPR-067`(d)). That rect goes to `R1821`, not to a blitter

**Confidence.** High (the dispatch, its argument, the class table and the occlusion test are named instructions; the `OffsetRect` model is read at instruction level) / **Medium** (that the `vt+0x30` dispatch is the drawable's *sprite* draw — it is not, [EXP-0039]; the row named a real dispatch and mis-assigned its job, and nothing it measured about the argument depended on that)

**Original status.** ● active (amended)

**Evidence.** [EXP-0036](../experiments/EXP-0036-sprite-geometry/), [EXP-0039](../experiments/EXP-0039-unit-frame/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-SPR-042

**What reaches the sprite draw call's `frame` argument — three arms and two gates, none of which is "`Index`, always".** The stack slot `L10288` pushes as argument 3 of the call through the vtable slot at `+0x3c` (`TERR-SPR-038`) has exactly four writers, all in `R0379`'s type-3 pass, and they resolve in this order. **(a) Plain arm — `frame = Index` verbatim**: `L10289` reads `[fffffdcc]` and `L10290` stores it to `[fffffdc4]`, where `[fffffdcc]` was loaded at `L10291` from the class record's `+0x10`. **(b) Animated arm — `frame = Index + T[phase]`**: `L10292` loads the class's own int32 array from `+0x2c`, indexed by `phase`, `L10293` adds the element at `4 × phase`, `L02613` stores it; `phase` is the signed division at `L10294` by `[fffffdbc]` — a **signed** remainder of `(animCtr + worldCol + worldRow·worldCol)`, i.e. `animCtr + col·(row+1)`, by the array's length at `class+0x3c`, with `animCtr = CMapView+0xa70` (the same counter the water phase uses, unshifted here). Entered only when `class+0x3c != 0` (`L10295`) **and** the cell's four corner tile words OR to `0xc000` in bits 15..14 (`L10296`, gate built at `L10297..L10298`) — see `TERR-TILE-044`. **(c) Dead-object arm — `frame = 0` and the sprite is swapped**: when the cell's own tile word has bit 13 set and `class+0x44 != -1`, `L07436..L10299` (and the identical block at `L10300..L10301`) replace `File` with `classes[DeadObject].File`, set the `Index` slot to 0 and the frame slot to 0. **(d) The global override**: `L10224` compares `[L05659]` with `0x0` and `L10225` stores `0` to `[fffffdc4]` — when animations are disabled (`-noanimation` / `-detail0`, `TERR-ANIM-009`'s flag, static default 1) **every** static object draws frame 0 regardless of its `Index`. All four draws of a cell push this one slot. Corpus, 82 classes on the corrected `.reg` framing: `Index ∈ 0..8`, and over the **75** classes whose sprite ships, **0** put `Index` outside the sheet's frame count and **0** put any `Index + T[k]` outside it (the other 7 have no frame count to be outside of — `ALM-CLS-042`). The rival "the frame is always 0" is refuted by art, not by arithmetic: nine classes share `stones/sprites.256` and their `DescText` runs *Big stone 1..3 / Rock 1..3 / Small Stone 1..3* against frames 0..8, which the render renders

**Confidence.** High (each arm, each gate and the modulus is a named instruction in a function read end to end, and the slot's writer set is exhaustive) / Medium (the corpus in-range figures — 82 classes, one registry) / **Unknown** (whether arm (b) is ever entered at runtime — `TERR-TILE-044`)

**Original status.** ● active (figure corrected)

**Evidence.** [EXP-0038](../experiments/EXP-0038-object-frame/)

**Amended.** The table ledger carried the status "● active (figure corrected)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-SPR-043

**The anchor's frame index is not always 0 — the body pass uses the drawn frame, the shadow pass uses frame 0 (amends `TERR-SPR-040`).** One type-3 cell issues up to four draws, and they build the `TERR-SPR-040` anchor **twice**, with different frame arguments to the calls through slots `+0x20`/`+0x24`: passes 1 and 2 push `0x0` (`L07845`, `L07846`; `L07847`, `L07848`) and subtract the sun shear from `dstX`, i.e. they are the **shadow** draws (`vt+0x3c` at `L10274`/`L10302`, gated by `[L06416]`); passes 3 and 4 push the frame slot itself (`L07849`, `L07850`; `L07851`, `L07852`) and subtract no shear, i.e. they are the **body** draws (`vt+0x14` at `L10303`, `vt+0x34` at `L10304`). The odd/even split is the two sprite objects at `spriteTable[File]+0x04` and `+0x08` — every object sheet ships as `sprites.256` **and** `spritesb.256` — with `[L04367]` gating the `+0x08` pair. Consequence, since frames of one sheet need not be the same size (`SPR256-FRAME-023`): a class whose `Index` selects a differently-sized frame has its **shadow displaced from its body** by `((frame0W − frameW)/2, (frame0H − frameH)/2)`. Eight shipped classes do: `Rock 1`…`SmallStone 3` (frame 0 `64×80`, drawn `64×64`, 8 rows) and `Pointer 2`/`Pointer 5` (`16×32` vs `24×24`, −4 columns and 4 rows). **`TERR-SPR-040`'s formula is unchanged and its "frame index 0" is true of the block it read** — the shadow pass — and false of the body

**Confidence.** High (both anchor blocks and both frame arguments are named instructions; the shadow/body split is decided by which `dstX` subtracts the sun shear, also named) / Medium (that `[L04367]` selects the `spritesb` variant — the pairing is exact over the shipped tree but no instruction names the file)

**Original status.** ● active

**Evidence.** [EXP-0038](../experiments/EXP-0038-object-frame/)

### TERR-TILE-044

**The tile word's top three bits: bit 13 is mutable at runtime and shared by three consumers; bits 15..14 gate a whole second draw path and are 0 on every shipped cell.** **Bit 13** (`0x2000`, `TERR-DIRT-017`'s impassable flag) is read at three independent places with one consistent meaning — the dirt composite in the terrain pass, the `DeadObject` substitution in the object pass (`TERR-SPR-042`(c)), and *exclusion* from `R0483`'s per-cell registry of cells holding a still-standing destructible object (`L05660`: a non-zero `& 0x2000` skips; `L05662`: `+0x44` against `-1`). It is **written at runtime**: `R1656` is a run-length decoder that rewrites it across the whole interior from `(8,8)` — the same 8-cell sim border as `ALM-TERR-016` — via `L09239` (`& 0xdfff`, OR in the value, 16-bit store) with the value normalised to exactly `0x2000` and alternated by `L10305` (XOR with `0x2000`). So a census of bit 13 in a `.alm` describes the initial state only (7 464 of 880 704 cells over 38 maps). **Bits 15..14** are never part of the strip index (`g = (w & 0x1fff) >> 6`) and are read as the four-corner OR `(w[idx] \| w[idx+1] \| w[idx+W] \| w[idx+W+1]) & 0xc000` — **but not *only* as `== 0xc000`: [EXP-0040] finds `R1504` testing for `0x8000` (`L07843`) *before* testing for `0xc000` (`L07844`), so the pair carries three distinguished states, not two. This row's "read only as … `== 0xc000`" clause is retracted.** That test gates `TERR-SPR-042`'s animated arm, and in `R1814` — one of the four terrain renderers of `TERR-EDGE-024` — it decides whether a cell is **drawn at all** (`L10306`: `== 0xc000` draws, else next cell); that renderer's only call site (`L10307`) is guarded by the animation flag and bracketed by an incremental scroll blit (`R1822`), i.e. it is the partial-repaint pass. The same idiom occurs in `R1657`, `R1823`, `R0218`, `R1307`, `R0310` — **and, [EXP-0039] adds, in `R1308` (`L10308`), CUnit's `vt+0x34`, on the cell under the unit's own `/256` position. That site was in this row's `imm:c000` scan output all along but not in its prose, because Ghidra had created no function there to name it** — the same blind spot that hid the unit `Draw`. A reader list taken from function names is only as complete as the function table. **[EXP-0040] re-ran the enumeration on a repaired function table (`tools/ghidra/RepairVtables.java`, 607 functions created) and adds three more, each running the identical four-word OR, `& 0xc000` and `== 0xc000`: `R0378` (`L10309`), `R1504` (`L10310`) and `R1505` (`L10311`). The first two were orphan code — invisible to any by-name sweep until the repair; the third had a function all along and was simply missed, which is why the corrective rule in `AGENTS.md` has a tooling half *and* a writing half. All three are `__thiscall` methods on a drawable reaching the map object via `this+0xe0`, and `R0378`/`R1505` store the gate's flag to `this+0x78` — the per-drawable visibility state `TERR-SPR-048` names as the unit sprite pass's guard. `+0x78` has several writers (`R1504` sets it to 0 from a bit-3 test on the `map+0x9b4` flag grid at `L10312` and returns before the tile-word arm), so the 0/880 704 census stays consistent with units drawing normally — but bits 15..14 now have a **second** consumer, which raises rather than answers the standing "who sets them" question.** Corpus: **0 of 880 704 shipped cells set either bit**, so 0 cells pass the gate and, from map data alone, no object ever animates and the partial-repaint renderer paints nothing. **Unknown: what sets them.** No instruction in the image ORs an immediate `0x4000`/`0x8000` into a memory operand — **re-run by [EXP-0040] on the repaired function table: the only two immediate-OR-into-memory hits image-wide are `L10313` and `L10314`, both dword ORs of `0x8000000x` in the statically-linked CRT, neither a 16-bit tile word** — and the one tile-word writer above preserves both, so either a writer exists that this enumeration's shape cannot see (a non-immediate store), or the path is dead. **⚠ ANSWERED by [EXP-0084], and the dichotomy was false: the writer is an *immediate* store, byte-wide.** `R0554` = `CUnit`/`CAirUnit` `vt+0x48` — reached once per tick per drawable through `vt+0x44`'s position test (`L02579`) and for **every** drawable every 32 animation ticks from `R0374` (`L02577`) — walks a 41×41 window around the drawable, masked by a dword table at `CMapView+0x17cc`, and executes four byte-ORs of `0xc0` into tile words (index `2 × cell`) per admitted cell at `idx`, `idx+1`, `idx+W+1`, `idx+W` (`L02581`, `L02582`, `L02583`, `L02584`) — **the gate's own four corners**. The immediate is `0xc0` on the word's high byte, so a sweep for `0x4000`/`0x8000`/`0xc000` as 16-bit operands could not see it: the miss was the operand *width*, not a missing instruction, and it is `AGENTS.md` enumeration rule 4 in mirror image — the field is a bit pair inside a `u16` and the writer addresses it as a byte. The clearer is bounded the same round: `imm:bfff` returns **2 hits image-wide**, of which `L01659` (16-bit `& 0xbfff`) in `R0374` clears **bit 14 only** across every cell of the map every 32 ticks, the other being an unrelated dword mask. That gives this row's own three-state finding a mechanism: bit 15 is set by the stamp and cleared by nothing, bit 14 is cleared periodically and re-set, so `R1504`'s `== 0x8000` before `== 0xc000` separates *stamped at some point* from *stamped this cycle*. **What this does NOT establish**: that the gate is ever satisfied in play. The stamp is guarded by `[[mapView+0x9b4]+0x38][player·2] & 8` and by `unit+0x102 != 0`, both unread, and the `CMapView+0x17cc` mask's contents are unread — so “the animated arm and the partial-repaint renderer are dead” is **withdrawn as unsupported**, not replaced by “they fire”. That is the successor experiment. **⚠ ANSWERED by [EXP-0085], and the answer is “they fire”:** the two guards are read (plus a *third* nobody had seen, `R0554`: byte `+0x15a` against `0x3`, not-below branch, the decay stage) and the mask turns out not to be a shape at all but a **per-drawable line-of-sight field** the map view `memset`s and rebuilds on every stamp (`TERR-FOG-080`). Because the stamp's immediate is `0xc0` — **both** bits into one byte — a single stamped word satisfies the four-corner OR, so the gate holds on exactly the cells the local player currently sees. Bits 15..14 are the **fog of war**, three-state, which is also why `R1504` tests `0x8000` first (`TERR-TILE-079`). This row's `0 of 880 704` is a fact about the **file**; the runtime grid is a different object and the census does not bound it. A consumer implementing only the plain arm reproduces every shipped map **at load time and never afterwards**

**Confidence.** High (bit 13's three consumers, the RLE writer's mask and the four-corner gate are each named instructions; the 0/880 704 census is exhaustive; **and the reader enumeration is re-run by [EXP-0040] on a repaired function table — the shape that made this row's list incomplete twice is now excluded, not merely noted**) / Medium (that `R1814` is *the* partial-repaint pass — read from its skip test and its call site, no second caller exists, but no name or comment attests it) / ~~**Unknown** (whether the animated arm and that renderer are reachable **in play**)~~ — **CLOSED to High by [EXP-0085]**: the guards are three, each a named instruction; the mask is a per-drawable line-of-sight field with one builder (`disp:17cc` 4 hits / 2 owners / 0 orphan) and one seed (`disp:24ec` 1 hit image-wide); and the gate holds wherever the local view has sight, because the stamp's single `0xc0` puts both bits in one word. Still **Medium** there: what sets the per-player bit 3 for a player other than the local one

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-SPR-047

**What reaches a unit's `frame` argument: nine animation states, six arms, and a default that draws the class id.** `R0553` switches on **`unit+0x74`** (`L10315`, bounded by a comparison with `0x8` (branch when above) then a jump through the table at `L10316` indexed by `4 × state`, a nine-entry table). The arms and what each computes — all bases per `SPR256-UNIT-024`, all reading the class at `[L02113][unit+0x20]`: **state 0** — if `IdlePhases != 0`, the idle block at `dir*IdlePhases + idleTimeline[unit+0x70]` (**no modulo**, `L02559`); else if `unit+0x15a == 1` the **last** dying frame and if `>= 2` a bone frame, both out of `classes[Dying]`'s sheet (`REG-UNITS-050`); else the standing frame. **State 1** — move: `MoveBeginPhases + moveTimeline[unit+0x70 mod class+0x50]`, the **only** arm that takes a modulus (`L02498`, a signed division by the value at `+0x50`, the built timeline's own length). **States 3, 7 and 8** — one shared attack arm (the game's own command strings distinguish `'Attack'`, `'Shoot'` and `'Cast'`). **State 5** — the standing frame, `frame = dir16` verbatim. **State 6** — dying, `dir*DyingPhases + unit+0x70/2` (`L10317`, a shift right by 1), from `classes[Dying]`. **States 2, 4 and anything above 8** fall to `L10318`, a load of the stack slot at offset `0x34`, written once at `L10319` from `+0x20` of the unit and never again — so **the frame drawn is the unit's own class id**, which is not a frame index by any reading. Consequently the `MoveBeginPhases` sub-block the sheet reserves is addressed by **no** arm. Corpus: the three built timelines index inside their own direction slot on 32 of 34 classes; Goblin and Goblin slinger both carry `MoveAnimFrame = 0..9` against `MovePhases = 8`, so two of their twenty move phases land on the next direction's first frames — the same pair whose `MovePhases`-vs-track-length disagreement `REG-UNITS-018` already records. **Amended by [EXP-0053]: this switch is DUPLICATED, not shared.** The routine read here is the **shadow** (`TERR-SPR-065`); the body `R0552` has its own nine-arm switch on the same `unit+0x74` with its own jump table at `L10320`, the same arm shape, and — arm by arm — the same arithmetic, except that its **default arm draws its own third argument (the light level) rather than the class id** (`L10321` vs `L10318`). So everything this row says about frame selection is true of both passes, and the corpus figures need no re-run; what it describes is the *shadow's* copy of the computation (`TERR-SPR-067`(c)). **Answers the Unknown `TERR-SPR-041` left**, on a different vtable slot than that row assumed — `TERR-SPR-048`

**Confidence.** High (each arm, the jump table, the modulus, the sign of the halving and the default's source slot are named instructions in a function read end to end — `evidence/rom-unit-frame-excerpt.md` §3) / Medium (the state→animation *names*: they are inferred from which block each arm addresses, and no string sits next to the jump table) / **Unknown** (whether states 2 and 4 ever reach a draw)

**Original status.** ● active (amended)

**Evidence.** [EXP-0039](../experiments/EXP-0039-unit-frame/)

**Amended.** The table ledger carried the status "● active (amended)".

### TERR-SPR-048

**Where a unit is drawn, and by which dispatch — the unit sprite pass is `vt+0x2c`, not `vt+0x30`.** **⚠ That headline is RETRACTED by [EXP-0053]: `vt+0x2c` is the unit's SHADOW.** The body is `vt+0x28` = `R0552`, dispatched at `L02996` over the same grid under the same guard (`TERR-SPR-065`, retracted.md). Everything measured below was measured correctly and is re-labelled as the *shadow's* dispatch, frame, destination and mirror. `R0379` issues two per-cell dispatches with the identical `(col, row, alt)` argument list, and `TERR-SPR-041` named the wrong one. The **sprite** pass is `L02993`, a call through the vtable slot at `+0x2c`, over the drawable grid `CMapView+0x90`, entered only when `obj+0x78 == 0` (`L10322`); the pass at `L07856`/`L10323`/`L10324` calls `vt+0x30` over all three grids (`+0x94`, `+0x8c`, `+0x90`) under `obj+0x7c != 0`, and CUnit's `vt+0x30` is `R1106`, whose entire body draws the **HP and mana bars** from `unit+0xfc`/`+0x100` and `+0x13c`/`+0x13e` through `R1824`. CUnit's `vt+0x2c` is `R0553` — vtable **`L02468`**, written by the constructor at `L10325` (a store of `L02468` into the object's first dword); `CAirUnit`'s vtable `L02585` names the same routine at the same slot. **The three arguments are never read**: the routine returns releasing `0xc` bytes over a `0x28 + 0x10` frame, so they sit at frame `+0x3c/+0x40/+0x44`, and the only two stack-offset-`0x3c` references (`L10326`, `L10327`) are five pushes deep and resolve to frame `+0x28`, a sprite vtable stashed at `L10328`/`L10329`. **⚠ The destination this row publishes is the SHADOW's — see `TERR-SPR-067`, and `retracted.md`. The body's differs by one term.** As read here (correctly, but of `R0553`): **`dstX = unit+0x60 − (frameW/2 + (CenterX − Width/2))`**, **`dstY = unit+0x64 − (frameH/2 + (CenterY − Height/2)) − unit+0x68`** (`L10330..L10331`, `L10332..L10333`) — to which the shadow adds a `− sunShear` on `dstX` (`L07951`/`L10334`) that this row did not record, and the **body** adds a `− unit+0x10` on `dstY` (`L10335`/`L10336`) that the shadow does not have — `TERR-SPR-040`'s anchor exactly, with the cell centre replaced by the unit's own position and `frameW`/`frameH` taken at the **drawn** frame (the calls through slots `+0x20`/`+0x24` with the drawn frame index), which is `TERR-SPR-043`'s body-pass convention. The sprite is `[L07265][class.File]` — the `units.reg` `Files` table, **not** the object table `[L10337]` — lazily loaded through `R1825` when `entry+0x10 == 0`, and the same entry carries two sprite objects at `+0x04` and `+0x08` with the second gated by `[L04367]`, the identical pairing `TERR-SPR-043` found for objects. **Amends `TERR-SPR-038`**: the blit's **sixth** argument is not always `0` — the unit path passes the horizontal-mirror flag there (`REG-UNITS-051`), and the object path's literal `0` is the not-mirrored case of the same parameter. Two blit entry points are used, `vt+0x3c` and, when the sun scalar computed at `L10338..L07952` is nonzero, `vt+0x1c` with `dstX` shifted by `scalar·k/2000` (`L10339` loads the constant `0x10624dd3`; `L10340` shifts right by `0x7`). Separately: `unit+0x08`/`+0x0c` are the unit's map position in **1/256 cell** units (`R1308` shifts both right by 8 to index the tile-word grid) and the draw path reads **neither** — so whether sub-cell position reaches the screen turns on how `+0x60`/`+0x64` are maintained, which is not established here

**Confidence.** High (both dispatches, the guards, the vtable and its constructor, the anchor terms, the `Files` indirection and the push order are named instructions; "the arguments are unread" is a frame-offset argument over the two references that could have been reads) / ~~**Unknown** (what the blit's fifth argument is for — the unit path passes the sun scalar, `TERR-SPR-038` reads the object path's as a shroud level, and the two were not reconciled)~~ — **resolved by `TERR-SPR-066`**: both paths pass the same thing, the sun shear; the object path's name for it was wrong and there was never a contradiction

**Original status.** ✖ partially retracted

**Evidence.** [EXP-0039](../experiments/EXP-0039-unit-frame/)

**Amended.** The table ledger carried the status "✖ partially retracted". `retracted.md` records a correction against this claim. The clause that all three status-bar grids require selection is also partially retracted by TERR-225: grids +0x8c and +0x90 additionally admit unselected, visible objects in show-health mode.

### TERR-PASS-049

**The map→sim ingest: three 256-stride byte planes, and the five arms that set the block byte.** `R0116` (the sim world constructor; object size `0xa4558`, `malloc`'d at `L00237`) clears four `0x10000`-byte regions — `world+0x00000` filled with **`0x01`** (`L10341` loads the constant `0x1010101`), `world+0x10000`, `world+0x20000` and `world+0x9451c` with 0 — then `R0278` walks the map's own `W×H` and writes, for `.alm` index `row·W+col`, into plane index **`(row<<8)\|col`** (the index being `col + row·0x100`, masked with `0xffff`; the 256 stride is fixed regardless of `W,H`): (a) the tile word tested against `0x2000` → store of the byte `1` at `+0x10000` (tile-word bit 13); (b) `CALL R0469(w & 0x3ff, &cost)`, then a comparison of the class with `8` → store of the byte `1` at `+0x10000` (terrain class 8); (c) a store of the cost byte, unconditionally, for **every** cell including blocked ones, `= 0xff` when the class is `≥ 13`; (d) a comparison with `0x200` after masking with `0x300` → store of the byte `1` at `+0x10000` (the water range, tested on the raw word, *not* through the classifier); (e) a non-zero test of the type-3 byte → store of the byte `5` at `+0x10000` from the type-3 plane at `map+0x10`; (f) a store of the byte at `+0x9451c` from the type-2 plane at `map+0x14`. All four block writes are **assignments**, so (e) wins outright. `R0470` then stamps `0x1f` into an **8-cell border** on all four edges (literal comparison with `0x8`; two loop shapes, for `W==H` and `W!=H`, the second also stamping the *array* edge rows/cols 0..7 and 248..255), and the ingest ends by copying the whole `+0x10000` plane to `+0x20000` (a repeated dword copy, `0x4000` dwords). Corpus, 38 maps / **880 704** cells: the block plane takes exactly four values — `0x00` 449 806, `0x01` 201 806, `0x05` 70 116, `0x1f` 158 976 — the `0x1f` count equalling the closed form `Σ W·H − (W−16)(H−16)` **to the cell**, and the cost byte never leaves `6..16` (`0xff` occurs 0 times)

**Confidence.** High (every arm is a named instruction with its immediate; the border depth is a literal and its count matches the closed form at difference 0; the plane stride is `EBP`'s own `0x100`; corpus 880 704 cells at 0 violations — and the MOV-vs-OR rival is separated by the instructions, not by the corpus, which cannot see it since `1\|5 == 5`) / Medium (that `map+0x14` is type-2 and `map+0x10` type-3: read off `R0478`'s case arms and cross-checked by swapping them, which makes 94.8 % of every map impassable — the *naming* still rests on `ALM-GRID-013`'s loader error string)

**Original status.** ● active

**Evidence.** [EXP-0041](../experiments/EXP-0041-terrain-passability/)

### TERR-PASS-050

**The terrain classifier `R0469`: which tile-word bits reach movement, and the 5-level cost blend.** Input is the tile word masked `& 0x3ff` — **bits 10–12 and 14–15 never reach this path at all** (0 of 880 704 shipped cells set any of them; this re-runs `ALM-GRID-012`'s stale census on the corrected EXP-0030 base). The split is read off the instructions and independently matches `TERR-IDX-003`'s renderer split: sub-cell `s = i & 0xf`, blend column `b = (i>>4) & 3`, strip group `g = (i>>6) & 0xf`. **Water subcell guard (amended):** subcell≥8 returns byte0xff while leaving cost output8 (TERR-WATERBOUND-164). Below 8, **water is taken first and taken raw**: `if ((i & 0x300) == 0x200)` returns class **9** — or class 1 for the single combination `s == 4 && (i & 0x30) == 0x10`, which **0** shipped cells hit — and leaves the cost at the function's entry literal **8**, so `CostWater` is never consumed on a water cell. Otherwise `s ≥ 14` rejects (class `0xff`); else the group's `(primary, secondary)` pair comes from `world+0x54156+2g` and the level `sel = world+0x540d6 + 16b + s`, `sel ∈ 1..5`: the class returned is the **secondary** for `sel ≤ 2` and the **primary** for `sel ≥ 3`, and the cost is `cost[sec]`, `(3·cost[sec]+cost[pri])>>2`, `(cost[sec]+cost[pri])>>1`, `(3·cost[pri]+cost[sec])>>2`, `cost[pri]` for `sel = 1..5`. Both tables are written as immediates by `R0116`, and **`g = 8..11` and `g ≥ 13` are never written** — the water groups are unreachable because of the early-out, but `g ≥ 13` *is* reachable and would read uninitialised heap (the world object is `malloc`'d, not zeroed): 0 of 880 704 shipped cells do it. Pairs: `0 (Grass,Land) · 1 (Cracked,Land) · 2 (Sand,Land) · 3 (Savanna,Land) · 4 (Stones,Land) · 5 (Cracked,Stones) · 6 (Flowers,Savanna) · 7 (Mountain,Stones) · 12 (Road,Land)` — an independent attestation of `TERR-SEM-004`'s labels from the *movement* side. Class census over 880 704 cells: Land 233 611, Grass 104 421, Flowers 10 685, Sand 24 119, Cracked 41 668, Stones 92 880, Savanna 118 009, Mountain 136 622, Water 87 584, Road 31 105, `0xff` **0**

**Confidence.** High (masks, shifts, table displacements and the five arithmetic arms are immediates in one routine with exactly one caller, `R0278` @`L10342`; the corpus never reaches a reject arm or an uninitialised pair) / Medium (the water test's *reading*: `(w & 0x300) == 0x200`, `g ∈ 8..11` and `(w & 0x3ff) ∈ [512,768)` agree on all 880 704 cells and are algebraically the same set for `g < 16`, so only the instruction separates them — and it is the first)

**Original status.** ● active, partially retracted: unconditional water-return prose; other predicates retained

**Evidence.** [EXP-0041](../experiments/EXP-0041-terrain-passability/), [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

**Amended.** The table ledger carried the status "● active, partially retracted: unconditional water-return prose; other predicates retained". `retracted.md` records a correction against this claim.

### TERR-PASS-051

**The rule: a cell blocks a mover iff `block[cell] & mover.mask` over the mover's `n×n` footprint — the block byte is a BITMASK, not the enum `ALM-TERR-016` published.** `R1331` (static plane `world+0x10000`) and `R0144` (dynamic plane `world+0x20000`) are the same routine over different planes: `n = mover->vt[0x1c]()`, `mask = *(byte*)(mover->[0x154] + 5)`, then `for dy,dx < n: if (grid[(cell + dy·0x100 + dx) & 0xffff] & mask) return 0`. The mask is set by **`R0471`** from `mover->vt[0x20]()`: `1 → 0x41` (`L06846`), `2 → 0x44` (`L06847`), `3 → 0x82` (`L06848`); the constructor `R0206` @`L06852` leaves **`0x41`**, so ground is the default. Bit meanings follow from what writes them: **bit 0** = blocks ground (set by all three terrain arms, by a type-3 object as part of `5`, and by the border); **bit 1** = blocks air, set by **nothing but the `0x1f` border**; **bit 2** = static object (`5 = bit0\|bit2`; `R0453` also `OR`s and `AND`s exactly `5`); **bits 3–4** appear only inside `0x1f`, and bit 4 is additionally preserved across a cell-record restore (`L10343`, `L10344`); **bit 5** = the cell has a record in the `world+0x540b8` hash table (100 buckets, `0x34`-byte records — `R0053`, `R0035`, `R1071`, `R1072`, `R0955`, `R0446`, `R0143`, `R1826`, `R1073`, `R0142`, `R1083`, `R1082`, `R1081` — **13 sites testing the byte at `+0x10000` against `0x20` over 13 distinct owners**, `EnumRefs disp:10000` on a repaired function table, 0 of them in orphan or undisassembled code); **bit 6 / bit 7** = a ground / an air occupant stands here, set and cleared only on the **dynamic** plane by `R0058`/`R1332`/`R0257`/`R0046`, again keyed by `vt+0x20` (`<3 → 0x40`, `==3 → 0x80`). Consequences over 880 704 shipped cells: a ground mover is blocked on **430 898** (48.93 %), the `0x44` mover on 229 092, an air mover on **158 976 — exactly the border**. The enum reading is refuted outright and not by counting: `0x41`, `0x44` and `0x82` equal **0** of the derived bytes, so under `==` no mover would ever be blocked anywhere. **External corroboration, recorded at the weight it earns:** of the 8 094 shipped type-6 placements, 7 831 (96.75 %) stand on cells this rule leaves free to a ground mover against 51.07 % of cells free, and all 263 exceptions fall in **7 of 41** placed classes, concentrated in Bee (30.4 %), Dragon (28.1 %) and Ghost (6.7 %) — what a per-mover mask predicts and a universal rule does not. It does **not** pin: the null model "only the border blocks" frees more cells and scores higher **AMENDED by [EXP-0044]**: the two virtuals are per-instance bytes of the **simulation actor** (`vt+0x1c` = `[obj+0x49]`, `vt+0x20` = `[obj+0x4a]`), not per-class values, and the base constructor writes `1` into both — ~~so on a map loaded from a `.alm` **every mover has a 1×1 footprint and the mask `0x41`**, and the `0x44`/`0x82` arms are unreachable~~ **(this consequence RETRACTED by [EXP-0049]: the spawn path streams both bytes from `Data.bin` — Ghost/Bee run mask `0x44` and Bat_Sonic/Dragon mask `0x82`, `TERR-MOVE-057` — which also explains this row's own 263-placement partition, 256 terrain-blocked cells under the ground mask being largely those classes' cells under their real masks)**. `TileSize` is *not* the footprint: it is what **`CUnit`'s** `vt+0x20` returns, a different hierarchy at the same offset (`TERR-MOVE-054`, `TERR-MOVE-055`). The 7-of-41-classes corroboration above therefore corroborates nothing about domains: 45 of the 263 fall on the 2 classes with `Z != 0`, and a `Z`-keyed rule would strand 218, against the null model's 4 **AMENDED by [EXP-0078] on two points.** (i) The three consequence counts above are the **ingest** plane's; after the structure pass (`TERR-STRUCT-068`/`071`) they read **441 540 / 242 179 / 158 976** — domain 3 unchanged, because no structure arm touches bit 1 — and split terrain / non-terrain they are 199 361 + 242 179 / **0** + 242 179 / **0** + 158 976, which is the form a consumer needs (`MOVE-DOM-026`). (ii) The **naming** this row withdrew is restored and now rests on three witnesses that do not depend on the mask: the `Data.bin` columns that produce the codes (`TERR-MOVE-057`), the `Z`/`CAirUnit` split falling on exactly the two code-3 classes (`REG-UNITS-061`), and the AI's own flier distance penalty at `L00225` (`MOVE-DOM-024`(e)). What must **not** be carried over with the name is “two grades of flight”: mask `0x44` carries the object bit and not the terrain bit, `0x82` neither, and the two disagree on 83 203 shipped cells

**Confidence.** High (the predicate, the mask setter, its default and the occupancy writers are each named instructions; the bitmask-vs-enum rival is excluded by the `TEST`/`JNZ` form together with the 0-cell equality count) / **Medium, was Unknown, was Medium** (calling `vt+0x20`'s codes *movement domains* — withdrawn by [EXP-0044], which found no producer of any value but 1 and no registry path to the field (`claims/retracted.md`); the producer was found by [EXP-0049] and the *name* attested twice more by [EXP-0078])

**Original status.** ● active, as amended

**Evidence.** [EXP-0041](../experiments/EXP-0041-terrain-passability/), **[EXP-0044](../experiments/EXP-0044-mover-domain/)**, **[EXP-0078](../experiments/EXP-0078-movement-domains/)**

**Amended.** The table ledger carried the status "● active, as amended". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-COST-052

**What the restored `Cost` scalars are consumed by — and it is never the block decision.** The blended byte `R0469` writes lands at `world + (row<<8) + col` (displacement 0, hence the Unknown clause below) and is read at two places. (a) **Per pathfinder step**, `R1326`/`R1327`/`R1328`: after the same `& mask` footprint test, an **orthogonal** step adds `cost[cell]` and a **diagonal** step adds `cost + (cost>>1)` — but only when `mover->vt[0x20]() == 1`; every other mover pays a flat **2** and **3**. (b) **Per move duration**, `R1088`: `speed = SpeedMultiplier (world+0x58db4) × unitSpeed`, tilted by the two cells' height difference from `world+0x9451c` **clamped to ±32** (`speed ± (\|Δh\|·speed)>>6`), then divided by the **mean of the two cells' cost bytes** (`(c₀+c₁)>>1`, replaced by **8** if that mean is 0) and clamped to `[1, 63]` before being stored at `mover->[0x154] + 0xa8`. ~~Both cost reads go through `R1087`~~ — **AMENDED by [EXP-0054]**: only the **duration** read (b) does. `R1087` returns `costPlane[cell]` and, if the cell carries a hash record whose byte 2 is set, **quarters it in place** (a shift right by 2, written back), and it has **exactly two call sites, both inside `R1088`** (`EnumRefs callto:R1087`, 2 hits / 1 owner / 0 orphan) — so the bit-5 quartering never reaches a **path** cost. The search side reads the plane directly (`L06760`), in **nine** places, not the three this row names: `R1326`, `R1327`, `R1328`, `R1329`, `R1330`, two copies inlined in `R0053` and two in the extractors `R0437`/`R0438` — exactly the widening this row's own Unknown clause said `disp:` could not see (`claims/retracted.md`; `MOVE-COST-002`). Corpus: the cost byte takes only `6..16` over 880 704 cells (58.9 % are exactly 8), and swapping the file's `Cost` values for the code's hardcoded defaults changes **128 694** cells (14.6 %) — so which of the two is in play is not academic. The file's win: `R0505` compares 15 characters and every `Cost*` key is ≤ 12

**Confidence.** High (the two consumers, the orthogonal/diagonal split, the `vt+0x20 == 1` gate, the ±32 clamp, the `[1,63]` clamp and the `>>2` fixup are named instructions) / **Unknown** (whether other readers of the cost plane exist. Its base displacement is **0**, so `EnumRefs disp:` — the instrument used for every other reader list here — structurally cannot see them; the readers above were reached through the routines that also touch `+0x10000`/`+0x20000` and through the two address-of-cell getters `R1349` and `R1351`, whose callers were enumerated, 1 and 2, both attributed. That is a reachability argument, not an enumeration — **and also [EXP-0054] found six more, exactly here**)

**Original status.** ● active (amended)

**Evidence.** [EXP-0041](../experiments/EXP-0041-terrain-passability/), **[EXP-0054](../experiments/EXP-0054-unit-movement/)**

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-PASS-053

**The block planes reach persistent state; the cost plane does not.** `R1360` is an MFC `Serialize` (it tests `CArchive+0x14 & 1` = `IsStoring` and calls the CRT archive helpers `R0272`/`R0271`). Storing: it sweeps a **flat run** of the dynamic plane — `world + 0x20807` for `0xe5e7` bytes, i.e. plane index `0x0807 … 0xEDED` — and for every cell whose **dynamic** byte is `> 0x0f` emits one `u32` `(cellIndex << 16) \| (dyn << 8) \| static`, reading the static byte at `[EBX + 0xffff0000]` = the same index in `world+0x10000`. Loading: `*(char*)(world + (v>>16) + 0x20000) = (char)(v>>8)` and `*(char*)(world + (v>>16) + 0x10000) = (char)v` (`L10345`/`L10346`). So **both** block planes are save state, byte for byte, for every cell carrying a bit above bit 3 — which on a freshly ingested map is exactly the `0x1f` border, and at runtime adds the bit-4 flag, the bit-5 record flag and bit-6/7 occupancy. The **cost** plane and the **height** plane are not touched by this routine. A consumer that hashes simulation state must therefore include the two block planes and must reproduce the runtime bits, not only the map-derived ones

**Confidence.** High (the archive-mode test, the sweep bounds, the pack/unpack arithmetic and the two store addresses are named instructions) / Medium (calling `[0x807, 0x807+0xe5e7)` "the interior": it is the literal range in the code, not one derived from `W`,`H`, and it is neither a rectangle nor aligned to the 8-cell border) / ~~**Unknown** (whether loading a save re-runs `R0278` before applying these bytes — the ingest's three callers were located but the load order was not traced)~~ — *answered by [EXP-0148]: it does. `R0414`'s load arm constructs the terrain from the loaded `.alm` at `L07067`, whose own tail is the call to `R0278` at `L08344`, and calls `R1360` only at `L07068` (`SAV-LOAD-057`)*

**Original status.** ● active

**Evidence.** [EXP-0041](../experiments/EXP-0041-terrain-passability/)

### TERR-MOVE-054

**Who the mover is: not `CUnit`.** The object whose `vt+0x1c`/`vt+0x20` `R1331`, `R0144`, `R1326`, `R1088` and `R0471` dispatch on is the **simulation actor**, a C++ family whose base vtable is `L00001` (derived: `L00002`, `L00003`) — the only class that allocates the `0xb4`-byte mover and stores it at `+0x154` (`R0185` @`L00963` (pushing `0xb4`) → `R0206` → `L06970` (the store into `+0x154`)). `CUnit` is **not** it: its own constructor writes *bytes* into `+0x154..+0x15b` (`L10347`, `L10348`, `L10349`, `L10350`, `L10351`), so `[CUnit+0x154]` is not a pointer, and `R0471`'s only two call sites (`L06850`, `L06851`) pass `(this->[0x154], this)` from methods of the `L00001` family. The nine drawable classes are recoverable **by name** from the MFC `CRuntimeClass` records at `L10352..L10353` — `CGameObject` `0x138`, `CStructure` `0x138`, `CBridge` `0x140`, `CVerticalWoodenBridge`/`CHorisontalWoodenBridge` `0x140`, `CBackPack` `0x138`, `CUnit` `0x1b0`, `CAirUnit` `0x1b0`, `CProjectile` `0x14c` — each vtable identified by the a stub that returns the `<CRuntimeClass>` address stub in its slot 0. **Every vtable offset in this row and the next is anchored to the vptr a constructor installs**, not to the start of a run of `.rdata` code pointers: `imm:L02468` has 3 hits and `imm:L13167` has 0, and taking the run boundary instead shifts every slot by one and makes `CUnit`'s `vt+0x1c` read as `TileSize` — a self-consistent wrong answer

**Confidence.** High (the class↔vtable pairing is a name string, a size and an installed immediate for each of nine classes; "`CUnit` is not the mover's owner" is two independent instruction facts — the byte writes at `+0x154` and the call sites' `this`)

**Original status.** ● active

**Evidence.** [EXP-0044](../experiments/EXP-0044-mover-domain/)

### TERR-MOVE-055

**What the two virtuals return, and that nothing in the image ever varies them.** For the simulation-actor family, `vt+0x1c` is `R0256` = a getter returning the byte at `+0x49` (the `n×n` footprint side) and `vt+0x20` is `R1344` = a getter returning the byte at `+0x4a` (the block-mask selector); the 17 sibling classes of the same family return the constants `1` and `0` (`R1827`, `R1828`). Both fields are written **once**, by the base constructor `R0185`: `L00821` (a store of the byte `0x1` at `+0x49`) and `L00820` (a store of the byte `0x1` at `+0x4a`). No other instruction in `rom.exe` stores into either field of this class — the two remaining immediate stores at those displacements (`L05637`, `L05428`) belong to the class with vptr `L05352`, whose own `vt+0x1c`/`vt+0x20` are the constant stubs. What *can* change them is four routines that take the fields' **addresses** and hand them to a field-by-field (de)serializer — `R0180` (`L04084`/`L10354`), `R0657` (`L04085`/`L10355`), `R1601` (`L10356`/`L10357`) and `R0210` (`L10358`/`L10359`, the loading branch of an MFC `Serialize` whose storing branch reads the same two bytes at `L08383`/`L10360`). ~~**Consequence for a session started from a `.alm`: every mover has footprint 1 and mask `0x41`**, so `R0471`'s `0x44` and `0x82` arms and the `vt+0x20 != 1` arms of `R1326`/`R1088` are unreachable~~ — **RETRACTED by [EXP-0049]** (`claims/retracted.md`): two of the four address-takers this row itself lists are the **Data.bin param streamers, run at spawn from inside the creation path** — `R0657` from `R0656` (humans), `R0180` from the base ctor's own continuation `R0185 → R0184` (units) — so on a `.alm` load `+0x49`/`+0x4a` are overwritten from the class's `tokenSize`/`movementType` parameters whenever those are not −1 (`TERR-MOVE-057`). **And no `units.reg` key can reach either field**: the class array `L02113` has 99 references over **20 distinct owners, every one inside `R0379…R0389`** — the drawable and loader modules — and none in the simulation module code-region where the mover, the predicate and the planes live

**Confidence.** High (each of the two getters, the two constructor stores and the four address-takers is a named instruction; the "only value is 1" claim is an immediate-store enumeration **plus** an independent byte-level scan of `.text` for every ModRM memory-write encoding covering those displacements, which does not use the disassembler at all — 3 and 2 byte-sized candidates respectively, all already attributed. **These clauses stand as amended: the streamer writes go through the pushed addresses, which is exactly the door the row's Medium clause left open**) / **~~Medium~~ (that a running game never holds another value) — overturned by [EXP-0049]**: the streams' source is the definition entry's param array, and 16 shipped unit entries carry codes 2/3 / ~~Unknown~~ (what codes 2 and 3 were meant to mean) → **answered**: `movementType` 2 = Ghost/Bee (mask `0x44`), 3 = Bat_Sonic/Dragon (mask `0x82`, air) — `TERR-MOVE-057`

**Original status.** ● active (amended)

**Evidence.** [EXP-0044](../experiments/EXP-0044-mover-domain/), **[EXP-0049](../experiments/EXP-0049-placeable-db/)**

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-MOVE-056

**The per-step movement speed, `R1088`, end to end.** Signature `__thiscall(world, actor, u16 srcCell, u8 facing)`. `dir = ((facing + 0x10) >> 5) & 0xff` (0..7, `L06859`); `dstCell = srcCell + (i16)world[0x58ec0 + 4*dir]` (`L06861`); the actor's raw speed is `actor->[0x70]->[0x3c]->[0x44]` when that byte is nonzero, else `(i16)actor->[0x8c]` (`L06933`/`L10361`; `+0x8c` is `16` at construction, `L10362`). If `vt+0x20() != 1` the raw speed is used as is. Otherwise: `d = clamp((i8)(height[src] − height[dst]), −32, +32)` over the height plane `world+0x9451c` (`L06866..L10363`); `v = SpeedMultiplier * speed` with `SpeedMultiplier` = `world+0x58db4` = `data/map.reg [Path Finding]`, shipped value 8 (`L06868`); `v += (v*d) >> 6` **arithmetic** shift, so uphill (`d < 0`) reduces and downhill increases (`L10364..L06871`); `c = (u8)(cost[src] + cost[dst]) >> 1`, a **byte** add, and `if c == 0 then c = 8` (`L06872..L05504`); `v /= c` signed (`L06873`). Then `v = clamp(v, 1, 63)` (`L06874`), `mover[0xa8] = v`, `mover[0xae] = dir`, and the per-axis steps `mover[0xb0]/[0xb1] = v*dx`, `v*dy` from the direction tables `world+0x58eb0`/`+0x58eb8`, multiplied by the **double `0.707` at `L06908`** and rounded through the CRT `ftol` when the step is diagonal (`dx*dy != 0`); finally `mover[0xaa] = ceil(256 / max(\|b0\|,\|b1\| chosen as b0 if nonzero, 1 if both zero))`. `cost(cell)` is `R1087` and is **not a pure read**: when `block[cell] & 0x20` and the cell's record in the `world+0x540b8` table has a nonzero byte at `record+0xe`, it shifts the stored cost byte right by 2 and **writes it back** (`L05498`/`L05499`). Corpus: over every orthogonally adjacent cell pair of the 38 shipped maps at speed 16 and `SpeedMultiplier` 8, `v ∈ [4..32]` with mode 16 (22.2 % of 1 740 396 pairs) — **neither clamp bound is ever reached**, so `[1,63]` is a guard, not a shaping term

**Confidence.** High (every term, its address and its arithmetic form — `SAR` vs `SHR`, byte-width add, signed `idiv` — is transcribed from the listing, which is quoted in full in the experiment's evidence; the corpus range is a re-implementation check, not the source) / Medium (the range figures: one parameter set, `v0 = 16`, on shipped terrain only)

**Original status.** ● active

**Evidence.** [EXP-0044](../experiments/EXP-0044-mover-domain/)

### TERR-MOVE-057

**What fills the simulation actor's `+0x49`/`+0x4a` (EXP-0044's open item): the placeable-definition database, at spawn — and block-mask codes 2 and 3 ARE produced.** Two of `TERR-MOVE-055`'s four address-takers are **Data.bin param streamers** run inside the creation path itself: humans `R0656` → `L10365` (an add of `0x8`) and `L03160` (the call to `R0657`), whose ordered reads end `slot 0x15 TokenSize → +0x49` (`L04085`), `slot 0x16 MovementType → +0x4a` (`L10355`); units base ctor `R0185` → `L00493 R0184` → `L10366 R0180(entry+8)`, slots `0x1f`/`0x20` → the same bytes (`L04084`/`L10354` — the very addresses EXP-0044 listed). Each read helper **skips the store when the param is −1** (an empty CSV cell), so the ctor's `1` survives exactly where the table is silent — all 210 parameterised humans. The shipped Units table stores `movementType` **2 on Ghost/Bee** (→ mask `0x44`) and **3 on Bat_Sonic/Dragon** (→ `0x82`, air), 4 entries each incl. difficulty variants, and `tokenSize` 2 on 11 / 3 on 4 entries — so `R0471`'s three arms and the `n×n` footprint walk are all live on shipped data, and EXP-0044's stranded-placement concentration in Bee/Dragon/Ghost is explained by their real masks. **Read this row's own gloss carefully — [EXP-0078] found it is where a consumer goes wrong.** “2 = Ghost/Bee (`0x44`), 3 = Bat_Sonic/Dragon (`0x82`, air)” reads as two grades of one thing and is not: terrain and objects are **different block bits**, `0x44` carries the object bit and not the terrain bit while `0x82` carries neither, so **domain 2 crosses water and stops at a tree and only domain 3 passes both** — they disagree on 83 203 of 880 704 shipped cells (`MOVE-DOM-026`). Also measured there: all 210 Humans rows carry −1, so every human is a ground mover, and no row in either table carries a value outside `{−1,1,2,3}` (`MOVE-DOM-028`)

**Confidence.** High (the streamer calls, their `entry+8` source, the ordered slot reads and the −1 gate are named instructions — `evidence/rom-placedb-excerpt.md` §6–7 — and the values are the shipped file's own bytes, `databin-units.csv`)

**Original status.** ● active (amended)

**Evidence.** [EXP-0049](../experiments/EXP-0049-placeable-db/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-MOVE-058

**The speed override `actor+0x70 → +0x3c → +0x44` (`TERR-MOVE-056`) is NOT the definition database; the DB supplies the fallback.** The fallback `(i16)actor+0x8c` is the class's `speed` parameter (humans slot 6 / units slot 8, streamed at spawn — construction default 16). The chain: `actor+0x70` = the **owner object** the actor is attached to at spawn (`R0152`: `actor+0x70 = owner`; called from the type-6 spawner `L10367`/`L10368`, the groups reader `L10369`, the order dispatch); the owner's ctor `R0149` allocates a **0x50-byte block** (`R0150`) into `owner+0x3c`; the override is that block's byte `+0x44`, **runtime state**: zeroed at the end of every pass of the per-owner unit loop (`R0200` @`L06891`, together with `+0x20`), stored from message/order bytes at `R0108`/`R0157` and as the constant 2 at `L10370` (`R0887`). No Data.bin value reaches it

**Confidence.** High (the fallback's source — the streamers write `+0x8c` from the `speed` slot — and every hop of the chain is a named instruction; writer list instrument: `EnumRefs` with a regex for an 8-bit store to `[E..+0x44]` whole image, 28 hits, 0 orphan, module-filtered by the `[+0x3c]` base) / ~~**Unknown** (the byte's *meaning* — which order/state sets it and what the values encode; only its writers are located)~~ — **CLOSED by [EXP-0093]**: `actor+0x70` is the AI **group** (`AI-GROUP-009`), `owner+0x3c` its `0x50`-byte AI record, and `+0x44` is the **minimum `Speed` over the group's members**, written by the two group-order setters beside `grpAI+0x20 = 4` and cleared with it (`MOVE-GROUP-030`). The two message/order writers this row names are those setters; the constant 2 at `L10370` belongs to a different object and is **not** this byte

**Original status.** ● active (amended)

**Evidence.** [EXP-0049](../experiments/EXP-0049-placeable-db/)

**Amended.** The table ledger carried the status "● active (amended)".

### TERR-LIGHT-059

**A sprite IS lit — and the sprite blit family splits in two, decided by which pointer its pixel loop reads.** The `.256` class dispatches on vtable **`L04339`** (stored by its own payload constructor `R0753` and by `R1829`; `EnumRefs imm:L04339` — 2 hits, 0 in orphan code — so the merged `FindVtables` run that puts `R1830` at `+0x3c` of the adjacent `L10048` names a **different class**). Four blit entry points. **`vt+0x14` `R1791` and `vt+0x34` `R1547` are LIT**: six args `(dstX, dstY, frame, level, shadeObj, mirror)`, `L10371` (a shift left by `0x9`), `L10372` (a load of the shade object's `+0x8`) and `L10373` (their sum) = `row = shadeObj[+8] + (level << 9)`, handed to a loop that reads the **source** palette index (`L10374` (a load of the source byte)) and looks it up (`L10099` (a load of the 16-bit table entry at `2 × index`)) — the same `LUT[(level<<9) + srcIndex*2]` shape as the terrain blitters (`TERR-LIGHT-011`), at **per-sprite** granularity: one row for every pixel of the blit. **`vt+0x1c` `R0795` (5 args) and `vt+0x3c` `R1831` (6 args) are NOT lit**: they take no shading object and no `<<9`; their eight pixel loops (`L03664/e240/eef0/f140` and `R1832/e750/f380/f690` — the four-way `mirror × [L03346]` split) **advance the source without ever reading it** and instead read the **destination** (`L10375` (a load of the 16-bit destination word)), a shift right by 3, index a *second* table `[L03341]` at row `arg × [L03340]`, and write it back: destination recolouring under the sprite's silhouette, i.e. the shadow/shroud pass. Which pass uses which: the object cell issues `vt+0x3c` twice (shadow, `TERR-SPR-043`) then `vt+0x14` + `vt+0x34` twice (body, `L10303`/`L10304`); a unit issues `vt+0x2c` = `R0553` for the shadow and `vt+0x28` = `R0552` for the body (`TERR-SPR-065`). The `vt+0x34` pair additionally reads the destination and `[L07813]` — the `spritesb` overlay blend, whose consuming instructions this locates and does not read

**Confidence.** High (the source-vs-destination read is one named instruction in each of the twelve loops, and the vtable is pinned by its constructor's immediate with an image-wide writer enumeration)

**Original status.** ● active

**Evidence.** [EXP-0053](../experiments/EXP-0053-sprite-lighting/)

### TERR-LIGHT-060

**Mode 2 — the exact per-entry transform every sprite table is built with.** `R1107` dispatches six modes through the jump table at `L05822` (raw bytes at file offset `0x27884`: `L10376 L10377 L10378 L10379 L10380 L05821` = modes 0…5), and **mode 2 is the arm at `L10378`**. Per level row and palette entry, per channel, in integers: `out = clamp(((pal_chan + skyTint_chan) × m × 2) / nLevels, 0, 255)` — `L10381` reads the RGBQUAD channel, `L10382` adds the tint, `L10383` multiplies by the row counter, `L10384` shifts left by `0x1`, `L10385` sign-extends and divides (signed, truncating toward zero, the divisor being `this[+4] = nLevels`), then the same two 16-bit clamps mode 3 uses, packed through the same DirectDraw bits/shifts (`TERR-LIGHT-019`). `m` is the row counter and counts **down** from `nLevels` to 1 while the destination cursor advances (the decrement at `L10386`), so **row `L` carries gain `2(nLevels − L)/nLevels`**: at `nLevels = 0x10`, row 0 → ×2.0, row 8 → ×1.0 (confirming `TERR-LIGHT-022`'s "neutral row is `nLevels/2`"), row 15 → ×0.125. The other four arms, for completeness: mode 0 = one row of `0x0000` then 255 × `0xffff`; mode 1 = one tint-only row with **no** level multiply; mode 4 reads `[L03346]` (unread here); mode 5 = a greyscale ramp `((B+G+R) × m × 2)/(3 × nLevels)`, which is what a unit takes on the `L10387` override

**Confidence.** High (every operation named to an address in an arm read end to end, and the mode→arm assignment is the dispatch table's own bytes)

**Original status.** ● active

**Evidence.** [EXP-0053](../experiments/EXP-0053-sprite-lighting/)

### TERR-LIGHT-061

**What feeds a sprite's level: `CMapView+0xb0`, a per-frame byte grid whose base value is a single global.** The blit's `level` argument is `grid[(row+3)*(visCols+6) + (col+3)]` — read at `L10388`/`L10389` on the object path (masked `0xff` at `L10390`, re-read for the `vt+0x34` pass at `L10391`) and at `L10392`/`L10393` on the unit path. The grid is `malloc((visCols+6)*(visRows+10))` at `L10394`/`L10395` in `R1280`, the **only** writer of the pointer (`EnumRefs disp:b0`: 131 hits, 78 owners, 1 in orphan code and that one not a `CMapView` displacement). Its **only** content writer is `R0608`, in three stages. **(1)** `memset(+0xb0, [L05650] >> 2, size)` — the whole visible window, from the sun model's **ambient** byte, `0x0e >> 2 = 3` in the daytime band (`L05653`, `L10396`, `L10397`). **(2)** per entry of the lit-object list at `this+0x9f0` (count `+0x9f8`, data `+0x9f4` — `TERR-EDGE-025`'s list): flag `0x1000` → byte `0` (`L05656`), flag `0x20000` → byte `0xc` (`L05657`), flag `0x8` → a `rand`-driven flicker between `0` and `0xc` seeded by `[this+0xa70]/2 + col*row`, plus a radius-1 stamp into `+0xa8` via `R1095` (`L10398…L10399`). **(3)** a sweep over every cell that overwrites it with `min( (Σ clamp0(a8corner − 0x20)) >> 4 , ambient>>2 )`, skipping cells whose four `+0xa8` corners are all `0xff` (a comparison of the sum with `0x37c` = 4 × `0xdf`, `L10400`) and substituting `ambient` for any single `0xdf` corner. `+0xa8` is a **light-source stamp** grid of `(visCols+7)×(visRows+11)` bytes, `0xff` = "no stamp", written only by `R1095(col,row,radius,value)` over the 41×41 distance table at `L10401`; the hardware-renderer path `R0541` names it by its fallback (`L10402` (a comparison with `0xff`) and `L10403` (a load of `[landscape+0x18]`) — the per-vertex terrain brightness). Consequences: the level is **uniform over the map** except at lit-object cells; lower level = brighter, so stage 3 can only brighten and the sole darkening mechanism is a `0x20000` object stamping `0xc`; and **objects and units read the same grid**, so no shipped frame can light one differently from the other. **Refutes** the reading that the sprite level is the shroud level (that is the silhouette family's argument, `[L04368]`/`[L04335]` into `[L03341]`) and the reading that it is sampled from the terrain brightness grid

**Confidence.** High (the read index, the allocation size, the memset and all three stages are named instructions, and the pointer's writer set is an image-wide enumeration whose instrument and blind spot are stated) / **Medium** (that the byte reaches the *unit* blit on every path: the per-cell dispatch `L02996` pushes it, but `R0552` also reads that argument slot as a frame index at `L10321`, forces it to `0` for the unit at `CMapView+0x98c` at `L10404`, and the second `vt+0x28` dispatch `L02994` pushes `0`)

**Original status.** ● active

**Evidence.** [EXP-0053](../experiments/EXP-0053-sprite-lighting/)

**Amended.** The sole stamp-writer and lit-object-only population clauses are partially retracted: MAGIC-269 and MAGIC-270 establish direct CProjectile path stores. The same-grid sentence does not guarantee equal lighting across drawable passes; the existing Medium exceptions stand. MAGIC-273 gives the named consumers and option boundary. The allocation, initialization and four-corner arithmetic stand. See claims/retracted.md.

### TERR-LIGHT-062

**The sprite ladder and the terrain ladder are the SAME ladder, and in the shipped daytime band the two are half a step apart.** `2(16 − S)/16 = (96 − (4S + 32))/32` identically, so mode 2's row `S` has exactly mode 3's row `L = 4S + 32` gain; built from the *same* palette the two rows are **bit-identical — 0 of 256 entries differ** at every one of `L = 32, 36, 40 … 64`. So a sprite level is a 4:1 quantisation of a terrain level with the domain shifted by 32, and the engine *can* give a sprite the gain of the ground it stands on. It very nearly does: flat daytime ground is `L = 46` → **×1.5625** (`TERR-LIGHT-013`'s `L = (R>>1) + ambient + 32 = 62`, `TERR-LIGHT-020`), a sprite is `S = 14>>2 = 3` → **×1.6250**, and the *exact* ambient quarter `14/4 = 3.5` would give `2(16−3.5)/16 = 1.5625`, the flat-ground gain to the last bit — the integer `SAR 2` at `L10396` rounds the sprite one half-step brighter. Not an arithmetic coincidence: both indices descend from the **same byte** the global at `L05650`, whose only two readers image-wide are `R0608` (sprites, `L05653`) and `R0468` (terrain, `L10405`), against seven writers all inside the sun model `R1813` (`EnumRefs refto:L05650`, 9 hits, 0 in orphan code). **Consumer consequence, and the reason this row exists:** a renderer that lights terrain and leaves sprites at their palette values does not produce a subtle mismatch, it produces a **38 % luminance deficit on the sprite alone** — measured on the review render, the shipped swordsman frame means 104.9 luma at the engine's row 3 and 65.2 with no table, against 106.1–118.3 for the terrain in the same crop — and mode-2 row `nLevels/2` is pixel-identical to no table at all

**Confidence.** High (the identity is arithmetic and the bit-equality a measurement over the shipped palette; the the global at `L05650` reader set is an image-wide enumeration) / **Medium** (that a live session runs at this ambient and θ — static observation only, the ceiling `TERR-LIGHT-029` carries)

**Original status.** ● active

**Evidence.** [EXP-0053](../experiments/EXP-0053-sprite-lighting/)

### TERR-LIGHT-063

**The terrain mapping's OUTPUT range, and it saturates to white on shipped data.** Per channel, mode 3: `out = clamp(((pal_chan + skyTint_chan) × (96 − L)) / 32, 0, 255)`, the `/32` truncating toward zero (a sign-extension bias masked with `0x1f`, added, then a shift right by 5 at `L10406…L10407`), then `R>>3 << 11 \| G>>2 << 5 \| B>>3` on the shipped RGB565 surface — re-derived here from `L10379` and agreeing with `TERR-LIGHT-019` instruction for instruction. The level census is re-run on the corrected base with the exact `TERR-LIGHT-028` span: **`[30..70]` over 38 maps / 859 768 interior vertices at θ = 0.78539815**, reproducing `TERR-LIGHT-029` exactly. So the brightest gain a shipped map can produce is `(96−30)/32 =` **×2.0625** and the darkest `(96−70)/32 = ×0.8125`. Against the one palette all 53 `terrain.3d` tiles share, pixel-weighted over 651 266 tile pixels: at **level 30** — 89 of 256 entries have a clipped channel, **33 pack to exactly `0xffff`**, and **16.99 % of pixels clip / 5.09 % go pure white**; at the flat daytime level 46 — 47 entries clip, 11 pure white, 10.59 % / 0.10 % of pixels, mean luma 119.49; at level 64 (×1.0) — 8 entries clip, 1 pure white, 0.30 % / 0.00 %, mean luma 79.48; at level 70 — nothing clips, max packed `0xce79`. **So the mapping does not stop short of white: bright shipped terrain reaching white is what the engine does.** Two refinements to `TERR-LIGHT-020`, neither a contradiction: its "0 % of pixels clipping at level 64" is true under the all-three-channels reading and 0.30 % under the any-channel one, and its 10.6 % / 119.2 at level 46 reproduce as 10.59 % / 119.49

**Confidence.** High (the transform is instruction-level in a function read end to end; the census is arithmetic over a measured palette and a re-measured level range, with θ, framing base and window quoted beside the figure)

**Original status.** ● active

**Evidence.** [EXP-0053](../experiments/EXP-0053-sprite-lighting/)

### TERR-LIGHT-064

**Which drawable gets which table, and the palettes are disjoint.** Terrain: one table for every tile, `R0919(0x60, 3, 1)` on the first non-null `tiles[]` slot, `[L10222] = obj + 0x14` (`L10408…L10409`; `TERR-LIGHT-018`/`022`, re-derived). Each **object** sheet: its **own** `(0x10, 2, 1)` table from its **own** BMP palette, built by the sprite loader right after construction (`L10410…L10411`), and the blit is handed `spriteObj + 0x14` (`L10412`). A **unit**: the shading object is selected on the class field `class+0x98` — `== 0` → one of **16 shared** tables `[L05816][sprite[+8]]`, built by the relight driver as `(0x10, 2, 1)` over a 16 × `0x400` palette blob (`L10413…L10414`, `L10238 … L10415`); `== 1` → the class's own `class+0x9c` from palette `class+0xac`; `> 1` → **a per-TIER table** `class+0x9c + (face−1)*4` from `class+0xac + (face−1)*4` (`L07205`, `L10416`, `L10417`) — **amended by EXP-0089**: that subscript is `unit+0x24 − 1`, and `unit+0x24` is the actor's `face`, not its owner (`PAL-FACE-005`); the *owner*-subscripted arm is the `== 0` one, so this row had the two arms swapped (retracted.md) — with two overrides, `[L05816]+0x40` and a per-draw `(0x10, mode 5, 0)` greyscale from `class+0xac` (`L10418`, `L10387`). The 16 KB blob is **not static data**: the VA `L10238` maps past the image's `0x1e2000` bytes, and `R0395` fills it with one `R0533(L10238, 0x4000)` read at `units.reg` load. Corpus: the 53 `terrain.3d` BMPs carry **1** distinct palette; the **1373** palette-bearing `.256` sheets carry **1225** distinct palettes and **0 of 1373** is the terrain palette (closest: max per-channel delta **201**). So terrain and sprites share the *arithmetic* (`TERR-LIGHT-062`) and nothing else — a renderer must key each sprite's ramp off that sprite's own palette

**Confidence.** High (each `(nLevels, mode, useTint)` triple is a push sequence at a named address; the palette census is a measurement with no model to be wrong about; the blob's non-staticness is a section-geometry fact) / High (**`class+0x98` is the LENGTH of the `class+0x9c` array** — the relight driver iterates `class+0x9c + i*4` for `i < class+0x98` to free them, `L10419…L10420`, which is exactly the domain the body's selector splits on — so `0` = owns none/use the 16 shared, `1` = one own, `n > 1` = one per **tier** subscripted `face−1`) / **Unknown, three clauses of which EXP-0089 CLOSES**: `class+0x98` is the `units.reg` key `Palette` (`PAL-KEY-002`); the 34-class split is 18 (`Palette` 0, every human/hero) / 3 (`Palette` 1) / 13 (`Palette` 4, every tiered monster); and the 16 KB node is `graphics\units\humans\human.pal` (`PAL-OWN-007`), whose 16 sub-palettes differ from each other on exactly one fixed 55-index band. **EXP-0182 CLOSES what the two mode-5 overrides are for**: both are the `stone_curse` arm of the unit draw. `L05815` gates them on `R0615(this, 0x30)` and they replace the shade table the blit takes, `[[L05816] + 0x40]` when the class's `Palette` is 0 and a per-draw greyscale ramp otherwise, so a stone-cursed unit is drawn grey (`MAGIC-STONEDRAW-084`). What stays Unknown is `[L03341]`, which one of them reaches and which is another lane's surface

**Original status.** ● active (amended)

**Evidence.** [EXP-0053](../experiments/EXP-0053-sprite-lighting/), [EXP-0089](../experiments/EXP-0089-tier-hue/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-SPR-065

**The unit BODY pass is `vt+0x28`, not `vt+0x2c` — `vt+0x2c` is the unit's shadow (amends `TERR-SPR-048`).** `R0379` runs the drawable grid `CMapView+0x90` twice, both times at `idx = (row+3)*(visCols+6) + (col+3)` under the same `drawable+0x78 == 0` guard, and the two passes differ in exactly one argument: `L02993` (a call through the vtable slot at `+0x2c`) is handed the `CMapView+0xc0` four-corner altitude (`TERR-SPR-039`) and `L02996` (a call through the vtable slot at `+0x28`) is handed the `CMapView+0xb0` light level (`L10392…L10421`). CUnit `vt+0x28` = **`R0552`** (`L02468+0x28`; `CAirUnit`'s `L02585+0x28` is the same routine — `EnumRefs callto:R0552` returns exactly those two `.rdata` slots and nothing else, 0 in orphan code), and it is the routine that issues the lit blits: `L07247` (`vt+0x34`) and `L07248` (`vt+0x14`) over `[L07265][class.File]`'s `+0x04`/`+0x08` sprites, plus `L05824`/`L05826` and `L07246`/`L10088` for the `unit+0x194` overlay sheets. `R0553`, which `TERR-SPR-048` published as "the unit sprite pass", contains **no lit blit at all**: its 17 indirect calls are only `+0x20`, `+0x24`, `+0x1c`, `+0x3c` — the width/height getters and the silhouette family — so what that routine draws is the unit's **shadow**, sheared by the sun scalar it computes at `L10338…L07952`. Everything `TERR-SPR-047`/`048` established about it stands (the nine-state frame switch, the anchor, the `Files` indirection, the mirror flag): it is the shadow's frame, destination and mirror, and the shadow must track the animation exactly as the body does. `R0552`'s own three arguments sit at frame `+0x58/+0x5c/+0x60`, the first two are never referenced, and the third is the level. **Population amendment:** these late +0x90 caller sites are the CAirUnit selector3 route, not universal CUnit registration; ordinary and alternate CUnit share the drawing targets but use other phases (ANIM-CATEGORY-084, ANIM-AIRPASS-086).

**Confidence.** High (both dispatches with their differing argument, the `callto` enumeration of the body routine's entry points, the six lit blit sites and the exhaustive indirect-call list of the shadow routine are all named instructions or image-wide enumerations with their instrument stated)

**Original status.** ✔ promoted (partially retracted)

**Evidence.** [EXP-0053](../experiments/EXP-0053-sprite-lighting/); [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/)

**Amended.** The table ledger carried the status "✔ promoted (partially retracted)". `retracted.md` records a correction against this claim.

### TERR-SPR-066

**The silhouette blit's fifth argument is the SUN SHEAR, which reconciles the contradiction `TERR-SPR-048` left open (amends `TERR-SPR-038`).** `TERR-SPR-038` wrote the `vt+0x3c` push list as `0, level, [L04368], frame, dstY, dstX` and `TERR-SPR-048` recorded that the unit path passes "the sun scalar" in the same slot, calling the two irreconcilable. They are the same thing and the word `level` was wrong: on the object path that slot is the frame local at `0xfffffdf8`, written **once**, at `L07860`, from the calls to `R1820` (`L10422`) and `R0812` (`L03742`), a multiplication by `[L10423]` (`L10424`) and a call to `R0279` (`L10425`) — an `__ftol` of a sun-angle expression, exactly the unit path's `L10338…L07952`. The blitter consumes it as a **16.16 per-row X slope**: `L10426` (a load of the argument at frame offset `+0x20`), `L07857` (a multiplication), `L07858` (a mask with `0xffff0000`) and `L07859` (a shift right by `0xf`), accumulated per row at `L10427` — the shadow's lean. The *fourth* argument is the shroud level (`[L04368]`, or `[L04335]` on the unit path's alternate arm), and it indexes the destination-recolour table `[L03341]` at stride `[L03340]`. The unit's `vt+0x1c` arm carries no per-row slope and instead offsets `dstX` by `shear/2000` (`L10339` loads the constant `0x10624dd3`; `L10340` shifts right by `0x7`), so the two arms are a skewed and a translated shadow. **A consumer implementing `vt+0x3c` from `TERR-SPR-038` alone would feed a brightness into a geometry parameter**

**Confidence.** High (the slot's single writer, the FPU call chain and the blitter's accumulator are named instructions)

**Original status.** ● active (amends `TERR-SPR-038`, `TERR-SPR-048`)

**Evidence.** [EXP-0053](../experiments/EXP-0053-sprite-lighting/)

### TERR-SPR-067

**The unit BODY places itself one term away from its own shadow — the rule `TERR-SPR-048` published is the shadow's.** `TERR-SPR-065` moved the body to `vt+0x28` = `R0552`; this row reads that routine's destination, anchor and frame, because `TERR-SPR-048` had read all three out of `R0553`. **(a) Destination.** The body ignores arguments 1 and 2 exactly as the shadow does (frame `+0x58`/`+0x5c`, no reference anywhere in the body) and builds its destination from the unit's own fields — but with an extra subtrahend: **`dstX = unit+0x60 − anchorX`**, **`dstY = unit+0x64 − anchorY − unit+0x10 − unit+0x68`** (`L10428…L10429`; the two position values passed as arg1 and arg2 at `L10430`/`L10431`), repeated identically at all four of its draw sites (`L10432…L10433`, `L10434…L10435` for the `unit+0x194`/`+0x198` overlay sheets; `L10428…L10429`, `L10436`, `L10437` for the `+0x04`/`+0x08` main sprites). The shadow's is the same expression **without** `unit+0x10` and **with** a `− sunShear` on `dstX` instead (`L07949…L07950`, the shear loaded at `L07951` -- **AMENDED by [EXP-0129]: that address loads the *pixel* term `ftol(tan(theta)*(2*floor(frameH/2) - anchorY))`, not the 16.16 slope, which lives at frame base+0x18 and is pushed as blit argument 5 at `L07953`; see `TERR-SHDW-130`, `TERR-SHDW-131`, retracted.md**). So on the same unit and the same frame the two destinations differ by `(sunShear, −unit+0x10)` — a renderer that uses one rule for both draws the shadow under the wrong feet or the body on the ground. **(b) Anchor: unchanged and shared.** Both routines compute `anchorX = frameW/2 + (CenterX − Width/2)`, `anchorY = frameH/2 + (CenterY − Height/2)` from the same class fields — `class+0x2c` Width, `+0x30` Height, `+0x34` CenterX, `+0x38` CenterY — with `frameW`/`frameH` fetched through `vt+0x20`/`vt+0x24` at the **drawn** frame, not frame 0 (body `L10438…L10439` and `L10440…L10441`; shadow `L10442…L10331` and `L10443…L10444`). `TERR-SPR-040`'s formula and `TERR-SPR-043`'s body-pass convention hold for the body verbatim. **(c) Frame: duplicated, not shared.** Each routine has its **own** nine-arm switch on `unit+0x74` with its own jump table — body `L10445`/`L10446`/`L10447` (a jump through the table at `L10320` indexed by `4 × state`), shadow `L10315`/`L10448`/`L10449` (a jump through the table at `L10316` indexed by `4 × state`) — and the two tables have the same shape arm for arm (states 2 and 4 share one, states 3/7/8 share one, six distinct each). Compared as multisets of instruction shapes with branch targets, calls and pushes removed and registers normalised (`tools/ghidra/armcmp2.py`, whose order-blindness is stated there): **states 1 and 3/7/8 are identical**; states 0, 5 and 6 differ only in *where* the direction `(unit+0x6c − 8) >> 1 & 7` is formed — the body reloads it inside each arm, the shadow uses the copy its prologue made, and both prologues form it identically (body `L05809`/`L10450`/`L10451`/`L05810`) — plus the shadow's hoisted `[L02113]` load in state 6. Both state-0 arms read `unit+0x15a` and `class+0x94`, so `REG-UNITS-050`'s corpse substitution is in both. The **default arm (states 2 and 4) is the one genuine divergence**: the shadow draws the class id (`L10318`, from a slot written once from `+0x20` of the unit — `TERR-SPR-047`), the body draws **its own third argument**, i.e. the light level (`L10321`, a load of the stack slot at offset `0x60`). Both are nonsense as frame indices, which is the same evidence `TERR-SPR-047` used to suspect those two states are never reached, now doubled. **(d) Which field is which.** `EnumRefs` with regexes for accesses at `0x10` and `0x68` from the register holding `this` — 176 hits, 109 owners, 1 in orphan code and not in this block; blind spot: accesses through that register only, which suffices here because every method in `L05122..L05123` is `__thiscall` and copies `this` into that register. On `this`, inside that block: `unit+0x68` is read by all four drawing routines (`R1405` vt+0x38, the body ×5, the shadow ×4, `R1106` the bars ×1) and `unit+0x10` by **three of the four — every one except the shadow** (the shadow's two reads of `+0x10` at `L10452`/`L10453` follow a reload of that register from the sprite table and are `spriteEntry+0x10`, the lazy-load flag). No instruction inside the block writes either field. So `+0x68` belongs to the plane the shadow lies in and `+0x10` lifts everything drawn *at* the unit off it; `TERR-SPR-041`'s `OffsetRect` picking rectangle subtracts both, i.e. it bounds the body

**Confidence.** High (every term, both jump tables and the anchor field set are named instructions; the arm comparison is a committed instrument with its blind spot stated; the reader split is an image-wide enumeration with its instrument and blind spot stated) / **Medium** (that `+0x10` is a height above ground and `+0x68` a ground-plane term — the reader split and the picking rectangle both point that way and no instruction says so; the discriminator is a writer for either field, which does not exist in this block)

**Original status.** ● active (amended)

**Evidence.** [EXP-0053](../experiments/EXP-0053-sprite-lighting/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-STRUCT-068

**How a placed structure makes its footprint impassable — and it is BOTH planes, from the object's own constructor, not the map ingest.** The sim class is the one the image names `Building` (`CRuntimeClass L10454`, `0x6c` bytes, vptr `R0492`; `Shop` derives from it — `SHOP-CLS-001`), and the chain is `R0485`/`R0693` (its ctors) → **`R0486`** (resolve the class in the `Data.bin` Buildings table and take the footprint; `L10455` (the call to `R0488`) with `this = [L04624]`, the world, **no null check**) → **`R0488`** (walk the `w×h` rectangle) → **`R1357`** (attach one cell: `L07898`/`L07899` (a store of `obj` into `+0xc`) into the cell record's payload) → **`R0453`** (recompute that cell). The recompute writes the pair **bit 0 \| bit 2** into the **static** plane `world+0x10000` *and*, three instructions later, into the **dynamic** plane `world+0x20000`; both arms are structured the same way, one load of the operand and then a read-modify-write per plane — set: `L10456` (loads the operand `0x5`) → `L10457` (ORs it into the static byte)/`L10458` (stores it at `+0x10000`) and `L10459` (ORs it into the dynamic byte)/`L10460` (stores it at `+0x20000`); clear: `L10461` (loads the operand `0xfa`) → `L10462` (ANDs the static byte with it)/`L10463` (stores it at `+0x10000`) and `L10464` (ANDs the dynamic byte with it)/`L10465` (stores it at `+0x20000`) (`TERR-STRUCT-071`). Call-site counts, `EnumRefs callto:` on the repaired table: `R0486` 2 (both ctors), `R0488` **1**, `R1357` **1**, `R1358` **1** (the destructor) — so there is exactly one route in and one route out. `R1357` refuses a cell that already carries a building (`L10466` (the non-zero branch) → return 0) and `R0488` **aborts the remaining footprint** on that refusal (`L10467`); over 38 maps / 3141 placements / 17 057 attached cells that fires **2** times. **AMENDED by [EXP-0081]: the last hop of that chain is CONDITIONAL, and the arrow above hides it.** `R1357` calls `R0453` exactly once, at `L07176`, on the branch that had to **create** the cell record; on the branch where a record already existed and was free it stores the building at `L07898` and returns 1 at `L07900` **without writing a plane byte** (`TERR-STRUCT-076`). 0 of the 17 057 shipped attachments take that branch, because the ingest creates no records, so nothing this row measured changes — but "attach one cell → recompute that cell" is not what the routine does. **The named rival — "buildings are baked into the static plane at map load, exactly like water" — is refuted for structures and true for scenery**: `R0278`'s type-3 arm does bake the `.alm` object layer with the same two bits (`TERR-PASS-049`(e)), but no arm of it is a structure arm

**Confidence.** High (each hop is a named instruction with its immediate; the both-planes fact is the two stores' own displacements; the rival is excluded by *absent* instructions in a complete reference enumeration of both planes — `EnumRefs disp:10000` 78 hits / 33 owners and `disp:20000` 60 / 19, **0 in orphan or undisassembled code**, with every `LEA` hit followed to its store and both address-of-cell getters' callers enumerated, `R1349` 1 and `R1351` 2)

**Original status.** ● active, as amended

**Evidence.** [EXP-0070](../experiments/EXP-0070-structure-passability/), **[EXP-0081](../experiments/EXP-0081-bridge-crossing/)**

**Amended.** The table ledger carried the status "● active, as amended". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-STRUCT-069

**The cell record: the plane byte is a cache, the record is the storage.** `world+0x540b4` is an MFC map of `0x34`-byte payloads keyed by the `u16` cell index (`R1537` find, `R1596` new node, payload at `node+0xc`; `world+0x5402c` is the scratch every toucher copies it through, which is why no payload field has an absolute displacement anywhere in the image). Payload map, read off `R0453` and `R1357`: **`+0x00`** the cell's terrain-baseline **cost** byte and **`+0x01`** its terrain-baseline **static block** byte, both snapshotted from the live planes when the record is created (`L10468`, `L10469`); **`+0x04`** a ground occupant → dynamic bit 6; **`+0x08`** an air occupant → dynamic bit 7; **`+0x0c`** the **building**; `+0x10` set/cleared by `R0938`/`R0447` and read by neither plane arm; `+0x14…+0x28` six pointers, each non-null one **quadrupling** the cell's cost byte in place (`L05495` (a shift left of the byte by `0x2`)), of which **`+0x20`** additionally ORs the same block pair; `+0x2c` a byte. `R0453` **returns immediately if the cell has no record** (`L10470` (the zero branch)), and otherwise rebuilds both plane bytes from scratch: assign the two baselines (`L10471`, `L05492`), OR-with-`0x20` into both, then re-apply occupants, the building, `+0x20`, and the saved bit 4 (`L10472` saves it, `L10473`/`L10474` restore it). So a plane byte is a pure function of (terrain baseline, occupants, building) recomputed per cell, and the detach restoring `+0x00`/`+0x01` (`L07061`, `L10475`) is what makes removing a building give the terrain — water included — back exactly

**Confidence.** High (the field map, the two snapshots, the recompute order and the early-out are named instructions; the container is identified by its own find/insert helpers) / **Unknown** (what fills `+0x14…+0x28`, and therefore whether the cost-×4 and the `+0x20` block arm are live. The instrument that would find a writer is not `disp:` — no payload offset survives as a displacement, all access is through the `LEA world+0x5402c` base, whose 28 owners were enumerated and read; a writer could still store into a hash **node** at `node+0x20`, which no sweep here covers)

**Original status.** ● active

**Evidence.** [EXP-0070](../experiments/EXP-0070-structure-passability/)

### TERR-STRUCT-070

**A structure's footprint is a rectangle plus TWO 32-bit masks, and it is not the mover's `n×n`.** `R0486` copies, from the class's `Data.bin` Buildings entry, **param 0 → `obj+0x60`** (width) and **param 1 → `obj+0x61`** (height) at `L10476`/`L10477`, **param 4 → `obj+0x64`** at `L10478` and **param 5 → `obj+0x68`** at `L10479`. The table ships its own column titles (`DAT-GRAM-003`) and names them: `sizeX`, `sizeY`, **`Passability`**, **`BuildingPresent`**. `obj+0x68` is the **attach set** — `R0488` @`L10480` tests it and only calls the attach for a set bit; `obj+0x64` is the **blocking set** — `R0453` @`L07901` tests it and picks OR-with-`5` or AND-with-`0xfa`. Bit index is `(row − objRow)·w + (col − objCol)` with the anchor at the footprint's top-left (~~`obj+0x10` byte 0 = col, byte 1 = row~~ — **REFUTED by [EXP-0081]: `obj+0x10` is a POINTER** to a 12-byte position object, and the two bytes are `[*(obj+0x10) + 0]` and `[+1]`; the arithmetic is unchanged, the field named is not, `TERR-STRUCT-075`), and the `SHL`'s count is masked to 5 bits, so a footprint over 32 cells **aliases**: the 11×4 `Castle`, the only shipped entry with `w·h > 32`, folds bits 32..43 onto 0..11 (87 aliased cells over 7 placements). The caller-supplied `(w,h)` arm (`ALM-CLS-036`'s `kind == 0x21` extension) sets `obj+0x68 = 0xffffffff` and **`obj+0x64` = 0 on both sides of its own `w·h > 32` branch** (`L10481` and `L10482`, the same immediate) — an extension-sized object occupies its whole rectangle and blocks none of it. Corpus, 66 shipped entries: `Passability ⊆ BuildingPresent` on **66/66**, `BuildingPresent` = the full rectangle on **64 of the 65** with `w·h ≤ 32` (`Cave` 3×3 is the exception, 63 not 511), and **31 entries have at least one occupied-but-not-blocking cell**

**Confidence.** High (the four parameter reads and the two mask tests are named instructions; the column titles are the shipped file's own; `TERR-MOVE-055`'s `n×n` is excluded because `vt+0x1c` is never called on this path) / **Medium** (that the `.alm` type-4 `X`/`Y` reach `obj+0x10` unmodified: the instruction fixes only the object-local half — `objCol + col`, `col ≥ 0` — and the corpus prefers the top-left anchor over a centred one by putting **789 of 1152 bridge-deck cells (68.5 %) on water against 50.0 %**, which is what a deck running onto both banks predicts (8/12 = 66.7 %) but is not decisive; the store into `obj+0x10` was not read) — **CLOSED to High by [EXP-0081]**, which read that store (`L02083`) and the whole chain behind it down to the loader's `L02090`/`L02091 SAR ...,0x8`; the anchor is now carried by absent instructions (nothing subtracts `(w−1)/2`) and not by the 68.5 % figure

**Original status.** ● active, as amended

**Evidence.** [EXP-0070](../experiments/EXP-0070-structure-passability/), **[EXP-0081](../experiments/EXP-0081-bridge-crossing/)**

**Amended.** The table ledger carried the status "● active, as amended". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-STRUCT-071

**Something SUBTRACTS from the block plane: a footprint cell whose `Passability` bit is clear has bits 0 and 2 erased on both planes — this is how a bridge crosses water.** `R0453`'s other arm loads the mask once and applies it **twice, once per plane**, exactly mirroring the OR-with-`5` pair above it: `L10461` (loads the operand `0xfa`) (`b1 fa`), then **`L10462` (ANDs the static byte with it)** (`22 d1`) → `L10463` (stores it at `+0x10000`), and **`L10464` (ANDs the dynamic byte with it)** (`22 d1`, the same immediate still live in that register) → `L10465` (stores it at `+0x20000`). There is no AND-with-immediate-`0xfa` instruction anywhere in the routine — `0xfa` reaches the byte only through a register, which is what a collapsed reading of this block loses. The same arm assigns the cell's movement cost from **`world+0x5417b`** (`L07045`/`L07046`) — the fifth entry of the `world+0x54177..0x54180` cost table, i.e. **`CostCracked`**, shipped value **6** (`ALM-TERR-043`); that global has exactly two referrers, this reader and the world constructor's writer `L10483` (`EnumRefs disp:5417b`). Because the recompute restores the record's terrain **baseline** first, the clearing is against whatever the ingest had put there, water included. Corpus, 38 maps / 3141 type-4 placements: **1113 cells are opened**, of which **949 were water**, 145 a type-3 object and 19 Mountain — and **0** were the tile-word bit-13 flag or the border, neither of which any structure sits on. **908 of the 1113 come from the four bridge classes**, whose shipped masks are literally deck patterns: `Horisontal Bridge` 6×4 = two solid rows with two open rows between them, `Vertical Bridge` 4×6 = solid outer columns with two open columns down the middle, and the variable-size `Vertical Wooden Bridge` opening its whole extent. The remainder are **doorways** — `Church` 4×3 opens two of its front corners, and 31 of 66 entries open at least one cell. A consumer that treats the block plane as monotone (terrain OR structures) is wrong on 1113 shipped cells and cannot place a bridge at all

**Confidence.** High (both stores, the mask test that selects between the arms, and the cost source are named instructions with their immediates; the corpus figure is the pass reproduced cell for cell against the ingest planes, 0 footprint cells off-plane; the swapped-column rival is excluded by call order — `R0488` tests `obj+0x68` before `R0453` ever sees `obj+0x64` — not by the counts; **every address in this row was re-verified against the raw PE bytes with no disassembler in the loop**, and the block's extent is pinned twice over by its own branch displacements: `L07902 74 20` (`JZ`) lands on `L07903` and `L10484 eb 28` (`JMP`) on `L05493`)

**Original status.** ● active (citation corrected in place 2026-07-31 — the assertion did not change)

**Evidence.** [EXP-0070](../experiments/EXP-0070-structure-passability/)

**Amended.** The table ledger carried the status "● active (citation corrected in place 2026-07-31 — the assertion did not change)".

### TERR-STRUCT-072

**When it happens, and that it is undone.** The attach runs **inside the constructor**, once, and never again: `R0485` @`L02200` and `R0693` @`L10485` are `R0486`'s only two call sites (**amended by [EXP-0091]: two call sites, THREE callers — the `Shop` ctor `R0490` reaches the first of them by calling `R0485` at `L02198`, `ALM-CLS-063`**), and `R0486` skips the whole thing when `[obj+0x40] == 0` or the class id exceeds the Buildings table's count (`L10486`, `L10487`). Nothing in the tick loop touches it. **The ingest strictly precedes it**, in one routine and by a dependency, not by convention: `R0128` does `L00237` (pushing `0xa4558` and calling `new`) → `L00242` the world ctor (→ `R0116` → `R0278`/`R0470`) → `L10488` (storing `world` into `[L04624]`) → `L09682` the `.alm` type-4 walker `R0489`, whose `Building` ctors read that same global with **no null test** — and the ingest's four block writes are *assignments*, so a structure attached first would be erased. **Removal:** `~Building` `R1833` @`L10489` calls `R1358`, which walks the same footprint, clears `payload+0x0c`, recomputes each cell, and when the payload has gone empty restores the two baselines (`L07061`, `L10475`) and frees the record. **Save/load does not re-attach and does not need to:** `Building::Serialize` `R1585` stores and loads `+0x40/+0x42/+0x44/+0x46/+0x48/+0x60/+0x61/+0x64/+0x68` and calls neither `R0486` nor `R0488`, and MFC's `CreateObject` ctor `R1834` constructs by the name `"null"` (`L10490`), which misses the table and skips the attach — while the world's own `Serialize` `R1360` calls the **cell-record map's** Serialize on both branches (`L10491` storing, `L10492` loading) beside `TERR-PASS-053`'s two plane sweeps. A structure cell reads `0x25`, above that sweep's `> 0x0f` threshold, so both the byte and the record travel — which is exactly what `SAV-CELLREC-017` measured from the other side, `u16 key + 0x34` payload records whose key set equals the static-plane bit-5 set with symmetric difference **0 in all four saves**

**Confidence.** High (the two call sites, the two guards, the load ordering with its global publication, the destructor call and the Serialize field list are named instructions; the record map's presence in the stream is independently attested by `SAV-CELLREC-017`) / **Medium** (that the serialized payload restores `payload+0x0c` as a live object pointer — the map's element Serialize was not read, only the two calls to it)

**Original status.** ● active

**Evidence.** [EXP-0070](../experiments/EXP-0070-structure-passability/)

### TERR-PASS-073

**The two block planes' complete writer list, the instrument, and where the bit semantics stop.** Instrument: `EnumRefs disp:10000` (78 hits / 33 owners) and `disp:20000` (60 / 19) on the repaired function table, **0 hits in orphan or undisassembled code**, with the blind spot closed by hand — a store through a pointer the routine `LEA`'d first carries no displacement, so all 12 static and 8 dynamic `LEA` hits were followed to their stores, and the two address-of-cell getters (which hand the plane pointer back as a *return value*, the shape no displacement sweep can see) had their callers enumerated: `R1349` 1, `R1351` 2. Writers of **`+0x10000`**: `R0278` (4 assignments — the ingest arms), `R0470` (border `0x1f`), `R0116` (memset 0, and the `REP MOVSD` copy to `+0x20000`), `R0453` (5), `R0057` (2), `R1357` (OR-with-`0x20`), `R0938` (OR-with-`0x20`), `R1358` (baseline restore), `R1360` (Serialize load). Writers of **`+0x20000`**: `R0116`, `R0453` (7), `R0057` (2), `R0058` (OR-with-`0x40`/OR-with-`0x80`), `R1332` (AND-with-`0x7f`/AND-with-`0xbf`), `R1360`. **Bit semantics as the instructions carry them:** bit 0 blocks ground, bit 1 blocks air (border only), **bit 2 = an object/building occupies this cell — set and cleared only ever together with bit 0, as the literal `5` and the mask `0xfa`**, which is what makes `TERR-MOVE-057`'s mask `0x44` (Ghost, Bee) mean *passes terrain, stopped by buildings*; bit 5 = the cell has a record, set at `L10493`, `L10494`, `L10495`; bits 6/7 dynamic-plane occupancy. **Where they stop:** bits 3 and 4 still arrive only inside the border literal `0x1f` — bit 4 is deliberately carried across every recompute (`L10472` → `L10473`/`L10474`) and re-ORed by `R0057` (`L10496`/`L10497`) and `R1358` (`L10498`/`L10499`), so it is a live runtime flag with no located origin. Bit 5 is cleared on the **static** plane when a record is freed (the baseline assignments `L10500`, `L10475`) — corroborated from outside the image by `SAV-CELLREC-017`, whose serialized record keys equal the static bit-5 set exactly in all four saves; no instruction in either sweep restores the **dynamic** baseline, so whether the dynamic bit 5 survives a freed record is **not** established here — what a clearing instruction would look like is an `AND` with `0xdf` (0 such hits on either plane, `EnumRefs text:0xdf`) or an assignment into `+0x20000` (3 sites, all listed above)

**Confidence.** High (the writer lists, with the instrument and both of its blind spots named and closed; the bit-2 semantics are the OR-with-`5`/AND-with-`0xfa` pair itself) / **Unknown** (bit 3's origin outside the border; the dynamic plane's bit 5 after a record is freed). *(Bit 4's origin located 2026-08-15 by [EXP-0174], `TERR-PASS-148`: `R1080` writes it from the area module's per-cell add, gated on the inner effect being a damaging one, and its only consuming reader is the RLE serializer `R1655`. No mover mask contains it. **Both writer lists above are incomplete**, and the two causes are different. `R1080` builds each address with `SHL`/`ADD` and stores through a bare register (`L07987`, `L07988`), so it is neither a displacement hit nor an `LEA` hit and escapes the hand-closed blind spot as well. The others were **`LEA` hits this row's own instrument produced and the sentence dropped**: static — `R1075` (`L05450` (an OR of `0x20` into the byte), a fourth bit-5 setter), `R0227` (`L07989`), `R0447` (`L07990`); dynamic — `R0257` (`L07011`) and `R0046` (`L07991`, `L07992`), both of which `MOVE-PLANE-005` already listed as writers in another ledger. `claims/retracted.md`.)*

**Original status.** ● active (amended)

**Evidence.** [EXP-0070](../experiments/EXP-0070-structure-passability/), [EXP-0174](../experiments/EXP-0174-area-movement/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-STRUCT-074

**A bridge deck is free to a ground mover, and the crossing is load-bearing — the model's own prediction, measured.** `TERR-STRUCT-071`'s clearing arm is the whole mechanism, re-read here instruction for instruction: `L07903` (a load of the static byte at `+0x10000`) · `L10461` (loads the operand `0xfa`) · `L10462` (ANDs the static byte with it) · `L10463` (a store at `+0x10000`), the same three again for `+0x20000` at `L10501`/`L10464`/`L10465`, then `L07045` (a load of the byte at `+0x5417b`) · `L07046` (a store of it into the cell's cost byte) for the cost. A ground mover's mask is `0x41` (`R0471` @`L06846`), so bit 0 cleared is a free cell, and the search relaxation tests the same byte with the same mask on the same plane (`L10502` (a test of the plane byte at `+0x10000` against the mask)), expanding the **full 3×3 with no corner rule** — `L10503` (an increment) / `L10504` (an increment) against a comparison with `0x2`, the destination cell tested alone. **What was never measured:** counting opened cells is not crossing. Over the 38 shipped maps, 8-connected reachability under mask `0x41` — free cells **449 806 → 439 164** (blocked **430 898 → 441 540**, the two figures `TERR-PASS-051` carries); bridge-deck cells free **391 of 1 299 → 1 299 of 1 299**; and **92 of the 104 bridge placements join two ground components the ingest stage leaves disjoint** (largest component: `Islands.alm` 22 189 → 35 135, `Tomb.ALM` 15 470 → 33 700, `Cross.ALM` 15 549 → 39 575, `scn:121.alm` 516 → 1 814). Under the ingest plane alone **no** bridge joins anything. The bit index reaching `L10505` (a shift of `1` left by the bit index) is correct only **modulo 256** — `L10506` (a subtraction) and `L10507` (a multiplication) are 32-bit operations on registers still holding the cell index — and that suffices because `SHL` masks its count to 5 bits; a consumer gets the right answer from the published formula for the reason of the mask, not of the arithmetic

**Confidence.** High (the deck arm, the mask, its setter and the relaxation's neighbourhood are named instructions with their immediates; the two readings that leave a deck blocked — "no clearing arm" and "every attached cell blocks" — are separated by the corpus at **0** bridge joins each against 93, and by the `TEST`/`JZ` pair) / ~~**Medium** (the *polarity* rival: reading a SET `Passability` bit as opening scores **89** joins against 93 … only the test at `L07901` and the zero branch at `L07902` separate them)~~ — **RAISED TO High by [EXP-0085]**, whose reading also shows this row's *reason* for the Medium to have been wrong twice over. The `JZ` is `74 20`: its `rel8` fixes the target at `L07903`, the block that begins `b1 fa` = the load of `0xfa`, so **bit CLEAR opens and bit SET blocks** (`TERR-STRUCT-078`) — the shipped title `Passability` is indeed inverted. And the corpus *does* separate the readings; bridge joins were the wrong statistic, near-degenerate because a bridge is nearly symmetric under the two. Over the 66 Buildings rows the code's reading puts **35 classes blocking every cell they attach and 2 blocking none**, the rival **2 and 35**, with `Well`, `Tower`, `Tower 2` and `Cave` walk-through. "The corpus does not separate them" was a statement about the statistic measured, not about the corpus / **Medium** (the connectivity figures as a statement about *play*: they are the plane's own reachability, with no occupancy, triggers or mission gates)

**Original status.** ● active (amended)

**Evidence.** [EXP-0081](../experiments/EXP-0081-bridge-crossing/), **[EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/)**

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### TERR-STRUCT-075

**`obj+0x10` is a POINTER to a 12-byte position object — not two coordinate bytes — and with it the footprint anchor is read from the image instead of fitted to the corpus.** The base actor constructor allocates it and stores its address: `L07893` (pushing `0xc`) · `L07894` (the call to `R0747`) (operator new) · `L07895` (the call to `R1010`) (its copy ctor) · **`L02083` (the dword store into `+0x10`)**. Its shape is fixed by `R0288`, which writes every byte of it: `L02175` (a byte store at `+0x0`) = **`+0` col**, `L10508` (a byte store at `+0x1`) = **`+1` row**, `L10509` (a 16-bit store at `+0x2`) = **`+2` the packed `(row<<8)` + col** the cell-record map is keyed by, `L10510`/`L10511` (a byte store at `+0x4` / `+0x5`) with `L00967` (loading `0x80`) = the **half-cell sub-position**, `L10512` (a dword store at `+0x8`) = the world. Both footprint routines dereference it: `L07896` (a load of the dword at `+0x10`) then `L02176` (a byte load at `+0x0`) / `L02177` (a byte load at `+0x1`), and `L07897` (a load of the dword at `+0x10`) then the bytes at `+0x0` / `+0x1` of it. `R0488` adds `col` in `[0,sizeX)` and `row` in `[0,sizeY)` to those bytes and **nothing anywhere on the chain subtracts `(w−1)/2`**, so the anchor is the object's own cell and the footprint runs right and down — top-left, by absent instructions. The `.alm` half closes too: `R0489` pushes the in-memory type-4 record's `+0x00` and `+0x02` into `R0288` (`L02171` (a byte load at `+0x0`), `L02170` (a byte load at `+0x2`)), and the record loader divides both `u32` file coordinates by 256 first — `L02090` (a shift right by `0x8`) and `L02091` (a shift right by `0x8`) before `L02168` (the call to `R0481`). **This retires `TERR-STRUCT-070`'s Medium clause** ("the store into `obj+0x10` was not read"; the corpus's 68.5 %-vs-50.0 % deck-on-water preference no longer carries the anchor) **and refutes that row's parenthesis** "`obj+0x10` byte 0 = col, byte 1 = row" — the arithmetic it produced is unchanged, the field it named is not what it said

**Confidence.** High (the allocation, the store, the copy constructor, the two dereferences and the two `SAR`s are named instructions with their immediates; the centred rival is excluded by an absent subtraction over a chain read end to end) / **~~Medium~~ → **High as amended by [EXP-0091]** (which of `L02090`/`L02091` is the low axis: this row's "the instructions do not say, the corpus does — 0 of 17 057 footprint cells land off-plane" is **withdrawn on both halves**. The instructions do say — the chain from the two `LEA`s that filled those locals to the cell key `R0488` packs has no permutation in it, `ALM-OBJ-061`. And the off-plane statistic could never have decided it either: 37 of 38 shipped maps are square, so it is 0 under both readings on 37 of them. Re-made as connectivity the corpus *does* separate them — bridge joins 93 vs 5)

**Original status.** ● active (contested — `registry.md` C-7 against `ALM-OBJ-019`'s destination for the `.alm` `+0x12` field; the *pointer* is not what is contested, the other row's destination is)

**Evidence.** [EXP-0081](../experiments/EXP-0081-bridge-crossing/)

**Amended.** The table ledger carried the status "● active (contested — `registry.md` C-7 against `ALM-OBJ-019`'s destination for the `.alm` `+0x12` field; the *pointer* is not what is contested, the other row's destination is)". `retracted.md` records a correction against this claim, which adds `superseded`. The `contested` qualifier is kept: the table ledger's status named a contest (`registry.md` C-7 against `ALM-OBJ-019`'s destination for the `.alm` `+0x12` field), and no record in this ledger or in `retracted.md` closes it.

### TERR-STRUCT-076

**The attach chain reaches the recompute on ONE of `R1357`'s two paths, not unconditionally.** The routine first walks the record map's bucket chain inline (`L10513` (an unsigned division by the dword at `+0x8`), `L10514` (a comparison of the 16-bit value at `+0x8`)). If the cell has **no** record (`L10515` (a jump to `L10516` when equal)) it creates one — zeroed payload (`L10517` (zeroing a value) / `L10518` (a repeated dword store)), snapshots the two terrain baselines (`L10468` (a byte store at `+0x0`) cost, `L10469` (a byte store at `+0x5402d`) static block), sets `L10493` (an OR of `0x20` into the byte), stores the building at `L07899` (a dword store at `+0xc`) and ends **`L07176` (the call to `R0453`)**. If the cell **already has a record** whose `payload+0x0c` is null, it stores the building at `L07898` (a dword store at `+0xc`), copies the payload back (`L10519` (a repeated dword copy)) and returns 1 at `L07900` — **writing no plane byte at all**. On that path a bridge deck stays blocked until something else recomputes the cell. The recompute's caller set is also larger than the two routines `TERR-STRUCT-068`/`069` name: **`EnumRefs callto:R0453` on the repaired function table — 15 hits over 11 distinct owners, 0 in orphan or undisassembled code** — `R0458` (4 sites), `R0057`, `R0938`, `R0447`, `R1357`, `R1358`, `R1075` (2), `R1078`, `R0227`, `R1359`, `R1835`

**Confidence.** High (both exits, the guard, the baseline snapshots and the single call are named instructions; the enumeration states its instrument and its blind spot — `callto:` sees `.rdata` code pointers as well as call references and is blind only to a target reached through a computed pointer) / **Unknown** (whether the no-recompute path is reachable in play: at map load no record exists before the type-4 walk, so **0** of the 17 057 shipped attachments take it, and what might create a record mid-session on a cell a building is later placed on was not traced)

**Original status.** ● active

**Evidence.** [EXP-0081](../experiments/EXP-0081-bridge-crossing/)

### TERR-STRUCT-077

**The `.alm` extension kind is a bridge class, and it opens its whole rectangle.** `kind == 0x21` — the type-4 extension discriminator (`ALM-OBJ-019`) — is class id **33**, whose shipped Buildings row is named **`Vertical Wooden Bridge`**. For that kind `R0486` takes the caller-supplied `(w,h)` instead of the table's (`L02193` (a byte store at `+0x60`), `L02194` (a byte store at `+0x61`)) — **amended by [EXP-0091] on the gate and on which is which: the arm is selected by the SUM of the two caller bytes at `L07853`…`L07854`, not by the kind (`TERR-STRUCT-090`), and the first caller byte is file `+0x14` while the second is file `+0x18` (`ALM-OBJ-062`)**, writes `0xffffffff` into both masks (`L10520`, `L10521`) and then **overwrites `obj+0x64` with 0 on both sides of its own `w·h > 32` branch** — `L10481` and `L10482`, the same immediate, so `L10522` (a comparison with `0x20`) decides nothing and the branch is dead. `BuildingPresent` stays `0xffffffff`, so every cell of the rectangle attaches, and `Passability = 0` sends every one of them to the clearing arm: an extension-kind object is an **author-sized hole in the block plane**, and because both masks end all-ones and all-zeros the `SHL`'s 5-bit aliasing cannot affect it at any size. Corpus: **8 placements** over `Cross.ALM` (3), `Horror.alm` (4) and `scn:81.alm` (1) — the same 8 the independent type-4 walk finds extensions on — **147 deck cells, 28 free before the pass and 147 after, and all 8 join two ground components the ingest separates** (`TERR-STRUCT-074`). So the corpus carries **two independently-sourced bridge mechanisms**: the Buildings table's `Passability` column (`Horisontal Bridge` 49 placements, `Vertical Bridge` 47) and this arm

**Confidence.** High (the six stores, the dead branch's two identical immediates and the class-id identity are named instructions and the shipped table's own row; the corpus count is exact and agrees with the extension count taken from the `.alm` walk) / Medium (that all 8 extension records in the corpus are class 33: it follows from the discriminator being the kind itself, but no record was decoded field by field here — **closed by [EXP-0091]: all 8 are decoded, all are class 33, and all carry two non-zero extent bytes**)

**Original status.** ● active (amended)

**Evidence.** [EXP-0081](../experiments/EXP-0081-bridge-crossing/), [EXP-0091](../experiments/EXP-0091-placement-axes/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-STRUCT-078

**The recompute's polarity, decided by one branch displacement — a SET `Passability` bit BLOCKS, and the shipped column title is the inverse of what the code does with it.** `L10523` (loading `0x1`) · `L10524` (copying the bit index) · `L10505` (a shift of `1` left by the bit index) makes the value exactly `1 << bitIndex`, so `L07901` (a test of the blocking set at `+0x64` against that bit) sets ZF iff the mask bit is **clear**. `L07902` is `74 20`: `rel8` fixes the `JZ` target at `L10525 + 0x20 = L07903` under any decode, and `L10484 eb 28` fixes `L05493`, so the fall-through block is exactly 32 bytes and the taken block exactly 40 and neither can be misaligned. The taken block begins with a load of the byte at `+0x10000` / the load of `0xfa` — and the fall-through begins `b2 05` / `0a da`. **Bit CLEAR → AND-with-`0xfa` on both planes plus `cost := CostCracked`, the cell is OPENED; bit SET → OR-with-`5` on both planes, the cell BLOCKS.** Every byte cited here was re-read out of the PE section table with no disassembler in the loop (`evidence/rom-bytes.txt`). **Note where this repo already stood:** `formats/terrain/format.md` prints this polarity as flat pseudocode while `TERR-STRUCT-074` graded it **Medium** — the spec was ahead of the ledger and neither said so, which is a second thing a consumer weighing grades would have got wrong. **The rest of `R0453`, whose one line the spec flags `UNSOURCED`, is now sourced**, in execution order: the cell record is looked up in the map at `world+0x540b4` and **13 dwords (52 bytes) of its payload are `REP MOVSD`-copied into the scratch at `world+0x5402c`** (`L10526`/`L10527`) — a miss returns with no plane write at all; bit 4 of the *existing* static byte is saved (`L10472` (a mask with `0x10`)); the snapshot's byte 1 becomes the static block byte and the snapshot's byte 0 the cost (`L10471`/`L05492`) — this is the terrain **baseline**; OR-with-`0x20` (bit 5, "this cell has a record") goes to the static byte and the result is **copied wholesale into the dynamic plane** (`L10495`…`L10528`), which is where the dynamic plane's initial value comes from; scratch `+0x4` non-null sets dynamic bit 6 and `+0x8` non-null dynamic bit 7 (`L10529`, `L10530`); scratch `+0xc` is the **building pointer** and, when non-null, runs the footprint arm above; then **six** dwords at scratch `+0x14`…`+0x28` each quadruple the cost byte (a shift left by `0x2`) when non-null (`L10531` (a comparison of the dword with zero) / `L05495` (a shift left of the byte by `0x2`) / `L10532` (an add of `0x4`) / `L10533` (a decrement) / `L07044` (the loop branch), the loop counter loading the count `0x6` at `L05494`), of which **`+0x20`, the fourth, additionally ORs 5 into both planes** (`L05481`…`L05488`) — the `formats/terrain/format.md` line flagged `UNSOURCED`, now carried by those instructions; finally the saved bit 4 is ORed back into both planes (`L10534`…`L10474`), so **bit 4 is the one block bit a recompute cannot destroy**. Corpus, 66 shipped Buildings rows: under this reading **35 of 66 classes block every cell they attach and 2 block none**; under the rival those two figures **swap to 2 and 35**, and `Well`, `Tower`, `Tower 2` and `Cave` open their entire footprint. `Passability` a subset of `BuildingPresent` reproduces 66/66

**Confidence.** High (the two branch targets are fixed by their own `rel8` bytes and re-verified against the raw PE; every arm's immediate and displacement is named; the rival is excluded by the taken block's first bytes, not by a census) / Medium (the 66-row corpus figures — corpus agreement, and recorded as agreement rather than as the discriminator)

**Original status.** ● active

**Evidence.** [EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/)

### TERR-TILE-079

**Tile bits 15..14 are the FOG OF WAR, and the gate they feed is satisfied exactly where the local view currently has line of sight — so the animated arm and the partial-repaint renderer are live, not dead.** The stamp `TERR-TILE-044` names writes the immediate **`0xc0`** — bits 15 *and* 14 together — into the high byte of each of the four cells the gate ORs (`L02581`…`L02584`), so **one stamped word in a quad already carries both bits and the four-corner OR passes `& 0xc000 == 0xc000`**. Bit 14 is cleared map-wide on a period by `L01659` (a 16-bit mask with `0xbfff`) and **bit 15 by nothing at all**. The instrument for that absence, with its own reading: `imm:3fff` returns **25 hits / 16 owners / 0 orphan** and `imm:7fff` **66 / 41 / 0**; of those 91, exactly **8 have a memory destination and all 8 are whole-`dword` `MOV`s of a `0x7fff`/`0x7fffffff` sentinel** into a stack local or an object field — **not one `AND` of a 16-bit word anywhere in the image**. The two hits inside map/drawable routines were read rather than assumed: `L10535`/`L10536`/`L10537`/`L10538` in `R1518` strip the two-bit RLE opcode off a run word (a test of `0x4000` / a test of `0x8000` two instructions above), and `L10539` in `R0579` is the `CDQ` / a mask with `0x3fff` / `ADD` / a shift right by `0xe` signed-divide-by-0x4000 idiom. **Blind spot, stated:** a clear reaching the word through a register would carry no immediate on the storing instruction (`AGENTS.md` rule 5's shape), so this is an absence of *immediate* clears, not a proof that bit 15 can never be cleared. The pair therefore carries three states with a mechanism: **`00` never seen · `10` explored, not currently visible · `11` currently in sight**, which is exactly why `R1504` tests `== 0x8000` before `== 0xc000` (`TERR-TILE-044`'s three-state finding, now explained). **The stamp's guards are three, not two.** `vt+0x48` is `R0554`, whose whole body before the tail-jump is `R0554` (a comparison of the byte at `+0x15a` with `0x3`) / `L10540` (the not-below branch) — a drawable at decay stage 3 or beyond (`ANIM-DEATH-007`'s server decay stage) stamps nothing. The body at `L02578` then tests (1) `L10541` (a test of bit `0x8` of the byte at `2 × index` into the array at `[[mapView+0x9b4]+0x38]`, the index being `[drawable+0x14]+0x4`), i.e. **bit 3 of a per-player `u16` on the LOCAL participant's record, indexed by the drawable's owning `Player+0x04`**; and (2) `L10542` (a comparison of the 16-bit value at `+0x102` with zero) — the drawable's sight. `R0709` writes `flags[localPlayer] = 0x0a` at session setup (`L07802`) and its per-participant refresh **preserves bit 3** (`L07803` (a mask with `0x8`)) while recomputing bits 0..2 from three per-participant dwords; the only other writer image-wide is a message arm at `L07804`  — *amended by [EXP-0088]: there is a **third**, opcode 45's wholesale `memcpy` at `L07807`, which no store-form or `disp:` sweep can witness; the `L07804` arm is opcode 33 (`TERR-FOG-089`)*. So **in single player exactly one player's units reveal: the local one**. The per-tick caller `vt+0x44` = `R1402` repeats those three and adds two more — `(drawable+0x50 & 0x1f) == 0x10` **and** `(drawable+0x54 & 0x1f) == 0x10` (`L10543`/`L10544`), and `drawable+0x8`/`+0xc` differing from `drawable+0xc0`/`+0xc4`, the position at which it last stamped (`L10545`…`L10546`, written back at `L10547`/`L10548`) — so a standing unit re-stamps only from the map-wide pass and a moving one eight times per cell per axis. `callto:R0554` and `callto:R1402` each return **2 `.rdata` slots and 0 code callers** (the `CUnit` and `CAirUnit` vtables). **Consequence for a consumer: `TERR-SPR-042`(b)'s animated-object arm and `R1814`'s partial repaint fire on the cells the local player can currently see** — `TERR-TILE-044`'s "0 of 880 704 shipped cells" describes the *file*, and the runtime grid is a different object. A map that ships bits 15..14 set is permanently revealed there

**Confidence.** High (each guard, each immediate and the four stores are named instructions; the gate's satisfiability follows from the single `0xc0` immediate rather than from any census; the bit-15 clearer's absence is an `imm:` sweep, which does see orphan code, with its blind spot named — bytes never disassembled) / Medium (that bit 3 means "shared vision": its setter for a non-local player is one located message arm whose opcode is unread)

**Original status.** ● active (amended)

**Evidence.** [EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-FOG-080

**The stamp's 41×41 mask is not a shape — it is a per-drawable line-of-sight field, rebuilt from scratch on every stamp.** `R0284` (called at `L10549` with the drawable as its argument) `memset`s `mapView+0x17cc` to zero over **`0x1a44` = 6724 = 41·41·4 bytes** (`L10550`…`L10551`), seeds the centre `mask[20][20]` — which *is* `mapView+0x24ec`, `disp:24ec` returning **1 hit image-wide** (`L10552`) — with **`(sight >> (8 − k)) + (1 << (k − 1))`** where `k = mapView+0x3f38 = 7` and `mapView+0x3f34 = 1 << k = 128`, both written once by the `CMapView` constructor (`L10553`, `L10554`), then walks **Chebyshev rings r = 1..19** (`L00938` (a comparison with `0x14`, branching when not below)), four edges per ring, **stopping at the first ring in which every call reported blocked**. One cell is `R0285`: **`mask[dx][dy] = mask[pred[dx][dy]] − ( cost[dx][dy] + h(cell) − h(observer) )`, visible iff `> 0`** (`L10555`…`L00936`, then `L00937` (a comparison with `0x0`, branching when not above)), where `h` is the map's altitude plane `map+0x10` read **signed** (`L10556` (a sign-extending load)) and `h(observer)` is sampled once at the drawable's own cell (`L10557`). Both tables are built once by `R0283` from the constructor: **`pred`** at `mapView+0xaa8`, 41×41 `int8` pairs, the step one cell toward the observer classified into three zones by `j < i>>1` (`L00929`) and `j > i<<1` (`L10558`) and mirrored into four quadrants; **`cost`** at `mapView+0x3210`, 41×41 `int16`, **`ftol(128 · sqrt(i² + j²) / max(i,j))`** (`L10559`…`L10560`) — the Euclidean length of one Bresenham step, so the axis step is exactly **128** and the diagonal **181**, the whole field is in **1/128 cell**, and the revealed region is a **disc, not a square**. `disp:17cc` returns **4 hits over 2 owners, 0 in orphan or undisassembled code**: three in the builder, one in the stamp. Ring cells outside a **7-cell** margin (`L10561`…`L10562`, against `map+0x4`/`map+0x8`) are never called and their mask entry stays 0, so the outermost window ring is unreachable in principle — at the byte field's ceiling (255 cells) on flat ground **0 of the window's 160 edge positions** are lit. **Consequences a consumer must not miss:** the terrain term is `h(cell) − h(observer)` **evaluated per cell**, not a slope along the ray, so a constant-height plateau beside the observer costs sight (127 → 104 revealed at sight 6) while raising the whole map including the observer costs nothing (127 → 127); and because descending ground *returns* budget, altitude dominates — over 7 221 sampled observers on the 38 shipped maps a sight-6 unit reveals ~~**127 cells on flat ground, 137 on average over real terrain, 58 at the worst position and 1 154 at the best**, a 19.9× spread from standing somewhere else~~ — *every figure in this clause is **withdrawn** by [EXP-0120]: the transcription behind them stopped four instructions short of the end of `R0283` and therefore ran on a table with two cells' predecessors wrong. Flat ground at sight 6 is **145**, not 127, and the whole flat series is `TERR-FOG-117`. The recurrence, the tables' construction, the memset, the seed, the ring bound and the margin in this row are unaffected and were re-read instruction by instruction; only the region figures moved*

**Confidence.** High (the memset size, the seed, the ring bound, the margin, the recurrence and both tables' construction are named instructions with their immediates; `k` and `S` are the constructor's own two stores; the "fixed shape" rival is excluded by the `memset` at the head of every stamp) / Medium (the reveal figures over the corpus — the transcribed algorithm re-executed on shipped altitude grids, with no units, no shroud draw and no runtime session)

**Original status.** ● active (figures withdrawn)

**Evidence.** [EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/), **[EXP-0120](../experiments/EXP-0120-sight-writers/)**

**Amended.** The table ledger carried the status "● active (figures withdrawn)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-FOG-081

**The sight value that drives the fog has TWO widths, and which one you get depends on whether the actor is a hero.** The drawable's `+0x102` — the stamp's second guard and the field builder's only input — is copied from **`actor+0xa4`** by the state sync (`L10563` (a load of the 16-bit value at `+0xa4`) / `L10564` (a 16-bit store at `+0x102`)); the `CUnit` constructor leaves it **0** (`L10565`), so a drawable with no sight never stamps. For a **hero** the derive `R0280` writes it as `ftol(((mind + reaction) / 25 + 4) × 256)` (`L10566`…`L00919`), whose three FPU constants read out of the PE are **25.0, 4.0 and 256.0** — which both confirms `HERO-SIGHT-007` from the image's own data and fixes the field's unit as **1/256 cell**. For a **non-hero** the streamer `R0180` registers **`actor+0xa5`** — the *high byte* of that same `u16` — as its **slot 11** (`L00921` (an add of `0xa5`), the eleventh registration after `+0x84`, `+0x86`, `+0x88`, `+0x8a`, `+0x96`, `+0x98`, `+0x9c`, `+0x9e`, `+0x8c` and `[+0x154]+0xa`), and title 11 of the `Data.bin` **Units** table is **`scanRange`** (Humans title 9, `ScanRange`). So a monster's sight is a whole number of cells and a hero's is not, and `AI-SIGHT-006`'s "radius `actor+0xa5`" is that same field read at whole-cell resolution. Corpus: `scanRange` runs 4..12 over the 56 parameterised Units rows (7 the mode, 18 of them) and 4..7 over 210 Humans rows (6 the mode, 150) — **the ring bound of 19 is never reached by shipped data**. Because the field builder halves the value into 1/128, a hero's low bit of 1/256 is discarded

**Confidence.** High (the copy, the derive's three constants, the constructor's zero and the streamer's slot are named instructions, and the slot's ordinal is counted against the table's own column titles) / Medium (that the Units column feeding slot 11 is `scanRange`: the ordinal match is exact and the shipped values are cell-sized, but the store itself goes through a pointer-taking helper carrying no displacement, so no `disp:a5` sweep can witness it — EXP-0072 names that same blind spot)

**Original status.** ● active

**Evidence.** [EXP-0085](../experiments/EXP-0085-tile-gate-and-polarity/)

### TERR-FOG-082

**How a consumer BRANCHES on the three states, and why the fourth is unreachable.** `TERR-TILE-079` establishes the states; this is the instruction that turns them into what is drawn. `R1657` reads the tile plane `[CMapView+0x80]+0xc` per lattice vertex and classifies: `L10567` (a 16-bit load at `2 × index`) · `L10568` (a mask with `0xc000`) · `L10569` (a comparison with `0xc000`) → shroud level **0** (`L10570`), else `L10571` (a comparison with `0x8000`) → level **8** (`L10572`), else level **0x10** (`L10573`). So the mapping is `11 → 0` (**no shroud drawn at all**), `10 → 8` (half), `00 → 16` (black), and the branch a renderer needs is on the **pair**, not on either bit alone. The combination `01` — currently visible but never seen — is not a state: bit 14 is never set without bit 15, because the stamp writes the single immediate `0xc0` (`TERR-TILE-079`), and bit 15 is cleared by nothing. That is what licenses the classifier's `== 0x8000` where a reader who assumed independent bits would write `& 0x8000`, and it is why `01` needs no arm anywhere in the image. **Independent agreement with `TERR-TILE-079`'s bit-15 absence, by a different instrument**: a whole-image `re:` sweep over store forms returns a 16-bit mask with `0x7fff` **0 hits** and a 16-bit mask with `0x3fff` **0 hits**, against a byte OR of `0xc0` 4 hits / 1 owner and a 16-bit mask with `0xbfff` 1 hit. Two sweeps of different shape, the same answer **Partially retracted by EXP-0349 (ALM-TILEVIEW-122, TERR-TILECLEAR-169, TERR-DRAWSTAMP-170): raw pair 01 can survive the selected light parser. The global unreachability and no-arm-anywhere conclusions are withdrawn. The exact shroud compare/store mapping, including the default level16 for raw01, and the bounded immediate-writer observations stand.**

**Confidence.** High (three cited compares with three cited stores, read end to end). The `01`-unreachability inherits `TERR-TILE-079`'s **Medium** blind spot — a clear of bit 15 reaching the word through a register would carry no immediate on either sweep The global 01-unreachability inference is withdrawn; the local compare/store relation remains High.

**Original status.** ● partially retracted

**Evidence.** [EXP-0088](../experiments/EXP-0088-fog-of-war/), [EXP-0349](../experiments/EXP-0349-tile-word-consumers/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### TERR-FOG-083

**What the shroud pass consumes is not the tile word but a per-frame projection of it: `CMapView+0xa0`, one dword per lattice vertex, memset and rebuilt every frame.** `R1280` allocates `+0xa0` and `+0xa4` at `+0x70` bytes each (`L10574`, `L10575`), `+0x70 = (cols+7)·(rows+11)·4` (`L10576`…`L10577`); `R1836` frees them and nothing else stores either pointer. `R1657` is the **only** content writer: `memset(+0xa0, 0, +0x70)` at `L10578`, then the `TERR-FOG-082` classification per vertex, then `memcpy(+0xa4, +0xa0, +0x70)` at `L10579`. `R0379` reads it at `(r+3)·(cols+7)+(c+3)` and its three neighbours (`L10580`…`L10581`) and dispatches four ways on the four corner levels (`L10582`…`L10583`): all-equal-**0** draws **nothing at all** (`L10584`, the cell is fully visible); all-equal-**0x10** takes `R1496` for an axis-aligned quad or `R1508`; all-equal-**8** takes `R1497` or `R1510`; anything else takes `R1837`/`R1509`, `TERR-FOG-037`'s ramp, with the four levels as the last four arguments. **So the boundary is per-VERTEX, and a cell whose corners disagree is a gradient** — a consumer that stores one fog value per cell cannot reproduce any cell on the frontier

**Confidence.** High (allocation sizes, the single content writer, the memset/memcpy pair and the four-way dispatch are all cited instructions; "only content writer" is a `disp:a0` enumeration over 159 hits / 86 owners, 0 orphan, whose blind spot — an access through a pointer `LEA`'d first — is closed here by the fact that the pointer's only readers are the four routines named)

**Original status.** ● active

**Evidence.** [EXP-0088](../experiments/EXP-0088-fog-of-war/)

### TERR-FOG-084

**The shroud table `[L03341]` is 17 rows and its law is one line: `out_channel = (in_channel × (16 − L)) >> 4`. This CLOSES `TERR-FOG-037`'s Medium — level 16 is exactly black and level 8 exactly half.** `R0788` sizes the allocation `17 · stride · 2` bytes (`L07810` (a shift left by `0x4`) · an addition · a shift left by `0x1`, stride `[L03340]` = 65536 or 8192 by `[L03346]`), and fills row `L` by accumulating `channelIndex × (16−L)` and `>>4` per channel (`L07811`…`L07812`), saturating at `0xff` across the `<<(8−bits)` widening. **The discriminator is that the engine implements the same law twice more, without the table:** `R1496` fills level 16 with `PXOR MM0,MM0`/`MOVQ` (or `FLDZ`/`FISTP`) — pure zero — and `R1497` fills level 8 with `(dst>>1) & [L07813]`, that mask being built at `L10585`…`L10586` as `(0x7f >> (8−bits)) << shift` per channel. Re-executed over the **whole 16-bit pixel space** in `tools/fogprobe -mode lut`: RGB565 (mask `0x7bef`) and RGB555 (mask `0x3def`) both give **65536/65536 EXACT** on all three of level 0 = identity, level 8 = `(px>>1)&mask`, level 16 = 0. A rival curve — `(16−L)/17`, a gamma ramp, an arbitrary per-level gain table — agrees at rows 0 and 16 and fails row 8 on almost every pixel

**Confidence.** High (the row count and the fill are cited instructions, and the law is checked against **two independent implementations in the same binary** over the complete input space — the alternatives are ruled out, not merely unobserved)

**Original status.** ● active

**Evidence.** [EXP-0088](../experiments/EXP-0088-fog-of-war/)

### TERR-FOG-085

**The clock: the map-wide clear runs on the PRESENTATION tick with period 32 — neither simulation clock touches the fog — and one global can freeze it.** `TERR-TILE-079` names the clear (`L01659`) and leaves its period unstated. `R0374` clears bit 14 over `[this+0x84]·[this+0x88]` words of `[this+0x80]+0xc` — **the whole map, not the visible window** — then re-runs `vt+0x48` on every element of `this+0x9b8` (`L02577`). Its **one** caller (`callto:R0374`, 1 hit, 0 orphan) is `R0334`, which increments `CMapView+0xa70` at `L02610` and gates: `L02576` (a comparison of the dword at `L01661` with zero) · `L10587` (the non-zero branch) · `L10588` (a mask with `0x1f`) · `L10589` (the non-zero branch) · `L01627` (the call to `R0374`). `+0xa70` is `ANIM-CLOCK-001`'s counter, incremented once per `0x401` play-loop tick, so **neither `server+0x00` nor `server+0x04` appears anywhere on this path**: the fog is presentation state on the presentation clock, not hashed and not paced by the simulation. Period **32 ticks ≈ 2 s** at the shipped default speed index, ≈ 4 s at the slowest (`SESS-CLOCK-005`; `ANIM-IDLE-010`'s 250 ms per four `Bee` frames fixes the tick at 62.5 ms). Because `TERR-TILE-079`'s `vt+0x44` re-stamps a *moving* drawable within the period, newly-seen ground brightens promptly while ground a unit has **left** stays lit until the next clear — the visible asymmetry is a consequence of the clear being the only periodic half. `[L01661]` is written only by `R0509` from a message parameter with value space `{0,1,2}` (`L10590` → 0, `L10591` → 1), and **a nonzero value skips the clear entirely**, so the live layer freezes and the map only ever grows

**Confidence.** High (the gate, the period, the map-wide extent and the single caller are cited instructions plus a `callto:` enumeration with 0 orphan hits). **Unknown**: what `[L01661]`'s three values mean — the arm at `L10592` and its two other readers `R0877` `L10593` / `R0916` `L10594`, where nonzero forces a local to **7**, would settle it

**Original status.** ● active

**Evidence.** [EXP-0088](../experiments/EXP-0088-fog-of-war/)

### TERR-FOG-086

**The fog gates two things — the shroud pixels and whether a drawable is drawn at all — and reaches nothing in the simulation.** Beside the renderer, the consumer that matters is `R0378`: it **ORs** the four corner tile words (`L10595`…`L10596`), a mask with `0xc000`, a comparison with `0xc000`, a non-zero flag result, and stores the result into the drawable's `+0x78` (`L10597`) — the field `TERR-SPR-048` names as the sprite pass's guard — then latches `+0x10c` on the transition via `+0x7c`. So **a unit on a cell with no currently-visible corner is not drawn**, and the test is an OR over the corners, i.e. one visible corner suffices. `imm:c000` returns 85 hits / 21 owners and **every game-module owner lies in `R1657 … L10598`** — the client/view modules; the others are CRT `0xc0000000` status constants **Partially retracted by EXP-0349 (ALM-TILEVIEW-122, TERR-DRAWGATE-171): the no-individually-visible-corner implication is false over the admitted raw-word domain. The original OR-before-compare expression stands. One 0xc000 corner suffices but is not necessary: separate 0x8000 and 0x4000 corners also aggregate to0xc000. The stated immediate-reference census and its bounded simulation-negative limits are retained.**

**Confidence.** High for the drawable gate (one function read end to end, five cited instructions). **Medium** for the negative "no simulation routine reads the fog": `imm:c000` excludes displacements and cannot see a byte-wide test on the word's high half — the exact shape that hid this pair's own writer for three experiments. The sweep-immune half of the argument is that the simulation's visibility input is a **different array**, `fog+0x2a008`, whose 9 hits / 5 owners under `disp:2a008` all lie in `R0114…L10599` (`TERR-FOG-088`), and that passability reads the block planes, bounded by `MOVE-DOM-027` to `R0144…R0227` High applies to the exact aggregate operations; the individual-corner equivalence is withdrawn.

**Original status.** ● partially retracted

**Evidence.** [EXP-0088](../experiments/EXP-0088-fog-of-war/), [EXP-0349](../experiments/EXP-0349-tile-word-consumers/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### TERR-FOG-087

**No shipped map authors either fog bit** — that half stands. ~~**and a save does not carry the fog: loading restores a fully unexplored map**~~ — **REFUTED by [EXP-0150]** (`TERR-FOG-145`, `SAV-FOG-061`). Corpus, `tools/fogprobe -mode bits`, both preserved roots, **unaffected**: **EN 38 maps / 880 704 cells, RU 34 maps / 677 696 cells; bit 14 set on 0 cells in either; bit 15 set on exactly 1 cell per map and that cell is index 3 on every one of the 72** — bytes 6..7 of the type-1 payload, i.e. the high half of the section head's `f32`, the same `[u32 typeId][f32]` overlay artefact `TERR-LIGHT-016` was withdrawn for. **Outside cells 0..3: 0 and 0.** So both bits are pure runtime state, and the `.alm` supplies a plane with bit 15 clear everywhere. The persistence conclusion rested on three arguments and two of them are wrong. (1) "No writer of bit 15 exists outside the stamp (`TERR-TILE-079`)" — there is one, `L07966` (a 16-bit OR into the tile word) in the save-load path `R0099`; and `TERR-TILE-079`'s instrument was an `imm:` sweep for bit-15 **clears**, which cannot see a write with a register source. (2) The arithmetic rival — "even a 1-bit layer is 8 192 B against 25 596–27 935 B of decoded stream, attributed term by term" — measured the **compressed body** and the record is in the **uncompressed tail**, where it is a run-length encoding costing 4–1 092 bytes, not a plane. (3) The `.alm` supplying the tile plane on both arms of `R0512` (`ALM-RDR-059`) is correct and is what makes the load arm's OR-only write reproduce the saved state. **A consumer must persist explored terrain across a save/load, and must not persist bit 14**

**Confidence.** ~~High~~ for persistence — the confidence cell claimed the argument "pairs a code enumeration with a corpus instrument immune to that enumeration's blind spot", and both instruments shared the blind spot: neither could see a write with no immediate, and the corpus instrument measured the wrong region of the file. **High retained** for the census: it is exhaustive over both corpora, separates the header overlay from authored cells, and is untouched. See `claims/retracted.md`

**Original status.** ● active (persistence clause refuted)

**Evidence.** [EXP-0088](../experiments/EXP-0088-fog-of-war/), [EXP-0150](../experiments/EXP-0150-save-fog-record/)

**Amended.** The table ledger carried the status "● active (persistence clause refuted)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-FOG-088

**The engine implements the same line-of-sight algorithm TWICE, on two objects, and only one of them is the player's fog — so the AI's vision and the fog are two mechanisms that a consumer must not merge.** `TERR-FOG-080` reads the **client** copy: `R0284`/`R0285` over `mapView+0x17cc`, seeded from the drawable's sight, `k = mapView+0x3f38 = 7` written by the `CMapView` constructor. The **server** copy is a second, structurally parallel implementation in an object embedded at **`world+0x58ee8`** — the base pinned by `L00277` (an add of `0x58ee8`) on `[session+0xa50]` in `R0110`, which is why `AI-SIGHT-006`'s byte map at `+0x2a008` **is** `world+0x82ef0`. Layout, read off `R0134`/`R0136`: `+0x22000` a 64×64 `u16` parent-step table (two `i8`, the server's `pred`); `+0x24000` a 64×64 `i32` accumulator, whose centre cell `[20·64+20]` is the address `AI-SIGHT-006` calls the field `fog+0x25450`; `+0x28000` a 64×64 `i16` step-cost table (the server's `cost`); `+0x2a000`/`+0x2a004` = `1 << k` and `k`; `+0x2a008` the `0x10000`-byte visibility byte map; `+0x3a008` the back-pointer to the world. Same recurrence — parent's budget minus step cost minus `h(cell) − h(observer)`, visible iff `> 0`, rings outward, stop when a whole ring is blocked. **The two differ in five ways that matter and in one that is a customisation trap.** Storage: a `0x10000`-byte byte map vs two bits of the map's own `u16` plane. Clock: per acquisition and per group on the AI's `% 16 == 6` full-tick slot (`AI-TICK-008`) vs the presentation tick's `% 32` (`TERR-FOG-085`). Consumer: a target population vs shroud pixels and the drawable hidden flag. Lifetime: the byte map is cleared before every use by `R0145` — whose **two** callers are `R0114` `L00275` (clear, then stamp **one actor**) and `R0110` `L00276` (clear, then stamp **every member of a group**), so the same array is per-actor or group-shared depending on the caller — against bit 15, which is cleared by nothing. And the trap: **`k` has two sources.** The view's is the compiled constant at `L10553`/`L10554`; the server's is `[Scanning] ScanShift` from `World\Data\map.reg` (`R0274` `L00899` (pushing `0x7`)), which ships as **7** in both roots. They agree today only because the shipped value equals the code default — **editing `ScanShift` moves the AI's sight precision and leaves the player's fog untouched**

**Confidence.** High (the base and the two `k` sources are cited instructions; the layout is read off two functions end to end; each of the five differences is carried by a cited instruction on both sides). ~~**Medium** that the two recurrences are the *same* algorithm rather than two similar ones — the server side was read end to end here and the client side is `TERR-FOG-080`'s reading, not re-derived, so the comparison is between one listing and one published row~~ — *discharged to **High** by [EXP-0120], which re-derived the client side from its own listings and executed both: nineteen clauses compared instruction by instruction, the four tables identical in 0 of 1681 differing cells, the seeds algebraically equal for a whole-cell sight, and 43 640 interior corpus observers with 0 disagreements (`AI-SIGHT-093`). The differences this row lists are unchanged and **two more** are now named: the seed's width, which truncates a hero's sight on the server only (`AI-SIGHT-094`), and the in-bounds inset (`TERR-FOG-118`)*

**Original status.** ● active (amended)

**Evidence.** [EXP-0088](../experiments/EXP-0088-fog-of-war/), **[EXP-0120](../experiments/EXP-0120-sight-writers/)**

**Amended.** The table ledger carried the status "● active (amended)".

### TERR-FOG-089

**The reveal-permission array has THREE writers, not two, and the one `TERR-TILE-079` could not see is a wholesale `memcpy` — so the two message opcodes are now named.** The array is `[[mapView+0x9b4]+0x38]`, one `u16` per player, whose bit 3 is `TERR-TILE-079`'s stamp guard 1. The dispatcher `R0509` decodes its opcode as `L02521` (a subtraction of `0x3`) · `L02522` (a comparison with `0xbb`) · `L10600` (a jump through the table at `L02524` indexed by `4 × case`), so an arm's opcode is `3 + (slot − L02524)/4`. **Opcode 33** (`0x21`; arm entry `L07808`, slot `L07809`) walks the participant list, matches a participant by `Player+0x14` against `[mapView+0x9a4]`, and writes **one** entry — `L07804` (a 16-bit store at `2 × index`) — which is the arm `TERR-TILE-079` located and left unread. **Opcode 45** (`0x2d`; arm entry `L07805`, slot `L07806`) is the writer that row's enumeration missed: after a per-participant pass that starts at index **16** (`L10601` (a store of `0x10` to the frame local at `0xfffff6b8`)) and is bounded by the message's own count at `+0xa`, it resizes the array (`L10602` (the call to `R1622`) on `[mapView+0x9b4]+0x34`, argument `[msg+0xa]`) and then **copies the whole array out of the message body**: `L10603`…`L10604` load `[[mapView+0x9b4]+0x38]`, and `L07807` (the call to `R0544`) is `memcpy(dst, msg+0xe, [msg+0xa]·2)`. So the per-player flags — bit 3 included — are **broadcast wholesale**, not accumulated. Why the earlier sweep could not see it is the failure mode `docs/INSTRUMENT.md` rule 4 names in one line: **a wholesale copy of a structure carries no displacement on any instruction**, so no `disp:38` or store-form sweep can witness it, and the arm's own destination is three `MOV`s away from the `CALL`

**Confidence.** High (both opcodes are computed from the dispatcher's own `SUB`/`JMP` and each arm's `.rdata` slot, each slot found by `refto:` on the arm entry with 0 orphan hits; the `memcpy`'s three arguments are named instructions). **Medium**, unchanged from `TERR-TILE-079`, that bit 3 *means* shared vision — naming the two opcodes says who writes the word, not what the bit is called; neither arm's sender was read, and no third instrument was applied

**Original status.** ● active

**Evidence.** [EXP-0088](../experiments/EXP-0088-fog-of-war/)

### TERR-STRUCT-090

**The footprint-override arm is selected by the SUM of the resolver's two byte arguments, not by the extension kind — and it has exactly ONE reachable caller in the image.** `R0486` does `L10605` (a load of argument `+0x8`) · a mask with `0xff` · `L10606` (a load of argument `+0xc`) · a mask with `0xff` · `L07853` (an addition) · `L10607` (a zero test of the sum) · `L07854` (a greater-than flag) · `L10608` (a comparison of the frame local at `-0x4` with zero) · a jump to `L10476` when zero. The kind byte at `obj+0x40` is not read by that test at all; it was read forty instructions earlier, for the table lookup. So `TERR-STRUCT-077`'s framing — "for that kind `R0486` takes the caller-supplied `(w,h)`" — is right only because of what reaches those two arguments: **an extension record whose two low bytes are both zero would take the table arm**, and a class-33 placement would then get the shipped `1×1` row instead of an author-sized rectangle. Corpus: all 8 shipped extension records carry non-zero bytes in both (3×4, 4×5, 5×4, 7×3, 5×7, 3×3, 3×4, 3×6), so nothing measured under the kind-gated reading changes. **The arm's caller set is one, not three.** `R0486`'s two call sites are reached by three callers (`ALM-CLS-063`), and of those, `R0490` pushes `L02196` (pushing `0x0`) / `L02197` (pushing `0x0`) and `R0693` pushes `L10609` (pushing `0x0`) / `L10610` (pushing `0x0`); the non-extension `.alm` constructor `R0481` zeroes `rec+0x0a`/`+0x0c` outright (`L10611`, `L10612`). **The only route to the override is `R0489`'s Building arm carrying a `kind == 0x21` record's two extension bytes**

**Confidence.** High (the seven instructions of the test, and the four literal-zero pushes plus the two zeroing stores that close every other route; the rival — a `CMP` against `0x21` selecting the arm — has no instruction anywhere in the routine, which is what a reading transcribed from the *discriminator* rather than from the *branch* would have needed)

**Original status.** ● active

**Evidence.** [EXP-0091](../experiments/EXP-0091-placement-axes/)

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-STRUCT-100 | (rom.exe) | High / Medium | ● active | [EXP-0092](../experiments/EXP-0092-structure-art/) |
| TERR-STRUCT-101 | (rom.exe) | High | ● active | [EXP-0092](../experiments/EXP-0092-structure-art/) |
| TERR-STRUCT-102 | (rom.exe) | High | ● active | [EXP-0092](../experiments/EXP-0092-structure-art/) |
| TERR-STRUCT-103 | (rom.exe) | High | ● active | [EXP-0092](../experiments/EXP-0092-structure-art/) |
| TERR-STRUCT-104 | (rom.exe) Flat selects the structure body phase; both body phases admit signed +0x78<2. | High | ✔ promoted (partially retracted) | [EXP-0092](../experiments/EXP-0092-structure-art/); [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |
| TERR-STRUCT-105 | (rom.exe) | High | ● active | [EXP-0092](../experiments/EXP-0092-structure-art/) |
| TERR-STRUCT-106 | Drawable lift uses footprint-centre signed interpolation; the unconditional structure cell-centre and four-corner-mean consequence is partially retracted by TERR-221. | High / Medium | ● active (partially retracted) | [EXP-0092](../experiments/EXP-0092-structure-art/) |
| TERR-STRUCT-107 | (rom.exe) | High | ● active | [EXP-0092](../experiments/EXP-0092-structure-art/) |
| TERR-LIGHT-108 | the global at `L06260` is the `ShowTimeFlow` game option — a shipped, persisted, on-by-default user setting, not a debug switch. | High / Medium | ● active | [EXP-0116](../experiments/EXP-0116-daylight-cycle/) |
| TERR-LIGHT-109 | The sun angle's span is ±0.78539815, not ±π/2 — `TERR-LIGHT-030` and `TERR-LIGHT-014` are both wrong about it by a factor of two (see retracted.md). | High | ● active | [EXP-0116](../experiments/EXP-0116-daylight-cycle/) |
| TERR-LIGHT-110 | Five stores write the sun angle, and TWO of them compute it — the night band sweeps as well as the day band, which no prior row records. | High / Unknown | ● active | [EXP-0116](../experiments/EXP-0116-daylight-cycle/) |
| TERR-LIGHT-111 | The sun angle reaches the rest of the engine through ONE cache, on the terrain object, written by its only reader. | High | ● active | [EXP-0116](../experiments/EXP-0116-daylight-cycle/) |
| TERR-LIGHT-112 | `R1820` is the shear function: a dead band around zero, then a two-thirds scale. | High / Unknown | ● active | [EXP-0116](../experiments/EXP-0116-daylight-cycle/) |
| TERR-LIGHT-113 | Every shadow pass in the image derives its shear from the sun angle — and the unit body computes it and throws it away. | High | ● active | [EXP-0116](../experiments/EXP-0116-daylight-cycle/) |
| TERR-LIGHT-114 | The relight cadence: one caller, once per 20 in-game minutes, or forced — and `TERR-LIGHT-022`'s condition is confirmed with its clock now named. | High | ● active | [EXP-0116](../experiments/EXP-0116-daylight-cycle/) |
| TERR-SIGHT-115 | The sight object at `world+0x58ee8` holds FOUR grids, not three, and its init builds three tables of which only two serve line of sight. | High | ● active | [EXP-0117](../experiments/EXP-0117-sight-region/) |
| TERR-SIGHT-116 | The playable rectangle is `world+0x58ee0..0x58ee3` = `(8, 8, W-9, H-9)`, written at map load, and the sight walk tests it as BYTES while indexing the map in 32-bit — which disagree at the edge of a 256-wide map. | High / Medium | ● active | [EXP-0117](../experiments/EXP-0117-sight-region/) |
| TERR-FOG-117 | `TERR-FOG-080`'s flat-ground region is the region of a table it never finished building: `R0283`'s last four instructions were not transcribed, and the correct figure at sight 6 is 145 cells, not 127. | High | ● active | [EXP-0120](../experiments/EXP-0120-sight-writers/) |
| TERR-FOG-118 | The playable rectangle the two implementations walk is not the same rectangle: the AI's is one cell narrower on every side, and its bounds test is byte-wide. | High / Medium | ● active | [EXP-0120](../experiments/EXP-0120-sight-writers/) |
| TERR-LIGHT-119 | The complete sky-light schedule: what every arm of `R1813` writes to the tint bytes and the two intensity fields, and how. | High | ● active | [EXP-0123](../experiments/EXP-0123-sky-light-schedule/) |
| TERR-LIGHT-120 | The routine reads five addresses and none of them is a table - the sky-light values are never looked up. | High | ● active | [EXP-0123](../experiments/EXP-0123-sky-light-schedule/) |
| TERR-LIGHT-121 | Six arms, not five: each twilight band splits in two at its own `CMP`, and two branches in the region are unreachable. | High | ● active | [EXP-0123](../experiments/EXP-0123-sky-light-schedule/) |
| TERR-LIGHT-122 | One multiplier, two post-shifts: the ramps divide by 120 and the hour selector by 60, and the ramp's clock is `t mod 120`, not the day arm's `(t+360) mod 720`. | High | ● active | [EXP-0123](../experiments/EXP-0123-sky-light-schedule/) |
| TERR-LIGHT-123 | The six arms are one continuous 24-hour programme, and that is a property the routine never asserts. | High | ● active | [EXP-0123](../experiments/EXP-0123-sky-light-schedule/) |
| TERR-LIGHT-124 | `L10231`, `L10613`, `L10614` are R, G, B in ascending address order - the rival BGR reading is excluded twice over. | High | ● active | [EXP-0123](../experiments/EXP-0123-sky-light-schedule/) |
| TERR-LIGHT-125 | What the schedule buys: flat ground runs level `[45,64]` over an in-game day and every sprite on the map dims through six steps together. | Medium | ● active | [EXP-0123](../experiments/EXP-0123-sky-light-schedule/) |
| TERR-LIGHT-126 | The two shroud levels are on the day/night schedule too, and one is always twice the other. | High | ● active (amended) | [EXP-0123](../experiments/EXP-0123-sky-light-schedule/) |
| TERR-LIGHT-127 | The sky tint can only ever add: it is taken unsigned and applied before the level multiply, so nothing in the day/night cycle darkens a colour channel. | High | ● active | [EXP-0123](../experiments/EXP-0123-sky-light-schedule/) |
| TERR-LIGHT-128 | The customisation limit (G2): the whole day/night schedule is 70 immediate fields, 181 bytes, inside one routine in `.text` - there is no data file, registry key or table to change. | High | ● active | [EXP-0123](../experiments/EXP-0123-sky-light-schedule/) |
| TERR-SHDW-129 | The silhouette blitter shears about the BOTTOM edge of its own rectangle, and the sign says which way a shadow leans. | High | ● active | [EXP-0129](../experiments/EXP-0129-shadow-destination/) |
| TERR-SHDW-130 | The three shadow casters' destinations, term by term. | High | ● active | [EXP-0129](../experiments/EXP-0129-shadow-destination/) |
| TERR-SHDW-131 | The corpus's shear "addition vs subtraction" is not a contradiction: it is one rule written in two coordinate directions, and the caster's scalar moves the blitter's pivot to the row the caster wants. | High / Unknown | ● active | [EXP-0129](../experiments/EXP-0129-shadow-destination/) |
| TERR-SHDW-132 | The unit's flat blit arm is the sheared arm's own rule collapsed to a single 32-pixel step, not a different placement. | High | ● active | [EXP-0129](../experiments/EXP-0129-shadow-destination/) |
| TERR-SHDW-133 | A shadow is placed with ONE height for the whole silhouette, sampled once, at the caster's own cell - established positively. | High | ● active | [EXP-0129](../experiments/EXP-0129-shadow-destination/) |
| TERR-SHDW-134 | One slope per invocation; the terrain height enters neither the slope nor anything it multiplies. | High | ● active | [EXP-0129](../experiments/EXP-0129-shadow-destination/) |
| TERR-SHDW-135 | `vt+0x3c` is a family of four cores behind two globals, and it never reads a source pixel - the "silhouette" is literal. | High / Medium | ● active | [EXP-0129](../experiments/EXP-0129-shadow-destination/) |
| TERR-SHDW-136 | A structure strip carries BOTH shear terms at once, and on 15 of 66 shipped classes `ShadowY` is not a pixel height at all - it is a suppression sentinel enforced by displacement, with no branch anywhere. | High / Medium | ● active | [EXP-0129](../experiments/EXP-0129-shadow-destination/) |
| TERR-SPR-137 | (rom.exe) The map painter has nine cell sweeps and one non-cell collection walk in its ten-phase composition. | High | ✔ promoted (partially retracted) | [EXP-0130](../experiments/EXP-0130-shadow-bounds/); [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |
| TERR-SPR-138 | (rom.exe) Drawable cell sweeps walk rows ascending and columns descending, with a narrower marker/bar window and a column-major shroud. | High | ✔ promoted (partially retracted) | [EXP-0130](../experiments/EXP-0130-shadow-bounds/); [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |
| TERR-SPR-139 | (rom.exe) The +0x90 shadow and body dispatches are separate complete cell sweeps, conditionally populated through CAirUnit selector3. | High | ✔ promoted (partially retracted) | [EXP-0130](../experiments/EXP-0130-shadow-bounds/); [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |
| TERR-SPR-140 | The only bound on a drawn silhouette besides its own frame is ONE global device-space clip rectangle, and the blitter enforces it per row and per pixel — there is no per-cell destination window anywhere on the draw path. | High | ● active | [EXP-0130](../experiments/EXP-0130-shadow-bounds/) |
| TERR-SPR-141 | The shroud cull `R1401` has exactly ONE caller in the image, and it culls an object's shadow and body TOGETHER — one boolean per object, nothing per row. | High | ● active | [EXP-0130](../experiments/EXP-0130-shadow-bounds/) |
| TERR-FOG-142 | The shroud sweep is the LAST cell loop in the painter, after every sprite — so a shadow reaching a fogged cell is drawn in full and then covered, neither clipped to the fog nor skipped. | High / Medium | ● active | [EXP-0130](../experiments/EXP-0130-shadow-bounds/) |
| TERR-LIGHT-143 | Nothing stops the shadow map being applied twice, so two overlapping silhouettes COMPOUND: the intersection is `((16-L)/16)²` of the ground, not `(16-L)/16`. | High / Medium | ● active | [EXP-0130](../experiments/EXP-0130-shadow-bounds/) |
| TERR-SPR-144 | `[L06416]` is the engine's shadow switch, and it reaches nothing else: seven reads image-wide, all seven inside the map-view painter, each one guarding a shadow dispatch. | High / Medium | ● active | [EXP-0130](../experiments/EXP-0130-shadow-bounds/) |
| TERR-FOG-145 | Tile bit 15 survives a save and a load, and its second writer is the save-load path. | High / Medium | ● active | [EXP-0150](../experiments/EXP-0150-save-fog-record/) |
| TERR-CELLREC-146 | What keys the cell record's four occupant slots, and what each of the six area-layer slots holds. | High / Medium | ● active | [EXP-0174](../experiments/EXP-0174-area-movement/) |
| TERR-FOOTPRINT-147 | Successful entry stores an actor's pointer in every cell record of its `n × n` footprint; refusal stops iteration without local rollback. | High | ● active (amended, partially retracted) | [EXP-0174](../experiments/EXP-0174-area-movement/) |
| TERR-PASS-148 | Bit 4's origin, located: the area module writes it, one serializer reads it, and no mover mask contains it. | High / Medium | ● active | [EXP-0174](../experiments/EXP-0174-area-movement/) |
| TERR-LIGHT-149 | The `TERR-LIGHT-015`/`TERR-LIGHT-023` Unknown, answered field by field: over the searched population no reachable read observes the stored value of any of the four light-side scalars — payload `+0x08`, `+0x0c`, `+0x10`, `+0x14` — while ... | Medium | ● active | [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |
| TERR-LIGHT-150 | `P+0x2c`, the landscape slot payload `+0x0c` is read into, has no access of any kind in the searched population: the builder's own store is the only site that names it. | Medium | ● active | [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |
| TERR-LIGHT-151 | Why "overwritten before read" is the map-load answer and not just one routine's internal order: the load path drives the relight past every branch but one, and the values it copies in cannot come from a file. | High / Medium | ● active | [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |
| TERR-LOAD-152 | `R0516`'s argument is the map's own tile-group mask, and it is the seventh scalar of the landscape builder's case-0 read sequence. | High / Medium | ● active | [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |

### TERR-STRUCT-100

**(rom.exe)** A placed structure has **no anchor pixel and no canvas** — the two values every unit and static object is placed by (`TERR-SPR-040`, `REG-OBJ-046`). Its anchor is the **footprint's min-col / min-row cell**: the client's create arm writes `this+0x08 = msg[0x0d]*256 + 128` and `this+0x0c = msg[0x0e]*256 + 128` (`L10615`/`L10616`) — the exact centre of one cell — and `R0614` derives `this+0x34 = this+0x08 >> 8`, `this+0x38 = this+0x0c >> 8`. The body pass places image cell `(k, c)` at **`dstX = screenCol*32`**, **`dstY = screenRow*32 - this+0x10 - this+0x68`**, where `screenCol`/`screenRow` are the walked cell (`R1794`, `L10617`…`L10618`, `L10619`) and `R0794` passes that pair to the blitter as the frame's **top-left**. `this+0x10` is written `0` by the create arm (`L10620`) for every structure, bridge or not — all three allocation arms converge on that fill — so in practice `dstY = screenRow*32 - lift`. `this+0x68` is the altitude lift shared with units (`TERR-SPR-038`). Corpus: **3141 of 3141** shipped type-4 records over 38 maps carry `0x80` in the low byte of both fine coordinates, and 3141/3141 footprints lie inside the map rectangle

**Confidence.** High for the formula (every term is a named instruction and the destination's meaning is fixed by the blitter's own argument use) and for the cell anchor (one create instruction plus 3141/3141; the live rivals — an anchor pixel like `CenterX`/`CenterY`, or `ShadowY`/`SelectionY1` as an origin — are refuted by the *absence* of any such key read on the body path, a routine read end to end). Medium that `this+0x10` is 0 **at draw time**: the create site is decisive, but a store-form sweep at displacement `0x10` returns 1997 hits over 1024 owners and cannot enumerate later writers — the instrument's blind spot, stated so it is not mistaken for an enumeration

**Original status.** ● active

**Evidence.** [EXP-0092](../experiments/EXP-0092-structure-art/)

### TERR-STRUCT-101

**(rom.exe)** A structure is drawn **once per footprint cell, and each call draws a vertical strip of image rows** — not one image, and not one image per tile. `vt+0x38` (`R1838` -> `R0613` arm 1) registers the drawable into `CMapView+0x94` over its **whole `TileWidth x TileHeight` rectangle**, so the renderer's cell walk reaches it once per occupied cell with `(screenCol, screenRow, lightByte)`. Inside, `R1794` computes the local column `COL0 = view+0x5c - this+0x34 + arg1` and local row `ROW0 = view+0x60 - this+0x38 + arg2`, then loops `k` from `rowTop = ROW0 - TileHeight + FullHeight` **down to** a limit that is `rowTop` itself when `ROW0 != 0` and **`0`** when `ROW0 == 0` (`L10621`…`L10622`), blitting frame `k*TileWidth + COL0` and stepping `dstY` up by 32 each time. So the image grid's bottom `TileHeight` rows map one-to-one onto the footprint rows, and the top `FullHeight - TileHeight` rows — the **overhang** — are drawn together with the structure's own back row, above it. Corpus: `FullHeight >= TileHeight` on 66/66 (differences `0 x38, 1 x20, 2 x6, 3 x2`), and **1284 of 3141** shipped placements have an overhang; when `FullHeight == TileHeight` the loop is exactly one blit per cell

**Confidence.** High (the loop bounds, the stride and the `ROW0 == 0` special case are named instructions in a routine read end to end; the rival "one image sized `TileWidth*32 x FullHeight*32`" dies on `SPR256-STR-040`'s 130/130 uniform 32x32 frames, and the rival "one frame per *footprint* cell" dies on the `limit = 0` arm and on `FullHeight > TileHeight` for 28 classes)

**Original status.** ● active

**Evidence.** [EXP-0092](../experiments/EXP-0092-structure-art/)

### TERR-STRUCT-102

**(rom.exe)** The frame arms and their order (`R1794`, per image cell): **(1)** if `(i16)this+0xfc <= 0` — the current health, filled from the create message's `+0x12` word — **and** the class's `Indestructible` is 0, the frame is `spriteFrameCount - TileWidth*FullHeight + cellIndex`, the **ruin** grid (`L10097`…`L10623`); **(2)** otherwise, if the class's animation timeline is non-empty *and* `[L05659]` (the animations-enabled global, `TERR-ANIM-009`) is set, `phaseFrame = timeline[this+0x70]` and, when that is non-zero **and** `AnimMask[cellIndex] != '-'`, the frame is `TileWidth*FullHeight + (phaseFrame-1)*live + rank` (`L10624`…`L10098`); **(3)** otherwise the base frame `cellIndex`. `this+0x70` is advanced by `vt+0x3c` (`R1839`) as `phase = (phase+1) mod class+0x3c`, **gated `this+0x78 == 0`** — so a structure's animation only runs while its cell is in the local player's sight (`TERR-TILE-079`), unlike a unit's, which is driven by the tick (`ANIM-CLOCK-001`). Max health comes from `Data.bin`'s Buildings row (`[L06410] + idx*0x1c`, column 3, defaulting to 1000 when 0), not from the registry

**Confidence.** High (both gates and all three frame expressions are named instructions; the `Indestructible` conjunct is a second test-and-branch in the same arm, and the animation gate's dependence on the fog state is one comparison of `[this+0x78]` with 0 in `R1839`)

**Original status.** ● active

**Evidence.** [EXP-0092](../experiments/EXP-0092-structure-art/)

### TERR-STRUCT-103

**(rom.exe)** The structure's **shadow** is `vt+0x2c` (`R1840`) and it is the *same* loop over the same frames with the same `dstY`; the only differences are the blit entry point — `CSprite256 vt+0x3c` with 6 arguments and the destination-recolour global `[L04368]` (`L10625`), the shadow family `TERR-LIGHT-059` separates — and the destination **X**, which is sheared per image row: `dstX = screenCol*32 + ftol( tan(sunAngle) * ((FullHeight - k)*32 - ShadowY) )`, the two doubles being `[L10626] = 32.0` (the row height) and, for the blit's own 16.16 per-row slope argument, `[L10627] = 65536.0` (`TERR-SPR-066`). So **`ShadowY` is the pixel height at which the shear is zero** — where the structure meets the ground — and it is the only use of that key anywhere on the draw path. The `b`-overlay repeats the shadow through `[L04335]`. The two `VariableSize` bridge subclasses override `vt+0x2c` with a bare `0xc`-byte-releasing return (`R1841`): a wooden bridge casts no shadow

**Confidence.** High (the FPU sequence is read instruction by instruction, both constants are dumped from the PE through its own section table, and the zero-shear reading is forced by the subtraction's position rather than fitted)

**Original status.** ● active

**Evidence.** [EXP-0092](../experiments/EXP-0092-structure-art/)

### TERR-STRUCT-104

**(rom.exe) Flat selects the structure body phase; both body phases admit signed +0x78<2.** Structures register their prepared footprint in view+0x94 (TERR-STRUCT-101, ANIM-REGISTER-083). The early body callL10628 takes Flat!=0, while main-cell callL02987 takes Flat==0. The latter precedes selector2 CUnit shadow/body calls in that cell. The later +0x78==0 testL10629 guards rectangle/mark work after the body, not the body itself; the former main-body sight-only clause is retracted (ANIM-DRAWGATE-087). Structure shadows form an earlier separate phase. Row/column ordering applies within each phase, not across every drawable category: selector4/CBackPack bodies precede the main cell phase and selector3 bodies follow it (ANIM-CELL-085, ANIM-AIRPASS-086). The recorded seven Flat classes and structure image-strip/corpus facts are unchanged.

**Confidence.** High for the opposite Flat tests, the actual body-gate anchors and static conditional category ordering. Native pixel overlap and registration collision order remain Unknown.

**Original status.** ✔ promoted (partially retracted)

**Evidence.** [EXP-0092](../experiments/EXP-0092-structure-art/); [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/)

**Amended.** The table ledger carried the status "✔ promoted (partially retracted)". `retracted.md` records a correction against this claim.

### TERR-STRUCT-105

**(rom.exe)** The two wooden bridges are **separate C++ classes with their own frame selector**, not registry data — which is what `VariableSize` actually means. The client's create arm branches on the class id byte: `0x21` (ID 33, `Vertical Wooden Bridge`) and `0x25` (ID 37, `Horisontal Wooden Bridge`) each allocate **`0x140`** bytes (against `0x138` for a plain `CStructure`) and call `R1842(w, h)` with two further message bytes, which lands them in `this+0x138` / `this+0x13c`; the vtables installed are `L10630` (`CVerticalWoodenBridge`) and `L10631` (`CHorisontalWoodenBridge`), the names being the image's own RTTI strings at `L10632` / `L10633`. Their `vt+0x28` (`R1843`, `R1844`) discards the grid entirely and picks a **nine-patch** frame from the cell's position in that `w x h` rectangle — corner / edge / interior — with one blit and no loop: 9 frames for the vertical bridge, 14 for the horizontal, whose bottom edge additionally cycles frames 8..11 on `[view+0xa70] & 3`, the animation counter `ANIM-CLOCK-001` names. Their sheets carry exactly **9** and **14** frames. Corpus: the `.alm` type-4 size extension exists only for `kind == 0x21` (`ALM-OBJ-019`), so of 3141 shipped records **8** carry it, all of class 33, and **class 37 is placed 0 times on any shipped map** — the 14-frame selector is unreachable from a shipped `.alm`

**Confidence.** High (two allocation sizes, two ctor call sites, two vtable installs and two selectors read whole; the 9/14 frame counts of the two sheets match the two selectors' index ranges exactly, which is the discriminator — a data-driven rival predicts the grid identity `SPR256-STR-041` measures, and these are the only two classes it misses)

**Original status.** ● active

**Evidence.** [EXP-0092](../experiments/EXP-0092-structure-art/)

### TERR-STRUCT-106

The affected former wording below is partially retracted. See the
**Amended.** paragraph for the retained facts and correcting claim.

**(rom.exe)** `drawable+0x68`, the altitude lift every sprite is raised by, is computed in `R0614` (`L10634`…`L10635`) as a **bilinear interpolation of the containing cell's four corner heights at the object's sub-cell position**, the position used being the **footprint centre** `+0x58`/`+0x5c` = `(TileWidth<<7) + fineX - 128`, `(TileHeight<<7) + fineY - 128`, and the grid being `[[view+0x80]+0x10]` read as signed bytes. This refines `TERR-SPR-039`/`TERR-SPR-040`'s "the mean of the cell's four corners": at the exact centre of a cell both fractional weights are 16/32 and bilinear *is* the mean, which is why the mean reading held — but a drawable at any other sub-cell position is lifted by the interpolant, not the mean. For a structure the two always agree, because `TERR-STRUCT-100` puts every one of the 3141 shipped placements exactly on a half-cell

**Confidence.** High for the arithmetic (four sign-extended loads and three interpolations, read instruction by instruction in one routine). Medium for the consequence that a structure always samples a cell centre — that follows from the corpus and the `(TileWidth<<7)` term rather than from a test in the code

**Original status.** ● active

**Evidence.** [EXP-0092](../experiments/EXP-0092-structure-art/)

**Amended.** TERR-221 partially retracts the unconditional cell-centre/mean consequence. The sampled position is the footprint centre; its weights depend on footprint dimensions. Three signed truncations are not an unconditional four-corner mean. The named arithmetic and signed loads stand.

### TERR-STRUCT-107

**(rom.exe)** **A structure's drawing does not depend on its owner.** The body call passes `(dstX, dstY, frame, lightByte, 0)` and the shadow `(dstX, dstY, frame, [L04368], slope, 0)`; neither reads `this+0x14`, the owning participant the create arm stores at `L10636`. Enumerating the whole `CStructure` vtable `L10637` — 21 slots, every one with a function, `FindVtables` over `.rdata` on the repaired table — the **only** slot that reads a per-owner value is `vt+0x34` (`R1307`), which indexes the per-player colour array `[L05816]` by `[this+0x14]+0x08` and fills a rectangle at the object's fine position scaled by the view's minimap factors: the **minimap blip**. There is no team tint, no per-owner palette and no owner-selected shade object on the world sprite — where a unit has an unresolved per-owner arm (`TERR-LIGHT-059`'s `L10417`), a structure has none

**Confidence.** High (an enumeration over one class's complete method surface, with the instrument named and every slot resolving to a function; a vtable enumeration's blind spot — a method reached other than through the table — does not bear on a claim about what *these* slots do)

**Original status.** ● active

**Evidence.** [EXP-0092](../experiments/EXP-0092-structure-art/)

### TERR-LIGHT-108

**the global at `L06260` is the `ShowTimeFlow` game option — a shipped, persisted, on-by-default user setting, not a debug switch.** `EnumRefs refto:L06260` (reference manager; blind to a store through a register-held pointer, which is why every site pushing `L06260` was followed) = **14 hits / 10 owners / 0 orphan**. Four of them push its **address** as argument 5 of an import call, and `StrDump fnstr:` puts the name beside the push: `L10638` with `L10639` = **"ShowTimeFlow"** into `[L10640]` — arg order `(hKey, name, 0, 4, &flag, 4)`, i.e. `RegSetValueExA` with `dwType = REG_DWORD` — and `L10641` with `L10642` into `[L10643]`, arg order `(hKey, name, 0, 0, &flag, &cb)` with `cb` preset to 4 at `L10644`, i.e. `RegQueryValueExA`. The savegame carries it too, under section **"GameOptions"** key **"ShowTimeFlow"** (`L10645`/`L10646` save, `L10647`/`L10648` load). **The default is 1**: `R1845` — the constructor of the option block at `L06413`, reached as `R1846` = a stub loading `L06413` as `this` and jumping to `R1845` — stores the dword `0x1` into `[L06260]` at `L10649` unconditionally, and it **runs before `main`**: `callto:R1846` returns one caller *in orphan code* and `callto:L10650` returns **zero**, so a function-table sweep cannot settle whether it runs; a whole-image **raw dword scan** (a different instrument, immune to function tables) finds `L10650` exactly once, at `L10651` in `.data`, inside a **135-entry contiguous run of in-image code pointers** — the C++ static-initializer table. Runtime toggle: `L10652` (a comparison of the dword at `[L06260]` with zero) / `L10653` (a flag set when equal) / `L10654` (a store of that flag into `[L06260]`), then a push of `1` and the call to `R0532` — a **forced** relight. All three dense switches of `R0819` were resolved from their own index bytes: switch 1 (`key−8`) sends `0x4e` to `L06253`, switch 2 (`key−0x41`, guarded) sends `0x4e` to `L06243`, switch 3 (`key−0x20`) sends `0x4e` to the toggle, so **key `N` reaches it down either branch of switch 2's guard**. The arm's gate `[L00627] != 0` is session lifetime (`R0231` sets 1, `R0229` clears), not a debug flag. **So the computed-angle path is reachable in a shipped campaign mission and is the default**

**Confidence.** High for the identification, the default store and its static-init reachability (every step is a named instruction, the switch resolution is a brute force over the tables rather than a transcription, and the one input a function-table sweep could not settle was settled by an instrument with a different blind spot) / **Medium** that a given install *runs* with it on: the image fixes the default and `RegQueryValueExA` leaves `lpData` untouched when the value is absent, but what a particular profile or savegame contains is not a fact about the image

**Original status.** ● active

**Evidence.** [EXP-0116](../experiments/EXP-0116-daylight-cycle/)

### TERR-LIGHT-109

**The sun angle's span is ±0.78539815, not ±π/2 — `TERR-LIGHT-030` and `TERR-LIGHT-014` are both wrong about it by a factor of two (see retracted.md).** The two doubles are read from the PE through its own section table on **both** roots: `[L07946]` = `4a d8 12 4d fb 21 e9 bf` = **−0.78539815** (not −π/2) and `[L07945]` = `18 39 35 9d 46 df 61 bf` = **−0.0021816615277777777** (negative). `L10655` (a reverse subtraction of `[L07946]`) computes `src − ST0`, so θ = `−0.78539815 + m·0.0021816615`, and `720 × 0.0021816615 = 1.5707963 = π/2` — **the total sweep is π/2, running −π/4 → +π/4**, not π running −π/2 → +π/2. `TERR-LIGHT-014`'s "`+π/2 (_L07947) … −π/2 (_L07946)`" is the same error: `[L07947]` = **+0.78539815**, the same truncated-π quarter with the sign flipped. Consequence for every consumer of `TERR-LIGHT-013`/`-028`: `stepH = 32/cos\|θ\|` never exceeds `32/cos(0.7854) = 45.25` and the lateral `tan\|θ\|` never exceeds `1.0`, where the retracted span allowed both to diverge

**Confidence.** High (both constants dumped as raw bytes from the PE on both roots, and the operand order of the `FSUBR` is the instruction's own)

**Original status.** ● active (corrects `TERR-LIGHT-030`, `TERR-LIGHT-014`)

**Evidence.** [EXP-0116](../experiments/EXP-0116-daylight-cycle/)

### TERR-LIGHT-110

**Five stores write the sun angle, and TWO of them compute it — the night band sweeps as well as the day band, which no prior row records.** `EnumRefs refto:L07948` = **6 hits / 2 owners / 0 orphan**: five writes in `R1813`, one read. The five arms, selected by `hour = (t/60) % 24` where `t` is the routine's argument: **cycle off** `L10656` → literal `+0.78539815`; **day, hour 6…17** `L10657` → computed, `m = (t+0x168) % 0x2d0`, θ = `−0.78539815 + m·0.0021816615`; **dawn, hour 2…5** `L10658`/`L10659` → literal `−0.78539815` (high dword `0xbfe921fb`; both sub-arms converge on `L10660`); **dusk, hour 18…21** `L10661`/`L10662` → literal `+0.78539815` (high dword `0x3fe921fb`; the 18…19 arm reaches it by the jump at `L10663` and the 20…21 arm falls through); **night, hour 22,23,0,1** `L10664` → **computed**, `m = (t+0x78) % 0xf0` (`L10665` forms `t + 0x78`, `L10666` loads the divisor `0xf0`), θ = `+0.78539815 − m·0.0065449845833333332` (`L10667` multiplies by `[L10668]`, `L10669` reverse-subtracts `[L07947]`) — a 240-minute sweep back from `+π/4` to `−π/4`, the same arc traversed three times faster while it is dark. Re-executed over a whole in-game day: **960 of 1440 full ticks carry a computed angle**, 480 a literal one. What this row does **not** establish: the tint and intensity schedule of the dawn and dusk arms, which compute through magic-division sequences this round did not read — `TERR-LIGHT-014` still owns those, at Medium

**Confidence.** High for the arm structure and both computed expressions (one routine read whole, every immediate re-read from the PE on both roots, and the band boundaries are the routine's own `CMP`s) / **Unknown** for the dawn/dusk tint arithmetic, explicitly not re-derived

**Original status.** ● active

**Evidence.** [EXP-0116](../experiments/EXP-0116-daylight-cycle/)

### TERR-LIGHT-111

**The sun angle reaches the rest of the engine through ONE cache, on the terrain object, written by its only reader.** `_L07948` has exactly **one read** image-wide, `L10670` in `R0468`, and the next four instructions copy the double onto the object: `L10670` (a load of the low dword at `[L07948]`) / `L10671` (a load of the high dword at `[L10240]`) / `L10672` (a store of the low dword at `+0x20`) / `L10673` (a store of the high dword at `+0x24`). `R0532` invokes `R0468` as `[view+0x80]`, so the cache lives at `[view+0x80]+0x20`. **The consequence a consumer must not miss: nothing reads the global again.** Every later use of "the sun angle" — the per-vertex relief's `stepH` and lateral shear, and every shadow shear (`TERR-LIGHT-113`) — reads the cache, so what the screen shows is the angle as of the **last relight** (`TERR-LIGHT-114`), never the current tick's

**Confidence.** High (the read enumeration is a reference-manager sweep over an absolute-addressed global with 0 orphan hits, and the copy is four consecutive named instructions)

**Original status.** ● active

**Evidence.** [EXP-0116](../experiments/EXP-0116-daylight-cycle/)

### TERR-LIGHT-112

**`R1820` is the shear function: a dead band around zero, then a two-thirds scale.** Its whole input is `R1820` (a load of the double at `+0x20`) — `TERR-LIGHT-111`'s cache. Read arm by arm with the four constants dumped from the PE on both roots (`[L05652]` = **0**, `[L10674]` = **−0.05**, `[L10675]` = **+0.05**, `[L10676]` = **0.66666666666666663**): `−0.05 < θ < 0` returns **−0.05**; `0 < θ < +0.05` returns **+0.05**; otherwise it returns **θ · 2/3** (`L10677` (a multiplication by `[L10676]`)). The comparisons are a floating-point compare followed by a status-word read and a test of bit `0x1` (C0, i.e. `<`) or of bits `0x41` (C0 or C3, i.e. `≤`), so the band is open at both ends and `θ = 0` takes the scale arm and returns 0. Re-executed over the day band, the dead band is entered on **22 + 22 of 720** full ticks — it does not swallow the sweep. **Why the band exists is inferred from its shape and from nothing in the image**

**Confidence.** High for the arithmetic (one routine read whole, all four constants re-read from the PE on both roots, both `TEST` masks distinguished) / **Unknown** for its purpose

**Original status.** ● active

**Evidence.** [EXP-0116](../experiments/EXP-0116-daylight-cycle/)

### TERR-LIGHT-113

**Every shadow pass in the image derives its shear from the sun angle — and the unit body computes it and throws it away.** `EnumRefs callto:R1820` = **10 hits / 5 owners / 0 orphan**, and all ten load the receiver the same way, loading `[<view> + 0x80]` as `this`, the view being `this` in `R0379` (`L10678`, `L10679`) and `[actor+0xe0]` in the four code-region draw routines. `actor+0xe0` is written from `R0509`'s own `this` (`L10680`, `L10681`, `L10682`) and `R0509` is invoked on `campaign+0xd0` — the same receiver `R0454` and `R0819` pass to `R1678`/`R1679`, which is what makes it the same object `R0532` hands to `R0468`. Each pass then takes the tangent — `FPTAN` at `L10683` (unit shadow `R0553`), `L10684` and `L10685` (structure shadow `R1840`), `L10686`/`L10687` (`R1796`), and the CRT `tan` at `L03742` (object shadow, inside `R0379`) — and multiplies by **65536.0** (`[L10627]`, or `[L10423]` on the object path) into the blitter's 16.16 per-row X slope (`TERR-SPR-066`). So **`TERR-STRUCT-103`'s `tan(sunAngle)` is exactly `tan(R1820(cached θ))`**, and `TERR-SPR-048`'s "sun scalar" is the same quantity. **The unit BODY `R0552` calls the same routine at `L10688` and discards the result one instruction later — `L10689` (a floating-point store-and-pop)** — which is `TERR-SPR-067`'s "the shadow's expression without the shear" seen as a *discarded* computation rather than an absent one. Re-executed over the day band, `tan(shear)` spans **−0.5774 … +0.5754**, crossing the dead band near noon; with the cycle off it is fixed at **+0.57735025728078282** = `tan(0.5235987666…)`, since `0.78539815 · 2/3` is that argument. **A consumer drawing a fixed-lean shadow is wrong on the default settings, and one leaning the body is wrong always**

**Confidence.** High (the call enumeration states its instrument and has 0 orphan hits; every receiver load, `FPTAN`, `FMUL` and the body's `FSTP` is a named instruction; the object identity is closed through `campaign+0xd0`'s shared receiver rather than assumed)

**Original status.** ● active

**Evidence.** [EXP-0116](../experiments/EXP-0116-daylight-cycle/)

### TERR-LIGHT-114

**The relight cadence: one caller, once per 20 in-game minutes, or forced — and `TERR-LIGHT-022`'s condition is confirmed with its clock now named.** `EnumRefs callto:R1813` = **1 hit / 1 owner / 0 orphan**: `L10690`, inside `R0532`, which `TERR-LIGHT-022` already establishes as the relight driver's only caller. The routine is gated on `campaign+0x3dc & 1` (the map screen) at `L10691`…`L10692`, then runs unforced only when `(campaign+0x3e0 & 0xf) == 0` **and** `((campaign+0x3e0 >> 4) + 0x168) % 0x14 == 0`; its argument `[argument +0x8] != 0` bypasses the modulo but **not** the map-screen guard. With `campaign+0x3e0` identified as sub-ticks (`SESS-TICK-026`), the first test selects the sub-tick on which the full tick rolls and the second selects one full tick in 20 — **one unforced relight per 20 in-game minutes, 72 per in-game day**. On each, in order: `R1813(minutes)`, `R0468(1,1,0,0)` on `[view+0x80]` (the per-vertex grid), `R1819`, `R1368` (the shading-table relight driver, `TERR-LIGHT-018`), then `view+0xdc = 1`, `view+0x74 = 1` and a `0x404` post. So a changed angle reaches the screen only through that whole chain; **there is no path that re-shades without rebuilding the per-vertex grid**, which is what makes `TERR-LIGHT-022`'s single shared table sufficient

**Confidence.** High (the caller enumeration states its instrument; both gates and the whole call sequence are named instructions in one routine read whole)

**Original status.** ● active (extends `TERR-LIGHT-022`)

**Evidence.** [EXP-0116](../experiments/EXP-0116-daylight-cycle/)

### TERR-SIGHT-115

**The sight object at `world+0x58ee8` holds FOUR grids, not three, and its init builds three tables of which only two serve line of sight.** Layout, read off `R0274`, `R1847`, `R0273`, `R0136` and `R0134`: `+0x00000` a 256x256 `u16` plane, the only region the init's repeated dword store clears (`L00898`, count `0x8000`); `+0x20000` a 64x64 `u16` grid holding `di*di + dj*dj`, built by `R1847`; `+0x22000` the 64x64 `u16` step grid; `+0x24000` the 64x64 `i32` accumulator; `+0x28000` the 64x64 `i16` cost grid; `+0x2a000` `1 << k` and `+0x2a004` `k`; `+0x2a008` the `0x10000`-byte visibility byte map (= `world+0x82ef0`); `+0x3a008` the back-pointer to the world. All three 64-wide grids are written only over the 41x41 window `di,dj` in `-20..+20`, and the ring walk visits only `r <= 19`, so no unwritten cell is ever read. **The squared-distance grid is not part of the sight predicate**: `disp:20a28` returns **4 hits over 3 owners** — the builder, plus `R0409` and `R0410`, which compare it against `(r+1)*(r+1)` (`L10693`/`L10694` (a comparison of the 16-bit value), with `r = [[actor+0x154]+0x8]`) and `OR` a `u16` mask taken from `[actor+0x14]+0x2c` into the `+0x00000` plane. That is a **disc test**: no accumulator, no height term, no predecessor chain. Merging it with the line-of-sight region would give a consumer a radius where the game has a budget. This extends `TERR-FOG-088`, whose layout lists the three sight grids and not this one

**Confidence.** High (every offset is an arithmetic identity from its own scale factor in a cited instruction, the clear's extent is its own count, and the fourth grid's readers come from a named sweep with 0 orphan hits)

**Original status.** ● active

**Evidence.** [EXP-0117](../experiments/EXP-0117-sight-region/)

### TERR-SIGHT-116

**The playable rectangle is `world+0x58ee0..0x58ee3` = `(8, 8, W-9, H-9)`, written at map load, and the sight walk tests it as BYTES while indexing the map in 32-bit — which disagree at the edge of a 256-wide map.** `R0278` writes the four bytes at `L00478`, `L10695` (the literal `0x8`, twice) and `L10696`, `L10697` (a subtraction of `0x9` from each of the low bytes of `[world+0x50000]` = W and `[world+0x50004]` = H), plus the same two corners packed as words at `+0x58ee4` = `0x808` and `+0x58ee6`. `disp:58ee0..3` returns 7/7/8/8 hits over 4/4/5/5 owners, 0 orphan: this writer, the ring walk `R0134` (four compares per edge, sixteen per ring), `R0301`, `R0155` and `R1655`. The same routine fills the height plane `world+0x9451c` from the `.alm` type2 Altitudes grid (`ALM-GRID-013`), source index `(row*W + col) & 0xffff`, destination `row*256 + col` (`L10698` (a 16-bit multiply by the value at `+0x50000`) … `L10699`), which is why the sight march's `origin + (b<<8) + a` indexes it directly. **The disagreement:** `R0134`'s four compares are 8-bit (a byte comparison against the value at `+0x58ee0`), while `R0136`'s height and visibility indices are exact 32-bit, so a true column of `-9..-11` wraps to `245..247`, passes `<= W-9` when `W = 256`, and reads the previous row. It needs a ring of radius `>= 17`, so it cannot occur until sight extends far past `scanRange` down a slope: re-execution over 81 048 EN observers finds **0** such cells at `scanRange` 4, **1 096** at 7, **10 649** at 12 and **172 756** at 19; **7 of the 38 EN maps and 4 of the 34 RU maps are 256x256**. An implementation that bounds-checks in 32-bit is correct where the original is not

**Confidence.** High (all four stores, both compare widths and the height-plane copy are cited instructions; the wrap is arithmetic on the compare width) / **Medium** for the four hit counts, which quantify how often it bites and come from this round's transcription rather than from a running original

**Original status.** ● active

**Evidence.** [EXP-0117](../experiments/EXP-0117-sight-region/)

### TERR-FOG-117

**`TERR-FOG-080`'s flat-ground region is the region of a table it never finished building: `R0283`'s last four instructions were not transcribed, and the correct figure at sight 6 is 145 cells, not 127.** After the `JMP` at `L10700` that closes both of the builder's loops, `L00939`…`L00940` write four literal bytes — `[mapView+0x118a] = 0xff`, `[+0x118b] = 0x00`, `[+0x10e6] = 0x01`, `[+0x10e7] = 0x00`. Since the array is based at `mapView+0xaa8` with row stride `0x52 = 41·2` and column offset `0x28`, `0x118a` is row **21** and `0x10e6` row **19**: the cells `(+1, 0)` and `(−1, 0)`, which the slope test leaves pointing **diagonally** because `(1,0)` satisfies neither `j < i>>1` nor `j > 2i`. They are the client's copy of the four stores `AI-LOS-087` found on the server, instruction for instruction. **The identification is a reproduction, not an inference:** re-executing this builder with those four stores omitted and nothing else changed returns **59 / 127 / 223 / 479 / 847 / 1183 / 1505 / 1521** cells at axis reaches **3 / 5 / 7 / 11 / 15 / 18 / 19 / 19** and diagonal reaches **3 / 4 / 6 / 8 / 11 / 13 / 17 / 19** — all eight rows of `EXP-0085/evidence/los-field.txt` section B, cells and both reaches, **including the last two where the ring bound binds first and the patched and unpatched runs coincide anyway**. With the four stores the same run gives **69 / 145 / 249 / 521 / 897 / 1253 / 1505 / 1521**, and the sight-6 figure is then the **145** `AI-LOS-089` measured on the server, because it is the same algorithm (`AI-SIGHT-093`). Corrected flat-ground counts for `scanRange` 1..6 at `k = 7`: **9 / 21 / 45 / 69 / 105 / 145**. The row's *terrain* consequences — the per-cell `h(cell) − h(observer)` term, the plateau asymmetry, the corpus spread — rest on the recurrence and not on the tables, and are not disturbed in shape; their **numbers** were measured on the defective table and are withdrawn

**Confidence.** High (the four instructions are cited and re-read from raw bytes on both roots; the defect is reproduced exactly over eight independent rows, and a wrong diagnosis would not predict the two rows where the two runs agree)

**Original status.** ● active

**Evidence.** [EXP-0120](../experiments/EXP-0120-sight-writers/)

### TERR-FOG-118

**The playable rectangle the two implementations walk is not the same rectangle: the AI's is one cell narrower on every side, and its bounds test is byte-wide.** Client, `R0284`: each ring cell's absolute column and row are tested `>= 7` and `< map+0x4 − 7` / `< map+0x8 − 7` with **32-bit** compares (`L10561`…`L10562`, comparing the 32-bit frame local at `-0x30` with `0x7` and the later bounds), so columns `7 .. W−8` are called. Server, `R0134`: the same test runs against the four **bytes** at `world+0x58ee0..3`, written by the map load `R0278` as `8`, `8`, `(u8)(W−9)`, `(u8)(H−9)`, so columns `8 .. W−9` are called and a negative absolute column wraps into the byte range instead of failing. Measured: over 79 160 observers at `scanRange = 6`, stride 4, on 72 shipped maps and both roots, **13 362 disagree — and every one of them stands within 27 cells of a map edge**; of the 43 640 that do not, **0 disagree**. So the whole of the remaining difference on a whole-cell sight is this inset, and a consumer that shares one implementation between fog and vision must still carry two rectangles

**Confidence.** High (both tests are cited instructions with their immediates, and the 27-cell separation is an executed census that leaves no interior counterexample) / Medium for the census figures — a transcription re-executed on shipped altitude grids, `k = 7`, `scanRange = 6`, stride 4, no session

**Original status.** ● active

**Evidence.** [EXP-0120](../experiments/EXP-0120-sight-writers/)

### TERR-LIGHT-119

**The complete sky-light schedule: what every arm of `R1813` writes to the tint bytes and the two intensity fields, and how.** With `m = t mod 120` (the routine's own load of the divisor `0x78` and signed division, at `L10701`/`L10702`/`L10703`/`L10704`) and `q(k) = (k*m)/120`: **cycle off** `L10237` R=0 G=0 B=0, `0x494`=14, `0x498`=32, `0x49c`=4, `0x4a0`=2, all seven immediates. **day (6..17)** `L10705` the same seven values, again all immediates. **dawn 1 (hours 2,3)** `L10706` R=`q(24)`, G=12 literal (`L10707`), B=`48 - q(48)`, `0x494`=`32 - q(4)`, `0x498`=`8 + q(12)`. **dawn 2 (hours 4,5)** `L10708` R=`24 - q(24)`, G=`12 - q(12)`, B=0 literal (`L10709`), `0x494`=`28 - q(14)`, `0x498`=`20 + q(12)`. **dusk 1 (hours 18,19)** `L10710` R=`q(24)`, G=0 literal (`L10711`), B=`q(8)`, `0x494`=`14 + q(10)`, `0x498`=`32 - q(12)`. **dusk 2 (hours 20,21)** `L10712` R=`24 - q(24)`, G=`q(12)`, B=`8 + q(40)`, `0x494`=`24 + q(8)`, `0x498`=`20 - q(12)`. **night (22,23,0,1)** `L10713` R=0, G=12, B=48, `0x494`=32, `0x498`=8, all immediates. Both twilight bands close on a shared tail writing `0x49c`=6 and `0x4a0`=3; `0x49c`/`0x4a0` are 4/2 off and by day, 8/4 at night. This is the arithmetic `TERR-LIGHT-110` graded **Unknown** and the schedule `TERR-LIGHT-014` held at Medium with no values

**Confidence.** High - one routine read whole by a decoder whose linear tiling is the falsifier (302 instructions closing exactly on `L10714`), every immediate re-read at the address of the instruction carrying it, and the six arms verified against a continuity property the routine never asserts (`TERR-LIGHT-123`)

**Original status.** ● active (supersedes `TERR-LIGHT-014`'s colour schedule; closes `TERR-LIGHT-110`'s Unknown)

**Evidence.** [EXP-0123](../experiments/EXP-0123-sky-light-schedule/)

### TERR-LIGHT-120

**The routine reads five addresses and none of them is a table - the sky-light values are never looked up.** A linear decode of `R1813..L10714` tiles the range exactly (302 instructions, both roots), so the absolute-operand census taken from it is complete for the routine by construction rather than by search. **READ**: `L06260` (the `ShowTimeFlow` flag, `R1813`) and four `.rdata` doubles used only by the two computed-angle arms - `L07947` `L10669`, `L07945` `L10715`, `L07946` `L10655`, `L10668` `L10667`. **WRITE**: `L10231/491/492` (7 sites each), `L05650`/`498` (7 each), `L04368`/`4a0` (5 each), `L07948` (5), `L10240` (3). So every tint and intensity value is either an immediate in the instruction stream or the output of a magic division - there is no colour ramp, no per-hour table and no interpolation between stored endpoints. The three **literal** angles are likewise immediates inside `.text` (`L10656`, `L10658`, `L10661` each store the two halves as immediate stores to a 32-bit memory operand), so in this routine *literal* and *table read* are physically distinguishable and three of the five angle stores take the literal

**Confidence.** High - the census is derived from a decode whose completeness is testable (a mis-sized operand desynchronises the stream and the tiling verdict goes false) rather than from a reference-manager query

**Original status.** ● active

**Evidence.** [EXP-0123](../experiments/EXP-0123-sky-light-schedule/)

### TERR-LIGHT-121

**Six arms, not five: each twilight band splits in two at its own `CMP`, and two branches in the region are unreachable.** `TERR-LIGHT-110`'s five arms are correct for the *angle*, which the two halves of each twilight band agree on - but they write different colour ramps. The selector, all named instructions: `L10716 JL`/`L10717 JGE` peel the day band; `L10718 JL L10713` and `L10719 JGE L10713` send hours <2 and >=22 to night; `L10720 JGE L10721` splits dawn into `{2,3}` at `L10706` and `{4,5}` at `L10708`; `L10722 JGE L10723` peels dusk; `L10724 JGE L10712` splits it into `{18,19}` at `L10710` and `{20,21}` at `L10712`. Each pair converges on a shared tail - `L10660` for dawn, `L10725` for dusk - which writes `0x49c`, `0x4a0` and the literal angle, which is why the angle sees four arms where the colour sees six. **Two branches cannot be taken**: `L10726 JL L10727` tests `hour < 18` at a point only hours 18..21 reach, and its target `L10727 JL L10728` jumps to the epilogue, which writes nothing. Dead in the shipped image on the reachability the guards themselves establish

**Confidence.** High - every guard is an instruction in a routine read whole, and the split is corroborated by the two halves writing demonstrably different formulas (`TERR-LIGHT-119`)

**Original status.** ● active (extends `TERR-LIGHT-110`)

**Evidence.** [EXP-0123](../experiments/EXP-0123-sky-light-schedule/)

### TERR-LIGHT-122

**One multiplier, two post-shifts: the ramps divide by 120 and the hour selector by 60, and the ramp's clock is `t mod 120`, not the day arm's `(t+360) mod 720`.** All five divide sequences in the routine load the same `0x88888889` and add the dividend back; the head shifts by 5 (`L10729`) and all four ramp arms by 6 (`L10730`, `L10731`, `L10732`, `L10733`). Re-executed from the multiplier and shift **read out of the PE** and compared against exact integer division: /60 over inputs 0..2 000 000 and /120 over 0..10 000, **0 mismatches on either root**. The ramp phase is `m = t mod 120` from the arms' own load of the divisor `0x78`, sign-extension and signed division; because `1440 mod 120 = 0` and the argument is `fullTicks + 360` with `360 mod 120 = 0` (`SESS-TICK-026`), `m` runs 0..119 across each two-hour half on every in-game day. The day arm alone re-adds `0x168` and works modulo `0x2d0`. **The division is signed**: a negative argument would run every ramp backwards, and nothing was found that can present one

**Confidence.** High - the divisor is established by exhaustive agreement with integer division over the whole reachable input range, not by recognising a constant; both immediates are read from the image rather than transcribed

**Original status.** ● active

**Evidence.** [EXP-0123](../experiments/EXP-0123-sky-light-schedule/)

### TERR-LIGHT-123

**The six arms are one continuous 24-hour programme, and that is a property the routine never asserts.** Taking the value each arm leaves at its band's last minute against the value the next arm opens with, over all six joins, the largest step in any of R, G, B, `0x494`, `0x498` is **1**. Two joins are exact: night's `(0,12,48,32,8)` is *precisely* what dawn 1 opens with, and day's `(0,0,0,14,32)` is *precisely* what dusk 1 opens with. Over a whole in-game day every one of the five values stays in `[0,48]` - no 8-bit store wraps and no subtraction goes negative, although four of them are 8-bit `SUB`s (`L10734`, `L10735`, `L10736`, `L10737`) that would wrap if it did. This is the round's discriminating test: it is not stated anywhere in the image, and a numerator, divisor or sign read wrong breaks it at once - an interim reading that took the post-shift for /60 produced `B = -47` at the end of dawn 1 and a 95-step discontinuity at that join

**Confidence.** High - a property computed over all 1440 minutes from constants read at their own addresses, which the published decode satisfies and the rival decodes do not

**Original status.** ● active

**Evidence.** [EXP-0123](../experiments/EXP-0123-sky-light-schedule/)

### TERR-LIGHT-124

**`L10231`, `L10613`, `L10614` are R, G, B in ascending address order - the rival BGR reading is excluded twice over.** The three bytes have exactly **one** reader outside `R1813`, found by a whole-image raw dword scan (8 references each, 7 of them the routine's own stores): `R1107` at `L10738`/`L10739`/`L10740`, `TERR-LIGHT-018`'s shading-table builder. In its terrain arm (mode 3) `[0x490]` is added to palette entry **byte 2**, `[0x491]` to byte **1**, `[0x492]` to byte **0** (`L10741`/`L10742`/`L10743` against `L10744`/`L10745`/`L10746`). Under the game's own palette entry order `[B,G,R,x]` (`SPR256-PAL-011`, pinned there by a known-answer test on a shipped VGA palette) that makes them R, G, B. **Independently**, each result is packed with the bits/shift pair `TERR-LIGHT-019` already assigns to that channel - `L06281`/`L10234` for `[0x490]` (R), `L10235`/`L10236` for `[0x491]` (G), `L06282`/`L06471` for `[0x492]` (B). Two instruments with different failure modes, one answer: the night sky is **blue**, not red

**Confidence.** High - the competing model is named and excluded, and by two routes that do not share an assumption; corpus agreement plays no part

**Original status.** ● active

**Evidence.** [EXP-0123](../experiments/EXP-0123-sky-light-schedule/)

### TERR-LIGHT-125

**What the schedule buys: flat ground runs level `[45,64]` over an in-game day and every sprite on the map dims through six steps together.** Evaluating `TERR-LIGHT-013` (`R = [0x498]`, `L = (R>>1) + [0x494] + 0x20`, a flat vertex having `s = 0` on both axes so its level is `ftol(L - R*sin(pi/6))`) on this round's schedule: the expression collapses to `[0x494] + 32` for even `[0x498]` and `[0x494] + 31` for odd, so the flat-ground level is **46 by day and with the cycle off** - the figure `TERR-LIGHT-013` already publishes by a different route - **64 at night**, and `[45,64]` over the day. The slope amplitude `[0x498]` itself falls 32 to 8 from day to night, so terrain relief **flattens** as it darkens. On the sprite side `TERR-LIGHT-061`'s per-frame grid is memset from `[0x494] >> 2`, which takes the values 3 (day) to 8 (night) through 5, 6, 7: a single map-wide byte, so objects and units dim in step and never independently

**Confidence.** Medium - the schedule is this round's at High, but both transforms it is evaluated through are `TERR-LIGHT-013` and `TERR-LIGHT-061`, so the figures inherit those rows' standing; nothing here was observed running

**Original status.** ● active

**Evidence.** [EXP-0123](../experiments/EXP-0123-sky-light-schedule/)

### TERR-LIGHT-126

**The two shroud levels are on the day/night schedule too, and one is always twice the other.** `TERR-SPR-066` already establishes what `L04368` and `L04335` *are*: the **fourth argument** of the silhouette blit `vt+0x3c`, an index into the destination-recolour table `[L03341]` at stride `[L03340]` - `0x49c` on the object path, `0x4a0` on the unit path's alternate arm (and `HERO-APPEAR-056` for the second hero sheet). What was not established is that they **move with the band**. `R1813` writes both in every arm: **4/2** with the cycle off (`L10747`/`L10748`) and by day (`L10749`/`L10750`); **6/3** in both twilight tails, the 6 arriving as a value set once in the prologue (`L10751`, loading `0x6`) and stored at `L10660`/`L10725`, the 3 as an immediate at `L10752`/`L10753`; **8/4** at night, the 8 arriving as a loaded value (`L10713`, loading `0x8`) and stored at `L10754` alongside `0x498`, the 4 at `L10755`. So a shadow is recoloured through row 2 of that table by day and row 4 at night on the unit path, and 4 -> 8 on the object path. **`0x49c` = 2 * `0x4a0` in all four settings**, and no instruction computes one from the other - both are written as separate constants, which is what makes the relation a fact about the values rather than about the code. Complete reader enumeration from a whole-image raw dword scan: `0x49c` 11 references, `0x4a0` 14, the non-routine ones at `L10756`, `L10757`, `L10758` and `L10759`, `L10760`, `L04336`, `L10761`, `L10762`, `L10763`, `L10764`, `L10765`, `L10766`, `L10767`. **Still open**: why the object path's level is exactly twice the unit path's, and whether any consumer reads both in one frame

**Confidence.** High for the values, the schedule and the doubling relation (named immediates in a routine read whole, plus a raw scan whose blind spot is stores through a register-held pointer); the identification is `TERR-SPR-066`'s, not this round's

**Original status.** ● active (extends `TERR-SPR-066`)

**Evidence.** [EXP-0123](../experiments/EXP-0123-sky-light-schedule/)

**Amended.** [EXP-0449] (`TERR-191`) narrows one clause: a unit's shadow is not recoloured only through row 2 by day and row 4 at night. An ordinary unit's first silhouette reads `[L04368]` (4 by day, 6 in twilight, 8 at night) and only its second silhouette, or an invisible owner's first, reads `[L04335]` (the `0x4a0` values). The second sentence of the open question (whether any consumer reads both cells in one frame) is answered: one call of `R0553` does, for a unit without effect key `0x26`. The reason for the factor of two stays open. The values, the schedule and the doubling relation are unchanged.

### TERR-LIGHT-127

**The sky tint can only ever add: it is taken unsigned and applied before the level multiply, so nothing in the day/night cycle darkens a colour channel.** The builder loads each of the three bytes and masks it a mask with `0xff` (`L10768`, `L10769`, `L10770`) before the per-entry loop, so a value is a 0..255 addend and never a signed offset - and `TERR-LIGHT-119`'s schedule never produces one above 48 in any case. The transform is `TERR-LIGHT-019`'s `out = clamp(((chan + tint)*m)/32, 0, 255)`: the tint is inside the multiply, so it is scaled by the level along with the palette channel rather than added to the result. **Consequence for a consumer**: all darkening across the cycle is done by `0x494`/`0x498` through the level, and the tint only ever pushes a channel up and can only ever push it towards `TERR-LIGHT-063`'s saturation - a night sky is blue because +48 blue survives a dim that costs the other two channels more in absolute terms, not because red and green were subtracted

**Confidence.** High - the mask, the operand order and the clamp are named instructions in a loop read end to end, and the schedule's range is measured over every minute of a day

**Original status.** ● active

**Evidence.** [EXP-0123](../experiments/EXP-0123-sky-light-schedule/)

### TERR-LIGHT-128

**The customisation limit (G2): the whole day/night schedule is 70 immediate fields, 181 bytes, inside one routine in `.text` - there is no data file, registry key or table to change.** Enumerated and re-read at their own addresses, each verified to *end* on an instruction boundary of the linear decode: 9 fields for the cycle-off arm, 7 for day, 6 for night, 4+5 for the two dawn halves, 3+4 for the two dusk halves, 7 in the two shared twilight tails, 5 band-boundary immediates, and the divider's multiplier and post-shift at each of the five divide sites. All 70 are byte-identical between the two lawful roots. **So the schedule is not data.** Its class is *compiled constant*: lifting it means patching `.text` at listed offsets, and the schedule's **shape** is fixed harder than its values - the number of bands, the two-hour half-band, the affine-plus-floor-divide form and the `120` modulus are encoded in the control flow and in the numerator scales (`LEA`/`SHL` operands), not in an immediate a table could redirect. Changing a shipped file's bytes: **none** - `terrain.3d`, the `.alm` maps and the `.reg` registries are untouched by any of it, and `TERR-LIGHT-023`'s map-stored light fields have no reader on this path. Two further limits a consumer inherits: only **24** distinct `(R,G,B,0x494,0x498)` tuples are reachable through the unforced relight cadence out of **170** the arithmetic can produce (`TERR-LIGHT-114`'s `t mod 20 == 0` samples each 120-minute ramp six times), and the forced path bypasses the modulo and can land on any of the 170

**Confidence.** High - the enumeration is checked against the routine's own instruction boundaries rather than against a listing (the check caught two mis-cited offsets in this round's first table), and the reachable-state counts are computed over every minute of an in-game day

**Original status.** ● active

**Evidence.** [EXP-0123](../experiments/EXP-0123-sky-light-schedule/)

### TERR-SHDW-129

**The silhouette blitter shears about the BOTTOM edge of its own rectangle, and the sign says which way a shadow leans.** The `CSprite256 vt+0x3c` thunk `R1831` (returning with `0x18` bytes released, six arguments) hands **seven** to the core, `cdecl` (the caller releases `0x1c` bytes): `(dstX, dstY, width, height, pixelRun, shroudLevel, shear)`, where `width`/`height` are `frame+0x0`/`frame+0x4` of the **drawn** frame and are added by the thunk - the caller never passes a size, so `TERR-SPR-066`'s "16.16 per-row X slope" is bounded by the frame's own height and by nothing the caster chooses. In `R1832` the whole shift is computed first, `totalShift = trunc(height*shear/65536)` (`L10771` multiplies, `L10772`/`L10773`/`L10774` add a sign bias masked with `0xffff`, `L10775` shifts right by `0x10`), and the destination pointer is **pre-advanced by all of it** before a single row is drawn - `L10776`/`L10777` adding the frame local at `-0xc` to the destination pointer (twice: 16 bpp). The accumulator (the argument at frame offset `+0x1c`) starts at **0** (`L10778`) and each row walks the pointer back **left**: `acc += shear`, then `whole = acc & 0xffff0000`, `step = whole >> 15`, `acc -= whole`, then the destination pointer is moved back by `step`. So the column of image row `r` is **`dstX + trunc(h*shear/65536) - floor(r*shear/65536)`** - re-executed instruction for instruction over 448 `(shear, height)` pairs spanning the day band and far outside it, **0 rows disagree** (`tools/shearwalk`). The row that lands on the `dstX` passed in is therefore `r = h`: **the pivot is the bottom edge of the blit rectangle**, the top row displaced by the whole `h*tan(theta)`, and a positive shear puts the top to the **right**. It is a shear and not a translate: the two halves of a silhouette lean opposite ways about the pivot row. **Row bound**: the row counter, set to `height` at `L10779`, decremented per row and by the RLE's row-skip opcode, which advances the accumulator by exactly that many rows (`L10780`, `L10426..L10781`) - a skipped row shears identically to a drawn one. **Column bound**: the frame local at `-0x4` is below `width` at `L10782`. **Mirroring does not flip the lean**: `R1848`/`R1849` run with the direction flag set from the right end of the row but keep the same pre-add of the frame local at `-0xc` to the destination (`L10783`) and the same per-row backward step (`L10784`). The fast path checks nothing inside the loop, and its accept test is **tight**: over the same sweep the loop never writes left of `dstX-abs(totalShift)` and writes at most one column right of `dstX+abs(totalShift)`, which the test's own `dstX+width` bound absorbs since the last column of a row is `dstX+width-1`

**Confidence.** High (the argument list is closed through the thunk rather than inferred from a caller, every term of the accumulator is a named instruction, and the closed form is checked by re-execution rather than by inspection)

**Original status.** ● active

**Evidence.** [EXP-0129](../experiments/EXP-0129-shadow-destination/)

### TERR-SHDW-130

**The three shadow casters' destinations, term by term.** Writing `T = ftol( tan(theta) * (2*floor(frameH/2) - anchorY) )` - a **pixel** count - and `s = ftol( tan(theta) * 65536 )` - the 16.16 slope of `TERR-SHDW-129`: **(a) unit** `R0553` (`CUnit vt+0x2c`): `dstX = unit+0x60 - anchorX - T`, `dstY = unit+0x64 - anchorY - unit+0x68` (`L07949..L07950`, `T` loaded at `L07951` from the slot written at `L07954`). **(b) placed object**, inside `R0379`: `dstX = col*32 + 16 - anchorX - T`, `dstY = row*32 + 16 - anchorY - alt` (`L10275..L10276`, `T` at `L10785`). **(c) structure strip** `R1840` (`CStructure vt+0x2c`): `dstX = col*32 + ftol( tan(theta) * ((FullHeight - k)*32 - ShadowY) )` and `dstY = row*32 - this+0x68 - this+0x10`, decremented by 32 per strip (`L10786..L10787`, `L10685..L10788`, `L10789`); there is **no `+16`** and no anchor, which is `TERR-STRUCT-100`'s "no anchor pixel and no canvas" seen from the draw side. All three push the same six-argument list `(dstX, dstY, frame, shroudLevel, s, mirrorFlag)` - `L10790`, `L10274`, `L10625`. **The anchor is shared and the arithmetic is exact**: `anchorX = floor(frameW/2) + CenterX - floor(Width/2)`, `anchorY = floor(frameH/2) + CenterY - floor(Height/2)`, the two `SAR r32,1` being separate truncations - so `T`'s multiplicand `floor(frameH/2) + floor(Height/2) - CenterY` equals `frameH - anchorY` for **even** `frameH` and one less for odd. The class record differs by caster and the field offsets differ with it: `CUnit` class `+0x2c` Width / `+0x30` Height / `+0x34` CenterX / `+0x38` CenterY (`L10791`, `L10792`), the object class `+0x18` / `+0x1c` / `+0x20` / `+0x24` (`L10793`, `L10794`, `L10282`, `L10795`) - a consumer sharing one struct across both reads garbage. `frameW`/`frameH` come through `vt+0x20`/`vt+0x24` **at the drawn frame**, not frame 0, on all three paths, which is `TERR-SPR-043` holding for the shadow as well as the body

**Confidence.** High (every term is a named instruction and every stack slot is tracked through the intervening pushes, which is where `TERR-SPR-067` mis-named one of these two sun-derived quantities as the other)

**Original status.** ● active

**Evidence.** [EXP-0129](../experiments/EXP-0129-shadow-destination/)

### TERR-SHDW-131

**The corpus's shear "addition vs subtraction" is not a contradiction: it is one rule written in two coordinate directions, and the caster's scalar moves the blitter's pivot to the row the caster wants.** `TERR-SPR-038` and `TERR-SPR-067` publish a **subtraction** on `dstX`; `TERR-STRUCT-103` publishes an **addition**; nothing reconciled them. `TERR-SHDW-129` fixes the blitter's own pivot at the bottom edge of the blit rectangle, so a caster that wants any other pivot must pay the difference in `dstX`. Substituting the unit and object rule `dstX = base - T` into the blitter's `X(r) = dstX + trunc(h*s/65536) - floor(r*s/65536)` gives **`X(r) = base + tan(theta)*(anchorY - r)`** - the shear vanishes at `r = anchorY`, the **anchor row**. Measured against that identity over 120 `(shear, frameH, anchorY)` triples with both of the engine's truncations in place, the residue is **at most 1.99 px and does not scale with `frameH`** (`maxDev/frameH` <= 0.06 and falling). That is what discriminates: a pivot left at the frame **bottom** would deviate by `tan(theta)*(frameH - anchorY)`, up to **55 px** at `frameH = 96` and `tan = 0.577`, i.e. two orders of magnitude outside the measured residue. The structure's `+` is the same relocation read upward instead of downward: `(FullHeight - k)*32` **grows as `k` falls**, and `k` falls as `dstY` rises - `L10789` (a subtraction of `0x20`) and `L10796` (a decrement) are in one basic block - so it is a height **above the image bottom** where `anchorY` is a depth **below the frame top**. Moving a pivot up costs a subtraction and moving it down costs an addition; the two casters do exactly one each. **Amends `TERR-SPR-038` and `TERR-STRUCT-103`**: both signs are right as written and neither may be "corrected" into the other. A consumer that copies one caster's sign onto the other draws the shadow leaning from the wrong end. **Left open**: for the structure the pivot is `X = col*32` at the screen row `dstY_k + h` of the strip with `(FullHeight - k)*32 = ShadowY`, where `h` is that strip frame's own height; if the strip frames are 32 px that is `ShadowY - 32` px above the image bottom, one cell row below what `TERR-STRUCT-103`'s wording implies, and the frame heights are not read here

**Confidence.** High (the reconciliation is forced by the direction of the loop's own decrement rather than chosen, and the anchor-row pivot is discriminated by a residue that is bounded and shown not to grow with the frame height - corpus agreement plays no part) / **Unknown** (where `ShadowY` sits relative to the structure image's bottom, pending the strip frame height)

**Original status.** ● active

**Evidence.** [EXP-0129](../experiments/EXP-0129-shadow-destination/)

### TERR-SHDW-132

**The unit's flat blit arm is the sheared arm's own rule collapsed to a single 32-pixel step, not a different placement.** `R0553` has **four** arms, two on `vt+0x3c` and two on `vt+0x1c`, crossed with the two shroud globals `[L04368]` and `[L04335]`. The `vt+0x3c` arms push `dstX` as `TERR-SHDW-130` computes it (`L10797`, and `L10798` is the other pair's). The `vt+0x1c` arms **add `T` back** and then translate by `shear/2000`: `L10799` (loads the constant `0x10624dd3`), `L10800` (multiplies by the stack argument at offset `0x20`), `L10801` (shifts right by `0x7`) and `L10802` (adds) is signed division by 2000, then `L10803` adds the `dstX` that already has `-T` in it, and `L10804/L10805` add `T` - the two `T` cancel. So the flat arm draws at `unit+0x60 - anchorX + shear/2000`, and since `65536/2000 = 32.8`, that translate is **one cell row's worth of lean**: `tan(theta)*32.8`, against the sheared arm's `tan(theta)*anchorY` at the top of the sprite. `TERR-SPR-066` recorded the `/2000` as making "a skewed and a translated shadow"; the cancellation is what makes them the *same* shadow at two levels of detail, and it is why a consumer must not implement the flat arm as "the shadow with the shear argument set to zero" - that would lose the 32-px offset and put the flat shadow under the sprite instead of beside it

**Confidence.** High (the divide-by-2000 idiom, both `ADD`s and the slot they read are named instructions, and the cancellation is arithmetic on terms this ledger already names)

**Original status.** ● active

**Evidence.** [EXP-0129](../experiments/EXP-0129-shadow-destination/)

### TERR-SHDW-133

**A shadow is placed with ONE height for the whole silhouette, sampled once, at the caster's own cell - established positively.** A silhouette covers pixels belonging to other cells, so a per-row or per-strip height lookup would be the natural way to make it follow the ground. There is none. **(a)** `R0553` (unit shadow) **never reads any of its three `Draw(col,row,alt)` arguments**: the frame is a `0x28`-byte reservation + 4 pushes, so they sit at base+0x3c/+0x40/+0x44, and the only two reads at stack offset `0x3c` in the whole routine (`L10326`, `L10327`) both occur with five pushes pending and therefore address base+0x28 - a local written at `L10328`. **(b)** `R1840` (structure shadow) reads `col` and `row` once each before any push (`L10806`, `L10807`) and **never reads `alt` at all**: stack offset `0x38` is read nowhere in the routine, at any stack depth. **(c)** The object path samples the smoothed altitude plane **exactly once**, `[view+0xc0]` at index `(col+4)+(row+4)*(cols+8)` (`L10280..L10808`) - `TERR-SPR-041`'s field and index, re-derived - and that value reaches `dstY` (`L10809`) and nothing else. **What rules the alternative out** is not failing to find a second sample: it is that the argument slots are provably unread with the stack tracked through every push, and that `dstX` and the shear are each traced to a complete set of named terms, none of which is a field read. The consequence is sharp: **a shadow does not drape.** It is one flat sheared stamp placed by the caster's own cell height, so a shadow crossing a slope, a cliff edge or a neighbouring cell of different altitude does not bend, and a consumer that samples terrain per silhouette row is not reproducing the original

**Confidence.** High (an absence carried by an enumeration of every read of the argument slots with the stack tracked, plus a complete term-by-term trace of both destination components - the only instruments that can carry an absence)

**Original status.** ● active

**Evidence.** [EXP-0129](../experiments/EXP-0129-shadow-destination/)

### TERR-SHDW-134

**One slope per invocation; the terrain height enters neither the slope nor anything it multiplies.** The 16.16 argument `s = ftol(tan(theta)*65536)` is computed **once** on every path and never recomputed per row, per cell or per strip: unit `L10338..L07952` in the prologue, structure `L10810..L10811` **before** the strip loop, object `L10422..L07860` once per drawable. The structure's per-strip block re-enters `R1820` at `L10812` with the **same receiver** (the stack argument at offset `0x20` -> `view+0x80`, the one cache of `TERR-LIGHT-111`) so the tangent is the same value each time; what is per-strip is only the *multiplicand* `(FullHeight - k)*32 - ShadowY`, and the slope pushed at `L10813` is the one word written at `L10811`. The unit computes its **pixel** term `T` twice, once per sprite group (`L10814..L10815` for the `+0x194`/`+0x198` overlay sheets, `L10816..L10817` for the `+0x04`/`+0x08` main sprites), because `T` depends on the drawn frame's own height - but the slope is shared by all four draws. **Altitude appears in none of it**: `s` is a function of the cached sun angle alone; `T`'s multiplicand is `floor(frameH/2) + floor(Height/2) - CenterY`, all frame and class geometry; the structure's is `FullHeight`, `k` and `ShadowY`, all class geometry. So the lean of a shadow is independent of the ground it falls on, and identical for every drawable in the frame - which is the same thing `TERR-LIGHT-113` says from the sun's end, now closed from the blitter's

**Confidence.** High (each computation site and its receiver is a named instruction, and the absence of a height term is read off a complete enumeration of the multiplicands rather than from not having noticed one)

**Original status.** ● active

**Evidence.** [EXP-0129](../experiments/EXP-0129-shadow-destination/)

### TERR-SHDW-135

**`vt+0x3c` is a family of four cores behind two globals, and it never reads a source pixel - the "silhouette" is literal.** `R1831` dispatches on `arg6` (the mirror flag `TERR-SPR-038`'s amendment identified) and on `[L03346]`: not mirrored gives `R1832` or `R1850`, mirrored gives `R1848` or `R1849` (`L10818`, `L10819`, `L10820`, `L10821`). All four carry the identical five-instruction shear accumulator (`L10822`/`L10823`, `L10824`/`L10825` and their clipped twins), so the geometry of `TERR-SHDW-129` is the family's, not one core's. In a **drawn** run the core reads the pixel **already on screen**, `L10826` (loads the 16-bit destination word), `L10827` (shifts right by `0x3`) and `L10828` (loads the table entry at `2 × index`) where the table base is `[L03341] + level*[L03340]*2`, and writes it back - while the source bytes of that run are consumed at `L10829` (advancing the source pointer) and **never looked at**. The sprite therefore contributes only run lengths, i.e. a mask; every shadow pixel's colour comes from the destination through the level table, which is `TERR-SPR-066`'s destination-recolour table seen from the consumer side. The clipped path does the same by another route, `L10830` (loads the destination word), `L10831` (shifts right by `0x2`) and `L10832` (masks with `0xfffe`) - `(d>>3)*2` written as `(d>>2) & ~1` - and adds the only per-pixel bounds tests in the routine, tracking screen X in the frame locals at `-0x4` + `-0x10` and screen Y in the frame local at `-0x14`. Three globals gate whether a shadow is drawn at all: the object shadow is skipped entirely on `L10833` (a comparison of `[L06416]` with zero, skipping when zero), the structure's `b`-overlay shadow on `L10834 [L04367]`, and `[L05659]` zeroes the structure's frame-variant selector at `L10835`

**Confidence.** High (the dispatch, the four cores' shared idiom, both recolour forms and the three gating compares are named instructions) / **Medium** (that the three gating globals are user-facing settings rather than internal state - nothing here reads their writers)

**Original status.** ● active

**Evidence.** [EXP-0129](../experiments/EXP-0129-shadow-destination/)

### TERR-SHDW-136

**A structure strip carries BOTH shear terms at once, and on 15 of 66 shipped classes `ShadowY` is not a pixel height at all - it is a suppression sentinel enforced by displacement, with no branch anywhere.** **(a) Both terms.** At `L10836` (a load of the stack slot at offset `0x14`) no push is pending, so that slot is the one written at `L10811` from `ftol(tan(theta)*65536.0)` - the **live 16.16 slope**, pushed into the blit's slope argument at `L10813` and again at `L10837` for the `b`-overlay. So `TERR-STRUCT-103`'s per-strip term and `TERR-SPR-066`'s per-image-row slope are **not alternatives**: one strip is displaced rigidly by one integer *and* each of its rows is displaced again inside the blit, the two composing into `TERR-SHDW-131`'s single pivot. A consumer implementing only the per-strip term draws a stack of rectangles that step but do not lean. **(b) `ShadowY` is read once and never tested.** It is `class+0x30` (`REG-STR-080`), its key name is the literal `ShadowY` at `L10838` pushed at `L10839`, and the loader stores the getter's result raw - `L10840` (the call to `R0452`) and `L10841` (a store of the result at `+0x30`) - with no clamp. On the draw path the only read is `L10842` (an integer load of the dword at `+0x30`); all fifteen compares in `R1840` were enumerated (`L10843`, `L10844`, `L10845`, `L10846`, `L10847`, `L10848`, `L10849`, `L10850`, `L10851`, `L10852`, `L10853`, `L10854`, `L10855`, `L10856`) and **none touches it**. **(c) The shipped range breaks the pixel reading.** Over all 66 classes: 28..55 on 51 of them, then `10000` on 4 (`well1`, `well2`, `well3`, `magic`) and `20000` on 11 (`cave`, `bridge1v`, `bridge2`, `bridge3`, `bridge4`, `Grave1`..`Grave4`, `Teleport`, `campfire`) - against a maximum `FullHeight*32` of **192 px** over the same roster. **(d) What those values do.** `R1820` returns `+-0.05` in a dead band and `theta*2/3` outside it (constants at `L05652..L10676` read through the PE's section table: `0.0`, `-0.05`, `+0.05`, `2/3`), so the smallest non-zero magnitude reaching `FPTAN` is `0.0333`. With `ShadowY = 20000` the strip's `dstX` sits **660 to 11 400 px** from `col*32`, outside every viewport width the engine opens (`MISSION-VIEW-020`), and the blitter's own reject test (`TERR-SHDW-129`) discards it. The fifteen are the flat and the hollow - four bridges, four graves, a cave, three wells, a teleport, a campfire, `magic` - which is the data-side counterpart of the bare `0xc`-byte-releasing return `TERR-STRUCT-103` records for the two `VariableSize` bridge subclasses, needing no code. A consumer that treats `ShadowY` as a pixel height, or that clamps it, gives fifteen classes a shadow the original does not draw

**Confidence.** High (the slope's push site is a named instruction at a stack depth tracked through the pushes; the field's key name is read from the PE at the address the loader pushes; the compare enumeration is over a routine read end to end; the census is the whole shipped roster) / **Medium** (that the two magnitudes are *authored* suppression rather than an accident - the discriminator offered is the semantic one, that all fifteen are flat or hollow objects, and a semantic argument is corroboration, not proof)

**Original status.** ● active

**Evidence.** [EXP-0129](../experiments/EXP-0129-shadow-destination/)

### TERR-SPR-137

**(rom.exe) The map painter has nine cell sweeps and one non-cell collection walk in its ten-phase composition.** The former ten-double-nested-loops-plus-list count is corrected by ANIM-WALKORDER-088. Retain phase labels:1 structure shadow;2 flat structure body;3 selector4/CBackPack shadow/body within a cell;4 non-flat structure, selector2, static-object and area-selector1 work within a cell;5 selector3 shadow;6 collection;7 area-selector0;8 selector3 body;9 marker/bar;10 shroud. The complete function isR0379..L01621 and the lowest direct back-edge target remainsL10857. The prefix contains no direct backward branch; its terrain/rect work precedes these cell loops. Registration/storage populations come from ANIM-REGISTER-083 and ANIM-CATEGORY-084; shared draw functions do not merge those populations.

**Confidence.** High for the original direct back-edge set, loop headers, grid loads and dispatch anchors. No claim that every virtual call paints pixels or that this enumerates nested helper internals.

**Original status.** ✔ promoted (partially retracted)

**Evidence.** [EXP-0130](../experiments/EXP-0130-shadow-bounds/); [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/)

**Amended.** The table ledger carried the status "✔ promoted (partially retracted)". `retracted.md` records a correction against this claim.

### TERR-SPR-138

**(rom.exe) Drawable cell sweeps walk rows ascending and columns descending, with a narrower marker/bar window and a column-major shroud.** Phases1..5,7..8 start row=-4, continue while row<view+0x68+8, and walk columns view+0x64+3 down through -4. Phase9 instead starts row=-3, continues while row<view+0x68+7, and walks columns view+0x64+2 through -3 (L10858..L10859). The old claim that this phase repeated the wider header is retracted. Phase10 starts column0 below view+0x64 and row0 below view+0x68+4, both ascending. Grid indexing and the static-object caller identify the axes. A later cell in one drawable sweep has greater row or equal row and smaller column; another phase is a separate ordering dimension (ANIM-WALKORDER-088).

**Confidence.** High for the named induction variables, bounds and axis arithmetic. Visited cells are not synonymous with admitted dispatches or opaque pixels.

**Original status.** ✔ promoted (partially retracted)

**Evidence.** [EXP-0130](../experiments/EXP-0130-shadow-bounds/); [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/)

**Amended.** The table ledger carried the status "✔ promoted (partially retracted)". `retracted.md` records a correction against this claim.

### TERR-SPR-139

**(rom.exe) The +0x90 shadow and body dispatches are separate complete cell sweeps, conditionally populated through CAirUnit selector3.** ShadowL02993 precedes the non-cell collection walkL02994, retained-area selector0 sweepL02995 and bodyL02996. The intervening work is one collection walk plus one cell sweep, not two cell sweeps. Both late cell paths check the same coordinate bounds, non-null entry and +0x78==0; the shadow sweep has an additional outer global-enable gate. Thus shadow calls precede body calls in the control flow when enabled, but the old identical-admitted-pixels/unconditional overpainting wording is narrowed. Ordinary/alternate CUnit uses earlier cell compositions (ANIM-CELL-085, ANIM-AIRPASS-086, ANIM-DRAWGATE-087).

**Confidence.** High for the static gate/grid/dispatch order. No native stable-population, pixel coverage or unconditional body-over-shadow result.

**Original status.** ✔ promoted (partially retracted)

**Evidence.** [EXP-0130](../experiments/EXP-0130-shadow-bounds/); [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/)

**Amended.** The table ledger carried the status "✔ promoted (partially retracted)". `retracted.md` records a correction against this claim.

### TERR-SPR-140

**The only bound on a drawn silhouette besides its own frame is ONE global device-space clip rectangle, and the blitter enforces it per row and per pixel — there is no per-cell destination window anywhere on the draw path.** The rectangle is the four globals `L01503` (left), `L01504` (top), `L01505` (right), `L01506` (bottom). `R0373(const RECT*)` copies four dwords into them and `R1822(l,t,r,b)` does the same by value; there is **no `IntersectRect`, no clamp against the surface and no second rectangle** — the setter replaces. Image-wide those four addresses carry 70/89/78/89 references over 39–41 owners with **0 in orphan or undisassembled code** (`EnumRefs refto:`, repaired function table), and every reference that is not one of those two setters plus `R1691` is a READ, by the whole `R0352…R1784` blitter family. `R0379` sets it four times, always to `*(CMapView+0xf4)` (`L10860`, `L10861`, `L10862`, `L10863`), each time after restoring a saved copy taken by `R0372` = `CopyRect(dst, &clipRect)` through user32 `[L01508]`; **the sweeps themselves set nothing.** In `R1832`, the sheared silhouette blit `(dstX, dstY, w, h, src, level, shear)`: `L10864…L10865` forms `span = abs((shear*h)>>16)`, then four tests take the unclipped loop only if `dstX-span >= left`, `dstX+w+span <= right`, `dstY >= top` and `dstY+h <= bottom` (`L10866`, `L10867`, `L10868`, `L10869`). Otherwise `L10870…L10871` tries a whole-sprite reject on the same four edges and, failing that, the loop tests the **current row** against top and bottom (`L10872`, `L10873`) and **each pixel's** sheared X against left and right (`L10874`, `L10875`). So a silhouette leaving the rectangle is **cut**, never dropped — the opposite of the terrain blitters, which drop a cell whole when its 32-px X range is not entirely inside (`formats/terrain`, EXP-0035). What is **not** established here: the runtime value of `CMapView+0xf4`

**Confidence.** High (both setters, the four globals with their image-wide reference sets and orphan count, all four bounding-box tests and both clip loops are named instructions)

**Original status.** ● active

**Evidence.** [EXP-0130](../experiments/EXP-0130-shadow-bounds/)

### TERR-SPR-141

**The shroud cull `R1401` has exactly ONE caller in the image, and it culls an object's shadow and body TOGETHER — one boolean per object, nothing per row.** `EnumRefs callto:R1401` on the repaired function table returns `L02992` inside `R0379`: 1 hit, 1 owner, **0 in orphan or undisassembled code**. At that site `L10876` (a zero test) and `L10877` (a branch to `L10878` when non-zero) skips past **all four** blits of sweep 4's inlined draw — the two silhouette calls `L10274` and `L10302` (calls through the vtable slot at `+0x3c`) and the two lit calls `L10303` (slot `+0x14`) and `L10304` (slot `+0x34`). `TERR-SPR-041` read the routine correctly and stated its scope more narrowly than the code does: the test culls the **shadow as well as the body**, and it is the *only* per-draw shroud test in the image — sweeps 5 and 8 have six guards each, the same six, and no call among them (`TERR-SPR-139`), so no unit shadow is subject to it or to anything like it. The routine's own shape: a column window `col-1 … col+2` clamped to `[0, view+0x64]` (`L10879…L10880`) crossed with a row window `R0377(x, dstY) … R0377(x, dstY+frameH)+1` clamped to `[0, view+0x68+4]` (`L10881…L10882`), returning 1 only if `view+0xa0[(r+3)*(view+0x64+7)+(c+3)] == 0x10` at every cell in it (`L10883`) — and level 16 is exactly black (`TERR-FOG-084`), so the cull fires exactly where the draw would be invisible

**Confidence.** High (the caller enumeration states its instrument, its orphan count and its blind spot; the branch target, the four dispatches it passes and the routine's two windows are named instructions)

**Original status.** ● active

**Evidence.** [EXP-0130](../experiments/EXP-0130-shadow-bounds/)

### TERR-FOG-142

**The shroud sweep is the LAST cell loop in the painter, after every sprite — so a shadow reaching a fogged cell is drawn in full and then covered, neither clipped to the fog nor skipped.** `TERR-FOG-083` established what that sweep consumes and how it dispatches; what it did not establish is **where in the frame it runs**, and that is what decides a silhouette's visible edge. It is sweep 10 of `TERR-SPR-137` (`L10884`/`L10885`, back-edges `L10583`/`L10886`), the final cell loop before the routine's epilogue, and it is clipped to the same `*(CMapView+0xf4)` as the sprite block (`L10862`, `TERR-SPR-140`). Combined with `TERR-SPR-141`, the **only** fog-driven suppression of a draw anywhere in the image is the object path's all-or-nothing cull, and it fires only when *every* covered cell is at level `0x10`: a silhouette crossing a partly fogged region is drawn whole and then overpainted, so **the visible edge of a shadow at fog is the shroud's boundary and not the shadow's** — a per-vertex gradient (`TERR-FOG-083`), never a cell boundary. On a fully visible cell (all four corners 0) the shroud draws nothing at all (`L10584`), so there the shadow stands as drawn

**Confidence.** High (the sweep's position in the back-edge set, its bounds and its clip are named instructions) / **Medium** (that being drawn later is *sufficient* to hide the shadow — what the six shroud blitters put on a partly-lit cell is `TERR-FOG-083`'s ramp, and this row establishes ordering rather than opacity)

**Original status.** ● active

**Evidence.** [EXP-0130](../experiments/EXP-0130-shadow-bounds/)

### TERR-LIGHT-143

**Nothing stops the shadow map being applied twice, so two overlapping silhouettes COMPOUND: the intersection is `((16-L)/16)²` of the ground, not `(16-L)/16`.** `TERR-FOG-084` established the table's law and `TERR-LIGHT-059` that the silhouette blits read the destination rather than the source; neither asks what happens when two silhouettes land on the same pixel, and the sun shear (`TERR-SPR-066`) makes that the ordinary case for two casters standing close. Read at the pixel loop, `R1832` writes `table[level][dst >> 3]` back over the destination with **no guard of any kind** — `L10826`, `L10828` and `L10887` (load the 16-bit destination word, load the table entry at `2 × index`, store it back to the previous pixel) — and there is no per-pixel state anywhere on the path that could carry one: the destination is the only memory the pass reads. The clipped path computes the same address by a different route (`L10831`/`L10832` (shift right by `0x2`, mask with `0xfffe`) against the fast path's `L10827` (shift right by `0x3`)), provably equal because the index's upper half is 0 from the mask `0xc03f` at `L10888`, so clipping does not change the arithmetic either. The rival — an idempotent table — is refuted by `TERR-FOG-084`'s law itself: `(16-L)/16` is multiplicative and its only fixed points are 0 and, through the `0xff` clamp, a saturated channel. **A consumer that composites shadows into a coverage mask and applies the darkening once draws something the original never draws**, and the difference is exactly where two casters' silhouettes cross

**Confidence.** High (the absence of a guard is the complete pixel loop of the one routine, and the two lookup forms are shown equal from their own instructions) / **Medium** (that the doubled step is *perceptible* rather than merely computed — that depends on the runtime shroud level and the palette, and the falsifier is a running original)

**Original status.** ● active

**Evidence.** [EXP-0130](../experiments/EXP-0130-shadow-bounds/)

### TERR-SPR-144

**`[L06416]` is the engine's shadow switch, and it reaches nothing else: seven reads image-wide, all seven inside the map-view painter, each one guarding a shadow dispatch.** `EnumRefs refto:L06416` gives 14 hits over 5 owners with **0 in orphan or undisassembled code**. The seven reads are all comparisons of the dword at `[L06416]` with zero in `R0379`, and each zero branch lands past exactly one shadow draw: `L10889` skips sweep 1 entirely (the branch to `L10890`, the whole vtable-slot-`+0x2c` loop), `L10891` skips the call through slot `+0x2c` at `L02983`, `L10892` skips the call through slot `+0x2c` at `L02985`, `L10893` skips the call through slot `+0x2c` at `L02988`, `L10833` skips the call through slot `+0x3c` at `L10274`, `L10894` skips the call through slot `+0x3c` at `L10302`, and `L10895` skips sweep 5 entirely (the branch to `L10896`). **No body dispatch, no terrain call and no shroud call is under any of them**, and the seven cover every shadow draw the ten sweeps make (`TERR-SPR-137`). The four writers are all in `R0326` (`L10897`, `L10898`, `L10899`, `L10900`) and the address itself is handed to `R1226` and `R1217` and stored at `+0x28` of a record by `R1227` — the shape of an options binding, the same shape `TERR-LIGHT-108` found behind `ShowTimeFlow`. **For G2 this is the cheapest seam on the render path**: one dword removes every silhouette and changes nothing else, and its class is a runtime global — no shipped file's bytes move. What is **not** established: what sets it, and therefore what it is called

**Confidence.** High (the reference enumeration states its instrument and orphan count; all seven guards and their branch targets are named instructions) / **Medium** (that `R0326` is an options reader — the writer set and the two address hand-offs point that way and no instruction says so)

**Original status.** ● active

**Evidence.** [EXP-0130](../experiments/EXP-0130-shadow-bounds/)

### TERR-FOG-145

**Tile bit 15 survives a save and a load, and its second writer is the save-load path.** `TERR-FOG-087` concluded that a save does not carry the fog and that loading restores a fully unexplored map, resting the persistence half on three arguments, one of which was "no writer of bit 15 exists outside the stamp". There is a second writer: `L07966` (a 16-bit OR into the tile word) in `R0099`, the save-load path, where the value ORed in is `0x8000` or `0` and the pointer walks the tile plane one `u16` per cell. It is fed from the save file's own uncompressed tail: the call to `R0452` at `L08361` reads `Fog/FirstState` and the call to `R1061` at `L08362` reads `Fog/Data`, the run-length encoding of the same bit written by `R0084` (`SAV-FOG-061`). **Why the absence argument could not have seen it:** `TERR-TILE-079`'s instrument was an `imm:` sweep, and it was a sweep for bit-15 **clears** (`imm:3fff`, `imm:7fff`), not for writers; the write here has a **register** source and therefore carries no immediate at all — the same blind spot that retracted `TERR-TILE-079`'s own writer-set clause and that `docs/INSTRUMENT.md` rule 4 names. What survives of `TERR-FOG-087` intact: the shipped-map census (**EN 38 maps / 880 704 cells, RU 34 / 677 696; bit 14 on 0 cells, bit 15 on exactly 1 cell per map and that cell is index 3 on all 72**, the section head's `f32` overlay), and the conclusion that no shipped map authors either bit. That census is what makes the load arm's OR-only write correct: the plane the `.alm` supplies has bit 15 clear everywhere, so ORing the record in reproduces the saved state exactly. **Consequence for a consumer: explored terrain must be persisted across a save/load, and bit 14 must not be** — the save masks `0x8000` alone, and bit 14 is re-derived within 32 ticks by the map-wide clear at `L01659` and the stamp (`ANIM-TICK-011`)

**Confidence.** **High** (the writer is a named instruction in a routine read end to end, and the corpus discriminates: both restart saves carry one run and 0 set cells, two saves minutes apart in one session carry 152 and 248, and the longest carries 3475 of 20 736) / **Medium** that bit 15 is the only tile bit a save restores: only the `Fog` section was traced into the plane, and no image-wide enumeration of tile-plane writers was run

**Original status.** ● active

**Evidence.** [EXP-0150](../experiments/EXP-0150-save-fog-record/)

### TERR-CELLREC-146

**What keys the cell record's four occupant slots, and what each of the six area-layer slots holds.** The four-slot payload layout is **not new**: `formats/terrain/format.md` has carried it since EXP-0070, down to `+0x10` being written by `R0938`/`R0447` with no plane arm reading it, and EXP-0085 sourced the six-slot cost loop; that spec's own open item is *"what each of the six holds is still open"*, and its `+0x4`/`+0x8` are labelled *ground occupant* / *air occupant* with nothing said about what decides which. This row answers both and identifies `+0x10`. **The six layer slots are `payload+0x14 + 4*idx` with `idx = R1074(spellId)`**, giving spell 3 to `+0x14`, 7 to `+0x18`, 8 to `+0x1c`, **19 to `+0x20`**, 12 to `+0x24` and 17 to `+0x28` (`MAGIC-WALLBLOCK-045`; `tools/areamove` derives the mapping from the image's own byte and jump tables). The record's payload is 52 bytes, copied into the scratch at `map+0x5402c` by a repeated dword copy of `0xd` dwords from `node+0xc` in all 28 routines that touch it (`EnumRefs disp:5402c`, 28 hits / 28 owners / 0 orphan). Layout as the instructions use it: `+0x00` the cost byte saved when the record was created, `+0x01` the static block byte saved then, `+0x02` the count of occupied area-layer slots, `+0x04` an actor whose movement domain is 1 or 2, `+0x08` an actor whose movement domain is 3, `+0x0c` a structure, `+0x10` a sack, `+0x14`…`+0x28` the six area-layer slots (`MAGIC-MAPLAYER-040`), `+0x2c`…`+0x2f` four bytes read by `R0458`'s trigger arm. The domain selects the slot through the same getter in both directions: `R1851` loads the vtable and jumps through the slot at `+0x20`, a tail jump to `MOVE-DOM-024`'s domain getter; `R0458` calls it at `L10901` and branches at `L10902` when not above (domain 0 or below cannot occupy), compares the byte at frame local `-0x30` with `0x2` at `L10903` (domains 1 and 2 take `+0x4`, stored at `L07121`) and with `0x3` at `L10904` (domain 3 takes `+0x8`, stored at `L07122`); the removal `R0057` calls `vt+0x20` at `L07106` and clears `+0x8` at `L07112` and `+0x4` at `L07113`. **Both actor slots hold at most one actor** — `L02040` (a comparison of the dword at `+0x4` with zero) and `L02041` (a comparison of the dword at `+0x8` with zero) return 0 without storing when taken. Slot `+0xc` is written only by `R1357` (`L10905` tests, `L07898` stores). Slot `+0x10` is written only by `R0938` (`L07983` tests, `L07984` stores), whose one caller `R0942` is the sack registration `ITEM-SACK-010` already identifies. Accessors, by the displacement each returns: `+0x4` — `R0035`, `R1071`, `R1826`; `+0x8` — `R1072`; `+0xc` — `R1073`; `+0x10` — `R0955`, `R0446`, `R0143`. What each slot does to the planes is `R0453`'s: `+0x4` sets dynamic bit 6, `+0x8` dynamic bit 7, `+0xc` sets or clears bits 0 and 2 per the structure's own mask, `+0x10` is never read there and has **no** effect on either plane

**Confidence.** High (the domain branch is the same three-way test in the add and the remove, read off two full listings; the slot-to-accessor mapping is the returned displacement in each listing; the layer mapping is read out of the shipped image's own two tables; `tools/areamove` re-asserts the branch, both slot stores and both accessor returns as literal byte strings, 0 mismatches on both roots) / Medium for the **name** *sack* on `+0x10`, which rests on `ITEM-SACK-010`'s independent identification of `R0938`'s caller rather than on anything in this round

**Original status.** ● active

**Evidence.** [EXP-0174](../experiments/EXP-0174-area-movement/)

### TERR-FOOTPRINT-147

**Successful entry stores an actor's pointer in every cell record of its `n × n` footprint; refusal stops iteration without local rollback.** The structure half of this is already promoted (`TERR-STRUCT-076`, `formats/terrain/format.md`); the actor half is one of the three items `formats/move/format.md` lists as *not specified* — *"what `R0058`, `R0050` and `R0057` do on a cell boundary"* — and is what this row adds. `R0050(map, actor)` is the only caller of `R0458` (`EnumRefs callto:R0458`, 1 hit / 1 owner / 0 orphan). It reads the footprint side once at `L05505` (a call through the vtable slot at `+0x1c`) into the frame local at `-0x4`, then runs two nested loops both bounded by it (`L05506` and `L05507` compare the counters with the frame local at `-0x4`) and calls `R0458` once per covered cell at `L05508`, pushing `y0 + outer` (`L07079` adds the frame local at `-0x10`) and `x0 + inner` (`L07080` adds the frame local at `-0x14`). A refusal from any covered cell stops further iteration and returns false: the non-zero branch at `L07116` falls through to `L10906`, which zeroes the result. `R0050` has five callers (`EnumRefs callto:R0050`, 5 hits / 5 owners / 0 orphan), among them `R0047`, the sub-cell step of `MOVE-STEP-010`, and `R0039`, the arrival of `MOVE-REFRESH-012`. Structures use the same shape over a **rectangle**: `R0488(map, obj)` loops rows `0..obj+0x61` and columns `0..obj+0x60`, keeps a running row-major sub-cell index, and calls `R1357` for each sub-cell whose bit is set in `obj+0x68` (`L10480` tests the bit in `obj+0x68`, `L10907` calls `R1357`); a refusal aborts at `L10467`. `R0453` computes the per-cell bit for the **solid** mask `obj+0x64` with the same numbering and the same stride `obj+0x60` (`L10908`…`L07902`), which is what makes a structure both present in and selectively passable over its own rectangle. Consequence for `R0453`: dynamic bits 6 and 7 are rebuilt per cell from the record's own slot, not from a separate footprint walk

**Confidence.** High (both loops with their bounds, both refusal exits and both masks are named instructions in full listings; the two mask usages are cross-checked against each other by the identical stride and numbering; `tools/areamove` asserts the loop bounds, the per-cell call and both mask tests as byte strings) Earlier successful cell writes and mover caches are not undone (`SAV-CELLFAIL-583`); the former all-or-nothing wording is withdrawn in `claims/retracted.md`.

**Original status.** ● active (amended)

**Evidence.** [EXP-0174](../experiments/EXP-0174-area-movement/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### TERR-PASS-148

**Bit 4's origin, located: the area module writes it, one serializer reads it, and no mover mask contains it.** `TERR-PASS-073` recorded bit 4 as a live runtime flag with no located origin, carried across every recompute. It is written by `R1080(map, x, y)`, four instructions that OR `0x10` into `map+0x10000` at `L05456` and `map+0x20000` at `L05457` for one cell. `EnumRefs callto:R1080` returns 2 hits in 2 owners: `L05458` inside `R1079` and `L05459` inside `R0630`. At `L05458` the call is gated by `L10909` (a comparison of the argument at frame offset `+0x10` with zero), the per-cell add's fourth argument; the cloud painter computes that argument at `L10910`…`L10911` as `innerEffect->IsKindOf(L08689) && innerEffect+0x5c > 0`, `+0x5b`/`+0x5c` being the byte fields `fire_ball`'s divide operates on. **`wall_of_earth` never reaches it**: its arm at `L10912`…`L05455` returns before that block (`MAGIC-WALLBLOCK-045`). Reader side: no mask value in `R0471`'s complete write population contains bit 4, and in `EnumRefs disp:10000` (78 hits / 33 owners) plus `disp:20000` (60 / 19), 0 orphan, the only routines that read the bit are the three that preserve it across a recompute (`R0453` at `L10472`, `R0057` at `L10913`, `R1358`) and **`R1655`**, which run-length encodes it over the map interior into a `CArchive` (`L10914`, a mask with `0x10`; first cell `map+0x10808`, bounds `map+0x58ee2` and `map+0x58ee3`; `EnumRefs callto:R1655` 1 hit, `R0131`). So bit 4 marks a cell carrying a damaging area effect, it is transmitted, and it is never consulted by movement. **Bit 5 is reconciled at the same time**: `MAGIC-WALLEARTH-042` calls it the bit meaning *this cell has a spell record*; three record creators outside the area module set it (`R0938` at `L10494`, `R1357`, `R1351`) and every creator funnels into `R0453`, which ORs it unconditionally at `L10495` and copies the byte to the dynamic plane at `L10528`. Its meaning is `TERR-PASS-073`'s — the cell has a record — and eleven routines gate a hash lookup on it

**Confidence.** High for the origin and the gate: the writer is four instructions, the enumeration is reported with hit, owner and orphan counts, and the gating expression is read at instruction level / Medium for the **name** *damaging area effect*: the gate names the class at `L08689` and the field `+0x5c`, but what `R0131` does with the encoded plane was not read, so the flag's purpose on the receiving side is not established

**Original status.** ● active

**Evidence.** [EXP-0174](../experiments/EXP-0174-area-movement/)

### TERR-LIGHT-149

**The `TERR-LIGHT-015`/`TERR-LIGHT-023` Unknown, answered field by field: over the searched population no reachable read observes the stored value of any of the four light-side scalars — payload `+0x08`, `+0x0c`, `+0x10`, `+0x14` — while the seventh read of the same case-0 sequence is consumed.** Both rows leave "whether any other path reads the stored values" open. Each file scalar has two destinations — the map object `M` (`ALM-META-091`) and the landscape object `P` (`TERR-LIGHT-023`) — and the three-way verdict differs between them. **Payload `+0x08` (θ) → `M+0x18`, `P+0x20`:** never read on `M`; on `P`, *overwritten before its first read*, because `L10672` stores `[L07948]` there on `R0468`'s entry block and the load path forces that routine (`TERR-LIGHT-151`). **`+0x0c` → `M+0x1c`, `P+0x2c`:** *never read*, on either — zero accesses of any kind at `P+0x2c` in the whole searched population (`TERR-LIGHT-150`). **`+0x10` → `M+0x20`, `P+0x1c`:** never read on `M`; on `P` its one read is `L10915`, inside `R0468` and after that routine's own store at `L10916`. **`+0x14` → `M+0x24`, `P+0x1d`:** the same shape, store `L10917`, read `L10918`; and the file cannot deliver this one intact in any case, because case 0's `read(P+0x1d, 4)` writes four bytes into a one-byte slot, so its bytes 1..3 land on `P+0x1e/0x1f/0x20`, overwriting the low byte of the θ slot with `+0x14`'s byte 3 — which is 0 on all 72 shipped maps. **`+0x18` → `P+0x28`:** *consumed* — it is the terrain tile-group mask (`TERR-LOAD-152`). That last field is what makes the four negatives a result about these fields rather than about the instrument: the seventh scalar of the same seven-read sequence has a reader four instructions after the object is published, and the same search finds it

**Confidence.** **Medium** (three of the four negatives are bounded by the reachability search's blind spots, enumerated in `ALM-META-092` and `TERR-LIGHT-150`; the fourth, the `+0x14` spill, is read straight off the `lea` displacements at `L10919`/`L10920` and the constant 4 pushed before each call)

**Original status.** ● active

**Evidence.** [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

### TERR-LIGHT-150

**`P+0x2c`, the landscape slot payload `+0x0c` is read into, has no access of any kind in the searched population: the builder's own store is the only site that names it.** `TERR-LIGHT-015` and `TERR-LIGHT-023` give the destination and call it unread by lighting without bounding the search. Interprocedural pointer reachability over a deliberately over-approximating seed set — **every** load of a dword from `base+0x80` into a 32-bit register in `.text`, 262 seeds, with no test that the base is a landscape view, so every unrelated class holding a dword at `+0x80` is in the population — entered 305 bodies and recorded 444 uses. Inside the window `P+0x18..P+0x2f` it found 42 accesses and **0 at `P+0x2c`**: no read, no write, no push of the address. The builder's `L10921` (forming the address `base+0x2c`) feeding the call at `L02393` is the only site anywhere that names the slot. For scale, an every-byte-offset decode of `.text` counts **1256** memory operands at displacement `+0x2c` on some base, so the emptiness is a property of where this pointer reaches, not of a rare displacement

**Confidence.** **Medium** (bounded negative. The seed set over-approximates the population but the propagation under-approximates it, and the evidence names every place the pointer leaves the search: 622 unresolvable `esp` slots, 52 indirect calls reached with the pointer live, 16 callees without an `ebp` frame, 13 pointer escapes, 37 stops at unresolved jump tables and 3 at unresolved indirect jumps. A consumer behind any of those would not appear)

**Original status.** ● active

**Evidence.** [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

### TERR-LIGHT-151

**Why "overwritten before read" is the map-load answer and not just one routine's internal order: the load path drives the relight past every branch but one, and the values it copies in cannot come from a file.** (a) *Ordering, computed over the control-flow graphs rather than asserted.* `R0514` publishes the landscape object at `L10922` (a store into `+0x80`), and **every path from that store reaches the call to `R0532` at `L10923` before leaving the body**. Of the four direct calls in between — `L10924`, `R1687`, `L10925`, `R0515`, bodies of 9, 17, 20 and 24 instructions — none reaches `R1820`, the only reader of `P+0x20` outside the relight. The driver is called with argument 1, and `L10926` (a comparison of the argument at frame offset `+8` with zero) is that argument's own test, so the cadence branch at `L10927` falls through. Over that body's graph exactly **one** branch then has a successor from which the call to `R0468` at `L10928` is unreachable: `L10692`, the `S->0x3dc & 1` gate (`TERR-LIGHT-114`). (b) *Provenance of what is copied in.* On `R0468`'s entry block, before its first branch at `L10929`, `[L07948]/[L10240]` go to `P+0x20/0x24` and `[L05650]/[L02293]` to `P+0x1c/0x1d`. Across the whole image those eight addresses take **53 stores, every one inside `[L10237, L10664]`** — inside `R1813`, the sky-light schedule — and **their address is never materialised**: of 61 immediate hits in an every-byte-offset decode, 58 are the displacement fields of those same absolute instructions, the remaining 3 are base+disp decodes starting one byte earlier, and **0** is an address formation, a push of an immediate or a load of an immediate into a register; all 18 aligned dwords in the image holding one of the addresses are the same displacement fields inside `.text`, none in `.rdata` or `.data`

**Confidence.** High for (b) — the live alternative is that the file value reaches these slots by a second route, writing the globals themselves before the relight copies them. Two independent scans exclude it: the only writes are 53 absolute stores in one routine, and nothing anywhere loads one of those addresses into a register, so no register-indirect store, pointer table or callback array can carry a file byte there / **Medium** for (a) — the dominance and reachability are mechanical over the two bodies, but two indirect transfers inside the window are not followed, `L10930` and `L10931`, both calls through the vtable slot at `+0x48`, the second posting message `0x403`

**Original status.** ● active

**Evidence.** [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

### TERR-LOAD-152

**`R0516`'s argument is the map's own tile-group mask, and it is the seventh scalar of the landscape builder's case-0 read sequence.** `TERR-LOAD-002` reads the tile loader without its input. `ALM-RDR-059` says `R0514` hands the object to both `R0515` and `R0516`; **only the first is true** — `R0515` is passed the object (`L10932` loads `+0x80`, `L10933` pushes it), the tile loader is passed a *value*: `L10934` (loads the object from the frame local at `-0x4c`), `L10935` (reads `+0x80` of it), `L02163` (reads `+0x28` of that), `L02164` (pushes it), `L10936` (calls `R0516`) and `L10937` (releases `4` bytes of arguments), i.e. `P+0x28`, which case 0 of `R0480` fills from payload `+0x18` (`TERR-LIGHT-023`). Its prologue — a `0x538`-byte frame reservation plus seven pushes plus the return address — puts that argument at stack offset `0x558`, and the body reads it there as a bit mask: the stack slot at offset `0x14` walks from `L10218` in steps of `0x10` while below `L10219`, exactly **32** iterations, with the stack slot at offset `0x10` as the counter `i`, and iteration `i` is entered only when `mask & (1<<i)` — `L10938` (loads 1), `L10939` (shifts it left by `i`), `L10940` (loads the mask from stack offset `0x558`), `L10941` (tests the mask against the shifted bit) and `L10942` (branches to `L10943` when the bit is clear). Each entered iteration composes one path from `(i>>2)+1` and `(i&3)*4` and fills four consecutive `tiles[]` slots. **One bit therefore selects one group of four tile variants: bit `i` → group `(i>>2)+1`, variants `(i&3)*4 .. +3`.** Since the loader's own index is `(G-1)*16 + V` (`TERR-LOAD-002`), iteration `i` fills exactly `tiles[4i .. 4i+3]`, which is the block the four render routines reach with strip group `g = i` and blend column `b = 0..3` (`TERR-IDX-003`). So one mask bit is one value of `g`: a clear bit leaves all four images that tile-word group can select null. Bits 32 and above are not examined. Corpus, 72 maps on both preserved roots: eight distinct masks — `0x1fff`×44, `0x0fff`×10, `0x00ff`×4, `0x0fbf`×4, `0x1fdf`×4, `0x0f7f`×2, `0x1f1f`×2, `0x1fbf`×2 — never a bit above 12, which decodes to groups 1–3 entire plus group 4 variants 0–3: exactly the shipped file set, whose gaps `TERR-LOAD-002` records as groups 5–8 and `tile4` variants ≥ 4

**Confidence.** High (the alternatives the question leaves live are that `P+0x28` is unread and that the argument is something other than a per-slot selector — a count, a quality level, a version. The loop excludes both: the argument is tested `mask & (1<<i)` where `i` is the same counter that indexes `tiles[]` and forms the filename's group and variant numbers, so no other reading of it survives the instruction sequence. The corpus agreement checks the decoding; it is not the evidence for it) / **Medium** for the *name* tile-group mask — the loop fixes the arithmetic, and that an author chose the bits to pick terrain families is a reading of the shipped value set

**Original status.** ● active

**Evidence.** [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-STREAM-157 | The metadata+0x0c read reaches M+0x1c inR0478 and P+0x2c inR0480 through the same constructed wrapper class described by ALM-STREAM-100. | High | ● active | [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/) |

### TERR-STREAM-157

The metadata+0x0c read reaches M+0x1c inR0478 and P+0x2c inR0480 through the same constructed wrapper class described by ALM-STREAM-100. Twelve original-x86 cases run those two read prefixes over six scalar bit patterns and preserve each exact word with a complete synthetic I/O cut. The wrapper itself does not numerically interpret or read back that destination. Successful-short/EOF cuts show why full transfer needs a complete lower read. This resolves only the selected stream edge of ALM-META-092/TERR-LIGHT-150; archive container+34, later M/P publication, aliases, first use and scalar meaning remain Unknown.

**Confidence.** High for conditional local storage and supplied-memory execution; later lifetime and meaning Unknown

**Original status.** ● active

**Evidence.** [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/)

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-TILECONTROL-163 | Original R0278 including R0469, R0470 and the final 65,536-byte block-plane copy ran on 136 word controls plus one object control, in a synthetic 32x32 map at interior (16,16), with fixed constructor tables, synthetic cost ... | High | ✔ promoted | [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |

### TERR-TILECONTROL-163

Original R0278 including R0469, R0470 and the final 65,536-byte block-plane copy ran on 136 word controls plus one object control, in a synthetic 32x32 map at interior (16,16), with fixed constructor tables, synthetic cost vector [255,8,8,9,14,6,12,11,16,8,6], height 7 and no structures/occupants. Word 0214 returns class 1/cost 8 but raw water still writes block 1; word 03ff returns invalid class/cost 255 but block 0. Toggling bits 10..12 or 14..15 changes none of these selected cost/block values; bit 13 on 0000 writes block 1. A nonzero type 3 cell replaces block 1 with 5; the x=7 border control is 31. Static/dynamic copies agree. The actual static/dynamic TEST consumers, with a synthetic 1x1 footprint, return pass iff block AND mask is zero for masks 0x41/0x44/0x82 over six independently supplied block bytes. This proves the bounded predicate, not native movement or later cell-record state.

**Confidence.** High for selected instruction execution and independent cost/block/object/border controls

**Original status.** ✔ promoted

**Evidence.** [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-WATERBOUND-164 | The water classifier tests subcell < 8 at L09155/L09156 before the class 9/class 1 arm. | High | ✔ promoted | [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |

### TERR-WATERBOUND-164

The water classifier tests subcell < 8 at L09155/L09156 before the class 9/class 1 arm. Across all 256 water inputs in the 1024-input classifier domain, under each of two fixed synthetic cost tables, 128 subcells 8..15 return byte 0xff, 124 return 9, and 4 return 1; every water case leaves the classifier cost output at 8. Ingest separately changes cost to slot 0 (constructor default 0xff) for return >= 13 and separately sets block 1 from the raw water predicate. The subcell rejection already appeared in the terrain reference; the unconditional water-return prose in TERR-PASS-050 is partially retracted. Non-water groups 13..15 remain allocation-dependent where their pair is read: synthetic pair residues 1 and 10 change the selected class. No terrain meaning is assigned to that residue.

**Confidence.** High for original unsigned guard, complete finite classifier controls and residue rival; native allocation state Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-GFXBOUND-165 | Full render kernels R0539, R1815 and R1816 extract g=(w & 0x1fff)>>6, b=(w>>4) & 3 and sub=w & 15. | High | ✔ promoted | [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |

### TERR-GFXBOUND-165

Full render kernels R0539, R1815 and R1816 extract g=(w & 0x1fff)>>6, b=(w>>4) & 3 and sub=w & 15. Bits 6..12 are in g; the former 6..9 parenthetical of TERR-IDX-003 is partially retracted, its arithmetic retained. Each kernel ran 136 one-bit controls; R0539 additionally ran 128 groups times 4 variants. With animation disabled, water groups 8..11 select group 8. For non-water 0000 toggles, bit 10 reaches slot 64, bit 11 reaches 128, and bit 12 reaches 256. Groups >= 32 attempt addresses outside the loader-owned 128-pointer array at L10218..L10944; execution stops before those reads. A separate original first-pixel-buffer load at L10945 faults at address 0x10 for a null selected pointer and reads a supplied pointer for a synthetic bitmap. This is neither a graceful rejection nor a native crash/pixel witness. A valid loaded resource for off-corpus groups remains Unknown.

**Confidence.** High for named arithmetic and array-boundary discrimination; selected field-read fault is emulator-only; actual resources/pixels Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-SLOTADMIT-166 | The original 3d loader loop L10946..L10947 covers 32 metadata-mask bits, four slots each. | High | ✔ promoted | [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |

### TERR-SLOTADMIT-166

The original 3d loader loop L10946..L10947 covers 32 metadata-mask bits, four slots each. All 128 initial slots were synthetic non-null objects. Masks zero, all-one and each one-bit mask produce exactly the corresponding 4i..4i+3 slots under successful allocation/bitmap/destruction service cuts. Masked-out non-null slots are destructed and zeroed at L10948/L10949; enabling a mask bit requests variants, not proof of usable bitmap pixels. Formatter calls retain group (i>>2)+1 and variant (i AND 3)*4+j. The original caller L10934 loads P+0x28 and passes that value at L10936. This agrees with TERR-LOAD-152 and separates admission, stale slot clearing and the renderer's larger index domain.

**Confidence.** High for loop, arguments and 34 controlled masks; resource constructor/lifetime implementation and native load outcomes excluded

**Original status.** ✔ promoted

**Evidence.** [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-WATERPASS-167 | Renderer R1814 shares the word extraction but has a caller-specific non-water skip: L10950 jumps to L10951, then L10952, before its lookup L10953. | High | ✔ promoted | [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |

### TERR-WATERPASS-167

Renderer R1814 shares the word extraction but has a caller-specific non-water skip: L10950 jumps to L10951, then L10952, before its lookup L10953. Under the 136 fixed-word controls with animation disabled, 29 water cases reach a synthetic populated slot and 107 non-water cases skip. A bit that changes g from water to non-water can suppress this pass without demonstrating that another terrain renderer omits the cell. The three full kernels are separately scoped by TERR-GFXBOUND-165; no new universal renderer equivalence follows.

**Confidence.** High for original branch and 136 fixed-state controls; scheduling and visible repaint effect Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-RLEBIT-168 | Render-grid updater R1656 selects a zero/nonzero initial packet word as flag 0/0x2000, walks alternating run lengths over the interior beginning (8,8), and its write kernel L10954..L10955 computes (w & 0xdfff) \| flag. | High | ✔ promoted | [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |

### TERR-RLEBIT-168

Render-grid updater R1656 selects a zero/nonzero initial packet word as flag 0/0x2000, walks alternating run lengths over the interior beginning (8,8), and its write kernel L10954..L10955 computes (w & 0xdfff) | flag. Eight original-kernel controls over seeds 0000/ffff/4000/8000 and flags 0/0x2000 preserve all other bits. The pointer resolves through [[view+0x80]+0x0c], not a simulation byte plane. Only the write kernel is executed; packet removal L10214, run producer/dispatch, first update and later ordering remain Unknown in this experiment.

**Confidence.** High for inspected selected decoder and executed write kernel; packet/lifetime integration excluded

**Original status.** ✔ promoted

**Evidence.** [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-TILECLEAR-169 | The original R0374 prefix obtains [[view+0x80]+0x0c], sets view+0xdc/+0xe0, and runs AND word, 0xbfff over view.width times view.height words. | High | ✔ promoted | [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |

### TERR-TILECLEAR-169

The original R0374 prefix obtains [[view+0x80]+0x0c], sets view+0xdc/+0xe0, and runs AND word, 0xbfff over view.width times view.height words. All 136 supplied-word cases clear bit 14 only. For high-pair seeds 00/01/10/11, the immediate post-clear pairs are 00/00/10/10. Execution stops at L10956 before drawable enumeration/restamping, so these are immediate-clear states, not the final state of the whole event or a measured native period.

**Confidence.** High for prefix, exact width and 136 one-bit controls; native cadence and post-event state Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-DRAWSTAMP-170 | Both original vtable data ranges L02468 and L02585 bind +0x44/+0x48 to R1402/R0554 and +0x20/+0x24 to R0586/L10957. | High | ✔ promoted | [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |

### TERR-DRAWSTAMP-170

Both original vtable data ranges L02468 and L02585 bind +0x44/+0x48 to R1402/R0554 and +0x20/+0x24 to R0586/L10957. The latter getters read drawable class+0xd0 through L02113, separately from simulation actor virtual +0x1c. With that original dispatch and a synthetic class size 1, the stamp reaches the four high-byte stores that OR in 0xc0 only for decay < 3, permission bit 3 and nonzero sight. A fixed synthetic LOS service admits one cell; the stamp writes both bits at its four corners. The per-tick wrapper additionally requires both low-five coordinate bits 16 and a changed saved stamp position: changed-x/changed-y controls stamp, stationary/uncentered/decay 3/permission 0/sight 0 controls do not. The LOS builder R0284 is cut; no native visibility extent, timing or globally complete writer set is claimed.

**Confidence.** High for original dispatch, independent layout, four seed stores and 12 guard controls; LOS service and native schedule Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-DRAWGATE-171 | Drawable R1504, with permission bit 3 clear and original dimension getters over a synthetic 1x1 footprint, ORs the four light-grid words before masking with 0xc000. | High | ✔ promoted | [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |
| TERR-METALIFE-175 | Landscape loader R0480 constructs a distinct 0x30-byte P with vtable L10958. | High | ✔ promoted | [EXP-0357](../experiments/EXP-0357-metadata-scalar-lifetimes/EXP-0357.md); generated receiver table, checks and vectors |

### TERR-DRAWGATE-171

Drawable R1504, with permission bit 3 clear and original dimension getters over a synthetic 1x1 footprint, ORs the four light-grid words before masking with 0xc000. All 256 independent high-pair corner combinations yield 225 state 0, 15 state 1 and 16 state 2: combined 0xc000 -> 0, 0x8000 -> 1, 0x0000/0x4000 -> 2. A corner 0x8000 plus another 0x4000 reaches 0 even when no individual corner is 0xc000; therefore OR-before-compare is not an all-corners-current test or an any-single-0xc000 test. Raw 0x4000 survives the light-parser bit 13 clear, although the known stamp writes 0xc000 and periodic clear removes 0x4000. The gate is executed only through its computed state at L10959; downstream redraw/native visibility is excluded.

**Confidence.** High for original getter/gate execution and all 256 discriminating corner controls; downstream/native presentation Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

### TERR-METALIFE-175

Landscape loader R0480 constructs a distinct 0x30-byte P with vtable L10958. Publication L10922 and binder R0515 store P at view+0x80; the binder copies P+4/+8 into dimensions and does not copy the P header. On replacement,L10960 dispatches old P through slot+4, bound by L10961 to L10962 and member destructor L10963. The latter reads/frees arrays at P+0x0c/+0x10/+0x14/+0x18; the wrapper optionally frees P. Twelve original-x86 cases preserve six words/two fills at P+0x2c through binding and observe no scalar read or write in the executed binder/destructor instructions before their free cuts poison memory. The view still holds the address at that stop. System free receives P and remains a boundary; post-publication message L10931 uses session+0xd4 slot+0x48, whose receiver/native delivery was not bound.

**Confidence.** High for conditional local binder/destructor instructions and original class dispatch; other P aliases, native LOAD/first use and semantic role Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0357](../experiments/EXP-0357-metadata-scalar-lifetimes/EXP-0357.md); generated receiver table, checks and vectors

## Terrain claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-FAMILY-187 | Selector mask 0x02 enables optional 3d precomposition before ordinary terrain loading. | High / Medium | ✔ promoted | [EXP-0370](../experiments/EXP-0370-terrain-family/) |
| TERR-FAMILY-188 | The two terrain families have equal paired geometry but different palette and index content. | Medium | ✔ promoted | [EXP-0370](../experiments/EXP-0370-terrain-family/) |

### TERR-FAMILY-187

**Selector mask 0x02 enables optional 3d precomposition before ordinary terrain loading.** In `R0516`, clear `L10964 & 0x02` enters ordinary loading at `L10965`; set enters `terrain.3d`, and a nonzero map group mask admits precomposition. Graphics object `+b84 == 8` selects the palette/index arm. Both normal surface arms delete/zero the temporary 128 tile pointers and dirt before falling through to the ordinary family. The final table is from that ordinary pass; the independent surface table `L10221` survives source cleanup and is passed by the map caller to `L10966` under mask 0x02. The normal startup body writes EBX=0 at `L10967` before `R1691`, assuming its intervening callees preserve nonvolatile EBX. All 12 occurrences in the complete file-backed embedded-address scan for L10968..L10969 resolve to absolute selector references. The only other direct write is `OR 1@L10970`, behind mask 0x02 and object `+12b0 != 0`; it cannot enable mask 0x02. The selected object initialization chain reaches an explicitly named IDirect3D creation operation. Its `-systemmemory`/`-emulation` members and the separate `-safevideo` global do not set this selector. This establishes neither hardware-only mode nor native activation. No enabling writer is found within the stated address-form population; computed pointers, bulk copies and native lifecycle remain Unknown.

**Confidence.** High for local branches, source cleanup and named writes: 25 decoded routines/funclets, 9,461 instructions, 45 detached original vectors including a rejecting branch mutation. Medium for the bounded producer search. Full startup, precomposition and native graphics execution were not performed.

**Original status.** ✔ promoted

**Evidence.** [EXP-0370](../experiments/EXP-0370-terrain-family/)

### TERR-FAMILY-188

**The two terrain families have equal paired geometry but different palette and index content.** In each pinned EN/RU graphics archive, both exact prefixes contain 53 BMPs with identical names/geometry: dirt 32×128, 36 strips 32×448 and 16 water strips 32×256. All use uncompressed 8-bpp, a 40-byte header, 256 palette entries and pixel offset 1078. The 52 tile bfSize fields understate resource length by 1024; each dirt has two bytes after its geometric pixels. Every corresponding EN/RU file is byte-identical within its family. Between families, all 53 file/palette pairs differ, including 255 RGB palette entries per pair; 647,325 of 651,264 logical indices and 647,586 RGB samples differ. All 3d files share one palette; 52 ordinary tiles share another and ordinary dirt differs from them in 41 RGB entries with reserved bytes equal. For 3d minus ordinary source RGB, ranges are R -69..113, G -58..99, B -58..164, so neither palette-only/reindexing-only equivalence nor uniform channel increase fits. Native rendered brightness and renderer correspondence are not inferred.

**Confidence.** Medium for this complete pinned corpus. The custom parser and independent Pillow decode agree on all 212 logical index/RGB planes; synthetic controls discriminate unused palette changes, reindexing, used-colour changes and signed deltas.

**Original status.** ✔ promoted

**Evidence.** [EXP-0370](../experiments/EXP-0370-terrain-family/)

## Unit shadow arms

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-191 | The unit shadow's first main-pair silhouette reads `[L04368]`, or `[L04335]` while the unit holds effect key `0x26` and its owner row has bit 3; the second silhouette always reads `[L04335]`. | High | ● active | [EXP-0449](../experiments/EXP-0449-shadow-arms/) |
| TERR-192 | The hero sheet pair `+0x194`/`+0x198` replaces the main pair when drawable `+0x18c` bit 0 is set and always shears; two silhouettes recolour overlapping pixels twice, and shipped sheared pairs overlap only at unequal frame sizes. | High / Medium | ● active (amended) | [EXP-0449](../experiments/EXP-0449-shadow-arms/) |
| TERR-193 | In `R0553` only the drawable's runtime class picks the flat arm: a `CAirUnit` instance draws both main-pair silhouettes through `vt+0x1c`; every other unit, and the hero pair, uses `vt+0x3c`. | High / Medium | ● active (amended) | [EXP-0449](../experiments/EXP-0449-shadow-arms/) |
| TERR-194 | Of 68 displacement stores through `+0x18c` in `rom.exe`, 45 reach a drawable and 5 set bit 0; 23 reach other classes; the dispatcher holds six, not four; three more, via an address formed by adding 0x18c to a base, preserve bit 0. | High / Medium | ● active | [EXP-0458](../experiments/EXP-0458-shadow-writers/) |
| TERR-195 | A `Sonic Bat` or `Dragon` shadow translates each sheet by its own size and the two flat silhouettes do not overlap at the routine's placement in any frame of either sheet; a one-pixel shift would overlap them. | High / Medium | ● active | [EXP-0458](../experiments/EXP-0458-shadow-writers/) |
| TERR-196 | All eight blit cores, flat and sheared, read the destination pixel and write its recolour back through two writes with no coverage buffer, so a pixel stamped by both silhouettes is recoloured twice. | High / Medium | ● active | [EXP-0458](../experiments/EXP-0458-shadow-writers/) |

### TERR-191

**(rom.exe)** `R0553` is `vt+0x2c` of both `CUnit` and `CAirUnit` (`.rdata` slots `L10971` and `L10972`; a raw dword scan for `R0553` finds only those two). All 835 instructions (`R0553`..`L10973`) were read. `L04327 TEST byte [ESI+0x18c],1` sends a clear bit to the main pair at `L10974` (the class's sheet-table entry `+0x04` and `+0x08` through `[L07265]`) and a set bit to the hero pair of `TERR-192`.

The index of each main-pair silhouette is chosen by one lookup, `R0615(0x26)` (`L05798` for the first silhouette, `L05799` for the second): a linear search of the unit's effect list `[this+0x128]` (count `[this+0x12c]`) for an entry whose `dword >> 16` equals the key, returning the index or -1 (`HERO-FIGURE-062`). Kind `0x26` is `invisibility` (`MAGIC-ACTOR-066`).

| silhouette | key `0x26` | extra gate | cell read |
|---|---|---|---|
| first, sheet `+0x04` | absent | none | `[L04368]` (`L10975`, `L10976`) |
| first, sheet `+0x04` | present | owner row bit 3 set, else the routine returns with nothing drawn (`L10977`..`L10978`) | `[L04335]` (`L10979`, `L10980`) |
| second, sheet `+0x08` | absent | `[L04367]` nonzero (`L10981`..`L10982`) | `[L04335]` (`L10983`, `L10984`) |
| second, sheet `+0x08` | present | none: not drawn (`L10985 JGE L04365`) | none |

The owner row is `[[view+0x9b4]+0x38]` indexed by `[[unit+0x14]+4]`; bit 3 marks the local player's own entry on exactly one index (`UNIT-VISBIT-044`, `MAGIC-INVISOWN-085`). An invisible unit's shadow is therefore drawn only on its owner's client. `[L04367]` is the Smoothing option cell: default 1 (`L10986`), flipped by the Ctrl+O arm (`L10987`..`L10988`, `MENU-057`, `TOWN-SMOOTH-460`).

Within this routine the two cells are read at nine sites: `[L04368]` at `L10989`, `L10990`, `L10991` and `[L04335]` at `L10992`, `L10765`, `L10766`, `L10767`, `L10761`, `L10762` (raw dword scan, operand addresses). Their values come from `TERR-LIGHT-126`: `[L04368]` is 4 by day, 6 in twilight and 8 at night, `[L04335]` half of it. So an ordinary unit's first silhouette uses the larger level and its second silhouette the smaller one; the levels that `TERR-LIGHT-126` states for "the unit path" (2 by day, 4 at night) belong to `[L04335]`, which an ordinary first silhouette does not read. This answers `TERR-LIGHT-126`'s open question whether any consumer reads both cells in one frame: one call of this routine does, for a unit without key `0x26`. The reason for the factor of two stays open.

Not read by the selection: terrain height or cell (`TERR-SHDW-133`), the unit's owner beyond the bit 3 gate, the frame, the animation state, and the day phase, which enters only through the two cells' values and the shear `TERR-SHDW-134` computes.

**Confidence.** High. Every push, every load of the two cells and each gate is a named instruction in a routine read end to end; `R0615` was read whole; the alternatives (index follows the sheet, the owner, the terrain, the day phase, an option) each have no reader in the routine, and the option cell gates only the second silhouette's presence. The weakest input is the raw dword scan, which sees absolute operands and not a cell reached through a register-held pointer.

**Unknown.** Whether `R0553` is issued for every drawable (its caller's gates were not read; `ANIM-AIRPASS-086` separates the air sweep). The pixels the original draws were not observed; the level-to-darkness law is `TERR-FOG-084`'s.

### TERR-192

**(rom.exe)** The hero pair is `+0x194`/`+0x198`, the `sprites.256`/`spritesb.256` pair that `R0551` loads from a `units/heroes*` directory (`HERO-APPEAR-043`). It is drawn when the test of bit 1 of the byte at `+0x18c` at `L04327` is set and in place of the main pair: the branch ends at the return at `L10993` or at a jump to the common epilogue `L04365`, and `L10974` (the main pair) has one predecessor, the zero branch at `L10994`. Bit 0 is set by the hero arm of `R0590` (`L03173`, an OR of `9`) and on the character screen's own drawables, and is clear at creation for every map-authored human of the shipped classes (`ANIM-117`). A frame index at or above `[sheet+0x4]` ends the branch with no silhouette (`L10995`, a comparison of the frame index with `[sheet+0x4]` branching when not below); that field's meaning is not identified here.

Both silhouettes go through `vt+0x3c` only (`L04328`, `L04329`); the three `vt+0x1c` calls (`L10996`, `L10997`, `L10998`) lie in the main branch. The index rule is `TERR-191`'s: first silhouette (`+0x194`) `[L04368]` when key `0x26` is absent (`L05797`, `L10758`) and `[L04335]` behind the owner-row bit 3 gate (`L05828`, `L10760`) when present; second silhouette (`+0x198`) skipped when the key is present (`L04337`, `L04338`) or `[L04367]` is zero (`L04366`, `L10999`), else `[L04335]` (`L04336`). One pixel term `T` is computed from the first sheet's drawn frame (`L10815`) and reused by the second (`L11000`); each sheet's own width and height give its anchor (`L11001`..`L04331`).

Overlap. Between the two blits the routine calls only the sheet's size getters and `R0615`. This experiment's listing holds two cores, `L03664` (`vt+0x1c`, flat) and `R1832` (`vt+0x3c`, sheared): each reads the destination pixel, looks it up in `[L03341] + level * [L03340]` and writes the result back, and neither keeps a coverage buffer. The other six cores (flat `L03665`, `L11002`, `L11003`; sheared `R1850`, `R1848`, `R1849`) are not in the listing; that they read the destination and write back rests on `TERR-LIGHT-059` and `TERR-SHDW-135` (the sheared four share one accumulator and recolour idiom), not on this experiment's reading. The hero pair is drawn by the sheared four only. A pixel stamped by both silhouettes is recoloured twice, through the first silhouette's row and then the second's.

Population measured, `graphics.res` on both roots (identical results). The census models the sheared placement only (`TERR-SHDW-129`, `TERR-SHDW-130`); it has no flat arm and no `vt+0x1c` translation. It covers the hero pairs and the non-air main pairs; the `Sonic Bat` (109 frames) and `Dragon` (164 frames) sheets are in the counts below but are drawn through the flat arm of `TERR-193`, so their overlap is unmeasured: 51 `units/**/spritesb.256` nodes, each with a sibling `sprites.256`; 7982 frame indices. Hero pairs: `heroes` 16 and `heroes_l` 14 sheets, 4570 frames. Main pairs: `humans` 7 and `monsters` 14 sheets, 3412 frames. Each frame is decoded with the cores' run grammar (an opcode with bits `0xc0` clear stamps a run, `0x40` skips rows, other opcodes skip pixels; no control byte of `0xc0` or above occurs; every stream is consumed exactly).

- At the routine's own placement with zero shear, 0 of 547,347 pixels of the second sheet fall on a pixel of the first. Moving the second sheet one pixel left or right gives 229,741 and 229,396 (two pixels: 233,838 and 234,009), so the zero is a property of the data and the probe can see overlap.
- 5293 frame indices have equal sizes for the two sheets. They give 0 overlap at 10 effective sun angles (`|theta|` 0.0334 to 0.5236, both signs, `tan` applied), 7 class constants, mirrored or not: both silhouettes then move identically.
- 2689 frame indices have unequal sizes. Their silhouettes are sheared about different pivots, by about `tan * (hB - hA) / 2` pixels relative to each other, and overlap: up to 139 pixels in one frame at one (angle, class constant), and 7.0 to 22.2 percent of the second sheet's pixels summed over the 7 class constants (94,430 to 298,781 of 1,347,080), rising with the angle.

**Confidence.** High for the branch structure, the two exits, the gates and the absence of a flat arm in the hero branch. High for the recolour read-modify-write in the two listed cores; for the six unlisted cores it is the grade of `TERR-LIGHT-059` and `TERR-SHDW-135`, inherited and not re-read here. Medium to High for the equal-size zero and the exact-placement zero: both re-execute the placement formulas with swept class constants, with no run of the original, and the one-pixel shift control shows the probe sees overlap. Medium for the unequal-size overlap counts: they come from re-executing the placement formulas (`TERR-SHDW-129`, `TERR-SHDW-130`) over a swept class constant and angle, not from a run of the original.

**Unknown.** Which frame indices and angles occur in play, so how often the unequal-size overlap is drawn; the class constants of the hero classes were not read (swept instead).

**Amended.** Two Unknowns are answered. The flat-arm pairs were measured: the `Sonic Bat` and `Dragon` sheets do not overlap at the placement `TERR-195` states, and the six cores this card did not list recolour as the two it did (`TERR-196`). The stores that can set bit 0 were classified by receiver: five reach a drawable, none is a store this card's text omitted (`TERR-194`). The census and its sheared-arm counts stand. The sentence that the `Sonic Bat` and `Dragon` overlap is unmeasured is superseded by `TERR-195`.

### TERR-193

**(rom.exe)** The flat-versus-sheared choice in `R0553` is one slot, the stack slot at offset `0x2c`. It is written once, at `L11004`, from `L11005`/`L11006`/`L11007` (push `L07474`, pass `this`, call `R0214`), and read three times, each a zero test that selects `vt+0x1c` for nonzero and `vt+0x3c` for zero: the first silhouette at `L11008` (key present) and `L11009` (key absent), the second at `L11010`. `R0214` calls the instance's `vt+0x0` (`GetRuntimeClass`) and `R0216` walks the base-class chain (reading each class's `+0x10` base pointer) comparing against the argument: it is `CObject::IsKindOf`. The argument `L07474` is the `CRuntimeClass` record named `CAirUnit` (size `0x1b0`, base `L00620`, the `CUnit` record; identical bytes on both roots).

The arms and cells are therefore: flat with `[L04335]` (`L10996`), flat with `[L04368]` (`L10997`), flat second silhouette with `[L04335]` (`L10998`), and the same three through `vt+0x3c` (`L10790`, `L11011`) for every unit that is not a `CAirUnit`. The flat arm places the silhouette by `TERR-SHDW-132` and recolours through the same table, without the per-row shear.

The runtime class is itself fixed at creation by `units.reg` `Z` (`R0509`, `REG-UNITS-061`), so `Z` is the upstream selector; inside this routine no class-record field is read to choose the arm, only the runtime class. `CAirUnit` instances are the units whose `units.reg` `Z` is nonzero: `Sonic Bat` and `Dragon` among the shipped classes `REG-UNITS-061` measured. No class derives from `CAirUnit`: a raw dword scan of all sections for `L07474` returns three hits, all in `.text` (`L11012` in its `GetRuntimeClass`, `L11013` in a routine not listed or read here, `L11014` in `R0553`), whereas a derived class would hold the address as a base-class pointer in data. The one data hit for `L00620` is `L11015`, `CAirUnit`'s own base pointer.

None of the other candidates reaches the slot. The hero pair has no flat arm (`TERR-192`). The `+0x18c` bit picks the pair, not the arm. Within `R0553` the class record feeds the frame arms, the anchor (`+0x2c`..`+0x38`) and the mirror flag (`+0x108`), and none of those reads reaches a test of the slot. The frame size feeds the anchor and `T` only. The animation state (`[ESI+0x74]`) yields the frame and mirror. The sun enters through `tan(theta)` before the slot is written, and `[L04367]` gates only the second silhouette's presence.

**Confidence.** High that the drawable's runtime class is the only selector inside `R0553` (the `Z` field selects it upstream, at creation): one writer, three readers, `R0214` and `R0216` read whole, the record read from the PE on both roots. Medium for the class population: the `Z`/class identity is `REG-UNITS-061`'s (34 classes on its framing) and no map-authored or scripted creation path other than `R0509` was enumerated here. The no-derived-class clause rests on the raw scan, which would miss a base pointer assembled at run time.

**Unknown.** No runtime observation of either arm.

**Amended.** The creation population of `CAirUnit` is closed by `UNIT-141`: one constructor site, guarded by the hero-id range and the class record's `Z`. The `Medium` for the class population stands for the stale-global edge that claim names. The `L07474` raw scan clause is unchanged.

### TERR-194

**(rom.exe)** Instrument: a capstone sweep of `rom.exe` (SHA-256 recorded in the experiment's input manifest; one binary on both roots, the RU sweep byte-identical) for every instruction whose memory operand has displacement `0x18c` and a write to it, 68 displacement stores in 38 routines (the sweep has 39 starts; `L11016`, a data word inside the dispatcher, is a false start), the population `ANIM-117` counted. Each store was assigned a receiver by hand from its routine's constructor, allocation or vtable store, not by the displacement.

| receiver | stores | bit 0 |
|---|---|---|
| `CUnit` or `CAirUnit` drawable | 45 | 5 set, 5 cleared, 1 copied, 34 preserve it |
| other classes with a member at `+0x18c` | 23 | not the drawable flag word |

- Set. The hero arm of `R0590` for an id in `[0x20,0x40)` (`L11017`, `(old & 0x80) | 9`). The dialogue-speaker synthesizer `R0743` (`L11018`) ORs 1 when the `npc.reg` Flags token `Hero` is true. The character screen: `R0592` stores `0x29` or `0x2b` (`L02788`, `L07287`) and `R0591` ORs 9 (`L11019`).
- Cleared. The constructor (`L11020`), virtual `Init` (`L03487`), the speaker synthesizer's initial `0x48` (`L03489`), `R0590` for an id below `0x1a` (`L11021`) and `R0591`'s zero (`L10191`). The copy constructor `R0595` copies the word; its only caller is the `CAirUnit` copy constructor `R0596`, which has no caller and no vtable slot.
- Preserved. 34 of the 45 AND or OR a mask that excludes bit 0. Three further stores lie outside the 68 and the 34: an add of `0x18c` followed by an OR of `8` into the byte on a drawable from lookup `L11022` (`L11023`, `L11024`, `L11025`); they preserve bit 0, so 48 stores reach a drawable in all.
- The client dispatcher `R0509` holds six stores, all preserving bit 0: OR `0x20`, OR 8 three times, and OR `0x80` and AND `0x7f` at `L03220` and `L03221`. `ANIM-117`'s count of four was short by the last two.
- Other classes (23): a class with vtable `L07691` (3), a member of a class with vtable `L11026` (2), a dword array (3), a vtable `L11027` family (3), server handle objects (6), objects holding a free pointer at `+0x18c` (6). All eight `LEA` hits with displacement `0x18c` are `[ESP+0x18c]`; the five immediate `0x18c` operands are the three adds above, one allocation size and one stack adjustment.

**Confidence.** High for the list of 68 displacement stores and for the five bit-0 setters: each was read with 12 instructions of context. Medium for the receiver classification (45 and 23): each class is a hand annotation, some by pattern across sibling routines, and the 23 are unnamed; the script checks only that every swept site has a row and that the store text matches. Medium for completeness: the sweep sees displacement stores and not `REP MOVS` with a variable count (of 592 `REP MOVS` sites, 4 have a constant count of at least `0x18d` bytes, none `0x1b0`), `memcpy` of a containing structure, or an address computed in a register from a non-immediate. The two `R0743` and `R0592` drawables were not shown to reach `R0553`.

**Unknown.** The class names of the 23 other receivers; whether a variable-count copy carries bit 0 into a drawable.

### TERR-195

**(rom.exe)** `Sonic Bat` (ID 70, node `units/monsters/bat/sprites.256`, canvas 128 by 128) and `Dragon` (ID 71, 160 by 160) are the only `units.reg` classes with `Z` nonzero, on both roots (`REG-UNITS-061`). Their two main-pair silhouettes go through the flat arm (`TERR-193`).

- Placement. `TERR-SHDW-132`'s flat placement is applied to each sheet with that sheet's own width and height: `dstX = x - (floor(w/2) + class term) + shear/2000`, `dstY = y - (floor(h/2) + class term) - z`. The second silhouette's anchor is `L11028`..`L11029` and its call `L10998`; between the two blits the routine calls only the sheet loader, the size getters and `R0615`. Both silhouettes use the one unit position and `z`, so a shared term cancels.
- Overlap. Over 109 `Sonic Bat` and 164 `Dragon` frames, 134 of them with unequal sheet sizes, the second sheet's opaque pixels never fall on the first sheet's pixel at the routine's placement. The first and second sheets hold about 13.8 thousand and 3.9 thousand pixels for `Sonic Bat`, and about 385 thousand and 28 thousand for `Dragon`. A mirrored frame reflects each sheet inside its own box (`EDI` starts at the box's right end, `STD`), and the mirrored overlap at zero shift is also 0.
- Control. Shifting the second sheet one pixel left or right overlaps it on over 1,400 pixels of the `Sonic Bat` sheets and over 11,000 of the `Dragon` sheets, so the probe sees overlap.
- Both roots give identical counts. Frames are decoded with the cores' control-byte grammar; no stream is left unconsumed.

**Confidence.** High that no stamp is double-covered at the placement the formula gives, over this population (two classes, 273 frames, both roots), with the shift control. Medium for the placement itself: it re-executes `TERR-SHDW-132` and the anchors listed and was not observed in a run of the original.

**Unknown.** Which frames occur in play; the shadow's pixels as drawn.

### TERR-196

**(rom.exe)** The flat arm (thunk `R0795`, `vt+0x1c`) calls cores `L03664`, `L03665`, `L11002` and `L11003`. The sheared arm (thunk `R1831`, `vt+0x3c`) calls `R1832`, `R1850`, `R1848` and `R1849`. `TERR-192` listed two of the eight and left six to `TERR-LIGHT-059` and `TERR-SHDW-135`.

- Each core reads its control bytes with the mask `0xc03f`, tests `AH` against `0x40`, and recolours: it reads the destination pixel `[EDI]`, looks it up in `[L03341] + level * [L03340] * 2`, and writes the result back. The index is the pixel shifted right by 3 in some paths and the pixel itself in others.
- Each core has exactly two writes to a destination through the destination pointer, both 16-bit stores (the fast and the clipped path). None writes a byte, repeats a store or addresses a coverage buffer.
- A pixel covered by both silhouettes is therefore recoloured twice. For the flat pairs `TERR-195` finds no such pixel; for the sheared pairs `TERR-192` counts them.

**Confidence.** High that all eight cores recolour in place with two word writes and keep no coverage buffer: each core's whole range was scanned for memory writes. Medium for the recolour's value equivalence between cores (the index shift differs between paths; no table value was compared).

**Unknown.** The pixel values the original draws.

## Selected native cost restore

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-198 | With an original world receiver, the selected L11030 prefix copies found-record baseline+00 to Cost[low16(key)] before an unresolved call; 256 byte controls copy exactly, including zero. | High / Medium / Unknown | ● active | [EXP-0467](../experiments/EXP-0467-cost-write-route/) |

### TERR-198

In the selected native prefix, EBP is entry ECX and EBX is the argument
masked to 16 bits. The inline hash search uses receiver+540b8/+540bc,
(key>>4)%bucketCount and node word+08. Null table, empty bucket and exhausted
key search return 0 before plane writes. A hit copies 13 dwords from node+0c
to receiver+5402c. Load L11031 places scratch+00 in CL; store L07054 writes
that byte to receiver+low16(key). Static baseline+01 is copied separately
at L11032. Count, occupants and current Cost do not gate that local store.

TERR-PASS-049 and TERR-COST-052 identify the original world planes;
TERR-STRUCT-069 and SAV-CELLREC-017 identify the record baselines and scratch.
Those original-derived authorities supply semantics only conditional on a
world receiver. Synthetic memory does not establish the native incoming alias.

The complete selected 472-byte span at L11030..L11033 decodes to 143
instructions. Its preceding RET4/padding, prologue and two balanced RET4
exits support the entry/body layout. Raw relative branches target instruction
starts within it. Target-only E8/E9 .text and exact-entry-dword whole-image
checks found no incoming anchor; they do not prove dead code or an exhaustive
caller census. EN/RU executable bytes are identical, one independent input.

Original bytes executed on 282 fresh synthetic vectors per image. The 256
baseline values 0..255 copy exactly. Current costs 0/10/255, Static baselines,
layer counts, occupied+04, relocated receiver, key/high-word and chain controls
discriminate source, destination and predicate. Each image has 279 found-node
stops before CALL L08162 at L11034 and three zero-return misses with no
non-stack writes. Hit prefixes make 13 scratch dword writes, one Cost byte,
one Static byte and one bucket/list-link dword. Dynamic is unchanged up to
that boundary. No external call is executed or replaced by a stub.

Three conditional controls project the original owner SAV's verified key
5879, baseline 10, Static baseline 0, count 0 and occupant presence into
independent allocations. Current costs 0, 10 and 255 become 10 before the
external call. This is a local store exclusion while that record baseline
remains 10; it is neither fault-time heap state nor a native invocation witness.

**Confidence.** High for native address/width/source, the local hash-hit
predicate and the measured pre-call effects, including all 256 byte values.
Medium for callable entry/body identification without an incoming target
anchor. Unknown for original entry reach and receiver identity, selected
fault-time record and any effect at or after the external-call boundary.

**Unknown.** Concrete incoming aliases, computed callers, L08162 and later
R0525/R1355 effects, post-call bit-4 paths, C18 fault-time cost history
and other writers. The 70 observed instruction starts include the unexecuted
boundary call; only 69 distinct starts execute. This is not whole-body or
whole-program execution coverage. A zero baseline can produce a zero local
store, but its presence or use during C18 is unobserved.

## Selected cell-prefix return effects

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-200 | In selected x86-32 mode, original callee L08162 returns immediately without memory writes or ECX dereferences; the selected caller's call supplies the return-stack write. | High / Unknown | ● active | [EXP-0468](../experiments/EXP-0468-cellcall-effects/) |

### TERR-200

The selected direct call at L11034 targets L08162 and returns at L11035.
The exact callee has one instruction: it reads a four-byte return address
from stack-segment:entry ESP in the selected x86-32 mode, advances ESP by four and transfers control there. It has no
nested call, conditional branch, ECX dereference or memory store. The direct
relative call anchors its entry. Adjacent padding and later code are outside
this complete normal-return body.

The selected caller forms its found node pointer plus 12 in EDI, then passes
EDI in ECX. That route comes from original bytes; the node's native identity
remains conditional on TERR-198. The callee does not use that pointer, so no
node/world alias is needed to establish its absence of memory writes.

EN/RU images are identical and count as one independent input. Original bytes
execute on 15 synthetic vectors per image. Eight isolated normal returns
read one four-byte stack word, write no memory and preserve general registers
and flags; ESP increases by four. Five original caller contexts execute the
real direct call and callee, then stop before the return-site instruction:
one four-byte caller stack write, one callee read and restored caller ESP.
Two invalid-stack vectors stop on unmapped access, outside normal returns.
Mapped/unmapped/null ECX and overlapping stack/data controls discriminate the
return source and pointer independence. Fourteen image-free controls exercise
instruction, byte-identity, memory-width, external-call, cap and output fences.

The local callee is excluded as a Cost-writing transition while these exact
original instructions execute. A caller stack write is distinct from a callee
write. Native stack/plane disjointness is not established. No native invocation,
fault-time C18 heap or whole-program writer census is inferred.

**Confidence.** High for the anchored complete callee and its local
normal-return effects. Its entire one-instruction body excludes unread local
stores and ECX dereferences; original call/return controls separate the caller
stack store. Unknown for actual C18 execution and writer history.

**Unknown.** Original incoming reach/node/stack/segment identity, faults or interrupts,
modified code, later selected callees, post-return effects and other writers.

## Selected call argument frontier

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-202 | Before unresolved call L11036, selected callee R0525 copies a four-byte stack argument to a new stack slot without pointee access; its continuation and C18 transition remain Unknown. | High / Unknown | ● active | [EXP-0469](../experiments/EXP-0469-cellcall-frontier/) |

### TERR-202

The selected direct call at L11037 targets R0525 and returns at L11038.
Its original construction reads a dword at incoming EBP+0x540b8 into ECX,
pushes the value and calls the exact callee. At callee entry, that value is
at stack-segment:ESP+4. The first instruction reads this four-byte word and
writes it unchanged at stack-segment:entry ESP-4; ESP decreases by four.
General registers and flags are preserved. The value is not dereferenced.

The selected 11-byte local shell has four instruction starts. Its next
instruction, at L11039, directly calls L11036. All successful prefix
probes stop before that unresolved call, including its implicit stack write.
The local cleanup and near return are conditional on a suitable normal
return from the unmeasured target. A lexical return is not proof that the
external target returns or leaves storage unchanged.

Section-specific PE mapping and raw relative displacements bind this shell
and the 12-byte caller construction, 23 unique bytes per image. EN/RU images
are identical, one independent input. Eight isolated vectors per image
perform one four-byte read and one four-byte stack write before the frontier.
Eight exact caller contexts perform two reads and three writes, all width
four: caller argument, caller return address and callee forwarded argument.
Four invalid source/destination controls stop on unmapped access. Thirteen
image-free controls discriminate source, writes, boundaries and guard loss.

Normal prefix execution requires readable source storage at entry ESP+4 and
writable destination storage at entry ESP-4 in selected x86-32 mode. Null,
mapped and unmapped forwarded values all reach the boundary. Flat synthetic
mappings do not establish original receiver, segment or pointer identities.
The caller's EBP-based field access also depends on its native segment base.
Field/argument-slot and field/return-slot controls preserve the loaded value
while showing a same-value store or source-slot overwrite. A native stack/C18
alias is unproved; changed bytes alone do not count all writes.

The bounded prefix writes only its stack destination and has no pointed-value
access. The whole callee remains a possible C18 transition through its
unmeasured continuation. Its involvement and every native destination alias
remain unestablished. No native invocation or whole-program writer census
is inferred. The 512-byte target and 128-byte caller navigation windows are
not the claimed effect population.

**Confidence.** High for anchored call construction, stack forwarding and
absence of a pointed-value access before the nested call. The unchanged
one-instruction prefix excludes the named local alternatives; original caller
contexts separate its write from the caller's two stores. Unknown for external
continuation effects/preconditions and the actual C18 transition.

**Unknown.** Original reach, incoming receiver, native segments, pointee and
stack/C18 aliases, L11036 effects, later calls, faults, interrupts, modified
code, C18 history and other writers.


## Nested cell-call domains and effects

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-204 | In selected x86-32 mode, L11036 returns with stack-only effects for a null argument; non-null input reaches unresolved L11040 before pointee access. Caller domain and native aliases remain conditional. | High / Unknown | ● active | [EXP-0470](../experiments/EXP-0470-nested-cellcall/) |

### TERR-204

The original direct call at L11039 anchors L11036. Its argument is the
forwarded four-byte stack word established by TERR-202. Two register saves
shift ESP by eight; the selected argument load at current ESP+12 therefore
reads callee-entry stack-segment:ESP+4. A zero test branches to local cleanup
and near return. The complete null path executes eight instructions, reads
four stack words and makes two four-byte stores: saved ECX at entry ESP-4
and saved ESI at entry ESP-8. General registers are restored, ESP becomes
entry ESP+4 and flags reflect the zero test. No pointee access or external
call occurs on that path. Stack stores remain after register restoration.

A nonzero argument executes six instructions before unresolved direct call
L11041 to L11040. It reads one stack argument and makes three four-byte
stack writes: the two saved registers and literal 9 at entry ESP-12. The call,
its implicit return-stack write and all external continuation effects are
unexecuted. Pointed storage need not be readable to reach this boundary;
that is not a precondition rule for the unresolved target.

The selected 104-byte callee slice has 41 decoded starts and two local return
sites. Raw relative branches target admitted starts. The 11-byte incoming
shell and 102-byte caller-domain slice give 217 unique original bytes per
image, bound by section-specific PE translation. Nine distinct callee starts
execute across both domains; later instructions are not execution coverage.
EN/RU images are identical and count as one independent input.

Original caller construction reads incoming EBP+0x540b8, tests it against zero
and can branch directly to its final reload. Otherwise it reads the unsigned
count at EBP+0x540bc. Count zero skips indexed reads; positive count uses a
reloaded field value as the base of dword slots at index*4. The final field
reload supplies the call argument. A null/table-base interpretation at callee
entry is conditional on original reach, source stability and segment bases;
it does not prove pointer validity or exhaust all incoming domains.

Each image has 22 original-byte synthetic vectors: seven null returns, ten
nonzero external-boundary stops and five invalid source/destination/return
controls. Isolated null/nonzero vectors preserve the stated callee event
counts. Incoming-shell and caller-branch vectors separate construction stores
from callee stores. Four caller controls cover field zero, count zero, two
empty slots and a field/index-slot alias. In the last, an original index-slot
store changes a nonzero field to zero before the final reload. This excludes
an unconditional argument-stability inference, without proving a native alias.
A same-byte saved-register control records two writes with zero changed bytes.
Eighteen image-free controls discriminate branch, source, write and guard loss.

Normal null return requires readable argument, saved and return words,
writable save slots and a permitted continuation in selected x86-32 mode.
The nonzero prefix needs a readable argument and three writable stack slots.
Fixtures have flat segments. Native segment bases and physical stack/C18
storage relationships are unproved. The complete null path has architectural
stack effects; its actual destination aliases remain conditional. The nonzero
continuation remains possible as a C18 transition. Neither path is a native
invocation or actual writer-history witness.

**Confidence.** High for the anchored local zero/nonzero branch, argument
construction, complete null effects and nonzero effects before the first
external call. Exact local paths exclude wrong offset/branch, additional
null-path operands/calls and pre-frontier pointee access. High for selected
caller field/count/array construction and reload, without a native receiver
or stability inference. Unknown for external effects/preconditions and C18.

**Unknown.** Native reach, receiver/table identity and validity, segments,
physical stack/pointee/C18 aliases, intervening changes outside measured
fixtures, L11040 and later effects, faults/interrupts, modified code,
actual C18 history and other writers.

## Map art and placement in three missions

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-GMAP-206 | Mission 90 (Fortress Kargallas) is world-map object `MapObject15`, point (478, 298), rectangle (477, 287, 27, 36), with no `Picture` entry; the rectangle covers the castle drawn in the 640x480 `GMap.bmp`, in EN and RU. | High / Medium | ● active | [EXP-0479](../experiments/EXP-0479-dialogue-art/) |
| TERR-STRUCT-207 | Mission 90 Castle lies at columns 104..114, rows 9..13 inside the playable rectangle; its former 4 px top-margin projection is partially retracted by TERR-221. | High / Medium | ● active (partially retracted) | [EXP-0479](../experiments/EXP-0479-dialogue-art/) |
| TERR-PLACE-208 | Mission 100's `100.alm` places 2 of its 11 class-76 turtles on border-ring cells (43,136) and (53,137); by static read the loader's seat call returns 0 on a ring cell, after which the listing logs and deletes the actor. | High / Medium | ● active | [EXP-0479](../experiments/EXP-0479-dialogue-art/) |

### TERR-GMAP-206

- Source: `scenario.res::globalmap.reg`, section `MissionObjects`, key `Mission90` is object 15
  (`REG-GMAP-066`). Both roots are identical in these keys.
- `MapObject15`: `MapPoint` [478, 298], `MapRect` [477, 287, 27, 36], no `Picture`. Of the 5 objects
  with a marker picture (`TOWN-038`, `TOWN-041`), objects 3, 13, 14, 18 and 19 belong to missions 50,
  60, 70, 100 and 110; mission 100 is `MapObject13`, picture `onmap13`, rectangle 20x16 at (208, 247).
  Mission 90 is not among them.
- Mission 91 is `MapObject28`, point (167, 73), rectangle 1x1, no `Picture`.
- The 640x480 24-bit `main.res::graphics/global.map/gmap.bmp` is the one background. The
  rectangle (477, 287, 27, 36) is a hit area over background pixels. A rendering of the region in
  scratch, not stored, shows the castle drawn in the background under the rectangle with a place-name
  label at its foot; the label is localized, which is why the EN and RU region hashes differ.
- An object with no `Picture` entry builds no marker (`TOWN-041`: a marker only for a picture other than
  "nothing"), and `TOWN-120` gives the ordered consumer list of the world-map paint.
- The position is the `MapPoint` anchor with the rectangle for hit-testing (`TOWN-038`).

**Confidence.** High for the registry values and the absence of a `Picture` key (all 35 objects
read). Medium only for a completed-mission overlay drawn outside the `TOWN-120` consumer list.

**Unknown.** Whether a completed-mission overlay is drawn from another resource at run time, and whether
a missing `Picture` entry reads as "nothing" beyond the condition `TOWN-041` states.

### TERR-STRUCT-207

The affected former wording below is partially retracted. See the
**Amended.** paragraph for the retained facts and correcting claim.

- `90.alm`: 144x144; type-4 structure record 28 is kind 57, `DescText` Castle (`Structure56`),
  footprint 11x5 cells, `FullHeight` 5, anchor (104, 9). Footprint: columns 104..114, rows 9..13.
- Playable rectangle (8, 8, 135, 135) (`TERR-SIGHT-116`): 0 footprint cells outside; 1 free
  playable row above the footprint and 21 free columns to its right.
- Altitude lift (`TERR-SPR-039`, mean of the four corner heights): row 9 is 28 on every column; over
  the footprint 28 to 43. The top edge is 9 x 32 - 28 = 260 px; the highest view origin row is 8
  (`SESS-VIEW-030`), top 256 px, for 640x480, 800x600 and 1024x768 views alike. Margin 4 px.
- Across the 28 campaign maps, 951 type-4 structure records in each root have 0 footprints outside
  the playable rectangle.
- By `TERR-STRUCT-100`, `TERR-STRUCT-101` and `TERR-STRUCT-106` (one 32x32 frame per footprint
  cell, anchored at the footprint's minimum cell) the sprite does not leave the map and is not cut by
  the scroll band, by rule; no frame was captured, and the 4 px margin depends on the lift convention
  (`TERR-SPR-039`) and on the scroll band being the only camera limit.

**Confidence.** High for the placement numbers, the playable rectangle and the 951-record census.
Medium for "drawn whole": the strip and lift rules were applied to the records, not compared with a
captured frame.

**Unknown.** A captured frame of the castle at the top of the scroll band.

**Amended.** TERR-221 partially retracts shared lift 28, world top 260 and the 4 px top margin. The footprint-centre corners are 28,28,28,44; shared lift is 32. At the selected creation offset 0, world top and highest view top are both 256. Placement, playable rectangle and structure-census facts stand. Native full-frame visibility remains Unknown.

### TERR-PLACE-208

- Mission 100 (`100.alm`, 144x144) has 11 type-6 placements of the turtle family (class 76): 3 in group
  Beists and 8 in group Friends. Friends cells: (43,136), (51,131), (53,137), (61,129), (57,135), (55,131), (49,134), (57,128).
- The first and third Friends cells are in the border ring (rows or columns outside 8..135):
  tile word `0x00f5`, block byte `0x1f` from the ring rule in `formats/terrain/movement.md`,
  mover `Turtle` tokenSize 1, movementType 1, mask `0x41`, so `block & mask` is non-zero.
- The other 9 cells have block byte `0` and lie inside the rectangle; none is inside a structure
  footprint.
- The `refused` column of the evidence table is the probe's own port of `block & mask` with the ring
  block `0x1f` taken from the format page; it restates the rule and does not test it.
- Static read of `rom.exe`: at `L11042` the loader calls `R1145(actor, x, y, 0)`; a zero
  return reaches the log string `L11043` ("Can't place unit %d during loading on the tile(%d, %d)."), then the actor's
  virtual delete (`[vtable+4]`, argument 1) and the loop continues at `L11044`.
- The ring census over all 28 maps in each root, 2333 type-6 placements, finds 4 ring cells: the two
  above, a class-76 unit at (77,136) in `80.alm` (group Beits) and a class-64 unit at (54,136) in
  `90.alm` (group Monsters).

**Confidence.** High for the cell, tile and block values. Medium for the refusal and destruction: the
call chain from `R1145` to the predicate `R1287` and the block byte was read from
claims and a listing, and no run observed the log line.

**Unknown.** Whether the log string surfaces anywhere visible; the in-game turtle count at start.

## Structure destruction

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-STRUCT-210 | Of the 66 structure classes, 33 draw their ruin grid when published health is not positive, 2 draw their intact art, and 31 draw no change: 26 have maximum HP zero, 3 more are Indestructible and 2 use bridge selectors. | High / Unknown | ● active | [EXP-0480](../experiments/EXP-0480-structures/) |
| TERR-STRUCT-211 | 39 of 950 authored Building placements in the shipped missions carry placement word zero, so they start with HP 0; their ruin-grid draw at mission start is High / Unknown, because mission-start publication is not read. | High / Unknown | ● active | [EXP-0480](../experiments/EXP-0480-structures/) |
| TERR-STRUCT-212 | In 339 distinct files, structures with a nonzero placement word are below maximum HP only on four classes of mission 131; apart from two Switches, one Bee House at (131,84) reached HP 0. | High / Unknown | ● active | [EXP-0480](../experiments/EXP-0480-structures/) |

### TERR-STRUCT-210

- Gate (`TERR-STRUCT-102`, `SPR256-STR-041`, `UNIT-STRUCTCONT-078`): the client draws the ruin grid when
  its health word is not positive and the class is not Indestructible. The health word is the Building's
  word `+0x42` when its word `+0x44` is nonzero, else the literal 1000 (`L09288..L09289`).
- Per class, from `structures.reg` `Indestructible`, the sheet's grid count, and the `Data.bin` Buildings
  column 3 of the same 1-based row (`evidence/structure-classes.tsv`):

| Draw after HP is not positive | Classes | IDs |
|---|---:|---|
| ruin grid (`frames - TileWidth*FullHeight + cell`) | 33 | 1-7, 9-13, 17-26, 28, 29, 43, 44, 46, 47, 49-51, 56, 57 |
| intact art again (no ruin grid ships, base 0) | 2 | 14 Cave, 45 Inn 3 |
| no change: maximum HP zero, so the health word is 1000 | 26 | 27, 30-32, 34-36, 38-42, 48, 52-55, 58-66 |
| no change: Indestructible with positive maximum HP | 3 | 8 Well, 15 and 16 Magic Well |
| separate bridge selectors (`TERR-STRUCT-105`) | 2 | 33, 37 |

- Precedence: `evidence/structure-classes.tsv` tests maximum HP zero before Indestructible, so the table's
  count of 3 covers only classes with a positive maximum. In the committed `structures.csv` 31 classes have
  Indestructible 1: the 26 maximum-HP-zero classes, both bridges and Wells 8, 15 and 16.
- Classes 28 and 29 are the two Switches (maximum HP 1): the same ruin grid is their toggled state
  (`UNIT-STRUCTUSE-090`).
- Neither a goblin or orc dwelling (IDs 1-4, 300 HP) nor the Cave (ID 14, 1000 HP) is Indestructible.
  A destroyed dwelling draws its ruin grid; a destroyed Cave draws its unchanged art.

**Confidence.** High for the class table: each input is an instruction-level gate or a data census, and
the table is their join. Unknown: a captured frame, the drawable's health before any publication, whether
anything removes a destroyed structure (`UNIT-STRUCTBOUND-081`), and the RU `structures.reg` and sheets (the
Indestructible and grid columns come from the EN edition of the structure-art census; only `Data.bin` was
re-read from RU, with an identical table).

**Open conflict.** `SPR256-STR-041` says 4 of the 39 two-grid classes are Indestructible. The same census file
gives 6 (IDs 8, 15, 16, 39, 40, 66). This experiment does not correct that card; the difference does not
change the table above, where those classes fall in the ruin-grid or no-change rows by the same gate.

**Evidence.** [EXP-0480](../experiments/EXP-0480-structures/), `evidence/structure-classes.tsv`,
`evidence/static-facts.tsv`.

### TERR-STRUCT-211

- Placement word `+0x0c` of a type-4 record: zero stores HP 0 into the new Building and nonzero leaves
  the definition's maximum (`ALM-128`). Census over every `M7R` map of `scenario.res`, both roots
  (`evidence/authored-placement-words.tsv`): 950 non-Shop Building records, 39 with word zero; EN and RU
  agree on every map and kind.
- The 39 are on maps 111, 120, 150, 40, 50, 61 and 81, on kinds 6, 7, 9, 10, 11, 17-20, 23, 44, 46, 47
  and 51. All 39 are in the 33 ruin-grid classes of `TERR-STRUCT-210`.
- Drawing: the client gate of `TERR-STRUCT-210` applies to these classes, but no instruction read shows that
  mission start publishes the Building to the client (`SAV-1166`), and no frame was captured. The census, the
  word and the saved HP 0 are High; the draw at mission start is Unknown.
- Saved state agrees: 134 of 134 saved records whose authored word is zero hold HP 0 and none holds a
  positive HP (`evidence/placement-word-vs-saved-hp.tsv`).

**Confidence.** High: a census over both installs, with the saved records as a separate check.

**Evidence.** [EXP-0480](../experiments/EXP-0480-structures/), `evidence/authored-placement-words.tsv`.

### TERR-STRUCT-212

- Population: 339 distinct files (SHA-256) under the seat saves root and the GOG install root (root,
  `oldsaves*`, an engine profile and experiment directories that include engine-written files); 338 parse with the exact reader, one has magic `Bsg&`. 201 carry
  a world half: 7,232 Building-family records on 17 map names, joined to the authored type-4 record by
  cell and kind (`evidence/save-structure-join.tsv`).
- 6,734 of the 7,232 records join; the other 498 are in 19 generated-document files of maps 10 and 20,
  which match by record count only. Of the 6,734: 1,892 have maximum HP zero, 134 have authored word zero,
  4,669 have a nonzero word and full HP, 24 have partial HP (kinds 24 and 25 of mission 131) and 15 have
  HP 0 (`evidence/placement-word-vs-saved-hp.tsv`).
- Authored word nonzero and HP 0: 15 records in 8 files, three structures of mission 131.
  Bee House at (131,84), kind 24, maximum 100: HP 0 in 8 distinct files of one session lineage (they are
  not 8 independent observations) and 100 in 7 others, of which 4 have different footprint rows (`21/21`,
  `+0x48` = 0; `SAV-1164`).
  Switch at (97,91), kind 28, maximum 1: HP 0 in 6 files. Switch at (67,36), kind 29: HP 0 in 1 file.
  The Switch HP 0 is the toggled state (`TERR-STRUCT-210`) and the corpus cannot separate it from damage.
- Absent: no class other than kinds 24, 25, 28 and 29 is below maximum HP in the 6,734 joined records.
  The 994 records of kinds 1-4 (goblin and orc dwellings) and the 100 Cave records are all at maximum.
- No saved HP is negative: of 149 records with HP not positive, all are exactly 0. The 149 are 134 authored
  zeros plus the 15 records of the single in-play event above, not 149 independent kills.
- The 8 HP 0 files are taken as original-play files only by footprint signature (rows `25/25`, `+0x48` = 2);
  their producer chain is not established.

**Confidence.** High for the census over this population. Unknown: which weapon, spell or hit sequence
set HP 0 (physical strikes are unclamped, `UNIT-STRUCTDELIVERY-065`, so a physical kill could leave a
negative value that no file shows), the producer of every file, the cause of the HP 0 (no census of writers of `+0x42` was run beyond the damage
writers and the Switch use arm), and any structure outside the corpus.

**Evidence.** [EXP-0480](../experiments/EXP-0480-structures/), `evidence/nonpositive-hp.tsv`,
`evidence/placement-word-vs-saved-hp.tsv`, `evidence/census/`.

## Map edges

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-216 | The inspected camera-clamp bodies use dimensions, spans and origins, not grid content; 342 supplied original-x86 cases match their endpoint arithmetic. | High | ✔ promoted (branch candidate) | [EXP-0499](../experiments/EXP-0499-map-edges/) |
| TERR-217 | The four terrain routines admit rows 0..spanRows+3 and columns 0..spanCols-1; a nonempty eight-cell camera band keeps the stored outermost terrain cells out of that population. | High | ✔ promoted (branch candidate) | [EXP-0499](../experiments/EXP-0499-map-edges/) |
| TERR-218 | The software painter intersects its repaint clip with the absolute widget view and resets it to that view when the camera changes; this is a screen rectangle, not a playable-cell bound. | High | ✔ promoted (branch candidate) | [EXP-0499](../experiments/EXP-0499-map-edges/) |
| TERR-219 | In 18 supplied original-x86 cases, ordinary and mirrored indexed body blits write only the sprite/clip intersection and preserve every outside or rejected destination pixel. | High | ✔ promoted (branch candidate) | [EXP-0499](../experiments/EXP-0499-map-edges/) |
| TERR-220 | The enumerated 38 EN and 34 RU maps have nonempty camera bands at all three resolutions; first playable row 8 is terrain-blocked on 3 maps in each root and partly open on the other 35/31. | High | ✔ promoted (branch candidate) | [EXP-0499](../experiments/EXP-0499-map-edges/) |
| TERR-221 | Mission 90 Castle uses footprint-centre corners 28,28,28,44 and shared lift 32; its first strip at the selected creation offset meets the highest view top, with a computed margin of 0 px. | High / Medium | ✔ promoted (branch candidate) | [EXP-0499](../experiments/EXP-0499-map-edges/) |

### TERR-216

`R1677` through `L13363` clamps absolute targets then writes
pending deltas. `R1678/R1679` clamp axis deltas. The prefix
`R1280..L13364` of `R1280` recomputes spans and reduces an
excessive origin to its upper bound; it adds no lower clamp.
`SESS-VIEW-028/030` supply the viewport and eight-cell band.

`clamp-cases.tsv` contains 342 matching original-x86 executions over spans
15x15,20x18,27x24, dimensions 16,32,40,64,144,256 and six target pairs.
`clamp-accesses.json` preserves every non-stack read/write location.
Grid pointers are null. Data reads touch only the view's dimensions, spans,
origins, pending deltas and landscape-presence pointer, plus landscape W/H
for span recomputation. The absolute body stops before notification; the
span prefix stops before allocation.

**Confidence.** High for these complete arithmetic bodies/prefix and the
finite execution population. Raw PE ranges and access traces rule out a
content-dependent branch inside them. Synthetic crossed bands demonstrate
upper-after-lower behavior; they are not shipped frame observations.

**Unknown.** Session notification, native scroll cadence and native small
maps. Other operations can set or restore an origin.

### TERR-217

`R0539`, `R1814`, `R1815` and `R1816`
walk rows 0 through spanRows+3 ascending and columns spanCols-1 through 0
descending. Each adds view origins to obtain world cells. With
`SESS-VIEW-030`'s nonempty band, columns stay in 8..W-9 and the greatest
terrain row is H-5. At the bottom stop, four ordinary terrain rows H-8..H-5
are inside the movement border. The top, left and right strips are not
terrain-dispatched by these loops. Altitude can move admitted pixels.

`terrain-renderers.txt` retains the original instructions. The full software
redraw call at L13365 enters the ordinary terrain renderer; it has no
preceding black-clear call in that arm. Conditional color-zero grid lines
are separate from border treatment. Tile clipping can preserve pixels.
The mesh and drawable margins have different bounds. The smoothed-height
guard at L13366 compares an unshifted row; its sampled base subtracts 1.

**Confidence.** High for the four loop headers, world indices and nonempty
band consequence. The raw bytes are translated through their PE section.
The stored outermost ring is not a terrain-dispatch population in this
scope; a projection read is not a draw call.

**Unknown.** Native per-pixel coverage and preceding surface contents;
selector-mask-0x02 alternate rendering and invalid camera bands.

### TERR-218

`R1224` resolves both rectangle corners through the widget ancestry.
The map painter calls it at L13367 using view+8. It intersects view+0xf4
with that absolute view at L13368 through imported IntersectRect. When
the origin differs from saved view+0xec/+0xf0, L13369..L13370 copies
the entire view rectangle to view+0xf4. Full-redraw/overlay state has other
reset paths. Repaint work and child overlays can reduce the active area.

`R0373` copies four dwords into clip globals L01503..L01506.
The software painter installs view+0xf4 for drawable work; `TERR-SPR-140`
identifies its four installations. The source is an absolute screen view,
not movement bounds or a per-cell destination window.

**Confidence.** High for the named producer, imported rectangle operation,
camera-change reset and setter. Instructions and import names are read from
the same lawful image; no native active value is inferred from a screenshot.

**Unknown.** The precise active repaint subset at an arbitrary native frame,
alternate rendering and overlay population during that frame.

### TERR-219

Structure/body wrappers `R0794`, `R1791` and `R1547`
reach indexed .256 body blitters, including ordinary `R0796` and
mirrored `R0895`. Their rejection/inside tests use global screen
clip edges; partial drawing tests rows and pixels against the same edges.

`body-clip-cases.tsv` runs both bodies on a 64x64 16-bit destination with
supplied half-open clip (4,3,52,51) and a synthetic 40x40 literal-run sprite.
Nine placements per body exercise four crossed edges, an interior sprite
and four whole rejections. All 18 complete original-x86 executions match
every destination pixel. The interior case crosses the 32-pixel cell line.
No cell rectangle is supplied to these blitters. Outside pixels retain
the 0xa55a sentinel; accepted pixels match the indexed palette, reversed
for the mirrored body.

**Confidence.** High for the named instruction paths and supplied finite
population. Whole rejection and partial clipping are distinct cases.

**Unknown.** Native receiver binding, registration, frame selection, fog,
malformed streams and the alternate overlay body blitters. Shadow geometry
has its separate `TERR-SPR-140` contract.

### TERR-220

`tools/mapedges` enumerates every M7R member in both scenario containers
and all root ALM files. The resulting population is 38 EN maps/880,704 cells
and 34 RU maps/677,696 cells. Twenty source files and every map are hashed.
All 216 map/resolution camera bands are nonempty. All signed Altitudes bytes
are 0..127, with zero negative values.

Over columns 8..W-9, row 8 has three all-terrain-blocked maps per root:
140.alm,141.alm,91.alm. The other 35 EN and 31 RU maps have some terrain-open
cells. On 140.alm row 8, blocked/open is 240/0; on 90.alm it is 60/68,
including five water cells. Both locales have those same two row populations.
Stored row 0 is counted separately. Four full-side eight-cell strips include
each corner in two strips. `maps/regions.csv` preserves the populations.

Terrain-only classification uses `TERR-PASS-049/050`'s tile bit13, raw water
and class8 predicates. It excludes type3 objects, the stamped simulation
border, occupants and runtime flags. Unknown class pairs remain Unknown;
none occurs in this measured corpus.

**Confidence.** High for this exact install/data population and derived
counts. Strict framing/length checks and sorted enumeration prevent silent
map omissions. EN/RU are data populations, not two code witnesses.

**Unknown.** Native frame differences between blocked/open rows, loose maps
outside the two roots and runtime changes to movement or terrain state.

### TERR-221

The mission 90 Castle placement remains `TERR-STRUCT-207`'s kind 57,
anchor (104,9), 11x5 cells, FullHeight 5 in the 144x144 map. The shared lift is
sampled at footprint centre (109.5,11.5), not the anchor cell. Both map hashes
are 1af4b4ff040d48f4293020101d5042092328f2794254f438cfd8253aa6a006e1.
Its four signed heights are 28,28,28,44. `R0614` performs three
signed interpolations with truncation toward zero, yielding 32. Four
supplied original-x86 executions reproduce centre 32 / anchor 28 in both roots.

The selected creation writes drawable+0x10=0 at L10620. The strip body
`R1794` subtracts that field and the shared lift from row*32.
At origin row 8, the first Castle strip is (9-8)*32-0-32=0, independently
of the three resolution spans. The world top and highest view top are 256.

Footprint parity can change interpolation weights even when the placement
anchor is at a half cell. Nested truncation is not an unconditional mean:
supplied original-x86 corners 1,0,0,0 at weights 16/32 give lift 1, while
their exact four-corner mean is 0.25. This narrows `TERR-STRUCT-106`'s mean
consequence without changing its named arithmetic.

**Confidence.** High for measured corners, original integer arithmetic,
the selected create-time offset and controlled lift executions. Medium for
the first-strip placement applied to those inputs; no native full frame
was executed. The supplied owner crop has no measured camera origin.

**Unknown.** The current drawable+0x10 and frame in the owner observation,
subsequent offset writers and the complete native Castle silhouette.

## Unit health and mana bars

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-224 | Original unit bars select red below trunc(max HP/4), yellow below trunc(max HP/2), else green; four filled RGB rows use intensities 128,255,192,128 before surface packing. | High | ● active (branch candidate) | [EXP-0501](../experiments/EXP-0501-status-bar-colour/) |
| TERR-225 | Original unit bars overwrite unfilled rows with RGB grey 64,128,96,64, except unselected show-health bars: their unfilled pixels are untouched and their filled pixels blend with the destination. | High | ● active (branch candidate) | [EXP-0501](../experiments/EXP-0501-status-bar-colour/) |
| TERR-226 | The original unit mana bar has fixed blue filled rows of intensities 128,255,192,128 before surface packing; show-health blending also applies to mana, and nonpositive maximum mana skips its bar. | High | ● active (branch candidate) | [EXP-0501](../experiments/EXP-0501-status-bar-colour/) |
| TERR-227 | Original unit bars use inner width class.Right-class.Left-8 and signed trunc(width*current/maximum); a zero quotient becomes 1 for nonzero HP and for any mana, with no clamp to the inner width. | High | ● active (branch candidate) | [EXP-0501](../experiments/EXP-0501-status-bar-colour/) |

### TERR-224

The read path is CUnit/CAirUnit vt+0x30, `R1106`, and helper
`R1824`. Current HP H at unit+0xfc and maximum HP M at unit+0x100
are signed 16-bit inputs. The caller compares H with trunc(M/4), then
trunc(M/2), using signed comparisons at `L13343` and `L13344`.
For positive M, truncation is floor; equality advances to the next colour.
For M=101, H=24 is red, H=25 is yellow, and H=50 is green.

The four rows at Y-2, Y-1, Y and Y+1 have source RGB components
`[128,255,192,128]` in red alone, red and green together, or green alone.
These are generated channel values, not palette indices. A channel c
is packed as `(c >> (8-bits)) << shift`, ORed across R/G/B.
TERR-225 defines the conditional destination blend.

**Confidence.** High for this original 16-bit path. Complete caller/helper
instructions exclude an absolute-current selector, a third state selector
for the colour family, floating-point ratios, and rounding the ratio first.
Ghidra 12.1.2 and Capstone 5.0.7 agree on every instruction boundary in
seven bounded ranges. Unicorn 2.1.4 executes the original instructions
on 1056 bar vectors across two injected pixel formats; all interior pixels
are checked. EN/RU are one byte-identical executable, not two witnesses.

**Unknown.** Native display format choice and other bar paths are not measured.
The original process is not launched; caps and the numeric label are excluded.

### TERR-225

Helper arguments are `(L,R,Y,N,bright,middle,dark)`. Its interior starts
at L+4 and ends at R-4. The ordinary branch at `L13345` writes four
background rows with source RGB `(64,64,64)`, `(128,128,128)`,
`(96,96,96)`, `(64,64,64)`, then overwrites the filled N pixels.

When view+0xaa0 is nonzero and unit+0x7c is zero, the caller halves each
packed fill word with `(word >> 1) & mask`; the helper instead shades
the existing filled rectangle with shroud row 8 and adds the halved fill.
Its rectangle ends at L+4+N, so the unfilled pixels are not touched.
For packed destination B and source fill C, the result is
`((B >> 1) & mask) + ((C >> 1) & mask)`, with masks 0x7bef for RGB565
and 0x3def for RGB555. TERR-FOG-084 supplies the shroud-row law.
The four filled row intensities retain TERR-224's order.

Unit selection method `L13346` stores its argument at +0x7c
(`L13347`); AI-CURSOR-202 connects this flag to the selection count.
MENU-057 identifies view+0xaa0 as the Ctrl+H toggle. The map painter
admits selected objects in all three grids. In +0x8c and +0x90 it also
admits unselected objects when show-health is on and object+0x78 is zero
(`L13348..L13349`, `L13350..L13351`). The +0x94 grid requires
selection. This narrows TERR-SPR-048's selected-only wording for all grids.

**Confidence.** High for the named guards and original interior pixel writes.
The probe covers all four selection/show-health combinations in RGB565 and
RGB555. The original overwrite, additive and shroud-lookup pixel loops run;
only the row-8 lookup data is injected from TERR-FOG-084. The unfilled
population is the interior remainder for N between 0 and W. This is not
a global no-write claim about cap sprites, other render passes or N>W.

**Unknown.** Native framebuffer contents and the cap sprite pixels are unobserved.

### TERR-226

At `L13352..L13353` the caller reads signed maximum mana from
unit+0x13e and skips the mana bar when it is not positive. It reads
signed current mana from unit+0x13c at `L13354`.
The colour construction at `L13355..L13356` generates blue
intensities 255,192,128 with zero red and green, without comparing
current mana to any colour threshold. The helper lays them out as
`[128,255,192,128]`. Mana Y is health Y+4. The selection/show-health
predicate and pixel operations are the same as TERR-225.

**Confidence.** High for the complete mana arm and its shared helper.
The original instructions produce fixed blue across all measured current/max
mana ratios, including zero. A health-like red/yellow selector is excluded
from this arm, not from every mana display in the executable.

**Unknown.** Other mana display paths and native cap appearance are not measured.

### TERR-227

The caller forms L/R from class+0x84/+0x8c with the same screen offset,
then computes W=R-L-8 at `L13357..L13358`. Fill starts at L+4.
Maximum and current values are signed 16-bit. The multiplication keeps
the signed 32-bit low product; IDIV truncates its quotient toward zero.
For ordinary positive W and maximum, N=trunc(W*current/maximum).

HP quotient zero becomes 1 only if the original HP word is nonzero
(`L13359..L13360`); HP zero keeps N=0. Mana quotient zero always
becomes 1 (`L13361..L13362`), including current mana zero.
Maximum mana must be positive; the HP division has no zero-denominator
guard in this routine.

Neither quotient is clamped to W before the helper. The overfull
control current=150, maximum=100, W=17 passes N=25 for both bars.
The reached pixel helpers separately clip rectangles to the viewport;
a nonpositive rectangle width writes no pixels. These are raster bounds,
not saturation of the current/max ratio.

**Confidence.** High for the two complete arithmetic blocks and three reached
pixel writers. The probe distinguishes truncation from nearest rounding
with W=17, current=1, maximum=3 giving N=5; it checks HP zero, mana zero,
the one-pixel minimum, skipped nonpositive mana maxima and overfull lengths.

**Unknown.** Native admission of nonpositive HP maxima, negative widths,
negative current values and overflowing products is not established.

## Cell planes at mission start

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TERR-228 | A fresh mission writes the static and dynamic planes in this order: ingest, Building attach, type-6 occupy, type-9 cell casters, Sacks, then the join walk's occupy in the first sub-tick. | High / Medium | ✔ promoted | [EXP-0525](../experiments/EXP-0525-start-planes/) |
| TERR-229 | A map's type-9 record with X or Y nonzero and A below 4 creates a cell record at map load; on mission 10 two such cells carry rows 20/20 that no structure, actor or Sack explains. | High | ✔ promoted | [EXP-0525](../experiments/EXP-0525-start-planes/) |
| TERR-230 | The ingest cost of the 1,653 record cells at mission start is the map.reg Cost value; the recompute is the only later cost writer reached, writing CostCracked on structure cells the &0xfa arm opens. | High / Medium | ✔ promoted | [EXP-0525](../experiments/EXP-0525-start-planes/) |

### TERR-228

- The map loader `R0128` (`w01`) allocates 0xa4558 bytes (`L00237`), calls the world
  constructor `R0118` at `L00242`, publishes it to `[L04624]` (`L10488`), then calls the
  type-4 walker `R0489` (`L09682`), the type-6 spawner `R0151` (`L00958`), the walk
  `R0067` with argument 1 (`L00959`) and `R0461` (`L02208`), with no branch between them.
- `R0118` (`w05`) calls `R0116` at `L14273` and the ingest `R0278` at `L08344`, its last
  call before the return. The border writer `R0470` has two direct callers, `L14274` inside the
  ingest and `L14275` (`s01`). The ingest arms are `TERR-PASS-049`'s.
- Writers in that order, each with its condition:
  1. Building attach (`R0489`): one record per `BuildingPresent` cell, first building per cell, recompute with the `|5` or `&0xfa` arm (`TERR-STRUCT-071`, `TERR-STRUCT-076`).
  2. Type-6 placement: `R1145` places an actor only where `R1287` finds no common bit between the Mover mask and the two planes, then the occupy fills the domain slot over the n by n footprint and recomputes (`MOVE-120`, `MOVE-121`).
  3. `R0461`, first loop: type-9 cell casters (`TERR-229`), listed in `w02` (`L14276`..`L14277`; the loop's exit `JGE` at `L14278` goes to `L02209`). `w02` ends at `L14279`; the Sack loop and the type-9 loop's tail are not listed. The Sacks' place after the casters rests on `ITEM-OWNED-028`, which puts the type-8 lookups at `L04680`..`L02084`, past that exit, and on `SAV-SACKENTRY-590` for the refusal on dynamic bit 0 (that card does not name `R0461`). Medium.
  4. The fresh-map arm of `R0512` (`w04`) then calls `R0474`, `R1575` and `R0946` under `server+0x124`/`+0x11c` tests and `R0945` only when `server+0x120` is zero; `R0946` and `R0945` are the `.ini` and random-scatter Sack origins, neither of which runs on a shipped campaign map (`ITEM-SPAWN-026`, `ITEM-SPAWN-027`).
  5. First sub-tick: the join walk places the carried party through `R1145` and the occupy (`MOVE-122`).
- The walk `R0067` (`w03`) reaches a plane writer only in its opcode-`0x10003` arm (`L13792`..`L14280`): it moves every actor of the group the node names into a new object, calling the removal `R0865` per actor at `L07140` (`MOVE-089`). No node with opcode 65539 (`0x10003`) occurs in `EXP-0155-trigger-closure/evidence/fixtures.csv`: 28 maps per root, 1,602 EN and 1,600 RU rows, including the five maps of this census (counted from the committed table; no claim ID states it). The EN root holds 38 maps by `TRIG-DROP-013`'s count, so 10 EN maps are outside that table. Medium.
- Direct callers (`s01`, capstone linear sweep plus a raw dword scan of every section, 0 stored addresses for any target): recompute `R0453` 16 sites; claim set `R0257` 1 (`L07007`, transit start); `R0058` 3; `R1332` 1; claim clear `R0046` 7; bit-4 writer `R1080` 2, both in the area module (`TERR-PASS-148`); `R0227` 3 (`L12574`, `L14281`, `L14282`).
- Census agreement (`SAV-1223`): over 9 restart slots no saved bit lies outside these sources; no claim bit, no bit 4 outside the border.

**Confidence.** High for the loader order, the constructor's two calls, the type-9 first loop's condition and call, and the `0x10003` arm's removal call: each is an instruction in a listing read whole (`w01`..`w03`, `w05`). Medium for the Sacks' place after the casters and for "no other plane write in `R0461`" (the listing stops at `L14279`), for the fresh-arm gates of `w04` (listed only as call sites; `R0474` and `R1575` are identified as readers by `SAV-` claims and not read here), and for "no campaign map carries a `0x10003` node" (28 of 38 EN maps tabulated). Medium that no other writer runs between the ingest and the restart save: the callers at `L06784` (`R1332`) and `L01945` (`R0058`) lie in routines not identified here, the plane displacement populations of `TERR-PASS-148` are not reclassified, and the first sub-tick's actor and area ticks were not read. The census excludes a writer that leaves a visible bit in the window on these 9 files; it cannot see a write of the value already present.

**Unknown.** The owners of `L06784` and `L01945` and whether they run before the restart save. The body of `R0461` from `L14279` (the type-9 loop's tail, the loop body between `L14283` and `L02209`, and the Sack loop) and the bodies of `R0474` and `R1575`, which could write a plane. Whether the 10 EN maps and the RU maps outside the `EXP-0155` table carry a `0x10003` node.

### TERR-229

- `R0461`'s first loop walks the list at `map+0x2f0` (`w02`, `L14276`..`L14277`). A record whose dword `+0x08` or `+0x0c` is nonzero and whose word `+0x10` is below 4 (`L12573`..`L14284`) builds six bytes: record `+0x16` and `+0x18`, then bytes 0 and 2 of the first two tail elements (`L14285`..`L14286`), and calls `R0227` with the cell `(low byte +0x0c << 8) | low byte +0x08` (`L14287`..`L12574`). This is `ALM-T9CONTROL-175`'s first write; `R0227` creates or reuses the record, refuses a cell whose dynamic bit 0 is set and stores the six bytes at `+0x2c..+0x31` (`TRIG-CELLTAIL-035`).
- On the wire (`ALM-TRIG-049`) the six bytes are spellRaw bytes 0 and 2, then the kind and low words' low bytes of tail elements 0 and 1.
- Census (`SAV-1224`): 4 of the 5 maps hold type-9 records (1, 11, 3 and 4 records); 6 qualify (1 on 20.alm, 2 on 10.alm, 3 on 41.alm). Each qualifying cell holds a saved record whose tail `+0x2c..+0x33` equals the prediction byte for byte, 8 records over 4 files. The two 10.alm cells, 21,63 and 22,64, have no structure, actor or Sack; their rows are static 0x20 / dynamic 0x20 in both mission-10 files. The other four sit on Building cells.
- Record `+0x2c` holds 13 (10.alm), 6 (20.alm) and 6, 15, 24 (41.alm). None is 26, so `R0039`'s relocation arm does not apply to these cells.

**Confidence.** High: the condition, the six-byte build and the call are instructions in one listing, and the census matches every saved tail and both residual rows that the earlier model left unexplained.

### TERR-230

- Order: `R0118` calls `R0116` before the ingest (`TERR-228`). That `R0116` reads the ten `Cost` keys into `world+0x54177..0x54180` rests on `ALM-TERR-043`; its key reads are not in a committed window here (`w05` stops at the call at `L14273`). The ingest stores the classifier's blended cost for every cell (`TERR-PASS-049`).
- Census: the cost baseline `payload+0x00` of all 1,653 saved records equals the ingest cost computed with the map.reg values (`[8,8,8,14,6,12,8,16,8,6]`). On the 376 records where rom.exe's defaults give another value, the save holds the map.reg value 376 of 376 (`SAV-1224`). The 376 are counted per file (9 files, 5 maps); the distinct (map, cell) count is not computed, since the census does not list the cells. By map it lies between 219 (the largest per-file count of each map: 99 on map 20, 32 on map 10, 24 on map 41, 61 on map 30, 3 on map 141) and 376. Only these 1,653 record cells are observed; no cost plane is saved.
- After the ingest, the mission-start writers of `TERR-228` reach the cost plane only through the recompute `R0453`: it restores `payload+0x00`, writes `CostCracked` (6) on a structure cell the `&0xfa` arm opens, and shifts left once per layer slot (`MOVE-085`, `MOVE-086`). That the recompute is the only cost writer after the ingest rests on the caller scan `s01` (callers of named routines, 16 sites of the recompute), not on a pass over stores into the cost plane. Every saved record has layer count 0 and six empty layer slots, so no shift applies.
- The model in `tools/startplanes` therefore changes 13 to 25 cells per map from the ingest cost, all opened structure cells; these values are the model's, not observed.
- `R1087`'s divide needs a nonzero layer count (`MOVE-085`), which no start record holds.

**Confidence.** High that the ingest cost of the 1,653 record cells uses the map.reg values on these five maps: 376 discriminating baselines (counted per file over 9 files; between 219 and 376 distinct cells). Medium that the recompute is the only cost writer after the ingest among the routines of `TERR-228`: the basis is the caller scan `s01`, and a baseline captured at record creation cannot show a later write. Medium for the `Cost` key reads of `R0116` (`ALM-TERR-043`, not re-read), for the cost plane outside record cells and for the opened cells: no cost plane is saved, so those values are the model's, not an observation. A sweep of stores into the cost plane would bound the writer set.
