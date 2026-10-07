# Saved ownership and Group state

[Reference](format.md)

## Saved formation ownership

Formation LOAD copies the block into the current serialized Player's +30
allocation. Archive object index, traversal position, saved identity key,
word identifier+04 and the dword identifier+08 are separate identities.
Group+44 resolves its own saved key immediately; a key registered only later
is unavailable at that read. Later actor-owner correction does not recopy
Group+44. — SAV-PLAYERIDENT-830

Both Group movement reads use `Group+44 -> Player+30 -> byte+1f`.
The client formation command instead carries the active CPlayer's +04 word;
the simulation finds the first matching saved Player+04 and writes that
object. Trigger formation binding uses a temporary map keyed by Player+08,
then writes through the compiled record's resolved Player pointer. These
selectors can name different Players. — AI-FORMOWNER-314, AI-FORMCMD-315,
AI-FORMTRIGGER-316

The selected existing-player return path matches a name and sends that
Player's description before the others. The client active-pointer store is
conditional on its current Player array size being 1. Full incoming-name
selection, initial client-array state, delivery and first post-LOAD chronology
remain Unknown; there is no proved universal first-SAV-record selector.
— AI-FORMACTIVE-317

The measured snapshot contains 129 paths, 76 contents and 75 complete-reader
acceptances. Its 262 Player records have traversal equal to both identifiers;
all 764 Group owner keys agree with the enclosing Player. That agreement does
not distinguish these selectors, and includes preserved generated diagnostics.
One invalid-header path is an explicit refusal. — SAV-PLAYERPOP-831

## Saved Group continuation

Group LOAD restores the 80-byte AI block at `*(Group+3c)`, then replaces its serialized +4c
pointer with a fresh word list and restores that list. Constructor order 0/enabled1 are not a
replacement for saved AI values. The normal resume linker receives argument 0; its initial-stance
branch is skipped, but its Group+1c index insertion is not. Saved +1c remains a literal selector,
distinct from file-local object pointer keys. Dynamic first-save initialization remains Unknown.
— SAV-GRPIDENT-562, SAV-GRPAI-563

This is not whole-linker or first-tick preservation. A selected later link arm can write AI+48
through a resolved actor's Group. Group tick reads membership count before order+20 and writes
ff when empty. Arm 0 reads actor ownership, dispatches the actor, then searches the live member
list before advancing. The rate helper follows actor+70 to Group+3c and reads AI byte+44.
Transitive effects and first reached chronology remain Unknown. — SAV-GRPAI-563

One Group dispatch selects order+20 once. Arms 2/4/5 receive AI word+0a;
changing the byte in a returning callee does not restart selection. Every arm,
including an unrecognized byte, reaches a fresh current-membership withdraw
tail. The empty-group ff store is at entry only: becoming empty during an arm
does not itself rewrite the order. — SAV-GRPDISPATCH-568

Order 0's inline successor search and the helper used by order 3, ff and the
withdraw tail find the just-dispatched actor again in the current list. Missing
actor or final node ends that stage; a new successor can extend it. The tail
starts again at the current head. Actual callback mutations and deallocation
safety remain Unknown. — SAV-GRPMUTATE-569

The AI-owned path is not the actor's active patrol ring. The path-copy setter
walks AI+4c head-to-tail but prepends every point into actor order+90. Source
A,B,C becomes own,C,B,A; current order+02 is C, chosen before prepending own.
This is a setter rule, not a normal-LOAD reinitialization rule.
— SAV-GRPPATROL-570

Actor-order LOAD preserves raw cursor+02 and latch+04 while replacing list
pointer+90. Patrol consumes a nonzero latch before guard, then only guard
result 0/b reaches continuation. That continuation sets the latch even before
arrival. Arrival searches for the first equal cell value, advances or wraps;
missing value reaches an unguarded null read. The cursor is not a node index.
— SAV-PATROLCURSOR-571

