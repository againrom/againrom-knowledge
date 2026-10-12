<a id="rom2-alm-map-m7r0--identity-survey"></a>

# ROM2 ALM maps (`M7R\0`)

ROM2 ALM extends the typed ROM1 map container with record types 10–12 and
versioned metadata/placement fields. The preserved versions are 1300 and
1600. The primary loader accepts versions through 1600 and requires at least
three records. — R2-ASSET-002, R2-ASSET-003, R2-ASSET-020

<a id="result"></a>

## File and record headers

| Offset | File header (20 bytes) | Record header (20 bytes) |
|---:|---|---|
| 0x00 | Magic `M7R\0` | u32 tag |
| 0x04 | u32 header length, 20 | u32 header length, 20 |
| 0x08 | u32 dataSize | u32 declared payloadSize |
| 0x0c | u32 recordCount | u32 typeId |
| 0x10 | u32 formatVersion | f32 record scalar |

Integers are little-endian. Installed dataSize is `4*W*H+72`, with W/H at
type 0 payload `+0/+4`. Installed maps carry the types once each in order
`0,1,2,3,5,11,4,9,8,6,7,10,12`. This order and count are installed values;
the loader's switch has cases 0..12 and a non-rejecting default.
— R2-ASSET-002, R2-ASSET-003, R2-ASSET-017, R2-SESSION-010

## Payload extents

| Type | Extent at versions 1300/1600 | Count source or rule |
|---:|---|---|
| 0 | 660 actual bytes; declared size 644 | Fixed/versioned metadata programme |
| 1 | 2×W×H | u16 Tiles cells |
| 2, 3 | W×H each | Byte planes |
| 4 | 20 or 28 bytes per element | Type0 object count; 28 when `kind==0x21 \|\| (kind&0x1000000)!=0` |
| 5 | 76×nPlayers | Type0 player count |
| 6 | 48×nUnits | Type0 unit count |
| 7 | ROM1 three-counted-array grammar | Counts in this payload |
| 8 | ROM1 variable loot-record grammar | Count in type 0; modern head 20 bytes |
| 9 | ROM1 counted caster-record grammar | Count in this payload |
| 10 | 16×meta[+0x30] | Flat array at map+0x310 |
| 11 | 12×meta[+0x34]+84×meta[+0x38]+12×meta[+0x3c] | Arrays at map+0x338/+0x324/+0x34c respectively |
| 12 | 28+28×meta[+0x40] | Fixed head at map+0x374; array at map+0x360 |

Type4 `kind` is at element `+0x08`. Its bit-24 extension test is not
version-gated; do not replace it with a low-byte comparison. Type5 comprises
four leading fields of widths 4,4,4,32 and sixteen u16 diplomacy entries:
`4+4+4+32+16*2=76`.
— R2-ASSET-018, R2-ASSET-019, R2-ASSET-020, R2-ASSET-021,
R2-ASSET-022, R2-ASSET-023

The campaign type7 grammar is exactly consumed for all 46 maps per preserved
root. Full-word operation values, compilation and the supported added arms are
specified in [Campaign scripts](script.md). — R2-ENGINE-041

<a id="unit-record-type-6"></a>

## Unit record (type 6)

Each 48-byte record becomes a 0x54-byte row in the map's unit list. The
row slot is where the program keeps the field; readers are named where
the read bodies show one. The table holds for versions 1300 and 1600, the
two in the installed campaign maps; the block read below version 0x44c
uses other fields.

