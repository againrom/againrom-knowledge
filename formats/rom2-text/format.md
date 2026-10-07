<a id="rom2-text--identity-survey"></a>

# ROM2 text resources

ROM2 has CRLF-delimited string resources under `main.res:text/` and
`patch.res:patch.txt`. Their container form matches the positional resource
family. The native mission/dialogue conversion and selection paths below are
bounded separately from resource identity. Named character encoding and
rendered glyph identities remain Unknown. — R2-ASSET-010, R2-ASSET-011,
R2-ENGINE-049, R2-ENGINE-052

<a id="container-result"></a>

## Resource families

The fifteen ROM1 table names from `main.txt` through `credits.txt` remain
present, together with patch.txt. Additional resources include numbered
`missionNN.txt` files and `globalmap.txt`, `help.txt`, `itemserv.txt`,
`quest.txt`, `town.txt`, `docs/1.txt`. The installed mission numbers are sparse;
do not derive a contiguous list from the filename pattern. — R2-ASSET-010

<a id="encoding-result"></a>

## Stored-byte domain

The preserved ROM2 text-like archive nodes (`.txt`, `.ini`, `.lst` and
extensionless) contain high bytes in `0xe0..0xff` and 31 values of
`0xc0..0xdf` (`0xda` absent). They contain none in `0x80..0xbf`.
This is a property of that input population, not a valid-byte restriction.
— R2-ASSET-011

Only the `0xe0..0xef` portion overlaps the two source ranges of the ROM1
Russian display transform. ROM1 TEXT-CONV-001 is partially retracted and is
not a ROM2 conversion rule. A matching byte range does not establish a
matching encoding or atlas. — R2-ASSET-011, TEXT-CONV-001

<a id="not-yet-surveyed"></a>

## Reading and rewriting

Keep source strings as bytes and preserve resource/table ordering. Do not
convert using a guessed code page or reverse a presumed glyph transform.
Independent text-writer acceptance and consumers outside the paths below remain
Unknown. — R2-ASSET-010, R2-ASSET-011, R2-ENGINE-037, R2-ENGINE-052

## Campaign source and section spans

Campaign mode selects `main\text\mission%d.txt` using scenario current
location; other mode selects `main\text\quest.txt`. The whole source is
converted before CString storage. Plain briefing finds the first `#briefing`
substring and begins eleven bytes later. Failure and subobjective readers
format numbered keys and begin at match plus key length plus two. Each cuts
at next `#` or EOF; a missing key produces an empty string. This is a substring
parser, with no line anchoring or numeric token-boundary check.
— R2-ENGINE-049, R2-ENGINE-052

Failure caching begins at two. Empty entries two through four use text-table
slot `0x118 + number`; the first empty entry above four stops the cache.
Subobjective caching begins at zero and stops at the first empty entry. The
two installed mission 10 files each contain one briefing, two failure and
eleven event tags, and no subobjective section. An objective state value does
not supply a missing source label. UI failure selection uses `reason - 2`
as the cache index. Initial dispatch to the briefing request remains Unknown.
— R2-ENGINE-049

## Event and page selection

At UI message `0x433`, ordinary IDs format `event%d` and construct a dialogue
when UI bit eight is clear; otherwise the ID enters a deferred collection.
IDs 250, 253, 254 and 255 use separate branches; 255 forwards a failure reason.
Packet reachability and deferred replay are outside this bounded UI contract.
The dialogue loader lowercases its formatted `#<key>` search key, takes its
first substring match plus key length plus two, and cuts at next `#`. An
absent section creates an empty body. The copied body retains loaded bytes;
the `npc` control test uses a temporary lowercase copy. — R2-ENGINE-050

The part counter starts at zero and advances before reading. `<...>` headers
are scanned in source order; temporary lowercase headers match the substring
`part=<number>`. The first admissible alternative is selected. Hero flag bit
four satisfies `iamfemale`, its absence `iammale`; bit two satisfies
`iammage`, its absence `iamfighter`. `npc=` reads a decimal ID from five
characters. `npcalive=` requires a controlled-unit lookup result by NPC ID;
`npcdead=` requires no result. A combat-health meaning is not established.
— R2-ENGINE-051

After a matching header the body begins after LF and ends at NUL or before
the next header after backing up to CR. `sound=` reads a quoted name through
the next quote or an unquoted name through semicolon. Missing sound clears
the name; missing tune yields minus one. `tips=` updates a retained value.
Malformed headers, LF-only input and numeric-prefix collisions remain
acceptance Unknowns. — R2-ENGINE-051

