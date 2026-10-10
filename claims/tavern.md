# TAVERN — mercenary hire

Claims about the inn's hiring hall. The inn is a town building, not a file
format. Its other half, the mission-giver (`InnMission` / `InnNPC`, EXP-0060),
is covered by `REG-SCN-064`; the two halves share the view and the record and
nothing else. Format of this file: [registry.md](registry.md).

## Terms

The terms follow the engine's own indexing.

- A type is the 1-based subscript that `byte [unit + 0x15b]` (client) and
  `byte [unit + 0x14c]` (server) carry. It is also the `[npc<t>]` section number
  in `scenario.res::npc.reg` and the index into `[General] MercenaryCount`.
  There are fifteen types.
- A pool is how many units of one type exist.
- A shelf is what a given mission offers.

## Types, shelf, hire, price and level

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MERC-TYPE-001 | A mercenary is a `CUnit` carrying a 1..15 type id, and the type is the index everything else is keyed by. | High | ● active (amended) | [EXP-0062](../experiments/EXP-0062-mercenaries/) |
| MERC-SHELF-002 | The tavern offers `[Mission<n>] Mercenaries` ∩ a permanent unlock list ∩ a non-empty pool: three sets under three separate keys. | High / Medium / Unknown | ● active (partially retracted) | [EXP-0062](../experiments/EXP-0062-mercenaries/) |
| MERC-HIRE-003 | A hire is one flag per type, and it takes the whole squad of that type, not one man. | High | ● active | [EXP-0062](../experiments/EXP-0062-mercenaries/) |
| MERC-PRICE-004 | `cost(type, mission) = (PriceA + n · PriceB) × unitPrice(mission)`, and `PriceA`/`PriceB` are the `npc.reg` fields whose use `SHOP-NPC-012` left Unknown. | High / Medium | ● active | [EXP-0062](../experiments/EXP-0062-mercenaries/) |
| MERC-LEVEL-005 | A mercenary's level is a step function of the mission number alone, and it selects a different `Data.bin` template rather than scaling one. | High | ● active | [EXP-0062](../experiments/EXP-0062-mercenaries/) |
| MERC-CMD-007 | The tavern is a simulation class the image names, `Tavern`, driven by command `0x37` (rebuild the roster) and command `0x38` (spawn and pay for the hire). | High / Unknown | ● active | [EXP-0062](../experiments/EXP-0062-mercenaries/) |

### MERC-TYPE-001

- `R0746` builds the inn's list from the campaign document's
  `CMapPtrToPtr` (`[campaignScreen+0xd0] + 0x9bc`, size `+0x9c0`, count
  `+0x9c4`). It admits a value iff `CObject::IsKindOf` returns true for the
  `CRuntimeClass` at `L00620`, whose fields decode as name `L10118` =
  `"CUnit"`, object size `0x1b0`, schema `0xffff`, no `CreateObject`.
- The type id is `byte [unit + 0x15b]`. Three routines use it as a subscript:
  `word [record+0x60 + (t−1)·2]` at `L07506`, the same at `L07507`, and
  `dword [record+0x88 + (t−1)·4]` at `L07508`.
- On the server the same id is `byte [obj + 0x14c]`, written at `L09214` from
  the spawn loop's own counter.
- Fifteen is the loop bound of the roster builder `R1555`: `t = 1..2`
  builds `Unit` objects from the literals `"Catapult"` and `"Ballista"`
  (`L10119`, `L10120`), and `t = 3..15` builds `Human` objects from
  `"NPC%02d_%d"`. It matches the 15 elements `[General] MercenaryCount` ships.
- The inn draws each type with `graphics\interface\inn\Unit%d\sprites.16a` keyed
  by the same byte (`L10121`).

**Confidence.** High. The class is read off its own `CRuntimeClass` record, and
the three subscript computations are named instructions on the repaired table.
`refto:L00620` returns 13 hits in 12 owners, 0 in orphan or undisassembled code.

**Amended.** `TAVERN-TALKPIC-016` (EXP-0400) corrects this claim's attribution
of `L10122`: that address is the talk loop's non-Hero branch.

### MERC-SHELF-002

- `R1412`, the inn view's `vt+0x80`, walks `Mercenaries`
  (`campaignScreen+0x5e4`/`+0x5e8` = `record+0x9c`/`+0xa0`). It is the only
  reader of `Mercenaries` on an image-wide `EnumRefs disp:`. It keeps an element
  only if `R1606` finds it in the array at `record+0xac`
  (`m_pData +0xb0`, `m_nSize +0xb4`). `R0746` then drops any type whose
  pool is 0.
- Only `R1797`, an append, fills the `+0xac` list. Its two callers are
  vtable slot 0 of the two mission-record classes: `R1798` (main record,
  vptr `L10123`) and `R1799` (sub-mission record, vptr `L07494`).
  Each drains its own record's `[Mission<n>] EnableMercenary` (`record+0x30`,
  `REG-SCN-059`'s destination) into the main record's list, then
  `SetSize(0,−1)`s it. An unlock is consumed once and never cleared.
- The chooser `R1669` has exactly one caller, `L10124`, inside the
  mission-end arm of `R0701`. It picks the record whose mission number
  equals `record+0x118`, the accepted mission. `EnableMercenary` is therefore a
  completion reward.
- Corpus, all three roots (`evidence/offers.csv`): 13 values over 5 sections.
  `[Mission10]` unlocks nine types; side missions 41, 71, 111 and 121 unlock one
  each.
- Four types are named by a `Mercenaries` key before they are unlocked: type 10
  named at `Mission40` and unlocked by side mission 41, type 8 at 70 / 71, type
  1 at 10 (and 20, and 110) / 111, type 5 at 120 / 121.
- Being named is not being shown: `R1412` applies the filter before it
  builds the list, so the tavern draws nothing for a type not yet unlocked. On
  the main-mission ladder in ascending order, the first shelf that can show each
  of the four is 50, 80, 120, 130 respectively (`evidence/checks.txt` C8, all
  three roots). A consumer keeps "named by this mission" and "hireable in this
  mission's town" apart.

**Confidence.** High for the three sets: every step is a named instruction.
`callto:R1797` returns 2 hits and `callto:R1669` 1, both on the repaired table
with `.rdata` slots included. Both writers are reachable only through a vtable,
so an `imm:`-shaped sweep would report zero callers for each. Medium that no
fourth set narrows the shelf: the reader enumeration is `disp:` over the array
object and its data and size fields, image-wide, with 0 orphan hits, but it
cannot see a pointer laundered out of the record. Unknown when within a chapter
a side mission's unlock starts to bite.

