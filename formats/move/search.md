# Domains, path search and route extraction

[Reference](format.md)

## Structure

```
per unit, per move order (R0178, once per tick):
  mid-transit?          -> advance the sub-cell position and return
  target changed, or  N dynamic searches since the last static one?
                        -> STATIC search  (plane world+0x10000, no occupancy) -> static route
  no dynamic route, or M ticks since the last dynamic search?
                        -> next static waypoint adjacent AND blocked?
                             -> wait, facing it. no step, no search.
                           else DYNAMIC search (plane world+0x20000, with occupancy)
                                aimed a few waypoints down the static route -> dynamic route
  any search that fails to label its goal
                        -> substitute the labelled cell with the SMALLEST LABEL near the
                           goal (or near the target actor) and route to that instead
  step: claim the next cell on the shared plane, turn toward it, or advance the sub-cell
        position by the frozen per-axis step -- exactly once per SUB-TICK
```


## Movement domain (`MOVE-DOM-024…028`)

One byte, `actor+0x4a`, read through slot `+0x20` of all three simulation-actor vtables
(`R1344`). It is **not** a passability property with side effects — six things branch on it:

| what | rule | where |
|---|---|---|
| block mask | `1 → 0x41`, `2 → 0x44`, `3 → 0x82`; anything else stores **nothing** and leaves the constructor's `0x41` | `R0471` |
| occupancy bit | `< 3` sets/tests bit 6, `== 3` bit 7, on the dynamic plane | `R0058`/`ac70`/`ad20`/`af40` |
| step cost | `== 1` reads `cost[dst]`; every other value takes flat 2 straight / 3 diagonal | 9 sites, `MOVE-COST-002` |
| speed | `== 1` divides by `cost[cell]` and tilts by height; 2 and 3 take the raw class speed; 0 and `> 3` give **0** | `R1088`, `R1345` |
| AI targeting | a domain-3 candidate counts **one cell farther** to a decider that is not domain 3 | `L00225`, `AI-ACQUIRE-002` |
| corpse | `> 1` slams the decay counter to −1000 → immediate teardown, so no corpse | `L04398` |

**The mask is a cache, written twice in the image and never refreshed**: `R0206` (the mover
constructor, `0x41`) and `R0471`'s three arms, whose only two call sites are inside object
construction. It survives a save because `Unit::Serialize` hands the whole `0xb4`-byte mover to
`CArchive::Read`/`Write` (`R1348`) before its own storing/loading branch.

**What each domain may enter.** Terrain and objects are *different bits*, so the domains are not
three grades of one thing. Counts are the 38 shipped maps' static plane after the structure pass,
880 704 cells:

```
                                    dom 1  0x41   dom 2  0x44   dom 3  0x82
terrain      water                      77 845             0             0
             Mountain (class 8)        117 572             0             0
             tile-word bit 13            3 944             0             0
non-terrain  .alm type-3 object         69 676        69 676             0
             structure (Passability)    13 527        13 527             0
             8-cell border             158 976       158 976       158 976
             ---------------------------------------------------------------
             blocked                   441 540       242 179       158 976  (50.1 / 27.5 / 18.1 %)
             + runtime: bit 6 occupant  yes           yes            no
                        bit 7 occupant  no            no             yes
```

So: **domain 2 crosses water and mountain and is stopped by every object, building and ground
occupant; domain 3 is stopped only by the map border and by another domain-3 occupant.** They
disagree on 83 203 cells. 8-connected, the free set is in more than one piece on 23 of 38 maps for
domain 1 and on **0 of 38** for domain 3 (whose one piece is the whole interior).

**Nothing narrows the verdict afterwards.** The predicate is `block[cell] & mover[5]` over the
`n × n` footprint and nothing else: no height term anywhere on the search path (the height plane
has 6 references image-wide, none in the search), no corner rule in the extraction, no plane test
at step time, and no order-time check — the block planes have no reader outside
`R1349…R0227`, so no interface or campaign routine can run one.

**Who is what.** Every one of the 210 `Data.bin` Humans rows carries `movementType = −1` → ground.
Of 57 named Units rows, 16 are non-ground: Ghost and Bee (2), Bat_Sonic and Dragon (3), four
difficulty variants each. The two domain-3 classes are exactly the two with `units.reg` `Z != 0`,
i.e. the two `CAirUnit`s. The hireable roster is Catapult, Ballista and 13 Humans, so **no
player-ownable unit is a non-ground mover** — with one open exception, `Control Spirit`, which
spawns a caster-owned `Ghost` (domain 2).


