# Sacks, death and loot

[Reference](format.md)

## Sacks

| Off | Type | Meaning |
|-----|------|---------|
| +0x04 | u32 | runtime id, from `MOVE-ID-016`'s bitmap |
| +0x10 | ptr | position |
| +0x1c | i32 | total value = gold + `Σ item+0x1c`; recomputed by `R0937` |
| +0x3c | i32 | gold |
| +0x40 | ptr | a container of the same class an actor holds |

**One sack per cell.** `R0472(pos, container, gold)` looks the cell up first; when a
sack is already there it pours the incoming container into it and adds the gold. A new sack
that fails to register at its cell is deleted and the call returns 0.

The local registration key is `Sack+10 -> Position+02`, not Position's other
cell word. Dynamic bit 0 rejects before lookup. An existing nonzero Sack slot
rejects even when it already holds the supplied Sack. An existing empty slot
is written without plane recomputation; creation instead zeroes 52 bytes,
captures current Cost/Static, sets the record-present flag, stores the Sack
and recomputes. — SAV-SACKENTRY-590

Local removal refuses only a missing record before its write path. It clears
payload+10 without an occupied-slot or pointer-equality test, recomputes, and
tests the four occupant slots, layer count+02 and operation+2c for deletion.
Other residue and the layer pointers themselves do not retain the record.
— SAV-SACKREMOVE-591

Deletion restores Cost and Static baselines, preserving Static bit 4, but does
not replace Dynamic with restored Static or clear Dynamic bit 5. Dynamic keeps
the preceding recompute result, with the conditional bit 4 OR. This is a local
write-set, not proof that allocator or higher-level callbacks preserve every
other field. — SAV-SACKPLANES-592

The lookup used by merge/create first requires Static bit 5 and then a matching
node. Registration success dispatches the collection append; refusal dispatches
the deleting-destructor slot. Two known removal-caller slices ignore the local
return before subsequent collection/transfer calls. Complete merge, destructor,
transfer and exceptional effects are separate from these established cell
transitions. — SAV-SACKCALLER-593

**A sack does not tick and does not expire.** Its `vt+0x14` and `vt+0x18` are empty stubs and
it is never inserted into the actor tick list — it is created, merged into, and destroyed by a
pick-up. Nothing ages it.

A pick-up takes **everything**: the gold is credited to the looter's `Player`, the sack's whole
container is poured into `actor+0x7c` one unit at a time, the sack is unregistered and deleted.
There is no per-item take from a sack in the protocol, and no capacity, distance or ownership
test at execution.

The measured gold command `0x23` requires a resolved Player, its first actor
and positive amount not above that Player's purse. It debits Player `+38`
before calling the Sack merge/create wrapper with a null container. The local
arm ignores a zero wrapper return and contains no refund on it; external
rollback and client reconciliation remain Unknown. A gold-only new Sack has
a fresh empty container; an existing Sack adds to `+3c` and recomputes `+1c`.
Pickup credits positive `+3c` to the looter's owner Player (MENU-094).

## Ground-sack drawn frame

