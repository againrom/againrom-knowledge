# REG — `.reg` class registries

Claims about the `&YA1` registry format that `.reg` nodes and the application
state store share: its framing, 32-byte records, kinds and pool; the `rom.exe`
loader, lookups, typed accessors and writer that read and write it; the class
records the graphics loaders build from `units.reg`, `objects.reg`,
`structures.reg` and `projectiles.reg`; the contents of the other shipped
registries; and the Map Editor's `.ini` catalogues. Spec:
[`formats/reg/format.md`](../formats/reg/format.md). Format of this file:
[registry.md](registry.md).

## Terms

- A registry is a nested `&YA1` payload: a header, a flat array of 32-byte
  records and a trailing pool (`REG-FMT-031`). The install holds 44
  (`REG-LOC-038`).
- The old framing is the reading published before EXP-0032: a `0x20`-byte header
  and records `[A][kind][name:20][C]` from `0x20`. The corrected framing, "the
  EXP-0032 base", reads a `0x18`-byte header and records
  `[value][size][kind][name:16]` (`REG-REC-032`). Under the old framing each key
  was paired with the next key's value. EXP-0032/0033 re-derived this ledger
  from the loader and withdrew what the correction invalidated.
- A record's name sits at file `0x28 + 32i` under both framings (old
  `[A][kind][name:20][C]` from `0x20`, corrected `[name:16] @ 0x18+32i+0x10`).
  Key catalogues and name lists therefore stand everywhere in the repository;
  only value attributions moved. The old `C @ 0x3C + 32i` is the next record's
  `value @ 0x1C + 32(i+1)`.
- A kind is a record's type word, a bitfield (`REG-KIND-033`, `REG-KIND-034`).

EXP-0043 regenerated the name → sprite-path pairing (`REG-ROSTER-052`: the old
arrows are wrong on 34/34 classes, and 14 index outside the `Files` table) and
gave a per-key value table to the 38 registries that had none: `REG-CUT-053`,
`REG-SFX-057`, `REG-NPC-058`, `REG-SCN-059`, `REG-AI-060`. It also measured two
format-shape facts a reader applies: the sorted flag sits on the root alone
(`REG-KEY-054`), and a `kind` belongs to the record, not the key name
(`REG-KIND-056`). `data/map.reg` is the one registry whose value table is not in
this ledger: it is `ALM-TERR-043`, where EXP-0041's re-walk on the `0x18`
framing restored the `Cost`/`Pass` figures byte for byte, so their withdrawal
under `ALM-TERR-015` no longer applies.

## Embedded application-store parser and lookup boundaries

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-099 | The raw YA1 loader `R1389` reads header, records and pool without recognizing, filtering or validating key names. | High / Unknown | ✔ promoted | [EXP-0337](../experiments/EXP-0337-application-store/) |
| REG-100 | Sorted and unsorted name lookups have different case rules: the linear route folds ASCII case over 15 bytes under no locale, the sorted route compares bytes through NUL. | High / Unknown | ✔ promoted | [EXP-0337](../experiments/EXP-0337-application-store/) |
| REG-101 | Name matching enforces neither node kinds nor duplicate uniqueness, and the integer getter `R0452` accepts any hit with `kind & 0x0e == 2`. | High | ✔ promoted | [EXP-0337](../experiments/EXP-0337-application-store/) |
| REG-102 | The raw registry writer `R1390` preserves the tested unrelated entries, while its sort can change later lookup behavior. | High | ✔ promoted | [EXP-0337](../experiments/EXP-0337-application-store/) |

### REG-099

- Complete `R1389` reads six header dwords, bulk-reads `R*32` records, reads
  the pool length and bulk-reads that pool.
- It checks the signature and the allocation outcomes. It does not enumerate or
  validate names, sorted order, child kinds, duplicate names or child-block
  validity.
- Under successful synthetic allocation/stream services, the complete input
  records and pool remain byte-exact, including unrelated roots and leaves
  before, between and after known names.
- All 24 noncolliding additions across four root/leaf sorted-flag combinations
  leave nine selected integer lookups equal to their baselines when the
  advertised order is true.
- Evidence: `registry_load`; V001..V028 and the boundary vectors.

**Confidence.** High for this complete raw body and conditional boundary,
corroborated by matching EN/RU images and original-instruction vectors.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Unknown.** OS allocation/stream failures, arbitrary malformed extents and
application LOAD acceptance.

### REG-100

- With parent bit 4 clear, `R0505` scans with
  `R1391(query,child+0x10,15)`. Under its no-locale condition, ASCII A..Z
  fold and high bytes do not.
- With bit 4 set and a positive count, the lookup copies at most 15 query bytes
  plus NUL into a temporary record and calls `R1392` with comparator
  `L07393`. The complete comparator compares unsigned bytes through NUL with
  no case fold.
- In the paired controls, root and leaf case variants match only on their
  unsorted routes.
- A 15-byte stored name matches a longer query sharing that prefix on both
  routes. A raw 16-byte non-NUL name matches the linear route but misses the
  sorted truncated-key route. Empty names match on both tested routes.
- Node insertion `R1393` clamps to 15 bytes plus NUL, sets bit 28 for input
  lengths above 15 and clears the parent sorted bit. Typed setters may
  subsequently replace the kind, as `L07394` does.
- This narrows `REG-REC-032` and `REG-KEY-054`.
- Evidence: lookup/comparison/search/insertion bodies; V029..V062, V090..V095.

**Confidence.** High for the complete lookup/comparator/search/insertion
instructions and discriminating conditional vectors.

**Unknown.** Nonzero CRT locale state, arbitrary unterminated input and external
memory reads.

### REG-101

- The linear lookup returns the first equal child. Sorted blocks of two, three
  and four identical names return indices 0, 1 and 1 under the original bsearch,
  which provides no first-duplicate rule.
- Neither route checks parent bit 0 or skips child bit 30. With explicitly valid
  child addresses and counts, parent kind controls 0, 1, 2, 17, 18 and
  `0x40000001` all reach the selected integer.
- Getter `R0452` performs the root and leaf lookups and returns its caller
  default for a miss. For a hit it checks only `kind & 0x0e == 2`: kinds 2, 3,
  18, 19, `0x10000002` and `0x40000002` return the stored dword; kinds 0, 4 and
  6 reach the exception-throw boundary.
- Evidence: lookup, bsearch and integer getter; V063..V083.

**Confidence.** High for the complete direct routines and named conditional
controls. The parent-kind vectors are accessor tests with supplied valid child
ranges, not assertions that those kinds describe valid tree grammar. Native
error presentation and other typed accessors remain outside this claim.

### REG-102

- `R1390` calls `R1394` and writes the supplied registry's header, all
  `R*32` records and the pool.
- Original sort/qsort/comparator instructions and synthetic stream services
  retain each of six unrelated root/leaf additions.
- A lowercase root that matches the unsorted case-folding lookup before writing
  misses that lookup after the writer sorts and sets the root sorted bit.
- The recursive sort candidate at `L07395..L07396` is
  `records[parent.value + parent.size]`, independent of the loop counter. This
  is not evidence that every descendant list is conventionally sorted.
- Internal record-copy/filter helpers and a fresh application-state producer are
  distinct boundaries.
- Evidence: complete writer/sort/qsort/shortsort/swap/comparator; V084..V090.

**Confidence.** High for the supplied-registry writer body, exact
recursive-candidate arithmetic and seven conditional round trips. No general
heap-safety, arbitrary-tree recursion, stable duplicate order or ordinary
application resave claim follows.

