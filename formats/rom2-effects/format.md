# ROM2 unit Effect records

<!-- machine-page -->

This machine attaches, renews and ends an Effect on a unit. Its direct stores
are known for EN and RU; source admission, list cleanup and final recomputed
characteristics have separate boundaries. R2-ENGINE-1164..1171.

Compact field labels such as word40 and byteF4 use hexadecimal offsets. IDs
and counts are decimal unless prefixed0x.

## Record and identity

| Coordinate | Read use |
|---|---|
| Effect u16 +0x0c | Attach identity and unit mask bit index, with x86 shift count modulo32 |
| Effect u8 +0x3c | Tag index0..49 |
| Effect u8 +0x3d | Entry flags; masks1/2 retain an Effect, mask2 selects continuous timer behavior |
| Effect s16 +0x40 | Magnitude selected by renewal's signed maximum |
| Effect u16 +0x42 | Remaining duration selected independently by zero-extended maximum |
| Effect dword +0x44 | Source pointer copied for a new record, retained from the old record on renewal |
| Unit +0x20 | Standing-Effect list; different base from an item's child list at item+0x20 |
| Unit dword +0x144 | Effect ID mask |
| Unit dword +0x150 | flag dword passed to the notify call (called dirty flags in R2-ENGINE-454, Medium) |

The copy explicitly transfers tag, flags, both words, ID and source after its
unread base copy. A renewal changes only the two words directly, surrounded
by old end/start. R2-ENGINE-1164.

## Attach outcomes

By the code read in both roots, the direct attach branches are:

| Input | Local outcome |
|---|---|
| Entry flags&3 =0 | Apply incoming slot40 directly, notify, retain no new record in this arm |
| Entry passes; incoming20 finds28 or incoming28 finds20 | End (unlink Medium) the first opposite; clear its mask bit; return without copying, linking or starting incoming; a second standing opposite record is not examined |
| Entry passes; same u16 ID found | End old; signed max magnitude; u16 max remaining duration; start old unless ID6 |
| Entry passes; no same ID | Allocate/copy; start unless ID6; link; set mask bit; send for nonzero ID |

Tag and source do not form part of the direct same-ID comparison. List
iteration, mutators, base copy and delivery remain Medium. Allocation-failure
recovery is Unknown. R2-ENGINE-1164.

## Start, end and timer

By the code read in both roots, the direct slot and timer operations are:

| ID/path | Direct operation |
|---|---|
| ID6 start | Health-word subtraction after optional resistance scaling and source calls; initial attach/renewal skips start |
| ID6 end | No direct characteristic operation |
| ID15 with unit virtual30=0 | Start subtracts magnitude<<8 at wordA4; end adds it |
| ID14 with unit virtual30!=0 | Start adds magnitude<<8 at wordsA4/E4; end subtracts it |
| Other IDs | Dispatch tag with polarity+1/-1, then unit virtual58 |
| Continuous timer | Before decrement, start at duration divisible by 8; omit end at expiry |
| Duration>9600 | Return without decrement, after any continuous start |
| Duration0 | Decrement wraps to 65535; a later timer call is above 9600 |
| Decrement reaches 0 | Non-continuous end; clear mask; removal send for nonzero ID; set ended flag |

Timer cleanup/unlink and cadence are unread. These are High direct-body
facts; final characteristic and delivery stay Medium or Unknown.
R2-ENGINE-1165, R2-ENGINE-1166.

Generic tag dispatch covers50 keys. It uses signed word40 when flags&7 is
nonzero, otherwise dword40, multiplied by polarity. Additive fields, upper
and lower clamps, target predicates, unread key38..42 callees and packed
overwrites are distinct operations. Keys44..48 overwrite bytesFB/F9/FA from
the original value even for polarity-1. EN alone clears signed wordD8>24
before generic tag dispatch; the RU body lacks that preamble. All generic
arms reach unread unit virtual58. A universal inverse or final-value formula
does not follow. R2-ENGINE-1166.

## Two effects and source boundaries

By the code read in both roots, negative renewal keeps the less negative
value: max(-5,-2)=-2. Duration is
independent, so one cast can supply the magnitude and another the remaining
duration. The renewal branch retains the old record and does not add the
incoming object, so it supplies no separate standing-record expiry. The
row21/29 arm at power0/15 gives magnitudes1/2 or-1/-2; numeric cast examples
remain Medium because the Effects-string key speed (17) is assumed,
recompute, upstream accessors and some numeric helpers are unread. Read the
Effects strings of data.bin rows21/29 to settle the key. Displayed speed is
Unknown. R2-ENGINE-1167.

