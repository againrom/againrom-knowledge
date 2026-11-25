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
| 9 | Counted elements, grammar in [Type-9 records with mode mask 4](#type-9-records-with-mode-mask-4) | Count in this payload |
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

The installed script census covers 142 maps: EN 13 loose and 46 campaign,
RU 37 loose and 46 campaign. Campaign actions/checks agree at 1334/986 per
root; triggers differ, EN 715 / RU 714. Both campaigns use 29 ordinary action
IDs and 19 ordinary check IDs, plus full word 65538 in both lists.
— R2-ASSET-247, R2-ASSET-248

The script companion specifies all 184 trigger offsets: a 128-byte prefix
(label interpretation Medium), three check-ID pairs, four action IDs,
three comparison codes and the flag DWORD at+0xb4. Both loaders bulk-copy
the trigger; unlike nodes, it has no field permutation. Compiler use of the ID-map node,
prefix consumers and the flag of a zero-condition trigger remain open (no reader of the
prefix was found by a census of direct accessor calls, Medium); the miss path
of the ID map is read (find or insert, node value Unknown). — R2-ENGINE-1316, R2-ENGINE-1453, R2-ENGINE-1681

Every occurring ordinary ID is mapped through both dispatch tables.
Twenty-seven action and nineteen check bodies per root are read to RET;
the bodies of actions 33 and 34 are read (R2-ENGINE-1450); their delegated effects are Medium. Local map-slot/Session-DWORD/relation
stores, delegated action calls and check operands are specified at their
individual confidence boundaries. Unread helpers do not establish complete
player effects. — R2-ENGINE-1312, R2-ENGINE-1313, R2-ENGINE-1314,
R2-ENGINE-1315

The final effects of actions 12, 13, 16 to 24, 28 to 30, 32 to 34 and 38 are read
at their own grades and listed in the action table of [script.md](script.md);
the shipped action 21 and 24 nodes and their triggers are tabulated there too.
— R2-ENGINE-1740, R2-ENGINE-1741, R2-ENGINE-1742, R2-ENGINE-1743,
R2-ENGINE-1744, R2-ENGINE-1745, R2-ENGINE-1746, R2-ENGINE-1747,
R2-ENGINE-1748, R2-ENGINE-1749, R2-ASSET-379

The packed full-word action's downstream consumer remains Unknown.
No stored parameter among all 142 maps is 8000..9999, although the action
compiler and executor have a positive deferred-reference path. Its pointer
array producer and return meanings are unproved. Map-result safe limits
remain Unknown and differ from the separate scenario bank's initialized
span. The check compiler marks the same path and no check handler read addresses it
(Medium; R2-ENGINE-1684), and no vector-method consumer of the word-65538 collection was
found in the population named (Medium; R2-ENGINE-1683). — R2-ENGINE-1317, R2-ASSET-249, R2-ENGINE-1318, R2-ENGINE-1319

<a id="byte-planes-and-the-client-fog"></a>

## Byte planes and the client fog

The client copies the type 2 plane, W×H signed bytes, as its height plane:
one byte per tile cell, index `y*W+x`, the same index as the u16 fog words.
The header reader (EN `L2.02735`, RU `R2.0003`) copies 20 bytes and puts `typeId`
at `[ebp-0x34]` unless its by-value header argument carries `0x3e8` at its fifth
dword; the skipped case is not read (Medium). Switch arm 2 (EN `L2.07463`, RU
`L2.07464`) fills `asset+0x14`, and client map builder EN `L2.01385` / RU
`L2.01386` copies `asset+0x14` to `client+0x10` (High in both roots). — R2-ENGINE-645, R2-ENGINE-1666

Every 32nd client tick, when `[L2.00989]` is zero, the client clears fog
bit 0x4000 in all W×H words and calls each client object's stamp slot (EN
census; the RU bodies are equal and the RU callers of the stamp are
unread). A unit whose cell word differs from its last stamped one is also
stamped from its per-tick update, so a stamp can come before the 32nd tick
(Medium; whether it does for the first frame is Unknown). The stamp slot
runs the stamp for the static class and, through a wrapper, the unit
class. The stamp walks rings 1..19 around the object's cell with a budget of
`(sight >> 1) + 64` at the centre. Each
cell's budget is its predecessor's (one step back along x, along y or
diagonally) minus the step cost `trunc(128*sqrt(i²+j²)/max(i,j))` (128 on
an axis, 181 on the diagonal) and minus the cell's type 2 height less the
centre's. Cells with x or y below 7 or above W−8 or H−8 are skipped. A ring
with no positive cell ends the walk. — R2-ENGINE-643, R2-ENGINE-644,
R2-ENGINE-645

Light and Darkness lower the sight word that seeds this walk through a Medium chain
(R2-ENGINE-1240, R2-ENGINE-644); the unit sender carries the changed word to the unit's own
player; for Human and Humanoid units with word +0xe in 0x21..0x3f another player may receive it
(Unknown) (High for the arms, Medium for the reach, the delivery and the fog link).

For each cell whose budget stays positive the stamp ORs 0xc000 into the 2×2
fog words at that cell and its +x, +y and +x+y neighbours. The figure is
offset by `trunc(size/2)` cells from the walk centre, so it lies on the walk
centre for one-cell bodies. An object is drawn only when the bitwise OR of
the top two bits of the four words of its 2×2 block is 0xc000 (the unit
routine also tests byte +0x184, +0x1b8 and the owner gate). —
R2-ENGINE-646

The stamp runs only for owners whose word in the client player table has
the value bit 0x8. No type 5 diplomacy entry of the first mission map has that bit,
so which owners pass is Unknown. — R2-ENGINE-648, R2-ASSET-162

For the first mission map (`10.alm`, 80×80, identical in EN and RU) and a
hero at 11/41 with sight 1648 (the unmodified Start_MF template), a stamp
uncovers 155 cells over x 7..18 and y 35..48; sight 1617 gives 154
(Medium). The refresh stamps no earlier than the 32nd tick; what the first
frame holds is Unknown. — R2-ENGINE-647,
R2-ASSET-162

<a id="cell-planes-bridges-and-plane-load"></a>

## Cell planes, bridges and the plane load

Types 1, 2 and 3 are 2×W×H, W×H and W×H bytes in all 59 EN and 83 RU installed
maps. Above the 10-bit tile index a type 1 word holds only bit 0x2000; type 2
holds 0..127 (no negative value); type 3 holds 0 or one of 191 values from 1 to
209. — R2-ASSET-261

The loader allocates and reads type 1 as W×H×2 bytes whatever its record
says. It reads type 2 by its own size field and type 3 by the size field of
the last type 2 record, so a writer gives type 3 the byte count of the type 2
record and places it after it; the type 3 size field is not read. The results
of the stream reads are not tested. EN and RU bodies are masked-equal
(R2-ASSET-129). — R2-ENGINE-1338

The loader stores codes at map +0x8c: 1 before the stream opens (the final value 1
does not tell a failed open from a load that ran), then bad magic 2, fewer than three
records 3, version above 0x640 4; after the record loop a missing type 1 gives
5 and a missing type 2 gives 6, and a missing type 3 becomes a zero buffer of
the type 2 size. No value of a plane is checked in the bodies read (Medium).
What the caller does with a nonzero status is Unknown. — R2-ENGINE-1339,
R2-ASSET-020

