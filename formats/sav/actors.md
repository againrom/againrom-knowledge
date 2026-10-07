# SAV Unit, Humanoid and Human

[Format reference](format.md) · [Token](token.md) · [Human state](human-state.md)

## Unit programme

`Unit::Serialize` is `R0210`. Each row follows the previous one without
padding. Sources are offsets in the actor unless an indirection is shown.
`list32` and `Count` are different primitives. — SAV-UNITPROG-156

| Order | Wire | Source |
|---:|---|---|
| 1 | Token37 | Base Token |
| 2 | `list32<Effect>` | `+20` |
| 3 | `Count(n)`, `n*u16` | Static route at `+15c` |
| 4 | `Count(n)`, `n*u16` | Dynamic route at `+178` |
| 5 | raw 24 | `+a6` |
| 6 | raw 22 | `+be` |
| 7 | raw 24 | `+114` |
| 8 | raw 64 | `+d4` |
| 9 | raw 180 | Mover at `*(+154)` |
| 10 | raw 148 | Order at `*(+158)` |
| 11 | `Count(n)`, `n*u16` | List at `*(*(this+158)+90)`, inside order serializer `R0199` |
| 12 | u8 each | `+49,+4a,+4b,+4c` |
| 13 | raw 4 each | `+50,+54,+58` |
| 14 | u8 each | `+60,+61,+6c` |
| 15 | Two `objref` operations | `+74,+78` |
| 16 | CString | `+80`, name |
| 17 | u16 each | `+84,+86,+88,+8a,+8c,+8e,+90,+92,+94,+96,+98,+9a,+9c,+9e` |
| 18 | u8 each | `+a2,+a3` |
| 19 | u16 each | `+a0,+a4` |
| 20 | u8, u32, three u8 | `+12c,+130,+134,+135,+136` |
| 21 | u32, u8, u32, u32 | `+138,+13c,+148,+144` |
| 22 | `objref` | `+68` |
| 23 | u8 presence; if nonzero, direct container | `*(+7c)` |
| 24 | u8 presence; if nonzero, direct Spellbook | `*(+140)` |
| 25 | Four u32, then u8 | `+5c,+64,+44,+40,+48` |

Row 10's bytes `0x90..0x93` are Order `+90`, the pointer to row 11's list.
STORE writes its live bits; LOAD replaces it with a fresh list before row 11.
Any saved value, including 0, loads alike. — SAV-1113

Speed sits in row 17's `+8c` and in byte `+0a` of the row 9 Mover block. The
traced LOAD bodies do not recompute it, and SAVE writes both live values.
Whether the original re-derives speed after LOAD is Unknown; on one mission-140
pair the derive formula reproduces the saved values. — SAV-1116

The six raw block writes in rows 5..10 total 462 bytes. SAVE first copies
`u8 +14c` into `u32 +148`; this is a local SAVE mutation. Presence producers
emit 0/1 while the load arms accept every nonzero byte. The inventory presence
flag stores no container identity: a preconstructed container receives its
count, references, insertion index and load. — SAV-UNITPROG-156,
SAV-CITYSTORE-516, SAV-DEADLOAD-130

Humanoid appends raw XP24 at `+1cc`, twelve equipment `objref`s at
`+198+4*i` for `i=1..12`, then a Diary `objref` at `+1e4`. Human inherits
that whole programme and adds no bytes in its own store arm. Inventory,
equipment, Spellbook and Diary are separate constructs.
— SAV-HUMAN-043, SAV-CARRY-050 (death-container identity clause partially
retracted; the four mechanisms stand)

The array those twelve references come from is thirteen dwords wide, not
twelve. The Humanoid constructor zero-fills `+198+4*i` for `i=0..12`, then six
dwords at `+1cc+4*i` for `i=0..5`, then `+1e4` on its own — the same three
widths the store arm uses, from one routine. The destructor walks `i=1..12` and
deletes each non-null element, so the index the serializer skips is the index
the destructor skips: element 0 is allocated storage that the load arm never
writes, the store arm never emits and the destructor never frees. It is never
filled either, because the only routine that can place a non-null item there
refuses a slot number of 0 outright. What array indices 1 and 2 hold remains
Unknown.
— SAV-1028, SAV-HUMAN-043, SAV-CARRY-050 (death-container identity clause
partially retracted; worn-slot measurements stand), ITEM-EQUIP-006