By the code read in both roots, different IDs can coexist. When an opposite
20/28 record is found, its slot44 is called and mask bit cleared; incoming is
not copied, linked or started. Unlink rests on an unread helper (Medium), and
a second opposite record is not examined. IDs21/29 combine only under the
assumed key speed (17), nonzero virtual30 and preserving recompute.
IDs14/15 have opposite virtual30 predicates,
so one stable predicate admits only one of their direct operations even when
both records stand. The pair list is bounded, not an exhaustive census of
every source on every characteristic. R2-ENGINE-1168.

Light (ID15) and Darkness (ID14) carry signed magnitude P/30+1 and -1-P/30, duration 16 and
mark 2*row+8, renewed at every area pulse (Medium) (R2-ENGINE-1236). Virtual30 is 1 for Human and
Humanoid and 0 for Unit, Token, Building and Pointer, and in RU for Shop and Inn
(R2-ENGINE-1237). Darkness therefore lowers a Human's or Humanoid's sight byte +0xa5 and
word +0xe4 by 1+P/30, and Light lowers a Unit's sight byte by P/30+1, each when +0xa5 is at least that
amount; below it the Effect arms wrap the byte modulo 256 and no arm clamps, and whether a reader
clamps is not read except the ring walk's limit of 19 rings (R2-ENGINE-502). Light on a Human and
Darkness on a Unit change no sight. Neither arm tests the owner (R2-ENGINE-1238).
The 17 read sites of +0xa5 are the ray scans (one with no found caller; whether it runs is
Unknown) and the planner, movement and group radii; five routines overwrite the word, and +0xe4
has no unit reader outside the Effect arms in the direct-displacement census (R2-ENGINE-1239).
The same sight decides Prismatic Spray's extra victims. They are hostiles (bodies only when no living hostile is in sight) that the caster's `+0x70` members see, ranked by edge distance from the caster and then by turn, and admitted up to the
range byte (R2-ENGINE-1766, High / Medium).
Bit 0x40000 sends the word to the unit's own player; for Human and Humanoid units with word +0xe in
0x21..0x3f another player may receive it (Unknown), with a further cut for Units below key 0x6000
(R2-ENGINE-1240). Expiry runs the end slot and resends; a
save keeps the lowered word (R2-ENGINE-1241). The per-class result is Medium
(R2-ENGINE-1242) (High for the arms, Medium for the reach, the delivery and the fog link).

By the code read in both roots, a parsed timed item Effect that retains
constructor ID0 meets another ID0 even with another tag; the renewal branch
keeps the old tag and stores the merged words in that record.
This is Medium: blank and item-use bodies are read, parser prefixes and
helpers are incomplete, and shipped timed-potion data was not read. Actual
potions' complete identity/population remains Unknown. R2-ENGINE-1169.

Item child-start/end invoke slots40/44 directly. Conditional additive item
and spell contributions can be separated algebraically; clamps, EN reset and
unread recompute prevent proving general independence. Complete item/spell
characteristic behavior remains Unknown. R2-ENGINE-1170.

## Writer boundary

Identified field writers include construction, copy, parser, speed-arm fill,
same-ID renewal and duration tick. The dword copies overlap both words.
Dispel, save/load, server-message and all other writers were not censused.
Completeness remains Unknown. R2-ENGINE-1171.

## Speed and resistance after the fold

In RU the fold `L2.00102` adds an effect modifier to the Human speed word +0x8c; a sum below 0 zeroes the
modifier and keeps the negative word. In the bodies read the speed word has a floor of 6 only when load reaches
capacity and no upper bound; when the step takes that word it clamps it to 1..63 (Medium). The recompute clamps
each resistance word after the fold to 0..min(Spirit/2 + 70, 100); the fold's effect on the resistance words is
Unknown, and the damage resolver bounds no resistance word but raises its result to 0 (R2-ENGINE-1830).

## Effect words in computer-unit scoring

Bit 0x40000 of the unit effect word +0x144 is effect id 18, and a candidate with it set gets +0x7f in the first ROM2 scorer. The bit 0x40000 of the field mask +0x150 is a separate flag that the slot routine sets when a Light or Darkness field changes (R2-ENGINE-1496).

