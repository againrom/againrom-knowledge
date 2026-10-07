# Effect action gates

[Reference](format.md)

## Effect action gates

`MAGIC-ACTGATE-079`, `MAGIC-ACTKEY-080`, `MAGIC-MASKREAD-081`, `MAGIC-AIBIT-082`.

### The gate

`actor+0x144` is a bitmask of the effect ids currently attached to the actor; authored Potion effects carry id 0, so it is not a set of spell ids alone (`MAGIC-ATTACH-016` (amended, partially retracted and superseded in part in the ledger)).
`R0016`, the per-actor order machine, reads it before it reads anything about the order.

The gate at `L05213` applies these rules in order:

- when `actor+0x144` is zero the gate is skipped;
- when bit 20 (`0x100000`) of `actor+0x144` is clear the gate is skipped;
- otherwise it reads the order record through the pointer at `actor+0x158`, and when `ord+0x09` is non-zero it leaves it alone;
- when `ord+0x09` is zero it writes the value 4 there.

`ord+0x09` is the progress byte of `AI-PROGRESS-034` (whose four-value liveness clause is
superseded: `0xff` is a fifth live value). The order switch of `AI-ORDER-039` (whose arm-`0xb`
clause is retracted) — walk,
attack, cast at an actor, cast at a cell, and eleven others — is entered only while it is 0
(the branch at `L00547` skips the switch when it is non-zero). Progress value 4 routes to the arm at `L00543`, which sets
`actor+0x54 = 0x1a`, re-tests the same bit, calls nothing, and clears the progress byte only when
the bit is gone.

So one bit refuses every kind of action, and it does so above the point where the kind is chosen.
Value 4 is written at `L00546` and at no other instruction in the image.

### The one escape, and its bound

The refusing arm is not the end of the tick. Both its exits reach the machine's common tail at
`L00010`, which is not gated on the mask: when `mover+0x98` is non-zero it clears that flag and,
for any command state but 1, `0xa` and `0x17`, runs target acquisition `R0004`. If that finds
a target the reach test accepts, it calls `R0005`, whose whole body sets four fields: `ord+0x09` to 1, `ord+0x15` to 0, `actor+0x54` to 3 and `actor+0x5c` to the value of `ord+0x0c`. That is order kind 2's own attack install, replacing the parked 4. The gate does not re-park while the
byte reads 1, so progress arm 1 runs, holding `actor+0x54 = 3` and counting `ord+0x15` until it
exceeds 2 with `actor+0x136` set, then returning the byte to 0; the gate parks the actor again on
the following tick.

The escape cannot repeat inside one refusal. `mover+0x98` has three setters — `L01720` in
`R0178`, `L00161` in `R0043`, `L00162` in `R0055` — and all three are inside
movement executors that are reachable only from an order arm (`MOVE-GATE-039`). The refusing arm
calls none of them, and the tail clears the flag itself. **So a consumer must allow at most one
attack to start at the beginning of a refusal, and none afterwards.**

### What reaches the gate, and what does not

The bit index is the spell id. `R0612` sets `1 << effect+0x0c`, and the per-spell arm stamps
`effect+0x0c` from `spell+0x8`. The gate's immediate `0x00100000` is `1 << 20`, and entry 20 of the
image's spell-name array at `L05034` is `stone_curse`.

Nothing about the effect record reaches the gate: not the kind parsed from the `Effects` column
(`effect+0x3c`), not the mode (`+0x3d`), not the magnitude (`+0x40`), not the duration (`+0x42`).
Spell 20's own `Effects` column parses to `absorbtion=+5`, which is applied through the effect-kind
dispatch of § 5 and is unrelated to the refusal. **An implementation built from the column alone
produces damage absorption and no immobilisation.**

### The observable

| Quantity | While the bit is set |
|---|---|
| position | unchanged; every call site of the routine that displaces an actor lies inside one call of the order machine, and the refusing arm calls none of them (`MOVE-GATE-039`) |
| a step in flight | completes; the gate fires only from progress 0, which is the tick the actor stands on a cell centre (`MOVE-STEP-040`) |
| orders | none of the fifteen arms runs; the machine's tail still runs and may install an attack once, see above |
| act state `actor+0x54` | forced to `0x1a` every tick, which is outside the actor tick's own fifteen-arm switch, so no state arm runs either (`ANIM-PARK-039`) |
| the drawable | drawn grey, with the frame index replaced by the facing so it no longer follows the animation clock; keyed separately on the effect record's `+0x0e` (`MAGIC-STONEDRAW-084`, § 15) |
| end | when `R0673`'s per-tick countdown of `effect+0x42` reaches zero and clears the bit |

Duration before resistance is `ftol(1.025^power × SpellDuration × 16)` ticks — 160 at power 0, 262
at 20, 549 at 50, 1890 at 100 on the shipped row — then multiplied by `(100 − target+0xca)/100`
with a floor of one tick (`MAGIC-SING-019` d; a different clause of this id, the item-cast (Delivery/timing clause narrowed by MAGIC-CASTCLOCK-171.)
universal reach, was retracted).

Two further, spell-independent ways an actor stops acting, which a consumer must not confuse with
this one: `ord+0x09 = 0xff`, written when the queued command is `actor+0x50 == 0x17` and cleared
only by `R0146`; and `actor+0x54 == 0x10`, which makes the actor tick return at `L05756`
before the order machine runs at all.

### The AI side

The same bit is read by the AI without gating anything. `R0225` adds 127 to a melee target's
cost when it carries the bit. `R1048` will not cast spell 20 on a target that already carries
it, and `R0394` generalises that test to any spell with `1 << spell+0x8`.

### Customisation limits

The immediate `0x00100000`, the parked value 4, the act state `0x1a`, the six progress arms and the
255-byte index table are `.text` code and immediates. Changing which spell immobilises, or making a
second spell do it, changes `ROM.EXE`. The duration is data: the row's `Spell Duration` column and
the target's `protectionEarth`. The bitmask is one dword, so the spell id space is bounded at 32 by
the field width regardless of how many rows `Data.bin` carries.
