<a id="format-specifications-level-3"></a>

# Format and runtime references

These references describe the installed ROM1 resource families and their
promoted loader, writer and behavioral contracts. Each page keeps unknown
fields and limits explicit. Resource scope: `INV-CORPUS-001`, `INV-SIG-002`,
`INV-SCOPE-005`.

<a id="target-formats"></a>

## ROM1 files and resource contracts

| Format | Seen as | Nature | Reference and limits |
|---|---|---|---|
| **RES / LM** | `*.res`, `KIDS.LM` | Container archive, magic `26 59 41 31` | [Reference](res/format.md): Header, tree nodes, payload addressing and emission; unnamed header/node fields remain Unknown |
| **REG** | `*.reg` inside archives | Hierarchical typed store | [Reference](reg/format.md): Record/pool encoding, lookup, writing, inheritance and consuming registries; complete map.reg key table remains Unknown |
| **SPR256** | `*.256` inside archives | Palette-indexed sprite frames | [Reference](spr256/format.md): Frame/trailer layout, byte RLE and class sheet roles; named overlay and palette anomalies remain Unknown |
| **SPR16A** | `*.16a`, `*.16` inside archives | Word-RLE sprites plus byte-RLE fonts | [Reference](spr16a/format.md): Container, literal blend, font glyphs/advances and resource consumers; framebuffer globals and wider malformed-input behavior remain Unknown |
| **PAL** | `*.pal` inside archives | Palette tables and recoloring | [Reference](pal/format.md): BMP/flat layouts, shade tables, tier/owner selection and first-message state; unnamed runtime cases remain Unknown |
| **ALM** | `*.alm`/`*.ALM`, loose or in `scenario.res` | Typed map records, magic `M7R\0` | [Reference](alm/format.md): Header/version/order rules, metadata, grids, placements, script, loot and caster records; unnamed fields and full arbitrary-map authoring remain Unknown |
| **SAV** | `game*.sav` | Save game, magic `Asg&` | [Format reference](sav/format.md): complete structural read programme for the known city/world serializers; envelope, compression, object graph, write sequence and required relations. Field meanings and safe general authoring values remain partly unknown; complete reading does not establish a complete from-scratch writer (`SAV-FRAME-021`, `SAV-FULLREAD-252`, `SAV-WRITERAUDIT-380`) |
| **FAME** | `famehall.dat` | Hall-of-fame record list | [Reference](fame/format.md): Record layout, ordinary reader, insertion, display and writer; two opaque tails are carried and exceptional lifecycle remains Unknown |
| **MENU** | `main.res:graphics/mainmenu/*` and UI resources | Asset, hit-test and widget contracts | [Reference](menu/format.md): Mask indices, bitmap placement, command panels, repeated controls and frame families; native event routing and the complete frame population remain Unknown |
| **VIDEO/MUSIC** | `Allods/VIDEO*.RES`, `Allods/MUSIC.RES` | Video, music and sound resources | [Reference](video/format.md): Archive routes, resource names, selection and playback controls; decoder internals and hardware timing are outside this reference |
| **TERRAIN** | `graphics.res:terrain.3d` and ALM planes | Tile graphics, geometry and cell rules | [Reference](terrain/format.md): Tile-word lookup, 16-bpp lighting, fog, sprites, structures and passability; 8-bpp/legacy paths and named runtime boundaries remain Unknown |
| **DAT** | `world.res:data/data.bin` | Definition database | [Reference](databin/format.md): All eleven collection grammars and established consumers; selected raw-array/extra-byte meanings remain Unknown |
| **TEXT** | `main.res:text/*.txt`, `patch.res:patch.txt` | Byte string tables and UI text | [Reference](text/format.md): Splitting, display/input transforms, font/name/collision rules and localized consumers; complete Unicode and malformed-byte behavior remain Unknown |

<a id="not-a-file-format--simulation-areas-with-their-own-spec"></a>

## Runtime contracts

These areas define inputs, state and transitions rather than a standalone
file encoding. Their save fields link to the relevant stored formats.

