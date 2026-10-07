# Player input and commands

[Reference](format.md)

## Structure-use command

AI-STRUCTUSE-306 connects player order 0x24 to actor state 15, order 15 and action15.
The command stores a Building reference and approaches a width-dependent cell
with range 1,2 or 3. Local admission checks an orientation byte and Chebyshev
distance before selecting use; failure continues movement. Exact scheduler
latency, interruption and first dispatch after LOAD remain Unknown.
UNIT-STRUCTUSE-090 and UNIT-STRUCTUSE-091 describe the reached switch and
fountain effects. They resolve nearby objects again; retained order identity
alone does not establish which object receives the effect in ambiguous layouts.

## Player input contract (`AI-INPUT-121`, `AI-SELECT-122`, `AI-PANEL-123`, `AI-MINIMAP-124`, `AI-KEY-125`, `AI-CURSOR-126`, `AI-INPUT-127`)

Input is edge-driven. The base dispatcher maps physical Windows messages to each surface's
vtable, gives a capture object first refusal, otherwise stops at the first child containing the
point, and gives a focused child first refusal on keyboard input. The map therefore never receives
a point inside the right column. There is no mission `WM_MOUSEWHEEL`, `WM_CANCELMODE` or
`WM_CAPTURECHANGED` handler (AI-INPUT-121/127, SESS-INPUT-037).
SESS-INPUT-037's inventory-parent clause is partially retracted: the measured
opening route attaches the grid to the map view (MENU-088). The dispatcher
and physical-handler clauses stand. MENU-088's explicit Backspace selection-reset
clause is separately partially retracted by MENU-091; the measured edit arm
changes text/caret without selection-endpoint stores.

### Map mouse

| input | state / target | result |
|---|---|---|
| left down or double-click | marquee global zero | call OS `ClipCursor` and begin marquee; no engine capture or action yet |
| repeated left down or double-click | marquee global nonzero | consumed no-op; preserve the first origin and clip |
| left up | click-sized, no selected actor or select cursor | selection |
| left up | rectangle selection | select/toggle owned non-structures with strictly more than half overlap; equality fails |
| left up | attack cursor, `CUnit` target | order `0x19` |
| left up | attack cursor, other target or ground | move `0x16` |
| left up | move / swarm | `0x16` / `0x1a` |
| left up | defend, target / ground | `0x1b` / no order |
| left up | cast, unit / ground | spell `0x1e` / `0x1f`, item `0x25` / `0x26` |
| left up | pickup / town / patrol | `0x21` / `0x24` / `0x1d` |
| right down | map | set engine capture object and store pan origin |
| right move | non-zero signed `(delta pixels)/8` | pan by cells, mark drag |
| right up | drag / click | release only / send cancel `0x405` |
| right double-click | any | no-op |

The selection/order rows require an active marquee at `[L00210]`. Left up with that global
zero skips the `ClipCursor`, hit-update and selection/cursor-order prefix. It is a no-op when no
item is selected, but it still reaches the subsequent `session+0x3cc` item placement arm and can
emit `0x23` when `item+6 == 0xffff` or transfer record `0x22/0x32` otherwise. This edge ordering matters for repeated or unmatched button
messages (AI-INPUT-121).

The click/drag threshold is `screenW*10/640`, hence 10/12/16 pixels. The outer test sends an
axis strictly greater than the threshold to selection; the selection routine itself calls a
rectangle a click only when both axes are strictly smaller. This makes equality context-dependent
and must not be simplified to one `>=` test (AI-INPUT-121).

Selection: plain click replaces with the topmost intersecting non-structure regardless of owner,
or a structure whose class `+0x64` is zero. Shift-click and Shift-rectangle toggle an owned
qualifier only when the old selection summary at `view+0x144` has bit `0x04` clear. If a plain
click first selected a foreign actor, that bit is set and either Shift gesture over an owned actor
is a no-op; ground and foreign candidates also preserve the old selection. A plain rectangle
replaces only if at least one owned non-structure has strictly more than half overlap; equality
does not qualify. Alt after a single map selection expands to that object's group but does not
centre (AI-SELECT-122).

### Minimap mouse

Minimap actions occur on **left down**. Default sends camera `0x406`; move `0x16`; attack
`0x19` with any object id and `0x1a` without one; defend `0x1b` with an id and nothing without;
cast emits nothing; patrol `0x1d`. Left-drag repeats on every move. Right down and right-drag send
camera `0x406`. If both button bits are present on one move, the left test runs first and returns,
so only left semantics occur. No capture is taken. Left up only clears a special cursor; left
double-click, right up and right double-click are no-ops (AI-MINIMAP-124).

