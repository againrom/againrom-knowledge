# SAV save game (`Asg&`)

A ROM1 `.sav` contains a compressed simulation document, an uncompressed
application state store and a campaign record. The simulation document always
carries the Player/Group/actor roster; a byte selects whether terrain and the
current session follow. The grammar below describes the known original
serializers. Some raw field meanings and safe values for a new game state
remain Unknown. — SAV-FRAME-021, SAV-SHAPE-023, SAV-ROSTER-024,
SAV-FULLREAD-252, SAV-WRITERAUDIT-380

## Reference map

| Data | Reference |
|---|---|
| Compression, integer counts, strings, archive tags | [Encoding](encoding.md) |
| Simulation head, city/world branch, document trailer | [Document](document.md) |
| Every known class body and embedded collection | [Object programmes](objects.md) |
| Player fields, identity, outcome and settings | [Player](player.md) |
| Group, AI state, member order and patrol lists | [Groups](groups.md) |
| Diary arrays and owner reference | [Diary](diary.md) |
| Token head, Position and saved-address keys | [Token](token.md) |
| Unit/Humanoid/Human fields, definitions and experience | [Actors](actors.md) |
| Human live values, modifiers and producer order | [Human state](human-state.md) |
| Inventory, Item/Effect and equipment state | [Items](items.md) |
| Terrain overlays, cell payload, session and load order | [World](world.md) |
| Label, `&YA1`, shortcuts, Fog and Projectiles | [Application state](application.md) |
| Campaign arrays, mission selection, markers and tavern eligibility | [Campaign](campaign.md) |
| Emission sequence, required relations and compatibility limits | [Writing](writing.md) |

The campaign score inputs `+124/+128` are raw saved words. Their located
writers accumulate groups of 16 simulation sub-ticks and received hostile
corpse-stage transitions, respectively; their precise guards and reconstruction
limits are in the campaign reference. LOAD does not locally clamp or recompute
these counters. — FAME-021, FAME-022

## Notation

All integer fields are little-endian. `u8/u16/u32` name wire widths;
`i8/i16/i32` name a signed interpretation where established. `raw N` is an
N-byte transfer. Object offsets such as `Player+0x38` are offsets in the
original runtime object, **not file offsets**. A variable-size field moves all
following wire offsets. Runtime member offsets are hexadecimal, including
abbreviated `+38`; wire-offset plus byte-count columns are decimal unless
prefixed with `0x`. `Count`, `CString`, `objref` and `list32` are defined
in [Encoding](encoding.md). — SAV-PLAYER-028, SAV-MEMBER-036 (its
unread-classes Unknown superseded by SAV-EMBED-039), SAV-UNITPROG-156,
SAV-FULLREAD-252, SAV-758

## File envelope

```text
file offset        bytes                       contents
0x00               16                          header
0x10               blobBytes                   compressed document
0x10 + blobBytes   256                         label buffer
...                variable                    &YA1 state store
...                variable                    campaign record
...                                            physical EOF
```

| File offset | Width | Field | Rule |
|---:|---:|---|---|
| `0x00` | 4 | magic | `41 73 67 26` = `Asg&`; mismatch reports `Invalid save file.` |
| `0x04` | 4 | `blobEnd` | Writer patches the first offset after the blob; original reader reads and discards it |
| `0x08` | 4 | version | Writer emits `0x0BAD0002`; reader rejects smaller values with `Outdated save file.` |
| `0x0c` | 4 | `blobBytes` | Reader's allocation and read length; writer makes it `blobEnd - 16` |