**Unknown.** The first-shelf figures walk the main-mission ladder in ascending
order, a reading of the data and not of a routine. A side mission completed
during its own chapter's town visit would apply its `EnableMercenary` at
`L10124` with the record still on that chapter. Whether the record advances at
that point turns on `R1414`, which `REG-SCN-065` describes as deleting
the finished sub-mission and `SHOP-TOWN-022` glosses as advancing the main
mission. One measurement narrows this without settling it. The other mission-end
arm, `L10125`'s `R1414` + `R0785`, is dead: its guard reads
`[L10126]`, a global with 2 references image-wide, both reads, both in
`R0701`, and no writer. Its static value is 10 against a
signed less-or-equal-to-10 test, which therefore always jumps. Every mission end takes
`R1411` at `L07487`. This leaves `SHOP-TOWN-022`'s reading of the
reachable site at `L09711` untouched.

**Amended.** The clause "Four types appear on a shelf one main mission before
the side mission that unlocks them, so the intersection is observable in play"
is retracted ([`retracted.md`](retracted.md), EXP-0062). The first missions at
which the four can appear are 50 / 80 / 120 / 130, not 40 / 70 / 10 / 120.

### MERC-HIRE-003

- `R1442` (hire) and `R1443` (dismiss) are mirror routines over
  `dword [record+0x88 + (t−1)·4]`. Each is gated on `R1606` and has
  exactly one caller (`R1441`, `R1440`).
- The flag's only other writers are the record reset `R0757`
  (zero-fill 15) and `R1605` (zero-fill 15, `MERC-DEATH-006`).
- Nothing is committed inside the inn. Leaving it runs `R0786`, which
  calls `R1653`. That routine culls every non-hero `CUnit` from the
  document map and sends command `0x38` carrying `R1654`'s vector
  `out[i] = working[i] × hired[i]` (a 16-bit multiply by the word at `hired + i·4`, at
  `L10127`): the entire surviving pool of each hired type.
- The server's `R1142` then runs `for (i = 1; i <= n; i++)`, creating one
  object per head.
- The client's affordability gate is `R1556`. `record+0x138` is the local
  player's money (`[[view+0x68]+0x9b4]+0x0c`) minus the price of every type
  already hired, and `R1441` refuses a hire whose price exceeds it. The
  check is local only. The money moves once, on the server, by
  `R0449(−Tavern+0x9c, 0)` at `L04977`.

**Confidence.** High. Both flag routines, the vector's `IMUL`, the spawn loop's
bound and the single debit are read at instruction level. `callto:` was run on
all five routines, `.rdata` slots included.

### MERC-PRICE-004

- Server (`R1142`, `L07464`…`L10128`): `PriceA` =
  `R1800(R1285(L02112, t))` = `[npc + 0x18]` and `PriceB` =
  `R1801(…)` = `[npc + 0x1c]`, the two fields `R0499` writes
  (`SHOP-NPC-012`). `n` is the hired count for that type, and the third factor
  is `R0589`.
- Client preview (`R1557`, `L10129`…`L10130`):
  `(byte[unit+0x155] · avail + byte[unit+0x156]) × (i16)[unit+0x158]`.
- The mapping between the two is witnessed on both sides of the wire, not
  inferred from the arithmetic. `R0059` emits `{type, PriceA, PriceB}` in
  that order under mask `0x40000000` (`L10131`, `L07833`, `L07834`), and
  `R0509` receives them into `{+0x15b, +0x156, +0x155}` (`L10132`,
  `L10133`, `L10134`). So `+0x156` is `PriceA` and `+0x155` is `PriceB`.
- `+0x158` is not sent: the client runs the ladder itself at `L03283`.
- `unitPrice` is a jump table at `L10135` on `mission/10 − 3`, arms
  `10 15 20 40 60 80 100 600 800 1000 6000 8000 10000`, everything outside
  `0..12` → 0. Its first argument, the type id, is pushed and read by no
  instruction, so the factor is per chapter and not per mercenary.
- Corpus, all three roots identical (`evidence/types.csv`,
  `evidence/costs.csv`): `PriceA` 28..80 and `PriceB` 0..15, present on exactly
  the 13 types whose pool is nonzero. The five pool-1 types all carry
  `PriceB = 0`, a flat price where the per-head term cannot bite.

**Confidence.** High for both formulas, read at instruction level, and for the
wire mapping, witnessed by the sender's and the receiver's own field order.
`callto:R1800` and `callto:R1801` return 2 hits each on the repaired table,
and the `refto:L00620`-style orphan check is clean. Medium that this is the only
price path: `R0589` short-circuits to a flat 450 when
`[L03937] == 2`.

**Unknown.** Which mode writes 2 to `[L03937]`. The global has 5 writers,
all in the `L09260` app/menu family, and none was read.

### MERC-LEVEL-005

- `R1802(mission)` computes `k = mission/10 − 1`, rejects `k > 14`,
  indexes the 15-byte arm table at `L10136`
  (`00 00 00 00 00 01 01 01 01 02 02 02 03 03 03`) and jumps through `L10137`
  to four arms returning 1, 2, 3, 4.
- Missions 10–50 give level 1, 60–90 level 2, 100–120 level 3 and 130–150
  level 4.
- Both consumers build a class name from it: `R1142` at
  `L10138`/`L10139` and `R0645` at `L10140`. Each formats
  `"NPC%02d_%d"` (`L10141` / `L10142`) and hands it to `R0497`
  (`new Human(0x1e8)`).
- Types 1 and 2 take no level: they are `Unit(0x198)` objects named by the
  literals `"Catapult"` / `"Ballista"`.
- A level is therefore a different character record, not a levelled-up one.
- Census over `world.res:data/data.bin`, all three roots: all 33 reachable
  `NPC<t>_<lvl>` names exist, and so do the 19 unreachable ones plus `Catapult`
  and `Ballista`, 60/60. The reading requires this, because `R1555`
  builds every one of the fifteen types unconditionally at whatever level the
  town is at.

**Confidence.** High. The ladder is transcribed from its own arm table and jump
table, and both consumers' `sprintf` sites are named. The `Data.bin` census is a
whole-string count over the length-prefixed names on 3/3 roots, and the reading
predicted it before the file was opened.

### MERC-CMD-007

- `Tavern` is `L10143`-vtabled, `0xa0` bytes, `: Building : Token`, in the
  `.data` `CRuntimeClass` table (`EXP-0057/evidence/runtime-classes.csv`). The
  single-player instance is the static `L08104`, the counterpart of
  `SHOP-ENTRY-003`'s `L09667`.
- Command `0x37` (`R0645`, opcode written at `L09231`, payload
  `dword cmd+0x0a` = `R1802(mission)`) reaches
  `R1555(player, level)`. That routine destroys the previous roster,
  refunds `Tavern+0x9c` in full, and rebuilds one object for every type 1..15 at
  that level.
