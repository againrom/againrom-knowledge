<a id="item--the-item-its-container-and-the-sack--specification-partial"></a>

# Items, equipment and sacks

Items bind definition rows to instance fields and optional Effects. Containers
and equipment slots retain object identity; movement commands transfer those
objects. Sacks attach containers to map cells. — ITEM-CLASS-001,
ITEM-CARRY-015, ITEM-EFFGRAM-070, ITEM-EFFSAVE-077

The double-factor ladder and damage column in ITEM-SCALE-017 and
ITEM-DMGCOL-018 are retracted; use ITEM-LADDER-019 through ITEM-WEAPCOL-021.
The meanings of unnamed Data.bin columns, item+0x47, producers of item+0x44
outside the specified constructors, and the transition to actor tick state 2
remain Unknown. Stock generation and pricing are in [SHOP](../shop/format.md).

## Object and transfer model

| Object | Runtime size | Definition source |
|---|---:|---|
| Item | 0x50 | Magic Items |
| Armor / Shield | 0x68 each | Armors / Shields |
| Weapon | 0x84 | Weapons |
| Sack / container | 0x44 / 0x24 | Container ownership and map-cell attachment |

Construction binds a row, tier and material, computes item fields and parses
Effects. Equip/unequip applies the per-class and per-Effect rules. Transfers
move whole objects or split quantities while preserving the specified owner,
container and load updates; death and authored loot feed sack creation.
Archive field order is linked through each object's persistence contract.
— ITEM-CLASS-001, ITEM-CARRY-015, ITEM-EFFGRAM-070, ITEM-EFFSAVE-077

## Reference map

| Reference | Contents |
|---|---|
| <a id="at-a-glance"></a><a id="the-item-record"></a><a id="structure"></a><a id="item-fields"></a><a id="magic-item-value"></a> [Item fields and definitions](fields.md) | Class sizes, row binding, instance fields and price |
| <a id="equipment-cell-effects-item-effgram-070item-effsave-077"></a><a id="a-weapons-numbers--R0856"></a><a id="an-armours-and-a-shields-numbers--R0890-R0888-item-armfill-032"></a><a id="what-equipping-an-armour-or-a-shield-does-item-armfold-033"></a><a id="the-weapons-rows-columns"></a><a id="a-weapons-own-damage-line"></a> [Equipment grammar and item formulas](effects.md) | Effect syntax, scalar targets and item formulas |
| <a id="the-container"></a><a id="containers"></a><a id="whole-object-transfer-and-serialized-state"></a><a id="equipment"></a><a id="equipment-transfer-boundary"></a><a id="weapon-borne-spell-state-item-caststate-056"></a><a id="the-wear-rule-item-wear-055-item-wear-057"></a><a id="equipment-named-by-a-humans-row"></a> [Containers and equipment state](equipment.md) | Containers, slots, suitability and weapon-spell state |
| <a id="what-an-item-is-called--the-display-name-item-dispname-036item-namelimit-042"></a><a id="what-an-item-looks-like--item0x40-item-appear-023-item-appear-024"></a><a id="display-names-item-dispname-036item-namelimit-042"></a><a id="appearance-field-item0x40-item-appear-023-item-appear-024"></a><a id="the-registry--which-items-have-a-picture-item-pict-046item-pict-052"></a> [Names and pictures](appearance.md) | Name construction and resource lookup |
| <a id="moving-an-item--command-0x22"></a><a id="what-crosses-a-mission-boundary"></a><a id="move-command-0x22"></a><a id="carried-potion-and-scroll-activation"></a><a id="persistence"></a><a id="mission-boundary-transfers"></a><a id="document-access-item"></a> [Transfers, activation and persistence](transfer.md) | Move/use commands and save/mission transitions |
| <a id="the-sack"></a><a id="where-a-sack-comes-from"></a><a id="sacks"></a><a id="death"></a><a id="sack-producers"></a><a id="authored-staffs-dragons-and-death"></a> [Sacks, death and loot](sacks.md) | Map loot, corpse drops and sack lifecycle |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.
