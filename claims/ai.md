# Claim registry — AI (target acquisition and the diplomacy matrix)

Level 2 ledger. Index: [registry.md](registry.md) · spec: [`formats/ai/format.md`](../formats/ai/format.md).
Legend: `✔ promoted` (also in the spec) · `● active` · `✖ retracted`. IDs are permanent.

Not a file format — the **simulation area** that answers "whom does this unit fight, without being
told". It sits between `claims/move.md` (how it gets there), `claims/hero.md` and `claims/unit.md`
(what happens when it arrives) and `claims/alm.md` (what a scenario authored). Opened by
[EXP-0074](../experiments/EXP-0074-target-acquisition/).

Vocabulary, extended by [EXP-0082](../experiments/EXP-0082-ai-behaviour/): a **group** is the
`0x48`-byte `CObList` subclass at `player+0x24`, `grp+0x1c` its authored group id, `grp+0x44` its
owning `Player`, and `grpAI = grp+0x3c` its `0x50`-byte AI record — order `grpAI+0x20`, centre
`grpAI+0x28`, working radius `grpAI+0x2d`, base radius `grpAI+0x38`, AI-enable `grpAI+0x45`,
patrol path `grpAI+0x4c`. The group is the unit of AI: `actor+0x50`, the per-actor state, is
evaluated only while `grpAI+0x20 == 0`.

Vocabulary, fixed once here. The **actor** is the simulation object of vtable family
`L00001`/`L00002`/`L00003` (`TERR-MOVE-054`); `actor+0x14` is its owning `Player` and
`Player+0x28` is 0 exactly for a human participant (`UNIT-OWNER-009`); `actor+0x154` is the mover,
`actor+0x158` the **order block**, `actor+0x12c` reach, `actor+0xa5` sight. The **session** is the
`0xc320`-byte object at `[L00004]` — the AI's owner, holding the world at `+0xa50`, two scratch
lists at `+0xba4`/`+0xbc4` and the diplomacy matrix at `+0xa9c4`. The **world** is the `0xa4558`-byte
object `claims/terrain.md` describes.

**EXP-0243 reconciliation.** `AI-GROUPSEE-068`'s five-phase builder account survives and now has
its outside caller's exact use: Prismatic Spray consumes list A then list B in source order, with
the primary target outside both. `AI-DIPLO-084`'s formerly Medium “spell-cast family” identity is
now direct: the complete caller is the id-14 Prismatic Spray selector, which flips hostility before
building secondaries. Its secondaries pass the builder's directional first-member diplomacy test;
the primary does not.

## Commands during an action cycle

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-354 | States 3, 0xd and 0xe apply from actor fields at phase 5; clearing the pending order alone does not cancel their progress, but the executor tail can replace active fields. | High / Medium / Unknown | ● active | [EXP-0414](../experiments/EXP-0414-cycle-orders/) |
| AI-355 | The admitted non-self attack command changes the queued victim without directly writing the active victim or progress; the next cycle boundary depends on the executor tail and position. | High / Medium / Unknown | ● active | [EXP-0414](../experiments/EXP-0414-cycle-orders/) |
| AI-356 | Pick-up and manual cast setters install pending data immediately; their ordinary executor rows load action fields only at progress zero, after a retained strike's application and recovery. | High / Medium / Unknown | ● active | [EXP-0414](../experiments/EXP-0414-cycle-orders/) |
| AI-357 | Scroll casts share the manual cast setters but write actor+0x68 before the progress gate; cast and weapon-proc cleanup can clear that pointer during a running cycle. | High / Unknown | ● active | [EXP-0414](../experiments/EXP-0414-cycle-orders/) |

### AI-354

The raw-PE action table sends states 3, 0xd and 0xe to `L00005`. The body has phases 0, 5 and 7. Phase 0 loads charge and clears `actor+0x136`; phase 5 decrements `actor+0x6c` and applies at zero; phase 7 decrements recovery and at zero stores phase 0 and completion 1. Application calls read the current actor fields: `L00006` passes `actor+0x5c` to `R0001`; `L00007` passes `actor+0x5c` and Spell `actor+0x64` to `R0002`; `L00008` passes cell bytes `actor+0x60/+0x61` and `actor+0x64` to `R0003`. The 346-instruction range `L00005..L00009` has no undecoded byte and no `+0x158` operand. This bounded body reads no pending order or progress field.

`AI-350` and `AI-351` establish the progress-before-pending switch and the Stand Ground stores. The whole executor restores action 3 from progress 1, or 0xd/0xe from progress 2 and `ord+0x5c`, before actor dispatch. Clearing pending order alone therefore does not select a cancellation arm. With active fields retained, the application reads the loaded target rather than the new pending target. This is an application-call argument result; it promises no damage or successful per-spell effect.

The common tail is a live exception to field retention. Progress arms jump to `L00010`. When `mover+0x98` is nonzero, the tail clears it and, for command state other than 1, 0xa or 0x17, calls `R0004` at `L00011`. Its in-position target arm calls `R0005` at `L00012`. That complete leaf writes progress 1, counter 0, action 3 and `actor+0x5c = ord+0x0c` without changing phase `actor+0x58`. It can change the application argument before the actor body runs on that tick. State 1 and 0xa have different cleanup arms, and state 0x17 writes progress 0xff; none supplies a universal retention rule.

**Confidence.** High for the named body, table, field reads and direct tail writes. `DisasmFnNoBytes`, `DisasmRangeNoBytes` and `EnumRefs` on the owned repaired Ghidra 12.1.2 project are bound to the identical EN/RU image digest. `derive.py` independently checks relative branch targets, dispatch entries and load-bearing MOV/CMP displacements from the containing PE section. Medium for retained-field next-tick composition. The image-wide `disp:`/`imm:` and overlapping-width census is preserved, not a complete typed writer census: outside the named routines, bases, arithmetic aliases, embedded objects and bulk copies remain unexamined. No High image-wide absence is claimed.

**Unknown.** Native command/tick timing, whether the tail prerequisite is reachable during a native running strike or cast, effects of untraced callees and indirect calls, per-spell effect/target-loss rules, and other external writers. Death and physical application validity retain `HERO-CADENCE-115`'s separate scope.

### AI-355

Opcode 0x19 calls `R0006` at `L00013`. Its admitted, non-self arm stores command state 3 (`L00014`), `ord+0x0c = new victim` (`L00015`), reach in `ord+0x14` and pending order 0 (`L00016`). The member helper `R0007` clears action `actor+0x54` at `L00017`. Neither body directly stores progress `ord+0x09`, active victim `actor+0x5c`, phase `actor+0x58`, countdown or completion. The self-target and guard-refusal arms leave command state 0xc and the post instead (`AI-CMD-054`). The store table retains writes overlapping the named fields, with register bases resolved from the actor parameter and its `+0x158/+0x154` loads for these named sites.

When the tail and other writers retain active fields, a running phase-5 application passes the old `actor+0x5c` even though `ord+0x0c` differs (`AI-354`). It can still produce no damage when physical validity/reach refuses the target. Group order 0 calls `R0008`; command state 3 calls `R0009` with the queued victim. Its ordinary non-spell arm writes pending 5 and that victim at `L00018/L00019`, independent of progress. Its special-cast and target-acquisition arms are separate alternatives. Pending row 5 copies the victim into the active field only after progress zero and a successful position test (`AI-350`, `AI-PURSUE-040`).

Let T0 be the retained cycle's recovery-zero tick. If the counter/completion gate clears progress at T0+1, the common tail remains quiet, pending 5 has been reissued, and the position test passes, row 5 installs the new cycle and phase 0 starts at T0+2 (`HERO-CADENCE-112`). Reissue can occur before progress clears; it need not wait for another group evaluation after T0+1. Failed position, different group/actor arms, counter wrap, disabled AI and the common-tail replacement prevent a universal tick number. The common tail can change the active victim earlier without resetting phase, so old-victim retention is conditional, not the only static route.

**Confidence.** High for the named setter, helper, state-3 writer and pending-row stores. Medium for the conditional old-victim and T0+2 compositions; no runtime observation exists. The reference/store census preserves wider neighbouring stores and orphan rows without inferring their base from displacement alone.

**Unknown.** The actual native next-cycle tick, whether common-tail acquisition chooses the commanded victim, admitted-command timing, untraced group/allocation and target-selection callees, and every command not named here.

### AI-356

Pick-up opcode 0x21 validates the cell, stores the message word at inventory `+0x1c`, and calls `R0010` at `L00020`. Its direct stores are action 0 (`L00021`), command state 2 (`L00022`), queued cell `ord+0x0a` (`L00023`) and pending 0 (`L00024`); it writes neither progress nor active victim. Per-actor state 2 writes pending 1 while cells differ, or pending 7 at `L00025` when they agree. Only progress zero reaches row 7. On its first pass row 7 writes action 2 (`L00026`), whose actor body runs the sack lookup and transfer path. On the completion pass it calls the member helper, stores command state 0xc (`L00027`), pending 0 (`L00028`) and `ord+0x50 = 1` (`L00029`). This bridge has no direct progress or active-victim store. The contents transferred by its callees are outside this claim.

Actor-cast opcode 0x1e calls `R0011`; its admitted Spell arm clears action, writes command state 0xd, queued actor `ord+0x28`, Spell `ord+0x30`, range and pending 0 (`L00030..L00031`). Cell-cast opcode 0x1f validates the cell and calls `R0012`; its admitted arm clears action, writes command state 0xe, queued cell `ord+0x3c`, Spell and pending 0 (`L00032..L00033`). Neither setter body directly stores progress, phase or active victim. Their missing-Spell arms fall back to command state 0xc. The per-actor command arms reissue pending 8 or 9; row 8 writes progress 2, action 0xd, active victim from `ord+0x28`, Spell and `ord+0x5c=1` (`L00034..L00035`). Row 9 calls `R0013`, which writes progress 2, action 0xe, cell bytes, Spell and `ord+0x5c=0` (`L00036..L00037`). Its `ord+0x60==0` cleanup can clear action on that installation tick; progress 2 restores the variant later.

**Confidence.** High for the named direct stores and progress-zero placement of rows 7/8/9. Medium for the conclusion that these ordinary row effects follow a retained strike's application and recovery. It is conditioned on `AI-354`'s active-field retention, not a native dispatch-order observation. No direct nonzero-progress refusal is introduced by the command setters themselves.

**Unknown.** Native event timing, the common-tail prerequisite during these commands, untraced validation/transfer callees, and successful spell effects after the actor-side application call. Scroll pointer handling is `AI-357`.

### AI-357

Opcodes 0x25 and 0x26 enter the scroll arm at `L00038`. The arm obtains the inventory item, requires definition byte `+0x3c==0x29`, writes `actor+0x68 = item` at `L00039`, creates Spell from definition `+0x40`, and parks it in `actor+0x44` at `L00040`. It then uses `R0011` for actor target or validates and uses `R0012` for a cell. Failed cell validation returns the item and clears `+0x68/+0x44`. These pointer writes precede the executor's progress gate and read no progress byte. The setters move the parked Spell into pending data and clear `+0x44`; they do not immediately replace active Spell `+0x64`.

The dispatcher calls `R0014` for resolved commanded members before the opcode arm. This helper returns without clearing an in-flight scroll cast when action is 0xd or 0xe and completion is zero (`L00041..L00042`). Its later matching-scroll branch deletes active Spell and clears `+0x64/+0x68`, not progress or victim. Its Item/Spell callees and ordinary item ownership remain separate from those direct writes.

A new scroll pointer can therefore coexist statically with a retained old active Spell. The cast application reads old `+0x64` and then current `+0x68`; its cleanup clears both (`L00043/L00044`). An ordinary physical strike skips this cast-only cleanup. A weapon-backed conversion in the shared body instead overwrites both fields at start/application (`L00045/L00046`, `L00047/L00048`). The physical application routine's weapon-proc arm also overwrites them and clears them after its Spell call (`L00049..L00050`). These paths exclude a general assertion that scroll pointer changes wait for recovery or survive every running strike.

**Confidence.** High for the named pointer writes, conditions and cleanup; they are whole local paths with raw-PE branch/displacement verification. No completeness claim is made about Item/Spell indirect dispatch or every writer of these fields.

**Unknown.** Which interleavings occur in native play, whether a pending scroll remains valid after such cleanup, exact inventory/resource loss, and per-spell behavior. No runtime or save corpus probe was run.

## Stand Ground member orders

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-349 | `R0015`, the order 3 arm's call for a human participant's member with no group-assigned target, stores `ord+0x08` once, as 0 or 8, and never stores `ord+0x09`, `ord+0x0a` or `ord+0x0c`. | High / Unknown | ● active | [EXP-0411](../experiments/EXP-0411-stand-ground/) |
| AI-350 | The per-tick executor `R0016` runs no arm for `ord+0x08` of 0, 3, 0xd or 0xe, re-reads `ord+0x0c` on every row 5 pass, and stores `ord+0x08` only as 0 in its own body. | High / Unknown | ● active | [EXP-0411](../experiments/EXP-0411-stand-ground/) |
| AI-351 | Both order 3 setters store `ord+0x08 = 0` for every member and never store `ord+0x09`, `ord+0x0a` or `ord+0x0c`; their other member stores are `ord+0x14`, `ord+0x38`, `ord+0x50`, `ord+0x60`, `actor+0x50`, `actor+0x54` and the post word. | High / Unknown | ● active | [EXP-0411](../experiments/EXP-0411-stand-ground/) |
| AI-352 | A member holding `ord+0x08` of 5 or 1 when a Stand Ground setter runs holds 0 afterwards, with `ord+0x0a` and `ord+0x0c` unchanged, and a step already in progress ends at the next cell centre. | High / Medium / Unknown | ● active (amended, partially retracted) | [EXP-0411](../experiments/EXP-0411-stand-ground/), [EXP-0414](../experiments/EXP-0414-cycle-orders/) |
| AI-353 | Under group order 3 a human participant's member keeps pathing to its last victim between two arm evaluations until a stand or refused route ends the order; the next evaluation reissues the attack for a victim in reach, else 0 or 8. | Medium / Unknown | ● active | [EXP-0411](../experiments/EXP-0411-stand-ground/) |

### AI-349

`R0015(session, actor)` (129 instructions, one stack argument besides the session, read whole) is entered from the dispatcher's inline order 3 arm at `L00051` when a member's `ord+0x20` is 0 and its owner's `Player+0x28` is 0 (`AI-STAND-076`, `UNIT-OWNER-009`). Five paths reach `L00052`, whose store at `L00053` writes `ord+0x08 = 0`: `actor+0x140 == 0` (`L00054`), word `actor+0x9a == 0` (`L00055`), word `actor+0x9a <= word actor+0xa0` (`L00056`), a null result from `R0017(actor+0x140, 6)` (`L00057`), and an ally scan whose best value is not below 1.0 (`L00058`; the double at `L00059` is `00 00 00 00 00 00 f0 3f` on both roots). Two paths cast and call `R0018(actor, target, spell)`: the self heal when word `actor+0x94 != actor+0x96` (`L00060`; `ord+0x60 = 1` at `L00061`, call at `L00062`) and the ally heal (`ord+0x60 = 1` at `L00063`, call at `L00064`).

The ally scan takes the cell from `R0019(actor+0x10)`, the candidates from `R0020(session+0xa50, cell, byte spell[+9])` (`L00065`) and the diplomacy filter `R0021(session, actor, list, 1)` (`L00066`, `AI-FILTER-001`). A candidate replaces the running best, seeded 2.0 (`[EBP-8] = 0x40000000`, `L00067`), when its integer quotient `word +0x94 / word +0x96` (`IDIV`, `L00068`) is lower and its word `+0x98` is nonzero. Only a candidate below its maximum health has quotient 0, and the first such candidate in list order stays best.

`R0018` (73 instructions, read whole) receives a nonzero target from both sites: the actor (`L00069`) or the candidate in frame local `-0x24`, which is assigned only where the best is updated (`L00070`) and which best below 1.0 implies. It stores `ord+0x08 = 8` (`L00071`), `ord+0x28 = target` (`L00072`), `ord+0x30 = spell` (`L00073`) and `ord+0x14`, first `actor+0x12c` (`L00074`) and then `spell[+9]` (`L00075`). Its `R0022` branch (`L00076`) needs a zero target and is not reached from these sites. The store table (every instruction of the read routines whose destination is a memory operand at displacement `+0x8`, `+0x9`, `+0xa`, `+0xc` or `+0x14`, stack slots removed, each hit classified by its base register) has one order block hit in `R0015` and three in `R0018`; no path stores `ord+0x09`, `ord+0x0a` or `ord+0x0c`. `callto:R0015` = **5 hits / 5 owners / 0 orphan**: `R0022` (`L00077`), `R0023` (`L00051`), `R0024` (`L00078`), `R0025` (`L00079`) and `R0026` (`L00080`).

**Confidence.** **High** for the write set: both routines are read whole and the store table is complete over their bodies. `R0019`, `R0017`, `R0020` and `R0021` are read whole and store nothing to an order block: `R0017` stores spell book `+0x18` once (`L00081`) and `R0021` loads the order block once, to read `ord+0x71` (`L00082`, `L00083`), while its store hits are inlined `CObList` removal on the candidate list (`L00084`..`L00085`).

**Unknown.** The list iterators `R0027`, `R0028` and `R0029` and the spell book accessors `R0030` and `R0031` are not read. What word `actor+0xa0` is and what spell id 6 is stay unresolved (`AI-FOLLOWHEAL-118` selects the same id). `R0020` calls `R0032`, `R0033`, `R0034` and `R0035`, and `R0021` calls `R0036`, which are not read; neither `R0020` nor these receive the order block through a `+0x158` load. Whether the order 1, 2, 4 and 5 arms reach this routine under the same conditions is not read here.

### AI-350

`R0016(actor)` (513 instructions, read whole; `callto:R0016` = **2 hits / 2 owners / 0 orphan**: `R0037` at `L00086` and `R0038` at `L00087`) reads the progress byte `ord+0x09` first. An idle actor (`ord+0x09 == 0`) whose group's AI-enable byte `grpAI+0x45` is 0, while `receiver+0xb388` is 0, parks `actor+0x54 = 0x1a` and returns before any row (`L00088`..`L00089`). Progress nonzero dispatches through the byte table at `L00090` into the six slot table at `L00091`: strike and cast counters (`L00092`, `L00093`) that count `ord+0x15` and clear the progress only when `ord+0x15 > 2` and `actor+0x136` is set (`L00094`..`L00095`); the step (`L00096`), which calls `R0039` and clears the progress and `actor+0x54` when `R0040` finds the actor at a cell centre (`L00097`); and three slots that park `actor+0x54 = 0x1a`. Progress 0 computes `ord+0x08 - 1` (`L00098`) and jumps through the 15 slot table at `L00099`: `ord+0x08 == 0` wraps above 14 (`L00100 JA`) and reaches the shared tail `L00010` with no arm, and `ord+0x08` of 3, 0xd and 0xe point at the same tail. Both tables read identically from the PE image and through `EnumRefs vt:`; the EN and RU executables are byte identical (SHA-256 `942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`).

Row 5 (`L00101`) reads reach `actor+0x12c` and `ord+0x0c` (`L00102`) and tests position with `R0041` (facing equal to the 8 way direction to the target and edge distance not above the reach, `AI-PURSUE-040`). In position it stores `ord+0x09 = 1`, `ord+0x15 = 0`, `actor+0x54 = 3` and `actor+0x5c = ord+0x0c` (`L00103`..`L00104`). Otherwise it re-reads `ord+0x0c` (`L00105`) and calls the approach `R0042(actor, target)` (`L00106`), which calls `R0043` with stop distance `ord+0x14` (a turn and stand when the target lies inside it, otherwise a route) and stores `ord+0x09 = 3` when the actor is not at a cell centre afterwards (`L00107`). Row 5 therefore takes one cell step per progress cycle toward the current position of `ord+0x0c` for as long as `ord+0x08` stays 5. Row 8 (`L00108`) does the same for `ord+0x28`, testing position at distance `ord+0x14`, and casts directly (`ord+0x09 = 2`, `actor+0x54 = 0xd`) when the target is the actor itself or in position.

The executor's own body stores `ord+0x08` seven times, all as 0 (`L00109`, `L00028`, `L00110`, `L00111`, `L00112`, `L00113`, `L00114`; the two stores of a zeroed scratch value, `L00028` and `L00112`, follow the zeroing at `L00115` and `L00116`, with no jump target between and no use of that scratch value in the two callees between), and never `ord+0x0a` or `ord+0x0c`. Two callees store it: `R0044`, a stand at the current cell that the step routine's arrival block calls (`L00117`), stores `ord+0x08 = 0`, `actor+0x50 = 0xc` and the post word (`L00118`, `L00119`, `L00120`); and the reacquisition `R0004` stores `ord+0x0c` with `ord+0x08 = 6` (`L00121`, `L00122`) or `ord+0x08` of 0 or 0xb (`L00123`, `L00124`, `L00125`; `AI-327`, `AI-328`). The executor epilogue calls it when `mover+0x98` is nonzero for an `actor+0x50` other than 1, 0xa and 0x17 (`L00010`..`L00011`). A Stand Ground member's `actor+0x50` is 0xc (`AI-AUTHOR-015`).

**Confidence.** **High** for the dispatch, the tables, the row 5 and row 8 reads and the store enumeration over the executor, `R0042`, `R0043`, `R0039`, `R0041`, `R0040`, `R0045` and `R0046`, all read whole. `R0047`, `R0048`, `R0049` and `R0050`, the callees of the step routine, load no `+0x158` order block and store nothing at those five displacements.

**Unknown.** Routines called on the row 5 path and not read: `R0051` and `R0036` (under `R0041`), `R0052`, `R0053`, `R0054`, `R0055` and `R0056` (under `R0043`), and `R0057`, `R0058` and the callees of `R0050`. `R0059` (705 instructions, called from the step routine's arrival block) is listed and not read whole; it loads no `+0x158` and its five hits store to the global message record at `L00126`. Whether any of these writes an order block field is not established.

### AI-351

`R0060(group)` (121 instructions, read whole) is the setter the player's Stand Ground command reaches (`L00127` in `R0061`, `AI-AUTHOR-015`, `AI-PANEL-053`) and the one `R0062` calls for the `StandGround = Yes` key (`L00128`): `callto:R0060` = **2 hits / 2 owners / 0 orphan**. Its first member walk calls `R0007` (`L00129`) and stores `ord+0x50 = 0` (`L00130`) and `ord+0x38 = 0` (`L00131`); then `grpAI+0x20 = 0` (`L00132`). The second walk, per member, sets `mover+1` to `mover+0` when they differ (`L00133`), calls `R0045` when word `mover+0x80` is nonzero and the actor stands at a cell centre on another cell (`L00134`), and stores `ord+0x14 = actor+0x12c` (`L00135`), `ord+0x60 = 0` (`L00136`), `mover+0x7c = 0` (`L00137`), `actor+0x54 = 0` (`L00138`), `ord+0x50 = 0` (`L00139`), `actor+0x50 = 0xc` (`L00140`), the post word `ord+0x00 = R0019(actor+0x10)` (`L00141`) and **`ord+0x08 = 0` (`L00142`)**. After the loop `grpAI+0x20 = 3` (`L00143`).

`R0063(group)` (155 instructions, read whole; `callto:R0063` = **6 hits / 5 owners / 0 orphan**: `R0064`, `R0065` twice, `R0066`, `R0067` and `R0003`) makes three member walks. The first (`L00144`..`L00145`) does the turn cancel and the conditional release and stores `ord+0x14 = actor+0x12c` (`L00146`), `ord+0x60 = 0`, `mover+0x7c = 0`, `actor+0x54 = 0`, `ord+0x50 = 0` and `ord+0x38 = 0`; then `grpAI+0x20 = 0` (`L00147`). The second calls `R0007` and stores `ord+0x50 = 0` and `ord+0x38 = 0`. The third (`L00148`..`L00149`) calls `R0007` and stores `ord+0x50 = 0`, `actor+0x50 = 0xc` (`L00150`), the post word (`L00151`) and **`ord+0x08 = 0` (`L00152`)**. The routine ends with `grpAI+0x20 = 3` (`L00153`). `AI-AUTHOR-015` places `L00154` on the single player load path.

`R0007` (36 instructions), `R0045` and its callee `R0046` (148 instructions, footprint bits in the plane grid only), `R0068`, `R0069`, `R0070` and `R0071` (the member walkers) are read whole. Neither setter nor any of them stores `ord+0x09`, `ord+0x0a` or `ord+0x0c`; the store table lists `ord+0x14` and `ord+0x08` in the two setters and `ord+0x14` in `R0007` (`L00155`).

**Confidence.** **High** for the write set: both setters and the routines listed above are read whole, and the store table is complete over them.

**Unknown.** `R0072` (under `R0068`) and `R0073` and `R0074` (under `R0070`) are not read; they are called from the member walkers, not from an order block access.

### AI-352

A member holding `ord+0x08 = 5` (pursue, victim in `ord+0x0c`) or `ord+0x08 = 1` (walk to the cell in `ord+0x0a`) when a Stand Ground setter runs holds `ord+0x08 = 0` afterwards (`AI-351`), with `ord+0x0c` and `ord+0x0a` unchanged. Its `ord+0x14` becomes `actor+0x12c`, `ord+0x60` and `ord+0x50` become 0, `actor+0x50` becomes 0xc, its post word becomes the current cell and `mover+1` takes `mover+0`, which cancels a pending turn.

The setters do not store `ord+0x09`, so they do not directly cancel a progress in flight. `R0016` (`AI-350`) runs the progress switch before it reads `ord+0x08`: a step (progress 3) calls `R0039` each actor tick until the actor stands at a cell centre, where the arrival block discards the route (`L00156`..`L00157`) and the executor clears the progress; a strike or cast counter (progress 1 or 2) counts `ord+0x15` and clears itself as `AI-350` states. With the progress clear and `ord+0x08 == 0`, no pending arm runs. A step already under way therefore reaches the next cell centre. A running strike or cast follows its retained cycle only when the common tail and other writers do not replace active fields (`AI-354`). The next evaluation of the arm reissues `ord+0x08 = 5` for a member whose scorer target is nonzero and otherwise stores 0 or 8 for a human participant's member (`AI-349`) and 0xb for an AI owner's (`L00158`); `AI-CLOCK-080` bounds the decision cadence.

**Confidence.** **High** for the values after the command, which are stores at cited addresses. **Medium** for the next tick behaviour, which composes the progress switch with the tail: the step and counter arms are read whole and no runtime observation exists.

**Unknown.** Native command/tick timing and untraced external writers. The bounded strike/cast body, selected command routes and common-tail replacement are covered by `AI-354` through `AI-357`.

**Amended.** The unconditional in-flight strike consequence is narrowed in `retracted.md`: retaining the old active fields requires the executor common tail and other writers not to replace them (`AI-354`). The immediate setter values and step-arrival mechanism stand.

### AI-353

For a human participant's unit under group order 3 the pursuit order and the arm run on different clocks. The arm runs once per full tick (`AI-CLOCK-080`). A member whose scorer target is nonzero gets `R0009`, which on its default path stores `ord+0x08 = 5`, `ord+0x0c = target` and `ord+0x14 = actor+0x12c` (`L00018`, `L00019`, `L00159`; `AI-337` names the exceptions); a member whose scorer target is 0 gets `R0015` (`AI-349`). Between two evaluations `R0016` runs on every actor tick (`AI-350`), and row 5 re-reads `ord+0x0c` and paths to that target's current position whenever the actor stands at a cell centre with no progress pending. A victim that leaves reach after the last evaluation is therefore followed cell by cell until the next evaluation, a stand or a refused route ends the order.

At the next evaluation the scorer of `AI-REACH-072` tests every candidate against reach again. A victim still in reach is issued again with `ord+0x08 = 5`. Otherwise `R0015` replaces `ord+0x08` with 0 (no cast) or 8 (heal cast) and leaves `ord+0x0c` at the old victim; with 0 the executor runs no arm (`AI-350`), the step in progress ends at the next cell centre and the member idles (`AI-352`). No routine read on this path measures the distance walked. Two routines end a row 5 pursuit between evaluations. The stand `R0044`, which the step routine's arrival block calls at `L00117` behind a cell-record test (`L00160`) and two `R0048` tests, stores `ord+0x08 = 0` (`L00118`). The executor epilogue turns a refused route (`mover+0x98`, stored at `L00161` and `L00162`; `AI-ROUTE-045`, `MOVE-073`) into `ord+0x08 = 0` and `R0004`, which stores `ord+0x08 = 6` for a victim in reach and, for a human participant's member with none, 0 (`AI-327`, `AI-328`); row 6 turns instead of pathing (`AI-PURSUE-040`). A victim that stays in reach at every evaluation is therefore followed until an evaluation, a stand or a refused route replaces the order. `AI-REACH-072` fixes what the arm assigns, and the executor's pathing between assignments is not covered by it. The idle turn `ord+0x08 = 0xb` is written only for an AI owner (`L00158`); a human participant's member idles with 0.

**Confidence.** **Medium**: the composition rests on cited stores and reads (`AI-349`, `AI-350`, `AI-351`, `AI-352`) and on the two clocks of `AI-CLOCK-080`; no runtime observation exists.

**Unknown.** `AI-CLOCK-080` leaves open a second call site of the group dispatcher (`L00163` in `R0075`) that is not gated by the full tick modulo, so the evaluation cadence there is not established. How far a member walks between two evaluations depends on its cell step time, which is not measured here. Whether an ally accepted by the heal scan can lie outside the position test of row 8 is not established: both use `spell[+9]`, the scan as a cell radius (`R0020`) and the test as an edge distance (`R0041`). How often the stand's cell-record test holds during a pursuit and how often a route is refused is not established (`MOVE-073` leaves the every-approach-occupied case Unknown), and the unread routines on the row 5 path (`AI-350`) are not excluded as further exits.

## Spatial group activity

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-ACTIVITY-324 | The regular activity rebuild derives Group AI+45 from spatially active members. | High / Unknown | ✔ promoted | [EXP-0367](../experiments/EXP-0367-group-activation/) |

### AI-ACTIVITY-324

DriverR0076 callsL00164 on world+92ef4 before Group dispatch when its receiver+b388 is zero; its later gate reads that same override or Group AI+45. CompleteL00164 andL00165 clear the grid, reset the byte for each represented actor whose Player+28 dword is nonzero, and mark coverage around actors whose control word is zero. Coverage uses map coordinates shifted right3, offsets -2..2 in both coarse axes, excluding the four corners. A second traversal increments the represented nonzero-control actor's Group byte when Group AI dword+48 is nonzero or its coarse cell is marked. It increments per actor, wraps the byte at 256, retains groups absent from the supplied list, and rebuilds rather than accumulating over calls. Local constructorR0077 ends with receiver+b388=0.

**Confidence.** **High** for the named complete rebuilds, static driver/constructor joins,35 original-instruction vectors per identical EN/RU image and a rejecting branch mutation. Supplied actor-list lifecycle, first native post-LOAD scheduling, other indirect writers and border-coordinate admission remain **Unknown**. No permanent Boolean interpretation or global activation radius for all paths is claimed.

## Formation owner selection

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-FORMOWNER-314 | Both Group movement formation reads use the explicit current Group owner. | High / Unknown | ✔ promoted | [EXP-0326](../experiments/EXP-0326-player-formation/) |
| AI-FORMCMD-315 | The client formation command transports an identifier; the simulation selects by saved Player+04, not archive position or +08. | High / Unknown | ✔ promoted | [EXP-0326](../experiments/EXP-0326-player-formation/) |
| AI-FORMTRIGGER-316 | Trigger formation targets bind through Player+08 and a resolved record pointer. | High / Unknown | ✔ promoted | [EXP-0326](../experiments/EXP-0326-player-formation/) |
| AI-FORMACTIVE-317 | The selected return-to-game and client-description paths do not define the active player as the first SAV record. | High / Medium / Unknown | ✔ promoted | [EXP-0326](../experiments/EXP-0326-player-formation/) |

### AI-FORMOWNER-314

`L00166..L00167` and `L00168..L00169` follow `Group+44 -> Player+30 -> byte+1f`. Their prefixes obtain the Group from an argument. Varying Group+44 between two Players with distinct bytes changes the read while their identifiers and the active client object remain fixed. With `SAV-PLAYERIDENT-830` and `SAV-GRPOWNER-561`, a locally loaded Group can therefore read another registered Player's formation even though its members have been stamped with the enclosing Player. A null owner has no fallback in these three-instruction reads.

**Confidence.** **High** for the original dereferences and four controlled read slices. The slices start after preceding member operations; malformed/null-owner reachability, original acceptance, scheduling and later ownership changes remain **Unknown**.

### AI-FORMCMD-315

`R0078` emits opcode46/subcode2 with `cmd+05 = low16((view+9b4)->CPlayer+04)` and the requested value at cmd+0e. `L00170 -> R0079 -> L00171` compares that signed word with each simulation Player's signed word+04; first match returns, miss returns null. `L00172..L00173` remaps 0/1/2 to 0/2/1, defaults to 2, and passes the resolved Player to `R0080`. Eight lookup controls and five client-to-dispatch vectors distinguish opposite +04/+08 values, reversed lists and duplicate-ID first-match behavior. The client CPlayer and serialized simulation Player are separate objects; their identifier, not their address, crosses this command.

**Confidence.** **High** for selected encoding, complete lookup, reached subcode and write. Transport is captured and fed into the selected dispatcher with ordered-list callbacks cut. Delivery, authorization, duplicate-ID acceptance and the first actual post-LOAD command remain **Unknown**.

### AI-FORMTRIGGER-316

Resume linker instructions `L00174..L00175` insert each Player into a temporary map keyed by dword+08. Reference type3's arm `L00176` passes the authored value to `L00177`; the first Player result goes to compiled rec+38 at `L00178`. `L00179` and the dispatcher copy the 72-byte record; instant7 passes its resolved +38 Player and raw +08 parameter to `R0080` at `L00180`. This write does not reselect the active client or compare the command's +04 identifier. Two resolved-target vectors change only the selected Player's byte.

**Confidence.** **High** for the named map key, record transport, arm and setter operands. Temporary-map collision behavior, every rebinding path and full LOAD lifecycle remain **Unknown**. Vectors execute the final arm with a supplied resolved pointer, not the whole linker.

### AI-FORMACTIVE-317

Existing-Player search `L00181..L00182` compares the incoming CString with Player+18 in live list order, retaining the first accepted name match at `L00183`; credential rejection additionally depends on world+0c. The retained Player's +04 goes into the selected opcode96 description at `L00184..L00185`, before `L00186 -> R0081` sends other Players. The client description arm assigns its described CPlayer to view+9b4 only when current array size view+9a8 equals 1 (`L00187..L00188`), then inserts it under its own +04 (`L00189..L00190`). Name-match vectors retain the same object under reversed order; array-size controls separately discriminate the active-pointer gate.

**Confidence.** **High** for bounded predicates and normal-return controls; **Medium** for their composition after mission LOAD. Incoming join-name source on every LOAD route, initial client-array state, packet delivery order, fresh-player path, credentials and complete multiplayer behavior remain **Unknown**. This conditional account does not guarantee that every save resumes a specified Player.

## Structure use, spell book and quick slot controls

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-STRUCTUSE-306 | Order0x24 schedules approach and one local actor use action. | High / Unknown | ✔ promoted (amended) | [EXP-0325](../experiments/EXP-0325-structure-use/) |
| AI-SPELLIDENT-286 | Valid book cell `i` in 0..23 is a signed controller index. | High | ✔ promoted | [EXP-0294](../experiments/EXP-0294-quick-spell-identity/) |
| AI-SPELLPOP-287 | Selection recomputation `R0082` accepts each selected node's object only when object+7c is nonnull, increments view+140, and ORs object+18 into availability view+148 before session/class/type predicates. | High | ✔ promoted | [EXP-0294](../experiments/EXP-0294-quick-spell-identity/) |
| AI-SPELLCAP-288 | For session+3dc equal1 or carrying bit2, every accepted selected object is tested: runtime class-name `CUnit`, object+20 in17h/18h, nonzero object+18 OR200h into view+144. | High | ✔ promoted | [EXP-0294](../experiments/EXP-0294-quick-spell-identity/) |
| AI-SPELLGUARD-289 | Mouse book selection writes current+60 only after a nonnegative grid hit and its bit in view+148. | High | ✔ promoted | [EXP-0294](../experiments/EXP-0294-quick-spell-identity/) |
| AI-SPELLITEM-290 | Active controller item+74 switches actor/cell producers to opcodes25h/26h and writes the word at controller+78 into command+10. | High | ✔ promoted | [EXP-0294](../experiments/EXP-0294-quick-spell-identity/) |
| AI-QUICKASSIGN-278 | F5–F8 reach map branches `L00191/L00192/L00193/L00194`, sending message `0x417`, slots 0–3 and the Ctrl latch to the spellbook controller at `campaign+0xec`. | High | ● active | [EXP-0273](../experiments/EXP-0273-quick-slots/) |
| AI-QUICKINVOKE-279 | Plain quick-key message `0x417` copies a nonnegative stored index into controller `+0x60`; a negative value leaves it unchanged. | High | ● active | [EXP-0273](../experiments/EXP-0273-quick-slots/) |
| AI-QUICKOWNER-280 | The reached quick-slot owner is the spellbook controller constructed by `R0083` and retained at `campaign+0xec`, not an actor/group-relative array. | High / Medium | ● active | [EXP-0273](../experiments/EXP-0273-quick-slots/) |
| AI-QUICKSAVE-281 | Original writer `R0084` emits controller offsets `+0x64..+0x70` in order as `SpellBook/Shortcuts`, kind 6 and 16 payload bytes: F5, F6, F7, F8 signed zero-based spell indices, with `-1` unbound. | High | ● active | [EXP-0273](../experiments/EXP-0273-quick-slots/) |

### AI-STRUCTUSE-306

Complete `R0085` sets actor state `+50=15`, clears action `+54`, writes Building pointer to order `+68`, and sets destination from Building Position. Width byte `+60` selects range/packed-cell offset: widths1/2 use1/+0; widths3/4 use2/+0x101; other widths use3/+0x202. State15 copies to order15. Its pending arm `L00195..L00010` calls `R0086`; failure schedules movement through `R0087` and action1, success sets action15 unless already15, in which case it clears/completes through `R0007` and `R0088`. The complete admission helper requires the byte returned by unexpanded `R0089` to equal Mover byte0, then unsigned Chebyshev cell distance no greater than the supplied range. Actor action15 maps through two raw tables to `L00196`.

**Confidence.** High for named local stores, tables and branches, calibrated against the original PE. Unknown for the nested orientation helper, complete scheduling, interruption and first dispatch after LOAD. No claim of instant completion or exact visible delay.

**Amended.** The Unknown for the nested orientation helper closes for `R0089` alone: `AI-446` states its law (the `AI-444` branch tree on whole-cell differences from the anchor cell, sub-cell and size ignored) and that the point gate and the walk's stop arm use it. Complete scheduling, interruption and first dispatch after LOAD stay Unknown.

### AI-SPELLIDENT-286

Map targeting passes `i+1`; book producers write that ordinal byte to command+10 for opcodes1e/1f. Dispatcher `L00197` reads the byte at `L00198+4*ordinal`, rewrites command+10, then calls `R0017(actor+140,id)` and parks the result at actor+44. The statistics reader uses the same storage via `L00199+4*i`. Cell5 therefore carries command6 but resolves spell23. The 24 IDs are `1 2 3 4 5 23 24 16 15 14 13 12 6 7 8 9 10 25 26 22 21 20 19 18`; the preceding dword is 0. This closes AI-CMD-032's unread table contents without changing MAGIC-ICON-024's mapping.

**Confidence.** High for the bounded representation chain: byte stores, in-place table rewrite, table alias and actor lookup are independently decoded. Malformed indices, stale actor+44 and eventual effect success remain Unknown.

### AI-SPELLPOP-287

It is a union, not intersection or primary-only availability. The first accepted object becomes view+138. Book-command producers `R0090/R0091` separately include only selected objects with +7c nonnull and bit `1<<(ordinal-1)` in +18; item context bypasses the bit. Append helper `R0092` caps the list at 253.

**Confidence.** High for the complete local population loops and guards. Upstream selection restrictions, snapshot+18 refresh, arbitrary mixed-population reachability and delivery interleaving remain Unknown; no global health/mana absence claim.

### AI-SPELLCAP-288

The predicate uses the current loop object, not only the primary; one later eligible object can enable Cast. Other session states assign summary8. The primary owner alone supplies bit4. `R0093` returns 0 for bit4, elseEFh plus10h for200h. The panel separately blanks on zero accepted count or summary24h. Action5/9 needs the Cast mode bit; item action0Ah bypasses the mode refusal. This partially corrects AI-PANEL-061's primary-only200h clause.

**Confidence.** High for the reached per-object accumulator and distinct gates. Runtime class taxonomy beyond the literal tests and reachability of every synthetic mixed state remain Unknown.

### AI-SPELLGUARD-289

The plain shortcut path writes any nonnegative binding first; unavailable does not imply impossible as current. AI-QUICKINVOKE-279's later availability/arming guard stands. In the reached map Cast branch, item+74 takes precedence over current+60 and a negative result emits no cast. For book selection `R0094` readsL00200 by cell:14 of 24 entries are nonzero. A unit-like hit plus nonzero flag selects actor production; other hits select cell production; without a hit, a nonzero flag emits no order.

**Confidence.** High for the bounded selection and target-production branches, not all hover/target validity or runtime feedback. Malformed stored indices, global current-state writers and every input focus state remain Unknown.

### AI-SPELLITEM-290

This is an inventory item position, consumed byR0095, not a book ordinal. The1e/1f table-rewrite prologue is skipped. The item arm requires a retrieved item and first effect kind29h, constructs Spell from effect byte+40, then calls the shared actor/cell setters. The target predicate also uses another table,L00201, indexed by item+74.

**Confidence.** High for the reached operand-width/opcode fork and item constructor chain. Item UI snapshot synchronization, all invalid-item states and the broader refund/use lifecycle are not expanded.

### AI-QUICKASSIGN-278

Its handler `R0096` assigns on nonzero Ctrl: current signed spell index `+0x60` when not `-1`, otherwise a nonnegative mouse-grid hit from `L00202`. It stores to `+0x64+4*slot` and empties every other equal slot to `-1`; this moves a duplicate binding, not a swap. The hover assignment arm has no spell-availability test and does not set `+0x60`. Ctrl still passes through the key handler's later arming checks.

**Confidence.** High for the exact physical-key table and reached local assignment branches on the identical EN/RU executables. Original GUI reachability under every focus/modal state was not witnessed.

### AI-QUICKINVOKE-279

With the spellbook closed, the map handler tests the slot through `L00203`: exactly `-1` is empty, otherwise the stored-index bit must exist in selected-population mask `view+0x148`. An available slot calls `R0097(9)`, which maps to mode 5 and requires the Cast-capability bit before storing `view+0x99c`. Case 9 neither opens the book nor directly emits a cast order. With the book open it only selects. Empty `-1` leaves the local mode/current spell unchanged; an unavailable populated slot still changes current spell but requests no mode. Later map targeting adds one to the selected index for the cast producer.

**Confidence.** High for the reached local dataflow and guards. Generic child forwarding, selected-item override, malformed indices and actual runtime casting/feedback are not globally closed.

### AI-QUICKOWNER-280

Construction sets its four dwords `+0x64/+0x68/+0x6c/+0x70` to `-1`, separately from current spell `+0x60` and item override `+0x74`. Message `0x411` and controller right-down/up clear only `+0x60`; `L00204` clears the separate item override. Selected-population recomputation `R0082` rebuilds availability mask `view+0x148` without directly writing bindings; book open/close attaches/detaches the same controller.

**Confidence.** High for owner construction and these direct operations. Medium for selection carry across all callbacks; no-load mission entry, new-campaign reset and network/session interleaving remain Unknown. This is not an exclusive-writer census.

### AI-QUICKSAVE-281

It separately emits current spell `+0x60` as `SpellBook/Pressed` and book visibility as `IsOpen`. Restore paths `R0098` and `R0099` read the array and copy four dwords to the same controller, without per-index validation; `Pressed` is separately restored. This identifies the four meanings left open by `SAV-ORIGFAULT-335`, not full malformed-save compatibility. All 27 preserved-root saves (EN 23, RU 4) have four `-1` slots and `Pressed=-1`; they contain no populated round-trip witness.

**Confidence.** High for the bounded original save/load programme and exact corpus measurements. Populated-slot original reload behaviour and later invalid-value handling are unwitnessed; broader session persistence remains Unknown.

## Explicit Retreat

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-RETREAT-270 | Explicit Retreat admission has separate panel, producer and dispatcher gates. | High | ● active | [EXP-0272](../experiments/EXP-0272-immediate-retreat/) |
| AI-RETREAT-271 | Explicit Retreat installs a persistent actor state, not a single destination or the automatic threshold policy. | High | ● active | [EXP-0272](../experiments/EXP-0272-immediate-retreat/) |
| AI-RETREAT-272 | Retreat does not unconditionally interrupt an in-progress step, attack or cast. | High | ● active | [EXP-0272](../experiments/EXP-0272-immediate-retreat/) |
| AI-RETREAT-273 | The explicit Retreat state recomputes an away order through `R0100`; reaching one flee cell is not its own state terminator. | High / Medium | ● active | [EXP-0272](../experiments/EXP-0272-immediate-retreat/) |
| AI-RETREAT-274 | Automatic withdrawal shares execution helpers with explicit Retreat but not its admission or state installation; the two automatic helpers are not identical. | High | ● active | [EXP-0272](../experiments/EXP-0272-immediate-retreat/) |
| AI-RETREAT-275 | Later admitted commands can replace explicit Retreat, and zero HP is not an extra Retreat-refusal check. | High / Unknown | ● active | [EXP-0272](../experiments/EXP-0272-immediate-retreat/) |

### AI-RETREAT-270

Panel left-down `R0101` requires no selected item at `session+0x3cc`, active panel, an enabled hit cell and a changed overlay cell. `R0082` disables the panel when the counted inventory-bearing selection is empty or summary has `0x24` (foreign first owner or structure); `R0093` supplies Retreat bit `0x80` without a spell-capability requirement. `R0097(8)` calls `R0102` immediately, then clears the armed mode. The producer appends selected entries with nonnull `actor+0x7c`; its append helper caps the list at 253. The dispatcher requires the AI manager, nonempty list and a resolvable first actor. An absent first actor aborts the entire command; absent later actors are skipped. Resolution uses the requested player's actor list and rejects action `actor+0x54 == 0x10`, not all nonpositive HP. The Retreat arm has no own health-threshold, hostile-presence, movement or casting admission test.

**Confidence.** High for these concrete code paths and gates. Physical keyboard focus, network delivery and selection changes between producer and dispatch were not runtime-witnessed.

### AI-RETREAT-271

`L00205` calls `R0103` on the newly grouped command members. Its reset passes call `R0007`, align mover desired facing to current facing, conditionally clear a noncurrent reserved cell only at a cell centre, copy reach to `ord+0x14`, and zero `ord+0x60`, `mover+0x7c`, `actor+0x54`, `ord+0x50`, and `ord+0x38`. It writes `actor+0x50 = 0x16` at `L00206`, pending order `ord+8 = 0`, and group order `grpAI+0x20 = 0`. Neither `R0103` nor `R0007` clears order-progress `ord+9`, attack target `actor+0x5c`, or the queued cast target/spell slots. The dispatcher separately calls the conditional spell-cancellation helper `R0014`; universal spell-object destruction is not established.

**Confidence.** High for writes and nonwrites of the two complete setter/reset bodies. No claim that every referenced object survives the separate cancellation/lifecycle helpers.

### AI-RETREAT-272

After the setter leaves `ord+9` unchanged, `R0016` consumes nonzero progress before the pending-order switch: progress 1 restores action 3; progress 2 restores action `0xd` or `0xe` from `ord+0x5c`; both increment `ord+0x15` and clear progress only when it exceeds 2 and `actor+0x136` is nonzero. Progress 3 calls `R0039`, keeps movement action 1 and clears progress on the cell-centre predicate. Progress 4 retains the status hold until effect bit `0x100000` clears. Thus a new Retreat state can coexist with completion of an older action.

**Confidence.** High for static state-machine precedence, including both original tables and branch targets. Exact visible final-hit/cast timing, emitted effects and animation-phase outcomes are Unknown without an original runtime witness.

### AI-RETREAT-273

Group order 0 calls `R0008`, whose state `0x16` calls that helper at `L00207`. It collects the occupancy square about the actor with radius `mover+8`, applies the directional diplomacy/invisibility filter, and gates on a nonzero byte count of positive-HP entries. It then averages fine coordinates over the whole surviving list, divides by the low byte of the list count, calls `R0104` with distance 3, and writes ordinary pending move 1 plus its cell. A zero positive-HP count calls ordinary acquisition `R0022`, which can install a pursuit or idle/autoheal order. Neither this helper nor the direct actor-state arm clears state `0x16` or has a timer/arrival test. Route failure reaches the ordinary `R0004` fallback.

**Confidence.** High for the bounded local control flow and arithmetic. Medium for sustained in-game persistence: whole-session event interleaving, all transitive lifecycle effects, dense-list byte wrap and runtime corpse occupancy are not closed.

### AI-RETREAT-274

The group-tail tests `R0105` then `R0106` compare signed current HP with `ord+0x44` and `ord+0x40`, using radius `mover+8` and literal 2. They call `R0100` and `R0107` respectively, without installing state `0x16`. In `R0100`, `L00208` counts positive HP only for the branch at `L00209`; the averaging loop starts from the entire filtered list and its divisor is `listCount & 0xff`. `R0107` has only a nonempty-list gate and no positive-HP scan. This retracts `AI-WITHDRAW-028`'s living-only centroid and shared positive-HP gate, not its distance-3 geometry or absence of a passability test in the picker.

**Confidence.** High for the two whole helper bodies, the tail and the discriminating branch/loop dataflow. Whether a mixed or all-dead hostile list reaches either helper during ordinary play remains Unknown.

### AI-RETREAT-275

Move `R0108` finishes with group order 4, so group dispatch no longer uses the per-actor state-`0x16` arm. Target attack `R0006` installs actor state 3 or fallback `0xc`; cast-at-actor `R0011` installs `0xd` or `0xc`; cast-at-cell `R0012` installs `0xe` or `0xc`. These are concrete replacement paths, not a claim about every opcode. The resolver rejects completed-death action `0x10`; an earlier death phase with HP at or below zero can still resolve. The unit tick `R0037` checks HP before order execution, zeros state/action in its nonpositive branch and does not call `R0016` there. Retreat therefore supplies no resurrection path in the traced tick.

**Confidence.** High for the named replacement and death paths. Unknown whether ordinary selection/queue timing can deliver Retreat in every intermediate death phase, or what a missing target/new command does on untraced opcodes.

## Prismatic Spray candidate lists

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-SPRAY-266 | Prismatic Spray ranks the shared group candidate lists in their existing order, rather than performing a radius search. | High | ● active | [EXP-0243](../experiments/EXP-0243-prismatic-victims/) |
| AI-SPRAY-267 | Only Prismatic Spray's secondaries pass the group visibility, first-member diplomacy/invisibility and corpse fallback gates; its primary bypasses them. | High / Unknown | ● active | [EXP-0243](../experiments/EXP-0243-prismatic-victims/) |

### AI-SPRAY-266

`R0109` calls `R0110(session,caster+0x70)`, then traverses scratch list A head to tail followed by list B head to tail. Its stack pointer array holds 100 entries, while scores go to `[session+0xd74+4*i]`, whose capacity is Unknown; neither copy loop checks a bound. It scores A by `((edgeDistance<<8)+turnCost)&0xffff` and B by `((((edgeDistance<<8)+turnCost)&0xffff)<<8)`, then repeatedly selects the strict minimum. Strict `<` preserves source order on equal stored scores. The primary target is not admitted through either list: the selector appends it first and skips its duplicate among winners. It is therefore possible for a primary that fails the secondary filters to remain the first victim.

**Confidence.** High. The direct caller, complete selector and complete builder already carried by `AI-GROUPSEE-068` distinguish this from an unordered set and a spatial-radius collector. The score-scratch capacity is explicitly outside the claim.

### AI-SPRAY-267

`R0110` clears sight once, stamps every group member, walks the global actor list head to tail, applies `R0021(session,firstMember,A,0)`, and moves health-below-one entries A→B. When A is empty, the builder moves all B back into A and those corpses use the A score. Otherwise B remains a distinct source list and uses `((((edgeDistance<<8)+turnCost)&0xffff)<<8)`. Selection first requires that score below 65530; append later requires the saved threshold below 65000. The bytes therefore do not establish a categorical living-excludes-corpse rule: a sufficiently low B score can pass, and this round did not establish whether the required edge-distance geometry is reachable. `R0109` separately writes the primary target's alarm and calls the directional hostility flip before invoking the builder. Thus the old statement that the Prismatic collector consults no diplomacy mistook absence of a direct matrix load for absence of the builder call.

**Confidence.** High for the gate order, primary bypass, B formula and sentinel tests: the selector and builder are read whole. Unknown whether a B corpse can attain an appendable score beside living A in a valid runtime layout. The candidate is visible to any group member but diplomacy is decided by the first member, as `AI-GROUPSEE-068` already establishes.

## Mission input contracts

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-INPUT-121 | The map's physical left contract is inactive-down-to-start, active-up-to-act. | High | ● active | [EXP-0190](../experiments/EXP-0190-combat-controls/) |
| AI-SELECT-122 | Mission selection has four exact forms, with a gate on both Shift forms. | High / Medium | ● active | [EXP-0190](../experiments/EXP-0190-combat-controls/) |
| AI-PANEL-123 | The command-panel/key contract is: Attack mode 1 → target `0x19`, ground fallback move `0x16`; Move mode 2 → `0x16`; Guard immediate `0x17`; Defend mode 4 → target `0x1b`, ground no-op; Cast mode 5 → spell `0x1e/0x1f` or item `0x25/0x26`… | High | ● active | [EXP-0190](../experiments/EXP-0190-combat-controls/) |
| AI-MINIMAP-124 | The 160x158 minimap acts on left DOWN, not up: default moves the camera `0x406`; move emits `0x16`; attack emits `0x19` on any non-zero cell object id and otherwise `0x1a`; defend emits `0x1b` only with an id… | High | ● active | [EXP-0190](../experiments/EXP-0190-combat-controls/) |
| AI-KEY-125 | The complete decoded default mission key surface is committed in `keyboard.tsv`, except the four rows amended below. | High / Unknown | ● active (amended) | [EXP-0190](../experiments/EXP-0190-combat-controls/), [EXP-0446](../experiments/EXP-0446-toggle-notice/) |
| AI-CURSOR-126 | ROM1 contains a Patrol cursor and order but no shipped default gesture that leaves Patrol armed. | High | ● active | [EXP-0190](../experiments/EXP-0190-combat-controls/) |
| AI-INPUT-127 | The map's physical right contract is capture-to-pan, up-to-cancel. | High / Unknown | ● active | [EXP-0190](../experiments/EXP-0190-combat-controls/) |

### AI-INPUT-121

`WM_LBUTTONDOWN` clips to the map and stores a marquee origin only while global `[L00210]` is zero; a repeated down while it is nonzero is a consumed no-op, and double-click aliases this same gate. `WM_LBUTTONUP` closes the marquee, updates the hit and reaches selection or the cursor command only while that global is nonzero. It does **not** call `ClipCursor(NULL)`: `R0111` passes the session rectangle at `session+0x6c0`. An up without an active marquee skips clipping, hit update, selection and cursor action, but the later selected-item placement arm still runs when `session+0x3cc` is nonzero, emitting `0x23` for the `item+6 == 0xffff` branch or `0x22/0x32` through the alternate transfer builder. The threshold is `screenW*10/640` = 10/12/16 pixels. A rectangle strictly beyond it goes to selection; it is not discarded. At equality the result depends on whether the caller entered selection for another reason, because the outer test is `>` while the selection routine's click test is `<`

**Confidence.** High (the physical vtable slots, both `[L00210]` gates, post-gate selected-item branches/opcodes, the bytes that pass the address at offset `0x6c0` to `ClipCursor` and both threshold comparisons are read end to end and raw-asserted)

### AI-SELECT-122

Plain click replaces with the topmost intersecting non-structure regardless of owner, or a structure whose class `+0x64` is zero. Shift-click and Shift-rectangle toggle an owned qualifier only when the **old** selection summary has bit `0x04` clear; an owned candidate is a no-op when that bit is set. Therefore, after a plain click selects a foreign actor, Shift-click or Shift-rectangle over an owned actor does not add or toggle it. Ground and foreign candidates also preserve selection. Plain rectangle selects every owned non-structure whose drawable rectangle overlaps by strictly more than half, but preserves the old selection if none qualifies. Alt after a one-object click expands to its stored group without centring; E selects every owned exact-name `CUnit`; plain/Shift digit selects or augments a group; Ctrl+digit assigns; Alt+digit selects and centres with message `0x406`, and Ctrl wins over Alt

**Confidence.** High for the four selection forms, old-summary/ownership gates and modifier precedence (selection and key routines read whole, and the `view+0x144 & 4` gate is asserted from raw bytes); Medium for exact-name `CUnit` as a behavioural category because it is a string compare rather than a runtime-class test

### AI-PANEL-123

The command-panel/key contract is: Attack mode 1 → target `0x19`, ground fallback move `0x16`; Move mode 2 → `0x16`; Guard immediate `0x17`; Defend mode 4 → target `0x1b`, ground no-op; Cast mode 5 → spell `0x1e/0x1f` or item `0x25/0x26`; Swarm mode 6 → `0x1a`; Stand Ground immediate `0x18`; Retreat immediate `0x14`. An armed cursor is one-shot and clears after a consuming click. Repeated keydown reissues Guard, Stand Ground and Retreat; repeated A/M/D/S only reasserts a mode; C is suppressed after its UI opens. The eight-cell panel is inactive for no selection, foreign ownership or the structure flag; cast alone needs the spell-capable flag. **It is not the only item-order producer:** character-panel `WM_LBUTTONUP` and nested inventory-grid `WM_LBUTTONDBLCLK` can pass an admitted item through vslot `+0x7c`, store its slot/effect context, and call `R0097(0x0a)`, which remaps to Cast mode 5; the next map down/up emits item `0x25/0x26`.

Inventory-grid `WM_LBUTTONUP` is a separate transfer path: live `grid+0x84` and selected `session+0x3cc` call grid vslot `+0xa4`; its destination code is 2, and `L00211` emits `0x22/0x32` even when hit vslot `+0x88` returns `-1`, using merge/insert-before-gold/append fallback. Its sibling handlers scroll on left-down, build item/money drag state on left-button move, and consume every right-button edge as a no-op; none emits another order. The grid double-click's gold-sentinel arm opens Drop Gold; action button-up or Enter emits `0x23` with the resolver-returned amount and current cell, while cancel button-up or Esc emits none

**Confidence.** High (panel/key tables, the complete inventory-grid physical mouse vtable, three physical inventory-item producer paths, modal controls and raw item/gold gates, every direct caller of all fifteen producer targets with a physical or explicit non-mission classification, order builders and selection-summary predicates are read and asserted)

### AI-MINIMAP-124

The 160x158 minimap acts on left DOWN, not up: default moves the camera `0x406`; move emits `0x16`; attack emits `0x19` on any non-zero cell object id and otherwise `0x1a`; defend emits `0x1b` only with an id; cast discards its hit-test result and emits nothing; patrol emits `0x1d`. Left-drag repeats the current small-cursor action per delivered move. Right down and right-drag move the camera; there is no capture. When a mouse-move arrives with both `MK_LBUTTON|MK_RBUTTON`, the handler tests the left bit first, executes the left action and returns, so no right-camera pan occurs. Left up only clears a special cursor; both double-clicks and right up are no-ops. The object's identity is established from construction, rectangle and overview draw, not the `s*` filenames

**Confidence.** High (constructor/store, draw, physical vtable and handlers read; the left-test/early-return bytes and all producer opcodes re-measured on both roots)
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### AI-KEY-125

It includes Esc menus; F1 help; F2 save; F3 load or Diplomacy by phase; F5..F8 quick spells; Pause; arrows; A/M/G/D/C/S/T/R/E; Space, B/Q, I/backtick; digits and numpad digits with Shift/Ctrl/Alt; numpad speed; Ctrl+F/H/L/N/O/U/W settings; Enter text entry; focused-character-panel Tab message `0x412`; and Alt+S screenshot. F4 and F9 are explicit no-ops; no wheel route exists. Windows repeat is not filtered. Open text entry consumes map shortcuts; focused popup children take precedence; focus loss clears all three modifier latches. Ctrl+numpad plus sets `campaign+0x40c=1`, selecting the unpaced/max-speed owner loop with one sub-tick per idle callback and no deadline; minus clears it, resets phase/epoch and restores paced mode (`SESS-CLOCK-005`, `SESS-IDLE-007`). F12, Backspace's clear, character-panel `0x412` and Alt+B..Y except S's outbound `0x46` record have exact mechanisms but Unknown gameplay meaning

**Confidence.** High for the event/key/action mapping and repeat/focus predicates (frame message map plus frame/map/character-panel handlers and existing clock provenance read); Unknown only for the four mechanisms explicitly named

**Amended.** `keyboard.tsv` rows for Ctrl+F, Ctrl+L, Ctrl+O and Ctrl+W label the settings the wrong way round: W is retreat mode, F formation, L flying damage and O smoothing, with the line each posts in `MENU-057`. The H, N and U rows and the key set stand. `MISSION-MSGPOST-058` and `ANIM-NUM-020` already carried the correct map; the cause of the row error is not identified. `MENU-060`, `MENU-061`, `MENU-062` and `AI-378` since read three of the four named mechanisms: F12 toggles the fps readout, Backspace calls three routines with the arguments of `SetSize(0, -1)` on the map message line's arrays, and Alt+B..Y except S send a record whose receiver arm is the debug console behind a privilege byte. Character-panel `0x412` stays Unknown.

**Amended.** "No wheel route exists" is narrowed to "no wheel route found", Medium (`MENU-078`). Population: the game message maps, the frame procedure and the root dispatcher's message set. A plain 4-byte search finds seven `0x20a` hits; the library scroll-view message-map entry is reached by no game object, and two library sends of `0x20a` at `L00212` and `L00213` were not traced. F1 inside help scrolls with the arrow and page keys while the text control has focus (`MENU-078`).

### AI-CURSOR-126

Mode 8 maps to the patrol cursor; map and minimap consumers emit `0x1d`. The executable direct-call census for `R0097` contains 14 sites, including item route `L00214`; both its character-panel left-up and inventory-grid double-click entrances pass `0x0a`, which remaps to Cast mode 5, not Patrol. The corrected whole-image census begins with all 77 direct calls to the common enqueue routine in 64 producer functions, enumerates every direct caller of every producer, and uses physical traces to select the fifteen order/item targets reached by default input. It includes inventory-grid left-up reaching transfer call `L00211` through vslot `+0xa4` and the second `R0112` caller in Drop Gold, but no new arm routine.

All 42 direct call sites under the selected targets have a physical predecessor, inherit one through the shared command dispatcher, or are explicitly non-mission. The panel/custom/key routes leave R as immediate Retreat `0x14` with mode reset. The capability mask's `0xef` admits internal mode 8 but does not make it reachable. **G2:** adding a binding is an authored executable/UI change; it changes no shipped data file, while changing the cursor sheet changes `graphics.res` bytes

**Confidence.** High for default unreachability within the global enqueue/direct-caller plus physical vtable/pointer/construction census; a dynamically manufactured engine custom message is outside the shipped default contract

### AI-INPUT-127

Right down captures and stores the origin. With `MK_RBUTTON`, each non-zero signed-truncated `(current-last)/8` cell delta sends camera message `0x406` and marks a drag. Right up releases capture; a marked drag performs no cancel, while a click sends `0x405`, cancelling an armed mode or deselecting all when no mode is armed. Right double-click is a no-op. No mission handler exists for capture loss or cancel mode, so the image does not establish the external-loss outcome

**Confidence.** High for all in-image event edges and branches; Unknown for Windows-driven external capture loss, explicitly outside the handled set

## Target acquisition, the actor list and the diplomacy matrix

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-FILTER-001 | `R0021` is not an acquisition routine — it is a diplomacy predicate over a candidate list, and the relation it keeps is a *parameter*. | High / Medium | ● active | [EXP-0074](../experiments/EXP-0074-target-acquisition/) |
| AI-ACQUIRE-002 | The standing acquisition rule is `R0022`: everything the unit can see, filtered to enemies, nearest first and cheapest turn to break a tie — then discarded unless it is within reach. | High / Medium | ● active | [EXP-0074](../experiments/EXP-0074-target-acquisition/) |
| AI-LIST-003 | `world+0xa4554` is the global on-map actor list, walked linearly — there is no spatial index anywhere on the acquisition path. | High / Medium | ● active | [EXP-0074](../experiments/EXP-0074-target-acquisition/) |
| AI-DIPLO-004 | The relation a unit consults is a 50×50 byte matrix at `session+0xa9c4`, and it is empty until something fills it. | High / Medium | ● active (amended, superseded) | [EXP-0074](../experiments/EXP-0074-target-acquisition/) |
| AI-DIPLO-005 | A scenario authors the matrix, one row per player, and the `.alm` loader is the consumer `ALM-GRP-041` could not find. | High / Medium | ● active | [EXP-0074](../experiments/EXP-0074-target-acquisition/) |
| AI-SIGHT-006 | `actor+0xa5` is the sight radius and the guard leash — it is never an acquisition radius. | High | ● active (amended, partially retracted, superseded, contested) | [EXP-0074](../experiments/EXP-0074-target-acquisition/), **[EXP-0082](../experiments/EXP-0082-ai-behaviour/)**, **[EXP-0088](../experiments/EXP-0088-fog-of-war/)**, **[EXP-0120](../experiments/EXP-0120-sight-writers/)**, [EXP-0322](../experiments/EXP-0322-group-order-producers/) |
| AI-GUARD-007 | What closes the distance is a pursuit order, not a wider scan — and two clauses stop a unit engaging at all. | High | ● active | [EXP-0074](../experiments/EXP-0074-target-acquisition/) |

### AI-FILTER-001

the callee clears twelve bytes of stack arguments, so the signature is `__thiscall(this = session, me, list, wantBit)`: for each element it reads the signed diplomacy byte (`L00215`) at `session + 0xa9c4 + 50·myPlayerIdx + hisPlayerIdx`, keeps its low bit, and **unlinks the element when the bit equals `wantBit`** (`L00216`/`L00217`), by an inlined `CObList::RemoveAt` at `L00218`. An element the diplomacy test would keep is unlinked anyway if it is invisible (`cand+0x144 & 0x8000`, `L00219`/`L00220`) unless some member of the decider's group list `me+0x70` is within that member's `order+0x71` (`L00083`). **Callers: 15 hits, 15 distinct owners, 0 in orphan/undisassembled code** (`EnumRefs callto:R0021`, repaired function table).

Two pass `wantBit = 1` — `R0015` (`L00066`) and `R0113` (`L00221`) — and both then minimise `cand+0x94 / cand+0x96`, the health fraction, and cast: they are the heal routines, which is what fixes the polarity. **Three producers fill the list**, not one: `R0020(cell, r)` builds the occupants of the `(2r+1)²` block into `world+0x5400c` (12 of the 15 sites; `callto:R0020` = 14 hits / 13 owners / 0 orphan), `R0004` fills `session+0xba4` from the world actor list by **reach**, and `R0114`/`R0110` fill it from the same list by **sight**

**Confidence.** High (the signature is the twelve-byte stack clearance with the three pushes read at every site; the index arithmetic and the unlink test are single instructions; the polarity is discriminated by the two `wantBit = 1` callers doing the opposite thing, not by assumption). **Medium** that the caller list is complete — `callto:` sees direct calls and `.rdata` code pointers, so an indirect call through a computed pointer would be missed

### AI-ACQUIRE-002

It is the only caller of `R0114` (`callto:R0114` = 1 hit) and has **13 callers over 12 owners** itself, against `R0004`'s **one**, so it — not `HERO-TARGET-024`'s route — is what an unordered unit runs. Selection: `bestDist` is seeded `reach + 1` (the reach byte at `+0x12c` is read at `L00222` and incremented at `L00223`) and `bestTurn` `0xff`; per candidate `d = R0036(me, cand)` (`L00224`), **incremented again when the candidate's `vt+0x20` is 3 and the decider's is not** (`L00225`…`L00226`, i.e. a flier costs one more), `turn = R0115(R0051(me, cand), mover[0])` (`L00227`); `d < bestDist` takes it outright, `d == bestDist` takes it only if `d <= reach` **and** `turn <= bestTurn` (`L00228`…`L00229`); afterwards `bestDist > reach` clears the pick (`L00230`…`L00231`).

On success `order+8 = 6`, `order+0xc = target`, `order+0x14 = reach`. **The population has a second half nobody had described**: hostile candidates at `HP < 1` are moved to `session+0xbc4` (`L00232`, `L00233`, `L00234`), and when no living candidate survives **the dead are moved back and one of them becomes the target** (`L00235`; the same arm exists in `R0004` at `L00236`). With no target at all: a human owner's unit tries `R0015` (heal), an AI owner's goes to `order+8 = 0xb` (guard)

**Confidence.** High (every branch above is a cited instruction in one function read end to end, and the reach cap is decided by the seed `reach+1` plus the post-loop test, which no rival ordering satisfies). **Medium** that no *fourth* route exists — the enumeration behind that is `callto:R0021`, whose blind spot `AI-FILTER-001` states

### AI-LIST-003

The world is `0xa4558` bytes (the allocation size passed at `L00237`), so the field is its last dword. `disp:a4554` returns **19 hits over 14 owners, 0 orphan**, with exactly **two writers** — `L00238` in the world initialiser `R0116` and `L00239` in `R0117`, which has no caller (`callto:R0117` = 0 hits) — and both store a constructor argument, so the field is assigned once and never mutated. The single construction site passes the global `CObList` at `[L00240]`: read at `L00241` and passed to the constructor call at `L00242` (`R0118`). That global has **one writer** (`L00243`, in `R0119`, from the list constructor `R0120`) and one clear-to-zero (`L00244`) over **40 references across 28 owners** (`refto:L00240`).

Membership is the actor's own bit: `R0121` removes it and sets `actor+0x4c | 8`; `R0122` and `R0123` place it at a free cell, add it and clear that bit, printing `"Unit can't enter map - no free p…"` / `"Unit can't return to map - no fr…"` on failure. All three candidate producers walk it head to tail testing every element

**Confidence.** High (the size constant fixes the field's position, the two writers are the complete `disp:` population, and the construction site names the object; the identity is then confirmed from the other side by the add/remove pair and their strings). **Medium** on writer completeness: a `disp:` sweep cannot see a wholesale `REP MOVSD` of the world, and `imm:a4554` returns 0

### AI-DIPLO-004

The session constructor `R0124` zeroes `0x271` dwords there — 625 dwords = **2500 bytes = 50×50** — so with no further writing nobody is hostile to anybody; it also sets the four header bytes of the sub-object at `session+0xa9bc` and reads `World\Data\ai.reg` `[Scanning] MinimalGuardRange` (code default **10**, the constant pushed at `L00245`) into `session+0xa9b8` — *amended by [EXP-0082]: the **shipped** `ai.reg` sets it to **8**, so the default is never the effective floor; the file has exactly 4 records and `[Tasker] IntelligentCons…` = 15 has no literal in the image at all* — whose only reader is `R0125` (`L00246`/`L00247`/`L00248`), which raises a group's guard radius `grp+0x2d`/`grp+0x38` to it (`disp:a9b8` = 4 hits / 2 owners).

The row and column index is `Player+0x04`, a signed word (`L00249`, `L00250`), stride **50** (`LEA ×5` twice then `×2`). **Bit 0 = hostile; bit 1 blocks the flip**: `R0126` (`this = session+0xa9bc`, so the matrix is its `+8`) tests the low two bits and only then sets bit 0 on both `[A][V]` and `[V][A]` (`L00251`, `L00252`), notifying `R0127` for whichever side is a human participant. The engine reads the cell as three bits: the mask `0x7` at `L00253`. Whole-image sweep `disp:a9c4` = **32 hits / 14 owners / 0 orphan**, `imm:a9c4` = **0**

**Confidence.** High (the size, the base, the stride and the two bits are single instructions in two functions read end to end, and `R0126`'s `this` is fixed by `HERO-AGGRO-028`'s call site). **Medium** that the writer list is complete: the matrix is also reachable as `sub-object + 8`, which carries displacement `8` and is invisible to `disp:a9c4` — `R0126` is exactly such a site and was found through `HERO-AGGRO-028`, not through the sweep

**Amended.** The `MinimalGuardRange` default: 10 is the **code** default (the constant pushed at `L00245`); the **shipped** `ai.reg` sets it to **8**, so the code default is never the effective floor (superseded, EXP-0082). `retracted.md` holds the full entries.

### AI-DIPLO-005

`R0128` (the map loader — its own strings are `ALM-REQ-055`'s `"Tiles block not found"` / `"Altitudes block not found"`) iterates the player collection and, for each player `P` with index `idx = P+0x04` (= type-5 slot + 1), reads `k = 0…15` from `ALM-GRP-041`'s sixteen `u16` through `R0129(map+0x28, idx−1)` then `R0130(player+0x34, k)` and stores the **low byte** at `matrix[idx][k+1]` (the byte store at `L00254`, row base `session+0xa9c4 + 50·idx`); then it forces `matrix[idx][idx] = 2` (`L00255`). So the file's row is 0-based over players while the matrix is 1-based, and **column 0 is never written**. On a mission join `R0131` clones a reference player's row *and* column for `i = 2…15` (`L00256`/`L00257`), sets `matrix[ref][new] = 2`, and forces **0 in both directions between every pair of human participants** (`L00258`, `L00259`); `R0132` is the same with `2` in both directions.

**Corpus, 38 maps / 191 player records / 3056 row cells** (`tools/diplomacy`, replaying the store above): the values are exactly `{0, 1, 2}` — 2362 / 471 / 223 — so nothing truncates in the byte store (0 cells exceed `0xff`) and no bit above 1 is used; **no map ships more than 16 players**, so the 16-column row always covers the whole roster. The file spells its own diagonal as 2 in 186 rows, as **1 in 2 rows and 0 in 3**, so the forced diagonal is load-bearing rather than decorative. Over the 866 ordered off-diagonal pairs: **469 hostile, 37 carrying bit 1, and 102 disagreeing with their mirror on bit 0 across 19 of the 38 maps** — the relation the engine indexes is `[me][him]` and the shipped content uses that asymmetry (`Islands.alm`: `Guards → Monsters` = 1 while `Monsters → Guards` = 0)

**Confidence.** High (the store, its index arithmetic and the forced diagonal are cited instructions, and the corpus replay of that exact store closes over every shipped map). **Medium** that `{0,1,2}` is the whole value space and that no bit above 1 exists — that is a corpus statistic, and the engine's own `AND …,0x7` admits a third bit no shipped map uses

### AI-SIGHT-006

`disp:a5` returns **10 hits over 10 owners, 0 orphan**; two are writes (`L00260`, the actor constructor's default **5**, and `L00261`), one is an address computation of the field on a different object in the presentation module (`L00262` in `R0133`), and seven are reads. *Amended by [EXP-0120]: those two are not the writer set. The sweep is reproducible hit for hit and it was asked a question `disp:` on a byte cannot answer — **there are six writers**, and `+0xa4` is a `u16` sight in 1/256 cell of which `+0xa5` is the high byte (`AI-SIGHT-092`). The `complete writer set is two` clause is withdrawn.* The one that matters is `L00263` inside `R0134`, reached from `R0114` through `R0135`: it zeroes the `0x10000`-byte map at `world+0x82ef0`, marks the actor's own cell, and expands **ring by ring inside a 41×41 window** (`(cell>>8 & 0xff) − 0x14`, `(cell & 0xff) − 0x14`, loop bound `0x14`), each cell tested by `R0136` — *which [EXP-0112] reads whole and which is **not** an altitude comparison but an accumulator march carrying a budget from cell to cell, the altitude being one of its four terms (`AI-LOS-081`)* — against the height plane `world+0x9451c`, stopping when an entire ring is blocked; the radius enters as `(1 << (k−1)) + (scanRange << k)` at `fog+0x25450`.

`R0114` then walks `world+0xa4554` and keeps every actor whose cell is marked (`L00264`/`L00265`) — plus, for an AI-owned unit only, the remembered attacker's cell `order+0x58`, which is marked, counted in `order+0x5a` and forgotten after **20** ticks (`L00266`). The second consumer is `R0137` (AI state 0xb): a compare of the notice-radius byte at `+0xa5` (`L00267`) against the distance from the guard post `order+0x00`, and beyond it the unit is ordered to walk back. **Neither uses it to accept a target**: `AI-ACQUIRE-002`'s cap is `actor+0x12c`

**Confidence.** High (the sight stamp and the leash are cited instructions in two functions read end to end, and the 10-hit sweep is stated with its instrument). ~~**Unknown**: the remaining five readers `R0138`, `R0139`, `R0140`, `R0141`, `R0142` were located, not read~~ — *narrowed by [EXP-0082]: `R0140` takes the **max** of `actor+0xa5` over a group's members and adds it to the member's distance from the group centroid, which is where the group notice radius `grpAI+0x2c` comes from (`AI-RADIUS-014`); `R0138`/`R0139` are the follow pair (states 8 and 0x11) and use `actor+0xa5` as the **stop distance** of the follow order they emit, but only when `order+0x70` is 0 (`L00268`/`L00269`, `L00270`/`L00271`); a nonzero `order+0x70` is used instead. `R0141` and `R0142` are still located, not read* — *narrowed further by [EXP-0322]: `R0142` is no longer only located. It is a sight-radius-bounded ring search over candidate cells: `actor+0xa5` is its own loop bound (the ring-size comparison at `L00272`-`L00273`), each candidate is validated against the static plane (`R0143`) and `R0144`, and every candidate that passes is appended through the shared list class's own grow/insert primitive `R0052` (`SAV-632`) into a list based at the function's own `+0xa4538` (`SAV-GRPLIST-807`) — a fourth, list-building consumer of the byte, distinct from the group-radius and follow-distance uses already read above. `R0141` remains located, not read* — *and amended by [EXP-0088], which read `R0134`/`R0136` for the object rather than for the radius and corrects two mechanism sentences: **`fog+0x25450` is not a field.** It is `losAcc[20·64 + 20]` — the **centre cell of the 41×41 accumulator grid** at `fog+0x24000` (`0x25450 − 0x24000 = 0x1450 = 1300 dwords = 20·64+20`), so the seed `(1 << (k−1)) + (scanRange << k)` is written into the accumulator, not into a scalar, and `k` is `[Scanning] ScanShift` from `World\Data\map.reg`, shipping as **7** in both roots. **And what `R0134` zeroes is that accumulator, not the map:** the clear at `L00274` starts at offset `0x24000` and its count of `0x1000` dwords is `0x4000` bytes. The `0x10000`-byte byte map at `fog+0x2a008` (= `world+0x82ef0`) is cleared by a **separate** routine `R0145`, which has **two** callers — `R0114` `L00275`, which clears then stamps one actor, and `R0110` `L00276`, which clears then stamps **every member of a group**, so the same array serves a per-actor and a group-shared reading depending on the caller. The object's base is `world+0x58ee8` (`L00277`). `TERR-FOG-084`*

**Amended.** Two readings of the writer set stand: the row's `disp:a5` sweep with 10 hits and two writers (`L00260`, `L00261`), and `AI-SIGHT-092`'s six writers, which withdraws the two-writer clause. The writer-set clause withdrawn by EXP-0120. The writer-set clause only: There are **six** writers (refuted, EXP-0120). Two mechanism sentences: **Both are off by one structure.** (a) `R0134` zeroes `0x1000` **dwords** at `fog+0x24000` — `0x4000` bytes, the line-of-sight accumulator — not the `0x10000`-byte byte map (superseded, EXP-0088). The "still located, not read" residue on `R0142` only; the sight stamp, the leash, both named consumers, the writer-set withdrawal by EXP-0120 and the EXP-0082/EXP-0088 corrections all stand: `R0142` is a sight-radius-bounded ring search over candidate cells, not merely located: `actor+0xa5` is read at entry and used as its own loop bound (the ring-size comparison at `L00272`-`L00273`), each candidate cell is validated against the static plane (`R0143`) and `R0144`… (narrowed, EXP-0322). `retracted.md` holds the full entries.

### AI-GUARD-007

`R0009` is the engage routine: given a target, or else the first survivor of the occupancy block around its own cell, it sets `order+8 = 5`, `order+0xc = target`, `order+0x14 = actor+0x12c` — a target with a **stop distance**, which is the only construct in this area that survives being farther away than reach. Its 13 call sites over 9 owners are the AI's engage arms, `R0137`'s post scan among them. Both `R0009` and `R0004` refuse outright when `(actor+0x4c & 4) && actor+0x12c < 2 && Player+0x28 == 0` — a human participant's short-reach unit with that flag does **not** auto-engage; `R0004` writes the target into `order+0xc` first and then sets `order+8 = 0` (`L00121`, `L00122`, `L00123`), so the pointer is left behind but the order is idle.

And `R0146` (state 0x17) **overwrites both radii on the actor**: the byte stores at `L00261` (`+0xa5`) and `L00278` (`+0x12c`), both of the value 5, gated on the word at `+0xe` equalling `0x18`, immediately before calling `R0022` — which is why EXP-0072 saw sight and reach written together, and which means **`actor+0x12c` is not immutable class data**

**Confidence.** High (the order fields, the two suppression clauses and the paired radius write are cited instructions). ~~**Medium** on the *meaning* of `actor+0x4c` bit 2 and of `actor+0x0e == 0x18`~~ — *both settled by [EXP-0112]: bit 2 is `HERO-CLASS-013`'s **mage** flag, so this row's suppression clause is about a human participant's short-reach spellcaster and `R0113` is the caster AI; `actor+0x0e == 0x18` is `Humans` typeID 24, 57 named rows on both roots, tested only inside an arm `AI-STATE-043` shows has no writer (`AI-GATE-079`)*. ~~and on how `order+8` states 1/5/6/0xb are consumed~~ — *closed by [EXP-0099] (`AI-ORDER-039`, `AI-PURSUE-040`) and, for the group orders that re-issue them, by [EXP-0112] (`AI-REISSUE-077`)*

## The AI tick, groups, group orders and the actor state machine

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-TICK-008 | The AI is one slot of the full tick, it is driven per group, and the per-actor state machine is a single arm inside it. | High / Medium | ● active (contested) | [EXP-0082](../experiments/EXP-0082-ai-behaviour/) |
| AI-GROUP-009 | The unit of AI is a runtime *group*: a `0x48`-byte `CObList` subclass owning a `0x50`-byte AI record, formed at spawn by equality of the `.alm` type-6 group id. | High / Medium | ● active | [EXP-0082](../experiments/EXP-0082-ai-behaviour/) |
| AI-ORDER-010 | `grpAI+0x20` is the group order byte and it has eight live values; `0` is what hands control to the members' own states. | High / Medium | ● active | [EXP-0082](../experiments/EXP-0082-ai-behaviour/) |
| AI-STATE-011 | `actor+0x50` is the per-actor AI state, its switch has 27 arms of which 12 are distinct, and its value at construction is `0xb` — guard. | High / Medium | ● active (amended, superseded) | [EXP-0082](../experiments/EXP-0082-ai-behaviour/) |
| AI-GUARD-012 | Guard is the whole of “motion without an order”, and home is the cell the actor stood on when guard first ran. | High / Medium / Unknown | ● active (amended, partially retracted, superseded) | [EXP-0082](../experiments/EXP-0082-ai-behaviour/) |
| AI-PATROL-013 | There is a patrol state, it is a ring of waypoints, and the mission description file that authors one does not ship. | High / Medium | ● active (partially retracted) | [EXP-0082](../experiments/EXP-0082-ai-behaviour/), **[EXP-0083](../experiments/EXP-0083-patrol/)** |
| AI-RADIUS-014 | The aggression radius exists, it is separate from reach and from sight, and it belongs to the *group* — derived from geometry, floored by `ai.reg`, and jittered by ±1 every time it is used. | High / Medium | ● active (amended, superseded, contested) | [EXP-0082](../experiments/EXP-0082-ai-behaviour/), **[EXP-0083](../experiments/EXP-0083-patrol/)** |
| AI-AUTHOR-015 | A behaviour is not authored per unit anywhere. | High / Medium | ● active (amended, superseded) | [EXP-0082](../experiments/EXP-0082-ai-behaviour/), **[EXP-0083](../experiments/EXP-0083-patrol/)** |
| AI-DIFF-016 | The difficulty level does not reach behaviour — not as stats read by the AI, not as counts, not as aggression, not as cadence. | Medium | ● active | [EXP-0082](../experiments/EXP-0082-ai-behaviour/) |

### AI-TICK-008

`SESS-TICK-006` located `R0076` at `server+0x04 % 16 == 6` and left it unread; it is the whole AI. `R0147` reaches it at `L00279` after the signed modulo at `L00280`…`L00281` (the compare against 6), so the AI runs **once per full tick** — 992 ms at the shipped speed index (`SESS-CLOCK-005`) — and never per sub-tick. `EnumRefs callto:R0076`: **2 hits / 2 owners / 0 orphan** (`R0147`, `R0075`). `R0076(playerList)` walks players, then each player's group list at `player+0x24`, and calls the group order machine `R0023` once per group (`L00282`), gated on `session+0xb388 != 0` **or** that group's own `grpAI+0x45` (`L00283`), whose constructor default is 1.

Nothing in it iterates actors. The per-actor state machine `R0008` has **2 callers** (`EnumRefs callto:R0008`: 2 / 2 / 0) — `R0148` and `R0023` at `L00284`, the latter being jump-table **arm 0** of the group order — so **an actor's own AI state is evaluated only while its group's order byte is 0**. The routine also profiles itself and, when `[L00285]+0x118`'s `[+0]` is set, queues a per-group average into the UI (`L00286`…`L00287`)

**Confidence.** High (the slot, the two walks, the per-group gate and both caller enumerations are named instructions and printed counts on the repaired function table). **Medium** that the caller sets are complete — `callto:` sees direct calls and `.rdata` code pointers, so an indirect call through a computed pointer would be missed

**Amended.** Two readings of the cadence clause stand. The row reads it from `R0147`, which reaches the AI walk at `L00279` behind the signed `% 16 == 6` test. `AI-CLOCK-080` finds a second call to the walk, `L00163` in `R0075`, with no modulo in front of it; whether that routine runs in ordinary play is Unknown, and if it does the AI decision rate is not once per full tick.

### AI-GROUP-009

Constructor `R0149`: `grp+0x40 = 0`, `grp+0x44 = 0`, then `new(0x50)` then `R0150` then `grp+0x3c` (`L00288`…`L00289`). That record's constructor zeroes `0x14` dwords, sets `grpAI+0x20 = 0` (the order) and **`grpAI+0x45 = 1`** (AI enabled), and allocates a `0x1c`-byte `CList` with growBy 10 at `grpAI+0x4c` — the patrol path (`L00290`…`L00291`). Membership: the type-6 spawner `R0151` adds each actor to the world list, sets `actor+0x14 = Player`, then scans `player+0x24` for a group whose **`grp+0x1c` equals the loader struct's `+0x40`** (`L00292`/`L00293`); on a miss it allocates `0x48` (`L00294`), appends to `player+0x24` (`L00295`) and stores that value (`L00296`).

The struct's `+0x40` is the file's `+0x42` at instruction level — case 6 reads 4 bytes into `struct+0x40` at `L00297` and keeps its running maximum at `mapObj+0x90` (`L00298`/`L00299`) — so this is `ALM-UNIT-048`'s **group id** with a consumer, not a label from the trigger grammar. `R0152` (AddMember, 15 sites / 11 owners) also writes `grp+0x44 = actor+0x14`, the owning `Player`. **Corpus, 38 maps / 8094 placed units** (`tools/aibehav`, replaying that rule per owner-and-groupId pair): **2264 groups**, mean **3.58** members, **686 singletons (30.3 %)**, largest **133** (`Horror.alm`), most groups in one map **362** (`Forester.alm`), and only **5** units carry group id 0

**Confidence.** High (the allocation sizes, the ctor defaults, the match instruction and the loader's own store are cited instructions; the file-to-struct binding is a read boundary, not an inference). **Medium** that the corpus figures describe what the engine builds — they replay the match rule but do not observe the engine doing it

### AI-ORDER-010

`R0023` forces `0xff` when the group is empty (`L00300`), then dispatches on the order value (`L00301` to `L00302`): above `0x11` only `0xff` acts, `0x11` takes its own arm, 6 to `0x10` do nothing, and 0 to 5 index the jump table at `L00303`. The 6-entry table, **read out of the PE through its own section table** (`tools/aibehav`), is `L00304 L00305 L00306 L00307 L00308 L00309` — **0** per-member `R0008`; **1** `R0024` guard; **2** `R0025(grp, grpAI+0x0a)`, `MOVE-ORDER-023`'s group move; **3** `R0110` + `R0153` + per-member engage/heal — *labelled "attack" here and **amended to `Stand Ground` by [EXP-0112]**: the arm has no radius clip and no walk, and its scorer vetoes every candidate past reach (`AI-STAND-076`, `AI-REACH-072`)*; **4** `R0154`; **5** `R0026`; plus **0x11** `R0155` (`L00310`) and **0xff** per-member `R0022` (`L00311`).

After every arm a tail runs for every member: `R0105` then `R0100`, else `R0106` then `R0107` (`L00312`…`L00313`). Writers, from `EnumRefs` on both store forms filtered to the AI module: `0` — the ctor plus twelve command handlers over 8 owners; `1` — `R0125` only (`L00314`); `3` — `R0060` (`L00143`), `R0063` (`L00153`), `R0156` (`L00315`); `2` — `R0156` and one orphan run `ORPHAN[L00316..L00317]` at `L00318`; `4` — `R0108`; `5` — `R0157`; `0x11` — `R0158`, `R0156`. **A group under order 1 or 3 never evaluates `actor+0x50` at all**

**Confidence.** High (the bounds, the table and the arm targets are the routine's own instructions plus a PE read, and the table's extent is fixed by the bound of 5 tested at the dispatch). **Medium** on the writer enumeration's completeness — it is over two store forms (a byte store of an immediate at `+0x20`, and the form that stores a zeroed register), so a wider store spanning `+0x20` and its neighbour, or a copy of the whole record, is invisible

### AI-STATE-011

`R0008`: dispatches on the order value (`L00319` to `L00320`): above `0x1a` goes to the default at `L00321`, and 0 to `0x1a` index the jump table at `L00322`. The 27-entry table, read out of the PE: `0` clear `actor+0x54` and return; `1` re-issue the stored move (`order+8 = 1`); `2` walk to `order+0xa`, then `order+8 = 7` on arrival; `3` `R0009(actor, order+0xc)`; `4` occupancy block at `order+0xa`, diplomacy filter, engage, else walk there; **`0xa` `R0159` — patrol**; **`0xb` `R0137` — guard**; **`0xc` `R0022` — acquire with no leash**; `8` and `0x11` `R0138`/`R0139(actor, order+0x10)`, the follow pair, whose own 5-arm inner table sits at `L00323`; `0xd` `R0018`; `0xe` and `0xf` set `order+8` to 9 and 0xf; `0x16` `R0100`; `0x17` `R0146` (`AI-GUARD-007`'s radius overwrite); `0x18` and `0x19` write **hard-coded** cells (`0x4f5b`/`0x2f52`; `world+0xa453c`); `0x1a` stores `0x1a` into `actor+0x54`; `5,6,7,9,0x10,0x12..0x15` fall to the default `L00321`, which clears `actor+0x54` and runs `R0022`.

**The default value is written by the actor's own initialiser** `R0160`: the store of `0xb` into `+0x50` at `L00324`, a few instructions from `AI-SIGHT-006`'s sight default `L00260` and `UNIT-COMBAT-006`'s reach default `L00325`. **The `.alm` spawner writes no `+0x50`**: `EnumRefs re:` over the immediate-store form returns **81 hits / 55 owners / 0 orphan** and `R0151` is not among them

**Confidence.** High (the bound, the table and the constructor default are cited instructions and a PE read; the arm labels are each that arm's own call or store). **Medium** that the state's writer set is complete — the sweep is over one store form and cannot see a state assigned through a register or copied with the actor. *[EXP-0090]: **C-6 is resolved and this row's arm-2 reading is confirmed exactly** — `actor+0x50 = 2` has one consumer, this arm, and one player-side writer, `R0010` (order opcode `0x21`), which also supplies the `ord+0x0a` the arm walks to; `ITEM-CMD-007`'s "deferred order" and this row's "AI state" are one path (`ITEM-PICK-016`). The row's own **Medium** on writer-set completeness is **vindicated rather than lifted**: the register-form writers it could not see are real, and one of them, `L00026`, is the pick-up's bridge (`AI-PROGRESS-034`)*

**Amended.** The arm inventory: The table and the count stand (superseded, EXP-0087). `retracted.md` holds the full entries.

### AI-GUARD-012

Guard is the whole of “motion without an order”, home is the cell the actor stood on when guard first ran, and ~~there is no wander anywhere in the module~~. — *first clause amended by [EXP-0083]: on the eight maps whose script commands a patrol it is not, and the routine this row cites as an RNG destination site, `R0161`, is itself a patrol setter (it stores `0xa` into `+0x50` at `L00326`) rather than a stand-alone destination picker — dead, since `callto:R0175` on its only caller returns 0. The broad “no wander” clause remains withdrawn: [EXP-0087] identifies a Group-AI destination generator outside the **per-actor state machine** examined by [EXP-0082]. `R0155` picks one of eight compass directions and steps 20 cells, re-rolling until the result is inside the map rectangle (`AI-ROAM-025`). Its local Group-cell/counter program and Swarm2 evaluation call do not by themselves prove a producer for actor destinations or a composed wandering-motion result (`AI-SWARM2GATE-107`, `AI-MOVE-023`). **`Roam` ships 0 nodes over 38 measured maps**; this is an authored-node census, not proof that no shipped actor wanders through any path. Neither the former positive “it is the wander” wording nor a restored whole-image no-wander conclusion follows. See [`retracted.md`](retracted.md).* `R0137` (state `0xb`): if `order+0x00` is zero it is set to the actor's **current** cell (`L00327`…`L00328`) — so home is emergent, never authored; the distance from it (`R0162`, `L00329`) is compared against `actor+0xa5` (`L00267`) and beyond it the actor gets `order+8 = 1`, `order+0xa = order+0x00` — **walk back** (`L00330`/`L00331`); inside it, the occupancy block around the **post** is filtered by `R0021` with `wantBit 0` and any survivor goes to `R0009` (`L00332`…`L00333`); with nothing to fight, an actor away from its post walks back (`L00334`) and one at its post runs `R0022` (`L00335`).

The group arm `R0024` (order 1) repeats this for a whole group around `grpAI+0x28` and ends each member the same way (`L00336` engage, `L00337` `order+8 = 0xb`, `L00338` walk home). **The following per-actor RNG-site notes are not a no-wander proof**: the earlier enumeration cited 19 `rand()` sites in `L00339..L00340` (`EnumRefs callto:R0179`, 87 / 44 / 0 image-wide); its two described per-actor *destination* sites write constants — `R0161` builds x in {70,100} from `rand()*2/0x8000` and y in 60..70 from `rand()*11/0x8000` (`L00341`…`L00342`), and jump-table arm `0x18` writes the literals `0x4f5b` and `0x2f52` (`L00343`, `L00344`). Neither is a neighbourhood of the actor

**Confidence.** High (the post's initialisation, the leash test and both walk-back stores are cited instructions in two routines read end to end). The former **Medium** whole-image “no wander” conclusion is withdrawn. The local Group-cell generator does not establish composed actor wandering motion; that conclusion remains **Unknown**

**Amended.** Dependent Roam proof and zero-node-to-no-motion inference only: The Group-AI destination generator invalidates treating the per-actor scope as complete (narrowed). The first clause: Every instruction the row cites stands (superseded, EXP-0083). The no-wander clause, second overturn: The census is exact **and it is a census of the per-actor state machine** (refuted, EXP-0087). `retracted.md` holds the full entries.

### AI-PATROL-013

~~**There is a patrol state, it is a ring of authored waypoints, and no shipped file reaches it.**~~ — *the second clause is **refuted** by [EXP-0083]: the `.ini` census below is exact and stands, but the mission description file is not the only writer of `actor+0x50 = 0xa`. The `.alm`'s **own type-7 script** reaches the state — 14 nodes on 8 campaign maps (`AI-PATROL-017`) — and the ring it builds is two nodes, not an authored path (`AI-PATROL-018`). Read the original clause as: no shipped file reaches **`R0163`**. The rest of the row is untouched.* `R0159` (state `0xa`) clears the move order, **runs guard first** (the call to guard at `L00345`: a patrolling actor still fights), and when guard leaves the order idle or guarding it advances the path: `order+0x02` is the current waypoint; on arrival it finds that node in the list at `order+0x90` and takes the **next**, falling back to the head, so the path is a **ring** (`L00346`…`L00347`); then `order+8 = 1`, `order+0xa = waypoint`.

`R0159` has **one caller** — the jump table. The path is loaded by `R0163(group, sectionName)`, whose own strings are `Patrol path in section [`, `] not found`, `Error in patrol point [`, `] line `, `Patrol path loaded.`: it empties `grpAI+0x4c`, finds the named section of the **mission description file**, and for each line requires a `';'` at index 1 or above (`L00348`/`L00349`), `atoi`s the two halves, requires **`7 < x < 0x88` and `7 < y < 0x88`** (`L00350`…`L00351`), appends `x + (y << 8)`, and if the list is non-empty calls `R0164`, which puts every member into state `0xa`. `SESS-LOAD-009` gives that file's path: `World\Mission\<n>.ini`. **Corpus: 0 nodes ending `.ini` or under `world\mission` in any of the 8 live archives, and 0 files under any `Mission` directory in the GOG root (120 files), `gameversions/en` (120) or `gameversions/ru` (42).** So the state, its loader, its four operator strings and its ring-walk are reachable code with no shipped input

**Confidence.** High (the ring advance, the semicolon grammar and the four bounds are cited instructions in two routines read end to end; the corpus half is a walk of both preserved roots and every archive's node list, which is immune to a code-enumeration blind spot). **Medium** that a patrol can *only* be authored this way — `R0164` has one caller and `grpAI+0x4c` was not swept for other writers — *[EXP-0083] swept the **state** instead (`re:` on the store form, 4 hits / 4 owners / 0 orphan) and found three further setters, one of them live: `AI-PATROL-019`*

**Amended.** The "no shipped file reaches it" clause: The census is exact and stands: no `World\Mission\<n>.ini` ships (refuted, EXP-0083). `retracted.md` holds the full entries.

### AI-RADIUS-014

`R0140(grp, memberCount)` sums each member's fine coordinates (`R0165`/`R0166`), divides by the count, and stores the centroid at `grpAI+0x28` (cell) and `grpAI+0x24` (fine) (`L00352`…`L00353`); then over the members it takes `grpAI+0x2a` = max Chebyshev distance from the centroid (`R0167`), `grpAI+0x2b` = max member sight `actor+0xa5`, and **`grpAI+0x2c` = max of (distance + that member's sight)** (`L00354`/`L00355`/`L00356`, stored `L00357`/`L00358`/`L00359`). `R0125` installs it: `grpAI+0x2d = grpAI+0x38 = grpAI+0x2c`, then either raises both to the caller's override or, when that is 0, to **`[Scanning] MinimalGuardRange`** at `session+0xa9b8` (`L00360`…`L00361`).

The group guard arm re-rolls the working radius ~~on every run~~ (*corrected by [EXP-0083]: on the tick a latch flips — the roll is gated on `grpAI+0x30`, which the same arm sets (the latch at `+0x30` is read at `L00362`, compared at `L00363`, and the branch at `L00364` jumps to `L00365` when it differs), and the taken branch only increments the counter `grpAI+0x34`; the `+4` added at `L00366` is on the has-members edge alone, the empty-group edge rolling the same expression without it at `L00367`…`L00368`. `AI-GUARD-021` names both fields*): `grpAI+0x2d = grpAI+0x38 + 4 + r` with `r = rand()*3/0x8000 - 1` in {-1,0,1} (`L00369`…`L00370`; the divisor is `session+0x00`, a constant `0x8000` set at `L00371`).

It clips the **candidate list** around `grpAI+0x28` (`L00372`/`L00373`, drop at `L00374`); it is not a hit range — `AI-ACQUIRE-002`'s `actor+0x12c` still gates the blow. **Measured: the shipped `World\Data\ai.reg` sets `MinimalGuardRange = 8`, not the code default `10`** (the constant pushed at `L00245`), so no group's radius is ever floored at 10. The RNG is the CRT's LCG (state times 214013 plus 2531011, shifted right 16 and masked `0x7fff`), seeded once per session from `time` (`L00375`/`L00376`)

**Confidence.** High (every arithmetic step is a cited instruction in two routines read end to end, the divisor is a constant with a named writer, and the shipped value is a record in a 4-record file). ~~**Medium** that `grpAI+0x2c` is never authored — the claim rests on `R0140` being the only producer, which was not swept for~~ — *swept by [EXP-0112]: there is a **second** producer, `R0141`, and it is the one that runs — every guard tick of every group, as `R0024`'s first act. Neither is authored; what the sweep also shows is that `grpAI+0x2c`'s **only readers** are the two inside `R0125`, so the value this row computes is consumed at load and never again, and the radius the arm rolls is frozen at the last guard issue (`AI-RADFREEZE-075`)* — and **Medium** that the roll is once per latch flip, since `grpAI+0x30` was not swept for writers outside `R0024`

**Amended.** Two readings of two clauses stand. The roll cadence: the row's original wording, "on every run", against the tick a latch flips (`AI-GUARD-021`). The producer of `grpAI+0x2c`: `R0140` alone, against a second producer, `R0141`, that runs on every guard tick and whose result is never read again after load (`AI-RADFREEZE-075`). The cadence clause: The arithmetic is right; the cadence is not (superseded, EXP-0083). `retracted.md` holds the full entries.

### AI-AUTHOR-015

~~Two setters exist, and what picks between them is the map's roster order — or, for content that does not ship, two keys in the mission description file.~~ — *amended by [EXP-0083]: the per-unit half stands, because a command names a **group**; the rest does not. What the two setters decide is the **load-time** stance, and a shipped campaign map then commands its groups from its **own type-7 script** — 135 group-command nodes over 20 maps, nine of the eleven authored sub-commands, patrol among them (`AI-GROUPCMD-020`, `AI-PATROL-017`).* The load-time half, unchanged: `R0125` ends **guard**: every member `actor+0x50 = 0xb` (`L00377`) and `grpAI+0x20 = 1` (`L00314`). `R0060` and `R0063` end **aggressive** — *the name is **amended to `Stand Ground` by [EXP-0112]**, which reads the arm that order selects: no clip, no walk, and a reach-vetoing scorer (`AI-STAND-076`); the key that picks this setter is itself spelled `StandGround` two sentences below* —: every member `actor+0x50 = 0xc` (`L00140`, `L00150`) and `grpAI+0x20 = 3` (`L00143`, `L00153`).

Both setters are **also** reachable as player-issued group commands out of the dispatcher `R0061` — `L00378` → guard, `L00127` → aggressive (`EnumRefs callto:R0125` = 6 / 6 / 0, `callto:R0060` = **2 / 2 / 0**, `callto:R0063` = 6 / 5 / 0) — so the load-path choice below is the *initial* setting, not a permanent one. On the map-load path `R0128` calls `R0067(map, 1)` (the constant 1 pushed at `L00379`); that routine walks `[L00380]`, then `player+0x24`, and — gated on its argument and on **`server+0x0c == 0`, i.e. single player** (`L00381`…`L00382`) — reads `grp+0x44` (the owning `Player`) and compares `player+0x04` against **1**: the **first type-5 record**, since `ALM-GRP-041`'s loader stores `slot+1` there.

That player's groups get `R0063`; every other player's get `R0125(grp, 0)` (`L00383`…`L00384`). **Corpus: type-5 slot 0 is named `Self` on 38/38 shipped maps** (32 `Self`, 6 `self`). The second surface is `R0062`, which builds a group from a section of the mission description file and then reads `Patrol` — a section name, non-empty then `R0163` (`L00385`…`L00386`) — else `StandGround` compared against `Yes` (`L00387`…`L00388`), equal then `R0060`, otherwise `R0125`. `Yes` is the **comparand, not a default**: `R0168` takes three stack arguments plus the object and returns the global empty `CString` for an absent key (the address `L00389` pushed at `L00390`), so an absent `StandGround` means guard *[EXP-0125]: the open item on `R0067`'s gate is **closed, and it wraps the whole walk**: `L00391` and `L00382` both carry a `rel8` to `L00392`, the group-list advance, with the two setter calls the only instructions between. So a multiplayer load leaves every group at the constructor's 0 — and that is benign, because order 0 hands each member to its own `actor+0x50` machine, which writes the post instead (`AI-GATE-100`). This row's setters are also shown to anchor a post before they set the order (`AI-POST-095`, `AI-STANCE-098`).*

**Confidence.** High (both setters, the roster test, the single-player gate and the two key reads are cited instructions; the stack-argument count fixes the comparand reading, and the 38/38 slot-0 name is a corpus census that could have failed). **Medium** on `server+0x0c == 0` meaning single player — that is `SESS-MAP-010`'s reading of `map+0xd4 > 1`, quoted here rather than re-derived

**Amended.** The "two setters" clause: The per-unit half stands — a command names a **group** (superseded, EXP-0083). `retracted.md` holds the full entries.

### AI-DIFF-016

`UNIT-GATE-014` fixes the control (`precreate+0x1d0`, three buttons) and `UNIT-GATE-012` its state (`[L00285]+0x84`, value set 1/2/3, default 2), and `UNIT-GATE-013` gives its one consumer: the `.alm` spawner scaling the actor's health, toHit and defence. **`EnumRefs refto:L00285` returns 172 hits over 93 owners, 0 orphan; 21 lie in `L00339..L00340` and not one of them reads `+0x84`** — they take `+0x118` (the message/profiler sub-object: `L00286`, `L00393`, `L00394`, `L00395`, `L00396`, `L00397`), `+0x14` (the world sub-objects: `L00398`, `L00399`, `L00400`, `R0169`, `L00401`, `L00402`, `L00403`), `+0xa4 + 4*i` (`L00404`), and seven address computations of the field at `+0x88` on a neighbouring `CMap` in `R0170`.

Nothing else carries the level in: the only registry the AI module reads is `World\Data\ai.reg`, whose whole content is **4 records** — subkeys `Scanning` and `Tasker`, and the scalars `MinimalGuardRange = 8` and `IntelligentCons` = 15 — with no per-level dimension; the group count is content, the cadence is `SESS-CLOCK-005`'s speed index, and both radii are computed from geometry (`AI-RADIUS-014`). **`Tasker` and `Intelligent` occur zero times as literals in `rom.exe`** (against one each for `Scanning` and `MinimalGuardRange`), so that second scalar is dead config and cannot be the missing lever either

**Confidence.** **Medium** — this is a negative resting on a reference sweep, and `refto:` cannot see the singleton reached through a pointer another routine has already loaded, nor a whole-object copy. What it establishes at High is the narrower statement it is written as: **no instruction in the AI module names `[L00285]` and reads `+0x84`**. The `ai.reg` half is immune — a 4-record census and a literal count, not a code enumeration

## Patrol and the authoring surfaces

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-PATROL-017 | A shipped map authors patrol through its own `.alm` script, and fourteen nodes on eight campaign maps do. | High / Medium | ● active | **[EXP-0083](../experiments/EXP-0083-patrol/)** |
| AI-PATROL-018 | What the command builds is a two-node ring between the creature's own cell and the commanded one — and it opens `AI-ORDER-010`'s gate itself. | High | ● active | **[EXP-0083](../experiments/EXP-0083-patrol/)** |
| AI-PATROL-019 | Four routines put an actor into the patrol state; two are live, one has no input and two are dead. | High / Medium | ● active | **[EXP-0083](../experiments/EXP-0083-patrol/)** |
| AI-GROUPCMD-020 | The third authoring surface is the map's own script, it is the one shipped maps use, and it writes the group order directly. | High / Unknown | ● active (amended, superseded) | **[EXP-0083](../experiments/EXP-0083-patrol/)** |
| AI-GUARD-021 | `grpAI+0x30` is a has-members latch and `grpAI+0x34` a tick counter with no reader — the guard arm's unread pair, and neither is a wander. | High / Medium | ● active (amended) | **[EXP-0083](../experiments/EXP-0083-patrol/)** |

### AI-PATROL-017

The action vocabulary's opcode 6 is a second dispatch on the node's first parameter (`TRIG-GROUP-005`); the shipped editor catalogue `Description Instants.ini` names literal **14** `Group Command : Patrol`, with `Par1 = X`, `Par2 = Y` and `Par9` a `Target_Group`. The runtime honours it: `R0156` reads `Par0` at `rec+0x08`, the bound check at `L00405` to `L00406` limits the switch to `Par0 ≤ 17` (anything above goes to `L00407`), and the 17-entry table at `L00408` — **read out of the PE through its own section table** (`tools/aipatrol`) — has entry 13 = `L00409`, an arm that calls `R0171(group)` (`L00410`) and then `R0172(member, X, Y)` for every member (`L00411`).

**Corpus, 38 maps: 135 opcode-6 nodes over 20 campaign maps and 0 over the ten loose maps**, of which **14 are `Par0 = 14`, on 8 maps** — `scn:10` ×2, `scn:50`, `scn:60`, `scn:80` ×2, `scn:90`, `scn:100` ×3, `scn:131`, `scn:141` ×3 — naming groups whose **34** type-6 members are owned by `Villagers`, `Enemies`, `Beists`, `Monsters`, `Hima` and `Friends`, and by **no** type-5 slot 0, the `Self` of `AI-AUTHOR-015`. Every one hangs off a trigger whose condition pair is an opcode-`0x10002` constant node compared **with itself** under code 0 (`TRIG-COND-003`, `TRIG-CMP-006`), so it fires on the first evaluation pass; seven of the eight maps set `+0xb4 = 1` (fire once) and `scn:131`'s `"WhenDragonStartPatrol"` sets 0, re-issuing every full tick (`TRIG-FIRE-007`). `scn:10` is the mission a new campaign enters from hero creation (`SHOP-TOWN-023`), and its trigger `"Just Start"` commands two. The `ru` root carries the identical census

**Confidence.** High (the bound and the arm are the routine's own instructions, the table was read out of the PE rather than off a listing, and the census is a walk of every shipped map's own records — a map with no patrol node would have shown as one; the catalogue is a shipped file this repo did not author). **Medium** for the per-node member counts, which replay the group id against the type-6 records rather than observe the resolver, and which one node (`scn:100` group 25) makes ambiguous by naming a group id two owners share

### AI-PATROL-018

`R0171(group)` stops every member (`R0007`'s work inline: `order+0x14 = actor+0x12c`, `order+0x60 = 0`, `mover+0x7c = 0`, `actor+0x54 = 0`, `order+0x50 = 0`, `order+0x38 = 0`) and then, once, **the byte store of 0 at `L00412` into `grp+0x3c` plus `0x20`** (the zero is produced at `L00413`) — `grpAI+0x20 = 0`, the arm that hands each member to its own `actor+0x50` state machine. `R0172(actor, X, Y)` then writes **`0xa` into `+0x50` at `L00414`**, sets `order+0x00` to the actor's current cell (`L00415`), **empties** the waypoint list at `order+0x90` (`L00416`…`L00417`) and rebuilds it with exactly two nodes: the actor's **own current cell** (`L00418`) and `(Y << 8) | X` (`L00419`…`L00420`, the engine's own cell packing, `SAV-OBJ-014`); `order+0x02` is then the tail (`L00421`), `order+0x04 = 0`, `order+0x08 = 0`, and the visibility footprint is re-stamped (`L00422`).

So a shipped patrol is **not** an authored path: it is one commanded point and wherever the creature happened to be standing. `R0159` walks it (`AI-PATROL-013`), and `order+0x04` is a **re-anchor latch** — set on every advance (`L00423`), consumed at the next entry to move the guard post `order+0x00` to the actor's current cell (`L00424`…`L00425`) — which is why `AI-GUARD-012`'s leash never pulls a patroller home

**Confidence.** High (every store is a cited instruction in two routines read end to end, and the group-order clear is the one write that decides whether any of it executes)

### AI-PATROL-019

An `EnumRefs` pattern query for 4-byte stores of the immediate `0xa` into `+0x50` returns **4 hits / 4 owners / 0 orphan**: `L00414 R0172`, `L00426 R0173`, `L00326 R0161`, `L00427 R0164`. `R0172` serves **two** callers (`callto:R0172` = 2 / 2 / 0) — the script arm `R0156` and `R0174`, whose own single caller is the **player command dispatcher** `R0061` at `L00428`, one slot below the guard and aggressive commands `AI-AUTHOR-015` names at `L00378`/`L00127`; `R0174` clears the group order the same way (`L00429`) before the per-member loop, **so a player can order a patrol too**. `R0164` is the mission-`.ini` route (`callto:R0164` = 1, from `R0163`), which ships no input.

`R0173` has **0 callers**; `R0161` has one (`R0175`) whose own count is **0** — and it is the routine `AI-GUARD-012` cites as an RNG destination site, so that site is a patrol setter rather than a wander: it picks `x ∈ {70,100}` and `y ∈ 60..70` into `order+0x02` while the ring it builds holds the actor's cell and a literal `0` (`L00430`), which `R0159`'s search cannot find

**Confidence.** High for the four writers and each caller count (each is a printed enumeration on the repaired table, with 0 orphan hits, and each callee was read). **Medium** that the writer set is complete — the instrument is a regex over one store form and cannot see a state assigned through a register or copied with the actor, the same blind spot `AI-STATE-011` states

### AI-GROUPCMD-020

Nine of the eleven `Par0` literals the catalogue declares appear in shipped content, 135 nodes over 20 campaign maps and none over the ten loose ones: `1 Guard` ×1, `2 Swarm` ×6, `3 Stand Ground` ×24, `4 Move` ×29, `5 Swarm 2` ×40, `10 Attack` ×5, `11 Defend` ×6, `14 Patrol` ×14, `15 Follow` ×10. What each does to `grpAI+0x20`, from the arms and from an `EnumRefs` pattern query for byte stores of an immediate into `+0x20`: `2` → **2** (`L00431`, plus `grpAI+0x0a = (Y<<8)|X`), `3` → **3** (`L00315`), `17 Roam` → **0x11** (`L00432`, with `grpAI+0x0a` = the first member's cell), `4` and `5` through `R0108`/`R0157` → **4** (`L00433`) and **5** (`L00434`); `1` delegates to `AI-AUTHOR-015`'s guard setter `R0125`; and **`10`, `11`, `14` and `15` leave the group order at 0**, each by writing it — `L00435` and `L00436` in the arms themselves, `L00412` in `R0171` for `14` (and for `17`, before it raises `0x11`), `L00437` in `R0176` for `15` — every one a byte store into `+0x20` of a register zeroed at the routine's head, which is why the immediate-form store sweep does not see them. So on a shipped campaign map a group's order is **not** the load-time 1 or 3 for the whole mission; it is whatever the last command left. `Par0 = 18` (`Dwell`) has no case and no shipped node, and `17` has a case and no shipped node

**Confidence.** High for the census and for the order byte each arm writes (the sub-dispatch is one routine read whole against a PE-read table, the store sweep prints its own 75 hits / 32 owners / 1 orphan, and the catalogue is shipped). **Unknown** what group orders 2, 4, 5 and `0x11` *do* past their entry points `R0025`/`R0154`/`R0026`/`R0155` — the names `Swarm`, `Move`, `Swarm 2`, `Roam` are the catalogue's, not a reading of those arms

**Amended.** The "named from the catalogue" bound: The four arms are read (`AI-SWARM-022`…`AI-ROAM-025`) and three of the four names are misleading: order 2 `Swarm` is *walk to a cell and fight what you see*, order 5 `Swarm 2` is the same **without** the walk and falls back to order 4 when `AImanager+0xbb4` is clear… (superseded, EXP-0087). `retracted.md` holds the full entries.

### AI-GUARD-021

In `R0024`: `L00438` loads the group's member count, the latch at `+0x30` is read at `L00362`. Non-empty and latch 0 → roll `grpAI+0x2d = grpAI+0x38 + 4 + r` (`L00369`…`L00370`), `grpAI+0x34 = 0` (`L00439`), `grpAI+0x30 = 1` (`L00440`). Non-empty and latch set → **only** `grpAI+0x34++` (`L00365`…`L00441`) and the latch re-written. Empty and latch set → the same roll **without the `+4`** (`L00367`…`L00368`) and the counter zeroed; empty and latch clear → counter++; then latch = 0 (`L00442`). Nothing in the routine reads `+0x34`, and nothing in it picks a destination — so `AI-GUARD-012`'s Medium clause ("a wander built from a counter rather than `rand()` would not appear") is discharged for this arm, and `AI-RADIUS-014`'s "every run" is corrected to *the tick the latch flips* *[EXP-0125]: the `+r` term is now exact rather than assumed. The divisor `[EBX]` is `AImgr+0x00` = `0x8000` (`AI-RANGE-102`), so `3 × rand() / 0x8000 − 1` is `{−1, 0, +1}` in thirds, and the empty-members edge is the identical sequence without the `+4` (`AI-JITTER-103`).*

**Confidence.** High for the two branches and every store (one routine read end to end, and the member-count test at `L00438`/`L00443` is what separates the two edges). **Medium** that the roll is once per flip in play — `grpAI+0x30` was not swept for writers outside this routine, and the AI record is reachable as `grp+0x3c` so a displacement sweep on `0x30` would not be the instrument

**Amended.** The `+r` term is exact rather than assumed: the divisor is `AImgr+0x00` = `0x8000` (`AI-RANGE-102`), so `3 × rand() / 0x8000 − 1` is `{−1, 0, +1}` in thirds, and the empty-members edge is the same sequence without the `+4` (`AI-JITTER-103`). The row also corrects `AI-RADIUS-014`'s "every run" to the tick the latch flips.

## Group orders 2, 4, 5 and 0x11, withdrawal and the order writers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-SWARM-022 | Group order 2 (`Swarm`) is *walk to the commanded cell and fight what you see*. | High | ● active | **[EXP-0087](../experiments/EXP-0087-group-arms/)** |
| AI-MOVE-023 | Group order 4 (`Move`) walks each member to its *own* `ord+0x0a` and re-acquires on arrival. | High | ● active (superseded) | **[EXP-0087](../experiments/EXP-0087-group-arms/)**, [EXP-0386](../experiments/EXP-0386-engagement-under-occupancy/EXP-0386.md) |
| AI-SWARM2-024 | Group order 5 (`Swarm 2`) is `Swarm` minus the walk, and falls back to `Move`. | High / Medium | ● active (amended, superseded) | **[EXP-0087](../experiments/EXP-0087-group-arms/)** |
| AI-ROAM-025 | Group order `0x11` (`Roam`) updates a Group-AI cell and counter, then invokes Swarm 2 evaluation; the measured corpus contains no authored Roam nodes. | High / Medium | ● active (amended) | **[EXP-0087](../experiments/EXP-0087-group-arms/)** |
| AI-WITHDRAW-026 | Every group member is tested for withdrawal once per full tick, after the order dispatch, whatever the order. | High | ● active | **[EXP-0087](../experiments/EXP-0087-group-arms/)** |
| AI-WITHDRAW-027 | The two thresholds are the `Data.bin` Units columns the engine names itself — `Withdraw` and `Wimpy` — and only ranged monster classes carry a `Withdraw`. | High | ● active | **[EXP-0087](../experiments/EXP-0087-group-arms/)** |
| AI-WITHDRAW-028 | The flee cell uses a distance-3 away projection without a passability test in the picker. | High / Unknown | ● active (amended, partially retracted) | [EXP-0087](../experiments/EXP-0087-group-arms/), [EXP-0272](../experiments/EXP-0272-immediate-retreat/) |
| AI-CLASS-029 | A spawn-time classifier exists that would overwrite both thresholds, and it never writes anything. | High | ● active | **[EXP-0087](../experiments/EXP-0087-group-arms/)** |
| AI-CLASS-030 | A command overwrites both thresholds at runtime, and its per-class table is unreachable. | High / Medium | ● active | **[EXP-0087](../experiments/EXP-0087-group-arms/)** |
| AI-ORDER-031 | The complete immediate-form writer set of `grpAI+0x20`, with its instrument and both blind spots. | High / Medium | ● active (amended) | **[EXP-0087](../experiments/EXP-0087-group-arms/)**, [EXP-0322](../experiments/EXP-0322-group-order-producers/) |

### AI-SWARM-022

`R0025(grp, cell)` — the cell is the dispatcher's `grpAI+0x0a` (`L00306`). It calls `R0110(grp)` (build the candidate list) then `R0177(grp)` (score, winner into each member's `ord+0x20`), then per member: `L00444` target non-null → `R0009(member, target)`; else `L00445` compares the current cell with the commanded one and `L00446` calls `R0040` for idle — not both → `L00447` stores 1 into `ord+0x08` and `L00448` stores the commanded cell into `ord+0x0a`, an ordinary move order to that cell; both → `L00449` on `[member+0x14]+0x28` (`UNIT-OWNER-009`): non-zero (scenario owner) → `L00450` stores `0xb` into `ord+0x08` and, if `actor+0x4c` bit 2, `R0113`; zero (human participant) → `R0015`, the heal AI. **No formation, no spread, no per-member offset**

**Confidence.** High — one routine read end to end; every store cited; the entry the arm is reached by is the PE-read table `L00303[2]` = `L00306`

### AI-MOVE-023

`R0154(grp, ignored)` — it takes a second argument and **never reads it**; the destination reaches the members from the setter `R0108`, not from here. `L00451` recomputes the notice radius (`R0140`). Per member: arrived (the compare at `L00452` against the word at `+0xa`) **and** idle **and** `ord+0x50 == 0` → `L00453` stores `actor+0x12c` into `ord+0x14` (reach becomes the stop distance), `R0007` × 1, `R0088`, `ord+0x50 = 1`, route list `pth+0x90` cleared; else if the route list is non-empty and `pth+0x76` equals the current cell → three `R0007` calls, `L00454` stores `0xc` into `+0x50` (actor state 0xc), `ord+0x00` = current cell, `ord+0x08 = 0`; then `L00455`: `ord+0x50 != 0` → `R0022` (acquisition), else `grpAI+0x44 != 0` → `ord+0x6c = 1` and the move is re-issued (`L00456`)

**Confidence.** High — one routine read end to end; the unread second argument is a bounded statement about *this* routine, not about the order

**Amended.** Gloss superseded. The `mover+0x90` gloss, at both sites in its own prose: the first arrived branch's "route list `pth+0x90` cleared," and the `else if` branch's "the route list is non-empty and `pth+0x76` equals the current cell": `mover+0x90` is a plain boolean, not a list or a count, set to 1 only by `R0178` (`L00457`), and cleared to 0 at three sites this row's own routine and this round's own evidence together identify: `L00458`, inside the first arrived branch this row's own prose cites as "route list `pth+0x90` cleared"… (superseded, EXP-0386). `retracted.md` holds the full entries.

### AI-SWARM2-024

`R0026(grp, cell)`: `L00459 R0110(grp)`, then `L00460` reads the dword at `+0xbb4` and `L00461` branches to `L00462` — on zero it **tail-calls `R0154` with both arguments forwarded** (`L00463`). Otherwise `R0177(grp)` and the same per-member body as `AI-SWARM-022` **without** the move-to-the-cell branch: engage `ord+0x20`, or `ord+0x08 = 0xb` + cast for a scenario-owned mage, or the heal AI for a human participant's unit

**Confidence.** High for the body and the branch (one routine, cited stores). **Medium** that the fallback is or is not taken in play: what fills `AImanager+0xbb4` is not established — `disp:bb4` returns 15 hits / 13 owners, the ten in the AI module all reads and the five writes on a different object in the code-region module, which is consistent with a member of the inlined `CObList` at `+0xba4` (head at `+0xba8`) rather than an assigned flag

**Amended.** The Medium clause is closed by `AI-CANDCOUNT-105`; the field is the candidate collection's element count, the gate is "did the build find anything", and the fallback is the ordinary path. Two sweep figures in that clause are corrected in `retracted.md`. The Medium clause only — the body and the branch are untouched: The field is the **element count** of the candidate collection at `AImanager+0xba8`, so the gate is *"did the build that just ran find a candidate"* and the fallback is the ordinary path, not a mode (superseded, EXP-0126). `retracted.md` holds the full entries.

### AI-ROAM-025

`R0155(grp)`, reached by the out-of-table arm the branch at `L00464` to `L00310`. Per tick it takes the **maximum** over members of `R0167(memberCell, grpAI+0x0a)`, the Chebyshev distance (`L00465`, `L00466`). `L00467` compares the maximum with 10 — **max < 10** → re-roll; else `L00468` compares the byte at `+0x15` with `0x32` (unsigned, at-or-below keeps) — the group counter `grpAI+0x15` **> 50** → re-roll; otherwise keep the destination. The roll: `L00469` calls `rand` (`R0179`), shifts the result left by 3 and divides by the dword at the divisor slot → 0..7; `L00470` / `L00471` index two signed byte tables at `world+0x58eb0` and `world+0x58eb8`, which the world init `R0116` fills in code as `dx = {0,1,1,1,0,-1,-1,-1}` / `dy = {-1,-1,0,1,1,1,0,-1}` (`L00472`…`L00473`); `L00474`…`L00475` compute `(X + 20·dx, Y + 20·dy)` from `grpAI+0x0a`/`+0x0b`; `L00476`…`L00477` reject and re-roll unless the result lies inside the rectangle `world+0x58ee0..+0x58ee3` (the byte store of 8 at `L00478` into `world+0x58ee0`).

On acceptance `L00479` stores the accepted cell into the word at `+0xa` and `L00480` stores 0 into the byte at `+0x15`; then `L00481` calls `Swarm 2` evaluation `R0026` and `L00482` increments the counter after normal return. This is not a command-setter call. `AI-SWARM2GATE-107` forwards the cell only on its zero-candidate fallback to `AI-MOVE-023`, whose local body ignores that argument and uses current actor `ord+0x0a`. These bodies do not establish a producer copying the generated Group-AI cell into actor destinations. Reached movement can depend on member order and called helpers; neither guaranteed wandering motion nor absence of all other wander paths follows. The former genuine-wander/only-one headline is narrowed in `retracted.md`. **`Par0 = 17` ships 0 nodes over 38 maps**

**Confidence.** High for the local Group-AI cell/counter program and evaluation call, not a composed movement result — one routine read end to end, the two byte tables read from their filler rather than from memory, the dispatch entry from the PE. High for the corpus 0 (a walk of every shipped map's type-7). **Medium** that the loop terminates for a group cornered against the rectangle: the re-roll has no attempt bound

**Amended.** Genuine-wander/only-one headline and "executes it" wording only: The local program generates a Group-AI cell, resets/increments its counter and calls Swarm2 evaluation (narrowed, EXP-0087, EXP-0126). `retracted.md` holds the full entries.

### AI-WITHDRAW-026

The tail of `R0023` at `L00483` is reached from the jump table's fall-through, from the `0xff` arm and from the end of every case. Per member: `R0105` — `L00484` reads the signed current health (word at `+0x94`) and compares it with the dword at `ord+0x44`; proceed only when `health <= ord+0x44`, then `R0020(cell, pth+0x08)` and `R0021(actor, coll, 0)` keep the hostiles; non-empty → `R0100`. Else `R0106`, the same with `ord+0x40` and the **literal radius 2** (the constant 2 pushed at `L00485`) → `R0107`. `R0076` calls the dispatcher once per group per full tick for **every** player, gated only on `session+0xb388` or `grpAI+0x45`

**Confidence.** High — both gates and the tail read end to end; the tail's reachability is a control-flow fact of one routine, and the branch at `L00486` to `L00487` (short displacement 7) fixes the boundary

### AI-WITHDRAW-027

`R0180` streams nine consecutive slots whose titles read `typeID face tokenSize movementType dyingTime Withdraw Wimpy "See invisible" XPvalue` into `actor+0x0e`, `+0x4b`, `+0x49`, `+0x4a`, a local, **`ord+0x40`** (`L00488`), **`ord+0x44`** (`L00489`), `ord+0x71`, `actor+0x1c`. Corpus, both roots identical: 56 parameterised Units rows, **12 with `Withdraw` > 0 and they are exactly `Goblin_Sling` ×4 (30), `Orc_Bow` ×4 (60), `Bat_Sonic` ×4 (15)** — every one a ranged `EquipItem`, no melee class anywhere; 24 with `Wimpy` > 0 (`Goblin_Pike` 10, `Goblin_Sling` 8, `Bee` 3, `Squirrel` 5, `Bat_Sonic` 4, `Dragon` 63). The threshold is **absolute health**, and the four tiers share the tier-1 value while `healthMax` scales ×1.6 per tier, so a tier-1 ranged monster satisfies `health <= Withdraw` **always** and tiers 2/3/4 below 62 % / 39 % / 24 %. The **Humans** table carries neither column, so a hero and a mercenary never withdraw. Shipped placements: EN 1527 of 8094 can withdraw, RU 817 of 3991

**Confidence.** High for the binding (four of the nine slots are independently published by `TERR-MOVE-054`, `UNIT-STREAM-001` and `HERO-XP-010`, so a one-slot misalignment breaks all four) and for the census (a walk of every shipped map on two roots; a withdrawing melee class would have shown as one)

### AI-WITHDRAW-028

The former living-only centroid and identical positive-HP gate for `R0100`/`R0107` are retracted by `AI-RETREAT-274`: the first counts positive HP only to choose whether to proceed, then averages the entire filtered list; the second uses a nonempty-list gate without that scan. Both divide coordinate sums by the low byte of list count. The surviving geometry is `R0104(self, centroid, 3)`: fine-coordinate differences of zero become 1; the larger axis advances by signed `3*256`, the other is projected through `__ftol`; coordinates shift right 8 and clamp to `[8, dimension-9]`. The caller writes pending move 1 and its cell. No block-plane read occurs in the picker; route handling is separate (`MOVE-ALT-019`).

**Confidence.** High for the retained geometry and corrected helper dataflow. Unknown whether mixed/all-dead filtered lists occur in ordinary runtime occupancy.

**Amended.** Living-only centroid and shared positive-HP gate only; distance-3 geometry and picker passability result stand: `R0100` counts positive HP only for an entry gate, then sums the whole filtered list and divides by its low-byte count (refuted, EXP-0272). `retracted.md` holds the full entries.

### AI-CLASS-029

`R0181(actor)` reads `actor->vt+0x30()`; non-zero and a spellbook → `ord+0x4c = 2`, `ord+0x40 = healthMax`, `ord+0x44 = healthMax/4`; non-zero, reach > 1 and `vt+0x20() != 3` → `ord+0x4c = 1`, `ord+0x40 = healthMax/2`, `ord+0x44 = 0`; non-zero otherwise → all three zeroed; **zero → the branch at `L00490` to `L00491`, which is the epilogue: it returns writing nothing.** `vt+0x30` is `R0182` (returns 0) in the base actor vtable `L00001` and `R0183` (returns 1) in `Humanoid` `L00002` and `Human` `L00003`. Reachability: `callto:R0181` → **1** (`L00492`, the last act of `R0184`); `callto:R0184` → **1** (`L00493` in `R0185`); `callto:R0185` → **5**, and every one of the five installs `L00001` **before** the call (`L00494`<`L00495`, `L00496`<`L00497`, `L00498`<`L00499`, `L00500`<`L00501`, `L00502`<`L00503`), while both derived classes install their own vptr only after the base returns (`L00504`<`L00505`, `L00506`<`L00507`). **So on every construction path the classifier sees the base slot and takes its first exit**, which is what lets the streamed `Withdraw`/`Wimpy` survive

**Confidence.** High — three `callto:` enumerations on a repaired function table with 0 orphan hits, plus five address comparisons inside single routines. The instrument's blind spot is an indirect call through a pointer no `.rdata` slot holds; the whole chain here is `CALL rel32`

### AI-CLASS-030

`R0186(a0,a1,a2,player)` and `R0187(…)` walk `world+0xa4554` and set `ord+0x40` / `ord+0x44` to `healthMax × pct / 100` (`0x51eb851f` × / `SAR ,5`), choosing `pct` by `ord+0x4c` — 0 → a0, 1 → a1, 2 → a2, `> 2` → the byte is forced to 0. `player == 0` applies to every actor whose `Player+0x28` is 0; otherwise to that player's actors. Their single caller each is one sub-arm of `R0061` (`L00508`, `L00509`) that reads a mode from `cmd+0x0e` and passes `(v, 2v, 3v)` and `(0, v, v)` with **`v = 0 / 10 / 30`**. Because `AI-CLASS-029` leaves `ord+0x4c` at 0 on every path, the 2v and 3v columns are never selected: the command's effect is uniform — mode 0 disables the withdraw, mode 1 sets it to 10 % of maximum health, mode 2 to 30 %, and all three overwrite the authored `Wimpy`

**Confidence.** High for both routines and the three modes (read end to end, one caller each). **Medium** for the *name* of the command: the outer opcode was not traced past the sub-dispatch at `L00510`

### AI-ORDER-031

An instruction-pattern sweep for byte stores of an immediate into `+0x20` returns **75 hits / 32 owners / 1 in orphan code**; the group-order members are: `0xff` `L00300` (`R0023`, empty group) · `0` `L00511`/`L00512`, `4` `L00433` (`R0108`) · `0` `L00513`/`L00514`, `5` `L00434` (`R0157`) · `0x11` `L00515` (`R0158`, **`callto:R0158` → 0**) · `1` `L00314` (`R0125`) · `3` `L00143` (`R0060`) · `0` `L00516` (`R0164`) · `3` `L00153` (`R0063`) · **`2` `L00318`, inside `ORPHAN[L00316..L00317, vtPtrs=0]`** · `2` `L00431`, `3` `L00315`, `0x11` `L00432` (`R0156`). Blind spot 1: the sweep is on **immediates** and cannot see the store of a register-held zero into `+0x20` at `L00412` (`AI-PATROL-018`).

Blind spot 2: the orphan hit has **no containing function**, `callto:L00316` → 0, `refto:L00316` → 1 DATA reference at `L00517`, and `range:L13125:L13126` reports 453 orphan bytes with 177 undisassembled before them — so ~~its entry point and reachability are **not established**~~ — *amended by [EXP-0322]: the entry point is established, and reachability is not thereby restored. `L00518` opens by saving registers and loading the global singleton at `L00285`, with no conventional frame setup, and reads a global singleton's own `+0x11c` (must be `0`) and `+0x80` (`5` exits to `L00317`, `6` continues, anything else exits to `L00519`), then dispatches a **10-way table at `L00520`** — the same base this row's own `refto:L00316` reference already names — on a counter at `this+0xbcb0`. `L00318` is reached only from table entry **2** (`L00521`; `evidence/q1-second-order2-jumptable.txt` captures the unbiased table at `L00520`), which logs a message (`CALL R0188`, string `L00522`) and increments its own selector before falling into a member-clear loop (zeroing each list member's `+0x158+0x50`/`+0x38`) ending in the cited `0`-then-`2` write and a literal cell `0x1824` (`L00523`/`L00318`/`L00524`). Table entry **1** (`L00316`, the address this row cites — the dword at `L00517`, this row's own `refto:` hit) does not itself reach `L00318`: it tail-jumps, with a pushed order value `2`, to a shared tail block at `L00525`, reached by fall-through from the arm above it as well as by that jump and not a separate function, whose effects this experiment did not trace. `xcallers L00518` and `xrefs L00518:L00526` are both 0 — the same absence this row already reports against `L00316`, now confirmed against the routine's real entry. Table entries 0 and 3-9 and the effects of the shared tail block at `L00525` remain untraced: whether they hold further, uncited `+0x20` writers is Unknown. A third, independent, equally unreachable order-2 producer at a wholly separate address is `AI-ORDER-294`.*

**Confidence.** **Medium** as a completeness claim, for the two named reasons. High for each listed site individually

**Amended.** The orphan hit's "entry point and reachability are not established" clause only; the 75-hit/32-owner count, the blind-spot-1 register-form note and every other listed site stand: The containing function's entry is `L00518`, not the jump-table target `L00316` the original hit cited: a distinct prologue… (narrowed, EXP-0322). `retracted.md` holds the full entries.

## Player order vocabulary and order progress

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-CMD-032 | The player's order vocabulary is nineteen slots of which four are empty, and it is a different table from the session's. | High / Medium | ● active | [EXP-0090](../experiments/EXP-0090-commands/) |
| AI-CMD-033 | Every player order builds a brand-new group at group order 0 — so `grpAI+0x20 == 0` is not a precondition the player must arrange, it is what the dispatcher establishes… | High / Medium | ● active (amended) | [EXP-0090](../experiments/EXP-0090-commands/) |
| AI-PROGRESS-034 | `R0016` is the per-actor order-progress routine, it has two switches, and it is the bridge from the order object to the actor's act state. | High / Medium | ● active (amended, superseded) | [EXP-0090](../experiments/EXP-0090-commands/), [EXP-0179](../experiments/EXP-0179-effect-action/) |
| AI-DEAD-035 | Two routines sweep every actor in the world with no group test, and neither is reachable — positively, not for want of looking. | High | ● active | [EXP-0090](../experiments/EXP-0090-commands/) |

### AI-CMD-032

`R0061` selects on `cmd+0x04`: `>= 1` is the order space (`L00527`: opcode `cmd+0x09`, biased by `0x14`, bounded at `0x12`, above which it goes to `L00528`, then **19 direct dwords at `L00529`**), `== 0` the session space (`SESS-CMD-016`). The two overlap on the opcode byte and are otherwise unrelated — `0x22` is empty here and the item move there. Every arm opens by testing the AI manager global `[L00004]` and does nothing when it is null. **The table's *shape* was already published** — `SESS-CMD-008` gives both bounds and `ITEM-CMD-007` the 19 dwords at `L00529`; what neither has is a single one of its arms, and that is what follows. Read out of the PE by `tools/cmdtable`, with what each arm dispatches to: `0x14 → R0103`; **`0x15` EMPTY**; `0x16 → R0108(grp, col, row)`; `0x17 → R0125(grp, 0)` guard; `0x18 → R0060` aggressive; `0x19 → R0006(grp, target)`; `0x1a → R0189 → R0157(grp, col, row)`; `0x1b → R0190(grp, target, 0)`; `0x1c` the engine's own `"defend location comes"` then `R0108`; **`0x1d → R0174` patrol** (`AI-PATROL-019`); `0x1e → R0011(grp, target, spell)`; `0x1f → R0012(grp, col, row, spell)`; **`0x20` EMPTY**; **`0x21` the sack pick-up** (`ITEM-PICK-016`); **`0x22` EMPTY**; **`0x23` EMPTY**; `0x24 → R0085(actor, building)`, refusing with `"No building #"`; `0x25` and `0x26` **share** `L00038`, which takes item `cmd+0x10` out of `actor+0x7c`, requires an effect of kind `0x29` on it, parks the spell in `actor+0x44` and then routes to `R0011` (at a unit) or `R0012` (at a cell) on the opcode.

The spell of `0x1e`/`0x1f` is remapped first through a **stride-4 byte table at `L00198`** indexed by `cmd+0x10` (`L00197`), whose contents are not decoded here. **Nine of the fourteen live arms have the dispatcher as their only caller in the image** — `callto:` on each, 0 in orphan code: `R0174` 1/1, `R0103` 1/1, `R0006` 1/1, `R0189` 1/1, `R0190` 1/1, `R0011` 2/1, `R0012` 2/1, `R0010` 1/1, `R0085` 1/1 — while three targets are shared with the script module's `R0156` (`R0108` 3/2, `R0157` 2/2, `R0125` 6/6), and the script **never enters the dispatcher**, whose whole caller chain is `R0191` ← `R0192`/`R0193`. **That is a statement about the ROUTINES and it does not make the behaviours player-only** — a player-only routine can reach a shared behaviour one level down, and patrol is exactly that: `R0174` has one caller, but it tail-calls `R0172`, whose `callto:` is **2 / 2 / 0** — this arm *and* `R0156`, which is `AI-PATROL-017`'s script route.

The **pick-up** does survive the same test as a *behaviour*: the complete immediate-form writer set of `actor+0x50 = 2` is **3 hits / 3 owners / 0 orphan** — `L00022` in `R0010`, `L00530` in the dispatcher's own session-space `0x22` arm, and `L00531` in `R0194`, which is in the shop module and is a different class's `+0x50` — so **both writers that reach an actor are inside `R0061`'s two spaces**. **Customisation: the four empty slots are absent implementation and free** — nothing in a shipped `.alm`, `.res` or save carries an order opcode — **while the opcode's width (one byte at `cmd+0x09`) and the member list (`u8` count at `cmd+0x12`, `u16` ids from `cmd+0x13`) are the multiplayer wire and are a coupling**, with a 255-member ceiling on one order

**Confidence.** High (the table is read from the PE through the section table and printed beside its raw bytes; every arm's call, argument list and guard is a named instruction; six opcodes published from other areas land on their known arms) / **Medium** for the *labels* of the arms not read past their entry — `0x14`, `0x19`, `0x1a`, `0x1b`, `0x1e`, `0x1f`, `0x24` are named by the routine they call and by their final state store, not by a reading of the body

### AI-CMD-033

Every player order builds a brand-new group at group order 0 — so `grpAI+0x20 == 0` is not a precondition the player must arrange, it is what the dispatcher establishes — and the orders then divide into two families by what they leave behind. The order space's prologue resolves the commanded actors (`R0195` per id), calls `R0196(Player+0x24)` — which scans the player's group list until the first group failing `R0197`, deletes that one and exits the scan, then allocates `0x48` bytes with `operator new` and constructs with `R0149` — adds each actor with `R0152`, and appends the group with `R0198` (`L00532`…`L00533`). The group's AI block is constructed by `R0150`, whose store at `L00534` writes **0** into `+0x20`, the zero having been produced at `L00535`.

From each order routine's own final store: **guard leaves 1** (`L00314`), **aggressive 3** (`L00143`), **move 4** (`L00433`, orders `0x16`/`0x1c`), **`0x1a` 5** (`L00434`) — those four act through `R0023`'s group-order arm — while **ten live arms leave 0** and act by writing a per-actor `actor+0x50` state that only `R0023` arm 0 evaluates: `0x14 → 0x16`, `0x19 → 3` then `0xc`, `0x1b → 0xc`/`8`, `0x1d → 0xa` patrol (through `R0172`, `L00414`), `0x1e → 0xd`/`0xc`, `0x1f → 0xe`/`0xc`, `0x21 → 2`, `0x24 → 0xf`. `R0010` and `R0085` contain **no** `+0x20` store at all, so their group stays at the constructed 0. **This is `AI-ORDER-010`'s gate seen from the player's end**: `AI-AUTHOR-015`'s "every group on a shipped single-player map is set to order 1 or 3 at load" is about **authored** groups and says nothing about the group a command creates

**Confidence.** High (the allocation chain, the constructor's zero and each routine's final `+0x20` and `actor+0x50` stores are all named instructions, and the stored register is zero in every routine quoted, each zeroing it at its own head) / **Medium** that the family split is complete: it rests on each routine's *last* `+0x20` store in program order, which cannot see a store made by a callee this experiment did not read

**Amended.** Cleanup scope, see `retracted.md` and `SAV-GRPCMD-578`. Every-rejected-Group cleanup clause only: A rejected old Group is removed/deleted, then `L00536` jumps directly to allocation at `L00537`; at most the first rejected Group is deleted in that invocation (narrowed, EXP-0291). `retracted.md` holds the full entries.

### AI-PROGRESS-034

The actor tick calls it four instructions before it dispatches on `actor+0x54` (the call to `R0016` at `L00086`, then the read of `actor+0x54` at `L00538`), and no experiment had read the call. It first **preserves** `actor+0x54` when it is 2 or `0xf` and clears it otherwise (`L00539`…`L00540`, with the comparison value 2 set at `L00541`). It then dispatches on the **progress** byte `ord+0x09` — decremented, bounded at `0xfe` (above which it goes to `L00542`), then 255 index bytes at `L00090` into **6** dwords at `L00091`, of which only `1..4` are live (`L00092`, `L00093`, `L00096`, `L00543`) *(FIVE are live: [EXP-0179] reads the index table out of the PE and `0xff` maps to its own arm `L00544`, written at `L00545` and cleared only by `R0146`. The same routine also writes progress 4 at `L00546`, eleven instructions before the dispatch this row describes, from `actor+0x144 & 0x100000` — `MAGIC-ACTGATE-079`; see [`retracted.md`](retracted.md))* — and, when that byte is 0 (the branch at `L00547` to `L00548`), on the **pending order** byte `ord+0x08`: decremented, bounded at `0x0e`, then **15 direct dwords at `L00099`**, with arms `3`, `0xd` and `0xe` EMPTY.

**Arm 7 is the pick-up bridge** (`L00549`): `actor+0x54 = 2` on the first pass and, on the second — when the tick has already run its own arm 2 — the completion `actor+0x50 = 0xc`, `ord+0x08 = 0`, `ord+0x50 = 1`. Both bytes are **inside the order object's serialized `0x94`-byte raw block** (`R0199`: `ar.Write(this, 0x94)` at `L00550`, `ar.Read(this, 0x94)` at `L00551`), so a consumer that inserts a field before them cannot load a save the original wrote — **and because a wholesale block move carries no displacement, no store-form sweep over `+0x08`/`+0x09` can see the save path at all**. Adding a *value* inside `1..0x0f` is free; adding a *field* is not

**Confidence.** High for both switches, the preserve test and the arm-7 bridge (the two range tests, both table bases and the arm-7 body are named instructions; the tables are read from the PE) / **Medium** that arms `1..6` and `8..0xf` are what their calls suggest — they were classified by their calls and their `actor+0x54` stores, not read line by line

**Amended.** The progress-arm liveness clause only — both switches, the range tests, the table bases, the act-state preserve and the arm-7 bridge stand: The index table, read out of the PE by virtual address, maps `ord+0x09 == 0xff` to its own arm at `L00544`, distinct from the default `L00542`, and `0xff` is written to `ord+0x09` at `L00545` — inside `R0016` itself, on the tail's `actor+0x50 == 0x17` arm… (superseded, EXP-0179). `retracted.md` holds the full entries.

### AI-DEAD-035

`R0148` and `R0038` are the same loop twice: both walk `[[AImanager+0xa50]+0xa4554]+4`, the world's whole actor list, and call `R0008` (the per-actor AI state machine) / `R0016` (the order-progress routine) on every member, with **no `grpAI+0x20` test anywhere in either**. Were either live, every actor on the map would evaluate its own state regardless of its group's order — which is precisely the shape that would make the per-actor states unconditional and dissolve `AI-ORDER-010`'s gate from the other side. For both: `callto:` **0 hits**, `refto:` **0 hits**, `imm:<address>` **0 hits**. The `imm:` leg is what upgrades this past `callto:`'s known blind spot — a call through a pointer another routine loaded still needs the address materialised by some instruction, and it is not. The live callers are one each: `R0023` arm 0 for the state machine (`callto:R0008` = 2 / 2 / 0, the other hit being the dead sweep) and the actor tick `R0037` for the progress routine (`callto:R0016` = 2 / 2 / 0, likewise)

**Confidence.** High (three independent enumerations over the repaired table, each returning zero, and the alternative they rule out is named)

## Dying actors and the formation mode

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-DEAD-036 | `R0200` — the only routine in the image that clears the group rate byte — is unreachable, and the fourth instrument is what settles it. | High | ● active | [EXP-0094](../experiments/EXP-0094-group-rate/) |
| AI-FORM-037 | Every `Player` carries a formation mode, it is one byte with three behaviours, and both of the things that can write it are authored surfaces. | High / Medium | ● active (amended, partially retracted) | [EXP-0094](../experiments/EXP-0094-group-rate/) |
| AI-SPREAD-038 | The formation spread threshold is a compile-time `2`, it lives on the AI manager and not on the world, and no shipped byte carries it. | High / Medium | ● active | [EXP-0094](../experiments/EXP-0094-group-rate/) |

### AI-DEAD-036

`EnumRefs callto:R0200` **0 hits**, `refto:R0200` **0 hits**, `imm:R0200` **0 hits** — the same three `AI-DEAD-035` uses. Those three are not enough here, because **every dispatch table in this area lives in `.text`**, outside `callto:`'s `.rdata` window `L00552..L00553`, and a table slot is neither a reference nor an operand: `R0154`, which certainly runs, is invisible to all three for exactly that reason. So the fourth leg is a scan of the whole image for the little-endian dword `R0200`, read through the PE's own section table (`tools/grouprate -mode scan`), and it returns **nothing** — while its three **controls**, the known jump-table slots `L00304`, `L00554` and `L00555`, are each found at their table (`L00303`, `L00556`, `L00557`).

Nor is it fallen into: the preceding function ends with a 3-byte return that pops one stack argument at `L00558` and five `0x90` pad bytes separate it from the first instruction of `R0200` (`evidence/cited-bytes.txt`), and `EnumRefs range:R0176:L00559` reports 5 functions, **0 orphan bytes** and 32 undisassembled — the padding. Two live neighbours were run through the Ghidra instruments as controls and both came back non-zero (`callto:R0171` 2/1/0, `callto:R0174` 1/1/0), so the zeros are the routine's and not the tool's. Consequence: `MOVE-GROUP-030`'s "cleared to 0 when the group's pass ends" is retracted, and nothing clears `grpAI+0x44` (`MOVE-GROUP-037`)

**Confidence.** High (four independent enumerations, one of them immune to Ghidra's function table, each returning zero, with positive controls on both the instrument and the neighbourhood)

### AI-FORM-037

The `Player` constructor `R0201` allocates a `0x20`-byte settings block (the size pushed at `L00560`), constructs it with `R0202` — which zeroes it and writes **one** nonzero default, the store of 2 into the byte at `+0x1f` at `L00561` — and stores it at `Player+0x30` (`L00562`); its `CreateObject` thunk `R0203` allocates `0x70`, the `Player` runtime-class size, so the object is identified rather than assumed. The byte's whole surface is four instructions: `EnumRefs "re:.*\+ 0x1f\].*"` gives **8 hits / 8 owners / 0 orphan** image-wide, four of them `[ESP+0x1f]`, and the rest are that default, the single setter `R0080` (`L00563`), and the two reads in the group-move setters. Behaviour, from `MOVE-GATE-035`: **`0`** never in formation, **`2`** in formation iff the spread test passes, **any other nonzero** in formation unconditionally — so 254 of 256 values alias to one behaviour.

Two callers, both authored surfaces and both read out of the PE rather than off a listing: **player command opcode `0x46`, sub-code 2** (the second command dispatch's index byte table at `L00564` sends `0x46` to slot 21 = `L00565`; the player comes from `cmd+0x05` through `R0079`, the value from `cmd+0x0e` **remapped** `0→0, 1→2, 2→1`, default 2), and **trigger instant id 7**, which the shipped `Description Instants.ini` `[Group6]` names **`Set formation`** (instant table `L00566` slot 6 = `L00554`) and which writes `rec+0x08` **raw**, with no remapping. It is in a save: `Player::Serialize` hands the block to `R0204`, a raw `CArchive::Write`/`::Read` of `0x20` bytes (`L00567`/`L00568`).

Corpus, both roots: **one** `Set formation` node ships in the entire 38-map corpus — `scn:110.alm`, node id 29. ~~`Par0 = 2`, every other `Par` zero — i.e. the constructor's own default~~ *— amended by [EXP-0170]: the node's slot 0 is the **`Player` reference**, type 3, value 2; its slot 1 is the `Formation` int, type 1, value **0**. Under `TRIG-PARAM-030`, published after this row, a reference-typed slot takes no `p` index, so the arm's `p0` is slot 1 and the byte written to `+0x1f` is **0** — `MOVE-GATE-035`'s *never in formation*, not the default. The node's own author label, identical on both roots, is `"D - brigands dont use formations"`. The value recorded here was the player index.*

**Confidence.** High (the allocation, the ctor default, the setter, the two dispatch tables and the serializer are cited instructions and PE reads, and the field's surface is a complete image-wide enumeration; two published opcodes on the same table, `SHOP-MISSION-018`'s `0x3f` and `ITEM-CMD-007`'s `0x22`, land on their claimed arms) / **Medium** (that a shipped map's one node *means* what the catalogue's `Par0_NAME=Player` says: the runtime takes the value from `Par0` and the object from the resolved-target slot `rec+0x38`, and the compiled record's slot-to-parameter mapping past `rec+0x2c` was not derived) — *that Medium was correct and also [EXP-0170] resolves it against the row's own corpus reading*

**Amended.** The corpus clause only — the allocation, the constructor default, the setter, the two dispatch tables, the serializer and the three-behaviour reading of the byte are untouched and reproduce exactly: The node's slot 0 is the `Player` reference, type 3, value 2, and its slot 1 is the `Formation` int, type 1, value **0** (refuted, EXP-0170). `retracted.md` holds the full entries.

### AI-SPREAD-038

`EnumRefs disp:a824` → **3 hits / 3 owners / 0 orphan**: two compares (the two group-move setters) and exactly one write, the byte store at `L00569` inside `R0124`. The stored value is 2, loaded at `L00570` and not written again before the store; the same register still holds 2 at `L00571`, twenty-one instructions later. `R0124` is the constructor of the object whose `+0xa9c4` it `REP STOSD`s with `0x271` dwords, i.e. `AI-DIPLO-004`'s 50×50 diplomacy matrix, which is what fixes the object as the **AI manager**; `AImanager+0xa50` is the *world* (`AI-DEAD-035`), so `MOVE-GROUP-030`'s `world+0xa824` names the wrong object and a consumer looking for the field there would not find it.

What it bounds is a **Chebyshev distance in whole cells from a member to the group centroid**, compared `JBE`, so a group of two units four cells apart passes and five cells apart fails. Corpus, at the **shipped placement positions** (`tools/grouprate -mode groups`, replaying `AI-GROUP-009`'s grouping over 38 EN maps): **1255 of 2264** groups satisfy it, and of the **1578** with more than one member, **569** (36.1 %); 686 groups are singletons. The RU root's `rom.exe`-derived evidence is byte-identical

**Confidence.** High (a complete displacement enumeration with one writer, the value fixed by a register that no intervening instruction touches, and the owning object fixed by a second field of the same constructor that another claim already identifies) / **Medium** (the corpus figures: they measure where the *author* put the units, not where they stand when a script command fires)

## The per-actor order machine and pursuit

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-ORDER-039 | The per-actor *order* machine is `ord+0x08`, its fifteen-slot switch has twelve live arms, and all twelve are now read rather than classified by their calls. | High / Medium | ● active (amended, partially retracted) | [EXP-0099](../experiments/EXP-0099-actor-states/) |
| AI-PURSUE-040 | A pursuit is an order object, not a state: `ord+0x08` in {5,6}, `ord+0x0c` the target, `ord+0x14` the stop distance — and the in-position test is *facing plus edge-to-edge distance*, with reach re-read live. | High | ● active | [EXP-0099](../experiments/EXP-0099-actor-states/) |
| AI-BREAK-041 | A pursuit can be broken off, the input is the distance from the creature's *post* to the *target*, and for the default state that distance is a hard-coded 5 cells Chebyshev — not sight, not reach, not the group radius. | High / Medium | ● active | [EXP-0099](../experiments/EXP-0099-actor-states/) |
| AI-POST-042 | `ord+0x00` is the post, it is emergent rather than authored, and for a guard it never moves. | High | ● active (amended, partially retracted, superseded) | [EXP-0099](../experiments/EXP-0099-actor-states/) |
| AI-STATE-043 | The complete written value set of `actor+0x50`, and the five live arms nothing writes. | Medium | ● active (amended) | [EXP-0099](../experiments/EXP-0099-actor-states/), [EXP-0385](../experiments/EXP-0385-dying-and-crossing-tick/) |
| AI-THREAT-044 | The engage routine consults a timed three-slot threat cache before it selects anything, which is why an abandoned target can be re-taken on the next tick. | High / Medium | ● active (amended, superseded) | [EXP-0099](../experiments/EXP-0099-actor-states/) |
| AI-ROUTE-045 | The second thing that ends a pursuit is the route search failing, and `mover+0x98` is how it says so — closing `claims/move.md`'s "read but not traced to an effect". | High | ● active | [EXP-0099](../experiments/EXP-0099-actor-states/) |

### AI-ORDER-039

`R0016` reaches it only while `ord+0x09` is 0 (`L00572`…`L00547`) — the tick the actor stands on a cell centre, `R0040` testing `pos+0x4 == pos+0x5 == 0x80` — and dispatches on the order value, decremented and bounded at `0x0e`, through 15 direct dwords at `L00099` (`L00098`…`L00573`). Read out of the PE by `tools/aistates`: **15 entries, 3 on the default `L00010`, 12 distinct live arms**. `1` walk to `ord+0x0a` (`R0178`, stop 0); `2` attack `ord+0x0c` with **no distance test at all** — `ord+0x09 = 1`, `actor+0x54 = 3`, `actor+0x5c = target`; `3` empty; `4` close on the actor `ord+0x18` with stop distance `ord+0x14` (`R0043`); **`5` and `6` the pursuit pair** (`AI-PURSUE-040`); `7` the sack pick-up bridge (`AI-PROGRESS-034`); `8` cast `ord+0x30` at the actor `ord+0x28`, `actor+0x54 = 0xd`; `9` cast at the cell `ord+0x3c`; `0xa` turn in place until `mover+0x01 == mover+0x00`, then `ord+0x08 = 0`; **`0xb` the idle turn** `R0205` — `mover+0x01 = mover+0x00 + 0x21`, one eighth of a circle, ~~the quotient added to it being 0 for any clock value because the divisor is the session's own first dword~~ — *amended by [EXP-0109]: there is no clock value. The dividend is `190 * R0179()` and `R0179` is `rand()` (`AI-RAND-058`), so the dividend reaches `6 225 730`, above every vtable pointer seen in this module (`L00574..L00575`), and an identically-zero quotient is **not** established. The same routine's own entry gate was also unstated: the arm re-picks a facing only when `ord+0x54` is non-zero — the flag a blow sets (`AI-RETAL-056`) — or when `rand() < 0xcd`, about one evaluation in 160. See [`retracted.md`](retracted.md);* `0xc` walk to `ord+0x0a` then run the `actor+0x54 = 2` action there; `0xd`/`0xe` empty; `0xf` reach the cell `ord+0x0a` within `ord+0x14`, then `actor+0x54 = 0xf` and `R0088`.

The **progress** byte `ord+0x09` is a second switch, 5 live arms over 255 index bytes at `L00090`: 1 and 2 share a tail that increments the attack counter `ord+0x15` and, once it exceeds 2 with `actor+0x136` set, returns progress to 0; 3 is the step-along-the-path arm; 4 and the default park `actor+0x54 = 0x1a` *[EXP-0125]: the arm-`0xb` divisor is pinned — `AImgr+0x00` = `0x8000` — so the quotient is uniform on **0…189** and the stored facing is `mv+0x00 + 0x21 + U(0…189)` truncated to a byte. The row's remaining `0x21` = one-eighth wording describes only that term (`AI-TURN-104`, `AI-RANGE-102`).*

**Confidence.** High (both range tests, both table bases and every arm body are cited instructions; the tables are read from the PE on both roots and 7 branch anchors re-decode to the cited addresses, 0 disagree). **Medium** on the `0xb` divisor being the session's vtable pointer — the field was read as an operand, not traced to its writer, and also [EXP-0109] confirmed the gap rather than closing it: `R0124` writes no vtable at `this`
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Amended.** The arm-`0xb` clause only: There is no clock value: the dividend is `190 * rand()`, which reaches `6 225 730`, above every vtable pointer seen in that module (`L00574..L00575`), so an identically-zero quotient is **not established** and the quotient is likelier 0-or-1 (refuted, EXP-0109). `retracted.md` holds the full entries.

### AI-PURSUE-040

`R0009` writes 5 (`L00018`/`L00019`/`L00159`, `ord+0x14 = actor+0x12c`); `R0022` writes 6 (`L00576`). Arm 5 (`L00101`) loads `actor+0x12c` **itself**, not `ord+0x14`, and calls `R0041(self, target, reach)`, which is two tests: `R0051` returns the 8-way direction self→target and it must equal `mover+0x00`, the current facing (`L00577`); then `R0036` must return `<= reach` (`L00578`). `R0036` is **edge to edge**: each axis is `(2*cell + tokenSize + 0x1ff) * 128 + subcell`, the two token sizes are subtracted and the result clamped at 0, the axes maxed, then `SAR 8` and `INC` — so two touching actors measure **1**. In position → `ord+0x09 = 1`, `actor+0x54 = 3`, `actor+0x5c = target`.

Out of position → `R0042`, which reads `ord+0x14` (`L00579`) and calls `R0043`: inside the stop distance it **turns to face and stands** (the compare of the distance with the stop distance at `L00580`, then `R0056`), outside it it paths. Arm 6 is arm 5 plus `AI-GUARD-007`'s suppression clause (`L00581`, and on it `ord+0x08 = 0`) and a turn-only fallback instead of the path. **Nothing in either arm measures elapsed time, health, or how far the pursuer has come**

**Confidence.** High (every instruction cited is a PE-verified boundary; `L00101`/`L00582` is anchored by its own short branch of displacement `0x30` and `L00580`/`L00583` by a short branch of displacement `0x1c`)

### AI-BREAK-041

Nothing inside the pursuit arms ends them (`AI-PURSUE-040`); what ends one is that the `actor+0x50` arm rewrites `ord+0x08` from scratch on its own tick. For state `0xb` — the constructor default (`L00324`) — `R0137` does, in order: post ← my cell if `ord+0x00` is 0; **if `Chebyshev(me, post) >= actor+0xa5` → `ord+0x08 = 1`, `ord+0x0a = post`** (`L00267`…`L00331`); build the occupancy block of Chebyshev radius **`mover+0x08`** around **`ord+0x00`** (`L00584`/`L00585`/`R0020`), keep hostiles, engage the first survivor (`L00586` → `R0009`); block empty and not home → the same walk-home (`L00334`); block empty and home → `R0022`. The engage overwrites the sight test whenever it fires and the empty-block branch writes exactly what it writes, so **the `actor+0xa5` test is behaviourally dead** but for one corner — engage refuses *and* its acquire fallback writes nothing, which needs a human owner with nothing acquirable.

**`mover+0x08` is 5**, written once in the image by the mover constructor `R0206` (`L00587`), and read by nine instructions over eight owners, every one an occupancy-block radius: an instruction-pattern sweep for byte stores into `+0x8` returns **118 hits / 69 owners / 0 orphan** and the constructor's is the only one whose base is a mover — every other hit in the AI module follows a load of the field at `actor+0x158` and is an `ord+0x08` write. The metric is `R0162`, the max of the two absolute byte differences. The consequence a consumer must not invert: **how far the creature has chased is not measured** — it re-engages at any pursuit length while the target is within 5 of the post, and breaks off at any pursuit length once the target passes 5

**Confidence.** High (both routines read end to end, every cited address PE-verified with a branch anchor, and two rival models — a group radius and a pursuer leash — ruled out by the same listings rather than by absence of search). **Medium** that no *other* mechanism ends a pursuit: the enumeration behind that is a store-form sweep of `ord+0x08` and of the state arms, which cannot see a wholesale block write

### AI-POST-042

`R0137` sets it to the actor's current packed cell the first time it runs on a zero (the zero test of the word at `L00327`, anchored by its own branch to `L00588`). Its complete set of later writers is four: the three teardowns at the top of `R0008` — state 3 with `ord+0x0c` dead, states 8/`0x11` with `ord+0x10` dead, state 1 arrived — each writing the actor's *current* cell and then `actor+0x50 = 0xc`, `ord+0x08 = 0` (`L00589`, `L00590`, `L00591`); `R0016`'s epilogue for state 1 (`L00592`); and the patrol arm's re-anchor (`L00593`). **A state-`0xb` actor is in none of those**, so its post is the cell it stood on when guard first ran — its spawn cell for a scenario creature — and `AI-BREAK-041`'s 5-cell block is fixed there for the whole mission.

Patrol is guard with a moving post: `R0159` clears `ord+0x08`, re-anchors, **calls `R0137`** (`L00345`), and advances its ring only if guard left the order at 0 or `0xb` (`L00594`…`L00595`) — which is why `AI-PATROL-018`'s leash never pulls a patroller back *[EXP-0125]: **three clauses retracted.** (a) The writer set is not four — `ord+0x00` is at displacement **0** of `[actor+0x158]`, so a `disp:158` sweep sees the pointer load and never the store; the sweep that can see them finds at least **26 over 23 routines** (`AI-POST-096`). (b) It is not fixed for the mission: six reachable sites re-anchor it (`AI-POST-097`). (c) For a group under **order 1** the initialiser `R0137` never runs at all, because `AI-ORDER-010` gives that order's members no `actor+0x50` evaluation — the post is the spawn cell because the **setter** wrote it (`AI-POST-095`, `AI-GATE-100`). The field's identity, its emergence, and the value for an uncommanded creature all stand. See [`retracted.md`](retracted.md).*

**Confidence.** High (all five writers are cited instructions, the initialiser is branch-anchored, and the negative is over the complete `disp:158` population — 598 hits / 118 owners / 4 in orphan)

**Amended.** The writer-set completeness clause only — that `ord+0x00` is the post, and that it is emergent rather than authored, both stand: At least **26 stores over 23 routines** write it (refuted, EXP-0125). The "never moves" clause only: The guard setter has **six** reachable call sites (`callto:R0125` = 6 / 6 / 0), among them the player's group-command dispatcher `L00378` and the script's `L00596`, and each rewrites `ord+0x00` to where the unit is standing then; `R0207` does the same for one actor (refuted, EXP-0125). The mechanism for a load-time guard only — the value it reports, the spawn cell, is right: `R0137` is reached only through the per-actor `actor+0x50` machine (`callto:R0137` = 6 / 4 / 0), and `AI-ORDER-010` establishes that a group under order 1 **never evaluates `actor+0x50`** — so for the 95.6 % of shipped placements `AI-CENSUS-047` puts under order 1, the initialiser never runs (superseded, EXP-0125). `retracted.md` holds the full entries.

### AI-STATE-043

Two store-form sweeps: 4-byte stores of an immediate into `+0x50` (**81 hits / 55 owners / 0 orphan**) and 4-byte stores of a register into `+0x50` (**151 / 78 / 3 in orphan**). The displacement is shared with `ord+0x50`, so a hit counts only when the destination register is the actor rather than one just loaded from `+0x158`; that leaves the value set **{0,1,2,3,4,8,0xa,0xb,0xc,0xd,0xe,0xf,0x10,0x11,0x16}**. The register forms resolve: `L00597`/`L00598`/`L00599` = `0xc` (`EBP` at `L00600`), `L00140` = `0xc` (`L00601`), `L00150` = `0xc` (`L00602`), `L00603` = `1` (`L00604`), and **`L00206` = `0x16`** (`EBP` at `L00605`, in `R0103`, player order `0x14`) — which closes this ledger's open item 9 on a register-form writer for the withdraw state.

Against `AI-STATE-011`'s 27-arm table: `5,6,7,9,0x12..0x15` are unwritten *and* on the default, consistent; but **`0x17`, `0x18`, `0x19`, `0x1a` have live arms and no writer**, and `4` — also a live arm — is written at exactly one site, `L00606`, which is *inside arm `0x18`*. Five live arms therefore form a cluster reachable only from itself. `R0137` has a fourth caller, `R0175`, and `callto:R0175` is **0 hits**, so it adds no path either. **Amended by [EXP-0385]:** `0x10` is not unwritten — teardown `R0208` (`TRIG-REAP-017`'s own subject) writes `0x10` into `+0x50` at `L00607`, on the function's own unmodified `this` (reloaded at `L00608` from the same frame local `-0x3c` prologue stash the `+0x54` write at `L00609` shares; no `[X+0x158]` dereference intervenes, so this is `actor+0x50`, not `ord+0x50`).

`TRIG-REAP-017` already cited this instruction at High confidence; this row's own sweep should have counted the hit under its own stated pattern and disambiguation rule and did not — the miss is outside the three blind-spot categories this row's own confidence text names (`CArchive` restoration, scaled-index addressing, in-place arithmetic), so an instrument's negative claim about one value is bounded by what its scan actually executed over, not by which named category it later turns out to fall in

**Confidence.** **Medium**, and the instrument's blind spot is the reason: a store-form sweep cannot see a state restored by `CArchive` from a save, written through a scaled-index displacement, produced by arithmetic in place, or — per the `0x10` correction — missed by the sweep's own pattern for an unidentified reason. The *positive* half — that the listed values are written at the listed addresses — is High

**Amended.** Value-set clause only: `0x10` is written: teardown `R0208` stores it at `L00607`, on the function's own unmodified `this`, confirmed by a fresh whole-body read and matching `TRIG-REAP-017`'s own prior citation of the same instruction (amended, EXP-0385). `retracted.md` holds the full entries.

### AI-THREAT-044

`R0009` walks `ESI = 0x78` to `0x84` step 4 over `ord+0x78`/`+0x7c`/`+0x80` (`L00610`…`L00611`); a non-null slot whose companion dword at `+0xc` — `ord+0x84`/`+0x88`/`+0x8c` — is **greater than `R0179()`**, ~~the tick clock~~, short-circuits the whole selection into `R0209(self, target, cached)` — *amended by [EXP-0109]: `R0179` is the CRT `rand()`, read whole as a linear congruential generator returning `0..0x7fff` (`AI-RAND-058`), so the comparison this row cites correctly is against a **random draw**, not a clock, and the words "timed", "unexpired" and "expiry" below describe nothing in the image. The cache is a probabilistic short-circuit whose odds are the slot's stored value over 0x8000; whatever fills a slot fills it with a threshold, not a deadline. The structural half of the row is untouched. See [`retracted.md`](retracted.md).* — *further amended by [EXP-0387]: "threat cache" and "cached" misname what the slots hold and what `R0209` does with a match. The three slots are the actor's own class spellbook (`UNIT-SPELL-007`), and the value handed to `R0209` is proven by that callee's own body to be a spell id, resolved through `Spellbook::Get`, never dereferenced as a target or threat pointer (`AI-341`, `MAGIC-221`). The loop, slot bounds, comparison direction and short-circuit-the-acquisition-pipeline behaviour this row cites are unaffected; only the identity of what is cached is corrected. See [`retracted.md`](retracted.md).* Nothing anywhere records that a target was *abandoned*: the teardowns end a pursuit by changing the state, not by suppressing re-selection, so the same actor may re-enter a pursuit on the very next group tick — and while a cache slot is unexpired it will

**Confidence.** High for the structure and the read (the loop bounds, the call and the comparison are cited instructions). **Medium** for what the cache means: the write side was not swept, so which event fills a slot, and with what threshold, is unread — *and also [EXP-0109] narrows that negative without closing it: whole-image `EnumRefs disp:78` (438 hits / 210 owners) and `disp:84` (578 / 273) contain **no** access to `ord+0x78` or `ord+0x84` anywhere in the AI module, the only non-stack hits at those displacements below `L00612` belonging to an inventory class and to the order deserializer `R0210`; a `disp:` sweep still cannot see a wholesale copy of the order block*

**Amended.** The identity of the comparand, not the structure: `R0179` is the CRT `rand()` — a linear congruential generator read whole (`seed = seed * 214013 + 2531011`, return `(seed >> 16) & 0x7fff`) (superseded, EXP-0109). The "threat cache" naming and "cached" as a target/threat reference: The three slots are the actor's own class spellbook (`UNIT-SPELL-007`'s `Spell 1..3`/`Probability 1..3` columns), the identical writer and offsets, not a target-tracking cache (narrowed, EXP-0387). `retracted.md` holds the full entries.

### AI-ROUTE-045

`R0043` sets `mover+0x98 = 1` when the search leaves the path list empty (`L00161`, a ten-byte store of 1 into `+0x98`). `R0016`'s epilogue is its **only** consumer: the branch at `L00613`, a 32-bit displacement re-decoded from the bytes to `L00614` — skips everything when it is 0, else it is consumed and the actor's *state* selects the response: `1` → the teardown to `0xc` with the post reset (`L00615`); `0xa` → advance the patrol ring; `0x17` → `ord+0x09 = 0xff`; **everything else → `ord+0x08 = 0` and `R0004`** (`L00114`/`L00011`; `callto:R0004` is **1 hit / 1 owner / 0 orphan**). `R0004` then walks the whole `world+0xa4554` list keeping only candidates with `dist <= actor+0x12c` (`L00616`/`L00617`), diplomacy-filters, scores on **turn cost** rather than distance, and writes `ord+0x08 = 6` — or `0xb` for a non-human owner with nothing in reach (`L00124`), or `0` for a human one. So an unreachable target does not merely stall a chase: it cancels the order and drops the actor into a reach-bounded idle

**Confidence.** High (the flag's writer, its single reader and every branch of the epilogue are cited instructions, the epilogue's own `JZ` is a re-decoded rel32, and the consumer enumeration is a `callto:` with 0 orphan)

## Load-time census of group orders

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-CENSUS-046 | The load-time census: over every shipped map of both roots the group order byte takes exactly two values, and `0` is not one of them — so no creature *starts* a mission under `AI-BREAK-041`'s post rule. | Medium | ● active | [EXP-0102](../experiments/EXP-0102-group-order-census/) |
| AI-CENSUS-047 | What the script changes, bounded twice — a player meets the group-CENTRE rule, and the post rule is a curiosity. | Medium | ● active | [EXP-0102](../experiments/EXP-0102-group-order-census/) |
| AI-CENSUS-048 | The campaign in mission order — what the first missions actually put in front of the player. | Medium | ● active | [EXP-0102](../experiments/EXP-0102-group-order-census/) |
| AI-ROOT-049 | The two roots differ in four places and in none that moves this census. | Medium | ● active | [EXP-0102](../experiments/EXP-0102-group-order-census/) |

### AI-CENSUS-046

`tools/grouporder` enumerates maps by `M7R\0` signature rather than by filename — **EN 38 (28 in `scenario.res`, 10 loose), RU 34 (28, 6)** — builds the runtime group set as the distinct `(owner, +0x42)` pairs the spawner would form (`AI-GROUP-009`), and assigns each the order `AI-AUTHOR-015`'s load walk gives it: **3** for the groups of type-5 slot 0, **1** for every other slot. Result, weighted by placed creatures: **EN 8085 at order 1, 9 at order 3, 0 at order 0** over 8094 placements in 2264 groups; **RU 3982 / 9 / 0** over 3991 in 1372. The nine order-3 placements are on four campaign maps (`41` 3, `71` 4, `150` 1, `151` 1); slot 0 owns nothing on the other 24 campaign maps and on every loose map.

Five published figures reproduce to the digit from this independently written probe — `ALM-OWN-039`'s **8094** type-6 records, `AI-GROUP-009`'s **2264** groups / mean **3.58** / **30.3 %** singletons, `AI-AUTHOR-015`'s slot-0 name on **38/38** EN maps (32 `Self`, 6 `self`), `AI-GROUPCMD-020`'s node census below, and `REG-SCN-059`'s four `MissionObjects` entries with no `scenario.reg` section (`81 101 141 151`)

**Confidence.** **Medium** — a census, and `METHODOLOGY.md` caps corpus agreement there; the counting half is exact and reproducible, but reading a count as "the post rule is not met at load" runs through `AI-AUTHOR-015`, which is itself **Medium** on the `server+0x0c == 0` single-player gate. **This experiment did not read whether that gate wraps the whole walk or only its aggressive branch** — if it wraps the walk, a multiplayer load leaves every group at the constructor's 0 (`AI-CMD-033`) and the answer inverts to 100 % post rule

### AI-CENSUS-047

The type-7 action opcode 6 (`AI-GROUPCMD-020`) is the only surface that moves the byte on shipped content, and its reach is small: **135 nodes, identical on both roots, over 20 campaign maps and 0 of the loose ones**, naming **91 of 2264 groups** and **361 of 8094 placements** (4.5 %). Per sub-command, against each root's own PE table at `L00408`: `1 Guard` 1, `2 Swarm` 6, `3 Stand Ground` 24, `4 Move` 29, `5 Swarm 2` 40, `10 Attack` 5, `11 Defend` 6, `14 Patrol` 14, `15 Follow` 10; `17 Roam` and `18 Dwell` ship **0**, reproducing `AI-ROAM-025`. **45 of the 135 are named by no trigger's action slot at all** and so can never be dispatched — a counting fact, not a simulation one. Two bounds follow, pessimistic (every node fires) and carried-only: order **0**, the post rule, governs **68 down to 34 of 8094 placements (0.8 % down to 0.4 %)**; order **1**, the centre rule, **7739 up to 7819 (95.6 % up to 96.6 %)**; the rest are 106/92 at order 3, 88/86 at 5, 37/37 at 2, 10/14 at 4, and 46/12 in groups commanded to more than one order.

Restricted to creatures hostile to slot 0 in either direction of the map's own diplomacy rows (`AI-DIPLO-004`/`AI-DIPLO-005`) — the only ones whose break-off a player can observe — the centre rule takes **7293 of 7571, 96.3 %**, the post rule **52, 0.7 %**. The rival that the script makes the load-time value irrelevant fails by two orders of magnitude

**Confidence.** **Medium** — corpus, and it inherits `AI-GROUPCMD-020`'s reading of which order each arm leaves plus `AI-ORDER-010`'s writer enumeration, itself **Medium** on completeness. The keying rival was tested: 4 nodes name a group id that more than one owner slot carries, and the **28 placements** in groups reached only through such a node bound what that ambiguity can move
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration. The agreement between the per-root
data stands as corroboration: the map files whose node counts are cited are separate files in each root.

### AI-CENSUS-048

Ordering by the map's own number, with `scenario.reg`'s `[Mission<n>]` sections separating 15 main missions from 9 sub-missions (`REG-SCN-059`): **mission 10, the first, has 35 placements of which 21 are hostile, all 35 at order 1 at load, 6 opcode-6 nodes touching 3 placements, and 2 placements ever reaching order 0 — 32 of 35 under the centre rule.** **Eight of the 28 campaign maps carry no group command whatsoever** (`20 30 31 41 51 61 101 121`), so on those the load-time order stands unchanged for the entire mission. No campaign map has a majority of its creatures away from order 1: the largest order-0 share on any map is `100.alm`'s 20 of 113, and the largest non-1 share of any kind is `151.alm`'s 92 of 237 at order 3. Per-mission table in `evidence/mission.txt`, per-map row in `evidence/map-census.csv`

**Confidence.** **Medium** — the same census restricted to the campaign; the mission ordering is the map numbering and `scenario.reg`'s own section numbering, not a read of the mission-advance code

### AI-ROOT-049

All **28 campaign maps agree placement for placement, group for group and opcode-6 node for node**, and 30 of the 34 maps both roots ship have byte-identical type-5, type-6 and type-7 payloads. The four exceptions: (a) the RU root ships **6 loose maps against EN's 10** — `Beast`, `Cross`, `Kids2` and `Tomb` are EN-only; (b) RU's **`Horror.alm` is a four-record terrain-only container** — 262876 B, header record count **4**, types `0 1 2 3` only, while its own type-0 metadata still declares 415 structures and 1815 units and EN's file carries all ten records at 409386 B, so on the RU root that map has no roster, no placement, no script and no trigger; (c) `scn:100.alm` has the same 36 action nodes but **42 checks / 19 triggers on EN against 40 / 17 on RU**, a 1960-byte type-7 difference; (d) `scn:140.alm` differs in exactly one byte outside a text field — **trigger 5's fire-once flag at `+0xb4`, 1 on EN and 0 on RU**, so that trigger re-arms every pass on RU (`TRIG-FIRE-007`). The remaining differences are authored **text**: type-5 records differing only inside their 32-byte name (`lumoir` 2, `scn:100` 4, 0 outside the name in both) and one `lumoir` node differing only in a label field

**Confidence.** **Medium** — a census over two installs; the payload comparison is sha256 plus a field-masked byte diff, which shows *that* the bytes differ and where, not what the engine does with a truncated container. **Nothing here was run**: whether RU's `Horror.alm` loads at all is untested

## Click, cursor and attack order production

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-CLICK-050 | A map click is turned into an order by the CURSOR it was made under, not by what it hit — and the attack cursor is the only one that produces two different orders. | High | ● active (partially retracted) | [EXP-0109](../experiments/EXP-0109-attack-order/), [EXP-0190](../experiments/EXP-0190-combat-controls/) |
| AI-CLICK-051 | What an attack order carries, from its producer: one static command record, a target actor id, and every selected unit by id. | High / Medium | ● active (amended, superseded) | [EXP-0109](../experiments/EXP-0109-attack-order/) |
| AI-CURSOR-052 | The hostility test happens at HOVER, one step before the click, and it reads a view-side row rather than the session's diplomacy matrix. | High | ● active (amended, partially retracted) | [EXP-0109](../experiments/EXP-0109-attack-order/), [EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/) |
| AI-PANEL-053 | Ordering an attack takes two clicks, and the first one only ARMS a mode | High | ● active (partially retracted, superseded) | [EXP-0109](../experiments/EXP-0109-attack-order/), [EXP-0110](../experiments/EXP-0110-order-gates/), [EXP-0190](../experiments/EXP-0190-combat-controls/) |
| AI-CMD-054 | Order opcodes `0x19` and `0x1b` read as bodies rather than named by the routine they call, which lifts two of `AI-CMD-032`'s Medium labels. | High / Medium | ● active | [EXP-0109](../experiments/EXP-0109-attack-order/) |
| AI-STRIKE-055 | Everything that makes a unit attempt a blow arrives through two per-actor states, and the shipped campaign reaches all of the surfaces that write them. | Medium | ● active | [EXP-0109](../experiments/EXP-0109-attack-order/) |
| AI-RETAL-056 | Being struck issues no order and produces no target: it sets one flag whose single consumer makes the victim TURN. | High / Medium | ● active | [EXP-0109](../experiments/EXP-0109-attack-order/) |
| AI-ARBITER-057 | One arbiter, not two: the AI tick applies the same machine to every player's groups, and ownership enters only as a local clause inside shared routines. | High / Medium | ● active | [EXP-0109](../experiments/EXP-0109-attack-order/) |
| AI-RAND-058 | `R0179` is the CRT `rand()`, and two published rows read it as a clock. | High / Medium | ● active (amended, superseded) | [EXP-0109](../experiments/EXP-0109-attack-order/) |

### AI-CLICK-050

`R0211` is the map-click handler: ~~it discards the click as a drag when either axis moved more than `[L00618] * 10 / 0x280`~~ *[EXP-0190] reads the callee and corrects this: a rectangle strictly beyond that threshold is sent to `R0212` for multi-selection; it is not discarded*, returns to selection `R0212` when `view+0x140` is 0, and otherwise compares the current cursor `[L00619]` against each registered cursor's own handle at `cur+0x04`, one arm per cursor. `tools/attackorder` re-derives the whole map out of the PE by byte pattern on both roots (`evidence/attack-order-tables.txt`): **attack** → `R0213` opcode **`0x19`** when `view+0x990` is non-zero **and** `R0214(view+0x98c, L00620)` holds, else `R0215` opcode **`0x16`**, a plain move to the cell; **swarm** → `0x1a`; **move** → `0x16`; **patrol** → `0x1d`; **defend** → `0x1b`, and **nothing at all** when no actor is under the cursor; **select** → `R0212`, no order; **pickup** → `0x21`; **town** → `0x24`; **cast** → `0x25`/`0x1e` at a unit and `0x1f`/`0x26` at a cell.

`R0214` is `CObject::IsKindOf` — it calls MFC's virtual slot 0, `GetRuntimeClass`, on the object, then `R0216` — and `L00620` is a `CRuntimeClass`. **So the click-time test is a runtime-CLASS test, and no diplomacy is consulted at click time at all** (`AI-CURSOR-052` has where it is consulted). *[EXP-0190] also fixes the physical edge: `WM_LBUTTONDOWN` begins the marquee and `WM_LBUTTONUP` reaches this routine.*

**Confidence.** High for the cursor arms and physical edge. The drag-discard clause is retracted by EXP-0190. *[EXP-0110] read the second attack-builder caller (`AI-MINIMAP-062`) and swept the whole family (`AI-SURFACE-063`).*
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Amended.** Partially retracted by EXP-0190. The drag-discard clause only — the cursor arms, attack fallback and absence of a click-time diplomacy test stand: `R0211` sends a rectangle whose width or height is strictly greater than the threshold to `R0212`, and that routine performs rectangle selection (refuted, EXP-0190). `retracted.md` holds the full entries.

### AI-CLICK-051

`R0213` writes into the **static** buffer at `L00621` — there is no allocation — the byte `0x19` at `+0x09` (`L00622`), the issuing player's index `word [player+0x04]` at `+0x05` (`L00623`), `0` at `+0x07`, **the target actor id as a `u16` at `+0x0e`** (`L00624`, the value being `view+0x990`), and `0` at `+0x12`; it then walks ~~the selection array at `view+0x9b8` (`+0x04` data, `+0x08` count)~~ — *[EXP-0110]: a **`CMapWordToPtr`**, `+0x04` the bucket array and `+0x08` the bucket count, element count at `+0x0c` = `view+0x9c4`; the walk and its conclusion are unchanged (`AI-SELECT-065`)* — and calls `R0092(cmd, word[unit+0x04])` once per member, **for members only, and for at most 253 of them** (`AI-FANOUT-064`), which is the count-and-ids append `AI-CMD-032` describes from the consumer end.

It enqueues with `R0217(L00625, cmd)` and the record then travels `R0192`/`R0193` → `R0191` → the dispatcher `R0061` (`callto:R0061` = **1 hit / 1 owner / 0 orphan**). The move builder `R0215` is the same shape with a **cell** (`view+0x994`, `view+0x998`) where the target id sits. **Customisation:** the record is a fixed global, so one order is in flight at a time by construction, and the target is addressed by a `u16` **id**, not a pointer — the same wire the multiplayer path uses

**Confidence.** High (each store is a named instruction in one routine read end to end, and the enqueue's whole caller chain is a printed `callto:` count on the repaired function table). **Medium** that `L00621` is never written by anything else between fill and dispatch — no sweep of that address was run

**Amended.** The container's field layout only — every store and the conclusion stand: **It is a `CMapWordToPtr`, not an array.** `+0x04` is `m_pHashTable`, `+0x08` is `m_nHashTableSize` (buckets, not elements) and the element count is `+0x0c` = **`view+0x9c4`** — confirmed by arithmetic against the field the inlined `GetStartPosition` idiom, a negate followed by a self-subtract with borrow, actually loads at `L00626`… (superseded, EXP-0110). `retracted.md` holds the full entries.

### AI-CURSOR-052

`R0218` (the hit test) returns a capability mask and stores the hit in `view+0x98c` / `view+0x990`; `UNIT-HOVER-020` gives the mask's bits. `R0219` turns mask into cursor: with the gate `[L00627]` set, `mask & 3` — the `CUnit`/`CAirUnit` bits — chooses **attack** (`L00628`) and its absence chooses **swarm** (`L00629`) (`L00630`…`L00631`); with `[L00632]` set the cursor is **move** ~~regardless~~ (`L00633`) *-- “regardless” is refuted by `AI-CURSOR-227` ([EXP-0218]): six later stores on that same path replace the value, and the second cascade's first arm never reads the latch*; and ~~the hostility bit `0x4` appears only inside `mask & 0x24`~~ *-- refuted by `AI-CURSOR-227` ([EXP-0218]): `L00634` and `L00635` each test the bit alone, and the second of them selects **attack** with no modifier key held --* the `mask & 0x24` test's own effect is to **suppress the select cursor** (`L00636`…`L00637`). The cursor identities are not inferred: `R0220` registers 28 cursors by art path and `tools/attackorder` pairs each pushed art path with the global it is stored to, read out of the PE — the attack cursor's art is `graphics\cursors\attack\sprites.16a` and its global is `L00628` on both roots

**Confidence.** ~~**Medium**, and the weakest input is named: the three gate globals are read, not explained, so every clause standing on them is capped here~~ — *the cap is **lifted** by [EXP-0110] (`AI-KEYMOD-059`), which reads all three: they are **modifier-key-held latches**, set on WM_KEYDOWN of vkey `0x11` / `0x10` and on WM_SYSKEYDOWN of vkey `0x12`, cleared on the matching WM_KEYUP and wholesale on WM_KILLFOCUS. The four-routine family this cell names is right and incomplete — two further writers, `R0221` and `R0222`, both clear `[L00632]`.* **High** throughout, therefore: the cursor-to-art binding (a PE read on both roots), each `AND` and `MOV` cited above, and now the gates themselves. In consumer terms `[L00627]` held turns the hover into the attack cursor and `[L00632]` held forces the move cursor *inside the cascade at `L00633`, which `AI-CURSOR-226` ([EXP-0218]) establishes is the cast-aiming path and not the one an ordinary hover runs*; naming those keys Ctrl and Alt is the Windows virtual-key contract and never raises this grade
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Amended.** Partially retracted by EXP-0218. Two clauses of the body: the Alt "regardless" clause and the clause that the hostility bit `0x4` appears only inside `mask & 0x24`. The headline, the `R0218` hit-test description, the attack/swarm arms of the `L00633` cascade, the select-suppression effect of `mask & 0x24` at `L00636`-`L00637`, and the cursor-to-art binding all stand; the modifier-latch row is named in the “what is true instead” cell rather than here, because this cell's ids are read as overturned: `R0219` contains **two** mask-to-cursor cascades, not one (refuted, EXP-0218). `retracted.md` holds the full entries.

### AI-PANEL-053

~~— which the panel refuses unless the current selection supports it~~ — *[EXP-0110] read `R0093` and it never touches the selection: it returns `0xef` (every mode but cast) for any selection the local participant owns, `0` for one he does not, and adds cast on one flag. See `AI-PANEL-060`, `AI-PANEL-061`.* `R0097` is the command-panel handler: `mode = button` (buttons 9 and 0xa fold to 5, `L00638`/`L00639`), and `view+0x99c = mode` is reached **only** when `R0093(view)`'s mask — *a read of `view+0x144`, not of the selection* — carries `1 << (mode-1)` or the button is 0xa (`L00640`…`L00641`); otherwise the routine returns having armed nothing. The armed mode is one-shot: `R0211` clears `view+0x99c` and notifies `0x40b` after every consuming click (`L00642`…`L00643`).

Both tables are read out of the PE by `tools/attackorder`, identically on both roots. **Armed-mode table, 8 dwords at `L00644`, index `view+0x99c - 1`:** mode **1 attack**, 2 move, 3 and 7 on the join (no cursor), 4 defend, 5 cast, 6 swarm, 8 patrol. **Panel table, 6 dwords at `L00645`, index `button - 3`:** buttons 3, 7 and 8 **issue an order immediately** — `R0223` `0x17` Guard, `R0224` `0x18` ~~aggressive~~ *Stand Ground*, `R0102` `0x14` *Retreat* — while 4, 5 and 6 only arm; buttons 1 and 2 are below that table's base and therefore only ever arm. So **Guard, Stand Ground and Retreat are one-click commands and attack is a two-click targeted order**. *[EXP-0190] supplies the labels from `main.txt[0..7]`, proves R takes the immediate Retreat arm and finds no default route that leaves Patrol armed.*

**Confidence.** High for the tables and action routes. The “aggressive” label is retracted by EXP-0190; opcode `0x18` was always the same correctly read arm.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Amended.** Partially retracted by EXP-0110 and EXP-0190. The capability clause only — the two tables and the two-click finding stand: **`R0093` never touches the selection.** Twenty instructions, read whole: it returns 0 when `view+0x144 & 4`, otherwise `0xef`, with `0x10` added when `view+0x144 & 0x200` (superseded, EXP-0110). The labels of opcodes `0x18` and `0x14` only — the immediate arms and opcode bytes stand: The physical cells and their localised `main.txt` lines bind cell 6 / key T / opcode `0x18` to **Stand Ground**, and cell 7 / key R / opcode `0x14` to **Retreat** (label moved, EXP-0190). `retracted.md` holds the full entries.

### AI-CMD-054

`0x19` is **attack**: `R0006(group, target)` clears every member (`R0007`), zeroes `ord+0x50` and `ord+0x38`, sets **`grpAI+0x20 = 0`** (`L00646`), and then per member either `actor+0x50 = 3` with `ord+0x0c = target`, `ord+0x14 = actor+0x12c` and `ord+0x08 = 0` (`L00014`…`L00016`) — `AI-STATE-011`'s arm 3, the engage routine `R0009` — or, when `R0225(member, target)` returns `0xffffff`, it turns the member to face and sets `actor+0x50 = 0xc` with `ord+0x00 = own cell` (`L00647`…`L00648`), arm `0xc`, `R0022`, acquire with no leash. `0x1b` is **defend a unit**: `R0190` gives the *target* `actor+0x50 = 0xc` (`L00649`) and every other member `actor+0x50 = 8` with `ord+0x10 = target` (`L00650`/`L00651`), the follow pair. So a player attack order and an unordered AI unit converge on the **same two per-actor states**, and the difference is only which one the member starts in

**Confidence.** High (both routines read end to end; every state, order field and group-order store above is a cited instruction, and the two values reproduce `AI-CMD-033`'s independently derived `0x19` → 3 then `0xc`, and `0x1b` → `0xc`/8). **Medium** for what `R0225` computes — it is read as a guard, not traced

### AI-STRIKE-055

`actor+0x50 = 3` (engage a named target, `R0009`) and `actor+0x50 = 0xc` (acquire whatever is visible, `R0022`) are the only two `AI-STATE-011` arms that end in a target, and every cause reaches one of them: a **player order** `0x19`/`0x1b` (`AI-CMD-054`), the **load-time aggressive stance** `R0060`/`R0063`, which writes `0xc` to every member and `grpAI+0x20 = 3` (`AI-AUTHOR-015`), the **guard state** `0xb`, whose occupancy-block survivor goes to `R0009` (`AI-GUARD-012`), the **group guard arm** `R0024` under group order 1, and the **script** `Par0 = 10 Attack` (`AI-GROUPCMD-020`). Being struck is **not** on this list (`AI-RETAL-056`). What shipped content reaches, from the published censuses rather than re-derived here: every group starts at order **1** or **3** and none at 0 (`AI-CENSUS-046`, EN 8085 / 9 / 0 over 8094 placements), the script's `10 Attack` ships **5** nodes and `11 Defend` **6**, over 20 campaign maps and 0 loose ones, and 8 of the 28 campaign maps carry no group command at all (`AI-CENSUS-047`, `AI-CENSUS-048`)

**Confidence.** **Medium** — the *routing* of each named cause is High and each is a cited instruction in the row that owns it, but the claim that the list is **complete** is an enumeration over `actor+0x50`'s writers, and `AI-STATE-043` grades that set Medium for exactly the reason quoted there: a store-form sweep cannot see a state restored from a save or produced by arithmetic in place. The corpus half inherits `AI-CENSUS-046`/`AI-CENSUS-047`'s own Medium

### AI-RETAL-056

`R0226(session, attacker, victim)` is nine instructions (`evidence/rom-attackorder-excerpt2.md` section 7): `ord+0x54 = 1` (`L00652`), `ord+0x58 = R0019(attacker+0x10)`, the attacker's cell (`L00653`), `ord+0x5a = 0` (`L00654`), then `R0126(session+0xa9bc, attacker, victim)`, the diplomacy flip. `EnumRefs re:` over the whole `+0x54` access form finds exactly **one** read of `ord+0x54` — `L00655` in `R0205`, `AI-ORDER-039`'s order arm `0xb`, the idle turn — where a non-zero flag **skips the `rand()` gate at `L00656`** so the actor re-picks a facing on this evaluation instead of on about one evaluation in 160, and the flag is then cleared (`L00657`). It does not touch `actor+0x5c`, `ord+0x08` or `actor+0x50`.

**So the whole of retaliation is: turn now, become mutually hostile, and let ordinary acquisition do the rest** — and for an AI-owned victim only, `ord+0x58`/`ord+0x5a` additionally force the attacker's cell into the sight stamp for 20 ticks (`AI-SIGHT-006`). This closes one direction of `HERO-AGGRO-028`'s Medium on "the victim's own AI arms were not enumerated"

**Confidence.** High for the hook and the arm (both read whole, every field a cited instruction). **Medium** for "exactly one reader": `re:` prints one preceding instruction, so a load of `actor+0x158` two or more instructions before its use is invisible to the filter, and a wholesale copy of the order block carries no displacement at all

### AI-ARBITER-057

`R0076` walks the player list it is handed, then each player's group list at `player+0x24`, and calls `R0023` per group gated **only** on `session+0xb388` or that group's own `grpAI+0x45` (`L00658`…`L00282`). **It contains no owner test at all**: all sixteen of its `+ 0x28` accesses are `[ESP + 0x28]`, `[EDI + 0x28]` or `[EBX + 0x28]` on the tick's own timer sub-objects, and none is preceded by a `+0x14` load (listing in `evidence/rom-attackorder-excerpt2.md` section 11). Where ownership does act it is `UNIT-OWNER-009`'s `actor+0x14` then `owner+0x28` idiom, which `EnumRefs re:` finds at **22 sites over 21 distinct owners in `L00339..R0227`, 0 in orphan code** — among them `R0022` (`L00659`), `R0009` (`L00660`, `L00661`), `R0016` (`L00662`), `R0004` (`L00663`, `L00664`), `R0023` (`L00665`), `R0110`, `R0024`, `R0025`, `R0026` and `R0126` (`L00666`, `L00667`). **So the customisation seam for AI behaviour is one machine with 22 owner-conditioned clauses, not two behaviour trees**

**Confidence.** High for the tick (read whole; the absence is an exhaustive read of one routine, not a search) and for each listed site individually. **Medium** for the clause list's completeness — `re:` prints one preceding instruction, so an owner fetched more than one instruction before its use, or held in a register across a call, is not counted

### AI-RAND-058

Fourteen instructions, read whole: `R0228()` returns a block whose `+0x14` holds the seed; `seed = seed * 214013 + 2531011` is built from shifts, adds and subtractions of the old seed ending in the addend `0x269ec3` (`L00668` to `L00669`); the return is the new seed shifted right by 16 and masked with `0x7fff` — MSVC's linear congruential generator with `RAND_MAX` `0x7fff`. `AI-GUARD-012` calls it `rand()` and is right; **`AI-THREAT-044` calls it "the tick clock" and is not**, so that row's threat cache is gated by a **random draw in `0..0x7fff`**, not by an expiry, and its words "timed", "unexpired" and "expiry" do not describe the comparison it otherwise cites correctly.

**`AI-ORDER-039`'s arm-`0xb` clause fails for the same reason**: there is no clock value, and the quotient it calls "0 for any clock value" is `190 * rand() / [session]`, whose numerator reaches `6 225 730` while every vtable pointer this experiment saw in the AI module lies in `L00574..L00575`, all **below** it — so an identically-zero quotient is not established *[EXP-0125]: the Medium is **closed**. `R0124` writes no vtable at `this` because it writes a number: the store of `0x8000` into the first dword of the AI manager at `L00371`, the global at `[L00004]`. `0x8000` is `RAND_MAX + 1`, so `n × rand() / AImgr[0]` is `floor(n·rand()/32768)` and the arm-`0xb` quotient is **uniform on 0…189**, not merely unestablished as zero (`AI-RANGE-102`, `AI-TURN-104`).*

**Confidence.** High for the identification (a complete LCG read at instruction level, with both constants recovered from the shift-and-add chain rather than from an immediate, and the `0x7fff` mask fixing `RAND_MAX`). **Medium** for the arm-`0xb` consequence, because the session's own vtable address was not pinned — `R0124` writes no vtable at `this`, so the bound above is over the module's vtables and not over the divisor itself

**Amended.** The arm-`0xb` consequence's Medium only — the LCG identification and both constants stand: The divisor is pinned (superseded, EXP-0125). `retracted.md` holds the full entries.

## Modifier latches, the selection mask and the producer census

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-KEYMOD-059 | The three globals that gate the cursors are modifier-key latches: `[L00627]` vkey `0x11`, `[L00670]` vkey `0x10`, `[L00632]` vkey `0x12` — the last of them the system key, established from inside the image. | High | ● active | [EXP-0110](../experiments/EXP-0110-order-gates/) |
| AI-PANEL-060 | The routine `AI-PANEL-053` calls the selection capability mask never looks at the selection: it is twenty instructions over one view-side flags word, and the only mode any selection can be refused is CAST. | High | ● active | [EXP-0110](../experiments/EXP-0110-order-gates/) |
| AI-PANEL-061 | `view+0x144` is a selection-derived flags word rebuilt on every selection change, and only two of its six bits reach the panel — one of them ownership. | High / Medium | ● active (partially retracted) | [EXP-0110](../experiments/EXP-0110-order-gates/) |
| AI-MINIMAP-062 | There is a second click surface, it is the SMALL-cursor one, and its gates are not the map click's. | High | ● active | [EXP-0110](../experiments/EXP-0110-order-gates/) |
| AI-SURFACE-063 | The complete producer census: exactly three input surfaces reach the order-builder family, plus one producer with no cursor arm at all. | High / Medium | ● active (partially retracted, superseded) | [EXP-0110](../experiments/EXP-0110-order-gates/) |
| AI-FANOUT-064 | One order carries at most 253 unit ids and the 254th is dropped in silence. | High | ● active | [EXP-0110](../experiments/EXP-0110-order-gates/) |
| AI-SELECT-065 | `view+0x9b8` is a `CMapWordToPtr` from unit id to object, not an array, and the order carries every entry the per-object flag `+0x7c` marks. | High / Medium | ● active | [EXP-0110](../experiments/EXP-0110-order-gates/) |

### AI-KEYMOD-059

`tools/ordergates` reads the WM_KEYDOWN handler's own jump table out of the PE (`0xb9` index bytes at `L00671`, arms at `L00672`) and only **two** of its thirteen arms touch a global: vkey `0x11` reaches the store of 1 into the global at `L00627` (`L00673`) and vkey `0x10` reaches the store of 1 into the global at `L00670` (`L00674`). `R0229` clears them by the same vkeys and `[L00632]` by vkey `0x12` — its chain of subtract-and-branch steps on the vkey (first `0x10`, then 1, then 1) is decoded from the bytes, not assumed — and `R0230` clears all three together (`L00675`/`L00676`/`L00677`). The **message map** is read from the PE as `AFX_MSGMAP_ENTRY` records at `pfn-0x14`: `R0231` from `0x100` WM_KEYDOWN, `R0229` from `0x101` WM_KEYUP, `R0232` from `0x104` WM_SYSKEYDOWN, `R0230` from `0x008` WM_KILLFOCUS.

**Set on key-down, cleared on key-up of the same key, cleared wholesale on focus loss** — a key-held latch by construction, with no second model fitting it. That `0x12` is the **system** key needs no outside fact: `R0232`, the WM_SYSKEYDOWN handler, sets `[L00632]` on two independent routes — `nChar == 0x12` (`L00678`) and the three hotkey arms guarded by `nFlags & 0x2000` (`L00679`, `L00680`) — which agree only if `0x12` is the key that bit reports. **Consumer reading:** `[L00627]` is the modifier that turns the hover into an attack cursor, `[L00632]` the one that forces the move cursor (`AI-CURSOR-052`), and `[L00670]` the one that makes a marquee *add* to the selection instead of replacing it (`L00681`, a set-byte-on-zero, `R0212`). Two more writers `EXP-0109` did not name, both clearing `[L00632]`: `L00682` (`R0221`) and `L00683` (`R0222`)

**Confidence.** High for the **mechanism** and the vkeys — the jump table and the message map are PE reads on both roots with byte-identical bodies, every store and clear is a cited instruction, and `refto:` gives the three globals 19 / 17 / 11 references over 5 / 10 / 8 owners, 0 orphan, so no writer is unaccounted for. **Naming vkey `0x10`/`0x11`/`0x12` Shift/Ctrl/Alt, and `0x100`/`0x101`/`0x104`/`0x008` WM_KEYDOWN/KEYUP/SYSKEYDOWN/KILLFOCUS, is the Windows platform contract, not a reading of ours** — it is labelled, it never raises a grade, and every clause above survives its withdrawal because the mechanism does not depend on it. Owner testimony is the falsifier for the naming and the prediction is what this experiment ships in its place

### AI-PANEL-060

`R0093(view)` read whole (`evidence/rom-ordergates-excerpt.md` section 1): `L00684` loads `view+0x144`, `L00685` tests bit `0x4` of it — set **returns 0** (`L00686`); otherwise `L00687` stores the mask `0xef` in a local, then `L00688` reloads `view+0x144`, `L00689` tests bit `0x200` and when it is set `L00690` adds bit `0x10` to the mask. It touches `view+0x9b8`, `view+0x138` and `view+0x140` **not at all**. Against `AI-PANEL-053`'s mode numbering (1 attack, 2 move, 3 join, 4 defend, 5 cast, 6 swarm, 7 join, 8 patrol) and its arming test `1 << (mode-1)`, the constant `0xef` is **every mode except 5** — so **attack, move, defend, swarm and patrol are armable for any selection whose bit `0x4` is clear, whatever classes it holds**, and cast alone is conditional. **Customisation:** a consumer that offers the attack mode over any owned selection matches the original exactly; one that filters the panel by unit class shows a control the original does not withhold

**Confidence.** High (the routine is read end to end at instruction level, and `tools/ordergates` re-reads its immediates out of the PE on both roots by **linear decode** from the entry point — it must reach `RET` or print where it stopped, and its output matches the listing instruction for instruction). The alternative ruled out is named and was this experiment's own H2: a per-class walk over the selection, which predicts reads of `view+0x9b8` that the routine does not contain
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### AI-PANEL-061

Its sole writer is `R0082`, the selection-maintenance routine and one of `R0093`'s two callers (`disp:144` = 125 hits / 63 owners; every store in the view is inside `R0082`). Reset at `L00691`, then: **`0x1`** set at `L00692` after `strcmp` against `"CUnit"`; **`0x200`** set at `L00693` when the current accepted loop object's `+0x20` is `0x17` or `0x18` (primary-only wording corrected by `AI-SPELLCAP-288`) (`L00694`/`L00695`) **and** its `+0x18` is non-zero (`L00696`); **`0x8`** set at `L00697` when `[obj+0x18c] & 1`; **`0x2`** set at `L00698` after `strcmp` against `L00699` = `"CAirUnit"`; **`0x20`** set at `L00700` after `strcmp` against `L00701` = `"CStructure"`; and an arm that forces the whole word to `0x8` (`L00702`).

**Bit `0x4` is ownership**: `L00703` tests the selection count `view+0x140` against zero, then `L00704` loads the primary object `view+0x138`, `L00705` reads its `CPlayer` from `+0x14`, `L00706` compares that with the local participant at `+0x9b4`, and **only when they differ** `L00707` sets bit `0x4`. So the panel is disabled outright for a selection the local participant does not own, and that is the same comparison `UNIT-VPLAYER-022` settles. The routine then notifies the panel `0x409` with `R0093`'s mask, or `0x40a` with **zero** when `view+0x144 & 0x24` (`L00708`)

**Confidence.** High for each bit's set site and its guard — every bit-set above is a cited instruction with its `strcmp` literal dumped from the image (`StrDump`) or its compare quoted, in one routine read whole. **Medium** that these are all the writers: `disp:144` is a displacement sweep and cannot see a store through a computed base, though a store through a register base plus `0x144` is the only form present. **Not established**: what bits `0x1`, `0x2`, `0x8` and `0x20` are *for* — only where each is set. Only `0x4` and `0x200` reach `R0093`

**Amended.** Partially retracted: the primary-only Cast clause. Primary-only200h setter clause: Local-28 is the current accepted loop object (refuted, EXP-0294). `retracted.md` holds the full entries.

### AI-MINIMAP-062

`R0233` is a three-stack-argument handler reading a `CPoint` at stack offsets `0x34` and `0x38` and never reading its first argument, reached through exactly one `.rdata` vtable slot (`callto:R0233` = 1 hit / 0 owners / 0 orphan, slot `L00709`). Like `R0211` it dispatches on the current cursor `[L00619]` against each registered cursor's own handle — and it covers **exactly the six cursors whose shipped art paths begin `s`**, while `R0211` covers exactly the eleven that do not, a clean and complete partition of `AI-CLICK-050`'s registry. The arm map, re-read out of the PE by `tools/ordergates` on both roots: `L00710` **sdefault** notifies `0x406` with the computed cell and issues **no order**; `L00711` **smove** calls `R0215` opcode `0x16`; `L00712` **sattack** calls `R0234`, and on a non-zero id `R0213` `0x19`, else `R0235` **`0x1a` swarm**; `L00713` **sdefend** calls `R0234`, and on a non-zero id `R0236` `0x1b`, else nothing; `L00714` **scast** calls `R0234` and **discards the result**; `L00715` **spatrol** calls `R0237` `0x1d`.

It ends by clearing `view+0x99c` and notifying `0x40b` (`L00716`…`L00717`), the same one-shot armed-mode clear the map click performs. **Five gates differ.** No drag threshold and no `view+0x140` return-to-select. **No `IsKindOf("CUnit")` on the attack arm** — the map click requires `view+0x990 != 0` *and* the class test (`AI-CLICK-050`); here a non-zero id is enough, so a `CStructure` under the pointer is an acceptable target. A different hit test: `R0234(x,y)` walks the object map at `view+0x9b8` and returns the **key** of the first entry whose `+0x78` is 0 and whose `+0x34`/`+0x38` equal the two arguments (`L00718`…`L00719`), against the hover's screen-rect walk. A different attack fallback: `0x1a` swarm, where the map click falls back to `0x16` move. And cast is unreachable from it

**Confidence.** High for the arm map and the five differences (the cursor-to-builder pairing and each builder's opcode byte are read out of the PE through its own section table on both roots with identical bodies; the handler is read end to end; the `s`/non-`s` partition is a property of the shipped art paths `AI-CLICK-050` already publishes). **The window's identity is inferred from those art names, not established** — what the image fixes is a second window whose click dispatches the small cursor set, and `s` for the scaled overview is the shipped naming, not a reading
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### AI-SURFACE-063

`callto:` over all thirteen builders on the repaired table (`evidence/rom-ordergates-enum.md`), 0 orphan throughout: `R0211` the map click (`R0215` twice, `R0213`, `R0235`, `R0236`, `R0237`, `R0090` twice, `R0091`, `R2088`, `R2084`); `R0097` the command panel (`R0223`, `R0224`, `R0102` — the three buttons `AI-PANEL-053` says issue at once); `R0233` the small-cursor surface (`R0215`, `R0213`, `R0235`, `R0236`, `R0237`). Nothing else calls any of them. The thirteenth, **`R0238`, writes opcode `0x22`** (the byte store at `L00720`) and enqueues through the same `R0217` (`L00721`), and has **7 hits over 6 owners** — `R0239`, `R0240` (twice), `R0241`, `R0242`, `R0243`, `R0244` — **none of them a cursor arm**.

That is very likely not an *order* at all: `AI-CMD-032` has the order space's `0x22` **empty** and the session space's `0x22` as the item move, `ITEM-CMD-007` names `0x22` as the one command that moves everything, and those six owners sit in the shop, inventory and town modules. **Inferred, not established** — no instruction setting `cmd+0x04` on this path was read, and `cmd+0x04` is what chooses the space. So `AI-CLICK-050`'s open Medium is bounded rather than closed: the attack builder has two callers and both are now read, the three cursor/panel surfaces are the whole of the order side, and the one remaining producer is probably the inventory's

**Confidence.** **Medium**, and the instrument is named: `callto:` decompiles only hits with a containing function and prints `.rdata` slots separately, so it shares the blind spots `AI-FILTER-001` and `AI-ORDER-031` state — which is exactly how `EXP-0109`, sweeping one builder, reported two surfaces where there are three. High for each individual pairing quoted, every one a cited `CALL` site with its target

**Amended.** The complete three-input-surface census and “`R0097` the command panel” clauses only — each direct caller/target pair it did list stands: Character-panel `WM_LBUTTONUP` reaches vmethod `R0240` through vtable `L00722 + 0x7c`; at `L00214` it calls `R0097` with `0x0a`, which remaps to Cast mode 5, so the next map click emits item `0x25/0x26` (superseded, EXP-0190). `retracted.md` holds the full entries.

### AI-FANOUT-064

`R0092(cmd, id)` is fourteen instructions read whole: `L00723` reads the member count, a **single byte** at `cmd+0x12`; `L00724` compares it with `0xfd` and `L00725` branches to `L00726` when it is at least that (signed), **which returns having done nothing**; otherwise `L00727` stores the id as a 16-bit value at `cmd+0x13+2*count` and `L00728` increments the byte. So the ids are `u16` at `cmd + 0x13 + 2*i`, the record reaches `0x13 + 2*253 = 0x20d` bytes at most, and past 253 there is **no error, no truncation flag and no wrap** — the append is a no-op. **Customisation:** this is the one place where a consumer with an unbounded selection and the engine diverge without a symptom. The selection container itself is a growing `CMapWordToPtr` (`AI-SELECT-065`) and imposes no such bound, so a selection above 253 is representable and only the *order* truncates; and because the builder iterates in hash-bucket order, **which** 253 survive is decided by the hash, not by selection order

**Confidence.** High (the routine is read end to end; the guard, the stride, the base displacement and the increment are four cited instructions in one fourteen-instruction body, and `0xfd` is an immediate in the image, not a derived figure). The alternative ruled out — that the ceiling is a buffer size checked elsewhere — is excluded by the guard being inside the append itself and by the count's width being one byte

### AI-SELECT-065

The layout is fixed by three independent uses rather than asserted: `R0213`, `R0234` and `R0212` all reach it by adding `0x9b8` to the view pointer and then read `+0x04` as a bucket **array of pointers** (the element read at `L00729`), `+0x08` as the **bucket count** they compare an index against, and the assoc's `+0x00` `pNext` / `+0x04` `nHashValue` / `+0x08` `key` (a `WORD`) / `+0x0c` `value` — MFC's `GetNextAssoc` walk, bucket chain then forward scan. `m_nCount` is therefore `0x9b8 + 0xc` = **`view+0x9c4`**, and that is exactly the field all three feed to the inlined `GetStartPosition` idiom that turns a non-zero count into all ones (`L00626`, `L00730`, `L00731`), which yields `-1` for a non-empty map and `0` for an empty one.

`R0212` uses the same object as a **`Lookup`** of an on-screen id (`L00732`), which an array of selected units could not serve. The order builder walks the whole map and calls `R0092` with `word[obj+0x04]` **only when `[obj+0x7c]` is non-zero** (`L00733`…`L00734`). **So `AI-CLICK-051`'s conclusion stands and its container does not**: every selected unit travels by id, but `+0x04` is not the data pointer and `+0x08` is not the count

**Confidence.** High for the container's identity and layout (three routines reach it independently, each read at instruction level, and the `m_nCount` displacement is confirmed by arithmetic — `0x9b8 + 0xc` — against the field the `GetStartPosition` idiom actually loads, which is a check the layout could have failed). **Medium** for the writer set: `disp:9b4` gives `view+0x9b4` two stores over 123 hits / 78 owners / 0 orphan, but `disp:9b8` returns only **3** hits and none in the input path, because adding the constant 0x9b8 to a base carries it as an immediate — the readers are visible only to `imm:9b8` (51 hits). **Inferred, not established**: that `obj+0x7c` means *selected*. It is the flag the builder filters on; `disp:7c` is 790 hits over 433 owners, too common a displacement to close by sweep

## Facing before a strike

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-FACE-066 | An attacker must already be facing its victim, and the swing neither tests facing nor produces it — the mirror of `AI-RETAL-056`, where the victim's turn IS the consequence. | High / Medium | ● active (amended) | [EXP-0111](../experiments/EXP-0111-blow-display/) |
| AI-FACE-067 | The turn-to-face producer set is at least six routines (11 call sites of `R0056`), and none read is on the swing path. | High / Medium | ● active (amended) | [EXP-0111](../experiments/EXP-0111-blow-display/) |

### AI-FACE-066

`R0245` (the swing start) and `R0246` (the strike) are both read end to end and their complete guard sets are: both actors non-null with a non-null `+0x14` `Player`; the attacker's health `> 0`; and, in the strike only, `reach >= R0247(attacker, target)` (`L00735`..`L00736`, `HERO-REACH-025`). **Neither contains a facing test, a facing write, or a call to any turn routine.** The test lives one level up, in `R0041`: `R0051` returns the direction from self to target and it must equal `mover+0x00`, the current facing (the mover at `+0x154` is read at `L00737` and its facing byte compared at `L00577`), and only if it does is the edge-to-edge distance compared against the reach argument (`L00578`).

Every route into the attack act-state passes through it. An `EnumRefs` pattern sweep for 4-byte stores of the immediate 3 into `+0x54` is **7 hits / 4 owners / 0 orphan**; three are float stores to a stack slot at offset `0x54` in `R0248` and the four real ones are `L00092`, `L00738`, `L00739` and `L00740`. `L00739` sits inside `R0005`, whose `callto:` is **1 hit** and whose only caller `R0004` runs `R0041` five instructions earlier (`L00741`). `L00738` and the pursuit arms at `L00101` and `L00581` each sit on the taken branch of their own `R0041` (`L00742`, `L00743`, `L00744`). And `L00092` is **not a fourth route but the latch**: it is the arm for `ord+0x9 == 1` — byte table `L00090` into jump table `L00091`, decoded out of the image in `evidence/tables.md` — and `ord+0x9 = 1` is written only by those same passing branches (`L00745`/`L00103`, `L00746`, `L00747`). The test is re-run every tick, so a victim that circles its attacker drops it back to `R0042`, which inside the stop distance **turns to face and stands** (`AI-PURSUE-040`). **So facing is a precondition, maintained continuously by the approach, and never a consequence of the blow**

**Confidence.** **High** for the two routines' guard sets and for `R0041`'s two tests — all read end to end, every clause a named instruction, and the latch arm settled by a static table rather than by inference / **Medium** for *every route*: the store-form sweep is `AI-STATE-043`'s instrument and inherits its blind spot, a state restored from a save or produced by arithmetic in place being invisible to it, and `AI-PROGRESS-034`'s second switch was read only at the one arm this row needs

**Amended.** `L00738` is pending order 2 and has no `R0041` call; `L00742` is cast arm 8. The other `ord+9 = 1` writers follow a passing gate (`AI-411`). The sentence "The test is re-run every tick ... facing is a precondition, maintained continuously by the approach" holds only at progress 0: once `ord+9 == 1` no gate runs (`L00092`) and the body is not re-turned (`AI-408`). The guard sets and the two tests stand.

### AI-FACE-067

`R0056(actor, dir)` is the only routine in the image that aims an actor: it stores the wanted facing at `mover+0x01` (`L00748`), takes the shortest arc against `mover+0x00`, and — when `mover+0xa0` is 0 and that arc is at most `0x21` in either direction (the compares against `0x21` at `L00749` and `L00750`) — **snaps** `mover+0x00 := mover+0x01` and raises `mover+0xa4`; otherwise it hands off to `R0249`, which is what actually issues a turn. `EnumRefs callto:R0056` is **7 hits / 6 owners / 0 orphan**: `R0205` at `L00751` (the idle-turn arm whose `rand()` gate `AI-RETAL-056`'s struck flag skips), `R0016` at `L00752`, `R0250` at `L00753`, `R0043` at `L00754` (the approach's stand-and-face), `R0178` twice, and `R0054`.

**None of the six is `R0245`, `R0246` or `R0037`**, and none is reached from them. Two consequences a consumer must carry: the snap threshold means an actor already inside its reach usually corrects its facing within one tick and without any visible turn animation, since a turn under 45 degrees never becomes action code 5; and the turn belongs to the mover or to the order machine, never to the strike, so a consumer that turns the attacker as part of the swing has invented a coupling the original does not have

**Confidence.** **High** for the routine — read end to end, both immediates and the snap transcribed — and for the caller set, a `callto:` enumeration on the repaired function table reporting 0 orphan, which is the shape that can carry an *only* / **Medium** for *none is reached from them*: that is a reading of three call chains rather than an enumeration, and a `callto:` sweep cannot see an indirect call

**Amended.** A raw `E8` scan finds 11 call sites of `R0056`, not 7 (`AI-411`); "none is reached from the swing path" is Medium (raw `E8`/`E9`, blind to indirect calls). The routine's snap stands.

## Group sight, target scoring, guard and Stand Ground

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-GROUPSEE-068 | A group sees as one animal: the candidate list is built from a sight map every member stamps into, swept over the whole world actor list, filtered by the *first* member's diplomacy row — and the dead are parked rather than dropped. | High / Medium | ● active | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-SCORE-069 | `ord+0x20` is the group-assigned target; it is rewritten for every member on every evaluation and cleared when nothing scores, so nothing at the group layer is sticky. | High / Medium | ● active | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-PREF-070 | The group's target choice runs through a 4×4 preference matrix of compile-time constants at `AImanager+0xb94`, and two of its cells are an absolute veto. | High / Medium | ● active | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-COST-071 | The cost is `(distance term << 8) + turn cost`, so distance decides and facing breaks the tie — and a ranged candidate is scored as though it could not move. | High / Unknown | ● active | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-REACH-072 | The second scorer refuses anything past reach outright, so group order 3's arm never assigns a target beyond reach; the executor's pathing between assignments is not covered. | High | ● active (amended) | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-327 | The reacquisition routine a refused pursuit falls into selects its new victim by a footprint-blind Chebyshev cell distance, and the footprint-aware in-position test `AI-PURSUE-040` names re-enters only after a winner is already chosen… | High | ● active | [EXP-0373](../experiments/EXP-0373-refused-route-attacker/) |
| AI-328 | `R0004` never writes the order's target field on a failure exit: `ord+0xc` has exactly one `MOV` destination in the whole 268-instruction body, and it sits inside the candidate-found branch. | High | ● active | [EXP-0373](../experiments/EXP-0373-refused-route-attacker/) |
| AI-329 | The complete outgoing call inventory of both group-level candidate scorers is four direct targets and six virtual dispatches, and none of the four direct targets is the route search… | Medium | ● active | [EXP-0373](../experiments/EXP-0373-refused-route-attacker/) |
| AI-FLIER-073 | A melee ground creature never auto-selects a flying target, a ranged one always can, and shipped content places fliers on thirty of the thirty-eight maps. | High / Medium | ● active | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-GRPGUARD-074 | Group order 1, the arm 95.6 % of shipped hostile placements run, read end to end — and its walk-home target is the member's own post, not the group's centre. | High | ● active (amended, partially retracted) | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-RADFREEZE-075 | A group's notice radius is frozen at the geometry it had when guard was last issued, while the circle it draws follows the group every tick. | High / Medium | ● active (amended) | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-STAND-076 | Group order 3 is Stand Ground, not attack: its scorer refuses anything past reach and its arm issues no walk order; the executor keeps pathing to the last victim until the next evaluation, a stand or a refused route (Medium, `AI-353`). | High / Medium | ● active (amended) | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-REISSUE-077 | What re-issues a pursuit is the group arm itself, unconditionally, every AI tick — and a group with nothing but corpses in sight keeps attacking a corpse. | High / Medium | ● active | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-DEADROLL-078 | The guard radius roll exists three times in the image and two of the copies are dead. | High / Medium | ● active | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-GATE-079 | `AI-GUARD-007`'s class gate is the mage bit, while `actor+0x0e == 0x18` remains inside an actor-state arm with no located writer. | High / Medium | ● active (amended, partially retracted) | [EXP-0112](../experiments/EXP-0112-ai-execution/), [EXP-0133](../experiments/EXP-0133-person-health/), [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |
| AI-CLOCK-080 | Engagement runs on two clocks, and the decision one is the slow one. | High / Unknown | ● active | [EXP-0112](../experiments/EXP-0112-ai-execution/) |
| AI-LOS-081 | The line-of-sight test is an accumulator march, not an altitude comparison — it carries a budget from cell to cell and stops when the budget runs out. | High / Unknown | ● active | [EXP-0112](../experiments/EXP-0112-ai-execution/) |

### AI-GROUPSEE-068

`R0110(AImanager, group)`, which takes one stack argument besides the manager, read end to end (`evidence/rom-groupscan-excerpt.md` §1). Five phases. (a) `R0251` empties scratch list A; `R0145(world+0x58ee8)` clears the `0x10000`-byte visibility map **once** at `L00276`, before the member loop, and then each member's sight is stamped with `R0135(fog, member, R0019(member+0x10))` — so a candidate visible to any member is visible to every member. (b) For a member whose owner is not a human participant (`[member+0x14]+0x28 != 0`, `UNIT-OWNER-009`) and whose `ord+0x58` is non-zero, the remembered attacker's cell is **incremented** in the same byte map (`L00755`, `L00756`), `ord+0x5a` is incremented, and at `ord+0x5a > 0x14` — **20** — `ord+0x58` is cleared (`L00757`/`L00758`): the group-side twin of `AI-SIGHT-006`'s per-actor memory.

(c) `world+0xa4554`, the whole on-map actor list (`AI-LIST-003`), is walked head to tail and every actor whose own cell byte is non-zero is appended to A (`L00759`…`L00760`) — no neighbourhood, no spatial index. (d) `R0021(AImgr, firstMember, listA, 0)` — `AI-FILTER-001`'s diplomacy predicate with the group's **first member** as the decider, which also fixes whose `me+0x70` the See-invisible exception consults. (e) Every candidate with `word[cand+0x94] < 1` is unlinked into list B at `+0xbc4` (`L00761`…`L00762`); then **if A is empty and B is not, the whole of B is moved back into A** (`L00763`…`L00764`), so a group with only corpses in view still has a candidate list. The routine overwrites its own argument slot at `L00765` with `AImgr+0xba4`, which is why `[ESP+0x24]` names the group before that instruction and the scratch list after it — read the other way it inverts phase (e). Callers: `EnumRefs callto:R0110` = **5 hits / 5 owners / 0 orphan** — group orders 1, 2, 5, the dispatcher's inline order-3 arm, and `R0109` from outside the module

**Confidence.** High (one routine read end to end, every phase a cited instruction; the caller count is a printed enumeration on the repaired table). **Medium** that the caller set is complete — `callto:` sees direct calls and `.rdata` code pointers, and every dispatch table in this area lives in `.text` (`AI-DEAD-036`)

### AI-SCORE-069

`R0177` and `R0153` are **the same routine twice** — instruction for instruction identical apart from the scorer called at `L00766` (`R0225`) against the one called at `L00767` (`R0252`) (`evidence/rom-groupscan-excerpt.md` §2). Both open by testing the byte at `group+0xc` and the byte at `AImgr+0xbb4` for zero — the **low byte only** of each count — and on either being zero walk the members writing `ord+0x20 = 0` (`L00768`, `L00769`). Otherwise, per member, over every candidate: `cost = f(AImgr, member, cand)`, taken when `cost < best` **strictly** (the signed compare and branch at `L00770`), so the first candidate wins a tie and list order decides; `best` is seeded `0xffffff` (`L00771`) and, if it survives, `ord+0x20 = 0` (`L00772`), else `ord+0x20 = winner` (`L00773`).

**The complete surface of `ord+0x20`**, from a full disassembly of `L00339..L00340` read for every dword access at displacement `0x20` (stack slots and vtable `CALL`s removed, each survivor classified by its own base register): **6 writers**, all in these two routines, and **4 readers** — `L00774` group order 1, `L00444` order 2, `L00775` order 5, and `L00776` inline in the dispatcher for order 3, which belongs to no named routine and is the hit a routine-by-routine reading loses. The discriminator against `grpAI+0x20` is width: the group order byte is a **byte**, this is a **dword**

**Confidence.** High for the structure and for each cited site (both routines read whole; the seed and the `JGE` fix the tie rule). **Medium** as a completeness claim: the instrument is module-scoped and blind to the order block reached at another displacement, to a wholesale copy of it, and to anything outside `L00339..L00340`

### AI-PREF-070

All sixteen bytes are written once each, in the AI manager constructor `R0124`, at consecutive addresses `L00571`…`L00777` (`evidence/rom-scorer-excerpt.md` §5). Their values come from three registers none of which is rewritten between its load and the block: one holding 2 (loaded at `L00570`, the same register `AI-SPREAD-038` reads at `L00569`), one holding 1 (`L00778`) and one holding 0 (zeroed at `L00779`); the remaining four stores carry the immediate `4`. Indexed `M[member domain][candidate domain]` — the stride-4 axis is the member's, fixed by the address computation at `L00780`, which adds `0xb94` plus four times the member's `vt+0x20()` result to the candidate's value, which is what rules out a flat 16-entry table.

Rows, member domain first: `0: 2 1 2 4` · `1: 4 2 1 0` · `2: 2 1 4 0` · `3: 2 1 2 4`. The domain is `actor+0x4a` through `vt+0x20` (`MOVE-DOM-024`), 3 being the image's own word for flier. **A `0` cell returns the selection loop's own seed `0xffffff`** (`L00781`), and the loop compares `JGE`, so a vetoed candidate can never be taken. **The instrument lesson is the row's own**: `EnumRefs disp:b94` returns 19 hits over 12 owners and exactly **one** store in the AI module, which alone reads as "one byte, set to 2"; the other fifteen bytes are at fifteen other displacements. `docs/INSTRUMENT.md` rule 4 says to reduce a common displacement by the field's **width** — the same trap on the *array* side is that a sixteen-byte table has sixteen displacements and a sweep for one of them is not a sweep for the table

**Confidence.** High (sixteen cited stores in one routine, three register provenances each traced to a load with no intervening write, and the index law discriminated by a stride the rival model does not have). **Medium** that no other instruction writes the block: the sixteen `disp:` sweeps cannot see a `REP MOVSD` of the containing object

### AI-COST-071

`R0225(AImgr, member, cand)`, which takes two stack arguments besides the manager, read whole (`evidence/rom-scorer-excerpt.md` §3), lower is better. `d = R0167(world, memberCell, candCell)`, the Chebyshev distance in cells, as a byte. `k = cand->vt+0x20()`; **when `k == 1` and `cand+0x12c > 1`, `k := 0`** (`L00782`…`L00783`) — an ordinary ground creature holding a reach above 1 is indexed as the immobile column, which is what makes `M[1][0] = 4` a *prefer the archer* rule for a ground melee decider. A member with reach `> 1` takes `pref = AImgr[0xb94 + k]` — **row 0 only, the member's own domain never consulted** — and `d <= reach ? d := 1 : d += 1 - reach` (`L00784`…`L00785`); a member with reach `<= 1` takes `pref = AImgr[0xb94 + k + 4*(member->vt+0x20())]` and leaves `d` alone.

Then `cost = (d << 8) + R0115(world, mover[0], R0051(world, member, cand))` (`L00786` to `L00787`) — the same direction/turn pair `AI-ACQUIRE-002` names, folded into one scalar instead of a two-key comparison. Three modifiers follow: `pref == 1` → `cost += cost/2`, `pref == 4` → `cost -= cost/4`, both gated on `cand->vt+0x1c() < 2` (token size 1, `MOVE-DOM-024`'s footprint getter) **and** `word[cand+0x88] >= 0xf` (Mind, `R0253`'s operand); and `cand+0x144 & 0x100000` → `cost += 0x7f` clamped at 1 (`L00788`…`L00789`), `+0x144` being the active-spell bitmask `R0254` indexes, so **spell id 20 makes its bearer the least attractive target on the map**. The actor constructor `R0160` defaults `word[actor+0x88]` to **20** (the 16-bit store of `0x14` into `+0x88` at `L00790`) and the `Data.bin` Units collection carries no Mind column, so the Mind half of the gate is satisfied by default everywhere and only token size discriminates

**Confidence.** High (one routine read end to end, every arithmetic step and both table indexings cited; the three field identifications are each another ledger's published operand). **Unknown** which spell carries id 20 — the bit was read, `claims/magic.md`'s id space was not cross-walked

### AI-REACH-072

`R0252(AImgr, member, cand)` (`evidence/rom-scorer-excerpt.md` §4) is `R0225` with three differences. The candidate's domain is forced to 0 whenever `cand+0x12c > 1` — for **any** domain, not only domain 1 (`L00791`…`L00792`). **the compare of the distance term with 1 at `L00793` — a distance term above 1 returns `0xffffff` immediately** (`L00794`), before the turn cost is even folded in, so this scorer can only ever name a target the member could strike where it stands: adjacent for a melee member, within reach for a ranged one. And the preference multipliers are ungated and coarser — `pref == 1` → `cost <<= 1`, `pref == 4` → `cost >>= 1`. Its single consumer is `R0153` (`callto:R0252` = **1 hit / 1 owner / 0 orphan**), whose single consumer is the dispatcher's inline order-3 arm (`callto:R0153` = **1 / 1 / 0**, `L00795`). `R0225` has three (`callto:R0225` = **3 / 3 / 0**): `R0177`, the player attack order `R0006`, and the script dispatcher `R0156`. `MOVE-DOM-024` lists both routines as located and not read; they are read here

**Confidence.** High — both routines read end to end and all four caller counts are printed enumerations with 0 orphan hits

**Amended.** The refusal stands: at each evaluation the arm assigns no target past reach. The consequence that group order 3's members therefore never chase is narrowed. The executor keeps pathing to the victim assigned at the previous evaluation until the next evaluation, a stand or a refused route replaces the order (Medium, `AI-350`, `AI-353`). `retracted.md` holds the entry.

### AI-327

The reacquisition routine a refused pursuit falls into selects its new victim by a footprint-blind Chebyshev cell distance, and the footprint-aware in-position test `AI-PURSUE-040` names re-enters only after a winner is already chosen — never during selection. `R0004` (`AI-ROUTE-045`'s consequence of `mover+0x98`) builds its candidate list in one pass over `world+0xa4554`: `R0255(self+0x10, cand+0x10)` at `L00796` is the sole distance call, kept when `<= actor+0x12c` (`L00616`/`L00617`) — the same reach bound `AI-ROUTE-045` already named without identifying the routine, and the same test `HERO-TARGET-024`'s `EXP-0074` amendment cites by its bound load at `L00797`. That amendment already calls this helper a whole-cell Chebyshev here, and `UNIT-STRUCTUSE-090` already calls it an unexpanded distance helper at an unrelated site; neither states its structure.

It is 28 instructions, a leaf with no `CALL` at all, so it never dispatches `vt+0x1c` (the footprint getter `TERR-PASS-051`/`MOVE-DOM-024` name, `R0256` per `TERR-MOVE-055`) or any other virtual — unlike the edge-to-edge `R0036` that subtracts both actors' token sizes for `AI-PURSUE-040`'s in-position test and `AI-ACQUIRE-002`'s standing pick. The only exclusion **by identity** during selection is the pursuer's own self (the compare against the actor argument at `L00798` and the branch at `L00799`) — the *previous* victim is not tested against and can be re-selected. Identity is not the only narrowing: a candidate is also dropped unless it is no farther than the nearest candidate seen so far (the compare against the running best at `L00800` and the branch at `L00801`, the running best updated at `L00802`) — the gate `HERO-TARGET-024`'s amendment states as "keeps every candidate that was `<=` the running minimum when it was seen" — the diplomacy filter unlinks non-hostiles (`R0021` at `L00803`, `AI-FILTER-001`), and the living/dead partition at `L00804` (`cand+0x94 < 1`, `AI-FILTER-001`'s health word) diverts corpses to `AImgr+0xbc4` (`L00805`/`L00806`) unless nothing living survives, when they are restored and one becomes the target (`L00807`/`L00808`/`L00236`) — the same corpse fallback `AI-ACQUIRE-002` already names in this routine.

After the filters and turn-cost scoring, the winner is written to `ord+0xc` (`L00121`) and `ord+0x08 = 6`, and only then does `R0041(self, target, reach)` run once (`L00741`) — **the identical routine `AI-PURSUE-040` cites for ordinary pursuit's in-position test**, footprint-aware by way of its internal `R0036` call — to choose between an immediate strike (`actor+0x54 = 3`) and a fresh walk toward the new target (`actor+0x54 = 1`, via `R0250`, not read here). So footprint size decides nothing about *which* actor is picked; it only gates whether the actor already holding that pick can swing at it this tick

**Confidence.** High (the sole distance call, its reach comparison, the running-best and self-exclusion compares, the diplomacy-filter call and the single downstream in-position call are each a cited instruction from a 268-instruction body read whole and hashed in `evidence/instructions.tsv`; the call enumeration is complete rather than direct-only, because the body contains **no indirect call at all** (`reacq_indirect_calls` is empty), so no dispatch this probe cannot resolve can reach a footprint, route or occupancy routine during selection; the footprint-blindness of `R0255` is the absence of any `CALL`, direct or indirect, in its own 28-instruction body. What is new against `HERO-TARGET-024` and `UNIT-STRUCTUSE-090` is the structure, not the identity: the helper's zero-call leaf body, the completeness of the caller's own call enumeration, and the placement of the one footprint-aware test after the write rather than during selection. The rival this rules out is "the reacquisition scan reuses the same footprint-aware metric as ordinary pursuit" — it does not, for selection; it does for the tick that follows selection)

### AI-328

The write is `L00121`, guarded by the branch at `L00809` to `L00810` on a null best-candidate pointer. `AI-GUARD-007` already publishes that site and the two order-kind writes that follow it on the found path — `L00122` sets `ord+0x08 = 6`, and `L00123` sets it back to `0` when that row's suppression clause fires (`actor+0x4c & 4` and `actor+0x12c < 2` and `Player+0x28 == 0`, `L00811`…`L00812`) — and reads it the same way: the pointer is left behind and the order goes idle. This row adds the other exit. The no-candidate tail beginning at `L00810`, where `AI-ROUTE-045`'s two failure outcomes are written (`ord+0x08 = 0xb` for an AI owner at `L00124`, `0` for a human one at `L00125`), contains no store to `ord+0xc` either, and the single-write enumeration over the full body says there is no third site: both the suppressed success and the total refusal leave the stale pointer.

So a repeatedly-refused pursuit does not scrub its own memory of who it was chasing — the order *kind* moves on every refusal (`AI-ROUTE-045`), the *target* field only on a fresh success. `mover+0x98` itself — the flag this routine answers — is a plain one-shot boolean with exactly three setters and one consumer (`AI-ROUTE-045`, `MOVE-GATE-039`), not a counter or a timeout: nothing in either routine accumulates a refusal count, and the next route search attempt re-decides from scratch

**Confidence.** High for what this routine does, which is what the row claims (the single write site, its guard and the no-candidate tail's absence of any `+0xc` store are each a cited instruction over the routine's full 268-instruction body, and the body has no indirect call through which a write could be reached unseen; the flag's setter/consumer count is `AI-ROUTE-045`/`MOVE-GATE-039`'s own `EnumRefs` enumeration, cited rather than re-run). The rival this rules out is "a cancelled pursuit order is torn down wholesale, the way state `1`'s teardown to `0xc` is" (`AI-ROUTE-045`) — it is not: only the order *kind* byte moves inside this routine. **Bound:** the enumeration covers this routine's own instructions, not its callees. The AI-owner tail calls `R0205` (`L00813`), which was not read here, so this row does not say the pointer survives the tick — only that `R0004` does not clear it

### AI-329

The complete outgoing call inventory of both group-level candidate scorers is four direct targets and six virtual dispatches, and none of the four direct targets is the route search, the occupancy-block builder or either dynamic-plane claim setter. `R0225` (`AI-COST-071`) and `R0252` (`AI-REACH-072`), each read whole and hashed, make ten direct calls between them to exactly four addresses: `R0019` (the packed-cell getter, twice in each body), `R0167` (the Chebyshev cell distance `AI-COST-071` names), and `R0051`/`R0115` (the direction and turn-cost pair `AI-ACQUIRE-002` names). They make six indirect calls: `vt+0x20` at `L00814`, `L00815`, `L00816`, `L00817`, and `vt+0x1c` at `L00818`, `L00819`.

The split is asserted exhaustive rather than filtered — `scorer_call_inventory` in `evidence/measurements.json` fails the run unless direct plus indirect accounts for every `call` instruction in each body. Neither direct set contains `R0053` (route search, `MOVE-SEARCH-001`), `R0020` (occupancy-block builder, `AI-FILTER-001`'s producer), or `R0257`/`R0046` (the dynamic-plane claim setters, `MOVE-CLAIM-007`). The six dispatches were not followed: `TERR-MOVE-055` publishes `vt+0x1c` and `vt+0x20` as two-instruction field getters over `actor+0x49`/`actor+0x4a` for the simulation-actor family and as constant stubs for its 17 siblings, and `AI-COST-071` and `AI-REACH-072` read these same sites as the footprint and domain reads, but nothing in this experiment establishes the receivers' class, so no dispatch implementation is excluded here. Against `AI-COST-071`, which already names five of `R0225`'s call edges by what they read, what this row adds is the inventory as a closed set over both bodies and the four-address negative, which neither prior row states

**Confidence.** **Medium.** The direct half is exhaustive and structural: every `call` instruction in two fully disassembled, hashed bodies (116 and 104 instructions) is accounted for, and none of the four route/occupancy addresses is a direct target. The half that would carry High is missing. The rival — "the group re-score compensates for the per-actor layer's route-blindness by preferring a reachable candidate" — needs the six virtual dispatches to read nothing about reachability, and a dispatch's implementation cannot be excluded by a structural read of its caller. `TERR-MOVE-055`'s vtable enumeration narrows the residue to a receiver outside the simulation-actor family, which is why this is Medium and not Unknown

### AI-FLIER-073

`M[1][3] = M[2][3] = 0` (`AI-PREF-070`) is the veto, and three constructor facts decide who it binds. The actor constructor `R0185` defaults `actor+0x4a` to **1** and `actor+0x49` to **1** (`L00820`, `L00821`; `re:` over the byte-store form returns 2 hits / 2 owners / 0 orphan for each). The `Data.bin` **Humans** collection ships **no** `MovementType` cell at all — all **210** named rows on both roots carry the absent-cell sentinel, so the streamer skips the store (`DAT-SCHEMA-004`) and every human actor, the player's hero included, is domain 1. The **Units** collection splits **40 / 8 / 8** over domains 1 / 2 / 3, identical on both roots. And the ranged branch of both scorers indexes **row 0**, whose four entries are `2 1 2 4` and contain no zero — so reach, not class, is what decides whether a creature will chase something in the air.

Corpus, resolving every type-6 record through `ALM-CLS-052`'s rule (`tools/aiexec`): **EN 701 of 8094 placements are domain 3, over 30 of 38 maps; RU 261 of 3991, over 25 of 34** — the RU shortfall being the four EN-only loose maps and the terrain-only `Horror.alm` (`AI-ROOT-049`), not a rule difference. Domain histograms EN `1: 4878, 2: 1008, 3: 701` plus 1507 human placements with no cell; RU `1: 2597, 2: 494, 3: 261` plus 639

**Confidence.** High for the code half (the matrix cell, the veto's mechanism and both constructor defaults are cited instructions, and the ranged branch's row-0 indexing is a single `MOV` with no member term). **Medium** for the corpus half — `METHODOLOGY.md` caps a census there, and it also inherits `ALM-CLS-052`'s resolution rule rather than observing the resolver. **Not established**: whether a *reach* above 1 is ever streamed onto a domain-1 creature in a way that would put it on the ranged branch — `actor+0x12c`'s filler was not traced this round

### AI-GRPGUARD-074

`R0024(AImgr, group)` (`evidence/rom-grouparm-excerpt.md` §6), in order: `R0141` the geometry pass (`L00822`); `R0110` the candidate list (`L00823`); then, with the working radius `grpAI+0x2d` and `grpAI+0x28` held, every candidate whose `R0167(world, candCell, centroid)` exceeds the working radius is unlinked (`L00372`…`L00374`) — the clip `AI-RADIUS-014` names; then `R0177` (`L00824`); then `AI-GUARD-021`'s latch and radius re-roll. The per-member tail (`L00825`…`L00826`) has exactly three outcomes: **`ord+0x20 != 0` → `R0009(AImgr, member, target)`** (`L00336`), the engage; else the member's current cell is compared against **`ord+0x00`, its own post** (the compare at `L00827`) and, differing — or `R0040` reporting it not idle — **`ord+0x08 = 1`, `ord+0x0a = ord+0x00`** (`L00338`/`L00828`), walk to the post; else, at the post and idle, a human participant's unit goes to `R0015` (heal) and an AI owner's to **`ord+0x08 = 0xb`** (`L00337`), the *idle turn* order arm of `AI-ORDER-039` — not the guard state `actor+0x50 = 0xb` — followed by `R0113` when `member+0x4c & 4`.

This confirms `AI-GUARD-012`'s three cited addresses exactly and corrects its gloss in one place: the arm does not send a member home to `grpAI+0x28`; the centroid is the **candidate clip's** origin and the post is the walk's destination, and for a group under order 1 from load nothing has ever written `ord+0x00` (`AI-POST-042`'s writer set is complete and contains no group-order path) *[EXP-0125]: the parenthesis is **refuted**. The load-time guard setter `R0125` writes `ord+0x00` itself, `0x76` bytes before the `grpAI+0x20 = 1` that puts the group under order 1 — `L00829` from the member's current cell, `L00830` from `mv+0x06` for one in motion. The arm's three outcomes and every address above stand; what fails is that the post is unwritten. See [`retracted.md`](retracted.md) and `AI-POST-095`.*

**Confidence.** High (one routine read end to end, every branch and store a cited instruction, and the two objects told apart by their base registers — `[EDI + 0x3c]` the group AI record against `[ESI + 0x158]` the member's order block)

**Amended.** The parenthetical clause only — the arm is read correctly end to end and its three outcomes stand: The load-time guard **setter** writes it, in the same routine and 0x76 bytes before the `grpAI+0x20 = 1` that puts the group under order 1: the 16-bit store at `L00829` from the member's current cell, or `L00830` from `mv+0x06` for a member in motion (refuted, EXP-0125). `retracted.md` holds the full entries.

### AI-RADFREEZE-075

`R0141(AImgr, group)` (`evidence/rom-grouparm-excerpt.md` §8) is `R0140` with the member count taken from `[group+0xc]`: it sums each member's fine coordinates, divides, and stores the fine centroid `grpAI+0x24` (`L00831`) and the cell centroid `grpAI+0x28` (`L00832`), then over the members the maximum Chebyshev spread `grpAI+0x2a`, the maximum member sight `grpAI+0x2b` and the maximum of the two summed `grpAI+0x2c` (`L00833`/`L00834`/`L00835`). It is `R0024`'s first act and its only caller (`callto:R0141` = **1 hit / 1 owner / 0 orphan**), so it runs on every guard tick of every group. **And `grpAI+0x2c` is then never read again.** Whole-image `disp:2c` restricted to byte accesses and stripped of stack slots returns **19 hits**, of which four touch this record: two writers, `L00359` (`R0140`) and `L00835` (`R0141`), and two readers, `L00836` and `L00837`, **both inside the guard setter `R0125`**, which runs at map load, on the player's guard command and on the script's `Par0 = 1` (`callto:R0125` = 6 / 6 / 0).

What the arm actually rolls from is `grpAI+0x38`, whose three writers are all in that same setter (`L00838`, `L00839`, `L00361`) and whose four readers are the two live roll sites and two dead ones (`AI-DEADROLL-078`). So the radius is a load-time constant plus `AI-GUARD-021`'s ±1 jitter, and a group that spreads out, loses members or crosses the map keeps it. This **amends `AI-RADIUS-014`**, whose Medium rested on `R0140` being the only producer: there is a second, and it is the one that runs per tick *[EXP-0125]: the ±1 this row quotes is confirmed to the instruction and its distribution fixed — `{−1, 0, +1}`, each about a third (`AI-JITTER-103`, `AI-RANGE-102`).*

**Confidence.** High (both writer sets and both reader sets are complete enumerations with their instruments named, and the per-tick producer is fixed by a `callto:` with 0 orphan). **Medium** that the enumeration is complete for this record: a byte-displacement sweep cannot tell one object from another, and the attribution rests on each site's own base register

**Amended.** The ±1 the row quotes is confirmed to the instruction and its distribution fixed, `{−1, 0, +1}` each about a third (`AI-JITTER-103`, `AI-RANGE-102`). The row amends `AI-RADIUS-014`, whose Medium rested on `R0140` being the only producer of `grpAI+0x2c`; there is a second, `R0141`, and it is the one that runs per tick.

### AI-STAND-076

The arm is inline in the dispatcher at `L00307`…`L00840` (`evidence/rom-grouparm-excerpt.md` §7): `R0110` (`L00841`), `R0153` (`L00795`), then per member `ord+0x20 != 0` → `R0009` (`L00842`); else an AI owner gets `ord+0x08 = 0xb` (`L00158`) plus `R0113` on `member+0x4c & 4`, and a human participant's unit `R0015`. **There is no radius clip and no walk of any kind**, and `AI-REACH-072`'s scorer vetoes every candidate past reach — so the arm's whole behaviour is *hit what is already next to you, otherwise turn on the spot*. Three independent names agree and the ledger's label does not: the mission-description key that selects this setter is spelled **`StandGround`** (`AI-AUTHOR-015`, `R0062`), the shipped editor catalogue names `Par0 = 3` **`Stand Ground`** (`AI-GROUPCMD-020`, 24 nodes over the corpus), and the arm stands its ground.

`AI-ORDER-010`'s "attack" and `AI-AUTHOR-015`'s "aggressive" are **amended** to Stand Ground. The consequence a consumer must not invert: the load walk gives order 3 to the groups of type-5 slot 0 — the player's own `Self` — and order 1 to everyone else's (`AI-AUTHOR-015`, `AI-CENSUS-046`: EN 8085 / 9 / 0), so *the player's units stand and the scenario's units watch a circle*, which is the opposite of what the two labels suggest

**Confidence.** High for the behaviour (one arm read end to end against a PE-read dispatch table, and its scorer read whole). **Medium** for the name — three shipped or authored spellings agree, but a name is not a measurement, and the setters also write `actor+0x50 = 0xc` (`AI-AUTHOR-015`) which only group order 0 would ever evaluate. **Medium** for the executor clause of the headline: it composes row 5's per-tick re-read (`AI-350`) with the arm's once-per-full-tick clock (`AI-CLOCK-080`), and `AI-353` grades it Medium / Unknown; no runtime observation exists.

**Amended.** The name, the scorer veto and the absence of a walk order in the arm stand. Two consequences are narrowed. The members never taking a step: the executor row 5 issued at the previous evaluation keeps stepping toward its victim until the next evaluation, a stand or a refused route replaces the order (Medium, `AI-350`, `AI-353`). Turning on the spot as the targetless outcome: a human participant's member gets `ord+0x08` of 0, or 8 for a heal cast (`AI-349`), and the idle turn 0xb is written for an AI owner's member only. The setters replace every member's running order with 0 (`AI-351`, `AI-352`). `retracted.md` holds the entry.

### AI-REISSUE-077

Neither group arm carries any test that a target is still worth having: `AI-SCORE-069`'s scorer rewrites `ord+0x20` from scratch, and the arm calls `R0009` on whatever it finds there, which writes `ord+0x08 = 5`, `ord+0x0c = target` and `ord+0x14 = actor+0x12c` afresh (`AI-GUARD-007`). So the answer to *what ends a pursuit* differs by group order: under order **0** it is `AI-BREAK-041`'s post rule, and under orders 1/2/3/5 there is no break-off at all — only a re-score whose inputs may no longer contain the old target. Three consequences follow from routines already read rather than from new ones. **Death does not release a target**: `AI-GROUPSEE-068` phase (e) moves corpses aside, but restores every one of them when no living hostile is in sight, so the scorer is handed a corpse, `ord+0x20` takes it, and the engage is re-issued at it.

**Leaving sight does release one**, because the candidate list is rebuilt from a freshly cleared stamp each tick — except for an AI-owned member's remembered attacker cell, which is forced into the stamp for 20 group ticks (`AI-GROUPSEE-068` phase (b)). **And unreachability is handled one layer down**, by `AI-ROUTE-045`'s `mover+0x98`, which the order layer consumes on the actor tick without the group layer ever hearing of it

**Confidence.** High for the mechanism (it is a control-flow reading of two arms and one builder, all read end to end this round). **Medium** for "no other mechanism ends a group pursuit": the negative is over the two arms' own instructions plus `AI-SCORE-069`'s module-scoped enumeration of `ord+0x20`, and neither can see a wholesale copy of the order block

### AI-DEADROLL-078

An `EnumRefs` pattern sweep for byte stores into `+0x2d` returns **7 hits / 4 owners / 0 orphan** — the complete writer set of the group's working radius. Three are the guard setter `R0125` (`L00360`, `L00843`, `L00844`), two are `R0024`'s two latch edges (`L00370`, `L00368`, `AI-GUARD-021`), and the remaining two are `R0258` (`L00845`) and `R0259` (`L00846`) — **nine-instruction routines, byte-identical to each other**, each computing `grpAI+0x2d = grpAI+0x38 + 4 + (rand()*3/AImgr[0] - 1)`, i.e. `R0024`'s has-members edge extracted and duplicated. Both are unreachable on three independent instruments: `callto:` **0**, `refto:` **0**, `imm:` **0** for each — the triple `AI-DEAD-035` uses. This discharges the completeness half of `AI-GUARD-021`'s Medium for `grpAI+0x2d`: its writer set is now enumerated and every member classified

**Confidence.** **Medium**, and the reason is `AI-DEAD-036`'s: those three instruments are all blind to a `.text` dispatch-table slot, and this area's tables are all in `.text`. The fourth leg — a raw scan of the image for the little-endian dword through the PE section table — was **not** run, so "unreachable" is stated at the confidence three legs support and not the four that would make it High. High for the writer set itself and for the two routines' contents

### AI-GATE-079

`actor+0x4c & 4` is the fighter/mage axis and set is mage; it is derived from positive mana and can be set by spellbook construction, not copied verbatim from Data.bin. Thus the short-reach suppression is a human participant's mage not auto-engaging. Separately, 57 Humans rows on both roots author table `typeID=0x18`. A non-zero player-character constructor mode overwrites that value, but zero-mode map Humans retain it, so the overwrite is **not** a second reason the compare cannot match. The test's reachability depends only on the still-unlocated producer of `actor+0x50=0x17` (`AI-STATE-043`)

**Confidence.** High for the class bit and 57-row census; Medium for the `0x17` arm's unreachability, unchanged from `AI-STATE-043`

**Amended.** Corrected. The second reason for the `typeID==0x18` arm's unreachability and the class-bit provenance only — the 57-row census and missing state writer stand: The class bit is derived from positive mana (`HERO-HP-071`), and zero-mode map Humans can retain table `typeID=0x18` (refuted, EXP-0192). `retracted.md` holds the full entries.

### AI-CLOCK-080

The order layer steps a pursuit from the **actor tick**: `R0037` calls `R0016` four instructions before it dispatches on `actor+0x54` (`AI-PROGRESS-034`), and that is where `AI-PURSUE-040`'s arms 5 and 6 re-test position and re-path. The decision layer runs from the **AI slot**: `R0076` walks players and groups and calls `R0023` once per group (`AI-TICK-008`). So a group re-scores its targets once per full tick while the pursuit it ordered advances on every actor tick — every field the group layer writes is a decision the order layer then executes many times, which is why nothing in the group arms needs to be sticky (`AI-REISSUE-077`). One clause of `AI-TICK-008` is narrowed rather than confirmed: `callto:R0076` is 2 hits, and only the first, `L00279` in `R0147`, sits behind the signed `% 16 == 6` test (the mask of 15 at `L00847` and the compare against 6 at `L00848`).

The second, `L00163` in `R0075`, has **no modulo in front of it** — it is bracketed by two calls through the import slot at `L00849` (the clock) whose difference is stored, i.e. the pass is being timed. `R0147` has 4 callers over 4 owners, all in the pacer module `R0260`…`R0099`; `R0075` has **1** (`R0261`), so it is not dead

**Confidence.** High for the two clocks and for both dispatch paths (all cited instructions, and the four caller counts are printed enumerations with 0 orphan). **Unknown** whether `R0075` runs during ordinary play: its one caller was located, not read, and if it does, the AI decision rate is not `AI-TICK-008`'s once per full tick

### AI-LOS-081

`R0136`, four stack arguments, read whole (`evidence/rom-grouparm-excerpt.md` §9). For the cell named by its first two arguments it reads a **signed byte pair** out of one word of the grid at `fog+0x22000` — the low byte at `L00850`, the high at `L00851` — and uses them as a step back to a predecessor cell; loads that predecessor's running value out of the dword accumulator at `fog+0x24000`, stride 64 (`L00852`, the grid `TERR-FOG-084` identifies); subtracts a signed word from a third grid at `fog+0x28000` (`L00853`); subtracts the terrain height `world+0x9451c[cell]` reached through the world pointer cached at `fog+0x3a008` (`L00854`…`L00855`); adds the observer's altitude, the fourth argument (`L00856`); **stores the result into this cell's own accumulator slot** (`L00857`); and then branches on its sign — `<= 0` returns **1**, blocked, which is what stops `R0134`'s ring; `> 0` sets `fog[0x2a008 + cell] = 1` (`L00858`) and returns **0**.

This **amends `AI-SIGHT-006`**, which describes the cell as "tested against the actor's own altitude from the height plane": the altitude is one of four terms and the test is on an accumulated remainder, so a consumer implementing a per-cell altitude compare reproduces neither the shape of the visible region nor where it stops. Two smaller corrections ride along: the visible-cell byte is **set to 1**, while `AI-GROUPSEE-068`'s attacker-memory path **increments** the same byte, so the array is not a boolean; and the zero test at `L00859` is inert, its flags overwritten by the next instruction

**Confidence.** High (one routine read end to end; every grid base is an arithmetic identity from its own scale factor, and the two exits are distinguished by their own immediate return values). **Unknown** what fills the `fog+0x22000` step grid and the `fog+0x28000` cost grid: both are read here and neither was traced to a writer

## Diplomacy matrix writers and events

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-DIPLO-082 | The complete writer set of the matrix after construction is SIX routines, and two of them no `disp:a9c4` sweep can name. | High / Medium | ● active | [EXP-0114](../experiments/EXP-0114-diplomacy-change/) |
| AI-DIPLO-083 | The reached local registration routine writes row 0 and column 0 and uses a four-byte template for live-slot relations. The separate leave routine clears the live marker. | High | ● active (partially retracted) | [EXP-0114](../experiments/EXP-0114-diplomacy-change/), [EXP-0351](../experiments/EXP-0351-authored-player-control/) |
| AI-DIPLO-084 | The event flip has three callers and a fourth entry point, and one of them re-scans immediately — which is what `R0109` turns out to be. | High / Medium | ● active | [EXP-0114](../experiments/EXP-0114-diplomacy-change/) |
| AI-DIPLO-085 | A relation changed during a mission survives a save, because the matrix is serialised verbatim — and the byte count is exact. | High | ● active | [EXP-0114](../experiments/EXP-0114-diplomacy-change/) |
| AI-DIPLO-086 | A diplomacy change is not propagated — it is only observed, so it never interrupts anything already under way. | High / Medium | ● active | [EXP-0114](../experiments/EXP-0114-diplomacy-change/) |

### AI-DIPLO-082

`AI-DIPLO-004`'s enumeration is reproduced unchanged — `EnumRefs disp:a9c4` = **32 hits / 14 distinct owners / 0 orphan**, `imm:a9c4` = **0** — and then extended in the two directions it could not reach. **(a)** That row named its own blind spot, the `sub-object + 8` alias, and did not close it; the address of `session+0xa9bc` can only be taken by a `disp:a9bc` or an `imm:a9bc` instruction, so both sweeps are run — **3 hits / 3 owners** and **5 hits / 4 owners**, 0 orphan, **8 hits over 6 owners** in all — and every one is read. **(b)** A `disp:` sweep names the instruction that *forms* an address, not the one that *uses* it: twenty of the 32 hits are address computations rather than accesses, and following each to its use found `R0061` spilling the pointer at `L00860` and storing through it at `L00861`, a byte store with displacement **0**, invisible to any sweep for `0xa9c4`.

The six are: the `.alm` loader `R0128` and the two mission joins `R0131` / `R0132` (`AI-DIPLO-005`); the script action `R0262` (`TRIG-DIPLO-019`); session command `0x45` inside `R0061` (`SESS-DIPLO-024`); the join/leave pair `R0263` / `R0264` (`AI-DIPLO-083`); and the event flip `R0126` (`HERO-AGGRO-028`, `AI-DIPLO-084`) — **four by sweep, six by reading**. The nine **readers** and the bit each tests are listed in `evidence/rom-writers.md` §8; the widest is the loader's own mask of `0x7` at `L00253`, which is why the value space cannot be closed at `{0,1,2}` from the corpus

**Confidence.** High for each named writer (every store is a transcribed instruction) and for both enumerations, whose instrument and 0-orphan counts are stated. **Medium** for *six is all of them*: a wholesale `memcpy` / `REP MOVSD` of the session carries no displacement at all and no sweep here can see one — the session `Serialize` (`AI-DIPLO-085`) is exactly such a copy and was found by reading the serialiser, not by a sweep

### AI-DIPLO-083

`R0263(player)`, on `this = session+0xa9bc`, stores the low byte of the complete control dword `player+0x28` (`UNIT-OWNER-009` is partially retracted for its universal value/authorship interpretation) into **`matrix[idx][0]`** at `L00862`, overwrites it with 0 when the complete dword is exactly 2 (the compare against 2 at `L00863`, the overwrite at `L00864`), marks the slot live with **`matrix[0][idx] = 1`** at `L00865`, then walks the slots 1 to 49 skipping every slot whose `matrix[0][slot]` is zero and, for each live one, selects one of four bytes at `this+0..3` by whether `matrix[idx][0]` and `matrix[slot][0]` are both non-zero (`template[2]`), belong to different zero/nonzero classes (`template[0]` one way and `template[1]` the other) or are both zero (`template[3]`), writing `matrix[idx][slot]` at `L00866`/`925`/`940`/`952` and `matrix[slot][idx]` at `L00867`; finally the byte store of 2 at `L00868`, at offset `0x8` plus an index of `51*idx` built by shifts and adds,

**the forced diagonal again**. The constructor sets the template `{1, 1, 0, 0}` (`L00869` to `L00870`, the values 1 loaded at `L00778` and 0 zeroed at `L00779`, neither rewritten in between), so a joining player is **hostile in both directions to live players in the other zero/nonzero class, and neutral to the same class, before the forced diagonal**. The inverse is `R0264`, four instructions, whose byte store of 0 at `L00871` — `matrix[0][idx] = 0`, the slot goes dead. **This amends `AI-DIPLO-005`**, whose *"column 0 is never written"* is true of the `.alm` loader and false of the matrix, and whose forced diagonal turns out to be written on **three** independent paths rather than one: the loader (`L00255`), the mission joins (`L00872`, `L00873`, `L00874`) and here **Partially retracted by EXP-0351 (SESS-073): the universal owner enum/authorship and unequal-byte hostility conclusions are withdrawn. Input0x102 stores column byte2; opposite live byte5 both cells select template[2], as do other two-nonzero pairs. The local template, live-slot, diagonal and leave stores stand. Native registration/map-row/join ordering remains Unknown.**

**Confidence.** High (both routines read whole; the template's four bytes read out of the constructor with their register provenance; the diagonal's index recovered from the `LEA` arithmetic rather than assumed) High applies to those reached local operations, not native join ordering or a universal owner schema.

**Amended.** Owner interpretation and unequal-byte hostility only: The routine first stores the low byte, then replaces it with zero only for complete dword2 (partially retracted, EXP-0351). `retracted.md` holds the full entries.

### AI-DIPLO-084

`HERO-AGGRO-028` names `R0265` as the hook into `R0226`; `EnumRefs callto:R0226` returns **3 hits / 3 owners / 0 orphan** — `R0265` (`L00875`, the routine whose own string is `"Unknown magic damage type"`), `R0266` (`L00876`) and `R0267` (`L00877`) — all three guarding on both actors having an owner (`+0x14 != 0`). The fourth entry point does not go through `R0226` at all: **`R0109`** calls `R0126` directly at `L00878`, after the identical alarm preamble (`L00879` `[victim+0x158]+0x54 = 1`, `L00880` `+0x58 = the caster's cell`, `L00881` `+0x5a = 0`, matching `L00652`…`L00654` instruction for instruction), and then — alone among the six writers — calls the group candidate builder **in the same instruction stream**, the call at `L00882` to `R0110`.

Its chain is `R0268` → `R0269` → `R0109`, one caller at each step (`callto:` 1 hit / 1 owner / 0 orphan), and `R0268` subtracts a cost from `[target+0x9a]` while `R0269` also calls `R0003`, the routine that tests `matrix[caster][target] & 1` — the spell-cast family. So **casting at somebody makes the two players mutually hostile on the same rule as striking him**, and this closes what `DISCOVERY.md` recorded as *"the fifth caller of the candidate builder, lives outside the AI module, and is unread"*. The flip itself is unchanged and re-read here: the low-two-bit test at `L00251` before the bit-0 set at `L00883`, **separately gated for each direction** (`L00252` / `L00884`), so a pair can flip one way and not the other

**Confidence.** High for the caller enumeration and for both preambles (a printed enumeration with 0 orphan, and the two alarm sequences are instruction-for-instruction the same). **Medium** for *"the spell-cast family"* — that identification rests on the cost subtraction and on `R0003`'s diplomacy test, not on reading `R0268` end to end

### AI-DIPLO-085

The session `Serialize` `R0270` (`SAV-SESS-031`) carries the sub-object at `session+0xa9bc` in both directions with the same length immediate: the store path pushes `0x9cc` at `L00885`, offsets the object by `0xa9bc` at `L00886` and calls `R0271` at `L00887`; the load path pushes `0x9cc` at `L00888`, offsets by `0xa9bc` at `L00889` and calls `R0272` at `L00890`. **`0x9cc` = 2508 = 8 + 2500** — the four template bytes of `AI-DIPLO-083` plus their padding, and the 50x50 matrix of `AI-DIPLO-004` — so the whole object goes into the **world** half of the save and comes back byte for byte. Nothing on the load path re-derives it from the `.alm` type-5 row: the loader `R0128` runs on a **map** load, not on a save load. A consumer that rebuilds diplomacy from the map when restoring a save therefore reverts every script change, every join and every flip that had already happened

**Confidence.** High (the length is the routine's own immediate, appearing identically on the store and the load arm, and 2508 admits exactly one split into a documented 8-byte header and a documented 50x50 byte matrix — a rival in which only the template were saved would have pushed 8)

### AI-DIPLO-086

The script arm ends at the jump at `L00891` to `L00892`, its own epilogue: no notify, no re-scan, no order invalidation (`TRIG-DIPLO-019`). Session command `0x45` notifies the affected players and does not re-scan (`SESS-DIPLO-024`). The join and leave routines do neither. **Only `R0109` re-scans** (`AI-DIPLO-084`). Everything else downstream happens because a consumer re-reads the cell later, and the published group layer fixes when that is: the candidate list is rebuilt from a freshly cleared sight stamp and re-filtered through `AI-FILTER-001` on every evaluation (`AI-GROUPSEE-068`), `ord+0x20` is rewritten for every member every time (`AI-SCORE-069`), and that runs once per full tick (`AI-CLOCK-080`) — so for the **95.6 %** of shipped hostile placements sitting under a group order (`AI-CENSUS-047`) a change takes effect at the next AI tick.

But the order already issued keeps advancing on the **actor** tick, and under group orders 1/2/3/5 there is no break-off at all (`AI-REISSUE-077`). The consequence a consumer needs: **making a faction friendly mid-mission does not stop a unit that is already attacking — it stops that unit being re-selected**, and the attack ends when the next re-score no longer produces the target, not when the byte changes

**Confidence.** High that the five non-cast writers are bare (each arm read to its own epilogue, and the absence of a call is a property of transcribed instructions). **Medium** for the reach: it is a derivation over `AI-GROUPSEE-068` / `AI-SCORE-069` / `AI-REISSUE-077` / `AI-CLOCK-080` rather than a reading of a re-filter that exists, and no running session was observed — the write-up states the prediction and what would refute it

## Line-of-sight grids

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-LOS-087 | Both of `AI-LOS-081`'s untraced grids have ONE writer, it runs at world construction, and it depends on the map for nothing. | High | ● active | [EXP-0117](../experiments/EXP-0117-sight-region/) |
| AI-LOS-088 | What the two grids hold. Step grid, `fog+0x22000`, one `u16` per window cell, row stride `0x80`, addressed by `R0136` as `fog + 0x22000 + (a<<7) + 2b`… | High | ● active | [EXP-0117](../experiments/EXP-0117-sight-region/) |
| AI-LOS-089 | The region the predicate admits is a DISC of radius `scanRange`, not a square and not the 41x41 window. | High / Medium | ● active | [EXP-0117](../experiments/EXP-0117-sight-region/) |
| AI-LOS-090 | A cell that blocks is NOT pruned — the march walks past it — and the reason nothing is ever re-lit behind a blocker is a corpus fact about altitude, not a property of the code. | High / Medium | ● active | [EXP-0117](../experiments/EXP-0117-sight-region/) |
| AI-LOS-091 | The complete reader and writer sets of both grids — and the sweep that reads as 'nobody touches these' is the one an agent would run first. | High | ● active | [EXP-0117](../experiments/EXP-0117-sight-region/) |

### AI-LOS-087

`R0273` (`__thiscall`, no args) writes the step grid based at `fog+0x22a28`/`+0x22a29` and the cost grid at `fog+0x28a28`, walking `i = 0..20` outer (pointer stride `0x80`) x `j = 0..20` inner (stride 2) and storing **eight bytes per iteration** — the low and high byte of one `u16` in each of four quadrant mirrors. `callto:R0273` returns **1 hit**, `L00893` in `R0274`, the sight object's init; `callto:R0274` returns **4**, each preceded by the address computation of `world+0x58ee8` in one of the four world constructors (`R0275` `L00894`, `R0276` `L00895`, `R0277` `L00896`, `R0118` `L00897`), all of which run it **before** the map load `R0278`. So: not per tick, not per stamp, not on a terrain change — **once, before terrain exists**.

The init's own `REP STOSD` (`L00898`, `0x8000` dwords from `+0`) clears `+0x00000..+0x20000` and stops below all three tables. The only input is `k = [Scanning] ScanShift` of `World\Data\map.reg` (the constant 7 pushed at `L00899` is the compiled default; the key is present and **7** in both roots), which reaches the cost table by a floating-point multiply with the integer at `+0x2a000`, `1 << k`. **The builder ends with four literal stores** — `L00900`/`L00901` writing `0xff`,`0x00` at `fog+0x22aa8`/`+0x22aa9` and `L00902`/`L00903` writing `0x01`,`0x00` at `fog+0x229a8`/`+0x229a9` — which are the cells `(+1,0)` and `(-1,0)`. They are not redundant: on the `j == 0` axis the `+j` and `-j` quadrant addresses coincide, the last store wins, and `(i,j) = (1,0)` is the one axis cell falling in the diagonal zone, so both cells leave the loop pointing diagonally. Re-executing the builder with the listing's store order and **without** those four stores leaves **exactly 2** cells whose predecessor is not closer to the centre, and they are exactly those two

**Confidence.** High (the writer, its single caller, that caller's four call sites and the ordering against the map load are cited instructions; the four trailing stores are explained by a consequence of the store order that was derived before it was measured, and no other reading gives them a function)

### AI-LOS-088

What the two grids hold. **Step grid**, `fog+0x22000`, one `u16` per window cell, row stride `0x80`, addressed by `R0136` as `fog + 0x22000 + (a<<7) + 2b` — its **low byte is a signed column delta and its high byte a signed row delta**, values in `{-1,0,+1}`, naming the cell one Bresenham step **toward the centre**. The builder picks the pair by the ray's slope: `j < i>>1` (`L00904`) gives a pure column step, `j > 2i` (`L00905`) a pure row step, and the band between them the diagonal. Every one of the 1 680 non-centre cells therefore has a predecessor **strictly closer** to the centre in Chebyshev distance, so every chain terminates at the seeded cell — measured, not assumed. **Cost grid**, `fog+0x28000`, one **`i16`** per cell, same geometry, value `ftol((1<<k) * sqrt(i*i + j*j) / max(i,j))` (the floating-point square root, multiply by the integer at `+0x2a000` and divide by the larger axis at `L00906`, then the call to `R0279`, which is `__ftol` with the rounding mode set to truncation toward zero).

That is the ray's Euclidean length divided by its Bresenham step count: **the mean length of one step along that ray**, in units of `1/(1<<k)` cell. For the shipped `k = 7` the table runs **128..181** — exactly `1<<k` on both axes and `181` on the diagonals — so the accumulator seeded with `(1<<(k-1)) + (scanRange<<k)` is a **budget in 1/128 cell, equal to `scanRange` cells plus a half-cell rounding term**. The centre cell's cost is computed by a division by `max(0,0)` and is never read. **Neither grid is a copy or projection of any plane already decoded** — not `TERR-PASS-073`, not `TERR-COST-052`, not a fog plane; a consumer has nothing to reuse and must build both

**Confidence.** High (both addressing forms, the FPU sequence with its four operands and the zone branches are cited instructions; the closed form's own bounds `128..181` are reproduced by re-execution, and the chain property is a measured invariant over the whole window)

### AI-LOS-089

On flat ground with `k = 7` the march reaches exactly `scanRange` cells along the axes (each axis cell costs `1<<k`, so the seed `128*sr + 64` survives `sr` steps and no more) and `floor((128*sr+64)/181)` along the diagonals — measured visible-cell counts **9 / 21 / 45 / 69 / 105 / 145** for `scanRange` 1..6, and 1 253 at 19, where the ring bound `r <= 19` (the signed compare against `0x14` at `L00907`) begins to bind. The window is 41x41 but only rings 1..19 are ever walked, so its outermost ring is unreachable in principle. Terrain moves this a long way in both directions, because the per-cell term is `h(observer) - h(cell)` and descending ground **returns** budget: over the shipped corpus at `scanRange = 6`, `k = 7`, stride 3, **81 048 EN and 61 364 RU observers on 38 / 34 maps**, the visible count runs **29 minimum, 151.5 mean, 1 248 maximum** — a **43x spread from where the unit is standing**, with the mean *above* the 145-cell flat baseline.

At `scanRange = 12` the same sweep gives 84 / 517.9 / 1 521, the maximum being the whole reachable window. Unrolled, the recurrence says why the spread is that large and that asymmetric: `acc(n) = seed - sum(cost) - sum(h(cell)) + n*h(observer)`, so **the observer's altitude is added at every step of a ray while a cell's own height is charged once**, when the ray passes through it. Reach therefore scales with where the unit stands and barely with what is in the way — a single tall cell costs its height once and casts no shadow behind itself. **The consequence for a consumer, in one sentence: sight range is a budget, not a radius, and the term that dominates it is the observer's own altitude**

**Confidence.** High (the ring bound, the seed and the cost landmarks are cited instructions and closed forms) / **Medium** for every corpus figure: they are this round's transcription of four routines re-executed on shipped altitude grids, with no units, no session and no running original, and they carry their parameters — `k = 7`, stride 3, the stated `scanRange`, the 72-map corpus

### AI-LOS-090

`R0136` stores the failing value into the cell's own accumulator slot at `L00857`, **before** the `JLE` at `L00908`, and returns 1; `R0134` consumes that return only at `L00909`/`L00910`/`L00911`/`L00912`, to clear the per-ring all-blocked flag `[ESP+0x30]`, and the ring runs to its end regardless. A later cell whose predecessor is the blocked one reads the stored negative value and continues from it. Whether the accumulator can rise is arithmetic: `acc[cur] = acc[pred] - cost - h(cell) + h(observer)`, the smallest cost in the table is `1 << k = 128`, and the height plane is read **signed** (`L00855 MOVSX`), so a rise needs `h(observer) - h(cell) >= 128`. Corpus: the `.alm` type2 altitude bytes run **0..127 across all 72 shipped maps, with zero bytes >= 0x80**, so the largest rise available anywhere is **127 — one short of 128**.

The accumulator is therefore strictly decreasing along every ray on shipped data, the visible set is star-shaped, and a re-execution over 81 048 + 61 364 observers finds **0 cells** visible whose predecessor was evaluated and blocked, at every `scanRange` from 4 to 19. **This is a customisation limit and it is one unit wide:** a map authoring a single altitude byte >= 0x80 makes that term exceed the axis cost and the region stops being star-shaped, so an implementation that prunes at the first blocker is indistinguishable from the original on shipped maps and diverges on an authored one

**Confidence.** High (the store-before-branch, the flag's only consumers and the `MOVSX` are cited instructions, and the 128-vs-127 margin is arithmetic on them) / **Medium** that no re-lighting occurs, which rests on a corpus census of 72 maps and on nothing in the image enforcing the byte range

### AI-LOS-091

Step grid: written only by `R0273` (`disp:22a28` 3 hits, `disp:22a29` 3 hits, plus the four trailing stores at `0x22aa8`/`0x22aa9`/`0x229a8`/`0x229a9`, all one owner), read only by `R0136` (`disp:22001` **1 hit** image-wide for the high byte, `disp:440` for the low). Cost grid: written only by `R0273` (`disp:28a28`, 1 hit), read only by `R0136` (`disp:500`). **`disp:22000` and `disp:28000` each return 0 hits, and so do `imm:22000` and `imm:28000`** — the base is folded into a `SHL`, so it appears as the address computations `+0x440` (`L00913`) and `+0x500` (`L00914`), and `imm:440`/`imm:500` cannot see those either, because `EnumRefs imm:` **excludes displacements by construction** and an address computation's operand is a displacement.

Only `disp:440`/`disp:500` see them. Every power-of-two folding of both bases was swept — `11000 8800 4400 2200 1100 880 220` and `14000 a000 5000 2800 1400 a00 280`, in both modes — and all are 0 on this object. Blind spots, stated: a wholesale `REP MOVSD` of the containing structure carries no displacement, bounded here by `range:L13127:L13128` reporting **0 orphan bytes** over 24 functions with the class's only bulk moves (`L00915`, `L00916`, `L00917`) all read; and `callto:` cannot see a pointer table, closed by a raw whole-image scan (`tools/sightregion -mode ptrscan`, which walks the PE section table and is committed with the figure) finding **0 occurrences** of `R0274`, `R0273`, `R1847`, `R0136`, `R0134`, `R0145` or `R0135` as a dword anywhere in the 1 977 344-byte image. The visibility byte map `fog+0x2a008` has 9 accesses over 5 owners, **all in the simulation** — no renderer reads it, so the AI's sight and the drawn fog share no storage

**Confidence.** High (each sweep is quoted with its mode, its hit count and its owners; both instruments' blind spots are named, and each is closed by a second instrument of a different kind)

## Sight radius writers and the fog

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-SIGHT-092 | `actor+0xa4` is a `u16` sight radius in 1/256 cell whose high byte is `actor+0xa5`, and it has SIX writers — a `disp:a5` sweep sees two of them. | High / Medium | ● active | [EXP-0120](../experiments/EXP-0120-sight-writers/) |
| AI-SIGHT-093 | The drawn fog and the AI's vision are ONE algorithm written twice, and at a whole-cell sight they agree exactly — so `TERR-FOG-088`'s Medium is discharged. | High | ● active | [EXP-0120](../experiments/EXP-0120-sight-writers/) |
| AI-SIGHT-094 | Where the two DO diverge: the server truncates a hero's sight to whole cells and the drawn fog does not, so the fog can show a hero cells his own simulation does not grant him. | High / Medium | ● active | [EXP-0120](../experiments/EXP-0120-sight-writers/) |

### AI-SIGHT-092

The pairing is the constructor's own two stores: the byte stores of 5 into `+0xa5` (`L00260`) and of 0 into `+0xa4` (`L00918`) are the little-endian halves of `0x0500` = 5.0 cells, and the 16-bit store at `L00919` writes the pair as one word. The complete set **over the two forms an offset sweep can enumerate**, each with the sweep that finds it: (1) `L00260`, the constructor default, `disp:a5`; (2) `L00261`, a byte store of 5 into `+0xa5` in `R0146`, behind the compare of the word at `+0xe` with `0x18` at `L00920`, which leaves the low byte alone, `disp:a5`; (3) `L00919`, the hero recompute `R0280` (`HERO-SIGHT-007`'s `ftol(((mind + reaction)/25 + 4) × 256)`), a **word** store spanning the byte, `disp:a4` only; (4) `L00921` forms the address of `+0xa5` and `L00922` calls `R0281`: the `Data.bin` **Units** streamer's `scanRange` slot (`TERR-FOG-081`), `imm:a5` only; (5) `L00923` forms the address of `+0xa5` and `L00924` calls `R0281`: the **Humans** streamer's slot 8 (`DAT-HUMANS-008`), `imm:a5` only; (6) `L00925` forms the address of `+0xa4` and `L00926` calls `R0282`: the actor's stream **load** arm, `imm:a4` only. Both helpers store through a pointer — a byte store at `L00927`, a 16-bit store at `L00928` — so neither carries a displacement. `disp:a3` and `disp:a2` are 4 hits each, all byte-wide on their own field, so no wider store spans `+0xa5` from below

**Confidence.** High that each of the six is a writer (every one is a cited instruction re-read from raw bytes on both roots, and each helper's own store was read) / **Medium** that six is the whole set: the enumeration is complete over displacement form and immediate-computed-address form and over **nothing else** — an address reached in two steps, an address carried in a field, and a wholesale `memcpy` of the actor are invisible to every sweep run here. Refuting `two` needs one counterexample and has four; asserting `six` needs an instrument this round does not supply
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### AI-SIGHT-093

Nineteen clauses of `R0283`/`R0284`/`R0285` against `R0273`/`R0134`/`R0136` (`evidence/rom-two-implementations.md`): both slope tests (`j < i>>1` at `L00929`/`L00930`, `j > 2i` at `L00931`/`L00932`), all three zone deltas, the four-mirror store order, the cost form `ftol(S·√(i²+j²)/max(i,j))` with `S = 1<<k` (`L00933`…`L00934` / `L00906`…`L00935`), the `i,j = 0..20` bounds, the recurrence `pred − cost − h(cell) + h(obs)` stored **before** the test (`L00936`/`L00937`, `L00857`/`L00908`), *visible iff > 0*, the ring bound `r = 1..19` (`L00938`/`L00907`), four edges of `2r+1`, and — the clause `TERR-FOG-080` missed — the **same four trailing literal stores**, `L00939`…`L00940` against `L00900`…`L00903`, repairing the same two cells `(+1,0)` and `(-1,0)` from the same wrong values.

And the seeds are algebraically equal for a whole-cell sight: `(sr·256 >> (8−k)) + (1<<(k−1))` = `(sr<<k) + (1<<(k−1))` at `k = 7`. Executed by `tools/sightwriters`: the four tables differ in **0 of 1681** window cells, patched or unpatched; the two marches over the shipped altitude grids at `scanRange = 6`, stride 4, 72 maps and both roots give **43 640 observers more than 27 cells from any edge and 0 disagreements** (27 = window radius 20 + the wider inset). **A consumer needs one implementation, parameterised**

**Confidence.** High (every clause is a cited instruction on both sides, and the agreement is executed rather than described; the two live rivals — different tables, different seeds — are each refuted by their own measurement)

### AI-SIGHT-094

The client seeds from the whole `u16` — the signed 16-bit read at `L00941` of `+0x102`, the drawable field `TERR-FOG-081` traces to `actor+0xa4` — then an arithmetic right shift by `8 − k` at `L00942`, keeping 1/128-cell precision. The server seeds from the byte: the byte read of `+0xa5` at `L00263` and a left shift by it at `L00943`, discarding all eight low bits. For a monster the streamers write whole cells and the two coincide. For a **hero** they do not: `HERO-SIGHT-007`'s derive is `ftol(((mind + reaction)/25 + 4) × 256)`, a multiple of 256 only by accident. On flat ground at `k = 7`, `actor+0xa4 = 1535` (5.996 cells) reveals **145** cells to the fog and admits **105** to the stamp — a **40-cell** gap; 1408 gives 117 against 105, 1152 gives 77 against 69. The second divergence is `TERR-FOG-118`'s in-bounds inset. Nothing else differs

**Confidence.** High for the arithmetic (four cited instructions, two per side, and the region figures are this round's transcription of both) / **Medium** for the reach: whether a player-controlled hero is ever the subject of a server sight stamp is not established here — `R0134`'s two callers are taken from EXP-0117's call graph and were not re-read — and no runtime session witnessed the gap

## Guard post, stance setters and the radius roll

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-POST-095 | A load-time guard's post is written by the guard SETTER, not by the per-actor initialiser — which that member never reaches. | High | ● active | [EXP-0125](../experiments/EXP-0125-guard-post/) |
| AI-POST-096 | `ord+0x00`'s writer set is at least 26 stores over 23 routines, not four. | High / Medium | ● active | [EXP-0125](../experiments/EXP-0125-guard-post/) |
| AI-POST-097 | A guard's post is not fixed for the mission — six reachable sites re-anchor it, and `AI-BREAK-041`'s 5-cell block moves with it. | High / Medium | ● active | [EXP-0125](../experiments/EXP-0125-guard-post/) |
| AI-STANCE-098 | All four stance setters anchor a post, and only the guard forms have a mid-step branch. | High / Medium | ● active | [EXP-0125](../experiments/EXP-0125-guard-post/) |
| AI-LOAD-099 | The map load spawns the placements and then sets the stances, seventeen bytes apart in one basic block — and the order block starts zeroed. | High | ● active | [EXP-0125](../experiments/EXP-0125-guard-post/) |
| AI-GATE-100 | The stance walk's session gate wraps BOTH setter branches — closing `AI-AUTHOR-015`'s open item — and the inversion is benign for a reason that row did not have. | High / Medium | ● active | [EXP-0125](../experiments/EXP-0125-guard-post/) |
| AI-CENTRE-101 | A placement is centred, and that is what decides which branch the guard setter takes. | High / Medium | ● active | [EXP-0125](../experiments/EXP-0125-guard-post/) |
| AI-RANGE-102 | `AImgr+0x00` is `0x8000`, and `n × rand() / AImgr[0]` is the AI module's uniform-range idiom. | High / Medium | ● active | [EXP-0125](../experiments/EXP-0125-guard-post/) |
| AI-JITTER-103 | The guard radius jitter is exactly `{−1, 0, +1}`, in thirds — `AI-GUARD-021`'s ±1 confirmed rather than assumed — and the empty-members edge omits the `+4`. | High / Medium | ● active | [EXP-0125](../experiments/EXP-0125-guard-post/) |
| AI-TURN-104 | The idle turn's quotient runs 0…189, not 0 — so the arm re-picks a near-uniform facing over a byte, not one eighth of a circle. | High / Medium | ● active | [EXP-0125](../experiments/EXP-0125-guard-post/) |

### AI-POST-095

`R0125`, the routine `AI-AUTHOR-015` reads for `grpAI+0x20 = 1`, writes every member's `ord+0x00` first: `L00377` stores `0xb` into `+0x50`, then `L00944` calls `R0040` (which tests both sub-cell offsets against `0x80`: on a cell centre?), then either `L00829` stores the current packed cell from `R0019` (which returns the word at `+0x2`) or `L00830` stores it from `mv+0x06`; then `L00945` stores 0 into `ord+0x08`, no order. That is **33 bytes after** the state store and **0x76 bytes before** the store of 1 into `+0x20` at `L00314`, and the setter's member walk `L00946`…`L00947` has the same instructions as the arm's own `L00948`…`L00826`, so the population written is the population read.

`AI-GRPGUARD-074`'s "for a group under order 1 from load nothing has ever written `ord+0x00`" is **refuted**, and `AI-POST-042`'s mechanism is **corrected**: `R0137` is reached only through `R0008` arm `0xb`, `R0159`, `R0146` and `R0175` (`callto:R0137` = 6 / 4 / 0), all under the per-actor machine, which `AI-ORDER-010` establishes a group under order 1 never evaluates (`callto:R0008` = 2 / 2 / 0, the live one being `L00284`, inside arm **0**'s body at `L00304`). The **value** both rows give — the spawn cell — survives; the route to it does not, and three behaviours turn on the difference (`AI-STANCE-098`, `AI-POST-097`, `AI-GATE-100`)

**Confidence.** High (three routines read end to end, all 69 cited encodings re-read out of the PE section table by `tools/guardpost -mode cites` and 0 disagree, and every branch and call target fixed by its own rel8/rel32 rather than by a second decode)

### AI-POST-096

The field sits at displacement **0** of a block reached through `[actor+0x158]`, so a store into it carries no displacement of its own and `EnumRefs disp:` — the instrument behind `AI-POST-042`'s enumeration — cannot see it at all. The sweep that can is an instruction-pattern sweep for 16-bit stores through a bare pointer: **154 hits / 118 owners / 5 in orphan code** image-wide, reduced to the hits whose *preceding* instruction loads the base from `[<reg> + 0x158]`, which leaves **26 stores / 23 routines** (`evidence/enum-post-writers.md`). `AI-POST-042` names five of them; the eighteen it does not include `R0125`, `R0060`, `R0063` and `R0207` — that is, every stance setter. So the clause "Its complete set of later writers is four" is **retracted**

**Confidence.** High for the refutation — one counterexample is enough against a completeness clause and `L00829` is one, cited and re-read. **Medium** for the figure 26/23, which is a **lower bound**: the filter keeps only stores whose base was loaded in the immediately preceding instruction, and a wholesale copy of the block carries neither a displacement nor a load

### AI-POST-097

`EnumRefs callto:R0125` = **6 hits / 6 owners / 0 orphan**: the map-load stance walk `L00384`, the mission-description builder `L00388`, the **player's** group-command dispatcher `L00378`, the **script's** `L00596`, plus `L00949` and `L00950`. Every one runs the per-member body of `AI-POST-095`, so every one rewrites `ord+0x00` to wherever the unit is standing at that moment. `R0207` (`callto:` = 1 / 1 / 0, from `L00951`) is the same body for a single actor, with the same two-way branch at `L00952`…`L00953`. `AI-POST-042`'s "for a guard it never moves" and "fixed there for the whole mission" are therefore **retracted**: they hold for a creature nobody commands, which is most of a shipped map, and fail the moment Guard is issued again

**Confidence.** High (the caller enumeration is a `callto:` with 0 orphan, and the body it reaches is one routine read end to end). **Medium** that the enumeration is complete — `callto:` sees direct calls and `.rdata` code pointers, so a dispatch through a computed pointer would be missed

### AI-STANCE-098

`R0060` (`L00141`) and `R0063` (`L00151`) — the two Stand Ground setters `AI-STAND-076` names — write `ord+0x00` from `R0019` **unconditionally**, with no idle test at all, two instructions after their own `actor+0x50 = 0xc`. The guard forms `R0125` and `R0207` branch on `R0040` and, for a unit **not** on a cell centre, take `mv+0x06` instead (`L00954`, `L00955`) — the cell it is stepping **into**, not the one it is leaving. `R0054` is what puts a cell there: `L00956` reads the word at `+0x8` of the live path node `[actor+0x17c]` and `L00957` stores it at `mv+0x06`, with `mv+0x80` taking the same value once the cell is reserved. So guarding a walking unit anchors it one cell ahead of itself

**Confidence.** High for the four bodies and the branch (each read end to end, every store a cited and re-read instruction). **Medium** for `mv+0x06` being the step target: `R0054` was read whole, but that field's writer set was not enumerated

### AI-LOAD-099

`R0128`: the call to `R0151` at `L00958` (the `.alm` spawner, `callto:R0151` = **1 / 1 / 0**), the constant 1 pushed at `L00379`, and the call to `R0067` at `L00959` (the stance walk). No branch separates them. So "the actor's current cell" at the setter is the placement's own cell, which is what makes `AI-POST-095`'s post the spawn cell rather than a stale value. The block it writes into is allocated at `L00960` with size `0x94` and constructed by `R0286`, whose first act (`L00961` to `L00962`) is to clear `0x25` dwords is exactly `0x94` bytes, so `ord+0x00` is 0 until something writes it and `R0137`'s zero test of the word at `+0` tests a real zero. The sibling allocation two lines up is the allocation of `0xb4` bytes at `L00963` and the call to `R0206` at `L00964` for `[actor+0x154]`, likewise zeroed whole before `+0x5 = 0x41`, `+0x8 = 5`, `+0x9 = 0xff`, `+0xa = 0x10`

**Confidence.** High (both allocations, both constructors and the two calls are cited instructions re-read out of the PE, and the `CALL rel32` targets are computed from the bytes rather than taken off a listing)

### AI-GATE-100

In `R0067`: the compare of a dword at `+0xc` with zero at `L00381` and its branch on zero at `L00391`, and the same compare at `L00965` with its branch on non-zero at `L00382`, both short displacements to **`L00392`**, which is the group-list advance — and the calls at `L00154` (`R0063`) and `L00384` (`R0125`) are the only instructions between them and it. So when either test fails **no group is set at all** and every group keeps `R0150`'s `grpAI+0x20 = 0`. That is `AI-ORDER-010` arm 0, which runs each member's own `actor+0x50` machine, whose constructor default is `0xb` (`AI-STATE-011`, `L00324`), whose body is `R0137` — which writes the post at the first tick. **The post is written on both configurations, by different routines**, and `AI-POST-042`'s mechanism is the live one only here

**Confidence.** High (both gate tests and both `rel8` targets are re-read out of the image, and the two calls' position between the gate and its target is a property of the byte layout, not of a decode). **Medium** on `session+0x0c` meaning multiplayer — that is `SESS-MAP-010`'s reading, quoted rather than re-derived

### AI-CENTRE-101

The position record `[actor+0x10]` holds the packed cell at `+0x02` and two sub-cell offsets at `+0x04`/`+0x05`; `R0040` returns 1 exactly when both are `0x80`. It has **four** placement entry points and every one writes `0x80` into both: the reset `R0287` (`L00966`, cell 0), `R0288` (`L00967`), `R0289` (`L00968`) and `R0290` (`L00969`), the last of which additionally **clamps both coordinates into `[8, dim−9]`** — the lower bound 8 stored at `L00970`, the upper bound formed by subtracting 9 at `L00971`, and the bound 8 loaded at `L00972` — reading the map's dimensions from `[world+0x50000]` and `[world+0x50004]`. The only other writer of those offsets, `R0039` (`L00973`/`L00974`), re-centres on **arrival** and is reached from three movement routines (`callto:R0039` = 3 / 3 / 0). A unit that has been placed and has not moved is therefore centred, and the setter takes `L00975`

**Confidence.** High (all four entry points read end to end, every `0x80` store a cited and re-read instruction, and the clamp's bounds are the routine's own immediates). **Medium** that the four are the complete placement set — the sweep behind that is an instruction-pattern sweep for 16-bit stores at `+0x2` over the image, which cannot see a placement made by copying the record

### AI-RANGE-102

The AI manager is the global at `[L00004]`, `new`'d at `L00976`, constructed by `R0077` (`L00977`) and stored at `L00978`; `R0077` forwards the same `this` to `R0124` (`L00979`), which seeds the CRT generator (a call with argument 0 to `L00980`, then `R0291`) and then writes the first dword of the manager as `0x8000` at `L00371`. An instruction-pattern sweep for stores of `0x8000` through a bare pointer returns **1 hit / 1 owner / 0 orphan**, that store. `0x8000` is `RAND_MAX + 1` (`AI-RAND-058` fixes `RAND_MAX` at `0x7fff`), so every signed division by the first dword of the manager after a call to `rand` (`R0179`) yields `floor(n · rand() / 32768)`, uniform on `0…n−1`. A sweep for signed divisions by a dword through a pointer returns 86 hits / 50 owners, of which **19 divide by a bare pointer and 15 of the 19 are inside `L00339..L00340`**. This **closes `AI-RAND-058`'s Medium**: `R0124` writes no vtable at `this` because it writes a number there

**Confidence.** High for the constructor's store, the allocation chain and the counts (each a cited instruction re-read out of the PE, and the global's writer set is a printed enumeration). **Medium** that the field still holds `0x8000` at any given roll — the sweep is one store form and cannot see a register-form write, and the other 13 bare-pointer `IDIV` sites were counted rather than read

### AI-JITTER-103

`R0024`'s has-members edge calls `rand` (`L00369`), multiplies the result by 3 (`L00981`), divides by the first dword of the AI manager (`L00982`; the manager is reloaded at `L00983` and passed to `R0009` at `L00984`), subtracts 1 (`L00985`), adds that to the byte at `grpAI+0x38` (`L00986`, `L00987`), adds 4 (`L00366`) and stores the sum into the byte at `grpAI+0x2d` (`L00370`). With `AImgr[0] = 0x8000` (`AI-RANGE-102`) the quotient is `floor(3·rand()/32768)` ∈ `{0,1,2}` — 0 for `rand() < 10923`, 1 for `10923…21845`, 2 for `21846…32767` — so `grpAI+0x2d = grpAI+0x38 + 4 + {−1,0,+1}`, each about a third of the time. The **empty**-members edge at `L00988`…`L00368` is the identical sequence **without** the added 4, giving `grpAI+0x38 − 1 … +1`. Both dead copies `AI-DEADROLL-078` found, `R0258` and `R0259`, carry the has-members form

**Confidence.** High (one routine read end to end, every instruction of both edges re-read out of the PE, and the divisor's object fixed by the register's own reload and its later use as a `this`). **Medium** for the thirds, which inherit `AI-RANGE-102`'s Medium on the divisor still holding `0x8000` at the roll

### AI-TURN-104

`R0205`'s dividend is built from `rand()` in four steps at `L00989` to `L00990` (multiply by 3, shift left by 5, subtract the original, shift left by 1), `((3x · 32) − x) · 2` = **190 × rand()** — and `L00991` divides by the first dword of the routine's own `this`, which is the AI manager (the read of `+0xa50` at `L00992` is the same world field `R0137` reads at `L00993`). With `AImgr[0] = 0x8000` the quotient is `floor(190·rand()/32768)`, **uniform on 0…189**. The stored facing is then built from the current facing byte (`L00994`), `0x21` added (`L00995`) and the quotient added (`L00996`), stored as a byte at `+0x1` of the mover (`L00997`) — the current facing plus `0x21` plus that quotient, truncated to a byte. `AI-RAND-058` left "an identically-zero quotient is not established"; it is now positively **refuted**, and `AI-ORDER-039`'s surviving "one eighth of a circle" wording describes only the `0x21` term. The arm's entry gate is unchanged: `L00998`…`L00999`, `ord+0x54` non-zero or `rand() < 0xcd`

**Confidence.** High (the routine read end to end, the multiplier recovered from the shift-and-add chain rather than from an immediate, and the divisor's object fixed by `[EDI+0xa50]` matching a known AI-manager field). **Medium** for the range, inheriting `AI-RANGE-102`'s Medium on the divisor

## Candidate collections and the order 5 gate

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-CANDCOUNT-105 | `AImanager+0xbb4` is the element count of the candidate collection at `+0xba8` | High | ● active | **[EXP-0126](../experiments/EXP-0126-swarm2-gate/)** |
| AI-CANDLIST-106 | The AI manager carries two candidate collections, at `+0xba8` and `+0xbc8`, each four bytes inside an outer object at `+0xba4`/`+0xbc4` — and the pair is shared scratch, used by per-*actor* acquisition as well as by the group layer. | High / Medium | ● active | **[EXP-0126](../experiments/EXP-0126-swarm2-gate/)** |
| AI-SWARM2GATE-107 | Group order 5 runs order 4's arm whenever the group sees nothing, and that fallback is the ordinary path rather than an exotic one. | High | ● active | **[EXP-0126](../experiments/EXP-0126-swarm2-gate/)** |
| AI-CMDSET45-108 | The order-4 and order-5 command setters are one routine emitted twice, and they differ in exactly one byte in the whole image. | High | ● active | **[EXP-0126](../experiments/EXP-0126-swarm2-gate/)** |
| AI-BB4SWEEP-109 | No `disp:bb4` sweep can see a writer of `AImanager+0xbb4` — the complete write population is three instructions and every one carries displacement `0xc`. | High / Unknown | ● active | **[EXP-0126](../experiments/EXP-0126-swarm2-gate/)** |
| AI-CANDBYTE-110 | The order-5 gate tests the candidate count as a dword; the scorer tests only its low byte — so a count that is a non-zero multiple of `0x100` passes the gate and scores nothing. | Medium | ● active | **[EXP-0126](../experiments/EXP-0126-swarm2-gate/)** |

### AI-CANDCOUNT-105

— so what group order 5 branches on is *"did the build that just ran find anything"*, not an assigned flag. The collection's own `RemoveAll` `R0251` fixes the layout in its body: it walks a chain from `[this+0x4]` through `[node+0x0]` destroying each `[node+0x8]`, then sets the fields `+0xc`, `+0x10`, `+0x8` and `+0x4` to 0 (`L01000` to `L01001`), and frees the plex chain at `[this+0x14]` — count first. It is called on the object at `AImgr + 0xba8` (`L01002`) and on the object at `AImgr + 0xbc8` (`L01003`), so `+0xc` on the first collection *is* `+0xbb4`. The cross-read a coincidence would have to survive sits inside one routine: `R0110` reads the **second** collection's count at `AImgr+0xbd4` (`L01004`) and again at `+0xc` of a pointer set to `AImgr+0xbc8` at `L01003` (`L01005`) and not reassigned between — the same field through two bases, fourteen instructions apart, both used as *"is this collection empty"*.

The owner is the AI manager, the global at `[L00004]` (`AI-RANGE-102`): its constructor `R0077` keeps `this` in ESI (`L01006`) and calls `R0124` with ECX untouched (`L00979`), and that callee holds the same `this` from `L01007` to its epilogue while constructing both outer objects (`L01008` forms the address of `+0xba4` and `L01009` calls `R0120`; `L01010` does the same for `+0xbc4`). Every address and branch target above is re-read out of the PE by `tools/candcount` — 41 anchors, 9 recomputed branch targets, 0 failures

**Confidence.** High. `RemoveAll`'s own body defines `+0xc`; the same field is read through two different bases inside one routine; the owner is a cited constructor chain rather than an inference from displacement; and every citation is checked against the file's bytes rather than against a second decode (`docs/INSTRUMENT.md` rule 6). The reading this refutes — an assigned mode flag — is excluded by the write population, not merely unsupported (`AI-BB4SWEEP-109`)

### AI-CANDLIST-106

`R0032` is the outer object's add and is the whole proof of the `+4`: `L01011` adds 4 to the object pointer and `L01012` calls `R0292` (AddTail). The compiler emits the null-preserving form of the same adjustment where the outer pointer could be null — a null-to-zero mask applied to `AImgr+0xba8` when the outer pointer is `AImgr+0xba4` (`L01013` to `L01014`) — and its simple form (a zero test at `L01015` and the address of `+0xba8` at `L01016`). Stride `0x20` = 4 + the collection's `0x1c`, and the first outer object begins exactly where the preceding member ends: `AI-PREF-070`'s 4×4 byte matrix occupies `+0xb94`…`+0xba3`, sixteen bytes, written one at a time in this same constructor at `L00571`…`L00777`.

That adjacency also settles the sibling class for good — on it a *collection* spans `+0xb8c`…`+0xba7` (`L01017`…`L01018`), which the matrix's sixteen individual byte stores exclude here. `EnumRefs disp:` at `ba4 ba8 bb4 bc4 bc8 bd4` returns 23 / 15 / 15 / 20 / 4 / 4 hits, **0 orphan** in all six; **ten** AI-module routines reach the pair through their own `this`, each prologue cited — `R0114` (which restores `this` into EBX at `L01019`, before its first hit), `R0022`, `R0004`, `R0110`, `R0177`, `R0153`, `R0024`, `R0026`, `R0109`, `R0293` — plus six that touch only the outer objects: `R0294`, `R0124`, `R0295`, `R0296`, `R0297`, `R0298`. `R0022` is the per-actor acquisition rule of `AI-ACQUIRE-002`, so the pair is **not** per-group state and could not carry a per-group mode even if something wrote one

**Confidence.** High for the layout — the `+4` adjustment is a cited instruction in two forms and the stride is arithmetic on the sweeps' own hits. **Medium** for the owner enumeration: six `disp:` sweeps over the whole image, blind by construction to a wholesale copy of the containing structure and to any access that reaches a collection through a pointer saved earlier rather than through `AImgr + <disp>`

### AI-SWARM2GATE-107

`R0026` is `__thiscall(AImgr, group, cell)` with two stack arguments: `L01020` keeps the manager, `L01021` passes the group, `L00459` calls `R0110` to build the candidate list, and the very next instructions read the count that build produced (the dword at `+0xbb4`, `L00460`), test it for zero (`L01022`) and branch on zero with a 32-bit displacement of `+0x96` (`L00461`), which pins the target at `L00462` however the bytes are decoded. There `L00462` reloads the cell argument and `L00463` calls `R0154` to forward **both** arguments to order 4's arm (`AI-MOVE-023`), which never reads the second one. The value cannot be stale: `R0110` empties the collection at `L01023`, **before** its own empty-group early exit at `L01024`, so a group with no members reaches the gate at count 0 and falls through.

And the count is whatever that build produced, which by `AI-GROUPSEE-068` includes corpses — collection B is moved back into A when A ends empty — so a group whose only visible candidate is dead still runs order 5's own body. **The two zero-target cases are therefore different, and the difference is a walk:** *no candidate at all* takes the gate to order 4's arm and the members walk; *candidates the scorer vetoes* — `R0177` writing `ord+0x20 = 0` for every member (`AI-SCORE-069`, `AI-PREF-070`) — leaves the arm in its own body, where the read of `ord+0x20` at `L00775`, the zero test at `L01025` and the branch at `L01026` (to `L01027`) send each member to the store of `0xb` into `ord+0x08` at `L01028` plus the optional cast, or to the heal AI at `L00080` for a human participant's unit. **Neither of those branches walks**, so a vetoed order-5 group stands still while an empty-sighted one moves

**Confidence.** High for the mechanism: one routine read end to end, the branch target fixed by its own `rel32`, the argument forwarding cited, and the pre-clear ordering read out of `R0110`'s own instruction order. **Not established, and not establishable from the image:** how often the branch is taken in a shipped mission, which is a fact about play. `experiments/EXP-0126-swarm2-gate/EXP-0126.md` states three falsifiable predictions and what refutes each

### AI-CMDSET45-108

`R0108` (`Move`) and `R0157` (`Swarm 2`) are adjacent `0x2f0`-byte blocks, `0x2f0` apart. `tools/candcount` compares them in the file and folds away only a difference that is the tail of an `E8`/`E9` `rel32` whose two encodings differ by exactly the block distance — the same absolute target named from two places, with the opcode byte required to match. **29 bytes differ; 14 such branches account for 28 of them; one byte survives**, `L01029` = `0x04` against `L01030` = `0x05`, the immediate stored into the byte at `+0x20` — the group order itself (`AI-GROUPCMD-020`'s `L00433` and `L00434`). The fourteen shared targets are `R0007`, `R0068`, `R0069`, `R0140`, `R0019`, `R0167`, `R0140`, `R0007`, `R0068`, `R0069`, `R0299`, `R0300`, `R0301`, `R0301`.

**Consequence:** every per-member state an order-4 group holds when `R0154` runs, an order-5 group holds identically, so the fallback is coherent rather than degenerate — `Par0 = 5` is `Par0 = 4` plus the gate. What the shared routine *does* is already published: `MOVE-FORM-036` reads the formation arm and `MOVE-GATE-035` the two gates, both citing the two routines as parallel address pairs (`L01031`/`L01032`, `L01033`/`L01034`, …). This row does not re-derive them; it closes the question those pairs leave open, namely whether the two routines agree **everywhere else** as well

**Confidence.** High. The comparison is over the file's bytes rather than over two disassembly listings, and the fold rule is mechanical and stated: a listing diff of these blocks shows fourteen "differences" that are not differences, and dismissing them by eye is the same fold applied silently. **This is a sharpening, not a discovery** — [EXP-0094] had already read both routines end to end and found them to behave alike; what was not established is that they are the same bytes

### AI-BB4SWEEP-109

They are the store of the incremented count into the `+0xc` field at `L01035` (after `L01036` and the increment at `L01037`; node allocation, `R0302`) and the store of the decremented count into the `+0xc` field at `L01038` (after `L01039` and the decrement at `L01040`; node free, `R0303`, which then runs the *count reached zero, so RemoveAll* test at `L01041` and `L01042`); and the store of zero into the `+0xc` field at `L01000` in `RemoveAll` itself. All three are the shared list implementation, reached with `this` = `AImgr + 0xba8`; none is reachable by any sweep keyed on `0xbb4`, at any completeness. **The code-region hits are a different class**, and `AI-SWARM2-024`'s reading of them is corrected twice over: they are **four writes and one read** (`L01043`), not five writes; and `R0304` constructs six collections inline over `+0x844`…`+0xbc4` on a `0x1c` stride — the store at `L01044` of `L01045` into the field at `+0xba8` is a `CObject`-derived vtable (`imm:L01045` = 3 hits / 3 owners / 0 orphan) and the store at `L01046` into the field at `+0xbb4` is that collection's own count — so the two layouts coincide at exactly this displacement and diverge at the next, the sibling's run continuing at `+0xbc4` where the manager's second collection is at `+0xbc8`.

**The instrument shape**, a fourth kind for `docs/INSTRUMENT.md`'s family: a field inside an embedded object is addressed by two disjoint displacements — `outer + BIG` by everything outside it, `inner + small` by everything that writes it — and the tell is a sweep whose hits are all reads in a class big enough to embed things. The cure is to sweep the *neighbouring* displacements and read the layout off the pattern: `bd4 − bc8 = 0xc` on the second collection is what fixed `bb4 − ba8 = 0xc` on the first

**Confidence.** High as a negative claim about the instrument, because it does not rest on the sweep: it rests on reading the three routines that hold the arithmetic and on the `+4` / `+0xc` layout `AI-CANDCOUNT-105` establishes. **Unknown** whether some other instruction writes the field through a pointer neither path produces — a saved `AImgr+0xba8` used outside the list implementation would be invisible to `disp:bb4` and `disp:ba8` alike

### AI-CANDBYTE-110

the dword count at `+0xbb4` read at `L00460` and tested for zero at `L01022`, against the same `+0xbb4` count read at `L01047` with a low-byte zero test at `L01048` and a branch when zero (`rel32 +0xce`) at `L01049` to `L01050`, the arm that walks the members writing `ord+0x20 = 0` (`AI-SCORE-069`), and the same pair in the twin `R0153`. The group's own member count is read the same way, a low-byte zero test of the byte at `group+0xc`, which is that collection's `+0xc` in turn. The effect is bounded and disclosed rather than observed: at exactly `0x100` visible candidates order 5 would enter its own body, clear every member's target, and send each member down the `ord+0x08 = 0xb` or heal branch instead of engaging

**Confidence.** Medium. Both instructions are cited and the asymmetry is not in doubt; what is not measured is whether `0x100` simultaneously visible candidates is reachable on any shipped map, a census this experiment did not run. Recorded because it is a **limit**, and lifting an actor cap is exactly the kind of change that would make it reachable

## Defend and Follow states

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-DEFEND-111 | `actor+0x50 = 8` is *defend a unit*, and the whole of its engagement is fought on the protected unit's behalf, not on the defender's own. | High / Medium / Unknown | ● active | [EXP-0127](../experiments/EXP-0127-follow-pair/) |
| AI-FOLLOW-112 | `actor+0x50 = 0x11` is *follow*, it is the defend arm minus the fighting, and the two are not one behaviour parameterised. | High | ● active | [EXP-0127](../experiments/EXP-0127-follow-pair/) |
| AI-FOLLOWTAB-113 | The five-arm table at `L00323` is a two-way test wearing a jump table, and its index is a field re-read after the call rather than a return value. | High | ● active | [EXP-0127](../experiments/EXP-0127-follow-pair/) |
| AI-FOLLOWGAP-114 | Both arms enforce a minimum separation of a hard-coded 2, and the step-away target is a straight line away from the escorted unit, clamped to the playable rectangle. | High / Medium | ● active | [EXP-0127](../experiments/EXP-0127-follow-pair/) |
| AI-FOLLOWRANGE-115 | `ord+0x70` is the escort range, it is a closed field, and it can never be 0 at the arm — but only the displacement forms are closed, and a save load is not one. | High / Medium | ● active | [EXP-0127](../experiments/EXP-0127-follow-pair/) |
| AI-FOLLOWSET-116 | Three sites put an actor into an escort state, they are one shape, and two of them are literally one piece of source emitted twice. | High / Medium | ● active | [EXP-0127](../experiments/EXP-0127-follow-pair/) |
| AI-FOLLOWAUTH-117 | The editor's own two names for these states are *Defend* and *Follow*, both authoring surfaces reach them, and the resolved script node is addressed by parameter TYPE rather than by index. | High / Medium | ● active | [EXP-0127](../experiments/EXP-0127-follow-pair/) |
| AI-FOLLOWHEAL-118 | A defender carrying a spell book heals the unit it protects before it fights for it, and the spell is a hard-coded id 6. | High / Unknown | ● active | [EXP-0127](../experiments/EXP-0127-follow-pair/) |
| AI-FOLLOWDEATH-119 | Nothing guards `ord+0x10`: the per-actor machine's only "my subject has died" test is on a different state and a different field. | High / Medium / Unknown | ● active | [EXP-0127](../experiments/EXP-0127-follow-pair/) |

### AI-DEFEND-111

`R0008` arm 8 loads `ord+0x10` (`L01051`) and calls `R0138(self, that actor)` (`L01052`), returning with `0x8` bytes of arguments released, `this` = the AI manager. The arm reads the range `ord+0x70` (`L01053`) and the distance returned by `R0162` — `TRIG-DIST-014`'s Chebyshev — over the two actors' `+0x10` position objects (`L01054`), saving both. **`dist > range`** (the comparison at `L01055` and the branch to `L01056` at `L01057` when the distance does not exceed the range): `ord+0x08 = 4` (`L01058`), `ord+0x18 =` the protected unit (`L01059`), `ord+0x14 =` the range or, when the range is 0, `actor+0xa5` (`L01060`/`L00269`, selected by the zero test of the range at `L01061`) — that is `AI-ORDER-039`'s arm 4, close on the actor `ord+0x18` with stop distance `ord+0x14` — and **return**, so a defender outside its range never engages.

**`dist <= range`**: `R0305` when `actor+0x140` is set, else `R0306` (`L01056`…`L01062`), both `(self, protected unit)`; then the inner table (`AI-FOLLOWTAB-113`); then the crowding check (`AI-FOLLOWGAP-114`). **`R0306` is the cover**: it builds the occupants of the `(2r+1)²` block around **the protected unit's** cell with `r =` that unit's `mover+0x08` (`L01063`…`L01064`, `R0020`, `AI-FILTER-001`), filters them with `R0021(protectedUnit, list, 0)` (`L01065`) — so the survivors are what is hostile **to the protected unit**, which need not be hostile to the defender — and engages via `R0009` the candidate nearest *self* by `R0036`, preferring those whose virtual `[vt+0x20]` returns 3 (`L01066`/`L01067`) and falling back to any; an empty block is `R0022(self)` (`L01068`)

**Confidence.** High for the arm and for cover's polarity (both routines read end to end from a raw listing, and all 100 cited encodings re-read out of the PE with 18 branch/call anchors landing exactly — `evidence/cites.txt`, `anchors.txt`; the polarity is discriminated by the *arguments*, `ESI` = arg2 at `L01069` being the protected unit at every one of the three sites that use it). **Medium** that the block is `11×11` — `mover+0x08 = 5` rests on `L00587` plus an immediate-form sweep (57/31/0, its only mover-based hit), and a register-form store is invisible to that. **Unknown** what the virtual `[vt+0x20]` returns 3 for

### AI-FOLLOW-112

`R0008` arm `0x11` (`L01070`/`L01071`) calls `R0139(self, ord+0x10)`. Its out-of-range half is `AI-DEFEND-111`'s, store for store — `ord+0x08 = 4`, `ord+0x18 =` the unit, `ord+0x14 =` range or `actor+0xa5` (`L01072`, `L01073`, `L01074`, `L00271`). Its **within-range half is three instructions**: the comparison of the distance with `2` at `L01075` and the branch to `L01076` when it is not below at `L01077` → `R0022(self)` (`L01078`), the ordinary acquire-with-no-leash any unordered unit runs; below 2 → the step-away of `AI-FOLLOWGAP-114`. There is **no** call to `R0305`, **no** call to `R0306` and **no** indirect jump anywhere in the routine, so a follower never scans on the followed unit's behalf and never heals it — it fights only what it would have fought standing alone. `EnumRefs refto:L00323` = **1 hit / 1 owner / 0 orphan**, `L01079`, inside `R0138`, and `imm:L00323` = 0 as `docs/INSTRUMENT.md` rule 3 predicts for a SIB displacement: the inner table belongs to the defend arm alone

**Confidence.** High (the routine is 41 instructions read whole; the discriminating byte `L01075 3c02` re-reads out of the PE, and `L01077 7351` is a `rel8` fixing its target at `L01076` independently of any decode)

### AI-FOLLOWTAB-113

`L01080`/`L01081` read the byte at `+0x8` of the order block at `+0x158`, `L01082` subtracts `0x5`, `L01083`/`L01084` branch to `L01085` when the result exceeds `0x4`, and `L01079` jumps through the table at `L00323`. Read out of the PE through its own section table (`tools/followarm -mode table`) the five entries are **`L01086 L01086 L01085 L01086 L01086`** — four of five are the routine's own epilogue and the fifth is the same block the bound's `JA` falls through to. So the domain is `ord+0x08` in 5…9, which by `AI-ORDER-039` is exactly *what the engagement helper just started*: 5 and 6 the pursuit pair, **7** the sack pick-up bridge, 8 and 9 the two casts. The behaviour is one sentence — **a pursuit or a cast means the defender is busy and does nothing further this tick; anything else falls to the crowding check** — and a consumer needs no table to implement it. This is why the arm has an inner switch nobody had read: it is a compiler's emission for a `switch` four of whose cases `break`

**Confidence.** High (the table is a PE read, not a listing; its extent is fixed by `L01083 83f804` and its base by the SIB displacement `L00323` at `L01087`, both re-read; the arm map is `AI-ORDER-039`'s, quoted)

### AI-FOLLOWGAP-114

`L01085` (defend: the saved distance byte at stack offset `0x18` compared with `0x2`, taking the not-below branch) and `L01075` (follow: the same comparison) test the saved Chebyshev distance. Below 2, the escorted unit's cell is packed as an 8.8 fixed-point pair — the pair `((Y & 0xff) << 16) + (X & 0xff)` then shifted left by 8 (`L01088`), so the low word is `X*256` and the high word `Y*256` — and `R0104(self, thatPair, ord+0x70)` is called (`L01089`, `L01090`). That helper reads self's own cell **and sub-cell** (`[pos+0]`,`[pos+1]`,`[pos+4]`,`[pos+5]`, `L01091`…`R0307`) into the same 8.8 form, takes `dx = selfX − otherX` and `dy = selfY − otherY` (`L01092`/`L01093`), **forces a zero delta to 1** (`L01094`, `L01095`) so the answer is never the escorted unit's own cell, moves `ord+0x70` cells along the dominant axis *away*, rounds the minor axis onto the line through `R0279`, clamps both to `[8, dim−9]` from `[map+0x50000]`/`[map+0x50004]` (`L01096`…`L01097`) and returns `(Y<<8)|X`.

The caller then writes `ord+0x08 = 1` and `ord+0x0a =` that cell (`L01098`/`L01099`, `L01100`/`L01101`). So an escort orbits between **2** and `ord+0x70` cells of its subject. The clamp is `TERR-SIGHT-116`'s playable rectangle **recomputed inline** rather than read from `world+0x58ee0`, and in 32-bit rather than as bytes, so this site cannot exhibit the edge disagreement that row records. `callto:R0104` = **4 hits / 4 owners / 0 orphan**: these two, plus `R0100` and `R0107`, the withdraw pair

**Confidence.** High (the helper is 60 instructions read whole; the packing is checked at both ends, producer and consumer, and `[map+0x50000]`/`+0x50004` are `TERR-SIGHT-116`'s W and H). **Medium** on the rounding helper `R0279` being a plain float→int — it is read as a use, not traced

### AI-FOLLOWRANGE-115

Complete byte-store population, the `re:` search for an 8-bit store to `[E.. + 0x70]` with any source = **8 hits / 5 owners / 0 orphan**, of which two are stack slots (`L01102`, `L01103`) and **six** are the three assignment sites of `AI-FOLLOWSET-116`, each a pair: the argument (`L01104`, `L01105`, `L01106`) or the immediate **3** when the argument is 0 (`L01107`, `L01108`, `L01109`). Complete byte-read population, `re:MOV .L,byte ptr \[E.. \+ 0x70\]` = **7 hits / 4 owners / 0 orphan**, of which exactly **two** have an order block as their base — `L01053` and `L01110`, each on the instruction after the order-block pointer load from `+0x158` — the two arms themselves. `docs/INSTRUMENT.md` rule 8 (sweep by the width of the widest store that can *touch* the byte) is discharged rather than assumed: **`disp:6f`, `disp:6e` and `disp:6d` return 0 hits each, whole image**, so no 16-bit or 32-bit store reaches `+0x70` from below, and every dword hit at `disp:70` inside `L00339..L01111` was read and is a load on another base or a stack slot. Corpus, identical on **both roots**: Defend `1×2, 3×3, 4×1`; Follow `3×4, 5×5, 6×1` — **never 0**, so the 0→3 coercion is authored-unreachable, established positively

**Confidence.** High for the two reader instructions and for the six writers (three enumerations plus three zero-returning neighbour sweeps, all on the repaired table, with the width rule discharged in both directions). **Medium** that the field's write population is complete: `AI-PROGRESS-034` establishes that the order block's first `0x94` bytes are serialized with `ar.Read`/`ar.Write`, and a wholesale block move carries no displacement at all, so the **load** path writes this byte invisibly to every sweep above

### AI-FOLLOWSET-116

Each walks a group and splits on `member == the named unit`: the named unit itself gets `actor+0x50 = 0xc` — `AI-STATE-011`'s acquire-with-no-leash — with `ord+0x00 =` its own cell and `ord+0x08 = 0`; **every other member** gets the escort state with `ord+0x10 =` the named unit, `ord+0x08 = 0`, and `ord+0x70` per `AI-FOLLOWRANGE-115`. The three: `R0308` out of line (returning with `0xc` bytes released; split by the comparison at `L01112` and the branch to `L01113` at `L01114`; `L01115` → `0xc`, `L01116` → **8**, `L01117` → `ord+0x10`), and inlined in `R0190` (`L00649` → `0xc`, `L00650` → **8**, `L00651`) and `R0176` (`L01118` → `0xc`, `L01119` → **`0x11`**, `L01120`; split `L01121`/`L01122`). the `re:` search for a 32-bit store to `[E.. + 0x50]` with tail `0x8` = **2 hits / 2 owners / 0 orphan** and `…,0x11` = **1 / 1 / 0**, so those are the complete immediate-form populations.

The twin: `[L01123..R0176)` against `[L01124..R0174)`, 192 bytes at a delta of `0x1c0`, differ in **exactly 3 bytes** — `+0x32` and `+0x33`, the two changing halves of the same `R0019` call's rel32 operand (both resolving to `R0019`), and `+0x57`, the state immediate `08` against `11`. The same shape `AI-CMDSET45-108` found for the order-4 and order-5 setters

**Confidence.** High for the shape and the twin (the 192-byte comparison is a PE read printed hit by hit, and the two rel32s are shown resolving to one target). **Medium** that these are the only three: the immediate-form sweep cannot see a state assigned through a register, which is exactly the gap `AI-STATE-011`'s own Medium records and `AI-PROGRESS-034` confirmed is real

### AI-FOLLOWAUTH-117

Action opcode 6's sub-dispatch is on `Par0 − 1`, bound `L01125` comparing the index with `0x10`: **case 10 (`Par0 = 11`, catalogue "Group Command : Defend") → `L01126`**, which writes `grpAI+0x20 = 0` (`L00436`) and calls `R0308` per member (`L01127`) → state **8**; **case 14 (`Par0 = 15`, "Group Command : Follow") → `L01128`**, which calls `R0176` (`L01129`), whose own `L00437` writes `grpAI+0x20 = 0` → state **`0x11`**. The player-order road is `AI-CMD-054`'s opcode `0x1b`, `R0190`, reached once from `R0061` at `L01130` → state **8**. `R0308` has one further caller, `R0065` at `L01131`, unread here. Both group orders staying **0** is `AI-CMD-033`'s family split seen from the arm side, and it is what lets `R0023` arm 0 reach the per-actor machine at all — the world-wide sweep `R0148` that also calls it is dead code (`AI-DEAD-035`).

**Corpus, identical on both roots**, 38 maps with a parsable type-7: **6** Defend nodes and **10** Follow nodes, `AI-GROUPCMD-020`'s census reproduced. The two cases read the *same* three runtime offsets — `node+0x0c` the int, `node+0x30` the unit, `node+0x34` the group (`L01126`…`L01132`, `L01128`…`L01133`) — while the catalogue puts them in different `Par` slots (Defend: unit 1, range 2, group 9; Follow: group 1, unit 2, range 3), and Patrol's two ints land at `+0x0c` and `+0x10`; so the resolved record has typed slots, not indexed ones. Nine of the ten Follow nodes name unit `10001`; the named group holds one member in half of the sixteen nodes and eight at most

**Confidence.** High for the two cases, the census and the group-order stores (the arm table is `tools/aipatrol`'s PE read, the census runs on both roots and agrees exactly, and every cited encoding re-reads). **Medium** for the typed-slot reading — it rests on three commands agreeing, which is corpus agreement over the *image*, and no routine that fills the resolved record was read. `Description Instants.ini` exists only on the root that also ships the editor (`TRIG-CAT-026`), so the parameter **names** are one root's; the parameter **types** are in every node's own 796 bytes on both

### AI-FOLLOWHEAL-118

`L01056` reads the field at `+0x140` and `L01134` branches to `L01062` when it is zero, which picks `R0305(self, protected)` over `R0306`. `R0305` selects id **6** (`L01135` loads `0x6`) on two disjoint conditions: `self+0xa0 == 0` **and** `protected+0x94 < protected+0x96` (`L01136`…`L01137`), or `self+0xa0 < self+0x9c + 3` **and** `protected+0x94 < protected+0x96 >> 1` (`L01138`…`L01139`) — that is, at any damage when one term is clear and only below half health otherwise. It then looks the record up with `R0017(self+0x140, 6)` (`L01140`) and casts only if `record+0x0c <= self+0x9a` (`L01141` / the greater-than branch at `L01142`), setting `ord+0x60 = 1` and calling `R0018` (`L01143`/`L01144`) — `AI-STATE-011`'s arm `0xd`. Nothing cast falls straight through to `R0306` (`L01145`). `callto:R0305` = **1 hit / 1 owner / 0 orphan**, `L01146`: this wrapper exists for the defend arm and for nothing else. The follow arm never reaches it (`AI-FOLLOW-112`)

**Confidence.** High for the control flow and the id (the routine is 47 instructions read whole and every cited encoding re-reads; the two conditions are separated by `L01147 JNZ L01138`, a `rel8`). **Unknown** what spell 6 *is* — the id is taken from the instruction, not resolved against a spell table, and `+0x9a`/`+0x9c`/`+0xa0` are read as cost-side operands rather than identified

### AI-FOLLOWDEATH-119

`R0008`'s prologue reads `actor+0x50` (`L01148`), compares it with `0x3` (`L01149`) and branches to `L01150` when different (`L01151`) — it runs **only** for `actor+0x50 == 3` — and what it then tests is `ord+0x0c` (`L01152`), null-checked at `L01153` and probed for the dead marker at `L01154`, a comparison of the dword at `+0x54` with `0x10`, on which it halts the actor into `0xc` (`L00597`). Arms 8 and `0x11` hold their subject in **`ord+0x10`**, load it at `L01051` and `L01070`, and pass it to a routine that dereferences it at `L01155` / `L01156` with **no null test and no `+0x54` test** anywhere between. This experiment found nothing that clears `ord+0x10` or resets either state when the named unit dies; that absence is a **lower bound**, because `+0x10` is one of the commonest displacements in the image (the `re:` search for a 32-bit store to `[E.. + 0x10]` with a register source = 1466 hits / 773 owners / 8 orphan) and the EXP-0125 base-load filter returns **zero** here — the compiler interleaved the load and the store at all three known writers, so the filter that worked on `ord+0x00` cannot be applied to this field at all.

**Falsifiable prediction, and only a running original can settle it:** a group placed under script `Follow` or `Defend` whose named unit is then killed keeps the same behaviour aimed at the dead unit, re-issuing order 4 toward it while its object survives. Refuted by that group reverting to any other behaviour without a further script node firing, or by the followers dispersing rather than converging on the death site; a failure at the moment the object is released would establish a *stale* pointer instead, which is a different defect and is distinguishable by whether anything happens at all first

**Confidence.** High for the guard's shape and for the absence of a test on the two arms' path (four cited instructions, each re-read, and the gate `L01151` is a `rel8`). **Medium** for "nothing clears `ord+0x10`" — that is the lower bound above, and the instrument's failure is stated rather than worked around. **Unknown** what the game does, which is what the prediction is for

## Script attack order

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-SCRIPTATTACK-120 | Trigger group subcommand 10 orders one named member to acquire and every other member to attack it. | High | ● active (amended, superseded) | [EXP-0155](../experiments/EXP-0155-trigger-closure/) |

### AI-SCRIPTATTACK-120

The dispatcher clears `grpAI+0x20`, walks the referenced group, and calls `R0309(member, namedUnit)`. That helper first performs the common route/reservation cleanup. For `member == namedUnit` it sets `actor+0x50 = 0xc`, `ord+0x00 = own cell`, and `ord+0x08 = 0`: acquire-with-no-leash. For every other member it sets `actor+0x50 = 3`, `ord+0x0c = namedUnit`, `ord+0x14 = actor+0x12c`, and `ord+0x08 = 0`: engage the named target at the member's current reach. It does not resolve a hostile candidate or copy the target's owner relation. Both campaigns author five nodes and reference three. *[EXP-0170]: the dispatcher was **not** read end to end and the headline is false for one class of member. Between the group-order store and the per-member call the arm walks the group a second time and calls `R0225(member, namedUnit)`, comparing the result with `0xffffff` (`L01157`, `L01158`) — `AI-COST-071`'s target cost, whose only path to that value is a **0** cell of `AI-PREF-070`'s preference matrix. On the sentinel the arm skips `R0309` entirely and calls `R0088` (`L01159`): the vetoed member acquires **in place** instead of attacking the named unit. `AI-CMD-054` had already read the identical gate on the player's order `0x19`; it was not carried across to this arm. The row's other statements reproduce exactly, and the veto is authored-unreachable in the shipped campaign (`TRIG-GRPARM-047`).*

**Confidence.** High (dispatcher and helper read end to end; the self/non-self branch gives two distinct, already classified order shapes and excludes a group-wide acquire) — *amended by [EXP-0170]: the per-member gate was missed, so the High covered a routine that had not in fact been read whole*

**Amended.** The headline's universality and the completeness of the dispatcher reading; the helper's two branches, the field writes and the node counts are untouched and reproduce exactly: The dispatcher walks the group **twice** (superseded, EXP-0170). `retracted.md` holds the full entries.

## Minimap and command panel widgets

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-MINIMAP-156 | `R0310`, the minimap's own paint routine (`TOWN-091`, `AI-MINIMAP-062`), is read here to its own terminator, `RET` at `L01160`… | High | ● active (amended, partially retracted) | [EXP-0211](../experiments/EXP-0211-column-controls/) |
| AI-MINIMAP-157 | Widget 5 of the right-column container (`TOWN-086`, `TOWN-091`) and the minimap object of `AI-MINIMAP-062`/`AI-MINIMAP-124`/`SESS-VIEW-031` are the same object. | High | ● active | [EXP-0211](../experiments/EXP-0211-column-controls/) |
| AI-PANEL-158 | The command panel's own mouse-hit-test routine, `R0101`, is a distinct routine from `R0097`, the order-arming routine `AI-PANEL-053` calls "the command-panel handler." | High / Unknown | ● active | [EXP-0211](../experiments/EXP-0211-column-controls/) |

### AI-MINIMAP-156

`R0310`, the minimap's own paint routine (`TOWN-091`, `AI-MINIMAP-062`), is read here to its own terminator, `RET` at `L01160`, and it does call the `vt+0x18`/`vt+0x20`/`vt+0x24` external-object blit primitives `TOWN-091` states it does not use. `TOWN-091` reads roughly the first 210 of the routine's 623 disassembled lines and reports that the routine performs "a masked 2x2-block pixel copy/blend loop… not a discrete-icon blit through the `vt+0x18`/`vt+0x34` primitives every other widget in this experiment's paint routine calls." Within those same first ~210 lines the routine in fact makes three calls on the resource object at global `L01161`: `L01162` calls the vtable slot at `+0x18`, positioned from `this+0x64`/`this+0x68`; `L01163` calls the vtable slot at `+0x18`, on a second object reached through `this+0x60` or `this+0x64` chosen by a null test on `campaign+0xdc` (`L01164`-`L01165`); and, past line 210, `L01166` calls the slot at `+0x20` and `L01167` the slot at `+0x24` on the same `L01161` object.

The masked pixel loop `TOWN-091` describes runs afterward, reading a raw buffer at `[L01168]` with stride `[L01169]`, and is additional to the blit calls, not a replacement for them. Past line 210: at `L01170`-`L01171` the routine walks a count-and-pointer pair on the widget itself, `[this+0x9c0]` entries at `[this+0x9bc]`, and for the first entry whose own first dword is non-null it calls `entry->vt+0x34` with a destination point built from `this+0x60`/`this+0x64`/`this+0x68` and the caller-supplied scale — the same external-object `vt+0x34` sub-rect blit convention, on an array this experiment did not identify beyond "count and pointer fields on the minimap widget." At `L01172` the routine forks on a local flag into one of two structurally symmetric closing arms, `L01173`-`L01174` and `L01175`-`L01160`; each computes a rectangle from `this+0x60`/`0x64`/`0x68` and calls `R0311` once, then `R0312`, then returns. Neither arm performs any further blit.

**Confidence.** High for every cited instruction and control-flow edge (the routine reaches its own `RET`, 623 lines, and every branch target is walked). The identity of the `this+0x9bc` array (candidate: per-unit map icons) is Unknown; the loop's mechanism and blit convention are established, its semantic content is not

**Amended.** The clauses that the `this+0x9bc` array is a widget field and that the walk takes the first non-null entry are retracted ([`retracted.md`](retracted.md)); `MISSION-063` (EXP-0455) identifies the array as the bucket array of the map view's object map, walked node by node.

### AI-MINIMAP-157

`TOWN-091` cites `R0310` as widget 5's own paint routine, matching the class's `vt+0x2c` slot by the project's own construction/paint convention; `AI-MINIMAP-062` cites `R0233` at "slot `L00709`" as the minimap's left-button-down handler. A fresh 30-slot dump of the vtable at `L01176` (`tools/figurecontrols -mode dwords -n 30`) reads `[0x2c] -> R0310` and `[0x54] -> R0233` (absolute `L00709`) in the same table, and the same dump's `[0x4c] -> R0313` and `[0x60] -> R0314` reproduce the mousemove and right-down handlers this experiment separately read to their own terminators. `TOWN-091` names widget 5's functional identity an open question; this closes it.

**Confidence.** High (one vtable, both cited routines present at the cited offsets, re-dumped fresh; a shared vtable base is the strongest identity proof available to static reading, short of tracing the single constructor call in `R0315` argument by argument, which this experiment also did not need to do because the two paint/handler citations already pin the same table)

### AI-PANEL-158

A fresh 24-slot dump of the command panel's vtable at `L01177` reads `[0x54] -> R0101` (matching the physical `WM_LBUTTONDOWN` slot convention) and `[0x2c] -> R0316`, the panel's own paint routine. A search for raw-dword references to `R0097` finds **zero** — it is never installed as any vtable slot — against **14** direct `CALL rel32` sites in `.text`. `tools/figurecontrols`'s section-scoped raw-dword scan covers 1,904,808 of the image's 1,977,344 bytes (the 5 declared PE sections, capped at each section's own `VirtualSize`); a second, whole-file scan (`-mode rawfull`) covers all 1,977,344 bytes directly, including the PE header, inter-section alignment padding, and the appended CodeView/debug directory outside every section (72,536 bytes together), and also finds zero occurrences, on both roots.

`R0101`, read to both its `0xc`-byte-releasing returns, hit-tests all 8 cells through the global function pointer at `L01178`, gates cell 4 on `R0317` (`AI-PANEL-053`'s existing shared predicate), and on a hit posts message `0x40c` with the cell index (or `-1` to clear a stale selection first) to `[this+0x5c]->vt+0x48`. It does not call `R0097` directly, and no call site among the 14 was traced back to a `0x40c` handler in this experiment; the path connecting the panel's own click-to-cell translation to `AI-PANEL-053`'s order-arming routine remains unestablished.

**Confidence.** High for the two routines' distinctness and for `R0101`'s own read (both terminators reached, all cited addresses re-measured). Unknown for the `0x40c`-to-`R0097` call path
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

## Cursor manager, SetCursor and slot consumers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-CURSOR-172 | Setting the current cursor is one path through one manager object, and that object owns a `timeGetTime`-paced frame counter with a modulo wrap that never lets a reader see a frame at or above the registered count. | High / Medium | ● active (amended, partially retracted) | [EXP-0213](../experiments/EXP-0213-cursor-lifecycle/), [EXP-0217](../experiments/EXP-0217-cursor-composition-order/) |
| AI-CURSOR-173 | Exactly 5 call sites in the whole `.text` section reach Win32 `SetCursor`, `LoadCursorA` has none, and none of the 5 touches the cursor-manager object or any of its 29 constructed cursors. | High / Medium | ● active | [EXP-0213](../experiments/EXP-0213-cursor-lifecycle/) |
| AI-CURSOR-174 | No named Win32 API delivers the manager's own cursor set to the screen: `SetCursor` never receives one of its objects, and `SetClassLongA`/`SetClassLong` is not imported at all. | High / Unknown | ● active | [EXP-0213](../experiments/EXP-0213-cursor-lifecycle/) |
| AI-CURSOR-175 | Classifying all 69 callers of the set-cursor adapter by the nearest preceding load of one of the 28 registry slots gives a concrete surface for five of `AI-PANEL-053`/`AI-CURSOR-126`'s unnamed cursors. | Medium | ● active | [EXP-0213](../experiments/EXP-0213-cursor-lifecycle/) |
| AI-CURSOR-176 | A 29th cursor is registered at runtime, through the same constructor as the 28 static ones, and its storage field is one of the four the held-item state occupies on the session object. | High | ● active (amended) | [EXP-0213](../experiments/EXP-0213-cursor-lifecycle/) |
| AI-CURSOR-177 | The map view answers a private dispatcher of 132 message ids, `0x401`-`0x484`, and message `0x405` — already known as the right-up cancel (`MENU-COMBAT-019`) — arms `R0318`, which the corpus had not named. | High / Unknown | ● active | [EXP-0213](../experiments/EXP-0213-cursor-lifecycle/) |
| AI-CURSOR-178 | Beyond the two sites the corpus already names — `AI-PANEL-053`'s consuming-click clear and `AI-MINIMAP-062`'s minimap clear — `view+0x99c`… | Medium | ● active | [EXP-0213](../experiments/EXP-0213-cursor-lifecycle/) |
| AI-CURSOR-188 | All 28 cursor registry slots have a consumer reference, and none is reached by either of the two addressing forms a `disp32` sweep cannot see. | High | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |
| AI-CURSOR-189 | `AI-CURSOR-175`'s 21 unattributed slots are an artefact of a store-to-local idiom, not of an unread code path. | High | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |
| AI-CURSOR-190 | The eight edge arrows are selected by a screen-position test inside the mission-map hover routine, and the arrow numbering is a clockwise compass starting at north. | High | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |
| AI-CURSOR-191 | The six small cursors are selected in `R0319`, by the mission view's own armed-mode field, through an eight-entry jump table. | High / Unknown | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |
| AI-CURSOR-192 | The runtime-constructed held-item cursor is stored in `[sess+0x3d8]`, not `[sess+0x3cc]`. This corrects `AI-CURSOR-176`. | High | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |
| AI-CURSOR-193 | No path that sets no cursor clears the one already displayed. | High | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |
| AI-CURSOR-194 | The `SetCursor(NULL)` routine's address is installed as vtable slot 22 of `L01179`, and once the routine is entered the call cannot be skipped. | Medium | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |
| AI-CURSOR-195 | The second `0x99c` population `AI-CURSOR-178` flagged belongs to a different object, and the field there is an array index rather than a mode enum. | High / Unknown | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |
| AI-CURSOR-196 | The `town` cursor (slot 24, `L01180`) has exactly three consumer references, all in mission-view routines. | Medium | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |

### AI-CURSOR-172

The manager is the global at `L01181`; `L00619` (`[current-cursor]`, `AI-CLICK-050`/`AI-MINIMAP-062`) is its own `+0x04`. Every path that changes it goes through `R0320(cursorObj)`, which unpacks `cursorObj+0x4/0x8/0xc/0x10/0x14` and thiscalls `R0321(mgr, ...)`. `R0321` calls `R0322(mgr)` (decrements a manager-owned refcount at `mgr+0x20`), releases the manager's own two prior sub-objects at `mgr+0x8`/`mgr+0xc` through each one's `vt+0x4`, copies the five incoming fields to `mgr+0x4`/`mgr+0x18`/`mgr+0x1c`/`mgr+0x28`/`mgr+0x2c` (the fourth and fifth being the incoming object's own `+0x10` frame-count and `+0x14` period arguments — the same two `SPR16A-CURSOR-046` already publishes per cursor), unconditionally zeroes `mgr+0x24`/`mgr+0x30`/`mgr+0x34`, allocates two fresh sub-objects into `mgr+0x8`/`mgr+0xc`, and calls `R0323(mgr)`.

`R0323` increments the refcount and, only when it was `0` on entry and `mgr+0x4` is non-zero, first calls `R0324(mgr)`: this calls WINMM.dll `timeGetTime` (IAT `L00849`, ~~the sole caller of that import in `.text`~~ *one of 85 direct call sites of that import in `.text`, `AI-CURSOR-224`*), computes `elapsed = now - mgr+0x30`, and when `elapsed > mgr+0x2c` (the period) increments `mgr+0x24` (the frame index) and stores `now` to `mgr+0x30`; if the incremented index is not less than `mgr+0x28` (the frame count) it is reset to `0` in the same call, before any external reader can observe it. `R0324` has exactly one caller in the whole `.text` section (`L01182`, inside `R0323`; `go run ./tools/cursorlife -callers R0324`).

Because `R0321` resets `mgr+0x24`/`mgr+0x30` to `0` on every call before `R0323` runs, a single "set cursor" call can only be observed to leave the frame index at `0` or `1` (`0` again if the frame count is `1`); reaching a higher index needs `R0323` to run again later **without** an intervening `R0321` reset. `R0323` has **12** callers in `.text` (`go run ./tools/cursorlife -callers R0323`); only one (`L01183`) is inside `R0321`. The other 11 (`L01184`, `L01185`, `L01186`, `L01187`, `L01188`, `L01189`, `L01190`, `L01191`, `L01192`, `L01193`, `L01194`) were not traced in this experiment. `R0322`'s own decrement-to-zero arm (`L01195`-`L01196`) tail-jumps to `R0325`, a distinct, untraced routine, rather than to `R0324`

**Confidence.** High for the mechanism as traced — the `timeGetTime` source, the elapsed/period compare, the increment-then-wrap and the field mapping are single instructions read end to end and cross-checked against a Ghidra disassembly of `R0320`/`R0321` (`evidence/disasm-excerpts.md`). This excludes the live alternative that cursor animation is driven by a Windows timer message: `SESS-TIMER-022` already establishes, from a full PE import-directory walk, that `SetTimer`/`KillTimer`/`timeSetEvent`/`timeKillEvent` are absent from the whole image and that `timeGetTime` is the only clock imported; this experiment adds that ~~`timeGetTime`'s one call site is inside this exact chain~~ *one of `timeGetTime`'s 85 call sites is inside this exact chain, `L01197`; the uniqueness is withdrawn by `AI-CURSOR-224` and the exclusion of a timer-driven animation rests on `SESS-TIMER-022` alone*. **Medium** for whether the mechanism is ever observed to animate past frame index 1 in play: that depends on whether any of the 11 untraced callers of `R0323` runs on a cadence that does not immediately follow a `R0321` reset, which this experiment did not check. This answers `SPR16A-CURSOR-046`'s open question — "no code evidence establishes whether the cursor animator visits frames beyond its registered count" — for the negative half only: code exists that would advance a frame, and it never stores a value at or above the registered count

**Amended.** The `timeGetTime` uniqueness clause only, in the body **and** in the confidence paragraph; the manager path, the field mapping, the elapsed/period compare, the increment-then-wrap and the 12-caller enumeration all stand: A scan of the whole image for the stored dword `L00849` returns 108 hits, all in `.text`: 85 of the form of the 6-byte indirect call operand, a direct call through the slot; 22 that load `[L00849]` into a register; and one jump through `L00849` at `L01198`, an import thunk with 0 direct call sites (refuted, EXP-0217). `retracted.md` holds the full entries.

### AI-CURSOR-173

Instrument: `callSitesToIAT` (`tools/cursorlife`) scans `.text` for direct `FF15 <iat>` and `E8`-to-thunk forms against IAT slots `L01199` (`SetCursor`) and `L01200` (`LoadCursorA`); this is a complete population sweep of one call form pair against two fixed addresses, reproducible with `go run ./tools/cursorlife`. Result: `SetCursor` 5 direct, 0 via thunk; `LoadCursorA` 0 direct, 0 via thunk. The 5: `L01201`, inside `R0326`, a large routine reserving a `0x630`-byte frame that opens a registry key (`RegOpenKeyExA` on `HKEY_LOCAL_MACHINE`, `L01202`) before calling `SetCursor(NULL)` — a settings/display-mode probe with 0 direct `E8` callers in `.text`, so its own caller was not established.

`L01203`, in code this experiment's Ghidra pass could not attribute to a function boundary (orphan), not further investigated. `L01204`/`L01205`, both inside one routine, `R0327` (returning with `0x4` bytes released, one stack argument): it adds the argument to a nesting counter at `this+0xa0` and, entering (`this+0xa0>0`), pushes the fixed global `[L01206]` and calls `SetCursor`, saving the returned previous handle to `this+0xa4` on the `0->1` transition; leaving (`this+0xa0<=0`), it restores the saved handle from `this+0xa4` and calls `SetCursor` again. This is the push/pop shape of a busy-cursor nesting counter; its own callers were not traced. `L01207`, inside `R0328` (releasing `0xc` bytes, three stack arguments): it calls `L01208` for an object, and when that object's `+0x50` is non-zero pushes the fixed global `[L01209]` and calls `SetCursor`, returning `1`; otherwise it forwards to an unread routine, `L01210`, and returns its result. Neither `L01206` nor `L01209` lies inside the 28-slot cursor registry (`L01211`-`L01212`) or equals the manager (`L01181`) or `[current-cursor]` (`L00619`)

**Confidence.** High for the enumeration and the non-overlap with the manager's own state (a complete sweep of two fixed IAT slots against a fixed set of known cursor-object addresses; the alternative excluded is that some other, un-enumerated call form reaches `SetCursor` — the sweep covers both the direct and the thunked call forms, which is the complete set this repository's instruments recognise). **Medium** for identifying `R0327`/`R0328` as MFC's busy-cursor/`OnSetCursor` idiom: this is inferred from the stack-cleanup size (the `0x4`- and `0xc`-byte argument releases) and the push/pop or gate-then-fallthrough shape, not from a resolved symbol, vtable slot, or message-map entry

### AI-CURSOR-174

`AI-CURSOR-173` shows the 5 `SetCursor` call sites reach two fixed busy-cursor globals and a `NULL`, never the manager (`L01181`), `[current-cursor]` (`L00619`), or a registry slot. A walk of the complete PE import directory (`go run ./tools/panelmodal -mode imports`) finds no `SetClassLong`/`SetClassLongA`/`SetClassLongPtrA` entry from any DLL, so no code path in this image can install a cursor into the window class either. Between them these are the two ways a Win32 program ordinarily hands a custom `HCURSOR` to the desktop. **This is a falsifiable prediction, not yet checked against the running original**: if the manager's own cursor set never reaches the OS cursor by either mechanism, the system mouse cursor should stay whatever the desktop default is (or invisible, given the `SetCursor(NULL)` at `L01201`) while hovering the mission map, and the visible in-game cursor art must instead be the game's own sprite, drawn at the mouse position by its normal render loop. A refuting observation: a screen capture or a second monitor showing the true Windows cursor icon change to match the game's own art while the game window has focus

**Confidence.** High for the absence half (`SetCursor` call-site enumeration in `AI-CURSOR-173`; the import-table walk for `SetClassLong*` is exhaustive over the whole import directory, the same instrument `SESS-TIMER-022` uses for its own absence claim). This excludes the live alternative that some other `SetCursor`/`SetClassLong` call site was missed by construction, not by search depth — the import walk cannot miss a name for want of a containing function, and the `SetCursor` sweep covers both direct and thunked call forms. **Unknown** for the positive claim (a self-drawn sprite cursor): this experiment did not locate the render-loop blit that would confirm it, and only the running original can settle it

### AI-CURSOR-175

Instrument: `slotLoadSites`/`classifySetCursorCalls` (`tools/cursorlife`) finds every `MOV r32,[slot]` (forms `A1`/`8B05`) against the 28 registry addresses across `.text`, then pairs each of the adapter's 69 callers (`callSites` against `R0320`) with the nearest such load within 64 bytes back. Result, reproducible with `go run ./tools/cursorlife`: `default` 36 sites (scattered widely — the largest single group); `wait` 12 sites, all in `L01213`-`L01214`, each interleaved with a `default` site immediately after it (a set-wait/set-default pattern); `select` 9 sites; `cantput` 4 sites, all in `L01215`-`L01216`; `dice` 1 site (`L01217`); `backpack` 1 site (`L01218`); `sdefault` 1 site (`L01219`), 60 bytes before the minimap dispatcher's own established address range `R0233`-`L01220` (`AI-MINIMAP-062`/`AI-MINIMAP-124`); 5 sites (`L01221`, `L01222`, `L01223`, `L01224`, `L01225`) have no static slot-load within the window, clustered among the `cantput` sites, and are candidates for a dynamically-built cursor object such as the held-item drag cursor (`AI-CURSOR-176`) rather than a registered slot. This gives `sdefault` (small-default) exactly the "surface of its own rather than a per-hover test" shape the brief's own hypothesis named, and gives `wait`, `cantput`, `dice` and `backpack` concrete call-site populations for the first time

**Confidence.** Medium. The instrument is a proximity heuristic, not a verified data-flow trace: it does not confirm that the matched slot-load's own register is what reaches the call to `R0320`, so a load that happens to sit within 64 bytes of an unrelated call site would misclassify, and a genuine slot-load more than 64 bytes back, or reached through a form the scan does not recognise (e.g. a computed or indirect load), would misreport as "dynamic." The counts and addresses are exact and reproducible; the surface each cluster belongs to (which UI/game state actually triggers it) was not read from the containing functions and remains an inference from address locality and slot name alone

### AI-CURSOR-176

*(This headline was amended 2026-08-22 with the body below: as published it said the storage field matches `TOWN-348`'s "cursor-held item" by name, which the correction makes false -- the cursor is at `+0x3d8` and `TOWN-348`'s held item is at `+0x3cc`.)* The registry walk (`SPR16A-CURSOR-046`'s `R0220`, `R0220`-`R0329`) uses one shared constructor, `R0330`, 28 times. A whole-`.text` search for other callers of that same constructor (`callSites` against each distinct callee found in the registry) finds exactly one more call site, `L01226`, inside `R0331`, outside the registry routine's own address range. Unlike the 28 static registrations, its art path is a computed value rather than a `68 <addr>` literal string push, and its hotspot/period arguments differ from the registry's own values.

`R0331` stores the constructed cursor object into `[this+0x3d8]`, and `[this+0x3cc]` receives the routine's incoming argument 1. **Amended 2026-08-22 by `AI-CURSOR-192` ([EXP-0214](../experiments/EXP-0214-cursor-surfaces/)): as published this sentence named `[this+0x3cc]` as the cursor object's storage field, which is wrong.** The convergence with `TOWN-348` survives the correction and is unchanged in substance -- `[this+0x3cc]` still holds the cursor-held item, it is simply written from an argument rather than from the constructor. `TOWN-350` names `sess+0x3d8` in the same clear set as `sess+0x3cc`. `TOWN-348` independently names `[sess+0x3cc]` "the cursor-held item," reached from a completely different routine (`R0332`, the character figure widget's item-drop hit test, `EXP-0210`) that tests it before any of the widget's own rectangles. The two decode paths — this experiment's cursor-registration trace and `TOWN-348`'s figure-widget rect trace — converge on the identical field offset and name without either citing the other's evidence

**Confidence.** High. The live alternative excluded is that `[this+0x3cc]`/`[sess+0x3cc]` is a coincidental shared offset on two unrelated objects (the false-positive risk `docs/INSTRUMENT.md` names for a common displacement): it is excluded here because the offset was reached twice, by two structurally unrelated routines read in two separate experiments, and both readings independently name the same role (an item being carried by the cursor) rather than being asserted from the offset alone

**Amended.** The storage-field clause, in the body **and in the headline**; the constructor identity, the single extra call site, and the town ledger's independent reading of `[sess+0x3cc]` as the cursor-held item all stand; that row is named in the evidence cell, not here, because this cell's ids are read as overturned: `R0331` writes the constructor's return value to `+0x3d8` of the object (`L01227  89 86 d8 03 00 00`) and writes its own incoming argument 1 to `+0x3cc` of the object (`L01228  89 8e cc 03 00 00`, the stored value loaded from stack offset `0x18` at `L01229`) (narrowed, EXP-0214). `retracted.md` holds the full entries.

### AI-CURSOR-177

The dispatcher (prologue `R0333`) subtracts `0x401` from the message id, range-checks it against `0x83`, indexes a byte table at `L01230`, and jumps through a dword table at `L01231`; decoded whole with `go run ./tools/cursorlife -msgtable L01230:L01231:132:401`, reproducible against both roots (`evidence/msgtable.txt`). Of the 132 slots, 124 (ids `0x403`-`0x404`, `0x407`-`0x483` except `0x40c`, `0x40e`, `0x40f`) share one default arm (`L01232`); `0x401 -> L01233`, which opens by loading the frame local at offset -0x34 and then calling `R0334` (`R0334`'s own entry, confirmed by its own standard frame setup + SEH prologue at that address), `0x402 -> L01234`, `0x405 -> L01235` (`R0318` by literal jump-table value), `0x406 -> L01236`, `0x40c -> L01237`, `0x40e -> L01238`, `0x40f -> L01239`, `0x484 -> L01240` are the remaining distinct arms.

`R0318` unconditionally clears `view+0x99c` (the armed-mode field, `AI-PANEL-053`) when it is non-zero and, only when the cleared value was specifically `5` (Cast), first calls two further routines (`R0317`, `R0335`) before the clear — extending `MENU-COMBAT-019`'s "right up cancels through map message `0x405`" with the clear's own target field and the Cast-specific extra step

**Confidence.** High for the dispatcher's own structure and the `0x405 -> R0318` identity: both are a direct read of the index and jump tables, reproducible byte for byte on both roots, and this excludes the live alternative that `MENU-COMBAT-019`'s "map message `0x405`" names a different arm — the jump table has one entry for `0x405` and it is this one. **Unknown** for the semantic trigger of message `0x401`, the dispatcher's own first slot and `R0334`'s arm: no code path posting it was traced in this experiment
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### AI-CURSOR-178

Beyond the two sites the corpus already names — `AI-PANEL-053`'s consuming-click clear and `AI-MINIMAP-062`'s minimap clear — `view+0x99c` (the armed-mode field) is also cleared unconditionally by the selection routine and by three "immediate order" panel buttons that never arm it, and conditionally by message `0x401`'s arm. A `disp:0x99c` sweep across `.text` (`dispAccesses`, `tools/cursorlife`) finds every `MOV`/`CMP` form touching displacement `0x99c` off any base register. Beyond the two already-published sites: `L01241`, inside the selection routine `R0212`, unconditional. `L01242`/`L01243`/`L01244`, inside the command-panel handler `R0097` — the tails of the Guard/Stand Ground/Retreat "issue order immediately" arms `AI-PANEL-053` already names, each additionally clearing `view+0x99c` and notifying message `0x40b` right after its order call, even though none of these three buttons ever arms the field in the first place.

`L01245`, inside `R0334` (`AI-CURSOR-177`'s message-`0x401` arm), conditional on a found-flag set by walking three spatial collections at `view+0x9b8`/`view+0x9c4`/`view+0x9d4`/`view+0x9e0` (the latter three offsets not cited elsewhere in this ledger), followed by the same notify-`0x40b` step. A second, address-distinct population of `0x99c` hits exists at `L01246`-`L01247`, inside the town/shop code range this ledger's `AI-MINIMAP-156` and neighbours already work in; this experiment did **not** determine whether that population reads the same `view` object or a different struct that coincidentally shares the offset, and excludes it from the count above

**Confidence.** Medium. The `disp:` sweep is a complete scan of one instruction-form family (`docs/INSTRUMENT.md`'s enumeration checklist), so the mission-view sites listed are a real, reproducible population — but the sweep by construction cannot distinguish two different base objects sharing the same field offset, which is exactly the open question for the second population it also found. Because that question is unresolved, this row does not claim the enumeration is complete over "every clear of the armed mode," only over the mission-view sites this experiment could attribute

### AI-CURSOR-188

A whole-image little-endian dword scan for values in `[L01211, L01248)` returns **212 references**, all 212 in `.text` and **0** in `.data`, `.rdata` or any other section. **0** are SIB-indexed and **0** are unclassified; every one is a `disp32`-only load or store. 28 are in the registry constructor `R0220` and 28 in the destructor `R0329`, one per slot, leaving 156 consumer references, with at least one for each of the 28 slots. Slot `n` is the dword at `L01211 + 4n`, and joining those addresses against `EXP-0190`'s `cursor-construction.tsv` in construction order gives the slot-to-name map published as `SPR16A-CURSOR-067`. This scan is the instrument `docs/INSTRUMENT.md` rule 7 requires for a reachability question: it reads bytes, not instructions, so a pointer stored in a data table would have been found. **The residual blind spot, which this scan does not close:** a slot address computed at runtime rather than encoded in an instruction leaves no dword to find, so this is an absence of the two encoded forms, not an absence of every possible reach

**Confidence.** High. The live alternative excluded is the one this experiment's own author predicted, that `AI-CURSOR-175`'s 21 unattributed slots are reached through a SIB-indexed table read (a load addressed by a scaled index plus a 32-bit displacement) or through a pointer stored in data, neither of which a displacement sweep can see. A byte-level scan of every section cannot miss either form, and it returned 0 of both. `AI-CURSOR-189` names the actual cause

### AI-CURSOR-189

Both routines that select among many slots compute the chosen cursor into a stack local, branch to a common tail, and call the set-cursor adapter **once** from that tail. `R0219` holds 49 of the 212 registry references (`AI-CURSOR-188`) and contains exactly one adapter call, at `L01218`; the nearest preceding slot load is `L01249` (a load of the global at `L01212`, slot 27 `backpack`), 40 bytes back, so a 64-byte proximity window attributes that single call to `backpack` and nothing at all to the other 48 slot references in the same routine. `R0319` behaves identically: its adapter call at `L01219` has `L01250` (a load of the global at `L01251`, slot 17 `sdefault`) as nearest preceding load, which is why `sdefault` is the one small cursor `AI-CURSOR-175` did attribute. `AI-CURSOR-175`'s 7 attributions are correct and are exactly the routines that load a slot immediately before calling; its stated blind spot is confirmed and its cause is named here

**Confidence.** High. Both routines were decoded at instruction level from a verified prologue, in fragments covering the blocks named here rather than end to end, and the byte distance from each adapter call back to the nearest slot load is arithmetic on addresses in the committed listings. The live alternative excluded is that the unattributed slots are set somewhere `AI-CURSOR-175` did not look: `AI-CURSOR-188` enumerates every reference in the image, and all 49 of `R0219`'s lie inside that one routine

### AI-CURSOR-190

All eight arrow slots (9..16, `L01252`..`L01253`) are referenced exactly once each **inside `R0219`**, in the block `L01254`..`L01255`. **They are referenced elsewhere as well**: the whole-image scan gives 23 arrow consumer references over five routines -- `R0219` (8), `R0319` (4), `R0336` (3), `L01256` (4), `R0337` (4) -- so this row describes one routine's block and not the whole population. **`L01256` is corrected to `R0338` by `AI-CURSOR-205` ([EXP-0216](../experiments/EXP-0216-mission-cursor-selection/)): it names an address inside that routine's body, not a routine of its own, and the count of 4 is unaffected.** The block tests only four globals: `X = [L01257]`, `Y = [L01258]`, `W = [L00618]`, `H = [L01259]`.

Read whole: `X == 0` selects `arrow7` when `Y == 0` (`L01260`), `arrow5` when `Y >= H-2` (`L01261`), else `arrow6` (`L01262`); `X >= W-2` selects `arrow1` when `Y == 0` (`L01263`), `arrow3` when `Y >= H-2` (`L01264`), else `arrow2` (`L01265`); otherwise `arrow0` when `Y == 0` (`L01266`), `arrow4` when `Y >= H-2` (`L01267`), else the routine falls through to its normal hover path at `L01255`. Each of the eight stores into frame local `-0x28` and jumps to `L01268` in the common tail. The direction map is arrow0 north, arrow1 north-east, arrow2 east, arrow3 south-east, arrow4 south, arrow5 south-west, arrow6 west, arrow7 north-west. The right and bottom bands are 2 pixels deep (a subtraction of 2 before each comparison); the top and left tests are equality with 0

**Confidence.** High. Every branch in the block is a named instruction read from a verified instruction boundary, and the eight loads are the complete population **within this routine**. The direction map is corroborated twice from outside this block: `R0319` and `R0336` run the same `[L01258]` against 0 and `[L01259]-2` test and pick `arrow1`/`arrow3`/`arrow2` for the top, bottom and middle of a right-hand column; and `EXP-0190`'s `cursor-construction.tsv` gives the eight registered hotspots as a compass rose -- `arrow0` (15,5) top-centre, `arrow1` (22,8) top-right, `arrow2` (25,15) right, `arrow3` (23,22) bottom-right, `arrow4` (16,25) bottom, `arrow5` (8,23) bottom-left, `arrow6` (6,16) left, `arrow7` (8,9) top-left. That is a measurement in the shipped art, taken before this experiment and independent of any code reading, so the compass is not an interpretation laid over the branch tree. **This refutes `EXP-0213`'s H3**, the live alternative that the edge arrows have a surface of their own rather than a per-hover test: they are selected on the same surface, in the same routine, and through the same store-to-local path as every other selection `R0219` makes

### AI-CURSOR-191

`sdefault` (slot 17) and the five `.256` cursors (slots 18..22) have twelve consumer references, split 6/6 between two adjacent routines. `R0319` runs from `R0319` to its `RET` at `L01269`, followed by the jump table and twelve `0x90` padding bytes; **a different routine begins at `R0233`** (a `0x1c`-byte frame reservation, the callee-saved registers pushed, `this` kept, and the mission view loaded from `this+0x5c`), and that one is the minimap dispatcher `AI-MINIMAP-062` and `AI-CURSOR-175` already name. Its six references (`L00710`, `L00711`, `L00712`, `L00713`, `L00714`, `L00715`) **compare** each slot's `+0x4` against `[L00619]` to dispatch an order; they select nothing. The selection is `R0319`'s alone.

That routine loads the mission view from `this+0x5c` at `L01270` (`this` having been saved at `L01271`) and then uses that register as an object base, not as a frame pointer, so the view's `+0x99c` is the same armed-mode field `AI-CURSOR-178` names on the mission view. The selection: `view+0x140 == 0` gives `sdefault` outright (`L01272`); otherwise mode 0 gives `smove` and modes 1..8 index the jump table at `L01273`, whose eight dwords are `L01274, L01275, L01275, L01276, L01277, L01274, L01275, L01278`, mapping mode 1..8 to `sattack, smove, smove, sdefend, scast, sattack, smove, spatrol`; any mode above 8 gives `smove` (a comparison with `0x7` and an unsigned-above branch). A test of the view's `+0x144` against `0x24` then overrides the result with `sdefault` when either bit is set (`L01275`).

Four further branches reach `sdefault` from the block `L01279`..`L01280`, which this experiment did not decode. Finally `[sess+0x3cc] != 0` overrides everything with `[sess+0x3d8]`, the held-item cursor (`L01281`, `AI-CURSOR-192`). **Decoded by `AI-CURSOR-208` ([EXP-0216](../experiments/EXP-0216-mission-cursor-selection/)): the four branches are a zoom-scaled rectangle test against a widget read from `view+0x80`/`+0x84`/`+0x88`, the widget itself not identified. `AI-CURSOR-207` ([EXP-0216]) further finds that `R0319` selects among four edge-arrow cursors in a block before this one, at `L01282`-`L01283`.**

**Confidence.** High for the selection tree and the jump table, which are named instructions and eight dwords read from the image. High for the routine boundary at `R0233`, which is fixed by a prologue, a `RET` and twelve padding bytes in the committed hex dump, and which `AI-CURSOR-175` and `AI-MINIMAP-062` both already published -- **the first draft of this row called `R0319` the minimap dispatcher, which those two rows contradict**. The live alternative excluded for the selection scope is that some other routine also selects a small cursor: the twelve references are the complete population (`AI-CURSOR-188`) and the six outside `R0319` were read and are comparisons. **This refutes the `sdefault` half of `EXP-0213`'s H3**: `sdefault` is not selected by a surface of its own but by an armed-mode test. **Unknown** for the meaning of `view+0x140` and of the two bits `0x24` in `view+0x144`, which are read here but not traced to a writer. **Resolved by `AI-CURSOR-202` and `AI-CURSOR-203` ([EXP-0216](../experiments/EXP-0216-mission-cursor-selection/)): `view+0x140` is the selection-loop's own object count, and the `0x24` mask is the same one `R0082` tests on itself to blank the command panel.**

### AI-CURSOR-192

`R0331` is a thiscall with four arguments, pinned by its own epilogue constant a `0x10`-byte argument release at `L01284` and by the 20 bytes it pushes after the SEH prologue, so at `L01229` argument `n` sits at stack offset `0x14+4n`. It passes argument 3 as the art path to the shared cursor constructor `R0330` with hotspot `(0x28, 0x28)` and period `0x3b9aca00` (`L01285`..`L01226`), then writes the constructor's return value to `+0x3d8` of the object (`L01227`), argument 1 to `+0x3cc` of the object (`L01228`), argument 2 to `+0x3d0` (`L01286`) and argument 4 to `+0x3d4` (`L01287`). `TOWN-350` reads the same four fields from the opposite direction: `R0339` writes the item pointer to `sess+0x3cc`, the slot id to `sess+0x3d0` and the constant 1 to `sess+0x3d4`, and `R0340` zeroes `sess+0x3cc` **and `sess+0x3d8`** and sets the other two to -1.

A `disp:3d8` sweep gives **14 containing-routine labels** for the field: `R0219`, `L01288`, `R0341`, `R0331`, `R0340`, `R0319`, `R0336`, `R0338`, `L01256`, `R0337`, `R0342`, `R0343`, `L01289`, `L01290`. **`R0338` and `L01256` are the same routine, not two** (`AI-CURSOR-205`, [EXP-0216](../experiments/EXP-0216-mission-cursor-selection/)): the sweep's own containing-function lookup split one routine's `+0x3d8` references across a real prologue and a false one, so those 14 labels are **13 distinct routines**. Three were read: `R0342` sets it (`L01291` reads it from the object at `+0x3d8`, `L01221` calls `R0320`), and `R0319` (`L01281`) and `R0336` (`L01292`) apply it through the same idiom -- test `[+0x3cc]`, load `[+0x3d8]`, idempotence guard, adapter.

`R0340` clears it (`TOWN-350`). **`R0344` does not touch `+0x3d8`**: it restores slot 0 `default` (`L01293` loads the global at `L01211`, `L01294` calls `R0320`), which is the paired restore rather than a consumer of the field. ~~Seven~~ *eight* of the thirteen have now been read: `R0331` and `R0342` here, `R0340` in `TOWN-350`, `R0319` and `R0336` here, and `R0338` and `R0337` in [EXP-0216](../experiments/EXP-0216-mission-cursor-selection/) (`AI-CURSOR-205`, `AI-CURSOR-206`), which find the same test-`[+0x3cc]`/load-`[+0x3d8]` idiom in both. ~~The six not read are `R0219`, `L01288`, `R0341`, `R0343`, `L01289` and `L01290`~~ *-- `AI-CURSOR-230` ([EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/)) reads `R0219`'s reference: `L01295` reading the object's `+0x3d8` behind a comparison of `+0x3cc` with zero at `L01296`, the same test-`[+0x3cc]`/load-`[+0x3d8]` idiom this row names in `R0319` and `R0336`. **The five not read are `L01288`, `R0341`, `R0343`, `L01289` and `L01290`.***

**Confidence.** High. The live alternative excluded is `AI-CURSOR-176`'s own reading, that the constructed cursor goes to `[this+0x3cc]`. Three independent facts exclude it: the store at `L01227` names `+0x3d8` and takes the constructor's return register, while the store at `L01228` names `+0x3cc` and takes a value loaded from the argument area; the argument displacements are fixed by the routine's own `0x10`-byte argument release, so they do not depend on the constructor's calling convention; and `TOWN-350`, decoded in a different experiment from a different routine, lists `sess+0x3d8` in the clear set beside `sess+0x3cc` and gives `sess+0x3cc` a value `R0331` receives as an argument. `AI-CURSOR-176`'s headline, that a 29th cursor is built at runtime through the same constructor, is unaffected

### AI-CURSOR-193

Nothing outside the set-cursor adapter writes `[L00619]`: a whole-image dword scan of that address returns 23 references and **every one is a read**. So an exit that calls no adapter leaves the pointer unchanged, whatever routine it is in. Six such exits were read. In `R0219`: the entry gate (`L01297`, comparing the state word at `+0x3dc` with 1 and jumping to the exit `L01298` unless it is 1) returns when the surface state word is not exactly 1; the hit-test call to `R0218` followed by a jump to `L01298` (`L01299`) returns after a failed hit test; a zero test of the chosen-cursor local (frame local `-0x28`, `L01300`, jumping to `L01298`) returns when the chosen-cursor local, initialised to 0 at `L01301`, was never written; and a comparison of the chosen cursor's `+0x4` with the value saved at entry (frame local `-0xe4`, `L01302`, jumping to `L01298` on equality) returns when the chosen cursor's `+0x4` already equals `[L00619]` as saved at entry (`L01303`).

`R0319` carries the last two at `L01304` and `L01305`, and has **a seventh** at `L01306` (a jump to `L01307` when the call through `L01178` returns 0), inside a block this experiment did not decode. **The no-cursor state is a no-change state.** The idempotence guard also means the adapter is never called for a cursor already displayed

**Confidence.** High for the general statement, and it rests on the read-only scan rather than on the exit enumeration: `[L00619]` has 23 references image-wide and no writer outside the adapter, which `AI-CURSOR-172` independently establishes as the writer. The live alternative excluded is that some routine clears the global directly to hide the pointer -- a byte-level scan of every section cannot miss such a store, and there is none. **The exit enumeration itself is not complete**: the two routines were decoded in fragments, the seventh exit was found by the adversarial pass in a block this experiment skipped, and `R0211`, `R0336`, `L01256` and `R0337` also select among slots and were not read. **Corrected by `AI-CURSOR-204`..`209` ([EXP-0216](../experiments/EXP-0216-mission-cursor-selection/)): `R0336`, `L01256` (which is `R0338`, not a separate routine) and `R0337` are now read and each carries the same three no-op exit shapes named here; `R0211` does not select among slots — `AI-CLICK-050` establishes it compares the current cursor against each registered cursor's own `+0x4` and dispatches an order, and `AI-CURSOR-204` reads the one slot reference `AI-CURSOR-196` names inside it (`L01308`) as exactly that comparison — and the seventh exit's own block (`R0319`, `L01282`-`L01283`) is also now read.** `AI-CURSOR-194` gives the one `SetCursor(NULL)` site, which is outside the manager

### AI-CURSOR-194

`EXP-0213` reported `R0326` as having 0 direct callers and called that a weak negative. Direct `E8 rel32` callers in `.text` are indeed **0**. A whole-image dword scan finds exactly **one** stored pointer to `R0326`, at `L01309` in `.rdata`. A second scan finds exactly **one** reference to `L01179`, the start of the table that address lies in, at `L01310`, inside `L01311` (a 6-byte store of `L01179` to offset 0 of the object being constructed), which is a vtable install. `(L01309 - L01179) / 4 = 22`, so `R0326` is slot 22 of that vtable. The installing routine begins at `L01312` and is called from `L01313` inside the routine at `L01314`.

Within `R0326` (252 instructions, `R0326`..`L01315`, decoded with 0 stops from its standard frame setup and SEH registration prologue (`-1`, `L01316`)) the 6-byte indirect call at `L01201` through `L01199` with a pushed `0x0` argument at `L01317` is unconditional: no return occurs anywhere before it, and the highest branch target in the whole routine is `L01318`

**Confidence.** Medium. The two byte-level scans are complete for both pointer values, which excludes the alternative that nothing anywhere in the image can reach the routine. What is **not** established is that anything calls vtable slot 22 of `L01179`: no call through that slot was searched for or found, so reachability here is by construction, a constructor installing the table, rather than by a traced call. `AI-CURSOR-173` and `AI-CURSOR-174` stand -- this call passes no manager cursor and is not part of the manager's cursor set

### AI-CURSOR-195

`[base+0x99c]` is accessed in 15 routines. Two further sweeps separate them. Six routines touch both `+0x99c` and `+0x910` -- `L01319`, `L01320`, `L01321`, `L01322`, `L01323`, `L01324`, all in the code-region range. Two routines touch both `+0x99c` and `+0x98c` -- `R0219` and `R0211`, both in the code-region range and both named by `AI-CURSOR-178` as mission-view routines. **No routine touches both `+0x910` and `+0x98c`.** At the code-region sites the field is consumed by reading `+0x99c` and the array base at `+0x910`, forming `(v<<5) - v`, and loading the dword at `base + 4*((v<<5)-v) + 0x74` (`L01246`), which computes `((v*32) - v)*4 = v*124` -- an index into an array whose base is `+0x910`, with a 124-byte element stride. The sibling loop at `L01325` confirms the stride by adding `0x7c` per iteration. `AI-CURSOR-178` reads the mission view's `+0x99c` as a small enum cleared to 0. The two readings cannot describe the same field

**Confidence.** High. The live alternative excluded is that the two populations are the same field and differ only in which call sites each sweep happened to sample. Two facts exclude it: the routine sets are disjoint under `+0x910` versus `+0x98c`, so the split is a property of the routines rather than of the sampling; and the instruction shape at the code-region sites uses the value as a multiplied array subscript, which a 0-or-small-enum mode field cannot be. **Unknown** for what object the code-region population belongs to: its routines are named, its class is not. `+0x9a4` was tried as a discriminator first and discarded, being touched from both ranges

### AI-CURSOR-196

`AI-CURSOR-175` attributed no call site to it. The whole-image scan of `AI-CURSOR-188` finds three references outside the registry constructor and destructor: `L01326` and `L01327`, both inside `R0219`, the mission-map hover routine; and `L01308`, inside `R0211`, which `AI-CURSOR-178` names as a mission-view routine. There is **no** reference to slot 24 in any routine of the town-screen transition family (`TOWN-372`), and none in any data section. The cursor named `town` is selected while the mission map is on screen, not while the town screen is

**Confidence.** Medium. The reference population is complete and its containing routines are named, which excludes the alternative that the slot is set from a town-screen routine the proximity instrument missed. What is **not** established is the condition under which either site is chosen: both lie inside the region `L01328`..`L01329` of `R0219` that this experiment did not decode, so what the pointer is over when `town` is selected is unknown. **`AI-CURSOR-209` ([EXP-0216](../experiments/EXP-0216-mission-cursor-selection/)) reads all three references: the two inside `R0219` are each gated by the same selection-count/CUnit-flag/object-flag/hit-test-bit condition, with the hit-test routine itself and the two flag bits' game meaning still unread, and `L01308` is a comparison arm of `R0211`'s cursor-to-order dispatch (`AI-CLICK-050`), not a selection.**

## Cursor selection sites, edge arrows and composition order

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-CURSOR-202 | `view+0x140` is the mission view's own count of currently selected objects, incremented once per object the selection-rebuild loop accepts. | High | ● active | [EXP-0216](../experiments/EXP-0216-mission-cursor-selection/) |
| AI-CURSOR-203 | The `view+0x144` mask `0x24` that `R0319` tests to force `sdefault` is the same mask, on the same field, that `R0082` tests on itself to blank the command panel. This resolves `AI-CURSOR-191`'s other Unknown. | High | ● active | [EXP-0216](../experiments/EXP-0216-mission-cursor-selection/) |
| AI-CURSOR-204 | `R0336` selects one of three edge-arrow cursors on the mission view's right screen edge, falls back to `default` on a second hit test, and applies a held-item override — one of the "unread selector routines" `AI-CURSOR-193` names. | High | ● active | [EXP-0216](../experiments/EXP-0216-mission-cursor-selection/) |
| AI-CURSOR-205 | The routine `AI-CURSOR-190`/`192`/`193` cite as `L01256` does not exist as a separate function. It is a false function boundary from `tools/cursorsurf`'s unchecked whole-`.text` scan for `0xE8` bytes… | High | ● active | [EXP-0216](../experiments/EXP-0216-mission-cursor-selection/) |
| AI-CURSOR-206 | `R0337` repeats `R0338`'s arrow-and-bottom-edge selection exactly, without the held-item-legacy branch, and is a second, separate routine from the one beginning immediately after its own `RET`. | High | ● active | [EXP-0216](../experiments/EXP-0216-mission-cursor-selection/) |
| AI-CURSOR-207 | `R0319` also selects among four edge-arrow cursors, in a block before the armed-mode tree `AI-CURSOR-191` already describes; this is where the routine's 4 arrow references (`AI-CURSOR-190`) and its seventh no-op exit… | High | ● active | [EXP-0216](../experiments/EXP-0216-mission-cursor-selection/) |
| AI-CURSOR-208 | The block `AI-CURSOR-191` names but does not decode, `L01279`-`L01280`, is a zoom-scaled rectangle test against a widget rectangle read from `view+0x80`/`+0x84`/`+0x88`; what the widget is was not established. | Medium | ● active | [EXP-0216](../experiments/EXP-0216-mission-cursor-selection/) |
| AI-CURSOR-209 | Both `town`-slot references `AI-CURSOR-196` names share one condition: exactly one `CUnit` selected, that unit's own `+0x18c` bit `0x1` set, and a hover hit-test result (`R0218`, read whole by `AI-CURSOR-231`) with bit `0x800` set. | Medium | ● active (amended) | [EXP-0216](../experiments/EXP-0216-mission-cursor-selection/), [EXP-0220](../experiments/EXP-0220-pickup-cursor-route/) |
| AI-CURSOR-218 | The cursor's own blit is `L01330`, a call through the vtable slot at `+0x18` on the cursor object at `mgr+0x4`, inside `R0345`, and it writes into the back surface. This closes the Unknown half of `AI-CURSOR-174`. | High / Medium | ● active | [EXP-0217](../experiments/EXP-0217-cursor-composition-order/) |
| AI-CURSOR-219 | The cursor is the last write into the back surface before any pixel becomes visible, so it covers everything the widget tree paints. | High | ● active | [EXP-0217](../experiments/EXP-0217-cursor-composition-order/) |
| AI-CURSOR-220 | The answer does not differ per mission panel surface: the cursor is above the control panel, the character figure widget, the worn box, the pack bar and the minimap column alike, for the same structural reason. | High / Medium | ● active | [EXP-0217](../experiments/EXP-0217-cursor-composition-order/) |
| AI-CURSOR-221 | Both Esc panels are below the cursor, and the answer is the same for the mission panel and the town panel. | High / Medium / Unknown | ● active | [EXP-0217](../experiments/EXP-0217-cursor-composition-order/) |
| AI-CURSOR-222 | The same cursor blit is issued a second time, directly onto the primary surface, by a path that runs outside any present and uses its own backing store. | High / Medium | ● active | [EXP-0217](../experiments/EXP-0217-cursor-composition-order/) |
| AI-CURSOR-223 | Nothing moves the cursor's position in the composition order. What state changes is whether the cursor is composited at all, and over which rectangle. | High / Unknown | ● active | [EXP-0217](../experiments/EXP-0217-cursor-composition-order/) |
| AI-CURSOR-224 | Correction to `AI-CURSOR-172`: the WINMM import at IAT slot `L00849` has 85 direct call sites in `.text`, not one. The row's conclusion is unaffected. | High | ● active | [EXP-0217](../experiments/EXP-0217-cursor-composition-order/); corrects `AI-CURSOR-172` |
| AI-CURSOR-225 | Correction to `TOWN-209`: `R0346` and `R0312` are the back surface's Lock and Unlock, not a clip save and restore. | High | ● active | [EXP-0217](../experiments/EXP-0217-cursor-composition-order/); corrects `TOWN-209` |

### AI-CURSOR-202

`R0082` (`AI-PANEL-061`'s sole writer of `view+0x144`) resets `view+0x140` to 0 at `L01331`, in the same block that resets `view+0x144` (`L00691`) and `view+0x148` (`L01332`). The per-object loop that walks the current selection reads `+0x140` of the object at `L01333`, adds 1 at `L01334`, and stores the result back at `L01335`, once per pass through a loop bounded by a comparison with the per-source array size at `+0x8` (`L01336`, `L01337`) and gated by a zero test of frame local `-0x68` with a jump to `L01338` on a per-object test (`[obj+0x7c]` non-zero). The same iteration sets `view+0x138` to the first accepted object (`L01339`), matching `AI-PANEL-061`'s "primary object `view+0x138`".

`R0082` reads `view+0x140` twice more, both as integer comparisons: a comparison of `view+0x140` with zero (jump when not positive) at `L00703` (guarding ownership bit `0x4`) and a comparison of `view+0x140` with zero (jump when zero) at `L01340` (guarding the panel-blank notify, `AI-CURSOR-203`). `AI-CLICK-050` independently reads `view+0x140` the same way, as a plain selected/not test, at its own routine's entry. This resolves `AI-CURSOR-191`'s Unknown for `view+0x140`

**Confidence.** High. The live alternative excluded is that `0x140` is a boolean latch rather than a count: the store at `L01335` is an increment of the field's own prior value inside a loop that can run more than once per call, and every consumer treats it with integer comparisons (`==0`, `<=0`, `==1`) rather than a bit test

### AI-CURSOR-203

`AI-PANEL-061` names the writer of `view+0x144` and the set sites of bit `0x4` (ownership mismatch, `L00707`) and bit `0x20` (a selected `CStructure`, `L00700`), but does not connect the field to `R0319`'s use of it. `R0082` reads its own field at the end of the same routine: a comparison of `view+0x140` with zero jumping to `L01341` when not positive gates bit `0x4`, and separately a comparison of `view+0x140` with zero jumping to `L01342` when zero, followed by a mask of the flags with `0x24` and a jump to `L01343` when the result is zero, at `L01344`/`L00708` — when nothing is selected, or when something is selected and either bit `0x4` or bit `0x20` is set, the routine notifies the panel with message `0x40a` and two zero arguments (`L01342`); only when something is selected and neither bit is set does it fall through to notify `0x409` with `AI-PANEL-060`'s capability mask. `R0319` tests the identical mask on the identical field at `L01275` (a test of `+0x144` against `0x24`, jumping when zero) to override its own selection with `sdefault` instead of the armed-mode pick

**Confidence.** High for the correlation: both routines read the same struct field with the same immediate mask, one write site (`view+0x144`'s sole writer per `AI-PANEL-061`) and two same-mask readers producing an analogous neutral outcome. Not established beyond `AI-PANEL-061`'s own statement: the game-level meaning of bits `0x1`, `0x2`, `0x8` and `0x20`. `AI-PANEL-061` gives `0x20`'s set site, a `strcmp` against `"CStructure"`, and lists `0x20` in the same not-established set as `0x1`, `0x2` and `0x8`; this row adds one consumer for `0x20` and for `0x4` — the `0x24` mask read here — and none for `0x1`, `0x2` or `0x8`

### AI-CURSOR-204

Bounded `R0336`-`L01345` (`RET`), one direct caller (`L01346`, `-callers`/`-fnof` verified), NOP-padded on both sides. Two calls to `R0347` and a vtable dispatch through `[+0x7c]` fetch the mission view into a local; a hit test through the shared function pointer `[L01178]` (`AI-CURSOR-193`) on the current pointer position gates everything after — a miss (a jump to `L01347` when zero) is a no-op exit. On a hit, a right-edge test (`X >= [L00618]-2` and `[view+0x3dc]==1`) picks one of three registry slots by `Y`: `L01348` (`arrow1`, north-east) when `Y==0`, `L01253` (`arrow3`, south-east) when `Y >= [L01259]-2`, else `L01349` (`arrow2`, east) — matching `AI-CURSOR-190`'s count of 3 arrow references for this routine, with the names taken from `SPR16A-CURSOR-067`'s construction order.

Outside the right edge, a second call through the same hit-test pointer, on failure leaves the chosen-cursor local unset and on success picks `L01211` (`default`, slot 0). At the join, `[view+0x3cc]!=0` overrides the pick with `[view+0x3d8]` (`AI-CURSOR-192`). Three no-op exits reach the epilogue with no adapter call: the hit-test miss, the local still unset, and an idempotence match against the displayed cursor's own `+0x4` (`L01350`) — the same three shapes `AI-CURSOR-193` establishes for `R0319`. `AI-CURSOR-193`'s unread-routine list also names `R0211`; that routine is `AI-CLICK-050`'s order dispatcher, which compares the current cursor `[L00619]` against each registered cursor's own `+0x4` handle, one arm per cursor, and dispatches an order rather than selecting a cursor.

The one slot reference `AI-CURSOR-196` names inside it, `L01308`, is read here and is that arm: the global at `L01180` is loaded at `L01351`, then its `+0x4` is read, saved to frame local `-0x40`, compared with frame local `-0x1c` and a jump to `L01352` taken when different, through `L01353`, with no store into a chosen-cursor local. `L01351` is the target of a not-equal jump at `L01354`, so that boundary is fixed independently of the linear decode. `R0211` is not part of this population

**Confidence.** High. RET-bounded, single verified caller, every branch a named instruction from a verified boundary

### AI-CURSOR-205

The routine `AI-CURSOR-190`/`192`/`193` cite as `L01256` does not exist as a separate function. It is a false function boundary from `tools/cursorsurf`'s unchecked whole-`.text` scan for `0xE8` bytes; every reference attributed to it belongs to the routine at `R0338`. `-callers L01256` reports one caller, `L01355`. Disassembly of `L01356`-`L01357` shows `L01355` is not an instruction boundary: it is the displacement byte of a load from the frame local at -0x18 (`L01358`-`L01355`, three bytes), followed by a comparison of the dword at `+0x4` with zero. Reading the same five bytes as an `E8` relative call with a 32-bit displacement computes target `L01359 + 0x00047983 = L01256`, reproducing the phantom exactly.

`-xrefs L01256:L01256` finds 0 stored dword references to that address in any section, excluding a vtable or jump-table entry. `R0338` is a genuine bounded routine: prologue (a 0x18-byte local frame, four callee-saved register saves and a save of the receiver), preceded by nine `0x90` padding bytes (`L01360`-`L01361`), themselves preceded by a `RET` at `L01362` with three instructions between the two, one direct caller (`L01363`, inside `R0348`), and a single `RET` at `L01364`, after which `R0339` begins cleanly (the routine `AI-CURSOR-192`'s `TOWN-350` cross-reference already names). The disputed byte falls inside a load of the global at `L01257` (`L01365`-`L01366`), mid-instruction, with continuous flow from `R0338` through it and no `RET` in between.

`R0338`'s own structure: after the shared hit test, a state-bit test (bit `0x2` of `[view+0x3dc]`) branches to a held-item-legacy path unique to this routine among its three siblings — with `[view+0x3cc]!=0` (a zero test with a jump to `L01367` at `L01368`), it picks `[view+0x3d8]` when either `[[view+0xf0]+0x144]` is non-zero (a non-zero test with a jump to `L01369` at `L01370`) or `[[view+0x3cc]+0x18]==1` (a comparison of `[..+0x18]` with 1 and a jump to `L01369` when equal, at `L01371`), and registry slot 23, `cantput` (`L01372`, `SPR16A-CURSOR-067`), only when NEITHER holds: both branches jump to `L01369`, and the `cantput` load at `L01373` is their fall-through. With `[view+0x3cc]==0` the branch loads no cursor slot at all and jumps straight to the common tail at `L01367` (`L01374`); so do its other two outcomes, by `JMP` from `L01375` and `L01376`, so the whole legacy branch bypasses the held-item override at `L01377`.

`evidence/disasm-listings.txt` prints this sequence at `L01370`-`L01369`; an earlier draft of this sentence stated the polarity the other way round, and that draft error is not a further `tools/cursorsurf` artifact — the instrument printed the bytes and both branch targets correctly. On the other branch, the same right-edge test as `R0336` (`arrow1`/`arrow3`/`arrow2`) plus, uniquely among the three, a bottom-edge test (`Y >= [L01259]-2`, state`==1`) picking `L01378` (`arrow4`, south) — four arrow picks, matching `AI-CURSOR-190`'s count of 4 for "`L01256`". A second hit-test call on failure of both edge tests picks `default` on success and loads no cursor slot on failure, then the same held-item override, the same three-exit idempotence pattern, and a call to `R0320`.

`AI-CURSOR-190`'s arrow census, `AI-CURSOR-192`'s `disp:3d8` sweep, and `AI-CURSOR-193`'s unread-routine list each name `L01256`; the correct address in each is `R0338`, and no count, headline or conclusion those rows draw depends on the address being different from what it names — only the address token is wrong

**Confidence.** High. The live alternative excluded is that `L01256` is a genuine alternate entry point: excluded by the absence of any `RET` before it in a continuous decode from `R0338`, by zero stored dword references to it in any section, and by the disputed byte's own position mid-instruction inside a `MOV` whose other four bytes are a verified data address

### AI-CURSOR-206

Bounded `R0337`-`L01379` (`RET`), NOP-padded to `R0349`, where a structurally different routine (5-register prologue with a 0x20-byte local frame) begins. `-fnof`/`-callers R0337` report a caller at `L01380` as "in `R0337`". That call site is in neither `R0337` nor the routine at `R0349`. `R0349` ends at its own `RET` at `L01381`, which `TOWN-093` gives independently of this experiment, decoding that routine whole as widget 8's `vt+0x2c` handler. The listing `-dis L01385:L13129` in `evidence/disasm-listings.txt` prints that `RET`, the eleven `0x90` bytes that follow at `L01382`-`L01383`, a third routine beginning at `R0350` (opening with a read of its first stack argument, a `0x10`-byte frame reservation and a register save), the call site `L01380` itself and that routine's `0xc`-byte-releasing return at `L01384`, in one continuous decode with 0 stops; its start `L01385` is the target of the not-equal jump to `L01385` at `L01386` that `TOWN-093` names, so the range begins on a verified instruction boundary.

`L01380` is inside that third routine, two `RET`s past `R0337`. The tool's containing-function report groups the call site under the nearest preceding heuristic hit rather than the true boundary — a display-grouping artefact of the same unchecked-boundary heuristic as the `L01256` case, but one that fabricates no address. `go run ./tools/claim -k "L01380"` and `-k "R0350"` each return this row and no other, so neither the call site nor its containing routine is named elsewhere in the corpus; `-k "R0349"` returns this row and `TOWN-093`, where the `R0349` routine start is already published (`evidence/corpus-searches.txt`). `R0337`'s own selection tree is the same shape as `R0338`'s arrow/bottom-edge block: right edge picks `L01348`/`f4`/`e4` (`arrow1`/`arrow3`/`arrow2`) by `Y`, and a bottom-edge test (state`==1`, no right-edge match) picks `L01378` (`arrow4`) — the same four slots, matching `AI-CURSOR-190`'s count of 4 for this routine. It has no state-bit-`0x2` held-item-legacy branch: after the arrow tests, a second hit-test call on failure picks `default` on success or leaves the local unset, then the same held-item override and the same three-exit idempotence-guarded call to `R0320`

**Confidence.** High for the routine's own bounds and selection tree, read from a verified prologue to its `RET`. The live alternative for the boundary — that `L01380`'s containing routine really is `R0337` — is excluded by two intervening `RET`+padding+new-prologue sequences, at `L01379`-`R0349` and at `L01381`-`R0350`, both printed in `evidence/disasm-listings.txt` and the second's terminator also published by `TOWN-093`

### AI-CURSOR-207

`R0319` also selects among four edge-arrow cursors, in a block before the armed-mode tree `AI-CURSOR-191` already describes; this is where the routine's 4 arrow references (`AI-CURSOR-190`) and its seventh no-op exit (`AI-CURSOR-193`) live. Read whole from `L01282` to `L01283`. The block's fall-through end is the mission-view load `AI-CURSOR-191` cites, the load from `this+0x5c` at `L01270`. After loading the current cursor, a hit test through `[L01178]` on the current pointer position; a miss (a jump to `L01307` when zero) is `AI-CURSOR-193`'s seventh, previously-undecoded exit. On a hit: a right-edge test (`X >= [L00618]-2`, `[view+0x3dc]==1`) picks `L01348` (`arrow1`, NE) when `Y==0`, `L01253` (`arrow3`, SE) when `Y >= [L01259]-2`, else `L01349` (`arrow2`, E); otherwise a top-edge test (`Y==0`, state`==1`) picks `L01252` (`arrow0`, N).

Each of the four jumps directly to `L01387`, the convergence point `AI-CURSOR-191`'s `+0x144`-against-`0x24` override guards, bypassing the mission-view load, `AI-CURSOR-208`'s bounds-check block, the `view+0x140==0` test, the armed-mode jump table, the `0x24` override itself and the `sdefault` load at `L01250`. The held-item override at `L01388` is past `L01387` and does apply to an arrow pick. That is exactly 4 slot loads, matching `AI-CURSOR-190`'s count for this routine. Together with `R0336`, `R0338` and `R0337` (`AI-CURSOR-204`..`206`), these four routines' edge-arrow tests cover north, north-east, east, south-east and south; none of the four tests west, south-west or north-west, which are read only inside `R0219`'s own block (`AI-CURSOR-190`)

**Confidence.** High. Every branch is a named instruction from a verified boundary, the exit matches `AI-CURSOR-193`'s own citation of address and condition, and the slot count matches `AI-CURSOR-190`'s independently-obtained whole-image census exactly

### AI-CURSOR-208

Read whole from `L01270`, where the mission view is loaded. A value at `[this+0x68]` — read at `L01283` from `this` as saved at `L01271`, not the mission view, which is loaded from `this+0x5c` at `L01270` (`AI-CURSOR-191`) — is used twice. Its sign selects between two derivations of a half-width and half-height from `[[view+0x80]+0x4]` and `[[view+0x80]+0x8]`, the `view+0x80` load being at `L01389`: non-negative takes a left shift then a right shift by 1 (`L01390`, `L01391`), negative takes a right shift by 2 (`L01392`, `L01393`). Its value, clamped to 0 when negative (a zeroing at `L01394`), is the shift count used in the first derivation and on both differences.

Those combine with the current pointer position around two fixed constants: `0x48` against `X` (the constant `0x48` loaded at `L01395` and `L01396`, subtracted from `[L01257]` at `L01397`) and `0x52` against `Y` (the constant `0x52` loaded at `L01398`, subtracted from `[L01258]` at `L01399`). Each difference is kept twice, raw and shifted right by the same count (arithmetic right shifts by the count at `L01400`/`L01401`). Four branches reject the pick and force `sdefault` (slot 17, `L01251`) at `L01250`: the raw `X` difference negative (`L01279`), the raw `Y` difference negative (`L01402`), the shifted `X` difference greater than `[view+0x84]-0x10` (`L01403`), the shifted `Y` difference greater than `[view+0x88]-0x10` (`L01280`) — matching `AI-CURSOR-191`'s "four further branches reach `sdefault`" exactly.

Only when the transformed pointer position falls inside this rectangle does the routine continue to the `view+0x140==0` test and the armed-mode jump table. A keyword search of every ledger for "party figure" and for "figure row" returns one row each, this row itself (`evidence/corpus-searches.txt`); no other row uses either phrase. A search for "view+0x80"/"view+0x84"/"view+0x88" (`go run ./tools/claim -k`, the literal `+` escaped so it is not read as a regex quantifier) returns this row, and for `+0x80` also `AI-CURSOR-191` carrying this experiment's own amendment, plus rows in `anim`, `dialogue`, `magic`, `shop` and `terrain` that predate this experiment and are not about the mission view: the offsets recur on unrelated structs, `TERR-EDGE-024`'s own text naming `CMapView+0x84` as a grid width, an object the mission view is not shown to be.

~~No row names a widget on the mission view at these offsets; the widget this rectangle describes was not identified from any other routine or from the shipped art in this experiment.~~ *-- corrected by `AI-CURSOR-232` ([EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/)): the three offsets are not a widget's. `R0218` reads `[view+0x80]+0xc` as the tile-word plane `TERR-FOG-082` names as `[CMapView+0x80]+0xc`, and `[view+0x84]` as that plane's row stride, off the same register that carries `view+0x98c` and `view+0x144`. The mission view is `CMapView`, which is the identity this row records as not shown; `view+0x88` is still unread on either side.* The corpus was not searched for `+0x68` on the containing object, so this absence is scoped to the five searches above and says nothing about that field

**Confidence.** Medium. The branch structure and its four conditions are a direct instruction-level read. What remains open is the widget's game identity, which decides whether "outside this rectangle forces `sdefault`" is the fact worth publishing or a proxy for a more specific one, and the game-level meaning of `[this+0x68]`, which is read here only as a sign test and a shift count and is called a zoom value by analogy rather than from any established writer. No live alternative for the widget's identity was tested

### AI-CURSOR-209

Both sites are cited by operand address, as `AI-CURSOR-196` cites them. At the second site (`L01327`, the operand of the load of the global `L01180` at `L01404`, reached from a hit-test result stored in the frame local at -0x60): the routine tests `[view+0x140]==1` (`L01405`) and `[view+0x144]&0x1` (`L01406`) — both established fields, `AI-PANEL-061` — into frame local `-0x104`, which is copied to frame local `-0x68` and to frame local `-0x64` at `L01407`-`L01408`; then, only when frame local `-0x64` is set, it additionally tests `[[view+0x138]+0x18c]&0x1` (the selected object's own field) and the local at -0x60 `&0x800` (the hit-test result) at `L01409`-`L01410`, and picks `town` at `L01404` when the frame local at -0x64 survives that refinement (a zero test with a jump to `L01411` at `L01412`).

At the first site (`L01326`, the operand of the load of the global `L01180` at `L01413`, reached from a separate hit-test result stored in the frame local at -0x5c), the same shape recurs: the same `[view+0x140]==1` (`L01414`) and `[view+0x144]&0x1` (`L01415`) test into frame local `-0xf8`, copied to frame local `-0x58` and to frame local `-0x54` at `L01416`-`L01417`; the same two refinement tests verbatim (`L01418`-`L01419`) against the same `+0x18c` field and the same `0x800` bit; and the pick gated on frame local `-0x54` at `L01420`. The two committed listings run `L01421`-`L01422` and `L01422`-`L01423` with no gap between them, and within that window frame local `-0x54` is written only at `L01417` and `L01419`.

The two `view+0x144` bits are already named by `AI-PANEL-061`: bit `0x1` is set at `L00692` after a `strcmp` against `"CUnit"`, which with `view+0x140==1` makes the single selected object a `CUnit`; and the refinement's own `[[view+0x138]+0x18c]&0x1` is the same field and bit that routine folds into `view+0x144` bit `0x8` at `L00697`, so the refinement re-tests a predicate the flags word already records. `AI-PANEL-061` does not establish what either predicate means in game terms. `AI-CURSOR-196`'s third reference (`L01308`, inside `R0211`) is not a selection: it is the operand of the load of the global `L01180` at `L01351`, whose next instructions load the slot's own `+0x4`, compare it against the frame local at -0x1c and branch (`L01424`-`L01353`) — one arm of the cursor-to-order dispatch `AI-CLICK-050` establishes for that routine

**Confidence.** Medium. The whole condition — `view+0x140==1`, `view+0x144` bit `0x1`, `+0x18c` bit `0x1` and hit-test bit `0x800` — is read at instruction level and recurs at both sites independently, and the local carrying it is written at only two instructions inside the committed listings. Not established: ~~the game-level meaning of `+0x18c` bit `0x1`~~ *-- `PARTY-FLAG-003` ([EXP-0077](../experiments/EXP-0077-party-and-save/)) reads `CUnit+0x18c` bit `0x1` as the player-character flag, at Medium; `AI-CURSOR-242` ([EXP-0220](../experiments/EXP-0220-pickup-cursor-route/)) finds the same field and bit gating `R0219`'s `pickup` arms* or of ~~hit-test bit `0x800` (`R0218` itself was not read)~~ *-- that is true of this experiment and not of the corpus: `UNIT-HOVER-020` ([EXP-0109](../experiments/EXP-0109-attack-order/)) reads `R0218` and gives bit `0x800` at High -- it is set for a `CStructure` whose `[L01425][obj+0x20]` record has a non-zero `+0x8c`. `AI-CURSOR-231` ([EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/)) reads the routine whole and enumerates all seven of its bits*. "What is under the pointer" when `town` is chosen is answered in terms of these two fields' own bits: `+0x18c` bit `0x1`, read as the player-character flag at Medium (`PARTY-FLAG-003`), and hit-test bit `0x800`, read as a hit `CStructure` at High (`UNIT-HOVER-020`/`AI-CURSOR-231`)

**Amended.** The two clauses the row leaves unestablished are closed elsewhere. `+0x18c` bit `0x1` is read as the player-character flag at Medium (`PARTY-FLAG-003`), and `AI-CURSOR-242` finds the same field and bit gating `R0219`'s `pickup` arms. Hit-test bit `0x800` is read at High as a hit `CStructure` (`UNIT-HOVER-020`), and `AI-CURSOR-231` reads `R0218` whole and enumerates all seven of its bits.

### AI-CURSOR-218

`R0345` has one call site in the whole image, `L01426`, inside `R0351` (`SESS-COMPOSE-057`), reached with `L01427` loading the object `L01181` and the presented rectangle pushed at `L01428`. Read to its `0x4`-byte-releasing return at `L01429`, the routine returns immediately unless `mgr+0x4` is non-zero (`L01430`, and `L01431` jumping to `L01432` when zero) and `mgr+0x20` is non-zero (`L01433`, and `L01434` jumping to `L01432` when zero); `AI-CURSOR-172` establishes `mgr+0x4` as the current cursor object and `mgr+0x20` as the manager's refcount. It then runs at most one of two earlier blocks and always ends at the cursor block beginning `L01435`. That block is: `L01436`, an indirect call through the import slot `L01437`, USER32 `IntersectRect` by the image's own import directory, against the cursor's own rectangle, returning without drawing when it is empty (`L01438` jumping to `L01432` when zero); `L01439` calls `R0346`, the back-surface lock; `L01440` calls `R0352`, which copies the pixels under the cursor into the buffer object at `mgr+0xc`; the blit at `L01330`, whose five arguments are pushed at `L01441`-`L01442` as `0`, `0`, `mgr+0x24`, `mgr+0x14 - mgr+0x1c`, `mgr+0x10 - mgr+0x18`; and `L01443` calls `R0312`, the unlock.

`AI-CURSOR-172` establishes `mgr+0x24` as the animation frame index and `mgr+0x18` and `mgr+0x1c` as two fields copied from the cursor object, so the first two arguments are a position minus a per-cursor origin. `AI-MINIMAP-156` reads `vt+0x18` on the same kind of resource object as a sub-rectangle blit. `AI-CURSOR-174` records that its experiment "did not locate the render-loop blit" and grades that half Unknown; this is the blit (`evidence/disasm-listings.txt`, `evidence/composition-sites.txt` sections A and D)

**Confidence.** High for the call, its five arguments, its enclosing lock and its single caller. Every one is a named instruction in a listing read from `R0345` to its own `RET`, and the caller count pairs a direct-call scan of `.text` with a relative-branch scan and a stored-dword scan over every section's raw extent. The alternative excluded is that the cursor is a widget in the composition tree: `R0345` has 0 stored pointers anywhere in the image, so it sits in no vtable, and its only entry is from inside the present. **Medium** for reading `mgr+0x18` and `mgr+0x1c` as a hotspot: the subtraction is read here, the meaning of the two fields is not

### AI-CURSOR-219

`R0351` runs the overlay composite at `L01426` and then does nothing but copy: the back-surface lock at `L01444`, the primary-surface lock at `L01445`, the call to `R0353` at `L01446`, and the two unlocks (`SESS-COMPOSE-057`). Inside `R0345` the three overlay blocks converge on the cursor block: the `mgr+0x60` block ends with a jump to `L01435` at `L01447`, and the `mgr+0x9c` block falls through to the same address, so the cursor is composited after whichever of the two ran. Widget content reaches the screen through this routine: every widget presents through `R0354`, which calls `R0351` at `L01448` (`TOWN-063`, and 96 vtables carry that slot, `SESS-COMPOSE-059`). That is an enumeration of the present slot, not of everything in the image that can copy to the display surface, which is why this row is scoped to regions presented through `R0351`.

The remaining ten call sites of the generic copy `R0353` lie between `L01449` and `L01450`, inside `R0355` and `R0356`. Neither routine references `[L01451]` and neither is among the nine `R0357` call sites: both lock the third surface at `[L01452]` through `R0358` (`SESS-COMPOSE-056`, `SESS-COMPOSE-058`)

**Confidence.** High for regions presented through `R0351`. The alternative excluded is that something writes to the display surface after the copy inside this routine: `R0351` past `L01446` does nothing but unlock at `L01453`, restore `[L01168]` at `L01454`, unlock again at `L01455` and call `R0359` at `L01456`, and that is one listing read to the `RET`. **This row says nothing about a region that never enters `R0351`**, and it does not assert that no widget can ever write to the display surface: `SESS-COMPOSE-058` grades that Medium and names what it does not exclude. `R0360`, which receives the display surface object at `L01457`, was not read

### AI-CURSOR-220

Every widget's content reaches the screen through `R0354`, which calls `R0351` (`TOWN-063`, `SESS-COMPOSE-059`), and `R0351` copies the back surface at `[L01458]`. Whatever a widget painted was therefore in the back surface when the present ran, and the cursor is composited into the back surface immediately before that copy (`AI-CURSOR-219`). The argument is about the present's own order and does not need a claim that no widget can ever reach the display surface, which `SESS-COMPOSE-058` grades Medium. The worked instance is the minimap: `AI-MINIMAP-156` reads `R0310`, the paint routine of `campaign+0xd8`'s own widget, from entry to its `RET` at `L01160`, and finds it writing through `[L01168]` and closing with `R0312`; the back-surface lock site inside that routine is `L01459`.

Ten of the 85 back-surface lock sites fall between `L01460` and `L01461`: `L01462`, `L01463`, `L01464`, `L01465`, `L01466`, `L01459`, `L01467`, `L01468`, `L01469`, `L01470`. None of the nine primary-surface lock sites is in that range. `MENU-COMBAT-017` names the right column's children as `campaign+0xd8`, `campaign+0xdc`, `campaign+0xe0` and `campaign+0xe4` (`evidence/composition-sites.txt` section A)

**Confidence.** High for the uniformity, which rests on the present's own order rather than on reading each widget: whatever reaches the screen through `R0351` was in the back surface before the copy, and the cursor is the last write into the back surface (`AI-CURSOR-219`). The alternative excluded is that one of these surfaces is presented by a different routine that composites no cursor: every widget presents through `R0354`, which has one call to `R0351` at `L01448` and no other present. **Medium** for the attribution of individual back-surface lock sites to individual widgets: only `L01459` is attributed, and that through a published row which reads its containing routine end to end. The other nine were not attributed to a routine. **Medium** for the residual the excluded alternative does not cover, that one of these five writes to the display surface outside its present slot: that is excluded at instruction level only for the minimap, whose paint routine `AI-MINIMAP-156` reads end to end, and for the other four it rests on the nine primary-surface lock sites all falling outside `L01460`..`L01461` while ten back-surface lock sites fall inside it, which is a lock-site range observation and not a read of those routines

### AI-CURSOR-221

`MENU-ESC-010` names the two constructions and the shared `L01471` … call to `R0361`. Re-read here: the call to `R0362` at `L01472` is followed by a jump to `L01471` at `L01473`, the call to `R0363` at `L01474` by a jump to `L01471` at `L01475`, and at `L01471` both push the constructed object and reach the call to `R0361` at `L01476`. The rectangles pushed at `L01477`-`L01478` and `L01479`-`L01480` are `(100,60)-(440,400)` and `(100,100)-(440,340)`, which agree with `MENU-ESC-010`. Inside `R0361`, after the screen darkening `DLG-DIM-013` reads, the order is the call to `R0312` at `L01481` (back-surface unlock), the call to `R0364` at `L01482`, and at `L01483` the call through the object's vtable slot at `+0x34`, the panel's own draw-subtree slot. `R0364` builds the whole client rectangle through `CopyRect` and calls `R0351` at `L01484`, so the darkened screen becomes visible through the present, which composites the cursor last.

The panel's own draw at `L01483` runs after that present, so its content becomes visible through a later one. Every present in the widget path is `R0351`: 96 vtables carry `R0354`, which has one call to it (`SESS-COMPOSE-059`, `TOWN-063`), and this experiment found no second present routine. Three vtables in the image replace the draw-subtree slot (`SESS-COMPOSE-059`); one of them, base `L01485`, replaces the slot at `L01486` with `R0365`, whose whole body is the calls to `R0346` (`L01487`), `R0366` (`L01488`) and `R0312` (`L01489`) - the base draw bracketed by a back-surface lock and unlock (`SESS-COMPOSE-059`). `R0361` also calls `R0367` at `L01490`, gated on `[L00210]`; that routine reaches `R0368` at `L01491` and `R0323` at `L01187`, both cursor-manager routines, and neither changes the composition order

**Confidence.** High for the chain from either construction site to `R0351`, and for both rectangles: every step is a named instruction, and the two rectangles are immediates re-read here and already published by `MENU-ESC-010`. Both panels were measured, not one. The alternative excluded is that the panel's own content is composited into a frame after the cursor: the vtable call at `+0x34` (`L01483`) is the draw-subtree slot and it runs after the call to `R0364` at `L01482` has returned, so it cannot execute inside the present that call performed, and the present is the routine that composites the cursor. **Medium** for the step from that to "below the cursor on the screen", which needs the panel's content to reach the screen through `R0351` as well: `R0354` has one call to it and this experiment found no second present routine, but the search was a stored-dword scan for `R0354` and a direct-call scan for `R0351`, not an enumeration of everything that can copy to the display surface. **Unknown** whether vtable `L01485` is either Esc panel's own class: the constructors were not followed to a vtable store, and the row does not depend on it

### AI-CURSOR-222

`R0324` — the routine `AI-CURSOR-172` reads for its `timeGetTime` frame advance and does not read past it — continues at the call to `R0357` at `L01492`, the primary-surface lock, saves the pixels under the cursor into the buffer object at `mgr+0x8` (the `mgr+0x8` read at `L01493`, the call to `R0352` at `L01494`), issues the vtable call at `+0x18` (`L01495`) on `mgr+0x4` with the same five arguments as `AI-CURSOR-218`'s blit (`0`, `0`, `mgr+0x24`, `mgr+0x14 - mgr+0x1c`, `mgr+0x10 - mgr+0x18`), and unlocks at `L01496`. Its counterpart is `R0325`: the call to `R0357` at `L01497`, a restore from the `mgr+0x8` buffer, the call to `R0369` at `L01498`. `R0325` has 0 direct call sites and 0 stored pointers; its only reference in the image is the tail jump at `L01196` to `R0325` out of `R0322`'s decrement-to-zero arm, which is the "distinct, untraced routine" `AI-CURSOR-172` names.

So the manager's two sub-objects have one job each: `mgr+0x8` backs the primary surface and `mgr+0xc` backs the back surface, and which surface the one blit primitive writes to is decided by the enclosing lock (`SESS-COMPOSE-056`). `R0323`, the only caller of `R0324`, has 12 call sites, all listed by `AI-CURSOR-172`; one is `L01185` inside `R0370`, which itself has 260 call sites

**Confidence.** High for the second blit, its surface, its backing store and the two routines' reference sets: each is a named instruction in a listing read to its own `RET`, and the reference sets pair a direct-call scan with a relative-branch scan and a stored-dword scan. The relative-branch scan is what finds `R0325` at all, and it agrees with `AI-CURSOR-172`, which reached the same tail jump independently. **Medium** for this being a live per-motion path in play: the cadence of the 12 call sites was not traced here, which is the same gap `AI-CURSOR-172` records for its own reading of the same 12

### AI-CURSOR-223

Three kinds of gate were found, all inside the cursor manager. First, both overlay routines return without touching anything unless `mgr+0x4` and `mgr+0x20` are both non-zero: `L01430`/`L01433` in `R0345` and `L01499`/`L01500` in `R0359`, the same pair in the same order. Second, each overlay element is tested with `IntersectRect` against the rectangle being presented — `L01501` for the `mgr+0x60` element, `L01502` for the `mgr+0x9c` element, `L01436` for the cursor — and an empty result skips that element only. Third, the `mgr+0x60` block ends with a jump to `L01435` at `L01447`, so when it runs the `mgr+0x9c` block does not. Within a present the surviving elements are always in the order `mgr+0x60` or `mgr+0x9c`, then the cursor; no branch reorders them and no phase word selects a different routine. `R0345` and `R0359` have one call site each in the whole image

**Confidence.** High for the gates and for the fixed order, which are named instructions in two listings read to their own terminators, plus a call-site enumeration. The alternative excluded is a second compositor selected by state: `R0345` and `R0359` are the only routines the present calls for this purpose and each has exactly one call site, so there is no second one to select. **Unknown** what `mgr+0x60` and `mgr+0x9c` are. The first block does one save-under and one call to `R0371`; the second does four save-unders over four strips of one rectangle, which is the shape of an outline. Neither was identified

### AI-CURSOR-224

`AI-CURSOR-172` describes `R0324`'s clock read as "WINMM.dll `timeGetTime` (IAT `L00849`, the sole caller of that import in `.text`)" and its confidence paragraph repeats it as "`timeGetTime`'s one call site is inside this exact chain". A scan of the whole image for the stored dword `L00849` returns 108 hits, every one in `.text`: 85 of the form of the 6-byte indirect call operand, a direct call through the slot; 22 that load the resolved address into a register for a later call (11, 5, 4 and 2 into the four distinct destination registers EDI, ESI, EBX and EBP respectively); and one jump through `L00849` at `L01198`, an import thunk with 0 direct call sites. `L01197` is one of the 85. The row's conclusion does not rest on the clause: what excludes a Windows-timer-driven animation is `SESS-TIMER-022`'s import-directory walk, which finds `SetTimer`, `KillTimer`, `timeSetEvent` and `timeKillEvent` absent from the image, and that is unchanged. Only the uniqueness of the call site is withdrawn (`evidence/disasm-listings.txt`, final section)

**Confidence.** High. The count is a complete stored-dword sweep over the raw extent of every section against one fixed address, and the two forms are read out of the printed bytes rather than inferred. The alternative excluded is that the other 107 sites reference a different address that happens to print the same: the scan matches the four bytes of the address itself and prints six bytes of context on each side, and `FF 15` before it is unambiguous

### AI-CURSOR-225

`TOWN-209` reads the tip popup's `vt+0x30` and describes the pair as "`R0346()` (no arguments)" and "`R0312()` (no arguments, clip-restore-shaped)". `DLG-DIM-013`, published earlier, already names the same two routines as the DirectDraw Lock and Unlock. The correction is scoped to this one pair: `TOWN-209` also names `R0372` and `R0373` as a clip-save/clip-set pair, and that reading is correct. `R0373` copies a `RECT` from its argument into the four consecutive globals `[L01503]`, `[L01504]`, `[L01505]` and `[L01506]` and returns; `R0372` passes `L01503` to `CopyRect` at `L01507` (an indirect call through `L01508`), copying the same four out. `DLG-DIM-013` cites `R0346` reaching `vt+0x64` with the store of `0x6c` into the dword at `L01509` at `L01510` and `R0312` reaching `vt+0x80`.

This experiment reads them the same way and adds the population: `R0346` has 85 call sites and `R0312` 88, and `SESS-COMPOSE-056` shows the pair publishing the locked surface's pixel pointer and pitch into `[L01168]` and `[L01511]`. The correction does not change what `TOWN-209` establishes about the popup: the tile ids, the loops and the pitch constants stand, and the bracket it describes is the lock that makes the tile blits legal rather than a clipping state

**Confidence.** High. The alternative excluded is `TOWN-209`'s own reading, that `R0346` and `R0312` are the clip-save and clip-restore pair. That pair is `R0373` and `R0372`, which copy a `RECT` into `[L01503]`..`[L01506]` and back out through `CopyRect` at `L01507` and touch no surface. `R0346` instead stores the descriptor size `0x6c` into the dword at `L01509` at `L01510` and calls the surface object's `vt+0x64`, after which `[L01168]` holds a pixel pointer and `[L01511]` a pitch, both read by the caller as named instructions. The two functions are therefore distinct routines with distinct globals, not two readings of one. That two published readings of the lock pair now agree is corroboration and is not what the grade rests on

## Hover cursor cascades and the pickup route

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-CURSOR-226 | `R0219` contains two mask-to-cursor cascades, not one, and an ordinary hover runs the second: over a hostile unit with no modifier key held it selects `attack`. | High | ● active | [EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/) |
| AI-CURSOR-227 | Two clauses of `AI-CURSOR-052` are refuted. The hostility bit `0x4` is tested alone at two addresses, and at one of them it selects `attack` with no modifier key held; and the Alt latch does not force `move` "regardless". | High | ● active | [EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/) |
| AI-CURSOR-228 | The selection local (frame offset -0x28) of `R0219` has 48 writers, and on entry to the modifier cascade at `L00633` it holds `cast` or `move`, never the zero it was initialised to. | High | ● active | [EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/) |
| AI-CURSOR-229 | `R0219`'s first cascade has four gates, a `cast` default and three replacements, two of them driven by three parallel 24-entry boolean tables. Its own default and one of its gates read as a cast-target test, at Medium. | High / Medium / Unknown | ● active | [EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/) |
| AI-CURSOR-230 | Five replacements run after the ordinary cascade of `R0219` in a fixed order and four of the five after the first cascade, and the last two are this routine's own use of the held-item cursor that `AI-CURSOR-192` left unread. | High | ● active | [EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/) |
| AI-CURSOR-231 | The hover hit test's capability mask has exactly seven bits. Bit `0x400` is set when the four corners of the hovered cell do not OR to `0xc000` and bit `0x40` when they do and the cell carries a drawable on the `CMapView+0x98` plane.… | High / Medium / Unknown | ● active | [EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/) |
| AI-CURSOR-232 | The mission view of the `AI-CURSOR` and `AI-PANEL` rows and the `CMapView` of the `TERR` rows are one object. `AI-CURSOR-208`'s unidentified widget rectangle is read from the map's own grid fields. | High / Medium | ● active | [EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/) |
| AI-CURSOR-233 | `R0219` has one `RET` and exactly four exits that call no cursor adapter, so `AI-CURSOR-193`'s exit enumeration is now complete for this routine; and its entry gate is an equality test against a field `TOWN-352` reads as a bitmask. | High / Medium | ● active (amended, partially retracted) | [EXP-0218](../experiments/EXP-0218-hover-cursor-cascades/), [EXP-0219](../experiments/EXP-0219-mapview-panel-messages/) |
| AI-CURSOR-234 | The map view's dispatcher arms for `0x40e` and `0x40f` each call one of two AddChild/RemoveChild toggle routines, gated on the same child-existence predicates `TOWN-347`/`TOWN-351` already read for the widget's own paint logic… | High / Medium / Unknown | ● active (amended, partially retracted) | [EXP-0219](../experiments/EXP-0219-mapview-panel-messages/), **[EXP-0230](../experiments/EXP-0230-non-music-sfx-events/)** |
| AI-CURSOR-235 | `R0374` is a distinct routine from the map view's private dispatcher, not part of it, and it writes `this+0xdc=1` and `this+0xe0=1` unconditionally at its own entry… | High / Medium / Unknown | ● active | [EXP-0219](../experiments/EXP-0219-mapview-panel-messages/) |
| AI-CURSOR-236 | No reader of the map view's own `+0xe0` field was found. Two writers are now known -- `TOWN-347`'s rect C (`L01512`) and `R0374`'s unconditional entry write… | Unknown | ● active | [EXP-0219](../experiments/EXP-0219-mapview-panel-messages/) |
| AI-CURSOR-242 | Both `pickup` loads in `R0219`'s ordinary-hover cascade are gated by the identical stack slot (frame offset -0x68), and that slot's own construction reads no dedicated "sack" bit… | High / Medium / Unknown | ● active | [EXP-0220](../experiments/EXP-0220-pickup-cursor-route/) |
| AI-CURSOR-243 | `AI-CURSOR-226`'s `town` load at `L01404`, in the same `mask & 0x23` arm as the first `pickup` site, is one of `AI-CURSOR-209`'s two cited `town` sites; the other is a different address, in the routine's other cascade. | High | ● active | [EXP-0220](../experiments/EXP-0220-pickup-cursor-route/) |

### AI-CURSOR-226

The cascade `AI-CURSOR-052` reads at `L00633` is reachable only through `L01513`, which the jump-on-zero at `L01514` enters only when `view+0x99c` is zero, and which then passes four gates. Each gate is a conditional branch to `L01515` taken when its own condition FAILS, so the four listed here are the conditions a hover must SATISFY to stay on this path: `R0317(view)` non-zero (the zero test at `L01516` and the long jump-on-zero at `L01517`), the id `[[sess+0xec]+0x74]` when non-negative else `[[sess+0xec]+0x60]` non-negative (the comparison at `L01518` and the long jump-if-less at `L01519`), `view+0x144 & 0x200` set (the mask at `L01520` and the long jump-on-zero at `L01521`), and `view+0x144 & 0x4` clear (the mask at `L01522` and the long jump-on-non-zero at `L01523`).

An ordinary hover has `view+0x144 & 0x200` clear, so it fails gate 3 and leaves for `L01515` there at the latest. `TOWN-354` reads `R0317(v)` as `R0375(sess+0xf0,3) != 0` when `sess+0x3dc & 2` and `R0375(v,3) != 0` otherwise, and `AI-CURSOR-233`'s entry gate requires `sess+0x3dc == 1`, so on this routine's path gate 1 is the presence of the view's child with id 3. Every other hover runs the cascade at `L01515`, which recomputes the mask into the frame local at -0x60: `view+0x140 == 0` (`L01524`) gives `select` on `mask & 0x23` and `default` otherwise; `view+0x144 & 0x24` set (`L01525`, `AI-CURSOR-203`'s mask on `AI-CURSOR-202`'s field) gives the same pair; then `mask & 0x4` (`L00635`) gives `move` with the Alt latch set, `select` on `mask & 0x20` (`L01526`), and otherwise **`attack`**, slot 3, `L00628`, loaded at `L01527` and stored at `L01528`; then `mask & 0x23` (`L01529`) gives `move` with Alt, `attack` with Ctrl (`L01530`), `pickup` (`L01531`), `town` (`L01404`) or `select` (`L01411`); and empty ground (`L01423`) gives `swarm` with Ctrl (`L01532`), `move` with Alt, `pickup`, else `move` (`L01533`).

Each address named for a cursor in this enumeration is that slot's own load; the store into the cursor local (frame offset -0x28) is the next instruction, as with `L01527`/`L01528`. A hostile `CStructure` therefore gives `select`, not `attack`, because every hit `CStructure` carries bit `0x20` and `L01526` tests `0x20`. `UNIT-HOVER-020` states `0x20` and `0x800` as alternatives; they are not alternatives, and `AI-CURSOR-231` carries the correction. **The Shift latch `[L00670]` is read at exactly one address in the whole routine, `L01534`, inside the first cascade**, so Shift changes nothing on an ordinary hover. Slot names are `SPR16A-CURSOR-067`'s, latch identities `AI-KEYMOD-059`'s, mask bits `UNIT-HOVER-020`'s and `AI-CURSOR-231`'s

**Confidence.** High. The live alternative excluded is the reading the corpus currently supports: that `AI-CURSOR-052`'s cascade is what an ordinary hover runs. `evidence/control-flow.txt`, a mechanical extraction of every explicit branch edge of the routine from the committed listing, gives `L01513` exactly one predecessor and shows no branch from outside landing anywhere in `L01513`-`L00633`; the routine's one indirect branch is the jump table at `L01535`, whose eight targets were read out of the image and all lie in `L01536`-`L01537`, so the edge set for that region is complete rather than the explicit-branch subset. The `attack` result is stated with its guards: it holds when at least one object is selected, `view+0x144 & 0x24` is clear, no item is held, no marquee is active and the pointer is outside both child rectangles, which are the five replacements `AI-CURSOR-230` enumerates

### AI-CURSOR-227

`AI-CURSOR-052` states that "the hostility bit `0x4` appears only inside `mask & 0x24`, whose effect is to suppress the select cursor". Its reading of `L00636`-`L00637` is correct. The word "only" is not: `L00634` (a mask with `0x4`) tests the bit alone in the first cascade's tail, and `L00635` (a mask with `0x4`) tests it alone at the head of the second cascade, where a set bit with the Alt latch clear and `mask & 0x20` clear selects `attack` at `L01527`/`L01528` -- the opposite of suppression, and with no modifier key held. `AI-CURSOR-052` also states that "with `[L00632]` set the cursor is **move** regardless". Its store at `L01538` is correct and "regardless" is not: on that same path the local is replaced by `pickup` at `L01539` when the composite at frame offset -0x58 holds, by `town` at `L01540`, by `default` at `L01541` or `L01542`, by `[sess+0x3d8]` at `L01543` and by `backpack` at `L01544` -- six stores, the same six the `AI-CURSOR-052` entry in `claims/retracted.md` lists.

The marquee's own `default` store at `L01545` is not among them: this path ends with the 5-byte relative jump to `L01268` at `L01546`, which lands past it (`AI-CURSOR-230`); and the second cascade's first arm (`L01524`-`L01547`, taken when `view+0x140 == 0`) contains no read of `[L00632]` at all, so with the Alt latch set and nothing selected a hover over a unit gives `select`. Both refuted clauses were published at High. `AI-CURSOR-052`'s other clauses -- the hover-time hostility test, the `mask & 3` split under Ctrl, and the cursor-to-art binding -- are unaffected

**Confidence.** High. The live alternative excluded for the first clause is that the two `AND` sites are the low half of a wider `mask & 0x24` test written in two instructions: both encodings carry the immediate `0x04`, `evidence/disasm-listings.txt` prints their bytes, and neither is adjacent to a second `AND` on the same value. For the second clause the exclusion is a path rather than an encoding: the `JNZ` at `L01548` is the only edge past the `view+0x140 == 0` arm, and that arm's twelve instructions, `L01524` to `L01547` inclusive, are in the same listing

### AI-CURSOR-228

The complete writer set, each with the instruction that produced its value, is listed in `evidence/control-flow.txt`, derived mechanically from the committed listing: one immediate zero at `L01301`, 46 stores of a value loaded from one of `SPR16A-CURSOR-067`'s 28 registry slots, and one store of `[sess+0x3d8]` at `L01543`. Eight of the 48 lie in `L01549`-`L01550`, the range absent from `EXP-0216`'s committed listing. The same listing carries **49** references to a cursor-registry slot inside the routine, and they account for the writer set exactly: 46 are a `MOV` whose destination register the very next instruction stores into frame local `-0x28`, and three are a `CMP` of the local's current value against a slot -- `L01551` and `L01552` against `cast` `L01553`, `L01554` against `move` `L01555`. 46 slot loads plus the zero at `L01301` plus the `[sess+0x3d8]` store at `L01543` is the 48.

`L01556` (a store of the value loaded from the global `L01553` into frame local -0x28) dominates `L00633`: that cascade is entered only from the zero-branch at `L01557` or the fall-through from `L01419`, both inside `L01421`-`L01419`, whose only entries are the non-zero branch at `L01558`, the zero branch at `L01559` and the fall-through from `L01560`, all past the unconditional store at `L01556`. So the local holds `cast` there, or `move` where `L01561` (`mask & 0x400`) or `L01560` (`mask & 3` clear and `R0094([sess+0xec])` non-zero) replaced it. **Frame local `-0x8`, written at `L01562` as the bitwise OR of the Ctrl latch, the Alt latch and their own OR, has no reader in the routine**

**Confidence.** High. Two alternatives are excluded. A writer reaching the slot through a pointer would need its address, and the routine contains no `LEA` of frame local `-0x28`, so the 48 stores are the population rather than a sample of one addressing form. A path on which the `L01301` zero survives to `L00633` would need an entry into `L01513`-`L00633` that skips `L01556`; the extracted predecessor set contains no such edge, and the routine's only indirect branch targets eight addresses all below `L01563`

### AI-CURSOR-229

Past the four gates of `AI-CURSOR-226`, `L01556` writes `cast` (slot 7, `L01553`). `L01561` replaces it with `move` when `mask & 0x400`. `L01560` replaces it with `move` when `mask & 3` is clear and `R0094([sess+0xec])` is non-zero. After the modifier cascade, `L01564`-`L01565` replaces it with `move` when the local is still `cast`, `R0094([sess+0xec])` is non-zero, `mask & 0x4` is clear (`L00634`), `[[sess+0xec]+0x74]` is negative (`L01566` `CMP`/`SETGE`/`TEST`/`JNZ`) and `[L01567 + [[sess+0xec]+0x60]*4]` is non-zero (`L01568`). `R0094` is eight instructions with two returns: `[this+0x74]` non-negative returns `[L00201 + [this+0x74]*4]`, otherwise `[L00200 + [this+0x60]*4]`.

The three tables are 24 dwords each, adjacent at `L00200`, `L01567` and `L00201`, and all 72 values are 0 or 1 (`evidence/imports-and-tables.txt`). The index recomputation at `L01569`-`L01570` selects between `+0x74` and `+0x60` a second time although `L01571` has already established `+0x74` negative, so the index at `L01568` is always `+0x60`

**Confidence.** High for the gates, the branch structure, `R0094`'s body and the tables' shape and contents: all are instructions in the committed listing or the dword values read out of the image, and the alternative that some gate is off the path is excluded by the extracted predecessor set. **Medium** for reading the branch as a cast-target test. That rests on two facts and not on any decoded target predicate: the branch's own default is the `cast` slot, and one of its gates is `view+0x144 & 0x200`, which `AI-PANEL-061` and `AI-PANEL-053`'s retraction entry name as the bit that enables the cast mode. **Unknown** for what `[sess+0xec]` is and what the three tables hold: `go run ./tools/claim -k` returns no row for `L00200`, `L01567`, `L00201`, `campaign+0xec` or `sess+0xec` (`evidence/corpus-searches.txt`), `R0317` is read by `TOWN-354` at High -- it returns `R0375(sess+0xf0,3) != 0` when `sess+0x3dc & 2` and `R0375(v,3) != 0` otherwise, so gate 1 is the presence of a child with id 3, through the same accessor and the same index `AI-CURSOR-230` finds in the tail. `go run ./tools/claim -k "R0317"` returns seven rows, and this experiment ran that search (`evidence/corpus-searches.txt`) without reading its output

### AI-CURSOR-230

In order: (1) `L01572`-`L01545`, the marquee, which runs on the ordinary-cascade path only: the first cascade's tail ends with the 5-byte relative jump at `L01546` to `L01268`, which lands past it, and every edge into `L01572` comes from inside the ordinary cascade. Replacements (2) to (5) run on both paths. The marquee test is -- `[L00210]` non-zero (`AI-INPUT-121`'s flag) and `IsRectEmpty` false on the normalised copy of `[L01573]` gives `default`, and the identical test is repeated on the armed-mode path at `L01574`-`L01575`; (2) `R0375(this, 2)` non-null and `PtInRect(child+0x8)` gives `default`, loaded at `L01576` and stored at `L01541`; (3) the same with index 3, `L01577`/`L01542`; (4) `[sess+0x3cc]` non-zero gives `[sess+0x3d8]` at `L01543`, the held-item cursor of `AI-CURSOR-192`, reached through the same test-`+0x3cc`-then-load-`+0x3d8` idiom that row finds in `R0319` and `R0336`; (5) inside (4), `sess+0x3dc == 1`, `view+0x140 == 1`, `[[view+0x138]+0x18c] & 1`, `[[view+0x138]+0x14] == view+0x9b4`, the pointer's map column and row within 2 of the selected object's `+0x34` and `+0x38`, and `mask & 1` give `backpack` (slot 27, `L01212`) at `L01544`.

`R0376`, called twice for the two differences, is five instructions (`R0376`-`L01578`) and is `abs`. The three imports behind the indirect calls are `PtInRect` `[L01178]`, `IsRectEmpty` `[L01579]` and `CopyRect` `[L01508]`, resolved from the PE import directory rather than by name guess. This makes `R0219` the eighth of `AI-CURSOR-192`'s thirteen `+0x3d8` routines to be read; the five not read are `L01288`, `R0341`, `R0343`, `L01289` and `L01290`

**Confidence.** High for the order and the five replacements. The live alternative excluded is that they are arms of one chain rather than a sequence: each is a separate `CMP`/`JZ` whose taken and not-taken edges both converge on the next test, which the extracted predecessor set shows directly, so a later replacement overwrites an earlier one rather than being skipped by it. What is **not** established: what `R0375`'s children 2 and 3 are on the mission view. `TOWN-347` reads the same accessor with the same two indices on the town screen, and whether those are the same widgets was not tested

### AI-CURSOR-231

The hover hit test's capability mask has exactly seven bits. Bit `0x400` is set when the four corners of the hovered cell do not OR to `0xc000` and bit `0x40` when they do and the cell carries a drawable on the `CMapView+0x98` plane. `TERR-FOG-082`'s state mapping reads the first as no corner currently visible, at Medium. This closes the open item `UNIT-HOVER-020` names. `R0218`'s mask local frame local `-0x50` is written at exactly eight addresses in the whole routine: `L01580` zeroes it, then `L01581` bit `0x1`, `L01582` bit `0x2`, `L01583` bit `0x800`, `L01584` bit `0x20`, **and those two are not alternatives**: `L01585 74 0b JZ L01586` sends a zero `+0x8c` past the `0x800` store, and `L01587 eb 0e JMP L01588` lands the non-zero case ON the `0x20` store rather than past it, so a hit `CStructure` always carries `0x20` and carries `0x800` in addition when `+0x8c` is non-zero; the third arm, `+0x8c` zero with `frame local -0x80+0x64` non-zero, jumps to `L01589` and sets neither bit.

`UNIT-HOVER-020` reads these as `0x800` when `+0x8c` is non-zero and `0x20` **otherwise**, at High, and that clause is refuted. Continuing, `L01590` bit `0x4`, `L01591` bit `0x40`, `L01592` bit `0x400`. The last two: `L01593`-`L01594` builds the flat tile index `(X>>5 + view+0x5c) + (R0377(X,Y) + view+0x60) * [view+0x84]` into the tile-word plane `[[view+0x80]+0xc]` and ORs the `0xc000` field of the four words at index, index+1, index+1+stride and index+stride -- the same four-corner OR, the same mask and the same comparison with `0xc000` that `TERR-FOG-086` reads in `R0378`. the comparison with `0xc000` at `L01595` with a not-equal branch skips the `0x40` arm and the comparison with `0xc000` at `L01596` with an equal branch skips the `0x400` arm, so the two bits are the same comparison read both ways and cannot both be set.

The `0x40` arm additionally requires `[[view+0x98] + ((r+3)*([view+0x64]+6) + c+3)*4]` non-zero at `L01597`, where `c = X>>5` and `r = R0377(X,Y)` are the RAW view-relative pair and not the map column and row of the sentence before it: `view+0x5c` and `view+0x60` are added into frame local `-0x24` and frame local `-0x2c` only at `L01598` and `L01599`, after this index is built at `L01600`-`L01601`, while the tile index adds them into a separate temporary frame local `-0x38` at `L01602` and `L01603`. That is the same index form `TERR-SPR-137` gives for its five drawable planes (`CMapView+0x8c/+0x90/+0x94/+0x98/+0x9c`, all indexed `(row+3)*(visCols+6) + (col+3)`); `TERR-SPR-138` reads `view+0x64` as the visible-column count, which is what makes the two index forms the same expression.

The routine also writes `view+0x994` = `X>>5 + view+0x5c` and `view+0x998` = `R0377(X,Y) + view+0x60` at `L01604` and `L01605`, the hovered cell's map column and row, beside the `view+0x98c` and `view+0x990` that `UNIT-HOVER-020` names. `-fnof` gives the routine six direct `E8` callers and all six are inside `R0219`, at `L01299`, `L01606`, `L01607`, `L01608`, `L01609` and `L01610`, listed with every other call of that routine in `evidence/control-flow.txt`. The returned mask is unused at two of the six: `L01299` stores nothing and the jump at `L01611` leaves for the routine's exit, and `L01612` stores it into frame local `-0x5c`, which has no reader at or after that address. That scan finds direct relative calls only, so it says nothing about a call reached through a vtable slot

**Confidence.** High for the seven-bit enumeration: the routine was decoded whole from its prologue to its own return, the local is a stack slot in a fixed frame-pointer frame, and there is no address-of of it anywhere in the routine, which excludes a writer reaching it through a pointer. High for the `0x20`/`0x800` flow that refutes `UNIT-HOVER-020`: the live alternative excluded is that `L01587` jumps PAST the `0x20` block, and the displacement decides it -- `eb 0e` taken from `L01586` is `L01588`, the first instruction of that block, and `evidence/disasm-listings.txt` prints both addresses with their bytes. High for the two comparisons and for the mutual exclusion, which follows from both arms testing one accumulator against one constant with opposite branch senses. **Medium** for the game reading of bit `0x400` as "under fog": it converts the OR test into "at least one corner currently visible" using `TERR-FOG-082`'s state mapping, and inherits that row's own Medium blind spot on the unreachability of the `01` state -- were `01` reachable, a corner holding `0x4000` and another holding `0x8000` would OR to `0xc000` with no corner visible. **Unknown** which class of drawable occupies `CMapView+0x98`: `TERR-SPR-137` names the five planes and their index form and does not say what each holds

### AI-CURSOR-232

`R0218` receives the mission view as `this` and stores into `view+0x98c`, `view+0x990`, `view+0x994` and `view+0x998` (`UNIT-HOVER-020` for the first two, `AI-CURSOR-231` for the other two). `view+0x144` (`AI-PANEL-061`) is read by `R0219` off the same object and not by this routine: the token `0x144` does not appear anywhere in `R0218`'s committed listing. Off the same register it reads `[this+0x80]+0xc` as a word plane masked with `0xc000`, which is `TERR-FOG-082`'s `[CMapView+0x80]+0xc` tile plane; `[this+0x84]` as that plane's row stride (a multiplication by it at `L01613`); and `[this+0x98]` as one of `TERR-SPR-137`'s five drawable planes at that row's index form `(row+3)*(visCols+6) + (col+3)` with `visCols = [this+0x64]`, which is `TERR-SPR-138`'s inner loop bound.

`AI-CURSOR-208` reads `[[view+0x80]+0x4]`, `[[view+0x80]+0x8]`, `[view+0x84]-0x10` and `[view+0x88]-0x10` in `R0319`, where `view` is `[this+0x5c]` (`AI-CURSOR-191`), and records that no row names a widget at those offsets and that `TERR-EDGE-024`'s `CMapView+0x84` belongs to "an object the mission view is not shown to be". It is that object: `view+0x80` is the map, `[view+0x80]+0x4` and `+0x8` its dimensions, and `view+0x84` the tile-grid row stride

**Confidence.** High for the object identity. The live alternative excluded is that the mission view and `CMapView` are different structures whose fields happen to coincide: one routine reads or writes `+0x98c`, `+0x990`, `+0x80`, `+0x84`, `+0x64`, `+0x5c` and `+0x60` off one register in one frame, and the first two of those are mission-view fields by `UNIT-HOVER-020` (EXP-0109) while the rest are `CMapView` fields by `TERR-FOG-082` (EXP-0088), `TERR-SPR-137` and `TERR-SPR-138`. The exclusion is not a coincidence argument: the two vocabularies are reached off the SAME register, frame local `-0x110`, within one stack frame, so one pointer serves both. **Medium** for what `AI-CURSOR-208`'s rectangle then measures. `view+0x88` is not read in either listing here, so it is the map's row count by position in the pair rather than by a read, and the screen role of the rectangle -- what widget draws the map at origin `(0x48, 0x52)` scaled by `[this+0x68]` -- was not established

### AI-CURSOR-233

The routine spans `R0219`-`L01614` with a single return. The extracted predecessor set gives `L01298`, the epilogue, exactly four explicit predecessors -- the jump at `L01615` (the entry gate), the jump at `L01611` (the failed `PtInRect` at `L01616`), the zero branch at `L01617` (the selection local still zero) and the zero branch at `L01618` (the local's `+0x4` already equal to `[L00619]` as sampled at `L01303`) -- plus the fall-through from the call to `R0320` at `L01218`, which is the path that does call the adapter. Those are exactly the four `AI-CURSOR-193` names for this routine, and the enumeration is complete because the routine has one exit instruction and its predecessor set was extracted mechanically.

**The routine sets the cursor itself**: `L01218` is the set-cursor adapter (`AI-CURSOR-172`), and no caller is needed to turn the routine's result into a drawn cursor. `-fnof` gives four direct callers -- `L01619`, which it attributes to `L01620`, an address inside `R0379`, whose span `TERR-SPR-137` gives as `R0379`-`L01621`; that attribution is the false-boundary artefact `AI-CURSOR-205` describes, and the call site belongs to the map-view frame routine; and `L01622`, `L01623` and `L01624`, attributed by the same `-fnof` query to `R0374`'s nearest called start, which also contains every message-handler address `AI-CURSOR-177` lists for the map view's private dispatcher.

**Amended by `EXP-0219` (`AI-CURSOR-235`): `L01622` and `L01623` are not inside `R0374`.** `R0374` is a separate routine, `R0374`-`L01625`, with its own frame setup and a return at `L01625`; the map view's private dispatcher begins at the next instruction, `R0333`, with its own frame setup and its own return at `L01626`. `L01622` and `L01623` are inside that second routine, past `R0374`'s return. `R0374` has exactly one direct `rel32` caller (`L01627`); the dispatcher has none and is reached only through an `.rdata` vtable slot (`L01628`), so `-fnof`'s nearest-directly-called-start heuristic has no entry for the dispatcher and reports `R0374` instead, across the clean return boundary between them.

`L01624` is also not inside `R0374`: it is past the same return, inside a third routine, `R0239`-`L01629`, that `EXP-0219` (`AI-CURSOR-235`) locates but does not read. None of the four call sites was read here, so what decides *when* the routine runs is not established. That scan finds direct relative calls only. The entry gate is the comparison of `[sess+0x3dc]` with `0x1` at `L01297`, branching to the exit when equal; `sess` is the main-frame object (`SESS-OBJ-001`) and `TOWN-352` establishes `campaign+0x3dc` as a bitmask whose bit 1 is the shop, set without clearing bit 0 when a shop is opened from a mission. An equality test against 1 therefore makes the whole routine inert whenever any bit above bit 0 is set

**Confidence.** High for the exit set. The live alternative excluded is a fifth no-adapter exit reached by a path the explicit-edge extraction cannot see: the routine has one return, so every exit passes `L01298`, and the only fall-through into it is from the adapter call. **Medium** for the consequence drawn about the shop, which rests entirely on `TOWN-352`'s bitmask reading and on that row's own statement that bit 0 is not cleared; no state of `sess+0x3dc` was observed here

**Amended.** The containing-routine attribution for call sites `L01622`, `L01623` and `L01624`; the exit-set enumeration and the entry-gate reading stand: `L01622` and `L01623` are not inside `R0374` (refuted, EXP-0219). `retracted.md` holds the full entries.

### AI-CURSOR-234

The map view's dispatcher arms for `0x40e` and `0x40f` each call one of two AddChild/RemoveChild toggle routines, gated on the same child-existence predicates `TOWN-347`/`TOWN-351` already read for the widget's own paint logic, and the two arms are the same shape with two structural asymmetries. `AI-CURSOR-177` resolves `0x40e` to `L01238` and `0x40f` to `L01239`, inside the map view's private dispatcher (`R0333`-`L01626`, its own frame-setup/return, byte table at `L01230`, dword table at `L01231`); this experiment reads both arms and the four routines they call to each routine's own return. Arm `0x40e`: the call `R0380(this,2)` (`L01630`) tests presence of `[sess+0xe8]` as child id 2 of `this`; TRUE calls `R0381` (`L01631`, removes it), FALSE calls `R0382` (`L01632`, adds it).

Arm `0x40f`: the call `R0317(this,3)` (`L01633`) tests presence of `[sess+0xec]` as child id 3, through the host redirect `TOWN-347`/`AI-CURSOR-229` already document (`[sess+0xf0]` when `sess+0x3dc & 2`, else `this`); TRUE calls `R0335` (`L01634`, removes it), FALSE calls `R0383` (`L01635`, adds it). `R0380` and `R0317` are the same two routines `TOWN-347` names for the widget's own art-switch predicates: one predicate pair drives both the widget's paint state and the map view's own child presence. Each of the four toggle routines: gets the session object via `R0347` plus a vtable call at slot `+0x7c` (`SESS-OBJ-001`); adds or removes `[sess+0xe8]` as a child of `this` (`0x40e` pair) or `[sess+0xec]` as a child of `this` or of `[sess+0xf0]` under the same host-redirect test (`0x40f` pair), through `R0384`/`R0385` (`RemoveChild`/`AddChild`, `TOWN-344`); **amended by [EXP-0230], selects registry slot 7 (`ibook`) through `[L01636][7]` and passes that sample as the object pointer to `R0386`; the separately pushed `0xdc` is the terminal's priority/category byte, and rects D/E/F instead select slot 1**; and calls `R0387` (`SESS-VIEW-028`'s viewport-row recompute for child id 2 or 3).

Three of the four also reposition `[sess+0xec]` (child id 3) via `R0388` (the `SetRect` forwarder `TOWN-184`/`TOWN-222`/`TOWN-232`/`TOWN-280` establish), computing the destination from `this+0x10`/`this+0x14` offset by one of three constants (`0x55`, `0x5a`, `0xaf`); `R0335` does not reposition anything. In `R0381` and `R0382` (the `0x40e` pair) the reposition is gated on `R0317` returning non-zero and is skipped when it returns zero. In `R0383` (the `0x40f` "absent" routine) the `R0380` test at `L01637` selects only which offset constant applies (`0x5a`/`0xaf` versus `0x55`); the reposition itself runs on both of that test's branches. That row's own four callers, named only generically ("the open/close pairs the WindowProc runs for messages 0x40e/0x40f"), are exactly `L01638`, `L01639`, `L01640` and `L01641`, one per toggle routine.

Two asymmetries beyond the toggle direction. First, only the `0x40f` "present" routine (`R0335`) applies the host redirect to its own RemoveChild target; the "absent" routine (`R0383`) has one `[sess+0x3dc]&2` test (`L01642`-`L01643`), whose single taken branch both forces `[sess+0xec]`'s own rect to a fixed literal (operands `0x131`, `0x1e0`, `0x186`, `0`, not computed from `this`) and adds it under `[sess+0xf0]`; the not-taken branch adds it under `this` directly, with no literal-rect force; the `0x40e` pair applies no host redirect anywhere. Second, only the two `0x40f` routines additionally post message `0x408` to `[sess+0xd4]` through `vt+0x48` (`L01644`-`L01645` and `L01646`-`L01647`); neither `0x40e` routine posts anything past `R0387`. `sess+0xd4` is stated here only as the mechanical post target; `TOWN-086` names `campaign+0xd4` as the right-hand column's own container (a different constructor than the widget), and reconciling that with `R0387` reading session child-presence off the same object is outside this experiment

**Confidence.** High for every call site, predicate identity, toggle target, reposition mechanism, `R0387`'s four callers and the two structural asymmetries -- all instructions are read on both roots (byte-identical) to each routine's own return, and the alternative that the two pairs differ only in operand values is excluded by the added push of `0x408` and vtable call at `+0x48` and the doubled host-redirect test, both control-flow differences rather than value differences. This raises `TOWN-347`'s own Medium clause ("0x40e and 0x40f toggling the presence of child id 2 and child id 3... inferred from the widget's own paint predicate and not read from the map view's handler") to High for the toggling itself, now read from the handler directly. **Medium** for the game identity of `sess+0xe8`/`sess+0xec`: `TOWN-356` reads rect A's (`0x40e`, child id 2) icon pair as the open/closed backpack and rect B's (`0x40f`, child id 3) as the open/closed journal, and `SESS-VIEW-028` names its own four callers as these same four toggle routines, and separately states that mission entry itself invokes those routines through two other addresses (`L01648`, `L01649`) to restore "the store's `Inventory/IsOpen` and `SpellBook/IsOpen`"; the two published rows use different words for what is structurally the same object pair, and neither decodes the child class itself (`SESS-VIEW-028`'s own Medium). Unknown for what `this+0x10`/`this+0x14` are and for what message `0x408` does at `sess+0xd4`; neither was read past its own use site here

**Amended.** The child-toggle sound identity and claimed match to rects D/E/F only; toggle, reposition and post mechanics stand: The four routines select slot 7 `ibook` (refuted, EXP-0230). `retracted.md` holds the full entries.

### AI-CURSOR-235

`R0374` is a distinct routine from the map view's private dispatcher, not part of it, and it writes `this+0xdc=1` and `this+0xe0=1` unconditionally at its own entry, roughly once every 32 ticks of message `0x401`. This corrects `AI-CURSOR-233`'s attribution of two call sites to `R0374`. A scan of the declared PE sections for the dword `0x000000E0` (`corpus-scan-e0.txt`) gives 278 hits; over the complete file, including header, padding and overlay bytes, the same dword occurs 287 times (`-mode rawfull`), the 9 extra being bytes that cannot be an operand. Of the 278, 143 carry a preceding byte in the ModRM mod=10 range (`0x80`-`0xBF`), the disp32 register-indirect form; 20 of those 143 are the `MOV r32,imm32` opcode byte (`0xB8`/`0xB9`/`0xBA`/`0xBB`/`0xBD`/`0xBE`, a false positive, not a genuine ModRM byte), leaving 123.

One of those 123 is `TOWN-347`'s own already-published write, whose ModRM byte is at `L01650` (a store of 1 at `+0xe0` of the object; the dword hit itself, its disp32, is at `L01651`). A second, `L01652`, sits inside the address range `AI-CURSOR-233` attributes to `R0374`. Read from its own frame setup at `R0374`, the routine caches `this` at frame local `-0x2c` and, with no message-id test first, stores the dword `0x1` into `this+0xdc` (`L01653`) and into `this+0xe0` (`L01654`). It continues: a loop clearing bit `0x4000` from every word of the tile plane `[[this+0x80]+0xc]` over `[this+0x84] * [this+0x88]` entries (`L01655`-`L01656`; `this+0x80`/`this+0x84` are `AI-CURSOR-232`'s tile plane and row stride); a conditional walk of a list at `this+0x9b8`, gated on `this+0x9c4 != 0`, calling `[[entry]+0xc]`'s `vt+0x48` on each non-null entry and looping back to its own gating test at `L01657` (the jump to `L01657` at `L01658`), not to the tile-clear loop, which is self-contained at `L01659`-`L01656` and is never re-entered; and its own epilogue (restore of the saved registers, frame teardown, return) at `L01660`-`L01625`.

**This return ends the routine.** The very next instruction, `R0333`, is a separate frame setup beginning the map view's private dispatcher `AI-CURSOR-177` documents (byte table `L01230`, dword table `L01231`, own return at `L01626`, already read in full by this experiment for the `0x40e`/`0x40f` arms, `AI-CURSOR-234`). `L01622` and `L01623`, the two addresses `AI-CURSOR-233` gives as call sites of `R0219` "inside `R0374`", are inside this second routine, not inside `R0374`: both lie past `R0333`'s return-terminated boundary and are read directly in `disasm-dispatch-arms.txt`. `-mode callto` gives `R0374` exactly one direct `rel32` caller, `L01627`, and 0 raw dword occurrences; the same query on `R0333` gives 0 `rel32` callers and one raw dword occurrence, an `.rdata` slot at `L01628` -- the dispatcher is reached only through that slot, never by a direct call.

`AI-CURSOR-233`'s own tool (`-fnof`, nearest directly-called start) has no entry for `R0333` and reports the nearest routine that IS directly called, `R0374`, across the clean return boundary between them -- the same false-boundary class `AI-CURSOR-233` names for `R0379`/`L01619`. `L01627` is inside `R0334`, msg `0x401`'s own dispatcher arm (`AI-CURSOR-177`'s base message): it increments a tick counter at `[this+0xa70]`, and calls `R0374` only when `[L01661] == 0` and `([this+0xa70] & 0x1f) == 0` -- once every 32 calls to that arm. `L01624`, `AI-CURSOR-233`'s third address in the same sentence, is not inside `R0374`: it is past `R0374`'s own `L01625` return. It is inside a separate routine, `R0239`-`L01629`, its own frame-setup/`0xc`-byte-releasing-return pair, read with no intervening return before `L01624`; this experiment did not read that routine's body past locating its boundary

**Confidence.** High for `R0374`'s own span, its two unconditional writes, its one direct caller, the dispatcher's separate span and its own return, the correction to `AI-CURSOR-233` for `L01622`/`L01623`, and that `L01624` is not inside `R0374` -- every boundary is a frame-setup/return pair read at instruction level on both roots (byte-identical), and the `callto` asymmetry (one direct caller for `R0374`, zero plus one vtable slot for the dispatcher) is the same instrument used for the sibling case in `AI-CURSOR-233`. High for the msg-`0x401`/32-tick trigger, read directly in `R0334`. Medium that `L01624`'s containing routine is `R0239`-`L01629`: its own frame-setup/`0xc`-byte-releasing-return pair is read with no intervening return before `L01624`, but its body was not read past locating the boundary. Unknown for the tile-clear loop's and the list-walk's game effect, and for what `this+0xdc` and `this+0xe0` mean together as a pair given this second, periodic writer -- not read past the instructions cited

### AI-CURSOR-236

No reader of the map view's own `+0xe0` field was found. Two writers are now known -- `TOWN-347`'s rect C (`L01512`) and `R0374`'s unconditional entry write (`L01654`, `AI-CURSOR-235`) -- and this experiment traced neither writer to a consumer. Searched: the map view's private dispatcher, `L01238`-`L01626` (`disasm-dispatch-arms.txt`); this listing does not reach the dispatcher's own `R0333` start, and 86 bytes inside the dispatcher, `L01662`-`L01663` (the `0x406` arm's tail and the whole `0x40c` arm), are in no committed listing -- the dword scan below still covers that gap; `R0374`, `R0374`-`L01625` (`disasm-toggle-and-eaee.txt`, `AI-CURSOR-235`); the four `0x40e`/`0x40f` toggle routines and `R0387` (`disasm-toggle-routines.txt`, `disasm-predicates-and-tail.txt`, `AI-CURSOR-234`) -- `R0387`'s own listing stops at instruction 80 of 84 (instruction-budget exhaustion) and does not reach its own return; the widget's own paint routine and the base Control dispatcher, both already committed by `EXP-0210` (`experiments/EXP-0210-figure-panel-controls/evidence/listing-R0389-paint.txt`, `listing-R0390-dispatch.txt`); and the first 250 instructions of the map view's own constructor, `R0391`, which does not reach its own return (`disasm-e0-search.txt`) -- so the constructor's own tail past that budget is not covered.

A case-insensitive scan of every disassembly listing this experiment produced for the byte sequence `e0]` (a `+0xe0` memory operand written out by the tool) finds exactly two occurrences: the write at `L01654`, and `[ESI+0xe0]` at `L01664`, which is `sess+0xe0`, a different object -- `ESI` in that routine also reads `sess+0x3dc`, `sess+0xcc`, `sess+0x35c` and `sess+0x41c` -- not a read of the map view's field. Beyond the disassembly, a scan of the declared PE sections for the dword `0x000000E0` gives 278 hits (287 over the complete file, `-mode rawfull`; the difference is header, padding and overlay bytes), 123 with a preceding byte in the ModRM mod=10 range after excluding 20 `MOV r32,imm32` opcode bytes; three were traced to completion (`TOWN-347`'s write, `R0374`'s write, and the `sess+0xe0` false lead above); the remaining 120 were not traced (`corpus-scan-e0.txt`). `go run ./tools/claim -k` was run for `mapview\+0xe0`, `CMapView\+0xe0`, `sess\+0xe8` and `sess\+0xec` before this statement; none returns a row naming a reader of `+0xe0` (`corpus-searches.txt`)

**Confidence.** Unknown, scoped to the search stated. The population is not exhausted: the map view's constructor tail past instruction 250 and 120 of 123 raw-scan candidates were not read, so a reader may exist outside the routines this experiment checked

### AI-CURSOR-242

Both `pickup` loads in `R0219`'s ordinary-hover cascade are gated by the identical stack slot (frame offset -0x68), and that slot's own construction reads no dedicated "sack" bit: three ANDed tests ending in a two-armed OR over an already-named mask bit and an already-named hit-object field. The `mask & 0x23` arm's `pickup` load, `AI-CURSOR-226`'s `L01531`, is reached from the arm's own entry `L01529` through two sibling tests: a zero test of frame local `-0x20` jumping to `L01665` at `L01666` (Alt, shared with this arm's own `move` pick, `L01667`), then a zero test of frame local `-0x30` jumping to `L01668` at `L01665` (Ctrl, shared with this arm's own `attack` pick, `L01530`), then a zero test of frame local `-0x68` jumping to `L01412` at `L01668`, whose non-zero side loads `pickup` (slot 8, `L01669`, `SPR16A-CURSOR-067`) at `L01531` and stores it at `L01670`.

`AI-CURSOR-226` also names `L01423` for the routine's other `pickup` arm, "empty ground"; that address is the arm's own entry (a zero test of frame local `-0x30`, Ctrl tested first here, the opposite order from the `mask & 0x23` arm), not a `pickup` load. That arm's own `pickup` load is `L01671` (a load of the global `L01669`, stored at `L01672`), reached from `L01423` through a zero test of frame local `-0x30` jumping to `L01673` (Ctrl, this arm's own `swarm` pick, `L00629`, `L01532`), then a zero test of frame local `-0x20` jumping to `L01674` at `L01673` (Alt, this arm's own `move` pick, `L01675`), then a zero test of frame local `-0x68` jumping to `L01533` at `L01674`. Both `pickup` tests read the identical `-0x68` frame offset: the committed listing gives exactly eight accesses to that offset across this range -- three writes (`L01676`, `L01677`, `L01678`) and five reads (`L01679`, `L01680`, `L01681`, `L01668`, `L01674`) -- with no writer between the last write and either arm's comparison, and no address-of instruction anywhere in the committed listing that could alias a pointer into it, so one computed value governs `pickup` in both arms; only the Alt/Ctrl priority order and their own non-`pickup` targets differ.

frame local `-0x68`'s own construction, `L01682`-`L01678`, runs once before either arm. It starts as a copy of frame local `-0x104` = `[view+0x140]==1` (`L01405`) AND `[view+0x144]&0x1` (`L01406`) -- exactly one object selected and that object's class name is `CUnit` (`AI-PANEL-061`) -- stored at `L01407`/`L01676` (`AI-CURSOR-209` reads the same frame local `-0x104` copied one instruction later into its own composite frame local `-0x64`, `L01408`). If that copy is zero, a zero test of frame local `-0x68` jumping to `L01683` at `L01680` leaves frame local `-0x68` at zero. Otherwise frame local `-0x68` is overwritten with `[[view+0x138]+0x18c]&0x1` (`L01684`-`L01685`, stored `L01677`) -- the selected object's own field, the same field and bit `AI-CURSOR-209` cites for its `town` composite.

If that is zero, a zero test of frame local `-0x68` jumping to `L01683` at `L01681` again leaves frame local `-0x68` at zero. Otherwise frame local `-0x68` takes a third value built at `L01686`-`L01678`: frame local `-0x60` against `0x40`, equal jumping to `L01687` sets it when the whole hover hit-test mask (`AI-CURSOR-231`'s frame local `-0x60`) equals exactly `0x40` -- no other of its seven bits set, so no `CUnit`/`CAirUnit`/`CStructure` was hit; otherwise it sets frame local `-0x108` (stored into frame local `-0x68` at `L01678`) when `[view+0x98c] == [view+0x138]` (`L01688`-`L01689`, the hover hit-test's own hit-object pointer, `UNIT-HOVER-020`/`AI-CURSOR-231`, equals the selected object) AND frame local `-0x60` masked with `0x40` (`L01690`-`L01691`), else zero.

So `pickup`, in either arm, requires exactly one selected `CUnit`, that object's own `+0x18c` bit `0x1` set, and either the hover mask equal to exactly `0x40`, or the hit object is the selection itself with mask bit `0x40` also set -- on top of the two gates `AI-CURSOR-226` publishes that this whole cascade already requires to be reached at all: `[view+0x144]&0x24` clear (the zero branch at `L01692` to `L01682`; set, it picks `select`/`default` and leaves for `L01572`) and mask bit `0x4` clear (the zero branch at `L01693` to `L01694`; set, it routes to the `move`/`select`/`attack` block instead). The mask-equal-`0x40` sub-case fires only in the empty-ground arm (it requires `mask&0x23==0`, that arm's own gate); the `mask & 0x23` arm's `L01531` load can only be reached through the hit-object-identity sub-case, which requires `[view+0x98c]` -- `UNIT-HOVER-020`'s own "the hit", singular -- to equal `[view+0x138]`, the object this composite's own first refinement already confirms a `CUnit`.

`AI-CURSOR-209`'s `town` composite frame local `-0x64` instead requires mask bit `0x800`, which `AI-CURSOR-231` sets only inside the class test's `CStructure` arm of `R0218`. Whether that arm's own object-class assignment (`L01695`-`L01696`) excludes `0x800` from ever coexisting with a `CUnit` identity match is not read by this experiment, so the `L01668`-before-`L01412` test order's freedom from a real conflict between `pickup` and `town` rests on `UNIT-HOVER-020`'s singular-hit wording rather than on an excluded alternative. Bit `0x40` and the field `[view+0x98c]` are both already named (`AI-CURSOR-231`, `UNIT-HOVER-020`); the routine reads no bit and no field beyond what those rows already name, and has no dedicated "sack" test. `PARTY-FLAG-003` ([EXP-0077](../experiments/EXP-0077-party-and-save/)) reads `CUnit+0x18c` bit `0x1`, the same field `AI-CURSOR-209` cites for its `town` composite, as the player-character flag, at Medium; the earlier "hero" gloss for the same bit is itself refuted in that row

**Confidence.** High for the shared-slot fact, the frame local `-0x68` construction and both arms' full instruction sequences: one continuous listing, `evidence/disasm-listings.txt`, `L01515`-`L01572`, extracted independently for this experiment and also byte-identical to `EXP-0218`'s own committed listing over the same range; the manifest gives both preserved roots the same `rom.exe` image (`EXP-0123`), so the bytes disassembled do not depend on which root is scanned. The live alternative excluded for the shared-slot High is that the two a zero test of frame local `-0x68` tests read two different locals sharing a coincidental offset; excluded because the committed listing gives a complete, continuous account of every access to that offset across the whole span -- eight accesses, three writes and five reads, with no writer between the composite's last write and either arm's comparison -- and contains no address-of instruction anywhere that could alias a pointer into the slot. This experiment extracted no predecessor set of its own, so an edge entering the span from outside `L01515`-`L01572` is not excluded by this evidence. Medium for the `pickup`/`town` mutual-exclusion argument: `UNIT-HOVER-020` states the routine stores "the hit" in `view+0x98c`, singular, and `AI-CURSOR-231` sets mask bit `0x800` only inside the class test's `CStructure` arm, so a hit `CStructure` and the identity sub-case's required `CUnit` match are unlikely to be the same evaluation, but this experiment does not read the class-assignment block (`L01695`-`L01696`) that would exclude the alternative outright -- a second fitting model, that the walk sets `0x800` from one object while `view+0x98c` is later read as a different one, is not excluded by evidence this experiment commits. Medium for the `+0x18c` bit `0x1` reading, inherited from `PARTY-FLAG-003`'s own grade. Unknown for whether the mask-`0x40` drawable is specifically a sack: `AI-CURSOR-231` already leaves this open and this experiment does not close it

### AI-CURSOR-243

`AI-CURSOR-209` cites two operand addresses for the `town` slot load: `L01327` (the operand of the load of the global `L01180` at `L01404`) and `L01326` (the operand of the load of the global `L01180` at `L01413`). `AI-CURSOR-226` places its own `town` load, in the ordinary-hover cascade, at `L01404` -- the same instruction address and the same operand address as the first of the two. `AI-CURSOR-209`'s other site, `L01413`, is not in the ordinary-hover cascade: `AI-CURSOR-226` gives that cascade's entry as `L01515`, and `L01413` sits earlier, inside the routine's other cascade, in the tail `AI-CURSOR-230` gives as ending with the jump to `L01268` at `L01546`. One of `AI-CURSOR-209`'s two sites is `L01404`; the other is a distinct address in a distinct cascade of the same routine

**Confidence.** High. Both addresses are cited verbatim by the two rows compared, `AI-CURSOR-226`'s cascade boundary at `L01515` is itself graded High, and `AI-CURSOR-230`'s independently cited `L01546` tail address corroborates it; the comparison needs no new disassembly, only the four already-published addresses

## Order 2 producers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-ORDER-294 | A third static, structurally identical, unreachable producer of order 2 exists beside `AI-ORDER-031`'s two, and `grpAI+0x20` names no pursuit, cast or pickup value — those already-published behaviours belong to a different field. | Medium | ● active | [EXP-0322](../experiments/EXP-0322-group-order-producers/) |

### AI-ORDER-294

`R0157` (`AI-CMD-033`'s own order-5 setter) ends cleanly at `L01697` with a `0xc`-byte-releasing return; a six-instruction stub follows — two register loads, three pushes, a call to `R0301` — then a tail jump at `L01698` (to `L01699`, back into `R0157`'s own body); `L01700` then opens its own prologue — four callee-saved registers pushed with no frame-pointer setup, the stack argument at offset `0x14` taken as the list owner, a counter zeroed and `this` kept — walks a list from the owner's `+0x4` (empty-list short-circuit at `L01701`), and per member: normalizes two adjacent bytes at `[member+0x154]+0x00`/`+0x01`, conditionally calls `R0040`/`R0019`/`R0045` gated on `[member+0x154]+0x80`, copies `member+0x12c` into `[member+0x158]+0x14`, and zeroes `[member+0x158]+0x60`, `[member+0x154]+0x7c`, `member+0x54`, `[member+0x158]+0x50`, `[member+0x158]+0x38` — the same member-clear shape the amendment above reads at `L00523`/`L00318`, and the same prologue shape as this function's own neighbour `R0158`, `AI-ORDER-031`'s already-published `0x11` producer (`callto:R0158` → 0).

The walk ends (`L01702`); the owner's `+0x3c` field (the GroupAI pointer every producer in this family uses) takes order `0` then `2` (`L01703`/`L01704`); the caller's own two byte arguments, zero-extended (stack offsets `0x18` and `0x1c`) are packed into the cell at `+0xa` (`L01705`-`L01706`) — a **computed** cell, unlike the amendment's own hardcoded `0x1824` — and the routine returns (`0xc`-byte-releasing return, ten bytes of `0x90` padding follow, the same padded-not-fallen-into signature `AI-DEAD-036` uses to confirm an unrelated routine in this module is not entered by fall-through). `xcallers L01700` and `xrefs L01700:L01707` are both 0. `go run ./tools/claim -k` for `L01700`, `L01704`, `L13130` and `L13131`, run over every ledger before this row was drafted, returned zero rows; each now returns this row and no other (`evidence/q1-claim-search-order2-third-site.txt`), since publishing it necessarily put the addresses into the corpus it searches.

`AI-ORDER-031`'s own 75-hit/32-owner figure could include this address among owners its prose does not name individually, so this experiment excludes the address being *named* by any claim's text, not being one of that census's own uncounted hits. **Separately**: no value this experiment or `AI-ORDER-010`/`AI-ORDER-031`/`AI-GROUPCMD-020` finds for `grpAI+0x20` (`formats/ai/format.md`'s own eight live values, 0/1/2/3/4/5/0x11/0xff) is a pursuit, a cast or a pickup. Those three are already fully published at the **per-actor** order object, a different field at a different offset on a different structure: pursuit is `ord+0x08 ∈ {5,6}` (`AI-PURSUE-040`), casting is `ord+0x08 ∈ {8,9}` (`AI-ORDER-039`), the sack pick-up is `ord+0x08 = 7` (`AI-PROGRESS-034`). This is a reading of two already-published value spaces side by side, not a new instruction; `pipeline/SAV-COMPLETION.md`'s own wording is project text about the reimplementation, not ROM1 evidence, and is not itself a claim this experiment confirms or refutes

**Confidence.** **Medium** for the new site: one routine read whole, single entry, single exit, zero callers and zero stored references by the same two instruments `AI-ORDER-010`/`AI-ORDER-031` use, but a runtime-computed call target is invisible to both and was not excluded, and no full re-run of either census's own instrument over the whole image was made to settle whether the address is already an uncounted hit of theirs. **Medium** for the pursuit/cast/pickup separation: it rests on reading two already-High-confidence value spaces side by side; each of those two rows is itself Medium on its own writer-set completeness

## Group arm reachability, occupancy and the threat slots

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-332 | Group order 5's own per-member arm runs only while the group has a nonempty candidate list, and on each such full tick rewrites `ord+0x08` for a member with no group-assigned target… | High / Medium | ● active (amended) | [EXP-0385](../experiments/EXP-0385-dying-and-crossing-tick/EXP-0385.md) |
| AI-335 | For group order 4's Move arm, an occupied-but-passable destination is a race the routines read this round leave partly open, not a flat "never"; an impassable destination is replaced by a labelled substitute cell before the arm ever runs… | High / Medium / Unknown | ● active (amended) | [EXP-0386](../experiments/EXP-0386-engagement-under-occupancy/EXP-0386.md) |
| AI-336 | `R0020`, the occupancy-block builder `AI-BREAK-041`/`AI-FILTER-001` already name as a producer but whose own internal order neither states, visits the centre cell first, then every ring outward by increasing Chebyshev radius… | High | ● active | [EXP-0386](../experiments/EXP-0386-engagement-under-occupancy/EXP-0386.md) |
| AI-337 | Nothing in the guard-block engage pipeline tests health or dying stage, so a dying-but-not-torn-down actor remains occupancy-plane-visible, diplomacy-hostile and list-eligible through `R0020`… | High / Medium | ● active | [EXP-0386](../experiments/EXP-0386-engagement-under-occupancy/EXP-0386.md) |
| AI-340 | The group scorer's tie-break (`AI-SCORE-069`) and standing acquisition's tie-break (`AI-ACQUIRE-002`) run in opposite directions on a full tie, and neither tie-break compare itself consults the RNG… | High / Medium / Unknown | ● active | [EXP-0386](../experiments/EXP-0386-engagement-under-occupancy/EXP-0386.md) |
| AI-341 | `ord+0x78`..`+0x8c` is the actor's own class spellbook (`UNIT-SPELL-007`'s memory), not a generic "threat cache" of targets, and `R0009`'s ESI-loop draws `rand()` independently per slot (S2 against S1/S3/S4). | High / Medium | ● active | [EXP-0387](../experiments/EXP-0387-creature-spellbook/EXP-0387.md) |

### AI-332

Group order 5's own per-member arm runs only while the group has a nonempty candidate list, and on each such full tick rewrites `ord+0x08` for a member with no group-assigned target — with no test of the receiving member's own health or death stage — extending `AI-REISSUE-077`'s target-staleness axis with a second, previously untested one; membership persists for the member's entire dying window because removal is teardown-gated, not health-gated. `R0026` (group order 5's per-member `CObList` walk, 84 instructions, read whole) first gates on `[EDI+0xbb4] != 0` (`L00460`..`L00461`); when it is 0, no member is walked and the arm instead calls `R0154` at `L00463`, order 4's own fallback arm (`AI-SWARM2GATE-107`).

When the gate passes, per member: `L00775` reads the member's own group-assigned target (`ord+0x20`); nonzero, `L01708` calls `R0009` — `AI-REISSUE-077`'s own cited writer, independently re-read this experiment and confirmed to carry no `actor+0x94`/`actor+0x13c` test across its own 122 instructions; zero, the arm tests `[[this+0x14]+0x28]` (the owning `Player`'s participant-type flag): nonzero (non-human participant), `L01028` stores the byte `0xb` into `ord+0x08` directly — `formats/ai/orders.md`'s own idle-turn order — then also calls `R0113` (the optional cast, `AI-SWARM2GATE-107`) when `[this+0x4c]` bit 2 is set; zero (human participant), the arm instead only calls `R0015` — already identified as the heal AI's own entry point for a human participant's unit (`AI-SWARM2GATE-107`) — and writes no `ord+0x08` on this path; `R0015`'s own effect on `ord+0x08` is not further traced by this experiment and is Unknown.

None of the loop's 84 instructions test `actor+0x94` or `actor+0x13c` either. The loop is entered from `R0023` (`L01709`, confirmed inside that function's own body by re-walking it) — the routine `AI-CLOCK-080`/`AI-TICK-008` already establish runs once per full tick, phase-6 gated — so the rewrite cadence, when the gate passes, is once per full tick, not once per (faster) actor tick. Nothing removes a member from this walked list while merely dying: the reap sweep's own removal test (`TRIG-REAP-017`) is `actor+0x54 == 0x10`, written only on the teardown tick — at `L01710`, just before the teardown call `R0208`, and at `L00609` inside teardown itself — and `HERO-DYINGTICK-145` (this experiment) shows `actor+0x54` is force-cleared to 0, not set to `0x10`, on every tick the actor is merely dying — the reap sweep's test is false throughout dying and becomes true only on the same tick teardown fires, so the actor is a full group member for the entire window.

Corpus: the permitted archive's 48 dying records — 59 saves in 7 of the eleven pinned directories (`loadSaves` is non-recursive and admits only a direct-child `.sav` with the `Asg&` header, R5) — split 4 original (`2026-08-24/game9999.sav`, `ord+0x08 = 0` in all 4) and 44 resave (`2027-09-07`, `SAV-1059`, ROM1 resaves of Againrom-produced/modified documents); `ord+0x08` is `{0x00: 36, 0x0b: 12}` overall, with every `0x0b` in the resave population and every original at `0x00`. `0x0b` matches exactly the value the arm's own non-human-participant, null-target branch writes, consistent with that on the resave population only — the correspondence rests on a field that can be our own writer's value carried through the original's load and save, and the 4 original dying records give it no support.

No dying record in this corpus carries `5` or `6` (the pursuing values the brief's own premise names), though the alive control population carries `5`/`6` on 5 of 760 records, so this corpus did not happen to catch a record dying while under an individually-assigned pursuit target; the mechanism read is what supports that such a record's `ord+0x08` would be governed by this same arm regardless, when the group has candidates, since the order machine that would otherwise consume it never runs while dying (`HERO-DYINGTICK-145`)

**Confidence.** High for the mechanism: both routines read whole (84 and 122 instructions), the candidate gate and its own fallback target traced instruction by instruction, the cadence confirmed by independently re-walking `R0023`'s own body, and the removal trigger cited from `TRIG-REAP-017`'s own prior reading. Medium, explicitly bounded, for the corpus-explains-corpus-value inference — the `0x0b` match is supported only by the resave population (12 of 12 resave `0x0b` values, 0 of 4 original records showing it) — and for extending the mechanism's reach to a `5`/`6` record, which this corpus does not contain in the dying population to observe directly

**Amended.** The Unknown on `R0015`'s effect on `ord+0x08` is closed by `AI-349`: the call stores 0 or 8 exactly once, so the human participant branch writes `ord+0x08` through the callee, and "writes no `ord+0x08` on this path" holds for the arm's own instructions only. The four original dying records at `ord+0x08 = 0` lie within that write set for a human participant's member; their owners were not read.

### AI-335

For group order 4's Move arm, an occupied-but-passable destination is a race the routines read this round leave partly open, not a flat "never"; an impassable destination is replaced by a labelled substitute cell before the arm ever runs, not collapsed onto the actor's own cell as the original reading had it. `R0154`, read whole (146 instructions, 0 unresolved), reaches the "second test" (`L00455` compares the dword at `+0x50` of the order-block pointer with the saved value; the order-block pointer is loaded from `actor+0x158` immediately before, at `L01711`, not two instructions earlier, and the field tested belongs to that order block, not to the actor as a literal reading of `actor+0x50` would have it) from four paths converging at `L01711` — `L01712` (jmp), `L01713` (je), `L01714` (jne), and the fall-through from `L01715` — not three.

The top-of-loop arrival test (`L00452`, a 16-bit comparison with the dword's field at `+0xa`) passes only when the actor's own cell already equals `ord+0x0a`. **There is no one-actor-per-cell rule forbidding this while occupied**: `MOVE-WAIT-008` (High, active) already shows two units whose dynamic routes were computed before either claimed the cell can both end up standing on it, so the arrival test can pass against an occupied destination in that race, latching `ord+0x50 = 1` and acquiring — M1, not M2, for that case. Outside the race, the static search behind this arm (`R0178`) is occupancy-blind (`MOVE-PLANE-005`) and resolves `mover+0x76` to the literal requested cell whether or not it is occupied, so `mover+0x90` (`MOVE-072`, this round: "already at the resolved endpoint," not "route list non-empty") stays 0 while the actor has not itself reached that cell, and this arm's own test is not reached true by that path alone — the routine falls to `L01716` and re-issues.

**Whether the order can still be cancelled by giving up is a separate mechanism read for the first time this round, and it is Unknown for the ordinary occupied-but-open case, not excluded as "never."** `R0178` calls the dynamic-plane re-search `R0055` (read whole this round, 131 instructions, 0 unresolved) at `L01717`; that routine sets `mover+0x98 = 1` at `L00162` — `MOVE-GATE-039`'s own third setter, already published and independently reconfirmed by this round's own read — exactly when its own inner dynamic-plane search (occupancy live there, `MOVE-PLANE-005`) returns empty (`[actor+0x184] == 0`, tested at `L01718`) **and** the branch that picked that search's own goal cell was the one aiming straight at the final destination (`BL == 1`, set at `L01719`, taken only when the static list holds `MOVE-REFRESH-012`'s own `StaticIsntNeeded` (5) nodes or fewer — only near a route's own end, not on every tick; `MOVE-REFRESH-012` already names this routine's three target-selection branches, and this round adds what happens after the chosen search returns).

Whether an isolated occupied cell, terrain otherwise open, actually drives that inner search empty is not decided by this or any prior reading: it depends on picker A's own eight-ring substitution search (`MOVE-ALT-018`), a property of the surrounding cells this experiment did not vary. So the **gate mechanism is High** (one routine read whole this round, its branch and both conditions cited); the **occupied-destination give-up outcome through that gate is Unknown**, scoped to "only reachable near a route's own end," not a flat "never" and not a flat "sometimes" either. **The multi-tick claim — that the same pending move re-issues indefinitely without ever latching or giving up — is therefore Medium, not High**; High stays only for the single-tick, in-routine branch trace: that `L00455` is not reached true on that path absent the `MOVE-WAIT-008` race.

**An impassable destination is a different mechanism, not a worse case of the same one.** An impassable goal is unlabelled; `R0053`'s static tail runs picker A (`R0392`, bound `(D>>2)+4`) before anything collapses (`MOVE-ALT-018`, `MOVE-ALT-022`, both High and active — cited this round rather than re-derived), and only a picker-A miss (a zero answer) frees the static list and lets `mover+0x76` collapse onto the actor's own current cell, setting `mover+0x98 = 1` at `L01720` in that same tail. When picker A instead finds a labelled cell within the bound, the static list is non-empty and `R0178` takes `mover+0x76` from that substitute's own tail node (`L01721`), never from the actor's own cell: the member walks to, and latches at, the substitute — not at its own position.

**M4 (depends on occupied vs. impassable) stays TRUE, and now rests on the corrected mechanism**: the two cases diverge upstream, in the static search's own picker-A step, not in a later "did the goal collapse" test — an impassable destination usually resolves to a substitute cell and latches there; only the narrower sub-case where picker A also finds nothing collapses onto the actor's own cell

**Amended.** The nonempty-substitute and only-picker-miss-collapse clauses are
partially retracted. An admitted substitute equal to the seed leaves an empty
static route under the empty-list premise. The selected immediate caller then
takes the same zero-count branch as a picker miss. MOVE-080 through MOVE-083
bound storage and local result use; native order continuation remains Unknown.

**Confidence.** High for the single-tick in-routine branch trace and for the `R0055` gate mechanism (both read whole this round, 0 unresolved) / Medium for the multi-tick "never re-issues, never gives up" claim (the gate exists and its reachability is Unknown, not excluded) and for the impassable-case substitute mechanism (composes this round's own `R0154`/`R0178` reads with `MOVE-ALT-018`/`MOVE-ALT-022`'s already-published, not re-derived, picker-A behavior) / Unknown, explicitly, for whether an ordinary occupied-but-open destination actually drives the give-up gate, and for whether picker A finds a substitute for a given impassable destination in general (terrain-dependent, not decided by this reading)

### AI-336

`R0020`, the occupancy-block builder `AI-BREAK-041`/`AI-FILTER-001` already name as a producer but whose own internal order neither states, visits the centre cell first, then every ring outward by increasing Chebyshev radius; within a ring the four edges interleave per cross-axis offset, not edge by whole edge as this row first had it. Read whole this round (191 instructions, 0 unresolved). The centre (the argument at frame offset `+8`) is tested once before the radius loop (`L01722`..`L01723`). The radius loop (frame local `-0xc`, 1 to the argument at frame offset `+0xc` inclusive, exit branch at `L01724`) drives one `dy` sweep from `-radius` to `+radius` (`L01725`..`L01726`) that itself contains all four edge computations, not four separate sweeps: per `dy` it tests, in order, `(row+radius, col+dy)` south (`L01727`..`L01728`); `(row-radius, col+dy)` north (`L01728`..`L01729`); then, only when `|dy| != radius` (the comparison at `L01730` with an equal branch, skipping the two corners south/north already covered), `(row+dy, col+radius)` east (`L01731`..`L01732`); `(row+dy, col-radius)` west (`L01732`..`L01733`) — before the loop advances to the next `dy`.

At radius 1 the visiting order, as `(drow, dcol)`, is `(+1,-1)`, `(-1,-1)`, `(+1,0)`, `(-1,0)`, `(0,+1)`, `(0,-1)`, `(+1,+1)`, `(-1,+1)`; a whole-edge order (south entirely, then north, then east, then west) would instead visit `(+1,+1)` third rather than seventh, and this row's own prior reading described that whole-edge shape rather than the interleave the branch structure actually produces. A cell passing the occupancy test — `R0034`, not itself read whole this round, tests bit 5 (`0x20`) of the **static** plane byte `[world+cell+0x10000]` (`TERR-PASS-051`'s own "a cell record exists" bit, not a live dynamic-plane occupant test as this row first had it) — is looked up via `R0035` — `TERR-CELLREC-146`'s own routine, independently re-read this round: it gates on the same record-exists bit `TERR-PASS-051` names (`L01734` returns null when `[cell+0x10000] & 0x20 == 0`), hashes into the 100-bucket table at `world+0x540b8`/`0x540bc`, and copies the record's own 13 dwords (`TERR-CELLREC-146`'s 52-byte payload) into the scratch at `world+0x5402c`, returning its own `+4` field — and, when non-null, appended to the block's own candidate list via `R0032` (`AI-CANDLIST-106`'s own routine, independently re-read: a two-instruction forward, adding 4 and then calling `R0292`, the outer object's `AddTail`).

Because a smaller radius is always fully scanned before any larger radius begins, and every ring stops at radius `[ebp+0xc]` (`AI-BREAK-041`'s own `mover+0x08`, 5), **a candidate at a smaller Chebyshev distance from `ord+0x00` is always appended before one at a larger distance, regardless of cardinal direction**; among candidates tied at the identical distance, the append order follows the per-`dy` interleave above, not a south/north/east/west block order. What the list head decides once appended — whether `AI-BREAK-041`'s own "engage the first survivor" actually reads it — is answered by `AI-337`, not assumed here

**Confidence.** High for the scan-order structure itself, including the corrected per-`dy` interleave (one routine read whole, 0 unresolved branches, the ring bound and edge structure each a cited instruction, and the raster alternative excluded by the same read) / High for the occupancy test's own plane and bit identity, cross-checked against `TERR-PASS-051`'s and `TERR-CELLREC-146`'s own already-published citations of the same instructions

### AI-337

Nothing in the guard-block engage pipeline tests health or dying stage, so a dying-but-not-torn-down actor remains occupancy-plane-visible, diplomacy-hostile and list-eligible through `R0020` (`AI-336`) and `AI-FILTER-001`'s own diplomacy filter — and on the selector's own default path it can be engaged ahead of a living hostile deeper in the same list, because that selector tests no health field either. `R0020` (this round, whole body) tests only the **static**-plane bit-5 record-exists byte and the cell-record-exists bit (see `AI-336`'s own corrected plane wording — not a dynamic-plane occupancy test, as this row first had it); `AI-FILTER-001`'s own published text is a diplomacy relation plus a visibility test, nothing about health.

A body that has not been torn down keeps its footprint's occupancy-plane bits set — `HERO-DEATH-026`: "occupancy is released there and nowhere earlier", at the teardown call `R0208`'s own footprint walk, and nowhere in the dying arm before it — and stays on the live actor list until the same teardown tick (`HERO-DECAY-069`; `TRIG-REAP-017`, narrowed by this ledger's own `EXP-0385` entry). So the candidate is enumerated. Whether it wins the engage, this row first left open on a false premise: **`AI-BREAK-041`'s own "engage the first survivor" cites `R0009`, and that routine was not "unread by this or any prior experiment" — it was already read whole, 122 instructions, by `AI-332` (`EXP-0385`, active in this experiment's own base commit `f22b0c11`, which states it "carr[ies] no `actor+0x94`/`actor+0x13c` test across its own 122 instructions")**; this round independently re-reads the same committed body (`../EXP-0385-dying-and-crossing-tick/evidence/engage-writer-R0009-body.txt`) and confirms it.

Of the routine's four paths, the **default** path (`L01735`..`L00159`) writes `ord+0x08 = 5` and `ord+0x0c` = the candidate **list head** `AI-336` establishes, with no health read anywhere on that path. Three named exceptions leave the default: a human-owned mage at reach below 2 (`L01736`..`L01737`, calls `R0022`); the threat-cache draw (`L01738`..`L01739`, calls `R0209`); and an AI-owned mage's own cast check (`L01740`..`L01741`, calls `R0393`). So on the default path — the one every non-mage, non-threat-cached engagement takes — a dying-but-not-torn-down body at the list head is engaged ahead of a living hostile placed later in the per-`dy` interleave `AI-336` names.

The counter-example this row cited for "no health test anywhere" was `AI-ACQUIRE-002`'s own `R0022`, which does split by `HP < 1` — but the three cited addresses (`L00232`/`L00233`/`L00234`, plus `L00235`) are not inside `R0022`'s own body: `R0022`'s committed evidence contains only its own call to them, the call to `R0114` at `L01742`. Read whole this round (202 instructions, 0 unresolved), `R0114` is where the split actually happens: `L00232` compares the 16-bit value at `+0x94` with 1 and `L01743` branches when it is not below, which sends `HP < 1` candidates into a second list via two calls to `R0032` (`L00233`/`L00234`); the general list gets its own append at `L00235`. **This experiment's own original "not read whole by this or any prior experiment" was a search that should have found `AI-332`**: `go run ./tools/claim -k "R0009"` returns it directly, and returned it at the time this experiment was authored, since `AI-332` was already committed at this experiment's own base (`f22b0c11` is `EXP-0385`'s own merge commit) — see this write-up's own Corpus search, corrected.

**Guard post geometry, latches and SAV coverage are not a gap: `AI-GRPGUARD-074`/`AI-POST-095`/`AI-POST-097`/`AI-LOAD-099`/`AI-GUARD-021`/`AI-PATROL-018`/`AI-SIGHT-006` already give a complete field-by-field, routine-by-routine, tick-by-tick account for both the per-actor guard state (`actor+0x50 = 0xb`) and group order 1's own post (`ord+0x00`), and `SAV-PATROLCURSOR-571` already covers the one persisted cursor field; nothing in this experiment's own read set adds to or narrows any of those rows, and no new claim is published for that question**

**Confidence.** High for the occupancy-visibility half (composes `AI-336`'s own fresh read with `HERO-DEATH-026`/`HERO-DECAY-069`, both already High) / High for the default-path selection answer (composes `AI-332`'s own already-High whole-body read, independently re-confirmed this round against the same committed evidence, with `AI-336`'s own scan-order read — two independently-crossed High reads, not one High composed with an uncrossed citation) / Medium for which candidate is at the list head in any given crowded scene, a property of the specific scene this experiment did not enumerate

### AI-340

The group scorer's tie-break (`AI-SCORE-069`) and standing acquisition's tie-break (`AI-ACQUIRE-002`) run in opposite directions on a full tie, and neither tie-break compare itself consults the RNG — but acquisition as a whole is not RNG-free: its own no-winner branch for an AI-owned mage reaches `rand()`. `R0177` (116 instructions, 0 unresolved, this round): `L00770` compares the candidate cost with the running best and `L01744` branches when it is not below — strict `<` — so on an exact cost tie the **first**-scanned candidate is kept and every later equal-cost candidate is skipped; `AI-SCORE-069`'s own citation of the same instruction is confirmed rather than superseded. `R0022` (215 instructions, 0 unresolved, this round): on an exact distance tie, `L01745` compares two byte counts and branches when not below, then `L00228` compares a byte with `+0x12c` of the candidate and branches when above, then `L01746` compares a byte with the stack byte at offset `0x13` and branches when above — a replacement requires the new candidate's count within reach **and** its turn cost `<=` the stored best, not `<` — so on a FULL tie (identical distance and identical turn cost) the **later**-scanned candidate overwrites the earlier one; `AI-ACQUIRE-002`'s own citation of the `<=` form is confirmed.

The two compares themselves are self-contained integer branches: neither reads `R0179` (`AI-RAND-058`'s CRT `rand()`) or anything that calls it. But `R0022`'s own **fourteen** call instructions — **twelve direct, two indirect** (`L01747`/`L01748`, a flier-cost check `AI-ACQUIRE-002` already characterizes) — resolve the twelve direct calls to nine distinct addresses, only listed here, not read, except where read whole and cited by their own row: `R0114` (`L01742`, this round's own `AI-337`), `R0036` (twice, `L00224`/`L01749`), `R0051` (twice, `L01750`/`L01751`), `R0115` (twice, `L00227`/`L01752`), `R0113` (`L01753`, read whole this round), `R0015` (`L00077`), `R0040` (`L01754`), `R0019` (`L01755`), `R0046` (`L01756`) — **twelve**, not thirteen as this row first had it (the row's own list already summed to twelve; the headline count was an arithmetic error over its own list, not a re-count).

One of those nine, `R0113`, is the no-winner branch's own callee — read whole this round (120 instructions, 0 unresolved) — and it calls `R0394` at `L01757`, read whole this round (228 instructions, 0 unresolved), which calls `R0179` directly at `L01758` and `L01759`. So the chain `L01753 → L01757 → {L01758, L01759}` reaches the RNG from inside `R0022`, on the branch this routine takes for an AI-owned mage that ends the search with no winner (`L01760`..`L01753`) — after the tie-break, not inside it. The scorer's own single call, at `L00766` to `R0225`, resolves to the per-candidate cost function `AI-COST-071` names, whose own complete call inventory (four direct addresses, `R0019`/`R0167`/`R0051`/`R0115`, plus six unfollowed virtual dispatches) is `AI-329`'s already-published reading, cited here rather than re-derived; nothing in that closure reaches `rand()` either.

**Whether "list order" tracks entity creation order specifically is Unknown**: `AI-CANDLIST-106` establishes the shared scratch collection's own layout and its ten AI-module accessors, but neither this experiment nor any cited row traces what determines the order candidates are appended to it in the first place (world actor list order, spawn order, or something else) — `go run ./tools/claim -k "creation order"` returns two rows, `MAGIC-FIREPOISON-160` and `MOVE-TICK-017`, neither about a candidate list's own order, so this is a genuine gap rather than an unasked question

**Confidence.** High for both tie-break directions and for the tie-break compares' own RNG absence (both selection loops read whole this round, 0 unresolved, the compare/branch pair that decides each tie a self-contained instruction pair independent of how the compared value was computed) / High for the no-winner mage branch's own RNG reachability (the full three-hop chain read whole this round, 0 unresolved at every hop) / Medium for the group scorer's own full-pipeline RNG absence specifically: its single call goes to the cost function `AI-329` already reads at Medium, for six of its own dispatches unfollowed, so this round's own contribution is High for the selection loop and inherits `AI-329`'s Medium for the value it selects on / Unknown, explicitly, for list-order-vs-creation-order — searched and not found, not merely unasked

### AI-341

`R0009`, read whole (`evidence/functions.txt`): the loop (offset `0x78` at `L00610`, step `+4`, bound `0x84`) tests the three slots at `ord + offset` (`L01761`); for a non-null slot it calls `R0179` (`AI-RAND-058`'s CRT `rand()`, `0..0x7fff`) inside the loop body (`L01762`, one call per qualifying iteration, not once before the loop) and compares the draw against the value at `ord + offset + 0xc` (`L01763`, skipping when not below); a match copies the slot's id into the spell-id variable (`L01764`), and the loop always runs all three iterations, so later slots can overwrite it and there is no early `break`. Only after the loop does a zero test of that variable decide whether to call `R0209(session, actor, target, spell id)` (`L01739`) with the surviving value, which `MAGIC-221` proves is used exclusively as a spell id (`Spellbook::Get(actor+0x140, spell id)`), never dereferenced as a target or threat pointer anywhere in that callee's body.

It is the writer, offsets and `× 0x147` scale `UNIT-SPELL-007` publishes for a class's `Spell 1..3`/`Probability 1..3` columns: the same memory, not a coincidentally-shaped second field. The blind spot of a `disp:` sweep for this reader is a register-computed offset, not a wholesale block copy as `AI-THREAT-044`'s `[EXP-0109]` amendment speculated. The loop computes the field offset in a running variable (start `0x78`, step 4) and adds it to the order-block pointer, so its instruction stream has no literal displacement. A whole-image `EnumRefs disp:` sweep over the six literal displacements (`0x78`/`0x7c`/`0x80`/`0x84`/`0x88`/`0x8c`, 438/790/666/578/461/389 hits, matching `AI-THREAT-044`'s 438/578 counts) plus all eighteen narrower single-byte displacements in the same `0x78`..`0x8f` span (`evidence/refs.txt`) returns 203 hits inside the AI module (code-region), all at the six main displacements; each of the eighteen narrow single-byte sections contributes zero AI-module hits, and `0x8f`'s single image-wide match is a stack store outside the AI module (`L01765`, a byte store at stack offset `0x8f` plus index in `R0395`), not a call target-address literal.

The loop's own instructions are not among the 203. Disambiguating the 203 by base-register provenance (full-body reads in `evidence/functions.txt` plus a targeted `EnumRefs re:` context sweep) resolves every one to a structure other than the order block. 64 are ESP-relative stack slots in functions unconnected to `ord`/`actor`/`mover` (`R0076` 2, `R0396` 6, `R0397` 20, `R0398` 7, `R0399` 21, `R0400` 8). 109 load their base from the actor's mover pointer across 28 functions (the four read whole above plus the 24 listed below): 108 of those hits carry a direct `[actor+0x154]` load, and one, inside `R0022`, reaches the same pointer through a register copy (`actor+0x154`, `UNIT-CTOR-004`).

The mover's `mover+0x7c`/`mover+0x80` fields share the order block's numeric offsets. Full-body reads of `R0309`, `R0301` and `R0154`, plus `AI-340`'s/`AI-332`'s complete read of `R0022`, each show `[actor+0x154]` loaded or copied into a register shortly before the `+0x7c`/`+0x80` access, an idiom repeated across 24 further order-dispatch siblings (`R0085`, `R0088`, `R0308`, `R0172`, `R0173`, `R0007`, `R0010`, `R0044`, `R0401`, `R0013`, `R0402`, `R0011`, `R0012`, `R0158`, `R0006`, `R0125`, `R0060`, `R0103`, `R0190`, `R0176`, `R0200`, `R0171`, `R0063`, `R0207`, the family `TRIG-GRPARM-047` names through its `R0309`/`R0007` reads, which do not mention `ord+0x7c`/`+0x80` because both offsets belong to the mover there, not the order block); `actor+0x88` (Mind, word, excluded by `MAGIC-MIND-010`'s dword-only filter, identical instructions at `L01766`/`L01767`/`L01768`/`L01769`); `actor+0x8c` (Speed, byte, `UNIT-COMBAT-006`, `R0108`/`R0157`); a `0x24`-byte container `UNIT-CTOR-004` allocates at `actor+0x7c` (`R0262`, 7 hits, whose two stores at `L01770`/`L01771` install a fresh object of that type built by `R0403`, or null; plus `R0404`, `R0405` and `R0406`, 1 hit each); the pre-create control `UNIT-PANEL-010` names, reached through the global `[L00285]` (`R0169`, 1 of its 2 hits, the other being a Mind hit counted above); the order block itself, reached through that same global rather than directly through `actor` (`R0170`, 7 hits, `SAV-ACTORINPUT-547`'s post-read key repair, each hit forming `base+0x88` from the global at `L00285`, an `ord+0x88` access cached behind a global pointer, not a UI subsystem); a `+0x7c` field on an unrelated class's `this+0x30`/`this+0x3c` sub-objects (`R0407`, a destructor-shaped routine with no `actor`, `mover` or `ord` pointer in its body, 3 hits); and one generic boolean accessor with no order-block reference (`R0408`, 1 hit).

The writer `R0184` is partly visible to the same sweep, 1 hit each at `disp:78`/`disp:84` (base plus `4 x index` plus `0x78` or `0x84`, a literal-base-plus-index loop), so the field is census-visible on the write side and census-invisible on the read side. Class reachability: `R0184`'s `actor+0x4c` bit-setting instructions are an OR with `0x2` (`L01772`, all 12 spellbook classes, `Spell 1 > 0`) and, later in the same function, exactly one further OR with `0x6` (`L01773`, gated on `actor+0x0e ∈ {0x47,0x48}`, Dragon and Daemon only), which confirms `UNIT-SPELL-007`'s citation at the instruction level. For an actor without the mage bit, this loop is unconditional inside `R0009`; a mage-bit actor with reach `< 2` and owner `+0x28 == 0` skips it instead, and owner `+0x28 == 0` by itself blocks only `R0393` (`L01774`).

Only Dragon and Daemon carry the mage bit that can reach `MAGIC-AI-012`'s Mind-gated path (`R0393`) through `R0009`'s gate (`L01740`). No writer other than `R0184`'s spawn setup and `Order::Serialize`'s raw LOAD copy was found; the disambiguation covers the AI module: of the 203 AI-module hits, none writes the order block at these offsets; the hits that write at all write a different structure (chiefly `mover+0x7c`/`mover+0x80` in the order-dispatch family above), and every hit whose base register is the order block is confined to `R0009`'s loop, which only reads the pair and copies a matched id into `EBX`, never storing back into the struct. `SAV-1066` names the one writer this census cannot see: `Order::Serialize`'s `CArchive::Read(this,0x94)` (`L00551`) copies all six dwords wholesale from the archive on LOAD and addresses no single field by a literal offset.

This narrows `AI-THREAT-044`'s open item 17 ("what fills the threat cache") to a bounded negative: within this census, nothing *refills* a slot after `R0184`'s spawn-time write other than a SAV load rewriting the whole block; a level change or script reaching a slot through a pointer indirection or another register-computed access is not excluded, only unseen by every literal-displacement instruction this instrument reads

**Confidence.** High for the loop structure, the per-slot independent draw (S2), the spell-id identity of the value it hands to `R0209` (each a complete-body read) and the register-computed blind-spot mechanism (the loop's instructions contain no literal displacement, a positive fact about the bytes, not an absence claim) / High for the census figures themselves and their disambiguation (consistent with `AI-THREAT-044`'s counts; all 203 AI-module hits traced by base register to a structure other than the order block — four of the 28-function mover family read in full, the rest matched by the same two-instruction idiom rather than individually disassembled) / Medium for "no writer beyond spawn setup and `Order::Serialize`'s LOAD copy", bounded to what a `disp:`/narrow-displacement literal sweep plus this base-register disambiguation can see

## Open, written out

1. ~~**The AI state machine.**~~ — *closed by [EXP-0082]*: the switch has **27** arms, not twenty,
   and all of them are named (`AI-STATE-011`). ~~Still open from that item: how `order+8` states
   1/5/6/0xb are consumed — the pursuit itself — is unread~~ — *closed by [EXP-0099]*: all twelve
   live `ord+0x08` arms are read (`AI-ORDER-039`), 5 and 6 are the pursuit pair (`AI-PURSUE-040`)
   and `0xb` is the idle turn. It was never a `claims/move.md` area — the pursuit is decided in the
   AI module and only *executed* by the mover.
2. ~~**The group layer.**~~ — *closed by [EXP-0082]* for the parts that decide behaviour
   (`AI-TICK-008`, `AI-GROUP-009`, `AI-ORDER-010`). ~~Still open: group orders **2, 4, 5 and
   0x11**~~ — *all four read by [EXP-0087]* (`AI-SWARM-022`…`AI-ROAM-025`). Still open there: the
   candidate scorer `R0252`, and `R0225`, the *other* scorer~~ — *both read by
   [EXP-0112]* (`AI-COST-071`, `AI-REACH-072`), together with the builder `R0110`
   (`AI-GROUPSEE-068`), the assigner `R0177`/`R0153` (`AI-SCORE-069`), the matrix they
   index (`AI-PREF-070`) and group order 1 end to end (`AI-GRPGUARD-074`). ~~**Still open there:**
   group orders 2 and 5 were read by [EXP-0087] before the scorer was, so what `Swarm` and `Swarm 2`
   do with a `0xffffff` veto was never stated; and `R0109`, the fifth caller of the builder,
   is outside the AI module and unread.~~ — *order 5 closed by [EXP-0126]*: `AImgr+0xbb4` is the
   candidate collection's element count, so an empty list takes the arm out to order 4 and the
   members walk, while a list the matrix vetoes keeps it in its own body, whose two zero-target
   branches do not move (`AI-CANDCOUNT-105`, `AI-SWARM2GATE-107`). *`R0109` closed by
   [EXP-0114]* (`AI-DIPLO-084`): it is the spell-cast twin of `R0226`, and the only writer
   of the matrix that re-scans rather than waiting for the next AI tick. **Still open there: group
   order 2**, which has no gate and is the cheaper of the two to read.
3. ~~**The line-of-sight test** `R0136`~~ — *read whole by [EXP-0112]* (`AI-LOS-081`): it is
   an accumulator march, not a predicate on altitude. ~~Open in its place: what fills the per-cell
   step grid at `fog+0x22000` and the per-cell cost grid at `fog+0x28000` — both are read by it and
   neither was traced to a writer.~~ — *closed by [EXP-0117]*: one writer, `R0273`, run once
   from the sight object's init at world construction, from geometry and `ScanShift` alone
   (`AI-LOS-087`…`AI-LOS-091`). **Open in its place, and smaller:** `R0409` and
   `R0410` — the two routines that consume the *other* grid the same init builds — were
   placed but not read, so what the per-player `u16` mask they OR into `fog+0x00000` is used for is
   unknown; ~~and the flat-ground region measured here (145 cells at sight 6) does not equal the 127
   `TERR-FOG-080` publishes for the structurally parallel client implementation at the same `k`,
   which is either a real divergence between the two or a defect in one of the two readings.~~ —
   *adjudicated by [EXP-0120]*: 145 is right and the two are one algorithm; 127 is
   `TERR-FOG-080`'s transcription stopping four instructions short of the end of `R0283`
   (`TERR-FOG-117`, `AI-SIGHT-093`).
4. ~~`actor+0x4c` bit 2 and `actor+0x0e`~~ — *both settled by [EXP-0112]* (`AI-GATE-079`); bit 2 had
   been answered in `claims/hero.md` for two experiments and no AI row cited it.
5. Whether `R0146`'s reach overwrite is reachable outside a multiplayer session — *narrowed
   by [EXP-0112] rather than closed: its own `actor+0x0e == 0x18` gate names 57 shipped `Humans`
   rows, so the gate is not the thing that makes the arm unreachable; `AI-STATE-043`'s missing
   writer for state `0x17` is.*
6. **`grpAI`'s unread fields.** `+0x30`/`+0x34` are a latch and a counter the guard arm advances;
   `+0x08`, `+0x09`, `+0x0a`, `+0x3c` are written but their consumers were not traced; `+0x45` is now traced as a rebuilt spatial activity count (AI-ACTIVITY-324). ~~`+0x24`, `+0x2a`, `+0x2b`~~ — *sourced by [EXP-0094]*: they
   are `R0140`'s outputs, the group's **fine centroid** (`+0x24`), its **packed centroid
   cell** (`+0x28`), the **maximum Chebyshev spread** from a member to it (`+0x2a`), the maximum
   member `actor+0xa5` sight (`+0x2b`) and the maximum of the two summed (`+0x2c`). Only the
   centroid pair has a traced consumer — the formation gate (`MOVE-GATE-035`) — so what reads the
   three maxima is still open.
7. ~~**`actor+0x54`.** … The full value set is unknown.~~ — *closed by [EXP-0090]*: the actor tick
   `R0037` dispatches on `actor+0x54 − 1` over a **15**-byte index at `L01775` into 6
   dwords at `L01776`, so the consumed set is `1` → `L01777`, **`2` → the pick-up**
   (`ITEM-PICK-016`), `3`/`0xd`/`0xe` → `L00005`, `0xf` → `L00196`, and `4..0xc` on the
   default — nine of fifteen empty. Beside those, `0x10` is death (`ITEM-DEATH-012`, and
   `SESS-TICK-006`'s skip) and `0x1a` this machine's own arm `0x1a`. `R0016` **preserves**
   only 2 and `0xf` across a tick and clears everything else (`L00539`…`L00540`). Still open:
   what the arms for 1, 3/0xd/0xe and 0xf do.
8. ~~**The mission description file.** `World\Mission\<n>.ini` is the only surface that authors a
   behaviour~~ — *closed the wrong way by [EXP-0083]: it is not the only surface, and it is the one
   that does not ship. The `.alm`'s own type-7 script authors group behaviour on 20 campaign maps
   (`AI-GROUPCMD-020`) and patrol on 8 of them (`AI-PATROL-017`).* Still open from that item: the
   `.ini`'s grammar past `Patrol` / `StandGround`, which cannot be measured because none ships.
9. ~~**The four unread group arms.**~~ — *closed by [EXP-0087]*: `AI-SWARM-022`, `AI-MOVE-023`,
   `AI-SWARM2-024`, `AI-ROAM-025`. Opened in their place, each narrow: ~~what fills
   `AImanager+0xbb4`, the gate that decides whether order 5 is `Swarm 2` or `Move`~~ — *closed by
   [EXP-0126]: it is the candidate collection's element count, so the gate is "did this group see
   anything" and the fallback is the ordinary path (`AI-CANDCOUNT-105`, `AI-SWARM2GATE-107`). Still
   open from that item: how often the branch is taken in a shipped mission, which only a running
   original can say — the prediction is written out in that experiment*; whether
   `Roam`'s re-roll loop can fail to terminate for a group cornered against the map rectangle;
   what fills `pth+0x08`, the `Wimpy` radius; ~~whether `actor+0x50 = 0x16` — the withdraw's own
   state-machine arm, table slot 22 — has a register-form writer~~ — *closed by [EXP-0099]:
   exactly one, `L00206` with `EBP = 0x16` from `L00605`, inside `R0103`, the player's
   order `0x14` (`AI-STATE-043`)*; and the entry point and reachability of the orphan routine at
   `L00316..L00317`, which writes `grpAI+0x20 = 2` (`AI-ORDER-031`).
10. ~~**`Target_Group` when one id has two owners.**~~ — *closed by [EXP-0438]* (`AI-366`,
    `AI-367`): the resolver is an id-only map in which the later owner's group replaces the earlier
    one. ~~`AI-GROUP-009` keys a runtime group on (owner, groupId); the script names a group id alone.
    One shipped patrol node (`scn:100` group 25) names an id carried by two owners.~~ Still open there:
    whether a SAV restore rebuilds the player and group lists in the same order.
11. **`order+0x00` for a group-guard member.** `AI-GUARD-012` reads the per-actor arm, which sets the
    post to the current cell when it is 0; the group arm `R0024` has no such initialisation and
    its walk-home writes `order+0x0a = order+0x00` (`L00338`). Where that sends a member whose post
    was never set is unread.
12. **The order object's own bytes.** `R0016` dispatches on `ord+0x09` and `ord+0x08`
    (`AI-PROGRESS-034`) and the whole object moves as a raw `0x94`-byte block in a save
    (`R0199`). ~~What the other bytes of that block are — and which of `ord+0x09`'s four live
    values means what — is unread; the twelve live `ord+0x08` arms were classified by their calls,
    not read.~~ — *both switches read by [EXP-0099]* (`AI-ORDER-039`), and the fields they consume
    are named: `+0x00` post, `+0x02` patrol waypoint, `+0x04` patrol re-anchor latch, `+0x0a` cell,
    `+0x0c` target, `+0x14` stop distance, `+0x15` attack counter, `+0x18` followed actor, `+0x20`
    group-assigned target, `+0x28`/`+0x30` spell target and spell, `+0x3c` spell cell, `+0x40`/
    `+0x44` the withdraw thresholds, `+0x70` follow leash, `+0x78`..`+0x8c` the actor's own class
    spellbook slots (`AI-341`, narrows `AI-THREAT-044`'s "threat cache" naming), `+0x90` the
    patrol ring. Still open: the remainder of the `0x94` block.
13. ~~**What enqueues a command.**~~ — *closed for the **order** space by [EXP-0109], and not by
    sweeping the pool: the producers are found from the other end, by the opcode store
    the `re:` search for an 8-bit store to `[E.. + 0x9]` with an immediate source, which lands a one-routine-per-opcode builder family in
    `R0215..R0238`. Every order opcode has exactly one builder, the map click chooses
    among them by cursor (`AI-CLICK-050`) and the command panel issues three of them directly
    (`AI-PANEL-053`).* Still open: the **session** space, and with it whether the source-3 pick-up
    is ever emitted; and `R0233`, a second surface reaching five of the same builders, which
    was not read.
14. **A fault, stated not observed.** `R0159`'s ring search ends by zeroing its result (`L01778`) and then
    reading the dword through that pointer (`L01779`), a null dereference when `order+0x02` is not a member of
    `order+0x90`. Neither live setter can produce that state; the one that would is dead
    (`AI-PATROL-019`).
15. **The formation mode's third state.** `AI-FORM-037` shows that a value other than `0` or `2`
    puts a group in formation with **no** spread test, and that the trigger instant can author one.
    Whether the game's own interface can reach it — and what the shipped catalogue's `Par0`/`Par1`
    labels mean against a runtime that takes the value from `Par0` — are interface- and
    editor-layer questions, not questions about these routines.
16. **What reads `grpAI+0x2a`/`+0x2b`/`+0x2c`.** `R0140` publishes a group's maximum spread,
    maximum member sight and their sum on every call, and `EXP-0094` traced a consumer only for the
    centroid pair.
17. ~~**What fills the threat cache.**~~ Closed by `AI-341` (`EXP-0387`): the "threat cache" is
    `UNIT-SPELL-007`'s own class spellbook slots, written at spawn time by `R0184` and, on a
    SAV load, wholesale by `Order::Serialize`'s own `CArchive::Read(this,0x94)` (`L00551`,
    `SAV-1066`) — the writers a per-displacement census cannot see because that read addresses no
    single field by a literal offset. A fresh whole-image `disp:` census over all six displacements
    and eighteen narrower single-byte ones returns 203 hits inside the AI module, none of them the
    order block: 64 are ESP-relative stack slots, 109 resolve to the actor's own mover
    (`mover+0x7c`/`mover+0x80` at the same numeric offsets, an idiom shared by 28
    order-dispatch functions), and the rest resolve to `actor+0x88` (Mind), `actor+0x8c` (Speed), an
    `actor+0x7c` container `UNIT-CTOR-004` allocates, the pre-create control `UNIT-PANEL-010` names,
    the order block itself reached indirectly through that same global, an unrelated class's own
    sub-object field, or a generic accessor with no order-block reference — so once disambiguated by
    base register, the census finds no other write-shaped instruction to the order block at any of
    these offsets, and `R0009`'s own consuming loop never stores back into the struct either.
    The instrument's own blind spot is now identified: the reader computes the field offset in a
    register (`ESI`, stepped `+4` per iteration) with no literal displacement at all, not a wholesale
    block copy as prior passes speculated, and is invisible to the census for that reason — the
    spawn-time writer remains partially `disp:`-visible (1 hit each at `0x78`/`0x84`) while the reader
    is fully invisible and the SAV-load writer is invisible by a different mechanism (a block read,
    not an indexed store). Whether a slot can be *refilled* through some other path neither the
    literal-displacement instrument nor this disambiguation can see (a saved pointer, or another
    register-computed access) is not excluded, only absent from every instruction this census can
    read; a SAV load itself is no longer in that Unknown (`SAV-1066`).
18. **The five live `actor+0x50` arms with no writer.** `4`, `0x17`, `0x18`, `0x19`, `0x1a`
    (`AI-STATE-043`), `4` being written only from inside arm `0x18`. Two of them write hard-coded
    map cells (`AI-GUARD-012`), which reads as leftover test code — but the sweep is store-form and
    cannot see a state arriving from a save, so *argued* unreachable, not *shown*.
19. **Tick order inside one full tick.** `R0008` runs from the group AI slot and
    `R0016` from the actor tick. Which comes first decides whether a rewritten `ord+0x08`
    is executed on the same tick or the next, and therefore how many cells a creature overshoots
    before it turns round. `SESS-TICK-006` fixes the slot; the relative order was not read.
20. ~~**Whether the per-actor and group leashes can both be live.**~~ — *the corpus half closed by
    [EXP-0102]*. A group at order 0 runs `AI-BREAK-041`'s 5-cell post block; one at order 1 runs
    `AI-RADIUS-014`'s jittered radius about the centroid and never enters the per-actor machine
    (`AI-ORDER-010`). The census now exists: **at load the byte is never 0** — 8085 of 8094 EN
    placements are at 1 and the other 9 at 3 — and even assuming every shipped group command fires,
    the post rule reaches at most **68 of 8094** (`AI-CENSUS-046`, `AI-CENSUS-047`). A player meets
    the **centre** rule. Still open, and now load-bearing rather than a detail: **whether
    `R0067`'s `server+0x0c == 0` gate wraps the whole walk or only its aggressive branch** —
    if the whole walk, a multiplayer load leaves every group at the constructor's 0 and the answer
    inverts.
21. **The three cursor gates.** `[L00627]`, `[L00670]` and `[L00632]` decide, before any
    property of the hovered object is consulted, whether the attack, select and move cursors are
    available at all (`AI-CURSOR-052`). `R0230` writes one value to all three;
    `R0231` sets two and `R0229` clears all three. No caller of any of them was read,
    so whether they are a screen mode, a key state or a lobby flag is unknown — and every clause in
    this ledger that stands on the cursor is capped at Medium until one is.
22. **The selection capability mask** `R0093`. It decides which of the eight armed modes a
    given selection may enter (`AI-PANEL-053`), so it is what actually answers "can this unit be
    ordered to attack". It was located, not read.
23. **`view+0x9b4` against `actor+0x14`.** `R0212` compares an actor's owning `Player`
    (`0x70` bytes) with the view-side player record (`0x48` bytes) at `L01780`. Both readings are
    established at instruction level and they cannot both be the same object
    (`UNIT-VPLAYER-021`); which one that comparison is really made against is unread.

Items 24 to 28 were filed on 2026-08-15 from the *AI's own machine* door in
`docs/DOORS-OPEN.md`, when that door was retired to `docs/DOORS-CLOSED.md`. The wording is the
door's own.

24. **What fills `actor+0x12c`.** The field decides which branch of both scorers a creature takes,
    the melee row `[member][cand]` or ranged row 0, and only the melee rows carry the flier veto
    (`AI-PREF-070`, `AI-REACH-072`). It was never traced, so `AI-FLIER-073`'s reach is bounded by a
    field whose value space is unknown. One write is known, `L00278` in `R0146`
    (`AI-GUARD-007`), and it sits in a state arm `AI-STATE-043` finds no writer for, so it is not
    the ordinary source.
25. **Whether `R0075` runs in ordinary play.** It calls the AI walk with no `% 16` gate in
    front of it, which is the one thing that would make `AI-TICK-008`'s cadence wrong. Its callers
    were not read.
26. **Which spell carries id 20.** `AI-COST-071` shows that a candidate bearing it costs a flat
    +127 to target, which removes it from selection at any distance. Which authored spell holds
    that id is unread.
27. **How often group order 5's fallback fires in a shipped mission.** [EXP-0126] separates a
    vetoed group from a blind one (`AI-CANDCOUNT-105`), and the frequency of the walking arm in
    play is not a property of the image. Only a running original can answer it.
28. **Whether a player-controlled hero is ever the subject of a server sight stamp.** That is what
    bounds the hero-truncation divergence `AI-SIGHT-094` ([EXP-0120]).
29. **What the array at minimap `this+0x9bc`/`+0x9c0` holds.** `AI-MINIMAP-156` reads the loop that
    walks it and blits one entry per iteration through the same external-object convention used for
    the terrain layers; the array's own contents were not identified.
30. **The call path from the command panel's own `0x40c` post to `R0097`.** `AI-PANEL-158`
    shows the two routines are distinct and neither calls the other directly; `R0097` has 14
    direct call sites in `.text` and none was traced back to a `0x40c` handler.
31. **What sets the cursor outside the mission map.** `AI-CURSOR-175`'s classify-setters census
    shows several set-cursor call sites (e.g. `L01781`, `L01782`, `L01783`) in the
    `L01784`-`L01785` address range this ledger's own town/shop rows (`AI-MINIMAP-156` and
    neighbours) already work in, and a second `[current-cursor]` access cluster sits in that same
    range. Neither the containing functions nor the surface each belongs to (town, shop, school,
    tavern, main menu, dialogue, an inventory or character panel, any modal) was read in this
    experiment. A keyword search of every ledger for "town cursor"/"shop cursor" returned nothing
    at the time this item was written.
32. **The idle rule: what selects the cursor with nothing armed and the pointer over no actor.**
    `default` is registry slot 0 and the single most common target among `AI-CURSOR-175`'s 69
    classified set-cursor calls (36 of 69), which is suggestive but was not traced to a specific
    "no hit" code path. Whether a reachable no-cursor state exists, and what it shows if so, was
    not investigated.
33. **The tail-jump target `R0325`.** `AI-CURSOR-172`'s decrement helper `R0322`
    reaches it on the manager's refcount going `1->0` with `mgr+0x4` non-zero, instead of calling
    the frame-advance routine `R0324` the increment helper calls on the opposite
    transition. Not disassembled in this experiment.
34. **The 11 untraced callers of the frame-advance gate `R0323`, outside `R0321`.**
    `L01184`, `L01185`, `L01186`, `L01187`, `L01188`, `L01189`, `L01190`, `L01191`,
    `L01192`, `L01193`, `L01194` (`AI-CURSOR-172`). Whether any of these runs on a cadence
    that does not immediately follow a `R0321` reset decides whether the cursor animator is
    ever observed to advance past frame index 1 in play.

`EXP-0213` was allocated `AI-CURSOR-172`..`AI-CURSOR-187` (16 ids) and spent
`AI-CURSOR-172`..`AI-CURSOR-178`, 7 ids. **`AI-CURSOR-179`..`AI-CURSOR-187` (9 ids) are
returned unused**, none ever reissued. The next free `ai.md` id is therefore `AI-CURSOR-188`.

`EXP-0214` was allocated `AI-CURSOR-188`..`AI-CURSOR-201` (14 ids) and spent
`AI-CURSOR-188`..`AI-CURSOR-196`, 9 ids. **`AI-CURSOR-197`..`AI-CURSOR-201` (5 ids) are
returned unused**, none ever reissued. The next free `ai.md` id is therefore
`AI-CURSOR-202`.

`EXP-0216` was allocated ids `202`..`217` of `claims/ai.md` (16 ids, one number space across
all stems) and spent `AI-CURSOR-202`..`AI-CURSOR-209`, 8 ids. **`AI-CURSOR-210`..`AI-CURSOR-217`
(8 ids) are returned unused**, none ever reissued. The next free `ai.md` id is therefore
`AI-CURSOR-218`.

`EXP-0217` was allocated ids `218`..`225` of `claims/ai.md` (8 ids) and spent all eight,
`AI-CURSOR-218`..`AI-CURSOR-225`. **No id is returned unused.** The next free `ai.md` id is
therefore `AI-CURSOR-226`.

`EXP-0218` was allocated ids `226`..`233` of `claims/ai.md` (8 ids, one number space across
all stems) and spent `AI-CURSOR-226`..`AI-CURSOR-233`, all 8. None are returned. The next free
`ai.md` id is therefore `AI-CURSOR-234`.

`EXP-0219` was allocated ids `234`..`241` of `claims/ai.md` (8 ids) and spent
`AI-CURSOR-234`..`AI-CURSOR-236`, 3 ids. **`AI-CURSOR-237`..`AI-CURSOR-241` (5 ids) are
returned unused**, none ever reissued. A returned range stays retired; the next free
`ai.md` id is therefore past it, at `AI-CURSOR-242`.

`EXP-0385` was allocated ids `332`..`334` of `claims/ai.md` (3 ids) and spent `AI-332`,
1 id. **`AI-333`..`AI-334` (2 ids) are returned unused**, none ever reissued. A returned
range stays retired; the next free `ai.md` id is therefore past it, at `AI-335`.

`EXP-0386` was allocated ids `335`..`340` of `claims/ai.md` (6 ids) and spent
`AI-335`, `AI-336`, `AI-337`, `AI-340`, 4 ids. **`AI-338`..`AI-339` (2 ids) are
returned unused**, none ever reissued: Part 3 of the brief (a held pursuit refused
by bodies with no alternate approach) needed no new row even after this
experiment's own correction pass widened it. `EXP-0373`'s own `AI-327`/`AI-328`/
`AI-329`, plus the already-published `AI-ROUTE-045`/`AI-REISSUE-077`/`AI-GUARD-007`/
`MOVE-GATE-039`, answer what happens once `mover+0x98` is actually raised:
footprint-blind, route-blind reacquisition (**P1**, though overstated as a clean
"retarget" — it can re-select the same victim) and a genuine give-up (**P3**) once
nothing is in reach. They do not decide whether an every-approach-occupied pursuit
ever raises `mover+0x98` in the first place: this correction pass (`MOVE-073`) found
that it can, through the same dynamic-plane re-search gate `AI-335` names — but
whether it does for a given every-approach-occupied scenario is Unknown, and an
alternate-approach substitution (**P2**) is not refuted: the re-search's own
contact-ring picker is exactly what a rejected goal falls back to, and no reading
traces that picker for the specific case where every ring cell is itself occupied.
This is a correction to how Part 3 composes, not a new fact needing its own row —
`MOVE-073`, already filed under `claims/move.md`, carries it. A returned range
stays retired; the next free `ai.md` id is therefore past it, at `AI-341`.

`EXP-0387` was allocated ids `341`..`343` of `claims/ai.md` (3 ids) and spent `AI-341`,
1 id. **`AI-342`..`AI-343` (2 ids) are returned unused**, none ever reissued: the
identity/mechanism finding for `ord+0x78`..`+0x8c` and the class-reachability gate both
fit inside one row together with the `MAGIC-221`/`SAV-1066` cross-references that carry
the rest of the brief's ten questions. A returned range stays retired; the next free
`ai.md` id is therefore past it, at `AI-344`.

## Original reacquisition and Retreat boundaries

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-360 | In R0004, a later diplomacy-rejected or dead candidate can lower the running-nearest gate; permuting the same actors can change the surviving list and winner. | High | ● active | [EXP-0420](../experiments/EXP-0420-reacquisition-cycle/) |
| AI-361 | Reacquisition scores admitted candidates by the circular byte turn cost alone; equal scores replace the winner, so the last surviving equal-score candidate wins. | High | ● active (amended) | [EXP-0420](../experiments/EXP-0420-reacquisition-cycle/) |
| AI-362 | Reacquisition traverses the bound actor manager head to tail; the selected creator, return and unlink paths append or preserve survivor order, rather than sorting by actor ID. | High / Medium | ● active | [EXP-0420](../experiments/EXP-0420-reacquisition-cycle/) |
| AI-363 | The selected original load rebuild appends on-map actors in player, group and member order; it does not preserve the pre-load global actor-list sequence as a separate ordering input. | Medium | ● active | [EXP-0420](../experiments/EXP-0420-reacquisition-cycle/) |
| AI-364 | Group order 0 dispatches command state 0x16 without a progress gate; an executor invocation that clears nonzero progress skips pending dispatch, while the next invocation entered at zero can execute it. | High / Medium / Unknown | ● active | [EXP-0420](../experiments/EXP-0420-reacquisition-cycle/) |
| AI-365 | The executor's active route-failure tail is a same-invocation exception: command state 0x16 can reacquire and reinstall strike progress after its old progress was cleared, without resetting action phase. | High / Unknown | ● active | [EXP-0420](../experiments/EXP-0420-reacquisition-cycle/) |

### AI-360

The complete selected selector is R0004. Distance admission tests reach at
L00617 and the running minimum at L00801. Self exclusion at L00799
precedes the accepted candidate's append and minimum store at L00802.
Previously appended candidates are not removed when a later candidate lowers
that minimum. The diplomacy pass runs afterward at L00803. The signed-HP
partition begins at L00804 and restores parked dead candidates when the
living list is empty. Neither filter rebuilds the earlier distance admission.

Original-instruction replay uses two hostile living actors at cell distances
4 and 3 plus a distance-1 actor. All six permutations are measured for a
nonhostile third actor and all six for a dead hostile third actor. With the
nonhostile first, only it is admitted and diplomacy removes it, leaving no
winner. With it last, the two living hostiles survive. A dead actor first
can exclude both living actors and then win the corpse fallback. The actual
original diplomacy and partition instructions run; no hostility-filter hook
supplies these outcomes.

**Confidence.** High for the bounded original routine and these controlled
permutations. EN/RU are one identical executable population. The 268-instruction
selector has no unresolved transfer. Section mappings, instruction digests,
raw relative targets and load-bearing starts are preserved. Explicit synthetic
allocation, footprint and post-selection hooks do not move either filter or
the minimum store.

**Unknown.** Native occurrence of the controlled actor arrangements and their
full-game construction history; invisibility paths and external list writers
outside this selected population.

### AI-361

R0115 masks both arguments to a byte, subtracts them, takes the absolute
value and negates the low byte when it exceeds 128. Its low-byte result is
min(abs(facing-direction), 256-abs(facing-direction)). All 65,536 byte pairs
are replayed. The selector masks the return at L01786. The selected direction
helper R0051 returns one of eight byte headings, multiples of 32; it rounds
its sector before the final shift. The 81 controlled centre-coordinate
deltas also retain the coincident-centre result rather than assuming a
special no-direction sentinel. This byte-pair enumeration does not assume
which facings native movement can produce.

The selector compares only this cost at L01787. JG at L01788 skips a
larger cost; equal cost executes the winner store at L01789. A two-actor
equal-cost permutation selects the later survivor. Changing that private
branch to JGE makes the same loss control select the earlier survivor.
The distance-4 actor can win over an admitted distance-3 actor when its turn
cost is lower. Reversing their scan order can exclude the farther actor before
scoring. Distance gates admission; it is not a final score term here.

**Confidence.** High for this helper's complete 12-instruction body and this
selector's complete score loop. The byte enumeration is finite and includes
the 128 boundary. The source score data flow and private branch mutation
exclude strict first-winner ties and a distance-weighted final score.

**Controlled direction outputs.** The existing `direction.tsv` instruction
replay supplies the 81 heading bytes below. Both synthetic actors have
footprint size 1 and position fractions 128/128. The decider is at cell
50/50; each target cell is 50 plus the named dx/dy. The equal-centre entry
returns 224, rather than a no-direction sentinel. This table transcribes
the accepted replay output; no new original-runtime observation is implied.

| target dy \ target dx | -4 | -3 | -2 | -1 | 0 | 1 | 2 | 3 | 4 |
|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| -4 | 224 | 224 | 224 | 0 | 0 | 0 | 32 | 32 | 32 |
| -3 | 224 | 224 | 224 | 0 | 0 | 0 | 32 | 32 | 32 |
| -2 | 224 | 224 | 224 | 224 | 0 | 32 | 32 | 32 | 32 |
| -1 | 192 | 192 | 224 | 224 | 0 | 32 | 32 | 64 | 64 |
| 0 | 192 | 192 | 192 | 192 | 224 | 64 | 64 | 64 | 64 |
| 1 | 192 | 192 | 160 | 160 | 128 | 96 | 96 | 64 | 64 |
| 2 | 160 | 160 | 160 | 160 | 128 | 96 | 96 | 96 | 96 |
| 3 | 160 | 160 | 160 | 128 | 128 | 128 | 96 | 96 | 96 |
| 4 | 160 | 160 | 160 | 128 | 128 | 128 | 96 | 96 | 96 |

These outputs establish this finite centre-coordinate population. They do
not establish the precise sector boundary for arbitrary fractions, different
footprints, wrapping coordinates or floating-point states.

**Unknown.** Native facing reachability, native actor arrangement, and other
acquisition or group scorers, which are separate routines.

**Amended.** The closing statement that the 81-cell table does not establish the sector boundary for arbitrary fractions, different footprints, wrapping coordinates or floating-point states is narrowed by `AI-444`: the law of `R0051` over fine points, footprints and 16-bit wrap, replayed on 84,842 inputs, with no floating-point instruction in the bodies (`w01`..`w03`). The bound is that population; which sizes and coordinates occur in play stays Unknown (`AI-444`). The 81 outputs and the score law stand.

### AI-362

The selector reads manager+8 as head, node+0 as next and node+8 as actor at
L01790..L01791. The original server constructor stores the manager at
L00240 at L00243. The selected map-load call passes it to the world
constructor at L00241/L00242; L00238 binds it at world+0xa4554.
The manager constructor initializes the list at manager+4. The controlled
construction returns head, tail and count zero and block size 10.

R0411 calls the original tail append before allocating actor+4. The
selected type-6 creation walk calls that wrapper at L01792. R0412
links the previous tail's next pointer, or the empty head, then stores the
new tail. R0413 unlinks a selected node and keeps every survivor's relative
order. Original replay gives A,B,C; removing B gives A,C; reappending B gives
A,C,B. Selected leave/return paths R0121, R0122 and R0123 remove or
append, and selected death teardown appends to a different dead-list manager.

**Confidence.** High for these local pointer relations and the selected
manager binding. Medium for their composition into a complete session order.
This experiment reads selected complete bodies and preserves their outgoing
direct and indirect calls. It does not rerun an image-wide typed mutation
census or claim that every creation, alias, bulk copy or indirect lifecycle
path has been excluded. Synthetic node allocation supplies storage, not a
native pool/free-list chronology. MOVE-TICK-013 and MOVE-TICK-015 retain their
separately established scope.

**Unknown.** Full-game party, spawn, summon, event and return chronology for
an arbitrary session. A native controlled creation/return sequence followed
by a list observation would discriminate an external reorder.

### AI-363

The selected world-load rebuild R0414 walks players, then each player's
manager at +0x20. It skips actor+0x4c bit 8 and calls R0032 on L00240 at
L01793. Player load R0415 builds that per-player manager from the player's
group list and each group's member list, appending at L01794. The selected
actor-list and group serializers retain file order through head-to-tail
iteration and tail append. These are the directly traced load relations, not
an actor-ID sort or a claim that a save preserves the old global manager order.

**Confidence.** Medium. The selected whole bodies, their starts, section
mappings and relative transfers are bound to the installed original. Their
local iteration and append relations are direct source evidence. Complete
archive-object creation, fixups, callbacks and native post-load interleaving
are not replayed in this experiment. Earlier MOVE-TICK and SAV claims are
navigation and supporting authorities, not a second code population.

**Unknown.** The final native traversal after an arbitrary load and any
untraced post-load reorder. A controlled save whose global actor order differs
from player/group/member order, followed by a native list observation, would
discriminate the remaining alternative.

### AI-364

The group-order-0 caller R0023 calls R0008 at L00284 without testing
ord+9 or actor+0x58. Its original state table maps 0x16 to L01795, which
calls withdrawal R0100 at L00207. The separate selected global pass
R0148 also calls the actor-policy routine without a progress test; this
does not prove that pass occurs in ordinary play.

The actor tick first takes its HP branch. On its positive-HP route, getter
R0416 returns actor+0x3c == 0, and L01796 skips the executor when that
return is nonzero. With actor+0x3c nonzero, L00086 calls R0016 before
reading action and dispatching its body at L01797. The executor first reads
ord+9. For progress 1/2 it restores attack/cast action, increments the byte
counter and clears progress only when the incremented counter exceeds 2
and actor+0x136 is nonzero. The clear jumps to the common tail, not back to
pending dispatch. Progress 3's cell-centre clear and progress 4's released
status hold also jump to that tail. The hold arm retains action 0x1a.

If an invocation enters with progress already zero, its pending switch has no
retained-phase gate, but its entry activity gate still applies. When controller
+0xb388 and group AI+0x45 are both zero, zero progress returns early with action
0x1a. Nonzero progress bypasses that refusal, so a completion-clear invocation
can end at zero and the following invocation can newly refuse pending dispatch.
Setting group AI+0x45 admits the zero-progress executor with controller+0xb388
still zero. If an invocation itself clears nonzero progress, a quiet
tail leaves action zero for progress 1/2/3 and the next entered-zero invocation
can dispatch the pending order. The reproduced recovery sequence has body
phase 7->0 and completion 1 at T0; progress/action clear at T0+1; pending move
dispatch at T0+2. The pending flee order was already written while progress
was nonzero; another policy evaluation after clearing is not required for
that order to be present. An entered-zero invocation with retained phase 5
dispatches pending move immediately at the executor boundary. These two
meanings of 'progress returns to zero' must not be conflated.

**Confidence.** High for the named caller, table, tests and jumps. Medium
for the composed private sequence: occupancy, allocation, profile imports,
automatic-tail predicates and movement services have explicit hooks. The
recovery action region executes original instructions; full native tick
cadence, physical movement and application are not supplied by the hooks.
Unknown for an unconditional native first-tick or wall-clock promise. The
AI-CLOCK-080 caller alternative and untraced external writers remain live.

The selected R0147 route evaluates group policy at pre-increment server
counter modulo 16 equal 6, then invokes R0193, which increments the counter
and walks actor ticks. R0076 admits a group on controller+0xb388 or group
AI+0x45. Other selected wrappers L01798, R0148 and R0038 preserve
different caller boundaries. No universal frequency is inferred from the
first caller alone.

**Unknown.** Native command-versus-group-versus-actor timing; inactive group,
disabled executor and outside-writer occurrence; arbitrary cycle completion,
counter wrap and spell effects. Observe ord+9, actor+0x58/+0x136, both dispatch
entries and mover+0x98 around a separately authorized native Retreat sequence
to discriminate the remaining timing alternatives.

### AI-365

Every selected progress arm reaches the executor common tail at L00010.
Nonzero mover+0x98 is consumed before the actor body. State 1 has teardown,
0xa has patrol advance and 0x17 writes progress 0xff. State 0x16 takes the
remaining arm at L00114/L00011 and calls R0004. Its successful
in-position branch calls R0005, whose original leaf stores progress 1,
counter 0, action 3 and the newly selected active victim. It does not store
actor+0x58. The controlled composed case clears old progress and then
reinstalls strike progress on the same invocation, retaining phase 5.

**Confidence.** High for this bounded original path and private replay with
an explicit successful position result. The quiet-tail control ends at zero
and does not enter pending dispatch on that invocation. Unknown for native
reachability of the route-failure prerequisite during a retained cycle.

**Unknown.** Whether that prerequisite occurs during a native retained
strike/cast, and its actual victim or application effect. A native trace of
mover+0x98 and the actor's phase before the common tail would discriminate it.

## Script group-id resolution

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-366 | A script `Target_Group` id resolves through an id-only map built over players and their groups head to tail; a repeated id keeps the last group inserted, so the later owner's group wins and the earlier one is unreachable by id. | High / Medium / Unknown | ● active | [EXP-0438](../experiments/EXP-0438-resolver-regen-card/) |
| AI-367 | Opcode-6 sub-command 5 (`Swarm 2`) writes order and destination to the members of the one group its parameter resolved to and to no other group; in mission 70 that is owner 7's 10 records of group id 40, never owner 5's 5. | High / Medium | ● active | [EXP-0438](../experiments/EXP-0438-resolver-regen-card/) |

### AI-366

- The binder `R0067` builds its lookup maps before it resolves any
  script parameter. It loads the player list `[L00380]` (`L01799`) and
  walks it with the first/next iterators `L01800` and `L01801`
  (`L01802`, `L01803`). For each player it stores `playerMap[Player+0x08]`
  (`L01804`, `L01805`..`L00175`), then walks the player's group list
  `Player+0x24` with `L01806` and `L01807` (`L01808`..`L01809`).
  Both lists are appended at the tail: players by `R0417` through
  `R0418` (`L01810`), groups by `R0419` through `R0198`
  (`L00295`).
- The group map is filled by the address of frame local `-0x108` (`L01811`), the call to `R0420` (`operator[]`, `L01812`) and a store through the returned reference (`L01813`), with
  key `group+0x1c`, the group id (`AI-GROUP-009`). `R0420` calls the lookup
  `L01814`; only when it finds nothing does it allocate and link an entry
  (`L01815`..`L01816`); in both cases it returns the entry's value slot
  (`L01817`..`L01818`). The store at `L01813` therefore overwrites an
  existing entry. No test of an earlier value precedes it. The match test is
  the id alone; the owner is not part of the key.
- Resolution is `R0421(map, id)`: `R0420` again, then the entry's
  value (`L01819`..`L01820`); a null value formats the message literal
  at `L01821` with the id and passes it to `R0422` (`L01822`..`L01823`).
  The node-binding arm `L01824`..`L01825` calls it with the group parameter
  `[rec+0x48+4k]`, stores the result at node `+0x34` for the first group
  parameter and `+0x3c` for the second (`L01826`, `L01827`), and writes
  `1` to `(group+0x3c)+0x48` of the result (`L01828`, `L01829`).
- Order: the player list follows the type-5 array. `R0423` loops index
  1..n over the type-5 records (`L01830`..`L01831`) and appends one
  Player per record (the call to `R0424` at `L01832`); the Player's id at `+0x08`
  is the record's `+0x04` (`L01833`..`L01834`). The type-6 spawner
  resolves an owner byte to a Player with `R0425` (`L01835`,
  `L01836`), which compares `Player+0x08`, and appends the new group to that
  Player's list when no group of the same id exists there (`L00292`..`L00296`).
  A group is therefore one (owner, id) pair, and the map holds one pair per id.
- Mission 70 (`tools/groupresolve`, both roots): the type-5 records are
  index 1 Self, 2 Beists, 3 Nocturnal, 4 Villagers, 5 Monsters, 6 Friends,
  7 Enemis. Id 40 is carried by owner 5 (5 records) and owner 7 (10 records);
  the replay returns owner 7. Id 7 is carried by owner 2 (1) and owner 3 (2),
  and returns owner 3.
- Corpus population (`tools/groupresolve`, per root): EN reads 38 maps, 10
  standalone (`Beast`, `Cross`, `Forester`, `Horror`, `Islands`, `Kids`,
  `Kids2`, `LuMoir`, `Tomb`, `Waters`) and 28 in `scenario.res`, all 38 with a
  parsed type-7 record. RU reads 34 maps, 6 standalone (`Forester`, `Horror`,
  `Islands`, `Kids`, `LuMoir`, `Waters`) and 28 in `scenario.res`; 33 parse and
  `Horror.alm` is skipped because its type-7 record fails to parse. The four
  EN-only standalone maps are `Beast`, `Cross`, `Kids2` and `Tomb`. Over those
  populations 267 `Target_Group` parameters were scanned on EN and on RU (the
  tool prints the same count and the same 8 hits for both); 8 name an id carried by more than one owner,
  over four (map, id) pairs: `scn:70` id 40 and id 7, `scn:100` id 25 (owners
  3 and 6, returns 6), `scn:140` id 6 (owners 2 and 3, returns 3). Four of the
  eight are opcode-6 actions; the other four are condition nodes.

- Resolver paths: `EnumRefs callto:R0421` returns 4 hits over 1 owner, all in
  the binder (`L01837`, `L01838`, `L01839`, `L01840`; 0 orphan). The
  two at `L01839` and `L01840` are in a binder arm that was not read.

**Confidence.** High for the mechanism: the build loop, the overwriting store,
the id-only key and the single-pointer lookup are cited instructions, and they
exclude first-match, all-match and (owner, id) matching. Medium that no second
resolver path exists: it rests on the group map being local to the binder's
stack frame and on the four call sites of `R0421` all lying in the
binder; two of them were not read. High for the tail
appends. Medium for the corpus replay, which reads the type-5 and type-6
records with `tools/groupresolve` and does not observe the binder. Medium that
the condition nodes resolve through the same arm: they were counted, not read.

**Unknown.** Whether a SAV restore re-creates the player list and each group
list in the same order as the map load (the binder also runs from the restore
arm at `L01841`, whose list contents were not read). A save whose owners
carry one shared id, restored and then commanded, would discriminate it.

### AI-367

- Sub-dispatch `R0156` indexes the table at `L00408` by
  `rec+0x08 - 1` (`L01842`..`L01843`). Parameter 5 is entry 4, `L01844`:
  it loads the byte at `rec+0x10`, the byte at `rec+0x0c` and the pointer at
  `rec+0x34`, and calls `R0157(group, +0x0c, +0x10)` (`L01845`). The
  pointer is the group `R0421` stored (`AI-366`).
- `R0157` reads the member list of the group it was given (count at
  `+0x0c`, head at `+0x04`, `L01846`..`L01847`) and visits each member
  once through that list: it writes `member+0x158` fields (`L01848`..`L01849`)
  and advances through `R0068` and `R0069` (`L01850`..`L01851`). It
  names no other group and reads no player list.
- Mission 70: the two instants (start trigger node 25 and trigger 0 node 2)
  carry sub-command 5 with group id 40 and resolve to owner 7's group of 10
  records, uids 135..144. Owner 5's group, uids 89..92 and 96, is not
  addressed by either instant.

- `EnumRefs callto:R0157` returns 2 hits over 2 owners: `L01845` and
  `L01852` in `R0189`, which was not read.

**Confidence.** High for group-only writes through the sub-command 5 path: the
routine takes one group and iterates only its list. The statement "to no other
group" covers that path; the second caller `R0189` was not enumerated. Medium for the mission 70 outcome, which combines
that with the `AI-366` replay.

**Unknown.** What a later AI tick does to owner 5's group; this pass reads no
AI tick.

## Order setters, tick order and route refusal

Evidence for this section is a static read of `rom.exe` (one image on both lawful installs,
sha256 `942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`). No process was run.
`ord` is the order block at `actor+0x158`; `mover` is the block at `actor+0x154`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-370 | Script group sub-command 3 (Stand Ground) is the arm at `L01853` in `R0156`; it calls neither `R0060` nor `R0063`, stands each member through `R0088` and stores group order 3 itself at `L00315`. | High / Medium | ● active | [EXP-0440](../experiments/EXP-0440-order-pursuit/) |
| AI-371 | In one tick the AI slot and queue drain precede every member's executor pass; a spell application in the actor body can run a setter after it. `R0062`, `R0065`, `R0066`, `R0067` are unplaced. | High / Medium / Unknown | ● active | [EXP-0440](../experiments/EXP-0440-order-pursuit/) |
| AI-372 | In the `+0x158` census the only reads of the order word `ord+0x50` are two compares in `R0154`, the group order 4 arm; if reached, a completed pickup leaves 1 there, which sends that arm to `R0022`, not to a walk. | High / Medium | ● active | [EXP-0440](../experiments/EXP-0440-order-pursuit/) |
| AI-373 | In `R0043` and `R0055` the near search raises `mover+0x98` only when aimed at the final destination; an empty waypoint search retries next pass. Occupied ring cells unlabelled: Unknown. | High / Medium / Unknown | ● active | [EXP-0440](../experiments/EXP-0440-order-pursuit/) |
| AI-374 | In state 3 `R0008` tests only the victim's action word; 0x10, set at health -10 or less, stands the member down to state 0xc before dispatch, so arm 3 is not entered; the prologue reads no list. | High / Medium | ● active | [EXP-0440](../experiments/EXP-0440-order-pursuit/) |
| AI-375 | Move stores `actor+0x50` = 1 and pickup 2; the executor tail reads that value when it consumes `mover+0x98`. A pending walk installed when recovery reaches zero at T0 first calls `R0178` at T0+2; whether it writes a step is Unknown. | High / Medium | ● active (amended) | [EXP-0440](../experiments/EXP-0440-order-pursuit/) |
| AI-394 | In `R0053`'s own body the static list is empty on return only through three exits: start equals goal, picker A answers 0, or it answers the start cell; the extractor's 1000-step cap is a fourth, unexcluded exit. | High / Medium | ● active | [EXP-0489](../experiments/EXP-0489-troll-hints/) |
| AI-395 | The idle-turn arm `R0205` (order 0xb) and the turn routines it calls never store `ord+0x08` or `ord+0x0c` and never call reacquisition or route search, so order 0xb ends by no write of its own. | High / Medium | ● active | [EXP-0489](../experiments/EXP-0489-troll-hints/) |

### AI-370

- `R0156` indexes the table at `L00408` by `rec+0x08 - 1` (`AI-367`). Sub-command 3 is entry 2,
  `L01853`..`L01854`.
- The arm walks the member list once, calling `R0007` per member and storing `ord+0x50 = 0` and
  `ord+0x38 = 0` (`L01855`..`L01856`). It stores `grpAI+0x20 = 0` (`L01857`), walks the list again
  calling `R0088` (`L01858`), and stores `grpAI+0x20 = 3` (`L00315`).
- `R0088` per member stores `ord+0x14`, `ord+0x60 = 0`, `mover+0x7c = 0`, `actor+0x54 = 0`,
  `ord+0x50 = 0`, `actor+0x50 = 0xc` (`L01859`), the post word and `ord+0x08 = 0` (`AI-352`).
- `callto:R0060` returns 2 hits: `L00127` in `R0061` (queued opcode 0x18) and `L00128` in
  `R0062`. `callto:R0063` returns 6 hits over 5 functions: `R0064`, `R0065`,
  `R0066`, `R0067` and `R0003`. None of the 8 sites lies in `R0156`, whose one
  caller is `R0262` at `L01860`.
- `callto:R0088` returns 4 call sites in 3 functions: `R0016` (`L01861`), `R0154`
  (`L01862`) and `R0156` (`L01858`, `L01159`).

**Confidence.** High that the sub-command 3 arm contains no call to either setter and writes group order 3
itself: the arm is read instruction by instruction and the two enumerations list every direct call site in the
image. Medium that no callee of the arm reaches either setter: `R0007` and `R0071` were not
read whole, but the enumerated callers of the two setters exclude them.

**Unknown.** The other `R0156` arms (the table has 17 entries) were not read here.

### AI-371

- Both tick drivers run the AI slot first. `R0147` calls `R0076` at `L00279` and
  `R0193` at `L01863`; `R0075` calls them at `L00163` and `L01864`
  (`callto:R0193` 2 hits, `callto:R0076` 2 hits, `callto:R0427` 1 hit, `callto:R0426` 1 hit).
- `R0193` increments the tick counter, drains the command queue through `R0191` and
  `R0061` (`L01865`), calls `R0426` (`L01866`), then `R0427` (`L01867`), which
  calls each actor's slot through the virtual call through the vtable slot at `+0x18` (`L01868`). The AI slot runs on a phase test of the pre-increment tick counter, `& 0xf == 6`. The slot `R0037` calls the executor `R0016`
  at `L00086` unless `R0416` returns nonzero (`L01796`), and dispatches on `actor+0x54` after it
  (`L00538`).
- Script origin: `R0076` reaches `R0428` (`L01869`), `R0262` (`L01870`) and
  `R0156` (`L01860`); `R0262` also reaches `R0064` (`L01871`, `L01872`), which
  calls `R0063` (`L01873`). Queue origin: opcode 0x18 calls `R0060` at `L00127` (dispatch
  table `L00529`, index opcode - 0x14 = 4).
- Actor-body origin: the arm at `L00008` calls `R0003`, which calls `R0063` at `L01874`.
  That call lies after `L00086` in the same slot invocation.
- Pickup (`AI-356`): row 7 at progress 0 writes action 2 on its first pass (`L00026`) and the actor body
  for action 2 runs in that slot invocation. The completion pass is the member's next executor pass, which
  keeps action 2 (`L01875`..`L00540`). The queue drain and the AI slot of that next tick run before it.
  A Stand Ground setter reaching the member in between writes action 0 and `ord+0x08 = 0` (`AI-352`), so
  row 7 is not reached.
- Not placed in the tick: `R0062` (called from `R0429`), `R0065` (3 callers),
  `R0066` (called from `R0430` and `R0061`), `R0067` (called from `R0414`
  and `R0128`) and the other callers of `R0262`, `R0431` and `R0432`.
  `R0016` has a second caller, `R0038` (`L00087`), which has no direct caller.

**Confidence.** High for the order of the AI slot, the queue drain and the executor within one tick, and for
the actor-body origin following the executor. Medium that a setter running between a pickup's two passes
prevents completion: the writes are cited and the pass order is read from code, but no run was observed.

**Unknown.** The tick position of the unplaced callers above, which spell arm reaches `L00008`, whether
the sack transfer already finished on the first pickup pass, and the AI slot period beyond its phase test.

### AI-372

- Census instrument: `OrderBlockFieldCensus` over every instruction of the image (inside functions and
  orphan runs) with displacement 0x50 and the base register loaded from `[x + 0x158]` within 12 instructions
  (class ORD). It finds 91 ORD hits, 89 writes and 2 reads, and 422 other-base hits. Displacements 0x4d,
  0x4e, 0x4f, 0x51, 0x52 and 0x53 give 0 ORD hits.
- The two ORD reads are `L01876` and `L00455` in `R0154`, called from `R0023` at
  `L01877` (group-order table `L00303`, entry 4) and, as its second caller, from `R0026` at
  `L00463`. The first runs
  the stand-down sequence for a member at its queued cell and cell centre only when `ord+0x50` is 0. The
  second chooses `R0022` when the word is nonzero and `ord+0x08 = 1` with the queued cell kept when it
  is 0.
- A completed pickup stores `ord+0x50 = 1` (`L00029`, `AI-356`). Stand-down sequences store 0 and then 1
  the same way (`L01878`, `L01879`). `R0022` is the state 0xc acquire routine (`AI-CMD-054`).
- The other-base reads in functions that load `+0x158` include actor-state reads (`R0008` four,
  `R0016` one at `L01880`) and stack slots; `R0053`, `R0433` and `R0262` hit
  `[ESP + 0x50]`.

**Confidence.** High for the two reads and their branches. Medium that no other routine reads the word:
the census does not see indexed operands, a base reached through a copy chain longer than 12 instructions,
a call return or a stack slot, a bulk copy such as the serializer `R0199`, or raw bytes Ghidra never
disassembled.

**Unknown.** Readers through those blind spots; the 5 other-base reads in `R0434` (`L01881`,
`L01882`, `L01883`, `L01884`, `L01885`) and the orphan runs `L01886`, `L01887`, `L01888` and
`L01889`, which load `+0x158` and were not classified; and whether the group order 4 arm is reached for a
member whose pickup just completed.

### AI-373

- `R0043(actor, victim, stop)` returns early when the position offset bytes are not both 0x80
  (`L01890`..`L01891`) and stands and turns when the edge distance is within `stop`
  (`L00580`..`L00754`). Otherwise it increments `mover+0x78`. A changed `mover+0x7c` clears the static
  list and sets `mover+0x09 = 0xff` (`L01892`..`L01893`).
- It runs the full search `R0053(..., flag 1)` (`L01894`) when `mover+0x09` exceeds count/3 + 1
  (`L01895`) and the previous length `mover+0x8a` exceeds `ctx+0x585c8` (`L01896`); with a shorter
  previous length it rebuilds the list as the single destination cell (`L01897`..`L01898`). It sets
  `mover+0x98 = 1` when the static list is empty afterwards (`L00161`), then stores `mover+0x7c = victim`
  and `mover+0x09 = 0` (`L01899`..`L01900`).
- Each later pass runs the near search `R0055` (`L01901`) when the dynamic list is empty
  (`L01902`) or `mover+0x78` exceeds `ctx+0x585c0`, and `mover+0x00 == mover+0x01` (`L01903`..`L01904`).
  It then stores `mover+0x94`, `mover+0x96`, increments `mover+0x09`, clears `mover+0x78`
  (`L01905`..`L01906`) and calls the stepper `R0054` (`L01907`).
- `R0055` picks a target by the static list count N: the destination `mover+0x76` (kind 1) when N
  is at most `ctx+0x585c8`, the head node (kind 2) when the head is farther than `ctx+0x585c4`, otherwise a
  later node (kind 3). It calls `R0053(..., flag 0)` (`L01908`). It sets `mover+0x98 = 1` only when
  the dynamic list count at `actor+0x184` is 0 and the kind is 1 (`L01909`..`L00162`). When N is nonzero
  and the head is within `ctx+0x585c4` cells it then removes the head node (`L01910`..`L01911`).
- The search labels cells in a word plane initialised to 0xffff (`L01912`). A goal label of 0xffff sends
  the full search to Picker A `R0392` (`L01913`), and the flag-0 search to Picker B `R0435`
  when the victim argument is nonzero (`L01914`), else to Picker A with radius 8 (`L01915`).
  Picker B scans the ring of cells around the victim, up to 8 rings (`L01916`), keeps the cell with the
  smallest label below 0xffff and returns 0 when none has one (`L01917`); on a 0 return the search builds
  no list (`L01918`).
- Composition, read from code: when Picker B returns 0 on a kind-1 near search the dynamic list is empty,
  `mover+0x98` becomes 1 and the tail of the same executor pass consumes it. For `actor+0x50` = 3 that
  stores `ord+0x08 = 0` and calls `R0004`, whose stores `AI-350` lists (`AI-ROUTE-045`); the next pass
  runs whatever order it stored. A kind-2 or kind-3 search that returns empty leaves the flag clear. The next
  pass reaches `L01902` with the dynamic count 0, so the refresh counter is not tested; with the facing
  settled it runs the near search again, or the full search once `mover+0x09`, incremented at `L01919`,
  exceeds count/3 + 1. With the facing unsettled it runs only the stepper.

**Confidence.** High for the near-search statement and the control flow cited above, read in
`R0043` and `R0055` and as branch structure in `R0053` and `R0435`. Medium for
"only" in the headline: it covers those two routines, and `R0178` stores `mover+0x98 = 1` at
`L01720`, a third setter not analysed here. Medium for the composed outcome of a fully occupied ring.
Unknown for whether a cell holding another unit carries label 0xffff.

**Unknown.** How unit occupancy enters the label plane (the plane is tested against the mover byte
`mover+0x05` at `L01920` and similar sites; no unit-occupancy write was traced), what `R0054` does
with an empty dynamic list, what `R0436` computes, and the values of `ctx+0x585c0`, `ctx+0x585c4`
and `ctx+0x585c8`.

### AI-374

- `R0008` reads `actor+0x50` first (`L01148`). For 3 it loads `ord+0x0c`; if that is nonzero and
  `[ord+0x0c + 0x54] == 0x10` (`L01154`) it runs the stand-down: three `R0007` calls,
  `ord+0x50 = 0`, `actor+0x50 = 0xc` (`L00597`), `ord+0x00` = own cell word, `ord+0x08 = 0`
  (`L01921`) and `ord+0x50 = 1` (`L01878`). It stores nothing to `ord+0x0c`.
- It then re-reads `actor+0x50` (`L01922`) and dispatches through the table at `L00322` with 0xc, not
  entry 3 (`L01923`, `R0009`).
- Action 0x10 is stored by `R0208` (`L00609`), called from `R0037` at `L01924` when health
  is -10 or less (`L01925`), and directly at `L01710`. In the same pass `R0427` removes the actor from the world list
  (`L01926`) and adds it to the dead list (`L01927`). The state-3 prologue reads the victim object
  through the stored pointer and reads neither list.
- `R0008` is called from the group order 0 arm of `R0023` (`L00284`, `AI-355`), whose only
  caller is `R0076` (`L00282`). The stand-down therefore happens at the first AI slot after the
  victim's action word is 0x10.

**Confidence.** High for the prologue, its stores and the dispatch. Medium for the timing, which composes the
slot placement of `AI-371` with the cited health test and does not enumerate other callers of
`R0008` (no `callto:R0008` was run). The question's "state-3 arm" is read as `actor+0x50 == 3`
in `R0008`.

**Unknown.** Arm 3 itself (`R0009` and its callees) was not read for list or liveness tests. A victim
that leaves the world list by a route other than action 0x10 is unanswered. What a victim object that is freed before the next AI slot holds at `+0x54`; this read does not
show when the dead list releases it.

### AI-375

- Move, queued opcode 0x16, calls `R0108` (`L01928`, table `L00529` entry 2). Per member
  `R0301` stores `actor+0x54 = 0`, `ord+0x50 = 0` (`L01929`), `ord+0x08 = 1` (`L01930`), the
  queued cell `ord+0x0a`, `mover+0x90 = 0` and `actor+0x50 = 1` (`L00603`, from the constant 1 loaded at `L00604`).
- Pickup `R0010` stores `actor+0x50 = 2` (`L00022`, `AI-356`).
- `R0016` stores `actor+0x50` only at `L00027` (pickup completion, 0xc) and `L00615` (teardown,
  0xc, after the consumption test). The tail reads `actor+0x50` at `L01880`: 1 tears down to 0xc, 0xa and
  0x17 have their own arms, and any other value, 2 included, stores `ord+0x08 = 0` and calls `R0004`
  (`AI-350`).
- The census (`OrderBlockFieldCensus`, displacement 0x50) lists in `R0008` four stores to
  `actor+0x50`: `L00597`, `L00598` and `L00599` of 0xc and `L00606` of 4. That routine runs at the AI
  slot (`AI-374`).
- Walk timing: the executor runs before the actor body in `R0037` (`L00086`, then `L00538`).
  The body applies recovery zero on tick T0 after the executor ran. At T0+1 the strike or cast progress arm
  counts `ord+0x15` and clears progress and action only when `ord+0x15 > 2` (`L00094`, `JBE`) and
  `actor+0x136` is set (`L01931`..`L00095`), then ends at the tail. At T0+2 progress is 0, row 1 runs and calls `R0178` (`L01932`) (`AI-350`, `AI-354`,
  `AI-355`).

**Confidence.** High for the Move and pickup stores and the executor's own stores. Medium for the value at
consumption: it is the setter's value unless an AI-slot evaluation or another writer outside the census
population intervened. Medium for T0+2 as the tick of the first `R0178` call, which composes cited gates and
assumes the pending walk is already installed.

**Unknown.** That `ord+0x15 > 2` and `actor+0x136` hold at T0+1 was not shown. Whether the first row 1 invocation writes a position step or only turns when `mover+0x00`
differs from `mover+0x01` (`R0056` is reached from `R0178` at `L01933` and `L01934`),
writers of `actor+0x50` outside the census population, and native reachability of the sequence.

**Amended.** `MOVE-090` reads the first walk call: it writes a step in that call only when the facing byte already equals the direction to the first node, and otherwise turns, with the step at the next call. `AI-382` places the pickup transfer in the first pass.

### AI-394

- `R0053` has five returns (`L01935`, `L01936`, `L01937`, `L01938`, `L01939`). It calls `R0437` (`L01940`, `L01941`), `R0438` (`L01942`, `L01943`), `R0439` (`L01944`) and once `R0058` (`L01945`, not read). Its only direct stores to the mover's list fields `+0x160`..`+0x170` are the clear at `L01946`..`L01947`, entered when the static flag is nonzero (`L01948`, `L01949`).
- With the static flag set the list is therefore empty on return, by the instructions of this body, in three cases. (a) Start equals goal: `L01950` jumps to `L01951` after the clear. (b) The goal is unlabelled and picker A answers 0: `L01952`..`L01944` frees the list. (c) Picker A answers the start cell: the extractor adds no node for an endpoint equal to its seed (`MOVE-080`). A labelled goal returns through `L01935` or `L01936` with a start that differs from the goal, and its list is then the extractor's. The extractor `R0437` counts trace steps and tests `cmp si,0x3e8` (`L01953`..`L01954`, again at `L01955`..`L01956`); on `ja` it jumps to `L01957` and the teardown at `L01958` stores count 0. That cap is a fourth exit that this body's enumeration does not exclude. `R0043` tests the count at `L01959` and stores `mover+0x98 = 1` at `L00161` (`AI-373`).
- `D` is the larger of the coordinate differences of start and goal, stored as a word at `esp+0x28` (`L01960`..`L01961`). The static tail loads it (`L01962`), shifts right by 2 and adds 4 (`L01963`, `L01964`) and passes the result as picker A's limit (`L01913`). Picker A scans rings 1 through limit-1 around the goal and stops at the first ring holding a labelled cell (`MOVE-ALT-019`): radius `(D>>2)+3`.
- The search stores label 0 on the start cell (`L01965`). Picker A reads the same word array at `+0x30000` with index `y*256+x` (`L01966`..`L01965`, `L01967`..`L01968`). The start cell lies on ring `D` around the goal, inside the scanned rings only for `D` of 1 to 4; `D` of 5 scans radius 4.
- Composition: for `D` of 1 to 4 an unlabelled goal never gets answer 0, because ring `D` holds the start cell with the strict minimum label. The answer is a labelled cell of a nearer ring when rings 1 to `D-1` hold one, else the start cell, which is case (c). For `D` of 5 or more the start cell is outside the scan and case (b) occurs when no ring up to `(D>>2)+3` holds a label. For rings inside the map, an unlabelled goal leaves the list empty in cases (b) and (c) exactly when no cell within Chebyshev distance `min(D-1, (D>>2)+3)` of it is labelled. Picker A masks each probe `(y<<8)+x` with `0xffff` and has no bounds test (`L01967`, `L01969`, `L01970`, `L01971`), so probes with `x` outside 0..255 read the neighbouring row and probes with `y` outside the map wrap the plane. Case (a) is separate.
- Scope: the static arm of `R0043`. `R0178` stores the flag at `L01720` after the same tail and was not re-read. The near search's flag at `L00162` follows the dynamic arm, with picker A at bound 8 or picker B (`AI-373`, `MOVE-ALT-018`), and is not covered.

**Confidence.** High for the return and call enumeration, the list-field stores, the `D` computation, the limit arithmetic and the start label: each is a cited instruction in a body read whole, with instruction counts and hashes in the evidence. Medium for "only" (the extractor's cap and the unread `R0058` are outside the enumeration) and for the composed rule, which rests on the ring scan of `MOVE-ALT-019` and on the extractor result of `MOVE-080`, a private replay and not a native run.

**Unknown.** Whether the search budget keeps a trace at or below 1000 steps, so that the extractor's cap never empties a labelled goal's list. What picker A returns near the map border, where probes alias. How often play produces a goal with no labelled cell inside the radius. A native trace of `L00161` and `L00162` with the order-kind transitions would settle it. What leaves a passable cell unlabelled beyond the generation budget (`MOVE-SEARCH-001`) is not read here.

### AI-395

- `R0205` (52 instructions) stores `actor+0x54` = 0 (`L01972`), the wanted facing byte `mover+0x01` (`L00997`), `ord+0x54` = 0 (`L00657`) and `actor+0x54` = 1 (`L01973`). It calls `R0056` (`L00751`) and the random generator `R0179` (`L01974`, `L01975`) and nothing else.
- `R0056` (84 instructions) stores `mover+0x01` (`L00748`), `mover+0x00` (`L01976`), `mover+0xa4` (`L01977`, `L01978`), `mover+0xa0` (`L01979`, `L01980`) and `mover+0x9d` (`L01981`). Its only call is `R0249` (47 instructions, no call), which stores `mover+0x00` alone. `R0250` (17 instructions) stores `actor+0x54` = 1 (`L01982`) and calls `R0056`. `R0179` (14 instructions) stores the generator state (`L01983`) and calls `R0228`, which was not followed.
- No store in these five bodies has displacement `+8` or `+0xc` through a register other than `esp` or `ebp`, and none of them calls `R0004` or `R0053`. The arm and the turn routines change facing and action state only.
- Inference: the idle-turn arm ends order 0xb by no write of its own. An actor that a refusal left in order 0xb (`AI-ROUTE-045`) therefore keeps it until a routine outside these bodies stores `ord+0x08`: the executor stores it only as 0 in its own body (`AI-350`), and a group member's arm can reissue a pursuit on every AI tick (`AI-REISSUE-077`).

**Confidence.** High for the call and store enumeration over the five bodies, each read whole with instruction counts and hashes in the evidence. Medium for the persistence inference, which depends on the writers of `AI-350` and `AI-REISSUE-077`.

**Unknown.** What `R0228` does. What the stores to `mover+0xa0`, `mover+0xa4` and `mover+0x9d` mean for a later step. The entry gate and the facing range are `AI-TURN-104`'s.

## Slot draw: entry, gate and edge cases

Evidence is a static read of `rom.exe` (identical on both lawful installs; capstone disassembly, direct `CALL rel32` cross-references and displacement sweeps, indirect calls not followed). Every instruction quoted is asserted on both editions (`evidence/asserts-en.txt`, `evidence/asserts-ru.txt`: 170 rows, 0 mismatches). The next free `ai.md` id is `AI-378`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-376 | The engage selector reads no cast-state field of its own actor, so a pending cast does not stop a draw; a miss writes order kind 5 over the retained cast order, while the in-flight cast runs from actor fields. | High / Medium | ✔ promoted (amended) | [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md) |
| AI-377 | The slot draw skips a zero slot id before `rand()`, never matches a zero threshold, and has no owner test in its body; only a mage-bit actor with reach below 2 and `Player+0x28 == 0` is diverted. | High / Medium | ✔ promoted | [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md) |
| AI-378 | The Alt+letter record's receiver is the debug console `R0440`, which acts only for a Player with `+0x68` above `0x32`; six keys act (D, H, I, Q, T, U), 17 do nothing; `#Chicken` in chat sets `+0x68` to `0xff`. | High / Medium / Unknown | ● active (amended) | [EXP-0450](../experiments/EXP-0450-keyboard-remainder/) |

### AI-376

Instrument: a displacement sweep (`evidence/fieldreads-engage.txt`) for `actor+0x54`, `+0x58`, `+0x5c`, `+0x64` and `+0x6c` over `R0009`, `R0209`, `R0018`, `R0137`, `R0023`, `R0024` and `R0441`, plus `R0008` (`R0008`..`L00322`). The only reads of `+0x54` are `L01154` and `L01984`, which compare another actor's word (the target's and the partner's) with `0x10` (`AI-374`). The only writes to the member's own `+0x54` are `L01985`, `L01986`, `L00321` and, in `R0024`, `L01987`. No instruction in `R0009` tests the caster's cast state, the cast timer or the progress byte `ord+9`, so the selector is reached and draws whether or not the creature's own cast is in flight.

`R0009` has 14 direct call sites. Command state 3 reaches it at `L01988` with `ord+0xc` untested, state 4 at `L01989` after a target search, and states 8, 10, 11 and 17 through `R0306`, `R0137` and `R0139`; the group handlers reach it at `L00842`, `L00336`, `L01990`, `L01708` and `L01991`, and `R0393` at `L01992`. State 13 and state 14 re-arm kind 8 and kind 9 through `R0018` instead (`L01993`). The member executor path `R0076` gates on `[this+0xb388]` or `[group+0x3c]+0x45`.

A draw that matches no slot writes the default engage order, `ord+8 = 5`, `ord+0xc = target`, `ord+0x14 = actor+0x12c` (`L00018`, `L00019`, `L00159`). The cast fields `ord+0x28`, `ord+0x30` and `ord+0x60` are not written on that path, so `ord+0x60` keeps the retention value of the earlier cast order (`MAGIC-240`). The cast in flight does not read the order block: the install copies the target and spell into `actor+0x5c` and `actor+0x64` (`MAGIC-239`), and the unit tick reads those fields (`L01994`, `L01995`). A replaced order therefore does not cancel a cast that has already started.

**Confidence.** High for the absence of a cast-state read in the selector and in the routines swept, and for the default engage writes (complete-body reads). Medium for "does not stop a draw" as a runtime statement: the entry gates of the 14 callers were read only partly, and a gate outside the swept routines could still keep a casting actor from reaching the selector.

**Unknown.** The cadence of the AI-slot driver `R0075` and the gate on its call to `R0076` at `L00163` were not established; `AI-371` places the slot before the executor passes. Whether the apply routine `R0003` reads the order block was not read. The unguarded call site `L01988` was not traced to the writer of `ord+0xc`.

**Amended.** `AI-385` reads the cadence: `R0076` runs once per 16 ticks on both of its callers, `R0147` and the thread loop `R0075`.

**Evidence.** [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md), `evidence/fieldreads-engage.txt`, `evidence/xref-engage.txt`, `evidence/asserts-en.txt`

### AI-377

Gate (`L01996`..`L01997`): the selector tests the mage bit of the actor, `actor+0x4c & 4` (`R0442`). A non-mage actor goes straight to the slot loop. A mage-bit actor with weapon reach `actor+0x12c < 2` and `Player+0x28 == 0` skips the draw and calls `R0022`, the standing acquisition, and returns. A mage-bit actor with reach 2 or more, or with `Player+0x28 != 0`, draws. No instruction in the gate reads a spell, a slot or a mana value, and none tests the owner of a non-mage actor, so the selector body has no owner test. Whether an owner-0 creature reaches the selector through its callers is not established.

Loop (`L00610`..`L00611`): a slot id of 0 at `[esi+edi]` is skipped before `rand()` (`R0179`). An order block whose three ids and thresholds are all zero therefore makes 0 `rand()` calls, leaves `EBX = 0` and falls to the default engage write (`AI-376`). A nonzero id with threshold 0 makes one `rand()` call and never matches, because the signed compare `JGE` skips a draw that is not below the threshold. `rand()` returns at most `0x7fff` (`AI-RAND-058`), so a threshold above `0x7fff` would match every draw. A match whose `Spellbook::Get` is null writes nothing, including no default engage order (`MAGIC-221`), so the order block is left as it was.

Cast admission: `R0268` refuses only a caster with the mage bit, no item in `caster+0x68`, and cost above `caster+0x9a` (`L01998`..`L01999`; `MAGIC-CAST-003`). A non-mage creature is never refused or charged on that test. The body of `R0268` has no range, facing or ownership test (`MAGIC-239`, which lists the callees and callers not read). Within the selector, the predicate that admits a non-mage creature to cast from its slots is the draw: a nonzero slot id, a draw below its threshold and a non-null spell row.

**Confidence.** High for the gate, the loop, the zero and threshold cases and the body of `R0268` (complete-body reads, instruction assertions on both roots). Medium for the meaning of `Player+0x28 == 0`, which this ledger reads as a human participant: `UNIT-OWNER-009` is partly retracted and no reader of the field was enumerated here.

**Unknown.** Whether an owner-0 creature reaches the selector through the callers: the gates in `R0076` (`[this+0xb388]`, `[group+0x3c]+0x45`), the `R0075` cadence and `L01988` were not resolved. The reach byte `actor+0x12c` of Dragon and Daemon classes was not read, so whether any shipped mage-bit creature meets the reach test is open. Whether a shipped creature class carries an all-zero slot block was not counted here.

**Evidence.** [EXP-0444](../experiments/EXP-0444-creature-books/EXP-0444.md), `evidence/asserts-en.txt`, `evidence/listing.txt`

### AI-378

- Receiver: opcode `0x46`, sub-selector `0x80` of `R0061` calls `R0440` at `L02000` (the `0x80` arm of `SESS-PARAM-017`; `EnumRefs callto:R0440`: 1 hit). The arm tests `CMP byte [ESI+0x68],0x32` / `JBE` on its Player argument first (`L02001`), then the index minus 3 against `0x11` (`L02002`), the byte table `L02003` and the target table `L02004`.
- Keys that act, by letter: D `L00397` toggles the turn-tracing flag `[[L00285]+0x118]+0` and prints `Turn tracing turned on.` or `turned off.`; H `L02005` prints a six-line help (`<Alt-h>`, `<Alt-q>`, `<Alt-t>`, `<Alt-i>`, `<Alt-d>`, `<Alt-u>`); I `L02006` calls `R0399`, which prints `Last Turn Statistics:` and `Average Turn Statistics:` blocks (turn, segment, active, AI, script and activating counts); Q `L02007` (in-game label "Safe mode") toggles `[session+0xb388]` and prints `Safe mode turned on.` or `off.`; T `L00396` toggles the script-tracing flag at `+4` of the same block and prints `Script tracing turned on.` or `off.`; U `L02008` calls `R0400`, which prints `Mission units stats:` with per-group counts and experience. The other 17 keys of B..Y except S reach `L02009` and return.
- Output: each line goes through `R0443` on `L00522`, which builds a record with `+9 = 0x91`, the chat record of `SESS-CMD-016`, and sends it with `R0217`. Where a client shows such a line was not read.
- Gameplay effect, an inference through `AI-ARBITER-057` and `AI-350` (the console's own label is "Safe mode"): Q is the AI admission override; with `[+0xb388]` set, `R0076` dispatches every group and `R0016` no longer parks idle actors with `actor+0x54 = 0x1a` (`AI-350`). `EnumRefs disp:b388` (`rom-enum-disp.txt`): readers `R0016` and `R0076`, writers `R0124`, `R0077` and `R0440`. D and T flip two trace flags and print. H, I and U: no stores found other than string builds and stack arrays; one virtual call at `L02010` (`[EDX+0x30]`, in `R0400`) is unread.
- Gate: `R0444` is the test `+0x68 > 0x32`. Writers of the byte in the displacement-store census (`EnumRefs disp:68`, `rom-enum-disp.txt`, and `callto:R0445`; other operand forms are not excluded): `R0445` has 2 call sites, both in the chat handler `R0430`: `0xff` at `L02011` after the line matches the string at `L02012` (`#Chicken`), and `0` at `L02013` for each other privileged Player after the sender passes `R0444`; the other store is `0` at `L02014`, in `R0201`, taken here to be the Player constructor (not shown). If so, a fresh Player holds 0 and the console refuses it.

**Confidence.** High for the receiver, the gate, the two tables and each arm's flag write or print (read whole, strings dumped). Medium for how `R0430` compares the typed line to `#Chicken` (whole line or prefix), for the demotion arm at `L02013`, for `R0201` being the constructor, and for the AI effect of Q (an inference).

**Unknown.** The readers of the two trace flags, everything `R0399` and `R0400` print beyond the strings named, where a client displays the `0x91` lines, and whether the save loader writes `+0x68`: only `L02014` and `R0445` store the byte by displacement, so a loaded privileged state would need a wider store.

**Amended.** `#Chicken` is a prefix match with no further gate (`MENU-102`), so the Medium comparison statement above is High; `R0201` is the Player constructor and `Player::Serialize` writes no `+0x68` in its own body (`MENU-103`); the console has no participant-flag test and, by the writers found, is reachable only where `#Chicken` acts or for a Player privileged earlier in the process (`AI-397`, Medium); a LOAD leaves the byte at 0 by the constructor route (Medium).

## Attack cycle, pickup pass and pursuit thresholds

Evidence for this section is a static read of `rom.exe` (one image on both lawful installs,
sha256 `942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`) in Ghidra 12.1.2 headless; no
process was run. `ord` is the order block at `actor+0x158`; `mover` is the block at `actor+0x154`. `EXP-0451` was
allocated ids `381`..`386` of `claims/ai.md` and spent `AI-381`..`AI-385`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-381 | The stop routines store `actor+0x54 = 0` and clear neither the strike phase nor the progress byte; the strike continues unless progress is cleared elsewhere (Medium). | High / Medium | ● active | [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md) |
| AI-382 | The first pickup pass stores action 2 and the action 2 arm runs the whole sack transfer in that slot call; the second pass carries only completion stores. | High / Medium | ● active | [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md) |
| AI-383 | Group orders 3 and 5 branch per member on `ord+0x20` and `Player+0x28` without testing `ord+0x08` or `ord+0x0a`: a member with both zero reaches `R0015` (stores 0 or 8, `AI-349`), one with `Player+0x28` nonzero gets `ord+0x08 = 0xb`. | High / Medium | ● active | [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md) |
| AI-384 | In `R0178`, `R0043`, `R0055` and the search routine, no counter gives up: `mover+9` and `mover+0x78` only select a re-search, and `mover+0x98` is set at three empty-list sites. | High / Medium | ● active | [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md) |
| AI-385 | `R0076`, the AI pass that contains the group tail, runs once per 16 ticks on both of its callers: `R0147` when the tick counter and `0xf` equal 6, and the thread loop `R0075` once before each block of 16 ticks. | High / Medium / Unknown | ● active | [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md) |

### AI-381

- Stop routines: `R0088` stores `actor+0x54 = 0`, `actor+0x50 = 0xc`, the post word at `ord+0x00`, `ord+0x08 = 0`, `ord+0x14`, `ord+0x60`, `mover+0x7c` and `ord+0x50`. `R0007` stores `actor+0x54 = 0`. Neither routine contains a store to `actor+0x58`, `ord+0x09`, `ord+0x15` or `actor+0x136` (`rom-stop-executor.txt`).
- Executor: `R0016` keeps the action word only when it is 2 or 0xf and stores 0 otherwise (`L00539`..`L00540`); progress 1 restores action 3, progress 2 restores 0xd or 0xe from `ord+0x5c`, progress 3 gives 1 and progress 4, 5 or any other value gives 0x1a. It clears progress only when `ord+0x15 > 2` and `actor+0x136` is set.
- Slot: the actor slot `R0037` calls the executor at `L00086` unless `R0416` returns nonzero (the test `actor+0x3c == 0`), then dispatches on `actor+0x54` at `L00538` through the byte table `L01775` and the dword table `L01776`: action 1 to `L01777`, 2 to `L00009`, 3, 0xd and 0xe to `L00005`, 0xf to `L00196`, others to `L02936` (`rom-tables.txt`). The arm at `L00005` reads `actor+0x58` for its phases 0, 5 and 7 (`rom-strike-arm.txt`).
- Progress clear: the executor clears progress only when `ord+0x15 > 2` and `actor+0x136` is set (`L01931`..`L02015`); the stops write neither, so with the counter at 2 or less or the flag clear the progress byte survives the stop.
- Inference: a stop that clears `actor+0x54` while `actor+0x58` is loaded leaves the phase in place; the next executor pass rebuilds the action word from progress and not from the phase. Whether the strike then completes depends on the progress byte, which the stop does not clear and the executor clears only under the condition above.

**Confidence.** High for the stores, the tables and the executor mapping, each read from the listing and the PE bytes. Medium for the run-time outcome that the strike continues, a composition of the arm and the executor without a run, conditional on progress not being cleared elsewhere.

**Unknown.** What `actor+0x3c` holds, so which actors the executor skip applies to; other readers of `actor+0x54`, such as animation; and the outcome of a stop issued on the tick the skip applies.

**Evidence.** [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md), `evidence/listings/rom-stop-executor.txt`, `evidence/listings/rom-strike-arm.txt`, `evidence/listings/rom-tables.txt`

### AI-382

- First pass: executor row 7 at `L00549` stores action 2 at `L00026`. The arm for action 2 at `L00009` calls `R0446` (cell record lookup, returns the sack pointer), `R0447`, `L02016` and `R0448` (`rom-pickup-arm.txt`, `rom-pickup-callees.txt`).
- Transfer: by `ITEM-PICK-009` `R0448` credits `sack+0x3c` to `Player+0x38` through `R0449` (message 0x67) and moves the whole container into `actor+0x7c` with `R0450`; it recomputes the load with `R0451`, nulls `sack+0x40` and deletes the sack.
- Second pass: `R0007` clears the action word and the completion stores follow (`actor+0x50 = 0xc`, `ord+0x08 = 0`, `ord+0x50 = 1`).

**Confidence.** High for the call sequence, which is read from the listing. Medium for the combined outcome that the transfer needs no second pass, since it relies on `ITEM-PICK-009` for the contents of `R0448`.

**Unknown.** Whether the arm for action 2 is reached on every tick that the action word equals 2, and what the sack lookup returns for a cell whose record is already deleted.

**Evidence.** [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md), `evidence/listings/rom-pickup-arm.txt`, `evidence/listings/rom-pickup-callees.txt`

### AI-383

- Order 3 arm: `L00307`..`L02017` in `R0023`. Order 5: `R0026`, gated by `[EDI+0xbb4]`. Both call `R0015` for a member with `ord+0x20 == 0` and `Player+0x28 == 0`; neither reads `ord+0x08` or `ord+0x0a` before the call (`rom-group-dispatch.txt`, `rom-order5-eval.txt`).
- Branches: a member with `ord+0x20` nonzero goes to `R0009` (`L00842`, `L01708`), which was not read. A member with `ord+0x20 == 0` and `Player+0x28` nonzero gets `ord+0x08 = 0xb` (`L00158`, `L01028`) and, when a mask test on `actor+0x4c` passes, `R0113` (`L02018`, `L02019`), also not read. A member with both zero reaches `R0015` (`L00051`).
- Stores of `R0015`, published in `AI-349`: `ord+0x08 = 0` at `L00053` on five paths (`actor+0x140 == 0`, the word at `+0x9a` equal to 0, `+0x9a <= +0xa0`, a null `R0017(+0x140, 6)`, and a best ally at or above 1.0) and 8 through `R0018` on the heal path. Those stores apply to the `ord+0x20 == 0` and `Player+0x28 == 0` branch only; `R0018` has no listing in this experiment.
- Tail: the group tail (`R0105`, `R0100`, `R0106`, `R0107`) is the withdrawal logic of `AI-RETREAT-274`, `AI-WITHDRAW-026` and `AI-WITHDRAW-027`; `R0100` stores `ord+0x08 = 1` with a flee cell.

**Confidence.** High for the missing tests, the branches and the 0xb store, each an instruction in the listing. Medium for the consequence on a member holding a destination, which composes the arm with the meaning of `Player+0x28`, itself Unknown.

**Unknown.** The bodies of `R0009` and `R0113`; which Player value `Player+0x28` takes for each controller; whether the evaluation also runs for a member that already stands; the order 4 arm after a completed pickup.

**Evidence.** [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md), `evidence/listings/rom-group-dispatch.txt`, `evidence/listings/rom-order5-eval.txt`, `evidence/listings/rom-group-tail.txt`

### AI-384

- Values: `R0116` reads `[Path Finding]` through `R0452`; `MOVE-PARAM-006` gives the shipped values equal to the defaults: StaticScanAhead 5 (`+0x585b4`), DynamicScanAhead 3 (`+0x585b8`), StaticRefreshRate 16 (`+0x585bc`), DynamicRefreshRate 32 (`+0x585c0`), DynamicByStaticLookup 3 (`+0x585c4`), StaticIsntNeeded 5 (`+0x585c8`).
- Counters: `mover+9` is incremented at `L02020`, `L01919` and `L02021`, and `mover+0x78` at `L02022`; both are per-pass counters that trigger a re-search, and neither ends the walk. The give-up here means the flag `mover+0x98`.
- `R0178` re-searches when the target changed or `mover+9 > 16` (`L02023`..`L02024`) and runs the near search when the dynamic list is empty or `mover+0x78 > 32` (`L02025`..`L02026`).
- `R0043` does a full search when `mover+9 > N/3 + 1` (N is the static count `actor+0x168`; `L02027`..`L01895`) and the previous length `mover+0x8a` is above 5; at length 5 or less it rebuilds the single destination node. It sets `mover+0x98 = 1` when the static count is 0 after the search (`L00161`). `R0178` also sets it at `L01720`, where the list is empty.
- `R0055` picks kind 1 (destination) when N is 5 or less, kind 2 (head) when the head is farther than 3, otherwise kind 3 (the node three hops after the head); it sets `mover+0x98 = 1` only when the dynamic count is 0 and the kind is 1 (`L01909`..`L00162`).
- Search budget: d plus the larger of the scan-ahead value and d/4; static uses `+0x585b4` (`L02028`..`L02029`), dynamic `+0x585b8` (`L02030`..`L02031`). The label plane is `ctx+0x30000`, initialised to 0xffff (`L01912`); by `MOVE-PLANE-005` the dynamic search reads the occupancy plane, so an occupied cell is not expanded and keeps 0xffff.

**Confidence.** High for each threshold and its comparison, read from the listing, and for the absence of a give-up counter in the four routines named above; other routines were not searched. Medium for the occupied-cell outcome, a composition with `MOVE-PLANE-005`.

**Unknown.** Readers of `mover+0x98` outside the executor tail already described by `AI-375`; the effect of `R0453` on occupancy labels; any give-up held outside the searched routines (`rom-pursuit.txt`, `rom-walk.txt`, `rom-search.txt`).

**Evidence.** [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md), `evidence/listings/rom-pursuit.txt`, `evidence/listings/rom-pathcfg.txt`, `evidence/listings/rom-search.txt`, `evidence/listings/rom-census-54-94.txt`

### AI-385

- `R0147` calls `R0076` when the tick counter `[obj+4]` and `0xf` equal 6 (`L00847`..`L00279`), before `R0193`; it has 4 direct callers (`R0260`, `R0454`, `R0455`, `R0099`).
- `R0075` is a thread loop: one call of the AI pass at `L00163`, then 16 calls of `R0193` (`L01864`, loop to 0x10) on a 0x3e ms schedule. It is the body of the thread procedure `R0261`, which the launcher `R0456` pushes (`L02032`) and creates through `R0457` when `obj+0x2c` is nonzero. `R0147` returns early when `obj+0x2c` is zero (`L02033`..`L02034`), the flag the launcher tests.
- Both routes give a period of 16 ticks. Retreat admission is `AI-RETREAT-274`, `AI-WITHDRAW-026` and `AI-WITHDRAW-027`; no new claim is made for it.

**Confidence.** High for the period on each route, read from the listing. Medium for the two callers sharing the `obj+0x2c` gate; whether both routes run in one session is Unknown.

**Unknown.** Whether the launcher `R0456` is reachable: it has no direct call, no reference and no raw dword hit in the PE (`rom-enum-cadence.txt`, `rom-tables.txt`; the one hit for `R0261` is the launcher's own push at `L02032`); whether both routes run in one session; and which route a given session mode uses.

**Evidence.** [EXP-0451](../experiments/EXP-0451-attack-cycle/EXP-0451.md), `evidence/listings/rom-cadence.txt`, `evidence/listings/rom-enum-cadence.txt`, `evidence/listings/rom-tables.txt`

## Step into a held cell

Evidence is a static read of `rom.exe` (one image on both lawful installs) with a capstone sweep; no process was run. `EXP-0452` was allocated ids `387`..`389` of `claims/ai.md` and spent `AI-387`; `AI-388` and `AI-389` are unused.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-387 | A mid-transit walk call has no destination check: stepping into a cell a live unit holds releases blind, rewrites the position and refuses the occupy with no slot store and no rollback. | High / Medium | ● active | [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md) |

### AI-387

`R0178`, when the sub-bytes are not both 0x80, calls `R0039` at `L02035` with no read of the destination cell; `R0039` calls `R0047` (`L02036`) and recomputes the stride `+0x72` after a cell change. A transit loaded from a save therefore reaches the step with whatever the cell holds. On a crossing into a cell held by a live unit the step (`MOVE-088`, `MOVE-094`) does, in one call:

1. the release at the old footprint: a non-zero slot test at the old footprint that does not compare the holder;
2. the claim at the old cell (`L02037`);
3. the position rewrite into the held cell (`L02038`..`L02039`);
4. the occupy: `R0458` finds the slot taken (`L02040`, `L02041`), calls the one-instruction stub `R0459` and returns 0, with no slot store and no recompute; the occupy returns 0 and the step ignores it (`L02042`, `L02043`).

No code in the step restores the old position (`SAV-CELLFAIL-583`). Later, the mover's departure release at that cell is blind and clears the other unit's slot (`MOVE-088`); the recompute then runs for that cell.

**Confidence.** High for the missing destination check on the path from the walk to the step, and for the occupy refusal and untested return, asserted instructions. Medium for the consequence at the departure release: a composition of read branches with no run.

**Unknown.** Whether an executable path loads a transit into a held cell other than a save. What the other unit's code does with its cleared slot. Observed play.

**Evidence.** [EXP-0452](../experiments/EXP-0452-area-cost/EXP-0452.md), `evidence/listings/rom-transit.txt`, `evidence/listings/rom-step.txt`, `evidence/listings/rom-slots.txt`

## Melee reach to fliers

Evidence is a static read of `rom.exe` (one image on both lawful installs) and a census of 340 distinct lawful saves; no process was run. `EXP-0481` was allocated ids `390`..`393` of `claims/ai.md` and spent all four. A flier is an actor whose `vt+0x20` movement domain is 3 (`MOVE-DOM-024`); the matrix and veto are `AI-PREF-070`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-390 | The strike executor `R0246` and the reacquisition routine `R0004` compare distance with reach and read no movement domain, so reacquisition can take a flier within reach and the strike applies no domain rule of its own. | High / Medium | ● active | [EXP-0481](../experiments/EXP-0481-flier-melee/) |
| AI-391 | Automatic acquisition keeps a reach-1 ground creature off fliers on three routes: the scorer's zero cell (`AI-FLIER-073`), `R0460`'s list removal for reach below 2, and a standing-acquisition cost no flier can pass. | High / Medium | ● active | [EXP-0481](../experiments/EXP-0481-flier-melee/) |
| AI-392 | A creature with reach above 1 is refused a flier by no route read: the scorer's ranged row has no zero cell, `R0306` engages the nearest flier before any ground candidate, and standing acquisition admits one inside reach. | High / Medium | ● active | [EXP-0481](../experiments/EXP-0481-flier-melee/) |
| AI-393 | In 340 distinct lawful saves no domain-1 or domain-2 actor of reach at most 1 holds a domain-3 actor as combat target or order victim; the one domain-3 target is held by a domain-1 actor of reach 4. | Medium | ● active | [EXP-0481](../experiments/EXP-0481-flier-melee/) |

### AI-390

- `R0246` (the strike executor) loads the reach byte `+0x12c` at `L00735`, takes the distance from `R0247` (`L02044`) and refuses when reach is below it (`L02045`..`L00736`). Its body and `R0247`'s hold no `vt+0x20` call and no `+0x4a` load.
- `R0004` keeps a candidate when `R0255` returns at most the reach byte (`L00796`..`L00617`). The routine's 270 listed lines hold no `vt+0x20` or `vt+0x1c` call and no `+0x4a` load; its stack loads at `[ESP+0x1c]` are argument slots.
- `R0004` has one caller, `R0016` (`L00011`), the arm `AI-ROUTE-045` names for a failed path.

- `R0246` has one direct caller, `L02046`; the target's damage routine is called through `vt+0x4c` at `L02047`.

**Confidence.** High that the three routines named carry no domain read in their listed bodies. Medium for the whole route: the callees `R0255`, `R0005` (called at `L00012`), `R0051`, `R0032`, `R0041`, `R0205` and `R0250` were not read for a domain test. That a creature holding a flier as victim can strike it is an inference from the strike's reach-only compare.

**Unknown.** Whether those callees, the `vt+0x4c` damage routine or the layers between an order and `R0246` filter by domain. Whether any route outside the automatic ones (a player order, a script arm, retaliation after damage) hands a flier to a reach-1 creature as victim.

### AI-391

- Scorer `R0225` (`AI-PREF-070`): a melee member, reach byte `+0x12c` at most 1, reads cell `4 * member domain + candidate label` (`L02048`..`L02049`); cells `M[1][3]` and `M[2][3]` are 0 and return `0xffffff` (`L02050`..`L02051`), which the caller treats as no candidate (`L02052`..`L02053`). A byte scan of `.text` finds at least four call sites of the scorer: `L02054`, `L02055`, `L00766`, `L01157`; `EnumRefs` returned three and missed `L02055`.
- `R0460(actor, list)` selects every list element with `vt+0x20 == 3` when the actor's reach byte is below 2 (`L02056`..`L02057`) and unlinks it (`L02058`..`L02059`). A byte scan finds three call sites: `L02060`, `L02061` and `L02062` (`R0306`, which then falls to `R0022` on an empty list). The first two sit in streams repeating the `R0020` / `R0021` / `R0460` sequence and are unread consumers of the same removal; `EnumRefs` missed both.
- `R0022`, standing acquisition: best distance is seeded reach plus 1 (`L00222`..`L02063`); a candidate's distance from `R0036` rises by 1 when the candidate is a flier and the actor is not (`L00224`..`L00226`); a candidate is taken only below the best distance (`L01745`), so a flier needs a distance value below reach. `R0036` ends with an arithmetic shift right by 8 and then an increment by 1 (`L02064`..`L02065`), so its minimum is 1: a flier costs at least 2, the seed for reach 1 is 2, and the take test never passes. At reach 1 the cost rule is a full exclusion on the first pass. The routine has at least 13 call sites (`EnumRefs`); a byte scan finds 18, not resolved into instruction streams.
- Both cells and the list filter apply to reach at most 1 whatever the weapon class; for a domain-2 member the veto cell is `M[2][3]`, and fliers are domain 3 only. The domain-2 rows leave `M[1][2]` and `M[2][2]` nonzero (`1`, `4`).

**Confidence.** High for the scorer cells, the list removal and the cost arithmetic including the floor of 1, each a cited instruction. The scorer cells restate `AI-FLIER-073`; the list removal and the cost arithmetic are new here.

**Unknown.** Whether the second pass of `R0022` over the `+0xbc4` list (hostile candidates at health below 1, `L02066`..`L02067`) can name a flier corpse: that pass adds no flier term.

### AI-392

- The scorer's ranged branch (reach above 1) indexes the member-independent row 0 (`L02068`..`L00785`), whose cells are `2 1 2 4` (`AI-PREF-070`), none zero.
- `R0306` runs `R0460` first, which removes nothing for reach 2 or more (`L02056`..`L02069`). Over the remaining list it records the nearest candidate with `vt+0x20 == 3` (`L02070`..`L02071`) and engages it through `R0009` (`L02072`) before looking at any non-flier; with no flier it engages the nearest candidate (`L02073`).
- `R0022` with reach 2 or more admits a flier at distance up to reach minus 1 and its other candidates up to reach.

**Confidence.** High for the three instructions' effect on the candidate list. Medium that the list at `R0306` holds hostile units only: the producer `R0020` was not re-read.

**Unknown.** Hostile filtering in `R0020`.

### AI-393

- Instrument: `tools/flierpairs`, which walks each save's actors with the `Unit::Serialize` reader of `EXP-0147`, resolves the combat target dword `+0x5c` and the order victim dword `+0x0c` against actor identity keys of the same save, and tabulates the attacker's `+0x4a`, `+0x12c` and the target's `+0x4a`. Input: 340 distinct saves out of 507 `.sav` files (`evidence/pairs/save-digests.json`).
- 9065 actors: domain 0 reach at most 1 11, domain 0 reach 2 or more 4, domain 1 reach at most 1 4678, domain 1 reach 2 or more 2990, domain 2 reach at most 1 747, domain 3 reach 2 or more 635. Domain 3 actors occur in 59 saves.
- `+0x5c` is nonzero on 220 actors and resolves on 65; the order victim is nonzero on 112 and resolves on 103 (`evidence/pairs/target-pairs.tsv`).
- Resolved pairs with a melee attacker (reach at most 1), as records (saves): domain-1 attacker against domain-1 target 9 (9) by `+0x5c` and 87 (5) by order victim; domain-1 attacker against domain-2 target 1 (1) and 3 (2); domain-2 attacker against domain-1 target 1 (1) and 1 (1). No melee attacker holds a domain-3 target.
- The one domain-3 target by `+0x5c` is held by a domain-1 attacker of reach 4 (`evidence/pairs/target-pairs-detail.tsv`, save `60267c82072c`).

**Confidence.** Medium: a census over saves the owner and the install happened to produce; 155 of the 220 `+0x5c` values do not resolve, so the table covers 30 percent of the nonzero targets and cannot exclude a melee-on-flier hold among the rest. The null has little power: one domain-3 target appears among all 65 resolved `+0x5c` pairs and none among the 103 resolved order victims, and the 87 melee order victims come from 5 saves, so the absence of a melee-on-flier pair is expected whether or not a route exists.

**Unknown.** The unresolved 155 targets. Whether any save was taken while a melee creature held a flier, which the veto routes read here predict it would not.

## Debug console reachability

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-397 | The Alt console tests no participant flag, and the only writer found that raises the privilege byte is `#Chicken`, which a participant flag blocks; so a multi-participant session refuses every key unless a Player was privileged before. | High / Medium | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| AI-398 | A console reply is one `0x91` record with `+7 = 0` and `+0xa = 0`, which `R0217` treats as the all-connections path; it carries no notice 5, 6 or 7. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |

### AI-397

`R0440`, whose only caller is `L02000` in the `0x80` arm of `R0061`, requires `byte[cmd+4] == 0`, `[L00004] != 0` and `Player+0x68 > 0x32`, then the index tests of `AI-378`. It reads no `server+0x0c` (`rom-console.txt`). The only writer found that raises the byte is `#Chicken` (`MENU-102`; `callto:R0445` and the `disp:68` census, not other operand forms), and `R0430` returns before that test when `server+0x0c` is set (`MENU-100`). A fresh Player holds 0 (`MENU-103`).

So the console is usable on a map where the participant flag is 0 after one `#Chicken`, and refuses every key where the flag is set. The exception is a Player whose byte was raised earlier in the same process, which depends on the lifetime left Unknown in `MENU-103`.

**Confidence.** High for the three tests and the absence of a `server+0xc` read in the console and in the `0x46` arm (`rom-dispatch.txt`). Medium for "only writer" (`AI-378` left other operand forms unexcluded and `MENU-103` bounds its census) and for the multi-participant refusal, which also rests on `MENU-100`.

### AI-398

The console builds each reply line through `R0443`, a `0x91` record, and sends it with `R0217`. All 12 call sites in `R0440` push Player 0, so `+7` is stored as 0 (`L13191` to `L13192`). `R0217` tests `[arg+7]` at `L12414`: zero walks the connection list at `+0x18b8`, nonzero looks one connection up by id. `#event` differs: `R0188` passes its Player, so `+7` is that Player's id. The notices of `MENU-111` are not used by the console arms.

**Confidence.** High for the record fields and the two branches of `R0217`. Unknown for any filter inside the connection loop (`L13193`, `L13194`, not read) and for the client panel that shows the line.

## Ranged release facing

Evidence is a static read of `rom.exe` (one image on both lawful installs, sha256 `942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`); no process was run. `EXP-0493` was allocated ids `405`..`412` of `claims/ai.md` and spent all eight. Blocks: `actor+0x158` is the order block (`ord`), `actor+0x154` the mover; `mover+0` is the current facing byte and `mover+1` the wanted facing. Listings and scans are in `experiments/EXP-0493-ranged-release-facing/evidence/`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-405 | High: the load gate `R0041` tests sim byte `mover+0` equal to the 8-way heading (tolerance 0) and edge distance at most reach. Medium: a walker finishes its step and turns before it loads, for the five gated writers, not `L06163`. | High / Medium | ● active | [EXP-0493](../experiments/EXP-0493-ranged-release-facing/EXP-0493.md) |
| AI-406 | A move order during a loaded cycle does not interrupt it; an attack order given during the following walk acts at the next cell centre, then turn, load and shot follow in that order. | Medium | ● active | [EXP-0493](../experiments/EXP-0493-ranged-release-facing/EXP-0493.md) |
| AI-407 | Among read routines, automatic acquisition (`R0004`) loads through the same gate or turns without loading and none replaces the victim or resets the phase while `ord+9` is 1; five writers and `R0008` are unread. | Medium | ● active | [EXP-0493](../experiments/EXP-0493-ranged-release-facing/EXP-0493.md) |
| AI-408 | No `mover` access in the state-3 body, swing start and strike ranges or their 16 direct callees (five virtual calls unread); the countdown continues, the strike tests reach only, and 0x72 carries `(facing + 8) >> 4` at load. | High / Medium | ● active | [EXP-0493](../experiments/EXP-0493-ranged-release-facing/EXP-0493.md) |
| AI-409 | The shot builder reads the 0x72 direction; while the `+0x86` victim lookup succeeds the projectile driver re-reads the victim's point every tick, so projectile direction is not the stored direction. ShootDelay is `units.reg` data. | Medium | ● active | [EXP-0493](../experiments/EXP-0493-ranged-release-facing/EXP-0493.md) |
| AI-410 | A shot can leave over one heading from the drawn facing if the victim's bearing changes by more than one heading between load and shot (not computed); on identical delta vectors 8-way and 16-way indices differ by at most half a heading. | Medium | ● active | [EXP-0493](../experiments/EXP-0493-ranged-release-facing/EXP-0493.md) |
| AI-411 | Pending order 2 loads a cycle with no facing or distance test, but its only immediate setter `L13240` has no reference in the scanned population (`E8`/`E9`, dword pointers, `[reg+8]`, `[reg+9]` stores), so no read route reaches it. | Medium | ● active | [EXP-0493](../experiments/EXP-0493-ranged-release-facing/EXP-0493.md) |
| AI-412 | The client drops a 0x72 message while the drawable run counter `+0xa0` is nonzero, as for 0x6b and 0x6d; for arcs of 0x21 or more the 0x6d count equals the sim step count; for the one-heading snap the sequence is Unknown. | Medium / Unknown | ● active (partially retracted) | [EXP-0493](../experiments/EXP-0493-ranged-release-facing/EXP-0493.md) |

### AI-405

- `R0041(actor, target, reach)` compares `mover+0` with the result of `R0051` at `L00577`. `R0051` returns a multiple of 32 (pairs of 16-way sub-sectors, output shifted left by 4, boundaries at a 2:1 slope), so the equality is exact and a facing one heading off fails. Edge distance from `R0036` must be at most the reach (`L00578`).
- `R0016` runs once per actor tick before the body (`L00086`, gated at `L13241` by `+0x3c == 0`). The pending-order switch is read only at `ord+9 == 0`. Pending 5 and 6 (`L00101`, `L00581`) are the pursuit arms; each stores `ord+9 = 1` only after a passing gate (`L00743`, `L00744`).
- Outside the stop distance the arm paths; at a cell centre inside it `R0043` calls `R0051` then `R0056`. `R0056` snaps `mover+0 := mover+1` when `mover+0xa0 == 0` and the arc is below 0x21 (`L00749`, `L00750`), so a one-heading turn completes in one call; larger arcs rotate by the rotation speed `mover+0xa` and set `mover+0xa4` to ceil(arc / rate).
- A step already begun continues through `R0039` and ends at a cell centre (progress 3), so the turn follows the step.

**Confidence.** High for the gate's two tests and tolerance (on the sim byte `mover+0`, not the client drawn facing) and for the single executor call per tick. Medium for the sequence step, turn, load: the callees `R0039`, `R2130` and `R0036` were read for their effect, not end to end. Turn-before-load holds for the five gated `ord+9 = 1` writers, not for pending order 2 (`L06163`, `AI-411`).

**Unknown.** Register-form or indirect stores to `ord+9` and `ord+8`, and orders restored from a save.

### AI-406

- The move setter `R0301` and the common reset `R0007` both write `mover+1 := mover+0` and `actor+0x54 = 0`; neither writes `ord+9`. The move setter also writes `actor+0x50 = 1` and `ord+8 = 1`.
- The progress-1 arm at `L00092` sets state 3 again, increments the counter and clears progress and state only when the counter exceeds 2 and `+0x136` is set. So a loaded cycle runs to its recovery (`AI-RETREAT-272`) and the walk starts afterwards.
- During the walk the progress is 3 until a cell centre (`L00096`); the new attack order is consulted at progress 0 only, so it waits for the centre, then the gate of `AI-405` decides turn or load.

**Confidence.** Medium: assembled from `AI-405`, `AI-RETREAT-272` and the setters; no run observed the order of events.

**Unknown.** The exact tick count between the move order and the first walk step.

### AI-407

- `R0004` has one caller, the tail at `L00011`. The tail at `L00010` clears `mover+0x98` and branches on `actor+0x50`: 1 tears down to 0xc, 0xa and 0x17 have their own arms, every other value calls `R0004`.
- The routine admits candidates within reach, scores them by turn cost (`AI-361`) and writes `ord+0xc` and `ord+8 = 6`. If `R0041` passes (`L00741`) it calls `R0005` (progress 1, counter 0, state 3, victim); otherwise `R0051` then `R0250` (turn only) and state 1.
- `mover+0x98` is written only at `L01720`, `L00161` and `L00162` (movement internals), whose callers are executor arms or orphan routines. The tail is therefore not entered from a progress-1 invocation, and no read writer of `ord+0xc` or `actor+0x5c` runs during progress 1; the statement covers read routines only.

**Confidence.** Medium: a call-graph reading over raw `E8`/`E9` scans (`evidence/xrefs.txt`), blind to indirect calls.

**Unknown.** The group slot-level arm `R0008` for `actor+0x50 == 1`, other phase and victim writers (`L13242`, `L13243`, `L13244`, `L13245`, `L13246`: hurt, stun, cast cancel, death) which were not read (they store 0 to `actor+0x58` or `actor+0x5c`), the stores at `L13247`, `L13248`, `L13249` and `L13250` in the region of `R0008`, and the dedicated server.

### AI-408

- State 3 phase 0 (`L00005`): the swing start `R0245` (`L05237`) sends the message and phase 5 counts down from the charge (`+0x134`) plus a distance term. At zero it calls `R0001` then `R0246` (`L00006`). Phase 7 sets phase 0 and completion 1.
- Guards of `R0246`: target non-null, health above 0, reach at least the distance (`L00735`..`L00736`). No `+0x154` access exists in `L00005`..`L00009`, `R0245`..`R0001` or `R0246`..`R0563` (`evidence/calls-state3-swing-strike.txt`).
- `R0245` calls `R0549`, which sends 0x72 (via `R1064`) when reach exceeds 1 and 0x71 otherwise. Fields: `+0xc` = `R2131(mover)` = `(facing + 8) >> 4`, `+0xd` charge plus relax, `+0xa` actor id, `+0xe` victim client id.

**Confidence.** High for the bounded absence (no `mover` access in the three ranges nor in their sixteen direct callees, depth 1) and for the message fields. Medium for a universal "no facing test after the load". Medium that nothing else writes `mover+0` meanwhile: the callers of `R0056` (`AI-411`, 11 raw sites) are not reachable from this path.

**Unknown.** The five virtual calls `L13251`, `L13252`, `L13253`, `L02047` and `L04400`, and callees below depth 1.

### AI-409

- Client message arm `L09206` sets action 7, `+0x85 := msg +0xc`, phase 0 and the target `+0x86`. The driver `R0548` action-7 arm (`L02711`) calls the shot builder (`vt+0x58`, `R0603`) at phase equal to ShootDelay (`L02716`) and writes `+0x6c := +0x85` every tick after that call (`L02535`).
- ShootDelay (`EXP-0428`, `shot-classes.csv`): Catapult and Ballista 0, Dragon 8, Goblin slinger 10, Archer 21, Crossbowman 9, Orc Archer 22, Sonic Bat 2. For 0 the builder reads `+0x6c` before the first write.
- The builder looks the victim up by `+0x86` and builds nothing if absent; its origin offset is `ShootOffset[(+0x6c - 8) & 0xe]`.
- The projectile driver `R0558`, action-1 arm, copies the victim's point (`+0x58`, `+0x5c`, `+0x10`) to `+0x88`..`+0x90` on every call while the lookup by `+0x86` succeeds (on failure, `L13254`, the last copy and direction stay; every shot from the builder takes this arm, `+0x84 = 1` at `L13255`), derives the 16-way direction through `L02828`, stores it at `+0x85` and `+0x6c` (`L02827`..`L13256`) and steps the position by (target - position) over the remaining segments (`L05578`..`L05579`).

**Confidence.** Medium: read from listings; the ShootDelay table is cited from `EXP-0428`, not re-derived.

**Unknown.** Whether the ShootDelay and reach rows (`units.reg` data, cited from `EXP-0428`) are equal on the EN and RU installs. Whether the drawable body redraw uses `+0x6c` or `+0x85` at the shot tick (read for the builder only).

### AI-410

- Existence case: the victim moves after load. The body is not turned (`AI-408`) while the projectile homes (`AI-409`). Exceeding one 8-way heading needs the victim's bearing to change by more than one heading within ShootDelay ticks; no magnitude was computed.
- Not a case here: a ShootDelay-0 class reads a stale `+0x6c`, which feeds only the origin offset (`L13257`..`L13258`), not the direction; whether it is the old facing depends on the client turn count (`AI-412`) and is Unknown.
- Instrument `dircompare.py`: instruction-level ports of `R0051` (8-way) and `L02828` (16-way) over |dx|,|dy| at most 400 on identical vectors, 320800 cells each, give index differences of 0 or 1 sixteenth-index only (`evidence/dircompare-400.txt`), at most half a heading. The sim routine reads cell coordinates and the projectile routine world points, so the result holds on identical delta vectors only.

**Confidence.** Medium. The comparison does not bound the origin shift of `ShootOffset`, footprint, or sim and client point differences.

**Unknown.** The magnitude of those three terms; any runtime witness.

### AI-411

- `L06163` is pending order 2: it stores `ord+9 = 1` with no `R0041` call. A byte-immediate scan of L13259..L13260 and a `[reg+9]` displacement scan of L02478..L13260 find six `ord+9 := 1` writers (`L06163`, `L00103`, `L00746`, `L00747`, `L13261`, `L13262`); the other five follow a gate.
- The only immediate byte writer of `ord+8 = 2` is `L13240`; it has no `E8`/`E9` call reference and no dword reference (`evidence/imm-ord8-store2.txt`, `evidence/xrefs.txt`). The loader variants `L13263`, `L13264` and `L13265` have none either.
- `AI-FACE-066` names `L00738` among routes on a taken `R0041` branch; `L00738` is the pending-2 arm and `L00742` is cast arm 8. `AI-FACE-067` counts 7 hits and 6 owners for `R0056`; the raw scan finds 11 sites: `L00752`, `L13266`, `L00751`, `L13267`, `L00753`, `L01933`, `L01934`, `L00754`, `L07010`, `L13268`, `L13269`.

**Confidence.** Medium: negative reachability over the named scans (immediate stores, `[reg+9]` displacements, `E8`/`E9` calls, dword pointers). The ungated load is reachable only through `ord+8 == 2` and any restored order; the route is not called absent beyond that population. Three dword stores of 2 to `[reg+8]` (`L13270`, `L13271`, `L05400`) lie outside the order code and were not resolved.

**Unknown.** Indirect callers, register-form stores and orders loaded from saves or scripts.

`AI-FACE-066` and `AI-FACE-067` are narrowed in `claims/retracted.md` for these facts.

### AI-412

- The client 0x72 arm is accepted only when the drawable run counter `+0xa0` is 0; otherwise it logs "Overriding by Shoot" and returns. The 0x6b (move) and 0x6d (turn) arms have the same gate. Opcodes 0x6c, 0x6e, 0x6f and 0x70 are state sync with a mask whose bit 0x10 writes `unit+0x6c` (`L02537`); the tick-time projector masks seen are 1, 2, 409 and 20.
- For arcs of 0x21 or more the 0x6d count `mover+0xa4` equals the sim step count (`ceil(arc / rate)`). For a smaller arc `R0056` snaps `mover+0` at once and sets `mover+0xa4 = 1` (`L13272`..`L01977`), so the gate can pass on the next sim tick and send 0x72 while the drawable run counter may still be nonzero.

**Confidence.** Medium for the gate and the masks; Unknown for a drop after a one-heading snap turn, which no run observed.

**Unknown.** A full-mask re-projection landing mid-cycle.

**Amended.** The clause that the 0x6b and 0x6d arms share the 0x72 gate is refuted by EXP-0498: both apply their message while `+0xa0` is nonzero, replacing the running action (`ANIM-135`; see [`retracted.md`](retracted.md)). The 0x72 gate stands.

## Pursuit search cadence

Evidence is a static read of `rom.exe` (one image on both lawful installs) and CPU emulation (unicorn) of the image's own bytes for `R0043`, `R0053`, `R0055`, `R0054` and their leaves, with the step, the position write and the allocator replaced by stated models; no process of the game was run. `EXP-0494` was allocated ids `413`..`420` of `claims/ai.md` and spent `413`..`418`. `mover` is `actor+0x154`; a pass is one call of `R0043(actor, victim, stop)`; a sub-tick is one call of `R0193` (`SESS-TICK-006`).

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-413 | In the pursuit pass `mover+0x8a` is a word holding the static route count after the last full search or single-node rebuild, set to 0x00ff when the victim differs from `mover+0x7c`; the mover constructor zeroes it. | High / Medium | ● active (amended) | [EXP-0494](../experiments/EXP-0494-pursuit-search-cadence/EXP-0494.md) |
| AI-414 | On a pass by a centred mover whose victim differs from `mover+0x7c` and whose edge distance exceeds the stop distance, the full static search runs (threshold 5); an off-centre mover finishes its step and a mover within stop turns. | High | ● active | [EXP-0494](../experiments/EXP-0494-pursuit-search-cadence/EXP-0494.md) |
| AI-415 | For an unchanged victim a re-search is a full search only while the last count exceeds 5; at 5 or less it rebuilds the old route end as one node, count 1, until a route-end match clears `mover+0x7c` and forces a full search. | High / Medium | ● active | [EXP-0494](../experiments/EXP-0494-pursuit-search-cadence/EXP-0494.md) |
| AI-416 | With an empty dynamic list the stepper takes no step and only turns `mover+0` toward `mover+1`; position, claim and `mover+0x90` stay, and unless the near search set `mover+0x98` (`AI-373`) the next sub-tick passes again. | High | ● active | [EXP-0494](../experiments/EXP-0494-pursuit-search-cadence/EXP-0494.md) |
| AI-417 | At most one pass per sub-tick, none in transit or on arrival. Byte `mover+0x78` counts centred out-of-stop passes since the last near search; byte `mover+0x09` near searches and route-end matches since the last full search or rebuild. | High / Medium | ● active | [EXP-0494](../experiments/EXP-0494-pursuit-search-cadence/EXP-0494.md) |
| AI-418 | Modelled band case (rings 1..4 impassable): from Chebyshev 7 a pursuer with reach below 7 raises `mover+0x98` on its first pass; from 12 it walks to ring 5 and raises it there when reach is below 5, else stands in reach. | Medium | ● active | [EXP-0494](../experiments/EXP-0494-pursuit-search-cadence/EXP-0494.md) |

### AI-413

- Writers in `R0043`: `L13285` stores the word 0x00ff in the victim-change block (`L13286`..`L01893`), which also sets `mover+0x09 = 0xff`, frees the static list and stores the victim cell in `mover+0x76` and `mover+0x8c`. `L13287` stores the low word of the static count `actor+0x168` after either the full search (`L01894`) or the single-node rebuild (`L13288`..`L01898`).
- Readers: `L13289` (full search when the word is greater than `ctx+0x585c8`, shipped 5 by `MOVE-PARAM-006`) and `L13290` (the dynamic list is freed after a search when the word is greater than 5).
- The mover constructor `R0206`, called at `L13291` before `actor+0x154` is stored at `L13292`, zeroes 0xb4 bytes (`rep stosd`, 0x2d dwords) and then sets `+0x05 = 0x41`, `+0x0a = 0x10`, `+0x09 = 0xff`, `+0x08 = 5`. A new mover therefore has `mover+0x7c = 0` and `mover+0x8a = 0`.
- After a victim change the field is 0x00ff while `mover+0x09` is 0xff; the next read at `L13289` is in the same pass.
- Emulation (`q2-first-call.tsv`, 108 first passes on open ground): after a search from Chebyshev distance d the field equals d and the count left after the near search equals d - 1.

**Confidence.** High for the two writers, the two readers and their values in `R0043`, read whole. Medium for "only": a displacement scan of `.text` finds 11 other word stores at `+0x8a` (`L13293`..`L13294`), all outside the movement routines; their base objects were not resolved.

**Unknown.** Whether a save load or a script writes `mover+0x8a` or `mover+0x7c`. Other constructors of a mover than the call at `L13291` were not enumerated.

**Amended.** The constructor enumeration is narrowed by `MOVE-119`: the image holds four direct calls of `R0206` (`L00964`, `L14119`, `L14120`, `L13291`) and no stored address of it, from a capstone operand sweep of `.text` plus a raw dword scan of every section. Medium over the unreached bytes; the base objects of the 11 other `+0x8a` word stores are still unresolved.

### AI-414

- The pass first returns when the actor is not centred (`L01890`..`L01891`), then compares the edge distance `R0036` with the stop argument (`L00580`); at or below it it turns (`R0051`, `R0056`) and returns with no search.
- Past the stop test a victim different from `mover+0x7c` sets `mover+0x09 = 0xff` and `mover+0x8a = 0x00ff` and empties the static list. The test `mover+0x09 > N/3 + 1` with N = 0 passes (`L01895`), and `0x00ff > 5` selects the full search `R0053(..., flag 1)` at `L01894`.
- A fresh mover has `mover+0x7c = 0` (`AI-413`), so the first centred out-of-stop pursuit pass of any victim takes this branch. A mover off its cell centre never reaches it on that call: `L01890`..`L01891` calls `R0039` and returns, with no search and `mover+0x8a` unchanged. The `> 5` test is against `ctx+0x585c8`, shipped 5 (`MOVE-PARAM-006`).
- Emulation of the original bytes over stop distances 1, 5 and 7, three directions and Chebyshev distances 1 to 12 (108 cases): every pass that reached `L06116` ran the full search at `L01894` and never the rebuild at `L13295`; passes at or within the stop distance ran no search.

**Confidence.** High for the branch after the centre and stop guards: it is a cited instruction sequence and the emulation of the routine bytes agrees in all 108 centred cases.

**Unknown.** A pursuit started while `mover+0x7c` already holds the same victim (an earlier pursuit of it that left the field set) does not reset; `AI-415` governs it.

### AI-415

- With `mover+0x7c` equal to the victim, a re-search happens when `mover+0x09 > N/3 + 1` (N the static count). When `mover+0x8a` is greater than 5 it is the full search; otherwise `L13288`..`L01898` frees the static list and appends one node holding `mover+0x76`, the end of the previous route, and `L13287` then stores count 1.
- After a full search `mover+0x76` is the tail node's cell (`L13296`..`L13297`), and `mover+0x74` and `mover+0x8c` hold the victim's cell at that pass (`L13298`, `L13299`). The rebuild does not read the victim's current cell.
- The route-end match: when the actor's cell equals `mover+0x76` and `mover+0x8c` equals the victim's current cell (`L06963`, `L06964`), the pass increments `mover+0x09`, stores `mover+0x7c = 0` and returns without the stepper (`L13300`..`L06965`). The next pass, if the pursuit order continues and the mover is centred and outside the stop distance, sees a changed victim and runs the full search (`AI-414`).
- Emulated walk (`AI-418` set-up, start Chebyshev 12, reach 3): full search at the first pass (count 7), a second full search at (137,128) after three cell moves, on the fourth stepping pass (count 4), a rebuild at the next re-search (count 1), the route-end match on the route end, and a full search on the following pass.

**Confidence.** High for the branch structure, read whole. Medium for the emulated sequence, which depends on the harness's step model.

**Unknown.** How often play reaches the rebuild with a moved victim; the near search then aims at the old route end while picker B uses the victim's current position (`AI-373`).

### AI-416

- `R0054` tests the dynamic count `actor+0x184` (`L13301`); at 0 it calls `R0249(actor)` and returns (`L06979`..`L13302`). `R0249` writes only `mover+0`: it snaps to `mover+1` when the folded arc is below the rate `mover+0x0a`, otherwise moves `mover+0` by the rate in the direction that reaches `mover+1`.
- Emulation of `R0054` over six facing and rate cases compared every byte of the actor, mover and position record and the whole occupancy plane: only `mover+0` changed, and not at all when `mover+0 == mover+1`.
- The rest of the same pass, read in `R0043`: `mover+0x78` is incremented (`L13303`); when the near search ran, `mover+0x94` and `mover+0x96` take the victim and actor cells, `mover+0x09` is incremented and `mover+0x78` cleared (`L01905`..`L01906`). `mover+0x90` has no writer in `R0043`, `R0054` or `R0249`; its two writers in the movement range are in `R0178` (`L06947`, `L00457`).
- `R0042` then finds the actor centred, stores no progress 3, and stores `actor+0x54 = 1` (`L13304`..`L06119`); the next sub-tick runs the pursuit row again (`AI-417`). When the empty list came from a kind-1 near search, `L00162` has already stored `mover+0x98 = 1` and the executor epilogue of the same call cancels the order (`AI-373`, `AI-ROUTE-045`).

**Confidence.** High: both routines are short and read whole, and emulation of their bytes shows no other change.

### AI-417

- `R0193` runs once per sub-tick and calls each actor's slot `R0037`, which calls the executor `R0016` once (`AI-371`, `SESS-TICK-006`; 62 ms at the default speed index, `SESS-CLOCK-005`).
- In the executor, progress 3 calls the step routine `R0039` (`L06938`); when the actor is then centred it clears progress and jumps to the epilogue (`L13305`..`L13306`), so the arrival call runs no pending arm. The pursuit rows run only at progress 0 (`AI-350`), and row 5 reaches `R0043` once through `R0042` (`L06113`).
- A pass whose stepper starts a step calls `R0047`, which adds the step vector `mover+0xb0`, `mover+0xb1` to the sub-cell position; `R0042` then finds the actor off centre and stores progress 3 (`L00107`).
- So between two passes of a stepping pursuer lie the transit calls of one cell, ending with the arrival call. `mover+0x78` counts centred out-of-stop passes (`L13303`) and is cleared after a near search (`L01906`); `mover+0x09` counts near searches (`L01919`) and route-end matches (`L02020`) and is reset after a full search or rebuild (`L01900`). Both are bytes and wrap. A turning-only pass with an unchanged victim and a non-empty dynamic list increments `mover+0x78` and leaves `mover+0x09`.

**Confidence.** High for the executor and wrapper structure and the two counters' writers. Medium for the rate of passes per cell stepped: a pass whose stepper only turns (`AI-416`, or facing unequal at `L07009`) is repeated each sub-tick, and the number of transit calls per cell is the step speed (`MOVE-094`), not measured here.

**Unknown.** The second executor caller `R0038` (no reference found, `MOVE-GATE-039`). Catch-up sub-ticks of the pacer run the same body (`SESS-CLOCK-005`).

### AI-418

- Set-up (`q7_band.py`): a victim at a cell whose Chebyshev rings 1 to 4 carry a blocking bit of the mover mask in both planes, one or three attackers (side 1, mask 0x41, rate 0x20) in a column at Chebyshev 7 or 12, stop distance = reach 3, 5 or 6, 100 sub-ticks.
- Original bytes run: `R0043`, `R0053` with pickers A and B, `R0055`, `R0054`, the turn, direction and distance leaves, `R0041` and the occupancy writers. Modelled: the step (a transit of 6 sub-ticks ending with the arrival release and the dynamic-list free of `R0039`), `R1088`, `R0047`, the allocator, and the order layer (the harness stops an attacker at `mover+0x98 = 1` or in position).
- Chebyshev 7, every reach and group size: the first full search fails (picker A radius `(7>>2)+3 = 4`, `AI-394`), and `L00161` stores `mover+0x98 = 1` on the start cell.
- Chebyshev 12: the first search reaches ring 5 (radius 6). Reach 3: the attacker walks to ring 5; there the route-end match clears `mover+0x7c`, the next full search from distance 5 has radius 4 and fails, and `L00161` fires on the ring-5 cell (sub-tick 47 under the 6-sub-tick step). Reach 5: in position on ring 5. Reach 6: in position on ring 6. Each member of the group of three ends on its own row's cell.
- After `mover+0x98 = 1` the executor epilogue cancels the order and runs `R0004` (`AI-ROUTE-045`); the harness does not model that, so the cell at 100 sub-ticks is the cell at the flag.

**Confidence.** Medium: the searches and pickers are the original bytes, but the step model, the occupancy bit standing for impassable terrain and the absent order layer are the harness's.

**Unknown.** What the attacker does after the cancel within 100 sub-ticks (reacquisition by `R0004` and the group AI once per full tick); groups whose members' routes cross; a native run.

## Script Follow and what ends it

Evidence is a static read of `rom.exe` (one image on both lawful installs). `ord` is `actor+0x158`; `grpAI` is the group's AI block (`AI-ORDER-031`). The population is the routines named below; the per-actor machine `R0008` was not re-read. `EXP-0511` was allocated ids `437`..`442` of `claims/ai.md` and spent `437`..`438`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-437 | Script group command 15 (Follow) calls `R0176`: group order 0 from a register, the named unit `actor+0x50 = 0xc`, every other member `0x11` with subject `ord+0x10` and range `ord+0x70` (parameter, else 3). | High | ● active | [EXP-0511](../experiments/EXP-0511-m100-servant-amulet/) |
| AI-438 | A Follow is replaced by a join, teardown, an empty group or the mission-end reset, and suspended or replaced by a later group command; distance does not end it, and no writer was found for the subject's death. | High / Medium | ● active | [EXP-0511](../experiments/EXP-0511-m100-servant-amulet/) |

### AI-437

- The script group-command arm `R0156` loads the sub-command from `rec+8`, subtracts 1, bounds it at 0x10 and jumps through the 17-entry table at `L00408` (`L01842`..`L01843`). Entry 15 is `L01128`, which passes `rec+0x30` (the named unit) and calls `R0176` (`L01129`). Entry 4 is `L00555`, which calls `R0108` (`L13762`).
- `R0176` (142 instructions, read whole) first clears `ord+0x50` and `ord+0x38` of every member. It then stores `grpAI+0x20 = 0` from `bl` (`L00437`), a register-form store outside `AI-ORDER-031`'s immediate sweep.
- On the second walk, the named unit gets `actor+0x50 = 0xc` (`L01118`). Every other member gets `actor+0x50 = 0x11` (`L01119`), `ord+0x10 =` the named unit (`L01120`), `ord+0x08 = 0`, and `ord+0x70 =` the range byte, or 3 when it is zero (`L01106`, `L01109`).
- `0x11` is the follow arm (`AI-FOLLOW-112`); it has no leash and no distance end.

**Confidence.** High. The dispatch table is read out of the PE (`groupcmd-table.tsv`), and each store is a named instruction in a body read whole.

### AI-438

The writers that end or replace a Follow, by the field they write:

1. A later group command for the same group runs its own setter through `R0156`. The Move entry `R0108` writes group order 0 or 4 (`AI-ORDER-031`). Its member stores were not read in this experiment. Under order 4 the group runs `R0154` instead of the members' states; under order 0 the members' states run, so a member still at `0x11` would resume the follow unless the Move entry rewrote it (`AI-ORDER-010`).
2. A join (instants 19 and 22, `R0064`, `PARTY-JOIN-025`) reaches `R0063`, which writes `actor+0x50 = 0xc` (`L00150`) and group order 3 (`L00153`) (`AI-STATE-043`, `AI-ORDER-031`).
3. Teardown `R0208` writes `0x10` (`AI-STATE-043`).
4. An empty group is set to group order `0xff` (`L00300`, `AI-ORDER-031`).
5. The mission-end kept-actor reset `R1380` writes `actor+0x54 = 0` and `actor+0x50 = 0` (`L13763`, `L13764`).

The withdraw tails `R0100` and `R0107` (read whole) store only `ord+0x08 = 1` and `ord+0x0a` (`L13765`/`L13766`, `L13767`/`L13768`), or they call `R0022`. They write no `actor+0x50`, no `grpAI+0x20` and no `ord+0x10` (`stores.tsv`). Distance has no end condition (`AI-FOLLOW-112`). Nothing was found that clears `ord+0x10` when the subject dies (`AI-FOLLOWDEATH-119`).

**Confidence.** **High** for writers 2 to 5 and for the withdraw tails' store set. **Medium** for writer 1: the group order it writes is read, its effect on the members is not. **Medium** for the death clause, a lower bound bounded to the searches of `AI-FOLLOWDEATH-119` and the bodies read here. **Medium** for completeness: `AI-STATE-043`'s sweep is Medium, and the callees of the withdraw tails (`R0022`, `R0104` and others) were not read.

**Unknown.** Whether `R0108` rewrites the members' `actor+0x50` or `ord+0x10`, and so whether a later Move ends the follow or only suspends it. `R0003` (spell apply) also calls `R0063`; which group it passes was not read. Behaviour when the subject leaves the map was not read.

## Heading before an act

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-443 | Before an attack (executor rows 5, 6) or a cast at an actor (row 8) a unit turns to `R0051(actor, target)`, the byte the act gate compares; rows 5 and 6 call it with the gate's pair, and row 6 can leave before the gate and the turn. | High / Medium | ✔ promoted | [EXP-0522](../experiments/EXP-0522-target-heading/) |
| AI-444 | `R0051` heads between the two centres, fine point plus (size-1)*128 in 16 bits: eight bytes from the axis signs and 2:1 tests, an exact 2:1 to the diagonal, the zero vector 224; replayed on 84,842 inputs. | High | ✔ promoted | [EXP-0522](../experiments/EXP-0522-target-heading/) |
| AI-445 | An attack on a Building heads by `R0051` toward the position record `obj+0x10`; no rectangle or size byte enters, and for vtable `R0492` the token size is 1 (`UNIT-STRUCTREACH-063`). | High / Medium | ✔ promoted | [EXP-0522](../experiments/EXP-0522-target-heading/) |
| AI-446 | Before a cast at a cell (row 9) or a structure use (row 0xf) the walk's stop arm turns the unit to `R0089(actor, cell)`: the `AI-444` law on whole-cell differences from the anchor cell, sub-cell and size ignored. | High / Medium | ✔ promoted | [EXP-0522](../experiments/EXP-0522-target-heading/) |
| AI-447 | Unit headings come from three routines and a shot's from a fourth, the 16-way `L02828`; on every nonzero fine vector in [-64,64]^2 the shot index equals the unit heading on 8,320 and is one sixteenth off on 8,320. | High / Medium | ✔ promoted | [EXP-0522](../experiments/EXP-0522-target-heading/) |

### AI-443

- Row 5 (`L00101`) calls the gate `R0041(actor, ord+0x0c, actor+0x12c)` at `L00743`. On failure it calls `R0042(actor, ord+0x0c)` at `L00106`, which calls `R0043(actor, ord+0x0c, ord+0x14)` at `L06113`.
- Row 8 (`L00108`) skips the gate when the target `ord+0x28` is the caster (`L06104`). Otherwise it calls the gate with `ord+0x28` and reach `ord+0x14` at `L00742`; on failure it calls `R0042(actor, ord+0x28)` at `L06107`.
- `R0043` returns through `R0039` when the actor's sub-cell bytes are not both 0x80 (`L01890`..`L01891`). At a centre, with the edge distance `R0036` at most the stop byte (`L00580`), it calls `R0051(actor, target)` at `L06115` and passes the byte to the turn `R0056` at `L00754`. Beyond the stop byte it routes (`AI-414`).
- Row 6 (`L00581`) calls the gate with `ord+0x0c` and `actor+0x12c` at `L00744`. On failure it calls `R0051(actor, ord+0x0c)` at `L14168` and passes the byte to `R0250` at `L13335`, which calls the turn at `L00753`: it turns in place and does not approach.
- The gate calls `R0051(actor, target)` at `L14169` and compares the byte with `mover+0` at `L00577`. In rows 5, 6 and 8 the turn and the gate receive the same actor and the same order field, so a completed turn meets the facing test while both positions are unchanged.
- Row 6's entry (`L00581`..`L14170`, `w07`) tests byte `[esi+0x4c]` against `dl` and, when that test is nonzero, compares byte `[esi+0x12c]` with `bl`. When the byte is below `bl` it reads `[[esi+0x14]+0x28]`; on a null it stores 0 to `ord+8` (`L00109`) and jumps to `L00010` without the gate (`L00744`) and without the turn (`L14168`). Every other path reaches the gate at `L14171`. Rows 5 and 6 call `R0051` with the actor and the order field the gate receives, so the heading value is one function of that pair. Which attacker class takes row 5 or row 6, and which can take the row 6 exit, is not shown: `w07` begins after the loads of `dl` and `bl`.
- `R0051` has 15 direct callers and no dword reference in any section (`s01`). Besides the gate, the approach and row 6: acquisition `L01750`, `L01751` (`AI-ACQUIRE-002`); the Acid Stream and Teleport cell `L05167` (`MAGIC-238`); reacquisition `L14172` (`AI-361`); `L14173`, followed by a call of `R0250` at `L14174`; the face-and-mark arm `L14175`, turned at `L13267`; the group scorer `L14176`, `L14177` (`AI-COST-071`); the Prismatic Spray selector `L14178`, `L14179`, `L14180` (`MAGIC-242`); and the wrapper `R2199` at `L14181`.
- EN and RU `rom.exe` have one SHA-256, so every clause holds for both.

**Confidence.** **High** for each named call, argument and branch: `w06`, `w07`, `w08`, `w09` and `w14` list them. **Medium** that rows 5, 6 and 8 are the only routes that turn a unit before an attack or a unit-target cast. Population of `R0043` callers (`s02`, four): `L06113` (rows 5 and 8 approach) and `L06939` (read in `w07`: target `ord+0x18`, reach `ord+0x14`, a different order field) are read; `L14182` and `L14183` were not read. The reacquisition pair `L14173`/`L14174` is placed by adjacency and its body was not read; pending order 2 (`L06163`) installs action 3 with no gate and no turn (`AI-411`).

**Unknown.** Which attacker class (melee, ranged or neither) takes row 5 or row 6, and which can leave row 6 without turning. The writers of `[esi+0x4c]` and of `actor+0x12c`, and the values of `dl` and `bl` at `L00581`, settle it; none was read. Whether `R2199` is reached by a computed call; a direct `CALL` sweep and a raw dword scan find none.

### AI-444

- The centre helpers `R0862` and `R0863` (`w02`) return, for an object's position record `obj+0x10`, the fine coordinate `(cell << 8) + sub` from bytes `+0`/`+4` (x) or `+1`/`+5` (y) (`R0165`, `R0166`, `w03`), plus `(size - 1) << 7`, where size is the low byte of `obj->vt+0x1c` (`L14184`..`L06158`). The sum is truncated to 16 bits (`L14185`, `MOV ax, si`).
- `dx` = centre x of the target minus centre x of the actor (`L14186`); `dy` likewise (`L14187`), with y growing downward. With `a = |dx|` and `b = |dy|`, the branch tree `L14188`..`L06160` and the tail `L06161`..`L06092` give:

  | quadrant | `b > 2a` | otherwise | `a > 2b` |
  |---|---|---|---|
  | `dx > 0`, `dy <= 0` | 0 | 32 | 64 |
  | `dx > 0`, `dy > 0` | 128 | 96 | 64 |
  | `dx <= 0`, `dy > 0` | 128 | 160 | 192 |
  | `dx <= 0`, `dy <= 0` | 0 | 224 | 192 |

  The middle column applies when neither strict test holds, so `b = 2a` and `a = 2b` give the diagonal. A vertical vector gives 0 or 128, a horizontal one 64 or 192. The zero vector falls in the last row's middle column and gives 224; it is not a separate branch.
- The branch tree produces a 16-way index with boundaries at slopes 1/2, 1 and 2; the tail adds 1 to a non-zero index, clears bit 0 and shifts left by 4, so the slope-1 boundary vanishes and eight bytes remain.
- The difference is taken between the two footprint centres. The anchor is the record's cell; a size-n footprint's centre lies `(n-1)*128` fine units right of and below the anchor's fine point. A centred size-1 actor sees a size-2 target anchored on its own cell at `(128, 128)`, heading 96. A size-4 actor sees a size-1 target anchored one cell right and down at `(-128, -128)`, heading 224 (`a680-footprint-near.tsv`, `a680-wrap.tsv`).
- There is no clamp and no rounding to cells: a sub-cell offset moves the heading only through the signs and the 2:1 tests.
- A centre past 65,535 wraps: a size-3 target anchored at cell x 255, sub 200, has centre x 200, so an actor at cell x 10 heads 192 toward it (`a680-wrap.tsv`).
- Replay populations, each run of the original instructions equal to the table above (`summary.tsv`): cell offsets in [-12,12]^2 at sub 128, size 1 (625); offsets in [-2,2]^2 with sub bytes {0, 1, 127, 128, 129, 255} on each of four axes (32,400); offsets in [-12,12]^2 with 8 seeded sub draws each (5,000); sizes {1,2,3}^2 over offsets in [-6,6]^2 (1,521); vectors on and one unit beside slopes 2, 1/2 and 1 out to 65,280 fine units in four quadrants (8,544); axes and the zero vector (40); wrap and size-0/255 tokens (6); 20,000 seeded draws over every cell and sub byte with sizes 1 to 3; the 16,641 comparison vectors of `AI-447`; the 65 calls of its slope sweep (`shot-vs-unit-sweep.tsv`). The list totals 84,842. The [-4,4]^2 part of the 625-cell grid equals `AI-361`'s 81 outputs.
- `R0051`, `R0862`, `R0863`, `R0165` and `R0166` use no floating-point instruction: no x87 or SSE mnemonic appears in `w01`..`w03`.
- EN and RU `rom.exe` have one SHA-256.

**Confidence.** **High.** The three bodies are complete (`w01`..`w03`; `R0051` is 106 instructions whose only calls are the four centre calls), and the instruction replay equals the integer law on every input above. The boundary vectors out to 65,280 exclude a cell-based difference, a clamp and a magnitude limit; the axis and zero rows exclude a special axis or no-direction case.

**Unknown.** Which sizes and coordinates occur in play, and whether any shipped map places a footprint centre past 65,535. A census of the sizes and anchor cells in the shipped placements, with the size routine of each class, settles it.

### AI-445

- A Building is a rows 5 and 6 target like an actor (`UNIT-STRUCTREACH-063`), so its heading is `R0051(actor, building)` (`AI-443`). No listing of this experiment shows a Building target entering rows 5 and 6.
- `R0862` and `R0863` are called directly, not through the vtable, on the target pointer (`L06155`..`L06156`). They read `obj+0x10` and `obj->vt+0x1c`. `UNIT-STRUCTREACH-063` gives the Building vtable `R0492` with `+0x1c` = `R1827`, which returns 1, so the Building's centre would be its position fine point with no footprint term. This experiment did not list `R1827` or the vtable, although the preregistration promised that re-read.
- No instruction in `w01` or `w02` reads `+0x60`, `+0x61` or another field of the target besides `+0x10` and the vtable.

**Confidence.** **High** that `R0051` reads `obj+0x10` and `obj->vt+0x1c` and no other field of the target, so the rectangle does not enter: the bodies are complete (`w01`, `w02`). **Medium** that the token size of a Building is 1 (reused from `UNIT-STRUCTREACH-063`; `R1827` and the vtable were not re-read here, and other structure classes were not read) and that `+0x10` is a Building position record laid out like an actor's (that layout was not read).

**Unknown.** Which cell of the Building's rectangle its position record names, and its sub-cell bytes, so the heading relative to the drawn building.

### AI-446

- Row 9 (`L05042`) calls the point gate `R0086(actor, ord+0x3c, ord+0x14)` at `L06108`. On failure it jumps to `L14189` and calls `R0087(actor, ord+0x3c, ord+0x14)` at `L06112`. Row 0xf (`L00195`, `AI-STRUCTUSE-306`) does the same with `ord+0x0a` (`L13336`, `L14190`..`L06112`).
- `R0087` calls the walk `R0178(actor, cell, reach)` at `L06117` (`w15`). The walk takes the larger axis distance between the position word `+2` and the cell (`L14191`..`L14192`). At a cell centre with that distance at most reach and reach non-zero (`L14193`..`L07003`) it calls `R0089(actor, cell)` at `L14194` and the turn at `L01933` (`w16`). Off centre it calls `R0039` and returns (`L06897`..`L06898`).
- The point gate compares `mover+0` with `R0089(actor, cell)` at `L06109` (`w06`): the turn and the gate use one routine and one pair.
- `R0089` (`w04`, 99 instructions, no call): `dx` = the cell argument's low byte minus position byte `+0`, `dy` = its bits 8..15 minus position byte `+1` (`L14195`..`L14196`). Its branch tree and tail (`L14197`..`L14198`) are those of `AI-444`, so the `AI-444` table holds with whole-cell `dx`, `dy`. The sub-cell bytes, the position word `+2`, the argument's upper 16 bits and the footprint size do not enter. A target cell equal to the anchor cell gives 224.
- Replay: every cell difference in [-255,255]^2 (261,121) with seeded sub-cell bytes and `+2` word, and 20,000 seeded draws with a random upper word, all equal the model.
- `R0089` has six direct callers (`s01`): the wall direction `L14199` (`MAGIC-AREACELL-039`), the point gate `L14200`, the stop arm `L14194`, the wait arm `L14201`/`L14202` (`MOVE-WAIT-008`), and the wrapper `R2200` at `L14203`, which has no direct caller and no dword reference (`s02`).
- A size-n caster heads from its anchor cell, where `AI-444` would use its centre. EN and RU `rom.exe` have one SHA-256.

**Confidence.** **High** for the routine, its law and the call sites named. **Medium** that the stop arm is the only turn toward the cell before the cast: the walk was read through `L14204` and at its wait arm, and its reach-0 branch `L07004` was not read.

**Unknown.** The reach-0 arm, and whether the position word `+2` can differ from bytes `+0`/`+1` while the walk runs, which would let the distance and the heading name different cells.

### AI-447

- Four routines produce a direction, each its own body with no call to another (`bodies.tsv`): `R0051` between actor centres (`AI-444`); `R0089` between anchor cells (`AI-446`); `R1356` from the actor's fine point to a cell centre, by signs alone (`MOVE-115`); and the projectile helper `L02828` (`ANIM-138`), 16 values with boundaries at slopes 1/4, 3/4, 4/3 and 4. `L02828` has two direct callers, `L13633` and `L02827` (`s01`).
- `R0051` and `L02828` were both run on every fine vector in [-64,64]^2, in their own coordinate arguments. The shot index equals the unit byte divided by 16 on 8,320 nonzero vectors and differs by one sixteenth on 8,320; no vector differs by more (`shot-vs-unit.tsv`). The zero vector gives 4 for the shot and 224 (14) for the unit.
- Every comparison in both helpers is homogeneous in `|dx|` and `|dy|`, so agreement depends on the direction alone while the products do not wrap. For `0 < b <= a` the two agree when `b <= a/4` or `b > 3a/4`, and differ by one sixteenth when `a/4 < b <= 3a/4`. For `b > a` they agree when `b < 4a/3` or `b >= 4a`, and differ when `4a/3 <= b < 4a`. At `dx = 64` the shot turns from 4 to 3 at `|dy| = 17` and the unit from 64 to 32 at `|dy| = 32` (`shot-vs-unit-sweep.tsv`).
- EN and RU `rom.exe` have one SHA-256.

**Confidence.** **High** for the routine identities and for the comparison on the replayed vectors and the homogeneous extension. **Medium** for what a player sees: the shot's vector runs from its launch point (`MAGIC-263`, `ANIM-139`) and the unit's between centres, so in play the two are not measured on one vector.

**Unknown.** Whether shot coordinates and unit fine centres share one unit and origin; the comparison here holds only for equal difference vectors.

## Group and order blocks at mission start

`grpAI` is the 80-byte block at `Group+0x3c` (`AI-GROUP-009`); `ord` is the 148-byte order block at
`actor+0x158`. Evidence is a static read of `rom.exe` (one image on both lawful installs, SHA-256
`942e9b72…`); no process was run.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| AI-449 | The order constructor zeroes 148 bytes but sets `+0x71` 1; the AI constructor zeroes 80 but sets `+0x45` 1; a map Group comes from the spawner's miss arm; the join walk builds a Group only for a Player with none. | High / Medium | ✔ promoted | [EXP-0524](../experiments/EXP-0524-group-start/) |
| AI-450 | The load-time guard setter `R0125(grp, 0)` writes AI `+0x20` 1, the members' centroid `+0x24`/`+0x28` and `+0x00`, `+0x2a`..`+0x2c`, `+0x2d` = `+0x38` = max(`+0x2c`, MinimalGuardRange) and `+0x39` 0. | High | ✔ promoted | [EXP-0524](../experiments/EXP-0524-group-start/) |
| AI-451 | The Stand Ground setter `R0063` stores only byte `+0x20` of the AI block (0, 0, then 3); every other AI byte of a Group it sets keeps its earlier value. | High | ✔ promoted | [EXP-0524](../experiments/EXP-0524-group-start/) |
| AI-452 | The trigger binder `R0067` stores no Group-AI field but dword `+0x48` = 1, at eight sites that resolve a unit's Group or a Group; its body has no order-block access. | High / Medium | ✔ promoted | [EXP-0524](../experiments/EXP-0524-group-start/) |
| AI-453 | After placing the carried members the join walk reads key `Humans`; each entry becomes a Human or Unit placed into the head Group, which then gets `R0063`; a nonzero read result skips both. | High / Medium | ✔ promoted | [EXP-0524](../experiments/EXP-0524-group-start/) |
| AI-454 | The first sub-tick runs no Group dispatch and no activity rebuild, and the executor's tail stores no order byte for an idle member while Mover `+0x98` is 0, its value at placement. | High / Medium | ✔ promoted | [EXP-0524](../experiments/EXP-0524-group-start/) |

### AI-449

- `R0286` (`w01`, whole) clears `0x25` dwords (`L00962`), stores bytes `+0x08`, `+0x09`, `+0x4c` 0 and `+0x71` = 1 (`L14225`), and puts a new `0x1c`-byte word list (vtable `L08172`, grow-by 10) at `+0x90` (`L14226`), or 0 when the allocation fails (`L14227`). Direct callers (`s01`, operand sweep plus raw dword scan): `L14228` (carry importer, `PARTY-LOSS-006`), `L14159` (join walk, `MOVE-122`), `L14229` (actor constructor, `AI-LOAD-099`), `L14230`; no stored address.
- `R0150` (`w01`, whole) clears `0x14` dwords, stores bytes `+0x08`, `+0x09`, `+0x20` 0 and `+0x45` = 1 (`L14231`), and puts a new word list at `+0x4c` (`L00291`). Its one caller is `L14232` in the Group constructor `R0149`.
- `R0149` (`w02`, whole) builds the base and the `+0x20` list, stores `+0x40` = `+0x44` = 0 and `+0x3c` = the new AI block; it writes no AI field itself. `s01` finds 13 direct callers and no stored address. `R0152` (`w02`) detaches the actor from its old Group, appends it, stores `actor+0x70` and `Group+0x44` = `actor+0x14`; it has no AI access.
- Map placements: the type-6 spawner joins each unit to the Player's Group whose `+0x1c` equals the record's group id, or builds one (`L00294`..`L00296`, `w03`; `AI-GROUP-009`, `SAV-GRPFIRSTSAVE-579`).
- Carried members: the join walk `R0065` builds no Group for a member (bounded to the routines read: the walk head and its member loop). Its head (`w10`) takes the first Group of `player+0x24` (`L14233`); when there is none it logs `Oops - player has no groups` (string `L07360`, `t01`) and builds and appends one (`L14234`..`L14235`). While that Group's `R0197` returns 0 and the list holds more than one Group (`L14236` above 1) it calls `R1582` on it and takes the new first Group (`L14237`..`L07334`). Each member keeps its `actor+0x70` Group, the reused Player's (`PARTY-PERSIST-028`). Callers of `R0149` not read: `L14238` (the Group build in the carry arm `L07346`, a planned window never read), `L07379`, `L14239`, `L14240`, `L14241`, `L14242`, `L07279` (the Group LOAD), `L14243`, `L14244`, `L14245`, `L14246`.

**Confidence.** **High** for the constructor stores, the caller sets and the walk head's branches (complete bodies, `CALL rel32` targets from the bytes). **Medium** for the reading of `R0197` as an empty test, `R1582` as removal from the list and `L14233` as the first element: their bodies were not read. **Medium** for "no mission-start routine builds a Group for a carried member": it covers the routines read (the spawner arm and the join walk), and eleven of the thirteen direct callers of `R0149` were not read, the carry arm's `L14238` among them.

**Unknown.** The routine around `L14230`. Whether a Group the walk head removes is freed. The unread callers of `R0149` named above, of `R0063` (`L01873`, `L14247`, `L01874`) and of `R0125` (`L00378`, `L00596`, `L00949`, `L00388`, `L00950`) as to whether any runs at a mission start.

### AI-450

- `R0125(grp, n)` (`w06`, whole). First member walk: `R0007`, `ord+0x50` = 0, `ord+0x38` = 0 (`L14248`, `L14249`). Then `grpAI+0x20` = 0 (`L14250`).
- Second member walk: `mover+0x01` = `mover+0x00` when they differ; `R0045` only when Mover word `+0x80` is nonzero; `ord+0x14` = `actor+0x12c` (`L14251`); `ord+0x60` = 0; `mover+0x7c` = 0; `actor+0x54` = 0; `ord+0x50` = 0; `actor+0x50` = `0xb` (`L00377`); `ord+0x00` = the position word `+0x02` when `R0040` finds the actor centred, else Mover word `+0x06` (`L00829`, `L00830`); `ord+0x08` = 0 (`L00945`).
- Then `grpAI+0x20` = 1 (`L00314`) and `R0140(grp)` (`w09`, whole): for a nonempty Group it sums the members' fine coordinates (`R0165`, `R0166`) and divides each by the low byte of the member count; it stores the cell word `+0x28` from the two quotients' high bytes (`L14252`) and dword `+0x24` = (y << 16) + x (`L00353`). A second walk stores `+0x2a` = max `R0167(member cell, +0x28)`, `+0x2b` = max `actor+0xa5`, `+0x2c` = max (that distance + `actor+0xa5`) in bytes (`L00357`..`L00359`). An empty Group returns at `L14253` with no store.
- Then word `+0x00` = word `+0x28` (`L14254`), `+0x2d` = `+0x2c` (`L00360`), `+0x38` = `+0x2c` (`L00838`). With `n` = 0, when `+0x2d` is below `session+0xa9b8` both take that value (`L00844`, `L00361`); with `n` nonzero, when `+0x2d` is below `n` both take `n` (`L00843`, `L00839`). `+0x39` = 0 on both paths (`L14255`, `L14256`). `session+0xa9b8` is `MinimalGuardRange`, 8 in the shipped `ai.reg` (`AI-RADIUS-014`).
- The stance walk calls it with `n` = 0 for every Group of a Player other than the first type-5 record (`Player+0x04` 1) in single player (`L14257`..`L00384`, `w05`; `AI-AUTHOR-015`, `AI-GATE-100`).
- `R0007` (`w08`, whole) stores `mover+0x01`, `ord+0x14` = `actor+0x12c` (`L00155`), `ord+0x60` = 0, `mover+0x7c` = 0 and `actor+0x54` = 0, and calls `R0045` under the same test.

**Confidence.** **High**: the setter, the centroid routine and `R0007` are read whole and every store is a cited instruction. The restart-slot census recomputes `+0x24`, `+0x28` (hence `+0x00`) and `+0x2a` from the members' cells and matches them on 133 of 133 guard Groups; `+0x2b` is not recomputed and `+0x2c` is only bounded; `+0x2d`, `+0x38`, `+0x39` and `+0x45` follow from the stored values (`SAV-1218`). The Chebyshev rule for `+0x2a` rests on that fit (`R0167` was not read): Medium.

### AI-451

- `R0063(grp)` (`w07`, whole). Its member walks store `ord+0x14`, `ord+0x60`, `ord+0x50`, `ord+0x38`, `ord+0x00` (`L00151`) and `ord+0x08` = 0 (`L00152`), `actor+0x54` = 0 and `actor+0x50` = `0xc`, as `AI-351` lists.
- Its AI stores are three byte stores to `+0x20`: 0 at `L00147`, 0 at `L14258`, 3 at `L00153`. It loads `[grp+0x3c]` only for these, and calls neither `R0140` nor any other routine with the Group but the member walkers `R0069`, `R0070` and `R0071`.
- Calls on the mission-start path (`s01`): `L00154` in the stance walk for the Groups of the first type-5 record (`AI-AUTHOR-015`); `L14161` in the join walk, once per placed member on its `actor+0x70` (`MOVE-122`); `L14259` in the join walk's tail on the head Group (`AI-453`).

**Confidence.** **High**: one routine read whole; the restart-slot census holds order 3 and no other changed AI byte on the one map Group and the twelve carried Groups of Player 1 (`SAV-1219`).

**Unknown.** The bodies of the three member walkers, which receive the Group pointer (`AI-351` reads `R0069`, `R0070` and `R0071` whole).

### AI-452

- `R0067` was decoded from its entry to its `RET 8` at `L14260`, 1,428 instructions (`w05`, one window and four extensions). After the stance loop (`L01802`..`L14261`) it is the trigger binder (`TRIG-BINDORDER-101`).
- Its Group-AI stores are eight `MOV dword [x+0x48], 1` through a just-loaded `[g+0x3c]`: `L03530` and `L03531` (a unit reference, through `actor+0x70`); `L01828` and `L14262` (a Group reference); `L14263` and `L14264` (a unit reference) and `L14265` and `L14266` (a Group reference), these four in the binder's second record pass (`L14267`..) and only when that record's opcode `+0x40`, copied to `[ebp-0x1f8]` at `L14268`, is not 1 (`TRIG-BINDORDER-101`). Each needs a nonzero resolved reference and, for a unit, a nonzero `actor+0x70`.
- The body loads no `[x+0x158]` and makes no other store through a loaded Group-AI pointer. Its calls into the AI module are the two stance setters (`L00154`, `L00384`) and the trigger record helpers `R0506`..`R2178`.
- `AI-ACTIVITY-324`: a Group whose `+0x48` is nonzero counts every represented member as active in each rebuild.

**Confidence.** **High** for the eight stores and their gates, and for the absence of any other AI or order-block access in the body's own instructions (linear decode, every hit of `+0x3c]`, `+0x70]` and `+0x158]` read). **Medium** that the binder's callees write no AI or order field: `R2190`, `R0421`, `R0507`, the `00538xxx` helpers, `R0288`, `R1388`, `R0865` and `L14269` were not read.

### AI-453

- `w11` (`L14161`..`L08258`, the walk after its member loop). `L06602` calls `R0475(this+0x44, "Humans", list)` (string `L06600`, `t01`). A nonzero result jumps to the exit `L14270`, past the `R0063` call.
- Otherwise, per entry: a name, cut at `#` (`0x23`, `L14271`) into a name and a number; `new(0x1e8)` + `R0497` (Human); if its type word `+0x0e` is 0, it is deleted and `new(0x198)` + `R0501` (Unit) is tried; a second 0 deletes it and skips the entry. With a number, `R0907`; with `this+0x134` nonzero, `R1562`. Placement `R1145`: at the walk's start cell with radius 8 when the entry's first field is -1, else at the entry's three fields. A failed placement deletes the actor.
- A placed actor: `R0411` on `[L00240]`, `actor+0x14` = Player, AddTail to `player+0x20`, `R0152(G, actor)` with `G` the walk head's Group (`AI-449`); then `R0044(actor)` when `this+0x128` is 0 (`AI-350`: `ord+0x08` = 0, `actor+0x50` = `0xc`, the post), else `R0308(actor, player+0x34, 0)`; then `R0993(actor)`.
- When the list is empty, `R0063(G)` at `L14259`.

**Confidence.** **High** for the branch structure and the calls (one range read whole). **Medium** for which `R0475` result means an absent key: its body was not read.

**Unknown.** The `Humans` key's source object `this+0x44`, the meaning of `this+0x128` and `this+0x134`, and the stores of `R0308`, `R0907`, `R1562` and `R0993`. No restart-slot save in `SAV-1218`'s population holds a spawner-pattern member in Player 1's first Group.

### AI-454

- From (0, 0) the mission start's one stepper call runs no phase slot (`SESS-086`); the AI slot `R0076` holds the Group dispatch (`AI-TICK-008`) and the activity rebuild (`AI-ACTIVITY-324`), so neither runs before the restart save.
- The executor `R0016` sends an idle member (`ord+0x09` 0) whose Group has `grpAI+0x45` nonzero, with `ord+0x08` 0, to the shared tail `L00010` (`AI-350`). The tail (`w12`, read to `L00614`) stores nothing and jumps to `L00614` when Mover dword `+0x98` is 0 (`L14272`..`L00613`). Otherwise it clears it and, by `actor+0x50`: 1, the stand (`ord+0x50`, `actor+0x50` = `0xc`, the post, `ord+0x08` = 0); `0xa`, `ord+0x04` = 1, `ord+0x08` = 0 and, when `R0162` returns below 3, the patrol cursor `ord+0x02`; `0x17`, `ord+0x09` = `0xff`; other values `ord+0x08` = 0 and `R0004`.
- Mover `+0x98` is the constructor's 0 for both placement routes (`MOVE-119`, `MOVE-122`) and every other Mover byte is 0 in the restart-slot census (`SAV-1213`).
- With `grpAI+0x45` 0 and `receiver+0xb388` 0 the executor parks `actor+0x54` = `0x1a` and returns before any order store (`AI-350`).

**Confidence.** **High** for the tail's branch and stores (`w12`) and for the slot order (cited claims). **Medium** that nothing else in the first sub-tick writes a Group or order field: the actor tick `R0037` before `L00086`, its dispatch on `actor+0x54`, the drain `R0191` and `R0426` were not read. The census finds no AI or order byte that differs from the start path's stores (`SAV-1218`, `SAV-1219`, `SAV-1220`); a same-value write is invisible to it.