- Command `0x38` (`R1653`, `L09233`; `word cmd+0x0e` = the mission
  number, `dword cmd+0x0a` = element count, payload from `L10144`'s `memcpy`)
  reaches `R1142(player, vector)`. That routine reads element 0 as the
  mission, `RemoveAt(0,1)`, spawns `n` of each remaining type and debits the
  accumulated cost once.
- Both arms sit in `R0061` at `L08101` / `L08145`, and both call the
  static instance.

**Confidence.** High. Both senders' opcode stores and payload writes, and both
dispatcher arms, are read at instruction level, and the payload layouts match
field for field between sender and handler. Unknown for a map-placed `Tavern`.

**Unknown.** Whether a map-placed `Tavern` is reachable in play. Only the static
instance was followed, and no analogue of `SHOP-ENTRY-016`'s `server+0x14c` test
was looked for on this class.

## The pool across missions

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MERC-DEATH-006 | At every mission end a hired type's pool becomes its live tally, an idle type regains one head up to `MercenaryCount`, and every hire flag is cleared. | High | ● active (amended) | [EXP-0062](../experiments/EXP-0062-mercenaries/) |
| MERC-POOL-011 | The shelf filter and the inn's price preview read one location, the working-pool `CWordArray` at campaign-record `+0x5c`; the server's own `n` is a different storage. | High / Unknown | ● active | [EXP-0395](../experiments/EXP-0395-mercenary-pool-persistence/) |
| MERC-POOL-012 | Nothing inside the tavern writes the pool; its four element writers and three storage routines are reached only from record construction, new-campaign reset, mission end, LOAD and record destruction. | High / Medium | ● active | [EXP-0395](../experiments/EXP-0395-mercenary-pool-persistence/) |

### MERC-DEATH-006

- Death and recovery are an explicit restore, not an absent write, and they are
  two different mechanisms.
- At the end of every mission `R0701` calls `R0809` (`L07290`),
  whose first act is `record->R1374()` (`L10145`; `callto:R1374` = 1
  hit).
- `R1374` collects the local player's live units through `R1373`
  (owner `[unit+0x14] == doc+0x9b4`, corpse stage `byte[unit+0x15a] == 0`) and
  tallies them per type into a 15-element scratch array (`L10146`…`L10147`).
- It then merges per type (`L10148`…`L10149`):
  - a hired type gets `working[t] = the live tally`, so every man that died is
    gone from the pool for good;
  - a type not hired gets `working[t] += 1`, and only while
    `working[t] < MercenaryCount[t]` (an unsigned 16-bit compare with `word[record+0x74 + i]` that skips the increment when not below).
- Then `R1605` zeroes all fifteen hire flags (`L10150`, its only
  caller), so no hire survives a mission.
- The tally runs before `R0809`'s own cull. The cull keeps only `CUnit`s
  with `+0x18c & 1` (player character; `PARTY-FLAG-003`, `PARTY-CULL-004`),
  `(i16)[+0xfc] > −10` and the player's owner id, and drops every mercenary from
  the map.
- A consumer must implement the consequence: a squad wiped on a mission takes as
  many further missions to rebuild as it lost men, and only while it is left at
  home.

**Confidence.** High. The merge is six instructions on one branch. Both in-play
writers of the pool's elements are enumerated, this routine and the record
reset, and the cap is a read of the pristine `MercenaryCount` array. `callto:`
on both routines returns 1 hit each, `.rdata` slots included.

**Amended.** EXP-0077 corrected the cull's gloss of `+0x18c & 1` from hero to
player character (`PARTY-FLAG-003`, `PARTY-CULL-004`;
[`retracted.md`](retracted.md)). `MERC-POOL-012` (EXP-0395) narrows the writer
enumeration: it is complete for the in-play paths and not for the document path.
The campaign reader `R0434` writes the same elements directly at
`L08025` and resizes the array at `L08024`, so "both writers" is not the
whole writer population of the storage ([`retracted.md`](retracted.md)). The
merge, its arms, its cap and the flag clearing are unaffected.

### MERC-POOL-011

- `MERC-SHELF-002`'s collector `R0746` and `MERC-PRICE-004`'s client
  preview `R1557` each reach the campaign object through
  the call to `R0347` and its `vt+0x7c`. Each then subscripts the same `m_pData`
  at `campaignObject+0x5a8` = `record+0x60` by `type − 1`:
  - `L10151` loads the 32-bit pointer at offset 0x5a8, then
    `L07507` compares the 16-bit element at index `type − 1` with zero and branches on unsigned below-or-equal, which drops a
    type whose unsigned word is zero;
  - `L10152` loads the 32-bit pointer at offset 0x5a8, then
    `L07506` reads the 16-bit element at index `type − 1` and
    `L10153` stores it as a 32-bit value into the local slot, storing exactly the value
    that `L10154` (a 32-bit multiply by that slot) multiplies by `PriceB`.
- `MERC-PRICE-004`'s client `n` is therefore the working pool. `R1557`
  also reads the pristine array's `m_pData` at `+0x5bc` (`L10155`, `L10156`)
  into a second out-parameter: the inn holds both the current and the full
  headcount and prices on the first.
- Image-wide census over the campaign-object-relative form of all three record
  collections: `disp:5a4` 0 hits; `disp:5a8` exactly two, both reads, the two
  above; `disp:5ac` 0; `disp:5b8` 0; `disp:5bc` 1 read; `disp:5c0` 0. The
  bare-immediate form of all six is 0 instructions. The 739 `imm:` hits over the
  same six are code-region-shaped constants and code-region-shaped string
  addresses that the prefix-matching mode returns, inspected as such.
- The server's price operand is a different storage. `R1142` runs
  `L08142` calls `R0130` (`ElementAt` on its argument array),
  `L10157` loads the 16-bit element, and `L10158` multiplies it by the 32-bit local at frame offset -0x24:
  a stack `CWordArray` built from command `0x38`'s payload, whose elements are
  `working[i] × hired[i]` (`MERC-HIRE-003`). For a hired type the two agree
  numerically and are two locations.
- Width: `PARTY-MERC-007` carries the pool across a mission boundary as "fifteen
  integers" and `SAV-629` quotes it as a "fifteen-integer mercenary pool-count".
  It is fifteen u16 words at `+0x60`; the fifteen dwords at `+0x88` beside it
  are the hire flags. Both rows cite `MERC-DEATH-006`'s merge, so both mean this
  array; only the word width is corrected, and neither row's finding changes.
- Evidence files: `evidence/enum-abs-release.txt`,
  `evidence/d-pool-release.txt`, `evidence/corpus-searches.txt`. Also cites
  `REG-SCN-064`.

**Confidence.** High. A closed census, not plausibility, excludes two separate
storages for the two readers: the two `disp:5a8` hits are the image's only two,
and neither offset has any immediate form at all. It also excludes the server
reading the client's location: `L08142` is `ElementAt` on a stack array, not a
campaign-object displacement.

**Unknown.** `MERC-PRICE-004`'s flat-450 mode, which this claim does not touch.

### MERC-POOL-012

- `R1442` (hire) and `R1443` (dismiss) were read whole. Each is
  under thirty instructions, is gated on `R1606`, and touches exactly one
  location: the 32-bit store of 1 at `L10159` and the 32-bit store of 0 at `L10160`, both to the address `[object+0x88] + (t−1)·4`.
  Neither references `+0x5c`, `+0x60` or `+0x64`, so a hire sets a flag and
  leaves the headcount alone until mission end.
- The four element writers:
  - the 16-bit element store at `L10161` in the record reset
    `R0757`, copying `working[i] = pristine[i]` over fifteen iterations
    (`L07504`…`L07505`) after `L07470` has read `[General] MercenaryCount`
    into the pristine array;
  - the 16-bit element stores at `L10162` and
    `L10163`, the hired and idle arms of the mission-end
    merge `R1374` (`MERC-DEATH-006`);
  - the indirect read call at `L08025` in the campaign reader `R0434`, a raw read of
    `2·count` bytes straight into `m_pData` (the byte count is pushed at `L10164`).
- The three storage routines:
  - the empty-array constructor `R1549` at `L09117`;
  - `SetSize` (`R1622`, which zero-fills a grown tail and frees and NULLs
    `m_pData` at size 0) at `L10165` with 15 and at `L08024` with the
    document's own count;
  - the teardown `L10166`, which the record destructor `R1803` calls
    on each pool object at `L10167` and `L10168`. That routine is a
    destructor by its own SEH frame, unwind-state stores, element-free loops and
    base-vptr stores. `L10166`'s own body was not read, so its effect on the
    storage is a role name from its call position.
- The events reaching those seven:
  - the constructor `R0677` has one call site, `L03281` (`SAV-600`);
  - the reset has three, `L09120`, `L08431` (`SESS-START-034`'s new-campaign
    arm) and `L03810`, whose own only caller is `R0701` at `L10169`;
    all three are new-campaign initialisation;
  - the merge is reached only from `L10145`, at mission end;
  - the reader is reached from `L08112` and `L08116`, the two campaign LOAD
    drivers (`SAV-CAMPTAIL-070`).
- `R1654`, which builds the wire vector on the way out of the inn, reads
  `[ESI + 0x60]` and `[ESI + 0x88]` and writes only its argument.
- Record-relative `+0x64` has exactly one object reference in the whole campaign
  family, the SAVE count read; the other eleven hits at that displacement inside
  `L07481..L07482` are `[ESP + 0x64]` stack locals. So neither tavern reader
  bounds-checks the subscript against `m_nSize`.
- Evidence files: `evidence/enum-rel-release.txt`, `evidence/range-release.txt`,
  `evidence/d-pool-release.txt`, `evidence/d-drivers-release.txt`,
  `evidence/d-reset-callers-release.txt`, `evidence/callers-release.txt`,
  `evidence/callers2-release.txt`. Also cites `MERC-HIRE-003` and `REG-SCN-067`.

**Confidence.** High for the four element writes, the three storage operations
and for hiring not debiting the pool. Both flag routines are short and were read
whole, so "no write" is a reading of the routine rather than an absence
argument; it excludes the live alternative that a hire debits immediately and
the mission-end merge only restores. Medium that there is no eighth writer. The
campaign-object-relative census is closed. The record-relative census was
crossed against all 35 record-base owners, whose 47 `+0x5c`/`+0x60`/`+0x64`
instructions split into 20 `[ESP + ...]` stack locals and 27 object references
read in place. Every one resolves to `[EDI+0xd0]` or `[EDI+0xec]`, a field
compared against the constant `0x445`, a halved dimension in a centring
computation, or an indirect call through a slot.

**Unknown.** A routine that receives the record pointer as a plain argument
while appearing in neither census; 1 093 bytes of `L07481..L07482` are
undisassembled.

## Roster cells, pictures, statistics and voice

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TAVERN-MERCVOICE-008 | The fourteen mercenary bios ship on all three roots, but 16 of 21 mercenary voice parts are EN/RU-only; where the pre-release root carries voice, its byte size matches EN's, not RU's. | High / Medium | ● active (amended) | [EXP-0393](../experiments/EXP-0393-inn-text-closure/) |
| TAVERN-ORDER-015 | The mercenary cells follow the walk of the campaign document's actor map, not any persisted list. | High / Medium | ● active | [EXP-0400](../experiments/EXP-0400-tavern-roster-presentation/) |
| TAVERN-TALKPIC-016 | A talk-only cell draws the sheet of the object `R0741(InnNPC[j])` returns: `HeroMage`/`HeroFighter` when `+0x18c` bit 0 is set, else `Unit<+0x15b>`; the selected cell animates it. | High / Medium | ● active | [EXP-0400](../experiments/EXP-0400-tavern-roster-presentation/) |
| TAVERN-TALKSTATS-017 | Selecting a talk-only cell shows that object's statistics in the upper left panel only when its `+0x18c` bit `0x40` is clear, and a synthesised object is created with it set. | High / Medium | ● active | [EXP-0400](../experiments/EXP-0400-tavern-roster-presentation/) |

### TAVERN-MERCVOICE-008

- The fourteen bios, `text/inn/mercenary/npc*.txt`, ship on all three roots: 0
  asymmetric.
- Their voice does not: 16 of 21 `inn/mercenary/*.wav` parts are EN/RU-only.
- `npc09p3.wav` and `npc13p3.wav` ship on EN with no RU counterpart: a
  text/voice split inside the same bio, EN richer.
- A fourth, owner-supplied pre-release root carries none of the 21 mercenary
  voice parts except `npc04p1.wav`, `npc10p1.wav` and `npc35p1.wav`. For those
  three its byte size equals the EN file's exactly (190188/179332/630792 bytes
  respectively), while RU's own recording of the same three differs
  (174496/371180/634548).
