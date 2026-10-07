# SAV campaign record

[Format reference](format.md) · [Application state](application.md)

The application campaign at `+0x548` follows the `&YA1` store on the same
file. Writer `R1294` and reader `R0434` use the same counted sequence.
The scalar suffix is not the entire campaign; arrays and markers make its
size variable. — SAV-CAMPTAIL-070, SAV-CAMPPROG-071

## Wire programme

All counts in this record are **plain u32**, not MFC `Count`. Source offsets
below are relative to the campaign record, or the child/base record where
specified. Pointer/count pairs describe in-memory storage, not wire pointers.
— SAV-CAMPPROG-071

The base record is written once at the head and once per child:

| Order | Wire | Source | Meaning |
|---:|---|---|---|
| 1 | Six u32 | `+04,+08,+0c,+10,+14,+18` | Mission, MapObject, Payment, shop lower/upper bounds, announce latch |
| 2 | u32 count, count × u16 | Count `+24`, array `+20` | AddHero collection at `+1c` |
| 3 | u32 count, count × u16 | Count `+38`, array `+34` | EnableMercenary collection at `+30` |

After the main base record:

| Order | Wire | Count / payload source | Meaning |
|---:|---|---|---|
| 1 | u32 count; each child = base + u32 age | `+50 / +4c`, child age `+48` | Side-mission children |
| 2 | u32 count; count × u16; count × u16 | `+64 / +60` then `+74` | Working / pristine per-type mercenary counts |
| 3 | u32 count; count × u32 | `+8c / +88` | Per-type hire flags |
| 4 | u32 count; count × u16 | `+a0 / +9c` | Current main mission's Mercenaries shelf |
| 5 | u32 count; count × u16 | `+b4 / +b0` | Permanent mercenary unlocks |
| 6 | u32 count; count × u16 | `+c8 / +c4` | InnNPC |
| 7 | u32 count; count × u16 | `+dc / +d8` | InnMission |
| 8 | u32 count; count × u16 | `+104 / +100` | TCMission |
| 9 | u32 count; count × u16 | `+f0 / +ec` | ShopMission |
| 10 | u32 count; pairs of u32 value, u32 kind | `+14c / +148` | Documents: kind 1 text, kind 0 picture |
| 11 | Seven u32 | Table below, in listed order | Scalar suffix |
| 12 | u32 count; counted markers | `+160 / +15c` | Selected-mission marker cache |

The two mercenary count arrays share one serialized count. InnNPC/InnMission
are semantically paired but have separate wire counts. TCMission precedes
ShopMission on the wire despite its higher memory address. With every count
zero, the complete record is 104 bytes. — SAV-CAMPPROG-071,
SAV-CAMPAIGN-076, SAV-CAMPAIGN-077, SAV-CAMPAIGN-078, SAV-CAMPAIGN-079,
SAV-CAMPAIGN-080, SAV-CAMPAIGN-081, SAV-CAMPAIGN-082, SAV-CAMPAIGN-083,
SAV-CAMPAIGN-084, SAV-CAMPAIGN-085, SAV-CAMPAIGN-086 (its clause that load
rehydrates the picture pointer is superseded by SAV-609: world-map entry does)

## Scalar suffix

| Order | Runtime source | Meaning / restoration |
|---:|---|---|
| 1 | `+118` | Selected mission |
| 2 | `+114` | Raw value; both application load paths increment it into server difficulty `+84` |
| 3 | `+110` | AutoGetMission |
| 4 | `+11c` | LastMission |
| 5 | `+120` | First-MapPoint flag, computed on SAVE |
| 6 | `+124` | Accumulated groups of 16 simulation sub-ticks at admitted mission completion |
| 7 | `+128` | Received hostile corpse-stage transition counter; exact condition below |

`+120=1` restores MapPoint zero. Otherwise selected mission `+118` resolves
through the external mission-to-map-object registry. This stores a relation
to campaign map data, not an X/Y coordinate. `View/X` and `View/Y` in YA1 are
separate application fields. — SAV-CAMPPOS-072, SAV-892

A census of `+120` over 54 organically-numbered original mission saves finds it
set on 22 and not a pure function of the selected mission: two saves share the
same selected/main mission with opposite flag values. The writer and reader
instructions for the field are Unknown; no existing claim traces them.
— SAV-1099

`+114/+128` are raw passthrough, unlike computed `+120`. Constructors/reset
zero them; the local campaign reader tail does not consume them again.
The `+114` singleton chain reaches the three-way ALM placement law. Neither
field is spare storage. — SAV-598, SAV-599, SAV-600, SAV-601, SAV-602,
UNIT-GATE-012, UNIT-GATE-013

The admitted completion arm adds signed trunc(simulation sub-ticks/16) to
`+124` before its terminal-score predicate. This establishes groups of 16
sub-ticks, not wall-clock seconds or the exact mission-entry clock baseline.
Constructor/reset zero `+124/+128`; SAVE/LOAD preserves both raw words.
The named arithmetic and reader impose no positive-value or range clamp.
— FAME-021

