# Map presence and group operations

[Reference](format.md)

### Map presence — instants 16, 17, 18, 32, 33 (`TRIG-OFFMAP-041`…`TRIG-MAPGROUP-043`)

Being on the map is one bit on the actor and one membership in one list. Two routines own
both, and the script reaches them five ways.

```
R0121(actor):                        // instant 16, and instant 18's first half
    if actor.flags4c & 0x08: return         // idempotent
    map.clearFootprint(actor)               // n x n, n from actor vtable+0x1c
    actor.flags4c |= 0x08
    onMapList.remove(actor)                 // [L00240]+4 = world+0xa4554
    packet(0x74, actor.id, 1, addressee = 0)          // broadcast
    packet(0x74, actor.id, 0, addressee = actor.owner)

R0123(actor):                        // instant 17
    x, y = actor.pos.x, actor.pos.y         // RETAINED from before the removal
    if !place(actor, x, y, r = 0) && !place(actor, x, y, r = 3):
        log "Unit can't return to map - no free place"
        return 0                            // still off the map, nothing else changed
    onMapList.append(actor)
    actor.flags4c &= ~0x08
    notifyEachClient(actor); sendActorDescription(actor, mask 0x20)
    packet(0x74, actor.id, 0, addressee = 0)
    return 1

place(actor, x, y, r):                      // R1145
    actor.pos.set(x, y)
    repeat (r*r)/2 + 2 times:               // 2 at r = 0, 6 at r = 3; the counter is
        cx = x - r/2 + rand(r)              // raised after a FAILED attempt and then
        cy = y - r/2 + rand(r)              // compared JLE against (r*r)/2 + 1
        if map.tryOccupy(actor, cx, cy): commit; return 1
                                            // at r = 0 both attempts are exactly (x, y)
    for cy in [y - r/2 .. y + r/2]:         // skipped entirely when r == 0
        for cx in [x - r/2 .. x + r/2]:
            if map.tryOccupy(actor, cx, cy): commit; return 1
    return 0
```

`rand(n)` is `MISSION-DROP-002`'s inclusive `(rand() * (n+1)) / 32768` (that claim's
`mapObj` base name is superseded; its RNG bound stands).

```
instant 32:  for member in Target_Group: R0121(member)
instant 33:  for member in Target_Group: R0123(member)
instant 18:  R0121(first Target_Unit)
             R0122(second Target_Unit, first.pos.x, first.pos.y, r = 3)
```

`R0122` is `R0123`'s tail with the cell supplied by the caller and one
attempt at `r = 3`; its failure string is `"Unit can't enter map - no free place"`.

State transitions:

- **The return cell is never authored.** No node of opcode 17 or 18 carries a coordinate.
  Instant 17 returns the unit to the cell it stood on, and instant 18 sends the second unit
  to the cell the first stood on. Storing the cell at removal time is therefore mandatory:
  the engine keeps it because it never clears the position object.
- **The return can fail, and the failure is silent in the game.** The unit stays off the
  map, the trigger has already latched, and nothing retries.
- **Nothing else about the unit changes.** Group membership, the item container `+0x7c`, the
  owner `+0x14`, the position object `+0x10` and health are untouched, in both directions.
  A group whose members are all off the map still answers its full count to check 1, which
  is an unfiltered read of `group+0x0c` that only death lowers.
- **Instant 18 needs two `Target_Unit` parameters**, because `rec+0x3c` is the builder's
  second reference of any kind. All six shipped nodes declare two.

The shipped campaign uses the pair as reinforcement and ambush: on `131.alm` two triggers
whose condition is `constant 2 == constant 2` fire on the first pass and take five units off
the map, and a later trigger returns four of them when a group is wiped out and a mission
variable does not hold 1.

### The cell record instant 29 reaches

Instant 29's lookup is the same 52-byte dynamic cell record `TRIG-CELLTAIL-035`'s writer (amended;
its sole-reader clause is retracted) creates and `SAV-CELLREC-017` serializes.

