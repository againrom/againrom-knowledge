# SAV containers, Items and equipment

[Format reference](format.md) · [Object programmes](objects.md) · [Human state](human-state.md)

## Container

The container is a separate CObList-based object, not Item. Sack serializes
its `+40` container directly; Unit serializes its `+7c` container when its
presence byte is nonzero. — ITEM-CONT-004, SAV-670

| Wire | Runtime source | Meaning |
|---|---|---|
| u32 count, Item references | Container list | Current contents in list order |
| u32 | Container `+1c` | Insertion index |
| u32 | Container `+20` | Incrementally maintained running load |

Constructor values are insertion index 10000 and load 0. Generic insert uses
`index<count ? insert-before : AddTail`; that default therefore appends.
Ground-pickup ordering and an equip arm can change the insertion index.
A transfer helper copies both dwords; a reset pair zeroes both. Neither LOAD
nor SAVE recomputes the running load. Container `+1c/+20` are not Item's
numerically coincident price/Effect-list fields. — SAV-670, SAV-671, SAV-672,
ITEM-CONT-004, ITEM-LOAD-005, ITEM-STACK-003 (its enchantment-separator clause
is retracted; the running-load sum stands), SAV-CITYSTORE-516

## Item Effect list

Item's `+20` list is typed as Effect and admits Effect_DirectDamage by
derivation. Located Item producers are the `Effects=` grammar and the
magic-shop merge-or-append generator; the latter uses Item `+1c` price as its
budget. Actual production by other paths remains Unknown.
— SAV-774, SAV-775 (its clause placing the live damage-spell path outside
both located writers' call graphs is partially retracted; the writers and the
class bound stand), ITEM-EFFGRAM-070, SHOP-EFFALT-071, SHOP-MAGIC-007

Those two located Item producers use the default constructor, and their own
writer bodies store no nonzero value in Effect's trailing `+0c` byte. The live
attached-spell identity stamp uses different call sites. Whether an Item
producer reaches those sites through a callee is Unknown: SAV-776's former
spell-builder caller chain is partially retracted, and SAV-910 supplies no
whole-builder graph in its place. Do not infer attached-spell identity from an
Item Effect merely sharing the class programme. — SAV-776, SAV-910,
MAGIC-ATTACH-016, MAGIC-DMG-005

## Transfers and nested identity

| Operation | Located state effect |
|---|---|
| Whole extraction + nonmerge insert | Retains Item pointer; inspected list/container bodies contain no Token/Position store |
| Merge | Retains destination Item, adds counts, ORs incoming Token `+08`, deletes incoming Item |
| Pickup before merge | Sets incoming `+08=1`, so retained flags become `oldFlags OR 1` |
| New Sack | Adopts supplied container and constructs a separate Sack Position |
| Existing Sack | Drains and deletes supplied container |

Publishing an Item's compact record clears its Token `+08` pickup flag after
sending it as bit `0x40`. Actor entry reaches that writer for carried and
equipped Items subject to class, recipient ownership and type gates. An Item
it does not publish keeps a saved 1 through session entry. No other clearing
writer was located. — SAV-1114, SAV-1115, SAV-POSTLOAD-222

Quantity-one draining can split a stack. Death reaches both Sack branches.
No branch establishes copying Sack/actor Position into contained Items;
actor callbacks and aliases remain outside the local no-write result.
— ITEM-WHOLE-128, ITEM-MERGE-129, ITEM-GROUNDMOVE-130, ITEM-DEATH-012 (its
unconditional container-identity sentence is retracted)

Armor/Shield/Weapon equip bodies write actor slots/stats. Base Item equip can
consume/delete the Item and change an Effect mode. Transitive callback writes
to the full Token head remain Unknown. — ITEM-EQUIPMOVE-131

Weapon equip with a first kind-41 Effect reconstructs its owned Spell from
the Effect ID. Unequip deletes that Spell and clears Weapon `+80`.
For a valid nonzero ID:

| Spell field | Producer |
|---|---|
| `+08` | Effect ID |
| `+09` | Spells parameter 6, low byte |
| `+0a` | Spells parameter 18 equals 1 |
| `+0c` | Spells parameter 1, low word |

The Spell programme stores these fields and its new identity key. Same Item
identity does not preserve nested Spell identity; numerical address reuse can
hide a new allocation. — ITEM-SPELLMOVE-132

## Definition binding

The saved row byte `+0c` selects the definition entry `+3c` at every load.

| Class | Collection | Range check |
|---|---|---|
| Item | `L04587` | Compares with the collection size |
| Armor | `L04589` | None: `[coll+4] + 0x3c*row` |
| Shield | `L04591` | None: `[coll+4] + 0x3c*row` |
| Weapon | `L04593` | None: `[coll+4] + 0x3c*row` |

Weapons has 28 runtime entries. Row 0 is never serialized from `Data.bin`,
and row 23 carries no parameter array. The Weapon compact-record arm reads
parameter 15 of the entry without a guard. The resume route after the document
has been read projects the actor at Player `+34`. When that actor's `Unit+0x74`
Weapon is on row 0, the original terminates with an access violation at
`L08027`. Which Human that actor is, and so which other Humans' slots are
reached at resume, is Unknown. A row of 28 or more indexes past the array.
Whether another routine rejects such a row before first use is Unknown.
— SAV-1088, SAV-1087, SAV-1089

