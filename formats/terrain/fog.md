# Fog and visibility

[Reference](format.md)

## Visibility bits

Bits 15/14 carry explored/current visibility. They gate animated objects
and partial terrain repaint. Installed ALM words leave these bits clear;
the runtime modifies the light-parser tile plane at `[[view+0x80]+0x0c]`.
That pointer belongs to the render object layout, distinct from the main
reader's simulation input. — TERR-TILE-079 (amended, superseded), TERR-FOG-080 (superseded), TERR-FOG-081,
ALM-TILEVIEW-122

The light parser clears bit 13 only, so an authored high pair `01` can
survive that operation. The known stamp and clear have these local effects:

| Input pair (15/14) | After light parse | Immediately after clear | After an admitted stamp |
|---|---|---|---|
| 00 | 00 | 00 | 11 |
| 01 | 01 | 00 | 11 |
| 10 | 10 | 10 | 11 |
| 11 | 11 | 10 | 11 |

The clear column ends before restamping in the same event. A full event's
final grid and its native interval are separate observations.
— ALM-TILEVIEW-122, TERR-TILECLEAR-169, TERR-DRAWSTAMP-170

The global `01`-unreachability clause of TERR-FOG-082 is partially retracted.
Its shroud classifier remains exact: `11` selects level0, `10` selects level8,
and both `00` and raw `01` select the default level16.

Drawable `R1504` ORs the four corner words before testing `0xc000` and
`0x8000`. Combined `0xc000` sets drawable state 0; `0x8000` sets 1;
`0x0000` or `0x4000` leaves 2 in a one-cell footprint without the permission
bypass. A corner `0x8000` plus another `0x4000` passes as `0xc000`, although
neither word contains both bits. This is not an all-corners-current test.
Downstream redraw is separate from that stored state. — TERR-DRAWGATE-171

The stamp's drawable virtual dimensions read class `+0xd0`; they are not
the simulation actor's footprint virtual. Both named drawable vtables reach
the original decay, permission, sight and position guards. The observed
stores require an admitted LOS cell. Bounded guard controls used a fixed
synthetic LOS field; native LOS extent and event cadence remain Unknown.
— TERR-DRAWSTAMP-170

- **State.** The two fog bits per tile word encode `11` currently in sight, `10` explored but not in sight, `00` never seen.
- **Set.** The stamp is `R0554`, the ground and air unit method at vtable offset `+0x48`, which tail-jumps to the body at `L02578`. It sets the two top bits (`0xc0`) of the high byte in four tile-word positions, at byte offsets +1, +3, +W*2+1 and +W*2-1 from `tile + idx*2` (the four stores are at `L02581`, `L02582`, `L02583`, `L02584`). One immediate sets both bits, so a single stamped word satisfies the four-corner OR. A second writer of bit 15 is at `L07966` in the save-load path `R0099`, which restores the saved run-length record (`SAV-FOG-061`, `TERR-FOG-145`); its saved value is held in a register, not an immediate.
- **Clear.** Only bit 14 is cleared, map-wide, on a period: the mask-out of `0xbfff` at `L01659` in `R0374`. No bit 15 clear is identified in the immediate-form writer set; register-sourced clears are outside that negative. The shroud reader distinguishes explored `0x8000` from visible `0xc000`.

**Guards, all three, before anything is stamped**

0. At `R0554`: when the drawable's decay stage (byte at `+0x15a`) is 3 or more, nothing is stamped.
1. At `L10541`: bit 3 of a per-player 16-bit word must be set. The word sits in the array at `[[view+0x9b4]+0x38]`, on the local participant's record, indexed by the drawable's owning `Player+0x04`. `R0709` writes `flags[localPlayer] = 0x0a` at session setup, and its refresh loop preserves bit 3 (the mask at `L07803`). The other writers are the single-entry message arm at `L07804` and the bulk-copy message arm at `L07807` described below. In single player only the local player passes.
2. At `L10542`: the drawable's sight field `this+0x102` must be non-zero; the constructor leaves it 0.

The method at vtable offset `+0x44`, `R1402`, repeats guards 0 to 2 every tick and adds two more:

3. `(this+0x50 & 0x1f)` and `(this+0x54 & 0x1f)` must both equal `0x10`, an eighth-of-a-cell grid.
4. The pair `this+0x8`, `this+0xc` must differ from `this+0xc0`, `this+0xc4`: the drawable moved since the last stamp.

**The stamped set.** It lives at `CMapView+0x17cc`, a 41x41 grid of 32-bit values that `R0284` zeroes on every stamp.

