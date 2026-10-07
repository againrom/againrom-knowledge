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

Scalar indices below are compiled indices, not original catalog slot indices.
The meanings are scoped to the verified RU/EN executor paths.

| Arm | Literal operands | Actual effect or predicate | Authority |
|---|---|---|---|
| Instant 35 | scalar 0=index, scalar 1=value | SetVar(index,value) | R2-ENGINE-045 |
| Check 23 | scalar 0=index | GetVar(index) into map result register | R2-ENGINE-045 |
| Check 24 | scalar 0=n | GetVar(752+n) into map result register | R2-ENGINE-045 |
| Instant 36 | scalar 0=n, scalar 1=mode | Mode-dependent stores to 752+n and notifications 253/254 | R2-ENGINE-046 |
| Instant 37 | scalars 0..7 from X,Y,int×6 | Bad instant default, optional diagnostic packet 0x91, ordinary cleanup/return | R2-ENGINE-058 |
| Check 25 | scalars 0=X,1=Y,2=layer | Nonzero terrain cell effect-layer pointer | R2-ENGINE-053 |
| Check 26 | resolved unit, scalar 0=effectID | Nonzero active effect bit 1<<(effectID&31) | R2-ENGINE-054 |
| Check 27 | resolved unit, reads scalars 1/2 | Tile X/Y equality plus fractional bytes 128/128 | R2-ENGINE-055 |
| Instant 38 | item u16+0x6c | One extraction call per iterated unit, returned-item destruction and unit refresh | R2-ENGINE-056 |
| Instant 39 | resolved group+0x60 | Clear forced-activity DWORD in the group's record | R2-ENGINE-057 |

Instant 36 mode 1 stores 1 only from old 0. Mode 2 leaves old 4 unchanged;
otherwise it emits 253 only if old bit 2 is clear, then stores 3. Mode 4 always
emits 254 and stores 5. Other modes store their value directly. Objective rows
read752+row, omit zero values, and render bits 2/4/1 in that priority using
frames 11/12/10. Generic35/36 and EnterLocation reset are known writers;
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
and UI 430 respectively. The EN result tick and full acknowledgement scheduling
are Unknown. — R2-ENGINE-060

DLL EnterLocation stores the location-record pointer and clears slots 752..767.
GetCurrentLocation returns a pointer; the client wrapper reads its numeric ID.
The ordinary LeaveLocation mission10 arm marks completed-bank slot 896+10,
adds location20 and returns output 1; that selected case has no 753/754 guard.
Other continuation arms, persistence and complete native campaign playthrough
remain Unknown. — R2-ENGINE-048

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
