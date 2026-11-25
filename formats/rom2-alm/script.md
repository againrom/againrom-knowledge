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
reduced to ordinary opcode 2. — R2-ENGINE-042 (amended), R2-ENGINE-043 (amended)

The register is the running check index, the value is slot 0 whatever its type code, and the packed action bytes are
slots 0 and 1 whatever their type codes; see Constants and registers. — R2-ENGINE-1727

The bounded postcompile trace still leaves the packed collection's consumer
Unknown. The read caller invokes the type-9 stage and a conditional
owner/group pass; neither read establishes what consumes that collection.
This does not establish that it is unused. — R2-ENGINE-1317

The collection is a member of the compiler object, not of the map object whose own +0x28
vector the loader fills. No vector-method consumer of the compiler's member was found in
the compile function (EN L2.00201..L2.08170), the caller or the type-9 stage to EN L2.08171
(Medium). Of the 87 direct vector-call sites per root, 6 are named and 5 classified; the other
sites are unclassified, and those outside EN L2.04217..L2.04448 are taken to belong to other
classes by address only. The stage after EN L2.08171, the body of EN L2.02494 after L2.08228,
the function at EN R2.0081 and direct field reads are unread. — R2-ENGINE-1683

Ordinary check IR has result slot+0, operation+4 and scalar 0 at+8. Literal
nonreference values pack densely, including trailing scalar zeros. Literal
references resolve separately: first unit pointer+0x5c, group+0x60,
player+0x64 and second reference+0x68. Literal Target_Item becomes
low16(value+3608) at+0x6c. Failed literal reference resolution can clear the
ID mapping and omit executable IR. Dynamic references in 8000..9999 use a
different metadata path; its complete resolution/error contract is Unknown.
— R2-ENGINE-044

The check compiler marks a slot of type 2, 3 or 4 whose value lies in 8000..9999 exactly as
the action compiler does: mode 1 below 9000 and mode 2 from 9000, value minus the base,
quotient and remainder by 100 as bytes, and a slot count at IR+0x58. It leaves the reference pointer
unset and copies the metadata into the check IR (High). In the check loop (EN L2.03283, RU L2.05851),
every check copy with nonzero byte IR+0x58 is skipped before the check dispatcher runs (High); that the copy
routine (EN L2.08305, RU L2.08306) carries the compiler's count into that byte is not read (Medium), so that
every check carrying an 8000..9999 parameter is skipped is Medium. In checks 2, 3, 11, 12, 13, 20, 22 and 27
(EN and RU) no operand with displacement 0x3a or 0x44..0x58 has the IR as base and none addresses Session+0xd724
(High; the one operand in 0x3a..0x58 is [obj+0x42] on the Buildings object of check 20, EN L2.08341, RU L2.08342);
for the 19 EN handlers of R2-ENGINE-1312 and the check dispatcher the grade stays Medium (R2-ENGINE-1684), and the
19 RU handlers were read in part.
In the raw-DWORD 0xd724 scan of both roots (seven hits per root), the six indexed loads are in the
action executor and the one remaining hit (EN L2.07408, RU L2.08307) is the argument of a Session init fill routine called with
0 and 0xfa0 (Session+0xd724..0xe6c3, 0xfa0 bytes; EN L2.00341 is the memset of R2-ENGINE-996, High for EN; RU L2.00342 body unread, Medium); the EN session serializer skips +0xd724 (High; Medium
that no indirect pointer or out-of-range +0xc784 index >= 1000 writes nonzero values). —
R2-ENGINE-1684 (amended), R2-ENGINE-1318, R2-ENGINE-1451 (amended), R2-ENGINE-1697

No value in 8000..9999 under type 2, 3 or 4 is stored in the 142 installed maps, so the
installed campaign does not take the path. — R2-ASSET-249

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

