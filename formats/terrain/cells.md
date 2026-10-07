# Cell records and Building attachment

[Reference](format.md)

## Structures on the block plane (`TERR-STRUCT-068` (amended, superseded)…`072`, `TERR-PASS-073` (amended, partially retracted))

The three tile-word arms, the type-3 cell and the border above are the whole of what the **ingest**
writes. A placed structure never goes through it. The sim class is the one the image names
`Building` (`CRuntimeClass L10454`, `0x6c` bytes, vptr `R0492`; `Shop` derives from it), and it
attaches **in its own constructor**, after the ingest, to a per-cell record:

- Map load, `R0128`: the world constructor is called at `L00242` (`R0116`, which runs `R0278` and `R0470`). The world pointer is published to the global at `L04624` at `L10488`, with no null check later. `R0489`, the `.alm` type-4 walker, is called at `L09682`.
- `Building` constructor: `R0485` calls `R0486`, then `R0488`, `R1357` and `R0453`.
- `~Building`: `R1833` calls `R1358`, which sets `payload+0x0c` to 0, runs `R0453` and frees the record.

**The cell record**, `world+0x540b4`, keyed by the `u16` cell index, `0x34`-byte payload:

- `+0x00`, u8: the terrain-baseline cost byte, snapshotted at record creation (`L10468`).
- `+0x01`, u8: the terrain-baseline static block byte, snapshotted at record creation (`L10469`).
- `+0x04`, pointer: the ground occupant; sets dynamic bit 6.
- `+0x08`, pointer: the air occupant; sets dynamic bit 7.
- `+0x0c`, pointer: the `Building`.
- `+0x10`, pointer: set and cleared by `R0938` and `R0447`; no plane arm reads it.
- `+0x14` to `+0x28`: six area-effect layer slots. Each non-null slot shifts the cost left by 2, and `+0x20` also blocks. This is sourced by `TERR-STRUCT-074` and `TERR-STRUCT-078`: a loop of six iterations at `L05494` to `L10533` tests each slot against null and shifts the 8-bit cost byte.
  - Which slot is which (`TERR-CELLREC-146`): the slot is `+0x14 + 4*R1074(spellId)`, and the two tables that function dispatches on are read out of the shipped image by `tools/areamove`. `+0x14` is spell 3 Wall of Fire; `+0x18` spell 7 Freezing Cloud; `+0x1c` spell 8 Poison Cloud; `+0x20` spell 19 Wall of Earth, the blocking one; `+0x24` spell 12 Light; `+0x28` spell 17 Darkness.
  - So a Wall of Earth is the only area effect that writes passability, and it does so through this arm rather than through anything in the area module. The registration is `R1075`, which then calls the recompute at `L05479` or `L05480`; the removal is `R1078`, whose slot clear is at `L05489` and recompute at `L05490` (`MAGIC-WALLBLOCK-045`, `MAGIC-AREACOST-046`, `MOVE-085`).
- `+0x2c`, u8.

**Cost byte order and persistence** (`MOVE-084`, `MOVE-085`, `MOVE-086`). `R1087` divides a
cell's cost byte by four once per read whenever the record's count at `+0x02` is non-zero, for every
slot including Wall of Earth, and two reads with no recompute between them divide twice. A recompute
resets the byte from `+0x00`, or from `CostCracked` on the footprint-clearing arm (`TERR-STRUCT-071`), and shifts it left by two per non-null slot in 8 bits. A transit start reads
the cost before any recompute of its cells, and the crossing recompute follows by at least one tick.
No cost plane is saved: a loaded layered cell keeps the ingest byte until the next recompute, and a byte decayed to 0 reaches the two unguarded dividers. Two
inline readers, `R0039` and `R1345`, divide by the byte with no zero guard.

**Who takes `+0x04` and who takes `+0x08`** (`TERR-CELLREC-146`). Not *ground* and *air*
as such: the selector is the actor's movement-domain byte, read through `vt+0x20` in both directions
— `R0458` at `L10901` (`JBE` rejects 0 and below, `<= 2` takes `+0x04`, `== 3` takes
`+0x08`) and `R0057` at `L07106` (clears `+0x08` for 3, `+0x04` for 1 and 2). Domain 2 —
`Ghost` and `Bee`, `MOVE-DOM-028` — therefore shares the slot with ordinary ground movers. Both
slots hold at most one actor and a taken slot fails the entry (`L02040`, `L02041`). `+0x10` is
the **sack** slot: `R0938`'s one caller `R0942` is the sack registration of
`ITEM-SACK-010`, and no plane arm reads the slot. Accessors by returned displacement: `+0x04`
`R0035` / `R1071` / `R1826`; `+0x08` `R1072`; `+0x0c` `R1073`;
`+0x10` `R0955` / `R0446` / `R0143`.

