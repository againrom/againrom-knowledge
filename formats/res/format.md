<a id="res--lm-container-ya1--public-functional-specification"></a>

# RES / LM archive (`&YA1`)

A RES/LM archive stores file payloads and a tree of named 32-byte nodes.
The 24-byte header gives the node array's position and count. The inline
[REG store](../reg/format.md) uses the same magic with a different layout;
choose the grammar from the containing resource's role. — RES-HDR-002,
RES-NODE-007, RES-SCOPE-015

## Overview

A tail-registry RES/LM archive consists of:

```text
+---------+---------------------+-----------------------------+-------+
| header  | payload data        | node registry               | slack |
| 24 B    | [0x18, regOffset)   | nodeCount * 32-byte nodes   | 0..n  |
+---------+---------------------+-----------------------------+-------+
```

The registry location and size come from header fields, not from EOF. Shipped files can
contain trailing bytes after the registry, so deriving `nodeCount` from file length is
not compatible with all observed releases. Bytes past the node array are unreachable to
the original reader. — `RES-HDR-003`, `RES-HDR-004`, `RES-GEOM-028`, `RES-OPEN-026`,
`RES-OPEN-027`

The same container is seen under more than one extension, including a valid empty
archive with no nodes. The extension plays no part in the binary: one reader validates
the magic. — `RES-SCOPE-010`

## Header

All integers are little-endian.

| Offset | Type | Meaning |
|---|---|---|
| `0x00` | `u32` | magic `0x31415926` (`&YA1` as bytes) |
| `0x04` | `u32` | root node `off`: first top-level child index |
| `0x08` | `u32` | root node `size`: top-level child count |
| `0x0C` | `u32` | root node type/flags; bit `0x10` means children are sorted |
| `0x10` | `u32` | byte offset of the node registry |
| `0x14` | `u32` | number of 32-byte nodes |

The header is 24 bytes and payload data begins at offset 24. The magic value is the only
field the original validates. — `RES-HDR-002`, `RES-MAGIC-001`

The header's first 16 bytes have the same logical shape as a node's first 16 bytes, so
the header functions as the tree's virtual root node. — `RES-NODE-016`,
`RES-HDR-017`, `RES-HDR-018` (its `strncmp` label for the linear branch is superseded
by the case-insensitive CRT `_strnicmp` of `RES-LOOKUP-023` and `RES-041`; the bit-4
branch stands)

