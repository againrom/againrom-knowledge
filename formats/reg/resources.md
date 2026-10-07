# Campaign, sound and other registries

[Reference](format.md)

## The other registries (`REG-CUT-053`, `REG-SFX-057`, `REG-NPC-058`, `REG-SCN-059`, `REG-AI-060`, `REG-INNCLOSE-116`, `REG-INNSTAGE-117`)

`data/map.reg` has no complete published per-key table; the named consumers
below define only their own keys.

- **`sfx.res::sfx.reg`** — the **sound-slot table** (`REG-SFX-057`). `[Global] SfxCount = 564`
  is the **highest slot id, not the entry count**: `[Sfx]` holds **115** kind-0 strings named
  `Sfx<n>` over a sparse id space whose maximum is exactly 564. Each value is a path relative
  to `sfx.res`, no extension (`click00`, `units\sword`, `ambient\river`, `magic\firewall`).
  Installed nonzero units.reg Sound[] values identify defined slots.
  Ambient slots (`REG-OBJ-047`) include
  `0x3c..0x3e` = `ambient\bird1/2/3`, `0x46` = `ambient\crow`.
  The loader and consumers are pinned by `VIDEO-SFX-013`…`VIDEO-SFX-021`.
  `R0585` allocates `SfxCount+1` pointers,
  clears them and constructs present keys `Sfx1` through `Sfx<SfxCount>`; the data value is
  capacity, not a live-event count. The installed archive lacks three named paths: slot 62 `ambient\bird3`, slot 80 `ambient\wind1` and slot 81 `ambient\wind2`.
  Every non-zero inherited `units.reg Sound[]` value still resolves. Runtime selectors are
  separate: fixed slots, inherited elements and `500+value` formulas read this pointer array,
  while literal-path samples and conditional voice banks bypass it (`VIDEO-SFX-013`..`021`).
