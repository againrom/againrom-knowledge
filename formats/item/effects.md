# Equipment grammar and item formulas

[Reference](format.md)

## Equipment-cell effects (`ITEM-EFFGRAM-070`…`ITEM-EFFSAVE-077`)

The first `{` in a `Humans` or `Units` equipment cell ends the right-trimmed item head. The effect
tail begins after it and ends at the first `}`; absent close means end-of-string and text after the
first close is ignored. The tail is split on commas from left to right. Each accepted element is
appended to an ordered `CObList`; unknown elements are skipped independently, and valid duplicates
remain in order. The parser has no count comparison or fixed element budget.

An element is case-folded and has `key=operand[:suffix]`. The first `=` and first `:` are the
delimiters. Index 0 of the 50-entry key table is a rejected sentinel; kinds 1..49 are:

```text
 1 price              2 body               3 mind               4 reaction
 5 spirit             6 health             7 healthMax          8 healthRegeneration
 9 mana              10 manaMax           11 manaRegeneration  12 toHit
13 damageMin         14 damageMax          15 defence           16 absorbtion
17 speed             18 rotationSpeed      19 scanRange         20 protection0
21 protectionFire    22 protectionWater    23 protectionAir     24 protectionEarth
25 protectionAstral  26 fighterSkill0      27 skillBlade        28 skillAxe
29 skillBludgeon     30 skillPike          31 skillShooting     32 mageSkill0
33 skillFire         34 skillWater         35 skillAir          36 skillEarth
37 skillAstral       38 itemLore           39 magicLore         40 creatureLore
41 castSpell         42 teachSpell         43 damage            44 damageFire
45 damageWater       46 damageAir          47 damageEarth       48 damageAstral
49 damageBonus
```

| Kinds | Operand | `Effect+0x40..+0x43` |
|---:|---|---|
| 1..40, 49 | signed-decimal prefix | mode 0/8: i32; modes 1/2/4: observed i16 magnitude + u16 ticks |
| 41 | spell name, optional power | u16 spell id + observed i16 power |
| 42 | spell name, optional mode | u16 spell id + optional u16 ticks |
| 43..48 | first-hyphen `low-high`, optional mode | u8 low + u8 `(high-lowByte)` + optional u16 ticks |

Exact `permanent` and `singleuse` select modes 0 and 8. Otherwise searched substrings `charges`,
`duration` and `continuous` select 4, 1 and 2; their count is stored as `count<<4` in the high u16.
Unknown suffix text is mode 0. Kind 41 always treats its suffix as power. General decimal parsing
keeps a leading run from `+-1234567890`: empty is zero, `12junk` is 12, and a non-empty prefix for
which `%d` makes no assignment leaves an indeterminate value. Out-of-range input remains accepted,
but its concrete CRT result is Unknown. Cast power admits no minus, so `:-20` is zero. Range bytes
wrap and the second is spread, not maximum (`ITEM-EFFMODE-073`).

One Effect is 0x48 bytes: Token base, u8 kind `+0x3c`, u8 mode `+0x3d`, operand union `+0x40`,
transient dword `+0x44`. Effect copy preserves Token state `+0x0c` and transient `+0x44`; equality
compares only kind, the full operand dword and mode. Item equality first compares item code. Both
stackable means equal immediately, exactly one stackable means unequal, and only two non-stackable
items compare their ordered lists in lockstep. A Potion remains stackable with effects because its
item kind is 3 (`ITEM-STACK-003`, whose universal enchantment-separator clause is retracted;
`ITEM-EFFOBJ-072`).

For freshly parsed equipment Effects, Token state `+0x0c` is zero. Equip walks the list and
dispatches each value with multiplier +1; unequip walks it with -1. Every kind arm then calls target
`vt+0x50` recompute. Kinds 1..37 target price, actor attributes, pools, combat, protection and skill
fields; 38..40 diagnose not implemented; 41 is a general no-op; 42 teaches a missing Spell; 43 and
49 share damage-base; 44..48 install elemental base/spread. The per-kind target, clamp and class gates are specified by `ITEM-EFFDISP-075`.
Token-state values 8, 12 and 17 select special lifecycle paths and are not grammar kinds.

### Effect keys and the fields they add to (`ITEM-EFFARM-146`, `ITEM-EFFKEY-147`, `ITEM-EFFSYM-148`, `ITEM-EFFSHIP-149`)

