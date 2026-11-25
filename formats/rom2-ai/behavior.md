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
add/subtract magnitude<<8 at A4/E4 only when virtual30 is nonzero; Darkness
changes sight only for objects whose virtual +0x30 returns nonzero, and for a
Unit-class object it attaches and changes no sight (Medium for the class of any
named creature; R2-ENGINE-1237, R2-ENGINE-1238); the target-drop consequence
stays Unknown. Direct stores are High, complete
list/callee composition Medium. — R2-ENGINE-986, R2-ENGINE-1164,
R2-ENGINE-1166, R2-ENGINE-1167, R2-ENGINE-1168

The hook writes hit memory, session +0xbd98/+0xbdc4/+0xbdc8 and one record
appended to session list +0xbbec, which the RU script tick clears (Medium:
Medium: RU navigation listing and the EN drain, R2-ENGINE-1451). The hook body writes no command +0x20,
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
This path is Medium, not observed; the Darkness Effect changes sight only for objects whose virtual +0x30 returns
nonzero (R2-ENGINE-1237, R2-ENGINE-1238), whether the lowered sight drops the old
target is Unknown, and so are the +0xbbec readers.
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
| Mode-18 setter(group, r) | Set members' order 0xb, programme 18, run radius setup. Nonzero r above +0x2d raises +0x2d/+0x38 to r; r=0 and +0x2d below session +0xa8b8 writes 20 to both. Clear +0x39. | R2-ENGINE-843 High; R2-ENGINE-541 High setter body | Setter callers and r values; the arguments of the three Guard A setter callers per root (RU L2.07168, L2.07169, L2.07170; EN L2.07171, L2.07172, L2.07173) are unread |
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

## Area hit numbers

Every row is one number of the struck unit's reaction. A hit rewrites the six command fields and posts a record; the
fields' readers and the repeat behaviour are the rows below.