| Area | Reference | Contents and limits |
|---|---|---|
| **MOVE** | [Reference](move/format.md) | Domains, search costs, routes, reservations, step timing and refresh; complete blocked-cell verdict table and selected cell-boundary effects remain Unknown |
| **SHOP** | [Reference](shop/format.md) | Stock generation, candidate/enchantment pools, price arithmetic, tray/payment/return and retained state; partial-selection gesture and first-shelf timing remain Unknown |
| **TAVERN** | [Reference](tavern/format.md) | Mercenary types, shelf/hire gates, price/level rules and death recovery; selected pending-hire and native failure behavior remain Unknown |
| **HERO** | [Reference](hero/format.md) | Character generation, skills/experience, ordered derived stats, equipment, combat and party continuity; named producer/lifecycle gaps remain Unknown |
| **MAGIC** | [Reference](magic/format.md) | Spell definitions, books, power, casting, effects, projectiles and action gates; named consumers and runtime limits remain Unknown |
| **ANIM** | [Reference](anim/format.md) | Drawable clocks, action phases, frame selection, numerals and composition; selected message meanings and native clock appearance remain Unknown |
| **SESSION** | [Reference](session/format.md) | Construction, phases, clocks, speed, map load, mission end, input and viewport; selected network-client and lifecycle paths remain Unknown |
| **DIALOGUE** | [Reference](dialogue/format.md) | Announcement transport, text windows, paging, tags, wrap/clamp, outcome dispatch and mission70 report identity; selected tag effects and composed names remain Unknown |
| **TRIGGER** | [Reference](trigger/format.md) | ALM compilation, register/latch state, checks/actions, comparisons, binding and mission outcomes; dormant or unreferenced arms keep their own limits |
| **UNIT** | [Reference](unit/format.md) | Definition-to-instance creation, equipment, overrides, combat inputs, Building/cell interactions and presentation; selected field/callsite meanings remain Unknown |
| **ITEM** | [Reference](item/format.md) | Class fields, effect grammar, formulas, containers, equipment, transfers, activation and sacks; unnamed columns/flags and selected lifecycle paths remain Unknown |
| **AI** | [Reference](ai/format.md) | Input commands, ownership, diplomacy, sight, target selection, group/member states, orders and retreat; broader execution and saved-continuation gaps remain Unknown |
| **MISSION** | [Reference](mission/format.md) | Placement, mission entry/exit, persistent party, campaign state and mission70 authored report predicate (`PARTY-ORIGIN-010`, its only-reader clause for the flat index partially retracted); selected sentinel, fresh-campaign gates and native report timing remain Unknown |
| **TOWN** | [Reference](town/format.md) | Exterior reactions, tavern drawing/clocks, school presentation and bounded world-map Return input; wider room/world-map behavior remains outside this page |

## ROM2

The [ROM2 family index](rom2/README.md) covers stored layouts and selected
loader/header contracts. It does not establish complete ROM2 decoding or
runtime equivalence with ROM1. The structural compiled-binary relationships
in `claims/rom2-engine.md` have no standalone stored format page.

| Format | Seen as | Reference | Contents and limits |
|---|---|---|---|
| **RES** | `Root archives` | [Reference](rom2-res/format.md) | Header/tree addressing and standard emission; payload equivalence between world/world_srv remains Unknown (`R2-ASSET-001`) |
| **ALM** | `Loose and scenario.res maps` | [Reference](rom2-alm/format.md) | Record types 0–12, version gates and declared/actual extents; new field meanings and writer accounting remain Unknown |
| **Data.bin** | `world/world_srv data and root templates.bin` | [Reference](rom2-databin/format.md) | A–H layouts, C raw14 extension and bounded templates.bin divergence; wider table semantics remain Unknown |
| **REG** | `Nested &YA1 resources` | [Reference](rom2-reg/format.md) | 24-byte header, records and pool; ROM2 kind semantics, lookup and value consumers remain Unknown |
| **Sprite / palette** | `*.16a, *.16, *.256, *.pal` | [Reference](rom2-spr/format.md) | Frame/trailer/palette shapes and named residues; full pixel/RLE/color semantics remain Unknown |
| **TEXT** | `main.res text and patch.txt` | [Reference](rom2-text/format.md) | String-table structure and known byte ranges; a complete named encoding and decoder remain Unknown |
| **Completion video** | `video.res SMK/REG and main.res UI text` | [Reference](rom2-video/format.md) | Selected report/movie source, numeric plan and return boundaries; codec internals, physical EOF and media fidelity remain Unknown |
| **Session frame** | `Socket and DirectPlay receive paths` | [Reference](rom2-net/format.md) | Eight-byte header and admission/decompression bounds; payload/opcode grammar and ROM1 equivalence remain Unknown (`R2-SESSION-003`) |
| **SAV** | `game*.sav` | [Reference](rom2-sav/format.md) | Four original save points: exact envelope/tail and bounded first Player decode; Group/actor/current-party identity Unknown (`R2-SESSION-075`, `R2-SESSION-078`) |
| **Character** | `*.a2c` | [Reference](rom2-a2c/format.md) | Six scrambled sections, checksum, plus a byte-run codec; section meanings Unknown (`R2-SESSION-022`) |

<a id="explicitly-out-of-scope-not-rom1-native"></a>

## External formats

These identified standard or third-party formats have no native-format
reference here:

- `Map Editor.opt`: Microsoft OLE2 compound document (`D0 CF 11 E0 …`).
- `*.bmp`, `*.wav`: Windows BMP and RIFF WAVE.
- GOG/system wrappers: `goggame-*`, `unins000*`, `ddraw.dll`, `smackw32.dll`,
  `aqrit.cfg`, `webcache.zip`, `Help/*.htm`, `Hints/*.gif`, `*.ico`, `*.lnk`.
- `rom.exe`, `Map Editor.exe`: PE executables.
