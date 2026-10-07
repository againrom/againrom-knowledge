# Group behavior, guard and patrol

[Reference](format.md)

## Update cadence

The AI is **one slot of the full tick**: `server+0x04 % 16 == 6`. It is driven **per group**, never
per actor — the driver walks the player list, then each player's group list, and runs the group
order machine once per group, skipping a group whose AI-enable byte is 0. At the shipped speed
index that is one AI pass per about 992 ms.

## Group and actor behavior

The group order at grpAI+0x20 is evaluated first. Eight values have live
arms. Parenthesized names are authoring-catalog labels; the behavior column
defines what each arm does.

| value | what the group does |
|---|---|
| 0 | run **each member's own state** — the only value under which `actor+0x50` matters |
| 1 | **guard**: build the list, clip it to the notice radius around the group centroid, score it, engage; a targetless member walks back to **its own post `ord+0x00`** (not the centroid) and, once there and idle, takes the idle-turn order `ord+0x08 = 0xb`. This is the order the load walk gives **every other** player's groups |
| 2 (`Swarm`) | build the list, score it, then per member: engage the scored target; else walk to the commanded cell `grpAI+0x0a` if not already standing on it and idle; else cast (scenario owner) or run the heal AI (a human participant's unit). **No formation and no spread** |
| 3 (`Stand Ground`) | **stand ground**: build the list with **no radius clip**, score it with the *reach-bounded* scorer, engage what is already in reach. The arm assigns no target past reach and issues no walk order. Read from the code and not observed at run time, the pursuit it stored at the previous evaluation keeps stepping until the next one reissues or replaces it, or until a stand or a refused route ends it. A human participant's member left without a target gets `ord+0x08 = 0`, or `8` when it casts a heal, and an AI owner's gets the idle turn `0xb`. The setters store `ord+0x08 = 0` for every member. This is the order the load walk gives the **player's own** groups. — AI-STAND-076, AI-REACH-072, AI-349, AI-351, AI-353 |
| 4 (`Move`) | walk each member to **its own** `ord+0x0a` — the arm ignores the cell the dispatcher passes it; the destination arrives from the command's setter. On arrival: stop, clear `mover+0x90`, set the stop distance to the member's reach, re-acquire |
| 5 (`Swarm 2`) | **arm 4 with a pre-emption.** The arm builds the candidate list and reads its element count, `AImanager+0xbb4`: **zero — nothing visible — and it tail-calls arm 4 with both arguments forwarded**; non-zero and it runs 2's body without the walk. The command's setter is arm 4's setter to the byte (one immediate apart), so the per-member state the fallback runs on is arm 4's own. **A vetoed group is not the same as a blind one:** candidates present but all vetoed by the preference matrix leaves every `ord+0x20` at 0 inside this arm, and both of its zero-target branches — idle turn plus optional cast, or the heal AI — **stand still** |
| 0x11 (`Roam`) | keep a Group-AI cell in `grpAI+0x0a`; re-roll when the farthest member's distance is **< 10** or the byte counter`+0x15` is **> 50**, by stepping **20 cells** in one of eight random compass directions and rejecting anything outside the playable rectangle; reset the counter on acceptance, call arm 5 evaluation and increment the counter after normal return |
| 0xff | forced when the group is empty: run the standing acquisition for each member |

Values 6..0x10 and 0x12..0xfe do nothing. **After every arm the withdraw tail runs**, below.

Roam's local result is the Group-AI cell/counter program and the Swarm2 call,
not a proved copy into member destinations. Swarm2 forwards the cell only on
its zero-candidate Move fallback; Move ignores that argument and uses each
member's current `ord+0x0a`. Reached movement can depend on member order and
called helpers. These bodies establish neither guaranteed wandering motion
nor absence of all other wander paths. (`AI-MOVE-023`'s own gloss for one field
this Move arm reads, `mover+0x90`, is superseded — see the occupied-destination
paragraph below.) — AI-ROAM-025, AI-SWARM2GATE-107, AI-MOVE-023

Group order 2 has three named producers: the script Swarm subcommand and
the routines at `L00518` and `L01700`. Callers of the latter two routines
remain Unknown. Pursuit, cast and pickup use the per-actor order field
`ord+0x08`; they are not values written to the group-order field by the
player-command or script-authoring paths. — AI-ORDER-031, AI-ORDER-294

When a Move member's own `ord+0x0a` destination is occupied by another actor
but not itself impassable, the arm's own arrival test is not reached true by
the ordinary path: the static search behind the walk is occupancy-blind and
resolves the route onto the literal occupied cell as it would any open one,
so the member's own cell never equals it and the same move re-issues every
tick — except in the race `MOVE-WAIT-008` names, where two units whose
dynamic routes were computed before either claimed the cell can both end up
standing on it, and the arm latches normally there. Nothing directly measures
how long the destination has been occupied; the only path that can still end
the wait is the walk arm's own periodic dynamic re-search (`R0055`),
which can raise the give-up flag, but only near a route's own end and only
when its own inner search comes back empty against the goal — whether that
happens for an isolated occupied-but-open cell is Unknown, not excluded. An
impassable destination reaches the static search's substitution step, which
tries a labelled cell near the goal. AI-335's universal nonempty-substitute
and only-picker-miss-collapse clauses are partially retracted. A substitute
equal to the seed leaves an empty route from an empty input list. The
selected zero-radius immediate walk caller then takes the same zero-count
branch as no substitute: resolved=current cell,+90=1,+98=1. A nonempty route
instead supplies its resolved end. Native later order continuation is
Unknown; these local flags do not establish next-tick arrival. A requested
seed can also exit the caller before search and preserve prior fields.
— AI-335, MOVE-072, MOVE-082, MOVE-083

The **actor state** (`actor+0x50`) is a 27-arm switch. **Its value at construction is `0xb`, guard.**
The arms a consumer needs:

| value | what the actor does |
|---|---|
| 0 | idle |
| 1, 2 | re-issue or complete a stored move |
| 3, 4 | engage a stored target, or the first hostile in the block around a stored cell |
| 8, 0x11 | follow an actor; the stop distance is `order+0x70`, or the actor's sight when that is 0 |
| **0xa** | **patrol** |
| **0xb** | **guard** |
| **0xc** | acquire (the decision above) with no post and no leash |
| 0xd, 0xe, 0xf, 0x17 | further order forms; `0x17` also overwrites both of the actor's radii with 5 |
| 0x16 | **explicit Retreat** — installed by the player command through a register-form store; repeatedly calls the first automatic-withdrawal helper, without that tail's HP-threshold admission (`AI-RETREAT-271`, `AI-RETREAT-274`) |
| anything else | clear the sub-tick gate and acquire |

## Group engagement

Everything here is what a group order arm runs before it touches a member; a consumer that gets the
per-actor states right and this wrong reproduces neither who is attacked nor when.

**1. One sight map for the whole group.** The builder clears the `0x10000`-byte visibility array
once, then stamps *every* member's sight into it, so a candidate seen by any member is a candidate
for all. For an AI-owned member the cell it was last struck from is forced into the map as well and
expires after **20** group ticks.

**2. The sweep is global.** Every actor on the map is tested against that map — head to tail of the
world actor list, no neighbourhood and no spatial index.

**3. One decider.** The diplomacy filter runs with the group's **first member** as `me`, so one
player row governs the whole group and one member's group list governs the See-invisible exception.

**4. Corpses are parked, not dropped.** Candidates below 1 health move to a second list; if the
first list ends up empty **the whole second list is moved back**. A group that can see only corpses
therefore has a candidate list and will attack one.

**5. Score, assign, and forget.** Each member is given the cheapest candidate in `ord+0x20`, or 0.
This happens on every evaluation: there is no memory of last tick's target anywhere. On an exact
cost tie the **first**-scanned candidate keeps the assignment — the opposite convention from
standing acquisition's own tie rule (`AI-340`).

**6. The list's element count is a decision input, and it is read at two widths.** Group order 5
branches on the full dword `+0xbb4`; the scorer tests only its low byte, as it does for the group's
own member count. A count that is a non-zero multiple of `0x100` therefore passes the gate and
scores nothing. Runtime reachability of this count boundary remains Unknown.

The cost, lower being better:

```
d    = Chebyshev(member, candidate) in cells
k    = candidate movement domain, forced to 0 when the candidate's reach > 1
pref = M[member domain][k]        for a member of reach <= 1   (melee)
     = M[0][k]                    for a member of reach >  1   (ranged)
d   := 1        when d <= member reach          (ranged members only)
     = d + 1 - member reach   otherwise         (ranged members only)
cost = (d << 8) + turnCost(member facing -> direction of candidate)
```

then `pref == 0` -> **never pick this candidate**; `pref == 1` -> `cost * 1.5`; `pref == 4` ->
`cost * 0.75`, the last two only when the candidate is single-cell with Mind >= 15. A candidate
carrying active spell id 20 costs a flat **+127**.

**M**, indexed `[member domain][candidate domain]`, is a compile-time constant:

| member \ candidate | 0 | 1 | 2 | 3 (flier) |
|---|---|---|---|---|
| 0 | 2 | 1 | 2 | 4 |
| **1 (ground)** | 4 | 2 | 1 | **0** |
| **2** | 2 | 1 | 4 | **0** |
| 3 (flier) | 2 | 1 | 2 | 4 |

Two consequences a consumer will not guess. **A melee ground creature never auto-selects a flier**
— and every human actor is domain 1, because the Humans table ships no `MovementType` cell and the
actor constructor defaults the field to 1. **A ranged creature uses row 0 always**, which contains
no zero, so reach and not class decides whether anything will chase a flier. Fliers are placed on
30 of the 38 EN maps.

Group order **3** uses a second scorer that additionally refuses any candidate whose distance term
exceeds 1 — that is the whole reason its members never move.

**Two clocks.** The group decides once per full tick; the order it writes is executed by the actor
tick, which runs far more often. Nothing in the group arms is sticky because nothing needs to be.

**The notice radius is frozen.** The group's geometry — centroid, spread, max sight — is recomputed
on every guard tick, but the only reader of the derived radius is the routine that *issues* a guard
order. The circle follows the group; its size is whatever it was when guard was last issued.

## Guard

**Where the post comes from is decided one level up, and it is decided twice.** A group under
group-order 1 or 3 never runs the per-actor state machine at all, so for it steps 1–4 below are
not the live path; what anchors its members is the **stance setter** that put the group under that
order. Every stance setter, before it writes the group order byte, walks the members and writes
each one's post: from the actor's current cell if it is standing on a cell centre, and from the
cell it is **stepping into** if it is mid-move. Only the two guard forms have that second branch;
the two stand-ground setters take the current cell unconditionally. At map load the placements are
all centred — every placement entry point of the position record writes the centre offsets — and
the spawner runs seventeen bytes before the stance walk, so a load-time post *is* the spawn cell.
Steps 1–4 are the live path only for a group at order **0**: one built by a player command, or —
because the load-time stance walk's session gate skips both setters together — every group on a
multiplayer load.

1. If the guard post is unset, **it becomes the actor's current cell**. Home is emergent: nothing
   authors it, and it is not the spawn point recorded in the map — it is wherever the actor
   happened to be when guard first ran.
2. If the actor is farther from the post than its **sight** radius, it is ordered to walk back.
   *This step is behaviourally dead*: step 3 overwrites it whenever it fires and step 4 writes
   exactly the same order. It survives only when the engage refuses and the acquire finds nothing,
   which needs a human owner.
3. The occupancy block of **Chebyshev radius 5** around the **post** — not around the actor — is
   filtered for hostiles and any survivor is engaged with a pursuit order. The 5 is `mover+0x08`,
   a **hard-coded constant** written once by the mover constructor; no shipped byte carries it and
   no class varies it.
4. With nothing to fight: away from the post, walk back; at the post, run the standing acquisition.

**Which survivor is engaged starts from distance to the post, not from any authoring or
diplomacy order, but "the first survivor" is a list head, not always the nearest.** The
occupancy-block builder visits candidate cells by increasing Chebyshev ring outward from the
post; within a ring the four edges interleave per cross-axis offset rather than running edge by
whole edge, so a smaller-radius candidate always precedes a larger-radius one, but among
candidates tied at the same ring distance the append order is the interleave, not a clean
south/north/east/west block (`AI-336`). A cell qualifies by the **static** plane's own bit-5
"a cell record exists" test, not a dynamic-plane occupancy read, and the enumerated actor is the
cell record's own `+4` slot. Neither the block builder nor its diplomacy filter tests health or
dying stage, so a body not yet torn down stays list-eligible; the selector itself
(`R0009`, read whole by `AI-332`) engages the list head with no health test on its own
default path, so a dying-but-not-torn-down candidate there is engaged ahead of a living hostile
placed later in the interleave — three named exceptions (a low-reach human-owned mage, a
class-spellbook slot draw, an AI-owned mage's own cast check) leave that default path (`AI-337`;
the slot draw is a creature's own `Spell 1..3`/`Probability 1..3` columns, not a target cache,
`AI-341`).

**This is what ends a pursuit.** Nothing inside a pursuit order measures elapsed time, health, or
the distance the pursuer has covered; the decision is re-taken from scratch every group tick, and
its only geometric input is *post to target*. Back a target out of the 5-cell block and the actor
is ordered home on the next tick, however far it has already chased; keep it inside and the actor
re-engages, however far it has already chased. **The post moves only when something re-issues a
stance.** A creature nobody commands keeps its spawn cell for the whole mission; a unit given
Guard again somewhere else is re-anchored there, and so is the 5-cell block. Besides the stance
setters, `ord+0x00` is written by the three target-died teardowns, the arrived-branch of the
order-progress epilogue, and patrol's re-anchor. A consumer must not model the post as immutable,
and must not model it as *"set on the first tick of guard"* either — for a load-time group it was
set before the first tick ran.

The second terminator is the route search: when it returns an empty path the mover sets
`mover+0x98`, and the order-progress epilogue cancels the order and re-acquires within **reach**.
A consumer that omits this will leave units frozen against unreachable targets. The pursuit's own
walk routine has no `MOVE-WAIT-008`-style wait state for a step refused only by another body
standing on the destination cell — three named exits skip its own stepper call instead (not
centred; already within stop distance; the current route already matches both stored endpoints) —
and its own direct static-search failure setter fires only on a total search failure, the same
self-collapse the Move arm's own occupied-destination paragraph above describes. But the same
routine also calls the dynamic-plane re-search that can raise `mover+0x98` through the gate that
paragraph names, aimed this time at the target's own contact ring rather than a single cell, so
"never on mere occupancy" does not hold in general: what a pursuit held because every approach to
the target is occupied (an escorted target, for instance) actually does is Unknown, not identical
to an ordinary in-progress walk (`MOVE-073`).

This actor guard path does not define the Group Roam program above. An unordered actor stands on its post
until something hostile comes into the block around that post. Guard is the whole of motion for a
group left at its **load-time** order — it is not the whole of motion without a *player* order,
because the map's script can put a group on patrol (below).

Two fields of the group AI record belong to this arm and to nothing else: `grpAI+0x30` is a
**has-members latch** and `grpAI+0x34` a **tick counter with no reader**. The latch is what gates
the radius roll below.

## Patrol

State `0xa`. The actor **runs guard first**, so a patrolling actor still fights; if guard leaves it
idle, the patrol advances. The waypoints are a **ring**: on arrival at the current one the list is
searched for that cell and the next node is taken, falling back to the head.

**Two live surfaces put an actor into this state, and the one shipped maps use is the map's own
script.** The `.alm` type-7 action opcode 6 is a group command; its first parameter `14` is the
shipped catalogue's `Group Command : Patrol`, taking `X`, `Y` and a target group. The command:

1. stops every member and sets the **group order to 0**, which is what lets the per-actor state
   machine run at all;
2. per member: `actor+0x50 = 0xa`, guard post := the member's current cell, and the waypoint list
   is emptied and rebuilt as **exactly two nodes — the member's own current cell, then `(Y<<8)|X`**;
3. the current waypoint is set to the commanded cell.

So a shipped patrol is one commanded point and wherever the creature was standing, walked back and
forth forever. `order+0x04` is a **re-anchor latch**: set on every advance and consumed at the next
entry to move the guard post to the actor's current cell — which is why the guard leash in step 2 of
the previous section never pulls a patrolling actor home.

The same per-member routine serves a **player-issued** patrol command, a sibling of the guard and
aggressive commands in the same dispatcher.

The other surface is the **mission description file** `World\Mission\<n>.ini`, a section of `x;y`
lines with both coordinates strictly inside 8..135; a line without a `;` at index 1 or above is
skipped with an operator message, and a group whose path came out empty is not put into patrol. It
copies the group's path into each member and adds that member's own cell to it. **No such file ships
in the GOG install** — 0 nodes in any archive, 0 files under any `Mission` directory in either
preserved root — so no shipped patrol comes from *there*.

A consumer implementing the ring should note that the engine's own search dereferences null when the
current waypoint is not a member of the list. No live setter can produce that state.

## Engagement radius

Reach decides the pick and sight decides the population (above). The **notice radius is a third
thing and it belongs to the group**:

```
centroid   = mean of the members' fine positions
grpAI+0x2c = max over members of ( Chebyshev(centroid, member) + member.sight )
grpAI+0x2d = grpAI+0x38 = max( grpAI+0x2c, [Scanning] MinimalGuardRange )
on the tick grpAI+0x30 flips from 0:
           grpAI+0x2d = grpAI+0x38 + 4 + rand3
           rand3      = (3 * rand()) / 0x8000 - 1,  i.e. exactly {-1, 0, +1},
                        each on about a third of the RNG's range
           (the +4 is on that edge only; the empty-group edge rolls without it)
```

`0x8000` is not a literal in the roll: it is `AImgr+0x00`, written once by the AI manager's base
constructor and never again, and it is `RAND_MAX + 1`. **`n * rand() / AImgr[0]` is the module's
uniform-range idiom** — fifteen sites use it — so any such quotient is `floor(n·rand()/32768)`,
uniform on `0…n−1`. The idle-turn arm is the same idiom at `n = 190`.

The radius clips the group's **candidate list** around the centroid under group order 1. It is
never authored, and it never decides whether a blow lands.
