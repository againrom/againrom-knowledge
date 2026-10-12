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
| 0x10 | u32 | +0x50 | bits 1, 2, 4, 8 set four actor fields; read at version 0x47e and above |
| 0x14 | u32 | +0x10 | definition id; read at version 0x3db and above |
| 0x18 | u32 | +0x14 | owner, low word only |
| 0x1c | u32 | +0x44 | 1-based type-8 index; that type-8 element takes the record's 0x28 low word |
| 0x20 | u32 | +0x1c | no reader found |
| 0x24, 0x26 | i16 | +0x20, +0x22 | actor words +0x94 and +0x96; -1 keeps the default |
| 0x28 | u32 | +0x3c | low word read; two bytes below version 0x456, four from 0x456 |
| 0x2c | u32 | +0x40 | group number; the map keeps the maximum |

"No reader found" is bounded to the actor builder and the mission-load
row loop. Installed values: 0x08 1..107, 0x0a 1..6, 0x0c {0, 16}, 0x10
{0, 8}, 0x1c 0..4, 0x20 0..14 (10 in nearly every record), 0x2c 0..38.
— R2-ASSET-119, R2-ASSET-122, R2-ENGINE-038, R2-ENGINE-513

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
each unit group of an owner with a nonzero +0x2c, for which a guard
routine (`L2.02502`, unread) returns nonzero, gets a group object, with
defaults when no entry matches; every matching entry applies in file
order. The 46 `scenario.res` campaign maps per root have dword 0x84 = 1
and no type-10 entry; the root `.alm` maps were not read. Type-0 payload
dword 0x88 is stored in the map and the session; its meaning is Unknown.
— R2-ASSET-120, R2-ASSET-122, R2-ENGINE-509, R2-ENGINE-510

Type 11 entries key on a type-4 object id, the sign-extended type-4 file
word 0x12:

| Array | Width | Fields | Receiving type-4 kinds |
|---|---:|---|---|
| 11a | 12 | id, flags (2, 4, 0x10, 0x20, 0x40), word used under flag 2 | 0x43..0x45, 0x63..0x65, 0x6f..0x71; last match wins |
| 11b | 84 | id, five runs of four i32 forming four groups of five | 0x22, 0x23, 0x5d..0x5f, 0x69..0x6b; first match |
| 11c | 12 | id, nonzero gate, position in the type-7 first array | 0x46..0x4b; last nonzero match wins |

The 46 campaign maps per root carry no 11a or 11b entry and 40 11c
entries in 10 maps. — R2-ASSET-121, R2-ASSET-122, R2-ENGINE-512

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
5. Check each payload and complete record-chain endpoint against the input.

The ROM1 version ceiling of 1001 rejects the preserved ROM2 versions. Shared
field shapes do not imply ROM1 load acceptance. — R2-ASSET-002,
R2-ASSET-003, R2-ASSET-024

<a id="not-yet-surveyed"></a>

## Authoring limits and Unknowns

Type 10 and 11 fields are decoded to the program fields that receive
them above; what the type-11 receiving objects do with those fields during
a mission is Unknown, and so is the reader of unit-record 0x0a and 0x20
outside the two bodies read. EN and RU differ in one type-10 store: the
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
