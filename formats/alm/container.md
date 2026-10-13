# Container and load rules

[Reference](format.md)

## Structure

The header gives `recordCount`. Each record supplies `payloadSize` and
`typeId`. The standard record chain ends at EOF; record count and type order
are data, subject to the loader dependencies below. — ALM-FRAME-031 (amended; framing retained),
ALM-CORP-060

```
+----------+-------------------------------------------------+
|  header  |          recordCount records (chain)            |
|  20 B    |   each: 20-B header + payloadSize B payload      |
+----------+-------------------------------------------------+
0        0x14                                               EOF
```

**Required: types 1 and 2 only.** See *Acceptance contract* before implementing a reader.

## File header (20 bytes)

| Off | Type | Field | Notes | Claim |
|-----|------|-------|-------|-------|
| 0x00 | char[4] | magic | `4D 37 52 00` = `"M7R␀"` — compared by the selected readers | ALM-HDR-001 |
| 0x04 | u32 | **hdrLen** | `= 20` — this header's **own length**, and it is *used*: the header helper peeks 8 bytes, seeks back, then reads `dword[cursor+4]` bytes. At offset 0 that reads 20 | ALM-HDR-001, ALM-FRAME-031 (amended; framing retained) |
| 0x08 | u32 | dataSize | Installed version-990 value `4·W·H + 72`; the loader ignores it. The writer-side meaning of 72 is Unknown | ALM-HDR-001 |
| 0x0C | u32 | **recordCount** | the loader's loop bound, gated `≥ 3` and nothing more; `= 10` in every authored map, `= 4` in `ru/Horror.alm` | ALM-META-024, ALM-FRAME-031 (amended; framing retained), ALM-CORP-060 |
| 0x10 | u32 | **formatVersion** | Installed value990; primary/editor unsigned upper gate1001 and conditional payload thresholds951,985,987,989,990. The 1000 helper rule differs by reader; see below | ALM-META-024, ALM-FRAME-031 (amended; framing retained) |

`W`,`H` themselves live in the [type 0 record payload](metadata.md), not the header.
`rom.exe` `R0466` reads this header as one `Read(dest, 0x14)`.

## Record header (20 bytes) — repeated `recordCount`×, from 0x14

