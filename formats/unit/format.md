<a id="unit--the-non-hero-actor-template--instance--combat-inputs-partial"></a>

# Non-hero actors and structures

Actor creation binds Data.bin definitions, ALM overrides and owner state to
simulation fields. The Units arm streams its defined slots once and has no
Human derive routine. Building interaction and cell registration follow
separate class predicates. — UNIT-STREAM-001, UNIT-EQUIP-005,
UNIT-SPELL-007

Human state is in [HERO](../hero/format.md); item effects and equipment are
in [ITEM](../item/format.md).

<a id="what-this-covers-and-what-owns-the-rest"></a>

## Creation and interaction inputs

| Input | Role |
|---|---|
| Data.bin Units row | Base actor slots, equipment and spell definitions |
| ALM placement | Class, cell, owner, group and conditional overrides |
| Owner and difficulty | Creation-time health/combat adjustments |
| Building class and cell references | Interaction, damage and footprint registration |

Construct the actor, bind its definition, stream the Units slots, construct
equipment/spells and perform the one-time fold. The Units arm has no Human
derive routine. Later combat and Building operations use their own class
gates, while presentation consumes separate drawable state.
— UNIT-STREAM-001, UNIT-EQUIP-005, UNIT-SPELL-007

## Reference map

| Reference | Contents |
|---|---|
| <a id="which-file-carries-which-number"></a><a id="creation-in-order"></a><a id="the-slot-map-units"></a><a id="there-is-no-derive"></a><a id="the-combat-inputs"></a><a id="corpus-both-shipped-roots-en-and-ru-worldres-different-files"></a><a id="the-owner-actor0x14"></a><a id="the-setting-that-scales-a-unit"></a><a id="what-the-placement-record-overrides"></a><a id="open"></a><a id="definition-sources"></a><a id="creation-sequence"></a><a id="units-slot-map"></a><a id="derived-state-boundary"></a><a id="combat-inputs"></a><a id="installed-units-definitions"></a><a id="owner-field-actor0x14"></a><a id="unit-scaling-setting"></a><a id="placement-overrides"></a><a id="unknowns"></a> [Definitions and creation](creation.md) | Definition streams, constructors, equipment and difficulty |
| <a id="clickable-structures"></a><a id="scope-and-interaction"></a><a id="physical-orders-against-a-building"></a><a id="multi-cell-building-area-targets"></a><a id="mission-10-cell-entry-caster"></a> [Buildings, interaction and cell casters](structures.md) | Interaction selectors, damage, casters and footprint cells |
| <a id="what-the-information-display-is-given"></a><a id="what-class-it-is-drawn-as-unit-appear-030"></a><a id="the-name-a-unit-is-shown-by-unit-name-039unit-nametab-041"></a><a id="information-display-state"></a><a id="drawable-class-unit-appear-030"></a><a id="display-names-unit-name-039unit-nametab-041"></a> [Actor presentation and names](presentation.md) | Panel fields, names and drawable selection |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.