| Number | Value | Claim and grade |
|---|---|---|
| Command +0x54 after a hit | 1; the one reader found in the census is the look-around, which clears it | R2-ENGINE-1491 High, readers Medium |
| Command +0x58 after a hit | the attacker's cell; five sight copies read it | R2-ENGINE-1491 High, readers Medium |
| Command +0x5a after a hit | 0, so each hit restarts the sight count | R2-ENGINE-1491 High |
| Command +0x94 after a hit | the attacker; no reader in the census | R2-ENGINE-1491 High, absence Medium |
| Command +0x98 after a hit | the session tick count; one reader found in the census, Guard B case 4 | R2-ENGINE-1491 High, reader Medium |
| Command +0x9c after a hit | a cell up to 10 cells from the victim, away from the attacker | R2-ENGINE-1491 High, sign Medium, truncation of L2.01309 Medium |
| Away vector scale | 10.0 divided by the larger absolute axis difference | R2-ENGINE-1491 High |
| Sight passes that mark the cell | 21 after the last hit (the mark clears when signed +0x5a exceeds 0x14) | R2-ENGINE-1492 High RU, EN Medium |
| Look-around new heading | heading byte + 0x21 + draw below 0xbe; taken when the flag is set or a 0..0x7fff draw is below 0xcd | R2-ENGINE-1491 High |
| Posted record per hit | 0x54 bytes (21 dwords), appended; the script tick destroys the list | R2-ENGINE-1492 High chain, no reader Medium |
| Hit counter | none: +0x5a counts sight passes and is reset by a hit | R2-ENGINE-1492 High |
| Session dword +4 | counts calls of the session tick; the tick returns at once when its object's +0x94 is 0 (R2-ENGINE-1016) | R2-ENGINE-1493 Medium: the tick object is not shown to be `[L2.01317]` |
| Session word +0 | counts blocks of 16 calls of the session tick | R2-ENGINE-1493 Medium: the tick object is not shown to be `[L2.01317]` |
| Group pass spacing | 16 session ticks; the turn routine runs when the tick count's low nibble is 6 | R2-ENGINE-1493 Medium: the tick object is not shown to be `[L2.01317]`; one direct caller (R2-ENGINE-1016, R2-ENGINE-778) |
| Guard B case 4 window | tick difference below 10 with health below maximum | R2-ENGINE-1493 High RU, EN Medium |
| Guard B cases 1 and 2 window | stamp difference below 30 blocks with health below maximum; below 10 blocks without the health test | R2-ENGINE-1493 High RU, EN Medium |
| Step-away passes per hit | case 4: at most 1 (for 10 of the 16 phases of the hit tick the next pass is within 9 ticks); cases 1 and 2: every pass walks to +0x9c while the stamp is under 30 or 10 blocks old, whatever the hit | R2-ENGINE-1493 Medium |
| Guard A with no target | walks to its post or idles; reads none of +0x16, +0x94, +0x98, +0x9c | R2-ENGINE-1494 High, absence Medium |
| Seconds per session tick | Unknown; 16.13 and 32.08 counter values per second were measured in two EN launches under a debugger and the difference is open | R2-ENGINE-1493 Unknown, R2-ENGINE-1073 |
| Area spell point | no body read chooses a point: the offensive choice casts at the victim unit and the special cast at the victim's cell; the centre of the target-mode apply L2.00034 is Unknown; no unit test of the area | R2-ENGINE-1490 High for the bodies, Medium for the composition |
| Candidate score, scorer L2.00057 | preference 1: s + s/2; preference 4: s - s/4, only when the actor's vtable +0x1c result is below 2 and its word +0x88 is at least 15; effect bit 0x40000 of +0x144: +0x7f, floor 1 | R2-ENGINE-1495 High both roots (bodies); callers RU |
| Candidate score, scorer L2.05238 | adjusted distance above 1: 0xffffff; preference 1 doubles, preference 4 halves, no conditions, no effect test | R2-ENGINE-1495 High both roots (bodies); callers RU |
| Effect bit 0x40000 | bit of +0x144 is effect id 18; the +0x150 bit 0x40000 is the separate field-changed flag; the two scorers read no other effect word, and an effect that changes the Mind word, the range or the virtual kind is not excluded | R2-ENGINE-1496 High for the bodies, Medium for the rest |
| Group assignment target | every member of the group argument gets the lowest-scoring enemy of the list +0xa68, or 0; no idle test | R2-ENGINE-1497 High, completeness Medium |
| Writers of command +0x20 | 7 command-block stores in the assignment routines; two unclassified dword stores near the group code and four unclassified non-dword stores outside the group code range | R2-ENGINE-1497 High, Medium |
| Flee without Teleport, empty route | stop, then order 0xc at the own cell; the next Group pass decides again | R2-ENGINE-1498 High bodies, Medium composed; Medium that a flee walk reaches the order-1 branch (movement +0x98 nonzero and +0x50 = 1, R2-ENGINE-1020) |
| Withdraw v | message dword +0xe = 0, 1, 2 gives 0, 10, 30 (other 0); the owner is the object L2.01565 returns for message word +5 | R2-ENGINE-1499 High RU and EN; that +5 names a sender is Medium (R2-ENGINE-1179) |
| Withdraw thresholds | command +0x40 = HPmax x {0, 2v, 3v} / 100; command +0x44 = word +0x96 x v / 100 | R2-ENGINE-1499 High, R2-ENGINE-1421 |
| Withdraw outer message type, sender and control | Unknown; the outer message type that selects the arm is not recorded in the committed evidence | R2-ENGINE-1499 Unknown |

## Health and hostile post-step

The unit strike returns without a blow when a party or owner is missing, the attacker's health is not above 0, or its range byte +0x12c is below the distance; in RU it also returns when a non-mage Human-kind attacker whose attack type byte is not 5 meets a target of movementType 3 (Bat_Sonic, Dragon, Succubus rows); the EN strike body has no such test (the EN distance and callers are unread, so EN behaviour is Medium). The strike has no refusing test on the target's maximum-health word or health; it reads the target's health word to subtract the damage and to gate the notification (R2-ENGINE-1605, High for the body; the damage body is unread, Medium; Medium that a Human's byte +0xb6 is the weapon's attackType). In RU the Human threshold setter has one direct call site, in the template loader tail; a Unit takes nothing from it (R2-ENGINE-1606, amended by R2-ENGINE-1795: a Human built by `L2.00130` does not reach that loader and starts at 0, Medium; other writers of the command fields stay Unknown, R2-ENGINE-1001).

