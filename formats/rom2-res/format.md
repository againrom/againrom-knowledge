<a id="rom2-res-container-ya1--identity-survey"></a>

# ROM2 RES archive (`&YA1`)

The preserved ROM2 archives use the same tail-registry container layout as
[ROM1 RES](../res/format.md). This identity covers the stored envelope and
node geometry; it does not equate their payload contents or all runtime
lookup behavior. — R2-ASSET-001

<a id="result"></a>

## Layout

All integer fields are little-endian. The header is 24 bytes, and each node
is 32 bytes.

| Offset | Type | Field |
|---:|---|---|
| 0x00 | u32 | Magic `0x31415926` (`&YA1`) |
| 0x04 | u32 | Root first-child index |
| 0x08 | u32 | Root child count |
| 0x0c | u32 | Root kind/flags |
| 0x10 | u32 | Registry byte offset |
| 0x14 | u32 | Node count |

File nodes reference payload byte ranges; directory nodes reference child
node ranges. The installed payloads lie between byte 24 and the registry.
Use the explicit offset and count, not an EOF-derived node count.
— R2-ASSET-001; shared layout: RES-HDR-002, RES-NODE-007

## Read and write order

Read the header, then the counted node array. Resolve payload and child
ranges using the node type. The shared container emitter order is header,
payloads, node array, with final registry offset/count in the header; its
detailed field rules are in [RES](../res/format.md#write-sequence).
Original ROM2 writer and malformed-input behavior are not independently
specified by the layout identity. — R2-ASSET-001

<a id="clientserver-split"></a>

## Installed resources

The archive names are `MUSIC.RES`, `graphics.res`, `main.res`, `movies.res`,
`patch.res`, `scenario.res`, `sfx.res`, `speech.res`, `video.res`, `world.res`
and `world_srv.res`. Recognized entry extensions include `.16`, `.16a`,
`.256`, `.alm`, `.bin`, `.bmp`, `.dat`, `.pal`, `.reg`, `.txt`, `.wav`,
extensionless entries, and the additional `.pkt` and `.smk` families.
The `.pkt` members are `data/itemname.pkt` in both world archives;
`.smk` members occur in `video.res`. — R2-ASSET-001

Both world archives have `data/{ai.reg,data.bin,itemname.bin,itemname.pkt,map.reg}`.
Data.bin is 149971 bytes on the client and 152527 on the server. The other
four matching names have equal lengths; payload equality is Unknown.
— R2-ASSET-012

<a id="population-note"></a><a id="not-yet-surveyed"></a>

## Unknowns

Complete client/server payload equality, additional archive-pair relations,
native malformed-input acceptance and writer behavior remain unspecified.
These resource facts describe one preserved ROM2 locale.

## Selected movie backing source

The selected EN/RU movie source keeps member base/length and a cursor.
Its zero-count read still seeks the backing stream to base plus cursor;
nonzero reads clamp to the remaining member length. Movie open then passes
the source's backing handle to SmackOpen. This does not establish that
codec reads pass through the member clamp or stop at the member end.
Among 99 selected numbered teleport SMK keys in each measured video.res,
only 01.smk exists. The 6753596-byte EN/RU payload hashes differ; equal
size does not imply media identity. Same-stem 476-byte REG companions
have equal hashes. Loose override/alias precedence, physical EOF and
truncation behavior remain Unknown. — R2-ENGINE-096
