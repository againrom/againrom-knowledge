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

Bank slot 770 occupies prefix bytes 0xC08..0xC0B. The selected paired
writer/reader transfers the saved DWORD verbatim before catalog rebuilding.
NewGame's reset and later campaign writers are separate operations. Native
cold LOAD and post-load mutation remain Unknown (`R2-SESSION-154`, High
for the transfer and offset).

Town 3 has record ID 3 and kind 2. Its writer node contains these two
DWORDs followed by raw 16 bytes; current remains a list ordinal, not a
fixed value 3. The named fresh route offers towns 1, 2, then 3 on Leave50;
nonzero bank 775 can restore towns 2/3 independently of stage. The route
is conditional and full membership remains Medium (`R2-SESSION-155`).

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
40 adds type-1 IDs 50 and 60, 50 adds type-2 ID 3 and, when slot 780 is nonzero, type-1 ID 52, and 60 adds
type-1 ID 80. The
add helpers append a catalog pointer only when no node holds it. The output
DWORD is 1 for ID 10, 2 for ID 30, 3 for IDs 70 and 80 only under opposite
slot 777/778 gates, and 5 or 4 for ID 110 by slot 779. — R2-ENGINE-146,
R2-ENGINE-147, R2-ENGINE-148, R2-ENGINE-419

A nonzero slot 775 first clears availability and appends 42 unconditional fixed
records and one of four records chosen by slots 776 and 781: towns 2 and 3,
missions 21, 30..32, 40..43, 50..53, 60..63, 70..77, 80..87, 90..92 and
100..103, 110, and one of 93..96. — R2-ENGINE-419 The earlier controller/list closure excluded two allocation wrappers;
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
Rendered rows, actor lookup success and event scheduling remain Unknown. The independently measured ID2 music arm is
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

## Inn hire, fire and talk choice

TalkTo has two kind 0 arms and tests no stage, bank or current record in
them. NPC 5 at any topic stores 536=1, 556=3 and 774=2. NPC 0x2a3 with
topic 0x4e stores 770=2 and ORs 0x80000 into the record-3 category-3 mask.
Kind 1 toggles bank 511+NPC; kinds 2 and 4..7 store nothing; bit 31 is not
read. — R2-ENGINE-493

EnterInn makes 48 option stores: 46 constants and two tail stores. Of the
seven kind 0 constants, 5/79 reaches the NPC 5 arm and 675/78 the 0x2a3
arm; 207/9, 2108/8, 2110/39, 2108/48 and 23/49 reach none. — R2-ENGINE-494

| Action | Kind 1 entry | Other entry |
|---|---|---|
| Hire | text `npc<ID>accept<stage>`, set bit 31, TalkTo | text `npc<ID>reject<stage>` only |
| Fire | clear bit 31, TalkTo, no text | clear bit 31, TalkTo (no store for kind 2) |
| Button 1 on a hire seat | text `npc<key>about` | text `npc<key>about` |

The stage is bank 768. The table lists what the Hire and Fire bodies do
(High; whether the text is shown Medium, R2-ENGINE-343, R2-ENGINE-051). On
the generic inn page `L2.00744`, the center pointer routine `L2.00755` after
its hit slot returns nonzero, and button 0 on release, run Fire when the
copied word has bit 31, else Hire. EN and RU agree after masking addresses
and register names, except that RU skips the about text when the selection
is negative. The Kaarg and druid inn pages, the event class and the callers
of `L2.00755` are Unknown. — R2-ENGINE-495, R2-ENGINE-496

On the generic inn page, hire seats come before talk seats. For a hire seat
whose copied word has bit 31 the painter passes main-text key 257's string
to `L2.00720`; the text and what appears on screen are Unknown. No hire
seat is seated while the selection is -1. The Kaarg and druid painters also
read bit 31 and push key 257 and are not graded. — R2-ENGINE-497

The kind 1 toggle addresses the selection slots 512..531 that the mission
entry reads. Kind 2 marks availability 2, which departure writes for a
selected hire; Hire on it formats the reject text and stores nothing.
Medium for the composed meaning of kind 2: it rests on R2-SESSION-047 and
R2-ENGINE-381. — R2-SESSION-191

