# Staged area effects

[Reference](format.md)

## Staged area effects

### Cell aliases and direct-damage dispatch

The ring and single-pass blast read cell occupants in `+4`, `+8`, `+0xc` order and immediately
apply to each pointer. No object set suppresses a pointer encountered in another cell. Ring
coordinates pass the interior gate; blast Building coordinates are truncated independently to
bytes without that gate. The persistent cloud pulse reads only `+4`, not the Building slot
(`UNIT-AREAVISIT-071`).

Inner-effect polymorphism matters: base Effect `L02981+3c` is the ordinary timed-effect attach,
but DirectDamage `L02980+3c` is `R0564` itself. The read fire-sacrifice, acid-stream and
meteor-storm producers carry DirectDamage. They therefore do not gain duplicate suppression from
the base-effect non-stacking rule. Fireball alone takes the explicit copy/divide branch of the
local dispatcher; it is not the only direct-damage area route (`UNIT-AREADIRECT-072`). This narrows
the old base-class-only and universal-idempotence clauses in `MAGIC-AREAAPPLY-038` (amended and partially retracted in the ledger) and
`MAGIC-FIREDIV-047` (amended and partially retracted in the ledger); the unit n² storage and spell-2 normalization remain valid for that unit path.

For Building `R0492`, the size getter is 1 regardless of rectangle size. Every reached direct
call uses the separate Building damage core, subtracts a word at `+42`, clamps a negative signed
result to zero, and notifies for positive damage. With aliases held present, repeated calls also
continue at HP zero. Actual HP-to-destruction timing is Unknown; do not infer either immediate
detachment or an immortal target (`UNIT-AREAHP-073`, `UNIT-STRUCTDETACH-074`).

Settled by `MAGIC-AREAPULSE-037` (amended and partially retracted in the ledger) and corrected/completed by `MAGIC-RING-048`.
`Distribution system == 5` builds an `AreaEffect` with mode 2, stage byte `+0x4b = 0`,
timer word `+0x4c = 0` and orientation byte

```
orientation = (directionByte & 0xff) >> 5       0..7
```

The shipped rows selecting this mode are ids 4 `fire_sacrifice`, 9 `acid_stream` and 21
`meteor_storm` on both roots.

### Stage clock and cell consumer

The first stage runs on the first tick. A stage reloads the timer with 2. Each following tick
decrements the timer and returns while its old value is positive, so later stages run three ticks
apart:

```
stage ticks = 0, 3, 6, ...
```

The prior two-tick reading treated the reload value as the interval. For every coordinate entry the
consumer computes target cell plus offset and truncates x and y to bytes. The position helper accepts
the cell exactly when `8 <= x <= mapWidth-9` and `8 <= y <= mapHeight-9`. It clamps its stored x and y
to that range even when it rejects the cell. A rejected cell is neither sent nor applied. An accepted
cell receives one client message for picture `2*spellId+9`, then one application attempt for occupant
slots `+0x4`, `+0x8` and `+0xc`, in that order. The picture and simulation therefore use the same
accepted cell list. The sack slot `+0x10` is not read.

After the full list, the consumer increments the byte stage counter. It sets completion byte
`+0x40 = 1` when the incremented stage is greater than or equal to the arm limit. The common effect
driver reaps it on that tick. Staged effects register no map layer and perform no separate cleanup.

### Fire Sacrifice

Fire Sacrifice ignores the orientation. Its two stages are fixed ordered lists relative to the
target cell:

```
stage 0: (-1, 1) (-1, 0) (-1,-1) ( 0, 1) ( 0,-1) ( 1, 1) ( 1, 0) ( 1,-1)
stage 1: (-2, 1) (-2, 0) (-2,-1) (-1, 2) ( 0, 2) ( 1, 2)
         (-1,-2) ( 0,-2) ( 1,-2) ( 2, 1) ( 2, 0) ( 2,-1)
```

The centre and the four corners of the radius-2 square are not visited.

### Acid Stream

Acid Stream has six stages. Define the two ordered source tables:

```
A_s = [(x,s) for x=-s..s]             s = 0..4
A_5 = []
B_s = [(s-i,i) for i=0..s]            s = 0..5
```

Even orientations use A and odd orientations use B:

| orientation | transformed offset |
|---:|---|
| 0 | `(+dx,-dy)` from A |
| 1 | `(+dx,-dy)` from B |
| 2 | `(+dy,+dx)` from A |
| 3 | `(+dx,+dy)` from B |
| 4 | `(-dx,+dy)` from A |
| 5 | `(-dx,+dy)` from B |
| 6 | `(-dy,+dx)` from A |
| 7 | `(-dx,-dy)` from B |

Stage 5 is empty for even orientations and contains six cells for odd orientations. The client
burst lifetime is 18 ticks rather than the default 16.

### Meteor Storm

Meteor Storm has 32 one-cell stages. Each stage calls `rand(5)` twice, in x-then-y order:

```
dx = first rand(5) - 2
dy = second rand(5) - 2
```

Each coordinate is in `[-2,3]`. The orientation and the `Radius` column are not used. Cells can
repeat across stages. A sampled cell is accepted only within the eight-cell interior:
`8 <= x <= mapWidth-9` and `8 <= y <= mapHeight-9`. A rejected edge sample is clamped internally
but sends, applies and draws nothing. An accepted stage applies to occupants and sends picture 51
at that cell with lifetime 16. The client consumes no further placement RNG.

Picture 51 is stationary at the accepted cell. Its sixteen paced client ticks draw frame 8 at
vertical offsets `-28,-24,-20,-16,-12,-8,-4,0`, then frames 0 through 7 at zero offset. It ends on
frame 7 and does not wrap (`MAGIC-093`, `ANIM-046`).

### Table layout and customisation limits

The two fixed arms use stage records in `ROM.EXE`:

```
record size 0xa4
+0x00  i32 count
+0x04  i32 dx[20]
+0x54  i32 dy[20]
```

Fire Sacrifice and Acid Stream take terminal counts 2 and 6 from image dwords. Meteor Storm takes
the immediate 32 and overwrites one count-1 record. The 20-cell capacity, stage bounds,
8-orientation transform, random domain, one-byte stage counter, one-byte map coordinates and a 16-bit timer are
engine limits. Changing the shipped fixed geometry changes `ROM.EXE`, not `Data.bin`. An engine may
externalize replacement tables without changing any shipped file bytes. More than 20 cells, 255
stages, 8 directions or single-byte coordinates requires a wider engine representation.
