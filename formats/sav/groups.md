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

The selected dispatcher does not directly read Group `+20` elements or AI
`+4c` elements. Transitive callbacks, genuine first-SAVE values and a graph-wide
atomic snapshot remain Unknown; skipping the initial-stance branch proves no
whole-linker or first-tick preservation. — SAV-GRPAI-563,
SAV-GRPSAVENEXT-572, SAV-GRPFIRSTSAVE-579