## Area and cloud spells: cells, pulses, sight timeline

Rows of the Spells table with Distribution above 1 are 2, 3, 6, 7, 9, 14, 15 and 17 and no others; the cells,
pulses and occupant entry per pulse are (R2-ENGINE-1536, High for the census and the cell counts, Medium for the pulse counts, the multiplicity and the classes):

| Row | Spell | Cells per pulse | Pulses | Entry applied | Claim |
|---|---|---|---|---|---|
| 2 | Fire Ball | 3 by 3 = 9 | 1 | +4, +8, +0xc | R2-ENGINE-1536, R2-ENGINE-1537 |
| 3 | Wall of Fire | 10 or 9 by facing | 10 + ceil(floor(16P/10)/16), 16 updates apart | +4 | R2-ENGINE-1536 |
| 6 | Poison Cloud | diamond, 21 | 10 + ceil(floor(16P/10)/16), 16 updates apart | +4 | R2-ENGINE-1536 |
| 7 | Blizzard | 1 of 36 | 32, 3 updates apart | +4, +8, +0xc | R2-ENGINE-1533, R2-ENGINE-1668 |
| 9 | Acid Stream | cone, 25 or 21 in total | 6, at updates 1, 4, 7, 10, 13, 16 | +4, +8, +0xc | R2-ENGINE-1536 |
| 14 | Darkness | diamond, 57 | 15 + ceil(floor(16P/10)/16), 16 updates apart | +4 | R2-ENGINE-1536 |
| 15 | Light | diamond, 57 | 15 + ceil(floor(16P/10)/16), 16 updates apart | +4 | R2-ENGINE-1536 |
| 17 | Wall of Earth | 10 or 9 by facing | 10 + ceil(floor(16P/10)/16), 16 updates apart | applies nothing | R2-ENGINE-1536 |

P is the book power 0 to 255, so the pulse count runs from the first number to that number plus 26. A unit
of s by s cells stands at entry +4 of each covered cell: a wall or Poison Cloud applies once per covered shape
cell per pulse, Fire Ball once per covered cell of its square at damage divided by the square of the unit's size byte (equal to s if that byte equals s; its source is not read)
(R2-ENGINE-1536, R2-ENGINE-1537, Medium). Blizzard cell validity EN `L2.02717` / RU `L2.01700` (61 of 61 masked-equal)
admits a candidate `(x, y)` iff `8 <= x <= W - 9` and `8 <= y <= H - 9`, reading only ALM type-0 `W` and `H`
(`+0x50000`, `+0x50004`) and clamping out-of-band coordinates (High in both roots, R2-ENGINE-1668, amending R2-ENGINE-1533).

Light and Darkness reach only entry +4: kind 1 or 2 units (Units rows with movementType 1 or 2, which include Ghost and Bee; R2-ASSET-327, R2-ASSET-328) are reached; a kind 3 unit (Bat_Sonic, Dragon, Succubus; a flier by name only) at entry +8 and
a building at entry +0xc are not (R2-ENGINE-1535, R2-ENGINE-1530, Medium). On a reached unit, in EN a pulse
attaches an Effect of duration 16, the actor step decrements it once per update and runs the end at the 16th
step of the pulse, which adds the sight shift back and unlinks the Effect; the next pulse, 16 updates after
the first, attaches again (a unit on several cells of one pulse renews one Effect, R2-ENGINE-1236). Sight is base between the end and that pulse once per 16 updates, and 15 updates
after the last pulse it returns for good (R2-ENGINE-1534, R2-ENGINE-1165, Medium; the RU twins are unread; no
census of other callers of the timer slot).
The per-class amount stays as in R2-ENGINE-1242.

The Poison Cloud Effect (ID 6): the RU arm stores no +0x42 in the 256 instructions read, and the suffix stores none
(R2-ENGINE-1262; within one step the executor runs before that arm, R2-ENGINE-1763); the value 128 from the Effects string through the constructor is Medium (R2-ENGINE-1575; EN arm
unread). A cast onto a unit that carries the Effect keeps the record, leaves +0x42 at the maximum of the old value and
128, never adds, keeps the less negative magnitude and does not call start (R2-ENGINE-1576, Medium; RU). From the Shield
arm of the support choice to the end of the recovery draw, a Shield cast at the caster draws once on the read path in
RU, and a cast through the single-target object skips the hit hook when that object carries the cast
(R2-ENGINE-1577, R2-ENGINE-1578, Medium; the EN twins of `L2.00031`, `L2.02072`, `L2.08110` and `L2.02074` are unread).

