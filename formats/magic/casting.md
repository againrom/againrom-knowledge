# Spell objects and casting

[Reference](format.md)

## Arithmetic

`ftol` truncates toward zero (`R0279`, rounding-control 11); `IDIV` truncates toward zero.
`rand(n)` is `R0861`: `rand() * (n+1) >> 15`, uniform on `[0, n]` **inclusive**, 0 when
`n == 0`. Runtime `Data.bin` column *i* is title *i+1* throughout.

## Spell definitions

Direct-damage effect delivery to a Building differs from a physical actor action:
`R0564` selects the Building resolver, subtracts from word `+0x42`, clamps a
negative result to zero and notifies only for positive damage. The physical Building
strike does not clamp and has a different notification predicate. The shared resolver
does not make the two delivery paths interchangeable (`UNIT-STRUCTDELIVERY-065`). This
does not establish every spell's Building admission or area-effect multiplicity.

`World\Data\Data.bin` group **H spells** ships **28** rows, indices 1..28. `rom.exe` carries the
same names as a 29-entry array at `L05034` whose entry 0 is `unused_spell_0`; the names agree
and the ID is the array index. Spell 0 is refused at every entrance.

```
 1 Fire Arrow            2 Fire Ball             3 Wall of Fire        4 Fire Sacrifice
 5 Protection from Fire  6 Heal                  7 Freezing Cloud      8 Poison Cloud
 9 Acid Stream          10 Protection from Water 11 Drain Life        12 Light
13 Lightning            14 Prismatic Spray      15 Invisibility      16 Protection from Air
17 Darkness             18 Shield               19 Wall of Earth     20 Stone Curse
21 Meteor Storm         22 Protection from Earth 23 Bless            24 Haste
25 Control Spirit       26 Teleport             27 Curse             28 Slow
```

The columns the cast path reads, by runtime index (title in quotes):

| idx | title | used for |
|---|---|---|
| 0 | `Complication Level` | added when cast state `0x0d`/`0x0e` still has a non-null `actor+0x64` Spell at recovery; weapon-diverted casts clear it first |
| 1 | `Mana Cost` | the whole cost, cached at `spell+0x0c` |
| 2 | `Sphere` | the school, 1..5 = Fire/Water/Air/Earth/Astral — also the damage kind |
| 5 | `Delivery System` | 1 = attach now, 2 = a flying missile |
| 6 | `Max Range` | cached at `spell+0x09`, then raised by the power |
| 7 | `Spell Effect Speed` | the missile's speed; the delay is `distance / speed` |
| 8 | `Distribution system` | 1 = a point effect on one unit, else an area effect |
| 9 | `Radius, Length/2` | the area effect's size |
| 11 | `Area Effect Duaration` | the fallback duration when `Spell Duration` is 0 |
| 14 | `Spell Duration` | the duration in game seconds, ×16 for ticks |
| 16, 17 | `damageMin`, `damageMax` | the damage pair, scaled by the power |
| 18 | `Defensive` | cached at `spell+0x0a` as `(value == 1)` |

## Spell objects and spellbooks

```
Spell        0x14 bytes, vtable L05044
  +0x00  vtable
  +0x04  pointer to the Data.bin Spells row   (restored from the id on load)
  +0x08  u8   spell id 1..28
  +0x09  u8   stored cast range -- initialized from Max Range; updated by R0904 -> R0625
  +0x0a  u8   Defensive
  +0x0c  u16  Mana Cost
  +0x0e  u8   per-cast scratch: damage base
  +0x0f  u8   per-cast scratch: damage spread
  +0x10  u16  per-cast scratch: duration times 16 -- no routine that takes a Spell* reads it
              (each arm that needs a duration recomputes it from the column; over those
              twelve routines the only `word ptr [reg + 0x10]` is a stack argument); the
              spellbook hover fold reads it from a record (TEXT-096).
Serialize stores +0x08, +0x09, +0x0a as bytes and +0x0c as a word. Nothing else.

Spellbook    0x1c bytes, at actor+0x140
  +0x00  vtable
  +0x04  CObArray (0x14 B) SUBSCRIPTED BY SPELL ID; a null element means "not known"
  +0x18  int, the id of the last successful lookup
```

Learning a spell is the effect kind `teachSpell` (42) and nothing else: with no book it does
nothing, with the id already present it does nothing, otherwise it constructs `Spell(id)` and
stores it at `book[id]`, then sets `actor+0x150 |= 0x400000`. Allocating a book sets `actor+0x4c`
bit 1, and bit 2 — the mage bit — when the actor has a nonzero `manaMax`.

