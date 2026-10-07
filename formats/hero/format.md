<a id="hero--character-generation-and-the-derived-stat-graph"></a>

# Human characters and combat

Human state includes four primary attributes, six skill/experience slots,
equipment, modifiers and derived combat fields. Character generation and
map placement populate different initial inputs; the derive routine and
later consumers have explicit ordering. — HERO-STAT-001, HERO-COST-002,
HERO-BUDGET-004, HERO-EQUIP-017, HERO-MOD-016

Combat cadence is specified by HERO-CADENCE-112 through HERO-CADENCE-115.
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
fields from definitions. — HERO-STAT-001, HERO-EQUIP-017, HERO-MOD-016

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