## Script-set effect times and script casts

Script action 29 sets the remaining time +0x50 of the area effects of spell p2 at cell (p0, p1) to p3, and action 30
sets the duration +0x42 of unit U's effects of spell p0 to p1 (Medium; R2-ENGINE-1749). Actions 21 and 24 cast one
spell through a caster Unit with Mind 30 whose sphere skill holds the node's power word, so the spell power is the word
up to 255 (Medium); the caster casts once and is deleted (Medium; R2-ENGINE-1740, R2-ENGINE-1741, R2-ENGINE-1742). The shipped
nodes cast Poison Cloud, Blizzard, Wall of Earth and Stone Curse (R2-ASSET-379).

A cell event of kind 23, written by a type-9 record or script action 25 (R2-ENGINE-1775), casts no spell: placement
skips its cast branch for that kind, and a unit whose step ends at the cell centre is passed to the cell move toward the target cell with no
effect created in the bodies read (Medium for the absence of an effect and for units outside the three step callers;
R2-ENGINE-1778). No store to the event was found in the EN range of R2-ENGINE-1777 (Medium; whole-entry stores and RU
bodies other than the empty test not searched). Mission 53 writes one from a trigger, and both roots ship nine action 25 nodes with kind 23
(R2-ASSET-386, R2-ASSET-387).

## Acid Stream by heading and at the map edge

The RU heading d is bits 5..7 of the byte `L2.02635` returns: an octant from the offset between caster and
target by slope thresholds 1/2 and 2, equal sides and a zero offset going to a diagonal (a zero offset to heading 7)
(R2-ENGINE-1545, High for the RU instructions read except the one at `L2.08087`, whose continuation to the tail for
sectors 4 to 7 is Medium; Medium for the position getters and the EN twin). Cells of pulse p around
the origin (x, y), p = 0..4 for the straight headings (pulse 5 is empty) and p = 0..5 for the diagonal ones:

| d | Cells of pulse p | Claim |
|---|---|---|
| 0 | (x+i, y-p), i = -p..p | R2-ENGINE-1546 |
| 1 | (x+(p-i), y-i), i = 0..p | R2-ENGINE-1546 |
| 2 | (x+p, y+i), i = -p..p | R2-ENGINE-1546 |
| 3 | (x+(p-i), y+i), i = 0..p | R2-ENGINE-1546 |
| 4 | (x+i, y+p), i = -p..p | R2-ENGINE-1546 |
| 5 | (x-(p-i), y+i), i = 0..p | R2-ENGINE-1546 |
| 6 | (x-p, y+i), i = -p..p | R2-ENGINE-1546 |
| 7 | (x-(p-i), y-i), i = 0..p | R2-ENGINE-1546 |

A cell whose byte x or y is below 8 or above the map's size field minus 9 is skipped whole: no packet, no occupant
call, no damage; the other cells are unchanged, and no cell is clamped or shifted (R2-ENGINE-1546, High for the RU
body, Medium that the two fields are width and height and for EN, whose cone also differs at pulse 2, R2-ENGINE-797).
The origin (x, y) is the stage's origin byte pair; whether it is the target cell or the caster's cell plus an offset of
R2-ENGINE-810 is not read here (Unknown), so the table does not give which map cells are hit. The direction tables of
R2-ENGINE-810 are read again in the terrain constructor, where the unit step vector uses them (R2-ENGINE-1470); a unit step changes its cell on advance ceil(128/v) or floor(128/-v)+1 of the step, by its y vector component (x when it has none) (R2-ENGINE-1760).

## Stone Curse reach by object class