A roster record carries one 24-byte stat layout twice, the 22-byte layout once
and the whole 64-byte modifier block, and one routine per layout is both its
write site and its read site. Before the store/load split,
`Unit::Serialize` reaches four fixed spans by adding a literal offset to the object pointer:
`+a6`, `+be`, `+114` and `+d4`. Three helpers serve them, and `+a6` and `+114`
share one. Each helper tests store-or-load, then pushes a single literal length
and calls either the archive's raw span write or its raw span read — `0x18`
for the 24-byte layout, `0x16` for the 22-byte one, `0x40` for the modifier
block. Store and load cannot disagree, and neither arm computes anything.

Under the three-copy layout, `+a6` is the live copy and `+114` the base copy of
one 24-byte stat record, `+be` is the live copy of the 22-byte defence record,
and `+d4..+113` is the modifier block carrying both layouts' modifiers and the
five scalar modifiers. The file therefore holds two records of one 24-byte
layout, `0x6e` apart. The scalar words are stored folded, not raw: the store arm
emits fourteen consecutive u16 from `+84` to `+9e`, then `+a0` and `+a4`, and
the fold's five scalar adds land inside that set.

For a consumer: both forms of the 24-byte layout are in the file. For the
22-byte layout and for the five folded scalars the file carries live and
modifier but no base, so a base is recoverable only by subtraction, and only as
far as the fold is additive — one field of the 24-byte layout is assigned
rather than added and is available only because the base copy is stored. The
spawn adjustment writes the live copies and not the base copy, so live minus
modifier is not in general the base.

The live/base/modifier naming is carried from the actor-layout rows, not
established on the wire; what the wire shows is one helper serving two spans of
one width. What the individual bytes of the 22-byte and 64-byte spans mean
beyond those field lists is Unknown. In the permitted save corpus (90 files, 1,360 actor
records) 14 records in 7 files carry a nonzero modifier damage-kind byte; the 7 files are 4
distinct files (three are byte copies), so these are 8 records from one dated series. All are Humans whose six
damage-kind bytes and six protection words are 100, the output pattern of the one routine that
stores there. Which path produced them is Unknown (`HERO-MODDK-161`).
— SAV-1031, SAV-UNITPROG-156, SAV-CITYSTORE-516, SAV-HUMLOAD-445, HERO-MOD-016,
HERO-MODDK-161, UNIT-CTOR-004, UNIT-GATE-013

A city-shape save's roster is measured end to end in decoded-stream
coordinates. One preserved 6,110-byte stream: head `0..75`, Player list
`75..5696`, the Player record `87..5696` whose fixed field run is `94..145`
with `+3c` at 119, two group records at `149..3041` and `3041..4942`, two Human
records at `248..3029` and `3131..4930`, the Player's own raw 32 at
`4942..4974` and its Diary at `4974..5696`. Between the list head and the first
Human there is a Player record and a group record and no count of characters:
membership is the nesting, not a roster object. Each Human's variable-length
parts tile the record with no residue under the programme above.
— SAV-1028, PARTY-ROSTER-002 (its `+0x1c` group-id clause is superseded by
SAV-GRPFLD-060; the three-level roster stands), SAV-UNITPROG-156

The actor Diary reference may resolve to an existing archive object. The
typed LOAD does not establish exclusive ownership; a new object is registered
before its virtual serializer. The bounded actor-method search reaches SAVE
and destruction but no ordinary independent Diary array consumer. — SAV-982,
SAV-984

Actor `+1e4` also equals `+1cc+4*6`. The two selected progress branches lack
a local index<=5 check, so index6 would address the Diary pointer word as a
scalar. Native reach of that index and a resulting pointer change remain
Unknown. It defines neither a seventh skill nor Diary array meaning. — SAV-983

## Scalar meanings

| Runtime field | Wire width | Established meaning |
|---|---:|---|
| `+84,+86,+88,+8a` | 2 each | Body, Reaction, Mind, Spirit |
| `+8c` | 2 | Speed |
| `+8e` | 2 | Actor's own carried weight |
| `+90` | 2 | Derived load: own weight plus half the container's running weight |
| `+92` | 2 | Capacity |
| `+94,+96,+98` | 2 each | Signed health, maximum and regeneration period |
| `+9a,+9c,+9e` | 2 each | Signed mana, maximum and regeneration period |
| `+a2,+a3` | 1 each | Separately saved regeneration remainders, unsigned on reload |
| `+a4` | 2 | Sight radius in 1/256-cell units |
| `+12c` | 1 | Range, including equipment delta |
| `+130` | 4 | Aggregate experience, simulation i32 |
| `+13c` | 1 | Actor stage |
| `+6c` | 1 | Signed timer; countdown while dying |

