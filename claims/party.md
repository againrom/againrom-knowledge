# PARTY — ownership, the roster, and what crosses a map change

Not a file format. Claims about what a player owns and what of it is still the
player's on the next map. The ledger sits between four ledgers that each own
one piece and none of which owns the join: [`tavern.md`](tavern.md) owns the
mercenary pool, [`hero.md`](hero.md) the hero's stats, [`alm.md`](alm.md) the
map's type-5 group roster (`ALM-GRP-041`), plus [`sav.md`](sav.md) the save
container. These claims add the carrier: the engine has no party object, and
one thing that crosses a mission boundary is a serialized object graph; one
thing, not the only one (`PARTY-PERSIST-014`). Format of this file:
[registry.md](registry.md).

## Terms

Two vocabularies meet here and are kept apart.

- The party is two containers on the `Player` plus a pointer
  (`PARTY-ORIGIN-010`, the claim to read first). The group list `+0x24` is what
  a save stores, the flat index `+0x20` is what the placement walk reads, and
  `+0x20` is rebuilt from `+0x24` on load. Which one a consumer implements
  decides which half of the game breaks.
- A `Player` (`0x70` bytes, `CRuntimeClass` at `L07270`) is the simulation's
  owner object, one per map type-5 slot.
- A group (`0x48` bytes) is a container of actors held by a `Player`.
  `ALM-GRP-041`'s "group" is the map record a `Player` is built from, not this
  object.
- The world is the server object at `[L00285]`; `SESS-*` claims call it
  the server singleton.

## Ownership, roster and the client cull

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| PARTY-OWN-001 | Ownership is one pointer, `actor+0x14`, to a `Player`; the `Player` carries two identities, its slot `+0x04` and the map's type-5 id word `+0x08`. | High | ● active | [EXP-0077](../experiments/EXP-0077-party-and-save/) |
| PARTY-ROSTER-002 | There is no party object and no member list: the roster is `Player` → group collection → group → the group's own actor list, and membership is containment. | High / Medium / Unknown | ● active (amended, superseded) | [EXP-0077](../experiments/EXP-0077-party-and-save/) |
| PARTY-FLAG-003 | On the client, side membership is the dword flag word `CUnit+0x18c`: the end-of-mission cull keeps bit 0, the tavern tally reads bit 4 (mercenary), and bit 1 splits that bucket. | High / Medium / Unknown | ● active (amended, partially retracted) | [EXP-0077](../experiments/EXP-0077-party-and-save/) |
| PARTY-CULL-004 | At the end of every mission the client document is cut down to one player and its own surviving player characters; everything else is destroyed. | High | ● active | [EXP-0077](../experiments/EXP-0077-party-and-save/) |

### PARTY-OWN-001

- The load arm of `Player::Serialize` (`R0415`) walks the player's own
  groups and writes `actor+0x14 = player` for every actor it finds
  (the 32-bit store to offset 0x14 at `L07271`).
- Every consumer that asks "is this mine" compares that field against the
  local player object: the client's end-of-mission cull at `L07272`
  (a 32-bit compare against `doc+0x9b4`) and the tavern's live-unit collector at
  `L07273`/`L07274` (`SHOP`/`MERC-DEATH-006` already used the second).
- The `Player` carries two identities, and they are not the same number:
  - `+0x04` is `slot+1`: the index it is re-seated at in the document's player
    array (`L07275`) and the value the wire uses to name a player
    (`L07276`, `L07277`);
  - `+0x08` is the map's own type-5 id word, the value the owner lookup
    `R0425` compares its argument against (`L02234`,
    a 32-bit compare against the field at offset 0x8).
- `ALM-GRP-041` publishes that word from the loader's side; this claim is the
  consumer's side.

**Confidence.** High. The write is one instruction inside the serializer, and
the three reads are named instructions in three unrelated modules. The two
identities are told apart by their two distinct consumers, not by their values.

### PARTY-ROSTER-002

- The roster has three levels:
  `Player → group collection → group → the group's own actor list`.
- `Player+0x24` is a collection of groups, serialized by `R1340`. On
  load it reads a count, then `new(0x48)` + `R0149` per element
  (`L07278`, `L07279`).
- Each group is serialized by `R1143`. It writes its actor list through
  `R1341` (an object list: count via `R1369`, then
  `ar << element` per node), then its own group id at `+0x1c` and two object
  references at `+0x40`/`+0x44` that are id-fixed-up on load (`R1370`,
  `R1371`).
- The players hang off one global list, `[L00380]`, whose serializer is
  `R1372`.
- `SAV-OBJ-016` found inventories and units nesting under their owners with an
  untagged ~100-byte structure between sibling units: that structure is a group
  record.
- Count: the corpus's five `Player`s are the map's five type-5 slots
  (`SAV-OBJ-016`; the count is the fourth head u32, `SAV-HEAD-025`). Cap: a map
  may declare at most 16 (`ALM-META-025`, `ALM-GRP-041`'s sixteen diplomacy
  words).
- No cap on groups per player and none on actors per group was found. Both
  containers are unbounded MFC collections, and no compare against a limit was
  seen on either insert path.

**Confidence.** High for the three-level structure and the group record's
fields: every level is a serializer read end to end, and the shape it writes is
the shape `SAV-OBJ-016` measured in the file. Medium for the 5 and the 16: both
are facts about shipped maps carried from `alm.md`, not bounds read off an
instruction.

**Unknown.** A per-group or per-player member cap. The absence was found by
reading the two insert paths, not by an image-wide sweep.

**Amended.** `SAV-GRPFLD-060` (EXP-0149) supersedes the group-id clause at
`+0x1c` ([`retracted.md`](retracted.md)). `+0x1c` is an authored id only for
groups the map supplies. The group constructor `R0149` writes `+0x40`,
`+0x44`, `+0x3c` and the embedded list at `+0x20`, never `+0x1c`, so for a
group the engine creates at run time the serializer writes out uninitialised
memory. The three-level roster, the two serializers and the load-side
construction are untouched.

### PARTY-FLAG-003

- `R1373` partitions the local player's live units into two buckets in
  one pass:
  - bucket A takes those with the bit-0 test (`L07280`) and corpse stage
    `byte+0x15a == 0`;
  - bucket B takes those with the bit-4 test, mask 0x10 (`L07281`) and the same stage
    test. Bucket B only is split into two counters by a bit-1 test, mask 0x2
    (`L07282`): `+0x2c` if set, `+0x28` if clear.
- The tavern's pool tally `R1374` reads bucket B: it takes the array at
  `L07283` and the count at `L07284`, the `+0x18` / `+0x1c` fields of the
  collector at `L07285`, and subscripts it by the mercenary type
  `byte+0x15b` (`L07286`). So bit 4 is "mercenary" and bit 0 is not. The
  end-of-mission cull keeps bit 0 (`PARTY-CULL-004`).
- `MERC-DEATH-006`'s parenthetical gloss "`+0x18c & 1` (hero)" is right about
  the shipped effect and wrong about the field: the bit means player character.
- `R0592` sets the word as a whole-word literal `0x29`, or `0x2b` when
  `byte[EDI+0x74] & 0x40` (`L02788`, `L07287`). That routine builds a hero
  out of a character record it has just read with eighteen `CFile::Read`
  calls.
- Bits 2, 3 and 7 are `OR`ed in at other sites (`L07288`, `L03217`,
  `L07289`). Bit 5 is also set by the dispatcher's OR-ing in of 0x20 (`L03215`..`L03216`, `TERR-194`).
- Instrument: `EnumRefs disp:18c`, whole image, 232 hits over 81 owners, 0 in
  orphan or undisassembled code. Its blind spot applies here: the word is
  not copied wholesale at the eleven sites in `R0743` and
  `R0591` (they OR single bits or the constant 9, `TERR-194`), and any
  `REP MOVSD` of a containing struct would carry no displacement at all and be
  invisible.

**Confidence.** High for what the three bits do: each of the five tests above
is a named instruction, and the bucket-to-consumer binding is fixed by the two
static offsets the tally reads, not by which bucket looks apt. Medium that bit 0
means "player character" in general: one writer sets it, and this install ships
one campaign, so a second kind of bit-0 unit would not have been seen.

**Unknown.** The remaining bits, and whether the server actor carries a
counterpart field at all: none was looked for.

**Amended.** `PARTY-ENDCULL-026` reads the server's end-of-mission cull. There
is no counterpart flag word: the server's membership predicate at a mission
boundary is a range test on the actor's typeID `word+0x0e`.
Two clauses are withdrawn (`retracted.md`, `TERR-194`): bit 5 has a third setter in the
dispatcher, and the eleven `R0743` and `R0591` sites are masked-bit ORs and
constant stores, not wholesale copies.

