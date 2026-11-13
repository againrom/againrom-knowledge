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

The sub-object writes every entry it holds, with no cell test: the count,
then buckets 0..99 by `(key >> 4) mod 100`, chain order within a bucket (the
buckets of the nine files do not decrease; the order inside a bucket is not
tested, and the slot value that binds the map to this body is Medium). An
entry is created at the first field set on its cell, which gives plane one
bit 0x20, and deleted when a clear leaves it empty, at 8 create and 6 delete
sites among the direct calls in EN code (Medium that no other path creates or
deletes one). Refresh sets the cell's plane-two byte to its plane-one byte
with bit 0x20, so a refreshed entry cell is in the key list (Medium that every
entry cell is refreshed). The key list also holds cells without an entry. In
the nine mission saves of the two owner directories
every entry key is in the key list with plane-two bit 0x20:

| Mission | Files | Keys | Entries |
|---|---|---|---|
| 10 | 5 | 1,782..1,830 | 125..126 |
| 20 | 2 | 2,131..2,282 | 474..475 |
| 30 | 2 | 3,623..3,638 | 427..430 |

— R2-SESSION-310, R2-ENGINE-758

A raw60 entry, at stored entry offsets (field names from read uses are
Medium):

| Bytes | Field |
|---|---|
| +0x00, +0x01 | plane bytes copied from the cell at create and written back by refresh |
| +0x02 | layer count |
| +0x03 | zero at create; no access in the routines read (Medium) |
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

Each byte group comes from its own setter at SAVE: the plane bytes from the
cell at create, +0x02 from the effect-layer routine, +0x04/+0x08 from
placement, +0x0c from the building setter, +0x10 from the object setter,
+0x14..+0x2b from the effect setter, +0x2c/+0x30 from the footprint
routines, +0x34..+0x39 from the event writer and +0x3a from create. In the
nine mission saves (2,432 entries) bytes +0x02, +0x03, +0x10..+0x2b and
+0x34..+0x39 are zero in every entry, +0x00 is nonzero in every entry (6..16
in the mission 10 files, 6..8 or 6..9 in the others), +0x01 is 0 or 1 in the
mission 10 files and 0, 1 or 5 in the others, and +0x3a equals the key.
— R2-SESSION-311