The archive restores weight, load and sight without recomputation. The named
Unit constructor capacity is based on default Body 30 and stays 300 on that
arm; that is not a general relation to restored Body. The direct weight-delta
helper adds a 16-bit delta and recalculates load. Its 64000 threshold is a
signed compare; a wrapped negative load removes the speed penalty in the
located arithmetic; native acceptance of such loaded values remains Unknown.
— SAV-UNITFLD-049, UNIT-CTOR-004, SAV-635, ITEM-LOAD-005, HERO-SIGHT-007,
AI-SIGHT-092, SAV-792, SAV-793, SAV-REGENWIDTH-528, SAV-REGENWIRE-532

A stored load of 181 with own weight 178 and an empty container can survive
original load/resave; recomputing would destroy that state. Sight is not
restricted to whole cells. Its named constructor gives `0x0500`; LOAD can retain
zero without that becoming a general constructor/default value. — SAV-794, SAV-795, SAV-796

## Skills and experience

| Source | Wire position | Meaning |
|---|---|---|
| `+a8+2*i`, `i=0..5` | Bytes 2..13 of raw-a6 | Six u16 levels |
| `+130` | Scalar row 20 | Aggregate XP, simulation i32 |
| `+1cc+4*i`, `i=0..5` | First 24 bytes of Humanoid suffix | Six XP values, simulation i32 |

The level word and XP dword share index i. Initializers and award paths update
them and the aggregate, but LOAD copies all three independently. The word at
`+a6` is not level slot zero. `aggregate=sum(XP)` is not enforced by LOAD. The original programme fixes six slots; extending
the XP tail shifts the following references and requires an external versioned
extension. — SAV-HEROXP-063, SAV-HEROSKILL-064

Weapon-borne spell awards use the same fields; no separate item-cast skill or
award-source tag is stored. Progress belongs to the containing Human object,
not its Group position. Player `+34` selects a primary character without
changing other members' shape. — HERO-ITEMSKILL-096, SAV-HEROID-065

## Definition and presentation binding

The restored row byte `+0c` and type word `+0e` are separate selectors. Let
`U=[L08211]` and `H=[L08213]` be Units/Humans array bases.

| Exact loaded class | Definition pointer `+3c` after local serializer return |
|---|---|
| Unit | `U + 48*u8(+0c)` |
| Humanoid | 0 |
| Human | `H + 48*(u16(+0e)<33 ? u8(+0c) : 5)` |

