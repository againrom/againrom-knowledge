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
