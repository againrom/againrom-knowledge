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

The 0x431 arm passes cache entry `reason - 2`, through an accessor with no
bound, to the failure report, so a mission whose cache holds E entries
covers reasons 2..1+E: the `#failure<reason>` section, or text-table slot
`0x118 + reason` for an empty 2..4. A reason below 2 or above 1+E indexes
past the filled entries; what the screen shows then is Unknown. That entry
i holds the i-th filler result and that the report shows it under title
slot 0x8d are Medium (unread helpers). After one report the arm sets
+0x3e8 = 1 and shows no further report until +0x3e8 is cleared; where it is
cleared is Unknown. — R2-ENGINE-629

In both roots the 46 mission texts yield 214 cache entries, equal in key and
source: 3 to 8 per mission, with slot entries only in missions 10, 20 (3
and 4), 30, 32 and 40 (4). Every reason a triggered instant 5 stores has its
`#failure<reason>` section except mission 96's 5, past its three entries.
Mission 110 also holds a bare `#failure` tag. — R2-ENGINE-630

The events that fail a mission and the section each selects are composed in
R2-ENGINE-631.

## Event and page selection

At UI message `0x433`, ordinary IDs format `event%d` and construct a dialogue
when UI bit eight is clear; otherwise the ID enters a deferred collection.
IDs 250, 253, 254 and 255 use separate branches; 255 forwards a failure reason.
Packet reachability and deferred replay are outside this bounded UI contract.
The dialogue loader lowercases its formatted `#<key>` search key, takes its
first substring match plus key length plus two, and cuts at next `#`. An
absent section creates an empty body. The copied body retains loaded bytes;
the `npc` control test uses a temporary lowercase copy. — R2-ENGINE-050

The search runs on one shared text buffer that holds `town.txt`,
`globalmap.txt`, `quest.txt` or a mission file, whichever loader ran last;
which file it holds when TALK runs is Unknown. When the find returns -1 the
loader leaves a one-byte empty body and a clear control flag (EN High; RU
Medium, the RU find was not read). That the find returns -1 for an absent
section rests on an unread callee (Medium). The first-town TALK body
appears to ignore the builder's result and still call TalkTo (Medium).
Whether the page then draws an empty box or closes is Unknown.
— R2-ENGINE-343, R2-ENGINE-344

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
no speech name. EN event7 has four parts; RU has three. — R2-ENGINE-051

