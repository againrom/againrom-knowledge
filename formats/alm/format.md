<a id="alm-map-container-m7r--specification"></a>

# ALM map (`M7R\0`)

An ALM stores metadata, three cell planes, placed actors and structures, and
mission-script data in typed records. The file header and each record header
are 20 bytes in the version-990 form. Payloads start after their record
headers; there is no separate trailer. — ALM-HDR-001, ALM-FRAME-031 (amended; framing retained),
ALM-GRID-032

The primary loader requires Tiles and Altitudes. Other records have individual
absence, version and order rules. A map browser uses a different acceptance
test. — ALM-REQ-055, ALM-REQ-056 and ALM-ORD-057 (type-7 count clauses
partially retracted), ALM-RDR-059

All integer fields are little-endian. File/payload offsets are distinct from
the hexadecimal runtime member offsets used for loader destinations.

## Wire layout

| Offset | Size | Contents |
|---:|---:|---|
| 0x00 | 20 | Magic M7R\0, header length, dataSize, recordCount, formatVersion |
| 0x14 | Repeated | 20-byte record header followed by its payload |
| Record +0x00 | 20 | Tag u32, header length u32, payloadSize u32, typeId u32, opaque four-byte word |

At version 990 the metadata payload is 632 bytes; grids are 2WH, WH and WH
bytes; structure/player/unit records are 20(+8 extension),76 and 70 bytes.
Type7 has counted script arrays, type 8 has metadata-counted loot, and type 9
has counted caster records. The standard size is
`20+sum(20+payloadSize)`, with no trailer. — ALM-FRAME-031 (amended; framing retained),
ALM-SEC-004, ALM-PLACE-033, ALM-TRIG-044, ALM-SACK-065

## Read and write order

Read the header, then dispatch records by typeId. The primary loader gates
recordCount>=3 and version<=1001 and requires types 1/2. Metadata precedes
its dependent records; type 2 precedes type 3. For the complete version 990
form, write types `0,1,2,3,5,4,9,8,6,7`, derive the documented counts for types 5/4/6/8 from their
payloads and retain the documented caster/loot/unit ordering. A missing
record supplies no instances even when metadata has a nonzero count.
At version 1000 the main helper skips record-header reads while the editor
reads five words in a different order; the browser/landscape helpers follow
header length without testing this word. — ALM-HEADER-097
— ALM-REQ-055, ALM-REQ-056 and ALM-ORD-057 (type-7 count clauses partially
retracted), ALM-ORD-068, ALM-CORP-060

For complete primary-reader type-7 reads, each of the three arrays is bounded
by its own internal count, not metadata+0x28. Incomplete reads can retain old
bytes in the shared local and must be distinguished from wire counts.
— ALM-COUNT-195, ALM-STALE-196

## Reference map

| Reference | Contents |
|---|---|
| <a id="at-a-glance"></a><a id="trailer--there-isnt-one"></a><a id="invariants-hold-across-all-38-en-maps"></a><a id="notes--open-questions"></a><a id="structure"></a><a id="file-header-20-bytes"></a><a id="record-header-20-bytes--repeated-recordcount-from-0x14"></a><a id="record-roster--what-the-writer-emits-and-what-the-loader-requires"></a><a id="acceptance-contract--what-a-reader-must-and-must-not-require"></a><a id="end-of-file"></a><a id="reading-algorithm"></a><a id="write-sequence"></a><a id="unknowns-and-compatibility"></a> [Container and load rules](container.md) | Header fields, version/order rules, acceptance and emission |
| <a id="type-0-metadata-payload-632-bytes--alm-meta-008alm-meta-010"></a><a id="grid-layers-type1--type2--type3--alm-grid-012-alm-grid-013"></a><a id="type1-tile-word--terrain-resolution-romexe"></a> [Metadata and cell planes](metadata.md) | 632-byte metadata, tile/height/object planes |
| <a id="content-sections-type49--alm-cnt-017-alm-unit-018-alm-obj-019"></a><a id="type4--placed-structuresbuildings-alm-obj-019-alm-obj-034-alm-cls-036037"></a><a id="type5--playergroup-roster-alm-grp-020-alm-grp-041"></a><a id="type6--placed-units-alm-unit-018"></a><a id="type7--the-trigger-script-three-counted-arrays-alm-trig-044047"></a><a id="type8--authored-loot-and-type9--caster-payload-alm-sack-065-alm-trig-049"></a><a id="romexe-cross-reference-alm-code-023"></a> [Placements and mission records](placements.md) | Structures, players, units, scripts, loot and casters |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.

The record-header final word has no established numeric interpretation. The
selected editor writer supplies a local frame slot without initializing it;
within-map equality alone does not identify semantics. — ALM-HEADER-098,
ALM-WRITER-099, ALM-CENSUS-101