The nonzero `+128` writer increments when client state reception changes a
hostile drawable from old stage0/1 to stage2/3/4. Stage1 is the death/fall arm;
stage5 takes another arm. The increment checks the local CPlayer relation row,
with no killer or prior-ID predicate. New CUnit objects begin at stage0, so
first receipt of an existing corpse can satisfy this condition; repeat 2-to-3
cannot. The loaded counter baseline and subsequent native reconstruction order
must not be replaced by a once-per-death assumption. — FAME-022

`+110` = 0 in a mission-10 document leaves the party on the world map after
the quest; with 20, the value every original mission-10 save carries, the
next mission starts. — SAV-1094

`+110` = 0 still takes the same `!= -1` fork as a real mission number and calls
the routing loader (`R0755`) with request 0; against any nonzero current
main-progress record this is the documented lower-request no-op (`+110`'s own
sole reader is monotone, above), so the call itself writes nothing. What the
routing fork's own reader does next with that zero return — travel to the
current mission's own map object, or take the same city/object-0 arm the `-1`
default uses — is Unknown; neither reading is excluded by an observed world map
with nothing further, since both leave the party there. — SAV-1104

## Markers

```text
u32 value
u32 n             writer uses strlen(string)+1
n string bytes    includes terminating NUL
u32
u32
```

A marker occupies `17+strlen(string)` bytes. The seven scalars plus marker
count occupy 32 bytes only when marker count is zero. Both serializer arms
include the nonzero-marker grammar. The LOAD contract is narrowed to restoring
these bytes: picture-pointer reconstruction runs separately on world-map entry,
not synchronously in either SAV loader. — SAV-CAMPMARK-073, SAV-609

## Progress and collections

| Event | Located campaign effect |
|---|---|
| Ordinary main progression | Retains active main record; rejects a lower main request |
| Side-mission completion / age 2 | Removes its child record |
| Accept town candidate | Shortens its building array; keeps InnNPC/InnMission paired |
| Announce candidate | Sets the record's own announce latch |
| Complete mission | Drains EnableMercenary into permanent unlocks; clears hire flags |
| Activate town | Drains AddHero |
| Acquire document | Accumulates `(value,kind)` pair |

There is no completed-mission collection in this grammar. The marker cache is
world-map presentation state; surviving heroes are objects in the Player
roster, not campaign records. — SAV-CAMPAIGN-076, SAV-CAMPAIGN-077,
SAV-CAMPAIGN-078, SAV-CAMPAIGN-079, SAV-CAMPAIGN-080, SAV-CAMPAIGN-081,
SAV-CAMPAIGN-082, SAV-CAMPAIGN-083, SAV-CAMPAIGN-084, SAV-CAMPAIGN-085,
SAV-CAMPAIGN-086 (load-rehydration clause superseded by SAV-609; the
presentation-state classification stands)

The base record's `+0x04` mission scalar is not written through the archive's
insertion operator. `R1631` loads the archive vtable's `+0x40` slot and
calls it with the field's own address and the literal length 4; `R1632`
mirrors it through the `+0x3c` slot with the same pointer and length. Neither
routine touches a Player. The other progress-bearing field of a save,
`Player+0x3c`, is written by a different serializer into a different half of the
file, so progress is split across two routines and a consumer that implements
one restores half of it.

Both calls are indirect and the vtable's contents are not resolved, so raw write
and raw read are role names here. What supports them is that the store routine
takes one slot and the load routine the adjacent one with identical arguments,
and that the actor stat-span helpers branch on the same store/load test and
reach a matching pair of primitives by direct call.
— SAV-1027, SAV-CAMPPROG-071, SAV-CAMPAIGN-077

## Mercenary state

Mercenaries `+9c` is the current main mission's shelf, not a cumulative union.
The tavern reader consumes the stored array directly; it does not recompute
membership from the mission number. Mission changes can replace rather than accumulate this shelf. — SAV-606

Hire sets its type's hire-flag dword. Mission entry does not clear that flag;
mission end does. Zero flags mean no outstanding hire, not that no hire ever
happened: named producers are hire, dismiss, reset and mission-end zeroing.
— SAV-615, SAV-616, MERC-HIRE-003

Actor class/type/equipment signatures do not encode hire origin. The same
Human shape can come from hiring, authored map placement or scripted transfer.
Creation-order IDs are session-local rather than a provenance tag.
— SAV-617, SAV-622, SAV-629

The working pool at `+60` is the tavern's own headcount storage, and both
tavern readers — the shelf filter and the inn's count/price preview — read the
same element of it. The preview also reads the pristine array at `+74` into a
second value. The server's per-head price factor is a different storage: an
element of the command `0x38` vector. Nothing inside the tavern writes the
pool; hire and dismiss touch only the hire flags at `+88`, and the pool's
elements change at the record reset, at mission end and on LOAD.
— MERC-POOL-011, MERC-POOL-012

SAVE reads the count once from the working array's `m_nSize` at `+64` and
writes both payloads under it. The pristine array's own `m_nSize` at `+78` is
never read, so a second count between the two payloads is a shape this writer
cannot produce and a reader must not expect one. — SAV-1084