### Command cells and keys

| action | panel/key result | map click result |
|---|---|---|
| Attack / A | arm 1 | target `0x19`, otherwise move `0x16` |
| Move / M | arm 2 | `0x16` |
| Guard / G | immediate `0x17` | — |
| Defend / D | arm 4 | target `0x1b`, ground no-op |
| Cast / C | arm 5 | `0x1e/0x1f` or `0x25/0x26` |
| Swarm / S | arm 6 | `0x1a` |
| Stand Ground / T | immediate `0x18` | — |
| Retreat / R | immediate `0x14` | — |

Every consuming map/minimap action clears the armed mode. Guard, Stand Ground and Retreat repeat
on Windows key repeat; armed modes are idempotently reasserted. Although internal mode 8, the
patrol cursor and opcode `0x1d` exist, no shipped panel or key route leaves Patrol armed: R is the
immediate Retreat command. Exposing Patrol is an authored extension (AI-PANEL-123,
AI-CURSOR-126).

### Explicit Retreat execution

Ordinary-command Group construction has a narrower cleanup rule than a full sweep:
`R0196` deletes at most the first old Group rejected by `R0197`, then allocates.
The selected prologue does not directly mint Group+1c or populate the embedded
Group+20 word list. Constructor/list state and later behavioral callbacks are distinct
sources; Guard's GroupAI+20 order byte is not Group+20. — SAV-GRPCMD-578

Retreat is immediate command admission, not guaranteed immediate interruption.
The panel requires an active enabled cell, no selected inventory item and an eligible
inventory-bearing selection. Empty selection, a foreign first owner, or the structure
summary flag disables the panel. The producer includes entries with nonnull `actor+0x7c`
and caps the command list at 253. The dispatcher requires its AI manager and a resolvable
first actor; missing later actors are skipped. Its resolver rejects completed-death action
`0x10`, not every nonpositive-HP death phase (`AI-RETREAT-270`).

The command sets group order 0, actor state `0x16`, pending order 0 and current action 0.
It clears order fields `+0x38`, `+0x50`, `+0x60` and mover field `+0x7c`; it aligns desired
facing to current facing and releases a noncurrent reserved cell only at a cell centre.
It does not clear order-progress `+9` or all old target/spell slots. Conditional spell-object
cleanup is a separate dispatcher helper (`AI-RETREAT-271`).

Nonzero progress takes precedence over the next pending order. Progress 1 restores attack;
progress 2 restores the existing cast action. Both await counter `ord+0x15 > 2` and the
actor's animation flag `+0x136`. Progress 3 completes the current movement step to its cell
centre. Progress 4 remains held while its status bit is set. The setter's action reset is
therefore not proof that the old action was cancelled (`AI-RETREAT-272`).

Group order 0 dispatches state `0x16` without testing progress or retained
phase. Policy dispatch and pending execution are separate boundaries. An
executor invocation that clears nonzero progress jumps to its common tail
and does not execute the pending order on that invocation. A subsequent
invocation entered at zero can execute it even with a retained phase, provided
the activity gate admits it. With controller `+0xb388` and group AI `+0x45`
both zero, zero-entry progress instead returns action `0x1a`; a nonzero
completion-clear invocation can be followed by this new refusal. The selected
positive-HP actor caller also requires actor `+0x3c` nonzero. A conditional
recovery-zero sequence executes pending move at T0+2; native cadence and
command interleaving remain Unknown (`AI-364`).

An active route-failure tail is a same-invocation exception. State `0x16` can
reacquire and reinstall strike progress after clearing the old progress,
without resetting retained action phase. Native occurrence of this prerequisite
during a running cycle remains Unknown (`AI-365`).

While group order is 0, state `0x16` recomputes a flee order through the first withdrawal
helper below. Its radius is mover field `+8`, not a player-supplied destination. A zero
positive-HP count falls back to ordinary acquisition, which can choose a pursuit, idle or
autoheal order. The local state arm and helper have no timer, arrival termination or state
`0x16` reset; route failure also uses ordinary acquisition. Sustained whole-session
persistence and visible completion timing remain Medium/Unknown (`AI-RETREAT-273`).