The roles of +0xbe0c and +0xbe14, Unknown in R2-ENGINE-1313 (amended), are the win counter and the failure reason;
see Victory and defeat. — R2-ENGINE-1729

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
- 4: Session+0xbe0c += 1; the win counter (R2-ENGINE-1729).
- 5: Session+0xbe14 = p0; the failure reason (R2-ENGINE-1729).
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
| 6 | p0 selects nested 18-cell switch; arms pass G, U, low8(p1/p2) | Arms read: see Action 6 selectors (R2-ENGINE-1726); play meaning of the orders Medium |
| 12 | I lookup; nonzero return to U+0x7c collection helper; notify U | Final effect below (R2-ENGINE-1747) |
| 13 | I extraction helper on U+0x7c; nonzero return gets virtual slot+4(1); notify U | Final effect below (R2-ENGINE-1747) |
| 16, 17 | U to EN L2.04906 or L2.01848 respectively | Final effects below (R2-ENGINE-1743, R2-ENGINE-1744) |
| 18 | U to L2.04906, then its location and constant 3 to L2.07384 on V | Final effect below (R2-ENGINE-1744) |
| 19, 22 | U,P or each iterated G member,P to EN L2.07385 | Final effect below (R2-ENGINE-1745) |
| 20 | Calls U helper; allocates/constructs 0x24-byte object and replaces U+0x7c | Final effect below (R2-ENGINE-1748) |
| 21 | low8(p0..p3), p4, p5 (or 99 when p5=0) to EN L2.03575 | Final effect below (R2-ENGINE-1740, R2-ENGINE-1741) |
| 23 | p0,0 to EN L2.01052 on P | Final effect below (R2-ENGINE-1746) |
| 24 | low8(p0/p1), U, p2/p3 to EN L2.03574 | Final effect below (R2-ENGINE-1742) |
| 25 | Six-byte cell event (low8 p0, low8 p1, source, target) for cell (low8 p2, low8 p3), source and target (p2, p3) twice when p4=0, else (p4, p5) and (p6, p7); p8, p9 unread; written by L2.03589 on [Session+0xa50] (EN and RU read) | The writer replaces the cell's event (R2-ENGINE-1775, R2-ENGINE-1776); event 23 passes a unit at the cell centre to the cell move toward the target (R2-ENGINE-1778; Medium for units outside the three step callers; nothing observed in the original); the cast consumer uses (+0x36,+0x37) literally (R2-ENGINE-1698) |
| 28 | Calls helper with U/V collections, replaces U+0x7c, notifies both | Final effect below (R2-ENGINE-1748) |
| 29 | p0+256*p1 lookup; six returned pointers; matching u16+0x0c==p2 gets u16+0x50=low16(p3) | Final effect below (Medium; R2-ENGINE-1749) |
| 30 | Iterates U+0x20; matching u16+0x0c==p0 gets u16+0x42=low16(p1) | Final effect below (Medium; R2-ENGINE-1749) |
| 32 | Iterates G and calls EN L2.04906 | Final effect below (R2-ENGINE-1743) |
| 33 | L2.01848 on each member of G | Final effect below (R2-ENGINE-1450, R2-ENGINE-1744) |
| 34 | p0 6, 15, 16 store low16 p1 at U+0x94, +0xbe, +0xc0; other p0 store nothing; notify L2.07789 | Final effect below (Medium; R2-ENGINE-1450, R2-ENGINE-1749) |
| 38 | I extraction per iterated terrain unit, optional virtual slot+4(1), refresh call | Final effect below (R2-ENGINE-1747) |

The delegated action table is a wrapper contract in both roots, with EN
callee addresses for reference. It does not assign callee semantics from
editor names. — R2-ENGINE-1314 (amended)

### Final action effects

Ordered by scenario nodes per root (equal in EN and RU; R2-ASSET-247). Each row
states what the action does in the game; a row that rests on an unread helper carries
(Medium), and so does every row that depends on the caster lifetime or an unread getter.

| Action | Nodes | Effect | Claim |
|---:|---:|---|---|
| 13 | 21 | Takes item I out of U's bag and destroys it (take quantity Medium); sends the bag message 0x76, which stores nothing in its own body | R2-ENGINE-1747 |
| 16 | 20 | Takes U off the map (Medium, unread `L2.03601`) and out of the unit registry and sets +0x4c bit 0x08; does nothing when the bit is set | R2-ENGINE-1743 |
| 28 | 18 | Moves every item of U's bag into V's bag and gives U an empty bag; when V's count is 1024 or more (Medium, unread count helper), U's items stay in its old bag, which nothing owns (Medium) | R2-ENGINE-1748 |
| 32 | 13 | Action 16 on each member of G (Medium as action 16) | R2-ENGINE-1743 |
| 19 | 10 | Gives U to player P: new owner, owner list, a new group of its own, cleared +0x18 bits, +0x204 = 1 | R2-ENGINE-1745 |
| 23 | 10 | Adds p0 to player P's +0x3c (named gold, Medium), without a bound, and sends message 0x67 | R2-ENGINE-1746 |
| 33 | 9 | Action 17 on each member of G (Medium as action 17) | R2-ENGINE-1744 |
| 12 | 8 | Builds item I and puts it into U's bag; sends the bag message | R2-ENGINE-1747 |
| 20 | 7 | Drops U's bag as a sack at U's cell with gold 0 (Medium, EN sack drop unread) and gives U an empty bag | R2-ENGINE-1748 |
| 38 | 7 | Action 13 for every unit of one Session list (that it is every unit is Medium) | R2-ENGINE-1747 |
| 21 | 5 | Casts spell low8(p4) once from cell (p0, p1) at cell (p2, p3) with power word p5, 99 when p5 = 0; the caster is then deleted (Medium) | R2-ENGINE-1740, R2-ENGINE-1741 |
| 22 | 4 | Action 19 on each member of G | R2-ENGINE-1745 |
| 29 | 4 | Sets remaining time p3 on the area effects of spell p2 at cell (p0, p1) (Medium) | R2-ENGINE-1749 |
| 24 | 3 | Casts spell low8(p2) once from cell (p0, p1) at unit U with power word p3, no default; the caster is then deleted (Medium) | R2-ENGINE-1740, R2-ENGINE-1742 |
| 30 | 3 | Sets duration p1 on U's effects of spell p0 (Medium) | R2-ENGINE-1749 |
| 17 | 2 | Puts U back on the map at the cell its stored position gives (Medium, unread getters), else within 3 cells (Medium, unread free-cell test); re-adds it to the registry and clears bit 0x08 | R2-ENGINE-1744 |
| 18 | 2 | Takes U off the map (Medium) and puts V within 3 cells of U's cell (Medium, unread free-cell test) | R2-ENGINE-1744 |
| 34 | 1 | Stores p1 as U's health, defence or absorption for p0 6, 15 or 16 (Medium) | R2-ENGINE-1749 |

