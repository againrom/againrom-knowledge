# SAV Player

[Format reference](format.md) · [Groups](groups.md) · [Diary](diary.md)

## Wire fields

Player begins with a CString name from `+18`. In the table, `s` is the first
byte after that encoded string. The fixed scalar prefix is 51 bytes.
Offsets in the source column are runtime Player offsets.
— SAV-PLAYER-028

| Wire offset | Type | Source | Meaning or transform |
|---:|---|---|---|
| `s+0` | u16 | `+04` | Registered slot ID; LOAD copies it literally |
| `s+2` | u32 | `+08` | Separate stored identifier; trigger lookup key |
| `s+6` | raw 8 | `+10` | Meaning Unknown |
| `s+14` | u8 | `+44` | Stored colour/shade selector; constructor 0, ALM-copy rule below, literal LOAD |
| `s+15` | u32 | `+28` | Zero for a human participant |
| `s+19` | u16 | `+2c` | Participant publication mask |
| `s+21` | u32 | `+38` | Money XOR `0x5c073f4d` |
| `s+25` | u8 | `+3c` | Mission-outcome latch |
| `s+26` | u8 | `+3d` | Retained-actor mission-entry placement latch |
| `s+27` | u32 | `+48` | XOR `0x5c073f4d`; gameplay meaning Unknown |
| `s+31` | u32 | `+50` | Constructor zero; other meaning/producer Unknown |
| `s+35` | u16 | `+54` | In-memory dword, `min(v,0x7fff)` on SAVE |
| `s+37` | u16 | `+4c` | In-memory dword, `min(v,0x7fff)` on SAVE |
| `s+39` | u32 | `+58` | AI mana-floor percentage; default 95 |
| `s+43` | u32 | `+34` | Participant's starting-character saved-address reference |
| `s+47` | u32 | `this` | This Player's saved-address identity key |

The XOR applies in both directions. The two saturated fields do not preserve
values above 32767. These transforms belong to the named Player fields; do not
apply them to other classes' fields. A value already transformed in memory
still serializes in that representation. — SAV-PLAYER-028, SAV-OBF-029,
SAV-OBFCEN-038

After the prefix, emit **all three** suffix members in order:

| Order | Wire | Source |
|---:|---|---|
| 1 | u32 group count; that many direct Group bodies | Player `+24`, `R1340 -> R1143` |
| 2 | raw 32 | Allocation pointed to by Player `+30`, `R0204` |
| 3 | Direct Diary body | Object pointed to by Player `+40`, `L08081` |

The Diary is required even when its arrays are empty. Omitting it desynchronizes
the next counted record. — SAV-MEMBER-036 (its unread-classes Unknown
superseded by SAV-EMBED-039, which reads the Diary's two members as a
`CDWordArray` and a `CWordArray`), SAV-PLDIARY-054, SAV-662, SAV-663

## Inventory gold scalar source

The measured Drop Gold path changes a client Player purse and queues a command
before server admission. SAVE's named Player body instead reads current server
`+38`; LOAD restores the same field later used by gold debit/credit. Its separate
Sack body writes/reads current `+3c`. Native SAVE/LOAD/pickup continuation and
SAVE-before-queue-execution remain Unknown (MENU-095).

## Identity and ownership

The archive reference chooses the Player object to deserialize. LOAD registers
its saved identity against that object and copies settings into its own `+30`
allocation. Stored `+04`, stored `+08` and traversal position do not redirect
this copy. Constructor `+04=0`, new-roster registration and LOAD are distinct
producers; the located LOAD/list bodies do not invoke registration.
— SAV-PLAYERIDENT-830, SAV-919

On a plain single-player mission LOAD the Player/client list is populated by
this pure archive copy-in, with no selection step inside the traced LOAD
bodies. The one documented mechanism that does select an active Player by name
match (an existing-Player search retaining the first accepted name) scopes
itself to the return-to-game/client-description route, not this LOAD route.
No claim traces a call from the mission-command dispatcher — the routine
issuing the first post-LOAD command — to either that mechanism or an
equivalent specific to this route; whether "the" restored player is the first
or only list entry by convention, rather than through any search, is untested.
— SAV-1105

Registration chooses a slot and sets `+2c = 1 << (slot mod 16)` only when
`+28==0`, otherwise zero. The ALM-to-session path separately assigns the colour
byte as `lowByte(CPlayer+08+1)`, from the ALM type-5 colour word at wire `+00`.
The Player constructor instead initializes `+44=0`; LOAD restores the saved
byte without running the ALM-copy formula. Colour, slot ID and mask remain
separate fields with separate producers. ALM-GRP-041 is partially retracted
for its inherited universal owner interpretation; its colour-field mapping
is retained here.
— SAV-664, SAV-PLAYER-028, ALM-GRP-041, PAL-SHADE-012, PAL-SHADE-013,
ALM-PLAYER-069