The terrain build then reads each cell at y×W+x and writes the plane at
+0x10000 with stride 256: 1 for type 1 bit 0x2000, for tile class 8 and for
`(tile & 0x300) == 0x200`; 5 for a nonzero type 3 byte, stored last; the type 2
byte goes to +0x944f4. Only bytes are set, never cleared. The border routine
then writes 0x1f over the 8-cell ring of the plane, so a 5 stored there from
type 3 is overwritten; the type 3 buffer itself is unchanged. Server tile
classifier EN `L2.06958` / RU `L2.07473` (High in both roots; writes no height
plane) decodes `tile = word & 0x3ff`: for `(tile & 0x300) == 0x200` it returns
class `1` if `(tile & 0xf) == 4` and `((tile & 0x30) >> 4) == 1`, else `9` for
`low < 8` (terrain byte `8`) and `0xff` for `low >= 8`; otherwise for `low < 14`
it reads weight `w` (`1..5`) from `this+0x540a6 + mid*16 + low`, class pair
`(c0, c1)` from `this+0x54126 + grp*2`, and blends `T[c0]`/`T[c1]` from
`this+0x54146` (`T[0] = 255`; `T[1..10]` are the destinations of the 10 `[Terrain]`
reads of R2-ENGINE-1472, code defaults `8, 8, 9, 14, 6, 12, 11, 16, 8, 6`; their
values come from `map.reg` only if the config object is built from the file, Medium
there) by `w` (`w=1: T[c1]`;
`w=2: (T[c0]+3*T[c1])>>2`; `w=3: (T[c0]+T[c1])>>1`; `w=4: (3*T[c0]+T[c1])>>2`;
`w=5: T[c0]`), returning `c0` if `w > 2` else `c1` (per-map distinct `u16` sets:
530 EN and 589 RU values in groups `0..12`; EN equals RU in all 173 instructions of
the routine). — R2-ENGINE-1340, R2-ENGINE-1341, R2-ENGINE-1472, R2-ENGINE-1665

| `mid` | Offset | Weights `low = 0..13` | Claim |
|---:|---|---|---|
| 0 | `+0x540a6` | 2, 3, 2, 4, 3, 4, 2, 2, 2, 2, 4, 4, 4, 4 | R2-ENGINE-1665 |
| 1 | `+0x540b6` | 3, 5, 3, 3, 1, 3, 2, 4, 2, 2, 4, 2, 4, 4 | R2-ENGINE-1665 |
| 2 | `+0x540c6` | 2, 3, 2, 4, 3, 4, 2, 4, 2, 2, 4, 2, 4, 4 | R2-ENGINE-1665 |
| 3 | `+0x540d6` | 5, 5, 5, 5, 5, 5, 2, 2, 2, 2, 4, 4, 4, 4 | R2-ENGINE-1665 |

| `grp` | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 12 | Claim |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|
| `(c0, c1)` at `+0x54126 + grp*2` | (2,1) | (5,1) | (4,1) | (7,1) | (6,1) | (5,6) | (3,7) | (8,6) | (10,1) | R2-ENGINE-1665 |

Type-3 plane readers: client map builder EN `L2.01385` / RU `L2.01386` zero-fills
`client+0x14` (`W×H` bytes) and copies only the interior rows `y = 8..H-5`,
columns `x = 8..W-9` from `asset+0x10` (High; downstream client readers of
`client+0x14` remain Unknown). On the server, byte `5` (`0x01 | 0x04`) in
`+0x10000`/`+0x20000` blocks footprint masks `0x41` (kind 1) and `0x44` (kind 2)
and admits `0x82` (kind 3) at `L2.02720`/`L2.03523` (Medium composition).
Blizzard cell validity EN `L2.02717` / RU `L2.01700` admits `(x, y)` iff
`8 <= x <= W-9` and `8 <= y <= H-9`, reading only type-0 `W` and `H`
(`+0x50000`, `+0x50004`) and clamping out-of-band coordinates (High in both
roots); that no other check gates Blizzard placement is Medium, because the routine
is shared with Summon positioning (R2-ENGINE-789) and its other callers are unread.
Other readers of the type-3 plane are not excluded. — R2-ENGINE-1667, R2-ENGINE-1668

The route search (RU `L2.05174`, EN `L2.05175`) floods a word grid at +0x30000 + 2 × key from the unit's cell. With
its flag 1 (walk and approach, Medium) it tests plane +0x10000. With flag 0 (repath) it tests plane +0x20000, which
the terrain build fills through a copy call whose body is unread (R2-ENGINE-1340) and which R2-ENGINE-809 rebuilds
with occupancy bits and structure masks. A neighbour is skipped when its plane byte AND the unit's mask byte (0x41, 0x44
or 0x82 for unit kinds 1, 2 and 3; Medium for the live mask) is nonzero. A unit of size 1 tests the 8 neighbours; a
larger unit tests a 3 × 3 neighbourhood through the footprint test. A step costs, for kind 1, the terrain byte
at +key straight and one and a half times it diagonal (rounded down), and for other kinds 2 straight and 3
diagonal; a neighbour's word is replaced when the new cost is below it. No worker or helper read takes the type-2
byte (+0x944f4); `L2.03871`, `L2.05228`, `L2.05231`, `L2.05229`, `L2.02566`, the stand-in and the callers of the
search are unread (Medium). — R2-ENGINE-1804, R2-ENGINE-871, R2-ENGINE-809

What each value of plane +0x10000 as the terrain build stores it closes, for flag 1 (Medium composition; the flag-0
plane is not derived here):

| Cell content | +0x10000 | Kind 1 (0x41) | Kind 2 (0x44) | Kind 3 (0x82) | Claim |
|---|---:|---|---|---|---|
| type-1 bit 0x2000; tile class 8; `(tile & 0x300) == 0x200` | 1 | closed | open | open | R2-ENGINE-1805 |
| nonzero type-3 byte | 5 | closed | closed | open | R2-ENGINE-1805 |
| 8-cell border ring | 0x1f | closed | closed | closed | R2-ENGINE-1805 |
| a cell none of the three stores touches | 0 (initial value and other writers of +0x10000 not read here) | open | open | open | R2-ENGINE-1805 |

The table is the test of one cell. A unit of size 1 applies it to each of 8 neighbours; a larger unit applies it to
every cell of its size × size footprint from the neighbour's key (RU `L2.08532`), so a neighbour is closed when any
footprint cell is closed. A unit whose virtual +0x20 is not 1, 2 or 3 keeps mask 0x41 (R2-ENGINE-809) and takes the
step cost of the other kinds. — R2-ENGINE-1805, R2-ENGINE-1804

Tile class 8 is group 7 (`(tile & 0x3c0) >> 6 = 7`) with weight above 2: `low` in {1, 3, 4, 5, 10, 11, 12, 13} for
`mid` 0, {0, 1, 2, 3, 5, 7, 10, 12, 13} for `mid` 1, {1, 3, 4, 5, 7, 10, 12, 13} for `mid` 2 and {0..5, 10..13} for
`mid` 3. In the route-search workers and helpers read (RU eight workers and the footprint test, EN three workers) a terrain
byte of 255 (class 13 and above) enters only kind 1's step sum and no test closes a cell on it; `L2.03871`,
`L2.05228`, `L2.05231`, `L2.05229` and the stand-in are unread (Medium). What a unit of
each kind does on these cells in play is Unknown. — R2-ENGINE-1805, R2-ENGINE-1665

The loader and the build compare no plane size with W×H and no value with a
range (Medium). All 142 installed maps are square and none is wider than
256. What a short plane or a map wider than 256 columns does in play is
Unknown. — R2-ENGINE-1342

