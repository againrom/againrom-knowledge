# Orders, formation and execution

[Reference](format.md)

## Player orders

Nineteen opcodes, `0x14..0x26`, dispatched by `R0061`'s order space (19 direct dwords at
`L00529`; [SESSION](../session/format.md) owns the command itself). **Four are empty** — `0x15`, `0x20`, `0x22`,
`0x23` — and every live arm first tests the AI manager `[L00004]` and does nothing when it is
null.

**Every order builds a new group.** Before the switch, the prologue resolves the commanded actors,
allocates a fresh `0x48`-byte group onto the player's list and fills it with exactly them. The
group's AI block is constructed with **`grpAI+0x20 = 0`**, so a commanded group always *starts* at
group order 0 — which is the gate the per-actor state machine runs behind. Nothing about the load-time
1-or-3 stance survives a player order.

The orders then split by what they leave that byte at:

| opcode | dispatches to | leaves `grpAI+0x20` | per-actor `actor+0x50` |
|---|---|---:|---|
| `0x14` | `R0103` | 0 | `0x16` |
| `0x16`, `0x1c` | `R0108(col,row)` | **4** — `Move` | from the argument |
| `0x17` | `R0125` — **guard** | **1** | `0xb` |
| `0x18` | `R0060` — **aggressive** | **3** | `0xc` |
| `0x19` | `R0006(target)` | 0 | `3`, then `0xc` |
| `0x1a` | `R0157(col,row)` | **5** — `Swarm 2` | from the argument |
| `0x1b` | `R0190(target)` | 0 | `0xc`, `8` |
| `0x1d` | `R0174(col,row)` — **patrol** | 0 | `0xa` |
| `0x1e`, `0x25` | `R0011(target, spell)` | 0 | `0xd`, then `0xc` |
| `0x1f`, `0x26` | `R0012(col,row, spell)` | 0 | `0xe`, then `0xc` |
| `0x21` | `R0010(col,row)` — **pick up a sack** | *unwritten* → 0 | `2` |
| `0x24` | `R0085(building)` | *unwritten* → 0 | `0xf` |

**Four orders act through a group arm** (guard, aggressive, move, `0x1a`); **ten act by writing a
per-actor state** and leaving the group at 0, which is the only value under which
`R0023` evaluates `actor+0x50` at all. That is why patrol and the pick-up are per-actor
behaviours and guard and aggressive are group behaviours — the difference is one byte written at the
end of the order routine.

**Nine of the fourteen live arms have the dispatcher as their only caller in the image** —
`0x14`, `0x19`, `0x1a`, `0x1b`, `0x1d`, `0x1e`, `0x1f`, `0x21`, `0x24` — while the map's script
reaches `R0108`, `R0157` and `R0125` directly and **never enters the
dispatcher**.

Patrol shares `R0172` between player and script routes; the
player helper `R0174` tail-calls it (`AI-PATROL-017`). Pickup's
immediate-form actor+0x50=2 writers that reach an actor both belong to
`R0061`. Broader computed-value producers are outside that writer set.

The received object-target command (`cmd+4 == 3`) resolves signed-word id `cmd+0x0e`
against the unit collection first, then the structure collection. Opcode `0x19` forwards
that pointer to the shared attack order. The scorer must not return `0xffffff`, and the
target must not be the member itself, before the member retains it at `ord+0x0c`.
The physical branch proceeds through order 5, shared approach and action 3, copying the
pointer to `actor+0x5c`; no separate structure action or per-strike footprint lookup is
present in these bodies. This is a conditional received-command route, not proof that
every Building instance is safely admitted or that ordinary hover produces the command:
the scorer/AI still carry actor-oriented field reads on the supplied object.
Claim: `UNIT-STRUCTORDER-062`.

**Customisation.** The four empty order slots are absent implementations and free. The order
opcode's ceiling `0x26` is a compiled constant. Both `actor+0x50` and `actor+0x54` are serialized, 4 bytes
each, in the scalar run of `Unit::Serialize` (`SAV-UNITPROG-156`), so a new state value in either
needs no new wire field and is written into every save. `ord+0x08` and `ord+0x09` are serialized
too, inside a raw `0x94`-byte block, so a new field before them breaks every save the original
wrote while a new *value* in them does not.

