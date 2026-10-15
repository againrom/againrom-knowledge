# Pointer-hover help

## Shared timing and dispatch

ROM1 has a common cursor controller which polls the current pointer, accumulates
elapsed `timeGetTime` milliseconds, and requests text from the widget beneath
the pointer. Its normal hover request occurs when elapsed time first crosses
500 ms. The test requires `new >= 500`, `new < 25500` and unsigned `old < 500`.
It is a threshold crossing, not a repeated request on every subsequent frame.
The actual display occurs at the next admitted controller poll, so 500 ms is
the threshold rather than a measured native frame timestamp. — TEXT-HOVER-048

The controller walks children in reverse stored order, recursively, then tests
the current widget's screen rectangle. It calls the selected widget's text
getter at virtual offset `+0x14`, copies the returned string and displays a
nonempty result. A null/empty child result does not make the dispatcher ask the
parent for another string. The normal path requires a root widget, no active
marquee and the controller's help-enable flag. Getter-specific state checks
still apply. — TEXT-HOVER-048, TEXT-HOVERSET-049

An admitted pointer-position change hides help, clears elapsed time and records
a new clock sample. The observed polling branch does not do this while the
right mouse button is held. At accumulated elapsed time 25500 ms the visible
box is cleared, giving about 25 seconds after the normal reveal. The separate
hide routine, also called after mouse button messages `0x201..0x206`, clears
visibility without restarting elapsed time; a hidden box
does not automatically reappear at an unchanged pointer position. Entering a
marquee also hides it. These are the identified controller routes, not a claim
about every focus or modal transition. — TEXT-HOVER-048

## Text targets and sources

Indices below are zero-based. `main.txt` is `main.res::text/main.txt`, whose
local indices equal global indices. Other files use their table-local index.
Source availability and a getter branch establish a static route; native
visibility under every state is not implied. — TEXT-HOVERSET-049,
TEXT-HOVERTEXT-052

| Surface | Hover target | Source |
|---|---|---|
| Command panel | Eight command icons | `main.txt[0..7]` |
| Character panel | Open/close book and backpack; doll/statistics; menu | `main.txt[8..14]`, chosen by current state |
| Character navigation | Previous/next hero or portrait | `main.txt[52..53]` or `[121..122]`, by mode |
| Statistics card | Four attributes, health, mana, damage, attack, armour, defence, weight, sight, speed, skills, resistances and experience | `main.txt[155..180]`; monster weapon resistance `[188]`; stat-specific visibility/ownership gates |
| Statistics card | Monster's assigned spell list | `main.txt[192]`, joined with names from `spell.txt` and runtime spell bits |
| Precreation | Difficulty, four hero templates, continue/back and name field | `main.txt[247..256]` |
| Final character generator | Four attributes, free points and ten class-specific skill pictures | `main.txt[155..158]`, `[273]`, `[171..180]`; each attribute row also has a label-and-value and two cost rectangles (below) |
| Town | Shop, school, tavern, mission exit and menu hotspots | `main.txt[233..237]`, in hotspot order rather than index order |
| School | Current five weapon or magic skill icons | `main.txt[171..180]`, with class and visual-slot permutation `1,2,4,3,5` |
| Tavern | Candidate statistics and equipped items | Shared statistics getter and item formatter |
| Shop | Shopkeeper and four stock groups | `main.txt[61..65]` |
| Shop and inventory | Scroll arrows; backpack, transaction table and shelf backgrounds | `main.txt[54..60]` |
| Shop, backpack and equipment | Item under the pointer, or gold | Item formatter; `main.txt[74]` for the inventory gold sentinel |
| Spellbook | Available spell cell | `spells.txt` name plus `main.txt` field labels and the live values of the selected actors (below) |
| World map | Available site marker | `sites.txt[marker index]`, subject to discovery/mission availability |
| Map-selection list | Map row and three metadata columns | The row's description text, then `dialogs.txt[134]`, `[135]`, `[136]` by cursor column (below) |

Character and generator mappings are established by TEXT-HOVERCHAR-050.
Room, navigation and item-area mappings are established by TEXT-HOVERROOM-051.
The composed text and local-table branches are established by TEXT-HOVERTEXT-052.
Existing TEXT-UI-039 and MENU-COMBAT-019 still apply to their own narrower rows.

The inspected family comprises 96 aligned widget tables carrying the same
three base methods: 28 distinct `+0x14` getters, of which 21 contain specialized
text-return paths, six return zero, and one reads the widget's owned hint
string at `+0x3c` for 69 tables. That inherited getter suppresses its string
when widget flag `0x20` is set; the setter at `+0x18` copies its argument.
Computed replacement tables and control families with different base methods
remain outside the inventory. The six null getters do not prove that their
entire screens lack help. — TEXT-HOVERSET-049

## Spellbook caption