### PARTY-CULL-004

- `R0809` has a single caller, `R0701` at `L07290`
  (`MERC-DEATH-006`). It walks the document's object map `doc+0x9b8`, and for
  each `CUnit` (`IsKindOf` against the `CRuntimeClass` at `L00620`,
  `L07291`) applies three tests, all of which must pass:
  - `[+0x18c] & 1` (`L07292`);
  - `(i16)[+0xfc] > −10` (`L07293`, a signed compare with -0xa that skips on less-or-equal);
  - `[+0x14] == [doc+0x9b4]` (`L07272`).
- A survivor is reset in place: `byte+0x15a = 0` (corpse stage),
  `byte+0x84 = 0`, `dword+0xa0 = 0`, then `vt+0x14(0)`
  (`L07294`…`L07295`).
- Everything else is destroyed through `vt+0x04(1)` and unlinked: every
  `CUnit` that fails any test, every non-`CUnit` object, the second map
  `doc+0x9d4` in its entirety, `doc+0x3f3c` and `doc+0x80`.
- The player array `doc+0x9a0` is then emptied from index 1 upward of every
  player except `doc+0x9b4` (`L07296`, a 32-bit compare against `doc+0x9b4` that skips on equal), and the survivor is written back at index `[player+0x04]`
  (`L07275`).
- The kill list is built first and drained second: a `CWordArray` of map keys
  at `[EBP-0x28]`, appended by `SetAtGrow(m_nSize, key)` and drained with
  `RemoveAt(0,1)`.

**Confidence.** High. The three tests, the three resets and the two loops are
named instructions in one routine with one caller. The "everything else" clause
is not an inference: the `IsKindOf`-false arm at `L07297` appends to the same
kill list as the failed-test arm at `L07298`.

## The carry across a map change

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| PARTY-CARRY-005 | What crosses a map change is one compressed `CArchive` graph with one root actor plus sixteen bytes of `Player` scalars, the same bytes in the network packet and the `.chr` file. | High | ● active (amended) | [EXP-0077](../experiments/EXP-0077-party-and-save/) |
| PARTY-LOSS-006 | The carry importer re-homes the actor: a carried character arrives at full pools, off the map, with no position, order or route, and named after the account. | High / Medium / Unknown | ● active | [EXP-0077](../experiments/EXP-0077-party-and-save/) |
| PARTY-MERC-007 | Mercenaries do not cross a map change as objects; what crosses is a count per type in the tavern's mission record, so one party needs two persistences. | High | ● active (amended) | [EXP-0077](../experiments/EXP-0077-party-and-save/) |
| PARTY-SESSION-008 | Apart from zeroing stores, `R0119` is the only writer of the player-list and actor-registry globals; the world global has three writers, and the campaign's mission-to-mission edge does not rebuild the server. | High / Medium / Unknown | ● active (partially retracted) | [EXP-0077](../experiments/EXP-0077-party-and-save/), [EXP-0100](../experiments/EXP-0100-party-origin/), [EXP-0165](../experiments/EXP-0165-join-persistence/), [EXP-0312](../experiments/EXP-0312-tail-and-registers/), [EXP-0341](../experiments/EXP-0341-human-load-first-use/) |
| PARTY-GROUP-009 | A client's group is the one whose `Player` the client is assigned, named by slot and nothing else; hiring a mercenary changes no ownership field. | High / Medium | ● active (amended) | [EXP-0077](../experiments/EXP-0077-party-and-save/) |

### PARTY-CARRY-005

- The exporter `R1375(actor)` opens a `CMemFile`, calls
  `R1110(ar, actor)` (`WriteObject`, one root: a single `ar << actor`),
  pads the memory file to an even length and compresses it with
  `R1376`. That is the save writer's compressor, and the image has
  exactly 2 call sites for it: this one and `R0804`.
- It lays the result into the static packet at `L07299`:
  - `byte+0x09 = 0xbe`;
  - `word+0x07 = [actor+0x14]+0x04`, the owning player's slot;
  - `dword+0x0a` = the payload word count;
  - `memcpy(pkt+0x0e, &{player+0x38, +0x48, +0x4c, +0x54}, 16)`;
  - the compressed graph at `pkt+0x1e`.
- The same 16 + N bytes are written to `chr\<profile>\<u><u>.chr` (format
  string `L07300`) when the profile name at `player+0x64` is non-empty, and
  to `<u><u>.chr` when it is not.
- The client re-sends them: `R1377` builds a `0xbe` packet in the same
  static buffer out of `campaign+0x114` / `campaign+0x118` (`R0821` at
  `L07301`, guarded by `[ESI+0x118] != 0` at `L07302`). When the blob is
  absent it sends the chargen message `0x48` instead (`R0822`).
- The client's packet carries the stored bytes at `pkt+0x0e` (`L07303`) and
  their word count at `dword+0x0a` (`L07304`), but its header differs from
  the exporter's: it stamps its own slot at `word+0x05` (`L07305`,
  `PARTY-GROUP-009`) and writes `word+0x07 = 0` (`L07306`).
- The server's arm (`R0061` at `L07307` / `L07308`) can take the
  payload from the wire or read a `.chr` off disk into the same packet field
  before calling the importer.
- The importer `R1378` has one decompressor call, and the image has
  exactly 2 of those: this one and the save loader.

**Confidence.** High. Both call-site enumerations are `EnumRefs callto:` on the
repaired table with `.rdata` slots included, and each returns 2 hits / 2
owners, 0 orphan. The single root is the single `WriteObject` at `L07309`.
The 16 bytes are one `memcpy` with a named source, and the importer assigns
those four dwords back to the same four `Player` offsets at
`L07310`…`L07311`.

**Amended.** The client packet is corrected ([`retracted.md`](retracted.md)).
The former wording called it "the identical `0xbe` packet". EXP-0077's listing
of `R1377` puts the slot at `word+0x05` and zero at `word+0x07`, where
the exporter puts the slot at `word+0x07` (`L07312`). The 16 + N payload
bytes are unchanged.

### PARTY-LOSS-006

After `ReadObject`, the importer `R1378`:

- overwrites the actor's own name with the receiving `Player`'s CString
  (`L07313`, `[actor+0x80] = player+0x18`);
- frees three sub-objects and installs fresh ones:
  - `actor+0x154` freed and replaced by `new(0xb4)` + `R0206`
    (`L07314`…`L07315`);
  - `actor+0x158` released via `R1379(1)` and replaced by `new(0x94)` +
    `R0286`;
  - `actor+0x10` freed and replaced by `new(0xc)` + `R0287`;
- zeroes five fields `+0x5c`, `+0x64`, `+0x44`, `+0x68`, `+0x40`
  (`L07316`…`L07317`);
- calls `R1380`;
- copies `word+0x96 → word+0x94` and `word+0x9c → word+0x9a`, a pair of
  "current ← maximum" restores;
- sets `byte+0x13c = 0`, clears bit 3 of `byte+0x4c` (AND with 0xf7), sets
  `word+0x18 = 0`, and calls `vt+0x50`.

The actor's name comes from the account rather than from the object. The four
`Player` scalars `+0x38`, `+0x48`, `+0x4c`, `+0x54` ride separately because
they are not in the actor. The complementary loss is `PARTY-MERC-007`.

**Confidence.** High for the field list and the replacements: twenty-odd
consecutive instructions in one routine, each naming its offset. Medium for the
reading of `+0x94`/`+0x9a` as current-from-maximum and of
`+0x40`/`+0x44`/`+0x5c` as map-bound state: the pattern is the same one the
world `Serialize`'s no-map arm applies to every actor it restores
(`SAV-ROSTER-024`), which is corroboration from a second site rather than a
reading of either field.

**Unknown.** What `+0x154`, `+0x158` and `+0x10` are; only that they are
per-map and are never carried.

### PARTY-MERC-007

- The reason is structural rather than a rule.
- Mercenaries are not in the carried graph: the graph has one root, and that
  root is a player character.
- They do not reach it as references: they hang off the `Player`'s groups
  (`PARTY-ROSTER-002`), and the `Player` is not serialized into the blob; only
  four of its dwords are.
- On the client the same units are dropped by `PARTY-CULL-004`, because they
  carry `+0x18c` bit 4 and not bit 0 (`PARTY-FLAG-003`).