**An actor occupies every cell of its footprint** (`TERR-FOOTPRINT-147` (amended, partially retracted)), the same shape
the building attach uses. `R0050(map, actor)` reads the footprint side once
(the virtual call through slot `+0x1c` at `L05505`), runs two nested loops both bounded by it (`L05506`,
`L05507`) and calls `R0458` once per covered cell at `L05508`; a refusal from any covered
cell stops further iteration (`L07116` to `L10906`, which returns 0), without local
rollback of earlier cells or mover+72/+82..85. Its five callers include the
sub-cell step `R0047` and the arrival `R0039`, so the `n x n` record entries are
rewritten on every cell transit, and `R0453` derives dynamic bits 6 and 7 per cell from the
slot rather than from a separate footprint walk.

This is prefix-preserving failure, not atomic entry. A conflict at each ordinal
of a2x2 footprint leaves the earlier successful actor slots in place under the
selected normal-return paths. — SAV-CELLFAIL-583

Existing domain 1/2 actor entry checks the trigger before testing occupied+04;
its caster call can precede an eventual refusal. Domain 3 checks+08 without that
trigger arm. A missing record takes creation instead:52 zeroed bytes, current
cost/static baselines in+00/+01, then refetch and actor store. The creation path
does not revisit the existing-record trigger branch. Existing record reuse keeps
its baseline, tail and residue. — SAV-CELLENTRY-582

Sack registration is a separate programme. It reads Position+02 through
Sack+10, refuses Dynamic bit 0, and rejects every nonzero existing+10 slot.
Writing an empty existing+10 returns without recompute. Missing-record creation
zeroes 52 bytes, captures current Cost/Static before setting Static bit 5, stores
the Sack, then recomputes. — SAV-SACKENTRY-590

Sack removal clears a present record's+10 without testing zero or identity,
then recomputes. Missing records return 0. Its deletion predicate tests the four
occupant slots, byte+02 plus byte+2c, not the six layer pointers or other residue.
— SAV-SACKREMOVE-591

The deletion arm restores Cost and Static from payload+00/+01 and preserves
current Static bit 4. Dynamic is not copied from the restored baseline and its
record-present bit 5 is not cleared: it retains the preceding recompute result,
plus the conditional bit 4 OR. The recompute itself never reads Sack+10; its
other payload inputs remain active. Allocation-release effects are outside this
direct write-set. — SAV-SACKPLANES-592

Sack lookup separately requires Static bit 5 before hash lookup. The creation
caller uses that lookup; registration's caller dispatches append or deletion,
and two selected removal callers continue without testing removal's result.
Those dispatches do not establish complete caller side effects.
— SAV-SACKCALLER-593

Detach tests the selected actor slot only for nonzero, not equality to the
actor argument. It clears and recomputes before testing whether four occupant
slots, layer-count+02 and operation+2c are all zero. Only that predicate enters
record deletion; other tail/residue bytes do not retain it. Successful detach
copies current Position cell/fractions into mover+86..89; missing-node/zero-slot
refusal returns before those stores. — SAV-CELLLEAVE-584

Entry samples Position cell-X/Y into caller locals before `R1345` and
passes their low bytes to its domain-1 inline cost read. After normal return,
its explicit mover stores are +72, +82, +83, +84 and +85, in that order.
+82/+83 receive the local cell bytes; +84/+85 receive full-X/Y getter low
bytes sampled after the call, Position+04/+05 under `SAV-TOKENPOS-074`.
Word+72 receives the call's AX. Virtual callback effects and pointer stability
remain Unknown; this order does not establish an atomic snapshot or a fixed
runtime+72 value.
After recompute and optional removal, detach copies Position+00/+01/+04/+05
to+86/+87/+88/+89, reloading actor+10 for each byte.
— SAV-CELLFAIL-583, SAV-CELLLEAVE-584, SAV-1161