Across 119 original saves, the 2471 Weapons carry rows 2–26 only; none carries
row 0 or a row of 28 or more. No original Armor or Shield carries row 0. In one
generated document, changing only the `Unit+0x74` Weapon's row byte from 0 to
13 moved the fault from `L08027` to `L04760`, the Armor compact-record
arm, reached from the Humanoid equipment array. That document also holds an
Armor on row 0. A written Weapon therefore names a row from 1 to 27 that has a
parameter array, and correcting the Weapon row alone does not make such a
document load. — SAV-1089

An Armor on row 0 in the Humanoid equipment array faults the same way, at
`L04760` in the Armor compact-record arm, reached from walker `R0669`'s
loop over `actor+0x198+4i`. In one generated document the only such Armor sat
at index 7; changing its row byte from 0 to 15, together with the Weapon's,
let LOAD reach the map. No original Armor carries row 0. — SAV-1091

## Equipment order

| Path | Local order |
|---|---|
| Armor removal | Clear actor slot before Effects |
| Shield removal | Effects before slot clear |
| Weapon removal | Effects; direct fields/derive; delete owned Spell; final slot clear |
| Weapon attach | Prepare Spell before displacement; derive before timing/range and weight |
| Armor/Shield attach | Store slots and add defence before positive-weight refresh; Armor flags before Effects, Shield flags after |

Weapon saved byte `+50` is separate from its definition binding. The current
byte, reread after derive, supplies the minus-one delta added/subtracted at
actor range byte `+12c`; arithmetic wraps, so removal need not invert arbitrary
prior state. Melee's active selector uses the low byte of definition parameter
5 cached after displacement and slot store. Ranged kinds 11/12 and removal
assign zero. — SAV-EQUIPORDER-552, SAV-HUMEQUIP-447, HERO-EQUIP-017,
ITEM-ARMFOLD-033 (its common flag/Effect order clause is retracted; the
per-class order above replaces it)

State-0 Effect apply/remove walks forward and derives after each normal
general dispatch. Its iterator saves a next-node pointer but later reads that
node's payload/link, so transitive mutation is not excluded. Command22 adds
container operations and a final zero-weight refresh, which derives only when
the quotient changes. Humanoid/Human wrappers add no trailing derive; Unit
wrapper `+54` is Token value, not derive. Atomic equipment snapshots, callback
purity and the state observed by an actual SAVE remain Unknown.
— SAV-EQUIPEFFECT-553, SAV-EQUIPCALL-554, SAV-EQUIPOBS-555, ITEM-EQUIP-006
(its universal wrapper and recompute clauses are retracted)

## City producers

Carried-to-table transfer updates container bookkeeping without a local
actor-load refresh. Sale completion and individual table-to-inventory return
refresh load and call Human derive only if the signed16 load/capacity quotient
changes. Bulk cancellation uses another return loop. Successful school
purchase changes skill/base/XP and unconditionally derives.
— SAV-CITYMOVE-512, SAV-CITYSALE-513, SAV-CITYRETURN-514, SAV-CITYDERIVE-515

SAVE transfers current Human spans, scalar words, XP and container tails in
their stored order. It also assigns Human `u32 +148 = u8 +14c`. Embedded
serializers and intervening callbacks remain open; preserving raw fields does
not imply that SAVE is wholly free of mutation or supplies a safe initial
state vector. — SAV-CITYSTORE-516

## Attached Effect source and actor credit

Ordinary Effect construction clears transient source `+44`; copy preserves
it; continuous same-ID refresh retains the old source while replacing the
counter. A separate template failure arm does not write `+44`, without an
established allocator value or ordinary failure-path reach. — SAV-906

Unit `+20` stores typed Effect references. A newly read Effect runs its factory
then loads a 44-byte body which omits `+44`. A back-reference aliases the
existing loaded object. The selected post-load entry repairs Position but
the selected Unit/Human lifecycle does not reconstruct that attached source.
Saved remaining duration and transient source identity are separate.
— SAV-907

Nonzero Token8 periodic HP payload reads Effect `+44` after changing HP.
Null skips source dispatch; negative source HP clears it; zero/positive HP
admit source virtual `+64`. Death instead reads Unit `+40` and dispatches
source virtual `+48` before award virtual `+60`. Null periodic source can
coexist with a mapped credited actor. Destruction, intervening callbacks and
first native-consumer chronology remain Unknown. — SAV-909, SAV-910

For the selected source classes, Unit `+64` is empty; Humanoid/Human `+64`
keeps the Effect source as the progress `+5c` receiver. Conditional cases
with different Effect source and victim credit reach that source's actual
progress stores. Their numerical calculations and notification returns are
supplied cuts, so a first native post-LOAD award remains Unknown.
— SAV-958, SAV-961

The death path's first source callback can enter Player mutation and
notification before credit is reread for `+60`. Whole notification effects
are unresolved; an unchanged source across that boundary cannot be assumed.
The removed victim's own attached-list loop separately clears each reached
Effect source before its teardown tick, without proving that every Effect
which refers to the victim is cleared. — SAV-959, SAV-962