- What is carried instead is the count per type, in the mission record the
  tavern keeps: `MERC-DEATH-006`'s merge, and the tavern's own rebuild of the
  whole roster on the next town visit (`MERC-CMD-007`, command `0x37`).
- A consumer implements two different persistences for one party: the hero as
  an object, the squads as fifteen integers.

**Confidence.** High. Each half is a published, instruction-level claim in
another ledger. This claim adds that the two are exhaustive: the carried
payload is one `WriteObject` and sixteen bytes, so there is no third channel.

**Amended.** `MERC-POOL-011` refines the width: it names that array's storage
and the width of its elements.

### PARTY-SESSION-008

- The world is the object `[L00285]`; `SESS-*` claims call it the server
  singleton.
- `R0119` is the only writer of the player list `[L00380]` and the
  actor registry `[L00240]`, apart from `R0967`'s `MOV …,0x0`
  stores.
- Those stores are four absolute zeroings, `L07318` `[L00004]`,
  `L00244` `[L00240]`, `L07319` `[L00380]` and `L07320`
  `[L04624]`, the only absolute operands among the routine's 130
  instructions. Two hit the actor registry and the player list; none hits
  `[L00285]`.
- The world global `[L00285]` has three writes (`PARTY-PERSIST-014`):
  `L07321` in the allocator `R1304`, `L07322` in `R0119`, and
  `L07323`, the teardown `R1303`'s `= 0`. None is in `R0967`.
- `R1304` allocates `0x174` bytes and makes both constructor calls,
  `R0967` and `R0119`, on the new object (`L07324`/`L07325`,
  both passing the new object as the receiver): it builds one world and destroys none. The per-mission
  starter `R0099` is not among its five callers and runs `R0512`
  on the existing object (`L06605`).
- The campaign's mission-end arm reaches no `R1303`, and `R0423`
  then reuses the surviving human `Player` (`PARTY-PERSIST-028`). A carry across
  that edge is therefore not only a serialization.
- `R0512`, the routine `R0099` calls, is read whole by
  `SESS-LOAD-009` and holds the mission-number parse and the save-versus-map
  branch. It finishes by zeroing `world+0x00` and `world+0x04`, the two counters
  `SESS-TICK-004` names and `SAV-HEAD-025` measures in the file, and by setting
  `world+0x2c = 1` (`L07326`). That is the flag the save writer tests to
  decide whether a world half exists at all (`SAV-SHAPE-023`).
- The mission `.ini` it builds a path for, `World\Mission\<n>.ini`, does not
  ship: no `.ini` entry exists in any of the eight `.res` archives (0 of 4 080
  entries), and no `World` directory exists on disk.

**Confidence.** Medium for the single-writer statement on `[L00380]` and
`[L00240]`: it rests on an `EnumRefs refto:` sweep with 0 orphan hits whose
output is not in EXP-0077's committed record, and the sweep's third address,
re-enumerated by `PARTY-PERSIST-014`, has writers outside `R0119`. High
for `world+0x2c`'s single write. High for the four zeroing-store addresses:
one image read twice (two instruction listings of the routine) gives the same four. High for the
world global's three writes, each at a named address (`PARTY-PERSIST-014`; its
own Medium rests on the mission-edge question, not on the write list). High
for the mission-to-mission edge not rebuilding the server: `PARTY-PERSIST-028`
reads the mission-end arm end to end and grades it High. High for the
`.ini`'s absence: a whole-corpus
enumeration of archive entry names plus a directory listing, the instrument
class `AGENTS.md` calls immune.

**Unknown.** What `R0966` does when the file is missing, and therefore
what a shipped campaign loses by its absence. What the two other zeroed
globals, `[L00004]` and `[L04624]`, hold is not read.

**Amended.** Two clauses are withdrawn ([`retracted.md`](retracted.md)).
`PARTY-PERSIST-014` (EXP-0100) replaces the writer set of `[L00285]`; the
former wording made `R0119` the only writer of all three globals apart
from `R0967`'s three `MOV …,0x0`. `PARTY-PERSIST-028` (EXP-0165)
replaces the lifecycle: it reads the mission-end arm the AMBIGUOUS entry left
unread, and that arm reaches no `R1303`. The withdrawn text read:
"nothing survives in memory, because the single-player server is destroyed and
rebuilt at every mission boundary", "calls `R0967` on the old world" and
"A carry is therefore a serialization or it does not happen". The lifecycle was
graded High when published and Medium after EXP-0100. The player-list and
actor-registry writers, `world+0x2c` and the `.ini`'s absence stand.

### PARTY-GROUP-009

- The client stamps `word[pkt+0x05] = word[[doc+0x9b4]+0x04]` on both messages
  it sends about its character: the carry `0xbe` (`L07276`) and the chargen
  `0x48` (`L07327`). That value is `slot+1` (`PARTY-OWN-001`).
- The server's arm resolves the slot to a `Player` and passes it to the
  importer as the second argument. The imported actor's name, money and three
  other scalars are then written from and to that object (`PARTY-LOSS-006`).
- Hiring a mercenary changes no ownership field. `MERC-HIRE-003`'s spawn loop
  `R1142` creates one object per head on the server, and the object's
  owner is the player it was created for. The hire flag
  `dword[record+0x88 + (t−1)·4]` lives in the campaign mission record, not on
  any unit.
- The three ways a unit becomes a player's are one mechanism and three origins:
  the field is always `actor+0x14`, and it is written by the map loader for
  what the map gives, by the spawner for what is bought, or by the importer for
  what is brought.

**Confidence.** High for the slot stamp and the hire's non-effect on ownership:
both are named instructions, the second in a published claim. Medium that the
resolution in the server's arm is by slot: the arm's player argument was read
as a live value at the call site, and the lookup that produced it was not
traced to its own compare. The discriminator is one listing of
`R0061`'s `0xbe` prologue.

**Amended.** The origin count is corrected ([`retracted.md`](retracted.md)).
The former wording read "one mechanism and two origins" beside the three
writers it names: the map loader, the spawner and the importer. The count is
three, one per writer. The slot stamp, the slot resolution and the hire's
non-effect on ownership are untouched.

## Where a player's units come from

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| PARTY-ORIGIN-010 | The party is two `Player` containers: the group list `+0x24` is saved, and the flat index `+0x20` is derived from it on load and is the list the placement walk places from. | High | ● active (partially retracted) | [EXP-0100](../experiments/EXP-0100-party-origin/) |
| PARTY-WRITE-011 | Eleven routines write the player's two containers by call, and the group-list and group-member appends share one owner set, so the two containers are always written together. | High / Medium / Unknown | ● active (amended) | [EXP-0100](../experiments/EXP-0100-party-origin/) |
| PARTY-INSTALL-012 | Giving a player a unit is one straight-line sequence of five writes, duplicated rather than shared at every entry point; a mercenary spawn writes four, without the hero pointer. | High / Medium | ● active (amended) | [EXP-0100](../experiments/EXP-0100-party-origin/) |
| PARTY-GATE-013 | No hero, no mission: `R0131` rejects a client whose `Player+0x34` is null, and that is the only membership test on the way in. | High | ● active | [EXP-0100](../experiments/EXP-0100-party-origin/) |
| PARTY-PERSIST-014 | The single-player carry has a memory arm: `R0423` can reuse the surviving `Player` for the new map's slot 1, and the campaign start selects that arm. | Medium | ● active (amended) | [EXP-0100](../experiments/EXP-0100-party-origin/) |

### PARTY-ORIGIN-010

- The `Player` constructor `R0201` builds both containers (`L07328`,
  `L07329`):
  - `+0x20` is `new(0x20)` + `R0120`, a four-byte vtable (`L06791`)
    wrapping one `CObList` at its own `+0x04`: the player's flat actor index;
  - `+0x24` is `new(0x1c)` + `R1381`, the group list `PARTY-ROSTER-002`
    names.
- Besides the two containers, the party is a pointer and two back-pointers.
  `+0x34` is the hero pointer, zeroed there (`L07330`). `actor+0x14` is the
  owning player (`PARTY-OWN-001`) and `actor+0x70` the owning group.
- `+0x20` is derived, not stored. `Player::Serialize`'s load arm iterates
  `[player+0x24]`'s groups and, for every actor in every group, appends it to
  `[player+0x20]` and stamps `actor+0x14` (`L07331`, `L07332`, `L01794`,
  `L07271`). That is why `SAV-PLAYER-028`'s seventeen fields contain `+0x24`
  and `+0x34` and not `+0x20`.
