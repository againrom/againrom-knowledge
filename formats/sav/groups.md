# SAV Groups and AI

[Format reference](format.md) · [Player](player.md) · [Actors](actors.md)

## Wire programme

A Group is serialized directly inside its Player's u32-counted Group list;
it has no archive class tag or saved-address definition of its own.
— SAV-OBJ-016, SAV-MEMBER-036 (its unread-classes Unknown superseded by
SAV-EMBED-039)

| Order | Wire | Original source |
|---:|---|---|
| 1 | `Count(n)`, `n*u16` | Embedded list at Group `+20` |
| 2 | raw 80 | AI object pointed to by Group `+3c` |
| 3 | `Count(n)`, `n*u16` | List pointed to by AI `+4c` |
| 4 | u32 count, typed Unit references | Group actor list |
| 5 | u32 | Group `+1c`, literal selector |
| 6 | u32 | Group `+40`, remapped key |
| 7 | u32 | Group `+44`, remapped owner key |

Byte `0x4c` of the AI raw80 is the AI word-list pointer. STORE writes its live
bits; LOAD replaces it with a fresh list before reading the list. Any saved
value, including 0, loads alike. — SAV-1113

The AI constructor allocates that `0x1c`-byte list and stores its address, or
0 when the allocation fails. STORE dispatches through the word, so a completed
STORE saw a heap address. No other direct store nulls it in the ranges scanned
(Medium). It is the object reference a Player record carries in each Group's
AI raw80. — SAV-1391

Members are written head-to-tail and loaded with append. Changing Group
detaches an actor first and appends it anew. List position does not identify
an actor; a Group can split or combine without changing its actors' keys.
— SAV-EMBED-039, SAV-WLIST-040, SAV-GRPORD-058, SAV-GRPFLD-060

## LOAD order and ownership

| Step | Local action |
|---:|---|
| 1 | Register the enclosing Player's saved key before its Groups |
| 2 | Construct a fresh Group |
| 3 | Restore Group list and AI80/list; AI LOAD replaces its saved `+4c` pointer with a fresh word list |
| 4 | Read each Unit reference, detach its old Group, append, stamp actor `+70`, copy actor `+14` to Group `+44` |
| 5 | Restore literal Group `+1c`; immediately remap `+40` and `+44`, nulling missing keys |
| 6 | Append the completed Group to Player `+24` |
| 7 | After settings and Diary, remap Player's hero, flatten Group members into Player `+20` and stamp actor `+14` with the enclosing Player |

Step 7 does not recopy Group `+44` or actor `+70`. Group owner and enclosing
Player can therefore differ at that local boundary. A later Player key
registration does not retroactively repair an earlier Group lookup.
— SAV-GRPLOAD-560, SAV-GRPOWNER-561, SAV-PLAYERIDENT-830

## Construction and selector

The 0x48-byte Group constructor preserves incoming `+1c`, builds the list at
`+20`, supplies `+3c`, and clears `+40/+44`. The list has its vtable at `+20`,
zero fields at `+24,+28,+2c,+30,+34` and grow-by 10 at `+38`. Neither the
selected base constructors nor the ordinary-command prologue supply a `+1c`
default. — SAV-GRPNEW-576, SAV-GRPCMD-578

A fresh allocation does not establish zero for preserved `+1c`. Its dynamic
first-SAVE producer is not fully specified. — SAV-GRPALLOC-577,
SAV-GRPFIRSTSAVE-579

An authored-map miss sets `+1c` from loader `+40`. LOAD restores the saved
selector and normal resume indexes Groups by it before its initial-stance
gate. The field is not a remapped pointer, and an uninitialized dynamic value
is not established as harmless. The ordinary wrapper deletes at most the
first rejected Group; Guard and its selected summary write AI state through
`+3c`, not Group `+1c/+20`. — SAV-GRPIDENT-562, SAV-GRPCMD-578,
SAV-GRPFIRSTSAVE-579

Group `+40` has a known remap rule but Unknown gameplay meaning. Group `+20`
has known framing but Unknown element meaning and ordinary mutation/consumer
rules outside its constructor and LOAD, including computed access.
— SAV-GRPFLD-060, SAV-GRPLIST-807