```
R1083(map, key):                     // key is 16 bits: (y << 8) + x
    if !(map.cellFlags[key] & 0x20): return null      // flag array at map+0x10000
    bucket = (key >> 4) % map.bucketCount             // map+0x540bc
    for node in map.buckets[bucket]:                  // map+0x540b8
        if node.key == key:                           // word at node+0x08
            copy 52 bytes from node+0x0c to map+0x5402c
            return map+0x5402c + 0x14
    return null
```

So dwords 5..10 of the record are six area-effect pointers, and the copy is of pointers:
instant 29's write reaches the live effect. The flag array is 65536 entries and the key is
16 bits with `x` in the low byte, so **the addressable cell space is 256 × 256**. The
largest shipped campaign map is exactly 256 × 256.

**Instant 29 sends no packet.** Neither does instant 21: both mutate an object outside the
session's register file and neither calls a notification helper, so the change is invisible
to a client until the effect's own tick, or the cast actor's own ticks, change what is
drawn. The five map-presence arms all notify.

### Action 6: the three member helpers, and the gate on subcommand 10

All three arms first stop every member and set `grpAI+0x20 = 0`, then walk the group again.

**Subcommand 10 scores each member against the named unit before it issues anything.** The
second walk calls the target-cost routine `R0225(member, namedUnit)` and compares the
result with `0xffffff`, the value returned only for a **0** cell of the 4x4 preference matrix
at `AImanager+0xb94` (`AI-PREF-070`, `AI-COST-071`). The matrix's only zeros are
`M[1][3]` and `M[2][3]`, so the veto needs a candidate of movement domain 3.

```
for member in Target_Group:
    if cost(member, Target_Unit) == 0xffffff:      # preference matrix cell is 0
        member.state    = 0xc                      # acquire with no leash
        member.ord[0x00] = member.cell
        member.ord[0x08] = 0
    elif member == Target_Unit:
        member.state    = 0xc                      # same disposition
        member.ord[0x00] = member.cell
        member.ord[0x08] = 0
    else:
        member.state    = 3                        # engage the named unit
        member.ord[0x0c] = Target_Unit
        member.ord[0x14] = (u8)member.reach
        member.ord[0x08] = 0
```

A vetoed member therefore **does not attack the unit the script named**. The same gate is in
the player's attack order `0x19` (`AI-CMD-054`). No shipped node reaches it: every authored
`Victim` is a Humans actor and therefore domain 1 (`TRIG-GRPARM-047`).

**Subcommands 11 and 15 are one shape with two constants** (`AI-DEFEND-111`,
`AI-FOLLOWSET-116`, `AI-FOLLOWRANGE-115`):

```
for member in Target_Group:
    if member == Target_Unit:
        member.state    = 0xc ; member.ord[0x00] = member.cell ; member.ord[0x08] = 0
    else:
        member.state    = 8 for subcommand 11, 0x11 for subcommand 15
        member.ord[0x10] = Target_Unit
        member.ord[0x08] = 0
        member.ord[0x70] = (u8)p1 if (u8)p1 != 0 else 3
```

`p1` is read as a **byte** out of the node's 32-bit int, in the arm and again in the helper, so
an authored 256 behaves as 0 and coerces to 3. The player's defend order passes a literal 0 and
so is always range 3, which makes the script surface the wider of the two.

Authored-but-unreferenced nodes do not make any of the three helpers reachable.

### Instants 19 and 22: the ownership move

Both dispatch to one routine, `R0064`; instant 22 calls it once per member of the
`Target_Group` (`PARTY-JOIN-025`).

```
instant 19: giveActor(Target_Unit,  Target_Player)
instant 22: for member in Target_Group: giveActor(member, Target_Player)

giveActor(actor, newOwner):
    if actor.group:        actor.group.remove(actor)
    actor.owner.flat.remove(actor)          // Player+0x20
    actor.owner = newOwner                  // actor+0x14
    newOwner.flat.append(actor)             // Player+0x20
    g = newGroup()                          // 0x48 bytes
    newOwner.groups.append(g)               // Player+0x24
    g.insert(actor)
    actor.visMask &= ~oldOwner.mask         // actor+0x18, both owners cleared
    actor.visMask &= ~newOwner.mask
    broadcast(actor)
```

