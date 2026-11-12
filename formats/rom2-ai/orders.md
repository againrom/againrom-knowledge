# Planner and executor

## Order storage and planning

| Input/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Player setter | The 20 read setters, including both order-1 setters and order 3, match individually in EN/RU. Reset and stop chains are read. Direct stores replace order/parameters and prepared byte; no direct primary write occurs in the setter bodies. | R2-ENGINE-541 High RU bodies, Medium completeness; R2-ENGINE-1105 and R2-ENGINE-1107 High local bodies, Medium complete input/callee effects | Other input routes, indirect writes and eventual execution; read their callers and run the next action |
| Unit-order selector | RU L2.01725 uses the 28-entry table L2.01729. Relevant routes: 1 walk, 3 attack, 0xb guard, 0xc idle, order 0xa (name not established), 0x16 flee, 0x18 planner L2.02864. | R2-ENGINE-555 High table, Medium per-order closure | Branch choices inside other planners; planner cadence |
| Attack target has +0x4c bit 8 | Skip slot draws; run L2.01733(unit,0), whose RU path refreshes/ranks candidates and chooses retarget or default. | R2-ENGINE-651 High local planner; R2-ENGINE-808 High RU stores, Medium composed target choice | Event meaning and unread candidate/mage helpers; not universal attack cancellation |
| Ordinary attack slots | Draw for each nonzero slot 0..2; the last draw below its signed threshold wins. Pick 23 casts only beyond 2 cells. Threshold is Probability * 0x147, draw bits 16..30 of LCG seed*0x343fd+0x269ec3. | R2-ENGINE-651 High draw/loop, Unknown caller rate; R2-ENGINE-652 High stores, Medium row binding | Initial seed and actual call frequency |
| Remaining attack route | With no target (RU body; in EN L2.03461 reads target +0x4c at L2.05580 before any null test, so a null target faults there) or a self target take first unit from L2.02117; prepare +8=5, +0xc=target and +0x14=range, with computer-mage helper first. | R2-ENGINE-505 High RU body, Medium helper meanings; R2-ENGINE-936 High EN null-target read; R2-ENGINE-1012 High exception, Medium composed chain: separate EN fault in the LOAD-driven session update, on the unit send path | Mage helper and candidate eligibility; global activity-5 writer census; RU null target |
| Struck by an area blow | The occupant routine calls the hit hook (caster, occupant, 1); the hook body stores no unit order, prepared +8, +0x0c or target +0x20 (its callees and the EN append are unread). The next prepared action comes from the planners and Group assignment; when the owner's diagonal has bit 0 clear and low bits set (mission 10: 0x12), the hostility filter keeps a same-owner caster out of the hostile list in the read bodies (Medium). | R2-ENGINE-985 High occupant call, Medium caster field set by the executor's cast; R2-ENGINE-987 High hook stores; R2-ENGINE-989 Medium composition | Posted-record readers; untraced +0x20 writers |
| Idle target selection | Rank eligible candidates by nearest edge distance, least turn and later full tie. A found target prepares activity 6; no target takes support/default. Activity-6 admission body is now read in EN/RU. | R2-ENGINE-501 High RU choice, Medium helper/composed no-step; R2-ENGINE-1106 High local admission, Medium full path | Mage/support effects and complete candidate population; read those callees and run a nearby-hostile control |

## Prepared actions used by these routes

Prepared dispatch occurs only after entry admission, no hold reinstall, primary
+9 still zero and valid nonzero +8. Clearing +9 in a primary arm does not
dispatch +8 inside that arm. R2-ENGINE-556 High local tests, R2-ENGINE-560 Medium
end-to-end composition; planner cadence remains Unknown under R2-ENGINE-558.

