# DAT — `Data.bin`, the placeable-definition database

Claims about `world.res:data/data.bin`, the database at runtime `L02110` that
resolves every ALM placement to its parameters (`ALM-CLS-036`, `ALM-CLS-038`).
Spec: [`formats/databin/format.md`](../formats/databin/format.md). Format of this
file: [registry.md](registry.md).

## Loading and wire grammar

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DAT-LOC-001 | The placeable-definition database at `L02110` loads from `World\Data\Data.bin`, shipped as the archive node `world.res:data/data.bin` (88 327 B). | High | ● active | [EXP-0049](../experiments/EXP-0049-placeable-db/) |
| DAT-OBJ-002 | The database holds eleven by-value collections; every entry begins `{vptr, CString name, CDWordArray params}` and its size depends on the element class. | High | ● active (amended) | [EXP-0049](../experiments/EXP-0049-placeable-db/) |
| DAT-GRAM-003 | `Data.bin` is eight serialized groups of column titles and entries, and this grammar tiles the shipped file with 0 residue. | High | ● active | [EXP-0049](../experiments/EXP-0049-placeable-db/) |

### DAT-LOC-001

- `R0679(this=L02110, "World\Data\")`, called at `L03285` and
  `L03286` with that literal prefix, calls `R0680`. It opens a loose
  `<path>Data.bin`, else the RES node `"World\Data\Data.bin"`, prints
  `"Loading static data."`, wraps a load-mode `CArchive` and runs the database
  Serialize `R0681`.
- When neither exists it prints `"StaticData files not found"` and
  `"Parsing .csv files"`, and `R0682` parses eleven CSV tables:
  `Spells Armors Materials Shapes Magic Weapons Shields "Magic Items" Units Humans Buildings`
  (`.csv`, `;`-separated, `rem` rows and `goto` skipped). It then prints
  `"Writing new .bin file"` and `R0683` rewrites `Data.bin`. No CSV
  ships.
- No `.reg` file is read on the path: the database is not a registry
  projection.
- A guard at `this+0xf0` makes the load run once.