Every action takes the first copied option with the seat's low 16 bits, and
constants precede the tail (High). A speaker therefore speaks its first
option (Medium). Medium, from the stage-to-body map of R2-ENGINE-161: only
NPC 2022 can have two options in one call (74..77 with 93..96 when the
current ID is 2, 959 is nonzero, 970..973 are 0, any of 980..983 is nonzero
and 989..992 are 0). Medium: at stage 50, with 533 at 1 or 2, 947 at 0 and
the current ID equal to 553, the topic 51 word with NPC key 2 precedes
hire 2, so Hire on hire 2 formats its reject text. — R2-ENGINE-498

Medium: within one visit of the page, a repeated TALK with the same hero
selects the same alternatives, because each builds a new dialogue from part
1. No inn section header carries `npcalive=`, `npcdead=` or `tips=`; other
header tokens were not enumerated. The file the buffer holds when TALK runs
is Unknown (R2-ENGINE-344). — R2-ENGINE-499

In `text/town.txt` each of the 46 talk keys first hits its own line-leading
tag in EN; in RU 45 do, and `#npc2006talk110` is not at a line start. It holds `about` for hires 1..7
and accept and reject sections by hire and stage. — R2-ASSET-115

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

### Shop parameters and current stock

Town-2 shop parameters occupy relative Scenario-prefix bytes
0x10a0..0x10ef and town-3 parameters occupy 0x10f0..0x113f. Each row holds
four 20-byte category cells after the 4096-byte bank. The selected paired
EN/RU transfer and offset arithmetic are High; these are not absolute
SAV-file offsets (`R2-SESSION-156`).

Current stock persistence, rebuilding and cold LOAD remain Unknown. The
selected owner/base serializers do not directly access the shop pointer,
but another serializer does access the same field offset and raw client
payload producers remain unclosed. No SAV byte range has been identified
as current stock (`R2-SESSION-157`).

A load through the load entry (`L2.00826` / `L2.00119`, reached from the
load page) does not keep the stock. In a town the load entry sends shop
refill 0x3f before the town entry, so all four categories are drawn again
from the restored current location's restored record; in a mission the next
town arrival draws. The `.text` references to the static shop owner are
four per root; the alternate owner serializer `L2.01660` transfers `+0x6c`
and is unclosed. The shop vtable's slots 2..4 are no-op stubs, read as
serialize slots by analogy to the owner table, so no SAV byte is expected to
hold the lists (`R2-SESSION-187`, High for the load order, Medium for the
absence). Callers of the session load that bypass the load entry were not
scanned.

## Conditional party producers

Mission-ID40 writes bank773=23. The selected EN location-leave consumer
passes the request and stage768 to the addition producer; an existing
owner-list identity23 bypasses creation. This conditional link does not
establish exact joining time or an unconditional roster. — R2-ENGINE-379

For hire h=1..20, selection is bank[511+h], availability is bank[531+h],
and town is bank[551+h]. Read mission-ID writers enable1/3/6/4/7 after
IDs20/50/60/70/100, with town gates2/3/2/3/2. NPC22/topic30 TALK enables
hire2/town2; the kind0/ID5 writer enables hire5/town3. These TALK writes do
not construct controlled identity22 or prove recruitment. Stage768 and
mission-ID remain separate selectors. The inn visit that offers
NPC22/topic30 builds identity22 through the inn selector below (Medium).
— R2-ENGINE-380

When app +0x5d8 is 2, the inn selector passes every EnterInn option word
and stage768 to the same addition producer, before the inn page is shown
(Medium). At stage30 the NPC22/topic30 word has low word 22, inside the
controlled range 21..30. The producer builds a Human with the hero flag from
key 10220, the row Fighter_Hero at Medium, stores WORD+0x14c=22 and appends
it to the owner's actor list, unless that list already holds identity22.
Whether the build completes in a town (the map registration calls) is
Unknown, and so is whether this actor is the companion the talk presents.
The NPC22 topic40 and topic62 words repeat the request at their stages.
— R2-ENGINE-633, R2-ENGINE-635