With Mind 30 the spell power of actions 21 and 24 equals the power word for 1
to 255, is clamped to 255 above, and is 0 for a word of 32768 or more (Medium).
— R2-ENGINE-1741, R2-ENGINE-1742

### Shipped caster actions 21 and 24

Both roots hold the same scenario nodes; RU also holds one loose map node. Spell
names follow R2-ENGINE-1536 and R2-ENGINE-025 (Medium; that p4 is a Spells row number is not read).

| Map | Node | Action | Source | Target | Spell | Power | Trigger |
|---|---:|---:|---|---|---|---:|---|
| 53 | 14 | 21 | (15, 121) | cell (13, 22) | Blizzard | 100 | none |
| 60 | 6 | 21 | (23, 66) | cell (23, 66) | Poison Cloud | 100 | 3, four times |
| 71 | 5 | 21 | (66, 47) | cell (66, 50) | Wall of Earth | 99 | 28 |
| 95 | 26 | 21 | (66, 25) | cell (66, 25) | Poison Cloud | 100 | 0 |
| 95 | 27 | 21 | (69, 27) | cell (69, 27) | Poison Cloud | 100 | 0 |
| 96 | 15 | 24 | (26, 69) | unit 20 | Stone Curse | 100 | 0 |
| 96 | 16 | 24 | (26, 69) | unit 21 | Stone Curse | 32768 | 0 |
| 96 | 17 | 24 | (26, 69) | unit 22 | Stone Curse | 32768 | 0 |
| Valley2 (RU loose) | 2 | 24 | (105, 27) | unit 556 | Stone Curse | 255 | 0 |

Map 60 trigger 3 fires when check 15 for player 1 and cell (23, 66) gives at
most 2; triggers 71/28, 95/0 and 96/0 compare constant 0 with constant 0. These
four have latch flag 1 and fire once. RU Valley2 trigger 0 needs check 21 on
object 551 and R[74] both 0 and has latch flag 0. No shipped action 21 node has
p5 = 0. — R2-ASSET-379

### Check result contracts

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

## Remaining action and check bodies, tick order and unit counts

The executable parses action numbers 1..39 and check numbers 1..27. Every check cell and 38
action cells reach a wrapper call (R2-ENGINE-1312). The cell of action 37 (EN L2.00511,
RU L2.00510) leads to a message-formatting arm that calls L2.04048 (EN) and posts through
L2.00513 when the diagnostic object is set; its meaning and whether the out-of-range default
is the same arm are Unknown. Bodies of the numbers that occur in no installed map are read in
both roots except action 37, equal after masking; delegated helper effects stay Medium, and
the x/y reads of checks 2, 3, 22 and 27 are Medium.
Pointer names U/V/P/G are as in R2-ENGINE-1313. — R2-ENGINE-1450

