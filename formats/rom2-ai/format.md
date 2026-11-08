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
| Saved continuation | SAV carries the 184-byte command and 80-byte Group programme plus lists, and unit +0x50/+0x54/+0x58 scalars. R2-ENGINE-823 is an EN transfer read; RU field offsets are placed in the same raw blocks by equal layout, not by a RU transfer read. | R2-ENGINE-823 Medium composed LOAD; R2-ENGINE-937 High EN postload keeps a zero +0xc, Medium writer completeness; R2-ENGINE-936 High EN fault of saved order 3 with command +0xc 0 in the session turn | Complete key fixups, client restoration and later overwrites |

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
