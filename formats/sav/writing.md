# Writing SAV files

[Format reference](format.md) · [Document](document.md) · [Application state](application.md)

A writer must satisfy the byte grammar and the relations between saved
objects. Opaque fields still need their actual state values; the grammar
does not supply a complete set of safe defaults. — SAV-WRITERAUDIT-380,
SAV-WORLDFRONT-432

## Emission sequence

1. Prepare Player/Group/actor and nested-object values. Preserve the intended
   alias graph and distinguish saved-address keys from gameplay IDs.
2. Emit the [document programme](document.md) in order: head, Player graph,
   dead list, shape byte, selected world half and trailer.
3. Include direct base/embedded serializers as well as tagged references.
   Each Player includes Groups, settings32 and Diary; each Unit includes all
   raw blocks and selected optional members.
4. For a world document, use the external ALM baseline and emit the saved
   terrain/cell overlays, terrain key, session4374 and counted world objects.
5. Emit `0xBADFACE1`, its global dword and trailer state400. Pad an odd logical
   endpoint by one byte, compress words and prefix `outWords`.
6. Write magic, placeholder `blobEnd`, version `0x0BAD0002`, `blobBytes` and
   blob. Patch `blobEnd` to the blob endpoint. Append label256, YA1 and the
   complete campaign record.
7. Validate physical/decoded intervals and the relations below. Loading the
   graph does not by itself validate the values consumed by the next action.

Archive indices are assigned by first use. Saved-address definitions need
unique nonzero keys; a reference uses its target's key and that field's
specific lookup timing/policy. The missing-key rule is narrowed by field:
Token's named resolver writes null, while stage-zero order repair retains the
raw key. `SAV-DOC-053`'s 7-of-18 endpoint count, its attribution of every
failure to a Unit subtree and its reading of the 400-byte tail as codec padding
are withdrawn; its top-level order stands. — SAV-FRAME-021, SAV-CODEC-022,
SAV-DOC-053, SAV-FULLREAD-252, SAV-ARCHREL-253, SAV-PTRMAP-035,
SAV-HUMRESUME-460

## Required relations

Preserve actor Diary arrays independently and retain the intended archive
aliases. The typed member reader admits an existing object, so actor ownership
does not require a distinct allocation. A bounded search found no ordinary
actor-owned array consumer, but does not justify omitting its state. — SAV-847,
SAV-982, SAV-984

The fixed six-entry XP tail ends before actor `+1e4`. A conditional index6 in
the located progress expression aliases that Diary pointer word; ordinary
reach is unmeasured. Do not infer an extra safe slot from the expression's
missing local upper check. — SAV-983

| Surface | Writer obligation |
|---|---|
| Envelope | `blobEnd=16+blobBytes`; blob includes u32 `outWords` |
| Word codec | Even decoded length, exact word count, complete opcode operands |
| CArchive | One shared class/object index; correct typed base; intended null/new/alias references |
| Player/Group | Intended owner/hero relations and head-to-tail actor order; early Group references need already-bound targets |
| Position | Matching terrain key for the named world rebind; coordinates/list position do not substitute for identity |
| Terrain | External baseline precedes overlays; cell keys correspond to block rows with static bit5 |
| Spellbook | n-1 reference operations for slots 1..n-1, each representing its intended object or null |
| Unit/Humanoid/Human | Six raw blocks, order list, presence-gated members, Humanoid XP and thirteen references |
| Shortcuts | Canonical kind-6 array holds four dwords /16 bytes |
| World application state | Producer emits Fog and Projectiles; empty Projectiles still includes FreeIndex and IDs |
| Campaign | Exact count/field order; shared count for the two mercenary-count arrays; MapPoint relation; NUL-inclusive marker strings |

A lookup can null a miss or retain its raw word, and some lookups occur only
after other objects load. Do not replace these rules with one generic fixup.
— SAV-GRPLOAD-560, SAV-GRPOWNER-561, SAV-TOKENLOAD-092, SAV-CELLLOAD-112,
SAV-UNITPROG-156, SAV-HUMAN-043, SAV-SPELLBK-041, SAV-ORIGFAULT-335,
SAV-PROJSTORE-428, SAV-CAMPPROG-071, SAV-CAMPPOS-072

## What the original write path drops or mutates

Three places in the write path emit fewer elements than the structure holds.
Two of them are not losses and one is.

The empty-group prune is not a serialization loss. The writer removes, from a
qualifying Player's group collection, every group whose element count is 0. The
argument to the removal is the group; the receiver is reloaded from the Player
in the two instructions before the call, and is the Player's own live group
collection — the same object the iterator is walking. The removal therefore
mutates live memory rather than a serialization copy, memory loses the group
too, and file and memory agree afterwards. A consumer must treat this as a live
mutation of the save action, not as a wire omission.
— SAV-1029, SAV-ROSTER-024, SAV-CITYSTORE-516

The equipment array's skipped index 0 is not a loss. Over the whole image, the
indexed form of that displacement is 18 instructions in 10 owners; exactly two
of them can leave a non-null element, and one refuses a slot number of 0 at its
own first test while the other runs inside the `1..12` loop. Every other indexed
write stores zero. The slot is never filled, so skipping it drops nothing. The
census finds only instructions whose encoded displacement is `0x198`; a write
through a previously computed address would not appear in it.
— SAV-1029, SAV-CARRY-050 (death-container identity clause partially
retracted; the twelve worn slots stand), ITEM-EQUIP-006

