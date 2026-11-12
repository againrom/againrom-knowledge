# Fields and SAV coordinates

## Relation coordinates

The hit hook's relation-helper boundary is closed separately from its
memory stores: the matrix is Session +0xa8c4+70*source+target, one byte per
directed cell. The complete relation helper gates each direction by that
cell's low two bits; its nonzero-third-argument reverse gate reads and can
increment victim command DWORD +0xac. These are native state coordinates,
not a new SAV layout proof. Hit-memory and clock/geometry grades stand.
— R2-ENGINE-845, R2-ENGINE-981

One owner index i makes both helper tests read the diagonal M[i][i]. The
hook also stores session +0xbd98=1, +0xbdc4=attacker and +0xbdc8=victim and
appends a 0x54-byte copy of +0xbd98 to session list +0xbbec. An area object
keeps its caster at +0x3c, copied to its attached object +0x44; an occupant
gets caster pointer +0x40 and byte +0x48 from the attached object +0xc.
Mission 10 stores owner word 5 for Troll.4 and the four Bat_Sonic.4 and
diagonal 0x12 in Player row 4, so their composed same-owner hit leaves
M[5][5] unchanged (Medium). — R2-ENGINE-985, R2-ENGINE-987, R2-ENGINE-988,
R2-ENGINE-989, R2-ENGINE-990

The listed hover Player's relation words use array +0x38 and data +0x3c, indexed
by the hovered owner's key. The read hostility helper returns word &1.
The server's querying-unit/candidate predicate reads the corresponding
outgoing matrix cell. Complete acquisition, delivery and attack admission
remain Medium. — R2-ENGINE-982

Offsets below are relative to the named object, not physical file offsets.
Command is the object at unit +0x1c4; movement is at unit +0x1c0; programme is
at Group +0x3c. Command +8, command +9, movement +9 and unit +0x50 are distinct.
SAV command coordinates are within the raw 184-byte block; programme coordinates
are within its raw 80-byte block. These transfers have Medium composed LOAD
confidence, with complete key fixups and client restoration Unknown.
R2-ENGINE-823 is an EN transfer read; RU field offsets are placed in the same
raw blocks by equal layout, not by a RU transfer read. This bound applies to
every SAV coordinate below.
R2-ENGINE-541, R2-ENGINE-554, R2-ENGINE-823.

## Group programme fields

| Input/state field | Read/write and SAV coordinate | Claims and clause grades | Unknowns |
|---|---|---|---|
| byte +0x20 | Programme selector; same byte in SAV programme block. | R2-ENGINE-882 High dispatch; R2-ENGINE-1040 High named setters, Medium writer completeness; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Callers of the named setters and whether a player order reaches them |
| point +0xa | Programme-2 goal; within SAV programme block. Width not strengthened here. | R2-ENGINE-507 High RU dispatch; R2-ENGINE-882 High body, Medium meaning; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Complete point producers |
| dword +0x24; word +0x28 | Mean x/y packed as y<<16 or x and y<<8 or x; same SAV block coordinates. | R2-ENGINE-841 High stores, Medium helpers; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Load-time centre initialization |
| bytes +0x2a/+0x2b/+0x2c | Largest member distance, largest +0xa5, largest sum; same SAV coordinates. | R2-ENGINE-841 High writes, Medium semantic inputs; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Reader census |
| bytes +0x2d/+0x38 | Working and baseline leash radius; same SAV coordinates. | R2-ENGINE-843 High writes, Medium draw; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Every setter and draw denominator |
| +0x30/+0x34; byte +0x39 | Member-presence latch/counter writes and setter clear; same SAV block positions. | R2-ENGINE-843 High writes, Medium broader meanings; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Full type/meaning beyond cited accesses |
| byte +0x45; dword +0x48 | Active and forced activity; same SAV block positions. | R2-ENGINE-057 High branches/stores, Medium complete population; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Full activity semantics |

## Unit and command fields

