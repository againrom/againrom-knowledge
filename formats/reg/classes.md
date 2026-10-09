# Graphics class registries

[Reference](format.md)

## `units.reg` — unit classes (`REG-UNITS-018`, `REG-ROSTER-019` (partially retracted, corrected by `REG-ROSTER-052`))

`units.reg` contains `Unit0..Unit33`: `[Global] UnitCount=34` and
`[Global] FileCount=33`, with a shared `Files` descriptor section.
Per-class keys are listed below. — REG-UNITS-018, REG-ROSTER-052

| Key | Meaning |
|-----|---------|
| `ID` | class id |
| `File` | **scalar index into `Files[]`** — *not* a path. `units/` + `Files[File]` + `.256` (`REG-VAL-029`; see *Resolved values* below) |
| `Index` | frame index / base |
| `Palette` | palette ref |
| `DescText` | human-readable name (`Human Two-Handed Swordsman`, `Ogre (Lord)`). **Nothing in `rom.exe` reads it** (`UNIT-NAME-039`), installed strings are ASCII and fit 31 bytes plus NUL in the 0x20 inline field — so it is editor-facing metadata with a 31-character ceiling (`REG-DESC-096`) |
| `InfoPicture` | portrait leaf, **without directory or extension** — the engine builds `graphics\infowindow\<InfoPicture>.bmp`, or `<InfoPicture><tier>.bmp` with the digit dropped at tier 1, from `class+0xd8` (`UNIT-PICT-036`). Read only when the drawable's `+0x18c & 0x11` is clear, i.e. for `ID >= 0x1a`; on the thirteen human classes below that it is **dead data** (`UNIT-PICT-035`) |
| `Parent` | inherit-from class |
| `InMapEditor` | editor-visible flag |
| `MovePhases`, `MoveBeginPhases` | walk / walk-start phase counts |
| `AttackPhases`, `DyingPhases`, `BonePhases`, `IdlePhases` | per-state phase counts |
| `Flip` | **selects the sheet layout, not a blit option** — see *The loaded class record* below (`REG-UNITS-051`) |
| `Width`, `Height` | frame canvas |
| `CenterX`, `CenterY` | draw anchor |
| `SelectionX1/X2/Y1/Y2` | selection box |
| `Z` | draw layer |
| `TileSize` | footprint |
| `Dying`, `AttackDelay` | death behaviour / attack cooldown |
| `Projectile`, `ShootDelay`, `ShootOffset` | ranged-only (present on ~9–11 classes) |
| `AttackAnimTime/Frame`, `MoveAnimTime/Frame`, `IdleAnimTime/Frame` | per-phase kind-6 tracks |
| `Sound` | sound ref |

Each class selects `Files[File]`, not `Files[ID]`. The sprite path is
`units/` + that descriptor + `.256`. `Unit33` (Unarmed Fighter with Shield)
has `File=21` and shares `Unit21`'s sheet. The retracted ID-index rule in
REG-ROSTER-019 must not be used. `Flip=1` selects its own sheet-layout arm.
`Projectile`, `ShootDelay` and `ShootOffset` supply ranged/caster state.
`Anim*Frame`, `Anim*Time`, `ShootOffset` and `Sound` arrays carry independent
element counts. — REG-VAL-029, REG-ROSTER-052, REG-UNITS-051

### Human and creature cast geometry

The EN/RU `graphics.res::units/units.reg` members have the same bytes.
Human body IDs 1,2,3,4,5,7,8,9,10,11,12,13,14,15,23,24 all use
Center `(64,78)`, Selection `(48,48,80,90)` and TileSize 1.
Their effective `ShootOffset` values are:

| Class IDs | Eight stored pairs |
|---|---|
| 23 mage, 15 xbowman | 57,75;45,66;44,53;52,43;68,42;80,50;81,62;73,73 |
| 24 mage_st | 52,79;36,64;37,46;54,33;75,34;91,46;91,65;75,79 |
| 14 archer | 61,90;36,80;28,58;40,39;65,30;85,40;96,60;87,81 |
| Other twelve human IDs | Empty, including parent fallback |