LOAD sizes both arrays from the document's count and reads both payloads out of
the document. Neither campaign load driver calls the record reset or the record
constructor, and neither reads `[General] MercenaryCount`; the registry
sections they do read are Objects and Projectiles, both pushed as literals,
and the General section is pushed nowhere on either driver. A restored pool is
therefore the document's own, not a recomputation from campaign state and not a
constructor default, and a serialized count of 0 leaves both arrays empty with
a NULL `m_pData` that nothing on the load path repopulates. — SAV-1085

Across 119 admitted saves over the four ROM1 roots and the preserved owner
saves, the serialized element count is 15 without exception, and the two arrays
are equal in every file but one. Corpus agreement of that kind does not exclude
a producer the corpus never exercises. — SAV-1086

## Tavern eligibility and selection

The stored main-mission shelf is narrowed to permanently unlocked types before
the tavern builds its offers. The client then collects live CUnit objects of
those types with a nonzero unsigned working-pool word at `type-1`. These are
separate conditions; shelf membership
does not check the pool index's range. A type must have a corresponding
working-pool entry. — MERC-SHELF-002, SAV-928

| Input / consumer | Rule |
|---|---|
| Campaign shelf `+9c`, count `+a0` | Supplies candidate types |
| Permanent unlocks `+b0`, count `+b4` | Filters the shelf before client collection |
| Working pool `+60`, count `+64` | Collector tests the unsigned u16 at `type-1` |
| Client document map | Zero map count returns an empty result; nonzero count traverses buckets/next links even for an empty eligible set |
| Live map entries | Require a finite coherent map, acyclic stable links, valid objects and returning services |
| Mercenary selection on inn activation | Stores 0 for a nonempty result, -1 for an empty result |
| Caption / price refresh | Signed checks `selection <= count-1` / `selection < count` admit -1; neither is a lower-bound check |
| InnNPC | Separate append loop; it does not supply an empty-mercenary fallback selection |

The persisted `InnNPC` order is the talk-only cells' order, at roster
position `mercCount + j`. The campaign record names which mercenary types are
offered but not their cell order: that comes from the ids of stock units built
after load, which the record does not store. — SAV-1111, SAV-1112

The local empty-selection stores and indexed consumers are established.
Whether selection/storage survive intervening UI/resource calls unchanged
remains conditional. Do not infer a universal empty-save failure from the
negative-index condition. — SAV-928, SAV-929

Saved eligible types do not prove that live stock was published, selected and
kept valid through entry. Full original click order and the first failing
consumer remain Unknown. The [Player command prerequisite](player.md#tavern-command-prerequisite)
is a separate relation. — SAV-932

## Terminal campaign

Nonzero LastMission with selected mission divisible by ten reaches an earlier
mission-end branch to score production and message `0x428`. It locally
bypasses ordinary reward/advance and enters the terminal UI branch.
This does not establish a final usable city or terminal save.
— SAV-890

The terminal handler adds the UI child at frame+11c to the manager at frame+cc.
Constructor/vtable bindings select credits activation, then manager display;
after those returns the handler sets mask40 in frame+3dc. This frame+11c is
a UI pointer, distinct from campaign LastMission+11c. Resource success,
mutable child/focus receivers and native arrival remain separate boundaries.
— SAV-970

A delivered terminal close notification clears mask40 and reads LastMission
through frame+664. Nonzero requests the FAME screen, whose activation sets
mask1000. Its selected close notification clears that bit and requests the
menu message. That message requires the entire mask to be zero before the
campaign reset and subsequent menu activation. Reset clears named campaign
fields and requests the ordinary initial mission loader; its argument is not
a terminal number. Earlier callbacks, remaining mask bits and queued
delivery prevent an unconditional native credits-to-menu conclusion.
— SAV-971

The selected F2 control requires mask0/1 and mode2; the SAVE-dialog opening
message independently requires mask0/1. Retained credits/FAME masks therefore
inhibit these controls. A matching SAVE-dialog close result445 reaches the
SAVE wrapper, but a separate message442 also reaches that wrapper without a
local mask test. Its normal terminal producer/order is unproved. Actual
terminal SAVE availability, current or retained campaign/documents, and
acceptance of a later native LOAD remain Unknown. — SAV-972

The inner registry loader rejects a request above
`10*ScenarioMissionCount` before campaign stores/grants. Its outer higher-request
caller still writes current `+04` and selected `+118` and reports success.
Those numbers can therefore change with older registry/documents retained;
the terminal dispatcher may bypass the path entirely. An inferred increment
does not supply a successor campaign. — SAV-891

Campaign SAVE/LOAD has no separate terminal grammar or local LastMission
rejection. On the selected missing-map lookup, default object one becomes
zero-based index zero. Later callbacks and UI remain acceptance boundaries.
Native terminal-city LOAD remains Unknown. — SAV-892

## Extension boundary

Existing counts admit more entries of their existing shape. Added fields,
changed widths or reordered fields are outside the original fixed programme;
there is no version-selected campaign grammar arm. Preserve the programme or
version an extension outside it. Stored `+114/+128` cannot be repurposed as
padding. — SAV-CAMPPROG-071, SAV-599, SAV-600, SAV-602