The reached EN entry loop requires nonzero selection and availability,
selects a subset of 20 hire slots, and passes the nonempty list and stage
to the actor creator. The paired creator bodies assign each input ID
to actor WORD+0x14c and register it. The supplied list has no extra local
numeric cap; complete entering membership and placement remain Unknown.
— R2-ENGINE-381

The read EN transition instead selects its reset branch by actor class
WORD+0x0e in 33..63. Other classes take detach/destroy calls. Identity
WORD+0x14c is not that predicate. Unread earlier callees, combat-death
removal, dismissal and return remain Unknown. — R2-ENGINE-382

The scenario SAVE/LOAD bank carries the selection, availability, town
gates and deferred request separately from Player/Group/Unit members.
Bank773 is at 0xC14, selection512..531 at 0x800..0x84F, availability532..551
at 0x850..0x89F and town552..571 at 0x8A0..0x8EF. Native transfer and saved
pointer relations do not establish executed cold LOAD preserving the
same roster or literal binding. — R2-SESSION-158

The hire loop runs only in the mission-entry case whose current location
record has type 1. A pattern pass finds two sites per root that post message
0x468 (Medium for the enumeration): the end of a global-map travel and the
victory arm when no town is available. The two sender bodies are read in both
roots (High). Passes over `E8` calls, `E9` jumps and `.rdata`, `.text`,
`.data` and RU `.st1073` dwords find one creator caller per root, so LOAD
creates no hire among the routes searched (Medium; register calls are not
excluded). EN and RU agree on the entry case and the senders.
— R2-ENGINE-523, R2-ENGINE-533

A hire takes the Humans row with serverID 10000+10*ID, else the Units row,
highest index first; when its family has suffix rows, the row keyed
9998+10*ID+stage/10 (stage bank 768, truncated) replaces it. Installed rows
exist for hires 1..7 only; hires 8..20 get Units index 0. The join is
Medium. — R2-ENGINE-524, R2-ASSET-125

A hire actor carries WORD+0x14c = ID and its row's typeID as class WORD+0x0e
(15, 5, 23, 100, 102, 101 or 26, Medium) in its own Group under the owner the
Players collection's head getter returns (the hero's where the map build
reuses it, Medium). Every hire class is outside 33..63, so the victory
transition detaches hires (Medium). The unit-state sender clears the field
bit that sends +0x14c when Session +0x74 is nonzero (Humans hires always,
Units hires when the key is 0x6000 or more; EN only). The creator writes no
cell. A store census finds no campaign writer that sets Session +0x74, so in
a campaign the bit stays set (Medium). — R2-ENGINE-525, R2-ENGINE-526,
R2-ENGINE-527, R2-SESSION-198

Party placement and positioning bodies are read, and that a campaign hire
reaches them is Unknown. Party placement fixes one start point for the
party: a random word of a drop list (x low byte, y high byte), or x and y
each rand(70)+30 when the list gives 0. It reads its +0x74 and drop list
through the receiver of the join, assumed the Session object (identity not
read). The radius is trunc(max(5, sqrt(n)+4)) for n actors in the owner's
list (Medium; the constants are read in EN only). In list order, each actor
whose returned cell pointer is not itself starts at the point (Medium; RU
calls two routines where EN calls one wrapper, and the behaviour of all
three is not read); the hero is tried at radius 0 first, every other actor,
hires included, at the party radius. A search tries trunc(r*r/2)+2 random
cells within r/2 of the point, y drawn before x in both roots (the range of
each draw is read in EN only), then the square x-major; a failure leaves the
actor unplaced on the last cell tried and logs an error. A join routine
runs this placement when its flag is set and the owner's placed byte +0x41
is 0; the victory transition clears that byte for every owner it keeps. No
read ties a campaign hire to the join: the packet that drives it, its
sender, its dword +0xa and its order against the creator are unread, and
two EN callers of positioning are unclassified. The source of the drop list
is Unknown. — R2-ENGINE-529, R2-ENGINE-530, R2-ENGINE-531, R2-ENGINE-532

Scenario.dll stores availability and town literals only for hires 1..7.
After NewGame hires 8..20 stay unavailable (Medium), so every hire that
enters has template rows. — R2-SESSION-199, R2-ENGINE-524