The getter reads the spell cell under the pointer and answers only when that cell's bit is set in
the selected-population spell mask. The text is up to five parts joined as `#`-separated lines:
the spell name from `spells.txt` with the mana line (`main.txt[117]`), then damage `[118]`,
range `[123]` and duration `[124]` lines, then one caption. A value line is omitted when its
maximum is zero, shows one number when the minimum equals the maximum and a range otherwise; the
duration is a stored value times 0.0625 with one decimal. The seven caption blocks `[182..187]`
and `[217]` are tried in the order 182, 183, 184, 185, 187, 186, 217 and share one destination, so
the last block with a value is the only caption shown. Captions 185 and 187 carry a percent sign,
183, 185 and 187 a leading plus. — TEXT-080

The values come from a fold over the selected actors that are `CUnit` instances with a spell mask:
damage is summed over actors, range and duration are minima and maxima, mana is the last actor's.
Each caption is a function of an actor level (skill plus stat minus 30, clamped to 0..100) and of
the spell id, defined for ten ids; ids 17, 27 and 28 are not in the book's 24-cell id table, so
the spellbook getter and its fold never fill caption 187. The separate item
formatter uses its descriptor's stored-power byte and the six-caption contract
below. — TEXT-081, TEXT-127, TEXT-133

An actor whose `+0x7c` is zero is skipped before it is counted. When the session's `+0x3dc` is not
1 and lacks bit 2, a counted actor sets the flags to 8 and the class tests and the per-cell fold
are skipped, so no value lines appear. The spell record's level is computed in byte arithmetic and
is a byte: it uses 100 when skill plus stat is below 30 and wraps again at 286 and above, unlike
the caption level. — TEXT-097

The record fields behind the value lines: byte `+9` is the Max Range column plus power/30
(power/3 for id 26), `+0xe` and `+0xf` the damage minimum and spread times power/30 + 1, plus word
`+0x10` a duration in sixteenths from the Spell Duration or Area Effect Duration column. — TEXT-096
TEXT-096's universal per-cast-at-actor-power clause is partially retracted;
its popup record-field formulas remain valid. Applied power and separate
damage/lifetime operands can differ from the popup. — TEXT-117

## Ordinary Scroll descriptions

The ordinary Scroll writer emits marker 51, kind 41 with the low byte of the
spell id, and tag 50 with the low byte of stored Effect power. Tag 50 constructs
and fills that Spell directly. It does not run the spellbook actor fold or
clamp this byte to 0..100. Ordinary construction separately admits ids 1..27
and powers 1..100 subject to its price window. Wider powers and id 28 in
descriptor controls are not ordinary stock. — TEXT-127, TEXT-134

Ordinary Scroll keys have packed class 14. That class selects the inline
fragment ` %s %s%s`, with main.txt[90], singular spell.txt[id-1] and
main.txt[91], after the keyed itemname.txt name. It also suppresses the marker's
magic heading. The class test is shared by MagicItems rather than restricted
to a Book subtype. Text bytes are localized inputs; there is no grammatical
case conversion in this formatter. — TEXT-128

The ordered suffixes are damage, range, duration and one special caption.
Every suffix begins with `#`. Damage uses main[118] and `#%s: %d-%d`, only when
the filled spread byte is nonzero. Range uses main[123] and `#%s: %d`, only
when the class-14 filled range byte is nonzero. Duration uses main[124] and
`#%s: %5.1f`, only when the final duration word is nonzero. Its five-character
minimum width can add a padding space to the literal's existing space.
This complete arm has no mana or caption-217 call. — TEXT-129

Let p be the unsigned power byte and T a truncating integer conversion.
The filled damage values are B=u8(T(p16*(p/30+1))) for positive p16, else 0,
and S=u8(T(p17*(p/30+1)-B)) for positive p17, else 0. The factor is stored as
binary64. The displayed endpoints are B and B+S, with no byte narrowing of
the sum. Range starts from u8(p6) and adds floor(p/3) for id 26; other ids
add floor(p/30) only when the base is nonzero. The range result wraps to a
byte. These are filled values, not raw table captions. — TEXT-130, TEXT-131

With positive p14, duration D is u16(T(min(1.05^p*3*16,65000))) for id 15,
or u16(T(1.025^p*p14*16)) otherwise. With no positive p14 but positive p11,
D=u16((p11<<4)+floor((p<<4)/10)); otherwise D=0. If global dword L06743
equals 4 and id is 20, the formatter replaces D with floor(D/7)+10. The
displayed value is then D*0.0625. The meaning and ordinary-context reach of
that global remain Unknown, so both variants are part of this contract.
Exact native floating and one-decimal rounding boundaries remain unproved.
— TEXT-132

The special helpers run in order 182,185,187,183,186,184 and replace one
shared destination; the last nonzero value wins. Their ordinary p=0..100
values and label formats are:

| main index | Spell ids | Value | Literal |
|---|---|---|---|
| 182 | 24; 7 and 28 | floor(p/15)+1; its negative | `#%s: %d` |
| 183 | 5,10,16,22 | floor(p/2) | `#%s: +%d` |
| 184 | 12,17 | -(floor(p/30)+1) | `#%s: %d` |
| 185 | 23 | floor(4p/5)+20 | `#%s: +%d%%` |
| 187 | 27 | floor(4p/5)+20 | `#%s: +%d%%` |
| 186 | 14 | min(floor(p/20)+2,7) | `#%s: %d` |

