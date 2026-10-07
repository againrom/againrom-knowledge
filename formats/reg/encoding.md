# Encoding and lookup

[Reference](format.md)

## Structure

A `.reg` payload is **itself an `&YA1`** (same magic as the outer RES container,
`REG-LOC-016`) but used as a **hierarchical key-value store**, not a file archive. After a
24-byte header it is a flat array of fixed 32-byte **records** forming a key/subkey tree,
then a `u32` pool length, then a **pool** holding string and array values (`REG-FMT-031`).

```
+---------+-----------------------------+-------+--------------------------+
| header  |   records (32 B each)       |poolLen|   string / array pool    |
| 0x18 B  |   key & subkey tree         | u32   |   DescText, File, arrays |
+---------+-----------------------------+-------+--------------------------+
0        0x18                    0x18+R*32   +4                          EOF
```

### Header (0x18 B) — `REG-FMT-031`

The header contains six u32 fields. Its first four fields form a virtual
root node without a stored name. Opening the registry supplies that name from
the file path. — REG-FMT-031

| Off | Field | Notes |
|-----|-------|-------|
| 0x00 | magic | `&YA1` = `0x31415926` LE; mismatch → the game's `"bad signature"` |
| 0x04 | root `value` | index of the root's first child record |
| 0x08 | root `size` | number of top-level children |
| 0x0C | root `kind` | **17** = `0x11` in installed files = subkey (bit 0) \| sorted (bit 4). The raw loader does not require this value (`REG-099`, `REG-101`). |
| 0x10 | `R` | total record count |
| 0x14 | pool waste counter | loaded/stored as registry `+0x30`; deletion and pooled replacement account bytes here — REG-104, REG-107 |

Then `R × 32 B` of records at `0x18`, a `u32` **pool byte length**, and the pool:

```
poolStart = 0x18 + R*32 + 4
registryEnd = 0x18 + R*32 + 4 + poolLen
```

## Record (32 bytes)

The record layout in REG-REC-032 is retained; its lookup clause is partially
retracted, as narrowed by REG-100.

| Off | Type | Field | Notes |
|-----|------|-------|-------|
| 0x00 | u32 | — | zeroed on insertion, retained by raw copy; no further meaning established — REG-104 |
| 0x04 | u32 | `value` | int32 · pool byte offset · child-block **start** index · **low dword of a double** |
| 0x08 | u32 | `size` | byte length · child count · **high dword of a double** |
| 0x0C | u32 | `kind` | bitfield, see below |
| 0x10 | char[16] | `name` | 15 significant characters + NUL |

Record `i` starts at `0x18 + 32*i`; its value is at `0x1c + 32*i`. These
REG-REC-032 offsets are retained; its lookup clause is narrowed by REG-100.

### Kind — a bitfield, not an enum (`REG-KIND-033`)

```
type       = kind & 0x0E          masked with `0x0E` and compared by every accessor
bit 0      = node is a subkey
bit 4      = children are sorted -> the lookup bsearch()es instead of scanning
bit 28     = insertion truncated a long name; a typed setter can replace the kind
bit 30     = deletion mark, tested by descent, endpoint and copy
```

Descent with another component requires `kind & 0x40000001 == 1`; endpoint
selection rejects `kind & 0x40000010 != 0`. Bit 0 alone does not reject a
terminal. Deletion marks bit 30, skips repeated accounting and recurses for
active subkeys, with a separate global early-stop gate. Copy omits marked
children and clears copied parent bit 4. Matching a name does not establish
acceptance by those later consumers. — REG-104

### Kinds and their storage (`REG-KIND-034`, producer absence partially retracted)

| Kind | Meaning | Storage |
|---:|---|---|
| 0 | string | `poolBase + value`, `size` bytes **including** the NUL |
| 1 | subkey | children = records `[value, value + size)` |
| 2 | int32 (signed) | `value` |
| **4** | **double** | **`value` = low dword, `size` = high dword — 8 bytes in the record, no pool access** |
| 6 | int32[] | `poolBase + value`, `size/4` LE int32s |
| 8 | string array | pool dword count, then NUL-terminated byte strings; count and record byte size are separate consumer bounds — REG-105 |
| 10 | double[] | `poolBase + value`, `size/8` LE doubles |

