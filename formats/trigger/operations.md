# Operation vocabulary

[Reference](format.md)

## Operation vocabulary

The promoted core state and outcome operations below are `TRIG-ACT-004` (instants
3, 8, 4 and 5) and `TRIG-COND-003` (check 19 and build-time opcode `0x10002`):

```
instant 3   slot[p0] = p1        set a mission variable
instant 8   slot[p0]++           increment it
instant 4   win
instant 5   lose
check   19  slot = slot[p0]      read a variable back into a comparable slot
check   0x10002  (build-time)    declare a variable and preset it to value[0]
```

A condition's constant operand is itself a check node with opcode `0x10002`, given a slot and
preset once at build. There is no immediate operand anywhere: **a comparison is always slot
against slot.**

Additional operation-specific rules
(`TRIG-SACK-022`…`TRIG-GIVEALL-025`):

```
check   14   slot = (a sack exists at cell (p0,p1)) ? 1 : 0
             p0 and p1 are read as BYTES; the sack itself is never yielded
instant 2    broadcast a "show message p0" packet to every client
             writes no slot, no counter, no latch and no actor field
instant 20   move the unit's whole item container to a sack at the unit's OWN cell,
             merging into a sack already there, and give the unit a fresh empty one
instant 28   move the whole container from the first Target_Unit to the second,
             destroy the source, re-seat the giver, notify each owner separately
```

The container is `actor+0x7c` (`ITEM-CONT-004`) and the routine that moves it **deletes**
it, which is why both instants must re-seat the giver. A consumer maintaining a cell→sack
index for pick-ups already has everything check 14 needs; one that does not will implement
the check against the wrong structure.

Three more (`TRIG-ADDITEM-027` (amended)…`TRIG-NEAREST-029`):

```
instant 12   obj = itemFactory(rec+0x40); if obj, add it to the Target_Unit's container
             a CREATION, not a transfer -- nothing is taken from anywhere.
             The owner notification runs even when the code decodes to null.
instant 23   Target_Player money += p0, then a packet carrying the NEW balance
             addressed to that player alone. An accumulate, no clamp, no floor.
check   15   slot = min over every unit the Target_Player owns of
             chebyshev(unit cell, (p0,p1)); 0xff when the player owns none
             p0 and p1 are read as BYTES. There is no unit parameter.
```

Check 15 returns 0xff when no qualifying actor exists. Comparing that result
with `<=constant` differs from treating an empty result as zero.

### The three item arms (`TRIG-TAKEITEM-038`, `TRIG-XFERITEM-039`, `TRIG-ITEMTEST-040`)

Instants 11, 12 and 13 all end by resyncing a unit's inventory to its owner, and they are
easy to confuse. They differ in what they do to the container and in how many packets they
emit.

```
instant 11   take ONE unit of item (u16)rec+0x40 out of the FIRST Target_Unit's
             container; notify that unit's owner UNCONDITIONALLY; then, only if
             something was taken, add it to the SECOND Target_Unit's container and
             notify that owner too. One packet on a miss, two on a hit.
             No shipped map authors this opcode.
instant 12   obj = itemFactory((u16)rec+0x40); if obj, add it to the Target_Unit's
             container. A CREATION -- nothing is taken from anywhere. One packet,
             emitted even when the code decodes to null.
instant 13   take ONE unit of item (u16)rec+0x40 out of the Target_Unit's container
             and FREE it through the object's own scalar deleting destructor
             (vtable slot +0x04, argument 1). One packet, emitted even on a miss.
check   12   1 iff the Target_Unit's container holds item (u16)rec+0x40, else 0
check   17   the same 64 bytes as check 12, emitted twice. Only 17 is authored.
```

The take is `R1522` and arms 11 and 13 are its only callers in the image. It
**detaches**: a stack of `n > 1` is split, one unit off through the item's `vt+0x40`, and a
count of `0` or `1` is unlinked from the container list outright. Either way the container's
running load `+0x20` drops by `(i16)+0x4a * (u16)+0x42`. So instant 13 removes exactly one
unit of a stack, not the stack, and a consumer must unlink before freeing. Checks 12 and 17
run the same lookup without detaching, so either is the non-destructive predicate for
"would instant 13 destroy something here".

The item field is a **word** at `rec+0x40`, written by the builder as `0xe18 + V` from the
authored `Item` value `V`, and it is matched against `(u16)item+0x40` on the object. Instant
11's destination comes from `rec+0x3c`, which the builder fills with the *second* reference
of whichever kind reached its own second slot -- unit, group and player each have their own
first/second flag and all three store to that one field. So `rec+0x3c` holds a unit only when
the node declares two `Target_Unit` parameters, which is what the editor's own declaration
for the opcode does.

