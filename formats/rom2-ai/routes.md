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
| Tail order 0xa | Set command +4=1 and +8=0; within distance 3 of command +2 take next point in command list +0xb4. | R2-ENGINE-872 High local body; R2-ENGINE-1624 High bodies, Medium composition | Retry duration at an unreachable point |
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

Light and Darkness lower the sight byte +0xa5 that the ray scans and the planner, movement
and group routines read: Darkness for a Human or Humanoid occupant, Light for a Unit occupant,
by 1+P/30 cells while the byte is at least that amount (below it the Effect arms wrap the byte
modulo 256 and no arm clamps; whether a reader clamps is not read except the ring walk's limit of
19 rings, R2-ENGINE-502), with no owner test (R2-ENGINE-1236, R2-ENGINE-1237, R2-ENGINE-1238,
R2-ENGINE-1239). The unit's own player receives the changed word (Medium; for Human and Humanoid
units with word +0xe in 0x21..0x3f another player may receive it, Unknown) (R2-ENGINE-1240), expiry
adds the shift back (R2-ENGINE-1241), and the per-class result is Medium (R2-ENGINE-1242)
(High for the arms, Medium for the reach, the delivery and the fog link).

Area reach of these Effects (Medium throughout): Light and Darkness pulse onto cell entry +4 only, so kind 1 or 2 units
(Units rows with movementType 1 or 2, which include Ghost and Bee; R2-ASSET-327, R2-ASSET-328, R2-ENGINE-1530) are
reached and kind 3 units (Bat_Sonic, Dragon, Succubus; a flier by name only) and buildings are not
(R2-ENGINE-1535); in EN sight is base once per 16 updates between the end of one pulse's Effect
and the next pulse (R2-ENGINE-1534). Blizzard passes the kind 1 or 2, kind 3 and building occupants of
its one sampled cell per pulse to a damage object that admits all three (R2-ENGINE-1533; cell validity band `8..W - 9`, `8..H - 9` confirmed High in both roots, R2-ENGINE-1668; resolver rejection by
target flag +0x4c bit 8 is unread in EN); the cells and pulses of the other area spells
are in R2-ENGINE-1536; Fire Ball hits a multi-cell unit once per covered cell (R2-ENGINE-1537, Medium).

The orders page's static continuation adds bounded clauses. The selector calls
the stop chain for order 3 when the target's +0x54 is 0x10 (R2-ENGINE-1121; the
order-0xc ending of the chain is R2-ENGINE-1020; the meaning of 0x10 is Medium; the immediate stores of 0x10, mov with displacement 0x54, are two per root and on the death path, register and other store forms unread, R2-ENGINE-1547, amending R2-ENGINE-1121)
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
EN item kind of a staff (R2-ENGINE-1141). The weapon constructor L2.05910 is
called by the record loader L2.07611 and four other routines (R2-ENGINE-1367,
Medium). No installed data.bin item row and no itemname.pkt record has 12 as Slot or
kind except the boots (Medium that either is the engine's item type; the object
+0x44 type of a built item is read only through R2-ENGINE-711 and
R2-ENGINE-265); a Staff is kind 1, attackType 3, physical 0 to 0 (R2-ASSET-297, R2-ASSET-298),
and its cast pays no mana (R2-ENGINE-1441, Medium).

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
at defended +0x1c0 (5 for every unit read, R2-ENGINE-1440). The search takes the centre cell first, then square
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

The reset's call RU L2.02822 is the footprint pass L2.02410(unit, 0, 1). It
moves no unit: it clears the unit's footprint reference from the ring of map
cells around its position, sets it on the position square and stores movement
+0x80 = the position cell (R2-ENGINE-1213). A unit that dies in the middle of
a step is cleared from both cells of the step (R2-ENGINE-1215). By a reference
scan of allods2.exe in both roots (Medium) the pass has one direct call site
outside the terrain code, in the idle planner tail (R2-ENGINE-1216).

