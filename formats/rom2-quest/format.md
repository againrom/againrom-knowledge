# ROM2 quest objective machine

<!-- machine-page -->

An objective's state is one `scenario.dll` bank slot, 752 + n, written by script
instant 36 and shown as row n of the briefing window: a 24x24 mark from
`interface/subobj.256` beside the text `#subobjective<n>`. EN and RU clients
read equal. R2-ENGINE-046 (amended; its frame and banner clauses are partially
retracted), R2-ENGINE-888, R2-ENGINE-889, R2-ENGINE-890.

## States and transitions

States are numbered by their writers. No read body names a state.

| State | Written by | Row shown | Mark (frame of `subobj.256`) and text style | Claims and clause grades | Unknowns |
|---|---|---|---|---|---|
| 0 | EnterLocation clear of 752..767; NewGame bank clear | No; the row takes no height | none | R2-ENGINE-046 High (amended); R2-ENGINE-045 High; R2-ENGINE-888 High | none for the row |
| 1 | instant 36 mode 1 from state 0 only | Yes | frame 0 (value 0xa); style global three | R2-ENGINE-046 High (amended); R2-ENGINE-888 High; R2-ENGINE-889 High selection, Medium frame drawing and depiction | The style colours; what frame 0 depicts in play |
| 3 | instant 36 mode 2 from any state but 4; event 253 when bit 2 was clear | Yes | frame 1 (value 0xb); style global one | R2-ENGINE-046 High (amended); R2-ENGINE-889 High selection, Medium frame drawing and depiction; R2-ENGINE-891 High arm, Medium sound | The sound entry 0xe; the colours; sprite method +0x18 body (R2-ENGINE-889) |
| 5 | instant 36 mode 4 from any state; event 254 on every write | Yes | frame 2 (value 0xc); style global two | R2-ENGINE-046 High (amended); R2-ENGINE-889 High selection, Medium frame drawing and depiction; R2-ENGINE-891 High arm, Medium sound | The sound entry 0x10; the colours; sprite method +0x18 body (R2-ENGINE-889) |
| other m | instant 36 mode m (not 1, 2, 4) stores m | Yes if m is nonzero | by bit: 2 before 4 before 1; with none of them no mark and the previous row's style | R2-ENGINE-046 High (amended); R2-ENGINE-888 High | Whether any installed node uses such a mode is outside the maps read |

## Window and text

| Part | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Text cache | Mission entry and LOAD rebuild it from `#subobjective0` up to the first empty text; row i pairs with slot 752 + i. `#subobjective0` is line breaks only in the 34 text members per root that carry subobjective tags; in maps 21 and 31 the objectives that are written are 1 and 2 (R2-ENGINE-890, R2-ENGINE-894, R2-ENGINE-895) | R2-ENGINE-890 High loop, Medium row 0 caching; R2-ENGINE-049 Medium | The emptiness test body |
| Window | Briefing window class, title text 0x44 and the cached briefing; slot 30 builds the rows; one creator call per root | R2-ENGINE-892 High constructor, Medium creator census and base layout | The player action that opens it; the text above the rows |
| Rows | In cache order, nonzero rows only; mark at x 48..72, text from x 80 to width - 48, wrapped; 10 between rows; the window grows, no scroll; button 0x445 closes | R2-ENGINE-888 High, Medium line height | The line-height and wrap helpers |
| Notifications | 254 and 253 play a sound when idle; 250 plays one and posts text 0x61 to the message line; none writes state | R2-ENGINE-891 High arms, Medium sounds; R2-ENGINE-047 High | Which sounds entries 0xe and 0x10 are; the status helper's meaning and 30000 as milliseconds (R2-ENGINE-891) |
| Writers | Direct references in `scenario.dll`: the EnterLocation clear (indexed, covering 752..767) holds the only direct reference to the 752..767 range; no other direct reference to 753..767 | R2-ENGINE-896 High census, Medium completeness | Computed-index writers; bodies receiving the three bank-base pushes at `D2.00223`, `D2.00225`, `D2.00224` (R2-ENGINE-896) |