`+34` selects the Player's own starting character in its Group graph.
Character generation and carry assign it; mercenary spawn does not. It is not
a Group field or a rule that every Human is a hero. Actor order is not identity.
— SAV-HERO-059, SAV-HEROID-065, SAV-GRPORD-058

Actor `Token+14` identifies its owning Player on the owner path;
LOAD's Player suffix also rebuilds membership from Groups. Discordant Group
`+44` can independently resolve to another Player or null. There is no proven
fallback from that missing owner to containment, or universal active-client
selection by first record/identifier 1. — SAV-OWNER-048, SAV-GRPLOAD-560,
SAV-GRPOWNER-561, SAV-PLDIARY-054

## Mission outcome and entry latch

| `+3c` | Meaning | Producer |
|---:|---|---|
| 0 | Mission in progress | Constructor, mission-start/reset paths |
| 1 | Mission complete | `session+b3ac==1`; log `Logic - Mission Complete`, announcement `0xb5` |
| 2 | Mission failed | No living hero or `session+b3b4==1`; log `Logic - Mission Failed`, announcement `0xb4` |

The full-tick reporter tests defeat before victory and acts only on
`+28==0` with nonzero `+3d`. `session+b3ac` is a world-half win counter;
`Player+3c` is the outcome retained in both document shapes.
— SAV-FLAG-027

The field's own wire instructions: the store reads the byte into `DL`, pushes
it and calls the archive's one-byte write primitive; the load pushes the field's
address and calls the one-byte read. `+3d` follows immediately through the same
pair. Neither arm tests the value, so a byte outside `{0,1,2}` survives a round
trip unchanged. The other progress-bearing field of a save, the campaign
record's `+0x04` mission scalar, belongs to a different serializer in the
uncompressed tail; the two never meet.
— SAV-1027, SAV-FLAG-027, SAV-1030

At session join, zero packet `+0a` and zero saved `+3d` admit placement of
unseated retained actors. Placement replaces Position cell, packed cell,
sub-cell and terrain, with possible further clamping. Seated actors are
skipped. Both placement and no-placement arms then set `+3d=1`; mission
teardown clears it on the surviving human Player. The narrowed contract governs
actor placement; `+3d` alone supplies no universal city/world shape rule.
— SAV-POSTLOAD-220

## Settings and consumers

The 32-byte settings allocation is copied raw. Its constructor clears it and
sets byte `+1f=2`, the formation mode. The other bytes are not specified as
padding or safe replacements for stored state.
— AI-FORM-037 (its corpus clause is retracted; the constructor byte stated here stands), SAV-662,
SAV-663

| Consumer | Player relation |
|---|---|
| Group movement formation | Group `+44` -> Player `+30` -> byte `+1f` |
| Client formation command | Active client CPlayer `+04` sent as a word; first simulation Player with matching `+04` |
| Trigger formation write | Temporary map keyed by Player `+08`; compiled record `+38` holds the resolved pointer |

Client CPlayer and serialized simulation Player are separate objects.
The selected return path sends the incoming-name match first, then the others;
the client sets its active pointer when the current array size equals 1.
Incoming-name provenance on every LOAD route, initial client array state,
packet delivery and normalization of discordant relations remain Unknown.
— AI-FORMOWNER-314, AI-FORMCMD-315, AI-FORMTRIGGER-316, AI-FORMACTIVE-317

Command `0x46`, parameter 3 produces `+58`: signed inputs 0/1/2 map to
100/50/0; 3..100 are copied; negative or above 100 preserves the old value.
It then dispatches every actor's derive. Unit derive consumes the percentage
through actor `+14`. SAVE/LOAD does not apply this command mapping, and
parameter 1 is the separate withdraw arm. — SAV-726, SAV-666,
UNIT-DERIVE-003, HERO-MP-006

The `+50` constructor writes zero; other ordinary producers remain Unknown.
LOAD still restores the saved value.
— SAV-665

## Diary client projection

The restored Player Diary and the client's received word cache are separate
objects. Opcode186 replaces the CWordArray in the supplied client dispatch
receiver; its arm performs no second Player lookup. The frontend stores
that client object at `+d0`. — SAV-995

The conditional opcode4 resume setup uses its matched Player for the Diary
builder when Player `+34` is nonnull and preceding calls return normally.
The call is outside the optional initialization arm. Later attributed events
use the existing Diary notification gate; the additional text-command call
does not prove another ordinary refresh trigger. Native delivery and first
refresh order remain Unknown. — SAV-997

## Tavern command prerequisite

The reached command37 arm selects a Player by signed word `+04` and returns
without stock construction when no match exists. On a match, the stock
handler first adds the tavern refund to Player `+38` and requests event67,
even for a zero refund. Earlier UI/transport and actual client-selected
identity remain Unknown. — SAV-918

The first failing tavern instruction and actual client-selected Player state
remain Unknown. Matching stored `+04/+08/+28` values alone supplies no complete
native tavern-construction contract. The campaign's
[eligibility and selection](campaign.md#tavern-eligibility-and-selection)
consumers impose separate conditions. — SAV-920