A key name does not determine its kind. `Mercenaries`, `InnNPC`, `InnMission`,
`EnableMercenary` and `AddTextDocument` can hold a bare int32 or an int32 array.
Integer-array getters accept a scalar as one element but read only its low
byte. `L07560`, `R1430` and `R1061`/`R1431` produce one-, two- and
four-byte elements; wider scalar destinations zero-extend the byte. For array
inputs they instead read one, two or four bytes of each four-byte pool item,
for unsigned `size >> 2` elements. The four-byte getters accept type 0 with
size below two as empty; the byte/word getters reject those controls.
— REG-KIND-056, REG-KEY-045, REG-106

Kind 8 has producer `L07399`, using string pointers with length metadata;
a new item requires `4 + sum(length+1)` bytes. Getter `L07400` sizes output
from the count dword and scans strings to the byte end. Count/string mismatch
can exceed the supplied output range. Native caller reach and string-object
construction beyond the named cuts remain Unknown. — REG-105

The type-8 converter skips four pool bytes and copies unsigned
`min(size-4,capacity-2)`, then places NULs at capacity minus two/minus one.
A shorter copy can leave a gap. Size below four underflows; small capacities
and the final destination NUL scan can exceed the supplied logical range.
Type 0 uses bounded string copying; type 2 converts the signed dword in base
10. Selectors 6/7 reach a throw, rather than returning the diagnostic as the
converted value. — REG-108

The game's own names for the types, from its error strings: *int*, *double*, *int array*,
*double array*, *string array or single string*.

### Kind 4 — the double (`REG-DBL-035`)

- Reader `R1397` masks the kind with `0xe`, requires it to equal 4, and loads the 64-bit float at `node+4`.
- Writer `R1398` stores `atof(text)` as a 64-bit float at `node+4` and sets `node->kind = 4`.

Kind 4 stores the binary64 value directly at record `+4`; it has no pool allocation. — REG-DBL-035

The direct setter `L07565` also writes both inline dwords. The string
converter calls a formatter with that pair, then converts the low dword as
an integer into the same buffer. A supplied formatter-return control proves
the later overwrite; native formatter behavior remains Unknown. Double-array
getter `L07567` copies unsigned `size >> 3` qwords, dropping partial final
bytes; setter `L07566` writes eight bytes per item. — REG-109, REG-110

### Pool encoding (`REG-VAL-025`, as amended)

Kind 0/6/10 records address `poolBase + value`; `size` is a byte length.
Values do not depend on a cursor or the preceding record. — REG-VAL-025

Pooled setters update size on growth but retain it on equal/shrinking
replacement, adding unused bytes to the pool-waste counter. The suffix and
larger extent can survive raw writer/readback. Existing matching kinds retain
flags; new typed records replace the whole kind. Byte/word source setters
zero-extend into four-byte pool entries. — REG-107

Record capacity grows when at most `count+1`, to
`capacity+(capacity>>2)+1`. Pool capacity grows when at most `used+request`;
outside its waste-sensitive compaction arm it becomes
`capacity+(capacity>>3)+request`. These are capacity rules, not wire size
limits. Overflow, allocation failure and native compaction remain Unknown.
— REG-111

### Tree walk (`REG-VAL-028`, as amended)

A subkey addresses the contiguous child range `[value,value+size)` in the
record array. The root has the same relation. — REG-VAL-028

## Text (`REG-TEXT-036`, `REG-TEXT-037`)

The loader transfers the record array and pool verbatim. The string accessor
copies from `poolBase + value`; it performs no code-page conversion.
— REG-TEXT-036, REG-TEXT-037

The unsorted name lookup's `_strnicmp` fast path folds `A`–`Z` only and leaves
bytes `>= 0x80` untouched. The sorted lookup instead compares raw bytes with no
case folding (`REG-100`). The CRT locale-dependent arm remains outside that
conditional fast-path statement.

Registry text has no established non-ASCII display mapping. Keep its bytes
distinct from any display encoding chosen by a consuming application.
— REG-TEXT-036, REG-TEXT-037

## Name matching, ordering and the 15-character clamp

