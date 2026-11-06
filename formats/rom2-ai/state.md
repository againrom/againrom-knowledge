# Fields and SAV coordinates

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
| byte +0x20 | Programme selector; same byte in SAV programme block. | R2-ENGINE-882 High dispatch; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | All external setters |
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
| command word +0xa; +0xc; byte +0x14; byte +0x15 | Prepared goal, action target, range, phase; SAV same block positions, with target key fixups open. | R2-ENGINE-505 High RU local writes, Medium helper meanings; R2-ENGINE-543 High phase; R2-ENGINE-823 Medium EN transfer; RU coordinates by equal layout | Complete target resolution |
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
