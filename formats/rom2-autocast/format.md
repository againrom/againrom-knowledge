# ROM2 automatic casting

The selected RU support routine reads a level from the actor's owning Player
preferences. The level is separate from the command's selected offensive spell.
Three cumulative masks select Minimum, Average and Maximum. The local EN/RU
cycle and constructor stores agree; full policy equivalence between the two
images remains Unknown. — R2-ENGINE-345, R2-ENGINE-350, R2-SESSION-150

Retained local High and bounded Medium grades use first observations before
the 160-window hard limit. Acquisition later reached a conservative minimum
of 169 windows and breached the reserve and stopping rules; the completed
late reproduction is excluded from grade rationales. Coordinate-getter
meanings and units remain Unknown; High range evidence is the local packed-byte
arithmetic. — R2-ENGINE-345, R2-ENGINE-346, R2-ENGINE-347, R2-ENGINE-348,
R2-ENGINE-349, R2-ENGINE-350, R2-ENGINE-351, R2-SESSION-150, R2-ASSET-085

## Control and setting

| Level | Level mask | EN label | RU source view |
|---|---|---|---|
| Minimum | `0x08` | Minimum | Минимальный |
| Average | `0x18` | Average | Средний |
| Maximum | `0x38` | Maximum | Максимальный |

The installed `text/dialogs.txt` zero-based physical rows 171..173 supply these labels; row 170
is `AutoCast mode:` in EN and `Автокастинг:` in a CP1251 view of RU source.
That view does not establish rendered glyphs or the runtime code page. The
selected RU options code populates control `0x8d` with rows 171..173 and applies
selection indices 0/1/2 as these masks. The source-to-native-table join is
Medium. — R2-ASSET-085, R2-ENGINE-345

Option apply independently adds affinity masks `0x01`, `0x02` and `0x04`
from the three checkbox values to the selected level mask. Previous affinity
bits are replaced by those values. This interpretation is Medium.
— R2-ENGINE-345

The cycle routines EN `L2.01501` and RU `L2.01502` change masked level
`0x08 -> 0x18 -> 0x38 -> 0x08` and preserve low affinity bits 0..2. The global
DWORD is EN `L2.01503`, RU `L2.01504`. The selected RU input arm calls the
cycle, but its physical key is Unknown. The support reader reaches byte
`+0x20` through actor owner `+0x14`, Player pointer `+0x34`. This supports a
Player preference scope at Medium; a complete selection, Group or unit-control
contract is Unknown. Class eligibility and per-unit setting availability are
Unknown; the selected admission helper requires the named spell in the actor's
book. — R2-ENGINE-345, R2-ENGINE-348

Global constructors initialize preference DWORD `+0x24` to `0x19`, which
combines the Average mask and own-target bit. This is High for the local
constructor stores. Registry overrides and per-Player initialization prevent
using that store as the proved default of a new hero, hire, summon or converted
unit. Those defaults remain Unknown. — R2-SESSION-150, R2-ENGINE-345

## Support policy

RU `L2.01512` stops after its first prepared support order. Admission requires
the spell in actor book `+0x140` and signed current mana `+0x9a` at least its
signed cost `+0x0c`. Each stage also has the independent minimum-mana gate
below. These are thresholds on the same current value, not additional mana
charges. Behavior is Medium because list construction, call reachability and
later execution remain incomplete. — R2-ENGINE-346, R2-ENGINE-347,
R2-ENGINE-348

| Stage | Additional gate | Attempt order |
|---|---|---|
| Minimum | level bit `0x08`; mana >=5 | self Heal 24 if HP differs from maximum; self Shield 27 if absent; Heal 24 on a selected candidate |
| Average | level bit `0x10`; mana >=10 | for each candidate, Bless 20 then Haste 21 |
| Maximum | level bit `0x20`; mana >=40 | for each candidate, Protection from Air 13, Earth 19, Fire 4, Water 8, then Invisibility 12; fallback self Invisibility 12 |