The footprint pass keeps +0x2c and +0x30; a nonzero event kind acts when an
occupant is placed on the cell (kind 0x17 stores the target pair into the
unit's command word) or arrives there; refresh reads +0x3a as the cell to
write back. — R2-ENGINE-673

Session is 4,000, 1,000, 48, 400 and 4,908 raw bytes, then `u8`, `u8`, `u32` and
three `u32`: 10,374 bytes against ROM1's 4,374. Only the first and fifth
widths differ from ROM1. — R2-SESSION-019

| Session bytes | Part | Written by | At LOAD |
|---|---|---|---|
| 0..3999 | 1,000 DWORD registers | session init, the script compile, the script tick, script nodes | kept; the script compile stores each 65538 check's parameter 0 again |
| 4000..4999 | trigger latches | the trigger runner | kept |
| 5000..5047 | performance timer: counter start 0..7, stop 8..15, frequency 16..23, ticks 24..27, milliseconds 28..31, frequency / 1000 at 32..35 | the session constructor and the turn routine | kept; the next turn rewrites start and stop |
| 5048..5447 | 400 bytes cleared by init; no other user found | session init | kept |
| 5448..10355 | diplomacy: bytes 1, 1, 0, 0, owner DWORD, 70 x 70 matrix | the diplomacy constructor and methods | kept; the owner DWORD is replaced at post-load |
| 10356, 10357 | two bytes with zero stores only | session init | kept |
| 10358..10361 | a counter, a divisor of the debug statistics; equal to register 999 in the nine files (Medium that it is a turn count) | the turn routine and, twice per call, a routine with no direct caller | kept |
| 10362..10373 | the result-tick DWORDs of R2-ENGINE-623 and R2-ENGINE-624 | script nodes | kept |

In the nine mission saves the frequency is 10,000,000, bytes 5048..5447 and
10356..10357 are zero, bytes 10358..10361 equal register 999, the diplomacy
bytes 0..3 are 1 1 0 0 and the matrix diagonal holds 0x12 at 8, 7 and 6
cells in missions 10, 20 and 30. The meaning and readers of bytes
5048..5447 and 10356..10357 beyond the transfer are Unknown. — R2-ENGINE-760,
R2-SESSION-312

Eight Againrom-written mission-10 files and their original resaves store the
diagonal as 2 with no cell holding bit 0x10. The EN server sets a Player's
own relation bit 8, which the fog stamp needs, only from the diagonal's bit
0x10; see "EN mission-10 sight, map reveal and command timing". —
R2-SESSION-339, R2-ENGINE-826

The first 4,000 bytes are 1,000 DWORD registers, register n at bytes
4n..4n+3. Session init sets register 901 (bytes 3604..3607) to 1 and
register 999 (bytes 3996..3999) to 0; one routine adds 1 to 999 before it
runs the triggers. No other reader of either was found in the scanned code
and in the 209 + 48 + 197 map operands of check 19 and instants 3 and 8.
The five EN
mission-10 saves hold 901 = 1 and 999 = 1, 1, 17, 141 and 0.
— R2-SESSION-265

The six register compares of the trigger runner `L2.00517` read registers
named by check slots for check ids present in the map, below 54 in both
46-map corpora; a trigger naming an absent id is unread. Registers 100..999
have direct writers only in session init, the script tick (999) and
instant 1 (902). In the nine mission saves they are nonzero only at 901 (1)
and 999 (1, 1, 17, 141, 0; 1, 91, 5, 0). — R2-SESSION-313

On LOAD the terrain post-load pass replaces each nonzero pointer DWORD by
its saved-key lookup and keeps the saved number on a miss; bytes
+0x00..+0x03 and +0x34..+0x3b, the session block, the Player raw blocks and
the Building base word +0x0c pass without a value test. — R2-ENGINE-676

LOAD inserts each stored entry without a key, plane or count test; a
repeated key replaces the earlier entry. Plane one and plane two of a listed
cell come from its key, not from the entry; a cell outside the key list
keeps the planes the map construction gives; plane zero is not stored. No
LOAD step read re-runs create or refresh (Medium for later steps).
— R2-ENGINE-759

## Building casters

The list `[Session+0x7c]+0x10` holds the building casters built from type-9 records. The document serializer carries it
in both directions through `R2.0173`: count, then each object; LOAD clears the list and reads each object back and
appends it (the clear, count and object reader bodies are unread; Medium that LOAD rebuilds the list from the file).
Whether the map build that LOAD runs also runs the caster stage is Unknown. The caster stores, after its base, byte
`+0x3c` and the 6-byte block that `+0x40` points to. The list sites are read in RU; the EN site is Medium. This list
is the "extra" list of the paragraphs above; its observed counts in five saves are zero (R2-SESSION-235, R2-ENGINE-601).
The layout row "ROM2-only object list" and the Unknown "Nonempty Sack/extra restoration" name the same list. — R2-ENGINE-1011

Bytes 0 to 3 of the block are the spell, power and cell bytes of the map record; bytes 4 and 5 are not written by the
arm. At run time the caster's target is chosen and the cast entry `L2.05951` called on the chosen unit. — R2-ENGINE-1008

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

In EN, the unit-list routine of the tick adds 1 to the byte at Player +0x44
plus the victim's identity word +0x14c (identities 1..0x9ff) while that byte
is below 14, after it adds 1 to +0xa4c (victim class 0x21..0x3f) or +0xa48
(any other victim); a character-file load copies 0xa00 bytes to +0x44, LOAD
reads the 2,560 bytes from the save, and the constructor clears them. That the
routine counts kills is Medium: the meaning is taken from its conditions (a
unit in state 0x10 with a live killer), not from a run. — R2-ENGINE-772

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

### Repeated-victim history

The read RU and EN Unit SAVE/LOAD bodies have no explicit transfer or reset
of the repeated-victim array at Unit +0x1a8. The Unit save loads the
dwords at +0x1c0 and +0x1c4 as receivers of helper calls; the widths 180
and 184 and the raw-block reading are R2-ENGINE-209's, and the helpers were
not re-read. On that basis the two blocks are separate pointees (Medium)
and do not cover that embedded header. No file
offset or width for history is established. Omission is Medium: helper,
global-alias and archive-construction effects remain unclosed. Under a
fresh constructor (R2-ENGINE-1173) and unchanged helper behavior, the code
read in both roots, with no run of the original, has LOAD start with an
empty list; a reused destination can keep its prior history. Native LOAD
and the next award remain Unknown. — R2-SESSION-410, R2-ENGINE-1173, R2-ENGINE-209

Recursive `.sav` pathname enumeration found no file in the two named ROM2
EN/RU install trees (pathname match only; archives and other extensions
not searched). That bounded directory result is High; the original saves
read by earlier claims were not re-examined for list bytes, so a byte
witness from these two roots and runtime persistence remain Unknown.
— R2-SESSION-411

The raw 180 block is the movement block at Unit +0x1c0. Byte 0 is the
facing and byte 1 the heading; LOAD reads the block back over the one the
constructor filled. In five original saves 355 of 356 Unit and Human
records hold the two bytes equal. High for the RU transfer and the
values; Medium for EN and for a loaded unit keeping them. — R2-SESSION-362

The class a summoned creature belongs to (vtable `L2.02314` in RU) has this
programme in slot +8, read from the class and not from a file. Its owner
+0x14 and +0x1a0 are stored; no step covers +0x148, the caster key that the
summoned-creature cap counts. Which units a mission save writes, and what
+0x148 holds after LOAD, are Unknown. — R2-SESSION-326

A unit that spell row 22 builds from a corpse with the constructor has the
same class and programme. Once the unit is placed, owner +0x14 is the
caster's, +0x1a0 is 1 and +0x148 is 0. The packed u16 +0x14c holds the
template serverID (657, 664 or 628) when the template loader runs, which
is inferred from a string layout. A unit the arm reuses from the dead-unit pool keeps its own class,
whose programme was not read. Which units a mission save writes is
Unknown. — R2-SESSION-327

A unit that Summon builds gets +0x1a0 = 1 before its placement. Placed,
it gets a new id at +4, owner +0x14 and key +0x148 from the caster
(+0x14 and +4), a new Group at +0x70 in the owner's group list with order
code +0x50 = 0xb, and the coverage word +0x1a4 from the recount. The
state sender ORs a field mask into the dword +0x114 + 4 * id for a player
id 16..31 that does not cover the unit while Session +0x74 is nonzero, and
clears that dword when it sends. The cell test reads footprint mask byte
`[+0x1c0]+5` (0x41 from the block constructor) and tokenSize +0x49. A
placement that fails leaves +0x1a0 = 1, no id, owner or Group store, and
the unit off the grid. The cast's recovery byte +0x6c on the caster
includes a draw of 0..3, taken in the actor step that runs the apply.
Session +0x74 has two known stores: server init's and the map build's.
Which of these a mission save writes is
Unknown. — R2-ENGINE-788, R2-ENGINE-789, R2-ENGINE-790, R2-ENGINE-791,
R2-SESSION-329

An Item writes the 38-byte base programme that Unit also calls first,
then u32-counted references and
u16, u16, u8, u8, u8, u16, u16, u8. Weapon adds raw 24, raw 22, u8 and a
reference; Armor adds raw 22 and u8. — R2-ENGINE-210 (amended)

The Weapon reference is the weapon's Spell: LOAD reads it with class record
RU L2.01180 / EN L2.05316, named Spell (size 0x14, schema 1). The EN LOAD bodies read (Weapon
`L2.05355`, Spell `L2.02975`) test nothing on it; the object reader `L2.03455`
and the RU Spell bodies are unread. — R2-ENGINE-904, R2-ENGINE-210

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
not read. That dword is the first of an object a separate routine constructs
(Medium: the constructor body and the object's other readers are unread), and
the skill-from-experience loop reaches it only for E at or above T(149)
(R2-ENGINE-1159). — R2-ENGINE-485 (High for the law, Medium for the exact values and
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

A run measured the hero field the Hall score reads as 62653614; that dword occurs once each in two saves at file
offsets 0x2010 and 0x200f. Its block is not established (Medium). — R2-SESSION-390

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

At death the bag first takes unit +0x78, +0x74 when its row parameter 15
is not 0, and the slots; a template name holding "NPC" (Medium) or +0x1a0
empties it. The item roll then adds at most one generated item with
quantity 1, and a nonempty bag or nonzero gold drops in the sack of the
unit's cell. Whether the death's walk of unit list +0x20 and its virtual
calls +0x38, +0x48 and +0x4c drop anything else, and how a sack already on
the cell bounds the merged bag, are Unknown. — R2-ENGINE-943, R2-ENGINE-940

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
R2-SESSION-035, R2-SESSION-036, R2-SESSION-037. The two final u32 of the pair
record are the Hall +4 and Hall +8 fields the Hall of Fame score reads. — R2-SESSION-384

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
can construct identity23 or reuse its existing match. Leaving mission 40
through the leave arm is the request; runtime acceptance remains Unknown.
— R2-ENGINE-379, R2-ENGINE-106, R2-ENGINE-744

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

## Town SAVE writer

With the application mode word 2 the save driver first has the server write the whole
envelope and closes the file: header, encoded document, then blobEnd patched
at offset 4. It then reopens the file and appends the 256-byte label from the
application, the REG store, the pairs, the Scenario record and the nine
cells. An odd document length gets one byte, the low byte of the rounded
length. Callees of the writer (the pre-save list walk, the archive, the
encoder, the REG writer) are unread, so their effects are Medium.
— R2-SESSION-283

| Document field | Written from |
|---|---|
| head u32, u32 | server `+4`, `+0` |
| head CString | server `+0x90` |
| head 11 x u32 | server `+0x174`, `+0x17c`, `+0x180`, `+0x184`, `+0x188`, `+0x18c`, `+0x190`, `+0x194`, `+0x1a0`, `+0x19c`, `+0x198` |
| mission, difficulty | server `+0xd4`, `+0xd8` |
| Players metadata | Player map `+0x20` |
| worldPresent | literal 1 when server `+0x94` is nonzero, else 0 |
| sentinels | literal `0xBADFACE1` twice |
| trailer | 400 bytes of the server's `+0x170` object |

The dead list is written by an unread virtual slot (Medium), and the Player
list writer is unread (Medium). Field meanings of the head members and the
trailer remain Unknown. — R2-SESSION-284

The REG store of a town SAVE holds `CurrentState/InBattle` (application
`+0x404` bit 0; 0 in the four town files, the writers of the bit are
Unknown), `Character/Name`, six `GameOptions`, `View/X` and `View/Y`,
`SpellBook` `IsOpen`, `Pressed` and an always empty `Shortcuts`,
`Objects/Selection`, `Objects/Group0..9` written when the selection groups
are nonempty (not present in the four files) and `Inventory/IsOpen`.
`Projectiles` and `Fog` are written only when InBattle is set. The values of
the unread getters and the record layout are Medium. — R2-SESSION-285

In the four EN town files every byte of the decoded document and of the
whole physical file has a writer source; the inherited groups (class
programmes, REG, pairs, Scenario, cells) are Medium, a source and not a
meaning. All four store InBattle 0 and 22 REG records; the town branch is
reached from the load entry only through the unread `L2.01829` and the
InBattle getter (Medium). — R2-SESSION-286

Two town files, game0003 and game9998, have equal documents and REG values;
the bytes after the NUL of REG record names differ (Medium that they are
uncleared memory), so they are not stable between saves. Label bytes after
the NUL are zero. — R2-SESSION-287

The LOAD body returns 1 only on open failure, a magic other than
`0x26677342` and a version below `0x0BAD0002`, and tests no length or sum
itself; failures inside its callees and the archive were not followed, and
the second caller (the map caller) was not read. The load entry does not test
the result. The town tail seeks past the label by blobEnd. The town branch
does not request `SpellBook/Shortcuts`, `Inventory/IsOpen`, `Projectiles` or
`Fog`; that holds for the body read, not for the image. What the player sees
after a refused file is Medium to Unknown. — R2-ENGINE-723

Nine EN and RU writer and town-LOAD bodies are equal in instruction sequence
after masking every operand of 0x1000 and above, apart from application field
offsets; callees are not compared. — R2-ENGINE-724

In a town LOAD the Player load arm and the base and Unit fields with a
visible reader call are read into their own runtime offsets (`+0xa48`
through the XOR, `+0xa4c` and `+0xa54` sign-extended, `+0x30` also copied to
`+0x32`). The hero's raw blocks, the Unit sub-objects `+0x1c8` and `+0x1e4`
and the base's raw 12 bytes go to serializers that were not read (Medium).
Two stored values are discarded: the base `+0x18` word, which the load arm
sets to 0 after reading it, and Unit `+0x40`, `+0x44` and `+0x5c`, which the
document's zero-worldPresent arm sets to 0 for each Player's list; the arm
restates R2-ENGINE-616 and R2-ENGINE-664, and the four town-like documents
of R2-SESSION-234 hold world byte 0 (Medium that every town SAVE does).
Saved addresses serve as map keys and lookup inputs, and the Player's `+0x38`
is replaced in place. Whether a later step overwrites a field read into
place is Medium: the embedded serializers of the hero's raw blocks, the town
branch past the tail readers and the town entry's callees are unread. The
record order is, as a composition of the serializers' programs, the order of
first references (Medium). — R2-SESSION-314

In RU, the bodies behind the base's raw 12 bytes and the hero's raw blocks
request 0xc, 0x18, 0x16, 0x40, 0xb4 and 0xb8 raw bytes through two unread
helpers and also call routines that were not read; the 0xb8 body also rebuilds
its sub-object through two unread routines, and one more body sets one bit of a
global table. The raw blocks are restored as stored (Medium; RU only), except
that sub-object, which is Unknown.
— R2-SESSION-325

The Player prefix fields `+0xa44..+0xa5c` are stored as 0, `+0xa5c` as 0x5f,
at one site of an unread construction body; a character-file load sets
`+0xa48`, `+0xa4c`, `+0xa54` and the unsaved `+0xa58`; play-time sites
change `+0xa48` and `+0xa54`. The census covers memory operands with a
displacement only (309 rows per root); address formation by `add` or `lea`,
the other Player prefix fields and the groups H, B and T were not censused.
The construction, the sources and the events are Medium. — R2-ENGINE-762

### Stored values of the original town files

Three owner save copies hold 25 files; 11 are town files: their current
Scenario record has kind 2, worldPresent is 0 and mission is 0. Each of the 14
files with a kind 1 current record has worldPresent 1. The 11 stand at five
stages: town1/10 (five files), town2/30 (two), town2/40 (one), town2/50 (two)
and town3/60 (one). That every original town SAVE has worldPresent 0 is
Medium. — R2-SESSION-320

Fields equal in every one of the 11 town files that holds them (High for these
files; Medium as a rule of town files of copies 1 and 2 (EN) and copy 3, whose
root is not established; no RU file; fields are named by offset, no meaning is
assigned). Party Human rows hold for every Human present (the hero in 11 files,
identity 22 in 5, identity 23 in 3):

| Group | Field | Stored value |
|---|---|---|
| Player prefix | `+4`, `+8` | 1, 1 |
| Player prefix | `+0x2c`, `+0x41` | 0 |
| Player prefix | `+0x30` | 2 |
| Player prefix | `+0xa4c`, `+0xa50`, `+0xa54` | 0 |
| Player prefix | `+0xa5c` | 0x5f |
| Party Human | `+0x18` | 2 |
| Party Human | `+0x40`, `+0x44`, `+0x5c` | 0 |
| Party Human | `+0x78` | null reference |
| Party Human | `+0x98`, `+0x9e` | 100, 50 |
| Party Human | `+0xd4..+0xd7` | zero |
| Base, Armor, Weapon, Shield, Item | `+8`, `+0x18` | 0 |
| Base, Effect | `+0xc`, `+0xe`, `+0x14`, `+0x18` | 0 |
| Base, Item | `+0x1c` | 0xFFFFFFFF (one item in 6 files; no other kind holds it; Medium as a rule of Items) |
| Keys | Player `+0x38` | the hero's saved address |
| Keys | Human `+0x14`, Group `+0x44`, nonzero Armor or Weapon `+0x14` | the Player's saved address |
| Keys | Group `+0x40` | 0 |
| Tail | pair count, row | 1, (1, 1) |
| Tail | nine client cells | words 0, length 0 |

Base rows hold per record kind, in the files that hold the kind: Shield 1
record in 1 file, Item 6 records of one item in 6 files, Armor 64 records of 25
items in 7 files, Effect 57 records of 17 roles in 10 files, Weapon 34 records
of 10 items in 11 files. Saved addresses are distinct multiples of 8 in each
file. A copy's town files of one Player name share some saved addresses, across
stages too (copy1 game0002 and game0003 share 7 of 9); no saved address occurs
in two copies. — R2-SESSION-321

Every comparison of town2/30 with a later stage, and of copy1 with copy3, also
differs in copy: town2/30 occurs in copy2 only and town2/40 to town3/60 in copy3
only. The within-copy stage transitions are copy2 town1/10 to town2/30 and copy3
town2/40 to town2/50 to town3/60. Fields that differ between stages and are
equal within the stages tested (town1/10, five files in two copies, and
town2/30, two files; the two town2/50 files are identical, so that stage adds no
independent test): Player `+0x40` (0 at town1/10 and town2/40, 1 at the other
three), `+0xa44` (1 at town1/10, then 2: town1/10 against every later stage),
`+0xa48` (0, 6, 6, 23, 46 by stage), the raw 2,560 bytes (zero at town1/10,
nonzero at 1 to 13 byte positions later), byte 32 of the raw 36 bytes (town1/10
against every later stage), the two DWORDs after the pair rows (0/0, 184/6,
350/6, 781/24, 1487/60), the current index, `Objects/Selection` and the hero's
Group `+0x1c`. The availability entries (kind/id) are equal per list index
within a stage; the lists differ in count. `Objects/Selection` is `[2]` in
copy3 game0024 only (one file; it equals the hero's base `+4` in ten files that
do not discriminate and identity 22's in game0024, Medium, a matching number).
The rows of members 22 and 23 (`+0x8e`, raw blocks), Group(id22) and places 1,
2, 5, 6 differ between the stages that hold them and were tested within a stage
on the identical town2/50 pair only; member 22's `+0x8e` (400 at town2/30, 314
at town2/40 to town3/60) is a copy2 against copy3 difference, not a change
between copy3 stages. The stored `+0xa4c`, `+0xa50`, `+0xa54` and `+0xa5c` equal
the construction stores of R2-ENGINE-762 (Medium; the three zeros are not
specific, `+0xa5c` 0x5f is). No tie of a stage difference to one write event
is established: the writers below are read, the events between two saved points
are not. R2-SESSION-314's LOAD discards the stored Human `+0x40`, `+0x44`,
`+0x5c` and base `+0x18`. A stage difference may be a copy difference (Medium).
— R2-SESSION-322

Gold, the hero's `+0x8e` and raw 184 block, the group count, the availability
count and the keys also differ between town files of one stage (gold is also
about 99.9 million in copy3 against 44, or 1000 in one file, in copies 1 and 2:
a visible copy difference); one of the 12 same-stage pairs
(copy3 town2/50) is equal in the 400 compared fields and the listed raw blocks,
while the physical files differ in 100 of 8,633 bytes. Whether EN and RU town
files differ is Unknown: no file of the population is established as RU.
— R2-SESSION-323

### Play-time writers of the town file fields

EN addresses; the events are named from the conditions of the read bodies
(Medium), and completeness is Medium: stores through computed pointers and
whole-object copies were not searched.

The unit-list routine `L2.04418`, called from the tick method `L2.04407`,
takes each unit in state 0x10 (tested again after a virtual call) whose killer
is live with class 0x21..0x3f. A victim of class 0x21..0x3f adds 1 to the
killer's owner `+0xa4c`; any other adds 1 to `+0xa48`; after either, the
`+0x44` byte of the victim's identity gains 1. A unit method (vtable
`L2.04420`, slot `+0x9c`; no direct call in EN `.text`, and 11 of 13
`call [reg + 0x9c]` sites unread) also adds 1 to `+0xa48`. The unit death adds
1 to `+0xa54` when the dying unit is the owner's hero. A character-file load
copies 0xa00 bytes to `+0x44` and sets `+0xa48`; LOAD restores both; the
constructor stores 0. — R2-ENGINE-772

The join routine stores `+0xa44` as the requested number or, for a new Player,
the smallest of 1..15 that no Player with `+0x2c` 0 holds (16 when none is
free); the map Player builder stores map player entry `+0xc` + 1.
— R2-ENGINE-773

Of 24 byte stores at displacement 0x40 in EN `.text`, the Player `+0x40`
stores are the constructor (0), the result tick per Player (0, 2, the low byte
of `+0xbe14` of the object at `L2.01066`, or 1), the unit death (identity −
19, when Session `+0x74` is zero), a dispatcher handler (0 when the Player has
a hero) and the transition loop (0, when Session `+0x74` is nonzero). 13 rows
on unidentified objects are not excluded. — R2-ENGINE-774

The two DWORDs after the pair rows are `+4` and `+8` of the pair object at
application `+0x598`: `+4` grows by Session `+4` / 16 at `L2.04422`, the one
site found that calls its adder (Medium: no call scan of the adder), `+8` by 1
at `L2.04425`, the one direct call of its adder; three sites clear both.
— R2-ENGINE-775

The writers are consistent with the stored values of the 11 town files and do
not explain them. For a counter (`+0xa48`, the tail DWORDs, the `+0x44` bytes)
any non-negative value lies in the writer's set. `+0xa48` (0, 6, 6, 23, 46)
and both tail DWORDs do not decrease within a copy; the nonzero `+0x44`
positions are below 0xa00, a bound the unit list, the character-file copy and
LOAD all give; the zero `+0xa4c` and `+0xa54` fit no counted event; `+0x40` 1
is a result-tick value. `+0xa44` 1 lies in the join's set (the smallest free
number) and in the builder's (entry `+0xc` = 0); 2 lies in the builder's set
only if a map entry `+0xc` is 1, which was not read, and the order of the two
writers is unread. Tail `+8` differs from `+0xa48` at town2/50 and town3/60,
so they are not equal counts of one event over the same span; resets are not
excluded. No writer of the hero's Group `+0x1c` was found in the RU Group
constructor chain. The number of events between two saved points was not read
(Medium). — R2-SESSION-324

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

By the code read, a file written at stage 50 before the next mission exit is
expected to hold 773 = 23 at 0xC14 and to keep identity 23 as a companion in
a Player Group (Medium); no file read holds identity 23 or a nonzero 773.
— R2-SESSION-302

The selected EN transition tests actor class WORD+0x0e, retaining the
reset branch for 33..63 and calling the detach/destroy route otherwise.
It does not test identity WORD+0x14c here. Earlier cleanup callees and
combat-death removal remain unclosed; this branch alone cannot identify
all surviving campaign members. — R2-ENGINE-382

In the owner's EN Kaarg town file written, by the owner's account, by the
automatic save of the abort-and-return from mission 30, Player 0's group
list holds two Groups of one Human each: the hero (+0x14c low word 21,
class +0xe 35, +0xc 28) and identity 22 (33, 30). No hire Human is stored
although selection 512 = 1 and availability 532 = 1 stay in the bank; the
mission-30 files of the same run (`game0012`, `game9999`) hold the hire as
a third Group (+0x14c 1, class 15, +0xc 264). Every member's owner key is
the Player's. No town file written after a victory was read.
— R2-SESSION-295

The EN group list equals the RU one after masking; LOAD rebuilds each Group
and appends it to the list (High for the sequence; Medium that it is the
Player's list). — R2-ENGINE-732

A companion is kept in a town file as a Human in a Player Group, with its
identity word in the low word of the packed +0x14c field and its class in
+0xe. In the one town file read, an abort-and-return save, no hire Human is
stored while bank cells 512+i, 532+i and 552+i stay, and the hire is built
at the next mission entry; whether a victory save keeps a hire Human is
Unknown (Medium for the rule). — R2-SESSION-296

## AreaEffect local scalar suffix

Both inspected AreaEffect serializers write four bytes for object fields
+0x4c..+0x4f, then a little-endian countdown word +0x50, after the base prefix
and before the owned effect +0x48. Pulse counter +0x4f and countdown +0x50
are local state needed by a running Blizzard. The RU base locally writes
completion +0x40 and +0x41 as two bytes after the ancestor prefix.
These are object-relative fields, not absolute SAV offsets. Ordinary SAVE
reachability, nested damage encoding, RNG persistence and cold LOAD continuation
remain Unknown; no preserved save was read (R2-SESSION-161, amended).

The ancestor prefix, in both roots, writes the position object +0x10, dword
+4, word +0xc (spell row), word +0xe (2*row+9), dword +8 (area mode 0, 1 or
2), word +0x18 (zeroed at LOAD), dword +0x1c, the object's own save key and
the owner key +0x14. The programme is the same for every mode. The three
bodies read do not write the area's caster +0x3c, the words +0x42, +0x44
(zeroed at construction) and +0x46, or the map cell registrations of a mode-1
area. The owned effect's +0x44 holds a copy of the caster; whether the nested
payload writes it, and whether the map saves the registrations elsewhere, is
Unknown (Medium for the dword width of the writer). — R2-SESSION-330

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

In the 16 original-made saves of 2026-10-30, all 150 non-hero Humans store an
empty +0x80 name, including 6 saved headman records (+0x14c 2112; 3 map placements, maps 20
and 30) and 2 saved records of a word-1 hire (+0x14c 1, probably one hire);
no other hire ID is saved. Each of the 16 heroes (+0x14c 21) stores 5 or 9
bytes. — R2-SESSION-352

Human LOAD stores base words +0xc (Humans row index) and +0xe (type), Unit
bytes +0x49..+0x4c (face +0x4b, gender in bit 7 for a non-hero) and the low
word of the u32 that also carries +0x204 and +0x1a0 (+0x14c). The u8
transfers of +0x49..+0x4c and the packed +0x14c rest on unread archive
helpers (`L2.03712` and `L2.03713` for the bytes), so their widths are
Medium. The client takes its type from +0xe and its face from +0x4b; that the
client look after LOAD is rebuilt from them and the stored item references is
an inference (Medium). — R2-ENGINE-682

A Human's definition is the unsigned u16 pair +0xc, +0xe at Human payload
bytes 16 and 18 (base bytes 16 and 18). SAVE writes the stored words as they
are; the template load sets them at creation (Medium that nothing else changes
them). It sets +0xc to the first Humans row whose name it finds equal to the
template name (the compare is Medium) and +0xe from that row, or to 0x21 or
more for a hero (the exact hero value is Medium). — R2-ENGINE-739

| +0xe | LOAD sets row pointer +0x3c to | Note |
|---|---|---|
| below 0x21 | Humans element +0xc | +0xe is not checked against the row |
| 0x21 or more | Humans element 5 (`Man_Sword2Handed` if the collection is the `data.bin` Humans table, Medium) | +0xc is not used to select the row |

The Human LOAD bodies read make no bound check, serverID lookup or name
search; +0x14c is not used for the row. — R2-ENGINE-740

The Humans table stores 290 elements in both roots; elements 1..289 are
loaded rows and element 0 is never loaded. In the Human LOAD bodies read a
+0xc of 0 or above 289 is not rejected; what it yields is Unknown.
— R2-ENGINE-741

In the 22 owner saves all 216 Humans store +0xc in 1..289 (highest 274) and a
+0x14c word equal to that row's serverID word; 29 store a hero type (33 or
35) and so load element 5. — R2-SESSION-298

A record for which slot +0x30 is false loads template row pointer +0x3c as
Units element +0xc, with no bound check (that slot +0x30 separates Humans is
Medium; its body is unread). Units element 0 is built with an array base of 0
and is not loaded by the Units load loop, so a saved +0xc of 0 gives a row
whose array base stays 0 unless another route writes it (Medium).
— R2-ENGINE-1013

One save written by another program stores Unit +0xc 0 in 22 of its 260 Unit
records. In the 16 original-made saves of 2026-10-30 (7 of them hold no Unit
record; the 330 Units lie in 9 files of missions 10, 20 and 30), 0 of 330 Units and 0 of
166 Humans store it, and each stored row has values. — R2-SESSION-374

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

Book index i holds the Spell whose ID is i: in the 25 owner files of the three 2026-10-30/31
copies, all 176 nonnull elements of the 38 present book records (11 positions repeated across
consecutive saves) have ID equal to index (854 null). The server book getter stores the last
fetched nonnull index in book DWORD +0x18. In 34 of the 38 books +0x18 is 0 or one of the book's
own IDs, which fits that store (Medium); in 4 books it holds other values and its meaning there is
Unknown. The slot masks (client +0x148, +0x150, object +0x18, +0x1c) are not book fields and the
SAVE body serializes no mask (Medium). On a unit update whose flags hold 0x100000 the server passes
the book's row mask to a writer that was not read, and the client converts a stored mask to slot
bits. REG `SpellBook/Pressed` holds panel +0x60: the engine writes a slot index 0..23 or -1 there,
not a spell ID (High for the writes); LOAD restores the saved value through an unread getter with
no range test (Medium), and no owner save's `Pressed` value was read. — R2-SESSION-318,
R2-ENGINE-766, R2-ENGINE-768

In a mission F4..F12 (0x73..0x7b) with none of Ctrl or Shift, with Ctrl, and with Shift reach the
view key routine's arm `L2.04354` through its first key table (F10 through the Alt routine
`L2.01073`); the Alt forms of F4..F9, F11 and F12 are not forwarded by the bodies read.
— R2-ENGINE-1178, R2-ENGINE-1181

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
from list +0x3a4: a 0x254 or 0x208 byte object (what it is, Medium) and then
a 0x1e4 byte object keyed word entry +0x3c + 0x6000. The 0x1e4 object is the client object of
the map-built unit with that +4. The constructor beyond these loops is
unread. — R2-ENGINE-616, R2-ENGINE-615, R2-ENGINE-951

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

