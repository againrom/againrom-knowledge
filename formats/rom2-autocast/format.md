# ROM2 automatic casting

The selected RU human support routine reads the owning Player's preference
byte. Three cumulative masks select Minimum, Average and Maximum; the
command's selected offensive spell is separate. Local order routes, attack
range and primary dispatch have static evidence. Update cadence, physical
cycle modifier, final lifecycle values and whole-policy EN/RU parity remain
Unknown.
— R2-ENGINE-345, R2-ENGINE-350, R2-ENGINE-361, R2-ENGINE-366,
R2-SESSION-152, R2-SESSION-153, R2-ENGINE-396, R2-ENGINE-397,
R2-ENGINE-400, R2-ENGINE-401, R2-SESSION-163, R2-SESSION-164

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
modifier/state and complete writer population are not joined. The local cycle
preserves target bits 0..2 and changes 8 -> 24 -> 56 -> 8. This is a bounded
numeric input route, with Medium physical-control confidence. The separate
object+0x54 storage clause of R2-ENGINE-359 is partially retracted.
— R2-ENGINE-345, R2-ENGINE-359, R2-ENGINE-366, R2-ENGINE-400

The acquired EN/RU key bodies require application DWORD+0x404 (EN) or
+0x41c (RU) to equal 1 and view+0x80 to be nonnull. Nonzero application+0xc0
forwards input through another object's virtual+0x6c before the key arms.
Paired options fragments form the same level masks 8/24/56 and add target
choices 1/2/4. These local tests and writes are High; complete physical input
and transport composition remain Medium. Direct flag and reduced adjacent-base
scans do not establish a complete indexed or overlapping writer census.
— R2-ENGINE-400

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
preferences constructor instead sets byte+0x20 to 0xff. Paired EN/RU
serializers request a 36-byte raw block in both directions. Complete successful
preference-block LOAD restores the saved byte; later refresh and the final
cold-LOAD byte remain Unknown. Aligned EN/RU setup excerpts both resend the
application policy, but complete delivery and transition precedence are
unjoined. New-game, mission-entry, town and postmission winning values remain
Unknown. No complete town/departure parent chain was acquired.
— R2-ENGINE-345, R2-SESSION-150, R2-SESSION-152, R2-SESSION-153,
R2-ENGINE-401, R2-SESSION-163, R2-SESSION-164

## Casters and support attempts

The acquired RU idle and attack callers test mage flag actor+0x4c &4 and
owner Player kind+0x2c == 0 before human support. The attack route can attempt
support after ordinary attack preparation when command range+0x14 is below 2.
Nonzero owner kind reaches separate computer policy. These routes do not
establish admission for every class, hire, summon, converted unit or current
order. Support in all moving, holding and retreating states and attempts per
simulation tick remain Unknown. — R2-ENGINE-351, R2-ENGINE-362,
R2-ENGINE-396, R2-ENGINE-398

The attack helper chooses the selected spell's range when the spell exists
and current mana covers its cost; otherwise it uses actor attack range+0x12c.
The attack planner writes this range to command+0x14 before its conditional
support call. The unsigned byte below 2 is a range gate. Support execution
relative to the actor's strike in one parent update remains Unknown.
— R2-ENGINE-397

## Scheduling, orders and actor states

The acquired RU scheduler traverses the reached actor collection and calls
its order selector. The qualifying idle human mage path has one support call
per wrapper invocation. The scheduler's parent and interval are unjoined.
The wrapper and support body do not test primary+9; primary dispatch still
takes priority over prepared+8. Local calls do not establish attempts per
tick or a universal busy/cooldown contract. — R2-ENGINE-396

The acquired numeric order table routes order 2 to coordinate movement,
3 to attack, 12 to idle, 13 to targeted cast and 14 to another cast helper.
Order 27 sets actor state 27 without a planner; default routes through the
idle wrapper. The order-2 branch prepares movement while destination differs
and an arrival order when equal. It does not directly call human support.
Named hold and retreat input bindings and other support events while visibly
walking remain Unknown. — R2-ENGINE-398

Actor state 16 exits the acquired actor step. Its living branch requires
positive health; actor+0x3c = 0 makes the acquired command-skip predicate true.
The dispatcher tests effect mask 0x40000, then installs primary 4 only when
primary+9 is zero. Its primary-4 branch sets actor state 27 and tests the mask
again, retaining the command while present and clearing it when absent.
Named stun mapping, Slow timing and invisible-caster admission remain Unknown.
The proposed Invisibility effect's absence gate is not a universal caster
suppression rule. — R2-ENGINE-399

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
helpers and successful movement remain unread. The acquired primary branch
increments phase and calls an unread cast-delivery helper at phase 3.
Nonzero actor completion byte+0x136 clears primary+9. Complete effect
application and parent timing remain Unknown. — R2-ENGINE-363, R2-ENGINE-396

Several sampled setters clear pending byte+8 and the support flag; the two
order-1 setters store byte+8 = 1. Interruption after primary cast
initialization is read in R2-ENGINE-545. Whether the former order resumes and
how often support tries remain Unknown. — R2-ENGINE-362, R2-ENGINE-363,
R2-ENGINE-541, R2-ENGINE-545

## Selected offense, computer policy and locales

RU selected-spell policy reads command byte+0x78 and dispatches IDs 1..28;
it does not iterate the book to choose a replacement spell in that body.
The selected owner-kind branches use different computer offensive and
support routines. Random distribution and admission of every computer unit
remain Unknown. — R2-ENGINE-350, R2-ENGINE-351

EN/RU acquired third maps, cycles, receiving byte widths, options masks and
36-byte preference widths agree. Their key-mode field offsets differ:
EN+0x404 and RU+0x41c; both test literal 1. Both aligned setup excerpts resend
application policy. The earlier paired constructors and registry transfers
remain bounded local evidence. Full support, order/status admission, cadence,
SAV precedence, feedback and lifecycle equivalence was not measured.
— R2-ENGINE-345, R2-SESSION-150, R2-ENGINE-366, R2-ASSET-085,
R2-ENGINE-400, R2-ENGINE-401, R2-SESSION-164

## Blizzard candidate boundary

The inspected RU offensive candidate loop tests IDs 1..28, a nonnull spell,
sufficient mana and cached spell flag +0x0a=0. Row 7 has no explicit exclusion.
The stored Defensive value -1 is not bound to that cached flag in the read
population. Actual Blizzard admission, selected choice, caller reachability
and EN choice parity remain Unknown. The local conditional test is Medium;
no original cast was observed (R2-SESSION-162).

## Player order machine

A player order reaches the unit through a setter that resets the unit state,
replaces the order code and parameters at once and stores prepared byte +8: 0,
except in the two order-1 setters, which store 1; no setter writes primary
byte +9 (R2-ENGINE-541). The executor reads +9 first (1 strike, 2 cast, 3 step,
4 hold; the 239 byte-table cells for other values are unread) and prepared
byte +8 only when +9 is zero (R2-ENGINE-542). A cast passes init, the
cast-routine try, wind-up and recovery (R2-ENGINE-543); the mage's mana is
subtracted at the cast routine's entry, once per call that reaches the delivery
(R2-ENGINE-544). A new order during a cast or a step between cells is stored
and, in the setter windows read, acts after the running action ends; order 1
acts at the next executor call, the other kinds after a planner pass
(R2-ENGINE-545, R2-ENGINE-546). Most planners, the command word +0xb0 and the
strike and hold cells are unread. States, EN/RU equality and Unknowns:
R2-ENGINE-547, R2-ENGINE-548, R2-ENGINE-549.