The EN load calls the Players reader (the unit records, Medium), then the
Buildings reader, then makes a new terrain object and reads the per-cell
records into it (R2-ENGINE-1220). The EN Building Serialize calls neither
the building init nor the setter (R2-ENGINE-1221). By reference scans of the
EN executable (Medium), no call site of the ground place, the footprint pass
or the positioning routine lies in the read load bodies, and, in the bodies
read, the footprint fields of an entry after a load are the file's numbers,
replaced by objects on a key hit; the routines that make a unit or a
building from the file are unread (R2-ENGINE-1222).

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
of the cast routine, set the stamp; a cast through the unit sender stamps it, so K is 1 for 5 passes after a cast (Medium: 5 or 6 by tick order) and the restoration does not pause; a move or strike also stores a stamp, so K stays 1 while a unit keeps moving or striking (R2-ENGINE-1565). Effect tag 8 adds to +0xde without a
clamp; tag 11 adds to +0xe2 only for a unit with mask 4 of +0x4c (named the mage bit, Medium) (R2-ENGINE-1564, R2-ENGINE-1206;
formulas High, sender meanings and writer completeness Medium).

An item's tag 11 value M adds M percent to the mana rate of a mage: each pass restores max * (M + 100) * K / D hundredths, D being the Units row's MP regeneration when the reader stores it (114 of 180 EN rows hold 0 and the effect of a 0 on D is Unknown), so an item of M 50 turns 2 % of the maximum per pass (D 50, K 1) into 3 %; the interval is the world pass of 16 session ticks (R2-ENGINE-1566).

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

## Computer caster spell and target choice

A computer mage (mage bit, owner +0x2c nonzero; "computer" at Medium) chooses
book spells in three routines: the offensive choice, called by the attack
planner after its slot draws, and the heal routine with its support choice,
called by the idle wrapper. Each has one direct caller in each root; the
priority routine has no direct reference in any form searched
(R2-ENGINE-1244, R2-ENGINE-1250). The heal routine heals the caster when
hurt, else the first wounded unit of the Heal-radius area list around the
caster that the relation filter keeps (bit 2), else calls the support
choice: Shield on the caster when its effect bit 27 is clear, then one drawn
list member of the caster's owner, healed when hurt or given a drawn
Defensive book spell at full mana (R2-ENGINE-1245, R2-ENGINE-1246). The
offensive choice casts a drawn book spell of ids 1..28 whose Defensive value
is not 1 at the victim, unless Mind above 59 and a draw below 30 of 100 send
it back to the planner; the three routines never choose Slow (29) from the
book (the loops stop at 28); that no other routine chooses a book spell is
Medium (R2-ENGINE-1244, R2-ENGINE-1247; per-spell census R2-ENGINE-1248). The EN bodies equal RU
after masking (R2-ENGINE-1249). How often these routines run per simulation
tick outside the Group pass is Unknown.

No body read chooses a point for an area spell. The offensive choice casts at the
victim unit; the special cast uses the victim's cell (for Acid Stream, a cell
beside the caster toward the victim). No body read tests the own or allied
units in the area. The centre a target-mode area spell takes is Unknown
(R2-ENGINE-1490, Medium for the composition).

In a Group pass (once per 16 counter values, RU; EN Medium) programmes 1 and 2
(a member at its post, with L2.00490 nonzero) and 18 (the same, with two
further gates), 3 and 5 reach the heal routine for a member without a target,
programme 0xff when the idle planner finds no candidate, and programme 19 for
every member; runs outside the Group pass are Unknown (R2-ENGINE-1420). The
post-step then decides break-off for the same member: when Teleport is in the
book and mana covers its cost it casts Teleport at a cell scanRange + 1
(back-off: scanRange - 1) from the hostiles' mean point; otherwise it walks 3
cells away from that point. Each spell routine also holds an Invisibility arm
and a walk-to-cell arm behind a repeated Teleport admission; the callers' entry
test makes them unreachable while the admission is pure (Medium). With no
hostile it takes the idle planner (R2-ENGINE-1418). Break-off runs once per
Group pass per member from the post-step; flee also runs from selector cell
0x16 (cadence Unknown). The bodies read keep no flight state; each decision is
re-made (R2-ENGINE-1419). Withdraw and Wimpy are hit points; of 180 populated
Units template rows, 40 book Heal, all by knownSpells, and of them only
Succubus has a Wimpy (R2-ASSET-293; Medium: books of heroes, summons and loaded
units are not read). The player threshold rewrite
L2.02445 sets +0x40 to HPmax x {0, 2v, 3v}[class] / 100 (R2-ENGINE-1421). No
heal, break-off or Group arm body read reads the difficulty word
(R2-ENGINE-1422, Medium).

