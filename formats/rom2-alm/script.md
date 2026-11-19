# ROM2 campaign scripts

Both preserved campaign corpora contain 46 ALM maps. Type 7 is three counted
arrays: action nodes of 796 bytes, check nodes of 796 bytes and triggers of
184 bytes. Per root there are 1334 actions and 986 checks. Node DWORD+0x40 is
the full operation, +0x44 its ID, +0x48 a separate word, +0x4c the ten DWORD
parameters and +0x74 their ten DWORD type codes. Actions and checks have
separate ID namespaces. — R2-ENGINE-041

## Compilation and state owners

Action word 65538 appends low8(parameter 0) | low8(parameter 1)<<8 to a compiler
u16 collection and creates no executable action. Its downstream consumer is
Unknown. Check word 65538 initializes a local check-result register with
parameter 0 and creates no executable check. Mission10's action pair is 11/41;
its special checks initialize 0/4/1/2/3/1229. These full words must not be
reduced to ordinary opcode 2. — R2-ENGINE-042, R2-ENGINE-043

The bounded postcompile trace still leaves the packed collection's consumer
Unknown. The read caller invokes the type-9 stage and a conditional
owner/group pass; neither read establishes what consumes that collection.
This does not establish that it is unused. — R2-ENGINE-1317

Ordinary check IR has result slot+0, operation+4 and scalar 0 at+8. Literal
nonreference values pack densely, including trailing scalar zeros. Literal
references resolve separately: first unit pointer+0x5c, group+0x60,
player+0x64 and second reference+0x68. Literal Target_Item becomes
low16(value+3608) at+0x6c. Failed literal reference resolution can clear the
ID mapping and omit executable IR. Dynamic references in 8000..9999 use a
different metadata path; its complete resolution/error contract is Unknown.
— R2-ENGINE-044

Literal Unit references 10001..11000 take a separate EN/RU native path with
offset reference-10001. It requires Session +0x74==0 and nonnull selected
owner +0x38, then compares actor WORD(+0x14c)-21 in that owner's actor list.
Thus literal 10001 compares that word to 21, and 10002 to 22. Actor presence,
uniqueness and complete alternate-list coverage remain Unknown. References
below 10001 and above 11000 take different native paths. These numeric
identities are separate from runtime entity keys. — R2-ENGINE-107

The selected map actor builder matches signed row WORD(+0x14) to owner
DWORD(+8), which the owner builder copies from owner-row DWORD(+4). Generated
owner WORD(+4) is a separate coordinate. Successful registration assigns the
actor's owner pointer and groups actors by actor-row DWORD(+0x40). Template
contents, skipped rows and complete entering membership remain Unknown.
— R2-ENGINE-108

Literal group compilation visits the current owner list, then each owner's
group list, storing each group pointer under its complete DWORD number.
Duplicate numbers select the last visited group. The selected owner
registration appends owners; every other list/load producer's order remains
Unknown. An owner-key maximum is not the lookup contract.
— R2-ENGINE-085

DLL scenario variables are distinct from map result registers at
Session+0xc784+4*slot. ScenarioGetVar/SetVar use static DWORD base+4*index
with no explicit guard in the verified export bodies. NewGame initially clears
1024 DWORDs and later explicitly sets slot 768=10. Safe external index bounds
and all final defaults are Unknown. — R2-ENGINE-045

## Added declared arms

Action 10 dispatches to the one-direction relation writer in the body read
(RU L2.05822 / EN L2.05823). With compiled
scalars p0/p1/p2, it stores M[p0][p1] = old & 0xfc, then adds low8(p2)
modulo 256. For p2 in 0..3 the old bits 2..7 survive. There is no local
lock-bit test and no direct reciprocal-cell store in that body. The action then invokes
the relation packet producer for returned owner entries matching either
endpoint. Complete iteration, queue delivery and effects of unread callees
remain Medium or Unknown. — R2-ENGINE-979, R2-ENGINE-980