- This is the hiring hall's own text/voice census, distinct from the
  mission-giver half this ledger excludes (`DLG-INNVOICE-030`).

**Confidence.** High for the presence plus byte-size figures: direct per-root
reads over a named, closed set of 21+14 paths
(`evidence/dialogue-presence.csv`). Medium for reading the three size matches as
unlocalized or reused audio: consistent with a byte-identical payload, but only
size was compared, not a payload hash, so exact identity is not established.

**Amended.** One of the fourteen text files, `npc35`, is the town gate line, and one
of the 21 voice parts, `npc35p1.wav`, is its voice; the other thirteen text files
are type bios (`TAVERN-BUTTON-020`, `TOWN-475`). The presence and size figures
stand.

### TAVERN-ORDER-015

- `R0746` visits the map's buckets from index 0 and each bucket's chain
  from its head (`L10170`..`L10171`). It keeps a `CUnit` whose type is in
  the unlock-filtered `Mercenaries` list and whose working pool is non-zero
  (`L10172`..`L10173`), and appends in walk order. The inn view copies the
  list unchanged (`L08031`..`L10174`).
- The bucket is `((key & 0xffff) >> 4) % size` (`L10175`..`L10176`), and the
  client inserts at the bucket head (`L10177`..`L10178`). Units with
  consecutive ids inside one 16-id block therefore come out newest first.
