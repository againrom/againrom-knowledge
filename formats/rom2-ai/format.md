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

Area reach of these Effects (Medium throughout): Light and Darkness pulse onto cell entry +4 only, so kind 1 or 2 units
(Units rows with movementType 1 or 2, which include Ghost and Bee; R2-ASSET-327, R2-ASSET-328, R2-ENGINE-1530) are
reached and kind 3 units (Bat_Sonic, Dragon, Succubus; a flier by name only) and buildings are not
(R2-ENGINE-1535); in EN sight is base once per 16 updates between the end of one pulse's Effect
and the next pulse (R2-ENGINE-1534). Blizzard passes the kind 1 or 2, kind 3 and building occupants of
its one sampled cell per pulse to a damage object that admits all three (R2-ENGINE-1533; cell validity band `8..W - 9`, `8..H - 9` confirmed High in both roots, R2-ENGINE-1668; resolver rejection by
target flag +0x4c bit 8 is unread in EN); the cells and pulses of the other area spells
are in R2-ENGINE-1536; Fire Ball hits a multi-cell unit once per covered cell (R2-ENGINE-1537, Medium).

Scorer L2.05238, candidate filtering and programme setters have separate
authorities: R2-ENGINE-1111, R2-ENGINE-1112 and R2-ENGINE-1113. The clock's local
rate selector is not a measured delivered cadence (R2-ENGINE-1114). Per-body
EN coverage, retained differences and ambiguous twins remain explicit
(R2-ENGINE-1115). The planner table names the read or run for each open cell.

The orders page's static continuation adds bounded clauses. The selector calls
the stop chain for order 3 when the target's +0x54 is 0x10 (R2-ENGINE-1121; the
order-0xc ending of the chain is R2-ENGINE-1020; the meaning of 0x10 is Medium; orders 8 and 0x11 are tested alike and the cast orders 0xd and 0xe are not, R2-ENGINE-1548; the immediate stores of 0x10, mov with displacement 0x54, are two per root and on the death path, register and other store forms unread, R2-ENGINE-1547, amending R2-ENGINE-1121)
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
and its cast pays no mana (R2-ENGINE-1441, Medium). A weapon's shop price adds its cast spell apart; an
armour's or shield's spell enters the effect size S of the surcharge trunc((1.5^(S/70) + 1) x S x 50)
(R2-ENGINE-1824, High for the order; Medium for the price).

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

Prismatic Spray takes its extra victims from Group sight of the caster's
`+0x70` members. The candidates are hostiles (bodies only when no living hostile is in sight; Medium), ranked by
edge distance from the caster, then by the turn from the caster's facing, and
admitted up to the spell's range byte. The target comes first and the list is
cut to the cap (R2-ENGINE-1766, High / Medium). Drain Life stores the caster
as the victim's killer and stores its spell id. A unit at health 0 or below
falls and dies at -10 or below, and the kill route may credit the stored killer
(R2-ENGINE-1767, High / Medium; the credit is Medium). In RU a mage's spell hit stores its Sphere in victim
byte +0x48, which the kill reads as a spell id; an RU spell kill offers its experience to award index 1 for
Spheres 1 to 4 and index 2 for Sphere 5 (R2-ENGINE-1826, Medium for the credit in play). A non-mage's weapon
Spell fires only when the blow's damage and the target's health after it are above 0, except Spell id 2, which
fires on every strike that reaches the cast test (R2-ENGINE-1829, High / Medium).

A unit with +0x54 = 0x10 or with health 0 or below restores nothing; below 0
the refill takes 1 health on one pass in four, and corpse ageing lowers
health separately, so the total loss per pass is not this figure alone.
Corpse ageing takes 1 health on every second world pass (Medium for the Session counter identity; a third caller ages units outside the world pass); a corpse of a class other than 0x21..0x24 reaches
kind 5 below -600 and is released (R2-ENGINE-1765, High / Medium). The stage tests are strict: kind 2, 3, 4 and
5 below -10, -20, -40 and -600. In RU a corpse of class 0x21..0x24 is held at -61 with kind 4 and never gets kind 5
by ageing; in EN it stops losing health below -60 but gets kind 5 when its health is below -600 at an ageing call
(R2-ENGINE-1825, High / Medium). The
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