In mission 10, trigger index 3 Ogre annoyed names action 19, opcode 10 with
(7,1,1), type 1 for each operand. The two whole map members are byte-equal.
Stored relation defaults between raw owner rows 1 and 7 are zero. Composing
their native identity and literal action compilation is Medium: the script
can produce M[7][1]=1 while M[1][7]=0. Owner 1's outgoing hostility test
then stays clear until an admitted hit (F=0, or the F nonzero gates) or
another writer changes it. The owner identity of the hero and ogre units and
the trigger's firing condition are Unknown. Runtime cursor, click, delivery
and the hero's next action are Unknown. — R2-ENGINE-983,
R2-ENGINE-984

Scalar indices below are compiled indices, not original catalog slot indices.
The meanings are scoped to the verified RU/EN executor paths.

| Arm | Literal operands | Actual effect or predicate | Authority |
|---|---|---|---|
| Instant 5 | scalar 0=reason | Session +0xbe14 = reason | R2-ENGINE-623 |
| Check 18 | resolved unit | +0xbe14 + 1 while unit +0x54 = 0x10; no check result | R2-ENGINE-627 |
| Instant 35 | scalar 0=index, scalar 1=value | SetVar(index,value) | R2-ENGINE-045 |
| Check 23 | scalar 0=index | GetVar(index) into map result register | R2-ENGINE-045 |
| Check 24 | scalar 0=n | GetVar(752+n) into map result register | R2-ENGINE-045 |
| Instant 36 | scalar 0=n, scalar 1=mode | Mode-dependent stores to 752+n and notifications 253/254 | R2-ENGINE-046 (amended), R2-ENGINE-891 |
| Instant 37 | scalars 0..7 from X,Y,int×6 | Bad instant default, optional diagnostic packet 0x91, ordinary cleanup/return | R2-ENGINE-058 |
| Check 25 | scalars 0=X,1=Y,2=layer | Nonzero terrain cell effect-layer pointer | R2-ENGINE-053 |
| Check 26 | resolved unit, scalar 0=effectID | Nonzero active effect bit 1<<(effectID&31) | R2-ENGINE-054 |
| Check 27 | resolved unit, reads scalars 1/2 | Tile X/Y equality plus fractional bytes 128/128 | R2-ENGINE-055 |
| Instant 38 | item u16+0x6c | One extraction call per iterated unit, returned-item destruction and unit refresh | R2-ENGINE-056 |
| Instant 39 | resolved group+0x60 | Clear forced-activity DWORD in the group's record | R2-ENGINE-057 |

Instant 36 mode 1 stores 1 only from old 0. Mode 2 leaves old 4 unchanged;
otherwise it emits 253 only if old bit 2 is clear, then stores 3. Mode 4 always
emits 254 and stores 5. Other modes store their value directly. Objective rows
read752+row, omit zero values, and test bits 2/4/1 in that priority for
mark values 11/12/10, which draw subobj.256 frames 1/2/0 (R2-ENGINE-889; the
former frame clause of R2-ENGINE-046 is partially retracted). Generic35/36 and
EnterLocation reset are known writers;
an acknowledgement transition to 4 and semantic names for states4/5 remain
Unknown. — R2-ENGINE-046

Check 25 packs low16(X+256Y). The terrain helper requires cell flag0x20 and
an existing cell record, then returns its layer-pointer array. Scalar 2 selects
a DWORD pointer without an explicit arm bound. RU effect removal counts six
layer pointers and maps effect IDs 3/6/17/15/14 to layers 0/2/3/4/5. Layer 1,
full producer domains and malformed operands remain Unknown.
— R2-ENGINE-053

Check 26's unit DWORD+0x144 is linked to timed-effect registration and expiry,
which set/clear the bit selected by the effect's u16 ID. Full duration,
duplicate-effect and removal behavior remains Unknown. — R2-ENGINE-054

Check 27's catalog declares Unit,X,Y, but the literal compiler packs X/Y into
scalars 0/1 while the executor reads scalars 1/2. No check 27 occurs in either
46-map campaign corpus. The intended authored coordinate convention is
Unknown; using intuitive indices0/1 would change the observed native predicate.
— R2-ENGINE-055

Instant 38 iterates the terrain-provided unit list and calls extraction once
per unit. Quantity<=1 unlinks the matching item; larger quantities invoke a
virtual extraction callback. The complete list membership and callback's
returned quantity remain Unknown. No drain-all-copies behavior follows from
this loop. — R2-ENGINE-056