## Formation

Every `Player` owns a `0x20`-byte settings block, allocated and constructed by the `Player`
constructor (at `L00560` it requests a `0x20`-byte allocation through `R0202`, and at `L00562` it stores the result in the `Player+0x30` field). Its **last
byte**, `+0x1f`, is the **formation mode**; the constructor writes the only nonzero default it
has, `2`. The block is serialized raw, `0x20` bytes, by `R0204`, so the mode is in every
save.

| value | behaviour |
|---:|---|
| `0` | never in formation; the group centroid is not even computed |
| `2` | in formation **iff** every member is within `AImanager+0xa824` cells (Chebyshev) of the group centroid — the threshold is a compile-time `2` |
| anything else | in formation **unconditionally**; no spread test runs |

One routine writes it, `R0080(player, mode)`, and it has two callers, both authored
surfaces:

- **player command `0x46`, sub-code 2** — the player from `cmd+0x05` through `R0079`, the
  value from `cmd+0x0e` **remapped** `0→0`, `1→2`, `2→1`, anything else `→2`, so the wire can only
  produce `{0,1,2}`;
- **trigger instant id 7**, which the shipped `Description Instants.ini` names **`Set formation`**
  and which writes the parameter **raw** — so a map may author a value the UI cannot.

What the mode decides is both halves of a group Move / Swarm-2 order: whether each member is sent
to `target + (memberCell − centroidCell)` or every member to the same cell, **and** whether the
group's rate override is written at all. [MOVE](../move/format.md) owns that; the gate is the same
local flag in both cases.

`scn:110.alm` authors `Set formation = 2`, equal to the constructor default.

## Order execution

The positive-HP actor-tick path calls `R0016` at `L00086` before dispatching
`actor+0x54`, only when the dword `actor+0x3c` is nonzero. Zero skips that call. The
opposite gate clause in `AI-405` is partially retracted and corrected by `AI-455`.
The actor's effect list and health test precede the executor (`AI-455`).

The executor distinguishes command policy, actor action, pending order and retained
progress. Their fields are separate (`AI-456`).

| field | width | executor meaning |
|---|---|---|
| `actor+0x154` | dword pointer | Mover |
| `actor+0x158` | dword pointer | order block |
| `actor+0x50` | dword | command policy evaluated by the AI |
| `actor+0x54` | dword | action dispatched by the actor tick |
| `actor+0x58` | dword | attack/cast phase |
| `actor+0x136` | byte | cycle completion |
| `ord+0x08` | byte | pending order |
| `ord+0x09` | byte | retained progress |
| `ord+0x15` | byte | strike/cast progress counter |
| `ord+0x5c` | dword | actor-target versus cell-target cast selector |

When progress is zero and both manager `+0xb388` and group-AI byte `+0x45` are zero,
the executor sets action `0x1a` and returns. Nonzero progress bypasses that activity
refusal. Otherwise it preserves actions `2` and `0xf` and clears other actions.
Effect bit `actor+0x144 & 0x100000` installs progress `4` only from progress zero
(`AI-456`).

| entered `ord+0x09` | operation | exit |
|---:|---|---|
| `0` | dispatch pending `ord+0x08` | common tail |
| `1` | action `3`; increment byte counter `ord+0x15` | counter above `2` and completion nonzero clear progress and action, then tail |
| `2` | action `0xd` for nonzero `ord+0x5c`, else `0xe`; same byte counter | same completion gate, then tail |
| `3` | call `R0039`; action `1` | cell centre clears progress and action, then tail |
| `4` | action `0x1a` | effect bit `0x100000` absent clears progress, then tail |
| `5..254` | clear progress; action `0x1a` | tail |
| `0xff` | action `0x1a`; retain progress | tail |

This is the complete byte dispatch, including default values (`AI-456`). A turn
can have action `1` and progress `0`; a crossing has progress `3`. Action `1` alone
does not identify a crossing. Its actor-body arm clears attack/cast phase
`actor+0x58` to zero (`AI-456`). Progress zero has no general centre test; centre checks
belong to individual pending paths.