## Guard radius source

The installed Scanning/MinimalGuardRange value is 8 in both roots; the getter
fallback is 10. Guard A's zero-argument floor copies that value; Guard B's
zero-argument branch writes 20 when setup's radius is below it. These writes
are on the group controller. One raw Guard B caller in RU (L2.07155) and two in EN (L2.07156,
L2.07157) are not interpreted, the complete
source-loader chain and the other writer forms remain open. R2-ENGINE-1252
(High data and lookup bodies; Medium delivery, caller and writer completeness).

## Conditional guard distances

With the installed floor delivered as 8 and r=0, Guard A starts at max(b,8),
where b is setup's +0x2c byte. Guard B keeps b when b>=8 and otherwise writes
20. Under the Medium draw range -1..1, baselines 8 and 20 become 11..13 and
23..25 after the first-member latch, affecting the next leash. The no-+4 helper
can write 7..9 and 19..21 through R2-ENGINE-843's no-member arm (reachability
Unknown); the working radius on that path is not bounded below by the setup floor. A nonzero r,
LOAD, another writer or byte overflow prevents a universal minimum claim.
The same radius filters the listed enemy before engage and pursuit; it does
not establish sight detection distance. R2-ENGINE-1253 (High local stores and
conditional arithmetic; Medium delivered bounds; global minimum Unknown).

## Pursuit reference and return

The leash follows the current truncated member mean, so it does not give a
fixed maximum journey from the post. A listed enemy at distance exactly equal
to the radius survives that filter; other sight/score exclusions can still
clear assignment. Target-free Guard A prepares the post return; Guard B's
HP/time and case-5 goals can precede its post return. Visible turn timing and
arrival stay Medium through the planners and active action. R2-ENGINE-1254
(Medium composition in both roots; no new target-writer census).

## Summoned-unit count pass

The world pass counts nonzero caster keys and compares a 1..100 draw with
n, then kills the first linked-list unit of each marked key, at most one per
key per pass. For n=1 that test cannot fire; for n=2 its conditional chance
is 328/32768. It runs at Session +4 mod 16 = 12, with delivered time still
Medium. The read Summon arm supplies no admission-count test; a universal
absence of a cap is unproved. Count-map helpers, list population, death
callees and visible departure keep the composition Medium. R2-ENGINE-1255
(High new helper bodies and arithmetic, Medium composition, global cap Unknown).

## Departure and LOAD boundary for summoned units

The bounded departure census retains the count-pass death and the selected
mission-transition detach/destroy path. An independent timer, caster-dependent
cleanup and general defeat fate remain Unknown. The Unit serializer omits
+0x148 at its cited grade; this alone proves neither disappearance nor immunity
after LOAD. Field candidates, indirect routes and saved-unit selection remain
unclosed. R2-ENGINE-1256 (Medium bounded composition, completeness and runtime
outcome Unknown).

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

The executor runs before the actor substate arm in each step. A strike started at step s counts down from s + 1 and
hits at s + c when it does not convert to state 13 (Medium), c being byte +0x134 plus a distance term; a held attack in range hits every c + r + 2 steps (Medium)
(R2-ENGINE-1763: High bodies, Medium time and period).

Across a mission LOAD the actor-step fields +0x58, +0x6c and +0x134..+0x136 are read back into place, so a caster saved in wind-up or recovery resumes its own count (R2-ENGINE-1401: High for the RU body, Medium for EN; that the count runs to its end is Medium, postload callees unread) and begins its next order no earlier than R2-ENGINE-1262's s + c + r (Medium). A loaded timed area effect with a nonzero stored +0x4c continues from the stored +0x50 (R2-ENGINE-1400: High for a stored +0x4c nonzero, Unknown for +0x4c 0), and the read effect serializers store no caster field (R2-ENGINE-1398, Medium for the values after construction; what the effect +0x14 owner key names is Unknown), so a loaded effect reaches its target with no stored caster link (R2-ENGINE-1399, Medium); the HP amount without a caster is Unknown.

