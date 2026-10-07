# Metadata and cell planes

[Reference](format.md)

## Receive source before the light-reader path

The frontend obtains a receive object through global managerL00625 and
entries requested from manager+0x18b8. For each selected source S, the receive
gate sums buffer+0x0f bytes under S+0x1030 locking. This value differs from
the actual buffer remaining-length expression, B+0x7f8 minus B+0x7fc.
The counter's meaning and native producer remain Unknown. — ALM-RECVSOURCE-207

The source reader traverses the list at S+0x1014, requests copies from
B+0x10+cursor, advances the cursor and requests removal/release of exhausted
short buffers. Copy and release helpers remain boundaries. The immediate
code does not prove an end-to-end complete read. — ALM-RECVREAD-208

Received types6/7/8 select shared objectL02436. With its shown constructor
vtable, the reader requests 5 bytes at object+0x0a, then unsigned
object+0x0e plus 1 bytes at object+0x0f; it tests neither read return. Only
type6 selects the frontend light-load arm, and only when view+0x80 is zero
does that arm submit object+0x0f toR0514. The actual input producer,
complete string contents, later aliases and native ordering remain Unknown.
— ALM-RECVOBJ-209

The selected type4 request uses different shared objectL02444. Its shown
serializer submits9 bytes from object+9 into an indexed send-side region.
The laterL02443 body requests a lookup and adds the submitted size to a
counter; it is not itself a payload-copy or list-append operation. The
unexpanded storage/flush helpers do not establish delivery to the receive
list or a causal type4-to-type6 pairing. These boundaries assign no new unit
or gameplay meaning to an ALM metadata scalar. — ALM-SUBMIT-210

## type-0 metadata payload (632 bytes) — `ALM-META-008…ALM-META-010`

The type 0 payload is 632 bytes: 48 bytes of scalar fields, a 64-byte map
name, two u32 values, a 64-byte description and seven further 64-byte text
slots. — ALM-META-008

The offsets below are relative to the payload. The fixed reads total 632 bytes. — ALM-FRAME-031 (amended; framing retained)

| Off | Type | Field | Notes | Claim |
|-----|------|-------|-------|-------|
| +0x00 | u32 | **W** | map width — read into the map object | ALM-HDR-001, ALM-META-024 |
| +0x04 | u32 | **H** | map height (`W≠H` occurs, e.g. 112×144) — read into the map object | ALM-HDR-001, ALM-META-024 |
| +0x08 | f32 | **angle** | Radians; loaded into map `M+0x18` and terrain `P+0x20`. The terrain store fills four bytes of a double. Forced relight replaces it from the sun globals before the known read. Installed angles are whole degrees: `±45, ±44, ±18, 36, 28, 25, 22`; the common π/4 encoding is `0x3f490fda`. | ALM-META-027 (amended; payload-angle clause retained), ALM-META-091, ALM-META-092, ALM-CORP-093, TERR-LIGHT-149, TERR-LOAD-152 |
| +0x0c | 32-bit word; signed integer in editor arithmetic | **Starting Time (editor minutes)** | Loaded into `M+0x1c`, terrain `P+0x2c` and the editor `E+0x1c`. The editor requests slider positions0..96, loads signed `value/15` and accepts `position*15`. Its `(value/60)%24` lighting path reaches palette colors and render-buffer stores. Default360 is 06:00. The identified normal game M/P paths do not read this scalar; a universal or native game effect remains unestablished. Native editor transactions are unobserved; slider bounds are not an ALM validation range. | ALM-EDITTIME-205, ALM-SCALAR-198, ALM-EDITORSCALAR-203 |
| +0x10 | u32 | scalar (stored) | Loaded into map `M+0x20` and terrain ambient byte `P+0x1c`. Forced relight overwrites the terrain value before its known read. Installed values span `0..33`; this is not a validation bound. | ALM-META-026, ALM-META-091, ALM-META-092, ALM-CORP-093 |
| +0x14 | u32 | scalar (stored) | Loaded into map `M+0x24` and terrain range byte `P+0x1d`; the four-byte terrain store also writes `P+0x1e/0x1f/0x20`. Relight replaces the range before its known read. Installed values span `27..64`; this is not a validation bound. | ALM-META-026, ALM-META-091, ALM-META-092, ALM-CORP-093 |
| +0x18 | u32 | bitmask | **Terrain tile-group mask.** The map loader discards its local copy; terrain stores it at `P+0x28`. Bit i selects group `(i>>2)+1`, variants `(i&3)*4 .. +3`. Installed values use bits `0..12`: groups 1–3 and group 4 variants 0–3. | ALM-META-026, TERR-LOAD-152 |
| +0x1c | u32 | **#players** | player-record count = `type5_size / 76` ; installed values `3..9` (editor caps at 16) | ALM-META-025 |
| +0x20 | u32 | **#objects** | object-record count (base records; extensions are additional bytes, not records) | ALM-META-025 |
| +0x24 | u32 | **#units** | unit-record count = `type6_size / 70` in the version-990 form | ALM-META-025 |
| +0x28 | raw 32-bit word | count-local residue, purpose Unknown | Read4 into EBP-0x50, overwritten by each complete internal type-7 count read. Short reads can retain prior bytes; this does not make it an external type-7 count | ALM-META-025 (discard shorthand narrowed), ALM-COUNT-195, ALM-STALE-196 |
| +0x2c | u32 | **#type 8 records** | count used as the case-8 loop bound; `0 ⟺ type8 empty` | ALM-META-025 |
| +0x30 | char[64] | **name** | NUL-terminated ASCII. Installed names use at most 21 bytes and can be empty; 21 is not a field-width limit. | ALM-META-010 |
| +0x70 | u32 | scalar (stored) → `map+0xd4` | **Multiplayer-mode source:** `[L00285]+0xc = (map+0xd4 > 1)`, read during player construction. A value 1 selects single-player loading even in a loose map such as RU `Horror.alm`; it does not identify `scenario.res` membership. The map-selection list loads it into a record column and lists a map only when it is above 1. A player-slot/MP-capacity interpretation beyond these consumers remains unestablished. | ALM-META-026, ALM-MODE-070, TEXT-083 |
| +0x74 | u32 | scalar (stored) | standalone `1..5`, campaign `1`; the map-selection list shows it in its last column | ALM-META-026, TEXT-083 |
| +0x78 | char[64] | **description** | NUL-terminated code-page text (ASCII or Windows-1251). Installed text uses at most 36 bytes; 36 is not a field-width limit. The map-selection list hover returns it, with line feeds read as `#`. | ALM-META-010, TEXT-083 |
| +0xb8 | 7×64 B | **text slots** | fixed array of 7 slots `[3×u32 prefix][char[52] text @+12]`, default `"<None>"`. Loaded (part of a 512-B block read from +0x78). Campaign maps fill slots 4/6 with trigger/quest text (`"mission complit"`, `"Start1"`, …) | ALM-META-028 |