**Confidence.** High. Every hop is a named call next to its own string.
`EnumRefs refto:L02110` on the repaired function table returns 45 hits in 25
owners, 0 in orphan or undisassembled code; the two `World\Data\` callers are
the only owners that reach `R0679`.

### DAT-OBJ-002

- Collections sit at `L02110 + {0x00 Materials, 0x14 Shapes, 0x28 Shields,
  0x3c Armors, 0x50 Weapons, 0x64 MagicItems, 0x78 Magic, 0x8c Units,
  0xa0 Humans, 0xc8 Buildings, 0xdc Spells}`. Each collection object is 0x14 B.
  The CSV parser's per-table advance of the collection pointer by that offset and the Serialize use the same
  displacements. There are eight element classes.
- Every entry begins `{vptr @+0x00, CString name @+0x04, CDWordArray params @+0x08}`.
  The array's `pData` sits at entry `+0x0c`: this is the u32 parameter array at
  `+0x0c` that EXP-0037 reported.
- Entry size by class: buildings and magic 0x1c (the bare base); spells 0x20
  (+1 CString); units and humans 0x30 (+ a CStringArray of 2 or 10 strings);
  armors, shields and weapons 0x3c (+10 raw bytes + a second dword array);
  magic items 0x40; materials and shapes 0x68 (name + 9 doubles, no params).
  The constructor, vtable and `vt+8` Serialize of each class are tabulated in
  the evidence excerpt §4.
- Elements are stored by value (`elementAt = base + i×size`), so a definition
  handle passed on is `entry+4`, the name field.

**Confidence.** High. Each size is the constructor's own stride
(`R0684`-family SetSize `×0x30/0x1c/0x20/0x3c/0x40/0x68`), and each
Serialize body was decompiled. The collection-to-offset map appears twice: in
`R0681`, and independently in `R0682`'s per-CSV arms.

**Amended.** EXP-0037's "each entry 0x30 bytes" described only the two
collections it saw; the building entries at `L02102` are 0x1c.

### DAT-GRAM-003

- Eight groups in fixed order, the serialize order of `DAT-OBJ-002`:
  A Shapes+Materials, B Magic, C Armors+Shields+Weapons, D MagicItems, E Units,
  F Humans, G Buildings, H Spells.
- Each group starts with a CStringArray of the CSV column titles: u16 count and
  length-prefixed CStrings, MFC `CStringArray::Serialize` on a per-group static
  (`L03287/L03288/L03289/L03290/L03291/L03292/L03293/L03294`).
- Then each collection writes a u32 entry count (`R0685`/`R0686`)
  and its entries through the element's virtual `vt+8`: `[CString name]` and a
  per-class payload. A param array is a u16 count followed by raw u32s.
- Groups C–H skip entry 0: it is allocated, never serialized and never matched,
  because the loops run from 1. Every consumer index is therefore 1-based, and
  `R0487` returns `count−1`.
- Shipped counts: 5 shapes, 16 materials, 50 magic, 30 armors, 9 shields,
  27 weapons, 49 magic items, 118 units (56 parameterised), 215 humans (210),
  66 buildings, 28 spells. The walk consumes 88 327 of 88 327 bytes with 0
  residue (`tools/placedb -mode parse`).

**Confidence.** High. The grammar is transcribed from `R0681` and the
eight Serialize bodies, and the tiling rejects the live rivals: u16 collection
counts break 0x66 bytes in, u32 param counts break in the first parameterised
entry, and serializing groups C–H from 0 over-reads. Exactly one width and base
assignment closes, and it is the one the instructions spell.

## Parameter slots and actor creation

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DAT-SCHEMA-004 | Column titles name the parameter slots: parameter `i` is CSV column `i+1`, and an empty cell stores −1. | High / Medium | ● active | [EXP-0049](../experiments/EXP-0049-placeable-db/) |
| DAT-SCHEMA-007 | Every collection's titles map to slot numbers fixed by four instruments; `Data.bin` Units carries a unit's simulation numbers and `units.reg` its drawable. | High / Medium | ● active | [EXP-0072](../experiments/EXP-0072-unit-stats/) |
| DAT-HUMANS-008 | The Humans spawn streamer consumes 23 slots in order, and slot 16 reaches `actor+0x0e`. | High | ● active (amended, partially retracted) | [EXP-0104](../experiments/EXP-0104-combat-columns/), [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |
| DAT-HUMANS-009 | Humans rows carry no absorption, protection, resistance, physical-damage or attack-kind column; those exist only in Units. | High / Medium | ● active | [EXP-0104](../experiments/EXP-0104-combat-columns/) |
| DAT-ACT-006 | Actors are created from the database through a Humans arm and a Units arm, and a −1 parameter keeps the constructor default. | High | ● active (amended, partially retracted) | [EXP-0049](../experiments/EXP-0049-placeable-db/), [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |

### DAT-SCHEMA-004

- The extra string fields (equipment lists) are parsed from the row's last cell
  (`R0687`, `,`-split with `{...}` groups). Numeric cells go to
  `params[col−1]` (`R0688`). An empty cell stores −1.
- For the two tables the simulation consumes, the spawn streamers
  `R0657` (humans) and `R0180` (units) read the slots in order
  into named actor fields; the full maps are in the evidence excerpt §7.
  - Humans: `0x10 typeID` (the `ALM-CLS-038` search key), `0x18 serverID` (the
    type-6 `+0x10` id, `R0498`), `0x15 TokenSize → actor+0x49`,
    `0x16 MovementType → actor+0x4a`.
  - Units: `0x1d typeID` and `0x1e face` (the second search arm),
    `0x1f tokenSize → +0x49`, `0x20 movementType → +0x4a`.
  - Named actor fields also include HP → `+0x96`, speed → `+0x8c` and
    rotationSpeed → the mover's `+0xa`.
- Buildings: `0 sizeX`, `1 sizeY` (the `ALM-CLS-036` footprint pair),
  `2 scanRange → obj+0x48`, `3 healthMax → obj+0x44`, `4 Passability`,
  `5 BuildingPresent`. The two trailing titles (`Start ID`, `Tiles`) have no
  stored parameter: 6 params on 66/66 rows.

**Confidence.** High for the streamed slots and buildings params 0–3: ordered
reads into fields whose meaning other experiments pinned. Medium for the
title-only slots (`Passability`, `BuildingPresent` and the non-streamed tables):
the shipped schema names them, and their consumers are unread.

### DAT-SCHEMA-007

- `evidence/databin-slots.csv` (EXP-0072) lists each title of all eleven
  collections with its slot index under the `DAT-SCHEMA-004` law. For Units and
  Humans it adds the actor field, store width and instruction address the spawn
  streamer writes.
- A group's title array is one array shared by its sibling collections:
  `Armors + Shields + Weapons` share 18 titles and `Shapes + Materials` share 11.
  The `Shapes`/`Materials` element kind has no param array (its record is 9
  doubles, `R0689`), so those 11 are titles, not slots.
- The param arrays are full: all 56 parameterised Units rows carry exactly 55
  values and all 210 Humans rows exactly 26. Every title except the trailing
  string column has a stored cell on every row, in both shipped roots.
- Four instruments fix the slot numbering independently of the shipped titles:
  the two streamers' consumption order; `R0690`, which reads the same two
  slots as the displacements `+0x74`/`+0x78` (`0x1d·4`, `0x1e·4`);
  `R0691`/`R0692`, which fetch slot `0x21`/`0x17` on demand and
  land on each table's `dyingTime`; and `HERO-EQUIP-017`'s weapon fill, which
  reads slots 3/6/7/8/9 as `weight`/`@.physicalMin`/`@.physicalMax`/`@.toHit`/`#.deIrnce`.