| Input/state field | Read/write and SAV coordinate | Claims and clause grades | Unknowns |
|---|---|---|---|
| unit dwords +0x50/+0x54/+0x58 | Order, actor state, action substate; separately transferred scalars. | R2-ENGINE-541 High order; R2-ENGINE-543 High substate flow; R2-ENGINE-823 Medium EN LOAD; RU coordinates by equal layout | Later overwrites and complete restoration |
| unit words +0x94/+0x96 | Health/max; read by Guard B and post-step. This page establishes no physical SAV offset for them. | R2-ENGINE-844 High comparisons; R2-ENGINE-506 High predicates | Complete save coordinate here |
| command word +0; word +2; +4 | Post; point word +2 and flag +4 as read by the order-0xa tail (order name not established). Same positions within raw SAV command block. | R2-ENGINE-844 High post; R2-ENGINE-872 High order-0xa tail; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | All post/point/flag producers |
| command bytes +8/+9 | Prepared/primary selectors; SAV command bytes +8/+9. | R2-ENGINE-542 High local flow; R2-ENGINE-553 High full primary table; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Planner/executor cadence |
| command word +0xa; +0xc; byte +0x14; byte +0x15 | Prepared goal, action target, range, phase; SAV same block positions, with target key fixups open. EN LOAD keeps a saved +0xc = 0, and unit order 3 then passes it to the planner, which faults. | R2-ENGINE-505 High RU local writes, Medium helper meanings; R2-ENGINE-937 High EN postload, Medium writer completeness; R2-ENGINE-936 High EN fault; R2-ENGINE-543 High phase; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout; R2-ENGINE-1012 High exception, Medium composed chain: separate EN fault in the LOAD-driven session update, on the unit send path | Complete target resolution |
| movement bytes +0, +1, +0xa, +0x9d; dword +0xa0; byte +0xa4 | Facing, target heading (multiples of 32), RotationSpeed, turn counter, turning flag and step count, written by the turn; the cast and strike admission tests compare +0 with the heading to the target or cell. This page establishes no SAV coordinate for them. | R2-ENGINE-927 High turn writes, Medium cadence and snap threshold (L2.01523 unread); R2-ENGINE-926 High test; R2-ENGINE-928 High cell heading; R2-ENGINE-543 High cast flow | RotationSpeed values; how the client shows the facing |
| command word +0x16 | Route-failure stamp, constructor 65500; same SAV word. | R2-ENGINE-846 High constructor; R2-ENGINE-873 High stamp; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Untraced writers and clock restore |
| command +0x20 | Assigned target; same SAV block position, not a physical address guarantee after LOAD. | R2-ENGINE-881 High writes, Medium census; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Remapping all stored keys |
| command +0x40/+0x44 | Withdraw/Wimpy health thresholds; same SAV block positions. | R2-ENGINE-506 High Human setup; R2-ENGINE-777 High stat stores, Medium row delivery; R2-ENGINE-884 Medium creature thresholds; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Creature parameter values and all writers |
| command +0x50 | Reset-chain marker, distinct from unit order +0x50; raw SAV command position +0x50. | R2-ENGINE-541 High local distinction; R2-ENGINE-555 Medium inherited reset-chain effect; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | All meanings of marker |
| command +0x54; word +0x58; byte +0x5a | Look-around flag; attacker-cell memory; builder-run age. SAV same positions. | R2-ENGINE-845 High stores; R2-ENGINE-502 High builder; R2-ENGINE-875 Medium census; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Other event consumers |
| command +0x7c+4i / +0x88+4i, i=0..2 | Spell slots and probability thresholds, raw SAV positions. | R2-ENGINE-652 High stores, Medium row binding; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | No-slot initialization and later writers |
| command +0x94; dword +0x98; word +0x9c | Attacker reference; last-hit server dword; away-cell goal. SAV same positions; pointer fixups remain open. | R2-ENGINE-845 High stores, Medium geometry; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Untraced pointer readers and full key resolution |
| command byte +0x9e; word +0xa0 | Guard-B counter and alternate goal, same SAV coordinates. | R2-ENGINE-876 High local path, Medium census; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | External starts/restored nonzero values |
| command word +0xb0 | Cell-record destination (x low byte, y high); same SAV word, not an order. | R2-ENGINE-557 High writer/consumer, Medium completeness; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Cell-record kind 0x17 |
| command list +0xb4 | Point word list +0xb4 as read by the order-0xa tail follows command raw transfer; it is not a stable serialized pointer. | R2-ENGINE-872 High tail accesses; R2-ENGINE-823 Medium EN transfer/list replacement; RU coordinates by equal layout | Complete list restoration; order name not established |