- The placement walk `R0065` places from `[player+0x20]+4`
  (`L06654`…`L06595`), after a hygiene pass that reads and prunes the
  group list `+0x24` (`L07333`…`L07334`, `PARTY-GATE-013`).
- The walk is not the only reader of `+0x20`: the server's end-of-mission cull
  walks it (`PARTY-ENDCULL-026`), and the join removes an actor from the old
  owner's `[+0x20]+4` (`PARTY-JOIN-025`).
- A consumer that implements the groups alone starts every mission with nobody
  standing on the map, and one that implements the flat index alone loses the
  party at the first save.
- The two globals `[L00380]` and `[L00240]` are the same class
  (`R0119` at `L07335`/`L07336`), so "my units" and "all units" are
  the same shape at two scopes.

**Confidence.** High. Every field is a named instruction in a constructor or a
serializer read end to end. The derivation is not an inference but the load
arm's own direction of copy, groups → index. The rival "one collection seen
twice" is excluded by the two `new`s of different sizes with different
constructors.

**Unknown.** The complete reader set of `+0x20`: EXP-0100's commands hold no
sweep of its readers.

**Amended.** The exclusivity clause is withdrawn
([`retracted.md`](retracted.md)). The former wording read "read only by the
placement walk" and "`R0065` reads only `[player+0x20]+4`". The walk's
own hygiene pass reads `+0x24` (`PARTY-GATE-013`, EXP-0100), and
`PARTY-ENDCULL-026` and `PARTY-JOIN-025` (EXP-0165) read `+0x20`. The two
containers, the derivation of `+0x20` from `+0x24` on load and the walk's
placement from `+0x20` stand.

### PARTY-WRITE-011

- Instrument: `EnumRefs` over direct calls to `R0032` (the flat index's append), 32
  hits, 20 owners, 0 in orphan or undisassembled code. It prints the
  instruction that loaded `ECX`, so a hit with no function name cannot be lost.
- Twelve hits, in eleven routines, load `ECX` from a `+0x20` displacement.
  Seven routines are read at instruction level:
  - `R0823` ×2 (chargen, `SESS-HERO-014`; the repair arm `L07337`
    and the fresh-hero arm `L07338`, `PARTY-INSTALL-012`);
  - `R0061` (the `0xbe` carry arm, `PARTY-CARRY-005`);
  - `R1142` (the mercenary spawn, `MERC-HIRE-003`, `PARTY-INSTALL-012`);
  - `R0415` (`Player::Serialize`, the load arm);
  - `R0065` (the `.ini` "Humans" arm);
  - `R0064` (the mid-mission join, `PARTY-JOIN-025`);
  - `R0066` (the `AddHero` install behind command `0x49`,
    `PARTY-ADDHERO-017`).
- Four are classified by their append site alone: `R0151` (the map's
  own type-6 spawner, `MISSION-ARM-006`; its append at `L06603` is not
  read), `R0062` (`L07339`), `R0003` (`L07340`) and
  `R0432` (`L07341`).
- The group list's append (`EnumRefs` over direct calls to `R0198`: 13 hits, 11
  owners, 0 orphan) and the group's own actor append (`callto:R0152`: 15 hits,
  11 owners, 0 orphan) have the same owner set. That is the evidence that the
  two containers are always written together.
- `callto:R0292`, the raw `CObList::AddTail` beneath the three append helpers,
  returns 4 hits, 4 owners, 0 orphan: the three helpers plus `R0914`. So
  nothing else reaches these lists by call.
- Blind spot: a wholesale `REP MOVSD` of a `Player` carries neither call nor
  displacement and is invisible to every sweep here.

**Confidence.** High for the enumeration being complete by call: three sweeps
on the repaired table, 0 orphan hits in each, and the raw-`AddTail` sweep closes
the route around the helpers. Medium that these eleven are the complete set of
origins: seven are read at instruction level and four are classified by their
append site alone.

**Unknown.** Whether any path copies a whole `Player`.

**Amended.** Two corrections ([`retracted.md`](retracted.md)). The former text
listed six routines with their roles and named five "which EXP-0100 did not
read", against its own "five were read at instruction level and six were
classified by their append site alone". The five read are `R0823` at
both sites, `R0061`, `R1142`, `R0415` and
`R0065`; the sixth classified routine is `R0151`, whose
instructions no cited record lists. EXP-0100's `DisasmFn` list holds neither it
nor `R1142`; the read status of `R1142` rests on four write
addresses quoted in EXP-0100's record (`L07342`, `L07343`, `L07344`,
`L07345`) and on `MERC-HIRE-003`. `PARTY-JOIN-025` (EXP-0165) reads
`R0064` and `PARTY-ADDHERO-017` (EXP-0153) reads `R0066`: each
appends the actor to `+0x20`, a new group to `+0x24` and the actor to that
group, as their append sites classified them. The three sweeps, their counts
and the shared owner set are untouched.

### PARTY-INSTALL-012

- The `0xbe` carry arm (`L07346`…`L07347`) writes, in order:
  1. `actor+0x04 = R0936() & 0xffff`, a fresh runtime id;
  2. `actor+0x14 = player`;
  3. `player+0x34 = actor`;
  4. `[player+0x20].AddTail(actor)`;
  5. `new(0x48)` + `R0149` → `[player+0x24].AddTail(group)` →
     `R0152(group, actor)`.
- The chargen arm `R0823` writes the identical five
  (`L07348`…`L07349`, `player+0x34` at `L07350`).
- The mercenary spawn `R1142` writes four of them (`L07342`,
  `L07343`, `L07344`, `L07345`): everything but `+0x34`, which is what
  makes a mercenary not a hero.
- A carried character therefore arrives in a group of its own, never in the
  group it left.
- `R0823`'s first arm is a repair rather than a creation. On
  `player+0x34 != 0` it tests `byte[actor+0x13c]`:
  - if that is 0, it returns the hero untouched (`L07351`, `L04476`,
    `L04477`), the state `PARTY-LOSS-006`'s importer leaves a carried actor
    in;
  - a nonzero value restores `word+0x94 ← +0x96` and `word+0x9a ← +0x9c`,
    re-indexes into `+0x20` and rebuilds a group if `actor+0x70` is null
    (`L04479`…`L07352`).
- Its one non-dispatcher caller is `R0132`'s revive arm, which calls it
  with all-zero stats and then re-runs the placement walk (`L03850`,
  `L06590`).

**Confidence.** High for the two sequences and the branch: three routines read
at instruction level, each write naming its own offset; the duplication is a
fact about two listings, not a reading. Medium for the reading of `+0x13c` as
the flag that separates "repair a dead hero" from "leave a carried one alone":
the two states are inferred from `PARTY-LOSS-006`'s zeroing and from the revive
caller, and the field itself was not read.

**Amended.** Narrowed by EXP-0524: "a carried character therefore arrives in a
group of its own" holds for an actor these arms install. At a mission start a
carried member keeps the Group of the reused Player, which can hold 3 or 4
carried members and AI bytes that appear to predate that mission
(`SAV-1219`, `AI-449`). Medium for the mechanism: 9 distinct restart-slot
files and 12 carried Groups were read, and the carry arm `L07346` with its
Group build `L14238` was not.

### PARTY-GATE-013

- `R0131`'s single caller is the session dispatcher's `0x04` arm
  (`L06672`, inside the arm at `L06671`). It tests `player+0x34` first
  (`L07353`). On null it logs
  `"Client %s tries to enter mission without Hero. Rejected."`
  (`L07354`/`L07355`), sends `"You can't enter mission without Hero"`
  (`L07356`) and returns without placing anything.
- An empty flat index is not tested and is not an error: it places nobody and
  the mission runs.
- Past the gate the routine clears `player+0x3c` (`SAV-FLAG-027`'s outcome
  latch, `L07357`) and calls the walk once, under a second latch:
  `byte player+0x3d`, set to 1 immediately after (`L07358`, `L07359`). That
  byte is the tenth field of `Player::Serialize`'s record (`SAV-PLAYER-028`),
  so a player placed and then saved is never re-placed on load.
- The walk begins with a hygiene pass on the group list that nothing else does:
  `GetHead`, and if there is none, log `"Oops - player has no groups %-["`
  (`L07360`) and create one. It then deletes empty groups from the head until
  the head is non-empty or is the last one (`L07333`…`L07334`).

**Confidence.** High. One routine and one loop were read at instruction level,
and three of the four arms are named by the image's own strings. The `+0x3d`
latch is a store two instructions after the call and a field of a published
serializer record.

### PARTY-PERSIST-014