In ROM2 EN (RU bodies unread), base +4 is the structure key. The map build
stores +8 = type-4 ID and +4 = +8 + 0x6000; the client's static object holds
the same sum from the map. LOAD reads +4 from the save with no value test;
no recompute was found (Medium: the object builder is unread). Of the fields
named above, LOAD rebuilds Building +0x3c, base +0x18 and the link keys and
enters the self key in the session; it keeps +4 as stored. Message 0x24
finds the clicked structure by +4. — R2-ENGINE-857, R2-ENGINE-617,
R2-ENGINE-618

Base +4 is also the unit key, in both roots. The map build stores +8 = ALM
unit word +0x3c and +4 = +8 + 0x6000. A unit created later takes the lowest
clear ID of a 0x6000-bit map whose ID 0 is reserved, so 1..0x5fff. LOAD keeps
the saved +4 (Medium). LOAD marks a nonzero loaded +4 below 0x6000 in that
map through the base serializer (R2-ENGINE-857; Medium that every loaded unit
reaches it). — R2-ENGINE-950, R2-ENGINE-953, R2-ENGINE-857

After LOAD the join (Medium that it serves LOAD) sends message 6 and then
every registry unit. Message 6
gives each ALM unit entry a client object under key +0x3c + 0x6000 with
object bit +0x1b8 0x80 set, also for an entry the server skipped. A unit
message carries +4 at word +0x12 and the client looks it up. For keys of
0x6000 or more the send-all mask holds no create bit, so the unit binds to its
map object: field 1 sets object +0x104 from unit +0x94 (the shown health,
Medium) and field 0x10 clears bit 0x80. Such a key with no map object gets no
client object from send-all (whether it is drawn is Unknown). A key below
0x6000 creates its object. — R2-ENGINE-951, R2-ENGINE-953

