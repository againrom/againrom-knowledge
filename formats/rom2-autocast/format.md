# ROM2 automatic casting

The selected RU human support routine reads the owning Player's preference
byte. Three cumulative masks select Minimum, Average and Maximum; the
command's selected offensive spell is separate. Control, receiving byte
stores and candidate/order routes have bounded static evidence. Final
lifecycle values and whole-policy EN/RU parity remain Unknown.
— R2-ENGINE-345, R2-ENGINE-350, R2-ENGINE-361, R2-ENGINE-366,
R2-SESSION-152, R2-SESSION-153

## Controls and feedback

| Level | Level mask | EN label | RU source view |
|---|---|---|---|
| Minimum | 0x08 | Minimum | Минимальный |
| Average | 0x18 | Average | Средний |
| Maximum | 0x38 | Maximum | Максимальный |

The selected RU options combo 0x8d uses dialogs source indices 171..173 and
applies selected indices 0/1/2 as these masks. Controls 0x83/0x84/0x85 use
indices 167/168/169, named Own/Allies/Neutral in EN and Своих персонажей /
Союзников / Нейтральных in a CP1251 RU view. Apply independently adds target
masks 1/2/4 to the selected level mask. The selected loader binds the dialogs
receiver to main/text/dialogs.txt; complete native encoding and rendered RU
glyphs remain Unknown. — R2-ENGINE-345, R2-ENGINE-360, R2-ASSET-085,
R2-ASSET-088

The selected EN/RU third key maps route input 0x55 (U) to the cycle arm.
That arm cycles only when an unnamed global flag is nonzero. Its physical
modifier/state and earlier routing are not fully joined. The local cycle
preserves target bits 0..2 and changes 8 -> 24 -> 56 -> 8. This is a bounded
numeric input route, with Medium physical-control confidence. The separate
object+0x54 storage clause of R2-ENGINE-359 is partially retracted.
— R2-ENGINE-345, R2-ENGINE-359, R2-ENGINE-366

After cycling, the selected RU arm requests patch-pool indices 115..117
and passes the text with argument 2000 to the view's text-state helper.
The receiver loads patch/patch.txt, separate from main.txt. The text helper
loads its third argument and forwards it with receiver+0x2c to L2.01545,
outside the acquired ranges. The final field holding the value, its units,
actual feedback strings, persistent panel state and sound remain Unknown;
main.txt autohealing labels do not supply this binding. The object+0x54
storage clause of R2-ENGINE-359 is partially retracted.
— R2-ENGINE-359, R2-ASSET-088

## Owner scope and lifetime sources

The sender carries a Player identity and application preference DWORD.
Compatible local EN/RU receive arms use packet word +5 to look up a Player,
then copy packet byte +0xe to that Player's preferences+0x20. No actor or
Group identity occurs in the sender payload. This supports owner scope at
Medium; complete subtype dispatch, selection/Group interaction and ownership
changes remain Unknown. — R2-ENGINE-345, R2-ENGINE-361

Application constructors set DWORD+0x24 to 0x19 (Average plus own target),
and registry AutoCasting transfers four bytes there. The acquired RU Player
preferences constructor instead sets byte+0x20 to 0xff. Player serialization
includes that byte in a 36-byte raw block. A selected RU mission-setup excerpt
resends the application DWORD. These are distinct sources; none alone
establishes the final new-game or cold-LOAD value. Mission/town preservation
and later overrides remain Unknown. — R2-ENGINE-345, R2-SESSION-150,
R2-SESSION-152, R2-SESSION-153

## Casters and support attempts

The sampled RU idle and attack callers test mage flag actor+0x4c &4 and
owner Player kind+0x2c == 0 before human support. The attack route can attempt
support after ordinary attack preparation when command range+0x14 is below 2.
Nonzero owner kind reaches separate computer policy. These routes do not
establish admission for every class, hire, summon, converted unit or current
order. Support in all moving, holding and retreating states and attempts per
simulation tick remain Unknown. — R2-ENGINE-351, R2-ENGINE-362

RU support stops after the first prepared order. Every attempt needs the
spell in actor book+0x140 and current mana at least its spell cost. Each stage
has an independent threshold on the same current mana; thresholds are not
additional deductions. — R2-ENGINE-346, R2-ENGINE-347, R2-ENGINE-348