| Offset | Width | Row slot | Use |
|---:|---|---:|---|
| 0x00, 0x04 | i32 | +0x00, +0x04 | x, y; shifted right by 8 into the placement call (the 1/256 tile unit is Medium) |
| 0x08 | i16 | +0x08 | read by the mission-load row loop, as a lookup index into a collection whose record chooses the object build; no actor-builder read (Medium) |
| 0x0a | u16 | +0x0c | no reader found |
| 0x0c | u32 | +0x48 | flags; bit 0x10 selects the person arm |
| 0x10 | u32 | +0x50 | bits 1, 2, 4, 8 set actor +0x194, +0x19c, +0x198, +0x1a0; read at version 0x47e and above. Bits 1, 2 and 4 put the unit in Inn offer populations; bit 8 removes its death drop and the experience a human attacker is paid for it (Medium) |
| 0x14 | u32 | +0x10 | definition id; read at version 0x3db and above |
| 0x18 | u32 | +0x14 | owner, low word only |
| 0x1c | u32 | +0x44 | 1-based type-8 index; that type-8 element takes the record's 0x28 low word |
| 0x20 | u32 | +0x1c | no reader found |
| 0x24, 0x26 | i16 | +0x20, +0x22 | actor words +0x94 and +0x96; -1 keeps the default |
| 0x28 | u32 | +0x3c | low word read; two bytes below version 0x456, four from 0x456 |
| 0x2c | u32 | +0x40 | group number; the map keeps the maximum |

"No reader found" is bounded to the actor builder and the mission-load
row loop. Installed values in the campaign maps: 0x08 1..107, 0x0a 1..6,
0x0c {0, 16}, 0x10 {0, 8}, 0x1c 0..4, 0x20 0..14 (10 in nearly every
record), 0x2c 0..38. In the root `.alm` maps: 0x0a 0..7, 0x0c {0, 4, 16,
20}, 0x10 {0, 1, 2, 3, 4, 5, 8, 9, 11, 12} in EN and also 10 in RU (the builder tests each bit separately; R2-ENGINE-513), 0x18 2..15 with high
word 0, 0x1c 0..58, 0x20 0..14, 0x2c 0..555.
— R2-ASSET-119, R2-ASSET-122, R2-ASSET-132, R2-ENGINE-038, R2-ENGINE-513

Actor +0x1a0 also receives 1 from two unique-spell arms and, when
Session+0x74 is set, from the map build for the units of owners whose
diplomacy toward the owner named "Self" has bits 0..2 clear. An actor with
+0x1a0 set drops no bag or gold at death and pays no experience to a human
attacker (R2-ENGINE-480; Medium here; a creature attacker's award was not
read for the flag). The unit copy omits the flag (Medium: its head is
unread); the save keeps it (Medium). Actor +0x194, +0x19c and +0x198 select
Inn offer blocks 2, 4 and 6 (Medium: the receiver of the offer generator is
the Inn by address inference); RU and EN agree in those blocks. — R2-ENGINE-535, R2-ENGINE-537

<a id="group-and-object-records-types-10-and-11"></a>

## Group and object records (types 10 and 11)

Type 10 entries are four i32, keyed on a unit group number:

| Offset | Field |
|---:|---|
| 0x00 | group number (unit record 0x2c) |
| 0x04 | percent, scaled by a global setting; the product is the reload value of the group object's countdown (Medium) |
| 0x08 | flags 1, 2, 4, 8 |
| 0x0c | position in the type-7 first array; the count of nodes before it with dword +0x40 below 0x10002 is passed to the routine `L2.02490` (Medium) |

The program runs the type-10 pass when type-0 payload dword 0x84 exceeds 1
(gate body High; "only" Medium, other writers of the gate are open). Then
each unit group of an owner with a nonzero +0x2c whose dword +0xc is not 0
(read by `L2.02502`; that it is the member count is Medium) gets a group
object, with defaults when no entry matches; every matching entry applies
in file order. The 46 `scenario.res` campaign maps per root have dword
0x84 = 1 and no type-10 entry. The 13 EN and 37 RU root `.alm` maps have
dword 0x84 of 4..16, above 1, and type-10 entries; the gate opens for them
if they are built by this routine (Medium: the loader of a root map is
unread). Type-0 payload dword 0x88 is
stored in the map and the session; its meaning is Unknown.
— R2-ASSET-120, R2-ASSET-122, R2-ASSET-130, R2-ENGINE-509, R2-ENGINE-510,
R2-ENGINE-538

