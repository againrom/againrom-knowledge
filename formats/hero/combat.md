# Combat and regeneration

[Reference](format.md)

## Hit resolution (`R0265`, the actor's `vt+0x4c`)

For a unit target, the decoded blow and direct-damage effect use this resolver (`HERO-DAMAGE-022`).
Building targets use a different resolver (`UNIT-STRUCTDAMAGE-064`). Its second argument is a `0x16`-byte
combat block of the `+0xa6` layout — the attacker's own live block for a melee strike, an
`Effect_DirectDamage`'s copy at `+0x48` for a spell ([MAGIC](../magic/format.md)).

```
rand(n)     = rand() * (n+1) >> 15          uniform on [0, n] INCLUSIVE ; 0 when n == 0

dmg   = A.dmgBase + rand(A.dmgSpread)
        bless (spell 23) active on the attacker and rand(100) < its magnitude
                                            -> dmg = A.dmgBase + A.dmgSpread, no roll
        curse (spell 27) likewise           -> dmg = A.dmgBase, no roll
hit   iff  A.toHit + (rand(200) - 100) > target.defence
      or   attacker+0x4c & 0x10                 set by attackKind == 3: Bat_Sonic, Bee
      or   the same roll >= 90              (the roll < -90 arm is inert: it re-assigns 0)
miss  -> the physical component is 0, the elemental one still runs if A has no physical damage

dmg  -= target.absorption                   FLAT, and physical only ; floored at 0
dmg   = ftol( dmg * (100 - target[0xce + A.activeSkill]) / 100 + 0.75 )

second component, run whenever that pair is nonzero -- NOT gated on the hit,
and no absorption is subtracted from it:
        v = A[+0x11] + rand(A[+0x12])
        v = ftol( v * (100 - target.protection[2]) / 100 + 0.75 ) ; max(v, 0)
        A[+0x11]/A[+0x12] = actor+0xb7/+0xb8, written only by the units streamer's
        attackKind == 1 arm -- which NO shipped class takes, so this pair is 0 in play

elemental component, applied iff the swing landed OR A has no physical damage at all:
        v = A.elemBase + rand(A.elemSpread)
        p = target.protection[A.elemKind]   kind 1..5 = Fire/Water/Air/Earth/Astral
        v = ftol( v * (100 - p) / 100 + 0.75 ) ; max(v, 0)

result = max(sum, 0) ; the melee strike subtracts it from target+0x94
side effect: target+0x40 = the attacker, target+0x48 = A.elemKind when the attacker is a mage
```

The six bytes at `+0xce … +0xd3` are therefore **resistance by weapon class**, read from the
target and indexed by the attacker's own active byte (`attacker+0xb6` for a melee strike; 0 =
General, 1..5 = blade/axe/bludgeon/pike/shooting; `HERO-DKIDX-162`), and they are never
re-derived for a human. No shipped Armor or Shield block and no effect-arm body writes the modifier bytes (the live
bytes are recomputed from the modifier; a block loaded from a SAV is Unknown). The writers found
are the `#modify … +god` order and two spawn paths (all six bytes
and six protections set to 100), the archive and the constructor zero fill (`HERO-MODDK-161`). The
active byte is the `Weapons` attackType for types below 10 and 0 otherwise; the other callers of
the resolver were not read. A protection is a
percentage, clamped to `[0,100]` by step 16, so 100 is immunity.

## Regeneration (`R0654`, the actor's `vt+0x14`)

The regeneration route first requires dword `actor+0x54 != 16` and signed
health>0. Health0 is untouched; negative health belongs to the separate decay
route. The helper computes signed `i32(server.sub - actor.due)` from server
`+04` and actor `+138`; a result>80 selects local rate 3, otherwise 1. This is
the arithmetic on an invocation, not a new scheduler or post-load timing rule.
— HERO-REGEN-021 (amended), SAV-REGENORDER-531

| arm | current / maximum / period | modifier / remainder | extra local gates |
|---|---|---|---|
| health | `+94 / +96 / +98` | `+de / +a2` | current<maximum, period!=0, signed server full-counter remainder modulo4==0 |
| mana | `+9a / +9c / +9e` | `+e2 / +a3` | current<maximum; no period or full-counter filter |