- Seed: the centre entry `mask[20][20]` (at `view+0x24ec`) is `(sight >> (8-k)) + (1 << (k-1))` with `k = view+0x3f38 = 7`, so the unit of the field is 1/128 cell.
- Walk: Chebyshev rings r = 1..19, four edges each; the walk stops at the first fully blocked ring.
- Cell rule (`R0285`): `mask[dx][dy] = mask[pred[dx][dy]] - (cost[dx][dy] + h(cell) - h(obs))`, and the cell is visible iff the result is above 0. `h` is the signed-byte grid at `map+0x10` (the `.alm` type-2 Altitudes grid). `h(obs)` is sampled once, so the term is per-cell altitude against the observer's, not a slope along the ray.
- Tables built once in the `CMapView` constructor by `R0283`: `pred` at `view+0xaa8` is a 41x41 grid of signed byte pairs, one step toward the observer, in three zones (`j < i>>1` gives (-1,0), `j > i<<1` gives (0,-1), otherwise (-1,-1), mirrored). `cost` at `view+0x3210` is a 41x41 grid of 16-bit values `ftol(128 * sqrt(i^2+j^2) / max(i,j))`: 128 on an axis and 181 on a diagonal, so the revealed region is a disc.
- Margin: cells outside `[7, W-7) x [7, H-7)` are never evaluated and stay 0.
- Clip: 19 cells. The maximum shipped scanRange is 12, so the clip never binds on shipped data.

**Sight.** `drawable+0x102` is copied from `actor+0xa4` by the state sync (`L10563`, `L10564`). For a hero, `R0280` writes `ftol(((mind+reaction)/25 + 4) * 256)`, in units of 1/256 cell. For a non-hero, the streamer's slot 11 targets `actor+0xa5`, the high byte, in whole cells (`Data.bin` Units title 11 `scanRange`, Humans title 9 `ScanRange`).

Consumer consequences: `Index` is in range on 82/82 shipped classes but a decoder should still
bounds-check it; `DeadObject` is a subscript into the same class array; `FireObject` is **not** a
class reference (`REG-OBJ-047`). **And the animated arm is not dead**: it fires on exactly the
cells the local player can currently see, so a shipped map's fires and trees animate inside the
field of view and hold frame 0 outside it. A consumer that implements only the plain arm
reproduces every shipped map at load time and never afterwards.

**Shroud / fog** (`TERR-FOG-037` (amended, superseded)): after the terrain, the same cell quad is darkened in place by
`R1509` (Gouraud level ramp, `dst = LUT[level][dst]`), or by `R1508` (level `0x10`,
degenerate quads filled with colour 0) or `R1510` (level `8`, degenerate quads halved as
`(px>>1) & mask`). These reuse the terrain quad and the same step-table edge walk.

**What the renderer branches on, and the levels it turns the pair into**
(amended `TERR-FOG-082`, `TERR-FOG-083`, `TERR-FOG-084`, `TERR-FOG-085`). The shroud pass does not read the
tile word directly: `R1657` projects the pair onto a **per-vertex** dword grid every frame.

- **Grid.** `CMapView+0xa0` holds one 32-bit value per lattice vertex, `(cols+7)*(rows+11)` entries. `R1280` allocates it together with `+0xa4` at `+0x70 = (cols+7)*(rows+11)*4` bytes. It is zeroed at `L10578`, filled, then copied to `+0xa4` at `L10579`. `R1657` is its only content writer and `R1836` frees both.
- **Classification per vertex.** The 16-bit word is read from the plane at `[CMapView+0x80]+0xc` (the light-parser render plane) at `L10567`, and its top two bits (mask `0xc000`, at `L10568`) decide the level:
  - `0xc000` gives level 0 (`L10570`), state `11` in sight, no shroud drawn at all;
  - `0x8000` gives level 8 (`L10572`), state `10` explored, half brightness;
  - anything else gives level `0x10` (`L10573`), state `00` never seen, black.
  Raw `01` takes the same default arm as `00`. The known stamp writes `11`; the periodic clear maps `01` to `00` and `11` to `10`. This does not rule out raw authored `01` or an additional writer.
- **Dispatch.** `R0379` reads the four corner vertices (`L10580` to `L10581`) and branches (`L10582` to `L10583`) on all four equal to 0, all equal to `0x10`, all equal to 8, or otherwise. A cell whose corners disagree is a gradient; one fog value per cell cannot draw it.
- **Table.** The pointer at `L03341` is built once by `R0788`: 17 rows (L = 0..16) of `stride` 16-bit entries, where stride is 65536 or 8192 depending on the flag at `L03346` (row addressing at `L07810`). `LUT[L][px]` scales each channel index by `v * (16-L) >> 4`, saturates at `0xff` and repacks. Row 0 is the identity, row 8 is `(px>>1) & [L07813]`, row 16 is 0. These were checked against the engine's own two table-free fast paths over all 65536 pixel values, RGB565 (mask `0x7bef`) and RGB555 (`0x3def`): 65536/65536 on all three rows. The word at `L07813` is the sum over channels of `(0x7f >> (8-bits)) << shift` (`L10585` to `L10586`).
- **Clock.** `R0374` clears bit 14 over the whole map (W*H words), then re-stamps every drawable in `+0x9b8`. Its one caller `R0334` gates it twice: a non-zero word at `L01661` skips the clear and freezes the layer (`L02610`); and the counter at `CMapView+0xa70` (`ANIM-CLOCK-001`) must be a multiple of 32 (mask `0x1f` at `L10588`). So the clear runs once every 32 presentation ticks, about 2 s at the default speed index and about 4 s at the slowest. Neither `server+0x00` nor `server+0x04` appears on this path: the fog is not simulation state, is not hashed, and stops when rendering stops.
- **Gates.**
  1. The shroud pixels above.
  2. `R0378` ORs the four corner words; when the top two bits of the result are not both set it writes `drawable+0x78 = 1`, the sprite pass's guard (`L10595` to `L10597`). Aggregate `0xc000` stores 0. One `0xc000` corner suffices but is not necessary: separate `0x8000` and `0x4000` corners also pass. `+0x7c` latches `+0x10c` on change.
  3. AI visibility uses a separate array (below); passability reads the block planes (`MOVE-DOM-027`). The direct immediate-reference negative does not cover byte-wide tests of the high half.
