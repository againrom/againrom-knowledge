<a id="dat-databin--the-placeable-definition-database--specification"></a>

# Data.bin definition database

`world.res:data/data.bin` serializes eleven definition collections in eight
groups. A loose `World\Data\Data.bin` takes precedence. If both sources are
absent, the original loads eleven semicolon-separated CSV tables from that
directory and writes Data.bin, including their column titles. — DAT-LOC-001

All integer fields are little-endian. Arrays are counted; groups C–H reserve
collection index 0 and omit that entry from the stream. — DAT-GRAM-003

<a id="at-a-glance"></a>

## Structure

Eight class-groups, each `[column-title string array][1..3 collections]`; a
collection is `[u32 count][entries]`. Groups C–H write entries `1..count-1` — entry
0 is reserved in these C–H definition collections. Their serialized
definition indices begin at 1; this is not an indexing rule for A/B or
for every consuming lookup.

| Group | Collections | Serialized definitions in the installed file | Payload after name |
|---|---|---|---|
| A | Shapes, Materials | 5,16 | Nine binary64 values (72 bytes) |
| B | Magic | 50 | Parameter array |
| C | Armors, Shields, Weapons | 30,9,27 | Parameter array, ten raw bytes, second dword array |
| D | MagicItems | 49 | Parameter array, one raw byte, CString |
| E | Units | 118 (56 parameterized) | Parameter array, two CStrings |
| F | Humans | 215 (210 parameterized) | Parameter array, ten CStrings |
| G | Buildings | 66 | Parameter array |
| H | Spells | 28 | Parameter array, CString |

For A/B the collection count equals the serialized definition count. For
C–H it includes the reserved entry 0, so it is one greater than the numbers
shown above. Read that stored count rather than using the installed numbers
as constants. Installed title counts for A–H are 11,30,18,4,57,28,9,24.
The ten group-C raw bytes are five u16 material masks, one per Shapes row.
— DAT-GRAM-003, DAT-MATMASK-020

| Primitive | Encoding |
|---|---|
| CString | u8 byte length;0xff selects a following u16 length; then the string bytes |
| Title array | u16 count followed by that many CStrings |
| Parameter/second dword array | u16 count followed by that many u32 values |
| Collection count | u32 |

These are the established bounded primitive forms. Broader MFC length
escapes are outside this reference. Integers and binary64 values are
little-endian. — DAT-GRAM-003

## The schema law

Each group shares one title array among its collections. Shapes and Materials
have nine binary64 values per entry and no parameter array. Other named
parameter arrays carry explicit counts. The stored Units and Humans parameter
arrays contain 55 and 26 values respectively in their parameterized rows.
— DAT-SCHEMA-007

Param slot `i` is CSV column `i+1` (column 0 is the entry name); the trailing
"equipment" column, when the class has string slots, is `,`-separated with `{...}`
groups and fills the entry's extra strings instead. **An empty cell is stored as −1,
and the named parameter-streaming readers skip the store on −1**. In those
readers, −1 preserves the constructor default; other consumers can use −1
as a literal value, including the document-item price below.
— DAT-SCHEMA-004, DAT-ACT-006 (unconditional Humans typeID clause partially
retracted; −1 law retained)

## The placement-facing tables

- **Buildings** — index = `structures.reg` `ID` (1..66; `DAT-BLD-005`). Slots:
  `0 sizeX, 1 sizeY` (footprint in tiles — overridden by the ALM type-4 `kind==0x21`
  extension), `2 scanRange → obj+0x48`, `3 healthMax → obj+0x44`, `4 Passability`,
  `5 BuildingPresent` (the two sub-cell masks, `TERR-STRUCT-070` (amended, partially retracted) and
  `TERR-STRUCT-078`; a set Passability bit blocks). Resolved from an ALM type-4 `kind` by
  direct 1-based subscript (`ALM-CLS-036`/`ALM-CLS-053`), or by name
  (`R0694`, backwards, `"Invalid building %s created"`).
- **Humans** — searched on slot `0x10 typeID` for type-6 keys `< 0x40 ∉ {26,27}`,
  and on slot `0x18 serverID` for the type-6 `+0x10` override id. Streamed at spawn
  (`R0657`): stats/HP/mana/speed/skills, `0x15 TokenSize → actor+0x49`,
  `0x16 MovementType → actor+0x4a`; ten equipment strings resolved by name against
  Weapons/Shields/Armors; `0x19 knownSpells` a bitmask. Slot `0x10` streams to
  `actor+0x0e`. `R0656` overwrites it with `gender+0x21/+0x23` only when the
  constructor-mode argument is non-zero. ALM definition-id and explicit typeID arms
  pass zero; the npc arm passes the exact `Hero` flag result (`PARTY-M20-031`). A map
  Human can therefore retain its authored table typeID.
- **Units** — searched on slots `0x1d typeID` + `0x1e face` for all other type-6
  keys. Streamed at spawn (`R0180`): slots **0–37 only**, in order — stats,
  the two regeneration periods, the damage pair through the `attackKind` switch,
  protections, resists, `0x1f tokenSize → +0x49`, `0x20 movementType → +0x4a`
  (shipped values 2 = Ghost/Bee, 3 = Bat_Sonic/Dragon — the `0x44`/`0x82` block
  masks). The **treasure and spell slots (38–54) are not streamed**: the spell pairs
  are read by `R0184` into the spellbook and the order block
  (`UNIT-SPELL-007`) and the treasure columns by the kill payout (`HERO-KILL-027`).
  Complete slot → actor-field map with widths and instruction addresses:
  [`formats/unit`](../unit/format.md), `UNIT-STREAM-001`.

## Consumers beyond placement