A type 4 element of kind 0x21 or 0x1000021 carries two extra dwords, an extent
rectangle (EN 41, RU 170; kind 0x25 never occurs). That such an element is a
bridge is the working label of this experiment (Medium; no string or caller was
read). Under the assumed footprint of the extra dwords no type 1 bit marks its
cells, and the game's reader of the element is unread (Medium). The editor's
creator (EN, 256 instructions read) leaves through a jump outside the part read
for an x or y below 8 or an extent reaching W−8 or H−8, builds the object, and
writes no cell plane in the part read; whether the id it searches is the new
object's id is not read, and what the rest of the editor writes is Unknown.
— R2-ASSET-262, R2-ASSET-263

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
| 0x0a | u16 | +0x0c | no reader found in the census of 21 references and of the 7 direct calls of the row routines per root (Medium); 0..7, 0 only in the person arm (R2-ASSET-313, R2-ENGINE-1510, R2-ENGINE-1801) |
| 0x0c | u32 | +0x48 | flags; bit 0x10 selects the person arm |
| 0x10 | u32 | +0x50 | bits 1, 2, 4, 8 set actor +0x194, +0x19c, +0x198, +0x1a0; read at version 0x47e and above. Bits 1, 2 and 4 put the unit in Inn offer populations; bit 8 removes its death drop and the experience a human attacker is paid for it (Medium) |
| 0x14 | u32 | +0x10 | definition id; read at version 0x3db and above; names the Units row (flag 0x10 clear) or Humans row by serverID, see [the definition tables](map-definitions.md#type-6-definition-id-to-databin-row) (R2-ENGINE-1512, R2-ASSET-316) |
| 0x18 | u32 | +0x14 | owner, low word only |
| 0x1c | u32 | +0x44 | 1-based type-8 index; that type-8 element takes the record's 0x28 low word |
| 0x20 | u32 | +0x1c | no reader found in the census of 21 references and of the 7 direct calls of the row routines per root (Medium); 0..14, 10 in 6,255 of 6,395 EN Units-arm records (R2-ASSET-313, R2-ENGINE-1510, R2-ENGINE-1801) |
| 0x24, 0x26 | i16 | +0x20, +0x22 | actor words +0x94 and +0x96 in both arms, after the Units arm's scaling by object +0xd8 (difficulty, Medium, R2-ENGINE-1002); -1 keeps the default; an equal pair other than -1 is reset to -1 in the person arm only (R2-ENGINE-1511) |
| 0x28 | u32 | +0x3c | low word read; two bytes below version 0x456, four from 0x456 |
| 0x2c | u32 | +0x40 | group number; the map keeps the maximum |

"No reader found" is bounded to the actor builder, the mission-load
row loop and the direct calls of the row routines (R2-ENGINE-1801). Installed values in the campaign maps: 0x08 1..107, 0x0a 1..6,
0x0c {0, 16}, 0x10 {0, 8}, 0x1c 0..4, 0x20 0..14 (10 in nearly every
record), 0x2c 0..38. In the root `.alm` maps: 0x0a 0..7, 0x0c {0, 4, 16,
20}, 0x10 {0, 1, 2, 3, 4, 5, 8, 9, 11, 12} in EN and also 10 in RU (the builder tests each bit separately; R2-ENGINE-513), 0x18 2..15 with high
word 0, 0x1c 0..58, 0x20 0..14, 0x2c 0..555.
— R2-ASSET-119, R2-ASSET-122, R2-ASSET-132, R2-ENGINE-038, R2-ENGINE-513

The actor builder and the mission-load loop are the only bodies read that take fields from the row (the builder: x,
y, 0x14, 0x18, 0x24, 0x26, 0x28, 0x2c, 0x0c, 0x10; the loop: the same without 0x10, with 0x08); none of 21 direct
references to the list per root reads the slot of 0x0a or 0x20; a route that hands the list or a row to another body
is unread (Medium). — R2-ENGINE-1510

The two list routines the row loop calls on the list (EN `L2.08531`, `L2.07902`; RU `L2.02552`, `L2.02553`) have 7
direct call sites per root and no stored dword, all in the row loop, the map teardown and the actor builder. So no
body reaches a row through a direct call of them and reads the slot of 0x0a or 0x20; a computed call or a copy of a
row is outside that census (Medium). — R2-ENGINE-1801

Value sets in the installed maps (EN 59 maps, 7,461 records; RU 83, 15,482), flag 0x10 clear is the Units arm, set the
person (Humans) arm. 0x0a: EN Units arm 1..6, person arm 0..7; RU 1..7 and 0..7. 0x20: EN Units arm 0, 2, 4, 6, 8, 10,
12, 14 and the person arm only 10; RU has odd values and 532 person-arm zeros. 0x18 owner: 1..15, high word 0, never
above the map's player count. 0x0c is 0, 4, 16 or 20, with bit 0x10 set in EN 1,066 and RU 1,847 records. No field of
the record is read as a level or an experience value in the two bodies read (the actor builder and the mission-load loop); where the game takes a placed unit's level
is Unknown. — R2-ASSET-313, R2-ASSET-314

0x24 and 0x26: EN 6,301 records hold -1 in both, 1,155 an equal pair, 5 only 0x24; RU 12,965, 2,512 and 5. The equal
value is the row's healthMax in 917 of 935 EN Units-arm pairs (RU 1,934 of 2,168). In the person arm an equal pair is
reset to -1 by the builder and so has no effect (EN 220, RU 344 records). That the actor words are current and
maximum health is Medium. — R2-ASSET-315, R2-ENGINE-1511

In every installed map (EN 59, RU 83) the 0x28 high word equals the low word of x and the low word is never 0;
the game reads the low word only and keys its unit table by it. The id never repeats within a map and never
equals a type-4 id; types 4 and 6 are in ascending id order. 90 EN and 343 RU units name a sack through 0x1c,
and each named sack holds the unit's id at +0x04 and its x and y. A unit is skipped with a log line, and the map
is not rejected, when its owner is not found, when the built actor's word +0x0e is 0 (set by the object build; not a record field) or when it cannot be placed on its tile.
— R2-ASSET-253, R2-ENGINE-1325, R2-ENGINE-1326

**Type 8 (sack).** Head: item count (+0x00), owner id (+0x04), x (+0x08), y (+0x0c), +0x10; then per item a u32
id, a u16 word and a u32 binding. +0x04 is a unit id, which the game resolves through the unit table, or 0 for
a ground sack (25 per root, x low byte 128), which the game places at x>>8, y>>8 with +0x10. A unit-owned sack
uses none of x, y and +0x10 in the game. The item id's high word is 0; its low word splits into A (bits
8-11), B (12-15), C (5-7) and D (0-4), and an id the builder cannot build is skipped. A = 1 builds a weapon, A = 2 a
shield, A = 3..13 an armour, A = 0xe a magic item whose row is the low id byte (table names by use, Medium); A = 0 and A = 0xf build nothing. The
installed maps hold EN 139 and RU 327 items, all with a row in range, and for every weapon, shield and armour the
row's Slot parameter equals A; none has A of 0, 3, 11, 13 or 15 (rows in [the definition
tables](map-definitions.md#sack-item-to-databin-row)). — R2-ENGINE-1513, R2-ENGINE-1514, R2-ASSET-319. The item word is 0 in
the 139 EN and 327 RU items of the 59 EN and 83 RU `.alm` maps of the root folder and the `scenario.res` members
(`world.res`, saves and user maps are outside that census). The loader appends it to the sack's list +0x28; stage C
(EN only when Session +0x74 is 0, RU always; R2-ENGINE-1328) gives, for a unit-owned sack, an item whose word is
nonzero to the owner's slot +0x44 (equip, a displaced or refused item going to the bag; Medium), and an item
whose word is 0 to the owner's bag (Medium); the effect in play is Unknown (R2-ENGINE-1803). The binding is 0 or a 1-based type-9 position; that element
has cell (+8, +0xc) zero, +0x00 equal to the item's 0-based position and +4 nonzero, or the item is skipped.
A sack owner naming no unit is not tested before use; the result is Unknown.
EN `L2.08239` and RU `L2.08240` (113/113 masked-equal) share one version-gated
grammar (`>= 0x3dd`: the arm reads the `+0x10` dword when the version is not below
`0x3dd`; no installed map has version `0x3dd`), consuming every installed map's
type-8 payload with 0 bytes left over across all 59 EN and 83 RU maps.
— R2-ASSET-254, R2-ASSET-359, R2-ENGINE-1326, R2-ENGINE-1327, R2-ENGINE-1328

| Field | Width | Where the loader keeps it | Game use | Claim |
|---|---|---|---|---|
| head +0x00, item count | u32 | loop bound | one item read per count | R2-ENGINE-1803 |
| head +0x04, owner id | u32 | sack +0x3c | nonzero: the owner from the unit table; 0: a ground sack in a new bag | R2-ENGINE-1803 |
| head +0x08, +0x0c, x and y | u32 | sack +0x40, +0x44 | ground sack only: placed at x >> 8, y >> 8 | R2-ENGINE-1803, R2-ENGINE-1328 |
| head +0x10 | u32 | sack +0x48 from version 0x3dd, else 0 | ground sack only: passed on | R2-ENGINE-1803, R2-ENGINE-1328 |
| item id | u32 | appended to the sack | low word decoded to A, B, C, D; a null build skips the item | R2-ENGINE-1327, R2-ENGINE-1513 |
| item word | u16 | sack list +0x28 | unit-owned sack, stage C (EN only when Session +0x74 is 0, RU always): nonzero to slot +0x44 (equip, Medium), 0 to the bag (Medium) | R2-ENGINE-1803 |
| item binding | u32 | sack list +0x14 | type-9 test; a failed test skips the item | R2-ENGINE-1327, R2-ENGINE-1803 |

Actor +0x1a0 also receives 1 from two unique-spell arms and, when
Session+0x74 is set, from the map build for the units of owners whose
diplomacy toward the owner named "Self" has bits 0..2 clear. An actor with
+0x1a0 set drops no bag or gold at death and pays no experience to a human
attacker (R2-ENGINE-480; Medium here; a creature attacker's award was not
read for the flag). By the code read, the flag is the second refusal of the
Human award, after the attacker class test (R2-ENGINE-1156); by the code read,
for a Human killer, a kill pays through the killer's repeated-victim list
(R2-ENGINE-1157, R2-ENGINE-1158). The unit copy omits the flag (Medium; the complete
selected bodies are now read, but callee effects remain unclosed, R2-ENGINE-1174);
the save keeps it (Medium). Actor +0x194, +0x19c and +0x198 select
Inn offer blocks 2, 4 and 6 (Medium: the receiver of the offer generator is
the Inn by address inference); RU and EN agree in those blocks. A template
name holding "NPC" (Medium: the substring helper is unread) also empties
the bag at death; with +0x1a0 clear such a
unit still rolls gold and an item. — R2-ENGINE-535, R2-ENGINE-537, R2-ENGINE-943

The repeated-victim array at Unit +0x1a8 has a 20-byte header in RU:
buffer +4, size +8, capacity +0xc and growth +0x10. Its getter returns
buffer+4*i. Capacity is dynamic; five entries is the award's policy, not
inline storage. These local RU reads are High; the effects of the unread
allocation, copy and fill helpers L2.03271, L2.06284, L2.04512, L2.00342
and EN L2.03812, and the unread EN array bodies, remain Medium.
— R2-ENGINE-1172

Five selected constructors call the array initializer in both roots; RU
header fields start at zero. Complete creation paths remain Medium.
The site RU L2.06420 / EN L2.06447 is array cleanup in Unit destruction,
not unit copying. — R2-ENGINE-1173

The selected Unit assignment bodies have no explicit array copy or clear.
By the code read, with no run of the original, and subject to unread
callees, assignment keeps the destination's history;
an empty fresh destination and a reused destination differ (Medium).
Field and helper-reference sweeps leave unclassified bases, virtual calls
and bulk copies, so no complete lifecycle census is claimed.
— R2-ENGINE-1174, R2-ENGINE-1175

Each root's 46 campaign maps hold 34 type-6 records of the Humans row
`2U_PeasantLeader` on maps 20, 21, 30, 31, 40, 42, 43, 52, 61, 71, 73, 75, 85,
92, 93, 94, 95, 96, 102 and 103. All carry definition id 2112 and a zero 0x10
field; the row's parameter 16 (typeID) is 1. The census is High for the 46
`scenario.res` maps of each root; the row join is Medium. — R2-ASSET-219

Composed (Medium): these placements show `npcnames.txt` line 2111, `Elder` /
`Староста`; the type label `Peasant` / `Крестьянин` shows only with an empty
unit name and NPC word 0. — R2-ENGINE-912

<a id="group-and-object-records-types-10-and-11"></a>

## Group and object records (types 10 and 11)

Type 10 entries are four i32, keyed on a unit group number:

| Offset | Field |
|---:|---|
| 0x00 | group number (unit record 0x2c) |
| 0x04 | percent e1: the group object's countdown reload is e1 × S / 100 container passes; S is 100 unless a settings parser stores 20..500 (EN read; Medium) |
| 0x08 | flags 1, 2, 4, 8 |
| 0x0c | position in the type-7 first array; the count of nodes before it with dword +0x40 below 0x10002 is passed to the routine `L2.02490` (Medium) |

The program runs the type-10 pass when type-0 payload dword 0x84 exceeds 1
(gate body High; "only" Medium, other writers of the gate are open). Then
each unit group of an owner with a nonzero +0x2c whose dword +0xc is not 0
(read by `L2.02502`; that it is the member count is Medium) gets a group
object, with defaults when no entry matches; every matching entry applies
in file order. The 46 `scenario.res` campaign maps per root have dword
0x84 = 1 and no type-10 entry. The 13 EN and 37 RU root `.alm` maps have
dword 0x84 of 4..16, above 1, and type-10 entries. The map build has one
direct call site per root, in the map caller; that a root map reaches the pass
through it, and that a load does not, is Medium (the load routines and the callers
of the pass are unread). Type-0 payload dword 0x88 is
stored in the map and the session; the editor's map-properties routine computes it from the
largest unit level (the EN body read; Medium for when it runs; the save routine also sets map `+0xdc` to 1 and
rewrites it in its placed-item loop, and whether that loop gives the same value is not read), and its meaning in the
game is Unknown.
— R2-ASSET-120, R2-ASSET-122, R2-ASSET-130, R2-ASSET-255, R2-ENGINE-509, R2-ENGINE-510,
R2-ENGINE-538, R2-ENGINE-1816, R2-ENGINE-1819

Flag 4 and flag 8 also put the group in Inn offer blocks 3 and 5 through
group +0x48 and +0x4c. In block 3 RU skips groups of the Inn's own owner
and EN does not. `0xcc` slot 5 (EN `L2.08245`, RU `L2.02521`; first 17 of 152 instructions compared
here, masked-equal) and respawn method EN `L2.08237` / RU `L2.08238` (156
instructions, equal; the compared range of 234 includes later routines) re-create
members only when `L2.02502(group) <= 0`, pass `+0xbc` (flag 2; RU stores spread,
EN keeps 0) to `L2.00312`/`L2.00313`, copy `+0xc4`/`+0xc8` (flags 4 and 8) to
the recreated `group+0x48`/`+0x4c`, and call script action executor
`L2.07383`/`L2.02490` with `+0xb8` (flag 1 + position) when `>= 0` (High for the
read bodies; Medium for caller/reader completeness across 37 direct `.text`
`0x310` sites per root: 20 in map constructors, loader, teardown, the type-10 pass
and unwind thunks, 17 on other bases that were not read). Slot 5 runs once per container pass, which runs once per 16
session ticks in RU (Medium; EN tick unread), so the countdown counts container passes, and the check repeats on every
pass until the group is empty. Per member the respawn calls a template restore and the positioning routine, with the
object's position and radius `+0xbc` when `+0xbc` is nonzero, else with the member's own position and radius 0, which
is every EN member (where the member then stands is Unknown); a successful placement reloads `+0xb4` from `+0xb0`, and
no field read here bounds the cycles. — R2-ENGINE-510, R2-ENGINE-537, R2-ENGINE-538, R2-ENGINE-1670, R2-ENGINE-1817,
R2-ENGINE-1818

Script action 31 filters the Buildings collection by vtable slot +0x38 (RU: returns 0 for Pointer, Shop, Inn and, by
R2-ENGINE-517, Building, and 1 for Outpost; EN: returns 0 for Building and Pointer and 1 for Outpost; Building is the
base of four class records per root and Outpost of none). It writes p1 into +0xb4 of every Outpost whose group number
(+0xac) equals p0, for every owner and with no loop exit (the filter passes only the Outpost in the classes with a
record; EN Shop and Inn slot 14 unread, Medium). Outpost::slot5 decrements +0xb4 once +0xc0 is 1 and triggers respawn
when it reaches 0 or below. p1 counts container passes without S; while the countdown runs, the check fires on the
(p1 + 1)-th pass (Medium), and a re-formation reloads +0xb0. — R2-ENGINE-510, R2-ENGINE-537, R2-ENGINE-538,
R2-ENGINE-1450 (amended), R2-ENGINE-1701 (amended), R2-ENGINE-1820, R2-ENGINE-1821

Over all 142 installed maps (0 campaign entries; 674 EN and 1,696 RU root
entries), type-10 flags use only bits 1, 2, 4 and 8 (bit 1: 7 EN / 26 RU;
bit 2: 668 / 1,665; bit 4: 358 / 937; bit 8: 67 / 109) and position (`+0x0c`)
takes `-1` (666 EN / 1,668 RU) or `0..9` (`0`: 1/1, `1`: 3/7, `2`: 3/7,
`3`: 1/5, `4`: 0/2, `5`: 0/2, `6..9`: 0/1 each); every bit-1 entry has
`pos >= 0`, while 1 EN and 2 RU entries have `pos >= 0` with bit 1 clear.
— R2-ASSET-360

| Type-10 flag mask | Bits | EN | RU | Claim |
|---:|---|---:|---:|---|
| 0 | none | 6 | 26 | R2-ASSET-360 |
| 2 | 2 | 268 | 649 | R2-ASSET-360 |
| 3 | 1, 2 | 7 | 22 | R2-ASSET-360 |
| 4 | 4 | 0 | 5 | R2-ASSET-360 |
| 6 | 2, 4 | 326 | 881 | R2-ASSET-360 |
| 7 | 1, 2, 4 | 0 | 4 | R2-ASSET-360 |
| 10 | 2, 8 | 35 | 62 | R2-ASSET-360 |
| 14 | 2, 4, 8 | 32 | 47 | R2-ASSET-360 |

Percent e1 over the same entries: 120 in 494 EN and 1,009 RU entries, above 120 in 166 and 662, below 120 in 14 and
25; 22 EN and 72 RU distinct values, 5..7500 EN and 5..30000 RU. CROSS2.ALM holds 238 entries per root and
SHALLOW.ALM 108, the most. An entry that matches a group object sets its reload to e1 × S / 100; the counts are entries,
not group objects built. With S = 100 an e1 of 120 fires on the 121st container pass (Medium).
— R2-ASSET-397

| Group object field | Use in the read bodies | Claim |
|---:|---|---|
| +0x14 | owner, from group +0x44 | R2-ENGINE-510 |
| +0x6c | member list; the respawn walks it | R2-ENGINE-1818 |
| +0x8c | template list, one copy per member, walked in step with +0x6c | R2-ENGINE-1818 |
| +0xac | group number; slot 5 looks the group up by it, action 31 matches it against p0 | R2-ENGINE-1820 |
| +0xb0 | countdown reload: 120 × S / 100, or e1 × S / 100 from a matching entry | R2-ENGINE-1816 |
| +0xb4 | countdown in container passes; action 31 writes p1 into it | R2-ENGINE-1817, R2-ENGINE-1821 |
| +0xb8 | argument of `L2.02490` at build and after a re-formation when 0 or above (set under flag 1), else -1 | R2-ENGINE-510 |
| +0xbc | placement radius around the object; 0 places a member at its own position | R2-ENGINE-1818 |
| +0xc0 | 1 once no member has byte +0x13c below 3; the countdown runs only then | R2-ENGINE-1817 |
| +0xc4, +0xc8 | copied to group +0x48 and +0x4c after a re-formation | R2-ENGINE-1818 |

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
all have gate 0. 11a flag 1 occurs in 143 EN and 158 RU root entries (23 EN and 30 RU distinct
flag masks; bits `0x08` and `>= 0x80` never occur); in direct `.text` `0x338`
scans of EN and RU `allods2.exe`, all 17 hits per root lie in map ctors,
dispatcher case 11a, teardown, SEH unwinds and the type-4 pass (which stores
nothing from bit 1), so no reader of flag 1 was found in that population
(Medium; indirect and other-binary routes unread); under flag 2 the word is
3661..3680. — R2-ASSET-121, R2-ASSET-122, R2-ASSET-130, R2-ASSET-132,
R2-ASSET-361, R2-ENGINE-512, R2-ENGINE-1671

| 11a mask | EN | RU | 11a mask | EN | RU | 11a mask | EN | RU | Claim |
|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|
| `0x00` | 0 | 2 | `0x31` | 7 | 7 | `0x61` | 16 | 22 | R2-ASSET-361 |
| `0x02` | 0 | 1 | `0x32` | 0 | 2 | `0x62` | 15 | 17 | R2-ASSET-361 |
| `0x14` | 26 | 26 | `0x33` | 17 | 21 | `0x63` | 11 | 11 | R2-ASSET-361 |
| `0x20` | 21 | 28 | `0x34` | 0 | 2 | `0x64` | 0 | 1 | R2-ASSET-361 |
| `0x21` | 2 | 2 | `0x35` | 6 | 6 | `0x70` | 11 | 36 | R2-ASSET-361 |
| `0x22` | 51 | 62 | `0x36` | 10 | 10 | `0x71` | 8 | 12 | R2-ASSET-361 |
| `0x23` | 26 | 26 | `0x37` | 18 | 18 | `0x72` | 55 | 60 | R2-ASSET-361 |
| `0x24` | 2 | 4 | `0x41` | 8 | 8 | `0x73` | 13 | 14 | R2-ASSET-361 |
| `0x25` | 4 | 4 | `0x43` | 7 | 7 | `0x74` | 0 | 3 | R2-ASSET-361 |
| `0x30` | 11 | 15 | `0x60` | 31 | 49 | `0x76` | 0 | 1 | R2-ASSET-361 |

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
no earlier node below 0x10002 failed. In the same state-15 arm, kinds 76
and 77 (0x4c, 0x4d, with kinds 28 and 29; first lookup) toggle signed building word +0x42 between
0 and 1 and call notify (unread, Medium) with no unit or inventory write;
kinds 15 and 16 (0x0f, 0x10; second lookup directly at unit +0x10) require
signed word +0x42 above 0, decrement it by 1, call notify, construct a named
Item from literal strings EN `L2.08314` / `L2.08316` (RU `L2.08315` /
`L2.08317`; the texts are "Potion Big Healing" and "Potion Big Mana"; the Data.bin MagicItems effects they name are Unknown)
and pass it to unit vtable +0x44 (EN read; RU twin unread, Medium), where Item vtable +0x38 attaches child
effects, recomputes, destroys the Item and returns 0 so it does not enter
unit +0x7c. — R2-ENGINE-517, R2-ENGINE-518 (amended), R2-ENGINE-519,
R2-ENGINE-520, R2-ENGINE-1695

A map click on cursor `L2.02987`, which EN chooses only on hover flag 0x800
(a structure whose record has +0x8c set), is an order, not a walk. The
hover routine walks its screen-rectangle list from the last entry down and
names the object of the first accepted rectangle meeting the cursor box
(x-1,y-1,x+1,y+1), coordinate span 2 per axis. The former 3x3 size label in
R2-ENGINE-605 is retracted; the read wrapper calls IntersectRect. List
store/index arithmetic is read, while allocation and complete indirect-route
coverage remain Medium (R2-ENGINE-849, R2-ENGINE-850); a structure whose kind record has +0x8c 0 and +0x64
set is skipped, and +0x8c set gives hover flag 0x800. The literal rectangle of a
Pointer kind and its +0x8c remain Unknown. The read map constructor uses
CStructure for ordinary structure kinds; linking Pointer through kind numbers
remains Medium (R2-ENGINE-853), so that a signpost click takes this path
is Medium. On cursor `L2.02987` a click sends message type 0x24
with the hovered object's ID and the IDs of the units that `L2.02988`
accepts (that these are the selected units is Medium). The server finds the
building whose dword +4 is that ID and sets order 0xf: command +0x68 = the
building, target +0xa = word +2 of its position object, range +0x14 = 1 for
building byte +0x60 of 1 or 2, 2 with target + 0x101 for 3 or 4, else 3 with
target + 0x202 (that +0x60 is the Buildings sizeX is Medium). Prepared 0xf
calls the state-15 handler with +0x68. Until the range test passes the
handler calls an approach step and sets state 1 (Medium that the unit walks
toward the target); when it passes, the unit enters state 15. — R2-ENGINE-605,
R2-ENGINE-606, R2-ENGINE-607, R2-ENGINE-519

The 46 campaign maps per root hold 13 Pointers and no Shop or Inn. The 35
gate-1 11c entries name 12 Pointers; for each, the action node at index
position exists and has operation 2. That the Pointer selects that node at
run time is Medium; the thirteenth Pointer has only a gate-0 entry. The
root `.alm` maps are censused: they hold 8 Pointers, 50 Shops and 50 Inns
(EN) and 8, 151, 132 (RU). Type-4 kinds 0x46..0x4b occur in the campaign
maps, so the two-value kind vocabulary of R2-ASSET-019 holds only for the
files it described. — R2-ASSET-123, R2-ENGINE-520, R2-ASSET-019, R2-ASSET-237

The Shop and Inn slot-5 bodies read in RU do not read the 11b or 11a
fields; slot 5 runs once per container pass (R2-ENGINE-1817), so the bodies run
once per 16 session ticks (Medium).
— R2-ENGINE-521, R2-ENGINE-1817

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
neither. Theme values are -1..16, where `t` in `0..16` directly indexes entry `t`
(`music\B00.wav`..`music\B16.wav`, archive keys `b00.wav`..`b16.wav`) of the
mission music list in both roots (full table on the music page). — R2-ASSET-084, R2-ENGINE-1672

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

<a id="editor-write-sequence"></a>

## Editor write sequence

`ROM2 Map Editor.exe` saves a full map with one routine (EN and RU bodies
masked-equal). It writes the file header `M7R\0`, 20, `4*W*H+72`, record count
and version 0x640, then the records in the order `0,1,2,3,5,11,4,9,8,6,7,10,12`.
The count is 13 when the map's placed-item container is non-empty (container
identity from loop shape, Medium). When it is empty the writer leaves out
types 4 and 6 and writes 11.