| Action | Body | Check | Body |
|---:|---|---:|---|
| 1 | Session+0xc710 and register 902 each add 1 | 2 | R=1 when U's low8 x is in p0..p2 and y in p1..p3 |
| 7 | helper L2.07788 on P with low8 p0 | 3 | R=1 when terrain metric (U, p0, p1) is at most p2 |
| 9 | returns at once | 11, 13 | return at once, no result |
| 11 | removes item from U+0x7c (unlinks qty <= 1; splits qty > 1 through slot +0x40, read for Weapon, other item classes Medium), subtracts weight, notifies U, inserts/merges into V+0x7c, notifies V (R2-ENGINE-1699) | 12 | R=1 when the item is in U's collection |
| 14, 15 | same call: global L2.00994 with U and 0xffb | 20 | count of P's Buildings in [[[L2.00805]+0x7c]] with i16+0x42 above 0; does not count units (R2-ENGINE-1696) |
| 26 | word +0x42 of the IR+0x70 object = low16 p0 | 22 | R=1 when terrain cell (p0, p1) exists; helper called on it |
| 27 | cell move EN L2.03593 / RU L2.00093 on U to (p0, p1) (cell test L2.03600, then detach, position store, placement, command reset and broadcast calls; helper bodies unread), then Guard setter EN L2.00441 / RU L2.02815 when owner+0x2c != 0, whether or not the move was accepted (R2-ENGINE-1700) | 27 | R=1 when low8 x = p1, y = p2 and helper true |
| 31 | filters Buildings by vtable +0x38 (RU: 0 for Pointer, Shop, Inn and, by R2-ENGINE-517, Building, 1 for Outpost; EN: 0 for Building and Pointer, 1 for Outpost, R2-ENGINE-1820), writes p1 into +0xb4, the remaining respawn countdown read by Outpost::slot5, of every Outpost whose group number equals p0, for every owner and with no loop exit (the filter passes only the Outpost in the classes with a record; EN Shop and Inn slot 14 unread, Medium; R2-ENGINE-1701); p1 counts container passes without S and the check fires on the (p1 + 1)-th pass (Medium, R2-ENGINE-1821) | | |
| 33 | action 17's helper on each member of G | | |
| 34 | p0 6, 15, 16 stores p1 at U+0x94, +0xbe, +0xc0 | | |

Rows for the other numbers are in the tables above. — R2-ENGINE-1450 (amended), R2-ENGINE-1696, R2-ENGINE-1699, R2-ENGINE-1700, R2-ENGINE-1701 (amended), R2-ENGINE-1820, R2-ENGINE-1821

One script tick adds 1 to register 999, runs the check loop, runs the trigger runner and
empties the Session+0xbbec list, and the turn routine runs the tick before the loop over the
group list (R2-ENGINE-936), also while Session+0xbbe8 is nonzero (the field's meaning is not
read). The check loop skips any check whose copied IR+0x58 byte is nonzero (High); the check compiler
counts 8000..9999 parameters under types 2, 3 and 4 into that byte, and that the check copy keeps the count is Medium (R2-ENGINE-1697). — R2-ENGINE-1451 (amended)

Group list membership by unit state and script checks 1, 5, 8 and 20 (R2-ENGINE-1452 amended,
R2-ENGINE-1696): an unconscious unit (`+0x54 = 0` after the actor step: health above -10, or health -10 or below while the `+0x6c` countdown has not run out; `+0x13c = 1`) stays in its
group list, while any dead unit (`+0x54 = 0x10`, whether leaving a body `+0x13c` in `2..4` or no
body `+0x13c = 5`, `+0x94 = -10001`) is detached from its group list by the kill walker
(EN `L2.04418`, RU `L2.04426`) via `owner->L2.01670` -> `group->L2.04986`, only when `[unit+0x14] != 0` (units with a zero owner pointer are not read). Living and unconscious units are counted by checks 1, 5 and 8; dead units are not. Check 20 walks the
Buildings collection `[[[L2.00805]+0x7c]]` and tests Building word `+0x42 > 0`; it does not count
units in any state.

| Unit state | Stays in group list | Check 1 (`G+0xc`) | Check 5 (`+0x54 != 0x10`) | Check 8 (one per member node of the groups of `P`) | Check 20 (Buildings) | Claim |
|---|---|---|---|---|---|---|
| Unconscious (`+0x54 = 0` after the actor step) | Yes (High) | Counts (Medium) | Counts / 1 (High) | Counts (Medium) | No (High) | R2-ENGINE-1696 |
| Dead leaving body (`+0x54 = 0x10`, `+0x13c` in `2..4`) | No (High) | No (Medium) | No / 0 (High) | No (Medium) | No (High) | R2-ENGINE-1696 |
| Dead leaving no body (`+0x54 = 0x10`, `+0x13c = 5`) | No (High) | No (Medium) | No / 0 (High) | No (Medium) | No (High) | R2-ENGINE-1696 |

The Medium cells for checks 1 and 8 name the RU group-count removal helpers (the unlink twins of `L2.04533` and `L2.08333`), which were not read; in EN the chain from the detach to the unlink of the member node and the decrement of `G+0xc` was read. Check 8 counts member nodes and does not read `G+0xc`.

The check-ID map is hash find-or-insert: a check ID absent from the map inserts a node.
— R2-ENGINE-1453 (amended)

The new node's value is zeroed, so that ID names register 0, the first check's register (Medium). — R2-ENGINE-1728

## Names, text fields and the shared text buffer

Every action and check node carries a 64-byte name at file +0x00 and ten 64-byte text fields at
+0x9c..+0x31b; every trigger carries a 128-byte prefix. In the 46 campaign maps of both roots all
names, all text fields and all prefixes are nonempty (longest 45, 23 and 56 bytes). — R2-ASSET-364