A progress-`3` call that reaches the next centre clears progress and jumps to the
common tail. It runs no pending-order row on that call. The next admitted call
entered at zero can dispatch the pending order. The tail or another writer can
replace that order before this happens (`AI-459`).

| `ord+0x08` | what it does | inputs |
|---|---|---|
| `1` | walk to `ord+0x0a` | the cell |
| `2` | attack `ord+0x0c` **with no distance test** | the target |
| `4` | close on the actor `ord+0x18` | stop distance `ord+0x14` |
| `5` | pursue and attack `ord+0x0c` | facing, reach `actor+0x12c`, stop distance `ord+0x14` |
| `6` | as `5`, plus the auto-engage suppression, and turn instead of path | same |
| `7` | the sack pick-up bridge | — |
| `8` | cast `ord+0x30` at the actor `ord+0x28` | facing, `ord+0x14` |
| `9` | cast at the cell `ord+0x3c` | facing, `ord+0x14` |
| `0xa` | turn in place until `mover+0x01 == mover+0x00` | the desired facing |
| `0xb` | idle: re-face to `facing + 0x21 + (190·rand()/0x8000)`, a near-uniform byte, and only when `ord+0x54` is set or on `rand() < 0xcd` | — |
| `0xc` | walk to `ord+0x0a`, then the `actor+0x54 = 2` action | the cell |
| `0xf` | reach `ord+0x0a` within `ord+0x14`, then `actor+0x54 = 0xf` | the cell, `ord+0x14` |
| `3`, `0xd`, `0xe` | empty — the switch's own default | — |

A pending order of `0` is not a row. The index `ord+0x08 - 1` wraps above 14 and the tick ends in
the shared tail with no arm (`AI-350`). The Stand Ground setters store `0` for every member and
leave `ord+0x0a` and `ord+0x0c` as they were (`AI-351`). The heal AI stores `0`, or `8` when it
casts, for a human participant's member that has no group-assigned target (`AI-349`). A step under
way ends at the next cell centre (`AI-352`). Its unconditional strike consequence is
partially retracted: a counted strike or cast finishes its progress first only when
the common tail and other writers retain its active fields (`AI-354`).
Row `5` re-reads `ord+0x0c` on every pass (`AI-350`). Read from the code and not
observed at run time, a pursuit stored at one group evaluation therefore keeps stepping toward its
victim until the next evaluation reissues or replaces it, or until a stand or a refused route ends
it (`AI-353`).

### Commands during motion

Stop `R0088`, move `R0301` and the attack setter's common reset
replace requested facing byte `mover+1` with current facing byte `mover+0`.
They clear the action but do not directly clear progress `ord+9`, current facing,
position or the active-turn dword `mover+0xa0`. A turn loses its old requested
heading at the command-time facing; a retained crossing keeps its progress
selector (`AI-457`, `AI-458`).

| command | immediate stored order | effect on a retained crossing |
|---|---|---|
| stop | policy `0xc`, pending `0`, post equal to the command-time packed cell | crossing continues to its next centre; with no intervening replacement, pending `0` starts no further walk (`AI-457`, `AI-459`) |
| move | policy `1`, pending `1`, queued cell `ord+0x0a` | old crossing finishes before a later entered-zero call takes the new destination (`AI-457`, `AI-459`) |
| admitted non-self attack | policy `3`, queued target `ord+0x0c`, pending `0` | crossing retains progress; policy evaluation must install pending `5` before its gate/approach can run (`AI-458`, `AI-459`) |

Reset and move release a reserved cell only when `mover+0x80` is nonzero, the actor
is centred and that reserved cell differs from its current packed cell. An
off-centre command skips that release. These setters do not reverse the crossing
or stop it at the command-time sub-cell. Its final centre is the old step
destination, subject to the crossing and common-tail handling (`AI-457`, `AI-459`).

