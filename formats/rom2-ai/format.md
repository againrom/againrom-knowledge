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
| Route failure | Failed full search selects a substitute and may stamp Guard B; an empty route sets movement +0x98 and the common tail replans by unit order. | R2-ENGINE-871 High branches; R2-ENGINE-872 High local writes; R2-ENGINE-873 High tests, Medium range interpretation | Composition Medium, R2-ENGINE-874; how often a failed search empties the list: R2-ENGINE-872 Unknown; grid identity, helper/list effects and observed trajectory |
| Saved continuation | SAV carries the 184-byte command and 80-byte Group programme plus lists, and unit +0x50/+0x54/+0x58 scalars. R2-ENGINE-823 is an EN transfer read; RU field offsets are placed in the same raw blocks by equal layout, not by a RU transfer read. | R2-ENGINE-823 Medium composed LOAD | Complete key fixups, client restoration and later overwrites |

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