The loader puts the name at object+0 and the text fields at object+0x98 (High). The readers
found are two compile-failure messages of the script compiler, a check and an action, each
passing the node ID and the node pointer to the formatting call; the format string was read in
an uncommitted pass and the message sink is unread (Medium). No reader of the ten text fields
or of the trigger prefix was found; this is a census of direct accessor calls (13 sites per
root) and a displacement scan for +0x98 over `.text` whose 174 sites outside the loader region
were not classified, so it is Medium. — R2-ENGINE-1681

Mission entry loads `mission<N>.txt` into the one global string that the dialogue lookup of the
TALK action also searches (High for the loader and the TALK body; Medium for the lookup's use of
that string), and each text load replaces that string. Which file the string holds at a given
TALK is not settled (Unknown). The 46 inn keys are in `town.txt`; the mission files carry the
section tags of R2-ENGINE-037; the key of a TALK from any other source is not enumerated.
— R2-ENGINE-1682, R2-ASSET-115

In the 46 campaign maps no script operand is a zone record by the bodies read (Medium). The
operand kinds are group, unit, player, item and one type-9 object; trigger-list accessor calls
are in the loader and the compiler only. — R2-ENGINE-1685 (amended)

The type-9 object is the structure whose type-4 id equals the value: the compiler keys each
object of its +0x24 list by DWORD+8, the type-4 id (Medium that the list is the Session
Buildings list). — R2-ENGINE-1725

## Source slots and compiled positions

The action compiler loops the ten source slots in order and sends each by its type code (node +0x74+4*slot).
Types 2, 3 and 4 store a resolved group at IR+0x60, player at +0x64 and unit at +0x5c; a second reference of the same
type goes to +0x68. Type 8 stores u16 +0x6c = low16(value + 3608) and type 9 a pointer at +0x70. Every other type code
(0, 1, 5, 6, 7) appends the value to the next scalar position pN at IR+8+4*N, N counting only those slots. Values
8000..9999 under types 2, 3 and 4 skip the arms (R2-ENGINE-1318). This is the rule R2-ENGINE-044 states for checks.
The 46 campaign maps give each ordinary ID the same type vector in every node, except action 6 (six vectors, told apart
in the installed maps by the value of slot 0: 3; 1; 15; 10; 11; 4, 5 or 14; the compiler reads type codes only) and check
word 65538 (two). English and Russian campaign maps agree. In a check, type 9 (check 21, s0=T) takes the lookup EN L2.08172
and stores the pointer at IR+0x70; the lookup is read in R2-ENGINE-1725. — R2-ENGINE-1680, R2-ASSET-363