All eight pool/max/period/modifier loads are signed16; the remainder loads
are unsigned8. Let `i32`/`i16` mean signed interpretation of the low32/16 bits,
`u8` mean low8 bits, and `T(a,b)` mean signed division truncated toward zero.
For a reached arm, with factor 2 for health and 1 for mana:

```
base = i32(current * 100 + remainder)
n = i32(maximum * factor)
n = i32(n * (modifier + 100))
n = i32(n * rate)
acc = i32(base + T(n, period))
q = T(acc, 100)
remainder = u8(acc - q * 100)       # first persistent store
current = i16(q)                    # second persistent store
current = min(i16(current), maximum) # signed read-back, third store
```

The `+100` sum is 32-bit, not narrowed to a word. The routine actually performs
two separate `IDIV100` operations, storing the first remainder before the
second quotient. There is no lower clamp. The final upper-bound store does not
clear the remainder. Negative remainders become bytes 157..255 and reload
unsigned; initial stored bytes 100..156 are also consumed without validation.
Health32766/max 32767/period 1/modifier 0/remainder 0/rate 1 yields 32764 after
quotient 98300 narrows, not 32767. — SAV-REGENWIDTH-528, SAV-REGENSTORE-529

Health period 0 skips its arm. Reached mana period 0 causes a divide fault before
either mana store. Products and accumulators can wrap without fault. Under
stable gate operands, valid memory and rate 1/3, the admitted product cannot
equal INT_MIN modulo2^32, so signed division overflow is excluded; injected
arm entries or corrupt locals are outside that proof. The operating-system
fault/recovery outcome remains Unknown. — SAV-REGENFAULT-530

The health stores at `L08217/L08218/L04038` precede callback
`L08224 -> R0059`. Mana then rereads its operands, without another
health test or rate calculation, and stores at `L08219/L08220/L04040`
before callback `L08225`. The local body is not transactional. An unchanged,
normally returning first callback permits a later mana divide fault after
the health stores; actual callback mutations and loaded chronology remain
Unknown. — SAV-REGENORDER-531

The six pool/max/period words and two remainder bytes are serialized at their
existing widths; the modifier words are inside the earlier raw64 block at
`+d4`. Serialization does not normalize these values or establish a safe
initial vector. — SAV-REGENWIRE-532

The modifiers scale the base rate rather than adding a pool amount. No stat
is directly read: Spirit contributes through `manaMax`, Mind not at all.
— HERO-REGEN-021 (amended)

## Combat loop

`R0037` is the actor's `vt+0x18`, run once per tick for every actor on the tick list
(`claims/move.md`). The attack lives in one arm of it, as a three-phase cycle:

```
actor+0x58   sub-phase        actor+0x6c   countdown byte
0            wind-up          = attackChargeTime + extra,  then sub-phase 5
                              extra = (dist*256 + 128)/200 when dist > 1, else 0
5            counting down    the blow lands on the tick the byte reaches EXACTLY 0
                              -> R0001 -> unit: R0246 -> target->vt+0x4c
                                                Building: R0563 -> R0653
                              then sub-phase 7
7            recovery         = attackRelaxTime + rand(3) + humanoidPenalty
                              at 0 -> sub-phase 0 and completion byte +0x136 = 1

humanoidPenalty = clamp(IDIV(runtimeWeaponWeight + 5*(30-Reaction), 12), 0, 12)
                  only with an equipped weapon and the Humanoid predicate; otherwise 0

one blow every   attackChargeTime + attackRelaxTime + rand(3) + extra
                 + humanoidPenalty + 2 actor ticks
                 the final 2 cross the completion-latch/order-machine boundary
                 62 ms per tick at the shipped speed index (16 tps)
```

Movement speed does not enter; Reaction does through `humanoidPenalty`. Neither the number of actors,
the tick order, nor adjacency changes the stored tick count: an actor
whose countdown is not zero cannot strike this tick, and one whose countdown is zero strikes
whatever `actor+0x5c` holds if it is in reach.