## Movement and clock fields

| Input/state field | Rule and stored boundary | Claims and clause grades | Unknowns |
|---|---|---|---|
| movement +9 | Retry byte; separate from command primary +9. | R2-ENGINE-874 Medium retry composition; R2-ENGINE-554 High local base distinction | SAV coordinate not established by these cards |
| movement +0x6c/+0x74/+0x76/+0x7c/+0x8c | Walk mode, cached goal, destination, target cache and route-end target cell in the read paths. | R2-ENGINE-883 High walk writes; R2-ENGINE-872 High destination writes; R2-ENGINE-874 Medium approach composition | Complete field lifecycle and SAV representation |
| movement dword +0x98 | Empty-route flag set by walk, approach, repath; read/cleared in common executor tail. It is not command +0x98, the hit clock. | R2-ENGINE-872 High bodies, Medium census | Untraced readers/writers and SAVE/LOAD coordinate |
| route +0x1c8; local list +0x1e4 | Path lists used by search/repath; not command objects. | R2-ENGINE-871 High branches; R2-ENGINE-872 High emptiness tests; R2-ENGINE-879 High builder, Medium list effects | Complete list owner/serialization layout |
| server word +0 / dword +4 | Stamp clock / hit clock respectively. Saved head-second/head-first interpretations retain Medium. | R2-ENGINE-845 High hit read; R2-ENGINE-846 High stamp read; R2-SESSION-343 Medium clock relation | Seconds per Group tick, wrap behaviour, precise restored clocks |

## Saved controls

| Population/state | Bounded observation | Claims and clause grades | Unknowns |
|---|---|---|---|
| Five original mission-10 saves, 150 member blocks | Every +0x16=65500, +0x9e=0, +0xa0=0; +0x9c nonzero only in three blocks with +0x54=1. No saved stamp/case-5 witness. The single saved command +8=5 block (game0006, Group 6, programme 18, target set, +0x16=65500) is the single saved case the stamp's condition could apply to. | R2-SESSION-346 High generated values, Medium state reading | Whether that block's substitute was in range or goal reachable is Unknown; tables not reproduced by this composition; movement block absent from that save reader |
| Same five files, 60 nonhero Groups | All programme 18; 33 with +0x30=1 have radius-baseline 3..5, 27 with +0x30=0 have equality. | R2-SESSION-343 High counts, Medium clock/pursuit interpretation | No proof of jitter distribution or continuous trajectory |
| Thirteen mission-10 documents | Five original files have stamp 65500 and relation diagonal 0x12; eight compared written/resaved files have stamp 0 and diagonal 2. These are parser-path facts, not runtime control equivalence. | R2-SESSION-339 High values | Writer diagnosis outside this page; next live case |
| Five mission-10 snapshots, grouped Unit-class | 138/139 unit orders are 0xb; one is 0, as generated by the cited claim's experiment, not reproduced here. | R2-ENGINE-884 High count, Medium generalization to continuation | Other ticks, creations and orders |

## Health thresholds, difficulty and attack order

Withdraw and Wimpy are Units columns (R2-ASSET-229). The stat reader stores them in command +0x40 and +0x44 for a unit whose template loads (Medium: loader condition, script-created creatures; R2-ENGINE-1001). The map arm scales a non-Human record by the difficulty word when Session +0x74 is 0 and the value is not 2 (R2-ENGINE-1002). In the bodies read, the only direct caller of the threshold post-step is the group tick (byte scan; indirect and jump routes unread, Medium; R2-ENGINE-1003). The order-3 arm passes command +0xc to the planner with no null test (R2-ENGINE-1004). In the bodies read, LOAD leaves the thresholds as saved (R2-ENGINE-1005).