An exhaustive direct-call census of the shared list class's two mutation entry
points (18 sites) finds none targeting Group `+20` specifically; every site
targets a different field of a different object. Only the constructor and
`Group::Serialize` touch `+20` directly — construction and archive round-trip,
not population. 0 of 567 corpus Group records ever carry a nonzero `+20`
element count, so this framework has no populated instance to read element
meaning from. — SAV-1107

## Dispatch and patrol state

| State | Local consumer or producer |
|---|---|
| AI order byte `+20` | Empty Group sets it to `ff` at dispatch entry; nonempty dispatch reads current order once |
| AI word `+0a` | Passed to order arms 2/4/5 |
| AI `+44` | Rate helper reads it through actor `+70 -> Group+3c` |
| AI `+48` | A later conditional linker arm writes it |
| AI `+4c` path | Path-copy setter reverses it into actor-order `+90`, selects head as cursor `+02`, then prepends actor's cell |

All order bytes reach the shared withdraw tail, which rereads count/head.
Becoming empty after dispatch entry does not reset the byte in that invocation.
Order 0, order 3, `ff` and withdraw re-find the dispatched actor in the current
list before advancing: a missing actor ends the stage; a newly appended
successor can be visited. Withdraw restarts at the current head even if the
order stage ended early. — SAV-GRPAI-563, SAV-GRPDISPATCH-568,
SAV-GRPMUTATE-569

Named ordinary producers can select `0,1,2,3,4,5,0x11,0xff`. The archive
restores the order byte literally. — SAV-GRPORDER-806

Actor-order LOAD restores its raw 148 bytes and replaces `+90` with a fresh
list. It does not replace saved cursor/path values with the Group path.
On arrival the walker finds the first equal cursor cell, then advances or
wraps; a missing value has an unguarded null edge. The `+04` re-anchor latch
is set whenever Guard admits patrol continuation. — SAV-GRPPATROL-570,
SAV-PATROLCURSOR-571

## SAVE and Unknowns

SAVE emits current list values, AI80/list, actor order and trailing selectors
in the wire order above. List-node identities are not saved. AI/order raw
blocks carry current allocation pointers which LOAD replaces again. The
Group list store precedes member serialization; the selector read follows
it. Callback mutation can therefore affect different fields at different
points in one SAVE. — SAV-GRPSAVENEXT-572, SAV-GRPFIRSTSAVE-579

## Restart-slot values

All ten restart-slot saves of the owner root hold sub-tick 1, full tick 0.
— SAV-1221

In them a map Group of Player 2 or above holds order 1 and the guard setter's
block: centroid `+24`/`+28` and `+00` and spread `+2a` recomputed from its
members' cells, sight `+2b` and notice `+2c` fitting their bounds, `+2d` = `+38` = max(`+2c`, 8), `+39` 0, `+45` 1,
`+48` 0 or 1, every other byte 0 and both lists empty. — SAV-1218, AI-450,
AI-452

Every Player-1 Group holds order 3. A map Group of Player 1 holds only `+20`
3 and `+45` 1. In 12 carried Groups of 9 distinct files, 7 hold AI bytes (`+0a`,
`+24`..`+2c`, `+44`) that no store read writes, from about four independent
lineages; `+45` is 0 on two. Their producer is Unknown, and whether they predate
the mission start is Medium (the carry arm is unread). A carried Group can hold
several carried members. — SAV-1219, AI-449, AI-451 The unconditional map admission reading of `AI-449` is superseded by `AI-476`; constructor facts and matched-record membership stand.

Every member's order block holds `+00` = its cell word, `+08` = `+09` = 0,
an empty `+90` list, `+14` and `+71`; a Unit adds `+40`/`+44` and up to one
spell slot. — SAV-1220

The selected dispatcher does not directly read Group `+20` elements or AI
`+4c` elements. Transitive callbacks, genuine first-SAVE values and a graph-wide
atomic snapshot remain Unknown; skipping the initial-stance branch proves no
whole-linker or first-tick preservation. — SAV-GRPAI-563,
SAV-GRPSAVENEXT-572, SAV-GRPFIRSTSAVE-579