## Spell power

```
power = clamp( skill[Sphere] + Mind - 30, 0, 100 )       skill[] is actor+0xa8 + 2i
      = the item's castSpell effect +0x42                when casting from an item -- UNCLAMPED
f     = power/30 + 1

damage pair   base   = ftol( damageMin * f )   ONLY when the column is > 0, else 0
              spread = ftol( damageMax * f ) - base    ditto -- a SPREAD, not a maximum
              so a -1 column scores 0, which is what lets a spell reach its own arm at all
duration      ftol( 1.025^power * SpellDuration * 16 ) ticks
              invisibility instead: min( ftol(1.05^power * 3 * 16), 65000 ) -- the literal 3,
              NOT its own SpellDuration column, which is read only as the > 0 gate
              when SpellDuration is 0: (AreaEffectDuaration << 4) + (power << 4)/10
cast range    MaxRange + power/30       (Teleport: MaxRange + power/3)
              ONLY when the MaxRange column is nonzero (L05191): Fire Sacrifice and
              Shield ship 0 and gain nothing at any power
spray victims min( power/20 + 2, 7 ) final victims, primary included
              This lives in R0269, whose only caller is the cast's spell-14 arm. It is
              PRISMATIC SPRAY's output-list cap, not a radius. The ten AreaEffect spells take
              their radius from the Radius column, unscaled; Prismatic secondaries instead
              come from the caster group's visibility population.
```

Mind is worth exactly as much as the school skill. **There is no `+3 × skill` to-hit term and no
`skill/5` damage-floor term for a spell** — the weapon-skill terms of [HERO](../hero/format.md) step 11 have no
counterpart here, and a spell's damage never passes the to-hit roll at all.

**The domain, and the one place the clamp is missing.** Both terms are hard-bounded: a school skill
is clamped to `[0,100]` twice in the derive and a stat is clamped to 100 by every effect arm, so
`skill + Mind - 30` runs over `[-30, 170]` and **the `[0,100]` clamp binds from `skill + Mind = 130`
upward — it is reachable, not defensive**. Three routines compute the expression and they do not
agree: `Spell::Apply` and the Prismatic Spray fan clamp; `R0253` returns `max(x, 0)` as a byte
with **no upper clamp**, and its single caller uses it for the spray victim cap. That makes no observable
difference — the victim count saturates at 7 at power 100, exactly where the clamp begins, and over all
10 201 reachable `(skill, Mind)` pairs the two forms never disagree — but a consumer that clamps
everywhere is right by luck, not by construction. An untouched stat is capped at **50**
([HERO](../hero/format.md) step 0), so **without a cap-raising effect Mind contributes at most 20 points of
power**.

At the ceiling, in the integers the engine stores: `f` tops at **4.333**, so damage is 4.25×–4.33×
the `Data.bin` columns; a duration column multiplies by `1.025^100 = 11.81`; range gains at most
**3** cells and Teleport **33**; Prismatic Spray reaches **7 total victims**. Nothing overflows its store on the
shipped table — the largest damage base is 43 and spread 87 against byte fields, the largest
duration 6 312 against a `u16`.

## Cast distance and client target form

The shipped Fire Ball definition has Sphere1 and Max Range10. The resolver
initializes Spell+9 from Max Range; the ordinary power producer calculates
10+floor(clamp(skill[Sphere]+Mind-30,0,100)/30), hence 10..13. Its general
non-Teleport branch leaves a zero base unchanged. Player-command handlers
and the unit-order constructor **copy the stored Spell+9** into order+0x14,
replacing weapon reach. These copies do not prove that current power was
recomputed before the distance decision (`MAGIC-REACH-178`). The admitted
manual setters clear order+0x60 from the register their heads zero. Parent
states 0x0d/0x0e reissue kind 8/9 and reload reach from the saved Spell+9 on
each evaluation, preserving the progress byte and the zero flag. This is a stored-value reload, not a power
recalculation (`MAGIC-253`).

For child order8, a non-self target must first match the mover's current
facing. The typed actor predicate then uses Position cell bytes+0/+1,
fractions+4/+5 and virtual+0x1c tokenSize. For each axis:

```
centre = u16(((size + 2*cell + 511) << 7) + fraction)
gap = max(abs(centreCaster-centreTarget) - (sizeCaster+sizeTarget)*128, 0)
distance = 1 + floor(max(gapX, gapY)/256)
```