### The campaign-reachable closure arms (`TRIG-CLOSURE-037`, amended)

The last campaign-reachable helpers are exact enough to implement directly. Parameters below
are the builder-packed `p` sequence, not the authored `Par` positions.

```
check 9:
    if subject.order.pending != 5: slot = 0
    else:                          slot = subject.order.target.mapUnitID

instant 21: castFromCell(p0, p1, p2, p3, spell=p4, power=(p5 != 0 ? p5 : 99))
instant 24: castFromCellAtUnit(p0, p1, Target_Unit, spell=p2, power=p3)

instant 29:
    key = (u16)(((u16)p1 << 8) + (u16)p0)     // a 16-bit ADD, not an OR
    for effect in sixEffectSlotsOfCellRecord(key):
        if effect != null && (u32)effect.spellID == p2: effect.life = u16(p3)

instant 30:
    for effect in Target_Unit.attachedEffects:
        if effect.spellID == u8(p0): effect.duration = u16(p1)

instant 34:
    if p0 == 6:  Target_Unit.health     = u16(p1)
    if p0 == 15: Target_Unit.defence    = u16(p1)
    if p0 == 16: Target_Unit.absorption = u16(p1)
    notify(Target_Unit)                 // also for every unsupported selector
```

The temporary-caster builders place `power` in `skill[Sphere]`. The accessor
`R0915` reads byte 0 of Spell-row parameter 2, whose title is `Sphere`, and the caller
masks that result to a byte before indexing the six-word skill vector. Shipped spells use only
Sphere 1..5; a custom Sphere-0 row writes General.

Instants 21 and 24 allocate a temporary actor at the source cell. A failed cast attempt retries
on a later actor tick; a successful one counts down, applies the spell, and removes the actor from
`session+0x2c`. The original save format serializes the resulting effects but not that in-flight
actor. A deterministic consumer therefore includes it in live simulation state without inventing
save persistence (`TRIG-CAST-033`).

These helpers are not script-exclusive (`TRIG-CASTACTOR-044`, exclusivity clause
retracted). A map-authored type-9 cell record also reaches the unit-target helper
through actor footprint attachment (`UNIT-M10CELL-054`, `UNIT-M10ENTRY-055`).
The two direct constructor callers do not constrain their own upstream callers.

Instant 25 writes a persistent actor-entry trigger:

```
cell = (u8(p3) << 8) | u8(p2)
tail = {u8(p0), u8(p1), 0, u8(p3), 0, u8(p3)}
writeOrCreateCellRecord(cell, tail)
```

Actor attachment accepts `tail[0] != 0 && tail[0] != 26`. It uses `tail[0]` as the spell id,
`tail[1]` as power and `tail[2:4]` as the temporary caster's source x/y; the entering actor or its
current cell is the target, selected through the spell table. A separate arrival reader accepts
`tail[0] == 26` and treats `tail[4:6]` as relocation x/y. The installed nodes write spell IDs 3 and 9, selecting actor-entry
casting. Entry into either restored cell and completion of its effect remain
Unknown at runtime. `TRIG-CELLTAIL-035` is amended: its sole-reader and shipped no-effect clauses are
retracted. — TRIG-CELLTAIL-035, SAV-CELLLOAD-111

Check 4's shipped selector 6 reads signed current health; check 21 reads signed current building
health. Their field identities are supplied by `HERO-HEALTH-032`, `ALM-CLS-053` and
`SAV-BLDG-037`.

### Reachability and Unknowns

Every successfully built check is evaluated. Actions and action-6 subcommands
require a surviving trigger reference to their node ID. Dormant and unreferenced
arms remain separate from campaign-reachable operations. EN/RU use the same
operation classes; their authored node counts can differ. `TRIG-CLOSURE-037` is amended: `TRIG-TAKEITEM-038`
supersedes its completeness clause. — TRIG-CLOSURE-037,
TRIG-INSTCENSUS-046

The **editor's** names for these arms live in `Description Checks.ini` / `Description
Instants.ini` at the EN install root. Those are Map Editor files — `rom.exe` contains none
of their literals and the RU root ships neither them nor the editor — so a name in them is a
hypothesis about an arm, never a reading of it (`TRIG-CAT-026`). `Get sack` is the standing
example: it names the right index and misdescribes the result. The same names and kinds do
reach the runtime, because the editor stamps them into every authored node
(`node+0x9c + 64*i` name, `node+0x74 + 4*i` kind).


See [map presence and group operations](groups.md) for placement/removal,
cell record operations, member commands, ownership moves and population checks.