The page plays its name at SoundSpeechPos on the first page and on each
Next. An empty name takes the first family word of the section key in the
order talk, about, accept, reject, then keeper, man or guard, and becomes
`<family>\%04d%c%sp%02d` (`about\%04dt000p%02d`; keeper, man and guard use
`town\`): the header's NPC, a letter, the key's trailing digits padded with
`0` to three, and the part. The letter is `t` unless the header has a hero
token or the NPC is 21; then the hero's female and mage bits give `f`, `a`,
`m` or `s`. Only `town.txt` sections carry a family word (EN 176, RU 174; Medium: the
census model, while the unit lookup it needs is unread); the global-map, quest and
mission files request no speech through this page route; other script audio
routes are not characterized. Every requested name equals a `speech.res` key
in both roots after the census's mapping of `\` to `/` and its case fold (the
archive lookup is unread, Medium); EN holds 12 and RU 20 keys that no
page or shop name format can request. The shop name formats and the shop
route are outside this page. — R2-ENGINE-802, R2-ASSET-208

The ordinary event dialogue starts as a 580x240 rectangle at (30, 120). Its
placement rounds the width to 8 + 96k and the height to 104 + 64k (488x232)
and centres the box on the screen: (268, 268) at 1024x768 (the numbers on
screen are Medium). Enter and Esc post `0x470` as the button does. The modal
push sets UI bit eight (EN +0x404, RU +0x41c), so later events defer while
the box is open; in EN the bit also stops the edge routine and six panel
redraws per pass (seven call sites, two of them alternatives). Whether the simulation stops is Unknown. — R2-ENGINE-608,
R2-ENGINE-609

The first-mission signpost, 10.alm Pointer id 42, runs node 25: operation 2,
parameter 19, event 19 (the node step is Medium). `#event19` occurs once in
each mission10.txt; each body is one `<` header line and one text line, both
ended by CR LF. The EN line is the 29 bytes of "To Great Mage Skrakan's
Tower"; the RU line is 23 bytes. — R2-ASSET-152

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
its ID supplies the mission path; the two IDs are distinct. R2-ENGINE-073 is
partially retracted only for its world-map record-choice clause; the clauses
here stand. — R2-ENGINE-073

In both selected result ticks, the session failure register takes priority
over the win counter. Victory's general-table report uses slot 0x8c and
its button sends 0x41d. With campaign mode and win flag set, the client
calls DLL LeaveLocation and consumes its outputs. Failure's first report
instead uses cached mission text at reason minus two. Its selected close
choices lead to menu or save-slot selection; those branch bodies contain
no direct call through the bound LeaveLocation pointer. Other failure
selectors, actual later loading and automatic initial briefing dispatch
remain Unknown. The claim's
repeated mission-10 reason 19 is retracted; the failure sources are
composed in R2-ENGINE-631. — R2-ENGINE-074

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

## Inn card and unit-type rows

The inn statistics card uses zero-based `main.res:text/main.txt` rows 15..20
(attribute, Health and Mana labels), 83 and 84 (`Enemy`, `Mercenary`), 190
(`spellcaster`) and 191 (`Armor piercing`) in EN and RU. Unit type names are
`text/unitname.txt` rows indexed by the type number: rows 1, 15 and 24 are
Peasant, Crossbowman and Mage. — R2-ASSET-145

## Unit names from npcnames.txt

`main.res:text/npcnames.txt` has 2371 CRLF lines in each root (EN 2030 empty,
RU 2021). Lines 0..19 are `Mercenary` / `Наемник`, except RU line 6
`Катапульта`; lines 20..26 hold the hero and companion names; lines 2107..2110
are `Peasant` (RU `Крестьянин`, `Крестьянка`, `Крестьянин`, `Крестьянка`) and line 2111 is `Elder` /
`Староста`. — R2-ASSET-218

EN client unit messages 0x6c, 0x6e, 0x6f and 0x70 store the optional word as
unit +0x1dc and split a 24-byte packet name at `|` into +0xec and +0xf8. When
+0xec is then empty and +0x1dc is nonzero, +0xec takes `npcnames.txt` line
+0x1dc - 1 (zero-based line index Medium: accessor `L2.02400` unread; that all four
message types reach the step is Medium). A packet name wins over the word's line. — R2-ENGINE-908

RU runs the same step; its line index is Medium because the RU pool loader is
unread. — R2-ENGINE-909

The EN server sends the packet name only when the server unit's string +0x80
is non-empty; the RU sender is Unknown. — R2-ENGINE-910

The sheet's first-line test is strlen of +0xec in both roots, so the sheet
draws +0xec and +0xf8 when +0xec is non-empty and the type label otherwise.
The objective formatter names a unit by line +0x1dc - 1 when +0x1b8 & 0x11 is
set, else by its type name (kinds 1, 4 and 11 by the case tables; Medium). No EN
hover or dialogue-heading reader of +0xec was found among the 161 linear-sweep
sites, of which the sheet, formatter, message line and overlay were opened
(Medium; RU not censused); the hover line source is Unknown. — R2-ENGINE-911

Composed (Medium): the 34 placed headmen show line 2111, `Elder` / `Староста`;
their type label, `unitname.txt` line 1 (`Peasant` / `Крестьянин`), shows only
with an empty +0xec and word 0. The server name is witnessed empty only for 3
EN placements on maps 20 and 30; the RU sender is Unknown. — R2-ENGINE-912

Composed (Medium; IDs 1..7, IDs 8..20 unread): a hire with an empty server name
shows line ID - 1, `Mercenary` in EN and `Наемник` in RU (`Катапульта` for ID 7),
not the type label. No RU save exists. — R2-ENGINE-915

Unknown: the RU server-name sender; the writer of server `+0x80` for placed
units; the hover line source; the unit-message arm's earlier exits; hires
8..20. — R2-ENGINE-910, R2-ENGINE-911, R2-ENGINE-908, R2-ENGINE-914

## Menu, save and load page rows

`main.res:text/dialogs.txt` has 176 CRLF rows in EN and RU. The save and
load pages use zero-based rows 0 and 1 (`Ok`, `Cancel`), 24..28 (load page
texts and `<<< Empty Slot >>>`), 29..33 (save page texts), 144, 151
(`Load Saved Game`), 157 and 158 (`Delete` and its tip). The in-game menu
uses rows 34, 35, 37, 77 and 40 for Save, Load, Sound Options, Abort Game and
Return; the Cut Scenes popup uses row 153 (`View Cutscenes`).
`patch.res:patch.txt` has 119 EN and 214 RU rows; rows 50 and 51 belong to
the delete confirmation around the save label, and rows 102 and 103 to the
web server list popup. The row texts are not quoted. The RU row roles are
Medium.
— R2-ASSET-098, R2-ENGINE-441

## Automatic-casting labels

The installed `main.res:text/dialogs.txt` source uses zero-based rows 167..169
for EN `Own`, `Allies`, `Neutral`; row 170 is `AutoCast mode:` and rows
171..173 are `Minimum`, `Average`, `Maximum`. A CP1251 view of the corresponding
RU source gives `Своих персонажей`, `Союзников`, `Нейтральных`, `Автокастинг:`,
`Минимальный`, `Средний`, `Максимальный`. Selected RU options code requests
rows 171..173 for control `0x8d`. The source-to-native-table binding is Medium;
runtime code page and rendered RU glyphs remain Unknown. — R2-ASSET-085

EN stored ASCII labels and indices retain High. The explicit RU source view
and source-to-native binding retain Medium. Their first support predates the
160-window limit; the later completed comparison at a conservative minimum
of 169 windows is excluded from confidence. — R2-ASSET-085

The separate `text/main.txt` rows 218..220 name EN `Autohealing off`,
`Occasional autohealing on`, `Frequent autohealing on`; they are not established
as the three options labels by the selected binding. Their input-dispatch and
feedback correspondence remain Unknown. — R2-ASSET-085

The selected RU loader binds the options dialogs receiver to
main/text/dialogs.txt. Target checkbox IDs 0x83/0x84/0x85 use source indices
167/168/169; combo 0x8d uses 171/172/173. A separate key-feedback receiver
loads patch/patch.txt and requests indices 115..117. That patch text population
was not read, so its feedback labels remain Unknown. The main.txt autohealing
labels cannot supply those values. Native pool mutation, conversion and
rendered glyphs remain Unknown. — R2-ASSET-088

## Spell text

`spell.txt` holds 29 spell names and a final empty line in both roots; spell
row r is line r-1 on items, in spell lists and in the mask-driven `Spells: `
list. The EN names of lines 15 and 21 differ from the data rows' names.
`spells.txt` holds 24 hover entries and a final empty line, in spellbook
slot order; a fixed 24-entry executable table lists one row number for each
slot, and the match of those numbers to the `spells.txt` names, which rests
on an uncommitted name order, is Medium. Rows 9, 14, 15, 28 and 29 have no
slot. Hover lines add `main.txt` labels 117, 118,
123, 124, 182..187 and 217. Both files have the same line counts in EN and
RU and differ in text bytes. — R2-ASSET-103, R2-ENGINE-444, R2-ENGINE-449

## Mission speaker picture and crop operands

The read EN/RU dialogue route resolves a client unit by its NPC word. For
the non-paperdoll branch it forms `graphics\infowindow\<InfoPicture><N>.bmp`
from the unit type and its +0x28 value, omitting N when it is 1. The composed
route is Medium: string/registry helpers and the creature variant tail remain
open. Missing NPC-registry records do not by themselves select an empty
portrait. Failed bitmap-open handling reaches error/termination calls;
an unresolved actor's allocated surface has no established pixel contents.
— R2-ENGINE-582, R2-ENGINE-818

The read face builders pass native destination `(8,7)`. Source operands are
`(36,140,108,232)`, or `(X,144-Y,X+72,240-Y)` for type 0x40..0x66 when
its program-table pair has Y other than -1. These operand tuples are High;
their mapping to decoded BMP and final child pixels is Medium. The requested
boxes are 72 by 92 or 72 by 96 in an 88 by 108 result surface. Runtime clip
and final child-painter behavior remain Unknown. — R2-ENGINE-819

Mission 10 names keys 21,616,623,651,2000 in both languages. Conditional on
the measured placements reaching matching client units and variants, keys
616,623,651 select Ogre.bmp, Troll4.bmp, Dragon4.bmp; predicted top-origin
BMP boxes are `(10,44,82,140)`, `(15,21,87,117)`, `(7,49,79,145)`.
The join, variant conversion and pixel prediction are Medium. Exact hero
and paperdoll layers for 21 and 2000 remain Unknown. — R2-ENGINE-821

The read npc.reg and units.reg key-name census supplies no key containing
rect, crop or portrait. This is a bounded Medium negative; other names,
registries and reader routes remain open. — R2-ENGINE-820

## Dialogue page speech opening

The selected EN/RU dialogue open constructs part counter zero and calls
modal open, which invokes the real dialogue +0x80 method after a no-op
+0x78. The page increments to part 1; 0x470 advances it later. Explicit
nonempty sound= wins over the talk/about/accept/reject/keeper/man/guard
family default; the request is speech\<name>.wav at SoundSpeechPos.
The helper for NPC word 21 searches for a class-tested object with that word
and has a fallback call. EN's read fallback writes flags 0x56, selecting
speaker letter s. Its RU body is unread. Full living-hero identity and
composition through collection/modal helpers remain Medium.
— R2-ENGINE-802, R2-ENGINE-865

The sound loader/play code depends on an object pointer cleared on
selected EN/RU startup failures and release. Complete lifetime and audible
output remain open. Name resolution has a disk fallback after archive
miss in the read EN path; normalized archive-key equality is not the
whole contract. — R2-ENGINE-866, R2-ENGINE-870