The low byte of distance must be <= order+0x14. This is a maximum-axis
footprint-gap comparison; it differs from the melee strike distance formula.
Self-target bypasses facing and distance. Success installs action0x0d with
Spell/target; failure reaches the unit approach entry (`MAGIC-REACH-179`).

For child order9, facing comes from current Position+0/+1 to the point, but
the distance is max(abs(dx),abs(dy)) from Position+2/+3. It does not use
footprint size or fraction bytes in that metric. Success installs action0x0e
with Spell/cell; failure reaches the point approach entry. Neither selected
typed predicate nor its heading helpers reads the world or altitude. With
matching facing, centred1x1 actors and stored range10, axis and diagonal
distance10 admit and 11 fail on flat/up/down synthetic height configurations;
range11 and 13 move the boundary to 11/12 and 13/14. This is conditional on
reaching those child-order arms, not proof of unchanged whole-game casting
when height changes (`MAGIC-REACH-180`).

Ordinary Fire Ball click selection uses the **point** builder even over an
actor: selector index1 has target-type0, click adds 1, and the command remap
retains spell2. An explicit unit-target cast order therefore does not describe
the ordinary Fire Ball click. The selected spell cursor arm also has a
separate visibility-field gate: four tile words OR/masked to0xc000 retain
the cast choice; a different aggregate sets capability0x400 and changes it
to move. Mixed0x8000/0x4000 corners satisfy the aggregate, so this is not an
individual-corner visibility rule. These fragments have no caster-distance
comparison (`MAGIC-REACH-181`).

Unknown: gates between the parent AI state and these child orders; which
power snapshot the next order receives; altitude-to-client visibility and
projected screen-picking updates for a particular map; complete arrival and
native release. Existing altitude-sensitive sight evidence does not supply
that missing joined cast route. The local no-altitude result must not erase
an upstream visibility or order gate.

## Cast sequence

```
Spell::Cast(caster, target, x, y)                                   R0268
  if the caster is a mage AND caster+0x68 == 0   (i.e. not an item cast):
        if spell.manaCost > caster.mana:  refuse, nothing happens
        caster.mana -= spell.manaCost
  a fighter is never charged; an item cast is never charged; the cost never scales
  casting at anyone but yourself sets any active invisibility's remaining duration to 1
  DeliverySystem == 2:  delay = distance(caster, target) / SpellEffectSpeed
                        Lightning and Prismatic Spray: delay = 5, flat
  else                  delay = 0
  send the cast-animation packet carrying that presentation delay
  ordinary phase5 later prepares the simulation payload after charge
```

The packet delay is not a simulation apply queue. Delivery2 creates a separate
SpellTransport whose own countdown releases its child effect. IDs13/14 set that
counter to 10; ID14 selects and prepares its victim list during admission, while
the later unit-target wrapper suppresses a duplicate. — MAGIC-DELIVERY-170,
MAGIC-CASTCLOCK-171

The cast itself runs through the actor's common action phases. Phase 0 loads charge
`actor+0x134`; phase 5 applies when that countdown reaches zero; phase 7 loads relax
`actor+0x135 + U[0,3]` and the equipped-Humanoid weight/Reaction penalty from
the [hero format](../hero/format.md). It also adds this Spell row's `Complication Level` when
state is `0x0d` or `0x0e` and `actor+0x64` is still non-null. For an uninterrupted retained
order cast, application-to-application cadence is

```
charge + relax + U[0,3] + humanoidPenalty + ComplicationLevel + 2 actor ticks
```

The caster weapon-diversion route enters state `0x0d`, but application clears `actor+0x68` and
`actor+0x64` before phase 7 tests the Spell pointer. Its cadence therefore omits
`ComplicationLevel`:

```
charge + relax + U[0,3] + humanoidPenalty + 2 actor ticks
```

The final two ticks are the completion latch crossing the order-before-action boundary, not another
stored timer. Insufficient mana returns from phase 0 before changing phase, countdown or the
completion latch, so it writes no failed-cast recovery. If a prior completion left that latch set,
the retained order attempts admission for three actor ticks, skips one while progress consumes the
stale latch, then repeats. With a zero latch, refusal continues to be attempted every actor tick.

## AI cast selection

**A creature can reach a cast through two of at least four independent mechanisms, not only the
Mind-gated mage check below.** The two this page traces are the engage selector's own three-slot
loop and the Mind-gated walk. Two more are already published elsewhere: the defender's own Heal
path (`R0305`, `AI-FOLLOWHEAL-118`) and `R0113` → `R0394` (`AI-340`);
`R0015`, which calls `R0018` twice, has not been read at all (`EnumRefs
callto:R0018`: 13 hits / 8 owners).