Sight messages carry +4 at word +0xa: 0x74 sets and 0x69 clears object bit
0x80 (read as hidden, Medium). With Session +0x74 nonzero a player whose
WORD +0x30 bit is clear in unit +0x1a4 (read as cannot see, Medium) receives
0x74 and the unit's other fields wait in unit +0x114. Message 0x6a removes the object by +4. A campaign keeps Session +0x74
zero (Medium), so sight does not gate these messages there. The body that
detects a sight change is Unknown. — R2-ENGINE-952, R2-ENGINE-953

In the 16 original-made saves every unit outside players.0 and in the dead
list stores +4 = +8 + 0x6000; the 21 units of players.0, the hero's player,
store 1, 2 or 3, the hero 1. No unit of the hero's player carries a map key.
— R2-SESSION-360

The five original mission-10 files store +4 = +8 + 0x6000 in all 16
structures (Pointer 24618). game1000 and game1236 store 2..17, and the
original's resaves game0049 and game0048 write 2..17 back. With such a file
loaded, a sign click's lookup returns 0 and no order, action or text box
follows; that composed path is Medium. — R2-SESSION-344, R2-ENGINE-858

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
client object list holds it; that list is keyed by unit +4, which LOAD keeps
from the save (Medium). Pause state and the camera meaning of View
are Unknown. — R2-SESSION-244, R2-SESSION-245, R2-ENGINE-621,
R2-ENGINE-951

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

