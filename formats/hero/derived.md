# Derived attributes and modifiers

[Reference](format.md)

## Derived-stat sequence, `R0280` (`vt+0x50` of both human classes)

Reached **only** through the `.rdata` slots `L03967` and `L03968`. The order is normative:

```
 0  stat caps    stat_i = min(stat_i, 50 + (int8)actor[0xd4 + i])            i = 0..3
 1  healthMax    h = body * (fighter ? 2 : 1)             -> +0x96 UNCONDITIONALLY (L04429),
                 [skip the two lines below if that product is 0 -- L03889 tests the
                  value just stored, NOT the Data.bin HealthMax column, which is gone]
                 h = ftol( h + log1.1(XP/5000 + 1) * (fighter ? 2 : 1) )
                 h = ftol( h * (pow(1.1, body)/100 + 1) )                    -> +0x96
 2  manaMax      [gate FIRST, on the streamed Data.bin ManaMax column at L04434:
                  if it is 0, current mana = 0 and the whole of 2 is skipped]
                 m = spirit * 2                                              -- overwrite second
                 m = ftol( m + log1.1(XP/5000 + 1) * (fighter ? 1 : 2) )
                 m = ftol( m * (pow(1.1, spirit)/100 + 1) )                  -> +0x9c
 3  sight        ftol( ((mind + reaction)/25 + 4) * 256 )                    -> +0xa4 u16, 1/256 cell
 4  capacity     body * 10 + 1                                               -> +0x92
 5  speed        reaction < 12 ? reaction : reaction/5 + 12
                 + 10 if typeID == 0x13 or 0x15                              -> +0x8c
 6  load         weight(+0x8e) + (money < 64000 ? money/2 : set load = 32000) -> +0x90
 7  overload     if load >= capacity: speed -= load/capacity ; speed = max(speed, 6)
 8  damage       dmgSpread = dmgBase = ftol( pow(1.1, body) / 20 )           -> +0xb5, +0xb4
                 the pair is (base, SPREAD): the roll is base + U[0, spread] -- combat.md: hit resolution
                 NOTE the Body term is added to BOTH bytes, so it raises the roll's
                 minimum once and its maximum twice; step 11's skill term hits ONLY
                 the base, and nothing on either path is multiplicative
 9  toHit        ftol( (pow(1.1, body) + pow(1.1, reaction)) / 5 )           -> +0xa6
10  skills       skill[i] = min(100, base[i] + bonus[i])                     i = 1..5
11  active skill toHit += 3 * skill[active] ; dmgBase += skill[active]/5     active = +0xb6, 0 = none
12  zero         memset(actor+0xbe, 0, 0x16)   -- defence, absorption, both resistance arrays
13  defence      reaction / 3                                                -> +0xbe
14  protections  prot[i] = spirit / 2                        i = 1..5, +0xc4 .. +0xcc (u16)
15  modifiers    R0840(actor+0xd4, actor)             -- modifier fold below
16  current      current health/mana clamped to their maxima
17  mover        [actor+0x154] + 0xa = low8(speed)            -- turning byte, see below
18  clamps       defence, absorption and load floored at 0;
                 prot[i] = clamp(min(spirit/2 + 70, prot[i]), 0, 100);
                 skill[i] = clamp(skill[i], 0, 100)
19  effects      if actor+0x140: walk the spell list
```

The mover assignment is an exact byte transfer:L12999 loads actor+154,
L03931 reads BYTE actor+8c and L03932 writes the allocated mover's byte+a.
This byte controls the reached facing update; it replaces any earlier table
rotation value when this derive reaches the store. The modifier fold's
negative-speed branch clears modifier+d8 while leaving the already stored
WORD+8c unchanged. Full callback effects remain Unknown — MOVE-RATE-053.
Effect selector 18 directly increments mover+a modulo256, then calls the same
target's vt+50. That intermediate write does not establish a surviving Human
bonus, because the Human/Humanoid slot selects this derive — MOVE-RATE-054.

Order is load-bearing at four points: the caps run first, so everything downstream sees the capped
value; health and mana are computed **before** the skills are restored, so a skill change reaches them
only on the next recompute; the `memset` runs between defence's inputs and defence's write, so all of
`+0xbe..+0xd3` restart from zero every time; and the equipment pass runs **between** the bare values
and the clamps, so the clamps bound equipment rather than stats.

**Current health `+0x94` and current mana `+0x9a` are live state, not derived.** Only the maxima are.

**`БРОНЯ` (absorption, `+0xc0`) is derived from no stat** — zeroed at step 12 and never rewritten by
the derive. It and `ЗАЩИТА` (defence, `+0xbe`) are two separate fields with two different sources,
and both are supplied at step 15.

### Step 0 in full — what bounds a stat *in principle*