The flee test and the flee body use the radius of the movement block's byte +8
plus 5. The byte is the block constructor's constant 5 (R2-ENGINE-1438); the
traced loads read it at 11 sites and write it nowhere, and all 356 saved Unit
and Human records hold 5 (R2-ENGINE-1439, Medium). So the flee radius of R2-ENGINE-506 and
R2-ENGINE-1418 is 10 cells for the test and the walk arm; the Teleport arm lists 5 cells, the
back-off radius is the constant 2, and the guard list radius is 5
(R2-ENGINE-1440). A computer healer that meets both the break-off test and the
heal on one Group pass walks: the post-step's store replaces the prepared heal
with the walk (R2-ENGINE-1442, Medium). When the flee finds no living hostile
in its list it runs the idle planner instead, which may prepare a heal; the
Teleport arm's store is unread (Medium / Unknown). A heal cast during a walk
that is under way is Unknown (R2-ENGINE-1442).

## Guard and summoned-unit boundaries

Guard setup reads installed Scanning/MinimalGuardRange=8, with getter fallback
10; the full source-loader and writer census remain Medium (R2-ENGINE-1252).
For r=0, Guard A uses max(b,8); Guard B keeps b>=8 or writes 20. Baselines 8
and 20 give next-leash 11..13 and 23..25 under the Medium draw interpretation.
Nonzero arguments, LOAD and byte narrowing keep a universal minimum Unknown
(R2-ENGINE-1253).

The leash uses the current member mean. Losing target assignment can prepare
return to the post, or Guard B's earlier special goals, but does not establish
a fixed journey from home or an immediate visible turn (R2-ENGINE-1254,
Medium composition). The count pass calls death for the first matching
linked-list unit after its draw<n test, at most one per caster key per pass;
its map helpers and visible effects keep the composition Medium
(R2-ENGINE-1255). This does not close caster loss, an independent timer,
mission failure or LOAD-dependent removal (R2-ENGINE-1256, Medium bounded
census, completeness Unknown).

## Cast wind-up and recovery

A cast whose routine try succeeds at actor step s runs the apply in step s + c
and returns to substate 0, with completion byte +0x136 = 1, in step s + c + r; a
count of 0 lasts 256 steps. c is caster byte +0x134. r = low byte(+0x135 + t + d
+ k): t = 0 unless a Human caster holds a weapon, else clamp((weapon word +0x4a +
(30 - Reaction) * 5) / 12, 0, 12); d is a 0..3 draw; k is the spell's
Complication byte on a book cast and 0 on an item or staff cast. In the arms and suffix read, neither count reads the target; the per-spell arms are unread. The defaults are 8 and 4; the Units and Humans readers and
weapon equip write the table values (its handling of -1 is unread), unequip restores 8 and 4, and one unit builder
writes 1 and 1; the writer census covers direct operands only. The apply's
common suffix builds the effect object and adds it, or its projectile, to a
world list; a Delivery 2 spell's effect object enters the effect list max(1, floor(distance / Effect Speed)) steps after the apply (Medium cadence; the effect object's own step is Unknown), a Delivery 1 spell's at the apply, and the walk to range is Unknown (R2-ENGINE-1567, R2-ENGINE-1568, R2-ENGINE-1569). R2-ENGINE-1260 (High
stores, Medium completeness), R2-ENGINE-1261 (High expression, Medium names),
R2-ENGINE-1262 (High step table, effect step Unknown), R2-ENGINE-1263 (High
terms, Medium column bindings).

