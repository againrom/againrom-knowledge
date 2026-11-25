<a id="hero--character-generation-and-the-derived-stat-graph"></a>

# Human characters and combat

Human state includes four primary attributes, six skill/experience slots,
equipment, modifiers and derived combat fields. Character generation and
map placement populate different initial inputs; the derive routine and
later consumers have explicit ordering. — HERO-STAT-001, HERO-COST-002,
HERO-BUDGET-004, HERO-EQUIP-017 (amended), HERO-MOD-016

Combat cadence is specified by HERO-CADENCE-112 through HERO-CADENCE-115 (amended). One actor routine runs a blow, a hand cast and an item cast on one phase byte and one
countdown; the branch differences are HERO-196, the distance term
HERO-197 (a blow adds it; Medium for the meaning of the distance; an item cast diverted at the phase-5 zero keeps it), the damage moment HERO-198 and shipped repeat intervals HERO-199.
Spell behavior is in [MAGIC](../magic/format.md); saved raw state is in
[SAV](../sav/format.md).

## Inputs and state transitions

| Input/state | Consumer |
|---|---|
| Body, Reaction, Mind, Spirit | Derived-stat graph, with caps before dependent formulas |
| Six skill/experience slots | Active weapon/school skill and award rules |
| Equipment and modifier block | Add/remove folds followed by the applicable derive path |
| Health, mana, target and action time | Regeneration, admission and ordered combat phases |

Character generation assigns class/sex/name and purchased attributes; map
placement has separate overrides. Equipment, spell effects and experience
can change inputs later. Apply each producer's specified recompute order;
restoring serialized state is a separate path, not a request to recreate all
fields from definitions. — HERO-STAT-001, HERO-EQUIP-017 (amended), HERO-MOD-016

## Reference map

| Reference | Contents |
|---|---|
| <a id="integer-conventions--read-this-before-any-formula-below"></a><a id="1--the-four-primary-stats"></a><a id="2--character-generation"></a><a id="6--the-class-discriminator"></a><a id="the-name-hero-name-079-hero-typed-080"></a><a id="integer-conventions"></a><a id="primary-attributes"></a><a id="character-generation"></a><a id="what-character-generation-puts-in-his-hand"></a><a id="class-discriminator"></a><a id="name-hero-name-079-hero-typed-080"></a><a id="character-generation-controls-hero-chargen-082hero-chargen-085"></a><a id="starting-templates-and-non-item-state"></a> [Character generation and identity](generation.md) | Attribute purchase, class/sex/name and identity |
| <a id="3--the-six-skill-slots"></a><a id="4--experience-and-how-a-skill-level-moves"></a><a id="skill-slots"></a><a id="experience-and-skill-levels"></a><a id="the-raise--R0810-vt0x5c-of-both-human-classes"></a><a id="the-loss--R0825-charged-by-conditional-repair"></a><a id="the-purchase--R0838-one-command"></a> [Skills and experience](experience.md) | Six-slot storage, level rules and awards |
| <a id="5--the-derive-R0280-vt0x50-of-both-human-classes"></a><a id="5a--step-15-the-modifier-fold-R0840actor0xd4-actor"></a><a id="5b--who-fills-the-modifier-block"></a><a id="derived-stat-sequence-R0280-vt0x50-of-both-human-classes"></a><a id="step-0-in-full--what-bounds-a-stat-in-principle"></a><a id="when-the-derive-runs-on-a-map-placed-person-and-what-its-inputs-are-hero-hp-071-hero-hp-072"></a><a id="modifier-fold-R0840actor0xd4-actor"></a><a id="modifier-producers"></a><a id="local-equipment-sequence"></a> [Derived attributes and modifiers](derived.md) | Ordered formulas, caps, equipment and effect folds |
| <a id="6a--how-the-character-is-drawn-hero-appear-040hero-appear-046"></a><a id="6b--the-twelve-slots-the-u16-in-each-and-where-the-sheets-come-from-hero-appear-047hero-appear-055-hero-figure-062"></a><a id="6c--composing-the-figure-two-surfaces-and-one-of-them-is-a-hit-map-hero-figure-057hero-figure-064"></a><a id="world-appearance-hero-appear-040hero-appear-046"></a><a id="figure-equipment-slots-hero-appear-047hero-appear-055-hero-figure-062"></a><a id="figure-composition-and-hit-map-hero-figure-057hero-figure-064"></a> [World and character-panel appearance](appearance.md) | World sprite and panel selection |
| <a id="6d--persistent-companion-joins"></a><a id="persistent-companion-joins"></a> [Persistent companion joins](party.md) | NPC/mercenary joins and mission continuity |
| <a id="7--resolving-one-hit-R0265-the-actors-vt0x4c"></a><a id="8--regeneration-R0654-the-actors-vt0x14"></a><a id="10--the-combat-loop-around-that-hit"></a><a id="hit-resolution-R0265-the-actors-vt0x4c"></a><a id="regeneration-R0654-the-actors-vt0x14"></a><a id="combat-loop"></a> [Combat and regeneration](combat.md) | Attack phases, damage, regeneration and callbacks |
| <a id="9--what-a-consumer-still-cannot-reproduce"></a><a id="unknowns"></a> [Unknowns](limits.md) | Unspecified fields and bounded execution paths |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.

## Network character file

A network character lives in the network character file in the current directory. Its name is
built from the format string at `L17914`: a fixed prefix, the slot index as four zero-padded
decimal digits, and the character-file extension. The text was read outside the committed
evidence (Medium). — HERO-213

| order | source | bytes | note |
|---|---|---|---|
| 1 | tag `AlCh` (`0x68436c41`) | 4 | the list accepts only this tag |
| 2 | store `+0x08` | 8 | |
| 3 | store `+0x10` | 32 | name, read by the list at file `+0x0c` |
| 4 | store `+0x30` | 32 | |
| 5 | flags `+0x70` | 4 | file `+0x4c`; bit 0 set ends the file here after row 8 |
| 6 | store `+0x74` | 4 | |
| 7 | store `+0x90` | 28 | starts with the four creation stats (Medium) |
| 8 | store `+0xac` | 28 | |
| 9 | store `+0x78` | 56 | starts with 16 Player-scalar bytes (Medium). The null-argument write stores 100, 0, 0, 0 in the object; whether those bytes reach the file depends on flag bit 0, whose setter was not read (Unknown) |
| 10 | six blocks of a second object | 16, 32, 16, 16, 4, 4 | meanings unread |
| 11 | store `+0xac` | 32 | |
| 12 | store `+0x8c` | 4 | four bytes packed from the frame |
| 13 | twelve slot dwords, each nonzero one followed by its record | 48 + records | |
| 14 | archive count, then the archive | 4 + 2 × count | the hero graph |

The creation pages are shared with the campaign; outside the campaign phase Play writes this
file. In the campaign phase the writer `R2314` returns at once; other writers of the file
were not searched (Medium). The client rewrites the file when it is connected and `L17915`
accepts a character packet; the match of the server's packet to that handler is Medium. The
server exports from six direct call sites: join, reporter (network mode), the primary's
teardown (network mode, after adding 1 to `player+0x54`), the actor tick, disconnect (when
`server+0x150 == 0` in network mode) and `R2129`, a function that sets four stats to 10,
10, 10, 1 (its role is unread). Register-indirect calls were not searched. — HERO-214, HERO-215
