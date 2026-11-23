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
the map inserts a new map node when compiled; the slot it then names is Unknown (R2-ENGINE-1453).

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

Script operation reads add these bounds. The zero-check-ID rule covers the first left ID of a trigger as well: a zero first left ID excludes the whole trigger from compilation (High as code; R2-ENGINE-1316). No installed campaign trigger has one, and one campaign trigger (map 42, index 16, both roots) names left check ID 0xFFFFFFFF, which is no check ID of its map (R2-ASSET-248). Every installed script node has a nonempty name and ten nonempty text fields and every trigger a nonempty prefix (R2-ASSET-364); the readers found are two compile-failure messages, a check and an action, that receive the node pointer, and no reader of the text fields or the trigger prefix was found (Medium; R2-ENGINE-1681). In the 46 campaign maps no script operand is a zone record by the bodies read; the record that the type-9 operand of check 21 selects is unread (Medium; R2-ENGINE-1685). Action 36's body calls GetVar/SetVar on 752+p0 with the arms of R2-ENGINE-046 (partially retracted, amended grade; High for the instructions, Medium for delegated effects; R2-ENGINE-1313). The read action wrappers expose operand widths and delegated calls only; they add no objective-state claim (R2-ENGINE-1314). Checks 23 and 24 read scenario GetVar(p0) or GetVar(752+p0) with Medium helper results (R2-ENGINE-1315). The map-slot actions and check 19 have no local index guard, which proves no safe slot range (Medium; R2-ENGINE-1319).

Bodies of the numbers no installed map uses (except action 37, Unknown), the order of one script tick and the bodies of checks 1, 5, 8 and 20 are in R2-ENGINE-1450, R2-ENGINE-1451 and R2-ENGINE-1452.
The ROM2 computer-unit gaps read after these wrappers (area aim, hit orders, guard clocks, scores, group targets, flight, withdraw) are R2-ENGINE-1490 to R2-ENGINE-1499. The script tick destroys the list of hit records; no reader of a record was found (Medium) (R2-ENGINE-1451, R2-ENGINE-1262, R2-ENGINE-1491, R2-ENGINE-1492), and Guard A reads no hit field (R2-ENGINE-1494). A Delivery 2 cast adds its effect object to the effect list max(1, floor(distance / Effect Speed)) steps after the apply of R2-ENGINE-1262 (Medium cadence; the effect object's own step is Unknown) (R2-ENGINE-1567, R2-ENGINE-1568, R2-ENGINE-1569), and a map building caster's cast builds a Unit that casts from state 13 (R2-ENGINE-1561). That caster lists units within its record radius (R2-ENGINE-1560, Chebyshev (Medium)) and starts a cast once per 16 session ticks (R2-ENGINE-1562); the Delivery 2 spell rows are in R2-ASSET-333, which also holds the Units MP regeneration that sets a mage's mana rate when the reader stores it; 114 of 180 EN rows hold 0 and the effect of a 0 is Unknown (R2-ENGINE-1566).
Guard B reads the hit stamps in session ticks and blocks of 16 ticks (Medium: the tick object is not shown to be the global the guard reads; R2-ENGINE-1493), and of the effect words the two scorer bodies read only bit 0x40000 of +0x144, a candidate-score input (an effect that changes the Mind word, the range or the virtual kind is not excluded, Medium; R2-ENGINE-1496).

From the Shield arm of the support choice to the end of the recovery draw, a Shield cast at the caster draws once on the read path in RU, the recovery draw of the step that calls the apply (R2-ENGINE-1577, R2-ENGINE-1578, Medium; callers, list helpers, callees and virtual calls unread; the EN twins of `L2.00031`, `L2.02072`, `L2.08110` and `L2.02074` are unread). A Poison Cloud Effect applied again to a unit that carries it ends with remaining duration 128, the maximum of the standing value and 128, never a sum (R2-ENGINE-1575, R2-ENGINE-1576, Medium; RU).

## Threshold setup moment

In RU the Human threshold setter has one direct call site, in the template loader tail; a Unit takes nothing from it (R2-ENGINE-1606, Medium that a Human reaches that loader; other writers of the command fields stay Unknown, R2-ENGINE-1001).