## Registry location and class key catalogues

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-LOC-016 | The unit, object, structure, projectile and material class definitions are five `.reg` nodes in `graphics.res`, 5 of the install's 44, with literal ASCII key names. | High / Medium | ● active (amended) | [EXP-0006](../experiments/EXP-0006-registry-keys/) |
| REG-FMT-017 | A `.reg` node is a nested `&YA1` holding a flat array of 32-byte records and a trailing pool. | High | ● active (amended) | [EXP-0006](../experiments/EXP-0006-registry-keys/), [EXP-0011](../experiments/EXP-0011-reg-values/), [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-UNITS-018 | `units.reg` holds 34 classes under one key catalogue, and its kind-6 keys are per-phase animation tracks whose lengths need not equal the `*Phases` scalars. | High / Medium | ● active (amended) | [EXP-0006](../experiments/EXP-0006-registry-keys/), [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-ROSTER-019 | The 34-unit `DescText` roster read from the `units.reg` pool stands; its name-to-sprite-path pairing is retracted and replaced by `REG-ROSTER-052`. | High | ● active (partially retracted) | [EXP-0006](../experiments/EXP-0006-registry-keys/), [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/), [EXP-0043](../experiments/EXP-0043-reg-restorations/) |
| REG-OBJ-020 | `objects.reg`'s keys carry the RT-02 object frame and destroyed/burning references, and the key catalogue is unaffected by the EXP-0032 framing shift. | High | ● active (amended) | [EXP-0006](../experiments/EXP-0006-registry-keys/), [EXP-0033](../experiments/EXP-0033-reg-value-tables/) |
| REG-OBJ-039 | `objects.reg` holds 82 objects, not 56: `[Global] ObjectCount = 82`, and `ID` equals the section index on 82/82. | High | ● active | [EXP-0033](../experiments/EXP-0033-reg-value-tables/) |
| REG-STR-021 | `structures.reg` holds 66 structures, and its keys include the RT-05 structure fields. | High | ● active (amended) | [EXP-0006](../experiments/EXP-0006-registry-keys/), [EXP-0033](../experiments/EXP-0033-reg-value-tables/) |
| REG-STR-040 | `structures.reg` holds 66 structures with `ID == section index + 1`; `FullHeight` is 1..6, `TileHeight` 1..6 and `TileWidth` 1..11. | High | ● active (partially retracted) | [EXP-0033](../experiments/EXP-0033-reg-value-tables/) |
| REG-PROJ-022 | `projectiles.reg` holds 31 projectiles, with keys including `Phases`, `RotationPhases`, `Homing`, `Flip`, `Palette` and `Width/Height`. | High | ● active (amended) | [EXP-0006](../experiments/EXP-0006-registry-keys/), [EXP-0033](../experiments/EXP-0033-reg-value-tables/) |
| REG-PROJ-041 | `projectiles.reg` holds 31 projectiles over a sparse `ID` domain `1..62`; `SFX` reads `1` wherever present, and `File` is a path string. | High / Medium | ● active | [EXP-0033](../experiments/EXP-0033-reg-value-tables/) |
| EDITOR-023 | The install's Map Editor ships its own catalogues: 22 trigger condition opcodes, 44 trigger action opcodes and a 1088-slot placeable-template id→name table. | High | ● active | [EXP-0006](../experiments/EXP-0006-registry-keys/) |

### REG-LOC-016

- The five nodes: `units/units.reg`, `objects/objects.reg`,
  `structures/structures.reg`, `projectiles/projectiles.reg`,
  `units/material.reg`.
- The RT-03 key names (`AttackPhases`, `MovePhases`, `DyingPhases`, `Flip`,
  `Projectile`, …) are literal ASCII in the record array: the labels are the
  game's own.
- These five are 5 of the install's 44 registries (`REG-LOC-038`).

**Confidence.** High for the key names: a direct byte observation, not a fit;
the names are literally present in the record array, at a file offset the
EXP-0032 correction did not move. Medium for the scope: "five" was the set then
examined and was wrong by a factor of nine (`REG-LOC-038`), the same error shape
as the count in this claim's own successor.

**Amended.** EXP-0032 widened the scope. The clause that the class definitions
live in five `.reg` nodes described the nodes then examined; the install holds
44 registries (`REG-LOC-038`, [`retracted.md`](retracted.md)).

### REG-FMT-017

- A record is
  `[value:u32 @+0x04][size:u32 @+0x08][kind:u32 @+0x0C][name:char16 @+0x10]`
  (`REG-REC-032`).
- Kinds: 1 = subkey (`value` = child block start, `size` = count); 2 = int32
  (`value`); 0 = string and 6 = int32[] (`value` = pool offset, `size` = byte
  length); 4 = double (`REG-DBL-035`); 10 = double[].

**Confidence.** High, as amended. Every field and every kind is carried by
`REG-REC-032`/`REG-KIND-034`, that is by an explicit type test and an explicit
writer in `rom.exe`. The corpus closure that preceded them held under the wrong
framing too, which is why this claim needed amending.

**Amended.** EXP-0032 found the published record framing 4 bytes early. The old
reading `[A][kind][name:20][C]` on a `0x20` header paired each key's name with
the next key's value; the record above replaces it, amended in place
([`retracted.md`](retracted.md), `REG-REC-032`). The kind list is the amended
one, with kind 4 = double from `REG-DBL-035`.

### REG-UNITS-018

- `[Global] UnitCount` reads 34 on the corrected framing, matching the 34
  `Unit*` sections. `[Global] FileCount` reads 33, matching the 33 `Files`
  children.
- Keys:
  `ID File Index Palette Sound DescText InfoPicture Parent InMapEditor Move/MoveBegin/Attack/Dying/Bone/IdlePhases Flip Width Height CenterX/Y SelectionX1..Y2 Z TileSize Dying AttackDelay Projectile ShootDelay ShootOffset {Attack,Move,Idle}AnimTime/Frame`.
- Ranged-only keys appear only on the ~9–11 ranged classes.
- Kind-6 arrays are per-phase animation tracks, for example
  `AttackAnimFrame=[0..6]`.
- `REG-UNITS-049` (EXP-0039) turns this catalogue into a field map: every key's
  offset in the loaded class record, its default, and the three timelines the
  loader builds out of the kind-6 tracks. `SPR256-UNIT-024` states what each
  `*Phases` scalar counts.
- A track's length need not equal the matching `*Phases` scalar. On the EXP-0032
  base `AttackPhases` agrees on 24/24 classes carrying both, `MovePhases` on
  15/18 (`Goblin` 8 vs 10, `Ghost` 3 vs 4, `Goblin slinger` 8 vs 10) and
  `IdlePhases` on 4/5 (`Ghost` 3 vs 4).
  `MoveBeginPhases`/`DyingPhases`/`BonePhases` have no `*AnimTime`/`*AnimFrame`
  key at all, so there is no length to compare.
- The other two registries agree less: `objects.reg` `Phases` vs
  `len(AnimationTime)` 7/16, `structures.reg` `Phases` vs `len(AnimTime)` 12/14.
- The loader iterates the array's own element count, never `Phases`:
  `R1395` uses `*local_294`.
- EXP-0032's two examples, `Unit0` `AttackPhases=7`/`MovePhases=8` matching
  their 7- and 8-element `Anim*Time`/`Anim*Frame` arrays, are two of the
  agreeing cases.

**Confidence.** High, as amended, for the catalogue and the counts: the key
catalogue is a name list, and names did not move under the EXP-0032 correction;
34 and 33 were re-read on the corrected base and each matches an independent
structural count. Medium for "ranged-only keys appear only on the ~9–11 ranged
classes": an approximate figure with no exact count attached, unparameterised as
written.

**Amended.** EXP-0032 corrected the counts: the published `UnitCount` 33 was an
off-by-one, and the old framing read `FileCount` as 0 and `Unit0`'s
`AttackPhases`/`MovePhases` as 128 and 2 against the same 7- and 8-element
arrays. EXP-0037 withdrew the clause that a kind-6 track's length "agrees with
the matching `*Phases` scalar", which had been promoted into
`formats/reg/format.md`; the agreement counts above replace it
([`retracted.md`](retracted.md)). A consumer reads the array length and treats
`*Phases` as a separate scalar.

### REG-ROSTER-019

- Roster: `Human Swordsman`, `Goblin`, `Dragon`, `Ogre (Lord)`, `Turtle`,
  `Human Archer/CrossBowMan`, `Catapult 1/2`, `Hero of the Lance/Sword`, …
- Re-derived verbatim on the corrected framing: EXP-0032
  `evidence/units-roster.txt`, 34 `DescText` values and a 33-entry `Files`
  table.

**Confidence.** High for the 34 names, re-derived.

**Amended.** The `DescText` name → sprite `File` path pairing this claim
asserted is retracted and replaced by `REG-ROSTER-052` (EXP-0043), which
regenerated it on `Files[class.File]` ([`retracted.md`](retracted.md)). The
published arrows were wrong on 34/34 classes, not merely stale, and 14 of them
indexed outside the 33-entry `Files` table: the pairing was falsifiable from the
shipped bytes alone at any time since EXP-0006. The names stand; the arrows are
superseded.

### REG-OBJ-020

- Keys include `Index` (frame), `Phases`, `DeadObject`, `FireObject`, `IconID`,
  `AnimationTime/Frame`, `Width/Height`, `CenterX/Y`, `Parent`: the RT-02 object
  frame plus destroyed/burning references.
- Record names sit at the same file offset under both framings, so the key
  catalogue is unaffected by the EXP-0032 shift.

**Confidence.** High. The immunity argument is checked and correct: `0x28 + 32i`
under both splits, which is what makes a name list safe where a value list is
not.

**Amended.** EXP-0033 corrected the object count (`REG-OBJ-039`): it was never a
framing casualty, but `FileCount` transposed for `ObjectCount`.

### REG-OBJ-039

- Re-derived on the EXP-0032 base, amending `REG-OBJ-020`'s count.
  `[Global] ObjectCount = 82` matches 82 distinct top-level `Object*` subkeys
  and, as an independent cross-check, 82 distinct `ID` values `0..81`.
- For this registry `ID == section index` on 82/82 (EXP-0037), so the two
  possible keys coincide here and nowhere else (`REG-KEY-044`).
- The published 56 is `[Global] FileCount`, the size of the sprite-file
  `Files[]` leaf table, not the object count. It was a transposed-counter slip
  present since EXP-0006/0011, not a framing casualty; EXP-0011's own raw
  observations recorded "objects → 82 + Global".
- Measured ranges, present on 43/82 unless noted: `Width` 32..128 [4], `Height`
  32..128 [5], `DeadObject` −1..67 [22] (35/82), `FireObject` −2..−1 [2]
  (14/82).

**Confidence.** High. The count carries two independent structural cross-checks,
82 `Object*` subkeys and 82 distinct `ID`s, on a framing read from the loader.
The ranges are measurements on that framing, each with its presence denominator.

### REG-STR-021

- Keys include `AnimMask`, `Phases`, `Flat`, `Indestructible`, `Usable`,
  `VariableSize`, `LightPulse`, `LightRadius`, `ShadowY`, `SelectionX1..Y2`: the
  RT-05 structure fields.
- The count holds on the EXP-0032 base (`REG-STR-040`).

**Confidence.** High for the keys and the count.

**Amended.** EXP-0033 withdrew the ranges `FullHeight (0..20)` and the bundled
`TileWidth/TileHeight (1..6)`; `REG-STR-040` gives the corrected ranges.

### REG-STR-040

- Re-derived on the EXP-0032 base, amending `REG-STR-021`'s ranges.
  `[Global] Count = 66` matches 66 `Structure*` blocks and the `ID` domain
  `1..66`: 66 distinct, and `ID == section index + 1` on 66/66. The count holds.
- `FullHeight` = 1..6 [6 distinct], present on all 66, not `0..20`. The
  published `0..20` is `SelectionX1`'s own corrected range (`0..20` [13]): the
  old `FullHeight` column read the next key's true value, the EXP-0032
  mechanism.
- `TileHeight` = 1..6 [6]: the old bundled figure holds, now attached to the
  right key. `TileWidth` = 1..11 [6]: wider than the bundled `1..6`, which was
  `TileHeight`'s range.
- Only `objects.reg` is 0-based. `units.reg`'s `ID` domain is `1..80` with 34
  distinct: sparse and 1-based, equal to neither the section index (0/34) nor
  index+1 (0/34), as `REG-PROJ-041` states.

**Confidence.** High. The count is cross-checked twice, and the `FullHeight`
correction demonstrates itself: the withdrawn `0..20` is the neighbour key's own
corrected range, the EXP-0032 mechanism reproducing itself.

**Amended.** The parenthetical "unlike `units`/`objects`' 0-based `ID`" is
refuted by EXP-0037: only `objects.reg` is 0-based, and `units.reg` is sparse
`1..80` ([`retracted.md`](retracted.md)). A consumer that believed the
parenthetical indexed the units array off by one on every class.

### REG-PROJ-022

- Keys include `Phases`, `RotationPhases`, `Homing`, `Flip`, `Palette`,
  `Width/Height`.
- The count holds on the EXP-0032 base (`REG-PROJ-041`).

**Confidence.** High for the keys and the count.

**Amended.** EXP-0033 withdrew the range `SFX (12..62)` (`REG-PROJ-041`).

### REG-PROJ-041

- Re-derived on the EXP-0032 base, amending `REG-PROJ-022`.
  `[Global] Count = 31` matches 31 `Projectile*` blocks and the `ID` domain
  `1..62`: 31 distinct, non-contiguous. The same shared-namespace pattern
  appears in `units.reg`'s `ID 1..80`/34 and `structures.reg`'s `ID 1..66`/66.
- `SFX` is present on 23/31 projectiles and reads `1` in every one: a presence
  flag, not a variable id, on this base. The published `12..62` is withdrawn and
  not re-attributable: no single re-measured neighbour key reproduces it closely
  enough to name. `Width` 12..128 was tried and is close in shape, not identity,
  so the old range is reported as superseded, not re-explained.
- `Width`/`Height` measure `12..128`, present on 23/31.
- `File` is a kind-0 string, the sprite path (for example `archer\arrow`), on
  31/31. Unlike `units.reg`/`objects.reg`, projectiles store the path directly
  rather than a `Files[]` index.

**Confidence.** High for the count and the `SFX` withdrawal. Medium for the
non-re-attribution; see the EXP-0033 Confidence.

### EDITOR-023

- `Description Checks.ini` lists 22 trigger condition opcodes,
  `Description Instants.ini` 44 trigger action opcodes, and `Templates.ini` a
  1088-slot placeable-template id→name table.
- Type system: `Target_Unit/Group/Player/Item/Structure/Building`, `X`, `Y`,
  `int`, `Enum`, `Const`.

**Confidence.** High. An enumeration of shipped plain-text `.ini` files: the
counts are what the files list, with no model interposed. Corpus: this install's
Map Editor.

## Header, pool, values and subkey tree

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-VAL-024 | A registry header is `0x18` bytes and a `u32` length prefixes the pool, so `poolStart = 0x18 + R*32 + 4`; the closure holds on 44/44 registries. | High | ● active (amended) | [EXP-0011](../experiments/EXP-0011-reg-values/), [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-VAL-025 | Each kind-0/6/10 record stores its pool byte offset in its own `value` field and its byte length in `size`; `rom.exe` reads `poolBase + node->value`. | High | ● active (amended) | [EXP-0011](../experiments/EXP-0011-reg-values/), [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-VAL-026 | Old-framing `record[i-1].C` is corrected-framing `record[i].value`; there is no snapshot slot, and a kind-2 key before a pool item keeps its scalar. | High | ● active (amended) | [EXP-0011](../experiments/EXP-0011-reg-values/), [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-VAL-027 | A kind-2 scalar is a signed int32 in the record's own `value` field at `+0x04`. | High / Medium | ● active (amended) | [EXP-0011](../experiments/EXP-0011-reg-values/), [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-VAL-028 | A subkey's `value` is its child block start index and `size` its child count, and the blocks cover the record array exactly once on all 44 registries. | High | ● active (amended) | [EXP-0011](../experiments/EXP-0011-reg-values/), [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-VAL-029 | Resolved `units.reg` has 34 classes and a 33-entry `Files` table, and a class's sprite is indexed by its own `File` scalar, not by `ID`. | High / Medium | ● active (amended) | [EXP-0011](../experiments/EXP-0011-reg-values/), [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-VAL-030 | On the corrected base `units.reg` `Width`/`Height` are 128×128 on 17 classes (Dragon 160×160), `CenterX` 64 (Dragon 80) and `CenterY` 58…84. | High | ● active (partially retracted) | [EXP-0011](../experiments/EXP-0011-reg-values/), [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/), [EXP-0033](../experiments/EXP-0033-reg-value-tables/) |

### REG-VAL-024

- `rom.exe`'s loader performs six 4-byte reads (`magic`, root `value`, root
  `size`, root `kind`=17, record count `R`, one more word) before bulk-reading
  `R×32` bytes of records.
- The pool is length-prefixed by a `u32`, not overlapped by the last record's
  `C`. `poolStart = 0x18 + R*32 + 4` is the same number the old rule gave
  (`0x20 + R*32 − 4`), which kept the error invisible.
- `@0x1C` is record 0's `value`, not a "leaf-table first index" header field;
  record 0 happens to be the `Files` subkey in `units.reg`.
- Structural closure `0x18 + R*32 + 4 + poolLen == payloadLen` holds on 44/44
  registries over the whole install.

**Confidence.** High, as amended.

**Amended.** EXP-0032 corrected the header from `0x20` to `0x18` bytes, the pool
from an overlap of the last record's `C` to a `u32`-prefixed block, and `@0x1C`
from a header field to record 0's `value`.

### REG-VAL-025

- `value` is at `+0x04` and `size` at `+0x08`. `rom.exe` reads
  `poolBase + node->value` with no cursor.
- Pool items tile the pool exactly, verified as a direct-offset tiling with no
  gap, overlap or tail on all 44 registries.

**Confidence.** High, as amended.

**Amended.** EXP-0032 replaced the running-cursor mechanism. The cursor was a
symptom of the 4-byte shift and reproduced the right addresses only because it
walked the same items in the same order.

### REG-VAL-026

- The 259/163/396/31/16 = 100% match stands and is now explained by the framing
  rather than only observed.
- Human Archer reads `Projectile=1`, `ShootDelay=21`, `AttackDelay=19` under the
  correct framing.
- The `~` flag should be removed from tools.

**Confidence.** High, as amended.

**Amended.** EXP-0032 kept the fact and moved it to the right record. It
withdrew the consequence first drawn from the match: a "snapshot slot" in which
a kind-2 key before a pool item loses its scalar.

### REG-VAL-027

- Under the old framing every scalar was attributed to the key one position too
  early: the values were real, the key each belonged to was wrong.
- `FireObject = −1` and the `Width`/`Height`/`Center*`/`Selection*` sanity check
  must be re-measured on the corrected base before being relied on.

**Confidence.** Medium overall: High for the mechanism; the per-key attributions
were not re-measured.

**Amended.** EXP-0032 moved the value from the `C` at `+0x1C` to the record's
own `value` field at `+0x04`.

### REG-VAL-028

- `rom.exe` computes `&records[parent->value]` and iterates `parent->size`
  entries (`value` at `+0x04`, `size` at `+0x08`).
- The root block plus every subkey block cover the record array exactly once,
  with 0 unreachable and 0 doubly-reached records, on all 44 registries. The
  direct rule needs no tiling carve-out.
- `A`-on-inheriting-classes no longer arises.

**Confidence.** High, as amended.

**Amended.** EXP-0032 replaced the rule "`C` = block end, children `[C−A,C)`,
tile by sorted ends", the shifted reading of the same data, which needed a
tiling carve-out.

### REG-VAL-029

- `DescText` reproduces the EXP-0006 roster, re-derived verbatim on the
  corrected base.
- 36 root children: `Files` + `Global` + 34 classes.
- Human Archer has `File=8` → `humans\archer\archer` and `ID=14`.
- `Flip=1` on 10 classes still holds.
- `Projectile`/`ShootDelay`/`ShootOffset` are present on the ranged classes and
  carry readable values.

**Confidence.** High for the roster, the counts and the `File` index. Medium for
the ranged-class keys: the per-key ranged-class census was not re-run.

**Amended.** EXP-0032 replaced the "indexed by `ID`" reading, which was the
4-byte shift showing through: the shifted `ID` slot held the true `File` value.

### REG-VAL-030

- `units.reg` `Width`/`Height` = 128×128 on 17 classes, Dragon 160×160;
  `CenterX` = 64, Dragon 80; `CenterY` per class 58…84.
- The SPR256 comparison is resolved by `REG-VAL-043` (EXP-0033), re-measured
  with both the corrected canvas and the corrected `Files[class.File]` sprite
  index: 18/18 within-canvas, 1/18 exact.

**Confidence.** High for `Width`/`Height`/`CenterX`/`CenterY`, and for the
SPR256 comparison through `REG-VAL-043`.

**Unknown.** The placeholder-phase-count observation remains unverified on the
new base; it was not the re-measurement's target.

**Amended.** EXP-0032 retracted the figures in part: they were measured on the
4-byte-shifted framing, so each named key reported the next key's value. The
published "`128×64`, Dragon `160×80`" is withdrawn for good, and no downstream
table should quote it. The SPR256 comparison built on those numbers
(`0/18 exact`, `only 4/18 fit`) is withdrawn and replaced by `REG-VAL-043`
([`retracted.md`](retracted.md)).

## Loader framing, record layout, kinds and text bytes

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-FMT-031 | Transcribed from `rom.exe`'s loader `R1389`, a registry is a `0x18`-byte header, `R×32` bytes of records, a `u32` pool length and the pool. | High | ● active | [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/), [EXP-0040](../experiments/EXP-0040-enumeration-audit/) |
| REG-REC-032 | A registry record is 32 bytes, `value`, `size`, `kind` and a 16-byte name from `+0x04`, and record `i`'s value word is file `0x1C + 32i`. | High | ● active (amended, partially retracted) | [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-KIND-033 | The `kind` word is a bitfield, not an enum: `kind & 0x0E` is the value type, bit 0 marks a subkey and bit 4 marks sorted children. | High | ● active | [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-KIND-034 | Kinds 0 string, 1 subkey, 2 int32, 4 double, 6 int32[] and 10 double[] each have an explicit type test and writer in `rom.exe`, with defined storage. | High / Unknown | ● active (partially retracted) | [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-DBL-035 | Kind 4 is a little-endian IEEE-754 `double` held in the record's own `value` and `size` fields, 8 bytes at `+0x04`, with no pool access. | High | ● active | [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-TEXT-036 | `rom.exe` performs no byte-to-character conversion on `.reg` text, so a code page for `.reg` is a display convention, not a transcription. | High | ● active (amended) | [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-TEXT-037 | In this install no `.reg` name or string value holds a byte `>= 0x80`, so the corpus cannot tell apart code pages that agree on ASCII. | High | ● active | [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |

### REG-FMT-031

- `+0x00` magic `0x31415926`, `+0x04` root `value`, `+0x08` root `size`, `+0x0C`
  root `kind` (=17), `+0x10` record count `R`, `+0x14` one further word, `+0x18`
  the `R×32` record array read verbatim, then a `u32` pool byte length, then the
  pool read verbatim.
- Exact structural closure on 44/44 registries in the install, with 0 invariant
  violations. The previously published `0x20` framing produces 82 violations
  over the same 44.

**Confidence.** High. Transcribed from `R1389`'s read sequence, not
fitted. Here the corpus also discriminates: 0 violations against 82 for the
alternative, over the same files.

### REG-REC-032

- In absolute file terms record `i` is `[0x18+32i, 0x18+32(i+1))`. Its value
  word is `0x1C + 32i`, eight bytes before the kind word, not the `0x3C + 32i`
  word after the name, which is record `i+1`'s value. The same word serves an
  integer key and a string key's pool offset.
- Stride 32 B (a left shift by 5 in the loader's allocation, in the lookup and in the
  node-insert); `value` `u32 @+0x04`, `size` `u32 @+0x08`, `kind` `u32 @+0x0C`,
  `name` `char[16] @+0x10` NUL-padded.
- The unsorted lookup arm of `R0505` compares names over 15 characters
  (`_strnicmp(key, node+0x10, 0xf)`). Its sorted arm truncates the query to 15
  bytes but compares stored names case-sensitively through NUL (`REG-100`).
- Four independent attestations:
  - loader `R1389`: six 4-byte header reads, then `Read(records, R<<5)`;
  - lookup `R0505`: `records + node[+4]*32`, iterate `node[+8]`, compare
    `node[+0x10]`;
  - accessors `R0452` / `R1396`: `*(u32*)(node+4)`;
    `poolBase + *(int*)(node+4)` with length `*(u32*)(node+8)`;
  - the node-insert writer `R1393`: it places a new child at
    `records[parent[+4] + parent[+8]]`, increments `parent[+8]`, writes the name
    at `node+0x10` clamped to 15 characters, so the field is 16 B and cannot be
    20, and stores `0` into `node+0x00`.
- `+0x00` is a real field, not name padding: zero in 4621/4621 records and read
  by no accessor found.
- This is the node record the RES ledger publishes. `RES-NODE-016` gives
  `[+0 reserved][+4 off][+8 size][+0xc type/flags]`, and `RES-HDR-017` quotes
  the instruction
  `R0505 @L07397`, which loads the node's 32-bit field at +4, shifts it left by 5 and adds the result to the field at +0x28 of the object. A `.reg`
  record is that node plus a 16-byte name: one class, one lookup. The two
  ledgers contradicted each other from EXP-0017 until this claim, although the
  RES reading was already in the repository.
- `RES-NODE-019` records no `[node+0]` dereference in the four scanned lookup,
  path-walk and finalize-sort functions.

**Confidence.** High. Four independent `rom.exe` attestations (loader, lookup,
two accessors and the writer), each named to a function, with the 15-character
clamp fixing the field size from the write side; no fifth reading fits.

**Unknown.** The purpose of `+0x00`, and whether another node-touching function
reads that word.

**Amended.** `REG-100` partially retracts the blanket lookup-comparison clause,
"the lookup compares names over 15 characters using `_strnicmp`, so 15
significant characters": only the unsorted route uses the 15-byte case-folding
comparator, and the sorted route truncates its query to 15 bytes plus NUL, then
compares stored bytes case-sensitively through NUL. The restatement of
`RES-NODE-019` as "`[node+0]` is never dereferenced" is narrowed to the four
scanned functions. The record offsets, stride, producer clamp and corpus counts
stand ([`retracted.md`](retracted.md)).

### REG-KIND-033

- The value type is tested as a mask with 0xe in every accessor and as
  `(kind>>1)&7` in the string converter.
- With bit 4 set the lookup `bsearch`es the children instead of linear-scanning
  (a test of bit 0x10 at `L07398`).
- The root node's `kind = 17 = 0x11` is therefore subkey | sorted, not a magic
  constant, which is why every `.reg` in the install carries `17` at file
  `+0x0C`.
- The node-insert `R1393` clears bit 4 with `kind & 0xffffffef` when it
  appends a child, and sets bit 28 (`0x10000000`) when the name it was handed
  exceeded 15 characters and was truncated. No shipped record carries bit 28.

**Confidence.** High. Each bit is pinned to the instruction that tests or writes
it. "Bitfield, not enum" is decidable only from code; the corpus, where every
root reads the same `17`, is what would keep it looking like a magic constant.

### REG-KIND-034

- Storage, each kind attested by an explicit type test and an explicit writer:
  - `0` string: `value` = pool byte offset, `size` = NUL-inclusive length;
  - `1` subkey: `value` = child block start, `size` = child count;
  - `2` int32: `value`, signed;
  - `4` double (`REG-DBL-035`);
  - `6` int32[]: `value` = pool offset, `size` = byte length, `size/4` elements;
  - `10` double[]: `size/8` elements.
- The string converter has one further case, `(kind>>1)&7 == 4`, that is kind
  8/9: skip 4 bytes of the pool item, copy `size−4`, terminate with two NULs. No
  record of that kind exists in the corpus.
- Corpus counts over all 44 registries: kind 0 ×873, kind 1 ×512, kind 2 ×2792,
  kind 4 ×122, kind 6 ×322; kinds 8 and 10 ×0.

**Confidence.** High for kinds 0/1/2/4/6/10 and the flags.

**Unknown.** The semantics of kind 8/9, for which this claim made no guess;
`REG-105` has since located kind 8's producer (see Amended).

**Amended.** `REG-105` partially retracts the producer absence ("no writer
exists anywhere in the binary") and the exhaustive-enumeration wording: original
setter `L07399` writes kind 8 as a pool dword count followed by terminated
strings, consumed by `L07400`. The other established storage and the corpus
counts stand ([`retracted.md`](retracted.md)).

### REG-DBL-035

- Reader: `R1397` tests `(kind & 0xe) == 4`, then
  a 64-bit floating-point load from the node's field at +0x4.
- Writer: the text importer `R1398` does
  `*(double*)(node+4) = atof(text); node->kind = 4`.
- Cross-check: the string converter passes `*(u32*)(node+4)` and
  `*(u32*)(node+8)` to `sprintf` as one double. The game's own error string is
  `"not an double associated with specified key"`.
- 122 instances exist in the install: `startfade`/`endfade` under `Fading<n>` in
  31 of the 33 video registries of `Allods/VIDEO4.RES`+`VIDEO8.RES`. Every one
  is `0x0000000000000000` (0.0, ×61) or `0x3FF0000000000000` (1.0, ×61), forming
  a fade-in then a fade-out whose count matches each file's own `nFadings`.
- Every other candidate 8-byte window yields denormal garbage (`8.49e-314`,
  `9.02e-314`, `5.30e-315`).

**Confidence.** High. Reader and writer are both named to an instruction, and
the corpus excludes the alternative windows by producing denormals: the bytes
themselves rule out the competing placements.

### REG-TEXT-036

- The loader bulk-reads the record array and the pool verbatim (`Read(ptr, n)`,
  no conversion call in `R1389`). The string accessor `R1396` does
  `strncpy(dst, poolBase + node->value, min(node->size, cap))`. No code page is
  named anywhere on the path.
- Whole-image reachability sweep:
  `MultiByteToWideChar`/`WideCharToMultiByte`/`LCMapStringA`/`GetACP`/`GetOEMCP`
  are called only from the statically linked MFC/CRT range
  code-region, never from the registry class
  `L07401–L07402`. The `CharToOemA`/`OemToCharA` thunks have zero
  callers, and `setlocale`/`_setmbcp`/`IsDBCSLeadByte` appear nowhere.
- On a repaired function table (EXP-0040: the vtable blind spot closed, 607
  functions created), 0 references to any code-page entry point fall inside
  `L07401–L07402`, and 0 of the 107 references image-wide lie in orphan
  code. Four references sit below code-region, the lowest at `L07403`.
- The only byte transform identified on this path is the unsorted name lookup's
  `_strnicmp`, whose fast path folds `A`–`Z` only
  (compares against 0x41 and 0x5a, then adds 0x20) and leaves bytes `>= 0x80`
  untouched. The sorted comparator performs no case folding (`REG-100`).
- CP866 and CP1251 are both defensible renderings, and neither is the game's.

**Confidence.** High. This is an absence, established the only way an absence
can be: a whole-image reachability sweep of every conversion entry point, plus
the one surviving byte transform read and bounded to `A`–`Z`. EXP-0040
re-checked it on the repaired table and confirmed it.

**Amended.** EXP-0040 corrected the caller range from code-region to
code-region, a correction of precision, not of conclusion. `REG-100`
amended the lookup scope: the case-folding `_strnicmp` is the unsorted route
only. The no-conversion result and the import-reference census are unchanged and
were not remeasured for that amendment.

### REG-TEXT-037

- Sweep parameters: every regular file under `ROM_DIR`; every file beginning
  `&YA1` opened; nested payloads classified by root `kind == 17`; each record's
  name taken as the 16 bytes at `+0x10` truncated at the first NUL, and each
  string value as its kind-0 pool item.
- 0 bytes `>= 0x80` occur in 4 621 name fields or 873 string values. The
  observed byte range is `0x20…0x7A` in both (59 / 65 distinct values).
- No shipped `.reg` datum exercises a code page.

**Confidence.** High for this install and this sweep.

## Install inventory and re-measured values

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-LOC-038 | The install holds 44 `.reg` registries in 12 `&YA1` containers, and every nested `&YA1` payload in it is a registry. | High | ● active | [EXP-0032](../experiments/EXP-0032-reg-text-and-kinds/) |
| REG-MAT-042 | `material.reg` is 16 `Material0..Material15` blocks, each carrying exactly one `Path` string, with no `[Global]` counter. | High | ● active | [EXP-0033](../experiments/EXP-0033-reg-value-tables/) |
| REG-VAL-043 | `units.reg`'s per-class canvas contains the SPR256 sprite's maximum frame bounds on 18 of 18 classes carrying `Width`/`Height`, exactly on 1/18. | High | ● active | [EXP-0033](../experiments/EXP-0033-reg-value-tables/) |

### REG-LOC-038

- Beyond the 5 in `graphics.res` that `REG-LOC-016` located: `scenario.res`
  (`globalmap`, `npc`, `scenario`), `sfx.res` (`sfx`), `world.res` (`data/ai`,
  `data/map`), and 33 video-cutscene registries in `Allods/VIDEO4.RES` (18) and
  `Allods/VIDEO8.RES` (15).
- The cutscene registries carry `Common{startx,starty,nFadings,nPanaramings}` /
  `Fading<n>{startframe,endframe,startfade,endfade}` /
  `Panaraming<n>{startframe,endframe,stepx,stepy}` sections, consumed by
  `R1399`.
- 0 nested payloads were rejected by the classifier.
- Kind 4 lives here: 31 of the 33 cutscene registries carry it (`LOGOS/1c.reg`
  and `LOGOS/buka.reg` do not). The statement "kind 4 is absent" was true of
  `graphics.res` and false of the install.

**Confidence.** High. An exhaustive enumeration of the install: every regular
file opened and every nested `&YA1` classified, 0 rejected, so the count is a
listing, not a sample. It is the standing example of why a scope figure needs
its sweep stated: `REG-LOC-016`'s "five" was five examined.

### REG-MAT-042

- Re-verified unchanged on the EXP-0032 base. The pre-EXP-0032 observation of
  this registry's shape did not depend on which value word was read, since it
  never resolved a scalar.

**Confidence.** High for the structure only: 16 blocks and one `Path` string
each, re-verified on the corrected base, with the immunity argument stated
rather than assumed.

### REG-VAL-043

- The canvas is `Width`/`Height` = 128×128, Dragon 160×160, both re-measured by
  EXP-0032 (`REG-VAL-030`).
- The sprite is resolved via `Files[class.File]`, `REG-VAL-029`'s corrected
  index, not `Files[class.ID]`.
- The exact case is `Ogre (Lord)`: canvas 128×128 = sprite max frame 128×128.
- The pre-EXP-0032 finding, "0/18 exact, only 4/18 fit within it", used the
  wrong canvas numbers (measured one key early) and the wrong sprite (indexed by
  `ID` instead of `File`) at once, and does not survive either correction alone.
- This resolves the SPR256 canvas comparison `REG-VAL-030` left open.

**Confidence.** High. Both inputs are re-measured on the corrected base, and the
old result fails under either correction alone, so the agreement is not a
coincidence of two errors cancelling.

## Class arrays and Parent inheritance

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-KEY-044 | The `units.reg` and `structures.reg` class arrays are keyed by `ID` and `objects.reg`'s by append order, so `Parent` and `.alm` placement ids are `ID` subscripts used directly. | High / Medium | ● active (amended) | [EXP-0037](../experiments/EXP-0037-alm-class-binding/) |
| REG-KEY-045 | `Parent` inheritance falls back on a missing record for scalar keys but on an empty result for array keys, so an explicit `= ""` does not clear an inherited array. | High / Unknown | ● active | [EXP-0037](../experiments/EXP-0037-alm-class-binding/) |

### REG-KEY-044

- Loaders: `R0395` (`units.reg` → the global at `L02113`), `R1395`
  (`objects.reg` → the global at `L02099`), `R1400` (`structures.reg` →
  the global at `L01425`). Each formats the section name `"<Prefix>%d"` from a 0-based
  counter running to `[Global]<Count> − 1`, looks it up by name and allocates
  the class: `0x12c` / 0x6c / 0xa4 bytes. The size 0x12c pushed at `L07404` is `DescText`'s
  `0x10c + 0x20`.
- Each stores the `ID` key at class `+0x04` (structures: `+0x0c`) and the
  counter itself at class `+0x08` (units/objects).
- The array subscript is per registry:
  - `units`: `L07405` reads the class's own `ID` from offset 0x04 of the class object, grow
    `m_nSize` to `ID+1`, `L07406` stores the class pointer at the array slot four bytes times `ID` ⇒
    `[L02113][ID] = class`, MFC `CObArray::SetAtGrow`;
  - `structures`: `L07407` reads the `ID` from offset 0x0c of the class object, `L07408` ⇒
    `[L01425][ID] = class`, the same;
  - `objects`: `L07409` reads `m_nSize` and `L07410` keeps a copy of it, `L07411`
    ⇒ `[L02099][m_nSize] = class`, `CObArray::Add`: append order, which for
    this registry equals both the section index and the `ID`.
- In all three the loop counter is a separate spill slot (stack offsets +0x1c /
  +0x14 / +0x20) and never the subscript. `m_nSize` is `maxID+1`,
  not the section count: `units.reg` 81 slots for 34 classes, 47 NULL, index 0
  among them; `structures.reg` 67 for 66, index 0 NULL; `objects.reg` 82 for 82,
  no holes. A consumer iterating the units array must null-check.
- `Parent` is an `ID`: `L07412/65` and `L07413/28` load `array[Parent]` as a
  plain scaled-index dword, with no call and no scan. The loader's next
  instruction reads `[parent+0x08]` to format the parent's section name, an
  indirection needed only because `Parent` is not itself a section number.
- The structures loader never reads `Parent`: `"Parent"` (`L07414`) is pushed
  at exactly two addresses image-wide, `L07415` and `L07416`, the code-side
  counterpart of `Parent` occurring on 0/66 structures.
- Inheritance is eager and per key, at load, not lazy at an accessor. The
  parent's resolved field is the default argument of
  `R0452(section, key, default)`, so a key in the child's own section
  wins and an absent one falls back (19 scalar sites `L07417`…`L07418`,
  `File` at `+0x0c` among them). The array keys
  (`Move`/`Idle`/`AttackAnimTime`/`AnimFrame`, `ShootOffset`) instead re-read
  the parent's own `.reg` section by that formatted name when the child's array
  comes back empty. The two guards differ, so an explicit `= ""` does not clear
  an inherited array (`REG-KEY-045`).
- Every consumer located subscripts directly: `R1401`
  (`[L02099][arg]` → `Files[class+0x0c]`), `R0379`
  (`[L02113][unit+0x20]`, `[L02099][byte]`), `R0218` /
  `R0212` (`[L01425][obj+0x20]`), `R0484`, `R0483`. For
  units and structures those subscripts are `ID`s.
- No search-by-`ID` over any of the three arrays exists in the image, the
  absence EXP-0037 established. For units and structures none is needed, because
  the array is the by-`ID` index.
- `ID` domains on the EXP-0032 base: `units.reg` 1..80, 34 distinct,
  `ID == index` on 0/34 and `index+1` on 0/34 (sparse, 1-based); `objects.reg`
  0..81, `ID == index` on 82/82; `structures.reg` 1..66, `ID == index + 1` on
  66/66.
- A `.alm` placement that stores an `ID` (`ALM-CLS-036` type-4 `kind`,
  `ALM-CLS-038` type-6 `+0x08`) is used directly as the array subscript, with no
  translation. The type-3 grid's `code − 1` (`ALM-CLS-035`) is an `objects.reg`
  `ID`, numerically unchanged there.
- Evidence: `evidence/disasm-regloaders.txt`,
  `evidence/reg-units-parent-resolution.csv`.

**Confidence.** High for the three append sites, the `Parent` subscript and the
eager-inheritance mechanism, at instruction level. The corpus discriminates the
units reading twice more: of the 16/34 units classes carrying `Parent` (values
`{3,12,21,26}`), the `ID` reading resolves 16/16 to a class already appended
that supplies the geometry the child omits, while the section-index reading has
4/16 parents not yet appended when the child loads (sections 1, 2, 3, 19, where
the subscript exceeds the array's `m_nSize`) and makes `Unit3` its own parent.
High for the three `ID` domains, a measurement over each whole registry. Medium
that `unit+0x20` and `obj+0x20` hold an `ID`: it follows from the array keying,
but the instruction that writes `+0x20` is unlocated. `objects.reg` cannot
witness the keying question: its `ID` equals its section index on 82/82 and its
loader is the one that appends in order, so both readings coincide there.

**Unknown.** The instruction that writes `unit+0x20` and `obj+0x20`, an open
item of EXP-0037.

**Amended.** A 2026-07-27 re-audit withdrew two clauses EXP-0037 had published:
(a) that the three graphics class arrays are section-index tables, so the array
subscript is the section index; (b) that `Parent`, resolved as
`[L02099][Parent]`, is a section index, not an `ID`. Downstream of both, it
stated that a `.alm` placement storing an `ID` must be translated; that
consequence is reversed. Only `objects.reg`'s array is a section-index table; it
is the registry the clause was read in and then generalised from. EXP-0039
corrected the units class size from `0x11c` to `0x12c` (`REG-UNITS-049`). Both
corrections are in [`retracted.md`](retracted.md).

### REG-KEY-045

- Both mechanisms are eager, per key, at load (`REG-KEY-044`); the test differs.
- Scalars (`File` and the 19 `Index`…`AttackDelay` sites) go through
  `R0452(section, key, default)`, which substitutes the `default` (the
  parent's resolved field) only when the record is not found:
  `L07419` tests the found-record pointer for zero and branches to `L06696`, which takes the `default` argument from the stack, versus
  `L07420` reading the value at offset 0x04 of the found record. A scalar written explicitly as `0`
  returns `0` and overrides the parent. A scalar present with a non-int kind
  throws (`L07421`) rather than defaulting.
- Arrays (`Move`/`Idle`/`AttackAnimTime` and `*AnimFrame`, `ShootOffset`) go
  through `R1061(section, key, dest)`, and the loader then tests
  `dest.m_nSize == 0`: the array's length after the read, and nothing about
  presence. `L07422` reads the 32-bit value at stack offset +0x28, where the stack offset +0x20 is the destination
  `CArray` and `+0x08` is `m_nSize`. The offsets are fixed by `R1198` =
  `CArray::SetSize`, which at `L07423/b8/bb` frees `m_pData` and zeroes
  `m_nMaxSize`/`m_nSize` when `nNewSize == 0`. Nonzero ⇒ skip; zero ⇒ `L07424`
  re-reads the same key from the parent's section into the same destination.
- A present, empty record reaches length 0 by a success path: `L07425`
  computes `kind = (record+0x0c >> 1) & 7`, and `kind == 0 && size < 2` takes
  `L07426 SetSize(0,-1); return 1`. A kind-6 record of size 0 reaches the same
  length via `L07427 SetSize(size>>2)`.
- Shipped corpus: `units.reg`'s `AttackAnimTime`/`AttackAnimFrame` occur on 33
  records each, 32 of kind 6 and one of kind 0. That pair is `Unit33` "Unarmed
  Fighter with Shield" (`ID` 2, `Parent = 3`) at kind 0, size 1, the NUL of an
  empty string (records 758/759, `tools/regdump -mode records`). `Unit33`
  therefore inherits `Unit0`'s `AttackAnimTime = [2 2 1 1 1 2 2]`, and its
  `= ""` changes nothing.
- The format has no representation of "clear this array". A typed loader must
  treat an empty array key as absent, and must not treat an empty scalar the
  same way: the two are not symmetric. A non-empty string where an array is
  expected is a hard error (`L07428`), not a clear.
- Evidence: EXP-0037 (added 2026-07-27), `evidence/disasm-reginherit.txt`.

**Confidence.** High for the guard, both accessors' absence and presence
behaviour, and the `SetSize` field offsets, read off the instructions. The one
corpus case that could discriminate is pinned to a measured `kind`/`size` rather
than argued over both branches. The reading was not settled from the corpus,
which is consistent with either: `Unit33`'s output is not observable in the
evidence.

**Unknown.** The absent-key residual. `L07429` returns without touching the
destination, and the loader reuses only two `CArray` objects, the one at stack offset +0x20 for
the three `*AnimTime` and the one at +0x3c for the three `*AnimFrame`, draining them
between pairs with `RemoveAt(0,1)`. For an absent key the guard therefore reads
the previous pair's residue, normally 0 but not zeroed by the accessor. Whether
any shipped class reaches a stale nonzero length is not established.

## Object and unit class records

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-OBJ-046 | `R1395` fills a `0x6c`-byte `objects.reg` class record and run-length expands the two kind-6 tracks into an animation timeline it builds rather than stores. | High / Medium | ● active | [EXP-0038](../experiments/EXP-0038-object-frame/) |
| REG-OBJ-047 | `DeadObject` is a class subscript, while `FireObject` is a two-valued flag that is only ever compared, never dereferenced. | High / Medium / Unknown | ● active | [EXP-0038](../experiments/EXP-0038-object-frame/) |
| REG-UNITS-049 | `R0395` fills a `0x12c`-byte `units.reg` class record and builds three animation timelines by run-length expansion; unlike objects, `File` inherits. | High / Medium | ● active | [EXP-0039](../experiments/EXP-0039-unit-frame/) |
| REG-UNITS-050 | `Dying` is a `units.reg` `ID` naming the class whose sheet carries the corpse, and `unit+0x15a` is the corpse stage, which the server's state sync sets. | High | ● active (amended) | [EXP-0039](../experiments/EXP-0039-unit-frame/), [EXP-0071](../experiments/EXP-0071-animation-driver/) |
| REG-UNITS-051 | `Flip` is not a mirror flag on the art: it halves the sheet, and the missing directions are produced at runtime by reflecting the stored half. | High | ● active | [EXP-0039](../experiments/EXP-0039-unit-frame/) |

### REG-OBJ-046

- Fields: `+0x04` `ID`, `+0x08` the loader's own 0-based counter (the `%d` of
  `Object%d`), `+0x0c` `File`, `+0x10` `Index`, `+0x14` `Phases`, `+0x18`
  `Width`, `+0x1c` `Height`, `+0x20` `CenterX`, `+0x24` `CenterY`,
  `+0x28..+0x38` an MFC `CArray<int>` (`vptr`, `m_pData @+0x2c`,
  `m_nSize @+0x30`, `m_nMaxSize`, `m_nGrowBy`), `+0x3c` a copy of `m_nSize`
  (`L07430/L07431`), `+0x40` `FireObject`, `+0x44` `DeadObject`, `+0x48`
  `InMapEditor` (default 0, not −1), `+0x4c` `DescText` read as `0x1f` bytes.
- Every scalar's default is the parent class's resolved field when
  `Parent != -1` (`REG-KEY-045`), except `File`, whose default is the literal
  `-1` unconditionally (`L07432 PUSH -0x1`): `File` does not inherit.
- The array at `+0x2c` is not a registry key. The loader run-length expands the
  two kind-6 tracks into it, appending `AnimationFrame[i]` exactly
  `AnimationTime[i]` times (`L07433..L07434`,
  `CArray::SetAtGrow(m_nSize, B[0])` at `L07435`). Unsigned `JBE`/`JC` make a
  non-positive time contribute nothing; both heads are dropped per round, and
  the loop ends when either track empties. The resulting length is stored at
  `+0x3c`.
- The renderer takes the animation phase modulo that length (`TERR-SPR-042`).
- Corpus: 82 classes, timeline lengths 0 ×52, 24 ×2, 28 ×21, 105 ×7. Only 16
  sections spell `AnimationTime`/`AnimationFrame` and 39 inherit them, so
  `REG-KEY-045`'s array fallback is load-bearing: without it 39 of the 82
  classes lose their timeline and 7 lose `FireObject = -2`.

**Confidence.** High for the field map, the expansion and `File`'s
non-inheritance: each key→offset pair is a `PUSH`/`CALL`/`MOV` triple in a
listing read end to end, the expansion loop is reproduced instruction by
instruction, and `File`'s non-inheritance is one named immediate. Medium for the
timeline-length distribution, a measurement over one registry.

### REG-OBJ-047

- `DeadObject` (`class+0x44`) is an index into the array the type-3 code
  indexes. The map renderer does `classes[class+0x44]` and takes its `File`
  (`L07436..L07437`, twice), and `R0483` tests it to decide whether a
  cell holds a still-standing destructible object (`L05662`).
- `DeadObject` value space over the 82 shipped classes after inheritance: −1 on
  54, and on the other 28 a valid section index: 21 distinct targets, 0 out of
  range, and every one of the 21 has a `DescText` ending "(dead)".
- `FireObject` (`class+0x40`) has one reader in the whole image, `R0484`,
  the ambient-sound loop, and there it is only compared:
  `L07438` compares the field at offset 0x40 of the object with zero and branches when it is less, then `L07439` compares the field at offset 0x40 of the object with -0x2 and branches when it is not equal.
  It is never a subscript, never an operand of arithmetic and never stored.
- The two comparison outcomes feed two counters whose ratio picks the sound
  slot: `L07440` compares the two counters (the second held in a stack local at frame offset -0x6c) against `rand()/(0x7fff/total)` chooses
  slots `0x3c..0x3e` (60–62) when the `>= 0` count wins and `0x46` (70)
  otherwise, played non-looping through `R0386(vol, pan, 0, 0xdc, 0)`.
- `FireObject` value space: `{-2 ×21, -1 ×61}`, never `>= 0`, so slot 70 is the
  only reachable one and the `>= 0` branch is dead on shipped data. The 21
  classes carrying −2 are exactly the 21 that are some other class's
  `DeadObject`, in both directions.
- Neither key occurs in any of the other 43 registries in the install
  (whole-install sweep, `evidence/reg-key-domains.csv`).

**Confidence.** High for the asymmetry, decidable from the instructions alone:
one field is an array subscript, the other appears only in `CMP`, and both
reader sets are whole-image enumerations, which an "is never dereferenced" claim
needs. High for both value spaces, an exhaustive read of the only registry that
carries the keys. Medium that −2 means "burnt-out remains": the biconditional
with `DeadObject` targets is exact and the consumer is a fire ambience, but no
instruction or string names it.

**Unknown.** What `FireObject >= 0` would mean; no shipped class uses it.

### REG-UNITS-049

- the size 0x12c pushed at `L07404`; `0x10c + 0x20` for `DescText` is exactly that size.
- Fields: `+0x04` `ID`, `+0x08` the loader's own 0-based counter (the `%d` of
  `Unit%d`), `+0x0c` `File`, `+0x10` `Index`, `+0x14` `MovePhases`, `+0x18`
  `MoveBeginPhases`, `+0x1c` `AttackPhases`, `+0x20` `DyingPhases`, `+0x24`
  `BonePhases`, `+0x28` `IdlePhases`, `+0x2c` `Width`, `+0x30` `Height`, `+0x34`
  `CenterX`, `+0x38` `CenterY`.
- Timelines: `+0x3c` an MFC `CArray<int>`, the Move timeline (`m_pData +0x40`,
  `m_nSize +0x44`), with its length copied to `+0x50`; `+0x54` the Attack
  timeline with its length at `+0x68`; `+0x6c` the Idle timeline with its length
  at `+0x80`.
- Then `+0x84..+0x90` `SelectionX1/Y1/X2/Y2`, `+0x94` `Dying`, `+0x98` `Palette`
  followed by `Palette` palette objects from `+0x9c` and their `0x400`-byte
  buffers from `+0xac`, `+0xbc` a `CArray` for `Sound`, `+0xd0` `TileSize`,
  `+0xd4` `Projectile`, `+0xd8` `InfoPicture` read as `0x10` bytes, `+0xe8` a
  `CArray` for `ShootOffset`, `+0xfc` `ShootDelay`, `+0x100` `AttackDelay`,
  `+0x104` `Z`, `+0x108` `Flip`, `+0x10c` `DescText` read as `0x20` bytes.
- With `Parent != -1`, each scalar's default is the parent's resolved field at
  the same offset (`REG-KEY-045`). With no parent the fallback is the literal
  `-1`, except `IdlePhases`, `Dying`, `Palette`, `Projectile`, `ShootDelay`,
  `AttackDelay`, `Z` and `Flip`, which default to `0`, and `TileSize`, which
  defaults to `1` (`L07441` supplies the constant 1).
- Unlike `objects.reg` (`REG-OBJ-046`), `File` here inherits: `L07442` loads
  the parent's `+0x0c` as the default.
- The three timelines are the objects loader's run-length expansion run three
  times (`L02603`, `L07443`, `L07444`): append `<track>AnimFrame[i]`
  exactly `<track>AnimTime[i]` times via `CArray::SetAtGrow(m_nSize, …)`,
  dropping both heads each round, ending when either track empties.
- The move arm takes the animation phase modulo the resulting length
  (`TERR-SPR-047`); the idle and attack arms take no modulus.
- Corpus, 34 classes on the EXP-0032 framing: expanded lengths 4…24 for Move,
  8…28 for Attack, 0…14 for Idle (`evidence/unit-timelines.csv`).

**Confidence.** High for the field map, the expansion and `File`'s inheritance:
each key→offset pair is a `PUSH`/`PUSH`/`CALL`/`MOV` quadruple in a listing read
end to end, with key names resolved from the operand addresses
(`evidence/rom-unit-frame-excerpt.md` §5); the expansion loop is reproduced
instruction by instruction; `File`'s inheritance is one named load. Medium for
the length distributions, one registry.

### REG-UNITS-050

- The dying and corpse arms of `R0553` do not merely pick a different
  block; they replace the class: `L07445`/`L07446`/`L07447` load
  `class+0x94` (`Dying`), subscript the global at `L02113` with it (so it is an `ID`,
  `REG-KEY-044`), park it in the slot the sprite lookup later uses, and read
  that class's `File`, `Flip`, `MoveBeginPhases`, `MovePhases`, `AttackPhases`,
  `DyingPhases` and `BonePhases` in place of the unit's own.
- `DyingPhases`/`BonePhases` therefore describe frames in the sheet of whoever
  names this class as their `Dying`, not necessarily the class's own. For 20 of
  34 classes that is itself. All six swordsman/axeman variants and both pikemen
  die as `Human Swordsman with shield` (`ID 4`), both clubmen and both archers
  as `Unarmed Fighter` (`ID 1`), and both mages as `Human Mage` (`ID 24`).
- Every one of the 34 classes has `Dying != 0`, every target resolves to a
  loaded class, and every target has `DyingPhases > 0`
  (`evidence/unit-corpse-chain.csv`).
- `unit+0x15a` is the corpse stage:
  - `0` alive (the standing frame);
  - `1` freezes on the last dying frame, `(dir+1)·DyingPhases` past the dying
    base, index `DyingPhases−1` of the direction's slot
    (`L07448` increments the phase index and `L07449` multiplies it by the value at offset 0x20 of the object);
  - `>= 2` draws bone frame `stage − 2`
    (`L07450` masks the stage with 0xff and `L07451` adds it plus 7 to the running frame index, the `+7`/`+0xe`
    being `S−2`).
- When the target class's `BonePhases` is `0` the whole corpse branch falls back
  to the standing frame (`L02547`).
- At stage `3` two further per-unit passes switch themselves off
  (`R1402`/`R0554`, each comparing the byte at offset 0x15a of the unit with 0x3 and returning when it is not below), and
  CUnit's `vt+0x34` stops at stage `> 1` (`L07452`).
- 5 of the 34 classes carry no `BonePhases`: the 5 that carry `IdlePhases` (the
  bat, the bee, the dragon, the ghost and the Death Star). The bone and idle
  blocks can therefore share a base (`SPR256-UNIT-024`).
- Nothing on the client advances the stage. The whole-image scan for
  `[reg + 0x15a]` finds five writes: three zeroings, the constructor's, and one
  at `L02541` that copies it wholesale from `[L02532]` under bit 3 of a
  mask. That write is in the same routine and the same shape as HP (`+0xfc`),
  mana (`+0x13c`), facing (`+0x6c`) and position (`+0x08`), and the mask is the
  field-present mask of the server's state-sync message (opcodes
  `0x6c`/`0x6e`/`0x6f`/`0x70`, arm `L02530`).
- The value is the simulation actor's `actor+0x13c`, the decay stage
  `HERO-DEATH-026` reads: 1 at death, 2/3/4 as health passes -10/-20/-40, 5
  below -600. No instruction increments it because the stage is not the client's
  to advance (`ANIM-DEATH-007`).

**Confidence.** High for the class substitution, both corpse formulas, the
`BonePhases == 0` fallback and the three stage thresholds, named instructions
(`evidence/rom-unit-frame-excerpt.md` §3, §6, §7); the writer set is a
whole-image enumeration. High for the `Dying` domain and the bone/idle
exclusivity, an exhaustive read of the only registry carrying the keys. High for
what advances the stage: EXP-0071 traced the one non-zeroing write to the state
sync's mask bit 3 and the value to `actor+0x13c`. The enumeration's shape still
cannot see a write through a computed pointer, the same blind spot
`TERR-TILE-044` names.

**Amended.** EXP-0071 answered what advances the stage: nothing on the client
does; the server's state sync writes it.

### REG-UNITS-051

- `class+0x108` is read at five sites in `R0553`, and at none of them is
  it passed to a blitter: every one is a branch selecting between two hard-coded
  layout constants.
- With `Flip == 0` the sheet stores 16 standing frames and 8 directions per
  animated block, and nothing is mirrored.
- With `Flip != 0` it stores 9 and 5, and the missing half is produced at
  runtime by reflecting the stored half: `L07453` compares the frame index with 0x8 and branches when it is less or equal, then
  `L07454` computes 0x10 minus the frame index for the 16-way standing index, and
  `L07455` compares the frame index with 0x4 and branches when it is less or equal, then `L07456` computes 0x8 minus the frame index for the
  8-way ones.
- Each reflection also sets a flag slot to `1` (`L07457`, `L07458`,
  `L07459`, `L07460`, and the same pair inside the corpse arms), handed to
  the sprite blit as its sixth argument (`L07461`, `L07462`); the object
  path passes a literal `0` there (`TERR-SPR-038`, `TERR-SPR-048`).
- The facing is `unit+0x6c`, and the two direction indices derived from it are
  `dir16 = (facing − 8) & 0xf` and `dir8 = ((facing − 8) >> 1) & 7` (`L07463`,
  `L02539`, `L02538`). The standing block is 16-way and every other block
  8-way, on both layouts.
- Corpus: `Flip ∈ {0 ×24, 1 ×10}` over the 34 classes. The constant pairs make
  the sheet-tiling test of `SPR256-UNIT-024` exact on 33 of them.

**Confidence.** High. All five reads, both reflection idioms and both direction
derivations are named instructions, and "never passed to a blitter" is an
enumeration of the field's readers inside the only routine that reads it. High
that the flag reaches the blit as argument six: it is the same slot, pushed
last, at two call sites.

## Restored tables and corpus censuses

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-ROSTER-052 | A `units.reg` class's sprite is `units/` + `Files[class.File]` + `.256`, restoring the pairing `REG-ROSTER-019` owed and refuting the one it published. | High | ● active | [EXP-0043](../experiments/EXP-0043-reg-restorations/) |
| REG-CUT-053 | The 33 video-cutscene registries are one schema with two optional section families, not 33 formats. | High / Medium | ● active | [EXP-0043](../experiments/EXP-0043-reg-restorations/) |
| REG-KEY-054 | The `kind` sorted bit is set on exactly one node per registry, the root, so every lookup below the root is a linear scan. | High | ● active (partially retracted) | [EXP-0043](../experiments/EXP-0043-reg-restorations/) |
| REG-NAME-055 | No shipped registry name reaches 16 bytes and no sibling pair collides under the 15-character compare, so the name clamp never bites on shipped data. | High / Medium | ● active | [EXP-0043](../experiments/EXP-0043-reg-restorations/) |
| REG-KIND-056 | A record's `kind` belongs to the record, not the key name: 9 key names carry more than one kind across the install's 44 registries. | High / Unknown | ● active | [EXP-0043](../experiments/EXP-0043-reg-restorations/) |
| REG-SFX-057 | `sfx.res::sfx.reg` is the sound-slot table: `[Global] SfxCount = 564` is the highest slot id, and `[Sfx]` holds 115 sparse `Sfx<n>` path strings. | High / Medium | ● active | [EXP-0043](../experiments/EXP-0043-reg-restorations/) |

### REG-ROSTER-052

- The index is `REG-VAL-029`'s corrected one, read off the loader at `L07442`
  into `class+0x0c` (`REG-UNITS-049`).
- Emitted for all 34 classes on the EXP-0032 base
  (`evidence/units-roster-restored.csv`): 34/34 carry their own `File` and
  `DescText` keys, so nothing inherits here; 34/34 resolve to a `.256` node that
  exists in `graphics.res`; `File` equals the section's own `%d` on 33/34.
- The exception is `Unit33` "Unarmed Fighter with Shield", `File = 21`, which
  shares `Unit21` "Unarmed Fighter"'s sheet: the only shared entry in the table.
- Measured against the withdrawn `Files[class.ID]` pairing on the same walk, the
  two differ on 34/34, and `ID` is inside the 33-entry table on only 20/34. 14
  of the published arrows indexed past the end of `Files[]`, and the other 20
  named another class's sheet (`Human Archer` `ID` 14 → `monsters\legg\sprites`,
  `Human Mage` `ID` 24 → `monsters\ghost\sprites`).
- The old pairing was refutable from the shipped bytes alone at any time since
  EXP-0006: no framing correction was needed to see 14 out-of-range subscripts.

**Confidence.** High. The resolution rule is code-read and already published
(`REG-VAL-029`, `REG-UNITS-049`). This claim applies it mechanically and adds
three discriminators against the alternative: 14/34 of the old arrows out of
range, 34/34 of the new ones opening a node that exists, and a render that draws
frame 0 of both candidate sheets for every class, where a wrong arrow is a
picture that does not match its own `DescText`.

### REG-CUT-053

- Grouping all 44 registries by a schema signature that strips digits from every
  section and key name yields 14 distinct shapes. The 33 cutscene files occupy 3
  of them, each a subset of the next:
  - `Common{startx,starty,nFadings}` +
    `Fading<n>{startframe,endframe,startfade,endfade}` ×27;
  - the same plus `Common.nPanaramings` +
    `Panaraming<n>{startframe,endframe,stepx,stepy}` ×4 (`M10/01`, `M50/01` in
    both containers);
  - `Common` alone with `nFadings = 0` ×2 (`LOGOS/1c.reg`, `LOGOS/buka.reg`), a
    degenerate instance of the first, not a different format.
- Every declared count agrees with its own section count: `nFadings` 33/33,
  `nPanaramings` 4/4. `startx == starty == 0` on 33/33. `startframe <= endframe`
  on 65/65 sections.
- The kind-4 census reproduces `REG-DBL-035`: `startfade`/`endfade` take only
  `0.0` (×61) and `1.0` (×61), with `Fading1` always fading in and the last
  `Fading` always out.
- Each registry sits beside a `.smk` node of its own basename in its own
  container directory, 33/33. The 15 node paths present in both `VIDEO4.RES` and
  `VIDEO8.RES` have byte-identical resolved content, 15/15, while their `.smk`
  siblings differ in size: the two containers are two encodings of one set of
  cutscenes sharing one timing sidecar.

**Confidence.** High for the schema and the counts: a listing of every cutscene
registry in the install and every section, with the two declared counters
cross-checked against counts derived by walking section names, a different path
from the counters. Medium that these values are playback timing for the sibling
video: the 1:1 file pairing, the fade-in/fade-out shape and the pan step are
strong, but `R1399`, the consumer `REG-LOC-038` names, was not read.

### REG-KEY-054

- `REG-KIND-033` pinned bit 4 to the instruction that tests it (a test of bit 0x10
  at `L07398`: `bsearch` when set, linear scan when clear). This claim is the
  census of which nodes set it.
- Over all 44 registries the root's `kind` is `17 = subkey|sorted` on 44/44, and
  0 of the 512 subkey records in the corpus set bit 4. The `bsearch` path is
  reachable only for a top-level section name; every per-key lookup inside a
  section is a scan.
- Sections are stored in lexicographic, not numeric, order
  (`MapObject1, MapObject10, …, MapObject2`), so the order in which a loader's
  `"<Prefix>%d"` counter visits sections (`REG-KEY-044`) is not the record
  order.
- A writer advertising bit 4 must emit children ordered for the actual
  case-sensitive comparator (`REG-100`). With bit 4 clear, the reader scans
  linearly.
- An earlier census found root lists sorted under its case-insensitive ASCII
  15-character comparison on 44/44, while only 81 of 512 non-root lists passed.
  Those ordering counts stay scoped to that instrument; no case-sensitive census
  replaces them.

**Confidence.** High. An exhaustive census over all 4 621 records: the
denominator is the population, and the instrument is a corpus census, the immune
one. The consuming instruction is `REG-KIND-033`'s, not re-read here.

**Amended.** `REG-100` partially retracts the comparator identity and the
unconditional writer obligation. The root lists were described as sorted under
the lookup's own comparator, case-insensitive ASCII over 15 characters, with
every writer bound to emit root children sorted. The actual sorted comparator is
case-sensitive byte comparison through NUL, and sorting is a data obligation
only when bit 4 is advertised. The root/subkey flag population and the
numeric-versus-lexicographic ordering stand ([`retracted.md`](retracted.md)).

### REG-NAME-055

- `REG-REC-032` fixed the name field at 16 bytes with 15 significant characters,
  from the writer's clamp (`R1393`) and the lookup's
  `_strnicmp(key, node+0x10, 0xf)`.
- Census over all 4 621 records: name lengths run 1..15, 0 records are 16 or
  longer, and 69 are exactly 15, the length at which a longer authored name
  would have been cut.
- The falsification the clamp invites fails: 0 sibling pairs anywhere in the
  install are indistinguishable under the lookup's 15-character case-insensitive
  compare.
- Of the 6 fifteen-character keys in the registries EXP-0043 tabulates,
  `rom.exe` carries a longer identifier sharing all 15 characters for 3, which
  recovers the author's own name: `MinimalGuardRan` → `MinimalGuardRange`,
  `AddPictureDocum` → `AddPictureDocument`, `ScenarioMission` →
  `ScenarioMissionCount`.
- The other three (`AddTextDocument`, `EnableMercenary`, `IntelligentCons`) have
  no longer literal, so whether they were truncated at all is undecidable from
  the image.

**Confidence.** High for the length histogram and the collision hunt, exhaustive
over the corpus. Medium for the un-truncation: a printable-ASCII-run census over
`rom.exe`, immune to the function-table blind spot, shows only that the image
contains the longer string; no call site was read pushing it to `R0452`.

### REG-KIND-056

- Four of the 9 are the empty-string-versus-array pair `REG-KEY-045` explains,
  where a present-but-empty kind-0 record reaches array length 0 by a success
  path: `units.reg` `AttackAnimTime`/`AttackAnimFrame` (kind 0 ×1, kind 6 ×32)
  and `structures.reg` `AnimTime`/`AnimFrame` (kind 0 ×52, kind 6 ×14).
- The remaining five occur only in `scenario.reg`, where a key is a bare int32
  in some `[Mission<n>]` sections and an int32 array in others: `Mercenaries`
  (kind 2 ×3 / kind 6 ×12), `InnNPC` (5/8), `InnMission` (5/8),
  `EnableMercenary` (4/1), `AddTextDocument` (1/1).
- A typed reader for these keys must accept a bare scalar where an array is
  expected and treat it as a one-element list. Reading only kind 6 loses 18 of
  the 31 `InnNPC`/`InnMission` records.

**Confidence.** High for the census over every leaf record in the install; the
denominator is the population.

**Unknown.** How `rom.exe`'s array accessor `R1061` treats a kind-2
record. `REG-KEY-045` read its kind-0 and kind-6 arms; one `DisasmFn` on
`R1061` settles it. The one-element-list reading is what the shipped data
forces, not what an instruction was seen to do.

### REG-SFX-057

- Two sections: `[Global] SfxCount = 564` and `[Sfx]` holding 115 kind-0 strings
  named `Sfx<n>`.
- `SfxCount` is the highest slot id, not the entry count. The ids are sparse
  (1..8, 11..16, 50, 60..62, 70, 80, 81, 90, 100, 110, …, 564), and 564 is the
  largest present: the same shape as `Templates.ini`'s own `Number = 1088`
  against 165 labelled slots.
- Each value is a path relative to `sfx.res` with no extension (`click00`,
  `units\sword`, `ambient\river`, `magic\firewall`).
- Every nonzero element of a `units.reg` class's `Sound[]` array is an existing
  slot, 153/153, with 17 further elements reading `0`.
- `REG-OBJ-047`'s ambient-sound slots resolve: the `>= 0` arm's `0x3c..0x3e` are
  `ambient\bird1/2/3` and the `-2` arm's `0x46` is `ambient\crow`.

**Confidence.** High for the table, an exhaustive read of the registry; both
resolution censuses carry their denominators, and the `Sound[]` one covers every
class in `units.reg`. Medium that slot 70 being a crow corroborates
`REG-OBJ-047`'s Medium reading of `FireObject = -2` as burnt-out remains: the
slot numbers are code-read there and the sample name is now known, but nothing
in the image names the intent.

## Campaign registries: `npc.reg`, `scenario.reg` and `ai.reg`

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-NPC-058 | `scenario.res::npc.reg` holds 105 `npc<n>` sections, four archetype blocks and a `Multiplayer` block; `Flags` is an 11-token list and `DataBinID` a `Templates.ini` id. | High / Medium / Unknown | ● active (amended) | [EXP-0043](../experiments/EXP-0043-reg-restorations/), [EXP-0062](../experiments/EXP-0062-mercenaries/), [EXP-0141](../experiments/EXP-0141-npc-fields-and-tag/) |
| REG-SCN-059 | `scenario.reg` and `globalmap.reg` are the campaign spine: 24 mission sections and 35 map objects, and every id in them resolves, including 143/143 npc references. | High / Unknown | ● active (amended) | [EXP-0043](../experiments/EXP-0043-reg-restorations/), [EXP-0062](../experiments/EXP-0062-mercenaries/) |
| REG-AI-060 | `world.res::data/ai.reg` is the install's smallest registry, 4 records and two int32 keys, and its `[Tasker]` half is named by no literal in `rom.exe` or `Map Editor.exe`. | High / Medium | ● active | [EXP-0043](../experiments/EXP-0043-reg-restorations/) |

### REG-NPC-058

- 110 top-level sections:
  - 105 `npc<n>` (ids 1..132, sparse);
  - four character-archetype blocks
    `MaleFighter`/`MaleMage`/`FemaleFighter`/`FemaleMage`, each carrying exactly
    `Face Body Reaction Mind Spirit Skill` as int32;
  - `Multiplayer`, carrying four kind-6 face-id lists `FacesMM`(6)
    `FacesMF`(15) `FacesFM`(4) `FacesFF`(8).
- Per-npc keys, with presence out of 105 and measured domain: `Flags` 105/105
  (kind 0), `Face` 86 (1..30), `PortraitX1` 48 (7..69), `PortraitY1` 48
  (8..72), `PortraitX2` 10 (48..54), `PortraitY2` 10 (15..32), `Picture` 33
  (64..80), `DataBinID` 23, `PriceA` 13 (28..80), `PriceB` 13 (0..15).
- `Flags` is a comma-separated token list with `!` negation over exactly 11
  tokens: `Human, Mage, Female, Face, Picture, Platoon, Hero, Me, Start, MySex, MyClass`.
  There are 15 distinct combinations, and all 11 tokens occur as whole ASCII
  strings in `rom.exe`.
- `DataBinID` is an id in `Templates.ini`'s space (`EDITOR-023`'s 1088-slot
  placeable-template table, the install's own text). All 23/23 lie inside the
  declared `1..Number=1088`. The 13 at or above the table's lowest labelled
  slot (103) name a labelled slot on 13/13: `509 → M10_Witch`,
  `510 → M10_Merchant`, `511 → M30_Healer`, `523 → M40_Hima`,
  `520 → M130_Veglud`, `521 → M151_Lord`, `522 → M71_PoorGuy`,
  `1011 → A_PeasantGuard3`, `1012 → U_PeasantLeader3`,
  `1028 → F_BrigandLeader3`, `1036 → M_EvilFemale2`, `1051 → F_Knight3`,
  `1056 → F_Knight4`.
- Full table: `evidence/values-scenario_npc.csv`.
- `PriceA`/`PriceB` (EXP-0057): the loader `R0499` stores them at
  `npc+0x18` / `npc+0x1c` of the `0x30`-byte record in the array
  the global at `L03415`. `EnumRefs refto:L03415` returns 20 hits, 7 owners, 0 in
  orphan/undisassembled code, and a `DisasmFn` pass over all six non-loader
  owners finds no `[reg+0x18]`/`[reg+0x1c]` operand. They are not the shop's:
  nothing in the shop machinery touches the npc array (`SHOP-NPC-012`).
- They are the two terms of the mercenary hire price,
  `cost = (PriceA + n·PriceB) × unitPrice(mission)`, read by `R1142`
  (`L07464`…`L07465`) and by the object serializer `R0059`, both
  through `R1285(L02112, id)`: the array object, not its data pointer
  (`MERC-PRICE-004`). The 13 npcs that carry them are exactly the 13 mercenary
  types with a nonzero `[General] MercenaryCount`, 15/15.

**Confidence.** High for the tables, the presence counts and the `Flags`
vocabulary: exhaustive over the registry. Medium that `DataBinID` is a
Templates.ini id: 13/13 above slot 103 is corpus agreement, not a discriminator,
and the remaining 10 entries, `26` on `npc21`..`npc24` and `42`..`46` on
`npc25`..`npc30` with `46` on both `npc29` and `npc30`, are all below the
table's lowest labelled slot and name nothing shipped.

**Unknown.** No gloss is put on `Body`/`Mind`/`Reaction`/`Spirit`/`Skill`
beyond their measured values.

**Amended.** The `Portrait*` presence count read
"`PortraitX1/X2`+`PortraitY1/Y2` 58", the sum of two of the four counts and not
any key's own count; EXP-0141 replaced it with the four figures above
([`retracted.md`](retracted.md), NARROWED). The split is the corpus side of
`X2`/`Y2` having no reader (`REG-NPC-089`). EXP-0057 had described `EnumRefs
refto:L03415` as enumerating every route to a record and located no reader of
`PriceA`/`PriceB`; EXP-0062 (`MERC-PRICE-004`) found the readers, which go
through the array object and which a `refto:L03415` sweep therefore could not
see. EXP-0141 discharges the `Face`/`Picture`/`Portrait*` Unknown
(`REG-NPC-088`, `REG-NPC-089`): none of them indexes a picture resource;
`Picture` and `Face` are stores into a synthesised actor's typeID and face, and
`Portrait*` is a Win32 `RECT`. The Confidence paragraph read "the remaining 10
values — `26`×4 and `42`..`46`", nine integers for ten entries;
`evidence/values-scenario_npc.csv` gives `46` twice
([`retracted.md`](retracted.md), CORRECTED).

### REG-SCN-059

- `scenario.res::scenario.reg`: `[General] TotalMissions = 15`,
  `ScenarioMissionCount = 15` (name un-truncated from `rom.exe`,
  `REG-NAME-055`) and `MercenaryCount`, a 15-element kind-6 array.
- 24 `[Mission<n>]` sections: the 15 main missions numbered by tens
  (`Mission10`..`Mission150`) plus 9 sub-missions.
- Per-mission keys and domains: `Mercenaries` 15, `InnNPC` 13, `InnMission` 13,
  `ShopMinPrice` 13 (0..5000), `ShopMaxPrice` 13 (1000..10 000 000), `Payment` 9
  (700..700 000), `EnableMercenary` 5, `ShopMission` 4, `TCMission` 3,
  `AddTextDocument` 2, `AddHero`/`AddPictureDocument`/`AutoGetMission`/`LastMission`
  1 each.
- `scenario.res::globalmap.reg`: `[General] ObjectCount = 35` matches 35
  `[MapObject<n>]` sections (1-based, 1..35), each carrying `MapPoint` = 2 ints
  and `MapRect` = 4 ints, and 5 of them a `Picture` string (`onmap13`,
  `onmap14`, `onmap18`, `onmap19`, …). `[MissionObjects]` holds 28
  `Mission<n>` keys mapping a mission number to a map-object number.
- Resolution censuses, each over the whole registry:
  - every `MissionObjects` value names an existing `MapObject<n>`: 28/28;
  - every npc-shaped reference in `scenario.reg` names an existing `npc.reg`
    `[npc<n>]` section: `InnNPC` 22/22, `AddHero` 1/1, `Mercenaries` 107/107,
    `EnableMercenary` 13/13, i.e. 143/143;
  - every `scenario.reg` mission has a `MissionObjects` entry: 24/24. Four
    entries there (`81`, `101`, `141`, `151`) name no `scenario.reg` section.
- `R1297` is the mission-record loader, and a `FindStrXref` sweep makes
  it the only referrer of every per-mission key literal:
  - `ShopMinPrice` → `record+0x10`, `ShopMaxPrice` → `record+0x14`
    (`SHOP-MISSION-018/019`);
  - `ShopMission`/`InnMission`/`TCMission` → three `u16` arrays merged into one
    sub-mission list (`REG-SCN-062`);
  - `Payment` → `+0x0c`, `AutoGetMission` → `+0x110`, `LastMission` →
    `+0x11c`, `Mercenaries` → `+0x98`, `EnableMercenary` → `+0x30`, `InnNPC` →
    `+0xc0`, `AddHero` → `+0x1c`;
  - `AddTextDocument`/`AddPictureDocum` element-wise into
    `R1403`/`R1404`.
- `ScenarioMissionCount` is read by the same routine as the campaign's
  mission-number bound: `L07466` reads it (default −1) and
  `L07467…L07468` rejects any request above `10 × count`.
- `MercenaryCount`'s reader is `R0757`, the record's reset
  (`L07469`/`L07470`). It copies the 15 `u16` elements to `record+0x60`,
  beside a parallel 15-`dword` array at `+0x88` that it zeroes. `record+0x60`
  is the working copy's `m_pData`, and the `+0x88` array is the per-type hire
  flag.

**Confidence.** High for the tables and all five resolution censuses:
exhaustive, each with its denominator. The npc census is the kind that fails
visibly if the id space were wrong: 143 subscripts into a 105-member sparse
set. High for what `MercenaryCount` counts (EXP-0062, `REG-SCN-067`): the 15
elements are one per mercenary type 1..15, a pool count, not one per main
mission, which is why the per-mission reading failed 12/15. All three
consumers are read at instruction level.

**Unknown.** Whether `TotalMissions` is read at all: no printable-ASCII run in
`rom.exe` or `Map Editor.exe` equals it, while its neighbour
`ScenarioMissionCount` is present, holds the same 15, and is read as the
mission-number bound.

**Amended.** EXP-0059 located the key readers this claim had left unlocated,
all in `R1297`. `MercenaryCount`'s reader was first located but not
interpreted; `REG-SCN-067` (EXP-0062) answers that Unknown with the per-type
reading, the working array and the hire flag above.

### REG-AI-060

- 4 records, empty pool, two sections of one int32 each:
  `[Scanning] MinimalGuardRan = 8` and `[Tasker] IntelligentCons = 15`. Both
  names are exactly 15 characters, i.e. clamped (`REG-NAME-055`).
- `rom.exe` carries `MinimalGuardRange` and `Scanning` as whole strings, which
  recovers the first key's authored name.
- Neither `Tasker` nor any printable-ASCII run beginning `IntelligentCons`
  occurs anywhere in `rom.exe` or in `Map Editor.exe`, so no site in either
  executable can look that key up by a literal name.

**Confidence.** High for the content: four records read in full. Medium for
the absence. The instrument is a printable-ASCII-run census over the
executables in hand; it involves no disassembly, so it is immune to the
function-table blind spot, but it is blind to a name assembled at run time or
stored non-contiguously. A cross-reference sweep on
`"Scanning"`/`"MinimalGuardRange"` and their loader would settle whether the
section is read at all.

## `units.reg` `Z` and the unit class

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-UNITS-061 | `units.reg` `Z` picks a unit's C++ class, `CUnit` or `CAirUnit`, and the class's registration override changes the unit's draw plane. | High / Medium | ✔ promoted (partially retracted) | [EXP-0044](../experiments/EXP-0044-mover-domain/), [EXP-0333](../experiments/EXP-0333-drawable-cell-passes/) |

### REG-UNITS-061

- `R0509` allocates `0x1b0` unit objects at two sites. `L07471` loads the `units.reg` class array
  `L02113`, `L07472` subscripts it, and
  `L07473` compares the 32-bit field at offset 0x104 of the class object with zero, testing `Z` (`class+0x104`,
  `REG-UNITS-049`). The branch chooses `CUnit` (`Z == 0`, `CALL R0593`) or
  `CAirUnit` (`Z != 0`, `CALL R0594`).
- The two classes are MFC `CRuntimeClass`-registered at `L00620`/`L07474`,
  both `0x1b0` bytes, `CAirUnit` derived from `CUnit`. Their vtables differ in
  four of 33 slots; three are `GetRuntimeClass`, the destructor and the
  class-name stub.
- The single behavioural override is `+0x38`. `CUnit` (`R1405`)
  registers the sprite in draw layer 2 when the corpse stage
  `[this+0x15a] < 2` and `[this+0x18c] & 0x80` is clear, else layer 4.
  `CAirUnit` (`R1406`) registers layer 3 unconditionally and carries no
  corpse arm. Both do additional rectangle-helper work after registration; that
  call is not a sprite blit.
- Registration selectors 2/3/4 mean view+0x8c/+0x90/+0x9c, respectively. The
  literal `CUnit` condition above, not an alive/dead shorthand, determines its
  plane. `CAirUnit` construction separately writes +0x10=16. Shared body/shadow
  methods do not imply one pass population (ANIM-CATEGORY-084,
  ANIM-AIRPASS-086).
- `Z` is therefore a z-order/altitude flag on the drawable, not a movement
  property: it never reaches the simulation actor's `+0x4a` (`TERR-MOVE-055`),
  and the simulation module never reads the class array at all.
- Corpus, 34 classes on the EXP-0032 framing with `Parent` resolved: `Z != 0`
  on exactly 2, `Sonic Bat` (ID 70) and `Dragon` (ID 71), both `Z = 96`. Every
  other class carries `Z = 0` or no `Z` key (default 0). Their 715 placements
  over the 38 shipped maps account for 45 of the 263 that EXP-0041 found
  stranded under the ground mask.

**Confidence.** High for the test, both allocation arms and the one
behavioural override: named instructions, with the class identities taken from
the `CRuntimeClass` records rather than from inference. Medium for reading `Z`
as "flying": layer 3 plus the two class names carry it, and the value 96 is
otherwise unexplained. EXP-0078 adds two independent witnesses, neither from
the drawable side: the `Z != 0` pair is the `Data.bin` `movementType == 3` pair
(`MOVE-DOM-028`), and the AI's target chooser adds one cell to a candidate's
distance when its `vt+0x20` is 3 and the decider's is not
(`L00225…L00226`, `MOVE-DOM-024`(e)), the image's own arithmetic for a
flier. The value 96 is still unexplained.

**Amended.** The registration selector shorthand and helper-blit wording are
narrowed ([`retracted.md`](retracted.md), EXP-0333, `ANIM-CATEGORY-084`). The
former wording was "All the class change buys is the draw layer; both
registration methods then blit through the rectangle helper". It is replaced
by the selector destinations, the literal stage/flag predicate, the
`CAirUnit` `+0x10` write and the non-blit helper stated above; common drawing
bodies do not merge the populations. The `Z`/corpus and simulation-domain
observations are not remeasured or upgraded.

The clause "only site in `rom.exe` that allocates a `0x1b0` unit object" is withdrawn
(`retracted.md`, `UNIT-141`): five sites in four routines allocate `0x1b0`, and the others
build `CUnit`. The `CAirUnit` arm also requires the creation mask to lack the hero range
(`UNIT-141`).

## Mission records, routing and town offers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-SCN-062 | `ShopMission`, `InnMission` and `TCMission` are lists of mission numbers that one loader turns into one shared sub-mission array; `ShopMission` is not a shop parameter. | High / Medium / Unknown | ● active (amended, partially retracted) | [EXP-0059](../experiments/EXP-0059-shop-mission/), [EXP-0060](../experiments/EXP-0060-town-screen/) |
| REG-SCN-063 | `AutoGetMission` and `LastMission` are the campaign's routing keys, and `[Mission10] AutoGetMission = 20` is what keeps the player out of the town after missions 10 and 20. | High / Unknown | ● active | [EXP-0060](../experiments/EXP-0060-town-screen/) |
| REG-SCN-064 | `InnMission`, `ShopMission` and `TCMission` are the offer lists of the tavern, shop and school, each with exactly one reader, and `InnNPC[i]` pairs with `InnMission[i]`. | High / Medium / Unknown | ● active (amended, partially retracted) | [EXP-0060](../experiments/EXP-0060-town-screen/), [EXP-0394](../experiments/EXP-0394-inn-selection-loop/EXP-0394.md) |
| REG-SCN-065 | Sub-missions are appended at each main-mission load and dropped after one further load or on completion; an entry is offered, not only listed, once its `+0x18` latch is set. | High | ● active | [EXP-0060](../experiments/EXP-0060-town-screen/) |
| REG-GMAP-066 | `Scenario\GlobalMap.reg` is keyed `[MissionObjects] Mission<n>`, read as a 0-based object index; the city is `MapObject1` (index 0), which no mission occupies. | High / Medium | ● active (amended, partially retracted) | [EXP-0060](../experiments/EXP-0060-town-screen/) |
| REG-SCN-067 | `MercenaryCount`, `Mercenaries` and `EnableMercenary` share one 1..15 mercenary-type space, and `[Mission<n>] Payment` is the mission's cash reward. | High / Medium | ● active | [EXP-0062](../experiments/EXP-0062-mercenaries/) |

### REG-SCN-062

- `R1297` reads all three keys into `u16` arrays at `record+0xd4` /
  `+0xe8` / `+0xfc` (`L07475` / `L07476` / `L07477`), then runs three
  consecutive loops over them that are identical instruction for instruction.
- Each element is tested `value % 10 != 0`
  (the sequence clears the high half, loads the 16-bit element, sign-extends, divides by 10, tests the remainder and branches when it is zero). A
  multiple of ten is dropped. Any other value has its own `[Mission<n>]`
  section opened for `AddHero`, `Payment` and `EnableMercenary`, is given its
  `GlobalMap.reg` `[MissionObjects]` object via `R1324`, and is appended
  as a 0x4c-byte sub-mission record to one shared CArray at `record+0x48`.
- No instruction records which of the three arrays an entry came from. The
  sub-record's `+0x10`/`+0x14`, where a shop window would sit, are never
  written by any of the three loops.
- These records are the campaign's offered side missions. `R1407(n)`
  routes `n % 10 == 0` to load mission n (`R0755`) and everything else
  to `R0756`, a linear search of the array by the record's own `+0x04`.
  `R1408` returns the next entry whose `+0x18` is nonzero; it is called
  from `R1409`, which holds a mission record at its own `+0x548`.
- Each entry carries a counter at `+0x48` that every mission load increments
  (`L07478…L07479`), and is deleted at 2 (`L07480`).
- Corpus, identical in all three roots (`EXP-0059/evidence/submissions.csv`):
  `ShopMission` ships 4 elements over 24 sections, `Mission30 → 31`,
  `Mission90 → 91`, `Mission110 → 110`, `Mission130 → 130`, so two of the four
  are multiples of ten. The same test drops 14 of `InnMission`'s 22 elements
  and 0 of `TCMission`'s 3.
- The value a building hands to `R0785` is unfiltered. Its
  `R1407` step routes `n % 10 == 0` to `R0755(n)`, which accepts
  `n` whenever `n >= record+0x04`, after which `R1410` latches the main
  record's `+0x18`. A multiple of ten is how a building hands the player the
  main mission: the inn does it in 11 of the 13 town chapters and the shop in
  the other two (`REG-SCN-064`).
- Four admitted values (`81`, `101`, `141`, `151`) name no `[Mission<n>]`
  section, so their sub-record takes every default. They are the same four
  `REG-SCN-059` found dangling in `[MissionObjects]`.

**Confidence.** High for the loader, read at instruction level, and the three
loops, which are the same instructions on three arrays; the corpus census is
exhaustive over the registry and reproduces in all three roots. Medium that the
array is the town's offer list rather than some other queue:
`R0756`/`R1408`/`R1407` are read and their callers
enumerated, but the screen that renders the list was not read.

**Unknown.** Whether anything distinguishes an inn-, shop- or TC-supplied entry
downstream. An `EnumRefs disp:` sweep over the record's class range
`L07481..L07482` (75 functions, 0 orphan bytes, 1093 undisassembled) finds
only the ctor, the copy/destroy pair, the reset, this loader and the two
streaming routines touching `record+0xe8`, so nothing outside the class reads
the ShopMission array; the instrument is blind to an access through a pointer
laundered out of the class. EXP-0060 answers it: the sweep was complete and the
question was asked in the wrong place. Provenance is not in the sub-mission
record but in which array the screen reads: `InnMission` is read only by the
inn view, `ShopMission` only by the shop view, `TCMission` only by the school
view (`REG-SCN-064`).

**Amended.** The clause that the two multiples of ten "are therefore inert" is
retracted by EXP-0060 ([`retracted.md`](retracted.md)): the `% 10` test
governs only whether a value becomes a sub-mission record, and the unfiltered
building route above hands a multiple of ten on as the main mission.

### REG-SCN-063

- `R1297` reads `[Mission<n>] AutoGetMission` into `record+0x110` with
  default -1 (`L07483` the literal, `L07484` the store), and `LastMission`
  into `record+0x11c` with default 0 (`L07485` / `L07486`).
- `record+0x110` has exactly one reader, `R1411`, itself called from
  exactly one site (`L07487`, the mission-end arm of `R0701`). Its two
  arms are the campaign's fork:
  - `!= -1` calls `R0755(that mission)`, which loads that mission's
    record, then `R1410(record+0x04)`, which latches the announce flag;
    the party travels to the new mission's own map object;
  - `== -1` returns 0, and the caller sends the party to global-map object 0,
    the city.
- Corpus, identical in all three roots (`EXP-0060/evidence/missions.csv`):
  `AutoGetMission` ships once in the whole campaign, `[Mission10] = 20`, and
  `LastMission` ships once, `[Mission150] = 1`. So missions 10 and 20 are the
  only two whose end does not pass through the town, which is what makes
  `SHOP-TOWN-023` true.

**Confidence.** High for both readers, both defaults and the fork's two arms,
read at instruction level; `callto:R1411` returns 1 hit, 0 in orphan code.

**Unknown.** What reads `record+0x11c`: the `LastMission` store is read at
instruction level, its consumer was not searched.

### REG-SCN-064

- The record is embedded at `campaignScreen+0x548`, so the four arrays sit at
  `+0x608` (`InnNPC`), `+0x61c` (`InnMission`), `+0x630` (`ShopMission`),
  `+0x644` (`TCMission`) of the campaign screen.
- An image-wide `EnumRefs disp:` over all four objects and their data/size
  fields, plus `imm:630` for the two forms that add 0x630 to the base pointer, returns 0 hits in
  orphan or undisassembled code and these readers:
  - `ShopMission`: `R0705` only, the shop view's `vt+0x80`;
  - `TCMission`: `R0706` only, the school view's `vt+0x80`;
  - `InnMission`: `R0702` only; `InnNPC`: `R1412` +
    `R0702`, both inn-view methods.
- Shop and school each take element 0 on entry, hand it to `R0785` and
  `RemoveAt(0,1)` the array (`L07488`…`L03652`, `L07489`…`L03655`).
- The inn is per-NPC only for selections at or past a bound; below it a
  sibling arm reaches the mercenary roster instead (`REG-118`; the bound is
  under Amended). `R0702` reads `InnMission[i]` and `InnNPC[i]` at the
  same subscript (`L07490`…`L07491`), speaks `inn\NPC\npc%02dm%d`, and
  queues the mission. `R1413` later removes the accepted value from both
  arrays at the same index (`L07492` `InnMission`, `L07493` `InnNPC`).
- So `InnNPC[i]` pairs positionally with `InnMission[i]`, which closes
  `REG-SCN-059`'s unpaired `InnNPC`.
- `InnMission[i] == 0` is the sentinel for an NPC with nothing to give: the
  zero arm speaks a line keyed by the live main mission and queues nothing
  (`L03600`, `L03601`).
- `TC` is the school: the view that reads `TCMission` is the one
  `music\schoolw.wav` / `music\schoolm.wav` and `SFX\Town\School\Point.wav`
  belong to.
- Corpus, all three roots identical (`EXP-0060/evidence/offers.csv`,
  `timeline.txt`): the three keys name 26 distinct missions over the 13 town
  chapters, each by exactly one building and none by two: tavern 19, shop 4,
  school 3. They cover every main mission 30…150 and every side mission
  31…151 with no gap.
- The tavern's 19 is its 22 `InnMission` elements minus the 3 that are `0`,
  the no-mission sentinel (`[Mission50]`, `[Mission110]`, `[Mission130]`, each
  at index 0); 19 + 4 + 3 = 26. `len(InnMission) == len(InnNPC)` in 22/22
  sections.

**Confidence.** High for each key's single reader, established by an
image-wide displacement enumeration on the repaired table, with each consuming
routine read at instruction level; the pairing and the zero sentinel are named
instructions. Medium that no fourth building exists: the three views were
reached from the campaign screen's slot table and the town's own five-way hit
test, not from an exhaustive sweep of every view class.

**Unknown.** What happens when one mission number appears in two of the three
keys; the corpus never does it. From the code it would be appended twice, since
both loops append unconditionally and no dedup exists, and resolved once by the
acceptor. That is a reading of the loops, not an observation.

**Amended.** The REG-118 correction pass (2026-09-22) scoped "per-NPC" and
`R0702`'s single-reader count to `curSel >= InnNPC.Size()`: `REG-118`
finds `R0702` entered only above that bound, and below it "a sibling arm
reads `InnNPC.m_pData` through the embedded `record`'s own `+0xc4`
displacement", a route this claim's campaign-screen-relative `EnumRefs` sweep
never searched, reaching the mercenary roster `MERC-HIRE-003`/`MERC-TYPE-001`
publish rather than a second per-NPC read. That amendment clause, with the
wording "the inn is per-NPC for selections at or past `InnNPC`'s own size", is
partially retracted ([`retracted.md`](retracted.md), EXP-0400, `TOWN-468`): the
bound is the inn view's mercenary count `view+0xc8`, and the arm below it reads
the view's mercenary `CUnit*` array `view+0xc4`, not `InnNPC`; the per-NPC
reader serves positions from the mercenary count on. The offer lists, their
single readers, the pairing, the sentinel and the corpus counts stand.
Separately, the tavern figure was published on 2026-07-30 as "tavern 21",
neither the element count 22 nor the offer count 19; it was repaired in place,
with no mechanism change and no `retracted.md` row.

### REG-SCN-065

- Appended at every main-mission load by the three `% 10 != 0` loops of
  `R1297` (`REG-SCN-062`), as 0x4c-byte records with vptr `L07494`:
  mission at `+0x04`, map object `+0x08`, `Payment` `+0x0c`, announce latch
  `+0x18`, `AddHero[]` `+0x1c`, `EnableMercenary[]` `+0x30`, owning record
  `+0x44`, age `+0x48`.
- Aged and dropped: the pass at `L07495`…`L07496`, which runs before the
  appends, increments every surviving element's `+0x48` and deletes it once the
  counter reaches 2. An offer lives through the load that created it and one
  further main-mission load, and no longer.
- Removed on completion: `R1414` deletes the finished sub-mission by
  `memmove` and decrements `record+0x50`.
- Offered: the latch `+0x18`, whose only writer is `R1415` (`L07497`,
  0 to 1, guarded by its own current value), reached only through
  `R1410`. `callto:R1415` = 1 hit; `callto:R1410` = 2 hits
  (`R0785`, `R1411`).
- Consumed: `R1416` resets the cursor `record+0x44` and yields the main
  mission first if the record's own `+0x18` is set; then `R1408` walks
  the array yielding only latched entries. `-1` means nothing is on offer.
- Bounded: `R1297` rejects any mission number above
  `10 × [General] ScenarioMissionCount` (`L07466`…`L07468`).
- `R1417`, which also reads `+0x18`, has 0 callers and 0 `.rdata`
  slots: dead code.

**Confidence.** High. Every step is a named instruction, and the two latch
enumerations ran on the repaired table with `callto` covering `.rdata` slots,
so virtual dispatch is included. They are blind to a call through a register
Ghidra resolved to no target.

### REG-GMAP-066

- The key is `[MissionObjects] Mission<n>`, not `[Mission<n>] MissionObjects`.
  `R1324` formats `"Mission%d"`, then pushes `"MissionObjects"` last
  (`L07498`), so that is the section, with default 1 (`L07499`). It returns
  the value minus one (`L07500`, a decrement): a 0-based index into the global-map
  view's object array at `view+0xc4`.
- The record caches it at `record+0x08`; a sub-mission record caches its own at
  `elem+0x08`.
- `[General] ObjectCount = 35` matches 35 `[MapObject<k>]` sections, 1-based,
  each with `MapPoint` (2 ints) and `MapRect` (4 ints), and 5 of them a
  `Picture` string. `R1325` compares that string against the literal
  `nothing` to decide whether the object gets a global-map icon.
- Numbering: `MapObject<k>` is the 1-based section name and `k − 1` the
  0-based index `R1324` returns. Every object below is named by section.
- Corpus, all three roots identical (`EXP-0060/evidence/globalmap.csv`): 28
  `Mission<n>` keys, values 2–35, 0 of them 1.
- `MapObject1` carries `MapPoint 215 234` / `MapRect 184 189 110 48`, the
  largest rect in the file. It is the position `R1409` tests the party
  against (`L07501`…`L07502`) before building the offered-mission list, and
  the position `R0701` travels to at `L07503` when a mission ends with
  no `AutoGetMission`.
- `MapObject20`…`MapObject35` (indices 19…34) are the file's only `1×1` rects,
  sixteen of them, each named by exactly one `[MissionObjects]` key.
  `MapObject20`/`21`/`22` carry the main missions 20, 40 and 80; the thirteen
  side missions occupy `MapObject23`…`MapObject35` contiguously, one each.
- The other twelve main missions sit on the large, sometimes-pictured objects
  2…19; only `MapObject2` carries two (`Mission140` and `Mission150`).
- Four of the thirteen side entries, `Mission81`, `Mission101`, `Mission141`,
  `Mission151` → `MapObject27`/`29`/`32`/`33`, are `REG-SCN-059`'s dangling
  keys: present in `[MissionObjects]`, with no `[Mission<n>]` section in
  `Scenario.reg`.

**Confidence.** High for the section/key order and the default, which are the
push order in the listing, and for the object-0 identity, which has two
independent consumers read at instruction level. The object partition is a
complete census of the file's own 35 `MapObject` sections against its own 28
`[MissionObjects]` keys (`evidence/globalmap.csv`), reproducing in all three
roots. Medium that `MapObject1` is the city as a matter of art rather than of
code: no instruction names it. The code says it is object index 0, the only
object no mission is placed on, and the one the game returns to.

**Amended.** The sentence "Objects 20–35 are all `1×1` rects and every one of
them is the target of a *side* mission" was published at High; its second half
is retracted ([`retracted.md`](retracted.md), EXP-0060). It fails on
`MapObject20`/`21`/`22` under either numbering convention. The numbering rule
and the object partition above replace it.

### REG-SCN-067

- `[General] MercenaryCount` is not per mission. It is a 15-element pool
  indexed by mercenary type, 1..15, which is also the `[npc<t>]` section
  number.
- `R0757` reads it into a `CWordArray` at `record+0x70` and copies it
  element for element into a working array at `record+0x5c`
  (`L07504`…`L07505`), zeroing a parallel 15-`int` hire-flag array at
  `record+0x88` in the same loop.
- Every consumer subscripts all three by `byte [unit+0x15b] − 1` (`L07506`,
  `L07507`, `L07508`). This is why `REG-SCN-059`'s per-main-mission reading
  failed 12/15: the array was right and the index space was wrong.
- `[Mission<n>] Mercenaries` (`record+0x98`) is that mission's shelf of type
  ids. `[Mission<n>] EnableMercenary` (`record+0x30`) is a permanent unlock
  list, drained into `record+0xac` by vtable slot 0 of both record classes and
  applied on mission completion (`MERC-SHELF-002`).
- `[Mission<n>] Payment` (`record+0x0c`, `REG-SCN-059`'s destination) is
  returned by vtable slot 1 of both classes, `R1418` and
  `R1419`, each a bare return of the 32-bit field at offset 0xc of the object. `R1420` picks
  the finished mission's record and hands the value to `R1298` at
  `L07509`, on the mission-end arm, immediately before the `EnableMercenary`
  step.
- Corpus, all three roots identical (`EXP-0062/evidence/types.csv`,
  `offers.csv`, `checks.txt`): `MercenaryCount` = `[1 1 1 1 1 4 4 3 3 3 0 3 4 3 0]`.
  All 107 `Mercenaries` elements and all 13 `EnableMercenary` elements lie in
  1..15.
- `MercenaryCount[t] > 0` iff `[npc<t>]` carries both `PriceA` and `PriceB`:
  15/15. The two zero slots (types 11 and 15) are exactly the two `npc`
  sections without them, offered by no mission and unlocked by none.

**Confidence.** High for the index space: three named subscript computations
plus the reset's copy loop. The `MercenaryCount`↔`PriceA`/`PriceB` coincidence
is a discriminator, not corpus agreement: a per-mission array has no reason to
agree slot for slot with a different registry's presence set. Medium for
`Payment` as the reward: the getter, the picker and the call site are named
instructions, but `R1298` was read only as far as its call; nothing was
checked about sign or currency.

## `structures.reg`, picture fields and `projectiles.reg`

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-STR-080 | In `rom.exe`, the `structures.reg` class record is `0xa4` bytes filled field by field by `R1400`, and it has no `Parent` key, so a structure class inherits nothing. | High | ● active | [EXP-0092](../experiments/EXP-0092-structure-art/) |
| REG-STR-081 | In `rom.exe`, `structures.reg` `File` is a path string, not a `Files[]` index, and the two sheets it names are opened lazily on first draw; all 66 classes resolve. | High | ● active | [EXP-0092](../experiments/EXP-0092-structure-art/) |
| REG-STR-082 | In `rom.exe`, `AnimMask` is a picture of the sheet's frame grid, `FullHeight × TileWidth` cells, not of the footprint; `'-'` marks a cell that never animates. | High | ● active | [EXP-0092](../experiments/EXP-0092-structure-art/) |
| REG-PICT-083 | `units.reg InfoPicture` and `structures.reg Picture` are the only registry fields naming a picture; their join to `graphics\infowindow` fails on both roots, differently. | High / Medium | ● active | [EXP-0138](../experiments/EXP-0138-actor-pictures/) |
| REG-PROJ-086 | In `rom.exe`, `projectiles.reg` loads field by field into an array keyed by `ID`, so `ID` is an address, not a label: 31 shipped rows fill a 63-slot array. | High / Medium | ● active (amended, superseded) | [EXP-0139](../experiments/EXP-0139-spell-pictures/) |
| REG-PROJ-087 | In `rom.exe`, `projectiles.reg` `Palette` is a boolean, "this sheet carries its own colour table", not an index; it matches the trailer's bit 31 on 31 of 31 rows. | High | ● active | [EXP-0140](../experiments/EXP-0140-cast-art-drawn/) |

### REG-STR-080

- The size 0xa4 pushed at `L07510` sizes the record. `R1400` fills it in this
  order:
  - `+0x04`/`+0x08` the two `CSprite256` sheets (lazy, `REG-STR-081`);
  - `+0x0c ID`, `+0x10 TileWidth`, `+0x14 TileHeight`, `+0x18 FullHeight`,
    `+0x1c Phases`, `+0x20..+0x2c SelectionX1/Y1/X2/Y2`, `+0x30 ShadowY`;
  - `+0x34` the `AnimMask` string, `+0x38` its live-cell count, `+0x3c` the
    expanded animation timeline's length, `+0x40` the `CArray` holding it;
  - `+0x54 Picture[0x10]`, `+0x64 Indestructible`, `+0x68 DescText[0x20]`,
    `+0x88 VariableSize`, `+0x8c Usable`, `+0x90 Flat`, `+0x94 LightRadius`,
    `+0x98 LightPulse`;
  - `+0x9c` the `CString` `"graphics\structures\" + File`, `+0xa0` the
    sheets-loaded flag.
- Defaults are `-1` for `ID..SelectionY2` and `0` for `ShadowY` and the five
  flags.
- There is no `Parent` key in this loader at all, so unlike
  `objects.reg`/`units.reg` (`REG-KEY-045`) a structure class inherits nothing.
- The array is `SetAtGrow(ID)`-keyed (`L07408`), as `REG-KEY-044` states.
- `AnimTime`, `AnimFrame` and `AnimMask` are read only when `Phases > 1`
  (`L07511` compares the `Phases` value with 1 and branches when it is less or equal). The timeline is the same run-length expansion
  `REG-OBJ-046` publishes for objects (`AnimFrame[i]` appended `AnimTime[i]`
  times).

**Confidence.** High. Every field is the store of the read value into the class object at its field offset following its own
key `PUSH`, listed in `evidence/listings.md` §1. The absent `Parent` is an
absence over a routine read end to end, not a sweep.

### REG-STR-081

- Unlike `units.reg`/`objects.reg` (`REG-VAL-029`, `REG-ROSTER-052`), `File` is
  not an index into a `Files[]` table. `R1396` reads it into a `0x100`
  local at `L07512` and passes it to the class constructor, which stores
  `"graphics\structures\" + File` at `+0x9c`.
- The sheets are opened on first draw: every draw arm tests `class+0xa0` and
  calls `R1421`, which constructs `<path>.256` into `+0x04` and
  `<path>b.256` into `+0x08` and sets the flag.
- Corpus: all 66 classes resolve, under the whole-path lower-casing
  `RES-CASE-036` established; 12 of the 66 spell `File` with a capital that no
  archive node carries.

**Confidence.** High: two `PUSH`es of the suffix literals in one routine, the
flag's single writer in the same routine, and 66/66 resolving in
`graphics.res`.

### REG-STR-082

- The loader allocates exactly `FullHeight * TileWidth + 1` bytes for
  `AnimMask` (`L07513`…`L07514`) and counts the bytes that are not `'-'`
  into `class+0x38` (`L07515`/`L07516`).
- The draw reads `AnimMask[k·TileWidth + c]` for image cell `(k,c)`. `'-'`
  means this cell never animates: draw its base frame. Anything else means this
  cell has one frame per animation phase, at the rank it holds among the live
  cells (`SPR256-STR-041`).
- Corpus: 14 of the 66 classes spell a non-empty `AnimMask`, and its length is
  `TileWidth × FullHeight` on 14/14. The rival `TileWidth × TileHeight` differs
  on 6 of those 14 and fails all 6.
- The alphabet over all 14 masks is exactly two bytes: `'+'` (78 occurrences)
  and `'-'` (56).
- Ten further classes spell `Phases > 1` and then spell neither `AnimMask` nor
  `AnimTime`/`AnimFrame`, so `Phases > 1` alone does not mean animated: with an
  empty timeline the phase frame is 0 and the null mask pointer is never
  dereferenced.

**Confidence.** High. The buffer size is one `IMUL`/`INC` pair, the live count
one comparison with 0x2d, and the corpus discriminates the only live rival 14/14 vs 8/14.

### REG-PICT-083

- `units.reg InfoPicture` (`class+0xd8`, `REG-UNITS-049`) and
  `structures.reg Picture` (`class+0x54`, `L07517`/`L07518`) are `0x10`
  bytes each. The engine's leaf is `<name>.bmp` or `<name><tier>.bmp`, with the
  digit dropped at tier 1 (`UNIT-PICT-036`).
- The join takes every section of both registries against
  `graphics\infowindow` over tiers 1..4.
- EN: 34 unit classes, of which 33 carry `InfoPicture`; 66 structure classes,
  all 66 carrying `Picture`. 56 nodes are named by `units.reg` and 27 by
  `structures.reg`; `horse.bmp` is a hardcoded literal (`HERO-DOLL-078`);
  `root.bmp` and `teleport.bmp` are named by nothing. 56+27+1+2 = 86, the whole
  tree.
- Thirteen unit classes name a value with no node at any tier
  (`SwordsMan SwordsManSh 2HSwordsMan AxeMan AxeManSh 2HAxeMan ClubMan ClubManSh PikeMan PikemanSh Archer XBowMan`,
  twelve distinct names over thirteen sections). They are exactly the classes
  `UNIT-PICT-035` shows can never reach the formatter.
- RU: the unit half is identical cell for cell. The structure half is not:
  thirteen of the 66 sections name `magic`, `ruins` or `sphinx`, and
  `GRAPHICS.RES` ships no such node. So the RU root ships thirteen dangling
  picture references, and both roots ship nodes no class names.
- Cross-check over the sixteen classes whose id is `>= 0x1a`, the ones the
  portrait arm can reach: `units.reg Palette` equals the number of shipped
  picture files for that class's name, 16 of 16, at both of its values (`4` on
  thirteen classes, `1` on `Catapult1`, `Catap2` and `daemon`). This is
  `PAL-LIMIT-009`'s four-tier model predicting a file count in a tree it says
  nothing about.

**Confidence.** High for the counts and the two set comparisons: exact archive
and registry measurements on both roots, joined through the engine's own
formatter rather than by pattern. Medium for `Palette == files`: sixteen
agreements over two values is corpus agreement, carried as corroboration rather
than as the reason the tier is the tier.

### REG-PROJ-086

- `R0606` reads `[Global] Count` (31), then for each `i` in
  `0..Count-1` builds section `Projectile%d`, takes `File` as a string and
  `A16` as the constructor's second argument, and allocates `0x3c` bytes
  through `R1422`, which stores `A16` at `+0x34` and zeroes `+0x38`, the
  not-yet-loaded flag.
- Fields and defaults: `+0x14 = ID` (default -1), `+0x10 = Phases` (-1),
  `+0x18 = RotationPhases` (0x10), `+0x1c = Width` (0x40), `+0x20 = Height`
  (0x40), `+0x28 = Palette` (0), `+0x2c = Homing` (0), `+0x30 = Flip` (0),
  `+0x24 = SFX` (0).
- Every default is its own `PUSH` and is the live value wherever the key is
  absent. On the shipped file that is 8 rows for `Width`, `Height`,
  `RotationPhases` and `SFX`, 7 for `A16`, 3 for `Palette`, 19 for `Homing`
  and 29 for `Flip`.
- It grows the array to `ID + 1` when needed (`R1423(ID + 1, -1)`) and
  stores `[L02822][ID] = record`. The shipped `ID` domain 1..62 over 31
  rows makes a 63-slot array with 32 permanently null slots.
- Every consumer indexes by `ID`: `R0556`, `R0558`,
  `R0603`, `R0509` and `R0607` (`EnumRefs refto:L02822`:
  19 hits / 9 owners / 0 orphan). `R1423` is the grow and
  `R1368` the release. `R1366` is the teardown: it calls `vt+4`
  on every record, frees the array and releases both shade tables and both
  smoke sprites.
- Outside the loop and outside the registry the same routine loads
  `projectiles.pal` into the global at `L01248`, `projectile_.pal` into
  the global at `L07214`, and `graphics\projectiles\smoke%d\sprites.16a` for 0 and 1
  into the global at `L02869`.
- Corpus, both roots, parsed identically: `Count = 31` matches 31 blocks and 31
  distinct `ID`s. Two blocks name the same `File` (`catap2\sprites`, `ID` 6
  and 7). The tree holds 32 `projectiles/*` directories; the two the registry
  never names are `smoke0` and `smoke1`.
- The other consumer of the same id space is `units.reg`: 9 of the 34 classes
  carry a `Projectile` key, values 1 2 3 4 5 6 7 10 12. 10 and 12 are
  `MAGIC-PIC-026`'s `fire_arrow` and `fire_ball` pictures, so two unit classes
  shoot a spell's art.
- G2: `Count` and `ID` are both engine limits and both already elastic; the
  array grows to `ID + 1` with no clamp. The id a cast can reach is fixed at
  `2*spellId + 8` or `+9` (`MAGIC-PIC-026`), so slots above 65 are reachable
  only through a `units.reg` `Projectile`. Adding a row changes
  `projectiles.reg`'s bytes and nothing else.

**Confidence.** High for the fields: each is its own key string and `PUSH`ed
default in a decompiled loader whose every store target is a distinct
displacement, and the by-`ID` indexing is confirmed independently by two
consumers reading the global at `L02822` with the object's `+0x20`. Medium for the
per-key present-or-absent counts, which are a corpus census.

**Amended.** The claim named `R1366` as the grow ("`R1366` the
grow and `R1368` the release"). EXP-0140 superseded that label
([`retracted.md`](retracted.md)): `R1366` is the teardown, and the grow
is `R1423`. Every field, default, count and consumer is untouched.
EXP-0140 also settles what `Palette` selects (`REG-PROJ-087`).

### REG-PROJ-087

- `REG-PROJ-086` records the field at `+0x28` with default 0 and three shipped
  rows defaulted, and records `projectiles.pal`/`projectile_.pal` being loaded
  outside the loop.
- The lazy loader `R0604` gates on it:
  `L07519` loads the 32-bit field at offset 0x28 of the object, tests it and branches to `L07520` when it is zero. When
  `Palette` is zero the routine builds no lookup at all. When it is non-zero it
  calls `R0919(this, 0x10, mode, tint)` on the sprite it has just built,
  with `(4, 0)` on the `.16a` arm and `(2, 1)` on the `.256` arm, and again on
  the `b` sibling at `+0x08`.
- The draw reads the same field and picks the blit slot from it
  (`ANIM-CAST-027`).
- The corpus half is a two-sided check rather than a census:
  `tools/castart -mode arms` walks each sheet reading the trailer's bit 31
  instead of assuming an offset, and `Palette != 0` <-> embedded palette
  present holds 31/31 on both preserved roots and the live install.
- The three rows that default are exactly the three sheets shipped with bit 31
  clear, `archer\arrow`, `xbowman\arrow`, `orc\arrow`, whose frames therefore
  begin at offset 0, not 1024.
- Two consequences for a decoder: the walker in `tools/spellart` starts every
  `.16a` at 1024 unconditionally and would mis-frame any sheet in that state;
  and because the `b` sibling is drawn only on the `Palette != 0` arm,
  `archer\arrowb.256`, `xbowman\arrowb.256` and `orc\arrowb.256` are loaded and
  never drawn.
- The companion field `A16` at `+0x34` selects more than the extension.
  `R0604` builds a `.256` sprite (vtable `L04339`) for 0 and a
  `.16a` sprite (`L07521`) for non-0. Those two vtables differ in exactly
  two slots, the deleting destructor and `+0x18`, so the class chooses the blit
  and, through the branch above, the shade mode too.
- G2: `Palette` is a registry value and free, but not independent. Setting it
  to 1 on a sheet with no embedded table builds a lookup from 1 024 bytes of
  frame data; clearing it on a sheet that has one silently redirects that sheet
  to the shared `projectiles.pal`. Adding art that needs its own colours means
  shipping a palette-bearing sheet, which changes that file's bytes.

**Confidence.** High. The gate and both build calls are read as raw
instructions with their `PUSH`ed arguments. The 31/31 agreement is not
corpus-only, because each side is read from a different place: one from the
registry the loader parses, one from the four bytes the constructor tests.

## `npc.reg` faces, portraits and the archetype block

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-NPC-088 | `npc.reg` `Face` and `Picture` index no picture resource: they are the `typeID` and face of a dialogue actor synthesised for the npc, each store gated by a `Flags` token. | High / Medium | ● active | [EXP-0141](../experiments/EXP-0141-npc-fields-and-tag/) |
| REG-NPC-089 | `npc.reg` `PortraitX1/Y1/X2/Y2` is a Win32 `RECT`, and only its `left` and `top` are ever read. | High / Medium / Unknown | ● active | [EXP-0141](../experiments/EXP-0141-npc-fields-and-tag/) |
| REG-NPC-090 | `npc21`..`npc24` are a four-slot block `rom.exe` knows by literal number at three sites, and `DataBinID` has one located reader, `R1155`. | High / Medium / Unknown | ● active | [EXP-0141](../experiments/EXP-0141-npc-fields-and-tag/) |
| REG-NPC-091 | The `Portrait*` window is cut from the 160 x 240 `infowindow` bitmap itself, and it frames the creature's head only if the canvas keeps the BMP's bottom-up row order. | High / Medium | ● active | [EXP-0141](../experiments/EXP-0141-npc-fields-and-tag/) |

### REG-NPC-088

Sources: `rom.exe` and the shipped corpus.

- `R0743(npcId)` allocates `0x1b0` bytes, constructs at `R0593`, and
  folds the section's `Flags` tokens into `object+0x18c`: `Hero`→bit 0,
  `Mage`→bit 1, `Female`→bit 2, `Human`→bit 4. `MySex`/`MyClass`/`!MySex`/`!MyClass`
  copy bits 1 and 2 from the player's own drawable at `[this+0x3f54]+0x18c`.
- It then takes the record `((void**)[L03415])[npcId]` and performs exactly
  two field copies:
  - `L07522` loads the 32-bit field at offset 0x10 of the record and `L07523` stores it at offset 0x20 of the object under the
    `Flags` token `Picture` (`L07524`, tested by `R1286`);
  - `L07525` loads the 32-bit field at offset 0xc of the record and `L03494` stores the record pointer at offset 0x24 of the object when the
    token `Start` is absent.
- `object+0x20` is a drawable's typeID (`UNIT-APPEAR-030`, `ALM-CLS-054`) and
  `object+0x24` its face byte (`PAL-FACE-005`). The portrait that follows is
  `UNIT-PICT-036`'s existing formatter with no new mechanism: `R0740`
  reads the same two fields back (`L07526`, `L07527`), subscripts the
  `units.reg` class-by-`ID` array `[L02113 + 4*typeID]` (`REG-KEY-044`),
  takes `class+0xd8` = `InfoPicture` (`REG-UNITS-049`) and formats
  `graphics\infowindow\%s%d`, dropping the digit at `face == 1`.
- With `Start` present the face comes instead from one of the four archetype
  sections, chosen by a jump table at `L07528` on `object+0x18c & 6`:
  `0`→`MaleFighter`, `2`→`MaleMage`, `4`→`FemaleFighter`, `6`→`FemaleMage`.
  Each is read as `Face` from `scenario\npc.reg` reopened at `L07529`; the
  loader caches the same four at `L07530`/`74`/`78`/`7c`.
- The `Flags` list is the record's schema. Over 105 sections on both roots the
  token and the key are biconditional: `Picture` token ⟺ `Picture` key (33/33
  and 72/72), `Face` token ⟺ `Face` key (86/86 and 19/19).
- The domain splits along `R0740`'s own branch `L03434`, a mask of the flags with 0x11
  (`Hero` or `Human`). Where `Picture` is present, `Face` is 1..4, a tier, and
  the section carries `!Human`. Where it is absent, `Face` is 1..30, the
  section carries `Human` or `Hero`, and it takes the composed-doll arm
  (`UNIT-PICT-035`) instead.
- Corpus join, both roots identically: all 33/33 sections carrying `Picture`
  resolve as a `units.reg` `ID`, that class carries `InfoPicture`, and
  `graphics.res` ships the exact leaf the formatter builds. The sections fall
  in (class, tier) order: `npc100..103` Goblin 1..4, `npc104..107` Orc 1..4,
  through `npc132` daemon 1 (`evidence/picture-join-*.csv`).
- G2: both fields are registry values and free to change, but neither is
  independent. `Picture` must name a live `units.reg` `ID` whose class carries
  an `InfoPicture`, and a `Face` above the tier count for that name resolves to
  a node that does not ship. Adding a portrait means adding a `graphics.res`
  node, which changes that file's bytes; re-pointing an existing npc changes
  only `npc.reg`.

**Confidence.** High. Both stores are named instructions with their
displacements, the `Flags` gates are the `PUSH`ed key strings of the same
routine, and the destination fields are identified by two prior claims that
read them from unrelated paths. Medium for the 33/33 corpus join, which is
corroboration and not the reason for the grade.

### REG-NPC-089

Sources: `rom.exe` and the shipped corpus.

- The loader `R0499` reads the four keys in the order `PortraitX1`,
  `PortraitY1`, `PortraitX2`, `PortraitY2`, each with default -1, into four
  consecutive stack dwords.
- The record constructor `R1424` copies them whole with
  `L07531` forming the address 0x20 bytes into the record, `L07532` and `L07533` pushing the two rectangle pointers, and
  `L07534` calling through the dword at `[L01508]`: the import `CopyRect` from
  `USER32.dll`, whose parameters are `LPRECT` and `const RECT*`.
- The mapping is fixed by a type, not by the key names: `record+0x20` = `left`
  = `PortraitX1`, `+0x24` = `top` = `PortraitY1`, `+0x28` = `right` =
  `PortraitX2`, `+0x2c` = `bottom` = `PortraitY2`.
- The rectangle's only reader is `R0740`. It loads all four
  (`L07535`, `L07536`, `L07537`, `L07538` into the frame locals at -0x30, -0x2c,
  -0x28, -0x24), then uses `left` four times and `top` five; `right` and
  `bottom` are never read again. This is a complete enumeration of four
  adjacent locals inside one function body.
- `L03605` compares the stack local at frame offset -0x2c with -0x1, testing `top` against the loader's own
  key-absent default and picks between two six-argument draw calls:
  `(8, 7, 0x24, 0x8c, 0x6c, 0xe8)` when the key is absent and
  `(8, 7, X1, 0x90-Y1, X1+0x48, 0xf0-Y1)` when it is present. Both are made on
  the global `[L07539]` through vtable slot `+0x34` and on the caller's own
  canvas through `+0x38`.
- Both slots (`R1425`, `R1426`) append `this->buf+8`, `buf[0]`
  and `buf[1]` and call `R0797` / `R0798`, which clip argument 1
  against `[L01503]` and push the shortfall into argument 3. So args 1–2 are
  a destination point and 3–6 a source rectangle.
- The canvas is `R0744(0xa0, 0xf0)`, `w*h*2 + 8` bytes with `w` at
  `[0]` and `h` at `[4]`: 160 x 240, 16-bit.
- The window is therefore a fixed 72 x 96 whose blit-space `top` is
  `0x90 - Y1`, against a fixed 72 x 92 default at `(36,140)`. Those are canvas
  rows, not picture rows: `REG-NPC-091` shows the canvas is the 160 x 240
  `infowindow` BMP kept in its own bottom-up order, so `PortraitY1` is an
  ordinary offset from the top of the picture and the window's visual top is
  `Y1 - 1`.
- Corpus, both roots identically: `PortraitX1` and `PortraitY1` are present on
  48 sections, `PortraitX2` and `PortraitY2` on 10 (`npc31..35`, `npc41..44`,
  `npc54`). The derived source rectangle lies inside the 160 x 240 canvas on
  48/48, out of a 32-bit domain (`evidence/portrait-geometry-*.csv`).
- G2: `X1`/`Y1` are free within `0 <= X1 <= 88` and `0 <= Y1 <= 144`. The
  72 x 96 window size, the 160 x 240 canvas and the `(8,7)` destination are
  engine constants, and moving any of them is a code change. `X2`/`Y2` are
  inert: editing them changes nothing this image does.

**Confidence.** High for the `RECT` identification (an import's parameter
type), the field order (the loader's own read order into consecutive slots),
the dead pair (an exhaustive read of four locals in the single reader) and the
source/destination split (the clip arithmetic of the leaf). Medium that
`R0740` is the only reader. The instrument is `EnumRefs refto:L03415`
(20 hits / 7 owners / 0 orphan) plus `imm:` and `refto:L02112` on the array
object (10 / 7 / 1 orphan, disassembled and read). Its blind spot: a wholesale
`memcpy` of a record carries no displacement, and a route computing the array
address rather than naming it is invisible to both modes.

**Unknown.** Whether `X2`/`Y2` were ever live in a shipped build: absence of a
reader in this image is established, intent is not.

### REG-NPC-090

Sources: `rom.exe` and the shipped corpus.

- Three independent sites:
  - `R0711` gates one of its two npc calls on
    `L07540` comparing the 32-bit local at frame offset -0x24 with 0x15 and branching when it is less, and
    `L07541` comparing it with 0x18 and branching when it is greater;
  - `R1155` carries the identical pair at `L07542`/`L07543` on its
    own argument;
  - `R0908` reaches the same four as
    `L07544` loading the dword at offset 0x54 plus 4 times the index in the array at
    `[L03415]`, i.e. `array[id + 21]`.
- The corpus agrees from the other side. The four sections carrying the
  `Start` token, the only ones whose face is taken from an archetype rather
  than from their own `Face` key (`REG-NPC-088`), are exactly `npc21`
  `Hero,Me,Start`, `npc22` `Hero,Mage,!MySex,Start`, `npc23`
  `Hero,!Mage,!MySex,Start`, `npc24` `Hero,!MyClass,MySex,Start`, on both
  roots. All four carry `DataBinID` 26 and no `Face`, `Picture` or `Portrait*`.
- `R0908` also reads `record+0xc` (`L07545`) under the `Flags` token
  `Face`, and compares the four cached archetype values as a table:
  `L07546` loading the dword at the table `L07530` plus 4 times the index, the only read of that
  table.
- `DataBinID` (`record+0x14`): its one located reader is `R1155` at
  `L07547` loading the dword at offset 0x14 of the record and
  `L07548` storing it at offset 0xa of the message body. It writes the value into a message
  body whose byte tag at `+0x9` is set to `0x49` two instructions later. This
  is consistent with `REG-NPC-058`'s reading of the field as a `Templates.ini`
  placeable id, and it is the first instruction in the image that consumes it.
- G2: the `21..24` range is an engine constant in three places. An npc outside
  it cannot take the archetype-face path or whatever the gated call at
  `L03420` provides; moving the block is a code change, not a data change.

**Confidence.** High for the three gate sites, the `+0x54` subscript and the
`DataBinID` store (named instructions with their immediates and
displacements), and for the corpus census of `Start` (exhaustive over 105
sections on both roots). Medium that `R1155` is `DataBinID`'s only
reader: same instrument and same blind spot as `REG-NPC-089`.

**Unknown.** What the message tagged `0x49` is: the tag byte and the field's
landing offset are read, the message's own grammar is not.

### REG-NPC-091

Sources: the shipped corpus and renders of it.

- Every one of the 86 `graphics.res::infowindow/*` nodes, on both roots, is a
  `BM` file declaring 160 x 240, 24 bpp, compression 0 with a positive height,
  i.e. a bottom-up DIB. That is the exact shape of the canvas
  `R0744(0xa0, 0xf0)` allocates (`REG-NPC-089`).
- So the blit's source rectangle `(X1, 0x90-Y1)..(X1+0x48, 0xf0-Y1)` names a
  region of that bitmap, and exactly two readings survive the instructions:
  - BOTTOM-UP: the canvas holds the file's rows unflipped, canvas row `r` is
    visual row `239-r`, and the window's visual top is `Y1-1`;
  - TOP-DOWN: the canvas is flipped at load and the visual top is `0x90-Y1`
    directly.
- The two coincide only where `Y1 = 72` and are up to 122 rows apart
  elsewhere.
- `tools/npcportrait` renders both for all 33 sections that carry `Picture`
  and `PortraitX1`/`PortraitY1`. Under BOTTOM-UP the 72 x 96 window frames the
  head and shoulders on 33/33; under TOP-DOWN it lands on legs, wings or empty
  background. For `npc132` (`daemon`, `Y1 = 11`) the TOP-DOWN window runs to
  visual row 229 and holds almost no figure. Both PNGs are byte-identical
  between the two preserved roots.
- `PortraitY1` is therefore an ordinary top-down offset from the top of the
  picture, and the engine's `0x90 - Y1` / `0xf0 - Y1` is the bottom-up
  conversion of a 96-tall window (`240 - (Y1 + 96)` and `240 - Y1`). This also
  fixes the default rectangle `(0x24, 0x8c, 0x6c, 0xe8)` as visual rows
  8..100, the top of the frame, not the bottom.
- G2: an author retargeting a portrait sets `PortraitX1`/`PortraitY1` in
  picture coordinates with the top-left corner as origin, bounded by
  `0 <= X1 <= 88` and `0 <= Y1 <= 144`; nothing else in the path is a data
  value.

**Confidence.** High for the bitmap census (an exhaustive header read of all 86
nodes on both roots) and for the two readings being the only survivors (the
blit's rect is fixed by `REG-NPC-089`, and a DIB has exactly two row orders).
Medium for BOTTOM-UP being the live one: 33/33 is a corpus agreement judged by
eye, and the instruction that would settle it, whether the loader that fills
the canvas from the BMP walks the rows forwards or backwards, was not read. The
falsifier is cheap: read that loader, or observe one dialogue portrait of
`npc132` on a running original. Under BOTTOM-UP it shows a horned head, under
TOP-DOWN a nearly empty pane.

## `DescText`, mission documents and `AddHero`

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-DESC-096 | `units.reg` `DescText` is identical on both roots and pure ASCII, and its 31-byte shipped maximum fills the `0x20` field it is read into: a limit as well as dead data. | High / Medium | ● active | [EXP-0143](../experiments/EXP-0143-actor-names/) |
| REG-SCN-097 | `AddTextDocument` and `AddPictureDocument` are one append onto one campaign-lifetime collection, and the two readers differ in three bytes. | High / Medium | ● active | [EXP-0151](../experiments/EXP-0151-mission-documents/) |
| REG-SCN-098 | `[Mission30] AddHero = 22` is consumed on the first town activation after mission 30 becomes current. | High | ● active | [EXP-0153](../experiments/EXP-0153-addhero-consumer/) |

### REG-DESC-096

- Measured over both installs (`evidence/desctext-roots.txt`): 34 of 34 classes
  carry the key, the value is byte-identical `en` to `ru` on 34 of 34, and no
  value carries a byte >= 0x80 on either root.
- The longest value plus its NUL is 31 bytes, against the `0x20` the loader
  reads into at `class+0x10c` (`REG-UNITS-049`; `0x10c + 0x20 = 0x12c` = the
  allocation): one byte of headroom on the longest shipped class.
- Two models fit those bytes, and the corpus alone cannot separate them:
  display text the Russian translator never touched, or metadata for the editor
  nothing draws. Two independent facts separate them, and the second model
  survives: nothing in the image reads the field, and the name a class is
  shown by is a line of `main\text\unitname.txt` whose text disagrees with
  `DescText` on several classes, `Death Star` against `Daemon` and `Ghost`
  against `Spirit` (`UNIT-NAME-039`, `UNIT-NAMETAB-041`).
- G2 limit: a `DescText` may be at most 31 characters. The class record is a
  fixed `0x12c` allocation with the string inline at the end, so a longer value
  writes past the object rather than being truncated by a `CString`. Lifting
  the limit changes the allocation size and every offset after it: a code
  change, not a file change.

**Confidence.** High for the corpus (both roots, all 34 classes, whole values)
and for the 31-vs-`0x20` arithmetic (the loader's own push of 0x12c and its
`0x20`-byte read). Medium for "a longer value overruns rather than truncates":
the loader's string-read helper was not read, only its byte count, so a clamp
inside that helper would change the failure mode without changing the limit.

### REG-SCN-097

- Corpus, both roots identical, over the whole registry:
  `[Mission10] AddTextDocument` kind 6 = `{1,2,3}`,
  `[Mission50] AddTextDocument` kind 2 = `{4}`,
  `[Mission60] AddPictureDocum` kind 2 = `{1}`. That is 4 text elements over 2
  sections and 1 picture element over 1 section, of 24 `[Mission<n>]`
  sections. `[General]` is exactly
  `{MercenaryCount, ScenarioMissionCount, TotalMissions}`.
- `R1297` reads each key as a `u16` array and calls one routine per
  element.
- `R1403` and `R1404` are the same growable-array append over a
  60-byte element on the same four fields of the same object: `+0x148` data,
  `+0x14c` size, `+0x150` capacity, `+0x154` grow-by, growth
  `min(max(size/8,4),1024)` at `L07549`…`L07550`. Neither clears the array.
- They differ only in the reject polarity (`L07551` `JNZ` on a non-zero
  `+0x04`, `L07552` `JZ` on a zero one), the kind literal constructed with
  (`L07553` pushes the constant 1, `L07554` pushes the constant 0) and the exception frame.
- Element `+0x00` is therefore the registry value and `+0x04` the kind, 1 text
  and 0 picture, with the append deduplicated on the pair.
- The owner is the campaign record built by `R0677` (six such arrays;
  this one's vptr at `+0x144`). `R0755`, the loader's only caller, runs
  it only for a mission strictly higher than the current one and returns 0 for
  a lower one, so the collection only grows.

**Confidence.** High for the corpus: exhaustive, both roots, every value
resolving to a shipped resource 5/5. High for the append and the kind
assignment: both routines read whole at instruction level, and the kind is
discriminated by each routine's reject polarity agreeing with its own literal,
so a swapped reading contradicts itself. Medium for "only grows": the monotonic
guard is one compare read in a decompiled caller, and the campaign's own
mission sequence was not driven.

### REG-SCN-098

- Town activation calls `R1384`, which walks `AddHero[]`, calls
  `R1155` per element and clears the array.
- The special `npc21..24` branch reads `npc22` tokens `Mage,!MySex`. An
  existing female character selects Humans row 28 `PC_Fergard`; an existing
  male selects row 29 `PC_Reniesta`.
- The stored `DataBinID = 26` is replaced by `26 + selector`; the name comes
  from localized `text/npcnames.txt` at `20 + selector`.
- Mission 20 makes mission 30 current before town activation, so the companion
  exists before mission 30 starts and cannot be added twice from the cleared
  array.

**Confidence.** High. The caller, array walk, clear, selector arithmetic, two
table lookups and shipped registry values are direct evidence; completion-time
and direct-`DataBinID` alternatives are excluded by the instruction path.

## Shared container field operations

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-104 | Shared registry consumers apply different kind masks: descent requires `(kind & 0x40000001) == 1`, and the endpoint rejects `kind & 0x40000010 != 0`. | High / Unknown | ✔ promoted | [EXP-0348](../experiments/EXP-0348-container-field-operations/EXP-0348.md) |
| REG-105 | Kind 8 has an original string-array producer: setter L07399 writes a pool dword count followed by NUL-terminated byte strings, which getter L07400 reads back. | High / Unknown | ✔ promoted | [EXP-0348](../experiments/EXP-0348-container-field-operations/EXP-0348.md) |
| REG-106 | Integer-array getters copy 1, 2 or 4 bytes from each kind-6 pool dword for unsigned `size >> 2` elements, and on kind 2 read only the low scalar byte. | High / Unknown | ✔ promoted | [EXP-0348](../experiments/EXP-0348-container-field-operations/EXP-0348.md) |
| REG-107 | Pooled typed setters keep a record's size when a value shrinks: an equal or smaller write adds the unused bytes to registry +0x30 without reducing record +8. | High / Unknown | ✔ promoted | [EXP-0348](../experiments/EXP-0348-container-field-operations/EXP-0348.md) |
| REG-108 | Converter R1396 dispatches on `(kind >> 1) & 7`: type 0 copies, type 2 formats a signed decimal, type 8 copies pool bytes from offset four, and types 12/14 throw. | High / Unknown | ✔ promoted | [EXP-0348](../experiments/EXP-0348-container-field-operations/EXP-0348.md) |
| REG-109 | Binary64-to-string conversion ends in a later integer conversion: after formatter R1196 returns, L07555 formats the low dword in radix 10 into the same buffer. | High / Medium | ✔ promoted | [EXP-0348](../experiments/EXP-0348-container-field-operations/EXP-0348.md) |
| REG-110 | The direct double setter and getter move both inline dwords at +4/+8, and the double-array pair moves eight bytes per item for unsigned `size >> 3` items. | High / Unknown | ✔ promoted | [EXP-0348](../experiments/EXP-0348-container-field-operations/EXP-0348.md) |
| REG-111 | Record reserve L07556 grows to `capacity+(capacity>>2)+1` 32-byte slots; pool reserve L07557 requests `capacity+(capacity>>3)+request` bytes outside compaction. | High / Unknown | ✔ promoted | [EXP-0348](../experiments/EXP-0348-container-field-operations/EXP-0348.md) |

### REG-104

- Descent R1427 requires `(kind & 0x40000001) == 1` before another
  component. Endpoint R1428 rejects `kind & 0x40000010 != 0`.
- Deletion R1429 marks bit 30, recurses for active subkeys, and accounts
  active non-inline leaf size at registry +0x30; already marked nodes do not
  repeat it. Global L07558 equal to 1 stops after the first child.
- Copy L07559 includes active leaves/subkeys under mask results 0/1, omits
  marked children, clears copied parent bit 4, and otherwise retains leaf
  flags.
- These are consumer rules, not physical removal or a closed enum.
- Evidence: V001..V132, V384.

**Confidence.** High for complete bodies and bounded one-bit controls.

**Unknown.** Malformed/cyclic trees.

### REG-105

- Setter L07399 writes a pool dword count followed by NUL-terminated byte
  strings. The source array uses pointer +4/count +8 and string lengths at each
  data pointer minus eight; a new item needs `4 + sum(length+1)` bytes.
- Getter L07400 sizes output from the count, then scans according to record
  byte size. Count 1 with two strings in an eight-byte item reaches a second
  destination-element boundary; count 3 with the same item makes two
  assignments.
- This retracts REG-KIND-034's producer-absence/exhaustive-enumeration clauses
  ([`retracted.md`](retracted.md)).
- Evidence: V342..V346, V366..V373.

**Confidence.** High for complete original bodies and conditional
setter/getter/raw-readback controls.

**Unknown.** Native caller reach and native string construction.

### REG-106

- L07560/R1430/R1061 and R1431 copy 1/2/4 bytes from each kind-6
  pool dword, for unsigned `size >> 2` elements.
- On kind 2, all read only the low scalar byte, and wider destinations
  zero-extend it. Scalar 0x12345681 yields 81/0081/00000081, while R0452
  returns the full dword.
- Four-byte getters also accept type 0 with size below 2 as empty; byte/word
  controls throw. Partial final element bytes are ignored.
- R1431 is an array getter with inline resize logic, correcting
  RES-CODE-020 ([`retracted.md`](retracted.md)).
- Evidence: V267..V298, V381.

**Confidence.** High for complete bodies and width/length controls.

**Unknown.** Native allocation and application consequences.

### REG-107

- String, integer-array, string-array and double-array setters update size and
  request a pool segment on growth; equal/shrink adds unused bytes to registry
  +0x30 without reducing record +8.
- The 32-byte replacement controls retain that extent and the untouched pool
  suffix through raw writer/readback.
- Matching existing records retain kind bits 28/30; insertion assigns the whole
  supported kind and clears the insertion long-name flag.
- Byte/word array setters zero-extend into four-byte elements; dword/double
  setters copy four/eight bytes.
- Evidence: V307..V351, V383.

**Confidence.** High for complete bodies and 45 setter/writer/loader controls.

**Unknown.** Native compaction/failure and malformed metadata.

### REG-108

- Type 0 copies `min(size,capacity)`. Type 2 formats the signed dword in base
  10 before bounded copying; its final destination NUL scan computes the return
  even at zero capacity.
- Type 8 computes unsigned `min(size-4,capacity-2)`, copies from pool offset
  four and writes NULs at capacity minus two/minus one, preserving any gap
  after a shorter copy. Size below four underflows; small capacity or truncated
  mapped sources can reach logical memory boundaries.
- Types 12/14 reach a throw, not a returned diagnostic.
- Evidence: V209..V265, V380.

**Confidence.** High for complete converter/copy/integer-format bodies and
guarded controls.

**Unknown.** Native error handling, arbitrary memory and double formatting.

### REG-109

- L07561 calls formatter R1196 with both inline dwords. The normal return
  at L07562 loads the low dword, and L07555 calls L07563 with radix 10
  into the same buffer before copying.
- With supplied formatter text FORMAT-CUT and low dword 0x12345678, the
  remaining original code returns 305419896.
- REG-DBL-035 inline binary64 storage stands.
- Evidence: V266; L07564..L07555.

**Confidence.** High for the instructions and the supplied normal-return
control. Medium for native end-to-end conversion, because formatter
success/output/failure is cut.

### REG-110

- Direct double setter L07565 writes both inline dwords at +4/+8; getter
  R1397 checks masked type 4 and loads the qword.
- Double-array setter L07566 emits eight bytes per item; getter L07567
  copies unsigned `size >> 3` qwords. Sizes 0/1/7 select zero items, 8/9 one
  and 23 two; partial final bytes are ignored.
- Tested wrong types reach the throw cut.
- Evidence: V299..V306, V317..V321, V347..V365, V382.

**Confidence.** High for complete instructions and width/extent controls.

**Unknown.** Exceptional floating values and native caller reach.

### REG-111

- Record reserve L07556 grows when capacity is at most count+1, to
  `capacity+(capacity>>2)+1`, requesting 32 bytes per slot.
- Pool reserve L07557 grows when capacity is at most used+request. Outside
  its compaction arm it requests `capacity+(capacity>>3)+request` bytes and
  advances used by request. With used 4/request 4, capacities 7/8/9 become
  11/13/9.
- The waste-sensitive arm calls L07568 before rechecking.
- Evidence: V374..V379.

**Confidence.** High for complete arithmetic and boundary controls under
synthetic allocation.

**Unknown.** Overflow, failure and native compaction.

## Inn NPC text closure

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-INNCLOSE-116 | `text/inn/npc/*.txt` matches `InnNPC`/`InnMission` addressing exactly on EN only; RU ships seven unaddressed files, and a fourth root lacks 20 addressed ones. | High | ● active | [EXP-0393](../experiments/EXP-0393-inn-text-closure/EXP-0393.md) |

### REG-INNCLOSE-116

- The addressed set is built from `REG-SCN-064`'s decoded reader
  (`npc<InnNPC[i]>m<stage>.txt`, zero arm on `InnMission[i]==0`) and
  enumerated over the shipped `npc/` tree. The addressing is `scenario.reg`'s.
- EN: 22 addressed combos and 22 shipped files match exactly in both
  directions (0 addressed-not-shipped, 0 shipped-not-addressed).
- RU: the `InnNPC`/`InnMission` arrays are the same 22 as EN's, reproducing
  `REG-SCN-064`'s "registries identical" from an independent walk, so RU also
  addresses 22. RU's `npc/` tree carries 29 files: the same 22 plus seven the
  registry never names
  (`npc23m70.txt npc24m80.txt npc25m111.txt npc29m150.txt npc29m151.txt npc43m140.txt npc44m50.txt`).
- These seven are the file half of `TEXT-ROOT-014`'s published RU-only count of
  seven `text/inn/npc/npc*.txt`, named here with a mechanism: they are
  orphaned, unreachable through the inn's own addressing, not merely present.
- A fourth, owner-supplied pre-release data root, outside
  `REG-SCN-063`/`REG-SCN-064`'s own live/EN/RU triple, inverts the shape: its
  registry addresses 24 combos of which only 4 ship. The 20
  addressed-not-shipped names are files that root's own `scenario.reg` expects
  but does not carry.
- Registry key vocabulary does not differ between the three roots measured:
  all 14 keys these roots' `[Mission<n>]` sections use anywhere
  (`Mercenaries`, `InnNPC`, `InnMission`, `ShopMinPrice`, `ShopMaxPrice`,
  `Payment`, `EnableMercenary`, `ShopMission`, `TCMission`, `AddTextDocument`,
  `AddHero`, `AddPictureDocum`, `AutoGetMission`, `LastMission`) are present
  on all three, 0 root-exclusive.
- Extends `REG-SCN-064` and `TEXT-ROOT-014`.

**Confidence.** High for the closure counts: two closed, named populations (the
shipped node list and the registry-addressed set) compared in both directions,
and which direction each root's mismatch runs is a discriminator a uniform
"always closed" or "always leaky" alternative would not predict. High for the
vocabulary census: exhaustive over every `[Mission<n>]` section on all three
roots; a root-exclusive key would read as a "no" in one column of
`evidence/vocab.csv`, and none does.

## Inn selection reader and fourth-root stage keys

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| REG-118 | `R0702` (release) and `R1432` (demo) run only when `curSel` reaches the gate count, read `InnMission` and `InnNPC` as `u16` at `idx = curSel − count`, and never check `InnMission`'s size. | High / Unknown | ● active (partially retracted) | [EXP-0394](../experiments/EXP-0394-inn-selection-loop/) |
| REG-INNSTAGE-117 | Where the fourth root's per-stage registry disagrees with EN/RU, the keys that move differ by stage, one stage differs on a single key, and one stage breaks `REG-SCN-064`'s equal-length invariant. | High / Unknown | ● active | [EXP-0393](../experiments/EXP-0393-inn-text-closure/) |

### REG-118

- Neither function is a per-portrait reader. Its callers enter it only when
  `curSel >= n`, where `n` is the gate count: release the field at offset 0xc8 of the view, demo the
  `CArray::GetSize()` accessor `R1433` on `record+0xc0`. Inside it the read
  subscript is `idx = curSel − n`, applied to both `InnMission` and `InnNPC`,
  read as `u16` words.
- Release, `R0702` (`evidence/d-inn-release.txt`), the exact range
  `REG-SCN-064` names, with the subscript decoded:
  - `L07490` reads the 32-bit field at offset 0xb8 of the record: `curSel`;
  - `L07569` reads the 32-bit field at offset 0xc8 of the record: `n`;
  - `L07570` reads the 32-bit field at offset 0x620 of the campaign screen: `InnMission.data`,
    `campaignScreen+0x61c`'s CArray data pointer;
  - `L07571` subtracts `n` from `curSel`: `idx`;
  - `L07572` reads the 32-bit field at offset 0x60c of the campaign screen: `InnNPC.data`, `campaignScreen+0x608`'s
    CArray data pointer, read as `u16` words;
  - `L07573` doubles `idx`: word index to byte offset;
  - `L07574` reads the 16-bit element at the data pointer plus the byte offset: `InnMission[idx]`;
  - `L07491` reads the 16-bit element at the data pointer plus the byte offset: `InnNPC[idx]`.
- Demo, `R1432` (`evidence/d-inn-demo.txt`,
  `evidence/d-accessors-demo.txt`): three independent
  `idx = curSel(record+0xb8) − InnNPC.Size()` computations
  (`L07575`, which subtracts the size from `curSel`, and others), with `R1433` the
  `CArray::GetSize()` accessor called on `record+0xc0`. Each feeds a one-line
  accessor: `R1434` (`record+0xd4` base, `InnMission[idx]`) or `R1435`
  (`record+0xc0` base, `InnNPC[idx]`). Both resolve through
  `R1436`, which forms the element address as the data pointer plus 2 times the index, the same word-array element arithmetic as
  the release build's inline form.
- All four callers, located and disassembled (`evidence/callers-release.txt`,
  `evidence/callers-demo.txt`, `evidence/d-dispatch-release.txt`,
  `evidence/d-dispatch-demo.txt`):
  - release `R0703`: `L07576` comparing `curSel` with the size and `L07577` branching to `L07578` when it is less
    skips the call; falling through calls `R0702` at `L07579`;
  - release `R1437`: `L07580` comparing `curSel` with the size and `L07581` branching to `L07582` when it is greater or equal →
    the call to `R0702` at `L07582`;
  - demo `R1438`: `L07583` comparing the field at offset 0xb8 of the record (`curSel`) with the size and `L07584` branching to `L07585` when it is greater or equal
    → the call to `R1432` at `L07586`;
  - demo `R1439`: `L07587` comparing the field at offset 0xb8 of the record (`curSel`) with the size and `L07588` branching to `L07589` when it is less
    skips; falling through calls `R1432` at `L07590`.
- All four gate identically: `curSel < n` never reaches the reader. That range
  goes to a sibling path (release `R1440`/`R1441`; demo the same
  call sites' other arms), which reads `record+0xc4` as an array of 4-byte
  pointers indexed directly by `curSel`, not by `idx`.
- In release that element is dereferenced at `+0x15b`
  (`L07591` loading the byte at offset 0x15b of the element, `MERC-TYPE-001`'s `CUnit` type-id offset) and
  pushed with the literal `inn\mercenary\npc%02d` (`L07592`) to
  `R1442`/`R1443`, `MERC-HIRE-003`'s hire/dismiss pair, called
  from these same two functions. `R0702`'s own literal is
  `inn\NPC\npc%02dm%d` (`L07593`, `REG-SCN-064`). The sibling arm is the
  mercenary roster those two claims describe, not a second `InnNPC`.
- Demo's sibling arm reads the same `+0x15b` type-id byte through a different
  accessor (`R1444`) and a different downstream call (`L07594`) that
  `MERC-HIRE-003` does not name. Only the release identification holds at
  caller-address precision; the demo match is offset-only.
- No instruction in any of the eighteen disassembled functions compares `idx`
  or `curSel` with `InnMission`'s own size field.
- Measured independently over all four permitted roots (`tools/innselbound`,
  `evidence/arraylen.csv`): the stage×root pair `REG-INNSTAGE-117` names (demo
  `Mission40`, `InnNPC=[22 90 84]` length 3, `InnMission=[40 41]` length 2) is
  the sole length mismatch of 39 pairs measured. The demo's `rom.exe` differs by
  SHA-256 from the binary EN, RU and the live install share
  (`evidence/input-manifest.json`).
- Registry file-pool byte adjacency is not evidence for any of this. The loader
  (`R1445`, `evidence/d-accessors-demo.txt`) allocates each screen-side
  array fresh at its own record's declared element count, which decouples file
  layout from what a live read can reach.
- Extends `REG-SCN-064` (amended) and `REG-INNSTAGE-117`; reconciled against
  `MERC-HIRE-003` and `MERC-TYPE-001`.

**Confidence.** High for the subscript formula, the four callers' gating and
the absence of any `InnMission` size check. All are named instructions,
cross-checked in two independently compiled binaries whose formula and gate
shape agree (the size subtracted from `curSel` in release against `curSel` subtracted from the size in demo; a less-than skip and
a greater-or-equal skip in different orders, same net gate). The sibling arm's release
identity as the mercenary roster rests on an exact caller-address and
struct-offset match to `MERC-HIRE-003`/`MERC-TYPE-001`.

**Unknown.** No traced instruction, in either build or any of the four callers,
places an upper bound on `curSel` itself before it reaches this gate. Demo
`R1439` selects its three-way region dispatch, including the arm that
reaches `R1432`, through an untraced hit-test call,
the call target `R1446` at `L07595`, seen and named, not disassembled; no
release-build analogue of that hit test was located. Whether real mouse or
keyboard input can drive `curSel` to a value producing `idx >= InnMission.Size()`
at demo `Mission40` (for example `curSel=5`, one word past `InnMission`'s
2-element allocation) is not established either way by EXP-0394.

**Amended.** [`retracted.md`](retracted.md) (EXP-0400, `TOWN-468`) withdraws,
for the release build, the identity of the bound and of `record+0xc4`, and the
Unknown about one field serving both readings. Former wording:
"the load at `L07569` of the count field at offset 0xc8 of the record (`InnNPC.Size()`)"; the sibling path "reads
`record+0xc4` — the same static field, `InnNPC.m_pData` — as a 4-byte-pointer
array", reaching the mercenary roster "through `record`'s own `+0xc4`
displacement rather than the campaign-screen-relative one `REG-SCN-064`'s sweep
used"; and "why one field serves both readings is not established here", that
is, why `record+0xc4` (`InnNPC.m_pData`, `campaignScreen+0x60c`, the field
`REG-SCN-064` names) is read as `u16` words on the gated arm and as a `CUnit*`
array on its sibling, which would have needed a heap-allocation trace or
dynamic observation EXP-0394 did not run. The former release gate
`curSel >= InnNPC.Size()`, its `curSel < InnNPC.Size()` sibling range, the
subscript `idx = curSel − InnNPC.Size()` and the "beyond the last portrait"
reader rest on that identity. Replacement (`TOWN-468`): in `R0702`,
the inn view is the first working object and the document the second; `view+0xc8` is the mercenary
count and `view+0xc4` the mercenary `CUnit*` array, while `InnNPC` is read from
the document at `+0x60c`/`+0x610`, so the release subscript is
`idx = curSel − mercCount`. The demo build was not re-read. The subscript
arithmetic, the four callers' gate shape and the absent `InnMission` size check
stand.

### REG-INNSTAGE-117

- EN, RU and the fourth root carry the identical 24 `[Mission<n>]` sections:
  same names, same count.
- EN and RU agree on every one of 14 measured per-stage keys (`InnNPC`,
  `InnMission`, `AddHero`, `AddPictureDocum`, `AddTextDocument`,
  `AutoGetMission`, `EnableMercenary`, `LastMission`, `Mercenaries`, `Payment`,
  `ShopMaxPrice`, `ShopMinPrice`, `ShopMission`, `TCMission`) at all 13
  non-empty (`InnNPC`/`InnMission`-bearing) stages: 0 differences.
- The fourth root, the owner-supplied pre-release snapshot `EXP-0393` also
  reads, disagrees with EN/RU at 5 of those 13 stages: Mission30, 40, 50, 60,
  90. The keys that move differ by stage:
  - Mission30: `ShopMaxPrice` alone (`1000`→`3000`); `InnNPC`, `InnMission`
    and `Mercenaries` match EN/RU there;
  - Mission40 and Mission50: `InnNPC`, `Mercenaries` and `ShopMaxPrice`
    together, and Mission50 also `AddTextDocument`;
  - Mission60: `InnNPC`, `InnMission` and `ShopMaxPrice`, but not
    `Mercenaries` (`[14 6 10 13 4]` on both);
  - Mission90: `InnNPC`, `InnMission`, `Mercenaries`, `LastMission` and
    `Payment`, but not `ShopMaxPrice` (`[100000]` on both).
- No fixed pair or triple of keys always moves together, and a stage can
  disagree on exactly one key.
- At Mission40 the disagreement is also a length mismatch: `InnNPC` carries 3
  elements (`[22 90 84]`), `InnMission` 2 (`[40 41]`). It is the sole stage, of
  39 stage×root combinations measured (13 stages × 3 roots), where the two
  arrays differ in length. `REG-SCN-064`'s "22/22 sections" equal-length figure
  stands on its own live/EN/RU scope; this is a fourth-root exception, not a
  correction to it.
- `LastMission`, which `REG-SCN-063` states "ships once… identical in all three
  roots", ships twice on the fourth root (`[Mission90]=1` and
  `[Mission150]=1`) against once on EN/RU (`[Mission150]=1` only).
  `REG-SCN-063` never located `record+0x11c`'s reader, so no runtime
  consequence is claimed for the second occurrence, only that the registry
  carries one.
- Extends `REG-SCN-063` and `REG-SCN-064`.

**Confidence.** High for the per-stage census: exhaustive over 24 sections × 3
roots × 14 measured keys, compared field by field on every key. High for the
Mission40 length mismatch and the `LastMission`-twice fact, both direct reads
of the registry's own record count and value, not inference.

**Unknown.** `REG-SCN-064` establishes the same-subscript read
(`InnMission[i]`/`InnNPC[i]`) but not which array's own length bounds
`R0702`'s loop. Whether the third `InnNPC` entry (npc 84) at Mission40
is ever addressed, silently ignored, or read one element past `InnMission`'s
own extent is Unknown; no disassembler was run in EXP-0393. What a second
`LastMission=1` does at runtime is Unknown. `REG-118` extends this claim with
the reader's decoded subscript.
