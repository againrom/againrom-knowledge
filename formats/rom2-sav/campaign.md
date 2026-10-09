# Native campaign record and selected client tail

The native scenario record has a fixed raw prefix and a counted reference
selection. Its writer and reader use the same widths, but the list reader
reconstructs membership from catalog pointers rather than restoring every
incoming record byte. — R2-SESSION-031, R2-SESSION-032, R2-SESSION-033

## Scenario record

| Order | Stored representation | Native relation |
|---|---|---|
| 1 | 4096 raw bytes | 1024 DWORD slots copied from/to one bank |
| 2 | 320 raw bytes | 4 by 4 cells; each transfers 20 bytes in row-major loop order |
| 3 | u32 count | Count of writer availability-list nodes |
| 4 | count records of u32 ID, u32 kind, raw 16 | ID is runtime +4, kind is runtime +0, raw payload is runtime +8 |
| 5 | u32 index | Writer compares each node's object pointer with the current pointer; default is zero |

The native export names identify the availability list and current pointer.
They do not assign meanings to the bank, grid or 16-byte payload. These producer
cones and arbitrary live values remain Unknown. — R2-SESSION-031, R2-SESSION-032

## List restoration and current

After reading the raw bank and grid, the reader clears availability and rebuilds
the catalog pointer list. It reads each 24-byte record into a temporary. Kind 1
selects a catalog kind-1/ID search; every other kind selects a catalog kind-2/ID
search. Matching catalog pointers are appended only when that pointer is absent.
The incoming raw 16-byte payload is not copied into the selected catalog objects.
Incoming order, duplicate records and unmatched keys can therefore produce a
different list count or current ordinal. Complete valid membership is Medium
because catalog-object and allocation producers remain unexamined.
— R2-SESSION-033

The writer's zero-default index does not distinguish an unmatched current pointer
from a match at the first node. The reader advances a cursor until its counter
equals the incoming index, then dereferences that cursor to assign current.
There is no local count/range/null test on this selection path. Native read,
exception, allocation and invalid-input responses remain Unknown; this is not a
claim that a malformed file is accepted or crashes. — R2-SESSION-034

## Departure transition

The ScenarioLeaveLocation export stores -1 to its output DWORD, returns when
current is null, and routes record type 2 to a type-2 routine and every other
type to the ordinary routine. The type-2 routine removes current from the
available list only for ID 1 and clears current; it writes no output and no bank
slot. The ordinary routine is one body with a 12-case switch on IDs 10..110 and
a nine-case switch on slot 768 values 30..110. All other selectors reach the
join without a case body. — R2-ENGINE-143, R2-ENGINE-144, R2-ENGINE-145

Ordinary departure adds available records by ID: 10 adds type-1 ID 20, 20 adds
type-2 ID 2 and, when slot 772 is nonzero, type-1 ID 21, 31 adds type-1 ID 32,
40 adds type-1 IDs 50 and 60, 50 adds type-2 ID 3 and, when slot 780 is nonzero, a fixed record whose type
and ID are unresolved, and 60 adds type-1 ID 80. The
add helpers append a catalog pointer only when no node holds it. The output
DWORD is 1 for ID 10, 2 for ID 30, 3 for IDs 70 and 80 only under opposite
slot 777/778 gates, and 5 or 4 for ID 110 by slot 779. — R2-ENGINE-146,
R2-ENGINE-147, R2-ENGINE-148

A nonzero slot 775 first clears availability and appends 42 unconditional fixed
records and one of four records chosen by slots 776 and 781. The types and IDs behind those records are not
resolved. The earlier controller/list closure excluded two allocation wrappers;
their selected local contracts are now measured. Their failure callbacks,
small-block internals and aliasing remain open. — R2-ENGINE-149,
R2-ENGINE-150, R2-ENGINE-218

Every ordinary departure zeroes slot 773, normalizes 532+i against 512+i, zeroes
512+i, removes current, clears current and stores 1 at slot 896+ID before the
ID switch. The case stores are per ID. IDs 70 and 80 set slots 777 and 770 and
slot 778 on both paths of their output gate, which skips only the output 3 store.
Slot 768 gains ten when the ID is divisible by ten,
whether or not the ID has a case body. — R2-SESSION-047, R2-SESSION-048,
R2-SESSION-049

The stage and ID cases store immediate values into an eight-entry, 20-byte-stride
array after the bank; its meaning and consumers are Unknown. The store at
slot 896+ID has no ID bound check, so IDs of 128 or above would address outside
the bank; whether such IDs occur is Unknown. — R2-SESSION-050,
R2-SESSION-051, R2-SESSION-052

## Inn stage, options and town 2 TALK