`0x08` is the root's child count. `0x04` is a node-array index; it is
not a count or checksum. `0x0c` carries the directory type and optional sorted
flag. — RES-HDR-005, RES-HDR-012 (its pointer to a `rom.exe` writer routine is
superseded: `rom.exe` holds no tail writer, and `0` versus the owned-node count
is the external packer's roots-first or roots-last order, RES-HDR-017,
RES-GEOM-028), RES-HDR-013 (its writer-generation reading of bit 4 is
superseded by RES-HDR-018's children-sorted flag)

Header bit 31 also controls original open: clear seeks to `regOffset`, set
reads the table at the current stream position after six header dwords.
Bounded one-bit controls establish this flag gate. Neither the complete header
kind nor each node kind is validated as a two-value enum. — RES-039

## Node record

Each registry node is 32 bytes:

| Offset | Type | File node | Directory node |
|---|---|---|---|
| `0x00` | `u32` | reserved | reserved |
| `0x04` | `u32` | payload byte offset | first child node index |
| `0x08` | `u32` | payload length | child count |
| `0x0C` | `u32` | ordinary base type `0`; consumer flags apply | ordinary base type `1`; `0x10` selects sorted children; consumer flags apply |
| `0x10` | `char[16]` | NUL-terminated name | NUL-terminated name |

Observed writers may leave non-semantic padding values after a short name. Readers
should use the terminator/name bound rather than treating padding bytes as content. —
`RES-NODE-007`, `RES-NODE-008`, `RES-NODE-019` (its clause that the four scanned functions
are all that touch a node is partially retracted; their absence of `[node+0]` reads stands)

`0x00` is reserved. It is zero on every node of every examined tail-registry archive,
which refutes reading it as a per-node hash, id or checksum, and no instruction in the
four scanned lookup, path-walk and finalize-sort functions reads it. Whether another
node-touching function reads it is Unknown. — `RES-NODE-011` (its never-read clause
narrowed to the scanned functions), `RES-NODE-014`

<a id="minimal-read-algorithm"></a>

## Read sequence

```text
require u32(file, 0x00) == 0x31415926

rootOff     = u32(file, 0x04)
rootCount   = u32(file, 0x08)
rootFlags   = u32(file, 0x0C)
regOffset   = u32(file, 0x10)
nodeCount   = u32(file, 0x14)

for i in [0, nodeCount):
    node[i] = parse32(file[regOffset + i*32 : regOffset + (i+1)*32])

root.children = node[rootOff : rootOff + rootCount]
```

A file node addresses `file[node.off : node.off + node.size]`. A directory node addresses
`node[node.off : node.off + node.size]`.

Defensive implementations should bounds-check these ranges even where the original is
permissive; such hardening is a safety choice, not an additional format invariant.

## Directory lookup

A directory's `0x10` flag selects sorted lookup; clear selects a linear scan.
The sorted comparator compares unsigned bytes without case folding. The
linear comparator folds ASCII under the named no-locale state.
RES-LOOKUP-023's unresolved-comparator clause is partially retracted; its
linear evidence stands. RES-HDR-018's `strncmp` label for the linear branch is
superseded; its bit-4 branch stands. — RES-HDR-018, RES-LOOKUP-023, RES-041

Path processing has several important compatibility properties:

- both `/` and `\\` are accepted as component separators;
- archive-manager input paths are lowercased before resolution in the reached path;
- sorted lookup compares preprocessed query bytes to stored bytes without folding;
- unsorted lookup folds ASCII under the named no-locale state;
- no Unicode normalization or general code-page conversion is part of the archive
  grammar. — `RES-TEXT-021`, `RES-PATH-025` (partially retracted), `RES-CASE-036`, RES-041

Manager queries `R/KEY` and `r/key` both become `r/key`. Sorted lookup then
matches stored `key` and misses stored `Key`; unsorted lookup matches either.
The query fold does not modify the stored name. — RES-041

After name matching, descent with another component requires
`kind & 0x40000001 == 1`; the endpoint rejects bits 4/30 but not bit 0 alone.
Deletion marks bit 30; shared copy omits marked children. Those decisions
differ from raw lookup acceptance. The later wrapper applies the existing
bit-29 loose-file sentinel. — RES-040, REG-104, RES-MASK-035

RES-CODE-020's `R1431` subkey-creator label is partially retracted: that
shared-library address is a REG array getter. It does not define RES payload
representation. — RES-CODE-020, REG-106

The read path performs no byte-to-character conversion. Names remain byte
strings; no general non-ASCII code-page mapping is established.
— RES-TEXT-022, REG-TEXT-036

<a id="container-digest-versus-selected-member"></a>

## Archive identities and resolver

An opened archive has an identity derived from its archive filename. Resource paths are
normally of the form:

```text
<archive-identity>\<path-inside-archive>
```

The resolver searches registered archives in registration order and returns the first
match, then falls back to configured loose-file directories. Because normal paths carry
an archive identity, the shipped archives mostly form disjoint namespaces rather than a
single overlay-by-priority system. — `RES-SET-032`, `RES-ORDER-033`, `RES-IDENT-034`

The leading path segment is compared case-sensitively with the archive's
stored name; subsequent components use the parent-selected comparator.
The resolver lowercases the path before these comparisons. Its loose-file
tier uses the process working directory captured before entry.
— RES-IDENT-024, RES-DIR-037, RES-041

The original also supports an update-list mechanism that can mark an archive node so
resolution falls through to a loose file. A compatible implementation may model this as
an explicit override from archive member to filesystem resource. — `RES-MASK-035`

<a id="acceptance-versus-corpus-invariants"></a><a id="publication-boundary"></a>

## Validation

Use explicit registry offsets and counts. Payload packing, tree reachability
and lack of cycles are properties of the installed archives, not checks made
by the original loader. EN and RU header values and registry slack can differ.
— RES-ACCEPT-031, RES-GEOM-028, RES-HDR-029, RES-HDR-030, RES-TREE-009

For a structurally bounded archive, check ranges without integer overflow:

```text
regOffset >= 24
regOffset + nodeCount*32 <= fileSize
file node: off + size <= regOffset
directory node: off + size <= nodeCount
```

These are structural safety checks; the original loader does not enforce all
of them. Bytes following the registry are not required to be absent.

## Two `&YA1` families

The same magic is used by two structurally distinct storage families:

1. this **tail-registry RES/LM archive**, where `0x10` is a registry byte offset;
2. the **inline REG-style store**, where the same position belongs to a different record
   framing.

Do not select the parser from the magic alone. Use the surrounding file/context and the
appropriate specification. — `RES-SCOPE-015`

<a id="writer-guidance"></a>

## Write sequence

A conservative writer for the tail-registry form should:

1. reserve/write the 24-byte root header;
2. append payload bytes and record each file node's byte range;
3. emit a flat 32-byte node array with directory child ranges;
4. set the root child range in the header;
5. write `regOffset` and `nodeCount` explicitly;
6. set the sorted-child flag only when the relevant child range is actually sorted under
   the lookup comparison used by the consumer.

There is no interoperability need to reproduce incidental padding/slack values from the
original packer.

## Unknowns

The complete malformed-input behavior and a general non-ASCII entry-name
mapping remain unspecified. Reserved node words have no established
semantic value beyond the writer and lookup rules above.
