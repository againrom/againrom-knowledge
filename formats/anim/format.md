<a id="anim--what-drives-a-units-animation--specification-partial"></a>

# Drawable animation state

Each drawable advances a presentation state block using a game-tick driver.
Action messages fill that block; the state selects phases and frames through
the class sheet. Simulation attack/movement timing remains a separate input.
— ANIM-CLOCK-001, ANIM-OBJ-008 (its object-arm unreachability clause is
withdrawn), SPR256-UNIT-024

The complete-period formula in HERO-CADENCE-023 is retracted. This reference
does not infer simulation duration from a frame count. The meanings that
distinguish messages 0x86/0x8a/0x8b/0x8c from 0x6b/0x71 and the every-32-tick
helper remain Unknown; native clock appearance has not been established.

## Drawable state

| Runtime member | Meaning |
|---|---|
| +0x6c / +0x70 | Facing / phase |
| +0x74 / +0x84 | Draw state / action code |
| +0x94 / +0xa0 | Action clock / remaining ticks |
| +0x88 / +0x8c | Remaining X/Y displacement in 1/256-cell units |

Action messages populate the drawable. The paced0x401 tick advances it;
idle/attack/cast/death timelines use ticks, while walking accumulates distance
and advances once per 1/16 cell. Drawing resolves the resulting state through
class phase blocks. Registration selects the collection and composition pass.
— ANIM-CLOCK-001, ANIM-OBJ-008 (unreachability clause withdrawn), ANIM-WALK-013

## Reference map

| Reference | Contents |
|---|---|
| <a id="the-one-thing-a-consumer-must-not-get-wrong"></a><a id="the-state-block-on-the-drawable"></a><a id="the-driver-once-per-game-tick-per-drawable"></a><a id="the-eight-action-arms"></a><a id="the-idle-run--state-0-and-the-fidget"></a><a id="the-messages-that-fill-the-block"></a><a id="the-death-chain-end-to-end"></a><a id="what-a-consumer-needs"></a><a id="the-act-state-a-refused-actor-is-parked-in-anim-park-039"></a><a id="state-and-identity"></a><a id="drawable-state-block"></a><a id="tick-driver"></a><a id="action-states"></a><a id="idle-and-fidget"></a><a id="directions"></a><a id="action-messages"></a><a id="death-and-corpse-phases"></a><a id="objects"></a><a id="state-requirements"></a><a id="refused-actor-state-anim-park-039"></a> [Actor action phases](actors.md) | Tick driver, action messages, phase/frame selection |
| <a id="the-damage-numeral-identity-merge-and-pacing"></a><a id="damage-numerals"></a> [Damage numerals](numerals.md) | Damage text lifetime and notification paths |
| <a id="a-projectile--a-different-animation-model-entirely"></a><a id="the-projectile-frame-clock"></a><a id="the-polyline-draw-of-pictures-34-and-36"></a><a id="reviewed-spell-effect-presentation-anim-044anim-047-late-pass-scope-amended"></a><a id="projectile-state"></a><a id="projectile-clock"></a><a id="picture-3436-polylines"></a><a id="spell-effect-presentation-anim-044anim-047-late-pass-scope-amended"></a> [Projectiles and effect presentation](projectiles.md) | Projectiles, stone state and effect drawing |
| <a id="drawable-registration-and-cell-composition"></a> [Drawable registration and composition](composition.md) | Registration selectors, phase/cell passes and shadows |
| [State-projector callers](state-projector.md) | Effective-mask admission, finite per-site predicates, restoration and client stage bounds |

The state-projector contract distinguishes a projector call from an ordinary
state-sync send, and server stage from the drawable's preceding stage.
Its finite caller table and native helper bounds are ANIM-129, ANIM-130,
ANIM-131, ANIM-132 and ANIM-133. ANIM-126's omitted-mask clause is narrowed;
ANIM-095's three-helper Unknown is resolved for native receiver bindings.

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.
