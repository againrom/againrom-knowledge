# Skills and experience

[Reference](format.md)

## Skill slots

`actor+0xa8 + 2i`, `i = 0..5`. Slot 0 is `Skill.General`; slots 1..5 are one shared set that the class
renames, `Data.bin`'s own column titles giving both names at once:

| slot | fighter | mage |
|---|---|---|
| 1 | Blade | Fire |
| 2 | Axe | Water |
| 3 | Bludgeon | Air |
| 4 | Pike | Earth |
| 5 | Shooting | Astral |

Base copies live at `+0x116 + 2i` (`i = 1..5`) and bonuses at `+0xe8 + 2i`. There is no
sixth..tenth slot: `+0xb4` immediately follows and is written as the damage byte.

The word at `+0x116` is inside the serialized block and both school price arms index it when their
slot is 0, but it is **not** a General base copy: all twelve restore, snapshot and bonus-fold loops
start at 1 and end at 5. Five are in the play-award routine, two each in purchase and derive, and
one each in load, creation and death loss. A General purchase changes live `+0xa8` and `xp[0]`
without refreshing `+0x116`.

Slot 0 is the ranged-weapon accuracy skill. A ranged weapon copies `skill[0]` to the to-hit
modifier and sets active skill to 0. Melee weapons select one of slots 1..5 instead. General
experience also enters the six-slot total used by the health and mana maximum formulas. It is not
included in the five-slot victim XP-value sum.

The spell Sphere is a raw index into this vector on three readers and one writer. The writer's
accessor `R0915` takes the Spell row at `+0x04`, pushes parameter index 2, calls the row
accessor and returns byte 0; the paired `Spells` title is `Sphere`. Its caller masks the result to a
byte before indexing the vector. Both shipped roots have 0 Sphere-0 rows among 28 spells; custom
Sphere 0 therefore activates an existing General seam without changing the spell schema.

The bonus slots are written by **effects only**, and by two class-gated sets of six arms over the
same six fields: `fighterSkill0`/`skillBlade…skillShooting` behind the fighter predicate,
`mageSkill0`/`skillFire…skillAstral` behind the mage one. An effect naming a fighter skill does
nothing to a mage and vice versa. Equipment reaches them only through an item's effect list.
For a recomputing human, the two slot-0 effect kinds are inert: they write bonus `+0xe8`, while the
derive restores and folds bonuses only for indices 1..5.

## Experience and skill levels

Two fields per slot, and they are **not** two views of one number: the level `+0xa8 + 2i` and the
experience `+0x1cc + 4i`. `actor+0x130` is their running sum.

```
S(n)   = ftol( (pow(1.1, n) - 1) * 1000 )       what a slot at level n accounts for
S^-1(x)= ftol( log_1.1( x/1000 + 1 ) )          used by the loss ONLY
invariant  actor+0x130 == sum over i = 0..5 of xp[i]      preserved by all four writers
at creation  xp[i] = S(skill[i])                          R0899, both build paths
```

### the raise — `R0810`, `vt+0x5c` of both human classes

```
gate    typeID (+0x0e) in [0x21, 0x3f]             player-character mode, not every Human
        source, if given and not already dead (health >= 0):
          same owner (+0x14)                        -> nothing
          diploMatrix[srcPlayer][myPlayer] & 2       -> nothing
amount  ftol( amount * (mind/30 + 0.25) )           signed; zero is not an early refusal
slot    +0x4c & 4 ? the caller's slot, which must be > 0
                  : the caller's slot must be 0, and the slot used is +0xb6 (active), > 0
        skill[slot] must be < 100, tested BEFORE the award
cap     amount = min(amount, S(level+1) - S(level))         negative amounts remain negative
                                                          [L00285]+0x0c ? that / 5 : that
pay     xp[slot] += amount ; actor+0x130 += amount          can reduce XP; no level-decrement arm
raise   if xp[slot] > S(skill[slot]):  skill[slot] += 1     ONE level, never two
after   a raise re-snapshots base[1..5] and calls vt+0x50 (the whole derive);
        on the +0x4c & 4 arm it also re-powers every spell in the book,
        power = clamp(skill[Sphere] + Mind - 30, 0, 100)
no raise -> skill[i] = min(base[i] + bonus[i], 100) for i = 1..5
```