The point apply attaches the Effect to a target whose class slot `+0x2c` is true (Unit, Humanoid, Human) and also to
one whose slot `+0x34` is true: Building and Pointer reach the attach, Token is refused. The slot `+0x2c` bodies of the Unit and Human vtables are `mov eax, 1; ret` (R2-ENGINE-1565, High for the
slot bodies). A ROM2 mage restores max*(M+100)*K/D hundredths of mana per 16-tick pass, M the `+0xe2` value and D the
Units MP regeneration (R2-ENGINE-1566, High for the formula, Medium for the source of M and the Human D join, Unknown for D 0).
A Building's actor slots do
not walk its Effects, so no timer end is read for it; Token slot `+0x34` returns 0 here and slot `+0x2c` by R2-ENGINE-811,
and the class sizes are those of R2-ENGINE-517 (slot `+0x38` of these classes is read in RU, and in EN for Building, Pointer and Outpost, which returns 1 as in RU: R2-ENGINE-1701 (amended), R2-ENGINE-1820; action 31 writes p1 into +0xb4 of every Outpost whose group number equals p0, for every owner and with no loop exit (the filter passes only the Outpost in the classes with a record; EN Shop and Inn slot 14 unread, Medium), counted in container passes: R2-ENGINE-1820, R2-ENGINE-1821; R2-ENGINE-1550, High for the slots, Medium for the composition
and for whether a cast request can select such a target). The hold is the actor mask of R2-ENGINE-454 and R2-ENGINE-786.
State 16, which makes the actor step skip its Effect walk (R2-ENGINE-786), is stored by two immediate `mov` stores per
root with displacement 0x54, both on the death path; a unit removed from play is not given it among those stores (RU
linear decode of `.text`; EN immediate-0x10 form only). Register stores (29 in RU, 27 unread), other store forms and the
EN register stores are not read (R2-ENGINE-1547, High for the sites, Medium that no other writer exists).

## Strike refusal

The RU unit strike refuses a blow from a non-mage Human-kind attacker of attack type other than 5 against a movementType 3 target, and from any attacker for a missing party or owner, health at or below 0 or a range below the distance (R2-ENGINE-1605, Medium that a Human's attack type byte is the weapon's attackType; the EN strike body has no such test, EN behaviour Medium).

## Idle-mage casting

A player mage on order 0x17 casts support spells in the part of the RU body read (R2-ENGINE-1627); the per-cell application of an area effect is R2-ENGINE-1537.

## Pocket flight, Teleport flight, staged areas, fallen units, hostile Heal, support relations, hired thresholds

In RU a refused flight in a pocket closed by terrain walks toward a cell at most 3 cells away for a unit inside the 8 .. size - 9 band; the route is not emptied when the route search's stand-in is another pocket cell, and then the order-1 branch is not entered and each Group pass stores the walk order again (R2-ENGINE-1787, amending R2-ENGINE-1498, R2-ENGINE-879 and R2-ENGINE-871; Medium); the own-cell stand-in case is open, and the outcome in play is Unknown. A computer mage that can pay Teleport casts it at a cell scanRange + 1 cells from its own cell, away from the hostiles' mean point along the major axis, clamped to the map band; the move happens when the landing test accepts the cell (Medium for the arguments; High for the arithmetic of `L2.02454`); with the hostile outside the flee list the idle planner gives prepared 6 on a hostile in range, else 0xb and, for a damaged mage with Heal, a self-Heal, and no walk (R2-ENGINE-1788, amending R2-ENGINE-1418 and R2-ENGINE-506; the caster's cell 40 ticks after the cast is Unknown).

Acid Stream and Blizzard accepted at one wind-up both attempt on area update 1; a first cell outside the band or rejected moves the first paint to update 4 or later (R2-ENGINE-1789, amending R2-ENGINE-812; Medium).

A fallen unit keeps running its Effect timers in each actor step before the health test; the fall sets the counter +0x6c to virtual +0x74 - 1 (the row value at index 33; for a Human the value at index 23 or the constant 8; counted in actor steps), and the death loop ends Poison with no further start (R2-ENGINE-1790, Medium; amending R2-ENGINE-1576 and R2-ENGINE-1534). From -3, a Poison Cloud of power P takes m = 2 + floor(2P/45) on each eighth timer call, and the unit dies on the first step after the counter at which health is -10 or below; the refill's loss below 0 comes on top (R2-ENGINE-1791, Medium):

| m (power P) | Start that first brings health from -3 to -10 or below, for word `+0xc6` = 0, ignoring the refill's loss (Medium) | Claim |
|---|---|---|
| 2 (P 0..22) | 4th | R2-ENGINE-1791 |
| 3 (P 23..44) | 3rd | R2-ENGINE-1791 |
| 4..6 (P 45..112) | 2nd | R2-ENGINE-1791 |
| 7 and above (P 113 and above) | 1st | R2-ENGINE-1791 |

The RU Heal arm refuses a target whose Player the caster's Player row marks hostile (bit 0), before its health tests; a refused Heal runs none of the arm's callees, while the apply head's cast award has run under its gates (bit 2, no item, mage, row parameter 2 nonzero) and the mana was paid at the cast (R2-ENGINE-1792, Medium for the argument roles; it amends R2-ENGINE-487 and the partially retracted R2-ENGINE-026).

The player support sweep heals only members of the builder's list +0xaa8 (the caster's bit 0 toward them clear) that pass the affinity; with Neutral alone the diagonal 0x12 refuses own units other than the caster, and a Player at relation 3 is not in the list the Heal walk reads (builder `L2.01515`) (R2-ENGINE-1793, amending R2-ENGINE-360 and R2-ENGINE-1200; Medium in play: the diagonal can be rewritten by script action 10, and `L2.01514` and `L2.01520` are unread). The relation matrix holds one byte per ordered pair of Player numbers at Session + 0xa8c4 + 70a + b; bit n is value 2^n:

