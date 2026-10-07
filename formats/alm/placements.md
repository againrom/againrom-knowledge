# Placements and mission records

[Reference](format.md)

## Content sections type 4–9 — `ALM-CNT-017`, `ALM-UNIT-018`, `ALM-OBJ-019`,
`ALM-GRP-020`, `ALM-TRIG-021`, `ALM-CODE-023`; payload framing is
`ALM-FRAME-031`. The type 8 geometry interpretation in `ALM-TRIG-022` is retracted.

Like the grids, the content payloads are pure: the first record starts at payload+0.
`X`/`Y` fields are `u32` fixed-point `/256` (integer part = tile, low byte usually
`0x80` = tile centre).

**Placement anchor → terrain cell is a bare `>>8`** (`ALM-PLACE-033`): no origin is
subtracted, no border inset added, no rounding applied. The engine shifts the object
anchors in the loader itself (`SAR ,8`) and the unit anchors when it walks the loaded
unit list. Combined with the grid base above, a placement at `(X,Y)` stands on tile
`(X>>8, Y>>8)` of every grid layer.

### type 4 — placed structures/buildings (`ALM-OBJ-019`, `ALM-OBJ-034`, `ALM-CLS-036/037`)

`count = meta+0x20` is the loader's case-4 loop bound. Each record is 20
bytes, plus an eight-byte extension when `kind == 0x21`.

| Off | Type | Field | Notes |
|-----|------|-------|-------|
| +0x00 | u32 | **X** | `/256`, `0x80`-centred. **This is the axis the tile grid strides by 1** and the block plane's **low** byte — the same axis as `meta+0x00` and as the Buildings column `sizeX` (`ALM-OBJ-061`) |
| +0x04 | u32 | **Y** | `/256`. The other axis: the grid's row index and the plane key's high byte |
| +0x08 | u32 | **kind** | Installed values are `1..66` (65 distinct, none zero). Full u32 equality to `0x21` selects the eight-byte extension; the loader stores low16, the walker tests that word for `0x22`/`0x23`, and either constructor receives low8. These are three different width boundaries (`ALM-127`) |
| +0x0c | two raw bytes | **placement word** | Preserved at placement-record+0x06 on complete input transfer. Shop interprets signed i16 ×1000 as its stock price cap; Building only tests zero versus nonzero here. The universal durability label is retracted (`ALM-127`, `ALM-128`, `SHOP-CAP-004`) |
| +0x0e | u32 | **owner** | low u16 used: a **1-based physical slot** in the type-5 array |
| +0x12 | u16 | **trigger target id** | sign-extended into the live object's `+0x10`; it is the id a type-7 `Target_Structure` parameter names (`ALM-TRIG-046`) |
| +0x14 | u32 | *ext, low axis* | present iff `kind == 0x21`. **Low byte only** → `obj+0x60`, the extent along the `+0x00` axis (`ALM-OBJ-062`). Bytes `+0x16`/`+0x17` are not consumed |
| +0x18 | u32 | *ext, high axis* | **low byte only** → `obj+0x61`, the extent along the `+0x04` axis; overrides the definition's own footprint — `structures.reg`'s `VariableSize`. Bytes `+0x1a`/`+0x1b` are not consumed |

The override is selected by `(w & 0xff) + (h & 0xff) > 0` at `L07853`…`L07854`, **not**
by the kind (`TERR-STRUCT-090`): an extension record with two zero bytes would take the
table arm. Installed extensions use class 33 and two nonzero extent bytes.

