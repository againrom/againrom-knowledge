<a id="rom2-single-player-save-bsg"></a>

# ROM2 single-player save (`Bsg&`)

A ROM2 `game*.sav` has the ROM1 [`Asg&` envelope](../sav/format.md) with a
different magic word. The simulation document shares its head, Player-list
position, world-present byte and trailer with ROM1; the world half, the
Player body and the physical tail differ. This page lists the differences
only; where it says "as ROM1", the [ROM1 SAV references](../sav/format.md)
apply. Four original save points have measured envelope, head and physical-tail
coverage. Group/actor programmes and current-party identity remain Unknown.
— R2-SESSION-017, R2-SESSION-018, R2-SESSION-075, R2-SESSION-078

## Envelope

| File offset | Width | Field | Rule |
|---:|---:|---|---|
| `0x00` | 4 | magic | `42 73 67 26` = `Bsg&` (0x26677342); ROM1 is `Asg&` |
| `0x04` | 4 | `blobEnd` | Writer patches the first offset after the blob; the reader discards it, the slot list seeks to it |
| `0x08` | 4 | version | Writer emits `0x0BAD0002`; reader rejects values of `0x0BAD0001` and below |
| `0x0c` | 4 | `blobBytes` | Length of the blob |

The blob is `u32 outWords` plus word-codec packets, as ROM1
[Encoding](../sav/encoding.md#word-codec). A 256-byte label follows the blob.
The slot list reads the label at `blobEnd`. — R2-SESSION-017

## Decoded document

```text
head                  u32, u32, CString, 11 x u32, mission u32, difficulty u32
Players               u32 metadata, then list of Player (below)
dead actors           list
worldPresent          u8
if worldPresent != 0:
    list              first world list
    list              second world list
    terrain           sparse key list, sub-object, u32 identity key
    session           10,374 raw bytes
    list              third world list
    list              ROM2-only object list
u32 0xBADFACE1, u32 0xBADFACE1
trailer state         400 bytes
```

The three lists before and after terrain and session sit where ROM1 has
Buildings, SpellEffects and Sacks; their classes were not read. The ROM2
reader reads both trailer words unconditionally and compares neither. ROM1 reads
the second only when the first equals 0xBADFACE1. The ROM2 world-half load builds the map from the named scenario before
reading terrain. — R2-SESSION-018

### Terrain and session

Terrain stores one `u32` per cell whose byte in plane two is above 0x0f, over
cell indices `0x807` up to `0xedee` (exclusive). The key is
`(index << 16) | (plane-two byte << 8) | plane-one byte`. The key list is
followed by one sub-object written through its own virtual slot and a `u32`
identity key. The ROM1 cell rows do not appear. List count form and the
sub-object layout are Unknown.

Session is 4,000, 1,000, 48, 400 and 4,908 raw bytes, then `u8`, `u8`, `u32` and
three `u32`: 10,374 bytes against ROM1's 4,374. Only the first and fifth
widths differ from ROM1. — R2-SESSION-019

## Player body

The scalar prefix has the same 51 bytes and type order as the ROM1 Player,
with different runtime offsets and the same XOR key on the two obfuscated
fields. After it, ROM2 writes 2,560 raw bytes, then the group list, then 36 raw
bytes. ROM1 has the group list, 32 raw bytes and a Diary; ROM2 has no Diary.
— R2-SESSION-020

## Selected RU Group boundary

The selected RU Player list calls each Group body directly. Group first
invokes a virtual serializer on embedded +0x20, then a direct helper on its
+0x3c pointer. That helper requests 80 raw bytes and invokes a second virtual
serializer. Group then reaches its member path and three final scalar calls.
Both virtual programmes and their extents remain Unknown. The 80-byte request
is not new authority for the unread transfer helpers. — R2-ENGINE-191

The member writer calls the published u32 helper for its count, then repeats
a call with the archive and selected item pointer. Group load reads its count
and items in its own inline loop; it does not call that writer body. The
writer body's separate reader branch is not reached by this Group load path.
Reference/body helpers remain unread. Native class, alias and membership
grammar, actual hero/party identity and complete LOAD remain Unknown.
— R2-ENGINE-192

Walking only the first published Player prefix in four frozen revisions fits
group count 1 and decoded first-Group starts 2707 in A/C and 2708 in B/D.
These are conditional corpus boundaries. The first member offset and complete
Group extent remain Unknown because the embedded programmes are unresolved.
Record order and same-offset coincidence establish no current-player identity.
— R2-SESSION-083

## Physical tail

The [native campaign and selected client tail](campaign.md) gives the paired
record contracts, catalog membership restoration, current-index limits and
buffer transfer rules. Field meanings and complete live-state coverage remain
Unknown. — R2-SESSION-031, R2-SESSION-032, R2-SESSION-033, R2-SESSION-034,
R2-SESSION-035, R2-SESSION-036, R2-SESSION-037

After the label ROM2 writes the [`&YA1` store](../rom2-reg/format.md) with the
ROM1 root names (`Character`, `GameOptions`, `SpellBook`, `Objects`,
`Inventory`, `Projectiles`, `Fog`). It then writes in order:

| Block | Layout |
|---|---|
| Pair list | `u32 count`, `count` pairs of `u32`, then two `u32` |
| Scenario record | 4,096 raw, 320 raw (4 by 4 cells of 20 bytes), `u32 count`, per node `u32`, `u32`, raw 16, then `u32` current index |
| Nine blocks | each `u16`, `u16`, `u32 length`, `length` raw bytes |

No ROM1 campaign record is found in this position; no ROM1 campaign routine was
compared byte for byte, so that the three blocks replace it is Medium. Field
meanings outside the campaign bank and selected location fields below remain
Unknown. — R2-SESSION-021, R2-SESSION-023

### Scenario campaign state

The 4,096-byte bank holds 1,024 DWORD slots. At DLL NewGame return all are
zero except slot 768=10. The initial current/available record has raw type 2
and ID 1. Later client character choice writes slots 776/781 from its flag
bits 0x40/0x80; the DLL return is not the complete client initialization state.
— R2-SESSION-023

Ordinary Leave normalizes nonzero slots 532+i to 1 or 2 according to incoming
512+i, then clears 512+i for i=0..19. It sets completed slot 896+ID and clears
current. The selected mission-10 case adds record type 1, ID 20 availability and movie
output 1. The common divisible-by-ten rule advances stage slot 768 by ten.
Slot 773 is cleared; nonzero incoming slot 775 invokes an availability
restoration branch. These are conditional native transitions, not a general
linear mission sequence. — R2-SESSION-023

The [departure transition](campaign.md#departure-transition) gives the per-ID
additions, output values, bank stores and the slot 768 stage rule of the
ordinary Leave routine. Catalog identities behind the restoration branch, the
output consumer and the array after the bank remain Unknown. — R2-ENGINE-145,
R2-ENGINE-146, R2-ENGINE-148, R2-ENGINE-149, R2-SESSION-047, R2-SESSION-048,
R2-SESSION-049, R2-SESSION-050, R2-SESSION-052

The [second town entry](campaign.md#second-town-entry-and-talk) records the
constructor of catalog record kind 2 / ID 2, the EnterInn stage cases and the
shared continuation that stage 20 selects. The stage and TALK section used by
ID 2 remain Unknown. — R2-ENGINE-159,
R2-ENGINE-161, R2-SESSION-059

The selected first EnterInn caller's final local setup pushes two frame
addresses and produces no scalar catalog or stage value. The preceding EN/RU
member compare is measured, but the ABI, buffers, unread callee and alias effects,
indirect switch destinations, all-path entry, catalog/stage relation and runtime
entry stage remain Unknown. — R2-SESSION-071, R2-SESSION-072

The selected type-1 entry filter takes additional ID i+1 only when both
512+i and 532+i are nonzero. The native creation gateway uses ID and stage
to choose a template. Intervening bank writes and complete entering-party
membership remain Unknown. The selected producers below extend the bounded
construction and map-population paths.
— R2-ENGINE-076

### Selected controlled actor construction

The selected native character producer reuses owner DWORD(+0x38) when it is
nonnull. Otherwise it constructs an actor from flag-selected native template
names, assigns a separate entity key and owner pointer, registers owner actor
and group collections, then stores owner +0x38=actor. Session +0x74==0 selects
actor WORD(+0x14c)=21. The measured EN command uses opcode 0x48; indirect
delivery, template contents and the complete initial roster remain Unknown.
— R2-ENGINE-105

The selected EN location-leave arm passes nonzero scenario slot 773 and stage
slot 768 to an addition producer. Its EN/RU branch for IDs 21..30 skips new
creation when the selected owner's actor list already contains the same
WORD(+0x14c). A missing match constructs and registers a human with that word.
Slot-773 authorship, the join schedule and runtime acceptance remain Unknown.
— R2-ENGINE-106

Selected EN/RU map owner construction can reuse the existing owner for its
first row under a Session +0x74 and count-wrapper condition. A fresh owner
constructor clears +0x38. Map actors are assigned by authored owner DWORD(+8)
and grouped by an actor-row key. These conditional reuse mechanisms do not
establish complete campaign carryover or SAV restoration. — R2-ENGINE-108

## Four original save points

All four measured files satisfy `Bsg&`, version `0x0bad0002`,
`blobEnd=16+blobBytes` and exact declared codec output. Symbols identify the
four frozen inputs within this corpus and imply no filename role or ordering.
— R2-SESSION-075

| Symbol | Physical bytes | Decoded bytes | Parsed decoded prefix | Opaque decoded remainder | Player-list count |
|---|---:|---:|---:|---:|---:|
| A | 6348 | 4486 | 2707 | 1779 | 1 |
| B | 26886 | 72178 | 2708 | 69470 | 8 |
| C | 6372 | 4486 | 2707 | 1779 | 1 |
| D | 26503 | 72178 | 2708 | 69470 | 8 |

The first reference introduces schema 1 class `Player`, shared class index 1
and object index 2 under shared archive-index framing; those indices are
inferred protocol state. Its admitted prefix and raw2560 block lead to Group count 1
in all four. The inline ROM2 Group programme is unsupported, so the remaining
document is opaque. Subsequent Players, actor references, world-present flag,
inventory/equipment, stats and progression are unparsed. Player-list counts
and equal saved-address words do not establish active-party roles.
— R2-SESSION-078

A/C decoded documents compare byte-for-byte equal. B/D have equal parsed
first Player state but different first head DWORDs and opaque document bytes.
The equal name length/hash and raw2560 hash across all four establish no hero
or actor identity. Opaque differences have no assigned field meaning.
— R2-SESSION-079

## Related

The character file is a separate format: [ROM2 character file](../rom2-a2c/format.md).