The stop's post is its packed cell when the command arrives, which can precede the
crossing's cell-boundary change. It need not equal the centre reached afterwards.
The stop does not itself order a return to that post (`AI-457`). Move installs its
pending order immediately; attack only replaces policy and queued target. The
native delay until attack policy reissues pending `5`, external writers and
command/tick interleaving remain Unknown (`AI-458`).

### Commands during a running cycle

States 3, 0xd and 0xe share phases 0 (charge), 5 (application) and 7 (recovery).
Phase 5 passes the current actor victim `+0x5c`, or cell bytes `+0x60/+0x61`,
and active Spell `+0x64` to the application routine. This body does not read the
pending order. Nonzero strike/cast progress restores the action before this
dispatch; clearing pending order alone does not cancel it (`AI-354`). Successful
damage and per-spell effects still require their own validity rules.

| command | direct setter effect | ordinary executor effect |
|---|---|---|
| admitted attack at another actor, 0x19 | command state 3, queued victim replaced, action 0, pending 0; no direct progress or active-victim store | pending 5 copies the queued victim only at progress zero and in position (`AI-355`) |
| pick-up, 0x21 | command state 2, queued cell, action 0, pending 0; no direct progress or active-victim store | row 7 at progress zero selects action 2, then completes with command state 0xc and pending 0 (`AI-356`) |
| actor cast 0x1e, cell cast 0x1f | command state 0xd/0xe, queued target and Spell, action 0, pending 0; no direct progress or active-victim store | rows 8/9 at progress zero load progress 2 and active target/Spell (`AI-356`) |
| scroll actor/cell cast, 0x25/0x26 | writes item `actor+0x68` before using the same cast setters | old cast or weapon-proc cleanup can overwrite or clear that pointer (`AI-357`) |

With the active fields retained, the running strike passes the old victim even
when the queued victim differs. For a retained cycle with recovery-zero tick T0,
the new in-position cycle can start at T0+2 if the counter/completion gate clears
progress at T0+1 and pending 5 is already reissued. There is no universal native
tick number (`AI-355`). The executor common tail can run with progress nonzero:
`mover+0x98 != 0` can lead to acquisition and an active-victim replacement before
application without resetting phase. Whether that prerequisite occurs during a
native cycle, command/tick interleaving, untraced external writers and per-spell
target loss remain Unknown (`AI-354`).

Explicit Retreat's group-policy dispatch has no progress gate. Its ordinary
pending order is skipped on an invocation that itself clears nonzero progress,
then can run on an admitted invocation entered at zero. The entry activity
predicate can newly refuse that following invocation (`AI-364`). The active
route-failure tail can instead reinstall strike progress on the clearing
invocation while retaining phase (`AI-365`). Native interleaving remains Unknown.

