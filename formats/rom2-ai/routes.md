# Search, substitutes, stamp and empty routes

## Route search

| Input/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Walk / approach / repath | L2.05182 calls search L2.05174 (EN L2.05175), flag 1 for walk/approach and 0 for repath. | R2-ENGINE-871 High bodies, Medium caller census | Indirect callers |
| Search goal equals unit cell | Search exits. A zero goal is not otherwise rejected in the read walk/search bodies. | R2-ENGINE-883 High local tests, Medium trajectory | Grid results for goal 0 |
| Goal word remains 0xffff; flag 1 | Choose substitute L2.05179(unit,0,goal,(distance>>2)+4); send notice if sufficiently far; run stamp test; nonzero substitute: L2.05181; zero: clear route list +0x1c8. | R2-ENGINE-871 High sequence; R2-ENGINE-873 High stamp operands | A nonzero substitute leaves a route unless the unread helpers L2.05231, L2.05229, L2.05228, L2.05185 empty it (R2-ENGINE-879, Medium); a substitute equal to the unit's cell stays open |
| Goal word remains 0xffff; flag 0 | Use L2.05189(unit,target) or substitute with limit 8, then builder L2.05190. | R2-ENGINE-871 High calls, Medium helper meanings | Those helper/builder bodies and exact branch selection |
| Approach retry | Re-search when movement retry byte +9 exceeds route length/3+1 or cached target +0x7c differs. At route end with +0x8c equal target cell, increment retry and clear cache. | R2-ENGINE-874 Medium composition including read approach flow | Actual retry frequency and route/list outcomes |

## Substitute and route construction

| Input/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Substitute radius r | Enumerate square rings 1..r-1 around goal. Stop on first ring holding a L2.05180 word below 0xffff; choose strictly least word there, first probed on a tie. With none, return 0. | R2-ENGINE-878 High algorithm, Medium grid interpretation | L2.05180 unread: least word does not prove reachable cell |
| Size-aware stamp distance | Per axis: if a<b, max(b-a-size+1,0), else a-b; return maximum axis. | R2-ENGINE-878 High body | Reading the range comparison as combat reach remains Medium |
| Notice | Substitute farther than 2 from goal if L2.05188(goal) succeeds, otherwise farther than 1: notice message 0x92 with (1,0,owner). | R2-ENGINE-871 High branch; R2-ENGINE-878 High goal test; R2-ENGINE-879 High message body | Delivery/display effect |
| Nonzero substitute/end cell | Append end cell to route +0x1c8; trace toward own cell by least neighbour word plus cost. Straight/diagonal cost 2/3, or kind-1 cell byte and 1.5 times it. Bounds start at 8 and end at map dimension+8. | R2-ENGINE-879 High body, Medium helper semantics | Grid/list operations, no-candidate pass and stalls |
| Trace exceeds 1000 steps | Clear route. Otherwise finalize through L2.05231. When substitute equals own cell, body appends that cell and does not enter the trace loop. | R2-ENGINE-879 High body, Medium empty-list composition | Whether finalization empties a one-cell route |

## Stamp and empty-route tail

| Input/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Flag-1 failure after substitute/notice; programme 18; prepared +8=5; distance(substitute,goal,size)>L2.01728(unit) | Store command word +0x16=server word +0. Substitute 0 is not an exclusion. Range helper returns the slot-spell range when command byte +0x78 is set and unit word +0x9a reaches the spell's word +0xc, else unit +0x12c. | R2-ENGINE-873 High tests/stores; R2-ENGINE-878 High distance; R2-ENGINE-846 High constructor 65500; range interpretation Medium | Untraced stamp writers; original occurrence |
| Next Guard B case 1/2 matches | Walk to command +0x9c before target case; target +0x20 remains stored by this path. Expiry permits case 3 again if a target remains. The cycle continues while a failed full search finds the substitute beyond range (Medium). | R2-ENGINE-874 Medium cycle; R2-ENGINE-844 High order/stores, Medium clock interpretation | Exits are inferred, not read at their writers: successful search, +0x20 cleared by leash or assignment with writers not censused, target death. The order held when the tail fires and the planner re-preparing the attack are unread; route construction returning an empty list and +0x9c before a first hit stay open. No run (mission 10, hero on an unreachable cell in sight of Troll.4) records the cycle (R2-ENGINE-874); duration in seconds is unmeasured |
| Same-owner area blow on a member | The hook rewrites +0x98 and +0x9c on each pulse; a member with target +0x20 stays in the target case, one without a target and below max health takes the +0x9c walk while server dword +4 is within 10 of +0x98. | R2-ENGINE-989 Medium composition; R2-ENGINE-844 High order/stores | Darkness sight effect on the kept target; observed trajectory |
| +0x9c=0 during window | Prepared move hands cell 0 unchanged through walk to search. In the five saves +0x9c is 0 in 147 of 150 blocks, none stamped (R2-SESSION-346); a stamp on a never-hit creature is not saved. | R2-ENGINE-883 High bodies, Medium trajectory; R2-SESSION-346 High values, Medium state reading | Whether route, substitute or empty list results; not a proven constructor value |
| Walk/approach route empty; or repath local list +0x1e4 empty for destination +0x76 | Set movement dword +0x98=1; walk/approach also set +0x76 to own cell. | R2-ENGINE-872 High local writers, Medium census | Untraced base accesses and how often route empties |
| Common tail sees movement +0x98 | Clear flag, read unit order +0x50; do not directly change command +0xc or assigned target +0x20. | R2-ENGINE-872 High body, Medium sole-reader census | One traced reader is not a global reader proof |
| Tail order 1 | Stop then reset to unit order 0xc, post=own cell, prepared +8=0. | R2-ENGINE-880 High local stores, Medium helper meanings | Reservation effects |
| Tail order 0xa | Set command +4=1 and +8=0; within distance 3 of command +2 take next point in command list +0xb4. | R2-ENGINE-872 High local body | List/helper effects; order name not established |
| Tail order 0x18 | Primary +9=255. | R2-ENGINE-872 High | Later planner clear conditions |
| Tail every other order | Prepared +8=0, call re-pick L2.02425: keep units within attack range and a running nearest distance; hostile least-turn result prepares 6 (later turn tie wins), else computer look-around 11 or player idle 0. | R2-ENGINE-872 High tail; R2-ENGINE-504 High RU local re-pick, Medium EN/helpers | Candidate-list filters and planner/executor timing |
| Measured grouped creatures | 138/139 Unit-class records in five mission-10 snapshots store unit order 0xb, selecting the last tail branch at those saved ticks. Programmes calling L2.01576 can instead store order 0xc. | R2-ENGINE-884 High counted values, Medium continuation inference | Snapshots of one run are not all live orders or runtime validation |
| Other read programme bodies | 0,2,3,4,5,17,19,0xff directly read none of command stamp/hit fields or movement +0x98. | R2-ENGINE-882 High direct body absence, Medium transitive effects | Unread callees; not a universal no-step-away proof; units in no group: the executor reads the group record without a test; every run unit having a group is inferred (R2-ENGINE-882, Medium) |

