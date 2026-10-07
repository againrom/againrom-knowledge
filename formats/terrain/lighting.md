# Terrain lighting and fog

[Reference](format.md)

## Terrain lighting (`TERR-LIGHT-011` (amended, superseded)…`TERR-LIGHT-013` (amended), `TERR-LIGHT-015`,
`TERR-LIGHT-018…TERR-LIGHT-023`, `TERR-LIGHT-028`, `TERR-LIGHT-029`,
`TERR-LIGHT-109`, `TERR-LIGHT-119…TERR-LIGHT-128`)

Terrain is **relief-shaded** (Lambertian slope shading), computed at runtime — the `.alm` stores
no lightmap:

- **Blit.** Each pixel is `LUT16[(level<<9 & 0xfffffe00) + srcIndex*2]`; `level` is bilinearly
  interpolated from four per-vertex brightness bytes (`R1506`/`R1507`). The LUT is a
  palette brightness ramp rebuilt on light change.
- **Per-vertex brightness** (`R0468`, grid at `P+0x18`, `P` = CMapView+0x80) — exact level:
  ```
  tanT   = tan|θ|                            H[] read as signed bytes
  stepH  = 32.0 / cos|θ|
  Δh₁    = H[x ∓ tanT, y+1] − H[x, y]        forward,  ONE row   (∓ = −tanT if θ≥0, +tanT if θ<0)
  Δh₂    = H[x, y] − H[x ∓ tanT, y−1]        backward, ONE row
  sᵢ     = atan2(Δhᵢ, stepH)
  axisᵢ  = clamp( L − R·sin(π/6 − sᵢ), 0, 95 )
  byte   = ftol( 0.5 · (axis₁ + axis₂) )     ← the shading level the blit LUT indexes
  ```
  `R = [L02293]` (day 32), `L = (R>>1) + [L05650] + 32` (day 62). Flat day vertex → 46.

  Both height differences span one cell along the row axis. The x±1 reads
