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
resolved. The measured closure has no unresolved direct callee other than two excluded
unread callees, and no indirect
edge in the controller or list helpers. — R2-ENGINE-149, R2-ENGINE-150

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

## Second town entry and TALK

The catalog record kind 2 / ID 2 has one constructor, D2.00016, which
initializes object D2.00018; it is one of 52 constant-argument constructor call
sites, identical in EN and RU. Whether the constructor is also reached by
registration or by a save-loaded record is Unknown. — R2-ENGINE-159

The EnterInn export has case bodies for stages 10, 30, 40, 50, 60, 70, 80, 90,
100 and 110 and no dedicated body for stage 20. Stage 20 selects the shared
continuation D2.00019, which seven case bodies jump to and the stage-110 body
falls into. The continuation is 295 instructions and holds 15 packed-entry
stores behind compares of bank slots: kind 0 NPC 5 topic 79, kind 0 NPC 675
topic 78, kind 3 NPC 2022 topics 74..77, 84..87 and 93..96, and kind 3 NPC 22
topic 62. Whether a stage-20 call satisfies those compares is Unknown. Stage 10
and stage 30 hold packed TALK entries (kind, topic, NPC) with compares of bank
slots between them. The stage that the client passes for ID 2 is Unknown.
— R2-ENGINE-161

One client body has a separate arm per ID 1, 2 and 3, selected by a compare chain
on the record member at +4; the arm for ID 2 pushes a music resource key in both
locales, and indirect calls in that body are unresolved. The EN text member has
one section each for NPC 22 topic 30 and NPC 2108 topic 31, which match the
stage-30 packed entries by key scheme. The choice-action producer, the RU
sections and the first-town arm key are Unknown. — R2-ENGINE-160,
R2-SESSION-059

The selected first EnterInn caller compares EN member+0x5d8 and RU
member+0x63c with 2; its unequal direct branch skips the local call setup.
This extends the earlier unmeasured RU clause for this caller only. The other
three RU windows, the member's meaning and changes across preceding unread calls
or aliases remain Unknown.
— R2-SESSION-060, R2-SESSION-072

The final local setup pushes frame-0x14 and then frame-0x94, leaving ECX at
the second address. It produces no scalar catalog or stage value. The full
ABI, buffer contract, indirect switch destinations, all-path entry and effects
of preceding callees or aliases are Unknown; no catalog/stage relation or runtime
route is established.
— R2-SESSION-071

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