For an actor without the mage bit, the engage selector's own three-slot loop over
`ord+0x78`/`+0x7c`/`+0x80` (each paired with a probability dword at `ord+0x84`/`+0x88`/`+0x8c`) is
unconditional inside `R0009`; a mage-bit actor with reach `< 2` and owner `+0x28 == 0` skips
it instead (`AI-341`). The loop draws one independent `rand()` per non-empty slot and compares it
against that slot's own stored threshold — the threshold is the class's Probability column scaled
by 327 (`UNIT-SPELL-007`), so a slot's own match probability is `P = p·327/32768` (99.8 % at
`p = 100`) — a match short-circuits straight into the cast dispatcher (`R0209`) with the
slot's spell id; no mage bit, no Mind gate, no mana or `Defensive` filter, and the loop always runs
all three slots so a later slot's match can overwrite an earlier one (`AI-341`). These three slots
are not a separate cache: they are the same class-spellbook columns ([UNIT](../unit/format.md),
`UNIT-SPELL-007`) every one of the twelve spellbook classes carries, set once at spawn; within the
AI module, no writer other than that spawn setup and `Order::Serialize`'s own raw LOAD copy on a
SAV load was found (`AI-341`, `SAV-1066`). Of the twelve unit classes `R0184` sets up, only
Dragon and Daemon receive the mage bit there that can reach the second mechanism below; a Human
with positive streamed `ManaMax` also carries the same bit through a different writer
(`HERO-CLASS-013` as amended, `MAGIC-AI-012`). Only when the loop finds no match does the actor fall through
toward the mage-only choice.

Only a **mage** reaches that choice (`actor+0x4c & 4`, `[actor+0x14]+0x28 != 0`, and a non-null
target after acquisition, `L13016`), and then:

```
if (Mind > 59 and rand()*100/0x8000 < 30):
        re-run the engage selector (R0009) in full; this call itself orders nothing, but
        the re-entered call can find a slot match, or reach this same branch again and recurse
otherwise, over ids 1..28 through Spellbook::Get:
        keep a spell iff  ManaCost <= current mana  AND  Defensive == 0
        if the kept list is non-empty: pick one uniformly, order it
        if it is empty: cast id 28 (Slow) anyway if the spellbook holds it — unfiltered by mana
                or Defensive, because it is only the walk's own last-tested id, not a deliberate
                fallback spell — otherwise cast nothing
```

So **a monster never casts a defensive spell** through the ordinary pick, never chooses by school,
skill or power, and Mind above 59 makes it cast *less often*, not better. `0x8000` is the AI
class's own RNG range, so both rolls are exact. The "hand back" branch re-runs the engage selector
rather than simply returning, so this tick can still end in a cast through that selector's own slot
loop, or recurse into the same Mind-gated choice again. The "otherwise" branch is a second,
independent cast attempt over the full spellbook (not the three slots above) that can itself issue
an order (`MAGIC-AI-012`).