The parent kind's bit 4 selects two different comparisons. With bit 4 clear,
`R0505` linearly scans children using `_strnicmp(query, child+0x10, 15)`.
With bit 4 set and a nonempty list, it copies at most 15 query bytes into a
NUL-terminated temporary record and calls `bsearch` with `L07393`. That
comparator compares unsigned bytes to NUL, **case-sensitively**. A root or leaf
case variant can therefore match in an unsorted container and miss in a sorted
one. Nonzero CRT locale state is not covered by the unsorted ASCII-fold result.
The lookup clause of REG-REC-032 and the comparator clause of REG-KEY-054
are narrowed by this result. — REG-100

The node-insert helper clamps a new name to 15 bytes plus NUL and marks an
overlong input with kind bit 28; typed setters may later replace the kind.
The raw loader does not clamp names. A 15-byte stored name
matches a longer query sharing that prefix on both lookup routes; a raw name
occupying all 16 bytes without NUL matches that query on the 15-byte linear
route but misses the sorted route's truncated temporary key. Arbitrary
unterminated names and comparison reads outside the record remain outside the
valid producer-name domain. — REG-099, REG-100

The installed registry roots set bit 4; their ordinary section children do
not. Section names therefore use sorted lookup, and per-key names ordinarily
use linear lookup. REG-KEY-054's unconditional comparator/writer clauses are
partially retracted; use REG-100's comparator rules.

- Sections are stored in **lexicographic**, not numeric, order — `MapObject1, MapObject10,
  MapObject11, …, MapObject2`. The order in which a loader's `"<Prefix>%d"` counter visits
  sections is **not** the record order.
- A container advertising bit 4 must order its children for the case-sensitive
  comparator. A container with bit 4 clear takes the linear route. — REG-100

Neither lookup route checks a child's kind before matching its name. The linear
route selects the first equal child; sorted duplicate blocks of 2, 3 and 4
equal names select indices 0, 1 and 1 in the original bsearch body. Lookup does
not enforce the parent subkey bit or skip a child carrying bit 30. The selected
integer getter checks only `kind & 0x0e == 2` after two name lookups; missing
names return its caller's default, while a mismatched value type reaches the
exception-throw boundary. These are accessor facts, not structural or
application acceptance rules. — REG-101

Some stored keys are 15-byte prefixes of longer authoring names:
`MinimalGuardRan` → `MinimalGuardRange`, `AddPictureDocum` →
`AddPictureDocument`, `ScenarioMission` → `ScenarioMissionCount`.
The stored prefix is the lookup name. — REG-NAME-055

## Raw parsing and writing of unrecognized entries

The raw loader reads six header dwords, the full `R*32` record block, a pool
length and the pool. It checks the signature and allocation results but does
not enumerate names, validate child kinds, enforce ordering or deduplicate
records. Unrecognized entries remain in the raw registry. Known lookups remain unchanged
when entries do not collide under the selected comparator and the advertised
order remains true. This does not establish application LOAD acceptance.
— REG-099, REG-100, REG-101

The raw writer sorts through `R1394`, then writes the supplied registry's
header, complete record array and pool. Unrecognized entries can survive this direct round trip. Sorting can alter lookup behavior: a lowercase
root matched the unsorted case-folding route before writing and missed the
case-sensitive route after writing set the root's sorted bit. The sort helper's
recursive candidate is indexed by `parent.value + parent.size`, rather than
its loop counter; it does not establish a conventional recursive sort of every
descendant. Internal insertion/copy helpers and an application's fresh-registry
producer are separate paths. — REG-102

## Read and write sequence

1. Read six little-endian header dwords; check magic `0x31415926`.
2. Read `R` records of 32 bytes from `0x18`.
3. Read u32 `poolLen`, then exactly that many pool bytes.
4. Interpret each value by its own kind. A subkey refers to records; string
   and array kinds refer to pool bytes; an integer or double is inline.
5. Resolve names with the parent kind's bit-4 comparator. Check record and
   pool ranges before dereferencing them in a bounded decoder.

To emit a registry, assign contiguous child ranges, encode each typed value,
build the pool and write the header, records, u32 pool length and pool.
Set each sorted flag only when its child block has the required unsigned-byte,
case-sensitive ordering. The original raw writer's sorting behavior is stated
above; it is not a general recursive normalization guarantee. Preserve unknown
values and the header's unnamed final dword when rewriting an existing store.
— REG-FMT-031, REG-KIND-034 (amended; storage clauses retained), REG-099, REG-100, REG-102