Across a mission LOAD the actor-step fields +0x58, +0x6c and +0x134..+0x136 are read back into place, so a caster saved in wind-up or recovery resumes its own count (R2-ENGINE-1401: High for the RU body, Medium for EN; that the count runs to its end is Medium, postload callees unread) and begins its next order no earlier than R2-ENGINE-1262's s + c + r (Medium). A loaded timed area effect with a nonzero stored +0x4c continues from the stored +0x50 (R2-ENGINE-1400: High for a stored +0x4c nonzero, Unknown for +0x4c 0), and the read effect serializers store no caster field (R2-ENGINE-1398, Medium for the values after construction; what the effect +0x14 owner key names is Unknown), so a loaded effect reaches its target with no stored caster link (R2-ENGINE-1399, Medium); the HP amount without a caster is Unknown.

The ordinary script action wrappers are read independently of this page: local stores and calls are High, delegated effects Medium, and the bodies of actions 33 and 34 are read (R2-ENGINE-1450 amended), with action 11 moving items from U+0x7c to V+0x7c (R2-ENGINE-1699), action 27 running cell move and Guard setter (R2-ENGINE-1700), and action 31 setting Outpost respawn countdown +0xb4 (R2-ENGINE-1701). Unconscious units stay in their group list while dead units (+0x54 = 0x10, body or no body) are detached by the kill walker; checks 1, 5 and 8 count every unit still in its group list (living and unconscious units, not dead units; EN High, RU Medium for checks 1 and 8), check 20 counts Buildings rather than units (R2-ENGINE-1452 amended, R2-ENGINE-1696), and the check loop skips a check whose copied IR+0x58 byte is nonzero (High; that the 8000..9999 mark reaches that byte is Medium; R2-ENGINE-1451 amended, R2-ENGINE-1697). They change no rule above (R2-ENGINE-1314).

Script action 6 gives a group one of eleven method calls by its selector p0 (1, 2, 3, 4, 5, 10, 11, 14, 15, 17, 18); the method bodies are read in EN and RU, the other selectors do nothing, and the arms set member orders 0xb, 0xc, 1, 3, 8, 0xa and 0x11 and the group programme value; what these orders and programme values do in play is Medium (R2-ENGINE-1726, amending R2-ENGINE-1314).

Script check 2 reads x and y as bytes +0 and +1 of the object at unit +0x10 and is an inclusive rectangle test on them (High); checks 3 and 27 read the same getters (Medium) (R2-ENGINE-1730, amending R2-ENGINE-1450).

The final effects of the other read actions (R2-ENGINE-1314 amended): actions 21 and 24 build the caster Unit of a map building caster, which has no mage bit, casts once and is deleted (Medium) by the world phase when its recovery ends (R2-ENGINE-1740). Action 21 casts spell low8(p4) from cell (p0, p1) at cell (p2, p3) with power word p5, or 99 when p5 is 0 (R2-ENGINE-1741); action 24 casts spell low8(p2) from cell (p0, p1) at unit U with power word p3 (R2-ENGINE-1742); that the spell power equals the word up to 255 is Medium. Both roots ship five action 21 nodes and three action 24 nodes in the scenario maps (R2-ASSET-379). Actions 16 and 32 take a unit off the map and out of the unit registry (R2-ENGINE-1743); actions 17 and 33 put it back at the cell its stored position gives through unread getters (Medium) or within 3 cells, and action 18 takes U off and puts V within 3 cells of U's cell (R2-ENGINE-1744). Actions 19 and 22 give a unit to player P, each in a new group of its own (R2-ENGINE-1745). Action 23 adds p0 to player +0x3c, named gold (Medium) (R2-ENGINE-1746). Actions 12, 13 and 38 build or destroy an item in a bag (R2-ENGINE-1747). Action 20 drops a unit's bag as a sack at its cell (Medium); action 28 moves U's bag into V's, and when V's count is 1024 or more (Medium, unread count helper) U's items stay in its old bag, which nothing owns (Medium) (R2-ENGINE-1748). Actions 29, 30 and 34 set an area effect's remaining time, an effect's duration, or a unit's health, defence or absorption (Medium; R2-ENGINE-1749).

