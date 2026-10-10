<a id="rom2-single-player-save-bsg"></a>

# ROM2 single-player save (`Bsg&`)

A ROM2 `game*.sav` has the ROM1 [`Asg&` envelope](../sav/format.md) with a
different magic word. The simulation document shares its head, Player-list
position, world-present byte and trailer with ROM1; the world half, the
Player body and the physical tail differ. This page lists the differences
only; where it says "as ROM1", the [ROM1 SAV references](../sav/format.md)
apply. Four original save points have measured envelope, head and physical-tail
coverage. The first Group member is walked to its payload end; later Group
bytes, other actors and current-party identity remain Unknown.
— R2-SESSION-017, R2-SESSION-018, R2-SESSION-075, R2-SESSION-078,
R2-SESSION-099

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

The body does not write the cheat flag (Player +0xa78); the constructor
sets it to 0. That a loaded Player therefore starts with cheats locked is
Medium: the load-time construction path was not read. — R2-SESSION-135

Gold is prefix field +0x3c, stored XOR 0x5c073f4d like +0xa48. The prefix
ends with Player +0x38, the controlled hero's saved address, and the
Player's own saved address. Load passes the latter to a map insert and
+0x38 to a map lookup after the groups are read; that these register and
resolve saved addresses is Medium, since the two map callees were not
read. — R2-ENGINE-323

The four frozen saves store gold at decoded offset 113 (A/C) and 114 (B/D)
as 1543978149, which unmasks to 1000. — R2-SESSION-146

## Selected RU Group boundary

The selected RU Player list calls each Group body directly. Group first
invokes embedded +0x20, then a helper on its +0x3c pointer; the helper invokes
its own +0x4c pointer before the Group member path. — R2-ENGINE-191

Under the selected construction paths, both embedded virtual +8 calls bind
to one counted two-byte-element programme. Other runtime classes/vptr
mutations and complete ordinary-save reachability remain unproved.
— R2-ENGINE-199

The count prefix c(n) is 2 bytes for n<65535, otherwise 6 bytes: u16 65535
then u32 n. With complete successful archive transfers and normal helper
returns, embedded load consumes c(n)+2*n. Store writes declared count n but
traverses L linked nodes, giving c(n)+2*L; no local n==L check is shown.
The intervening raw request is 80 bytes. Its reader can return short and the
caller ignores that return; transport/refill targets remain unread.
The reached local embedded CFG shows no explicit reference/alias transfer.
Unread append, lifecycle and transport/refill callees can still affect archive
or global state. A caller that passes no archive operand does not rule out
those effects.
These conditions prevent unconditional native LOAD authority.
— R2-ENGINE-200

Walking the published first-Player prefix gives Group starts 2707 in A/C
and 2708 in B/D. Record order and same-offset coincidence establish no
current-player identity. — R2-SESSION-083

With the selected constructor target and complete successful transfer,
both embedded counts are 0 in each frozen revision. Member count 1 then
lies at 2791 in A/C and 2792 in B/D; the first member operation starts at
2795/2796. Short-read, transport/refill and other-class alternatives remain
Unknown beside this conditional boundary. The walk stops before members;
complete Group extent, member identity and native LOAD remain unproved.
— R2-SESSION-091

The member writer calls the published u32 helper for its count, then repeats
a call with the archive and selected item pointer. Group load reads count
and items in its own inline loop; the writer body's separate reader branch
is not reached by this Group path. Reference/body helpers, native class,
alias and membership grammar, actual hero/party and complete LOAD remain
Unknown. — R2-ENGINE-192

### First Group member

Every archive object reference is a u16 tag, or 0x7fff and a u32 tag. A
clear top bit is an earlier object index, 0 for null, with no payload.
0xffff introduces a class: u16 schema, u16 name length (below 64) and the
name. Other tags name an earlier class. A new class and then the new object
take the next two indices of one shared counter, before the object's own
payload. Load matches the name against registered class records and checks
the caller's expected class. — R2-ENGINE-207

Human adds no bytes to Humanoid. Humanoid is the Unit programme, 24 raw
bytes, then twelve Item references for slots 1..12. — R2-ENGINE-208

| Unit part | Bytes |
|---|---|
| base | raw 12, u32, u16, u16, u32, u16, u32, u32 saved address, u32 (38) |
| references +0x20 | u32 count, then references |
| two word lists | count (u16, or 0xffff then u32), 2 bytes per element, each |
| raw blocks | 24, 22, 24, 64, 180, 184, then a third word list |
| scalars | 4 x u8, 3 x raw 4, 3 x u8 |
| references | +0x74, +0x78 |
| name | counted string |
| scalars | 14 x u16, 2 x u8, 2 x u16, u8, u32, 3 x u8, u32, u8 |
| packed | u32 (u16 +0x14c, two flags), u32 |
| reference | +0x68 |
| gate +0x7c | u8; if set: u32 count, references, u32, u32 |
| gate +0x140 | u8; if set: u32, u32 n, references for 1..n-1 |
| tail | 4 x u32, u8 |

