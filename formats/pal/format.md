<a id="pal-pal-palette--specification"></a>

# PAL palette resources

The palette loader consumes either a BMP color table at byte `0x36` or
sixteen consecutive raw color tables. A table contains 256 four-byte BGR0
entries. The consuming path selects the resource shape and shade-table mode.
— PAL-FILE-001, PAL-MODE4-010 (partially retracted only for its native-unreachability
assertion; the resource and mode contract stands).

<a id="two-shapes-told-apart-by-how-the-engine-opens-the-file"></a><a id="corpus"></a>

## Resource layouts

```
A — the per-class palette                     B — the shared owner palettes
   graphics\units\<class dir>\palette.pal        graphics\units\humans\human.pal
   graphics\units\<class dir>\palette2.pal
   graphics\units\<class dir>\palette3.pal
   graphics\units\<class dir>\palette4.pal

   open(path); seek(0x36); read(0x400)          open(path); read(0x4000)
                                                   ^ no seek
   -> a BMP's 256-entry colour table            -> 16 consecutive 1024-byte colour tables
   shipped size 17462 B (54 + 1024 + 128*128)   shipped size 16384 B exactly
             or 26678 B (dragon, 160*160)       first two bytes 00 00 — it is NOT a BMP
```

A colour table entry is 4 bytes, `[B, G, R, 0]` — the same layout as a `.256`'s own leading
1024 bytes (`SPR256-PAL-003`, `SPR256-PAL-011`). The engine reads bytes 0..2 of each entry and
ignores byte 3.

<a id="how-a-palette-becomes-the-draws-lookup"></a>

## Shade-table construction

Identical for every sprite in the game — the tier path is not special:

```
ShadeTable::Build(palette, nLevels = 0x10, mode = 2, useTint = 1)

  obj+0x04 = nLevels
  obj+0x08 = malloc(nLevels << 9)          nLevels rows of 256 u16

  for level = 0 .. nLevels-1:
      m = nLevels - level
      for i = 0 .. 255:
          R = clamp((pal[i].R + tintR) * m * 2 / nLevels, 0, 255)     truncating divide
          G = clamp((pal[i].G + tintG) * m * 2 / nLevels, 0, 255)
          B = clamp((pal[i].B + tintB) * m * 2 / nLevels, 0, 255)
          row[i] = (R>>3)<<11 | (G>>2)<<5 | (B>>3)                    RGB565

  tint = (L10231, L10613, L10614), forced to (0,0,0) when useTint == 0

blit:  dst = table[level * 256 + srcIndex]
```

Mode is `arg3`; the six arms are a jump table at `L05822`: 0 white silhouette, 1 per-channel
tint with no gain, **2 the sprite mode above**, 3 the terrain ladder, **4 the `.16a` mode**,
5 the same gain applied to `(R+G+B)/3` (greyscale — what the unit **shadow** uses).

## Mode 4 — the `.16a` table, and the one arm the tint never reaches

```
for level = 1 .. 16:                                 // table row = level - 1
    for i = 0 .. 255:
        chan = clamp((pal[i].chan * level) >> 4, 0, 255)     // or / 18, see below
        row[i] = pack(R, G, B) into the framebuffer's widths and shifts
```

Mode-4 rules (PAL-MODE4-010, partially retracted only for its native-unreachability
assertion; the predicate and local operations retain their instruction scope):

- **No tint term appears on this arm**, on either branch. `useTint` is passed and ignored, so a
  `.16a` sprite does not change colour with the daylight cycle the way a `.256` unit does.
- The `>> 4` / `/ 18` pair is **not a display mode**. `R0788` sets the flag that picks it
  from `GlobalMemoryStatus`: `dwTotalPhys < 24000000`. The same flag halves the companion
  destination table from 65 536 entries a row to 8 192. The condition is an environment check, not a palette-field value.
- Normal-memory source generation uses `/16` and the decoder selects destination
  row `1+L`; their weights are complementary. The low-memory source branch uses
  `/18`, while the decoder selects destination row `L` with `old >> 3` and no
  added row. The normal complementarity rule does not apply to that branch.
  — SPR16A-080; [sprite encoding](../spr16a/encoding.md).
- The memory predicate does not establish whether a native original can run
  in that system state. Native low-memory reachability remains Unknown; the
  native-unreachability clause of PAL-MODE4-010 is partially retracted.

## The two shared projectile tables