The sources and LOAD handling of terrain raw60 and session raw10374 are
assigned under Terrain and session. Still open there: plane zero of an
entry cell after LOAD, +0x03 users through a record pointer, effect and
event bytes (zero in all nine mission saves), readers of session bytes
5048..5447 and 10356..10357, and computed-index register writes.
— R2-ENGINE-759, R2-SESSION-311, R2-ENGINE-760, R2-SESSION-313

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

A unit under Stone Curse holds an Effect in its +0x20 references: base u16
+0xc and u16 +0xe, u8 tag +0x3c, u8 flags +0x3d, u32 +0x40 holding the u16
values +0x40 and +0x42, and a trailing u16 +0xc. Its u32 +0x144 and the raw
184 command block, whose +9 is 4 once the hold is installed, are stored with
it. The field list is High. The values are predicted from code, not read in
a save (Medium): +0xc 18, +0xe 44, tag 0, the timed flag, bit 18 (0x40000)
of +0x144, and the remaining duration in the high word +0x42 of the u32.
LOAD restores these bytes. That the countdown resumes from the stored
duration and that the client mark returns is Medium; no save holding the
effect was read. — R2-SESSION-328

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

### Unit Effect marks

A unit Effect's base +0xe holds the mark identifier 2*(+0xc)+8 that the
0x88 message carries; u16 +0x42 is the remaining duration. Apply gives the
Effect a timed flag in byte +0x3d (1 duration, 2 continuous) for spell rows
4, 6, 8, 12..15, 18..21 and 27..29, so, of the arms read, only these
Effects carry a mark to attach; the Heal and Drain Life arms build no
Effect, and their marks 56 and 60 are set on the client by the cast object.
The server expiry decrements +0x42 unless it is above 9600 and sends
0x89 at 0; for a continuous Effect it also calls slot +0x40 at every eighth
step. Which duration an attached Poison Cloud, Light or Darkness Effect
holds, and whether the single-target delivery object attaches its Effect,
are Unknown. In the
EN saves game0004..0006 the hero holds Shield (+0xc 27, +0xe 62, durations
563, 302 and 50) and, in game0005, Haste (21, 50, duration 196); a player-7
unit holds Shield. Its u16 +0x40 is 5 for the hero's Shield, 14 for player 7's
and 2 for Haste; its meaning was not read. Each holder's unit base +0x18 is 2
and Player 0 +0x30 is 2. The client marks and overlay list are not stored. After
LOAD a held mark returns, of the writers read, only through a new 0x88 (the
effect update arms also write marks, but only for a live effect object, and
none is restored; of the 38 EN sites that add the mark list's offset, the
others, with four in functions of an unread class, build, free, look up or
read the list, and the opcode 0x6c arm passes it to an unread routine);
with the saved +0x18 and +0x30 the notify resend test fails for player 0
unless LOAD clears it, and whether the overlay reappears is Unknown. — R2-SESSION-273, R2-SESSION-274, R2-ENGINE-695,
R2-ENGINE-697, R2-ENGINE-700, R2-ENGINE-701

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

