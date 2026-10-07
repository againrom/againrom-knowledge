# SAV Human live fields and modifiers

[Format reference](format.md) · [Actor wire programme](actors.md)

## Live and modifier fields

The raw-a6 tail is `+0xb4..+0xbd` (10 bytes), raw-be tail is
`+0xc0..+0xd3` (20), and raw-d4 is `+0xd4..+0x113` (64). These 94 bytes
carry live state and modifiers. — SAV-HUMRUN-444

| live fields | modifier fields | role |
|---|---|---|
| `+0xb4/+0xb5` | `+0xf4/+0xf5` | physical base/spread bytes |
| `+0xb6` | `+0xf6` is not folded | active skill index, written directly by weapon paths |
| `+0xb7/+0xb8` | `+0xf7/+0xf8` | second damage base/spread bytes |
| `+0xb9/+0xba/+0xbb` | `+0xf9/+0xfa/+0xfb` | elemental base/spread and assigned kind |
| `+0xbc/+0xbd` | `+0xfc/+0xfd` | unnamed final bytes, not used by the named folds |
| `+0xbe/+0xc0` | `+0xfe/+0x100` | defence/absorption words |
| `+0xc2..+0xcc` | `+0x102..+0x10c` | six protection words; slot-zero gameplay meaning remains Unknown |
| `+0xce..+0xd3` | `+0x10e..+0x113` | six damage-kind resistance bytes |

## Modifier layout

The modifier prefix is four signed stat-cap bytes `+0xd4..+0xd7`, then
seven words: speed `+0xd8`, capacity `+0xda`, maximum health `+0xdc`, health
regeneration percentage `+0xde`, maximum mana `+0xe0`, mana regeneration
percentage `+0xe2`, and sight `+0xe4`. The attack modifier begins with to-hit
`+0xe6`, General shadow `+0xe8` and five class-skill modifiers `+0xea..+0xf2`.
The named derive/award loops skip the General shadow. — SAV-HUMRUN-444

## Equipment folds

The attack fold adds one word and six bytes, assigns elemental kind,
and skips skills, the active-index mirror and the final two bytes. The
defence fold adds eight words and six bytes. Additions wrap at their storage
widths before any later derive clamp. — SAV-HUMFOLD-446

Weapon kinds 11/12 assign the to-hit modifier from General; removal subtracts
the weapon's own to-hit word. This is not an inverse for arbitrary prior
modifier state. Armor/Shield explicitly add their blocks into both defensive
copies; Weapon invokes derive after direct modifier writes. Reconstructing
all modifiers as a sum of currently equipped objects is therefore not an
established authoring law. — SAV-HUMEQUIP-447

## Current amounts and regeneration

Current health/mana, regeneration remainders and earned XP are distinct from
these modifier words. Admitted health/mana arithmetic changes current
amounts and remainders while leaving the selected 94 bytes unchanged; a raised
active skill can change live damage through derive without changing its
modifier. First post-load admission remains Unknown. — SAV-HUMMUT-448

Regeneration consumes current/max/period at `94/96/98` and `9a/9c/9e`, plus
modifiers `de/e2`, as signed 16-bit; the stored remainder bytes `a2/a3` reload
unsigned. Modifier plus 100 is 32-bit. Products and the accumulator wrap at 32
bits, followed by signed division truncated toward zero. The exact admitted
formula and local gates are in the [hero regeneration reference](../hero/format.md).
— SAV-REGENWIDTH-528, HERO-REGEN-021 (amended)

Each reached arm stores the remainder's low byte, then the quotient's low
word, then the signed minimum of that stored word and the maximum. There is
no lower clamp or remainder reset at the upper bound. Thus the serialized
pair is not necessarily a normalized hundredths value: negative remainders
become unsigned bytes 157..255. With unchanged returning callbacks, mana=0,
max=101, period=1 and modifier=-101 produces current=-1/remainder=255 and the
next admitted call produces current=0/remainder=54. — SAV-REGENSTORE-529

Health period=0 skips its arithmetic. Reached mana period=0 faults before any
mana store; preceding health stores lie before an intervening callback and
are not rolled back by this local body. Product/accumulator overflow wraps;
signed quotient overflow is excluded by the stable admitted word/rate domain.
Actual callback effects, first-load admission and saved post-fault state
remain Unknown. — SAV-REGENFAULT-530, SAV-REGENORDER-531