The Spellbook's skipped index 0 is writable storage rather than unreachable
storage, so whatever occupies it is dropped, and the difference from the
equipment array is one guard instruction. The array setter has no test of the id
at all beyond one against the current size: a short array is grown to `id+1` and
the element stored unconditionally. The serializer then writes the element count
and emits references from index 1, so an element at index 0 is written as a
count with no reference. It comes back null rather than as residue: the load
arm's own resize allocates the new span and zero-fills it through a call whose
middle argument is the literal 0.

Of the 33 sites that reach the setter, 25 push a literal id inside the spawn
routine and one more pushes a literal outside it; a further site in the spawn
routine takes its id from a table and skips the whole construct-and-store when
the value is 0 or less. Seven sites push a non-literal index, and one of the
seven was traced to that index's origin: the teach-spell effect arm uses the
product of the effect's own magnitude field and the routine's second argument as
the array index, its two preceding guards test only for a missing book and an
already-occupied slot, and a zero factor yields index 0 — not a truncation
artefact. The spell object built on that path is returned unconditionally; its
id resolver takes a zero arm, emits a message and returns normally. The other
six non-literal sites were not traced and neither add to nor subtract from this.
Whether shipped effect data ever supplies a zero factor is Unknown: no
effect-table census has been run, and no preserved save witnesses a Spellbook of
count 1.
— SAV-1029, SAV-SPELLBK-041, MAGIC-BOOK-002, MAGIC-SPELL-001

## Values without a complete authoring contract

| Fields / relation | Known rule | Remaining Unknown |
|---|---|---|
| Human live values and modifiers | Restored literally; named later consumers may read before derive | Complete current-state producers, first consumer and safe residual values |
| Human constructor tails | Named constructors preserve `bc/bd` and clear `fc/fd` | Intervening writes and first-SAVE values |
| Player/Group/Diary raw fields | Defined widths, ownership and field-specific LOAD rules | Safe arbitrary defaults and complete city-consumer prerequisites |
| Tavern campaign lists | [Shelf, unlock and pool filters](campaign.md#tavern-eligibility-and-selection); empty collection selects -1 | Live stock publication, preserved selection/lifetime and full click order |
| Eleven world-head values | Constructor zero, literal restore, named conditional consumers | All-path first-SAVE values and unnamed consumers |
| Position/cell keys | Defined relations and class/path-specific repair | First use on paths without repair; safe unresolved-key placeholders |
| Session and trailer raw state | Exact transfer and named repairs | Unnamed field meanings and valid new-state values |
| Terminal campaign | Same counted grammar; credits/FAME close chain and gated reset/menu; selected F2/SAVE-dialog inhibition | Normal SAVE availability, current or retained campaign/documents, and native LOAD acceptance |

Definition rows describe initial data; accepted current values are not
necessarily derivable from those rows alone. Constructor zero, a zero seen in
a save, or another save's accepted value is not a general default.
— SAV-HUMLOAD-445, SAV-HUMGAPS-449, SAV-HUMFIRST-465, SAV-HUMNEW-505,
SAV-900, SAV-LIVEPROD-413, SAV-WHEADLIMIT-525, SAV-TOKENLOAD-094,
SAV-CELLLOAD-113, SAV-790, SAV-791, SAV-892, SAV-928, SAV-929, SAV-932

The received Diary word cache belongs to a separate client object. Its
opcode186 replacement is a runtime projection, not an additional field in
the serialized Diary grammar. Preserving both saved Diary arrays remains
necessary; constructing a client cache does not prove native LOAD delivery,
first refresh or panel redraw. — SAV-DIARY-042, SAV-995, SAV-998

Terminal UI receiver bindings identify credits and the common manager;
conditional dismissal selects FAME, then a menu message that requires a zero
remaining UI mask before campaign reset. These are local instruction relations.
The selected F2 and SAVE-dialog opening controls admit mask0/1, excluding
retained credits/FAME bits. A separately aligned message442 SAVE call has no
local mask test, so those controls do not prove global SAVE unavailability.
No native terminal output supplies safe current-state values or a later LOAD
acceptance witness. — SAV-970, SAV-971, SAV-972

## Compatibility and extensions

Structurally complete no-world files can load and later fail on the shop
route. The responsible value/consumer contract is not fully identified;
Player/Group/Diary and Human state cannot be declared independently sufficient.
Reducing the party can also leave an authored mission's hero reference
unresolved. Native next-action behavior remains a separate check from framing.
— SAV-ORIGVALUE-399, SAV-LIVEPGD-415, SAV-ORIGMISSION-400

The known descriptor programmes do not define runtime-built or unregistered
derived serializers. Nonminimal primitive encodings have known reader arms,
but do not establish complete malformed-file or allocation acceptance.
— SAV-READPOP-255, SAV-CONTSER-189, SAV-758

Keep original field widths/order and the six Human skill/XP slots. Added
fields require a versioned extension outside these fixed programmes.
Unrecognized YA1 entries can survive a raw registry write but have no general
fresh-application-SAVE preservation guarantee. Label residue, physical suffix
and decoded alignment are likewise not specified extension storage.
— SAV-HEROXP-063, SAV-CAMPPROG-071, SAV-915, REG-102, SAV-EXTSURV-239
