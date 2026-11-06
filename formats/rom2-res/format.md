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

`sfx.res` holds the unit voices as `characters/<voice>/<gesture>/<n>.wav`
(gestures select, attack, move, swarm with n = 1..4; retreat, defend and
pickup are built with n = 1) and `characters/<voice>/easy.wav`, `hard.wav`,
`dead.wav`. EN has 88 such keys, gesture files for six named voices only,
and also `aldor` defend and pickup with n = 2, which no built name reaches;
RU has 198, gesture files for all twelve named voices and easy, hard and
dead for the six generic ones too. — R2-ASSET-185

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

## Selected speech lookup and record projection

The read EN archive component parser accepts slash and backslash. Child
comparison is byte-sensitive after a prior normalizer; selected EN/RU
zero-locale case branches lowercase ASCII A..Z. The EN directory helper
tries disk search bases after archive miss, with its concrete file method
calling CreateFileA using read/OPEN_EXISTING for flags 0. Full mount
precedence, locale behavior and RU's concrete disk opener remain Medium
or Unknown. — R2-ENGINE-870

Both clients mount ten names at startup through one routine, in order: a
built `cs.res` path, `main.res`, `patch.res`, `world.res`, `music.res`,
`video.res`, `sfx.res`, `movies.res`, `scenario.res`, `speech.res`; then
they read `update.lst`. An exception-handler block after the `music.res`
mount clears the music-available word; that a failed mount reaches it is
Medium (mount body not interpreted).
`graphics.res` is not among them. Which archive wins when two hold one
name is Unknown. — R2-ENGINE-921

Each root's 491 itemname-record projection through the shop speech
formula forms 852 distinct normalized relative names, 283 present and
569 absent in that root's speech key set. The projection spans all three
shop stems and ordinary item parts 1..3. It is not the admitted/mutated
stock population and proves neither universal native coverage nor 569
actual failed requests. — R2-ASSET-212