## Path search — `R0053`

A **double-buffered label-correcting wave**. Not Dijkstra (no extract-min), not A\* (no `h`).

```
label:    u16 plane at world+0x30000, 65536 cells, 256-stride, 0xffff = unlabelled
frontier: two lists used alternately, both 4096 entries, counts u16 at
          world+0x54008 and world+0x5400a
            footprint 1  -> u16 packed cells at world+0x545b4 / world+0x565b4
            footprint >1 -> byte pairs   x world+0x50008 / y world+0x51008
                                     and x world+0x52008 / y world+0x53008
seed:     label[src] = 0, src pushed to list A
generation:  for every cell in the current list, for each of the 9 neighbours
             (centre included, scan order dx = -1,0,+1 outer, dy = -1,0,+1 inner):
               skip if (blockPlane[c] & mover.mask) != 0 over the mover's n×n footprint
               g = label[cur] + step(cur -> c)
               if g < label[c]:  label[c] = g;  append c to the other list
             then swap the lists
stop when:   label[goal] != 0xffff  (tested between generations)
          or the current list is empty
          or generations >= budget
```

**Step costs** (`MOVE-COST-002`), where `cost[]` is the byte plane at `world+0`:

| mover | straight | diagonal |
|---|---|---|
| `movementType == 1` (ordinary ground) | `cost[dst]` | `cost[dst] + (cost[dst] >> 1)` |
| any other movement type | `2` | `3` |

The diagonal is 3/2 of the straight step **truncated**, and the cost byte is the *destination*
cell's. There is no heuristic and no distance term in any label.

**Budget** (`MOVE-TERM-003`), with `D = max(|Δx|, |Δy|)`:
the caller-substitution clause is partially retracted; these budget relations
stand.

| search | footprint 1 | footprint > 1 |
|---|---|---|
| static | `max(StaticScanAhead, D>>2) + D`, or **1000** under the override below | `StaticScanAhead + D` |
| dynamic | `max(DynamicScanAhead, D>>2) + D` | `DynamicScanAhead + D` |

**The 1000-generation override is not an edge case — implement it or a human player's units will
refuse long detours.** It applies to the static footprint-1 search when both hold
(`L06770`…`L06775`):

1. `actor+0x14` — the owning `Player` — has a complete dword zero at+0x28.
   A nonzero value whose low byte is zero does not qualify. This retains
   `UNIT-OWNER-009`'s local predicate; its universal authorship/value-space
   interpretation is partially retracted. Authored copies and serializer
   load can supply the complete field. — ALM-140
2. the **goal's own `tokenSize × tokenSize` footprint** is free of the mover's domain mask
   `mover+0x5` on the static plane `world+0x10000`. A `tokenSize <= 0` skips the scan and takes
   the override.

Outside the override, one generation advances the wave by one ring, so the slack over the
straight-line distance is only `max(scalar, D>>2)` rings. A detour needing more fails, and the
search then substitutes a goal near the requested one (below) and extracts to that instead.
**There is no node budget, no frontier capacity check and no closed set.**


## Substitute goals (`MOVE-ALT-018…022`)

When the loop ends with `label[goal] == 0xffff`, `R0053` picks a substitute **itself** — the
two pickers are called from its own tail and from nowhere else in the image — and feeds it to the
ordinary route extraction as the destination. The substitute is **not** written back to the actor:
the ordered target (`mover+0x74`/`+0x76`) and the order block are untouched, so the next search
starts from the same request again.

```
label[goal] != 0xffff  -> ordinary extraction to the goal
staticFlag != 0        -> ring picker around the requested cell, limit = (D>>2) + 4   [altTarget ignored]
staticFlag == 0        -> altTarget != 0 ? contact-ring picker(mover, altTarget)
                                         : ring picker around the requested cell, limit = 8
picker returned 0      -> no route (static: the route list is freed)
```

`altTarget` — the search's eighth argument — is a pointer to a **target actor**, not a cell. It is
supplied only by the go-to-actor order (`R0043` → `R0055`) and is consulted only on
the dynamic branch.

**Both pickers choose by the same rule and read only one plane: the label plane `world+0x30000`.**
A candidate qualifies iff it carries a label, i.e. iff the wave that has just failed reached it —
which already folded in the plane it ran on, the mover's footprint and its mask. Among candidates
the **smallest label** wins, so the substitute is the cell *cheapest to reach from the mover* under
the step costs above, **not** the cell nearest the click. Neither picker contains or calls a
passability predicate, a cost read or an occupancy test.