The ordinary script action and check wrappers are read independently of this page: local stores and calls are High, delegated effects Medium, and the bodies of actions 33 and 34 are read (R2-ENGINE-1450), their delegated effects Medium. The 19 ordinary check bodies are read the same way, with Medium helper results (R2-ENGINE-1315), and the map-slot actions and check 19 have no local index guard, which proves no safe slot range (Medium; R2-ENGINE-1319). They change no rule above (R2-ENGINE-1313, R2-ENGINE-1314).

The final effects of the other read actions (R2-ENGINE-1314 amended): actions 21 and 24 build the caster Unit of a map building caster, which has no mage bit, casts once and is deleted (Medium) by the world phase when its recovery ends (R2-ENGINE-1740). Action 21 casts spell low8(p4) from cell (p0, p1) at cell (p2, p3) with power word p5, or 99 when p5 is 0 (R2-ENGINE-1741); action 24 casts spell low8(p2) from cell (p0, p1) at unit U with power word p3 (R2-ENGINE-1742); that the spell power equals the word up to 255 is Medium. Both roots ship five action 21 nodes and three action 24 nodes in the scenario maps (R2-ASSET-379). Actions 16 and 32 take a unit off the map and out of the unit registry (R2-ENGINE-1743); actions 17 and 33 put it back at the cell its stored position gives through unread getters (Medium) or within 3 cells, and action 18 takes U off and puts V within 3 cells of U's cell (R2-ENGINE-1744). Actions 19 and 22 give a unit to player P, each in a new group of its own (R2-ENGINE-1745). Action 23 adds p0 to player +0x3c, named gold (Medium) (R2-ENGINE-1746). Actions 12, 13 and 38 build or destroy an item in a bag (R2-ENGINE-1747). Action 20 drops a unit's bag as a sack at its cell (Medium); action 28 moves U's bag into V's, and when V's count is 1024 or more (Medium, unread count helper) U's items stay in its old bag, which nothing owns (Medium) (R2-ENGINE-1748). Actions 29, 30 and 34 set an area effect's remaining time, an effect's duration, or a unit's health, defence or absorption (Medium; R2-ENGINE-1749).

The bodies of the numbers no installed map uses are read in R2-ENGINE-1450 (amended), except action 37, whose cell is a message-formatting arm (Unknown): action 11 transfers the item out of U+0x7c into V+0x7c (R2-ENGINE-1699), action 27 runs terrain cell move and Guard setter for nonzero owner+0x2c (R2-ENGINE-1700), and action 31 writes Outpost remaining respawn countdown +0xb4 matching group number +0xac (R2-ENGINE-1701, amended; p1 goes into +0xb4 of every Outpost whose group number equals p0, for every owner and with no loop exit; the filter passes only the Outpost in the classes with a record, EN Shop and Inn slot 14 unread, Medium: R2-ENGINE-1820; p1 counts container passes without the respawn scale, Medium: R2-ENGINE-1821). Unconscious units stay in their group list while dead units (+0x54 = 0x10, body or no body) are detached by the kill walker; checks 1, 5 and 8 count every unit still in its group list (living and unconscious units, not dead units; EN High, RU Medium for checks 1 and 8), and check 20 counts Buildings with i16 +0x42 > 0 rather than units (R2-ENGINE-1452 amended, R2-ENGINE-1696). The script tick runs before the loop over the group list, skips checks whose copied IR+0x58 byte is nonzero (that the 8000..9999 mark reaches that byte is Medium), and empties the Session+0xbbec list (R2-ENGINE-1451 amended, R2-ENGINE-1697).