## Step vector, step time and parameters from map.reg

The step vector is the block at unit `+0x1c0`. EN `L2.06918` / RU `L2.07803` stores into it (R2-ENGINE-1470): speed
word `+0xa8`, heading word `+0xae` (dir 0 is dy = -1, dir 2 is dx = +1, 0 to 7; "north, clockwise" names this order),
signed x and y bytes per advance `+0xb0` and `+0xb1`, advance count word `+0xaa`. In EN `L2.03582..L2.06966` it has one
call, the step start of R2-ENGINE-1213; a census outside that range and of the other stores into the block is not made
(the arrival zeroes `+0xa8` and `+0xaa`), so that it is the only writer of the step values is Medium. Straight step:
the vector is speed * (dx, dy). Diagonal step: each axis is the speed * 0.707 truncated toward zero (speed 1 gives
(0, 0), speed 2 gives (1, 1), speed 63 gives (44, 44)); how a step with a zero vector ends is Unknown. The count is 256
divided by the larger axis byte, rounded up; the absolute-value helper `L2.04968` is unread and taken as an absolute
value, so the count is Medium (R2-ENGINE-1470).

Speed for a unit whose kind (`L2.03577`) is 1 (R2-ENGINE-1471): base * M, then the slope term, then divided by the mean
of the two cell terms, ((t(cell) + t(next)) and 0xff) >> 1, the sum kept in a byte (a mean of 0 counts as 8); M is
`SpeedMultiplier`, default 8. The term `L2.03599` has three return paths and one stores the cell byte shifted right by 2
back into the terrain byte array; which cells take it is Unknown. The slope term uses the byte at `this+0x944f4` that
R2-ENGINE-1340 reads as a height (Medium; tile classifier `L2.06958`/`L2.07473` indexes `+0x54146..+0x54150` and writes no height, R2-ENGINE-1665; client/server type-3 readers in R2-ENGINE-1667): dh < 0 (called climbing, the next cell higher) subtracts speed * |dh| / 64,
dh >= 0 adds speed * dh / 64, dh clamped to 32. Any other kind uses the base speed unchanged. The speed is clamped to
1..63, so a straight step takes from 5 to 256 advances (counted in advances, not game ticks; Medium for the helper).
No group term was found in the bodies read; the object that holds the base speed byte is not identified
(R2-ENGINE-1471).

| Section | Key | Default | Destination |
|---|---|---|---|
| Scanning | ScanShift | 7 | dword `+0x2a004` |
| Terrain | CostLand, CostGrass, CostFlowers, CostSand, CostCracked, CostStones, CostSavanna, CostMountain, CostWater, CostRoad | 8, 8, 9, 14, 6, 12, 11, 16, 8, 6 | bytes `+0x54147..+0x54150` |
| Path Finding | SpeedMultiplier | 8 | dword `+0x58d84` |
| Path Finding | StaticScanAhead, DynamicScanAhead, StaticRefreshRate, DynamicRefreshRate, DynamicByStaticLookup, StaticIsntNeeded | 5, 3, 16, 32, 3, 5 | dwords `+0x58584..+0x58598` |

All 18 keys are read by EN `R2.0113` through the config object named `World\Data\map.reg` (R2-ENGINE-1472: High for
the name string and the 18 calls, Medium that the object is built from the file). R2-ASSET-012 lists an entry `map.reg`
of 1020 bytes in `world.res` and `world_srv.res`; that the name reaches that entry is Medium. The values in it and the
readers of the cost bytes are Unknown. The search bounds are code constants: 8, 8, width - 9, height - 9.

Questions answered at the grade of earlier claims: which cell a step occupies and when it changes
(R2-ENGINE-1213, R2-ENGINE-673); the refused placement and the removal that does not compare (R2-ENGINE-1215);
when a route is rebuilt (R2-ENGINE-872, R2-ENGINE-871, R2-ENGINE-874, R2-ENGINE-1108). Unknown: the passability
test that refuses entry into the source cell during a step, and the tie rule between equal routes, whose bodies
(the path helpers at RU `L2.05231..L2.05230`) are unread.

