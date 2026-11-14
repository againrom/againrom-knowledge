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

## Static clause boundaries

The [planner page](orders.md) distinguishes local unit stores and incomplete
writer census (R2-ENGINE-1104, R2-ENGINE-1105), initiation versus HP-application
distance (R2-ENGINE-1106), and the paired setters with RU input chains
(R2-ENGINE-1107). Empty-route and arrival signals remain separate
(R2-ENGINE-1108). Hit-hook stores do not close scheduled effects
(R2-ENGINE-1109); Darkness slot arithmetic does not guarantee target retention
(R2-ENGINE-1110).
Light and Darkness lower the sight byte +0xa5 that the ray scans and the planner, movement
and group routines read: Darkness for a Human or Humanoid occupant, Light for a Unit occupant,
by 1+P/30 cells while the byte is at least that amount (below it the Effect arms wrap the byte
modulo 256 and no arm clamps; whether a reader clamps is not read except the ring walk's limit of
19 rings, R2-ENGINE-502), with no owner test (R2-ENGINE-1236, R2-ENGINE-1237, R2-ENGINE-1238,
R2-ENGINE-1239). The unit's own player receives the changed word (Medium; for Human and Humanoid
units with word +0xe in 0x21..0x3f another player may receive it, Unknown) (R2-ENGINE-1240), expiry
adds the shift back (R2-ENGINE-1241), and the per-class result is Medium (R2-ENGINE-1242)
(High for the arms, Medium for the reach, the delivery and the fog link).

Scorer L2.05238, candidate filtering and programme setters have separate
authorities: R2-ENGINE-1111, R2-ENGINE-1112 and R2-ENGINE-1113. The clock's local
rate selector is not a measured delivered cadence (R2-ENGINE-1114). Per-body
EN coverage, retained differences and ambiguous twins remain explicit
(R2-ENGINE-1115). The planner table names the read or run for each open cell.

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
at defended +0x1c0. The search takes the centre cell first, then square
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
of the cast routine, set the stamp. Effect tag 8 adds to +0xde without a
clamp; tag 11 adds to +0xe2 under an unread test (Medium) (R2-ENGINE-1206;
formulas High, sender meanings and writer completeness Medium).

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
tick is Unknown.