Instant 39 clears group.record+0x48; the constructor's initial value is 0 and
compiler references can set it to 1. Recount marks active groups from a coarse
actor-proximity bitmap or this override. Reset clears group active bytes only
in rosters whose owner DWORD+0x2c is nonzero. Under Session gate 0, an inactive group
can skip group AI; an inactive unit without an explicit command enters action
kind 27, while a nonzero command continues processing. Groups in active cells
remain active after the clear. — R2-ENGINE-057

On the observed ordinary mission tick route, the activity object is embedded
at terrain+0x92ecc, where terrain is Session.DWORD(+0xa50). The bitmap is
activity-object+0x400, hence terrain+0x932cc. The recount counters are
activity-object+0x1614/+0x1624, hence terrain+0x944e0/+0x944f0.
These offsets use the embedded object as their local base. — R2-ENGINE-057

The bitmap producer requires actor.owner.DWORD(+0x2c)==0 and ORs
owner.u16(+0x32) over a 5x5 neighborhood centered on tile(X>>3,Y>>3).
Its other owner branch can mark a group active under the literal actor
i16(+0x94)<i16(+0x96) predicate. The owner fields' classification, other
producer exceptions and equality with GUI fog are Unknown. Clearing the
override is not an unconditional immediate suspension. — R2-ENGINE-057

The compared actor words are signed current and maximum refill values.
Strict current<maximum sets that actor's group active on the nonzero-owner
branch; the health label retains a Medium interpretation boundary.
— R2-ENGINE-086

Recount rebuilds coverage from its actor collection and from each zero-class
owner's nonnull +0x38 actor, then marks collection actors. The additional
owner producer runs before marking. Later loops update other receivers'
coverage masks. Collection membership, owner+0x38 producers and every
command-byte producer remain Unknown. A nonzero command byte bypasses
the selected activity gate; later execution checks remain separate.
— R2-ENGINE-087

No instant 37 occurs in either 46-map corpus. The verified default optionally
delivers its report through the ordinary packet/message route and returns;
other music paths and arbitrary authored runtime behavior remain Unknown.
— R2-ENGINE-058

## Ordinary event and result dependencies

Instant 2 emits packet 0xb6 with its event ID. The actual receiver tables in both
clients forward it to UI 433. Values250/253/254/255 take separate branches;
ordinary values format event%d when UI mode bit 8 is clear or enter a deferred
queue otherwise. This establishes the event interface. Text-page parsing and
speech are separate consumer surfaces. — R2-ENGINE-047

The ordinary trigger runner compares two map result registers with signed
codes0..5: ==,!=,>,<,>=,<=. Both stored operands are check IDs; the right operand
is not an immediate value. Conditions short-circuit their conjunction. A nonzero
trigger flag skips an already-set byte latch. Otherwise the runner clears that
byte before evaluation and sets it on a pass before executing actions in stored
order. All 16 mission10 stored DWORD flags at+0xb4 are 1. — R2-ENGINE-059

Mission10 trigger7 tests result(check 12)<=result(check 4), then executes
action IDs 9,23,12,28 in that order. Both operands are original check IDs.
— R2-ENGINE-059

Instant 4 increments the session win counter+0xbe0c; instant 5 stores a failure
reason+0xbe14; instant 8 increments a selected map result register. The verified
RU result tick prioritizes a positive failure reason and emits 0xb4, otherwise
a nonzero win counter emits 0xb5. The receiver forwards these to UI 433 selector255
and UI 430 respectively. Full acknowledgement scheduling is Unknown. Its former
mission-10 reason 19 is retracted: mission 10 stores 0 and 5. — R2-ENGINE-060

DLL EnterLocation stores the location-record pointer and clears slots 752..767.
GetCurrentLocation returns a pointer; the client wrapper reads its numeric ID.
The ordinary LeaveLocation mission10 arm marks completed-bank slot 896+10,
adds location20 and returns output 1; that selected case has no 753/754 guard.
Other continuation arms, persistence and complete native campaign playthrough
remain Unknown. — R2-ENGINE-048

## Failure reasons