- **`scenario.res::npc.reg`** (`REG-NPC-058`) — **105** `npc<n>` sections (ids 1..132, sparse),
  four archetype blocks (`MaleFighter`, `MaleMage`, `FemaleFighter`, `FemaleMage`, each
  `Face Body Reaction Mind Spirit Skill`), and `Multiplayer` with four kind-6 face lists
  (`FacesMM` 6, `FacesMF` 15, `FacesFM` 4, `FacesFF` 8). Per-npc keys are Flags, Face, Picture, PortraitX1/Y1/X2/Y2,
  DataBinID, PriceA and PriceB. Installed ranges are Face 1..30, Picture 64..80,
  PortraitX1=7..69, Y1=8..72, X2=48..54, Y2=15..32, PriceA=28..80 and
  PriceB=0..15; none is a field-width limit. Flags is a comma-separated list
  with `!` negation over Human, Mage, Female, Face, Picture, Platoon, Hero,
  Me, Start, MySex and MyClass. DataBinID identifies a Templates.ini slot
  (declared 1..1088), for example 509=M10_Witch and 1051=F_Knight3. The
  command producer writes it to message+0xa behind tag 0x49.
  — EDITOR-023, REG-NPC-090

  The in-memory record is **`0x30`** bytes, allocated only when `Flags` is a non-empty string,
  and held in an array indexed by the section number itself (`npc<i>` → element *i*, 256 slots):

  | Offset | Key | Default |
  |-------:|-----|--------:|
  | `+0x08` | `Flags` (CString) | `""` |
  | `+0x0c` | `Face` | 0 |
  | `+0x10` | `Picture` | 0 |
  | `+0x14` | `DataBinID` | 0 |
  | `+0x18` | `PriceA` | 0 |
  | `+0x1c` | `PriceB` | 0 |
  | `+0x20` | `RECT { PortraitX1, PortraitY1, PortraitX2, PortraitY2 }` | -1 each |

  **`Face` and `Picture` name no picture resource** (`REG-NPC-088`). They are copied into a
  dialogue actor's **typeID** (`+0x20`) and **face** (`+0x24`), each gated on the matching token
  of the same section's `Flags` list, and the portrait that follows is the ordinary
  `graphics\infowindow\<InfoPicture><face>.bmp` leaf (`UNIT-PICT-036`) with the digit dropped at
  face 1. With the `Start` token the face comes instead from one of the four archetype sections,
  chosen by the player's own sex/class bits. Where `Picture` is present `Face` is 1..4, a tier;
  where absent it is 1..30 and the section takes the composed-doll path (`UNIT-PICT-035`).

  **`Portrait*` is a Win32 `RECT`** — the loader hands the four consecutive values to `CopyRect`,
  which fixes the order as `left, top, right, bottom` — and only `left` and `top` are read
  (`REG-NPC-089`). They cut a fixed **72 x 96** source window, `(X1, 144-Y1)`…`(X1+72, 240-Y1)`,
  out of a 160 x 240 16-bit canvas; `PortraitY1` absent (the loader's `-1`) selects a fixed
  72 x 92 default at `(36, 140)`. `PortraitX2`/`PortraitY2` have no reader in the image.

  `npc21`..`npc24` are a four-slot block the image knows by literal number — two routines gate on
  `21 <= n <= 24` and a third subscripts `array + id*4 + 0x54` — and they are exactly the four
  sections carrying `Start` (`REG-NPC-090`).
- **`scenario.res::scenario.reg`** (`REG-SCN-059`) — `[General] TotalMissions = 15`,
  `ScenarioMissionCount = 15`, `MercenaryCount` a 15-element array; **24** `[Mission<n>]`
  sections (15 main, numbered by tens, plus 9 sub-missions). Keys include Mercenaries, InnNPC, InnMission, ShopMinPrice,
  ShopMaxPrice, Payment, EnableMercenary, ShopMission, TCMission,
  AddTextDocument, AddHero, AddPictureDocument, AutoGetMission and LastMission.
  Installed ranges are ShopMinPrice 0..5000, ShopMaxPrice 1000..10000000 and
  Payment 700..700000; these are data values, not limits. NPC references name
  npc.reg sections. MercenaryCount is a 15-element pool indexed by mercenary
  type, 1..15, which is also the `[npc<t>]` section number; it is neither per
  mission nor a roster length (`REG-SCN-067`). TotalMissions has no named
  literal consumer.

  The reached shop and training show tails consume the head mission word only
  when their campaign offer-array count is positive. Shop header/data/count
  are `+630/+634/+638`; training uses `+644/+648/+64c`. They name
  `shop\npc31m<word>` or `training\npc34m<word>`, call the shared dialogue
  builder, then register and remove the head after normal return without a
  result test. This local consumer order does not establish native array
  population or exception handling. — DIALOGUE-048

  **Cross-root closure of `InnNPC`/`InnMission` against the shipped inn text, and per-stage
  divergence across a third comparison population, over four roots — the live/EN/RU triple this
  page's other rows already cover plus a fourth, owner-supplied pre-release data snapshot**
  (`REG-INNCLOSE-116`, `REG-INNSTAGE-117`): the 24-section structure and the 14-key vocabulary above are identical
  on all four; EN's registry addressing is closed against the shipped `text/inn/npc/*.txt` set in
  both directions, RU's leaves 7 shipped files unaddressed, and the pre-release root addresses 20
  files it does not ship. The pre-release root's per-stage arrays disagree with EN/RU at 5 of 13
  non-empty stages (Mission30/40/50/60/90); which other keys move disagrees per stage — Mission30
  moves `ShopMaxPrice` alone, Mission60 moves `ShopMaxPrice` but not `Mercenaries`, Mission90 moves
  `Mercenaries` but not `ShopMaxPrice` — no fixed subset of keys moves together at every
  disagreeing stage. `InnNPC`/`InnMission` are unequal length at exactly one of 39 stage×root pairs
  measured (Mission40, pre-release root).

  **The tavern reader `R0702` is bounded by the inn view's mercenary count, not by
  `InnMission`** (`REG-118`, bound corrected by `TOWN-468`): it is entered only when the
  selection index is at or past the mercenary count, reads `InnMission` and `InnNPC` at
  `idx = curSel − mercCount`, and no instruction in either compiled build compares that index
  against `InnMission`'s own size. `REG-118`'s reading of that bound as `InnNPC.Size()`, and of the
  mercenary array as `InnNPC`'s data, is partially retracted: the view and the campaign document
  are two objects. Below the bound the selection reaches the mercenary roster `MERC-HIRE-003`/
  `MERC-TYPE-001` describe, so the selectable range is mercenary cells first, then one per-NPC
  cell per `InnNPC` element (`TOWN-468`). Whether real input can drive the selection index far
  enough to expose the pre-release root's shorter `InnMission` at Mission40 is not established.

- **`scenario.res::globalmap.reg`** (`REG-SCN-059`) — `[General] ObjectCount = 35` matching
  **35** `[MapObject<n>]` sections (1-based), each `MapPoint` = 2 ints, `MapRect` = 4 ints, 5
  with a `Picture` string; plus `[MissionObjects]`, **28** `Mission<n>` keys naming a map-object
  number. Each installed key names a MapObject. All scenario.reg missions appear here;
  additional mission keys are 81,101,141,151. The two geometry keys have
  different consumers (`TOWN-116`, `TOWN-118`): `MapRect` becomes an indexed-query hit rectangle;
  `MapPoint` is a world-route endpoint and marker anchor. All 35 shipped MapPoints land on
  index-2 graph nodes in `graphics.res::global.map/pathmap.bmp`; that bitmap has 46 further
  nodes with no MapObject. A MapObject whose `Picture` is `"nothing"` receives a static return-enable
  bit, while a picture-bearing object's indexed return is gated by the runtime marker cache (`TOWN-123`).
  `MapPoint` coordinates and PathMap pixels are therefore one coupled data seam: moving one
  without placing an index-2 node at the same coordinate leaves route lookup without an endpoint.
- **`world.res::data/ai.reg`** (`REG-AI-060`) — 4 records, empty pool: `[Scanning]
  MinimalGuardRange = 8` and `[Tasker] IntelligentCons = 15`. Neither `Tasker` nor any run
  beginning `IntelligentCons` occurs in `rom.exe` or `Map Editor.exe`.
- **The 33 video-cutscene registries** (`REG-CUT-053`) — **one schema with two optional section
  families**, in three shipped shapes:

  ```
  Common { startx, starty, nFadings [, nPanaramings] }
  Fading<n>     { startframe, endframe, startfade, endfade }   (kind 4 doubles)
  Panaraming<n> { startframe, endframe, stepx, stepy }
  ```

  nFadings/nPanaramings give the counts of the corresponding sections.
  Common-only resources set nFadings=0. Installed startx/starty are 0;
  startframe<=endframe. Fades use endpoints 0.0/1.0 (kind 4 binary64), first
  fading in and last fading out. A registry accompanies the .smk resource
  with the same basename; VIDEO4/VIDEO8 can share registry values while
  containing differently sized videos. The complete meaning of frame indices
  relative to the video remains Unknown. — REG-CUT-053, REG-DBL-035

## Map Editor text catalogs (`EDITOR-023`)

The install ships parallel game-authored vocabulary as plain text at its root:
`Description Checks.ini` (22 trigger **condition** opcodes, id→name→params),
`Description Instants.ini` (44 trigger **action** opcodes), `Templates.ini`
(1088-slot placeable-template id→name table). Trigger types: `Target_Unit`,
`Target_Group`, `Target_Player`, `Target_Item`, `Target_Structure`,
`Target_Building`, `X`, `Y`, `int`, `Enum`, `Const`.

## `AddHero` town activation

A mission record stores `AddHero` as a `u16` array. Town activation walks the current record's array, sends one command `0x49` per value, and then clears the array. The array is not a mission-completion action.

The shipped value `[Mission30] AddHero = 22` selects `npc22`. Its `Mage,!MySex` flags select Humans row 28 (`PC_Fergard`) when the existing primary character is female and row 29 (`PC_Reniesta`) when the existing primary character is male. The displayed name is the corresponding localized entry in `text/npcnames.txt`. The special `npc21..24` branch replaces the stored `DataBinID` with `26 + selector`.