The gate +0x140 list with n > 0 is Medium; the rest is High.
— R2-ENGINE-209

An Item writes the 38-byte base programme that Unit also calls first,
then u32-counted references and
u16, u16, u8, u8, u8, u16, u16, u8. Weapon adds raw 24, raw 22, u8 and a
reference; Armor adds raw 22 and u8. — R2-ENGINE-210

In all four saves the first member is a new Human (class index 3, object
index 4). Its payload runs 2806..4024 in A/C and 2807..4025 in B/D,
1218 bytes. The byte fit is High; reading it as the native LOAD result
inherits the conditional Group framing above. — R2-SESSION-099

Unit +0x74 holds a new Weapon (103 bytes). Humanoid slot 7 holds a new
Armor (77 bytes); slots 8, 9, 10 and 12 hold Armor by earlier class.
All other references are null; no earlier-object alias occurs.
— R2-SESSION-100

A/C and B/D member bytes are equal. A and B differ only in the raw 12,
raw 180, raw 184 and +0x50 raw 4 Unit blocks. — R2-SESSION-101

#### Hero fields

The producer builds the hero from a Humans row: template load copies the
row's parameters in order into Body, Reaction, Mind, Spirit (+0x84..+0x8a),
HealthMax (+0x96), ManaMax (+0x9c), speed (+0x8c), the six skills
(+0xa8+2i), typeID and face (+0xe, +0x4b), timing and token fields. The
producer then writes the chosen attributes, the name (+0x80), the owner
Player (+0x14) and +0x14c. Binding each read to its parameter title is
Medium. — R2-ENGINE-319

School skills sit at +0xa8+2i, their base at +0x116+2i, the main skill at
+0x23c, each school skill's experience at +0x23c+4i (i = 1..5) and the sum
at +0x130. — R2-ENGINE-320

Recompute derives HP and mana maxima, speed, sight (+0xa4), defence
(+0xbe) and resistances (+0xc2+2i) from the attributes, then adds the item
modifier totals held at +0xd8..+0x113. — R2-ENGINE-321

The hand weapon is reference +0x74, an Armor sits in the Humanoid slot its
own +0x58 names, the bag is gate +0x7c and the spellbook gate +0x140 is set
only when ManaMax is positive. — R2-ENGINE-322

The four frozen saves hold one Start_MF hero: attributes 41/34/24/17, Axe
20 and Shooting 10, experience 7320, HP 149/149, mana 0/0, speed 18. Face
+0x4b = 32 disagrees with the static template load, which writes 5.
— R2-SESSION-144

Hero +0x14 and Group +0x44 hold the Player's saved address; Player +0x38
holds the hero's. That load maps them back to the Player and the hero is
Medium. — R2-SESSION-145

Each of the five Armor objects carries +0x58 equal to its slot; the item
identities are Unknown. — R2-SESSION-148

Twenty-nine positions have no direct store in the selected writer
bodies, among them +0x8, +0x18, +0x78, +0x8e, +0x98, +0x9e, +0xd4..+0xd7,
the raw 12 and 184 blocks and the raw 180 block except +0xa. Unread
callees of the Unit constructor (the base constructor and the
constructors of the embedded objects at +0xa6, +0x114 and +0xd4) are
candidate writers of the base fields, +0xbb..+0xbd, +0x114..+0x12b and
+0xd4..+0xd7. The main skill +0x23c is set from the row only when the
template name contains `_Hero` or `Start_`. — R2-ENGINE-319

### After the first member

After its members the Group writes u32 +0x1c, +0x40 and +0x44 (a saved
Player address) and ends; the Player then writes 36 raw bytes and ends.
— R2-ENGINE-324

In the four frozen saves these are 0, 0 and 254852752 from 4024 (A/C) and
4025 (B/D); the first Player ends at 4072 and 4073. B/D continue with
0x8001, an earlier-class tag; what follows is Unknown. — R2-SESSION-147

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
Unknown. — R2-SESSION-021

### Scenario campaign state

The 4,096-byte bank holds 1,024 DWORD slots. At DLL NewGame return all are
zero except slot 768=10. The initial current/available record has raw type 2
and ID 1. Later client character choice writes slots 776/781 from its flag
bits 0x40/0x80; the DLL return is not the complete client initialization state.
— R2-SESSION-023 (the initial TALK clause is partially retracted; the
NewGame-return contract stands)

Ordinary Leave normalizes nonzero slots 532+i to 1 or 2 according to incoming
512+i, then clears 512+i for i=0..19. It sets completed slot 896+ID and clears
current. The selected mission-10 case adds record type 1, ID 20 availability and movie
output 1. The common divisible-by-ten rule advances stage slot 768 by ten.
Slot 773 is cleared; nonzero incoming slot 775 invokes an availability
restoration branch. These are conditional native transitions, not a general
linear mission sequence. — R2-SESSION-023 (the initial TALK clause is
partially retracted; these Leave clauses stand)