These are geometry inputs, even when hero appearance loads art from its
dynamic body path. The normal cast consumer and eight-direction deltas
are defined in [magic presentation](../magic/presentation.md#cast-origin-and-human-equipment).
The registry's empty-array case uses an unscaled Selection fallback.
— MAGIC-263

Of 118 stored Units rows, 56 have parameters and twelve have a positive
spell slot. Only Goblin_Sling.4 (79), Orc_Bow.4 (65) and Bat_Sonic.4 (70)
have nonempty arrays in this caster-definition population. Classes 79/65
use `59,77;44,67;42,53;51,42;66,39;80,45;85,59;77,72`.
Class 70 uses eight `(64,64)` pairs equal to its Center. The other nine
rows use their own Selection/Center fallback. Stored spell slots do not
prove native visible-cast reachability. — MAGIC-264

Phase counts and animation-array lengths are independent. The loader iterates
the kind-6 array's `size/4` elements. For example, Goblin movement has 8 phases
and 10 entries; Ghost movement/idle have 3 phases and 4 entries. MoveBegin,
Dying and Bone phase scalars have no corresponding animation array key.
The unconditional equal-length clause of REG-UNITS-018 is retracted.
— REG-UNITS-018, REG-KIND-034

`Width`/`Height` are class canvas bounds, not sprite pixel extents. The named
human/monster classes use 128×128; Dragon uses 160×160. `CenterX` is 64 or
80 respectively, while `CenterY` is per-class. The earlier 128×64/160×80
reading is retracted. The sprite can occupy less than the class canvas.
— REG-VAL-030, REG-VAL-043

### `Z` selects the drawable class (`REG-UNITS-061`, registration scope amended)

`R0509` allocates `0x1b0` bytes at two sites and branches
on `class+0x104` = **`Z`** (five sites in four routines allocate `0x1b0` bytes; the others build `CUnit`, `UNIT-141`; `REG-UNITS-061`, amended):

- At `L07473` the routine reads `Z`, the field at `class+0x104` of the class selected by subscript from the class array at global `L02113`.
- `Z == 0` (branch at `L13017`): it allocates `0x1b0` bytes and calls `R0593`, building a `CUnit` (`CRuntimeClass` `L00620`).
- `Z != 0` (`L09187`): it allocates `0x1b0` bytes and calls `R0594`, building a `CAirUnit` (`CRuntimeClass` `L07474`).

`CAirUnit` derives from `CUnit` and overrides one behavioural vtable slot of 33
(`vt+0x38`): registration selector 3. CUnit selects 2 only when unsigned corpse
stage `+0x15a<2` and bit `0x80` of `+0x18c` is clear, otherwise 4. This is the amended form of the
alive/dead shorthand of `REG-UNITS-061`; the alternate flag matters at stage 0.
Selectors 2/3/4 store into map-view `+0x8c/+0x90/+0x9c`, respectively
(`ANIM-REGISTER-083`, `ANIM-CATEGORY-084`). Their shared drawing functions do not
imply one phase population: only selector 3 has the late separated shadow/body
sweeps (`ANIM-AIRPASS-086`). CAirUnit construction also writes `+0x10=16`; this
does not decode the registry's `Z=96` value. `Z` is a drawable-side class flag; it
does not reach the simulation, whose mover fields are per-instance bytes no registry key can touch
(`TERR-MOVE-055`). Installed nonzero `Z` values belong to `Sonic Bat` (ID70) and `Dragon`
(ID71), both `Z=96`.

### The loaded `units.reg` class record (`REG-UNITS-049`, `REG-UNITS-050`, `REG-UNITS-051`)

`R0395` allocates **`0x12c`** bytes per class and fills, in this order:

```
+0x04 ID            +0x08 section counter  +0x0c File         +0x10 Index
+0x14 MovePhases    +0x18 MoveBeginPhases  +0x1c AttackPhases
+0x20 DyingPhases   +0x24 BonePhases       +0x28 IdlePhases
+0x2c Width         +0x30 Height           +0x34 CenterX      +0x38 CenterY
+0x3c CArray<int> Move   timeline   (m_pData +0x40, m_nSize +0x44)   len copy +0x50
+0x54 CArray<int> Attack timeline   (m_pData +0x58, m_nSize +0x5c)   len copy +0x68
+0x6c CArray<int> Idle   timeline   (m_pData +0x70, m_nSize +0x74)   len copy +0x80
+0x84 SelectionX1  +0x88 SelectionY1  +0x8c SelectionX2  +0x90 SelectionY2
+0x94 Dying        +0x98 Palette      +0x9c.. palette objects  +0xac.. 0x400-B buffers
+0xbc CArray Sound (m_nSize +0xc4)
+0xd0 TileSize     +0xd4 Projectile   +0xd8 InfoPicture[0x10]
+0xe8 CArray ShootOffset (m_nSize +0xf0)
+0xfc ShootDelay   +0x100 AttackDelay +0x104 Z  +0x108 Flip  +0x10c DescText[0x20]
```

Defaults when there is no `Parent`: **`-1`**, except `IdlePhases`, `Dying`, `Palette`,
`Projectile`, `ShootDelay`, `AttackDelay`, `Z` and `Flip` (**`0`**) and `TileSize` (**`1`**).
**`File` inherits here** — the opposite of `objects.reg`, where its default is the literal
`-1` (`REG-OBJ-046`).

The three timelines at `+0x3c` / `+0x54` / `+0x6c` are **built, not stored**: the same
run-length expansion `REG-OBJ-046` describes, run once per `<track>AnimTime`/`<track>AnimFrame`
pair. Only the **move** arm of the renderer takes the animation phase modulo the length; the
idle and attack arms index the array raw.

**`Dying` is an `ID`, and it names the class whose sheet holds the corpse.** The dying and
corpse draw arms substitute `classes[Dying]` wholesale — its `File`, its `Flip` and all six of
its phase counts. A live unit's own `DyingPhases`/`BonePhases` therefore describe frames drawn
for whoever names *it*; 20 of 34 classes name themselves, and the humans name a sibling.
The live object's corpse stage is a byte: `0` alive, `1` the last dying frame, `>= 2` bone
frame `stage - 2`. See `SPR256-UNIT-024` for what the counts index into.

**`Flip` selects between two sheet layouts.** `0` -> 16 standing frames, 8 stored directions,
never mirrored. Nonzero -> 9 and 5, with the other half drawn by reflection and a mirror flag
handed to the blit. It is never passed to a blitter as data.

## Which key a class table is indexed by (`REG-KEY-044`)

The three graphics class registries are loaded into flat arrays — `units.reg` →
the global at `L02113` (`R0395`), `objects.reg` → the global at `L02099` (`R1395`),
`structures.reg` → the global at `L01425` (`R1400`) — by walking
`i = 0 … [Global]<Count> − 1` and formatting the section name `"<Prefix>%d"`. `ID` is
stored on the class (`+0x04`; structures `+0x0c`) and the counter `i` on the class too
(`+0x08`; structures have no such field). **But the array subscript the class is stored
at is per-registry, and only `objects.reg`'s is the section index** (the all-three-section-index clause of `REG-KEY-044` is amended):

| registry | array | store | subscript | `m_nSize` | slots vs classes |
|---|---|---|---|---:|---|
| `units.reg` | the global at `L02113` | `SetAtGrow(ID)` @ `L07406` | **`ID`** | `maxID+1` = 81 | 34 populated, **47 NULL** (index 0 among them) |
| `objects.reg` | the global at `L02099` | `Add()` @ `L07411` | append order **= section index = `ID`** | 82 | 82, no holes |
| `structures.reg` | the global at `L01425` | `SetAtGrow(ID)` @ `L07408` | **`ID`** | `maxID+1` = 67 | 66 populated, index 0 NULL |

| registry | `[Global]` counter | sections | `ID` domain | relation to the section index |
|---|---|---:|---|---|
| `units.reg` | `UnitCount` = 34 | `Unit0..Unit33` | `1..80`, 34 distinct | sparse ID; neither index nor index+1 |
| `objects.reg` | `ObjectCount` = 82 | `Object0..Object81` | `0..81`, 82 distinct | ID=index |
| `structures.reg` | `Count` = 66 | `Structure0..Structure65` | `1..66`, 66 distinct | ID=index+1 |

So a `.alm` placement's `ID` (a type-4 record's `kind` is a `structures.reg` `ID`, a
type-6 record's `+0x08` a `units.reg` `ID` — `ALM-CLS-036`, `ALM-CLS-038`) is used
**directly as the array subscript; no translation.** A units-array consumer must
null-check instead, because 47 of the 81 slots are empty. **No lookup by `ID` over these
arrays exists anywhere in the image** — and for units and structures none is needed,
because the array *is* the by-`ID` index.

`Parent` is an **`ID`** as well (`array[Parent]`, `L07412/65` and `L07413/28`), read
by the units and objects loaders only — the structures loader never reads the key, and
Structures have no Parent key. Inheritance is **eager, per key, at load**; nothing
resolves lazily at an accessor. **The guard is not the same for scalar and array keys**
(`REG-KEY-045`):

| keys | mechanism | inherits when | an explicit empty value |
|---|---|---|---|
| `File` + the 19 scalars `Index`…`AttackDelay` | the parent's resolved field is passed as the *default* argument of `R0452(section, key, default)` | the record is **absent** (`L13018`) | **overrides** the parent — `0` is returned as `0`; a non-int kind *throws* |
| `Move`/`Idle`/`AttackAnimTime`, the three `*AnimFrame`, `ShootOffset` | `R1061` reads into a `CArray`; the loader then re-reads the **parent's section** by the name formatted from the parent class's `+0x08` | `dest.m_nSize == 0` **after** the child read (`L07422/0b`) — presence is never tested | **does not** clear it: a present record with `kind == 0 && size < 2` takes a *success* path `SetSize(0,-1)` (`L07426`), so the guard fires and the parent's array is inherited |

An empty array key is treated as absent and cannot clear an inherited
array. This rule does not apply to scalars. For example, Unit33 (`ID=2`,
`Parent=3`) stores AttackAnimTime/AttackAnimFrame as kind 0, size 1 (`""`),
and inherits Unit0's `AttackAnimTime=[2 2 1 1 1 2 2]`.

`objects.reg` uses ID=section index and its loader appends in section order.

## Sibling registries

The per-key values below use the `REG-FMT-031`/`REG-REC-032` framing and
`REG-VAL-043`. The displaced per-key values in partially retracted
`REG-VAL-030` are invalid.

- **`objects.reg`** — **82** objects (`REG-OBJ-039`; the previously published 56 was
  `[Global] FileCount`, the 56-entry sprite-file table size, transposed for
  `ObjectCount` — a slip independent of the framing bug that the framing re-check
  surfaced). Keys: `Index`, `Phases`, `DeadObject` (−1..67), `FireObject` (−2..−1),
  `IconID`, `AnimationTime/Frame`, `Width/Height` (32..128), `CenterX/Y`, `Parent`
  (`REG-OBJ-020`, `REG-OBJ-039`). The **in-memory class record** is `0x6c` bytes and
  `REG-OBJ-046` reads it off the loader `R1395`:

  ```
  +0x04 ID        +0x08 section counter   +0x0c File      +0x10 Index
  +0x14 Phases    +0x18 Width             +0x1c Height    +0x20 CenterX
  +0x24 CenterY   +0x28 CArray vptr       +0x2c m_pData   +0x30 m_nSize
  +0x34 m_nMaxSize +0x38 m_nGrowBy        +0x3c = m_nSize +0x40 FireObject
  +0x44 DeadObject +0x48 InMapEditor(def 0) +0x4c DescText[0x1f]
  ```

  **`File` does not inherit** — its `R0452`
  default is the literal `-1`, unlike every other scalar here, whose default is the
  parent's already-resolved field. And the array at `+0x2c` is **built, not stored**: the
  loader run-length expands the two kind-6 tracks, appending `AnimationFrame[i]` exactly
  `AnimationTime[i]` times, and `+0x3c` holds the result's length. That length is the
  modulus of the renderer's animation phase and the array is what it indexes
  (`TERR-SPR-042` (amended, superseded)). Shipped lengths: 0 ×52, 24 ×2, 28 ×21, 105 ×7 — and 39 of the 82
  classes get theirs only through the `Parent` array fallback.

  `DeadObject` and `FireObject` are **not the same kind of field** (`REG-OBJ-047`).
  `DeadObject` is a subscript into this same class array — −1 on 54 classes, otherwise one
  of 21 section indices, all in range, all naming a class whose `DescText` ends "(dead)".
  `FireObject` has one recovered reader in `rom.exe`, the ambient-sound loop
  `R0484`: its whole shipped value space is `{-2 ×21, -1 ×61}`, and the reader only
  compares it against `0` and `-2` to pick which sound slot a burnt cell feeds. The
  21 classes carrying −2 are exactly the 21 that are some class's `DeadObject`. Neither
  key occurs in any other registry in the install.
- **`structures.reg`** — 66 structures (held, `REG-STR-040`, partially retracted; see
  retracted.md). The **in-memory class record**
  is `0xa4` bytes and `REG-STR-080` reads it off the loader `R1400`:

  ```
  +0x04 CSprite256 <File>.256    +0x08 CSprite256 <File>b.256   +0x0c ID
  +0x10 TileWidth   +0x14 TileHeight   +0x18 FullHeight   +0x1c Phases
  +0x20 SelectionX1 +0x24 SelectionY1  +0x28 SelectionX2  +0x2c SelectionY2
  +0x30 ShadowY     +0x34 AnimMask*    +0x38 live cells   +0x3c timeline length
  +0x40 CArray(timeline)             +0x54 Picture[0x10]  +0x64 Indestructible
  +0x68 DescText[0x20]  +0x88 VariableSize  +0x8c Usable  +0x90 Flat
  +0x94 LightRadius  +0x98 LightPulse  +0x9c CString path  +0xa0 sheets-loaded
  ```

  The structure loader has no `Parent` key and performs no class inheritance
(`REG-KEY-045`). `File` is a path string: the class stores
`"graphics\structures\"+File` and lazily opens `<path>.256` and
`<path>b.256` on first draw (`REG-STR-081`).

`AnimMask` indexes the sprite grid, not the footprint. Its buffer is
`FullHeight*TileWidth+1` bytes, `'-'` marks a cell that never animates,
`class+0x38` counts the other cells, and drawing indexes it by
`k*TileWidth+c`. Installed masks use `{'-','*'}`. `AnimTime`, `AnimFrame`
and `AnimMask` are read only when `Phases>1`; the timeline uses the same
run-length expansion as objects.reg. `Phases>1` alone does not establish
animation: the keys may be absent. — REG-STR-082, SPR256-STR-041

  Installed dimensions are `FullHeight=1..6`, `TileHeight=1..6` and
`TileWidth=1..11`; these are not format limits. The FullHeight/TileWidth
ranges in partially retracted `REG-STR-040` are replaced by these values.
Other keys include `AnimMask`, `Phases`, `Flat`, `Indestructible`, `Usable`,
`VariableSize`, `LightPulse`, `LightRadius`, `ShadowY`, `SelectionX1..Y2`,
`IconID` and `Picture` (`REG-STR-021`). Picture occupies 16 bytes at
`class+0x54` and names `graphics\infowindow\<name>.bmp`, without a tier
suffix (`UNIT-PICT-037`). The RU install lacks the `magic`, `ruins` and
`sphinx` leaves referenced by structure classes (`REG-PICT-083`).

- **projectiles.reg:**31 installed definitions. Present SFX values are 1,
  a presence flag; the 12..62 interpretation in `REG-PROJ-041` is withdrawn.
  Installed Width/Height are 12..128, not limits. File is a kind 0 string path
  such as `archer\arrow`, not a Files[] index. Other keys are Phases,
  RotationPhases, Homing, Flip and Palette. — REG-PROJ-022, REG-PROJ-041
- **material.reg:**16 installed Material sections, each with Path and no
  Global counter. — REG-MAT-042

## `projectiles.reg` — and why its `ID` is an address

`REG-PROJ-086`, read from the loader `R0606` rather than fitted. Its routine label is
superseded: `R1366` is the teardown and `R1423` the grow. Every field, default, count
and consumer stands.

```
[Global] Count = 31
[Projectile<n>]   n = 0 .. Count-1        (the section index is NOT the id)
  File            string, e.g. archer\arrow  -- a path, as in structures.reg, not a Files[] index
  A16             1 -> the sheet is .16a, 0 -> .256          default 0
  ID              the key everything else uses                default -1
  Phases          frames per facing                           default -1
  RotationPhases  facings in the sheet                        default 16
  Width, Height   the DRAW's centring offsets, not the art    default 64, 64
  Palette         0 -> use the shared projectiles.pal          default 0
  Homing          stored at +0x2c and read by nothing          default 0
  Flip            halves the sheet: 9 stored facings, 7 mirrored   default 0
  SFX             1 wherever present on the shipped file       default 0
```

The loader grows its array to `ID + 1` and stores the record at `array[ID]`, so **`ID` is the
address**: the shipped 1..62 domain over 31 rows leaves 32 null slots, and a null slot makes every
consumer draw nothing rather than fault. Two rows may name one `File` (the shipped file does, twice
for `catap2\sprites`). Two `projectiles/*` directories — `smoke0` and `smoke1` — are loaded by
literal path outside the loop and are in no section.

The id space is shared with `units.reg`, whose `Projectile` key selects a unit's shot from the same
array; 9 of the 34 shipped classes carry one, and two of those point at a spell's art.