With operand `S = multiplier × operand` (the signed low word when the mode byte has bit 1, 2 or 4,
else the signed dword), an arm adds `S` to its field at the field's own width. "Modifier" is the
0x40-byte object at `actor+0xd4`; "live" is the actor's own field. Humanoid-gated keys write the
modifier on the Human families and the live field on the Unit family.

| kinds | field |
|---|---|
| 1 price | `actor+0x1c` i32 |
| 2 body, 4 reaction, 3 mind, 5 spirit | live `+0x84`, `+0x86`, `+0x88`, `+0x8a` (cap 100), modifier byte `+0xd4`, `+0xd5`, `+0xd6`, `+0xd7` (skipped for mode bit 8) |
| 6 health, 9 mana | live `+0x94`, `+0x9a`, capped at `+0x96`, `+0x9c` |
| 7 healthMax, 10 manaMax (mage only) | live `+0x94`, `+0x9a` and modifier `+0xdc`, `+0xe0` |
| 8 healthRegeneration, 11 manaRegeneration (mage only) | modifier `+0xde`, `+0xe2` |
| 12 toHit | modifier `+0xe6` |
| 13 damageMin, 43 damage, 49 damageBonus; 14 damageMax | modifier byte `+0xf4` (base); `+0xf5` (spread) |
| 15 defence, 16 absorbtion, 17 speed, 19 scanRange | modifier `+0xfe`, `+0x100`, `+0xd8`, `+0xe4` (humanoid), else live `+0xbe`, `+0xc0`, `+0x8c`, `+0xa4`; scanRange shifts `S` left 8 |
| 20..25 protection0, Fire, Water, Air, Earth, Astral | modifier word `+0x102`, `+0x104`, `+0x106`, `+0x108`, `+0x10a`, `+0x10c`, floor 0 |
| 26 fighterSkill0, 27..31 Blade..Shooting (fighter only) | modifier `+0xe8`, `+0xea`..`+0xf2` |
| 32 mageSkill0, 33..37 Fire..Astral (mage only) | modifier `+0xe8`, `+0xea`..`+0xf2` |
| 18 rotationSpeed | mover `+0x0a` byte |
| 44..48 damageFire..damageAstral | assign modifier `+0xf9`, `+0xfa` from the operand bytes and `+0xfb` = 1..5 |
| 0, 38..41; 42 | no state change; teachSpell adds a missing Spell |

No instruction in an arm body stores into or points at the modifier damage-kind bytes
`+0x10e..+0x113`; each arm tail calls `vt+0x50`, which recomputes the live bytes `+0xce..+0xd3`
from the modifier. That no effect apply or remove changes the modifier bytes, callees included,
is Medium (`HERO-MODDK-161`). Remove is the exact inverse of apply except at the caps and floors above, for kinds
2..5 with mode bit 8, for kinds 44..48 (the multiplier is ignored, so remove restores nothing) and
for kind 42 (remove passes the negated id). A signed modifier speed above `0x18` is zeroed before
every dispatch. The shipped equipment cells carry 266 accepted effects per root over 16 keys, equal
in EN and RU, none of them belonging to an arm body that stores into the damage-kind block. Whether a
non-humanoid actor folds the modifier protections into its live block is Unknown.


### A weapon's numbers — `R0856`

Every scaled value is `ftol(row.column × shape.f × material.f + 0.5)`, where `shape` is the
5-row quality table at `L04635` subscripted by `item+0x45` and `material` the 16-row table at
`L02110` subscripted by `item+0x46` (the bound check at `L04637` limits `+0x46` to 15, and
`R0956`, which writes `+0x45`, searches `this + 0x14`).

**Double-factor layout.** The ladder in `ITEM-SCALE-017` and damage column in
`ITEM-DMGCOL-018` are retracted; use `ITEM-LADDER-019`…`ITEM-WEAPCOL-021`. Both tables share one schema (`Data.bin` group A: a name plus nine f64) and the
runtime record is **`0x68` bytes** (`R0959` multiplies the record count by `0x68` at `L04636`). The
serialized tail is `0x48`, so the head is `0x20` and double *j* sits at **`+0x20 + 8j`**. Group A
ships eleven titles for one name and nine doubles, so one title has no slot: the correspondence
is *j* ↔ shipped title *j+1*, leaving `Level` unserialized and putting the two identically-zero
doubles on the non-numeric `Abbreviation` and `Materials` columns.