- `units.reg` is the client drawable's record (`REG-UNITS-049`), and the
  simulation module never reads that class array (`REG-UNITS-061`); `Data.bin`
  Units is the simulation actor's. Three pairs that look like duplicates sit on
  different objects with different consumers:
  - `units.reg` `TileSize` is `CUnit`'s `vt+0x20`; `tokenSize` is the actor's
    `+0x49` footprint (`TERR-MOVE-054`).
  - `AttackDelay`/`ShootDelay` pace the drawable;
    `attackChargeTime`/`attackRelaxTime` are the simulation countdown
    (`HERO-CADENCE-023`).
  - `Dying`/`DyingPhases` name a corpse sheet (`REG-UNITS-050`); `dyingTime` is
    how long the corpse holds its cells.
- The join is `units.reg ID == Data.bin typeID` (`ALM-CLS-052`'s resolution
  rule, re-measured 33/34 by `tools/placedb -mode verify`). The tier `face`
  exists only on the `Data.bin` side.
- `gameversions/ru/WORLD.RES` is a different file (`8ba0d479…` vs `7b44c526…`)
  with a different `Data.bin`. Its Units collection agrees on every name, every
  equipment string and all 56×55 parameters: 0 differences. Humans differ on 16
  parameters over 6 slots: one row's skills, `typeID` on 8 rows and `face` on 4.

**Confidence.** High for the slot law, which is `DAT-SCHEMA-004`'s; this claim
adds two displacement-bearing instructions that carry the numbering without any
title. The `units.reg`/`Data.bin` division rests on `REG-UNITS-061`'s
enumeration of the class array's 99 references. Medium, unchanged from
`DAT-SCHEMA-004`, for the title-to-meaning reading of the eight tables whose
consumers are unread.

### DAT-HUMANS-008

- `R0657`, read whole, maps `0..3 → +0x84/+0x86/+0x88/+0x8a`;
  `4 → +0x96`; `5 → +0x9c`; `6 → +0x8c`; `7 → [+0x154]+0x0a`; `8 → +0xa5`;
  `9 → +0xbe`; `10..15 → +0xa8+2i`; `16 → +0x0e`; `17 → +0x4b`; `18 →` a dead
  local; `19/20 → +0x134/+0x135`; `21 → +0x49`; `22 → +0x4a`.
- Its second loop copies base skills for `i=1..5` only.
- Every helper skips a store on −1 and advances either way.
- The constructor later fetches gender slot `0x12`, but writes
  `gender+0x21/+0x23` only for a non-zero constructor-mode argument, so slot 16
  is not always overwritten. Mission 20's zero-mode actors keep the table values
  `0x17` and `0x0a` in original saves (`PARTY-M20-030`).

**Confidence.** High for the complete streamer map and the conditional
overwrite; the original-save words separate it from the unconditional reading.

**Amended.** The overwrite of slot 16 was first stated as unconditional; that
clause is retracted, and the 23-slot map and widths stand
([`retracted.md`](retracted.md), EXP-0192).

### DAT-HUMANS-009

- From `DAT-HUMANS-008`'s complete map, the Humans arm streams no
  `absorbtion`, no `prot Fire..Astral`, no `res.Blade..res.Shooting`, no
  `physicalMin`/`physicalMax` and no `attackKind`. These five column families
  exist only in `Units`; a consumer must not invent them for a Humans-arm
  placement.