| Record header field | Value written |
|---|---|
| `+0x00` tag | One map field, the same on every record |
| `+0x04` header length | 20 |
| `+0x08` payloadSize | Type 0: the constant 644. Other types: the payload length: from W and H (1, 2, 3), from the type-0 counts (5, 6, 10, 11, 12), summed over the elements (4, 9, 8) or from map fields (7), per R2-ASSET-231 |
| `+0x0c` typeId | The type |
| `+0x10` scalar | Not stored by either writer (read); every installed map carries one value on all 13 of its records, which fits an unwritten stack slot (inference) |

The type-0 record is 660 bytes, in the loader's field order. `+0x18` is the
loaded value with `1<<((cell>>6)&15)` of every masked cell ORed in, then
`0xf00` ORed in when the value exceeds 0xff, then capped at 0x1fff. Both
writers do this. The routine does not clear it first; whether the loader or
`L2.07298` leaves it clear is not read. `+0x1c`, `+0x20`, `+0x24` and
`+0x30..+0x40` are the counts that size types 5, 4, 6, 10, 11 and 12. What the
type-10, 11 and 12 counts count was not read in the helpers that compute them
(Medium). `+0x20` counts the editor's CStructure objects, `+0x24` its CUnit objects
and `+0x2c` its sack list; in every installed map they equal the type-4, type-6 and
type-8 record counts, and `+0x28` equals the sum of the three type-7 counts
(R2-ASSET-255). The game bounds its type-4, type-6 and type-8 loops by them and
overwrites its copy of `+0x28`; it compares none with a payload size (Medium,
R2-ENGINE-1324). The editor removes entries from three
lists before it writes them (R2-ASSET-231). Before writing, the editor sets the
first `8*W` tile cells to the last cell's value and masks every cell with
0x23ff.