| Input/state | Rule | Claims and clause grades | Unknowns |
|---|---|---|---|
| Each member after programme route | Health <= command +0x44 and hostile within movement byte +8 + 5 (10 cells; the Teleport arm lists 5 cells): flee. Otherwise health <= +0x40 and hostile within 2: back-off helper L2.02451. | R2-ENGINE-506 High predicates, Medium helper destination; R2-ENGINE-877 Medium bounded route independence; R2-ENGINE-1440 radius | All threshold writers; units outside this Group post-step |
| Human-table threshold setup | Spellbook: class 2, +0x40=HPmax, +0x44=HPmax/4. Else range>1 and kind!=3: class 1, HPmax/2 and 0. Else class 0, both zero. | R2-ENGINE-506 High local stores | Further writer effects |
| Creature thresholds | Stat reader can supply Withdraw/Wimpy to command +0x40/+0x44, keeping defaults for -1. Human-only setup does not imply creature thresholds stay constructor values. | R2-ENGINE-777 High reader, Medium delivered row; R2-ENGINE-884 Medium creature composition | Creature row values and complete writer population |
| Flee chosen | Teleport (23) admitted by L2.01518 (in the book, mana at least its cost): cast at a cell scanRange + 1 from the hostiles' mean point; otherwise walk 3 cells away from that point. Each spell routine also holds an Invisibility (12) arm and a walk-to-cell arm behind a repeated Teleport admission; the callers' entry test makes them unreachable while the admission is pure (Medium). No living hostile: idle planner. Back-off is the same with radius 2 and scanRange - 1 (R2-ENGINE-1418). | R2-ENGINE-506 High flee body, Medium centre/away interpretation | Spell admission and concrete destination |
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
and its cast pays no mana (R2-ENGINE-1441, Medium).

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
(R2-ENGINE-1767, High / Medium; the credit is Medium).

A unit with +0x54 = 0x10 or with health 0 or below restores nothing; below 0
the refill takes 1 health on one pass in four, and corpse ageing lowers
health separately, so the total loss per pass is not this figure alone.
Corpse ageing takes 1 health on every second world pass (Medium for the Session counter identity; a third caller ages units outside the world pass); a corpse of class
0x21..0x24 stops at -61 with kind 4 and stays, and any other corpse reaches
kind 5 below -600 and is released (R2-ENGINE-1765, High / Medium). The
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

From the Shield arm of the support choice to the end of the recovery draw, a Shield cast at the caster draws once on the
read path in RU, the recovery draw of the step that calls the apply; the bodies of the cast, the Shield arm and the
Effect constructor hold no draw call, and the single-target object skips the hit hook for a Defensive spell
(R2-ENGINE-1577, R2-ENGINE-1578, Medium; callers, list helpers, callees and virtual calls unread; the EN twins of
`L2.00031`, `L2.02072`, `L2.08110` and `L2.02074` are unread). A Poison Cloud Effect applied again to a unit that
carries it ends with remaining duration 128, the maximum of the standing value and 128, never a sum (R2-ENGINE-1575,
R2-ENGINE-1576, Medium; RU).

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

The bounded summoned-unit departure census retains the count pass and the
selected transition cleanup; caster loss, timers, general defeat fate and
LOAD outcome remain Unknown. R2-ENGINE-1256 (Medium bounded composition).

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

The bodies of the numbers no installed map uses are read in R2-ENGINE-1450 (amended), except action 37, whose cell is a message-formatting arm (Unknown): action 11 removes the item from U+0x7c into V+0x7c (R2-ENGINE-1699), action 27 runs cell move and Guard setter for nonzero owner+0x2c (R2-ENGINE-1700), and action 31 sets Outpost respawn countdown +0xb4 by group number +0xac (R2-ENGINE-1701, amended; p1 goes into +0xb4 of every Outpost whose group number equals p0, for every owner and with no loop exit; the filter passes only the Outpost in the classes with a record, EN Shop and Inn slot 14 unread, Medium: R2-ENGINE-1820; p1 counts container passes without the respawn scale, Medium: R2-ENGINE-1821). Unconscious units stay in their group list while dead units (+0x54 = 0x10, body or no body) leave on the kill walker; checks 1, 5 and 8 count every unit still in its group list (living and unconscious units, not dead units; EN High, RU Medium for checks 1 and 8), check 20 counts Buildings rather than units (R2-ENGINE-1452 amended, R2-ENGINE-1696), and the check loop skips a check whose copied IR+0x58 byte is nonzero (High; that the 8000..9999 mark reaches that byte is Medium; R2-ENGINE-1451 amended, R2-ENGINE-1697).

