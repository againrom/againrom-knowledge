# Spell training and weapon casting

[Reference](format.md)

## Training events

**Training is event-based, not one award per release.** Every apply from a book —
`caster+0x4c & 2` and `caster+0x68 == 0` — calls `caster->vt+0x68(target, spellId)`, an empty stub on
the monster vtable and `R1053` on both hero ones. It awards `round(manaCost / 2)` into the
spell's `Sphere`. An item cast keeps `actor+0x68` non-null and gets **none** of that immediate award.

Later `vt+0x64` calls are independent: positive direct damage supplies the resolved amount; Drain
Life supplies the positive capped transfer; Slow id 28, Stone Curse and Curse id 27 each supply
`trunc(target.healthMax × 0.03)` once per accepted target before attachment; Poison Cloud id 8
supplies every nonzero signed continuous tick while its recorded caster has health `>= 0`. Negative
caster health clears the pointer; zero health does not. The award is
`trunc(target.XPvalue × 0.5 × amount / target.healthMax + 1)`. Before that calculation, the damage
caller requires a victim owner with `+0x28 != 0`, and refuses when owner `+0x5c` and multiplayer
`server+0x0c` are both nonzero. The sink first requires recipient
`typeID` in `[0x21,0x3f]`; low-type map Humans receive no award. After that gate, with
`actor+0x4c & 4` set it maps the actual spell id to `Sphere`; with the bit clear the sink substitutes
`actor+0xb6`, the current weapon skill. Thus an admitted caster item trains a school only through
admitted later events, while an admitted fighter rider trains the equipped weapon for its physical
hit and every spell-side event.

There is no early refusal when Mind scaling produces zero. The chosen arm adds zero and runs the
ordinary refold. Nor is a negative amount refused: a custom negative item power can invert Poison's
signed magnitude, so the tick heals through subtraction and still calls the award path. The positive
cap leaves a negative value unchanged; the chosen arm subtracts it from slot and aggregate XP but
has no level-decrement arm. Shipped positive-power Poison items do not exercise that G2 boundary.

`vt+0x60` is a delayed kill award. Its caller requires a victim owner and applies the same
owner-`+0x5c`/multiplayer refusal, but not the damage caller's owner-`+0x28` test. The direct-damage
resolver leaves prior attribution for a null source or null source definition, clears `victim+0x40`
for a definition-bearing source without an owner, and otherwise writes the source. It then writes
damage kind for a bit-4 source or zero for a clear-bit source. A
PointEffect enters its post-payload tail only with a recorded caster and `+0x41 != 0`, making it
non-Defensive only. A now-null caster definition actively clears prior `victim+0x40`; otherwise a
non-null owner admits the recorded-caster and actual-id writes, while a null owner leaves prior
state. An AreaEffect uses a separate tail: Wall of Earth returns before payload, Light after payload
but before attribution, and every other arm requires a nonzero target movement-domain byte plus a
recorded caster. A null definition or null owner clears `victim+0x40`; a surviving owner admits the
recorded-caster and actual-id writes. The caller zero-extends the byte, so ordinary domains 1..3 and
custom `0xfe` pass, while only zero fails. Movement domain is not health, so a lethal application on
an ordinary domain is attributed. Drain Life jumps to the apply epilogue
without either envelope and writes no
fresh attribution. Poison Cloud seeds id 8 during area application, but later ticks do not refresh
it; intervening combat can therefore change its eventual kill recipient. An **area** application
remains per target, and Poison Cloud is per qualifying tick; one release may pay zero, one or many
awards.

Installed Data.bin castSpell weapon entries include Human staffs and the
Catapult/Ballista Unit weapons. Authored powers are
`{1,5,10,15,25,30,34,35,40,50,60,63,65,70,82,90,98,99}`; they all enter the same unclamped signed
`i16` item-power route. The live shop is an additional producer: unflagged generated weapons choose
ids `{1,11,13,14,20}` and a price-derived random power capped at 100. Plain Unit vtable
`L00001` maps kill, damage and cast-award slots to empty
stubs, so those siege actors can release the rider but never train. Humanoid recipients then face
the type-id and later gates above.

An item apply passes the weapon's existing `Spell` to the same fill as a book cast. It reloads and
power-adjusts serialized `Spell+0x09` Max Range, then rewrites `+0x0e` damage base, `+0x0f` spread
and `+0x10` duration scratch. The item effect, `weapon+0x80` pointer and `Spell` identity survive;
the four filled fields do not remain unchanged.

