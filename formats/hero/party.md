# Persistent companion joins

[Reference](format.md)

## Persistent companion joins

There are five shipped persistent-companion producer origins. Mission 30 town activation consumes
`AddHero=22` and constructs Fergard or Reniesta. Map instants transfer an existing `Hero` actor in
missions 40, 70, 100, and 140. A dynamic npc with `DataBinID=26` selects Humans server id

```
26 + sex/class selector + 4*floor(mission/40)
```

The selector bits are sex at bit 0 and class at bit 1. `MySex`, `MyClass`, and their inverted forms
derive them from the primary hero. The ten conditional source variants are recorded by
`HERO-JOIN-120` and compared across roots by `HERO-JOIN-128`.

Brian is mission-40 unit 6, `npc25`, fixed Humans row 42 `PC_Paladin`. Mission-70 Naira exists for a
male primary and is unit 151, `npc23`, Humans row 31 `PC_Naira_2`. Their implementation payload is:

| field | Brian | Naira |
|---|---|---|
| class / sex / actor type / definition type | fighter / male / `0x21` / 5 | fighter / female / `0x22` / 14 |
| face / figure / weapon body / geometry | 1 / `mfighter` / `swordsman2h` / 5 | 1 / `ffighter` / `archer` / 14 |
| Body / Reaction / Mind / Spirit | 41 / 39 / 25 / 21 | 39 / 41 / 22 / 26 |
| six skill levels | 0 / 25 / 3 / 0 / 0 / 1 | 0 / 24 / 5 / 0 / 0 / 38 |
| six skill XP values | 0 / 9834 / 331 / 0 / 0 / 100 | 0 / 8849 / 610 / 0 / 0 / 36404 |
| aggregate XP | 10265 | 45863 |
| max HP / MP | 157 / 0 | 177 / 0 |
| speed / sight raw / capacity / equipped load | 18 / 1679 / 411 / 457 | 20 / 1669 / 391 / 179 |
| to-hit / defence / absorption | 97 / 76 / 5 | 177 / 88 / 1 |
| physical base / spread | 20 / 15 | 17 / 11 |
| physical / fire / water / air / earth / astral protection | 0 / 10 / 10 / 10 / 10 / 10 | 0 / 32 / 13 / 29 / 13 / 13 |
| known spells / carried item instances | none / 0 | none / 0 |
| occupied worn slots | 1, 6, 7, 8, 9, 10, 12 | 1, 4, 5, 6, 7, 8, 9, 10, 12 |

Skill order is General, Blade/Fire, Axe/Water, Bludgeon/Air, Pike/Earth, Shooting/Astral. Sight is
the raw 1/256-cell word. Item codes, effects, prices, weights, damage and defence are in
`HERO-JOIN-125`.

The map ownership routine transfers the same actor object. It does not reconstruct the Human or
write `Player+0x34`. Current health and other mutable state must therefore be copied from the live
source actor rather than reset to the source-initial value. The transferred type remains in
`0x21..0x24`, so the actor keeps the ordinary Human inventory, worn slots, skill, spellbook and save
capabilities. The join does not make the companion primary.

Claims: `HERO-JOIN-120` through `HERO-JOIN-128`.