The failure register Session +0xbe14 has eleven raw-displacement users in
each root: the Session init stores 0, instant 5 stores scalar 0, check 18
and a detector that has no EN reference (R2-ENGINE-628) add 1, a
serializer passes the value and its address to two unread helpers, and the
result tick reads it. Other
encodings of the access are outside the population. — R2-ENGINE-623

The result tick, equal in EN and RU after masking, works on the owner while
its byte +0x41 is set and +0x2c is zero, and, while the tick object's +0x74
is zero, stops sending once owner +0x40 is 2 or more. It sends reason 2 when owner +0x38 is null or that actor's
byte +0x13c is set (and reason 0 when owner byte +0x42 is also set); else a
positive +0xbe14 is sent and its low byte stored at +0x40; else a nonzero
win counter sends 0xb5 and stores 1. Byte +0x13c and +0x42 are Unknown in
meaning. — R2-ENGINE-624

Outside the script, the death of a class 0x21..0x3f actor, with
`[L2.00805]` +0x74 zero and owner +0x2c zero, stores owner +0x40 = identity
+0x14c − 19 and sends it: identity 21 gives reason 2, identities 22..30
give 3..11. Hire classes, at Medium, lie outside the range. — R2-ENGINE-625

Check 18 adds 1 to +0xbe14 on each check-loop pass while the unit at IR
+0x5c has state +0x54 = 0x10, and writes no check result. Each corpus holds
one check 18: map 10 ID 1 on unit reference 32, named by no trigger;
whether it runs is Unknown. — R2-ENGINE-627

Each corpus holds 81 instant-5 nodes, equal in both roots, parameter 0 of
type 1 and no other parameter. They store 0 or 5..9; map 96's triggered 5
and map 40's untriggered 7 exceed their mission's failure cache, and no
node stores 19. — R2-ENGINE-630

Which events fail a mission and the text each brings are composed in
R2-ENGINE-631.

## Activity owner field producers

Selected EN/RU owner construction sets DWORD+0x2c=1 and u16+0x30=0.
The record builder sets +0x2c to zero for ordinal 1 when Session+0x74==0.
Under +0x2c==0, registration computes +0x30 from signed runtime owner key+4
and copies it to +0x32; nonzero registration clears +0x30 without a +0x32
assignment. Archive loading copies loaded +0x30 to +0x32 and assigns actor
owner pointers at +0x14. — R2-ENGINE-071

The mask stores the low word of a DWORD one shifted by signed key%16,
with the x86 shift count masked to five bits. Negative-key reachability,
other writers, human/AI classification, authored control identity and GUI
fog remain Unknown. — R2-ENGINE-071

## Selected construction and literal-list dependencies

The selected owner actor collection's native append chain links each new
node after the old tail. Its iterator starts at the head and follows next,
returning the pointer at node+8. The literal Unit helper compares identities
in that nonnull payload prefix and returns its first encountered match.
Null payloads terminate the primary loop; other collection mutators and
actual initial members remain Unknown. — R2-ENGINE-123

The character producer's new-object branch requests 0x254 (596) bytes and
calls a constructor whose parameter-record loader runs before registration.
The producer later writes owner+0x38 and, under Session+0x74==0, explicitly
stores actor WORD+0x14c=21. The existing-object branch has a separate reuse
route. Full native class, base initialization, later callback effects and
ordinary-start reachability remain Unknown. — R2-ENGINE-124

Selected initialization sets Session+0x74 from the signed predicate mode<2.
The accepted map-object branch overwrites it from signed DWORD(+0xd8)>1;
that map field's authored origin and meaning remain Unknown. The selected
owner reuse comparison to 1 reads the global owner collection's DWORD+0x0c,
through a complete getter, rather than owner+0x24. It supplies no actor roster
size or complete initial member population. — R2-ENGINE-125

The measured construction/list dependencies leave initial native classes,
complete selected-owner membership and successful literal Unit resolution
Unknown. Indirect delivery, unexpanded virtual targets, map metadata loaders
and other collection producers remain open. These conditional contracts do
not establish a complete ordinary single-player start or mission transition.
— R2-ENGINE-127

A second static pass decodes 24 further EN/RU routine pairs (48 complete
bodies) reached by direct edges from those dependencies. 15 pairs match in
normalized shape and 9 differ. — R2-ENGINE-135