- `R0423` builds one `Player` per type-5 record, except that for roster
  slot 1 it reuses the player already in the global list when two conditions
  hold: `[L00285]+0x0c == 0`, and the list holds exactly one player
  (`L07361`…`L07252`, three compares in a row, `GetSize` then `GetHead`).
- That mode flag is written once, in the server's own constructor, as
  `arg0 < 2` (`L06608` compares the first argument with 2, `L06609` sets the flag when it is less, and `L06610` stores it). The campaign-start routine passes 2 (`R1382`,
  the argument 2 pushed at `L07362` → the call to `R1304` at `L07363`), so on the campaign path
  the flag is 0 and the reuse arm is the live one.
- The server object is not rebuilt per mission. `EnumRefs refto:L00285` gives
  172 hits, 93 owners, 0 orphan, with exactly 3 writes: `L07321` the
  allocator, `L07322` its constructor, `L07323` the teardown's `= 0`.
- `EnumRefs` over direct calls to `R1304` gives 5 hits, 5 owners, 0 orphan, and the
  per-mission starter `R0099` is not one of them: it calls
  `R0512` with `ECX = [L00285]`, the existing object (`L06605`).
- This contests `PARTY-SESSION-008`'s "a carry is therefore a serialization or
  it does not happen", whose premise was that `R1304` destroys the old
  world. Both of its constructor calls, `R0967` at `L07364` and
  `R0119` at `L07365`, are made on the newly allocated object
  (`L07324`, passing the new object as the receiver), so that routine builds one world and destroys none.

**Confidence.** Medium. The reuse arm, the mode flag and the write enumeration
are instruction-level, and the campaign's argument is a literal `PUSH 2`, but
whether the shipped campaign's mission→mission edge always avoids the teardown
was not established. Discriminator: `R0701` is the campaign state
machine and holds both `R0099` (`L07366`) and four calls to the
teardown `R1303`. The `L07366` arm jumps straight to the common exit
`L06230` without reaching any of them, but which arms the campaign drives, in
what order, was not read.

**Amended.** `PARTY-PERSIST-028` (EXP-0165) reads the mission-end arm: it
reaches no `R1303`, so the memory carry is the live one on the campaign's
mission-to-mission edge. Whether the `0xbe` carry also fires on that edge, and
what builds the human `Player` before the campaign's first map load, stay open.

## The purse, and AddHero companions

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| PARTY-MONEY-015 | `Player::Player` stores a zero purse at `L07367`; the conclusion that nothing credits it before the first mission is withdrawn, because the participant factory stores 100. | High | ● active (partially retracted) | [EXP-0151](../experiments/EXP-0151-mission-documents/), [EXP-0156](../experiments/EXP-0156-start-money/) |
| PARTY-MONEY-016 | Each `Player` owns one 32-bit purse at `+0x38`, and the two persistence paths, the save record and the mission carry, carry that same field. | High / Medium | ✔ promoted | [EXP-0152](../experiments/EXP-0152-money-cycle/) |
| PARTY-ADDHERO-017 | `[Mission<n>] AddHero[]` creates a companion player character, neither the primary character nor a mercenary, when the town view activates. | High / Medium | ● active | [EXP-0153](../experiments/EXP-0153-addhero-consumer/) |
| PARTY-MONEY-018 | The constructor's 32-bit zero purse, the purse-free chargen opcode `0x48` and the scoped mission and data negatives stand; the conclusion that a new campaign reaches play with `Player+0x38 == 0` is withdrawn (`PARTY-MONEY-024`). | High | ● active (partially retracted) | [EXP-0154](../experiments/EXP-0154-campaign-start/), [EXP-0156](../experiments/EXP-0156-start-money/) |
| PARTY-MONEY-024 | A fresh campaign starts with 100 in the human participant's `Player+0x38`; zero is only the constructor intermediate. | High / Medium | ✔ promoted | [EXP-0156](../experiments/EXP-0156-start-money/) |

### PARTY-MONEY-015

- The constructor fact stands: `Player::Player` stores 0 at `L07367`.
- The registry, type-5, mission-10 type-8 and instant-23 censuses remain
  correct in their stated scopes.
- Participant factory `R0913` directly tests `Player+0x38` and stores
  `0x64` when it is zero (`L07368`…`L07369`). Its one direct caller is the
  join routine `R0910`, which then calls `R0449` with delta 0 at
  `L04967`.
- The original call-site instrument could not see the direct store it was
  being used to exclude.

**Confidence.** High retained for the constructor and the four scoped data
negatives.

**Amended.** EXP-0156 refutes the complete no-credit conclusion
(`PARTY-MONEY-024`, [`retracted.md`](retracted.md)); it had been graded
Medium. The withdrawn headline read: "The purse starts at zero and nothing at
or before the campaign's first mission credits it." Two lawful
original-runtime saves decode to 100, including an automatic save at sub-tick
1.

### PARTY-MONEY-016

- `Player::Serialize` loads the dword at `L07370`, passes it through the XOR
  involution `0x5c073f4d` and writes one u32 (`SAV-OBF-029`). This record is in
  the campaign half and therefore in both save shapes.
- Independently, the mission carry copies sixteen bytes from
  `{Player+0x38,+0x48,+0x4c,+0x54}` into the packet and character file, and its
  importer writes those four dwords back to the same offsets
  (`PARTY-CARRY-005`).
- The purse is therefore per participant, not a field on the hero or one
  campaign-global balance.
- Arithmetic readers decide signed interpretation: the shop affordability gate
  uses a signed `JGE`, while instant 23's bare `ADD` has no clamp and wraps
  modulo 2^32 (`TRIG-MONEY-028`).
- G2: widening the purse changes the `Player` memory layout, save member width,
  sixteen-byte carry block and packet ABI. It does not require changing a
  shipped map or registry file.

**Confidence.** High for object, offset, width, save and carry: two
independently read paths name the same field and width. Medium for signed
interpretation as a general law: individual operations are read, but no type
declaration survives in the image.

### PARTY-ADDHERO-017

- `R1383` calls `R1384` on the current mission record. That
  routine sends one command `0x49` per `u16` and clears the array.
- Command `0x49` calls `R0066`, which assigns the existing player as
  owner, appends the actor to `Player+0x20`, creates and appends a group to
  `Player+0x24`, and inserts the actor into that group.
- It does not write `Player+0x34` and does not use the mercenary constructor or
  type field. The actor is therefore a companion player character, not the
  primary character and not a mercenary.
- The serialized group list preserves it in saves. Complete mission-to-mission
  memory persistence retains `PARTY-PERSIST-014`'s Medium limit.

**Confidence.** High for activation, insertion, classification and save
persistence: complete instruction paths, with the distinct primary and
mercenary alternatives excluded. Medium for uninterrupted mission-to-mission
memory persistence.

### PARTY-MONEY-018

- The retained facts: the constructor still writes a 32-bit zero, opcode `0x48`
  still carries no purse, and the scoped mission and data negatives stand:
  the four selected `Humans` rows and the direct weapon branch do not call
  the purse delta helper, and mission 10 has no registry money input, type-8
  gold, instant-23 grant, mission-entry delta call or first-tick delta call on
  either root.
- What failed is the jump from those facts to playable state: the participant
  factory writes 100 directly before hero creation, outside the enumerated
  delta-helper callers.
- The first-sub-tick automatic save and a later explicit save both decode to
  100.

**Confidence.** High for the retained subclaims, each in its own scope: the
grade [`retracted.md`](retracted.md) records for them. The start value and the
complete no-grant conclusion, graded Medium, are retracted.

**Amended.** Partially retracted by EXP-0156: the claim's own prediction failed
(`PARTY-MONEY-024`, [`retracted.md`](retracted.md)). The withdrawn headline
read: "A new campaign reaches the first playable mission with
`Player+0x38 == 0`, independent of character-generation choices." The
retraction entry keeps the constructor, field-width, chargen-packet and scoped
mission and data observations.

### PARTY-MONEY-024

- `Player::Player` stores 0 at `L07367`.
- The participant factory `R0913`, whose only direct caller is join
  routine `R0910`, then, at `L07368`…`L07369`, tests `Player+0x38` against zero and stores 0x64 into it only when it is zero. It therefore preserves an
  existing nonzero purse and gives 100 to a fresh one.
- Join next calls `R0449(player,0,1)` at `L04967`; its store at
  `L07371` adds zero, leaving 100, and emits opcode `0x67`. The client copies
  packet `+0x0a` into its local participant balance at `L04963`.