The Unit wire order is six two-byte members `94,96,98,9a,9c,9e`, then the two
one-byte members `a2,a3`; within this scalar run their offsets from member +84
are 16,18,20,22,24,26,28,29. The earlier raw 64 block at `d4` contains modifier
words `de/e2` at block offsets 10/14. These width-preserving transfers do not
rebuild or normalize the values. Wire order is not arithmetic store order,
and no safe invented period/remainder vector follows from it.
— SAV-REGENWIRE-532

## LOAD and derive

The named archive/load-hook bodies do not rebuild modifiers from
equipment references or directly invoke derive. A later derive consumes
the restored modifiers and active index. The exact first computed post-load
consumer is still Unknown, including city/world order and callbacks.
— SAV-HUMLOAD-445

The attack initializer clears 22 bytes, not its full 24-byte serialized extent;
the final two survive that helper. The modifier constructor separately clears
all 64 bytes. Safe authored values for the unnamed tails remain Unknown.
Capacity and second-component modifiers have fold consumers but no nonzero
producer in the selected paths; the General shadow and active-index mirror
are not interchangeable with their live counterparts. Later aliases, bulk
copies and non-enumerated producers remain open. — SAV-HUMGAPS-449

`R1569 -> R1482 -> R1560` repairs separate Position, actor-reference,
mover and order keys under the [field-specific key rules](token.md#identity-map). With ordinary
disjoint receivers it touches none of the selected 94 bytes and invokes no
Human derive. World resume depends on
YA1 InBattle, frontend mode/authority and the server run gate. City resume
drains queued commands around client restoration. Earlier document calls,
queues and callbacks remain outside that hook's negative scope.
— SAV-HUMRESUME-460 (reference-repair shorthand superseded by the
field-specific rules)

Group LOAD first invokes the embedded Group+20 list serializer, then the
separate Group+3c state serializer. Its actor archive read distinguishes an
existing object from a newly constructed object: the new object is registered
before its receiver+8 serialization and before Group insertion. Insertion can
remove old Group membership before assigning actor+70. These callbacks and
archive aliases precede later Player/world registration; they do not establish
the first Human stat read. — SAV-946

The world-present suffix repairs a copied thirteen-dword manager record:
ten nonzero keys at +4..+28 are replaced only on lookup hits. The following
manager callback only sets its own +a9c0 self-pointer. Another manager walks
its own entries and calls their +24; those entries are not established as
Humans. The no-world arm bypasses these callbacks, clears actor +40/+44/+5c,
and leaves the server run gate zero. Earlier archive and later frontend work
remain separate boundaries. — SAV-947

Resume setup sends opcode4 before the optional server bootstrap. If that
packet reaches the ordinary dispatcher unchanged, its ID selects a Player by
WORD+4 and its resume argument1 supplies zero to the entry setup flag. The
reached setup can project that Player's current +34 actor before ordinary
actor subticks. Earlier setup callbacks, transport arrival and continued
identity of +34 with the restored Human remain unproved. — SAV-948

The first computed read after LOAD remains Unknown. Earlier archive/world/UI
callbacks, phase6 Group work, phase12 projection/full ticks, earlier actors
and queued packets can precede the selected ordinary consumer. A same-Human
ordered trace with values at the first read or an explicit skipped tick is
still missing. There is no universal derive-before-read guarantee. — SAV-949

## Conditional consumers

These admitted paths compute from saved fields without a local derive call.

| admitted path | selected-field computation | boundary |
|---|---|---|
| actor-state sender, effective damage mask `200` | `(b4+b9+b7) mod 256`, `(b5+ba+b8) mod 256` | computed presentation, not combat; recipient/class mask applies |
| phase-12 Human full tick, positive health and deficit/period gates | signed `de/e2` regeneration | the same list iteration projects state before actor full tick |
| base Effect generic apply/remove | signed `d8` test/possible clear before eventual derive | attached Effect identity, mode and cadence select the path |
| special Effect identity 17 on Human | `e4` and `a4` add/subtract signed magnitude times 256 | modular word changes without generic derive |
| admitted strike | signed `c0` absorption, indexed resistance, `c6` secondary protection, `bb` elemental selector | physical/secondary/third-component gates differ |

The sender's virtual class test is not derive. The chargen routine that does
derive constructs a new temporary Human, so it cannot establish loaded-actor
ordering. Base Effect identity 8 apply also computes from `c6` without generic
derive; other Effect subclasses and earlier command order remain open.
— SAV-HUMPROJECT-461, SAV-HUMTICK-462, SAV-HUMSTRIKE-463

## Index bounds

Derive reads `word[a8+2*b6]` for every nonzero byte index before clearing its
secondary/elemental damage and defensive live fields. No upper bound occurs
in that indexing step. Indices 10,13,32,39,42 select respectively
`bc/bd`, `c2/c3`, `e8/e9`, `f6/f7`, `fc/fd` and affect computed to-hit/damage.
The resolver independently indexes a resistance byte at `target+ce+b6`
without a six-slot local bound. Ordinary production and native acceptance of
these out-of-range indices remain Unknown.
— SAV-HUMINDEX-464

An admitted third damage component expects `bb=1..5`; zero reaches a
diagnostic arm rather than selecting protection slot zero. No ordinary
`c2` damage arm is established in this resolver. Nonzero capacity and
secondary-modifier producers, ordinary meanings of residual shadows/tails,
and the absolute first loaded computed read remain Unknown. There is no
unconditional derive-before-read contract or safe zero vector.
— SAV-HUMSTRIKE-463, SAV-HUMFIRST-465

## Construction and first SAVE

The named Human block constructors preserve `bc/bd` while clearing `fc/fd`;
both pairs are serialized. Allocation contents and later Effect/indexed writes
do not supply a general first-SAVE vector. — SAV-HUMALLOC-504, SAV-HUMNEW-505,
SAV-HUMNEWSAVE-507

The shipped nonempty Human weapon definitions produce selectors 1..5. Signed attackType values below 10 are narrowed to
a byte; values at least 10 and removal produce 0. Positive types 10/42 do not
produce the corresponding malformed tail aliases. Negative/custom inputs,
archive restore, starting-skill aliases and other lifecycle producers need
their own value contracts. No safe zero vector follows. — SAV-HUMSEL-506,
SAV-HUMNEWSAVE-507

## Setup helper and starting-skill aliases

After its hero-setup gate, `R1562` writes six words at `102..10c` and
six bytes at `10e..113`, all 100, then calls Human derive through `+50`.
Its own writes preserve the eleven selected bytes `bc/bd`, `da/db`, `e8/e9`,
`f6`, `f7/f8`, `fc/fd` up to that dispatch. The callee's transitive execution
and the later mission lifetime remain separate boundaries. — SAV-898

The starting-skill input is distinct from active selector `b6`. Its local
store is `word[a8+2*(index&255)] = lowByte(level)` after clearing slots 1..5.
Indices 10/25/32/39/40/42 respectively overlap `bc/bd`, `da/db`, `e8/e9`,
`f6/f7`, `f8`, `fc/fd`. The selected packet arm forwards its skill byte;
ordinary production of these alias indices and earlier validation are not
established by those local windows. — SAV-899

First-SAVE values of these residual fields remain Unknown after intervening
callbacks and indexed writes. Constructor clears do not supply a safe vector.
— SAV-900

The named creation UI hit-test returns one of five marker indices or no match.
Its handler supplies skill index 1..5; draft initialization supplies 1. These
inputs exclude the residual aliases on that path. Attribute and preset handlers
forward stored selection+1 without a local bound. Other draft writes, current
frame identity and subsequent admission remain separate conditions. — SAV-934

The expanded modifier folds read capacity `da/db` into live `92` and secondary
damage `f7/f8` into live `b7/b8`. Their four local bodies preserve the eleven
residual bytes. Spellbook update loads its receiver as `u32[Human+140]`, uses
that same Human as a read argument and each Spell as its write receiver: the
derived Spell fields are bytes `09/0e/0f` plus word `10`. The row-derived skill
read uses a full 32-bit index. A separate write loads `mover = u32[Human+154]`
and stores `u8[mover+0a] = u8[Human+8c]`. Aliased storage, numeric callees and
the complete derive lifetime remain open. — SAV-935

Before the new builder returns, the same Human reaches item callbacks, owner
and group collections, further helper calls and state publication. Those handoffs
do not establish an empty interval before SAVE. Old-item and attached-Effect
receivers, allocation contents and later writes still prevent a general
first-mission-SAVE value contract for the six residual spans. — SAV-936