```
disp    +0x20        +0x28      +0x30   +0x38   +0x40     +0x48    +0x50      +0x58         +0x60
title   Abbreviation Materials  Price   weight  @.damage  @.toHit  #.defence  #.absorption  MagCap
```

```
w+0x4a  weight  = col3 × shape[+0x38] × material[+0x38]    title: weight
w+0x60  dmgBase = col6 × shape[+0x40] × material[+0x40]    title: @.damage   <- @.physicalMin
w+0x61  dmgSprd = col7 × shape[+0x40] × material[+0x40] − w+0x60               <- @.physicalMax
w+0x52  toHit   = col8 × shape[+0x48] × material[+0x48]    title: @.toHit
w+0x6a  defence = col9 × shape[+0x50] × material[+0x50]    title: #.defence
w+0x48          =        shape[+0x60] × material[+0x60]    title: MagCap  (no column, no +0.5)
w+0x50  reach   = col0xb verbatim, or 1 when that column is −1 (no factor)
```

The armor and shield fills read factor `+0x58` (`#.absorption`); the
weapon fill does not. Price uses factor `+0x30`. — ITEM-LADDER-019,
ITEM-ARMFILL-032, SHOP-PRICE-011

**`w+0x60` is re-read as an already-rounded unsigned byte before the subtraction**
(the byte at `w+0x60` is loaded at `L04003` and the floating-point subtraction is at `L04004`), so the spread is
`round(max×f) − round(min×f)` and not `round((max−min)×f)`. The two encodings are
indistinguishable while one weapon is equipped and diverge as soon as a second source contributes,
because the actor's fold sums bases and spreads **separately** (`HERO-FOLD-035`).

`Common`'s `@.damage` factor is `0.2000`; omitting a leading shape word
selects this factor. For example: `Iron Short Sword` = shape
`Common` (0.2) × material `Iron` (1.0) over columns `23..40` gives **`(5, 3)`**, `@.attackType` 1,
`+0x52` 5, price 250. `Uncommon Steel Two Handed Sword` = `0.24 × 1.3351 = 0.3204` over `41..82`
gives **`(13, 13)`**.

Weapon's name constructor `R0665` separates an effect tail, then
parses a leading shape word, a leading material word and a remainder looked
up in `Weapons` through `R0666(L04593, …)`. Runtime collection
index 0 is rejected. `BareHands` is the first serialized Weapons definition;
that ordinal is distinct from the reserved runtime index. It can be named
by a Units `EquipItem` cell and is not the unarmed-Human default.
— ITEM-NAMEPARSE-040, ITEM-DMGCOL-018 (damage-column clause retracted;
use ITEM-WEAPCOL-021), HERO-BARE-037

### The serialized Weapon attack tail

W52 is the 24-byte block at Weapon `+0x52`. Its bytes `[19:22]` are Weapon
`+0x65`, `+0x66` and `+0x67`. The name and direct-code Weapon constructors
unconditionally initialize all three to zero through the embedded attack
initializer, before shape/material/row selection and fill. This is explicit
initialization and does not depend on new allocation contents. — ITEM-157

The named fill, price, text-Effect attachment, Equip and removal bodies have
no direct local store to these three Weapon bytes on valid, nonaliased paths.
The melee Equip branch reads them independently and assigns actor
`+0xf9/+0xfa/+0xfb`; removal clears the actor bytes. Transitive callback
preservation remains Unknown. — ITEM-158, ITEM-WEAPCOL-021

Weapon copy initializes the block and then copies 46 current source bytes
from `+0x52`, carrying any source tail over the zeros. SAVE exports the current
24-byte W52; LOAD restores the same raw block. Constructor defaults therefore
do not constrain a copied or loaded Weapon's tail. Other writers, callbacks
and the complete construction-to-first-SAVE lifetime remain Unknown. — ITEM-159

### An armour's and a shield's numbers — `R0890`, `R0888` (`ITEM-ARMFILL-032`)

The same routine twice, offset by two bytes because the armour spends `+0x50` on its `Slot`. The
three item fills share the `Data.bin` group C title list, so a column index means the same thing
in all three, and each factor's own group A title names the same attribute:

```
armor+0x50  slot       = col4 verbatim (no factor)              title: Slot
armor+0x52  defence    = col9  x shape[+0x50] x material[+0x50]  title: #.deIrnce  x #.defence
armor+0x54  absorption = col10 x shape[+0x58] x material[+0x58]  title: #.absorbtion x #.absorption
armor+0x48             =         shape[+0x60] x material[+0x60]  title: MagCap (no column)
armor+0x4a  weight     = col3  x shape[+0x38] x material[+0x38]  title: weight

shield+0x50 defence    = col9   ... same factors, same order ...
shield+0x52 absorption = col10  ...
shield+0x48            =        ...
shield+0x4a weight     = col3   ...
```

**The absorption term is the one scaled term in any of the three fills with no `+0.5`.** Every
other scaled value in this file is `ftol(x + 0.5)`; `armor+0x54` and `shield+0x52` are `ftol(x)`
(`L13001` → `L13002`, `L13003` → `L13004`, against `L13005` and `L13006`). The shield fill does not read the `Slot` column. Installed `Shields`
definitions carry `Slot = 2`; that value does not participate in the fill.

### What equipping an armour or a shield does (`ITEM-ARMFOLD-033`, event-order clause retracted)

Both classes embed one **0x16-byte modifier block**, `memset` to zero by
`R0977` → `R0839`, at `armor+0x52` and `shield+0x50`. The two scaled words above
are its members `+0x00` and `+0x02`; **nothing writes the other twelve.** `Armor::Equip` and
`Shield::Equip` pass the block's **address** to `R0845` twice — once with `actor+0xfe`,
once with `actor+0xbe` — and that routine folds fourteen members (`HERO-FOLD-033`). So:

```
an armour or shield contributes  defence (block +0x00) and absorption (block +0x02)
and contributes exactly zero to  the six protections  actor+0xc2..+0xcc
                                 the six damage kinds actor+0xce..+0xd3
```

There is **no per-slot difference**: the slot byte indexes `actor+0x198 + 4i` and is read nowhere
else, and neither block add sits under a branch on it. The Unit equip wrapper's
`vt+0x54` call (`L13007`) is **not** an equipment recompute. Humanoid/Human
equip/removal wrappers omit that call (`SAV-EQUIPCALL-554`; `ITEM-EQUIP-006`,
whose universal wrapper and recompute clauses are retracted).
The humanoid `+54` target `R0978` sums
`actor+0x1cc + 4i` for `i = 1..5`, the skill-experience caches, and stores `ftol(sum x 0.01)` into
`actor+0x1c`; the base class's slot is a bare `return actor+0x1c`, and the caller discards the
result. Equipment is **stored on the actor, never recomputed from the worn array.**

The helper set does not specify a common event order. Armor removal refreshes
negative weight, subtracts both defensive blocks, clears its slot, sets flags,
then removes Effects. Shield removal removes Effects after those block
subtractions but before flags and slot clear. Shield attach sets flags after
Effects; Armor attach sets them before (`SAV-EQUIPORDER-552`; `ITEM-ARMFOLD-033`,
whose common event-order clause is retracted and replaced by this order).
Weapon attach prepares its Spell before old-weapon eviction and calls derive
before timing/range and weight stores; Weapon removal starts with Effects and
clears the owned Spell and actor slot near the end (`HERO-EQUIP-017`).

Weapon's serialized byte `+50` is the current range operand, not the active
skill selector. After derive, attach adds `u8(Weapon+50)-1` to actor range
byte `+12c`; removal subtracts it, both with byte wrapping. Removal timing
is a literal 8/4 assignment. The active melee selector instead uses the low
byte of cached definition parameter 5; ranged 11/12 and removal assign zero
(`SAV-EQUIPORDER-552`, `SAV-HUMEQUIP-447`). Callback mutation remains Unknown.

Both Effect walks are forward. Each state 0 Effect's general normal-return
dispatch derives before the next Effect. One next-node pointer is prefetched;
later node contents remain live, so this is not a pre-collected batch
(`SAV-EQUIPEFFECT-553`). Command22 takes the source object, calls the actor
wrapper, reinserts the displaced/removed object and refreshes zero weight.
Derive from that final refresh is conditional, not an unconditional final pass
(`SAV-EQUIPCALL-554`). Named reads can occur before later local stores; a
direct-body no-store result does not establish transitive callback purity or
an atomic equipment update (`SAV-EQUIPOBS-555`).