Class records name the owner object class Player (size 0xa80, base CObject)
and an actor chain Human > Humanoid > Unit > Token > CObject (size 0x254).
The Player create slot calls the owner constructor. That the character
producer's 0x254-byte object is class Human is a Medium inference: the Human
create slot calls a different constructor. — R2-ENGINE-136

Heap Player construction is reached through the dispatcher and one caller
body. A second body builds a stack-frame Player and allocates no heap object.
The Player factory has no direct caller. Which route builds the ordinary-start
owner is Unknown. — R2-ENGINE-137

In the measured map builder the owner-build, actor-build and literal-reference
caller calls occur in that order with no exit between the first and last. The
caller has two direct callers. Jump tables and the map builder's callers keep
consumer reachability Medium. — R2-ENGINE-138

Ten of 13 actor-list append sites in the new population load their receiver
from a reg+0x24 field. The base object at those sites is not identified, so no
site is shown to append to the owner's collection. — R2-ENGINE-139

The new population holds 109 EN / 113 RU indirect calls and 14 decoded jump
tables per locale. Outside .text, raw pointers to the selected entries exist only in a
.data class-record slot and five .rdata cells per locale. All indirect targets
are Unknown. — R2-ENGINE-140

Complete initial membership and order of the owner's actor collection, the
startup predicate and successful literal resolution remain Unknown after this
pass. — R2-ENGINE-141

## Startup caller and concrete collection receiver

The bounded replay has 21 frontier complete pairs plus one attribution miss
(22 new pairs), and 17 recaptured prior pairs. Twenty frontier pairs and the
miss match normalized shape; the remaining frontier pair uses different EN/RU
receiver field coordinates. The 14 cells include the miss's stopping sentinels. The miss's
compact selector is not enumerated. Each stopping DWORD overlaps its first
four bytes; the decoded arms do not prove a complete caller CFG or broader
reachability. — R2-ENGINE-151

The complete map caller calls the map builder under Session+0x1ac==0;
its nonzero arm takes a separate load route. Local source/alternate/map failure
returns are 3, 4 and 5, without an established UI interpretation. Ordinary caller
values and callee return effects remain Unknown. — R2-ENGINE-152

The verified outer wrapper calls that map caller only after initialization
returns 0. Its map-call argument is prepared through an unexpanded helper.
The other measured caller requires its first argument to equal zero and a nonzero field
at receiver EN+0x5d4 / RU+0x638, then clears Session+0x1ac and calls it.
Those containing objects' equivalence and runtime predicates are Unknown.
— R2-ENGINE-153

The nonnull allocation arm of selected initialization constructs an empty
global owner collection; the null allocation arm stores zero. New
owner registration tail-appends a pointer to it. Fresh owner's actor collection
at +0x24 is separately allocated and locally empty; owner+0x38 is cleared.
Allocation, empty initialization and conditional insertion are separate facts,
with intervening effects and entering membership open. — R2-ENGINE-154

Map actor lookup compares signed row WORD+0x14 to owner DWORD+8 copied from
map owner record DWORD+4. A head match returns the same first owner selected
by the literal helper. Runtime owner WORD+4 is a different coordinate;
actual row values, loaded/reused owners and mismatch traversal remain Unknown.
— R2-ENGINE-155

After nonzero positioning return, the selected map actor code registers the
pointer globally, sets actor+0x14 to the looked-up owner and tail-appends to
that owner's collection at +0x24. Global pointer search returns do not gate that global
append. Position-world return, virtual/constructor effects and actual first
accepted row remain Unknown; this local path proves no complete uniqueness,
initial order or final numeric identity. — R2-ENGINE-156

The alternate caller iterates each owner's+0x24 and copies payloads with
`(actor BYTE+0x4c & 0x08)==0` through the global actor receiver. Source iteration
and destination append use different storage coordinates. Runtime pointer
aliasing, archive/virtual construction and initial source members remain Unknown. — R2-ENGINE-157

Positive ordinary-start members, the first accepted row, total order,
uniqueness, final actor numeric identity, owner+0x38 population and literal
success remain Unknown at the measured outer-caller, lookup-next, world-result
and callback boundaries; runtime receiver aliasing is also unproved.
— R2-ENGINE-158

