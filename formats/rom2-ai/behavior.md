# Group and member decisions

## Directed relation updates and admission

Script action 10 changes M[p0][p1] by mask-and-add and invokes the relation
producer for matching endpoint records. In the action and producer bodies
read, it directly changes neither the reciprocal cell nor a unit order. Packet word Q for receiving P carries
M[P][Q] in its low bits and reverse hostility as bit 0x40. Local stores and
packet arithmetic are High; iteration and delivery remain Medium or Unknown.
— R2-ENGINE-979, R2-ENGINE-980

The hit relation helper first ORs bit 0 into M[attacker][victim] only when
that cell's low two bits are zero. It treats the reverse cell separately,
also requiring low two bits zero. A zero third argument admits the reverse
write immediately (RU call sites read: two of three pass 0; which game
actions reach each is Medium). A nonzero argument additionally requires
victim command +0xac signed >10, health below half max, or state 0x10;
otherwise it only increments that counter. The complete bodies are High in
both roots; the hook-to-helper edge and its pushes are read in both roots
and equal after masking; complete combat entry and notification delivery
remain Medium. — R2-ENGINE-981, R2-ENGINE-987

The listed hover block tests a Player's indexed relation word &1 and derives
flag 4 (that Player as the local player is Medium, inherited). The selected
cursor fragment uses that flag for an attack cursor; upstream gates are
unread. An attack
cursor click sends 0x19 without another local relation test. The server
acquisition predicate tests M[querying owner][candidate owner]&1, not the
reverse cell. Local predicates are High, composed cursor/attack/acquisition
remain Medium. Mission 10 can therefore have M[7][1]=1 and M[1][7]=0 after
the ogre action, until an admitted hit (F=0, or the F nonzero gates)
or another writer changes it. The owner identity of the hero and ogre units
and the trigger's firing condition are Unknown. That case is Medium, with
observed cursor, delivery and next action Unknown. Two-root code
comparisons cover only the listed windows. — R2-ENGINE-982,
R2-ENGINE-983, R2-ENGINE-984

## Same-owner area blow

An area pulse stores the caster as area +0x3c when the caster's own +0x3c
is nonzero. Each occupant is then handled with no owner or relation test:
row 17 nothing; row 2 the damage copy, other rows the area object's attach;
row 15 nothing more; every other row a caster pointer in occupant +0x40 when
the caster has an owner (not for row 14) and the hit hook (caster,
occupant, 1). The RU body alone skips the hook when a session flag
and the caster owner's +0x2c are nonzero and the occupant owner's +0x2c is
zero; it cannot be met when caster and occupant share one owner object (an
equal owner index is Medium evidence for that). The occupant body is High in both roots; the caster
field at the apply is Medium. — R2-ENGINE-985

By the code read in both roots, the Effect attach keeps the old same-ID
object and kind, maximizes signed
magnitude and zero-extended duration independently, cancels Bless against
Curse, or adds a copy and its bit in unit +0x144. The direct body reads no
owner. Negative renewal keeps the less negative magnitude. Darkness's slots
add/subtract magnitude<<8 at A4/E4 only when virtual30 is nonzero; final sight
and target effects remain Unknown. Direct stores are High, complete
list/callee composition Medium. — R2-ENGINE-986, R2-ENGINE-1164,
R2-ENGINE-1166, R2-ENGINE-1167, R2-ENGINE-1168

The hook writes hit memory, session +0xbd98/+0xbdc4/+0xbdc8 and one record
appended to session list +0xbbec, which the RU script tick clears (Medium:
navigation listing; EN tick unread). The hook body writes no command +0x20,
+8 or +0x0c; its callees and the EN append are unread. For one owner index i the
relation helper reads M[i][i] in both directions: low bits set, nothing is
written and +0xac is not incremented; low bits clear, bit 0 is ORed into the
diagonal. Group sight marks the remembered attacker cell for a member
whose owner +0x2c is nonzero and lists units on it; the hostility filter
then decides. —
R2-ENGINE-987, R2-ENGINE-988