`Unit::Serialize` preserves training as six levels, six per-skill experience dwords and one
aggregate. It also stores `actor+0x68` as an object reference, while `actor+0x64`, kill-credit
`actor+0x40` and attribution `actor+0x48` are raw `u32`, raw `u32` and `u8`. The credited actor is
remapped separately through the saved-address map by the world LOAD lifecycle;
the same lifecycle repairs+64/+68, with missing keys cleared. Wire representation
does not establish pointer validity after later callbacks. — SAV-908

**Prismatic Spray** (`R0269`) selects at most `min(power/20 + 2, 7)` total victims and
applies to each in selected order. Its secondaries come from group sight; this count is not a
radius and the ten AreaEffect spells use a different path.

The per-spell arms are **power consumers too** — twelve distinct expressions in all, of which the
four above are the shared ones. The magnitudes an arm writes to `effect+0x40`:

```
Freezing Cloud        -(power/15 + 1)          Shield              power/10 + 3
Light                   power/30 + 1           Haste / Slow      +-(power/15 + 1)
Darkness               -(power/30 + 1)         Bless / Curse     +-((power*4)/5 + 20)
the four Protections    power/2                Poison Cloud       ftol(template * f)
Fire Sacrifice       two totals, each capped at 0x200
```

## Weapon casting

An item may carry an effect of kind `castSpell` (41). That effect is **never applied** — the
dispatch's arm for it is a bare jump to the epilogue. It is read as data instead:

- `R0858(item)` finds it by kind, and its `+0x42` is the power, replacing the [spell-power](casting.md)
  expression;
- `effect+0x40` is the spell **id**, a byte; `R1011` reads it and `R1012` hangs a
  constructed `Spell` on `weapon+0x80`. Nothing anywhere resolves that spell by **name** — the
  name-taking twin `R1045` is dead code, unreached by a call, by a pointer table in any
  section, or by any `rel32` in `.text`;
- **which class fires it decides which of two paths runs, and both exist.** The predicate is one:
  `actor+0x4c & 4` (set = caster, fixed at its write — it is OR'd in only when the streamed
  `ManaMax` column is positive, and the spellbook is allocated in the same arm).

```
actor tick R0037, state 3 (attack), weapon != 0, weapon+0x80 != 0:

  caster    (actor+0x4c & 4)  -> state := 0x0d, +0x64 = weapon+0x80, +0x68 = weapon
                                 the strike is NOT called at all
                                 R0268 validates; wind-up = actor+0x134 frames
                                 then Spell::Apply                       <- the cast
  non-caster                  -> R0001 -> R0246            <- the strike
                                 ordinary rider iff dealt > 0 and post-hit health > 0
                                 Fire Ball rider when either test fails too
                                 Spell::Apply                            <- the rider
                                 then physical XP iff dealt > 0, saved pre-hit gate clear,
                                      and post-spell health > -10
```

`R0001` has exactly one caller and is the tick's **else**, so a caster wielding such a
weapon reaches no to-hit roll, no damage roll and no absorption: it **cannot miss** and deals no
weapon damage. Both paths clear `+0x64`/`+0x68` afterwards, and an item whose kind is `0x0e` is
destroyed with its `Spell` after the cast — a staff is not.

The selected-inventory item Cast command is not a third shipped weapon path. Its UI arm requires
descriptor bits `0x10|0x01`; the effect walk supplies bit 4, but the Weapon descriptor clears the
byte and supplies only bits 1 and 2 from `sutableFor`. Bit 0 comes only from a base MagicItems
descriptor. The server's latent item-order arm would accept a crafted kind-2 command, so this is an
engine/client G2 boundary, not a Data.bin switch.

Fire Ball is the fallback rider whenever the positive-damage/positive-health conjunction fails. It
therefore triggers after a miss or zero damage and on living as well as dead targets. The item spell
runs before physical experience; if it moves health to `-10` or below, the later physical award is
suppressed.

Spell id 14 separates the routes. The caster validator runs the Prismatic Spray fan during
admission while the item context is live; the later common wrapper refuses id 14 and prevents a
second fan. The fighter rider never enters that validator and reaches only the refusing wrapper,
so its id-14 rider applies nothing.

The **generated** starting kit gives a caster a staff and only a staff (`Wood Staff
{castSpell=Fire_Arrow:20}` above tier 10, `:10` at or below, plain or `Uncommon` by gender) and a
non-caster one of five swords/axes/maces/pikes/bows. That is a distribution, not a rule: nothing on
the equip path refuses an item on account of class.