Script action 6 gives a group one of eleven method calls by its selector p0 (1, 2, 3, 4, 5, 10, 11, 14, 15, 17, 18); the method bodies are read in EN and RU, the other selectors do nothing, and the arms set member orders 0xb, 0xc, 1, 3, 8, 0xa and 0x11 and the group programme value; what these orders and programme values do in play is Medium (R2-ENGINE-1726, amending R2-ENGINE-1314).

The roles of the Session DWORDs that script actions 4 and 5 write are the win counter (+0xbe0c) and the failure reason (+0xbe14) that the result tick reads, sent as victory when the counter is nonzero and as defeat when the reason is above 0 (Medium as the player sees them; R2-ENGINE-1729, amending R2-ENGINE-1313).

Script check 2 reads x and y as bytes +0 and +1 of the object at unit +0x10 and is an inclusive rectangle test on them (High); checks 3 and 27 read the same getters (Medium) (R2-ENGINE-1730, amending R2-ENGINE-1450).

Script action 25 writes one cell event on the single cell (low8 p2, low8 p3), not a rectangle: event byte low8 p0, power low8 p1, and source and target pairs that are (p2, p3) twice when p4 is 0, else (p4, p5) and (p6, p7); the write replaces an earlier event of that cell (R2-ENGINE-1775, R2-ENGINE-1776, amending R2-ENGINE-1730 and R2-ENGINE-1314). With event 23, a unit whose step ends at the centre of that cell is passed to the cell move, which puts it at the centre of the target cell unless the move test refuses it (Medium for units outside the three step callers; nothing observed in the original), and occupant kinds 1 and 2 also keep the target in command word +0xb0 for one retry by the executor (R2-ENGINE-1778). The power byte and the source pair are unread for event 23 (Medium; R2-ENGINE-1779), and no store to a written event was found in the EN range of R2-ENGINE-1777 (Medium; whole-entry stores and RU bodies other than the empty test not searched). Mission 53 writes such an event from trigger 25 (R2-ASSET-386); both roots ship nine action 25 nodes, all with event 23 (R2-ASSET-387).

EN `L2.06918` stores the speed, heading, the per-advance x and y bytes and the advance count into the block at unit `+0x1c0` (R2-ENGINE-1470). In EN `L2.03582..L2.06966` it has one call (the step start of R2-ENGINE-1213); that it is the only writer of the step values is Medium. The count is 256 divided by the larger axis byte rounded up (Medium: the absolute-value helper `L2.04968` is unread). For a unit whose kind (`L2.03577`) is 1 the speed comes from the base speed, the multiplier, the slope between the two cells (Medium: the byte read as a height, R2-ENGINE-1340) and the mean of two cell bytes (the sum kept in a byte); for any other kind it is the base speed. The result is clamped to 1..63 (R2-ENGINE-1471); no group term was read. The 18 path and terrain parameters are read from `World\Data\map.reg` with code defaults (R2-ENGINE-1472; tile classifier `L2.06958`/`L2.07473` indexes `+0x54146..+0x54150` and writes no height, R2-ENGINE-1665; client/server type-3 readers in R2-ENGINE-1667); the full tables are on the route page.

A step changes the cell on advance ceil(128/v) when its y vector component (x when it has none) is positive and floor(128/-v)+1 when negative, from fine bytes 0x80 and a heading that already matches, advance 1 being the step start tick in that case (R2-ENGINE-1760, Medium for the premises). The static route is rebuilt on a new goal or after 17 repaths, the local path when empty or after 33 walk calls, unless the wait decision of R2-ENGINE-1762 returns 1 or 2, and after a re-plan (R2-ENGINE-1761), with the shipped `map.reg` values of R2-ASSET-383. At a held next cell a one-cell unit turns and waits for a walking occupant or a standing one about to leave, else it re-plans (R2-ENGINE-1762).