`graphics\projectiles\projectiles.pal` and `projectile_.pal` are shape-A BMPs (5 174 B =
54 + 1 024 + 64×64) whose colour tables the projectile registry loader lifts. **Neither is a
sprite palette**: each becomes a 16-level **mode 2** shade table with `useTint = 0`, held in
the global at `L01248` and the global at `L07214` (`PAL-PROJ-011`). Only the first is ever read by a blit — it is
the external table the draw hands `vt+0x14` for the three projectile sheets that ship without a
palette of their own. `projectile_.pal`'s table has exactly one reader in the image, its own
teardown; the two colour tables differ on 255 of 256 entries, the dead one being largely grey.
Neither name occurs anywhere in the install's other executable.

## Which palette a unit is drawn through

One `units.reg` key decides, and it decides between two different mechanisms:

```
Palette = class+0x98            (units.reg int key "Palette"; no-parent default 0)

  Palette == 0    ->  the 16 SHARED tables, subscript = the drawable's OWNER
                      table[k] built once at units.reg load from human.pal + k*0x400
                      a 17th, table[16], is (0x10, mode 5, useTint 0) from sub-palette 0

                      k = CPlayer+0x08 of the record at drawable+0x14. Within the
                      vptr-stamped constructor family it has two stores (L07210,
                      L07211). The real vtable surface is six dwords — five function
                      entries plus null; a 16-slot dump crosses into the next class
                      (PAL-SHADE-012, ALM-CPLAYER-090). The value arrives on session
                      message 0x96 as the byte at msg+0x0c (PAL-SHADE-013):

                        Player+0x44 != 0 ->  k = Player+0x44 - 1
                                             = the ALM type-5 record's +0x00 word
                                             RECONNECT ONLY: a join allocates a fresh
                                             Player (L07242) whose +0x44 is 0 from the
                                             constructor (L07239), and +0x44's only
                                             non-zero writer is L02129, in the map load
                        otherwise        ->  k = Player+0x04 & 0x0f
                                             = the session slot, 1-based, first free
                                             = the type-5 record's 1-based ORDINAL, and
                                             the rule every faction of a shipped
                                             single-player mission takes (PAL-RULE-021)

                      The interface entry is created by the FIRST 0x96 for an index and
                      reused without a re-slot afterwards (PAL-FIRST-022, L07253 /
                      L07254), so a later 0x96 cannot change an existing colour.

                      k is read unmasked at L07218; the safe authored range is 0..15.
                      Shipped maps use 0..10 and one 13; 11, 12, 14 and 15 never occur.

  Palette == 1    ->  class+0x9c[0], built lazily from class+0xac[0]; no subscript

  Palette  > 1    ->  class+0x9c[face - 1], built lazily from class+0xac[face - 1]
                      face = the actor's Data.bin tier column, reaching the drawable
                      as unit+0x24
```

Shipped values over the 34 `units.reg` classes: **0 on 18** (every human and hero class),
**1 on 3** (`Catapult 1`, `Catapult 2`, `Death Star`), **4 on 13** (every monster class that
ships tiers). `class+0x9c` and `class+0xac` are `0x10` apart, so **each array holds four
entries** and `Palette <= 4` is a structure-width limit, not a file one.

## Where `face` comes from

```
Data.bin  Units column 30 / Humans column 17   "face"
   -> actor+0x4b  (u8; low 6 bits are the face, bits 6..7 are flags on the human arm)
   -> state message, the SECOND of two bytes under sync mask bit 0x4000
   -> [L04218]
   -> unit+0x24
   -> class+0x9c[unit+0x24 - 1]
```

The same byte subscripts the `(typeID, face)` prototype cache at `L07204` as
`typeID * 4 + face`, in two routines that must agree — the draw's `R0590` and the cache
builder `R0690`, the latter reading the value straight out of the `Data.bin` row.

Shipped `face`: Units 1..4 (four rows per creature family, one `typeID` each), Humans 1..31.
Over the 262 `Data.bin` rows whose `typeID` resolves to a `units.reg` class, **0** carry a `face`
outside `[1, Palette]` on a class that owns palettes.

<a id="reproducing-the-recolour--what-a-consumer-must-and-must-not-do"></a>

## Recolor procedure

- **Read the file.** No closed form reproduces a tier palette from tier 1. Fitted at their own
  optima over the 39 shipped (class, tier >= 2) pairs: per-channel affine misses by >= 18 of 255
  on the best class, a general 3x3 matrix + offset (which contains every hue rotation, saturation
  scale and channel swap) by >= 4, an HSV rotation has an inter-quartile hue spread up to 357.9
  degrees, and an index remap is impossible because at most 191 of 256 tier colours occur anywhere
  in tier 1.
- **Tier 1 is not a recolour.** `palette.pal` is byte-identical to the sheet's own embedded
  palette on all 256 entries of each installed class.
- **Never touch palette entry 0.** It is identical across every installed tier of every class and
  no standard `.256` emits it: 0 of 9 541 025 literal pixels over the 1370 standard sprites
  (`SPR256-PAL-012`, its transparent-key clause superseded and its population amended from 1376
  to 1370). It is *reserved*, not a key: no blitter tests it (`SPR256-KEY-044`), so a decoder
  must not treat it as transparent.