The negative consumer statements above are limited to the named map/terrain
paths. Indirect calls, aliases and the skipped-relight/message-path cases
remain Unknown; absence of a reader is not a semantic default.
— TERR-LIGHT-149, TERR-LIGHT-150, TERR-LIGHT-151

The record's `typeId = 0` and opaque four-byte word are [record-header
fields](container.md) at `+0x0c/+0x10`; they are not payload fields.

The whole type-0 record is read by `rom.exe`'s `.alm` loader (`R0478`, case 0) in
16 reads totalling exactly **632** bytes: twelve `u32` (`W/H`, the light `angle`, three
stored scalars, the discarded bitmask, the five content-record counts), `0x40` (`name`),
two `u32`, and one `0x200` block covering `description` + the 7 text slots. The `π/4`
constant is `+0x08` (`0x3F490FDA`).

## Grid layers type 1 / type 2 / type 3 — `ALM-GRID-012`, `ALM-GRID-013`,
`ALM-GRID-014`, `ALM-GRID-032`

Each grid record stores **`W·H` cells starting at payload+0** — the payload being the
bytes after the record's 20-byte header — with `payloadSize` exactly `2·W·H` (type 1) /
`W·H` (type 2, type 3). The payload is **pure grid**: nothing is overlaid on it, no cell is
lost, and the record's `typeId`/opaque word live in the header (`ALM-GRID-032`). Cell layout is
row-major, `W` cells per row:

```
cell (col, row)  ->  element index  row*W + col      (0 <= col < W, 0 <= row < H)
type1: u16 at payload + 2*(row*W + col)
type2: u8  at payload +    row*W + col
type3: u8  at payload +    row*W + col
```

The world ingest addresses `tiles[row*W+col]`, `heights[row*W+col]` and
`objects[row*W+col]` from the payload origin. — ALM-GRID-032