- A human's damage comes from `HERO-COMBAT-011`'s derivation,
  `ftol(1.1^body / 20)` into both `+0xb4` and `+0xb5`, so the bare roll is
  `[d, 2d]`, plus whatever `Weapon::Equip` adds. Its absorption comes from
  armour alone.
- The one combat column the arm streams is `Defence`, slot 9 → `+0xbe`
  (`L03295`). `R0280`'s `memset(actor+0xbe, 0, 0x16)`
  (`HERO-ORDER-014`) overwrites it and the eleven bytes after it each time the
  recompute runs, replacing it with `reaction / 3`.
- `toHit` is not a column either: `+0xa6 ← +0xa8` makes it the `Skill.General`
  column until the same recompute replaces it with
  `ftol((1.1^body + 1.1^reaction) / 5)`.
- Corpus, both roots: 215 rows, 210 parameterised, 26 parameters each. Slots
  0…22 are streamed; 23…25 (`DyingTime`, `serverID`, `knownSpells`) are fetched
  on demand.
- Example, the row mission 1 needs: `serverID` 512, `M10_Brigands`, health 15,
  Body 5, Reaction 20, `Defence` 4, every skill 0, equipment `Wood Club` and
  `Leather Boots`, and no damage number anywhere in the row (`UNIT-M10-018`).

**Confidence.** High for the absences: the map is the whole routine, which has
no other store, so a column the streamer does not read is not stored. High for
the `memset`, a cited instruction in a routine read end to end. Medium that the
`Defence` column is dead in practice on an `.alm`-placed human.

**Unknown.** Whether `R0280` runs at spawn on the Humans arm. Every
equip path can reach it through `R0451`.

### DAT-ACT-006

- Humans: `R0497(new(0x1e8), entry+4=name, …)` calls `R0656`,
  which walks Humans from 1 by name, parks the definition at `actor+0x3c`,
  streams the params, resolves ten equipment strings and expands `knownSpells`.
- The stream writes slot 16 `typeID` to `actor+0x0e`. Only a non-zero
  constructor-mode argument makes `L03163` overwrite that word with
  `gender+0x21` or `gender+0x23`. Definition-id and explicit typeID placements
  pass zero; npc placements pass the exact `Hero` flag lookup (`PARTY-M20-031`).
- Units: the spawner's Units arm constructs through `R0501` and streams
  `entry+8`.
- Every stream helper reads `param[cursor++]` and skips its store when the value
  is −1, so an empty CSV cell keeps the constructor default.
- Buildings by name go through `R0693` → `R0694(L02102, name)`.

**Confidence.** High. The conditional, all three placement arguments, the
stream calls and the −1 gates are named instructions; mission 20 supplies saved
counterexamples to the unconditional wording.

**Amended.** The Humans `typeID` assignment was first stated as unconditional;
that clause is retracted, and the database lookup, equipment, spell and −1 laws
stand ([`retracted.md`](retracted.md), EXP-0192).

## Buildings, items and names

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DAT-BLD-005 | The Buildings table is `structures.reg` re-keyed: entry index equals the `structures.reg` `ID`, 1..66. | High / Medium | ● active | [EXP-0049](../experiments/EXP-0049-placeable-db/) |
| DAT-MATMASK-020 | The ten raw bytes of a group-C entry are five `u16` material masks, one per `Shapes` row (tier). | High | ● active | [EXP-0144](../experiments/EXP-0144-item-pictures/) |
| DAT-ITEMNAME-010 | Item row names in `Data.bin` are identical on both roots, so the database neither localises nor stores an item's displayed name. | High | ● active | [EXP-0142](../experiments/EXP-0142-item-names/) |
| DAT-NAMES-016 | A `Data.bin` row name is an engine identifier, never display text, and is the same on both roots. | High | ● active | [EXP-0143](../experiments/EXP-0143-actor-names/) |
| DAT-DOC-021 | `MagicItems[28]` is the document access item's data identity, but no Humans row can author it as starting equipment. | High | ✔ promoted | [EXP-0154](../experiments/EXP-0154-campaign-start/) |

### DAT-BLD-005

- The count is 67 with entry 0 skipped, which gives the `R0486` guard
  `kind ≤ count−1 = 66` and explains `ALM-CLS-036`'s 1-based rule.
- All 3141/3141 shipped type-4 `kind`s land on a named entry.
- `structures.reg`'s 66 IDs land on named entries in the same order. 44/66
  names are byte-identical case-insensitively; the other 22 paraphrase the same
  object (`"Goblin's Hut 1"`→`"Goblin Hut"`, `"Windmill 1"`→`"Mill"`,
  `"Saving Place"`→`"Multiplayer Church"`; full list in `evidence/verify.txt`).
  Two authored vocabularies name one roster.