| Off | Type | Field | Notes | Claim |
|-----|------|-------|-------|-------|
| +0x00 | u32 | tag | installed value `7`; stored to the map object (`+0x08`) and never read again | ALM-SEC-002, ALM-FRAME-031 (amended; framing retained) |
| +0x04 | u32 | hdrLen | constant `20` (= this header's own length); read, not validated | ALM-SEC-002 |
| +0x08 | u32 | payloadSize | payload byte length; the payload follows immediately | ALM-SEC-002 |
| +0x0C | u32 | **typeId** | the record's type, `0..9`; **this is the word the loader's `switch` dispatches on** (10-entry jump table) | ALM-SEC-003 (amended; type dispatch retained), ALM-FRAME-031 (amended; framing retained) |
| +0x10 | byte[4] | locally unconsumed word | No value consumer or surviving copy in the enumerated normal reader paths. Next complete overwrite/frame release ends local ownership. The editor writer supplies its own uninitialized stack slot. Numeric representation and original intent remain Unknown. | ALM-HEADER-098, ALM-CENSUS-101, ALM-HDRLIFE-206 |

Each record payload starts immediately after its 20-byte header and has no
additional type identifier. The loader reads the header with `Read(dest,0x14)`
except when `formatVersion == 1000`, then dispatches the payload grammar by
the header's typeId. A payload may be empty or shorter than eight bytes.

## Record roster — what the writer emits, and what the loader requires

The complete authored order is `0,1,2,3,5,4,9,8,6,7`. The primary loader
dispatches the declared recordCount, requires only types 1/2 and permits
other record sets. The four-record RU Horror.alm satisfies these structural
checks. — ALM-SEC-003 (amended; type dispatch retained), ALM-REQ-055, ALM-REQ-056, ALM-ORD-057,
ALM-META-058, ALM-RDR-059, ALM-CORP-060

Sizes by type (`ALM-SEC-004`):

| typeId | payloadSize | Role |
|--------|-------------|------|
| 0 | 632 (constant) | **metadata** — decoded below (`ALM-META-008…ALM-META-010`) |
| 1 | `2·W·H` | **grid: tile word** (u16/cell) — decoded below (`ALM-GRID-012` (amended), `ALM-GRID-032`) |
| 2 | `W·H` | **grid: relief/height** (u8/cell) — decoded below (`ALM-GRID-013`, `ALM-GRID-032`) |
| 3 | `W·H` | **grid: static-object placement layer** (u8/cell) — cell code `c ≠ 0` is `objects.reg` section index `c − 1` (`ALM-GRID-014`, `ALM-CLS-035`) |
| 4 | `20·#4` (+8·ext) | **placed structures/buildings** — 20-B records, `#4 = meta+0x20` (the loader's own case-4 loop bound); every record is a `structures.reg` class (`ALM-OBJ-019`, partially retracted for its `+0x12` destination; `ALM-PLACE-033`, `ALM-CLS-036`) |
| 5 | `76·#5` | **player/group roster** — 76-B named records, `#5 = meta+0x1c` (`ALM-GRP-020`, `ALM-PLACE-033`) |
| 6 | `70·#6` | **placed units** — 70-B records, `#6 = meta+0x24` (`ALM-UNIT-018` (partially retracted), `ALM-PLACE-033`) |
| 7 | variable | **trigger effect/instant list + Drop table** — `[u32 entryCount][named blocks]` (`ALM-TRIG-021`) |
| 8 | variable, sometimes **0** | **authored loot**, `meta+0x2c` records with no count word. The box/circle reading in `ALM-TRIG-022` is retracted; use `ALM-SACK-065` |
| 9 | variable (small) | **caster list** — `[u32 count@+0][records]` (`ALM-CODE-023`, `ALM-TRIG-049` (partially retracted)) |

The three fixed-record sections carry their **count in the type-0 metadata**
(`+0x1c`=type 5, `+0x20`=type 4, `+0x24`=type 6, `ALM-CNT-017`); the variable trigger
sections carry their own count word.

Metadata `+0x1c/+0x20/+0x24/+0x2c` supplies the counts for types 5/4/6/8.
**Metadata +0x28 is not a type-7 loop bound after complete internal reads.**
Type 7 reads three internal counts; type 9 reads its own count. Do not derive
a metadata+0x28 writer obligation from type-7 cardinalities. An incomplete read
can leave old bytes in the reused count local; that is not format semantics.
— ALM-COUNT-195, ALM-STALE-196; ALM-REQ-056 (type-7 clause partially retracted)

## Acceptance contract — what a reader must and must not require

Placed-structure footprint registration must not be reused as the physical-attack token
size. The base Building attack vtable supplies size 1 even though `+0x60/+0x61` carry its
registered rectangle (`UNIT-STRUCTREACH-063`). Its destructor detaches the footprint, but
the physical strike's HP write does not itself invoke that destructor. The HP-to-destruction
scheduler and derived-class overrides remain Unknown in this bounded route
(`UNIT-STRUCTSTOP-066`); no new type-4 byte layout is implied.

Read from `R0478` and its consumers (`ALM-REQ-055`, `ALM-REQ-056`, `ALM-ORD-057`,
`ALM-RDR-059`). The engine's own names for the two mandatory records are in its error strings.
The type-7 count clauses of ALM-REQ-056 and ALM-ORD-057 are partially retracted;
the unrelated requirements below are retained.

| record | absent → | why |
|---|---|---|
| **type 1** *(“Tiles”)* | **reject** — loader status 5, `"Tiles block not found"` | the world builder `R0278` reads `map+0x0c` as u16/cell with no null test |
| **type 2** *(“Altitudes”)* | **reject** — loader status 6, `"Altitudes block not found"` | same, `map+0x14` as u8/cell |
| type 3 | **default**: a `W·H` plane of zeroes | same consumer reads `map+0x10` as i8/cell; nonzero = “an object blocks this cell”, so all-zero is a valid empty plane |
| type 5 | **default**: one group record — `new(0x48)`, `R0510(1,1)`, then `CPlayer+0x0c = 5000` (`ALM-REQ-056`, type-7 clause partially retracted; player-default clause retained), which is the field a record's own `+0x08` fills (`ALM-SCALAR-087`). Every shipped map authors `0` in that position on record 0 (`ALM-SCALAR-089`) | the group list must be non-empty |
| type 0 | **default**: `W = H = 16`, empty name and text, `+0x70 = +0x74 = 1` | see the actual dimension and metadata-count dependencies below; no type-7/type-9 count dependency on type 0 is established |
| type 4, 6, 8 | **skip when absent** | their present-record loops use metadata counts; absence does not itself run a loop |
| type 7, 9 | **skip when absent** | when present, they read their own internal counts; type 7 has three, not a metadata+0x28 bound — ALM-COUNT-195 |
| `typeId >= 10` | **skip** by `Seek(payloadSize)`, no error | the switch's `default` arm |

Header gates, in order: magic `M7R␀`; `recordCount >= 3`; `formatVersion <= 1001`. A duplicate
`typeId` is **not** rejected — the case runs again and overwrites the pointer.

**Ordering.** The permutation is not enforced, but three precedence relations are, because two
lengths and four metadata-count bounds are carried in state from earlier records:

1. **type 0 before type 1** — the tile grid is allocated `W·H·2` from the *metadata's* `W`/`H`
   (default `16,16`), not from the record's own `payloadSize`.
2. **type 2 before type 3** — type 3's allocation and read length is the **type-2 record's**
   `payloadSize`.
3. **type 0 before types 4/5/6/8 for their metadata-count bounds.** Type 7
   reads its own three counts and type 9 its own first count. This removes only
   the false count dependency, not other object/pointer ordering requirements.
   — ALM-COUNT-195; ALM-ORD-057 (type-7 clause partially retracted)

**Four readers ship, with three different acceptance tests** (`ALM-RDR-059`). Beside the loader,
`rom.exe` has the `*.alm` browser scan's accept test — magic, `recordCount >= 2`, a type-0 record,
and type-0 payload **`+0x70 >= 2`**, so a map below that threshold is silently unlisted
(`ALM-META-058`) — and a light parser that handles typeIds 0..3 only, walks to EOF rather than by
count, and treats its two header failures as advisory message boxes without stopping. `Map
Editor.exe` has a separate reader and a distinct version-1000 header permutation;
partial reader mechanics are established, but its complete native acceptance
contract remains Unknown. — ALM-HEADER-097, ALM-FRONTIER-200

## End of file

The standard chain ends after its final payload. There is no separate
eight-byte trailer; ALM-TRL-005 and ALM-TRL-030 are retracted. The final two
words in a type 7 payload with no conditions or triggers are its zero count
words. — ALM-FRAME-031, ALM-TRIG-044

## Reading algorithm

```
assert bytes[0:4] == "M7R\0"
hdrLen    = u32(bytes, 0x04)             # == 20
dataSize  = u32(bytes, 0x08)             # == 4*W*H + 72
count     = u32(bytes, 0x0C)             # the loader's loop bound; gated >= 3, NOT == 10
fmtVer    = u32(bytes, 0x10)             # == 990
off = 0x14
for k in 0..count-1:
    tag, hdrLen = u32(off), u32(off+4)    # == 7, 20
    size        = u32(off+8)
    typeId      = u32(off+12)             # <- the loader's switch; >= 10 is skipped by seek
    opaqueWord = bytes(off+16, 4)
    payload     = bytes[off+20 : off+20+size]
    off        += 20 + size
assert off == len(bytes)                  # no trailer
# acceptance: require types 1 and 2 only. Default type3 to a W*H zero plane, type5 to one
# group record, type0's fields to W=H=16 / empty text / +0x70=+0x74=1. Skip the rest.
# Do NOT require 10 records, the full set {0..9}, or the physical order.

# grids (payload = pure data):
#   type1: u16 cell (col,row) at payload + 2*(row*W + col)
#   type2: u8  cell (col,row) at payload +   (row*W + col)
#   type3: u8  cell (col,row) at payload +   (row*W + col)
#
# content records (meta = type-0 payload):
#   type5: 76-B records, count = u32(meta,0x1c)            -> id @ rec+0, name @ rec+0x0c
#   type6: 70-B records, count = u32(meta,0x24)            -> X @ rec+0, Y @ rec+4 (/256),
#          class = units.reg ID @ rec+8, owner slot @ rec+0x14
#   type4: 20-B records, count = u32(meta,0x20)            -> X @ rec+0, Y @ rec+4,
#          kind = structures.reg ID @ rec+8, owner slot @ rec+0x0e;
#          kind == 0x21 -> 8 extra bytes follow (footprint override);
#          kind in {0x22,0x23} -> Shop, else Building
#
# class binding:
#   type3 cell code c != 0  ->  objects.reg    section index c - 1
#   type4 kind              ->  structures.reg ID (= section index + 1)
#   type6 rec+8             ->  units.reg      ID (sparse 1..80; NOT the section index)
#   owner (type4 +0x0e, type6 +0x14) -> 1-based physical slot in the type-5 array
#   type7: [u32 nAct][nAct x 796][u32 nCond][nCond x 796][u32 nTrg][nTrg x 184]
#          node    : name[64], opcode, id, 0, value[10], type[10], pname[10][64]
#          trigger : name[64], 64 B junk, 3x(condId,condId), 4x actionId, 3x opcode, flag
#          a trigger slot holds a node's id (node+0x44), NOT its array index
#   type9: [u32 count][count x (26 + 6n)], n = u32 at rec+0x16; X,Y = tile indices
#   type8: NO count word; meta+0x2c records of (20 + 10n), n = u32 at rec+0; the
#          head is 20 B at formatVersion >= 0x3dd and 16 B below it;
#          X,Y at rec+8/+0xc are 0x80-centred /256 anchors
#
# placement -> tile: cell = (X >> 8, Y >> 8). Nothing else: no origin, no inset.
```

## Write sequence

1. Select the format version. This reference's fixed-size placement layouts
   describe version 990; apply each documented version branch for other values.
2. Build metadata and payloads. Derive counts from the actual records and use
   `W*H` cells in each present grid. A metadata count alone creates no payload.
3. Emit the 20-byte file header: magic, header length, dataSize, recordCount,
   formatVersion. Version-990 installed maps use `dataSize=4*W*H+72`; its
   decomposition remains Unknown.
4. Emit each record header followed by its payload. The installed full-map
   order is `0,1,2,3,5,4,9,8,6,7`. Preserve metadata before dependent records,
   type 2 before type 3, and type 9 before type 8 before type 6.
5. Emit exact type-specific counts, optional fields and extensions. Do not add
   a trailer or duplicate counts where the grammar uses metadata.
6. Validate record endpoints, grids, class/owner relations and script IDs.
   A type 7 trigger refers to a node's ID, not its array position.

The structural size relation is `20 + sum(20+payloadSize) = fileSize` for
the standard complete form. Types 1 and 2 are required by the primary loader;
the complete installed type set is not an acceptance requirement.
— ALM-FRAME-031 (amended; framing retained), ALM-ORD-057, ALM-ORD-068, ALM-TRIG-044,
ALM-TRIG-045, ALM-TRIG-046, ALM-TRIG-047 (amended: opaque-region gloss withdrawn), ALM-CORP-060

## Map editor writer

The EN map editor's writer emits types 4 and 6 only when the map holds at least
one object. A new map therefore has 8 records in the order `0,1,2,3,5,9,8,7`:
type 0 632 bytes, type 1 `2*W*H`, types 2 and 3 `W*H`, type 5 152 (two
players), type 9 4, type 8 0 and type 7 12. Every record header carries the
map's tag word: 5 for a new map, and for a loaded map the tag read from its
type-1 record header (7 in every installed record header). W and H are the
dialog values plus 16: 32..256 in steps of 8, 80 by default. — ALM-NEWMAP-237

The writer applies the tile and count normalisations of the
[metadata page](metadata.md) on every save. All 72 installed maps (38 EN, 34 RU)
carry `dataSize = 4*W*H+72` and one record-header `+0x10` value per map.
— ALM-SAVE-238

A save can also change records of a loaded map:

- A loaded map whose objects were all deleted is written without types 4 and 6.
- A unit's wire `+0x18` is stored minus 1 on load and written plus 1, so it
  round-trips when the stored value is below the loot count; otherwise it is
  stored as -1 and written as 0 (the number of installed units out of range was
  not measured; ALM-SAVE-255). Wire `+0x08` and `+0x0a` are widened to signed 32 bits and
  written back as 16 bits.
- A player's id is its record ordinal plus 1 and is not stored in the file. A
  unit whose wire owner matched no player is bound on load to the first
  player and is written with owner 1, for a file whose type-5 record precedes
  the unit records (otherwise the default player, built with id 1, is the one
  stored); none of the 12,085 installed units is unmatched. — ALM-SAVE-256
- A loot record whose owner object is not found raises a message and is moved
  to unit 1. Its position is written from the owner object, not from the loaded
  position; the owner is written as 0 for one object class. Of the installed
  loot (180 EN, 176 RU), 43 records per root have an owner equal to a unit id
  and the same position as that unit; 137 EN and 133 RU have wire owner 0; at load
  each loot whose `+0x3c` is 0 when the fix-up loop runs, up to the count
  `[ebp-0x18]`, gets an object of the class with program text `CBackPack` at its
  loaded position, so the save lookup finds it and writes owner 0 and that
  position (Medium: the loop bound, the loop body and the unit link `+0xd6` are
  unread, and the counts are of the wire owner). A loot referenced
  by a unit takes that unit's id; a unit's `+0xd6` at or above the loot count is
  written as 0. — ALM-SAVE-255, ALM-SAVE-271
- Type-9 entries with `+8` non-zero, `+4` zero and the 16-bit `+0x16` zero are
  removed from the list by `L14474` and not written; the entry that moves into
  the removed slot is written but not added to the record size. No installed
  entry meets the predicate (0 of 792 EN and 334 RU). — ALM-SAVE-254
- On load, a type-9 entry with `+8` and `+0xc` non-zero and the word `+0x10`
  equal to 4 is dropped when `L15129` returns 0; whether it does for any of
  the 206 EN and 60 RU installed entries meeting the condition is bounded: the
  helper keeps an entry only if an object exists whose rectangle holds the cell
  and whose key is non-zero. 48 EN and 11 RU entries sit on a type-4 record's
  cell, but none is proved kept (the reader may have made no object for the
  record, its key may be 0, its footprint is unread); 158 EN and 49 RU stay
  undecided because the footprint tables are filled at runtime; none is proved
  dropped. — ALM-SAVE-253, ALM-SAVE-260
- A type-4 record is 20 bytes, or 28 when its dword at `+8` is `0x21`. The
  reader creates no structure for a cell of 7 or less or beyond the map bounds,
  and stores a position of `256 * cell + 0x80`. All 71 installed type-4
  payloads parse exactly this way. — ALM-SAVE-262
- A plain type-4 record (kind other than `0x21`) makes an object whose word
  `+4` and hash key are the wire id word, whose position is
  `256 * cell + 0x80`, whose owner is the player of ordinal I (the first player
  when I is above the player count) and whose kind is the low byte
  sign-extended; a repeated key replaces the earlier object in the hash and the
  body read has no test that refuses id 0. A record of kind `0x21` makes an
  object of the class with program text `CVerticalWoodenBridge` (Medium) whose id is the
  smallest id from 1 that no object holds and whose key is the first free key
  from 1; the wire id word is not used and the two extra dwords are its extent
  (the getter reading is Medium). — ALM-SAVE-268, ALM-SAVE-269
- The writer emits the structures in ascending key order (key 0 last) from the
  objects: X, Y, kind, A, the owner's id and the object's own `+4`, and the
  extent of a bridge (Medium). A plain record therefore round-trips (Medium) when its object was
  made; the footprint-table rejection and the order of the installed records
  are not measured. All 3,133 EN and 1,680 RU plain records have a non-zero,
  distinct id, an owner within the player count and a position on a cell
  centre. — ALM-SAVE-270, ALM-SAVE-268
- After the records the reader builds a default player when none exists and an
  object of class text `CBackPack` for each loot whose `+0x3c` is 0 (Medium); the
  two buffers it allocates after the read are one byte per cell from the height
  bytes and, in the code read, touch no record (Medium: their callees are
  unread). — ALM-SAVE-272,
  ALM-SAVE-271
- The owner lookup at save returns the first object whose word at `+4` equals
  the loot owner, including 0; no installed unit has id 0 and no plain type-4
  record has wire id 0 (a plain type-4 object's id is that wire id), and the
  137 EN and 133 RU owner-0 loot find the object made at load (Medium, as above).
  — ALM-SAVE-261, ALM-SAVE-268, ALM-SAVE-271
- The reader stores a loaded player's `+4`, `+8`, `+0xc`, `+0x10`, `+0x30` and
  first 16 words from the wire; the array count 0x20, the vtables and the 16
  words after the first 16 are constructor values, and whether the writer's
  emission depends on them is Unknown. The block copy takes its backward path
  when the destination lies inside the source range; its byte tails are unread
  (Medium). Five rel32 call sites of the caster removal exist (the writer and
  four loops that compare a type-9 entry's `+0x12` with an object's `+4` or walk
  a loot's dword array, Medium); a dword scan of the program sections for the editor's dialog handlers found only
  message maps and vtables (`ALM-DLGPLAYERS-286`). The writer emits exactly 16 relation words per player
  (`ALM-ROSTERIO-292`). A name
  listing for `Map Editor*` and `*edit*` to depth 2 finds no map editor program
  in the RU root. — ALM-SAVE-263, ALM-SAVE-264
- Every action node, condition node and trigger of type 7 is requested whole
  from the same offsets by the reader and the writer, so a load and save without
  an edit writes them unchanged (High for the instruction pairs, Medium end to
  end). — ALM-SAVE-252
- The writer `L02295` stores the file version word 990 whatever the loaded
  version. — ALM-SAVE-255

The first two are High. The statements on players, loot and type-9 entries are
High for the reader and writer instructions read and Medium for code outside
the windows read. The read primitive behind the record reads is a `ReadFile` call into the
destination buffer (ALM-SAVE-272); the footprint tables and any other writer of
a map file are Unknown. — ALM-SAVE-239, ALM-SAVE-252, ALM-SAVE-272, ALM-SAVE-273

## Unknowns and compatibility

| Field or path | Limit |
|---|---|
| Header dataSize | Installed `4*W*H+72`; ignored by the loader; writer accounting of 72 Unknown |
| Type0 +0x0c | Editor Starting Time in minutes and its color consumer are established; complete ROM1 gameplay effect and native transactions remain Unknown — ALM-EDITTIME-205 |
| Type0 scalars +0x10/+0x14/+0x74 | Exact widths/destinations retained; general authoring meanings Unknown |
| Type5 +0x08 | Loaded to `CPlayer+0x0c`; copy use is known, complete gameplay consumer effect Unknown |
| Type6 +0x1c/+0x22/+0x26/+0x28/+0x2a/+0x30/+0x33/+0x34/+0x35 | Known stored fields; no named read in the described original field-consumer set; authoring roles Unknown |
| Trigger +0x40 | 64 raw bytes; no meaningful text grammar established |
| Trigger Target_Item | Authored values exist; complete value domain/meaning Unknown |
| Type9 +0x00 | Exact tag width known; meaning Unknown |
| Type0 text slots | Fixed layout; complete trigger-to-slot relation Unknown |
| Editor reader | Separate reader; complete acceptance contract Unknown |

— ALM-HDR-001, ALM-META-091, ALM-META-092, ALM-CORP-093,
ALM-SCALAR-087, ALM-SCALAR-088, ALM-SCALAR-089, ALM-CPLAYER-090,
ALM-TAILDIR-081, ALM-TAILU16-082, ALM-TAILVER-084,
UNIT-PLACESKILL-086, ALM-TRIG-047 (amended: opaque-region gloss withdrawn), ALM-TRIG-049, ALM-TRIG-050

The placeable-definition source is [Data.bin](../databin/format.md), including
footprints and actor fields. The earlier unknown source-file clause is
resolved by DAT-LOC-001 and DAT-ACT-006 (unconditional Humans typeID clause
partially retracted; database lookup retained).

The first three rows of campaign maps 131 and 150 contain type 3 residue,
including 64 codes whose `c-1` exceeds the 82-entry object array. Bounds-check
all object codes; an in-range code in those rows does not establish a valid authored placement. — ALM-CLS-051

The preserved RU loose `Horror.alm` declares four records and has no type 4 or
type 6 payload despite nonzero metadata counts. Decode no placements from
those absent records; native runtime continuation remains Unknown. A Building's
rectangle, selected presence-mask cells and successfully attached cells are
distinct quantities. — ALM-CORP-060, UNIT-AREAPOP-075, UNIT-STRUCTCELL-070

## Reader and writer distinctions

The main and the editor readers test file+0x10 as an unsigned word. Their payload
read thresholds are 951 (two Unit words),985 (one Unit byte),987 (one Unit
dword),989 (one type8 dword) and 990 (one type9 dword). These tests describe
read programs, not successful complete maps at every value. — ALM-HEADER-097

At 1000, mainR0467 makes no record-header read and preserves its destination.
EditorL02294 reads wire words 0..4 into header dwords2,4,0,3,1 respectively.
Other values select20-byte reads. The browser and landscape helper instead
peeks 8 bytes, rewinds 8, and requests hdrLen bytes without a version test.
Short or oversized headers have no new safety guarantee. — ALM-HEADER-097

Record-header final-word mutations in the selected main/landscape prefixes
do not alter their read program or scalar copy. No numeric interpretation or
equality enforcement was identified there. Earlier f32/enum wording in
ALM-META-009, ALM-META-027, ALM-FRAME-031 and ALM-SEC-003 is partially
withdrawn for this word only. — ALM-HEADER-098

Selected editor writer L02295 emits file-word 990 and dataSize 4*W*H+72.
It submits a 20-byte record header whose last four bytes come from a frame
slot with no local initializer or input-header copy. The incoming frame bytes
and a native editor round-trip remain Unknown. — ALM-WRITER-099

The wrapper clamps reads to length-cursor and advances its cursor before the
underlying read. A successful-short lower read can leave old destination
bytes while the wrapper has advanced further. Full-width field storage
therefore requires a complete lower read. The loose-file receiver is resolved;
archive container+34 and later aliases remain open. — ALM-STREAM-100

The preserved corpus contains 72 maps and 714 record headers. All file words
are 990; all 72 maps have equal raw record words internally. This equality is
a corpus property. — ALM-CENSUS-101

## Count provenance, header residue and closure boundary

In the primary reader, metadata+0x28 first occupies EBP-0x50. Each type-7
internal Read4 then writes the same slot before its own signed loop comparison.
With complete reads, only the respective internal count determines the node,
condition or trigger population. With EOF/short reads, metadata or a preceding
count can survive in that slot. Do not infer a wire field's meaning from such
residual bytes. — ALM-COUNT-195, ALM-STALE-196

The ordinary primary record-header read places its final four bytes at
EBP-0x24. A partial next header can preserve the previous word's unread suffix;
the version-1000 helper preserves all of it without any read. The browser/light
helper instead uses hdrLen after an eight-byte peek. None of these transfer
mechanics establishes a numeric representation for the final word.
— ALM-HDRREAD-197

The primary upper version gate is unsigned <=1001, after unsigned recordCount>=3.
The browser's header predicate instead requires unsigned count>=2 without a
version check there. The light reader walks to EOF and treats count<2 as an
advisory message rather than an early return. Primary payload thresholds are
unsigned >=951, >=985, >=987, >=989 and >=990; evaluating an isolated window
above 1001 does not imply primary whole-file admission. — ALM-READERS-199

The editor interprets metadata+0x0c as Starting Time in minutes and uses it
in its color pipeline. The identified normal game M/P paths do not read that
scalar. Record-header+0x10 is locally unconsumed in the named readers. These
are separate semantic results; neither native acceptance nor universal
all-image non-use follows. — ALM-FRONTIER-200, ALM-EDITTIME-205, ALM-HDRLIFE-206

The complete editor header helper uses five Read4 calls at version1000,
placing wire words0..4 at destination dwords2,4,0,3,1. Other tested values
use one Read20. The helper does not validate returned counts before returning;
short lower transfers retain old destination bytes. The last wire word is
therefore at runtime header+4 at 1000, not+0x10. Isolated helper controls do
not imply whole-file version admission or a numeric field meaning.
— ALM-EDITORHDR-202

The post-load P route is bound for the constructor-produced application,
frame, view and child families. Its identified view aliases and P getters
use dimensions, planes and angle, with no scalar reader. Nonzero app+0x20
still overrides the primary frame, and the host fallback is not a primary-frame
identity proof. Native selection and external member replacement remain
unobserved. — ALM-FRONTIER-200

The complete named reader lifetimes include the incidental header address left
in ESI and its private stack saves. The ordinary wire word is locally
unconsumed; version1000 moves the editor destination to header+4, while the
primary performs no header transfer. Original stream wrappers forward the
destination without retaining it. The next complete overwrite or normal
frame release ends local ownership. The editor writer's last word has no
defined initializer on the shown path. These conclusions assume ordinary ABI,
DF=0, constructor-produced receivers and in-bounds buffers. They establish
neither a numeric type nor native I/O values. — ALM-HDRLIFE-206
