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
| 7 | Blizzard | 1 of 36 | 32, 3 updates apart | +4, +8, +0xc | R2-ENGINE-1533 |
| 9 | Acid Stream | cone, 25 or 21 in total | 6, at updates 1, 4, 7, 10, 13, 16 | +4, +8, +0xc | R2-ENGINE-1536 |
| 14 | Darkness | diamond, 57 | 15 + ceil(floor(16P/10)/16), 16 updates apart | +4 | R2-ENGINE-1536 |
| 15 | Light | diamond, 57 | 15 + ceil(floor(16P/10)/16), 16 updates apart | +4 | R2-ENGINE-1536 |
| 17 | Wall of Earth | 10 or 9 by facing | 10 + ceil(floor(16P/10)/16), 16 updates apart | applies nothing | R2-ENGINE-1536 |

P is the book power 0 to 255, so the pulse count runs from the first number to that number plus 26. A unit
of s by s cells stands at entry +4 of each covered cell: a wall or Poison Cloud applies once per covered shape
cell per pulse, Fire Ball once per covered cell of its square at damage divided by the square of the unit's size byte (equal to s if that byte equals s; its source is not read)
(R2-ENGINE-1536, R2-ENGINE-1537, Medium).

Light and Darkness reach only entry +4: kind 1 or 2 units (Units rows with movementType 1 or 2, which include Ghost and Bee; R2-ASSET-327, R2-ASSET-328) are reached; a kind 3 unit (Bat_Sonic, Dragon, Succubus; a flier by name only) at entry +8 and
a building at entry +0xc are not (R2-ENGINE-1535, R2-ENGINE-1530, Medium). On a reached unit, in EN a pulse
attaches an Effect of duration 16, the actor step decrements it once per update and runs the end at the 16th
step of the pulse, which adds the sight shift back and unlinks the Effect; the next pulse, 16 updates after
the first, attaches again (a unit on several cells of one pulse renews one Effect, R2-ENGINE-1236). Sight is base between the end and that pulse once per 16 updates, and 15 updates
after the last pulse it returns for good (R2-ENGINE-1534, R2-ENGINE-1165, Medium; the RU twins are unread; no
census of other callers of the timer slot).
The per-class amount stays as in R2-ENGINE-1242.

The Poison Cloud Effect (ID 6): the RU arm stores no +0x42 in the 256 instructions read, and the suffix stores none
(R2-ENGINE-1262); the value 128 from the Effects string through the constructor is Medium (R2-ENGINE-1575; EN arm
unread). A cast onto a unit that carries the Effect keeps the record, leaves +0x42 at the maximum of the old value and
128, never adds, keeps the less negative magnitude and does not call start (R2-ENGINE-1576, Medium; RU). From the Shield
arm of the support choice to the end of the recovery draw, a Shield cast at the caster draws once on the read path in
RU, and a cast through the single-target object skips the hit hook when that object carries the cast
(R2-ENGINE-1577, R2-ENGINE-1578, Medium; the EN twins of `L2.00031`, `L2.02072`, `L2.08110` and `L2.02074` are unread).

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
R2-ENGINE-810 is not read here (Unknown), so the table does not give which map cells are hit.

## Stone Curse reach by object class

The point apply attaches the Effect to a target whose class slot `+0x2c` is true (Unit, Humanoid, Human) and also to
one whose slot `+0x34` is true: Building and Pointer reach the attach, Token is refused. The slot `+0x2c` bodies of the Unit and Human vtables are `mov eax, 1; ret` (R2-ENGINE-1565, High for the
slot bodies). A ROM2 mage restores max*(M+100)*K/D hundredths of mana per 16-tick pass, M the `+0xe2` value and D the
Units MP regeneration (R2-ENGINE-1566, High for the formula, Medium for the source of M and the Human D join, Unknown for D 0).
A Building's actor slots do
not walk its Effects, so no timer end is read for it; Token slot `+0x34` returns 0 here and slot `+0x2c` by R2-ENGINE-811,
and the class sizes are those of R2-ENGINE-517 (R2-ENGINE-1550, High for the slots, Medium for the composition
and for whether a cast request can select such a target). The hold is the actor mask of R2-ENGINE-454 and R2-ENGINE-786.
State 16, which makes the actor step skip its Effect walk (R2-ENGINE-786), is stored by two immediate `mov` stores per
root with displacement 0x54, both on the death path; a unit removed from play is not given it among those stores (RU
linear decode of `.text`; EN immediate-0x10 form only). Register stores (29 in RU, 27 unread), other store forms and the
EN register stores are not read (R2-ENGINE-1547, High for the sites, Medium that no other writer exists).

## Strike refusal

The RU unit strike refuses a blow from a non-mage Human-kind attacker of attack type other than 5 against a movementType 3 target, and from any attacker for a missing party or owner, health at or below 0 or a range below the distance (R2-ENGINE-1605, Medium that a Human's attack type byte is the weapon's attackType; the EN strike body has no such test, EN behaviour Medium).

## Idle-mage casting

A player mage on order 0x17 casts support spells in the part of the RU body read (R2-ENGINE-1627); the per-cell application of an area effect is R2-ENGINE-1537.