The candidate Heal path requires the selected list count to exceed one. It
chooses the first strict minimum of signed integer `HP / maximumHP`, then
requires that integer quotient to be below 1. Equal scores keep the first
candidate. It does not rank wounded targets by a floating health fraction.
Complete list membership and order remain Unknown. — R2-ENGINE-346

Average and Maximum are target-major: all listed spells for one candidate are
tried before advancing to another candidate. Their corresponding effect bit
must be absent. Bless has an additional predicate; target Invisibility excludes
self and needs a nonzero masked relation value. The fallback self Invisibility
is a separate attempt. — R2-ENGINE-347

## Affinity and range

RU `L2.01519` accepts the same actor unconditionally. Other acceptance branches
use the same owner pointer with preference mask `0x01`, relation mask `0x02`
with preference mask `0x02`, or `(relation & 3) == 0` with preference mask
`0x04`. The selected options
labels bind these bits to Own, Allies and Neutral at Medium. They do not define
every relation-matrix value. — R2-ENGINE-348, R2-ASSET-085

RU `L2.01520` passes the packed values returned by `L2.01521` to `L2.01522`.
The metric isolates each value's two unsigned low bytes, takes their absolute
differences through `L2.01523` and returns the maximum. The wrapper compares
that byte with spell range byte `+9`. Exact local arithmetic is High. The
getter's returned-value meaning, coordinate units, candidate visibility and
any sight filtering inside the list builders remain Unknown. Self Heal, self
Shield and fallback self Invisibility bypass this range wrapper.
— R2-ENGINE-348

## Prepared order and pending condition

RU `L2.01524` prepares actor command byte `+8 = 8`, target pointer `+0x28` and
spell pointer `+0x30`. It records the spell ID at target command `+0xa4` and the
current tick at target command `+0xa8`; actor command `+0x14` ends with the
spell range byte. These local writes are explicit; resulting behavior is Medium.
Successful preparation does not
establish that an effect has already occurred. — R2-ENGINE-349

The listed Bless, Haste and extended attempts skip a candidate when its pending
field equals literal 20 and the absolute difference from the current tick is
at most 20. This is a pending-target condition, not a proved universal spell
cooldown. Self Heal, self Shield and fallback self Invisibility do not contain
it. Movement to cast, cast cadence, player-order interruption and resumption
remain Unknown in the selected preparation and partial command-eight paths.
— R2-ENGINE-349

## Offensive spell and computer Players

RU `L2.01527` handles the already-selected spell ID in command byte `+0x78`.
Its tables dispatch IDs 1..28. It does not iterate the book to choose another
offensive spell. Its selected target arms and the support level are distinct;
the complete user-selection and execution contract remains Unknown.
— R2-ENGINE-350

The selected dispatch branches distinguish owner Player DWORD `+0x2c == 0`
from nonzero. The nonzero path reaches RU `L2.00056`, which builds an
affordable non-defensive book-spell list and calls an unread index helper.
Its distribution is Unknown. A Mind-dependent early branch calls `L2.01530`
and returns zero; the recursive outcome is Unknown. Separate support routines
are `L2.01533`/`L2.01513`. These selected branches do not establish
that every computer Player shares the three human support levels. — R2-ENGINE-351

## Persistence and locale boundary

EN Player pointer `+0x34` reaches a preferences serializer that transfers 36
raw bytes. That interval includes the support preference byte `+0x20`.
The local width is High; ordinary SAV reach and the field interpretation are
Medium. Global Windows registry `AutoCasting` is a four-byte value for global
preferences `+0x24`. Mission and town preservation and later overrides remain
Unknown. — R2-SESSION-150

The paired EN/RU constructor and level-cycle sequences agree after address
normalization. The 389-instruction registry spans have one non-address
difference: an embedded-base addition is `0x41c` in EN and `0x434` in RU.
The local AutoCasting transfers still have the same width and offset; this
does not prove equivalent external callees. Label source strings differ by locale. The full support,
offense, command execution, SAV and lifecycle paths were not compared across
both images, so whole-policy equivalence remains Unknown. — R2-ENGINE-345,
R2-SESSION-150, R2-ASSET-085