## Selected initialization field boundary

The selected initialization's explicit stores through its saved entry receiver
are one DWORD at +0xf8 and indexed DWORDs at +0xfc..+0x16c, ending at byte
+0x16f. Its complete local body has no memory operand displacement +0x174;
pointer-mediated and external mutations remain open. This does not prove that
the field is unchanged. — R2-ENGINE-167

Its first direct target receives unchanged ECX without explicitly using it.
The target passes a fixed global address, zero and 0xc00 to an unread callee,
then writes that global DWORD as one. Global/Session aliasing and the opaque
callee's ECX use remain Unknown. The queue stopped after this target: its own
opaque call and the initialization call with receiver Session+0x78 are
eligible sites left unexpanded. — R2-ENGINE-168

Local signed argument predicates store, through the symbolic saved entry
receiver, +0x74 from argument<2, +0x70 from argument>0 and +0x1a8 from the
freshly stored +0x74. Tested argument value 2 therefore gives local values 0, 1
and 0. These fields do not supply +0x174, and the saved receiver is not proved
to persist through the earlier opaque and indirect calls. — R2-SESSION-067

The numeric Session DWORD+0x174 before selected map-actor insertion remains
Unknown. The indirect call, unexpanded eligible callees, the store of the saved
receiver to a fixed global and the store through [Session+0x7c]+0x0c leave
mutations and runtime aliases open; no composed numeric value or
ordinary-start path is proved. — R2-SESSION-068

## Conditional hire admission

A reached EN mission-entry loop selects hire IDs1..20 only when both
bank[511+ID] and bank[531+ID] are nonzero. Its creator assigns the input
ID to actor WORD+0x14c and registers the actor under the selected owner.
These hire identities1..20 differ from the 21+k compared by literal
10001+k. The body has no additional numeric cap on its supplied list,
but complete mission admission and placement remain Unknown.
— R2-ENGINE-381

## Installed operation population and dispatch

Each scenario corpus has 29 ordinary action IDs and 19 ordinary check IDs,
plus the full word 65538 in both lists. Counts below are per root and agree
between EN and RU. They count stored nodes, including unreferenced nodes,
rather than executions. — R2-ASSET-247

| Action ID | Nodes | Check ID | Nodes |
|---:|---:|---:|---:|
| 2 | 448 | 1 | 58 |
| 3 | 48 | 4 | 8 |
| 4 | 46 | 5 | 62 |
| 5 | 81 | 6 | 205 |
| 6 | 102 | 7 | 62 |
| 8 | 197 | 8 | 14 |
| 10 | 60 | 9 | 5 |
| 12 | 8 | 10 | 50 |
| 13 | 21 | 14 | 1 |
| 16 | 20 | 15 | 46 |
| 17 | 2 | 16 | 5 |
| 18 | 2 | 17 | 37 |
| 19 | 10 | 18 | 1 |
| 20 | 7 | 19 | 209 |
| 21 | 5 | 21 | 24 |
| 22 | 4 | 23 | 5 |
| 23 | 10 | 24 | 3 |
| 24 | 3 | 25 | 24 |
| 25 | 9 | 26 | 3 |
| 28 | 18 | 65538 | 164 |
| 29 | 4 | — | — |
| 30 | 3 | — | — |
| 32 | 13 | — | — |
| 33 | 9 | — | — |
| 34 | 1 | — | — |
| 35 | 4 | — | — |
| 36 | 140 | — | — |
| 38 | 7 | — | — |
| 39 | 6 | — | — |
| 65538 | 46 | — | — |

The check switch subtracts one from IR operation+4, admits unsigned 0..26
and dispatches through EN L2.03276 / RU L2.03277. The action switch similarly
admits 0..38 through EN L2.00509 / RU L2.00508 after reference preparation.
Every occurring ordinary ID has a read table cell and direct handler call.
Twenty-seven action and all nineteen check bodies are read through RET in
each root. Action 33 is only a prefix; action 34's body is unread. Neither
is assigned an operand/effect contract here. — R2-ENGINE-1312