| Kind | ID | Type codes (s0..s9) | Reference slots | Scalar slots (p0, p1, ... in slot order) | Nodes, EN campaign | Claim |
|---|---:|---|---|---|---:|---|
| action | 2 | 1,0,0,0,0,0,0,0,0,0 | none | s0-9 | 448 | R2-ASSET-363 |
| action | 3 | 1,1,0,0,0,0,0,0,0,0 | none | s0-9 | 48 | R2-ASSET-363 |
| action | 4 | 0,0,0,0,0,0,0,0,0,0 | none | s0-9 | 46 | R2-ASSET-363 |
| action | 5 | 1,0,0,0,0,0,0,0,0,0 | none | s0-9 | 81 | R2-ASSET-363 |
| action | 6 | 7,0,0,0,0,0,0,0,0,2 | s9=G | s0-8 | 8 | R2-ASSET-363 |
| action | 6 | 7,1,0,0,0,0,0,0,0,2 | s9=G | s0-8 | 1 | R2-ASSET-363 |
| action | 6 | 7,2,4,1,0,0,0,0,0,0 | s1=G s2=U | s0,3-9 | 14 | R2-ASSET-363 |
| action | 6 | 7,4,0,0,0,0,0,0,0,2 | s1=U s9=G | s0,2-8 | 1 | R2-ASSET-363 |
| action | 6 | 7,4,1,0,0,0,0,0,0,2 | s1=U s9=G | s0,2-8 | 7 | R2-ASSET-363 |
| action | 6 | 7,5,6,0,0,0,0,0,0,2 | s9=G | s0-8 | 71 | R2-ASSET-363 |
| action | 8 | 1,0,0,0,0,0,0,0,0,0 | none | s0-9 | 197 | R2-ASSET-363 |
| action | 10 | 1,1,1,0,0,0,0,0,0,0 | none | s0-9 | 60 | R2-ASSET-363 |
| action | 12 | 4,8,0,0,0,0,0,0,0,0 | s0=U s1=I | s2-9 | 8 | R2-ASSET-363 |
| action | 13 | 4,8,0,0,0,0,0,0,0,0 | s0=U s1=I | s2-9 | 21 | R2-ASSET-363 |
| action | 16 | 4,0,0,0,0,0,0,0,0,0 | s0=U | s1-9 | 20 | R2-ASSET-363 |
| action | 17 | 4,0,0,0,0,0,0,0,0,0 | s0=U | s1-9 | 2 | R2-ASSET-363 |
| action | 18 | 4,4,0,0,0,0,0,0,0,0 | s0=U s1=V | s2-9 | 2 | R2-ASSET-363 |
| action | 19 | 4,3,0,0,0,0,0,0,0,0 | s0=U s1=P | s2-9 | 10 | R2-ASSET-363 |
| action | 20 | 4,0,0,0,0,0,0,0,0,0 | s0=U | s1-9 | 7 | R2-ASSET-363 |
| action | 21 | 5,6,5,6,1,1,0,0,0,0 | none | s0-9 | 5 | R2-ASSET-363 |
| action | 22 | 2,3,0,0,0,0,0,0,0,0 | s0=G s1=P | s2-9 | 4 | R2-ASSET-363 |
| action | 23 | 3,1,0,0,0,0,0,0,0,0 | s0=P | s1-9 | 10 | R2-ASSET-363 |
| action | 24 | 5,6,4,1,1,0,0,0,0,0 | s2=U | s0-1,3-9 | 3 | R2-ASSET-363 |
| action | 25 | 1,1,5,6,5,6,5,6,0,0 | none | s0-9 | 9 | R2-ASSET-363 |
| action | 28 | 4,4,0,0,0,0,0,0,0,0 | s0=U s1=V | s2-9 | 18 | R2-ASSET-363 |
| action | 29 | 5,6,1,1,0,0,0,0,0,0 | none | s0-9 | 4 | R2-ASSET-363 |
| action | 30 | 4,1,1,0,0,0,0,0,0,0 | s0=U | s1-9 | 3 | R2-ASSET-363 |
| action | 32 | 2,0,0,0,0,0,0,0,0,0 | s0=G | s1-9 | 13 | R2-ASSET-363 |
| action | 33 | 2,0,0,0,0,0,0,0,0,0 | s0=G | s1-9 | 9 | R2-ASSET-363 |
| action | 34 | 4,1,1,0,0,0,0,0,0,0 | s0=U | s1-9 | 1 | R2-ASSET-363 |
| action | 35 | 1,1,0,0,0,0,0,0,0,0 | none | s0-9 | 4 | R2-ASSET-363 |
| action | 36 | 1,1,0,0,0,0,0,0,0,0 | none | s0-9 | 140 | R2-ASSET-363 |
| action | 38 | 8,0,0,0,0,0,0,0,0,0 | s0=I | s1-9 | 7 | R2-ASSET-363 |
| action | 39 | 2,0,0,0,0,0,0,0,0,0 | s0=G | s1-9 | 6 | R2-ASSET-363 |
| action | 65538 | 5,6,0,0,0,0,0,0,0,0 | none | s0-9 | 46 | R2-ASSET-363 |
| check | 1 | 2,0,0,0,0,0,0,0,0,0 | s0=G | s1-9 | 58 | R2-ASSET-363 |
| check | 4 | 4,1,0,0,0,0,0,0,0,0 | s0=U | s1-9 | 8 | R2-ASSET-363 |
| check | 5 | 4,0,0,0,0,0,0,0,0,0 | s0=U | s1-9 | 62 | R2-ASSET-363 |
| check | 6 | 4,4,0,0,0,0,0,0,0,0 | s0=U s1=V | s2-9 | 205 | R2-ASSET-363 |
| check | 7 | 4,5,6,0,0,0,0,0,0,0 | s0=U | s1-9 | 62 | R2-ASSET-363 |
| check | 8 | 3,0,0,0,0,0,0,0,0,0 | s0=P | s1-9 | 14 | R2-ASSET-363 |
| check | 9 | 4,0,0,0,0,0,0,0,0,0 | s0=U | s1-9 | 5 | R2-ASSET-363 |
| check | 10 | 3,3,0,0,0,0,0,0,0,0 | s0=P s1=V | s2-9 | 50 | R2-ASSET-363 |
| check | 14 | 5,6,0,0,0,0,0,0,0,0 | none | s0-9 | 1 | R2-ASSET-363 |
| check | 15 | 3,5,6,0,0,0,0,0,0,0 | s0=P | s1-9 | 46 | R2-ASSET-363 |
| check | 16 | 4,8,5,6,0,0,0,0,0,0 | s0=U s1=I | s2-9 | 5 | R2-ASSET-363 |
| check | 17 | 4,8,0,0,0,0,0,0,0,0 | s0=U s1=I | s2-9 | 37 | R2-ASSET-363 |
| check | 18 | 4,0,0,0,0,0,0,0,0,0 | s0=U | s1-9 | 1 | R2-ASSET-363 |
| check | 19 | 1,0,0,0,0,0,0,0,0,0 | none | s0-9 | 209 | R2-ASSET-363 |
| check | 21 | 9,0,0,0,0,0,0,0,0,0 | s0=T | s1-9 | 24 | R2-ASSET-363 |
| check | 23 | 1,0,0,0,0,0,0,0,0,0 | none | s0-9 | 5 | R2-ASSET-363 |
| check | 24 | 1,0,0,0,0,0,0,0,0,0 | none | s0-9 | 3 | R2-ASSET-363 |
| check | 25 | 5,6,1,0,0,0,0,0,0,0 | none | s0-9 | 24 | R2-ASSET-363 |
| check | 26 | 4,1,0,0,0,0,0,0,0,0 | s0=U | s1-9 | 3 | R2-ASSET-363 |
| check | 65538 | 1,0,0,0,0,0,0,0,0,0 | none | s0-9 | 160 | R2-ASSET-363 |
| check | 65538 | 7,0,0,0,0,0,0,0,0,0 | none | s0-9 | 4 | R2-ASSET-363 |