Other ids return zero, which suppresses the helper's write. The item arm can
show captions for ids 17 and 27 even though the spellbook cell table omits
them; it omits the id-18 caption helper 217. — TEXT-133

The complete numeric constructor-candidate envelope has 2,700 spell/power
pairs and 5,400 localized predictions. Binary64 and decimal80 arithmetic
models differ on 76 pairs, including id 3/power 10 and id 8/power 1. Encoded
formulas and field selection have High confidence; exact numeric predictions
have Medium confidence. Neither model is an original execution transcript.
Actual rendered cards and all placed/serialized/custom reachability remain
Unknown. — TEXT-134

## Attribute rows in character generation

Each of the four attribute rows is tested against five rectangles in order (17 distinct rectangles). The first returns `main.txt[155+i]`,
the second `[273]`, the third the attribute label and value as `label = value`, the fourth a
signed cost of the next point as a negative number and the fifth a signed refund of the current
point as a positive number. The two signed numbers use the point-buy cost function and are
comma-grouped in threes from the right once they exceed three characters (four with a sign
character): `-1234` becomes `-1,234`. — TEXT-082, TEXT-098

Rectangles are panel-relative and half-open. The value, cost and refund rectangles are 20 pixels
high at x 82..102, 107..127 and 132..152, with row tops 54, 86, 118 and 150. The second test is
one rectangle (46,181)..(123,203) shared by all rows, below the value, cost and refund rectangles. The label rectangle starts at x 16 and its
right and bottom edges follow font metrics. By the traced stores the hover answers only between
the start and the teardown of the owning screen. — MENU-066

## Map-selection list

Over the first 300 pixels of a row the hover text is the map's description. Over the three
columns that follow it is `dialogs.txt[134]`, `[135]` or `[136]`, chosen only by the pointer's
x position and so independent of the map. The columns display the map size and two dwords of the
`.alm` metadata (payload `+0x70` and `+0x74`). The loader returns success only for a map whose
first dword is above 1. — TEXT-083

The size column prints the map's first two header dwords, width and height, each minus 16. — MENU-067

## Inherited hint binding

A control's inherited hint is a CString filled when it is constructed, from a text parameter of
one of two base constructors, or later by the base setter. Of 172 constructor call sites
that forward a text parameter, 63 pass null, 51 pass the address of a zero-initialised data
cell, three pass an executable literal, 52 read `dialogs.txt` and three read `patch.txt` entries
52..54. Seventeen final classes of the 69 inherited-getter tables are reached; the rest of the
population is not traced. Eight setter calls in one routine pass `dialogs.txt[23]` and `[75]`.
— TEXT-084, TEXT-085

The eight setter calls are on the four controls of the Sound Options dialog. `dialogs.txt[117]`
is the hint of a text edit box and `patch.txt` 52..54 are the hint and caption of three Game
Options controls. — MENU-068

Of the 52 inherited-getter tables not reached, 46 are built with a null hint; none is among the
9 of 252 setter receivers that resolve, and whether any shows help is not established. — MENU-069

The 55 constructor sites that read a text table sit in 14 builder routines, and 114 of the 172
forwarding sites pass no hint at construction; a later setter can still write one. — MENU-096

One forwarding constructor creates children. Its hint goes to its own field and to two child controls; its
third child, a button, gets the empty cell. It sets flag 0x20 on one child, which makes the getter answer
empty, and pairs each setting with a write of flag 4 on that child. Five uses read `dialogs.txt` 111..115.
— TEXT-105

The hit test tries the last stored child first, recurses, and falls back to the control's own rectangle.
It and its callees read no flag word, so only the chosen control's getter can answer empty. — TEXT-106

## Presentation

The string's `#` bytes are line separators. The normal layout measures each
source line, keeps the maximum width and places the box above the pointer:
width is measured width plus 11 pixels, height is `14 * lineCount + 5`.
It shifts left at the right screen edge and down at the top screen bound.
Text uses `font2`, with a five-pixel left inset, four-pixel top inset and
14-pixel line pitch. The decoded path splits authored lines rather than
automatically wrapping them by words. Long descriptions already carry these
separators in both installed languages. — TEXT-HOVERPAINT-053

This lifecycle is separate from introductory `text/tips/*.txt` panels, whose
constructor, controls and persistent `TipsMode` gate are documented in the town
claims. Disabling an introductory panel is not evidence that the hover
controller's distinct enable field is disabled. `TOWN-184` is amended: only
the room named in one aside of that claim is retracted, and the operand-order
finding cited here stands. — TEXT-HOVERPAINT-053, TOWN-184, TOWN-185, TOWN-186

## Limits

The timings, branches and resources above come from static executable/data
inspection and paired resource measurements. Native frame timing, all focus
transitions, all inherited hint assignments and a universal inventory of every
possible control remain unproved. A title or tooltip-shaped string in a table
alone does not establish a control binding. — TEXT-HOVERSET-049