A second routine writes terrain only: records 0–3 with record count 4. Its
type-0 record is 644 bytes and declares 644; the full writer's type 0 carries
four more dwords, `+0x34` to `+0x40`. At version 0x640 the loader consumes 660,
so the next header is read 16 bytes late (Medium, not run). A global flag in
the save routine selects between the two routines (EN read; in RU only the two
call sites, `L2.07296` and `L2.07297`, are found).

Re-heading an installed map with these header values (payloads, tag and scalar
kept) gives the same file for 56 of the 142 installed maps (both roots; 87
distinct contents). The other 86 are the version-1300 maps; each differs only
in the version field. The census tests the file header, the record order and
the header lengths, not any payload, and sees only the loose `.alm` files of
the two root folders and the `scenario.res` members. Every installed map
carries one scalar and one tag on all its records, and the tile and count
fingerprints of the save. What a newly created map holds before any placement
remains Unknown.
— R2-ASSET-231, R2-ASSET-232, R2-ASSET-233, R2-ASSET-234

**Placed elements.** The editor keeps structures, units and ground sacks in one id map. A new element takes the
lowest free id from 1, map `+0x8c` becomes id + 1 and the position is cell × 256 + 0x80
(R2-ASSET-256). These are EN reads; RU is read for the document creators only. A new unit copies the template unit with the current player as owner and, when the template
holds a sack, a copy of that sack with `+0x04` = the unit id and its bound type-9 elements. A new structure takes
the template's kind, x, y and word `+0x0c`, the current player when the template has none, and, for the Shop, Inn
and Pointer classes, a default type-11 entry with the structure's id (Medium: class-to-list link by list identity). A ground sack is created only with at
least one item; its `+0x04` is written as 0. Dialog fields, default record contents and the object tool are
unread (R2-ASSET-257). Deleting a unit (EN bodies read; RU removes only the sack at the unit's sack index; ground-sack delete and move are
read in EN only) removes its sacks and bound type-9 elements and renumbers the later
bindings and sack indices; deleting a structure removes its mode-4 type-9 elements and its type-11 entries.
No other id changes, map `+0x8c` is not lowered and type-7 and type-10 records are not touched, so a script
reference to a removed id stays (Medium: the bodies read). Moving changes x and y only; the unit's 0x28 high word
and its sack's coordinates follow at save (R2-ASSET-258). The editor's own reader refuses a structure within 8
cells of the map edge (R2-ASSET-259).