## Join membership and empty Groups

Owner-transfer instants 19/22 give each transferred actor a separate fresh
Group. A successful town AddHero grant also receives a fresh Group per actor.
The command-0x38 hire loop creates one Group per nonzero type and appends all
actors of that type to it. Catapult and Ballista share that Group path.
— PARTY-039, PARTY-040

The Player Group archive writes every currently listed Group, including an
empty one, and LOAD reconstructs that partition in order. The named loops do
not merge Groups or renumber the literal +1c selector. Runtime allocation
addresses are new on LOAD. — SAV-1234

At the next join walk, the read collection helpers find the selected head,
repair links and decrement count. The prefix tests zero head members and list
count above one. With no transitive list mutation, it removes the leading empty
prefix until a nonempty head or one Group remains. Later empty Groups behind a
nonempty head are not visited by this prefix. Local selection/unlink is High;
post-callback progression remains Medium. AI-449's map-admission clause is
superseded; its constructor facts stand. — AI-449, AI-479, AI-628, AI-629

Player removal requests Group deletion after count zero only when the Player
reached through Group+44 has +28 zero. Teardown and the mission-end discard arm
reach that path. Ordinary member detach clears actor+70; the read owner-transfer
and tavern rebuild paths have no local empty-Group destruction arm. Retention
through their unread callbacks is Medium. — PARTY-041, PARTY-042

The file writer invokes live empty-Group pruning for selected Player+28-zero
lists before archive construction. Its iterator saves the successor before
returning the current Group. This is separate from the inner serializers'
inclusion of every currently listed Group, including empty ones. Final saved
population and full town chronology remain bounded. — SAV-1234, SAV-1251

The named tavern handler commits Groups and members before one debit. Rebuild
credits the last tracked cost before detaching marked mercenaries; same-visit
undo activation and source-Group retention through callbacks remain Unknown.
— PARTY-043

The three raw sections' rel32/dword target scans give bounded caller candidates,
not a complete unlink/free census. The binder caller and transitive clear paths
remain unread. — AI-630

No explicit upper-count rejection was found in the named append/archive bodies
and forwarding wrappers. Pool allocation, actual archive primitives, overflow
and a program-wide Group/member limit remain Unknown. This negative is Medium
over the read bodies. — SAV-1252

## Selector binding after LOAD

World LOAD empties the Player list and appends the saved Players in file
order; each Player's Groups follow in file order. The map reload and the
selector binder run after that, so the binder sees every saved Group,
including Groups created during play and emptied Groups. — SAV-1337, SAV-1234

After LOAD a script Group id binds to the last saved Group carrying that `+1c`,
Players in file order, then Groups in file order; the owner is not part of the
key. Two Groups of one Player with one id: the later one. Medium: the binder
and the local append order are High, and that nothing between the Player-list
LOAD and the binder reorders the lists is Medium. — SAV-1338

The owner-transfer, town AddHero and hire bodies store no `+1c` after Group
construction. In 25 original mission saves 43 of the 45 human-Player Groups
carry a `+1c` that no map authors (15 zero, 1 small, 27 at least 256). The
census does not identify the creator of those 43 Groups; that they hold
allocation residue is an inference, Medium. — SAV-1339

In those 25 saves every scripted id binds after LOAD to the owner the map
start binds; 38 of 164 bind to an empty Group. No unauthored Group carries a
scripted id of its map. Each scripted id has exactly one carrier, so this does
not separate last-wins from first-wins. — SAV-1340

In 27 complete walks of original-authored saves no Group of a `+28`-zero
Player is empty (49 Groups);
`+28`-one Players hold 51 empty Groups of 405, 20 behind a nonempty Group, in
25 mission saves. A truncated walk is a further witness. The corpus contains
no town save with an AI Player or an empty Group. — SAV-1341

The human Player is the first Player (slot 1, id 1, `+28` zero) and holds the
hero; no member in 27 complete walks (1,004) carries another owner key than
its enclosing Player's. A member
saved in another Player's Group is restored as that Player's actor. —
PARTY-047