| Bit or value | Use | Claim |
|---|---|---|
| bit 0 | the row's Player holds the column's hostile: hostility predicate, support builder, Heal arm | R2-ENGINE-1794 |
| bit 1 | ally: support filter, affinity Allies test | R2-ENGINE-1794 |
| low two bits 0 | affinity Neutral test | R2-ENGINE-1794 |
| bit 4 (0x10) | copied in reverse into the client word by the EN server routine `L2.01852` (EN read, RU not read) | R2-ENGINE-1794, R2-ENGINE-826 |
| 0x12 on the diagonal | map build and the original's saves | R2-ENGINE-1794 |
| (old & 0xfc) + low8(p2) | script action 10, which can store 3 | R2-ENGINE-1794 |

For a Human built by `L2.00130`, a hire or a map Human, no direct call reaches the loader or the threshold setter: the Unit constructor's default name has length 0 and the setter would read the Unit slot +0x30; command +0x40, +0x44 and +0x4c are 0 after creation, with a book or a bow (R2-ENGINE-1795, amending R2-ENGINE-1606, R2-ENGINE-506, R2-ENGINE-1001, R2-ENGINE-1002 and R2-ENGINE-777; Medium).

## Cell planes in the route search

The route search (RU `L2.05174`, EN `L2.05175`) floods its word grid at +0x30000 + 2 × key from the unit's cell. With flag 1 (walk and approach) it tests plane +0x10000, with flag 0 (repath) plane +0x20000, and it skips a neighbour whose plane byte AND the unit's mask byte is nonzero; for unit kind 1 a step costs the terrain byte straight and one and a half times it diagonal, for other kinds 2 and 3 (R2-ENGINE-1804; the flag roles are Medium, R2-ENGINE-871). In the size-1 arm the loop tests the goal word, the frontier and the half-pass limit once per pair of one list-A and one list-B half-pass (R2-ENGINE-1876). The grid helpers are read: `L2.05185`(x, y) and `L2.05180`(key) address the same grid word for one `this`, `L2.05230` the terrain byte and `L2.02402` returns unit virtual +0x20 (R2-ENGINE-1804, amending R2-ENGINE-871, R2-ENGINE-878 and R2-ENGINE-879). With the plane values the terrain build stores (R2-ENGINE-1340), for flag 1 and a unit of size 1, value 1 closes kind 1, value 5 kinds 1 and 2, and the border ring's 0x1f all three kinds; in the route-search workers and helpers read, a terrain byte of 255 enters only kind 1's step sum and no test closes a cell on it (R2-ENGINE-1805, Medium; the outcome in play is Unknown). For a larger unit the footprint loop is RU `L2.08532` (EN `L2.08533`), which calls the AND of two bytes RU `L2.03520` (EN `L2.02720`) (R2-ENGINE-1804, narrowing the footprint tester naming of R2-ENGINE-1667). The EN location routine `L2.07308` keeps `(y << 8) + x`, which has the form of the plane key (R2-ENGINE-1802; EN body read, Medium that the word is the plane key, RU body unread).

The RU reads of the route search loop, the group target scan, the defender list, the hostile filter and the Teleport admission, and the older cards they amend, are cited here: R2-ENGINE-502, R2-ENGINE-1418, R2-ENGINE-1440, R2-ENGINE-1804, R2-ENGINE-1879, R2-ENGINE-1880.