- **The owner band is separate and narrow.** The 16 sub-palettes of `human.pal` differ from
  sub-palette 0 on exactly the same 55 indices for all fifteen k: `4, 10, 13, 15-17, 21, 23-24,
  28, 36, 41-42, 44-45, 55, 81, 89, 104, 106, 109, 114, 116, 118-120, 123-125, 127, 141, 144-147,
  150-151, 162-164, 168-174, 176-179, 181, 193, 211, 225`.
- **The sixteen slots are seven hues and a grey at two brightness levels**, `k` paired with
  `k + 8` (`PAL-SLOT-015`). Saturation-weighted mean hue over the band, then mean saturation and
  value: `k=0` 235.8 blue, `k=1` 112.1 lime-green, `k=2` 354.5 red, `k=3` 22.0 red-orange, `k=4`
  287.7 violet-magenta, `k=5` 39.0 orange-yellow, `k=6` 174.5 cyan, `k=7` grey; `k=0..6` at
  saturation 0.510 and value 0.573, `k=8..14` the same hues at 0.561 and **0.423**, `k=15` grey
  at value 0.303.
- **The band is not the same fraction of every figure.** Complete decode of two sheets:
  `swordsman.256` 152 frames, 7 374 of 56 503 opaque pixels in the band (**13.05 %**), on the
  sleeves of both arms; `mage/sprites.256` 144 frames, 36 419 of 80 809 (**45.07 %**), on the
  hood and the mantle (`PAL-BAND-016`). A consumer that recolours a whole figure, or that
  recolours nothing, is visibly wrong in different amounts per class.
- **The doll and the dialogue speaker carry no owner shade.** `R0745`, the compositor
  both are drawn by, contains no reference to `L05816` or `L02113` and takes three
  surface parameters and no palette (`PAL-FIGURE-014`).
- **The world body draw applies the selection to every sprite it blits.** `R0552` holds
  the selected object in one stack slot and pushes it as the second of six arguments at all five
  of its blit call sites — `L05824` and `L05826` on `drawable+0x194`, `L07246` on `+0x198`,
  `L07247` and `L07248` on an object from `[L07265]`. The `drawable+0x18c & 1` test at
  `L07262` chooses the sprite source, not whether the shade is applied, so a map-placed
  humanoid is owner shaded (`PAL-BLIT-024`, `UNIT-SPRITE-042`).

**Discriminating check.** `units/monsters/goblin/palette*.pal`, entry 128, at the shipped daytime
level 3 (`nLevels = 16`, `m = 13`, tint 0):

| tier | `[B,G,R]` at entry 128 | `table[3][128]` |
|---|---|---|
| 1 | `6e 7f 9d` | `0xfe76` |
| 2 | `6e 94 9d` | `0xff96` |
| 3 | `6e 75 9d` | `0xfdf6` |
| 4 | `75 80 95` | `0xf697` |

Worked for tier 2: `R = 157*26/16 = 255` (clamped), `G = 148*26/16 = 240`, `B = 110*26/16 = 178`
→ `(255>>3)<<11 | (240>>2)<<5 | (178>>3)` = `0xf800 | 0x0780 | 0x0016` = `0xff96`. An
implementation that rounds instead of truncating, that applies the gain per pixel instead of per
table row, or that reads `[R,G,B,0]` instead of `[B,G,R,0]`, gives a different word here.

<a id="open"></a>

## Unknowns

- What the global sprite table `[L07265]` holds, and therefore whether the sprite a
  map-placed humanoid is drawn from is an equipment-composed figure or a class sheet. The shade
  reaches it either way (`PAL-BLIT-024`); what is unread is the object, and the index that
  selects it, which is computed before the `drawable+0x18c` branch
  (`PAL-RULE-021`…`PAL-BLIT-024`).
- Whether the shade object the shadow arm selects (`L10418` / `L10387`) is used at all: the
  silhouette blits `vt+0x1c` / `vt+0x3c` index `[L03341]` and never read the object handed to
  them. The contents of that global are a separate unresolved consumer relation.
- `palette_.pal`, which ships beside every `palette.pal` (byte-identical to it on 11 of 13
  classes, different on `ghost`) and matches no name the loader builds. The same `_` sibling
  pattern in the projectile directory is **not** a parallel: there the loader does build both
  names, and the `_` one is dead at the reader rather than at the name (`PAL-PROJ-011`).
- `units/heroes/human.pal` and `units/heroes_l/human.pal` — two further 16384-byte blobs no
  located call site opens; only `graphics\units\humans\human.pal` is read into `L10238`.
