<a id="rom2-format-specifications-level-3"></a>

# ROM2 format references

These pages describe the preserved ROM2 locale's known layouts and loader
contracts. Shared shapes do not import unverified ROM1 behavior. Field
semantics and writer boundaries remain explicit on each page.

| Surface | Reference | Defined boundary |
|---|---|---|
| Archives | [RES](../rom2-res/format.md) | Header, node array and payload geometry; payload identity separate |
| Maps | [ALM](../rom2-alm/format.md) | Headers, 660-byte metadata programme, version rules, type 0–12 extents and type-12 music areas |
| Definitions | [Data.bin](../rom2-databin/format.md) | Eight groups, 14-byte C block, client/server counts and parameter schema |
| Inline registry | [REG](../rom2-reg/format.md) | Header, record array and pool; value and lookup meanings Unknown |
| Sprites and palettes | [Sprite/palette](../rom2-spr/format.md) | Frame container and palette shapes; pixel decoding Unknown |
| Text | [Text](../rom2-text/format.md) | Resource families and stored-byte domain; conversion Unknown |
| Completion presentation | [Report/movie](../rom2-video/format.md) | Selected sources, numeric plans and client return boundaries; codec/physical EOF Unknown |
| Music | [Music](../rom2-music/format.md) | Archive keys and wave format, screen keys, mission areas, player and settings; audible output Unknown |
| Town presentation | [Town](../rom2-town/format.md) | ID/resource selection, Kaarg square layers, mask, clocks and room art; full room layout and shell text destinations Unknown |
| Character generator | [Generator](../rom2-chargen/format.md) | Pre-create and detail screens, budget, skills, tips and the hero at town 1; item effects and the hero's SAV bytes Unknown |
| Session records | [Header](../rom2-net/format.md) | Eight-byte header and transport length rules; payload grammar Unknown |
| Single-player save | [Save](../rom2-sav/format.md) | `Bsg&` envelope, document, Player and tail differences from ROM1; field meanings Unknown |
| Character file | [Character](../rom2-a2c/format.md) | `.a2c` sections, scramble, checksum and codecs; section meanings Unknown |

ROM2 game implementation and protocol equivalence are outside these references.
The structural binary relationship recorded in `claims/rom2-engine.md` is not
a stored format and has no independent format page.