- `Mercenaries`, the unlock list and the pools decide membership only.
- Owner screenshot of the original EN tavern for a mission-130 SAV
  (`Mercenaries` `14,6,10,13,4,8,7,3,1,9,2,12,5`): the cells read
  `10 9 8 7 6 5 4 3 2 1 14 13 12` from position 0. That is descending type in
  two runs, the pattern of stock units created in type order with the id block
  boundary between types 10 and 12.
- The owner's statistics panels tie the three 420,000 cells (types 3, 4, 5) to
  `NPC04_4`, `NPC03_4`, `NPC05_4` field for field. Under the seat's
  left-to-right reading this puts types 5, 4, 3 at positions 5, 6, 7, as the
  rule predicts.
- Evidence files: `evidence/listing.txt`, `evidence/model-scores.txt`,
  `evidence/roster-grid.csv`, `evidence/mage-templates.csv`. Cites
  `MERC-SHELF-002`, `MERC-LEVEL-005`, `SAV-928` (the same walk, for membership),
  `MERC-CMD-007`, `SAV-918`, `SAV-926` (command `0x37` at tavern entry rebuilds
  types 1..15 in fixed type order, so the ordering ids are assigned after any
  load).

**Confidence.** High that the order is the actor-map walk. Rivals excluded on
the same 14 cells: `Mercenaries` key order matches 2, ascending type 2,
descending type 1 (`evidence/model-scores.txt`). Medium for the id-block split:
the stock units' ids were not derived from code, so where a split falls is
fitted to one screenshot, and the mage tie rests on the seat's reading of the
owner's click order.

### TAVERN-TALKPIC-016

- Loader `L10179`..`L10180`. When `+0x18c` bit 0 is set, bit 1 picks
  `HeroMage` over `HeroFighter`.
- The object is a live `CUnit` passing the npc section's `Flags` terms
  (`DLG-SPEAKER-023`) or, when none passes, the synthesised object with
  `+0x18c = 0x48` plus the token bits and `+0x15b` = the npc id (`L03489`,
  `L10181`).
- Paint advances a cell's frame `(frame+1) % frameCount` only while its position
  equals the selection and more than 125 ms have passed
  (`L10182`..`L10183`); other cells hold their frame.
- A sheet that does not ship aborts: the `.16a` constructor formats
  `"FATAL ERROR: can't load "` + path, shows it and calls `abort`
  (`L03528`..`L03529`, `R0754`).
- Over the 22 `InnNPC` elements per root (EN = RU):
  - npc22 draws `HeroMage`, npc23 and npc25 `HeroFighter` on either arm;
  - on the synthesised arm npc2, 30, 32, 41, 52, 62 and 64 draw their shipped
    `Unit<id>`;
  - npc90 (Mission40) and npc59 (Mission70) have no `Unit<id>` sheet. They are
    the only Human candidates a stock mercenary template of the stage's level
    passes on sex, class and face (`NPC10_1`, `NPC08_2`), paired with side
    missions 41 and 71, which unlock types 10 and 8.
- The mission-130 SAV's candidate npc25 draws
  `graphics\interface\inn\HeroFighter\sprites.16a`, a shared sheet, not a
  per-character one.
- `L10122` is this loader's non-Hero talk branch; this corrects
  `MERC-TYPE-001`'s attribution of that address.
- Evidence files: `evidence/listing.txt`, `evidence/talk-candidates.csv`,
  `evidence/inn-sheets.csv`. Also cites `REG-NPC-088`, `DLG-SPEAKER-022`,
  `MERC-CMD-007`, `SAV-926` and `MERC-LEVEL-005` (the level-1 type-10 stock unit
  `NPC10_1` at Mission40).

**Confidence.** High for the rule and for the Hero candidates' sheets. Rival
excluded: a sheet keyed by the `InnNPC` value alone, which would give npc25 the
unshipped `Unit25` and abort. Medium that npc90 and npc59 resolve to a live
stock mercenary: `Platoon`, `+0x15a` and the walk order were not evaluated, and
no running original was observed at those stages.

### TAVERN-TALKSTATS-017

- `R1804` passes `view+0xec[sel − view+0xc8]` to `R0892`. That
  routine blits `LeftStats.bmp` and `LeftPicture.bmp`, and calls the shared
  character/unit panel `R0877` with rect `(12,0)-(172,238)` only when
  `+0x18c & 0x40` is clear (`L10184`..`L10185`).
- Below the panel it draws the object's composed figure through a temporary
  `allods%d.$$$` file when `+0x18c & 0x11` is set (`L10186`..`L10187`), else
  its `graphics\infowindow\<InfoPicture>.bmp` (`L10188`..`L10189`).
- The synthesiser stores `0x48` (`L03489`) and afterwards only ORs in token
  bits: `1` for `Hero`, `0x10` for `Human`, `4`, `2` and the hero-relative sex
  and class bits. None clears a bit.
- The client class setter keeps only bit `0x80` of the old value and sets `9`
  for heroes or `0x18` for human classes (`L10190`..`L07801`).
- A synthesised candidate therefore shows no statistics. Below the panel, one
  with `Hero` or `Human` shows its composed face figure. One with neither
  (`0x48 & 0x11 = 0`) takes the `infowindow` arm with the class name the object
  holds, and the synthesiser sets `+0x20` only under a `Picture` token.
- The one such shipped record is npc2 (Mission110, `Flags` `Platoon`), identical
  on both roots.
- For the mission-130 SAV the owner saw statistics for npc25's cell, so that
  cell holds a live actor passing `Hero,Face,!Female,!Mage`. When several pass,
  the first in the actor-map walk is taken. The owner identifies Brian.
- Evidence file: `evidence/listing.txt`. Cites `TEXT-UI-034`, `DLG-SPEAKER-022`,
  `HERO-APPEAR-053` (names the same two literals in `R0892`).

**Confidence.** High for the gate and the value at creation. Rival excluded:
statistics drawn from the npc record itself; no `npc.reg` field reaches
`R0877`, which takes the object. Medium that the synthesised object keeps
bit `0x40` for its lifetime. The nine immediate dword stores to `+0x18c` in the
linear listing are `0x48` (`L03489`), `0x29`/`0x2b` (`L02788`/`L07287`),
`0` (`L03487`, `L10191`) and four `0` stores (`L10192`, `L10193`,
`L10194`, `L10195`) set aside as another class's field: their routines zero
`+0x190` beside it, call an import on `this+0x3b0`, or pass the old value to an
import before zeroing it. Register stores such as the mask-with-0xf7 merge at
`L10196` (clears bit 3 only) were not all traced, nor other virtual slots or
pointers to the object. Medium also that no other path sets bit `0x40` on a live
actor.

**Unknown.** npc2's picture, and whether a live actor answers its `Platoon`
term.