### The `Weapons` row's columns

Group C ships eighteen titles for a name and seventeen slots, so `slot i` = shipped title `i+1`
with no spare. The consumed columns are:

- slot 2, `Price`: read by `R0962` into `item+0x1c`;
- slot 3, `weight`: the fill's `w+0x4a`;
- slot 5, `@.attackType`: read by `R0850` (`L04005`), then branched on at 10, 11 and 12;
- slot 6, `@.physicalMin`: the fill's `w+0x60`;
- slot 7, `@.physicalMax`: the fill's `w+0x61`;
- slot 8, `@.toHit`: the fill's `w+0x52`;
- slot 9, `#.deIrnce`: the fill's `w+0x6a`;
- slot 0xb, `@.range`: the fill's `w+0x50`; `Weapon::Equip` at `L04638` adds `w+0x50 − 1` to `actor+0x12c`;
- slot 0xc, `@.charge`: `Weapon::Equip` at `L04639` stores it to `actor+0x134`;
- slot 0xd, `@.relax`: `Weapon::Equip` at `L04640` stores it to `actor+0x135`;
- slot 0xe, `2 handed`: `Weapon::Equip` at `L04641` compares it with 2; this is the shield-removal gate.

`Weapon::Equip` has **three** arms on slot 5. `< 0xa` is melee: it adds `w+0x60`/`w+0x61` into
the modifier pair `actor+0xf4`/`+0xf5`, adds `w+0x6a` into `+0xfe` and `w+0x52` into `+0xe6`,
**assigns** `actor+0xf9/+0xfa/+0xfb` from `w+0x65/+0x66/+0x67`, and sets `actor+0xb6` to the
attack type. `== 0xb` and `== 0xc` add `w+0x60`/`w+0x61` into `actor+0xf9`/`+0xfa` instead and
assign the kind byte `+0xfb` to 1 or 2, leaving `actor+0xb6` zero — so the shipped
`Flame Thrower` feeds the *second* damage component, not the first.

### A weapon's own damage line

The game draws it, and it reads the weapon's own two bytes. `R0808` takes the equipped
weapon at `actor+0x74` and appends two attributes through `R0917`, which stores a
`(tag, value)` byte pair at `this+0x0c + 2 × this+0x09`:

- At `L04644` the weapon's byte at `+0x61` (the spread) is read as the gate: a weapon with zero spread draws no damage line.
- At `L04642` the byte at `+0x60` becomes attribute tag `0x0d`.
- At `L04643` the byte at `+0x61` becomes attribute tag `0x0e`.

`R0964` formats the list — the index is `tag − 1`, bounded above by 0x32, and goes through the
byte table at `L04646` into the jump table at `L04647`. The `0x0d` arm at `L04648` consumes
**both** pairs: it prints the base, steps over the next tag byte, adds the spread to the base (at `L04650`) and
prints `base + spread`. So the item's line is `[base, base+spread]` — the same composition the
character sheet uses (`HERO-SHEET-038`) — and tag `0x0e`'s own table entry is the default arm,
which nothing reaches. The shipped `Iron Short Sword` draws `5-8`.

### Equal effect kinds on one item (`ITEM-154`, `ITEM-155`, `ITEM-156`)

The item information formatter appends one line per stored attribute pair into its output buffer
and compares no pair with another (`ITEM-154`). An Armor or Shield packer writes its base defence
as tag 15 and, when above zero, its base absorption as tag 16 before the Effect pairs, and an
Effect pair carries its kind as its tag, so a Defence or absorbtion Effect repeats a tag
(`ITEM-155`). Among the direct-call appends enumerated, only the Magic Items shop generator
merges an equal kind: it overwrites the operand dword of the existing node and discards the new
Effect. The map loader, the cell parser, the item copy and the book makers append with no kind
test (`ITEM-155`).

Each Effect node adds its own operand to the modifier word, so two Defence nodes of +1 add 2
there; the recompute into displayed defence is not established (`ITEM-156`). The campaign
equipment cells and the linked campaign map recipes repeat no kind; the loose-map records' item
links are not established (`ITEM-156`).