Flag 4 and flag 8 also put the group in Inn offer blocks 3 and 5 through
group +0x48 and +0x4c. In block 3 RU skips groups of the Inn's own owner
and EN does not. The object re-creates members only when the group's
dword +0xc is 0 or below (Medium, R2-ENGINE-510: no slot-5 body was read
for this page); RU places them over the member spread, EN at
radius 0 (Medium). — R2-ENGINE-537, R2-ENGINE-538

Type 11 entries key on a type-4 object id, the sign-extended type-4 file
word 0x12:

| Array | Width | Fields | Receiving type-4 kinds |
|---|---:|---|---|
| 11a | 12 | id, flags (2, 4, 0x10, 0x20, 0x40), word used under flag 2 | 0x43..0x45, 0x63..0x65, 0x6f..0x71; last match wins |
| 11b | 84 | id, then per category i (0..3): mask 0x04+4i, minimum price 0x14+4i, maximum price 0x24+4i, draws 0x34+4i, quantity bound 0x44+4i | 0x22, 0x23, 0x5d..0x5f, 0x69..0x6b; first match |
| 11c | 12 | id, nonzero gate, position in the type-7 first array | 0x46..0x4b; last nonzero match wins |

The 46 campaign maps per root carry no 11a or 11b entry and 40 11c
entries in 10 maps. The root `.alm` maps carry 376 EN and 477 RU 11a
entries, all on Inn kinds, and 452 EN and 555 RU 11b entries, all on Shop
kinds except five per root (entries 1, 4, 11, 16 and 19, all with id 265)
that name no type-4 record; their 11c entries
all have gate 0. 11a flag 1 occurs there and no field receives it; under
flag 2 the word is 3661..3680. — R2-ASSET-121, R2-ASSET-122,
R2-ASSET-130, R2-ASSET-132, R2-ENGINE-512

An Inn uses its 11a fields to build a player's offer list: flags 0x20,
0x40 and 0x10 admit unit class ranges, flag 4 enables one block and the
flag-2 word is the argument of another (Medium: the receiver of the
offer generator is the Inn by address inference); what an offer is to the
player is Unknown, and so are other readers of Inn +0xc4..+0xd4 and Shop
+0x70..+0xbf, for which no displacement scan ran. When Session+0x74 is set, a Shop fill takes its four category
entries from the Shop's 11b fields, in the order of the Scenario.dll
ordinal-15 record; otherwise from that record. — R2-ENGINE-536,
R2-ASSET-131, R2-ENGINE-539

The kinds that receive type-11 fields build structure classes: 0x46..0x4b
Pointer (0x70 bytes), 0x22, 0x23, 0x5d..0x5f, 0x69..0x6b Shop (0xc4) and
0x43..0x45, 0x63..0x65, 0x6f..0x71 Inn (0xec), each a subclass of the 0x6c
Building class. The Data.bin Buildings entry at the same index carries the
matching name, "Pointer 1".."Pointer 6", "Shop", "Shop 2" and the
Multiplayer shops and inns (that the kind indexes Buildings is Medium).
— R2-ENGINE-517, R2-ASSET-124

A Pointer stores the 11c count at +0x6c; its Serialize slot transfers the
dword (that this writes a save is Medium). An actor step with health above
0 and unit state 15 looks up the first building within distance 2 of the
point (x-1, y-1) of the unit position (the origin and "first" are Medium).
When that building is a Pointer with +0x6c above 0, the arm calls the
action executor with +0x6c on each such step and stores nothing; a latch
inside the executor is excluded only for operation 2. For a Shop and an Inn
the arm calls `L2.02566` with 0x83 and 0x84 (what the call does is unread).
One order handler sets state 15 when `L2.02567` returns nonzero (arrival in
range is Medium) and calls `L2.01576`, which sets [r+0x50] to 1, on its next
call in state 15. Compiled slots follow node order; that a slot advances
only on a successful compile is Medium, so +0x6c names the 11c node only if
no earlier node below 0x10002 failed. — R2-ENGINE-517, R2-ENGINE-518,
R2-ENGINE-519, R2-ENGINE-520

