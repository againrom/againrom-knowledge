# Transfers, activation and persistence

[Reference](format.md)

## Move command `0x22`

One command, five fields (`R0238`, partially retracted `SHOP-TRAY-025`): `cmd+0x0c` source code,
`cmd+0x0e` source index, `cmd+0x0d` destination code, `cmd+0x10` destination position,
`cmd+0x12` quantity. The command layout stands; the universal take/split clause is withdrawn.

| Code | Source | Destination |
|------|--------|-------------|
| 1 | an equipment slot (`slot = cmd+0x0e + 1`) | equip; the displaced item goes to `cmd+0x0e` of the container |
| 2 | the actor's container, index `cmd+0x0e`, quantity `cmd+0x12` | the actor's container at index `cmd+0x10` |
| 3 | **the ground** — validates a sack on the actor's own cell, sets `container+0x1c = cmd+0x10` and `actor+0x50 = 2`; refuses with `"Invalid pickup order - no sack there."` | **the ground** — see below |
| 4 | a shop (`4..8` selects tray or shelf) | a shop; stamps `item+0x14 = Player` |

Dropping to the ground:

```
destX = cmd+0x10 & 0xff ; destY = (cmd+0x10 >> 8) & 0xff
if |actorCol - destX| <= 2 && |actorRow - destY| <= 2 :  drop at (destX,destY)
else:                                                    drop at the actor's own cell
```

Never refused — out of range it lands underfoot. Money is opcode `0x23` and uses the same
window; it refuses an amount `<= 0` or one exceeding `Player+0x38`, debits `Player+0x38`, and
adds the amount to the sack's `+0x3c`. CastSpell item use is orders `0x25`/`0x26`;
ordinary Potion use is session `0x22` destination 1 (`ITEM-USE-113`).

### Carried Potion and Scroll activation

`ITEM-USE-112`, `ITEM-USE-113` and `ITEM-USE-114` distinguish input, reservation and commit. Mission backpack
double-click and dropping a carried item on the character panel reach the same
action. The shop backpack's double-click uses its borrowed character panel too.
The action requires one owned selected member and descriptor/class admission.
Ordinary Potion and Scroll descriptors admit both classes.

Without castSpell display bit 0x10, usable bit 1 emits a one-unit transfer to equip.
Base Item use returns the item unchanged at HP<=0. At HP>0 it applies every Effect
in order and destroys the detached item, even if healing was already capped or an
effect's class check made it a no-op. With both display bits it instead arms cast
mode and stores the carried slot. Arming does not remove the server item.

Map-up uses the item-specific target table, not simply the Spells target column.
It emits unit order `0x25` or point order `0x26`. The server removes one item and reserves
it at actor+68, building a Spell from its first kind-41 Effect. Point admission
requires Spells parameter 4 ==2 and otherwise restores at the original slot.
Missing/wrong first Effect has no local restore; no ordinary malformed-input
route is established.

Accepted cast-start sets actor+136=0; completion destroys the category-14 Item
after invoking apply, without an effect-success refund. Earlier cancellation can
restore it only under the act/+136/first-effect/id-match checks; cancellation
after accepted start and actor teardown can destroy it instead. Pathfinding and
the complete stale-target failure behavior remains Unknown. Shop arming is established, but casting
while the town/shop covers the target map remains Unknown.

`ITEM-VALUE-115`: the general Book value reader uses only its first Effect's
spell id, or zero if empty. It does not select the last Effect. Scroll has a
separate kind 41 summation. Shelf constructors override the stored price, and the
Item copy preserves that price and ordered Effect data without re-pricing.

## Persistence

`Effect::Serialize` writes the 37-byte Token head, u8 kind, u8 mode, u32 operand and u8 Token
state `+0x0c`: 44 bytes of body. It omits transient Effect `+0x44`. The effect-list serializer
writes u32 count followed by archive object references in list order; load clears the list and
appends that count in archive order (`ITEM-EFFSAVE-077`).

```
Item::Serialize      head, effects, u16 +0x40, u16 +0x42, u8 +0x44, u8 +0x45, u8 +0x46,
                     u16 +0x48, u16 +0x4a, u8 +0x47; on load, resolve +0x3c from +0x0c
Armor/Shield/Weapon  Item::Serialize, then their own tail, then the same resolve against
                     their own collection
container            the CObList elements, then +0x1c, then +0x20
Sack::Serialize      head, +0x3c, container
Unit::Serialize      ... actor+0x74, actor+0x78, container(actor+0x7c) ...
Humanoid::Serialize  Unit::Serialize, raw 24 bytes at +0x1cc, the twelve slots
                     actor+0x198+4i (i = 1..12), one reference at +0x1e4
```

Item copy/split preserves the deep Effect list but leaves the serialized `Item+0x47` byte
indeterminate. Saving such a clone writes that byte; its meaning remains unknown.

## Mission-boundary transfers

`PARTY-CARRY-005`'s carried blob is a single `WriteObject` on one actor. That invokes the
chain above, so **a hero's equipment and backpack cross inside it**; `PARTY-LOSS-006`'s strip
list touches none of `+0x74`, `+0x78`, `+0x7c` or `+0x198…+0x1e4`. A mercenary is never a root
of that graph and is not reachable from one, so nothing of it crosses but `MERC-DEATH-006`'s
fifteen pool integers — and if it dies first its gear is destroyed by the `"NPC"` test above.
A consumer must implement two persistences for one party.

## Document access item

`MagicItems[28]` is internal row `Quest Documents`, packed code `0x0e1c`. The display name is
`Valuable Documents` in EN and `Официальные Документы` in RU. Its `Data.bin` price is -1.

An eligible inventory action checks `(code & 0x0f00) == 0x0e00` and `(code & 0x1f) == 28`, then
posts campaign message `0x463`. The campaign arm opens the document panel. The panel binds to the
campaign record at `campaign+0x548` and reads its document collection. The item is an access trigger;
it does not contain the document entries.

Player-hero construction calls `R0993` after starting-skill and derived-stat application and
before runtime-id assignment. When `server+0x0c == 0` and `[L03937] != 2`, the helper constructs
an Item from the exact row name `Quest Documents` and appends it to the hero's `+0x7c` container.
The campaign construction path establishes the first gate. The second gate was not measured in a
fresh campaign, so actual item presence at mission 10 start is Unknown (`ITEM-DOC-054`,
`ITEM-DOC-069`). The campaign record already contains text documents 1, 2, and 3; a non-empty
collection does not establish item possession or panel access.

No mission-start routine read moves an item between party members on the campaign's
mission-to-mission edge: the carried actors keep their own containers, and an existing primary
hero is not given a second stack. The negative is bounded by 11 unread owners in the start closure,
the untraced SAV-load start path and the unfollowed indirect calls (`PARTY-037`). A member removed
by the mission-end culls takes its stack out of the party, and a second `Quest Documents` stack is
made only when a primary hero is constructed fresh (`PARTY-038`).

The class-14 code uses the low byte as its row; this UI gate uses only five
low bits. Rows `28+32n` therefore alias at the gate. The packed representation
allows at most 255 nonzero class-14 rows. The item producer, destination and
both producer gates are executable rules; a data-only control for changing
them is not established.