The two mechanisms do not share one choke point. `R0209`'s own 28-entry jump table
(`L05156`, keyed by spell id) sends 16 of its 28 ids to the shared writer `R0018`, which
writes `ord+0x08 = 8`, the Spell pointer at `ord+0x30`, and the target at `ord+0x28` (`MAGIC-221`) —
the same fields [ORDERS](../ai/orders.md) documents for a player-issued cast; 10 ids reach a kind-9
order (`ord+0x08 = 9`, target's own cell at `ord+0x3c`) that `R0209` writes itself, bypassing
`R0018`; the remaining 2 ids write no order at all, only `ord+0x60 = 1`. A null
`Spellbook::Get` lookup writes nothing, not even `ord+0x60` (`MAGIC-221`). The per-slot loop's own
match reaches this table through `R0209`; the Mind-gated walk's own match bypasses
`R0209` entirely and calls `R0018` directly, so it can only ever produce a kind-8,
target-actor order — never the kind-9, target-cell order the slot path's own table arms can
produce for the identical spell id (`MAGIC-221`).

Each table arm fixes the target of a slot-drawn cast (`MAGIC-237`). Fire Arrow, Drain Life,
Lightning, Prismatic Spray, Stone Curse, Curse and Slow are aimed at the engage victim. The
Protections, Heal, Invisibility, Shield, Bless and Haste are aimed at the caster itself. Fire Ball,
Wall of Fire, Freezing Cloud, Poison Cloud, Light, Darkness, Wall of Earth and Meteor Storm are
aimed at the cell the victim stands on. Acid Stream and Teleport are aimed at the cell one step
from the caster toward the victim; Teleport's search for a free cell around the victim runs but
its result is overwritten (`MAGIC-238`). Fire Sacrifice and Control Spirit write no order. No
slot-drawn spell is aimed at an ally; none of the 28 arms is.

A failed kind-8/9 reach or facing test selects approach rather than immediate
release. Approach whose route search leaves a non-empty path re-evaluates the
cast. The former unconditional pending-order-retention wording is narrowed by
two paths. A route search that leaves the path list empty sets mover+0x98
(`L01720` point, `L00161` actor); for a manual cast's parent state
0x0d/0x0e the executor epilogue then clears the pending kind and calls
reacquisition, which writes kind 6, 0x0b or 0.
The point helper's obstacle branch can write pending kind 10 or 0; that local
replacement does not itself clear the parent cast command. Whether a later
parent evaluation reissues the cast after either path, native route
completion and feedback remain Unknown (`MAGIC-239`, amended). `ord+0x60` is read by the two cast install arms: a zero value runs the
stop-and-reset at install, a nonzero value keeps the order armed, and every creature cast selector
found writes 1 (`MAGIC-240`). A selector pass that finds no slot match writes a kind-5 engage order and
leaves `ord+0x60` alone; the cast already in flight runs from actor fields and is not cancelled by
it (`AI-376`). The selector body does not read the creature's own cast state; its callers' gates were read only in part (`AI-376`). A zero slot id is
skipped before `rand()`, a zero threshold never matches, and a non-mage creature is neither refused
nor charged by `R0268` (`AI-377`). The slot path also issues its order with no
mana test at selection, unlike the Mind-gated walk and `R0305`. Unknown: whether repeated path failure ends the walk, a dead target at the kind-8 arm, a null target at the cell arms from the unguarded call site `L01988`, the reach of Dragon and Daemon (`actor+0x12c < 2`), and whether the selector is reached during a cast. Whether the cast applies identically to a player's after the cast state begins was not traced.

## Mind and Spirit

Enumerated image-wide (`EnumRefs disp:88`: 461 hits / 235 owners / 4 orphan, all four stack frames;
`disp:8a`: 36 / 13 / 0), with the actor's `u16` width used as the filter and every byte- and
word-wide hit read:

```
Mind    -> the power term, one for one with the school skill    (the three spell-power sites above)
        -> Prismatic Spray's victim cap, through the one unclamped helper
        -> whether an AI mage casts at all                      (AI cast selection above)
        -> Control Spirit copies the TARGET's Mind into the Ghost
Spirit  -> the mana pool, through manaMax        \ each in ONE hop, inside the derive
        -> the five protections, through spirit/2 /
        -> Control Spirit copies the TARGET's Spirit into the Ghost
```

Neither stat is read directly by the mana-cost comparison in `Spell::Cast`
(`R0268`). In that specific body, the explicit refusal is the mana gate;
there is no separate cooldown, interruption or random-failure check. These
local observations do not describe the earlier order range/facing and client
visibility gates. Mind contributes to the stored-range producer, while the
refresh timing and which power snapshot an order receives remain Unknown
(`MAGIC-REACH-178`, `MAGIC-REACH-179`, `MAGIC-REACH-180`, `MAGIC-REACH-181`).

The separate spellbook result remains: **the book's capacity is not bounded
by a stat, what may be learned is not chosen by one, and learning is not
gated by one**.

The blind spot this enumeration cannot close: a wholesale `REP MOVSD` copy of an actor would carry
no displacement and no `disp:` sweep can see it.

## Prismatic Spray heading and facing

The heading the selector uses is one of eight bytes, multiples of 0x20, from the signs and a 2:1 magnitude test of
the fine centres of caster and candidate; each centre is the fine point plus 128 per extra footprint cell, so
the sub-cell offset is included. The score is the edge gap shifted left 8 plus the folded facing difference, masked
to 16 bits: a ranking term (`MAGIC-242`). No facing test is found in the three routines read: the cast routine, its Prismatic Spray
helper and the selector. Book and scroll casts are gated before the cast by the executor's unit-target
and point-target rows, a weapon cast by the attack rows but not by the plain row that installs action 3
(`MAGIC-243`, the arm-to-row map Medium). Which code issues that plain row is Unknown.