apply the tanT shear within rows y±1; stepH lengthens the baseline along
the sun azimuth. Each arm clamps its level before the two levels are
averaged with 0.5. A central height difference or averaging before the clamps
changes the result. Axis 2 selects its lateral column by `y==1`, not by θ's
sign: the zero flag left by a decrement replaces the comparison flags. — TERR-LIGHT-028

  Written region = the **strict interior `{1..W-2}×{1..H-2}`** only (`R0468(1,1,0,0)`; the `±1`
  neighbour's sign follows the sun azimuth, so starting at 1 / ending at W-2 keeps both in bounds).
  The **outer ring of vertices is NOT computed** and has no one-sided fallback (`TERR-EDGE-025`, correcting an
  earlier "first row/col one-sided" reading); it is left at the non-zeroing allocator's value. The grid
  is allocated by `R0480` (`malloc(W·H)`, unpadded) but never read from the file (computed).
  Constants read from the x87 disassembly: `32.0`, `π/6`, `95.0`, `0.5`.
- **The shading table** (`R1107`, called via `R0919` from the relight driver
  `R1368` as `(nLevels=0x60, mode=3, useTint=1)`). It is allocated `malloc(nLevels<<9)` —
  **96 rows × 256 entries × u16 = 49 152 B, row stride 512** — and filled per channel with

  ```
  out_chan = clamp( ((palette_chan + skyTint_chan) * (96 - level)) / 32, 0, 255 )   # trunc toward 0
  entry    = (R>>(8-rBits))<<rShift | (G>>(8-gBits))<<gShift | (B>>(8-bBits))<<bShift
  ```

  Integer, linear per channel; no gamma, no cross-channel term. The bits/shifts are the DirectDraw
  surface's (derived at runtime by `R1242`; shipped = **RGB565**). The palette is the tile
  object's own BMP palette (`obj+0x20`); the table lands at `obj+0x1c` and the renderers pass
  `*([L10222] + 8)`.

  **Row 64 is unattenuated** (`96−L = 32`); `L=0` is ×3.0 and `L=95` is ×1/32, so the level byte is
  an *attenuation* index. 96 rows is exactly the `[0,95]` clamp of the per-vertex byte above.
  **Flat daytime ground is level 46 → ×1.5625** — the shipped art is authored dark for that
  so flat daytime rendering uses level 46, not a ×1.0 multiplier. Only the **sky tint** enters the table; the intensity bytes shape the level.
  One table serves all terrain (the driver rebuilds the first tile slot and `break`s) — sound
  because all 53 `terrain.3d` palettes are byte-identical. For what this mapping actually puts on
  screen, including where it **clips to white**, see [terrain output ranges](sprites.md); mode2 sprite tables are described
  on that same page.
- **Sun** (`R1813`): day/night cycle — hour `= (t/60)%24` sets the sky RGB tint + intensity
  bytes and sweeps the sun angle **θ ∈ [−0.78539815, +0.78539815]** — a truncated-π *quarter* each
  way, so the whole sweep is π/2. The ±π/2 interpretation is retracted; use these quarter-π bounds
  (`TERR-LIGHT-109`). The clock `t` is **`fullTicks + 360`**, `fullTicks = campaign+0x3e0 >> 4` where
  `campaign+0x3e0` is the server sub-tick counter shipped as command `0x64`: **one in-game minute is
  one full tick**, a mission starts at 06:00, and an in-game hour is ≈ 59.5 s of real time at the
  shipped speed index (`SESS-TICK-026`, `SESS-TICK-027`). Five arms write θ:

  | band | hours | store | θ |
  |---|---|---|---|
  | cycle off | — | `L10656` | literal `+0.78539815` |
  | day | 6…17 | `L10657` | **computed** `−0.78539815 + m·0.0021816615`, `m = (t+360) % 720` |
  | dawn | 2…5 | `L10658` | literal `−0.78539815` |
  | dusk | 18…21 | `L10661` | literal `+0.78539815` |
  | night | 22,23,0,1 | `L10664` | **computed** `+0.78539815 − m·0.0065449846`, `m = (t+120) % 240` |

  The **colour** sees six arms, not five: each twilight band splits in two at `L10720` /
  `L10724` and each half is a 120-minute ramp. With `m = t mod 120` and `q(k) = (k·m)/120`
  (one magic division, multiplier `0x88888889`, post-shift **6**, i.e. /120):

  | band | hours | R `L10231` | G `L10613` | B `L10614` | ambient `L05650` | amplitude `L02293` | shroud `0x49c`/`0x4a0` |
  |---|---|---|---|---|---|---|---|
  | cycle off | — | 0 | 0 | 0 | 14 | 32 | 4 / 2 |
  | night | 22,23,0,1 | 0 | 12 | 48 | 32 | 8 | 8 / 4 |
  | dawn 1 | 2,3 | `q(24)` | 12 | `48 − q(48)` | `32 − q(4)` | `8 + q(12)` | 6 / 3 |
  | dawn 2 | 4,5 | `24 − q(24)` | `12 − q(12)` | 0 | `28 − q(14)` | `20 + q(12)` | 6 / 3 |
  | day | 6…17 | 0 | 0 | 0 | 14 | 32 | 4 / 2 |
  | dusk 1 | 18,19 | `q(24)` | 0 | `q(8)` | `14 + q(10)` | `32 − q(12)` | 6 / 3 |
  | dusk 2 | 20,21 | `24 − q(24)` | `q(12)` | `8 + q(40)` | `24 + q(8)` | `20 − q(12)` | 6 / 3 |

  Every value above is an **immediate in `.text`** or that quotient — the routine reads five
  addresses in total (the `ShowTimeFlow` flag and the four `.rdata` doubles the two
  computed-angle arms use) and **no table** (`TERR-LIGHT-119`, `TERR-LIGHT-120`). The three
  bytes are `R, G, B` in ascending address order (`TERR-LIGHT-124`), and the tint is taken
  **unsigned** and added to the palette channel *inside* the level multiply, so it can only
  brighten — all darkening is the level's (`TERR-LIGHT-127`). The six arms form one
  continuous programme: the largest step across any band join is **1**, and night's tuple is
  exactly dawn 1's opening tuple (`TERR-LIGHT-123`). Flat-ground level therefore runs
  **`[45,64]`** over a day against 46 by day, the map-wide sprite level 3 → 8, and terrain
  relief **flattens** as it darkens because `L02293` falls 32 → 8 (`TERR-LIGHT-125`).

  **A consumer must not interpolate this per minute.** The relight cadence below samples each
  120-minute ramp **six** times, so only **24** distinct `(R,G,B,0x494,0x498)` tuples are
  reachable through the unforced path out of the 170 the arithmetic can produce; the forced
  path bypasses the modulo and can reach any of them (`TERR-LIGHT-128`).

  The sun moves during 960 of the 1440 full ticks in a day (`TERR-LIGHT-110`).
Relight runs when `(clock&0xf)==0 && ((clock>>4)+0x168)%20==0`, once per 20
in-game minutes (about 19.8 seconds), or when forced. Its order is sun,
vertex grid, tables, redraw (`TERR-LIGHT-114`). With the cycle disabled,
ambient 0x0e/range 0x20 gives L62/R32. ShowTimeFlow defaults to 1 before main,
persists in the registry and save, and is toggled by N (`TERR-LIGHT-108`).

`R0468` copies global θ to terrain+0x20. The shadow helper reads
that cache and returns θ×2/3, with dead bands returning−0.05 for
−0.05<θ<0 and+0.05 for 0<θ<0.05. Shadow passes convert its tangent to a
16.16 per-row X slope. The slope spans−0.5774..+0.5754 during a day and
is+0.57735 with the cycle off. The unit-body draw discards the helper's
result and does not lean. — TERR-LIGHT-111, TERR-LIGHT-112, TERR-LIGHT-113

Type0 light fields use payload offsets+0x08 (f32 angle),+0x10 and+0x14
(intensities), stored to terrain+0x20,+0x1c,+0x1d. The angle store fills
only four bytes of a double. Relight overwrites all three from sun globals
before reading them; they do not define initial lighting. These are the amended
light-meaning clauses of `ALM-META-009` and `TERR-LIGHT-015`; the separate
record-word interpretation of `ALM-META-009` is also withdrawn.
The separate scalar+0x0c goes to terrain+0x2c and has no identified consumer.
Map-object copies also have no identified reader. These negatives retain
bounded Medium confidence, indirect-call limitations and the skipped-relight/
message-path caveats of `TERR-LIGHT-149`, `TERR-LIGHT-150` and
`TERR-LIGHT-151`. The load path forces relight after publishing the terrain
object behind its stated gate. Known sun-global stores belong to
`R1813`; a file source is not established. Payload+0x18 is separately
consumed as the tile-group mask. — TERR-LIGHT-023, TERR-LOAD-152

The per-axis clamp bounds levels to 30..89 for any height field. The floor 30
is the sine saturation at slope−π/3. Installed grids give 30..70 at
θ=0.78539815 and a ceiling 81 at the ends of the daily θ sweep. Level ranges
depend on θ. The displaced-grid statistics associated with `TERR-LIGHT-016`
and the 30..74 approximation in `TERR-LIGHT-029` are withdrawn; they do not
change the clamp or the current formula.