Script action 6 gives a group one of eleven method calls by its selector p0 (1, 2, 3, 4, 5, 10, 11, 14, 15, 17, 18); the method bodies are read in EN and RU, the other selectors do nothing, and the arms set member orders 0xb, 0xc, 1, 3, 8, 0xa and 0x11 and the group programme value; what these orders and programme values do in play is Medium (R2-ENGINE-1726, amending R2-ENGINE-1314).

The roles of the Session DWORDs that script actions 4 and 5 write are the win counter (+0xbe0c) and the failure reason (+0xbe14) that the result tick reads, sent as victory when the counter is nonzero and as defeat when the reason is above 0 (Medium as the player sees them; R2-ENGINE-1729, amending R2-ENGINE-1313).

Script check 2 reads x and y as bytes +0 and +1 of the object at unit +0x10 and is an inclusive rectangle test on them (High); checks 3 and 27 read the same getters (Medium) (R2-ENGINE-1730, amending R2-ENGINE-1450).

Script action 25 writes one cell event on the single cell (low8 p2, low8 p3), not a rectangle: event byte low8 p0, power low8 p1, and source and target pairs that are (p2, p3) twice when p4 is 0, else (p4, p5) and (p6, p7); the write replaces an earlier event of that cell (R2-ENGINE-1775, R2-ENGINE-1776, amending R2-ENGINE-1730 and R2-ENGINE-1314). With event 23, a unit whose step ends at the centre of that cell is passed to the cell move, which puts it at the centre of the target cell unless the move test refuses it (Medium for units outside the three step callers; nothing observed in the original), and occupant kinds 1 and 2 also keep the target in command word +0xb0 for one retry by the executor (R2-ENGINE-1778). The power byte and the source pair are unread for event 23 (Medium; R2-ENGINE-1779), and no store to a written event was found in the EN range of R2-ENGINE-1777 (Medium; whole-entry stores and RU bodies other than the empty test not searched). Mission 53 writes such an event from trigger 25 (R2-ASSET-386); both roots ship nine action 25 nodes, all with event 23 (R2-ASSET-387).

EN `L2.06918` stores the speed, heading, the per-advance x and y bytes and the advance count into the block at unit `+0x1c0` (R2-ENGINE-1470). In EN `L2.03582..L2.06966` it has one call (the step start of R2-ENGINE-1213); that it is the only writer of the step values is Medium. The count is 256 divided by the larger axis byte rounded up (Medium: the absolute-value helper `L2.04968` is unread). For a unit whose kind (`L2.03577`) is 1 the speed comes from the base speed, the multiplier, the slope between the two cells (Medium: the byte read as a height, R2-ENGINE-1340) and the mean of two cell bytes (the sum kept in a byte); for any other kind it is the base speed. The result is clamped to 1..63 (R2-ENGINE-1471); no group term was read. The 18 path and terrain parameters are read from `World\Data\map.reg` with code defaults (R2-ENGINE-1472; tile classifier `L2.06958`/`L2.07473` indexes `+0x54146..+0x54150` and writes no height, R2-ENGINE-1665; client/server type-3 readers in R2-ENGINE-1667); the full tables are on the route page.

A step changes the cell on advance ceil(128/v) when its y vector component (x when it has none) is positive and floor(128/-v)+1 when negative, from fine bytes 0x80 and a heading that already matches, advance 1 being the step start tick in that case (R2-ENGINE-1760, Medium for the premises). The static route is rebuilt on a new goal or after 17 repaths, the local path when empty or after 33 walk calls, unless the wait decision of R2-ENGINE-1762 returns 1 or 2, and after a re-plan (R2-ENGINE-1761), with the shipped `map.reg` values of R2-ASSET-383. In RU the local search stores every cell from the unit's next cell to the local goal, none when its trace passes 1000 steps; a step takes the first cell, and the list is cleared at each cell centre; the search then runs again at the next walk call unless the wait decision holds the unit (R2-ENGINE-1827, Medium). At a held next cell a one-cell unit turns and waits for a walking occupant or a standing one about to leave, else it re-plans (R2-ENGINE-1762). In RU, in the bodies read, the Human speed word `+0x8c` that the step speed falls back on has a floor of 6 only when load reaches capacity and no upper bound, and an effect modifier may leave it at 0 or below; when the step takes that word it clamps it to 1..63 (R2-ENGINE-1830, Medium for the step).