**Bit 4** (`TERR-PASS-148`). Written by `R1080`, four instructions that OR `0x10`
into both planes for one cell, called from the area module's per-cell add and blast on a fourth
argument meaning *the inner effect does damage*. `wall_of_earth` never reaches it. No mover mask
contains bit 4, and its one consuming reader is `R1655`, which run-length encodes it over the
map interior into a `CArchive`. It is transmitted state, not passability.

The recompute reads the record through a **52-byte snapshot**: `L10526` calls the map lookup and
`L10527 REP MOVSD` copies 13 dwords from `record+0x0c` into the scratch at `world+0x5402c`, and
every `payload+N` below is really `scratch+N`. A lookup miss returns at `L10470` having written
no plane byte at all.

**The recompute** `R0453(world, cellIndex)` — the only routine that turns a record into plane
bytes, and a no-op on a cell that has no record:

1. The terrain baseline is restored first: static is taken from `payload+0x01` and cost from `payload+0x00`.
2. Static gets bit 5 set (`0x20`: this cell has a record), and dynamic is set equal to static.
3. A non-null `payload+0x04` (ground occupant) sets dynamic bit 6 (`0x40`); a non-null `payload+0x08` (air occupant) sets dynamic bit 7 (`0x80`).
4. If `payload+0x0c` is set, a `Building` stands here. Its position is read through the pointer at `obj+0x10` (a pointer, see "the position object"). The footprint bit index is `(cellRow - pos[1]) * obj[0x60] + (cellCol - pos[0])`; the shift count is masked to 5 bits, which is the only reason the engine's own 32-bit intermediate is harmless.
   - If bit `obj[0x64] & (1 << bit)` is set, the cell blocks: static and dynamic both gain `0x05`.
   - Otherwise the cell opens: static and dynamic both lose `0x05` (mask `0xfa`), and cost becomes `costTable[5]` (`CostCracked`, shipped 6).
5. Polarity is settled by branch displacement (`TERR-STRUCT-078`). It is decided at `L07902`, a branch whose short displacement of `0x20` fixes the target at `L07903`, the block that loads the operand `0xfa`. The branch is taken when the bit test finds the bit clear, so a clear bit takes the mask-with-`0xfa` arm and a set bit falls through to the set-`5` arm. The fall-through measures exactly `0x20` bytes and the taken block exactly `0x28` (the jump at `L10484`), so neither can be misaligned by a decode. The shipped `Data.bin` column title "Passability" is therefore the inverse of what the code does with it: a set bit is impassable.
6. Both arms load the operand once and then read-modify-write each plane separately. The block arm loads `5` at `L10456` and ORs it into the planes at `L10457` to `L10460` (plane stores to the planes at `+0x10000` and `+0x20000`). The open arm loads `0xfa` at `L10461` and ANDs it into the same two planes at `L10462` to `L10465`. There is no immediate-operand AND with `0xfa` in the routine: `0xfa` reaches the byte only through a register.
7. For each of the six slots from `payload+0x14` to `+0x28`, a non-null slot shifts the cost left by 2.
8. If `payload+0x20` is non-null, static and dynamic both gain `0x05`.
9. If the old static byte had bit 4 (`0x10`) set, static and dynamic regain it: bit 4 is carried across every recompute.

So a block byte is **derived state**, recomputed per cell from (terrain baseline, occupants,
building). A consumer must keep the baseline, not only the current byte, or it cannot demolish.

The selected original entry prefix at L11030 also copies a found record's
baseline+00 to Cost[low16(key)] and baseline+01 to Static before the selected
call. This is conditional on an original world receiver. In 256 byte
controls the cost store copies the baseline exactly, including 0; current
cost and layer count do not supply its value. Three controls matched to saved
baseline 10 write 10 from current costs 0/10/255. Missing records return without
plane writes. Incoming aliases, native entry reach, subsequent effects and
the fault-time writer remain Unknown; entry/body layout lacks an incoming
target anchor. — TERR-198

The first selected callee at L08162 returns immediately in the selected
x86-32 mode. It reads the four-byte return word from the stack and writes no memory; the incoming
node-payload pointer is not dereferenced. The selected caller's call supplies
the return-stack write. This excludes that unchanged callee as a local
Cost-writing transition, without establishing native C18 reach or writer
history. Faults, interrupts, later callees and native aliases remain Unknown.
— TERR-200

