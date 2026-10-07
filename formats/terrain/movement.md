# Passability and cell records

[Reference](format.md)

## Tile word → movement (`TERR-PASS-049`, `TERR-PASS-050` (partially retracted), `TERR-PASS-051` (amended, superseded), `TERR-COST-052`, `TERR-PASS-053`)

The same word drives a second, wholly separate reading: whether a **ground unit** may stand on the
cell. It is the sim side, not the render side, and it uses a different bit split — the whole
decision runs on `w & 0x3ff` plus bit 13, so **bits 10–12 and 14–15 never reach it** (0 of 880 704
shipped cells set any of them). Sim-side grid layout, the ingest and the persistence are specified
in [ALM](../alm/format.md) → "Runtime passability"; what belongs here is the tile word's own part.

`rom.exe R0469` maps the masked word to a terrain class and a movement cost:

The water subcell rejection below is retained. TERR-PASS-050 is amended
because its water-return prose omitted this guard. Ingest applies its raw
water block test independently of the returned class. — TERR-WATERBOUND-164

```
i = w & 0x3ff

if (i & 0x300) == 0x200:                     # the water range, taken before anything else
    if (i & 0xf) >= 8:      return 0xff, 8   # reject      (0 shipped cells)
    if (i & 0xf) == 4 and (i & 0x30) == 0x10:
                            return 1,    8   # Land        (0 shipped cells)
                            return 9,    8   # Water; the cost is the literal 8, never CostWater

s = i & 0xf ; b = (i >> 4) & 3 ; g = (i >> 6) & 0xf     # same split as the render mapping
if s >= 14:                 return 0xff, 8   # reject      (0 shipped cells)

pri, sec = pair[g]                           # world+0x54156 + 2g   (immediates, below)
sel      = level[b][s]                       # world+0x540d6 + 16b + s, in 1..5
cls      = sec if sel <= 2 else pri
cost     = [ cost[sec],
             (3*cost[sec] + cost[pri]) >> 2,
             (  cost[sec] +   cost[pri]) >> 1,
             (3*cost[pri] + cost[sec]) >> 2,
             cost[pri] ][sel - 1]            # cost[k] = world+0x54176+k, [0] = 0xff
return cls, cost
```

`pair[g]` is the same primary/secondary pairing the render mapping's terrain names come from,
which is what makes this an independent attestation of them from the movement side:

```
g  0 (Grass, Land)   g  4 (Stones, Land)     g  8..11  never written (water early-out)
g  1 (Cracked, Land) g  5 (Cracked, Stones)  g 12 (Road, Land)
g  2 (Sand, Land)    g  6 (Flowers, Savanna) g 13..15  never written — AND REACHABLE:
g  3 (Savanna, Land) g  7 (Mountain, Stones)           a map with g >= 13 reads uninitialised
                                                       heap. 0 shipped cells do it.

level[b][s], s = 0..13            cost[1..10] from world.res:data/map.reg  (Cost only)
b=0 [2,3,2,4,3,4,2,2,2,2,4,4,4,4]   Land 8 Grass 8 Flowers 8 Sand 14 Cracked 6
b=1 [3,5,3,3,1,3,2,4,2,2,4,2,4,4]   Stones 12 Savanna 8 Mountain 16 Water 8 Road 6
b=2 [2,3,2,4,3,4,2,4,2,2,4,2,4,4]
b=3 [5,5,5,5,5,5,2,2,2,2,4,4,4,4]
```

A ground mover is blocked by tile bit 13 (`w&0x2000`), Mountain class 8,
raw water bits `(w&0x300)==0x200`, a nonzero type 3 object cell, or the
eight-cell border. Bit13 also controls the render-path dirt composite;
that graphical use does not replace the other movement blockers.

### Cost, block and later state

With the constructor's default cost vector and an empty interior cell,
word `0x0214` produces classifier class 1 and cost 8 but block 1. Word
`0x03ff` produces cost 255 but block 0. Cost is not the block predicate.
The ingest copies its final static block plane to the dynamic plane;
objects replace the block with 5 and the border replaces it with 31.
The actual footprint tests apply `block & mask`, independently of cost.
These are immediate-ingest rules. Later structures, cell records and
occupants can change both planes. — TERR-TILECONTROL-163

Mission 100 authors two turtle placements on border cells; by static read the loader's seat call
returns 0 for such a cell and the listing then logs and deletes the actor (`TERR-PLACE-208`).