## Authoring a map script

The rules below are what a map editor needs to write nodes and triggers that compile as the installed maps do. Counts
are per root over the 46 campaign maps unless a loose-map count is given. — R2-ASSET-375

### Type-9 operands

Check 21 is the one installed user of type 9. Its slot value names a structure by its type-4 id (file +0x12). The
compiler keys every object of its +0x24 list by DWORD+8, which is that id, looks the value up and stores the object
at IR+0x70; check 21 then sets its register to the object's sign-extended word +0x42 (High for the bodies; Medium that
the list is the Session Buildings list). A value that names no structure of the map fails the node's compile (Medium:
one constructor after the zero fill is unread). A zero type-4 file word +0x0c clears Building +0x42 (R2-ENGINE-1288);
what +0x42 holds is Unknown. — R2-ENGINE-1725

| Measure | Count | Claim |
|---|---|---|
| Check 21 values that are type-4 ids of the same map (kinds 0x1c / 0x1d) | 24 (12 / 12); RU loose 1 (0x1d) | R2-ASSET-375 |
| Action nodes with a type-9 slot | 0 | R2-ASSET-375 |

### Action 6 selectors

Action 6 takes the selector in p0 and the group G; U, p1 and p2 follow in slot order. Selectors 6..9, 12, 13 and 16,
0 and values above 18 do nothing. Order codes are those of the player-order setters; what each order and group
programme value does in play is Medium. — R2-ENGINE-1726

| p0 | Operands | Effect | Campaign nodes | Claim |
|---:|---|---|---:|---|
| 1 | G, p1 | each member guards (order 0xb); programme 1; record +0x2d and +0x38 raised to p1, or to Session byte +0xa8b8 when p1 is 0 | 1 | R2-ENGINE-1726, R2-ASSET-375 |
| 2 | G, p1, p2 | members reset; programme 2; record cell = (p1, p2) | 0 (loose EN 2, RU 3) | R2-ENGINE-1726, R2-ASSET-375 |
| 3 | G | each member gets order 0xc; programme 3 | 8 | R2-ENGINE-1726, R2-ASSET-375 |
| 4 | G, p1, p2 | each member moves (order 1) to (p1, p2), keeping its offset from the group centre when [[G+0x44]+0x34]+0x1f is 2; programme 4; record cell = (p1, p2) | 20 | R2-ENGINE-1726, R2-ASSET-375 |
| 5 | G, p1, p2 | as 4, programme 5 | 32 | R2-ENGINE-1726, R2-ASSET-375 |
| 10 | G, U | each member attacks U (order 3) when its score for U is valid, else stops | 1 | R2-ENGINE-1726, R2-ASSET-375 |
| 11 | G, U, p1 | each member gets order 8 on U with byte p1 (3 when 0); a member equal to U stops | 7 | R2-ENGINE-1726, R2-ASSET-375 |
| 14 | G, p1, p2 | each member gets order 0xa with a path to (p1, p2) | 19 | R2-ENGINE-1726, R2-ASSET-375 |
| 15 | G, U, p1 | each member gets order 0x11 on U with byte p1 (3 when 0); a member equal to U stops | 14 | R2-ENGINE-1726, R2-ASSET-375 |
| 17 | G | programme 0x11; record cell = the first member's cell | 0 (RU loose 1) | R2-ENGINE-1726, R2-ASSET-375 |
| 18 | G, p1 | as 1, programme 0x12, default range 20 | 0 | R2-ENGINE-1726, R2-ASSET-375 |

A move target is clamped to four bytes of the map object, read as its movable bounds (Medium). — R2-ENGINE-1726

### Constants and registers

The compiler gives check IDs registers by a running index k over the check list in file order: a check of word 65538
and a check that compiles take register k and advance it; a check of a word above 65538 compiles nothing, advances
nothing and so shares the next check's register. Word 65538 stores its slot-0 value in its register at compile time
and reads no type code. An action of word 65538 adds value slots 0 and 1, as bytes, to a compiler collection whose
consumer is Unknown (R2-ENGINE-1683). — R2-ENGINE-1727

