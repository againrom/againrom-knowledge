# Cursor, item and spell resources

[Reference](format.md)

## Cursor sheets (`SPR16A-CURSOR-046`)

The executable registers 23 `.16a` cursor sheets. Paths, hotspots, opaque period arguments and
file geometry are tabulated by `SPR16A-CURSOR-046`. Command-relevant entries:

| cursor | frames | size | hotspot | period argument |
|---|---:|---|---|---:|
| default | 1 | 32x32 | 5,5 | 2000000000 |
| move | 5 | 32x32 | 15,15 | 100 |
| swarm | 5 | 44x44 | 21,21 | 100 |
| attack | 10 | 32x32 | 3,3 | 100 |
| defend | 8 | 32x32 | 15,13 | 100 |
| select | 1 | 32x32 | 3,4 | 100 |
| patrol | 8 | 32x32 | 8,25 | 100 |
| cast | 14 | 32x32 | 15,15 | 100 |
| pickup | 15 | 32x32 | 12,13 | 66 |

All eight edge arrows are one 32x32 frame; small-default is one 16x16 frame; cantput is one
64x64 frame; town and backpack are one 32x32 frame; dice and wait are 15 and 10 32x32 frames.
All payloads are identical across EN/RU. The period argument's downstream time unit is not part of
the format contract. Frame count, dimensions and pixels are data; hotspot and period are executable
constants.

## The item-icon subtree (`SPR16A-ICON-027`)

Installed `inventory/` .16a icons each have one 80×80 frame with bit 31 set.
Literal levels are 1..15, corresponding to alpha 2/16..16/16 under
SPR16A-ALPHA-025's normal-memory arm; its low-memory generalization,
format-level 15-step restriction and native-unreachability assertion are
partially retracted. SPR16A-080 retains the local addressing contract;
native reachability remains Unknown.
These dimensions and levels describe the installed icon
set, not all sprites: other .16a resources have mixed geometry, and adjacent
.256 equipment sheets include zero-frame resources. — SPR16A-ICON-027,
SPR16A-BOUND-016

## Projectile sheets

The 31 installed projectiles.reg rows select 24 .16a sheets and 7 .256 sheets
through the A16 key. Frame geometry is uniform within each of those sheets.
The extension is selected by the registry. — SPR16A-PROJ-024

The consumer-facing law is the draw's: a sheet must hold `Phases * 9` frames when the registry's
`Flip` is set and `Phases * RotationPhases` when it is not. The installed exceptions are `healing` (7 `Phases`, 8 frames; its mark renderer reaches all eight) and `goblin\arrow`
(one frame plus an 8-byte residue before the count trailer, the `SPR256-EXC-017` shape).

`Width`/`Height` are centering fields and need not equal frame dimensions.
Absent values use the loader's `0x40` default; installed frames include
12×12 and 64×96 dimensions.

**How one of these sheets is drawn** (SPR16A-ALPHA-025, partially retracted for
its low-memory generalization and native-unreachability assertion;
SPR16A-080 retains the local addressing contract; REG-PROJ-087).
`A16` picks the extension *and* the C++ class —
the `.256` and `.16a` sprite vtables differ in exactly two slots, the destructor and `+0x18`, so
the class is what chooses the blit. The registry's `Palette` is a **boolean**, "this sheet carries
its own colour table": non-zero makes the loader build the sprite a 16-level lookup (mode 4, no
tint, for `.16a`), zero makes it build nothing and the draw hands the blit the shared table from
`projectiles.pal` instead. Installed `Palette` values agree with trailer bit 31. The three absent-key
`.256` sheets have that bit clear and begin frames at byte 0 rather than 1024
(`REG-PROJ-087`). Installed projectile literals use levels 1..15; a sheet
need not contain an opaque level-15 pixel (`SPR16A-PROJ-026`).

## Spell art: which sheets one cast can reach

Spell ID maps to projectile record `2*id+8` for the cast picture and
`2*id+9` for the burst picture. The installed bindings are below.
— SPR16A-CAST-028