R[s] below is Session DWORD+0xc784+4*s. pN is compiled DWORD+8+4*N;
U/G/P/V are resolved pointers at IR+0x5c/+0x60/+0x64/+0x68, and I is
u16+0x6c. These are compiled positions, not source parameter slots.
Read stores and call arguments are High; complete effects of unread helpers
and collection membership are Medium or Unknown. — R2-ENGINE-1313,
R2-ENGINE-1314, R2-ENGINE-1315

Actions 3 and 8 respectively store R[p0]=p1 and increment R[p0]. The other
read local effects, each High for the instructions and Medium for delegated
effects:
- 4: Session+0xbe0c += 1; the role of the field is Unknown.
- 5: Session+0xbe14 = p0; the role of the field is Unknown.
- 10: byte at Session+0xa8c4+70*p0+p1 = (old AND 0xfc)+low8(p2), then, for
  each entry of the Session+0xc77c iteration whose word +4 equals p0 or p1,
  a call to EN L2.01852 (R2-ENGINE-980 at its grade).
- 35: SetVar(p0,p1).
- 36: GetVar/SetVar on 752+p0 with the arms of R2-ENGINE-046 (partially retracted; amended
  grade).
- 39: zero stored at [group+0x3c]+0x48.

No accepted-index domain follows from these DWORD accesses.
— R2-ENGINE-1313

| Action | Additional read operand/call contract | Boundary |
|---:|---|---|
| 2 | p0 and zeros to event producer | Event interface above; delivery separate |
| 6 | p0 selects nested 18-cell switch; arms pass G, U, low8(p1/p2) | Nested table and method effects not fully read |
| 12 | I lookup; nonzero return to U+0x7c collection helper; notify U | Creation/mutation helpers unread |
| 13 | I extraction helper on U+0x7c; nonzero return gets virtual slot+4(1); notify U | Extraction quantity and callback unread |
| 16, 17 | U to EN L2.04906 or L2.01848 respectively | Final effects unread |
| 18 | U to L2.04906, then its location and constant 3 to L2.07384 on V | Callees unread |
| 19, 22 | U,P or each iterated G member,P to EN L2.07385 | Ownership effect and membership unproved |
| 20 | Calls U helper; allocates/constructs 0x24-byte object and replaces U+0x7c | Constructor and old collection effects unread |
| 21 | low8(p0..p3), p4, p5 (or 99 when p5=0) to EN L2.03575 | Creation and parameter meanings unread |
| 23 | p0,0 to EN L2.01052 on P | Final effect unread |
| 24 | low8(p0/p1), U, p2/p3 to EN L2.03574 | Final effect unread |
| 25 | Byte descriptor from p0..p7, with p4=0 repetition branch; packed low8(p2/p3) to terrain helper | Terrain meaning unread |
| 28 | Calls helper with U/V collections, replaces U+0x7c, notifies both | Complete transfer semantics unread |
| 29 | p0+256*p1 lookup; six returned pointers; matching u16+0x0c==p2 gets u16+0x50=low16(p3) | Returned pointer kinds unproved here |
| 30 | Iterates U+0x20; matching u16+0x0c==p0 gets u16+0x42=low16(p1) | Collection identity/membership unread |
| 32 | Iterates G and calls EN L2.04906 | Final effect unread |
| 33, 34 | Dispatch targets EN L2.07410/L2.07411, RU L2.07412/L2.07413 | Operands/effects Unknown |
| 38 | I extraction per iterated terrain unit, optional virtual slot+4(1), refresh call | Existing extraction boundary above |

The delegated action table is a wrapper contract in both roots, with EN
callee addresses for reference. It does not assign callee semantics from
editor names. — R2-ENGINE-1314