**Which C++ class a record becomes** (`ALM-CLS-037`, read from `R0489` at
`L02105`…`L12973`; the names are the MFC runtime-class table's own, `SHOP-CLS-001`):

```
stored kind word == 0x22 or 0x23 -> Shop, 0x74 bytes (R0490, vtableR0491)
any other stored kind word     -> Building, 0x6c bytes (R0485, vtableR0492)
```

`0x22`/`0x23` are `structures.reg` `ID` 34/35 = **`Shop 1` / `Shop 2`**, the two `Usable`
shops. `kind == 0x21` — the
*extension* discriminator — is an **object**: the two discriminators are different
questions with adjacent immediates. **`Shop` derives from `Building` and its constructor
opens by calling `Building`'s at `L02198`**, so a shop resolves and attaches a footprint
like any other placement — with the two extent bytes pushed as literal `0`, i.e. always
from the table (`ALM-CLS-063`).

**How `kind` reaches a class.** The engine does not subscript `structures.reg` with it
directly: `R0486` uses the byte as a **1-based** index into a placeable-definition
table (`L02102`, guarded `kind != 0 && kind <= count − 1`), and that definition supplies
the footprint (`sizeX × sizeY` in tiles + the two per-cell masks), the HP pair and one
further byte. That table is the **Buildings collection of `world.res:data/data.bin`**
(`DAT-LOC-001`, `DAT-BLD-005`); the 1-based law is the file's own entry-0 skip.

#### Placement-word consumers and production

| Reached runtime class | Read and transform | Result before later simulation |
|---|---|---|
| Building, including the kind33 extension arm | Setup copies definition column3 into words +0x44 and +0x42. The walker tests signed placement-record+0x06 for zero; only zero is stored into object+0x42 | Zero overrides the initialized current field. Every nonzero word leaves it at the definition value; there is no proportional nonzero assignment. +0x44 remains the definition value (`ALM-128`) |
| Shop, stored kinds34/35 | Sign-extend placement-record+0x06; multiply by 1000; setter writes object+0x70 and forwards to template+0x90 | The template's price maximum receives the signed product. This operation does not clear the inherited +0x42 field. The campaign town has a separate producer (`ALM-128`, `SHOP-CAP-004`, `SHOP-MISSION-018`) |

The loader stores a placement pointer in map+0x304. After constructing and
registering the runtime object, the walker frees that placement and clears
its array slot. The original nonzero Building word is not passed to its
constructor or registration. Only the selected state above survives this
walker; this is not an image-wide no-alias claim. Building's +0x14 virtual
can later increment +0x42 for selectors15/16 under its existing counter and
range predicates, so zero is not established as an irreversible state.
No other runtime subclass is selected by this type4 walker. — ALM-127,
ALM-128, UNIT-STRUCTZERO-080

| Editor operation | Placement-word rule |
|---|---|
| Complete type4 read | Read2 bytes, sign-extend for the factory argument, then retain low16 at CStructure+0xf6 and +0xfa. Both ordinary and kind33 factories call the same initializer; the latter constructs an extent-capable editor object (`ALM-129`) |
| Selected type4 write | For each object admitted as an editor CStructure, emit2 bytes directly from +0xf6. No per-word scaling, zero normalization or signed-range check occurs at this write site (`ALM-129`) |
| Confirm existing-structure property | The selected handler opens dialog226 with +0xf6 sign-extended into its integer control1041. If accepted and the selected class remains nonzero, clamp the integer only above 3000 at this local store, then retain low16 in +0xf6/+0xfa. A successful property confirmation can therefore change32767 to 3000 before saving; an unchanged direct write retains0x7fff (`ALM-130`) |

Full input transfer and successful factory admission are preconditions of
the editor preservation rule. UI input validation, delivery of the selected
property command, other editor producers, short-read recovery, archive
backing-stream identity, and native behavior after the first field consumer
remain Unknown. No shared health, charge or level interpretation follows
for all structure definitions. — ALM-127, ALM-128, ALM-129, ALM-130

### type 5 — player/group roster (`ALM-GRP-020`)

`count = meta+0x1c`. **76-byte** fixed records. Every field below is a read
boundary in the loader's case-5 sequence (`4+4+4+32+16×2 = 76`).
This field map retains ALM-GRP-041, partially retracted for the
inherited universal owner interpretation; the control field has the
separate predicates below.

| Off | Type | Field | Notes |
|-----|------|-------|-------|
| +0x00 | u32 | **colour slot** | → `player+0x08`, carried `+1` to `Player+0x44`; the unit body draw indexes the 17 shade objects with it (`ALM-PLAYER-069`, `PAL-SHADE-012`, `PAL-SHADE-013`). Sparse and unordered, e.g. `Cross.ALM` `8 2 3 5 1 6 4 13`. The owner lookup `R0425` compares the record's own 1-based ordinal (below), **not** this word |
| +0x04 | u32 | **control word** | Four-byte Read destination `CPlayer+0x30`; complete copies reach `Player+0x28`. No Boolean validation occurs in these local transfers. The old 0/1 inventory is a corpus result; `UNIT-OWNER-009` is partially retracted for treating it as a general value/authorship bound. `ALM-139`, `SESS-072`, retained corpus/ordinal chain in `ALM-PLAYER-069` |
| +0x08 | u32 | **scalar; gameplay effect Unknown** | Loaded into CPlayer+0x0c. Preserve all four bytes. The installed 0/5000 values are a corpus observation, not a two-value schema. — ALM-SCALAR-087, ALM-SCALAR-089 |
| +0x0c | char[32] | **name** | NUL-terminated ASCII: `Self, Monsters, Villagers, Neutral, Enemy, Beasts, Guards, Peasants, Orcs, Trolls, …` |
| +0x2c | u16[16] | **relations** | one per editor player slot (the editor caps at 16) — the diplomacy row |

The game loads each record into the map-owned `CPlayer`, object size `0x48`.
Its vtable at `L02273` has five function entries and a null; the EN editor
has a separate four-function-plus-null table at `L02281`. These are not the
session `Player` type. The game destinations are `+0x08` colour, `+0x30`
control word, `+0x0c` scalar, `+0x10` name, `+0x34` diplomacy array and
`+0x04` loader ordinal. Non-copy constructors initialize the scalar to zero;
the loader supplies the authored word or its absent-section default. —
ALM-CPLAYER-090, ALM-SCALAR-087, ALM-EDITOR-152, UNIT-VPLAYER-022

The game map-copy constructor can allocate independent scalar-bearing
CPlayers in a second map. Destroying the original roster does not destroy
those copies. Native reach of that constructor and any later gameplay
consumer remain Unknown. A world field also retains an address after map
destruction; the address does not establish a live object. The earlier
selected direct-displacement absence is not a whole-image exclusion. —
ALM-COPY-151, ALM-ALIAS-154, ALM-SCALAR-088

The EN editor copies this scalar and writes its current four-byte value.
Roster dialog control `0x407` is bound to a signed integer with a requested
validation range of `0..2147483647`; selected handlers transfer that value
to the roster scalar. Adding an entry initializes its scalar to `5000`.
The dialog edits the live roster while retaining a map copy for restoration
after a modal result other than `1`. Actual accept/cancel, validation and
file round-trips remain unobserved. These producer facts establish no game
currency or gameplay interpretation. — ALM-EDITOR-152, ALM-EDITOR-153

A record's **physical slot + 1** is what type-4 `+0x0e` and type-6 `+0x14` store
(`ALM-OWN-039`); the loader itself writes `slot+1` to `player+0x04`. The
`+0x00` color slot is a separate value.

Both non-copy constructors initialize `CPlayer+0x30` to zero; its copy
constructor preserves all four bytes. Type5 requests four bytes into it
without testing the read result before the next field. A full transfer is
conditional on the stream supplying four bytes; the wrapper's short-read
boundary remains as stated by `ALM-STREAM-100`. — ALM-139

Roster transfer reuses the manager's first Player only for roster slot1,
server mode zero and manager count1. This reuse preserves `Player+0x08`;
the new-object arm copies the authored ordinal before registration. Both
arms copy the full control word. Participant join instead conditionally
looks up `Self` and writes control zero on either its reused or new object.
These local rules do not determine the native order of the two operations.
— SESS-072

| Consumer | Operation on the session word | Scope |
|---|---|---|
| Actor clauses | Whole-dword zero tests, including route-budget eligibility and victim experience exclusion | 27 catalogued owner reads include 26 predicate sites and one read whose flags/value die before the shared call; the set is bounded — ALM-140 |
| Matrix join | Store low byte, then replace it with zero only for full word `2` | `0x100` maps to zero; `0x102` maps to two. Template selection tests the stored byte for zero — SESS-073 |
| Interface message | Emit full-word `!=0`, then change only destination `CPlayer+0x30` bit0 | This message-derived CPlayer has a different origin from the authored record — SESS-074 |
| Player persistence | Write/read the full dword through four-byte buffer operations | Full native save/reload is a separate observation — SESS-074 |

The alias search does not cover every advanced pointer or whole-object
copy. Native arbitrary-map admission, transfer/join identity, later UI and
editor property authoring remain Unknown. The separate wire+0x08 scalar
lifetime is unchanged by these control-word results. — SESS-075

### type 6 — placed units (`ALM-UNIT-018`)

`count = meta+0x24`. Version 990 uses 70-byte records. File offsets differ
from runtime offsets: file+0x2c..+0x40 maps with displacement−4, so file+0x35
is runtime+0x31 and file+0x3b is runtime+0x37 (`ALM-TAILMAP-079`). The
four u16 values at+0x20..+0x27 use sentinel 0xffff (`ALM-TAILU16-082`).
The inventory-sentinel interpretation of the byte runs is retracted: those
runs are stat overrides, and installed bytes+0x2c..+0x3f do not use 0xff.
— ALM-TAILRUN-083

| Off | Type | Field | Notes |
|-----|------|-------|-------|
| +0x00 | u32 | X | `/256`, `0x80`-centred |
| +0x04 | u32 | Y | `/256` |
| +0x08 | i16 | **class** | a units.reg ID; installed IDs lie in 1..80 and are not section indices |
| +0x0a | u16 | **class (2nd key)** | only consulted when `+0x08 ∈ {0x1a,0x1b}` or `≥ 0x40`; also the `npc.reg` index on the NPC path |
| +0x0c | u32 | flags | arm-specific: bit 0 selects NPC inside the Human band; ordinary Human bit 2 controls actor+0x4b high bit behind the secondary-key rule below; bit 7 supplies a constructor equipment gate (`ALM-FLAGPATH-109`, `UNIT-PLACEGEAR-098`) |
| +0x10 | u32 | **definition id** | present only when `formatVersion > 0x3da`; when nonzero and `!= 0xcdcdcdcd` it **overrides** `+0x08`/`+0x0a` (matched on the definition's parameter `0x18`). **But it is read only when `+0x0c` bit 0 is clear** — `R0151` tests the NPC flag *outside* the definition-id test, so a record carrying both takes the npc arm and this field is never read (`MISSION-ARM-006`) |
| +0x14 | u32 | **owner** | low u16 used: 1-based physical slot in the type-5 array |
| +0x18 | u32 | **type-8 link** | 1-based, bounded by `meta+0x2c`; `0` = none. The loader writes this record's `+0x40` word into entry `[value−1]` of `mapObj+0x2dc` |
| +0x1c | u32 | — | installed values 0..14; not a field-width limit. No direct reader in the established game record-holder family; stream escape remains bounded. The editor reads, copies and writes an opaque dword. Meaning Unknown (`ALM-TAILDIR-081`, `ALM-PLACESTREAM-111`, `ALM-PLACEEDITOR-112`) |
| +0x20 | u16 | **current health** | absent value `0xFFFF`; applied to `actor+0x94` (`UNIT-PLACE-034`), signed and unclamped, with no skip for 0 (`UNIT-144`). Optional in installed maps (`ALM-TAILU16-082`): 30 records authored, 21 at 0 and 9 at -10 (`UNIT-143`). A 0 enters the death arm of `HERO-DEATH-026`; the frame-0 draw is Unknown |
| +0x22 | u16 | — | absent value `0xFFFF`; authored on exactly the records `+0x20` is and with a different value set, and **no reader** (`ALM-TAILU16-082`) |
| +0x24 | u16 | **current mana** | absent value `0xFFFF`; applied to `actor+0x9a` (`UNIT-PLACE-034`). Installed value 0xffff; the consumer exists but installed maps leave it absent (`UNIT-PLACEIDLE-088`) |
| +0x26 | u16 | — | Installed value 0xffff, with **no reader** (`ALM-TAILU16-082`) |
| +0x28 | u16 ×2 | — | present only when `formatVersion > 0x3b6`; they land at `rec+0x4c`/`rec+0x4e` and **nothing reads either** — `disp:4e` is empty image-wide. `+0x28` is `0` on every record; `+0x2a` has multiple installed values (`ALM-TAILVER-084`) |
| +0x2c | u8 ×4 | **stat overrides** | Body, Mind, Spirit, Reaction in that file order → `actor+0x84`/`+0x88`/`+0x8a`/`+0x86`, zero = absent (`UNIT-PLACE-034`, file order fixed by `ALM-TAILMAP-079`) |
| +0x30 | u8 | — | lands at `rec+0x2c`; **no reader**, and `0` on every record of both roots (`ALM-TAILRUN-083`) |
| +0x31 | u8 ×2 | — | → `actor+0xbe` and `actor+0xc0`; `+0x32` present only when `formatVersion > 0x3d8`. Installed+0x31 is 0; +0x32 can be nonzero (`UNIT-PLACE-034`, `UNIT-PLACEIDLE-088`) |
| +0x33 | u8 ×2 | — | land at `rec+0x2f`/`rec+0x30`; **no reader**, and `0` on every record of both roots (`ALM-TAILRUN-083`) |
| +0x35 | u8 ×6 | **skill overrides** | → `actor+0xa8 + 2i`, zero = absent — but the loop's index starts at **1**, so only `+0x36`…`+0x3a` are applied and `+0x35`, the `Skill.General` slot, is stored and never read (`UNIT-PLACESKILL-086`) |
| +0x3b | u8 ×5 | **elemental resistances** | all five applied to `actor+0xc4 + 2i`, zero = absent, after the spawner's re-derivation call (`UNIT-PLACERESIST-087`) |
| +0x40 | u16 | **unit id** | what a type-7 `Target_Unit` parameter names (`ALM-TRIG-046`). Also copied into the linked type-8 entry's `+0x3c` |
| +0x42 | u32 | **group id** | what a type-7 `Target_Group` parameter names (`ALM-TRIG-046`); shared by group members. The map keeps its running maximum at `mapObj+0x90`, i.e. the next free **group** id (`ALM-UNIT-048`; the unit/group-label clause of `ALM-UNIT-040` is retracted) |

The record is **70 bytes only at `formatVersion == 990`**: three of the reads above are
version-gated (`> 0x3b6`, `> 0x3d8`, `> 0x3da`), so a different version yields a different
stride (`ALM-UNIT-040`).

**Who ever reads a placed-unit record.** The loader stores the record's *pointer* into a
4-byte-element `CObArray` at `mapObj+0x318`, and exactly three routines can hold one: the
loader itself, the placement spawner `R0151`, and the map object's destructor,
which frees the record without reading a field. Between them they read `+0x00`, `+0x04`,
`+0x08`, `+0x0c`, `+0x10`, `+0x14`, `+0x20`, `+0x24`, `+0x28`…`+0x2b`, `+0x2d`, `+0x2e`,
`+0x31+i` (`i=1..5`), `+0x37+i` (`i=0..4`), `+0x3c`, `+0x40`, `+0x44` and `+0x48` of the
runtime record — nothing else (`ALM-TAILHOLD-080`). The no-reader statements are limited to this placed-record holder family.

Installed placements leave the current-mana and actor+0xbe overrides
absent. Other tail overrides are active, so the full field program is
required when those values are authored. — UNIT-PLACEIDLE-088

**Which shipped file a placement reads** (`MISSION-ARM-006`, `R0151`). The outer
discriminator is the class key, not a flag, and the four arms are ordered:

```
+0x08 >= 0x1a                          -> data.bin Units,  matched on params 0x1d/0x1e
+0x08 <  0x1a and +0x0c bit 0          -> npc.reg [npc<+0x0a>] -> its DataBinID
                                          -> data.bin Humans on param 0x18 (serverID)
+0x08 <  0x1a, bit0 clear, +0x10 != 0  -> data.bin Humans on param 0x18, +0x10 being the id
+0x08 <  0x1a, bit0 clear, +0x10 == 0  -> data.bin Humans on param 0x10 (typeID)
```

### Placement flag consumers and initialization boundary

The local flags consumers preserve the resolution order above. Constructor
arguments and actor stores are distinct operations (`ALM-FLAGPATH-109`):

| Selected arm | Constructor third argument | Secondary key and bit 2 |
|---|---|---|
| Human type key | exactly `flags & 0x80` | copy low byte to actor+0x4b, then replace its high bit from bit 2 |
| Human definition ID | exactly `flags & 0x80` | same stores only when the secondary low word is nonzero; 0x100 passes despite its zero low byte |
| NPC in Human band | literal zero; separate Hero mode is another argument | neither store above |
| Units | different constructor, without this argument | neither store above |

The Human initializer tests that third argument for zero. On the ordinary
matched-definition path, nonzero skips the ten definition-equipment cells;
zero allows their nonempty entries to construct/equip. A special identifier
prefix can also suppress this loop, so zero alone does not promise equipment.
This local gate does not establish native inventory or lifetime behavior
(`UNIT-PLACEGEAR-098`). No purpose is assigned to the other 29 bits.

The installed version-990 census has flags 0/1/4/5 on both roots. Counts are
7297/14/782/1 across 8094 EN placements and 3913/14/63/1 across 3991 RU
placements. Bit 7 is absent from this population; corpus absence does not
make it invalid (`ALM-FLAGCORP-110`).

The game stream wrapper forwards the interior record destination unchanged
after clamping length. Its loose-file receiver forwards it to ReadFile;
neither selected read body stores the destination durably. The archive
receiver construction, imported service behavior and other aliases remain
open. The holder-family negatives above are not global lifetime exclusions
(`ALM-PLACESTREAM-111`).

At version 990 the EN editor's paired reader and writer each transfer 28
fields totaling 70 bytes. Wire+0x1c uses editor object+0x66; a reached copy
constructor also copies that dword. The questioned tail members are
wire+0x22/+0x26/+0x28/+0x2a to editor+0xfa/+0x10c/+0xf8/+0x10e,
wire+0x30/+0x33/+0x34 to editor+0x112/+0x113/+0x114, and the six-byte
wire+0x35 run to editor+0x118. Their field meanings and GUI property
producers remain Unknown. Paired operations on a supplied object do not
establish a native editor round trip (`ALM-PLACEEDITOR-112`).

Spawner-return overrides are conditional state. The direct initialization
caller later processes loot; its actor equip dispatch can reach Weapon
equip and another Human derive. Actual first-tick survival, the complete
actor lifetime and later save persistence remain Unknown
(`UNIT-PLACEFRONTIER-099`).

The definition lookup `R0498` searches backwards, skips index 0,
and returns 0 on a miss. A miss therefore returns a valid index rather than
an error sentinel. — MISSION-DEF-007

Those four arms also select the Humans constructor mode. Both direct Humans arms pass zero. The
npc arm queries the literal flag `Hero` on its `npc.reg` section and passes that Boolean. The
constructor always streams the Humans row's slot-16 `typeID` to `actor+0x0e`, then replaces it with
the player-character `gender+0x21/+0x23` value only when this mode is non-zero. Consequently a
Humans placement is not by itself a persistent party actor: mission 20's npc without `Hero` and its
three definition-id Humans retain table types outside the mission-end keep band
(`PARTY-M20-030`, `PARTY-M20-031`).

### type 7 — the trigger script: three counted arrays (`ALM-TRIG-044`…`047`)

The primary loader reads each internal count into EBP-0x50 before that array's
own comparison. Each complete Read4 replaces metadata+0x28 or the previous
internal count; the resulting populations are A/C/T and the consumed size is
`12 + 796*A + 796*C + 184*T`. Metadata+0x28 is not an additional loop bound.
EOF or short lower reads may preserve old local bytes, including a prior count,
which describes malformed-input execution rather than an alternative grammar.
— ALM-COUNT-195, ALM-STALE-196

The payload is **not** one list of nodes. It is three arrays, back to back, each preceded
by its own `u32` count:

```
[u32 nAct ][ nAct  x 796-byte node    ]   the "THEN" vocabulary, Description Instants.ini
[u32 nCond][ nCond x 796-byte node    ]   the "IF"  vocabulary, Description Checks.ini
[u32 nTrg ][ nTrg  x 184-byte trigger ]   binds conditions to actions
```

`ALM-TRIG-021`'s `entryCount` is `nAct`; the previously reported "8-byte trailer" is
`[u32 nCond = 0][u32 nTrg = 0]`, which is what a skirmish map with no script looks like.

**Node — 796 bytes, identical for an action and a condition:**

| Off | Type | Field | Notes |
|-----|------|-------|-------|
| +0x000 | `char[64]` | editor label | the map author's own name; NUL-terminated inside a fixed buffer whose tail is uninitialised editor memory |
| +0x040 | `u32` | **opcode** | the `ID` declared in the `.ini` (`EDITOR-023`) |
| +0x044 | `u32` | **id** | unique within each installed list; what a trigger references |
| +0x048 | `u32` | — | installed value 0 |
| +0x04c | `u32[10]` | `value[Par0..Par9]` | **indexed by slot, not packed** |
| +0x074 | `u32[10]` | `type[Par0..Par9]` | `0` = slot unused |
| +0x09c | `char[64][10]` | `pname[Par0..Par9]` | the `.ini`'s `Par<i>_NAME`; `"<None>"` when unused |

Parameter **type codes** (`ALM-TRIG-046`), and what a value of that type names:

| Code | `.ini` type | The value is |
|---|---|---|
| 1 | `int`, `Enum` | a literal |
| 2 | `Target_Group` | the type-6 record's `+0x42` group id |
| 3 | `Target_Player` | 1..8 |
| 4 | `Target_Unit` | the type-6 record's `+0x40` unit id |
| 5 / 6 | `X` / `Y` | a plain tile coordinate in `[0,W)×[0,H)` for installed records |
| 7 | `Const` | the literal the `.ini` declares for this variant; it is what selects among the eleven signatures sharing opcode 6 |
| 8 | `Target_Item` | domain Unknown (values 2..36) |
| 9 | `Target_Structure` | the type-4 record's `+0x12` word |

`Target_Building` is declared by the editor and placed by no shipped map.

**Trigger — 184 bytes:**

| Off | Type | Field |
|-----|------|-------|
| +0x00 | `char[64]` | name |
| +0x40 | 64 B | opaque bytes; typed role and native producer Unknown (ALM-TRIGSTORAGE-163) |
| +0x80 | `u32[3][2]` | three `(left, right)` **condition ids** — complete pairs in installed records |
| +0x98 | `u32[4]` | up to four **action ids** |
| +0xa8 | `u32[3]` | one **comparison code** per pair — `0 ==`, `1 !=`, `2 >`, `3 <`, `4 >=`, `5 <=`, dispatched through a 6-entry table; the three pairs are **ANDed with short-circuit** and a code above 5 is permanently false (`TRIG-CMP-006`) |
| +0xb4 | `u32` | the **once flag**: `1` = fire at most once per session (the byte latch at `session+0xbec4+index` gates re-entry), `0` = re-evaluate and re-fire every full tick (`TRIG-FIRE-007`) |

Trigger references contain node IDs, not array indices. Resolve action
IDs within the action list and condition IDs within the condition list.
— ALM-TRIG-045, ALM-TRIG-046, ALM-TRIG-047 (amended: heap/non-text gloss withdrawn)

The game and the editor readers each transfer the complete184-byte record when
the lower read completes. The game initializer and the editor no-buffer arm
zero-fill all requested records. Editor growth preserves the old records and
zero-fills only the newly added records; unused capacity is outside that span. Existing-record dialog initialization
and apply copy46 dwords through dialog+0x338; the second64 bytes are at
dialog+0x378..0x3b7. The writer submits one184-byte item per record. Native
creation, UI/name-copy behavior and a durable editor round trip remain Unknown.
— ALM-TRIGSTORAGE-163, ALM-TRIGEDITOR-164, ALM-TRIGZERO-165

### type 8 — authored loot, and type 9 — caster payload (`ALM-SACK-065`, `ALM-TRIG-049`)

Neither is a named-node tree, and **type 8 has no count word at all**. Both walks consume
the complete payload.

Type8 is authored loot. Type9 is caster data. The box/circle interpretation
of type 8 in retracted `ALM-TRIG-022` is invalid: trigger geometry reads the
trigger's own parameters. `R0461` consumes two distinct lists: case 9
fills `mapObj+0x2f0` for spellbook/building-caster processing; case 8 fills
`mapObj+0x2dc` for container creation and sack placement through
`R0472`. On installed single-player campaign maps, the type 8 path is
the only load-time sack source. — ALM-SACK-065, ALM-SACK-066,
ALM-LIM-067, ALM-ORD-068

```
type9: [u32 count] then count records of 26 + 6*n bytes
       +0x00 u32 tag
       +0x04 u32 X · +0x08 u32 Y
       +0x0c u16 A · +0x0e u16 B · +0x10 u16 C
       +0x12 u32 spellRaw · +0x16 u32 n
       then n elements of 6 B = (u16 kind, u16 low, u16 high)

type8: no count word - the record count is meta+0x2c (the loader's case-8 bound).
       Records of 20 + 10*n bytes; the head is 20 B only at formatVersion >= 0x3dd,
       16 B below it (L02137), and every shipped map is 0x3de.

       +0x00 u32 n      element count. A LOCAL: the arm never stores it (L02139)
       +0x04 u32 owner  -> obj+0x3c. 0 = a sack on the ground; non-zero = the id of an
                          actor that already exists, looked up in a map R0461
                          builds from the registry L00240 keyed by actor+0x8.
                          OVERWRITTEN at load by case 6 for any record a type-6
                          placement's +0x18 links to (L02127) -- 43 of 43 agree
       +0x08 u32 X      -> obj+0x40. L02140 shifts the fixed-point value down by 8 bits -> position byte 0
       +0x0c u32 Y      -> obj+0x44. L02141 shifts the fixed-point value down by 8 bits -> position byte 1
       +0x10 u32 gold   -> obj+0x48. R0472 arg 3; L02142 adds it to sack+0x3c
                          (version-gated: 0 below 0x3dd)
       then n elements of 10 B:
       +0x00 u32 item   low u16 is a PACKED code -- R0519 cuts bits 8..11 (class),
                        12..15, 5..7, 0..4 and R0520 allocates 0x84 Weapon (1),
                        0x68 (2), 0x68 (3..13), 0x50 Item (14), null otherwise. Class 14
                        takes the whole low BYTE as its index instead of bits 0..4
       +0x04 u16        read ONLY on the owner!=0 arm (L02223): 0 -> append to the
                        actor's carrier at actor+0x7c, non-zero -> actor vt+0x3c
       +0x06 u32 link   1-based index into the type-9 list, 0 = none (L02224)

       A record with owner==0 and a non-empty container becomes a Sack (new 0x44) at the
       cell; if a sack is already there R0472 merges the container and adds the gold,
       so two records on one cell become one sack.
```

`ALM-EFFLINK-072` and `ALM-EFFPOP-073` read the link operation and type 9 consumer end to end. The loader and consumer each
subtract one from a positive file link before indexing the type 9 list; zero means none and file
value one reaches zero-based record zero. Case 8 writes the referring element ordinal and owning
type 8 pointer to the target type 9 object's runtime `+0x00/+0x04`; these are runtime fields, not
wire offsets.
The consumer accepts a linked recipe only when its file X/Y are both zero and the backlink exists.
It then appends Effects in this order:

1. non-zero A becomes kind `low8(A + 43)`, with B/C as its operands (`ALM-T9CONTROL-175`);
2. non-zero low word of `spellRaw` becomes kind 41, or kind 42 for a Book;
3. every tail element becomes `(kind, low, high)` in file order, with input kind 41 remapped to
   runtime kind 49.

This construction finishes before the owner-zero branch. A ground sack and actor stock therefore
receive the same enchanted item operation; the link is not a ground-only or owner-only field.

### `rom.exe` cross-reference (`ALM-CODE-023`)

The separate type-9 cell-entry arm is `UNIT-M10CELL-054`. On version-990
records, X or Y nonzero and unsigned A below 4 select it. The target key is
`u8(X) | (u8(Y)<<8)`. Its six-byte payload is:

| Byte | File source | Role |
|---|---|---|
| 0 | low byte of spellRaw | spell/operation |
| 1 | byte 2 of spellRaw | power |
| 2,3 | low bytes of tail entry 0 words 0,1 | caster source x,y |
| 4,5 | low bytes of tail entry 1 words 0,1 | operation-26 relocation x,y |

This arm reads two tail entries without a length guard. Their third words do
not feed this payload. These coordinate roles do not replace the linked-item
Effect interpretation: X=Y=0 is that other arm's eligibility condition. A bit 2
selects a separate building-caster branch with a building-key lookup; the cell
arm performs no such lookup. Mission 10 has exactly two cell selections,
(22,64) and (21,63), both spell 13 / power 1, and no building-caster selection.
The remaining nine records have zero X/Y. Runtime admission and lifetime are
`UNIT-M10ENTRY-055` through `UNIT-M10LIFE-057`.

### Type9 control predicates

`ALM-T9CONTROL-175` bounds the complete builder's use of A. Let N mean X or Y
is nonzero, and L mean the runtime backlink at+4 is nonzero. The first loop
executes these tests in order:

| Predicate | Operation |
|---|---|
| N and zero-extended A<4 | Cell write; L does not gate it |
| not L and A&8 | Actor lookup and Spellbook insertion |
| N and not L and A&4 | Building lookup and VirtualCaster construction |

Failed actor lookup skips the rest of that record. Thus A=12 reaches both
searches only when the actor lookup succeeds; it is not a pair of independent
operations. A=8 needs no nonzero coordinate. A=0x8000 does not pass A<4.
The later linked-item arm requires not N and L and treats nonzero A as
arithmetic data: kind low8(A+43). The local store maps 212/213/65535 to 255/0/42;
this does not establish that every resulting kind is valid at later consumers.

`ALM-T9ACTOR-176` fixes the inputs. The key is full32 B|(C<<16) for both
lookups. The actor path creates actor+140 storage only if absent, constructs
an entry from low8(spellRaw) and inserts it at low16(spellRaw), deleting a
previous occupied slot through its own virtual+4. The building constructor
receives no A argument. Its copied fields and bounded lifetime are in
[Buildings and casters](../unit/structures.md). Native admission and additional
record aliases remain Unknown.

`%d.alm` is built + loaded in `R0099`→`R0473` (`CMap` family).
`R0065` reads groups (type 5, `"…no groups"`) + the drop location (type 7, `"…no
drop location in .alm…"`); `R0474` reads Outpost/repopper effects (type 7). All
via a generic named-node accessor family (`R0475` find-by-name, `R0476`
get-param, `R0477`/`IsEmpty` iterate). This historical cross-reference
does not classify types 8/9 as named-node trees: their distinct record consumers
are `ALM-SACK-065`, `ALM-EFFREC-071` and `UNIT-M10CELL-054`.