```
stat_i = min(stat_i, 50 + (int8)actor[0xd4 + i])     i = 0..3, and the smaller is WRITTEN BACK
```

The `50` is a hardcoded immediate, four times over. It is **not** class-dependent, **not** per level
and **not** in any file. The modifier byte is signed and is written by **effects only**: an effect
that raises a stat raises the modifier by the same amount (unless its own flag bit 3 is set, in which
case the next recompute takes the raise back), and clamps **the stat alone to 100**.

```
a stat no effect has ever touched      capped at 50   (chargen can only reach 43)
the ceiling in principle               100            imposed by the effect arm, not by the cap
fighters vs mages                      identical      the class flag reaches neither operand
```

The `100` is the same compare against 0x64 in all four arms, and each arm owns one modifier byte:
`+0xd4` Body, `+0xd5` Reaction, `+0xd6` Mind, `+0xd7` Spirit. Each of the four was enumerated
image-wide and each returns the same shape — **4 hits, 2 owners, 0 in orphan code**: the derive
reads it twice, the effect dispatch reads and writes it. Because the cap is step 0 and every effect
arm ends in a full recompute, **whatever reads a stat sees the capped value**, and an untouched
stat is therefore 50, not its nominal maximum.

Consumer note: the modifier is a **byte** and is not clamped, so a long enough chain of raises wraps
it negative and the cap with it.

### When the derive runs on a map-placed person, and what its inputs are (`HERO-HP-071`, `HERO-HP-072`)

The same routine is the `vt+0x50` of both human vtables, and on the `.alm` placement path it runs
**twice** before the actor is playable:

- In `R0656`, the call at `L04443` runs `R0899`, which fills `actor+0x130` from the six skill columns; the virtual call at `L04444` (vtable offset `+0x50`) is the derive; then `L04437` copies `actor+0x96` into `actor+0x94`.
- In `R0151`, `L12688` to `L12693` apply the placement record's stat overrides; the virtual call at `L04446` (vtable offset `+0x50`) is the derive again, on the overridden Body; then `L12694` copies `rec+0x20` into `actor+0x94`, only when `rec+0x20` is not -1.

So for such an actor the **`Data.bin` `HealthMax` column is dead**: the streamer writes it into
`+0x96` (and `+0x94`), step 1 overwrites `+0x96` at `L04429`, and the surviving
`+0x94 = min(+0x94, +0x96)` clamp at `L03902` is itself overwritten by `L04437`. The inputs
that decide a placed person's health maximum are Body (capped at 50 by step 0, the block being
zeroed at spawn), the six skill columns through step 1's experience term, and whether the `ManaMax`
column is positive — which sets the class bit at `L04439` and therefore the `fighter ? 2 : 1`
multiplier. Over the 215 shipped `Humans` rows the derived maximum equals the column on **1** and
reaches **20.2x** the column on one row. The `-1` defaults for such an actor are **not** the base
constructor's: `R0876` runs after it and raises Mind and Spirit to 30, healthMax and health
to 50, and speed to 16.

## Modifier fold `R0840(actor+0xd4, actor)`

The actor carries **three parallel copies of one layout**, built together by every constructor:

| live | modifier | base | contents |
|---|---|---|---|
| `+0xa6` | `+0xe6` | `+0x114` | `[toHit u16][6 skill u16][dmgBase u8][dmgSpread u8][active u8][2 u8][elemBase u8][elemSpread u8][elemKind u8]` |
| `+0xbe` | `+0xfe` | — | `[defence u16][absorption u16][6 protection u16][6 damage-kind u8]` — 0x16 B |

The modifier block is `actor+0xd4 … +0x113`, exactly 0x40 bytes: four stat-cap bytes, then
`+0xd8` speed, `+0xda` capacity, `+0xdc` healthMax, `+0xde` healthRegen, `+0xe0` manaMax,
`+0xe2` manaRegen, `+0xe4` sight, then the two sub-blocks at `+0x12` and `+0x2a`. The fold is:

```
speed += +0xd8 ; capacity += +0xda ; healthMax += +0xdc ; manaMax += +0xe0 ; sight += +0xe4
if the resulting speed < 0:  the MODIFIER +0xd8 is zeroed, not the speed
+0xbe block += +0xfe block   (all 0x16 bytes)
+0xa6      += +0xe6 ; dmgBase += +0xf4 ; dmgSpread += +0xf5 ; four more bytes added, the last ASSIGNED
```

The six skill words are **not** in this fold — the derive adds `+0xe8 + 2i` itself at step 10.
The complete attack fold is 53 instructions with seven `ADD`s (one word and
six bytes) and one byte assignment. It also skips active-index mirror `+0xf6`
and final bytes `+0xfc/+0xfd`; the active live index is written separately.
— HERO-FOLD-035 (amended), SAV-HUMFOLD-446