Later admitted Move installs group order 4. Target attack replaces actor state with 3 or
fallback `0xc`; the two cast commands replace it with `0xd`/`0xe` or fallback `0xc`.
Nonpositive HP is handled before order execution in the unit tick and clears actor state
and action. Retreat does not bypass that death branch (`AI-RETREAT-275`).

The character panel is a second physical producer of Cast. Its `WM_LBUTTONUP` handler sends
selected item kinds 1, 2 and 5..8 to vslot `+0x7c`. The inventory grid's
`WM_LBUTTONDBLCLK` handler also forwards a resolved regular item to this same vslot. For an admitted owned single selection whose
item flags contain both `0x10` and `0x01`, `R0240` stores the inventory slot and effect
`0x29` context, then calls `R0097(0x0a)`; the routine remaps that value to mode 5. The next
normal map down/up therefore uses the item branch and emits `0x25` for a unit or `0x26` for
ground. The same character-panel vmethod contains the sibling transfer producers `0x22/0x32`.
This path is not a ninth command-panel cell and must not be lost by enumerating only the eight-cell
object.

The inventory grid also produces a transfer on physical `WM_LBUTTONUP`. Live content at
`grid+0x84` is hit-tested through vslot `+0x88`. With no selected item at `session+0x3cc`, the
handler stops after that hit-test. With a selected item, the signed hit result is passed unchanged
to vslot `+0xa4 = R0242`. This grid's destination vslot returns code 2. Its insertion
routine merges an equal stack first, inserts at a hit in `0..count-1`, and otherwise inserts before
a trailing gold sentinel or appends. A margin or outside hit of `-1` therefore remains an item
transfer. `L00211` emits `0x22/0x32`, then clears the selected item and posts `0x46d`.

The grid's other physical mouse slots do not hide another order. Left-down scrolls the inventory
viewport by one through one of two edge strips. Left-button move over a live cell, with no item
already selected, selects a regular item or gold amount and builds its drag cursor; Shift changes
the requested amount from one regular item or 1000 gold to the source stack's quantity. That move
emits no order. Right down, right up and right double-click are consumed no-ops. Only the later left-up reaches the transfer builder.

The same inventory-grid double-click has a separate gold-cell arm. A cell whose `item+6` word is
`0xffff` opens the prebuilt 296x168 Drop Gold modal at `(100, H minus 200)` only while the local purse is
positive. Enter or an inside left-down/up on `Ok` posts `0x445`; Esc or the cancel button posts
`0x446`. The action parses the edit, uses the amount returned by the owner resolver, subtracts it from the purse and emits
opcode `0x23` with that amount and the current map cell. Cancel closes without an order. The direct
caller set of `R0112` is exactly the map selected-item arm and this dialog
(`AI-PANEL-123`, `SESS-INPUT-037`). SESS-INPUT-037's inventory-parent clause
is partially retracted; these money/action clauses stand.

Other default mission keys:

| key | result |
|---|---|
| Esc | mission/town menu, unless focused child handles it |
| F1 / F2 | help / save |
| F3 | load in phase 2, Diplomacy otherwise |
| F4 / F9 | explicit no-op |
| F5..F8 | Ctrl assigns quick-spell slots 0..3; plain key selects the binding, with gated Cast-mode arming when the book is closed |
| F12 | toggle an unidentified global |
| Pause | modal pause text |
| numpad `+/-` | speed index, clamped 0..8, each step posting a speed line in a campaign session (`MENU-058`); Ctrl+plus selects the unpaced/max-speed idle loop, Ctrl+minus restores paced mode and resets its epoch |
| digits/numpad digits | select group; Shift augments; Ctrl assigns; Alt selects and centres; Ctrl wins |
| arrows | pan one cell per repeat |
| E | select all owned exact-name `CUnit` objects |
| Space | toggle/close inventory and spellbook together |
| B or Q / I or backtick | spellbook / inventory |
| Enter | open and focus text entry; subsequent map shortcuts are suppressed |
| Tab with character panel focused | send panel custom message `0x412`; precise effect Unknown |
| Ctrl+W/F/H/U/L/N/O | retreat mode, formation, show health, autoheal, flying damage, day/night, smoothing; each posts its new state's line (`MENU-057`, `MENU-059`) |
| Alt+S | screenshot |

Modifier latches set on down, clear on up and all clear on focus loss. Keydown does not inspect the
repeat flag. Backspace clears three unidentified global containers; Alt+B..Y except S emits an
unidentified outbound record `0x46`. Tab's `0x412` route exists only when focus dispatch reaches
the character panel. Unlisted keys have no mission action (AI-KEY-125).