## SAV and LOAD

SAV carries the bank raw and LOAD restores it (R2-SESSION-031). LOAD rebuilds
the text cache (R2-ENGINE-890) and the read route does not pass EnterLocation,
so slots 752..767 survive a load (R2-SESSION-348, High for the save values,
Medium for the route; Unknown: the callers of the three EnterLocation call sites). The 16 original-made saves hold slot 754 only.

The selected EN/RU SAVE and LOAD bodies request a 4096-byte bank transfer, including the
objective cells (R2-SESSION-368, High for the requests). This static read did not execute stream
restoration or later callbacks; the first cold-LOAD town frame remains Unknown.

## Maps 21 and 31

Both maps write objectives 1 and 2 with instant 36 modes 1 and 2 only (no mode
4, so no state 5) and have no check 24 or check 23, so no condition in these two
maps reads an objective slot (R2-ENGINE-894, R2-ENGINE-895: High wiring; the
condition and end labels Medium). A condition row with a zero check ID is not compiled (R2-ENGINE-893);
the others compare result registers as R2-ENGINE-059 states. A nonzero check ID absent from
the map inserts a new map node when compiled (R2-ENGINE-1453, amended); the node's value is zeroed, so the ID names register 0, the first check's register (Medium; R2-ENGINE-1728).

| Map | Objective | Opens | Completes | Claims and clause grades | Unknowns |
|---|---|---|---|---|---|
| 21 | 1 | first runner pass | Get sack (60,18) reads 0 | R2-ENGINE-894 High wiring, Medium condition meaning | Check arm bodies; whether T4 can pass after T9 |
| 21 | 2 | a controlled unit (reference 10001) within the check-6 distance (constant 4) of unit 46 | same distance, sack nonzero, register 70 nonzero | R2-ENGINE-894 High wiring, Medium condition meaning | as above |
| 31 | 1 | within 4 of unit 101 while player 4 has units | within 4 of unit 101 when player 4 has none; the mission ends complete | R2-ENGINE-895 High wiring, Medium condition meaning and same-pass end | Check arm bodies; runner order |
| 31 | 2 | within 4 of unit 102 while player 4 has units | register 75 = 1, set near unit 100 after register 76 | R2-ENGINE-895 High wiring, Medium condition meaning | as above |

The Opens and Completes cells give the first test of the trigger only, other tests omitted: map 21 objective 1 T0 and T9, objective 2 T1 and T4 (R2-ENGINE-894); map 31 objective 1 T2 and T5, objective 2 T3 and T13 (R2-ENGINE-895). Map 31 objective 1 opens also with register 71 = 0 and completes also with register 74 = 0 (set by its own T5). All condition names are catalog labels. Map names: R2-ASSET-215. Grades belong to the cited clauses; a row is no
stronger than the lowest clause it names. The frame and banner clauses of
R2-ENGINE-046 are partially retracted; R2-ENGINE-889 and R2-ENGINE-891 replace
them.

Script operation reads add these bounds. The zero-check-ID rule covers the first left ID of a trigger as well: a zero first left ID excludes the whole trigger from compilation (High as code; R2-ENGINE-1316). No installed campaign trigger has one, and one campaign trigger (map 42, index 16, both roots) names left check ID 0xFFFFFFFF, which is no check ID of its map (R2-ASSET-248). Every installed script node has a nonempty name and ten nonempty text fields and every trigger a nonempty prefix (R2-ASSET-364); the readers found are two compile-failure messages, a check and an action, that receive the node pointer, and no reader of the text fields or the trigger prefix was found (Medium; R2-ENGINE-1681). In the 46 campaign maps no script operand is a zone record by the bodies read; the type-9 operand of check 21 selects the structure whose type-4 id is its value, and check 21 reads that structure's word +0x42, whose meaning is Unknown (Medium; R2-ENGINE-1685, amended by R2-ENGINE-1725). Check 2 is an inclusive rectangle test on a unit's x and y, and no installed map uses it (R2-ENGINE-1730, amending R2-ENGINE-1450 and R2-ENGINE-1685). Action 36's body calls GetVar/SetVar on 752+p0 with the arms of R2-ENGINE-046 (partially retracted, amended grade; High for the instructions, Medium for delegated effects; R2-ENGINE-1313). The read action wrappers expose operand widths and delegated calls only; they add no objective-state claim (R2-ENGINE-1314). Action 6's selector arms give group orders and programme values; the method bodies read write no objective slot (R2-ENGINE-1726, amending R2-ENGINE-1314). Actions 4 and 5 write the win counter and the failure reason that the result tick turns into victory or defeat (R2-ENGINE-1729, amending R2-ENGINE-1313). Checks 23 and 24 read scenario GetVar(p0) or GetVar(752+p0) with Medium helper results (R2-ENGINE-1315). The map-slot actions and check 19 have no local index guard, which proves no safe slot range (Medium; R2-ENGINE-1319).