- Chargen opcode `0x48` contains stats and appearance/class/sex but no purse,
  and resolves this already existing `Player`. Thus no purse write occurs after
  chargen confirmation and before play; the playable state inherits the join
  value.
- Two lawful original-runtime mission-10 saves independently decode the human
  Fergard record by XOR with the constant `0x5c073f4d`: 100 at sub-tick 1/full tick 0
  and 100 at sub-tick 129/full tick 8.
- Last-writer vocabulary: `L07369` is the last value-changing start write;
  `L07371` is the later physical store and UI notification with delta zero.

**Confidence.** High for value, owner, ordering, save representation and
notification: every hop is a named instruction, `callto:R0913` is 1 hit / 1
owner / 0 orphan, and two different-tick original-runtime saves discriminate the
constructor rival. Medium that the experiment's typed writer list is image-wide
complete: `disp:38` is 889 accesses / 367 owners / 2 orphan, and whole-object
copies remain a blind spot. The start chain does not depend on that universal.

## Joins and the mission-end boundary

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| PARTY-JOIN-025 | A mid-mission join is one routine, `R0064`, reached from trigger instants 19 and 22: `PARTY-INSTALL-012`'s install sequence without the hero pointer or a fresh runtime id. | High | ● active (amended) | [EXP-0165](../experiments/EXP-0165-join-persistence/) |
| PARTY-ENDCULL-026 | The server's end-of-mission cull `R0826` decides which players and actors exist on the next map; it keeps an actor iff `0x21 <= word[actor+0x0e] < 0x40`. | High | ● active (partially retracted) | [EXP-0165](../experiments/EXP-0165-join-persistence/) |
| PARTY-BAND-027 | Withdrawn: the server survival band `[0x21,0x40)` was said to be the Humans creation arm's typeID range; the band is real, but that arm's output is conditional (`PARTY-M20-031`). | High | ✖ retracted | [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |
| PARTY-PERSIST-028 | The campaign's mission-to-mission edge preserves the surviving human `Player`, its name and every actor that first passes the client and server filters, not an unfiltered roster. | High | ● active (amended, superseded) | [EXP-0165](../experiments/EXP-0165-join-persistence/), [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |
| PARTY-JOINCORPUS-029 | The shipped join corpus holds 28 runtime instant-19/22 nodes on 16 maps, 22 of which hand actors to player 1; the handed Humans are not all on the kept side. | High / Medium | ● active (amended, partially retracted) | [EXP-0165](../experiments/EXP-0165-join-persistence/), [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |

### PARTY-JOIN-025

- Trigger instants 19 (`Change unit's owner`) and 22 (`Change group's owner`)
  both dispatch to `R0064`: arm `L07372` once, and arm `L07373` once
  per group member (`TRIG-ACT-004`'s table).
- In order, the routine:
  1. removes the actor from its current group if it has one (`L07374`,
     `R1140`);
  2. removes it from the old owner's flat index (`L07375`, `R1133` on
     `[old+0x20]+4`);
  3. sets `actor+0x14 = newPlayer` (`L07376`);
  4. appends to `[new+0x20]` (`L07377`, `R0032`);
  5. allocates `new(0x48)` + `R0149` (`L07378`, `L07379`);
  6. appends that group to `[new+0x24]` (`L07380`, `R0198`);
  7. inserts the actor into it (`L07381`, `R0152`).
- It does not write `player+0x34` and does not allocate a fresh runtime id,
  which is what separates it from the `0xbe` carry arm and from chargen.
- It then clears both the old and the new owner's `+0x2c` mask from
  `actor+0x18` (`L07382`, `L07383`) and broadcasts through `R0059`.
- A joined actor arrives in a group of its own, as a carried character does.

**Confidence.** High. One routine read end to end; every field is named by its
own instruction, and the two absences are two of the five writes
`PARTY-INSTALL-012` lists for the other entry points.

**Amended.** Narrowed by EXP-0524: the comparison "as a carried character
does" holds for the install arms of `PARTY-INSTALL-012`, not for a carried
member at a mission start, which appears to keep its existing Group
(`SAV-1219`, `AI-449`; Medium: 9 distinct files and 12 carried Groups, the
carry arm `L07346` unread). The joined actor's own Group stands.

### PARTY-ENDCULL-026

- Its call site is the campaign state machine's mission-end arm, `L07384`,
  nineteen instructions after `PARTY-CULL-004`'s client cull.
- Three passes:
  1. per player with `+0x28 == 0` and a hero, drop the hero from the map's cell
     structure;
  2. empty the tick list `[L00240]+4` (`MOVE-TICK-009`) and expire every
     effect hanging off each of the player's actors;
  3. the roster cull.
- The roster cull destroys a `Player`, removes it from `[L00380]` and calls
  `vt+0x04(1)` through `R1385` when `+0x28 != 0`, or when the `CString`
  at `player+0x18` compares byte-equal to the literal `"Self"` at `L07385`
  (`L07386`...`L07387`; the compare is `R1386` to `R1387` to
  the CRT byte loop at `R0751`).
- A surviving player has its placement latch cleared (`byte+0x3d = 0`,
  `L07388`, the field `PARTY-GATE-013` names), and then its flat index `+0x20`
  is walked.
- An actor is kept iff `0x21 <= word[actor+0x0e] < 0x40` (`L07389`,
  `L07390`). Otherwise it is destroyed through `R1388`, which unlinks it
  from its group, drops that group from `player+0x24` if it is left empty and
  its owner is human, removes it from `+0x20` and calls `vt+0x04(1)`.
- A kept actor is reset in place: `word+0x94` from `+0x96`, `word+0x9a` from
  `+0x9c`, `byte+0x13c = 0`, `R1380`, then `+0x5c`, `+0x64`, `+0x44`,
  `+0x68`, `+0x40` zeroed (`L07391`...`L07392`).
- Neither `R0826` nor the reset helper `R1380` writes an
  inventory field, not the sack `+0x7c` (`TRIG-ADDITEM-027`) and not the
  twelve worn slots `+0x198` (`SAV-CARRY-050`), so a survivor keeps what it
  carried.
- This is the server counterpart `PARTY-FLAG-003` recorded as never looked for,
  and it is not a flag word: the server asks the actor's typeID.

**Confidence.** High. One routine was read end to end with its five helpers;
each test and each store is a named instruction, both removals are calls to the
object's own `vt+0x04(1)`, and the literal was read out of the PE through its
section table.

**Amended.** The inventory clause is partially retracted by EXP-0511
(`PARTY-M100-034`, [`retracted.md`](retracted.md)). The reset helper's first
callee `R0014` appends a castSpell cast-slot item to the pack `+0x7c`
under the three conditions `PARTY-M100-034` names (item byte `+0x44` not 2,
state not 0xd or 0xe with `+0x136` = 0, matching spell id), and `R1380`
deletes a class-14 item still in the cast slot. The pack and worn slots are
otherwise unwritten, and every other clause stands.

### PARTY-BAND-027

- `R0656` first streams `Humans` slot 16 into `actor+0x0e`, and
  overwrites it with `gender+0x21` or `+0x23` only when its constructor-mode
  argument is non-zero.
- Mission 20 supplies zero on its definition-id arm and supplies the exact
  `npc.reg` `Hero` flag on its npc arm. Its four transferred Humans therefore
  retain `typeID` `0x17,0x0a,0x0a,0x0a` and all fail the band. See
  `PARTY-M20-031`.

**Confidence.** High for the refutation: the conditional constructor branch,
both spawner arguments, both culls and two original-save boundaries
discriminate it from the creation-arm model.

**Amended.** Retracted as a whole by EXP-0192 (`PARTY-M20-030` to
`PARTY-M20-032`, [`retracted.md`](retracted.md)). The withdrawn text read: "The
survival band `[0x21,0x40)` is the Human creation arm's own typeID range" and
"A script handover of a Human is permanent".

### PARTY-PERSIST-028

- `R0701` reaches the client cull and then `R0826` before town
  construction, and reaches no teardown.
- `R0423` consequently reuses the lone surviving player and skips the
  map name assignment.
- Repetition of that reuse is unbounded, but membership at each boundary is
  conditional: mission 20's four transferred Humans are gone before reuse
  (`PARTY-M20-031`).

**Confidence.** High for the ordered boundary and reuse arms, read end to end.
Original saves independently witness the narrowed membership.

**Amended.** Narrowed by EXP-0192 (`PARTY-M20-031`);
[`retracted.md`](retracted.md) records the unqualified roster-persistence
reading as superseded. The withdrawn sentence read: "Persistence is therefore
unbounded across boundaries". Reuse is unbounded only for the Player and the
actors that survive both filters; it is not a bypass around them. The surviving
Player reuse arm and the name rule stand.

### PARTY-JOINCORPUS-029

- Mission 20 alone supplies four counterexamples to the previous partition of
  the handed Humans as the kept side: one npc-arm Human without the `Hero` flag
  and three definition-id-arm Humans, all with out-of-band table typeIDs.
- Mission 40's measured recruit still survives, for the narrower reason that
  `npc25` has the exact `Hero` flag, so its npc arm passes a non-zero
  constructor mode and performs the player-character overwrite.
- The published mission-40 saves and `AddHero` census otherwise stand.

**Confidence.** High for the corrected mission-20 and mission-40
classifications at instruction level. Medium for the corpus counts and the save
observation.

**Amended.** Narrowed by EXP-0192 (`PARTY-M20-030`, `PARTY-M20-031`);
[`retracted.md`](retracted.md) records the Humans-as-kept-side partition as
refuted: a creation arm does not decide the cull side. The node counts, the
mission-40 saves and the `AddHero` census stand.

## Mission 20's transferred actors

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| PARTY-M20-030 | Mission 20 transfers exactly four live actors to player 1 when `T08` fires: group 16 units 136–139. | High / Medium | ● active | [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |
| PARTY-M20-031 | A transferred Human survives a mission boundary by constructor mode and resulting typeID, not by being a Human. | High | ● active | [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |
| PARTY-M20-032 | Mission 20's transferred actors have no town, tavern, later-mission or post-boundary save identity. | High / Medium | ● active | [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |

### PARTY-M20-030

- Unit 136 is `npc52`, runtime name Sarindar, `Humans[201] M10_Merchant`,
  table `typeID=0x17`, with a Bronze Amulet and Rare None Robe.
- Units 137–139 are three unnamed instances of definition 114,
  `Humans[58] NPC14_1`, table `typeID=0x0a`, each with Iron Mace, Hard Leather
  Small Shield, Uncommon Hard Leather Helm, Bronze Cuirass, Uncommon Hard
  Leather Mail, Bronze Bracers, Leather Gauntlets and Leather Boots.
- EN and RU resolve the same rows, typeIDs and equipment.
- An original mission-20 save independently holds the same four under the human
  player with runtime ids 111–114, class keys 201/58, type words `0x17/0x0a`,
  2/6 worn slots and empty packs.

**Confidence.** High for identities, count and fields: the
trigger/group/placement/Data.bin join agrees on both roots, and an original
save observes a different representation. Medium for the translated display
name, observed in the EN save only.

### PARTY-M20-031

- `R0656` streams Humans slot 16 to `actor+0x0e`; only a non-zero mode
  argument enters `L03163`'s overwrite to a player-character value.
- `R0151` passes zero from the definition-id and explicit typeID arms,
  and passes the exact `"Hero"` flag lookup from the npc arm.
- The server cull keeps only `[0x21,0x40)`. Independently, the client
  classifier gives mission 20's `0x17` and `0x0a` actors no keep bit, and the
  client cull requires that bit.
- Mission 20's sole reachable victory route calls the client cull, then the
  server cull, before town; its authored VIP condition is unreferenced. All four
  actors are therefore detached and destroyed at completion if still present.
- Brian in mission 40 survives because `npc25` is flagged `Hero`, not because
  it took the Humans arm.

**Confidence.** High. Two independent culls agree; the constructor mode and each
spawner arm are named instructions; mission-20 and mission-40 data discriminate
the competing creation-arm rule.

### PARTY-M20-032

- Closely paired mission save `game0009` contains all four under Danath. Town
  save `game0010` has matching Player and Danath allocator identities but none
  of the four; matching values do not prove that both files came from one
  process. Its second actor is fresh town `AddHero=22` companion Reniesta.
- Mid-mission save/load serializes the four through the Player group graph. A
  town save cannot recreate records already culled, and load remaps pointer
  identities between processes.
- Tavern stock is independent, even though the client classifier puts these
  actors in the bit-4 bucket the live tally reads: the simulation and client
  constructors clear the mercenary type, the broadcaster omits that field group
  for zero, and the tally matches only 1..15.
- Mission 20 lists locked type 1 and therefore has an empty shelf. Mission 30
  offers type 14, whose fresh level-1 roster object reuses the same `NPC14_1`
  template and equipment as the three removed guards. Matching dress is
  template reuse, not actor conversion.
- Later mission-31 and mission-40 saves contain only Danath and Reniesta under
  the human Player. Because they are from another session, they corroborate
  roster absence but establish no identity comparison.

**Confidence.** High for the static cull, campaign creation and shelf/template
paths. Medium for the paired-save boundary; for the zero-type stock dependency,
because no image-wide writer sweep is committed; for later-session absence; and
for the predicted fresh type-14 runtime identity, which was not viewed in the
tavern UI.

## Mission 100's servant and the kept-actor reset

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| PARTY-M100-033 | Mission 100's servant, unit 135 (owner 5, group 21), is a mode-0 placement and is destroyed by the client cull at the mission end, whether or not group 21 was handed to player 1. | High / Medium | ● active | [EXP-0511](../experiments/EXP-0511-m100-servant-amulet/) |
| PARTY-M100-034 | The kept-actor reset returns a castSpell cast-slot item to the pack except when its byte +0x44 is 2, the actor's +0x54 is 0xd or 0xe with +0x136 = 0, or the spell id differs; a class-14 item left there is deleted. | High / Medium | ● active | [EXP-0511](../experiments/EXP-0511-m100-servant-amulet/) |

### PARTY-M100-033

- `100.alm` type 6, EN and RU identical: unit 135 is record 68, flags 0, so the definition arm places it, owner slot 5 "Friends", group 21, the only member.
- The mission-end arm of `R0701` calls the client cull `R0809` at `L07290`, then the server cull `R0826` at `L07384`, then `R1660` and `R1658(-1)` (window `L07928`..`L13769`).
- The client cull keeps a `CUnit` only when `[+0x18c] & 1` is set and the owner is the local player (`PARTY-CULL-004`). Bit 0 is set only by the hero arm with a non-zero constructor mode (`UNIT-140`), which a flags-0 placement does not take (`PARTY-M20-031`).
- Action 16 (group 21 to player 1) is in no trigger slot (`TRIG-M100-095`). A join would change the owner test but not bit 0, so the servant fails the client cull either way.

**Confidence.** **High** that the client cull destroys the servant (two named tests; the placement mode is read from the map). **Medium** for the server-cull result, which needs the typeID of definition 1009; that value was not read.

### PARTY-M100-034

The server cull resets a kept actor through `R1380` (`PARTY-ENDCULL-026`). `R1380` (67 instructions, read whole) runs in this order:

1. It calls `R0014` (`L13775`). That helper (84 instructions, read whole) exits when `actor+0x68` is null, when the item's byte `+0x44` is 2, when the item's first effect (as returned by `R2173`, not read) is missing or its kind byte is not `0x29` (`L13776`), or when `actor+0x54` is 0xd or 0xe while `+0x136` is 0. Otherwise it compares the spell id byte at `[+0x64]+8` with the effect word `+0x40` (`L13777`). When they are equal it deletes the spell, clears `+0x64`, appends the item to the pack `actor+0x7c` (`R0929`, `L13778`) and clears `+0x68`. When they differ it logs string `L13779`.
2. When `+0x136` is 0, `+0x68` is non-null and the item's class (`R1056`: `(word +0x40 >> 8) & 0xf`) is 14, it deletes the item through `vt+4(1)` (`L13780`), clears `+0x68` (`L13781`) and deletes the spell at `+0x64` (`L13782`).
3. It clears `+0x64` and `+0x58`, sets `+0x136 = 1`, and stores `+0x54 = 0` and `+0x50 = 0` (`L13783`..`L13764`).

The two bodies write no other actor field (`stores.tsv`). An item in neither path stays in the pack or worn slot it was in. The server cull then zeroes `+0x68` (`PARTY-ENDCULL-026`). A cast-slot item of another class and another effect kind is therefore unlinked without a return to the pack. A class-14 castSpell item that fails any of the three tests of step 1 (byte `+0x44` = 2, state 0xd or 0xe with `+0x136` = 0, or a spell id mismatch) is still in `+0x68` at step 2 and is deleted there when `+0x136` is 0.

**Confidence.** **High** for the order, the tests and the stores (two bodies read whole). **Medium** that the tested effect is the item's first: `R2173` was not read. **Medium** that no callee writes the pack or the worn slots: `R2173` and the `vt+4` destructors were not read. Effect kind `0x29` is castSpell (`MAGIC-ITEM-007`).

**Unknown.** Which items reach `actor+0x68` at a mission end in play.

## Inventory across a mission start

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| PARTY-037 | No mission-start routine read moves an item between party members on the campaign's mission-to-mission edge; 11 unread closure owners, the SAV-load path and three census blind spots bound this. | High / Medium / Unknown | ● active | [EXP-0519](../experiments/EXP-0519-hero-build-fallbacks/) |
| PARTY-038 | A member the mission-end culls remove takes its stack out of the party, and a second `Quest Documents` stack is made only when a primary hero is constructed fresh. | High / Medium / Unknown | ● active | [EXP-0519](../experiments/EXP-0519-hero-build-fallbacks/) |

### PARTY-037

The inventory is the actor's own container at `actor+0x7c` (`ITEM-CONT-004`). On the campaign's
mission-to-mission edge the surviving actors are the same objects (`PARTY-PERSIST-028`); the only
boundary routine that writes an inventory field is the end-of-mission reset, which touches the
cast slot `+0x68` alone (`PARTY-M100-034`). No mission-start routine read rebuilds inventories. The scope is the campaign's
mission-to-mission edge and the routines named below; it does not cover the SAV-load start path
or the unread owners.

- The type-5 player builder `R0423` (read whole): its reuse arm (`L07361`..`L13990`,
  `L13991`..`L02132`) writes `player+0x44` and `player+0x28` and touches no actor.
- `R0823` (read whole) with `player+0x34 != 0` returns that hero (`L07351` → `L13992`
  → `L13993`). Its repair arm (`L13994`..`L13995`, taken when `byte[hero+0x13c] > 0`) calls
  `R0825`, copies the two pools and re-installs the hero in `+0x20` and a group; it makes
  no container call. The `Quest Documents` producer `R0993` (`L13987`) is on the
  construction arm only.
- `R0993` (read whole) tests `server+0x0c == 0` (`L13996`) and `[L03937] != 2`
  (`L13997`), builds an item from the literal at `L13998` through `R2076`, and appends
  it to its argument's `+0x7c` (`L13999`). It reads no other actor's container and does not look
  for an existing stack.
- Container calls: 45 direct `E8` calls in `.text` to the add `R0929`, insert `R0930`, drain
  `R0450` and take `R1522` (`callto.tsv`).
- Direct-call closure from `R0423` and the placement walk `R0065`, depth 6: 419
  bodies, 125 indirect calls not followed. Two owners call a container routine: `R0907`, the
  `Humans.Hero` `.ini` parser, whose file does not ship (`ITEM-SPAWN-026`), and `R0993`.
- Closure widened to the per-mission starter `R0099` and `R0512`, depth 6: 1731
  bodies, 1004 indirect calls not followed, 12 owners besides the container routines. Each of the
  12 is a walked body, so the closure shows a direct-call path from the start roots to it
  (depth 6 or less); the TSV carries no parent edge, so that path is not recorded, and it may be
  spurious (a fall-through or tail-jump walk). `R0993` is read whole. The other 11 are not
  read by this experiment: `R0946`, `R0945`, `R0907`, `R0461`,
  `R0472`, `R0262`, `R0061`, `R0407`, `R0014`, `R0430`
  and `R0448`. Their roles (the order dispatcher `R0061`, the pickup `R0448`,
  the script instants in `R0262`, `Give All` `R0407`, item creators, the chat
  handler `R0430`, the cast-slot return `R0014`) come from earlier claims
  this card does not cite, not from a read here. Whether any of them runs at a mission start is
  Unknown.
- Script corpus, both roots (EXP-0160 `corpus-nodes.txt`): the file holds item-reference nodes
  only, 50 instant nodes per root in 10 ALM files (7 of operation 12, 43 of operation 13) plus
  check nodes, and no instant of operation 11, 20 or 28. No node in it carries code `0x0e1c`. The
  filter and count were a third data pass, one over the preregistered budget. The three `Give All` nodes move a map unit's container into hero ordinal 10001
  (`TRIG-GIVEALL-025`); no shipped node authors `Drop all` (`TRIG-DROPALL-024`).

A save loaded at a mission start restores each actor with its own container record
(`SAV-CARRY-050`) and re-indexes the actors from the groups (`PARTY-ORIGIN-010`). The load path
was not traced by this experiment: its closure roots did not include the `Player::Serialize` load
arm.

**Confidence.** High for what the three routines read whole do (`R0423`, `R0823`,
`R0993`) and for the existence of the 45 direct `E8` calls. Medium for the negative over
the whole start: the closures do not follow indirect calls; 11 closure owners were not read; the
call census ran no raw dword scan of the four routine addresses (tables and callbacks), did not
look for tail `jmp rel32`, and does not see item moves through the base list methods on
`actor+0x7c` that bypass the four routines (`ITEM-CONT-004` explains why those routines keep the
load field, which makes a bypass unlikely and does not exclude it). Medium for the script census:
its population is the item-reference nodes above, and the pass was over budget. Unknown for the
SAV-load start path. The EN and RU
`rom.exe` are one image, and the script corpus is identical on both roots, so the answer does
not differ between them.

### PARTY-038

- The server cull keeps an actor iff `0x21 <= word[actor+0x0e] < 0x40` (`PARTY-ENDCULL-026`); the
  client cull keeps `CUnit+0x18c` bit 0 (`PARTY-CULL-004`). Mercenaries do not cross as objects
  (`PARTY-MERC-007`), and a mode-0 `Humans` actor does not cross (`PARTY-M20-031`). A culled
  actor is destroyed through `R1388` and `vt+0x04(1)`; slot `+0x04` of the `Human` vtable
  `L00003` is `R2195` (read whole), which calls `R2196`.
- Direct-call closure from `R2196` and `R1388`, depth 6: 69 bodies, 20 indirect calls
  not followed, and no call to the container add, insert, drain or take, or to the sack makers
  `R0871`, `R0944` and `R0472`. The stack is therefore not handed to another
  actor or dropped into a sack by any direct call; it leaves the party with its holder.
- The `0xbe` carry serializes one root actor (`PARTY-CARRY-005`), so on that path a companion and
  its stack do not cross.
- `R0993` runs only on `R0823`'s construction arm, `player+0x34 == 0`
  (`PARTY-037`). A primary that already exists gets no second stack at a mission start; a primary
  constructed while another member holds a stack would hold a second one.

**Confidence.** High for the cull membership (cited claims, each read at instruction level) and
for the construction-arm gate (`R0823` read whole). Medium that the stack is destroyed with
its holder rather than moved: the destructor closure is direct calls only, and the `vt+4`
destructors of the contained items were not read.

**Unknown.** Whether any shipped flow constructs a primary hero while a companion already holds
`Quest Documents`. Whether `Quest Documents` can reach the cast slot `+0x68`, where the end reset
deletes a class-14 item (`PARTY-M100-034`). The client's inventory mirror at a mission start was
not traced.

## Open questions

- The remaining bits of `+0x18c` (`PARTY-FLAG-003`), and the eleven sites that
  copy the whole word between drawables.
- What the three replaced sub-objects are (`PARTY-LOSS-006`: `+0x154`,
  `+0x158`, `+0x10`), and the two group references `+0x40`/`+0x44`
  (`PARTY-ROSTER-002`).
- A second campaign. Every figure about how many (five players, one carried
  root, fifteen mercenary types) is a fact about the one campaign this install
  ships. The structures are read from code and do not depend on it; the counts
  do.
- Whether the `0xbe` carry also fires on the campaign's mission-to-mission edge,
  and what builds the human `Player` before the campaign's first map load, the
  one boundary `R0423`'s reuse arm cannot serve (`PARTY-PERSIST-014`,
  `PARTY-PERSIST-028`).
- The four writers of `player+0x20` classified only by their append site
  (`PARTY-WRITE-011`): `R0151`, `R0062`, `R0003` and
  `R0432`. Each is one listing.
- `actor+0x13c`, the flag that decides whether `R0823` repairs a hero or
  leaves it alone (`PARTY-INSTALL-012`), and `actor+0x70`'s writers. One writer
  of `+0x13c` is named: the mission-end cull sets it to 0 on every survivor
  (`PARTY-ENDCULL-026`), the state that makes the repair arm return a carried
  hero untouched. What sets it non-zero, other than corpse decay, is not read.