A map building caster lists the units within its record radius of its cell, Chebyshev (Medium), 1 to 9 cells, with no range compare against the spell (R2-ENGINE-1560). It starts a cast once per 16 session ticks through a built Unit that casts from state 13 with wind-up 1 and recovery 1 (R2-ENGINE-1561 amended, R2-ENGINE-1562); the built Unit casts once and is deleted when its recovery ends (Medium; R2-ENGINE-1740). It stops when the resolver is called with a blow or Blizzard value not below the building's health; no other route is shown to stop it (R2-ENGINE-1563). The Delivery 2 spell rows and Effect Speeds are in R2-ASSET-333.

## Computer-unit gaps

A struck ROM2 unit rewrites six command fields and posts a 0x54-byte record to the session list. The look-around is the one reader of command +0x54 found in the census (Medium), and no reader of a record was found in the chain that destroys the list (Medium) (R2-ENGINE-1491, R2-ENGINE-1492, R2-ENGINE-1262, R2-ENGINE-1451). Guard B reads the hit stamps in session ticks and in blocks of 16 ticks, and Guard A reads none of the hit fields (R2-ENGINE-1493, R2-ENGINE-1494). No body read chooses a point for an area spell: the aim is the victim unit or the victim's cell (R2-ENGINE-1490). Two scorers rank candidates, and of the effect words they read only bit 0x40000 of +0x144; an effect that changes the Mind word, the range or the virtual kind is not excluded (Medium) (R2-ENGINE-1495, R2-ENGINE-1496). The group assignment targets every member of the group with no idle test (R2-ENGINE-1497). A flee without Teleport walks, and an empty route ends in a stop and order 0xc at the unit's cell (R2-ENGINE-1498). A withdraw message sets v = 0, 10 or 30 for the owner object that message word +5 resolves to (that the word names a sender is Medium), and the outer message type, the sender and the control are Unknown (R2-ENGINE-1499, R2-ENGINE-1421). The numbers are in the area hit numbers table of the behavior page.

From the Shield arm of the support choice to the end of the recovery draw, a Shield cast at the caster draws once on the read path in RU, the recovery draw of the step that calls the apply (R2-ENGINE-1577, R2-ENGINE-1578, Medium; callers, list helpers, callees and virtual calls unread; the EN twins of `L2.00031`, `L2.02072`, `L2.08110` and `L2.02074` are unread). A Poison Cloud Effect applied again to a unit that carries it ends with remaining duration 128, the maximum of the standing value and 128, never a sum (R2-ENGINE-1575, R2-ENGINE-1576, Medium; RU).

## Strike refusal and threshold setup

The RU unit strike refuses a blow from a non-mage Human-kind attacker of attack type other than 5 against a movementType 3 target, and from any attacker for a missing party or owner, health at or below 0 or a range below the distance (R2-ENGINE-1605, Medium that a Human's attack type byte is the weapon's attackType; the EN strike body has no such test, EN behaviour Medium). In RU the Human threshold setter has one direct call site, in the template loader tail; a Unit takes nothing from it (R2-ENGINE-1606, amended by R2-ENGINE-1795: a Human built by `L2.00130` does not reach that loader and starts at 0, Medium; other writers of the command fields stay Unknown, R2-ENGINE-1001).

## Player order fields