At mission end a client update of a unit with +0x1dc in 1..20 and signed
health word below -10 sets availability bank[531+ID] to 0 where the update
carries +0x14c (Medium). The ordinary departure then turns a nonzero
availability to 2 when the selection is nonzero and to 1 when it is zero,
the 2 of an unselected hire turns to 1 at the next ordinary departure, and
selection is cleared. A dead hire's availability stays 0 until another
writer of R2-SESSION-048 or R2-ENGINE-380 stores a nonzero value; a
selected survivor rests one departure (Medium as a composition; the
victory-arm link is not read here). Session +0x74, whose nonzero value
would block the update, is 0 in a campaign (Medium). — R2-SESSION-195,
R2-SESSION-198

In RU the hire's Group is, at Medium, in the Player group list a mission
save writes and LOAD rebuilds (the list's caller and collection identity
are unread); the Unit body carries class, owner and hire ID. EN is
Unknown. — R2-SESSION-196

## Catalog and campaign graph

The catalog holds 49 records: towns 1, 2 and 3 (type 2) and 46 missions
(type 1) with IDs 10, 20, 21, 30..32, 40..43, 50..53, 60..63, 70..77,
80..87, 90..96, 100..103 and 110. Each record carries a map x and y; three
further type-1 records with IDs 10, 20 and 30 are built but not in the
catalog. — R2-ENGINE-418

The available list changes only through departure, TalkTo kind 3, NewGame
and LOAD by direct call; EnterLocation sets current only. The stage switch
writes shop records, not the list. One list method reached through a
pointer table stays open. — R2-ENGINE-420, R2-ENGINE-421

| Record | Producer | Condition |
|---|---|---|
| town 2 | Leave20 | none |
| town 3 | Leave50 | none |
| 21 | Leave20 | slot 772 nonzero |
| 30, 31 | inn at stage 30 | 31: slot 927 zero |
| 32 | Leave31 | none |
| 40..43 | inn at stage 40 | 41..43: slots 937..939 zero |
| 50, 60 | Leave40 | none |
| 51 | inn at stage 50 | slot 533 nonzero, 947 zero |
| 52 | Leave50 | slot 780 nonzero |
| 53 | inn at stage 50 | slot 949 zero |
| 61, 63 | inn at stage 60, town 2 | slots 957, 959 zero |
| 62 | town 2, stage 60 or more | 949 nonzero, 958 zero |
| 70, 71 | inn at stages 60..80, town 3 | 771 nonzero, 966 zero; 71 also 967 zero |
| 72, 73 | inn at stage 70, town 2 | 968, 969 zero |
| 74..77 | town 2, any stage | 959 nonzero, 970..973 zero; one by 776/781 |
| 80 | Leave60 | none |
| 81 or 82 | inn at stage 80, town 2 | 776 nonzero or zero; 977 or 978 zero |
| 83 | inn at stage 80, town 3 | 536 nonzero, 979 zero |
| 84..87 | town 2, any stage | one of 970..973 nonzero, 980..983 zero; one by 776/781 |
| 90 | inn at stage 90, town 2 | none |
| 91, 92 | inn at stage 90, town 3 | 987, 988 zero |
| 93..96 | town 2, any stage | one of 980..983 nonzero, 989..992 zero; one by 776/781 |
| 100, 102, 103 | inn at stage 100, town 2 | 102: 998 zero; 103: 987 nonzero, 999 zero |
| 101 | inn at stage 100, town 3 | 997 zero |
| 110 | inn at stage 110, town 2 | none |

Inn rows are kind-3 options that TALK admits. The slot-775 restoration
appends 43 of the 49 records (not town 1, missions 10 and 20, or three of
93..96) and LOAD adds any record whose ID the bank holds. Slot 896+ID is
the completion flag of mission ID; each option whose body tests it (all
inn options except 30, 40, 90, 100 and 110) is withheld once its own flag
is set. Options gated by an exact stage are offered only until the next
decade departure (High for the gates, Medium for that consequence). — R2-ENGINE-424, R2-SESSION-173, R2-SESSION-174

Slot 775 has no writer in the read population except clears and LOAD.
— R2-ENGINE-425