A map building caster lists the units within its record radius of its cell, Chebyshev (Medium), 1 to 9 cells, with no range compare against the spell (R2-ENGINE-1560), and the pick of R2-ENGINE-1008 draws one of them. It starts a cast once per 16 session ticks through a built Unit that casts from state 13 with wind-up 1 and recovery 1 (R2-ENGINE-1561 amended, R2-ENGINE-1562); the built Unit casts once and is deleted when its recovery ends (Medium; R2-ENGINE-1740). It stops when the resolver is called with a blow or Blizzard value not below the building's health; no other route is shown to stop it (R2-ENGINE-1563). The Delivery 2 spell rows and Effect Speeds are in R2-ASSET-333.

## Computer-unit gaps

A struck ROM2 unit rewrites six command fields and posts a 0x54-byte record to the session list. The look-around is the one reader of command +0x54 found in the census (Medium), and no reader of a record was found in the chain that destroys the list (Medium) (R2-ENGINE-1491, R2-ENGINE-1492, R2-ENGINE-1262, R2-ENGINE-1451). Guard B reads the hit stamps in session ticks and in blocks of 16 ticks, and Guard A reads none of the hit fields (R2-ENGINE-1493, R2-ENGINE-1494). No body read chooses a point for an area spell: the aim is the victim unit or the victim's cell (R2-ENGINE-1490). Two scorers rank candidates, and of the effect words they read only bit 0x40000 of +0x144; an effect that changes the Mind word, the range or the virtual kind is not excluded (Medium) (R2-ENGINE-1495, R2-ENGINE-1496). The group assignment targets every member of the group with no idle test (R2-ENGINE-1497). A flee without Teleport walks, and an empty route ends in a stop and order 0xc at the unit's cell (R2-ENGINE-1498). A withdraw message sets v = 0, 10 or 30 for the owner object that message word +5 resolves to (that the word names a sender is Medium), and the outer message type, the sender and the control are Unknown (R2-ENGINE-1499, R2-ENGINE-1421). The numbers are in the area hit numbers table of the behavior page.

From the Shield arm of the support choice to the end of the recovery draw, a Shield cast at the caster draws once on the read path in RU, the recovery draw of the step that calls the apply (R2-ENGINE-1577, R2-ENGINE-1578, Medium; callers, list helpers, callees and virtual calls unread; the EN twins of `L2.00031`, `L2.02072`, `L2.08110` and `L2.02074` are unread). A Poison Cloud Effect applied again to a unit that carries it ends with remaining duration 128, the maximum of the standing value and 128, never a sum (R2-ENGINE-1575, R2-ENGINE-1576, Medium; RU).

## Strike refusal and threshold setup

The RU unit strike refuses a blow from a non-mage Human-kind attacker of attack type other than 5 against a movementType 3 target, and from any attacker for a missing party or owner, health at or below 0 or a range below the distance (R2-ENGINE-1605, Medium that a Human's attack type byte is the weapon's attackType; the EN strike body has no such test, EN behaviour Medium). In RU the Human threshold setter has one direct call site, in the template loader tail; a Unit takes nothing from it (R2-ENGINE-1606, Medium that a Human reaches that loader; other writers of the command fields stay Unknown, R2-ENGINE-1001).

## Routes of player orders

In the RU body, at Medium for play, guard walks back to its post (R2-ENGINE-1620, R2-ENGINE-1621, R2-ENGINE-1622). Patrol retries an unreachable point (R2-ENGINE-872, R2-ENGINE-1623, R2-ENGINE-1624). The pick-up cell test of the bodies read tests no route (R2-ENGINE-1183, R2-ENGINE-1625, R2-ENGINE-1626). The routes of standing orders, idle casting and look-around are not stated here (R2-ENGINE-1627, R2-ENGINE-1628, R2-ENGINE-1629, R2-ENGINE-1630). Area effects: R2-ENGINE-1537.