Group SAVE consumes its current embedded word list, AI block/list, members and
tail fields. No direct Group-dispatch instruction dereferences either Group
word-list payload; this is not a whole-call-tree absence. Group+20 semantics,
unnamed AI fields and original LOAD-to-dispatch-to-SAVE chronology remain
Unknown. — SAV-GRPSAVENEXT-572

## Scope

| | owner |
|---|---|
| the blow itself — timer, roll, damage, experience | [`claims/hero.md`](../../claims/hero.md), [`claims/unit.md`](../../claims/unit.md) |
| walking there, pathing, tick order | [`formats/move`](../move/format.md) |
| the block/cost planes and the cell-record hash | [`formats/terrain`](../terrain/format.md) |
| the `.alm` type-5 player record's own layout | [`formats/alm`](../alm/format.md) |
| **this file** — the decision, from "nobody told me" to "that one" | — |

## Object state

| | what |
|---|---|
| **session** | `[L00004]`, `0xc320` bytes. `+0xa50` world, `+0xa9b8` `MinimalGuardRange`, `+0xa9c4` diplomacy, `+0xba4` / `+0xbc4` the live / dead candidate lists — each an outer object whose list sub-object begins **four bytes in**, at `+0xba8` / `+0xbc8`, so their element counts are `+0xbb4` / `+0xbd4` (`list + 0xc`). The pair is **shared scratch**: per-actor acquisition rebuilds it too, so nothing in it survives another routine's call |
| **world** | `0xa4558` bytes. `+0x10000` block plane, `+0x5400c` the occupancy scratch list, `+0x540b8` the cell-record hash, `+0x82ef0` the visibility map, `+0x9451c` heights, `+0xa4554` the actor list |
| **actor** | `+0x10` position, `+0x14` `Player`, `+0x4c` flags (bit 3 = off map), `+0x5c` combat target, `+0x94`/`+0x96` health, `+0xa4` sight (`u16`, 1/256 cell; `+0xa5` is its whole-cell high byte), `+0x12c` reach, `+0x144` state bits (`0x8000` invisible), `+0x154` mover, `+0x158` order block |
| **order block** | `+0x00` guard post cell, `+0x02` current patrol waypoint, `+0x08` state, `+0x0c` target, `+0x14` stop distance, `+0x20` the group's assigned target, `+0x58`/`+0x5a` remembered attacker cell and its age, `+0x70` follow stop distance, `+0x71` see-invisible radius, `+0x90` the patrol waypoint list |
| **group** | `0x48` bytes, a `CObList` (`+0x04` head, `+0x0c` count) in `player+0x24`. `+0x1c` the authored group id, `+0x3c` the AI record, `+0x44` the owning `Player` |
| **group AI record** | `0x50` bytes. `+0x00` current guard post, `+0x20` the group order, `+0x24`/`+0x28` fine and cell centroid, `+0x2a` spread, `+0x2b` max member sight, `+0x2c` derived notice radius, `+0x2d` working radius, `+0x38` base radius, `+0x45` activity count used as a nonzero gate (constructor 1; refresh below), `+0x48` force-activity dword, `+0x4c` the patrol path |


## Spatial activity refresh

Before Group dispatch, the regular driver callsL00164 when receiver+b388 is
zero. That session constructor ends with the override zero. The complete
rebuild consumes its registered actor list twice: Player+28==0 actors mark
coverage; other actors first reset their Group AI+45 and then increment it when
covered or when the Group AI dword+48 is nonzero. Coverage is a five-by-five
area in eight-map-cell blocks, with the four corners excluded. The Group byte
counts represented active members, wraps at 256, and is rebuilt on the next
call. Groups absent from the supplied list are not reset. This is a proximity
activation mechanism distinct from target acquisition and line of sight.
— AI-ACTIVITY-324

The named bodies establish the refresh and its driver position, not every
actor-list enrollment/removal path, the first native post-LOAD event, or
acceptance of source coordinates beyond the measured safe grid.
— AI-ACTIVITY-324
