# Containers and equipment state

[Reference](format.md)

## Containers

`0x24` bytes: an MFC `CObList` (`+0x00..+0x1b`) plus two dwords.

| Off | Type | Meaning |
|-----|------|---------|
| +0x1c | u32 | the index the **next** `Add` inserts at; constructed to 10 000, i.e. "the end" |
| +0x20 | u32 | the **load** = `Σ (i16 weight × u16 count)` over the elements |

Homes: `actor+0x7c` (`R0185` at `L04604`), `sack+0x40`. There is **no slot count and
no capacity**: an insert appends whenever `index >= count`, and nothing in the class or in any
arm of the move command compares the count or the load against a maximum.

What the load does, and the only thing it does (`R0280`):

```
load = actor.ownWeight(+0x8e)
if actor.container:
    if container.load < 64000:  load += container.load / 2      (truncated toward zero)
    else:                       load  = 32000
capacity = body*10 + 1                                          (actor+0x92)
if load >= capacity:  speed -= load / capacity;  speed = max(speed, 6)
```

### Whole-object transfer and serialized state

Whole extraction (`qty >= count`) returns the existing Item pointer; only
`qty < count` reaches split. With count 1 and a non-merging destination, the
inspected take/add/list bodies do not store any serialized Item Token or
Position field. This is a bounded direct-body result, not a transitive callback
guarantee. Container draining takes one unit repeatedly, so a larger stack can
still be split. — ITEM-WHOLE-128

A merge requires equal codes and both objects stackable. It retains the
destination Item, adds the incoming count, ORs incoming flags into destination
Token `+0x08`, and deletes the incoming Item. It does not directly copy the
incoming Position or `+0x14` into the retained Item. — ITEM-MERGE-129

Pickup writes each incoming Item `+0x08 = 1` before the drain. A pickup merge
therefore produces `oldDestinationFlags OR 1`. GiveAll and XferItem direct
bodies do not perform that stamp. New-Sack adoption retains the supplied
container; an existing Sack instead drains/deletes it. Sack Position is its
own allocation, copied from the supplied drop Position, not a rewrite of the
contained Item Positions. — ITEM-GROUNDMOVE-130

## Equipment

Fourteen pointer fields, all outside the container.

```
slot 1  -> actor+0x74                       any actor
slot 2  -> actor+0x78                       any actor
slot 3..12 -> actor+0x198 + 4*slot          Humanoid only, gated on actor->vt+0x30()
                                            (false arm: "Error - Trying to takeoff armor
                                             from non humanoid")
```

The UI index in the command is `slot - 1`. **The array is thirteen dwords, `actor+0x198` through
`actor+0x1c8`, and element 0 is dead storage**: the constructor `R0876` clears `0..12`,
while `Humanoid::Serialize`, the death strip and the destructor all run `1..12`, and
`Armor::Equip` refuses a piece whose `Slot` is 0 outright. So no code path can fill index 0 and
none reads it.

**`Armor::Equip` indexes the array with the slot number and does not skip 1 or 2**, so
`actor+0x19c` and `actor+0x1a0` are ordinary slots — written by `Equip`, round-tripped by both
arms of the serializer, stripped on death and deleted by the destructor. The one thing that
cannot reach them is the **take-off command**, which spends those two numbers on `actor+0x74`
and `actor+0x78` instead. No shipped `Armors` row uses either.

**Which slot an item takes is a property of the item, never of the actor or of where the item
was named.**

```
Weapon::Equip  R0850  -> actor+0x74            no refusal arm exists
Shield::Equip  R0849  -> actor+0x78            refuses when shield+0x0c == 0
Armor::Equip   R0848  -> actor+0x198 + 4*part  refuses when armor+0x50 == 0, and when
                                                      actor->vt+0x30() is 0
armor+0x50 = param 4 of the piece's Armors row (R0890 L04716); shipped title "Slot",
             values 1..12; R0890 discards a row above 12 WITHOUT restoring +0x50 to 0
```

`Armors`, `Shields` and `Weapons` share one 18-title array. The equipment
and descriptor paths consume these two slots:

```
slot  4  "Slot"        armor+0x50; which of the twelve fields the piece takes
slot 15  "sutableFor"  a TWO-BIT MASK over consumer classes: bit 0 = fighter, bit 1 = mage.
                       R0984 (Armor), R1032 (Shield) and R1033 (Weapon)
                       each clear the item descriptor byte element+0x08 and copy bit 0 into
                       its bit 1 and bit 1 into its bit 2; the death gate tests the whole
                       value against 0. Shipped: Weapons 0 on three monster attacks (Boulder
                       Thrower, Flame Thrower, Sonic Beam), 1 on nineteen martial weapons,
                       2 on the two staves, 3 on BareHands and Plasma Sword, and -1 on rem,
                       whose 17 parameter cells are all -1, so both bits are set; Shields
                       all 1; Armors 1 on nineteen metal rows, 2 on nine cloth rows, 3 on
                       Amulet and Ring.
```