Bodies of the numbers no installed map uses (except action 37, Unknown), the order of one script tick and the bodies of checks 1, 5, 8 and 20 are in R2-ENGINE-1450 (amended), R2-ENGINE-1451 (amended) and R2-ENGINE-1452 (amended), extended by R2-ENGINE-1696 (group list membership by unit state; check 20 counts Buildings rather than units), R2-ENGINE-1697 (check loop skips a check whose copied IR+0x58 byte is nonzero; that the 8000..9999 mark reaches that byte is Medium; Session init passes Session+0xd724 to a fill routine with value 0, EN High, RU Medium), R2-ENGINE-1699 (action 11 transfers out of U+0x7c into V+0x7c), R2-ENGINE-1700 (action 27 cell move and Guard setter) and R2-ENGINE-1701 (action 31 Outpost respawn countdown +0xb4).

The final effects of the other read actions narrow the wrapper boundaries of R2-ENGINE-1314 (amended). Action 21 casts spell low8(p4) once from cell (p0, p1) at cell (p2, p3) with power word p5, or 99 when p5 is 0, through a caster Unit that the world phase deletes when its recovery ends (Medium; R2-ENGINE-1740, R2-ENGINE-1741). Action 24 casts spell low8(p2) once from cell (p0, p1) at unit U with power word p3 (R2-ENGINE-1742). That the spell power equals the word up to 255 is Medium. Both roots ship five action 21 nodes in the scenario maps, four of them in triggers of maps 60, 71 and 95, and three action 24 nodes in map 96; four of those five triggers fire once (R2-ASSET-379). Actions 16 and 32 take a unit off the map and out of the unit registry (R2-ENGINE-1743); actions 17 and 33 put it back at the cell its stored position gives through unread getters (Medium) or within 3 cells, and action 18 puts V within 3 cells of U's cell after taking U off (R2-ENGINE-1744). Actions 19 and 22 give a unit to player P in a new group of its own (R2-ENGINE-1745). Action 23 adds p0 to player P's +0x3c, named gold (Medium) (R2-ENGINE-1746). Actions 12, 13 and 38 build or destroy an item in a bag (R2-ENGINE-1747). Action 20 drops a unit's bag as a sack (Medium); action 28 moves U's bag into V's unless V's count is 1024 or more (Medium, unread count helper), when U's items stay in its old bag, which nothing owns (Medium) (R2-ENGINE-1748). Actions 29, 30 and 34 set an area effect's remaining time, an effect's duration, or a unit's health, defence or absorption (Medium; R2-ENGINE-1749).