The 46 campaign maps per root hold 13 Pointers and no Shop or Inn. The 35
gate-1 11c entries name 12 Pointers; for each, the action node at index
position exists and has operation 2. That the Pointer selects that node at
run time is Medium; the thirteenth Pointer has only a gate-0 entry. The
root `.alm` maps were not censused. Type-4 kinds 0x46..0x4b occur in these
maps, so the two-value kind vocabulary of R2-ASSET-019 holds only for the
files it described. — R2-ASSET-123, R2-ENGINE-520, R2-ASSET-019

The Shop and Inn slot-5 bodies read in RU do not read the 11b or 11a
fields; no caller of slot 5 was read, so when they run is Unknown.
— R2-ENGINE-521

<a id="music-areas-type-12"></a>

## Music areas (type 12)

Type 12 is one head record and `meta[+0x40]` area records, each seven
little-endian i32:

| Offset | Field |
|---:|---|
| 0x00 | x, tiles |
| 0x04 | y, tiles |
| 0x08 | radius, tiles |
| 0x0c..0x18 | themes 0..3: index into the mission music list, -1 for none |

The loader admits the head only when its theme 0 is at least 0. A record
at (0,0) is the default for a position outside every area, so on a map
with an admitted head every mission track is an area or head theme; a
record with all four themes -1 is ignored. Area selection and the theme draw are in
[ROM2 music](../rom2-music/format.md). — R2-ASSET-084, R2-ENGINE-336

All 46 campaign maps per root have a (0,0,0) head with 1..4 themes and
2..33 areas (470 per root, equal in EN and RU). Two EN and six RU root
maps have areas and a head with themes -1; the other root maps have
neither. Theme values are -1..16. — R2-ASSET-084

## Version conditions

| Case | Condition | Transfer |
|---:|---|---|
| 0 | version>0x47d | One extra u32 |
| 0 | version>0x4cd | Three extra u32 |
| 0 | version>0x513 | One extra u32 |
| 0 | version>0x487 | Two extra u32 |
| 6 | version at or above 0x47e | Extra u32; otherwise zero in memory |
| 6 | version at or above 0x3db | Extra u32; otherwise zero in memory |
| 6 | version<0x44c | Legacy block, including nested 0x3b6/0x3d8 gates; otherwise omitted |
| 6 | version at or above 0x456 | Four-byte final field instead of the two-byte form |
| 8, 9 | version>0x3dd | One extra field on each named case |

Type0's ordered width sum at the preserved versions is
`48+4+12+4+64+4+4+8+512=660`; when none of its gates fire it is 632.
Type6 is `36+4+4+0+4=48`; the ROM1 version-990 branch is
`36+0+4+28+2=70`. Whether the case 6 0x3db gate and ROM1's published 0x3da
gate refer to the same field remains Unknown. — R2-ASSET-024

## Decode sequence

1. Read the file header and its magic/count/version gates.
2. Walk record headers in file order. For type 0 at versions 1300/1600,
   consume the complete 660-byte programme even though payloadSize says 644.
   Advancing only by the declared size misplaces the next header by 16 bytes.
3. Retain metadata dimensions and all count fields before their dependent cases.
4. Use each record's own case grammar and version conditions. For the
   preserved versions, the other record payloads use their declared extents.
   The loader itself reads types 10, 11 and 12 by their type-0 counts in
   signed loops and never consults the payload size: a count of 0 or below
   reads nothing, and a count that disagrees with the size misframes the
   records after it (Medium). Below versions 0x47e, 0x4ce and 0x514 the
   respective counts are 0. — R2-ASSET-129
5. Check each payload and complete record-chain endpoint against the input.

The ROM1 version ceiling of 1001 rejects the preserved ROM2 versions. Shared
field shapes do not imply ROM1 load acceptance. — R2-ASSET-002,
R2-ASSET-003, R2-ASSET-024

