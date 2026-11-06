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

## Maps 21 and 31

Both maps write objectives 1 and 2 with instant 36 modes 1 and 2 only (no mode
4, so no state 5) and have no check 24 or check 23, so no condition in these two
maps reads an objective slot (R2-ENGINE-894, R2-ENGINE-895: High wiring; the
condition and end labels Medium). A condition row with a zero check ID is not compiled (R2-ENGINE-893);
the others compare result registers as R2-ENGINE-059 states.

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