It writes no recruit flag, does not touch `Player+0x34`, and does not allocate a new runtime id.
The actor therefore lands in a group of its own, which is what a save of a joined actor shows.

Whether the actor survives the mission is **not** decided here. At mission end the server keeps an
actor only when `0x21 <= actor.typeID < 0x40`; the client first applies an independent keep-bit
filter. The range is **not** the output of every Humans placement. The Humans constructor streams
Data.bin slot 16 and overwrites it with a player-character value only in non-zero constructor mode:
definition-id and explicit typeID placements pass zero, while an npc placement passes its exact
`Hero` flag result. Mission 20 therefore transfers four Humans that retain `0x17,0x0a,0x0a,0x0a`
and are all removed (`PARTY-M20-030`, `PARTY-M20-031`). Mission 40's `npc25` survives because that
npc record is flagged `Hero`. The transfer node itself writes no persistence flag.

### The two population checks count the LIVING, and neither tests health

`1 How many units contains this group` and `8 How many units this player have` are the only two
checks that count a population, and a consumer that implements either as an unfiltered walk over
the map's placements gets a number that never falls. Both are unfiltered — and both answer
*living* anyway, because the container they measure is one death removes from.

```
check 1   slot = *(u32*)(group + 0x0c)      five instructions, L12353, no loop, no test
check 8   slot = count of every member of every group of that player, body = increment by one per member
```

`group + 0x0c` is the group's own `CObList` element count: zeroed by the list constructor,
raised by one in the node allocator, lowered by one in the node freer, and returned by the
class's own `GetCount`. The `AddTail` behind `AddMember` is taken on the group itself, so the
counted list is the group's base and not the patrol list embedded at `group + 0x20`.

**The removal is per SUB-tick, not per full tick, and it comes after the script pass — but its
own trigger is teardown, not death.** The sub-tick body ticks every actor on the on-map list and,
for one whose own per-tick routine has just torn it down (`actor+0x54 == 0x10`), removes it from
its group **before** unlinking it from the world and appending it to the dead list. The tick
driver takes the phase from `server+0x04` before the sub-tick increments it, so on the phase-6
iteration the whole script pass runs first and the reap follows in the same iteration. A member
torn down during the script pass is gone from the group by the end of that sub-tick, and the next
pass is sixteen sub-ticks away: **no pass ever sees a torn-down actor in a group.**

The population checks read current list membership. Teardown removal updates that list in the
same sub-tick, before the next script pass; death alone does not — a merely dying actor's own
per-tick routine force-clears `actor+0x54` to `0`, never to `0x10`, on every tick it is not yet
torn down (`HERO-DYINGTICK-145`).

**The population checks and the group's own order-issuing arm read the identical list, so both
see a dying (corpse-stage) member for the whole dwell between the killing blow and teardown, not
just the sub-tick death lands on.** The removal test this section cites (`actor+0x54 == 0x10`) is
the reap sweep's own trigger, and it is teardown, not death. `group + 0x0c` (check 1's own field,
above) is the same `CObList` the order-5 per-member arm walks — count `[ebp+0xc]`, head `[ebp+4]`
inside `R0026` — so a dying member counted by check 1 or check 8 is the same member the
order machinery still walks, and while that arm's own candidate-list gate is open it keeps
rewriting the member's own pending-order byte once every full tick with no test of its health or
death stage (`AI-332`). "No pass ever sees a [torn-down actor] in a group" therefore holds only for
teardown; a merely dying, corpse-stage member is present and reachable, by both the population
checks and the order machinery, for the same window — until teardown, not until death.

Two lifetime rules go with it. An emptied group is **destroyed** when its owning player's
`+0x28` is zero (a human participant) and **kept** otherwise, so a scenario group survives its
last member and reads 0 — which is what makes a `== 0` comparison reachable at all. And check 1
has **no null path**: it dereferences its group unconditionally. What keeps that safe is the
binder, which never builds a check whose `Target_Group` failed to resolve; the trigger that names
the unbuilt check then resolves to **slot 0**, which is the hazard that replaces it.