The next selected unresolved callee at R0525 copies its four-byte stack
argument to a new stack slot, then reaches a call to L11036. The argument
comes from incoming EBP+0x540b8 through the selected caller's push. No
pointed-value access occurs before that boundary. The nested call is not
executed by the bounded probe. Native receiver, segment and stack/C18 aliases,
external effects and C18 writer history remain Unknown. The callee cannot
be excluded as a transition through its unmeasured continuation. — TERR-202

The nested L11036 callee returns locally for a zero stack argument: two
register-save stack writes and four stack reads, with no pointee access or
external call. A nonzero argument reaches unresolved L11040 after three
stack writes and one stack read. Its caller conditionally null-tests the
field or uses it as a counted dword-table base, then reloads it for the call;
intervening source changes can alter the supplied value. Native reach,
receiver/table/segment identities, physical stack/C18 aliases and nonzero
continuation effects remain Unknown. — TERR-204

**The footprint** is a rectangle plus two 32-bit masks, all four from the class's `Data.bin`
Buildings entry (`R0486`, params 0/1/4/5, the file's own column titles):

```
obj+0x60  u8    sizeX  the extent along the plane key's LOW byte  (param 0)
obj+0x61  u8    sizeY  the extent along its HIGH byte             (param 1)
obj+0x64  u32   "Passability"     the BLOCKING set    tested by R0453 @L07901
obj+0x68  u32   "BuildingPresent" the ATTACH set      tested by R0488 @L10480
```

`R0488` walks `row = 0..sizeY-1`, `col = 0..sizeX-1`, bit index running continuously, and
attaches `cell = low16(((objRow+row) << 8) + objCol+col)` for each set bit of `BuildingPresent`.
This is addition, not bitwise OR: synthetic out-of-byte coordinates can carry into the other
coordinate. Neither registration body clips against the authored map dimensions
(`UNIT-STRUCTCELL-070`).
`sizeX` bounds the **inner** loop and is added to the position object's byte 0 — the same axis the
tile grid strides by 1 and the `.alm` type-4 record's `+0x00` carries (`ALM-OBJ-061`); the shipped
masks corroborate it, since `Horisontal Bridge` (6×4) and `Vertical Bridge` (4×6) are deck patterns
only under this assignment.
`R1357` refuses a cell that already carries a building and the walk then aborts. A footprint over 32 cells aliases — the 11×4 `Castle` folds bits 32..43
onto 0..11. The `.alm` extension arm overrides `(w,h)` from the record and sets
`Passability = 0`, `BuildingPresent = 0xffffffff`; `kind == 0x21` is class id 33,
`Vertical Wooden Bridge`, so the arm exists to let a map author size a bridge, and it opens every
cell of the rectangle. **The arm is selected by `(w & 0xff) + (h & 0xff) > 0`** at
`L07853`…`L07854`, not by the kind (`TERR-STRUCT-090`) — the caller-supplied bytes are file
`+0x14` → `obj+0x60` and `+0x18` → `obj+0x61` (`ALM-OBJ-062`), and the other two callers of
`R0486` push literal zeros, so this arm has exactly one reachable caller. The
`Passability = 0` store happens on **both** sides of the arm's own `w·h > 32` test
(`L10481`, `L10482`, same immediate): the branch is dead (`TERR-STRUCT-077` (amended, superseded)).

Registration is not transactional. A collision aborts immediately and leaves earlier accepted
references in place; the constructor ignores the return value (`UNIT-STRUCTCELL-070`). Thus the
rectangle, mask-selected cells and successfully attached cells are different sets. Installed collisions do not establish a partial-prefix runtime case; the
conditional non-transactional rule is retained (`UNIT-AREAPOP-075`).

`~Building` independently calls `R1358`. This walks the current dimensions and mask again,
returns at a missing cell record, and clears an existing record's `+0xc` without checking pointer
identity. It does not replay a saved successful-attachment list. Therefore cleanup of a partially
registered object has a conditional ownership hazard; actual destructor reach and HP-to-destruction
ordering remain Unknown (`UNIT-STRUCTDETACH-074`). Ring and blast consumers read each current
cell reference anew, so a surviving alias can be visited repeatedly (`UNIT-AREAVISIT-071`).

The cell accessor does not read Building HP. Recompute reads Position, width and
the blocking mask, but no HP or class gate. Conditional on an unchanged reference,
mask and other cell inputs, HP 1, 0 and -1 therefore produce the same lookup and
blocking/opening contribution (`UNIT-STRUCTNEXT-079`). This conditional consumer
contract is not a lethal-hit lifetime rule: the bounded caller search leaves the
actual reference/mask survival boundary Unknown (`UNIT-STRUCTBOUND-081`).

**Two call sites, three callers** (`ALM-CLS-063`). `R0486` is called only from
`R0485` (`Building`'s ctor) and `R0693`, but the `Shop` ctor `R0490` opens by
calling `R0485` at `L02198` — so a `kind ∈ {0x22,0x23}` placement attaches a footprint too,
always from the table. On the shipped maps that is 450 cells over 50 placements: a 3×3 square with
the bottom-left cell open as the doorway, 400 blocking cells and 373 cells that a ground mover could
otherwise walk through.

**The attach does not always recompute** (`TERR-STRUCT-076`). `R1357` has two exits and only
the one that had to **create** the cell record ends `L07176 CALL R0453`. If a record was
already there with a null `payload+0x0c`, the building pointer is stored (`L07898`) and the
routine returns 1 at `L07900` with no plane byte written — the cell keeps its old bytes until
something else recomputes it. At map load no record exists before the type-4 walk, so 0 of the
17 057 shipped attachments take that exit; a consumer that places a building at runtime must decide
what to do about it. The recompute's callers are eleven routines, not two:
`R0458` (×4), `R0057`, `R0938`, `R0447`, `R1357`, `R1358`,
`R1075` (×2), `R1078`, `R0227`, `R1359`, `R1835`.

**The position object** (`TERR-STRUCT-075` (superseded, contested)). `obj+0x10` is a pointer, allocated and stored by the
base actor constructor (a `0xc`-byte allocation at `L07893` and `L07894` through `R0747`, and the store into `obj+0x10` at `L02083`) and written whole by `R0288`:

```
+0x00  u8   col           +0x01  u8   row
+0x02  u16  (row<<8)|col  the packed cell index the record map is keyed by
+0x04  u8   0x80          +0x05  u8   0x80    the half-cell sub-position
+0x08  ptr  the world
```

The chain from the file is complete **and unpermuted** (`ALM-OBJ-061`): the `.alm` type-4 record's
`+0x00` is read into the local at `[EBP-0x68]` (`L02165`) and `+0x04` into `[EBP-0x74]`
(`L02166`); those are `SAR ...,0x8`-ed at `L02091`/`L02090` and pushed **last** and
second-to-last, so they become the in-memory record's `+0x00` and `+0x02`; `R0489` pushes
`byte[rec+0x00]` last into `R0288` (`L02171`, `L02170`), which makes it `[ESP+0x4]` and
therefore **byte 0**. So file `+0x00` is `col`. **The anchor is the object's own cell and the
footprint runs right and down** — nothing on that chain subtracts `(w-1)/2`, which is what excludes
the centred rival.

A Building can make terrain passable: its AND0xfa mask clears two ingest
bits in each plane. Structure registration must run after terrain ingest
before movement consumes the final planes. The masks read as pictures (`#` blocks, `o` opens):

```
38 Horisontal Bridge 6x4      52 Vertical Bridge 4x6      11 Church 4x3      27 Magic Symbol 3x3
   ######                        #oo#                        o##o               ooo
   oooooo                        #oo#                        ####               ooo
   oooooo                        #oo#                        o###               ooo
   ######                        #oo# (x6)
```

The ground search expands the full 3×3 neighborhood with no diagonal
corner rule. Bridge openings can connect regions disconnected after terrain
ingest; passability is the final state after structure attachment.
— TERR-STRUCT-074

**Save/load.** `Building::Serialize` (`R1585`) round-trips `+0x40/+0x42/+0x44/+0x46/+0x48/
+0x60/+0x61/+0x64/+0x68` and does **not** re-attach; MFC's `CreateObject` ctor builds by the name
`"null"`, which misses the table and skips the attach. It does not need to: the world's Serialize
(`R1360`) stores both plane sweeps (`TERR-PASS-053`) **and** the cell-record map itself
(`L10491` storing, `L10492` loading) — the `u16 key + 0x34` records described by `SAV-CELLREC-017`.
The saved record keys correspond to occupied bit 5 cells in the named save
contract; broader arbitrary-state equivalence is not established.