The [departure transition](campaign.md#departure-transition) gives the per-ID
additions, output values, bank stores and the slot 768 stage rule of the
ordinary Leave routine. Catalog identities behind the restoration branch, the
output consumer and the array after the bank remain Unknown. — R2-ENGINE-145,
R2-ENGINE-146, R2-ENGINE-148, R2-ENGINE-149, R2-SESSION-047, R2-SESSION-048,
R2-SESSION-049, R2-SESSION-050, R2-SESSION-052

The [inn contract](campaign.md#inn-stage-options-and-town-2-talk) separates
catalog identity, bank stage and returned options. EnterInn reads bank768;
its client ABI passes an option array and count pointer. The named fresh
ordinary departures10/20/30 yield stages20/30/40, so town2 is first admitted
at 30 and an admitted revisit after mission30 uses 40. Arbitrary bank
writers, restored state and runtime scheduling remain Unknown.
— R2-ENGINE-159, R2-ENGINE-215, R2-SESSION-107

Stage30 returns NPC22/topic30 kind3, NPC2108/topic31 kind3 only at 927=0,
and NPC2110/topic39 kind0. Stores fill the option buffer; kind3 TALK admits
catalog type1/ID-topic pointers to availability. NPC22/topic30 also
sets533=1,553=2. The selected client lists kinds0/3 as talk actors and
passes the complete word after dispatching npc%dtalk%d text.
— R2-ENGINE-216, R2-ENGINE-219, R2-ENGINE-220

The fifteen fixed continuation predicates require current-record ID and
bank values; NPC2022 families have no stage gate. The named early writer
path enables none at stages20/30, but all live bank states remain open.
A separate dynamic tail can emit kind1/ID1 immediately after Leave20.
The initial topic10 TALK also writes 769=1, correcting the partially
retracted availability-only clause of R2-SESSION-023.
— R2-ENGINE-161, R2-ENGINE-217, R2-ENGINE-221, R2-SESSION-109,
R2-SESSION-110

The [remaining inn bodies](campaign.md#remaining-inn-bodies-and-pre-50-visits)
give all 25 additional stores and their gates: stages40/50, shared60/70/80,
90,100,110. Kind3 chooses type1/ID-topic admission through TalkTo; the two
kind0 topics48/49 have no selected bank or catalog effect. Native contracts
are High, while composed presentation remains Medium.
— R2-ENGINE-223, R2-ENGINE-224, R2-ENGINE-225, R2-ENGINE-226,
R2-ENGINE-227, R2-ENGINE-228, R2-ENGINE-229, R2-ENGINE-230

With incoming DWORD stage s, Leave31/32 preserve s and Leave30 gives s+10
modulo 2^32; this local algebra is High. With no intervening stage writer,
Leave31/32 would retain 40 after the named Leave30 path; the generated
postponed-side-mission path retains 70. Postponing mission50 after Leave40
permits named town2 visits at 60/70 through Leave60/80. Leave50 unconditionally
calls add-town(3); the nonzero 775 restoration also adds town3.
The full 46-map authored35/36 population targets neither 768 nor775.
Positive composed routes are Medium; an exhaustive live stage/town set
remains Unknown because indexed writers and raw restoration are unclosed.
— R2-SESSION-111, R2-SESSION-112, R2-SESSION-113, R2-SESSION-114

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

After the character generator, a town-1 save carries scenario slots 776
(mage) and 781 (female) at bank offsets 0xC20 and 0xC34. The Player's
gold 1000 at +0x3c is Medium: the order of the setting call before the
first save was not traced. The template contents are recorded on the
data.bin page; the hero record's SAV bytes remain Unknown.
— R2-SESSION-131

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

| Symbol | Physical bytes | Decoded bytes | Published first-Player prefix | Remainder after that prefix | Player-list count |
|---|---:|---:|---:|---:|---:|
| A | 6348 | 4486 | 2707 | 1779 | 1 |
| B | 26886 | 72178 | 2708 | 69470 | 8 |
| C | 6372 | 4486 | 2707 | 1779 | 1 |
| D | 26503 | 72178 | 2708 | 69470 | 8 |

The first reference introduces schema 1 class `Player`, shared class index 1
and object index 2 under shared archive-index framing; those indices are
inferred protocol state. Its admitted prefix and raw2560 block lead to Group count 1
in all four. The selected construction/full-transfer model extends the
conditional Group walk to its first member operation, and the first member
is walked to its payload end. The rest of the Group and the remaining
document are opaque. Subsequent Players, world-present flag and progression
are unparsed; first-member field meanings are Unknown. Player-list counts
and equal saved-address words do not establish active-party roles.
— R2-SESSION-078, R2-SESSION-091, R2-SESSION-099

A/C decoded documents compare byte-for-byte equal. B/D have equal parsed
first Player state but different first head DWORDs and opaque document bytes.
The equal name length/hash and raw2560 hash across all four establish no hero
or actor identity. Opaque differences have no assigned field meaning.
— R2-SESSION-079

## Related

The character file is a separate format: [ROM2 character file](../rom2-a2c/format.md).