*Ring picker* `R0392(world; mover, 0, cell, limit)` — rings `r = 1 … limit-1` around the
requested cell; per ring, `i = -r..r` and four probes `(x+i, y+r)`, `(x+i, y-r)`, `(x+r, y+i)`,
`(x-r, y+i)`; the **whole ring** is scanned and the strict minimum kept; the loop stops after the
first ring that yielded anything. The centre is never probed. The mover's footprint side is fetched
and branched on but the two branches are the same code. Returns the packed cell or 0.

*Contact-ring picker* `R0435(world; mover, target)` — the box
`x ∈ [tx-nM-k, tx+nT+k]`, `y ∈ [ty-nM-k, ty+nT+k]` for `k = 0..7`, `nM`/`nT` the mover's and
target's footprint sides; at `k = 0` its perimeter is exactly the set of mover origins whose
footprint touches the target's without overlapping. The walk enters where the line between the two
actors' fine footprint centres crosses the box and proceeds in **both** directions to close the
ring; the strict minimum label wins. Returns the packed cell or 0.

On the **static** branch only, the substitute's Chebyshev distance from the request is compared with
1 — or 2 when the requested cell carries a cell record (`TERR-PASS-051` (amended, superseded) bit 5) with a non-zero
`+0x0c` — and exceeding it queues a UI message. It does not cancel the move.


## Route extraction — `R0438` (dynamic) / `R0437` (static)

Walk downhill through the label field from the terminating cell to the seed, choosing at each step
the 3×3 neighbour minimising `label[nb] + step(nb -> current)` with the same four cost arms.
Constraints: the neighbour must be labelled, and `8 <= x <= W+8`, `8 <= y <= H+8`
(`world+0x50000` = W, `world+0x50004` = H). Passability is **not** re-tested and there is **no corner
rule** — a diagonal between two blocked cells is legal.

**Tie-break, and it is not symmetric:** the straight/centre arm accepts a candidate on `<=`, the
diagonal arm only on `<`. Among equal-cost neighbours the **last straight one in scan order** wins,
and a diagonal never displaces an equal straight.

Output: a doubly-linked list of 12-byte nodes on the actor — `+0x00` toward the goal, `+0x04` toward
the unit, `+0x08` the packed cell `(y<<8)|x` — built goal-first, the unit's own cell unlinked before
returning, consumed from the tail. A walk exceeding **1000** steps discards the whole route.

Two routes per unit, one per plane:

| | static (unit-blind) | dynamic (unit-aware) |
|---|---|---|
| plane read | `world+0x10000` | `world+0x20000` |
| first / last / count | `actor+0x160` / `+0x164` / `+0x168` | `actor+0x17c` / `+0x180` / `+0x184` |
| free list / pool | `actor+0x16c` / `+0x170` | `actor+0x188` / `+0x18c` |

First/last here follow node+0 from the near-seed end to the resolved end.
Older movement prose calls those ends tail/head. For an initially empty list,
extraction with endpoint=seed allocates and removes one node, leaving both
end pointers and count zero. Ordinary one-, two- and three-transition
controls retain their non-seed cells. Extraction prepends to an existing
list; it does not clear that list. — MOVE-080

Static search clears its old list before its requested-seed shortcut.
Dynamic search alone preserves an old list on that shortcut and on a picker
miss. A nonzero picker answer equal to the seed reaches extraction and can
leave an empty route. The empty-input premise matters for interpreting
dynamic count. Private topology replay uses declared actor/heap/ownership
services; native caller list invariants and picker B remain Unknown.
— MOVE-081


## Parameters — `data/map.reg` `[Path Finding]`

Read by `R0116`. The shipped file carries all seven equal to the code defaults, in both the
EN and RU releases.

| key (file clamps names to 15 chars) | world | default = shipped |
|---|---|--:|
| `SpeedMultiplier` | `+0x58db4` | 8 |
| `StaticScanAhead` | `+0x585b4` | 5 |
| `DynamicScanAhead` | `+0x585b8` | 3 |
| `StaticRefreshRate` | `+0x585bc` | 16 |
| `DynamicRefreshRate` | `+0x585c0` | 32 |
| `DynamicByStaticLookup` | `+0x585c4` | 3 |
| `StaticIsntNeeded` | `+0x585c8` | 5 |