## Template keys and the original client health divisor

A saved HPmax word alone does not establish the original client maximum
after LOAD. The EN Unit serializer writes current word +0x14c OR
(DWORD +0x204 shifted 16) OR (DWORD +0x1a0 shifted 24) as one u32.
LOAD restores the low word and Boolean byte bands; this local path does
not derive the template key from saved +0xc. The Unit template initializer
takes +0x14c from Units parameter 55. Complete archive transfer remains a
condition of the composed restoration. — R2-ENGINE-816

The client may reconstruct its maximum through an incoming template key.
The EN sender can omit an explicit maximum when it equals template
parameter 4, while sending +0x14c separately. For the supplied row-92
record, the Medium omitted-maximum explanation requires owner to differ
from recipient, no scaling at the loaded difficulty, and class neither 26
nor 27. When flag 0x20 is clear and 0x80 is set, client reconstruction looks
up that key and copies temporary Unit +0x96 to client +0x108; a separately
received maximum can overwrite it. Key 0 occurs in none of the 180
nonempty installed EN Units rows (61 other serialized rows are empty).
The table at L2.00811 being that Data.bin Units array is inherited, not
read here. The live default-row and packet delivery composition is Medium.
— R2-ENGINE-815, R2-ENGINE-817

One supplied mission-10 document has thirty zero packed actor states.
Five original-made mission-10 documents have 150 actor records, each
with a nonzero low-word key; every saved HPmax in both populations is
nonzero. Twenty-eight supplied records are Units: twelve have positive
HP and sixteen are in the dead list. This is a bounded key census, not
a claim that all twenty-eight reach the same draw. — R2-SESSION-335

The corresponding EN dump shows a signed division by zero at IDIV
L2.04928 with client divisor +0x108 equal to zero (High). The stack words
and installed code place it inside LOAD entry L2.00826 after server reader
L2.00289 returned (Medium: the module list holds no MD5). The inferred
current HP 1229 matches one of the thirty Human/Unit records of the
supplied save: Player 4/Group 0/member 0, Unit archive object 26, ID 32,
row 92. Its packed state at decoded 18166 is 0; the matching original
records have low key 623. The fault object's identity is an inference from
the number 1229, Medium. It assumes saved +0x94 was copied through
R2.0075 into client +0x104 and still held that value at the first draw;
regeneration or a later writer is not excluded. The causal reconstruction
remains Medium because its heap bytes and actual packet flags are absent.
— R2-ENGINE-814, R2-SESSION-336

## Saved attack order without a target

A saved unit order +0x50 = 3 with command +0xc = 0 stops the EN original.
LOAD keeps both values (EN transfer R2-ENGINE-662, Medium): +0x50
is a raw dword, the command block is raw, and the command postload looks up
only nonzero dwords +0xc, +0x10, +0x18, +0x20, +0x28, +0x30 and +0x68. The
read order-3 writers never store order 3 with target 0: the setter itself
stores nothing for a self target (L2.03312 unread), and its one direct
caller reads the target before it. Writer completeness is Medium: disp32,
register, byte and copy stores, indirect calls, stored setter addresses and
writers of +0xc alone were not searched. — R2-ENGINE-937

When the unit's group is dispatched (session +0xbbe8 = 0 needs group
[+0x3c]+0x45 nonzero; mode 0 runs the member loop; either arm of L2.02403
reaches the dispatch), the order-3 arm passes command +0xc to attack
planner L2.03461, which reads target byte +0x4c at L2.05580 before any null
test. A target of 0 reads address 0x4c and the client stops with exception
0xc0000005. The stop is in the session turn after LOAD, not in LOAD
(Medium: the module row holds no MD5; the routines holding L2.05596 and
L2.05598 were not read from their starts). — R2-ENGINE-936

The order-3 arm passes the saved target to the planner with no null test in EN and RU (R2-ENGINE-1004).

The dragonout `game0057.sav` (mission 30) holds exactly one such record
among 107 Unit and Human records: the Human "Igles" of player 0, group 176
(mode 0, +0x45 = 1). Order 3 is at decoded offset 22342; the command block
at decoded 22152 holds +8, +9, +0xc and +0x14 all 0. None of the 496 Unit
and Human records of the 16 original-made saves holds order 3; the
original mission-30 records of the same row and type hold order 12. The
join of the dump's unit to this record is by elimination over the order-3
route only, Medium. — R2-SESSION-356

Unknown: whether the stop was on the first turn after LOAD; RU (planner
L2.01530 not reread); the other ten callers of L2.03461.

## EN mission-10 command inputs and resave preservation

The listed EN command path sends selected client-unit IDs with a local
Player ID. Server lookup uses the Player roster and rejects a found Unit
whose DWORD +0x54 is 16. This is a local admission predicate, Medium for
the composed selection, delivery and attack path; client controlled-set
restoration and additional target gates remain unclosed. — R2-ENGINE-822

The listed SAV programmes transfer Player prefix fields, Unit ID/owner
keys and DWORD +0x50/+0x54/+0x58, raw position 12, command 184 and Group
programme 80 bytes. Player LOAD attaches group members to its roster and
sets their owner; later key fixups and client projection remain Medium or
Unknown. These widths do not establish whole LOAD success. — R2-ENGINE-823

The listed EN mode-1 no-target arm passes the Unit command word +0 as
its prepared move goal when position differs or arrival is incomplete.
A zero post is therefore a conditional zero-goal explanation, Medium;
actual target choice, pathing and the complete mode-18 route remain
unclosed. — R2-ENGINE-824

In the supplied mission-10 game0029 and original resaves game0038/game0039,
Gildarius Unit state50/state54/state58 is 0/0/0 and a saved owner-key
match with Player 1 holds, equal at the same parser path in all three files.
The shown state54 check does not explain lost hero
control. Its specific saved-byte cause remains Unknown. Seven other live
groups have mode 1 and all 13 members have command post 0, equal at the
same parser path in game0029, game0038 and game0039. Five
original-made mission-10 comparisons use mode 18 for the corresponding
nonhero groups. Exact measured values are High; restored-state and
behavioural joins are Medium. — R2-SESSION-337

Eight complete mission-10 walks partition 11 decoded sections each.
Both resaves change head, players, dead, buildings, terrain and session;
spellEffects, sacks, extra, worldPresent and sentinels-and-trailer remain
byte-equal to game0029. The census covers every parsed field and archive
record kind, including opaque blocks by hash. Parser-path equality is
High for bytes, Medium for semantic state equivalence after reference,
group or list changes. — R2-SESSION-338