The [backpack sheet](../spr256/format.md#backpack-sheet-frame-selection-spr256-077) has a
client-side frame field, `CBackPack+0x20`, separate from the Sack's own `+0x1c`. A client
message dispatcher's opcode `0x7a` case writes it: it copies a message byte into
`CBackPack+0x20` and clamps it to 0..5. A dispatcher-wide scan finds five further
stores of a register value into a `+0x20` field outside this case, one already published as a different
message's Building field and four of unverified object class, so this case is shown to write
the field, not shown to be its only writer. — `ITEM-136`

The producer found for that opcode is a generic notify routine's Sack branch, reached only
after a null/non-null recipient split, a null-object exit, and a virtual call through a
vtable slot where only a zero return proceeds (Sack's own slot there always returns zero) —
this is not a direct `IsKindOf` gate at entry, the class test comes after that vtable check.
Past the gate it reads the recomputed `Sack+0x1c` (the same field the Sacks table above
lists) and writes `_ftol(log10(Sack+0x1c))` — not a literal threshold table — into the
message byte the dispatcher copies. Neither weight nor a contained-item count is read
anywhere in that routine. A per-byte-offset scan of the whole `.text` section for both a
direct-immediate and a register-mediated write of that opcode found one direct site (this
routine's own) and no register-mediated candidate; neither scan can catch a value computed by
arithmetic or loaded from a table, so this is the sole producer found, not a proof there is
only one. — `ITEM-137`

Because `_ftol` truncates toward zero rather than flooring, and `log10(x)` for `x` on
`[1, ∞)` is non-negative, truncation and flooring agree over that whole range — but "the
input is non-negative" does not by itself exclude zero, and `Sack+0x1c = 0` is a real point in
the domain. `log10`'s own zero-input branch does not compute an ordinary logarithm: it
discards the operand and returns the extended-precision negative-infinity bit pattern with an
internal error code 2; a negative-input branch returns a quiet-NaN pattern with error code 1.
That branch is reachable at the instruction level rather than excluded by type: the load is
a signed 32-bit integer read, so a negative operand is representable. Whether the
recomputed field can ever hold one is bounded by `ITEM-SACK-010`'s own recompute and was not
verified here. What `_ftol` produces from
either error value, and hence what frame a zero-value Sack draws, was not traced.

The mapping ascends in six bands at the powers of ten — `Sack+0x1c` in `[1, 10)` gives index
0, `[10, 100)` gives 1, `[100, 1000)` gives 2, `[1000, 10000)` gives 3, `[10000, 100000)`
gives 4, and `100000` and above clamp to 5 — a ladder at fixed decimal thresholds, not a
smooth curve, and the same `fldlg2`/`fxch`/`fyl2x`/`_ftol` idiom the shop screen's price
plaque selector uses for its own digit-count index (`SHOP-SCREEN-037`). No sign flip, table
reversal or subtraction from a fixed frame count sits anywhere between the `fyl2x` result and
the clamped store, so within this ladder a larger recomputed value never yields a lower frame
index. Whether a higher index also draws a visibly *bigger* sack sprite is a separate,
unmeasured question: it depends on the backpack sheet's own frame ordering, which no claim in
this repository establishes, so the index direction is decoded fact while the visual-size
direction remains an owner report this experiment did not verify against the sheet. This
mapping's reach is also bounded by `ITEM-137`'s own scope: it describes the producer that scan
found, not every path that could write the opcode. One concrete trigger path from a value
change to this notifier call is read end to end: the death/drop wrapper `R0944`
through the shared create/merge routine and the sack registry lookup to that notify call
(`TRIG-DROPALL-024` independently reaches the same wrapper chain for its own purpose); whether
every value change reaches the notifier by some other path was not traced. — `ITEM-138`

## Death

`R0208`, in order:

```
1. state := 16; leave the world
2. unequip actor+0x78 into the container
3. unequip actor+0x74 into the container ONLY IF the weapon's Data.bin parameter 15
   ("sutableFor") != 0 -- i.e. unless the weapon suits no class at all
4. actor->vt+0x44()  -- the STRIP. Empty body on the base actor; on both humanoid
   classes R0981: for i = 1..12 inclusive, unequip actor+0x198+4i into the
   container. A null slot costs nothing: vt+0x40 returns 0 and the append returns
   before storing. This is step 4 and everything below it sees the result.
5. suppress := templateName contains "NPC"   OR   (multiplayer AND Player+0x5c != 0)
6. if suppress:  DELETE the container outright -- with the armour already in it --
                 and install an empty one
7. gold := 0; if typeID > 0x40 and rand()%100 < param[0x26]:
               gold := param[0x27] + rand()%param[0x28]
8. if container is non-empty OR gold != 0:
       no existing Sack: the new Sack ADOPTS the container object itself
       existing Sack: DRAIN into its container, deleting the source container
9. the corpse is given a fresh empty container
```

The container identity is retained only by the new-Sack adoption branch.
Strip/reinsertion and an existing-Sack drain apply the whole/split/merge rules;
Weapon unequip also changes its owned Spell. `ITEM-DEATH-012` (its unconditional
container-identity sentence is retracted), `ITEM-GROUNDMOVE-130` and
`ITEM-SPELLMOVE-132` distinguish these outcomes. A mercenary (`NPC%02d_%d`)
leaves nothing at all. Because the strip precedes the
emptiness test, a body that wore anything leaves a sack even if it carried nothing.

**Death is not overridden per class.** `vt+0x18`, the tick that reaches `R0208`, is
`R0037` on all three actor vtables. Only `vt+0x44` differs, and its empty base body is
structural rather than stylistic: `Unit` is `0x198` bytes, so the loop's first read on a base
actor would be four bytes past the end of the object.

**Disposal is a different path from dropping.** `Humanoid::~Humanoid` (`R0982`) also
walks `1..12`, but calls `worn[i]->vt+0x04(1)` -- the deleting destructor -- and never touches
the container. On a normal death it finds twelve nulls, because the strip ran first.

## Sack producers

Five owners, complete on the repaired function table:

```
R0944   the wrapper used by death and by both drop arms
R0061   command 0x23, at the requested cell and at the hero's cell
R0946   the mission .ini's "Items" list, at map load  (no .ini ships)
R0945   random treasure at map load, and a multiplayer top-up to W*H/400
R0461   .alm type-8 authored loot, at map load
```

The repaired `EnumRefs` reference-manager and `.rdata` populations contain three direct
packed-item-factory owners: `.alm` type-8 load, script instant 12 and a transient UI formatter that
immediately destroys its object. The sack maker has six direct calls in five owners; the death
inventory routine has one caller. Computed targets remain the bounded blind spot of those image
reference populations (`ITEM-PRODUCER-091`).

## Authored staffs, Dragons and death

Every Human equipment cell containing `castSpell` is a slot-0 Weapon cell: 46 per root, belonging
to 36 ordinary template names and ten with the exact uppercase substring `NPC`. Every underlying
weapon is `Staff` or `Shaman Staff`, both `sutableFor = 2`. Across the 28 campaign maps, 76 Human
placements resolve to those definitions per root. Death moves the same worn Weapon into the actor's
container before producing a sack for 49 non-`NPC` placements; it destroys that container for the
27 `NPC` placements. This split is for the shipped single-player campaign maps. In multiplayer,
`Player+0x5c != 0` is a second suppression input; loose-map non-`NPC` rows therefore leave death
eligibility unevaluated (`ITEM-AUTHCAST-086`, `ITEM-AUTHDROP-087`, `ALM-MODE-070`;
`ITEM-DEATH-012`, whose container-identity sentence is retracted; its suppression inputs stand).

Dragon rows 112–115 author `Flame Thrower` in slot 0, leave slot 1 empty and
have no castSpell. `Flame Thrower` has `sutableFor=0`, so the death gate does
not move it into the container. No Dragon staff or enchanted-item death
producer is established in type 8 stock, same-cell ground elements,
item/container script targets, the shared death body, or its enumerated
direct/reference callers. This negative is limited to those families.
Flame Thrower is innate-weapon state. — ITEM-DRAGON-088, ITEM-DRAGDROP-089
