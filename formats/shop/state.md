# Classes and state requirements

[Reference](format.md)

## The runtime class table

`.data` carries 28 MFC `CRuntimeClass` records (`{name, objectSize, schema, CreateObject, base,
next}`, 24 bytes). The shop-relevant hierarchy, with the image's own names:

```
CObject
├── Player 0x70                     ; +0x38 = money
├── CMultiShopShelf 0x1c   CMultiShopInstance 0xa0   CMultiShopTemplate 0x98
└── Token 0x3c
    ├── Building 0x6c  ├── Outpost 0xac  ├── Tavern 0xa0  └── Shop 0x74
    ├── Item 0x50      ├── Armor 0x68    ├── Shield 0x68  └── Weapon 0x84
    └── Sack 0x44
```

The full 28-row runtime-class population, including the unit/spell/effect side, is bounded by
`SHOP-CLS-001`.


## State requirements

- Refresh clears and refills stock when nobody has the shop open: on each
  city-screen entry in single player, or every 180 ticks in multiplayer.
- Stock is not saved. The generator uses a clock seed; a stable mission seed
  does not reproduce the original selection state.
- A one-participant campaign session routes shop access to the town shop.
- Sell payment rounds the whole stack; displayed totals halve per unit and
  can exceed payment by `floor(qty/2)` for odd item prices.
- Both sides share five trade places. Failed affordability stops the commit.
- Randomly drawn stock is not deduplicated. Literal potions and returned
  items use merging insertion; equal-code Potions ignore effect differences.

These are the same lifecycle, rounding and duplicate rules specified in the
sections above. Shop `+0x0c` is a varying price parameter, not durability.


## Document access item exclusion

Shop generation cannot create the document access item at campaign start. No town or stock
generation occurs before mission 10, and mission 10 routes directly to mission 20. Later live stock
generation uses weapons, armor, shields, spells, and six literal potions. The MagicItems name-filter
routine has no live caller. `Quest Documents` also has price -1, below the live pool floor 0
(`SHOP-DOC-029`).
