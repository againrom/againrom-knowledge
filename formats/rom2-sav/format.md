<a id="rom2-single-player-save-bsg"></a>

# ROM2 single-player save (`Bsg&`)

A ROM2 `game*.sav` has the ROM1 [`Asg&` envelope](../sav/format.md) with a
different magic word. The simulation document shares its head, Player-list
position, world-present byte and trailer with ROM1; the world half, the
Player body and the physical tail differ. This page lists the differences
only; where it says "as ROM1", the [ROM1 SAV references](../sav/format.md)
apply. Nine frozen EN saves now have complete structural document and physical
partitions under the selected successful-transfer grammar. Four omit the world
half; five include it. Numeric controlled/owner-key matches find one
Player-1 Group Human per file; which Human is the active party member stays
Unknown. Native LOAD completion and first-frame party state remain Unknown.
The earlier four-point measurements below retain their
own population and grades.
— R2-SESSION-017, R2-SESSION-018, R2-SESSION-075, R2-SESSION-078,
R2-SESSION-099

The [nine-save LOAD contract](#nine-save-load-contract) gives the complete
partitions, Human fields, book populations and remaining restoration Unknowns.
— R2-SESSION-233, R2-SESSION-234, R2-SESSION-235, R2-SESSION-236,
R2-SESSION-237, R2-SESSION-238

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

The earlier positional identification of Buildings, SpellEffects and Sacks
was Medium. Selected native programmes and five complete mission-buffer walks
now give Building/Pointer and SpellEffect-derived instances; observed Sack
and extra-list counts are zero. — R2-SESSION-235, R2-ENGINE-603

The ROM2
reader reads both trailer words unconditionally and compares neither. ROM1 reads
the second only when the first equals 0xBADFACE1. The ROM2 world-half load builds the map from the named scenario before
reading terrain. — R2-SESSION-018

### Terrain and session

Terrain stores one `u32` per cell whose byte in plane two is above 0x0f, over
cell indices `0x807` up to `0xedee` (exclusive). The key is
`(index << 16) | (plane-two byte << 8) | plane-one byte`. The key list is
followed by one sub-object written through its own virtual slot and a `u32`
identity key. The ROM1 cell rows do not appear. The key count is u16, or
0xffff then u32, followed by four bytes per key. Under the selected
constructor, the sub-object has the same escaped count followed by u16 key
and raw 60 bytes per entry. — R2-ENGINE-603

A raw60 entry, at stored entry offsets (field names from read uses are
Medium):

| Bytes | Field |
|---|---|
| +0x00, +0x01 | plane bytes copied from the cell at create and written back by refresh |
| +0x02 | layer count |
| +0x03 | no access found |
| +0x04, +0x08 | occupant of kind 1..2, occupant of kind 3 |
| +0x0c | building |
| +0x10 | object |
| +0x14..+0x2b | six effect-layer pointers |
| +0x2c, +0x30 | footprint occupants, kind other than 3 and kind 3 |
| +0x34..+0x39 | cell event: +0x34 kind; +0x35..+0x39 arguments passed to `L2.03574` / `L2.03575`; for kind 0x17, +0x38 and +0x39 are read as a destination cell |
| +0x3a | the entry's own cell key, u16 |

In the five EN mission-10 saves every +0x3a equals the entry key, bytes
+0x34..+0x39 are zero, and every nonzero DWORD at +0x04, +0x08, +0x0c, +0x2c
and +0x30 is a saved self key of the same save. — R2-SESSION-266

The footprint pass keeps +0x2c and +0x30; a nonzero event kind acts when an
occupant is placed on the cell (kind 0x17 stores the target pair into the
unit's command word) or arrives there; refresh reads +0x3a as the cell to
write back. — R2-ENGINE-673

Session is 4,000, 1,000, 48, 400 and 4,908 raw bytes, then `u8`, `u8`, `u32` and
three `u32`: 10,374 bytes against ROM1's 4,374. Only the first and fifth
widths differ from ROM1. — R2-SESSION-019

The first 4,000 bytes are 1,000 DWORD registers, register n at bytes
4n..4n+3. Session init sets register 901 (bytes 3604..3607) to 1 and
register 999 (bytes 3996..3999) to 0; one routine adds 1 to 999 before it
runs the triggers. No other reader of either was found in the scanned code
and in the 209 + 48 + 197 map operands of check 19 and instants 3 and 8; the
six register compares of the trigger runner `L2.00517` were not censused.
The five EN
mission-10 saves hold 901 = 1 and 999 = 1, 1, 17, 141 and 0.
— R2-SESSION-265

On LOAD the terrain post-load pass replaces each nonzero pointer DWORD by
its saved-key lookup and keeps the saved number on a miss; bytes
+0x00..+0x03 and +0x34..+0x3b, the session block, the Player raw blocks and
the Building base word +0x0c pass without a value test. — R2-ENGINE-676

## Player body

The scalar prefix has the same 51 bytes and type order as the ROM1 Player,
with different runtime offsets and the same XOR key on the two obfuscated
fields. After it, ROM2 writes 2,560 raw bytes, then the group list, then 36 raw
bytes. ROM1 has the group list, 32 raw bytes and a Diary; ROM2 has no Diary.
— R2-SESSION-020

The 2,560 raw bytes are Player +0x44; the only reader found among the seven
EN `push 0xa00` sites passes them to a 0xba message builder. They are zero
in the five EN mission-10 saves except one byte of Player 0 in game0006. A
census of loaded +0x1c0 pointers in the terrain routines hits 27 offsets
read (heuristic; the base is not shown to be a Unit); LOAD remaps the
DWORD +0x7c of that block by saved key.
— R2-ENGINE-675

Pointer +0x6c is 24 in the five EN saves. No reader of a Building's base
byte +0x0d and no constructor store to its word +0x0c was found in the read
EN bodies (terrain module listing, serializers and constructor chain);
seven of 16 structures hold a nonzero high byte in four of the
saves. — R2-ENGINE-674

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

### Automatic-casting preferences

The selected EN Player call passes its `+0x34` pointer to preferences
serializer `L2.01536`, which transfers 36 raw bytes. This interval includes
the support preference byte `+0x20` read through the owning Player in the
selected RU support policy. The local transfer width is High; ordinary SAV
reach and the autocast field join are Medium. Mission and town preservation,
per-Player initialization and later overrides remain Unknown.
— R2-SESSION-150

The retained High local-width and Medium field-join grades use first support
before the 160-window limit. Acquisition later reached a conservative minimum
of 169 windows; the completed late comparison is excluded from those grades.
— R2-SESSION-150

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

A hero taken from a character record or through the reuse path derives each
school skill from its experience E through the threshold table
T(L) = trunc((1.1^L - 1) x 1000), and the loaders and the reuse path
rebuild the sum from E; at template load E = T(skill). The level from E is
the largest L with T(L) <= E for E below the dword past the table, which was
not read. — R2-ENGINE-485 (High for the law, Medium for the exact values and
for the set of readers), R2-ENGINE-489 (High for the arithmetic of the
bodies read, Medium for the roles)

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

#### Hero inventory

The bag at gate +0x7c is an item list: u32 count, Item references head to
tail, u32 insert index (+0x1c, 0x7fff appends) and u32 weight sum (+0x20).
LOAD empties the existing list and appends the references in saved order
(Medium: the clear and append helpers were not read); +0x1c and +0x20 are
read back, not recomputed. Worn items are the
references +0x74, +0x78 and slots 1..12, never bag entries. — R2-SESSION-165,
R2-ENGINE-402

| Record | Bytes in order |
|---|---|
| Item | base 38 (u32 +0x1c price, u32 +0x14 owner); u32 count and Effect references; u16 ID +0x40, u16 quantity +0x42, u8 type +0x44, u8 class +0x45, u8 material +0x46, u16 +0x48, u16 weight +0x4a, u8 +0x47 |
| Effect | base 38; u8 tag +0x3c, u8 flags +0x3d, u32 +0x40 (two u16 values), u16 +0xc |

An Item without effects is 54 bytes after its reference tag; each Effect
adds its reference tag and 46 bytes. Item +0x3c, +0x4c, +0x4d and +0x50
are not stored; LOAD rebuilds +0x3c from +0xc. The order and widths are
High; constructor values for the unstored fields are Medium; no save with
a nonempty bag was read. — R2-SESSION-165, R2-ENGINE-403

Unit initialization builds the bag. An empty bag stores count 0, +0x1c
32767 and +0x20 0, the constructor values; the four frozen saves hold
exactly this for their one hero. No first-Player Group member of the four
frozen saves holds a nonempty bag; the bag of S01 was not reached and
other Players were not walked (138 files examined by magic). In the
bodies read, unit death, a character file load and two script instants
replace the bag with a new empty list; respawn and the victory transition
create a bag only when the pointer is null.
— R2-ENGINE-410, R2-SESSION-169

EN and RU write the same bytes at Medium: the EN serializer bodies were
read, their list and reference callees were not. RU items carry one more runtime dword
(+0x54, not stored), so RU Weapon and Armor fields sit 4 bytes above EN's.
— R2-ENGINE-407

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

Bank slot 770 is at bytes 0xC08..0xC0B of the Scenario record's first
4096-byte bank, not at an absolute SAV-file offset. Selected ScenarioLoad
reads its saved raw bytes before catalog reconstruction. Native cold-LOAD
delivery and later client writes remain Unknown (`R2-SESSION-154`,
High for transfer and offset).

Town-3 availability records store u32 ID 3, u32 kind 2, then raw 16 bytes.
The final current index is a list ordinal with default zero, not a fixed
town-3 value. Reconstruction uses catalog kind/ID matching; complete
membership and accepted visits remain open (`R2-SESSION-155`, High for
identity and representation, Medium for full membership).

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
The mission-ID40 departure arm supplies slot773=23; the selected consumer
can construct identity23 or reuse its existing match. The join schedule
and runtime acceptance remain Unknown. — R2-ENGINE-379, R2-ENGINE-106

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

## Save and load screens

Both pages enumerate `\game*.sav` in the directory an unread call returns
(read as the current directory, Medium) and list a file only when its first
four bytes are `Bsg&`; the row shows the 256-byte label read at `blobEnd`.
In EN, rows run newest last-write time first; the RU order is Medium, since
its comparator is unread. A file without `Bsg&` is not listed and cannot be
loaded from the page. The load page lists `game9998.sav` and `game9999.sav`;
the save page does not (Medium: the name compare is an unread call).
— R2-SESSION-179, R2-ENGINE-438, R2-ENGINE-439

The save page's first row is `<<< Empty Slot >>>`. Saving there writes
`gameNNNN.sav`, four digits, with the smallest number not used by a listed
file (Medium: the sort comparator is unread). Saving over a listed row, or
typing a name equal to a listed label (the compare rule is unread), rewrites
that row's file with no confirmation. In application mode 2 the save driver
appends the typed label after the blob (Medium); in other modes the name
travels in a type-7 record and no label is appended. The save page runs in
any mode. Delete asks for confirmation and removes the file. The main menu
disables Load Game when no `game*.sav` exists in that directory.
— R2-SESSION-179, R2-ENGINE-440, R2-ENGINE-437

## Related

The character file is a separate format: [ROM2 character file](../rom2-a2c/format.md).

## Automatic-casting value sources

The acquired RU Player preferences constructor sets byte+0x20 to 0xff.
The Player serializer passes pointer+0x34 to a 36-byte transfer, which includes
that byte. Paired EN/RU preference serializers request 36 bytes in both
archive directions. Complete successful preference-block LOAD restores its
saved byte+0x20 over the constructor byte. Short transfer handling, later
application refresh and final ordinary-SAV cold-LOAD precedence remain Unknown.
The separate application constructor's DWORD 0x19 is not the Player constructor
value. — R2-SESSION-152, R2-SESSION-164

A selected RU mission-setup excerpt resends the application's AutoCasting
DWORD. Its complete transition/receiver schedule was not established.
Constructor, saved value, registry/application input and later resend remain
distinct carriers; final new-game, mission, town and cold-LOAD bytes are Unknown.
— R2-SESSION-153

### Shop parameters and current stock

Town-2 shop parameters occupy relative Scenario-prefix bytes
0x10a0..0x10ef and town-3 parameters occupy 0x10f0..0x113f. Each row holds
four 20-byte category cells after the 4096-byte bank. The selected paired
EN/RU transfer and offset arithmetic are High; these are not absolute
SAV-file offsets (`R2-SESSION-156`).

Current stock persistence, rebuilding and cold LOAD remain Unknown. The
selected owner/base serializers do not directly access the shop pointer,
but another serializer does access the same field offset and raw client
payload producers remain unclosed. No SAV byte range has been identified
as current stock (`R2-SESSION-157`).

A load through the load entry (`L2.00826` / `L2.00119`, reached from the
load page) does not keep the stock. In a town the load entry sends shop
refill 0x3f before the town entry, so all four categories are drawn again
from the restored current location's restored record; in a mission the next
town arrival draws. The `.text` references to the static shop owner are
four per root; the alternate owner serializer `L2.01660` transfers `+0x6c`
and is unclosed. The shop vtable's slots 2..4 are no-op stubs, read as
serialize slots by analogy to the owner table, so no SAV byte is expected to
hold the lists (`R2-SESSION-187`, High for the load order, Medium for the
absence). Callers of the session load that bypass the load entry were not
scanned.

## Campaign membership and serialized state

The scenario bank carries availability and selection flags, not the actor
objects. Slot773 occupies bytes0xC14..0xC17 of the 4096-byte bank. Hire
selection slots512..531 occupy0x800..0x84F, availability532..551 occupies
0x850..0x89F, and town gates552..571 occupy0x8A0..0x8EF. The selected
scenario SAVE/LOAD pair transfers these bytes; the actors and controlled
pointer use the Player/Group/Unit programs. Cold LOAD preserving the same
roster and literal binding remains Unknown. — R2-SESSION-158

The selected EN transition tests actor class WORD+0x0e, retaining the
reset branch for 33..63 and calling the detach/destroy route otherwise.
It does not test identity WORD+0x14c here. Earlier cleanup callees and
combat-death removal remain unclosed; this branch alone cannot identify
all surviving campaign members. — R2-ENGINE-382

## AreaEffect local scalar suffix

Both inspected AreaEffect serializers write four bytes for object fields
+0x4c..+0x4f, then a little-endian countdown word +0x50, after the base prefix
and before the owned effect +0x48. Pulse counter +0x4f and countdown +0x50
are local state needed by a running Blizzard. The RU base locally writes
completion +0x40 and +0x41 as two bytes after an unexamined ancestor prefix.
These are object-relative fields, not absolute SAV offsets. Ordinary SAVE
reachability, nested damage encoding, RNG persistence and cold LOAD continuation
remain Unknown; no preserved save was read (R2-SESSION-161).

## Nine-save LOAD contract

The selected EN Ok message 0x419 reaches the load entry after clearing the old
session. It creates a mode-2 server, decodes the document, then reopens the
physical tail. Document order is head, Players with all Groups and nested
actors, dead actors, worldPresent, optional world half, two sentinel DWORDs
and raw400 trailer. Zero worldPresent skips the world half. Nonzero reads
Buildings, SpellEffects, constructs the map from the saved scenario name,
then reads terrain, session, Sacks and the extra Token list. Session reading
is conditional on its global pointer being null. — R2-ENGINE-601

The client reads inline REG at blobEnd+256 and requests
CurrentState/InBattle. Nonzero selects mission setup; zero selects the town
continuation and posts 0x42e, whose measured entry reaches the town-display
routine L2.00266; view dispatch is Unknown. Both selected tail paths
read REG, pairs, Scenario and nine buffers in order. The label is skipped as
state. Incoming catalog raw16 is consumed without being copied into catalog
objects. Transport failure, later overrides and first-frame delivery remain
Unknown. The listing also contains a 0x421 arm and a zero-worldPresent
field-reset loop whose effects are unexamined. — R2-ENGINE-601,
R2-SESSION-033

The five mission-like files store CurrentState/InBattle 1 and take the
mission branch; the four town-like values are unread. — R2-SESSION-244,
R2-ENGINE-601

### Complete measured documents

All nine frozen EN revisions have exact structural document and physical
partitions under the selected programmes and complete successful transfers.
This framing result proves no native LOAD acceptance. — R2-SESSION-233

| Save | Decoded bytes | Players | worldPresent | Mission | Incoming current kind/ID |
|---|---:|---:|---:|---:|---|
| game0000 | 4486 | 1 | 0 | 0 | 2/1 |
| game0001 | 72178 | 8 | 1 | 10 | 1/10 |
| game0002 | 4418 | 1 | 0 | 0 | 2/1 |
| game0003 | 4408 | 1 | 0 | 0 | 2/1 |
| game0004 | 72188 | 8 | 1 | 10 | 1/10 |
| game0005 | 72482 | 8 | 1 | 10 | 1/10 |
| game0006 | 72936 | 8 | 1 | 10 | 1/10 |
| game9998 | 4408 | 1 | 0 | 0 | 2/1 |
| game9999 | 72092 | 8 | 1 | 10 | 1/10 |

The four world-absent documents have one Player and an empty dead list. The
five world-present documents each have 15 Building objects and one Pointer.
Only game0006 has one dead Unit and two world-effect records, SpellTransport
and PointEffect, with a nested Effect_DirectDamage. Sack and extra counts
are zero throughout the five mission documents. — R2-SESSION-234,
R2-SESSION-235

Building payload is base38, raw22 and scalar widths 1,2,2,2,1,1,1,4,4:
78 bytes. Pointer adds a DWORD: 82. Selected Sack is base38, a DWORD and
the bag programme; these files exercise no Sack body. — R2-ENGINE-603

Game0000, game0004 and game9999 each have one codec word-suffix byte after
the document. It equals the low byte of the even decoded length. The
selected RU writer rounds an odd archive length up and appends one byte from
that rounded-length variable before compression. The EN writer was not read,
so the cause in the EN saves is an inference from the byte match. The
document reader reads no record there. — R2-SESSION-233

### Measured Human and book state

All 14 Humans are walked separately. Each Player-1 record has one Group Human
whose self key equals Player +0x38 and whose owner key equals Player's self
key. Five other Humans belong to Player 8. These numeric matches do not
establish final active-client membership. — R2-SESSION-236

Name is Unit +0x80's CString. Attributes are +0x84..+0x8a, current/max HP
+0x94/+0x96 and mana +0x9a/+0x9c. Skills are +0xa8+2i, bases +0x116+2i,
main skill and five experience DWORDs are the Humanoid raw24 at +0x23c,
and the sum is +0x130. Parameter-title joins retain their existing Medium
grade. Worn items and ordered bag entries are separate. — R2-SESSION-236,
R2-ENGINE-319

The base raw12 block stores tile X/Y in bytes 0/1 and fractions in bytes 4/5.
Player-1 positions are 0/0 in the four town-like files, 11/41 in
game0001/0004/0005/9999 and 21/41 in game0006; fractions are 128/128.
Selected RU base LOAD copies this block and registers saved keys. Its
measured map helpers compare complete DWORD keys and return the stored
object, separately from archive indices. Later placement remains Unknown.
— R2-SESSION-236, R2-ENGINE-604

A present book at Unit gate +0x140 has DWORD +0x18, DWORD size n and references
for slots 1..n-1; slot 0 is omitted. Each new Spell payload is nine bytes:
u8 ID, u8 range, u8 defensive, u16 cost and u32 self key. LOAD rebuilds the
row pointer from ID. Damage/duration are outside this payload. Book DWORD
+0x18 and later cache precedence remain Unknown. — R2-ENGINE-602

| Player-1 saves | Book | Size | Nonnull book IDs |
|---|---|---:|---|
| game0000, game0001 | absent | 0 | none |
| game0002, game0003, game0004, game9998, game9999 | present | 28 | 24,27 |
| game0005, game0006 | present | 30 | 1..29 |

The full books occupy 327 bytes after the gate; sparse Player-1 books occupy
80. IDs 1..29 equal the published spell domain. Weapon-attached spells and
Player-8 books are separate. Native book delivery and castability remain
Unknown. — R2-SESSION-237, R2-ENGINE-019

### Mission-like continuation

Mission setup with argument 1 skips the map caller and the autosave, sends
packet type 4, runs one server tick when app +0x5d4 is set and pumps client
messages until a type 0x64 message returns after a type 6 message built the
client map; then it restores REG state, sends 0x408 through slot +0x48 of
app +0xd4 and stores the tick period and time stamps at app
+0x418/+0x414/+0x424. The messages before the 0x64, whether the clock runs
and what receives 0x408 are Unknown. — R2-ENGINE-613

The world half names the map asset by session +0x90, prefixed `Scenario\`
when session +0xd4 is nonzero. The asset gives terrain size and tiles
through the terrain constructor, is passed to the script compiler and is
freed; the document reader feeds the Units, Buildings, SpellEffects and
Sacks readers. What the terrain-layer reader, the second world object and
the script compiler take from the asset is Unknown. The client constructor
builds two families of objects from the asset named in message 6: static
objects from list +0x390, keyed entry +0x10 + 0x6000, and unit objects
(Medium) from list +0x3a4, a 0x254 or 0x208 byte object and then a 0x1e4
byte object keyed word entry +0x3c + 0x6000. The constructor beyond these
loops is unread. — R2-ENGINE-616, R2-ENGINE-615

LOAD rebuilds rather than reads: Building +0x3c from byte +0x40, base
+0x18 cleared, each base self key entered in session +0xdc, the base link
key and (RU Human) the DWORD keys at +0x40, +0x44, +0x5c, +0x64 and +0x68
replaced by objects or 0 through the unit-registry walk of slot +0x24.
Consumers of the loaded Building and Pointer fields after LOAD are not
searched. — R2-ENGINE-617, R2-ENGINE-618

| Building offset | Width | Note |
|---|---:|---|
| base (position +0x10, +4, +0xc, +0xe, +8, +0x18, +0x1c, self key, owner +0x14) | 38 | +0x18 cleared on LOAD |
| +0x4a | 22 | raw |
| +0x40 | 1 | selects +0x3c on LOAD |
| +0x42, +0x44, +0x46 | 2 each | |
| +0x48, +0x60, +0x61 | 1 each | |
| +0x64, +0x68 | 4 each | |

— R2-ENGINE-618

REG restore destinations: GameOptions to client +0xab0..+0xabc, app
+0x41c and global ShowTimeFlow; View X/Y to client +0x5c/+0x60; SpellBook
IsOpen and Pressed; Objects Selection and Group0..9 by object ID;
Inventory IsOpen; Projectiles FreeIndex and Prj%d objects; Fog runs into
bit 0x8000 of the client map words, read only when InBattle is nonzero.
Speed 0..8 selects a divisor 8..32 and app +0x418 = 1000/divisor.
— R2-ENGINE-614, R2-ENGINE-620

In the five files View is 8/29, Selection is [1], the Player-1 Human's
base +4 and that of no other object in the Players and Buildings field
table, and the book is open only in game0005 and game0006; game0006 alone
restores a projectile. The restore asks for object 1 to be selected if the
client object list holds it. Pause state and the camera meaning of View
are Unknown. — R2-SESSION-244, R2-SESSION-245, R2-ENGINE-621

Spell 29 (Slow) has data, an apply arm and cast ID 66, and the EN table
L2.02137 has no row for it: the table omits rows 9, 14, 15, 28 and 29, and
the client reads 24 bits of the slot mask. Row 29 is equal in the client
`world.res` (identical in both roots) and the RU server `world_srv.res`; the
EN install has no server data file. The game0005 and game0006 books hold
it. — R2-ENGINE-619, R2-SESSION-237

### Semantic limits

All nine physical tails end exactly after inline REG, one pair, Scenario and
nine zero-length buffers. Current indices are zero; only game9998 has two
availability records. Incoming keys support the four town-like and five
mission-like classifications. Current-pointer delivery remains Unknown.
— R2-SESSION-238

Complete framing leaves most head/raw-Player fields, embedded Group state,
position bytes outside the coordinates, portions of Unit raw blocks, terrain
raw60, session raw10374, trailer400 and most REG/bank/grid/raw16 meanings
unassigned. Nonempty Sack/extra restoration and live effect continuation are
unobserved. Cold LOAD with saved-key, field, effect/target-tick and first-frame
witnesses would settle delivered state. — R2-SESSION-233, R2-SESSION-235,
R2-SESSION-236, R2-SESSION-238, R2-ENGINE-601, R2-ENGINE-603,
R2-ENGINE-604

## Restored effects and command continuation

Selected EN LOAD readers restore SpellEffect bytes +0x40/+0x41;
SpellTransport adds nested +0x48/+0x4c and countdown word +0x50;
PointEffect adds nested Effect +0x4c and saved target DWORD +0x48;
AreaEffect adds bytes +0x4c..+0x4f, word +0x50 and nested Effect +0x48.
These are object-relative offsets. The list is rebuilt, then postload
walks its members and nested records. Some saved-key misses become zero;
position-link misses preserve their number. The bounded reader census
does not close computed callers or bulk copies. — R2-ENGINE-659

The selected EN world tick decrements the restored transport countdown,
queues a nested effect at expiry and marks the transport complete. Point
application dispatches its payload to an eligible target and completes.
Same-tick versus next-tick execution, live target delivery and exact HP
change remain Unknown. game0006 has one top-level transport, countdown 1,
with a nested point and DirectDamage; it is the two world-effect records
already counted, not two top-level members.
— R2-ENGINE-660, R2-SESSION-260, R2-SESSION-235

The Unit +0x1c4 command object stores raw184 then a word list. LOAD replaces
the saved list pointer with a newly allocated list. Unit +0x50/+0x54/+0x58
are separate raw DWORDs. EN postload remaps seven nonzero command keys at
+0x0c,+0x10,+0x18,+0x20,+0x28,+0x30,+0x68 only on lookup hits; misses
retain their number. This narrows the earlier unread-subobject Unknown.
— R2-ENGINE-662, R2-ENGINE-617

The read command consumer uses bytes +8/+9 and optional word +0xb0.
Zero selectors and zero +0xb0 enter no selector arm; movement state and
earlier gates can still act. All five measured EN heroes have Unit +0x50=12;
game0001/0005/9999 have zero selectors, game0004 has 8/2 and game0006 has
6/1. The saved DWORD alone does not prove an active order or idle state.
— R2-ENGINE-663, R2-SESSION-259

The bounded read gate set includes worldPresent, archive class/index checks,
saved-key lookup outcomes, completion/countdown, projectile remaining
segments, Unit HP/condition and manager/group flags. It is not a complete
LOAD failure census. Fifteen selected EN/RU programmes have equal instruction
shape after masking operands >=0x400000; callees, complete scheduling and
RU projectile continuation are outside that comparison. Save values are EN-only.
— R2-ENGINE-664, R2-ENGINE-665

The separate REG projectile has its own continuation contract: a stored
object with zero segments requests removal on its first read update path,
conditional on unchanged entry state. Actual first-frame behavior remains
Unknown. — R2-ENGINE-661

## Mission LOAD view and idle gates

The EN REG restore writes View/X and View/Y to client +0x5c/+0x60.
Selected constructor, resize, panel-layout and draw writers store to
the same fields and can change the origin whenever they run; their order
against the REG stores is unread. The bounded writer population does not
close the last writer at first presentation or identify a followed unit.
— R2-ENGINE-667

The app +0xd4 root slot +0x48 is L2.03545 in EN. Its forwarding path
conditionally reaches four direct children, and a store runs only when
the child's own base call (L2.00913) returns zero for message 0x408 and
the forwarder delivers the message to that child. Their mapped 0x408
branches store +0x6c=-100, +0x60=1 and +0x5c=1 in three distinct child
objects; the fourth takes its local default. Full child delivery and the later
visible effects remain Medium/Unknown. — R2-ENGINE-668, R2-ENGINE-613

Acquired EN unit slots select through L2.03528/L2.03529, storing +0x80
and +0x114. The entry opens or closes the book and inventory from REG
SpellBook and Inventory IsOpen through controls that are separate from
the Selection restore. Client key-1 identity/membership and actual first-frame panels remain
Medium/Unknown. — R2-ENGINE-670, R2-ENGINE-614, R2-ENGINE-621

EN schedulers read app +0x418 in a millisecond deadline compared with
timeGetTime. Campaign idle dispatch bypasses its local mission tick
helpers on UI mask 0x4008; otherwise +0x434 selects an immediate or
timed helper. Actual clock advance before the first input remains
Unknown. The modal-bit simulation Unknown is narrowed at that caller:
in app mode 2 the dispatcher skips the local tick helpers when
+0x404 & 0x4008 is set; whether the modal box stops the simulation in a
loaded mission stays Unknown.
— R2-ENGINE-669, R2-ENGINE-671, R2-ENGINE-620, R2-ENGINE-609

The acquired EN/RU mission entries have 1002 instructions each and
24 app-field differences after address masking. Room-bind bodies have
55 each and one UI offset pair, +0x404/+0x41c. After those substitutions
their instruction shapes agree. No other newly read body or referenced
data is covered by that equality. — R2-ENGINE-672