## Roster and button clicks

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TAVERN-CLICK-019 | In the tavern a click on an occupied roster cell selects it and a double click hires or dismisses a mercenary or opens a talk dialogue; no routine read gives the candle, cauldron or tender a click test. | High / Medium | ● active (amended) | [EXP-0409](../experiments/EXP-0409-town-figures/) |
| TAVERN-BUTTON-020 | The three tavern buttons are rectangles that act on release over the pressed one: the upper hires or dismisses, the middle opens the bio or talk dialogue, the lower leaves for the town. | High | ● active | [EXP-0409](../experiments/EXP-0409-town-figures/) |
| TAVERN-FIGURE-021 | Without the tip popup, a press on the tavern's candle, cauldron or tender reaches the roster; only occupied-cell pixels select: none on the candle, 5456/15120 on the cauldron and 13680/38160 on the tender with 18 cells. | High / Medium | ● active | [EXP-0412](../experiments/EXP-0412-town-square/EXP-0412.md), `evidence/figure-presses.tsv`, `evidence/windows.tsv`, `evidence/d-tavern.txt`, `evidence/d-dispatch.txt`, `evidence/popup-presses.tsv`; extends TAVERN-CLICK-019, TOWN-467, TOWN-468 |
| TAVERN-LINES-022 | Read hire refusals request sound alone, and read training refusal branches post no message line. No Sleep action was found in the inspected room handlers, import callers or either shipped text corpus. | High / Medium | ● active | [EXP-0412](../experiments/EXP-0412-town-square/EXP-0412.md), `evidence/d-tavern.txt`, `evidence/d-school.txt`, `evidence/d-shop.txt`, `evidence/d-town.txt`, `evidence/poster-callers.tsv`, `evidence/api-summary.tsv`, `evidence/text-search.tsv`, `evidence/exe-ascii-search.tsv`, `evidence/code-anchors.tsv`; extends TAVERN-CLICK-019, TAVERN-BUTTON-020, MISSION-MSGPOST-058 |

### TAVERN-CLICK-019

- A click reaches the child that holds the point through the base dispatcher's
  children broadcast (`TOWN-139`, `TOWN-211`), and the view's own `+0x54`
  forwards to the same child (`TOWN-013`, `TOWN-014`). Of the three children
  (`TOWN-064`) the left panel's `+0x54` (`R1805`) returns 0, and the
  right-button slots `+0x60`, `+0x64` and `+0x68` are default stubs in all three
  (`evidence/refs.txt`, `evidence/d-tavern.txt`).
- Roster `+0x54`, `R1806`, runs the cell test `R1528`: `PtInRect` over the
  cell rectangles `i` below `view+0xc8` plus `view+0xf0`, so only occupied cells
  count, mercenary cells first (`TOWN-066`, `TOWN-467`, `TOWN-468`). A hit
  stores `i` at `view+0xb8`, refreshes the first button caption (`R0600`,
  `TOWN-391`), requests sound `+0xa8` (`R1807`) and returns 1. A miss returns
  0. The roster's left-up slot `+0x58` is a stub.
- Roster `+0x5c`, `R1437`, is the double-click slot (`TOWN-211`, `TOWN-406`).
  It runs `+0x54` first and returns 0 when that misses. On a hit, a selection
  below the mercenary count `view+0xc8` reads the type byte `unit+0x15b` and
  calls dismiss `R1440` when `[campaign+0x5d0][type-1]` is nonzero, hire
  `R1441` otherwise, then refreshes the caption. Hire is refused with sound
  `+0xa0` when the price exceeds the money left (`MERC-HIRE-003`). A selection
  at or above the count calls the talk routine `R0702`. Every hit returns 1.
- `R0702` formats `inn\NPC\npc%02dm%d` (literal `L07593`) from
  `InnNPC[idx]` and `InnMission[idx]` at `idx=sel-count`. When `InnMission[idx]`
  is 0 the second number is the live main mission. It opens the dialogue through
  `R0695`, queues a nonzero mission and requests sound `+0xb4` (`REG-118`,
  `REG-SCN-064`, `DLG-ZEROARM-029`). The text nodes are `main.res`
  `text/inn/npc/npc<NN>m<M>.txt`: 22 nodes and 17499 bytes on EN, 29 nodes and
  22758 bytes on RU (`evidence/node-families.tsv`, `evidence/nodes.tsv`).
- The candle (160,48)-(240,168), the cauldron (420,160)-(480,412) and the tender
  (240,152)-(420,364) are placed by painter addends (`TOWN-407`). No click slot
  read tests them. The candle overlaps no cell rectangle, the cauldron overlaps
  cells 11 and 17, and the tender overlaps cells 7 to 11 and 13 to 17. A click
  there reaches the cell test, which selects only an occupied cell
  (`evidence/rooms.tsv`). The tip popup that the tavern's enter builds at
  `L10197` over (0,0)-(312,200) while the tips option is on (`TOWN-015`)
  contains the candle whole and overlaps the tender; `evidence/rooms.tsv` has no
  popup row and the popup's controls were not read for mouse messages
  (`TOWN-207`).
- The hero figure in the right column is the shared character panel
  (`SHOP-FIGURE-041`, `TOWN-349`).

**Confidence.** High for the slot words, the cell test, the double-click routing
and the talk routine: each was read at instruction level, and both roots run one
program (`evidence/rom-program-digest.tsv`). Rival excluded: a hit test on the
candle, cauldron or tender inside the roster child, whose click slots are the
two above and stubs. Medium that no other routine acts on a click in the figure
rectangles: the tavern view's own message handler `R1808` (`vt+0x48`) and the
controls of the tip popup were not read for mouse messages. High that a double
click reaches `+0x5c`: the base dispatcher maps message `0x203` to it
(`TOWN-211`, `TOWN-406`), and the frame's window class requests it (`MENU-143`).

**Unknown.** The files bound to the sounds `+0x98` to `+0xb4`, the voice nodes of
the talk family, and what the tip popup's controls do with a left click.

**Amended.** `MENU-143` reads the window class style and the frame's `0x203`
arm; the double-click routing clause rises from Medium to High. `MENU-146`
states that the double click repeats the roster press before acting.

### TAVERN-BUTTON-020

- The buttons child (vtable `L08008`, rectangle (480,0)-(640,238), `TOWN-064`)
  has three rectangles that its initializer `R1809` stores at `+0x90`,
  `+0xa0` and `+0xb0`: index 0 (484,44)-(624,90), index 1 (484,91)-(624,137) and
  index 2 (484,138)-(624,184). They are the upper, middle and lower bands
  `TOWN-214` draws the buttons in (`evidence/d-capstone.txt`,
  `evidence/exe-anchors.tsv`). The hit test `R1810` runs `PtInRect` on them
  with the point relative to the parent origin and returns the first index or
  -1.
- Left down `+0x54`, `R1811`, stores the hit index at `+0xc0` (-1 on a miss),
  requests sound `parent+0xb0` for index 2 and returns 1 for any point. Double
  click `+0x5c`, `L10198`, repeats `+0x54`. The pointer slot `+0x4c`,
  `L10199`, only updates the highlight `+0xc4` (`L10200`) and returns 0.