Three feeds, all `vt` slots of the same two vtables and all **stubs** on the non-human one:
`vt+0x60` a kill, `ftol(victim+0x1c * 0.5)`; `vt+0x64` a landed hit,
`ftol(victim+0x1c * 0.5 * dealt / victim.healthMax + 1)`; `vt+0x68` a cast,
`round(manaCost/2)`. `victim+0x1c` is the `Units` `XPvalue` column. Each resolves the slot as
*"the spell's `Sphere` if a spell id was supplied and `+0x4c & 4` is set, otherwise 0"*. The
clear-bit arm accepts that zero and substitutes the current weapon skill at `+0xb6`; a fighter's
weapon-borne spell damage therefore trains its weapon, not a school, only after the typeID gate
above. A low-type map Human receives no item hit, damage or kill award. Item context suppresses the
`vt+0x68` cast feed, but later direct damage, Drain Life, Slow, Stone Curse, Curse and Poison Cloud
can each reach `vt+0x64`. On the fighter route those spell-side events run before the physical XP
test. That later test rereads post-spell health and refuses at `-10` or below, so the physical blow
does not necessarily add its own award. Their cadence follows affected targets and ticks.

The damage feed first requires a victim owner with `+0x28 != 0`; kill has no `+0x28` test. Both
require a victim owner and refuse when owner `+0x5c` and multiplayer are both nonzero. Poison ticks
reject only a zero signed amount and retain a recorded caster at health zero, clearing it only below
zero. A custom negative item power can therefore turn Poison into a healing tick that still reduces
the credited Humanoid's slot and aggregate XP. The shipped Catapult and Ballista carry castSpell
weapons, but use the non-human vtable's three stubs: they can release the rider and never train.

The kill feed's second argument is the final victim attribution byte. The resolver retains prior
state for null source/definition, clears the actor for an ownerless defined source, and otherwise
writes source plus damage kind for bit 4 or zero for clear bit. A non-Defensive PointEffect can
replace it with actual id after payload; an AreaEffect has a distinct writer gated by a nonzero,
zero-extended target movement-domain byte.
Lethal payload damage does not change either condition for ordinary shipped direct-damage shapes.
Drain Life constructs neither envelope and writes no attribution. Poison Cloud can seed id 8 on
area application but does not refresh attribution on later ticks, leaving both eventual kill paths
dependent on the victim's runtime history.

Consequences a consumer must carry:

- **No play award raises `Skill.General`.** Both arms exclude slot 0, and every restore /
  re-snapshot loop runs `i = 1..5`. Opcode 61 can raise it when supplied slot 0, although the
  shipped school cannot supply that selection.
- **A `+0x4c & 4` carrier earns nothing from melee or missile combat** — those arrive with slot 0
  and its arm refuses that.
- **The first award into a slot always raises it**, because creation leaves `xp[i] == S(skill[i])`
  and the test is a strict `>`.
- **A slot at 100 accrues nothing at all** — the ceiling test precedes the payment.
- **Item-spell progress uses these same six slots and fields.** There is no item proficiency or
  separate save state.

### the loss — `R0825`, charged by conditional repair

For `i = 0..5` (**General included**): `xp[i] = xp[i] * 9 / 10`, then `skill[i] = S^-1(xp[i])`; the
total is rebuilt and `base[1..5]` re-snapshotted. Its one caller `R0823` runs it only for
an existing primary with **zero-extended byte** `actor+0x13c != 0`, then restores HP and MP.
The loss tail clears the stage. This is not a campaign death or reload charge.

There are three caller sites of that accepter (`HERO-DEFEAT-136`):