Selected bodies enter a wrapped text control and can select an NPC portrait.
The button emits `0x470` to advance; exhaustion optionally posts `0x45b` with
tips and closes through `0x445`. Nonempty speech names form
`speech\<name>.wav`; default names exist for particular town-dialogue key
families. The two mission 10 files have no explicit sound headers and their
event keys match none of those default families, so this page route selects
no speech name. EN event7 has four parts; RU has three. Other missions and
other script audio routes are not characterized. — R2-ENGINE-051

## Mission conversion and font1 index

The native whole-file mission loader appends NUL and calls
`CharToOemA(buffer, buffer)` before storing the converted string. Static
evidence identifies that import, not a named Windows code page or its result.
The font selector is initialized from the last non-NUL byte of `main\id`
minus ASCII zero; it is zero in installed EN and one in installed RU.
Other selector writers remain unenumerated. — R2-ENGINE-052

When the selector equals one, the byte mapper adds `0x30` to `0x80..0xaf`
and `0x10` to `0xe0..0xef`; other bytes pass through. Under other selector
values all bytes pass through. Drawing then subtracts `0x20` and truncates to
a byte for the glyph/advance index. Width measurement uses the same mapper;
space and tilde paths have separate handling. — R2-ENGINE-052

The dialogue selects font1 with spacing two and the resource base
`graphics\font1\font1`; `.16` supplies sprites and `.dat` four-byte advances.
The two installed font1 nodes match across locales and the advance node has
224 entries. Pixel identity, Unicode glyph labels, the source encoding and
runtime Windows conversion remain Unknown. — R2-ENGINE-052

## Controlled NPC keys and membership

Selected EN/RU native paths populate the lookup word at client unit+0x1dc
from zero, a builder argument, a copied unit word, server-unit u16+0x14c or
an optional packet WORD. This NPC word differs from packet u16+0x12, the
entity key used to insert/remove the controlled map at receiver+0x9d0.
The selected synthetic builder body contains no direct map insertion.
— R2-ENGINE-069, R2-ENGINE-070

The selected packet branch creates an absent unit when flags&0xa0 is nonzero.
State byte 5, selected by unit i16+0x104 below -600, reaches removal only for
entity keys below 0x6000. A reset loop separately deletes/removes map entries.
Complete control transitions and combat-death-to-absence equivalence remain
Unknown. Preserve the lookup meaning of `npcalive=` and `npcdead=`.
— R2-ENGINE-070

Native name parsers transfer parameter slots 55/24 into server-unit+0x14c;
the second replaces a WORD above 10000 by (word/10)%1000. The association
with the Data.bin Units/Humans serverID schema is Medium. A native load-time
join, NPC registry section-ID equality and arbitrary authored key domains
remain Unknown. — R2-ENGINE-072

## First campaign dialogue and result routes

The selected fresh campaign takes initial raw type 2/ID 1 from the DLL.
Stage-10 EnterInn supplies packed 0x300a0205. TALK formats npc517talk10
from low-word ID and 12-bit topic, then immediately forwards that packed
value to TalkTo. Its three-bit kind 3 unlocks type-1 ID 10. Both town
resources contain that header once. Page acknowledgement is not the unlock
gate in this selected body. Availability supplies the location pointer and
its ID supplies the mission path; the two IDs are distinct. — R2-ENGINE-073

In both selected result ticks, the session failure register takes priority
over the win counter. Victory's general-table report uses slot 0x8c and
its button sends 0x41d. With campaign mode and win flag set, the client
calls DLL LeaveLocation and consumes its outputs. Failure's first report
instead uses cached mission text at reason minus two. Its selected close
choices lead to menu or save-slot selection; those branch bodies contain
no direct call through the bound LeaveLocation pointer. Other failure selectors, actual later loading and automatic
initial briefing dispatch remain Unknown. — R2-ENGINE-074

## First campaign movie table

The enabled native dispatcher indexes cutpaths.txt by DLL movie output.
The table's CRLF rows supply stems; installed index 1 is teleport. The
dispatcher tries files 01..99 under the selected stem, skips missing files
and stops on native player return zero. Mission-10 Leave supplies output 1.
This establishes a conditional first-movie route, not later movie ordering
or audiovisual equivalence between locales. — R2-ENGINE-075

## Selected success report sources

The native report stores the supplied general-table text pointer at +0x7c;
its actual vtable initialization method passes that pointer to a text
control. The selected startup source join maps flat slot 0x8c to main.txt
row 140. Both main sources have 369 CRLF rows. The report's dialogs
table indices 0x2b and 0x9a select rows 43 and 154 in 176-row sources;
their constructed commands are 0x41d and 0x446. Selected EN strings are
`Mission Completed`, `~Victory!` and `Continue`. RU row hashes and lengths
differ; a named encoding or rendered correspondence is not established.
The CRLF loader calls CharToOemA. Live table mutation, the full startup
schedule and complete generic input bindings remain Unknown. — R2-ENGINE-093