- Left up `+0x58`, `R0703`, acts only when `+0xc0` is 0, 1 or 2 and the hit
  test at the release point returns the same index. It clears `+0xc0` and
  returns 1 on every path. The actions:
  - Index 0, upper: for a selection below the mercenary count it dismisses
    through `R1440` when `[campaign+0x5d0][type-1]` is nonzero and hires
    through `R1441` otherwise, then refreshes the caption. For a talk cell it
    does nothing.
  - Index 1, middle: for a selection at or above the count it calls `R0702`
    (`TAVERN-CLICK-019`). Otherwise it formats `inn\mercenary\npc%02d` (literal
    `L07592`) with the type byte `unit+0x15b`, opens it through `R0695` and
    requests sound `+0xb4`.
  - Index 2, lower: `L10201` sends `0x445` through the view's `vt+0x48` and
    posts `0x42e` to the window `[object+0x1c]`, which reopens the town view
    (`TOWN-477`).
- The captions are `main.txt` slot 258 or 259 for the upper button by hire state
  and empty for a talk cell, slot 242 for the middle and slot 232 for the lower
  (`TOWN-391`, `TOWN-383`).
- The bio nodes are `main.res` `text/inn/mercenary/npc<NN>.txt` for NN 01 to 10,
  12, 13 and 14: 13 nodes on each root (EN 3026 bytes, RU 2603 bytes). The
  family holds one more node, `npc35`, which is the town gate line (`TOWN-475`;
  14 nodes, EN 3249 bytes, RU 2836 bytes). No node exists for types 11 and 15 on
  either root (`evidence/node-families.tsv`, `evidence/nodes.tsv`). This narrows
  `TAVERN-MERCVOICE-008`, whose fourteen bios include `npc35`.

**Confidence.** High for the three rectangles, the hit test, the three slot
bodies and the action switch: each was read at instruction level, the
initializer's stores were checked byte for byte against the rectangle fields,
and both roots run one program (`evidence/rom-program-digest.tsv`). The index
order agrees with the caption order of `TOWN-391` and `TOWN-392`. Rival
excluded: an action on release without the pressed index, which `L10202` rules
out.

**Unknown.** The sound files bound to `parent+0xb0` and `+0xb4`, and what
`R0695` shows for a bio node that does not exist.

### TAVERN-FIGURE-021

The base figure routes and no-action results below assume the tip popup is absent and no capture child is set.

- Route. The campaign window procedure sends the press to the root view `[campaign+0xcc]`. The root's children broadcast gives it to the tavern view, whose message slot `R1808` forwards to the base handler `R0716` and the base dispatcher `R0390`. That slot (54 instructions, read whole) has arms only for messages `0x402`..`0x45a` and no mouse arm. The dispatcher's children broadcast gives the press to the roster child (id `0x450`, vtable `L10203`, (160,0)-(480,480)). The candle (160,48)-(240,168; 9600 px), the cauldron (420,160)-(480,412; 15120 px) and the tender (240,152)-(420,364; 38160 px) are the painter rectangles of `TAVERN-CLICK-019`. All three lie inside the roster's rectangle and overlap neither the left panel (0,0)-(160,480) nor the buttons child (480,0)-(640,238) (`evidence/windows.tsv`).
- Left-down. The roster's `+0x54`, `R1806`, runs the cell test `R1528` (`TAVERN-CLICK-019`). The 18 cells are 48x64 px in 6 columns (x 176 to 464) and 3 rows (y 288 to 480), numbered 0 to 17 from the bottom left, row by row. Only occupied cells count, in index order from 0. A hit selects the cell (`view+0xb8`), refreshes the first button's caption and requests sound `+0xa8`, and the routine returns 1. A miss returns 0. The tavern view's redelivery `R0790` then offers the press to the roster once more with the same result, and the press ends unconsumed (`TOWN-479`).
- Pixels pressed (`evidence/figure-presses.tsv`, identical on both roots). With no cell occupied every pixel of all three figures is unconsumed. With 18 cells occupied the candle's 9600 px are unconsumed because it overlaps no cell. The cauldron has 5456 px that select a cell (cell 17: 2816 px; cell 11: 2640 px) and 9664 unconsumed. The tender has 13680 px that select a cell (cells 14, 15 and 16: 3072 px each; cell 13: 2048; cells 8, 9 and 10: 576 each; cell 7: 384; cell 17: 256; cell 11: 48) and 24480 unconsumed. With fewer occupied cells only the cells below the count select: the cauldron selects nothing while the count is 11 or less, the tender nothing while it is 7 or less, and the candle never.
- Left-up. The roster's `+0x58` and the tavern view's `+0x58` are the stub `R1812`, which returns 0, so a release on any figure pixel is unconsumed and does nothing. A double click (`0x203`) reaches the roster's `+0x5c`, `R1437`, which runs `+0x54` first and returns 0 on a miss; on a selecting pixel it hires, dismisses or opens the talk dialogue as for the cell itself (`TAVERN-CLICK-019`).
- Player-visible result. A press on a figure outside an occupied cell posts no text, opens no dialogue, plays no sound and writes no state. On a cell's pixels it selects that cell exactly as a press on the cell does.
- With the tip popup shown (`TOWN-480`), the popup (160,0)-(472,200) lies over the candle, the top 48 rows of the tender and the top 40 rows of the cauldron. Its check box (200,160)-(348,176) consumes a press on 320 px of the candle and 1728 px of the tender and writes the tips flag `[L03631]` on down. Its Close button (352,160)-(432,178) consumes down on 1224 px of the tender and 216 px of the cauldron, registers the press and can post `0x45a` on release inside (`TOWN-480`). These are known actions outside every roster cell. On the other figure pixels under the popup the press passes to the roster, whose cells (y 288 and below) lie 0 px under the popup, so it ends unconsumed (`evidence/popup-presses.tsv`).

**Confidence.** High for the route, the cell test, the cell geometry, the pixel counts and the popup overlaps: each was read at instruction level or computed from committed anchors and archive sizes, and both roots run one program (`evidence/rom-program-digest.tsv`). For the popup-absent figure routes, alternative (a), a named window acting on the press, holds only through the roster's cell test. Alternative (b), a handler that does nothing, is the result on every other pixel. Alternative (c), another base route, was not found: the left panel's `+0x54` (`R1805`) returns 0 and the figures lie outside it. With the popup shown, the measured check-box and Close overlaps instead have the known actions above. Medium that no other routine acts on the press: the roster's children were enumerated from its constructor only, the figure rectangles are painter rectangles with no opaque-pixel mask, and the model assumes no capture child at the press. The occupancy states 0 and 18 were run; intermediate counts follow from the prefix rule of `R1528`.

**Unknown.** The occupancy a shipped save gives the roster at the moment of the press, and the pixels of the figures that are transparent in their frames.

### TAVERN-LINES-022