In the RU body, at Medium for play, the order fields read by guard, patrol, pick-up and standing orders are in R2-ENGINE-872, R2-ENGINE-1183, R2-ENGINE-1620, R2-ENGINE-1621, R2-ENGINE-1622, R2-ENGINE-1623, R2-ENGINE-1624, R2-ENGINE-1625, R2-ENGINE-1626, R2-ENGINE-1627, R2-ENGINE-1628 and R2-ENGINE-1629; the per-cell effect list is in R2-ENGINE-1537.

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

## Cell planes in the route search

The route search (RU `L2.05174`, EN `L2.05175`) floods its word grid at +0x30000 + 2 × key from the unit's cell. With flag 1 (walk and approach) it tests plane +0x10000, with flag 0 (repath) plane +0x20000, and it skips a neighbour whose plane byte AND the unit's mask byte is nonzero; for unit kind 1 a step costs the terrain byte straight and one and a half times it diagonal, for other kinds 2 and 3 (R2-ENGINE-1804; the flag roles are Medium, R2-ENGINE-871). In the size-1 arm the loop tests the goal word, the frontier and the half-pass limit once per pair of one list-A and one list-B half-pass (R2-ENGINE-1876). The grid helpers are read: `L2.05185`(x, y) and `L2.05180`(key) address the same grid word for one `this`, `L2.05230` the terrain byte and `L2.02402` returns unit virtual +0x20 (R2-ENGINE-1804, amending R2-ENGINE-871, R2-ENGINE-878 and R2-ENGINE-879). With the plane values the terrain build stores (R2-ENGINE-1340), for flag 1 and a unit of size 1, value 1 closes kind 1, value 5 kinds 1 and 2, and the border ring's 0x1f all three kinds; in the route-search workers and helpers read, a terrain byte of 255 enters only kind 1's step sum and no test closes a cell on it (R2-ENGINE-1805, Medium; the outcome in play is Unknown). For a larger unit the footprint loop is RU `L2.08532` (EN `L2.08533`), which calls the AND of two bytes RU `L2.03520` (EN `L2.02720`) (R2-ENGINE-1804, narrowing the footprint tester naming of R2-ENGINE-1667). The EN location routine `L2.07308` keeps `(y << 8) + x`, which has the form of the plane key (R2-ENGINE-1802; EN body read, Medium that the word is the plane key, RU body unread).

The RU reads of the route search loop, the group target scan, the defender list, the hostile filter and the Teleport admission, and the older cards they amend, are cited here: R2-ENGINE-502, R2-ENGINE-1111, R2-ENGINE-1112, R2-ENGINE-1185, R2-ENGINE-1188, R2-ENGINE-1418, R2-ENGINE-1440, R2-ENGINE-1804, R2-ENGINE-1877, R2-ENGINE-1878, R2-ENGINE-1879, R2-ENGINE-1880.

In EN, Drain Life with a computed v of 0 or below skips the hit hook and the killer store (R2-ENGINE-1892, High EN arm, Medium RU and callers); in EN, action 6 selectors 1 and 18 store order 0xb at the member's cell and a circle record whose radius is raised to p1 when p1 is nonzero (selector 1 else session byte +0xa8b8, selector 18 else 0x14), and no member flag in the bodies read (R2-ENGINE-1893, High EN, Medium RU and unread helpers); the flee Teleport admission is R2-ENGINE-1880.

The group programme byte stores are 17 byte sites per image with no immediate from 6 to 16 (High as a count, Medium as a complete cover; copies and loads of a group record are not covered; R2-ENGINE-1894); no body with the shape of a Druid patrol stage was found among the stores, the dispatch and six order handlers read (Medium; R2-ENGINE-1895); EN L2.08843 and L2.08844 are the only coordinate walkers the order dispatcher reaches, and EN L2.08845, with thresholds on x and a destination, has no reference in either image (High for the bodies, Medium for meaning, reachability Unknown; R2-ENGINE-1896); the position getter census is Medium and covers direct calls only (R2-ENGINE-1897); in EN a Darkness caster can be a hook target when it is the occupant picked (Medium; R2-ENGINE-1898).
