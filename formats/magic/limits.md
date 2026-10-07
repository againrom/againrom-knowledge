# Unknowns

[Reference](format.md)

## Remaining boundaries

- The complete visibility rejection for an invisible actor remains Unknown.
  Located consumers of actor+0x144 bit 15 cancel it; they do not by themselves
  define the visibility test.
- The full meaning of caster+0x3c remains Unknown. Its nonzero gate suppresses
  non-projectile delivery and removes the caster from Effect+0x44.
- The mapping between Distribution system values and the labels Point,
  Round, Long, Phase and Hang On Unit remains Unknown.
- The producer of block+0x11/+0x12 damage scratch, which can be written on
  a missed swing, remains Unknown.
- Complete spellbook/weapon-spell caller coverage, the meaning of the
  Control Spirit selector actor+0x13c, and the AI's picked-spell helper
  R0018 remain Unknown outside the named paths.

The Poison valid-text numeric path is specified by MAGIC-POISONINPUT-157,
with substituted CString/decimal-library behavior. Native parser behavior
and MAGIC-EFFECT-015's separate Stone Curse conversion boundary remain
Unknown. Area-effect geometry is in [staged effects](area.md), and projectile
motion/presentation in [trajectories](projectiles.md) and
[animation](../anim/projectiles.md).