Catalog kind2/ID2 has a measured constructor, but a catalog ID is not an
EnterInn stage. The client passes an option-array pointer and count pointer.
EnterInn reads bank768 for every catalog record. All four measured local
EN/RU call setups use this ABI, guarded by mode member+0x5d8/+0x63c equal
to 2. Complete mode producers and runtime dispatch remain Unknown.
— R2-ENGINE-159, R2-ENGINE-215, R2-SESSION-060, R2-SESSION-071,
R2-SESSION-072

NewGame returns 768=10. Ordinary departures with IDs divisible by ten
increment it by ten; type2 departures do not. On the named fresh path,
Leave10 gives20, Leave20 gives30 and admits town2, and Leave30 gives40.
An admitted town2 revisit after mission30 therefore uses stage40.
Restored bank, arbitrary writers and the complete event schedule remain
Unknown. — R2-SESSION-049, R2-SESSION-107

The export retains its ten stage case bodies and shared continuation.
Stage30 returns kind3 NPC22/topic30, kind3 NPC2108/topic31 only when
slot927=0, and kind0 NPC2110/topic39. The stores fill the caller buffer
and increment its count; they add no catalog record. Kind3 TALK later
searches catalog type1/ID-topic and admits matching pointers to
availability only if absent. NPC22/topic30 also sets533=1 and 553=2;
NPC2108/topic31 has no special bank store; NPC2110/topic39 kind0 makes no
state change in the selected export. — R2-ENGINE-147, R2-ENGINE-161,
R2-ENGINE-216, R2-ENGINE-219

The selected client copies packed words in order. Kinds1/2 enter one actor
vector; other kinds, including 0/3, enter its talk vector. The TALK action
finds the first word matching the selected actor's low16 key, formats
npc%dtalk%d, dispatches text, then sends the same full word to TalkTo.
The stage30 keys are npc22talk30, npc2108talk31 and npc2110talk39.
Rendered rows, actor lookup success, dialogue alternatives and event
scheduling remain Unknown. The independently measured ID2 music arm is
separate from this option contract. — R2-ENGINE-160, R2-ENGINE-220,
R2-SESSION-059

The fifteen fixed continuation stores use these predicates:

| Option topics / NPC | Kind | Predicate |
|---|---:|---|
| 79 / 5 | 0 | currentID3; signed768>60;774=1 |
| 78 / 675 | 0 | currentID3; signed768>60;770=1 |
| 74..77 / 2022 | 3 | currentID2;959!=0; all970..973=0 |
| 84..87 / 2022 | 3 | currentID2; any970..973!=0; all980..983=0 |
| 93..96 / 2022 | 3 | currentID2; any980..983!=0; all989..992=0 |
| 62 / 22 | 3 | currentID2; signed768>=60;958=0;949!=0 |

For each NPC2022 family, flags776/781 zero/zero select74/84/93,
zero/nonzero select75/85/94, nonzero/nonzero select76/86/95 and
nonzero/zero select77/87/96. These families have no stage compare.
Each kind3 option requests available type1/ID-topic on TALK; kind0 adds
no catalog record. — R2-ENGINE-217

The three named callees do not interpret these packed options. D2.00011
reads record ID. D2.00035/00036 allocate/release storage blocks for list
nodes through selected heap paths. EnterInn calls only the ID getter.
Allocation callbacks, small-block internals and aliasing remain Unknown.
— R2-ENGINE-218

The published action populations for maps 10/20/30 contain0/1/0 instant35
nodes; the map20 node conditionally writes 772=1. Their 4/6/8 instant36
nodes target753..756. None of these selected nodes targets a fixed inn
completion family. Known NewGame/Leave10/Leave20 writers keep those
families zero, so that named stage20/30 path admits no fixed continuation
option. This Medium result does not cover every live bank or later visit
at the same stage. — R2-SESSION-108, R2-SESSION-109

The separate dynamic tail scans i=0..19. It emits kind1/2, low ID i+1,
when552+i=currentID and 532+i is 1/2, with bit31 iff signed512+i>0.
An admitted EnterInn call with currentID2 immediately after Leave20 and
without an intervening writer emits kind1/ID1. The selected topic30 TALK
also enables kind1/ID2 under the same call and writer conditions. The whole
continuation is therefore not empty on that stage30 call. — R2-ENGINE-221

The initial kind3/topic10 TALK also writes 769=1. The former
availability-rather-than-bank clause of R2-SESSION-023 is partially
retracted; its NewGame and Leave facts stand. — R2-SESSION-110

## Remaining inn bodies and pre-50 visits