No shipped Armor or Shield block and no effect-arm body writes the modifier's six damage-kind
bytes `+0x10e..+0x113`: the shipped blocks add zero and no arm body stores there. A sweep of the
image for displaced stores finds one routine that stores a constant, `R1562`, which sets the six protection words and
six damage-kind bytes to 100 and then recomputes. Its callers are the `#modify … +god` order and
two spawn paths; the constructors clear the object and the archive copies it whole
(`HERO-MODDK-161`).

The fold is not a universal loaded-state boundary. Before its later clears,
derive reads `word[a8+2*b6]` for any nonzero byte `b6`; the selected indexing
step has no upper bound. Malformed synthetic indices 10,13,32,39,42 make
`bc/bd`, slot-zero `c2/c3`, General shadow `e8/e9`, mirror/secondary pair
`f6/f7`, and unnamed `fc/fd` affect to-hit/damage. Ordinary producer reach of
those indices is Unknown. The damage resolver also has an unchecked
`target+ce+attacker.b6` byte index. — SAV-HUMINDEX-464

The actor-state sender computes wrapped sums of live physical/secondary/
elemental damage bytes with an effective damage mask, without a local derive
call. This is presentation, not the combat formula. Phase-12 list order puts
that actor's projection before its regeneration read of `de/e2`; ordinary
subticks walk Effects before orders/actions. Generic Effect dispatch reads
`d8` before its own later derive, while special Effect identity 17 on a Human
changes `e4/a4` by signed magnitude times 256 without generic derive.
The admitted strike independently consumes live absorption, resistance,
secondary protection and elemental selector under distinct component gates.
— SAV-HUMPROJECT-461, SAV-HUMTICK-462, SAV-HUMSTRIKE-463

None of these conditional routes establishes the absolute first read after
loading: earlier commands, callbacks and other receivers may already have
changed the actor. Nonzero capacity/secondary-modifier producers and ordinary
residual-field semantics remain Unknown; source-byte passthrough and the
temporary chargen Human's derive do not close those authoring inputs.
— SAV-HUMFIRST-465

Fresh hero and hired Human creation share one Human constructor chain. Its
embedded block calls leave live-tail `bc/bd` unchanged and clear modifier-tail
`fc/fd`. A 488-byte Human allocation rounds to 496, above the image's480-byte
small-block threshold, and reaches imported `HeapAlloc` with flags0; returned
payload contents are not established. Both pairs reach raw SAV writes, but
the constructor clear does not close later Effect/indexed writes or first-save
values. — SAV-HUMALLOC-504, SAV-HUMNEW-505, SAV-HUMNEWSAVE-507

The 156 shipped nonempty Human weapon-definition cells per root all join to
attackType1..5. Weapon equip copies the low byte only for signed type below 10;
every type at least 10 and removal set active selector 0. Thus positive type 10/42
does not yield the known tail aliases. This finite producer bound excludes
neither negative/custom values nor archive, starting-skill and later lifecycle
aliases. It does not establish a default vector. — SAV-HUMSEL-506,
SAV-HUMNEWSAVE-507

The mode-gated hero helper writes six defensive modifier words at `102..10c`
and six resistance bytes at `10e..113`, all 100, then calls Human `+50`.
Its own writes preserve the selected eleven residual bytes up to that call;
the subsequent derive and elapsed mission lifetime are not a proved
preservation interval. The gate tests nonzero server `+12c` or `+134`;
their ordinary campaign settings remain Unknown. — SAV-898

Starting-skill input is separate from active selector `b6`. Its local indexed
word store can overlap all selected residual spans at indices
10/25/32/39/40/42. The measured packet arm forwards the skill byte to creation,
but that conditional transport does not prove an ordinary out-of-range source
or a reached active-index consumer. — SAV-899

The earlier constructor clears the nine selected modifier bytes. That
clear does not specify allocation contents before it, survival of the live
tail through later calls, or safe first-save values. Ordinary residual use
and the first selected-field reads/writes after normal load remain Unknown.
— SAV-900, SAV-HUMRUNTIME-476

## Modifier producers

**Equipment** has class-specific stores. Armor/Shield add into the modifier and
live defensive copies. Weapon writes modifiers and calls derive:

```
Armor   slot = item+0x50, stored at actor+0x198 + 4*slot ; block item+0x52 -> +0xfe and +0xbe
Shield  single slot actor+0x78                           ; block item+0x50 -> +0xfe and +0xbe
        equipping one drops a two-handed weapon
Weapon  single slot actor+0x74, melee (attackType < 10):
          dmgBase mod += w+0x60 ; dmgSpread mod += w+0x61 ; defence mod += w+0x6a ;
          toHit mod += w+0x52  ; active skill = low byte of cached definition parameter5
        ranged (attackType 0xb / 0xc):
          the damage goes to +0xf9/+0xfa, +0xfb = 1 or 2, toHit mod = skill[0], active skill = 0
unequip subtracts additive terms, but ranged toHit assignment is not inverted:
          removal subtracts w+0x52 from the General value assigned to +0xe6
```