Bits 10–12 and 14–15 do not affect the selected ingest rules, even though
the render grid retains them. Off-corpus non-water groups 13–15 can read
constructor-unwritten terrain pairs when the subcell is below 14. Their
native allocation contents remain Unknown; synthetic residues produce
different classes. — TERR-TILECONTROL-163, TERR-WATERBOUND-164

### Who is asking — the mover (`TERR-MOVE-054`…`TERR-MOVE-057` (amended, superseded), of which `TERR-MOVE-055` is amended)

The `n×n` footprint and the mask are **not** properties of a `units.reg` class. Both come from the
simulation actor instance the predicate is called on (base vtable `L00001`, derived `L00002`,
`L00003` — the only family that allocates the `0xb4`-byte mover and stores it at `+0x154`):

- Vtable offset `+0x1c` is `R0256`, which returns the byte at `+0x49`: the footprint side n.
- Vtable offset `+0x20` is `R1344`, which returns the byte at `+0x4a`: the mask selector, 1, 2 or 3.
- The base constructor `R0185` sets the byte at `+0x49` to 1 (`L00821`) and the byte at `+0x4a` to 1 (`L00820`) as constructor defaults.

The constructor's1 values are defaults. Data.bin Units/Humans streamers
overwrite tokenSize/movementType when the parameter is not−1. Installed
Ghost/Bee use domain 2 (mask0x44), Bat_Sonic/Dragon use 3 (mask0x82), and
footprints include sizes 2 and 3. The universal 1×1/0x41 consequence is
retracted by `TERR-MOVE-057`. The units.reg drawable class array is not a
simulation input: TileSize belongs to drawable CUnit, while Z selects the
drawable class. The simulation-domain clause of `REG-UNITS-061` is retained
as amended. — TERR-MOVE-054, TERR-MOVE-055

### Movement speed (`TERR-MOVE-056`)

`R1088(world, actor, u16 srcCell, u8 facing)`; the direction tables are written as
immediates by the world constructor.

```
dir = ((facing + 0x10) >> 5) & 0xff                                   0..7
dst = srcCell + (i16)world[0x58ec0 + 4*dir]
v0  = actor->[0x70]->[0x3c]->[0x44]  if nonzero  else (i16)actor->[0x8c]   (16 at construction)
if vt+0x20() != 1:  v = v0
else:
  d = clamp((i8)(height[src] - height[dst]), -32, +32)      height plane world+0x9451c
  v = SpeedMultiplier * v0                                  world+0x58db4, map.reg = 8
  v = v + ((v * d) >> 6)                                    SAR: uphill slows, downhill speeds up
  c = (u8)(cost[src] + cost[dst]) >> 1 ; if c == 0 then 8   byte-width add
  v = v / c                                                 signed idiv
v = clamp(v, 1, 63)
mover[0xa8] = v ; mover[0xae] = dir
dx = (i8)world[0x58eb0+dir] ; dy = (i8)world[0x58eb8+dir]
mover[0xb0] = dx*dy != 0 ? (i8)ftol(v*dx*0.707) : (i8)(v*dx)      double 0.707 @L06908
mover[0xb1] = dx*dy != 0 ? (i8)ftol(v*dy*0.707) : (i8)(v*dy)
s = |mover[0xb0]| or |mover[0xb1]| if that is 0, or 1 ; mover[0xaa] = ceil(256 / s)
```

`cost(cell)` = `R1087` is **not a pure read**: with `block[cell] & 0x20` and a nonzero byte
at `record+0xe` in the `world+0x540b8` table it shifts the stored cost right by 2 and writes it
back. The read happens once per transit start for each of the two cells, before any recompute of
them, and every actor tick precedes every area-effect tick (`MOVE-084`); two reads with no recompute
between them divide twice, for every layer slot including Wall of Earth (`MOVE-085`). No cost plane
is saved, the recompute restores the byte from `+0x00` or from `CostCracked` (the sweep of other
stores is partial), and two inline readers divide by the byte with no zero guard, which decay can
produce (`MOVE-086`). With installed maps, v0=16 and SpeedMultiplier=8, the resulting
speed spans 4..32 (mode16); these are data-derived values within the clamp.

## Cell records

[Cell records and Building attachment](cells.md) specify the runtime payload,
registration, recompute, masks, detach and saved state.