The six other distinct stage bodies emit 25 packed options: 23 kind3
mission offers and two kind0 dialogue offers. These stores fill the caller
buffer. A later kind3 TALK requests type1/ID equal to topic; NPC is an actor
key. Kind0 has no admitted catalog type or ID in the selected paths.
The table gives native store order and every extra gate. A nonzero
predicate compares against zero; all admitting branches fall through
except the topic82 bank776=0 arm, which takes its equality branch.
— R2-ENGINE-223, R2-ENGINE-224, R2-ENGINE-225, R2-ENGINE-226,
R2-ENGINE-227, R2-ENGINE-228, R2-ENGINE-229, R2-ENGINE-230

| Stage | Kind | Topic / NPC | Predicate beyond dispatch | Claim |
|---|---:|---|---|---|
| 40 | 3 | 40 / 22 | None | R2-ENGINE-223 |
| 40 | 0 | 48 / 2108 | None | R2-ENGINE-223 |
| 40 | 3 | 41 / 2015 | 937=0 | R2-ENGINE-223 |
| 40 | 3 | 42 / 2111 | 938=0 | R2-ENGINE-223 |
| 40 | 3 | 43 / 2004 | 939=0 | R2-ENGINE-223 |
| 50 | 0 | 49 / 23 | None | R2-ENGINE-224 |
| 50 | 3 | 51 / 2 | 533!=0;947=0 | R2-ENGINE-224 |
| 50 | 3 | 53 / 2019 | 949=0 | R2-ENGINE-224 |
| 60/70/80 | 3 | 70 / 680 | currentID3;771!=0;966=0 | R2-ENGINE-225 |
| 60/70/80 | 3 | 71 / 681 | currentID3;771!=0;966=0;967=0 | R2-ENGINE-225 |
| 60/70/80 | 3 | 83 / 677 | currentID3;768=80;979=0;536!=0 | R2-ENGINE-225 |
| 60/70/80 | 3 | 61 / 2006 | currentID2;768=60;957=0 | R2-ENGINE-225 |
| 60/70/80 | 3 | 63 / 2109 | currentID2;768=60;959=0 | R2-ENGINE-225 |
| 60/70/80 | 3 | 73 / 2004 | currentID2;768=70;969=0 | R2-ENGINE-225 |
| 60/70/80 | 3 | 72 / 2108 | currentID2;768=70;968=0 | R2-ENGINE-225 |
| 60/70/80 | 3 | 81 / 2010 | currentID2;768=80;776!=0;977=0 | R2-ENGINE-225 |
| 60/70/80 | 3 | 82 / 2009 | currentID2;768=80;776=0;978=0 | R2-ENGINE-225 |
| 90 | 3 | 90 / 2004 | currentID2 | R2-ENGINE-226 |
| 90 | 3 | 91 / 681 | currentID3;987=0 | R2-ENGINE-226 |
| 90 | 3 | 92 / 2003 | currentID3;988=0 | R2-ENGINE-226 |
| 100 | 3 | 100 / 2006 | currentID2 | R2-ENGINE-227 |
| 100 | 3 | 102 / 2109 | currentID2;998=0 | R2-ENGINE-227 |
| 100 | 3 | 103 / 2005 | currentID2;987!=0;999=0 | R2-ENGINE-227 |
| 100 | 3 | 101 / 681 | currentID3;997=0 | R2-ENGINE-227 |
| 110 | 3 | 110 / 2006 | currentID2 | R2-ENGINE-228 |

Topic71 inherits966=0: a nonzero 966 skips both 70 and 71. The shared
body's stage60/70/80 label alone does not satisfy its current-ID or bank
gates. These local native predicates are High; actual campaign bank values
and scheduling remain Unknown. — R2-ENGINE-225

All 23 remaining kind3 topics call the same TalkTo AddMission arm. With
the measured initial recordID10 they write no additional bank slot:
bank769=1 requires a topic equal to that initial record ID, and none of
these topics is 10. TalkTo has no direct stage store. Its kind1 arm toggles
bank[511+low16]; inn-emitted kind1 low IDs1..20 bound it to slots512..531,
outside stage slot768. Kind0 topics48/49
also write no bank and add no catalog record; the selected client route
uses npc2108talk48 and npc23talk49. Native effects are High, composed
presentation is Medium, and a displayed dialogue was not observed.
— R2-ENGINE-229, R2-ENGINE-230

With incoming DWORD stage s, ordinary Leave30 gives s+10 modulo 2^32;
Leave31 and Leave32 preserve s. This local algebra is High. The common
prelude stores into bank[512+i] and bank[532+i] for i=0..19 and bank[896+ID].
For IDs31/32 these address slots512..551 and 927/928, never stage slot768.
— R2-SESSION-111

