<a id="reg--registry-nested-ya1-key-value-tree--specification"></a>

# REG registry (`&YA1`)

REG stores a hierarchy of named typed values. The header is followed by a
32-byte record array, a pool-byte count and a string/array pool. Integer fields
are little-endian. Names and string values are byte strings. — REG-FMT-031,
REG-REC-032 (layout retained, lookup narrowed by REG-100), REG-TEXT-036

The record layout in REG-REC-032 remains valid. Its lookup clause and the
comparator clause of REG-KEY-054 are narrowed by REG-100: sorted lookup is
case-sensitive; the unsorted fast path folds ASCII case.

## Wire layout

| Offset | Size | Contents |
|---:|---:|---|
| 0x00 | 24 | Six u32: magic, root child start/count/kind, record count R, pool waste counter |
| 0x18 | 32×R | Records: unknown u32, value u32, size u32, kind u32, name[16] |
| 0x18+32×R | 4 | Pool byte count P |
| 0x1c+32×R | P | String/array pool |

File size is `0x1c+32*R+P`. Subkeys reference contiguous record ranges;
strings and arrays reference byte ranges in the pool. Kind 4 stores a binary64
inline in value/size. — REG-FMT-031, REG-REC-032 (lookup clause narrowed by
REG-100), REG-KIND-034

## Read and write order

Read the six header words, R records, pool length and pool. Interpret values
by kind, then resolve names using the parent kind's sorted flag. To write,
assign child ranges and pool offsets, emit those same blocks and advertise
sorted children only when their case-sensitive ordering is correct. Preserve
unknown fields when rewriting. The raw writer's sorting is not a general
recursive normalization guarantee. — REG-099, REG-100, REG-102

## Typed consumer boundaries

Header `+0x14` maps to registry `+0x30`, the pool-waste counter. Kind 8 has
an original setter: a dword string count followed by terminated byte strings.
REG-KIND-034's producer-absence clause is partially retracted. Its getter
sizes output by count and scans using the record byte extent. — REG-104,
REG-105, REG-107

Pooled setters retain a larger existing size on shrink, so old pool suffixes
can survive raw writer/readback. Matching existing records retain flags;
insertion assigns a whole supported kind. Integer-array getters have one-,
two- and four-byte destinations, but their scalar arm keeps only the low
byte. Doubles remain inline; double arrays use eight-byte items. Capacity
growth is separate from used size. — REG-106, REG-107, REG-110, REG-111

Converter capacity is not a universal memory bound: the return scans for NUL,
and type 8 subtracts from unsigned size/capacity before copying. Binary64
conversion reaches a formatter then converts the low dword as an integer;
native formatter behavior remains unverified. Path, copy and deletion impose
separate flag predicates after lookup. — REG-104, REG-108, REG-109

## Reference map

| Reference | Contents |
|---|---|
| <a id="at-a-glance"></a><a id="cross-checks-that-the-framing-has-to-satisfy"></a><a id="structure"></a><a id="header-0x18-b--reg-fmt-031"></a><a id="record-32-bytes"></a><a id="kind--a-bitfield-not-an-enum-reg-kind-033"></a><a id="kinds-and-their-storage-reg-kind-034"></a><a id="kind-4--the-double-reg-dbl-035"></a><a id="pool-encoding-reg-val-025-as-amended"></a><a id="tree-walk-reg-val-028-as-amended"></a><a id="text-reg-text-036-reg-text-037"></a><a id="name-matching-ordering-and-the-15-character-clamp"></a><a id="raw-parsing-and-writing-of-unrecognized-entries"></a><a id="read-and-write-sequence"></a> [Encoding and lookup](encoding.md) | Header, record, kind, pool, comparator and writer |
| <a id="unitsreg--unit-classes-reg-units-018-reg-roster-019-partially-retracted-corrected-by-reg-roster-052"></a><a id="z-selects-the-drawable-class-reg-units-061-registration-scope-amended"></a><a id="the-loaded-unitsreg-class-record-reg-units-049-reg-units-050-reg-units-051"></a><a id="which-key-a-class-table-is-indexed-by-reg-key-044"></a><a id="sibling-registries"></a><a id="projectilesreg--and-why-its-id-is-an-address"></a> [Graphics class registries](classes.md) | Unit/object/structure/projectile keys and inheritance |
| <a id="the-other-registries-reg-cut-053-reg-sfx-057-reg-npc-058-reg-scn-059-reg-ai-060"></a><a id="map-editor-text-catalogs-editor-023"></a><a id="addhero-town-activation"></a> [Campaign, sound and other registries](resources.md) | Campaign, sound, NPC, cutscene, plus editor catalogs |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.