## Type-9 records with mode mask 4

An element is, in file order: u32 `+0x00` (only at version 0x3de and above), u32 `+8`, u32 `+0xc`, u16 mode at
`+0x10`, u32 key at `+0x12`, u32 at `+0x16` (bytes `+0x16` and `+0x18` used), u32 count, count × (u32, u16) entries.
Read from RU `L2.05946` and EN `L2.07314`, equal under the address mask. Mode values in the census are 0, 1, 2, 4 and
5. "Mode mask 4" is mode & 4. — R2-ASSET-230, R2-ENGINE-1292

The population of 142 maps (EN 59, RU 83) has 99 EN and 322 RU mask-4 records, on 9 EN and 29 RU map files. Maps 72
and 81 of `scenario.res` hold 27 and 3 in both roots. The 9 mode-5 records are at cell 0,0 and have a key above 32767,
and the 412 mode-4 records have a nonzero cell. — R2-ASSET-230, R2-ENGINE-1009

The mission-load stage builds one object for an element with mode & 4, a nonzero cell and `+4` zero, when the key
names a building by its `+8`; otherwise an error is reported and nothing is built. The object copies the building's
location and owner pointer, the element's `+0x00`, the spell byte `+0x16`, the power byte `+0x18` and the cell low
bytes. `+4` is set by the type-8 arm. Neither the mode-4 stage arm (EN `L2.07317..L2.08244`, RU `L2.05960..L2.05961`)
nor the caster constructor (`L2.05956`/`L2.05954`) reads the trailing count or `(u32, u16)` entries at `+0x1a`
(High for the read bodies, Medium that no other routine reads `+0x1a`); all 412 installed mode-4 elements (98 EN,
314 RU) have count 0; of the 9 mode-5 records the 1 EN record has count 0 and 7 of
the 8 RU records carry count 1..3 (the sack binding of 5 RU mode-5 elements is R2-ASSET-254). — R2-ENGINE-1006, R2-ENGINE-1669, R2-ASSET-230