Named fresh paths before playing50 permit these visits. This is a Medium
positive reachability result, not an exhaustive live stage set. Town1 is
removed on its type2 departure; the named stage20 point has no available
town. Leave20 admits town2. Leave40 admits both 50 and 60, so postponing 50
allows Leave60 then Leave80 to reach60 then 70 while town2 remains available.
With no intervening stage writer, Leave31/32 after the named Leave30 path
would retain 40; this composition is not a generated named visit. The
generated postponed-side-mission path retains 70 through Leave31/32.
— R2-SESSION-112, R2-SESSION-113

| Admitted town | Bank stage | Inn stage body | Named writer |
|---|---:|---|---|
| type2/ID1 | 10 | stage10,D2.00207 | NewGame |
| type2/ID2 | 30 | stage30,D2.00208 | Leave20 |
| type2/ID2 | 40 | stage40,D2.00146 | Leave30 |
| type2/ID2 | 50 | stage50,D2.00155 | Leave40 |
| type2/ID2 | 60 | shared60/70/80,D2.00162 | Leave60 while postponing 50 |
| type2/ID2 | 70 | shared60/70/80,D2.00162 | Leave80 while postponing 50 |

The measured constructors identify type2 IDs1,2,3. Leave50 unconditionally
calls add-town(3); nonzero 775 restoration also appends town2 and town3 regardless
of stage. Its complete775 producer set is Unknown; no pre-50 town3 visit
was observed. Any admitted town uses global bank768 dispatch, with its
current ID changing inner predicates. Topic70 in the shared body needs
currentID3 and 771!=0; Leave50 is a named771=1 writer. — R2-SESSION-113

The full 46-map corpus per locale contains four authored instant35 and
140 instant36 nodes. Their literal targets are 772/780/779 and 753..756;
none targets768 or 775. This bounded negative is High. Indexed SetVar,
packet-key writers, aliases, raw restoration and other operations remain
outside that population. A complete writer/availability graph or authorized
original observation would settle the exhaustive pre-50 state set.
— R2-SESSION-114, R2-SESSION-112

## Selected client tail

After the store, the selected EN/RU writer and reader arms transfer the pair
record, scenario record and nine buffer records in that order. Explicit native
module loading and ordinal-12/13 bindings connect these arms to ScenarioSave and
ScenarioLoad. The selected reader arm requires client mode value 2 and a nonzero
argument. Other routes and complete store-value production remain Unknown.
— R2-SESSION-037

The pair record is u32 count, count pairs of u32, and two final u32. Runtime rows
are 60 bytes, but the selected pair helpers transfer only runtime +0 and +4.
The reader resizes the array and reads those fields. The other 52 bytes, the row
constructors and all field meanings remain Unknown. — R2-SESSION-035

Each of the nine buffer records is u16, u16, u32 byte length and that many raw
bytes. The runtime cell is 12 bytes: the first two WORDs, length at +4 and a
buffer pointer at +8. The buffer pointer itself is not emitted by this helper.
The reader's nonzero old-length arm clears the previous buffer; nonzero incoming
length selects allocation and a raw read. Buffer grammar, producer cones and
complete length validation remain Unknown. — R2-SESSION-036

These contracts do not establish a complete live-World serializer, owner-save
acceptance, post-load gameplay or presentation. — R2-SESSION-033,
R2-SESSION-034, R2-SESSION-035, R2-SESSION-036, R2-SESSION-037

## Four measured physical tails

In four original save points, exact label-end framing gives inline stores of
742, 1050, 742 and 938 bytes for symbols A, B, C and D. Record counts are
22, 28, 22 and 28; pool lengths are 10, 126, 10 and 14. Every store is followed
by one pair `(1,1)`, two zero suffix DWORDs, the scenario record, and nine
buffers whose two WORDs and lengths are all zero. The ninth buffer ends at EOF.
Registry values, pair meanings, grid and raw availability payloads remain
Unknown. — R2-SESSION-076

| Symbol | Document mission | Incoming availability kind/ID | Incoming current index | Nonzero bank slots |
|---|---:|---|---:|---|
| A | 0 | 2/1 | 0 | 768=10 |
| B | 10 | 1/10 | 0 | 754=1, 768=10, 769=1 |
| C | 0 | 2/1, 1/10 | 0 | 768=10, 769=1 |
| D | 10 | 1/10 | 0 | 768=10, 769=1 |

These stored keys support a bounded city-like A/C and mission-like B/D
classification through the selected prior location grammar. Incoming current
index zero does not prove a native current-pointer match or catalog restoration.
No filename-to-owner-description mapping, dialogue, chronology or causality is
established. — R2-SESSION-077

A/C's decoded documents are byte-for-byte equal despite different inline-store
hashes, bank slot 769 and availability. B/D differ at bank slot 754 and the first head DWORD, with
additional opaque document and registry differences. The slot values are
preserved; no complete producer, actor state or progression meaning is assigned.
— R2-SESSION-079