### Quick-slot state and lifecycle

The four signed quick-slot values live on the shared spellbook controller at
`campaign+0xec`, offsets `+0x64/+0x68/+0x6c/+0x70`. Construction writes `-1`
to all four. Current spell `+0x60` and selected-item override `+0x74` are
separate fields. The reached keys do not index actor/group-owned bindings.
Message `0x411` and controller right-down/up clear current spell only;
selection-summary recomputation has no direct binding write. Book open/close
attaches/detaches the same controller. Selection carry beyond these bodies is
Medium; no-load mission entry and new-campaign reset remain Unknown.
— AI-QUICKOWNER-280

F5–F8 send `0x417`, slot 0–3 and the Ctrl latch to that controller. Ctrl copies
current spell when it is not `-1`; otherwise it uses a nonnegative spell-grid
mouse hit. Each assignment empties all other equal bindings to `-1`, so a
duplicate moves rather than swaps. The hover arm neither checks current spell
availability nor changes the current-spell field. Ctrl still reaches the
key handler's subsequent arming checks. — AI-QUICKASSIGN-278

A plain key copies a nonnegative binding into current spell. If the book is
closed, an available stored-index bit in `view+0x148` requests action 9,
which maps to mode 5 only when Cast capability is present. It does not open
the book or directly emit a cast order. An empty `-1` slot leaves local mode
and current spell unchanged; an unavailable nonempty slot changes current
spell but requests no new mode. Neither local branch cancels an existing
mode. An open book skips mode arming. Generic child dispatch and malformed
saved indices remain outside this bounded result. — AI-QUICKINVOKE-279

The original save programme writes the four indices in F5–F8 order to
`SpellBook/Shortcuts` and restores them to the same controller. It separately
saves/restores `SpellBook/Pressed`. The reached slots are persistent save
state. Native save/reload behavior with populated slots remains Unknown. — AI-QUICKSAVE-281

### Book identity and selection predicates

For valid book cell `i=0..23`, current spell and quick bindings hold `i`.
Map targeting passes `i+1`; opcodes 1e/1f carry this ordinal as a byte at
command+10. The dispatcher rewrites it through the byte at
`L00198+4*(i+1)` before actor-book lookup. That is the same storage as the
statistics table at `L00199+4*i`. Cell 5 means command 6 and intrinsic spell 23,
not spell 6. Lookup uses actor+140 and returns the sparse ID-indexed element;
its pointer is parked at actor+44. Malformed indices and stale actor spell
pointers remain outside this valid-cell contract. — AI-SPELLIDENT-286

Selection recomputation ORs object+18 into view+148 for every selected object
with nonnull+7c, before class/type/session predicates. Count and the primary
object use this same accepted population. The two book producers independently
filter those selected objects by the chosen cell bit; the append helper caps
the command at 253 members. Snapshot freshness and all upstream selection
restrictions remain Unknown. — AI-SPELLPOP-287

Cast capability is narrower than availability. In session states+3dc==1 or
with bit 2 set, each accepted `CUnit` with+20 equal 17h/18h and nonzero+18 can
OR200h into view+144. This is the current loop object, not only the primary
object; the primary-only clause of AI-PANEL-061 (partially retracted) is corrected. Other session
states assign summary8. The owner comparison still uses the primary. The
mode helper returns 0 for summary bit 4, otherwiseEFh plus 10h for 200h; the
panel separately disables on zero count or summary24h. Item action0Ah bypasses
the mode-bit refusal. Static predicates do not prove every synthetic mixed
population is reachable. — AI-SPELLCAP-288

Mouse cell selection requires the cell's availability bit. A nonnegative
shortcut can become current before its later arming guard, so unavailable
does not mean impossible as current. The reached map Cast branch takes
nonnegative item+74 before current+60 and rejects a negative result. Its
book target flag is indexed by cell at L00200:14 of 24 entries are nonzero.
A unit-like hit plus nonzero flag routes to actor production, other hits to
cell production; without a hit a nonzero flag emits no order.
— AI-SPELLGUARD-289

Item opcodes 25h/26h instead carry the word from controller+78 at command+10.
The consumer retrieves that inventory item and constructs a Spell from its
first kind 29h effect's byte+40. It skips the book ordinal translation and
the producer's learned-spell bit filter. Its target predicate uses L00201.
Item snapshot synchronization and the complete use/refund lifecycle are
separate boundaries. — AI-SPELLITEM-290