| Prepared +8 | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| 0 | Bypass prepared dispatch; most setters need a planner to prepare a later action. | R2-ENGINE-560 Medium composition; R2-ENGINE-541 High setter stores | Whether/when planner runs |
| 1, walk | Fifteen-cell prepared table RU L2.01571 / EN L2.02887 takes L2.05251, calls L2.02634(goal +0xa,0), then walk L2.02867. | R2-ENGINE-883 High table/bodies | Grid and final route outcome |
| 2, direct strike | Call strike initializer L2.02407: primary +9=1, phase +0x15=0, unit state 3 and unit +0x5c=command +0xc. | R2-ENGINE-556 High local initializer | Range helper and common tail interactions |
| 5, approach/attack | RU L2.02437 / EN L2.06208 recomputes range with L2.01728 and tests L2.01574 against command target +0xc. Admitted: strike/state 3; otherwise call approach wrapper/state 1. | R2-ENGINE-1106 High local comparisons and stores, Medium helper/path composition | Virtual sizes, search/step give-up and execution timing; read those helpers or run the trajectory |
| 6, idle attack | Strike when L2.01574 admits: facing [+0x1c0]+0 equal to the 8-way heading to the target and the target within range. Otherwise turn toward it through L2.02409, which writes only facing, heading and turn counters; no step call. | R2-ENGINE-501 High RU local choice; R2-ENGINE-926 High EN/RU bodies; R2-ENGINE-927 High turn body, Medium snap threshold (L2.01523 unread) and one call per step | Executor cadence; RotationSpeed values |
| 8/9, target/point cast | A self-targeted 8 initializes at once. Otherwise the helper (L2.01574, L2.02567) admits only with the facing equal to the heading to the target or cell (the caster's own cell gives 224) and within range; admitted: primary 2 and state 13/14; else approach: step when not centred, turn when centred in range (for 9 only when range is not 0), route otherwise. Turn: within 32 of 256 at once (Medium: L2.01523 unread), else RotationSpeed per call. Wind-up and recovery precede completion. | R2-ENGINE-543 High local flow, Medium apply meanings; R2-ENGINE-926 High EN/RU bodies, Medium wait cadence; R2-ENGINE-928 High bodies, Medium own-cell run; R2-ENGINE-927 High turn, Medium per step and snap threshold | Apply branches and transport; no general successful-cast guarantee; how the facing reaches the client |
| 0xb, look-around/support | Guard-at-post path writes this prepared byte for owner +0x2c nonzero. Look-around turns/new heading consumes hit flag. | R2-ENGINE-844 High write, Medium helper mapping; R2-ENGINE-504 High RU body, Medium EN/helper meaning | Support/mage callees; this does not write unit order 0xb |
| 3/13/14 | The executor's prepared table does nothing for these values. | R2-ENGINE-542 High local dispatch | Their producers and wider meanings |

## Primary actions

| Primary +9/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| 0, admitted | Before prepared dispatch, a nonzero command word +0xb0 and centred position can invoke relocation and clear the word. It is a cell, not a queued order. | R2-ENGINE-557 High local consumption, Medium writer completeness; R2-ENGINE-705 High RU relocation stores | Cell kind 0x17; rejection and later effects |
| 1, strike | State 3; increment phase +0x15. Phase>2 and completion unit +0x136 nonzero clears +9 and state. | R2-ENGINE-553 High | Actor substate/common-tail effects with a new order |
| 2, cast | State 13 if unit target +0x5c nonzero, otherwise 14; phase 3 posts record; same completion test. Actor substates 0 -> wind-up 5 -> recovery 7 -> completion. | R2-ENGINE-553 High executor; R2-ENGINE-543 High local flow, Medium stage meaning | Full apply/record semantics |
| 3, step | Advance L2.02848, state 1; position offsets +4/+5 both 0x80 clears primary and state. | R2-ENGINE-553 High local arm; R2-ENGINE-546 High centre test, Medium uninterrupted motion | Movement helpers and flag/new-order coexistence |
| 4, hold | Executor installs it when effect mask 0x40000 is set with +9 zero. State 27; clear +9 when bit is gone. | R2-ENGINE-553 High | Effect meaning and complete interruption |
| 5..254 | Clear +9, state 27, then common tail. | R2-ENGINE-553 High all 255 table cells | Tail may change primary again |
| 255 | Keep +9, state 27, then common tail. Planner L2.02864 owns both read 255-to-0 clears. | R2-ENGINE-553 High table/body; R2-ENGINE-554 High local clears, Medium global coverage | Conditions selecting clears and other writers |
| Any arm except inactive early return | Run common empty-route tail; it can reset, re-pick or write primary 255. Local precedence does not prove that every newly stored order survives. | R2-ENGINE-872 High tail branches; R2-ENGINE-556 High local flow; R2-ENGINE-560 Medium retention | Planner timing, helper effects and flags during an arriving order |

## Health thresholds, difficulty and attack order

The stat reader stores them in command +0x40 and +0x44 for a unit whose template loads (Medium: loader condition, script-created creatures; R2-ENGINE-1001). The map arm scales a non-Human record by the difficulty word when Session +0x74 is 0 and the value is not 2 (R2-ENGINE-1002). In the bodies read, the only direct caller of the threshold post-step is the group tick (byte scan; indirect and jump routes unread, Medium; R2-ENGINE-1003). The order-3 arm passes command +0xc to the planner with no null test (R2-ENGINE-1004). In the bodies read, LOAD leaves the thresholds as saved (R2-ENGINE-1005).

## Composed route of a unit's next action

The rows are bounded to the RU bodies each claim names. A row is no stronger than the lowest cited clause. An order value not listed takes the row with its selector arm, R2-ENGINE-555. The Group pass runs when the session counter +4 mod 16 is 6, once per 16 values of that counter (R2-ENGINE-778, R2-ENGINE-1016). Seconds per step are Unknown and the actor step's dispatch is unread. Rows beginning Zero hold when the unit's Group reaches the selector, as above; which programme a player's Group holds is Unknown outside the mission-10 saves (R2-ENGINE-1040). The per-unit loop L2.01724 has no reference found, and both owner +0x2c branches of the wrapper reach the selector. R2-ENGINE-1017. The loop's negative also covers Jcc rel32 and rel8 forms (R2-ENGINE-1124, Medium census). Before its dispatch the selector calls stop chain L2.01576 for order 3 whose target has +0x54=0x10, orders 8/0x11 whose command +0x10 object has it, and order 1 with +0x90 set at the destination cell (R2-ENGINE-1121; High tests, Medium state meaning).

| Owner +0x2c / order / stage | Next action | Claims and clause grades | Gap and the read or run that closes it |
|---|---|---|---|
| Any owner, programme route member | The Group programme route decides as on the behavior page; programme 0 reaches the selector through the wrapper. | R2-ENGINE-507 High, R2-ENGINE-882 High, R2-ENGINE-1017 Medium, R2-ENGINE-1040 High body, Medium hero identity | Programme of a player's Group in other missions and at mission start: only mission-10 files are walked (R2-ENGINE-1040). L2.06060 is reached from input, from L2.04542 and from the script action switch (R2-ENGINE-1124, R2-ENGINE-1123); read the eight callers of L2.04542 and action table L2.06288, or run mission start |
| Zero / order 1 / walking | Planner L2.02401 prepares +8=1 and +0xa and reads no candidate list in its own body (R2-ENGINE-1042). Candidate list use later in the walk is not read. | R2-ENGINE-555 High store; R2-ENGINE-1042 High body | Whether anything reads a candidate list later in the walk; a run walking past an adjacent enemy |
| Zero / order 1 / movement +0x98 set at the tail | Stop chain stores unit order 0xc, state 0 and prepared +8=0. Empty route sets +0x98; the read own-cell arrival branch separately sets movement +0x90. | R2-ENGINE-1020 High local tail; R2-ENGINE-1108 High local stores, Medium route composition; R2-ENGINE-1121 High selector test: order 1 with +0x90 set and the destination equal to the own cell calls stop chain L2.01576 | Order 0xc from that chain: R2-ENGINE-1020 High, R2-ENGINE-1042 High (RU); the step continuations between the +0x90 store and the next selector pass are unread; run arrival with no enemy |
| Zero / order 0xc / standing | Idle planner L2.01731 reads the candidate lists; nearest then least turn; found goes to L2.01525: a player-owned mage first tries L2.01527, otherwise activity 6 with +0xc = target; none takes L2.01531 and L2.02911 (+8=0 or the helper route). Its own body stores no order and no +0x20. | R2-ENGINE-1018 High body, Medium helpers; R2-ENGINE-501 High RU choice (its Medium for the L2.01525 clause is amended by R2-ENGINE-1043, High whole body); R2-ENGINE-555 High for the store of 6 | L2.02404 gathers world units on marked sight cells and filters them (R2-ENGINE-1127 High body, Medium grid helpers); L2.01527 beyond RU L2.06314 and its EN continuation (R2-ENGINE-1127 Medium) |
| Zero / any order but 1, 0xa, 0x18 / movement +0x98 set | Re-pick chooses a within-range least-turn target, stores +0xc/+8=6, then activity 6 admits a strike or requests a turn. The owner/predicate/range-below-2 exception can clear +8. | R2-ENGINE-1019 High body, Medium predicate meaning; R2-ENGINE-1105/1106 High initializer and whole activity-6 body, Medium complete path | Predicate/kind bindings; turn wrapper L2.02408 only turns and sets state 1 (R2-ENGINE-1127 High); compare live next actions |
| Nonzero / no candidate at a re-pick | Stores +8=0xb, then L2.02430, which turns toward the stored heading or picks a new look heading and sets state 1 or 0, with no order or target store. | R2-ENGINE-1019 High body; R2-ENGINE-1127 High EN/RU body | Random draws L2.02431/L2.01532 per call; run a computer unit with no candidate |
| Any / hit | The hook body stores the struck unit's command hit fields and no +0x20, +8 or +0xc (R2-ENGINE-845, R2-ENGINE-987; callees unread); it also calls the relation helper. The +0xbbec records have four appenders and no read operand found in a RU sweep. | R2-ENGINE-987 High body; R2-ENGINE-1021 Medium; R2-ENGINE-1123 High read cases, Medium negative | Seven case routines, L2.00516 and L2.06061 read no +0xbbec operand; the other 20 case routines, L2.02490 and L2.06283 are unread; run striking a unit and watching +0x20 and +8 |
| Any / writers of the order fields | Linear-sweep census by traced base class; the prior untraced +0x20 writes stay untraced individually. | R2-ENGINE-1022 Medium | Reading the entry of each boundary-class function |
| Group programme | Constructor programme 0 and mission-10 saved 3/4 retain their original grades. RU input type 0x18 reaches L2.06060: member order 0xc, programme 3. Programme and member order are different fields. | R2-ENGINE-1040 High constructor/body, Medium hero identity/completeness; R2-ENGINE-1107 High RU cell/arm; R2-ENGINE-1113 High named paired setter stores | Script action switch L2.06286 calls eleven Group methods, seven with a programme store read in R2-ENGINE-1113 (R2-ENGINE-1123 High body, table unread); L2.04542 sets programme 0 then 3 from eight unread sites (R2-ENGINE-1124); mission-start state needs a run |
| Hero, programme 3 | Route L2.02466 per member: L2.01530 if +0x20 set, else L2.01531; no direct call to the selector; in the census of R2-ENGINE-1017 the selector is reached only if loop L2.01724 runs, an indirect or computed route not excluded. | R2-ENGINE-1041 High body, Unknown composition; R2-ENGINE-1017 Medium census; R2-ENGINE-1040 Medium hero identity; R2-ENGINE-1124 Medium six-form census, High caller body | L2.02466's only site is the Group tick switch L2.02456 (table L2.02462 unread); run counting hits of L2.01724 and L2.02942 with the hero walking in mission 10 |
| Walk prepare; order 0xc setter | Walk prepare stores +8=1/+0xa. Order-0xc setter stores the own cell and clears +8. The three empty-route flag set branches and the separate +0x90 arrival store are now locally read. | R2-ENGINE-1042 High prepare/setter; R2-ENGINE-1108 High local branches, Medium route composition | Occupied goal/all neighbours blocked and arrival-to-reset path; read search/occupancy/step continuations and run those controls |
| Standing, found target | L2.01525: owner +0x2c zero with L2.00044 true tries L2.01527; otherwise +8=6, +0xc=target, +0x14 from L2.01728. | R2-ENGINE-1043 High body, Medium helper meanings; R2-ENGINE-1127 Medium partial L2.01527 | L2.01527 beyond RU L2.06314: its read part stores command +0x60=1 before cast preparation or clears +8; kinds other than the mage branch |
| Attack order endings | The selector calls stop chain L2.01576 for order 3 when target +0x54=0x10; the chain ends in order 0xc and stores command +0x50=1. L2.01530 itself stores no order: with target +0x4c bit 8 set it runs idle planner L2.01731 through L2.01733(unit,0); with a null or self target it first takes the first unit of L2.02117 and runs the idle planner only when none is found; a slot draw prepares a cast with command +0x60=1; otherwise +8=5 or a mage helper. | R2-ENGINE-1121 High tests, agrees with R2-ENGINE-504 (High, RU); state 0x10 as a dead or removed target Medium (R2-ENGINE-518, R2-ENGINE-399, R2-ENGINE-364); R2-ENGINE-1020 High, R2-ENGINE-1042 High (RU); EN L2.02686 whole; R2-ENGINE-1120 High EN/RU bodies, Medium callee meanings; R2-ENGINE-1122 Medium byte-store bit-8 census, effects of the bit in R2-ENGINE-393, R2-ENGINE-477, R2-ENGINE-616; a call of the resolver on a target with the bit returns 0 before the roll (R2-ENGINE-1148 High, EN/RU, before the roll and before any draw; the generator arithmetic of the roll is equal in EN and RU, R2-ENGINE-1150) | Writer of target state 0x10 and meaning of bit 8 (read L2.06235 and the clears); a run striking until the target dies |
| Script switch and +0xbbec clear | Table L2.03277 maps field +4 to the 27 arms (cells 5..7 permuted). Seven read case routines store into result array +0xc784. L2.06061 calls L2.06283 per 0x54-byte element; L2.06062 calls L2.06284 on each node of a chain (read as a free, Medium, R2-ENGINE-1045). L2.00516 tests trigger conditions against the result array and runs actions through L2.02490. | R2-ENGINE-1045 High shape, Medium clear meaning; R2-ENGINE-1123 High read bodies, Medium negative; R2-ENGINE-1021 Medium | Case routines from L2.06275 on, L2.02490, L2.06283 and action table L2.06288 |
| Hero of a player, counted run (EN, one mission-10 launch) | Loop L2.04213 and wrapper L2.02403 had 0 hits while the hero stood, read order 1 and stood; EN L2.04984 was entered with the hero. Actor step L2.00263 ran once per scheduler counter value for each of 30 units. A ground click wrote order 1 and moved the hero to another Group object; the left record read programme 0xff; order read 0xc afterwards. | R2-ENGINE-1072 High counts, Medium EN correspondence and route; R2-ENGINE-1073 High counts and each launch's rate under a debugger (16.1 and 32.1 counter values per second, unexplained), Unknown outside a debugger; R2-ENGINE-1074 High values, Medium new Group object, save Group identity, that the unit is the hero, and +0x50, +0x20, +0x14 as order, programme and owner in EN (RU meanings) | A hostile in reach; other missions; RU; the writers of order 1 and programme 0xff or 4; goal-cell class of a blocked route |
| Hero of a player, static continuation (EN and RU) | Idle/programme routes retain their earlier scope. Scorer L2.05238 and assignment L2.05173 branches, Group-sight/filter bodies and named programme setters are read; no universal pursuit-ending or all-kind behavior follows. | R2-ENGINE-1075/1076 High local bodies, Medium callees; R2-ENGINE-1111/1112/1113 High named branches/stores, Medium complete choice/writer population | Remaining planner endings, virtual kinds and caller selection; read those bodies and run pursuit loss/next action |
| Hero of a player, second counted run (EN, save game0006) | With no attack order the hero read order 0xc, state 3 and another owner's unit at +0x5c about 6 s after the click at (400,190), and order 0xc afterwards. Loop L2.04213 and wrapper L2.02403 stayed at 0 hits; the idle planner was entered with the hero from L2.06123 and L2.06124 only. | R2-ENGINE-1078 High counts and values, Medium field meanings (+0x50, +0x20, +0x14 in EN are RU meanings), caller match, that the unit is the hero and that its owner is hostile, R2-ENGINE-556 High RU initializer | Distance to that unit: a run that reads a position field. What wrote the target and order 0xc: a write watch. A hostile near with no strike: a further counted run. |

## Static continuation by owner, kind and stage

EN/RU below means the named local bodies were read in both roots. It does not
promote inherited caller paths or unread virtual kind implementations. Any-owner
rows start at an already reached routine. Player/computer labels retain the
older owner-field binding. R2-ENGINE-1115 states the remaining correspondence gaps.

| Owner kind / unit kind / order and stage / roots | Next action in the read path | Claims and clause grades | Remaining gap and settling read or run |
|---|---|---|---|
| Any / read unit layout / order-0xc setter reached / EN/RU | Reset, store unit order 0xc, own cell in command +0, prepared +8=0. The direct immediate store occurs once per EXE in the reduced sweep population. | R2-ENGINE-1104 High local store, Medium census; R2-ENGINE-1107 High paired setter | Register/alias/copy writers and other images: classify the complete census and other admitted images |
| Any / read unit layout / strike initializer reached / EN/RU | Primary 1, phase 0, state 3, unit +0x5c=command target +0xc. | R2-ENGINE-1105 High local stores | Which live caller wrote a sampled target: read caller dispatch or run a write watch |
| Any / read unit layout / target-cast initializer / EN/RU | Primary 2, state 13, unit +0x5c=command +0x28; command +0x5c=1. Point cast clears command +0x5c, without proving a unit-target clear. | R2-ENGINE-1105 High local stores, Medium complete lifecycle | Unit-target clears on ordinary completion: read the remaining lifecycle writers |
| Any / read virtual-size operands / prepared 5 / EN/RU | Recompute range, require facing equal to heading and edge-distance result within range; admit strike or request approach/state 1. | R2-ENGINE-1106 High local tests/stores, Medium path; R2-ENGINE-1125 High approach wrapper: L2.01578 with command +0x14, step +9=3 when L2.00490 is zero, state 1; R2-ENGINE-1139 High arithmetic: the result is the Chebyshev distance from the attacker's cell to the nearest cell of the target's body, the approach goal is the target's packed word +2, the anchor cell for writer L2.02120 (Medium for other writers), the approach range is command +0x14 while admission recomputes it, Medium centre-only scheduling, no run; R2-ENGINE-1140 High: the Unit and Human size slot returns byte +0x49 (tokenSize 2 for typeIDs 66 and 68, 3 for 71) | Size slots of other classes, L2.05182 and L2.00490: read them; run positions at initiation/application |
| Any owner +0x2c / weapon with a Spell, mage predicate true / prepared 5 admitted, state 3, no earlier planner branch (target +0x4c bit 8, a winning slot draw, L2.00056 or L2.01527 casting) / RU, EN field-aware equal | By the code read, the actor arm stores the weapon's Spell at +0x64 and the weapon at +0x68, sets state 0xd and calls the cast routine in the same step; in the bodies read there is no mana test, no distance re-test and no sight test; application follows the wind-up. With owner +0x2c zero the selected-spell policy runs first, and L2.01728 returns that spell's +9 when +0x78 is nonzero and affordable. | R2-ENGINE-1144 Medium composition, no run; R2-ENGINE-1136 High read bodies, Medium unread callees; R2-ENGINE-1138 High cast routine, Medium apply tail; R2-ENGINE-1142 High read bodies, Medium +0x78 writer; R2-ENGINE-1120 High branches; R2-ENGINE-900 High L2.02338 | Release distance in a run: a Staff mage attacking an Ogre and a one-cell unit; L2.01527's tail, L2.01512, the 92 unread callees of R2-ENGINE-1136 and the apply tail: read them |
| Any / book cast at a unit / order 0xd, prepared 8 / EN/RU | L2.01524 stores prepared 8, target, spell and range = spell +9; the prepared-8 handler admits by facing and edge distance against spell +9, except for a self target. The eight RU scan hits of the packed-cell metric L2.01520 lie in L2.06342..L2.06343, after support entry L2.01512; they are not read as call sites and computed calls are outside the scan. | R2-ENGINE-1143 High bodies and scan hits, Medium site functions | Functions holding the metric's scan hits: read L2.01512 to its end |
| Group member or idle actor / unit larger than one cell / candidate lists / EN/RU | On the marked-cell path Group sight and the idle builder add a unit when the sight-grid byte at its anchor cell is set and test no other body cell; an empty live list is refilled from fallback +0xa88 with no cell test, and the command cell +0x58 is also marked. | R2-ENGINE-1137 High tests, Medium anchor corner and packed word; R2-ENGINE-1112 fallback shape; R2-ENGINE-502 inherited marking grades | Other grid readers and position-object writers: read them; run a large unit with only a non-anchor cell in sight |
| Player branch / kind passing the earlier mage predicate / order 0xc, idle finds target / RU inherited caller, EN/RU activity 6 | Mage helper can handle it first; otherwise prepare 6. Activity 6 has a range-below-2 exception, then the shared facing/range admission; rejection requests a turn. | R2-ENGINE-1043 High local producer, Medium kind/helper meaning; R2-ENGINE-1106 High consumer, Medium composition | Complete mage helper and per-kind outcome: read its full body and run matched kinds |
| Computer branch / assigned member / programme attack route / RU inherited route, EN/RU local consumer | The route can prepare attack; the shared admission starts strike. Programme is not the member order. | R2-ENGINE-882/1076 High routes, Medium helpers; R2-ENGINE-1106/1113 High local tests/stores, Medium composition | Every programme/kind route and pursuit ending: read the remaining planner callees and run target loss |
| Any / target predicate +0x2c nonzero / state-3 application / EN/RU local bodies | Reject null/owner-null or nonpositive-HP attacker; reject application distance beyond unit +0x12c. Admitted paths subtract target slot +0x54 result from HP. RU has an additional kind/item rejection that EN lacks. | R2-ENGINE-1106 High local stores/tests, Medium whole strike path; R2-ENGINE-1125 High RU slot +0x34 path (word +0x42 minus L2.01703's result), Medium amount | Virtual damage, L2.01703, L2.02340 and the EN counterpart of L2.02316 (not found): read them and run next action |
| Any / walking member / search returns empty route / EN/RU | Store own cell as movement destination and +0x98=1; the executor tail branches by order. Order 1 resets to 0xc/state 0; 0xa advances points; 0x18 sets primary 255; other orders re-pick. | R2-ENGINE-1108 High local branches, Medium composition | Why search is empty: read search/occupancy; run occupied goal and fully blocked neighbours |
| Any / walking member / own cell equals movement destination / EN/RU | The read walk branch stores movement +0x90=1. For order 1 the selector then calls stop chain L2.01576 when +0x76 equals the own cell. | R2-ENGINE-1108 High local store; R2-ENGINE-1121 High EN/RU selector test; R2-ENGINE-1020 High, R2-ENGINE-1042 High (RU), EN L2.02686 whole for the order-0xc store | Other +0x90 consumers and cadence: read the remaining +0x90 census sites and run unobstructed arrival |
| Different owners / struck unit / area pulse reaches hook / EN/RU hook, inherited area | Write hit memory and event record. Local hook absence of order/target stores does not exclude scheduled changes. | R2-ENGINE-985/986 High local bodies, Medium binding; R2-ENGINE-1109 High stores, Medium negative/composition | Posted-record/aliased readers: read the script cases, then run the next tick with order fields watched |
| Any / virtual slot +0x30 nonzero / Darkness Effect 14 applies or expires / EN slots, RU bodies masked-equal | Add or subtract magnitude<<8 at words +0xa4/+0xe4; +0xa5 is the high byte. Same-ID renewal keeps the less negative signed magnitude and independently maximizes remaining duration. Inherited sight reads +0xa5. No retained-target guarantee follows. | R2-ENGINE-1110 and R2-ENGINE-1166 High slot arithmetic; R2-ENGINE-1167 High comparison, Medium composed characteristic; R2-ENGINE-502 inherited reader grades | Kind predicate, +0xe4 readers and candidate consequence: read them and run a pulse on an attack holder |
| Group member / virtual kinds unresolved / scorer L2.05238 and assignment L2.05173 / EN/RU | Adjust distance for range, reject adjusted distance>1; score distance/turn/preference; choose strictly least score, first on tie, else clear assigned target. | R2-ENGINE-1111 High arithmetic/stores, Medium complete choice; R2-ENGINE-1126 High constructor bytes 2,1,2,4,4,2,1,0,2,1,4,0,2,1,2,4, Medium object link and writer completeness | Kinds behind virtual slot +0x20 and the unresolved dword stores: read the vtables and sites |
| Group member / marked-cell population / sight and leash before assignment / EN/RU named bodies | Gather world units on marked cells, apply relation/invisibility/health filters and fallback; leash removes candidates beyond radius, after which assignment may clear a target. | R2-ENGINE-1112 High local branches, Medium eligibility/end composition | Grid/world producers and other target writers: read them; run pursuit ending and next action |
| Player input / selected read units / message cells 0x12..0x26 / RU input, EN/RU setters | Named input chains reach 13 distinct order values. Types 0x1e/0x25 can reach order 0xd; 20 setters cover 17 values in total. | R2-ENGINE-1107 High local chains/setters, Medium complete input population | EN cells 0, 1, 2, 4, 5, 6, 8, 10, 11, 15, 18 call the RU method addresses minus 0xa039 and cells 3, 14, 16, 17 take the default as in RU (R2-ENGINE-1128 High; method bodies not compared); order 4 comes from the selector's order-0x19 arm, 0x11 from a script action; 0x18/0x19 setters have no six-form reference (Medium). EN cells 7, 9, 12, 13, 19, 20 unread |
| Any / Group record / named programme setter reached / EN/RU | Store programme constant independently of member order, including programme 4/5 with member order 1 and programme 3 with member order 0xc. | R2-ENGINE-1113 High named stores, Medium writer completeness | Script actions call the named setters (R2-ENGINE-1123); register/bulk/load writers and initial programme: classify remaining stores and run mission start |
| Single-player scheduling / reached clock body / EN and RU separately | EN has a selectable period and deadline-based tick calls; RU bodies use different app offsets. Session counter increments and Group gating retain their earlier local grades. | R2-ENGINE-1114 High local code, Medium delivered cadence; R2-ENGINE-778/1016 inherited grades; R2-ENGINE-1129 High imports (timeGetTime both roots) and RU mode selector L2.01223, Medium period writer | RU period writer (functions at L2.06337 and L2.06329), mode field meanings and historical speed choice: read them; run period/clock trace after LOAD |

Writer census and body correspondence remain bounded: R2-ENGINE-1104 and
R2-ENGINE-1115. No unlisted unit kind or root acquires an unconditional row.

## Attack with a spell staff

By the code read, a unit that passes the mage predicate and holds a weapon
with a Spell, ordered to attack a unit, and not taken elsewhere by an earlier
planner branch (target +0x4c bit 8, a winning slot draw, L2.00056, or L2.01527
for owner +0x2c zero), enters the weapon's Spell cast at the first admitted
cell centre: the Chebyshev distance from its cell to the nearest cell of the
target's body is at most its range, which is +0x12c unless an affordable
selected spell sets it (R2-ENGINE-1144, Medium; no run was made). If the
hero's base +0x12c is 1 (unread; the default of 1 is a creature's) and a Staff
(range 6) adds 5 (RU equip body), that is 6 cells to a one-cell target and 6 to
the nearest cell of an Ogre, Troll or Dragon; measured to their anchor cell it
is 7 or 8 from the high-x/high-y side (R2-ENGINE-1139, R2-ENGINE-1140). No
read body on the route tests sight, and 92 direct callees are unread
(R2-ENGINE-1136); no read body re-tests a distance after admission
(R2-ENGINE-1138). A computer unit's target can enter through lists that, on
their marked-cell path, test only the anchor cell; their callers were not
counted (R2-ENGINE-1137, R2-ENGINE-1136). Command +0x78 is read and not
written in the bodies read; who writes it is Unknown (R2-ENGINE-1142). A book
cast at a unit compares the same distance with spell +9, except for a self
target; the packed-cell metric's scan hits lie after the support entry
(R2-ENGINE-1143, Medium). A staff's item kind is 1 in EN (R2-ENGINE-1141).

## Keyboard orders and the defend order

The mission letter keys G, T, R, P and L send input types 0x17, 0x18, 0x14,
0x13 and 0x12 to the selection; on the RU input chain these become orders
0xb, 0xc, 0x16, 0x1a and 0x17 (Medium for the order numbers: RU chain only,
conditional on the handler's admission, R2-ENGINE-1107), while A, M, D and S
only set a cursor mode (R2-ENGINE-1179). Order 0x1a searches square rings
around the unit for an accepted cell, stores it at command +0xa with
prepared +8 = 0xc, and with none calls the order-0xc setter
(R2-ENGINE-1183). Order 8 approaches the defended unit beyond range r,
command byte +0x70. Within r a defender with a book plans spell 0x18 on the
defended unit when (defender word +0xa0 is 0 and defended +0x94 < +0x96) or
(+0xa0 nonzero, below defender +0x9c + 3, and +0x94 < +0x96 / 2), the spell
is in its book and its word +0xc is at most defender +0x9a
(R2-ENGINE-1185); in every other case it runs the attack helper over an
area list around the defended unit (R2-ENGINE-1188). At distance below 2, when the prepared
byte left by the helpers is 7 or outside 5..9, it steps r cells away along
the line between the two cells, clamped to 8..size - 9; no read body reads a
facing field (R2-ENGINE-1184).

The application key routine handles Pause, Esc, F1, F2, F3, Num+ and Num-
(R2-ENGINE-1181); Num+ and Num- step the game speed index 0..8 through the
setter that R2-ENGINE-620 describes (R2-ENGINE-1182). The view key
routine and its three key tables are read in R2-ENGINE-1178 and
R2-ENGINE-1180 (the third table's arms); 39 of the 46 key
entries of the RU install text have an arm in the read routines
(R2-ENGINE-1186).

## Cursor modes, the order panel and the Defend list

This path reads no list stored for the defended unit; it reads the map's
shared scratch list, rewritten by each call. Each call rebuilds it from the
units on the cells within r of the defended unit's cell, where r is byte +8 of the object
at defended +0x1c0. The search takes the centre cell first, then square
rings. The list is then filtered with the defended unit as the reference
(R2-ENGINE-1188; the EN search body and the filter meanings are Medium).
With the list empty, order 8 reads as calling the idle planner for the
defender, which stores no order in its own body; the ordinary selected branch
prepares 6, and the default branch prepares 0xb or 0 or runs a mage helper.
The order 8 planner has one direct caller; the scan excludes only direct
references, not calls through computed addresses (R2-ENGINE-1189; the composed
path is Medium).

View mode +0x9b4 sets the cursor:

- modes 1, 2, 4, 5 and 6 set the attack-or-move, move, defend, cast and
  0x1a cursors through the mode setter;
- mode 8 maps to the 0x1d cursor in the table, but no setter path holds
  mode 8, so that cursor needs another source (Unknown);
- modes 3 and 7 set none.

A map click sends by cursor:

- type 0x19 on a hovered unit, or 0x16 at the point;
- 0x16;
- 0x1b on a hovered unit;
- the cast;
- 0x1a at the point;
- 0x1d at the point, only if another source sets the 0x1d cursor (Unknown).

A click with a nonzero mode stores mode 0 and, by the pushes in the window
(the call itself lies past the window end, Medium), posts 0x40b to the panel;
a click with mode 0 does neither (R2-ENGINE-1190).

The EN order panel has 8 buttons in two rows of four in its hit-test
geometry. Pressing an enabled button i (message 0x201, read as a left press)
asks the view for mode i + 1 through the mode setter (the wParam + 1 argument
is read in an older listing, Medium). The setter stores the mode when the
allowed-mode mask has its bit (Medium: L2.06497 as the allowed-mode set). For
arguments 3, 7 and 8 it reaches three immediate senders (types 0x17, 0x18 and
0x14), each followed by mode 0 (R2-ENGINE-1191, R2-ENGINE-1179). Which buttons are enabled, and the RU
panel, are Unknown.

Prepared 4 reads as approaching command +0x18. Prepared 0xc reads as walking
to cell +0xa, waiting one call there and clearing prepared. The order 0x1a
planner stores the cell its ring search added last, if the list append counts
the node (R2-ENGINE-1192, RU; Medium).

In EN, key release L2.06589 clears the Shift, Ctrl and Alt flags and sends
nothing to the application's +0xcc object in its own body; it calls unread
routines. The view key routine sits in view slot +0x6c. The key cell bytes are
the same in both roots; in EN the application routine forwards 0xc0 and drops
0xbc..0xbf (R2-ENGINE-1193). The RU E and Alt routines have the EN
instruction structure (the E routine's class handle and callee are Medium),
and the RU speed cells match EN (R2-ENGINE-1194).

The Ctrl+N, O, K, T and C toggles are process globals (section not read).
The census finds no direct writer of T except its key arm; a write through a
computed address is not excluded (R2-ENGINE-1195; persistence over a mission
change or restart is Medium).

## The L key order (0x17)

The L message (type 0x12) carries the sender in word +5 and the selected
ids. With message byte +4 of 1 or more (Unknown for L, which does not write
it) the dispatcher looks each id up in the sender's unit list and drops an
id it does not find or whose unit has +0x54 = 0x10. The found units leave
their Groups and join one new Group; the Group method then calls the order
0x17 setter on each member and sets programme 0x13. Both roots are read;
the sender's list and the stretch to the type switch are Medium. The path
for byte +4 below 1 is not read, and that L takes it is not excluded
(R2-ENGINE-1198).

The setter tests nothing. It runs the reset (the stop routine and command
+0x50 = 0), then stores order 0x17, command word +0 = the unit's own cell
and prepared +8 = 0. In the bodies read it writes no target, primary +9 or
+0xc; its unread callees are not excluded (R2-ENGINE-1199).

On the Group tick, programme 0x13 calls the support routine on each member:

- a player unit that is not a mage gets prepared 0, which the executor maps
  to no handler: on this route it reads as standing, with no walk, turn or
  attack call (Medium); whether the per-unit loop `L2.01724`, whose caller
  is Unknown, reaches the selector and the idle planner is Unknown;
- a player mage with mana 5 or more and bit value 8 heals itself or the
  first wounded ally in list order with spell 0x18 and casts 0x1b on itself;
  buffs 0x14 and 0x15 need bit 0x10 and mana 10; buffs 0xd, 0x13, 4, 8 and
  0xc need bit 0x20 and mana 40; RU read to its end, EN for the first 256
  instructions;
- a computer member gets prepared 0xb, or 8 when its mage branch casts, and
  a computer mage heals with 0x18.

A primary action in progress, the low-health flight and the executor tail
on an empty route still act. No direct route from the group tick reaches the
selector under programme 0x13 (R2-ENGINE-1200; the readings "player" and
"computer" and the composed "reads as standing" are Medium; prepared 0xb
under this programme is Unknown).

The census finds no compare with 0x17. The order ends when another write to
unit +0x50 replaces it: another input order, a script action on the Group,
or a programme handler that stores order 0xc (R2-ENGINE-1201; the negatives
are Medium). The direct references give one route, the L key through the
input arm; no panel button reaches it by a direct reference, and other
builders of a type-0x12 message were not searched (R2-ENGINE-1202; computed
calls are not excluded). A save holds unit +0x50, Group record +0x20 and
command +8; command word +0 lies in the saved block but no save claim names
it, and whether a loaded unit resumes the route is Unknown (R2-ENGINE-1203).
