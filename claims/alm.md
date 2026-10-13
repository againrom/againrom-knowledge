# Claim registry — ALM (.alm map)

Level 2 ledger. Index: [registry.md](registry.md) · spec: [`formats/alm/format.md`](../formats/alm/format.md). Legend: `✔ promoted` (also in the spec) · `● active` · `✖ retracted`. IDs are permanent.

> **Confidence re-graded by the standing sweep (2026-07-25).** Every row below was re-judged
> against the evidence it already cites, under the discriminating-evidence rule
> (the research `METHODOLOGY.md` confidence scale). No new bytes were read. Where a
> clause rests on corpus agreement alone it now reads Medium; where a figure was measured on a
> base that later moved it is flagged in place. See `../docs/REGISTRY-LOG.md` → *Standing corrections*.

> ### ⚠ EXP-0030 — corpus-wide framing correction (read this before using any offset below)
>
> [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) proved from `rom.exe` and the shipped
> bytes that EXP-0007's section split is **8 bytes too early**. The true framing is a **20-byte
> file header** and, per record, a **20-byte header `[tag=7][hdrLen=20][payloadSize][typeId][f32]`**
> followed by `payloadSize` bytes of **pure payload** (`ALM-FRAME-031`). Consequences:
> - **Every payload offset published before EXP-0030 must be read −8.** The *byte positions* in the
>   file are unchanged; only the labels move. E.g. type-0 `W@+0x08` → `+0x00`; type-6 `X@rec+8` →
>   `rec+0`; type-7 `entryCount@+8` → `+0`.
> - **The grids gain 8 bytes at the front**: cell `(col,row)` = element `row·W+col` of the record
>   payload = **`+4` cells (type1 u16) / `+8` cells (type2, type3 u8)** relative to the old payload
>   (`ALM-GRID-032`). *This is a real decode bug in anything built on `ALM-GRID-011`.*
> - There is **no `[typeId][f32]` overlay on the grids, no per-section `f0`/`f1`, and no 8-byte
>   file trailer**.
>
> Rows amended by EXP-0030: `ALM-HDR-001`, `ALM-SEC-002`, `ALM-SEC-003`, `ALM-SEC-004`,
> `ALM-TRL-005`, `ALM-HDR-006`, `ALM-META-008`, `ALM-META-010`, `ALM-GRID-011`, `ALM-CNT-017`,
> `ALM-UNIT-018`, `ALM-OBJ-019`, `ALM-GRP-020`, `ALM-TRIG-021`, `ALM-TRIG-022`, `ALM-META-024`,
> `ALM-META-025`, `ALM-META-026`, `ALM-META-027`, `ALM-META-028`, `ALM-HDR-029`, `ALM-TRL-030`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-T9CONTROL-175 | Type9 A has interacting predicates in the complete builder and an arithmetic use on the linked-item arm. | High / Unknown | ✔ promoted | [EXP-0354](../experiments/EXP-0354-type9-lifetime/) |
| ALM-T9ACTOR-176 | The A&8 path supplies an actor Spellbook entry before the A&4 building path; neither receives A as a constructor parameter. | High / Unknown | ✔ promoted | [EXP-0354](../experiments/EXP-0354-type9-lifetime/) |

### ALM-T9CONTROL-175

**Type9 A has interacting predicates in the complete builder and an arithmetic use on the linked-item arm.** For zero-extended u16 A at runtime record+10, N=(X or Y nonzero) and L=(backlink+4 nonzero), R0461 first writes a cell when N and A<4, then searches an actor when not L and A&8, then searches a building when N and not L and A&4. Failed actor lookup jumps to the next record, so A=12 with an absent actor suppresses the building search. The later eligible linked arm requires not N and L and appends kind low8(A+43) only when A!=0; A=212/213/65535 produces 255/0/42 at that store. Setting bit15 cannot make the zero-extended comparison negative. The established wire frame and cell-entry route remain separate.

**Confidence.** **High** for the body-scoped predicates, order and the byte store: complete original listing, 456 original-instruction controls over 38 words/12 contexts, and 38 linked arithmetic controls discriminate independent flags and a low-bit enumeration. Native acceptance, Effect-kind validity and other record aliases remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0354](../experiments/EXP-0354-type9-lifetime/)

### ALM-T9ACTOR-176

**The A&8 path supplies an actor Spellbook entry before the A&4 building path; neither receives A as a constructor parameter.** Both lookups use full32 B plus C shifted16 from runtime record+12. Building keys are populated from the enumerated object's+8 and compared full-width throughL02074/L02075. On the actor path a missing actor+140 causesR0462 construction, actor+4c bit1 to be set, and bit2 when signed word+9c is nonzero. It constructsR0463(low8(spellRaw)) and inserts at low16(spellRaw) throughR0464; an occupied slot first receives virtual+4(1). For a constructed Spell entry, table L05044 binds this toL02076. The building constructorL02077 receives only ECX=this; its caller later copies fields.

**Confidence.** **High** for these local inputs, stores and typed entry dispatch, supported by complete bodies and explicit lookup/entry service cuts. Other existing entry classes, actor execution, allocation failures and native Spellbook lifetime remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0354](../experiments/EXP-0354-type9-lifetime/)

## ALM claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|----|-------|-----------|--------|----------|
| ALM-HDR-001 | File header (12 B): `char[4]` magic `4D 37 52 00` ("M7R␀") + `u32 version = 20` + `u32 dataSize = 4·W·H + 72` (a precomputed size; `W,H` live in the type-0 section, not the header). | High / Medium | ● active (amended) | [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| ALM-SEC-002 | The body after the header is a chain of length-prefixed sections; each section header is 20 bytes `[u32 f0][u32 f1][u32 tag = 7][u32 hdrLen = 20][u32 payloadSize]` + `payloadSize` bytes of payload. | High / Medium | ● active (amended) | [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| ALM-SEC-003 | Exactly **10 sections**; each is typed by the leading `u32` of its payload; typeIds are the fixed set `{0..9}` in fixed physical order `0,1,2,3,5,4,9,8,6,7`. | High / Medium | ● active (partially retracted) | [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0056](../experiments/EXP-0056-alm-record-set/) |
| ALM-SEC-004 | Section-size profile by type: `type0` = 632 B metadata (holds `W` @payload+8, `H` @payload+12, `π/4`, scalars, ASCII map name + description); `type1` = `2·W·H`; `type2`,`type3` = `W·H` (full-map grid layers); `type4–9` variable ... | High / Medium / Unknown | ● active (amended) | [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0056](../experiments/EXP-0056-alm-record-set/) |
| ALM-TRL-005 | An 8-byte trailer (two u32) follows the last section: **zero** in all 10 standalone root maps; small non-zero ints in the 28 embedded campaign maps (2nd u32 usually 1, 1st ∈ {0,1,2,3}). | — | ✖ retracted | [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| ALM-HDR-006 | Per-section header fields `f0`,`f1` located but not interpreted (`type0` constant `10,990`; grid layers `0,0`; other sections variable / pointer-like). | — | ✖ retracted | [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| ALM-LOC-007 | `scenario.res` is an `&YA1` archive holding **28 embedded campaign maps** (`N.alm`, e.g. `10.alm`..`151.alm`) + 3 nested `&YA1` nodes; all 28 satisfy the container model (tiling, `dataSize`, typeId order) — 38 maps total conform | High / Medium | ● active | [EXP-0007](../experiments/EXP-0007-alm-container/) |
| ALM-META-008 | type-0 (632 B) field layout, constant across 38 maps: `typeId@+0=0`, `f32 selectorA@+4`, `W@+8`, `H@+12`, `f32 angle@+10`, u32 scalars @`+14,+18,+1c,+20 (low-bit bitmask),+24,+28,+2c,+30,+34,+78,+7c`, `char[64] name@+38` ... | High / Medium / Unknown | ● active (amended) | [EXP-0010](../experiments/EXP-0010-alm-grids/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0056](../experiments/EXP-0056-alm-record-set/) |
| ALM-META-009 | The two type-0 floats: `+0x10` is an **angle** (radians, ≈`[-π/4..+π/4]`, exact `π/4` in most maps, also `π/5, ±π/10, -π/4`), plausibly a light/sun angle; `+0x04` is a **3-value discrete selector** `{-1.50133,-1.91545,-1.92320}` that ... | High / Medium | ● active (partially retracted) | [EXP-0010](../experiments/EXP-0010-alm-grids/) |
| ALM-META-010 | `+0x38` = ASCII **map name** (64-B NUL-terminated field; ≤21 B used; empty on most campaign maps; e.g | High / Medium | ● active (amended) | [EXP-0010](../experiments/EXP-0010-alm-grids/) — offsets −8 per [EXP-0030]: name `@+0x30`, description `@+0x78` (verified landing on the string starts) |
| ALM-GRID-011 | The three full-map grid sections store `W·H` cells at **payload+0** with `payloadSize` exactly `2·W·H` (type1, u16) / `W·H` (type2, u8) / `W·H` (type3, u8); the `[u32 typeId][f32]` identifier is **overlaid on the first 8 bytes** (no ... | High | ● active (partially retracted) | [EXP-0010](../experiments/EXP-0010-alm-grids/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| ALM-GRID-012 | type1 = "Tiles" | High / Medium | ● active (amended, superseded) | [EXP-0020](../experiments/EXP-0020-alm-grid-semantics/), **[EXP-0041](../experiments/EXP-0041-terrain-passability/)** |
| ALM-GRID-013 | type2 = "Altitudes" (height) — decided. | High / Medium | ● active | [EXP-0020](../experiments/EXP-0020-alm-grid-semantics/) |
| ALM-GRID-014 | type3 | High / Medium | ● active (amended) | [EXP-0020](../experiments/EXP-0020-alm-grid-semantics/), [EXP-0037](../experiments/EXP-0037-alm-class-binding/) |
| ALM-TERR-015 | Terrain vocabulary = `world.res:data/map.reg` | High | ● active (amended, superseded, partially retracted) | [EXP-0020](../experiments/EXP-0020-alm-grid-semantics/), **[EXP-0041](../experiments/EXP-0041-terrain-passability/)** |
| ALM-TERR-016 | The runtime **passability/occupancy grid** (`sim+0x10000`, 256×256, copied to `+0x20000`) is derived per cell from type1+type3 (`rom.exe R0278`): **1** = terrain-impassable (Mountain, Water, or a bit-13 tile), **5** = type3 ... | High | ● active (amended, partially retracted) | [EXP-0020](../experiments/EXP-0020-alm-grid-semantics/), **[EXP-0041](../experiments/EXP-0041-terrain-passability/)** |
| ALM-CNT-017 | The type-0 metadata is a **manifest of the fixed-record content-section counts**: `type0 payload +0x24 = #type5`, `+0x28 = #type4`, `+0x2c = #type6` (38/38 maps); the variable sections (type7/8/9) carry their own count word instead. | High | ● active (amended) | [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0037](../experiments/EXP-0037-alm-class-binding/) |
| ALM-UNIT-018 | type6 = placed units | High / Medium | ● active (amended, partially retracted) | [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0037](../experiments/EXP-0037-alm-class-binding/) |
| ALM-OBJ-019 | type4 = placed objects/structures | High / Medium | ● active (amended, partially retracted) | [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0037](../experiments/EXP-0037-alm-class-binding/), **[EXP-0081](../experiments/EXP-0081-bridge-crossing/)**, [EXP-0542](../experiments/EXP-0542-map-field-readers/) |
| ALM-GRP-020 | type5 = player/group roster | High / Medium | ● active (amended) | [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| ALM-TRIG-021 | type7 = the trigger effect/instant list | Medium | ● active (amended) | [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), **[EXP-0045](../experiments/EXP-0045-alm-triggers/)** |
| ALM-TRIG-022 | type8 / type9 = serialized marker trees | Medium / Low / Unknown | ● active (amended) | [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), **[EXP-0045](../experiments/EXP-0045-alm-triggers/)** |
| ALM-CODE-023 | `rom.exe` map-load/placement path (own Ghidra project): the `%d.alm` filename is built and loaded in `R0099` → `R0473` (`CMap` family); `R0065` reads the **groups** (type5 — `"Oops - player has no groups"`) and the ... | High / Medium | ● active | [EXP-0019](../experiments/EXP-0019-alm-content/) |
| ALM-META-024 | In `rom.exe` the `.alm` reader `R0478` reads a 20-byte file header, `strcmp`s the magic `M7R␀` (the global at `L02085`), then reads **`type0.f0` as the record count** (loop bound, gated `≥3`) and **`type0.f1` as the format ... | High | ● active (amended) | [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0056](../experiments/EXP-0056-alm-record-set/) |
| ALM-META-025 | The type-0 scalars `+0x24/+0x28/+0x2c/+0x30/+0x34` are **element counts** the loader uses as loop bounds for the content records. | High / Unknown | ● active (amended) | [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/) |
| ALM-META-026 | The type-0 `+0x20` low-bit bitmask is **read into a local and discarded** by the map loader (`local_60`, referenced nowhere else) — read-but-unused by `rom.exe`'s map parse. | High / Medium / Unknown | ● active (amended, superseded, partially retracted) | [EXP-0018](../experiments/EXP-0018-alm-metadata/) — offsets −8 per [EXP-0030]: discarded bitmask `@+0x18`, stored `+0x0c/+0x10/+0x14` and `+0x70/+0x74`, [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |
| ALM-META-027 | The type-0 float `+0x10` (`angle`, ≈±π/4) is read into the map object as a stored **light/sun angle** (corroborated by the Map-Editor Help "Show Time / Light" settings). | High / Medium | ● active (partially retracted) | [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |
| ALM-META-028 | The type-0 `+0xc0` region (**`+0xb8` per [EXP-0030]**) is a fixed **array of 7 × 64-byte text slots** (`[3×u32 prefix][char[52] text @slot+12]`, default `"<None>"`), read into the map object as the tail of the 512-byte description block. | High / Unknown | ● active (amended) | [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0056](../experiments/EXP-0056-alm-record-set/) |
| ALM-HDR-029 | For sections other than type0, the 20-byte header words `f0/f1` are **not map fields**: `0` where the writer's source buffer was zeroed (type1/4/5), and adjacent record-framing / uninitialised residue elsewhere — decisively ... | High | ✖ retracted (superseded) | [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| ALM-TRL-030 | The 8-byte file trailer is **not read by the map parser** (`R0478` reads exactly `count` records and returns; no trailing read). | — | ✖ retracted | [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| ALM-FRAME-031 | Corrected container framing (supersedes the EXP-0007 split). | High | ● active (partially retracted) | [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| ALM-GRID-032 | Grid base and cell→index (the four-cell question, answered). | High | ● active | [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| ALM-PLACE-033 | Placement anchor → terrain cell is a bare `>>8`. | High / Medium | ● active | [EXP-0030](../experiments/EXP-0030-alm-grid-origin/) |
| ALM-OBJ-034 | type-4 extension discriminator = `kind == 0x21`. | High | ● active (amended, superseded) | [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0037](../experiments/EXP-0037-alm-class-binding/) |
| ALM-CLS-035 | type3 = the static-object placement layer, and its code names a class. | High / Medium / Unknown | ● active (amended) | [EXP-0037](../experiments/EXP-0037-alm-class-binding/), **[EXP-0045](../experiments/EXP-0045-alm-triggers/)** |
| ALM-CLS-036 | type4 `kind` (file `+0x08`) is the class-selecting field, and its domain is `structures.reg`'s `ID`. | High | ● active (amended) | [EXP-0037](../experiments/EXP-0037-alm-class-binding/), **[EXP-0049](../experiments/EXP-0049-placeable-db/)** |
| ALM-CLS-037 | The type-4 object/structure discriminator is `kind ∈ {0x22, 0x23}`, read from the walker. | High / Medium | ● active (amended) | [EXP-0037](../experiments/EXP-0037-alm-class-binding/) |
| ALM-CLS-038 | type6 names its unit class by `units.reg`'s `ID`, and resolves it through a definition collection, not a registry subscript. | High | ● active (amended, partially retracted) | [EXP-0037](../experiments/EXP-0037-alm-class-binding/), **[EXP-0049](../experiments/EXP-0049-placeable-db/)** |
| ALM-OWN-039 | Owner, both record types: type-4 file `+0x0e` and type-6 file `+0x14` are the same field — a 1-based physical slot in the type-5 array. | High / Medium | ● active | [EXP-0037](../experiments/EXP-0037-alm-class-binding/) |
| ALM-UNIT-040 | The type-6 70-byte record, field by field, from the loader's own read sequence | High / Medium | ● active (amended, superseded) | [EXP-0037](../experiments/EXP-0037-alm-class-binding/) |
| ALM-GRP-041 | The type-5 76-byte record, field by field | High / Unknown | ● active (partially retracted, superseded) | [EXP-0037](../experiments/EXP-0037-alm-class-binding/), [EXP-0074](../experiments/EXP-0074-target-acquisition/), [EXP-0180](../experiments/EXP-0180-shade-select/), [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/), [EXP-0351](../experiments/EXP-0351-authored-player-control/) |
| ALM-CLS-042 | Shipped maps place `objects.reg` *fire* variants, and those classes' art does not ship. | High / Medium / Unknown | ● active (amended) | [EXP-0037](../experiments/EXP-0037-alm-class-binding/) (added 2026-07-27) |
| ALM-TERR-043 | `world.res:data/map.reg` in full — and the half of it the engine never reads. | High / Medium | ● active | [EXP-0041](../experiments/EXP-0041-terrain-passability/) |
| ALM-TRIG-044 | The type-7 payload is THREE counted arrays, not one. | Medium | ● active | [EXP-0045](../experiments/EXP-0045-alm-triggers/) |
| ALM-TRIG-045 | The 796-byte node record, field by field | Medium / Unknown | ● active | [EXP-0045](../experiments/EXP-0045-alm-triggers/) |
| ALM-TRIG-046 | The nine parameter type codes, and the identifier domain each names. | Medium / Unknown | ● active (amended) | [EXP-0045](../experiments/EXP-0045-alm-triggers/) |
| ALM-TRIG-047 | The 184-byte trigger record, field by field. | Medium / Unknown | ● active (partially retracted) | [EXP-0045](../experiments/EXP-0045-alm-triggers/) |
| ALM-UNIT-048 | The type-6 record's two id words, corrected: `+0x40` (u16) is the UNIT id and `+0x42` (u32) is the GROUP id. | Medium | ● active | [EXP-0045](../experiments/EXP-0045-alm-triggers/) |
| ALM-TRIG-049 | type-9 = `[u32 count]` + `count` records of `26 + 6n` bytes | Medium | ● active (amended, partially retracted) | [EXP-0045](../experiments/EXP-0045-alm-triggers/), [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/) |
| ALM-TRIG-050 | type-8 has NO count word | Medium / Unknown | ● active (amended) | [EXP-0045](../experiments/EXP-0045-alm-triggers/) |
| ALM-CLS-051 | `ALM-CLS-035`'s 64 residual type-3 cells are head-of-buffer residue, not placements — and the contaminated span is wider than 64 cells. | Medium / Low / Unknown | ● active | [EXP-0045](../experiments/EXP-0045-alm-triggers/) |
| ALM-CLS-052 | The type-6 → definition resolution, closed over the shipped file. | High / Medium | ● active | [EXP-0049](../experiments/EXP-0049-placeable-db/) |
| ALM-CLS-053 | The type-4 `kind` law, completed: `kind` subscripts the Buildings table of `Data.bin`, 1-based because the file skips entry 0. | High | ● active | [EXP-0049](../experiments/EXP-0049-placeable-db/) |
| ALM-CLS-054 | The client creation message's class byte becomes drawable `+0x20` at `L02143`, and its server source is `actor+0x0e`. | High / Medium | ● active (amended, partially retracted) | [EXP-0049](../experiments/EXP-0049-placeable-db/), [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |
| ALM-REQ-055 | The open path requires exactly two records: `type1` and `type2`. Nothing else, and never a count. | High | ● active | [EXP-0056](../experiments/EXP-0056-alm-record-set/) |
| ALM-REQ-056 | Per-section behaviour on absence, section by section. | High | ● active (partially retracted) | [EXP-0056](../experiments/EXP-0056-alm-record-set/), [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/) |
| ALM-ORD-057 | The physical order `0,1,2,3,5,4,9,8,6,7` is not enforced, but three precedence relations are | High | ● active (partially retracted) | [EXP-0056](../experiments/EXP-0056-alm-record-set/), [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/) |
| ALM-META-058 | type-0 payload `+0x70` (→ `map+0xd4`) is a gate operand twice over, and the `+0x78` text region is consumed as one string. | High / Medium | ● active | [EXP-0056](../experiments/EXP-0056-alm-record-set/) |
| ALM-RDR-059 | Four `.alm` readers ship, with three different acceptance tests — and only one of them acts on the loader's status word. | High / Medium / Unknown | ● active (amended) | [EXP-0056](../experiments/EXP-0056-alm-record-set/), [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |
| ALM-CORP-060 | `ru/Horror.alm` is `en/Horror.alm`'s first 262 876 bytes with two bytes changed — and it fixes what the 38-map corpus actually was. | High | ● active | [EXP-0056](../experiments/EXP-0056-alm-record-set/) |
| ALM-OBJ-061 | Which of a type-4 placement's two fixed-point coordinates is the column, read from the image end to end — `TERR-STRUCT-075`'s "the instructions do not say" is withdrawn. | High | ● active | [EXP-0091](../experiments/EXP-0091-placement-axes/) |
| ALM-OBJ-062 | The type-4 extension's eight bytes, field by field — and three quarters of them are dead. | High / Medium | ● active | [EXP-0091](../experiments/EXP-0091-placement-axes/) |
| ALM-CLS-063 | The `kind ∈ {0x22,0x23}` class DOES reach the footprint resolver — through its base-class constructor, which is why a call-site count could never show it. | High | ● active | [EXP-0091](../experiments/EXP-0091-placement-axes/) |
| ALM-LIM-064 | What a placement lets you customise, and what lifting each limit costs. | Medium | ● active | [EXP-0091](../experiments/EXP-0091-placement-axes/) |
| ALM-SACK-065 | The `.alm` type-8 section is the map's authored loot, and every field has a consuming instruction. | High | ● active | [EXP-0128](../experiments/EXP-0128-alm-items-section/) |
| ALM-SACK-066 | The corpus, and the three properties that discriminate the reading from its rivals. | High / Medium | ● active | [EXP-0128](../experiments/EXP-0128-alm-items-section/) |
| ALM-LIM-067 | What an authored-loot record lets a consumer customise, and what lifting each limit costs. | Medium | ● active (amended, partially retracted) | [EXP-0128](../experiments/EXP-0128-alm-items-section/), [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/) |
| ALM-ORD-068 | The shipped physical record order is `0,1,2,3,5,4,9,8,6,7`, and it is not a writer's habit - two forward references in the loader require it. | High / Medium | ● active | [EXP-0128](../experiments/EXP-0128-alm-items-section/) |
| ALM-PLAYER-069 | The type-5 record's `+0x00` word is the faction's colour slot and its `+0x04` word is `Player+0x28`; the word the owner lookup matches is neither, it is the record's own 1-based ordinal. | High / Medium | ● active (amended, superseded) | [EXP-0180](../experiments/EXP-0180-shade-select/), [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/) |
| ALM-MODE-070 | The type-0 payload `+0x70` scalar is the map object's `+0xd4`, and its consumer is one test: the session's multiplayer flag is `(map+0xd4 > 1)`. | High | ● active | [EXP-0181](../experiments/EXP-0181-owner-colour-rule/) |
| ALM-EFFREC-071 | Every version-990 type-9 record has one exact `26+6n` wire frame; an eligible linked zero-coordinate record separately supplies an ordered map-item Effect recipe. | Medium | ● active | [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/) |
| ALM-EFFLINK-072 | A type-8 element's type-9 link is 1-based, zero means none, and the linked operation is the same for ground and actor stock. | High / Medium | ● active (amended) | [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/) |
| ALM-EFFPOP-073 | Every positive type-8→type-9 link in the shipped campaign is enumerated: 74 per root. | Medium | ● active | [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/) |
| ALM-M40-074 | Mission 40's type-8 element at cell `(96,84)` is record 2 element 1, code `0x0e1f`: class 14, runtime row 31, `Quest Item31`, value 10,000, link zero and no linked Effect. | High / Medium | ● active | [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/) |
| ALM-TAILMAP-079 | The type-6 record's file→runtime map, re-derived: 28 stream reads, not 27, and 18 of them are not the identity. | High / Medium | ● active | [EXP-0296](../experiments/EXP-0296-unit-placement-tail/) |
| ALM-TAILHOLD-080 | Exactly three routines can hold a pointer to a type-6 runtime record, and between them they dereference 20 distinct displacements of it. | High / Medium | ● active (amended) | [EXP-0296](../experiments/EXP-0296-unit-placement-tail/) |
| ALM-TAILDIR-081 | The type-6 file `+0x1c` is a `u32` whose domain is exactly `0..14`, and no routine that can hold a record pointer reads where it lands. | Medium | ● active (amended) | [EXP-0296](../experiments/EXP-0296-unit-placement-tail/) |
| ALM-TAILU16-082 | The four type-6 `u16` at file `+0x20`…`+0x27` use `0xFFFF` as the absent value, and only two of the four are read by anything that can hold the record. | Medium | ● active (amended) | [EXP-0296](../experiments/EXP-0296-unit-placement-tail/) |
| ALM-TAILRUN-083 | `ALM-UNIT-018`'s `0xFF`-sentinel inventory reading of the type-6 byte runs is wrong on both the values and the role. | High / Medium | ● active (amended) | [EXP-0296](../experiments/EXP-0296-unit-placement-tail/) |
| ALM-TAILVER-084 | The version-gated type-6 pair at file `+0x28`/`+0x2a` lands at `rec+0x4c`/`rec+0x4e`, and no routine that can hold a record pointer reads either. | High / Medium | ● active (amended) | [EXP-0296](../experiments/EXP-0296-unit-placement-tail/) |
| ALM-SCALAR-087 | The type-5 `+0x08` word lands in `CPlayer+0x0c`; within the seven-function CPlayer family the field has one direct read, in copy constructor `R0503`, and that constructor has two direct call sites. | High / Unknown | ● active (amended, partially retracted) | [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/), [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md) |
| ALM-SCALAR-088 | A provenance-selected direct-displacement population contains 17 map-side routines plus 7 CPlayer-family routines; no map-side hit was attributed to a read of the map-loaded `CPlayer+0x0c`, while the family contains `ALM-SCALAR-087`'s ... | Medium / Unknown | ● active (amended, partially retracted) | [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/), [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md) |
| ALM-SCALAR-089 | The type-5 `+0x08` word takes exactly two values over both preserved roots, every record 0 authors `0`, and the loader's absent-section default authors `5000` into the record-0 position. | High / Unknown | ● active (amended) | [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/), [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md) |
| ALM-CPLAYER-090 | The interface `CPlayer` vtable at `L02273` is six dwords: five function entries followed by a null at `L02274`; `L02275` begins the next class. | High | ● active (amended) | [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/), [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md) |
| ALM-META-091 | Case 0's sixteen reads, and the four map-object slots the stored type-0 scalars land in. | High | ● active | [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |
| ALM-META-092 | Over the searched population nothing reads the four stored scalars back off the map object, and the one reader with no object to store them in steps over them. | Medium | ● active (amended) | [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |
| ALM-CORP-093 | What the four stored type-0 scalars carry, over 72 shipped maps on both preserved roots | High / Unknown | ● active | [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/) |
| ALM-HEADER-097 | The selected main reader and the editor reader compare file-header+0x10 as an unsigned dword: upper gate1001, and payload gates951,985,987,989,990. | High / Unknown | ● active | [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/) |
| ALM-HEADER-098 | Record-header+0x10 is four raw bytes, not an established f32 or enum. | High / Medium / Unknown | ● active | [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/) |
| ALM-WRITER-099 | Selected editor writerL02295 writes file+0x10 from literal990 atL02296, and computes header dataSize as 4*W*H+72. | High / Medium / Unknown | ● active | [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/) |
| ALM-STREAM-100 | Main/browser construct wrapperL02299 with vtableL02300; landscape callerR0514 constructs the same wrapper beforeR0480. | High / Unknown | ● active | [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/) |
| ALM-CENSUS-101 | Read-only header inventory covers38 EN and 34 RU ALMs:16 loose files plus 56 scenario archive members,714 record headers. | Medium | ● active | [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/) |
| ALM-EDSCALAR-102 | The editor readerL02303 case0 requests four bytes for payload+0x0c into its object+0x1c, through the actual direct-file class/vtableL02304 Read+34=L02305 or the fallback wrapper class. | High / Unknown | ● active | [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/) |
| ALM-FLAGPATH-109 | Placement flags have arm-specific consumers, including a conditional bit-2 write. | High / Unknown | ✔ promoted | [EXP-0347](../experiments/EXP-0347-unit-placement-flags/), byte-free instruction/decoder tables, call frontiers and synthetic vectors |
| ALM-FLAGCORP-110 | The admitted 72 maps author flags 0, 1, 4 and 5 only. | Medium | ✔ promoted | [EXP-0347](../experiments/EXP-0347-unit-placement-flags/), byte-free instruction/decoder tables, call frontiers and synthetic vectors |
| ALM-PLACESTREAM-111 | The case-6 interior destination reaches an underlying read; the previous stream escape is narrowed, not globally closed. | High / Unknown | ✔ promoted | [EXP-0347](../experiments/EXP-0347-unit-placement-flags/), byte-free instruction/decoder tables, call frontiers and synthetic vectors |
| ALM-PLACEEDITOR-112 | The EN editor reads, copies and offers the questioned Unit fields to its writer as opaque values. | High / Medium / Unknown | ✔ promoted | [EXP-0347](../experiments/EXP-0347-unit-placement-flags/), byte-free instruction/decoder tables, call frontiers and synthetic vectors |

### ALM-HDR-001

File header (12 B): `char[4]` magic `4D 37 52 00` ("M7R␀") + `u32 version = 20` + `u32 dataSize = 4·W·H + 72` (a precomputed size; `W,H` live in the type-0 section, not the header). **Amended by [EXP-0030] (`ALM-FRAME-031`): the header is 20 bytes** — the same three fields plus `u32 recordCount = 10 @+0x0C` and `u32 formatVersion = 990 @+0x10`, which EXP-0007 attributed to section 0's `f0`/`f1`. Further: **`+0x04 = 20` is the header's own LENGTH, and it is read and used** — `R0465` peeks 8 bytes at the cursor, seeks back 8, and reads `dword[cursor+4]` bytes; its first call runs with the cursor at file offset 0, so it reads `dword@+0x04` = 20 bytes as the file header (unambiguous: no sequential-cursor ambiguity at offset 0). `dataSize = 4·W·H+72` still holds 38/38 but **loses its explanation** (EXP-0007 read it as `12 + 3·20 + 4WH`; corrected that is `20 + 60 = 80`); the loader never reads the field — open

**Confidence.** High (the 20-byte header and that `+0x04` is its live length: `R0465` peeks 8 B, seeks −8 and reads `dword[cursor+4]` bytes, with the cursor at file offset 0 on the first call; the magic is `strcmp`ed — `ALM-META-024`) / **Medium** (`dataSize = 4·W·H + 72` — an unexplained arithmetic identity over the 38-map corpus with no consumer read; restored to High by a `rom.exe` site that writes or validates the field)

**Original status.** ● active (amended)

**Evidence.** [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-SEC-002

The body after the header is a chain of length-prefixed sections; each section header is 20 bytes `[u32 f0][u32 f1][u32 tag = 7][u32 hdrLen = 20][u32 payloadSize]` + `payloadSize` bytes of payload. The chain **tiles the file exactly**: `12 + Σ(20 + payloadSize) + 8 == fileSize` (0 gaps/overlaps/OOB). `hdrLen = 20` is self-describing (= header length), pinning the 20-byte stride. **Amended by [EXP-0030]: the 20-byte header is `[tag=7][hdrLen=20][payloadSize][typeId][f32]`, i.e. it starts 8 bytes later and ends with the two words EXP-0007 read as the payload's head; there are no `f0`/`f1` fields.** The tiling is the same arithmetic re-split: `20 + Σ(20+payloadSize) == fileSize`, 38/38. `hdrLen = 20` self-describes **both** splits, so it never pinned this one

**Confidence.** High (chain/stride — carried by `ALM-FRAME-031`'s `Read(hdr,0x14)` in `R0466`/`R0467`, **not** by the tiling: the exact tiling is corpus evidence and holds under both splits, i.e. Medium on its own) / **corrected** (the split)

**Original status.** ● active (amended)

**Evidence.** [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

**Amended.** The table ledger carried the status "● active (amended)". Narrowed by
`ALM-NEWMAP-237`: the tag word is `E+8`, 5 in an editor-created map; the constant 7 is the
installed value.

### ALM-SEC-003

Exactly **10 sections**; each is typed by the leading `u32` of its payload; typeIds are the fixed set `{0..9}` in fixed physical order `0,1,2,3,5,4,9,8,6,7`. Every payload begins `[u32 typeId][f32 per-map-constant]` (the f32 is byte-identical across a map's sections). **No ASCII section tags**. **Amended by [EXP-0030]: `typeId` and the `f32` are the last two words of the record HEADER, not the head of the payload** — the loader's `switch` reads header dword 3 (`[EBP-0x28]`, 10-entry jump table at `L02078`) and the payload begins immediately after. Count/order/constancy unchanged, 38/38. The "every payload begins `[typeId][f32]`" half is **impossible on shipped data**: twelve records have `payloadSize < 8` — `type8 size=0` on 8 maps, `type4 size=0` on `scn:31`, `type9 size=4` on `scn:30/31/61` — so an 8-byte payload-head identity cannot exist there. (EXP-0007's parser recorded a section only when `size >= 4`, which hid it.) **Amended by [EXP-0056]: "exactly 10" and "the fixed set `{0..9}`" are writer regularities, not a contract, and the corpus that graded them was the EN root + `scenario.res` only.** The loader's loop bound is the header's `recordCount`, gated **only `>= 3`** (`L02079`); the RU root ships a magic-valid, exactly-tiling, `dataSize`-consistent **four-record** `Horror.alm` with `recordCount = 4` and typeIds `{0,1,2,3}` — a counterexample from shipped bytes (`ALM-CORP-060`). What the open path does require is `{type1, type2}` (`ALM-REQ-055`), and the physical order is enforced only as three precedence relations (`ALM-ORD-057`). Count/order/`f32`-constancy still hold 71/72 over both roots, `ru/Horror.alm` excepted **Partially retracted by EXP-0346 (ALM-HEADER-098, ALM-WRITER-099): the record-header word has no demonstrated f32/enum interpretation, map-object numeric consumption or payload echo. Framing and unrelated metadata-angle clauses stand.**

**Confidence.** High (`typeId` = record-header dword 3 via the 10-entry jump table at `L02078`; the loop bound is the file header's `recordCount` — `ALM-META-024`) / **Medium** (the `f32`'s byte-identity across a map's records — corpus, a writer convention no instruction enforces) / **✖** (that ten records, the full set `{0..9}`, or the physical permutation are required — refuted by [EXP-0056] from code and from a shipped file)

**Original status.** ● partially retracted

**Evidence.** [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0056](../experiments/EXP-0056-alm-record-set/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### ALM-SEC-004

Section-size profile by type: `type0` = 632 B metadata (holds `W` @payload+8, `H` @payload+12, `π/4`, scalars, ASCII map name + description); `type1` = `2·W·H`; `type2`,`type3` = `W·H` (full-map grid layers); `type4–9` variable (`type7` ≈ 808/992 B, payload begins ASCII `"Drop"`; `type8` sometimes empty). Payload internals deliberately not decoded. **Amended by [EXP-0030]: the sizes are unchanged but they are sizes of PURE payload** — type0 = 632 B of fields (the loader's case-0 reads sum to exactly 632), type1/2/3 = `2WH`/`WH`/`WH` of pure grid; the type7 payload begins with its count word, and the ASCII `"Drop"` EXP-0007 saw at its head is 8 bytes further in. **Amended by [EXP-0056] on where each grid case takes its length from** — it is *not* the record's own `payloadSize` in two of the three cases: case 1 allocates and reads `this+0x00 * this+0x04 * 2` (the multiplication at `L02080` and the shift left at `L02081`), i.e. `W·H·2` from the **type-0** record, with the constructor's `16,16` as the fallback; case 2 uses its own `payloadSize` and **saves it** to the frame local at `-0x4c`; case 3 allocates and reads the frame local at `-0x4c`, the **type-2** record's size. So `type1 = 2·W·H` and `type3 = W·H` are what the loader *demands* of the writer, not what it measures — and they are why type-0-before-type-1 and type-2-before-type-3 are real requirements (`ALM-ORD-057`)

**Confidence.** High (`type0` = 632 and `type1`/`type2`/`type3` = `2WH`/`WH`/`WH` — case 0's read lengths sum to exactly 632, and the grid cases' `malloc` + `Read` lengths are the named instructions above) / **Medium** (the `type4–9` size profile — a corpus description with no consumer bound behind it; and "`type7` ≈ 808/992 B" is unparameterised: that is the range over the **10 standalone skirmish maps**, campaign missions reach ~100 KB — `ALM-TRIG-021`) / Unknown (payload contents)

**Original status.** ● active (amended)

**Evidence.** [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0056](../experiments/EXP-0056-alm-record-set/)

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-TRL-005

An 8-byte trailer (two u32) follows the last section: **zero** in all 10 standalone root maps; small non-zero ints in the 28 embedded campaign maps (2nd u32 usually 1, 1st ∈ {0,1,2,3}). Meaning unknown. **Resolved by [EXP-0018]: the map parser never reads it (see ALM-TRL-030)**. **RETRACTED by [EXP-0030]: there is no trailer.** Those 8 bytes are the last 8 bytes of the **type-7 record's payload**; the parser "never reads them" only in the sense that it hands the whole payload to the trigger-tree parser. The file ends exactly at `20 + Σ(20+payloadSize)`, 38/38

**Confidence.** ✖

**Original status.** ✖ retracted

**Evidence.** [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

### ALM-HDR-006

Per-section header fields `f0`,`f1` located but not interpreted (`type0` constant `10,990`; grid layers `0,0`; other sections variable / pointer-like). **Resolved by [EXP-0018]: type0.f0 = record count, type0.f1 = version (ALM-META-024); other sections' f0/f1 are framing residue (ALM-HDR-029)**. **RETRACTED by [EXP-0030]: no such fields exist.** `type0`'s pair is the **file header's** count+version; every other section's pair is the **last 8 bytes of the preceding record's payload** (e.g. type2's "f0/f1" are the last four tile words of the type1 grid — 152/152 structurally valid tile words across 38 maps)

**Confidence.** ✖

**Original status.** ✖ retracted

**Evidence.** [EXP-0007](../experiments/EXP-0007-alm-container/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

### ALM-LOC-007

`scenario.res` is an `&YA1` archive holding **28 embedded campaign maps** (`N.alm`, e.g. `10.alm`..`151.alm`) + 3 nested `&YA1` nodes; all 28 satisfy the container model (tiling, `dataSize`, typeId order) — 38 maps total conform

**Confidence.** High (the roster is an **enumeration, not a fit**: the 28 `N.alm` nodes are read out of `scenario.res`'s own directory, and the only alternative — a different node set — is excluded by the listing itself; corpus = the GOG install, 10 root maps + 28 embedded = 38) / **Medium** (that all 38 "satisfy the container model" — the model quoted here is the *pre-EXP-0030* one; the identity was re-verified on the corrected framing as `20 + Σ(20+payloadSize)`, 380 record headers, 0 violations — `ALM-FRAME-031`(f))

**Original status.** ● active

**Evidence.** [EXP-0007](../experiments/EXP-0007-alm-container/)

### ALM-META-008

type-0 (632 B) field layout, constant across 38 maps: `typeId@+0=0`, `f32 selectorA@+4`, `W@+8`, `H@+12`, `f32 angle@+10`, u32 scalars @`+14,+18,+1c,+20 (low-bit bitmask),+24,+28,+2c,+30,+34,+78,+7c`, `char[64] name@+38`, `char[64] description@+80`, and a 440-B trailing region of default-empty 64-B text slots @`+c0`. Only typeId is corpus-constant; every other scalar is per-map. **Meanings resolved by [EXP-0018]: ALM-META-025 (count fields +24/+28/+2c/+30/+34), ALM-META-026 (+20 discarded, +14/+18/+1c/+78/+7c stored), ALM-META-028 (+c0 text-slot array)**. **Amended by [EXP-0030]: every offset in this row is −8** (`typeId`/`selectorA` are header words, not payload) — `W@+0x00`, `H@+0x04`, `angle@+0x08`, scalars `+0x0c/+0x10/+0x14/+0x18(discarded)/+0x1c/+0x20/+0x24/+0x28/+0x2c`, `name@+0x30`, `+0x70/+0x74`, `description@+0x78`, text slots `@+0xb8`. Re-derived independently from the loader: case 0's read lengths sum to exactly **632**. **[EXP-0056] publishes each of the sixteen reads' destination**, which sharpens the row in three places: the description and the trailing text region are **one 0x200 read** into `map+0xdc`, not two; `+0x08/+0x0c/+0x10/+0x14` land in the map object at `this+0x18/+0x1c/+0x20/+0x24` (this row's `angle` and first three scalars) while `+0x1c/+0x20/+0x24/+0x28/+0x2c` land in **stack slots that are nothing but the case-5/4/6/7/8 loop bounds** — confirming `+0x18`'s discard from the other side, since it is the one scalar that reaches neither; and `+0x70`/`+0x74` → `map+0xd4`/`+0xd8`, of which `+0x70` has two consumers (`ALM-META-058`)

**Confidence.** High (layout — case 0 issues *sequential* `Read`s whose lengths sum to exactly 632, so every field boundary quoted here is a read boundary) / **Medium** (that only `typeId` is corpus-constant and every other scalar per-map — corpus 38/38) / Unknown → mostly resolved (meanings; see the per-field rows)

**Original status.** ● active (amended)

**Evidence.** [EXP-0010](../experiments/EXP-0010-alm-grids/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0056](../experiments/EXP-0056-alm-record-set/)

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-META-009

The two type-0 floats: `+0x10` is an **angle** (radians, ≈`[-π/4..+π/4]`, exact `π/4` in most maps, also `π/5, ±π/10, -π/4`), plausibly a light/sun angle; `+0x04` is a **3-value discrete selector** `{-1.50133,-1.91545,-1.92320}` that the loader echoes into every section's payload+4 (the EXP-0007 `perMapConst`). **Refined by [EXP-0018] (ALM-META-027): `+0x10` is read into the map object as a stored light/sun angle; `selectorA` is the per-record `perMapConst`** **Partially retracted by EXP-0346 (ALM-HEADER-098, ALM-WRITER-099): the record-header word has no demonstrated f32/enum interpretation, map-object numeric consumption or payload echo. Framing and unrelated metadata-angle clauses stand.**

**Confidence.** High (both fields exist and are read into the map object — `ALM-META-027`) / **Medium** (that `+0x08` is an *angle*: supported by its value set (`π/4`, `π/5`, `±π/10`) and by the same field being consumed as `tan θ` in `R0468` — `TERR-LIGHT-028` — but no instruction names it a sun angle) / **Medium** (`selectorA`'s 3-value domain — corpus 38/38)

**Original status.** ● partially retracted

**Evidence.** [EXP-0010](../experiments/EXP-0010-alm-grids/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### ALM-META-010

`+0x38` = ASCII **map name** (64-B NUL-terminated field; ≤21 B used; empty on most campaign maps; e.g. `Beast  Land`, `Crossroads of Mystery`, `Heroes' Tomb`); `+0x80` = **description** (64-B NUL-terminated; ASCII or Windows-1251; ≤36 B used)

**Confidence.** High (two 64-B NUL-terminated `char` fields at those offsets — case-0 read boundaries, and the corrected offsets were re-verified landing on the string starts, 38/38; a wrong offset yields mid-string garbage, so the evidence does discriminate) / **Medium** (the *labels* "name" and "description" — inferred from content and length; no consumer read distinguishes the two fields)

**Original status.** ● active (amended)

**Evidence.** [EXP-0010](../experiments/EXP-0010-alm-grids/) — offsets −8 per [EXP-0030]: name `@+0x30`, description `@+0x78` (verified landing on the string starts)

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-GRID-011

The three full-map grid sections store `W·H` cells at **payload+0** with `payloadSize` exactly `2·W·H` (type1, u16) / `W·H` (type2, u8) / `W·H` (type3, u8); the `[u32 typeId][f32]` identifier is **overlaid on the first 8 bytes** (no separate prefix — this is what makes `dataSize = 4·W·H + 72` exact). Cell count exact **38/38**. **PARTIALLY RETRACTED by [EXP-0030] — see `ALM-GRID-032`: there is no overlay.** The cell counts and `payloadSize` identities stand, but the grid starts at the **record payload**, which is 8 bytes later than the payload EXP-0007/0010 delimited: cell `(col,row)` = element `row·W+col` there = **`+4` cells (type1) / `+8` cells (type2/type3)** relative to this row's "payload+0". No cell is clobbered and none is lost; the four words this row called an overlay are header fields

**Confidence.** High (sizes/counts) / ✖ (the "payload+0" base and the overlay)

**Original status.** ✖ partially retracted

**Evidence.** [EXP-0010](../experiments/EXP-0010-alm-grids/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

**Amended.** The table ledger carried the status "✖ partially retracted". `retracted.md` records a correction against this claim.

### ALM-GRID-012

**type1 = "Tiles"** (`rom.exe` error "Tiles block not found"): a per-cell 16-bit **tile-index word** — **bits 0–9 = tile index**, **bit 13 (`0x2000`) = impassable flag**, bits 10–12/14–15 unused (**re-run by [EXP-0041] on the corrected EXP-0030 base: 0 of 880 704 cells set any of them, 0 exclusions**). The terrain **class is derived** from the index by the loader (`rom.exe R0469`): strip group (bits 6–9) → hardcoded (primary,secondary) terrain-pair (`world+0x54156`); blend variant (bits 0–5) → primary/secondary pick + 5-level passability blend; index range `[512,768)` = Water. Enum (1-based in lookup; = `map.reg` `Terrain` record order): Land,Grass,Flowers,Sand,Cracked,Stones,Savanna,Mountain,Water,Road. *(Supersedes the EXP-0010 low/high byte split: the "coherent high byte" was the strip-group bits.)*

**Confidence.** High (index→class and the impassable bit: `R0469` masks `word & 0x3ff` and indexes the strip-pair table at `world+0x54156`; `R0278` turns bit 13 into block value 1) / **Medium, and the stale flag is now cleared** (the bit-population census was measured on the **pre-EXP-0030 grid base** under the "overlay corner excluded" carve-out; **[EXP-0041] regenerated it** over all 880 704 cells with 0 exclusions — bits 10–12 and 14–15: **0**, bit 13: **7 464** (0.847 %), which reproduces the old 0.85 % — so the figures were never wrong, only unverified) / **and the "impassable flag" gloss is corrected by [EXP-0041]**: bit 13 is **one of three** tile-word arms that block a cell, and the smallest of them — 7 464 cells against 136 622 for terrain class 8 and 87 584 for the raw water test `(w & 0x300) == 0x200`, pairwise overlaps 860/370/0 (`TERR-PASS-049`)

**Original status.** ● active (amended)

**Evidence.** [EXP-0020](../experiments/EXP-0020-alm-grid-semantics/), **[EXP-0041](../experiments/EXP-0041-terrain-passability/)**

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-GRID-013

**type2 = "Altitudes" (height) — decided.** Named by the `rom.exe` loader's "Altitudes block not found" error; the ingest (`R0278`) copies it into the sim **height buffer** (`this+0x9451c`). Corpus cross-ref clinches it: resolving every cell's terrain (via the `rom.exe` tables) and averaging type2 gives **Mountain highest (55.1, SD 37.1) and Water lowest (35.2, SD 20.0)**, monotone across 9 terrain types on 38 maps — a height field, **not** illumination (the metadata `+0x10` `π/4` is a separate sun angle; hill blobs are radial, undirected — EXP-0010)

**Confidence.** High (that this block is the height source: the loader's own error string "Altitudes block not found" names it, and `R0278` copies it into the sim height buffer `this+0x9451c`) / **Medium ⚠ stale base** (the corpus cross-ref *figures* — Mountain 55.1 SD 37.1, Water 35.2 SD 20.0, 38 maps — were computed on the **pre-EXP-0030 base**, where type1 and type2 are relatively misaligned by 4 cells, and were **not regenerated**. The *direction* survives on independent evidence: EXP-0030's η² of type-2 height explained by type-1 strip group peaks at the corrected alignment, 0.335 vs 0.133 at the old base, 38 maps pooled) / **⚠ the height *range* published elsewhere as "0..~246" is refuted**, not merely stale: see `TERR-LIGHT-016` — every map's old-base maximum is the third byte of its own `selectorA` `f32`, and `TERR-LIGHT-028`(d) finds 0 of 880 704 height bytes ≥ `0x80`

**Original status.** ● active

**Evidence.** [EXP-0020](../experiments/EXP-0020-alm-grid-semantics/)

### ALM-GRID-014

**type3** is a **sparse object/feature occupancy overlay** (0 = empty; 86–98% zero in 38/38; 98 distinct codes 1..246). The optional 4th grid (`rom.exe` parser stores it to `mapObj[4]`, zero-fills if absent); the ingest marks every **nonzero** cell as **object-occupied** — runtime block value **5** (vs terrain-block **1**, border **0x1f**). Only nonzero-ness is read on the movement path; the specific code (structured, step-2/3 runs; several near-universal) identifies the placed object/feature but its exact code→registry mapping is not pinned here. **RESOLVED and RE-LABELLED by [EXP-0037] (`ALM-CLS-035`): the code is `objects.reg` section index + 1**, read by the renderer at three sites. This layer is the map's **static-object placement layer**, and its movement effect is *derived* from it; "occupancy overlay" was the consequence mistaken for the thing. The stale census was **re-run on the corrected base and reproduces exactly** — 86.19–97.94 % zero, 98 distinct codes 1..246 over 71 099 nonzero cells — so the `246`-is-a-header-byte suspicion below is **wrong**

**Confidence.** High (blocks movement — `R0278` writes block value 5 for any nonzero cell, and only nonzero-ness is tested there) / Medium (overlay role — superseded by `ALM-CLS-035`'s placement-layer reading) / **Medium** (the census, now re-derived on the corrected base by [EXP-0037]; the stale-base flag and the `246` suspicion are withdrawn — the figures were never wrong, only unverified) / ~~Low (exact code→object identity)~~ → **see `ALM-CLS-035`**

**Original status.** ● active (amended)

**Evidence.** [EXP-0020](../experiments/EXP-0020-alm-grid-semantics/), [EXP-0037](../experiments/EXP-0037-alm-class-binding/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### ALM-TERR-015

**Terrain vocabulary = `world.res:data/map.reg`** `Terrain` section: **10 named classes** in fixed record order — Land,Grass,Flowers,Sand,Cracked,Stones,Savanna,Mountain,Water,Road — each with a `Cost` and a `Pass` scalar. `rom.exe` (`R0116`) opens `World\Data\map.reg`, reads the ten scalars into consecutive bytes at `world+0x54177` indexed by the enum, and blends them across tile transitions. **[EXP-0041] corrects what those ten scalars are and what the blend does**: they are the ten `Cost*` keys, not `Pass`; the table is `world+0x54176+class` with `[0x54176] = 0xff`; the blend is a 5-level interpolation between the strip group's two terrain classes (`TERR-PASS-050`); and it produces a **movement cost**, never a passability decision — the ten `Pass*` scalars have **no string and no reader in `rom.exe` or `Map Editor.exe`** and the impassable set is hardcoded (`ALM-TERR-043`). **⚠ The published figures `Cost = [0,0,0,0,0,0,0,1,1,1]` and `Pass = [8,8,14,6,12,8,16,8,6,7]` were WITHDRAWN by the 2026-07-25 sweep and are now RESTORED by [EXP-0041]: `Cost = [8,8,8,14,6,12,8,16,8,6]`, `Pass = [0,0,0,0,0,0,0,1,1,1]`.** The withdrawal is not merely replaced but *explained* — pairing name(record i) with value(record i+1), which is exactly the pre-EXP-0032 framing, regenerates both withdrawn arrays byte for byte from the same 31 records, trailing `7` included (it is `ScanShift`, the record that follows `Terrain`). They were read with our own `.reg` tooling in EXP-0020, i.e. on the **pre-EXP-0032 record framing**, under which every key was paired with the wrong record's value (`REG-REC-032`, `REG-VAL-030`'s withdrawal). They were never regenerated: EXP-0033 recorded `data/map.reg` as having "no published per-key value table before EXP-0032" and left it out of frame — **that is wrong, this row is such a table**, published in the ALM ledger rather than the REG one, which is how the base move passed under it. The *names* and their order are unaffected (they come from the name pool, and `rom.exe` reads exactly ten in this order). That re-walk is done

**Confidence.** High (the 10-class vocabulary, its order, and the `world+0x54177` table — `R0116`) / High as restored by [EXP-0041] (the values are the record array's own bytes on the framing `REG-REC-032` pins, and the old figures regenerate from them under the old pairing — which is what makes this a restoration rather than a second guess)

**Original status.** ● active (amended)

**Evidence.** [EXP-0020](../experiments/EXP-0020-alm-grid-semantics/), **[EXP-0041](../experiments/EXP-0041-terrain-passability/)**

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-TERR-016

The runtime **passability/occupancy grid** (`sim+0x10000`, 256×256, copied to `+0x20000`) is derived per cell from type1+type3 (`rom.exe R0278`): **1** = terrain-impassable (Mountain, Water, or a bit-13 tile), **5** = type3 object-occupied, **0x1f** = the 8-cell map border (`R0470`); the type2 Altitudes grid is copied to the height buffer. **Amended by [EXP-0041] on two points.** (a) There are **three** derived planes, not one: this block plane, a **movement-cost** plane at `sim+0x00000` pre-filled `0x01` and overwritten for every cell, and the height plane at `sim+0x9451c` — so "all three grids feed one derived passability map" is wrong, they feed two derived planes plus a copy. (b) **These values are a bitmask, not an enum.** The consumer is a nonzero test of the grid byte against a **per-mover** mask byte (`0x41` / `0x44` / `0x82`, `R0471`), and `1 = bit0`, `5 = bit0\|bit2`, `0x1f = bits 0..4`. Read as enum values, none of the three masks equals any of them, so no mover would ever be blocked anywhere; read as bits, a ground mover is blocked on 430 898 of 880 704 shipped cells and an air mover on exactly the 158 976 border cells (`TERR-PASS-051`)

**Confidence.** High (mechanism — the three block values are immediates written by `R0278`, and the 8-cell border by `R0470`; **re-read at instruction level by [EXP-0041]**, which adds the cost plane, the assignment-not-OR form of the four writes and the border's two loop shapes) / **~~Medium (per-value edge cases)~~ → superseded**: the values are bits and their consumer is `TERR-PASS-051`

**Original status.** ● active (amended)

**Evidence.** [EXP-0020](../experiments/EXP-0020-alm-grid-semantics/), **[EXP-0041](../experiments/EXP-0041-terrain-passability/)**

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-CNT-017

The type-0 metadata is a **manifest of the fixed-record content-section counts**: `type0 payload +0x24 = #type5`, `+0x28 = #type4`, `+0x2c = #type6` (38/38 maps); the variable sections (type7/8/9) carry their own count word instead. Independent second proof of the 76/20/70 strides; explains the EXP-0010 "count-like" `+0x28`/`+0x2c`. **Amended by [EXP-0030]: the offsets are −8** — `+0x1c` = #type5, `+0x20` = #type4, `+0x24` = #type6 (re-verified 38/38 at the corrected offsets)

**Confidence.** High (`+0x1c` = #type5 and `+0x24` = #type6 — the loader uses these two as its case-5/case-6 loop bounds, and the stride identities hold 38/38) / **High as amended by [EXP-0037]** (`+0x20` = #type4 — it *is* case 4's loop bound: the value is read into the frame local at offset -0x10 by case 0's ninth read and case 4's loop counter is compared against that local at `L02082`, so this is no longer a corpus identity only; `ALM-META-025`'s Medium on the same field is superseded)

**Original status.** ● active (amended)

**Evidence.** [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0037](../experiments/EXP-0037-alm-class-binding/)

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-UNIT-018

**type6 = placed units**: 70-B fixed records, count `= meta+0x2c`; `X @ rec+8` / `Y @ rec+12` (`u32`, fixed-point `/256`, low byte usually `0x80` = tile centre); records carry `0xFF` sentinel words (empty inventory-slot pattern). Verified: `70·count == payloadSize` **38/38** and all **8094/8094** records land inside `[0,W)×[0,H)` — a wrong stride/offset would scatter coordinates OOB. **Amended by [EXP-0030]: `X @ rec+0x00`, `Y @ rec+0x04`, count `= meta+0x24`** (offsets −8; the file bytes read are identical). The loader allocates a `0x50`-byte unit struct and reads `X`,`Y` into its `+0/+4`; the reads total exactly 70 B at version 990 **Amended by [EXP-0037]: the non-coordinate fields are no longer Medium** — the whole 70-byte read map, the class keys (`+0x08`/`+0x0a`/`+0x10`), the owner (`+0x14`), the `meta+0x2c`-bounded `+0x18` and the unique id `+0x42` are published as `ALM-UNIT-040` / `ALM-CLS-038` / `ALM-OWN-039`; the `0xFF` sentinel runs are the two reads at file `+0x35` (6 B) and `+0x3b` (5 B) **Refuted by [EXP-0296]: there is no `0xFF` sentinel anywhere in this record.** No byte of file `+0x2c`…`+0x3f` is `0xFF` on any of the 12 085 records of the two preserved roots; the only sentinels are the four `0xFFFF` `u16` at `+0x20`…`+0x27`, and the two byte runs are the skill and elemental-resistance override vectors (`ALM-TAILRUN-083`, `UNIT-PLACESKILL-086`, `UNIT-PLACERESIST-087`)

**Confidence.** High (layout/tiling/xy — case 6 reads exactly 70 B per record and stores the first two dwords into the unit struct's `+0x00`/`+0x04`; `70·count == payloadSize` 38/38 and 8094/8094 anchors inside `[0,W)×[0,H)`. The EXP-0030 relabelling moved the record base and the field offsets by the same 8 bytes, so the *file* bytes behind both figures are unchanged) / Medium (unit label)

**Original status.** ● active (amended)

**Evidence.** [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0037](../experiments/EXP-0037-alm-class-binding/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-OBJ-019

**type4 = placed objects/structures**: 20-B base record (`+0x00 u16`≈100, `+0x02 u16` kind/owner, `+0x06 u16` incrementing id, `+0x08 X`, `+0x0c Y` (`u32`/256), `+0x10 value`), count `= meta+0x28`; a minority of objects carry an extra 8 B = a second `[u32 /256 coord][u32 value]` pair. An adaptive base-20+extension walk consumes each payload **exactly** (36/37 maps; scn:81's trailing 8 B = last-record extension); base-20 alone tiles 34/37 (the 3 with in-stream extensions are Cross/Horror/scn:81). **Amended by [EXP-0030]: offsets −8 and the field order is read from the loader** — `X @ rec+0x00`, `Y @ rec+0x04`, `kind @ rec+0x08` (u32), `+0x0c` u16, `+0x0e` u32, `+0x12` u16 (the loader reads 4,4,4,2,4,2 = 20 B); count `= meta+0x20`. The **extension discriminator is `kind == 0x21`** (`ALM-OBJ-034`), with which the walk tiles **38/38** — Cross/Horror were mis-walks, not anomalies (**the earlier 36/37 and 34/37 tiling figures in this row are superseded and should not be quoted**) **Amended by [EXP-0037]: the "objects/structures" label is wrong and the fields are named.** Every type-4 record resolves against **`structures.reg` alone** — `kind` is that registry's `ID` (`ALM-CLS-036`) — and the object/structure split is between two C++ classes chosen by `kind ∈ {0x22,0x23}` (`ALM-CLS-037`), not between two registries. Field roles: `+0x0c` = a `×1000` scalar the structure arm stores at `obj+0x70` — **not a durability: [EXP-0057] reads its consumer and it is the shop's stock value cap** (`SHOP-CAP-004`), `+0x0e` = the **owner** slot (`ALM-OWN-039`), `+0x12` = sign-extended into `obj+0x10` (**role answered by [EXP-0045]: it is the id a type-7 `Target_Structure` parameter names — 16/16 references land in this field's value set, `ALM-TRIG-046`**; **not** the per-map incrementing id this row guessed — sequential on only 5/38 maps; **the DESTINATION `obj+0x10` is CONTESTED by [EXP-0081] — `registry.md` C-7**: that field is a pointer to a 12-byte position object, written by the base actor constructor at `L02083` and dereferenced by both footprint routines, so a sign-extended `i16` cannot live there. The field's *role* is untouched; where it is stored is now unsourced on both sides), extension = the footprint override `w×h` in tiles — **and also [EXP-0081] names the class it belongs to: `kind == 0x21` is class id 33, `Vertical Wooden Bridge`, whose whole rectangle is opened rather than blocked (`TERR-STRUCT-077`)**

**Confidence.** High (base layout/count/tiling — the loader reads `4,4,4,2,4,2` = 20 B and branches on `rec+0x08 == 0x21`; with that rule the walk consumes every payload exactly, 38/38) / ✖ (the "objects *or* structures" label) / **Medium, was unstated** (the `+0x12` destination, which is refuted: see Amended)

**Original status.** ● active (amended, contested)

**Evidence.** [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0037](../experiments/EXP-0037-alm-class-binding/), **[EXP-0081](../experiments/EXP-0081-bridge-crossing/)**

**Amended.** The table ledger carried the status "● active (amended, contested)". `retracted.md` records a correction against this claim. `ALM-STRUCTREAD-228` reads the `+0x12` destination: the walker stores it at `obj+0x08` (`L14613`), not `obj+0x10`. The destination clause "sign-extended into `obj+0x10`" is refuted in `retracted.md`, which closes registry contest C-7; the field's role (`ALM-TRIG-046`) stands.

### ALM-GRP-020

**type5 = player/group roster**: 76-B fixed records, count `= meta+0x24` (38/38 exact tiling); each record has a NUL-terminated ASCII **name** at `rec+0x14` — observed `Self, Monsters, Villagers, Neutral, Enemy, Beasts, Guards, Peasants, Brigands, Goblins, Orcs, Trolls, Ogres, Merchant, Paladin, …` (factions/groups) — and a scalar `{0, 5000}` at `rec+0x10`. **Amended by [EXP-0030]: offsets −8** — name `@rec+0x0c`, scalar `@rec+0x08`, count `= meta+0x1c` (re-verified 38/38) **Amended by [EXP-0037] (`ALM-GRP-041`): the whole record is read from the loader** — `+0x00` is the id word the owner lookup compares, the name field is exactly `0x20` bytes, and `+0x2c` holds **sixteen u16**, one per editor player slot

**Confidence.** High (stride/count — `76·count == payloadSize` 38/38 with `count` the loader's own case-5 bound; and the name field, because a NUL-terminated ASCII faction name begins at `rec+0x0c` in every record of all 38 maps, which a wrong offset or stride would not produce; the in-record offsets are now read boundaries — `ALM-GRP-041`) / Medium (per-field diplomacy semantics)

**Original status.** ● active (amended)

**Evidence.** [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-TRIG-021

**type7 = the trigger effect/instant list** (the "THEN" actions), incl. the **Drop table**: `[u32 typeId=7][f32 perMapConst][u32 entryCount][named-node blocks]`, `entryCount @ +8` (Beast/Cross=1, scn:51=10, scn:120=24, scn:10=28, scn:150=45 — not the string length). Each block is a named editor node (name string + id + parameter values + inline parameter-name strings + `"<None>"` target sentinels) matching `Description Instants.ini`; every map's default first entry is `"Drop Location"` — **corrected by [EXP-0030]: false on 5/38 maps** (`scn:10` = `Send message "mission starts"`, `scn:110` = `StartPoint`, `scn:111` = `MessStart`, `scn:120` = `StartLoc`, `scn:91` = `Phase out unit`), so "first entry" is not a structural invariant, only a skirmish-map default. Minimal (808/992 B — the two values seen across the 10 standalone skirmish maps) and up to ~100 KB on the 28 scripted campaign missions. **Amended by [EXP-0030]: the payload is `[u32 entryCount @ +0][named-node blocks]`** — `[typeId][f32]` are record-header words, not payload, so `entryCount` is at payload `+0`. The last 8 bytes of this payload are what EXP-0007 called the file's 8-byte trailer **Amended by [EXP-0045]: the payload is not one list but THREE counted arrays** — `[u32 nAct][nAct x 796]` `[u32 nCond][nCond x 796]` `[u32 nTrg][nTrg x 184]` (`ALM-TRIG-044`). This row's `entryCount` is the FIRST of the three, which is why its figures reproduce (Beast/Cross 1, `scn:51` 10, `scn:120` 24, `scn:10` 28, `scn:150` 45); and the "last 8 bytes of this payload" that `ALM-TRL-005` called a file trailer are the second and third arrays' count words, zero on the skirmish maps because those carry no conditions and no triggers. The "named-node block" is the fixed 796-byte record of `ALM-TRIG-045`; the "`<None>` target sentinels" are unused parameter-name slots

**Confidence.** Medium (structure + editor vocabulary + rom.exe match) / ~~Low (full serialized leaf grammar)~~ → **superseded by `ALM-TRIG-044…047`, at Medium: the grammar closes over the corpus but no consuming instruction was read**

**Original status.** ● active (amended)

**Evidence.** [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), **[EXP-0045](../experiments/EXP-0045-alm-triggers/)**

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-TRIG-022

**type8 / type9 = serialized marker trees**. type8 holds `0x80`-centred `X/Y` box/circle **marker regions** (the geometry behind the "Is unit in a box / circle" checks), empty on most skirmish maps, present on campaign missions. type9 is a `[u32 count @ +8][records]` list of **tile-coordinate markers** with a recurring `0x2376` word. Both are read by the same generic named-node parser as type7. **Amended by [EXP-0030]: type9's count is at payload `+0`** (offsets −8). type8's emptiness also explains `ALM-HDR-029`: when type8's `payloadSize` is 0, the bytes that follow its header are the next record's header, which is why type6's apparent `f1` is the map's `selectorA` float **Amended by [EXP-0045]: neither is read by a named-node parser and type8 has no count word at all.** type9 = `[u32 count]` + `count x (26 + 6n)` (`ALM-TRIG-049`); type8 = `meta+0x2c` records of `20 + 10n`, no leading count (`ALM-TRIG-050`); both walks consume every payload exactly, 38/38. The `0x2376` word this row noticed is the third `u16` of a type-9 tail element. The "same generic named-node parser as type7" clause is **withdrawn** — it came from `ALM-CODE-023`, which located the accessors rather than reading them, and neither record shape has a name field **The box/circle ROLE clause is refuted by [EXP-0080] (retracted.md):** the runtime checks that test a box and a radius read the check node's own plain parameter slots, and the only consumer of type-8/type-9 anywhere in the image is `R0461`, the spellbook / building-caster builder. The record shapes in `ALM-TRIG-049`/`ALM-TRIG-050` are untouched.

**Confidence.** **The EXP-0080 role clause is itself half wrong, by [EXP-0128]:** `R0461` is indeed the only consumer of either section, but it has **two** loops - the first over `map+0x2f0`, the type-9 list, which is where the two `spellbook` / `building - caster` strings live, and the second over `map+0x2dc`, the **type-8** list, which builds item containers and calls the sack maker `R0472` at `L02084`. type-8 is the map's **authored loot** (`ALM-SACK-065`); only type-9 is the caster payload. Low/Medium (roles + leading structure) / **Medium** (the two record shapes — each is the unique model in its `(head, nOff, elem)` family that closes 38/38, but no consuming instruction was read) / Unknown (every header field except the coordinates)

**Original status.** ● active (amended)

**Evidence.** [EXP-0019](../experiments/EXP-0019-alm-content/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), **[EXP-0045](../experiments/EXP-0045-alm-triggers/)**

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### ALM-CODE-023

`rom.exe` map-load/placement path (own Ghidra project): the `%d.alm` filename is built and loaded in `R0099` → `R0473` (`CMap` family); `R0065` reads the **groups** (type5 — `"Oops - player has no groups"`) and the **drop location** (type7 — `"Warning - no drop location in .alm - random used"`); `R0474` reads the **Outpost/repopper** effects (type7). All three access their data through a generic named-node accessor family (`R0475` find-by-name, `R0476` get-parameter, `R0477`/`IsEmpty` iterate), confirming type7–9 are a serialized tree of named nodes with typed parameters. ~~The M7R magic is not string-compared in the exe~~ — **this sentence is withdrawn: it contradicts `ALM-META-024`, which names the `strcmp` and its operand the global at `L02085` inside `R0478`.** Nothing else in the row depends on it

**Confidence.** High (the strings and their call sites — located by xref, a direct observation about the binary) / **Medium** (that the accessor family *establishes* the type7–9 node/parameter grammar — the three accessors were located, not read at instruction level, and locating a function is not establishing a claim from it) / Medium (field-by-field mapping)

**Original status.** ● active

**Evidence.** [EXP-0019](../experiments/EXP-0019-alm-content/)

### ALM-META-024

In `rom.exe` the `.alm` reader `R0478` reads a 20-byte file header, `strcmp`s the magic `M7R␀` (the global at `L02085`), then reads **`type0.f0` as the record count** (loop bound, gated `≥3`) and **`type0.f1` as the format version** (gated `≤1001`, used in version-conditional field branches at 0x3b6/0x3d8/0x3da/0x3dd/0x3e9). All 38 maps: count=10, version=990. Header `@0x04=20` and `@0x08=dataSize` are read-but-**not validated** by the loader (only the magic is). **Amended by [EXP-0030]: the count and version are FILE-header fields at `+0x0C`/`+0x10`, not "type0.f0/f1"** — `R0466` does `Read(hdr,0x14)` and the gates read `[EBP-0x3c]`/`[EBP-0x38]` inside that 20-byte buffer. Each record then gets its own `Read(hdr,0x14)` via `R0467` (skipped iff version == 1000), and the `switch` dispatches on that header's **dword 3** = `typeId`. **Amended by [EXP-0056]: the gates' failure codes have names in the image, and the code is not what the caller tests.** `[map+0x8c]` is initialised to `1` and set to `2`/`3`/`4` by the magic/count/version gates and to `5`/`6` after the loop; `R0128` maps them to `"File not found"`, `"Not a map file"`, `"Wrong block number"`, `"Map version too new (update loader!)"`, `"Tiles block not found"`, `"Altitudes block not found"` and a `"Unknown error"` default — but its own accept test is `[map+0x14] != 0`, the type-2 grid pointer, and it reads `+0x8c` only to pick the message. So the count gate is a `>= 3` **minimum**, not a `== 10` check (`ALM-REQ-055`), and two of the three call sites test neither (`ALM-RDR-059`)

**Confidence.** High (every clause tied to a named read, gate or stack slot; the corpus figures count=10/version=990 are what those gates see, not the evidence for them)

**Original status.** ● active (amended)

**Evidence.** [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0056](../experiments/EXP-0056-alm-record-set/)

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-META-025

The type-0 scalars `+0x24/+0x28/+0x2c/+0x30/+0x34` are **element counts** the loader uses as loop bounds for the content records. Corpus-universal size-identities: **`+0x24` = #player records (type5, 76 B each, 38/38)**, **`+0x2c` = #unit records (type6, 70 B each, 38/38)**, `+0x28` = #object records (type4, ≈20 B), `+0x34` = #type8 records (`0 ⟺ type8 empty`, explaining EXP-0007's empty type8), `+0x30` = #type7 records. **Amended by [EXP-0030]: offsets −8** — `+0x1c`=#type5, `+0x20`=#type4, `+0x24`=#type6, `+0x28`=(claimed #type7), `+0x2c`=#type8. **Also corrected by [EXP-0030]: `+0x28` is NOT the case-7 loop bound and is not the type-7 record count.** Case 7 (`L02086`) does `Read(&[EBP-0x50], 4)` — it reads its own count **from the stream** into the very local the metadata value was parked in, and loops on that; the two differ on **32/38** maps (`scn:10` 57 vs 28, `scn:100` 97 vs 36). Case 9 likewise reads its own stream count; only case 8's bound is the metadata value (`+0x2c`). So `+0x28` is read-and-discarded, like the `+0x18` bitmask — its meaning is **Unknown** **Amended by [EXP-0037]: `+0x20` is High** — it is the local `[EBP-0x10]` that case 4 loops on (`L02082`) — and **`+0x2c` does double duty**: besides bounding case 8 it is the range check on every type-6 record's `+0x18` field (`ALM-UNIT-040`) **Narrowed by EXP-0358 (ALM-COUNT-195/ALM-STALE-196): read-and-discarded presumes subsequent complete count reads. Successful-short/EOF reads can retain metadata or prior-count bytes in the reused local; that malformed-input effect is not metadata count semantics.**

**Confidence.** High (players/units/type8, and `+0x20` per [EXP-0037]) / **Unknown** (`+0x28`)

**Original status.** ● active (amended)

**Evidence.** [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### ALM-META-026

The type-0 `+0x20` low-bit bitmask is **read into a local and discarded** by the map loader (`local_60`, referenced nowhere else) — read-but-unused by `rom.exe`'s map parse. `+0x14/+0x18/+0x1c` and `+0x78/+0x7c` are read into the persistent map object (stored config scalars; `+0x78 ∈ {1,4,8,12,16}` — **1 on all 28 campaign (single-player) maps**, `{4,8,12,16}` on the 10 standalone MP maps, a player-slot/MP-capacity value). **`ALM-MODE-070` amends the corpus statement and identifies the consumer**: the `+0x70` scalar is the map object's `+0xd4` and the session's multiplayer flag is `(map+0xd4 > 1)` (`L02087`/`L02088`/`L02089`). All 28 `scenario.res` maps carry 1 on all three roots, but the converse fails — the RU root's loose `Horror.alm` carries 1 where the EN file of the same name carries 16, so the value does not identify a campaign map. The 28/10 split is the EN corpus; RU has 34 maps and a `1:29 4:1 8:2 16:2` histogram. **Amended by [EXP-0298]: the three stored scalars have destinations and no reader.** `+0x0c` → `M+0x1c`, `+0x10` → `M+0x20`, `+0x14` → `M+0x24` (`ALM-META-091`); nothing in the searched population loads any of them back off the map object, and the browser reader steps over payload `+0x08..+0x2f` entire (`ALM-META-092`); their corpus value sets are `ALM-CORP-093`

**Confidence.** Medium (bitmask discarded / scalars stored) / Unknown (precise meaning of the remaining scalars; `+0x70`'s consumer is `ALM-MODE-070`, High)

**Original status.** ● active (amended)

**Evidence.** [EXP-0018](../experiments/EXP-0018-alm-metadata/) — offsets −8 per [EXP-0030]: discarded bitmask `@+0x18`, stored `+0x0c/+0x10/+0x14` and `+0x70/+0x74`, [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-META-027

The type-0 float `+0x10` (`angle`, ≈±π/4) is read into the map object as a stored **light/sun angle** (corroborated by the Map-Editor Help "Show Time / Light" settings). `+0x04` (`selectorA`) is consumed as the per-record `perMapConst` enum (its bits reappear as adjacent-section framing, e.g. `type6.f1`). **Amended by [EXP-0030]: the angle is at payload `+0x08`, and `selectorA` is not a payload field at all** — it is the 5th word of every record header (which is why it "reappears": type6's apparent `f1` is type8's header `f32` when type8 is empty). **Amended by [EXP-0298]: the destination is `M+0x18`, and on the map object the value is never read** (`ALM-META-091`, `ALM-META-092`). The corpus reading is also sharper than "≈±π/4": all ten shipped values are whole degrees, `±45 ±44 ±18 36 28 25 22` (`ALM-CORP-093`). The angle's *other* destination, `P+0x20`, is overwritten before it is read (`TERR-LIGHT-149`, `TERR-LIGHT-151`) **Partially retracted by EXP-0346 (ALM-HEADER-098, ALM-WRITER-099): the record-header word has no demonstrated f32/enum interpretation, map-object numeric consumption or payload echo. Framing and unrelated metadata-angle clauses stand.**

**Confidence.** High (float read into map object) / Medium (light-angle naming)

**Original status.** ● partially retracted

**Evidence.** [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### ALM-META-028

The type-0 `+0xc0` region (**`+0xb8` per [EXP-0030]**) is a fixed **array of 7 × 64-byte text slots** (`[3×u32 prefix][char[52] text @slot+12]`, default `"<None>"`), read into the map object as the tail of the 512-byte description block. Standalone maps leave slots default/empty; campaign maps fill slots 4/6 with in-game trigger/quest strings (`"mission complit"`, `"Start1"`, `"1st Finger"`, …)

**Confidence.** High (slot layout + loaded — the region is inside case 0's read sequence, whose lengths sum to exactly 632, and the 7×64 B sub-structure is fixed by the `"<None>"` defaults landing at `slot+12` in every slot of every map) / Unknown (per-slot role) / **[EXP-0056]: neither consumer subdivides the region.** The loader reads payload `+0x78`..`+0x277` as **one 0x200 read** into `map+0xdc` — description and the seven slots together — and the `*.alm` browser's reader `R0479` reads the same 0x200 block, rewrites every `'\n'` to `'#'`, and assigns it as **one NUL-terminated CString**. The 7×64 sub-structure is not refuted (the `"<None>"` defaults still land where this row says) but it is a writer's model with no consumer behind it

**Original status.** ● active (amended)

**Evidence.** [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0056](../experiments/EXP-0056-alm-record-set/)

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-HDR-029

For sections other than type0, the 20-byte header words `f0/f1` are **not map fields**: `0` where the writer's source buffer was zeroed (type1/4/5), and adjacent record-framing / uninitialised residue elsewhere — decisively, **`type6.f1` = the per-map `selectorA` float bits** (Beast `0xBFF52D9D`, Forester `0xBFF62B6D`, Kids/Islands `0xBFC02B6D`). The loader's parse is driven by count+version+typeId+size, not by these words. **RESOLVED (and the premise retracted) by [EXP-0030]: there are no `f0/f1` words.** A record's apparent pair is the **last 8 bytes of the previous record's payload** — which is exactly why they are `0` after zero-tailed payloads (type0's empty text slots → type1; type3's sparse overlay → type5) and why `type6.f1` is the `selectorA` float (type8 precedes type6 and is empty, so its header's `f32` lands there). The last four "type2.f0/f1" words are valid tile words on 152/152

**Confidence.** High (the observation) / ✖ (the "header fields" reading)

**Original status.** ✖ retracted (superseded by ALM-FRAME-031)

**Evidence.** [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

**Amended.** The table ledger carried the status "✖ retracted (superseded by ALM-FRAME-031)".

### ALM-TRL-030

The 8-byte file trailer is **not read by the map parser** (`R0478` reads exactly `count` records and returns; no trailing read). Standalone maps: all-zero (10/10). Campaign maps: writer-set small ints (2nd u32 usually 1; 1st u32 ∈ {0,1,2,3}) — scenario-linkage tokens consumed outside the geometry/entity parse. **RETRACTED by [EXP-0030]: there is no trailer to read.** The parser reads `count` records and stops because the last record's payload already ends at EOF; those 8 bytes are the tail of the type-7 trigger payload

**Confidence.** ✖

**Original status.** ✖ retracted

**Evidence.** [EXP-0018](../experiments/EXP-0018-alm-metadata/), [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

### ALM-FRAME-031

**Corrected container framing (supersedes the EXP-0007 split).** A `.alm` is a **20-byte file header** `[char[4] "M7R␀"][u32 20][u32 dataSize][u32 recordCount][u32 formatVersion]` followed by `recordCount` records, each a **20-byte header `[u32 tag=7][u32 hdrLen=20][u32 payloadSize][u32 typeId][f32 perMapConst]`** + `payloadSize` bytes of **pure payload**. **No per-section `f0/f1`, no payload-head identity, no file trailer**; the file ends exactly at `20 + Σ(20+payloadSize)`. Proof: (a) **twelve shipped records have `payloadSize < 8`** (`type8 size=0` ×8 maps, `type4 size=0` on `scn:31`, `type9 size=4` ×3), so the old model's mandatory 8-byte payload-head identity cannot exist — a refutation from bytes alone, which also kills the read-overrun variant; (b) `R0465` peeks 8 B, seeks −8, then reads **`dword[cursor+4]`** bytes — with the cursor at file offset 0 on its first call, that reads **20** bytes as the file header, so `hdrLen` is a live self-describing field and `tag` is the header's first word; (c) `R0466`/`R0467` each `Read(dest, 0x14)` and the reader's `switch` dispatches on the record header's **dword 3** via a 10-entry jump table at `L02078` (under the old split that word is the constant `hdrLen=20` and the switch would be dead); (d) case 0's field reads sum to **exactly 632** = the type-0 `payloadSize` (the old split needs 640); (e) the second, independent parser `R0480` uses the same header helper and switches on the same word; (f) corpus 38/38 on every framing identity, 0 violations over 380 record headers. The only seeks on the path are that peek/rewind and the `default` case's skip — there is no per-section re-seek. Everything EXP-0007…0019 located keeps its **file** position; only labels move **−8**. Residual: `dataSize = 4·W·H+72` holds 38/38 but is no longer *explained* by the framing (see `ALM-HDR-001`) **Partially retracted by EXP-0346 (ALM-HEADER-098, ALM-WRITER-099): the record-header word has no demonstrated f32/enum interpretation, map-object numeric consumption or payload echo. Framing and unrelated metadata-angle clauses stand.**

**Confidence.** High (six independent proofs, listed (a)–(f) in the claim: one is a **refutation from bytes alone** — twelve records too small to hold the old model's mandatory payload head — and three are instruction-level. This is the row the rest of the ALM ledger's Highs are inherited from)

**Original status.** ● partially retracted

**Evidence.** [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim. Narrowed by `ALM-NEWMAP-237`: the header tag `7` is the installed value; the editor writes the map-carried word `E+8`, which is 5 in a map it creates.

### ALM-GRID-032

**Grid base and cell→index (the four-cell question, answered).** All three grid layers are addressed from the **record payload's first byte** with `index = row·W + col` — no offset, no padding, no overlay (`rom.exe` sim ingest `R0278` reads `tiles[row·W+col]`, `heights[row·W+col]`, `type3[row·W+col]`; the loader does `malloc(W·H·2)` + `Read(ptr, W·H·2)` with the destination un-adjusted, and `R0480` agrees). Because the record payload starts **8 bytes later** than `ALM-GRID-011` placed it, a decoder using the old base is **+4 cells off in X on type1 (u16)** and **+8 cells off on type2/type3 (u8)**. Corpus, from map bytes alone: placement anchors' water rate is minimised at **+4 cells** on type1 (2-D argmin exactly `(+4,0)` on 19/26 non-degenerate maps); their type-3 occupancy rate is minimised at **+8 bytes** on 29/35; the type1↔type2 mutual alignment (η² of height by strip group, **no placements involved**) peaks at **+4 bytes** on 32/38; and the derived-impassability rate over 11 235 placements is **4.86 %** at `(4,8)` versus 14.00 % at `(4,4)` and 20.69 % at `(0,0)` — the u16-vs-u8 split that only a byte-level framing shift can produce. The four "missing" cells are inside the payload: the words EXP-0007 read as the next section's `f0/f1` are valid tile words on **152/152**

**Confidence.** High (the ingest's own indexing read instruction-level, and — unusually — a corpus that *does* discriminate: four independent statistics with argmins at the same shift, including one (η²) computed without any placement data, so it cannot be an authoring artefact. Every figure carries its denominator)

**Original status.** ● active

**Evidence.** [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

### ALM-PLACE-033

**Placement anchor → terrain cell is a bare `>>8`.** No origin is subtracted, no border inset added, no rounding applied. Type-4 objects: `rom.exe` `R0478` case 4 shifts the record's first two dwords right by 8 bits (`L02090`/`L02091`, and `L02092`/`L02093` in the extension arm) and passes them to `R0481`/`R0482`, which store them as **u16 cell coordinates**. Type-6 units: the loader stores the raw `/256` anchor in the unit struct's `+0x00/+0x04` and the routine walking the loaded unit array (`mapObj+0x318`, `R0151`) converts with the same bare shift by 8 (`L02094`/`L02095`, `L02096`/`L02097`). A whole-binary scan of **all 82** shift-by-8 sites finds no `.alm` coordinate biased before or after the shift, and the only border in the ingest is symmetric 8 cells on all four sides (`sim+0x58ee0..3 = 8, 8, W-9, H-9`), which cannot yield an X-only 4

**Confidence.** High (objects + the shift) / Medium (that `R0151` is *the* unit spawn walker, vs one of several consumers)

**Original status.** ● active

**Evidence.** [EXP-0030](../experiments/EXP-0030-alm-grid-origin/)

### ALM-OBJ-034

**type-4 extension discriminator = `kind == 0x21`.** The loader reads the 20-byte object record as `4,4,4,2,4,2` and then tests the third dword (`rec+0x08`) against `0x21` (the comparison at `L02098`); on equality it reads **8 more bytes** and constructs the object through `R0482` (two extra u16 fields) instead of `R0481`. With this rule the type-4 walk consumes every payload **exactly 38/38** — resolving the `ALM-OBJ-019` open item and dissolving its Cross/Horror/scn:81 "anomalies" **Amended by [EXP-0037]: this is the *extension* discriminator only.** `kind == 0x21` records take the **object** arm of the spawn walker; the object/structure discriminator is a different pair of immediates, `0x22`/`0x23` (`ALM-CLS-037`). What the 8 extra bytes are is now read too: their low bytes override the object's footprint `w×h` in tiles (`ALM-CLS-036`)

**Confidence.** High (one named `CMP` immediate, plus the walk it turns from 34/37 into 38/38 — a rule read from code that then explains anomalies the corpus alone had only been able to carve out)

**Original status.** ● active (amended)

**Evidence.** [EXP-0030](../experiments/EXP-0030-alm-grid-origin/), [EXP-0037](../experiments/EXP-0037-alm-class-binding/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-CLS-035

**type3 = the static-object placement layer, and its code names a class.** A nonzero cell code `c` resolves to `objects.reg` **section index `c − 1`** — the `%d` of `Object%d`, not the `ID` key and not `c` itself. `rom.exe` does the arithmetic at three independent sites: the map renderer `R0379` (`cVar4 = (char)local_1c0 + -1` then `[L02099][local_1c0 & 0xff]`, reading `File @+0x0c`, `Index @+0x10`, `DeadObject @+0x44`), `R0483 @L02100` (`[L02099] + -4 + bVar2*4`), and `R0484 @L02101` (`bVar2 = cVar1 - 1` then `FireObject @+0x40`). This resolves the `ALM-GRID-014` open item ("exact code→object identity"); occupancy is *derived* from the layer (`ALM-TERR-016`), it is not what the layer is. Census re-run on the corrected base, 38 maps: **86.19–97.94 % zero, 71 099 nonzero cells, 98 distinct codes 1..246** — i.e. the pre-EXP-0030 figures reproduce **exactly**, so `ALM-GRID-014`'s stale-base flag is cleared rather than corrected, and the suspicion that `246` was an injected header byte is **wrong**

**Confidence.** High (the `− 1` and the array: three consuming sites read at instruction level, plus the render, which draws every code's `DescText` — under the rival `c` reading every tree becomes its own *burning* variant, because `Object%d` interleaves live/fire/dead. **Corpus alone would not have carried it**: `c − 1` scores 99.9100 % in-range against the rival's 99.8256 %, a 60-cell gap) / Medium (the census figures — corpus 38/38) / **Unknown** (64 of 71 099 cells, in `scn:131`/`scn:150` rows 0..2 only, index past the 82-entry array under `c − 1`; whether `rom.exe` reads them at runtime is not established. A consumer must bounds-check)  / **the 64-cell Unknown is answered by [EXP-0045] (`ALM-CLS-051`)**: they are head-of-buffer residue, not placements, and the contaminated span is wider than the 64 cells that happen to fall out of range

**Original status.** ● active (amended)

**Evidence.** [EXP-0037](../experiments/EXP-0037-alm-class-binding/), **[EXP-0045](../experiments/EXP-0045-alm-triggers/)**

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-CLS-036

**type4 `kind` (file `+0x08`) is the class-selecting field, and its domain is `structures.reg`'s `ID`.** The loader keeps its low 16 bits at the in-memory record's `+0x04`; the `CObject` ctor `R0485` stores the **low byte** at `obj+0x40`; `R0486` then uses that byte as a **1-based** subscript into the placeable-definition table `L02102` (`ECX=L02102` at `L02103`/`L02104`, guard `kind != 0 && kind <= R0487() = count−1`), parks the definition at `obj+0x3c`, and copies its parameters: params 0/1 → `obj+0x60`/`+0x61` = the footprint **in tiles** (`R0488` walks exactly that rectangle against the `obj+0x68` mask), param 2 → `obj+0x48`, param 3 → `obj+0x44`/`+0x42`. **When the 8-byte extension is present (`kind == 0x21`, `ALM-OBJ-034`) its two low bytes override the footprint** — **named by [EXP-0091]**: file `+0x14` → `obj+0x60`, file `+0x18` → `obj+0x61` (`ALM-OBJ-062`), and the arm is chosen by the **sum** of those two bytes, not by the kind (`TERR-STRUCT-090`) and `obj+0x68` becomes `-1` (full rectangle) — which is what `structures.reg`'s `VariableSize` key is for. Corpus, 3141 records / 38 maps: `kind ∈ 1..66`, **65 distinct, never 0** — exactly `structures.reg`'s `ID` domain (`1..66`, and `ID == section index + 1` on 66/66). **Amended by [EXP-0049]: the populator is found and the parameters are named** — the table is the Buildings collection of `world.res:data/data.bin` (`DAT-LOC-001`), params 0/1 are the columns `sizeX`/`sizeY`, param 2 `scanRange`, param 3 `healthMax`; the 1-based law is the file's own entry-0 skip (`DAT-GRAM-003`), and the `ID` ordering is `DAT-BLD-005`, corroborated name-for-name against `structures.reg`

**Confidence.** High (that `kind` is the class-selecting field, that it is 1-based and bounded by another table's count, and the extension's footprint override — all read from `R0486`'s own instructions; and that the domain is `1..66`/65-distinct/never-0 over 3141 records, which excludes the `objects.reg` reading, whose `ID`s run `0..81`) / **High as amended by [EXP-0049]** (the `structures.reg`-`ID` ordering — was Medium "populator not found": the populating file is read, its index law is the skip-0 grammar, and 66 ordered names agree with `structures.reg`'s 44/66 verbatim, `DAT-BLD-005`)

**Original status.** ● active (amended)

**Evidence.** [EXP-0037](../experiments/EXP-0037-alm-class-binding/), **[EXP-0049](../experiments/EXP-0049-placeable-db/)**

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-CLS-037

**The type-4 object/structure discriminator is `kind ∈ {0x22, 0x23}`, read from the walker.** `R0489` walks `mapObj+0x304` and tests the in-memory record's `u16` at `+0x04`: the sign-extended 16-bit value is read at `L02105` and compared with `0x22` at `L02106`, taking the structure arm on equality, and otherwise compared with `0x23`, taking the object arm on inequality. On a match it builds a **`0x74`-byte `CStructure`** (`R0490`, vtable `PTR_R0491`, plus a `0x98`-byte sub-object at `+0x6c` and `obj+0x70 = (i16)rec+0x06 × 1000`); otherwise a **`0x6c`-byte `CObject`** (`R0485`, `PTR_R0492`). **`kind == 0x21` — `ALM-OBJ-034`'s *extension* discriminator — takes the OBJECT arm**, so the two discriminators are different questions with adjacent immediates. `0x22`/`0x23` are `structures.reg` `ID` 34/35 = `Shop 1`/`Shop 2`, the two `Usable` shops. Corpus: 50 of 3141 records (26 + 24) take the structure arm. **Amended by [EXP-0057] on two points.** (a) The two classes are named **`Shop`** (`0x74`) and **`Building`** (`0x6c`) by the MFC runtime-class table `Shop::GetRuntimeClass` returns (`R0493` returns the class-table address `L02107`); "`CStructure`"/"`CObject`" were this experiment's own labels (`SHOP-CLS-001`). (b) **`+0x0c` is not a durability and is not the constant `10000`** — this row's "they carry `+0x0c = 10000`, i.e. `obj+0x70` = the ctor default `0x989680`" is withdrawn: `L02108` sign-extends the 16-bit value at record `+0x06`, multiplies it by `0x3e8` and passes it to `R0494`, which forwards it to the stock generator's **value cap** `template+0x90`, and over those same 50 EN records it takes 23 distinct values (caps 4 000 … 10 000 000) with only 12 at `10000` (`SHOP-CAP-004`, retracted.md). Consequence for `ALM-OBJ-019`'s label: **type-4 is not "objects *or* structures resolved against two registries"** — every record is a `structures.reg` class, and the split is between two C++ classes

**Confidence.** High (two `CMP` immediates and the two constructors they select, named to the instruction; the rival — `0x21` as the object/structure discriminator — is refuted by those same instructions, which would otherwise be dead) / Medium (that `+0x0c × 1000` is a durability/HP quantity rather than some other scale — inferred from the ctor default it matches, no instruction names it)

**Original status.** ● active (amended)

**Evidence.** [EXP-0037](../experiments/EXP-0037-alm-class-binding/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### ALM-CLS-038

**type6 names its unit class by `units.reg`'s `ID`, and resolves it through a definition collection, not a registry subscript.** File `+0x08` (`MOVSX`-ed i16 → record `+0x08`) is the primary key and file `+0x0a` (u16 → record `+0x0c`) the secondary one: `R0151 @L02109` calls `R0495(L02110, (byte)rec+0x08, (byte)rec+0x0c)`, which searches a node collection matching **parameter `0x10` == `rec+0x08`** when `rec+0x08 < 0x40 && ∉ {0x1a, 0x1b}`, and otherwise matches **parameters `0x1d`/`0x1e`** against both words; the index it returns is used as `R0496(L02111, idx) + 4` → the live `CUnit` ctor `R0497(new(0x1e8), …)`. Two overriding paths: file `+0x10`, when nonzero and `!= 0xcdcdcdcd`, is a definition **id** found by `R0498` on parameter `0x18`; and `rec+0x48 & 1` diverts to the `scenario/npc.reg` table `L02112` indexed by `rec+0x0c` (that table is filled by `R0499` from `npc_%d` sections). Which registry key `+0x08` is, is settled by the corpus: 8094 records / 38 maps, domain `1..80`, **inside `units.reg`'s 34-element `ID` set on 8094/8094**, inside its 34 section indices on only **1447/8094 (17.88 %)**. The live `CUnit` holds a units.reg **`ID`** at `+0x20` (`[L02113][unit+0x20]` in `R0379`) — **corrected 2026-07-27, this row previously said "section index" and "a translation happens between the two"**: that array is keyed by `ID` (`REG-KEY-044`), so the map's key and the array's key are the same and **nothing is translated**. Medium, because the instruction that *writes* `+0x20` is still unlocated **Amended by [EXP-0049] on three points.** (a) `L02111` is the **Humans** table of `world.res:data/data.bin` and `L02114` (`+0x8c`) the **Units** table (`Humans.csv`/`Units.csv` — `DAT-OBJ-002`); parameter `0x10` is the Humans column **`typeID`**, `0x18` is **`serverID`**, `0x1d`/`0x1e` are the Units columns **`typeID`**/**`face`**. (b) This row's `R0496(L02111, idx) + 4 → R0497(new(0x1e8))` describes only the **humans** arm; the units arm is in the same function — `L02115 new(0x198)`, `L02116/L02117` `R0500(L02114, idx)+4` → ctor `R0501` (`DAT-ACT-006`). (c) The resolution closes over the shipped file: `ALM-CLS-052`

**Confidence.** High (that `+0x08`/`+0x0a` are the class-selecting fields and `+0x10` the overriding definition id — read from `R0151`/`R0495`/`R0498`; and that the key is the `ID` and **not** the section index, which the corpus excludes outright, failing on 82 % of records) / **High as amended by [EXP-0049]** (the `0x10`-is-the-class-key clause — was Medium "populator not found": the populating file, its `typeID` column and the streamer that copies it into the drawable-class subscript `actor+0xe` are read, `DAT-SCHEMA-004`) / ~~Unknown (what fills `L02110`)~~ → **answered: `DAT-LOC-001`**

**Original status.** ● active (amended)

**Evidence.** [EXP-0037](../experiments/EXP-0037-alm-class-binding/), **[EXP-0049](../experiments/EXP-0049-placeable-db/)**

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-OWN-039

**Owner, both record types: type-4 file `+0x0e` and type-6 file `+0x14` are the same field — a 1-based physical slot in the type-5 array.** Both are handed to `R0425(this=[L00380], id)`, which walks the group list comparing `node+0x08` and returns the group (`L02118`/`L02119` for type-4 → `obj+0x14`, failure printing `"Warning - Building without owner"` at `L02120`; the type-6 site at `L02121`/`L01836`, failure formatting `L02122`). The corpus discriminates between the two live readings of the value: **1-based slot** holds on **3141/3141** type-4 and **8094/8094** type-6 records, while the type-5 record's own id word (file `+0x00`, which the loader parks at `player+0x08`) fails on **22.70 %** / **27.65 %** — e.g. `Forester.alm` uses owner `3` and its roster ids are `7 2 5 4`

**Confidence.** High (the field's identity and its consumer — one lookup, two call sites, both named; and the slot-vs-id question is decided by the bytes: the id reading is refuted on 713 + 2238 records) / **Medium** (that the runtime group object's `+0x08` is exactly the loader's `player+0x04 = slot+1` — `R0065`, which builds that list, was not read)

**Original status.** ● active

**Evidence.** [EXP-0037](../experiments/EXP-0037-alm-class-binding/)

### ALM-UNIT-040

**The type-6 70-byte record, field by field, from the loader's own read sequence** (`R0478` case 6: `new(0x50)` then 27 sequential reads summing to exactly `0x46` at `formatVersion == 990`; the three version gates are `CMP …,0x3db`/`0x3b7`/`0x3d9` at `L02123`/`L02124`/`L02125`, so at other versions the record is a different length). `+0x00` X, `+0x04` Y (`/256`); `+0x08` i16 class key, `+0x0a` u16 secondary key (`ALM-CLS-038`); `+0x0c` flag word → `rec+0x48` (bit 0 = NPC path, bit 2 → `unit+0x4b` high bit, bit 7 passed on); `+0x10` definition id (present only when version > 0x3da); `+0x14` owner (`ALM-OWN-039`); **`+0x18` a 1-based index bounded by `meta+0x2c`** — the loader stores it minus one, sets it to `-1` when `0` or `≥ meta+0x2c`, and otherwise writes the record's `+0x40` word into entry `[value−1]` of `mapObj+0x2dc` (`L02126`…`L02127`); `+0x1c` u32; `+0x20/+0x22/+0x24/+0x26` u16; `+0x28/+0x2a` u16 (version > 0x3b6); `+0x2c..+0x34` nine bytes (`+0x32` version > 0x3d8); `+0x35` six bytes and `+0x3b` five bytes — `ALM-UNIT-018`'s `0xFF`-sentinel region; `+0x40` u16; **`+0x42` u32, a unique id whose running maximum the map keeps at `mapObj+0x90`** (`L00299`). Corpus, 8094 records / 38 maps: the `+0x18` guard is satisfied **8094/8094** (43 records use a nonzero value). **Amended by [EXP-0045]: `+0x40` and `+0x42` are labelled the wrong way round here.** `+0x40` is the **unit id** and `+0x42` the **group id**: on `scn:110` `+0x40` takes 39 distinct values over 39 records and `+0x42` takes 19, and the type-7 trigger parameters resolve `Target_Unit` against `+0x40` and `Target_Group` against `+0x42` (169/169, `ALM-TRIG-046`). The instruction fact stands — `L00299` does keep a running maximum of `+0x42` at `mapObj+0x90` — but what it maximises is the **group** id, which is what an engine that allocates groups at runtime needs. `ALM-UNIT-048` **Amended by [EXP-0296]: the read count is 28, not 27.** The widths still sum to exactly 70 and every field label here stands, but 18 of the 28 mappings are not the identity — 17 land in the record at a different offset and one passes through a stack local — so file offset and runtime offset must not be used interchangeably. From file `+0x2c` to `+0x40` the shift is `−4`, which puts the six-byte run listed above at `rec+0x31`…`+0x36` and the five-byte run at `rec+0x37`…`+0x3b` (`ALM-TAILMAP-079`). The `0xFF`-sentinel label those two runs carry above is refuted, not merely inherited: they are the skill and elemental-resistance override vectors, and no byte of file `+0x2c`…`+0x3f` is `0xFF` on any of the 12 085 records of the two preserved roots (`ALM-TAILRUN-083`, `UNIT-PLACESKILL-086`, `UNIT-PLACERESIST-087`). The Medium clause below — the roles of `+0x1c` and `+0x20`…`+0x2b` — is answered by `ALM-TAILDIR-081`, `ALM-TAILU16-082`, `ALM-TAILVER-084` and `UNIT-PLACE-034`

**Confidence.** High (the read map, the `+0x18` guard and the `+0x42` max — each tied to a named instruction, and the reads sum to exactly 70; and the `+0x18` bound is `meta+0x2c` because that is the local case 0's twelfth read filled, `[EBP-0x18]`) / Medium (the *roles* of `+0x1c`, `+0x20..+0x2b`, `+0x40` — located as read destinations only)

**Original status.** ● active (amended)

**Evidence.** [EXP-0037](../experiments/EXP-0037-alm-class-binding/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record. `ALM-UNITREAD-227` reads the stream read slot behind the "read destinations only" caveat of the Medium clause: no holder dereferences `rec+0x1c`, `+0x22`, `+0x26` or `+0x4c`; the caveat that remains is array append `L02248` and the accessors.

### ALM-GRP-041

**The type-5 76-byte record, field by field** (`R0478` case 5: `new(0x48)` + ctor `R0502`, then four reads and a 16-iteration loop, `4+4+4+32+16×2 = 76`). `+0x00` → `player+0x08`: ~~**the word the owner lookup `R0425` compares against**~~ — *[EXP-0180]: the field mapping is exact and the appositive names the wrong field. `R0425` compares the session `Player`'s `+0x08`, which `R0423` fills at `L02128` from this record's `+0x04` of the loaded `CPlayer`, i.e. the record's own 1-based ordinal. This word is the faction's **colour slot**: `L02129` carries it to `Player+0x44` as `+1`, the join sender puts `+0x44 - 1` on session message `0x96`, and the unit body draw indexes the 17 shade objects with it (`ALM-PLAYER-069`, `PAL-SHADE-012`, `PAL-SHADE-013`)*; `+0x04` → `player+0x30`; `+0x08` → `player+0x0c`, the `{0,5000}` scalar `ALM-GRP-020` noted; `+0x0c` a **32-byte** NUL-terminated ASCII group name → `player+0x10`; `+0x2c` **sixteen u16** read through `R0130(player+0x34, i)` — one per editor player slot (`ALM-META-025` records the editor's 16-player cap), i.e. the diplomacy row — **confirmed by EXP-0074, which read the consumer this row said it did not have**: the map loader `R0128` fetches element `k` through this very `R0130(player+0x34, k)` and stores its low byte at `session+0xa9c4 + 50·(player+0x04) + k + 1` (`L02130`, `L00254`), then forces the diagonal to 2 (`L00255`); that byte is what the acquisition filter `R0021` tests bit 0 of (`AI-DIPLO-005`, `AI-FILTER-001`). Over the 38 shipped maps the sixteen words take **only the values 0, 1 and 2** (2362/471/223 of 3056), so the u16 never exceeds the byte the engine keeps, and no map ships more than 16 players. The loader additionally stores `slot+1` at `player+0x04`. Corpus: names and ids in `evidence/alm-type5-roster.csv`; the id word is sparse and not ascending (e.g. `Cross.ALM` `8 2 3 5 1 6 4 13`), so it is an id and not a slot number **Partially retracted by EXP-0351 (ALM-139, SESS-072): only the inherited universal human-participant interpretation of wire+04 is withdrawn. The field map, name width, ordinal/colour distinction, diplomacy row, corpus observations and separately scoped scalar clauses stand.**

**Confidence.** High (every field is a read boundary in a sequence summing to exactly 76, and the name field's length is fixed at `0x20` by the read itself — `ALM-GRP-020` guessed at it from content) / **High** (that the 16 u16 are the diplomacy row — EXP-0074 read the consumer, and the store's index arithmetic is the same base and stride the filter reads) / ~~Unknown (`+0x04`, and the `{0,5000}` scalar's role)~~ — *[EXP-0180]: `+0x04` has the confirmed destination `Player+0x28`; the universal human-participant interpretation inherited from `UNIT-OWNER-009` is partially retracted, while the destination and following corpus observations stand — `R0423` stores it twice, `L02131` and `L02132`. Over the corpus it takes only 0 and 1, and the split is exact: every one of the 28 embedded campaign maps authors 0 on record 0 and every loose map authors 1, with every record above ordinal 0 authoring 1 on every map (`ALM-PLAYER-069`). ~~The `{0,5000}` scalar at `+0x08` stays Unknown~~ -- *[EXP-0297 correction]: the destination is `CPlayer+0x0c`; the seven-function family has one direct read in copy constructor `R0503`, which has two direct call sites inside `R0504`. The bounded population is 17 map-side plus 7 family routines, not 24 map-side routines; no map-side direct-displacement read was attributed, but computed pointers and the whole-image mover/serializer/`memcpy`-shaped surface remain unsearched (`ALM-SCALAR-087`, `ALM-SCALAR-088`). The 357-record corpus law is `ALM-SCALAR-089`; consumer effect stays **Unknown**. `R0423` itself does not read the field*

**Original status.** ● partially retracted

**Evidence.** [EXP-0037](../experiments/EXP-0037-alm-class-binding/), [EXP-0074](../experiments/EXP-0074-target-acquisition/), [EXP-0180](../experiments/EXP-0180-shade-select/), [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/), [EXP-0351](../experiments/EXP-0351-authored-player-control/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-CLS-042

**Shipped maps place `objects.reg` *fire* variants, and those classes' art does not ship.** Resolving all 71 099 nonzero type-3 cells over the 38 maps through the established key (`code − 1`, `ALM-CLS-035`), **33 cells** — codes 5, 8, 14, 20, at 4 / 2 / 5 / 22 cells respectively (`evidence/alm-type3-code-hist.csv`) — land on `objects.reg` section indices 4, 7, 13, 19, four fire-variant classes with **empty `DescText`** whose `Files[class.File]` paths are `vallen1\fire\fvallen`, `bush1\fire\fbush1`, `iva1\fire\fiva` and `oak1\fire\foak` (`evidence/reg-objects-rows.csv`). Those four sprite files are reported present in **none** of the install's eight archives and nowhere on disk. **AMENDED by [EXP-0038], which re-ran that absence rather than inheriting it and widened it from the placements to the registry: there are SEVEN such classes, not four** — sections 1, 4, 7, 10, 13, 16, 19, i.e. every `<species>\fire\f<species>` path (`pine1`, `vallen1`, `bush1`, `elka1`, `iva1`, `maple1`, `oak1`), each the middle member of a live/fire/dead triple, each carrying the 105-entry animation timeline and its live sibling's `DeadObject`. All seven are absent both at the node the loader would open (`graphics.res!objects/<File>.256`) and from **every** archive node in the install — re-measured, `evidence/run-reg.txt`. **Four** is the number that are *placed*; that figure and the 33 cells stand. So the cells are not corrupt and not out of range: the placement is valid, the class exists in the registry, and only its art is absent. **Consumer consequence: a type-3 decoder must tolerate a resolved class with no loadable sprite** rather than treat it as bad data — distinct from `ALM-CLS-035`'s 64 genuinely out-of-range cells. What the engine draws at such a cell is **not established**; the codes were found while generating EXP-0037's review page, which also showed that the same "nameless and spriteless" signature covers **30 344** cells (42.7 %) under the rival `code` key, so it is the ~920× ratio and not the signature's presence that discriminates the two keys

**Confidence.** Medium (**that the 33 cells resolve to those four named classes is re-checked here against this experiment's own committed evidence** — the four code counts sum to exactly 33 and the four `FilePath`s reproduce — a whole-corpus measurement, capped at Medium because it sits on top of `ALM-CLS-035`'s key rather than on an instruction) / **High** (that the art is absent — the cheap restoration this row asked for was done by [EXP-0038]: every archive in the install re-walked for all seven fire paths, 0 hits, run recorded in `EXP-0038/evidence/run-reg.txt`. Still an absence in our own file index rather than in `rom.exe`'s loader, but the index is now exhaustive over the install rather than a by-product of a review page) / **Unknown** (what the engine renders there, and whether the fire variants are reachable at runtime at all — [EXP-0038] narrows this: their frames would be selected by `TERR-SPR-042`'s animated arm, whose tile-word gate no shipped cell opens, `TERR-TILE-044`)

**Original status.** ● active (amended)

**Evidence.** [EXP-0037](../experiments/EXP-0037-alm-class-binding/) (added 2026-07-27)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### ALM-TERR-043

**`world.res:data/map.reg` in full — and the half of it the engine never reads.** Walked on the corrected EXP-0032 framing (`0x18` header, 32-B records at `0x18+32i`, pool after the record array): **31 records, 3 sections**. `Terrain` = 20 records **interleaved** `Cost`,`Pass` per class in the `ALM-TERR-015` order, `Path Finding` = 7, `Scanning` = 1. Values: **`Cost = [8,8,8,14,6,12,8,16,8,6]`**, `Pass = [0,0,0,0,0,0,0,1,1,1]` (Land…Road); `SpeedMultiplier 8`, `StaticScanAhead 5`, `DynamicScanAhead 3`, `StaticRefreshRate 16`, `DynamicRefreshRate 32`, `DynamicByStaticLookup 3`, `StaticIsntNeeded 5`, `ScanShift 7`. Three facts a key list alone would miss. (1) **The ten `Pass*` scalars are read by nothing we ship.** A whole-image ASCII scan finds **10/10 `Cost*` key strings and 0/10 `Pass*` in `rom.exe`**, and 0/10 of either in `Map Editor.exe`; the only `Cost`/`Pass`/`Terrain` strings in `rom.exe` are the ten `Cost*` names and `Terrain`, so there is no `Pass%s` built at runtime either. Impassability is **hardcoded** — terrain class 8 and the raw water bit test (`TERR-PASS-049`) — and reading it out of `Pass` costs 4.6 points on the shipped-placement check. (2) **The key names in the file are truncated to 15 characters and `rom.exe` asks for longer ones, yet the lookups still hit**: `R0505` truncates the *request* to 15 on both its paths (`strncmp(...,0xf)` on the linear one; a copy into a 16-byte buffer with `buf[15]=0` before the `bsearch` on the other), so `DynamicByStaticLookup` (21 chars) matches the stored `DynamicByStatic`. All seven `Path Finding` keys are found, not defaulted. (3) **`rom.exe` carries its own defaults for all ten `Cost*` keys** — `[8,8,9,14,6,12,11,16,8,6]` — which differ from the file at Flowers (9 vs 8) and Savanna (11 vs 8); the difference reaches 128 694 of 880 704 cells, so the distinction is load-bearing. Destinations: `Cost` → `world+0x54176+class` with `[0x54176] = 0xff` for class 0; `Path Finding` → `world+0x585b4..0x585c8` and `SpeedMultiplier` → `world+0x58db4`; `ScanShift` is read only by `R0274` and its use is not traced

**Confidence.** High (the record walk is the EXP-0032 framing applied to 31 records with their raw bytes recorded; the ten `Cost` reads, their order, their defaults and their destination bytes are the `PUSH` operands and `MOV`s of `R0116` @`L02133…L02134`; the 15-character truncation is read in `R0505`; and the `Pass` absence rests on a **whole-image byte scan**, which — unlike a reader enumeration — no code path, orphan or undisassembled, can evade) / Medium (that `Pass` is therefore *unused by the game* rather than merely unread by these two binaries: only `rom.exe` and `Map Editor.exe` were scanned)

**Original status.** ● active

**Evidence.** [EXP-0041](../experiments/EXP-0041-terrain-passability/)

### ALM-TRIG-044

**The type-7 payload is THREE counted arrays, not one.** `[u32 nAct]` + `nAct` action records of **796 B**, then `[u32 nCond]` + `nCond` condition records of **796 B**, then `[u32 nTrg]` + `nTrg` trigger records of **184 B**. The walk consumes every payload exactly, **38/38**, 0 residue. The two node lists hold the "THEN" vocabulary (`Description Instants.ini`, 44 opcodes) and the "IF" vocabulary (`Description Checks.ini`, 22 opcodes) respectively; the third binds them. Corpus totals: 797 actions, 680 conditions, 421 triggers. This is what `ALM-TRIG-021` saw the front of — its `entryCount` is `nAct`. It also **names EXP-0007's "8-byte file trailer" byte for byte**, and reproduces the distribution `ALM-TRL-005` published before either the framing or the grammar was known: on a map with no triggers the last two words are `[nCond][nTrg]`, both 0 (all 10 standalone maps, 6 of which also have nTrg = 0 and 4 of which end in a trigger whose two words are 0); on a campaign map they are the last trigger's `+0xb0` **comparison code** and `+0xb4` word — which is exactly `ALM-TRL-005`'s "2nd u32 usually 1, 1st ∈ {0,1,2,3}" (26 of 28 campaign maps end in `1`; the 1st word is 0/1/2/3 on 27 of 28, and on `scn:111` it is `0x02020101` in a trigger whose third pair is unused)

**Confidence.** **Medium** (corpus, and the corpus **does** discriminate inside the model family: every `(A,B)` with `A ≤ 2000`, `B ≤ 1200` — 150 000 models — was walked over all 38 maps and **exactly one**, `A=796 B=184`, consumes every payload exactly. Capped at Medium because no `rom.exe` instruction that reads these arrays was examined this round; `R0477` is what would settle the iteration bound)

**Original status.** ● active

**Evidence.** [EXP-0045](../experiments/EXP-0045-alm-triggers/)

### ALM-TRIG-045

**The 796-byte node record, field by field**, identical for an action and a condition: `+0x000 char[64]` the map author's own label (NUL-terminated inside a fixed buffer whose tail is uninitialised editor memory) · `+0x040 u32` **opcode** = the `.ini`'s `ID` · `+0x044 u32` **id**, unique within its list on 38/38 lists of each kind · `+0x048 u32` **0 on all 1477 nodes** · `+0x04c u32 value[10]` · `+0x074 u32 type[10]` · `+0x09c char[64] pname[10]`. **Parameters are stored by `Par<i>` SLOT, not packed**: 125 nodes carry an empty slot ahead of a used one (804 empty slots), and all 125 are the `Group Command` opcode whose `Par9` is the target group while `Par3..Par8` are unset — which a packed list cannot produce. Cross-check against the shipped catalogs, which this repo did not author: **1476 of 1477** nodes match a declared group on all ten `(name, type)` slots, 0 nodes carry an undeclared opcode, and all **155** `Const` literals equal the value the `.ini` declares — the literal that also selects among the eleven signatures sharing opcode 6 **Confirmed and completed by [EXP-0080] from the loader itself.** `R0478` case 7 reads `0x40, 4, 4, 4, 0x28, 0x28, 0x280` into `+0x000, +0x040, +0x044, +0x318, +0x048, +0x070, +0x098` — so this file framing is right, the record is 796 either way, and the **runtime** struct is a *permutation* of it in which the file's `+0x48` word lands at `+0x318` and everything after it moves down four bytes. The Unknown is answered: the **opcode word wins** — the runtime dispatches on `+0x40` and reads parameters by their type code, so a node whose block disagrees with its opcode is built (if its references resolve) and silently measures the wrong thing, its unused plain slots reading 0 because `R0506` zero-initialises the record.

**Confidence.** **Medium** (a 2880-slot agreement with a prediction this repo did not write, plus the slot-gap refutation of the packed rival; still corpus only — `R0476` (get-param) is the read that would pin the three arrays and their indexing) / **Unknown** (which of two disagreeing words wins where they conflict: `scn:120`'s condition `"PalNearChief"` stores opcode **7** with opcode **6**'s parameter block, and under opcode 7 its second value would be an X of 10005 on an 80×80 map)

**Original status.** ● active

**Evidence.** [EXP-0045](../experiments/EXP-0045-alm-triggers/)

### ALM-TRIG-046

**The nine parameter type codes, and the identifier domain each names.** `1` = `int` (and `Enum`, which serialises as an int) · `2` = `Target_Group` · `3` = `Target_Player` · `4` = `Target_Unit` · `5` = `X` · `6` = `Y` · `7` = `Const` · `8` = `Target_Item` · `9` = `Target_Structure`; `Target_Building` is declared by the editor but placed by no shipped map. Derived by pairing each declared `Par<i>_Value` with the code the bytes carry in the same slot — **0 conflicts** over 2880 slots. What the values name, from the map's own records: **X/Y are plain tile coordinates**, inside `[0,W)×[0,H)` on **505/505**; `Target_Group` is the type-6 record's **`+0x42`** word (169/169 distinct references resolve); `Target_Unit` is the type-6 **`+0x40`** word; `Target_Structure` is the type-4 **`+0x12`** word (16/16), which answers that field's "role Unknown" in `ALM-OBJ-019`; `Target_Player` takes 1..8 **[EXP-0080] adds the band structure `Target_Unit` actually has**: `R0507` sends a value `< 10001` to the map's own unit table, `10001..11000` to a hero ordinal resolved against the live player list (`R0508`, which returns 0 outright in multiplayer), and `> 11000` to a static name table at `L02135`; the failure strings are `"Can't resolve unit %d."` and `"Can't resolve hero %d."`. With the bands honoured **every** `Target_Unit`, `Target_Group`, `Target_Structure` and `Target_Player` reference in all 38 maps resolves — 0 of 1304 buildable nodes would be rejected by the builder.

**Confidence.** **Medium** (X/Y and `Target_Group` are whole-corpus exact against a domain another record type publishes — a coordinate out of grid or an id absent from the unit records would show, and none does; the `Target_Structure` identification is exact but rests on 16 references) / **Unknown** (`Target_Item`: values 2..36, only 62.5 % inside the type-4 `+0x12` set, so its domain is not called here)

**Original status.** ● active

**Evidence.** [EXP-0045](../experiments/EXP-0045-alm-triggers/)

**Amended.** EXP-0565 settles the `Target_Item` domain at instruction level: the binder stores `(V + 0xe18) mod 0x10000` as an item code (`TRIG-ITEMBIND-121`); for V 0..231 that code names MagicItems row 24 + V, with completed creation/binding and pack traversal, first-match order and whole-pack absence graded Medium (`TRIG-ITEMKEY-124`); the installed values 2..36 are rows 26..60 (`TRIG-ITEMCORPUS-125`). The type-4 `+0x12` comparison above does not bear on that domain.

### ALM-TRIG-047

**The 184-byte trigger record, field by field.** `+0x00 char[64]` name · `+0x40` 64 opaque bytes; the historical corpus found no printable strings in 421 records, but the heap-address and universal non-text interpretation is **partially retracted** (ALM-TRIGSTORAGE-163, ALM-TRIGEDITOR-164, ALM-TRIGZERO-165) · `+0x80` three `(u32 left, u32 right)` **condition-id pairs** · `+0x98` four `u32` **action ids** · `+0xa8` three `u32` **comparison codes**, one per pair, alphabet `{0,1,2,3,4,5}` · `+0xb4 u32` ∈ {0,1}. **The slots hold a node's `+0x044` id, not its array index**: 735/735 action and 1092/1092 condition references resolve by id, against 85.03 % / 92.03 % by 1-based index — decided by the 16 maps whose action ids are not `1..nAct` dense. The three condition slots are **pairs**, not a flat list: their nonzero counts are 398/398, 116/116, 36/36 and **1263/1263** pairs are both-set or both-clear, with zero half-empty pairs **Both Unknowns closed by [EXP-0080] at instruction level.** The six comparison codes are `== != > < >= <=` in exactly that order (the 6-entry table at `L02136`, each arm one `CMP` and one conditional jump), a code above 5 evaluates false, and the three pairs are **ANDed with short-circuit** — there is no OR. `+0xb4` is the **fire-once flag**: set, the pattern is skipped once its latch byte at `session+0xbec4` is non-zero; clear, the latch is re-armed every pass and the trigger fires again every full tick its conditions hold (`TRIG-CMP-006`, `TRIG-FIRE-007`).

**Confidence.** **Medium** (the id-vs-index question is decided by the bytes on 16 maps and the pair structure is exceptionless over 1263 pairs; no consuming instruction read) / **Unknown** (what each comparison code *means* — the corpus shows codes 3 and 5 sit almost entirely under distance and count checks and codes 0/1 under variable and boolean checks, consistent with `== != > < >= <=` in that order, but nothing here establishes it; the role of `+0xb4`, which is 0 on the single trigger of every standalone map; the 64 bytes at `+0x40`; and whether the pairs are ANDed)

**Original status.** ● partially retracted (opaque-region interpretation only)

**Evidence.** [EXP-0045](../experiments/EXP-0045-alm-triggers/)

**Amended.** The table ledger carried the status "● partially retracted (opaque-region interpretation only)". `retracted.md` records a correction against this claim.

### ALM-UNIT-048

**The type-6 record's two id words, corrected: `+0x40` (u16) is the UNIT id and `+0x42` (u32) is the GROUP id.** `ALM-UNIT-040` labelled `+0x42` "a unique id" from the loader's running-maximum instruction alone. It is not unique: on `scn:110` it takes 19 distinct values over 39 unit records while `+0x40` takes 39. The type-7 trigger vocabulary decides which is which — a `Target_Group` parameter resolves against `+0x42` (169/169) and a `Target_Unit` against `+0x40` — and it explains the instruction rather than contradicting it: an engine that spawns and re-populates groups at runtime needs to allocate **group** ids, so `mapObj+0x90` maximising `+0x42` is exactly right

**Confidence.** **Medium** (the distinctness counts and the 169/169 group resolution are whole-corpus and the rival — `+0x42` as the unit id — is refuted by both; but the *labels* rest on the trigger grammar of `ALM-TRIG-046`, itself Medium, rather than on an instruction that dereferences either field)

**Original status.** ● active

**Evidence.** [EXP-0045](../experiments/EXP-0045-alm-triggers/)

### ALM-TRIG-049

**type-9 = `[u32 count]` + `count` records of `26 + 6n` bytes**, `n` the `u32` at record `+0x16`. Head: `+0x00 u32 tag` · `+0x04 u32 X` · `+0x08 u32 Y` · `+0x0c u16 A` · `+0x0e u16 B` · `+0x10 u16 C` · `+0x12 u32 spellRaw` · `+0x16 u32 n`, then `n` elements of 6 bytes, each `(u16 kind, u16 low, u16 high)`. Consumes every payload exactly, **38/38**, including the three maps whose type-9 payload is the bare 4-byte count. **Amended by EXP-0227:** the prior apparent `u32` at `+0x0c` was the adjacent A/B words. X/Y are an eligibility pair; the former claim that the tail's first two words repeated them is refuted. For an eligible type-8 link, the consumer appends A, spellRaw and then the tail elements as ordered Effects (`ALM-EFFREC-071`)

**Confidence.** **Medium** for the unique wire shape, corpus, and linked field semantics; the semantic listing is not committed

**Original status.** ● active (amended)

**Evidence.** [EXP-0045](../experiments/EXP-0045-alm-triggers/), [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-TRIG-050

**type-8 has NO count word**: its record count is `meta+0x2c` (the loader's own case-8 bound, `ALM-META-025`), and the payload is that many records of **`20 + 10n`** bytes with `n` the record's **leading `u32`**. Head: `+0x00 u32 n` · `+0x04 u32` · `+0x08 u32 X` · `+0x0c u32 Y` (**`0x80`-centred `/256` anchors**, as `ALM-TRIG-022` said — `(X>>8, Y>>8)` in grid on **180/180**) · `+0x10 u32`, then `n` elements of 10 bytes. Consumes every payload exactly, **38/38**, including the eight maps whose type-8 is empty **[EXP-0080] adds the version gate the corpus could not show**: the loader reads the fifth head dword only when the map's `formatVersion >= 0x3dd` (989) and zeroes it otherwise (the version comparison at `L02137`), so the 20-byte head is a **990-only** shape and a consumer meeting an older map must read 16.

**Confidence.** **[EXP-0128] reads the arm and discharges every Unknown in this row.** Case 8 is `L02138` in the loader's own jump table at `L02078`; the head fields are `+0x00` the element count (a local, the loop bound at `L02139`), `+0x04` an actor id with **0 meaning "drop on the ground"**, `+0x08`/`+0x0c` the coordinates - now with the consuming instructions `L02140`/`L02141` shifting the coordinate down by 8 bits into a position byte - and `+0x10` **gold**, added to the sack's own purse at `L02142`. The 10-byte element is `u32` item code \| `u16` (read only on the actor arm) \| `u32` 1-based type-9 link. The corpus figure that stands unchanged is the closure; what changes is that it is no longer the evidence. **Medium** (the same three-parameter sweep returns **exactly one** closing model; the nearest fixed-stride rival — 30 bytes — closes 23/38, which is what makes the variable tail a measurement rather than a preference; corpus only) / **Unknown** (`+0x04`, `+0x10`, and the 10-byte element)

**Original status.** ● active (amended)

**Evidence.** [EXP-0045](../experiments/EXP-0045-alm-triggers/)

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-CLS-051

**`ALM-CLS-035`'s 64 residual type-3 cells are head-of-buffer residue, not placements — and the contaminated span is wider than 64 cells.** All 64 lie in **rows 0..2** of exactly two maps (`scn:131` 34, `scn:150` 30), i.e. inside the 8-cell border ring the ingest overwrites with block `0x1f` (`ALM-TERR-016`), so no mover ever stands on them. `scn:131`: read as `u16` the affected head holds **5** distinct nonzero words — 20, 84, 148, 212, 276, all blend variant 20 with strip groups 0..4 — and **56/56** occurrences are live type-1 tile words **of that same map's own tile grid**; read as `u8` placement codes the same bytes give 34 out-of-range values. `scn:150`: no `u16` tile structure, but the head is a nested set of symmetric radial profiles — 3 of its 6 nonzero runs of length ≥ 5 are exact palindromes, against a corpus base rate of **24/546 (4.40 %)** in rows 3.. of all 38 maps — the signature `ALM-GRID-013` identifies with the height grid. **Consumer consequence: bounds-check `code − 1`, and do not treat an in-range code in these two maps' rows 0..2 as a placement either**

**Confidence.** **Medium** for `scn:131` (56/56 is exceptionless and the rival reading produces 34 impossible values from the same bytes) / **Low** for `scn:150` (the palindrome and nesting signature is suggestive at 3/6 against 4.40 %, but the source buffer is not identified and the run appears nowhere in that map's own height grid) / **Unknown** (what wrote them, and whether `rom.exe` reads them at all — `R0379`'s handling of a code past the array end would settle the consumer question)

**Original status.** ● active

**Evidence.** [EXP-0045](../experiments/EXP-0045-alm-triggers/)

### ALM-CLS-052

**The type-6 → definition resolution, closed over the shipped file.** The search `R0495(L02110, k=(byte)file+0x08, k2=(byte)file+0x0a)`: when `k < 0x40 && k ∉ {0x1a,0x1b}` walk **Humans** from 1 matching `param[0x10] (typeID) == k`; otherwise walk **Units** from 1 matching `param[0x1d] == k && param[0x1e] == k2` (`typeID`,`face`); entries with empty names are skipped; 0 = not found. The override `R0498(L02110, file+0x10)` walks Humans **backwards** matching `param[0x18] (serverID)`. Corpus, 8094 type-6 records / 38 maps: 15 take the NPC arm (out of this DB's scope); of the 8079 in scope **all resolve** — 1490/1490 by `serverID`, 2/2 through the humans arm, 6587/6587 through the units `(typeID,face)` arm — and the `0x40`/`{26,27}` carve-out partitions the shipped populations exactly (no human `typeID ≥ 0x40` or ∈ {26,27}; no named unit `typeID < 0x40` outside them). `units.reg` cross-check: 33/34 IDs resolve through the same rule; ID 2 (`Unarmed Fighter with Shield`) has **no definition entry** and is placed by no shipped map

**Confidence.** High (both search bodies and their boundary immediates are read — the boundary constants are carried by the instructions alone, since the disjoint populations mean the corpus could not discriminate a wrong cut inside `(27,0x40)` — and the whole-corpus resolution is exceptionless; instrument for "all": `tools/placedb -mode verify` over the EXP-0030 framing) / Medium (that Data.bin `typeID` ≡ `units.reg` `ID` *as a general law* — exact on every placed key and on 33/34 registry IDs, and the drawable path consumes it as one, `ALM-CLS-054`; but nothing forces an unplaced future value to agree)

**Original status.** ● active

**Evidence.** [EXP-0049](../experiments/EXP-0049-placeable-db/)

### ALM-CLS-053

**The type-4 `kind` law, completed: `kind` subscripts the Buildings table of `Data.bin`, 1-based because the file skips entry 0.** All **3141/3141** shipped `kind`s land on a named building entry (1..66); the guard `kind ≤ R0487() = count−1 = 66` (`ALM-CLS-036`) is the file's own count (67) minus the skipped null entry; params consumed at spawn: `sizeX`/`sizeY` (footprint, overridden by the `kind == 0x21` extension), `scanRange` → `obj+0x48`, `healthMax` → `obj+0x44`/`+0x42`. The Data.bin names corroborate the `structures.reg`-`ID` identity name-for-name (`DAT-BLD-005`)

**Confidence.** High (rests on `DAT-GRAM-003`/`DAT-BLD-005`'s instruction-read grammar plus the exceptionless corpus)

**Original status.** ● active

**Evidence.** [EXP-0049](../experiments/EXP-0049-placeable-db/)

### ALM-CLS-054

**The client creation message's class byte becomes drawable `+0x20` at `L02143`, and its server source is `actor+0x0e`.** `R0509` takes the client Human path for `0x20 ≤ id < 0x40`; otherwise it bounds-checks the `units.reg` table and dispatches by class `Z`. On the server, both database streamers write their table `typeID` to `actor+0x0e`. A Humans constructor overwrites that word with `gender+0x21/+0x23` only when its constructor-mode argument is non-zero; definition-id and explicit typeID placements pass zero, while npc placements pass the exact `Hero` flag result (`PARTY-M20-031`). Thus `0x20..0x3f` is the player-character constructor-mode range, not the range of every Human placement

**Confidence.** High for both endpoint stores, the client gate and the conditional server producer; Medium only for the untraced message field between them

**Original status.** ● active (corrected)

**Evidence.** [EXP-0049](../experiments/EXP-0049-placeable-db/), [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/)

**Amended.** The table ledger carried the status "● active (corrected)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-REQ-055

**The open path requires exactly two records: `type1` and `type2`. Nothing else, and never a count.** `R0478` iterates the *file header's* `recordCount` (the loop-bound comparison at `L02144`), gated only `>= 3` (`L02079`, else status 3 “Wrong block number”) and `formatVersion <= 1001` (`L02145`, else status 4); it dispatches per record on the record header's dword 3 through the 10-entry table at `L02078`, and `typeId >= 10` is skipped by `Seek(payloadSize, SEEK_CUR)` at `L02146` with no error recorded and no counter kept. The code after the loop (`L02147`..`L02148`) is the whole contract: **five** tests, and only the first two are failures — `[this+0x0c] == 0` (type 1) → status **5** “Tiles block not found”, `[this+0x14] == 0` (type 2) → status **6** “Altitudes block not found”, each returning immediately. **Why those two and no others:** the consumer `R0278(world, map)` dereferences `map+0x0c` (u16/cell), `map+0x10` (i8/cell) and `map+0x14` (u8/cell) unconditionally in its per-cell loop, with no null test on any of them — the third is manufactured rather than demanded (`ALM-REQ-056`). The loader also never rejects a **duplicate** typeId: the case simply runs again, overwriting the pointer

**Confidence.** High (every clause is a named instruction with its immediates; the rival “the ten-record set is required” is killed twice — by the absence of any count test in the fully quoted post-loop block, and by a shipped four-record file the loader accepts (`ALM-CORP-060`). The *reason* for the two is the consumer's own unconditional dereference, not an inference from the shipped corpus)

**Original status.** ● active

**Evidence.** [EXP-0056](../experiments/EXP-0056-alm-record-set/)

### ALM-REQ-056

**Per-section behaviour on absence, section by section.** `type1`,`type2`: **fail** inside the loader (`ALM-REQ-055`). `type3`: **defaulted** — `malloc(<frame local at -0x4c>)` plus a byte-by-byte zero fill (`L02149`), where the frame local at `-0x4c` is the type-2 record's `payloadSize`, i.e. `W·H`; semantically exact, since `R0278` reads a nonzero byte of that plane as “an object blocks this cell” (`TERR-PASS-050`), so an all-zero plane is a valid empty one. `type5`: **defaulted** — `new(0x48)`, ctor `R0510(1,1)`, `+0x0c = 5000`, `+0x30 = 0`, appended to the list at `this+0x28` (`L02150`..`L02151`). `type0`: **defaulted** — `name[0] = 0`, text block `[0] = 0`, `this+0xd8 = 1`, `this+0xd4 = 1` (`L02152`..`L02153`), on top of the constructor's `W = H = 16`; but see `ALM-ORD-057`, because a type-0-less file with any of 1 or 4..9 present is undefined, not defaulted. `type4`,`type6`,`type7`,`type8`: **skipped**, their loops being bounded by type-0 count fields that are 0 when the records are absent. `type9`: **skipped**, its count being the first word of its own payload. `typeId >= 10`: **skipped by seek**. Corroboration that the count-vs-size coupling holds in shipped data: over 72 `.alm` files, 15 zero-length records exist (2 type-4, 13 type-8) and every one has a zero bound — **0 disagreements** either way; had any disagreed, the sequential parse would desynchronise and `ALM-FRAME-031`'s tiling would fail **Partially retracted by EXP-0358 (ALM-COUNT-195/ALM-STALE-196): type7 is not bounded by metadata+0x28 after complete reads, and the type0-less safety statement must not use that false count dependency for type7 or type9. Their own stream counts are separate; incomplete reads may preserve old local bytes. Other absence/default/corpus clauses stand.**

**Confidence.** High (each arm is a quoted instruction sequence, and the type-3 default's *semantic* exactness rests on the consumer's own test `if (cVar2 != 0)`) / **High** (the empty-record / zero-bound coincidence — not corpus agreement standing in for a mechanism but a consistency check between two independently read mechanisms, where a single disagreement would have broken the framing claim)

**Original status.** ● partially retracted

**Evidence.** [EXP-0056](../experiments/EXP-0056-alm-record-set/), [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### ALM-ORD-057

**The physical order `0,1,2,3,5,4,9,8,6,7` is not enforced, but three precedence relations are** — each carried by one stack slot or object field with exactly one writer. (a) **type-0 before type-1**: case 1 allocates and reads `this+0x00 * this+0x04 * 2` (the multiplication at `L02080` and the shift left at `L02081`), and `this+0x00`/`this+0x04` are written only by case 0, the constructor leaving them at `0x10`/`0x10` (`L02154`/`L02155`) — so a type-1 record ahead of type-0 gets a 512-byte grid. (b) **type-2 before type-3**: case 3's `malloc` and `Read` length is `[local-0x4c]` (`L02156`), written only by case 2 (`L02157`); ahead of type-2 it is uninitialised stack. The post-loop type-3 default reads the same slot and is safe only because the status-6 exit fires first whenever case 2 did not run. (c) **type-0 before types 4..8**: their loop bounds are the frame locals at offsets -0x54/-0x10/-0x14/-0x50/-0x18, written only by case 0's reads 8..12. The prologue reserves a `0x178`-byte frame with pre-loop stores to the locals at offsets -0x4, -0x1c, -0x20 and -0x58 and the `this` fields only — none of the five. The shipped order satisfies all three; the rest of the permutation the loader does not observe **[EXP-0128] adds two more, and they are cross-section pointer writes rather than stack slots:** case 6 indexes the type-8 list at `L02158` and case 8 indexes the type-9 list at `L02159`, so **9 must precede 8 and 8 must precede 6**. The shipped physical order is `0,1,2,3,5,4,9,8,6,7` on 38/38, which departs from the typeId sequence in exactly the two places a dependency exists (`ALM-ORD-068`). **Partially retracted by EXP-0358 (ALM-COUNT-195): clause(c) falsely includes type7 and says local -0x50 has only its metadata writer. Three internal type7 reads overwrite it. The metadata-count precedence subset is types4/5/6/8, not 4..8; type9 also supplies its own count. Other dimensions/planes and caster/loot/unit pointer-order clauses stand. Incomplete-read residue does not restore a format dependency.**

**Confidence.** High (each relation is a named `malloc`/`Read` length or loop bound plus an exhaustive account of that slot's writers, read from the listing rather than inferred; the rival “the whole permutation is enforced” is killed by the dispatch being on `typeId` and not on position, and the rival “nothing is enforced” by these three slots)

**Original status.** ● partially retracted

**Evidence.** [EXP-0056](../experiments/EXP-0056-alm-record-set/), [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/)

**Amended.** The table ledger carried the status "● partially retracted". `retracted.md` records a correction against this claim.

### ALM-META-058

**type-0 payload `+0x70` (→ `map+0xd4`) is a gate operand twice over, and the `+0x78` text region is consumed as one string.** (a) The `*.alm` **browser scan**'s accept test `R0479` ends `return (entry+0x0c < 2) ? 0 : 1` on the value read from payload `+0x70`; its caller `R0511` destroys the entry and never `PostMessage`s it on a `0`, so a map with `+0x70 < 2` is **silently absent from the list**. (b) The map-load path `R0128` sets `[[L00285]+0x0c] = (1 < map+0xd4)` and `[[L00285]+0x120] = ([+0x0c] == 0)`. Census over 72 files: **56/56** embedded campaign maps carry `1`; 15 of the 16 loose root maps carry 4, 8, 12 or 16, scaling with side (80×80 → 4, 144×144 → 8, 256×256 → 12 or 16); the sole loose map below 2 is `ru/Horror.alm`, whose 1 is a patch (`ALM-CORP-060`). **Extended by [EXP-0058]: the chain has a third link, and it is consumer-visible.** `[L00285]+0x14c` is written by exactly two instructions image-wide, both `= ([+0x0c] == 0)` (`L02160`, `L02161`), and its one consumer is the shop-by-id resolver, which returns the static town shop under it **without reading the id** — so on any map whose `+0x70` is 1, a shop placed on the map is built by the type-4 walker and can never be opened (`SHOP-ENTRY-016`). This field therefore decides three things, not two: whether the browser lists the map, `+0x120`, and which shop object a session reaches. `+0x74 <= +0x70` on 15/16 loose maps, `ru/Horror.alm` again the exception. (c) payload `+0x78`..`+0x277` is **one 0x200 read** into `map+0xdc`, and `R0479` reads the same block, rewrites every newline to `#` and assigns it as a single NUL-terminated CString

**Confidence.** High (both gates and the 0x200 read are named instructions; that a `+0x70 < 2` file is unlisted is the caller's own `if (iVar3 == 0)` arm) / **Medium** (calling `+0x70` a *participant count* — the two gates fix what it is compared against, not what it means, and the 72-file census supports the reading without discriminating it against a team or slot count. `+0x74` has no consumer in frame and is not named)

**Original status.** ● active

**Evidence.** [EXP-0056](../experiments/EXP-0056-alm-record-set/)

### ALM-RDR-059

**Four `.alm` readers ship, with three different acceptance tests — and only one of them acts on the loader's status word.** (1) `R0478`, the ten-case loader: magic, `recordCount >= 3`, `version <= 1001`, then `{type1,type2}` (`ALM-REQ-055`). Three callers: `R0128` (map load, via `R0512`), `R0414` (the world `Serialize`'s restore arm, `SAV-*`), and `R0277`, which `EnumRefs callto:` reports with **0 hits** including `.rdata` code pointers. **Only `R0128` tests anything** — `[map+0x14] != 0`, refusing at open with “Map error : ‹reason›, in ‹path›”; the other two hand the object straight to the world builder, so on the save-restore path an absent type-1/2 becomes a **null dereference in `R0278`**, i.e. downstream. (2) `R0479`, the `*.alm` browser's accept test: magic (a second “M7R” literal, `L02162`), `recordCount >= 2`, a type-0 record must exist, and `+0x70 >= 2` (`ALM-META-058`); it walks **to EOF**, not by count. (3) `R0480`, the light parser, reached from a `switch` on byte `+9` of the items `R0513` yields inside `R0509` (case 6 → `R0514`, which hands the object to `R0515` and `R0516` — **amended by [EXP-0298] (`TERR-LOAD-152`): only `R0515` receives the object; `R0516` receives a value read out of it, `P+0x28` (`L02163` reads it and `L02164` passes it as the argument), which is the terrain tile-group mask**): magic and `recordCount >= 2`, both reported by `AfxGetApp()->DoMessageBox` (`R0517`) which **does not stop the parse**; handles typeIds 0..3 only (case 0 reads 7 fields then falls through to a `Seek(0x25c)`, 28 + 604 = 632), seeks past everything else, walks to EOF, and finishes by clearing bit 13 across the type-1 grid. (4) `Map Editor.exe` carries **this** family's two literals (`0xe84a4`, `0xe84c4`) and none of the ten-case loader's six — a byte-level observation only; its loader is unread

**Confidence.** High (the three readers in `rom.exe` are read end to end, and the “only one caller checks” clause is the presence or absence of a compare at three named sites) / **Medium** (`R0277` being dead — instrument `EnumRefs callto:`, which sees `.rdata` code pointers and attributes orphan hits by address, but is blind to bytes Ghidra never disassembled) / **Unknown** (`Map Editor.exe`'s own reader: its own area, and the literal search bounds nothing but the strings)

**Original status.** ● active (amended)

**Evidence.** [EXP-0056](../experiments/EXP-0056-alm-record-set/), [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

**Amended.** The table ledger carried the status "● active (amended)".

### ALM-CORP-060

**`ru/Horror.alm` is `en/Horror.alm`'s first 262 876 bytes with two bytes changed — and it fixes what the 38-map corpus actually was.** EN `Horror.alm` (409 386 B, sha256 `426d7f02…c6719ccd`) is byte-identical to the GOG install's; RU `Horror.alm` (262 876 B, sha256 `153e76eb…db9ed272`) covers exactly the EN prefix ending at EN's record-4 header offset, and differs from it at **two offsets only**: `0x0c`, the file header's `recordCount`, 10 → 4; and `0x98` = type-0 payload `+0x70`, 16 → 1. Its four records — typeIds 0, 1, 2, 3 — satisfy every `ALM-FRAME-031` identity, `dataSize == 4·W·H + 72` included, and its type-0 record still advertises 7 type-5 / 415 type-4 / 1815 type-6 records that are not in the file — inert, since those figures are read only as the loop bounds of cases that never run. **Corpus scoping:** `ALM-SEC-003`'s 38 maps were the EN root's 10 loose files plus `scenario.res`'s 28; the RU root ships **6** loose files (Beast, Cross, Tomb, Kids2 dropped) and the same 28, so the union over both preserved roots is 44 distinct names / 72 walked files, of which **71 carry ten records and one carries four**. No reading of *why* the RU release ships it truncated is published

**Confidence.** High (hashes, a byte diff and the walk; nothing here is an interpretation)

**Original status.** ● active

**Evidence.** [EXP-0056](../experiments/EXP-0056-alm-record-set/)

### ALM-OBJ-061

**Which of a type-4 placement's two fixed-point coordinates is the column, read from the image end to end — `TERR-STRUCT-075`'s "the instructions do not say" is withdrawn.** The chain has no permutation in it and every hop is a named instruction. (a) The loader's case-4 read order fixes the locals: the read at `L02165` targets the local at frame offset -0x68 for file `+0x00` and the read at `L02166` the local at -0x74 for file `+0x04`. (b) the local at -0x68 is loaded at `L02167`, shifted down by 8 bits at `L02091` and pushed **last** before the call to `R0481` at `L02168`, so it is that routine's first stack argument, which `L02169` stores as a 16-bit value at the in-memory record's `+0x00`; file `+0x04` becomes `rec+0x02` by the same argument one push earlier. (c) `R0489` pushes `byte[rec+0x02]` then `byte[rec+0x00]` into `R0288` (`L02170`/`L02171`, and `L02172`/`L02173` on the other arm), and that routine pops `0xc` bytes of arguments on return and copies its first argument (`L02174`) into a byte of the object (`L02175`) — so **`rec+0x00` is position byte 0**. (d) `R0488` adds its inner counter, bounded by `obj+0x60`, to byte 0 (`L02176`) and its outer, bounded by `obj+0x61`, to byte 1 (`L02177`), then packs the key by shifting the first sum left by 8 bits (`L02178`) and adding the second — byte 0 is the key's **low** half. (e) The same low half is the tile grid's stride-1 axis: `L02179` reads the meta record's first dword straight into `mapObj+0x00`, and `R0278` bounds its **outer** counter by that word (`L02180`), multiplies the **inner** counter by it for the grid index (`L02181`, a multiplication by the 16-bit width word at `+0x50000`), and seeds the plane index from the outer counter (copied at `L02182`) stepping `+0x100` per inner iteration (`L02183`). So **type-4 `+0x00`, meta `+0x00`, the plane's low byte, `obj+0x60` and the Buildings column `sizeX` are one axis**; the other four fields are the other. What the instructions do *not* fix is which axis an authoring tool called "X" — that is nomenclature, and no consumer needs it. **The corpus separates the rivals too, contrary to `TERR-STRUCT-070`'s hedge**, once the statistic is re-made: off-plane cells cannot discriminate when 37 of 38 maps are square, but re-running the whole pass with the anchor transposed puts **178** footprint cells off the map, **29** inside the border ring, raises attach refusals **2 → 17** and collapses bridge joins **93 → 5**. Independently, the shipped `data.bin` masks are deck patterns only under this assignment: `Horisontal Bridge` (6×4) opens two full-width rows across its 6-cell axis and `Vertical Bridge` (4×6) two columns down its 6-cell axis — each named for the direction its deck then runs — while the transposed reading dissolves both into alternating noise and gives `Training School` a 4-wide row on a 3-wide building

**Confidence.** High (nine named instructions with their immediates form an unbroken chain; the rival needs a permutation — an `XCHG`, a swapped push, a second store into byte 0, or `mapObj+0x04` bounding the outer loop — and none exists over a chain read end to end; the corpus and the shipped mask geometry agree from two directions that are not the disassembly)

**Original status.** ● active

**Evidence.** [EXP-0091](../experiments/EXP-0091-placement-axes/)

### ALM-OBJ-062

**The type-4 extension's eight bytes, field by field — and three quarters of them are dead.** The loader reads all 8 into **one** local (an 8-byte read requested at `L02184` into the local at frame offset -0x80) and then pushes only two words out of it: the 16-bit word at the local's offset 0 (`L02185`) = file **`+0x14`** and the 16-bit word at its offset 4 (`L02186`) = file **`+0x18`**. Push order makes them `R0482`'s args 7 and 8, which `L02187`/`L02188` store at the in-memory record's `+0x0a` and `+0x0c`. `R0489`'s Building arm takes the **low byte** of each (the byte at record `+0xa` read at `L02189` and pushed as ctor arg 3, the byte at `+0xc` read at `L02190` as arg 4), `R0485` forwards them reversed (`L02191` then `L02192`), and `R0486` writes `obj+0x60` at `L02193` from arg 3 and `obj+0x61` at `L02194` from arg 4 (byte stores). **So file `+0x14` → `obj+0x60` (the `ALM-OBJ-061` low axis) and file `+0x18` → `obj+0x61`** — which is what `formats/alm/format.md` already printed, previously with no claim under it. Two riders the spec did not carry: **file `+0x16`, `+0x17`, `+0x1a` and `+0x1b` reach no instruction on this path**, and the effective width of each extent is **one byte**, twice over (`R0482` keeps a `u16`, the walker takes its low byte, `obj+0x60` is a byte). Corpus, all 8 extension records on 3 maps: every dword's upper three bytes are `00`, and 7 of the 8 are non-square (3×4, 4×5, 5×4, 7×3, 5×7, 3×3, 3×4, 3×6)

**Confidence.** High (six named stores and the two push orders that connect them; a transposed reading would require `R0482` to store arg 7 at `+0x0c` or `R0485` to push in the other order, both of which are single instructions that say otherwise) / **Medium**, and it is the corpus that is weak here (transposing the two extents leaves 5 of the 8 shipped bridges still joining their two banks — the shipped maps discriminate this only 3 times, so the reading rests on the image; and with every upper byte `00` the corpus cannot separate `u32`, `u16` and a byte at all)

**Original status.** ● active

**Evidence.** [EXP-0091](../experiments/EXP-0091-placement-axes/)

### ALM-CLS-063

**The `kind ∈ {0x22,0x23}` class DOES reach the footprint resolver — through its base-class constructor, which is why a call-site count could never show it.** `R0490`, the `0x74`-byte `Shop` constructor `R0489` selects at `L02195`, opens by calling the `Building` constructor: the two literal-zero arguments pushed at `L02196`/`L02197` · the position object · the kind byte · the call to `R0485` at `L02198`, and only afterwards installs its own vptr (`L02199` stores the vtable `R0491` in the object's first dword). `R0485` reaches `R0486` at `L02200`. So `TERR-STRUCT-072`'s **two call sites** is right and complete, and the resolver nevertheless has **three callers**: `R0489` twice over, once directly and once through `R0490`. `EnumRefs callto:` on the repaired table, 0 hits in orphan or undisassembled code: `R0486` 2 sites, `R0485` 2 (`L02201`, `L02198`), `R0490` 1 (`L02195`). Because the Shop pushes literal zeros for the two footprint bytes it always takes the resolver's **table** arm. Corpus: ids 34/35 are 3×3 with `BuildingPresent` 511 and `Passability` 447 — the full square attached, with bit 6 (the bottom-left cell) opened as the doorway. Over 11 maps the **50** placements attach **450** cells, block **400** and open **50**; deleting the arm frees **373** cells to a ground mover and merges a component boundary on two maps (`Horror.alm` 141 → 139, `Waters.alm` 33 → 32). The published 3141-placement / 17 057-cell figures were therefore right, but they assumed this rather than tested it

**Confidence.** High (one direct `CALL` immediate inside the constructor's own listing, which no enumeration instrument is needed to see; the rival — a Shop that initialises its own base — is excluded by that instruction, and its corpus cost is 450 cells)

**Original status.** ● active

**Evidence.** [EXP-0091](../experiments/EXP-0091-placement-axes/)

### ALM-LIM-064

**What a placement lets you customise, and what lifting each limit costs.** Decisive test: does lifting it change the bytes of a file the original wrote? (1) **Map extent 256×256** — position bytes are `u8` and the cell key is `u16`; the `.alm` already carries `u32` coordinates so no map file changes, but `TERR-PASS-053` packs the cell index into the top 16 bits of a save word, so a larger map changes the bytes of a `.sav` the original wrote: **semantic coupling, dangerous**. (2) **Footprint extents 255×255** — `obj+0x60`/`+0x61` are bytes; the extension's file fields are `u32` with three zero bytes in all 8 shipped records, so widening is **absent implementation, free**. (3) **Footprint mask 32 cells** — `obj+0x64`/`+0x68` are `u32` and `SHL` masks its count to 5 bits (`TERR-STRUCT-070`'s `Castle` aliasing); widening is **data width** in `world.res:data/data.bin`, and only the fixed classes need it because the extension arm writes both masks as constants. (4) **Buildings `sizeX`/`sizeY` 255** — read as `byte [EAX]` from the table: **data width** in `data.bin`. (5) **Class id 255** — `obj+0x40` is a byte (`L02202`) and the guard is `kind ≤ count−1`; the `.alm` field is `u32` and the shipped domain is 1..66, so classes 67..255 are a **compiled constant, free**. (6) **Only `kind == 0x21` may be variable-size** — the immediate is free to change, but the type-4 stream's record **length is inferred from `kind`** (`L02098`) and there is no per-record "has extension" bit, so making a second class variable-size makes the payload untileable by every reader that exists: **semantic coupling, dangerous**. (7) **The shop doorway** is one clear bit of `Passability` in `data.bin` — **data width**, free to move, and it is the only cell of a shop a unit can stand on

**Confidence.** Medium (each width and guard is a named instruction, but "does the byte of a shipped file change" is an argument about consequences over the corpus we hold, not a measurement — no modified file was written or loaded)

**Original status.** ● active

**Evidence.** [EXP-0091](../experiments/EXP-0091-placement-axes/)

### ALM-SACK-065

**The `.alm` type-8 section is the map's authored loot, and every field has a consuming instruction.** The record loader `R0478` dispatches at `L02203`, an indexed jump through the table at `L02078`; the table read out of the image gives arm 8 = `L02138`, and that arm allocates `0x4c` bytes per record (size `0x4c` passed at `L02204`, ctor `R0518`), reads five head fields through the stream's virtual read slot at vtable offset `+0x3c`, and appends the record to **`map+0x2dc`** (call to `L02205` at `L02206`). Case 9 appends to `map+0x2f0` (`L02207`). The consumer is **`R0461`**, whose only caller is `L02208` inside the `.alm` loader `R0128` (`callto:R0461` = 1 hit, 0 orphan) and which is **unconditional** there: its first loop walks `map+0x2f0`, its second (`L02209`..`L02210`) walks `map+0x2dc` and calls the sack maker `R0472` at `L02084`. **File `+0x00` `u32 n`** is the element count, a local, and the bound of the loop at `L02139`. **`+0x04` -> `O+0x3c`** (`L02211`): `L02212` tests it and a sack is made only when it is **0**; non-zero is a key into a map the routine fills from the actor registry `[L00240]` as `map[*(actor+0x8)] = actor` (`L02213`..`L02214`). **`+0x08` -> `O+0x40`** (`L02215`): `L02140` reads `O+0x40` and `L02216` shifts it down by 8 bits, stored as a byte and pushed **last** into `R0288`, so by `ALM-OBJ-061`'s argument-copy routine (pops `0xc` bytes, copies the first argument into a byte) it is **position byte 0**. **`+0x0c` -> `O+0x44`** (`L02217`), the same three instructions at `L02141`/`L02218`, position byte 1. **`+0x10` -> `O+0x48`**, read only when the version comparison at `L02137` passes: `L02219` pushes it as `R0472`'s third argument, and `L02220`, `L02142` and `L02221` add it to the sack's own gold field (`+0x3c`), so it is **money**. Then `n` elements of 10 B: **`+0x00 u32`**, whose low `u16` is the item code (a 16-bit read at `L02222`, passed into `R0519` on the `Data.bin` root `L02110`); **`+0x04 u16`**, read at `L02223` on the `owner != 0` arm **only**; **`+0x06 u32`**, a **1-based** index into the type-9 list (decremented by 1 at `L02224` before `map+0x2f0` `ElementAt`), 0 meaning none. `R0472` looks the cell up first (`R0446`, `TRIG-SACK-022`'s lookup) and on a miss allocates `0x44`, `ITEM-SACK-010`'s `Sack`; on a hit it merges the container and adds the gold, so two records on one cell become one sack

**Confidence.** High (every field is named with both the instruction that reads it and the instruction that consumes it; the two call enumerations are complete on the repaired table with `.rdata` slots included and 0 orphan hits; the nearest rival record shape - a fixed 30-byte stride, indistinguishable on any map whose records each carry one element - is excluded by the image rather than by the corpus, because case 8's element-loop bound is a value read from the stream at `L02225`, not a constant)

**Original status.** ● active

**Evidence.** [EXP-0128](../experiments/EXP-0128-alm-items-section/)

### ALM-SACK-066

**The corpus, and the three properties that discriminate the reading from its rivals.** Both preserved roots, `ALM-CORP-060`'s map set: **30 of 38** maps on the larger root carry a type-8 record and **28 of 34** on the smaller; the walk consumes the payload exactly **38/38** and **33/33**; **137** / **133** records have `+0x04 == 0` and are therefore sacks on the ground, **43** / **43** have it non-zero; **181** / **177** elements in all; **636 074** / **125 074** gold over the ground records; **0** cells outside `[0,W)` x `[0,H)` before the byte store and **0** the byte store would wrap; **0** of 74 type-9 links out of range; `formatVersion` is `0x03de` on **every** shipped map, so the 16-byte head the `>= 0x3dd` gate implies is never exercised by shipped content. The two roots differ by exactly the two loose maps the smaller drops (3 sacks / 11 000 gold and 1 sack / 500 000 gold); the 28 archived campaign maps agree. **Three properties a wrong reading would have failed.** (a) *Axes.* Over the corpus **23 of 28** check-opcode-14 nodes - `TRIG-SACK-022`'s cell-to-sack existence test, whose two parameters are read as bytes - name a cell a type-8 ground record occupies, and **0 of 28** do so with the two axes swapped. (b) *The element boundary.* The `+0x04` field is read by exactly one arm of the image, the `owner != 0` arm at `L02223`; in the shipped data it is non-zero on **5 of 25** elements of `owner != 0` records and **0 of 156** of `owner == 0` records, so the authored bytes respect a boundary only that code path reads. (c) *The class nibble.* `R0520` returns null for a class of 0 or 15, and a wrong nibble position would put about an eighth of the codes there; over **181** elements there are **0** nulls and **0** elements with any bit set above the low `u16`. **Five live check-opcode-14 nodes name a cell no record occupies** and are not explained: one is cited by no trigger and is a dead editor node, four are cited - `scn:80` "Get sack1" at (98,13), `scn:90` "Get sack 97-132" at (97,132), "Get sack 66-33" at (66,53) and "Get sack 22-113" at (22,113), the last two disagreeing with their own author-written labels

**Confidence.** High for the three discriminating properties (each is a prediction the shipped bytes could have refuted and did not, and each comes from a different mechanism - a runtime lookup, a single-armed field read, and an allocation guard) / **Medium** for the census figures themselves, which are corpus agreement

**Original status.** ● active

**Evidence.** [EXP-0128](../experiments/EXP-0128-alm-items-section/)

### ALM-LIM-067

**What an authored-loot record lets a consumer customise, and what lifting each limit costs.** Decisive test: does lifting it change the bytes of a file the original wrote? (1) **Position 0..255 per axis** - the file carries `u32` fixed point but the shift down by 8 bits at `L02216` is stored into a byte, and the sack's cell key is `TRIG-SACK-022`'s `u16`; widening changes no map file, but `TERR-PASS-053` packs a cell index into a save word, so it is the same **semantic coupling** `ALM-LIM-064` already records for placements - dangerous. (2) **Records per map 2^32-1** - the bound is `meta+0x2c`, a `u32` in the type-0 record with no compiled cap; shipped maps use 0..23. **Absent limit, free.** (3) **Elements per record 2^32-1** - likewise a stream `u32`; shipped records use 0..3. **Absent limit, free.** (4) **Gold `u32`** - added straight into the sack's gold field; shipped values reach 500 000. **Absent limit, free.** (5) **Item classes 1..14 only** - `R0520` is a compiled if-chain returning null outside that window, and the shipped domain is {1,2,4..10,12,14}; classes 0, 3, 11 and 13 are a **compiled constant**, free to give a meaning, while 15 and above are unreachable because the field is four bits. (6) **Item index 5 bits for classes 1..13, 8 bits for class 14** - the mask `0x1f` at `L02226` against the mask `0xff` at `L02227`; widening the first is a **compiled mask**, free, and it is the sharpest limit in the record: 32 rows per class against 256 for magic items. (7) **One sack per cell** - `R0472` merges on a hit rather than refusing. That is behaviour rather than a width, but it means no `.alm` can author two distinguishable sacks on one cell, and a consumer that wants them needs a second mechanism. (8) **The `+0x06` link is 1-based**: file value 0 means none, while file value 1 addresses zero-based type-9 index 0 (`ALM-EFFLINK-072`)

**Confidence.** Medium (each width and guard is a named instruction and each shipped domain is measured, but no modified `.alm` was written or loaded, so every cost is an argument about consequences over the corpus we hold)

**Original status.** ● active (amended)

**Evidence.** [EXP-0128](../experiments/EXP-0128-alm-items-section/), [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-ORD-068

**The shipped physical record order is `0,1,2,3,5,4,9,8,6,7`, and it is not a writer's habit - two forward references in the loader require it.** `ALM-ORD-057` established three precedence relations carried by stack slots; there are two more, and both are cross-section pointer writes. Case **6** at `L02228`..`L02127` takes the unit object's `+0x3c` word, indexes `map+0x2dc` - the **type-8** list - by the placement's already-decremented `+0x18` link, and stores the id into that record's `+0x3c`, the same in-memory slot the type-8 record's own file `+0x04` lands in (`ALM-SACK-065`); the guard is `link < meta+0x2c` and the else arm writes `unit+0x44 = -1` (`L02229`). Case **8** at `L02230`..`L02231` indexes `map+0x2f0` - the **type-9** list - by an element's `+0x06` and writes both the element index and the owning type-8 object into it. Neither indexed read is defined unless the filling case has already run, so **9 must precede 8 and 8 must precede 6**. The corpus obeys exactly that: **38 of 38** maps on the larger root and **33 of 33** full maps on the smaller carry the physical order `0,1,2,3,5,4,9,8,6,7` (the 34th is `ALM-CORP-060`'s truncated four-record file, `0,1,2,3`). The two orders differ from the typeId sequence in two places, and **both** are exactly where a dependency exists - `5` before `4`, and `9,8` before `6`. **The cross-write is a bijection with a file field.** Over the whole corpus there are **43** type-6 records with a non-zero `+0x18` and **43** type-8 records with a non-zero file `+0x04`, and each of the 43 is the one the other names: **0** disagreements in either direction, on both roots. Because case 8 runs first, the file's `+0x04` is then **overwritten** by case 6 for every linked record, so the engine's effective value is the type-6-derived one and the file field is redundant on exactly the 43 records it is non-zero on. A consumer needs only one of the two, but must not read a non-zero `+0x04` as a ground sack

**Confidence.** High for the two precedence relations (each is an indexed read of a list another case fills, both named instructions) and High for the bijection (43 = 43 with 0 errors either way over 180 records is a prediction the shipped bytes could have refuted in two directions and did not) / **Medium** for "the order is required rather than merely observed": nothing here loads a permuted file, and the loader would not reject one - it would read an unfilled list

**Original status.** ● active

**Evidence.** [EXP-0128](../experiments/EXP-0128-alm-items-section/)

### ALM-PLAYER-069

**The type-5 record's `+0x00` word is the faction's colour slot and its `+0x04` word is `Player+0x28`; the word the owner lookup matches is neither, it is the record's own 1-based ordinal.** `R0478` case 5 loads each 76-byte record into a `CPlayer` and stores the ordinal itself (a dword store at `L02232` of `i + 1` into `CPlayer+0x4`), then appends to the list at `map+0x28` (`L02233`). At map load `R0423` walks that list and, for each element, fills a session `Player`: `L02128` stores `Player+0x8` from `CPlayer+0x04` (**the ordinal**), `L02129` stores the byte `Player+0x44` from `CPlayer+0x08 + 1` (**ALM `+0x00`**, the colour), the name from `CPlayer+0x10`, and `L02131`/`L02132` store the dword `Player+0x28` from `CPlayer+0x30` (**ALM `+0x04`**). The owner lookup `R0425` compares at `L02234` the dword `Player+0x8` with the sought ordinal, that is `Player+0x08`, the ordinal — so `ALM-GRP-041`'s appositive on `+0x00` names the wrong field, while its field mapping is exact. `Player+0x44` is the colour byte the join sender puts on session message `0x96` as `+0x44 - 1` (`PAL-SHADE-013`), and that byte becomes the shade slot the unit body draw indexes (`PAL-SHADE-012`, `PAL-OWN-007`). **Corpus, both preserved roots plus the live install:** 191 type-5 records over 38 EN maps, 166 over 34 RU maps (`ru/Horror.alm` carries no type-5 record, `ALM-CORP-060`). `+0x00` takes only `0..10` and one `13`; 0 records exceed 15, which matters because the draw indexes the 17-entry table with it unmasked (`PAL-LIMIT-017`). 48 records over 20 of the 38 EN maps carry a `+0x00` that differs from their own ordinal, and three maps carry two records with the same `+0x00` — `scn:131.alm` (`Knights`, `ghost`, both 6), `scn:50.alm` (`Wild monsters`, `Mage`, both 6), `scn:90.alm` (`Enemies`, `Dragon`, both 9). `+0x04` takes only 0 and 1, and the split is exact: **every one of the 28 embedded campaign maps authors 0 on record 0, every loose map authors 1**, and every record above ordinal 0 authors 1 on every map. Record 0 is named `Self` or `self` on all 38. This closes `ALM-GRP-041`'s Unknown on `+0x04`; ~~`+0x08`, the `{0,5000}` scalar, is still read by nothing~~ -- *[EXP-0297 correction]: the earlier replacement was still too broad. The selected population is 17 map-side plus 7 CPlayer-family routines. The family has one copy read; its constructor has two direct call sites inside an owner with no direct references. No map-side direct-displacement read was attributed, but computed pointers and whole-object movement remain outside the search, so consumer effect is **Unknown** (`ALM-SCALAR-087`, `ALM-SCALAR-088`). `R0423` itself does not touch `CPlayer+0x0c`*

**Confidence.** High (the roster/colour chain is a named-store reading and the corpus figures are a complete walk of every `M7R\0` payload of each root) / **Medium** (that the `+0x00` word reaches the wire in a shipped single-player mission: the join sender takes the `+0x44` arm only when `Player+0x44` is already non-zero, and whether map load precedes the join is an ordering static reading does not settle — `PAL-SHADE-013`)

**Original status.** ● active (amended)

**Evidence.** [EXP-0180](../experiments/EXP-0180-shade-select/), [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-MODE-070

**The type-0 payload `+0x70` scalar is the map object's `+0xd4`, and its consumer is one test: the session's multiplayer flag is `(map+0xd4 > 1)`.** The ALM reader's type-0 case writes it at `L02235` (a 4-byte read, `&map+0xd4`), bracketed by the 64-byte name at `L02236` (a 64-byte read, `&map+0x94`, payload `+0x30`) and `+0x74` at `L02237`. The map loader `R0128` then computes the signed test `map+0xd4 > 1` at `L02087`/`L02088` and stores the result at `L02089` into the dword at `+0xc`, so `[L00285]+0xc` — the flag `R0423` reads at `L02238` before deciding whether the first faction reuses the existing session `Player` — is that scalar's `> 1` test. Corpus over both preserved roots and the install: every one of the 28 `scenario.res` maps carries 1 on every root. The converse does not hold, and the counterexample is a deliberate edit rather than an accident. `ALM-CORP-060` established that `ru/Horror.alm` is `en/Horror.alm`'s first 262 876 bytes with **two** bytes changed: the file header's `recordCount` at `0x0c`, 10 → 4, and `0x98`, which is type-0 payload `+0x70`, **16 → 1**. So the one payload field the RU release altered on that map is this one, and altering it is what makes the engine load the map with the multiplayer flag clear. `+0x70 == 1` therefore identifies a map that loads single-player, not a map stored in `scenario.res`. EN and install histogram `1:28 4:2 8:2 12:1 16:5` over 38 maps; RU `1:29 4:1 8:2 16:2` over 34. This identifies the consumer of the scalar `ALM-META-026` left with its role undetermined; the wider reading of the value space as a player-slot count is untouched

**Confidence.** High (the reader's push, the loader's compare and the flag store are named instructions re-read from raw image bytes on three roots with 0 mismatches; the one-file counterexample is an exact measurement over both roots, which is what discriminates against the campaign/standalone reading)

**Original status.** ● active

**Evidence.** [EXP-0181](../experiments/EXP-0181-owner-colour-rule/)

### ALM-EFFREC-071

**Every version-990 type-9 record has one exact `26+6n` wire frame; an eligible linked zero-coordinate record separately supplies an ordered map-item Effect recipe.** The universal wire form is `u32 tag, u32 x, u32 y, u16 A, u16 B, u16 C, u32 spellRaw, u32 n, n*(u16 kind,u16 low,u16 high)`. On the linked arm, the consumer appends nonzero A as kind `A+43`, nonzero spell low word as kind 41 (42 for Book), then each element in file order, remapping element kind 41 to 49. The walk tiles 199 campaign type-9 records per root and every loose record. The Effect statement is restricted to the linked zero-coordinate arm and does not reclassify the section's separate caster arm

**Confidence.** **Medium** (universal framing closes the complete corpus; linked semantics come from a full static reading whose semantic listing is not committed)

**Original status.** ● active

**Evidence.** [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/)

### ALM-EFFLINK-072

**A type-8 element's type-9 link is 1-based, zero means none, and the linked operation is the same for ground and actor stock.** The loader and consumer each subtract 1 before indexing `map+0x2f0`; the loader writes element ordinal and owning type-8 pointer back into runtime type9 `+0/+4`. Consumption requires type9 x/y zero and that backlink nonzero. The complete Effect list is built before the owner-zero branch, so both destinations receive the same enchanted object. This excludes 0-based and embedded-id models through two asserted decrement sites; the destination statement comes from the read consumer body

**Confidence.** High for 1-based arithmetic (both decrement byte ranges are asserted on both executables) / **Medium** for backlink, complete construction and shared destination, whose semantic listing is not committed

**Original status.** ● active

**Evidence.** [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/)

**Amended.** `ALM-T9TAG-312` commits the case-8 loop (`w18`): for each linked element the loader stores the element ordinal at `T9(link-1)+0x00` (`L15874`) and the type-8 object at `+0x04` (`L02231`). The backlink write, which this card grades Medium for want of a committed listing, is now read in a committed listing; the grade here is not changed.

### ALM-EFFPOP-073

**Every positive type-8→type-9 link in the shipped campaign is enumerated: 74 per root.** Across 28 maps per root there are 176 type-8 records / 177 elements and 199 type-9 records; all 74 links are in range, every target has x=y=0, and every linked type-9 record has exactly one reference. Three targets carry nonzero file tags 1, 2 and 4, which do not gate application because the loader overwrites runtime `+0`. EN/RU linked rows and constructed effect sequences agree

**Confidence.** **Medium** (complete corpus with exact tiling and explicit alternatives, but corresponding campaign content supplies parity rather than independent variation)

**Original status.** ● active

**Evidence.** [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/)

### ALM-M40-074

**Mission 40's type-8 element at cell `(96,84)` is record 2 element 1, code `0x0e1f`: class 14, runtime row 31, `Quest Item31`, value 10,000, link zero and no linked Effect.** Both preserved campaign archives carry the same relation. The positive value is ordinary class-14 row data (`ITEM-MAGVAL-090`), not evidence of enchantment

**Confidence.** **Medium** for the exact two-root corpus identity / High for zero taking no link operation (`ALM-EFFLINK-072`)

**Original status.** ● active

**Evidence.** [EXP-0227](../experiments/EXP-0227-authored-enchantments-death-drops/)

### ALM-TAILMAP-079

**The type-6 record's file→runtime map, re-derived: 28 stream reads, not 27, and 18 of them are not the identity.** `R0478` case 6 (`L02239`…`L02240`) allocates `0x50` per record and performs **28** reads (a width argument followed by a call through the stream's virtual read slot at vtable offset `+0x3c`) whose widths sum to exactly 70 at `formatVersion == 990`; deduplicated and enumerated by address they run `L02241`…`L00297`. Ten of the 28 mappings are the identity and **18 are not**. Seventeen land in the record at a different offset: file `+0x0c` → `rec+0x48`, `+0x18` → `rec+0x44` (then `−1` at `L02242`), `+0x28`/`+0x2a` → `rec+0x4c`/`+0x4e`, the nine bytes `+0x2c`…`+0x34` → `rec+0x28`…`+0x30`, `+0x35` → `rec+0x31`, `+0x3b` → `rec+0x37`, `+0x40`/`+0x42` → `rec+0x3c`/`+0x40`. The eighteenth is `+0x0a` → `[local-0x98]` → `rec+0x0c`. Relocating the version-gated pair to `rec+0x4c`/`+0x4e` shifts file `+0x2c`…`+0x40` by `−4` and the `u32` at `+0x42` by `−2`, so the six skill bytes at file `+0x35` are `rec+0x31`…`+0x36` and the five resistance bytes at file `+0x3b` are `rec+0x37`…`+0x3b`. Three reads land in a stack local first, and only the second of them is not the identity: `+0x08` → `[local-0x9c]` → `rec+0x08` (`MOVSX`, `L02243`), `+0x0a` → `[local-0x98]` → `rec+0x0c` (`L02244`), `+0x14` → `[local-0xa0]` → `rec+0x14` (`L02245`). The three version gates reproduce `ALM-UNIT-040`'s `0x3db` / `0x3b7` / `0x3d9`. **`UNIT-PLACE-034`'s file attribution, which its own confidence line grades Medium, survives**: the nine-byte file run at `+0x2c` does land at `rec+0x28`…`+0x30`, so Body ← file `+0x2c`, Mind ← `+0x2d`, Spirit ← `+0x2e`, Reaction ← `+0x2f`. `ALM-UNIT-040`'s "27 sequential reads" is one short; its byte total and every field label stand

**Confidence.** **High** for the count and the map — the reads are an address-by-address enumeration of one idiom over one range, and the live alternative (`ALM-UNIT-040`'s 27) is excluded by that enumeration rather than by re-summing widths / Medium for the map being the same at every version the loader accepts: every shipped map is at 990, so the gated branches were read and never exercised

**Original status.** ● active

**Evidence.** [EXP-0296](../experiments/EXP-0296-unit-placement-tail/)

### ALM-TAILHOLD-080

**Exactly three routines can hold a pointer to a type-6 runtime record, and between them they dereference 20 distinct displacements of it.** Case 6 stores each `0x50`-byte record's **pointer**, not its bytes, into the 4-byte-element `CObArray` embedded at `mapObj+0x318` (`L02246` offsets to `0x318` and `L02247` calls `L02248`; the element address is `this+0x04` plus 4 times the index in `R0521`). `imm:318` returns 19 hits over 11 owners and `disp:318` 9 over 8; after classification the owners of *this* array are `R0478` (loader), `R0151` (placement spawner) and `R0522` (map destructor). The rest are the array's construction (`R0523`, `R0504`, `R0524`), five code-region unwind funclets, two hits on a type-7 796-byte node base (`L02249`, `L02250`), and the unrelated code-region family. `callto:` on `R2121`, `R1338`, `R0521`, `L13132` and `L02248` adds no owner. In the spawner, **50** loads of the record slot resolve to a dereference and **0** escape — the pointer is never pushed as an argument and never used as `this`. The loader dereferences 12 times, at `+0x08`, `+0x0c`, `+0x14`, `+0x3c`, `+0x40`, `+0x44` only, and escapes twice: the base handed to the stream reader (`L02251`) and the array store (`L02252`). The destructor dereferences **nothing** — it fetches the element and frees it (`L02253`…`L02254`, a call to `R0525`). Union of every displacement dereferenced over the record's whole lifetime, the two indexed reads shown with their index range: `+0x00`, `+0x04`, `+0x08`, `+0x0c`, `+0x10`, `+0x14`, `+0x20`, `+0x24`, `+0x28`…`+0x2b`, `+0x2d`, `+0x2e`, `+0x31+i` (`i=1..5`), `+0x37+i` (`i=0..4`), `+0x3c`, `+0x40`, `+0x44`, `+0x48`. This is the reduction rule every `disp:` negative on this record rests on: a displacement is not a type, so a sweep hit counts only when its owner can hold a record pointer

**Confidence.** **High** for the holder set reachable through the container — the alternative, a fourth routine reading the array, is excluded by the `imm:`+`disp:` enumeration of every `+0x318` site, and a fourth routine receiving a record pointer by argument is excluded by the spawner's zero escapes and the destructor's single escape to `free` / Medium for the stronger reading that no routine obtains one by any other route: case 6 hands a record-interior pointer to the stream's read slot (vtable `+0x3c`) 28 times per record and that routine was not read

**Original status.** ● active

**Evidence.** [EXP-0296](../experiments/EXP-0296-unit-placement-tail/)

**Amended.** `ALM-UNITREAD-227` reads the stream read slot (`R0533`, `R0534`) that this card lists as unread: it keeps no record pointer. `ALM-UNITREAD-313` then reads the array append `L02248`, `R2270` and the accessors `R1338`, `R2121`, `R0521` and `L13132`: none dereferences a record, except that the grow routine `L15875` is unread. The grades stay as graded.

### ALM-TAILDIR-081

**The type-6 file `+0x1c` is a `u32` whose domain is exactly `0..14`, and no routine that can hold a record pointer reads where it lands.** Read at `L02255` into `rec+0x1c`. Over 12 085 records of both preserved roots it takes 15 distinct values, `0..14` with no gap and nothing above. Even values dominate by margins that differ sharply between the roots: EN authors 6967 even against 1127 odd (`0:1203 2:735 4:611 6:1023 8:816 10:1143 12:710 14:726` against `1:140 3:134 5:174 7:203 9:159 11:148 13:169`, 5.1× counting only the nonzero), RU 3857 against 134 (24.5× nonzero), although both roots ship the same `scenario.res`. The per-map parity split was not measured, so which maps carry the odd values is Unknown. 6891/8094 EN and 3421/3991 RU records are nonzero. `disp:1c` returns 3276 hits over 1236 owners image-wide, of which exactly two are owned by a routine that can hold a record pointer — `L00292` and `L00296`, both based on `[EBP-0x10]`, the object the same block compares against `rec+0x40` at `L00293`. Ten hits sit in orphan runs and are listed rather than dropped; four of those are `[ESP + 0x1c]` stack locals. The surviving explanation is a field written for a consumer outside this image: `Map Editor.exe` carries a second, unread `.alm` loader (`ALM-RDR-059`)

**Confidence.** **Medium** — the negative is bounded by `ALM-TAILHOLD-080`'s reduction rule and inherits its stream-reader blind spot, and the domain is corpus agreement at one `formatVersion`

**Original status.** ● active

**Evidence.** [EXP-0296](../experiments/EXP-0296-unit-placement-tail/)

**Amended.** `ALM-UNITREAD-227` reads the stream read slot (`R0533`, `R0534`) that this card lists as unread: it keeps no record pointer. The grades keep the remaining caveat, array append `L02248` and the accessors `R1338` and `R2121`, which `ALM-UNITREAD-227` also names.

### ALM-TAILU16-082

**The four type-6 `u16` at file `+0x20`…`+0x27` use `0xFFFF` as the absent value, and only two of the four are read by anything that can hold the record.** `+0x20` and `+0x24` are read behind `MOVSX` / `CMP -1` / `JZ` at `L02256` and `L02257` and reach `actor+0x94` and `actor+0x9a` (`UNIT-PLACE-034`). `+0x22` and `+0x26` have none: `disp:22` returns 20 hits over 8 owners with 2 in orphan runs and `disp:26` 11 over 7, and no hit of either is owned by a record-pointer holder. Corpus, 12 085 records over both roots: `+0x24` and `+0x26` are `0xFFFF` on **12 085/12 085**, so the mana override has a consumer no shipped map exercises. `+0x20` and `+0x22` are authored on exactly 30 records per root over the same seven `scenario.res` maps (`scn:30` ×4, `scn:71` ×9, `scn:81` ×1, `scn:100` ×1, `scn:120` ×1, `scn:131` ×5, `scn:140` ×9), and the cross-tab is per record, not per map: 30 both-present / 8064 both-absent on EN and 30 / 3961 on RU, so neither is ever authored without the other. They are not copies of each other — `+0x20 ∈ {0, 65526}` against `+0x22 ∈ {0, 65526, 10, 5}`

**Confidence.** **Medium** — every count is corpus agreement at one `formatVersion`, and the two negatives inherit `ALM-TAILHOLD-080`'s reduction rule and its blind spot

**Original status.** ● active

**Evidence.** [EXP-0296](../experiments/EXP-0296-unit-placement-tail/)

**Amended.** `ALM-UNITREAD-227` reads the stream read slot (`R0533`, `R0534`) that this card lists as unread: it keeps no record pointer. The grades keep the remaining caveat, array append `L02248` and the accessors `R1338` and `R2121`, which `ALM-UNITREAD-227` also names.

### ALM-TAILRUN-083

**`ALM-UNIT-018`'s `0xFF`-sentinel inventory reading of the type-6 byte runs is wrong on both the values and the role.** No byte of file `+0x2c`…`+0x3f` is `0xFF` on any of the 12 085 type-6 records of the two preserved roots — the same population that row was measured over. The record's only `0xFFFF` sentinels are the four `u16` at `+0x20`…`+0x27` (`ALM-TAILU16-082`), which is where `EXP-0019`'s pre-`EXP-0030` offsets land after that correction's −8 shift. The two runs are stat vectors: file `+0x35` (6 B) → `rec+0x31` is the skill override run (`UNIT-PLACESKILL-086`) and file `+0x3b` (5 B) → `rec+0x37` the elemental-resistance override run (`UNIT-PLACERESIST-087`), each read by one indexed loop of `R0151` behind a zero guard and stored into the actor's own stat vectors; no inventory structure is touched by either. The shipped corpus authors them sparsely: `+0x35`…`+0x3a` are nonzero on 3, 4, 4, 4, 4 and 119 records and `+0x3b`…`+0x3f` on 7, 6, 8, 26 and 4 (EN). The three byte positions between and after the runs are the record's remaining holes, each with its own bounded negative and each `0` on 12 085/12 085 records: file `+0x30` → `rec+0x2c` (`disp:2c`, 1025 hits over 423 owners, 3 in orphan runs, **0** owned by a record-pointer holder), file `+0x33` → `rec+0x2f` (`disp:2f`, 13 over 10, 0 holder hits), and file `+0x34` → `rec+0x30` (`disp:30`, 1206 over 457, 2 in orphan runs, and **2 hits inside a holder** that are not record reads — `L02146`, an indirect call through a vtable slot at `+0x30`, and `L02258`, a dword store of 0 at `+0x30` — both outside case 6's `L02239`…`L02240` and neither among the loader's 12 resolved record dereferences, which use `+0x08`, `+0x0c`, `+0x14`, `+0x3c`, `+0x40` and `+0x44` only)

**Confidence.** **High** for the role — the two loops' own destination stores — and for the `0xFF` census, which is complete over the population the refuted row used and would be overturned by one `0xFF` byte anywhere in that 20-byte window / Medium for the claim that no other `formatVersion` uses those bytes differently, and for the three holes' negative, which is bounded by `ALM-TAILHOLD-080`'s reduction rule and inherits its stream-reader blind spot

**Original status.** ● active

**Evidence.** [EXP-0296](../experiments/EXP-0296-unit-placement-tail/)

**Amended.** `ALM-UNITREAD-227` reads the stream read slot (`R0533`, `R0534`) that this card lists as unread: it keeps no record pointer. The grades keep the remaining caveat, array append `L02248` and the accessors `R1338` and `R2121`, which `ALM-UNITREAD-227` also names.

### ALM-TAILVER-084

**The version-gated type-6 pair at file `+0x28`/`+0x2a` lands at `rec+0x4c`/`rec+0x4e`, and no routine that can hold a record pointer reads either.** Read at `L02259` and `L02260` behind a version check that skips them below `0x3b7`. `disp:4c` returns 414 hits over 219 owners with 1 in an orphan run and **none** owned by a record-pointer holder; `disp:4d` and `disp:4e` return **0 hits in the whole image**. Corpus, 12 085 records over both roots: `+0x28` is `0` on 12 085/12 085; `+0x2a` is nonzero on 348 per root with exactly four values (`1` ×117, `111` ×117, `64464` ×75, `64768` ×39) over seven maps (`scn:40` ×39, `scn:60` ×3, `scn:71` ×75, `scn:80` ×106, `scn:90` ×11, `scn:100` ×113, `scn:131` ×1), every one inside `scenario.res`, which is why the two roots agree exactly

**Confidence.** **High** for `rec+0x4e` having no direct-displacement reader anywhere — `disp:4e` is empty image-wide, so no base register can matter, and the only two indexed accesses to a record have bases `0x31` and `0x37` with maxima `0x36` and `0x3b` / Medium for `rec+0x4c`, whose 414 hits are excluded by `ALM-TAILHOLD-080`'s reduction rule rather than by an empty sweep, and for every corpus count

**Original status.** ● active

**Evidence.** [EXP-0296](../experiments/EXP-0296-unit-placement-tail/)

**Amended.** `ALM-UNITREAD-227` reads the stream read slot (`R0533`, `R0534`) that this card lists as unread: it keeps no record pointer. The grades keep the remaining caveat, array append `L02248` and the accessors `R1338` and `R2121`, which `ALM-UNITREAD-227` also names.

### ALM-SCALAR-087

**The type-5 `+0x08` word lands in `CPlayer+0x0c`; within the seven-function CPlayer family the field has one direct read, in copy constructor `R0503`, and that constructor has two direct call sites.** The loader's third 4-byte record read targets `obj+0x0c` at `L02261`/`L02262`; the absent-section arm writes 5000 to the same displacement at `L02263`. The family has direct field accesses at `L02264` (write zero), `L02265` (copy read), `L02266` (copy write), and `L02267` (write zero). The copy constructor is called at `L02268` and `L02269`, both inside `R0504`; `callto:R0504` has zero direct call or `.rdata` references, which does not exclude indirect reach. The CPlayer table has five function entries plus a null, and the two shared `CObject` stubs occupy slots `[2]` and `[3]` ALM-COPY-151 establishes distinct scalar storage in both owner arms when that owner is entered; it does not establish a native entry event.

**Confidence.** **High** for the destination, four family access sites and two direct call sites: the decoded operations are asserted in scratch and eleven independently mapped two-root ranges are bound by SHA-256 without publishing their bytes / **Unknown** for any consumer effect outside this family because other computed/advanced-pointer and whole-object mover paths remain open; the named owner copy and source-roster teardown are now bounded by ALM-COPY-151

**Original status.** ● active (amended)

**Evidence.** [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/), [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-SCALAR-088

**A provenance-selected direct-displacement population contains 17 map-side routines plus 7 CPlayer-family routines; no map-side hit was attributed to a read of the map-loaded `CPlayer+0x0c`, while the family contains `ALM-SCALAR-087`'s copy read. This is not a whole-image consumer exclusion.** The whole-image literal `+0x0c` census is 7013 hits / 2798 owners / 49 in orphan or undisassembled code. The 17 map-side routines are `R0526`, `R0414`, `R0128`, `R0423`, `R0523`, `R0504`, `R0524`, `R0478`, `R0522`, `R0277`, `R0118`, `R0116`, `R0278`, `R0489`, `R0151`, `R0067`, `R0461`; the seven family routines are `R0527`, `R0502`, `R0528`, `L02270`, `R0503`, `R0510`, `R0529`. The two known live loader callers destroy their map-owned lists before return at `L02271` and `L02272`, but a copy through an unsearched mover is not excluded. Computed access is used on this class's holders, and the earlier whole-image copy/assignment/serializer/`memcpy`-shaped search was not performed. ALM-COPY-151 closes the two selected copy-owner arms conditionally; ALM-ALIAS-154 distinguishes a retained map address from a live object. Neither extends this 17+7 negative to all 80 later surveyed functions

**Confidence.** **Medium** for the bounded direct-displacement absence: the population and attribution are explicit / **Unknown** for consumer effect and for native survival through the remaining computed/indirect escape paths. Supersedes the unscoped clauses in `ALM-PLAYER-069` and `ALM-GRP-041`

**Original status.** ● active (amended)

**Evidence.** [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/), [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-SCALAR-089

**The type-5 `+0x08` word takes exactly two values over both preserved roots, every record 0 authors `0`, and the loader's absent-section default authors `5000` into the record-0 position.** The complete walk covers 38 EN maps / 191 records and 34 RU maps / 166 records, with 0 tiling failures and `ru/Horror.alm` alone carrying no type-5 section. Across all 357 records, 97 carry 0 and 260 carry 5000; all 71 record-0 entries carry 0. Of the 286 later records, 260 carry 5000 and 26 carry 0. All 56 records with `+0x04 = 0` carry `+0x08 = 0`, while 41 records carry `+0x04 = 1` with `+0x08 = 0`. The absent-section default disagrees with every authored record-0 value; that fact does not establish inertness or any consumer effect. ALM-EDITOR-153 separately establishes an editor property source and a new-entry default, without converting the installed value set into a general allowed-value set

**Confidence.** **High** for counts and value set: a complete two-root walk enumerates every section and independently reproduces the 191 / 166 totals / **Unknown** for consumer effect; the corpus alone cannot distinguish no consumer from an unobserved one

**Original status.** ● active (amended)

**Evidence.** [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/), [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### ALM-CPLAYER-090

**The interface `CPlayer` vtable at `L02273` is six dwords: five function entries followed by a null at `L02274`; `L02275` begins the next class.** The five entries are `R0527`, `R0528`, the shared `CObject` stubs `R0530` and `R0531`, and `R0529`. The runtime-class record at `L02276` gives object size `0x48`. Three constructors stamp `L02275` at `L02277`, `L02278`, and `L02279`, and their routines begin at `L02280`, after the seven-function CPlayer family. A 16-dword dump crosses this boundary and attributes ten next-class entries to CPlayer; `PAL-SHADE-012` now uses the six-dword surface. This is the game table: the EN editor has a distinct four-function-plus-null table atL02281 (ALM-EDITOR-152), so matching scalar layouts do not transfer the game method surface to the editor

**Confidence.** **High**. The null boundary, next-vptr constructor sites, and five resolved entries agree in the compact static measurements; the corresponding two-root address ranges are published only as SHA-256 digests

**Original status.** ● active (amended)

**Evidence.** [EXP-0297](../experiments/EXP-0297-roster-scalar-reader/), [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md)

**Amended.** The table ledger carried the status "● active (amended)". `retracted.md` records a correction against this claim.

### ALM-META-091

**Case 0's sixteen reads, and the four map-object slots the stored type-0 scalars land in.** `R0478` case 0 is `[L02282, L02283)`, entered through the ten-entry table at `L02078`. It issues **sixteen** `Read(dest,len)` calls whose lengths tile the payload contiguously from `+0x00`, so file order fixes the mapping with no assumption about field sizes. Six land in `ebp`-relative locals of the loader itself (`+0x18→[ebp-0x5c]`, `+0x1c→[ebp-0x54]`, `+0x20→[ebp-0x10]`, `+0x24→[ebp-0x14]`, `+0x28→[ebp-0x50]`, `+0x2c→[ebp-0x18]`); ten land on `M`, the object the caller allocates with `new(0x32c)` and passes in `ecx`: `+0x00→M+0x00`, `+0x04→M+0x04`, **`+0x08→M+0x18`**, **`+0x0c→M+0x1c`**, **`+0x10→M+0x20`**, **`+0x14→M+0x24`**, `+0x30→M+0x94` (0x40 B), `+0x70→M+0xd4`, `+0x74→M+0xd8`, `+0x78→M+0xdc` (0x200 B). The four bold destinations are the ones `ALM-META-026` calls stored config scalars and `ALM-META-027` calls the stored light angle; neither row gives a destination. The loader pre-initialises three of them before the record loop — `M+0x18 = 0` (`L02284`), `M+0x24 = 0x20` (`L02285`), `M+0x20 = 0x20` (`L02286`)

**Confidence.** High (the rival is a mapping shifted by one read, which is what hand-pairing pushes with calls produces. Three anchors that fall out of this table without being used to build it exclude it: `ALM-META-026`'s published `+0x70/+0x74 → map+0xd4/+0xd8`, `ALM-META-058`'s single 0x200 read into `map+0xdc`, and the lengths ending at exactly `+0x278` — the same payload end the *other* type-0 reader reaches by a different route, 7×4 + `Seek(0x25c)` = 632 (`ALM-RDR-059`). A one-read shift breaks all three at once)

**Original status.** ● active

**Evidence.** [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

### ALM-META-092

**Over the searched population nothing reads the four stored scalars back off the map object, and the one reader with no object to store them in steps over them.** Interprocedural pointer reachability — places are the eight general-purpose registers plus frame- and stack-relative slots, moves, address formations and immediate adds carry a byte delta, a memory operand whose base holds the pointer is recorded as delta + displacement, a direct call seeds the callee's `this` and stack arguments, each body run to a CFG fixed point — from the three call sites of `R0478` (`L02287`, `L02288`, `L02289`), each seeded both at the caller's `new(0x32c)` result and at the loader's return value, plus two hand seeds for callees with no frame pointer: **100 entered bodies, 548 recorded uses, 7 accesses in `M+0x18..M+0x27`, of which 0 are reads.** Three are the pre-initialisations of `ALM-META-091`; four are the argument push that hands a slot's address to `Read`. The browser accept test `R0479` (`ALM-RDR-059`) does not read them either: its stream calls, enumerated over its whole body, open with two 4-byte reads, then `Seek` with `0x28` (`L02290`), then the 0x40-byte name — so it steps over payload `+0x08..+0x2f` entire, forty bytes and ten dwords, and no read in that body touches any of them

**Confidence.** **Medium** (a bounded negative, scoped to that population and to what the analyser can follow. Every place the pointer leaves the search is enumerated in the evidence and is the blind spot: 19 recorded stops at indirect calls reached with the pointer live, over 13 distinct sites, ten of which are the `Read` vtable slot itself; 7 callees entered without a frame pointer; one pointer escape at `L02291` (a store into the field at `+0x540d0`); five stops at unresolved jump tables. A consumer behind any of those would not appear. The builder-side copies of the same file bytes are a separate destination and a separate result, `TERR-LIGHT-149`)

**Original status.** ● active

**Evidence.** [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

**Amended.** EXP-0564 narrows the pointer escape at `L02291` to a bounded displacement-scan negative: the field at
world `+0x540d0` is stored and reloaded only in `R0277`: a `.text` displacement scan finds those two uses and an
immediate scan none (`ALM-META-311`, Medium), and `R0278`, which receives the map from the world constructor,
reads only `M+0x00`, `+0x04`, `+0x0c`, `+0x10` and `+0x14` (`ALM-MAPLIFE-226`). The other blind spots stand.

### ALM-CORP-093

**What the four stored type-0 scalars carry, over 72 shipped maps on both preserved roots** (loose `.alm` plus every `M7R` member of `scenario.res`; EN 38, RU 34; each file re-framed per `ALM-FRAME-031` rather than assumed). `+0x08`, the angle: ten distinct dwords, and every one is a **whole number of degrees** when its four bytes are read as an IEEE single — `±45` (41 maps and 2), `±44` (10 and 2), `±18` (4 and 6), `36`, `28`, `25`, `22` (one or two each); the modal `0x3f490fda` is one ULP below `float32(π/4)` = `0x3f490fdb`, and is the only one of the ten that is not the correctly-rounded `float32` of its whole degree, so "≈π/4 in most maps" and "45°" are the same observation. `+0x0c`: five values only — `360`×62, `480`×2, `645`×2, `720`×4, `1080`×2 — and `360` is `0x168`, the constant `R0532` adds to the shifted sub-tick clock at `L02292` before handing it to the sun model, whose unit is minutes (`SESS-TICK-026`, `TERR-LIGHT-014`); read that way the five are `06:00`, `08:00`, `10:45`, `12:00`, `18:00`. `+0x10`: sixteen values in `0..33`, mode `21`×20. `+0x14`: eight values in `27..64`, `64`×37 and `63`×22 — `64` is the unattenuated level of the terrain shading table (`TERR-LIGHT-020`) and is above anything the sun model writes to the global at `L02293`

**Confidence.** High (the census itself: every type-0 record on both roots, framing re-checked per file) / **Unknown** (what any of the four means. The minute-clock reading of `+0x0c` and the attenuation reading of `+0x14` are agreements between a shipped value set and a published constant, and no instruction reads either field — `ALM-META-092`, `TERR-LIGHT-149` — so nothing discriminates them from an editor-side annotation the engine ignores)

**Original status.** ● active

**Evidence.** [EXP-0298](../experiments/EXP-0298-map-scalar-consumers/)

### ALM-HEADER-097

The selected main reader and the editor reader compare file-header+0x10 as an unsigned dword: upper gate1001, and payload gates951,985,987,989,990. Main helperR0467 performs no record-header read at exactly1000 and leaves its destination untouched. Editor helperL02294 instead reads five4-byte words into destination offsets8,16,0,12,4 at 1000; otherwise it reads 20 bytes. BrowserR0479 and landscapeR0480 useR0465, which peeks 8 bytes, rewinds 8, then reads header-length bytes; their bodies do not test file+0x10. The main/editor reject branches and the landscape message-only branches are distinct. Original x86 boundary cases stop before main post-loop plane checks or at individual payload joins; no complete alternative-version map acceptance is claimed.

**Confidence.** High for instruction-derived read/gate distinctions; complete alternative-version acceptance Unknown

**Original status.** ● active

**Evidence.** [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/)

### ALM-HEADER-098

Record-header+0x10 is four raw bytes, not an established f32 or enum. Main stores them at frame EBP-0x24, browser at EBP-0x14, and the landscape record buffer at entryESP-0x10; the editor ordinary helper stores them at its caller frame EBP-0x24. No numeric use or equality enforcement was identified in the selected reader bodies.24 original-x86 record mutation cases vary one or all three synthetic record words over six structurally different patterns in main/landscape prefixes: no read of that four-byte slot, unchanged parse widths and scalar destination. This is a bounded negative; rendering, simulation, save/reload and later aliases were not observed.

**Confidence.** High for byte positions and measured prefix results; Medium for bounded absence; numeric/semantic role Unknown

**Original status.** ● active

**Evidence.** [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/)

### ALM-WRITER-099

Selected editor writerL02295 writes file+0x10 from literal990 atL02296, and computes header dataSize as 4*W*H+72. Its repeated record-header buffer is EBP-0x6c, length20, with the last four bytes at EBP-0x5c. The complete body has zero direct overlapping operands at that slot and ten address-taking sites for the same containing header; the byte writerL02297/L02298 receives those bytes with the other header words. There is no local initialization or input-record copy into that slot before the first header submission. Stack residue therefore fits the observed within-map equality without any numeric meaning. The actual incoming stack bytes and native editor output were not observed; neither bit preservation from an input map nor arbitrary editor versions is promised.

**Confidence.** High for selected writer stores, buffer and first-submission source; Medium for all-path residue interpretation; native round-trip Unknown

**Original status.** ● active

**Evidence.** [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/)

### ALM-STREAM-100

Main/browser construct wrapperL02299 with vtableL02300; landscape callerR0514 constructs the same wrapper beforeR0480. Its Read+3c isR0533: clamp requested count to unsigned length-cursor, seek underlying stream+10 to baseOffset+cursor, advance wrapper cursor by the clamped count before the underlying read, then forward the unchanged destination. The selected loose-file resolver path constructsL02301/vtableL02302 and copies its object pointer through the returned tuple into wrapper+10; underlying ReadR0534 forwards to imported ReadFile. The archive tuple instead takes container+34, whose constructor remains unexpanded. Synthetic four-byte reads yield full12345678, EOF-clamped aaaa5678 with cursor2, or successful-short aaaa5678 with wrapper cursor4 and return2. OS guarantees and archive receiver identity remain Unknown.

**Confidence.** High for named local loose-file chain and original-x86 transfer distinctions; broader receiver/lifetime coverage Unknown

**Original status.** ● active

**Evidence.** [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/)

### ALM-CENSUS-101

Read-only header inventory covers38 EN and 34 RU ALMs:16 loose files plus 56 scenario archive members,714 record headers. Every file word+0x10 is000003de. Every map has byte-identical record words internally, including the four-record RU Horror map. EN record-word frequencies are bfc02b6d:270, bff52d9d:60, bff62b6d:50; RU frequencies are 280,4,50 respectively. These are bit-pattern counts, not floating-point semantics or enforced equality. The full manifest records18 corpus install inputs and each map row has its own SHA-256.

**Confidence.** Medium: corpus-only equality and values

**Original status.** ● active

**Evidence.** [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/)

### ALM-EDSCALAR-102

The editor readerL02303 case0 requests four bytes for payload+0x0c into its object+0x1c, through the actual direct-file class/vtableL02304 Read+34=L02305 or the fallback wrapper class. Selected writerL02295 submits four bytes from its receiver+0x1c atL02306 after W,H and object+0x18. Writer callerL02307 obtains that receiver from document+0x50. ConstructorL02308 invokes readerL02303 on its own entry object through a verified jump thunk. These are separate local object bindings: document publication and intervening editor aliases/control changes remain unexpanded. No conversion, scalar meaning or native unchanged round-trip is established by sharing the offset.

**Confidence.** High for local reader/writer offsets, widths and receiver bindings; cross-event preservation and meaning Unknown

**Original status.** ● active

**Evidence.** [EXP-0346](../experiments/EXP-0346-alm-reader-transactions/)

### ALM-FLAGPATH-109

**Placement flags have arm-specific consumers, including a conditional bit-2 write.** Wire+0x0c reaches record+0x48 through ALM-TAILMAP-079. The complete spawner `R0151` reads it at `L02309/L02310/L02311/L02312/L02313`. Existing MISSION-ARM-006 resolution order stands. Within the Human band, bit 0 selects NPC. Ordinary Human type-key and definition-ID arms forward exactly `flags & 0x80` as constructor `R0497` third stack argument at `L02314/L02315`. NPC `L02316` passes literal zero there, alongside the separate NPC Hero mode; Units `L02317` calls `R0501` without this argument. The type-key Human arm always copies the secondary key low byte to actor+0x4b, then sets/clears its high bit from bit 2. The definition-ID arm performs those stores only when the secondary key low **word** is nonzero; secondary 0 preserves constructor state, while secondary 0x100 passes the word guard despite its zero low byte. NPC and Units do not perform those bit-2 stores. The remaining 29 flag bits have no consumer in the selected spawner; this establishes no universal purpose or lifetime irrelevance. Original-x86 one-bit probes use constructor cuts and corroborate argument/store scope; their patterned actor is not a native default.

**Confidence.** **High** for the five original reads, exact arguments and guarded stores; Unknown for unread-bit purposes and native downstream lifetime.

**Original status.** ✔ promoted

**Evidence.** [EXP-0347](../experiments/EXP-0347-unit-placement-flags/), byte-free instruction/decoder tables, call frontiers and synthetic vectors

### ALM-FLAGCORP-110

**The admitted 72 maps author flags 0, 1, 4 and 5 only.** The read-only version-990 census contains 38 EN maps / 8094 placements and 34 RU maps / 3991 placements. EN value counts are 0:7297, 1:14, 4:782, 5:1; RU counts are 0:3913, 1:14, 4:63, 5:1. Bit 0 is set on 15 records per root, bit 2 on 783 EN and 64 RU, and all other 30 bits, including bit 7, on zero records. Manifested inputs are scenario.res and loose ALM files; archive members are selected by the map signature. A map with no type-6 section contributes zero placements. This is an installed domain, not a valid-flags mask.

**Confidence.** **Medium**; exhaustive agreement over the manifested installed population at one version.

**Original status.** ✔ promoted

**Evidence.** [EXP-0347](../experiments/EXP-0347-unit-placement-flags/), byte-free instruction/decoder tables, call frontiers and synthetic vectors

### ALM-PLACESTREAM-111

**The case-6 interior destination reaches an underlying read; the previous stream escape is narrowed, not globally closed.** Loader `R0478` constructs `L02299`, installing vtable `L02300` whose +0x3c is `R0533`. That complete wrapper clamps length to remaining extent, seeks the underlying stream, advances its cursor and forwards the unchanged destination and clamped length at `L02318`. It stores no destination in its own fields. Opener `R0535` copies the source tuple into wrapper+0x10/+0x14/+0x18. Provider `R0536` tries archive tuples first; `R0537` supplies archive-object+0x34 as the underlying receiver, whose construction remains unclosed here. The separate loose-file arm constructs `L02301`, installs `L02302`, and selects read `R0534` at +0x3c. That body forwards the destination to imported ReadFile with null OVERLAPPED and a stack byte-count result. No durable destination store is present in these two read bodies. Imported service behavior, the archive receiver, exceptions and other aliases remain frontiers; ALM-TAILHOLD-080/ALM-TAILDIR-081 do not become universal no-reader claims.

**Confidence.** **High** for local forwarding and the loose-file vtable binding; Unknown for complete destination retention or all native record lifetimes.

**Original status.** ✔ promoted

**Evidence.** [EXP-0347](../experiments/EXP-0347-unit-placement-flags/), byte-free instruction/decoder tables, call frontiers and synthetic vectors

### ALM-PLACEEDITOR-112

**The EN editor reads, copies and offers the questioned Unit fields to its writer as opaque values.** The EN editor image is 1855488 bytes, SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`; no editor image exists in the admitted RU root. In loader `L02303` at version 990, wire+0x1c is a four-byte destination at object+0x66 (`L02319`). In writer `L02295`, `L02320` passes that same four-byte member to `L02297`. Insertion `L02321` clones through `L02322` and base `L02323`, which copies object+0x66 as a dword. Matching reader/writer members for the previously unresolved tail are wire+0x22 to object+0xfa, +0x26 to +0x10c, +0x28 to +0xf8, +0x2a to +0x10e (two bytes each), +0x30 to +0x112, +0x33 to +0x113, +0x34 to +0x114 (one byte each), and +0x35 to +0x118 (six-byte run). The two finite reader/writer intervals each issue 28 calls totaling 70 bytes. Seven original-x86 pairs varying only wire+0x1c across 0, 1, 2, 14, 15, 0x100 and 0xffffffff emit the same synthetic 70 bytes when run on the same supplied object with stream/sink cuts. These are operation pairs, not native editor round trips: insertion/copy execution, UI property producers and intervening editor events are unobserved. Field meaning and property authorship remain Unknown.

**Confidence.** **High** for the bound member/width operations and explicit copy stores; Medium for generalizing the selected operation pairs; native editor authorship and semantic purpose Unknown.

**Original status.** ✔ promoted

**Evidence.** [EXP-0347](../experiments/EXP-0347-unit-placement-flags/), byte-free instruction/decoder tables, call frontiers and synthetic vectors

## ALM claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-127 | The type4 placement word retains 16 raw bits, while the class path has three width boundaries. | High / Unknown | ✔ promoted | [EXP-0350](../experiments/EXP-0350-structure-placement-word/) |
| ALM-128 | Building's placement word is a zero override at this load stage. | High / Unknown | ✔ promoted | [EXP-0350](../experiments/EXP-0350-structure-placement-word/) |
| ALM-129 | The selected editor read/write path preserves the type4 word. | High / Unknown | ✔ promoted | [EXP-0350](../experiments/EXP-0350-structure-placement-word/) |
| ALM-130 | Confirming an existing editor structure property has a different word rule from direct writing. | High / Unknown | ✔ promoted | [EXP-0350](../experiments/EXP-0350-structure-placement-word/) |

### ALM-127

**The type4 placement word retains 16 raw bits, while the class path has three width boundaries.** On complete transfer,L02324 requests 2 bytes;L02325/L02326 store low16 at placement+06. Full wire kind32==33 selects the extension, low16 stored kind selects Shop only for 34/35, and the constructor receives low8. All other stored kinds take the Building constructor, not additional runtime subclasses. The walker registers the constructed object, frees the placement atL02327 and clears its slot atL02328. Nine kind controls include0x10022 reaching Shop and0x122 reaching Building; no invalid-definition acceptance is claimed.

**Confidence.** **High** for selected original instructions and 63 synthetic read/branch controls; complete transfer, valid definition admission and arbitrary alias lifetime remain bounded prerequisites/Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0350](../experiments/EXP-0350-structure-placement-word/)

### ALM-128

**Building's placement word is a zero override at this load stage.** SetupR0486 copies definition column3 into+44 and+42. After its constructor,L02329..L02330 writes+42 only when placement+06 is zero; every nonzero word leaves the initialized value. The constructor does not receive the word. Shop takes the separate existing signed i16*1000 cap chain and does not execute this zero override. The selected Building virtual+14 isR0538, whose established selector15/16 and counter/range predicates can increment current0; no irreversible-zero or universal field meaning follows.

**Confidence.** **High** for the named local relation, separately read setup, resolved virtual target and 63 original-instruction controls with explicit synthetic constructor defaults; native lifecycle and all-definition semantics **Unknown**

**Original status.** ✔ promoted

**Evidence.** [EXP-0350](../experiments/EXP-0350-structure-placement-word/)

### ALM-129

**The selected editor read/write path preserves the type4 word.** ReaderL02331 requests 2 bytes and forwards a sign-extended word to ordinary factoryL02332 or kind33 factoryL02333. Both reached factories callL02334, which stores low16 at CStructure+f6/+fa. WriterL02295 admits CStructure-derived objects andL02335..L02336 emits 2 bytes directly from+f6. No scaling or zero normalization occurs at that site. Twenty-one synthetic original-instruction read/initializer/write cases retain all seven tested word patterns for ordinary, extension and Shop representatives.

**Confidence.** **High** for the selected instruction relation conditional on full transfer and factory admission; native round-trip acceptance, backing I/O, intervening callbacks and other producers **Unknown**

**Original status.** ✔ promoted

**Evidence.** [EXP-0350](../experiments/EXP-0350-structure-placement-word/)

### ALM-130

**Confirming an existing editor structure property has a different word rule from direct writing.** The selected mode0x8043/CStructure handlerL02337 uses dialog226, integer control1041 and sign-extended+f6. After modal result1 and nonzero class selection,L02338 applies signed upper clamp3000, thenL02339/L02340 store low16 in+f6/+fa. Eight post-confirmation instruction cases retain0,1,2999,3000,-1,-32768 but replace3001 and 32767 by 3000. A direct write retains0x7fff; this property path can change it to0x0bb8.

**Confidence.** **High** for the local guarded store, constructor/control relation and eight instruction controls; UI validation, natural command delivery and accepted negative input **Unknown**

**Original status.** ✔ promoted

**Evidence.** [EXP-0350](../experiments/EXP-0350-structure-placement-word/)

## ALM claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-TILEMAIN-121 | The main type-1 arm at L02283 requests 2 * M.width * M.height bytes, stores the allocated pointer at M+0x0c, and submits that destination and count to stream virtual +0x3c at L02341. | High | ✔ promoted | [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |

### ALM-TILEMAIN-121

The main type-1 arm at L02283 requests 2 * M.width * M.height bytes, stores the allocated pointer at M+0x0c, and submits that destination and count to stream virtual +0x3c at L02341. M.width/height are at +0/+4. No tile-bit operation occurs in this arm. In 136 supplied-memory cases with a complete synthetic read, all 16 word bits survive into M+0x0c, and the simulation ingest leaves that source word unchanged. This is conditional transfer, not evidence that an incomplete read, every reader, or a native load succeeds.

**Confidence.** High for the named original caller and conditional transfer; lower read and native transaction excluded

**Original status.** ✔ promoted

**Evidence.** [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

## ALM claims (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-TILEVIEW-122 | The light type-1 arm L02342 uses P.width/height at +4/+8, stores the allocated pointer at P+0x0c and requests 2WH bytes through virtual +0x3c. | High / Unknown | ✔ promoted | [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv` |

### ALM-TILEVIEW-122

The light type-1 arm L02342 uses P.width/height at +4/+8, stores the allocated pointer at P+0x0c and requests 2WH bytes through virtual +0x3c. Its end-of-parse loop L02343 applies AND word, 0xdfff at L02344 over that plane. All 136 word cases retain bits 0..12/14/15 and clear bit 13 after a complete synthetic read. Binder R0515 stores P itself at view+0x80 and dimensions at view+0x84/+0x88; it does not copy the tile plane. Renderer R0539 reads [[view+0x80]+0x0c]. This light-parser layout is distinct from M in ALM-TILEMAIN-121; equal +0x0c offsets do not establish shared allocation with the simulation input.

**Confidence.** High for original ranges, pointer binding and all 16 one-bit controls; later ownership and native ordering Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0349](../experiments/EXP-0349-tile-word-consumers/), `evidence/instructions.tsv`, `evidence/controls.tsv`, `evidence/boundaries.tsv`

## Authored Player control

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-139 | ALM type5 wire+04 requests four bytes into CPlayer+30 atL02345. | High / Unknown | ✔ promoted | [EXP-0351](../experiments/EXP-0351-authored-player-control/) |
| ALM-140 | The selected Player-control catalog contains 52 sites, including 27 actor-owner reads: the 25 sites retained by UNIT-OWNER-009, XOR-separatedL02348 andL02349 before the shared call. | High / Unknown | ✔ promoted | [EXP-0351](../experiments/EXP-0351-authored-player-control/) |
| ALM-COPY-151 | The ROM1 map-copy owner creates independent scalar cells in both CPlayer-copy arms. | High / Unknown | ✔ promoted | [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md); generated checks, call/reference tables and original-x86 vectors |
| ALM-EDITOR-152 | The EN editor copies the roster scalar and writes its current four-byte value. | High / Medium / Unknown | ✔ promoted | [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md); generated checks, call/reference tables and original-x86 vectors |
| ALM-EDITOR-153 | The EN roster dialog has a scalar producer and a map-copy restoration path. | High / Medium / Unknown | ✔ promoted | [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md); generated checks, call/reference tables and original-x86 vectors |
| ALM-ALIAS-154 | A retained map pointer is not evidence that its CPlayer objects remain alive. | High / Unknown | ✔ promoted | [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md); generated checks, call/reference tables and original-x86 vectors |
| ALM-METACOPY-183 | Game map-copy owner R0504 reads source M+0x1c at L02373 and stores destination+0x1c at L02374. | High / Unknown | ✔ promoted | [EXP-0357](../experiments/EXP-0357-metadata-scalar-lifetimes/EXP-0357.md); generated checks, reference counts and vectors |
| ALM-METALIFE-184 | Owning constructor R0277 stores M at world+0x540d0, calls grid ingest R0278, then destroys and frees M before return, leaving the address in the holder. | High / Medium / Unknown | ✔ promoted | [EXP-0357](../experiments/EXP-0357-metadata-scalar-lifetimes/EXP-0357.md); generated M-access/frontier tables and checks |
| ALM-METAEDITOR-185 | On its base-open success arm, editor load handler L02376 constructs a 0x32c-byte E through L02308/L02303, installs its return at document+0x50 at L02377, and, when E+0x0c is nonzero, constructs a backup through ... | High / Unknown | ✔ promoted | [EXP-0357](../experiments/EXP-0357-metadata-scalar-lifetimes/EXP-0357.md); generated checks and vectors |
| ALM-METADEFAULT-186 | The editor new-map wrapper L02380 calls base constructor L02381, whose L02382 writes 0x168 to E+0x1c. | High / Unknown | ✔ promoted (amended, partially retracted) | [EXP-0357](../experiments/EXP-0357-metadata-scalar-lifetimes/EXP-0357.md); generated command window, checks and vectors |
| ALM-COUNT-195 | For the primary ROM1 reader, metadata payload+0x28 is read atL02385 into EBP-0x50, then overwritten by each complete four-byte type7 internal count read atL02386/L02387/L02388. | High / Unknown | ● active | [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/) |
| ALM-STALE-196 | At the primary scalar/header/count seams, callers do not check the returned byte count. | High / Unknown | ● active | [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/) |
| ALM-HDRREAD-197 | Primary record helperR0467 ordinarily requests 20 bytes; record+0x10 occupies EBP-0x24 in the measured caller. | High / Unknown | ● active | [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/) |
| ALM-SCALAR-198 | Primary metadata+0x0c Read4 atL02392 reaches M+0x1c; complete light readerR0480 Read4 atL02393 reaches distinct P+0x2c. | High / Unknown | ● active | [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/) |
| ALM-READERS-199 | Primary header gate controls95 combinations: unsigned recordCount>=3, then unsigned version<=1001. | High / Unknown | ● active | [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/) |
| ALM-FRONTIER-200 | The identified normal ALM header/metadata frontier is discharged with explicit semantic limits. | High / Medium / Unknown | ✔ promoted | [EXP-0358 semantic result](../experiments/EXP-0358-alm-headers-metadata/SEMANTIC-RESULTS.md); closure.tsv; semantic-header, semantic-m, semantic-m-cuts, P and semantic-v products |

### ALM-139

ALM type5 wire+04 requests four bytes into CPlayer+30 atL02345. The reader does not test the returned count before the next field. Both non-copy CPlayer constructors write dword zero; copyR0503 reads/writes the complete dword atL02346/L02347. The actual wrapper vtableL02300 binds Read+3c toR0533, which forwards the destination after clamping length; ALM-STREAM-100 retains the short-read/native-I/O boundary. Sixteen arbitrary-word controls reproduce the supplied full read and complete copy; constructor controls retain zero. These are local representation rules, not a Boolean validation or a native arbitrary-map acceptance result.

**Confidence.** High for named instructions and conditional local transfers; native I/O and the editor authoring Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0351](../experiments/EXP-0351-authored-player-control/)

### ALM-140

The selected Player-control catalog contains 52 sites, including 27 actor-owner reads: the 25 sites retained by UNIT-OWNER-009, XOR-separatedL02348 andL02349 before the shared call. Twenty-six actor sites feed whole-dword zero/equality clauses;L02349 feeds an unconditionalR0008 call whoseL02350 overwrites flags andL01148 overwrites EAX before their use. This read is not another control branch. SiteL02351 tests the full dword and returns AL=0 on nonzero, preserving EAX high bytes; the named callerL02352 stores only AL andL02353 tests that byte. Other clauses reach route-budget eligibility, experience suppression, placement/cull gates, mover state, thresholds, group bookkeeping and relation notification. The table names each local condition and next operation. Ten representative predicate slices use16 word patterns each, including zero-low-byte nonzero values. No complete AI tick or universal consumer enumeration is claimed.

**Confidence.** High for catalogued local instructions and conditional branch controls; completeness and native joined effects Medium/Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0351](../experiments/EXP-0351-authored-player-control/)

### ALM-COPY-151

**The ROM1 map-copy owner creates independent scalar cells in both CPlayer-copy arms.** At L02268 it copies source-map+0x3c into a fresh 0x48-byte object, installs the result at destination-map+0x3c and appends it through destination-map+0x28. At L02269 it copies later source roster pointers and appends each result. Constructor R0503 explicitly copies source+0x0c to destination+0x0c at L02265/L02266. The selected pointer-array append copies a pointer value, not the 0x48-byte payload. Sixteen original-x86 cases cover eight scalar patterns and two destination fills: two distinct destination objects retain both scalar values after the original source roster destructor dispatches vptr+0x4 to R0528 and a named free cut poisons the source objects. Relation-array copying is an explicit cut at advanced receivers+0x34. Native entry into R0504 is unobserved: every file byte was searched for its absolute address and every .text byte for E8/E9 relative encodings, with zero hits. Computed dispatch remains open.

**Confidence.** **High for the named original instructions and conditional scalar independence; Unknown for native owner reach, later consumers and universal no-survival. The synthetic destruction witness does not establish a native lifetime.**

**Original status.** ✔ promoted

**Evidence.** [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md); generated checks, call/reference tables and original-x86 vectors

### ALM-EDITOR-152

**The EN editor copies the roster scalar and writes its current four-byte value.** Its CPlayer table L02281 has four function entries followed by null; the game table is different. Copy constructor L02354 reads+0x0c at L02355 and writes the destination at L02356. Editor map-copy L02357 allocates distinct 0x48-byte CPlayers at its two calls through L02358, stores the first at map+0x3c and uses map+0x28 as the roster. Writer L02295 computes the selected pointer from map+0x2c, advances it by0c and calls L02297 at L02359 with size 4 and count 1. Sixteen original-x86 cases copy the scalar, poison the source, and execute the original buffered writer path: eight bit patterns under two destination fills emit the unchanged word and advance its synthetic output buffer by 4. Locks and the distinct embedded relation-array copy are named cuts; no file flush or original editor process runs.

**Confidence.** **High for explicit scalar-copy, roster-holder and buffered-write operations; Medium for applying the bounded pair to an intervening editor transaction; native round-trip and gameplay meaning Unknown.**

**Original status.** ✔ promoted

**Evidence.** [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md); generated checks, call/reference tables and original-x86 vectors

### ALM-EDITOR-153

**The EN roster dialog has a scalar producer and a map-copy restoration path.** Open handler L02360 binds dialog+0x198 to document-map+0x28, snapshots the map through L02361 -> L02357, and supplies dialog 225. Data exchange binds integer control 0x407 to dialog+0x18c and asks signed validation 0..0x7fffffff. Its original indirect call at L02362 uses dialog-vtable L02363+0x84 -> L02364 -> L02365. Initialization copies selected CPlayer+0x0c to dialog+0x18c. Selection-change, add and accept handlers write dialog+0x18c back to selected CPlayer+0x0c; add then initializes the new entry scalar to 5000. The add message row is message 0x111/notification 0/control 0x40e -> L02366 -> L02367. After modal return!=1, L02360 destroys the edited map, clones the saved map, installs its pointer at document+0x50 and calls the unresolved document+0x5c refresh receiver at L02368. Sixteen original-x86 cases establish initialization/store/default slices and the signed validator failure edge; no text parsing, modal event, refresh callback or native accept/cancel is executed.

**Confidence.** **High for bound instructions, table/roster receivers and selected scalar transfers; Medium for the inferred native accept/cancel preservation policy; native UI validation and emitted-file effects remain Unknown.**

**Original status.** ✔ promoted

**Evidence.** [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md); generated checks, call/reference tables and original-x86 vectors

### ALM-ALIAS-154

**A retained map pointer is not evidence that its CPlayer objects remain alive.** In game constructor R0277,L02291 stores the loaded map pointer at world+0x540d0,L02369 reloads it,L02370 calls map destructor R0522 and L02371 frees the map. This body does not clear the stored address before returning. The roster destructor follows map+0x28 and dispatches each CPlayer vptr+0x4 at L02372; ALM-COPY-151 separately demonstrates independent copies. The 80-function survey and call table do not resolve every computed world/map alias or every indirect receiver. The selected ordinary mission and restore loader callers still destroy their original map-owned objects before returning; no all-image no-copy or no-consumer conclusion follows.

**Confidence.** **High for the local store/destruction sequence; Unknown for post-return computed alias use and for any resulting gameplay effect.**

**Original status.** ✔ promoted

**Evidence.** [EXP-0352](../experiments/EXP-0352-cplayer-scalar-lifetime/EXP-0352.md); generated checks, call/reference tables and original-x86 vectors

### ALM-METACOPY-183

Game map-copy owner R0504 reads source M+0x1c at L02373 and stores destination+0x1c at L02374. Twelve original-x86 cases vary six bit patterns and two destination fills; the copied word survives poisoning the source map. This is distinct from the CPlayer copies inside the same owner. The owner's address has zero absolute matches over 1,977,344 file bytes and zero possible E8/E9 references over 1,660,928 text bytes in the admitted ROM1 image. Those representation bounds do not resolve computed/native entry.

**Confidence.** High for local four-byte independent copy; native owner reach and gameplay role Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0357](../experiments/EXP-0357-metadata-scalar-lifetimes/EXP-0357.md); generated checks, reference counts and vectors

### ALM-METALIFE-184

Owning constructor R0277 stores M at world+0x540d0, calls grid ingest R0278, then destroys and frees M before return, leaving the address in the holder. Borrowing constructor R0118 passes its supplied M to the same ingest without this publication/destruction sequence. The ingest REP MOVSD moves world byte planes, not the M header; world-relative callees R0469 and R0470 were unexpanded in this original population; their local frontier is subsequently narrowed by ALM-INGEST-201, without closing later aliases or native first use. The four named mission pre-teardown calls enter R0489/R0151/R0067/R0461: 32 M-bearing call sites address containers at M+0x304/+0x318/+0x40/+0x5c/+0x78/+0x2dc/+0x2f0. Expansion through 17 accessor bodies records 29 M-relative uses and no access overlapping M+0x1c..1f; returned pointed storage and two unresolved switch tables remain frontiers. The selected L02375 call receives container+0x14, giving M+0x54/M+0x70; no-frame R0540 allocates 4+count*stride, links the previous block head, and replaces that head. It does not copy a whole M or receive M+0x1c on these calls.

**Confidence.** High for named instructions and receiver distinctions; Medium for bounded negative; post-return aliases and native first use Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0357](../experiments/EXP-0357-metadata-scalar-lifetimes/EXP-0357.md); generated M-access/frontier tables and checks

### ALM-METAEDITOR-185

On its base-open success arm, editor load handler L02376 constructs a 0x32c-byte E through L02308/L02303, installs its return at document+0x50 at L02377, and, when E+0x0c is nonzero, constructs a backup through L02361/L02357 and publishes it at document+0x54. The base copy explicitly reads E+0x1c at L02378 and stores destination+0x1c at L02379. Writer caller L02307 selects document+0x50;L02306 submits one four-byte item from receiver+0x1c. Twelve original-x86 slice cases execute publication, scalar copy and buffered output over six patterns/two fills; output equals the word after the source is poisoned and the buffer advances four bytes. Constructor/copy slices and lock cuts are not a native editor transaction.

**Confidence.** High for local publication/copy/writer bindings and conditional storage; intervening replacement/control policy and native emitted-file result Unknown

**Original status.** ✔ promoted

**Evidence.** [EXP-0357](../experiments/EXP-0357-metadata-scalar-lifetimes/EXP-0357.md); generated checks and vectors

### ALM-METADEFAULT-186

The editor new-map wrapper L02380 calls base constructor L02381, whose L02382 writes 0x168 to E+0x1c. The existing-file loader uses a separate constructor path. Corrected controls reinitialize the separate creation slice and record its actual E receiver, incoming word and remaining-object fill at L02382: six declared words crossed with fills 5a/a5 give 12 distinct observed pairs and produce 0x168. The first candidate's claimed varied default-input population is withdrawn: all 12 of those slices entered with 0xdddddddd after source poisoning; see [retracted.md](retracted.md). This is not an input validation range or evidence for a time unit. In the selected 16 command-table rows L02383..L02384 and their outer handler bodies, no exact E+0x1c control binding was established. Indirect dialogs, redraw helpers and other commands were not exhausted.

**Confidence.** High for the separate default producer; scalar exposure/recomputation policy and native UI outcomes Unknown

**Original status.** ✔ promoted (amended control population)

**Evidence.** [EXP-0357](../experiments/EXP-0357-metadata-scalar-lifetimes/EXP-0357.md); generated command window, checks and vectors

**Amended.** The table ledger carried the status "✔ promoted (amended control population)". `retracted.md` records a correction against this claim. The qualifier follows that record.

### ALM-COUNT-195

For the primary ROM1 reader, metadata payload+0x28 is read atL02385 into EBP-0x50, then overwritten by each complete four-byte type7 internal count read atL02386/L02387/L02388. Their original signed loop comparisons atL02389/L02390/L02391 and original container operations yield the internally declared A/C/T cardinalities in M+0x40/+0x5c/+0x78. 768 original-CPU cases cross six independent metadata patterns, all64 count triples0..3 and two fills; read visits/cardinalities and type7 consumption12+796*A+796*C+184*T follow internal counts, never the previous metadata word after full reads. This independently remeasures ALM-META-025 and refutes the contradictory type7 clauses of ALM-REQ-056/ALM-ORD-057. Field purpose outside these loops remains Unknown.

**Confidence.** High for complete-read loop source, original storage/cardinality and measured population; other field purpose Unknown

**Original status.** ● active

**Evidence.** [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/)

### ALM-STALE-196

At the primary scalar/header/count seams, callers do not check the returned byte count. Original wrapperR0533 clamps to remaining length and advances logical cursor before original CFileR0534 calls the controlled ReadFile import; successful-short returns do not roll it back. EOF before any type7 count can preserve metadata+0x28 or the prior internal count in EBP-0x50:18 EOF controls build the stale1/2/3 population despite zero-byte count/payload transfers. Six zero-byte-short controls carry wire count0 but preserve old2 and build2/2/2 with adequate synthetic payload. Sixty short-metadata cases still yield1/2/3 after complete internal reads. This is residual storage under the explicit I/O contract, not a format count rule, valid-map acceptance, or a native OS guarantee.

**Confidence.** High for named original instructions, observed residual sources and cursor contract; native OS behavior and malformed-map continuation Unknown

**Original status.** ● active

**Evidence.** [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/)

### ALM-HDRREAD-197

Primary record helperR0467 ordinarily requests 20 bytes; record+0x10 occupies EBP-0x24 in the measured caller. At file version1000 the helper reads nothing and retains the old destination.66 length/version/fill controls and 10 sequential-header controls distinguish complete transfer, partial suffix from the previous header and total no-read preservation. Browser/light helperR0465 instead peeks 8, rewinds 8 and requests hdrLen:36 bounded16/20/24-length/truncation controls confirm this, without a version test.72 independent primary header/scalar controls impose no within-map equality. The field remains byte[4]/opaque: transport does not establish f32, signedness, enum, pointer or any gameplay consumer. Global alias/lifetime completeness is not claimed.

**Confidence.** High for measured transfer/skip/reuse mechanics; numeric representation and semantic role Unknown

**Original status.** ● active

**Evidence.** [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/)

### ALM-SCALAR-198

Primary metadata+0x0c Read4 atL02392 reaches M+0x1c; complete light readerR0480 Read4 atL02393 reaches distinct P+0x2c. The measured browser window reads W/H and seeks40 bytes from payload+8 to+0x30, skipping both disputed metadata words; light reads the first28 bytes then skips604.240 metadata EOF/short cases separate untouched destination suffixes from wire bytes. Twelve original scalar-copy slicesL02394..L02395 log actual source word and independent destination fill at entry, execute the original load/storeL02373/L02374, then preserve the copy after source poisoning. These extend transport observations, not the complete copy owner's native caller, a time unit, editor producer completeness or gameplay effect.

**Confidence.** High for named read/copy/skip mechanics; complete lifetime and gameplay meaning Unknown

**Original status.** ● active

**Evidence.** [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/)

### ALM-READERS-199

Primary header gate controls95 combinations: unsigned recordCount>=3, then unsigned version<=1001. Browser header controls95 combinations: unsigned recordCount>=2 without a version rejection in the predicate.144 complete light-reader cases with W=H=0, type0 and skipped type7 reach EOF even at version1002/0x80000000/0xffffffff; count1 logs its advisory message but continues.70 primary payload windows confirm unsigned thresholds951(two u16),985(one byte),987(one dword),989(type8 dword; explicit0 below),990(type9 dword). Window high-bit controls bypass the earlier whole-file upper gate and do not prove accepted maps. The editor upper gate from ALM-HEADER-097 remains inherited; its complete header helper now has separate fresh ALM-EDITORHDR-202 controls. Browser later listing/CString/UI processing is outside the dynamic population.

**Confidence.** High for named predicates/windows and bounded light-reader execution; complete primary/browser/editor native acceptance Unknown

**Original status.** ● active

**Evidence.** [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/)

### ALM-FRONTIER-200

**The identified normal ALM header/metadata frontier is discharged with explicit semantic limits.** ALM-HDRLIFE-206 closes the named header-local and incidental-register lifetimes; ALM-EDITTIME-205 establishes the editor Starting Time producer and color effect. Mission/save-resume M is destroyed after ingest and five map consumers. Both placement switches have eight admitted entries. The borrowing world constructor returns world; the CString edge retains only independent name bytes from M+0x94. The owning world+0x540d0 holder has only its publication/reload literal references and outlives the pointed-to allocation, not its ownership. The actual whole-P render argumentR0541 reads dimensions/planes and creates no scalar escape. The thirty-slot V class and identified app+0xd0, drawable+0xe0/copy, UI+0x80, child+0x30 and focus+0x34/+0x38 aliases are followed through 63 additional P getters in 32 methods. Four local P spills and two syntactic return candidates have checked dispositions; no live scalar-specific edge remains in this declared graph. The selected M/P paths do not read the queried scalar and end their pointer ownership without a new scalar copy. Complete type7 reads still use internal counts; the editor produces metadata+0x28 as the three cardinalities' sum. The 2858 prior case definitions remain separate recorded original-function/emulated evidence.

**Confidence.** **Medium** for normal-path negative semantic closure over explicit receivers, call conventions and in-bounds objects. **High** positive/editor and local mechanical facts retain their own claims. Universal computed-alias absence, native receiver/event ordering, complete editor transactions, malformed buffers and externally installed callbacks remain **Unknown**. The new static continuation is not covered by the sole earlier confidence review.

**Original status.** ✔ promoted

**Evidence.** [EXP-0358 semantic result](../experiments/EXP-0358-alm-headers-metadata/SEMANTIC-RESULTS.md); closure.tsv; semantic-header, semantic-m, semantic-m-cuts, P and semantic-v products

## Authored Player control (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-INGEST-201 | Complete originalR0278 and calleesR0469/R0470 execute in 304 additional controls:288 cases cross six M+0x1c words, two fills, four shapes, two synthetic table variants and world+0x540d0 as M/null/unmapped0xdeadc0de;16 controls ... | High / Medium / Unknown | ● active | [EXP-0358 continuation](../experiments/EXP-0358-alm-headers-metadata/ALIAS-RESULTS.md); reproducible304-case archive and static inventory |
| ALM-EDITORHDR-202 | The original editor header helper has a distinct version1000 transfer program. | High / Unknown | ✔ promoted | [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/EDITOR-RESULTS.md) |
| ALM-EDITORSCALAR-203 | Fresh editor scalar controls distinguish a direct read, independent copy, buffered writer and separate creation default. | High / Unknown | ✔ promoted | [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/EDITOR-RESULTS.md) |
| ALM-EDITCOUNT-204 | The editor metadata+0x28 producer adds its three type7 cardinalities before buffered output. | High / Unknown | ✔ promoted | [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/EDITOR-RESULTS.md) |
| ALM-TRIGSTORAGE-163 | Complete game and the editor Trigger reads transfer184 bytes, including the opaque span+0x40..0x7f. | High / Unknown | ✔ promoted | [EXP-0353](../experiments/EXP-0353-trigger-storage/) |
| ALM-TRIGEDITOR-164 | The selected editor dialog and writer copy/submit the whole184-byte Trigger. | High / Unknown | ✔ promoted | [EXP-0353](../experiments/EXP-0353-trigger-storage/) |
| ALM-TRIGZERO-165 | The game Trigger initializer zero-fills count*184 bytes before per-record debug checks. | High / Unknown | ✔ promoted | [EXP-0353](../experiments/EXP-0353-trigger-storage/) |
| ALM-TRIGMETA-166 | The selected editor writer produces metadata+0x28 as the modulo2^32 sum of its three type7 cardinalities. | High | ✔ promoted | [EXP-0353](../experiments/EXP-0353-trigger-storage/) |
| ALM-RECVSOURCE-207 | The selected frontend receive path reads from entries requested through manager+0x18b8. | High / Unknown | ✔ promoted | [EXP-0360](../experiments/EXP-0360-alm-receive-source/) |
| ALM-RECVREAD-208 | The source reader consumes a linked sequence of buffer payloads, with separate cursor and removal operations. | High / Unknown | ✔ promoted | [EXP-0360](../experiments/EXP-0360-alm-receive-source/) |
| ALM-RECVOBJ-209 | Received type6 selects a shared object and an in-place variable-length reader, separate from the type4 request object. | High / Unknown | ✔ promoted | [EXP-0360](../experiments/EXP-0360-alm-receive-source/) |
| ALM-SUBMIT-210 | The selected type4 request serializes nine bytes into an indexed send-side region;L02443 is a counter update request. | High / Unknown | ✔ promoted | [EXP-0360](../experiments/EXP-0360-alm-receive-source/) |
| ALM-EDITTIME-205 | The EN editor interprets metadata+0x0c as Starting Time in minutes and uses it in its color pipeline. | High / Unknown | ✔ promoted | [EXP-0358 semantic result](../experiments/EXP-0358-alm-headers-metadata/SEMANTIC-RESULTS.md); editor-semantic products |

### ALM-INGEST-201

Complete originalR0278 and calleesR0469/R0470 execute in 304 additional controls:288 cases cross six M+0x1c words, two fills, four shapes, two synthetic table variants and world+0x540d0 as M/null/unmapped0xdeadc0de;16 controls omit five register/EFLAGS-preserving observation bridges. All return1, preserve scalar/alias and agree on whole-world output within the same fill/shape/table group after excluding only the varied alias word. All five original classifier targets are visited; loss controls preserve output. The three complete bodies contain373 instructions, two direct calls and one five-target computed jump. Callees receive world, not M; the final original REP copies world planes. Two complete runs and their raw observations agree bytewise. This narrows the named ingest frontier of ALM-METALIFE-184; it does not prove no discarded reads, no other alias consumers, native reachability, a semantic unit or the complete wire-to-gameplay effect chain.

**Confidence.** High for named local instruction/receiver/copy facts; Medium for bounded no-output-effect population; overall field meaning Unknown

**Original status.** ● active

**Evidence.** [EXP-0358 continuation](../experiments/EXP-0358-alm-headers-metadata/ALIAS-RESULTS.md); reproducible304-case archive and static inventory

### ALM-EDITORHDR-202

**The original editor header helper has a distinct version1000 transfer program.** CompleteL02294 executes five Read4 calls to destination offsets8,16,0,12,4 at version1000 and one Read20 for the other six tested versions. In 588 crossed controls it returns after supplied short transfers without checking returned counts; the unread destination suffix retains its prior fill. At 1000 the last complete wire word lands at header+4.

**Confidence.** **High** for the complete helper and measured versions, words, fills and per-read caps. The synthetic Read callback is a named lower-service contract, not Windows I/O; whole-file admission, numeric meaning and native effects remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/EDITOR-RESULTS.md)

### ALM-EDITORSCALAR-203

**Fresh editor scalar controls distinguish a direct read, independent copy, buffered writer and separate creation default.** SixtyL02396..L02397 cases transfer0..4 bytes to E+0x1c with explicit old-byte provenance. Twelve cases executeL02378/L02379, poison source storage and retain the independent copy through original buffered writerL02306. Twelve separate entries restore all six actual prior words and both fills immediately beforeL02382; every store writes0x168.

**Confidence.** **High** for these original local operations, backed by actual-entry telemetry and twice-reproduced output. Writer lock/unlock are named cuts with a synthetic buffer. These controls alone do not identify a unit. ALM-EDITTIME-205 separately establishes the original editor control and minutes-to-color chain; native transactions and ROM1 gameplay effects remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/EDITOR-RESULTS.md)

### ALM-EDITCOUNT-204

**The editor metadata+0x28 producer adds its three type7 cardinalities before buffered output.** OriginalL02398/L02399/L02400 read E+0x4c/E+0x80/E+0x68;L02401/L02402 add them, andL02403 invokes the original buffered writer. All64 triples0..3 crossed with two fills write their sum and advance output by 4 bytes. This is a producer result, separate from the game's internal-count loop-bound rule.

**Confidence.** **High** for the 128 bounded executions and named original arithmetic/output. Native SAVE, invalid cardinality admission and a gameplay use of this metadata word remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0358](../experiments/EXP-0358-alm-headers-metadata/EDITOR-RESULTS.md)

### ALM-TRIGSTORAGE-163

**Complete game and the editor Trigger reads transfer184 bytes, including the opaque span+0x40..0x7f.** GameL02404/L02405 submits0xb8 through the constructed stream Read slot after selecting the array record; editorL02406/L02407 submits the same span at E+0x7c+index*0xb8. The legacy 0-of-421 printable-string corpus result does not prove pointers or a universal non-text type.

**Confidence.** **High** for byte-verified widths and carrier addresses in the selected bodies. Complete transfer is conditional on the lower read; native delivery and the opaque span's typed role remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0353](../experiments/EXP-0353-trigger-storage/)

### ALM-TRIGEDITOR-164

**The selected editor dialog and writer copy/submit the whole184-byte Trigger.** Existing-record initializationL02408/L02409 copies0x2e dwords to dialog+0x338, so the queried span is dialog+0x378..0x3b7. ApplyL02410/L02411 copies0x2e dwords back. WriterL02412/L02413 submits one0xb8-byte item from E+0x7c+index*0xb8.

**Confidence.** **High** for these original copy spans and submission arguments. UI/name-copy tailL02414->L02415, new-record lifetime, actual input policy and a native durable LOAD/edit/SAVE transaction remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0353](../experiments/EXP-0353-trigger-storage/)

### ALM-TRIGZERO-165

**The game Trigger initializer zero-fills count*184 bytes before per-record debug checks.** OriginalL02416 computes count*0xb8 andL02417 callsL02418 with zero and that span. The editor no-buffer arm zero-fills newCount*0xb8 atL02419. On growth,L02420 first preserves oldCount records, thenL02421 zero-fills only the newCount-oldCount tail; spare capacity is outside that span. These explicit initialization paths include the second64 bytes.

**Confidence.** **High** for the selected zero-fill operations. They do not identify the origin of bytes read from an editor file or establish native dialog creation; the complete producer lifetime remains **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0353](../experiments/EXP-0353-trigger-storage/)

### ALM-TRIGMETA-166

**The selected editor writer produces metadata+0x28 as the modulo2^32 sum of its three type7 cardinalities.** InL02295, loadsL02398/L02399/L02400 select E+0x4c/E+0x80/E+0x68; addsL02401/L02402 combine them andL02403 submits four bytes after ten prior scalar words. The type7 writer separately submits each internal count. This producer does not make the metadata word the game parser's complete-read loop bound.

**Confidence.** **High** for named original arithmetic, field order and buffer submission. Complete-read count-source correction is inherited from ALM-META-025/091; no fresh count-mutation or native transaction is claimed here.

**Original status.** ✔ promoted

**Evidence.** [EXP-0353](../experiments/EXP-0353-trigger-storage/)

### ALM-RECVSOURCE-207

**The selected frontend receive path reads from entries requested through manager+0x18b8.** AtL02422/L02423,R0509 passes globalL00625 toR0513. The latter requests entries throughR0542/R0543 and passes each nonzero pointed-to S toL02424, returning the first nonzero result.L02424 first callsL02425(S). The latter locks S+0x1030, traverses S+0x1014 and sums each B+0x0f byte throughL02426; that sum is not the remaining-byte expression B+0x7f8 minus B+0x7fc inL02427. A nonzero sum permits a one-byte read request followed byL02428. The one-byte read's return is not tested.

**Confidence.** **High** for these pinned instruction and receiver paths; helper success, counter meaning, native initialization, producers and ordering **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0360](../experiments/EXP-0360-alm-receive-source/)

### ALM-RECVREAD-208

**The source reader consumes a linked sequence of buffer payloads, with separate cursor and removal operations.** Under the S+0x1030 lock,L02429 selects the head at list+4 throughL02430.L02431 advances its iteration pointer through node+0 and returns node+8, which the reader dereferences as B. It compares signed B+0x7f8 minus B+0x7fc against the requested length. For an exhausted short buffer it submits the remaining count toL02432, advances destination/count, and requests head removal and release throughL02433/L02434.L02432 bounds cursor+count against B+0x7f8, passes destination, B+0x10+cursor and count toR0544, then advances B+0x7fc. Copy, removal and release callees remain explicit frontiers; no native success or complete ownership proof is implied.

**Confidence.** **High** for the named list, argument, comparison and cursor operations. Native buffer invariants, lower copy/release behavior and complete-read guarantees **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0360](../experiments/EXP-0360-alm-receive-source/)

### ALM-RECVOBJ-209

**Received type6 selects a shared object and an in-place variable-length reader, separate from the type4 request object.** The 191 admitted selector values atL02428 use72 targets. Types6/7/8 selectL02435 and globalL02436; the shown initializerL02437/L02438 installsL02439. Its+0x0c slotL02440 requests 5 bytes into R+0x0a, then unsigned byte R+0x0e plus 1 bytes into R+0x0f. Neither read return is tested.L02441 sets R+9 from the originally read type byte and returns R. Across the 188 admitted frontend selector values, only 6 reachesL02442; it passes R+0x0f toR0514 only while view+0x80 is zero. GlobalL02436 has no file-backed initial value; constructor execution and later vtable replacement remain native conditions.

**Confidence.** **High** for the table, constructor-slot, read-argument and conditional-consumer facts. Complete input, untouched/stale bytes after failed reads, native lifecycle and request/response pairing **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0360](../experiments/EXP-0360-alm-receive-source/)

### ALM-SUBMIT-210

**The selected type4 request serializes nine bytes into an indexed send-side region;L02443 is a counter update request.**R0545 writes type4 at Q+9, destination zero at Q+7 and its argument at Q+0x0a in globalL02444. ConstructorR0546 installsL02445; its+8 slotL02446 passes Q+9 and the+0x10 slot's constant9 toL02447. The latter selects S+0x10+(S+0x1010 shifted left11), requests storage throughL02448, and on its short-capacity arm callsL02449 before a remainder call. These helpers remain frontiers.L02443 requests a key fromL02450 and lookup throughL02451 at manager+0x18d4; on success it adds the submitted size to the pointed-to dword. The selected body does not itself request a payload copy or list append. Distinct source formulas do not establish a live delivery path to S+0x1014 or a causal type4-to-type6 relation.

**Confidence.** **High** for the concrete serialization, indexed receiver and conditional addition; actual storage, transport, aliases, callback effects and native causal order **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0360](../experiments/EXP-0360-alm-receive-source/)

### ALM-EDITTIME-205

**The EN editor interprets metadata+0x0c as Starting Time in minutes and uses it in its color pipeline.** Original document+0x50 is E; the Light commandL02452 passes E to dialog+0x5c. Resource137 names control1082 Starting Time; DDXL02453 binds that same trackbar. InitializationL02454/L02455 reads E+0x1c and signed-divides by 15; the requested position range is 0..96. IDOK dispatch reachesL02456, which stores 15 times the returned position atL02457. After accepted modal return and enabled globalL02458,L02459 passes E+0x1c toL02460. That call-free helper computes signed (value/60)%24 and lighting outputs. Its RGB bytesL02461..0x48a enter mode2 palette generationL02462, whoseL02463 stores packed colors. RebuildL02464 publishes palette objects atL02465; rendererL02466 reads their color data andL02467 writes those colors to its clipped render buffer. Existing loader/copy/writer evidence binds E+0x1c to wire+0x0c. Defaults360 encode06:00. The UI requests quarter-hour steps; native clamping and actual control returns are not measured.

**Confidence.** **High** for the positive static control, arithmetic and render-buffer dataflow in the pinned EN editor;23 selected complete bodies,8 slices and 86 independently checked instruction sites. Native editor transactions/presentation, extra renderer receivers and ROM1 M/P gameplay effects remain **Unknown**.

**Original status.** ✔ promoted

**Evidence.** [EXP-0358 semantic result](../experiments/EXP-0358-alm-headers-metadata/SEMANTIC-RESULTS.md); editor-semantic products

## Authored Player control (continued)

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-HDRLIFE-206 | The last record-header wire word has no value consumer or surviving original object/global copy in the enumerated normal reader paths. | High / Medium / Unknown | ✔ promoted | [EXP-0358 semantic result](../experiments/EXP-0358-alm-headers-metadata/SEMANTIC-RESULTS.md); semantic-header products |

### ALM-HDRLIFE-206

**The last record-header wire word has no value consumer or surviving original object/global copy in the enumerated normal reader paths.** Its actual locals are primary frame offset -0x24, browser -0x14, light entry stack offset -0x10 and the editor's -0x24; editor version1000 instead uses -0x30. Primary version1000 transfers no record header. Complete reader/frame inventories find no direct operand overlapping those locals. The incidental H address left by the separate file-header REP copy is followed through 70 ROM and 93 editor callee entries from 138/170 alias-live root calls. Constructor-produced stream/app/map callbacks and exact switches leave zero unbound normal original calls. Ten ROM and 22 editor normal saved-register bodies restore their 21/37 return sites with no semantic use or surviving saved-slot escape. Next complete header overwrite or normal frame release ends local ownership, without erasing bytes. The original editor writer submits its own preexisting last-word stack storage ten times without a defined initializer or incoming-header copy.

**Confidence.** **High** for the pinned local transfers, complete-body inventories, exact saved slots and conditional original call bindings; **Medium** for normal-path no-use under ordinary ABI, DF=0, constructor-produced receivers and in-bounds buffers. Numeric representation/intent, native transactions, exceptions, malformed inputs and external pointer replacement remain **Unknown**. The allocation-failure callback requires a handler installed outside the identified loader; its original slot is zero.

**Original status.** ✔ promoted

**Evidence.** [EXP-0358 semantic result](../experiments/EXP-0358-alm-headers-metadata/SEMANTIC-RESULTS.md); semantic-header products

## Original mission survey: catapult owner, loot and opening view

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-CATAPULT-213 | In 90.alm the catapult at (102,14) belongs to player 8 "Catapult", which is non-hostile to Self and holds 0 or 2 toward every player; no trigger names it. | High / Medium | ● active | [EXP-0478](../experiments/EXP-0478-mission-scripts/) |
| ALM-LOOT-214 | In 111.alm the only unit with authored loot is the brigand leader, unit 91 at (80,116); the one Troll-class unit, unit 119 at (10,73), has no type-8 link, and the six type-8 records are the leader's stock and five ownerless ground ... | High / Medium | ● active | [EXP-0478](../experiments/EXP-0478-mission-scripts/) |
| ALM-VIEW-215 | In 140.alm the Rood Glaen body (unit 443, authored hp 0) and the tower near it are 227 cells from the Drop Location; the script and map data author no reveal, so neither is visible at start under local-only stamping. | High / Medium | ● active | [EXP-0478](../experiments/EXP-0478-mission-scripts/) |

### ALM-CATAPULT-213

In 90.alm the catapult at (102,14) belongs to player 8 "Catapult", which is non-hostile to Self and holds 0 or 2 toward every player; no trigger names it. Population: `90.alm`, identical in the preserved EN and RU `scenario.res`. The catapult is the type-6 record at (102,14), placement class 26 (Data.bin Units row 26, `Catapult`, weapon `Boulder Thrower`, spell `Fire_Ball:70`), unit id 290, group 114, owner slot 8, player name `Catapult`. The castle Troll stands at (107,15). The players are Self, Monsters, Beists, Nocturnal, Villagers, Friends, Enemies, Catapult, Dragon (1-based). The type-5 `Rel` rows fill the matrix by `AI-DIPLO-004`. Bit 0 is hostile and bit 1 locks the relation. The Catapult row is `0,2,2,2,2,2,2,2,2`, so the Catapult is non-hostile to every player; Self and the four players holding 0 toward it (Monsters, Beists, Nocturnal, Villagers) can flip their cell toward it when it is attacked (`HERO-AGGRO-028`), and Friends, Enemies and Dragon cannot. The relation to the human is `Self -> Catapult` 0 and `Catapult -> Self` 0. Monsters, Beists, Nocturnal and Villagers hold 0 toward the Catapult; Friends, Enemies and Dragon hold 2. The diplomacy-change instant (10) is authored in none of 90, 110 and 120. No trigger of 90 reads unit 290 or group 114, so no script orders the catapult to act. **Confidence.** **High** for the owner and the matrix (type-4/5/6 decode, 2 locales). **Medium** that it does not attack at start: the matrix and the empty script exclude a hostile target or a scripted order, while the standing acquisition of a non-hostile unit against hostile castle defenders was not run. **Unknown.** Whether the catapult engages Enemies or Monsters owned units by its own AI; the matrix shows no hostility to them either.

**Confidence.** High / Medium

**Original status.** ● active

**Evidence.** [EXP-0478](../experiments/EXP-0478-mission-scripts/)

### ALM-LOOT-214

In 111.alm the only unit with authored loot is the brigand leader, unit 91 at (80,116); the one Troll-class unit, unit 119 at (10,73), has no type-8 link, and the six type-8 records are the leader's stock and five ownerless ground sacks, none linked to a unit. Population: type-6 records (81) and type-8 records (6) of `111.alm`, EN; RU identical. Type-8 loot with a nonzero owner stocks an existing actor; the link is the type-6 `+0x18` word (`ALM-SACK-065`). Unit 91 (class 14, owner 2, group 4) carries gold 10000 and two items (0x7173, 0xe08). The five other type-8 records have owner 0 and lie on the ground at (9,72), (83,83), (56,94), (15,13) and (62,131). The sack at (9,72) is adjacent to the troll at (10,73) and the sack at (83,83) to the ghost at (83,82); no field links either sack to that unit, and a player can attribute an item found beside a creature to it. The troll-class placements are one Troll, class 68, `class2` 3, Data.bin Units row 86, owner 3 Beists, group 13, unit 119 at (10,73), and one Ogre, class 66, `class2` 3, row 82, unit 53 at (53,49). Units with x >= 85 are Bat at (85,86), (102,82) and (94,98), Bee at (90,132), (95,135), (102,133) and (101,124), and the Orc at (87,36); no Troll, Goblin, Ghost or Catapult row stands there. The map data gives no compass direction for the x and y axes, so no troll is placed in an "east" or "west" and the geometry is stated only as cell coordinates. A troll's money comes from the death draw of `MISSION-DROP-074`, tier 3 row 86 for unit 119. **Confidence.** **High** for the loot link table, the type-8 record list and the unit census (decoded records, both locales). **Medium** that unit 119 is the troll the owner saw (the only Troll-class unit in the map). **Unknown.** The palette meaning of `class2` 3 and the compass direction of the axes.

**Confidence.** High / Medium

**Original status.** ● active

**Evidence.** [EXP-0478](../experiments/EXP-0478-mission-scripts/)

### ALM-VIEW-215

In 140.alm the Rood Glaen body (unit 443, authored hp 0) and the tower near it are 227 cells from the Drop Location; the script and map data author no reveal, so neither is visible at start under local-only stamping. Population: `140.alm`. Rood Glaen is `npc26`, type-6 record 124, unit 443 at (242,115), owner 5, authored hp 0x0000. Trigger T00 fires when his Health is at least 1: message 8, owner change to Self, Mission Complete. The tower is the type-4 record at (243,112), kind 55, owner 4. The Drop Location is (15,114). The map has 24 action nodes (18 bound to a trigger, 6 unreferenced) and 12 triggers. The nodes are message, change owner, mission complete, group command, phase in and out of a group, spell effect and Drop Location; none is a fog or reveal operation. The action vocabulary has 34 opcodes (`TRIG-ACT-004`) and the arm list of `EXP-0080` names no fog or reveal arm, but arms 7, 23 and 24 call routines that were located and not read. No shipped map authors fog bits (`TERR-FOG-087`); the stamp comes only from the local participant's drawables (`TERR-TILE-079`). **Confidence.** **High** for the data and for the absence of a reveal node among the 24 action nodes of this map. **Medium** for the absence of any reveal operation in the whole vocabulary (three arms unread) and that neither object is visible at start: bit 3 shared vision (`TERR-FOG-089`) and the radius of the Drop Location group were not run. **Unknown.** Visibility under a shared-vision player flag.

**Confidence.** High / Medium

**Original status.** ● active

**Evidence.** [EXP-0478](../experiments/EXP-0478-mission-scripts/)

## Map editor: a new map and save

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-NEWMAP-236 | A new EN-editor map saves type-0 +0x08 0x3f490fda, +0x0c 360, +0x10 21, +0x14 63, +0x1c 2, +0x20..+0x2c 0, +0x70 4, +0x74 1, name "Multiplayer Map", +0x18 the plane's group mask; its text slots are not written. | High / Medium | ✔ promoted (amended) | [EXP-0543](../experiments/EXP-0543-editor-writes/) |
| ALM-NEWMAP-237 | A new EN-editor map saves 8 records 0,1,2,3,5,9,8,7 with tag 5 in every record header; W and H are the dialog values plus 16, 32..256 in steps of 8, 80 by default. | High / Medium | ✔ promoted | [EXP-0543](../experiments/EXP-0543-editor-writes/) |
| ALM-SAVE-238 | Every EN-editor save sets the top 8 tile rows to the last cell's word, clears tile bits 10-12 and 14-15 and recomputes type-0 +0x18..+0x2c; all 72 installed maps already satisfy the plane and mask rules. | High / Medium | ✔ promoted | [EXP-0543](../experiments/EXP-0543-editor-writes/) |
| ALM-SAVE-239 | An EN-editor save drops types 4 and 6 when no object exists, round-trips unit wire+0x18 in range, moves orphan loot to unit 1, passes some casters to unread helper L14474 and out of the size sum, binds an unmatched owner to E+0x3c. | High / Medium | ✔ promoted (amended) | [EXP-0543](../experiments/EXP-0543-editor-writes/) |
| ALM-SAVE-240 | The EN editor's reader stores type-0 +0x08..+0x14, the name, +0x70, +0x74 and the 512-byte text region where its writer reads them, so an unedited save writes those loaded bytes back unchanged. | High / Medium | ✔ promoted (amended) | [EXP-0543](../experiments/EXP-0543-editor-writes/) |
| ALM-SAVE-252 | The EN editor's reader stores every script node (796 bytes) and trigger (184 bytes) whole, in file order, and its writer emits the same offsets, so an unedited load and save writes both lists back byte for byte (Medium end to end). | High / Medium | ● active | [EXP-0554](../experiments/EXP-0554-editor-script-and-newmap/) |
| ALM-SAVE-253 | Reader case 9 stores a type-9 entry's fields where the writer reads them, but drops an entry whose +8 and +0xc are non-zero and whose word +0x10 is 4 when L15129 returns 0; +4 and +0x12 are not on the wire. | High / Medium | ● active (amended) | [EXP-0554](../experiments/EXP-0554-editor-script-and-newmap/) |
| ALM-SAVE-254 | At save the type-9 size loop passes an entry with +8 non-zero, +4 zero and 16-bit +0x16 zero to L14474 (thunk L14524), which removes it from the array; the next entry is skipped by the size sum but still written. | High / Medium | ● active (amended) | [EXP-0554](../experiments/EXP-0554-editor-script-and-newmap/) |
| ALM-SAVE-255 | Every EN-editor save writes file version 990 and, per type-8 loot, a position taken from its owner object and owner 0 for one class; all 72 installed maps are version 990 and 43 of 43 unit-owned loot per root keep their position. | High / Medium | ● active (amended) | [EXP-0554](../experiments/EXP-0554-editor-script-and-newmap/) |
| ALM-SAVE-256 | On load a player's id is its file ordinal (+4 = index + 1), not a stored field, and E+0x3c is the first player, so a unit whose owner matches none is bound to player 1 and saved with owner 1; none of 12,085 installed units is unmatched. | High / Medium | ● active (amended) | [EXP-0554](../experiments/EXP-0554-editor-script-and-newmap/) |
| ALM-NEWMAP-257 | A new EN-editor map's text region after its first byte comes from an allocation nothing in the path clears (High), and its name after the NUL does unless copy R0260 pads (Medium), so saved blocks can carry heap residue. | High / Medium | ● active (amended) | [EXP-0554](../experiments/EXP-0554-editor-script-and-newmap/) |
| ALM-NEWMAP-258 | The two players of a new EN-editor map are built by L14488 with id and +0x08 equal to 1 and 2, names Self and Enemy, +0x30 equal to 0 and 1, +0x0c equal to 0 and a 16-word relation row holding 2 at its own slot and 1 elsewhere. | High / Medium | ● active (amended) | [EXP-0554](../experiments/EXP-0554-editor-script-and-newmap/) |
| ALM-SAVE-260 | Reader case 9 drops a type-9 entry when L15129 (L15130) finds no object whose footprint holds the cell; of 206 EN and 60 RU entries, 48 and 11 sit on a type-4 record's cell, 158 and 49 are undecided; none is proved kept or dropped. | High / Medium | ● active (amended) | [EXP-0557](../experiments/EXP-0557-editor-helpers/) |
| ALM-SAVE-261 | L14520 returns the first object whose word +4 equals its argument, 0 included; no installed unit or plain type-4 record has id 0 (the type-4 object id is inferred), so 137 EN and 133 RU owner-0 loot are Medium. | High / Medium | ● active (amended) | [EXP-0557](../experiments/EXP-0557-editor-helpers/) |
| ALM-SAVE-262 | A type-4 record is 20 bytes, or 28 when its dword +8 is 0x21; the reader makes no structure for a cell at or below 7 or outside the width-8 and height-8 bounds; all 71 installed payloads parse exactly. | High / Medium | ● active (amended) | [EXP-0557](../experiments/EXP-0557-editor-helpers/) |
| ALM-SAVE-263 | L15212 takes the backward path when dest lies inside the source range, R0260 copies to the NUL, L15251 zero-fills a new array, and L15219 gives a loaded player a vtable, +0x0c, +0x30 and a zeroed 32-word array. | High / Medium | ● active (amended) | [EXP-0557](../experiments/EXP-0557-editor-helpers/) |
| ALM-SAVE-264 | L14524 has five rel32 callers in .text; four sit in loops over the type-9 array or a loot's dword array; the open handler L02376 builds the map through L02308 and L15303 (unread); the RU root has no map editor program. | Medium | ● active (amended) | [EXP-0557](../experiments/EXP-0557-editor-helpers/) |
| ALM-SAVE-268 | Reader case 4 builds a plain structure via L02332: +4 and the hash key are the wire id, +0x1a the kind, +6/+0xa cell*256+0x80, +0x12 the owner; class text "CStructure", 0x106 bytes; the body read has no test refusing id 0. | High / Medium | ● active | [EXP-0558](../experiments/EXP-0558-editor-structures/) |
| ALM-SAVE-269 | A kind 0x21 record is built by L02333: +4 is the smallest unused id from 1, the hash key the first free key from 1, never the wire id; its two extra dwords become +0x106 and +0x10a; the class name text "CVerticalWoodenBridge" is Medium. | High / Medium | ● active | [EXP-0558](../experiments/EXP-0558-editor-structures/) |
| ALM-SAVE-270 | The writer emits each structure from +6, +0xa, +0x1a, +0xf6, the owner's +4 and +4 (a bridge adds +0x106 and +0x10a, Medium), keys from 1 so key 0 is last; the round trip is Medium: kind range, footprint rejection and record order are open. | High / Medium | ● active | [EXP-0558](../experiments/EXP-0558-editor-structures/) |
| ALM-SAVE-271 | At load each loot whose +0x3c is 0 when the fix-up loop runs, up to the count [ebp-0x18], gets an object (class text "CBackPack") with kind 3, a new id and its loaded position; the save then writes owner 0 and that position. | Medium | ● active | [EXP-0558](../experiments/EXP-0558-editor-structures/) |
| ALM-SAVE-272 | L15303 is the reader L02303; its first file object (CreateFileA/ReadFile) reads into the records unchanged; L02308, L14494, L14495 touch two map-size buffers in the code read; L15252, L15261 look like floating-point code. | High / Medium | ● active | [EXP-0558](../experiments/EXP-0558-editor-structures/) |
| ALM-SAVE-273 | [L15131] is a pointer array indexed by an entry key and filled with entry objects built field by field by a keyed reader; its source, key text and footprint derivation are unread, so footprints and the type-9 load test stay undecided. | Medium / Unknown | ● active | [EXP-0558](../experiments/EXP-0558-editor-structures/) |

### ALM-NEWMAP-236

Population: the EN `Map Editor.exe`, SHA-256
`697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`, read statically. Path:
new-map command `L14475`, allocation `0x32c` bytes through operator new `L14476`, wrapper
`L02380`, constructor `L02381`, writer `L02295`.

| type-0 field | writer source | new-map value | store |
|---|---|---|---|
| +0x08 | `E+0x18` | `0x3f490fda` | `L14477` |
| +0x0c | `E+0x1c` | `0x168` (360) | `L02382` (`ALM-METADEFAULT-186`) |
| +0x10 | `E+0x20` | `0x15` (21) | `L14478` |
| +0x14 | `E+0x24` | `0x3f` (63) | `L14479` |
| +0x18 | recomputed from the tile plane (`ALM-SAVE-238`) | `1 << ((w & 0x3ff) >> 6)` for the one word `w`; `0xf00` for Water `0x211` | writer |
| +0x1c | `E+0x30`, the roster count | 2 | two `CPlayer` appended (`L14480`, `L14481`) |
| +0x20 | objects of class `L14482` in the map's object map | 0 | writer count |
| +0x24 | objects of class `L14483` | 0 | writer count |
| +0x28 | `E+0x4c + E+0x80 + E+0x68` (`ALM-EDITCOUNT-204`) | 0 | empty lists |
| +0x2c | `E+0x2e4`, the loot count | 0 | empty array |
| +0x30 (64 bytes) | `E+0x94` | `Multiplayer Map` and its NUL | `strcpy` from `L14484` |
| +0x70 | `E+0xd4` | 4 | `L14485` |
| +0x74 | `E+0xd8` | 1 | `L14486` |
| +0x78 (512 bytes) | `E+0xdc` | first byte 0 | `L14487` |

The two players are built by `L14488` (not read) with arguments (1, 1) and (2, 2); the second
gets the name `Enemy` (from `L14489`) at `+0x10` and 1 at `+0x30`. The roster count 2 is set
by the array-grow routine `L14490` (`L14491`, presumed, not read). The name store is a
string copy `R0260` (presumed, not read). The two object counts come from a class test
`L14492` against `L14482` and `L14483` (presumed, not read). The write routine
`L02297` is also presumed.

The bytes after the name's NUL, and `E+0xdd`..`E+0x2db` (the rest of the 512-byte region, which
holds the seven text slots of `ALM-META-028`), are written by none of the five bodies read on the
path. The block comes from operator new `L14476`, which calls `L14493` (presumed `malloc`,
not read) and fills nothing.

**Confidence.** High for the literal stores and the writer's sources. Medium for the roster
count 2 (`L14490`), the name copy (`R0260`), the object counts 0 (`L14492`) and the byte
stream the write routine (`L02297`) emits: each depends on a library-style helper whose body
was not read. Medium for the text region: the path read is the handler, wrapper, constructor, `L14494` and
`L14495`, and no census of other writers between creation and save was made.

**Unknown.** The content of the unwritten bytes; the other fields of the two `CPlayer` records.

**Amended.** The bytes after the name's NUL and `E+0xdd`..`E+0x2db` come from an allocation nothing in the path clears, and the writer saves them unchanged (`ALM-NEWMAP-257`); the two players' fields are in `ALM-NEWMAP-258`. The clauses "the content of the unwritten bytes" and "the other fields of the two `CPlayer` records" of the Unknown above are answered there.


### ALM-NEWMAP-237

1. Writer `L02295` emits types 4 and 6 only when `E+0x310` (the object count) is non-zero
   (`L14496`, `L14497`, `L14498`), so a new map has `recordCount` 8 and the order
   0,1,2,3,5,9,8,7. Sizes: type 0 632, type 1 `2*W*H`, types 2 and 3 `W*H`, type 5 152 (two
   players of 76 bytes), type 9 4 (a zero count), type 8 0, type 7 12 (three zero counts).
2. Every record header carries `E+8` as its tag (`L14499`); the constructor stores 5 there
   (`L14500`). The reader's case 1 stores the type-1 record's tag into `E+8` (`L14501`), so a
   loaded map keeps its own tag (7 in every installed record header). `ALM-FRAME-031` and
   `ALM-SEC-002` state the tag as the constant 7, which is the installed value only.
3. The file header carries `0x14`, dataSize `4*W*H+0x48`, the record count and 990
   (`ALM-WRITER-099`); the record-header `+0x10` word is a stack slot the writer never stores
   (`ALM-WRITER-099`).
4. `W` and `H` are the two edit values plus 16 (`L14502`..`L14503`). `OnOK` clamps each edit
   value to 16..240 and rounds it down to a multiple of 8. Both spin controls carry
   `UDS_SETBUDDYINT` and `UDS_AUTOBUDDY` (dialog `0x82`), and `OnInitDialog` (`L14504`) sets
   range 16..240 and position 64 after the base call `L14505`, so the edits show 64 and the
   default map is 80 x 80.

**Confidence.** High for the record conditions, the tag stores and the size arithmetic. Medium
for the type-5 size 152 and the record count 8 as bytes: the roster count passes through
`L14490` and the stream through `L02297`, neither read (presumed array grow and write
routine). Medium for the default 80: it rests on the base `OnInitDialog` (`L14505`, not read) running the data
exchange before the position is set, and on the up-down control's buddy update.

**Unknown.** Native dialog behaviour.

### ALM-SAVE-238

Writer `L02295`, on every call (loaded or new map):

1. `L14506`..`L14507`: cells `0 .. 8*W-1` of the tile plane, the top 8 rows, take the word of
   the last cell `W*H-1`.
2. `L14508`: every tile word is ANDed with `0x23ff`, clearing bits 10..12 and 14..15.
3. `L14509`..`L14510`: `+0x18` is the OR of `1 << ((w & 0x3ff) >> 6)` over all cells; a
   value above `0xff` gains `0xf00`; a value above `0x1fff` raises a message box and becomes
   `0x1fff`.
4. `+0x1c`..`+0x2c` are written from the roster count, the two object-class counts, the
   type-7 sum and the loot count (`ALM-NEWMAP-236`'s table).

Steps 1 and 2 store into `E+0x0c`, so the open map changes too. The reader's case 0 reads
`+0x18`, `+0x1c`, `+0x20`, `+0x24`, `+0x28` and `+0x2c` into locals (`L14511`..`L14512`);
none reaches the writer.

Installed population: 38 EN and 34 RU maps (loose `.alm` and every `M7R` member of
`scenario.res`). In all 72 the top 8 rows equal the last cell, no cell has a bit of 10..12 or
14..15, `+0x18` equals the recompute, dataSize is `4*W*H+72`, every record-header tag is 7, and
each map carries one record-header `+0x10` value across its records.

**Confidence.** High for the writer and reader instructions. Medium that the installed maps were
last written by this writer: the checks agree with it in every map, but another producer with the
same rules exists. RU `Horror.alm` has 4 records (`ALM-CORP-060`), a structure this writer cannot
emit (it writes types 0 to 3, 5, 9, 8, 7 and types 4 and 6 only with objects), and it satisfies
the plane and mask rules.

**Unknown.** Whether an installed map ever carried a different top band or mask before its last
save. What produced RU `Horror.alm`.

### ALM-SAVE-239

1. Types 4 and 6 are written only when `E+0x310` is non-zero (`L14496`, `L14497`,
   `L14498`): a loaded map whose objects were all deleted loses both records.
2. Type-6 wire `+0x18`: the reader stores it at `+0xd6` and subtracts 1 (`L14513`..
   `L14514`); the writer writes `+0xd6` plus 1 and restores the member. The wire word
   round-trips. Wire `+0x08` and `+0x0a` are read as signed 16-bit and stored as dwords at
   `+0x1a`, `+0x1e` (`L14515`..`L14516`) and written back as 16 bits.
3. Type-6 owner: the reader matches wire `+0x14` against each player's `+4` (`L14517`); with
   no match the unit is bound to the player at `E+0x3c` (`L14518`), whose identity on a loaded
   map is Unknown (the constructor sets `E+0x3c` to 0; reader case 5, which may set it, was not
   read). The writer writes the owner player's `+4`, so an unmatched owner is written as the
   id of whichever player `E+0x3c` names.
4. Type 8: for each loot record the writer looks up its owner unit by `+0x3c` through `L14519`
   (`L14520`, not read). With no match it shows a message box (format at `L14521`), takes
   unit 1, stores the loot index into that unit's `+0xd6` (`L14522`) and sets the loot's
   `+0x3c` to 1 (`L14523`).
5. Type 9: before the size sum, entries with `+8` non-zero, `+4` zero and the 16-bit `+0x16`
   zero are passed to `L14524` (`L14474`, not read) and left out of the size
   (`L14525`..`L14526`); the loop index is not adjusted after the call. The listing shows
   the skip in the size sum; it does not show a deletion. `L14524` has four other call
   sites in the same family (`L14527`, `L14528`, `L14529`, `L14530`), so its role is not
   settled by the writer call alone. The write loop would still write a surviving entry.

**Confidence.** High for the conditions, the `+0xd6` pair and the owner match, which are named
instructions. Medium for steps 3 to 5 as file changes: who `E+0x3c` is (reader case 5 not read),
what `L14520` matches, and whether `L14474` deletes the entry were not read. Medium is kept.

**Unknown.** The player at `E+0x3c` on a loaded map: a read of reader case 5 settles it.
Whether `L14474` deletes the entry, and whether a left-out caster is still written by the
later write loop: a read of `L14474` and `L14524` settles both.

**Amended.** Item 3: `E+0x3c` is the first player and an unmatched owner is saved as owner 1 (`ALM-SAVE-256`). Item 4: the owner lookup `L14520` is the first object in the hash at `E+0x304` whose word `+4` equals its argument (`ALM-SAVE-255`). Item 5: `L14474` removes the entry from its array, the write loop does not write it, and the entry that moves into its slot is skipped by the size sum (`ALM-SAVE-254`); no installed entry meets the predicate. The Unknowns of this card on the three points are answered there; the Medium of items 3 to 5 stays only where those cards keep it. Item 2 is narrowed: wire `+0x18` round-trips only when the stored `+0xd6` is below the loot count; the reader stores -1 for a value that is -1 or not below it and the writer writes that as 0 (`ALM-SAVE-255` item 4). The number of installed units whose `+0x18` is out of range was not measured.

**Amended.** EXP-0562: item 5's predicate (`+8` non-zero, `+4` zero, word `+0x16` zero) is met by entries the editor's caster dialogs can make (spell None, index 0; `ALM-SAVE-254` Amended), so "no installed entry meets the predicate" does not bound maps saved after editing; whether `+4` is zero in such an entry rests on the constructor `L15206`, unread (Medium).

### ALM-SAVE-240

Reader `L02303` case 0 (`L14531`..`L14532`) reads W and H into `E+0`, `E+4`; `+0x08`,
`+0x0c`, `+0x10`, `+0x14` into `E+0x18`, `E+0x1c`, `E+0x20`, `E+0x24`; 64 bytes at `+0x30` into
`E+0x94`; `+0x70`, `+0x74` into `E+0xd4`, `E+0xd8`; 512 bytes at `+0x78` into `E+0xdc`. The
writer reads each of these from the same E offset (`ALM-NEWMAP-236`'s table). The record tag
round-trips through `E+8` (`ALM-NEWMAP-237`); the record-header `+0x10` word does not.

**Confidence.** High for the reader and writer instruction pairs. Medium for "unchanged":
the editor's dialogs write these fields by design, and no census of other writers was made.

**Unknown.** Reader cases 4, 5, 7, 8 and 9 were not read here; their round trip rests on
`ALM-PLACEEDITOR-112`, `ALM-129` and `ALM-TRIGEDITOR-164` where those apply.

**Amended.** The Unknown on reader cases 4, 5, 7, 8 and 9: cases 5, 6, 7, 8 and 9 were read in `ALM-SAVE-252` to `ALM-SAVE-256`; case 4 (structures) was not read.

**Amended.** Case 4 (structures) is read for the reader in `ALM-SAVE-262`: a type-4 record is 20 bytes, or 28 when its dword at `+8` is `0x21`, and `L02332` and `L02333` create no object for a cell of 7 or less or beyond the width-8 and height-8 bounds. The writer's type-4 emission was not re-read there, so the type-4 round trip stays Unknown.

**Amended.** `ALM-SAVE-270` reads the writer's type-4 emission: the round trip of a plain record is Medium (the kind range, the footprint rejection and the record order are open). The reader-side statements here are unchanged.


### ALM-SAVE-252

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`, 1,855,488 bytes), read statically: reader
`L02303` case 7 (`L15196`..`L15197`), the writer `L02295` head (`L02295`..`L15198`) and its type-7
emission (`L15199`..`L15200`).

1. Reader. It reads the action count (4 bytes), then per node allocates `0x31c` bytes (constructor `L15201`) and
   reads 0x40 bytes to `+0x00`, 4 to `+0x40`, 4 to `+0x44`, 4 to `+0x318`, 0x28 to `+0x48`, 0x28 to `+0x70` and
   0x280 to `+0x98`; these seven reads cover the 0x31c bytes once. The node is appended at the tail of the list at
   `E+0x40`. The condition count and nodes are read the same way into the list at `E+0x5c`. Then it reads the
   trigger count, sizes the array at `E+0x78` to that many 0xb8-byte elements (freed for 0, otherwise zero-filled)
   and reads 0xb8 bytes into each element. No compare, clamp, id remap or opcode test occurs in these bodies: a
   node with opcode `0x10002` (the drop node of `MISSION-DROP-002`) is an ordinary node.
2. Writer. The record size is `(E+0x4c + E+0x68) * 0x31c + E+0x80 * 0xb8 + 12`; it writes the count `E+0x4c`,
   walks the list from its head and writes the seven fields from the offsets of item 1 in the reader's order, then
   the same for the list at `E+0x5c` (count `E+0x68`), then the count `E+0x80` and each trigger's 0xb8 bytes. The
   writer head reads the lists only to add `E+0x4c + E+0x80 + E+0x68` into type-0 `+0x28` (`ALM-EDITCOUNT-204`).
3. Consequence. With no edit, every node and trigger byte written is the byte read, in the same order, and the
   record size equals a canonical loaded payload (`[nAct]` nodes, `[nCond]` nodes, `[nTrg]` triggers,
   `ALM-TRIG-044`).

**Confidence.** High for items 1 and 2 (named instruction pairs). Medium for the unedited-save statement: the load
command and the code between a load and a save (the load handler `L02376` and any dialog) were not read.

**Unknown.** A file with two type-7 records: the second read appends nodes and resizes the trigger array to its own
count (read, not exercised).
### ALM-SAVE-253

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`), reader case 9 (`L15202`..`L15203`) and the writer's type-9
emission (`L15204`..`L15205`), read statically; derived counts over 38 EN and 34 RU installed maps.

1. Reader. It reads the entry count; per entry it allocates 0x2e bytes (constructor `L15206`), reads 4 bytes to
   `+0x00` only when the file version is at least `0x3de`, then 4 to `+0x08`, 4 to `+0x0c`, 6 to `+0x10`, 4 to
   `+0x16`, a count `m` and `m` elements of 6 bytes appended to the array at `+0x1a` (elements at `+0x1e`, count at
   `+0x22`). When `+0x08` and `+0x0c` are non-zero and the word at `+0x10` is 4, it calls `L15129(E, +0x08, +0x0c)`
   (`TERR-253`: thunk to `L15130`, body not read here); a zero result skips the append, so the entry is not in the
   array at `E+0x2f0` (the body frees nothing; that the helper looks up a structure at the cell is an inference from the description of its caller in `TERR-253`, Medium). Otherwise the hash at `E+0x304` is probed with the low 16 bits of the
   result, `+0x12` is set from the word at `+4` of the found object and the entry is appended. When the condition is
   false the entry is appended unchanged. If the probe finds nothing, `+0x12` is taken through a local the path
   does not assign (`L15207`..`L15208`).
2. Writer. It writes the entry count `E+0x2f8`, then per entry 4 bytes from `+0x00`, 4 from `+0x08`, 4 from `+0x0c`,
   6 from `+0x10`, 4 from `+0x16`, the count `+0x22` and the `m` 6-byte elements from `+0x1e`. `+0x04` and `+0x12`
   are not written. `+0x04` is set only by reader case 8: for a loot triple with `cc != 0` the array element
   `cc - 1` gets `+0x00` = the triple's index and `+0x04` = the loot object (`L15209`..`L15210`).
3. Installed maps (`corpus.txt`): the load-drop precondition (`+0x08` and `+0x0c` non-zero, first 16 bits of the
   six bytes equal to 4) holds for 206 of 792 EN and 60 of 334 RU entries. Whether the lookup fails for any of them
   was not decided. Every installed map is version 990.

**Confidence.** High for the field pairs and the drop branch (named instructions in both bodies). Medium for the
stale-local reading of the failed probe. Unknown below.

**Unknown.** Whether `L15129` returns 0 for any installed entry. If an entry were dropped, the loot triples that
follow in file order (type 8 is read after type 9) index the shorter array by their file numbering; the index
arithmetic of that case is not settled.

**Amended.** `ALM-SAVE-260`: `L15129` is read; of the 206 EN and 60 RU entries, 48 and 11 sit on a type-4 record's cell, 158 and 49 stay undecided, none proved kept or dropped (conditions in `ALM-SAVE-260` item 4).

### ALM-SAVE-254

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): `L14474` read whole, the writer's type-9 loops
(`L15204`..`L15211`), `L15212` read to its forward path; 38 EN and 34 RU installed maps for the census.

1. Thunk and body. `L14524` is `JMP L14474`. `L14474(E, i)` works on the array at `E+0x2f0` (data `+4`,
   count `+8`): it moves the `count - i - 1` pointers above `i` down one slot with `L15212`, lowers the count by 1,
   then for each remaining slot `k >= i` whose entry has `+4 != 0` sets element `entry+0` of the loot object's array
   (`[+4] + 0x18`) to `k + 1`. It calls nothing that frees the entry (the one-element loop at `L15213` has no
   call). No other branch.
2. Size loop. For `i` below the live count `E+0x2f8` (re-read each pass) an entry with `+8 != 0`, `+4 == 0` and the
   16-bit word at `+0x16` equal to 0 goes to `L14474(E, i)` and adds nothing; any other entry adds
   `0x1a + 6 * [+0x22]`. The index is not adjusted after the call, so the entry that moved into slot `i` is neither
   tested nor added.
3. Write loop. It uses the live count and array, so the removed entry is not written, and the entry that moved into
   slot `i` is written without having been added to the record size. The record size field is then smaller than the
   bytes written by `0x1a + 6m` per skipped entry. The removal also changes the open map, as the other save
   normalisations do.
4. Installed maps (`corpus.txt`): the removal precondition (`+0x08 != 0` and the low 16 bits of `+0x16` equal to 0)
   holds for 0 of 792 EN and 0 of 334 RU entries, with or without the loot link, so no installed map is changed by it.

**Confidence.** High for items 1 to 3 (every branch of `L14474` and of both loops read). Medium for `L15212` as
a memmove: its forward path (taken for a destination below the source) was read, its backward branch `L15214`
and its tails were not. Whether the game's loader reads the type-9 record size is not read.

**Unknown.** The four other call sites of `L14524` (`L14527`, `L14528`, `L14529`, `L14530`, `ALM-SAVE-239`)
were not read.

**Amended.** `ALM-SAVE-263` reads both branches of `L15212`; `ALM-SAVE-264` locates the four other rel32 call sites in loops that compare a type-9 entry's `+0x12` with an object's word `+4` or walk a loot's dword array (Medium: read at the call sites only; a delete role is inferred).

**Amended.** EXP-0562 (`ALM-EDITCAST-295`, `296`, `297`): the editor's caster dialogs make entries that can meet the item 2 predicate. Add in the Traps Manager and in the Unit Spellbook accepts the spell combo's first item None (index 0; `r02` of EXP-0562: 29 strings, None first, resources 282 and 283) and stores that index in word `+0x16`, with `+8` the clicked cell's x; Cast from structure appends an entry with spell 0 when it is opened and is read to keep it on Cancel (Medium: the base-class Cancel is unread). With `+4` zero (the entry constructor `L15206` is unread, Medium) the writer drops such an entry and the entry that moves into its slot is skipped by the size sum. Item 4's census (0 of 792 EN and 0 of 334 RU installed entries) is unchanged; it states installed maps, not maps the editor makes.

### ALM-SAVE-255

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): writer head and type-8 emission (`L15205`..`L15215`),
`L14520`, reader head and cases 6 and 8; 38 EN and 34 RU installed maps.

1. Version. The writer stores `0x3de` (990) in every file header (`L02296`). The reader refuses a version above
   `0x3e9`, fewer than 3 records or a magic mismatch, and gates fields on the file version: `0x3b7` (unit `+0xf8`,
   `+0x10e`), `0x3d9` (`+0x110`), `0x3db` (`+0x159`), `0x3dd` (loot `+0x48`), `0x3de` (type-9 `+0x00`). A file
   below a gate is written with the field from the constructor. All 72 installed maps are version 990.
2. Owner lookup. `L14519` (`L14520`) returns the first object in the hash at `E+0x304` whose word at `+4` equals
   its argument, or 0.
3. Writer, per loot. It looks the owner up by `loot+0x3c`; when none is found it shows the message at `L14521`,
   takes the object for id 1, stores the loot index in that object's `+0xd6` and sets `loot+0x3c` to 1. It then
   writes the count `[loot+8]`, the owner (0 when the class test `L14492(owner, L15216)` holds, else
   `loot+0x3c`), the owner object's `+6` and `+0x0a`, `loot+0x48`, and per triple the dword array `[loot+4]`, the
   word array `[loot+0x2c]` and the dword array `[loot+0x18]`.
4. Reader. Case 8 reads the count, `+0x3c`, `+0x40`, `+0x44`, `+0x48` (version 0x3dd or later) and the triples. The
   loaded `+0x40` and `+0x44` are therefore not written back: the position comes from the owner object. Reader case 6
   (`L15217`..`L15218`) sets `loot[unit+0xd6].+0x3c` to the unit's word `+4` when `unit+0xd6` is in range, so a
   loot referenced by a unit takes that unit's id; an `+0xd6` at or above the loot count is stored as -1 and written
   as 0.
5. Installed maps (`corpus.txt`): loot EN 180, RU 176. Owner 0: EN 137, RU 133. Owner equal to a type-6 unit's id:
   EN 43, RU 43, and in all 86 the wire position equals that unit's position. The type-4 structure layout assumed
   for a further match was refuted by the size check (3 EN and 1 RU type-4 payloads are not a multiple of 20), so
   no structure match is reported.

**Confidence.** High for items 1 and 4 (named instructions). Medium for items 2 and 3: the thunk `L14519` is not
read in this experiment; its pairing with `L14520` is carried from `ALM-SAVE-239`, the body of `L14520` is read
and the call convention matches, and item 3 is the branch taken on that lookup. Medium for item 5 (the census
matches the loot owner against the unit id, the 16-bit word at unit wire offset 64, as the writer's `+4` word).

**Unknown.** The body of the thunk `L14519`. Whether `L14520` finds an object for owner 0 (an object whose word `+4` is 0): if not, the 137 EN and
133 RU owner-0 loot records take the message-and-unit-1 branch at every save. Which objects pass the class test.

**Amended.** `ALM-SAVE-261` reads the thunk `L14519` and the class test; no installed unit has id 0 and no plain type-4 record has wire id 0 (the type-4 object id is an inference, Medium). `ALM-SAVE-262` explains the type-4 size check (28-byte kind `0x21`).

**Amended.** `ALM-SAVE-271` bounds the Unknown on owner 0 (Medium): the reader builds a class `L15216` object for each loot whose `+0x3c` is 0 when the fix-up loop runs, up to the count `[ebp-0x18]`, so the save lookup finds it, writes owner 0 and writes that object's position, which is the loaded `+0x40` and `+0x44`; item 4's statement that the loaded position is not written back holds for loot with an owner object of another origin, and for owner-0 loot the fix-up reaches. The loop bound, the loop body `L15408`..`L15409` and the interaction with the unit link `+0xd6` are unread. The class test objects are answered there.

### ALM-SAVE-256

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): reader head, cases 5 and 6, the new-map constructor
`L02381` and the writer's type-5 and type-6 emission; 38 EN and 34 RU installed maps.

1. Case 5. Per type-5 record of 76 bytes the reader allocates a 0x48-byte player (constructor `L15219`, thunk
   `L15220`, not read), reads 4 bytes to `+0x08`, 4 to `+0x30`, 4 to `+0x0c`, 0x20 to `+0x10` and sixteen 2-byte
   words to the array at `+0x38`, sets `+0x04` to the record index plus 1 (`L15221`..`L15222`) and appends the
   player to the array at `E+0x28`. After the loop `E+0x3c` is set to element 0 of that array (`L15223`..`L15224`).
   The writer emits `+0x08`, `+0x30`, `+0x0c`, `+0x10` (0x20) and the sixteen words; `+0x04` is not on the wire.
2. Initial and other stores. `E+0x3c` is 0 at the reader's start (`L15225`). In the reader windows read (head,
   cases 4 to 9) no other store to `E+0x3c` exists. A new map sets `E+0x3c` to its first player (`L15226`).
3. Case 6. A unit's wire owner (4 bytes, offset `0x14` at version `0x3db` or later) is compared with `+0x04` of each
   player; a match stores that player at unit `+0x12`; no match stores `E+0x3c` (`L15227`..`L15228`). The
   writer writes `+0x04` of the unit's player. An unmatched owner therefore becomes player 1 and is saved as owner 1.
4. Installed maps (`corpus.txt`): the type-5 record precedes the type-6 record in 71 of 72 maps (the 72nd, a
   four-record RU map, has no type 5 or type 6); owners outside `1..playerCount` occur in 0 of 8,094 EN and 0 of
   3,991 RU units.

**Confidence.** High for items 1 to 3 (named instructions). Medium for "no other store to `E+0x3c`": only the reader
and constructor windows were read, not every editor command. Medium for the order dependence: if a type-6 record
preceded the type-5 record, `E+0x3c` would still be 0 when case 6 runs (not seen in the corpus).

**Unknown.** The constructor `L15219`; `E+0x3c` stores outside the windows read.

**Amended.** `ALM-SAVE-263` item 4 reads the constructor `L15219`: it stores the vtable `L02281`, `+0x0c` = 0, `+0x30` = 0 and the base `L15246` at `+0x34`, and sizes that array to a count of 0x20 words. The reader stores `+0x04`, `+0x08`, `+0x0c`, `+0x10`, `+0x30` and the first 16 words from the wire; the array count, the vtables and the 16 words after the first 16 are constructor values. The `E+0x3c` stores outside the windows read stay Unknown (`ALM-SAVE-264`: no displacement scan was run).

**Amended.** `ALM-SAVE-268` item 4 and `ALM-SAVE-272` item 3 read two more stores to `E+0x3c`: the creation body `L02332` (`L15410`) and the reader's post-loop default player (`L15411`), each building a player by `L14488(1, 1)` when `E[0x3c]` is 0. Item 2's "no other store" holds for the windows of this card only. Item 3's "an unmatched owner becomes player 1" and "`E+0x3c` is the first player" hold for a file whose type-5 record precedes the use; otherwise the default player, built with `+4` = 1, is the one stored.
### ALM-NEWMAP-257

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): new-map handler `L14475`, wrapper `L02380`, constructor
`L02381`, operator new `L14476`, `L14493`, the small-block allocator `L15229`, its page commit `L15230`,
the free path `L15231`, the writer head; 38 EN and 34 RU installed maps for a census.

1. Allocation. The handler pushes `0x32c` and calls operator new (`L15232`); `L14476` calls `L14493`, which
   calls `L15233` and `L15234`. A size at most `[L15235]` (`0x3f8` in the image; its only store is the setter
   `L15236`, which has no direct caller in a call scan of `.text`) goes to `L15229`, otherwise `HeapAlloc` with
   flags 0. `0x32c` is 812, so the small-block path applies if the image value holds at runtime (Medium: a store through an unscanned indirect call is not excluded). Neither branch clears the block. `L15229` takes a free block of the size class from its
   group lists, splits it, writes header and footer words and returns the block plus 4; it contains no fill.
   `L15230` commits fresh pages with `VirtualAlloc` (the OS supplies zero pages) and writes only header, footer and
   link words; the free path `L15231` writes link and header words and calls no fill (its call list: two indirect calls
   through a register at `L15237` and `L15238`, `HeapFree` at `L15239`, `L15212` at `L15240`). A freed block
   therefore keeps its bytes.
2. Fill. The wrapper calls the constructor and allocates the two shade planes; the constructor stores the embedded
   collections, the dialog values, `E+0x18`..`E+0x24`, `E+0x8c`, `E+0x90`, the players, a copy through `R0260` (not read; a string copy is presumed, Medium) of
   `Multiplayer Map` to `E+0x94` (16 bytes with the NUL), the byte `E+0xdc` = 0, `E+0xd8` = 1 and `E+0xd4` = 4. No
   call to a fill routine exists in `L15232`..`L15241`, `L02380` or `L02381`.
3. Output. The writer emits 0x40 bytes from `E+0x94` and 0x200 bytes from `E+0xdc` unchanged. The 48 bytes after the
   name's NUL and the 511 bytes after the first text byte, which hold the seven text slots of `ALM-META-028`, are
   whatever the block held; for the name tail this holds unless `R0260` pads (Medium).
4. Installed maps (`corpus.txt`; shipped-file observation, not the editor's output): bytes after the first NUL are non-zero in the name of 29 of 38 EN maps (632
   bytes) and 29 of 34 RU maps (637), and in the text region of 29 of 38 EN (4,848) and 29 of 34 RU (4,889). The
   other 14 maps have all-zero tails. The text-region count is shipped-file content that the slot defaults explain
   (`ALM-META-028`) and is not residue evidence; the census does not say how those bytes were written.

**Confidence.** High that no store in the paths read clears the text region; Medium for the name tail (copy `R0260` unread). Medium that a given session's bytes are
non-zero: that depends on the heap history, which static reading cannot observe. Medium for the free path, whose
middle (`L15242`..`L15243`) was read by its call list only.

**Unknown.** Other writers of `E+0x94` and `E+0xdc` between creation and save (the map-properties dialogs write
them by design; no census of writers was made). The allocator's callee `L15244` past `L15245`. The copy `R0260` and the player base constructor `L15246` (not read).

**Amended.** `ALM-SAVE-263` reads `R0260` as a copy to the NUL (the part from `L15304` to the exit `L15305` is unread, so the Medium for the name tail stays) and `L15246` (a vtable and four zero words). `ALM-STRCOPY-293` reads the copy whole: the name tail is not stored by it (High).

### ALM-NEWMAP-258

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): constructor `L02381` (player stores `L15247`..`L15248`),
the player constructor `L14488` (thunk `L15249`) and the writer's type-5 emission.

1. `L14488(p, a, b)` stores the vtable, `+0x04` = `a`, `+0x08` = `b`, `+0x0c` = 0, `+0x30` = 0, copies the string
   at `L15250` (`Self`) to `+0x10`, sizes the array at `+0x34` to 0x20 bytes (`L15251`, not read) and sets word
   `i` of `+0x38` for `i` from 0 to 15 to 2 when `i == a - 1`, else 1. The new-map constructor calls it with
   `(1, 1)` and `(2, 2)`; the second player's name is `Enemy` (`L14489`) and its `+0x30` is set to 1.
2. Type-5 record per player (76 bytes, the writer's order): `+0x08` (1, 2), `+0x30` (0, 1), `+0x0c` (0, 0), the
   32-byte name (`Self`, `Enemy`; the bytes after the NUL are not stored by the copy), the sixteen words (player 1:
   2 then fifteen 1; player 2: 1, 2, then fourteen 1). `+0x04` is not on the wire.

**Confidence.** High for the stores and the writer's field order. Medium for the word row (the array's data pointer
comes from `L15251`, not read) and for the copy semantics of `R0260` (not read; a string copy is presumed).

**Unknown.** The constructor `L15219` that the reader uses for loaded players. The copy `R0260` and the player base constructor `L15246` (not read).

**Amended.** `ALM-SAVE-263` reads `L15251`, `L15246` and `L15219`. Item 1's "0x20 bytes" is a count of 0x20 words: `L15251(this+0x34, 0x20, -1)` allocates `2n` bytes (0x40) and sets size and capacity to `n = 0x20`, zero-filled before `L14488` sets sixteen words. The Medium for the copy `R0260` stays (its tail is unread; `ALM-STRCOPY-293` reads it whole, and the writer's emission is 16 words, `ALM-ROSTERIO-292`). Reservation: whether the writer's sixteen-word emission depends on the array count 0x20 or on the 16 words after the first 16 is Unknown.

### ALM-SAVE-260

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`), read statically: `L15129`, `L15130`, reader case 9 (`L15202`..`L15306`), `L15307`; 38 EN and 34 RU installed maps (`corpus2.txt`, rule R3 of Amendment 1).

1. Thunk and helper. `L15129` is `JMP L15130` (`w01`). `L15130(E, x, y)` takes the point `(x*256+0x81, y*256+0x81)`. With `E+0x310` (the object count) zero it returns 0. Otherwise it walks keys 0, 1, 2, ... through the hash at `E+0x304` (bucket `(key & 0xffff) >> 4` modulo the bucket count; node word `+8` is the key, node `+0xc` the object), stopping after `E+0x310` objects were found. For each object it tests, by `L14492` class: class `L14482` (rectangle from the object position at `+6` and `+0xa` to the position plus the footprint dwords `+0x10` and `+0x14` of the record `[L15131][object +0x1a]`, times 256), class `L14483` (record `[L15308][object +0x1a]`, dword `+0xd0` on both axes), class `L15216` (position plus `0x100`). The first object whose rectangle holds the point (all four bounds inclusive) makes it return the key; the end returns 0. A key of 0 cannot be told from a miss. That a type-4 object belongs to class `L14482` is inferred from its allocation size `0x106` (`w10`), which class `L15216` also has; the link is the unread `L15309` and the constructor thunk `L15310` (Medium).
2. Reader. Case 9 calls it with the entry's `+8` and `+0xc` only when both are non-zero and the word at `+0x10` is 4. A zero result goes to the next entry without appending the entry to the array at `E+0x2f0` (`L15311`, `L15312`, `L15313`, `L15314`). A non-zero result probes the hash with its low 16 bits, takes the found object's word `+4` into the entry's `+0x12` and appends. Case 8 indexes that array by `cc - 1` with no bound test (`L15209`..`L15210`), so a dropped entry shortens the array the loot triples index.
3. Tables. `[L15308]` is a pointer array filled at runtime (its size-set routine stores it at `L15315`, `L15316` and `L15317`, window `w07`). The stores to `[L15131]` are at `L15318`..`L15319`, outside every window read, and are known from the displacement scan `s01` only. The contents of both tables were not read. The footprint of a structure or unit type is therefore Unknown here.
4. Installed entries (`corpus2.txt`). 206 EN and 60 RU entries meet the call condition. A type-4 record sits on the entry's cell in 48 EN and 11 RU (class SA). An SA entry is kept only if (a) `L02332` made the object (it makes none when `cell + footprint` reaches width-8 or height-8, `w10` `L15320`, `L15321`; `corpus2.py` tests a cell of 7 or below only), (b) the first matching object's key is non-zero (a key of 0 reads as a miss; SA includes the kind `0x21` records, 8 EN and 1 RU, whose ids come from the unread `L15322`) and (c) its footprint, from the unread table, is at least one cell. Record order (type 4 before type 9 in each file) was not tested. None of the 48 and 11 is therefore proved kept. `corpus2.txt` prints no line for the counters UA (an entry on a type-6 unit cell) and ND (no object at or before the point), so no entry was counted in either class and no entry is proved dropped; 158 EN and 49 RU stay undecided (UN).

**Confidence.** High for items 1 and 2 (every branch of `L15130` and the case 9 call path read), except the class of a type-4 object (Medium, item 1). Medium for item 4 as kept or dropped: it needs the footprint tables, the creation of the object and its key (conditions a to c). The class names `CStructure`, `CBackPack` and `CUnit` are program text at the class records; no role beyond the class test is claimed.

**Unknown.** The contents of `[L15131]` and `[L15308]`; whether any of the 158 EN and 49 RU undecided entries is dropped; the key assigned to the first object (key 0); whether `L02332` made an object for each of the 48 EN and 11 RU type-4 records on a type-9 cell.

**Answers.** The Unknown of `ALM-SAVE-253` on whether `L15129` returns 0 for an installed entry: the helper is read, and the answer is bounded in item 4.

**Amended.** `ALM-SAVE-268` reads the constructor chain: a plain type-4 object is of class `L14482` (item 1 no longer an inference), its key is the wire id (0 for no installed record) and a bridge's key is the first free key from 1 (`ALM-SAVE-269`), so condition (b) of item 4 holds for every installed type-4 record; conditions (a) and (c) still depend on the unread footprint table (`ALM-SAVE-273`). The counts 48, 11, 158 and 49 are unchanged.

### ALM-SAVE-261

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): `L14519`, `L14520`, `L14492`, `L15323`, the class records at `L14482`, `L15216`, `L14483`, reader cases 4 and 6 and the structure-add routines `L02332` and `L02333`; 38 EN and 34 RU installed maps (`corpus2.txt`, R4 and R5).

1. `L14519` is `JMP L14520`. `L14520(E, id)` returns 0 when `E+0x310` is 0; otherwise it walks the hash at `E+0x304` bucket by bucket and returns the first object whose word at `+4` equals `id`. The comparison has no special case for 0. It does not test an object's class.
2. `L14492(obj, record)` calls the object's virtual slot 0 for its class record and returns 1 when the chain `[record +0x10]` reaches the argument. The three class records have size words `0x106` (`L14482`), `0x106` (`L15216`), `0x175` (`L14483`) and the same base `L15324`; the other two are not derived from `L15216`. Objects of class `L15216`, or of a class with base `L15216`, pass the writer's owner test. Users of the class record in `.text` (immediate scan, `s02`): `L15325`, `L15326`, `L15327`, `L15328`, `L15329`, `L15330`. Which code creates such an object was not read; reader cases 4, 6 and 8 do not.
3. Object ids on load. A unit's word `+4` is the wire u16 at offset 64 (the summed read sizes of case 6). A type-4 structure's hash key is the wire u16 at offset 18 (`f`, stored in the read body). That `f` is also the object's word `+4`, the id `L14520` compares, is an inference (Medium): `L02332` passes `f` to `L15309`, which stores `+4` and the type word `+0x1a` and is unread. A structure of kind `0x21` takes its id from `L15322(E)` (`L15331`; unread).
4. Installed maps (`corpus2.txt`). Owner-0 loot 137 EN, 133 RU. Units with id 0: 0 and 0. Type-4 records with wire id `f` 0 (kind other than `0x21`): 0 and 0. Records of kind `0x21`: 8 EN, 1 RU. Loot with a non-zero owner: 43 per root match a unit id. `corpus2.txt` prints no line for a match with a type-4 id or for no match, so none was counted in those classes.

**Confidence.** High for items 1 and 2 (bodies read) and for the wire sources in item 3 as the values the reader stores (unit `+4`, structure hash key); Medium that a type-4 object's `+4` is `f`. Medium for the owner-0 outcome: no unit has id 0 and no plain structure has wire id 0 in the installed maps (the structure object id is inferred from `f`), so the lookup at save finds none unless a kind-`0x21` structure (`L15322` unread) or an object made by an editor command carries id 0; the writer's branch on a failed lookup is `ALM-SAVE-255`.

**Unknown.** `L15322`; the code that makes a class `L15216` object; the type-8 emission was not re-read in this experiment (`ALM-SAVE-255` carries it).

**Answers.** `ALM-SAVE-255` items 2 and 3: the thunk `L14519` is read (item 1); the Unknown on which objects pass the class test is answered in item 2 as far as the class chain.

**Amended.** `ALM-SAVE-268` settles item 3: the plain type-4 object's word `+4` is the wire id word (`L02334`); `ALM-SAVE-269` reads `L15322` and the kind `0x21` ids. `ALM-SAVE-271` finds the code that makes a class `L15216` object: the loot fix-up at load, so the owner-0 lookup at save finds an object.

### ALM-SAVE-262

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): reader case 4 (`L15332`..`L15333`), `L02332`, `L02333` (read to `L15334`); 38 EN and 33 RU type-4 payloads (one RU map has none) (`corpus2.txt`, rules R1 and R2). The format surface is the type-4 record of `formats/alm/container.md`.

1. Wire layout read by case 4: u32 `a`, u32 `b`, u32 `c` (a kind), u16 `d`, u32 `e`, u16 `f` (20 bytes); when `c` is `0x21`, two more dwords follow (28 bytes).
2. Reader. For `c` other than `0x21` it calls `L02332(E, cell(a), cell(b), byte c, signed d, e, signed f)`; for `0x21`, `L02333(E, cell(a), cell(b), signed d, e, g, h)`. `cell(v)` is `(v - 0x80 + (0xff if negative)) >> 8` (arithmetic). `L02332` creates nothing when either cell is 7 or less, when `cell(a)` plus the footprint dword `+0x10` of `[L15131][c]` reaches the word `E[0]` minus 8, or `cell(b)` plus footprint `+0x14` reaches `E[4]` minus 8; `L02333` has the same bounds with `g` and `h`. A created structure is `0x106` bytes (kind `0x21`: `0x10e`) at position `cell*256+0x80` on both axes, so an `a` or `b` that is not `256k+0x80` is stored rounded; `f` is the hash key, stored in the read body (that it is also the object id is an inference, Medium: the `+4` store is in the unread `L15309`; kind `0x21`: the hash key is a generated id); `e` selects the owner player (`players[e-1]` when `e` is at most the player count, otherwise the first player, created if absent).
3. Installed maps. The 20/28 parse ends exactly at the payload size in 71 of 71 type-4 payloads (3,141 EN and 1,681 RU records; kind `0x21` in 8 EN, 1 RU). `corpus2.txt` prints no line for the counters of a parse mismatch, of a record with `a` or `b` off `256k+0x80` and of a record with a cell of 7 or less, so none was counted in those classes. This explains the 3 EN and 1 RU payloads that are not multiples of 20 (`ALM-SAVE-255`).

**Confidence.** High for items 1 and 2 for the reader, except the object id and class of the created object (Medium), and for item 3 as a shipped-file observation. Medium that `E[0]` and `E[4]` are the map width and height (the code names them; the type-0 load is `ALM-SAVE-240`). The writer's type-4 emission was not re-read here, so the type-4 round trip is Unknown. The footprint tables and the bodies of `L15309`, `L15335` are unread.

**Unknown.** The writer's type-4 record layout; what a structure whose key is already in the hash does to the earlier object (the node's object pointer is replaced, `L15336`..`L15337`).

**Answers.** The Unknown of `ALM-SAVE-240` on reader case 4 (structures) is answered for the reader; `ALM-SAVE-255` item 5: the type-4 size check is explained by the 28-byte kind.

**Amended.** `ALM-SAVE-268` and `ALM-SAVE-269` read `L15309` (`L02334`) and `L15335`: the wire id is the object's `+4` and the hash key for a plain record; a bridge takes both from allocators and its wire id is not used. A repeated key replaces the node's object pointer without freeing it in that body. `ALM-SAVE-270` gives the writer's type-4 layout. Item 2's `e` clause (the owner player) stays as stated.

### ALM-SAVE-263

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): `L15212` (`L15212`..`L15338` and `L15214`..`L15339`), `R0260` (`R0260`..`L15304`), `L15251` (to `L15340`), `L15246`, `L15219`.

1. `L15212(dest, src, n)`: when `dest` is above `src` and below `src + n` it takes the branch at `L15214` (backward `REP MOVSD` with the direction flag set, tails through the tables at `L15341` and `L15342`); otherwise it copies forward (`REP MOVSD`, tails through `L15343`, `L15344`, `L15345`, `L15346`). The selection of the backward path when `dest` lies inside the source range is read; the byte tails of both paths are unread (Medium).
2. `R0260(dest, src)` jumps to a copy loop that copies source bytes up to the NUL, aligned dword at a time with a NUL-byte test (`L02415`..`L15347`). No store beyond the NUL was seen in the part read; the tail from `L15304` to the exit `L15305` was not read.
3. `L15251(this, n, grow)` with `n` non-zero and no data: allocates `2n` bytes (`L14476`), zero-fills them (`L15348(ptr, 0, 2n)`) and sets size `+8` and capacity `+0xc` to `n`. `L15246` stores the vtable `L15349` and zero in `+4`, `+8`, `+0xc`, `+0x10` and nothing else.
4. The loaded-player constructor `L15219` (thunk `L15220`) calls `L15246` at `this+0x34`, stores vtable `L02281`, `+0xc` = 0, `+0x30` = 0 and calls `L15251(this+0x34, 0x20, -1)`. The reader stores `+4`, `+8`, `+0xc`, `+0x10`, `+0x30` and the first 16 words of the array at `+0x38` from the wire (`ALM-SAVE-256`). The array count and capacity 0x20 (the `2n` allocation of `L15251` holds 32 words), the vtable at `+0x34` and the 16 words after the first 16 are constructor values, not wire values. Whether the writer's sixteen-word emission depends on them is Unknown.

**Confidence.** High for items 3 and 4 and for the branch selection of item 1; Medium for the exact tail bytes of item 1 and for item 2 (tails unread). The word row of `ALM-NEWMAP-258` starts zeroed (item 3) before `L14488` sets it.

**Unknown.** The jump-table tails of `L15212` and the tail of `R0260`. Whether the writer's sixteen-word emission depends on the array count 0x20 or on the 16 words after the first 16.

**Answers.** The Unknowns of `ALM-SAVE-254` on the backward branch of `L15212`, of `ALM-SAVE-256` and `ALM-NEWMAP-258` on `L15219`, `L15246`, `L15251`, and of `ALM-NEWMAP-257` on the copy `R0260`, are answered here as stated.

**Amended.** `ALM-STRCOPY-293` reads `R0260` whole: it stores through the NUL and nothing past it, so the unread tail of item 2 is settled. `ALM-ROSTERIO-292` reads the writer's type-5 emission: 16 words from the data pointer with bound 0x10, independent of the array count 0x20.

### ALM-SAVE-264

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): rel32 call scan of `.text` for `L14524` and `L14474` (`s03`), the windows around the four other call sites (`w13`, read only around the sites), the open handler `L02376` and thunks (`w16`, `w17`, `w18`); a name listing of `gameversions/ru` for editor programs.

1. Callers. A `.text` rel32 scan finds `L14524` called from `L14527`, `L14528`, `L14529`, `L14530` and `L14526` (the writer), and `L14474` has no rel32 caller other than the thunk. Indirect calls, computed jumps and pointer tables are not covered; the dword scan for `L14524` and `L14474` as pointers was not run. At `L14527` and `L14529` a loop over the type-9 array (`E+0x2f4`) compares each entry's `+0x12` with the word `+4` of an object, calls `L14524` with the index and steps the index back; at `L14528` and `L14530` a loop over a loot's dword array (`[loot+0x18]`) calls `L14524(cc - 1)` for each non-zero `cc`. That these calls delete entries bound to an object or a loot is an inference from the call sites. The containing functions (`L15350`, `L15351`, `L15307`) were not read line by line.
2. Open handler. `L02376(path)` calls `L15352`, allocates `0x32c` bytes and runs the thunk `L15353` (to `L02308`), which calls `L15303(E, path)` (unread) and allocates two buffers of width times height bytes at `E+0x320` and `E+0x324`. The handler then calls `L15354` (to `L14494`) and `L15355` (to `L14495`) on the map and builds a second `0x32c` object through `L15356(E)` (to `L02361`) at `this+0x54`. Those bodies and `L15303` were not read.
3. Node reads. The reader calls the virtual slot `+0x34` of the object at `[ebp-0x20]` with a pointer and a size for every field; the class behind that slot was not resolved.
4. RU. A name `ls` for `Map Editor*` and a `find -maxdepth 2 -iname '*edit*'` under the RU root list no map editor program (consultation rows 10 and 11; names only, no other pattern or depth, no committed listing). The claims of this ledger about the editor are for the EN program only.

**Confidence.** High for item 1 as the rel32 scan result, within the blind spots it states. Medium for the loop descriptions and their delete role in item 1 (call-site neighbourhood) and for item 4 (the name patterns stated). Items 2 and 3 are Unknown beyond what is stated.

**Unknown.** `L15303`, `L14494`, `L14495`, the read primitive at slot `+0x34` and the branches `L15252` and `L15261` (not reached); the other writers of `E+0x3c`, `E+0x94`, `E+0xdc` (no displacement scan was run); the pointer-table callers of `L14524` and `L14474` (no dword scan was run).

**Answers.** The Unknown of `ALM-SAVE-254` on the four other call sites: located in item 1, loops read at the call sites only. The Medium of `ALM-SAVE-252` for the unedited load and save stays: the load handler's calls are unread.

**Amended.** `ALM-SAVE-272` reads `L15303` (the reader `L02303`), `L02308`, `L14494`, `L14495`, `L15252`, `L15261` and the class behind slot `+0x34` (a `ReadFile` file object): items 2 and 3 are answered there. The dword scan for pointer callers of `L14524` and `L14474` was not run in `EXP-0558` either.

**Amended.** EXP-0562 reads the bodies around three call sites: `L14527` lies in the structure delete `L15350`, `L14528` and `L14529` in the unit delete `L15351`; each calls `L14524` to remove a type-9 entry (`ALM-EDITLINK-298`, `ALM-EDITDEL-299`). The delete role of those three loops, Medium in item 1 here, is read in those bodies. `L14530` (in `L15307`) is not read there. EXP-0562's scans target other values, so the dword scan for pointer callers of `L14524` and `L14474` stays unrun.

### ALM-SAVE-268

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): reader case 4 (`L15332`..`L15412`, w05), `L02332` whole (w03, w04), the init `L02334` (w02), the thunk and constructor chain `L15413`, `L15310`, `L15414`, `L15415`, `L15416` (to `L15417`, w10 to w12), the class getters `L15418` (w39) and the class test `L14492` (w38), the default-player constructor `L14488` (w20); rel32 scans of `.text` (s01, s02); 38 EN and 34 RU installed maps (`corpus.txt`, C1, C2, C4, C5).

1. Call path. Reader case 4 reads per record X u32, Y u32, K u32, A u16, I u32 and ID u16 (20 bytes). For K other than `0x21` it calls the thunk `L15413` (to `L02332`) as `L02332(E, cx, cy, byte K, sign-extended A, I, sign-extended ID)`, with `cx = (X - 0x80 + (0xff if negative)) >> 8` and `cy` likewise from Y. The only rel32 caller of `L15413` in `.text` is `L15419` (scan s02; pointer tables, indirect calls and computed jumps are not covered, and no dword scan was run).
2. Creation gate. `L02332` returns without an object when `cx <= 7`, `cy <= 7`, `cx` plus the dword `+0x10` of the record `[L15131][K]` reaches `E[0] - 8`, or `cy` plus its dword `+0x14` reaches `E[4] - 8` (unsigned compares). The table contents are unread (`ALM-SAVE-273`).
3. Object. It allocates `0x106` bytes and runs `L15310` (to `L15414`: the base constructor `L15415` to `L15416`, then the vtable `L15420`) and `L15309` (to `L02334`). Slot 0 of that vtable returns the class record `L14482` (program text name `CStructure`, size word `0x106`). `L02334` stores `+4` = the low word of the id argument, `+0x1a` = the kind (a sign-extended byte), `+6` and `+0x22` = `cx*256+0x80`, `+0xa` and `+0x26` = `cy*256+0x80`, `+0xe`, `+0x66`, `+0x6a` and `+0xfe` = 0, `+0xfa` and `+0xf6` = the low word of A. The base constructor was read to `L15417`; its remaining stores are unread.
4. Owner. `+0x12` = `E[0x2c][I-1]` when `E[0x30] >= I`, otherwise `E[0x3c]`; when `E[0x3c]` is 0 a player is built by `L14488(1, 1)` (its `+4` and `+8` are 1) and appended to the list at `E+0x28`. An I of 0 indexes `E[0x2c][-1]`; no installed record has it.
5. Hash. The object is stored in the hash at `E+0x304` under the key `low16(ID)` (bucket `(key >> 4)` modulo the bucket count `[hash+8]`; node word `+8` the key, node `+0xc` the object). A key already present keeps its node and its object pointer is replaced by the new object; nothing in the body frees the old object. No test refuses key 0.
6. Installed maps (`corpus.txt`, shipped files only). Plain records: 3,133 EN and 1,680 RU. Wire id 0: 0 and 0. Distinct ids: 3,133 and 1,680 (no repeated key). I of 0: 0 and 0; I above the type-5 player count: 0 and 0 (type 5 precedes type 4 in 38 of 38 EN maps and in 33 of 33 RU maps that carry both, so `E+0x3c` is player 0 as in `ALM-SAVE-256`). X or Y off `cell*256+0x80`: 0 and 0. Cell of 7 or less: 0 and 0. The dimension bounds of item 2 were not measured: the corpus counters for them took W and H from the type-1 payload, which is the wrong record (the reader case that stores `E[0]` and `E[4]` is the type-0 record), so those counters are void.

**Confidence.** High for items 1, 4 and 5 and for item 3 as far as read, for the EN program. Medium for the base constructor tail (item 3) and for whether a footprint rejects a record (item 2). Item 6 speaks for the shipped files, not for what the editor produces in a session.

**Unknown.** The stores of `L15416` from `L15417`; the footprint tables; whether any installed record is rejected by the dimension bounds; whether an editor command, outside the load path, changes an object's `+6` or `+0xa` before a save.

**Answers.** The Unknown of `ALM-SAVE-262` on the stores of `L15309` and on a repeated key, and the Medium of `ALM-SAVE-261` item 3 on the type-4 object id: `+4` is the wire id word.

### ALM-SAVE-269

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): reader case 4 (`L15332`..`L15412`, w05), `L02333` whole (w06, w07), `L15421` whole (w07, w08), `L15422` whole (w08, w09), the thunks `L15423`, `L15335`, `L15322`, `L15424` (w02, w07), `L15425` (w11), the vtable getters `L15426`, `L15427` and `L15428`, `L15429` (w39, w41, w42), the class records and name strings (w37, w42); rel32 scans s01 and s02; 38 EN and 34 RU installed maps (`corpus.txt`, C3, C5).

1. Call path. For K equal to `0x21` the reader reads two more dwords E1 and E2 and calls the thunk `L15423` (to `L02333`) as `L02333(E, cx, cy, sign-extended A, I, E1, E2)`. The ID word is read and not passed. The only rel32 caller of `L15423` in `.text` is `L15430` (scan s02; same blind spots as `ALM-SAVE-268` item 1).
2. Gate. `L02333` returns 0 without an object when `cx <= 7`, `cy <= 7`, `cx + E1 >= E[0] - 8` or `cy + E2 >= E[4] - 8` (unsigned). No table is read.
3. Object. It allocates `0x10e` bytes, runs `L15335` (to `L15425`: the plain constructor `L15310`, vtable `L15431`, `+0x106` = E1, `+0x10a` = E2), stores the vtable `L15432` and runs `L15309` with the kind `0x21`, position `cell*256+0x80` and A as in `ALM-SAVE-268` item 3. Slot 0 of `L15432` (`L15433` to `L15426`) returns the class record `L15434` (size `0x10e`; its name pointer `L15435` is read in `w37`, `w39`). The name text `CVerticalWoodenBridge` and the base names `CBridge` and `CStructure` are program text read in `w42`, past the declared window cap, and are Medium. The getters in slots `+0x18` and `+0x1c` (`L15428`, `L15429`, read in `w41`, `w42`, past the cap) return `+0x106` and `+0x10a`; that reading is Medium.
4. Ids. The id argument is `L15424(E)` (to `L15422`): the loop returns the smallest integer from 1 that no object in the hash holds in its word `+4` (1 for an empty hash). The hash key is the word returned by `L15322(E)` (to `L15421`): the smallest key from 1 absent from the hash; it also raises `E[0x8c]` to the key plus 1 when the key reaches it. Neither is 0. `+0x12` is set as in `ALM-SAVE-268` item 4. The node is stored as in item 5 there (an existing key is reused).
5. Installed maps (`corpus.txt`). Kind `0x21` records: 8 EN, 1 RU; wire id 0: 0 and 0; the dimension bounds of item 2 were not measured (the corpus counter used the wrong record, see `ALM-SAVE-268` item 6). For 9 of 9 the wire id equals the id the allocation gives when every plain record of the map already holds its wire id (counter C3, counted over all plain wire ids of the map as `corpus.py` does, not over earlier records only) (type 4 precedes type 6 in 38 of 38 EN and 33 of 33 RU maps that carry both). That rule offers the allocator more used ids than a record order would, so the equality of the allocated and wire ids is Medium.

**Confidence.** High for items 1, 2 and 4 and for the stores of item 3 for the EN program (every branch of `L02333`, `L15421` and `L15422` read). Medium for the class name text, the base names and the getters of slots `+0x18` and `+0x1c` in item 3: the windows that read them (`w41` after its first lines, `w42`) lie past the window cap of the preregistration, which was passed in `w41` (`EXP-0558.md` Deviations). Medium for item 5.

**Unknown.** The file order of the plain records against the bridge records inside the payload; whether an editor command creates a bridge of the horizontal class (`L15436`, name text `CHorisontalWoodenBridge`); the stores of `L15416` from `L15417`.

**Answers.** The Unknown of `ALM-SAVE-261` on `L15322` and the Unknown of `ALM-SAVE-262` on the hash key of a kind `0x21` record.

### ALM-SAVE-270

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): the writer `L02295` from its head to `L15437` (w13 to w18), the class test `L14492`, the getters `L15428`, `L15429`, the player constructor `L14488` (w20); 38 EN and 34 RU installed maps (`corpus.txt`, C1 to C5). The part of the writer from `L15437` to `L15438` (units tail and later records) is unread.

1. Key walk. The writer finds objects by looking keys up one after another: `t` starts at 1 (`L15439`, `L15440`, `L15441`) and the lookup uses the low word of `t`, so the objects are met in ascending key order and a key-0 object only when `t` reaches 65536. The pre-pass stops when `E[0x310]` objects were found; the structure emission stops after the number of objects of class `L14482` counted by the pre-pass.
2. Record. The type-4 record is written only when `E[0x310]` is not 0, with the size `0x1c` for each object of class `L15442` (the bridge) and `0x14` for each other object of class `L14482`. Per object it writes, with the 4- or 2-byte size shown: `+6` (4), `+0xa` (4), `+0x1a` (4), `+0xf6` (2), the owner's `+4` (`[obj+0x12]+4`, 4), `+4` (2); for a bridge also the results of the slots `+0x18` and `+0x1c` (`+0x106`, `+0x10a`, 4 each; Medium: the getters `L15428` and `L15429` were read in `w41`, `w42`, past the window cap).
3. Consequences. X and Y come from the object, not from the file. The id written for a bridge is its allocated `+4` (`ALM-SAVE-269`), not the wire id. I is the `+4` of the owner player, which is the record ordinal plus 1 (`ALM-SAVE-256`) or 1 for the default player (`ALM-SAVE-268` item 4). The kind is the sign-extended low byte read.
4. Round trip. A plain record created by `ALM-SAVE-268` is written with the X, Y, A and ID it was read with, the same I when the player of ordinal I exists, and the same kind when the kind fits a signed byte. In the installed maps the position, id, I and A conditions hold for every plain record (`ALM-SAVE-268` item 6; A is non-zero in 2,979 EN and 1,589 RU records and the word survives the sign extension); the kind range, the creation rejection by footprint (`ALM-SAVE-273`) and the order of the records were not measured: the writer orders them by key, and whether the installed payloads are in ascending id order was not counted.

**Confidence.** High for items 1 and 3 and for item 2 except the bridge extent fields (Medium, see item 2) for the bodies read in the EN program. Medium for item 4: it composes the reader, the writer and shipped-file counts, and the kind range, the footprint rejection and the record order are open.

**Unknown.** Whether the installed type-4 payloads are in ascending id order; the kind range of the installed records; the writer from `L15437`.

**Answers.** The Unknown of `ALM-SAVE-262` on the writer's type-4 record layout; the Unknown of `ALM-SAVE-263` on whether the writer's emission depends on the constructor values stays open (the player records were not re-read); the Unknown of `ALM-SAVE-240` on the type-4 round trip, which is Medium here (kind range and record order open).

### ALM-SAVE-271

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): the reader's end (`L15203`..`L15408`, `L15409`..`R1507`, w32 to w35; the list-resize code `L15408`..`L15409` is unread), the writer's type-8 emission (`L15205`..`L15215`, w17, w18), the constructor `L15443` and the getter `L15444` (w36, w39), the class records and name strings (w37); 38 EN and 34 RU installed maps (`corpus.txt`, C6).

1. Load fix-up (Medium). When the reader has read all records and the loot count `[ebp-0x18]` is above 0, the loop (`L15445`..`L15446`) runs while its index is below `[ebp-0x18]`; for each loot `L` (the array at `E+0x2e0`) whose `+0x3c` is 0 at that moment it builds an object: `0x106` bytes, constructor `L15443` (base `L15415`, vtable `L15447`), whose slot 0 (`L15444`) returns the class record `L15216` (program text `CBackPack`). `L15309` stores the kind 3, `+6` and `+0x22` = `L[+0x40]`, `+0xa` and `+0x26` = `L[+0x44]`, `+4` = `L15424(E)` (the smallest unused id from 1). `+0x12` = `E[0x3c]` (a default player is built if absent), `+0xd6` = the loot index, `L[+0x3c]` = the new id, and the object is stored in the hash under the key of `L15322` (`ALM-SAVE-269` item 4).
2. Save (High for the object it names). For each loot the writer looks the owner up by `L[+0x3c]`. For the object of item 1 the lookup succeeds; the class test `L14492(owner, L15216)` holds, so the owner is written as 0, and the position is the object's `+6` and `+0xa`. The count `L[+8]`, the dword `L[+0x48]` and per item the dword `L[+4][j]`, the word `L[+0x2c][j]` and the dword `L[+0x18][j]` follow (10 bytes per item). Only when a lookup fails does the writer show the message at `L14521`, take the object with id 1, store the loot index in its `+0xd6` and set `L[+0x3c]` to 1.
3. Installed maps (`corpus.txt`). Loot: 180 EN, 176 RU. Owner 0: 137 and 133. Owner equal to a unit id: 43 and 43. Owner equal to a plain structure id or only to a bridge wire id: 0 and 0. Owner matching nothing: 0 and 0. The counts are of the wire owner. The fix-up tests `+0x3c` after the record loop, and reader case 6 stores a unit's id into `loot[unit+0xd6].+0x3c` (`ALM-SAVE-255` item 4); a loot with wire owner 0 that a unit references would not be tested as 0, and no count of owner-0 loot referenced by a unit was made. So the failed-lookup branch is not taken for any installed loot, and the 137 EN and 133 RU owner-0 loot are written with owner 0 and the position they were read with only for the loot the fix-up reaches, if `L[+0x40]` and `L[+0x44]` hold the wire position (`ALM-SAVE-255` item 4; the case 8 reads were not re-read here).

**Confidence.** Medium for item 1 and for the statement that each loot with wire owner 0 gets an object: the source of the loop bound `[ebp-0x18]`, the loop body `L15408`..`L15409` and the unit link `+0xd6` are unread. High for item 2 for the object it names. Medium for item 3 (the case 8 stores of `+0x40` and `+0x44` are unread; the census is of shipped files and of the wire owner).

**Unknown.** The source of `[ebp-0x18]` and whether it equals the number of loot read; the loop body `L15408`..`L15409` (the inline list-resize code, entered by the jump at `L15448`..`L15449`); the interaction of the unit link `+0xd6` (reader case 6) with the `+0x3c` test, and the number of owner-0 loot a unit references; the stores of `L15416` from `L15417`.

**Answers.** The Unknown of `ALM-SAVE-255` on whether `L14520` finds an object for owner 0 and on which objects pass the class test: the object of item 1 does (Medium, bounded as above). The Unknown of `ALM-SAVE-261` on the code that makes a class `L15216` object.

### ALM-SAVE-272

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): the reader `L02303` head and sections jump (`L02303`..`L15450`, w22), the loop end and post-processing (`L15203`..`L15451`, `L15452`..`L15408`, `L15409`..`R1507`, w32 to w35), the thunks (w21), the file objects (w23 to w27), `L02308` (w28), `L14494` (w29), `L14495` (w30), `L15252` and `L15261` (w31).

1. `L15303` is a thunk to the reader `L02303(E, path)`. It sets the sub-objects of `E` to empty states (`E[0]`, `E[4]` = 0x10, `E[0x8c]` = 1, `E[0x3c]` = 0, the hash at `E+0x304` with `0x11` buckets), opens the file and reads a header that must match the string at `L15453`, hold at least 3 records and a version of `0x3e9` or less. It then loops over the records, each dispatched on its type word through the table at `L15454`.
2. File object. The first object tried (0x10 bytes, constructor `L15455`, vtable `L02304`) opens by slot `+0x20` (`L15456`) with `CreateFileA` for reading; the read primitive of slot `+0x34` (`L02305`) calls `ReadFile` straight into the destination buffer and returns the byte count. In this object's open and read bodies no byte is changed; a short read leaves the rest of the destination as it was (High for the first file object only). Only if that open fails does the reader try a second object (0x20 bytes, constructor `L15457`, vtable `L15458`) whose open (`L15459`) and read (`L15460`, a read from an inner buffer) were read to their first lines only.
3. After the records, the reader requires the cell array `E[0xc]` and `E[0x14]`, zero-fills `E[0x10]` if absent, builds a default player if `E[0x3c]` is still 0 and appends it to the list at `E+0x28`, sets `E[0x94]` and `E[0xdc]` to 0 and `E[0xd4]`, `E[0xd8]` to 1 when no type-0 record came (the flag `[ebp-0x1c]` is cleared at `L15461` and set to 1 by the first dispatch target, the case at `L14531`, which reads W and H into `E[0]` and `E[4]`; the same flag guards these stores at `L15462`..`L15463`), and runs the loot fix-up of `ALM-SAVE-271` (Medium).
4. `L02308` calls `L15303`, then allocates two buffers of `E[0] * E[4]` bytes at `E+0x320` and `E+0x324`. `L14494` computes one byte per cell from the height bytes `E[0x14]` and writes only `E+0x320`; `L14495` passes `E+0x324`, `E+0x320` and `E[0] * E[4]` to `L14591` (`L15464`), read as a copy from the call's arguments (inference; the body is unread). In the code read, none of them reads or writes a structure, loot, type-9 entry, player, node or trigger; the unread callees keep that negative Medium (the nine call sites of `L14494` reach six routines, `L02308` calls `L14476`, `L15465` and `L15466`, and `L14495` calls `L14591`).
5. `L15252` stores the registers and ST0, calls `L15261` and restores the control word unless it is `0x27f`; `L15261` switches on an exception code and calls `L15467` and `L15468`. They hold no map record (C runtime floating-point exception code is an inference from the control-word handling and the code switch).
6. Between the file read and a later save, the structure records change only as `ALM-SAVE-268`, `ALM-SAVE-269` and `ALM-SAVE-271` state; no other step of the paths read edits them.

**Confidence.** High for items 1 and 2 for the bodies read of the first file object, for the buffer sizes and for the per-cell byte store of item 4. Medium for the negatives of item 4 and for "unchanged" beyond the first file object (callees unread, second file object read to its first lines only), for item 3 and for item 5 (callees unread) and for item 6 (the second file object, the inline list-resize code `L15408`..`L15409`, the open handler `L02376` and `L02361` are unread; the node and unit bodies were read by other claims).

**Unknown.** `L15459` and `L15460` beyond the windows; `L15469`; `L15467`, `L15468`; `L15408`..`L15409`; `L02361`; the callees `L15470`, `L15471`, `L14602`, `L15472`, `L15473`, `L14603` of `L14494` (nine call sites); `L14591` (called at `L15464`); `L14476`, `L15465`, `L15466` (called by `L02308`); the second file object's open and read (`L15459`, `L15460`), reached when the first open fails.

**Answers.** The Unknown of `ALM-SAVE-264` on `L15303`, `L14494`, `L14495`, the read primitive at slot `+0x34` and the branches `L15252` and `L15261`. The Medium of `ALM-SAVE-252` for the unedited load and save stays for the nodes (this claim reads no node code).

### ALM-SAVE-273

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): the readers of the table in `L02332` (w03) and `L15474` (w11), the routine around the stores `L15318`..`L15475` (`L15476`..`L15477`, `L15478`..`L15476`, `L15479`..`L15480`, `L15481`..`L15478`; w43, w45 to w47); the second block `L15482`..`L15319` and the routine start before `L15478` are unread. A first window `L15483` decoded nothing (w44).

1. Shape. `[L15131]` is a pointer array with the count `[L15484]` and capacity `[L15485]`. The stores at `L15318` (free), `L15486` (allocate) and `L15475` (grow) manage the array. The store of an element is `L15487`: `[[L15131] + m*4] = S`, where `S` is the entry object built in the current loop pass and `m` is its dword `+0xc`.
2. Entry. A loop over entry numbers builds `S` (`0xa4` bytes, constructor `L15488`) from a keyed reader (calls with a section name and a key string; whether it reads a file is not shown by the windows): `S+0xc` (default -1), `+0x10` (-1), `+0x14` (-1), `+0x18`, `+0x1c`, `+0x20`, `+0x24`, `+0x28`, `+0x2c` (-1), `+0x30` (0), strings at `+0x54` and `+0x68`, and `+0x64`, `+0x88`, `+0x8c`, `+0x94`, `+0x98`, `+0x90`, each through a key string in `.rdata` (`L15489`.. `L15490`); the key text, the section name format at `L15491` and the file name are unread. `S+0x10` and `S+0x14` are the dwords `L02332` adds to the cell; a text buffer of `S+0x10 * S+0x18 + 1` characters is read when `S+0x1c` is above 1 and `S+0x38` counts its characters other than `-`.
3. Not reached. The values of `S+0x10` and `S+0x14` for any entry, the file they come from, the stores of `[L15308]` (`L15315`, `L15316`, `L15317`; the routine `L15492`..`L15493` was read by an earlier experiment and is not re-read here), the type-9 load test on the 206 EN and 60 RU entries (no rule was fixed, no corpus step run), and the hashes of the data files of both roots.

**Confidence.** Medium for items 1 and 2 as the reading of the windows named. Unknown for item 3.

**Unknown.** Everything in item 3; the window that settles the file name is the prologue of the routine before `L15478` and the strings at `L15491` to `L15494`.

**Answers.** Part of the Unknown of `ALM-SAVE-260` item 3: the stores `L15318`..`L15475` are array management and an element store, not footprint values.

## Map field readers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-MAPLIFE-226 | Three callers of loader R0478 create the map object; each passes it to a fixed set of readers and destroys it, and only R0277 stores its pointer, at world+0x540d0. EN and RU rom.exe are identical. | High / Unknown | ✔ promoted | [EXP-0542](../experiments/EXP-0542-map-field-readers/) |
| ALM-UNITREAD-227 | No holder of a type-6 runtime record reads unit file +0x1c, +0x22, +0x26, +0x28, +0x2a, +0x30, +0x33, +0x34 or +0x35, and the stream read slot keeps no record pointer. | High / Medium | ✔ promoted (amended) | [EXP-0542](../experiments/EXP-0542-map-field-readers/) |
| ALM-STRUCTREAD-228 | Type-4 file +0x12 is sign-extended into record+0x10 and stored at the Shop or Building object's +0x08 (L14613), not +0x10; +0x0c and +0x0e reach the shop cap, obj+0x42 and the owner at obj+0x14. | High / Medium | ✔ promoted | [EXP-0542](../experiments/EXP-0542-map-field-readers/) |
| ALM-HDRTEXT-229 | After the loader stores the 0x200 header block at M+0xdc, no holder of the map object reads it; the map browser reads the block up to the first NUL, below +0x40 on 72 map instances (38 EN, 34 RU; 28 scenario.res members per root). | High / Medium | ✔ promoted | [EXP-0542](../experiments/EXP-0542-map-field-readers/) |
| ALM-HDRSLOT-230 | Header text slot 4 equals a trigger name on the same 13 scenario maps in EN and RU; no instruction found reads a slot, so no runtime reference from a script record to a slot exists in the holder set. | High / Medium | ✔ promoted | [EXP-0542](../experiments/EXP-0542-map-field-readers/) |
| ALM-ROSTERREAD-231 | The loaded map's holders read roster CPlayer +0x04, +0x08, +0x10, +0x30 and +0x34 only; the type-5 +0x08 scalar at CPlayer+0x0c is read only by the copy constructor that R0504 calls. | High / Unknown | ✔ promoted | [EXP-0542](../experiments/EXP-0542-map-field-readers/) |
| ALM-NODEREAD-232 | After load only binder R0067 reads script nodes: opcode, id, values and types, and file +0x48 for conditions only; the label is read only as text in two failure logs, and the parameter names are never read. | High / Medium | ✔ promoted | [EXP-0542](../experiments/EXP-0542-map-field-readers/) |
| ALM-TRIGREAD-233 | After load only binder R0067 reads the 184-byte triggers: condition pairs, action ids, comparison codes and the once flag; no holder reads the name or the +0x40..+0x7f span. | High / Medium | ✔ promoted | [EXP-0542](../experiments/EXP-0542-map-field-readers/) |
| ALM-EDITDISP-277 | EN editor: the Units, Structures and Objects menu items set the tool word [L14535] and call unread helpers; no object is made in their bodies; it is made on the map view's mouse-up. | High / Medium | ● active (amended) | [EXP-0560](../experiments/EXP-0560-editor-authoring/) |
| ALM-EDITUNIT-276 | EN editor unit tool: a mouse-up on a free cell makes a unit from the template [L15495] with a new id; its loot is cloned. | High / Medium | ● active (amended) | [EXP-0560](../experiments/EXP-0560-editor-authoring/) |
| ALM-EDITUNIT-278 | EN editor unit dialog (template 227): right-click (message 0x205, inferred) on a unit edits a copy and replaces it; on a free cell it edits T; the routine behind thunk L15496 (taken as OK, inferred) writes T fields. | Medium | ● active (amended) | [EXP-0560](../experiments/EXP-0560-editor-authoring/) |
| ALM-EDITDISP-279 | EN editor: the Objects tool writes one byte per cell, not a loot record; copy and cut work on a cell rectangle; tool 0x8053 opens dialogs captioned as spell and trap ones (program text); no loot path or unit delete in the bodies read. | High / Medium | ● active (amended) | [EXP-0560](../experiments/EXP-0560-editor-authoring/) |
| ALM-EDITNODE-280 | EN editor Instance and Check manager dialogs: Add new gives a 0x31c-byte node the id last+1 (else 1) in the dialog's list this+0xdc; Delete unlinks it there; identity with the map's script is Medium. | High / Medium | ● active (amended) | [EXP-0560](../experiments/EXP-0560-editor-authoring/) |
| ALM-EDITCAST-294 | EN editor tool 0x8053, case read from L15566: an empty cell opens the Traps Manager, a unit the Unit Spellbook, a structure Cast from structure, another object none; that the case is the right click is inferred (Medium). | High / Medium | ● active | [EXP-0562](../experiments/EXP-0562-editor-triggers/) |
| ALM-EDITCAST-295 | EN editor Traps Manager: Add makes a type-9 entry at the clicked cell with kind 0..3 (From + 2 x To), spell, power and two coordinate pairs, structure id 0; Remove deletes it; kinds 4 and 8 are not listed. | High / Medium | ● active | [EXP-0562](../experiments/EXP-0562-editor-triggers/) |
| ALM-EDITCAST-296 | EN editor Unit Spellbook: Add makes a type-9 entry of kind 8 at the unit's cell with the unit id, spell index and power; Remove deletes it; the list shows kind-8 entries of that unit id at a nonzero cell. | High / Medium | ● active | [EXP-0562](../experiments/EXP-0562-editor-triggers/) |
| ALM-EDITCAST-297 | EN editor Cast from structure: opening it on a structure with no kind-4 entry appends one (spell None, power 0); OK stores spell, power and scan range, OK with spell None removes it; that Cancel keeps it is Medium. | High / Medium | ● active | [EXP-0562](../experiments/EXP-0562-editor-triggers/) |
| ALM-EDITLINK-298 | EN editor: the type-9 removals read (L14474 and the inline loops of three caster dialogs) renumber later entries' loot back-references and leave the removed entry's own link; the unit delete removes linked entries, then the loot slot. | High / Medium | ● active | [EXP-0562](../experiments/EXP-0562-editor-triggers/) |
| ALM-EDITDEL-299 | EN editor Delete key (command 0x8063): on a structure it removes the structure and its kind-4 entries; on a unit it removes the unit, takes its loot records out of the list with their linked entries, and removes its kind-8 entries. | High / Medium | ● active | [EXP-0562](../experiments/EXP-0562-editor-triggers/) |
| ALM-EDITPASTE-300 | EN editor Paste (command 0xe125) only selects the paste tool; its mouse-up writes the pasted rectangle into three cell layers, each gated by a global switch, and no list; the Select objects dialog stores a template index. | High / Medium | ● active (amended) | [EXP-0562](../experiments/EXP-0562-editor-triggers/) |
| ALM-EDITCLONE-301 | EN editor: no command that copies an existing unit with its loot or type-9 entries was found in the handlers read; the accelerator table has 53 entries and most handlers are unread, so this is Medium. | Medium | ● active | [EXP-0562](../experiments/EXP-0562-editor-triggers/) |
| ALM-EDITOBJ-302 | EN editor Select objects: objects classes with nonzero InMapEditor, inserted in load order (shown order Medium); a pick stores +4+1. | High / Medium | ● active | [EXP-0563](../experiments/EXP-0563-editor-palettes/) |
| ALM-EDITSTRUCT-303 | EN editor Select structures: VariableSize 0 classes inserted by ascending ID (shown order Medium); a placed pick's +0x1a is its ID. | High / Medium | ● active | [EXP-0563](../experiments/EXP-0563-editor-palettes/) |
| ALM-EDITUNIT-304 | EN editor unit list: every units class inserted by ascending ID (shown order Medium), no key filter; a pick sets template +0x1a to it. | High / Medium | ● active | [EXP-0563](../experiments/EXP-0563-editor-palettes/) |

### ALM-MAPLIFE-226

Population: EN and RU `rom.exe`, byte-identical (SHA-256
`942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`); 27 ranges decoded or dumped,
each compared EN against RU, all equal. M is the `0x32c`-byte map object that `R0478`
constructs and loads.

- `R0478` has three direct callers and no dword reference in any section: `R0128` (the
  mission map load), `R0414` (the save-restore arm) and `R0277` (the world constructor).
- `R0128` passes M to `R0423` (roster), `R0118`, `R0489` (type-4 walker), `R0151`
  (type-6 spawner), `R0067` (script binder) and `R0461` (types 8 and 9), then deletes it
  through `L14427`, which calls destructor `R0522`.
- `R0414` passes M to `R0118` and `R0067`, then deletes it through `L14427`.
- `R0277` stores M at `world+0x540d0` (`L02291`), passes it to `R0278`, calls `R0522`
  and frees M. `R0118` passes M only to `R0278`, which reads M `+0x00`, `+0x04`, `+0x0c`,
  `+0x10` and `+0x14` only.
- Displacement `0x540d0` occurs twice in the decoded image, both in `R0277`.
- M's containers: roster list `+0x28`, action nodes `+0x40`, condition nodes `+0x5c`, triggers
  `+0x78` (`0xb8`-byte elements), type 8 `+0x2dc`, type 9 `+0x2f0`, type-4 records `+0x304`,
  type-6 records `+0x318`.
- The other map constructors `R0523`, `R0504` and `R0524` have no direct call and no
  dword reference in any section.

This is the holder set the other claims of this topic search: M's creators, the routines they
pass M to, and the destructor.

**Confidence.** **High** for the three creators, what each passes M to and the destruction:
named instructions in ranges read in both roots; `R0525`, the destructor's only escape of each
record, is read as a free by its pairing with allocator `R0747` and its body was not read here
(`ALM-TAILHOLD-080`) / **Unknown** for a computed alias of `world+0x540d0` used after `R0277`
returns (`ALM-ALIAS-154`) and for indirect entry into the three uncalled constructors.

**Unknown.** A computed call or alias that reaches M or a map constructor without a direct call,
a displacement or a dword reference.

### ALM-UNITREAD-227

Loader case 6 reads each record into a `0x50`-byte heap record that is not zeroed. The fields
asked about land at: file `+0x1c` → `rec+0x1c`, `+0x22` → `rec+0x22`, `+0x26` → `rec+0x26`,
`+0x28` → `rec+0x4c`, `+0x2a` → `rec+0x4e`, `+0x30` → `rec+0x2c`, `+0x33` and `+0x34` →
`rec+0x2f` and `rec+0x30`, `+0x35` → `rec+0x31`.

- The holders are the loader, spawner `R0151` and destructor `R0522` (`ALM-TAILHOLD-080`,
  re-read here).
- Spawner slot `[ebp-0x30]`: written once (`L14614`), loaded 50 times, every load
  dereferenced, no load passed on. Displacements: `+0x00`, `+0x04`, `+0x08`, `+0x0c`, `+0x10`,
  `+0x14`, `+0x20`, `+0x24`, `+0x28`..`+0x2b`, `+0x2d`, `+0x2e`, `+0x31+i` (`i=1..5`),
  `+0x37+i` (`i=0..4`), `+0x3c`, `+0x40`, `+0x48`.
- Loader slot `[ebp-0x90]`: 39 loads. It dereferences `+0x08`, `+0x0c`, `+0x14`, `+0x3c`,
  `+0x40` and `+0x44`, passes interior addresses to the stream read slot, passes the record to
  array append `L02248` (`L02252`) and frees the spare record after the loop (`L14615`).
- The stream read slot: vtable `L02300` slot `+0x3c` is `R0533`. It seeks the inner
  object at `this+0x10` through that object's `+0x30` and reads through its `+0x3c`, and keeps
  no copy of the buffer address. `R0534` passes the buffer to the `ReadFile` import. That the
  inner object is the file object rests on `ALM-STALE-196`.
- No holder dereferences `rec+0x1c`, `+0x22`, `+0x26`, `+0x2c`, `+0x2f`, `+0x30`, `+0x31`,
  `+0x4c` or `+0x4e`.
- Census of records whose value is not the absent value (0, or `0xffff` for `+0x22` and
  `+0x26`), EN of 8094 / RU of 3991: `+0x1c` 6891 / 3421, `+0x22` 30 / 30, `+0x2a` 348 / 348,
  `+0x35` 3 / 3; `+0x26`, `+0x28`, `+0x30`, `+0x33`, `+0x34` 0 / 0.

**Confidence.** **High** for the three holders and the stream read slot: every load of each
record slot is listed / **Medium** for no reader anywhere in the image: array append
`L02248`, which receives each record pointer, and the accessors `R1338` and `R2121`
were not read.

**Unknown.** Whether `L02248` or an accessor copies a record. An authored `+0x35`, `+0x22`,
`+0x2a` or `+0x1c` value has no game effect found; the editor's use is outside this population.

**Amended.** EXP-0564 read `L02248`, `R2270`, `R2121`, `R1338`, `R0521` and `L13132`: each stores,
returns or counts a record pointer and dereferences no record (`ALM-UNITREAD-313`). Of the Unknown, only the grow
routine `L15875` remains unread.

### ALM-STRUCTREAD-228

- Loader case 4 reads `+0x00`, `+0x04`, `+0x08` (4 bytes each), `+0x0c` (2), `+0x0e` (4) and
  `+0x12` (2) (`L02165`..`L14616`). It calls `R0481`, or `R0482` when `+0x08` is
  `0x21`, with X shifted right by 8, Y shifted right by 8, the kind word, the `+0x0c` word, the
  `+0x0e` low word and `+0x12` sign-extended (`L14617`).
- The `0x14`-byte record: `+0x00` X, `+0x02` Y, `+0x04` kind, `+0x06` file `+0x0c`, `+0x08`
  file `+0x0e` low word, `+0x10` file `+0x12` as a dword (`L14618`), `+0x0a` and `+0x0c` the
  extension words or 0.
- Walker `R0489`, record slot `[ebp-0x18]`: kind `0x22` or `0x23` builds a Shop (`R0490`)
  and passes `(i16)rec+0x06 × 1000` to `R0494` (`SHOP-CAP-004`). Any other kind builds a
  Building (`R0485`); when `(i16)rec+0x06` is 0 it stores 0 at `obj+0x42` (`L02330`).
- Both arms then pass the `rec+0x08` word to `R0425` and store the result at `obj+0x14`
  (`L02119`), logging a message when it is 0. Then `obj+0x08 = rec+0x10`
  (`L14619`/`L14613`), the object goes to `R1337`, and the record is freed (`L02327`)
  and its array slot cleared.
- File `+0x12` therefore reaches `obj+0x08`. `obj+0x10` holds the position-object pointer
  (`TERR-STRUCT-075`). The high word of `+0x0e` is dropped.
- Census: `+0x0e` high word 0 and `+0x12` non-zero on all 3141 EN and 1681 RU records. Building
  `+0x0c` values EN `0`:162, `1`:1, `3`:2, `5`:4, `10`:1, `100`:2921; RU `0`:92, `1`:1, `3`:2,
  `5`:4, `10`:1, `100`:1557. Shop records: EN 50 with 23 values, RU 24 with 15 values.

- Readers of `obj+0x08`: binder `R0067` takes the manager object at `[[this+0x24]]` that the
  walker hands every structure to (`L14620`, `R1337`), iterates it (`L14621`, `L14622`)
  and pushes `elem+8` as the key of the map at `[ebp-0x5c]` (`L14623`..`L14624`, call
  `L14625`), then stores the element in the map slot (`L14626`). The type-9 arm
  (`L14627`..`L14628`) looks the parameter value up in that map and logs a failure when the
  result is 0 (`L14629`), which is the `Target_Structure` resolution of `ALM-TRIG-046`.

**Confidence.** **High** for the loader reads, the record constructors and the walker stores,
which are named instructions in ranges read in both roots; the walker's in-memory record is
`ALM-CLS-037` / **Medium** for the readers of `obj+0x08`: the binder chain is read, but that
`[[this+0x24]]` is the list `R1337` fills, and what the map helper `L14625` does with the key
address, rest on routines not read.

**Unknown.** Whether `R1337` changes `obj+0x08`; the bodies of `R1337`, `L14625`,
`L14621` and `L14622`; any reader of `obj+0x08` outside the binder.

### ALM-HDRTEXT-229

- Loader case 0 reads the `0x200` bytes at type-0 `+0x78` into `M+0xdc` (`L14630`); the
  absent-section arm writes 0 there (`L14631`).
- The decoded-image scan finds 96 displacement and 65 immediate hits of `0xdc`. In M's holder
  set (`ALM-MAPLIFE-226`) only the loader touches `M+0xdc`. Of the uncalled constructors,
  `R0523` and `R0524` zero it and `R0504` copies it through `R1604`.
- The map browser's reader `R0479` (one direct caller, `R0511`) reads the `0x200` bytes
  from the file into a stack buffer, rewrites each `'\n'` to `'#'` up to the first NUL and
  assigns the buffer to a CString (`ALM-META-028`).
- Census: the block's first NUL lies below `+0x40` on 72 map instances (38 EN, 34 RU; 28
  `scenario.res` members per root), so the
  browser's string ends inside the description and contains no text slot.

**Confidence.** **High** for the holder set and the browser reader / **Medium** for no other
reader in the image: a reader that opens the file itself was searched only through the direct
calls to `R0479` and the `0xdc` displacement. That the string ends at the first NUL rests on
the rewrite loop, which stops there, and on the assignment `R1774`, whose body was not read
(`ALM-META-058`).

**Unknown.** A file reader of the block outside the browser, found by neither scan.

### ALM-HDRSLOT-230

- Census per root of the seven slots (text at slot `+12`): EN 90 slot texts that are neither
  empty nor `"<None>"`, RU 91. EN: 13 equal a trigger name of the same map (all slot 4; one also
  equals a node label), 8 equal another slot, 69 equal nothing. RU: 13, 8 and 70.
- The 13 maps are the `scenario.res` members `10`, `30`, `40`, `41`, `50`, `51`, `60`, `70`,
  `91`, `100`, `111`, `150` and `151`, in both roots.
- No instruction in M's holder set reads `M+0xdc` after the loader (`ALM-HDRTEXT-229`). No
  script parameter type code is a text reference (`ALM-TRIG-046`).

**Confidence.** **High** for the counts: one walk of every map of both roots / **Medium** for no
runtime reference from a script record to a slot, which rests on `ALM-HDRTEXT-229`'s Medium.

**Unknown.** Whether the map editor uses the slot texts; `Map Editor.exe` was outside this
population.

### ALM-ROSTERREAD-231

- `R0423` reads each roster CPlayer's `+0x04`, `+0x08` and `+0x30` and passes the address
  of `+0x10`, the name.
- `R0128` passes each player's `+0x34` to `R0130` with an index below 16 and copies the
  returned byte into the session table at `[L00004]+0xa9c4`.
- The destructor deletes each player through its virtual `+4`.
- No holder reads `+0x0c`. Its one read is in copy constructor `R0503`, whose two call sites
  are in `R0504` (`ALM-SCALAR-087`); `R0504` has no direct call and no dword reference
  (`ALM-MAPLIFE-226`, `ALM-COPY-151`).

**Confidence.** **High** for the holders' reads; the destructor's virtual `+4` (`R0528`) was
not read, and "no holder reads `+0x0c`" leans on the four access sites of `ALM-SCALAR-087` for
it / **Unknown** for indirect entry into `R0504` and for any effect of the scalar
(`ALM-SCALAR-089`).

**Unknown.** A computed call into `R0504`, and a reader of a copied CPlayer.

### ALM-NODEREAD-232

Loader case 7 allocates each node as `0x31c` bytes, runs `R2209` and reads, in order:
`0x40` bytes to `+0x000`, 4 to `+0x040`, 4 to `+0x044`, 4 to `+0x318`, `0x28` to `+0x048`,
`0x28` to `+0x070` and `0x280` to `+0x098`. So file `+0x48` lands at runtime `+0x318`, the ten
values at `+0x48`, the ten types at `+0x70` and the parameter names at `+0x98`. Actions are
appended to `M+0x40`, conditions to `M+0x5c` (`R2186`).

- Binder `R0067` (`TRIG-BIND-010`) reads M only at `+0x40`, `+0x5c` and `+0x78`. Action slot `[ebp-0xa4]` and
  condition slot `[ebp-0xa0]` dereference runtime `+0x40` (opcode), `+0x44` (id; its address is
  the key passed to map lookup `R2053`), `+0x48`..`+0x6f` (values) and `+0x70`..`+0x97`
  (types).
- Runtime `+0x318` is read for conditions only (`L14632`) and goes to the compiled check's
  `+0x48`.
- The label `+0x000` is read only as the `%s` of `"Failed to create instant %d %s"`
  (`L14633`) and `"Failed to create check %d %s"` (`L14634`).
- No holder reads the parameter names `+0x098`..`+0x317`.
- Node constructor `R2209` zeroes the label's first byte, `+0x40`, the values, the types,
  each name's first byte and `+0x318`. The node copies `L14635` and `L14636` and the
  node-list serializer `L13861` have no direct call; `L13861` is referenced only from a
  vtable slot at `L13862`.

**Confidence.** **High** for the loader layout and the binder's reads, taking the appended
pointer as what the accessors `R2189` and `L14637` return (appends `R2186`; the walker's
in-memory record is `ALM-CLS-037`) / **Medium** for no other reader: the map lookups `R2053`
and `L14625`, node delete `L14638` and the serializer reached through `L13862` were not
read. The binder's failure strings and the `0x1388` threshold are `TRIG-BIND-010`.

**Unknown.** Whether any code dispatches the node list's serializer slot.

### ALM-TRIGREAD-233

- Loader case 7 sizes the trigger array at `M+0x78` (`L14639`) and reads `0xb8` bytes into
  each element (`L02405`, `ALM-TRIGSTORAGE-163`).
- Binder slot `[ebp-0x230]` dereferences `+0x80`, `+0x84` (condition id pairs), `+0x98`
  (action ids, compared with `0x1388`), `+0xa8` (comparison codes) and `+0xb4` (once flag,
  `TRIG-FIRE-007`). The read order and the once flag in the pattern are `TRIG-REC-011`; the
  comparison codes are `TRIG-CMP-006`; the binder's strings and threshold are `TRIG-BIND-010`.
- No holder reads the name `+0x00`..`+0x3f` or `+0x40`..`+0x7f`.
- The trigger array's serializer `L14640` writes whole elements through `L14641`; it has no
  direct call and is referenced only from a vtable slot at `L14642`.

**Confidence.** **High** for the binder's reads and the holder set, taking the appended pointer
as what the accessors return / **Medium** for no other reader: the serializer slot's callers were
not found and the element accessors (`R2189`, `L14637`) were not read.

**Unknown.** The role of `+0x40`..`+0x7f` stays Unknown (`ALM-TRIGSTORAGE-163`).

### ALM-EDITDISP-277

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`, `input-hash.txt`), static: the two menu resources 128 and 129 (`r02`), the message-map tables around the hits of two dword scans of `.rdata` (`s01`, `s02`; `w02`, `w03`), the thunk table `L08454`..`L15497` read at a 5-byte stride (`w04`), the tool selectors `L15498`..`L15499` (`w07`), the mouse handlers' heads (`w05`, `w06`).

1. Menu 129 (the document menu; strings are program text) holds, among others, `Objects` 0x8027, `Select objects` 0x8029, `Structures` 0x8044, `Select structures` 0x8045, `Units` 0x804e, `Select units` 0x804f, `Players` 0x8042, `Map Information` 0x8056, `Global description and settings` 0x8057, `Locate Unit/Structure` 0x8062. Neither menu holds an item for a loot record, a script node or a trigger.
2. Command entries are six dwords (`0x111`, 0, id, id, `0x0c`, handler) and the handler is a `JMP` thunk: 0x8027 to `L15498`, 0x8044 to `L15500`, 0x804e to `L15501`, 0x804f to `L15502`, 0x8029 to `L15503`, 0x8045 to `L15504`, 0x8042 to `L02360`, 0x8056 to `L15505`, 0x8057 to `L15506`, 0x8062 to `L15507`.
3. `L15498`, `L15500`, `L15501` and `L15508` store 0x8028, 0x8043, 0x804a and 0x8053 into the word `[L14535]` (the tool word), after calling `L15509(new, old)`, and call `L15510`; `L15498` and `L15501` also call `L15511`, `L15512`, `L15513` and `L15514`, and `L15508` calls `L15515`. These callees (`L15509`, `L15515`, `L15510`, `L15511`, `L15512`, `L15513`, `L15514`) are unread. The four bodies themselves make no object; what the callees do is Unknown. The menu IDs 0x8027, 0x8044 and 0x804e therefore select the tool words 0x8028 (objects), 0x8043 (structures) and 0x804a (units).
4. Mouse entries (messages 0x200, 0x204, 0x205, 0x201, 0x202) follow the command entries in one table run and point through thunks to `L15516`, `L15517`, `L02337`, `L14537` (button down) and `L15518` (button up). The button-up body branches on `[L14535]`: 0x8032 to `L15519`, 0x8043 to `L15520`, 0x804a to `L15521`, 0x8053 to `L15522`, 0xe125 to `L15523`, other values through a table of 0x22 entries or to `L15524`. The button-down body indexes a byte table by the tool word minus 0x8008; for 0x8028, 0x8043 and 0x804a it takes the target `L15525`, which sets `[L15526]` to 1 and stores `L15129(E, cx, cy)` in `[L15527]`.

**Confidence.** High for items 1 to 3 (tables and instructions read) as to the stores and calls in the four bodies; Medium for any statement that nothing else happens on these commands (the seven callees are unread). Medium for item 4's statement that the mouse entries and the command entries belong to one map of one class (the run is contiguous; the map header was not read). Medium for the thunk table: `w04` lists 84 lines that are not `E9`; the thunks used here were read.

**Unknown.** The toolbar, the accelerator resource (128, 424 bytes, unread) and any command source other than the two menus; the dialogs' buttons; where the dialogs with the captions `Instance Manager`, `Check manager` and `Traps Manager` (program text, resources 272, 277, 282) are opened; a `.data` scan for message maps (only `.rdata` was scanned, as dword scans for the IDs and the mouse messages).

**Amended.** The Unknown on where the dialogs `Instance Manager`, `Check manager` and `Traps Manager` are opened is narrowed by `ALM-EDITNODE-280` (two `rel32` callers located, their routines not identified) and `ALM-EDITDISP-279` item 3 (`Traps Manager` is reached from tool 0x8053). Item 3's reading of the Objects tool word 0x8028 as a placement tool is narrowed by `ALM-EDITDISP-279` item 1 (its button-up case writes one byte per cell).

### ALM-EDITUNIT-276

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`): the button-up body `L15518` from its head and its tool word 0x804a case `L15521`..`L15522` (`w06`, `w08`, `w09`, `w10`), the add routine `L15528` (`w11`), the thunk table (`w04`), a dword scan of `.text` for `L15495` (`s03`, 105 hits). `E` is the map object.

1. Preconditions. The case runs when `[L15526]` is non-zero (set by the button-down). The cell `cx`, `cy` is the event position divided by 32 (toward zero) plus the view's `+0x54`, and `L15529(x, y)` plus the view's `+0x58` (`L15530`, unread). When bit 4 of the key-flag argument is set the case calls `L15351(cx, cy)` (unread) and ends. Otherwise `key = L15129(E, cx, cy)` (`L15130`, unread) is looked up in the hash at `E+0x304` (bucket `key >> 4` modulo `[hash+8]`, node word `+8` the key, node `+0xc` the object). With no object the case places a unit; with an object it places only when the class test `L14492(obj, L15216)` holds (the class of `ALM-SAVE-271`); otherwise it ends. The two paths run the same stores.
2. Template. The unit is made from the global object `T = [L15495]`: `T+0x12 = E[0x2c][E[0x328] - 1]` (the player at the current-player index `E+0x328`), `T+6 = cx*256+0x80`, `T+0xa = cy*256+0x80`, `E+0x90 = max(E+0x90, T+0x7a)` (unsigned compare), when `[L15531]` is non-zero `T+0x66 = rand * 15 / 0x7fff` (`L14578`, unread; a C library `rand` is an inference), and `T+4` (word) `= L15424(E)` (`L15422`: the smallest id from 1 that no object holds, `ALM-SAVE-269` item 4; not read here). No range test, refusal or clamp of a field value occurs in these bodies; the values in `T` come from code outside them.
3. Loot clone. When `T+0xd6` is not -1, a 0x4c-byte loot record is built by `L15532` (unread) as a copy of `E[0x2e0][T+0xd6]` and appended to the list at `E+0x2dc` (the array grow `L15533`, unread); `T+0xd6` becomes its index. For each entry j of the copy's dword array at `+0x18` (count at `+0x1c`) that is non-zero, a 0x2e-byte object is built by `L15534` (unread) as a copy of the type-9 entry `E[0x2f4][entry - 1]` and appended to the list at `E+0x2f0` (grow `L15535`, unread); the entry is replaced by the new list count plus one. The copy's `+0x3c` is set to the word `T+4`.
4. Add. `L15536` is `L15528(E, T)`. It sets `T+0x22 = T+6`, `T+0x26 = T+0xa`, `T+0x6a = 0`, `T+0xe = 0`, byte `T+0x7e = 0` and `T+0x6e = 0`, allocates 0x175 bytes (the size word of class record `L14483`, `ALM-SAVE-261`) and builds the object by `L15537(T)` (`L02322`, unread). The hash key is `L15322(E)` (`ALM-SAVE-269` item 4: the first free key from 1), not the word `T+4`; the node (chunks from `L15538`, unread) takes the key at `+8` and the object at `+0xc`. A second routine at `L02321` with the same head takes the key from `T+4` (its caller was not found).
5. After the add the case calls `L15539` on `L15540` (`L15541`, unread), the view's virtual slot `+0x5c` with 1, `L14556(0, 0, 0)` and sets `[L15526]` to 0.

**Confidence.** High for the stores, branches and call sequence in the windows read (items 1 to 5), for the EN program. Medium that the object built by `L15537(T)` takes its fields (and the wire fields later written) from `T`: the constructor is unread. Medium for the id and key statements (they rest on `ALM-SAVE-269` for `L15422` and `L15421`). The construction of `T` (its source for class, flags, skills, and the dialog values) is Unknown.

**Unknown.** The constructor `L02322`; how `T` is filled (`s03` lists the stores: `L15542`..`L15543`, `L15544`..`L15545`, `L15546`..`L15547`); the copies `L15532`, `L15534`; `L15130`; the ctrl-click action `L15351`; clone and delete; the unit dialog; loot, script and trigger commands (Q2 to Q7 of the preregistration, except the unit placement).

**Amended.** The Unknown on how `T` is filled is narrowed by `ALM-EDITUNIT-278`: the dialog writes `T+0x12`, `T+0x1a`, `T+0x1e`, `T+0x66`, `T+0x7a`, `T+0x155` and `T+0x159` on its OK and notification routines, and a right-click on a unit loads `T` from a copy of that unit; the first construction of `T` and the constructor `L02322` remain unread.

**Amended.** EXP-0562 reads the ctrl-click action `L15351`: it is the unit delete (`ALM-EDITDEL-299`, `ALM-EDITLINK-298`), also reached from the Delete command 0x8063 through `L15713`. No handler read copies an existing unit with its loot or type-9 entries (`ALM-EDITCLONE-301`, Medium: most handlers of the accelerator table are unread). The Unknown on delete is answered for the unit and structure classes; the Unknown on clone stays open as that Medium negative.

### ALM-EDITUNIT-278

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`), static: the button-up routine `L02337` from its head to the end of its units case (`w12`, `w13`; `w12` stops decoding at `L15548`, ten bytes before `w13` starts at `L15549`; the gap holds the instruction the branch at `L15550`/`L15551` jumps to; its ten bytes are not shown, and reading it as a store of 0 by analogy with the identical branch at `L15552` is an inference), the Select units command `L15502` (`w22`), the unit dialog class `L15553` with its handlers `L15554`, `L15555`, `L15556`, `L15557` (`w14`, the jump table after `L15558` and the padding after `L15559` are data decoded as instructions), the thunk table (`w04`). `E` is the map object, `T` the global template `[L15495]`.

1. Dispatch. The routine `L02337` is called the right-button-up routine from the message number 0x205 of its message map entry (the Windows number is an inference, as in item 8). It clears `[L14539]` and `[L15560]` and branches on the tool word `[L14535]`: 0x8043 to `L15561`, 0x8009 to `L15562`, 0x8028 to `L15563`, 0x8032 to `L15564`, 0x804a to `L15565`, 0x8053 to `L15566`, 0xe125 to `L15567`, other values to `L15568`.
2. The units case (`L15565`) builds a dialog object on the stack with `L15569` (a thunk to `L15553`; its constructor passes the template word 0xe3, resource 227, whose controls are unread: the DIALOGEX parser of the first pass failed), stores `E+0x28` and `E+0x328` into it, and looks up the cell key `L15129(E, cx, cy)` in the hash at `E+0x304` as in `ALM-EDITUNIT-276`.
3. Cell holds an object and the class test `L14492(obj, L14483)` (the unit class of `ALM-SAVE-261`) holds: the old `T` is destroyed through its virtual slot `+4` with 1; `T` becomes a new 0x175-byte object built by `L15537(unit)`, a copy of the clicked unit; the loop over `E[0x2c]` (count `E+0x30`) finds the player whose `+4` equals the one of `T+0x12` and stores its index plus 1 into the dialog; the dialog is run (`L15570`; that it is a modal run and that its result 1 means OK are inferences from the MFC-shaped call and the compare of its result with 1). On result 1: the hash node of the cell key is unlinked and put on the hash's free list (`[hash+0x10]`), the bucket array is freed when the count reaches 0, the old unit object is destroyed through its virtual slot `+4`, a node with the same key and the object `T` (the edited copy) is inserted, `E+0x90 = max(E+0x90, T+0x7a)` (unsigned), `E+0x328` takes the dialog's player index, and `T` becomes another copy `L15537(T)`. On any other result nothing else is stored, so the clicked unit stays and `T` remains a copy of it. The unit's key and `T+4` (its id) are not written by this routine. A cell object of another class skips the dialog.
4. Free cell (reached by the branch at `L15550`/`L15551`, whose target instruction in the ten bytes at `L15548` is not shown; a store of 0 is inferred from `L15552`): `T+4 = L15424(E)` (the smallest unused id, `ALM-SAVE-269`), the dialog runs, and on result 1 `E+0x328` takes the index plus 1 of the player whose `+4` equals the one of `T+0x12`. No object is added and `E+0x90` is unchanged; `T` keeps the dialog values and feeds the next placement of `ALM-EDITUNIT-276`.
5. The command Select units (0x804f, `L15502`) builds the same dialog class, stores `T+0x12 = E[0x2c][E+0x328 - 1]` before it runs, and on result 1 sets `E+0x328` from the player found; it adds and changes no object.
6. The routine `L15555`, which the thunk `L15496` names (who calls the thunk, the dialog class's vtable and message map, is unread; the body ends in `L15571`/`L15572` with no base-class OK call visible; that it runs on OK is an inference), reads the controls (`L15573`, the call the manager dialogs also use) and stores into `T`: `T+0x1a` from `492984(this+0x334, sel)` where `sel` is the result of message 0x100c (flags 2) on the control at `this+0x350`, minus-one results replaced by 0; `T+0x7a = this+0x370`; `T+0x1e = this+0x378`, then for `T+0x1a >= 0x40` a value 0 becomes 1 and a value above 4 becomes 4 (the type test at `L15574`, `JL`, is a signed compare; the upper bound at `L15575`, `JBE`, is unsigned, and the two agree on the spin range 0..0xfe); `T+0x159` from messages 0x147 then 0x150 on the control at `this+0x80`; `[L15531] = this+0x37c`; in `T+0x155`, bit 2 is set when `this+0x384` is non-zero and cleared otherwise, bit 0x80 is set when `this+0x38c` is zero and cleared otherwise, bit 1 is set when `this+0x388` is zero and cleared otherwise; `T+0x66` from eight controls tested with message 0xf0: `this+0x314` gives 0, `+0x170` 2, `+0x1e8` 4, `+0x260` 6, `+0x29c` 8, `+0x2d8` 0xa, `+0x224` 0xc, `+0x1ac` 0xe, the later test in that order winning. The messages are the standard Windows numbers (inference); the control classes are unread.
7. The routine `L15556` (called on a selection notification) stores `T+0x1a` and `T+0x66` by the same tests; `L15557` stores `T+0x12` as the element of the array at `[this+0x5c]+4` selected by message 0x147 on the control at `this+0x134`, when the selection is not negative. The id shown in the dialog (`this+0x374`, `T+4`) is read by the initialiser and written by no routine read here.
8. Ranges. The data exchange routine `L15576` (a MFC-shaped sequence, bodies `L15577`, `L15578`, `L15579`, `L15580`, `L15581` unread) binds control 0x451 to `this+0x370`, 0x453 to `0x374`, 0x454 to `0x378`, 0x4b1 to `0x37c`, 0x4b3 to `0x380`, 0x4b4 to `0x384`, 0x4c7 to `0x388`, 0x4c6 to `0x38c`, and calls `L15578` on `this+0x378` with the bounds 0 and 0xfe. The initialiser `L15554` sends message 0x465 to the controls at `this+0xf8` with the packed pair (0, 0xfe) and at `this+0xbc` with (0, 0x3e80), fills the list at `this+0x350` from the table `[L15308]` (count `[L15582]`, entry `+4` the code compared with `T+0x1a`, entry `+0x10c` the text, null entries skipped), the combo at `this+0x80` from the arrays `[L15583]` and `[L15584]` (count `[L15585]`; the entry equal to `T+0x159` is selected) and the player combo at `this+0x134` from the array at `[this+0x5c]+4` (entry `+0x10` the text). Program text and table contents are unread.

**Confidence.** High for the stores, branches and call sequence inside the listed bodies (items 1 to 7), for the EN program. Medium for every statement that composes them with an inferred link: that the entry `L02337` is the right-button-up routine (message number 0x205), that `L15555` is the OK routine (identified only through the thunk `L15496`) and that `L15570` returns 1 for OK, and the free-cell branch (the ten bytes at `L15548` are not shown). So "on OK the edited copy replaces the unit under the same key and id", "on a free cell the dialog edits the template" and "the OK routine writes seven fields" are Medium. Medium for item 8's ranges (the exchange routines and the tables are unread; the message numbers are Windows constants, an inference). Medium that the dialog is the only writer of these `T` fields.

**Unknown.** The dialog class's vtable and message map (who calls thunk `L15496`); the ten bytes at `L15548`; resource 227 (controls, captions); the bodies of the exchange routines and of `L15537` (`L02322`); the table `[L15308]`; how `T` is first built (`ALM-EDITUNIT-276`); the meaning of `T+0x1e`, `T+0x7a`, `T+0x66`, `T+0x155`, `T+0x159`; whether the dialog offers a delete (no deletion of a unit occurs in the units case, the dialog routines read or the Select units command; the other window of the map view, `L15568`, is unread); a clone command (none located).

**Amended.** EXP-0562 reads `L15568`: the epilogue of `L02337` (`CALL L15714`, `RET 0xc`), not a command (`w13`). A unit delete exists: command 0x8063 runs `L15713`, then `L15351` for a unit (`ALM-EDITDEL-299`). Item 1's 0x8053 case is read (`ALM-EDITCAST-294`, Medium for the right-click entry resting on item 1).

**Amended.** `ALM-EDITUNIT-304` reads item 8's list fill: every non-null class, lParam `+4`, so `T+0x1a` is the class `ID`.

### ALM-EDITDISP-279

Population: the EN `Map Editor.exe` (SHA-256 as in `ALM-EDITDISP-277`), static: the right-button-up routine `L02337` (`w12`, `w21`), the tool 0x8028 button-up case `L15586` and the jump tables that reach it (`w23`, `w06`), the Select objects command `L15503` (`w24`), the Edit commands and the message map (`w15`, `w19`, `w20`, `s04`), the copy and cut routines `L15587`, `L15588` and the shared body `L15589` (`w16`, `w17`), the cell byte functions `L14592`..`L15590` (`w18`), the tool 0x8053 cases (`w21`) with the scan `s05` and the call scan `s07`.

1. Tool 0x8028 (the Objects menu item, program text `Objects`): the button-up case `L15586` runs when `[L15526]` is non-zero. Its jump-table index is the tool word minus 0x8007 (byte table `L15591`, target table `L15592`: byte 4 gives `L15586`). It computes the cell as in `ALM-EDITUNIT-276` and calls `L15593(E, cx, cy, v)` with `v = byte [L15594]`, or 0 when bit 4 of the key flags is set. `L15593` is `L15595`: it stores the byte `v` at `[E+0x10] + cy*[E] + cx`. The case reads and writes no loot record, no hash `E+0x304` and no list `E+0x2dc` or `E+0x2f0`. The grid at `E+0x10` is a byte per cell; its meaning is unread.
2. The right-button-up case of tool 0x8028 (`L15563`) calls `L15596`, a thunk to `L15503`, the Select objects command (menu 0x8029): it builds a dialog by `L15597`, runs it by `L15570`, and destroys it; the result is not read.
3. Tool 0x8053 (set by `L15508`; no menu item names it in menu 129): the right-button-up case `L15566` looks up the cell key in the hash at `E+0x304` and dispatches by class: `L14483` (units) builds a dialog by `L15598` (a thunk to `L15599`), `L14482` builds one by `L15600` (`L15601`), any other object builds one by `L15602` (`L15603`); each receives the cell coordinates, the address `E+0x2f0` and (for units) the id word at `+4` of the object, and is run by `L15570`. The constructors push the template words 0x11b, 0x11c and 0x11a (dword scan `s05`, hits at constructor entry plus 0x23, the offset of the template word in the constructors `L15553`, `L15604`, `L15605`), resources 283, 284 and 282, whose captions are program text `Unit Spellbook`, `Cast from structure`, `Traps Manager`. The left-button-up case `L15522` runs a dialog built by `L15606` (`L15607`, unread) for the object found in the cell.
4. Edit commands: the message map entries of `0xe122` (copy), `0xe123` (cut), `0xe125` (paste), `0xe12b` (undo) and `0x802e` (select) point through thunks to `L15587`, `L15588`, `L15608`, `L15609`, `L15610`. Copy and cut both call `L15589(archive, cut)` after `L15611`; for each cell of the selection rectangle (the object at `ebp-0xa8` fields `+0x70`..`+0x7c`) it writes or reads the value `L15612` returns (set by `L14567`), altitude values (`L15132`, set by `L15133`) and the byte of item 1, gated by the three words `[L15613]`, `[L15614]`, `[L15615]` (the menu items Copy Tiles, Copy Altitudes, Copy Objects, 0x805b..0x805d, are program text). With the cut flag, each cell is set from `L15612(0, 0)`, from `L15132(0, 0)` and to the byte 0. The body reads no displacement `0x304`, `0x2dc` or `0x2f0`.
5. Keys. The map run `L15616`..`L15617` holds command entries, the messages 0x114, 0x115, 0x200, 0x201, 0x202, 0x204, 0x205, 0x5 and no entry for the messages 0x100, 0x101 or 0x102 (tables `w15`, `w20`, `w03`, `w02`). The accelerator resource (id 128, 424 bytes) is unread, so a key bound to a delete command is not excluded.
6. Unit deletion. In the windows read (the units cases of the two mouse routines, the dialog class, Select units, Select objects, copy, cut) no body unlinks a unit from the hash at `E+0x304` except the replacement in `ALM-EDITUNIT-278` item 3. No clone command was located.

**Confidence.** High for items 1, 2 and 4 for the windows read, and for the dispatch targets of items 3 and 4. Medium for the template ids of item 3 (taken from a dword scan and the constructor shape, not from reading `L15599`, `L15601`, `L15603`), and for item 5 (the accelerator table and the frame's own map are unread). Item 6 is a bounded negative: the population is the bodies listed above, the instrument a linear decode with `rel32` scans for the handler thunks only, so the absence of a delete elsewhere is Unknown, not shown.

**Unknown.** The accelerator table; the frame and document message maps; `L15568`; whether the byte grid is the sack layer; whether the entries of the list `E+0x2f0` that the `Unit Spellbook`, `Cast from structure` and `Traps Manager` dialogs receive are the entries `ALM-SACK-065` names loot elements. The captions (program text) suggest spells and traps, while `ALM-EDITUNIT-276` item 3 shows that the entries of a loot record's array at `+0x18` index that same list (the clone reads `E[0x2f4][entry - 1]` and appends to `E+0x2f0`), which agrees with `ALM-SACK-065` (case 9 appends to `map+0x2f0`) and does not contradict it; the open question is a caption against a list the cloning code treats as loot elements, and the dialog bodies are unread. Also unread and outside the population: the Paste command (message map entry 0xe125, handler `L15608`) and the tool 0xe125 button-up case `L15523`, the Select objects dialog built by `L15597` (its result is unread), the tool 0x8053 left-button-up dialog `L15607`, the toolbar. The `Copy Objects` gate of the copy and cut body writes only the byte of tool 0x8028, which excludes loot for copy and cut and says nothing about paste.

**Amended.** EXP-0562 reads what item 5, item 6 and the Unknown list left open. Item 3: only an empty cell builds the dialog by `L15602`; any other class builds none (`ALM-EDITCAST-294`). Item 5: the accelerator table 128 is read (53 entries; the Delete key is command 0x8063, Shift+Insert is 0xe125; `ALM-EDITDEL-299`, `ALM-EDITPASTE-300`). Item 6: a unit delete exists, reached by the Delete key and by a mouse case with bit 4 of the key flags (`ALM-EDITDEL-299`), and it removes the type-9 entries its loot records link and detaches the records from the list (`ALM-EDITLINK-298`); no command that copies an existing unit was found (`ALM-EDITCLONE-301`, Medium). `L15568` is the epilogue of `L02337`, not a command. The template ids of item 3 are read from the three constructors (`ALM-EDITCAST-295`..`297`). The entries of `E+0x2f0` that the three dialogs make carry kinds in word `+0x10`: 0..3 traps, 4 structure casts and 8 unit spells. The row text "no loot path or unit delete in the bodies read" holds for the bodies read by EXP-0560 only.

**Amended.** The Unknown on the Select objects dialog built by `L15597` is narrowed by `ALM-EDITOBJ-302`: the dialog lists the objects classes with a nonzero `InMapEditor` field and sets the Objects tool byte `[L15594]` that item 1 writes into the cell.

### ALM-EDITNODE-280


Population: the EN `Map Editor.exe` (SHA-256 as in `ALM-EDITDISP-277`), static: the DIALOG resources 272, 277, 282, 283, 284 (`r03`, program text), the dword scans `s05` and `s08`, the immediate scan `s06`, the `rel32` call scan `s07` (`.text`, four values), the constructor heads `L15618`..`L15619` and `L15620`..`L15621` (`w25`, mostly padding before the constructors), the callers `L15622`..`R1873` and `L15623`..`L15624` (`w26`), the message maps `L15625`..`L15626` and `L15627`..`L15628` (`w27`), the Instance manager handlers Add new `L15629` and Delete `L15630` (`w28`), the node copy routine `L15631` (`w29`), the Check manager handlers Add new `L15632` and Delete `L15633` (`w30`, `w31`).

1. Resources (program text). 272 `Instance Manager` and 277 `Check manager` hold the same control set: a list box (1109) captioned `Instance types` / `Check types`, a combo (1197), the lists `Defined Instances` / `Defined Checks`, the buttons `Add new` (1180), `Delete` (1181), `Validate Instances` / `Validate Checks` (1220), a check box `Execute once` (1138), five edit pairs 1115..1119 and 1127..1131 under `Parameters`, and `Done!`.
2. Classes. The constructor of the Check manager is `L15604` (template word 0x115 pushed at `L15634`; vtable `L15635`; a string slot at `this+0xd8` is filled from `L15636`). The constructor of the Instance manager is `L15605` (template word 0x110 at `L15637`; vtable `L15638`; string slot from `L15639`). Their thunks are `L15640` and `L15641`. The only `rel32` callers in `.text` (`s07`, four hits including the two thunks) are `L15642` and `L15643`, each in a routine whose frame holds a dialog object built there; the routines are not identified.
3. Handlers. The Instance manager map: control 1180 (code 0) to `L15629`, 1181 to `L15630`, 1220 to `L15644`; the Check manager map: 1180 to `L15632`, 1181 to `L15633`, 1220 to `L15645` (thunks resolved in `w04`).
4. Add new (both managers, bodies read: Instance `L15629`..`L15646`, Check `L15632`..`L15647`; the Check body was read by its calls, compares and the id lines `L15648`..`L15649`, and its string literals differ: `L15650`, `L15651` against `L15652`, `L15653`): the handler reads the controls (`L15573` with 1), then for each of the five rows reads two control texts and, over the ten slots of 64-byte strings at `this+0x190`, stores `atoi` (`L15654`) of the second text into the dword `this+0x140+4*j` for the slot whose string is neither of two literals and equals the first text (`L15655` compared with zero); `this+0xf8` is the working node. The new node id is `this+0x13c`, set before the copy to the tail node's `+0x44` plus 1 when the list at `this+0xdc` is non-empty and to 1 otherwise (`this+0x13c` is the working node's `+0x44`). A 0x31c-byte node is allocated and built by `L15656(this+0xf8)` (a thunk to `L15631`, the same routine in both managers), appended to the list at `this+0xdc` (head `+4`, tail `+8`, count `+0xc`; node of the list built by `L15657`), the list box at `this+0x78` is cleared (0x14b), refilled with message 0x143 per node and the node's `+0x40` as item data (0x151), and the last entry is selected (0x14e).
5. The node copy `L15631` copies: the string at `+0` (the string-assign shaped call `R0260`), the dword `+0x40`, the id `+0x44`, ten dwords at `+0x48`, ten dwords at `+0x70`, ten 64-byte strings at `+0x98` (ending at `+0x318`) and the dword `+0x318`: 0x31c bytes, the size of the node record of `ALM-TRIG-045` (its field names are not used here). The copied offsets agree with the runtime node layout of `ALM-TRIG-045` exactly: `+0x40` the opcode, `+0x44` the id, ten dwords at `+0x48` and `+0x70`, ten strings at `+0x98` and the dword `+0x318`. That supports that the nodes have the script node's shape; it does not show that the dialog's list is the map's script.
6. Delete (Instance `L15630`..`L15658`, Check `L15633`..`L15659`, both read): the selected index is read (0x147), the controls are read (`L15573` 1), the list at `this+0xdc` is walked and the node whose `+0x44` equals `this+0x13c` is unlinked, its object freed (`L15465`), the matching list box line is deleted (0x144), the selection is set to the previous line (0x14e) and the selection handler (`L11534` / `L15660`, targets `L15661` / `L15662`, unread) is called. No other node's id is changed and no other list, trigger slot or map object is read or written in these bodies.
7. The dialogs 282, 283 and 284 (`Traps Manager`, `Unit Spellbook`, `Cast from structure`) have no `Add new` or `Delete` button; the dialog 282 holds `Trigger Position`, `From`, `To` and `Spell Power` groups (program text). They are not the managers of trigger records.

**Confidence.** High for items 1, 3 and 5 and for the read bodies of items 4 and 6 (instructions read, EN), as operations on the dialog's own list at `this+0xdc`. Medium that this list is the map's script node list, or that a Delete removes a node from a saved map: the owner of `this+0xdc`, whether it is a copy filled by the callers `L15643` and `L15642`, and whether it is written back are unread. Medium for the id rule's use of the node field `+0x44` as the id that the script stores (the field's role is taken from the loop that compares it, not from a consumer), for the meaning of the two literals of item 4, and for "the only callers" in item 2: that is `rel32` calls only (scan `s07` over `.text`), with indirect calls, computed jumps and `.rdata` or `.data` pointers unscanned. Unit-id parameters, structure, group or player references, type codes and checks at parameter entry were not read.

**Unknown.** The owner of the list at `this+0xdc` (the map's script lists, or a copy) and whether it is written back; the two routines that call the managers (`L15643`, `L15642`) and the command that reaches them; the selection handlers (`L15661`, `L15662`), the Validate handlers (`L15644`, `L15645`), the type-combo handlers (`L15663`, `L15664`), the other edit handlers; the literals `L15652`, `L15653`, `L15650`, `L15651`; the opcode lists and defaults; how a working node's type code and ten codes are set; every question on triggers (no trigger dialog, handler or record write was located: `TRIG-*-113`..`116` are not used).

**Amended.** EXP-0562 locates the trigger dialog this card found none of: the Trigger Manager (resource 273, `TRIG-EDITTRIG-117`) edits the saved trigger array, and its Manage Instances and Manage Checks buttons open these two managers on the lists `E+0x40` and `E+0x5c` themselves, so the managers edit the saved node lists (Medium: the base-class paths were not read). Item 7 stays true for the dialogs 282, 283 and 284: their `<<` and `>>` buttons (1208, 1209) add and remove type-9 entries, not triggers (`ALM-EDITCAST-295`, `296`). The caller of the managers is the Trigger Manager (`L15715`, `L15716`), which settles the Unknown on `L15643` and `L15642`. EXP-0565 reads the type-combo, selection and Validate handlers (Check manager by outline) and partial catalogue paths: the observed type assignments use `Par<i>_value` text; the display call passes radix 10; write handlers pass text to conversion and store its returned 32-bit value. No type-8 comparison appears in these handlers. Instance manager instructions are High, Check manager and decimal-conversion semantics Medium; accepted syntax and range are Unknown. Second catalogue-loop completion past `L15946` including `L15947` is unread (`TRIG-EDITITEM-126`).

### ALM-EDITCAST-294

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`, 1,855,488 bytes), static. Read: the right-button-up case of tool 0x8053 `L15566`..`L15567` (`w12`), the thunk table `L08454`..`L15497` (`w04`), the `rel32` call scan `s06` and the dword scans `s07`, `s08`, `s09` (`.rdata`, `.data`, `.text`) over 16 values (the four dialog constructors and their thunks, the delete routines and the Paste handler).

1. In the case for tool word `[L14535]` = 0x8053 the routine looks up the object at the clicked cell (`L15129`) and tests its class with `L14492`. No object: the Traps Manager (constructor thunk `L15602`, body `L15603`; caller `L15717`). Class `L14483` (unit): the Unit Spellbook (`L15598`, `L15599`; caller `L15718`). Class `L14482` (structure): Cast from structure (`L15600`, `L15601`; caller `L15719`). Any other class: no dialog in this case. Each dialog receives the list `E+0x2f0`, the cell, and for the two object dialogs the object's id (`this+0x11c`, `+0x120`, `+0x124`, `+0x128`; the Traps Manager has `+0x13c`, `+0x140`, `+0x144` and a zero at `+0x148`). The window `w12` starts at `L15566`, inside the branch: the entry from the mouse routine `L02337` to this case on the tool word is read in `ALM-EDITUNIT-278` item 1, and the mapping of the routine to the right-button-up (message 0x205) is from `ALM-EDITDISP-277` item 4 and `ALM-EDITUNIT-278` item 1, where the Windows message number is an inference. Neither was read again here.
2. The tool 0x8053 case of the button-up routine `L15518` (`L15522`, `ALM-EDITDISP-277` item 4) opens the Trigger Manager (`TRIG-EDITTRIG-117`); that it is the left-button-up is inferred in the same way (Medium).
3. Callers (`s06`, `rel32` over `.text`, raw offsets): each of the four constructor thunks has one direct caller, listed above; no call targets a constructor body directly. The dword scans find no pointer to a constructor or its thunk in `.text`, `.rdata` or `.data` (`s07`, `s08`, `s09`).
4. This replaces item 3 of `ALM-EDITDISP-279`, which said every other object builds the trap dialog: only an empty cell builds it.

**Confidence.** **High** for items 1 and 3 as read inside the case: which dialog the case builds for no object, a unit and a structure (instructions read, EN; the scans bound the direct and pointer references). **Medium** that the case is entered by a right click and that the Trigger Manager case of item 2 is entered by a left click: the entry from the mouse routine and the message numbers are taken from `ALM-EDITUNIT-278` and `ALM-EDITDISP-277` (inferred), not from a window of this experiment. **Medium** that the class constants name the unit, structure and loot sack classes: they are read from the `L14492` test sites and from the delete routines, not from a class name. Indirect calls through registers and computed jumps are not scanned.

**Unknown.** The cell-object lookup `L15130` (behind `L15129`) was not read here; which object a cell holds when several share it; what a loot sack under the cursor does (the class tests read here are the two named).

### ALM-EDITCAST-295

Population: the EN `Map Editor.exe` as in `ALM-EDITCAST-294`, static. Read whole: the Traps Manager class `L15603`..`L15553` (`w27`, template word 0x11a, resource 282), its message map `L15720` (`w28`) and the resolved thunks (`w04`).

1. Map (`w28`): control 1208 (`<<`, clicked) to `L15721` to `L15722` (Add); 1209 (`>>`) to `L15723` to `L15724` (Remove); the list box 1206, selection change (code 1), to `L15725` to `L15726`. OK is `L15727` (thunk `L15728`), identified by its ending in the base OK `L15571`; its vtable slot was not read.
2. Fields (`L15729`): list box 1206; spell combo 1217 (None and 28 spells, `r02`); From radio 1210 Tile / 1211 Structure, edits 1154 x, 1212 y, 1213 structure id; To radio 1215 Tile / 1216 Unit, edits 1155 x, 1156 y; spell power edit 1218 (bounds 0..100 through `L15578`) with spin 1214; trigger position edits 1157 x, 1219 y (program text, `r01`).
3. Listing (`L15730`): for each entry of `E+0x2f0` with (`+8`, `+0xc`) not both 0 and word `+0x10` neither 8 nor 4, a line with the spell name (item data = the entry's list index). The entry at the clicked cell is loaded into the controls. Entries with kind 4 or 8 never appear here.
4. Add (`L15722`, combo selection must not be -1): a 0x2e-byte entry (constructor `L15206`) with word `+0x16` = the combo index, word `+0x18` = the power, `+8` and `+0xc` = the clicked cell (the dialog's variables for edits 1157 and 1219 are first overwritten with the cell, `L15731`..`L15732`), word `+0x10` = From + 2 x To (From 1 = Structure; To 1 = Unit), dword `+0x12` = 0 (the caller's zero), and two 6-byte elements in the array at `+0x1a` (data `+0x1e`, count `+0x22`): the From pair and the To pair as words. The third word of each element is copied from an unwritten local (`L15733`, `L15734`), so its content is not set by the handler. The entry is appended to `E+0x2f0`.
5. Selection change (`L15726`) and OK (`L15727`) write the previously selected entry back from the controls: spell, power, `+8` and `+0xc` from the trigger position edits, both element pairs, the kind word, and `+0x12` from the structure id edit 1213. So the typed structure id and the trigger position reach the entry on a later write, not on Add.
6. Remove (`L15724`): the selected entry's list slot is closed by `memmove` (`L15212`), the count drops by 1, and for every later entry with `+4` nonzero the loot back-reference is renumbered (`ALM-EDITLINK-298`). The list box is rebuilt.

**Confidence.** **High** for items 1, 3, 4 and 6 and for the field values of items 2 and 5 (instructions read, EN). **Medium** for the meaning of the kind word (`TRIG-*` consumers were not read here: the reading of From and To follows the dialog's radio groups, not a consumer) and for the control-to-field mapping of item 2, which comes from the exchange routine's control ids matched to the template. The edit bounds and spin ranges are Windows control messages read from constants (inferred).

**Unknown.** What the element third word holds in a saved map; the consumers of kinds 0..3 in the running game; whether the entry constructor `L15206` zeroes `+4`; the base OK path `L15571`. Add accepts the combo's first item None (index 0, `r02`) and `+8` is the clicked cell's x, so an entry with word `+0x16` = 0 can be made; with `+4` zero (constructor unread) it meets the save-time removal of `ALM-SAVE-254` item 2. Whether the running game accepts such an entry is not read.

### ALM-EDITCAST-296

Population: the EN `Map Editor.exe` as in `ALM-EDITCAST-294`, static. Read whole: the Unit Spellbook class `L15599`..`L15735` (`w29`, template word 0x11b, resource 283) and its message map `L15736` (`w30`).

1. Map (`w30`): 1208 clicked to `L15737` to `L15738` (Add); 1209 to `L15739` to `L15740` (Remove); the list box 1206, selection change, to `L15741` to `L15742`. OK is `L15743` (thunk `L15744`), identified by its ending in the base OK `L15571`; its vtable slot was not read.
2. The caller sets `this+0x11c` = `E+0x2f0`, `+0x120` x, `+0x124` y, `+0x128` the unit id. The unit id edit 1213 is read-only (style bit 0x800, `r01`) and is filled from `this+0x128` (`L15745`).
3. Listing (`L15745`): entries with (`+8`, `+0xc`) not both 0, word `+0x10` = 8 and dword `+0x12` = the unit id, with the spell name per line. An entry of kind 8 at cell (0,0) is not listed.
4. Add (`L15738`, combo selection not -1): a 0x2e-byte entry with word `+0x16` = the combo index, word `+0x18` = the power (edit 1218, spin 1214), `+8` = x, `+0xc` = y of the unit's cell, word `+0x10` = 8, dword `+0x12` = the unit id; appended to `E+0x2f0`. There is no test for an equal entry already present, so Add twice makes two entries.
5. OK (`L15743`) writes the selected line's entry from the controls the same way; selection change (`L15742`) does so for the previous line and loads the new one. Remove (`L15740`) closes the slot, lowers the count and renumbers later entries' loot back-references (`ALM-EDITLINK-298`); the list box is rebuilt.

**Confidence.** **High** for items 1 to 5 as read (instructions read, EN). **Medium** that dword `+0x12` equals the unit's `Unit` id as the running game uses it: the dialog stores the caller's word `+0x128`, taken from the object at the cell (`ALM-EDITCAST-294`), not a consumer's read.

**Unknown.** The unit's id field the caller passes (read in `w12` only as a word); the consumers of kind 8 in the running game. Add accepts the combo's first item None (index 0, `r02`), so an entry with word `+0x16` = 0 can be made; with `+4` zero (constructor `L15206` unread) it meets the save-time removal of `ALM-SAVE-254` item 2.

### ALM-EDITCAST-297

Population: the EN `Map Editor.exe` as in `ALM-EDITCAST-294`, static. Read whole: the Cast from structure class `L15601`..`L15603` (`w26`), its empty message map and vtable (`w38`, `w39`: no command entry; `OnInit` slot `L15746` to `L15747` to `L15748`; OK slot `L15749` to `L15750` to `L15751`).

1. Controls (template 284, `r01`): spell combo 1217, structure id edit 1213 (read-only), power edit 1218 with spin 1214, scan range edit 1219 with spin 1004. The caller sets `this+0x11c` = `E+0x2f0`, `+0x120` x, `+0x124` y, `+0x128` the structure id.
2. `OnInit` (`L15748`): spin ranges are set by control message (power 0..100 as low `0x64`, scan range low `0xf0`, high 1; Windows constants inferred). The list is searched for an entry at the clicked cell with word `+0x10` = 4 and dword `+0x12` = the structure id. Found: the controls get its spell (`+0x16`), power (`+0x18`) and range (dword `+0`). Not found: a new 0x2e-byte entry is appended at once, with word `+0x10` = 4, `+8`, `+0xc` = the cell, `+0x12` = the id, spell 0; the spell combo shows None.
3. OK (`L15751`): the combo index 0 (None) removes the entry (slot closed, count lowered, later loot back-references renumbered, `ALM-EDITLINK-298`). Any other index stores word `+0x16` = the index, word `+0x18` = the power, dword `+0` = the scan range.
4. So opening the dialog on a structure and pressing OK with spell None leaves no entry (items 2 and 3). No body read removes the entry `OnInit` appended when the dialog is cancelled, so Cancel is read to keep it; the Cancel path is the base class (`L15572`) and was not read.

**Confidence.** **High** for items 1 to 3 as read (instructions read, EN) and for the first sentence of item 4 (OK with spell None leaves no entry). **Medium** for the Cancel consequence of item 4: the base-class Cancel path `L15572` was not read. **Medium** that dword `+0` is the scan range: the field is named from the dialog's control 1219 (`Scan Range` group, program text), not from a consumer.

**Unknown.** The bounds the data exchange enforces on the edits (`L15752` family and `L15578` were not read); what the running game does with a kind-4 entry whose spell is 0. The entry appended on open has word `+0x16` = 0 until OK; with `+4` zero (constructor `L15206` unread) it meets the save-time removal of `ALM-SAVE-254` item 2, and the dialog's Cancel is read to keep it (Medium).

### ALM-EDITLINK-298

Population: the EN `Map Editor.exe` as in `ALM-EDITCAST-294`, static. Read: `L14524` to `L14474`..`L15753` (`w36`), the structure delete `L15350`, the unit delete `L15351` (`w16`), and the inline removals of the three caster dialogs (`w26`, `w27`, `w29`).

1. `L14474(index)` removes a type-9 entry of `E+0x2f0`: closes the slot by `memmove`, lowers the count at `E+0x2f8`, and for each later entry `k` with `+4` nonzero sets the dword `[entry+4]+0x18 + 4 x [entry+0]` to `k + 1`. The removed entry is not read or freed there. It is reached through `L14524` from the five rel32 callers of `ALM-SAVE-264` (`L14527`, `L14528`, `L14529`, `L14530`, and `L14526`, the save-time normaliser of `ALM-SAVE-254`); the caller `L14530` lies in `L15307`, which this experiment did not read.
2. The three dialogs use the same loop inline after their `memmove` (Cast from structure OK with spell None, Traps Remove, Unit Spellbook Remove). The removed entry's own link into its loot record is not cleared in any of the three.
3. `L15350` (structure) and `L15351` (unit) call `L14524` for each entry they remove and decrement their own loop index. `L15351` (`w16`, `L15754`..`L15755`), for each loot record whose owner word is the unit's id, calls `L15756(rec, 0, -1)` and `L15251(rec+0x28, 0, -1)` (callees unread), then walks the record's link array `[rec+0x18]` and removes each named entry through `L14524`, then calls `L15756(rec+0x14, 0, -1)`, then closes the record's slot in `E+0x2dc` with `memmove`; it then renumbers the surviving loot records' ordinal at unit `+0xd6`. No destructor or free of the record object appears in the body. So the entries are removed while the record is still listed and the record is then detached, and no listed loot record keeps a link to a removed entry; that the record or its member arrays are freed is an inference about unread callees (Medium).

**Confidence.** **High** for items 1 and 2 and for the order of calls in item 3 as read (instructions read, EN). **Medium** that these are all the type-9 removals the editor makes: no scan of this experiment targets `L14524`, `L14474` or a decrement of the count at `E+0x2f8`, the callers come from the rel32 scan of `ALM-SAVE-264` (its dword scan for pointer callers was not run), `L15307` is unread, and another inline removal elsewhere in the program is not excluded. **Medium** that the unit delete frees its loot records (callees unread). **Medium** that no stale link can result: it depends on the entries made by the caster dialogs holding `+4` = 0 (the entry constructor `L15206` was not read), and the loot record layout (`+0x18` array of 1-based ordinals) is taken from the loop, which agrees with `ALM-EFFLINK-072`.

**Unknown.** The callees `L15756` and `L15251`; the function `L15307`; a dword scan for pointer callers of `L14524` and `L14474`; the constructor `L15206`; `L14519` (behind `L14520`), which the unit delete uses to find a unit by id; the effect of a stale link in the running game.

### ALM-EDITDEL-299

Population: the EN `Map Editor.exe` as in `ALM-EDITCAST-294`, static. Read: the accelerator table 128 (`r03`, 53 entries), the message map entries around `L15757` (`w14`), the Delete handler `L15713` (`w15`), `L15350` and `L15351` (`w16`), the left-button-down handler `L14537`..`L15758` only at the call sites `L15759` and `L15760` (`w32`, `w33`).

1. The Delete key (virtual key 0x2e, flags 0x03) is command 0x8063 (`r03` entry 41); the message map entry `L15757` sends it to thunk `L15761`, body `L15713` (`s07`: the only pointer to the thunk is `L15762`).
2. `L15713` builds a 0x32c-byte snapshot object (`L15356` to `L02361`) for undo, takes the cell from `[L15763]` and `[L15764]`, looks up the object (`L15129`), and calls `L15350` for class `L14482` (structure), else `L15351`.
3. Structure delete `L15350`: every `E+0x2f0` entry with `+4` = 0, word `+0x10` = 4 and `+0x12` = the structure's id is removed (`ALM-EDITLINK-298`), then the object's destructor slot and its hash node. Unit delete `L15351` (class `L14483` must hold): for every loot record whose owner word (`+0x3c`) is the unit's id, each entry its link array names is removed and the record is detached from the list (`ALM-EDITLINK-298` item 3: the order read, no free seen); then the entries with `+4` = 0, word `+0x10` = 8 and `+0x12` = the unit id are removed; then the unit is destroyed and its hash node unlinked. A cell with another class changes nothing in these bodies.
4. Other callers (`s06`): `L15759` and `L15760` in the left-button-down handler `L14537` call the same two routines when bit 4 of the handler's flags argument is set (`ALM-EDITUNIT-276` calls it the ctrl-click); `L15765` in the structure dialog case of the right-button-up (`w37`) calls `L15350` after the dialog's OK when the word `[L15766]` is 0. `L15767` and `L15768` are the two calls of item 2.
5. Shift+Delete is command 0xe123 and Ctrl+Insert is 0xe122 (`r03` entries 42, 46), cut and copy of a rectangle (`ALM-EDITDISP-279`); neither is a unit delete.

**Confidence.** **High** for items 1 to 3 and 5 as read (instructions read, EN). **Medium** for item 4: the bodies of `L14537` and `L02337` around the call sites were read only to the extent shown, and the meaning of bit 4 and of `[L15766]` is not stated by a consumer.

**Unknown.** What a delete of a loot sack does (no class test reaches `L15216` in the bodies read); the object lookup behind `L15129`; the contents of the undo snapshot.

### ALM-EDITPASTE-300

Population: the EN `Map Editor.exe` as in `ALM-EDITCAST-294`, static. Read whole: the Paste command handler `L15608`..`L15769` (`w17`), the load branch `L15769`..`L15770` (`w18`), the paste mouse-up case `L15523` (`w24`), the three cell setters `L14557`, `L15595`, `L14551` (`w35`), the Select objects dialog constructor and OK (`w25`), and the vtable and message map entries around the Paste command (`w31`, `w21`, `w22`).

1. Command 0xe125 (Shift+Insert, Ctrl+V, key `V` with flags 0x0b; `r03`) is in the message map at `L15771` and goes to thunk `L15772`, body `L15608`: it saves `[L15773]` = `[L14535]` and sets the tool word to 0xe125. It writes no map data.
2. The tool-0xe125 mouse-up (`L15523`) restores the tool word, reads the clipboard through `OpenClipboard`, `GetClipboardData` and `CloseClipboard` (`w34`), wraps the data in a memory archive and calls `L15611` (the load branch `L15769`). With no data it reports through `L15774` and writes nothing.
3. The load branch reads a rectangle's width x height words and two byte buffers, then loops over the rectangle and, for each of three global switches (`[L15613]`, `[L15614]`, `[L15615]`), writes one cell layer: a word through `L14551` (array at object `+0xc`), a byte through `L14557` (object `+0x14`) and a byte through `L15595` (object `+0x10`); the object-byte layer is written for a nonzero source byte unless `[L15775]` is set. The three cell layers are gated by three global switches. The switches match the Copy Tiles, Copy Altitudes and Copy Objects menu item texts (program text; `ALM-EDITDISP-279` item 4, which rests on the same text match); the menu handlers that set `[L15613]`, `[L15614]` and `[L15615]` were not read, so naming the layers tiles, altitudes and the object byte is Medium. No call touches `E+0x2f0`, `E+0x2dc`, `E+0x78`, the node lists or the object hash `E+0x304`. A cut (second argument nonzero) clears the source rectangle with `L15612`, `L14567`, `L15132`, `L15133` and `L15593`.
4. The Select objects dialog (template 185, constructor `L15776` with word 0xb9): its OK (`L15777`) reads the list selection (message 0x100c with flags 2), stores the template index + 1 as a word in `[L15594]` and closes. Opening it creates no object.

**Confidence.** **High** for items 1 and 2, for item 3 as to three cell layers written under three global switches (instructions read, EN), and for item 4. **Medium** for the names of the layers (menu item text, program text; the menu handlers were not read). **Medium** that the Paste path creates or moves nothing else: callees `L14495` (called after the loop), `L14494` and the view methods were not read.

**Unknown.** The menu handlers that set the three switches; whether the object-byte layer is the sack layer (`ALM-EDITDISP-279`); `L14495`; whether a pasted rectangle can carry a unit or structure through the clipboard format (the archive's writer is the copy path of `ALM-EDITDISP-279`).

**Amended.** Item 4's template index is narrowed by `ALM-EDITOBJ-302`: the word stored in `[L15594]` is the selected item's lParam plus 1, and the lParam is the class's `+4` value, not the row number. Rows skip the classes whose `+0x48` is 0, so the two differ after the first skipped class.

### ALM-EDITCLONE-301

Population: the EN `Map Editor.exe` as in `ALM-EDITCAST-294`. Searched: the accelerator table 128 (53 entries, `r03`), the `.rdata` dword scan `s03` for eight of its command ids (`0x8063`, `0xe145`, `0xe150`, `0xe151`, `0x8033`, `0x8034`, `0x8013`, `0x8014`; 24 hits, message maps), the handlers read in this experiment (`0x8063` Delete, `0xe125` Paste, the five dialogs), and `ALM-EDITUNIT-276`, `ALM-EDITDISP-279` of EXP-0560 (copy and cut of a cell rectangle; unit made from the template `[L15495]` with a clone of its loot).

1. No handler read copies an existing unit, its loot records or its type-9 entries. The only loot copy located is the clone made when a unit is placed from the template (`ALM-EDITUNIT-276`).
2. Of the 53 accelerator entries, 0x8063, 0xe122, 0xe123 and 0xe125 are read (Delete, copy, cut, Paste); the 0x80xx tool selectors set the tool word (`ALM-EDITDISP-277`); the commands 0xe100, 0xe101, 0xe103, 0xe104, 0xe145, 0xe150, 0xe151 and the 0x80xx commands not listed there have handlers that were not read.

**Confidence.** **Medium**: the negative covers the handlers read and the accelerator table, and the unread handlers of item 2 could hold a clone. Not searched: menu items without an accelerator beyond the two menus of EXP-0560, and commands reached only by toolbar or dialog buttons.

**Unknown.** The handlers of the commands in item 2; a native editor run, or reading them, settles it.

### ALM-EDITOBJ-302

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`, `input-hash.txt`), static, capstone linear decode: the objects loader `L15800`..`L15801` (`w07`, `w02`, `w04`, `w08`), the Select objects initialiser `L15802`..`L15803` with the notification handler `L15804` (`w10`), the second handler `L15777` and the Objects tool cases `L15805`, `L15586` (`w20`), the message map `L15806`..`L15807` (`w11`), the list-control helpers `L04847`..`L15808` (`w29`); scans `s02` to `s07` and `s11`. Addresses are the editor's.

1. List. The initialiser `L15802` sends message 0x102f with the objects array size `[L15809]` and message 0x1003 with a value taken from `[L15810]` (0 when it is 0) to the list control at `this+0x78`, then walks the objects array `[L15811]` from index 0 to the size. A class whose field `+0x48` is 0 gets no row. Every other class gets one row through `L04847` with the text at class `+0x4c`, then `L15812` with mask 4 and the class's `+4` value. The row is marked (mask 0xb, state 3) when the word `[L15594]` minus 1 equals the class's `+4`.
2. Order. The loader appends each class pointer at the array's current size in the order of its loop over the `Object%d` sections (`REG-EDLOAD-121`). It stores the `InMapEditor` value at `+0x48` (`REG-INMAP-122`) and tests it nowhere. Rows are therefore inserted in section order, without the classes whose `+0x48` is 0. The list shows that order only if the control does not sort; its style is in dialog resource 185, which was not read, and the control class is inferred from the message numbers.
3. Row value. `L04847` and `L15812` fill one 0x28-byte item (mask at 0, item at 4, text at 0x14, image at 0x1c, lParam at 0x20) and send 0x1007 and 0x1006 to the control at `this+0x1c` through `[L15813]`; `L15814` sends 0x1005 with mask 4 for one item and returns the lParam slot. A row's lParam is the class's `+4` value.
4. Pick. `L15804` (message map entry `L15815`: message 0x4e, code 0xfffd, control 0x400, thunk `L15816`) and `L15777` (entry `L15817`, thunk `L15818`) take the selected item (message 0x100c with flags 2; a negative result becomes 0), read its lParam through `L15814`, add 1 and store the word `[L15594]`.
5. Cell. The Objects tool's mouse-up case `L15586` writes the byte `[L15594]` into the cell through `L15593` (`L15595`, `ALM-EDITDISP-279` item 1). The cell byte is the chosen class's `+4` value plus 1, low byte.
6. Shipped registry (composition). The loader's first key, stored at `+4`, is named `ID` in a byte table that was not committed (EXP-0563 Deviations), and `objects.reg`'s `ID` equals the section index on 82/82 classes (`REG-KEY-044`). With item 2's order, the written byte is the section index + 1, the code `ALM-CLS-035` resolves as section index `c - 1` in `rom.exe`. Of the preregistered alternatives, O1 (a key filters the rows) holds for the program read; O0 (every section is a row) and O2 (the list does not come from `objects.reg`) are excluded.

**Confidence.** **High** for items 1 to 5 as instructions read (EN): the `+0x48` test, the insertion order, the text offset, the lParam value, the +1 store and the cell write (item 5 at `ALM-EDITDISP-279`'s grade). **Medium** for the order the list shows (resource 185's list style unread) and for the loader's binding to `objects.reg` (`REG-EDLOAD-121`). **Medium** for item 6: the name `ID` of the `+4` key and the name `DescText` of the `+0x4c` text rest on an uncommitted byte table, and the section-index equality on `REG-KEY-044`'s measurement. The message numbers are Windows list-view and notification constants (inference); the import behind `[L15813]` was not read; that `L15777` runs on OK is `ALM-EDITPASTE-300`'s reading.

**Unknown.** Resource 185's controls; who fills the image list `[L15810]` (the loop at `L15819` of `REG-INMAP-122` is a candidate whose role is unread); `L14505` and `L15571` (base-class shaped calls).

### ALM-EDITSTRUCT-303

Population: the EN `Map Editor.exe` as in `ALM-EDITOBJ-302`: the structures loader `L15820`..`L15821` (`w25`) and its key strings (`w26`, `w31`), the Select structures command `L15504` (`w12`, `w13`), the dialog class (`w15`, `w17`, `w18`, `w19`: `L15822`..`L15823`, `w21`: `L15823`..`L15824`), the Structures tool case `L15520` (`w20`), the placement `L15825`..`L15826` (`w22`, `w28`), the object initialiser `L02334`..`L15827` (`w29`), the list-control helpers (`w29`); scans `s08` to `s11`.

1. List. The initialiser `L15822` sends 0x102f with the structures array size `[L15484]` to the list at `this+0xb4`, fills the combo at `this+0xf0` (message 0x143) with the texts at `+0x10` of the entries of the list at `this+0x11c` and selects `this+0x120` (0x14e), then walks `[L15131]` from index 0 to the size. A null slot, or a class whose `+0x88` is not 0, gets no row. Every other class gets a row with the text at `+0x68` and the lParam `+0xc`, marked when `[L15766] & 0xfff` equals `+0xc & 0xfff`.
2. Fields and order. The loader stores the key at string `L15489` (`ID`) at `+0xc`, `DescText` (`L15828`) as 32 bytes at `+0x68` and `VariableSize` (`L15829`, default 0) at `+0x88`, and stores each class at array index `+0xc` (`REG-EDLOAD-121`). Rows are inserted by ascending `+0xc` value; a class with a nonzero `VariableSize` gets no row. The list shows that order only if the control does not sort; its style is in dialog resource 226, which was not read, and the control class is inferred from the message numbers. No `InMapEditor` key is read for structures (`REG-INMAP-122`).
3. Pick. `L15830` (notification 0xfffd from control 0x413) and `L15831` (notification 0xfff8 from control 0x413) store the selected item's lParam (`L15814`) into the word `[L15766]`, with no +1. `L15832` (command from control 0x410) keeps the combo selection (message 0x147) at `this+0x120`.
4. Placement. The Structures tool case `L15520` takes `idx = [L15766] & 0xfff`, the class `[L15131][idx]` and its extents `+0x10`, `+0x14`, calls `L15129` for each cell of that rectangle and then, through `L15833`, `L15825(x, y, idx, 0x64, E+0x328)`.
5. Object. `L15825` reads `idx` as a sign-extended byte and makes no object unless x > 7, y > 7 and x + class `+0x10` and y + class `+0x14` stay below the map's `E+0` - 8 and `E+4` - 8 (unsigned). It builds a 0x106-byte object (`L15310`), calls `L14610(x, y, idx)` when `[L14608]` is nonzero, and calls `L02334` (thunk `L15309`), which stores `+0x1a` = idx, word `+4` = the id from `L15424(E)` (`ALM-SAVE-269`), `+6` and `+0x22` = x * 256 + 0x80, `+0xa` and `+0x26` = y * 256 + 0x80, `+0xe`, `+0x66`, `+0x6a` and word `+0xfe` = 0, words `+0xfa` and `+0xf6` = 0x64. With p = `E+0x328`, `+0x12` = `E[0x2c][p - 1]` when p <= `E+0x30`; otherwise `+0x12` = `E+0x3c`, a 0x48-byte object built once by `L15249(1, 1)` and appended to the array at `E+0x28`. The object is linked into the hash `E+0x304` under the key `L15322(E)`. The chosen row's `+0xc` value is the placed structure's `+0x1a`.

**Confidence.** **High** for items 1 to 5 as instructions read (EN), item 2 as the insertion order, and for the names `DescText` and `VariableSize` (`w26`). **Medium** for the order the list shows (resource 226's list style unread) and for the loader's binding to `structures.reg` (`REG-EDLOAD-121`). **Medium** for the name `ID` of the `+0xc` key: its string `L15489` was read as window 122, past the cap of 120. Medium for the notification and message meanings (Windows constants, inference). Unread: `L15310`, `L14610`, `L15130` (behind `L15129`, its role in the cell loop), `L15249`, `L15538`, and the store into `E+0x328` after the dialog in `L15504`.

**Unknown.** Whether `+0x1a` is what the save path writes as the type-4 `kind` (`ALM-CLS-036`): no structure writer was read. How many shipped structures carry a nonzero `VariableSize`: the corpus rules did not count that key. A class value above 0x7f would be sign-extended by `L15825`; the shipped `ID`s run 1..66 (`REG-KEY-044`). The combo's list at `this+0x11c` and its use.

### ALM-EDITUNIT-304

Population: the EN `Map Editor.exe` as in `ALM-EDITOBJ-302`: the unit dialog initialiser `L15554`..`L15558` (`w27`, `w28`), the units loader ranges `L15834`..`L15835`, `L15836`..`L15837` and `L15838` (`w30`, `w31`), the list-control helpers (`w29`); scan `s12`.

1. List. The initialiser sends 0x102f with the units array size `[L15582]` to the list at `this+0x350`, then walks `[L15308]` from index 0 to the size. A null slot gets no row. Every other class gets a row through `L04847` with the text at `+0x10c` and the image equal to the row number, then `L15812` with mask 4 and lParam `+4`; the row is marked (mask 0xb, state 3) when the template's `+0x1a` (`T = [L15495]`) equals `+4`. The row number advances only for an inserted row. No other class field is tested.
2. Order. The units loader stores the `ID` key (`L15839`) at `+4` and each class at array index `+4` (`L15840`, `REG-EDLOAD-121`). Rows are inserted by ascending `ID`, one per loaded class with a distinct `ID`. The list shows that order only if the control does not sort; its style is in dialog resource 227, which was not read, and the control class is inferred from the message numbers. No `InMapEditor` value is read (`REG-INMAP-122`).
3. Rest of the initialiser. It copies `T+0x7a`, the word `T+4` and `T+0x1e` into the dialog's `+0x370`, `+0x374`, `+0x378`, calls `L15573(0)`, selects item `[this+0x60] - 1` of the player combo, fills the combo at `this+0x80` from `[L15583]` with item data from `[L15584]` (count `[L15585]`) and selects the entry equal to `T+0x159`, calls `L15841(this+0x64, 0)` when `this+0x388` is 0, and returns 1.
4. Pick. `L15555` and `L15556` store `T+0x1a` from `L15814` of the selected item (`ALM-EDITUNIT-278` items 6 and 7); `L15814` returns the item's lParam (`ALM-EDITOBJ-302` item 3), so `T+0x1a` becomes the chosen class's `ID`. The unit tool places a copy of `T` (`ALM-EDITUNIT-276`).

**Confidence.** **High** for items 1 to 3 as instructions read (EN), item 2 as the insertion order, and for the lParam mechanism of item 4. **Medium** for the order the list shows (resource 227's list style unread) and that the units loader reads `units.reg` (its path push is unread, `REG-EDLOAD-121`). **Medium** that item 4's store runs on OK (`ALM-EDITUNIT-278`'s inference) and that the row text is `DescText`: the editor's key behind `+0x10c` was not read; the game's units class has `DescText` at `+0x10c` (`REG-UNITS-049`) and the editor's units class is also 0x12c bytes, a size match. Medium for the shipped row count: 34 classes with 34 distinct `ID`s (`REG-KEY-044`) give 34 rows.

**Unknown.** Resource 227's controls; the units loader between `L15835` and `L15836` (its other keys); `L15573`, `L15841`, `L15842`.

## Map editor: players and map-information dialogs

Population: the EN `Map Editor.exe` (SHA-256 `697f7b2cddd0c8cd18e8b567e83e70c68a81e3cee31b407fa5d3e6f6526bcb9e`), static
reading. The RU root has no editor program in the name listings of earlier experiments. Addresses are the editor's.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-DLGPLAYERS-286 | The EN editor's Players command (menu 0x8042) runs L02360, which opens the template-225 dialog; its radios, combo, Add, Delete and Diplomacy controls reach L15665, L15666, L15667, L02367, L15668 and L15669. | High / Medium | ● active | [EXP-0561](../experiments/EXP-0561-editor-players/) |
| ALM-DLGPLAYERS-287 | The Players dialog requests +0x08 in 0..15, +0x0c in 0..0x7fffffff and a name of at most 32 characters; OK stores +0x30, +0x0c, +0x08 and the name and sets +0x04 to index+1, validated or not (Medium). | High / Medium | ● active | [EXP-0561](../experiments/EXP-0561-editor-players/) |
| ALM-DLGPLAYERS-288 | Add appends a player (+0x30 1, +0x0c 5000, name New Player N) with no 16-player test; Delete drops a non-first player's pointer without destroying it and touches no owner; a deleted slot's objects rebind at reload (Medium). | High / Medium | ● active | [EXP-0561](../experiments/EXP-0561-editor-players/) |
| ALM-DLGPLAYERS-289 | The Diplomacy dialog (template 270) shows and writes the current player's 16 relation words only, through 0->1, 1->0, 2->2 and other->0 on both directions; it mirrors nothing into other rows and Cancel writes nothing. | High / Medium | ● active | [EXP-0561](../experiments/EXP-0561-editor-players/) |
| ALM-DLGINFO-290 | The Global description dialog (menu 0x8057) writes the map name at E+0x94, the text region at E+0xdc, E+0xd8 (1..9999) and E+0xd4 (1..50) and nothing else; strings have no program-side length bound. | High / Medium | ● active | [EXP-0561](../experiments/EXP-0561-editor-players/) |
| ALM-DLGINFO-291 | The Map Information dialog (menu 0x8056) is a display of counts the handler computes; its handler stores nothing into the map object and ignores the modal result. | Medium | ● active | [EXP-0561](../experiments/EXP-0561-editor-players/) |
| ALM-ROSTERIO-292 | The type-5 reader and writer copy +0x08, +0x30, +0x0c, the 32-byte name and exactly 16 relation words without normalising them; 357 installed records (38 EN, 33 RU maps) hold 0, 1 or 2 words only. | High / Medium | ● active | [EXP-0561](../experiments/EXP-0561-editor-players/) |
| ALM-STRCOPY-293 | Copy R0260 stores a string through its NUL and nothing after it, so a dialog edit or a new map leaves the bytes after a name's or text's NUL as they were; the reader's type-0 defaults run only when no type-0 record was read. | High / Medium | ● active | [EXP-0561](../experiments/EXP-0561-editor-players/) |

### ALM-DLGPLAYERS-286

Message-map entries `{0x111, 0, id, last id, sig 0xc, pfn}` with a `JMP rel32` thunk in code-region as the pfn.

1. Menu 129 (`&Players` 0x8042, `&Map Information` 0x8056, `Global des&cription and settings` 0x8057) is routed
   by entries at `L15670`, `L15671` and `L15672` to thunks `L15673`, `L15674` and `L15675`, then to
   handlers `L02360`, `L15505` and `L15506`.
2. Handler `L02360` builds the dialog on the stack (thunk `L15676` to constructor `L15677`, template 225,
   program text `Modify Player`), stores the map object's roster array address in dialog `+0x198` and the current
   index in `+0x19c`, and calls the modal loop (`L15570`).
3. Dialog message map `L15678`: control 0x405 (Computer radio) to `L15665`, 0x404 (Human radio) to `L15666`,
   0x40d with code 1 (selection change) to `L15667`, 0x40e (Add) to `L02367`, 0x40f (Delete) to `L15668`,
   0x444 (Diplomacy) to `L15669`. OnInitDialog is `L15679`, OnOK `L15680`, data exchange `L02365`.
4. Raw scans over the program: a rel32 scan of `.text` for 25 values finds 24 bodies, each referenced once, from its
   thunk; `L15681` (the 25th value) has no reference and was not dword-scanned. Dword scans of `.text` and `.data` for
   the 49 values (24 bodies, 24 thunks, `L15682`) find 0 hits; the `.rdata` scan finds 19 hits, all message maps and
   vtables: 17 lie in tables listed in windows `w04`, `w11`, `w20`, `w27` and `w31`, and two (`L15683`, `L15684`)
   lie past the end of `w11` and are classified by position.

**Confidence.** High for the routing and the handler bodies read whole. Medium for "no other reference": the scans
see a rel32 target or a four-byte equal value, not a computed address or a base-plus-offset reference.

**Unknown.** Vtable slots `L15685`, `L15686`, `L15687` and `L15688` (not read); the edit controls' own
text limits; the 148 `.text` hits of the displacement scan for `+0x94`/`+0xdc` (not classified).

### ALM-DLGPLAYERS-287

Dialog members and their targets: `+0x188` (control 0x40a, spin 0..15) to player `+0x08`; `+0x18c` (0x407, 0..0x7fffffff)
to `+0x0c`; `+0x190` (CString, combo 0x40d, `DDV_MaxChars` 0x20) to the name at `+0x10`. `+0x194` is a plain integer
that `L15665` sets to 1 and `L15666` to 0, that initialisation and selection set from `+0x30`, and that OK writes back
to `+0x30`; the data exchange `L02365` has no routine for it. The ranges and the limit are arguments of the validation
calls `L15689` and `L15690`. The spin messages carry a 16-bit upper bound in the lParam low word, 0x9680, which is
10,000,000 truncated to 16 bits (the meaning of that message is an inference).

- Selection change (`L15667`): saves `+0x30`, `+0x0c`, `+0x08` of the current player, writes the name only for
  index other than 0, refreshes the combo item and reloads the fields.
- OK (`L15680`): `UpdateData(TRUE)` (`L15573`, unread); its result is not tested (`L15691` is followed at once by the
  member loads), as in the selection handler and Add. Then `+0x30`, `+0x0c`, `+0x08` and the name for the current player
  unconditionally, then `+0x04 = i + 1` for every player, then the base OK (`L15571`). The map object's current
  player index (`E+0x328`) becomes the dialog's index + 1 only on OK.
- A 32-character name leaves no room for a NUL inside the 32 bytes; the copy (`ALM-STRCOPY-293`) then stores its NUL
  on the low byte of `+0x30` (inference; the validity of `DDV_MaxChars` in the library body was not read).

**Confidence.** High for the dialog members, the validation arguments and the OK order. Medium for what OK stores when
validation fails, and for the ranges and the 32-character limit as enforced values: `L15573` and the failure tails of
`L15689` and `L15690` are unread (entry bodies of `L15692`, `L15693`, `L15690`, `L15689` read, tails not),
as is the name-overflow inference. Nothing was run.

**Unknown.** `L15573` and the library tails above; the edit controls' own limits; what OK stores after a refused value.

### ALM-DLGPLAYERS-288

Let n be the roster count before the action.

- Add (`L02367`): saves the current player (name unconditionally), builds `new(0x48)` through `L15219`, appends it
  with array SetSize(n+1) (`L14490`, `L15694`), then sets `+0x30` 1, `+0x0c` 5000, `+0x04` n+2 (OK renumbers it to
  n+1), `+0x08` n+1, relation word `i` 2 when `i` equals n and 1 otherwise, and the name `New Player %d` (program
  text at `L15695`) with n+1. The other players are untouched. The body has no test of n against 16; a word at
  index n of 16 or more is not on the wire (`ALM-ROSTERIO-292`).
- Delete (`L15668`): does nothing when the current index is 0 or less. Otherwise it moves the pointer array down over
  the slot and lowers the count by 1, with no destructor and no free; then sets `+0x04 = i + 1` for the remaining
  players. The other players' `+0x08` and relation words, and every unit and structure, are not touched.
- The writer emits an object's owner as the dword `+0x04` of the player object at `[object+0x12]` (windows for
  structures and units). A deleted player's objects keep the pointer, so they write its old `+0x04`; objects of later
  players write the shifted ordinal. For a unit, the reader (type-6 arm, window `w47`) binds `+0x12` to the roster entry
  whose ordinal matches, else to `E+0x3c`, the first roster entry (type-5 window `w41` and `ALM-SAVE-256`). The type-4
  structure arm of the reader was not read (`ALM-SAVE-268` states `+0x12` as the owner without the lookup).

**Confidence.** High for the Add and Delete bodies and the writer's emission. Medium for the composed outcome at reload
(units of a deleted player rebind to whichever player holds that ordinal, else to the first): the unit arm window lies
beyond the experiment's window cap. Unknown for structures at reload (arm unread).

**Unknown.** Whether the pointer to a deleted player is freed elsewhere; the SetSize branch for a size above the
current capacity (branch at `L15696`).

### ALM-DLGPLAYERS-289

- The dialog (class constructor `L15697`, template 270: program text `Diplomacy relations`, radios `Enemy`,
  `Neutral`, `Aliance`) holds 16 int members, each starting at -1; its message map `L15698` is empty.
- Handler `L15669` shows 16 rows: a name for each index below the roster count, each radio set from the current
  player's word through `f` (`L15699`): 0 to 1, 1 to 0, 2 to 2, anything else to 0.
- OK writes all 16 words of the current player's row through `g` (`L15700`): 0 to 1, 1 to 0, 2 to 2, anything else
  to 0. A row at or beyond the roster count is unchecked (-1) and so writes word 0. The player's own row is editable.
  Other players' rows are not written. Cancel writes nothing.
- A stored word above 2 shows as radio 0 and is written back as 1.

**Confidence.** High for the handler read whole to its return and for both mappings. Medium for the library DDX
bodies.

**Unknown.** The order of the three radios in the template against the member values 0, 1, 2; the tails of the library
exchange routines.

### ALM-DLGINFO-290

Handler `L15506` (read whole to `L15701`) copies `E+0x94` and `E+0xdc` into CStrings and `E+0xd8`, `E+0xd4` into
ints, opens the dialog (class `L15702`, template 274: program text `Global map settings`, labels `Map Name`,
`Map Description`, `Rec. Players`, `Min. Level`), and on OK calls `R0260(E+0x94, name)`, `R0260(E+0xdc,
text)`, then stores `E+0xd8` and `E+0xd4`.

- Data exchange `L15703`: CString controls 0x457 (name) and 0x458 (text) with no `DDV_MaxChars`; int 0x459 for
  `E+0xd4` with range 1..50 and int 0x45a for `E+0xd8` with range 1..9999 (file `+0x70` and `+0x74`; `ALM-NEWMAP-236`).
- No other type-0 field (`E+0x18`..`E+0x24`) is written by the dialog.
- The name area is 64 bytes and the text region 512: a name of 64 or more characters spills through `E+0xd4`,
  `E+0xd8` and into the text region; a text of 512 or more characters spills beyond `E+0x2db` (inference from the
  copy, the sizes and the missing bound).

**Confidence.** High for the writes and ranges. Medium for the absence of any bound: the edit controls' own text
limits were not read.

**Unknown.** The edit controls' limits; the order of the dialog's own controls on screen.

### ALM-DLGINFO-291

Handler `L15505` (read whole) counts tiles and objects, passes the counts into the dialog's members (class
`L15704`, template 269, 16 int-only exchange entries, empty message map, base OK), opens it, ignores the result
and stores nothing into the map object.

**Confidence.** Medium: helpers `L15705`, `L15706`, `L15707`, `L15708`, `L15709` and the target of thunk
`L15682` were not read, so a store through them is not excluded.

**Unknown.** What the 16 members display.

### ALM-ROSTERIO-292

1. Reader case 5 (`L15710` and its window) copies `+0x08`, `+0x30`, `+0x0c`, the 32-byte name and 16 words from
   the stream into the player. The writer (`L02295`, emission from `L15711`) writes the same fields, the 16 words
   from the data pointer with bound 0x10, whatever the array count (0x20). No diagonal forcing, resizing or clamping
   was seen in these windows. This answers the Unknown of `ALM-SAVE-263` on the emission.
2. Corpus, 72 installed maps (38 EN, 34 RU; one RU map has no type 5), derived counts: 357 type-5 records (191 EN, 166
   RU); every payload a multiple of 76; 9 records whose word at their own index is not 2 (5 EN, 4 RU); 0 words above
   2; all 357 records hold a word 0; 105 asymmetric pairs (59 EN, 46 RU); `+0x30` 0 in 56 and 1 in 301, no other
   value; `+0x0c` 0 in 97 and non-zero in 260; 123 names with a non-zero byte after the NUL; 0 names without a NUL.
   The map's count word equals the type-5 record count in 71 of 71 maps with a type 5.

**Confidence.** High for the reader and writer windows read and for the corpus counts. Medium for "no
normalisation anywhere": the reader's record-count bound and the code before `L15710` were not read.

**Unknown.** Other writers of roster fields (no displacement scan for roster offsets was run).

### ALM-STRCOPY-293

1. `R0260` (read whole) is an unbounded string copy: it stores bytes through the NUL and nothing past it, no
   padding. This settles the unread tail of `ALM-SAVE-263` item 2 and the Medium for the name tail of
   `ALM-NEWMAP-257`.
2. A dialog edit of the name or text region (`ALM-DLGINFO-290`, `ALM-DLGPLAYERS-287`) leaves the earlier bytes after
   the NUL. In 72 installed maps, 58 type-0 names (29 EN, 29 RU) hold a non-zero byte after the NUL; the text region
   has its first NUL at byte 0 in 67 maps and at byte 1..63 in 5.
3. The reader's defaults (name byte 0 and text byte 0 set to 0, `E+0xd4` and `E+0xd8` set to 1) are stored only when no
   type-0 record was read; a second initialiser at `L15712` stores the same defaults.

**Confidence.** High for item 1 and the corpus counts of item 2. Medium for item 3: the reader window that shows
the condition lies beyond the window cap; that `[ebp-0x1c]` is the type-0 flag is an inference from the stores.

**Unknown.** Whether another code path writes these strings.

## Game-side consumers of the editor's unexplained map fields

Population: the game program, EN `rom.exe` and RU `ROM.EXE` (one SHA-256,
`942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`), static reading: whole bodies by recursive
descent, linear windows, dword tables, one pointer-reachability run (the analyser and eight seeds of EXP-0298,
re-run), rel32 call scans, displacement and immediate scans of `.text`, dword scans of `.rdata` and `.data`. The
installed maps of both roots (38 EN, 34 RU; loose `.alm` files and the `M7R` members of `scenario.res`) for derived
counts. `M` is the map object of loader `R0478`; `O` is a map-browser entry; `T9` a type-9 runtime record.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| ALM-MAPLIST-310 | Type-0 +0x74 is read by R0479 into entry+0x10, posted by R0511 as message 0x461 and formatted by R2271 as the row's last column; among callers of L15876 and L15877 no other reader reads it. The route is Medium. | High / Medium | ● active | [EXP-0564](../experiments/EXP-0564-map-field-consumers/) |
| ALM-META-311 | Over the searched population no game routine reads M+0xd8 (type-0 +0x74), M+0x20 (+0x10) or M+0x24 (+0x14); the world+0x540d0 copy of M is only stored and reloaded in R0277; the light parser seeks over +0x74. | Medium | ● active | [EXP-0564](../experiments/EXP-0564-map-field-consumers/) |
| ALM-T9TAG-312 | Type-9 +0x00 is kept at T9+0x00 from format 990 on; case 8 overwrites it on each linked record; R0461 reads it on its unlinked building arm into caster+0x3c, read only by the serializer among the 14 slot bodies. | High / Medium | ● active | [EXP-0564](../experiments/EXP-0564-map-field-consumers/) |
| ALM-UNITREAD-313 | The type-6 array helpers L02248, R2270, R2121, R1338, R0521 and L13132 store, return or count a record pointer and dereference no record; except unread grow routine L15875; callers: loader, spawner, destructor. | High / Medium | ● active | [EXP-0564](../experiments/EXP-0564-map-field-consumers/) |
| ALM-CPLAYERC-314 | Roster accessors L15878 and R0129 are called from R0128, R0423 and the destructor on M+0x28, and from R0504 and the loader; the two session bodies read no CPlayer+0x0c; installed type-5 +0x08 is 0 or 5000. | High / Medium | ● active | [EXP-0564](../experiments/EXP-0564-map-field-consumers/) |
| ALM-CORP-315 | In 72 installed maps, type-6 +0x26 is 0xffff and +0x28, +0x30, +0x33, +0x34 are 0 in every record; type-0 +0x74 is 1..5; type-9 +0x00 is 0..12; type 9 precedes type 8 in all 71 maps that carry both. | High | ● active | [EXP-0564](../experiments/EXP-0564-map-field-consumers/) |

### ALM-MAPLIST-310

1. Browser test `R0479` (read whole) fills entry `O`: file `+0x00` to `O+0x14`, `+0x04` to `O+0x18`, a seek over
   `+0x08..+0x2f`, the 64-byte name to `O+4`, `+0x70` to `O+0x0c`, `+0x74` to `O+0x10`, and the 512-byte text, line
   feeds replaced by `#`, to `O+8`. It returns 1 only when `O+0x0c` is above 1 (signed compare). These are the reads
   `TEXT-083` lists, re-read here.
2. Its one rel32 caller, `L15879`, is in `R0511` (read whole), which loops FindFirstFileA and FindNextFileA
   over the pattern `\*.alm` (program text, `L15880`) while WaitForSingleObject on `this+0xe8` with timeout 0
   returns non-zero. Per file it allocates a 0x1c-byte entry (constructor `L06328`). An accepted entry goes out by
   PostMessageA(`[L06446]`, 0x461, entry, 0); a rejected one is deleted. `[L06446]` has one store in `.text`,
   at `L09312`.
3. Map-list constructor `L15881` (one rel32 caller, `L15882`) installs vtable `L06390` and passes `L13591`,
   which calls `R0511` with the list object, to `L15883`. The list's routine `R0526` (slot `+0x48`) switches
   on the message minus 0x444 (byte table `L15884`, jump table `L15885`) and sends 0x461 to `R2271(entry)`.
4. `R2271` (read whole) appends the entry to the array at `L+0x6c` (`L15877`) and formats `%s#%dx%d#%d#%d`
   (program text, `L15886`) from the name, `O+0x14`-16, `O+0x18`-16, `O+0x0c` and `O+0x10`, then adds the text
   to the list control through `R2144`.
5. Entry readers among the callers of the element routine `L15876` and append `L15877`: element routine `L15876` has five rel32 callers, `L15887` (list destructor), `L15888` (hover
   routine `L11131`), `L15889`, `L15890` and `L15891` (in `R0526`); each reads `O+0`, `O+4` or `O+8`. Append
   `L15877` has one caller, `L15892`. `O+0x10` is read at `L15893` only.
6. Row draw `L15894` (slot `+0x7c` of vtable `L11178`, one `.rdata` reference at `L15895`, no rel32 caller)
   splits the row text on `#` into four columns. Hover routine `L11131` returns text 0x88 of table `L06197`
   over the fourth column: `dialogs.txt[136]` (`TEXT-083`, which also records the row text and the column captions
   `[134]`, `[135]` and `[136]`).

**Confidence.** **High** for the bodies `R0479`, `R2271` and `R0511` (each read whole), the format string and
the call and scan results / **Medium** for the chain from the post to the displayed row and for "no other entry
reader". The route: 0x461 is posted to the window handle stored at `L09312`, and how it reaches the list's routine
was not read (immediate scan for 0x461: two `.text` hits, the post and a control argument at `L15896`; dword scans
of `.rdata` and `.data`: none). That `L15883` starts a thread is an inference. The reader census is the rel32
callers of `L15876` (five) and `L15877` (one); the entry travels as the `wParam` of a message to a window
procedure that was not read, and the census follows no entry through a copied pointer or an indirect call.

**Unknown.** The dispatcher of slot `+0x7c`; a reader of an entry through a copied pointer, an indirect call or the
unread window procedure.

### ALM-META-311

1. Reach run `map` (the EXP-0298 analyser and its eight seeds): 100 bodies, 548 uses. In `M+0xd4..M+0xdb` it finds
   the loader's read argument at `L15897`, the stores of the absent-type-0 defaults (`M+0xd8 = 1` at `L15898`,
   `M+0xd4 = 1` at `L02153`) and the `M+0xd4` read at `L02087` (`ALM-MODE-070`). It finds no read of
   `M+0xd8..M+0xdb` and none of `M+0x18..M+0x27`.
2. World constructor `R0118` (read whole; the run stops there for want of a frame pointer) passes its map
   argument to `R0278` only and dereferences neither argument. `R0277` (read whole) stores the map pointer at
   world `+0x540d0` (`L02291`), reloads it (`L02369`), destroys the map (`R0522`) and frees it. Displacement
   scan of `.text` for 0x540d0: two hits, those two. Immediate scan: none.
3. Light parser case 0 (window `L15899..L15900`) reads payload `+0x00..+0x18` into `P` and seeks 0x25c bytes,
   over `+0x1c..+0x277`, so `+0x74` has no `P` destination. `R0468` (read whole) stores `P+0x1c` (`L10916`) and
   `P+0x1d` (`L10917`) on its entry block, before its reads at `L10918` and `L10915` (`TERR-LIGHT-149`,
   `TERR-LIGHT-151`).

**Confidence.** **Medium**: a bounded negative over the reach population and the two scans. The run does not follow
its 19 indirect-call stops, 7 callees without a frame pointer, unresolved jump tables or loads from heap memory;
items 2 and 3 are body-scoped reads. The preregistered displacement scan of `.text` for `0xd8`, reduced by the
reach population, was not run, so a reader of `M+0xd8` outside the reach population is not excluded.

**Unknown.** A reader behind an indirect call or a computed offset.

### ALM-T9TAG-312

1. Loader case 9 (window `L02138..L15901`) builds each record with constructor `L15902` (read whole: dwords
   `+0x00..+0x0c` and words `+0x10..+0x18` zeroed). It reads file `+0x00` into `T9+0x00` only when the format version
   is at least 0x3de, then `+0x04` to `T9+8`, `+0x08` to `T9+0xc`, six bytes to `T9+0x10`, `+0x12` to `T9+0x16`, the
   element count into a local and the elements into the array at `T9+0x1a`, and appends `T9` to `M+0x2f0`
   (`L02207`).
2. Loader case 8, for each element whose link is not 0, stores the element ordinal at `T9(link-1)+0x00` (`L15874`)
   and the type-8 object at `+0x04` (`L02231`). In all 71 installed maps with both records type 9 precedes type 8
   (`ALM-CORP-315`), so a linked record's file value is replaced on every one of them.
3. The rel32 callers of the list routines `L15903` and `L15904` lie in the loader, in `R0461` and in
   destructor `R0522`. `R0461` (read whole) reads `T9+0x00` once, at `L15905..L15906`, on the arm taken
   when `T9+8` or `T9+0xc` is non-zero, `T9+0x04` is 0, a bit test passes and a building is found at the cell. It
   stores the low byte at caster `+0x3c` (`UNIT-T9CTOR-110`). So, on the caller census above, only an unlinked
   record's file value reaches a caster; that "only" rests on the Medium holder census (rel32 callers of the list
   routines and the reach run; no `disp:2f0` or `imm:2f0` scan of `.text` was run).
4. The caster class (vtable `L09291`; its 14 slot bodies read whole: 12 in `w21`, and the shared stubs `R0531`
   and `L04402` in `w24`, `RET` and `RET 4`, which read no field) reads `+0x3c` only in serializer `L08387`,
   on the arm that writes the byte to the archive (`L08298`); the other arm reads a byte into it (`R1581`).
   Which arm stores is an inference from the call pattern.

**Confidence.** **High** for items 1 and 2, the bodies of item 3 and the 14 slot bodies of item 4 / **Medium** for
the holder census of item 3, for any reader of `+0x3c` outside the 14 slot bodies and for no other reader: casters
are appended (`L12558`) to a collection whose consumers were not read, and indirect calls and copied pointers are
not followed. The holder set rests on the rel32 callers of the list routines (`s12`) and the reach run; no
displacement scan of `0x2f0` and no `.rdata` or `.data` dword pass for routine addresses without a caller was run.
Whether the archived byte reaches a SAV file is not established here (`SAV-CLASSSER-173`).

**Unknown.** A run-time effect of caster `+0x3c`; a reader of `+0x3c` outside the 14 slot bodies.

### ALM-UNITREAD-313

1. `L02248` (read whole) calls `R2270` (read whole) with the count at `this+8` and the pointer; `R2270`
   grows the array through `L15875` (not read) when the index is at or past the count and stores the pointer at
   `data[index]`. `R2121` returns `this+8`. `R1338` returns the result of `R0521`, the address `data+4*i`.
   `L13132` returns the upper bound. None dereferences a stored pointer.
2. rel32 callers in `.text`: `R2121` from `L15907`; `R1338` from `L15908`, `L15909`, `L15910` and
   `L15911`; `L02248` from `L02247`; `L13132` from `L15912`; `R0521` from `L15913`; `R2270` from
   `L15914`. The array constructor `L15915` has four callers, one of them `L15916` in the loader, and its
   destructor `L15917` is called from `L15918`, `L15919` and unwind funclets. `ALM-MAPLIFE-226` gives the
   roles: loader `R0478`, type-6 spawner `R0151`, destructor `R0522` and the map constructors `R0523`,
   `R0504` and `R0524`. That `L15907`, `L15908` and `L15909` lie in the spawner is taken from their
   addresses; no window of this experiment reads `R0151`.
3. Immediate scan of `.text` for 0x318: 19 hits, the call sites above plus two in loader case 7 (`L02249`,
   `L02250`) whose windows started inside an instruction and decoded nothing. Displacement scan: 9 hits
   (`L15920`, `L15921`, `L15922`, `L15923`, `L14632`, `L15924`, `L15925`, `L15926`, `L15927`), not
   classified by a window here.

**Confidence.** **High** for item 1, except grow routine `L15875`, and for the call list / **Medium** for the
holder set: the two case-7 immediates and the 9 displacement hits rest on `ALM-TAILHOLD-080`.

**Unknown.** Whether `L15875` copies a record.

### ALM-CPLAYERC-314

1. rel32 callers in `.text`: count `L15878` from `L15928`, `L15929`, `L15930` and `L15931`; element
   `R0129` from `L15932`, `L15933`, `L15934`, `L15935`, `L15936`, `L15937`, `L15938` and `L15939`.
   The `M+0x28` base (`ADD ecx, 0x28` before the call) is read at three callers: `R0128`, `R0423` (`w22`) and
   the destructor `R0522` (`w19`). For `R0504` (`L15931`, `L15936`) and the loader (`L15937`) the base is
   taken from the addresses: no window here covers it and `r01` records no `+0x28` use in either body. The three
   `R0526` calls run on another list; that it is the session player list is an inference.
2. `R0423` (read whole) reads each roster player's `+0x04`, `+0x08` and `+0x30` and passes `+0x10`, the name.
   `R0128` (read whole) passes `+0x34`. The destructor deletes each player through its virtual `+4`. None reads
   `+0x0c`; copy constructor `R0503`, called by `R0504`, reads it (`ALM-ROSTERREAD-231`).
3. The reach run shows `M+0x3c` used by the loader only. With no type-5 record the loader builds one player with
   `+0x0c = 5000` (`L02263`).
4. Installed type-5 `+0x08`: 0 in 51 EN records (38 maps) and 46 RU records (33 maps); 5000 in 140 EN records (35
   maps) and 120 RU records (30 maps).

**Confidence.** **High** for the call list, the bodies and the counts / **Medium** for no reader and for the `M+0x28`
base of the `R0504` and loader calls: a call through a pointer and a player copied elsewhere are not followed. This reproduces `ALM-ROSTERREAD-231` from windows read here.

**Unknown.** Any effect of `CPlayer+0x0c`.

### ALM-CORP-315

Rule (fixed before the run): every top-level `.alm` and every `M7R` member of `scenario.res` per root; framing per
`ALM-FRAME-031`; type 0 when its payload has at least 0x278 bytes; type 5 when its size is 76 times type-0 `+0x1c`;
type 6 when the format version is 990 and its size is 70 times type-0 `+0x24`; type 9 when the count-then-`26+6n`
walk consumes the payload exactly. Counts are maps with the value in at least one record, then records.

1. Kept: EN 38 of 38 maps for every type; RU 34 type-0, 33 type-5, type-6 and type-9 (one map has none of the
   three). No framing failure.
2. Type 0 `+0x74`: EN 1 (31 maps), 2 (1), 3 (2), 4 (3), 5 (1); RU 1 (29), 2 (1), 3 (1), 4 (2), 5 (1). On maps whose
   `+0x70` is above 1, the browser population: EN 10 maps (1: 3, 2: 1, 3: 2, 4: 3, 5: 1), RU 5 (1, 2, 3: 1 each; 4: 2).
3. Type 6, EN 8094 and RU 3991 records: `+0x26` is 0xffff and `+0x28`, `+0x30`, `+0x33`, `+0x34` are 0 in all;
   `+0x35` is 0 except 50, 70 and 99 (one record each per root); `+0x22` is 0xffff except 0, 5, 10 and 0xfff6
   (30 records per root); `+0x2a` takes 0, 1, 111, 0xfbd0 and 0xfd00; `+0x1c` takes 0..14.
4. Type 9 `+0x00`, EN 792 and RU 334 records: EN 0..10 and 12 (0 in 575 records); RU 0..6, 9 and 10 (0 in 266).
5. Every map with a type-8 and a type-9 record (EN 38, RU 33) has type 9 first.

**Confidence.** **High** for the shipped files under the rule. The counts say nothing about values an editor session
writes.