The turn routine, once per 16 session ticks (Medium), runs a consumer that walks the objects in the session list. For
each object the spell table at Session `+0xf8` gives T; a unit chosen by the relation bit against T `+0xa` is passed to
the call of `L2.05951`. The turn routine runs when Session `+4` mod 16 is 6, so a caster with a candidate casts once per 16
session ticks, with no mana, timer or cooldown test in the consumer, pick, fill and cast entry (the built unit has no mage bit, so its own cast pays no mana, and it casts once, Medium). Read in RU; the EN call site of the consumer and the
EN type-9 arm are Medium. — R2-ENGINE-1007, R2-ENGINE-1008, R2-ENGINE-1562, R2-ENGINE-1740

The candidates are the units of the map list whose cell is within `+0x00` cells (Chebyshev (Medium) distance, a square; the absolute value, the position-word encoding and the list builder are unread) of the record's
cell, then the relation filter; no Spell Max Range is compared in the bodies read. Map 72 holds radii 5 to 8 and spells
10, 2, 16, 6, 5 and 18, power 255 except two Stone Curse records. — R2-ENGINE-1560, R2-ASSET-334

The cast body builds a new Unit object in state 13 with the record's spell, the power byte as its skill in the spell's
sphere, wind-up 1 and recovery 1, aimed at the chosen unit, and the cast routine then applies as for any unit. The world phase deletes
that unit when its recovery ends, so it casts once (Medium). — R2-ENGINE-1561 (amended), R2-ENGINE-1740

Field ranges: spell 1..29; power 2..255; `+0x00` 1..9 on mode 4 and 0 or 2 on mode 5; cell low bytes up to 243 and
242. The consumer tests no range of the spell or power. `+0x00` is the radius in cells of the candidate fill; the power byte
is the skill of the built unit in the spell's sphere; with the unit's Mind of 30 it equals the spell power (Medium). — R2-ENGINE-1009,
R2-ENGINE-1560, R2-ENGINE-1561, R2-ENGINE-1740, R2-ENGINE-1741

Type-4 `+0x0e` is resolved at load into the building's `+0x14`. The built object copies that pointer, and the consumer
uses its word `+4` as the row of the relation matrix. A call of `L2.01703` whose computed value is not below its target's word `+0x42` removes the objects whose
location word `+2` equals the target's (the cell word only). That the target is the building class is Medium. In the RU
image the resolver has two direct callers: the blow HP application and Blizzard's damage object, so a building struck by a
blow or by Blizzard's object with a value not below its `+0x42` stops its casters; a script kill and the other spells'
hits are not shown to remove them (other paths not shown). — R2-ENGINE-1010, R2-ENGINE-1563

The objects are in the document's caster list and are saved and loaded with it; LOAD calls the list reader, whose bodies are unread (Medium that it rebuilds the list from the file).
— R2-ENGINE-1011

## Placement fields and catalog keys

What each placement names, per record type, in EN and RU unless stated.

**Type 4.** Element: x (+0x00) and y (+0x04), of which the loader keeps `x >> 8` and `y >> 8` (cell coordinates; the
EN location routine `L2.07308` keeps `(y << 8) + x`, which has the form of the plane key; EN body read, Medium that the
coordinates index the planes, R2-ENGINE-1802); kind
dword (+0x08); word +0x0c; owner word (+0x0e, low word of a dword); id (+0x12, i16). A kind of 0x21 or with bit
24 is stored as 0x21 and reads two more dwords; when the sum of their low bytes is above 0 they replace row
parameters 0 and 1 (the size, Medium; all 41 EN and 170 RU 28-byte elements of the census do, so the census
cannot show the other case). Shop, Inn and Pointer kinds select those classes; every other kind is a Building.
An unknown owner keeps the building and logs a warning. — R2-ENGINE-1288, R2-ENGINE-1289, R2-ASSET-237

Readers of each element field in the type-4 pass, per arm (EN and RU masked-equal; every load of the element in the
pass is followed by a dereference, and an element copy held elsewhere is not followed, Medium; other readers of map
list +0x390 are not censused, Medium):

| File offset | Field | Shop | Inn | Pointer | Building | Claim |
|---:|---|---|---|---|---|---|
| 0x00, 0x04 | x, y | location | location | location | location | R2-ENGINE-1802 |
| 0x08 | kind | class tests, kind byte | class tests, kind byte | class tests, kind byte | class tests, constructor argument | R2-ENGINE-1802 |
| 0x0c | word | not read in this pass | not read in this pass | not read in this pass | zero test; a zero is copied to Building +0x42, a nonzero value is not read | R2-ENGINE-1802 |
| 0x0e | owner word | owner resolve | owner resolve | owner resolve | owner resolve | R2-ENGINE-1802 |
| 0x12 | id | list compare, object +8 | list compare, object +8 | compare over map list +0x34c, object +8 | object +8 | R2-ENGINE-1802 |
| 0x14, 0x18 | extension (kind 0x21) | not read in this pass | not read in this pass | passed on | constructor arguments | R2-ENGINE-1802 |

Shop kinds are 0x22, 0x23, 0x5d..0x5f and 0x69..0x6b; Inn kinds 0x43..0x45, 0x63..0x65 and 0x6f..0x71; Pointer kinds
0x46..0x4b; every other kind takes the Building arm. — R2-ENGINE-1802