| Check | Read result contract | Boundary |
|---:|---|---|
| 1 | Helper result on G | Group population/count meaning Medium |
| 4 | p0=6 writes signed U.i16+0x94; other p0 writes no result | Other selectors unproved |
| 5 | U.DWORD+0x54!=16 as 0/1 | Meaning of other states unproved |
| 6 | Either unit state 16 gives 255; else low8 helper on U/V locations | Metric/getters unread |
| 7 | low8 helper on U location and low8(p0/p1) | Metric/getters unread |
| 8 | Count nested returns from P+0x28 | Membership unread |
| 9 | If [U+0x1c4].BYTE+8=5, DWORD at [[U+0x1c4]+0xc]+8; else 0 | Record-kind meaning unread |
| 10 | P/V owner-word-indexed relation byte AND 3 | Complete relation contract separate |
| 14 | Boolean terrain lookup of low8(p0/p1) | Lookup meaning unread |
| 15 | Minimum low8 location helper over P's nested collection, default 255 | Membership and metric unread |
| 16 | I lookup nonzero gives low8 location helper to low8(p0/p1); else 255 | Lookup and metric unread |
| 17 | Boolean I lookup on U+0x7c | Lookup semantics unread |
| 18 | State 16 increments Session+0xbe14 (role Unknown); no result store | Role of +0xbe14 Unknown |
| 19 | R[p0] | Source-slot limits unproved |
| 21 | Signed i16+0x42 through IR pointer+0x70 | Complete pointer/field meaning separate |
| 23, 24 | Scenario GetVar(p0) or GetVar(752+p0) | DLL contract above |
| 25, 26 | Layer pointer / active effect bit | Added-arm contracts above |

All nineteen local bodies are read independently in EN and RU. Helper
results, iterations and field meanings retain the boundaries in the table.
— R2-ENGINE-1315

## Complete stored trigger extent

Type 7 bulk-loads each trigger as 184 bytes into the compiler's record,
with no node-style field permutation. Per-byte value sets are measured
over every trigger in 142 maps: EN 715 campaign plus four loose; RU 714
campaign plus fourteen loose. Values alone do not name a field.
— R2-ASSET-248, R2-ENGINE-1316

| Offset | Stored extent | Game read |
|---|---|---|
| 0x00..0x7f | 128 bytes | Prefix copied by loader; label interpretation Medium; its separate consumers Unknown |
| 0x80/0x84, 0x88/0x8c, 0x90/0x94 | Three DWORD pairs | Left/right check IDs, not immediate right operands |
| 0x98,0x9c,0xa0,0xa4 | Four DWORDs | Action IDs in execution order; zero skipped |
| 0xa8,0xac,0xb0 | Three DWORDs | Corresponding signed-comparison code, installed values 0..5 |
| 0xb4 | DWORD | Latch skip flag (nonzero: skip a set latch); installed values 0 or 1 |

A zero first left check ID excludes the trigger from compilation (High as
code). No installed campaign trigger has one; EN 4 and RU 11 loose-map
triggers do (`triggers.tsv`). One installed campaign trigger (map 42, index
16, both roots) names left check ID 0xFFFFFFFF, which is no check ID of its
map; every other nonzero ID is the ID of a node of the same map.
The compiler skips a condition pair when either ID is zero. ID mapping and list-helper
failure/default behavior remain Medium. The stored flag is copied within
an admitted condition path; the zero-condition case's final flag is Unknown.
An action ID above 5000 causes a diagnostic and continues to mapping; this
is not a proved rejection bound. — R2-ENGINE-1316, R2-ASSET-248

## Dynamic references and variable limits

None of the 47,410 stored parameter DWORDs of these 142 maps lies in
8000..9999. This includes all ten slots and both lists, including special
nodes. No ID from that range occurs in this population, so there is no
installed-map-specific return value to enumerate. Other archives, saves,
user maps and runtime synthesis are outside the census. — R2-ASSET-249

The read action compiler nevertheless recognizes 8000<=value<10000 with
type 2/3/4. It retains the original value, records mode 1 below 9000 or
mode 2 otherwise, subtracts the corresponding base, and stores division-by-100
quotient/remainder bytes. The action prefix uses the remainder byte to load
Session+0xd724+4*index into the appropriate resolved pointer field. That
array's producers, mode/quotient meaning, check-side complete route and
failure behavior remain Medium or Unknown. — R2-ENGINE-1318

The read map-slot action 3/8 and check 19 bodies have no explicit local
index/value guard; this is a bounded observation about those bodies, not
a complete compiler/caller census. Safe map-slot capacity and index/value
limits remain Unknown. The scenario DLL's initial 1024-DWORD clear is a
different namespace and does not prove an external safe range or unbounded
admission. — R2-ENGINE-1319