The `rem` row carries a parameter array. The descriptor arms `AND` a stored `-1` with 1 and 2
without a sign test, and its all-zero material mask builds no item from it (`ITEM-WEAR-056`,
whose former no-parameter-array `rem` clause is retracted; `ITEM-PICT-050`).

### Equipment transfer boundary

Armor, Shield and Weapon retain their Item allocation through the named
equip/unequip slot operations. The inspected item-class bodies write actor
slots/stats, not Item Token/Position fields. Effect apply/remove and actor
recompute callbacks are an unresolved transitive write frontier. Base Item
equip is instead consumption: positive actor HP permits changing zero Effect
mode to 8, Effect attachment, then Item deletion. — ITEM-EQUIPMOVE-131

Weapon equip finds the first kind 41 Effect and, if present, deletes any old
owned Spell and reconstructs one from the low byte of Effect `+0x40`. For a
valid nonzero id, the new Spell's serialized `+0x09`, `+0x0a` and `+0x0c` come
from Spells row parameter 6 low byte, parameter 18 == 1, and parameter 1 low word.
Unequip deletes a nonnull owned Spell and clears Weapon `+0x80`, while leaving
the source Effect in place. No Item split is needed for this nested identity
change. An allocator may reuse the old numerical address. — ITEM-SPELLMOVE-132

### Weapon-borne spell state (`ITEM-CASTSTATE-056`, `+0x68`-only remap clause retracted)

A `castSpell` record remains an ordinary ordered Effect of kind 41. General effect dispatch does
nothing for that kind beyond target recompute. A separate list finder returns the first kind-41
record; Weapon construction uses its spell-id byte to create a distinct owned Spell at
`weapon+0x80`, and Weapon copy/split reconstruct it from id. The list record remains the source of
comparison and persistence while the derived object holds cast runtime state (`ITEM-CASTLINK-076`).

A weapon carrying `castSpell` keeps its constructed `Spell` at `weapon+0x80`. Both release routes
temporarily set `actor+0x64` to that `Spell` and `actor+0x68` to the weapon, then clear both after
the attempt. The non-null item field suppresses mana use and the immediate half-mana training
award. The item effect's `+0x40` is the byte spell id and `+0x42` is a raw signed-`i16` power. The
ordinary staff path decrements no charge and preserves the effect record, `weapon+0x80` pointer and
pointed-to object identity. It does rewrite that `Spell` on every apply: `+0x09` is reloaded from Max
Range and raised by `power/3` for Teleport or `power/30` otherwise, while `+0x0e`, `+0x0f` and
`+0x10` receive the damage base, spread and duration scratch. `+0x09` is serialized; the three
scratch fields are not.
Installed Data.bin castSpell weapons include Human staffs and Catapult/
Ballista Unit weapons. Authored powers are
`{1,5,10,15,25,30,34,35,40,50,60,63,65,70,82,90,98,99}` on each root; 10 and 20 are not an authored
boundary. Runtime stock is larger: the shop generates castSpell weapons with ids
`{1,11,13,14,20}` and price-derived random power capped at 100. The two siege Units can apply their rider, but their plain Unit vtable uses empty cast,
damage and kill award slots, so they gain no training.
`Unit::Serialize` stores `actor+0x68` as an archive object reference and `actor+0x64` as a raw
`u32`. The world LOAD lifecycle explicitly resolves both fields through the
saved-address map: hit replaces the value, miss clears it. This repairs the
raw Spell key as well as the item reference. Later pointer validity and native
resume remain untested. The+68 fixup looks up its then-current word; the
archive-resolved reference is not proved to survive this second lookup.
— SAV-908

A separate post-cast arm destroys an item of kind `0x0e` and its `Spell`; the shipped staff does not
have that kind. Prismatic Spray id 14 is route-dependent: caster admission applies the fan before
the common wrapper refuses the id, while the fighter rider reaches only that refusal. Training is
not stored on the item: the common sink first requires recipient type id in `[0x21,0x3f]`, after
which later damage awards enter the actor's school or current weapon slot.

The selected-inventory Cast UI cannot admit a Weapon despite accepting kind 2 into its vmethod. It
requires descriptor bits `0x10|0x01`; a castSpell Weapon has bit 4 from its effect but its class
descriptor supplies only suitability bits 1 and 2. Bit 0 is a MagicItems-class bit. The server arm
would accept a crafted weapon item order, but no Weapon row can make the shipped UI emit one.

### The wear rule (`ITEM-WEAR-055`, Human-constructor clause retracted; `ITEM-WEAR-057`)

The column decides whether a character may equip or use the item, and the whole rule is one
13-instruction predicate in the CLIENT. The simulation applies none of it.