- Values are committed in `evidence/databin-buildings.csv`, for example
  `Goblin Hut` 3×2, scan 4, HP 300, and `Church` 4×3, HP 30 000.

**Confidence.** High for the index law: the skip-0 grammar, the count guard,
3141/3141 and the ordered name agreement. The rival "ordered by something else"
has no surviving form when 66 ordered rosters agree name for name. Medium for
`Passability` and `BuildingPresent` as footprint cell masks: the value shapes
fit `sizeX×sizeY` bits, but the consumer was not read.

### DAT-MATMASK-020

- `DAT-OBJ-002` records the `Armors`/`Shields`/`Weapons` entry as
  `{vptr, name, params}` plus 10 raw bytes and a second dword array. The ten
  bytes sit at runtime `entry+0x1c`.
- The shop's candidate generator subscripts them by tier and tests one bit per
  material: at `L03296` it loads the 16-bit word at entry offset `0x1c` plus
  twice the tier, then (`L03297`..`L03298`) masks it with a single bit
  1 shifted left by the material index and skips the candidate when the
  result is zero.
- Ten bytes, five tiers, sixteen bits: every bit is addressable and none is
  spare. This is the only per-row constraint in the schema and was its last
  uninterpreted field.
- Corpus, both roots byte-identical: 367 bits set over the 66 parameterised
  rows; by tier `Common 79, Uncommon 97, Rare 80, Very Rare 73, Elven 38`; by
  material `Iron 24 … Bone 0 … None 41`. The five rows with no art, `BareHands`,
  `rem` and the three monster attacks, carry `0x0000` in all five words.
- The bits agree with the shipped `graphics.res` picture set on 5 280 of 5 280
  (row, tier, material) cells with 0 disagreements (`ITEM-PICT-050`).

**Confidence.** High. The six instructions are one read path with the stride,
the subscript and the bit test all named, and the reading is checked against an
independent artefact, the art tree, over the whole cross product.

### DAT-ITEMNAME-010

- All six name-bearing collections were walked on both roots under
  `DAT-GRAM-003`'s grammar: `Shapes` 5, `Materials` 16, `Weapons` 27,
  `Armors` 30, `Shields` 9, `MagicItems` 49. 0 of the 136 stored row names
  differ, and neither root has a high byte in them.
- The two `Data.bin` files do differ, on `Humans` parameters; the difference is
  not in these names.
- These names are the only strings a per-class definition carries. An item
  instance has three further degrees of freedom: its `+0x45` shape byte, its
  `+0x46` material byte and its class. The displayed name's source is
  `ITEM-DISPNAME-036`.

**Confidence.** High. A complete walk of both files with an exact string
comparison; the walk tiles both payloads with 0 residue.

### DAT-NAMES-016

- Measured over both installs (`evidence/databin-names.txt`): `Humans` has 215
  rows and `Units` 118. The name lists are identical `en` to `ru` row for row,
  and no name carries a byte >= 0x80 in either collection on either root. The
  longest name is 23 bytes in `Humans` and 15 in `Units`.
- 0 of 215 `Humans` names begin with `Hero`, and 0 of 215 contain a `.`: the two
  things `R0656` parses out of the name it is handed
  (`SESS-NAMEGATE-033`). The suffix grammar `SESS-HERO-014` reads therefore
  applies only to names the code composes, never to a shipped name.
- A display string would differ between the roots, as `TEXT-STRTAB-023`'s table
  does on 257 of 275 lines.

**Confidence.** High. Both roots, complete collections, name for name, on the
`DAT-GRAM-003` walk that closes with zero residue; the two absences are
exhaustive scans, not samples.

### DAT-DOC-021

- Both roots store row name `Quest Documents`, params `[-1,5]` and titles
  `Magic Items, Price, weight, Effects`, so the price is -1. The packed code is
  `0x0e1c`.
- All 215 stored `Humans` rows contain 909 non-empty equipment cells, with no
  exact or substring `Quest` match.
- The Humans equipment streamer constructs cell 0 as Weapon, cell 1 as Shield
  and cells 2..9 as Armor, so its positional grammar has no MagicItems arm.

**Confidence.** High. A complete both-root parse and the instruction-level
positional dispatcher.