The blob begins with its own `u32 outWords`, followed by the
[word codec](encoding.md#word-codec). That dword belongs to the blob; the
container header is exactly 16 bytes. A consistent written envelope has
`blobEnd = 16 + blobBytes`, with at least four blob bytes for `outWords`.
— SAV-HDR-001 (its 20-byte header and `0x14` body start superseded by
SAV-FRAME-021's 16-byte header), SAV-VER-002, SAV-FRAME-021, SAV-CODEC-022,
SAV-EXT-009

## Decoded document

```text
head                         clocks, map name, eleven values, mission, difficulty
Players                      list32<Player>, including Groups and actor graphs
dead actors                  list32<Unit>
worldPresent                 u8; writer 0/1, reader zero/nonzero
if worldPresent != 0:
    Buildings                list32<Building>
    SpellEffects             list32<SpellEffect>
    terrain                  block rows, cell rows, terrain identity key
    session                  4374 bytes
    Sacks                    list32<Sack>
trailer discriminator        u32; writer 0xBADFACE1
if discriminator == 0xBADFACE1:
    trailer scalar           u32
trailer state                400 bytes
word alignment               one byte only if the logical endpoint is odd
```

Player-list metadata precedes its count. See the complete
[head and ordered grammar](document.md). There are no separators between
counted records. Object programmes determine their endpoints; scanning for
class names or `0xBADFACE1` does not. — SAV-DOC-053 (partially retracted for
its 7-of-18, Unit-failure and terminal-padding clauses; the top-level order
stands), SAV-PLDIARY-054, SAV-TAGSCAN-158, SAV-FULLREAD-252

A no-world document retains the roster and trailer and omits the world half.
Its map name can still name the previous mission while its mission number is
zero. Determine the document shape from `worldPresent`. The mission outcome
is the separate saved byte `Player+0x3c`. — SAV-CITY-030, SAV-FLAG-027

## Read sequence

1. Check the envelope and consume exactly `blobBytes` bytes from file `0x10`.
2. Decode the blob as words. For structural validation, require exactly
   `outWords` output words and no truncated opcode operand.
3. Parse the document in the order above. Maintain one shared CArchive
   class/object index table and a separate saved-address map. Use the exact
   class programme at every reference; inherited and embedded bodies do not
   add tags of their own.
4. Consume the logical trailer, then classify the remaining decoded byte as
   word alignment only when the logical endpoint is odd. The original loader
   does not compare its document cursor with the decoded extent.
5. Read the 256-byte label, then the framed `&YA1` store, then the counted
   campaign record. Determine the store end from its header/records/pool;
   determine the campaign end from its grammar.
6. For runtime reconstruction, follow the [world load order](world.md#load-order):
   construct external ALM terrain before saved overlays, then perform each
   field's specific reference repair. Restore application and campaign state
   through their own consumers.

The complete structural reader validates byte 0 through EOF; the located
original application loaders make no EOF comparison after the campaign.
Structural closure does not by itself establish original load or next-action
acceptance. — SAV-FULLREAD-252, SAV-ARCHREL-253, SAV-CELLLOAD-108,
SAV-CELLLOAD-109, SAV-DECPAD-238, SAV-PHYSUFFIX-237, SAV-ORIGLOAD-332

## Write sequence

1. Build the current Player/Group/actor graph and choose the city or world
   shape. Supply established field values and external map/campaign relations;
   unresolved raw fields have no universal default.
2. Assign archive indices in first-use order. Assign distinct nonzero
   saved-address keys where the serializer defines identities, and write each
   reference as its target's key. Preserve the separate relation types.
3. Emit the exact document programme, including embedded Diary/AI/list bodies,
   the selected world half and the unconditional 400-byte trailer state.
4. Pad an odd logical document by one byte, encode its words and prefix the
   blob with `outWords`.
5. Emit the 16-byte envelope, blob, 256-byte label, `&YA1` application state
   and campaign record. Set `blobBytes`; patch `blobEnd` to the blob endpoint.
6. Validate the complete grammar and the
   [required cross-record relations](writing.md#required-relations).

[Writing](writing.md) lists the exact producer/loader differences, original
acceptance limits and remaining Unknowns. In particular, original application
restore reads four `SpellBook/Shortcuts` dwords; an empty array is insufficient.
— SAV-FRAME-021, SAV-CODEC-022, SAV-ARCHREL-253,
SAV-TRAIL-026, SAV-ORIGFAULT-335, SAV-WRITERAUDIT-380