Composed, in the read bodies: with the diagonal's bit 0 clear and low bits
set, and the owner's +0x2c nonzero so that group sight marks the caster's
cell, the filter drops the same-owner caster, assignment keeps the
least-score hostile, and a
member with a target stays in Guard B's target case; a member without one
and below max health walks to command +0x9c while the hit clock is fresh.
This path is Medium, not observed; what the Darkness Effect does to sight
range (it can drop the old target) and the +0xbbec readers are Unknown.
Mission 10's Troll.4 and the four
Bat_Sonic.4 share owner word 5, whose Player row stores diagonal 0x12 (High
fields, Medium names). — R2-ENGINE-989, R2-ENGINE-990

## Group tick

| Input/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Ordinary session, session +0xbbe8 zero | Recount activity. Group tick L2.00504 calls L2.02456 only if session +0xbbe8 or programme +0x45 is nonzero. Inactive Group and command +9 = 0 makes the unit executor set state 27 and return. | R2-ENGINE-057 High for branches; R2-ENGINE-507 High RU gate | Owner classification and other session gate modes |
| Activity marking | Owner +0x2c = 0 actors mark owner masks over a 5 by 5 coarse-cell neighbourhood. Actors of the alternative owner branch (owner +0x2c nonzero) below max health can activate their Group. A nonzero cell mask or forced +0x48 marks active. | R2-ENGINE-057 High local stores, Medium producer population | All marking exceptions; GUI fog equivalence is not established |
| Empty Group | Tick sets programme byte +0x20 to 0xff before dispatch. | R2-ENGINE-507 High RU; R2-ENGINE-882 High both dispatch bodies | Whether a guard's no-member radius arm can be reached |
| Programme dispatch | Byte table RU L2.02461 / EN L2.05245 maps 0..19 into jump table RU L2.02462 / EN L2.05246; 0xff has its own branch. Post-step follows the selected route or no route. | R2-ENGINE-882 High | Seconds per tick and relation to member executor calls |

## Programme routes

Order is preserved within the clauses each card reads. A named unread helper is
an explicit boundary; its internal decisions are not supplied here.

