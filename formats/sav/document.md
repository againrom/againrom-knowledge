# SAV simulation document

[Format reference](format.md) · [Encoding](encoding.md)

## Runtime source names

Document `world` is the server singleton `[L00285]`. Terrain
`[L04624]` is a separate object. The names in the source columns below
do not identify a single shared allocation. — SAV-657

## Head

Offsets in this table are relative to decoded byte 0. Let `p` be the first
byte after the encoded map-name CString. For short ANSI `10.alm`, `p=0x0f`.
The map name is reopened as `Scenario\` plus the name when the mission number
is nonzero. Mission start obtains that number with `atoi(mapName)`.
— SAV-MAP-005, SAV-HEAD-025

| Decoded offset | Wire | Original source | Meaning or rule |
|---:|---|---|---|
| `0x00` | u32 | `world+0x04` | Sub-tick counter; zeroed at mission start |
| `0x04` | u32 | `world+0x00` | Full-clock value; restore the separately stored value |
| `0x08` | CString | `world+0x28` | Map name |
| `p+0` | 11 × u32 | See ordered list below | Literal saved values |
| `p+44` | u32 | `world+0x80` | Mission number |
| `p+48` | u32 | `world+0x84` | Difficulty; LOAD retains a stored value only in 1..3 |
| `p+52` | u32 | `playerList+0x20` | Meaning Unknown |
| `p+56` | u32 | Player list count | Followed by that many `objref<Player>` operations |

The eleven-value source order is
`+11c, +124, +128, +12c, +130, +134, +138, +13c, +148, +144, +140`.
The final three are not in address order. They are not the difficulty field
and receive no Boolean normalization or difficulty clamp.
— SAV-STREAM-010, SAV-HEAD-025, UNIT-GATE-012, SAV-657, SAV-WHEADLOAD-521

## Ordered grammar

`R0414` writes the following after the head:

| Order | Wire | Source or exact serializer |
|---:|---|---|
| 1 | Player bodies, as counted by the head | `R1558`, typed Player archive references |
| 2 | `list32<Unit>` | Dead-actor manager `R1589 -> R1341` |
| 3 | u8 world-present discriminator | Writer derives 0/1 from `world+0x2c`; loader tests zero/nonzero |
| 4, world only | `list32<Building>` | `R1590 -> R1591` |
| 5, world only | `list32<SpellEffect>` | `R1122 -> R1109` |
| 6, world only | Block array, cell table, u32 terrain key | `R1360`; [World](world.md) |
| 7, world only | raw 4374 | Session `R0270` |
| 8, world only | `list32<Sack>` | `R1592 -> R1593` |
| 9 | u32 discriminator | Writer emits `0xBADFACE1` |
| 10, marker matches | u32 | Global `[L07886]` |
| 11 | raw 400 | Block pointed to by `world+0x118`, `R1491` |
| 12, odd endpoint only | 1 byte | Outer word-compression alignment; not a document member |

The Building, SpellEffect, Sack and dead-actor managers are reached through
four fields of the separate object at `world+0x14`: `+0`, `+4`, `+8`, `+c`
respectively. The dead list is not `world+0xc`. Every top-level list uses a
plain u32 count and complete archive references; an empty list is four bytes.
— SAV-ROSTER-024, SAV-SHAPE-023, SAV-DOC-053 (partially retracted for its
7-of-18, Unit-failure and terminal-padding clauses; the top-level order
stands), SAV-657

## Load-side validation

The load arm validates almost nothing, and this is a property of the whole arm
rather than of one field. After the store/load split the arm holds 24
conditional branches and the enclosing routine has one exit instruction, so
there is no error return for a rejection to reach. Exactly one value is
substituted: a world-head difficulty outside `1..3` is discarded and the
constructor's value kept. The Player list is read with no count test and no
return test, so a count of 0 yields an empty roster and the arm continues. A
`Player+0x3c` outside its own `{0,1,2}` domain is restored literally. A trailer
discriminator other than `0xBADFACE1` skips the following global read and falls
into the same common tail. None of these is a refusal.

This covers semantic malformation — an empty count, an out-of-range enum, an
unrecognised discriminator. It does not cover a truncated stream: the archive
read primitives are outside this reading, and a throw raised inside one would
not appear as a branch in the arm. Whether the arm's permissiveness becomes a
fault downstream is Unknown; no consumer instruction that dereferences an
emptied record without a guard has been named.
— SAV-1030, SAV-HEAD-025, SAV-FULLREAD-252, SAV-ROSTER-024

## City and world shapes

Both shapes contain the entire roster, dead-list framing and trailer.
A city/no-world save omits Buildings, SpellEffects, terrain, session and Sacks
from the conditional part. It can retain the previous map name with mission
number zero. Unit/Human objects in the roster remain legal irrespective of
whether the conditional world half is present. The located no-world path
restores and detaches actors without constructing terrain. The narrowed
placement-latch claim supplies no universal relation between `Player+3d` and
the city/world discriminator.
— SAV-CITY-030, SAV-ROSTER-024, SAV-SHAPE-023, SAV-POSTLOAD-220

Two preserved city-shape streams are measured end to end. File-absolute, the
3,544-byte stream: header `0..20`, word-coded document `20..2219`, label block
`2219..2475`, embedded store `2475..3234`, uncompressed campaign record
`3234..3544` with all 310 bytes consumed across 36 fields. The 3,223-byte
stream: `20..1940`, `1940..2196`, `2196..2955`, `2955..3223`, all 268 bytes
across 27 fields. Decoded, the two streams are 6,110 and 5,592 bytes: head
`0..75`, Player list `75..5696` and `75..5178` with one Player each, dead-actor
list of count 0, then the world-present discriminator at decoded offset 5,700
and 5,182 respectively, value 0. That offset is measured, not inferred.
— SAV-1026, SAV-CITY-030, SAV-SHAPE-023

No counted collection in either city-shape stream is sized by, or indexed by,
the mission space. Every variable-length collection carries its own count: 44
of them in the 3,544-byte stream, 40 in the 3,223-byte one, each count read off
the stream. The largest count anywhere is 119, six times per save — the
`CDWordArray` at `Diary+04` and the `CWordArray` at `Diary+18` of each of the
three Diary records a city save holds. 119 is the Units table's row count and
the subscript is the actor's class-dependent Units/Humans definition ordinal,
not a mission. The next largest is 19, the Spellbook, subscripted by spell id;
then 12, the carried container; every other count is 5 or less. The campaign
record's largest is 15. Every remaining span whose byte length is not an element
count is a fixed structure span whose length is a literal inside the emitting
helper. Both streams stand at campaign mission 30. A consumer must not look for
a completed-mission bitmap or a mission-indexed list in a city save: campaign
progress is the campaign record's own `+0x04` scalar plus one byte per Player at
`+0x3c`.

What this does not settle: a short list can carry mission ids as values without
being subscripted by them. `Unit+15c`, `Unit+178`, `Group+20`,
`*(Group+3c)+4c` and `*(*(Unit+158)+90)` are counted here and none of their
element meanings is established, and the interiors of the fixed spans are not
decomposed.
— SAV-1026, SAV-CAMPAIGN-077, SAV-667, SAV-668 (per-unit-type reading
superseded by SAV-844), SAV-844, SAV-662, MAGIC-SPELL-001

The roster is nested as Player -> Groups -> actors -> owned/equipped objects.
Group records are direct inline programmes, with no class tag. The map producer
creates Players from type-5 slots, actors from type-6 records plus the hero,
and Buildings from type-4 records. These relations do not fix the file counts.
Sacks and their contents are created at runtime.
— SAV-OBJ-016, SAV-MEMBER-036 (its unread-classes Unknown superseded by
SAV-EMBED-039), SAV-DIARY-042, SAV-HUMAN-043

## Trailer

The loader always consumes a discriminator and 400 bytes. It consumes the
intervening global dword only when the discriminator is `0xBADFACE1`. Thus
marker-present and marker-absent trailers occupy 408 and 404 bytes. A marker
inside an earlier raw member does not terminate that member.
— SAV-TRAIL-026, SAV-DOC-053 (its terminal-padding clause is partially
retracted; the 400 bytes are loaded state), SAV-FULLREAD-252, SAV-790

The 400-byte block is a separately allocated, exclusively owned world object,
zero-filled by its constructor and freed at world teardown. It is not session
storage or padding. Both archive directions transfer all `0x190` bytes.

| Block offset | Width | Meaning |
|---:|---:|---|
| `0x00` | 4 | Turn-trace toggle |
| `0x04` | 4 | Script-trace toggle |
| `0x08..0x18f` | 392 | Literal stored state; consumers Unknown |

Only the first two dwords have identified consumers. The remaining fields'
meanings and pointer/alias/computed access remain Unknown. The separate global `[L07886]`
has initialization and archive read/write sites but no established nonzero
ordinary producer.
— SAV-646, SAV-647 (its zero-direct-caller clause is partially retracted; the
toggle sites stand), SAV-648, SAV-649, SAV-790, SAV-791

Command 46 parameter 80 reaches the toggle helper only after Player resolution,
`byte[cmd+4]==0`, nonzero `[L00004]` and unsigned `Player+68 > 50`.
Subcommands 3/19 toggle the first/second dword: zero becomes 1, any nonzero
becomes 0. Ordinary UI emission and production of the gate byte remain Unknown.
The helper's debug help names subcommand 19 `Script tracing on/off`.
— SAV-694, SAV-647 (zero-direct-caller clause partially retracted), SAV-791

An odd logical endpoint receives one transport byte; its value is not
constrained or consumed by the document routine. An even endpoint receives
none. Fresh SAVE computes parity again rather than preserving a source pad.
— SAV-TRAIL-026, SAV-DECPAD-238

## Eleven world-head values

The constructor explicitly zeroes all eleven dwords. Its separate memset
`[+a4,+118)` and indexed initializer ending at `+114` do not cover them.
Mode setup writes `+c/+8/+150` elsewhere. Constructor zeros are not an
all-path first-SAVE vector. — SAV-WHEADINIT-520, SAV-WHEADLIMIT-525

| Field | Identified conditional consumer |
|---|---|
| `+11c` | Nonzero restricts map-unit creation to `Player+28==0`; zero also admits the later `Mission/Players` loader with `Monsters` fallback |
| `+124` | Zero admits the map-construction `Outposts` reader |
| `+128` | Zero selects `R0044(actor)`; nonzero selects `R0308(actor,Player+34,0)`, Defend |
| `+12c/+134` | Either admits a helper on a new hero; `+134` also admits it on eligible reconstructed actors |
| `+130` | Nonzero writes new-actor word `+8c=0x40` after its earlier derive |
| `+138` | Nonzero constructs `PlasmaSword` and dispatches new-actor `vt+3c` |
| `+148/+13c` | `+148` enables the modulo-5 full-clock watchdog; `+13c` selects actor `+c0` clearing or missing-actor countdown/replacement |
| `+144/+140` | Constructor and archive transport known; other consumers Unknown |

The hero helper writes six words 100 at actor `+102..+10c`, six bytes 100 at
`+10e..+113`, then dispatches `vt+50`. Its flag tests do not run on the
existing-hero return arm. The watchdog decrements a nonzero countdown;
finding zero on entry reloads 2 and attempts replacement.
— SAV-WHEADMAP-522, SAV-WHEADHERO-523, SAV-WHEADWATCH-524

No general post-load reset or nonzero ordinary producer is established for
these eleven values. Undecoded locations, aliases, computed calls and
unexpanded callees remain open.
— SAV-WHEADWATCH-524, SAV-WHEADLIMIT-525

A bounded transition reading closes two local leaves: entry's call on
`world+30` has no object store or call, and the later `world+170` accessor only
reads its dword. Mission-return cleanup
receives the pointer loaded from `world+14`, whose constructor value is
`L05906`; its three local clears are on that manager. Ordinary SAVE passes
the current global world pointer to the outer writer, which saves and reloads
that same argument for the document serializer. These local facts do not
prove identity or field preservation across intervening calls. — SAV-1070

The twelve-body expansion leaves the return's `world+6c -> L09064 ->
R1622` alias with arguments 0/-1. The last routine is the known word-array
sizing entry; its complete zero-size write/call footprint on this receiver
remains unexpanded. The return callback and other external effects also
remain open. A separate `world+48 -> L09069 -> L09070` chain writes through
the returned pointer only on the conditional mission-overlay path. That path
is not an ordinary shipped-campaign witness, and the selected random-scatter
arm is campaign-excluded. Both preservation and intervening mutation still
fit the bounded ordinary-route evidence; it supplies no all-path first-SAVE
vector. — SAV-1071, SAV-667, TRIG-INI-012, ITEM-SPAWN-026, ITEM-SPAWN-027