## Fields written by the read next-action bodies

The wrapper tests owner +0x2c at unit +0x14 and reaches the same selector on both branches (R2-ENGINE-1017, Medium). The idle planner's own body writes none of unit +0x50, command +8, +9, +0xc, +0x20 (R2-ENGINE-1018). The stop chain resets unit state +0x54 and stores order 0xc (R2-ENGINE-1020). Session list +0xbbec has appenders of record kinds 1 to 4, a routine that clears it in a navigation read and no read operand found in a RU sweep (R2-ENGINE-1021, Medium).
The Group pass runs only when session counter +4 mod 16 is 6 (R2-ENGINE-1016). The re-pick writes command +0xc and +8 and no unit +0x50 or command +0x20 (R2-ENGINE-1019). The write census is by base class and does not count unit-order writers (R2-ENGINE-1022).

Movement +0x98: census sets L2.02432, L2.02433, L2.02434 and clear L2.05192 (R2-ENGINE-1042, Medium census, arrival unread). L2.01525 stores command +8, +0xc and +0x14 (R2-ENGINE-1043). L2.05852, called with ecx = session +0xbbec, walks a chain from [this+4] and clears +4, +8, +0xc, +0x10 and +0x14 (R2-ENGINE-1045, helpers unread).

The programme-3 route L2.02466 makes no direct call to the selector; no reference to loop L2.01724 was found in the Medium census, and an indirect or computed route is not excluded (R2-ENGINE-1041).

Counted run (EN, one mission-10 launch, one hero): R2-ENGINE-1072 (loop L2.04213 and wrapper L2.02403 had 0 hits in the span; route serving the hero Unknown in launch 1), R2-ENGINE-1073 (actor step once per scheduler counter value per unit; the counter advanced 16.1 and 32.1 per second in the two launches), R2-ENGINE-1074 (a ground click wrote order 1 and moved the hero to another Group object).

Continuation reads: R2-ENGINE-1075 (EN return addresses L2.06121 and L2.06122 of L2.04984 lie in the idle planner's no-candidate arm and in the programme-3 route; which of the planner's six E8 callers served the hero is Unknown in launch 1), R2-ENGINE-1076 (RU programme 4 and Guard A bodies, including Guard A's arm L2.01733(member, 1), which does not call the idle planner; the empty-Group store of programme 0xff), R2-ENGINE-1078 (a second EN launch: with no attack order the hero read order 0xc and a strike state on another owner's unit, and the idle planner's entries by return address), R2-ENGINE-1077 (256-instruction masked comparison of RU and EN bodies; 16 pairs equal, seven differ at a later index, the EN twin of L2.05852 stays ambiguous).

## Coordinates of the continued machine

Unit order +0x50 and target +0x5c differ from command-object fields at the same
offsets. Census counts are operand populations, not resolved unit writers
(R2-ENGINE-1104); strike and target-cast initializers write unit +0x5c, while
point cast clears command +0x5c (R2-ENGINE-1105). The 20 paired setters store
17 values and their parameters (R2-ENGINE-1107). Initiation can use the range
producer's result, whereas the read HP-application gate uses unit +0x12c
(R2-ENGINE-1106).

Movement +0x98 empty-route and +0x90 arrival signals are distinct
(R2-ENGINE-1108). Hit memory, assigned target +0x20 and prepared target +0xc
remain distinct (R2-ENGINE-1109). Darkness adjusts words +0xa4/+0xe4; +0xa5
is the high byte of the former, with kind-specific readers still bounded
(R2-ENGINE-1110). Scorer L2.05238 and Group candidate lists decide the named
assignment stores (R2-ENGINE-1111, R2-ENGINE-1112); Group programme +0x20 is a
separate record field (R2-ENGINE-1113). EN/RU clock offsets differ
(R2-ENGINE-1114). None of these coordinates implies complete twin coverage
(R2-ENGINE-1115); open writer/reader populations are on the
[planner page](orders.md).

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