```
cast parity  (+8), 15 defined rows
  firebolt(10) fireball(12) p_fire(18) healing(20) poison_d(24) p_water(28) Drain(30)
  lightnin(34) chain(36) p_air(40) shield(44) p_earth(52) bless(54) teleport(60) Curse(62)

burst parity (+9), 8 defined rows
  fireexpl(13) firewall(15) smallxpl(17) freeze(23) poison(25) acid(27) wall(47) Meteor(51)

both parities defined: fire_ball and poison_cloud, and no other spell
neither parity defined: light, invisibility, darkness, stone_curse, haste,
                        control_spirit, slow
```

**Orientation.** Exactly two of the 23 spell-reachable rows carry `RotationPhases = 16` with
`Flip = 1`: `firebolt` and `fireball`, both `Phases = 4`, both walking 36 frames, which is
`Phases * 9` under the mirrored-facing law. Every other spell-reachable row is `RotationPhases = 1`,
so its frame index is the phase alone and it stores no per-direction row.

**Frame law.** Use `Phases*9` with Flip, otherwise
`Phases*RotationPhases`. `healing` has eight frames against `Phases = 7`;
its mark renderer reaches frames 0–7 regardless of that declared phase count. — SPR16A-PROJ-024, SPR16A-031

**Reachability.** Of the 31 shipped rows, `steam`(8) is named by no spell id at either parity and by
no `units.reg` `Projectile` key.

## Which consumer draws each spell-reachable sheet (`SPR16A-PART-030`)

`SPR16A-MARK-029` is partially retracted: spell reachability does not
imply the actor-mark renderer. The rows divide among actor marks, moving
cast objects, bursts and ground effects. — SPR16A-PART-030

**Cast parity, `2*id + 8`, 15 rows.** Ten are actor-bound marks drawn from the mark array on the
actor: `p_fire`(18) `healing`(20) `poison_d`(24) `p_water`(28) `Drain`(30) `p_air`(40) `shield`(44)
`p_earth`(52) `bless`(54) `Curse`(62). Five are moving cast objects: `firebolt`(10) `fireball`(12)
`lightnin`(34) `chain`(36) `teleport`(60). `healing` and `Drain` are in both groups, because the
cast-spawn switch also gives them a one-tick flight.

**Burst parity, `2*id + 9`, 8 rows.** All eight are burst art. Four of them are additionally painted
on ground cells under an area effect: `firewall`(15) `freeze`(23) `poison`(25) `wall`(47). The other
four are burst-only: `fireexpl`(13) `smallxpl`(17) `acid`(27) `Meteor`(51).

Every row in the mark group and the ground group has `RotationPhases = 1`. The two rotating sheets
are both in the cast-object group.

### Frame demand of the mark path

Corrected for Heal and Drain Life by `SPR16A-031`.

The frame index a mark carries comes from an engine immediate, not from the sheet. Each arm can
produce a bounded set:

| record index | sheet | frames the arm can index | declared `Phases` |
|---|---|---|---|
| 0x12 0x1c 0x28 0x34 | `p_fire` `p_water` `p_air` `p_earth` | 6 | 6 |
| 0x14 | `healing` | 8 | 7 |
| 0x18 | `poison_d` | 6 | 8 |
| 0x1e | `Drain` | 8 | 9 |
| 0x2c | `shield` | 5 | 5 |
| 0x36 | `bless` | 5 | 5 |
| 0x3e | `Curse` | 5 | 5 |

Seven of the ten agree exactly. `poison_d` carries two frames no actor mark can reach and `Drain`
carries one, frame 8. `healing` is the inverse registry exception: the sheet walks eight frames,
`Phases` is 7, and its mark reaches all eight, 0 through 7. The terminal phase is produced because
the builder tests the old phase below 7, then increments and copies it (`SPR16A-031`).

Nine of the ten rows are 12x12 in the registry and `healing` is 16x16; the mark draw subtracts
exactly `Width/2` and `Height/2`, so those are centring halves here as everywhere else.