The universal inverse rule does not hold: prior to-hit modifier 5, General 17
and item to-hit 11 give melee 16 then 5, but ranged 17 then 6 in the bounded
original instruction slices. Weapon directly writes modifiers and invokes
derive; Armor/Shield explicitly add into both defensive copies. Full runtime
cycles and transitive callback interleaving remain Unknown. The bounded local
event sequence is specified below.
— HERO-EQUIP-017 (amended), SAV-HUMEQUIP-447

### Local equipment sequence

Armor removal refreshes negative weight, subtracts defence, clears its slot and
sets flags, then removes Effects. Shield removal refreshes and subtracts first
too, but removes Effects before flag update and slot clear. Weapon removes
Effects first, then changes attack modifiers/active skill and derives;
range/timing, weight, flags, owned-Spell deletion and slot clear follow. On
attach, Weapon prepares its Spell before eviction and derives before
timing/range and weight. Armor/Shield store their slot and add defence before
weight and Effects; Armor sets flags before Effects, Shield after. Same-slot
eviction calls the item; opposite-hand two-handed eviction calls the actor
wrapper and reinserts before installing the new item (`SAV-EQUIPORDER-552`,
`HERO-EQUIP-017`, `ITEM-ARMFOLD-033`, whose common event-order clause is
retracted; its slot store and block add stand).

The active selector comes from definition parameter 5 cached after displacement
and the new-weapon slot store; melee copies its low byte, ranged 11/12 and
removal assign zero (`SAV-HUMEQUIP-447`). Range uses a different source:
after derive, attach/removal reread current Weapon byte `+50` and add/subtract
its minus-one delta in actor byte `+12c`. Removal assigns timing 8/4. This is
not a guaranteed inverse cycle across callbacks (`SAV-EQUIPORDER-552`).

State 0 Effects run in forward list order, on removal as well as attach. Each
normal-return general Effect dispatch reaches actor `+50` before the next item
Effect; an empty list does not. The iterator prefetches one next-node pointer,
not all remaining Effect values. Actual callback list mutations remain Unknown
(`SAV-EQUIPEFFECT-553`). Command22 adds container removal/reinsertion and a final
zero-weight refresh; that refresh derives only on a changed load quotient.
Humanoid/Human wrappers add no trailing derive; Unit's trailing `+54` is Token
value, not `+50` derive (`SAV-EQUIPCALL-554`; `ITEM-EQUIP-006`, whose universal
wrapper and recompute clauses are retracted). Derive and Effect reads therefore
occur between local stores; no atomic final-state or pure callback contract
follows (`SAV-EQUIPOBS-555`).

A weapon's numbers are `round(column × shapeFactor × materialFactor)` (`+0.5` then `ftol`); runtime
column *i* is `Data.bin` title *i+1*, so they are the shipped titles `@.physicalMin`,
`@.physicalMax`, `@.toHit`, `#.deIrnce`, `weight`. The ladder clause of
`ITEM-SCALE-017` is retracted; `ITEM-LADDER-019` supplies the factor layout:
read one f64 from the Shapes record indexed by `item+0x45` and the Materials
record indexed by `item+0x46`, at `record + 0x20 + 8×(title−1)` — so the damage factor is the column the game itself
names `@.damage`, the to-hit factor `@.toHit`, the defence factor `#.defence`. **Both ends of the
damage pair take the same `@.damage` factor**, so an item scales symmetrically; the actor does not. **The damage pair is not stored as it is
written**: `w+0x61 = round(physicalMax × s × m) − w+0x60`, i.e. the spread. On the defensive block
the first word is `#.deIrnce` (defence) and the second `#.absorbtion` (`БРОНЯ`).

**Effects**, one per typed increment. `R0843` takes `kind = effect+0x3c` (`0..0x31`) and
`v = param × effect+0x40`, jumps through a 50-entry table, and ends every arm with a full recompute.
The kinds are `Data.bin`'s Magic rows by index and by name: `price`; the four stats (each with its
cap byte); `health`/`healthMax`/`healthRegeneration`; `mana`/`manaMax`/`manaRegeneration` (the last
two mage-only); `toHit`; `damageMin`/`damageMax`; `defence`; `absorbtion`; `speed`; `rotationSpeed`;
`scanRange`; six protections; twelve class-gated skill slots; three lores; `castSpell`/`teachSpell`;
six damage elements; `damageBonus`. Fifteen arms branch on `target->vt+0x30()` — modifier copy when
the target recomputes, live field when it does not. `itemLore`, `magicLore`, `creatureLore` and
`castSpell` are **dead arms**.