The arithmetic wraps at 32 bits and has no collection-size check. The fixed
Human row-5 branch covers every zero-extended word 33..65535. These branches
do not rewrite the selectors. Class Item has a size-checked lookup; Armor,
Shield and Weapon have an unchecked stride-`0x3c` lookup (see
[Items](items.md#definition-binding)). — SAV-ACTORBIND-544, SAV-1088

Unit copies lowByte(`+148`) to `+14c`; exact Unit then clears the byte,
Humanoid preserves it, and Human retains a nonzero byte only with a nonnull
definition whose name contains `NPC`. The substring need not start the name;
the suffix preserves values such as 2/255 rather than Booleanizing them.
The dword remains unchanged until later SAVE widens the effective byte into
it. A sender uses the byte to select a separate collection at `L02112`.
Other character modes and full runtime remain outside this local result.
— SAV-ACTORDISPLAY-545

Archive creation invokes constructors first. Human's creator uses
`R0656("Man_Unarmed",0,0)`; constructor definition/equipment/type setup is
not rerun by the load suffix. Actor `+154/+158` allocation identities are
fresh; their 180/148 bytes come from SAV, and the order helper replaces its
`+90` list. These constructor outputs are not post-load defaults.
— SAV-ACTORCTOR-546

Footprint/domain `+49/+4a`, face/class `+4b/+4c` and mover mask `+5` restore
independently. That mask selects the passability plane. — TERR-PASS-051 (amended, superseded)

The selected post-read hooks do not normalize their consistency.
Stage zero admits key repair in the fresh order object for
`+0c,+10,+18,+20,+28,+30,+68`; hits replace, misses retain raw words, without
class/lifetime checks. Other stages skip that repair. — SAV-ACTORINPUT-547,
SAV-HUMRESUME-460 (reference-repair shorthand superseded by this field-specific
rule)

The repaired list excludes both the pending-order byte `ord+0x08` and the order-progress byte
`ord+0x09`; neither is touched by stage zero's own repair or by any other located LOAD-side
writer, and `Order::Serialize` copies its own fixed-size block through one raw archive
read/write pair, so a saved value at either offset sits unchanged after LOAD. A dying,
not-yet-torn-down actor (signed `+94 <= 0`, stage `+13c != 0`, `+54 != 0x10`) carries `ord+0x08`
confined to `{0x00, 0x0b}` and `ord+0x0c` zero across the permitted archive's own 48 such
records — 59 saves read from 7 of the eleven pinned directories, 4 original and 44 ROM1 resaves
of Againrom-produced/modified documents (`SAV-1059`) — because the per-tick routine's dying
branch never reaches the order machine that would otherwise consume either field, on any dying
tick. A crossing actor is action state `+54 == 1` **and** order-progress `ord+0x09 == 3`
together, not state `1` alone: `MOVE-STEP-040`'s own progress-3 arm is what sets state 1 for a
transit, and state `1` is also the arrival state of two idle-turn order arms at a different
progress value. It carries its own attack-cycle countdown `actor+0x6c` — an actor
field, outside this repaired order-object list regardless — cancelled, not frozen: every
state-`1` tick stores the attack sub-phase `+58 = 0` and never reaches the dispatch span that
reads or advances `+6c`, so a crossing actor's own countdown sits stale through the crossing and
is overwritten, not resumed, at the next charge start. Of the permitted archive's alive
population, 38 state-1/progress-3 (crossing) records exist, and every one carries
`actor+0x6c = 0`: the corpus holds no example of a crossing actor with a live countdown. —
HERO-DYINGTICK-145, HERO-CROSSHOLD-146

A creature's three class-spellbook slot pairs (`ord+0x78`..`+0x8c`, [MAGIC](../magic/casting.md),
`AI-341`) sit inside `Order::Serialize`'s own raw `0x94`-byte span but are not among the
key-repaired offsets above, so LOAD carries whatever the archive stored there unrepaired. No writer
other than `R0184`'s spawn setup and `Order::Serialize`'s own raw LOAD copy was found; the
disambiguation covers the AI module — so the carried value and a hypothetical class-derived rebuild
are currently byte-identical whenever nothing else has changed the slots since spawn (`SAV-1066`).

Exact Humanoid's null definition is not established as safe: its shared Unit
`vt+58` uses `+3c+8` without a local null guard. Exact-class acceptance, loaded
reach and the first frame/move/save chronology remain Unknown.
— SAV-ACTORLIMIT-548

## Mover and first event

Mover `R1348` transfers 180 bytes without field normalization, including
current/desired bytes `+0/+1`, rotation speed `+a`, counter byte `+9d`, active
dword `+a0` and estimate byte `+a4`. Unit passes its `+154` receiver and the
same archive. The selected Unit/Humanoid/Human serializers contain no later
direct replacement of rotation-speed byte `+a`; their virtual/embedded calls
and later Human derive are separate mutation boundaries. — SAV-TURNLOAD-822,
SAV-MOVRATE-866, MOVE-TURN-044

World LOAD rebuilds Player and global actor membership using the same actor
pointers, excluding actor `+4c` mask `0x08`. The insertion helper `R0032`
does not include the creator helper's `+4` assignment. Post-load `actor+24`
uses the base hook; stage zero reaches mover `+7c` reference repair. These
selected bodies provide no unconditional derive-before-read guarantee.
— SAV-LOADREG-878, SAV-LOADHOOK-879

Common tick processes attached Effect `+38` before HP/order admission.
The selected pending-order-10 turn requires positive signed HP, nonzero
`+3c` and its admitted order prefix. Its inactive snap skips the allocated
byte read; an admitted continuous Effect can call actor `+50` first.
Frontend, world, phase and command callbacks precede the selected actor suffix,
so the first actual restored mover event, its value and native timing remain
Unknown. — MOVE-EVENT-060, MOVE-EVENT-061, SAV-FIRSTMOVE-880

Because the mover is transferred raw, a document with rotation speed `+a` = 0
terminates the original with an integer divide by zero at `L09136` on the
first turn of an ordinary move. Original Humans carry 15..22 and Units 8..22.
A written mover block needs a nonzero rate. — SAV-1092

### Start state a written mission document needs

Hero control on the first LOAD depends on the document, not on the process
that loads it. A generated mission-10 document without control regained it
when the human participant's subtree took an original resave's values: hero
sub-cell `+4/+5` 128,128, no pending move order, an empty `+0x178` word list,
the hero's 462-byte block after that list and its 19-byte block,
Player body `+57`, group AI `+0x48` and session wire 1400..1415. Player body
`+57` alone and the group's own AI block `+0x48` alone are each individually
insufficient; the remaining hero Token/block/word-list/session subgroup, taken
together, is sufficient. Which single byte within that subgroup decides control
is Unknown. — SAV-1095, SAV-1100

An AI start state of zeros (post 0,0, 19-byte block `+4` 0, mover `+08/+09`
0,0, group centre and guard radius 0, diplomacy template 0) coincides with an
AI fight at mission start that the original start does not show. Copying the
mission-start original's values removed it. The driver is the actor's own
16-byte personal AI-start record, not its owner's AI group's 12-byte block:
reverting the personal record alone reproduces the fight regardless of the
group's state, and reverting the group alone does not, regardless of the
personal record's state. — SAV-1096, SAV-1101

### Client drawable after LOAD

The selected entry state message does not populate every numeric card field.
A non-owned Human outside types 33..63 receives mask `50fb`, preserving
current/max health and mana but omitting primary and combat values. Types
9/24 also take no prototype-stat copy. The subsequent numeric card reads
drawable bytes; local omitted fields retain their existing contents.
Other client writers and first-frame contents remain Unknown. — SAV-1125,
SAV-1126

One external mission-151 SAV and its reported original resave, 546 full-clock
units apart, retain the primary values of five Human candidates matching
owner screenshots by class and maxima. Those primary values differ from
the displayed ones. Screenshots do not identify a unique actor, and an
original resave does not establish a regenerated drawable cache. — SAV-1127

A native mission 151 control and its entry autosave share217 unique nonzero
authored ids with that external/resave pair. External/native primary words,
attack, maximum health, type and definition index match on all 217. Other
actor/member differences remain, including Human+1e4 Diary references and
live scalar values. Five zero-id actors per file are not identity-matched.
Exact member-span coverage is not complete semantic equivalence of all
document records. — SAV-1128

The four selected targets' absorption 200 is a map-script assignment to
actor+c0. Native observations0/10 do not establish a required saved default:
the repeatable setter coexists with delayed temporary casts whose effect
application can invoke Human derive, clear the defence/absorption block and
fold equipment modifiers. The last runtime writer at SAVE remains Unknown.
The separate +be field is defence; the +c0 server field is a word, even
where client projection reads only its low byte. — SAV-1128

The session-entry join sends every tick-list actor, and every dead-list actor
below stage 5, through the actor-state sender with a full mask. The client
state arm creates a drawable for an unknown runtime id whose message passes
the owner, class and `0x4001` tests. The drawable's position is the one the
message names. — SAV-1118

The sender names the Token Position only when both sub-cell bytes `+4/+5` are
`0x80`. Otherwise it names the mover word `*(+154)+06` as a cell centre. An
actor that is still mid-step when the entry send reads it gets its drawable
at the mover cell while the simulation keeps it at its Position. For a
restored actor that is the saved mover cell only if no tick before the entry
send moved it. The pre-entry tick runs when campaign `+0x6b8` is nonzero, and
the number of ticks before the send was not established. Whether the client
shows the drawable was not read. Ordinary move messages add deltas to the
drawable and do not restore an absolute position. — SAV-1119

Mission-141 documents of this project carry three Units with a non-centred
sub-cell and mover word 0, among them map unit 57; unless the pre-entry tick
moves it first, its drawable is created at map cell 0,0 (Medium). None of 76
distinct original-written documents of the owner save tree holds an actor
whose mover cell is more than one cell from its Position (Medium).
— SAV-1120, SAV-1121

## Dead actors and source references

Dead actors leave the owner graph and remain in the exact top-level dead list.
LOAD creates/resolves them through CArchive, then the later manager callback
runs `actor->vt+24`, repairing Position and `+5c,+64,+44,+68,+40` references.
It does not join an authored ALM record or derive actor state. Timer `+6c`,
signed health `+94` and stage `+13c` remain separately restored fields.
— SAV-DEATH-051 (tag-scan counts superseded by SAV-DEADLOAD-124; removal from
the owner graph stands), SAV-DEADLOAD-124, SAV-DEADLOAD-125, SAV-DEADLOAD-126

Stage-5 health is not restricted to -10001: the decay consumer retains every
signed value below -10000. Producers of the other observed residues remain
Unknown. Death moves/adopts or drains the old inventory into a Sack and gives
the corpse a fresh container; the saved corpse container is not the Sack
container identity. — SAV-DEADLOAD-127, SAV-DEADLOAD-129, SAV-DEADLOAD-130,
ITEM-DEATH-012 (its unconditional container-identity sentence is retracted;
ITEM-GROUNDMOVE-130 holds both Sack branches)

Unit `+40` is a raw old actor-address key; `+64` is temporary attribution;
`+68` is initially an archive object reference; signed `+48` is an attribution
byte. World lifecycle separately map-repairs `+40/+64/+68`, nulling misses;
no-world LOAD clears `+40`, and `+48` is retained at that boundary. The `+68`
second lookup uses its then-current word, so archive resolution alone does
not prove its survival. Ordinary item-cast completion clears `+64/+68`.
— SAV-908, MAGIC-ITEMKILL-117, ITEM-CASTSTATE-056 (its `+68`-only remap clause
is retracted; the cast lifetime stands)

Actor word `+50`, the AI state, is rewritten at death in two steps. The per-tick routine returns at once for
`+54` = `0x10`; otherwise it stores 0 into `+50` and `+54` on every tick with signed health at or below 0. Teardown
stores `0x10` into both and the actor moves to the dead list. A dying record therefore stores 0 and a dead-list
record outside every group stores `0x10`. LOAD copies the word as four raw bytes with no test, and the post-load
rebind and its direct-call closure hold no operand at `+50`. An actor in a group is appended to the tick list by
World LOAD unless `+4c` bit 3 is set; a dead one with `+54` not `0x10` gets 0 at its first tick, while a dead actor
only in the dead list is never ticked and keeps the loaded value. The one witnessed load is five records, `0xb` in
and 0 out, in one engine-to-original resave lineage. Other values, a grouped dead actor with `+54` = `0x10`, and the
producer of the exception records are Unknown. — SAV-1169, SAV-1170, SAV-1171

### Consequence receivers

The selected Unit callbacks `+48/+60/+64` are empty. Humanoid/Human use
`L08088/R1054/R0870`; the latter two keep the source actor as the
receiver of progress slot `+5c ->R0810`. Constructor and vtable provenance
identify these classes. A restored address-map hit or actor type word does
not validate an arbitrary receiver's class or allocation lifetime. — SAV-958

The source callback pair belongs to actor manager `R0427`, after its
admitted actor tick/removal work. Humanoid/Human `+48` can add victim `+1c`
to source-owner Player `+38` and enter notification dispatch. It bypasses that
helper for server `+0c==0`, a nonzero victim class predicate, or negative
source HP. The manager then rereads signed victim `+48` and victim `+40`
for `+60`, without another null/type/health check. Notification callbacks
remain an unresolved mutation boundary. — SAV-959

A `+60` refusal does not necessarily suppress the manager's later accounting.
After it returns, the manager rereads the current source and owner, updates
that Player's selected counter, and calls the Player-owned Diary with the
victim. Missing victim owner or the paired multiplayer/owner gate can refuse
progress while that later Diary operation still runs on the admitted prefix.
Arithmetic and notification cuts limit the joined evidence to Medium;
native awards and first post-LOAD recipients remain Unknown. — SAV-960

Before the current actor's tick, the manager clears nonzero actor `+40` when
the referenced source's `+14` is null. It dereferences that source to make
the check; this is not an allocation-lifetime test. — SAV-962

Before the callback pair, the removed victim's attached-list loop sets each reached
Effect's counter `+42=1` and source `+44=0` before its tick callback. This is
not a global walk of Effects referring to the removed actor. Nested removal,
derived Effect and notification callbacks, actual destruction, full LOAD
ordering and the `+68` second lookup remain open. — SAV-962