## EN mission-10 sight, map reveal and command timing

The listed EN server builds each client's relation words from the session
matrix and sets bit 8 of word Q from M[Q][P] bit 0x10 and copies any bit 8 already in M[P][Q]; for the own word the stored diagonal 2 has none. The
client replaces its whole relation array with these words. A saved
diagonal of 0x12 gives a Player's own word 0x3a; a saved diagonal of 2
gives 0x22, without the bit 8 that the fog stamp needs. Which server call
delivers the words after LOAD is unread, so the composed post-LOAD state
is Medium. — R2-ENGINE-826

Client message 0xaa value 1 marks every fog cell explored and visible,
borders included, and stops the fog refresh; value 0 clears the flag. The
senders found are the "#show map", "#hide map" and "#victory" chat
commands behind the unsaved cheat flag. The sender set rests on byte
scans and is Medium. — R2-ENGINE-827

The listed routine L2.03404, which looks like a map-dot drawer (Medium), needs the OR of the four fog cells under the object to hold both fog bits and, for
one list tag, owner relation bit 8. The world sprite renderer is not
identified. — R2-ENGINE-829

The listed EN Group mode-18 arm sends a member to its command word +0x9c
while command word +0x16 lies within 10 of the server word +0 restored
from the SAV head second DWORD (30 when its HP differs from max; the absolute value is assumed because L2.04968 is unread). Otherwise it
attacks, patrols or returns to its command post. The trajectory and the
LOAD restore are Medium. — R2-ENGINE-828, R2-SESSION-284

In 13 mission-10 files the five original-made store the diagonal 0x12,
eight cells with bit 0x10 and command +0x16 65500 in every Unit and Human
record. The eight Againrom-written files and their original resaves store
the diagonal 2, no bit 0x10 and command +0x16 0. Head second is 1 in
game1040 and 444 in game1041. — R2-SESSION-339

The original resaves keep the matrix and command +0x16 as read. The fog REG
of the original resaves game0040 and game0039 holds all 6400 cells
explored. Over
game1040 to game0040, 13 of 30 live units move, 11 nearer to cell (0,0);
over game1041 to game0041 none of 14 moves. Both resaves empty the
selection. No stored byte is shown to refuse command of game1040's hero.
— R2-SESSION-340

## EN mission-10 creatures, dead list and resaves

The mission-10 bats, squirrels and Foot_Animated are Units rows 97, 109
and 114 (typeIDs 70, 74 and 75), with 4, 6 and 10 records in each file;
the owner's creature names are tied to them by name (the bats also by the
Darkness spell, R2-ASSET-179): Medium. Of the three original-written files of
the sight run game0044 stores all of them alive, game0045 six
squirrels and four Foot_Animated in the dead list, game0046 all sixteen
there. Why one centipede appears and vanishes after a LOAD is Unknown.
— R2-SESSION-341

A killed unit leaves its Group, its owner's +0x24 list and the walked list
(the unit registry, through R2-ENGINE-772) and enters the corpse list, which the document writes and reads
through that list's slot 0 as a count and objects. In the bodies read,
LOAD keeps those units in the corpse list only; the list's class and the
Player and Group LOAD readers are Medium. — R2-ENGINE-830, R2-ENGINE-831

