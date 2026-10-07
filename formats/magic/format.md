<a id="magic--the-spellbook-the-cast-and-the-resistance"></a>

# Spells, effects and casting

Spell definitions, learned objects, actor state and cast inputs determine
application and presentation. Effects have their own duration, ownership and
action-gate rules. — MAGIC-SPELL-001, MAGIC-CAST-003, MAGIC-EFFMODE-009, (Delivery/timing clause narrowed by MAGIC-CASTCLOCK-171.)
MAGIC-CADENCE-126, MAGIC-CADENCE-127

Definition tables are in [Data.bin](../databin/format.md), actor state in
[HERO](../hero/format.md), and archive state in [SAV](../sav/format.md).

## Cast inputs and object state

| Input/state | Meaning |
|---|---|
| Spell+0x08/+0x09/+0x0a | ID, range and Defensive bytes |
| Spell+0x0c | u16 mana cost |
| Spell+0x0e/+0x0f/+0x10 | Damage base/spread and duration scratch; not serialized |
| Actor+0x140 | Spellbook indexed by spell ID |

Ordinary power is `clamp(skill[Sphere]+Mind-30,0,100)`; item castSpell
power takes its signed i16 path without that clamp. Cast gates run before
application, which chooses immediate or projectile delivery and then the
spell-specific effect/lifetime path. Presentation and action restrictions
consume the resulting effects separately. — MAGIC-SPELL-001, MAGIC-CAST-003, (Delivery/timing clause narrowed by MAGIC-CASTCLOCK-171.)
MAGIC-EFFMODE-009

## Reference map

| Reference | Contents |
|---|---|
| <a id="1--the-spell-table"></a><a id="2--a-spell-and-a-spellbook"></a><a id="3--the-power-what-a-school-skill-buys"></a><a id="4--casting"></a><a id="5a--what-decides-whether-an-ai-actor-casts"></a><a id="8--what-mind-and-spirit-do-and-what-they-do-not"></a><a id="arithmetic"></a><a id="spell-definitions"></a><a id="spell-objects-and-spellbooks"></a><a id="spell-power"></a><a id="cast-sequence"></a><a id="ai-cast-selection"></a><a id="mind-and-spirit"></a> [Spell objects and casting](casting.md) | Definition columns, spellbook, power, gates and cast cadence |
| <a id="5--applying"></a><a id="6--an-effects-duration-model"></a><a id="effect-ownership-and-credit"></a><a id="application"></a><a id="non-spell-potion-lifetime"></a><a id="prismatic-sprays-selected-victims-magic-spray-134magic-spray-137"></a><a id="effect-duration"></a> [Application and ordinary Effect lifetime](application.md) | Delivery, resistance, awards and effect lifetime |
| <a id="10--the-two-pictures-a-spell-has"></a><a id="11--what-a-cast-draws-and-on-which-tick-it-applies"></a><a id="12--the-map-object-a-cast-makes-and-how-long-it-lives"></a><a id="spell-icons"></a><a id="cast-presentation-and-application-tick"></a><a id="spelleffect-map-objects"></a> [Icons, cast presentation and map objects](presentation.md) | Icons, map objects and cast presentation |
| <a id="13--staged-area-effects"></a><a id="staged-area-effects"></a><a id="cell-aliases-and-direct-damage-dispatch"></a><a id="stage-clock-and-cell-consumer"></a><a id="fire-sacrifice"></a><a id="acid-stream"></a><a id="meteor-storm"></a><a id="table-layout-and-customisation-limits"></a> [Staged area effects](area.md) | Walls/clouds, duration, stages and area visitation |
| <a id="14--the-drawn-path-of-a-travelling-projectile"></a><a id="projectile-trajectories"></a> [Projectile trajectories](projectiles.md) | Missile creation, paths and impact |
| <a id="15--what-a-lasting-effect-draws-on-its-actor"></a><a id="actor-effect-marks"></a><a id="the-two-states-an-actors-marks-are-held-in"></a><a id="lifecycle"></a><a id="rebuild"></a><a id="drawing"></a><a id="the-four-protections"></a><a id="shield-bless-and-curse"></a><a id="poison_cloud-heal-and-drain_life"></a><a id="two-spells-that-change-the-actor-instead"></a><a id="customisation-limits"></a> [Actor effect marks](marks.md) | Protection, Shield, Bless and other actor marks |
| <a id="16--what-a-lasting-effect-does-to-the-actors-actions"></a><a id="customisation-limits-1"></a><a id="effect-action-gates"></a><a id="the-gate"></a><a id="the-one-escape-and-its-bound"></a><a id="what-reaches-the-gate-and-what-does-not"></a><a id="the-observable"></a><a id="the-ai-side"></a> [Effect action gates](gates.md) | Freeze/stone and actor action restrictions |
| <a id="9--what-a-consumer-still-cannot-reproduce"></a><a id="unknowns"></a> [Unknowns](limits.md) | Unknown consumers and runtime limits |
| <a id="fire-wall-and-poison-cloud-overlap"></a> [Fire Wall and Poison Cloud overlap](overlap.md) | Repeated application, attachment, stored power and expiry |
| <a id="weapon-casting"></a><a id="7--a-weapon-that-casts"></a> [Spell training and weapon casting](training.md) | Event awards, attribution and item-cast state |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.