## Health thresholds, difficulty and attack order

The stat reader stores them in command +0x40 and +0x44 for a unit whose template loads (Medium: loader condition, script-created creatures; R2-ENGINE-1001). The map arm scales a non-Human record by the difficulty word when Session +0x74 is 0 and the value is not 2 (R2-ENGINE-1002). In the bodies read, the only direct caller of the threshold post-step is the group tick (byte scan; indirect and jump routes unread, Medium; R2-ENGINE-1003). In the bodies read, LOAD leaves the thresholds as saved (R2-ENGINE-1005).

## Tail branch of a finished route

The flag branch is read at L2.02883: order 1 stores order 0xc through L2.02858 and L2.02636; orders 0xa and 0x18 keep their branches; every other value clears +8 and calls re-pick L2.02425 (R2-ENGINE-1020). The route from the Group pass to this tail is bounded by R2-ENGINE-1017.

The walk prepare L2.02401 stores command +8 and +0xa and reads no candidate list (R2-ENGINE-1042). Its movement +0x98 sites are census rows; arrival stays unread. The found-target branch L2.01525 is R2-ENGINE-1043.

The programme-3 route L2.02466 makes no direct call to the selector; no reference to loop L2.01724 was found in the Medium census, and an indirect or computed route is not excluded (R2-ENGINE-1041).

Counted run (EN, one mission-10 launch, one hero): R2-ENGINE-1072 (loop L2.04213 and wrapper L2.02403 had 0 hits in the span; route serving the hero Unknown in launch 1), R2-ENGINE-1073 (actor step once per scheduler counter value per unit; the counter advanced 16.1 and 32.1 per second in the two launches), R2-ENGINE-1074 (a ground click wrote order 1 and moved the hero to another Group object).