No message-line post was found in the tavern, school, shop or town routine bodies preserved for this experiment. The exported direct-reference census of the two message posters names 44 sites for `R0588`, in owners `R0819` (7), `R0509` (36) and `R0701` (1), and one site for `R1193` in `R0231` (`evidence/poster-callers.tsv`; `MISSION-MSGPOST-058`). Neither census names a room routine as a poster caller. This is a result over these bodies and direct references, not a proof that no computed code pointer can post elsewhere.

The hire refusal arm requests only the tavern sound slot `+0xa0`; hire success uses `+0x98` and dismissal uses `+0x9c`. No dismissal refusal arm was found in the read handler (`evidence/d-tavern.txt`). The two training refusal tests return through the same exit: purse after price below zero at `L10204`, and price at or below zero at `L10205`. Neither refusal branch itself posts text or requests sound (`evidence/d-school.txt`). A normal button's preceding presentation or input work is outside this narrower statement about refusal feedback. These paths provide no string-resource slot for a refusal line.

The tavern loader's other sound slots are `+0xa8` for the Helper.wav filename, `+0xb0` for Out.wav and `+0xb4` for Talk.wav. The preserved loader and consuming call sites establish the slot identities; the filename alone is not their authority (`evidence/d-tavern.txt`, `evidence/code-anchors.tsv`). These resolve the sound-file questions of `TAVERN-CLICK-019` and `TAVERN-BUTTON-020` for the read paths.

No Sleep action was found in the inspected tavern handlers. The imported `Sleep` function has five call sites in four owners, `R0075`, `L10206`, `L10207` and `L10208`, all in the read network population (`evidence/api-summary.tsv`). The executable ASCII-run scan searched 10044 runs and found four matching runs, including the import name; it did not identify a tavern action (`evidence/exe-ascii-search.tsv`). The sleep/rest/night-related text-node search covered 422 EN nodes and 434 RU nodes, with matches in 16 and 24 nodes respectively (`evidence/text-search.tsv`). Those matches include dialogue, event and item wording but supply no read tavern action route. The corpus search is corroboration, not authority to infer a handler from words. A line 'after Sleep' is consequently not identified by this evidence.

**Confidence.** High for the local hire and training refusal branches and the loader-to-slot associations: the preserved instructions distinguish a sound request from a message post and read the refusal tests themselves. Medium for absence of room message posts or a Sleep action beyond those paths: the direct-reference export misses computed pointers, and a bounded vocabulary/ASCII search cannot exclude an action named by other text or assembled dynamically. The recorded counts state the searched population. EN/RU instruction agreement is one program, not independent confirmation (`evidence/rom-program-digest.tsv`).

**Unknown.** Unread indirect message-poster callers, callers through code pointers outside the inspected tables, a differently named or dynamically composed action outside the read handlers, and the consumers of `0x472` and `0x444`.
## Pending mission identities

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TAVERN-023 | The inn appends each nonzero heard mission to a dword array without an identity comparison; its commit loop calls registration for every entry, including repeats. | High | ✔ promoted | [EXP-0415](../experiments/EXP-0415-dialogue-queue/) |
| TAVERN-024 | Repeated inn registration removes the first matching offer, sets one mission's announce flag and skips an existing journal identity; the offered-mission consumer reads flagged campaign records. | High / Unknown | ✔ promoted | [EXP-0415](../experiments/EXP-0415-dialogue-queue/) |

### TAVERN-023

`R0702` reads paired `InnMission`/`InnNPC` words. The nonzero mission arm
opens the dialogue at `L03607`, then passes the mission and current count
`view+0x118` to `L10209` with `this = view+0x110`. The zero arm opens dialogue
keyed by the live main mission and skips append (`REG-SCN-064`).

The embedded array holds data at `+0x114`, count at `+0x118`, capacity at
`+0x11c` and growth at `+0x120`. `L10209` compares the supplied index with
array `+8`; when needed it calls `L10210(index+1,-1)`, then stores the supplied
dword at `data[index]`. `L10210` allocates/grows, copies old dwords, zeroes
new storage and sets count. Neither compares stored mission identities.

`R0786` reads count and visits indices 0 through count-1. At `L10211` it
calls `R0785` on `campaign+0x548` for each stored dword. It does not compare
the current mission with an earlier entry and does not clear the array in
this loop. Its later UI teardown is outside the queue proof.

Original-code Unicorn vectors execute `L10209` and `L10210`, with explicit
allocation, free, zero-fill and overlap-safe copy hooks. The commit range
`L10212..L10213` runs with registration hooked to record arguments. Inputs
`31`, `31,31`, `31,41,31` and `0,31` produce exact stored and called sequences.
Zero is a container control, not a value the talk path appends.

**Confidence.** High for this append/container/commit path. The repeated and
mixed vectors distinguish retention from deduplication or replacement. The
callee hook proves commit calls, not registration effects; those use the
separate read in `TAVERN-024`. EN/RU have identical code, one observation.

**Unknown.** Native reentry/event ordering and aliases outside the named bodies.
No claim says the queue survives another inn visit or that repeated calls imply
repeated rewards, missions or visible journal entries.

### TAVERN-024

`R0785` calls `R1413`, `R1407`, `R1410` and `R1325` in that order.
`R1413` searches the u16 InnMission array at record `+0xd4` by identity and
removes the first match from both it and the paired InnNPC array at `+0xc0`,
using `L10214(index,1)`. A second call searches the remaining array again.
Thus a duplicated offer array could lose a second matching row; with one
matching offer, the second search finds none. Offer multiplicity is distinct
from the inn pending array's multiplicity.

`R1407` routes multiples of ten to `R0755`; an already-current main
identity takes its equality arm and writes selected `record+0x118` from the
current identity, without calling the mission loader `R1297`. Other
identities use `R0756`, a first-match search of 0x4c-byte side records whose
identity is `+4`, and write selected `record+0x118`.

`R1410` finds the current main record or the first side record with that
identity and calls `R1415`. The latter writes announce `+0x18 = 1` only
when it is zero. `R1408`, called by `R1409` and through `R1416`, walks
side records with a cursor at `record+0x44`, skips entries whose `+0x18` is
zero, and returns a flagged entry's `+4` identity. It consumes campaign flags,
not the inn pending array. Repeated flag writes do not add campaign records.

`R1325` compares the mission with each 0x14-byte journal record's first
dword at `record+0x15c`, count `+0x160`. A match at `L10215` skips the journal
append at `L10216..L10217`. This is an independent identity guard after the
other three calls, not deduplication of the pending queue or commit loop.

**Confidence.** High for the named comparisons, first-match removals, main
identity equality branch, flag latch, journal guard and first consumer. Unknown
for complete campaign consequences: the mission loader, registry/string
helpers, journal constructors/copy helpers and later mission execution are not
read here. No reward-count or native-play conclusion is drawn.

**Unknown.** Mutation by external aliases, preexisting duplicate campaign side
records or journal records, native interleaving and later reward effects. The
selected flow does not prove all repeated registrations are globally harmless.
