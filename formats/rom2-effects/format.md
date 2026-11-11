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