Action 25 writes one cell event on cell (low8 p2, low8 p3): event byte low8 p0, power low8 p1, and source and target pairs that are (p2, p3) twice when p4 is 0, else (p4, p5) and (p6, p7); it replaces an earlier event of that cell and creates the cell's record when absent, and the write is refused when bit 0 of the cell's byte at map +0x20000 is set (R2-ENGINE-1775, R2-ENGINE-1776, amending R2-ENGINE-1314 and R2-ENGINE-1730). With event 23 the cell moves units: a unit whose step ends at its centre is passed to the cell move, which puts it at the centre of the target cell unless the move test refuses it (R2-ENGINE-1778; Medium for units outside the three step callers; nothing observed in the original); the power byte and the source pair are unread for event 23 (Medium; R2-ENGINE-1779). No store to a written event was found in the EN range of R2-ENGINE-1777 (Medium; whole-entry stores and RU bodies other than the empty test not searched). Mission 53 writes such an event on cell (15, 121) with target (13, 22) from trigger 25, once check 38 on unit 89 and that cell is at most 4 (Medium; check helper unread; R2-ASSET-386); both roots ship nine action 25 nodes, all with event 23, in maps 53 and 93 to 96 (R2-ASSET-387).

The ROM2 computer-unit gaps read after these wrappers (area aim, hit orders, guard clocks, scores, group targets, flight, withdraw) are R2-ENGINE-1490 to R2-ENGINE-1499. The script tick destroys the list of hit records; no reader of a record was found (Medium) (R2-ENGINE-1451, R2-ENGINE-1262, R2-ENGINE-1491, R2-ENGINE-1492), and Guard A reads no hit field (R2-ENGINE-1494). A Delivery 2 cast adds its effect object to the effect list max(1, floor(distance / Effect Speed)) steps after the apply of R2-ENGINE-1262 (Medium cadence; the effect object's own step is Unknown) (R2-ENGINE-1567, R2-ENGINE-1568, R2-ENGINE-1569); within one step the executor runs before the substate arm, so a held attack in range hits every c + r + 2 steps (Medium; R2-ENGINE-1763); and a map building caster's cast builds a Unit that casts from state 13 (R2-ENGINE-1561, amended) once and is then deleted (R2-ENGINE-1740). That caster lists units within its record radius (R2-ENGINE-1560, Chebyshev (Medium)) and starts a cast once per 16 session ticks (R2-ENGINE-1562); the Delivery 2 spell rows are in R2-ASSET-333, which also holds the Units MP regeneration that sets a mage's mana rate when the reader stores it; 114 of 180 EN rows hold 0 and the effect of a 0 is Unknown (R2-ENGINE-1566).
Guard B reads the hit stamps in session ticks and blocks of 16 ticks (Medium: the tick object is not shown to be the global the guard reads; R2-ENGINE-1493), and of the effect words the two scorer bodies read only bit 0x40000 of +0x144, a candidate-score input (an effect that changes the Mind word, the range or the virtual kind is not excluded, Medium; R2-ENGINE-1496).

From the Shield arm of the support choice to the end of the recovery draw, a Shield cast at the caster draws once on the read path in RU, the recovery draw of the step that calls the apply (R2-ENGINE-1577, R2-ENGINE-1578, Medium; callers, list helpers, callees and virtual calls unread; the EN twins of `L2.00031`, `L2.02072`, `L2.08110` and `L2.02074` are unread). A Poison Cloud Effect applied again to a unit that carries it ends with remaining duration 128, the maximum of the standing value and 128, never a sum (R2-ENGINE-1575, R2-ENGINE-1576, Medium; RU).

## Threshold setup moment

In RU the Human threshold setter has one direct call site, in the template loader tail; a Unit takes nothing from it (R2-ENGINE-1606, amended by R2-ENGINE-1795: a Human built by `L2.00130` does not reach that loader and starts at 0, Medium; other writers of the command fields stay Unknown, R2-ENGINE-1001).

## Pocket flight, Teleport flight, staged areas, fallen units, hostile Heal, support relations, hired thresholds

In RU a refused flight in a pocket closed by terrain walks toward a cell at most 3 cells away for a unit inside the 8 .. size - 9 band; the route is not emptied when the route search's stand-in is another pocket cell, and then the order-1 branch is not entered and each Group pass stores the walk order again (R2-ENGINE-1787, amending R2-ENGINE-1498, R2-ENGINE-879 and R2-ENGINE-871; Medium); the own-cell stand-in case is open, and the outcome in play is Unknown. A computer mage that can pay Teleport casts it at a cell scanRange + 1 cells from its own cell, away from the hostiles' mean point along the major axis, clamped to the map band; the move happens when the landing test accepts the cell (Medium for the arguments; High for the arithmetic of `L2.02454`); with the hostile outside the flee list the idle planner gives prepared 6 on a hostile in range, else 0xb and, for a damaged mage with Heal, a self-Heal, and no walk (R2-ENGINE-1788, amending R2-ENGINE-1418 and R2-ENGINE-506; the caster's cell 40 ticks after the cast is Unknown).

Acid Stream and Blizzard accepted at one wind-up both attempt on area update 1; a first cell outside the band or rejected moves the first paint to update 4 or later (R2-ENGINE-1789, amending R2-ENGINE-812; Medium).

A fallen unit keeps running its Effect timers in each actor step before the health test; the fall sets the counter +0x6c to virtual +0x74 - 1 (the row value at index 33; for a Human the value at index 23 or the constant 8; counted in actor steps), and the death loop ends Poison with no further start (R2-ENGINE-1790, Medium; amending R2-ENGINE-1576 and R2-ENGINE-1534). From -3, a Poison Cloud of power P takes m = 2 + floor(2P/45) on each eighth timer call, and the unit dies on the first step after the counter at which health is -10 or below; the refill's loss below 0 comes on top (R2-ENGINE-1791, Medium):

| m (power P) | Start that first brings health from -3 to -10 or below, for word `+0xc6` = 0, ignoring the refill's loss (Medium) | Claim |
|---|---|---|
| 2 (P 0..22) | 4th | R2-ENGINE-1791 |
| 3 (P 23..44) | 3rd | R2-ENGINE-1791 |
| 4..6 (P 45..112) | 2nd | R2-ENGINE-1791 |
| 7 and above (P 113 and above) | 1st | R2-ENGINE-1791 |

The RU Heal arm refuses a target whose Player the caster's Player row marks hostile (bit 0), before its health tests; a refused Heal runs none of the arm's callees, while the apply head's cast award has run under its gates (bit 2, no item, mage, row parameter 2 nonzero) and the mana was paid at the cast (R2-ENGINE-1792, Medium for the argument roles; it amends R2-ENGINE-487 and the partially retracted R2-ENGINE-026).

The player support sweep heals only members of the builder's list +0xaa8 (the caster's bit 0 toward them clear) that pass the affinity; with Neutral alone the diagonal 0x12 refuses own units other than the caster, and a Player at relation 3 is not in the list the Heal walk reads (builder `L2.01515`) (R2-ENGINE-1793, amending R2-ENGINE-360 and R2-ENGINE-1200; Medium in play: the diagonal can be rewritten by script action 10, and `L2.01514` and `L2.01520` are unread). The relation matrix holds one byte per ordered pair of Player numbers at Session + 0xa8c4 + 70a + b; bit n is value 2^n:

| Bit or value | Use | Claim |
|---|---|---|
| bit 0 | the row's Player holds the column's hostile: hostility predicate, support builder, Heal arm | R2-ENGINE-1794 |
| bit 1 | ally: support filter, affinity Allies test | R2-ENGINE-1794 |
| low two bits 0 | affinity Neutral test | R2-ENGINE-1794 |
| bit 4 (0x10) | copied in reverse into the client word by the EN server routine `L2.01852` (EN read, RU not read) | R2-ENGINE-1794, R2-ENGINE-826 |
| 0x12 on the diagonal | map build and the original's saves | R2-ENGINE-1794 |
| (old & 0xfc) + low8(p2) | script action 10, which can store 3 | R2-ENGINE-1794 |

For a Human built by `L2.00130`, a hire or a map Human, no direct call reaches the loader or the threshold setter: the Unit constructor's default name has length 0 and the setter would read the Unit slot +0x30; command +0x40, +0x44 and +0x4c are 0 after creation, with a book or a bow (R2-ENGINE-1795, amending R2-ENGINE-1606, R2-ENGINE-506, R2-ENGINE-1001, R2-ENGINE-1002 and R2-ENGINE-777; Medium).
