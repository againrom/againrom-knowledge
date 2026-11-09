# ROM2 group machine, orders and routes

<!-- machine-page -->

This machine dispatches a Group programme, assigns member targets and prepares
actions. The unit executor and route search consume those actions. Programme,
unit order, prepared action and primary action are separate selectors.
R2-ENGINE-507, R2-ENGINE-541, R2-ENGINE-542, R2-ENGINE-555, R2-ENGINE-882.

## State and evaluation order

| Input/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Active Group | Session +0xbbe8 or programme active byte admits the Group tick; an empty Group gets programme 0xff. Dispatch precedes the health/hostile post-step. | R2-ENGINE-057 High for gates; R2-ENGINE-507 High RU; R2-ENGINE-882 High for both dispatch tables | Other session modes; seconds per Group tick |
| Guard A or B | Recompute member mean, gather sight, leash around that mean, assign targets, update radius, then decide each member. The radius update affects the next leash. | R2-ENGINE-841 High order, Medium helpers; R2-ENGINE-843 High writes | Full sight population and helper effects |
| Member action | Unit order +0x50 selects a planner; command +8 prepares an action, +9 selects the running primary action. Nonzero +9 takes precedence locally. | R2-ENGINE-541 High RU setters; R2-ENGINE-542 High executor; R2-ENGINE-555 High local table/stores, Medium closure | Planner/executor scheduling and end-to-end interruption |
| Same-owner area blow | An area pulse applies its object to each occupant and, outside rows 15 and 17, calls the hit hook (caster, occupant, 1) with no owner or relation test. When the owner's diagonal has bit 0 clear and low bits set (mission 10: 0x12), the hostility filter drops a same-owner caster and the read bodies keep the struck member's assigned target (Medium). | R2-ENGINE-985 High occupant body, Medium caster field set by the executor's cast; R2-ENGINE-989 Medium composition | Observed mission-10 case; posted-record readers; Darkness sight effect |
| Route failure | Failed full search selects a substitute and may stamp Guard B; an empty route sets movement +0x98 and the common tail replans by unit order. | R2-ENGINE-871 High branches; R2-ENGINE-872 High local writes; R2-ENGINE-873 High tests, Medium range interpretation | Composition Medium, R2-ENGINE-874; how often a failed search empties the list: R2-ENGINE-872 Unknown; grid identity, helper/list effects and observed trajectory |
| Saved continuation | SAV carries the 184-byte command and 80-byte Group programme plus lists, and unit +0x50/+0x54/+0x58 scalars. R2-ENGINE-823 is an EN transfer read; RU field offsets are placed in the same raw blocks by equal layout, not by a RU transfer read. | R2-ENGINE-823 Medium composed LOAD; R2-ENGINE-937 High EN postload keeps a zero +0xc, Medium writer completeness; R2-ENGINE-936 High EN fault of saved order 3 with command +0xc 0 in the session turn; R2-ENGINE-1012 High exception, Medium composed chain: separate EN fault in the LOAD-driven session update, on the unit send path | Complete key fixups, client restoration and later overwrites |

## Reference map

| Reference | Machine part |
|---|---|
| [Group and member decisions](behavior.md) | Tick, all programme values, Guard A/B, radius, struck-member memory and withdrawal |
| [Planner and executor](orders.md) | Order replacement, attack/idle planning, prepared and primary arms |
| [Search and empty routes](routes.md) | Full/repath failure, substitute, route construction, stamp and order-specific re-pick |
| [Fields and SAV coordinates](state.md) | Distinct bases, widths, stored coordinates and bounded save controls |

Grades belong to the cited clauses, not to this composition. A row's statement
is no stronger than the lowest clause named in it. A local High body does not
raise a composed trajectory above Medium. The former radius clause of
R2-ENGINE-507 is partially retracted; the current radius follows R2-ENGINE-843.
No runtime observation establishes the composed stamp cycle. R2-ENGINE-874,
R2-SESSION-343, R2-SESSION-346.

## Health thresholds, difficulty and attack order

Withdraw and Wimpy are Units columns (R2-ASSET-229). The stat reader stores them in the command's
+0x40 and +0x44 for a unit whose template loads (R2-ENGINE-1001; Medium: loader condition,
script-created creatures). The map arm, for a record without the person flag, halves HPmax, HP and
+0x44 when the difficulty word is 1 and Session +0x74 is 0; at any value other than 1 or 2 with
+0x74 = 0 it adds 50 to toHit and defence and multiplies HPmax by 1.5; at 2 it writes nothing
(R2-ENGINE-1002, RU; the table is read in RU, EN equal under the address mask, the EN doubles
unread). In the bodies read, the only direct caller of the threshold post-step is the group tick
(byte scan; indirect and jump routes unread, Medium; R2-ENGINE-1003). The order-3 arm passes command
+0xc to the planner with no null test (R2-ENGINE-1004). In the bodies read, LOAD leaves the
thresholds as saved (R2-ENGINE-1005, Medium).

## Route of a unit's next action

The RU read route from the Group pass to a unit's next action, with its standing-order and re-pick planners, is the composed route section of the [planner page](orders.md). Both owner +0x2c branches reach the selector through L2.01726 (R2-ENGINE-1017, Medium). The idle planner writes no unit order and no assigned target in its own body (R2-ENGINE-1018). The executor tail with movement +0x98 set stores order 0xc for a unit in order 1 (R2-ENGINE-1020).
The Group pass in this route runs only when session counter +4 mod 16 is 6 (R2-ENGINE-1016). The write census is by base class and does not count unit-order writers (R2-ENGINE-1022).
Programme byte writers, method L2.06060 and the saved values of players.0.group.0, read as the hero's Group (Medium): R2-ENGINE-1040 (High bodies, Medium completeness). Programme 3 route (no direct selector call) and the unknown reach of the selector for that Group: R2-ENGINE-1041. Walk prepare, the order-0xc setter and movement +0x98 sites: R2-ENGINE-1042. The found-target branch: R2-ENGINE-1043. Attack planner operands: R2-ENGINE-1044. The script switch L2.03275 and L2.05852: R2-ENGINE-1045. Each carries its grade and gaps on the [orders page](orders.md).

Counted run (EN, one mission-10 launch, one hero): R2-ENGINE-1072 (loop L2.04213 and wrapper L2.02403 had 0 hits in the span; route serving the hero Unknown in launch 1), R2-ENGINE-1073 (actor step once per scheduler counter value per unit; the counter advanced 16.1 and 32.1 per second in the two launches), R2-ENGINE-1074 (a ground click wrote order 1 and moved the hero to another Group object).

Continuation reads: R2-ENGINE-1075 (EN return addresses L2.06121 and L2.06122 of L2.04984 lie in the idle planner's no-candidate arm and in the programme-3 route; which of the planner's six E8 callers served the hero is Unknown in launch 1), R2-ENGINE-1076 (RU programme 4 and Guard A bodies, including Guard A's arm L2.01733(member, 1), which does not call the idle planner; the empty-Group store of programme 0xff), R2-ENGINE-1078 (a second EN launch: with no attack order the hero read order 0xc and a strike state on another owner's unit, and the idle planner's entries by return address), R2-ENGINE-1077 (256-instruction masked comparison of RU and EN bodies; 16 pairs equal, seven differ at a later index, the EN twin of L2.05852 stays ambiguous).
