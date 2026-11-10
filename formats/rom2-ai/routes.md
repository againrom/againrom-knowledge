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