After a campaign victory the client calls LeaveLocation, plays the movie
the output selects, then shows the credits when slot 768 is at least 120,
enters the first available record when no town is available, and
otherwise opens the global map. Outputs 1..5 select the cutpaths rows
teleport, stars, crown, knights and druids. In the read client
call sites a defeat does not call LeaveLocation (Medium). — R2-ENGINE-422, R2-ENGINE-423

## Global map

The 0x41d and 0x42d arms open the global map. It runs in route mode after
leaving town 1 and in select mode otherwise. No map control or key
closes it; it closes itself when a travel ends (Medium for other closers:
message 0x446 and an application arm that sends it were not read). — R2-ENGINE-428

Route mode draws `Umoir.bmp` and select mode `GMap.bmp`. BallMap dots
show the route; Flag, Cross, FlagOnMap, MissionFlag and YourFlag sprites
mark the party, the travel target and records. Flag1 is loaded, but the
record loop that would draw it is not reached. — R2-ENGINE-429, R2-ASSET-093

Each paint runs inside one surface lock. The background is an opaque copy
of the picture converted to the device format at load; BallMap dots skip
converted value 0; the sprites go through the .16a word-run painter with
the sprite's 16-row mode-4 shade table. Cross has 13 frames and its own
palette, byte-equal in EN and RU. — R2-ENGINE-465, R2-ASSET-109

Only GetAvailableLocations records are drawn, at their catalog points,
and only in select mode while no travel runs: the pointed record carries
MissionFlag and the others FlagOnMap; a record at the party point is not
drawn. Route mode and a running travel draw no record marker. — R2-ENGINE-430

The pointer targets the nearest available record by squared distance;
there is no hit area. The route to it is previewed, and the name and
description are shown. A left-button press starts the travel without a
confirmation. — R2-ENGINE-431

The travel extends the dotted route 8 points per paint while Cross
animates at the target, with Point1 and Point2 footsteps. At its end the
map closes, posts 0x468 and calls EnterLocation on the target record (the
first available record in route mode). A click or key skips to the end.
— R2-ENGINE-432

In the bodies read, Point2 starts first; each later start is the other
sound, once the current one has no channel, while the route still grows.
Pointing and the press start no sound in them; the arrival, retarget and
skip bodies were not read, so that nothing else starts one is Medium.
ScrollUp and ScrollDn are loaded and never started in the bodies read.
The map plays `music\map.wav` from its start. — R2-ENGINE-463

The idle message 0x402 paints the map at most once per 100 ms of
timeGetTime, with no catch-up. If the helpers `L2.02239` and `L2.02240`
both return the route's point count L, a travel ends on paint
max(15, floor((L-8)/8)+1); the helpers are unread and other paint
sources were not enumerated (Medium). — R2-ENGINE-464

Routes run between the index-2 pixels of `PathMap.bmp` (100) or
`UmoirPath.bmp` (81) by least summed segment length. In the EN image town 1 and
mission 10 lie on UmoirPath nodes; the other 47 catalog points lie on PathMap
nodes. — R2-ENGINE-433

Names come from `globalmap.txt` sections `#title<ID>` and `#town<ID>`,
descriptions from `#briefing<ID>` and `#description<ID>`; town 1 and
missions 10 and 20 have no section of their own. — R2-ENGINE-434

The lookup is a substring search that takes the first occurrence, so
`#title10` finds `#title100`; mission 10 is shown only in route mode,
which draws no text. A key found nowhere leaves the name empty and the
text block undrawn. — R2-ENGINE-462

EN and RU run the same bodies with the same resource names; the text
files differ in bytes and share their keys. — R2-ENGINE-435

Of 47 presentation body pairs, 43 are equal after masking and four differ
only in application-field displacements. — R2-ENGINE-466

Set-point stores the point it is given, and the 0x42d arm (for a type-2
record) and the 0x41d arm pass the catalog point of the current record:
the town for 0x42d, the won mission for 0x41d; after mission 20's victory
it is mission 20's point, a PathMap node. That every opening does so is
Medium: other openers of the map were not enumerated. Two LOAD bodies
store no party point; the rest of the load is unread (Medium). —
R2-SESSION-177, R2-SESSION-185