A map building caster lists the units within its record radius of its cell, Chebyshev (Medium), 1 to 9 cells, with no range compare against the spell (R2-ENGINE-1560). It starts a cast once per 16 session ticks through a built Unit that casts from state 13 with wind-up 1 and recovery 1 (R2-ENGINE-1561 amended, R2-ENGINE-1562); the built Unit casts once and is deleted when its recovery ends (Medium; R2-ENGINE-1740). It stops when the resolver is called with a blow or Blizzard value not below the building's health; no other route is shown to stop it (R2-ENGINE-1563). The Delivery 2 spell rows and Effect Speeds are in R2-ASSET-333.

## Computer-unit gaps

A struck ROM2 unit rewrites six command fields and posts a 0x54-byte record to the session list. The look-around is the one reader of command +0x54 found in the census (Medium), and no reader of a record was found in the chain that destroys the list (Medium) (R2-ENGINE-1491, R2-ENGINE-1492, R2-ENGINE-1262, R2-ENGINE-1451). Guard B reads the hit stamps in session ticks and in blocks of 16 ticks, and Guard A reads none of the hit fields (R2-ENGINE-1493, R2-ENGINE-1494). No body read chooses a point for an area spell: the aim is the victim unit or the victim's cell (R2-ENGINE-1490). Two scorers rank candidates, and of the effect words they read only bit 0x40000 of +0x144; an effect that changes the Mind word, the range or the virtual kind is not excluded (Medium) (R2-ENGINE-1495, R2-ENGINE-1496). The group assignment targets every member of the group with no idle test (R2-ENGINE-1497). A flee without Teleport walks, and an empty route ends in a stop and order 0xc at the unit's cell (R2-ENGINE-1498). A withdraw message sets v = 0, 10 or 30 for the owner object that message word +5 resolves to (that the word names a sender is Medium), and the outer message type, the sender and the control are Unknown (R2-ENGINE-1499, R2-ENGINE-1421). The numbers are in the area hit numbers table of the behavior page.

## Player orders in play

In the RU body, at Medium for play, a player's guard sets the group radius of R2-ENGINE-1621 on one path; whether that path or the unit-level order 0xb governs in play is Unknown (R2-ENGINE-1620, R2-ENGINE-1622). A pick-up order walks to a ring-searched cell; what takes the item is Unknown (R2-ENGINE-1183, R2-ENGINE-1625, R2-ENGINE-1626). Orders 0xc and 0x17 stand until another write in the bodies read (R2-ENGINE-1627). An idle mage casts by the arm table (R2-ENGINE-1628, R2-ENGINE-1629) and the look-around draws a heading (R2-ENGINE-1630). A patrol unit stays and retries at an unreachable point (Medium composition; R2-ENGINE-1624). A Fire Ball is applied once per covered cell of its square (Medium; R2-ENGINE-1537).

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

The route search (RU `L2.05174`, EN `L2.05175`) floods its word grid at +0x30000 + 2 × key from the unit's cell. With flag 1 (walk and approach) it tests plane +0x10000, with flag 0 (repath) plane +0x20000, and it skips a neighbour whose plane byte AND the unit's mask byte is nonzero; for unit kind 1 a step costs the terrain byte straight and one and a half times it diagonal, for other kinds 2 and 3 (R2-ENGINE-1804; the flag roles are Medium, R2-ENGINE-871). The grid helpers are read: `L2.05185`(x, y) and `L2.05180`(key) address the same grid word for one `this`, `L2.05230` the terrain byte and `L2.02402` returns unit virtual +0x20 (R2-ENGINE-1804, amending R2-ENGINE-871, R2-ENGINE-878 and R2-ENGINE-879). With the plane values the terrain build stores (R2-ENGINE-1340), for flag 1 and a unit of size 1, value 1 closes kind 1, value 5 kinds 1 and 2, and the border ring's 0x1f all three kinds; in the route-search workers and helpers read, a terrain byte of 255 enters only kind 1's step sum and no test closes a cell on it (R2-ENGINE-1805, Medium; the outcome in play is Unknown). For a larger unit the footprint loop is RU `L2.08532` (EN `L2.08533`), which calls the AND of two bytes RU `L2.03520` (EN `L2.02720`) (R2-ENGINE-1804, narrowing the footprint tester naming of R2-ENGINE-1667). The EN location routine `L2.07308` keeps `(y << 8) + x`, which has the form of the plane key (R2-ENGINE-1802; EN body read, Medium that the word is the plane key, RU body unread).