```
actor+0x5c    the combat target      a SECOND field; the destination is in the mover / order block
actor+0x54    3 = attacking, 1 = moving, 0xd/0xe = casting, 0x10 = torn down
actor+0x12c   reach, in cells        ctor default 1, raised by the weapon (weapon[0x50] - 1)
actor+0xa5    scanRange              read by the AI module, NOT by the attack path
actor+0x49    tokenSize (footprint)  vt+0x1c
actor+0x4a    movementType           > 1 leaves no corpse (Ghost, Bee, Bat_Sonic, Dragon)

in reach iff   attacker+0x12c >= R0247(attacker, target), where
               d = max(|dx|, |dy|)                                1/256-cell units
               d -= ((sizeA + sizeB) << 7) - 256
               result = (d <= 0x180) ? 1 : (d + 0x40) >> 8
an order also requires the actor to be FACING the target before state 3 is entered
out of reach -> a move-to-actor order and state 1; the strike sub-phases run only in state 3
```

**State 1 and state 3 are mutually exclusive, and a crossing tick cancels the attack-cycle's
own phase rather than freezing it.** The order machine's own crossing arm forces
`actor+0x54 = 1` unconditionally while a cell boundary is being crossed and clears it to `0`
(idle), not directly to `3`, on arrival. Every state-`1` tick's own walk handler stores the
attack sub-phase `actor+0x58 = 0` before jumping to the shared tail, without ever executing an
instruction inside the span holding all six of state `3`'s own `actor+0x6c` writers, so a
crossing tick cannot itself read or advance the countdown. That `+0x58 = 0` store cancels the
phase: the sub-phase switch's own `0` arm never reads `actor+0x6c` on any later tick either — it
either exits untouched or starts a fresh charge, which overwrites `actor+0x6c` with a new base
before anything else reads it. So the countdown's old value sits stale through a crossing and is
overwritten, not resumed, at the next charge start. A crossing actor is action state `1` **and**
order-progress `3` together, not state `1` alone (`MOVE-STEP-040`'s own progress-3 arm is what
sets state 1 for a transit; state `1` is also the arrival state of two idle-turn arms at a
different progress value). Of the permitted archive's alive population, 38
state-1/progress-3 (crossing) records exist, and every one carries `actor+0x6c = 0` — the corpus
contains no example of a crossing actor carrying a live countdown (`HERO-CROSSHOLD-146`).

The start routine itself does **not** reject out-of-reach targets: it computes distance for
extra delay. Application rechecks reach (`HERO-CADENCE-115`, `UNIT-STRUCTREACH-063`). The
Building branch reads only combat bytes `+0x13/+0x14`, with a positive-spread gate and flat
subtraction 5; ordinary physical `+0x0e/+0x0f` is ignored. A physical Building hit writes the
word HP without clamping, unlike the direct-effect Building consumer, which clamps at zero.
Reach below 2 versus at least 2 selects physical message `0x71` versus `0x72`, not another
server strike callback (`UNIT-STRUCTDAMAGE-064`, `UNIT-STRUCTDELIVERY-065`).

The acquisition clause of `HERO-TARGET-024` is retracted and replaced by
`AI-ACQUIRE-002`: unordered units use `R0022`; visibility supplies
the candidate population and reach bounds selection. Turn cost resolves
selection rather than nearest distance. Both acquisition routes can fall
back to corpses when no living enemy qualifies. See [AI acquisition](../ai/targeting.md).

**A hit makes the two players mutually hostile** (bit 0, both directions) — that, not a per-actor
assignment, is retaliation (`HERO-AGGRO-028`, bounded by `AI-DIPLO-004`'s bit 1; that row's
`MinimalGuardRange` default clause is superseded and does not touch the bit).

