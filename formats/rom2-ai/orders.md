# Planner and executor

## Order storage and planning

| Input/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Player setter | Reset, store unit order +0x50 and parameters immediately, then command +8=0 (two order-1 setters store 1). No direct primary +9 write in the read setters. | R2-ENGINE-541 High RU setters, Medium completeness | EN orders 1/3 setters; called effects and eventual execution |
| Unit-order selector | RU L2.01725 uses the 28-entry table L2.01729. Relevant routes: 1 walk, 3 attack, 0xb guard, 0xc idle, order 0xa (name not established), 0x16 flee, 0x18 planner L2.02864. | R2-ENGINE-555 High table, Medium per-order closure | Branch choices inside other planners; planner cadence |
| Attack target has +0x4c bit 8 | Skip slot draws; run L2.01733(unit,0), whose RU path refreshes/ranks candidates and chooses retarget or default. | R2-ENGINE-651 High local planner; R2-ENGINE-808 High RU stores, Medium composed target choice | Event meaning and unread candidate/mage helpers; not universal attack cancellation |
| Ordinary attack slots | Draw for each nonzero slot 0..2; the last draw below its signed threshold wins. Pick 23 casts only beyond 2 cells. Threshold is Probability * 0x147, draw bits 16..30 of LCG seed*0x343fd+0x269ec3. | R2-ENGINE-651 High draw/loop, Unknown caller rate; R2-ENGINE-652 High stores, Medium row binding | Initial seed and actual call frequency |
| Remaining attack route | With no target or self target take first unit from L2.02117; prepare +8=5, +0xc=target and +0x14=range, with computer-mage helper first. | R2-ENGINE-505 High RU body, Medium helper meanings | Mage helper and candidate eligibility; global activity-5 writer census |
| Idle target selection | Rank eligible candidates by nearest edge distance, least turn and later full tie. A found target prepares activity 6; no target takes support/default. | R2-ENGINE-501 High RU choice, Medium helper/composed no-step | EN activity-6 body and mage/support effects |

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
| 5, approach/attack | L2.02437 tests L2.01574 at command range +0x14; admitted: strike; otherwise L2.01577 and approach L2.01578. No distance cutoff appears in this activity body itself. | R2-ENGINE-505 High RU flow, Medium no-bound/helper meaning; R2-ENGINE-556 High acquired call | Search/step give-up and original trajectory |
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