The kind byte indexes by position the collection the building initialiser reads (the row is at the array
pointer held at collection +4, plus 0x1c × kind, with no bound test): kind 0, or a kind above the collection's
count minus one, gives a Building with no row and no registration. Row parameters 0..5 fill the size, a byte, two words and two
dwords of the Building; for a Building-arm kind, word +0x0c equal to 0 clears one of those words (+0x42). That the
collection is data.bin Buildings and which title each parameter is are Medium. — R2-ENGINE-1289, R2-ENGINE-1802

In the installed maps (EN 59, RU 83: the `.alm` files of the root folder and the `.alm` members of `scenario.res`;
other archives including `world.res`, saves and user maps are not in the population) every kind lies in 1..179;
data.bin Buildings stores 180 rows, and that this is the collection the loader indexes is Medium. Every owner
word lies in 1..player count of its map, and no id repeats within a map. Word +0x0c is 100 in 3,056 of 3,270 EN
elements and in 5,744 of 6,081 RU elements, and 0 in 203 EN and 312 RU. — R2-ASSET-237, R2-ENGINE-1289

The kind's row is listed per kind in [the definition tables](map-definitions.md#type-4-kind-to-databin-buildings-row):
EN 176 distinct kinds in 3,270 elements, RU 178 in 6,081, every kind with a row. The 28-byte elements are all kind
0x21 (EN 41, RU 170); their two extension dwords are 2..11 in EN and 2..17 in RU, in the pairs the page lists, and what reads the footprint
they replace, or Building +0x42 which a zero word +0x0c clears, is Unknown. — R2-ASSET-317, R2-ASSET-318

**Type 5.** File +0x00 plus one becomes Player byte +0xa44, the colour (R2-ENGINE-963); +0x04 AND 0x2 (value 3) sets Player +0xa60 to 1; +0x08 is
not read by the Player builder; the name is 32 bytes; the low byte of each of the 16 diplomacy words fills the
session matrix, whose diagonal is forced to 0x12. No comparison of the player count with a constant was found in
the type-5 case, the Player builder or the diplomacy routine as read (the type-0 count reader was not read).
— R2-ENGINE-1291

File +0x08 is 5000 or 0 in every installed record; the editor gives a new player 5000 in a field whose file
offset was not read. — R2-ASSET-238, R2-ENGINE-1293

**Type 6.** A record without flag 0x10 names a Units row by `serverID`, scanned from the last row down to
row 1; a miss builds the unit from row 0 (Medium). Installed definition ids are 600..999 without flag 0x10
and 1..10230 with it. — R2-ENGINE-1290, R2-ASSET-239

The lookup compares the full dword id with Units parameter 0x37 or, with flag 0x10, Humans parameter 0x18, from
count - 1 down to 1, and skips a row whose row-test result is zero (read as a row without a parameter array, Medium); its helpers are read in EN (High) and by shape in RU (Medium; the Humans lookup helpers are equal in RU under the mask, R2-ENGINE-742).
Every installed record resolves in the `world.res` Data.bin (Medium that this is the table the game loads; EN 241 and RU 250 distinct pairs, none matching two rows; table in [the definition
tables](map-definitions.md#type-6-definition-id-to-databin-row)). The Data.bin collections have the shapes in
R2-ASSET-320. — R2-ENGINE-1512, R2-ENGINE-1514, R2-ASSET-316, R2-ASSET-320

**Type 9.** Mode below 4 with a nonzero cell, mode & 8 (a unit's spellbook, keyed by unit) and mode & 4 take
separate arms. The grammar and the first two arm tests are read in the committed stage windows; the mode & 4
arm lies past their end and rests on R2-ENGINE-1006 at its grade. Installed maps have no mode 8; every mode-4 key equals a type-4 id of the same map.
In mode below 4, the stage indexes point-list entries 0 (source +0x36/+0x37) and 1 (target +0x38/+0x39) through EN `L2.08347` with no bounds check and no trigger-cell substitution. At cell entry, the consumer passes record source bytes (+0x36, +0x37) unchanged through the cast wrappers and position setter without testing for zero, so zero source bytes place the temporary actor at cell (0, 0) with sub-cell (0x80, 0x80). The following application targets the entering unit on the unit-targeted path (state 0xd) or record target bytes (+0x38, +0x39) on the point-targeted path (state 0xe); projectile spells (`param5 == 2`, other than 0xa and 0xb) compute flight delay from the distance between actor position (0, 0) and the target position in EN `L2.03854` (RU parity of `L2.05953`/`L2.02599`/`L2.08352`/`L2.08354` and per-spell reads inside `L2.01692` Medium).
— R2-ENGINE-1292, R2-ASSET-240, R2-ENGINE-1698

**Cell events.** The stage hands bytes `+0x16`, `+0x18` and the low bytes of entries 0 and 1 as a six-byte event to
the writer `L2.03589`, which stores it at `+0x34..+0x39` of the cell's terrain record, creating the record when absent
and replacing an earlier event; script action 25 is the writer's only other direct caller. Event 23 (`0x17`) passes a unit
whose step ends at the cell centre to the cell move toward the target cell `(+0x38, +0x39)`, which refuses it when the
move test fails (Medium for units outside the three step callers; nothing observed in the original);
its power and source bytes are unread (Medium). Both roots ship 16 EN and 18 RU type-9 records (mode below 4, cell
nonzero) whose event byte is 23 and that the stage hands to the writer (Medium that a move follows), 5 EN and 6 RU of them with the target equal to their own cell; 18 event-23 cells per root lie one cell right and
down of a Teleport Pad's base cell. — R2-ENGINE-1776, R2-ENGINE-1778, R2-ENGINE-1779, R2-ASSET-387

**Editor catalogs.** The EN map editor reads `objects.reg`, `units.reg` and `structures.reg` under
`graphics\` and `templates.bin`. That it opens `world.res` was observed in an uncommitted scan and is Medium,
not reproduced. Which catalog index becomes a kind or a definition id,
and the dialog-set fields of a new placement, are Unknown; what the editor stores for a new element is in
[Editor write sequence](#editor-write-sequence). — R2-ENGINE-1293, R2-ASSET-257

<a id="not-yet-surveyed"></a>

## Authoring limits and Unknowns

Type 10 and 11 fields are decoded to the program fields that receive
them above. The 11c count reader is the state-15 arm (R2-ENGINE-518; that it
is the only reader is Medium). Order 0xf, from message type 0x24, reaches the
handler that sets state 15 (R2-ENGINE-607, R2-ENGINE-519; other producers of
order 0xf were not searched). A Pointer kind's literal registered bounds remain
Unknown (R2-ENGINE-605, amended by R2-ENGINE-849/850; the former 3x3 size label
is retracted). What an Inn offer is and which event builds the offer list also
remain Unknown (R2-ENGINE-536). So does the reader of unit-record
0x0a and 0x20 outside the two bodies read. EN and RU differ in one type-10 store: the
RU group object stores the group's spread at +0xbc, the EN object keeps 0
there (R2-ENGINE-511). With 0 the positioning routine gets the member's own position; which cell the member takes
after the template restore is Unknown (R2-ENGINE-1818). Seconds per container pass are Unknown (R2-ENGINE-1817). Type 12 is decoded above (R2-ASSET-084). The editor's
full-map writer declares type 0's 644 bytes as a constant while it writes 660
(R2-ASSET-231). Why the constant is 644 remains Unknown. See [Editor write sequence](#editor-write-sequence).
— R2-ASSET-023, R2-ASSET-024, R2-ASSET-231

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