| Layer | Name | Cell | Encoding | Claim |
|-------|------|------|------------------------------------------------|-------|
| type 1 | **Tiles** | u16 LE | Bits 0–9 feed the simulation classifier; bits 0–12 feed graphics selection. Bit 13 feeds the simulation block assignment but is cleared by the light parser. Bits 14/15 carry render visibility state. An installed absence is not an unused-bit rule. | ALM-GRID-012, ALM-TILEMAIN-121, ALM-TILEVIEW-122 |
| type 2 | **Altitudes** | u8 | Height/altitude, copied to the simulation height buffer. Installed heights are below `0x80`; this is not a general u8 limit. The earlier range derived from the displaced grid origin is retracted. | ALM-GRID-013, TERR-LIGHT-016, TERR-LIGHT-028 |
| type 3 | **Objects** | u8 | Static-object code: 0 is empty; nonzero `c` selects `objects.reg` section `c-1`. Missing type 3 is zero-filled. World ingest additionally derives runtime block value 5 for nonzero cells. | ALM-GRID-014, ALM-CLS-035 |

### type 1 tile word — terrain resolution (`rom.exe`)

`terrainClass, cost = f(word & 0x3ff)` (`R0469`); the block assignment is separate:

- **strip group = bits 6–9** indexes a hardcoded (primary, secondary) terrain-type pair
  table at `world+0x54156`; **blend variant = bits 0–5** selects primary vs secondary
  and a **5-level** blend of their costs (auto-tiling terrain transitions).
- In the water range `(word & 0x300) == 0x200`, subcell 8–15 returns class
  `0xff`; below 8, subcell 4 with variant 1 returns class 1 and the others
  return class 9. All leave the classifier cost output at 8. Ingest replaces
  that cost with slot 0 (`0xff`) for an invalid class and independently sets
  block 1 for every raw-water word. — TERR-WATERBOUND-164
- **Bit 13** assigns block 1 before later object and border writes.
  It does not itself select a terrain class. — TERR-TILECONTROL-163

**Terrain enum** (1-based in the lookup; = `world.res:data/map.reg` `Terrain` record
order): `1 Land · 2 Grass · 3 Flowers · 4 Sand · 5 Cracked · 6 Stones · 7 Savanna ·
8 Mountain · 9 Water · 10 Road` (`ALM-TERR-015`). `map.reg` supplies a per-terrain `Cost`
and `Pass` scalar; `rom.exe` reads **only the ten `Cost*` keys**, into the byte table
`world+0x54176+class` (slot 0 = `0xff`), with its own hardcoded defaults
`[8,8,9,14,6,12,11,16,8,6]`.

**Restored by `ALM-TERR-043`** on the corrected `.reg` framing:

| class | 1 Land | 2 Grass | 3 Flowers | 4 Sand | 5 Cracked | 6 Stones | 7 Savanna | 8 Mountain | 9 Water | 10 Road |
|---|---|---|---|---|---|---|---|---|---|---|
| `Cost` | 8 | 8 | 8 | 14 | 6 | 12 | 8 | 16 | 8 | 6 |
| `Pass` | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 1 | 1 | 1 |

The displaced Cost/Pass vectors associated with superseded `REG-FMT-017`
are invalid. Use the per-class values in the table above.

**`Pass` is inert.** No `Pass*` key string exists anywhere in `rom.exe` or in
`Map Editor.exe`; impassability is hardcoded (terrain class 8, plus the raw water bit
test). `Cost` is a **movement cost**, consumed only by the pathfinder's per-step add and
by the move-duration divide (`TERR-COST-052`). `map.reg` also carries a `Path Finding`
section — `SpeedMultiplier 8`, `StaticScanAhead 5`, `DynamicScanAhead 3`,
`StaticRefreshRate 16`, `DynamicRefreshRate 32`, `DynamicByStaticLookup 3`,
`StaticIsntNeeded 5` — and `Scanning`/`ScanShift 7`. The `.reg` name field holds 15
characters and `rom.exe` asks for keys up to 21; both lookup paths truncate the request
to 15 first, so every key still resolves.

**Terrain graphic** (which picture a tile-word draws) is a separate render mapping,
specified in [TERRAIN](../terrain/format.md) (`TERR-IDX-003`, amended bit interval;
arithmetic retained, `TERR-SEM-004`): `g = (w & 0x1fff) >> 6` selects a
`terrain.3d/tileG-VV.bmp` strip (`G=(g>>2)+1`, `V=(g&3)*4+((w>>4)&3)`) and `w & 0xf`
selects the 32×32 sub-cell. `tile1/2` = land strips, `tile3` = animated water, `tile4` =
road; bit-13 tiles composite over `dirt.bmp`.

### Main and light-parser tile storage