Continuation reads: R2-ENGINE-1075 (EN return addresses L2.06121 and L2.06122 of L2.04984 lie in the idle planner's no-candidate arm and in the programme-3 route; which of the planner's six E8 callers served the hero is Unknown in launch 1), R2-ENGINE-1076 (RU programme 4 and Guard A bodies, including Guard A's arm L2.01733(member, 1), which does not call the idle planner; the empty-Group store of programme 0xff), R2-ENGINE-1078 (a second EN launch: with no attack order the hero read order 0xc and a strike state on another owner's unit, and the idle planner's entries by return address), R2-ENGINE-1077 (256-instruction masked comparison of RU and EN bodies; 16 pairs equal, seven differ at a later index, the EN twin of L2.05852 stays ambiguous).

## Admission, arrival and retained targets

The read walk's own-cell arrival stores movement +0x90; search-empty branches
store +0x98 and the common tail replans by order (R2-ENGINE-1108). A direct
order-0xc store in a reduced EXE sweep is not complete writer closure
(R2-ENGINE-1104). Target initialization and command-target clears must use
their actual bases (R2-ENGINE-1105). Strike initiation and application have
different distance calculations, with an additional RU application rejection
(R2-ENGINE-1106). Player input/setter and Group programme paths remain
separate (R2-ENGINE-1107, R2-ENGINE-1113).

Hit/event and Darkness field stores do not guarantee that an existing route
or assigned target survives the next scheduled action (R2-ENGINE-1109,
R2-ENGINE-1110). Scorer L2.05238 and filtered/leashed candidates supply
specific target-retention/loss paths (R2-ENGINE-1111, R2-ENGINE-1112), not all
pursuit endings. Clock delivery and EN body coverage retain their boundaries
(R2-ENGINE-1114, R2-ENGINE-1115). The [planner table](orders.md) names the
search/occupancy reads and runtime controls still needed.

The orders page's static continuation adds bounded clauses. The selector calls
the stop chain for order 3 when the target's +0x54 is 0x10 (R2-ENGINE-1121; the
order-0xc ending of the chain is R2-ENGINE-1020; the meaning of 0x10 is Medium)
and stops order 1 on the +0x90 arrival flag at the destination
(R2-ENGINE-1121). The attack planner stores no order; with target +0x4c bit 8
set it runs the idle planner, and with a null or self target it first takes the
first unit of L2.02117 and runs the idle planner only when none is found
(R2-ENGINE-1120). Byte +0x4c bit 8 has a bounded census of byte stores
(R2-ENGINE-1122). A script action switch reaches Group methods; the case
routines read store only into a result array
(R2-ENGINE-1123, R2-ENGINE-1124). The slot +0x34 blow and approach wrapper
(R2-ENGINE-1125), the constant preference bytes (R2-ENGINE-1126), the standing
helpers (R2-ENGINE-1127), the EN input cells and order-4/0x11 setters
(R2-ENGINE-1128) and the clock imports and RU mode selector (R2-ENGINE-1129)
keep the grades stated on the orders page. Root correspondence
remains bounded by R2-ENGINE-1115.

The orders page's staff section composes, by the code read, the attack with a
weapon that has a Spell (R2-ENGINE-1144, Medium; no run was made), unless an
earlier planner branch takes the unit elsewhere: admission by the Chebyshev
distance to the target's nearest body cell at a cell centre (R2-ENGINE-1139),
body sizes from byte +0x49 (R2-ENGINE-1140), no sight operand in the bodies
read and 92 direct callees unread (R2-ENGINE-1136), anchor-cell tests on the
marked-cell path of the sight lists, whose callers were not counted
(R2-ENGINE-1137), no distance re-test after admission in the bodies read
(R2-ENGINE-1138), command +0x78 read with its writer Unknown
(R2-ENGINE-1142), the book cast's spell +9 admission (R2-ENGINE-1143) and the
EN item kind of a staff (R2-ENGINE-1141).

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

## Health and mana restoration

These are readings of both executables; no run of the original measured a
restoration.

The refill's address stands at three dword sites per root, slot 5 of three
tables per root; two RU tables have a read store of their base and the other
bases are inferred from the entry offset (Medium). The world pass calls the
refill on each object of its list once per 16 session ticks, when Session
+4 is 12 mod 16. Session +4 is the counter the session tick advances by 1
per call in which its +0x94 is nonzero; Session +0 advances once between
two passes. At the scheduler periods of the Speed table one pass takes 496
to 2000 ms, 992 ms at Speed 4 (R2-ENGINE-1204; bodies and arithmetic High
in both roots, milliseconds Medium).

Mana is restored on every pass while below its maximum. Health is restored
on one pass in four, when the signed remainder of Session +0 by 4 equals
that of the unit id +4, and only when divisor +0x98 is not 0: 16 and 64
counter advances between two restorations (R2-ENGINE-1205; that neither
value is negative is Medium).

Each restoration adds hundredths: 2 * max * (+0xde + 100) * K / +0x98 to
health and max * (+0xe2 + 100) * K / +0x9e to mana, with C truncating
division. The remainders stay in bytes +0xa2 and +0xa3; the pool is capped
at its maximum after the remainder is stored. With the defaults 100 and 50
and modifiers 0 each restoration adds 2 % of the maximum. K is 3 when
Session +4 minus the unit's stamp +0x138 exceeds 80, else 1; the move step,
the strike senders and three further senders, called from the address range
of the cast routine, set the stamp. Effect tag 8 adds to +0xde without a
clamp; tag 11 adds to +0xe2 under an unread test (Medium) (R2-ENGINE-1206;
formulas High, sender meanings and writer completeness Medium).

A unit with +0x54 = 0x10 or with health 0 or below restores nothing; below 0
the refill takes 1 health on one pass in four, and corpse ageing lowers
health separately, so the total loss per pass is not this figure alone. The
body tests no owner, side, player, order, combat or cast. Which classes the
list holds is Medium (R2-ENGINE-1207). A census of word stores to health and
mana finds no other restorer of a living unit over time among the stores it
classified; nine RU word stores are unclassified and are not excluded. A
per-player pass revives a hero at health -40 or below to full health and
mana, and no pass over time was found calling the remaining full sets
(R2-ENGINE-1208, Medium). The RU Unit serializer holds the pools, maxima,
divisors, remainders, stamp and the raw block that covers the modifiers;
Session +0 and +4 in a save are Unknown (R2-ENGINE-1209).