- **Persistence.** Bit 15 is saved; bit 14 is recomputed. The save tail's `&YA1` registry stores `Fog.FirstState` (int32) and `Fog.Data` (int32[]) as run lengths over W*H cells in `idx = col + row*W` order. The runs sum to the cell count. Store is `R0084`; load is `R0099`, which ORs in the decoded word. The ALM plane supplies the initial state; the saved runs restore exploration. The 32-tick clear and stamp rebuild bit 14. The no-persistence clause of `TERR-FOG-087` is retracted; its earlier compressed-body size argument did not describe this tail record (`SAV-FOG-061`, `TERR-FOG-145`).
- **Edge.** The map border is black because it is never seen: level 16, the same mechanism at its maximum, and the 7-cell stamp margin of `TERR-FOG-080` (superseded) is why it is never lit. There is no separate edge treatment.
- **Reveal.** Permission to stamp is bit 3 of `[[mapView+0x9b4]+0x38][player]`, and that array has three writers: the setup store in `R0709`; dispatcher opcode 33 (`0x21`, arm `L07808`), which writes one entry at `L07804`; and dispatcher opcode 45 (`0x2d`, arm `L07805`), which resizes the array and copies it wholesale out of the message body at `L07807` (destination `[[view+0x9b4]+0x38]`, source `msg+0xe`, length `[msg+0xa]*2` bytes). The opcode is `3 + (slot - L02524)/4`, from the dispatcher's own jump-table indexing. A wholesale copy carries no displacement in the bulk copy.

TERR-FOG-086 is partially retracted only for its individual-corner equivalence.
The aggregate gate, field stores and latch described above remain supported;
the admitted raw domain is separate from the known stamp/clear-produced states.

**The AI's vision is a second implementation of this same algorithm, not this one**
(`TERR-FOG-088`). It lives in an object embedded at `world+0x58ee8` — so `AI-SIGHT-006`'s byte
map at `+0x2a008` **is** `world+0x82ef0` — with the same recurrence and its own `pred`
(`+0x22000`), `cost` (`+0x28000`) and accumulator (`+0x24000`, whose centre cell `[20·64+20]` is
the address that row calls the field `fog+0x25450`). Its `k` is **`[Scanning] ScanShift` from
`World\Data\map.reg`** (ships 7), while the view's `k` is the compiled constant at `L10553`;
they agree only because the shipped value equals the code default, so **editing `ScanShift` moves
the AI's sight and leaves the player's fog untouched**. Its byte map is cleared by
`R0145`, whose two callers stamp **one actor** (`R0114`) or **a whole group**
(`R0110`). Different storage, clock, consumer and lifetime — do not merge them.
`AI-SIGHT-006`'s two location clauses (the array `R0134` zeroes and the field
`fog+0x25450`) are superseded by `TERR-FOG-084`.

Both of the server's tables are built by **`R0273`**, called once from the object's init
`R0274` in each of the four world constructors, before the map load; there is no other
writer and no other reader than `R0136` (`TERR-SIGHT-115`, `AI-LOS-087`…`AI-LOS-091` —
the layout and the region are specified in [AI](../ai/format.md), which is where the stamp is
consumed). The same init builds a **fourth** grid at `+0x20000` holding `di² + dj²`, which serves
a disc test in `R0409`/`R0410` and has nothing to do with line of sight.
The view and server use the same predecessor/cost algorithm, including
four final stores repairing cells(+1,0) and(-1,0). At k7, flat-ground
scanRange1..6 yields 9,21,45,69,105,145 cells; range 19 yields 1253. The
127-cell radius 6 result in `TERR-FOG-080` is withdrawn. The AI seeds from
the high byte at actor+0xa5 and truncates fractional sight, while drawn fog
uses the full u16. Their playable rectangles also differ by one cell on
each edge. — TERR-FOG-117, AI-SIGHT-093, AI-SIGHT-094, TERR-FOG-118

**The playable rectangle** (`TERR-SIGHT-116`) is four bytes at `world+0x58ee0..0x58ee3` written by
the map load `R0278` as `(8, 8, W-9, H-9)`, with the same corners packed as words at
`+0x58ee4` = `0x808` and `+0x58ee6`. Every sight ring cell is tested against them; the compares
are **byte-wide** while the height and visibility indices are exact 32-bit, so on a 256-wide map a
true column of `-9..-11` wraps to `245..247`, passes, and reads the previous row.