The main type-1 arm requests `2*W*H` bytes into `M+0x0c`, with dimensions
at `M+0/+4`. With a complete lower read, it preserves the word before the
simulation ingest. The light parser has dimensions at `P+4/+8` and its own
tile pointer at `P+0x0c`; after reading, it applies `word &= 0xdfff` to every
tile. View binding stores `P` at `view+0x80`, and renderers use that pointer.
The binder copies dimensions, not the tile plane. Equal member offsets do
not prove that the two reader layouts share an allocation.
— ALM-TILEMAIN-121, ALM-TILEVIEW-122

The light parser preserves `0x4000` as well as `0x8000` and `0xc000`.
Later render-grid updates can replace bit 13, stamp bits 14/15 together or
clear bit 14. Load-time words and post-event words must be distinguished.
See [tile selection](../terrain/tiles.md) and also [visibility](../terrain/fog.md).
— ALM-TILEVIEW-122

**Runtime passability** — corrected and completed by `TERR-PASS-049`, `TERR-PASS-050` (partially retracted), `TERR-PASS-051` (amended, superseded).
The map load derives **three** 256×256 byte planes at fixed
stride 256, addressed `(row<<8)|col` regardless of `W`,`H`:

| plane | at | initial fill | built from |
|---|---|---|---|
| movement cost | `sim+0x00000` | `0x01`, overwritten for every in-bounds cell | the blended terrain cost |
| block bits | `sim+0x10000`, copied to `sim+0x20000` | `0` | type 1 + type 3 + the border |
| height | `sim+0x9451c` | `0` | type 2 (Altitudes) |

The block byte is a **bitmask**, not an enum: `1 = bit0`, `5 = bit0|bit2`,
`0x1f = bits 0..4`; bit 5 marks a cell carrying a runtime record, bits 6/7 a ground/air
occupant on the dynamic plane. A cell blocks a mover iff `block[cell] & mover.mask != 0`
over the mover's `n×n` footprint, mask `0x41` ground / `0x44` / `0x82` air. Bit 1 is set
by nothing but the border, so **only the border stops an air mover**. The block arms, in
the order the ingest applies them:

```
if (w & 0x2000)                block = 1     tile-word bit 13
if (classify(w & 0x3ff) == 8)  block = 1     Mountain
if ((w & 0x300) == 0x200)      block = 1     water range, tested on the raw word
if (type3[cell] != 0)          block = 5     static object  (assignment: this one wins)
8-cell border                  block = 0x1f
```

Both block planes are **save state** (`TERR-PASS-053`); the cost and height planes are
not. So is the cell-record map the next paragraph names.

**That is the whole of what the ingest writes — it is not the whole of what blocks.** A placed
**type-4 structure** never goes through it: it attaches, in its own constructor and strictly after
the ingest, to a per-cell record, and one routine then recomputes that cell's cost byte and *both*
block bytes from the record. Its footprint can **clear** bits 0 and 2 as well as set them, which is
how a bridge crosses water. Specified in [TERRAIN](../terrain/format.md) → "Structures on the block plane"
(`TERR-STRUCT-068` (amended, superseded)…`072`, `TERR-PASS-073` (amended, partially retracted)); the `kind`→class resolution is `ALM-CLS-036` above.

The record-header opaque word is separate from metadata payload+0x08.
No numeric interpretation follows from its bit-pattern census. — ALM-HEADER-098

For payload+0x0c, the editor loader wrapper returns the same E object whose
address is installed at document+0x50. When E+0x0c is nonzero, the load handler also creates a map
backup at document+0x54. The separate copy constructor explicitly copies
E+0x1c to destination+0x1c. The writer caller selects document+0x50 and
submits one four-byte item from that receiver+0x1c. Complete native load,
intervening control changes, replacement and emitted-file preservation
remain Unknown. — ALM-EDSCALAR-102, ALM-METAEDITOR-185

The game and landscape read prefixes transfer this scalar unchanged only
when their lower read completes. Resolving that stream edge does not close
later map/landscape aliases or establish meaning. — TERR-STREAM-157

### Metadata scalar lifetime bounds

Game copy owner R0504 reads M+0x1c at L02373 and stores the word in the
destination at L02374. The copy has independent scalar storage, but native
entry into this owner remains unresolved. A global no-reader or no-copy
interpretation is unsupported. — ALM-METACOPY-183