**`R0016` (this table's own dispatcher) never runs while an actor is dying.** The per-tick
routine's dying branch exits through its own shared tail before reaching the four instructions
that call it, on every tick from the first health-`<=`-0 tick to teardown, not only the first
(`HERO-DYINGTICK-145`). So `ord+0x08` on a dying actor is not frozen at whatever this table's own
machine last wrote it to: while a group order-5's own candidate-list gate is open, its per-member
arm keeps rewriting `ord+0x08` once every full tick regardless — for a member with no
group-assigned target, `0xb` (this table's own idle row) when the owning player is a non-human
participant, or a call into the heal AI, which stores `0`, or `8` when it casts, when the owner is
a human participant (`AI-349`); a member that does carry a group-assigned target goes through the
pursuit/engage writer (`5`/`6`/etc., this table's own rows) instead — with no test anywhere in
either arm of the receiving member's own health or death stage (`AI-332`). The arm's own
instructions write no literal `0`; the heal AI call does (`AI-349`). Group membership does not
end at the health threshold this table's own state machine reacts to: a dying actor stays a
walked member until teardown, not until death (`AI-332`).
The permitted archive's own dying population — 59 saves in 7 of the eleven pinned directories, 4
original and 44 ROM1 resaves of Againrom-produced/modified documents (`SAV-1059`) — is confined
to `{0x00, 0x0b}` for `ord+0x08`, with every `0x0b` in the resave population and every original at
`0x00`; `0x0b` matches the non-human-participant, null-target branch's own direct store,
consistent with that on the resave population only. No dying record in that corpus carries `5` or
`6`, though the mechanism argument above does not depend on that absence.

The in-position test that `5`, `6` and `8` share is **not** a centre-to-centre distance: it is the
actor's current facing equal to the 8-way direction to the target, **and** an edge-to-edge distance
that subtracts both token sizes, so two touching actors measure 1. Reach is re-read from
`actor+0x12c` at the test; `ord+0x14` is used only as the mover's stop distance, inside which the
actor turns to face and stands still.

The heading those turns aim at is the byte the facing test compares. Rows `5` and `8` turn through
the approach, row `6` turns in place, and all three call
`R0051(actor, target)` with the gate's pair (`AI-443`).

Row `6` clears pending order and leaves before the gate and the turn when actor
byte `+0x4c` has bit `4`, reach byte `+0x12c` is below `2`, and the owning Player
dword `+0x28` is zero. The complete executor fixes the entry constants missing
from `AI-443`'s bounded read (`AI-462`). Which policy writer selected row `5` or
`6`, and native occurrence of that exit, remain separate bounds.

`R0051` takes the difference
between the two footprint centres, each the fine point plus `(size-1)*128` in 16 bits, and returns one
of eight bytes: within each sign quadrant the axis byte when one magnitude exceeds twice the other,
otherwise the diagonal, so an exact 2:1 is diagonal; the zero vector gives 224 (`AI-444`). A Building target is measured toward its position record `+0x10`, not its rectangle; for the Building vtable the token size is 1 (`AI-445`; the size routine and the record's layout are Medium). Rows `9` and `0xf` turn at the walk's stop arm to `R0089(actor, cell)`, the same
law on whole-cell differences from the anchor cell (`AI-446`). A shot's flight direction is a
different 16-way helper; on one difference vector it equals the unit heading or is one sixteenth off
(`AI-447`).

Worked example — the sack pick-up, which is the whole chain in one order:

1. order `0x21` → `R0010`: `actor+0x50 = 2`, `ord+0x0a` = the sack's cell, `ord+0x08 = 0`.
2. `R0008` arm 2: not at `ord+0x0a` → `ord+0x08 = 1`, walk. **At it → `ord+0x08 = 7`.**
3. `R0016` pending arm 7 → **`actor+0x54 = 2`**.
4. the actor tick's `actor+0x54 == 2` arm loots the sack under the actor's feet.
5. next tick, arm 7 sees `+0x54` already 2 and completes: `actor+0x50 = 0xc`, `ord+0x08 = 0`.

So a unit told to pick something up ends the order **hunting**, not idle. The same relay carries
every other order that has a "walk there, then do a thing" shape.

### Setter origins, tick order and route refusal

- The script Stand Ground sub-command is its own arm: it stands each member through `R0088` and
  stores group order 3 itself, calling neither `R0060` nor `R0063` (`AI-370`).
- Within one tick the AI slot and the command-queue drain run before every member's executor pass; a
  spell application inside the actor body can run a setter after it. A setter between a pickup's two passes
  writes action 0 and pending 0, so the completion pass finds no row. Callers `R0062`, `R0065`,
  `R0066` and `R0067` are not placed in the tick (`AI-371`).
- The order word `ord+0x50` is read only by the group order 4 arm in the `+0x158` census. If reached, a
  completed pickup leaves 1, which sends that arm to the state 0xc acquire routine instead of a walk
  (`AI-372`).
- In `R0043` and `R0055` the route refusal flag is raised only by a full search that leaves
  the path list empty or by a near search aimed at the final destination; `R0178` also stores it. An empty waypoint search retries on the next pass. Whether
  another unit's cell blocks the contact ring is Unknown (`AI-373`).
- The full search's own body leaves the static list empty in three cases: start equals goal, picker A
  answers 0, or picker A answers the start cell; the route extractor's 1000-step cap is a fourth exit not excluded. Picker A scans rings up to Chebyshev radius `(D>>2)+3` around
  the goal, `D` being the start-goal distance. The search labels the start cell 0, and that cell is inside the
  scan only for `D` of 1 to 4. For rings inside the map, an unlabelled goal therefore leaves the list empty when no cell within
  `min(D-1, (D>>2)+3)` of it is labelled; near the border probes alias (`AI-394`).
- The idle-turn arm that a refused AI-owned pursuit can fall into, and the turn routines it calls, store neither
  `ord+0x08` nor `ord+0x0c` and call neither reacquisition nor route search; the arm ends order 0xb by no write
  of its own (`AI-395`).
- In state 3 a victim whose action word is 0x10 stands the member down to state 0xc before dispatch; the
  state-3 prologue reads no list. Arm 3 was not read for list tests, and other callers of the state machine
  were not enumerated (`AI-374`).
- Move stores `actor+0x50` = 1 and pickup 2; the executor tail reads that value when it consumes the
  refusal flag. A pending walk installed when recovery reaches zero at T0 first calls the walk routine at T0+2; that call writes a step only when the facing already equals the direction to the first node, otherwise it turns and the step follows at the next call (`AI-375`, `MOVE-090`).

### Stop, pickup, group evaluation, pursuit thresholds and cadence

- The stop routines clear the action word and neither the strike phase nor the progress byte; the strike
  continues unless progress is cleared elsewhere (Medium; `AI-381`).
- The first pickup pass stores action 2 and the same slot call runs the sack transfer; the second pass
  carries only the completion stores (`AI-382`).
- Group orders 3 and 5 branch per member on `ord+0x20` and `Player+0x28` without testing the pending order or
  destination: both zero reaches the evaluation (pending order 0, or 8 on the heal path, `AI-349`), and
  `Player+0x28` nonzero stores pending order 0xb (`AI-383`). Two routines on those branches are unread.
- In the walk, stepper, near-search and search routines no counter gives up: `mover+9` and `mover+0x78` only
  select a re-search against the `[Path Finding]` values, and the flag `mover+0x98` is set at three empty-list
  sites (`AI-384`).
- The AI pass that contains the group tail runs once per 16 ticks on both of its callers (`AI-385`).

### Transit into a held cell

A mid-transit walk call reaches the step with no destination check. A crossing into a cell a live unit holds releases
blind at the old footprint, rewrites the position into the held cell and has the occupy refused with no slot store;
the step does not test the refusal or roll back (`AI-387`).

## Retained targets at death and departure

Selected attack/follow policy prechecks test target action `0x10`, not HP or
off-map bit8. A match stands the actor down on its next eligible evaluation,
without directly zeroing its queued target or current strike pointer. Group
cadence and callback notification remain bounded. — AI-469

The normal actor manager clears source attribution on referred owner loss;
it has no direct sweep of other actors' order targets. Acquisition can select
a corpse alternative. Full notification and pointer lifetime remain Unknown.
— AI-470

A selected new attack changes Orders victim while retaining current target,
phase and countdown in its direct stores. Expiry uses current actor `+5c`.
Executor acquisition can relatch that pointer without clearing phase/countdown;
occurrence during a concrete replacement command remains Unknown. — AI-471

Departure removes the actor from the measured normal walker and retains its
pending fields locally. Global effects before normal actors is a composed
inference at Medium; the two procedure-local orders are High. Each actor's
attached effects locally precede its HP/order processing. Full-frame,
pixel-paint and remaining caller paths stay Unknown. Prepared spell source
cleanup uses owner loss rather than HP death. — MOVE-136, AI-474 (amended), MAGIC-297

The selected actor-cast policy queues a nonnull target without local HP,
terminal-action or off-map rejection; null target uses acquisition instead.
Replacement cast input, caller eligibility and pointer lifetime remain Unknown.
— MAGIC-298

## Temporary caster actions

The composed temporary-caster creation/evaluation interpretation is Medium. Direct
action stores `+0x54=0xd`/`0xe`, local calls, explicit separate-list append
and Unit vtable binding remain High. The positive-HP executor bypass is
High conditional on `actor+0x3c` still zero at its getter. Earlier zero
stores do not close unread same-actor calls `R1624`/`R0184` or builder
helpers. Action tables `L01775`/`L01776` are not decoded as semantic
values. The builder's own instructions assign no fallback Group order;
indirect assignment or registration remains open. Local ordinary activity
dereferences the Player, and the executor with override zero dereferences
the Group before its progress gate; these instructions remain High (`AI-478`).

Conversion binding is Medium because parameter table `L12450` is
undecoded; resolver miss, publication predicate and transfer/removal
instructions remain High. Builder helpers `R0289`, `R0462`,
`R0463`, `R0915` and `R0464` remain unread. Full scheduling,
indirect registration, null-owner spell outcomes and persistence remain
Unknown (`AI-477`, `AI-478`).

## A new order during a cast

Player order setters write the command state `actor+0x50` and the pending kind
`ord+0x08` at once. Only the move setter (opcodes 0x16, 0x1a, 0x1c) stores a
nonzero pending kind, 1; the other admitted setters store pending 0 and leave
the kind to the group pass. No setter read stores progress `ord+0x09`, phase
`actor+0x58`, countdown `actor+0x6c`, completion `actor+0x136` or the active
target and Spell (`AI-487`). The pre-arm helper acts only on a scroll pointer
(`AI-490`).

| progress `ord+0x09` | executor arm | action word set before the body |
|---|---|---|
| 0 | pending switch on `ord+0x08` | 0 unless 2 or 0xf |
| 1 | strike counter | 3 |
| 2 | cast counter | 0xd, or 0xe when `ord+0x5c` is 0 |
| 3 | step to the cell centre | 1 |
| 4 | status hold | 0x1a |
| 0xff | park, progress kept | 0x1a |
| 5..0xfe | progress cleared | 0x1a |

The progress-3 arm clears progress when the mover test is nonzero after the
step (the mover at a cell centre, `AI-350`), and the progress-4 arm when
actor bit 0x100000 is clear; the 0xff arm keeps progress (`AI-499`, whose clause that the step
stores `mover+0x98` is partially retracted). These are the arm exits (High). Every arm ends in the executor tail, which for a
nonzero `mover+0x98` can store progress 0xff (command state 0x17) or call
`R0004`, whose callee stores progress 1 (`AI-511`). The arms of progress 3, 4
and 0xff call no routine read that sets the flag, and every pass clears it, so the
tail's branch is not taken in such a pass through any path read; the callees
not read rest on the census and `MAGIC-ACTGATE-079` at their grade, and the entry
value of the flag is Medium (`AI-510`, `AI-516`). The step bodies read hold no
order-record operand (High); that a move, attack, cast or stop stored during a
step leaves it running to the cell centre, to be read at the first pass entered
at progress 0, is Medium: the setters of 0x14, 0x18, 0x19, 0x1e and 0x1f are
unread, the setters read store mover bytes, and the tail depends on the flag
(`AI-509`). Bit 0x100000 is set by the effect attach and cleared at expiry
(`L05766`, the only clear read); the attach also clears the bits of ids 0x17
and 0x1b when those two effects replace each other (`AI-512`). While the bit is
set no stored order is read (`AI-513`, Medium across ticks).
Progress 1 and 2 clear only when `ord+0x15 > 2` and
`actor+0x136` is set (`AI-488`). A player's cast installs once: the install pass stores progress 2,
then command state 0xc and pending 0, and the action body idles that tick
(`AI-489`). The next pass runs phase 0, where the mana is subtracted before
phase 5 and the wind-up countdown (`MAGIC-310`).

An order arriving while the cast is queued or approaching replaces it with no
mana spent, because the subtraction runs only at phase 0 (`MAGIC-310`). No
setter stores progress or phase, so an order arriving from the install pass to
the exit leaves the cast running; that it stays charged and unrefunded, and
that a repeated cast is charged again at its own admission, are Medium: the
census of `actor+0x9a` finds one subtracting store and no refund store among
the sites read, one site is unread, the application stores health 1 and mana 0
for one arm (`MAGIC-316`, `MAGIC-318`), the gates of the restore sequences that call the cut routine are
unread (`AI-503`), and the executor tail of `AI-354` can rewrite the active
target in some states (`MAGIC-311`). A move
acts at the first executor pass entered at progress 0; other orders act after
the group pass also writes their pending kind (Medium, `AI-491`). The scroll
pre-arm helper tests the action word and completion, not the cast, so the
install-to-admission gap and the exit tick are unprotected (`AI-490`). A
player's cast refused for mana stays at progress 2 with pending 0 and is
dropped or retried by the completion byte (`MAGIC-312`). The setter of
opcode 0x1b stores pending 0, as no direct callee read names `ebx` (`AI-514`),
and the setter of 0x1d reaches a member routine that stores command state 0xa and
pending 0 on every path (`AI-515`); the group-pass stores read are pending 4 and 1
for command state 8, 6 and 0xb for state 0xc and 1 for state 0xa, and the helpers
can leave other kinds, including 5 to 9 (Medium, `AI-514`, `AI-515`; these amend
`AI-501`). The group pass
writes pending 1 or 7 for command state 2 and 0xf for state 0xf (`AI-502`). A scroll cast pays no mana
(`MAGIC-317`).

## A unit acting without a new order

The routine that can start an attack or a cast for a unit without a new order
is chosen by its group's order byte (`AI-517`): order 0 runs the unit's own
state arm, orders 1, 2, 3, 4, 5 and 0x11 run a group arm, once per AI pass (`AI-526`). The table and the
arms are read (High); that no other path acts for a unit is Medium, because 7 of the 37 call sites of the three
acting routines are unplaced (`AI-517`). A player's own groups start at order 3, so the default idle unit runs the
order 3 arm (scorer `R0153`, veto `R0252`) and engages what is already in reach; the standing acquisition
below applies after an order that leaves group order 0 or 4 (`AI-517`, `AI-518`). Standing
acquisition builds its candidate set from the sight stamp, caps the pick at
reach, and stores pending 6, the target and the reach; it stores no progress,
phase, action word or completion, and the strike or cast starts at the next
executor pass at progress 0 (`AI-518`). A human owner's unit with no pick runs
the heal arm (`MAGIC-320`); a unit left in state 0xb engages what stands near its
post and returns to it otherwise; Guard stores 0xb with group order 1, so only the
patrol or a group of order 0 reaches that evaluator (`AI-525`, Medium). The engage
rolls the unit's spell slots each time it runs, and a hit stores a cast on the order
record (`MAGIC-322`; Medium that a human owner's unit casts this way in play); an armed cast cursor stores nothing on the unit (`MAGIC-321`).

The group pass of order 4 (Move) stores the move again for a member that has not
arrived and runs no acquisition before arrival or before the route ends on the
member's own cell; after arrival every pass runs acquisition and none re-issues
the move (`AI-519`). There is no group-pass attack-move; the executor tail path
is open (Medium): a member stood down to state 0xc is reacquired by the tail on a
later raise of the refusal word (`AI-524`). Orders 1 and 3 engage a scored victim
with pending 5; order 1 walks a member away from its post home, order 3 stores no
walk (`AI-520`). Order 5 runs the order 3 split when the candidate build found
something and the order 4 arm when it found nothing (`AI-520`, `AI-519`). Order 0
maps the command states to engage, acquire, re-issue and cast arms, and a
victim at action 0x10 stands a state 3 member down to state 0xc and acquisition
in the same call (`AI-521`). The executor tail reacquires for every state except
1, 0xa and 0x17 when the mover raises its refusal word (`AI-524`, Medium). A
Move stores `ord+0x50` = 0 for every member (`AI-519`).

A strike starts at progress 0 behind a gate that needs the facing byte equal to
the direction to the target and the edge distance within reach; the action body
then winds up, releases on a reach test with no facing or mover access, and
recovers, while progress 1 and 2 hold the unit (`AI-522`). The hit hook sets the
turn flag, the attacker's cell and the diplomacy flip and stores no pending kind
or target; the hit routine body, two of the hook's callers and the readers of the
attacker pointer are unread (`AI-523`, Medium).
