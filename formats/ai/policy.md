# Behavior selection and difficulty

[Reference](format.md)

## Behavior selection

Nothing in the `.alm` authors a behaviour **per unit**; everything below names a **group**.

### At load

The map contributes two things: the **group partition** (the type-6 group id, which puts units of
one player into one group) and the **roster order**. Two setters exist:

| setter | members | post it anchors | group order |
|---|---|---|---|
| guard | `actor+0x50 = 0xb` | current cell, or the cell being stepped into | 1 |
| aggressive / stand ground | `actor+0x50 = 0xc` | current cell, always | 3 |

Both also clear each member's order byte, so a stance issued mid-order cancels it.

On the map-load path, **in single player only**, the groups of the player in **type-5 slot 0** are
set aggressive and every other player's groups are set to guard. Installed EN
slot 0 names are Self/self. The single-player test and the argument test
gate **both** branches together: fail either and no group is set, so a multiplayer load leaves
every group at the constructor's `0` and hands every member to its own state machine — where the
per-actor guard arm anchors the post instead, at the first tick.

The mission description file can override this per group: key `Patrol` names a waypoint section, and
key `StandGround = Yes` selects the aggressive setter. An absent key means guard. No such file ships.

### At run time — the map's own script

Type7 action opcode 6 dispatches group commands on its first parameter. Eleven
signatures are authorable; Dwell has no executable case:

| par0 | Command | Resulting group order |
|---:|---|---|
| 1 | Guard | 1, via the load-time guard setter |
| 2 | Swarm | 2, with the target cell in `grpAI+0x0a` |
| 3 | Stand Ground | 3 |
| 4 | Move | 4 |
| 5 | Swarm 2 | 5 |
| 10 | Attack | 0 — per member, engage the named unit |
| 11 | Defend | 0 — per member, the follow states |
| 14 | Patrol | 0 — per member, state `0xa` (above) |
| 15 | Follow | 0 |
| 17 | Roam | `0x11` |
| 18 | Dwell | **no case — authorable and inert** |

A subcommand-10 attack is not a group-wide copy of one order. The helper first clears each
member's route and reservation state, then forks on identity: the named member enters state `0xc`
with its post set to its own cell and no pending order; every other member enters state 3 with the
named member at `ord+0x0c`, its own `actor+0x12c` at `ord+0x14`, and no pending order. The next
member-state pass therefore acquires on the subject and engages on everyone else
(`AI-SCRIPTATTACK-120`, whose headline universality and single-walk dispatcher reading are
superseded; the two helper branches stated here stand). The dispatcher walks the group twice.
Before the per-member call it compares the target cost of member and named unit with the sentinel
`0xffffff` (`AI-COST-071`, reached only through a zero cell of `AI-PREF-070`'s matrix); on the
sentinel it takes the named-member branch (`R0088`: state `0xc`, post its own cell, no
pending order) instead of the other-member branch, so that member acquires in place and never
engages the named unit. Every shipped `Victim` reference is a Humans unit in movement domain 1,
so no shipped node reaches it.

Subcommand 11 is the defend form already specified by `AI-DEFEND-111` and
`AI-FOLLOWSET-116`: the named member acquires, every other member stores it at `ord+0x10` and
defends it at the authored range. Subcommand 15 is the related follow form, which keeps the range but omits defending on the
subject's behalf.

A group's order is therefore **not** its load-time 1 or 3 for the whole mission: it is whatever the
last command left. The four commands that leave 0 are the ones that hand the group back to its
members' own states.

### Break-off rules

The current group order selects the break-off rule. Script commands can
replace the load-time order, and player-issued orders allocate a fresh group
at order 0. — AI-CENSUS-047

| Group order | Break-off rule |
|---:|---|
| 0 | Chebyshev5 about the actor post, ord+0x00 |
| 1 | grpAI+0x2d about group center, grpAI+0x28 |
| 2 | Swarm; no complete clip rule specified |
| 3 | No break-off |
| 4 | Move; no complete clip rule specified |
| 5 | Swarm2; see group-state rules |

## Ranged withdrawal

Once per full tick, for **every member of every group and after whatever the order arm did**, the
dispatcher runs two tests in this order and takes the first that fires:

| test | radius | threshold |
|---|---|---|
| 1 | the actor's own `pth+0x08` | `health <= ord+0x44` — the `Data.bin` Units column **`Wimpy`** |
| 2 | the literal **2** cells | `health <= ord+0x40` — the `Data.bin` Units column **`Withdraw`** |

Both collect the occupants of the square of that radius, keep the ones the diplomacy matrix calls
hostile, and require the set to be non-empty. Both thresholds compare against **current** health as
a signed word, so a value of 0 means *never* and a value equal to `healthMax` means *always*.

The first reaction helper counts entries with `health > 0` in a byte. A zero count calls
ordinary acquisition. Its averaging loop then sums fine coordinates over the **whole filtered
list**, not a living-only subset, and divides by the low byte of the list count. The second
helper has a nonempty-list gate without a positive-HP scan. Their actual mixed/all-dead
runtime populations and dense-count wrap remain Unknown (`AI-RETREAT-274`).

Both call the same distance-3 away picker and install pending move 1. The picker clamps to
`[8, dimension-9]`; a zero coordinate difference becomes 1. It has no passability test.
Route search, not the picker, handles inaccessible cells. These geometry clauses survive
the partial correction to `AI-WITHDRAW-028`, whose living-only centroid clause is retracted.

**Shipped threshold source, not explicit-command eligibility.** The **Humans** table has
neither column; that is not an immunity to explicit Retreat. Of the 56 parameterised Units rows in
both shipped roots, twelve carry a nonzero `Withdraw` and they are exactly the three ranged
families — `Goblin_Sling` (30), `Orc_Bow` (60), `Bat_Sonic` (15), four tiers each. Because the
threshold is absolute health and the four tiers of a family share the tier-1 value while `healthMax`
scales ×1.6 per tier, a **tier-1 ranged monster withdraws on every contact** and tiers 2/3/4 only
below 62 % / 39 % / 24 % of their own maximum. Twenty-four rows carry a nonzero `Wimpy`
(`Goblin_Pike` 10, `Goblin_Sling` 8, `Bee` 3, `Squirrel` 5, `Bat_Sonic` 4, `Dragon` 63).

A spawn-time classifier can
replace both numbers with `healthMax`, `healthMax/2` or `healthMax/4` according to a shooter/caster
class byte. It **never executes those writes**: its only call site is inside the base actor
constructor, before the derived vtable is installed, so its `is-humanoid` gate always reads 0.
The same class byte selects the per-class
percentages of the runtime command that resets both thresholds (three modes: off, 10 %, 30 %); since
nothing ever sets the byte, only the first column of that table is reachable.

## Difficulty

The three-way pre-create control does **not** reach behaviour. No instruction in the AI module reads
the state it sets; `ai.reg` has no per-level dimension; group counts come from the map, the cadence
from the speed index, and the radii from geometry. The level's only consumer is the spawner's
health / to-hit / defence adjustment ([UNIT](../unit/format.md)).

## Cost

The candidate sweep is **linear in the number of on-map actors**, twice: once to stamp
visibility, once to test each actor's cell. There is no spatial index anywhere on the path, and
the visibility map is cleared and rebuilt per decision.