Owning world constructorR0277 publishes M at world+0x540d0,
ingests its planes, then destroys and frees M before returning without
clearing that holder. Only its publication/reload instructions contain the
literal holder displacement in the pinned image. The separate borrowing
constructor uses its caller's M; mission and save-resume paths delete that M.
Five pre-teardown consumers and both exact eight-entry placement switches
use other map fields. The selected complete lifetimes contain no M+0x1c
read. This is a result for the identified receivers and call conventions,
not a global absence proof for arithmetic-built aliases or native callbacks.
— ALM-METALIFE-184, ALM-INGEST-201, ALM-FRONTIER-200

The editor's new-map constructorL02381 writes 360 to E+0x1c. This is a
creation default, not a constraint on loaded values. The Light dialog's
Starting Time control reads signed minutes/15 and accepts position*15.
Its time-to-color consumer identifies the default as 06:00. The earlier
sixteen-command search did not include this established control chain.
Native control returns and SAVE transactions remain unobserved.
— ALM-METADEFAULT-186, ALM-EDITORSCALAR-203, ALM-EDITTIME-205

Landscape P is a distinct 0x30-byte object. Its binder stores the pointer at
view+0x80 and copies dimensions, without copying the header. A subsequent
old-P deletion through the constructor's actual vtable frees four arrays and
P. Neither the binder nor those destructor bodies reads P+0x2c. The
allocator uses page descriptors outside the P payload. The actual whole-P
renderer, additional view holders and sixty-three new P getters now have
normal-path dispositions: dimensions, planes and angle are used, while this
scalar is not read. Native ordering and universal all-image non-use remain
Unknown. — TERR-METALIFE-175, ALM-FRONTIER-200

## Scalar transfer versus semantic closure

A complete primary Read4 maps metadata+0x0c to M+0x1c; the separate light
reader maps it to P+0x2c. The original explicit map-copy load/store preserves
all four bytes in independent storage even after source poisoning. The browser
instead skips this word and metadata+0x28 in its seek from payload+8 to+0x30;
the light reader skips metadata+0x28 after its first seven scalar reads.
These transfer operations alone do not establish a unit. The independent editor
control and color chain establishes minutes on its Starting Time scale; a complete
ROM1 gameplay role does not follow from that editor result.
— ALM-SCALAR-198, ALM-EDITTIME-205

Complete type-7 internal counts replace metadata+0x28 in their shared local and
produce their own three array populations. Incomplete reads can leave metadata
or a previous internal count there. Returned byte count, requested width, wrapper
cursor, and destination contents must be tracked separately; a surviving byte is
not necessarily from the current wire field. — ALM-COUNT-195, ALM-STALE-196

The editor's Starting Time control and its palette/render-buffer effect are
established independently of the transfer tests. ROM1's selected M and P
consumers read dimensions, planes and other lighting fields; they do not
inherit an editor effect merely because they retain the same wire word.
Metadata+0x28's editor count-sum producer is established; its game purpose
outside the measured loops remains Unknown. — ALM-FRONTIER-200, ALM-EDITTIME-205

### Bounded ingest controls

The 304 additional controls execute the original world ingest and both callees.
They separate six map-scalar words from M/null/unmapped values in the held-pointer
slot, across two fills, four synthetic shapes and two table variants. Five original
classifier branches are visited; sixteen controls remove the observation bridges
and retain identical output. Only the deliberately varied pointer word is excluded
from the whole-world comparison. This is a finite no-output-effect result, not proof
that no read is ever discarded or that the field is globally unused. The body
copies world byte planes, not the map header. — ALM-INGEST-201

Fresh original-editor controls place complete scalar Read4 output at E+0x1c
and retain the old suffix after short lower transfers. An independent copy
survives source poisoning and reaches the original buffered writer. The
separate new-map store writes0x168 after the actual prior word and fill are
verified at entry. The separate control/consumer chain identifies the editor unit
as minutes and the default as 06:00. Native control returns and complete editor
transactions remain unobserved. — ALM-EDITORSCALAR-203, ALM-EDITTIME-205

The original editor produces metadata+0x28 by adding its action, condition
and Trigger cardinalities. All64 triples0..3 under two fills write the sum
through the buffered writer. The game's complete type7 reads still select
the three internal count words for their own loops. Static arithmetic also
establishes modulo2^32 addition; native SAVE remains unobserved.
— ALM-EDITCOUNT-204, ALM-TRIGMETA-166

The post-load P route is bound for the constructor-produced application,
frame, view and child families. Its identified view aliases and P getters
use dimensions, planes and angle, with no scalar reader. Nonzero app+0x20
still overrides the primary frame, and the host fallback is not a primary-frame
identity proof. Native selection and external member replacement remain
unobserved. — ALM-FRONTIER-200