<a id="not-yet-surveyed"></a>

## Authoring limits and Unknowns

Type 10 and 11 fields are decoded to the program fields that receive
them above. The 11c count reader is the state-15 arm (R2-ENGINE-518; that it
is the only reader is Medium). Which order code puts a unit in state 15 is
Unknown (R2-ENGINE-519); so are what an Inn offer is and which event
builds the offer list (R2-ENGINE-536). So is the reader of unit-record
0x0a and 0x20 outside the two bodies read. EN and RU differ in one type-10 store: the
RU group object stores the group's spread at +0xbc, the EN object keeps 0
there (R2-ENGINE-511). Type 12 is decoded above (R2-ASSET-084). The writer's reason
for declaring 644 bytes while the metadata programme transfers 660 is also
Unknown; no complete original-compatible ROM2 ALM writer follows from this
loader grammar. — R2-ASSET-023, R2-ASSET-024

Installed W/H maxima are 256, object-count maximum 324 and unit-count maximum
785; these are not format limits. The player record has sixteen diplomacy
slots, while the installed maximum player count is 15. R2-ASSET-027 is
partially retracted only on attribution of the ROM1 comparison maxima:
478 objects belong to EN Beast.ALM, and 1815 units to EN Horror.alm.
Those comparison values are not admission limits either.

Complete ROM2 field meanings, arbitrary malformed-map acceptance, a ROM1
hard player-count rejection above its installed maximum of nine, and the
type-ID-to-definition consumer link remain Unknown. See
[ROM2 Data.bin](../rom2-databin/format.md).

## Player records and native activity fields

Native activity is runtime state. Its recount rebuilds coverage before the
mark pass and includes zero-class owners' nonnull +0x38 actors as extra
producers. Authored Player type is not proved equivalent to that runtime
class; every list membership and +0x38 producer remains Unknown.
— R2-ENGINE-087

The selected EN/RU owner-record builder walks argument+0x28 records, copies
a name and other fields, and separately sets owner DWORD+0x2c. It becomes
zero for record ordinal 1 under Session+0x74==0; the observed condition is
not a direct copy of an authored type DWORD. Registration derives owner
u16+0x32 through +0x30 from a runtime owner key, and archive loading copies
loaded +0x30 to +0x32. Actor DWORD+0x14 supplies the owner receiver.
— R2-ENGINE-071, R2-ENGINE-072

The existing player-record grammar supplies no proved authored owner-type
or activity-mask default. Native table-to-Data.bin identity, NPC registry
joins, human/AI classification and complete control-record mappings remain
Unknown for the selected producer chains. — R2-ENGINE-072

On the selected native new-owner path, owner record DWORD+4 is copied to
runtime owner DWORD+8. Map actor lookup compares signed actor-row WORD+0x14
to that field; a global-head match selects the first owner used by the literal
consumer. This is a conditional native receiver join, not a complete stored
record-type mapping or an ordinary-start member population. — R2-ENGINE-155

Actual first accepted row, entering owner actor contents/order, uniqueness,
final numeric identity and ordinary-start literal success remain Unknown at
outer caller predicates, lookup-next/world results, receiver aliasing and
callback effects.
— R2-ENGINE-158

The selected initialization does not establish the numeric Session DWORD+0x174
used by the map-actor insertion predicate. Its local stores, pointer-mediated
write, unexpanded eligible callees and opaque/indirect calls leave field
production and runtime aliases Unknown. — R2-SESSION-068

## Blizzard placement eligibility

The .alm-named campaign archive population has 46 parsed maps per root, with
4212 EN and 4234 RU type-6 placements. Map hashes agree for 42 pairs; 110.alm,
42.alm, 52.alm and 74.alm differ. Only 110.alm differs in placement count,
259 versus 281. These structural counts do not establish Blizzard casters.
The placement spellbook offset/bit and creature/item joins remain Unknown;
loose maps and magic-based discovery were outside this population
(R2-SESSION-162).