```
R0882(member, element) =
      (element+0x08 & 2) && !(member+0x18c & 2)      item allows fighter, member is no mage
   || (element+0x08 & 4) &&  (member+0x18c & 2)      item allows mage,    member is a mage

member+0x18c bit 1 = mage. R0590 writes it from (typeID - 0x21) & 2 for a typeID in
   [0x20, 0x40). In non-zero player-character constructor mode, Human builds typeID =
   sex + 0x23 for a mage, sex + 0x21 otherwise (L13008 / L03519), and bit 2 of the
   same dword is sex. Zero-mode map Humans retain their Data.bin typeID; the classifier's
   low-type arm sets the mage bit independently for 0x17 and 0x18 (`PARTY-M20-031`).

element+0x08 = the item descriptor byte the simulation builds and the client reads
   bit 0  0x01  the item is a MagicItems-backed Item, not an equipment piece
   bit 1  0x02  sutableFor bit 0
   bit 2  0x04  sutableFor bit 1
   bit 4  0x10  the item carries an effect of kind 0x29 (castSpell)
   bit 5  0x20  the item has at least one effect; cleared again when item+0x44 == 3 (Potion)

Two call sites, whole image (EnumRefs callto:R0882, 2 hits / 2 owners / 0 orphan):
   R0994 L04836  shop cell background only: false -> backinvg.bmp, the empty-cell
                          background. The icon is drawn either way and the item is still sold.
   R0240 L04839  the equipment doll's drop handler. False sets the refusal flag; the
                          routine returns 0 without reaching this->vt+0x80, the only call that
                          turns the drop into a move. No command 0x22 is produced.

Command 0x22's handler R0061 and all seven routines its equip destination reaches read
no class field and not this column. Their Data.bin parameter reads are 5, 0xc, 0xd, 0xe.
```

A `MagicItems` item takes its restriction from its ROW NAME instead. `R0995` writes
`0x07` into the descriptor byte and clears bit 1 when `item+0x44 == 5`; `item+0x44` is a prefix
test on the row name at `record+0x4` through `R0996` (`strstr`, `-1` for no match):
`Potion` = 3, `Book` = 5, `Scroll` = 4, otherwise 0. Shipped, 49 rows, both roots: 13 Potion,
5 Scroll, 5 Book, 26 neither. The five `Book` rows are the only mage-only `MagicItems`, and
they are refused a second time by the `teachSpell` gate (`ITEM-WEAR-058`).

A refusal returns `this` rather than the displaced item, and every caller treats a non-null
return as "put this in the container" — so a refused piece lands in the backpack. **No slot
restricts by type**: the one runtime-class test on the path, `Humanoid`'s `vt+0x38`
(`R0933`, testing against `Armor`), branches between two byte-for-byte identical calls.

One cross-slot rule: `Shield::Equip` reads param `0x0e` of the already-worn weapon's row and,
when it is `2`, takes the weapon off through `actor->vt+0x40` and appends it to the container
itself. `Weapon::Equip` carries the mirror arm for an already-worn shield.

### Equipment named by a `Humans` row

The ten trailing strings of a `Humans` row are ten equipment cells, and the cell **position**
picks the C++ class with the string never consulted:

- Cell 0 builds a `Weapon`: in `R0656` at `L04699` it allocates `0x84` bytes and calls constructor `R0665`.
- Cell 1 builds a `Shield`: at `L04700`, `0x68` bytes, constructor `R0973`.
- Cells 2 to 9 each build an `Armor`: `0x68` bytes, constructor `R0974`.
- Each cell then calls `actor->vt+0x3c(item)`, which on a Human is `R0975`: it calls `actor->vt+0x38(item)` and appends to `actor+0x7c` only the return value.

This is not how the `Units` arm does it — there the class is chosen by searching the string for
the literal `"Shield"` (`UNIT-EQUIP-005`). A name resolves through
`R0991` (cut from `{`) → `R0956` (a Shapes word) → `R0957` (a Materials
word) → `R1858` → `R0666` (exact lookup, 0 = not found). `R1858` puts back
the shape word the **material** implies — `"Soft "` for a material containing `Leather`,
`"Wooden "` for one containing `Wood` — and the `Weapon` constructor is the one that does not
call it. The `Shield` constructor alone then cuts the literal `" Shield"` off the residue.

```
actor->vt+0x38  Equip(item)    -> item->vt+0x38(actor); returns the displaced item
actor->vt+0x40  Unequip(item)  -> item->vt+0x3c(actor); returns the item
Item::vt+0x38                  -> refuses if actor.health <= 0; else attaches every effect of
                                  the item to the actor and DELETES the item (a consumable)
Item::vt+0x3c                  -> prints "Unknown item takeoff"; only the three subclasses
                                  implement a real take-off
```

Only the Unit wrappers add a trailing `actor->vt+0x54()`: equip `R0931` and removal
`R0932` make that call, and its Unit target only returns `actor+0x1c`. The Humanoid/Human
wrappers `R0933`/`R0934` omit it. `vt+0x54` is Token value, not an equipment recompute
(`ITEM-EQUIP-006`, whose universal wrapper and recompute clauses are retracted;
`SAV-EQUIPCALL-554`).