```
death, from the killing blow to the freed id  (HERO-DEATH-026, SESS-TICK-004,
                                               SESS-TICK-006, HERO-DWELL-065,
                                               HERO-FINISH-066, HERO-DYETICK-067,
                                               HERO-REVIVE-068, HERO-DECAY-069,
                                               HERO-ZERO-070)
  TWO clocks act on a dying body when stepping is admitted. Campaign outcome panels
  pause frontend stepping before teardown (SESS-DEFEAT-065); primary-fall reporting
  can also replace owned actors' HP with -50 (MISSION-DEFEAT-045).

  A. vt+0x18, R0037, once per SUB-TICK
  tick 0        health <= 0 : defence halved, cast interrupted,
                              countdown actor+0x6c = dyingTime - 1
                              the corpse STILL OCCUPIES ITS CELLS
  while > 0     countdown -= 1                       decremented ONLY while positive
  at 0          the byte PINS there -- nothing writes it again while dying -- and
                from this sub-tick on the tail re-runs on EVERY invocation:
                  movementType > 1 -> health = -1000    (Ghost, Bee, Bat_Sonic, Dragon)
                  if health <= -10 -> torn down          <-- a STANDING condition,
                    every footprint cell released, the reserved next cell released,
                    the actor unlinked from the tick list onto the dead list,
                    gold dropped: treasureMin.1 + rand(treasureMax.1) when
                                  treasure.1 Gold > rand(100)

  B. vt+0x14, R0654, once per FULL tick -- the SAME routine that regenerates
     the living, reached over the LIVE list, which a corpse is still in
  health > 0    regeneration
  health < 0    health -= 1 every FOURTH full tick        <-- what makes A's standing
                                                              condition come true
  health == 0   entered (JLE) and left untouched (JGE)    <-- NOTHING HAPPENS, EVER

  C. after the teardown, R0867 over the DEAD list, once per FULL tick
  then          one health point per two FULL ticks; stages at -10, -20, -40
  at -600       the runtime id's bitmap bit is freed and actor+0x4 = 0

  so a ground unit felled to -h, 1 <= h <= 9, and left alone:
    >= dyingTime-1 sub-ticks   ~437 ms   the countdown, then it pins
    (10-h) x 4 full ticks      up to ~36 s   B carries health to -10
                                             A tears it down: cells free
    590 x 2 full ticks         ~19.5 min     C: the id returns to the pool
  and a ground unit felled to EXACTLY 0 does none of it, for the whole mission.

experience, per LANDED blow and only to a human class (the monster slot is a stub)
  xp = ftol( target.XPvalue * 0.5 * dmg / target.healthMax + 1 )
  then the [experience](experience.md) Mind scaling; refused when the two share a player, or when the
  diplomacy byte has bit 1 -- and, past that, capped and credited to ONE slot by
  the [experience rules](experience.md), which is where it becomes a skill level
```

**`dyingTime` is a floor on the dwell, not the whole of it, and a corpse is not out of reach.**
The countdown is in **sub-ticks** — 62.5 ms each at the shipped default speed, so the human
default `dyingTime = 8` floors the dwell at ≈ 437 ms and not the ≈ 7 s a full-tick reading gives
(`HERO-DYETICK-067`). What ends the dwell is the standing `health <= -10`, and health keeps
moving after the killing blow: the strike tests only the **attacker**'s health and subtracts
unconditionally, so a body can be struck again (`HERO-FINISH-066`). A consumer must therefore
implement three things this spec used to imply away. **What ends the dwell without anyone doing
anything is the regeneration tick** — `vt+0x14` is dispatched over the *live* list, which a body
that has not been torn down is still in, and its dying arm removes one health point every fourth
**full** tick while health is strictly negative (`HERO-DECAY-069`). So a body felled at −1…−9
reaches −10 by itself in at most ≈ 36 s and is torn down then; nothing about it changes visibly in
the meantime, because the corpse stage byte does not move until −10. **The exception is one
value.** The arm is entered on `health <= 0` and left on `health >= 0`, so a ground unit reduced to
**exactly 0** — a blow whose damage equals the remaining health, which nothing prevents, since the
subtraction has no floor — is a corpse no clock can move: cells, reserved next cell and runtime id
held for the whole mission at corpse stage 1 (`HERO-ZERO-070`). Flying classes cannot reach it,
their arm forcing −1000 regardless. And the first tick's
effects are **reversible**: a heal that carries health back above 0 clears the corpse stage,
doubles the halved defence and floors health at 1, so a corpse inside the window is a live actor
with non-positive health (`HERO-REVIVE-068`). −10 is written at three sites, all in
`R0246` — the damage numeral, the experience/retaliation call, and the death arm's own
guard — and it is one line, not three rules.

Ranged is the same machine: the only melee/ranged test in the simulation is
`attacker+0x12c > 1`, and it picks between two **network opcodes** (0x71 melee, 0x72 shot). No
projectile actor carries the damage — the shot resolves through the same `R0265` after the
distance-proportional `extra` ticks.