The installed maps use word-65538 checks as the constants of comparisons; the registers named by actions 3, 8 and check 19 lie above
every check register (their role is not read). — R2-ASSET-375

| Measure | Count | Claim |
|---|---|---|
| Word-65538 checks; slot 0 of type 1 / type 7 | 164; 160 / 4 | R2-ASSET-375 |
| Word-65538 checks named by a trigger; values | 153; 0..1229 | R2-ASSET-375 |
| Trigger rows whose right side is a word-65538 check | EN 1341, RU 1337 (every row) | R2-ASSET-375 |
| Register operand p0 of actions 3, 8 and check 19 above the map's check registers | 48, 197, 208 (registers 70..83) | R2-ASSET-375 |
| Check 19 whose p0 is an ordinary check register | 1 | R2-ASSET-375 |
| Largest number of check registers in a map | 54 | R2-ASSET-375 |
| Action nodes of word 65538 (slot types 5 and 6) | 46 | R2-ASSET-375 |

### Node and trigger fields the game reads

The type-7 loader reads the counts, each node's seven fields and each 184-byte trigger, and tests no value but the
allocation result. The compiler reads of a node are the operation word, the ID, the ten type codes and values, the
node pointer for its failure message, and for a check the word at file +0x48, which it copies to IR+0x74; no reader of
IR+0x74 is found in the listings searched (Medium). Name and text fields reach only the failure messages
(R2-ENGINE-1681); the type-0 +0x28 count is overwritten (R2-ENGINE-1324); the trigger fields are R2-ENGINE-1316. A check
that fails to compile, such as a type-9 value naming no structure, gets the failure message and register 0 and takes
no index; a trigger naming such a check reads register 0. — R2-ENGINE-1728

| Measure | Count | Claim |
|---|---|---|
| Word +0x48: action nodes 0 / check nodes 0 / check nodes 1 | 1334 / 824 / 162 | R2-ASSET-375 |
| Check / action nodes; duplicate, zero or above-5000 IDs | 986 / 1334; 0 | R2-ASSET-375 |
| Triggers; flag 0 / 1 | EN 715, RU 714; EN 130 / 585, RU 129 / 585 | R2-ASSET-375 |
| Rows by comparison code 0..5 | EN 661, 159, 96, 24, 12, 389; RU 660, 158, 95, 24, 12, 388 | R2-ASSET-375 |
| Rows naming a check ID no node has | 1 | R2-ASSET-375 |
| Triggers whose first left check ID is 0 | campaign 0; EN loose 4, RU loose 11 | R2-ASSET-375 |

### Victory and defeat

A map states its outcome with two actions. Action 4 adds one to the win counter; the result tick then sends victory
when no failure is pending (Medium as the player sees it, R2-ENGINE-624). Action 5 stores p0 as the failure reason; a
reason above 0 is sent as defeat and selects the mission's failure text (Medium, R2-ENGINE-629), and p0 = 0 stores no
defeat. Check 18 also adds one to the failure reason. An actor byte +0x13c state sends reason 2 without a script node
(Medium; its meaning is Unknown, R2-ENGINE-624), and a companion's death sends its own reason (R2-ENGINE-625).
— R2-ENGINE-1729

| Measure | Count | Claim |
|---|---|---|
| Action 4 nodes; named by a trigger | 46; 46 | R2-ASSET-375 |
| Action 5 nodes; named by a trigger | 81; 75 | R2-ASSET-375 |
| Check 18 nodes; named by a trigger | 1; 0 | R2-ASSET-375 |

### Rectangular areas

Among the check bodies read (Medium), the script has one rectangle test: check 2 with unit U in its type-4 slot and the scalars left, top, right, bottom in
p0..p3 sets its register to 1 when U's x lies in p0..p2 and its y in p1..p3, bounds included, compared signed. x and y
are bytes +0 and +1 of U's position object (Medium that they are tile coordinates). No other read check takes four
corner scalars; action 25 writes one cell event for the single cell (p2, p3) and shapes no rectangle
(R2-ENGINE-1775). No script operand kind of the 46 campaign maps is a rectangle record (R2-ENGINE-1685), and the type-12 music
areas are a centre and a radius read by music (R2-ASSET-084); other map record kinds were not read for rectangles
(Medium). — R2-ENGINE-1730

| Measure | Count | Claim |
|---|---|---|
| Nodes of checks 2, 3, 22 or 27 in the 46 campaign maps of either root and the loose `.alm` files (EN 13, RU 37); no other install archive counted | 0 | R2-ASSET-375 |
| Operations with point operands (types 5 and 6) | checks 7, 14, 15, 16, 25; actions 21, 24, 25, 29 | R2-ASSET-375 |