| Programme/state | Per-member rule in the read order | Claims and clause grades | Unknowns |
|---|---|---|---|
| 0 | Group route L2.02463 runs each member through L2.01726, the unit-order selector wrapper. | R2-ENGINE-507 High RU route; R2-ENGINE-882 High body; R2-ENGINE-558 High wrapper, Unknown rate | Planner scheduling |
| 1, Guard A, L2.02464 | Guard preparation below; with target command +0x20 run the attack planner; otherwise, in the EN arm read by R2-ENGINE-824, walk to the post unless centred there. R2-ENGINE-844 states the case subset; the RU Guard A body was not read. | R2-ENGINE-841 High preparation; Guard A member decision: one sentence of R2-ENGINE-844 (cases 3 and 6 only; the card's High is for Guard B), Medium for Guard A; R2-ENGINE-824 Medium, EN mode-1 arm | RU Guard A body not read (R2-ENGINE-882); helper effects; final coordinate from a zero post |
| 2, L2.02465 | Sight, assignment L2.02424; target: planner; no target: walk to Group point +0xa, or L2.01733(member, 1). | R2-ENGINE-882 High body order, Medium callees; R2-ENGINE-507 High RU route | Exact arrival/helper conditions; no direct leash call does not exclude one in callees |
| 3, L2.02466 | Sight, assignment L2.05173 using scorer L2.05238; target: planner; no target: L2.01531. | R2-ENGINE-882 High body; R2-ENGINE-881 High assignment writes, Medium census | Scorer's complete decision; sight/support callees |
| 4, L2.02467 | Read body calls L2.04666, walk, L2.01576, L2.01733, L2.02114 and L2.01731. It does not access target +0x20. The reset chain can store unit order 0xc. | R2-ENGINE-882 High call/access body, Medium callees; R2-ENGINE-884 Medium order composition | Branch conditions connecting those calls; complete per-member execution |
| 5, L2.02468 | Sight; with enemies, assignment then planner on a target or L2.01733(member, 1); without enemies, programme 4. | R2-ENGINE-882 High body order, Medium callees | Sight, planner and move helper effects |
| 17, L2.02469 | Distance tests and L2.02421, then programme 5. | R2-ENGINE-882 High local sequence, Medium meanings; R2-ENGINE-507 High RU route | Wander destination and distance-branch semantics not specified by these cards |
| 18, Guard B, L2.02470 | Guard preparation, then the first matching numbered case below. | R2-ENGINE-841 High order; R2-ENGINE-844 High operands/order, Medium absolute value | Full path/clock composition |
| 19, L2.02471 | Each member takes L2.01531. | R2-ENGINE-882 High body, Medium callee meaning | Complete support choice |
| 0xff, L2.02472 | Each member takes L2.01733(member, 0). | R2-ENGINE-882 High body, Medium callees | Candidate/mage helper effects |
| 6..16, 20..254 | No programme route; post-step still runs. | R2-ENGINE-507 High RU; R2-ENGINE-882 High both tables | External setters and script meanings |

## Guard preparation and assignment

| Input/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Nonempty Guard A/B Group | L2.02475 writes programme +0x24 = meanY << 16 or meanX; +0x28 = meanY << 8 or meanX, with truncated integer means. Then +0x2a = maximum member distance, +0x2b = maximum member +0xa5, +0x2c = maximum sum. | R2-ENGINE-841 High bodies, Medium count/list helpers | Readers of +0x2a..+0x2c; load-time centre writer |
| Fresh mean available | Group sight L2.02473, leash L2.02474(mean +0x28, radius +0x2d), assignment L2.02424, radius step, member loop. | R2-ENGINE-841 High call order; R2-ENGINE-507 High/Medium sight description | Complete group-sight inputs |
| Listed enemy farther than radius | Remove it when Chebyshev distance from the current mean exceeds +0x2d. The member's post is not the leash centre. | R2-ENGINE-842 High operands/body; R2-ENGINE-878 High distance body | List helpers; unseen enemies |
| Enemies survive | Score each for each member; strictly least score wins, first on a tie. Write command +0x20; with no enemy or all scores excluded write 0. | R2-ENGINE-503 High arithmetic; R2-ENGINE-881 High stores, Medium writer completeness | Untraced stores, LOAD/constructor effects |
| Group scorer L2.00057 | Score = d * 256 + turn. Ranged members use d-range+1 beyond range, else 1. A 16-byte table supplies preference; 0 excludes. For size below 2 and Mind at least 15, preference 1 adds half, preference 4 subtracts a quarter. Effect 0x40000 adds 0x7f, floor 1. | R2-ENGINE-503 High arithmetic/table, Medium kind names | Semantic kind labels; target eligibility outside this scorer |
| Preference-table input | Candidate kind 1 with its range above 1 becomes kind 0. Ranged member indexes T[k], melee T[k+4*ownKind]. T rows are 2,1,2,4 / 4,2,1,0 / 2,1,4,0 / 2,1,2,4. | R2-ENGINE-503 High values, Medium kind names | Full meaning of kind values |
| No assigned target after leash | Guard A/B take their no-target clauses. A pursuit ending at this point is a composition, not a universal write census. | R2-ENGINE-842 Medium composed end; R2-ENGINE-881 Medium coverage | 141 RU untraced +0x20 writes; helpers and later execution |

## Radius latch

| Input/state after assignment | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| First member exists, +0x30 = 0 | +0x2d = +0x38 + draw(-1,1) + 4; +0x34 = 0; +0x30 = 1. | R2-ENGINE-843 High stores, Medium uniform {-1,0,1} interpretation | Draw denominator; seconds per tick |
| First member exists, +0x30 nonzero | Increment +0x34, set +0x30 = 1; radius unchanged. | R2-ENGINE-843 High | Meanings beyond these writes |
| No first member, +0x30 nonzero | +0x2d = +0x38 + draw(-1,1); +0x34 = 0; +0x30 = 0. | R2-ENGINE-843 High stores, Medium draw | Reachability after empty Group becomes 0xff |
| No first member, +0x30 = 0 | Increment +0x34, set +0x30 = 0. | R2-ENGINE-843 High | Same reachability boundary |
| Mode-18 setter(group, r) | Set members' order 0xb, programme 18, run radius setup. Nonzero r above +0x2d raises +0x2d/+0x38 to r; r=0 and +0x2d below session +0xa8b8 writes 20 to both. Clear +0x39. | R2-ENGINE-843 High; R2-ENGINE-541 High setter body | Setter callers and r values |
| Guard-A setup in the Summon path | L2.04403 sets programme 1 and unit order 0xb; L2.04666 and setup copy +0x2c to +0x2d/+0x38 and raise both to session MinimalGuardRange when lower. | R2-ENGINE-788 High stores, Medium config role/value | Configuration source/value; not all creation paths |

## Guard B cases

Take the first matching row, in this order. Difference denotes the subtraction
followed by the presumed absolute-value helper; that interpretation remains
Medium. The word and dword clocks are distinct. R2-ENGINE-828, R2-ENGINE-844.

| Case/input | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| 1: HP differs from max; word difference(command +0x16, server +0) below 30 | Prepare walk to command +0x9c, even with an assigned target. | R2-ENGINE-828 High order/operands, Medium trajectory; R2-ENGINE-844 High stores, Medium absolute value | Clock restore, helper and route outcome |
| 2: same word difference below 10 | Prepare the same walk, with no HP or hit-memory test. | R2-ENGINE-844 High operands/stores, Medium absolute value | Same clock/path boundaries |
| 3: command +0x20 nonzero | Run attack planner L2.01530 on the assigned target. The EN planner L2.03461 reads target +0x4c before any null test; this case passes only a nonzero target, while EN unit order 3 (L2.05581) passes command +0xc unchecked. | R2-ENGINE-844 High branch; R2-ENGINE-505 High RU body, Medium range/step interpretation; R2-ENGINE-936 High EN read order; R2-ENGINE-937 High read order-3 writers store no target 0, Medium completeness; R2-ENGINE-1012 High exception, Medium composed chain: separate EN fault in the LOAD-driven session update, on the unit send path | Planner helpers; whether target remains eligible |
| 4: no target; HP differs from max; dword difference(command +0x98, server +4) below 10 | Prepare walk to command +0x9c. | R2-ENGINE-844 High branch/stores, Medium absolute value; R2-ENGINE-845 High hit stores, Medium away-cell geometry | Live clock and destination result |
| 5: command byte +0x9e nonzero | Increment counter, walk to +0xa0, clear counter past 10. The read path that would start it from zero always jumps over initialization. | R2-ENGINE-876 High body/path, Medium census; R2-ENGINE-844 High walk | Untraced writers, LOAD and outside/indirect entries |
| 6: off post +0 or not centred | Prepare +8 = 1, goal +0xa = post. Otherwise run L2.01733(member,1): owner +0x2c nonzero writes +8 = 0xb; otherwise L2.02911. | R2-ENGINE-844 High body/arrival test, Medium helper meanings | Full support/mage path; +8 = 0xb is an activity, not unit order 11 |
| Guard A subset | R2-ENGINE-844 states only cases 3 and 6. The EN mode-1 no-target arm passes a zero post unchanged to prepared movement. | R2-ENGINE-844 one sentence, Medium for Guard A; R2-ENGINE-824 Medium, EN only; R2-ENGINE-883 Medium search composition | RU Guard A body not read (R2-ENGINE-882); actual zero-goal movement; prior setters |
| Walking home on later Group tick | Sight, leash and assignment run again before these cases; no read body establishes an unconditional return-home lockout. | R2-ENGINE-847 Medium composition | All writer/helper effects and original observed case |

## Struck member and look-around flag

| Input/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Hit hook L2.02074(attacker,victim,third) | Victim command +0x54=1, +0x58=attacker cell, +0x5a=0; call L2.02419, post copied event-1 record through L2.01727; +0x94=attacker, +0x98=server dword +4, +0x9c=away cell. No target +0x20, prepared +8 or +0x0c store in the hook body (callees and EN append unread). The record goes to the tail of session list +0xbbec, which the RU script tick clears. The area occupant routine is one caller, with third argument 1. | R2-ENGINE-845 High stores/calls, Medium geometry/caller roles; R2-ENGINE-981 High relation-helper body; R2-ENGINE-987 High EN/RU hook and RU append/tick, Medium list clear; R2-ENGINE-985 High occupant call | Posted-record readers before the clear; full caller coverage and notification delivery |
| Remembered cell | For a computer-owned actor, candidate builder marks it for runs 1..21, increments +0x5a and clears +0x58 after marking on run 21. Group sight also reads remembered cells: for a member whose owner +0x2c is nonzero it marks +0x58, increments +0x5a, clears +0x58 past 20 and lists units on marked cells before filter L2.02119. | R2-ENGINE-502 High local builder; R2-ENGINE-875 Medium group-sight census; R2-ENGINE-988 High group-sight body, Medium filter output list | Full group-sight predicates; builder-run cadence against the Group tick is Unknown |
| Candidate admission | Builder ray scan marks cells; hostility bit and visibility remove candidates; health below 1 separates dead candidates. | R2-ENGINE-502 High local arithmetic/filter, Medium creature scan range/table inputs | Complete group-sight population; shared builder does not prove identical pipelines |
| Look-around flag | +0x54 requests a new heading and is cleared by look-around. The hit hook does not directly order every Group to attack the attacker. | R2-ENGINE-504 High RU look-around, Medium EN/helper meaning; R2-ENGINE-845 High no direct target write, Medium broader inference | Event consumers and untraced target writers |
| Stored attacker pointer | +0x94 has no reader in the census with traced command bases; +0x58 has builder/group-sight readers, +0x98/+0x9c have Guard B readers. | R2-ENGINE-875 Medium census; High for the Guard B readers of +0x98/+0x9c; Medium for the builder and group-sight readers of +0x58 | 268 RU +0x94 sites with untraced bases; aliases and bulk copies |

## Health and hostile post-step

| Input/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Each member after programme route | Health <= command +0x44 and hostile within movement byte +8 + 5: flee. Otherwise health <= +0x40 and hostile within 2: back-off helper L2.02451. | R2-ENGINE-506 High predicates, Medium helper destination; R2-ENGINE-877 Medium bounded route independence | All threshold writers; units outside this Group post-step |
| Human-table threshold setup | Spellbook: class 2, +0x40=HPmax, +0x44=HPmax/4. Else range>1 and kind!=3: class 1, HPmax/2 and 0. Else class 0, both zero. | R2-ENGINE-506 High local stores | Further writer effects |
| Creature thresholds | Stat reader can supply Withdraw/Wimpy to command +0x40/+0x44, keeping defaults for -1. Human-only setup does not imply creature thresholds stay constructor values. | R2-ENGINE-777 High reader, Medium delivered row; R2-ENGINE-884 Medium creature composition | Creature row values and complete writer population |
| Flee chosen | Try spell 23 through admission L2.01518; otherwise move 3 cells away from living hostiles' centre. No living hostile: default route. | R2-ENGINE-506 High flee body, Medium centre/away interpretation | Spell admission and concrete destination |
| RU creation arm with session word +0xd8=1 | Halve command +0x44. | R2-ENGINE-884 High stores, Medium selection meaning | Meaning of +0xd8 and complete EN counterpart |

## Health thresholds, difficulty and attack order

Withdraw and Wimpy are Units columns (R2-ASSET-229). The stat reader stores them in command +0x40 and +0x44 for a unit whose template loads (Medium: loader condition, script-created creatures; R2-ENGINE-1001). The map arm scales a non-Human record by the difficulty word when Session +0x74 is 0 and the value is not 2 (R2-ENGINE-1002). In the bodies read, the only direct caller of the threshold post-step is the group tick (byte scan; indirect and jump routes unread, Medium; R2-ENGINE-1003). The order-3 arm passes command +0xc to the planner with no null test (R2-ENGINE-1004). In the bodies read, LOAD leaves the thresholds as saved (R2-ENGINE-1005).

## Reach of the per-unit selector and the event list

Selector L2.01725 is reached through L2.01726 from the programme-0 route; the per-unit loop L2.01724 has no reference found (R2-ENGINE-1017, Medium). The executor tail is bounded by R2-ENGINE-1020. Records appended to session list +0xbbec have no read operand found in a RU sweep; aliased readers stay open (R2-ENGINE-1021, Medium). Programme 3 route L2.02466 makes no direct call to the selector; in the census of R2-ENGINE-1017 (Medium) a hero in programme 3 reaches it only if loop L2.01724 runs, and an indirect or computed route is not excluded (R2-ENGINE-1041, Unknown composition).
The Group pass runs only when session counter +4 mod 16 is 6 (R2-ENGINE-1016). The idle planner ranks list candidates by distance then turn (R2-ENGINE-1018) and the re-pick writes +0xc and +8 (R2-ENGINE-1019). The write census counts no unit-order writer at +0x50 (R2-ENGINE-1022).

Group methods write programme byte +0x20 (R2-ENGINE-1040, High bodies, Medium completeness). A found target stores +8=6, +0xc and +0x14 at L2.01525, with a mage branch for owner +0x2c zero (R2-ENGINE-1043). The switch L2.03275 has 27 arms, each one call. Routine L2.05852, called with ecx = session +0xbbec at L2.06082, walks a chain and passes each record pointer to unread L2.06061; whether any arm or that helper reads a record is Unknown (R2-ENGINE-1045, R2-ENGINE-1021).

Counted run (EN, one mission-10 launch, one hero): R2-ENGINE-1072 (loop L2.04213 and wrapper L2.02403 had 0 hits in the span; route serving the hero Unknown in launch 1), R2-ENGINE-1073 (actor step once per scheduler counter value per unit; the counter advanced 16.1 and 32.1 per second in the two launches), R2-ENGINE-1074 (a ground click wrote order 1 and moved the hero to another Group object).

Continuation reads: R2-ENGINE-1075 (EN return addresses L2.06121 and L2.06122 of L2.04984 lie in the idle planner's no-candidate arm and in the programme-3 route; which of the planner's six E8 callers served the hero is Unknown in launch 1), R2-ENGINE-1076 (RU programme 4 and Guard A bodies, including Guard A's arm L2.01733(member, 1), which does not call the idle planner; the empty-Group store of programme 0xff), R2-ENGINE-1078 (a second EN launch: with no attack order the hero read order 0xc and a strike state on another owner's unit, and the idle planner's entries by return address), R2-ENGINE-1077 (256-instruction masked comparison of RU and EN bodies; 16 pairs equal, seven differ at a later index, the EN twin of L2.05852 stays ambiguous).

## Target choice and continuation boundaries

Assignment L2.05173 rejects adjusted distance above 1 and selects the first
strictly least score; preference/kind meanings remain bounded
(R2-ENGINE-1111). Group sight gathers marked-cell world units before relation,
effect and health filtering; leash loss followed by assignment is only a
bounded pursuit-ending path (R2-ENGINE-1112). Programme setters are separate
from member-order setters (R2-ENGINE-1113, R2-ENGINE-1107). The direct writer
census and named target writers do not close all live stores
(R2-ENGINE-1104, R2-ENGINE-1105).

Strike admission and HP application use distinct tests (R2-ENGINE-1106), and
empty-route/arrival branches have separate signals (R2-ENGINE-1108). A hit
hook's local stores do not prove later order retention (R2-ENGINE-1109).
Darkness adjusts overlapping sight-coordinate words for the read predicate,
with complete target effects still open (R2-ENGINE-1110). Delivered cadence
and per-body EN coverage keep their limits (R2-ENGINE-1114,
R2-ENGINE-1115). Remaining readers, predicates and runs are named on the
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