| Entry | Admission to repair |
|---|---|
| Reporter | Actual `server+0x0c != 0`, participant active, latch at least 2, entry-active byte nonzero, signed primary HP below -53; campaign paths are excluded |
| Session opcode `0x48` | Resolved player; no mode/latch restriction in this handler, so an existing staged primary can be charged if the command is delivered |
| Session opcode `0xbe` | The call occurs only with a null primary; it creates rather than charging an existing dead primary |

Observed `0x48` production is character acceptance, not a failure-panel command. Teardown,
direct healing and imported-character restoration also clear stage without this penalty.
Ordinary UI delivery of a staged-primary `0x48` remains Unknown. The loss arithmetic's prior
level-band result remains in `HERO-SKILLLOSS-075` (amended); it does not establish mode reachability.

### the purchase — `R0838`, one command

Server command opcode **61** and no other. Requires `server+0x0c == 0`; slot from the packet,
bounded `0..5`, so it is the only shipped server operation that reaches `Skill.General`; price
`ftol(1.1^word[actor+0x116+2*slot] * 200)` charged against `player+0x38`. For slots 1..5 the word
is the maintained base; for General it is the decoupled serialized word above. Then
`skill[slot] += 1` **with no ceiling**, `xp[slot] = S(new level) + 1`, base slots 1..5
re-snapshotted, derive called.

The first crossing of each fixed-width stage is below. Order is by the numeric level/base input;
live General and its serialized price shadow are independent inputs, so this is not one execution
timeline.

| order | input transition | wide calculation | stored/consumed value | Width boundary |
|---:|---|---|---|---|
| 1 | new level 152 → 153 | `S(152)=1,957,437,581`; `S(153)=2,153,181,439` | purchase stores low-dword `S(153)+1 = 0x8056f100`, signed `-2,141,785,856` | signed i32 consumption of the stored dword |
| 2 | new level 160 → 161 | stored `S(n)+1`: `4,195,942,440` → `4,615,536,784` | low dword wraps to `0x131b8090`, signed `320,569,488` | low-u32 experience storage |
| 3 | maintained price base 169 → 170 | `P(169)=1,978,763,028`; `P(170)=2,176,639,331` | low dword `0x81bce163`, signed `-2,118,327,965`, is compared to the purse with signed `JLE` | signed i32 price comparison |
| 4 | maintained price base 177 → 178 | `P(177)=4,241,654,286`; `P(178)=4,665,819,714` | low dword wraps to `0x161ac242`, signed `370,852,418` | low-u32 price storage |
| 5 | new level 385 → 386 | `S(n)`: `8.6334382265275638e18` → `9.4967820491803197e18` | the 64-bit integer-store input first exceeds signed-qword range | signed-qword conversion; native out-of-range result Unknown |
| 6 | maintained price base 402 → 403 | `P(n)`: `8.7274913946611507e18` → `9.6002405341272678e18` | the 64-bit integer-store input first exceeds signed-qword range | signed-qword conversion; native out-of-range result Unknown |
| 7 | new level `0x7fff` → `0x8000` | the stored word is 32768 | the sign-extending load passes `-32768` to `S` | signed i16 consumption of the stored word |
| 8 | old level `0xffff` + 1 | mathematical result 65536 | the 16-bit increment stores `0x0000` | u16 increment wraps to zero |

The total at `actor+0x130` is another signed dword:
because it accumulates all six slots, its sign and wrap thresholds depend on the other five values
and can precede the per-slot rows above. Between the first dword wrap and the qword-input boundary,
later sign and wrap transitions repeat modulo `2^32`; the table records the first crossing of each
width stage, not every repetition. Runtime behaviour downstream of every unexecuted
transition, including the exact CRT result after either qword-range crossing, remains Unknown.

**Store both fields.** Experience is not a display of the skill vector and the skill vector is not a
display of the experience: after any raise `xp[slot]` lies in `(S(n-1), S(n)]` for level `n`, and
only the loss ever brings them back into agreement.