A bat record stored with command spell slot +0x7c = 0 keeps it through
original resaves (game0038..game0041 are the original's resaves by the
owners' statements), and the listed Unit LOAD body calls no template loader
directly (35 direct targets, 5 indirect calls; callees and the object factory
unread), so the attack planner finds no Darkness slot (Medium). — R2-ENGINE-832

Two of the inputs of the L2.03404 test (SpellEffects records, relation
cells 1,5 and 5,1) hold no tag 0x20 and no bit 0x10 for the bats in 18 files
(cell 5,1 is 1 in eight files, 0 in ten); +0x184, +0x1b8 and fog were not
compared, and the role of L2.03404 as a map-dot drawer is Medium
(R2-ENGINE-829). Which value keeps the bats undrawn is Unknown.
— R2-ENGINE-833

Original resaves keep the Againrom-written creature values: +0x136 0, the
name, bag +0x1c 0 and, for dead records, +0x58 0 and +0x48 10; for the
bats, spell slots 0. — R2-SESSION-342

## Mission-10 group leash and command fields

The Group programme record holds the leash state of a guard group. Each
guard tick first writes +0x24 and +0x28 with the members' truncated mean
cell; +0x2a, +0x2b and +0x2c take the largest member distance from it, the
largest member byte +0xa5 and the largest sum of the two. The leash bound is
byte +0x2d, measured from +0x28. — R2-ENGINE-841

| Offset | Width | Field |
|---|---|---|
| +0x20 | u8 | programme mode; 18 for every nonhero group of the five original mission-10 saves |
| +0x24 | u32 | mean cell, y << 16 \| x |
| +0x28 | u16 | mean cell, y << 8 \| x |
| +0x2d | u8 | leash bound: +0x38 + U{-1, 0, 1} + 4 once +0x30 is 1 |
| +0x30 | u32 | 1 after the first guard tick with a member; 0 before it, and again after a tick with no member |
| +0x34 | u32 | guard ticks since +0x30 last changed |
| +0x38 | u8 | base radius: 10, 11, 12, 15 or 20 in mission 10 |
| +0x45 | u8 | active byte |

The mode-18 setter raises +0x2d and +0x38 together to its argument, or to
20 when the argument is 0 and +0x2d is below a session value; the jitter
draw's divisor is unread, so U{-1, 0, 1} is Medium. In the five original
mission-10 saves, 33 records store +0x30 1 with +0x2d - +0x38 of 3, 4 or 5
and 27 store +0x30 0 with +0x2d = +0x38, every group of the restart save
among them. +0x28 equals the stored members' truncated mean in 52 of 65
records (all 33 with +0x30 = 1). — R2-ENGINE-843, R2-SESSION-343

The Unit command block fields of the walk back and the hit memory:

| Offset | Width | Field |
|---|---|---|
| +0x0 | u16 | post cell |
| +0x8 | u8 | activity: 1 walk, 5 attack, 0xb at the post |
| +0xa | u16 | walk goal cell |
| +0x16 | u16 | stamp: server word +0 when a mode-18 attacker's route search finds no route and its substitute cell is farther from the goal than `L2.01728` (RU and EN); 65500 from the constructor |
| +0x20 | u32 | target; 0 when the leash leaves no enemy |
| +0x54 | u32 | 1 after a hit |
| +0x58 | u16 | the attacker's cell at the hit |
| +0x5a | u8 | 0 at the hit |
| +0x94 | u32 | the attacker |
| +0x98 | u32 | server dword +4 at the hit |
| +0x9c | u16 | the step-away cell beyond the victim |
| +0x9e | u8 | guard B case-5 counter; among the census-traced sites only in case 5, and no read body starts it from 0 |
| +0xa0 | u16 | guard B case-5 walk cell; among the census-traced sites only in case 5 |

A hit writes no target. All 150 unit records of the five original mission-10
saves store +0x16 = 65500, so none holds a stamp; game0006 stores
a group of two that reads as in pursuit, with the hit memory toward the hero's cell.
— R2-ENGINE-845, R2-ENGINE-846, R2-SESSION-343

The same 150 records store +0x9e = 0 and +0xa0 = 0, in the five saves of R2-SESSION-346.
— R2-ENGINE-873, R2-ENGINE-876, R2-SESSION-346

## Weapon Spell, book and the mage bit across LOAD

Unit byte +0x4c bit 4 is the mage bit: EN L2.01136 / RU L2.00044 return it,
EN L2.05310 / RU L2.01517 its negation (High in both clients). A unit with the
bit and a weapon Spell enters actor state 0xd from state 3; a unit without it
casts the weapon Spell only from the strike, after a blow with damage above 0
to a living target (or the Spell id is 2). The tests are High in both clients;
what state 0xd does is Unknown (the handler was not read). — R2-ENGINE-900,
R2-ENGINE-901

In EN only, the attack-order routine and idle support call the book choices and
support only for a mage (High for the gates; Medium for what the choices do,
since the bodies of L2.04141, L2.04143, L2.05314 and L2.05315 are unread). The
EN cast charges mana only to a mage casting without an item. The EN book order
and the server book-mask update contain no bit 4 test (the call L2.03061 in the
book order is unread). Unknown: whether three other mask callers and the
panel's item spells depend on bit 4; what the cast body L2.05342 does and what
a returned 1 at L2.05347 leads to; RU, where these routines are unread.
— R2-ENGINE-902, R2-ENGINE-903, R2-ENGINE-901

In an EN linear-sweep census of +0x4c byte writes, 7 sites in 6 routines set
bit 4: the creature template (reached from unit init only with a name), two
ensure-book routines that act only on a unit with no book, and three that act
when mana max +0x9c is above 0; none is in the Unit LOAD body (LOAD stores the
byte through a helper, so a displacement census would not show it either way).
In an EN linear-sweep census of 88 +0x84 store candidates, the classified
Weapon writers are the Weapon constructors, destructor, copy constructor (a new
Spell from the stored id) and two setters; none is in a LOAD body, and the
Weapon LOAD stores the Spell reference through the reference reader L2.05317,
which that census cannot see. Both censuses are Medium for completeness.
Unknown: the callers of L2.00201, R2.0088, L2.01818 and L2.01385; the 58 word
and dword +0x4c candidates and writes through L2.03713 and memcpy; the RU
censuses. — R2-ENGINE-905, R2-ENGINE-906

| game1200 record | Units row | +0x4c | Weapon Spell | Book size | Book IDs |
|---|---|---:|---|---:|---|
| players.3.group.0.members.0 | 178 Necro_Male1 | 2 | 26, range 7, cost 20 | 28 | 1, 26, 27 |
| players.3.group.0.members.1 | 178 Necro_Male1 | 2 | 26, range 7, cost 20 | 28 | 1, 26, 27 |
| players.3.group.5.members.0 | 180 Necro_Leader1 | 2 | 26, range 7, cost 20 | 28 | 1, 10, 26, 27 |

The table is the Againrom-written game1200: three living necromancers of
computer-owned Player 3, typeID +0xe 99 (Medium: a necromancer is identified
by +0xe 99 and the +0xc row; the 8 records with +0xe 99 include dead.33). The
four original-made row-178 records, two in each of game0012 and game9999, store the same Spell and book IDs with +0x4c
14 or 6 (bit 4 set), mana 150/150 and book +0x18 263 or 3. — R2-SESSION-350

In four original resaves of mission 10, all 64 squirrel and Foot_Animated
records keep the Againrom-written size-1 book and +0x4c 0; all store mana max
0 (Medium that LOAD and play keep a stored book in general: one map, one
creature set; none of these units stores +0x9c above 0, so the bit-4 setters
gated on it were not tested). — R2-SESSION-351

So, in the EN census and routes (RU not censused), after LOAD of game1200 the
original keeps each necromancer's Drain Life staff Spell and book, but with bit
4 clear the necromancer takes no state-0xd route and no book choice. The only
read route that reaches its weapon Spell is the strike; what the cast does
after its return is unread. The stored +0x4c decides it (Medium: three
mana-gated bit-4 routines have unread callers, L2.00201 is bypassed by a unit
with a book, and state 0xd and the four book-choice callees are unread). A cold
LOAD of game1200 in the original, watching for ranged casts and resaving to
compare +0x4c, would settle it and has not been run. — R2-ENGINE-907

### Human modifier to-hit and damage with held effects

The Human raw64 block at object +0xd4 (UD4) contains modifier to-hit word
at block +0x12/+0x13 (object +0xe6), damage base byte +0x20 (+0xf4) and
spread byte +0x21 (+0xf5). These differ from live attack +0xa6/+0xb4/+0xb5.
The read Unit transfer writes or reads the block before held-item references.
RU transfer and local order are High; the composed equipment-to-SAV result
and EN helper semantics are Medium. Whole-session last-writer and post-LOAD
order are Unknown. The coordinates are arithmetic on the raw 64-byte transfer.
— R2-SESSION-358

For descriptor parameter 5 below 10, the read equip bodies add Weapon
to-hit/base/spread to those modifier fields, modulo 65536/256/256. Sources
are RU +0x5a/+0x68/+0x69 and EN +0x56/+0x64/+0x65. Equip calls recompute,
then child start; removal calls child end before subtracting the Weapon
fields and recomputing. Call order is High in RU; in EN the slot +0x58
recompute and Weapon slots are bound by RU correspondence (Medium).
Complete child traversal, the parameter binding and all-event coverage are
Medium; map construction and initial modifier values are Unknown.
— R2-ENGINE-947

In both read Effect apply bodies, tag 49 shares the arm of tags 13 and 43
and adds the scaled operand to modifier base modulo 256. That arm writes
neither modifier to-hit nor spread. Flags 1/2/4 select signed16 +0x40, otherwise dword +0x40;
the generic start/end route supplies multiplier +1/-1 and excludes Effect
+0xc values 6, 14 and 15. The local branches are High. Other consumers of
tag 49 are Unknown. — R2-ENGINE-946

Recompute builds live attack from attributes and skills, then folds the
modifier block into it; it does not store the derived live base at the
modifier-base store listed above. Local fold stores are High; formulas
containing unread power/conversion helpers remain Medium. The earlier
weapon formula is amended for child and pre-existing modifier terms.
R2-ENGINE-479's formula stays Medium / Unknown.
— R2-ENGINE-948, R2-ENGINE-479 (amended)

The complete apply bodies are not identical: EN has six extra instructions
that clear signed Human word +0xd8 when greater than 24. The other 870
instructions match RU after absolute-address masking. This local difference
is High; callee equivalence is Medium and its full scheduling/player effect
is Unknown. — R2-ENGINE-949

### Effect renewal and construction boundary

By the code read in both roots, attach keeps the old same-ID object, tag,
flags and pointer44,
then independently maximizes signed magnitude and zero-extended remaining
duration between old end/start. A new copy transfers the two words with one
dword store. Direct bodies are High; base-copy, traversal and notification
composition remain Medium. — R2-ENGINE-1164

By the code read in both roots, entry is flags&3. Timer duration0 wraps to
65535 before a later >9600 return;
continuous start runs before that limit check and continuous expiry omits
end. The timer marks an ended record; cleanup/unlink is unread. These direct
operations are High in both roots. — R2-ENGINE-1165

A parsed timed item child retaining constructor ID0 meets another ID0 even
with a different tag; renewal uses the old tag. This source composition is
Medium: blank and item-use bodies are complete, parser prefixes and helpers
are incomplete, and shipped timed-potion rows were not read. — R2-ENGINE-1169

Item child-start/end call slots40/44 directly. Additive item/spell modifier
algebra is conditional; complete recomputation and independence are Unknown.
The read construction/renewal/timer stores do not form a complete writer
census, and no new save/load or message-writer proof is supplied.
— R2-ENGINE-1170, R2-ENGINE-1171

## Health and mana restoration state

The RU Unit serializer's fourteen u16 at +0x84..+0x9e cover health +0x94,
its maximum +0x96, the health divisor +0x98, mana +0x9a, its maximum +0x9c
and the mana divisor +0x9e. Its u8 +0xa2 and +0xa3 are the two remainders
in hundredths, its u32 +0x138 is the stamp that sets the factor K, and its
64 raw bytes from +0xd4 cover the modifier words +0xde and +0xe2 without
naming them. Whether a save holds Session +0 and +4, which set the health
phase and K, is Unknown. — R2-ENGINE-1209