Field consumers and their formulas are defined in the following references.

| Collection | Established use and authority |
|---|---|
| Shapes / Materials | The nine-double records feed item tier/material scaling and damage factors: `ITEM-LADDER-019`, `ITEM-DMGFACT-020`, [item specification](../item/format.md). The old ladder in `ITEM-SCALE-017` is retracted in favour of `ITEM-LADDER-019` |
| Weapons | Item scaling and equipped combat inputs: amended `ITEM-SCALE-017`, `HERO-EQUIP-017`; see [item](../item/format.md) with [hero](../hero/format.md) |
| Armors / Shields | Slot selection, armor/shield fields and the suitability mask: `ITEM-ARMSLOT-031`, `ITEM-ARMFILL-032`, partially retracted `ITEM-SUIT-035`. Its mask/title remain supported; its display-only consumer limit is superseded by `ITEM-WEAR-055`/`ITEM-WEAR-057`: the client refuses unsuitable equipment drops, while the raw equip path remains unrestricted. Group C's ten raw bytes are five u16 material masks, one per Shapes row, tested by the shop candidate loop (`DAT-MATMASK-020`) |
| Magic | The eligible enchantment pool and selection stages are `SHOP-EFFPOOL-061`, `SHOP-EFFWEIGHT-062`, `SHOP-EFFPAY-063`, `SHOP-EFFRANGE-064`, `SHOP-EFFCAST-065`, `SHOP-EFFPRICE-066`, `SHOP-EFFORDER-067`, `SHOP-EFFRETRY-068`, `SHOP-EFFCAP-069`, `SHOP-EFFBASE-070` and `SHOP-EFFALT-071`. `SHOP-EFFPRICE-066` is partially retracted for its nonlinear formulas; aggregation, retained-cast additivity, upper clamp and subsequent budget reads stand. See [shop](../shop/format.md) |
| Spells | Definition binding, cast inputs and effect construction are `MAGIC-SPELL-001`, `MAGIC-CAST-003`, `MAGIC-EFFECT-015`; individual spell consumers remain in the [magic specification](../magic/format.md)  (Delivery/timing clause narrowed by MAGIC-CASTCLOCK-171.) |
| MagicItems | Initial signed price assignment and Scroll/Book exceptions derived from the item's display string are amended `ITEM-MAGVAL-090` and `ITEM-VALUE-115`. Book uses the first Effect's spell id, not the last. The later -1 descriptor arm is `ITEM-WEAR-058`; the extra wire byte remains unnamed |

The equipment-cell parser is read: `ITEM-NAMEPARSE-040` specifies name/tier/material
parsing, with the enchantment grammar and effect construction refined by
`ITEM-EFFGRAM-070`, `ITEM-EFFPOP-071`, `ITEM-EFFOBJ-072` and `ITEM-EFFMODE-073`.

<a id="buildings-presence-mask-consumer"></a>

## Coverage boundaries

The second group-C dword array's role and group D's extra raw byte's meaning
remain Unknown. Both have known wire extents and must be retained. The Buildings
`Start ID`/`Tiles` titles have no stored numeric parameter (`DAT-SCHEMA-007`),
so they are not two missing numeric payload fields. Full server-to-client
creation-message semantics are a separate session contract. Human type-ID
streaming retains the conditional constructor overwrite described above
(`DAT-ACT-006` and `DAT-HUMANS-008`, each with its unconditional-overwrite
clause partially retracted; `PARTY-M20-030`, `PARTY-M20-031`).

EN and RU Data.bin payloads differ. Units entries match; Humans parameters
can differ. Preserve the values of the selected definition source.
— DAT-HUMANS-008 (unconditional slot-16 overwrite clause partially retracted;
23-slot map retained)

The Buildings presence mask is consumed as `(row*width+col)&31`, with byte width/height and a
32-bit mask; it is not an arbitrary-size bitset. Registration visits only its set positions and
can stop at the first occupied Building slot. The declared rectangle therefore does not prove
that all selected cells attached (`UNIT-STRUCTCELL-070`). Both preserved Data.bin inputs yield
66 identical numeric Buildings rows, but the corresponding installed placement populations
differ (`UNIT-AREAPOP-075`).

## Campaign-start document item

`MagicItems[28]` is named `Quest Documents` on both roots. Its parameter array is `[-1,5]` under
the titles `Magic Items`, `Price`, `weight`, `Effects`; price is therefore -1. Its packed item code
is `0x0e1c`.

The `Humans` table cannot place this MagicItems row in starting equipment. Cell 0 constructs a
Weapon, cell 1 a Shield, and cells 2..9 Armor. The installed starting-equipment cells contain no `Quest` entry. — DAT-DOC-021

## Read and write sequence

1. Process groups A through H in the fixed order above.
2. For each group, read its u16 title count and length-prefixed strings.
3. Read each collection's u32 stored count. Read all entries for A/B and
   entries `1..count-1` for C–H; preserve the reserved null slot in memory.
4. Read the entry name and its group's exact payload. The group-C second
   dword array has the same u16-count/u32-element framing as a parameter array.
5. Resolve parameters through the relevant definition consumer. A stored -1
   preserves the constructor default on the named streaming paths.

Emission uses the same order and count rules. Emit the shared titles once per
group, each collection count once, then its entries. Keep the ten-byte group-C
material-mask block, its second counted array, and the group-D extra byte even
when their interpretation is incomplete. No magic, section tags or per-entry
size fields are inserted by this grammar. Array and CString lengths must fit
the established primitive form; broader MFC encodings are not specified here.
— DAT-GRAM-003, DAT-SCHEMA-004, DAT-SCHEMA-007, DAT-MATMASK-020