| Stage | Additional gate | Ordered attempts |
|---|---|---|
| Minimum | bit 0x08; mana >=5 | self Heal 24 if HP differs from maximum; self Shield 27 if absent; listed Heal target |
| Average | bit 0x10; mana >=10 | for each candidate, Bless 20 then Haste 21 |
| Maximum | bit 0x20; mana >=40 | for each candidate, Air protection 13, Earth 19, Fire 4, Water 8, Invisibility 12; fallback self Invisibility 12 |

Candidate Heal requires list count greater than one. It retains the first
strict minimum of signed integer HP/maximumHP and requires that quotient
below 1. Equal scores preserve the first candidate. Average/Maximum are
target-major. Effect absence and spell range also constrain each attempt;
Bless has an additional unread predicate and candidate Invisibility excludes
self with its recorded relation condition. — R2-ENGINE-346, R2-ENGINE-347

## Candidates, relation and range

The selected support body feeds caster+0x70 to a spatial candidate builder,
then filters that result with a second builder. The population at +0x70 is
not proved to be the UI selection or Group. The spatial marker's callees
prevent a complete sight, occlusion or physical-radius contract.
— R2-ENGINE-364

The second builder excludes candidate state 0x10, separates relation bit 1
into a different list, and conditionally filters effect 0x1000. With that
effect, relation byte exactly 2 passes; otherwise a member of caster+0x70
must pass an unread helper comparison against its command byte+0x71.
Support consumes the relation-bit-1-clear list. These are local predicates,
with Medium complete membership confidence. — R2-ENGINE-364

The selected second-stage list appends at its tail and iterates next nodes,
so it preserves that builder's input order among admitted support candidates.
Heal ties and target-major buffs inherit this order. Upstream creation,
selection, distance, ID and LOAD priority remain Unknown. — R2-ENGINE-365

Later affinity accepts self unconditionally. Other observed branches accept
same owner with preference bit 1, relation bit 2 with preference bit 2, or
relation&3 == 0 with preference bit 4. The named self attempts bypass these
target toggles. Other consumers and the full relation grammar remain Unknown.
— R2-ENGINE-348, R2-ENGINE-360

Range uses the word returned by the position getter and compares a packed-byte
Chebyshev metric to spell byte+9. The getter reads its object's word+2; the
metric returns the maximum absolute difference between the unsigned low-byte
pairs. Physical coordinate units and complete sight semantics remain Unknown.
Self Heal, self Shield and fallback self Invisibility bypass this wrapper.
— R2-ENGINE-348, R2-ENGINE-364

## Prepared support and player orders

Support marks command+0x60 = 1 and prepares byte+8 = 8, target+0x28,
spell+0x30 and spell range+0x14. It records target pending spell ID and tick.
Successful preparation is not an applied effect. Buff attempts suppress
literal pending ID 20 within absolute tick difference 20; this is not a
universal per-spell cooldown. — R2-ENGINE-349, R2-ENGINE-363

The dispatcher handles a nonzero primary byte+9 before a prepared byte+8.
Prepared byte 8 indexes slot 7 and reaches RU L2.01552. The local self/pass
branch initializes primary cast byte 2, phase 0, actor state 13, target and
spell. The other branch prepares an approach and actor state 1. The approach
helpers and successful movement remain unread, and the cast-completion route
is not established. — R2-ENGINE-363

Several sampled setters clear pending byte+8 and the support flag. Their
player-input callers and interruption after primary cast initialization are
not joined. Whether the former order resumes and how often support tries
remain Unknown. — R2-ENGINE-362, R2-ENGINE-363

## Selected offense, computer policy and locales

RU selected-spell policy reads command byte+0x78 and dispatches IDs 1..28;
it does not iterate the book to choose a replacement spell in that body.
The selected owner-kind branches use different computer offensive and
support routines. Random distribution and admission of every computer unit
remain Unknown. — R2-ENGINE-350, R2-ENGINE-351

EN/RU acquired third maps, cycles, receiving byte widths and 36-byte preference
widths agree. The earlier paired constructors and registry transfers remain
bounded local evidence. Full options, support, target, command execution,
SAV precedence, feedback and lifecycle equivalence was not measured.
— R2-ENGINE-345, R2-SESSION-150, R2-ENGINE-366, R2-ASSET-085

## Blizzard candidate boundary

The inspected RU offensive candidate loop tests IDs 1..28, a nonnull spell,
sufficient mana and cached spell flag +0x0a=0. Row 7 has no explicit exclusion.
The stored Defensive value -1 is not bound to that cached flag in the read
population. Actual Blizzard admission, selected choice, caller reachability
and EN choice parity remain Unknown. The local conditional test is Medium;
no original cast was observed (R2-SESSION-162).
