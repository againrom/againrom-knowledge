# SAV — save game

Claims about the ROM1 save game (`.sav`, magic `Asg&`): its simulation
document, application state store and campaign record, the original serializers
and LOAD paths that write and restore them, and the producers, consumers and
lifetimes of the state they carry. Spec:
[`formats/sav/format.md`](../formats/sav/format.md). Format of this file:
[registry.md](registry.md).

## Diary notification receiver and local refresh

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-994 | The Diary message selects a connection before it serializes opcode, count and the words; a 17-word message is 39 serialized bytes. | High / Medium | ✔ promoted | [EXP-0345](../experiments/EXP-0345-diary-notification-refresh/) |
| SAV-995 | Opcode 186 replaces the dispatch receiver's client CWordArray without another Player lookup. | High | ✔ promoted | [EXP-0345](../experiments/EXP-0345-diary-notification-refresh/) |
| SAV-996 | Two shared character/unit-panel readers use the received words as local disclosure selectors. | High | ✔ promoted | [EXP-0345](../experiments/EXP-0345-diary-notification-refresh/) |
| SAV-997 | Resume setup, later Diary mutation and a text-command branch are distinct local refresh producers. | High | ✔ promoted | [EXP-0345](../experiments/EXP-0345-diary-notification-refresh/) |
| SAV-998 | Static receipt and panel consumption do not close native Diary refresh chronology. | High / Unknown | ✔ promoted | [EXP-0345](../experiments/EXP-0345-diary-notification-refresh/) |

### SAV-994

- The static packet at `L07299` is constructed by `L08054` through
  `L08055`, which installs vtable `L08056`.
- Transport `R0217` uses packet `WORD+7` for its addressed/broadcast branch.
  Addressed lookup `L08057` searches `manager+18b8` by connection `WORD+4`; a
  miss may fall back to `manager+18b4` when `manager+18b0` is zero.
- Concrete packet `virtual08=L08058` passes `packet+9` and
  `5+2*DWORD[packet+0a]` to connection writer `L02447`;
  `virtual10=L08059` returns the same length. The 17-word builder therefore
  supplies 39 serialized bytes and does not serialize `header+5`/`+7`.
- Client retrieval `R0513` iterates connections, `L02424` obtains an
  opcode, and `L02428` maps 186 through two tables to `L08060`. It selects
  the same static packet and `virtual0c=L08061`, which reads four count
  bytes, then `2*count` bytes when count is positive.
- Evidence: complete transport/packet bodies, tables, independent operand
  controls and typed round-trip cases. Supporting claim: `SAV-845`.

**Confidence.** High for the concrete local receivers, constructor/vtable
operands and framing. Medium for the two original-instruction round trips
joined by supplied streams; differing sender IDs produce identical synthetic
39-byte streams.

**Unknown.** Connection queue, refill, flush, synchronization and native
delivery remain cuts, not a claimed transport guarantee.

### SAV-995

- Client `R0509` uses `opcode-3`, byte table `L02523` and the dword table
  `L02524`: 186 selects index 43 and arm `L08062`.
- With R equal to the saved entry receiver (`this`), the arm sizes `R+3f58` from packet
  `DWORD+0a`, then copies `2*count` bytes from `packet+0e` to
  `DWORD[R+3f5c]`. The complete arm neither reads `packet+5`/`+7` nor selects
  `R+9b4`.
- Constructor `R0391` calls `R1549` at `R+3f58` and installs receiver
  vtable `L03067`; `R1549` installs the CWordArray vtable and clears its
  pointer, count, capacity and growth.
- Frontend `R0315` allocates `0x3f78`, constructs R and stores it at
  `frontend+d0`; `R0260` passes that pointer into `R0509`.
- Evidence: client switch, member constructor, receiving cases and call
  frontier.

**Confidence.** High for the local selection and stores, the complete
constructor/binding instructions, count 0/1/17/18 controls, distinct R/OTHER
objects and replacement after `ffff`. The vectors supply dispatch locals and an
existing capacity of 64; the zero-count free return is supplied. No native
allocation, delivery, callback preservation or first-after-LOAD order follows.

### SAV-996

- Both `R0877` and `R0916` obtain R from client drawable `U+e0`.
- For signed `U+20>=64` and `index=U+20-64` below signed `R+3f60`, they read
  `WORD[DWORD[R+3f5c]+2*index]`, shift by the x86 count from
  `4*(byte[U+24]-1)`, and retain its low nibble.
- This replaces their earlier owner/visibility fallback, rather than ORing or
  maximizing it. Nonzero `L01661` then overrides the selector with 7.
- The drawing body gates groups of formatting/draw calls at selector >0
  through >6 and has additional exact ==7 gates; the pointer-position reader
  gates the corresponding text returns. Nibble 15 is not clamped to 7.
- Evidence: complete reader bodies, operand controls and selector cases.
  Supporting claims: `TEXT-UI-034`, `TEXT-UI-035`.

**Confidence.** High for the two complete local bodies and 12 distinct
original-prefix controls, each executing both readers. The prefixes stop before
drawing or coordinate dispatch. The draw/text calls, their established
captions, every `U+e0` producer, native selected-drawable identity and when a
panel redraws are separate boundaries.

### SAV-997

- On normal return through the `Player+34`-nonnull arm of `R0131`,
  `L08063` calls `L08064` with that same Player and server manager
  `L00522`. This call is outside the setup argument's optional
  initialization branch, after earlier setup/projection/notification calls.
- `SAV-948`'s opcode 4 lookup therefore supplies a conditional resume route; it
  does not prove delivery or the first refresh.
- Later attributed accounting retains the conditional Diary notification of
  `SAV-845` and `SAV-960`.
- The other builder site, `L08065`, belongs to `R0430`. Its server caller
  is an opcode 145 text arm with an existing Player and a leading `#`; the
  local command branch uses `'#modify '`, `'self'`/`'army'` and
  `'+knowledge'` prefix predicates. It passes the supplied Player to the
  builder and does not by itself establish an ordinary LOAD/event trigger.
- Evidence: setup/command bodies, server/client tables, literal hashes and
  frontier.

**Confidence.** High for the complete caller bodies, branch membership and
final Player/manager arguments. The three E8 builder callers are a finite
original-image candidate census, not proof of no indirect producer.

**Unknown.** String-service returns, earlier returning callbacks, accepted
native command syntax, every LOAD continuation and absolute event order remain
bounded frontiers.

### SAV-998

- The selected population: 35 complete bodies, 15309 independently decoded
  instruction rows, 1125 calls including 96 computed calls, nine table spans
  with 755 entries, and 20 original-x86 synthetic cases.
- A file-backed scan yields 83 raw E8 candidates to five roots, no stored
  root-address candidates, and five immediate 186 candidates among 554699
  linear `.text` instructions with 23 skipped bytes.
- The declared Ghidra offset search yields 10 candidates, not a global consumer
  census.
- Receipt only resizes/copies the client array; the complete opcode 186 arm
  makes no direct draw, invalidation or message-post call.
- Evidence: population manifests, call frontier and falsifiable native
  prediction.

**Confidence.** High for the enumerated population and the local arm. The
synthetic stream and the supplied dispatcher/panel memory are not an original
runtime witness. No original process or save corpus was used.

**Unknown.** Connection delivery/ordering, partial stream failure and
static-packet interleaving, all callback preservation, complete
drawable-to-client/Player lifetime, the first native post-LOAD refresh and
visible refresh delay.

## Actor-owned Diary receiver frontier

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-982 | Within a bounded actor receiver frontier, the actor Diary member has archive and destructor escapes but no ordinary actor-owned Diary array consumer. | Medium | ✔ promoted | [EXP-0344](../experiments/EXP-0344-actor-diary-consumers/) |
| SAV-983 | The actor Diary pointer word overlaps a conditional progress index, without establishing a Diary array use. | High / Unknown | ✔ promoted | [EXP-0344](../experiments/EXP-0344-actor-diary-consumers/) |
| SAV-984 | Actor Diary LOAD can return an archive alias as well as a newly serialized object, so the member route does not prove exclusive ownership. | High / Unknown | ✔ promoted | [EXP-0344](../experiments/EXP-0344-actor-diary-consumers/) |

### SAV-982

- Constructor stores independently bind Humanoid/Human tables
  `L00002`/`L00003` and Diary `L08066`.
- Population: the 56 actor slots `+00..+6c`, member lifecycle, copy, archive
  and retained displacement candidates yield 85 complete bodies plus five
  single-instruction stack windows: 7351 independently decoded instructions,
  573 calls including 55 computed, and 1270 additional target/table/member
  controls, on one byte-identical EN/RU code population.
- At `L08067` the `actor+1e4` pointer receives `virtual+04`, conditionally
  `L08068` then `L08069` for the constructor-built Diary; this forwards its
  embedded word/dword collections to their destructors.
- SAVE `L08070` forwards that same member through `R1110` to `R1129`,
  whose first-object branch queries slot `00` and calls slot `08`.
- The 28 retained displacement candidates have local provenance categories:
  seven actor member sites, six embedded-collection sites, four cleanup thunks,
  five literals, five stack operands and one global-rooted load. These
  categories are not a whole-program receiver census.
- Evidence: complete bodies, receiver paths, call frontier and independent
  controls.

**Confidence.** Medium for the bounded negative. Whole-actor calls, arbitrary
aliases, dynamic classes, callbacks and rebased indices remain outside
transitive closure. Neither global dead storage nor pointer preservation
follows; `SAV-983` gives a specific overlapping address expression.

### SAV-983

- In Humanoid/Human slot `+5c` target `R0810`, the `actor+4c &4` predicate
  `R0442` selects either positive stack argument at frame offset +10 or unsigned
  `actor+b6`, with a selected signed skill-word check below 100.
- The indexed load/add/store and comparison at
  `L08071`/`L08072`/`L08073` or `L08074`/`L08075`/`L08076` use
  dword `actor+1cc+4*k`; neither branch locally caps k at 5.
- The independently decoded displacement/scale give k=6 at `actor+1e4`. The
  local admitted load/add/store therefore treats that pointer-sized word
  numerically if k=6 reaches it. It does not dereference either Diary backing
  array.
- By contrast, initializer `R0899` uses 0..5 and the aggregate `R0978`
  uses 1..5, ending before the pointer.
- Evidence: rebased aliases and complete progress/caller instruction bodies.

**Confidence.** High for the actual addressing, branch predicates and
arithmetic overlap only. No original instructions were executed. This is not an
extra legitimate skill slot.

**Unknown.** Native k=6 reach, the supplied award value, a loaded pointer
change, actor lifetime and any downstream array effect.

### SAV-984

- Humanoid LOAD forms `actor+1e4` and calls `L08077`, which requests
  descriptor `L08078` through `L05928` and stores the returned pointer.
- The existing-object arm reads the archive load-table entry and applies
  `R0214`/`R0216` to a nonnull typed result.
- The new-object arm obtains the descriptor factory result, registers it at
  `L08079`, then calls its `virtual+08` at `L08080` before returning it.
- Constructor-built Diary table `L08066` selects `L08081`; the derivation
  check accepts a compatible descendant rather than demanding exact descriptor
  identity.
- Copy `R1550` separately forwards `source+04`/`+18` to two array-copy
  helpers, each using its own source count, but no actor-member source for that
  copy is established.
- Evidence: typed archive and copy receiver paths. Supporting claim:
  `SAV-774`.

**Confidence.** High for the local typed member/branch/registration order and
the copy requests. `SAV-847` remains the independent array LOAD authority.

**Unknown.** Compatible dynamic classes, unique allocation/ownership, callback
effects, copy primitive effects, the first native loaded consumer and an
independent actor-owned array purpose.

## Terminal UI continuation and SAVE controls

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-970 | The selected terminal handler binds a credits child and the common UI manager. | High / Unknown | ✔ promoted | [EXP-0343](../experiments/EXP-0343-terminal-continuation/) |
| SAV-971 | The selected terminal close chain requests FAME and then a mask-gated reset/menu route. | High / Unknown | ✔ promoted | [EXP-0343](../experiments/EXP-0343-terminal-continuation/) |
| SAV-972 | Retained terminal presentation masks inhibit the selected F2/SAVE-dialog path, but that does not close all SAVE reachability. | High / Unknown | ✔ promoted | [EXP-0343](../experiments/EXP-0343-terminal-continuation/) |

### SAV-970

- UI construction `R0315` stores the `L08082` result at `frame+11c` and the
  `R0773` result at `frame+cc`; their final vptrs are `L08083` and
  `L03640`.
- In `L08084`, manager add `R0385` precedes child slot 80 -> `L03793` and
  manager slot 34 -> `R0366`.
- Credits activation selects the identifier `main\text\credits.txt` through
  `R1551`.
- After those calls return, the handler clears `frame+41c` and ORs DWORD
  `frame+3dc` with 40.
- `Frame+11c` is a UI pointer, distinct from campaign `LastMission+11c` at
  `frame+664`.
- Add publishes a mutable list/parent alias; selected base activation and
  manager traversal have further callbacks.
- Evidence: receiver bindings, complete selected bodies, pointer words and
  Ghidra memory comparisons. Supporting claim: `SAV-890`.

**Confidence.** High for instruction order, constructor assignments, slot words
and named resource selection in the matched ROM1 build.

**Unknown.** Resource success, callback preservation and natural final-main
arrival.

### SAV-971

- Terminal keyboard prefix `L03799` sends 445 to its own event slot.
- Base modal event `R0716` handles 445/446 with `modal+5c` nonzero by
  calling slot 84 before queuing 44c with that modal as sender.
- Delivered 44c -> `R0709` with `sender=frame+11c` clears bit 40 and reads
  DWORD `frame+664`; nonzero requests 429, zero requests 421.
- Message 429 binds `frame+108` to FAME activation `L03690` and subsequently
  sets bit 1000. Its selected close sender branch clears bit 1000 and requests
  421.
- Dispatcher `L08085` requires the entire `frame+3dc` DWORD to be 0 before
  `R0816`; that body calls campaign reset `R0757` before later menu
  activation.
- The reset clears named fields and requests `R0755(10)`, the ordinary reset
  request, not a terminal number.
- Evidence: close chain, frame dispatcher, selected reset/menu bodies and
  complete call frontier. Supporting claims: `FAME-DISPLAY-012`,
  `SAV-CAMPAIGN-085`.

**Confidence.** High for the conditional message/store/receiver ordering. This
does not prove an unconditional natural credits-to-menu route or a retained
terminal city.

**Unknown.** `Thread+7c` main-window resolution, actual queued delivery, widget
dispatch, cleanup/resource return, intervening aliases, retained LastMission
and the zero-mask prerequisite.

### SAV-972

- The frame message map and key tables bind F2 to `L06466`, which requires
  `frame+3dc` in 0/1 and `mode+6bc=2` before 41a; dispatcher 41a independently
  requires mask 0/1 before constructing the `frame+12c` SAVE dialog.
- Credits bit 40 and FAME bit 1000 violate both mask conditions while retained.
- In `R0709` a sender matching `frame+12c` and DWORD `sender+60=445` reaches
  the aligned `L03820` -> `R0084` SAVE call.
- The other aligned call, `L03821`, is in dispatcher message 442 and has no
  local mask test; its normal terminal producer/order is unproved.
- Selected LOAD-dialog completion requests 419 -> `R1284`, then the existing
  document loader and world/no-world continuation.
- Evidence: complete SAVE callers, key/message tables, dialog vtables, LOAD
  wrapper and explicit producer frontier. Supporting claim: `SAV-892`.

**Confidence.** High for aligned callers, control bindings and local predicates.
No native terminal output or forced-handler substitute was used.

**Unknown.** Full normal terminal SAVE availability, the accepted
filename/label, current or retained campaign/documents at SAVE, output world
presence and restart/LOAD acceptance.

## Ordinary Human producer and lifetime edges

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-934 | The named skill hit-test supplies indices 1..5, which excludes the residual aliases on that local input path. | High / Medium | ✔ promoted | [EXP-0340](../experiments/EXP-0340-human-producer-lifetime/) |
| SAV-935 | The expanded fold and spellbook edges distinguish Human reads from Spell writes. | High / Medium | ✔ promoted | [EXP-0340](../experiments/EXP-0340-human-producer-lifetime/) |
| SAV-936 | The live builder exposes the same Human to equipment, collections and publication before return; its first mission SAVE remains unobserved. | Medium / Unknown | ✔ promoted | [EXP-0340](../experiments/EXP-0340-human-producer-lifetime/) |

### SAV-934

- `R1216` compares one raster byte against five marker bytes and returns the
  first index 0..4 or -1.
- `R1552` rejects -1, retains the index and forwards index+1 to `R0832`.
- When current-thread `virtual+7c` returns the same object R, setter receiver
  `R+420` and `member+a0` address `R+4c0`, also read by packet producer
  `R0822`. The setter writes a nonzero input without an upper bound.
- Draft initialization `R1553` supplies literal 1; the three
  attribute/preset callers instead forward stored `selection+134` plus 1
  without a local range check.
- Two complete 256-pixel suffix populations, including duplicate markers,
  confirm the first-match return range.
- Evidence: UI instructions, caller/reference rows and hit-suffix vectors.
  Supporting claims: `HERO-GENERAL-088`, `SAV-899`.

**Confidence.** High for the original local loop, index transformation, literal
and offset relations. Medium for the bounded ordinary-input provenance:
current-thread/window/raster calls, preserved draft identity, other
selection/draft writes and later packet admission remain explicit cuts. The 512
suffix cases do not execute those calls or establish an ordinary producer of
indices 10/25/32/39/40/42.

### SAV-935

- Derive passes `actor+d4` and actor to `R0840`; its defence and attack calls
  fold `actor+fe` into `actor+be` and `actor+e6` into `actor+a6`.
- The four complete fold bodies write none of `bc/bd`, `da/db`, `e8/e9`, `f6`,
  `f7/f8` or `fc/fd`; `da` feeds live capacity and `f7/f8` feed the live
  secondary damage bytes.
- Derive passes the loaded book pointer `u32[actor+140]` as `this` with that same
  actor as its argument; the book traverses nonnull Spell entries from index 1.
- `R0904` receives that Spell as `this` and the Human as an argument, reads
  signed `word[actor+a8+2*rowParameter2]` using a full 32-bit row index, and
  forwards a clamped derived byte to `R0625` on the Spell receiver. That
  setter's member stores are `Spell+09`, `+0e`, `+0f` plus word `+10`.
- Eighteen original-instruction book-to-setter-cut cases preserve the entire
  supplied Human and reach the Spell receiver.
- Evidence: fold/book instructions, 18 spellbook-cut vectors and call frontier.
  Supporting claims: `SAV-HUMFOLD-446`, `SAV-898`, `SAV-900`.

**Confidence.** High for the local receiver, width and fold relations, and the
explicit pre-setter cases. Medium for the expanded derive boundary. This does
not close the whole derive, numeric helper callbacks, aliased book/Spell/mover
storage, later equipment/Effect writes or ordinary row-index population. No
full-helper-return or first-SAVE vector is established.

### SAV-936

- In the successful new branch of `R0823`, the retained Human local passes
  through quest-item grant and optional mode Weapon dispatch. It then receives
  `identifier+04` and `owner+14`, enters the owner's actor list and a new
  group, conditionally calls a helper on its `+10` pointee, is assigned to
  `owner+34` and reaches `R0059`.
- Human `virtual+3c` resolves through `Human+38` to the supplied Weapon's
  `+38=R0850`; that body calls derive and later attached-Effect dispatch
  `R0855`.
- The collection receiver is `u32[owner+20]`; append stores the actor pointer
  in `node+8`.
- Group linking stores `u32[actor+70]=group` and copies the loaded owner
  pointer as `u32[group+44]=u32[actor+14]`.
- The caller at `L03850` first works from an existing `owner+34` actor and
  reaches the builder after other helpers; the builder itself retests
  `owner+34`. This reference alone therefore establishes no additional
  ordinary allocation route.
- Evidence: constructor/vtable witnesses, insertion ordering, receiver
  resolutions and six-span frontier. Supporting claims: `SAV-HUMALLOC-504`,
  `SAV-HUMNEW-505`, `SAV-HUMNEWSAVE-507`, `SAV-900`.

**Confidence.** Medium for the bounded 48-body, two-window receiver/lifetime
frontier. The 530 synthetic p-code cases contain no allocation, complete
creation, original process or emitted SAV.

**Unknown.** Allocator-return bytes, complete old-item/Effect receiver
populations, aliased collection storage, callbacks/reentry, same-actor
preservation through all intervening writes, the first genuine mission SAVE
values and semantic safety.

## Source callback and consequence recipients

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-958 | The selected source callbacks have concrete class-dependent targets in the Unit, Humanoid and Human tables. | High / Unknown | ✔ promoted | [EXP-0342](../experiments/EXP-0342-effect-source-consequence/) |
| SAV-959 | Humanoid/Human `source+48` has a Player mutation and notification boundary before death credit is reread. | High / Unknown | ✔ promoted | [EXP-0342](../experiments/EXP-0342-effect-source-consequence/) |
| SAV-960 | The admitted death progress callback and later Player/Diary accounting can have different eligibility outcomes. | High / Medium | ✔ promoted | [EXP-0342](../experiments/EXP-0342-effect-source-consequence/) |
| SAV-961 | The selected periodic path reaches the Effect source's concrete progress receiver independently of victim credit. | High / Medium | ✔ promoted | [EXP-0342](../experiments/EXP-0342-effect-source-consequence/) |
| SAV-962 | The manager has distinct source-owner and attached-Effect clears, without closing global source lifetime. | High / Unknown | ✔ promoted | [EXP-0342](../experiments/EXP-0342-effect-source-consequence/) |

### SAV-958

- ROM1 default constructor stores at `L00494`/`L00505`/`L08086` install
  the Unit/Humanoid/Human tables.
- Unit slots `+48`/`+60`/`+64` select empty `L08087`/`R1554`/`R0869`;
  they do not call its separate nonempty `+5c`.
- Humanoid and Human select `L08088`/`R1054`/`R0870`, with both latter
  methods dispatching the same source receiver's `+5c` to `R0810`.
- A type word or saved-address map hit alone does not prove one of these
  classes or allocation validity.
- Evidence: constructor/table and instruction controls; source callback
  vectors.

**Confidence.** High for constructor/table provenance, complete target bodies,
independently read table dwords and actual conditional dispatch.

**Unknown.** Ordinary class/lifetime after LOAD.

### SAV-959

- Complete `L08088` bypasses the owner helper for `server+0c` zero, a nonzero
  `victim+30` predicate or negative source HP.
- Otherwise it sends `victim+1c` and literal 1 to `R0449` on `source+14`.
  The helper adds to that `Player+38` and sends its new value and identity
  through `R0623` to `R0217`; opcode 103 uses the supplied Player
  `identifier+04`.
- The manager `R0427` then freshly reads `victim+48` signed and `victim+40`
  for `source+60`, without a renewed null/type/health guard.
- No whole-notification preservation follows: `R0217` has unresolved packet
  virtuals and delivery/refresh callees.
- Evidence: `source+48`, owner/packet/transport bodies, fresh-read controls.

**Confidence.** High for the reached own stores, class predicates and
fresh-read order; 24 isolated original-p-code callback cases discriminate the
branches. Three hypothetical transport-return mutations demonstrate the
dependency but do not claim original mutation or native null dereference.

**Unknown.** Complete notification effects.

### SAV-960

- In `R0427`, `source+60` receives the then-current `victim+40` source;
  concrete Humanoid/Human progress dispatch keeps that actor as the `+5c`
  receiver.
- A missing victim owner or the paired multiplayer/`victim-owner+5c` gate can
  make `+60` return without progress.
- The manager nevertheless continues: it rereads `victim+40` and `source+14`,
  increments that Player's `+4c` for victim type 33..63 or `+48` otherwise,
  then calls `R1484` on that Player's `+40` Diary with the victim.
- Distinct A/B sources and Player Diaries discriminate the actual recipient
  from the victim and the separate actor-owned Diary; `Effect+44` is zero in
  every death-prefix case.
- Evidence: death-consequence vectors and recipient events. Supporting claims:
  `SAV-843`, `SAV-908`, `MAGIC-ITEMKILL-117`.

**Confidence.** High for the concrete receiver reads and separate
gates/stores. Medium for the 14 joined prefix cases: earlier
tick/removal/Effect callbacks are omitted, and numeric award results/threshold
and notification returns are supplied. The cases execute actual progress and
Diary stores but establish no native award amount, first-after-LOAD chronology
or transitive preservation.

### SAV-961

- Token 8 application `R0565` dispatches its admitted `source+64`.
- For Humanoid/Human, `R0870` retains that same source as receiver of
  `R0810` through `+5c`; the victim is an argument.
- Eight conditional cases with distinct `Effect+44` and `victim+40` execute the
  source's actual progress stores while leaving the separate Diary state
  untouched.
- Null/negative source controls, Unit's empty callback and the reached
  victim-owner gates bypass those stores.
- Evidence: periodic-to-progress vectors and actual callback events.
  Supporting claims: `SAV-909`, `MAGIC-AREASOURCE-161`.

**Confidence.** High for the concrete dispatch/receiver and own progress
stores. Medium for the conditioned eight-case join: numerical calculations and
progress-notification returns are explicit cuts, source allocation/first
post-LOAD history is supplied, and the case stops before payload notification
`R0561`. No new amount or full source-lifetime claim is made.

### SAV-962

- Before the current actor tick, complete `R0427` clears a nonzero
  `actor+40` when the referenced `source+14` is null; three no-cut cases
  execute this check. The source is dereferenced, not validated as an
  allocation.
- Selected removal/membership helpers then precede the removed victim's
  attached-list loop. Each reached element receives `counter+42=1` and
  `source+44=0` before `virtual+38`.
- The subsequent `source+48`/`+60` pair is inside this manager, after its Unit
  tick dispatch, not inside `R0037` itself.
- The finite population: 36 complete bodies, 2345 independently decoded
  instructions, 176 calls including 27 computed calls, 429 additional
  target/width/vtable checks and 52 p-code cases.
- Evidence: complete caller, call frontier, summary and verification.

**Confidence.** High for the local list-element stores, the actual containing
function and the enumerated population. No original process or save corpus was
used.

**Unknown.** Nested removal/list callbacks, other derived Effect targets, every
Effect referring to a removed source, actual destruction, complete
reference-registration order/`Unit+68` second lookup, source
notification/progress transitive effects and native LOAD-to-first-consequence.

## Tavern roster progress

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-926 | An empty eligible shelf does not skip command 37 stock construction. | High / Unknown | ✔ promoted | [EXP-0339](../experiments/EXP-0339-tavern-roster-progress/) |
| SAV-927 | A valid empty owner group differs from a null group, and stock construction also retains unchecked allocation and identifier paths. | High / Unknown | ✔ promoted | [EXP-0339](../experiments/EXP-0339-tavern-roster-progress/) |
| SAV-928 | The client collector can finish with no eligible entries; emptiness does not validate its document map or type-index arrays. | High | ✔ promoted | [EXP-0339](../experiments/EXP-0339-tavern-roster-progress/) |
| SAV-929 | Empty mercenary collection writes selection -1, which the caption and price-refresh upper-bound checks admit. | High / Medium / Unknown | ✔ promoted | [EXP-0339](../experiments/EXP-0339-tavern-roster-progress/) |
| SAV-930 | Initial party selection is independent of the mercenary shelf and requires a marked live party entry. | High / Unknown | ✔ promoted | [EXP-0339](../experiments/EXP-0339-tavern-roster-progress/) |
| SAV-931 | Ten decoded archive observations support a bounded association between the saved eligible list and the reported tavern outcome, not a culprit. | Medium / Unknown | ✔ promoted | [EXP-0339](../experiments/EXP-0339-tavern-roster-progress/) |
| SAV-932 | A complete original tavern click and the first failure in the reported N3 run remain unlocalized by the bounded roster probe. | Unknown | ✔ promoted | [EXP-0339](../experiments/EXP-0339-tavern-roster-progress/) |

### SAV-926

- Complete `R1555..R1142` refunds, walks the Player group, sizes stock to
  16 and clears indices 0..15; fixed loops then construct Unit types 1..2 and
  Human types 3..15.
- No shelf or permanent-unlock array is read in this body.
- The publication loop starts at 1 and rereads stock size.
- With successful named constructors, a valid finite old-group chain and
  callbacks preserving stock size, the empty-old-group and mixed-old-group
  conditional vectors both return with types 1..15 published. The mixed input
  removes its two nonzero-type actors and skips its type-zero actor.
- Evidence: server stock body and stock vectors. Supporting claims:
  `MERC-CMD-007`, `SAV-918`.

**Confidence.** High for the explicit loop bounds and the absent eligible-list
access in the complete body. Completion is conditional on named refund,
allocation, construction, identifier, detach/destruction and publication
prerequisites.

**Unknown.** Constructor internals, real resource success and whole command
chronology.

### SAV-927

- The stock cleanup passes `Player+20+4`, or zero, to `R0028`; head helper
  `L08089` reads `collection+4` without a null guard.
- `R1334` caches `node.next` before returning `node+8`, so current removal
  can progress only while remaining nodes stay valid.
- `R0936` scans successive dwords from `L06793` until one is not all ones,
  without a first-scan bound, then finds and marks a zero bit through
  `L08090`. The four-full-word synthetic prefix asks for its fifth word; this
  is not a capacity measurement.
- Null allocation arms join type stores at `L08091`/`L08092` through
  address `0000014c`.
- Evidence: iterator/head/next/remove/id bodies, stock vectors.

**Confidence.** High for the local accesses, loop condition and conditional
controls. No saved value alone proves constructor/lifetime validity.

**Unknown.** Ordinary bitmap exhaustion, allocation exceptions, transitive
destructor effects, complete chain production and the owner failure's reach.

### SAV-928

- `R0746` clears output arrays and returns immediately for document-map
  count zero.
- With a nonzero map count it walks bucket/next links even when the supplied
  eligible array is empty. It admits CUnit objects whose type is in that array
  and whose unsigned working-pool word at type-1 is nonzero.
- Finite coherent buckets, acyclic stable links and returning services bound
  the walk. Empty-map, empty-eligible, one/two-entry, duplicate-eligible and
  zero-pool vectors return.
- A self-cycle with empty eligibility reaches the instruction cap; a nonzero
  map count with no bucket node attempts the null read at `L08093`.
- Matching types 0/16 access outside a supplied 15-word pool at `L07507`;
  membership is not a type-range check.
- Evidence: client collector, actual pointer resize and collection vectors.
  Supporting claim: `MERC-SHELF-002`, with its partial retraction.

**Confidence.** High for the complete local filter/walk and the conditional
boundary controls. Synthetic cycle/corruption controls do not establish
ordinary saved or runtime reach. Classifier, allocation, live map construction
and callback preservation remain explicit prerequisites.

### SAV-929

- The inn constructor clears its pointer arrays; resize `L08094` frees/nulls
  storage on size zero.
- Activation stores selection 0 for a nonzero mercenary count and -1 for zero
  at `L08095..L08096`. Its separate InnNPC append loop supplies no
  empty-mercenary fallback selection.
- Caption `R0600`, called directly by activation at `L08097`, checks
  signed selection <= count-1 before the indexed read at `L08098`.
- Refresh `R1556` checks signed selection < count before `R1557` reads
  the same array.
- Neither test excludes -1. Conditional zero-count/zero-storage vectors read
  `fffffffc`; one/two-entry controls return.
- Adding an NPC preserves the empty-list caption failure under the measured
  field-preservation condition.
- Evidence: activation/constructor/resize/caption/refresh/price bodies and
  activation/refresh vectors.

**Confidence.** High for local stores, resize semantics and negative-index
conditions. Medium for the activation-to-caption preservation inference: the
list slice and caption execute separately, with intervening UI/resource callees
left as boundaries. A synthetic intervening selection=0 write changes the
empty-with-NPC caption outcome to return.

**Unknown.** Native failure attribution.

### SAV-930

- `R1373` builds its primary list from CUnit values owned by
  `document+9b4`, with flag bit 1 at `actor+18c` plus byte `actor+15a` zero.
- `L08099` returns the first primary index with flag bit `0x20`, otherwise
  -1. Empty and nonempty-unmarked vectors both return -1.
- Activation stores it at `view+bc` and indexes party data at `L08100`
  without checking for the sentinel. A no-party activation vector reaches this
  indexed read even with one eligible mercenary.
- Valid application/campaign/document/player objects, coherent pointer arrays
  and preserved object lifetime are additional runtime prerequisites.
- Evidence: party bodies, activation and four party-selection controls.

**Confidence.** High for the complete selected collectors/selector and the
local unchecked use.

**Unknown.** Actual flag production, the original first selection, callback
effects, full UI entry and absence of other earlier failures.

### SAV-931

- All ten exact-file inputs pass the full research reader.
- Eight labelled resaves are main/selected mission 30/30; only T4/T7 have saved
  eligible set `[14]`, and those are the two owner-reported opening cases. The
  other six have empty eligibility and were reported closing.
- The Documents factor changes none of the four matched comparisons; all eight
  InnNPC/InnMission arrays are empty.
- N3 city has shelf `[14]`, empty permanent unlocks and InnNPC/InnMission
  `[22]`/`[30]`, hence empty saved eligibility; the accepted town comparator is
  mission 40/40 with saved eligibility `[14,6]`. Those latter two are different
  campaign states.
- Runtime outcomes are owner reports for the labelled inputs; native reload of
  the resaves was not newly witnessed.
- Evidence: hashed input table, save-inputs/summary/projections and owner
  metadata digests. Supporting claims: `SAV-CAMPAIGN-083`, `MERC-SHELF-002`.

**Confidence.** Medium for the bounded list/outcome association and the
labelled metadata join. Exact parsing and saved values do not prove stock
publication, live selection, native construction or instruction reach.

**Unknown.** The observed hang/closing cause.

### SAV-932

- The selected population: 31 complete routines, zero extra windows, 32
  distinct conditional vectors and ten exact-file observations.
- Stock loops, client list collection, initial party selection and later
  caption/price indexing have separate prerequisites.
- A reached negative mercenary index, an earlier/later failing unresolved
  operation and a callback changing selection fit the available observations.
- The report supplies a reachable T0..T7 visible-pattern prediction and an
  instruction-level falsifier for the proposed negative-index attribution.
- No original process, native instruction witness or new owner repeat run was
  used.
- Evidence: method, alternatives, predictions and explicit service boundaries.

**Confidence.** Unknown for full click chronology and owner failure
attribution. The local High claims and the preserved observation agreement do
not raise that confidence.

**Unknown.** Whether the failing original click first reaches the measured
consumer with the relevant fields intact.

## Tavern command Player prerequisites

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-918 | The reached tavern command 37 arm resolves `Player+04` before its matched stock handler changes `Player+38`; absolute click-to-consumer chronology remains Unknown. | High / Medium / Unknown | ✔ promoted | [EXP-0338](../experiments/EXP-0338-tavern-player-prerequisites/) |
| SAV-919 | The tavern lookup key `Player+04` has distinct constructor, registration and archive LOAD producers; the measured LOAD list path does not directly re-register it. | High / Unknown | ✔ promoted | [EXP-0338](../experiments/EXP-0338-tavern-player-prerequisites/) |
| SAV-920 | Four named original resaves close under the exact reader but do not localize the N3 tavern failure. | Medium / Unknown | ✔ promoted | [EXP-0338](../experiments/EXP-0338-tavern-player-prerequisites/) |

### SAV-918

- Complete campaign helper `R1318` calls sender `R0645` before computed
  view activation. The sender stores opcode 37 and copies active client
  `CPlayer+04` into `command+5`.
- In the measured category-below-1 opcode switch, the command 37 table target
  is `L08101`.
- That arm calls `L00171`, whose predicate compares signed `Player+04` with
  the signed input word at `L08102..L08103`; it returns the first matching
  pointer or zero. The arm exits on zero; a match calls `R1555` on static
  tavern `L08104`.
- Its first call is `R0449(Player,Tavern+9c,0)`: a wrapped 32-bit addition
  to `Player+38`, followed by event 67 even for a zero refund. Only after that
  event call returns does the stock body inspect `Player+20` and its previous
  roster.
- The direct lookup and refund bodies do not first consume `+2c`/`+44`/`+48`.
- Eight conditional lookup vectors discriminate `+04` from
  `+08`/`+2c`/`+44`/`+48`; four refund vectors check the add and event
  arguments.
- Evidence: `functions.tsv`, dispatcher windows, `evidence/tables.tsv`,
  `evidence/vectors.json`. Supporting claims: `MERC-CMD-007`,
  `AI-FORMACTIVE-317`.

**Confidence.** High for these direct instructions and local ordering under the
stated synthetic collection/event services. Medium for the bounded entry
account.

**Unknown.** Earlier UI/transport calls, the actual command category/key, real
collection-helper effects, later publication and the owner's failure
instruction; the named offsets are not excluded from those unexpanded paths.

### SAV-919

- Player constructor `R0201` explicitly writes `word+04=0` and
  `dword+38=0`. The Player descriptor's factory `R0203` calls that
  constructor after successful allocation.
- Non-LOAD roster copy `R0423` calls registration `R0424` at `L08105`,
  before its `colour+44` write. Registration writes the chosen slot at
  `L07240`, with the map-owning zero-slot counter fallback at `L08106`, and
  derives or clears `+2c` as in `SAV-664`.
- Player LOAD instead passes its `+04` address to `L08107` at `L08108`; its
  `R0282` buffered arm copies exactly the next two bytes, without
  normalization, and advances by two. Three conditional buffered-reader vectors
  preserve arbitrary 16-bit values and adjacent destination bytes.
- Manager/list LOAD `R1372 -> R1558` reads a typed Player through
  `L08109`, then passes its returned pointer to append `R0417`; those
  complete LOAD bodies contain no direct call to `R0424`.
- The matched tavern refund subsequently changes the restored balance; `+38`
  restoration and XOR are `SAV-OBF-029`.
- Evidence: constructor/factory/registration/roster-copy and LOAD bodies,
  `evidence/tables.tsv`, buffered archive vectors. Supporting claims:
  `SAV-664`, `SAV-OBF-029`, `SAV-PLAYERIDENT-830`.

**Confidence.** High for explicit stores, direct call ordering and the local
two-byte copy. The no-registration negative is limited to those LOAD bodies,
not their full transitive callees.

**Unknown.** General archive dispatch, allocator/collection callbacks, all
later slot writes, active client identity, first-save values and post-LOAD
mutation.

### SAV-920

- N3 `game0067`/`game0068`/`game0069` and the accepted `game0066` have four
  distinct SHA-256 identities listed in the experiment, with 3/3/1/1 Player
  records and present/present/absent/absent world halves.
- Their first Player has `+04=1`, `+08=1` and `+28=0` in all four; `+44` is
  2/2/0/2.
- The remaining two Players in each world file have `+04`/`+08` of 2/2 and
  3/3, `+28=1`, and `+44` of 3 and 4. Both city files have one Player with
  saved slot 1.
- This does not observe `command+5`, active client identity, manager validity
  or instruction reach. The colour difference is a measured value difference,
  not a failure cause.
- The owner reported N3 shop/school success followed by tavern hang and process
  exit without an error dialog; no instruction or exception address was
  observed. The accepted comparator label supplies no new tavern witness.
- Evidence: `inputs.tsv`, `evidence/save-inputs.tsv`,
  `evidence/save-players.tsv`; an exact `tools/savfull -player-identity` read
  of each named file, with zero refusals.

**Confidence.** Medium for the bounded four-file observation and comparison.
No game process was launched or owner repeat requested.

**Unknown.** Different campaign histories, whole Player equivalence, native
construction, later scalar use, the first runtime consumer and the failure
cause.

## Embedded application-state ownership

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-914 | The two located application-state readers use temporary registries and named consumers. | High / Unknown | ✔ promoted | [EXP-0337](../experiments/EXP-0337-application-store/) |
| SAV-915 | The located ordinary SAVE reconstructs its application registry; a direct raw registry round trip is a different persistence boundary. | High / Medium / Unknown | ✔ promoted | [EXP-0337](../experiments/EXP-0337-application-store/) |
| SAV-916 | Raw parsing and successful typed lookup do not establish application acceptance: both readers unconditionally consume four Shortcuts words. | High / Unknown | ✔ promoted | [EXP-0337](../experiments/EXP-0337-application-store/) |

### SAV-914

- Complete `R0098` constructs a registry at `frame+3c` (`L08110`), loads
  it (`L08111`), reads campaign state on the same stream (`L08112`) and
  destroys the registry (`L08113`).
- Complete `R0099` uses `frame+24` at `L08114`, `L08115`, `L08116` and
  `L08117`.
- Constructor `R1450` initializes empty root/record/pool state; destructor
  `L08118` passes record and pool allocations to free.
- Both application bodies consume named Character, CurrentState, GameOptions,
  View, SpellBook and Objects state; the long body additionally reads
  Inventory, Projectiles/`Prj<id>` and Fog. Dynamic Objects/`Group<n>` lookups
  are bounded by ten.
- Complete caller `R1284` routes to the two readers after its state probe;
  the selected probe window itself reads CurrentState/InBattle with default 1.
- Neither complete application body enumerates all registry names.
- Seven conditional executions of `L08119..L08120` transfer the same eight
  GameOptions/View scalar values with and without unrelated roots/leaves.
- Evidence: complete application readers/route and registry lifetime bodies;
  load-state-probe caller window; V096..V103. Supporting claims:
  `SAV-CAMPTAIL-070`, `REG-099..REG-101`.

**Confidence.** High for direct ownership, named-consumer and original
scalar-prefix operations in the matching EN/RU images. The prefix starts after
the file/campaign/string joins; it is not a full LOAD run.

**Unknown.** Unexpanded helpers, exception unwinding, indirect aliases and
subsequent application callbacks.

### SAV-915

- Complete `R0084` constructs a fresh registry at `frame+18` (`L08121`)
  and calls typed setters with fixed names and values from current
  application/UI/world fields. It additionally emits nonempty
  Objects/Group0..Group9 lists and `Prj<id>` sections for live projectile IDs.
  The world guard controls Projectiles and Fog.
- At `L08122` the newly built registry reaches `R1390`; `L06699` writes
  campaign state to the same file, and `L08123` destroys the new registry.
- No loaded-registry pointer or arbitrary input-name enumeration is passed
  along this direct ownership path.
- A conditional cleanup/construction-prefix vector frees the loaded
  allocations and executes the original constructor plus the first SAVE
  setter: the resulting two-record CurrentState/InBattle store contains none of
  the previous unrelated entries.
- The remaining typed-setter programme is static evidence, not a fully
  interpreted or native resave.
- This narrows the fixed 28-record/19-leaf statement of `SAV-TAILEXT-062` to
  its measured corpus, without changing those counts.
- Evidence: application SAVE, writer/constructor/destructor and two direct
  caller windows; V103. Supporting claims: `REG-102`, `SAV-PROJSTORE-428`,
  `SAV-PROJLOAD-429`, `SAV-TAILEXT-062`.

**Confidence.** High for the complete direct SAVE flow, fresh construction and
named producer sources. Medium for loss of noncolliding entries that no
producer requests at an ordinary completed SAVE.

**Unknown.** Hidden effects of unexpanded helpers and a complete native
LOAD/SAVE cycle.

### SAV-916

- After `R1061(SpellBook,Shortcuts,destination)`, the short body reads four
  dwords at `L08124`; the long body reads four at `L08125`. Neither checks
  the getter return or the resulting count before those copies.
- Original constructor `L05586` initializes an empty destination. The getter
  returns 0 without changing it for a missing name; an empty kind-6 array
  returns 1 after resizing to zero.
- With the declared synthetic resize service, both reach the long consumer's
  unmapped-null read. One- and three-element arrays expose reads beyond their
  declared count; four and five cover all four reads.
- The kind-2 compatibility arm produces one dword from the leaf's low byte; a
  kind-4 control stops at the original exception-throw boundary.
- Evidence: both application readers, `registry_get_dwords`,
  `dword_array_construct`; V104..V111.

**Confidence.** High for the paired owning loops, the complete
getter/constructor and eight discriminating conditional vectors. The resize
service supplies array storage; it does not emulate the OS allocator.

**Unknown.** Actual out-of-bounds values, native crash/exception behavior,
full LOAD and later callbacks.

## Attached Effect source and Unit attribution

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-906 | Effect creation and attachment have distinct source writes. | High / Unknown | ✔ promoted | [EXP-0336](../experiments/EXP-0336-effect-attribution/) |
| SAV-907 | The ordinary new-Effect archive route constructs a zero source slot, and the Effect's own serializer does not restore that slot. | High / Medium / Unknown | ✔ promoted | [EXP-0336](../experiments/EXP-0336-effect-attribution/) |
| SAV-908 | Unit kill credit has a concrete post-load identity repair despite its raw u32 wire representation. | High | ✔ promoted | [EXP-0336](../experiments/EXP-0336-effect-attribution/) |
| SAV-909 | The periodic source and the later credited actor are different conditional consumers; zero HP is not negative HP. | High / Medium / Unknown | ✔ promoted | [EXP-0336](../experiments/EXP-0336-effect-attribution/) |
| SAV-910 | The bounded static experiment does not establish a native source-actor lifetime across SAV. | High / Unknown | ✔ promoted | [EXP-0336](../experiments/EXP-0336-effect-attribution/) |

### SAV-906

- Default constructor `R1049` clears `Effect+44` after its Token constructor
  call.
- Template constructor `R1047` clears it only after successful parsing and
  normal return from the temporary object's destruction; its null-parser arm
  does not write `+44`.
- Copy `R1001` copies the source dword.
- The selected area-builder window `L08126..L08127` admits the caster when
  `caster+3c` is nonzero, stores it in `envelope+3c`, then copies it to
  `inner+44` if the inner is present.
- The continuous same-id attachment branch changes only the retained counter.
  The non-continuous replacement's own stores between unapply/reapply change
  magnitude/counter and leave source unchanged; that statement does not assume
  the callbacks preserve source.
- Evidence: constructors, copy, attachment and envelope vectors. Supporting
  claims: `ITEM-EFFOBJ-072`, `MAGIC-POISONREFRESH-158`.

**Confidence.** High for the selected original stores, branches and
conditional p-code controls. Prior attachment claims remain scoped to their
measured branches.

**Unknown.** The null-parser arm's allocation contents, ordinary parser-failure
reach, omitted base/copy/destructor effects and complete spell-builder paths.

### SAV-907

- Unit `R0210` sends `Unit+20` to `R0951`, whose LOAD calls typed reader
  `L08128` with descriptor `L08129`.
- `L05928`'s new-object branch invokes runtime-class creation before virtual
  serialization; the descriptor's factory `R1559` allocates 72 bytes and
  calls default `R1049`.
- The 44-byte Effect programme restores Token, kind, mode, packed operand and
  id, omitting `+44`.
- An archive back-reference instead returns the existing object without
  rerunning its factory/serializer. It thus preserves that loaded object's
  alias, not a separately reconstructed source actor.
- Effect's own post-load slot `+24` is Token's position-rebind `R1560`; the
  selected Unit/Human lifecycle bodies do not walk `Unit+20` or reconstruct
  `Effect+44`.
- Evidence: descriptors/vtables, complete readers,
  factory/serializer/back-reference vectors. Supporting claims:
  `SAV-EFFCHAIN-046`, `ITEM-EFFSAVE-077`.

**Confidence.** High for the concrete class/factory/serializer/alias routes and
their own writes; scalar/archive/base calls in the synthetic cases are explicit
cuts. Medium for the joined source-reset conclusion through the selected
lifecycle, because indirect callers and later intervening writes are not a
native LOAD-to-first-payload witness.

**Unknown.** Actual runtime state after LOAD.

### SAV-908

- After raw restore of `Unit+40`, Unit lifecycle `R1482` passes its address
  to `R1108` at `L08130`.
- The helper looks the saved word up in the map at `[L00285]+88`; a hit
  replaces it with the mapped pointer, a miss writes zero. `Unit+48` is not
  changed there.
- The same lifecycle repairs raw `+64` through `L04868` and `+68` through
  `L04869`, both with the same hit/miss rule.
- Token LOAD registers old-object-address to new-object-address in that map.
- The document's world branch dispatches actor lifecycle; its no-world branch
  instead explicitly clears `actor+40`.
- Unit credit can therefore retain an actor identity relation independently of
  the omitted `Effect+44`.
- This withdraws the previous inferences that only `+68` is remapped and that
  portable credit requires new state; the evidence carries explicit amendments
  to `MAGIC-ITEMKILL-117` and `ITEM-CASTSTATE-056`.
- Evidence: Unit post-load and map-cut vectors. Supporting claim:
  `SAV-PTRMAP-035`.

**Confidence.** High for the concrete caller, fields, map selection and
hit/miss stores, independently decoded and exercised with different mapped
actors and absent keys. No class/lifetime validation by the map, complete
arbitrary-reference registration order, raw-pointer validity after later
callbacks or native resume is established.

### SAV-909

- A nonzero Token 8 HP payload writes victim HP before testing `Effect+44`.
- A null source skips source dispatch; helper `R1561` tests signed source
  HP<0, clearing `Effect+44` for negative HP and admitting HP 0 or positive HP
  to source `virtual+64`.
- Zero computed damage bypasses this source check. Duration-only controls
  decrement their counter without a periodic source call.
- Separately, the selected death prefix consumes repaired `Unit+40`, applies the
  same negative-HP predicate, and reaches `source+48` before `source+60`.
- The later `+60` argument window reads `Unit+40` and signed `tag+48` again, so
  the intervening `+48` callback is a lifetime/state boundary.
- A null Effect source can coexist with credit mapped to either A or B.
- Evidence: periodic, duration, map-to-death and callback-boundary vectors.
  Supporting claims: `MAGIC-AREASOURCE-161`, `MAGIC-ITEMKILL-117`.

**Confidence.** High for the reached local instructions and the 59-case
population's discriminators, including two actors and negative/zero/positive
health. Medium for joining the LOAD suffix to a later admitted death prefix:
the map result and chronology are supplied, and the preceding
owner/type/scheduler gates and actual award callbacks are unexecuted.

**Unknown.** Destroyed-source pointer validity and eventual awards.

### SAV-910

- Population: 36 complete bodies, one refused candidate body and four windows:
  4,621 body instructions plus 109 window instructions, 444 body calls
  including 40 computed calls, and 59 synthetic cases, over one executable
  population shared by EN/RU. Every retained instruction has a second-decoder
  check.
- Candidate `R1483` is inside the six-byte instruction at `L07736`, not an
  entry. The former named function and asserted caller chain in `SAV-776` are
  withdrawn, without replacing them with an unmeasured whole-builder graph.
- The selected source windows are explicitly partial.
- No save corpus, original process, OS allocation, source destruction or
  ordinary LOAD-to-next-consumer chronology was observed.
- Evidence: functions/windows, decoder audit, call frontier and verification.

**Confidence.** High for the enumerated population and the interior-address
discriminator. A later repair can fit these local bytes; no global absence or
safe-default claim follows.

**Unknown.** The unexecuted original lifecycle, transitive call preservation,
parser-failure reach, all source producers and a complete remaining repair
census.

## Residual Human creation-to-first-SAVE writes

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-898 | The mode-gated Human helper does not itself overwrite the eleven selected residual bytes before derive. | High / Unknown | ✔ promoted | [EXP-0335](../experiments/EXP-0335-human-first-save/) |
| SAV-899 | The conditional starting-skill store can address every selected residual span, but ordinary production of those indices is not established. | High / Unknown | ✔ promoted | [EXP-0335](../experiments/EXP-0335-human-first-save/) |
| SAV-900 | The expanded eleven-byte Human frontier does not choose an allocation-to-first-mission-SAVE model. | Medium / Unknown | ✔ promoted | [EXP-0335](../experiments/EXP-0335-human-first-save/) |

### SAV-898

- In live hero setup, nonzero `server+12c` or `server+134` reaches `R1562`;
  both zero bypass it.
- Its complete 33-instruction body writes six words at `actor+102..10c` and six
  bytes at `+10e..113`, all 100, then calls `virtual+50` at `L08131`. The
  supplied Human vtable resolves that call to `R0280`.
- Sixteen original-instruction cases cross fills 90/165 and three zero/one
  server flags. All actor bytes match only the named speed/defence changes, and
  `bc/bd`, `da/db`, `e8/e9`, `f6`, `f7/f8`, `fc/fd` remain unchanged at the
  dispatch cut.
- Evidence: complete bodies, mode-gate vectors and Human vtable.

**Confidence.** High for this local gate and own-write result. The caller's
prior name copy/construction and subsequent derive/insertion are not executed.
It is not a whole-helper-return or first-SAVE preservation claim.

**Unknown.** Ordinary configuration values and transitive writes.

### SAV-899

- The packet producer copies `draft+4c0`'s low byte to `packet+0e`; the
  selected dispatcher window forwards it to the hero builder, whose two local
  arms forward it with level 10/20.
- After an earlier removal boundary, `L08132` clears slots 1..5 and
  `L03936` writes `word[actor+a8+2*(index&255)]` from the level byte.
- Over all 256 byte indices, exactly 10/25/32/39/40/42 intersect the selected
  eleven bytes: `bc/bd`, `da/db`, `e8/e9`, `f6/f7`, `f8` and `fc/fd`
  respectively.
- Thirty-two original store-window cases and ten dispatcher cases confirm the
  addresses plus byte transport.
- Evidence: selected-aliases, skill-write and packet-skill vectors.

**Confidence.** High for the conditional forwarding and indexed-write
relations. The starting-skill index is distinct from `actor+b6`. No ordinary
producer, allocator provenance or consumer reach is inferred from injected
state.

**Unknown.** Earlier validation, ordinary out-of-range input reach, earlier
removal effects, subsequent calls and an actual creation-to-store execution.

### SAV-900

- Population: 24 full analyzed bodies, 5074 body instructions and 298
  instructions in two dispatcher windows, 447 body calls including 35
  computed, and ten reference rows.
- Eight archive controls execute the original mode getter and reach raw
  Write/Read with `actor+a6`/24 bytes or `actor+d4`/64 bytes, stopping before
  the raw call.
- Earlier embedded initialization preserves `bc/bd` and clears the nine
  selected modifier bytes; the mode-gate and indexed-store results do not close
  the intervening lifetime.
- Zero-return, nonzero-return and later-write models remain compatible for the
  live tail; a later producer or data dependence remains compatible after
  modifier clearing.
- Evidence: call-frontier, archive-cut vectors and six-span frontier.
  Supporting claims: `SAV-HUMGAPS-449`, `SAV-HUMNEW-505`,
  `SAV-HUMNEWSAVE-507`.

**Confidence.** Medium for this enumerated static frontier. The 66 synthetic
original-instruction cases include no allocation, full Human creation, emitted
SAV or original process. No safe default vector follows.

**Unknown.** Actual allocation survival, ordinary alias reach, full
callback/lifetime writes, first genuine SAVE values, first ordinary computed
use and semantic safety.

## Terminal campaign and the city persistence boundary

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-890 | The selected final-campaign dispatch bypasses ordinary advancement before requesting a separate UI handler. | High / Unknown | ✔ promoted | [EXP-0334](../experiments/EXP-0334-terminal-campaign/) |
| SAV-891 | An upper-bound loader refusal does not make the outer main-load operation transactional. | High / Unknown | ✔ promoted | [EXP-0334](../experiments/EXP-0334-terminal-campaign/) |
| SAV-892 | The selected SAVE/LOAD programmes contain no separate terminal campaign wire arm. | High | ✔ promoted | [EXP-0334](../experiments/EXP-0334-terminal-campaign/) |
| SAV-893 | The declared preserved corpus supplies no terminal-city witness: every one of its 40 no-world paths has LastMission 0. | Medium / Unknown | ✔ promoted | [EXP-0334](../experiments/EXP-0334-terminal-campaign/) |

### SAV-890

- In admitted message 41d of `R0701`, campaign is `frame+548`.
- After the direct store to `campaign+124`, a nonzero `campaign+11c` and a
  selected `campaign+118` divisible by 10 select `L03704` -> `R0802`,
  server `R1301(1)`, message 428 and the common dispatcher continuation.
- The alternate branch `L08133` contains the ordinary reward/advance calls,
  including `R1414`.
- The measured switch maps 428 to `L08134` -> `L08084`; that body calls UI
  virtuals through `frame+11c` and `frame+cc`, sets `frame+41c=0` and
  `frame+3dc` bit 40, and conditionally requests `music/inn_ssi.wav`.
- `Frame+11c` in that callee is a UI pointer, distinct from `campaign+11c`.
- Evidence: windows W03/W11/W15/W19/W21/W23/W24/W26. Supporting claims:
  `REG-SCN-063`, `FAME-PRODUCER-014`.

**Confidence.** High for these local predicates, ordering, switch and receiver
distinctions in the matched EN/RU image. The condition is stored LastMission
plus selected-main routing, not merely an absent successor.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Unknown.** Earlier and later callbacks, successful message delivery, a
displayed final city, a working city SAVE and the resulting terminal document.

### SAV-891

- `R1297` reads ScenarioMissionCount with default-1 and returns 0 for a
  signed request greater than its 32-bit count-times-10 limit, before its
  direct campaign stores or document-grant loops.
- `R0755`'s higher-request arm does not test that return: it writes the
  request to `campaign+04` and `+118` and returns 1.
- The selected campaign `vtable+0c` getter returns `+04`, so `R1563`
  requests `current+10`.
- The outer operation changes those two numbers without executing the inner
  registry-replacement/document-grant block. Retention of the old fields and
  documents is conditional on intervening helper side effects.
- Evidence: windows W02/W12/W13/W14/W29/W30/W31.

**Confidence.** High for the named static guard, getter and unconditional outer
stores. This conditional helper result does not establish a naturally produced
terminal scalar or a reachable post-completion city; `SAV-890` can bypass the
helper.

**Unknown.** Arbitrary callback side effects, overflowed configurations and
runtime arrival state.

### SAV-892

- Mode 2 frontend SAVE `R0084` calls document writer `R0804` and later
  campaign writer `R1294`. The latter writes current base/child fields,
  counted collections and document pairs, then the seven scalars including
  `selected+118` and `LastMission+11c`.
- Full reader `R0434` restores that grammar and those scalars without a
  local terminal-value rejection.
- Its position suffix uses the first MapPoint when `+120` is nonzero; otherwise
  `R1324` resolves the selected mission through GlobalMap/MissionObjects,
  with lookup default 1 and subtraction 1.
- Frontend `R1284`'s no-world return calls `R0098`, whose `L08112`
  restores the campaign before its remaining UI/application work and message
  42e.
- Evidence: windows W05/W08/W17/W18/W25/W32. Supporting claims:
  `SAV-CAMPAIGN-076/085`, `SAV-WHEADLOAD-521`.

**Confidence.** High for the selected archive, branch and raw scalar transfer
relations. First-point state, missing-key lookup, allocation/file callbacks and
later UI work remain distinct from original acceptance. No terminal input was
loaded by the original, and no terminal SAVE output or accepted numerical
terminal representation is established.

### SAV-893

- A read-only recursive inventory of EN, RU and preserved saves contains 129
  paths/76 distinct digests.
- The document walker closes 124 paths: 84 world and 40 no-world; four document
  refusals and one header refusal remain listed. The campaign grammar closes
  128 paths.
- The 40 no-world paths comprise 32 digests, with current=selected 30 on 39
  paths and 40 on one; every one has LastMission=0.
- The two registered `globalmap.reg` entries match, with 28 MissionObjects
  mappings, but this inventory establishes neither completion lineage nor
  original runtime acceptance.
- Evidence: `evidence/corpus-inventory.json`, `corpus-summary.json`,
  `globalmap-numeric.json`.

**Confidence.** Medium for the bounded corpus result. No-world means a
parser-classified shape, not a witnessed usable city. No absence outside these
roots or beyond these captured digests is claimed.

**Unknown.** Post-terminal fields, city reachability and subsequent native
acceptance.

## Mover byte transport and later producer boundary

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-MOVRATE-866 | The Unit archive prefix transports the exact allocated mover byte; a later derive is a different producer. | High / Unknown | ✔ promoted | [EXP-0329](../experiments/EXP-0329-mover-rate-binding/) |

### SAV-MOVRATE-866

- At `L08135..L08136` actor is the Unit serializer receiver saved in
  the frame local at -20, archive is its argument, and `L08137` passes `[actor+154]` as `this` to
  the call to `R1348`.
- This helper passes the mover address and 180-byte length to `R0271` when
  archive mode bit 0 is clear, or `R0272` when set.
- Five synthetic buffered SAVE/LOAD pairs per image execute this Unit prefix,
  both original archive-buffer branches and the selected original aligned copy.
  All 180 bytes match, including byte `+a` independently varied over
  0/1/7/19/255. The archive cursor advances 180 in each direction.
- A composed subsequent `R0249` invocation consumes the restored byte and
  changes facing only.
- Complete selected Unit/Humanoid/Human serializer bodies contain no direct
  `mover+a` replacement after this transfer; their other callees and enclosing
  events remain separate boundaries.
- Evidence: archive-prefix and restored-byte vectors.

**Confidence.** High for the exact receiver, byte range and admitted buffered
transfer. The vectors use synthetic memory and invoke the consumer explicitly.

**Unknown.** File callbacks, whole-LOAD preservation, first live dispatch and
whether Human derive rewrites the byte beforehand; `MOVE-RATE-053` supplies
that later local producer.

## Diary word-array consumer frontier

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-854 | The reproduced direct word-accessor population establishes no independent Diary word consumer. | Medium / Unknown | ✔ promoted | [EXP-0328](../experiments/EXP-0328-diary-word-consumers/) |
| SAV-855 | The sole raw direct caller of `R1142` supplies a freshly constructed stack word collection. | High / Unknown | ✔ promoted | [EXP-0328](../experiments/EXP-0328-diary-word-consumers/) |

### SAV-854

- The byte-identical original EN/RU images contain 14 raw E8 references to
  `R0130`, plus its one call to `L08138`; there is no raw E9 or
  stored-address hit to either entry in the file-backed sections. All 15 sites
  decode at selected instruction boundaries.
- Three wrapper sites are the known Diary initialization/decrement. The other
  11 have explicit receiver routes or unresolved source identities, rather than
  an established `Player+40` or `actor+1e4` Diary chain. Four belong to the
  argument-taking routine covered by `SAV-855`.
- All five raw E8 callers of backing-pointer helper `L08139` receive locally
  constructed stack collections; transitive value provenance is unclosed.
- The selected 44 windows contain 7,638 unique instructions.
- Separate linear shape censuses retain 28 positive immediate/displacement
  `0x1e4` candidates, 320 candidates that add `+0x18` to the object pointer (by addition or address formation), 368 non-stack dword
  loads at `+0x1c`, and 555 two-byte memory operands with scale 2. Only named
  receiver paths are classified; those shape populations are not whole-program
  alias proofs.

**Confidence.** Medium for the bounded negative. No additional caller event,
actor-owned word use or post-LOAD consumer is established.

**Unknown.** Indirect/computed calls, rebased or inlined access, unclassified
receivers, intervening callbacks and first live chronology.

### SAV-855

- The four shared-accessor sites of `R1142` are `L08140`, `L08141`,
  `L08142` and `L08143`.
- The only raw E8 reference is `L08144`; no raw E9 or stored-address
  reference to that entry is found.
- In the local `L08145..L08146` branch, `L08147` takes the address of
  the frame local at -0xcc and `L08148` calls CWordArray construction.
- SetSize receives the input dword at `+0x0a`; the caller requests a copy of
  twice that count from `input+0x0e` into the collection's backing pointer,
  then passes the stack collection as argument 2 to `R1142`, with receiver
  `L08104`.
- This rejects reading this direct caller's collection object as an existing
  Player or actor Diary.

**Confidence.** High for the stated local construction and argument flow. The
copy primitive and allocator internals are outside this receiver proof. It
does not establish a Diary consumer or a post-LOAD path.

**Unknown.** The branch's incoming event, input payload provenance, indirect
callers and any semantic relationship between copied words and Diary words.

## Diary counters and consumers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-842 | The located Diary writer saturates its dword at 17 but continues decrementing its word toward zero. | High / Unknown | ✔ promoted | [EXP-0327](../experiments/EXP-0327-diary-consumers/) |
| SAV-843 | The located live Diary call belongs to attributed actor teardown, not the health subtraction itself. | High / Unknown | ✔ promoted | [EXP-0327](../experiments/EXP-0327-diary-consumers/) |
| SAV-844 | Diary capacity comes from Units, but its actor index has class-dependent table identity. | High / Unknown | ✔ promoted | [EXP-0327](../experiments/EXP-0327-diary-consumers/) |
| SAV-845 | `Diary+2c` routes a conditional notification whose consumer reads the supplied Player's own Diary. | High / Unknown | ✔ promoted | [EXP-0327](../experiments/EXP-0327-diary-consumers/) |
| SAV-846 | The reached actor-owned Diary lifecycle differs from the live Player Diary mutation path; no additional actor-owned array consumer was established. | Medium / Unknown | ✔ promoted | [EXP-0327](../experiments/EXP-0327-diary-consumers/) |

### SAV-842

- `R1484` rejects an actor only when `virtual+30` is nonzero and its
  unsigned `+0c` byte exceeds 63.
- It then calls `R1564`, which applies the same gate and decrements a
  nonzero unsigned word at `Diary+18`. Afterward it reads the dword at
  `Diary+04` and returns if its unsigned old value exceeds 16; otherwise it
  increments it.
- The selected array accessors perform pointer arithmetic without a size
  check.
- Original-instruction vectors from defaults give `(17,1007)` after 17
  admitted calls and `(17,1006)` after 18, so the corpus complement relation
  in `SAV-668` is not a permanent invariant.
- A zero word still permits a dword increment, and an already oversized dword
  is not normalized.

**Confidence.** High for the complete local bodies and 21 mutation/sequence
vectors, with notification transport cut.

**Unknown.** Loaded malformed-index acceptance and global writer completeness.

### SAV-843

- The sole E8 caller is `L08149` in `R0427`, under actor `T+54==16`.
- Its tail selects `S=T+40`, requires S type word `+0e` in 33..63 and signed
  `health+94>=0`, and passes T to `S->Player(+14)->Diary(+40)`.
- Negative S health clears `T+40` and skips; zero health passes. An earlier
  prefix clears S when `S+14` is null.
- Two virtual effects on S precede the Diary call.
- `R0037` has a death/animation/decay path to `+54=16`.
- Standard strike `R0246` passes the attacker into target `virtual+4c`; its
  resolved `R0265` body can store that source at `target+40` without a
  target-death gate.

**Confidence.** High for the named local branches, the receiver chain and 7
teardown-tail vectors with two explicit virtual-call cuts.

**Unknown.** Whether every kill reaches this path, whether S always names the
lethal source, and which event runs first after LOAD.

### SAV-844

- `R0184` writes the matching Units ordinal to `actor+0c` at `L08150`;
  `R0656` writes the matching Humans ordinal at `L08151`.
- `Vtable+30` returns 0 on Unit and 1 on Humanoid/Human, so the Diary gate
  excludes only upper Human-family indices, not all indices above 63.
- The independently reached packet consumer selects Units row i, then
  parameters 29/30, named typeID/face by both original database schemas.
- Row 68 is typeID 79/face 1; row 72 is 80/1 and row 73 is 80/2, separating
  row identity from presentation type.
- The original tables have 119 Units slots and 216 Humans slots.
- This narrows `SAV-667`'s sizing-to-single-index-domain inference.

**Confidence.** High for the original producer instructions, virtual slots, the
complete database grammar and index-bound vectors.

**Unknown.** The meaning of every unobserved lower index and the full live
population.

### SAV-845

- `R1484` passes a nonnull `+2c` to `L08064` when an increment changes the
  half-count: local increments to 2,4,6,8,10,12,14,16. A null `+2c` suppresses
  notification but not mutation.
- `L08064` reads the supplied `Player+40` dword array and emits opcode 186
  with 17 words.
- `Player+68>10` fills all words with `0xffff`. Otherwise it scans
  `i=64..dword_count-1` and uses `q=min(low-byte(dword[i]>>1),7)`, Units
  typeID 64..80 as output index `typeID-64`, and face to shift q by
  `4*(face-1)`, then ORs the low word.
- Two other E8 call sites are `L08063` and `L08065`.
- It does not read the Diary word array.

**Confidence.** High for the complete local builder, the null/different-owner
controls and 10 packet vectors including byte narrowing and OR collision.
Packet transport is cut at `R0217`.

**Unknown.** Receiver UI, names and live refresh cadence.

### SAV-846

- Player construction supplies itself to `R1565`.
- Human initialization conditionally constructs `R1566` and stores the
  null-owner Diary at `actor+1e4`.
- Humanoid member construction, destruction and serialization
  clear/delete/archive that pointer; LOAD reads a Diary object into it.
- The selected live mutation follows `S+14` then `Player+40`, never `S+1e4`.
- The bounded image census has 15 raw E8 candidates to `L08152`, 14 to
  `R0130` and 28 linearly decoded immediate/displacement occurrences of
  `0x1e4`; seven are the named actor member sites. Other receivers and
  indirect/computed/inlined accesses are not closed.
- Diary copy constructor `R1550` has no raw E8 or stored-address candidate
  in the file-backed image scan.

**Confidence.** Medium: a bounded negative with explicit unclassified
populations, not a global dead-field proof.

**Unknown.** Actor-owned gameplay purpose.

## Diary LOAD remap and preserved Diary values

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-847 | Diary LOAD reads its two arrays with independent counts and explicitly remaps its trailing reference, writing zero when the lookup fails. | High / Unknown | ✔ promoted | [EXP-0327](../experiments/EXP-0327-diary-consumers/) |
| SAV-848 | Over the bounded original-only complete-reader census every Diary holds complementary values, and the census cannot establish the mutation law. | Medium | ✔ promoted | [EXP-0327](../experiments/EXP-0327-diary-consumers/) |

### SAV-847

- `L08081` calls `L08153` and `L08154`; each reads its own count, calls
  SetSize and reads 4n or 2m bytes.
- It then stores the streamed key at `L07737` and calls `R1371` at
  `L07738`. A lookup success replaces `+2c` with the mapped pointer; a failure
  writes zero.
- Three archive/map-cut vectors retain unequal lengths, noncomplementary values,
  a different mapped Player and a missing reference.
- Player Serialize's subsequent local continuation resolves the hero and
  reconstructs actors without a direct Diary reset.
- This corrects `SAV-669`'s clause that the Diary has no separate LOAD fixup
  ([`retracted.md`](retracted.md)).

**Confidence.** High for the original serializer and resolver and the 3 local
vectors; allocation and archive primitives are explicit cuts.

**Unknown.** First post-LOAD use, intervening callbacks and original acceptance
of malformed counts or references.

### SAV-848

- Direct shipped/owner SAV discovery reaches 85 paths and 61 hashes;
  complete-document parsing succeeds on 82 paths and 59 hashes. The three paths
  that fail that predicate are excluded with their 24 partial Diaries.
- The accepted path population has 465 Diaries, each with 119 elements per
  array: 333 Player-owned, all self-referenced, and 132 actor-owned, all
  null-referenced and default.
- All 55,335 element pairs satisfy word = 1024 - dword. 256 dwords are nonzero,
  maximum 7, at eight indices: 64, 68, 72, 73, 88, 92, 96, 100.
- Actor-owned is the exporter's Humanoid serializer category, not an exact
  concrete-class assertion.
- These are per-path counts, including duplicate original contents, and exclude
  nested generated diagnostic directories.

**Confidence.** Medium for corpus agreement; exact extraction and exclusions are
reproducible. The absence of a dword 17 witness does not strengthen an
unbounded-counter or permanent-complement model; `SAV-842` discriminates them at
instruction level.

## Saved Player identity and formation

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-PLAYERIDENT-830 | Player formation LOAD uses the serializer's receiver, independently of the archive object index and the two stored identifiers. | High / Unknown | ✔ promoted | [EXP-0326](../experiments/EXP-0326-player-formation/) |
| SAV-PLAYERPOP-831 | The measured preserved SAV snapshot agrees on Player order and identifiers, so it cannot select an identity model by itself. | High / Medium | ✔ promoted | [EXP-0326](../experiments/EXP-0326-player-formation/) |

### SAV-PLAYERIDENT-830

- Document LOAD calls the Player manager at `L08155`. `R1558` reads each
  Player reference through `L08109 -> CArchive::ReadObject`, then appends the
  returned object through `R0417`.
- Serialize slot `L08156` names `R0415`. That body restores +04 and +08,
  registers its saved key against its current receiver at `L08157..L08158`,
  and passes that receiver's +30 block to the literal 32-byte read at
  `L08159..L08160 -> R0204`.
- No active-player lookup selects this copy's destination.
- Group ownership uses a separate key: immediate lookup `R1370` can resolve
  an already-registered different Player, or yield null when that key is not yet
  present. The later actor-owner correction does not recopy Group+44
  (`SAV-GRPOWNER-561`).
- Receiver A/B and prior/missing-key vectors discriminate a first-record copy
  and an enclosing-owner fallback.

**Confidence.** High for the named original call chain, receiver, stores and
immediate-lookup law; the vectors' archive and map calls are explicit
normal-return cuts.

**Unknown.** Original acceptance of discordant records, later normalization and
the full loaded lifecycle.

### SAV-PLAYERPOP-831

- Recursive discovery beneath `gameversions/` reaches 129 paths and 76 distinct
  contents. The complete reader accepts 75 contents and refuses one path at its
  invalid header.
- Accepted distinct documents contain 262 Players and 764 Groups.
- All 262 Players have one-based traversal equal to saved signed +04 and saved
  +08. All 764 Group+44 keys equal the enclosing Player key and resolve to that
  object using only prior definitions.
- Formation is 0 in six Player records and 2 in 256; all six zeros belong to
  +28=0 records.
- These are manifest counts, including preserved generated diagnostics, not a
  fixed corpus size or original-acceptance evidence for every input.
- Archive index, traversal, both identifiers and saved key remain separate
  columns.

**Confidence.** High for exact-reader extraction over the named snapshot and the
explicit refusal. Medium for identity inference from its agreement alone: this
census establishes no permuted identifier, discordant Group owner or
multiple-human witness.

## Sack cell lifecycle

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-SACKENTRY-590 | Sack registration has different paths for an existing cell record and for a missing one. | High / Unknown | ✔ promoted | [EXP-0293](../experiments/EXP-0293-sack-cell-lifecycle/) |
| SAV-SACKREMOVE-591 | Sack removal clears the cell's Sack slot without testing whether it is occupied or matches the supplied Sack. | High | ✔ promoted | [EXP-0293](../experiments/EXP-0293-sack-cell-lifecycle/) |
| SAV-SACKPLANES-592 | Empty Sack-cell deletion restores Cost and Static but leaves Dynamic at the preceding recomputation result. | High / Unknown | ✔ promoted | [EXP-0293](../experiments/EXP-0293-sack-cell-lifecycle/) |
| SAV-SACKCALLER-593 | Sack lookup and its callers' effects are separate from the local cell transition; two removal callers do not test the removal return. | High / Medium / Unknown | ✔ promoted | [EXP-0293](../experiments/EXP-0293-sack-cell-lifecycle/) |

### SAV-SACKENTRY-590

- `R0938` keys the cell from `Sack+0x10 -> Position+0x02`. It tests Dynamic
  bit 0 before lookup and returns 0 when that bit is set.
- Existing record: copied as 52 bytes. Any nonzero payload+0x10 refuses,
  including the same Sack pointer. An empty +0x10 is set and the complete
  payload is copied back, returning 1 without calling recomputation or directly
  writing a plane.
- Missing record: creation zeroes 52 bytes and captures current Cost/Static into
  +0x00/+0x01 before Static is ORed with 0x20. It then refetches, stores the
  Sack and calls `R0453` before returning 1.
- The explicit failed-refetch branch returns 0 without a local rollback of the
  preceding creation writes.
- Registration does not directly write Position.

**Confidence.** High for the complete local branches, operands and finite
existing/missing/bit/identity controls. The refetch-failure edge is static-only
under the consistent selected hash helpers.

**Unknown.** Allocation failure, malformed object pointers, external mutation
and whole-runtime reachability.

### SAV-SACKREMOVE-591

- `R0447` uses `Sack+0x10 -> Position+0x02`; a missing hash record returns 0.
- A present record is copied, payload+0x10 is set to zero unconditionally and
  copied back, and `R0453` runs before the deletion predicate. Zero, equal
  and different incoming slot pointers therefore take the same local clear.
- Deletion tests exactly payload+0x04/+0x08/+0x0c/+0x10, byte+0x02 and
  byte+0x2c. Nonzero +0x03/+0x2d..0x33 does not retain the record, and the six
  layer pointers are not independently tested. A retained record returns 1.
- Position and the supplied Sack's fields are not directly rewritten.

**Confidence.** High for the selected complete removal instructions and finite
identity, missing, collision, each-retainer and residue controls. Count 0 with a
nonzero layer is an intentionally inconsistent local predicate control, not a
claimed valid original state. Transitive allocator effects and caller lifecycle
are not inferred.

### SAV-SACKPLANES-592

- After `R0453`, the deletion arm refetches the record, writes payload+0x00
  to Cost and payload+0x01 to Static, unlinks the hash node, and conditionally
  ORs the preserved current Static bit 4 into Static and Dynamic.
- It does not copy the restored Static byte to Dynamic or clear Dynamic bit 5.
- With payload baseline Cost 11/Static 0x02 and incoming Static 0x12, finite
  normal-return execution ends with no node and planes 11/0x12/0x32.
- Nonzero layer+0x20 with count 0 ends 11/0x12/0x37. This distinguishes the
  actual Dynamic write set from a full-plane restoration guess.
- The recompute body never reads Sack slot+0x10; its other occupant, layer and
  baseline inputs still determine its writes.
- Payload destructor `L08161(0)` calls the empty `L08162`. The final hash
  allocation release is an explicit normal-return cut.

**Confidence.** High for direct stores, ordering, the decoded recompute and
finite before/after values.

**Unknown.** Whole-allocator preservation, original loaded scheduling and
gameplay reachability of inconsistent layer/count controls.

### SAV-SACKCALLER-593

- `R0446(Position)` first requires Static bit 5, then a matching hash node,
  and returns payload+0x10 or 0.
- `R0472` uses this accessor before choosing existing-container merge versus
  Sack creation. A failed `R0942` returns before its gold addition and value
  recomputation.
- `R0942` calls cell registration, then `L08163` on success, or the Sack's
  deleting-destructor slot with argument 1 on failure.
- Two selected removal-caller slices, `L08164..L00196` and
  `L08165..L08166`, do not test the removal return before `L02016` and
  `R0448`.
- These are caller dispatch facts, not proofs of complete transfer or
  destruction effects.

**Confidence.** High for the accessor and registration-caller instructions and
the finite dispatch controls. Medium for the bounded larger-caller slice
interpretation.

**Unknown.** In this experiment: container merge, actual destructor and
collection effects, exceptional returns and the complete callsite population.

## Saved crossing local completion

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-CELLENTRY-582 | An existing cell's trigger can be dispatched before its actor entry is refused; a newly created cell takes a different path. | High / Unknown | ✔ promoted | [EXP-0292](../experiments/EXP-0292-saved-crossing/) |
| SAV-CELLFAIL-583 | A refused footprint entry has no local rollback of its completed prefix or its mover caches. | High / Unknown | ✔ promoted (amended) | [EXP-0292](../experiments/EXP-0292-saved-crossing/); [EXP-0466](../experiments/EXP-0466-entry-cost-order/) |
| SAV-CELLLEAVE-584 | The selected actor detach clears a nonzero domain slot without checking that its pointer equals the actor argument. | High / Unknown | ✔ promoted | [EXP-0292](../experiments/EXP-0292-saved-crossing/) |
| SAV-CROSSNEXT-585 | A local cell refusal does not prevent centered completion or normalize the next selected SAVE. | High / Unknown | ✔ promoted | [EXP-0292](../experiments/EXP-0292-saved-crossing/) |

### SAV-CELLENTRY-582

- For domain 1/2, `R0458` reads an existing 52-byte payload, checks for a
  nonzero operation other than 26, and calls a caster helper before testing
  occupied slot +04 at `L02040`.
- A nonzero slot then calls `R0459` and returns 0 without a local occupant
  store or recompute.
- Domain 3 tests +08 without that trigger arm.
- The miss path calls `R1351`, which initializes 52 bytes to zero, copies the
  current cost/static baselines into +00/+01 and inserts the record. Entry
  refetches it and stores the actor, without revisiting the existing-record
  trigger branch.
- The same creator's hit path only copies the existing payload to scratch and
  returns, preserving the saved baseline, tail and residue.
- This extends the retained field meanings of `SAV-CELLLOAD-111`; it is not a
  new layout.

**Confidence.** High for the named local ordering and the finite
new/existing/domain/occupied vectors. A synthetic mutating caster cut
demonstrates why caller preservation is not whole-callee preservation.

**Unknown.** Caster effects and callback scratch purity.

### SAV-CELLFAIL-583

- Before the per-cell loop, `R0050` writes mover+72 and +82..85. In store
  order:
  - +82 receives the cached low byte of Position cell-X getter `R0299`;
  - +83 the cached low byte of cell-Y getter `R0300`;
  - +84 the low byte of full-X getter `R0165`;
  - +85 the low byte of full-Y getter `R0166`.
- By `SAV-TOKENPOS-074` these last low bytes are Position+04/+05.
- The first two getter results are stored in caller locals before call
  `R1345`. After a normal return, mover stores run in order +72, +82, +83,
  +84 and +85; the full-coordinate getters run after the call. Word +72 receives
  that call's 16-bit return value. The probe's `0x1357` return is an explicit cut, not an original
  runtime value or calculation. `SAV-1161` verifies the native local source/order.
- All four getter bodies are accessor cuts in this probe; their source laws come
  from the cited prior claim. No atomic four-byte snapshot or callback purity is
  implied.
- Each successful `R0458` writes its own record. On the first false return,
  `L07116..L08167` returns 0 directly; no earlier-cell undo or cache restore
  occurs.
- A 2x2 synthetic footprint with a conflict at each ordinal preserves exactly
  the earlier row-major actor-slot prefix.
- The all-or-nothing headline of `TERR-FOOTPRINT-147` is therefore withdrawn,
  while its loop geometry remains valid ([`retracted.md`](retracted.md)).

**Confidence.** High for the complete local refusal path and five controlled
footprint cases per image.

**Unknown.** Unexpanded notification and caster effects, reachability of a
conflict in valid gameplay, and later repair.

**Amended.** `SAV-1161` distinguishes pre-call caller-local caching from
post-return mover stores. The prefix-preserving refusal result and its
normal-return scope stand; this is a wording clarification, not a retraction.

### SAV-CELLLEAVE-584

- `R0057` refuses a missing node or a zero selected slot. Domain 1/2 clears
  +04 and domain 3 clears +08 after only a nonzero test. It copies the changed
  52 bytes back and recomputes.
- Deletion then requires all four occupant dwords +04/+08/+0c/+10, layer-count
  byte +02 and operation byte +2c to be zero. Residue +03/+32/+33 and the
  remaining trigger-tail bytes do not retain a record.
- The deletion arm restores the cost/static baselines, invokes removal, and
  preserves plane bit 4.
- After recompute and any removal call, success writes, in order:
  - mover+86 from byte[Position+00] at `L07123`;
  - +87 from byte[Position+01] at `L08168`;
  - +88 from byte[Position+04] at `L08169`;
  - +89 from byte[Position+05] at `L07124`.
- These are direct field loads, not the entry probe's accessor cuts; Position is
  reloaded from actor+10 before each load. An early refusal omits those stores.
- Own-pointer and distinct-pointer inputs follow the same selected clear path.

**Confidence.** High for the exact local tests, stores and call order, with 18
detach vectors per image. Node allocator and removal internals are controlled
cuts.

**Unknown.** Malformed identity consequences, later repair and complete callback
behavior.

### SAV-CROSSNEXT-585

- In `R0047`, a false detach ends the old-footprint loop but continues
  through `R0058`, the new Position stores and entry. The return from
  `R0050` is not tested before the existing snap test.
- `R0039` then tests fractions, not entry success. The known centered cleanup
  clears mover+aa/+ac/+a8/+80/+a6 and the dynamic route at actor+178; the static
  route at +15c survives these direct instructions.
- Progress 3's center test can clear order+09/action+54 despite a refused entry.
- Under explicit returning cuts, the composed refusal vector retains the
  conflicting destination slot, centers the actor and empties its dynamic route,
  and the isolated selected serializers emit those current values.
- The Unit prefix serializes the static then the dynamic list before the
  180-byte mover. Token emits Position's 12 bytes, and the cell serializer emits
  the current node key plus 52 payload bytes.
- This is no rollback, atomic snapshot or original resave guarantee.

**Confidence.** High for the named caller branches and finite normal-return
instruction vectors. Existing step arithmetic, key repair and ordinary route
cleanup are reused, not newly inferred.

**Unknown.** First actual post-load scheduling, whole SAVE chronology, operation
26 relocation and unexpanded effects.

## Dynamic Group initialization and selected first SAVE

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-GRPNEW-576 | The selected Group and base constructors preserve incoming Group+1c but explicitly construct an empty embedded Group+20 list. | High / Unknown | ✔ promoted | [EXP-0291](../experiments/EXP-0291-dynamic-group-first-save/) |
| SAV-GRPALLOC-577 | The selected existing-page small allocation path does not clear a newly returned Group payload. | High | ✔ promoted | [EXP-0291](../experiments/EXP-0291-dynamic-group-first-save/) |
| SAV-GRPCMD-578 | The selected ordinary-command prologue neither mints Group+1c nor fills Group+20; its wrapper deletes at most the first rejected old Group per invocation. | High / Unknown | ✔ promoted | [EXP-0291](../experiments/EXP-0291-dynamic-group-first-save/) |
| SAV-GRPFIRSTSAVE-579 | For Group+1c and the Group+20 list, constructor provenance, authored-map provenance, LOAD and selected SAVE are different boundaries. | High / Unknown | ✔ promoted | [EXP-0291](../experiments/EXP-0291-dynamic-group-first-save/) |

### SAV-GRPNEW-576

- `R0149 -> L08170 -> R1333 -> L08171` writes the base through +18,
  not +1c.
- `R1545` stores vtable `L08172` at +20, zeros +24/+28/+2c/+30/+34 and
  writes grow-by 10 at +38.
- Group construction supplies +3c and zeros +40/+44.
- Two byte-pattern sentinels and two valid mapped source-pointer sentinels
  survive at +1c. An incoming populated embedded-list header becomes empty
  without dereferencing its old head. The pointer sentinels are instrument tags,
  not field semantics.
- Both lawful images decode identically.

**Confidence.** High for the named normal-return constructor writes and
controlled sentinel discrimination.

**Unknown.** Allocation contents, later callbacks, Group+20 element meaning and
first genuine SAVE values.

### SAV-GRPALLOC-577

- `R0196` requests 72 bytes through
  `R0747 -> R0722 -> R0748 -> R1567`.
- The image-initialized threshold is 480. Requests 72/80/28 round to 80/80/32
  and enter `R1568` with 5/5/2 units.
- The controlled existing-page success through `L08173` writes allocation
  metadata, returns `page+100` and makes zero writes into the returned payload.
  Two sentinels survive all 24 locale/size/route allocation cases.
- The alternative threshold-zero fixture reaches import slot `L08174` with
  flags 0 and the rounded size. Its returned contents are explicitly supplied by
  a cut, and the caller does not clear them afterward.

**Confidence.** High for the finite selected route, rounding, flags and caller
stores. Nothing is asserted about actual OS bytes, the runtime threshold, fresh
or reclaimed pages, failure handlers or every allocator route. A zero-filled
fixture is not a constructor default.

### SAV-GRPCMD-578

- `R0196` exits the scan after a zero-return `R0197` deletion through
  `L00536 -> L00537`. This contradicts the former every-rejected-Group
  clause of `AI-CMD-033` ([`retracted.md`](retracted.md)).
- Prologue `L00532..L08175` creates a Group, appends resolved actors and
  appends the new Group. Group append writes actor+70 and Group+44, not the
  queried fields.
- Guard `R0125` and its expanded direct summary `R0140` write GroupAI
  through +3c, not Group+1c or the embedded list.
- Controlled Guard callbacks can change both queried fields, and the remaining
  local setter preserves that change.

**Confidence.** High for the direct bodies, branch order and the finite
one-found, two-found and one-missing prologue vectors. The shared list and
delete, actor helper, detach and behavioral callbacks are cuts.

**Unknown.** Whole-command preservation and other setters.

### SAV-GRPFIRSTSAVE-579

- Authored miss `L08176..L08177` writes loader+40 to Group+1c at `L00296`.
- Group LOAD restores its list and the literal saved +1c. Controlled
  `[1234,abcd]`/`10203040` values survive the next selected store.
- Group SAVE emits the current Group+20 list, AI and members before reading
  +1c/+40/+44.
- Under preserving cuts, the composed command-prologue-to-store emits the
  incoming selector and an empty list.
- A changed actor-serialization cut installs a new list and selector after the
  embedded-list store: that invocation emits the old count 0 but later emits the
  changed selector `13579bdf`.
- The selected serializer has no direct normalization of either queried field.

**Confidence.** High for the selected instruction sequence and controlled source
discrimination. These are not real first-SAVE observations, atomic snapshots or
safe authored defaults.

**Unknown.** Whole-runtime chronology, transitive mutation and actual first-save
values.

## Restored Group AI continuation

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-GRPDISPATCH-568 | One Group dispatch selects its order once; an empty Group at entry forces ff, while a Group emptied during the dispatch does not. | High | ✔ promoted | [EXP-0290](../experiments/EXP-0290-restored-group-ai-continuation/) |
| SAV-GRPMUTATE-569 | The caller's next-member rule is an identity search in the current list, including for Stand Ground, ff and withdrawal. | High / Unknown | ✔ promoted | [EXP-0290](../experiments/EXP-0290-restored-group-ai-continuation/) |
| SAV-GRPPATROL-570 | The AI-owned patrol path and a member's active ring are separate ordered lists, and the path-copy setter reverses the source. | High / Unknown | ✔ promoted | [EXP-0290](../experiments/EXP-0290-restored-group-ai-continuation/) |
| SAV-PATROLCURSOR-571 | The saved patrol cursor is a cell value, not a persisted node position. | High / Unknown | ✔ promoted | [EXP-0290](../experiments/EXP-0290-restored-group-ai-continuation/) |
| SAV-GRPSAVENEXT-572 | The selected next-SAVE consumers read current ordered state, not a retained input payload or constructor defaults. | High | ✔ promoted | [EXP-0290](../experiments/EXP-0290-restored-group-ai-continuation/) |

### SAV-GRPDISPATCH-568

- The complete `R0023..L08178` reads the Group count before AI+20, writes
  ff only at `L00300`, and reads the byte at `L08179`.
- Its six-entry table and separate 11/ff arms lead to the shared tail at
  `L00483`; all other byte values also reach that tail. The cell argument for
  arms 2/4/5 comes from AI word+0a.
- A returning arm that changes AI+20 does not cause a second selection in this
  invocation.
- The tail rereads the current count and head. An arm that removes every member
  leaves its previous order until a later dispatch's entry test.
- Independently decoded EN/RU instructions give 512 zero/one-member byte-order
  vectors per root, including all 256 values. Controlled callbacks return
  normally and are not asserted pure in ROM1.

**Confidence.** High for the complete local switch, table, single selection and
fresh tail entry. Raw-image extraction and all local branch targets close the
selected body, not transitive callees or original loaded chronology.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### SAV-GRPMUTATE-569

- Order 0's inline search and helper `R0071` search the Group head-to-tail
  for the just-dispatched actor, then return its current successor; a missing
  actor or final node returns null.
- Order 3, ff and the shared withdraw tail call that helper at
  `L08180/L08181/L08182`.
- Sixteen controlled vectors per root show that removing the current actor
  terminates that stage, appending a later actor can extend it, changing the
  order does not reselect it, and the withdraw tail starts again from current
  membership. Its first successful predicate calls `R0100`; otherwise its
  second successful predicate calls `R0107`.
- Removing the current actor during a tail predicate still permits that actor's
  remaining selected tail calls before the successor lookup.

**Confidence.** High for these caller and helper programmes and the explicit
mutation-cut traces.

**Unknown.** Actual callback mutation, malformed or cyclic membership,
deallocation safety and termination under arbitrary repeated mutation.

### SAV-GRPPATROL-570

- In the complete setter `R0164..L08183`, the source is Group+3c -> AI+4c
  and the destination is actor+158 -> order+90.
- The source is walked head-to-tail, but each value is inserted at the
  destination head (`L08184..L08185`).
- The setter chooses order+02 from that head at `L08186/L08187`, then
  prepends the actor's own cell (`L08188..L08189`).
- Source A,B,C therefore produces own,C,B,A with current waypoint C, preserving
  duplicate values.
- An empty source reaches the unguarded first-value read at `L08190`; the
  published mission-INI route's nonempty admission remains separate.
- This setter is not a direct call of the ordinary Group dispatcher.

**Confidence.** High for the complete setter data flow, pointer bases and
controlled nonempty/duplicate/empty vectors. Stop, allocator and list primitives
are explicit cuts.

**Unknown.** Original invocation after LOAD and unexpanded mutation effects.

### SAV-PATROLCURSOR-571

- Actor-order LOAD `R0199` restores 148 raw bytes, replaces order+90 with a
  fresh counted-word list, and leaves the raw word+02 cursor plus the dword+04 latch
  intact at this local boundary.
- Walker `R0159` consumes a nonzero +04 by updating post+00 to the current
  cell and clearing +04, then calls guard.
- Only a guard result order+08 equal to 0 or b reaches the patrol continuation.
- That continuation sets +04 to 1 even when not yet arrived. On arrival it
  searches the active list for the first value equal to +02, takes its successor
  or wraps to the head, and writes +02, move order+08 and destination+0a.
  Duplicate cells therefore use first-equal semantics.
- A missing current value or an empty arrived list reaches `L01779` with null;
  no fallback precedes that dereference.

**Confidence.** High for the complete serializer and walker and eight
discriminating cursor vectors per root. Controlled unmapped reads identify the
local invalid-pointer edge, not an observed original crash or a claim that valid
shipped setters produce that state.

**Unknown.** Guard effects and original acceptance.

### SAV-GRPSAVENEXT-572

- Group `R1143` serializes the Group+20 word list, the current 80-byte AI
  block and the current AI+4c word list, then the current members plus dwords
  +1c/+40/+44.
- The word-list writer emits the count and walks nodes to null, writing each
  current 16-bit value; it does not persist node identity or an iterator.
- AI store writes its current 80 bytes before dereferencing its current +4c
  list. Actor-order store similarly writes 148 bytes before its current +90
  list.
- Controlled LOAD -> dispatch/patrol -> store vectors retain changed order and
  cursor values and newly allocated list pointers in the raw block, while
  emitting the ordered values separately.
- The direct Group dispatcher neither reads Group+20 list elements nor
  dereferences AI+4c; this does not make either list unused outside that
  selected body.

**Confidence.** High for the selected normal-return serializer sequence and
finite controlled vectors. Archive, allocation, actor serialization and
behavioral callbacks are explicit cuts; no atomic snapshot, safe authoring
defaults, all-path producer set or original resave is established.

**Unknown.** Group+20 semantics, the remaining unnamed AI fields, whole-linker
mutation, and first-loaded and next-SAVE runtime chronology.

## Saved Group reconstruction

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-GRPLOAD-560 | Group LOAD reconstructs membership through ordinary append before restoring its trailing fields. | High / Unknown | ✔ promoted | [EXP-0289](../experiments/EXP-0289-saved-group-rebuild/) |
| SAV-GRPOWNER-561 | Enclosing Player, actor Token owner and saved Group owner are not one restored value. | High | ✔ promoted | [EXP-0289](../experiments/EXP-0289-saved-group-rebuild/) |
| SAV-GRPIDENT-562 | A saved Group+1c is consumed on normal resume without being reminted as an authored id. | High / Unknown | ✔ promoted | [EXP-0289](../experiments/EXP-0289-saved-group-rebuild/) |
| SAV-GRPAI-563 | Restored Group AI is neither constructor defaults nor an untouched 80-byte pointer-bearing blob. | High / Unknown | ✔ promoted | [EXP-0289](../experiments/EXP-0289-saved-group-rebuild/) |

### SAV-GRPLOAD-560

- Player group-list loader `R1340` allocates and constructs a fresh 0x48-byte
  Group, calls `R1143`, then appends that Group to the collection after the
  serializer returns (`L07279/L08191/L08192`).
- Group first restores the embedded word list and the owned 80-byte AI block and
  list, and clears its actor CObList. It then reads each Unit reference and
  calls ordinary append `R0152` (`L08193/L08194`), not generic list LOAD
  `R1341`.
- Append detaches an existing `actor+70` link, AddTails, stamps
  `actor+70 = Group`, then copies the current `actor+14` to Group+44.
- After all actors, Group reads the literal dword+1c and separately reads and
  remaps +40 and +44 through the current identity map. Both remappers return
  null on a missing key; neither validates membership or a target class in its
  selected body.
- The order is the established SAV-GRPORD-058 rule; this claim adds restoration
  precedence and the separate continuation, not a new ordering.

**Confidence.** High for the complete named serializer and helper bodies and the
normal-return local precedence.

**Unknown.** Malformed references, allocation failure, reused nonfresh lists and
original acceptance.

### SAV-GRPOWNER-561

- Player registers its identity at `L08158` before loading groups.
- Token registers its own key and resolves actor+14 at `L08195/L08196`.
  Group append can read that owner, but `L08197/L08198` later replace
  Group+44 from its independent saved key.
- After all groups, the 32-byte Player block and the Diary, Player remaps its
  hero reference, flattens group members into Player+20 and stamps every visited
  actor+14 with the enclosing Player (`L01794/L07271`). That suffix does not
  recopy Group+44 or change actor+70.
- A mapped discordant Group owner can therefore differ from the corrected actor
  owner at this boundary. A missing Group key becomes null, without falling back
  to the enclosing Player.
- Nine bounded original-instruction vectors with explicit archive/MFC cuts
  discriminate the append-wins, Group-tail-wins and Player-normalizes-everything
  models.

**Confidence.** High for the local stores, reads and order. The vectors are
controlled instruction execution, not original SAV acceptance or proof of
arbitrary callee preservation.

### SAV-GRPIDENT-562

- Group LOAD copies Group+1c literally.
- The normal world-resume call `L08199/L01841` passes zero to `R0067`.
  Its group loop indexes each restored Group using +1c at
  `L08200/L01812/L01813` before testing that argument.
- Zero skips the owner+44-based initial-stance branch at `L00391`, not the
  index insertion.
- Authored map creation instead compares and writes Group+1c from the type-6
  loader value (`L00292/L00293/L00296`, already AI-GROUP-009).
- Constructor `R0149` and dynamic allocation wrapper `R0196` have no
  direct +1c store.
- This does not establish all dynamic first-save producers or the harmlessness
  of an arbitrary saved value. File-local pointer keys and this raw group
  selector are different mechanisms.

**Confidence.** High for the literal restore, the selected resume branch and
index, and the authored contrast.

**Unknown.** All-path dynamic initialization, collision consequences and first
normal-tick chronology.

### SAV-GRPAI-563

- `R1350` conditionally deletes the constructor-created owned word list,
  reads 0x50 raw bytes into `*(Group+3c)`, allocates a fresh word list, replaces
  AI+4c, then restores that list (`L08201/L08202/L08203/L08204`). Saved
  AI+4c pointer bits are therefore not used as the restored list identity.
- Group tick `R0023` reads the current membership count. Zero overwrites AI
  order byte+20 with ff before dispatch, while a nonempty Group reads the
  restored or current order.
- Arm 0 reads each actor+14, dispatches it, then searches the live Group list
  for that actor before taking the current successor.
- Rate helper `L08205` reads actor+70 -> Group+3c -> AI byte+44, otherwise
  actor word+8c.
- The selected later link arm can set AI+48 through a resolved actor's Group
  (`L03530/L03531`); skipping the initial stance does not prove that the
  whole linker preserves AI.
- Group+40's gameplay meaning and the embedded Group+20 word-list meaning remain
  unclassified.

**Confidence.** High for the wire and pointer replacement and the named local
consumers.

**Unknown.** Transitive dispatch effects, complete AI-field meanings, exact
loaded chronology and original observation.

## Equipment event order

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-EQUIPORDER-552 | The six accepted equipment methods, attach and removal for Armor, Shield and Weapon, do not share one event order. | High / Unknown | ✔ promoted (amended) | [EXP-0288](../experiments/EXP-0288-equipment-event-order/) |
| SAV-EQUIPEFFECT-553 | Equipment Effect traversal is forward and reads each next node incrementally, not a pre-collected batch or a reverse removal. | High / Unknown | ✔ promoted | [EXP-0288](../experiments/EXP-0288-equipment-event-order/) |
| SAV-EQUIPCALL-554 | Ordinary command 22 adds container and load events but no unconditional final Human derive. | High / Unknown | ✔ promoted | [EXP-0288](../experiments/EXP-0288-equipment-event-order/) |
| SAV-EQUIPOBS-555 | Intermediate equipment states have named consumers; there is no atomic final-state contract. | High / Unknown | ✔ promoted | [EXP-0288](../experiments/EXP-0288-equipment-event-order/) |

### SAV-EQUIPORDER-552

- Armor attach completes same-slot item removal, stores its slot, adds both
  defensive blocks, refreshes positive weight, sets flags, then applies Effects.
- Armor removal refreshes negative weight, subtracts both blocks, clears its
  slot, sets flags, then removes Effects.
- Shield attach removes the old shield, conditionally actor-removes and
  reinserts a two-handed weapon, stores the new shield, adds blocks, refreshes
  weight, applies Effects, then sets flags.
- Shield removal refreshes negative weight, subtracts blocks, removes Effects,
  sets flags, then clears its slot.
- Weapon attach prepares its owned Spell before old-weapon removal and
  conditionally actor-removes and reinserts a shield. It stores the new weapon,
  writes attack modifiers and the active skill, calls actor derive, then writes
  conditional timing and range fields, refreshes weight, sets flags and applies
  Effects.
- Weapon removal removes Effects first, writes attack modifiers and the active
  skill, and derives. It subtracts the currently reread `u8(Weapon+0x50)-1` from
  actor range byte `+0x12c` with byte wrapping (`L07756..L07757`), assigns
  timing bytes `+0x134=8` and `+0x135=4`, refreshes negative weight, sets flags,
  deletes and clears its nonnull owned Spell, and finally clears the actor
  weapon.
- Attach adds the same current-byte range delta after derive
  (`L04088..L04091`). Neither range operation directly reloads a definition
  parameter.
- Intervening callback mutation is not excluded, so the cycle is not a
  guaranteed restoration.
- Same-slot eviction uses item `+3c`; opposite-hand eviction uses actor `+40`.

**Confidence.** High for the named local control-flow ordering on normal-return,
valid, nonaliased paths, including the optional displacement arms.

**Unknown.** Runtime occurrence, callback mutation and exceptional completion.

**Amended.** The range-restoration scope is narrowed. Weapon removal was first
said to derive and then restore range and timing; [`retracted.md`](retracted.md)
replaces that shorthand with the local arithmetic above: a reread byte delta
subtracted with byte wrapping, literal timing 8/4, and no proven inverse cycle
restoration (EXP-0288).

### SAV-EQUIPEFFECT-553

- Both walkers use `L08206/L08207 -> L08208`: they save the current node's
  next pointer into iterator state before returning its Effect pointer, call
  that Effect, then fetch the pending node.
- Twelve original-instruction walker controls distinguish empty, single, AB, BA
  and two callback-cut mutations. Changing A's next after A was fetched still
  visits B,C; changing pending B's payload or link yields A,D. These injected
  cuts demonstrate iterator sensitivity, not actual Effect mutations.
- For Token state 0, each normal-return application or removal runs general
  `R0843` with multiplier +1/-1. All 50 static kind entries reach actor `+50`
  at `L06977`, including kind 41's empty local arm, before the next Effect.
- State 8 apply and states 12/17 apply/remove have different branches. Arbitrary
  Effect subclass overrides and transitive callbacks are not covered.

**Confidence.** High for the bounded iterator instructions, the twelve
controlled traces and all 50 normal-return kind paths.

**Unknown.** Actual list mutation by callbacks, runtime reach and unlisted
overrides.

### SAV-EQUIPCALL-554

- Carried source 2 first takes the requested quantity through `R0095`.
- Destination 1 calls actor `+38` at `L08209`, reinserts a nonnull same-slot
  return at the source index, then calls `R0451(0)` at `L04617`.
- Equipment source 1 calls actor `+40`; destination 2 inserts at the requested
  destination index, then refreshes zero weight at `L08210`.
- `R0451` compares the prior and refreshed signed load-capacity quotients and
  calls actor `+50` only when they are unequal.
- Raw vtables select Humanoid/Human wrappers `R0933/R0934`, which add no
  trailing `+50/+54`. Unit wrappers add `+54`, whose Unit target only returns
  actor `+1c`. The Humanoid/Human `+54` target is the already-published
  Token-value computation, not equipment derive.
- Command notification follows the selected load refresh.

**Confidence.** High for the named command arms, the wrapper split and the load
predicate.

**Unknown.** Full UI admission, quantity and split semantics outside the
published controls, notification side effects and runtime chronology.

### SAV-EQUIPOBS-555

- Armor/Shield negative-weight refresh precedes defensive subtraction. A changed
  quotient enters Humanoid derive while that item body has not yet removed the
  old defensive modifier contribution.
- Derive clears live defence, then folds retained actor `+fe` through
  `R0840 -> R0875 -> R0845`.
- Weapon's explicit derive precedes its timing/range and per-item weight update
  on both attach and removal.
- At the local Effect-removal boundary Armor has cleared its slot, Shield has
  not, and Weapon has not yet done its direct removals or owned-Spell deletion.
  These are store-order facts, not a claim that every Effect reads those slots.
- Actual reads include state 8 apply's current actor `+c6/+94` and general
  dispatch's current `+d8`; each state 0 Effect can enter derive before the
  next.
- Earlier callees can mutate the same fields.

**Confidence.** High for the named store, call and read frontier. No callback
purity is assumed.

**Unknown.** A complete intervening write set or runtime snapshot. Complete
transitive aliases, nested reentry, faults and nonreturns, first-SAV chronology
and runtime visible values also remain open.

## Saved actor definition binding

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-ACTORBIND-544 | Loaded Unit, exact Humanoid and Human have different definition laws; every saved Human type word 33..65535 selects row 5. | High / Unknown | ✔ promoted | [EXP-0287](../experiments/EXP-0287-actor-definition-binding/) |
| SAV-ACTORDISPLAY-545 | The saved display backing dword `+148` and the effective byte `+14c` can diverge during load. | High / Unknown | ✔ promoted (amended) | [EXP-0287](../experiments/EXP-0287-actor-definition-binding/), [EXP-0318](../experiments/EXP-0318-human-unit-residuals/) |
| SAV-ACTORCTOR-546 | Archive construction is not actor reconstruction from definitions after restore. | High / Unknown | ✔ promoted | [EXP-0287](../experiments/EXP-0287-actor-definition-binding/) |
| SAV-ACTORINPUT-547 | The selected load and hook slice retains saved mover and visual inputs, not a fresh definition projection. | High / Unknown | ✔ promoted (amended, superseded) | [EXP-0287](../experiments/EXP-0287-actor-definition-binding/) |
| SAV-ACTORLIMIT-548 | A complete local binding law proves neither first-consumer safety nor exact-class acceptance. | High / Unknown | ✔ promoted | [EXP-0287](../experiments/EXP-0287-actor-definition-binding/) |

### SAV-ACTORBIND-544

- Unit `R0210` dispatches `vt+30`. Unit returns 0 and writes
  `+3c = [L08211] + 48*u8(+0c)`; Humanoid/Human return 1 and clear `+3c`.
- Exact Humanoid only appends XP and references.
- Human `L08212` then writes
  `+3c = [L08213] + 48*(u16(+0e)<33 ? u8(+0c) : 5)`. The type word is
  zero-extended before the signed compare, so every value 33..65535 takes row 5.
- Lookup pairs `R0500/L08214` and `R0496/L07761` read no table size
  and perform wrapped 32-bit arithmetic, unlike Item's bounded load.
- Neither selector is rewritten by these suffixes.

**Confidence.** High for the complete named suffixes, the original vtable
predicates and the lookup helpers, independently decoded and executed in bounded
vectors.

**Unknown.** Valid external row contents and original input acceptance.

### SAV-ACTORDISPLAY-545

- Unit restores `u32 +148`, copies its low byte to `+14c`, then clears the byte
  for exact Unit. Exact Humanoid preserves it.
- Human keeps a zero. A nonzero byte is cleared if the newly selected `+3c` is
  null or name search `R0996(+3c+4,"NPC")` returns -1; otherwise the original
  byte survives, not boolean 1.
- These suffixes leave `+148` unchanged; the later local Unit store writes
  `u32 +148 = u8 +14c`.
- Sender `R0059` conditionally uses the effective byte to index a separate
  collection at `L02112`, not Data's actor row.

**Confidence.** High for the named stores, conditions and sender inputs. The
original instruction vectors injected the explicitly named string-search return.
A confirmation executes the actual CString and substring routines in single-byte
mode on NPC/xNPCy/npc/Ordinary with incoming 0/1/2/255: it preserves nonzero
values on a match and never enables zero.

**Unknown.** Complete backing-collection initialization, intervening mutations
and actual presentation.

**Amended.** The substring confirmation in the Confidence paragraph was added to
the injected-return vectors; EXP-0318 is cited beside EXP-0287.

### SAV-ACTORCTOR-546

- The three creator bodies reach Unit `R0901`, Humanoid `R0902` and Human
  `L08215`, with base chaining and distinct installed vtables.
- Unit allocates fresh mover and order objects at `+154/+158` and initializes
  selected selectors. Human additionally calls `R0656("Man_Unarmed",0,0)`
  before archive restore.
- The load restores Token row/type and Unit scalar and raw inputs, then runs
  `SAV-ACTORBIND-544`, not the name/table/equipment initializer or its
  nonzero-mode type rewrite.
- Actor `+154/+158` pointers are construction-owned, while their target
  180/148-byte contents are SAV-owned. Order load additionally replaces target
  `+90` with a newly allocated list.

**Confidence.** High for the creator, constructor and named load instruction
order.

**Unknown.** Transitive allocation and callback effects, and a whole-object safe
authored default vector.

### SAV-ACTORINPUT-547

- Unit loads `+49/+4a/+4b/+4c` literally. Footprint and domain getters read the
  first two, while sender `R0059` exports the low byte of type `+0e` and face
  `+4b` under mask `0x4000`.
- The raw mover restore includes mask byte `+5`.
- `R0471` maps domain 1/2/3 to mask `0x41/0x44/0x82` and otherwise preserves
  the mask. Its two located direct callers are constructor routes, not the
  selected serializer or hook bodies.
- Post-read `R1482/R1569` conditionally repairs mover and order
  identities, without direct mask-builder or actor-table calls.
- The order call `R0170` requires actor stage `+13c == 0`. Its seven slots
  (`SAV-HUMRESUME-460`) leave zero and missing keys unchanged and replace hits
  with mapped values. This is not blanket nulling or proof that every restored
  key becomes a valid pointer.
- This slice establishes no consistency normalization between the separately
  saved domain and mask.

**Confidence.** High for the named literal reads, getters, raw coverage and the
direct-body and callsite facts.

**Unknown.** Computed aliases, later refreshes and first post-load rendering and
movement.

**Amended.** The reference-repair shorthand is made field-specific: the order
helper runs only at stage `+13c == 0` and leaves missing keys unchanged.
[`retracted.md`](retracted.md) supersedes the shorthand and withdraws the
EXP-0277 report's "missing entries become null" generalization: `R1488`
returns false without writing its output, and the caller writes only on success.

### SAV-ACTORLIMIT-548

- Human conditionally consumes its freshly selected definition in the same
  serializer name check.
- Unit and exact Humanoid share `vt+58 -> L08216`, whose initial `+3c+8`
  parameter lookup has no preceding local null guard; Human overrides that slot
  with a `+1c`-based body. Humanoid's local null result is therefore not proof
  that the pointer is unused.
- No original was launched. A later callback or computed consumer can still
  alter or consume the same fields before a first frame, move or save.
- Static descriptor presence and bounded vectors do not establish that an
  exact-class Humanoid SAV reaches that method.

**Confidence.** High for the static consumer frontier and the vtable split.

**Unknown.** Global first-consumer order, exact-class production and acceptance,
malformed-row behavior, external payload parity and original load/state/resave.

## Save body codec and framing

Spec: [`formats/sav/format.md`](../formats/sav/format.md).

- The body `[0x14, u32@0x04)` is a word-oriented run/literal code (EXP-0046).
  Decoding it resolves `SAV-UNK-006`, corrects `SAV-PTR-003`'s "descriptor
  head", and puts EXP-0041's predicted block-plane records in the file.
- EXP-0048 walks the decoded stream end to end: the object framing closes in all
  four saves (393/393 tag words), every object is attributed, the map's own
  placements reproduce the population, and the stream tiles with ≤ 70 bytes
  unattributed per save.
- EXP-0046 and EXP-0048 read no `rom.exe` routine, so what they alone establish
  is carried by corpus arithmetic and named excluded rivals, never by a listing.
  Their corpus is four saves of one map (`scn:10.alm`, 80×80, two play
  sessions). That bounds every claim about shape and does not bound the
  identities that close on quantities the model is not given.
- EXP-0077 read five routines: the save writer `R0804`, the save reader
  `R1305`, the world `Serialize` `R0414`, the compressor
  `R1376` and the decompressor `R1515`. The decode of all four
  saves is identical under both framings. Three corpus-model statements about
  where the fields live are wrong, and the world half of a save is optional: see
  `SAV-FRAME-021`…`SAV-TRAIL-026`.

## Minimal generated-child instrument

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-MINCHILD-536 | The minimal generated-child instrument reaches ordinary children but not the selected profile-free AppContainer child. | Unknown | ✔ promoted | [EXP-0286](../experiments/EXP-0286-minimal-child-creation/) |

### SAV-MINCHILD-536

- Two three-call runs return four suspended baseline primaries through ordinary
  and empty-extended startup. Their queried tokens are primary, non-AppContainer
  and medium-integrity, and their exact job membership is observed. Their own
  token/job self-check exits with `0x2860003f` before the primary handles
  signal.
- Both SECURITY_CAPABILITIES submissions instead fail CreateProcessW with
  immediate numeric error 2, PID 0 and no child token, wait or exit.
- Pointer-escape hardening does not change either result.
- All four baseline job-accounting pairs read total 1/active 1 before resume and
  total 2/active 1 after primary exit. No member census identifies that extra
  count, so the configured active-process limit 1 is not a verified one-process
  population or a second-child denial.
- Only four fresh directory and executable descriptors receive one
  non-inheriting package RX ACE each; bytes and final configured descriptors
  survive read-only postchecks.
- No original process, profile, registry, privilege, ancestor ACL or install
  mutation occurs.

**Confidence.** Unknown. API, token and exit values are bounded instrument
observations only.

**Unknown.** Profile-free child feasibility, the object or cause behind error 2,
the unexplained job population, and all confinement and original-runtime
properties.

## Regeneration arithmetic and stores

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-REGENWIDTH-528 | Admitted regeneration consumes signed words and an unsigned remainder byte, not a widened normalized pool. | High / Unknown | ✔ promoted | [EXP-0284](../experiments/EXP-0284-regen-store-semantics/) |
| SAV-REGENSTORE-529 | Each regeneration arm writes the remainder, the narrowed pool, then the signed upper-bound result, and the remainder survives the bound. | High / Unknown | ✔ promoted | [EXP-0284](../experiments/EXP-0284-regen-store-semantics/) |
| SAV-REGENFAULT-530 | Health skips a zero regeneration period while a reached mana deficit arm faults on it; multiplication overflow is not a divide fault. | High / Unknown | ✔ promoted | [EXP-0284](../experiments/EXP-0284-regen-store-semantics/) |
| SAV-REGENORDER-531 | The regeneration routine does not transact health and mana together: health's stores and callback precede mana's tests. | High / Unknown | ✔ promoted | [EXP-0284](../experiments/EXP-0284-regen-store-semantics/) |
| SAV-REGENWIRE-532 | The persistent regeneration operands keep their Unit wire widths; the serializer applies no arithmetic normalization. | High / Unknown | ✔ promoted | [EXP-0284](../experiments/EXP-0284-regen-store-semantics/) |

### SAV-REGENWIDTH-528

- In `R0654`, current/max/period at `94/96/98` and `9a/9c/9e`, and rate
  modifiers `de/e2`, are sign-extended from 16 bits; `a2/a3` are zero-extended
  from 8 bits.
- Modifier plus 100, the products and the final accumulator operate modulo 2^32.
- The health product is `2*maximum*(modifier+100)*rate`; mana omits the 2.
- Signed IDIV occurs after both multiplications, truncates toward zero and
  leaves a remainder with the dividend's sign. The wrapped accumulator is
  divided by 100 separately for remainder and quotient.
- Twelve alternate scalar models have different stored outcomes on each arm
  across 24 discriminators.

**Confidence.** High for the identified image and the decoded arithmetic under
valid stable memory.

**Unknown.** Ordinary producer reach of synthetic edge operands.

### SAV-REGENSTORE-529

- Health stores are `L08217` byte `a2`, `L08218` word `94`, then `L04038`
  word `94`. Mana stores are `L08219` byte `a3`, `L08220` word `9a`, then
  `L04040` word `9a`.
- The comparison rereads the stored pool as i16; there is no lower clamp.
- Current 32766, max 32767, period 1, modifier 0, remainder 0 and rate 1 gives
  health 32764, not 32767.
- Negative remainders store their low byte and later reload unsigned: mana 0,
  max 101, period 1 and modifier -101 produce (-1,255), then (0,54) on another
  admitted call with unchanged inputs and returning nonmutating callbacks.

**Confidence.** High for the local stores and conditional repeated arithmetic.

**Unknown.** Runtime callback effects and actual repeat chronology.

### SAV-REGENFAULT-530

- Health period 0 skips its arithmetic at `L08221`. Mana has no period-zero
  gate and faults at `L08222`, before either mana store, when its deficit arm
  is reached with period 0.
- With stable gate operands and rate 1/3, an admitted numerator cannot equal
  INT_MIN modulo 2^32: the maximum nonzero power-of-two factors total 30 for
  health and 29 for mana, below 31. Signed quotient overflow is therefore
  excluded for the admitted typed domain, although products and accumulators
  wrap.
- An injected health arm with max -32768, modifier 32668 and period -1 reaches
  divide overflow at `L08223`, but the normal positive-health and deficit
  gates forbid that maximum.
- Invalid-memory probes fail closed and identify instruction cut points, not
  original operating-system outcomes.

**Confidence.** High for the local guards, arithmetic domains and
instruction-order cuts.

**Unknown.** Original fault handling, recovery and producer reach.

### SAV-REGENORDER-531

- State dword `+54==16` exits, and the regeneration route requires signed
  health > 0.
- The scalar helper subtracts actor `+138` from server `+04` modulo 2^32; a
  signed result > 80 selects local rate 3, otherwise 1.
- Health additionally requires a deficit, a nonzero period and a signed
  full-counter remainder of 0 modulo 4.
- After its three stores, call `L08224 -> R0059` precedes mana's deficit
  test and operand reads. The local body does not recheck health or recompute
  the rate before mana.
- Mana's three stores precede its own callback at `L08225`.
- If the first callback returns normally without changing selected state, a
  later mana zero-divisor fault leaves the earlier health stores committed.
- The callback itself, the scheduler and absolute post-load chronology are
  excluded.

**Confidence.** High for the local reads, gates, store and call order, and the
explicitly conditional composition.

**Unknown.** Actual callback effects and loaded chronology.

### SAV-REGENWIRE-532

- `R0210` serializes the six pool/max/period members as two bytes each in
  `94,96,98,9a,9c,9e` order, then `a2,a3` as one byte each.
- Modifiers `de/e2` occupy bytes 10..11 and 14..15 of the earlier raw 64-byte
  block at `d4`; helper `R1570` transfers that block on store and load.
- Direct original-source anchors verify all eight scalar read/write sites and
  both modifier reads.
- Wire order differs from arithmetic store order. These raw, width-preserving
  transfers establish neither an atomic pool/remainder update nor a safe
  invented value or first-load reconstruction.

**Confidence.** High for the bounded serializer widths and the arithmetic
interpretation.

**Unknown.** Post-load first-use state and safe authored defaults.

## City transfer, sale and school derive

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-CITYMOVE-512 | Carried-to-table transfer in opcode `0x22` has a container-update path distinct from actor derive. | High / Medium / Unknown | ✔ promoted | [EXP-0282](../experiments/EXP-0282-city-sale-producers/) |
| SAV-CITYSALE-513 | Sale refreshes the actor's load, but full Human derive runs only when the stored-load quotient changes. | High / Unknown | ✔ promoted | [EXP-0282](../experiments/EXP-0282-city-sale-producers/) |
| SAV-CITYRETURN-514 | Individual return and bulk cancellation do not share one actor-refresh sequence; only individual return calls the load helper. | High / Medium | ✔ promoted | [EXP-0282](../experiments/EXP-0282-city-sale-producers/) |
| SAV-CITYDERIVE-515 | A reached sale derive can rewrite current and derived Human fields without rebuilding the retained modifier block; a successful school purchase always reaches it. | High / Unknown | ✔ promoted | [EXP-0282](../experiments/EXP-0282-city-sale-producers/) |
| SAV-CITYSTORE-516 | Selected Human SAVE fields are read from current memory, but SAVE is not globally mutation-free. | High / Medium / Unknown | ✔ promoted | [EXP-0282](../experiments/EXP-0282-city-sale-producers/) |

### SAV-CITYMOVE-512

- In opcode `0x22`, `L08226 -> R0195 -> L08227 -> R1571` resolves an
  actor in the selected Player's actor list, not a Player receiver.
- Source code 2 clears actor dword `+150`, calls `R0095` on `actor+7c`, and
  subtracts signed16 Item weight times unsigned16 detached count from container
  dword `+20`; insertion index `+1c` is retained.
- A positive partial quantity reduces the source count and uses Item `vt+40` to
  detach the requested quantity; a whole transfer removes the element.
- Destination code 4 stamps Item `+14=actor+14` and calls
  `L08228 -> L08229 -> L08230 -> R0930` on the instance tray at `+78`,
  which adds weight and may merge counts and flags.
- The dispatcher reaches `L03270`'s state sender without `R0451` or actor
  `vt+50` in this arm.
- Item `vt+50` in insertion is `R0927` stackability, not Human `R0280`.
- No primary, current, base, modifier or derived Human store occurs directly in
  this carried-transfer arm.

**Confidence.** High for the typed receivers, the container arithmetic and the
named arm's direct calls and stores. Medium for end-to-end retained Human state,
because notification callbacks are not a complete mutation trace.

**Unknown.** Ordinary UI partial-pickup reach, invalid quantities and allocation
failure.

### SAV-CITYSALE-513

- Opcode `0x34` reaches `R1572` through `L08231/L08232`; the wrapper
  writes the requested actor into `instance+74`.
- The sale loop credits its Player for owned priced Items, removes tray nodes,
  clears Item ownership and returns Items through shelf merge.
- After notifications and purse refresh, `L08233` calls `R0451(0)` on that
  actor, even if no item passed the sell predicate.
- The helper first computes `trunc(s16(actor+90)/s16(actor+92))`. It then
  rewrites own weight `+8e` with zero increment, and refreshes word `+90` from
  own weight plus truncating signed32 container `+20/2` when below 64000, or
  assigns 32000 otherwise. Last it compares the new signed16 quotient with the
  stored old quotient.
- Unequal dispatches actor `vt+50` at `L08234`; equal returns without derive.
- There is no zero-capacity guard: the first IDIV faults before the actor
  stores. Word arithmetic wraps before the new quotient.
- Twelve original-instruction vectors per identical root distinguish equal and
  unequal quotients, negative inputs, threshold, wrapping and zero divisor;
  execution stops before derive.

**Confidence.** High for the complete helper, the sale call receiver and the
conditional dispatch. No unconditional recompute or safe-default inference
follows.

**Unknown.** Live event values, legality of synthetic edge inputs and
intervening callback effects.

### SAV-CITYRETURN-514

- Code `4 -> 2` calls `R0095` through the tray take wrappers, adds to
  `actor+7c` at `L08235`, then calls `R0451(0)` at `L08210`. Human derive
  therefore uses the same stored-load quotient condition.
- Opcodes `0x35/0x36` instead call `L08236 -> L08237 -> R1573`, returning
  owned Items with `R0929`, returning unowned Items to shelves, and clearing
  the tray with raw list `R1574`. `0x36` subsequently unregisters the
  instance.
- Neither those dispatcher arms nor the complete bulk helper contains the actor
  load-helper or derive call.
- Bulk list clear and sale's raw node removal are not the container's
  weight-subtract routine; no tray-load recomputation is inferred from them.

**Confidence.** High for the distinct direct call sequences and receiver
provenance. Medium for retained actor state across notifications and instance
destruction, whose full alias effects are not closed.

### SAV-CITYDERIVE-515

- Human/Humanoid `vt+50=R0280` caps primary words; rebuilds load, speed,
  capacity, maxima, sight and live attack/defence; restores and folds class
  skills 1..5; clamps current health and mana; and writes mover speed.
- Its direct stores do not rewrite base raw `+114/24`, XP or the active skill
  byte.
- Modifier `+d4/64` is consumed rather than reconstructed from equipment. Speed
  modifier `+d8` is the conditional exception, cleared by `R0840` if folded
  speed is negative.
- The admitted opcode `0x3d` school operation charges the Player, calls
  `R0838`, restores class skills, increments the selected skill, updates the
  selected `+1cc+4*i` XP and aggregate `+130`, snapshots base slots 1..5, then
  unconditionally calls actor `vt+50` at `L08238`.
- Failed admission and merely opening the school are not this operation. Shipped
  selections are slots 1..5.

**Confidence.** High for the selected direct stores, the modifier-fold
exception, typed derive dispatch and the admitted school sequence.

**Unknown.** Complete callback effects, live before/after values, unnamed
modifier producers and any authored initial vector.

### SAV-CITYSTORE-516

- Writer `R0804` calls document `R0414`; archive `R1129` uses object
  `vt+8`, giving Human `L08212 -> R0954 -> R0210`.
- Unit passes the existing spans `+a6/24`, `+be/22`, `+114/24`, `+d4/64` to raw
  helpers whose store arms write without recomputing them. Scalar words
  `+84..+9e` and Humanoid `+1cc/24` XP are likewise emitted from their existing
  values.
- Container `R0952` writes elements, insertion index `+1c` and stored load
  `+20`, not a new weight sum.
- Unit's store arm nevertheless writes `u32 +148=u8 +14c` at `L08239`; the
  Human `+3c/+14c` repairs are load-only.
- Embedded serializers at actor `+15c/+178`, referenced-object serializers,
  writer cleanup and queued callbacks remain outside a complete alias and
  chronology proof.
- This does not establish that the next SAVE's entry state equals the
  immediately preceding sale's exit state.

**Confidence.** High for the named store primitives, source spans, container
field order and the explicit `+148` write. Medium for the bounded
writer-to-Human frontier.

**Unknown.** Universal next-SAVE normalization, embedded receiver alias effects
and whole-object authoring.

## Profile-free AppContainer child boundary

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-CHILDBOUNDARY-496 | The selected profile-free AppContainer instrument did not reach a primary child: CreateProcessW returned file-not-found with PID 0. | Unknown | ✔ promoted | [EXP-0280](../experiments/EXP-0280-child-write-boundary/) |

### SAV-CHILDBOUNDARY-496

- One ordinary-sandbox host stopped on synthetic low-label assignment.
- One sanctioned host completed five parent controls, an inheritable canary
  handle and job configuration; CreateProcessW then returned
  `The system cannot find the file specified.` with PID 0.
- No child token, child operation or second-child denial was observed.
- Exact cleanup initially failed: thirteen native descriptors gained NULL-SACL
  metadata. The DACLs and file bytes matched their originals. Backed-up
  recreation of only those thirteen generated fixtures then restored the exact
  native owner/group/DACL/label SDDL and hashes, confirmed by a separate
  read-only postcheck.
- Both failed raw records remain unchanged.
- No game copy or process, profile, registry value, or pre-existing host, parent
  or install ACL was created or changed.

**Confidence.** Unknown. The creation failure does not reject or validate
AppContainer protection, establish a missing Windows feature, or explain an
original-game behavior.

**Unknown.** Primary-child boundary feasibility and every proposed child
control.

## Human allocation and first-save tail values

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-HUMALLOC-504 | At the image-default allocator threshold, an ordinary Human request reaches the OS heap boundary, not the small-block pool. | High / Unknown | ✔ promoted | [EXP-0281](../experiments/EXP-0281-human-firstsave/) |
| SAV-HUMNEW-505 | New hero and hired Human creation share one boundary: the embedded constructors leave the live attack tail `bc/bd` unwritten and clear the modifier tail `fc/fd`. | High / Medium / Unknown | ✔ promoted | [EXP-0281](../experiments/EXP-0281-human-firstsave/) |
| SAV-HUMSEL-506 | The shipped Human weapon-definition producer does not reach the known residual-tail aliases `b6=10/42`. | High / Medium | ✔ promoted | [EXP-0281](../experiments/EXP-0281-human-firstsave/) |
| SAV-HUMNEWSAVE-507 | Static new-Human creation does not yet select a first-save tail-value model. | Medium / Unknown | ✔ promoted | [EXP-0281](../experiments/EXP-0281-human-firstsave/) |

### SAV-HUMALLOC-504

- New hero `R0823` and hired Human `R1142` request `0x1e8` through
  `R1147 -> R0747 -> R0722 -> R0748 -> R1567`.
- The last body rounds to `0x1f0`, compares against image value `0x1e0` at
  `L08240`, and reaches imported `HeapAlloc` at `L08241` with flags 0 and
  size 496.
- Original p-code reaches that boundary in 12 steps; request 480 instead reaches
  the small-block lock boundary in 8.
- Four reference-manager hits to the threshold are reads, not a proof against
  aliased runtime writes.
- Synthetic small-pool reuse controls therefore do not establish ordinary Human
  allocation residue.

**Confidence.** High for the image-default routing and boundary arguments. The
imported allocation was not executed.

**Unknown.** OS-returned payload contents, the actual runtime threshold and
first-save residue.

### SAV-HUMNEW-505

- `R0497 -> R0902 -> R0901` executes the embedded attack, defence,
  modifier and base constructors before Human definition `R0656`.
- The live attack initializer excludes `bc/bd`, whereas the modifier's actual
  64-byte memset includes `fc/fd`.
- Four isolated original-p-code sequences with sentinels 90/165 preserve
  live-tail words 23130/42405 and clear modifier-tail word 0, changing exactly
  130 actor bytes across those four embedded calls. These are not whole-Human
  constructor runs.
- Named definition skill loops, Humanoid direct stores and ordinary
  attack/defence folds do not widen that boundary.
- Unit serialization passes the live 24-byte and modifier 64-byte runs to raw
  archive writes, including both pairs.

**Confidence.** High for the shared call chain, the embedded-call coverage and
the raw-store inclusion. Medium for the bounded creation write frontier.

**Unknown.** Equipment Effect dispatch, callback and alias closure, and the
complete first-save values.

### SAV-HUMSEL-506

- Weapon equip tests signed attackType less than 10. That arm stores its low
  byte into `b6`, while every value at least 10 stores 0. Removal stores 0.
- Both lawful roots have 156 nonempty Human cell-zero weapon joins, all
  resolved, with attackType counts 1:31, 2:21, 3:60, 4:12, 5:32. The sole
  negative Weapons table row, 23, is not joined.
- These equipment inputs select only ordinary class words and resistance slots,
  not the `b6=10/42` tail aliases.
- Positive type 10/42 therefore does not construct selector 10/42; synthetic
  negative values -246/-214 would, by low-byte truncation.

**Confidence.** High for the named signed branch and store. Medium for the
finite shipped-data producer bound. No global lifecycle invariant follows:
archive restore, unchecked indexed skill writes, direct starting-weapon literals
and other callbacks are separate admission surfaces.

### SAV-HUMNEWSAVE-507

- The 67-body export retains 751 calls, 45 computed.
- Equipment reaches `R0855`, which dispatches each attached Effect through
  slot 40; its receiver population and all resulting writes are not closed here.
- Starting-skill writer `L03936` uses `a8+2*byteIndex`, permitting wider
  aliases without a local six-slot bound.
- Later events before the first archive store are not a measured empty interval.
- Both the OS-supplied-zero and the retained-nonzero live-tail models remain
  compatible; an additional alias write is also unexcluded.
- The proven modifier clear is a creation-stage value, not a universal
  first-save default.

**Confidence.** Medium for the enumerated creation-to-store frontier. No safe
native-SAV default vector is published.

**Unknown.** Complete allocation dependence, concrete first-save values, all
ordinary selector producers and zero safety.

## World-head dwords

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-WHEADINIT-520 | The eleven world-head dwords receive explicit constructor zero stores, not zeros inferred from a SAV corpus. | High / Unknown | ✔ promoted | [EXP-0283](../experiments/EXP-0283-world-head-values/) |
| SAV-WHEADLOAD-521 | The eleven world-head words are restored literally before the world-half flag, including in no-world documents. | High / Unknown | ✔ promoted | [EXP-0283](../experiments/EXP-0283-world-head-values/) |
| SAV-WHEADMAP-522 | World `+11c` and `+124` select map construction work; the block is not inert padding. | High / Unknown | ✔ promoted | [EXP-0283](../experiments/EXP-0283-world-head-values/) |
| SAV-WHEADHERO-523 | World `+128,+12c,+130,+134,+138` drive distinct actor-construction decisions; none of their actor writes mutates the world head. | High / Unknown | ✔ promoted | [EXP-0283](../experiments/EXP-0283-world-head-values/) |
| SAV-WHEADWATCH-524 | World `+148` gates a watchdog and `+13c` chooses its recovery effects; an apparent reset is not a proved live world mutation. | High / Medium / Unknown | ✔ promoted | [EXP-0283](../experiments/EXP-0283-world-head-values/) |
| SAV-WHEADLIMIT-525 | World-head construction is closed more narrowly than the first-save value frontier; in the searched population archive restore is the only nonzero producer. | High / Unknown | ✔ promoted | [EXP-0283](../experiments/EXP-0283-world-head-values/) |

### SAV-WHEADINIT-520

- `R1304` allocates `174h` bytes, calls `R0967` with that pointer as `this`,
  then publishes its return at `L00285`.
- The constructor saves `this` in the frame local at -18 and reloads it for zero stores at
  `L08242,L08243,L08244,L08245,L08246,L08247,L08248,L08249,L08250,L08251,L08252`,
  corresponding to wire-order
  `+11c,+124,+128,+12c,+130,+134,+138,+13c,+148,+144,+140`.
- The nearby clear covers only `world+[a4,118)`; initializer `R0119`'s
  `a4+4*i`, `i=1..28`, also ends before the selected block. Its mode argument
  sets `+c`, `+8` and `+150`, not a selected word.

**Confidence.** High for the exact receiver, stores and range bounds in the
identical EN/RU executable.

**Unknown.** First-SAVE invariance through every intervening callee or alias.

### SAV-WHEADLOAD-521

- `R0414` reads the eleven world addresses through `R0686` in
  `SAV-HEAD-025` order with no local clamp or Boolean normalization; the later
  difficulty clamp is separate.
- Its leading `R1339` clears global `[L06793,L08253)`, not the world
  object.
- `R0512` with `+154 != 0` loads `game0000.sav`. `R1284` constructs mode
  2, restores the archive and can enter `R0099(1)`, whose argument skips
  new-mission construction and whose `+6b8` gate can run a tick.
- Literal restore is not evidence that post-load calls preserve or recompute
  every selected word.

**Confidence.** High for the straight head load and the named branch structure.

**Unknown.** Transitive post-load mutations, the first runtime consumer and
resave fidelity.

### SAV-WHEADMAP-522

- During `R0128`'s call to `R0151`, `L08254` reads `[L00285]+11c`;
  nonzero suppresses a map unit unless its resolved Player has `+28 == 0`.
- In ordinary `R0512` map construction, `L08255` calls the `Outposts`
  reader `R0474` on `world+44` only when `+124 == 0`.
- The later `L08256` calls `R1575` on that same subobject only when
  `+11c == 0`. `R1575` reads `Mission/Players` with a `Monsters` fallback.
- The `+11c` map-unit read can therefore precede the later mission-entry test.

**Confidence.** High for the named receivers and branch effects.

**Unknown.** Ordinary producers of nonzero values and all-path first-consumer
chronology.

### SAV-WHEADHERO-523

- `R0823` saves the world receiver separately from its actor. After the existing-hero
  arm has returned:
  - `+130 != 0` writes actor word `+8c = 40h`;
  - either `+12c != 0` or `+134 != 0` calls `R1562(actor)`;
  - `+138 != 0` constructs `PlasmaSword` and calls actor `vt+3c` with it.
- `R1562` writes six words 100 at actor `+102..+10c` and six bytes 100 at
  `+10e..+113`, then invokes `vt+50`.
- In `R0065`, `+134` independently selects that helper on an eligible
  reconstructed actor.
- `+128 == 0` calls session `R0044(actor)`; nonzero calls session
  `R0308(actor,Player+34,0)`, the published Defend operation, not Follow.
- None of those actor writes is a world-head mutation.

**Confidence.** High for the conditional instructions, receiver separation and
helper payload.

**Unknown.** Broad gameplay labels, computed-call effects, ordinary nonzero
producers and runtime outcomes.

### SAV-WHEADWATCH-524

- `R1576` calls `R1290` only when `+148 != 0`; that body tests the
  full-clock signed remainder by 5.
- With a selected actor, nonzero `+13c` and a count below two in the `Player+20`
  group's `+4` actor list permit actor word `+c0 = 0`.
- Without that actor but with a Player, `+13c` enables the global `L08257`
  countdown: a nonzero value is decremented; an invocation finding zero reloads
  2 and calls `R0065`.
- Separately, `L08258` clears incoming `this+148`, conditionally clears
  `this+13c`, then may subtract 25 from the selected actor's `+c0`.
- The reference-manager call/branch and aligned `.rdata` pointer query finds no
  caller for that routine, leaving even its world-compatible receiver unbound.

**Confidence.** High for the watchdog's conditional instructions and the
candidate routine's body. Medium for the candidate's world-layout
interpretation.

**Unknown.** The candidate's exact receiver and reachability, real-time cadence
and ordinary nonzero producers.

### SAV-WHEADLIMIT-525

- The finite `.text` raw-dword/listing census, global-reference queries and 30
  named bodies establish no nonzero producer other than literal archive restore,
  and identify no non-serialization world consumer for `+140` or `+144`.
- Those are search results, not global absence: undecoded locations, same-offset
  non-world receivers, data-driven offsets, interior-pointer aliases, computed
  dispatch and unexpanded callees remain.
- `R0826` has no selected-field direct store in its body; its `L08259` byte
  `+13c` clear is on an actor.
- Neither that body nor a constructor zero proves invariance across a mission
  boundary.
- A zero-until-restore model and an untraced-alias producer model both fit the
  evidence.

**Confidence.** High for the explicit scope and the actor/world distinction.

**Unknown.** All-path first-SAVE values, nonzero producer reach, complete
mutation and post-load coverage, world `+140/+144` semantics and original
load/action/resave compatibility.

## Original-process startup probes

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-INITGUARD-488 | The tested restricted-token profile does not safely extend the original startup checkpoint. | High / Unknown | ✔ promoted | [EXP-0279](../experiments/EXP-0279-original-init-paths/) |
| SAV-HUMRUNTIME-476 | Original-process startup isolation is reachable, but loaded-Human chronology remains unobserved. | High / Unknown | ✔ promoted | [EXP-0278](../experiments/EXP-0278-human-runtime/) |

### SAV-INITGUARD-488

- One synchronous thread probe used flags 9 and one verified restricting SID,
  granted only to an owned synthetic copy subtree.
- Of nine ordered controls, outside create, write-open and rename, and
  HKCU/Software KEY_SET_VALUE intent were denied. Deleting an ordinary outside
  fixture and opening a NULL-DACL outside fixture for write succeeded.
- All file targets were generated under the same owned canary root; no ROM1 copy
  or child was created.
- ACL readback contradicted the raw restoration flag. A native DACL-only repair
  then restored the one fixture to its untouched sibling's access SDDL, with all
  nine fixture ACLs retained and only that file changed.
- Source hashes and the preserved 181-file guard passed.

**Confidence.** High for this profile's two counterexamples and the final
fixture repair. No impossibility of safe execution and no permission blocker is
claimed.

**Unknown.** Other guards, primary-child isolation, actual original resource and
save paths, menu/load reachability and Human chronology.

### SAV-HUMRUNTIME-476

- Two bounded debugger-owned EN child sessions used a 139-file byte-identical
  disposable copy with the ordinary unmodified `game0000.sav` present.
- The first stopped on a misclassified Windows `STATUS_WX86_BREAKPOINT`, not an
  identified game fault.
- The second, corrected session reached `L07681`, verified its original
  INSTALLDIR stack buffer, changed only that process-local path, and observed
  the original `SetCurrentDirectoryA` return 1 at `L08260`. It then terminated
  before the following resource-directory registration and UI construction.
- The registry value and preserved inputs stayed unchanged.
- No normal SAV load or Human observation hook was reached or armed. Zero
  captured Human events is an instrument scope statement, not a no-read census.
- Sandbox canonicalization failed, but approved escalation passed guard tests
  and preflight; this is not an approval-blocked result.

**Confidence.** High for the two bounded process records and the successful path
checkpoint. No chronological model is selected.

**Unknown.** Effective later resource and save lookup, ordinary loaded actor
identity, first selected reads and writes, admission or authoring defaults.

## Human live fields and first post-load reads

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-HUMRUN-444 | The selected Human runs cover 94 bytes in three different roles, not three opaque scalar values. | High / Unknown | ✔ promoted | [EXP-0276](../experiments/EXP-0276-human-live-fields/) |
| SAV-HUMLOAD-445 | In the traced Human load bodies, archive restore is not a modifier rebuild. | High / Unknown | ✔ promoted | [EXP-0276](../experiments/EXP-0276-human-live-fields/) |
| SAV-HUMFOLD-446 | The attack fold has seven ADD instructions and one assignment in 53 instructions. | High / Unknown | ✔ promoted | [EXP-0276](../experiments/EXP-0276-human-live-fields/) |
| SAV-HUMEQUIP-447 | Ranged to-hit authoring is history-sensitive, and its removal is not the inverse of equip. | High / Unknown | ✔ promoted | [EXP-0276](../experiments/EXP-0276-human-live-fields/), [EXP-0288](../experiments/EXP-0288-equipment-event-order/) |
| SAV-HUMMUT-448 | Health and mana amounts, their regeneration modifiers and skill progress are not interchangeable saved state. | High / Unknown | ✔ promoted | [EXP-0276](../experiments/EXP-0276-human-live-fields/) |
| SAV-HUMGAPS-449 | Known consumers do not close all live Human authoring inputs. | High / Medium | ✔ promoted | [EXP-0276](../experiments/EXP-0276-human-live-fields/) |
| SAV-HUMRESUME-460 | Original resume and the expanded Human reference hook do not establish a universal derive-before-read boundary. | High / Unknown | ✔ promoted (amended, superseded) | [EXP-0277](../experiments/EXP-0277-human-postload/) |
| SAV-HUMPROJECT-461 | Actor-state projection computes from live Human damage bytes without a local derive. | High / Unknown | ✔ promoted | [EXP-0277](../experiments/EXP-0277-human-postload/) |
| SAV-HUMTICK-462 | First gameplay reads are branch-dependent: regeneration and Effect dispatch consume Human modifier state, and not every Effect derives. | High / Unknown | ✔ promoted | [EXP-0277](../experiments/EXP-0277-human-postload/) |
| SAV-HUMSTRIKE-463 | The admitted ordinary melee strike reads live selected Human fields without a derive in its dispatch prefix. | High / Unknown | ✔ promoted | [EXP-0277](../experiments/EXP-0277-human-postload/) |
| SAV-HUMINDEX-464 | A malformed active index can make residual Human fields computed inputs before derive's later clears. | High / Unknown | ✔ promoted | [EXP-0277](../experiments/EXP-0277-human-postload/) |
| SAV-HUMFIRST-465 | The post-load Human authoring frontier remains conditional despite full selected-byte accounting. | High / Medium / Unknown | ✔ promoted | [EXP-0277](../experiments/EXP-0277-human-postload/) |

### SAV-HUMRUN-444

- The raw-a6 tail is live `+0xb4..+0xbd` (10 bytes); the raw-be tail is live
  `+0xc0..+0xd3` (20); raw-d4 is `+0xd4..+0x113` (64), a modifier prefix plus
  attack and defence sub-blocks at `+0xe6` and `+0xfe`.
- Thirty-four field rows account for every selected byte without overlap.
- The two attack tails include unnamed final bytes, so the full serialized
  extent exceeds the named arithmetic payload.
- Current health, current mana and six earned-XP dwords are outside these
  ranges.

**Confidence.** High for serializer addresses, widths, complete byte accounting
and the named producer and consumer mappings.

**Unknown.** Meanings and producers marked Unknown in the field map remain
Unknown.

### SAV-HUMLOAD-445

- `R0210` passes the live and raw pointers through symmetric 24/22/64-byte
  helpers before its scalar load arm.
- `R0954` restores XP and equipment references, and `L08212` then restores
  class-definition state.
- Human hook `R1569` only delegates to Unit hook `R1482`, whose named
  calls repair references and mover/order state.
- None of these bodies directly invokes `R0280` or reconstructs
  `+0xd4..+0x113` from equipped references.
- Later derive reads the retained modifier block, including active `+0xb6` as an
  input rather than deriving it from the weapon reference.

**Confidence.** High for the direct-body and callsite ordering. This is not a
global no-rebuild claim.

**Unknown.** The first computed post-load consumer, nested callback reach,
city/world interleaving and runtime survival.

### SAV-HUMFOLD-446

- `R0846` adds u16 to-hit and six u8 damage values, assigns elemental kind at
  relative `+0x15`, and skips all six skill words, the active-index mirror
  `+0x10` and the final two bytes.
- With source `actor+0xe6`, it therefore does not restore active live `+0xb6`
  from `+0xf6`.
- The defence fold adds eight words and six bytes from `+0xfe` into `+0xbe`.
- A 64-byte modifier basis and a simultaneous wrap/assignment probe, executed
  from each lawful image, distinguish addition from copy, saturation and type
  addition: `250+10` wraps to 4, kind 3 becomes source kind 2, and active 5
  stays 5 despite mirror 1.
- `HERO-FOLD-035`'s eight-ADD/21-instruction counts are false.

**Confidence.** High for complete instruction bodies, exact widths and bounded
instruction-interpreter results, independently checked against 200 randomized
scalar-oracle vectors.

**Unknown.** Original-process execution and surrounding recompute effects.

### SAV-HUMEQUIP-447

- `R0850` melee adds item `+0x52` into actor `+0xe6`, but kinds 11/12 assign
  General `+0xa8` into `+0xe6`. The common removal body `R0853` subtracts
  item `+0x52` in all arms.
- Actual instruction slices with prior modifier 5, General 17 and item to-hit 11
  give melee 16 then 5, but ranged 17 then 6.
- Melee writes active `+0xb6` from the low byte of signed definition parameter 5
  cached in the frame local at -4 (`L04005..L04006`, `L04007/L04008`); ranged 11/12
  writes zero.
- Removal assigns zero at `L04199`; it does not restore a previous selector.
- Armor and Shield explicitly add their defence blocks into both live and
  modifier copies. Weapon instead invokes derive after the shown modifier
  stores.
- These paths refute the universal inverse and direct-dual-write clauses of
  `HERO-EQUIP-017`.

**Confidence.** High for the complete named writer paths and three controlled
instruction-slice outcomes. The synthetic values are discriminators, not SAV
defaults.

**Unknown.** Full item-effect/eviction event interleaving and runtime equipment
cycles.

### SAV-HUMMUT-448

- Effect dispatch `R0843` separates current amount writes from maximum
  modifiers `+0xdc/+0xe0` and rate modifiers `+0xde/+0xe2`.
- The admitted regeneration arithmetic reads the rate modifiers without writing
  the selected 94 bytes. With current 10, maximum/period 100 and rate 1,
  modifiers 0/50/-100 yield health 12/13/10 and mana 11/11.50/10.
- The `.50` is remainder byte `+0xa3=50`, not a floating saved mana value.
- Award `R0810` writes XP and aggregate outside these runs and invokes derive
  on a level raise.
- The active-skill derive slice with skill 9 versus 10, initial damage 2 and
  to-hit 20 produces damage 3/4 and to-hit 47/50; active zero produces 2/20.

**Confidence.** High for scoped instruction arithmetic and producer separation.
Interpreter windows omit tick gates, network callbacks and the rest of derive.

**Unknown.** Control admission, complete skill-award events and GUI/runtime
deltas.

### SAV-HUMGAPS-449

- `R1577` initializes only 22 of the serialized attack block's 24 bytes: a
  24-byte sentinel input leaves the final two unchanged.
- The modifier constructor `R0844` separately clears all 64 bytes, including
  its mirrored tail.
- No field-level use of live `+0xbc/+0xbd` or mirrored `+0xfc/+0xfd` was
  established in the 34 selected bodies.
- Capacity modifier `+0xda` and second-component modifiers `+0xf7/+0xf8` have
  known fold consumers but no nonzero producer in those bodies.
- `+0xe8` has effect writers but is skipped by the traced skill loops, and
  `+0xf6` is skipped by the attack fold.

**Confidence.** High for initializer coverage and the bounded named-body
distinctions. Medium for the inferred unused-tail and shadow interpretation,
because aliases, bulk copies and later routes remain live alternatives. No
zero-safety or whole-object authoring claim follows.

**Unknown.** The first loaded computed consumer and safe values for these
residual fields.

### SAV-HUMRESUME-460

- `R1284` loads the archive and reads YA1 InBattle through `R1578`.
  Nonzero selects `R0099(1)`; zero selects city restoration, with
  queued-command drains on both sides.
- World resume skips fresh construction. Its bootstrap `R0147` call requires
  frontend `+6bc != 3`, nonzero authority `+6b8`, then the callee's server `+2c`
  gate.
- Human slot `+24 = R1569` calls Unit `R1482` and Token `R1560`. The
  expanded Position, reference, mover and order key-repair helpers touch none of
  the selected 94 Human bytes with ordinary disjoint receivers, and invoke no
  Human slot `+50`.
- Stage `+13c != 0` skips the mover and order repairs.
- With stage zero, `L04566` calls order helper `R0170` on `actor+158`.
  Slots `+0c,+10,+18,+20,+28,+30,+68` each skip zero, look up a nonzero key
  through `R1488` in `[L00285]+88`, and write the mapped value only on a
  hit.
- A miss preserves the raw dword: for `+20`, `L07753` skips store `L07754`
  after lookup `L07752` returns zero. `R1488` returns zero without writing
  its output on a miss (`L07755`).
- Earlier document calls, queues, client/message dispatch and malformed pointer
  aliases are outside that negative scope.

**Confidence.** High for the named mode and call edges and the expanded
reference-hook closure.

**Unknown.** The absolute first loaded computed consumer and the global
derive-before-UI order.

**Amended.** `retracted.md` makes the reference-repair shorthand field-specific
and withdraws the EXP-0277 report sentence that missing entries become null; it
is not this order helper's law. The no-derive, retained-input and first-consumer
Unknown boundaries stand.

### SAV-HUMPROJECT-461

- `R0059` uses actor slot `+30`, not derive `+50`.
- For a Human, non-owner recipients outside type IDs `21..3f` mask out damage
  bit `200`. With that effective bit, packet bytes are `(b4+b9+b7) mod 256` and
  `(b5+ba+b8) mod 256`, narrowed by original helper `R0881`.
- Nine controls distinguish each input, wrap and mask-off; none writes selected
  Human fields.
- Entry `R0131` calls the sender at `L03739`.
- Phase-12 `R0860` calls `R0811` before actor slot `+14`; `R0675` then
  requires two nonzero helper results and nonzero actor slot `+2c` to request
  all-state projection.
- The `R0808` derive is on a freshly constructed temporary chargen Human, not
  a loaded actor.

**Confidence.** High for conditional computed presentation, arithmetic and
same-iteration ordering. Projection is not itself a combat consumer.

**Unknown.** Earlier derive and callback order, recipient population or gameplay
damage.

### SAV-HUMTICK-462

- Server `R0147` runs phase-12 full-tick work before ordinary `R0193`,
  which drains commands and world-object dispatch before global actor subticks.
- Human full-tick `R0654` excludes state 16. Positive health plus health
  deficit/nonzero period/full-tick modulo 4 admit signed `de` at `L08261`;
  positive health/mana deficit admit signed `e2` at `L08262`. Neither
  regeneration branch writes selected bytes.
- Actor subtick `R0037` calls attached Effect slot `+38` before health,
  orders and actions.
- Base Effect `R0673` requires duration/continuous bits, applies continuous
  effects every remaining-duration multiple of 8 and removes expired
  noncontinuous effects.
- Generic `R0843` first reads signed Human `d8`, clears above 24, then
  eventually invokes derive at `L06977`.
- Special identity `+0c=8` apply instead reads protection `c6`. Identity 17 on
  Human apply/remove updates `e4` and `a4` by signed magnitude times 256 modulo
  65536 without generic derive.

**Confidence.** High for named branch gates, call order and four
special-identity integer controls. Identity 8 x87 scaling was not executed.

**Unknown.** The first attached Effect and receiver, other subclasses and actual
resume timing.

### SAV-HUMSTRIKE-463

- Melee action `L00006 -> R0001 -> R0246` requires the target-class arm,
  actors and owners, positive attacker health and distance within reach.
  Defender slot `+4c = R0265` receives attacker `+a6`.
- Positive physical damage with nonspecial flags reads signed absorption `c0`,
  then the resistance byte at `target+ce+attacker.b6`.
- Secondary `b7/b8` has a separate nonzero-sum gate, independent of the physical
  hit, and reads `c6`.
- Admitted third damage switches `bb=1..5` to `c4,c6,c8,ca,cc`; `0,6,255` reach
  diagnostic `L04135`, not an ordinary slot-zero protection arm.
- Four absorption, four secondary-admission, two explicit post-RNG and eight
  selector controls retain exact integer boundaries. RNG, x87 and full-strike
  outcomes are not emulated.

**Confidence.** High for the conditional consumers, admission gates and bounded
controls.

**Unknown.** Any earlier effect or command derive, the diagnostic outcome and
whole-game slot-zero use.

### SAV-HUMINDEX-464

- `R0280` accepts any nonzero byte `b6` in the selected indexing step,
  reading signed `word[a8+2*b6]` at `L08263/L08264` before clearing `b7..ba`
  and the `be..d3` block.
- Synthetic indices 10, 13, 32, 39, 42 select `bc/bd`, `c2/c3`, `e8/e9`,
  `f6/f7`, `fc/fd`. Word 10 changes to-hit 20/base 2 to 50/4.
- Index 0 skips; 1/5 controls use class slots.
- Index 160 crosses the instrument's 0x1e8 allocation boundary; it is not an
  observed process fault.
- Separately, resolver `L04147` reads `target+ce+attacker.b6` without a
  six-slot local bound; twelve synthetic indices cover ordinary and
  residual-byte aliases.

**Confidence.** High for positive bounded address and arithmetic sensitivity and
the overwrite order. Late-window state is supplied, not attributed to a whole
derive execution.

**Unknown.** Ordinary producer reach, accepted malformed saves and safe
authoring values.

### SAV-HUMFIRST-465

- Thirty-four rows cover 94 bytes and distinguish replace-on-derive live fields,
  retained modifiers and surviving tails.
- Positive malformed-index aliases consume unnamed tails, the General shadow and
  the active mirror. This narrows the earlier bounded Unknown without
  establishing their ordinary meaning.
- Capacity `da` has a positive retained fold input (301 plus 0/7/-7 gives
  301/308/294), but no new nonzero native producer. The same producer gap
  remains for secondary modifiers `f7/f8`.
- Slot-zero protection folds but has no identified ordinary arm in the examined
  resolver.
- Both earlier-callback-derived and saved-live-field-first models remain
  compatible with the static frontier.
- Queue, receiver, Effect and counter observation is the next discriminator.

**Confidence.** High for the positive capacity slice. Medium for the bounded
load-to-consumer frontier. No zero-safety inference follows, and no byte-roundtrip inference.

**Unknown.** Absolute first chronology, residual producer invariants, admission
and whole-object authoring.

## Container header and coded body

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-HDR-001 | A save opens with magic `41 73 67 26` (`Asg&`) and little-endian u32 fields; the 20-byte header and `0x14` body start are superseded by `SAV-FRAME-021`'s 16-byte header. | High / Medium | ● active (amended, superseded) | [EXP-0009](../experiments/EXP-0009-sav-container/), [EXP-0046](../experiments/EXP-0046-sav-body/) |
| SAV-VER-002 | `u32@0x08 = 0x0BAD0002` in all 4 saves is a `rom.exe`-held constant paired with the magic at two sites; the reader compares it as a version (`SAV-FRAME-021`). | High / Medium | ● active (amended, partially retracted) | [EXP-0009](../experiments/EXP-0009-sav-container/), [EXP-0046](../experiments/EXP-0046-sav-body/) |
| SAV-PTR-003 | `u32@0x04` is the coded body's end and `u32@0x0C = u32@0x04 − 16` (4/4); `SAV-FRAME-021` reverses which of the two has a consumer. | High / Medium | ● active (amended) | [EXP-0009](../experiments/EXP-0009-sav-container/), [EXP-0046](../experiments/EXP-0046-sav-body/) |
| SAV-EMB-004 | Each save embeds one REG-style inline key-value `&YA1` store at `u32@0x04 + 0x100`, outside the coded body, holding 9 top-level state keys; no `M7R` snapshot. | Medium | ● active (amended, partially retracted) | [EXP-0009](../experiments/EXP-0009-sav-container/), [EXP-0046](../experiments/EXP-0046-sav-body/), [EXP-0150](../experiments/EXP-0150-save-fog-record/), [EXP-0223](../experiments/EXP-0223-sav-campaign-tail/) |
| SAV-MAP-005 | The save references its map by a length-prefixed ASCII name at decoded body `+0x08` (`06` then `"10.alm"`, 4/4); the map is referenced, not embedded. | High / Medium | ● active (amended) | [EXP-0009](../experiments/EXP-0009-sav-container/), [EXP-0046](../experiments/EXP-0046-sav-body/) |
| SAV-UNK-006 | `u32@0x10` exceeds the file size in every sample because it is the coded body's decompressed size in 16-bit words; it is the blob's first dword. | High | ● active (amended) | [EXP-0009](../experiments/EXP-0009-sav-container/), [EXP-0046](../experiments/EXP-0046-sav-body/) |
| SAV-PACK-007 | The coded body is a run/literal code over 16-bit words: opcode `n < 0x80` precedes `n` literal words, and `n ≥ 0x80` repeats the next word `n & 0x7f` times. | High / Medium / Unknown | ● active (amended) | [EXP-0046](../experiments/EXP-0046-sav-body/), [EXP-0048](../experiments/EXP-0048-sav-stream/) |
| SAV-SIZE-008 | `u32@0x10` equals the decompressed size of the coded body in 16-bit words, exactly on 4/4 saves; the dword lives inside the blob, not the container header. | High | ● active (amended) | [EXP-0046](../experiments/EXP-0046-sav-body/) |
| SAV-EXT-009 | The file is three regions: header, coded body ending at `u32@0x04`, and a verbatim tail to EOF; the coded region starts at `0x10`, not `0x14`. | High | ● active (amended) | [EXP-0046](../experiments/EXP-0046-sav-body/) |

### SAV-HDR-001

- Corpus model, 4/4 saves: magic `41 73 67 26` ("Asg&") @0x00, then LE `u32` at
  `0x04`,`0x08`,`0x0C`,`0x10`; the body begins @0x14 after a fixed 20-byte
  header.
- EXP-0046 gives three of the four fields a demonstrated role: `0x04` is the end
  of the coded body, `0x10` its decompressed word count, and `0x08` a constant
  `rom.exe` itself holds.
- The decode pins the body start: it fails from `0x11` and from `0x15..0x18`.

**Confidence.** High for the field layout and the body start. No save reader was
read for this grade, which was restored after the 2026-07-25 demotion. What
carries it is that a decode begun at `0x14` reproduces `u32@0x10` and `u32@0x04`
exactly in 4/4 saves and fails at every neighbouring start. The bound is
`0x12..0x14`, since `0x12`/`0x13` are that same field's zero high bytes and
opcode `0x00` is a no-op. Medium for `0x0C`: it is only ever `0x04 − 16`,
EXP-0046 showed it bounds nothing, and no consumer was known.

**Amended.** `SAV-FRAME-021` (EXP-0077, from the writer and the reader)
supersedes the framing: the header is 16 bytes, the blob begins at `0x10`, and
`0x10` is not a header field but the blob's own decompressed-word count
(`retracted.md`). The High was earned against the wrong map; read
`SAV-FRAME-021` first.

### SAV-VER-002

- `u32@0x08 = 0x0BAD0002`, identical across all 4 saves.
- The EXP-0046 whole-image byte scan finds `02 00 ad 0b` at 2 sites in `rom.exe`
  (file `0x0ce4fc`, `0x0ce726`), each `0x2a` bytes past one of that binary's
  four `"Asg&"` literals, and at 0 sites in `Map Editor.exe`, which also holds
  no `"Asg&"`.
- The value is a constant the writer holds, not something the map supplied.

**Confidence.** High that it is a `rom.exe`-held constant paired with the magic
at two sites: a whole-image byte scan cannot be evaded by any code path, the
same instrument `TERR-PASS-053`'s `Pass*` absence used. Medium, as graded here,
that it is a version: nothing observed here compared it for range or ordering,
and one shipped value cannot show a second.

**Amended.** `SAV-FRAME-021` lifts the Medium: the reader compares it for
ordering, `CMP …,0xbad0002` / `JGE` at `L07884`, failing to "Outdated save
file." `retracted.md` withdraws the not-a-version clause.

### SAV-PTR-003

- `u32@0x04` is a byte offset into the file, and `u32@0x0C = u32@0x04 − 16` in
  4/4 saves.
- The EXP-0000 "2 varying bytes" (`0x04..0x05`) are the low half of `0x04` and
  encode a position that scales with body size, not a version, slot id or
  counter.
- `u32@0x04` is where the coded body ends and the uncompressed tail begins.
- `u32@0x0C` is not the start of a "descriptor head": the coded body runs
  straight through it, and the 16 bytes there, with EXP-0009's `E1 AC DF BA`
  marker, are the body's own last opcodes. Stopping the decode at `u32@0x0C`
  leaves 209 words undelivered in every save.

**Confidence.** High for `u32@0x04` as the coded body's end: the decode
terminates there and nowhere else, 4/4, and the same decode independently
reproduces `u32@0x10`. Medium for the `−16` relation: measured 4/4, but with no
consumer found for `0x0C` it may be derived from `0x04` rather than stored for a
reader.

**Amended.** EXP-0046 withdrew the "descriptor head" (`retracted.md`).
`SAV-FRAME-021` reverses the no-consumer clause: `u32@0x0C` is the field with a
consumer, the reader's allocation size and read length, and `u32@0x04` is read
and discarded; the `−16` relation is the writer's own arithmetic
(`retracted.md`).

### SAV-EMB-004

- The nested `&YA1` sub-container begins at `u32@0x04 + 0x100`. It is the
  REG-style inline key-value `&YA1`, which the `.res` tail-registry parser
  rejects.
- It holds 9 top-level state keys
  (`Character CurrentState Fog GameOptions Inventory Objects Projectiles SpellBook View`;
  header `@8 = 9`).
- No `M7R` map snapshot is embedded.
- It lies after `u32@0x04`, outside the coded body, so the `0x100`-byte label
  region and the `&YA1` are the only part of the file stored verbatim.
- The `Fog` key is the explored-terrain record (`SAV-FOG-061`).
- The between-mission save carries 7 of the 9 sections, lacking `Fog` and
  `Projectiles`.

**Confidence.** Medium: 4 saves, one map, no reader read. The key names stand;
the extent was wrong.

**Amended.** EXP-0150 (`SAV-TAILEXT-062`) corrects the extent "running to EOF"
(`retracted.md`): the store ends at `0x18 + 32R + 4 + poolLen`, and a further
268 bytes follow it (310 on the between-mission save), so the tail is three
regions, not two. EXP-0223 identifies what follows the store: one campaign
record rooted at application `+0x548` (`SAV-CAMPTAIL-070`), not another `&YA1`
extent.

### SAV-MAP-005

- The name sits at `+0x08` of the decoded body: one length byte, then that many
  ASCII bytes.
- It resolves against the install (`scenario.res:10.alm`, 80×80) in 4/4.
- A literal run of the coded stream carries it intact, so EXP-0009 saw it "near
  the start" of the raw file.

**Confidence.** High: the offset is fixed 4/4 in the decoded stream, and the
name resolves to a shipped map whose dimensions the rest of this ledger's
figures then reproduce. Medium that `+0x00`/`+0x04` before it are what they look
like (see `SAV-STREAM-010`).

**Amended.** EXP-0046 gave the name its fixed home at decoded `+0x08`.
`SAV-HEAD-025` extends it: the field is `world+0x28`, and the load arm re-opens
it as `Scenario\` + name whenever the mission number is non-zero.

### SAV-UNK-006

- `u32@0x10` exceeds the file size in every sample: 28934–29157 against sizes
  25596–27935; high byte `0x71`.
- It is the decompressed size of the coded body, in 16-bit words. It exceeds the
  file because the body is compressed; `2 × u32@0x10` is the byte length the
  decoder emits.

**Confidence.** High, resolved; see `SAV-SIZE-008`.

**Amended.** EXP-0046 resolves EXP-0009's "not a file size/offset, meaning
unknown". `SAV-FRAME-021` relocates the field: the value is right and the
address is not; it is the blob's own first dword, inside the coded region
(`retracted.md`).

### SAV-PACK-007

- The body is not stored verbatim: `[0x14, u32@0x04)` is a run/literal code over
  16-bit words.
- One opcode byte `n`: `n < 0x80` → `n` literal words (`2n` bytes) follow;
  `n ≥ 0x80` → the next word is repeated `n & 0x7f` times.
- No other opcode form and no end marker: the stream ends when the span does.
- Observed operand domains over the four saves: literal counts 1..126, run
  counts 2..127. `0x00`, `0x7f`, `0x80` and `0x81` never occur. A run of 1 word
  costs the same 3 bytes as a literal of 1, which is consistent with an encoder
  that never emits one.
- Per save ≈ 1520 literals and ≈ 1600 runs.
- EXP-0048 census: over 12 516 opcode positions (3093/3132/3105/3186 per save)
  the four values occur 0 times; 82 distinct opcode values occur at all. This
  measured absence bounds what shipped saves exercise and says nothing about how
  the decoder treats those values.

**Confidence.** High for the decode of these four bodies. It is fixed by the
opcode alone, yet reproduces two quantities it is not given, `u32@0x10` and the
terminal offset `u32@0x04`, exactly, 8/8. Four named rivals score 0/4 each:
literal count in bytes; run repeats a byte; counts biased `n+1`; no run bit at
all. The run unit is a word rather than a byte because 145–150 runs per save
repeat a word whose two bytes differ. A byte-swapped-word variant, which keeps
the length, is excluded by content: it spells the map name `01a.ml`. Nothing in
`rom.exe` was read for this grade, so no instruction carries it. Medium that it
is the format's code rather than these files': no shipped save reaches a body of
65 535 words, so a wider count form would not have been seen. The discriminator
named was the codec routine in `rom.exe`.

**Unknown.** The behaviour of the four unobserved opcode values, as graded here.

**Amended.** `SAV-CODEC-022` confirms the decode instruction for instruction and
closes its Unknown: all four unobserved opcodes are legal, and the compressor's
output buffer is sized to its input (`retracted.md`). `SAV-FRAME-021` moves the
coded region's start to `0x10`.

### SAV-SIZE-008

- 29085 / 28988 / 29157 / 28934 against decoded word counts 29085 / 28988 /
  29157 / 28934: exact, 4/4, so `2 × u32@0x10` bytes are emitted.
- Resolves `SAV-UNK-006`. The "high byte `0x71` in all four" EXP-0009 noted is
  the shared magnitude of four saves of one map.

**Confidence.** High. An exact identity on four files; the alternative readings
once listed, an in-memory buffer size and a session serial, predict no such
identity. A reader needs this number to size its output buffer, which is what a
decompressed length is for.

**Amended.** `SAV-FRAME-021` and `SAV-CODEC-022` relocate the field: the
identity holds, and the dword lives inside the blob, written and read by the
codec, not by the container (`retracted.md`).

### SAV-EXT-009

- The coded body is `[0x14, u32@0x04)`; everything from `u32@0x04` to EOF is
  stored verbatim.
- The decoder consumes to exactly `u32@0x04` in 4/4 saves. Started anywhere in
  `0x11` or `0x15..0x18`, it fails both predictions.
- The file is therefore three regions, not the four EXP-0009 drew: header
  `[0, 0x14)`, coded body `[0x14, u32@0x04)`, uncompressed tail
  `[u32@0x04, EOF)` = a `0x100`-byte label region then the embedded `&YA1`.

**Confidence.** High. The terminal offset is reproduced 4/4 by a decode that is
not given it, and the start is excluded on both sides. The bound on the start is
`0x12..0x14`, because `0x12`/`0x13` are `u32@0x10`'s zero high bytes and opcode
`0x00` is a literal of no words.

**Amended.** `SAV-FRAME-021` corrects the start: the coded region is
`[0x10, u32@0x04)`. The three-region split and the terminal offset stand; only
the start moves, by the four bytes the codec reads as its own header
(`retracted.md`).

## Decoded stream: head, block array, objects and tail

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-STREAM-010 | The decoded body opens with two counters where `+0x04 == (+0x00) >> 4` (4/4) and the map name at `+0x08`, and carries `CArchive` new-class records naming 11 classes. | High / Medium / Unknown | ● active (amended) | [EXP-0046](../experiments/EXP-0046-sav-body/), [EXP-0048](../experiments/EXP-0048-sav-stream/) |
| SAV-BLOCK-011 | The decoded body carries `TERR-PASS-053`'s predicted block array: a `u16` count, then u32 records packing cell, dynamic and static bytes inside `R1360`'s window. | High / Medium / Unknown | ● active (amended) | [EXP-0046](../experiments/EXP-0046-sav-body/) |
| SAV-BLOCK-012 | The block array is a delta over the terrain ingest, not a plane: a cell whose only block bits are the map's own `0x01` or `0x05` is never written. | High / Medium / Low | ● active | [EXP-0046](../experiments/EXP-0046-sav-body/) |
| SAV-STREAM-013 | The decoded stream is one object graph under one index counter shared by classes and objects, from 1 in stream order; the walk closes on 393/393 tags in four saves. | High / Medium | ● active (amended, superseded) | [EXP-0048](../experiments/EXP-0048-sav-stream/), [EXP-0251](../experiments/EXP-0251-sav-full-reader/) |
| SAV-OBJ-014 | The corpus's two equal cell words are one packed cell stored twice inside the raw position object (413/413 instances), under `SAV-TOKENPOS-074`'s layout. | High / Medium | ● active (amended) | [EXP-0048](../experiments/EXP-0048-sav-stream/), [EXP-0145](../experiments/EXP-0145-save-members/), [EXP-0224](../experiments/EXP-0224-sav-token-position/) |
| SAV-ID-015 | The head's `+0x0c` id is a runtime id from the lowest free bit of bitmap `L06793`, not the map's id field; on these saves hero 1, buildings 2..19, units 20..54. | High / Medium | ● active (amended) | [EXP-0048](../experiments/EXP-0048-sav-stream/), [EXP-0055](../experiments/EXP-0055-tick-order/) |
| SAV-OBJ-016 | The saved population is the map's, object for object: 5 Players for the 5 type-5 slots, one Human or Unit per type-6 record plus the hero, 18 Buildings, items and sacks. | High / Medium | ● active (amended) | [EXP-0048](../experiments/EXP-0048-sav-stream/), [EXP-0055](../experiments/EXP-0055-tick-order/) |
| SAV-CELLREC-017 | The block array is followed by a counted table, `u16 count` + `count ×` (`u16 packedCell` + 52-byte payload), serialized by `R1521`. | High / Medium | ● active (amended) | [EXP-0048](../experiments/EXP-0048-sav-stream/), [EXP-0236](../experiments/EXP-0236-sav-cell-record-rebind/) |
| SAV-TAIL-018 | A fixed-size region sits between the cell-record table and the Sack records in all four saves, and the stream ends in zeros; `SAV-SESS-031` names the region at 4374 bytes. | High / Unknown | ● active (amended, superseded) | [EXP-0048](../experiments/EXP-0048-sav-stream/) |

### SAV-STREAM-010

- `+0x00 u32` and `+0x04 u32` are two counters satisfying
  `+0x04 == (+0x00) >> 4` exactly in 4/4: 806/50, 2992/187, 8533/533, and 1/0 in
  the save labelled "Restart last mission".
- `+0x08` is the map name (`SAV-MAP-005`).
- After the name come zeros `[0x0f, 0x3b)`, then four u32 `10 2 6 5`, identical
  in all four saves. The last equals the map's type-5 player count; one map
  cannot separate that from coincidence.
- Further in, the stream carries 11 records of the shape
  `FF FF | u16 schema = 1 | u16 len | len ASCII bytes` naming
  `Player Human Weapon Item Armor Diary Effect Shield Unit Building Sack`. This
  is the standard Microsoft `CArchive` new-class record, which makes the body's
  object graph readable without a disassembler.
- The class records appear in first-use order, which differs between saves:
  game0002 registers `Item`/`Effect` 3rd/4th, the restart save 4th/7th. The
  record order is state; the name set is not. All 44 records carry `schema = 1`.
- No u32 in the Player object or the head tracks the elapsed counter's delta
  over the 6 save pairs (hypothesis H6).

**Confidence.** High for the `>>4` identity, the 11 records, the four head u32s
and the per-save orders, which are exact measurements. Medium that `+0x00` is
elapsed time: monotone with save order and 1 on the restart save, suggestive but
not a reading.

**Unknown.** The unit of `+0x00`; the meaning of `10 2 6 5`; whether the 11
names are all the classes a save can carry, since this is one map and one party.

**Amended.** EXP-0048 mapped the rest of the fixed head. `SAV-HEAD-025` names
the four u32s as mission number, difficulty, `playerList+0x20` and player count,
and `+0x00`/`+0x04` as the sub-tick and full-tick counters `SESS-TICK-004` read
independently. `SAV-CLASS-033` amends the eleven-class count: the image carries
28 serializable classes, and `game0010.sav` adds a `Spell` record.

### SAV-BLOCK-011

- This is EXP-0041's `TERR-PASS-053` prediction, in the shape it predicted:
  `u16 count` then `count × u32` = `(cell << 16) | (dyn << 8) | static`. Cells
  are strictly increasing, every one inside `R1360`'s window
  `[0x807, 0x807+0xe5e7)`, and every `dyn > 0x0f`: 1890 / 1938 / 2068 / 1842
  records, 100 % on all three invariants.
- The window's low bound is exact to the cell. An 80×80 map has 2304 border
  cells, of which 640 (rows 0..7) plus 7 (row 8, cols 0..6) fall below `0x807`;
  the array carries all 1657 that remain, 4/4 saves.
- Byte 0 is the plane at `world+0x10000`, with byte 1 the plane at `+0x20000`, not
  the reverse. 0 records carry an occupancy bit (`0x40`/`0x80`) in byte 0,
  against 32–39 under the swap, and `TERR-PASS-051` says only the dynamic plane
  may hold one.
- Against a block plane derived from `scn:10.alm` alone, bits 0..3 agree on
  1771/1890 (1723–1949 of 1842–2068). The 119 that differ are the same 119 in
  all four saves. 89.9 % of them sit within two cells of one of the map's 18
  type-4 (building) anchors, which the terrain ingest never reads, against a
  10.2 % null over all 4096 interior in-window cells.
- External corroboration of byte 1: on the restart save the cells with `dyn` bit
  6 cover 35 of the map's 35 type-6 unit placements.

**Confidence.** High for the framing, the packing and the window: the
pack/unpack arithmetic and the sweep bounds are `TERR-PASS-053`'s named
instructions, and this claim adds that the bytes are there and that the low
bound is exact on a map whose border count is arithmetic, not fitted. Medium for
every figure that is a fact about one 80×80 map (1657, 119, the histograms), and
for reading `dyn` bit 6 as an occupant: the 35/35 hit shows it tracks units, but
`TERR-PASS-051` names it an occupancy bit, not this evidence. Medium for the 119
as building footprints: the distance evidence is strong against a stated null,
and the mechanism is not read.

**Unknown.** The 9 cells where the save clears bit 0 the map sets.

**Amended.** `SAV-1226` measures the cleared bits on two other mission-10 restart slots: 9 cells per file, each a structure footprint cell the `&0xfa` arm opens. Medium that these are the 9 cells this claim left Unknown: its four saves are not identified with those two files by hash, and the census does not count the 119 cells whose bits 0..3 differ from the ingest, which were measured on other saves of the map.

### SAV-BLOCK-012

- The serializer's `dyn > 0x0f` gate means a cell whose only block bits are the
  map's own `0x01` (terrain) or `0x05` (static object) is never written.
- Of the 4071 in-window cells whose map-derived block byte is nonzero, 2378 are
  absent (2365 in `game0002`), and a further 647 of the map's 2304 border cells
  lie below the window entirely.
- Applied to a zero plane, these records leave the map's water, mountains and
  placed objects passable and its top edge open.
- The four record sets are nested, `game9999 ⊂ game0000 ⊂ game0001 ⊂ game0002`,
  with 0 cells present only in the earlier member across all six pairs.
- What grows is one population: cells whose dynamic byte carries bit 5 while the
  static byte does not (0 / 47 / 98 / 232, none in the restart save).
  97.9–98.7 % have an 8-neighbour in the same set, against a null of 44.7–89.7 %
  built from the array's own interior cells.

**Confidence.** High for the counts, and for the sufficiency question they
answer: the array demonstrably does not determine the plane. Medium for the
consequence that a load therefore re-runs the ingest before applying these
bytes, which is `TERR-PASS-053`'s stated Unknown. This is an argument, not a
reading: the load path was not traced, and the only alternative it excludes is
"a loaded session has no terrain blocking". The discriminator is one listing of
`R1360`'s caller. Low for the growing population as movement trails: the
nesting and the restart-save zero are measured, the reading is not.

### SAV-STREAM-013

- `FF FF | u16 schema | u16 len | name` introduces a class. A u16 with bit 15
  set introduces a new object of the class whose index is the low 15 bits.
- Indices come from a single counter shared by classes and objects, starting at
  1, in stream order: a class's first instance begins where its record ends and
  takes the next index.
- The walk that enumerates the stream this way closes in all four saves: 393/393
  tagged objects carry a tag equal to `0x8000 | index-of-a-class`, 0 land
  anywhere else, and the values differ between saves exactly as the object
  counts do. Unit's tag is `0x804a` in the session-1 saves and `0x8047`/`0x8048`
  in session 2; Diary's walks from `0x800c` to `0x8014` as the hero's inventory
  grows.
- The enumeration itself is measured: map-anchored heads are found by the
  per-session `+0x08` word, Players by `0x8001`+CString, and equipped items
  (zero `+0x08`) are recovered by a deficit repair the tags themselves force.
  The hero's hidden second armor is recovered this way: Shield's observed tag
  exceeds the walk's index by exactly 1.

**Confidence.** High as the shared-index identity rule: the original 393
unprovided tags reject fixed registration ordinals and a classes-only counter,
and EXP-0251's exact call-site replay independently agrees on all 13,235 reached
archive transitions across 31 distinct documents. The null arm is directly
witnessed 8,939 times. Medium for the finite census: the plain-u16
back-reference arm remains statically present but original-produced-unwitnessed
at 0 events, and a class no measured save names may still appear.

**Amended.** EXP-0251 (`SAV-ARCHREL-253`) supersedes the "null arm unobserved"
clause of the confidence boundary (`retracted.md`); only back-reference
production remains unwitnessed. The shared counter and the 393-tag discriminator
stand.

### SAV-OBJ-014

- `SAV-TOKENPOS-074` resolves the prefix as
  `u8 cellX | u8 cellY | u16 ((cellY<<8)|cellX) | u8 subX | u8 subY | u16 unmanaged by the bounded constructor/copy family | u32 terrain pointer`.
- The word at `+0x00` is therefore the same little-endian packed cell stored
  explicitly at `+0x02`. This explains `cellA == cellB` on 413/413 instances
  without treating them as two cells.
- The restart-save binding remains 53/53, and the moving hero's `0x34 0x36`
  sub-cell pair is the corpus witness against an at-rest-only reading.
- The three former `+0x08` “state” values are raw terrain pointers; 0 on
  equipped items is a null pointer, not a format tag.

**Confidence.** High for the byte layout, the packed-cell identity, the
coordinate accessors and the terrain-pointer source: five constructors, two copy
paths, five accessors and the raw archive helper agree. Medium that no writer
outside the bounded family assigns semantics to `+0x06..+0x07`; no alias-aware
whole-image writer census was completed.

**Amended.** `SAV-TOKEN-034` narrows the old 16-byte head extent to a prefix of
the 37-byte `Token` head; the fields, offsets, packed-cell reading, 53/53 and
413/413 stand (`retracted.md`). `SAV-TOKENPOS-074` (EXP-0224) replaces the
two-cell and `+0x08` state reading.

### SAV-ID-015

- Corpus order: hero = 1, then the 18 type-4 records as 2..19 in file order,
  then the 35 type-6 records as 20..54 in file order, then runtime-created
  objects (ground sacks 55..58).
- The rival, binding by the map's own `+0x12`/`+0x40` id fields, whose value
  ranges overlap this one, resolves ids but matches 0 cells; the ordinal binding
  matches 53/53 on the restart save.
- The two id systems coexist in the record: the map's type-6 `+0x40` unit id is
  serialized at head `+0x13`, equal in 137/137 Human/Unit instances across all
  four saves. The best partial for any other tried map quantity is 109/137. This
  independently corroborates both the ordinal binding and `ALM-UNIT-048`'s
  labels.
- Dead units keep their object with id 0: game0002 lacks live objects for type-6
  ordinals 11/13/17 and instead carries three id-0 Units at other cells.
- The id is the actor's `+0x04` field, assigned at tick-list insert
  (`R0411`) as the lowest free bit of the bitmap `L06793`
  (`R0936`; id 0 pre-marked). It is sequential exactly while nothing has
  been freed, which on these saves' single session it had not.
- Corpse decay stage 5 frees the bit and zeroes the field (`R0867`,
  `L06794`/`L04100`), which is the mechanism behind the id-0 corpses; the
  next spawn then reuses the lowest freed id.
- Buildings draw from the same bitmap through direct allocator calls without a
  tick-list insert, which is how they interleave at 2..19.
- On load the id is read back and re-marked (`R0950`), so ids round-trip
  exactly (`MOVE-ID-016`).

**Confidence.** High as the identity on this corpus: the coverage is exact
(1+18+35 ids, no gaps, no extras below 55), and the `+0x13` anchor is 137/137
against a named 109/137 runner-up. EXP-0055's listings raise the id-0-corpse and
same-session-sequential readings to High. Medium as a law for saves with
mid-session churn: after any decay the sequence has reused holes, predicted by
the allocator and not yet observed in a shipped save.

**Amended.** EXP-0055 reads the mechanism: "creation-order" is the corpus's
shape, not the allocator's rule, which is lowest free bit with reuse after
decay.

### SAV-OBJ-016

- 5 Player objects = the map's 5 type-5 slots, each a CString name and its
  1-based slot id twice (`"Danath" 1`, `"Villagers" 2`, `"Rogues" 3`,
  `"Beasts" 4`, `"Nocturnal" 5`, identical 4/4).
- One Human or Unit per type-6 record, the C++ class chosen by the record's
  class key with no overlap: Human ← keys {1,7,10,11,14,24}, owners = slots 2,3;
  Unit ← keys {69,73,74}, owners = slots 4,5. Plus the hero.
- 18 Buildings = the 18 type-4 records; then items and sacks.
- Inventory objects serialize inline after their owner: a villager's Weapon
  follows its Human, and sack contents follow their Sack.
- ~100-byte untagged structures sit between some sibling units (`01 00 00 00`, a
  `col 0x80 row` triple, the packed cell); the same shape precedes each
  faction's first unit.
- Progression across the corpus: sacks 4/4/3/2 (looted), three monsters dead by
  game0002, humans drift off their spawn cells (16/14/12/12 still on them) while
  every surviving monster save one stays put.
- `Player::Serialize` is `R0415`. Its group set is written by
  `R1340` → `R1143` per group as a direct inline call, with no
  class record. The untagged ~100-byte structures between sibling units are the
  group records (group id at `+0x1c`, two serialized members, the actor-list
  count).
- Each group's `R1341` WriteObjects its actor list head→tail, which is
  why inventories and units nest under their owners.

**Confidence.** High for the counts and partitions: exact measurements closing
on map quantities not given (5, 18, 35 and the key-set split). The group-record
reading is carried by the writer's listing, `EXP-0055 evidence §7`. Medium for
"looted" and "corpses" as readings of the progression.

**Amended.** EXP-0055 reads the named discriminator, `Player::Serialize` and its
group writer, and identifies the untagged structures as group records.

### SAV-CELLREC-017

- The original four-save stride discriminator: counts 185/186/183/179, every
  tested key valid, neighbouring strides 51..57 score at most 14, and the key
  set equals the static-bit-5 block cells with symmetric difference 0.
- EXP-0236 reads the virtual serializer `R1521`: its store and load
  widths are exactly 2 and `0x34`, and its hash is the embedded object at
  terrain `+0x540b4`.
- `SAV-CELLLOAD-110`/`111` project the payload: two terrain baselines, one
  area-layer count, ten typed object identity keys, conditional trigger-tail
  state and bounded residue.
- Load overlays an ALM-constructed node or inserts an absent key before the
  later identity pass; it is not an opaque replacement table.

**Confidence.** High for framing, hash identity, widths, overlay and every
positive typed field: independent writer and reader instructions plus the
original corpus discriminator. Medium for the negative residue fields, whose
complete 28-owner scratch census can miss arbitrary aliases.

**Amended.** EXP-0236 and `SAV-CELLLOAD-110`/`111` close the former serializer
and payload Unknowns. `SAV-CELLREC-032` adds four bytes after the table, before
the session block.

### SAV-TAIL-018

- The region's length is identical in all four saves, measured as 4382 bytes,
  and it is ≥ 98.5 % zero.
- Its sparse content begins with the same u32 the object heads carry at `+0x08`,
  holds a 64-bit value written twice that differs per session, `10 27` (10000),
  and a pointer-shaped `e8 8f bc 04`.
- After the last object the stream's final ~404 bytes are all zero.
- With this region the whole stream tiles: head + 437 objects + block array +
  cell-record table + this region + tail leave 51–70 bytes unattributed per
  save, the region's own sparse nonzeros.

**Confidence.** High for the constant size, the zero shares and the tiling,
which are measurements. The 4382 is withdrawn: the writer's own `PUSH`ed lengths
sum to 4374, and the corpus reproduces that, 11/11.

**Unknown.** The region's meaning, as graded here; `SAV-SESS-031` closes it.

**Amended.** `SAV-TRAIL-026` attributes the zero tail: it is `world+0x118`'s own
serializer, after `0xbadface1` and one dword (`retracted.md`). `SAV-SESS-031`
(EXP-0086) names the region field for field as the session `Serialize`
`R0270`, written between `R1360`'s cell records and the sack list,
and re-measures it at 4374 bytes rather than 4382 (`retracted.md`). The region
is not in every save: the call sits inside the arm `SAV-SHAPE-023`'s shape byte
gates.

## Save writer, reader, codec and world halves

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-FRAME-021 | The container is a 16-byte header and a self-describing compressed blob from `0x10`, not a 20-byte header and a body; `u32@0x0C` is the read length. | High | ● active | [EXP-0077](../experiments/EXP-0077-party-and-save/) |
| SAV-CODEC-022 | Codec `R1515` reads a word-count dword, then run and literal opcodes until the source ends; the four opcodes shipped saves never use are all legal. | High | ● active | [EXP-0077](../experiments/EXP-0077-party-and-save/) |
| SAV-SHAPE-023 | A save is a campaign half plus an optional world half: `R0414` writes one byte from `world+0x2c`, and the world snapshot follows only when that byte is 1. | High / Medium | ● active (amended) | [EXP-0077](../experiments/EXP-0077-party-and-save/) |
| SAV-ROSTER-024 | The roster is written in the campaign half, before the shape byte, so every save carries the player list and every `Player`, group and actor. | High / Unknown | ● active | [EXP-0077](../experiments/EXP-0077-party-and-save/) |
| SAV-HEAD-025 | The decoded head is named field for field: two tick counters, the map name, eleven u32 fields, mission number and difficulty, then two player-list dwords. | High / Unknown | ● active (amended) | [EXP-0077](../experiments/EXP-0077-party-and-save/), [EXP-0521](../experiments/EXP-0521-session-clock/) |
| SAV-TRAIL-026 | The stream ends with three writes after the world half: marker `0xbadface1`, the dword `[L07886]`, and `world+0x118`'s serializer, measured as 400 zero bytes. | High / Unknown | ● active | [EXP-0077](../experiments/EXP-0077-party-and-save/) |

### SAV-FRAME-021

- Writer `R0804` serializes into a `CMemFile`, pads it to an even length
  (`L08265`, a test of the low bit of the length → `SetLength(len+1)`), and compresses `len/2` words
  with `R1376`.
- It then writes `"Asg&"`, a zero dword, `0x0BAD0002`, the compressed byte
  length, and the blob (`L07874`…`L07875`).
- It takes `GetPosition()` (`L08266`), optionally appends a `0x100`-byte
  region, seeks back to offset 4 and patches that position in (`L07880`
  `Seek(4,0)`, `L07881` `Write(&pos,4)`).
- Reader `R1305` mirrors it: magic, then a dword it reads and discards
  (`L07879`), then the version, then `u32@0x0C` → `malloc` and `Read` of
  exactly that many bytes from offset `0x10` (`L07876`…`L07877`).
- Consequences:
  - the coded region is `[0x10, u32@0x04)`, not `[0x14, …)`;
  - `u32@0x0C` is the consumed field (the allocation size and the read length),
    and `u32@0x04` is the one nothing reads;
  - `u32@0x04 − u32@0x0C = 16` is an identity, not a measurement;
  - EXP-0046's header field `u32@0x10` is not a header field: it is the blob's
    own first dword, which the decompressor reads (`SAV-CODEC-022`).
- `u32@0x08` is a version and is ordered: `CMP …,0xbad0002` / `JGE` at
  `L07884`, whose failure arm reports the literal "Outdated save file."
  (`L07885`), against "Invalid save file." (`L08267`) for a bad magic.
- Probe: all six framing checks pass 4/4, and the bytes the new framing emits
  are byte-identical to EXP-0046's.
- This amends `SAV-HDR-001` (header length, body start), `SAV-PTR-003` (which
  field has a consumer), `SAV-UNK-006`/`SAV-SIZE-008` (where the word count
  lives), `SAV-EXT-009` (the coded region's start) and `SAV-VER-002` (Medium →
  High that it is a version).

**Confidence.** High. Both directions of the format are read at instruction
level and agree field for field; the corpus then reproduces every relation the
pair predicts, 4/4, including the two the old framing had to measure.

### SAV-CODEC-022

- `R1515(src, srcLen, &dst, &dstWords)`: `*dstWords = *(u32*)src`;
  `*dst = malloc(*dstWords * 2)`; `i = 4`; `while (i < srcLen)` dispatch on
  `src[i] & 0x80` to `R1579` (run) or `R1580` (literal), adding
  each helper's return to `i`.
- Run: `n = op & 0x7f`, emit the following word `n` times, return 3.
- Literal: `n = op` (no mask), copy `n` words from `op+1`, return `2n+1`.
- The four values shipped saves never use are all legal and harmless: `0x00` is
  a literal of no words that costs one byte, `0x7f` a literal of 127, `0x80` a
  run of zero that costs three, `0x81` a run of one.
- There is no end marker: the loop is bounded by the source length and never by
  the output count.
- The compressor `R1376` is the mirror, and its rule is one compare:
  `word[i] == word[i+1]` → emit a run, else a literal (`L08268`…`L08269`).
- The compressor's output buffer is `malloc(srcWords * 2)`: sized to the input,
  with no headroom, so a body the code cannot shrink overruns it.
- The compressor writes the word count into the first dword of its own output,
  which is why the count survives a `.chr` file and a network packet that have
  no container header at all.
- This amends `SAV-PACK-007`, settling its four Unknown opcodes.

**Confidence.** High. The decode is transcribed instruction for instruction and
re-run against the corpus: 12 516 opcodes over four saves, the same output as
EXP-0046, the blob consumed exactly, and the declared word count reproduced 4/4.
The four opcodes' behaviour is read off the arms, not inferred from an absence.

### SAV-SHAPE-023

- The world `Serialize` `R0414` writes the campaign half unconditionally,
  then tests `world+0x2c` (`L07815`, a comparison of the 32-bit field at offset 0x2c of the world object with zero and a branch when zero)
  and writes one byte: `1` followed by the whole world snapshot, or `0` and
  nothing (`L08270` / `L07816`, `ar << (BYTE)`).
- The load arm reads that byte back (`L08271`, `R1581`) and branches on
  it (`L08272`, a branch to `L08273` when zero). Non-zero: destroy the live actors, rebuild
  the world's three sub-objects, re-open the map from the install by name, read
  the snapshot, and set `world+0x2c = 1`. Zero: `world+0x2c = 0` and no map at
  all.
- A mid-mission save and a between-mission save are therefore different shapes,
  and one byte in the stream tells them apart; nothing outside the file is
  needed.
- `R0512` sets `world+0x2c` to 1 when a map finishes loading
  (`PARTY-SESSION-008`), which makes the distinction "is a mission in progress".
- The four shipped saves are all mid-mission and all carry the world half, so
  that corpus exercises one shape of the two; the discrimination here is the
  listing's, not the corpus's.

**Confidence.** High that the two shapes exist and that one byte selects them:
the two arms are a `JZ` and its fall-through, they are exhaustive, and the
reader's branch is the same test on the byte it just read. Medium for the gloss
"between-mission": that `world+0x2c == 0` means no map loaded rests on
`R0512`'s single write of 1, and no save with the byte clear had been
observed. The discriminator named was one save taken from the town screen.

**Amended.** EXP-0086 (`SAV-CITY-030`) witnesses the byte-0 arm: `game0010.sav`,
3 544 bytes, mission number 0, one player, no block array, no cell-record table,
no session block, no `Unit`/`Building`/`Sack` class record. The Medium gloss is
discharged.

### SAV-ROSTER-024

- `[L00380]->R1372(ar)` writes the player list, and through it every
  `Player`, every group and every actor (`PARTY-ROSTER-002`). It is called
  before the shape byte, in both arms (`L08274` store, `L08275` load).
- When the shape byte is 0 the load arm still has actors to restore. It walks
  all of them and zeroes `actor+0x40`, `+0x44` and `+0x5c`
  (`L08276`…`L08277`): the same detachment the character import performs
  (`PARTY-LOSS-006`), from a second and independent site.
- The writer prunes before it writes: for every `Player` whose `+0x28` is 0 it
  removes from that player's group collection every group whose element count is
  0 (`L08278`, `R0197` = `GetCount`, → `R1582`). An empty group
  never reaches the file, and a reader must not expect the map's sixteen slots
  to survive a round trip.

**Confidence.** High: the call order is two instructions either side of the
branch, and the zeroing and the prune are each a short named sequence.

**Unknown.** `Player+0x28`, the flag that decides whether a player is pruned at
all.

### SAV-HEAD-025

- Store order (`L08279`…`L08280`): `u32 world+0x04`, `u32 world+0x00`,
  `CString world+0x28` = the map name, eleven `u32`
  (`+0x11c, +0x124, +0x128, +0x12c, +0x130, +0x134, +0x138, +0x13c, +0x148, +0x144, +0x140`,
  in that order, `+0x148` before `+0x144` before `+0x140`), `u32 world+0x80`,
  `u32 world+0x84`.
- Then `R1372` writes `u32 playerList+0x20` and the list's own element
  count.
- The four head dwords EXP-0048 could only quote (`SAV-STREAM-010`), `10 2 6 5`,
  are therefore mission number, difficulty, `playerList+0x20` and player count,
  and the eleven zeros are `world+0x11c…+0x140`.
- `world+0x80` is the mission number `R0512` parses out of the map file
  name (`PARTY-SESSION-008`). It is also a mode switch on load: when non-zero
  the map name is re-opened as `"Scenario\" + name` (`L08281`…`L08282`,
  literal `L08283`), which is why `10.alm` resolves against `scenario.res`.
- `world+0x84` is the difficulty. The load arm clamps it to 1..3 and silently
  keeps the old value outside that range (`L08284`/`L08285`,
  `UNIT-GATE-012`).
- Probe, 4/4: the head parses, `world+0x00 == world+0x04 >> 4`, the eleven
  dwords are all zero, the difficulty is in 1..3, and the mission number is the
  map name's numeric stem.
- This amends `SAV-STREAM-010` (the four u32s, the two counters) and
  `SAV-MAP-005` (the name's consumer and its `Scenario\` prefix).

**Confidence.** High: the store arm is a straight run of `ar <<` calls with no
branch, the load arm mirrors it, and the corpus reproduces the whole head under
that reading with five independent checks passing 4/4. High that `world+0x04` is
the sub-tick counter and `world+0x00` the full-tick counter it feeds: the `>>4`
identity holds 4/4, and `SESS-TICK-004` read both increments and the pacing
independently from a listing.

**Unknown.** What `playerList+0x20`'s `6` is.

**Amended.** The relation `world+0x00 == world+0x04 >> 4` is corrected: `SAV-1210` finds it false
on `2026-08-02/game0003.sav` (sub-tick 9343, full tick 584) and true on the other 27 of 28 files.
The full-tick count the stepper reaches is `(world+0x04 + 1) >> 4` on all 28 (`SESS-086`). The
field order, the field meanings and the other four checks stand.

### SAV-TRAIL-026

- After the world half: `ar << 0xbadface1` (`L08286`; the load arm tests it at
  `L08053` and reads one more dword only if it matches), `ar << [L07886]`,
  then `world+0x118 → R1491(ar)` (`L07783`), which is the last thing in
  the stream.
- Measured, 4/4: the marker occurs exactly once in each decoded body, the dword
  after it is 0 in all four, and what follows to the end is 400 bytes, every one
  zero. `game0001` has 401, of which 400 are zero; that file's body is
  odd-length before the writer's even pad.
- `SAV-TAIL-018`'s "final ~404 zero bytes" is this tail. The 4382-byte fixed
  region `SAV-TAIL-018` located is not: it sits before the sacks, inside the
  world half.

**Confidence.** High: the three writes are consecutive named instructions, and
the marker's uniqueness, the zero dword and the 400-byte run are exact corpus
measurements.

**Unknown.** `[L07886]`, and what `world+0x118`'s serializer holds; 400 zero
bytes is a shape, not a meaning.

## Player record, between-mission save and session block

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-FLAG-027 | The completed-mission flag is `Player+0x3c`, one byte with domain `{0,1,2}` in the campaign half of every save, set by reporter `R0132`. | High / Medium / Unknown | ● active | [EXP-0086](../experiments/EXP-0086-saves/) |
| SAV-PLAYER-028 | `Player::Serialize` is `R0415`, and its record is seventeen fields in one straight run ending in the object's own pointer. | High / Unknown | ● active | [EXP-0086](../experiments/EXP-0086-saves/) |
| SAV-OBF-029 | Two `Player` fields, `+0x38` (money) and `+0x48`, are XOR-obfuscated on the way out, and two, `+0x54` and `+0x4c`, are silently narrowed to `u16`. | High / Medium | ● active | [EXP-0086](../experiments/EXP-0086-saves/) |
| SAV-CITY-030 | The between-mission save `game0010.sav` witnesses `SAV-SHAPE-023`'s byte-0 arm: 3 544 bytes, no world-half structure, and the whole campaign-half roster. | High / Medium / Unknown | ● active | [EXP-0086](../experiments/EXP-0086-saves/) |
| SAV-SESS-031 | `SAV-TAIL-018`'s fixed region is the session `Serialize` `R0270`: 4374 bytes in the world half, located at the same byte two ways in 11/11 saves. | High / Unknown | ● active | [EXP-0086](../experiments/EXP-0086-saves/) |
| SAV-CELLREC-032 | Four bytes follow the 54-byte cell-record table, before the session block, in 11/11 saves; `SAV-TERRKEY-056` identifies them. | Medium / Unknown | ● active (amended, superseded) | [EXP-0086](../experiments/EXP-0086-saves/) |

### SAV-FLAG-027

- The reporter `R0132` has three writing arms, named by the image's own
  strings:
  - `L08287` sets 1 after logging `"Logic - Mission Complete"` (`L08288`)
    and raising announcement `0xb5`, on `session+0xb3ac == 1`;
  - `L08289` and `L08290` set 2 (`"Logic - Mission Failed"`, `L08291`,
    announcement `0xb4`) on `session+0xb3b4 == 1` and on no living hero
    respectively;
  - `L07927` sets 0 on the re-place arm.
- Lose is tested first (`L08292` before `L08293`), so 2 beats 1 in one tick.
- `Player::Serialize` stores the byte at `L08294` and reads it back at
  `L06703` (`SAV-PLAYER-028`); it is the ninth field of that record.
- Corpus, 12 owner-produced saves of one campaign: `+0x3c == 1` on exactly 3 and
  `0` on 9. The human participant is identified independently as the one player
  with `+0x28 == 0` (`UNIT-OWNER-009`), 12/12.
- The rival, that `TRIG-END-009`'s counter `session+0xb3ac` is the flag, is
  refuted by the third flagged file: `session+0xb3ac == 1` on 2 of the 3, and
  the third has no counter at all, because that save has no world half.
- The counter is a map-lifetime quantity and the latch is the campaign's. A
  consumer that implements the counter loses mission completion at the town
  screen, which is where the campaign needs it.

**Confidence.** High. Both directions of the field are named instructions, the
arms are named by the image's own strings, and the corpus partitions 3/9 with
three independent alignment anchors that a one-byte misalignment breaks
simultaneously: `+0x04 == +0x08` 45/45, `+0x28 == 0` for exactly one player per
save, and `+0x3c` inside the writers' value set 45/45. Medium that `{0,1,2}` is
the complete value set: the instrument is `EnumRefs re:` on the byte-wide
immediate store at displacement `0x3c` (39 hits, 22 owners, 0 orphan) plus the
two zeroing sites `R0201` (the constructor) and `R0131`, and it
cannot see a store through a register or a `memset` of the object. The owner
supplied a labelled set; the reading stands without it
(`EXP-0086/evidence/owner-testimony.md`).

**Unknown.** Whether `2` ever reaches a shipped file: no lost mission is in this
corpus.

### SAV-PLAYER-028

- Store arm `L08295`…`L08296`; load arm `L08297` mirrors it call for call.
- Widths are taken from the four helpers themselves: `L08298` = 1 byte,
  `L08299`/`R1583` = 2, `L08300`/`R0685` = 4, `R0271`/`R0272` =
  `CArchive::Write`/`Read` of the pushed length.
- In order: `CString +0x18` (the name), `u16 +0x04`, `u32 +0x08`, 8 raw bytes
  `+0x10`, `u8 +0x44`, `u32 +0x28`, `u16 +0x2c`, `u32 f(+0x38)`, `u8 +0x3c`,
  `u8 +0x3d`, `u32 f(+0x48)`, `u32 +0x50`, `u16 min(+0x54, 0x7fff)`,
  `u16 min(+0x4c, 0x7fff)`, `u32 +0x58`, `u32 +0x34`, and finally `u32 this`,
  the object's own pointer.
- This confirms `SAV-OBJ-016`'s "slot id twice" as `+0x04` and `+0x08` (45/45)
  and puts `UNIT-OWNER-009`'s `+0x28` in the file.
- Measured: `+0x28` is 0 for exactly one player per save and 1 for every
  scenario owner, 12/12.

**Confidence.** High: a branchless run of `ar <<` calls with a mirroring load
arm, re-measured over 45 records on 12 files where three anchors fail together
under any misalignment.

**Unknown.** `+0x44`, `+0x2c`, `+0x50`, `+0x34`, and what the stored `this` is
for; nothing here reads it. `+0x58` is no longer Unknown: `SAV-726` gives it a
non-constructor command producer, and `UNIT-DERIVE-003`/`HERO-MP-006` already
give it a consumer, the AI mana-floor percentage.

### SAV-OBF-029

- `R1273` is four instructions, `return x ^ 0x5c073f4d`
  (`L08301` combines the argument with the constant 0x5c073f4d by exclusive-or).
- `Player::Serialize` puts `+0x38` and `+0x48` through it in the store arm
  (`L08302`, `L08303`) and through it again in place in the load arm
  (`L08304` → `L08305`), which works because the XOR is an involution.
- `Player+0x38` is `SHOP-BUY-009`'s money, so a reader that skips the transform
  reads ≈1.54 billion gold on every save. This is the first defect a consumer
  ships if it copies the record verbatim.
- `+0x54` and `+0x4c` are dwords in memory written as `u16` after an explicit
  `CMP …,0x7fff` and clamp (`L08306`, `L08307`): a value above 32 767 comes
  back smaller with no message.
- Corpus: the decoded money is 100 on ten saves and 600 on the two most recent;
  the scenario players' `+0x38` decodes to 0, 44/44.

**Confidence.** High: the helper is four instructions, both call sites are
named, the load arm applies the same function to the same field, and the decode
produces small plausible integers where the raw dwords are all within 600 of the
key. Medium that `+0x48` is a quantity of the same kind: it decodes to 5…19 and
rises with the session, and nothing here reads it.

### SAV-CITY-030

- `game0010.sav` is 3 544 bytes against 27 935–46 489 for the eleven mid-mission
  saves, and decodes to 6 110 bytes against 58 314–83 802.
- Absent, each a world-half structure: the block-plane array (`SAV-BLOCK-011`),
  the 54-byte cell-record table (`SAV-CELLREC-017`), the session block
  (`SAV-SESS-031`), and the class records `Unit`, `Building`, `Sack`: 8 class
  records against 11.
- Different in the campaign head (`SAV-HEAD-025`): `world+0x80` (the mission
  number) is 0 while the map-name field still holds a stale `20.alm`; both tick
  counters are 0; `playerList+0x20` is 2 against 6; the player count is 1
  against 5.
- Present: the whole roster (the human participant's `Player`, its groups, the
  party's actors, their inventories) and a `Spell` class record no mid-mission
  save in this corpus carries.
- `Player+0x3d`, which `R0132` requires non-zero before it will report
  anything, is 0 here and 1 in all eleven others.
- The byte immediately before `0xbadface1`, which `SAV-TRAIL-026` fixes as the
  first thing written after the world half, is `0`.

**Confidence.** High that this is the byte-0 shape: every structure the shape-1
arm writes is absent and every structure the campaign half writes is present,
and the two arms are exhaustive, a `JZ` and its fall-through. Medium that the
`0` observed at `markerAt − 1` is the shape byte at that exact offset: the
roster was not walked to its end, so the byte's position is inferred from the
writer's order rather than measured.

**Unknown.** Why `Spell` appears only here.

### SAV-SESS-031

- Store arm `L08308`…`L08309`, in order and with the widths the four helpers
  give:
  - `0x190` = 400 bytes from `session+0xbd34` (`TRIG-STORE-002`'s 100 int result
    slots);
  - `0x3e8` = 1000 from `+0xbec4` (`TRIG-FIRE-007`'s fire-once latches);
  - `0x30` = 48 from `+0x08`;
  - `0x190` = 400 from `+0xa828`;
  - `0x9cc` = 2508 from `+0xa9bc`, which spans `AI-DIPLO-004`'s 50×50 matrix at
    `+0xa9c4`, so the matrix starts 1 856 bytes into the region;
  - `u8 +0xa48`, `u8 +0xa49`, `u32 +0xa4c`;
  - the three outcome integers `u32 +0xb3ac`, `u32 +0xb3b0`, `u32 +0xb3b4`.
- Sum: 4374.
- Two independent locators agree on the byte in 11/11 world-half saves: the
  world `Serialize`'s own consecutive calls (`R1360` at `L08310`,
  `R0270` at `L07814`), and a whole-stream search for a 2 500-byte window
  in the matrix's measured domain, which returns exactly one hit per save and
  none in the between-mission save.
- At the predicted offset the matrix validates 11/11: 0 bytes above 2 over
  2 500, and diagonal 2 at exactly the five occupied slots.
- `session+0xb3ac` (the win counter) reads 1 on the two labelled mid-mission
  saves and 0 on the other nine; `+0xb3b4` is 0 in 11/11; `+0xb3b0` is 0 in
  11/11 and is not named.

**Confidence.** High. The layout is a branchless run of named instructions whose
pushed lengths sum to a figure the corpus then reproduces; two locators derived
from different things agree 11/11; and the matrix check is a prediction that a
wrong length would have moved. The win counter's 2/11 split is a second,
independent witness of the outcome `SAV-FLAG-027` reads off the player.

**Unknown.** The remaining fields in `session+0x08`'s 48 bytes
(SAV-1078/SAV-1079 locate the timer divisor and frequency destination),
`+0xa828`'s 400, `+0xa48`/`+0xa49`/`+0xa4c`, and `+0xb3b0`.

### SAV-CELLREC-032

- The distance from the block array's last record to `SAV-SESS-031`'s region
  start is `2 + 54 × count + 4` in 11/11 saves (counts 174–197).
- A placement without the `+ 4` puts the win counter on the full-tick counter.
- This extends `SAV-CELLREC-017`.

**Confidence.** Medium: an exact, identical measurement on eleven files.
`R1360`'s own listing was not walked for this claim.

**Unknown.** What the four bytes are, as graded here; the discriminator named
was the tail of `R1360`.

**Amended.** `retracted.md` supersedes the "unattributed" label: EXP-0148
(`SAV-TERRKEY-056`) reads the four bytes as the terrain object's own identity
key, written inline by `R1360`. The `2 + 54 × count + 4` measurement
stands.

## Serializable classes and the Token head

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-CLASS-033 | The image carries 28 MFC serializable class descriptors, not the 11 four saves show, and `game0010.sav` already carries a twelfth, `Spell`. | High / Unknown | ● active | [EXP-0145](../experiments/EXP-0145-save-members/) |
| SAV-TOKEN-034 | The placeable head is 37 bytes written by one routine, `Token::Serialize` `R0950`: the 12-byte position object, then eight fields ending in `this` and `this+0x14`. | High | ● active (amended) | [EXP-0145](../experiments/EXP-0145-save-members/), [EXP-0224](../experiments/EXP-0224-sav-token-position/) |
| SAV-PTRMAP-035 | Token identity and its named reference sites use the writing process's object addresses as keys; absent-key-to-zero holds at these sites, not at every reference repair. | High | ● active (partially retracted) | [EXP-0145](../experiments/EXP-0145-save-members/) |

### SAV-CLASS-033

- `.data` holds 28 MFC-shaped class descriptors: name pointer, object size,
  schema `1`, `createObject` in `.text`.
- They name their own hierarchy:
  - `Token` (`L08311`) is the base of `Unit`, `Item`, `Effect`, `Building`,
    `Sack`, `VirtualCaster` and `SpellEffect`;
  - `Humanoid` sits between `Human` and `Unit`;
  - `Armor`/`Shield`/`Weapon` derive from `Item`;
  - `Outpost`, `Tavern` and `Shop` derive from `Building`;
  - `PointEffect`, `AreaEffect` and `SpellTransport` from `SpellEffect`;
  - `Player`, `Diary`, `Spell`, `Spellbook`, `TableLine` and the three
    `CMultiShop*` from `CObject`.
- `game0010.sav` introduces a `Spell` class record, so the eleven-class set is
  not even the observed set. A reader that knows only eleven names fails on a
  save carrying a spellbook, and by construction on any save from a mission with
  a shop, a tavern or an outpost.
- This amends `SAV-STREAM-010`'s “the four shipped saves introduce 11”, which
  was a true statement about four files and never about the format.

**Confidence.** High: a raw `.data` scan for the descriptor shape, a data scan
and therefore immune to Ghidra's function coverage. It is blind only to a
descriptor built at run time or living outside `.data`, neither of which is
claimed against.

**Unknown.** What the 17 unobserved classes write: their `Serialize` bodies are
unread, `Spell`'s included.

### SAV-TOKEN-034

- In file order: the 12-byte position object decoded by `SAV-TOKENPOS-074`,
  copied raw from `*(this+0x10)` by `R1494`; then `u32 this+0x04`,
  `u8 this+0x0c`, `u16 this+0x0e`, `u32 this+0x08`, `u16 this+0x18`,
  `u32 this+0x1c`, `u32 this`, `u32 this+0x14`.
- `SAV-ID-015`'s creation-order id is the `u32` at head +12, and its map unit id
  at head `+0x13` is the low `u16` of the `u32` at head +19.
- The routine is reached from `Building`'s, `Item`'s, `Effect`'s, `Sack`'s and
  `Unit`'s first instruction after the frame.
- `SAV-TOKENPTR-075` carries the remaining post-load pointer-rebind question.

**Confidence.** High: the store arm fixes the 37-byte programme; `Building`'s 40
own bytes make its 77-byte record chain in twelve saves; and the position helper
independently carries a literal width of 12.

**Amended.** EXP-0224 (`SAV-TOKENPOS-074`) decodes the first twelve bytes and
closes the earlier Unknown on them.

### SAV-PTRMAP-035

- The head's eighth field is `this` itself, written verbatim; the ninth is
  `this+0x14`, another object's address.
- The load arm binds them. After reading the eighth into a local it calls
  `SetAt` (`L08312`) on the map at `[L00285] + 0x88`: old address → the
  object just constructed. After reading the ninth into `this+0x14` it calls
  `R1371`, which looks that value up in the same map and overwrites the
  field with the result, or with `0` when absent (`L08313`).
- `Player::Serialize`'s seventeenth field and `Diary::Serialize`'s `+0x2c` are
  the same mechanism at other sites.
- For these sites an authored identity key must be unique and nonzero, and each
  reference must name its target's key; the numeric value is arbitrary but the
  identity is not.
- The absent-key-to-zero rule belongs to `R1371` and these named sites, not
  to every reference repair. Order helper `R0170` instead preserves a missing
  key (`SAV-HUMRESUME-460`, `SAV-ACTORINPUT-547`).

**Confidence.** High, and the address reading is discriminated rather than
assumed. Decoding the eighteen `Building` records of `game0000.sav` with the
programme gives keys `2c94570 2c94670 2c946f0 2c94770 2c947f0 2c94870 …`,
ascending in `0x80`/`0x100` steps for a class whose object size is 108 bytes: an
allocator's stride, not what a counter, a hash or an index would produce. The
same decode shows the reference field taking one of a small set of values shared
by groups of buildings (`2c8a670` for the first seven, `2c8bae0` for the next
four), pointing at a handful of longer-lived objects. That is consistent with
the owning `Player`, though nothing here names the target's class.

**Amended.** `retracted.md` withdraws the former universal clauses "every
cross-object reference" and "a reference to a key no object claims silently
becomes null": the shared key map does not imply a shared missing-key policy.
The field-specific rule above replaces them.

## Serialize bodies and record lengths

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-MEMBER-036 | The eleven `Serialize` bodies are read field by field, and a record's length is not a constant: counted lists, a string and presence flags vary it. | High / Unknown | ● active (amended, superseded) | [EXP-0145](../experiments/EXP-0145-save-members/), [EXP-0146](../experiments/EXP-0146-embedded-members/) |
| SAV-BLDG-037 | A `Building` record is exactly 77 bytes, `Token`'s 37 plus 40 of its own, and chaining consecutive instances confirms the length. | High | ● active | [EXP-0145](../experiments/EXP-0145-save-members/) |
| SAV-OBFCEN-038 | Within the disassembled `Serialize` call set, the XOR obfuscation and the `0x7fff` clamp occur only in `Player::Serialize` `R0415`. | Medium | ● active | [EXP-0145](../experiments/EXP-0145-save-members/) |
| SAV-EMBED-039 | The eight embedded-object sites a `.sav` record serializes hold four classes: an unnamed `u16` list at five sites, a `Spellbook`, a `CDWordArray` and a `CWordArray`. | High / Unknown | ● active | [EXP-0146](../experiments/EXP-0146-embedded-members/) |
| SAV-WLIST-040 | Five of the eight embedded sites are one class with no name anywhere in the image: a list of `u16`, vtable `L08172`, written as a count and two bytes per element. | High | ● active | [EXP-0146](../experiments/EXP-0146-embedded-members/) |
| SAV-SPELLBK-041 | `Spellbook::Serialize` `R1584` writes a `u32`, a count `n` and then `n-1` references, skipping index 0. | High / Unknown | ● active | [EXP-0146](../experiments/EXP-0146-embedded-members/) |
| SAV-DIARY-042 | A `Diary` record is a `CDWordArray`, a `CWordArray` and a `u32` reference; its only fixed part is those four bytes. | High / Unknown | ● active | [EXP-0146](../experiments/EXP-0146-embedded-members/) |
| SAV-HUMAN-043 | A `Human` record is a `Unit` plus 24 raw bytes and thirteen object references, the references `SAV-MEMBER-036` omits. | High / Medium | ● active | [EXP-0146](../experiments/EXP-0146-embedded-members/) |
| SAV-SPELL-044 | A `Spell` record is nine bytes, and its load binds an identity key in the `Token` identity map although `Spell` is not a `Token`. | High / Unknown | ● active | [EXP-0146](../experiments/EXP-0146-embedded-members/) |
| SAV-UNITLEN-045 | A `Unit` record has no fixed length; its floor is `609 + L` bytes, where `L` is the name's length. | High / Medium | ● active | [EXP-0146](../experiments/EXP-0146-embedded-members/) |
| SAV-EFFCHAIN-046 | `Effect` is 44 bytes; chaining cannot reach a second instance because each `Effect` is an element of its owner's `+0x20` list, not written consecutively. | High | ● active (amended, partially retracted) | [EXP-0146](../experiments/EXP-0146-embedded-members/), [EXP-0245](../experiments/EXP-0245-sav-unit-subtree-reconciliation/) |
| SAV-SERPOP-047 | The image carries 35 serializable class descriptors, not 28: seven more at schema 0, of which `CDWordArray` and `CWordArray` are `Diary`'s two members. | High / Medium / Unknown | ● active (amended) | [EXP-0146](../experiments/EXP-0146-embedded-members/), [EXP-0247](../experiments/EXP-0247-sav-nontoken-serializers/) |

### SAV-MEMBER-036

- Each routine is located from its class's constructor
  (a 32-bit store of the vtable address into the new object's first field) and read from vtable slot 2. Slots 3 and 4
  hold the same two empty stubs `R0531`/`L04402` in all eleven, which fixes
  slot 2 independently.
- The programmes are in the experiment's `evidence/programmes.md`.
- `Item` carries a counted list at `+0x20` (`R0951`: `u32` count, then
  that many `ar << CObject*`).
- `Unit` carries a `CString` at `+0x80`, fourteen `u16` at `+0x84`…`+0x9e`,
  three object references and two presence flags: a `u8` `1`/`0` at each of
  `this+0x7c` and `this+0x140`, each followed by a sub-object only when set.
- `Sack` carries a list too.
- `Shield::Serialize`'s store arm is empty: a `Shield` is an `Item` plus 22 raw
  bytes.
- `Human::Serialize` writes nothing of its own. It calls `Humanoid::Serialize`,
  which calls `Unit::Serialize` and appends 24 raw bytes from `+0x1cc`, then
  thirteen object references (`SAV-HUMAN-043`). Everything else in
  `Human::Serialize`'s body is the load arm.

**Confidence.** High for every sequence. Each store arm was disassembled in
full, and the tool-reduced sequence was then checked against the raw listing for
all eleven; that check caught `Weapon`'s object reference, `Unit`'s two embedded
sub-objects and `Shield`'s empty arm.

**Unknown.** What any field named only by offset means, and what the unnamed
`u16` lists hold (`SAV-EMBED-039`). No total length is claimed here for `Unit`,
`Human`, `Diary` or a group record.

**Amended.** Two clauses are narrowed ([`retracted.md`](retracted.md),
EXP-0146). The `Human` clause said `Humanoid::Serialize` appends only the 24 raw
bytes; it also writes thirteen `ar << CObject*` (`SAV-HUMAN-043`). The Unknown
clause listed seven embedded sites; the count is eight, with `Group+0x20`
dispatched first in `R1143` (`SAV-EMBED-039`). The ten other field
programmes, the presence flags, the counted lists and `Shield`'s empty arm are
untouched. `SAV-EMBED-039` also supersedes the Unknown clause that the classes
of the objects embedded at `Unit+0x15c`, `Unit+0x178`, `Diary+0x04`,
`Diary+0x18`, `Group+0x4c`, `*(Unit+0x140)` and `*(*(Unit+0x158)+0x90)` were
not read: it reads an unnamed `u16` list at `Unit+0x15c`, `Unit+0x178`,
`*(*(Unit+0x158)+0x90)`, `Group+0x4c` and `Group+0x20`, a `Spellbook` at
`*(Unit+0x140)`, a `CDWordArray` at `Diary+0x04` and a `CWordArray` at
`Diary+0x18`.

### SAV-BLDG-037

- `Building::Serialize` `R1585` = `Token` (37) + 22 raw bytes from
  `+0x52` + `u8 +0x40`, `u16 +0x42`, `u16 +0x44`, `u16 +0x46`, `u8 +0x48`,
  `u8 +0x60`, `u8 +0x61`, `u32 +0x64`, `u32 +0x68`: 40 bytes of its own, no
  count, no string, no branch.
- Buildings are written consecutively, so the length is testable. Stepping 77
  bytes from the class record and requiring the next word to be the instance tag
  chains 18/18 instances in six saves and 30/30 in five. Every chain stops on a
  `0000` word, the null arm of `ar << CObject*`, with no slack. This confirms
  `SAV-OBJ-016`'s 77.
- The head fields land where `SAV-TOKEN-034` says. The identity key at +29 is
  distinct for all 18 and all 30. The creation-order id at +12 runs 2..19 on the
  18-building map and 2..62 on the other: `SAV-ID-015`'s law, read at an offset
  the corpus model did not have.
- Three internal checks the length alone does not give:
  - `Building+0x40` reproduces the head's `u16` at +17 in 18/18 records.
  - The dword at head +19 runs `1 2 3 19 20 26 28 30 31 32 33 43 …`, the map's
    own type-4 record id (`SAV-ID-015`), read at the offset this model predicts.
  - `+0x42`/`+0x44` are an equal pair (`1000/1000`, `30000/30000`, `100/100`,
    `2000/2000`), and `+0x64`/`+0x68` pair the same way (`506/511`, `15/15`,
    `63/63`). Both have the shape of a current/maximum pair; that is a reading,
    not a decode.

**Confidence.** High. Two independent derivations meet exactly, at two different
map populations, with an exact terminator; 76 and 78 both fail on the second
instance. None of the three internal checks was used to fix an offset.

**Unknown.** What the two `+0x42`/`+0x44` and `+0x64`/`+0x68` pairs mean.

### SAV-OBFCEN-038

- Over the eleven `Serialize` bodies and every sub-serializer they call,
  a call to `R1273` (`return x ^ 0x5c073f4d`) occurs at four sites, `L08302`,
  `L08303`, `L08304` and `L08314`, all inside `R0415`, two in each
  arm.
- `min(v, 0x7fff)` occurs at four sites, `L08306`/`L08315` and
  `L08307`/`L08316`, also all inside it.
- The only other `0x7fff` in the disassembled set is `L08317`, the class-tag
  threshold inside `CArchive::WriteObject`.
- `SAV-OBF-029`'s two traps therefore do not generalise: for every other class
  the bytes in the file are the bytes in memory.

**Confidence.** Medium. The instrument is a grep over the transitive call set
that was disassembled: the eleven bodies plus `Token`, `Humanoid`, and the list,
raw-block and primitive helpers. It is blind to a transform applied outside
`Serialize`, that is, to a field already held obfuscated in memory, which a
corpus cannot see either. A High grade would need the writers of each field, not
the writers of the record.

### SAV-EMBED-039

- A site that dispatches with an indirect call through the dword at offset 0x8 of the object's table names no class in the
  listing. The class is reached from the instruction that builds the object: for
  an embedded member, the enclosing constructor's
  formation of the address `this+N` as the object, followed by a call to `<ctor>`; for a pointer member, the serializer's own
  load arm, which allocates a literal size and calls the constructor there. Then
  from the constructor's literal 32-bit store of `<vtable>` into the object's first field, then from slot
  0, `GetRuntimeClass`, which returns `<descriptor>`, and whose
  descriptor carries the class's name.
- `Unit+0x15c`, `Unit+0x178`, `*(*(Unit+0x158)+0x90)`, `Group+0x4c` and
  `Group+0x20` are one class: ctor `R1545` called with `0xa`, 28 bytes,
  vtable `L08172`, `Serialize` `R1546`.
- `*(Unit+0x140)` is a `Spellbook`: descriptor `L08318`, ctor `R0462`,
  vtable `L08319`, `Serialize` `R1584`.
- `Diary+0x04` is a `CDWordArray` (`L08320`, `Serialize` `L08153`) and
  `Diary+0x18` a `CWordArray` (`L08321`, `L08154`).
- All seven offsets `SAV-MEMBER-036` publishes are where it says. `Group+0x20`
  is an eighth site that claim does not carry, and it is first in a group
  record's file order.
- `Serialize(*(*(Unit+0x158)+0x90))` is inside `R0199`, not in
  `Unit::Serialize`.

**Confidence.** High. Each class was reached from an instruction that builds or
names the object, never from the size of the hole it fills. Slots 3 and 4 of all
four vtables hold the same two empty `CObject` stubs `R0531`/`L04402` that
fix slot 2 as `Serialize`, the fingerprint `SAV-MEMBER-036` used.

**Unknown.** What the u16 lists hold, and what any field named only by offset
means.

### SAV-WLIST-040

- `R1545` writes vtable `L08172` and zeroes six dwords, taking its one
  argument (always `0xa`) into `+0x18`. The object is 28 bytes with the shape of
  a linked list plus a block allocator.
- Its `GetRuntimeClass` (`L08322`) returns `L08323`, `CObject`'s own
  descriptor: name `CObject`, size 4, schema `0xffff`, null createObject. The
  class does not override it and has no runtime descriptor of its own.
- `Serialize` `R1546` writes `CObject::Serialize` (empty), then a count
  through `CArchive::WriteCount` `L08324` (`u16` when `n < 0xffff`, else
  `u16 0xffff` then `u32 n`), then walks the node chain from `*(this+0x4)`,
  emitting two bytes per element through `R1586`, whose `SHL 1` is the
  width.
- The load arm reads the count, then two bytes per element, and appends each
  through `R1587`.
- It is not one of `SAV-CLASS-033`'s 28 and cannot be named: a consumer must
  refer to it by its vtable.

**Confidence.** High. Both arms were disassembled in full. The element width is
read from the left shift by 1 in the element helper and from the load arm's
its 16-bit load, not inferred from a corpus stride.

### SAV-SPELLBK-041

- `R1584`: `w:u32 this+0x18`, then `w:u32 n`, where `n` is the element
  count of the array at `this+0x04` (`R0030`), then
  `for i = 1; i < n; i++` an `ar << CObject*` (`R1110`) of element `i`. A
  record carries `n` and `n-1` references.
- The load arm mirrors it: read the `u32`, read `n`, `SetSize(n, -1)` through
  `R1588`, then fill `i = 1 .. n-1` through `L08325`.
- The `Spellbook` is reached only through `Unit`'s presence flag at `+0x140`. It
  is serialized directly through its vtable, so it never appears as a class
  record of its own; the `Spell` objects inside it do.

**Confidence.** High. Both arms were read in full; the skipped index is the same
literal `1` in both loops and the same comparison against the array's own size.

**Unknown.** What the `u32` at `+0x18` counts.

### SAV-DIARY-042

- `Diary::Serialize` `L08081` dispatches into the embedded objects at `+0x04`
  and `+0x18`, then writes `w:u32 this+0x2c`, a reference fixed up on load
  through `R1371`.
- `Diary+0x04` is a `CDWordArray`: `Serialize` `L08153` writes the count
  through `L08324`, then `m_nSize << 2` raw bytes from `m_pData`.
- `Diary+0x18` is a `CWordArray`: `L08154`, the same with `m_nSize << 1`.
- The record is `count + 4n`, `count + 2m`, `u32`: parallel arrays of 4-byte and
  2-byte entries whose lengths need not agree.
- Both classes carry runtime descriptors at `L08320`/`L08321` with schema 0.

**Confidence.** High. Both array bodies are the runtime's own and were read in
both directions, including the `IsStoring` test, the inverted bit 0 of `ar+0x14`.

**Unknown.** What the two arrays hold. `SAV-OBJ-016`'s 1488-byte Diary extent is
neither confirmed nor refuted, since the length is a function of two counts.

### SAV-HUMAN-043

- `Humanoid::Serialize` `R0954` calls `Unit::Serialize` and writes raw `0x18`
  from `this+0x1cc`. It then runs `for i = 1; i < 0xd; i++`, emitting
  `ar << CObject*` of `*(this + i*4 + 0x198)`: twelve references, index 0
  skipped, the `Spellbook` idiom. Then it emits one more, `*(this+0x1e4)`.
- The load arm mirrors both loops against the same literal `0xd`.
- A walk built on `Unit + 24` desynchronises eight bytes into the first `Human`
  of a save and never recovers; a walk built on `Unit + 24 + 13 objref` tiles.
- In every one of fifteen preserved saves the thirteenth reference resolves to
  the `Diary` and the twelve to equipment.

**Confidence.** High for the thirteen references and the skipped index: both
arms, one literal bound. Medium for reading the twelve as equipment slots and
the thirteenth as the `Diary`: fifteen saves agree and nothing in the image
names them, which is corpus agreement and caps there.

### SAV-SPELL-044

- `Spell`'s descriptor `L08326` gives createObject `L08327`, which allocates
  `0x14` and calls the constructor `L08328`. Its
  the constructor's store of L05044 into the new object's first field gives the vtable; slot 2 is `R1013`.
- Record: `u8 +0x08`, `u8 +0x09`, `u8 +0x0a`, `u16 +0x0c`, `u32 this`. The last
  is the object's own address.
- The load arm reads the four fields, then binds that dword to the fresh object
  with `SetAt` (`L08312`) on the map at `[L00285] + 0x88`, the identity
  map `Token::Serialize` uses (`SAV-PTRMAP-035`), even though `Spell` derives
  from `CObject` and carries no 37-byte head.
- It then looks `+0x08` up in the registry at `L05046` and stores the result
  at `this+0x04`.
- `Spell` was one of the seventeen classes `SAV-CLASS-033` left unread, and the
  only one already observed in a save.

**Confidence.** High. The body was disassembled in full in both directions. The
nine bytes tile `game0010.sav` and `game9999.sav`, which carry five `Spell`
records each, inside a walk that closes on the whole object population.

**Unknown.** What the three bytes and the word mean.

### SAV-UNITLEN-045

- `Unit::Serialize` `R0210`, with every embedded class read, contains four
  counted lists (`+0x20` as `u32` count + objrefs; three u16 lists at `+0x15c`,
  `+0x178` and `*(+0x158)+0x90`), a `CString` at `+0x80`, three object
  references and two presence flags. No constant length exists.
- The fixed part sums to 603 bytes: 37 head + 4 list count + 2 + 2 + 462 raw
  (24+22+24+64+180+148 from `+0xa6`, `+0xbe`, `+0x114`, `+0xd4`, `*(+0x154)`,
  `*(+0x158)`) + 2 + 19 + 1 (the `CString`'s length byte) + 55 + 1 + 1 + 17.
- Adding `L` and three object references at two bytes each when null, a `Unit`
  cannot be shorter than `609 + L`.
- `SAV-OBJ-016`'s 621 is `609 + 12`, consistent with a twelve-character name.
  That is arithmetic, not a second measurement.
- `Human` adds 24 bytes and thirteen references, at least `24 + 26` more, so no
  save can hold a `Human` shorter than a `Unit`.

**Confidence.** High for the programme and the 603: each width is read from the
primitive's own store instruction, and the whole record is tiled against fifteen
files. Medium for `609 + L` reproducing `SAV-OBJ-016`'s 621: the arithmetic
works, but the 621 was measured a different way and only one name length makes
them meet.

### SAV-EFFCHAIN-046

- Three readings were enumerated before measuring: the length is wrong; the
  instances are not written consecutively; the first instance does not follow
  its class record. A programme-driven walk kills the first and the third.
- `Effect::Serialize` `L08329` is `Token`(37) + `u8 +0x3c` + `u8 +0x3d` +
  `u32 +0x40` + `u8 +0x0c` = 44, no branch and no count. The walk consumes 44
  for each instance and lands on the enclosing record's next field.
- Every class record in every save is followed immediately by its first
  instance's body, which the walk assumes and would break on.
- The second reading stands: an `Effect` is an element of the `+0x20` list of
  the object it is attached to. In `game0010.sav` it belongs to an `Item` whose
  own record continues for twelve further bytes after the `Effect` ends and only
  then introduces the next object.
- `SAV-BLDG-037` could chain `Building` because buildings are consecutive;
  nothing steps from one `Effect` to another.
- This refutes nothing in `SAV-OBJ-016`'s 232 except its interpretation as an
  `Effect` record length, which it cannot be.

**Confidence.** High. 44 is read from the complete body, and `game0010.sav`
walks 30 of 30 programme objects with nothing unattributed.

**Amended.** The support clause from the independent tag scan is withdrawn
([`retracted.md`](retracted.md), EXP-0245). A current repeat of the fifteen-file
population gives 12 of 15 exact first-Player agreements, and `SAV-TAGSCAN-158`
proves that tag-shaped raw words can move the scan without moving the programme.
The 44-byte programme and the non-consecutive placement stand.

### SAV-SERPOP-047

- `SAV-CLASS-033`'s scan (`savclass -mode all`) requires schema `1` and both the
  descriptor record and its name string in `.data`.
- Relaxing every filter over `.data` and `.rdata` (`-mode any`) finds 117
  descriptor-shaped records:
  - 82 at schema `0xffff` with a null createObject: dynamic only, so not
    creatable from a stream;
  - 28 at schema 1;
  - 7 at schema 0, serializable all the same: `CDib`, `CStringArray`,
    `CDWordArray`, `CWordArray`, `CByteArray`, `CMapStringToString`,
    `CMapStringToOb`.
- Six of the seven have their name strings in `.rdata`, the filter that hid
  them.
- `CDWordArray` and `CWordArray` are reached from a `.sav` object graph as
  `Diary`'s two members. That located route dispatches them directly and
  therefore emits no class record; 55 preserved SAV paths likewise contain no
  schema-0 class record.
- Their non-null creation routines still make generic archive object dispatch
  structurally possible.
- `SAV-CLASS-033`'s observed tag namespace survives; the former categorical
  "never" does not.

**Confidence.** High for the two-section census and the direct `Diary` route.
Medium for zero schema-0 class records in 55 paths and 31 distinct digests.

**Unknown.** Actual SAV class-record reach outside the located route, that is,
whether another SAV producer emits a schema-0 class record. A run-time-built
descriptor remains outside the census.

**Amended.** The clause that the two arrays never appear as class records
because they are dispatched directly is narrowed to the located `Diary` route;
reach through another producer is Unknown ([`retracted.md`](retracted.md),
EXP-0247, `SAV-CONTSER-188`, `SAV-CONTSER-194`). The 35-descriptor census and
the direct `Diary` members stand.

## Object walk: owner, unit fields and the document

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-OWNER-048 | The human participant's own objects are the subtree of the stream's first top-level record, and two further mechanisms agree, over 23 files (18 distinct). | High / Medium / Unknown | ● active | [EXP-0147](../experiments/EXP-0147-save-object-walk/) |
| SAV-UNITFLD-049 | The fourteen stat words of a `Unit` record are named, with their alignment fixed by two constructor immediates, and a monster's carry capacity is a fossil of the default body. | High / Unknown | ● active (amended) | [EXP-0147](../experiments/EXP-0147-save-object-walk/), [EXP-0311](../experiments/EXP-0311-unit-mover-route/) |
| SAV-CARRY-050 | What a unit carries and knows is four separate mechanisms: the `+0x7c` inventory, twelve worn armour references, the `Spellbook` and a thirteenth reference. | High / Medium | ● active (amended, partially retracted) | [EXP-0147](../experiments/EXP-0147-save-object-walk/), [EXP-0285](../experiments/EXP-0285-whole-item-transfer/) |
| SAV-DEATH-051 | Something no longer in play is removed from the owner graph and kept in the exact top-level dead list; stage and runtime id, not one health constant, identify the terminal state. | High / Medium | ● active (amended, partially retracted) | [EXP-0147](../experiments/EXP-0147-save-object-walk/), [EXP-0237](../experiments/EXP-0237-sav-dead-actor-projection/) |
| SAV-TOPLVL-052 | The EXP-0145 tag scan's 61 `Player` hits include 2 zero-run artefacts that `SAV-OBJ-016`'s doubled slot id rejects. | High / Unknown | ● active (amended, partially retracted) | [EXP-0147](../experiments/EXP-0147-save-object-walk/) |
| SAV-DOC-053 | The top-level document `Serialize` `R0414` enumerates everything a save contains, in order, each width taken from the primitive it calls. | High / Unknown | ● active (amended, partially retracted) | [EXP-0148](../experiments/EXP-0148-save-seen-record/), [EXP-0055](../experiments/EXP-0055-tick-order/), [EXP-0222](../experiments/EXP-0222-sav-residual-census/), [EXP-0245](../experiments/EXP-0245-sav-unit-subtree-reconciliation/), [EXP-0246](../experiments/EXP-0246-sav-token-class-serializers/) |
| SAV-PLDIARY-054 | `Player::Serialize` has a third tail call, a `Diary` at `Player+0x40`, and with it the object graph is walkable end to end. | High | ● active | [EXP-0148](../experiments/EXP-0148-save-seen-record/) |

### SAV-OWNER-048

Three independent tests over 23 files, 18 distinct:

- Structural. Walking consecutively from offset 75, the first record is a
  `Player` on 23/23, its slot id is 1 on 23/23 and its `+0x28` is 0 on 23/23:
  `SAV-PLAYER-028`'s human-participant field, with `UNIT-OWNER-009`'s meaning.
  The name is not constant (twenty-two carry `Danath`, one `Fergard`), so this
  is a position, not a string match.
- Referential. Every actor's `Token` head carries a reference dword at file
  33..36, which EXP-0145 read as `this+0x14` and `UNIT-OWNER-009` read, from the
  ALM spawner's own store `L06024`, as the owning `Player`. For 276 of 276
  actors reached across the corpus, that dword equals the identity key (the
  trailing `this`) of the `Player` whose group actor list the actor was walked
  inside: 0 mismatches, 0 orphans.
- Field. `+0x28 == 0`.

A consumer needs none of the scanning the third test implies: the party is
record one.

The rival that would also fit 276/276 is excluded by value. If `+0x14` named the
enclosing group rather than the `Player`, the nesting count would be identical,
but a group record carries no identity key (`SAV-OBJ-016`: groups are a direct
inline call with no class record), and the observed dword equals the `Player`'s
own.

**Confidence.** High for all three. Two are independent, a container relation
the writer emits by recursion against a pointer written 37 bytes into a
different record, and cannot agree 276/276 by accident; the first-record rule is
a position measured on every file. Medium that `+0x28 == 0` alone identifies the
human side: on `game0003` and `game0004` the tag scan reports a `Player` at a
run of zeroes whose `+0x28` also reads 0, and only `SAV-OBJ-016`'s doubled slot
id rejects it.

**Unknown.** Multiplayer: every save here has exactly one participant.

### SAV-UNITFLD-049

- `Unit::Serialize` writes `u16` at
  `+0x84 +0x86 +0x88 +0x8a +0x8c +0x8e +0x90 +0x92 +0x94 +0x96 +0x98 +0x9a +0x9c +0x9e`,
  EXP-0145's order unchanged.
- Alignment anchor: `UNIT-CTOR-004`'s `R0160` writes `0x64` to `+0x98`
  (`L04043`) and `0x32` to `+0x9e` (`L04044`). The file reads 100 and 50 at
  those two word positions on 100/100 `Human` records, and a run shifted by one
  word cannot place two different constants at two predicted positions at once.
- That fixes the reading
  `body, reaction, mind, spirit, speed, ?, ?, capacity, health, healthMax, healthRegen, mana, manaMax, manaRegen`,
  nine of which are `UNIT-STREAM-001`'s slots.
- Relations over 276 actors: `+0x94 <= +0x96` 276/276, `+0x9a <= +0x9c` 276/276,
  `+0x12c >= 1` 276/276.
- `+0x92` is 300 on 176 of 176 `Unit` records regardless of that unit's own
  body, and `10*(+0x84) + 1` on 100 of 100 `Human` records. `UNIT-CTOR-004`
  computes `capacity = body x 10` (`L08330` multiplies the body by the constant 0xa) from the
  constructor's default body of 30, before any template is streamed, and nothing
  recomputes it on the units arm. A monster's carry capacity is a fossil of the
  default; a person's is live.
- `UNIT-CTOR-004`'s sight `+0xa4 = 0` occurs in no record of the 276. Over 1,274
  records in the 31 preserved streams that no project-written input reaches, it
  occurs in no record (`SAV-796`). The three records that carry it in the wider
  1,420-record census are one generated candidate of this project's, a second
  copy of it, and the original's own resave of the first, agreeing in all 24
  scalar columns. The zero this claim looks for is not what the constructor
  leaves at word width, which is `0x0500` = 1280 and occurs 39 times in the same
  census.
- The sight field's values: on `Unit`, `1024 x {1, 1.25, 1.5, 2}` over the 176
  records measured here, and six values over 826 original-authored records,
  `1024 x {1, 1.25, 1.5, 1.75, 2, 2.25}`, which in the field's published
  1/256-cell unit is 4 to 9 whole cells. On `Human`, 1382..1658 over the 100
  measured here, and 1331..2037 over 448 original-authored records, 19 of them
  outside 1382..1658 (`SAV-795`).
- `+0x6c`, which `HERO-DEATH-026` loads at the first death tick, is non-zero
  with 25 distinct values on live `Unit` records at stage 0. A consumer must not
  read it as a countdown unless the stage byte says the actor is dying.

**Confidence.** High for the nine named words, the alignment anchor and the
capacity split: each is a named constructor instruction, and each relation is a
different arithmetic on a different pair, so one shift cannot preserve them all.
The six raw blocks are untouched: what lives inside them is `UNIT-COMBAT-006`'s
at its own confidence, not this claim's.

**Unknown.** `+0x18` (`SAV-636`); `+0x6c` on a live actor; the `+1` in the
human capacity relation, which no published claim attributes.

**Amended.** EXP-0321 narrows the two sight enumerations, which were stated as
closed sets: `1024 x {1, 1.25, 1.5, 2}` on `Unit` and 1382..1658 on `Human` hold
over this claim's 176 and 100 records, not over the preserved corpus (`SAV-795`,
[`retracted.md`](retracted.md)). `SAV-796` confirms the zero's absence over the
larger population. The rest of this claim's population arithmetic is unaffected.
`SAV-635` (EXP-0311) names four fields this card listed as Unknown: `+0x8e` is
the carried weight and `+0x90` the derived load (`ITEM-LOAD-005`,
`HERO-SIGHT-007`), `+0xa2` the health-regeneration remainder byte
(`HERO-REGEN-021`, `SAV-REGENSTORE-529`), and `+0xa4` the `u16` sight radius in
1/256-cell units (`AI-SIGHT-092`, `HERO-SIGHT-007`).

### SAV-CARRY-050

- Measured over 276 actors (100 `Human`, 176 `Unit`):
  - the inventory at `+0x7c` is presence-flagged and present on 276/276,
    matching `UNIT-CTOR-004`'s constructor allocating `0x24` bytes there
    unconditionally; it is non-empty on 28 `Human` and 0 `Unit`;
  - the twelve worn references at `+0x198 + 4i`, `i = 1..12`, are filled 254
    times over the 100 `Human` records;
  - the `Spellbook` at `+0x140` is present on 2 of 100 `Human` and 0 of 176
    `Unit`;
  - the thirteenth reference `+0x1e4` is non-null on 24 of 100.
- The twelve are worn armour by a consumer reading, not by corpus agreement, on
  evidence `ITEM-CORPSE-034` carries:
  - `R0981`, the `vt+0x44` the death routine dispatches, loops
    `L04745`, comparing the 32-bit local at frame offset -0x4 with 0xd, over
    `L04746`, loading the dword at offset 0x198 plus 4 bytes times the loop index, appending each into
    `actor+0x7c`;
  - `Armor::Equip` refuses `Slot == 0` outright (`L04710`, the "Illegal armor"
    arm);
  - all 30 shipped `Armors` rows carry `Slot` in 1..12;
  - `ITEM-CORPSE-034` names `R0954`, the serializer itself, as running
    the same 1..12 bound.
- On unsuppressed death, only a newly constructed Sack adopts the supplied
  `+0x7c` container (`L04607`). An existing Sack drains that container through
  `L04885 -> R0450` and deletes it at `L07774`. Contained Item identity
  follows the whole/split/merge predicates; no unconditional container-identity
  rule follows.

**Confidence.** High for the twelve worn slots: the strip loop, the equip guard
and the shipped `Slot` census are named instructions and a full-column
measurement. `SAV-HUMAN-043` capped the twelve at Medium for want of an
image-side name, and `ITEM-CORPSE-034` already carried one. High for the four
mechanisms being distinct: each is a different construct in the serializer, a
presence-flagged counted list, a fixed reference run, a presence-flagged
embedded object and a bare reference. High for the named new-Sack adoption,
existing-Sack drain and source-container deletion branches. Medium for the
thirteenth reference being the `Diary`: every save agrees, and it is non-null on
24 of 100, about one per save and consistent with the hero alone carrying one,
but no reader names `+0x1e4`.

**Amended.** `ITEM-GROUNDMOVE-130` (EXP-0285) refutes the unconditional clause
that the `+0x7c` container is the same object `ITEM-DEATH-012` turns into the
sack on death, and its consequence for a consumer that rebuilds it as a fresh
list; the branch-dependent account above replaces it
([`retracted.md`](retracted.md)). The inventory and equipment measurements and
the thirteenth-reference uncertainty are unchanged.

### SAV-DEATH-051

- Every actor reached in the original `Player` walk is alive, supporting the
  teardown unlink.
- A scan of actor-shaped bytes outside that graph is not the dead-list
  population. `SAV-DEADLOAD-124` walks the dead list's actual
  `u32 count + references` frame and finds 101 observations, stages
  `3:5, 4:71, 5:25`.
- Stage 5 and runtime id 0 coincide 25/25.
- The signed-health set is `-10001:5, -10007:5, -10011:4, -10014:4, -10017:7`.
  Complete `R0867` treats every value below -10000 identically and
  returns unchanged.

**Confidence.** High for removal from the owner graph, the exact frame and the
stage-5/id-0 coincidence, and High for the terminal consumer class. Medium for
the archival corpus distribution.

**Unknown.** The producer of each non--10001 residue (`SAV-DEADLOAD-129`).

**Amended.** The tag-scan counts, 267 records, 78 of them at a non-zero stage,
and stages `0:189,3:7,4:63,5:8`, are superseded by the exact dead-list walk and
not reused (`SAV-DEADLOAD-124`, EXP-0237, [`retracted.md`](retracted.md)).
Removal from the owner graph, signed health and the stage-5/id-0 relation stand.

### SAV-TOPLVL-052

- A consecutive walk from offset 75 with the archive's own index counter, under
  a `Player` programme without the `Diary` tail call, reads one `Player` record,
  one back-reference tag, then the `0000` null-objref word, and stops, on 23 of
  23 saves. On `game0011.sav` that is 2 038 bytes of a 57 972-byte decoded
  stream.
- The other `Player`s, every `Building`, every `Sack`, the cell records and the
  session block lie past that stop. `Building` chains (`SAV-BLDG-037`); `Player`
  does not under that programme.
- The EXP-0145 tag scan that reaches them reports 61 `Player` hits over the
  corpus, of which 2 are artefacts: on `game0003` and `game0004` it lands on a
  run of zeroes that parses as an empty `Player` whose `+0x28` reads 0,
  indistinguishable from the human participant by that field.
- `SAV-OBJ-016`'s published anchor, the slot id written twice at `+0x04` and
  `+0x08`, rejects both; with a printable-name test 59 of 61 survive.
- This refutes nothing in `SAV-STREAM-013`, whose tag scan reproduces 393/393
  tag words: a scan can attribute every object without any frame telling a walk
  where to go next.

**Confidence.** High for the chain measurement: a programme walk that tiles, run
on every file, ending on the same terminator word 23/23. High for the two
artefacts: each is re-derived by a published anchor, not by looking wrong.

**Unknown.** What the back-reference between the `Player` and the terminator
points at.

**Amended.** The headline "The stream does not chain above the first record: a
walk reaches one `Player` subtree and a terminator, and everything else needs a
pattern scan" is refuted by `SAV-PLDIARY-054` and `SAV-DOC-053` (EXP-0148,
[`retracted.md`](retracted.md)). The stop was the record programme running out,
not a frame in the file. This claim named a listing of the routine calling
`Player::Serialize` per player as the closing instrument; `SAV-DOC-053` reads
the document `Serialize` `R0414`, which writes the `Player` list through
`R1372`. The tag-scan error rate, the two zero-run artefacts and the
`SAV-OBJ-016` anchor stand.

### SAV-DOC-053

Store arm, in order:

- `u32 +0x04`, `u32 +0x00`, `CString +0x28` (the map name).
- Thirteen `u32`:
  `+0x11c +0x124 +0x128 +0x12c +0x130 +0x134 +0x138 +0x13c +0x148 +0x144 +0x140 +0x80 +0x84`.
- `R1372`: one `u32`, then `R1558`, a `u32` count and that many
  `ar << CObject*`: the `Player` list.
- At `L08331`, an indirect call through `[vtable+0]` on `*(*(this+0x14)+0xc)`, the dead-actor list
  manager (`this+0x14` points to the world and `world+0xc` to the manager). Its
  slot 0 is `R1589`, which serializes the embedded `CObList` at manager
  `+4` through `R1341` as `u32 count` followed by that many
  `ar << CObject*` actor references. The record is 4 bytes only when the
  dead-actor list is empty.
- The world-half byte (`SAV-SHAPE-023`). When it is 1: the `Building` list
  (`R1590`→`R1591`), the `SpellEffect` list
  (`R1122`→`R1109`), the terrain record (`R1360`,
  `SAV-TERRKEY-056`), 4374 raw session bytes (`R0270`, derived from the
  same pushed lengths `SAV-SESS-031` sums), and the `Sack` list
  (`R1592`→`R1593`).
- `u32 0xbadface1`, `u32 [L07886]` and `world+0x118 → R1491(ar)`,
  which reads or writes exactly 400 raw bytes.

If that logical endpoint is odd, the outer writer extends the memory stream by
one unconsumed byte before word compression; no byte is added at an even
endpoint. The 400 observed zeros are loaded state, not padding
(`SAV-TRAIL-026`).

The five top-level list serializers use one shape: `IsStoring`,
`GetHeadPosition`/`GetCount`, `u32` count, and `ar << CObject*` per node. Load
reads the count and `ar >> <class>` with the element class pushed as a
descriptor.

The `Player` list's first element begins at decoded offset 75 on 18/18, the
offset `SAV-OWNER-048` measured without the header programme.

**Confidence.** High. The enumeration is one routine read end to end. EXP-0055
independently identifies the dead-actor list through its construction,
death-path insertion, decay walk and serializer. The EXP-0245 audit attributes
every byte through padding on 30 of 31 distinct preserved saves. The remaining
digest reaches a `SpellTransport` branch absent from that audit after 49
successfully walked Unit-derived records, not inside a Unit field. EXP-0246
separately closes the class programme and nested graph.

**Unknown.** The meaning of the thirteen head `u32` past `SAV-STREAM-010`'s
four.

**Amended.** Three clauses are refuted (EXP-0245, `SAV-UNITCORP-157`,
[`retracted.md`](retracted.md)): that the programme attributes every byte on 7
of 18 saves, that the other 11 desynchronise inside a `Unit` subtree, and that
everything after the trailer dword is 400–401 bytes of the codec's word padding.
The current programme closes all eighteen distinct streams of that population.
The trailing bytes are `R1491`'s 400-byte record plus at most one
conditional alignment byte, as stated above. The top-level order stands.

### SAV-PLDIARY-054

- After the group list (`R1340` on `*(Player+0x24)`) and the 32 raw bytes
  (`R0204` on `*(Player+0x30)`), `L07961`, an indirect call through the dword at offset 0x8 of the object's table,
  dispatches `*(Player+0x40)`.
- The constructor builds it by pushing 0x30 at `L08332` and calling `R1565` at `L08333`
  and storing the constructed object into the 32-bit field at offset 0x40 at `L08334`. `R1565`'s own
  `L08335` stores L08066 into the object's first field, the `Diary` vtable
  (`SAV-MEMBER-036`'s table).
- `Diary::Serialize` `L08081` is `CDWordArray` at `+0x04`, `CWordArray` at
  `+0x18`, then `u32 +0x2c`, resolved through the identity map at load
  (`L07738`).
- Without that record a consecutive walk from offset 75 stops after one `Player`
  subtree. With it, the top-level list's five elements are five full `Player`
  records: `[75,3044] [3044,16734] [16734,17652] [17652,29546] [29546,32336]` on
  `game0001.sav (en)`.
- This refutes `SAV-TOPLVL-052`'s headline: what looked like a terminator was a
  short record programme.

**Confidence.** High. The class is reached from the instruction that builds the
object and from the constructor's literal vtable store, never from the size of
the hole it fills. The corrected programme lets 7 saves tile to the marker,
which the previous one could not do on any file.

## Seen state, terrain record and load order

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-SEEN-055 | Withdrawn: a save was said to carry no per-cell record of what the human participant has seen; `SAV-FOG-061` finds that record in the uncompressed tail. | — | ✖ retracted | [EXP-0148](../experiments/EXP-0148-save-seen-record/), [EXP-0150](../experiments/EXP-0150-save-fog-record/) |
| SAV-TERRKEY-056 | The terrain record is a count of packed block words, a 54-byte-stride cell hash map, and an inline `u32` identity key that accounts for `SAV-CELLREC-032`'s four bytes. | High | ● active | [EXP-0148](../experiments/EXP-0148-save-seen-record/) |
| SAV-LOAD-057 | A load builds the terrain from the `.alm`, ingest included, and then applies the saved block records over it. | High | ● active | [EXP-0148](../experiments/EXP-0148-save-seen-record/) |

### SAV-SEEN-055

- The record exists, in the save's uncompressed tail (`SAV-FOG-061`, EXP-0150).
  The embedded `&YA1` state store carries a section `Fog` whose `Data` leaf is a
  run-length encoding of tile bit 15 over `W × H` cells, written by
  `R0084` and read back into the live tile plane by `R0099` at
  `L07966`. On the save this claim's own corpus contains, 3475 of 20 736 cells
  are set.
- What stands is the enumeration. `SAV-DOC-053`'s reading of the document
  `Serialize` `R0414` is correct: exactly two of its ten top-level items
  are cell-keyed, and the tile plane is not written by any of them. The three
  size, shape and magnitude measurements against the compressed body stand.
- What fell is the step from there to the file. The file has two producers: the
  document writes the compressed body, and the view layer writes the tail's
  state store from a different routine at a different time. An enumeration of
  one writer bounds that writer, not the file.

**Confidence.** The High was earned on the compressed body and stated about the
file; it stands only for the enumeration of the compressed body. See
[`retracted.md`](retracted.md).

**Amended.** Retracted as a whole by EXP-0150 (`SAV-FOG-061`,
[`retracted.md`](retracted.md)). The withdrawn headline read: "A save carries no
per-cell record of what the human participant has seen, and loading restores a
fully unexplored map." The residual this claim carried, "Medium that no route
outside the document `Serialize` writes into the same file", is the clause that
was right.

### SAV-TERRKEY-056

- `R1360` writes `WriteCount(n)`, then `n` packed `u32`
  `(cellIndex << 16) | (dyn << 8) | static` over the sweep window
  `[0x807, 0x807+0xe5e7)`, for every cell whose dynamic byte is `> 0x0f`
  (`TERR-PASS-053`).
- Then the embedded object at `terrain+0x540b4`. Its vtable `L08336` is
  stored at `L08337`. Its slot 2 `R1521` is a hash map writing
  `WriteCount(count)` and, per node, 2 bytes at node+0x8 then 0x34 at node+0xc
  (`L08338` pushes 0x2, `L08339` pushes 0x34): `SAV-CELLREC-017`'s 54-byte
  stride, read from the routine.
- Then `L07962` stores `this` as an inline `u32` at the destination. The load
  arm reads it back at `L08340` and binds it in the identity map at
  `[L00285]+0x88` (`L07963`).
- The cell-record table is therefore `WriteCount + n × 54` with no tail, and the
  four unattributed bytes are the terrain's entry in `SAV-PTRMAP-035`'s scheme.
- Measured counts: block array 1 842–3 394, cell records 185–197.

**Confidence.** High. Both routines were read end to end. The terrain record's
total is what makes the whole stream tile onto the `0xbadface1` marker on 7
saves, which a length wrong by four would have moved.

### SAV-LOAD-057

- `R0414`'s load arm, in order:
  - `L08341`/`L02287` load the map named by the `CString` at `+0x28`;
  - `L08342` pushes 0xa4558, allocating 673 624 bytes, and
    `L07067`, a call to `R0118`, constructs the terrain object from that map,
    stored to `[L04624]` at `L08343`; the constructor's own tail is
    `L08344`, a call to `R0278`, the terrain ingest;
  - then `L07068`, a call to `R1360`, reads the records.
- The session object is built the same way, `L08345` pushing 0xc320, `R0077`,
  `[L00004]`, before `L07069`, a call to `R0270`.
- The save's cell records are therefore a delta over a freshly ingested map, as
  `SAV-BLOCK-012` argued without reading it. This listing is the discriminator
  `SAV-BLOCK-012` and `TERR-PASS-053` both named, and it answers
  `TERR-PASS-053`'s Unknown.

**Confidence.** High. One caller, read at instruction level, with the allocation
size and the ingest call in the same routine.

## Actor order, the starting character and Group fields

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-GRPORD-058 | Nothing orders the human participant's actor list: it is the group's `CObList` head-to-tail order, the order of appends, which is runtime state. | High / Unknown | ● active | [EXP-0149](../experiments/EXP-0149-actor-order/) |
| SAV-HERO-059 | `*(Player+0x34)` names the participant's own starting character and is in the file: it equals the named hero's identity key on 18 of 18 distinct saves. | High / Unknown | ● active | [EXP-0149](../experiments/EXP-0149-actor-order/) |
| SAV-GRPFLD-060 | Group+1c has a bounded initialization gap; +40 and +44 are serialized references, not universal constants. | High / Medium / Unknown | ● active (amended) | [EXP-0149](../experiments/EXP-0149-actor-order/), [EXP-0289](../experiments/EXP-0289-saved-group-rebuild/) |

### SAV-GRPORD-058

- `R1341`'s store arm is `L08089` GetHeadPosition, `R1369` GetCount
  written as a `u32`, then per node `R0069` GetNext and
  `R1110 ar << CObject*`: twenty-four instructions with no comparison in
  them.
- The load arm is `R0033` RemoveAll, the count, then per element `L08346`
  (`ar >> Unit*`, pushing the `Unit` descriptor `L08347`) and `R0412`
  AddTail. The file order is reproduced exactly and is the order of appends.
- `R0152`, the group's own append, first detaches the unit from
  `*(unit+0x70)` if it has a group (`L08348`, `L08349`), so a unit's
  position is destroyed and re-made whenever it changes group.
- Measured: the same five actors, by identity key, appear as one group of five
  in `game0000.sav (2026-08-02)`, five groups of one in `game0001.sav`, and
  three groups (3/1/1) in `game0007.sav`. One session gives three partitions and
  three orders, with the participant's own hero at index 0, last, and alone in
  the middle group respectively.
- A consumer must not read position as identity.

**Confidence.** High. Both arms were read end to end; the absence of a sort is a
property of one short routine, not a search; and the corpus produces three
orders of one actor set, which no stable key could.

**Unknown.** What re-partitions the groups.

### SAV-HERO-059

- It is field 16 of `Player::Serialize`'s seventeen: written at `L08350`
  through the `u32` primitive `R0685`, read at `L08351`, then resolved
  through the identity map at `[L00285]+0x88` by `L08352`, a call to `R1370`.
- In the decoded stream it is the dword 43 bytes past the `Player` name's
  `CString`, immediately before the record's own identity key.
- It equals the identity key of exactly one actor in the participant's own
  groups on 18 of 18 distinct saves. That actor is the named hero: `Danath` on
  17, `Fergard` on the one save whose participant is a different character.
- The rival "it names the last actor added" fails on `game0007.sav`, where
  `+0x34` names an actor that is neither first nor last in file order.
- The image side agrees and was not needed for the file position:
  `PARTY-INSTALL-012` reads `player+0x34 = actor` in the chargen and `0xbe`
  carry arms, and its absence from the mercenary spawn `R1142`, which is
  what makes a mercenary not a hero.

**Confidence.** High. A named store and a named load-side resolution, against
18/18, with a rival excluded by value rather than by count.

**Unknown.** More than one human participant: every save here has one.

### SAV-GRPFLD-060

- The constructor `R0149` has no direct +1c store, clears +40/+44 and
  constructs the two lists/AI block.
- Across the original 18-save population, the human participant's 29 Group
  records had +40=0; +1c was zero in 10 and had 15 distinct nonzero values
  otherwise.
- Ordinary append `R0152` writes Group+44 from the appended actor's current
  owner, and the 29 measured saved +44 values equal their enclosing Player keys.
  That is a producer and corpus agreement, not a LOAD equality law: Group LOAD
  later restores its own +44 key, and Player's later actor-owner correction does
  not recopy it (`SAV-GRPOWNER-561`).
- +1c is an authored selector on the map path and is still read on resume
  (`SAV-GRPIDENT-562`). The bounded constructor result does not prove every
  dynamic first-save value is uninitialized.

**Confidence.** High for the named append store and the measured 29-record
values. Medium for the historical uninitialized-memory explanation of dynamic
+1c.

**Unknown.** All-path initialization, and +40's consumer and gameplay meaning.

**Amended.** EXP-0289 narrows two clauses ([`retracted.md`](retracted.md),
`SAV-GRPOWNER-561`, `SAV-GRPIDENT-562`): the unconditional +44 equality ("+0x44
is not an independent fact", "so it repeats Token+0x14") and the
dynamic-initialization headline, which called a trailing dword uninitialized
memory. Producer and corpus equality did not establish restoration precedence,
and the headline exceeded its Medium scope. The 18-save observations are
unchanged.

## Fog record and the uncompressed tail

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-FOG-061 | A save carries a per-cell record of what the human participant has seen: a run-length `Fog` section in the uncompressed tail's state store, not in the compressed body. | High / Medium / Unknown | ● active | [EXP-0150](../experiments/EXP-0150-save-fog-record/) |
| SAV-TAILEXT-062 | The uncompressed tail is three regions, a `0x100`-byte label region, the `&YA1` state store ending at `poolStart + poolLen`, and a campaign record running to EOF. | High / Medium | ● active (amended) | [EXP-0150](../experiments/EXP-0150-save-fog-record/), [EXP-0223](../experiments/EXP-0223-sav-campaign-tail/) |

### SAV-FOG-061

- The embedded `&YA1` state store (`SAV-EMB-004`) holds a section `Fog` with two
  leaves: `FirstState` (kind 2, int32) and `Data` (kind 6, int32[]).
- `Data` is a run-length encoding of tile-plane bit 15 over exactly `W × H`
  cells in the plane's own linear order `idx = col + row·W` (`TERR-EDGE-024`).
  The first run carries the state in `FirstState`, and the state flips between
  runs.
- Store, `R0084`, transcribed: `L07964`, a load of the field at offset 0xc of the world half, takes the
  tile plane; `L08353`, a multiplication of the width by the field at offset 0x4 of the same object, makes the cell count `W·H`;
  `L08354`/`L08355` take `tile[0] & 0x8000`, and `L08356`, a call to `L07603`,
  writes it as `Fog/FirstState`. The loop compares the masked value (the cell word masked with 0x8000) against the
  running state (`L08357`, `L08358`), advances one `u16` per cell
  (`L08359`, advancing the cell pointer by one 2-byte cell) and appends each run length at `L08360`.
  `L07965`, a call to `L07606`, writes the array as `Fog/Data`.
- Load, `R0099`, transcribed: `L08361`, a call to `R0452`, reads
  `FirstState` into the running state; `L08362`, a call to `R1061`, reads the array;
  `L07966`, a bitwise or of the running state into the 16-bit tile word at the cell pointer, sets the state into each tile word;
  `L08363` advances the cell pointer by one 2-byte cell; `L08364` toggles bit 0x8000 of the running state
  between runs.
- The load arm only ORs, so it depends on the `.alm`-supplied plane carrying bit
  15 clear, which `TERR-FOG-087`'s census measured over 72 shipped maps.
- Corpus, 15 preserved saves:
  - `sum(runs)` is 6400 on every save of `10.alm` and 20736 on every save of
    `20.alm`, the exact extents of 12 and 13 shipped `scenario.res` maps,
    identical on both roots;
  - `FirstState` is 0 on 14 of 14;
  - run counts and set-cell counts ascend with play: 1 run / 0 cells on both
    restart saves, 29/152 and 35/248 on two saves minutes apart in one session,
    273/3475 on the longest;
  - the between-mission save carries no `Fog` section, by the store arm's own
    `L08365` test of the guard value and `L08366` branch on zero, the world-half guard.
- This refutes `SAV-SEEN-055` and the persistence clause of `TERR-FOG-087`.
  `SAV-DOC-053`'s enumeration of `R0414` is untouched and correct: it
  bounds that routine, and the file has a second producer.

**Confidence.** High. The encoding is read instruction by instruction on both
arms, and the decode reproduces a quantity it is not given, the map's exact cell
count, on 14 of 14 saves across two extents. The polarity is a field in the
file, not an assumption. The index-list and bitmap rivals are excluded by the
exact sums and by the varying element count. Medium that this is the only tile
bit a save restores: only the `Fog` section was traced into the plane.

**Unknown.** What the original draws from the restored bit. The render gate
`TERR-TILE-079` describes ORs four neighbouring corner words, so the drawn
extent may exceed the set cells.

### SAV-TAILEXT-062

- `SAV-EMB-004` reads the tail as a `0x100`-byte label region, then an `&YA1`
  running to EOF. The store's own framing does not reach EOF:
  `0x18 + 32R + 4 + poolLen` leaves a further 268 bytes on 13 saves and 310 on
  the one between-mission save.
- That region's first dword is 10 on the `10.alm` saves, 20 on the `20.alm`
  saves and 30 on the between-mission save, whose campaign head carries mission
  0 (`SAV-CITY-030`).
- The third region is the variable campaign record
  `SAV-CAMPTAIL-070`/`SAV-CAMPPROG-071` (EXP-0223). Its programme closes exactly
  at EOF on all 26 preserved owner saves. Their record-size histogram is
  `268:19, 308:1, 310:4, 368:2`, so neither 268 nor 310 is a format constant.
- The measured store is 28 records / 19 leaves on all 13 mid-mission saves and
  22 / 15 on the between-mission save: `Character/Name`,
  `CurrentState/InBattle`, `Fog/FirstState`, `Fog/Data`,
  `GameOptions/{Wimpy,ShowHP,FlyingHP,Formation,Speed,ShowTimeFlow}`,
  `Inventory/IsOpen`, `Objects/Selection`, `Projectiles/{FreeIndex,IDs}`,
  `SpellBook/{IsOpen,Pressed,Shortcuts}`, `View/{X,Y}`.
- The two sections missing from the between-mission save are `Fog` and
  `Projectiles`, consistent with the store arm's world-half guard at `L08365`.
- A consumer that parses the `&YA1` to EOF rather than to `poolStart + poolLen`
  reads 268 bytes of a different structure as pool.

**Confidence.** High for the three-region boundary and the measured leaf set:
structural identity plus the named world-half guard. High for the third region's
programme, by the paired writer/reader evidence in `SAV-CAMPPROG-071`. Medium
for the 26-file size histogram, which is one preserved owner corpus.

**Amended.** EXP-0223 closes the former Unknown on the third region's identity.
SAV-915 narrows the fixed-set implication: the complete producer also writes
nonempty `Objects/Group0..Group9` and top-level `Prj<id>` sections from live
collections. The 28/19 and 22/15 counts and the three-region boundary stand as
facts about the measured corpus, not general record/leaf constants
([`retracted.md`](retracted.md)).

## Saved Human progress

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-HEROXP-063 | A saved `Human` carries experience twice, an aggregate raw dword at `Unit+0x130` and six per-skill raw dwords at `Humanoid+0x1cc+4i`, both restored from the file and read as signed `i32`. | High | ● active (amended) | [EXP-0159](../experiments/EXP-0159-save-hero-progress/) |
| SAV-HEROSKILL-064 | The saved level vector is six `u16` at `Unit+0xa8+2i`, level slot `i` pairs with experience slot `i`, and `Unit+0xa6` is not slot 0. | High | ● active | [EXP-0159](../experiments/EXP-0159-save-hero-progress/) |
| SAV-HEROID-065 | Party progress is attached to each serialized `Human` object, not to its group, list position or primary status; the numeric `Token` key is identity within one file only. | High / Medium | ● active | [EXP-0159](../experiments/EXP-0159-save-hero-progress/) |

### SAV-HEROXP-063

- `Unit::Serialize` stores `+0x130` through the four-byte primitive at
  `L08367`..`L08368`; its load arm passes the address of the same field to
  the four-byte reader at `L08369`..`L08370`.
- `Humanoid::Serialize` calls `Unit::Serialize`, then its store arm writes
  `0x18` raw bytes from `+0x1cc` at `L08371`..`L08372`, and its load arm
  reads `0x18` directly back there at `L08373`..`L08374`: the six dwords,
  `i = 0..5`.
- The aggregate is therefore not a value the load derives after reading the six
  dwords.
- Over 306 player-subtree `Human` records in 22 distinct save byte streams, the
  stored aggregate equals the sum of the six stored dwords on 306/306.
- Controlled same-session pairs discriminate that invariant from a level-derived
  model:
  - Danath's skill vector stays `[0 11 0 0 0 1]` while experience slot 1 moves
    `1710→1745` and the aggregate `1738→1773`, with the other five slots and
    four other party members unchanged;
  - a second pair moves the same slot `1745→1762` and the aggregate `1773→1790`;
  - Naira keeps `[0 0 0 0 0 11]` across a no-world-half to world-half boundary
    while slot 5 and the aggregate both move `1648→1663`.

**Confidence.** High. Both store/load arms are read through their width-setting
calls; a programme walk, not a scan, locates 306 records; and three isolated
deltas refute both aggregate-only and per-skill-only storage. EN and RU each
decode 64/64 records with the invariant, but their relevant files are
byte-identical: root parity, not independent data variation.

**Amended.** EXP-0191 adds the signed `i32` interpretation: the byte widths
stand, and negative item-Poison awards make the signed interpretation
observable.

### SAV-HEROSKILL-064

- `Unit::Serialize` passes `this+0xa6` to `R1594` at
  `L08375`..`L08376`; that helper's two arms write and read one fixed
  24-byte block.
- `R0899` fixes the internal alignment independently: its six-iteration
  loop reads `word[+0xa8+2i]`, computes and writes `dword[+0x1cc+4i]`, then adds
  that same dword to `+0x130` (`L08377`..`L04454`).
- The award routine repeats the same subscript for skill and experience on both
  class arms (`L08071`/`L08378` and `L08074`/`L08379`), with no
  permutation.
- In the decoded player subtrees the preceding word `+0xa6` differs from skill
  slot 0 on 306/306 records, refuting the premise that the raw block starts with
  slot 0.
- The six-word level vector is independent state: the controlled Danath and
  Naira pairs change experience with the corresponding level word unchanged.
- Customisation boundary: the original record has exactly six skill words and
  exactly six experience dwords. A seventh paired slot cannot be represented in
  the 24-byte experience tail. Extending it moves every following `Humanoid`
  field, while reusing bytes in the `Unit` block destroys other serialized
  state. A versioned extension must retain the original six-slot reader.

**Confidence.** High. The block widths come from both archive arms; the internal
alignment and parallel indexing come from two independent writers; and
controlled file deltas exclude a level-derived or shifted-index interpretation.

### SAV-HEROID-065

- The complete path is document `Player` list → `Player::Serialize` → group
  actor `CObject*` list → `Humanoid::Serialize` → `Unit::Serialize`.
- On load the archive constructs the tagged object and invokes the inverse arms.
  Its `Token` head binds the saved key to that constructed object and resolves
  the saved owner key (`SAV-PTRMAP-035`), while the group list retains the
  returned object pointer.
- The owner reference equals the enclosing `Player` key (`SAV-OWNER-048`).
  `Player+0x34` additionally selects the primary character (`SAV-HERO-059`), but
  companions need no positional marker, because each is already a separate
  object in that owner graph.
- In the `2026-08-02` no-world-half/world-half pair the same five first-`Player`
  keys and owner survive while all five programme-walk positions change; one
  record's progress changes and the other four remain byte-for-byte equal.
- Naira independently moves from order 2 to 1 across the timestamped boundary
  while her own slot 5 changes, and temporary mission actors do not match her
  key.
- This refutes an order-indexed or one-party-vector restore.
- The numeric `Token` key is only file-internal identity, not a durable
  character id: later saves reuse four equal key values for objects with a
  different owner key. Same-session differential joins therefore require key,
  class, name and owner to agree. Loading one file must use the graph and mint
  unique nonzero keys, not preserve a numeric key across files.
- The archive's shared class/object index and the `Token` key map are two
  distinct identity layers that converge on the same constructed object.

**Confidence.** High for file-internal association: the established writer/load
chain is met by a five-object reorder in which no progress is reassigned. Medium
for a cross-save same-session join using the pointer-shaped key: the
class/name/owner controls discriminate the observed pairs, but key reuse proves
the number alone is not stable beyond that scope.

## Campaign record wire programme

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-CAMPTAIL-070 | The bytes immediately after the embedded `&YA1` state store are one campaign record rooted at application `+0x548`, not a fixed anonymous trailer. | High | ✔ promoted | [EXP-0223](../experiments/EXP-0223-sav-campaign-tail/) |
| SAV-CAMPPROG-071 | The campaign-record wire programme is variable, with a 104-byte minimum when every count is zero. | High / Medium | ✔ promoted | [EXP-0223](../experiments/EXP-0223-sav-campaign-tail/) |
| SAV-CAMPPOS-072 | The campaign record reconstructs current world-map position from the selected mission and one first-MapPoint flag; it stores no raw coordinate pair in this record. | High | ✔ promoted | [EXP-0223](../experiments/EXP-0223-sav-campaign-tail/) |
| SAV-CAMPMARK-073 | The 32 bytes after the document list are fixed only when the marker count is zero; each marker adds `17 + strlen` bytes. | High / Medium | ● active (partially retracted) | [EXP-0223](../experiments/EXP-0223-sav-campaign-tail/) |

### SAV-CAMPTAIL-070

- The outer writer calls the `&YA1` producer at `L08122`, keeps the same file
  object, forms `app+0x548` and calls `R1294` at `L06699`, with no seek
  between them.
- Both load paths consume the store and call the inverse `R0434` on the
  same object at `L08112` and `L08116`, again without an intervening seek.
- Computing the store end from its own `0x18 + 32R + 4 + poolLen` framing and
  starting the campaign programme there reaches EOF on 26 of 26 preserved owner
  saves.

**Confidence.** High. Two independent load entrances mirror the one writer,
exact bytes anchor all three outer sites, and the owner-corpus walk closes at
the boundary the code predicts rather than at a searched signature.

### SAV-CAMPPROG-071

- The complete programme, in order:
  - a base prefix of six `u32` and two counted `u16` arrays;
  - counted child records, each base plus one `u32`;
  - one count with two parallel `u16` arrays;
  - one counted `u32` array;
  - six counted `u16` arrays;
  - documents as a count plus `(u32 value,u32 kind)` pairs;
  - seven `u32` in offset order
    `+0x118,+0x114,+0x110,+0x11c,+0x120,+0x124,+0x128`;
  - the marker list.
- The writer and reader use the same widths, counts, offsets and order through
  their own epilogues.
- The programme consumes all bytes on 26 of 26 preserved owner saves, whose
  observed sizes are 268 (19), 308 (1), 310 (4) and 368 (2).

**Confidence.** High for the programme, order, widths and minimum: complete
inverse bodies plus exact anchors, and the corpus closes at EOF under the
programme. Medium for the size histogram, which describes one preserved owner
corpus and is not a universal population.

### SAV-CAMPPOS-072

- Before writing campaign `+0x120`, `R1294` compares live
  `view+0xf8/+0xfc` with the first MapPoint at `view+0xc4` and writes 1 only
  when both coordinates agree.
- On load, non-zero `+0x120` copies that first pair; zero passes selected
  mission `+0x118` through `R1324` and copies the indexed MapPoint.
- The record therefore persists a relation to campaign map data. The separate
  `&YA1` `View/X,Y` leaves are a different producer.

**Confidence.** High. The writer's comparison and both reader branches are read
to their stores, and the raw-coordinate rival predicts two coordinate fields
that the complete programme does not contain.

### SAV-CAMPMARK-073

- Seven `u32` are followed by a marker count. Every marker then writes
  `u32 value`, `u32 stringByteCountIncludingNUL`, that many string bytes, and
  two further `u32`: `17 + strlen` bytes.
- The inverse reader consumes the same shape, reading value, string and two
  `u32` per marker, and calls only the archive-read and string helpers.
- Picture-pointer rehydration is a separate routine, reached only from world-map
  entry, never from this reader or the campaign loader (`SAV-609`).
- All 26 preserved owner saves have marker count 0, so they exercise the 32-byte
  case and no dynamic non-zero-marker case.

**Confidence.** High for the variable wire rule: both arms, the string count's
`+1`, loop bounds and epilogues are anchored. Medium for zero markers in the
26-file corpus.

**Unknown.** A live non-zero-marker save/load witness remains open.

**Amended.** The clause that the inverse reader "rehydrates the omitted runtime
picture pointer after the loop" is retracted by `SAV-609` (EXP-0308,
[`retracted.md`](retracted.md)): its image-wide caller census finds the
rehydration routine `R1490` called only from world-map entry
`R1409`. The variable wire rule and the 26-file zero-marker corpus stand.

## Token position object

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-TOKENPOS-074 | The first twelve bytes `Token::Serialize` writes are one raw 12-byte position object: cell, packed cell, sub-cell, two omitted bytes and a raw terrain pointer. | High / Medium | ● active (amended) | [EXP-0224](../experiments/EXP-0224-sav-token-position/), [EXP-0235](../experiments/EXP-0235-sav-token-load-lifecycle/) |
| SAV-TOKENPTR-075 | Position persistence includes a raw terrain identity key, and whether load later repairs it depends on which collections receive post-load dispatch. | High / Medium | ● active (amended) | [EXP-0224](../experiments/EXP-0224-sav-token-position/), [EXP-0235](../experiments/EXP-0235-sav-token-load-lifecycle/), [EXP-0236](../experiments/EXP-0236-sav-cell-record-rebind/) |

### SAV-TOKENPOS-074

- `Token+0x10` points to a separately allocated 12-byte object whose maintained
  layout is `u8 cellX`, `u8 cellY`, `u16 ((cellY<<8)|cellX)`, `u8 subX`,
  `u8 subY`, two bytes omitted by the bounded constructor/copy family, and a raw
  terrain pointer.
- Accessors compute `fullX=cellX*256+subX` and `fullY=cellY*256+subY`; a
  separate getter returns the packed cell.
- `R1494` writes or reads all 12 bytes in one raw operation.

**Confidence.** High for the layout, coordinate laws and archive width: five
constructors, the mutators, five accessors and a literal push of 0x0c agree.
Medium for semantic non-use of `+0x06..+0x07`: the resolved family, lifecycle
paths and simple-alias census agree, but stored aliases, merged registers and
dynamic-size copies remain blind spots.

**Amended.** EXP-0235 extends the claim: `SAV-TOKENLOAD-095` supplies the
whole-image writer/consumer census for `+0x06..+0x07`.

### SAV-TOKENPTR-075

- Ordinary `Token` construction seeds position `+0x08` from the live terrain
  global `[L04624]`. Archive load overwrites it with all twelve raw saved
  bytes before terrain construction.
- `R1560` later calls `R1595`, which conditionally resolves that
  dword through the archive identity map after terrain registration: success
  writes the live terrain pointer, while zero or a miss remains unchanged.
- The complete world-load arm dispatches this method for live actors, dead
  actors and SpellEffects.
- Building and Sack vtables carry the same method, but their collections receive
  no post-load dispatch in that arm; Sacks deserialize after the actor lifecycle
  pass.

**Confidence.** High for the constructor source, the raw overwrite, the helper
semantics and the actor/dead/SpellEffect dispatch: complete routines and exact
vtables. Medium for the Building/Sack remainder: the complete load path and
known helper population contain no call, while a distinct alias writer, a later
simulation call or a running-original dereference remains untested.

**Amended.** EXP-0235 and EXP-0236 narrow the former Unknown on the key's later
lifecycle to the Building/Sack remainder graded Medium above.

## Campaign record meaning

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-CAMPAIGN-076 | Every collection in the campaign record has an owner-side meaning; none is an unnamed mission set, and two scalars remain unnamed. | High | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-077 | The serialized campaign record stores no completed-mission set: completed side missions cannot be reconstructed as a historical set from it. | High | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-078 | Main-mission progress is monotone at the campaign loader: `R0755` refuses a request below the current main mission. | High | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-079 | Campaign `+0x118` is the selected mission, not completed or available state. | High | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-080 | A configured town candidate and an offered mission are different persisted states: candidate arrays versus a latch at record `+0x18`. | High | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-081 | Town acceptance consumes the persisted candidate that produced it; enumerating or viewing an announced mission does not. | High / Medium | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-082 | A child record is a retained side-mission candidate, not a completion log: expiry and completion both erase the whole record. | High | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-083 | The persisted mercenary state is four different things: pool counts, per-type hire flags, the mission's shelf and permanent unlocks. | High / Medium | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-084 | `AddHero[]` is a consumed campaign grant, but it is not Brian's grant: Brian enters through a mission-40 map instant. | High | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-085 | Campaign documents are an append-only, pair-deduplicated collection, not mission completion state. | High / Medium | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-086 | The marker list is selected-mission presentation state, not progress or availability. | High / Medium | ● active (amended, superseded) | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-087 | The fixed `we have brian !!` save is not a mission-30 or mission-40 state: main 50, selected 41, no record or building array holding 30 or 40. | High / Medium | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |
| SAV-CAMPAIGN-088 | Brian's mission-40 handover has no identity-dedup guard; ordinary progress prevents the duplicate upstream. | High / Medium | ● active | [EXP-0232](../experiments/EXP-0232-sav-campaign-state/) |

### SAV-CAMPAIGN-076

- Base and child records carry mission, MapObject, Payment, shop bounds,
  announce latch, `AddHero[]` and `EnableMercenary[]`; children add age.
- The paired `u16` arrays are working and pristine 15-type mercenary counts,
  followed by 15 hire flags.
- The six arrays are `Mercenaries`, permanent unlocks, `InnNPC`, `InnMission`,
  `TCMission`, `ShopMission`.
- Then come documents, selected mission, one still-unnamed scalar,
  `AutoGetMission`, `LastMission`, the first-MapPoint flag, mission time, one
  still-unnamed scalar, and the marker cache.
- The wire writes `TCMission` before `ShopMission`, despite their reverse object
  offsets.

**Confidence.** High for field identity and order: the complete serializer is
joined to the registry loader and each established consumer. The two unnamed
scalars remain unnamed rather than inheriting a guess.

### SAV-CAMPAIGN-077

- Main progress is the current main record at `+0x04`; completing it advances by
  ten and replaces its registry-derived fields.
- A completed side mission is removed from the child array. Once removed, its
  completion is indistinguishable from expiry or never having been retained,
  except for state it moved elsewhere, such as documents or permanent mercenary
  unlocks.
- "Completed side missions" therefore cannot be reconstructed as a full
  historical set from this record.

**Confidence.** High. `SAV-CAMPAIGN-076` accounts for every serialized
collection, while `R1414` supplies the two completion arms; an
explicit-set rival has neither storage nor a membership consumer left in the
bounded object.

### SAV-CAMPAIGN-078

- `R0755(request)` returns 0 when `request < current`. On equality or
  `-1` it selects the current main record. Only a higher request calls the
  registry loader before writing both current and selected mission.
- A normal save at main mission 50 cannot make mission 30 or 40 current again.

**Confidence.** High. The compare, rejecting return, loader call and two stores
are one complete 31-instruction routine; the lower-load, equal-load and
higher-load alternatives are its three explicit arms.

### SAV-CAMPAIGN-079

- Accepting a main or child selects that mission; travel loads `<selected>.alm`.
- Completing a side mission ultimately calls `R0755(-1)`, selecting the
  unchanged main mission after the child is removed.
- The fixed witness therefore correctly stores selected 41 inside main record
  50, without implying that 30 or 40 is incomplete.

**Confidence.** High. Selection has accept, travel, completion and load
consumers that agree on one field, and the fixed save supplies the
discriminating main-50/selected-41 combination.

### SAV-CAMPAIGN-080

- `InnMission`, `ShopMission` and `TCMission` are registry-derived candidate
  arrays.
- `R1410` instead finds the main or child record and latches its own
  `+0x18`.
- `R1416`/`R1408` enumerate the latched main first and then
  latched children. Enumeration resets only its cursor, not the latches or
  candidate arrays.

**Confidence.** High. The three building consumers, the sole latch writer and
the complete enumerator are instruction-level paths. Conflating candidates with
latches predicts reads of the six-array block in an enumerator that reads only
record `+0x18`.

### SAV-CAMPAIGN-081

- Shop and school remove element 0 from `ShopMission` or `TCMission`; inn
  acceptance removes the matching index from both `InnMission` and `InnNPC`.
- The selected record's announce latch is then set separately.
- Saving after acceptance preserves the shortened building arrays; merely
  enumerating or viewing an announced mission does not shorten them.

**Confidence.** High for all three removal paths and their order relative to the
latch. Medium for the post-acceptance file sentence, because no new
running-original save was made in this experiment; persistence follows from the
established serializer.

### SAV-CAMPAIGN-082

- Loading a main record ages existing children and deletes one at age 2 before
  appending the new registry candidates.
- Completing a selected side mission removes that child immediately.
- Its mission, MapObject, rewards, announce latch and age are serialized
  together, so expiry and completion both erase the whole record.

**Confidence.** High. The age pass, append, completion search, destructor, move
and count decrement are bounded instruction paths, and the child writer confirms
the exact fields that disappear.

### SAV-CAMPAIGN-083

- Two 15-word arrays are the working and pristine per-type pool counts.
- The 15 dwords are per-type hire flags.
- `Mercenaries[]` is the current mission's shelf, and the permanent unlock array
  admits shelf types.
- Completing a mission appends that record's `EnableMercenary[]` to the main
  permanent list and clears the source.
- These are pool, one-mission hire, authored shelf and campaign reward state,
  not mission history.

**Confidence.** High for the type index, the writers and the completion drain.
Medium for treating the two equal arrays in an arbitrary file as working versus
pristine without the reset and live-tally paths; this claim uses those paths,
not equality alone.

### SAV-CAMPAIGN-084

- Town activation walks the current record's array, creates one companion player
  character per value, then clears the array.
- The only shipped key is mission 30 value 22, which produces Fergard or
  Reniesta by sex. Mission 40's Brian enters through a map instant instead.
- Re-entering town cannot repeat an ordinary mission-30 grant from the cleared
  saved array.

**Confidence.** High. The array walk and clear are one complete routine, and the
registry's sole key plus mission 40's separate `npc25` route discriminate the
two companions.

### SAV-CAMPAIGN-085

- Registry `AddTextDocument` and `AddPictureDocum` values append `(value,kind)`
  pairs to the same array and never clear it; kind 1 is text and kind 0 picture.
- The fixed witness holds text pairs `(1,1)` through `(4,1)`, reflecting grants
  from main missions 10 and 50 while the selected mission is 41.

**Confidence.** High for append, kind and dedup. Medium for the witness's
historical gloss, which joins registry authorship to one file rather than
observing each grant.

### SAV-CAMPAIGN-086

- Selecting a mission memoizes one marker only when its MapObject has a
  non-`"nothing"` picture.
- World-map paint and the picture-bearing rectangle gate read it, and save/load
  persist it. Its picture pointer is rehydrated by world-map entry
  `R1409`, the one caller of `R1490`, not by either SAV load
  routine or the campaign loader (`SAV-609`, `SAV-CAMPMARK-073`).
- No mission loader, completion arm, building offer consumer, announce
  enumerator or companion path reads the marker list.

**Confidence.** High. Its writer, two presentation readers and serializer are
complete, while the rival progress consumers are bounded and read other fields.
Medium for the corpus scope: all 55 decoded paths have zero markers, but some
are duplicate byte streams.

**Amended.** `SAV-609` supersedes the clause "load rehydrates its picture
pointer" ([`retracted.md`](retracted.md)). Its image-wide caller census finds
the rehydration routine `R1490` called only from world-map entry
`R1409`, which is reached through virtual dispatch and not from either
SAV load routine or the campaign loader. The persisted marker bytes still
round-trip through LOAD. Whether ordinary play reaches world-map entry after
every LOAD is `SAV-609`'s Unknown. The presentation-state classification, the
readers, the serializer and the corpus scope are unchanged.

### SAV-CAMPAIGN-087

- Its exact tail is main 50, selected 41, announced child 41 at age 1 with
  `EnableMercenary=[10]`, unannounced child 51 at age 0, and
  `InnMission=[0,50,51]`. No record or building array contains 30 or 40.
- The instruction path predicts that winning 41 appends unlock 10, removes child
  41 and selects 50, leaving child 51 and the same inn array. That continuation
  was not driven in the original.

**Confidence.** High for the exact parsed state and the transition operations.
Medium for their combined post-win prediction, because desktop safety prevented
the running-original witness; `evidence/verification.md` names the file values
that would refute it.

### SAV-CAMPAIGN-088

- Mission 40 instant 22 loops group 14 and calls `R0064` for `npc25`.
- That complete routine detaches the actor from its old group and owner index,
  rewrites `actor+0x14`, appends it to the destination actor list, creates a new
  group and inserts it, without walking or comparing the destination roster. An
  existing Brian would not stop the transfer.
- Normal play avoids that case because `SAV-CAMPAIGN-078` refuses to reload
  mission 40 after main progress reaches 50.
- Brian's saved identity remains in the Player/group graph, outside the campaign
  record.

**Confidence.** High for the absence inside the complete transfer routine and
for the upstream lower-mission guard. Medium for the forced-replay outcome,
because it was left as a falsifiable running-original prediction, not observed.

## World-half load order

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-CELLLOAD-108 | World-half load has a deserialize order and a later, different post-load dependency order; the archived sequence alone is not the dependency order. | High | ● active | [EXP-0236](../experiments/EXP-0236-sav-cell-record-rebind/) |

### SAV-CELLLOAD-108

- The complete load arm of `R0414` executes: live actors/Players, dead
  actors, actor-list rebuild, Buildings, SpellEffects, map resource, ALM terrain
  construction, terrain-global publication, the saved block/cell/terrain record,
  session, Sacks, triggers.
- Then post-load: live actors, dead actors, every cell record, session and
  SpellEffects.
- Actors, Buildings and SpellEffects therefore exist before terrain. Sacks exist
  after the saved cell bytes are overlaid but before cell object keys are
  repaired. SpellEffect post-load follows the cell pass.

**Confidence.** High. The 17 anchors lie in one complete 2,379-byte function
range with zero orphan or undisassembled bytes. Raw bytes from both identical
executable roots and the repaired listing agree, and the one-order rival
predicts calls in an order the instruction stream excludes.

## World-half cell table load

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-CELLLOAD-109 | A saved cell table overlays the cell hash constructed from the ALM; it neither replaces nor clears it. | High | ● active | [EXP-0236](../experiments/EXP-0236-sav-cell-record-rebind/) |
| SAV-CELLLOAD-110 | Ten cell-payload dwords are persisted archive identity keys that become typed live pointers on a successful lookup; a failed lookup leaves the saved value. | High / Unknown | ● active | [EXP-0236](../experiments/EXP-0236-sav-cell-record-rebind/) |
| SAV-CELLLOAD-111 | The cell payload mixes persisted baselines and a layer count, later-bound identity keys, two conditional live trigger forms and bounded residue. | High / Medium / Unknown | ● active | [EXP-0236](../experiments/EXP-0236-sav-cell-record-rebind/) |
| SAV-CELLLOAD-112 | Cell-record occupancy and `Token` terrain keys follow separate post-load mechanisms, leaving live Building/Sack cell pointers beside raw saved terrain keys. | High / Medium | ● active | [EXP-0236](../experiments/EXP-0236-sav-cell-record-rebind/) |
| SAV-CELLLOAD-113 | Every populated cell object key in 14 distinct preserved world-half saves names an exact serialized object in the same document: 2,373/2,373 resolved, 0 unresolved. | Medium | ● active | [EXP-0236](../experiments/EXP-0236-sav-cell-record-rebind/) |

### SAV-CELLLOAD-109

- Terrain construction runs first.
- For each saved entry, `R1521` zeroes 13 dwords, reads a two-byte key
  and 52 raw bytes, and looks the key up with `R1537`. It allocates
  through `R1596` only on a miss, then copies all 13 dwords over the
  matched or inserted node.
- A saved key overwrites its constructed payload, a saved-only key is inserted,
  and a construction-only node survives.
- The later `R1597` pass visits the resulting union.

**Confidence.** High. Store and load arms, the lookup branch, the conditional
allocation and both 13-dword copies are exact instructions. The
wholesale-replacement and objects-only reconstruction rivals require a clear or
an unconditional construction, and the complete routine has neither.

### SAV-CELLLOAD-110

- `R1598` visits `+0x04` (movement-domain-1/2 actor), `+0x08` (domain-3
  actor), `+0x0c` (Building), `+0x10` (Sack) and six SpellEffect layer slots
  `+0x14..+0x28`.
- Each nonzero dword is passed to the archive identity map at
  `[L00285]+0x88`. Success writes the returned live address; zero and a
  failed lookup leave the saved value unchanged.
- `R1597` performs this after Sacks deserialize, copying each node's 13
  dwords to scratch and back.
- Occupancy is therefore restored from saved keys, not reconstructed from each
  object's position, and an identity failure leaves raw residue rather than
  null.

**Confidence.** High. The five lookup sites, the six-iteration last site, the
success-only stores and the caller order are exact raw anchors. Independent
attach/detach/accessor families type all ten slots, and the 14-save join in
`SAV-CELLLOAD-113` agrees for every populated actor, Building and Sack key.

**Unknown.** Runtime failure behaviour after the preserved raw key.

### SAV-CELLLOAD-111

- `+0x00` is the captured cost baseline and `+0x01` the captured static-block
  baseline; recomputation reads both.
- `+0x02` is the occupied six-layer count, retained from the save until an
  ordinary layer attach or detach recounts it.
- The ten dwords `+0x04..+0x28` follow `SAV-CELLLOAD-110`'s identity lifecycle.
- Trigger writer `R0227` stores six bytes at `+0x2c..+0x31`. Two
  consumers read them:
  - On actor footprint attach, `R0458` requires
    `+0x2c != 0 && +0x2c != 26`, indexes the spell table by that byte and
    dispatches through `R1137` or `R1138`. `+0x2c` is the spell,
    `+0x2d` the power and `+0x2e/+0x2f` the temporary caster's source x/y; the
    target is the entering actor or its current cell.
  - Separately, `R0039` tests `+0x2c == 26` and reads relocation x/y from
    `+0x30/+0x31`.
- Only constructor-zero `+0x03` and `+0x32..+0x33` cross raw load and post-load
  unchanged with no semantic consumer in the enumerated population.

**Confidence.** High for the positive baseline, count, identity and both
conditional trigger consumers: independent writers and readers, exact stores and
the two mutually exclusive operation tests. Medium for treating `+0x03` and
`+0x32..+0x33` as ignored residue: all 28 owners of `world+0x5402c` plus the
immediate-address owner were read, with zero relevant orphan owners, but
arbitrary alias or pointer-arithmetic users outside that bounded access shape
can escape the census.

**Unknown.** Whether either restored shipped trigger is entered and completes a
spell in a running original.

### SAV-CELLLOAD-112

- `R1560` sends Position `+0x08` to `R1595`. A successful identity
  lookup writes live terrain; a miss preserves the raw saved key.
- The complete load arm invokes vtable slot `+0x24` for live actors, dead actors
  and SpellEffects, whose relevant vtables reach that helper.
- Building and Sack vtables also contain it, yet their list loaders have no
  post-load collection dispatch in the complete arm.
- Their cell slots are still repaired by the later all-cell pass. The result is
  a bounded asymmetry: live cell-record Building/Sack pointers beside raw saved
  terrain keys in their position objects.

**Confidence.** High for the helper, the vtable targets and the
actor/dead/SpellEffect dispatch. Medium for the Building/Sack remainder and the
resulting asymmetry: the full load routine and the known helper/call population
contain no dispatch, while a distinct alias write, a later simulation call and a
running-original dereference remain untested.

### SAV-CELLLOAD-113

- Field and class population: `+0x04` 208 Human and 381 Unit references; `+0x08`
  22 Unit; `+0x0c` 1,710 Building; `+0x10` 52 Sack.
- The walk emits 1,948 nonzero `Token` identity-key objects.
- The relevant position terrain dword equals that save's terrain identity on all
  208 Humans, 391 Units, 285 Buildings and 52 Sacks.
- No save has a nonzero `+0x14..+0x28` SpellEffect layer key, so that typed
  field is instruction-only, not corpus-exercised.

**Confidence.** Medium. The join is exact over 14 SHA-distinct saves from five
dated owner directories and agrees independently with the typed writers, but
corpus agreement cannot establish the later live lookup or cross-process address
use. The zero area-layer population is an explicit blind spot, not evidence that
those fields are always empty.

## Token lifecycle and the Position rebind

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-TOKENLOAD-092 | A dispatched `Token` lifecycle hook resolves Position `+0x08` as an archive identity reference; it does not install current terrain unconditionally. | High / Unknown | ● active | [EXP-0235](../experiments/EXP-0235-sav-token-load-lifecycle/) |
| SAV-TOKENLOAD-093 | The supported serialisation programme has twenty `Token`-lineage classes and six lifecycle bodies, and every body reaches the common Position rebind when invoked. | High | ● active | [EXP-0235](../experiments/EXP-0235-sav-token-load-lifecycle/) |
| SAV-TOKENLOAD-094 | World load dispatches the Token lifecycle only through three established populations; city load dispatches none. | High / Medium / Unknown | ● active | [EXP-0235](../experiments/EXP-0235-sav-token-load-lifecycle/) |
| SAV-TOKENLOAD-095 | Position `+0x06..+0x07` is raw-only in the resolved Position and Token-lifecycle population; zero is the Medium-confidence authoring placeholder. | Medium | ● active | [EXP-0235](../experiments/EXP-0235-sav-token-load-lifecycle/) |
| SAV-TOKENLOAD-096 | The two-root Token-position census holds 2,258 verified heads, of which 1,110 join to twelve completely walked world documents. | Medium | ● active | [EXP-0235](../experiments/EXP-0235-sav-token-load-lifecycle/) |

### SAV-TOKENLOAD-092

- World load reads raw roster Token records before that load arm constructs new
  terrain. Terrain load then binds its saved identity to the new terrain in
  `[L00285]+0x88`.
- `R1560` passes `Token+0x10` to `R1595`. Zero returns unchanged;
  a non-zero dword is looked up; a successful lookup overwrites `+0x08`; a miss
  returns unchanged.
- For a dispatched world object that will use terrain, a writer must put the
  terrain record's same saved identity in `+0x08`. There is no fixed pointer
  placeholder.

**Confidence.** High for the writer, its two no-write arms and the paired-key
authoring law: terrain registration, the hook and the complete 48-byte rebind
are instruction-level paths read as raw bytes in both byte-identical images.

**Unknown.** The rule as a universal placeholder for undispatched Buildings,
nested inventory/effects and the city transition. This claim does not infer
reach from vtable membership.

### SAV-TOKENLOAD-093

- The twenty classes are `Token` and nineteen descendants.
- Fourteen descriptors use `R1560` directly: Token, VirtualCaster,
  SpellEffect, Effect, Effect_DirectDamage, Building, Outpost, Tavern, Shop,
  Item, Armor, Shield, Weapon and Sack.
- Unit uses `R1482`; Humanoid and Human wrap it through `R1569`.
- PointEffect, AreaEffect and SpellTransport use `L05904`, `R1116` and
  `R1115`. Each calls the common hook first and then dispatches one or two
  nested Token references.
- Descriptor, base, creator, vtable, slot-2 serializer and slot-`0x24` target
  are recorded for all twenty.

**Confidence.** High for this structural map: all 28 schema-1 descriptors were
enumerated, base chains select exactly twenty, raw vtable entries reduce to six
complete bodies, and no class bypasses the common hook inside its body. Manager
membership and actual load-time reach are the separate bounded result
`SAV-TOKENLOAD-094`.

### SAV-TOKENLOAD-094

- After terrain identity registration, `R0414` invokes the global live
  CObList, the dead-actor CObList and the typed SpellEffect manager before Sacks
  deserialize. Their exact iterator bodies call every present element's vtable
  slot `+0x24`.
- The live-list insertion family admits Human, Unit and Sack, but loaded Sacks
  are not present for this pass.
- Building and loaded-Sack managers have no corresponding call in the complete
  document loader, and nested Item/Effect reach is not established.
- On the world-half-zero arm the loader constructs no terrain, invokes no
  lifecycle manager, zeroes three non-position actor fields and returns to the
  campaign record without a Position or terrain consumer.

**Confidence.** High for the three calls and the bounded city non-use: the
exhaustive branch and both manager vtables and bodies are instruction-level
evidence. Medium that no Building load callback exists beyond this loader and
the enumerated insertion surface; computed callbacks and later session paths are
blind spots.

**Unknown.** City-to-mission placement.

### SAV-TOKENLOAD-095

- Constructors maintain `+0x00..+0x05` and `+0x08`. Copy construction and
  assignment copy those same bytes and the pointer while omitting the middle
  word. Logical compare, distance, setters and getters omit it. Only the literal
  twelve-byte archive helper transfers it.
- A whole-image simple-alias census starts from 1,407 non-stack `object+0x10`
  seeds, rejects 2,172 stack-argument loads and finds no candidate access at
  either displacement.
- In 1,110 clean world records, 883 words are zero and 227 are non-zero, across
  89 non-zero pairs.

**Confidence.** Medium. Positive family instructions, varied corpus residue and
the bounded alias census discriminate a required-reserved-zero model, but
aliases stored and reloaded from memory, interprocedural returns, merged
registers and dynamic-size copies are blind spots. The `a55a` live-load mutation
is an unwitnessed prediction, so zero is not promoted to High.

### SAV-TOKENLOAD-096

- Clean class population: Unit 146, Human 178, Effect 43, Building 11, Item 25,
  Armor 388, Shield 18, Weapon 289 and Sack 12.
- All 1,110 have a derived packed cell. Position `+0x08` equals that file's
  terrain identity on 993, is zero on 72 and is another non-zero identity on 45.
- One SpellTransport occurs in an EN document whose variable Unit subtree
  prevents the world join. The other ten Token-lineage classes have zero corpus
  witnesses.
- The 27 file instances represent 26 distinct hashes.

**Confidence.** Medium. The programme walk and an independent tag scan must
agree before a head is counted, and clean pointer relations require the terrain
boundary. These are nevertheless finite owner-corpus populations, not
format-wide laws. Zero-count classes and the 45 other identities are not
assigned semantics by count alone.

## Dead-actor load

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-DEADLOAD-124 | The top-level dead-actor record is an exact archive list: `u32 count`, then that many actor object references. | High | ● active | [EXP-0237](../experiments/EXP-0237-sav-dead-actor-projection/) |
| SAV-DEADLOAD-125 | Every loaded dead actor receives one virtual post-load rebind, which repairs its saved terrain key and five object references without joining an authored map unit. | High | ✔ promoted | [EXP-0237](../experiments/EXP-0237-sav-dead-actor-projection/) |
| SAV-DEADLOAD-126 | Stage, signed health and signed timer are three independent stored scalars, restored directly and not derived during dead-actor load. | High | ✔ promoted | [EXP-0237](../experiments/EXP-0237-sav-dead-actor-projection/) |
| SAV-DEADLOAD-127 | The exact dead list holds 101 observations at stages 3/4/5 = 5/71/25, and every record joins one authored type-6 unit in each root. | High / Medium | ● active | [EXP-0237](../experiments/EXP-0237-sav-dead-actor-projection/) |
| SAV-DEADLOAD-128 | The enumerated actor-proven stage routes have seven causes, and one of them gives stage 5 outside ordinary decay. | High | ● active | [EXP-0237](../experiments/EXP-0237-sav-dead-actor-projection/) |
| SAV-DEADLOAD-129 | `R0867` treats every stage-5 health word below -10000 as one terminal consumer class; it does not distinguish the particular residue. | High / Unknown | ● active | [EXP-0237](../experiments/EXP-0237-sav-dead-actor-projection/) |
| SAV-DEADLOAD-130 | A dead actor stores carried-container contents, not allocation identity. | High / Medium | ✔ promoted | [EXP-0237](../experiments/EXP-0237-sav-dead-actor-projection/) |
| SAV-DEADLOAD-131 | The timer's eight actor-proven writers belong to the live dying/order tick; the named dead-load and decay routes do not consume it. | High / Unknown | ● active | [EXP-0237](../experiments/EXP-0237-sav-dead-actor-projection/) |
| SAV-DEADLOAD-132 | The complete dead-actor post-load chain has no stage-specific cell insertion. | High / Unknown | ● active | [EXP-0237](../experiments/EXP-0237-sav-dead-actor-projection/) |

### SAV-DEADLOAD-124

- Dead-manager slot 0 is `R1589`, which calls `R1341` on its
  embedded list.
- Store writes the count, then each `CObject*`. Load creates or resolves every
  element through the archive's shared class/object index and appends the
  pointer.
- The exact walk advances that shared index and reaches all 101 references in
  twelve preserved saves. No tag scan or measured skip participates.

**Confidence.** High. The complete manager table, both arms of the bounded
archive helper and the programme walk agree; every element is checked against
the record minted at its archive index.

### SAV-DEADLOAD-125

- Document load calls dead-manager slot 1 after world reconstruction;
  `R1599` walks every element and calls `actor->vt+0x24`.
- The three actor tables converge on `R1482`. It resolves position
  `+0x08` through `R1560` → `R1595`, and actor
  `+0x5c/+0x64/+0x44/+0x68/+0x40` through the archive identity map.
- No W40, runtime-id or ALM-record search occurs in this complete chain.

**Confidence.** High for the bounded virtual chain and its six rebinds: all
manager and actor vtable slots resolve, direct targets are enumerated, and raw
instructions carry every call. The absence is scoped to this chain; orphan code
and whole-object helpers are not silently excluded.

### SAV-DEADLOAD-126

- `Unit::Serialize` stores and loads timer `+0x6c` through a byte helper, health
  `+0x94` through a word helper and stage `+0x13c` through a byte helper, each
  with its own stored destination address.
- The post-load actor routine changes none of them; stage 0 only gates repair of
  two unrelated embedded objects.
- Consumers establish `i8`, `i16` and `u8` respectively.

**Confidence.** High. Both complete serializer arms and the complete post-load
routine exclude clamping, cross-field derivation and overwrite for these
destinations.

### SAV-DEADLOAD-127

- The saved Token low word at file `+0x13` matches exactly one W40 in the named
  map.
- EN and RU agree on record index, position, class pair, template, owner,
  player, W40 and UID for 101/101, with zero missing or ambiguous joins.
- Stage 5 has runtime id 0 on 25/25, and signed health
  `-10001:5, -10007:5, -10011:4, -10014:4, -10017:7`.
- These are repeated observations of six authored stage-5 units, not 25
  independent actors.
- ROM1 does not use this authored join during load (`SAV-DEADLOAD-125`).

**Confidence.** High for the exact corpus counts and one-to-one mappings,
mechanically emitted from exact cursors. Medium for treating the observed
distribution as representative beyond these twelve archival saves.

### SAV-DEADLOAD-128

- Direct stores: constructor 0, mission-end survivor reset 0, first dying tick
  1, revive 0, the decay ladder 2/3/4/5, and the spell-family stage-2-to-5 path.
  The archive load is the stored-address writer.
- Consumers: death/revive/decay, the spell arm, post-load character repair,
  mission outcome, actor message, full-state broadcast, serializer and stage-0
  post-load repair.
- The bulk broadcast includes dead actors only below stage 5.

**Confidence.** High for the enumerated direct, stored-address and virtual
routes; overlapping displacements, immediates, vtables, direct targets and
orphans are preserved in the census. No global absence is claimed past the named
whole-copy and unresolved-callback blind spots.

### SAV-DEADLOAD-129

- The complete routine reads health signed and returns immediately below -10000.
- Otherwise it may subtract one and advance stages at -10/-20/-40/-600, and it
  has exactly one terminal assignment, -10001, before id release.
- -10007, -10011, -10014 and -10017 are therefore neither alternate endpoints
  nor continued terminal decay in this routine.
- All 101 saved dead actors have zero effects.

**Confidence.** High for the entry equivalence class, the one endpoint and the
exclusion of in-routine post-terminal decay: complete raw body plus a 36-store
census.

**Unknown.** Which external health writer left each residue, and whether another
subsystem gives a residue meaning. Melee, poison, typed effects and other
writers survive as alternatives. No numeric-shape name is assigned.

### SAV-DEADLOAD-130

- The constructor creates a fresh 0x24-byte object at `actor+0x7c`. A nonzero
  presence byte loads `u32 count`, object references and two tail dwords into
  that object.
- Death had already handed the old container object to a sack and installed a
  fresh empty corpse container before stages 2 through 5.
- Load creates another fresh object and restores it in place.
- All 101 observed bodies carry presence 1, count 0 and tails `(10000,0)`.

**Confidence.** High for the allocation/transfer/load identity rule, from
complete constructor, teardown and serializer paths. Medium for the all-empty
distribution, an archival two-map corpus rather than controlled deaths.

### SAV-DEADLOAD-131

- All eight byte stores are in `R0037`: dyingTime minus one, the positive
  dying decrement and six movement/order writes.
- The archive restores the timer directly.
- Dead-manager serialization and post-load, actor rebind and `R0867`
  never read it; all 101 exact dead records carry zero.

**Confidence.** High for the eight stores and the bounded dead-load absence,
with overlapping displacement, immediate, stored-address, vtable and orphan
searches named.

**Unknown.** Uses outside those routes: a callback, whole-object copy or first
runtime tick could supply another use. The first actual post-load tick in a
running original is unwitnessed.

### SAV-DEADLOAD-132

- The chain restores position bytes, rebinds the position's terrain key and
  resolves actor object references.
- The document rebuilds the live tick list from `Player` actors and handles the
  dead list separately.
- No stage-3/4 versus stage-5 occupancy branch or insertion call occurs between
  manager dispatch and actor-rebind return.

**Confidence.** High for the bounded chain's absence, with raw instructions,
complete vtables, direct-call enumeration and orphan blind spots.

**Unknown.** Actual loaded cell occupancy: an earlier subsystem could rebuild
it, and no safe running-original load was witnessed.

## Building and Sack Position consumers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-POSLOAD-140 | The tested serializer, tick, placement and removal bodies contain no direct consumer of the Building/Sack Position terrain identity. | High / Medium / Unknown | ✔ promoted | [EXP-0244](../experiments/EXP-0244-sav-position-consumers/) |

### SAV-POSLOAD-140

- After the base Position serializer, the complete Building, Outpost, Tavern,
  Shop and Sack serializers contain 21/10/3/3/4 call instructions. None directly
  calls lifecycle `R1560`, rebind `R1595`, either terrain-clamping
  Position helper or any Building/Sack placement/removal helper. Outpost retains
  one unresolved indirect call.
- Building, Outpost, Tavern and Shop tick through the empty body `L08380`, and
  Sack through the empty body `R0940`.
- In the complete placement/removal bodies the reduced direct Position reads are
  Building `+0x00/+0x01` and Sack `+0x02`; terrain is supplied separately.
- The 19-function simple-alias reduction finds 13 direct accesses at those
  offsets and zero at `+0x06..+0x0b`, and records six escapes.
- The raw full-`.text` direct-call scan gives Sack removal two callers,
  `L08381` and `L08382`, where the targeted Ghidra project shows one.

**Confidence.** High for the five complete serializer bodies, the two complete
tick bodies, the four complete placement/removal listings and their positive
direct reads. Medium for excluding `+0x08` from the named helper paths: the
alias reducer is syntactic, six pointers escape, and the Ghidra project has
6,705 functions rather than the full repaired function population.

**Unknown.** Everything outside the bounded population, including the first
later post-load Position consumer, nested-call use of `+0x08`, resave behaviour
and the safety of an unresolved terrain key.

## Unit wire programme and tag-scan reconciliation

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-UNITPROG-156 | `Unit::Serialize` is one symmetric variable wire programme with six raw writes, not four. | High / Unknown | ✔ promoted (amended) | [EXP-0245](../experiments/EXP-0245-sav-unit-subtree-reconciliation/) |
| SAV-UNITCORP-157 | No Unit-derived subtree fails in the full current preserved SAV corpus of 55 file instances and 31 distinct SHA-256 streams. | High / Medium | ● active | [EXP-0245](../experiments/EXP-0245-sav-unit-subtree-reconciliation/), [EXP-0246](../experiments/EXP-0246-sav-token-class-serializers/) |
| SAV-TAGSCAN-158 | The `game0007.sav` 38-of-50 mismatch is a tag-scan failure, fully reconciled: the scan reads four raw Unit words as objects. | High | ● active | [EXP-0245](../experiments/EXP-0245-sav-unit-subtree-reconciliation/) |

### SAV-UNITPROG-156

- In store order after `Token`, the object list and two u16 lists, the raw calls
  are 24 from `+0xa6`, 22 from `+0xbe`, 24 from `+0x114`, 64 from `+0xd4`, 180
  from `*(+0x154)` and 148 from `*(+0x158)`. The last helper then serializes the
  pointed `+0x90` u16 list.
- The programme continues with the offset-only scalar runs, three archive
  references, `CString`, and the `+0x7c` and `+0x140` presence arms.
- The 19-byte first scalar run is bytes `+0x49`..`+0x4c`, dwords `+0x50`,
  `+0x54` and `+0x58`, plus bytes `+0x60`, `+0x61` and `+0x6c` (store
  `L08383`..`L08384`; load `L08385` and `L08386` for `+0x50` and
  `+0x54`). The per-actor order state `actor+0x50` and action field
  `actor+0x54` (`AI-AUTHOR-015`) are therefore serialized, 4 bytes each.
- Load repeats the same order and widths.
- Each load flag is masked to one byte and tested against zero, so every nonzero
  value takes the arm; store emits 0 or 1.
- `Humanoid` appends raw 24 and thirteen references.
- The six widths still sum to `SAV-UNITLEN-045`'s 462, so its arithmetic stands.

**Confidence.** High. Two independent raw-instruction listings agree on both
complete arms and every helper bound. Mutation 1 to 2 preserves the exact
endpoint, while neutralising a genuine archive tag breaks it.

**Unknown.** Meanings inside raw blocks, u16 lists and offset-only scalar runs.

**Amended.** The scalar-run bullet extends the programme with the run's offsets; the six raw writes, the widths and the 462 sum are unchanged.

### SAV-UNITCORP-157

- The audit reaches 2,300 Unit/Human record instances, 1,274 after digest
  reduction.
- Its exact document programme closes through the trailer, the terminal raw-400
  block and the conditional alignment byte on 53 of 55 instances, 30 of 31
  digests.
- Both failures are the same `game0018.sav` digest. They occur after 49
  successfully walked Unit-derived records, when a genuine archive reference
  introduces `SpellTransport`, for which this audit has no programme, at decoded
  offset 46,969.
- EXP-0246 publishes that class's full programme and exact nested graph
  (`SAV-CLASSSER-175`, `SAV-CLASSSER-177`); neither changes the observed Unit
  boundary.
- All eighteen distinct streams in EXP-0148's population close in this audit.

**Confidence.** High for this bounded population, every reached Unit boundary
and each reported endpoint: every digest, field cursor and first error is
mechanically emitted from the static programme. Medium for the observed count
ranges and any extrapolation beyond these digests; this audit alone does not
establish the remaining whole-document endpoint.

### SAV-TAGSCAN-158

- The first Player programme has 50 objects and the byte scan 44. 12 agree, 10
  exact objects have no scan offset, 28 same offsets carry the wrong scanned
  class, and four scan-only hits occur at decoded tag offsets 559, 1894, 3200
  and 4666.
- The first three of those words are `0x8004` and the fourth `0x8003`. All four
  lie inside four `*(Unit+0x154)` raw-180 spans, and no archive call consumes
  them.
- The scanner counts them as Humans, advances its invented shared counter and
  shifts later class indices.
- Inserting `0x8001` into another raw Unit span leaves the exact 30-object
  programme unchanged while changing the scan from 30 to 20; neutralising a
  genuine class tag makes the programme fail.

**Confidence.** High. Complete static field spans, every archive-call event,
exact mismatch attribution and opposite mutation outcomes exclude a missing Unit
arm and a genuine-object reading at all four sites. The scan remains a
candidate-word index, usable only when a separately proved archive boundary
confirms a hit.

## Token-subclass serializers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-CLASSSER-172 | All twenty enumerated schema-1 `.data` `Token`-lineage descriptors have published slot-2 programmes; eight, not the proposed nine, were unpublished at EXP-0246's base. | High | ● active | [EXP-0246](../experiments/EXP-0246-sav-token-class-serializers/) |
| SAV-CLASSSER-173 | Four Token-subclass bodies have fixed extents: `VirtualCaster` 44 bytes, `SpellEffect` 39, `Effect_DirectDamage` 68 and `Tavern` 81. | High / Unknown | ● active | [EXP-0246](../experiments/EXP-0246-sav-token-class-serializers/) |
| SAV-CLASSSER-174 | `PointEffect::Serialize` is `SpellEffect`, one typed object reference, then one raw identity dword with a load-side repair. | High / Unknown | ● active | [EXP-0246](../experiments/EXP-0246-sav-token-class-serializers/) |
| SAV-CLASSSER-175 | `AreaEffect` and `SpellTransport` have variable typed object-graph programmes, at least 47 and 45 bytes long. | High | ● active | [EXP-0246](../experiments/EXP-0246-sav-token-class-serializers/) |
| SAV-CLASSSER-176 | `Outpost::Serialize` ends in a counted embedded array of raw eight-byte elements; its extent is `95 + 8n`, or `99 + 8n` with the wide count. | High / Unknown | ● active | [EXP-0246](../experiments/EXP-0246-sav-token-class-serializers/) |
| SAV-CLASSSER-177 | The full permitted SAV corpus holds one distinct nested `SpellTransport` → `PointEffect` → `Effect_DirectDamage` graph and no record or raw name of five other EXP-0246 classes. | Medium | ● active | [EXP-0246](../experiments/EXP-0246-sav-token-class-serializers/) |

### SAV-CLASSSER-172

- A fresh scan finds 28 schema-1 descriptors in `.data`; base chains select
  `Token` plus nineteen descendants.
- Descriptor → creator → directly called constructor → literal vtable → slot 2
  resolves every one.
- Existing SAV rows cover eleven and `SHOP-SAVE-015` covers `Shop`, leaving
  exactly `VirtualCaster`, `SpellEffect`, `PointEffect`, `AreaEffect`,
  `SpellTransport`, `Effect_DirectDamage`, `Outpost` and `Tavern`.
- Their slot-2 targets are respectively `L08387`, `R1600`, `L08388`,
  `R1601`, `R1602`, `L08389`, `L08390` and `L08391`.

**Confidence.** High for this bounded structural population. `RepairVtables`
accounts for 1,875 executable pointer targets and 8,004 functions. The
descriptor scan has no class allow-list, retains non-Token rows and prints every
creator/constructor/vtable join. It is blind to runtime-built descriptors and to
descriptors outside `.data`, neither of which the schema-1 `.data` headline
claims against.

### SAV-CLASSSER-173

- `VirtualCaster::Serialize` `L08387` is `Token` + `u8 +0x3c` + six raw bytes
  at the buffer pointed to by `+0x40`: 44 bytes.
- `SpellEffect` `R1600` is `Token` + `u8 +0x40` + `u8 +0x41`: 39 bytes.
- `Effect_DirectDamage` `L08389` is `Effect` + 24 raw bytes from `+0x48`: 68
  bytes.
- `Tavern` `L08391` is `Building` + raw `u32 +0x9c`: 81 bytes.
- Each store/load pair visits the same base plus bytes in the same order.

**Confidence.** High. All four complete bodies are branchless past the archive
store/load split, raw listings fix every helper width, and the inherited totals
use the independently published 37-byte Token, 44-byte Effect and 77-byte
Building programmes.

**Unknown.** Field meanings inside the raw blocks.

### SAV-CLASSSER-174

- Store: `L08388` writes `ar << *(this+0x48)`, then raw `u32 this+0x44`.
- Load reads an object through descriptor `L08129` (`Effect`), then the dword
  into `+0x44`. `R1108` then looks that key up in the archive identity map:
  success replaces it with the live pointer and failure writes null.
- The body is at least 45 bytes and grows when the object reference introduces a
  class or object.

**Confidence.** High. Both arms and the complete identity helper are read. The
literal descriptor excludes an untyped-read model, and the post-read call
excludes a plain persistent scalar model.

**Unknown.** What the repaired pointer means to simulation.

### SAV-CLASSSER-175

- `AreaEffect` `R1601`: `SpellEffect`, four bytes at `+0x48..+0x4b`,
  `u16 +0x4c`, then `ar << *(+0x44)`. Load requires `Effect`. Minimum 47 bytes.
- `SpellTransport` `R1602`: `SpellEffect`, references `*(+0x44)` and
  `*(+0x48)`, then `u16 +0x4c`. Load requires `SpellEffect` then `AreaEffect`.
  Minimum 45 bytes.
- Both arms preserve that order.

**Confidence.** High. Complete bodies fix the order and primitive widths, and
the three typed readers push the exact `Effect`, `SpellEffect` and `AreaEffect`
descriptors. The archive reference grammar, not a condition in either body,
makes the extents variable.

### SAV-CLASSSER-176

- `L08390` calls `Building`, transfers raw dwords in order `+0x84`, `+0x88`,
  `+0x80`, `+0x8c`, then dispatches slot 2 of the embedded object at `+0x6c`.
- Its constructor calls `L08392`, which installs vtable `L08393`. Slot 2 is
  `L08394`: empty `CObject`, `WriteCount`/`ReadCount`, then `L08395`, whose
  `SHL 3` transfers `8n` raw bytes.
- The extent is `95 + 8n` for `n < 0xffff`, or `99 + 8n` for the wide count.

**Confidence.** High. Constructor, literal vtable, full slot table, both
serializer arms, both count arms and the element helper discriminate a fixed
raw-tail model.

**Unknown.** The embedded object's class name and element meanings; it exposes
only the `CObject` runtime descriptor.

### SAV-CLASSSER-177

- The census covers 55 paths and 31 SHA-distinct saves. EN `game0018.sav` and
  owner archive `2026-08-15/game0018.sav` are the same digest.
- In its 73,436-byte decoded body, class records begin at 46,949, 47,008 and
  47,064; bodies begin at 46,969, 47,025 and 47,089.
- Programme replay consumes 196, 136 and 68 bytes including nested bodies, and
  every nested start agrees with the independent exact class-record scan.
- `VirtualCaster`, `SpellEffect`, `AreaEffect`, `Outpost` and `Tavern` have zero
  exact records and zero raw-name hits.

**Confidence.** Medium. Exact framing, raw-name search, digest deduplication and
programme/class-record-offset agreement establish this manifest's histogram and
refute the no-instance hypothesis. Corpus agreement cannot establish future
reach or turn a zero into a format law.

## Non-Token serializers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-CONTSER-188 | All fifteen non-`Token` bodies in the static 35-descriptor population have exact framing. | High | ● active | [EXP-0247](../experiments/EXP-0247-sav-nontoken-serializers/) |
| SAV-CONTSER-189 | A load-created base `TableLine` is `CString`, then archive `Count(n)`, then `n` raw u32 values. | High / Unknown | ● active | [EXP-0247](../experiments/EXP-0247-sav-nontoken-serializers/) |
| SAV-CONTSER-190 | The remaining runtime arrays use element counts: `CStringArray = Count(n) + n CString`; `CByteArray = Count(n) + n raw bytes`. | High | ● active | [EXP-0247](../experiments/EXP-0247-sav-nontoken-serializers/) |
| SAV-CONTSER-191 | The two schema-0 string maps serialize current map enumeration, not sorted keys. | High / Unknown | ● active | [EXP-0247](../experiments/EXP-0247-sav-nontoken-serializers/) |
| SAV-CONTSER-192 | A `CDib` archive body is one BMP file image written directly through the archive's underlying file. | High | ● active | [EXP-0247](../experiments/EXP-0247-sav-nontoken-serializers/) |
| SAV-CONTSER-193 | `CMultiShopShelf`, `CMultiShopInstance` and `CMultiShopTemplate` each have an empty archive body. | High / Unknown | ● active | [EXP-0247](../experiments/EXP-0247-sav-nontoken-serializers/), [EXP-0058](../experiments/EXP-0058-shop-lifecycle/) |
| SAV-CONTSER-194 | None of the nine EXP-0247 candidate classes appears as an archive class record in the permitted corpus: zero in 55 paths and 31 distinct SAV digests. | Medium | ● active | [EXP-0247](../experiments/EXP-0247-sav-nontoken-serializers/) |

### SAV-CONTSER-188

- The population is eight schema-1 classes, `TableLine`, `Diary`, `Player`,
  `Spell`, `Spellbook` and the three `CMultiShop*`, plus the seven schema-0
  runtime classes.
- Seven programmes were already complete, including `CMultiShopTemplate`.
  EXP-0247 reads the other six non-empty bodies, establishes the Shelf and
  Instance empty bodies and independently confirms Template.
- No non-`Token` body-framing Unknown remains inside the `.data`/`.rdata`
  descriptor scan.

**Confidence.** High within the stated scan: EN and RU independently give 28
schema-1 plus seven schema-0 rows with non-null creators and 82 schema-`0xffff`
rows with null creators, and every one of the fifteen selected base-create
routes reaches a complete slot-2 body. A descriptor built at run time, outside
the two sections or with another shape stays outside the population.

### SAV-CONTSER-189

- Descriptor `L02471` creates 28 bytes. Constructor `L08396` stores vtable
  `L08397` and constructs the direct array member at `+8` with vtable
  `L08398`; slots `L08399` and `L08400` carry the two-part body.
- Ten vtables return the same `TableLine` descriptor but not the same
  serializer, so this grammar belongs to the descriptor's own creation route,
  not to every unregistered derived object that shares slot 0.

**Confidence.** High for the base route and boundary: creator, both
constructors, both vtable slots, both archive arms and the separate `4n`
transfer are complete.

**Unknown.** Whether an unregistered derived table is ever sent through generic
SAV object dispatch; no preserved class record witnesses any `TableLine`.

### SAV-CONTSER-190

- The shared `Count` is unsigned u16 below `0xffff`, otherwise u16 `0xffff` plus
  u32.
- `CDWordArray` and `CWordArray` already establish the same count followed by
  `4n` and `2n` raw bytes, which excludes a byte-length reading of the common
  prefix.

**Confidence.** High. Complete store/load arms call the same independently read
count pair; the word plus dword transfers discriminate element count from byte
count, which the byte array alone cannot. Malformed allocation and signed-loop
overflow are outside this framing claim.

### SAV-CONTSER-191

- `CMapStringToString` writes `Count(n)`, then `n` pairs of
  `CString key, CString value`.
- `CMapStringToOb` writes the same count and keys but sends each value through
  the archive object operation.
- Load reads the same sequence and inserts each entry. Object values therefore
  share the archive's class/object counter and identity map rather than forming
  inline raw pointers.

**Confidence.** High for the grammar and archive-reference behaviour: both
complete arms, bucket/chain loops, string helpers and object-reference helpers
agree.

**Unknown.** Entry meanings, and whether a SAV producer reaches either map.

### SAV-CONTSER-192

- Layout: a 14-byte file header, `bfOffBits-14` bytes of bitmap info and
  palette, then the derived pixel byte count.
- Store authors `BM`, `bfOffBits=0x36+4p` and `bfSize=bfOffBits+imageBytes`.
- Load requires `BM` and uses the offset, info header and palette rules but
  never validates `bfSize`.
- The segment grammar is symmetric; its redundant-header constraint is not.

**Confidence.** High. The archive flush/file handoff, complete file read/write
helpers, palette branches and both image-size arms discriminate inline archive
fields, element arrays and a stricter symmetric validator. No SAV class-record
witness establishes actual reach.

### SAV-CONTSER-193

- Their distinct constructor-owned vtables put the same `R0530: RET 4` in
  slot 2.
- The current `Shop` body does not dispatch stock.
- EXP-0247 adds the Shelf and Instance body proof to the already published
  Template result.

**Confidence.** High for all three bodies and the located Shop exclusion:
separate descriptor/create/vtable chains converge on the complete
one-instruction target.

**Unknown.** Another generic object producer.

### SAV-CONTSER-194

- `Spellbook`, `CDWordArray` and `CWordArray` are additional zero-class-record
  controls with proved directly embedded reach.
- Class-record absence therefore does not mean body absence.

**Confidence.** Medium. Every compressed body was decoded and every uncompressed
offset scanned for the full `ffff | schema | length | exact class name` prefix.
Positive observed classes and the three direct-member controls check the
instrument, but a finite owner corpus cannot establish format-wide absence.

**Unknown.** Actual SAV production of zero-witness classes, unless a located
enclosing serializer decides it.

## SAV producer reach

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-PRODROOT-204 | Within the materialised Ghidra direct-call-reference population, the located SAV document store reaches physical writer `R0804`. | High / Unknown | ● active | [EXP-0248](../experiments/EXP-0248-sav-producer-reach/) |
| SAV-PRODPOP-205 | Twenty of the static 35 stream-creatable descriptors have no corpus witness; the initialised image holds no second static population under the stated filter. | High / Unknown | ● active | [EXP-0248](../experiments/EXP-0248-sav-producer-reach/) |
| SAV-PRODGEN-206 | Three zero-witness classes, `Outpost`, `Shop` and `AreaEffect`, have native generic SAV producer paths. | High / Medium | ● active | [EXP-0248](../experiments/EXP-0248-sav-producer-reach/) |
| SAV-PRODDIRECT-207 | Six zero-witness bodies have proved direct SAV reach without their own class descriptor, so direct body reach does not imply a native generic route or a class record. | High / Medium | ● active | [EXP-0248](../experiments/EXP-0248-sav-producer-reach/) |
| SAV-LOADRESAVE-208 | Six zero-witness identities have a static typed-loader-return -> insertion/retention -> generic-writer route, conditional on the reader returning that identity. | High / Unknown | ● active | [EXP-0248](../experiments/EXP-0248-sav-producer-reach/) |
| SAV-PRODNOHIT-209 | The explicit 106-function producer slice locates no SAV parent for ten zero-witness classes; this is a bounded no-hit, not a universal absence. | Medium / Unknown | ● active | [EXP-0248](../experiments/EXP-0248-sav-producer-reach/) |

### SAV-PRODROOT-204

- The query returns two callers of that writer: `L08401` in `R0084` and
  `L08402` in `R0061`.
- For document serializer `R0414` it returns the store caller at `L08403`
  and the load caller at `L08404`.
- For compressor `R1376` it returns one other caller, the separate
  party/character exporter `R1375`, which has one returned caller and does
  not enter the SAV document path.
- None of the returned references has an orphan or undisassembled owner.

**Confidence.** High for the named returned edges and the store/load
distinction: raw writer and document listings, materialised-reference queries
and the separate exporter's caller/dispatch account agree.

**Unknown.** Unmaterialised image bytes and calls, a computed alias, and a
separately implemented writer or other implementation that uses neither root.

### SAV-PRODPOP-205

- Each byte-identical executable yields 28 schema-1 plus seven schema-0
  creator-backed `.data` rows and 82 schema-`0xffff` null-creator `.rdata` rows.
- The filter is the recurring layout plus a plausible size and schema.
  Every-byte section scanning finds no unaligned row and no additional recurring
  row outside those `.data`/`.rdata` populations.
- It retains one aligned false prefix at `L08405` (`usa`, size 1/schema 0,
  unmapped creator and next, zero raw pointer occurrences and zero materialised
  Ghidra refs).
- A fresh corpus selects the twenty zero-witness classes: each has zero records
  in 55 SAV paths / 31 distinct digests.
- Generic writer `R1129` gets the descriptor from computed object slot 0, so
  a run-time-built, copied, differently shaped or uninitialised compatible
  descriptor stays outside the census.

**Confidence.** High for the bounded executable and corpus populations:
independent EN/RU rows, the every-byte scan, the retained reject, positive
corpus controls plus byte-identical regeneration agree.

**Unknown.** Anything beyond the stated static filter, because the producer
consumes a computed descriptor pointer.

### SAV-PRODGEN-206

- `Outpost` has three non-load constructor → Building-manager-add pairs and
  `Shop` one. All reach the document Building list's generic store at
  `L08406`, with constructor vtables returning the exact descriptor.
- Map loader `R0128` reaches the `Shop` pair plus two `Outpost` pairs, one
  through multiplayer tail `R1603`. Session start `R0512` reaches the
  third `Outpost` pair after its non-load arm.
- Document load separately calls one of those `Outpost` routines, so it is
  dual-use rather than exclusively native.
- A branch in `R0003` constructs `AreaEffect`, stores it at
  `SpellTransport+0x44` and adds the transport to the saved SpellEffect list.
  Transport serializer `R1602` generically writes that nested pointer at
  `L08407`.
- No preserved save exercises any of the three exact identities.

**Confidence.** High for each static capability: complete
allocation/constructor/add/list/generic chains and exact vtable slots
discriminate an embedded-only model. Medium for occurrence in native play,
because branch execution and object survival to save are unwitnessed.

### SAV-PRODDIRECT-207

- `Token`, as the base of six SAV-reached descendants; its seventh located
  direct caller, `VirtualCaster`, is non-SAV.
- `Humanoid`, as the base of `Human`.
- `SpellEffect`, as the base of `SpellTransport`/`PointEffect`/`AreaEffect`.
- `Spellbook`, as a direct `Unit` member.
- `CDWordArray`/`CWordArray`, as direct `Diary` members.
- The exact typed readers for `Token`, `Humanoid` and `Spellbook` have zero
  materialised direct-call references.

**Confidence.** High for all six named direct/base chains: complete
explicit-body listings and constructor-owned member vtables fix each dispatch.
Medium for the absence of another native generic route, because unmaterialised,
computed or aliased parents stay outside closure.

### SAV-LOADRESAVE-208

- The Building-list reader expects `Building`, making `Outpost`, `Tavern` or
  `Shop` type-compatible if returned.
- The dead-list reader expects `Unit`, making `Humanoid` type-compatible if
  returned.
- The SpellEffect-list reader expects `SpellEffect`, making exact `SpellEffect`
  or derived `AreaEffect` type-compatible if returned. `SpellTransport` also has
  typed `SpellEffect`/`AreaEffect` fields.
- Each located load arm inserts or retains a returned pointer in state whose
  located store arm uses generic object dispatch.
- This does not prove that the original accepts a constructed stream, reaches a
  later save or actually resaves any identity.

**Confidence.** High for the static expected-descriptor, ancestry,
insertion/retention and located writer edges.

**Unknown.** Input acceptance, later-save reach and runtime resave.

### SAV-PRODNOHIT-209

- The ten classes: `CByteArray`, `CDib`, `CMapStringToOb`, `CMapStringToString`,
  all three `CMultiShop` classes, `CStringArray`, `TableLine` and
  `VirtualCaster`.
- Positive alternatives are present: all sixteen `CStringArray` calls belong to
  the separate `Data.bin` family, all six `TableLine` calls to table
  serializers, and native `VirtualCaster` construction inserts into a setup/map
  container rather than a document list.
- The four remaining runtime-library serializers have only materialised vtable
  refs; the three empty `CMultiShop` bodies are not dispatched by `Shop`.
- The slice contains 2,214 call rows and 118 computed calls.

**Confidence.** Medium. Mandatory function boundaries, every call row,
materialised Ghidra target/reference queries and positive non-SAV destinations
discriminate missing analysis from a different archive user.

**Unknown.** Unmaterialised, computed, aliased, run-time-built and separately
implemented SAV reach.

## Post-load session entry

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-POSTLOAD-220 | World resume and city-to-mission reach different pre-entry routes, and serialized `Player+0x3d` is the retained-actor placement latch. | High / Medium / Unknown | ✔ promoted (partially retracted, amended) | [EXP-0249](../experiments/EXP-0249-sav-postload-lifecycle/), [EXP-0254](../experiments/EXP-0254-sav-reconstruction-sufficiency/) |
| SAV-POSTLOAD-221 | The mandatory session-entry Position consumer is class-specific: Sack emits full X/Y, while Building emits cell X/Y and discards the sub-cell bytes. | High / Medium / Unknown | ✔ promoted (partially retracted, superseded) | [EXP-0249](../experiments/EXP-0249-sav-postload-lifecycle/), [EXP-0254](../experiments/EXP-0254-sav-reconstruction-sufficiency/) |
| SAV-POSTLOAD-222 | Session entry consumes actor Item references only after actor, recipient, type and ownership gates; nested Effect reach has two further Item gates. | High / Medium / Unknown | ✔ promoted | [EXP-0249](../experiments/EXP-0249-sav-postload-lifecycle/) |
| SAV-POSTLOAD-223 | Automatic Sack entry omits `Sack+0x40`; pickup moves quantity-one Items into the existing actor container and deletes the source container inside the move helper. | High / Medium / Unknown | ✔ promoted | [EXP-0249](../experiments/EXP-0249-sav-postload-lifecycle/) |

### SAV-POSTLOAD-220

- After `R1305`, resume calls `R0099(1)` while the two city UI arms call
  `R0099(0)`; only zero admits fresh mission construction `R0512`.
- Before session entry, nonzero campaign `+0x6b8` calls `R0147 -> R0193`,
  an ordinary data-dependent simulation tick.
- The later entry join `R0131` has one direct caller. It tests `Player+0x3d`
  only when `packet+0x0a == 0`, calls retained-actor placement only for zero,
  then writes 1 on both arms.
- Mission teardown clears the latch on the surviving human-Player branch.
- Corpus: all 52 world-half paths / 29 digests have first-Player latch 1; all
  three no-world-half paths / two digests have latch 0. This agreement is not a
  format law.

**Confidence.** High for the complete callers, the branch, the serialized byte,
the one-caller join and the write/reset sites. Medium for the corrected
finite-corpus distribution.

**Unknown.** The absolute first Position/reference event inside the pre-entry
computed tick, and a running-original transition.

**Amended.** The corpus clause is partially retracted
([`retracted.md`](retracted.md), EXP-0254, `SAV-SUFF-300`). It had read 50
world-half paths / 28 digests at latch 1 and five no-world-half paths / three
digests split between latch 0 and latch 1, refuting "city save always means
latch 0". The two city/latch-1 paths were duplicate copies of one world-half
`SpellTransport` digest, which the incomplete older walker restored to its
default no-world classification after stopping. The corrected distribution above
replaces it, so that refutation is itself refuted. The static route and the
placement-latch programme stand.

### SAV-POSTLOAD-221

- `L03114` walks the Building manager and `R1016` the Sack manager; both
  call `R0562`.
- Getter `R0165` reads Position `+0x00/+0x04` and returns `(cellX<<8)|subX`;
  `R0166` reads `+0x01/+0x05` and returns `(cellY<<8)|subY`.
- The Sack arm stores both returned words. The Building arm shifts each word
  right by eight and stores the low byte, so only `+0x00/+0x01` influence its packet,
  while `+0x04/+0x05` are read and discarded.
- Neither arm writes Position first or reads packed cell `+0x02/+0x03`,
  unmanaged `+0x06/+0x07` or terrain `+0x08`.
- A prior `R0147` tick can still consume, replace or destroy state, so this
  is mandatory-prefix order, not absolute-first order.
- All three no-world-half paths / two digests have top-level Building/Sack
  counts 0/0. Entry packets do not distinguish loaded occupancy from repaired or
  placed occupancy.

**Confidence.** High for both manager loops, the resolved class arms, the getter
arithmetic and the distinct Sack/Building stores. Medium for the corrected
3-path/2-digest city corpus.

**Unknown.** Earlier simulation, actual occupancy and offsets outside the
bounded prefix.

**Amended.** The no-world corpus count is superseded by EXP-0254
(`SAV-SUFF-300`; [`retracted.md`](retracted.md)). "All five no-world-half paths"
and "the 5-path/3-digest city corpus" become the three paths / two distinct
digests above; the two removed rows are one `SpellTransport`-bearing world
digest. The 0/0 Building/Sack observation and the sender result stand.

### SAV-POSTLOAD-222

- Mask `-1` admits equipment bit `0x80` and carried-container bit `0x800000`,
  but it does not make either walker unconditional.
- Equipment walker `R0669` gates the concrete-recipient arm on actor vtable
  `+0x30`, then branches on type, recipient and ownership.
- `R0059` calls carried-container walker `R0670` only when `actor+0x14`
  equals the recipient, and `R0670` admits only type IDs `[0x21,0x40)`.
- Every Item pointer actually reached then calls `R0883`. That routine
  dispatches vtable `+0x54` only when appearance `Item+0x40` is nonzero. Plain
  Item target `R0995` also bypasses `R0998` when `Item+0x1c == -1`; Armor,
  Shield and Weapon lack that second guard.
- Over 53 complete paths / 30 digests, actor ancestry contains 3,013 Items and
  146 effect-bearing Items/Effect references, with zero appearance-zero and zero
  plain-Item-minus-one blockers. Programme ancestry does not measure passage
  through the actor/recipient gates.

**Confidence.** High for all named gates, reached-Item consumption and Effect
traversal. Medium for the ancestry counts only.

**Unknown.** The runtime gate distribution, the pre-entry tick and future
Item-guard distributions.

### SAV-POSTLOAD-223

- `R1016 -> R0562` sends Sack Position/state without walking the
  container.
- Pickup `R0448` first enumerates its Items and writes each Item `+0x08 = 1`,
  then calls `R0450(actor+0x7c, Sack+0x40)`.
- That helper repeatedly extracts quantity one through `R0095(source,0,1)`
  and adds the returned Item through `R0929` to the existing destination.
- A stack split calls Item vtable `+0x40`. The four Item-class targets invoke
  copy constructors whose base copies the `Item+0x20` Effect list, so Effect
  references can be consumed and deep-copied before destination insertion.
- After draining, `R0450` cleans and deleting-destructs the source container.
- Only after it returns does the caller null `Sack+0x40`, remove and destroy the
  Sack and call the actor sender with mask `0xa08000`. That sender keeps
  `SAV-POSTLOAD-222`'s actor/recipient gates.
- The complete corpus has 218 Sack-ancestry Items and four effect-bearing Items
  summed over paths; ancestry does not measure split or sender-gate reach.

**Confidence.** High for the automatic omission, quantity-one extraction and
addition, split-copy and source-delete order. Medium for the 53-path/30-digest
ancestry counts.

**Unknown.** Pickup timing, stack-split occurrence and absolute first runtime
use.

## Slack regions and extension survival

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-LABELTAIL-236 | Source bytes after the first NUL in the fixed 256-byte SAV label do not retain provenance through the located load-to-save path. | High / Unknown | ● active | [EXP-0250](../experiments/EXP-0250-sav-extension-survival/) |
| SAV-PHYSUFFIX-237 | A physical suffix after the parsed campaign programme lies outside both located main-load read programmes and is discarded by the located fresh-save producer. | High / Unknown | ● active | [EXP-0250](../experiments/EXP-0250-sav-extension-survival/) |
| SAV-DECPAD-238 | The decoded document transport has only a conditional one-byte alignment gap, not a framed extension region. | High / Medium / Unknown | ● active | [EXP-0250](../experiments/EXP-0250-sav-extension-survival/) |
| SAV-EXTSURV-239 | None of the three independently bounded SAV slack candidates supplies a provenance-preserving 16-byte route through the located original writer. | High / Unknown | ● active | [EXP-0250](../experiments/EXP-0250-sav-extension-survival/) |

### SAV-LABELTAIL-236

- Both SAV dialogs read all 256 bytes into a local buffer, but their
  list/selection route passes a NUL-terminated string to `R1604`, which
  copies only through the first NUL into application label storage.
- The main writer later emits 256 bytes from that application buffer.
- The corpus has 235..254 post-NUL bytes per path and nonzero residue in 38/55
  paths, 469 bytes total. Residue is not evidence that the source tail survived:
  a fresh output can contain unrelated old bytes after its NUL.

**Confidence.** High for the complete bounded dialog-selection-writer provenance
path and the fixed capacities, which discriminate direct source-tail copying
from stale destination residue.

**Unknown.** Original acceptance of a mutated tail, and a distinct route outside
the two located dialogs and the main writer. No runtime input was sent.

### SAV-PHYSUFFIX-237

- Campaign reader `R0434` returns at its grammar endpoint; both complete
  outer load bodies then make no file read or EOF comparison.
- The writer opens with mode `0x1011`, whose helper selects `CreateFileA`
  disposition 2, writes the fixed label, state store and campaign, then closes.
- All 55 corpus paths already end exactly at campaign EOF.
- A source suffix therefore has no provenance route into a newly created output.

**Confidence.** High for non-consumption and discard on these complete static
paths: the reader endpoints, post-reader bodies, create/truncate mapping and
writer endpoint rule out misparsing, old-destination retention and source-suffix
copying within the named population.

**Unknown.** Actual original acceptance of a suffix, and malformed/error or
other loader paths; runtime was not witnessed.

### SAV-DECPAD-238

- Across 55 paths the gap after the exact logical trailer is 0 on 37 and 1 on
  18; its ten observed byte values reject a single fixed constant.
- The container reader decompresses the declared word stream, calls the document
  reader and cleans up without comparing its cursor to decoded extent.
- The writer independently extends an odd logical length by one byte before word
  compression.
- One byte cannot carry EXP-0250's 16-byte frame, and no
  source-pad-to-output-pad copy occurs in the located routines.

**Confidence.** High for the 0/1 capacity, the reader/writer mechanics and the
inability to carry the frame, established by complete paths and exact executable
anchors. Medium for the finite-corpus distribution.

**Unknown.** Original mutation acceptance, and whether allocator reuse can make
a runtime output byte coincide with its input; alternating fresh-process trials
are unwitnessed.

### SAV-EXTSURV-239

- The label tail and the physical suffix have room for the frame and are
  statically omitted by the semantic load path. The label tail loses source
  provenance at the NUL copy, and the suffix at create/truncate plus the
  campaign endpoint.
- The decoded gap is at most one byte.
- This rules out durable round-trip storage only for these regions and paths. It
  does not prove original acceptance or exclude semantically ignored fields,
  internal raw padding, runtime-built serializers, multiplayer-only shapes or
  another producer.

**Confidence.** High for the bounded negative: reader and writer boundaries
independently discriminate preservation from residue and destination retention
for all three candidates.

**Unknown.** Running-original load compatibility and every carrier outside this
population; no global absence is claimed.

## Independently generated SAV subsets

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-WRITER-284 | An independently generated no-world SAV subset has deterministic structural and checked-graph evidence, but not all-declared-value or running-original equivalence. | Medium | ● active | [EXP-0253](../experiments/EXP-0253-sav-original-writer-acceptance/) |
| SAV-WORLDWRITER-316 | An independently generated world-present SAV subset has deterministic structural and checked-graph evidence, but not original acceptance or semantic equality. | Medium | ● active | [EXP-0255](../experiments/EXP-0255-sav-world-writer-acceptance/) |

### SAV-WRITER-284

- `tools/savauthor` emits two complete files without copying corpus programme
  bytes: a 1,328-byte minimal graph and a 1,460-byte representative graph with
  four new-class records, one second object of an existing class, two plain-u16
  backreferences and five unique identity keys. Two generations are
  byte-identical.
- Its separate decoder consumes every outer, decoded, state-store and campaign
  byte. The pre-existing `tools/savdoc` independently closes both document
  programmes and every emitted class anchor.
- The equality oracle covers only the enumerated map/shape/label, selected
  Player/Human/graph, state-tree shape and campaign-extent subset.
- Eighteen exact mutations distinguish header/codec, new/existing class,
  backreference, identity, optional-arm, nested-count and tail failures.
- A confidence-review skill mutation from 1 to 99 and a campaign-scalar mutation
  from 0 to 1 still close both readers, refuting all-declared-value equivalence.
- Unverified: Human skills/stats/XP and runtime/class/type values, several
  Player/Group/Diary scalars, the Spell identity value, non-name state values,
  seven campaign scalars and consumed raw/offset-only bytes.
- The result is full structural consumption plus checked-subset equality only
  for the campaign half with `world-half = 0`. Each Human still has 450
  raw-block bytes with unknown ROM1 meanings.

**Confidence.** Medium. Gap-free byte rules, two separately implemented walks,
deterministic regeneration and 18 mutations exclude a copied skeleton and
accidental structural tiling for the exercised graph. The review counterexamples
preserve unchecked scalar alternatives, both readers still derive from the same
promoted ROM1 grammar, and no original process accepted or resaved the generated
bytes.

**Unknown.** The present world half, original loader acceptance and original
fresh-save survival.

### SAV-WORLDWRITER-316

- `tools/savauthor -world` emits complete 1,510-byte minimal and 1,850-byte
  representative files without copying corpus programme bytes. Two generations
  are byte-identical.
- The representative's archive-operator boundary:
  - seven new-class/first-object records, for Player, Human, Item, Spell,
    Building, Sack and Effect;
  - two prior-class/new-object Item records;
  - two plain backreferences from the first two equipment slots to the first two
    Human-container Items;
  - nine unique file-local identity keys.
- It includes one Building; one Sack containing an Item containing an Effect;
  one packed terrain block; one 54-byte cell overlay whose actor/Building/Sack
  keys resolve to the named objects; one terrain key; 4,374 session bytes; a
  25-record state store with a 6,400-cell Fog run; and the 104-byte campaign
  minimum with selected mission 10.
- The separate decoder consumes every byte and exactly compares the enumerated
  named edges, collections and checked values. The pre-existing `tools/savdoc`
  independently closes both documents and reaches all seven emitted class
  families.
- Twenty-two exact mutations distinguish framing, codec, archive
  tags/keys/backreferences, map/world selectors, collection/count/order,
  terrain/cell joins, Fog and campaign boundaries.
- An unchecked session-byte mutation, a third equipment alias and a reminted
  Spell identity all pass the bounded equality oracle, refuting whole-graph and
  all-declared-value equivalence.
- Unchecked: eleven world-head dwords; Player/Group/Diary offset-only or raw
  fields; Human runtime/class/type, skills, stats, XP, six raw Unit blocks and
  Position `+0x06..+0x07`; the remaining eleven equipment/reference slots and
  the Spell identity; Building's fixed 40-byte body; Sack scalar/tails;
  Item/Effect scalar bodies; all 4,374 session values; the final raw 400 bytes;
  non-name application state; campaign meanings beyond the selected-mission
  join; nonempty dead actors and SpellEffects; multiple Players and AI-owned or
  rare graphs; and external ALM/Data/campaign projection.

**Confidence.** Medium. Gap-free byte rules, deterministic regeneration, two
independently implemented walks, exact comparison of the enumerated named edges,
collections and values, and matched mutations exclude a copied skeleton and
accidental tiling for the exercised world shape. The surviving session,
third-alias and Spell-identity alternatives, the shared promoted grammar,
external reconstruction and the absent running-original witness cap the result
at structural checked-subset A.

**Unknown.** Original loader acceptance, visible loaded state and fresh-save
survival.

## Original acceptance of generated SAVs

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-ORIGLOAD-332 | The original EN process does not accept either selected independently generated positive fixture through a successful load transition. | High / Medium / Unknown | ✔ promoted | [EXP-0257](../experiments/EXP-0257-sav-original-acceptance/) |
| SAV-ORIGCHOOSER-333 | For the exact tested neighbours, chooser admission separates valid generated outer framing from invalid outer magic. | Medium / Unknown | ✔ promoted | [EXP-0257](../experiments/EXP-0257-sav-original-acceptance/) |
| SAV-SAVEDIR-334 | Launching a copied `rom.exe` with a disposable working directory does not isolate SAV discovery when the installed registry value exists. | High / Unknown | ✔ promoted | [EXP-0257](../experiments/EXP-0257-sav-original-acceptance/) |
| SAV-ORIGFAULT-335 | Both generated positives fail at the same first located incompatibility: they encode `SpellBook/Shortcuts` with zero bytes, while the original unconditionally consumes sixteen. | High | ✔ promoted | [EXP-0257](../experiments/EXP-0257-sav-original-acceptance/), [EXP-0273](../experiments/EXP-0273-quick-slots/) |

### SAV-ORIGLOAD-332

- The 1,460-byte representative no-world fixture and the 1,850-byte
  representative world-present fixture were each visible and selectable in
  `Load Game`. Each fresh disposable process then failed after the owner pressed
  Load.
- Windows recorded two Application Error events naming the exact disposable
  executable, both exception `0xc0000005` at image offset `0x0007822d`.
- The retained event metadata does not bind one record ID to one fixture, so
  that pairing is not claimed.
- This refutes acceptance only for these two exact programmes. It does not
  identify all incompatible fields, refute their research-reader structural
  closure, or establish RU behaviour.

**Confidence.** High for two exact same-boundary process failures from Windows
event metadata. Medium for fixture-to-failure attribution, from owner
observation and staged hashes without a committed screen capture or per-event
fixture binding.

**Unknown.** Post-load state, original resave, smaller fixtures and RU.

### SAV-ORIGCHOOSER-333

- Both positive files appeared and were selectable.
- The 1,850-byte world control, whose only declared role here is invalid outer
  `Asg&` magic, did not appear while staged as `game9000.sav`.
- This is a chooser-level result, not deep-loader rejection: the control could
  not be selected, and one mutation does not establish a complete chooser
  grammar.

**Confidence.** Medium. Owner-observed enumeration over three exact staged
digests, with no committed screenshot.

**Unknown.** Other malformed neighbours, the exact internal chooser predicate
and RU.

### SAV-SAVEDIR-334

- Startup `R0326` reads 32-bit HKLM `SOFTWARE\1C\Allods` value
  `INSTALLDIR` and on success calls `SetCurrentDirectoryA(INSTALLDIR)` at
  `L07681`.
- The machine value named the preserved GOG install, reproducing the owner's
  unexpected many-save chooser.
- The controlled run required an isolated registry value naming the disposable
  root, restored afterwards.
- Executable location and caller-supplied CWD alone are therefore not the save
  namespace.

**Confidence.** High for the read instruction path, the exact registry value and
the reproduced chooser consequence.

**Unknown.** Other compatibility layers and absent-key startup.

### SAV-ORIGFAULT-335

- Fault offset `0x0007822d` maps to VA `L08125` in `R0099`: the first
  the load from the source pointer plus the running offset in a four-dword copy immediately after
  `R1061("SpellBook","Shortcuts")`. The destination is the controller at
  campaign `+0xec`, offsets `+0x64..+0x70`.
- `R1061` sizes its returned array as record `Size >> 2`. Both exact
  generated state stores declare this kind-6 record with `Size=0`, producing
  zero elements and no data pointer, which the caller dereferences at the first
  unconditional copy.
- This establishes the direct failed-consumer mechanism, not repair sufficiency:
  supplying sixteen bytes may expose a later incompatibility.
- The four values are F5–F8 signed spell indices, with `-1` unbound
  (`AI-QUICKSAVE-281`).

**Confidence.** High for both event offsets, the instruction/callee mechanism
and the two exact fixture records. The claim does not certify corrected-fixture
acceptance, later boundaries or RU. The meaning of the four dwords comes from
`AI-QUICKSAVE-281`, which closes the Unknown this claim left on it.

## Full-document reader

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-FULLREAD-252 | One call-site-programmed reader consumes the complete frozen lawful SAV population from byte 0 through physical EOF without a tag or trailer search. | High / Medium / Unknown | ✔ promoted | [EXP-0251](../experiments/EXP-0251-sav-full-reader/) |
| SAV-ARCHREL-253 | Over 31 distinct documents, the reached archive calls reconstruct 4,296 objects under one shared class/object index state, with zero genuine back-reference events. | High / Medium / Unknown | ✔ promoted | [EXP-0251](../experiments/EXP-0251-sav-full-reader/) |
| SAV-IDCENSUS-254 | The pointer-map census, separate from CArchive identity, classifies 9,382 references over 31 distinct documents with zero duplicate nonzero definition keys. | High / Medium / Unknown | ✔ promoted | [EXP-0251](../experiments/EXP-0251-sav-full-reader/) |
| SAV-READPOP-255 | Both byte-identical executables join all 35 static creator-backed descriptors to the reader's programme; 19 identities lack a witness and 16 bodies are unexecuted. | High / Medium / Unknown | ✔ promoted | [EXP-0251](../experiments/EXP-0251-sav-full-reader/) |

### SAV-FULLREAD-252

- It closes 55/55 paths and 31/31 distinct digests: two world-half-absent and 29
  world-half-present documents, 21 decoded streams with no alignment byte and
  ten with one.
- Every distinct digest has zero decoded and physical interval gaps or overlaps.
- A tag-shaped raw mutation changes no archive event, and a second
  trailer-shaped dword inside raw object bytes changes no programme endpoint.
  Count, tag, schema, truncation and suffix poisons fail at the exact crossed
  field.
- Separate load-arm fixtures preserve framing when world byte 1 becomes nonzero
  2, when a non-`badface1` trailer omits the conditional scalar, in the Unicode
  CString arm and in the Unit-container-absent arm.
- Corpus Unit containers are 0/1,274 absent/present; embedded Spellbooks are
  1,254/20 absent/present.

**Confidence.** High for the named instruction-derived boundaries and the
scan-versus-programme discriminators. Medium for whole-reader closure, because
the accepted and produced language can exceed a finite corpus.

**Unknown.** Run-time-built descriptors, other loaders, malformed-input
acceptance and original production of the synthetic-only arms.

### SAV-ARCHREL-253

- Events: 346 new-class+first-object, 3,950 prior-class new-object, 8,939 null
  and zero genuine back-reference.
- A separately implemented state replay rereads all 13,235 call offsets and
  agrees on every transition.
- Low nonzero tags resolve only prior objects; a synthetic unallocated low tag
  is rejected rather than treated as a length or forward reference.
- The zero count is a corpus fact, not a ban on the statically present
  back-reference arm.

**Confidence.** High for the reached shared-index transition programme and the
exact replay, which discriminate separate counters and permissive unresolved
references. Medium for the finite event census.

**Unknown.** Original production and acceptance of a genuine back-reference.

### SAV-IDCENSUS-254

- Over the full document, 4,280 definitions and 9,382 references classify as
  5,722 uniquely resolved, 2,963 zero and 697 nonzero unresolved.
- All 697 unresolved observations are Position terrain keys, retained as
  unresolved under their exact delayed-preserve load mode. The reader does not
  guess a target from coordinates or record order.

**Confidence.** High for the bounded relation modes and the duplicate-key
rejection, which come from the reached load programmes rather than key
coincidence. Medium for the finite census.

**Unknown.** Later runtime repair and the meaning of every unresolved key.

### SAV-READPOP-255

- Each image joins the 35 creator-backed static schema-0/1 descriptors to the
  reader's exact body programme and expected serializer.
- Identity-root coverage is only 15 archive-object classes plus embedded-only
  `Spellbook`, leaving 19 descriptor identities without either witness.
- Programme execution is distinct: `Token`, `Humanoid` and `SpellEffect` run as
  bases of witnessed derived objects, leaving 16 body programmes entirely
  unexecuted.
- The wider descriptor-shape census is 117 per image: 7 schema-0, 28 schema-1
  and 82 schema-`ffff`.
- Static programme presence or base execution does not establish exact-class SAV
  production, input acceptance or resave reach.

**Confidence.** High for the bounded descriptor -> creator/runtime-class/vtable
-> serializer joins in both images. Medium for the finite identity/programme
witness split.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Unknown.** Run-time-built, copied, aliased, differently shaped or
out-of-section descriptors; actual use of the 19 unwitnessed identities;
execution of the 16 unwitnessed bodies.

## World reconstruction provenance

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-RECON-268 | The located resumable-world programme is a field-specific hybrid, not a SAV-only restore or an external-only reinitialization. | High / Medium / Unknown | ● active (amended, partially retracted) | [EXP-0252](../experiments/EXP-0252-sav-world-reconstruction/), [EXP-0287](../experiments/EXP-0287-actor-definition-binding/) |

### SAV-RECON-268

- Direct clocks, Player/actor/dead-object state, cell occupancy, session latches
  and campaign collections survive from the SAV.
- The saved map name selects an external map/ALM whose terrain and cell hash are
  constructed before the saved block/cell overlays.
- SAV session bytes load before the map trigger builder, whose rebuilt logic
  later consumes the resulting slots and latches.
- Item/Armor/Shield/Weapon definition pointers are re-derived from
  class-specific Data tables using the saved row index.
- Actor binding is class-specific (`SAV-ACTORBIND-544`): Unit uses its saved
  row, exact Humanoid clears the pointer, and Human uses its saved row below
  type33 or literal Humans row5 otherwise.
- Campaign MapPoint position is resolved relationally against external campaign
  data.
- City-shaped documents omit the world half and take fresh mission construction.
- No one provenance rule covers all owners.

**Confidence.** High for the bounded static routes and their discriminating
order: complete reader, initializer and rebind paths independently exclude both
pure models. Medium for the 55-path / 31-digest shape and object populations.

**Unknown.** Running-original equivalence.

**Amended.** The clause that Human definition pointers, like
Item/Armor/Shield/Weapon ones, are re-derived from the saved row index is
retracted ([`retracted.md`](retracted.md), EXP-0287, `SAV-ACTORBIND-544`). A
Human with saved type 33..65535 selects literal Humans row 5 without rewriting
its saved row; Unit uses the saved Units row; exact Humanoid leaves null after
its inherited Unit suffix. The field-specific hybrid model and the other owners
stand.

## World reconstruction and SAV sufficiency

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-RECON-269 | World reconstruction has non-interchangeable graph, terrain and late identity stages, and saved cell state overlays the external baseline rather than replacing it. | High / Medium / Unknown | ● active | [EXP-0252](../experiments/EXP-0252-sav-world-reconstruction/) |
| SAV-RECON-270 | SAV-only sufficiency for a resumable live world is refuted, but sufficiency of SAV plus matched external inputs is not established. | High / Unknown | ● active | [EXP-0252](../experiments/EXP-0252-sav-world-reconstruction/) |
| SAV-SUFF-300 | The exact current corpus holds 52 world-half paths / 29 distinct digests and three no-world-half paths / two digests, not the formerly published 50/5 path split. | High / Medium / Unknown | ● active | [EXP-0254](../experiments/EXP-0254-sav-reconstruction-sufficiency/) |
| SAV-SUFF-301 | A differing `.res` container digest does not imply that the SAV-selected external payload differs: the six corpus-selected maps are byte-identical across roots. | High / Medium / Unknown | ● active | [EXP-0254](../experiments/EXP-0254-sav-reconstruction-sufficiency/) |
| SAV-SUFF-302 | Across the located original routes, persisted bytes alone and external inputs alone are each insufficient; independent route information is not established. | High / Medium / Unknown | ● active | [EXP-0254](../experiments/EXP-0254-sav-reconstruction-sufficiency/) |
| SAV-SUFF-303 | No observationally sufficient full-world reconstruction vector is established. | High / Unknown | ● active | [EXP-0254](../experiments/EXP-0254-sav-reconstruction-sufficiency/) |

### SAV-RECON-269

- Archive graph reading constructs tagged objects and binds file-local keys.
- The named map/ALM then constructs terrain and the cell hash and binds saved
  terrain identity.
- Saved cell records overwrite matching keys or insert missing keys without
  clearing construction-only nodes.
- After session and Sack load plus trigger construction, the all-cell pass
  resolves ten typed occupancy/layer keys on successful lookup.
- Live/dead actor and SpellEffect post-load hooks repair their own bounded
  references at separate points.
- Numeric key value and list order are therefore transport, not durable
  cross-file identity.
- Corpus control resolves 2,373/2,373 populated actor/Building/Sack cell keys on
  14 paths. No nonzero SpellEffect-layer key exercises that arm.

**Confidence.** High for the complete positive ordering, overlay branches and
typed fixups. Medium for the finite exact-key join.

**Unknown.** Failed-lookup downstream behaviour, nonzero SpellEffect layers and
live first use.

### SAV-RECON-270

- Mandatory ALM/map construction, trigger compilation, definition binding and
  campaign relational lookup consume state absent from the SAV.
- The wider candidate, SAV plus matched external inputs, remains open:
  - a data-dependent ordinary simulation tick can precede the later proved entry
    consumers;
  - eleven head dwords, session spans/scalars, six Unit raw blocks, Position
    `+0x06..+0x07`, cell residue and the final raw 400-byte world block have
    proved framing but unresolved material meaning or first use.
- The two executables are byte-identical, while the EN/RU `main.res`,
  `scenario.res` and `world.res` pairs have different full-file digests.
  Executable equality alone therefore cannot prove root-independent derived
  state.

**Confidence.** High for the mandatory external dependencies and the enumerated
static sufficiency frontier. No corpus invariant is promoted to necessity.

**Unknown.** A causal-minimal sufficient vector, the root-specific relevant
projection and running-original equivalence.

### SAV-SUFF-300

- Every first `Player` in the exact world population has placement latch 1;
  every first `Player` in the no-world population has latch 0.
- Both false no-world/latch-1 rows were copies of one exact world document
  containing `SpellTransport`. The older partial walker had recorded the early
  Player fields, then stopped in that later class and restored the tentative
  world discriminator to its default false value.
- A fresh run of the byte-0-to-EOF reader reproduces all fourteen frozen
  EXP-0251 outputs byte-for-byte before this join.
- This corrects only the finite population, the counts
  [`retracted.md`](retracted.md) withdraws from `SAV-POSTLOAD-220`,
  `SAV-POSTLOAD-221` and `TRIG-SAVE-008`. It does not make latch value a format
  law.

**Confidence.** High for the exact discriminator, the failure mechanism and
fresh-reader equality. Medium for the 55-path finite distribution.

**Unknown.** Other producers and running-original transitions.

### SAV-SUFF-301

- The EN/RU `scenario.res` containers differ, but each of the six map members
  selected by the current SAV corpus (`10.alm`, `20.alm`, `30.alm`, `31.alm`,
  `40.alm`, `41.alm`) is byte-identical across roots.
- The measured `globalmap.reg`, `npc.reg`, `scenario.reg`,
  `world.res/data/map.reg` and `world.res/data/ai.reg` members are also equal,
  while `world.res/data/data.bin` differs.
- These payload hashes narrow the external-input candidate set. They do not
  establish equal runtime-derived state, complete semantic equality, or
  compatibility of crossed roots.

**Confidence.** High for the read-only member extraction, sizes and SHA-256
comparisons. Medium for relevance, because only located/named candidates were
inventoried.

**Unknown.** Runtime projections and unmeasured members.

### SAV-SUFF-302

- The 21-row candidate matrix independently rechecks eleven exact route anchors
  and sixteen source-evidence digests.
- SAV-only omits mandatory ALM construction, definition binding, trigger
  rebuilding and campaign lookup.
- External-only omits consumed graph, clocks, occupancy, session,
  inventory/effect and campaign state.
- I18/I19 locate distinct world-resume and no-world fresh-construction controls,
  but no crossing holds SAV bytes, external inputs and `world_half` fixed while
  varying route alone. A hybrid that reads SAV `world_half` and conditionally
  selects those controls therefore remains live.
- Complete SAV bytes, selected external payloads, constructor defaults, staged
  identity repairs and conditional pre-entry work are reached information
  classes, not a causal-minimal set.

**Confidence.** High for the two pure-model refutations and the located control
edges. Medium for finite corpus discriminators.

**Unknown.** Independent route information, causal minimality and routes outside
the searched anchors.

### SAV-SUFF-303

- The byte-complete fitting candidate includes the original's unresolved
  computed pre-entry tick and first normal tick, so it is circular rather than
  an independently executable reconstruction.
- A distinct convergence model fits the same static and corpus evidence:
  pre-tick graph, occupancy, trigger, clock or effect differences may converge
  only after the first normal tick.
- No lane-owned original window, pre/post-tick projection or crossed
  SAV/external/route runtime result was available.
- A differing declared projection after one controlled normal tick refutes the
  candidate. Equality already before that tick refutes the required-convergence
  alternative.
- Until one of those observations and the opaque consumer closure exists, a
  full-world implementation story is blocked. Bounded framing or subsystem work
  is a separate decision.

**Confidence.** High for the located conditional tick edge and the explicit
evidence gap.

**Unknown.** Pre-tick targets, post-tick convergence and joint sufficiency.

## First-tick convergence

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-FIRSTTICK-348 | The optional pre-entry transition and the normal paced loop call the same ordinary tick wrapper, so an executed pre-entry arm is one ordinary sub-tick, not a distinct load-only repair pass. | High / Unknown | ● active | [EXP-0258](../experiments/EXP-0258-sav-first-tick-convergence/) |
| SAV-FIRSTTICK-349 | The named direct Building and Sack tick virtuals cannot themselves change loaded state, and the entry projection does not traverse a Sack container. | High / Unknown | ● active | [EXP-0258](../experiments/EXP-0258-sav-first-tick-convergence/) |
| SAV-FIRSTTICK-350 | No value-level first-tick convergence, persistence, replacement, consumption or destruction result is established for the controlled Building, Sack or Sack-nested Item/Effect case. | High / Medium / Unknown | ● active | [EXP-0258](../experiments/EXP-0258-sav-first-tick-convergence/) |

### SAV-FIRSTTICK-348

- The load route tests campaign `+0x6b8` at `L08408`. Zero branches around the
  call; nonzero calls `R0147` at `L03780`.
- The normal paced loop calls that same wrapper at `L02626`.
- The wrapper returns without work when `server+0x2c == 0`. Otherwise its
  unconditional `L01863` call reaches `R0193`, which:
  - increments `server+0x04`;
  - drains queued commands;
  - dispatches every `world+0x2c` entry through `vt+0x18`;
  - walks the global ticking list through a second `vt+0x18` dispatcher;
  - publishes the incremented counter.
- On the executed arm, P1 is therefore after one ordinary sub-tick, and the
  first normal paced call is the next wrapper invocation. On the skipped arm no
  tick instruction lies between P0 and P1.

**Confidence.** High for the condition, the two callers and the complete
wrapper/sub-tick bodies.

**Unknown.** Receiver values, and whether either wrapper returns at
`server+0x2c == 0` in a loaded original.

### SAV-FIRSTTICK-349

- Building/Outpost/Tavern/Shop `vt+0x18 = L08380` and Sack
  `vt+0x18 = R0940` are complete no-op bodies.
- Later mandatory entry walks the Building and Sack managers. The Building
  sender reads Position cell/sub-cell and emits only cell coordinates; the Sack
  sender emits full coordinates; neither writes Position.
- The Sack sender omits `Sack+0x40` and therefore its nested Item/Effect graph.
- This bounds the named class virtual and entry-sender targets only. It does not
  establish P0=P1 or P1=P2, because the same ordinary sub-tick first drains
  commands and invokes two computed receiver populations whose loaded instances
  were not observed (`SAV-FIRSTTICK-348`).

**Confidence.** High for both complete no-op bodies and the already closed
manager senders.

**Unknown.** Computed receiver membership, indirect mutation and live graph
values at P0/P1/P2.

### SAV-FIRSTTICK-350

- Scratch regeneration produced a structurally closed 1,850-byte case with one
  Building and one Sack containing one Item containing one Effect. An
  independent reader reached all four selected class records.
- No original process or owned original window existed, no input was sent and no
  generated file entered an install.
- Static closure excludes direct mutation only in the two no-op class virtuals
  and excludes nested traversal only in the Sack entry sender
  (`SAV-FIRSTTICK-349`).
- Command handlers, the `world+0x2c` virtual population and the global
  ticking-list virtual population remain capable of acting before either
  observed boundary. No live population or P0/P1/P2 projection was acquired.

**Confidence.** High for the explicit static and runtime boundary. Medium for
the generated checked structural case.

**Unknown.** Dynamic state and convergence.

## Rare-fixture boundaries

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-BOUNDARY-364 | The 55-path / 31-digest preserved population supplies a multi-spell Spellbook, a nonempty top-level SpellEffect graph and dead lists, but no full twelve-slot equipment record or 256-wide terrain fixture. | High / Unknown | ✔ promoted | [EXP-0259](../experiments/EXP-0259-sav-boundary-fixtures/) |
| SAV-SPELLBK-365 | The preserved `game0021.sav` gives an exact high-count Spellbook boundary: stored count 29 followed by 28 consecutive `Spell` references. | High / Unknown | ✔ promoted | [EXP-0259](../experiments/EXP-0259-sav-boundary-fixtures/) |
| SAV-EFFECTGRAPH-366 | The preserved `game0018.sav` exercises one nonempty top-level SpellEffect graph. | High / Unknown | ✔ promoted | [EXP-0259](../experiments/EXP-0259-sav-boundary-fixtures/) |
| SAV-BOUNDARY-367 | Research-reader closure is not original load/resave evidence for the rare fixtures. | High / Unknown | ✔ promoted | [EXP-0259](../experiments/EXP-0259-sav-boundary-fixtures/) |

### SAV-BOUNDARY-364

- The actor census reaches 1,109 actor rows. The maximum populated
  worn-reference count is 7/12 and the maximum Spellbook count is 29.
- Exact documents reach a maximum dead-list count of 22, a maximum terrain block
  count of 4,117 and a maximum terrain cell count of 197. Selected map names
  remain six-byte mission names.
- The committed provenance/digest, exact-document and 1,109-row actor tables
  define the finite population. An OUT-honouring probe derives every maximum
  without receiving expected values.

**Confidence.** High for these mechanically reproduced finite maxima.

**Unknown.** Whether the absent requested cases are obtainable or representable
in other lawful saves.

### SAV-SPELLBK-365

- Slots 1 through 28 occupy decoded offsets 1,117..1,425.
- Each is an eleven-byte prior-class/new-object archive record tagged `0x8009`,
  and together they advance object indices 13 through 40.
- The complete document closes at EOF.

**Confidence.** High for count, tiling, class identity and endpoint, from the
exact archive event stream.

**Unknown.** Original load and resave.

### SAV-EFFECTGRAPH-366

- The root record is a new-class `SpellTransport`. Its first typed child is a
  new-class `PointEffect`, that child's typed effect is a new-class
  `Effect_DirectDamage`, and the transport's typed `AreaEffect` reference is
  null.
- Their nested decoded interval is 46,949..47,165, and the exact reader closes
  the whole document.

**Confidence.** High for the observed archive classes, nesting, null arm and
endpoints.

**Unknown.** Base `SpellEffect`, non-null `AreaEffect`, original load and
resave.

### SAV-BOUNDARY-367

The selected Spellbook, SpellEffect and 22-dead-actor documents close under the
exact byte-0-to-EOF parser. No lane-owned original process or foreground window
existed, so no original acceptance, loaded-state projection or fresh resave was
measured.

**Confidence.** High for exact parser closure and the bounded absence of a
runtime witness.

**Unknown.** Original behavior.

## Independent writer audit and original-load trials

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-WRITERAUDIT-380 | A complete claim-to-emission audit refutes completeness of the independent city/world writer and separates three claim-determined repairs from one claim-insufficient state surface. | High / Unknown | ✔ promoted | [EXP-0260](../experiments/EXP-0260-sav-next-load-boundary/) |
| SAV-ORIGWRITER-396 | An independently serialized no-world document reaches original EN load, stable shop use and a distinct resave. | High / Unknown | ✔ promoted | [EXP-0261](../experiments/EXP-0261-sav-minimal-original-load/) |
| SAV-ORIGMIN-397 | The smallest A/B/C-successful topology tested under the promoted no-world grammar is one Player, one Group, one Human and exactly two CArchive objects. | High / Unknown | ✔ promoted | [EXP-0261](../experiments/EXP-0261-sav-minimal-original-load/) |
| SAV-ORIGCANON-398 | An original resave preserves the accepted semantic graph while reminting file-local identities and canonicalizing selected application state. | High / Unknown | ✔ promoted | [EXP-0261](../experiments/EXP-0261-sav-minimal-original-load/) |
| SAV-ORIGVALUE-399 | Complete format closure and an original-produced resave do not establish safe Human values: a projected one-Human candidate loads, then fails on the shop route. | High / Medium / Unknown | ✔ promoted | [EXP-0261](../experiments/EXP-0261-sav-minimal-original-load/) |
| SAV-ORIGMISSION-400 | The one-Human minimum crosses from city `31.alm` into mission `30.alm` and produces complete world-present saves, but the transition is not script-silent. | High / Medium / Unknown | ✔ promoted | [EXP-0261](../experiments/EXP-0261-sav-minimal-original-load/) |

### SAV-WRITERAUDIT-380

- The 29-row matrix accounts for container, document, object, identity, world
  and tail requirements.
- The legacy fixtures used a zero-byte `SpellBook/Shortcuts`, admitted only one
  hard-coded Spell, and paired one world cell record with a block whose static
  bit 5 was clear.
- The audit writer accepts an explicit 16-byte Shortcuts payload and a vector
  that emits the witnessed count-29 Spellbook as 28 distinct records, and
  rejects any unequal static-bit-5/cell key set.
- Two generations of the corrected city/world candidates are byte-identical at
  1,997 / 2,383 bytes and preserve seven equipment aliases.
- The world candidate still has 25 state records and omits
  `Projectiles/FreeIndex` and `Projectiles/IDs`. Existing claims require 28
  records/19 leaves and name those leaves but do not publish their wire kinds or
  values.
- Opaque head, object, session, trailer, application-state and campaign values
  remain in the same explicit Unknown rather than receiving a claimed-safe zero.

**Confidence.** High for the exact legacy emissions, the three corrections,
deterministic hashes and matrix coverage from pinned repository inputs.

**Unknown.** Original acceptance, semantic value safety, Projectiles wire
details, loaded state and fresh resave.

### SAV-ORIGWRITER-396

- The writer parses named fields and graph relations from one lawful city
  witness. It then independently emits the Player/Group/Human and reachable
  object programmes, rebuilds the CArchive registry and identity relations,
  semantically serializes the 22-record state tree and campaign programme, and
  rebuilds the word codec. It copies no source tags, object records, physical
  offsets or compressed packets.
- The complete 25-object semantic candidate is 3,220 bytes (`df8dad3e…`).
- The owner loaded it with both heroes visible, opened the same shop, changed
  equipment and wrote a distinct 3,221-byte output (`66f2508a…`).
- Both complete readers close candidate and output with zero gaps or overlaps.

**Confidence.** High for the exact generated EN candidate, the bounded
owner-observed load/shop/save route and the complete resave.

**Unknown.** RU, arbitrary typed values, world-present authoring and broader
gameplay.

### SAV-ORIGMIN-397

- `primary-hero-core.sav` is 1,864 bytes (`4335c377…`). It retains Danath as
  `Player+0x34` and a present-empty Unit container, and has no Item, Effect,
  Spellbook, equipment or world half.
- The owner loaded it, observed Danath, opened the same shop and wrote
  `game0004.sav`, a distinct 1,860-byte output (`2dac9d3e…`) that retains the
  same one-Player/one-Group/one-Human graph.
- Both readers close both 2,104-byte decoded documents with 22 state records and
  zero gaps or overlaps.
- This is a minimum over the preregistered coherent topology reductions, not
  over malformed inputs or arbitrary compressed encodings.

**Confidence.** High for the exact topology, the bounded EN A/B/C observation
and the complete output.

**Unknown.** Byte minimality, malformed tolerance, RU and missions requiring
more party members.

### SAV-ORIGCANON-398

- The final generated/resaved pair retains two objects, one Player/Group/Human,
  all checked Danath values plus byte-identical 268-byte campaign records.
- Player/Human keys change `02c1ed70/04c56b88` to `02b4b100/059da938`.
- `Inventory/IsOpen` changes 1 to 0, while fourteen other checked state leaves
  remain semantically equal.
- The earlier 25-object and three-object accepted pairs likewise retain topology
  under reminted identities.
- Numeric identity keys are therefore not durable state; their resolved
  relations are.

**Confidence.** High for the three exact EN pairs and the enumerated
comparisons.

**Unknown.** Unenumerated state leaves, different routes and identity behaviour
in other producers.

### SAV-ORIGVALUE-399

- The 1,527-byte projected one-Human candidate and the original's distinct
  1,540-byte resave both load and close completely.
- The generated candidate hangs on the shop route. After the original-produced
  resave is loaded in a fresh process, that route crashes.
- The independently serialized source-faithful 25-object control, then its
  three-object and two-object reductions, open that shop and resave.
- Original identity reminting and state canonicalization are therefore not
  sufficient corrections. One or more invented Human values or retained
  relations in the projected candidate remain unsafe; the experiment does not
  isolate which field.

**Confidence.** High for the repeated bounded failure and the successful
controls. Medium for localizing the cause to the multi-field semantic
difference.

**Unknown.** The exact unsafe field.

### SAV-ORIGMISSION-400

- The automatic 28,339-byte output (`d8e6f1d2…`) and the explicit 29,152-byte
  output (`06799a42…`) close with zero gaps or overlaps.
- The owner transcribed the nonfatal messages `instant14 removecure2` and
  `cant resolve hero 10002`. `MISSION-CURE-026` identifies instant 14 as the
  item-removal action aimed at second-hero ordinal 10002, which the one-Human
  roster lacks.
- The mission still loaded and the explicit save succeeded. Format acceptance
  therefore does not imply that every mission's authored party references
  resolve.

**Confidence.** High for the owner-observed continuation and the exact world
outputs. Medium for the diagnostic-to-absent-hero causal join.

**Unknown.** Extended mission play and second-hero-free campaign semantics.

## Live-field differences between the projected and accepted programmes

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-LIVEFIELD-412 | The failed projected and accepted source-faithful two-object city programmes differ outside Human/Unit, refuting a Human-only difference model for this pair. | High / Unknown | ✔ promoted | [EXP-0262](../experiments/EXP-0262-sav-live-field-authoring/) |
| SAV-LIVEPROD-413 | The accepted programme's differing values do not yet have a complete no-SAV producer set: 35 of its 65 differences have no producer located within the searched bounds. | High / Unknown | ✔ promoted | [EXP-0262](../experiments/EXP-0262-sav-live-field-authoring/) |
| SAV-LIVEREADER-414 | Four reciprocal Player/Group/Diary versus Human/Unit diagnostic candidates are deterministic complete no-world documents, but reader closure establishes no original outcome. | High / Unknown | ✔ promoted | [EXP-0262](../experiments/EXP-0262-sav-live-field-authoring/) |
| SAV-LIVEPGD-415 | One exact accepted Player/Group/Diary-half trial did not reproduce the accepted EN shop result in the projected-base transplant. | High / Unknown | ✔ promoted | [EXP-0262](../experiments/EXP-0262-sav-live-field-authoring/) |

### SAV-LIVEFIELD-412

- A complete census accounts for 196 named values, relations, identities and
  physical consequences: 131 agree and 65 differ.
- The differing population is 50 Human, 11 Player, one Group and three state
  rows. By classification it is 59 logical values, one resolved graph relation,
  one external terrain relation, two file-local identities and two
  serializer-derived pool offsets.
- The embedded Player Diary alone changes from empty dword/word arrays and a
  null reference to 119 dwords, 119 words and a reference resolving to the
  enclosing Player.
- This is a difference census, not evidence that any Player, Diary or Human row
  causes the original shop outcome.

**Confidence.** High for the deterministic complete census of the two exact
programmes and the normalized graph relations.

**Unknown.** The original consumer, the causal field and other producers.

### SAV-LIVEPROD-413

- The accepted Human class key selects installed Humans row 26, independently
  named `PC_Danath`. The EN and RU `world.res:data/data.bin` rows agree on all
  26 measured parameters.
- Row-by-row producer accounting of the 65 differences (`SAV-LIVEFIELD-412`): 26
  have a named installed or runtime producer chain, four are
  serializer/identity-derived and 35 are Unknown within the searched definitions
  and published producer routes.
- The installed Body, Reaction, Mind, Spirit and HealthMax initializers are 40,
  36, 25, 17 and 50; the accepted current stats are 41, 35, 20, 15 and 131. The
  inventory names no route producing those five accepted values.
- Exact accepted values whose producer remains Unknown cannot be promoted from
  the source SAV into an independent writer.

**Confidence.** High for the bounded dual-root definition extraction, the exact
row counts and the claim-linked inventory.

**Unknown.** Producer routes outside the named bounds, and every causal
requirement.

### SAV-LIVEREADER-414

- Two fresh generations agree byte-for-byte at 1,657, 1,714, 1,714 and 1,661
  bytes.
- Each candidate retains base-local Player/Human identities, remaps root
  relations and differs from its base at exactly the preregistered 11-row or
  52-row partition.
- `savfull` closes each with one Player, two archive objects, 22 state records,
  zero gaps and zero overlaps. Independent `savdoc` closes the same documents
  with two agreeing objects and no mismatch row.
- The values are diagnostic transplants from the compared programmes, not
  independently produced authoring values.
- One candidate has a separate original-runtime result in `SAV-LIVEPGD-415`. The
  other three retain Unknown chooser, load, shop and resave outcomes.

**Confidence.** High for deterministic generation, the exact semantic partitions
and two-reader closure.

**Unknown.** Original acceptance of the three untested candidates, and safe
authoring.

### SAV-LIVEPGD-415

- The 1,657-byte `p-plus-pgd.sav` (`a5f2dc53…`) was hash-matched before the
  bounded original trial.
- Its chooser row appeared and the original reached the city and shop route, but
  the process crashed while opening the shop.
- The requested 30-second city stability and visible-Human subchecks were not
  separately reported, and resave was not attempted.
- This one observation does not distinguish PGD-half insufficiency from the
  preregistered hidden-runtime-state, candidate-identity or route-variance
  model, so PGD-half sufficiency remains Unknown.
- It also does not show that PGD is unnecessary, that Human/Unit is sufficient,
  that one field is causal, or that any SAV-derived value is safe to author.
- The three reciprocal candidates (`SAV-LIVEREADER-414`) and RU remain untested.

**Confidence.** High for the exact candidate partition, the chooser observation
and the reported C-boundary process crash.

**Unknown.** PGD-half sufficiency, repeatability, the omitted B subchecks,
reciprocal causality, RU and safe producers.

## Projectiles state and the independent world writer

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-PROJSTORE-428 | A world-present ROM1 save producer always emits a canonical `Projectiles` state subtree, even when the live projectile collection is empty. | High | ✔ promoted | [EXP-0263](../experiments/EXP-0263-sav-world-authoring/) |
| SAV-PROJLOAD-429 | The matching loader is stricter about wire type than about record presence: the whole `Projectiles` subtree is loader-optional. | High / Unknown | ✔ promoted | [EXP-0263](../experiments/EXP-0263-sav-world-authoring/) |
| SAV-PROJCORP-430 | The lawful 55-path / 31-digest SAV population agrees with the producer shape but scarcely exercises its value relation: one document has a nonempty `IDs`. | High / Medium | ✔ promoted | [EXP-0263](../experiments/EXP-0263-sav-world-authoring/) |
| SAV-WORLDSTATE-431 | An independent structural writer can close the canonical empty-Projectiles state tree, but only at level A. | High / Unknown | ✔ promoted | [EXP-0263](../experiments/EXP-0263-sav-world-authoring/) |
| SAV-WORLDFRONT-432 | Closing `Projectiles` does not yield an implementation-authorable independent world writer: the Human graph remains a synthetic-value boundary and several world values still lack producers. | High / Unknown | ✔ promoted | [EXP-0263](../experiments/EXP-0263-sav-world-authoring/) |
| SAV-WORLDDIAG-433 | One source-faithful world vector can be independently encoded to complete level A without replaying source records or compressed packets. | High / Unknown | ✔ promoted | [EXP-0263](../experiments/EXP-0263-sav-world-authoring/) |

### SAV-PROJSTORE-428

- The complete save arm writes `Projectiles/FreeIndex` through the YA1 integer
  setter from the manager word at `world+0xa0c`, so its wire kind is 2.
- It walks each live manager node, appends the node's u16 id, and writes a
  decimal `Prj<id>` section with sixteen kind-2 leaves:
  `x y z picture dir phase lastaction action actiondir actiontarget actionx actiony actionz actionphase actionsegments actionspell`.
- It then writes `Projectiles/IDs` through the YA1 integer-array setter. That
  setter assigns kind 6, allocates four bytes per source word and zero-extends
  each u16 to one u32.
- An empty collection therefore still produces `FreeIndex` and an empty `IDs`:
  28 records under nine state roots, not the earlier independent writer's 25
  (`SAV-WRITERAUDIT-380`).

**Confidence.** High. The world gate, manager fields, complete collection loop,
all twenty string references and both typed setter bodies are instruction-pinned
in the byte-identical EN/RU executable. The finite corpus independently exhibits
the same kinds and names (`SAV-PROJCORP-430`).

### SAV-PROJLOAD-429

- The world loader constructs an empty word vector, calls the integer getter
  with default 0 for `FreeIndex` and narrows the result into `world+0xa0c`. It
  then calls the integer-array getter for `IDs`.
- A missing integer leaf returns the supplied default. A missing array leaf
  returns 0 without changing the already-empty vector, and the non-positive
  count branches around every projectile allocation.
- Canonical kind 6 is read as `size/4` elements, taking the low u16 of each
  four-byte element. Kind 2 has a one-element compatibility arm; other kinds
  reach the typed error path.
- For each present id the loader allocates a 0x14c-byte projectile, reads the
  same sixteen leaves using the constructed object's fields as defaults, binds
  world, terrain and id, links a manager node, then calls `R0614`.
- Absence therefore does not explain an earlier load rejection. A nonempty
  vector reaches object construction and insertion before later ordinary play.

**Confidence.** High for the complete direct loader/helper paths and the first
reached consumer.

**Unknown.** Malformed error presentation, later projectile tick behaviour and
running-original mutations.

### SAV-PROJCORP-430

- All 29 distinct world-state documents have `Projectiles/FreeIndex` kind 2 and
  `Projectiles/IDs` kind 6. `FreeIndex` takes 13 distinct values.
- Twenty-eight have empty IDs.
- The sole nonempty case is `en/game0018.sav`, SHA-256
  `1e2eb21f47082ab05a8f7810148fdf722fe8aa4f69a57d79953bd8fe0025eb6b`. IDs
  contains one four-byte element whose value is 266, exactly one `Prj266`
  section carries all sixteen named leaves, and `FreeIndex` is 267.
- The upper sixteen bits of every observed array element are zero.

**Confidence.** High for the exhaustive deduplicated state census, the kinds,
the counts and the exact singleton. Medium for `FreeIndex = max(IDs)+1`, because
one nonempty document cannot discriminate that allocator relation from a
coincidence or a wider rule.

### SAV-WORLDSTATE-431

- Two fresh complete runs produced the same 1,606-byte source-independent
  document, SHA-256
  `a559e01a4c177b975a6fb191783b151ab0a24ada7edc021d9b5aca79ed69d448`, with 28
  state records, kind-2 `FreeIndex=0` and an empty kind-6 `IDs`.
- Its 1,510-byte comparator, SHA-256
  `61f0a2624d0b22a8d73399688414b8e3cc09b356af96c104efcabd7a81dae179`, retains
  the same declared surrounding programme but omits the subtree and closes at 25
  records.
- Complete local parsing and the original loader path (`SAV-PROJLOAD-429`)
  therefore refute the model that those three missing records alone force early
  load rejection.
- Neither file was submitted to ROM1, because both still carry unresolved
  independently invented non-state values.

**Confidence.** High for deterministic generation, the exact hashes, no-source
construction and complete local parser closure.

**Unknown.** Original chooser, load, first action and resave.

### SAV-WORLDFRONT-432

- The first already-observed synthetic-value boundary is the Human graph. A
  projected one-Human document loads and resaves, yet its first shop route
  hangs, and the original-produced resave later crashes on that route
  (`SAV-ORIGVALUE-399`); the unsafe field is Unknown.
- Beyond it, a live-state no-copy/no-guess accounting still lacks producers for:
  - eleven world-head dwords;
  - Position `+0x06..+0x07`, failed cell joins and computed tick receivers;
  - unnamed session spans/scalars;
  - campaign `+0x114` and `+0x128`;
  - the final global dword and 400-byte world object;
  - the root-specific external reconstruction projection.
- A source-faithful diagnostic can carry one ROM1-derived vector
  (`SAV-WORLDDIAG-433`), but that is not a derivation from current Againrom
  state and closes none of these producers.

**Confidence.** High for the explicit field-accounting stop, the observed Human
boundary and the diagnostic/control distinction.

**Unknown.** Causal-minimal sufficient values, live-Snapshot producers, first
computed consumer membership, EN/RU original compatibility and resave fidelity.

### SAV-WORLDDIAG-433

- The diagnostic decodes the original-produced EXP-0261 mission-30 transition
  (`d8e6f1d2…`, `SAV-ORIGMISSION-400`) into named document fields, 145 CArchive
  objects and their relations, 1,841 terrain blocks, 184 cell records, six
  bounded session families, a typed YA1 tree and campaign records, then rebuilds
  each programme.
- Its 64,074-byte decoded document is byte-equal to the source, but its outer
  stream is 253 independently emitted literal packets with no runs.
- Its 955-byte YA1 store changes 247 physical bytes by reordering records
  canonically, zeroing unused name padding and recalculating pool offsets, while
  preserving every path, kind and value.
- Two fresh runs produce the same 65,826-byte file, SHA-256
  `b3c7c0e81fd437dbe486e13a25af2b10672a822e99440b9e9ddfa7bf1c9e3dad`. The
  complete reader reaches EOF with 460 replayed archive events, zero unresolved
  identities, zero gaps and zero overlaps.
- This is a diagnostic source-faithful value vector, not a live-Snapshot writer
  or proof of original chooser/load/action/resave.

**Confidence.** High for independent physical encoding, deterministic
generation, semantic equivalence and complete-reader closure.

**Unknown.** ROM1 EN/RU B/C/D and implementation-authorable value producers.

## Campaign-record `+0x114` and `+0x128`

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-598 | Campaign-record `+0x114` and `+0x128` are raw, uncomputed passthrough in both the writer and the reader, and neither is read again inside the reader's own tail. | High | ✔ promoted | [EXP-0305](../experiments/EXP-0305-campaign-scalars/) |
| SAV-599 | `+0x114`'s reach outside the SAV is `UNIT-GATE-012`'s server-singleton chain plus the campaign record's own constructor and reset; no other consumer is located. | High / Medium | ✔ promoted | [EXP-0305](../experiments/EXP-0305-campaign-scalars/) |
| SAV-600 | Three of the campaign record's own methods read or write `+0x128` outside the SAV writer/reader: its constructor, its reset, and a floating-point reader. | High | ✔ promoted | [EXP-0305](../experiments/EXP-0305-campaign-scalars/) |
| SAV-601 | Campaign-record `+0x128` is not constant across the owner's preserved saves, refuting an inert/round-tripped reading for this field specifically. | High | ✔ promoted | [EXP-0305](../experiments/EXP-0305-campaign-scalars/) |
| SAV-602 | `+0x128`'s only located reader beyond the SAV writer/reader (`SAV-600`) never writes it, and no located instruction produces the corpus's observed non-zero values (`SAV-601`). | High / Unknown | ✔ promoted | [EXP-0305](../experiments/EXP-0305-campaign-scalars/) |

### SAV-598

- Writer `R1294`'s seven-scalar suffix writes
  `+0x118,+0x114,+0x110,+0x11c` by the identical
  shape (the address `record+off` formed, 4 pushed, that address pushed, the file passed as `this`, then the primitive called)
  (`+0x114` at `L08409`).
- It then computes `+0x120` from a live-view/first-MapPoint comparison
  (`L08410`..`L08411`, the same comparison `SAV-CAMPPOS-072` reads) before
  `+0x124` and `+0x128` (`L08412`) return to the plain shape.
- Reader `R0434` mirrors the order and shape (`+0x114` at `L08413`,
  `+0x128` at `L08414`). Immediately after, `L08415` moves on to the next
  field with no further reference to either.
- Later in the same function, `L08416` re-reads `+0x120` from the object to
  branch a position-restore arm, and `L08417` re-reads `+0x118` to feed a
  mission lookup. `+0x114` and `+0x128` are never re-read between the
  seven-scalar suffix and either of the function's two exits that return with 4 argument bytes cleaned.
- `SAV-CAMPAIGN-076` fixed both fields' wire position without naming them; this
  claim narrows their in-record treatment.

**Confidence.** High. Both inverse bodies are read through their own epilogues
at instruction level. The contrast with the computed write of `+0x120` and the
reader-tail reuse of `+0x120`/`+0x118` is the discriminator between
"passthrough" and "computed/reused" for the two fields in question.

### SAV-599

- Absolute-form census (`campaign+0x65c`): 7 `disp` hits, 6 owners, 0 `imm`. It
  reproduces `UNIT-GATE-012`'s own count exactly and is entirely that claim's
  already-named population: its default write, its new-campaign-arm write, its
  pre-create-screen UI read/reverse-store, and the two `INC`-and-store load
  paths `SAV-CAMPTAIL-070` names. It adds no population.
- Record-relative form: 49 `disp` hits, 29 owners. A corrected
  callee-reachability crosscheck, which walks each of the 35 `+0x548`
  record-base sites to its own resolved `CALL` target, resolves them to exactly
  four methods, no fifth:
  - the writer;
  - the reader;
  - the record's constructor `R0677`, which zeroes `+0x114` at
    `L08418`, storing a zero built by a self exclusive-or at `L08419`;
  - its reset `R0757`, which zeroes it again at `L08420`, storing a
    zero built by a self exclusive-or at `L08421`.
- Both load paths read `campaign+0x65c`, increment it, and store to
  `[L00285]+0x84`: Load A at `L08422`..`L08423`, Load B at
  `L08424`..`L08425`. `L08426` is not a Load B site: it is mid-instruction
  and decodes inside a different, adjacent function.
- The load driver's own unconditional world-restore call (`L08427`), which
  runs before either load path's mutually exclusive arm at `L08428`, reaches a
  guarded store to the same singleton field (`L08429`). The campaign `INC` is
  therefore always the last writer of that field before any reader runs.
- `UNIT-GATE-012` names three readers of that field: the unit-information panel,
  the world `Serialize` store arm, and the `.alm` spawner `UNIT-GATE-013`
  details. It grades its own reader enumeration Medium, not High. None of the
  three is a sole consumer.
- `field_114` is 1 in all 23 EN and 4 RU preserved owner saves walked for
  EXP-0305, so the `INC` always lands on 2, the neutral arm.

**Confidence.** High for the exhaustive addressing-convention census under both
forms, cross-checked by callee reachability rather than address proximity, and
for the constructor/reset zero-writes, each read at instruction level. Medium
for "the `INC` always lands on 2" in the walked corpus, scoped to this 27-file
snapshot and to `UNIT-GATE-012`'s own Medium-graded reader count.

### SAV-600

- Intersecting the 28 `disp:548` owners against the record-relative owners by
  function name cannot see a method that receives the record pointer directly as
  `this`. Such a method never computes `app+0x548`, so by construction it cannot
  appear in that intersection, and the intersection cannot support a universal
  negative.
- Walking each of the 35 `disp:548` sites to its own resolved `CALL` target
  instead finds three methods:
  - `R0677`, the constructor (one caller, `L03281`), zeroes `+0x128` at
    `L08430`;
  - `R0757`, the reset (three callers, including `SESS-START-034`'s
    new-campaign driver at `L08431`), zeroes it again at `L08432`;
  - `R0802` (one caller, `L03704`) loads `+0x128` once
    (`L08433 FILD`, never a store, in its fully disassembled body) into a
    floating-point product/ratio with the sibling field `+0x124` and an external
    object's own `+0x108`, converts the result to an integer, and passes it to
    a call to `L03688`, which itself owns no record-relative `+0x128` reference.
- The owner-name-overlap axis still disproves 6 (not 5) apparent record-relative
  name collisions among the 28 `disp:548` owners by direct register-provenance
  tracing. That axis resolves coincidental name overlap, a narrower and
  different population; it could never rule the three record methods above in or
  out, since none of them computes `app+0x548`.

**Confidence.** High for the positive existence claim: three named methods, each
read at instruction level with its own confirmed caller population from a
whole-image `callto:` sweep, touch `+0x128` outside the writer/reader, and none
of them writes a non-zero value. Not exhaustive: a `disp:548` site whose own
call is an indirect virtual dispatch (an indirect call through the dword at `[reg+0xc]`, a vtable slot)
is not followed by this one-hop callee walk, so a fourth, unlocated consumer
reached only through such a site is not excluded (`SAV-602`).

### SAV-601

- Across 23 EN and 4 RU saves walked by `tools/savcampaign`, `field_128` ranges
  over `{0,2,3,5,6,9,22,26,27,30,33}`.
- A 13-save cluster (EN `game0002,0003,0004,0006,0011,0012,0013,0015,0021`; RU
  `game0000,0001,0002,9999`) is byte-identical across 27 of
  `campaign-values.tsv`'s 33 reported columns. These are every campaign-state
  field the tool reports, not a hand-picked subset, including selected mission,
  `AutoGetMission`, `LastMission`, the first-MapPoint flag, mission time,
  document/child counts, mercenary rosters, and the base/list progress markers.
- The cluster differs only in per-file bookkeeping (`path`,
  `label_hex`/`label_ascii`, `file_bytes`, `campaign_offset`) and in `field_128`
  itself, which still takes 5 distinct values, `{0,2,3,5,6}`, within that
  identical-state cluster.
- `+0x128` is therefore not a deterministic function of any other reported
  campaign scalar, individually or jointly, within this population.

**Confidence.** High for non-constancy itself: a direct reading of the corpus,
dispositive against a fixed-value model. High, not Medium, for the
joint-constancy check: all 27 non-bookkeeping columns the tool reports were
compared for the cluster and found identical, an exhaustive comparison over the
reported population rather than a selective one. A field the tool does not
report at all remains untested.

### SAV-602

- `R0802` references `+0x128` exactly once, a load
  (`L08433`, a 32-bit integer load from offset 0x128 of the object onto the floating-point stack). No instruction in its fully
  disassembled body stores to that address.
- It multiplies that value by a ratio of the sibling field `+0x124` and an
  external, session-scoped object's own `+0x108` when `+0x124` is non-zero, or
  by a fixed constant times that same external field when `+0x124` is zero. In
  neither arm is `+0x128` a divisor, so `0` is an ordinary multiplicand, not an
  excluded or special input.
- Every value in the corpus range `{0,2,3,5,6,9,22,26,27,30,33}` is
  arithmetically consistent with having passed through this reader. But a reader
  cannot be the writer.
- The two located writers, `SAV-600`'s constructor and reset, both write `0`,
  not the observed range. EXP-0305 locates no third writer anywhere in the
  image, under either addressing convention or the corrected callee-reachability
  crosscheck.
- The leading unexcluded explanation is one of `docs/INSTRUMENT.md`'s named
  blind spots; neither is confirmed:
  - a `disp:548` site whose own call is an indirect virtual dispatch
    (an indirect call through the dword at `[reg+0xc]`, a vtable slot) that the corrected crosscheck's
    one-hop callee walk does not resolve;
  - a bulk `memcpy`/`REP MOVS*` copy from a differently-addressed source.
- The one concrete bulk-copy candidate found in the image was ruled out.
  `R0701`'s two `MOVSD.REP`/`MOVSB.REP` pairs at `L08434`/`L08435`
  and `L08436`/`L03824` copy a text-table line, read through the local
  string accessor `R0668` (`TEXT-STRTAB-023`, `TEXT-UI-033`), table
  `L06351` index `0x37`, into offset 0x134 of the object, and a second, directly-addressed
  string at `L03822` into offset 0x234 of the object. Neither destination is `+0x128`.

**Confidence.** High for the narrower checked facts: `R0802` is read-only
on `+0x128` in a fully disassembled function body, `+0x128` is never a divisor
in its computation, and its downstream call does not itself reference the field.

**Unknown.** The write mechanism. No alternative is discriminated; this claim
states the open gap rather than closing it.

## Mercenary shelf, hire flags, party census and marker pictures

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-606 | Over 55 files, `Mercenaries[]` equals the registry's `[Mission<n>] Mercenaries` value for the record's current main mission, with zero exceptions; it is not a cumulative union of missions played. | High | ✔ promoted | [EXP-0308](../experiments/EXP-0308-town-sav-homes/) |
| SAV-607 | The campaign record's fifteen-element hire-flag array (`+0x88`) is all-zero in every one of 55 accessible files (31 distinct SHA-256 byte streams), zero exceptions. | Medium | ✔ promoted | [EXP-0308](../experiments/EXP-0308-town-sav-homes/) |
| SAV-608 | No accessible save's Player-owned actor graph contains an individually addressable hired mercenary; every actor is drawn from one of nine fixed identities, none resembling the tavern's 1–15 type space. | High | ✔ promoted (partially retracted) | [EXP-0308](../experiments/EXP-0308-town-sav-homes/) |
| SAV-609 | Marker-picture rehydration is not LOAD-synchronous: `R1490`'s one caller is the world-map-enter routine, reached only through virtual dispatch, not either SAV Load routine or the campaign loader. | High / Unknown | ✔ promoted | [EXP-0308](../experiments/EXP-0308-town-sav-homes/) |

### SAV-606

- Corpus: 55 files (23 EN, 4 RU, 28 dated snapshots; 31 distinct SHA-256 byte
  streams, 24 redundant files), spanning every distinct main mission the
  population reaches.
- `tools/savcampaign -state` reports exactly five distinct
  `(main_mission, mercenaries)` pairs over all 55 files (14/8/4/3/2 over the 31
  distinct states): `10→{1}`, `20→{1}`, `30→{14}`, `40→{14,6,10}`,
  `50→{14,6,10,13}`. They match `EXP-0062`'s `offers.csv` registry values digit
  for digit at every point.
- Type 1 is present at main mission 10 and 20 and absent at 30, 40 and 50: a
  type present at an earlier mission can be absent later. This is a positive
  counter-example against a cumulative-union model, not agreement with a
  prediction.
- This corroborates `SAV-CAMPAIGN-083`'s "current mission's shelf" reading with
  a corpus size that claim's own text does not state.
- An image-wide `disp:`/`imm:` census of `campaignScreen+0x5e4`/`+0x5e8`
  (bracket-displacement addressing) reproduces `MERC-SHELF-002`'s own
  single-reader finding (`R1412`, two hits per address, both reads, zero
  writes) rather than adding to it.
- A full disassembly of that reader's loop (`L08437`-`L08438`) settles the
  remaining alternative:
  - `L08437` reads the stored count (`+0x5e8`);
  - `L08439` reads the stored array pointer (`+0x5e4`, the field the census
    shows has no writer);
  - `L08440` indexes it by the loop counter;
  - `L08441` calls a per-element membership test against the record;
  - `L08438` closes the loop on the stored count re-read at `L08442`.
- No instruction anywhere in `L08443`-`L08438` reads a mission-number field
  (`+0x110`/`+0x114`/`+0x118`/`+0x11c`).

**Confidence.** High for the exhaustive corpus check, the type-1 counter-example
falsifying cumulative union, and the reader decode refuting live per-visit
computation (Q1/H3): `R1412`'s own loop reads a stored count and a stored
array, never a mission-number field. Not resolved by this claim: which routine
writes `+0x9c`/`+0xa0` under record-relative (non-screen-relative) addressing.
The writer search is scoped to the tavern's own screen-relative addressing form
and found no writer there, so it rules out a screen-relative writer and the
read-time-computation alternative, not the writer's location.

### SAV-607

- `tools/savcampaign -state`'s `hire_flags` column is
  `[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]` across all 23 EN, 4 RU and 28 dated-snapshot
  files. It was checked separately per population, combined, and over the 31
  distinct byte streams the 55 files reduce to.
- This is what `MERC-DEATH-006`'s unconditional mission-end zero
  (`R1605`, its only caller) predicts.

**Confidence.** Medium. Corpus agreement with an already-published mechanism,
not a discriminated fact: no accessible save was taken mid-mission with an
outstanding hire, so this claim cannot separate "flags are correctly reset" from
"flags are never set outside a state this corpus cannot reach". `SAV-615`
supplies that mid-mission case from a different population; this claim's 55-file
scope is unchanged.

### SAV-608

- `tools/savunit -mode party` walks every actor reachable from a valid human
  Player's group subtree across all 55 files: 152 actors in total, 82 over the
  31 distinct SHA-256 byte streams, with the same nine identities.
- It finds exactly nine distinct `(classKey, name)` pairs: `26 Danath` (45),
  `27 Naira` (4), `28 Fergard` (6), `29 Reniesta` (16), `42 Brian` (7), `54 –`
  (9), `58 –` (42), `200 Witch` (9), `201 Sarindar` (14).
- `classKey 42`/Brian's seven occurrences (four distinct save states)
  corroborate `SAV-CAMPAIGN-087`/`088`'s single-file `game0020.sav` finding
  ("the save's Brian object is in the Player group graph, not this campaign
  record") across the full corpus, not one fixture.
- The two unnamed classKeys are read as not evidence of hired mercenaries (see
  Amended):
  - `54` appears only in three files (two distinct states), always exactly three
    per file, always runIDs 88–90;
  - `58` appears in fourteen files (seven distinct states), always exactly three
    per file, runIDs 112–116, across two different named lineages, Danath and
    Naira;
  - the stated discriminator is a fixed, small, per-file-invariant count and a
    narrow, contiguous runID band, not the scattered creation-order ids and
    variable per-save count a player-authored hire drawn from a 1–15 type space
    with a Data.bin-templated name (`MERC-LEVEL-005`) would take.
- Both classKeys' `owner` field equals the same file's own Player identity in
  every occurrence. That alone does not discriminate a hire from a structural
  entry: `MERC-DEATH-006`'s mission-end tally and `PARTY-MERC-007`'s
  group-attachment mechanism mean a hired mercenary would carry the identical
  owner value, since `tools/savunit -mode party`'s own filter
  (`tools/savunit/main.go:1083`) only ever prints owner-equals-Player rows. The
  fixed count and runID band are the discriminators, not self-ownership.
- Scope: the 55 files' 55 human-participant Players only (`+0x28 == 0`). The
  same population's whole Player-owned actor graph, human and non-human Players
  together, totals 1109 owner-matched actor records across 168 valid Players.
  This census is 152 of 1109, 14%; the other 957 belong to 113 non-human Players
  this claim does not search. Not searching them rests on the same
  `PARTY-MERC-007`/`MERC-DEATH-006` mechanism already cited for not searching
  the AI side, not on a finding about them.

**Confidence.** High, bounded to the 55-file accessible population's 55
human-participant Players. This is an exhaustive, targeted search for the
counter-example Q2/H1 predicts (an outlier classKey or name), not agreement with
a favored model: `PARTY-MERC-007`/`ITEM-CARRY-015`'s structural-exclusion
mechanism predicts exactly this outcome, and a single outlier actor would have
refuted it. None was found in any of the 55 files.

**Amended.** [`retracted.md`](retracted.md) withdraws the reasoning of the "not
evidence of hired mercenaries" clause. classKey 58, partially retracted: a hire
confirmed independently of classKey has exactly this fixed count and narrow
runID band (`SAV-614`..`SAV-617`, EXP-0309). classKey 54, narrowed: the
discriminator fails on mechanism grounds, but the conclusion stands on
`41.alm`'s authored Player-slot placement and structural exclusion of hire
(`SAV-624`, `SAV-627`, `SAV-628`, EXP-0310). The nine-identity census, the Brian
corroboration and the owner-field discussion stand.

### SAV-609

- An image-wide `callto:` census of `R1490` (`TOWN-123`'s marker-picture
  rehydration) finds one hit: `R1409` at `L08042`, an ordinary
  unconditional CALL inside that function's own body. Full disassembly confirms
  it: SEH-prologue MFC method, the `this` pointer kept in a frame register at `L08444`.
- `R1409` is `TOWN-117`/`TOWN-118`'s world-map-enter routine. The same
  census on it finds zero direct CALL references anywhere in the image; its only
  reference is one `.rdata` dword at `L08445`, a virtual-dispatch vtable
  slot.
- The same probe run's census on the campaign loader `R0434` reproduces
  `EXP-0305`'s two published Load call sites exactly (`R0098` at
  `L08112`, `R0099` at `L08116`). This cross-validates the addressing
  before concluding that neither load routine nor the campaign loader itself
  calls `R1490`.
- EXP-0308's independently selected 55-file population (31 distinct SHA-256 byte
  streams; a different selection rule from whatever `EXP-0232` used) also has
  zero non-zero-marker files, replicating `SAV-CAMPAIGN-086`'s 55-file corpus
  finding under a different population construction.

**Confidence.** High for the call-graph fact: an exhaustive image-wide caller
census of `R1490` finds exactly one caller, reached only by virtual
dispatch from world-map entry, not from either load routine or the campaign
loader. This discriminates two readings of `SAV-CAMPAIGN-086`'s "load rehydrates
its picture pointer". Read as LOAD-synchronous, that sub-clause is not supported
by this call graph. Read as "the persisted marker record round-trips through
LOAD and its picture becomes valid once more", the underlying claim stands, and
this claim supplies the missing timing detail rather than contradicting it. The
persisted marker bytes still round-trip through LOAD unchanged, which this claim
does not dispute.

**Unknown.** Per `TOWN-124`: whether ordinary play reaches world-map entry
immediately after every LOAD, which would make the timing distinction
practically invisible even though it is real at the instruction level.

## Witnessed hire

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-614 | Across a witnessed hire (`game0004.sav` to `game0005.sav`), the campaign record's only changed field is the hire-flag array; `Mercenaries[]`, pool counts, unlocks, inn/mission arrays, documents and markers are byte-identical. | High | ✔ promoted | [EXP-0309](../experiments/EXP-0309-hire-witness/) |
| SAV-615 | The three actors first seen in `game0005.sav` keep one identity (ident, runID) through `game9999.sav` into `game0006.sav`, and the hire-flag array stays set through mid-mission play. | High | ✔ promoted | [EXP-0309](../experiments/EXP-0309-hire-witness/) |
| SAV-616 | A hire confirmed independently of classKey reproduces classKey 58's published corpus signature field for field but not classKey 54's, refuting `SAV-608`'s non-hire reading for classKey 58. | High / Unknown | ✔ promoted | [EXP-0309](../experiments/EXP-0309-hire-witness/) |
| SAV-617 | The hire creates one new group under the Player (groups 1→2) holding exactly the three new actors (groupActors 1→4); every record the witness changes has an already-published writer. | High | ✔ promoted | [EXP-0309](../experiments/EXP-0309-hire-witness/) |

### SAV-614

- Population: a fourteen-file owner corpus outside SAV-606/607's 55-file
  population, at `gameversions/saves/2026-08-27/EXP-0261-owner-runs/`, outside
  this repository.
- `tools/savcampaign -state` over the two files reports fourteen top-level
  fields per record. Exactly one changes: `dword_array` index 13 goes `0`→`1`.
  The index is zero-based, `MERC-HIRE-003`'s own `dword[record+0x88+(t-1)*4]`,
  type 14 one-based.
- Byte-identical: `base`; `children`; `documents` (three entries); `markers`
  (empty); `parallel_u16` (both fifteen-element pristine/working mercenary
  pool-count arrays); and every one of the six `u16_arrays` entries:
  `mercenaries` `[14]` (`Mercenaries[]`, matching `SAV-606`'s own `30→{14}`
  pairing for main mission 30), `permanent_unlocks`, `inn_npc` (`InnNPC[]`),
  `inn_mission` (`InnMission[]`), `tc_mission` (`TCMission[]`) and
  `shop_mission` (`ShopMission[]`).
- `Player+0x38` money (`tools/savunit`, `SAV-OBF-029`'s field) moves 683→163, a
  520-point debit. It equals `MERC-PRICE-004`'s cost formula applied to type
  14's own already-published price fields:
  `(PriceA 28 + 3·PriceB 8) × unitPrice(30) 10 = 520`, digit for digit against
  EXP-0062's committed `evidence/costs.csv` row `30,14,NPC14_1,3,28,8,10,520`.
  That row was read as committed repository evidence, not by opening
  Data.bin/npc.reg directly, the same way `SAV-616` reads EXP-0308's committed
  evidence.
- Type 14 and the full three-head pool (`n=3`) are both fixed independently, by
  the flag index and `Mercenaries[]`, not assumed from the price match itself.
- The only no-hire control in this witness, `game0002.sav`→`game0003.sav` (a
  separate, earlier run), is city→mission-entry, not this claim's city→city
  shape; its campaign record and money stay unchanged across that transition.
- No city→city no-hire pair exists in this witness. Every other file pair
  carries a distinct Player identity (`tools/savunit`'s own `ident` field, a
  runtime pointer compared only within one session), so none besides the hire
  pair is a same-session before/after at all.
- What ties the observed flag flip to an actual hire call is `MERC-HIRE-003`'s
  own exhaustive two-writer caller census (hire, dismiss).

**Confidence.** High for the exhaustive field-by-field diff: every JSON key
`tools/savcampaign -state` emits is compared, not a subset, and the changed
index matches `MERC-HIRE-003`'s published address arithmetic exactly. This
discriminates against the alternative that a hire also touches `Mercenaries[]`
or the pool-count arrays, which this diff finds untouched. High for the
money-delta exact-value match: 520 is the formula's digit-for-digit output for
type 14's published price fields at the type and pool size fixed independently
(the flag index, `Mercenaries[]`), not an order-of-magnitude coincidence. This
discriminates against any unrelated cause of the same city-save's money change
landing on that exact figure.

### SAV-615

- `tools/savunit -mode party` shows idents `05910460`/`05910c40`/`05911810`
  (runID 17/18/19), owner equal to the same Player identity, present with
  identical typeWord/worn/classKey in `game0005.sav`, `game9999.sav` and
  `game0006.sav` alike: the same three objects, not three new objects that
  happen to share a classKey.
- `05910460` shows health 120 in `game9999.sav` and 117 in `game0006.sav`, a
  change during actual mission play, not a static placeholder value.
- `tools/savcampaign -state`'s `dword_array` index 13 reads `1` in all three
  files: the flag survives mission entry and stays set into mid-mission play.
- This is the case `SAV-607`'s confidence paragraph named as untested ("no
  accessible save was taken mid-mission with an outstanding hire"). It is
  evidence from a different population; `SAV-607`'s own sentence, scoped to its
  55-file corpus, is unaffected.
- The no-hire control `game0002.sav`→`game0003.sav` shows the same
  city→mission-entry transition producing zero new classKey-58 actors and an
  already-zero hire-flag array staying zero, ruling out mission entry alone as
  an account of either fact.

**Confidence.** High for the persistence and liveness facts: identical
`(ident, runID)` across three files is the discriminator against re-creation,
the health change is the discriminator against a placeholder, and the control
pair rules out mission entry alone as the account of the actor set or the flag.
Silent on mission end: `game0006.sav` is mid-mission, not post-mission, so
`MERC-DEATH-006`/`PARTY-MERC-007`'s published account of what happens when a
mission actually ends is untested by this claim.

### SAV-616

- The hire is identified by the hire-flag transition, the money debit and a new
  group under the Player (`SAV-614`, `SAV-617`).
- The witness's three confirmed-hire actors (`game0005.sav`) are typeWord
  `0x000a`, worn 6, exactly three per file, runIDs 17–19 (one contiguous band),
  health 120 at hire.
- `SAV-608`'s classKey-58 numbers, re-extracted from EXP-0308's committed
  `evidence/party/party-lifecycle.txt` (`awk` on column position, 42 rows over
  14 files): typeWord `0x000a` (42/42), worn 6 (42/42), a fixed
  exactly-three-per-file count, health in `{76,101,111,120}`. The typeWord, worn
  count and per-file count are the same, and 120 is one of that set's four
  values.
- The one difference, the runID band's absolute position (17–19 here, 112–116
  there), is what `MOVE-ID-016` predicts across two different sessions: runID is
  "a lowest-free-bit bitmap allocation — creation-ordered only until the first
  death decays," reset and re-marked per session. It is not expected to agree
  across files from different sessions, so it does not discriminate a hire from
  anything else.
- classKey 54's published signature (9 rows/3 files: typeWord `0x0003`, worn 7,
  health locked at 117) matches none of this, worn 7 not 6 and typeWord `0x0003`
  not `0x000a`, so this witness says nothing about classKey 54 in either
  direction.

**Confidence.** High for the narrow clause this refutes: `SAV-608`'s argument
that classKey 58's signature is "not... the scattered creation-order ids and
variable per-save count a player-authored hire... would take." A confirmed hire,
tied to classKey 58 only by coincidence of transition (never by matching the
signature itself, which would be circular), takes exactly that signature.

**Unknown.** Whether the 55-file corpus's own 14 classKey-58 files were
themselves earlier, unrecognized hires: no before/after pair exists for those
files in EXP-0309, so this claim does not certify them retroactively. classKey
54's status, in either direction.

### SAV-617

- `tools/savunit -mode party`'s header line reports `groups=1 groupActors=1` for
  `game0004.sav`'s Danath and `groups=2 groupActors=4` for `game0005.sav`'s
  Danath, in the transition `SAV-614` measures: net +1 group, +3 actors.
- Danath's own record (ident `059da938`, runID 1) is unchanged field-for-field
  between the two files.
- The three new records are class `Human` (`tools/savunit`'s own class tag).
  This matches `MERC-LEVEL-005`'s `R1142` … `new Human(0x1e8)`
  object-creation call inside the routine `MERC-HIRE-003` cites as the
  `0x38`-command server handler, with its `for (i = 1; i <= n; i++)` spawn loop.
- The new group and its actors match `PARTY-ROSTER-002`'s published container
  structure.
- Their runIDs (17, 18, 19) are contiguous, consistent with `MOVE-ID-016`'s
  lowest-free-bit allocator serving one tight allocation loop with no
  intervening allocation.
- Every record EXP-0309 finds changed has an already-published writer: the
  hire-flag array element and the money field (`SAV-614`), and the three actors
  and the one new group (this claim).

**Confidence.** High for the citation cross-check: each changed record maps onto
an already-published, address-level mechanism (`MERC-HIRE-003`,
`MERC-LEVEL-005`/`MERC-CMD-007`, `PARTY-ROSTER-002`) with no gap found, and no
new disassembly was run to reach this. Bounded to this one witness's changed
records, not a claim that every possible campaign-record or actor-graph field in
the whole format has a known writer.

## Origin of classKeys 54 and 58

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-622 | classKey 54 and classKey 58, the two unnamed identities in `SAV-608`'s census, are Data.bin `Humans[]` rows 54 and 58: tavern types 10 and 14 (`NPC10_1`, `NPC14_1`, both level 1). | High | ✔ promoted | [EXP-0310](../experiments/EXP-0310-classkey-corpus/) |
| SAV-623 | Across all 28 campaign maps, classKey 58/`NPC14_1` is the target of exactly two scripted party-transfer instants and classKey 54/`NPC10_1` of zero; the one live classKey-58 transfer is `PARTY-M20-030`'s mission-20 `T08` group. | High | ✔ promoted | [EXP-0310](../experiments/EXP-0310-classkey-corpus/) |
| SAV-624 | Both classKey 58 and classKey 54 are a campaign map's own directly-authored unit; the distinction between the two classes is the authored owner, not existence. | High | ✔ promoted | [EXP-0310](../experiments/EXP-0310-classkey-corpus/) |
| SAV-625 | The 55-file corpus's classKey-58 population (14 files, 7 distinct states) matches `PARTY-M20-030`'s transfer signature with zero exceptions; a within-corpus negative control fits its trigger gate, not an unconditional source. | High / Medium | ✔ promoted | [EXP-0310](../experiments/EXP-0310-classkey-corpus/) |
| SAV-626 | Hire is structurally excluded as classKey 58's origin at every state the 55-file corpus's classKey-58 population occupies, leaving `PARTY-M20-030`'s scripted transfer as the only located mechanism consistent with it. | High | ✔ promoted | [EXP-0310](../experiments/EXP-0310-classkey-corpus/) |

### SAV-622

- The row-number-is-classKey convention is `PARTY-M20-030`'s own: its text
  already reads `Humans[58]` and a save's "class key… 58" as the identical
  number, for row 58/201.
- `tools/humanequip -mode dest -rows 54,58,201 -full` against the lawful EN
  install resolves row 54 to `NPC10_1` (9 declared equipment cells), row 58 to
  `NPC14_1` (8 cells, already named by `PARTY-M20-030`) and row 201 to
  `M10_Merchant` (already named "Sarindar", 2 cells):
  `evidence/humans-destinations-54-58-201.txt`.
- `MERC-LEVEL-005`'s level-1 naming (`"NPC%02d_%d"` on `mission/10−1`) and
  `MERC-TYPE-001`'s fifteen-type roster fix type 10 and type 14 as the tavern's
  own subscripts for these two rows.

**Confidence.** High. A direct read of one already-vetted tool (EXP-0131)
against the lawful install; the row-is-classKey convention itself is reused from
`PARTY-M20-030`, not re-derived.

### SAV-623

- `tools/campaign -mode join -all` (110 lines, `evidence/join-all.txt`) finds
  `Humans[58]` in `10.alm`
  (`op22 group 16 -> player 1 ["Self"] bound=UNREFERENCED`, three units plus the
  same group's Sarindar) and in `20.alm`
  (`op22 group 16 -> player 1 ["Self"] bound=T08`, the same shape). There is no
  third map.
- `bound=UNREFERENCED` means no trigger in `10.alm` reaches that node. Only the
  `20.alm` instance, `PARTY-M20-030`'s own subject, is reachable in ordinary
  play.
- `Humans[54]` matches zero lines in the same 110-line output. This is an
  exhaustive whole-campaign search, not a check of only the map that shares a
  number with the corpus's classKey-54 `main_mission` value (`50.alm`); the map
  those files are on is `41.alm` (`SAV-624`, `SAV-627`).

**Confidence.** High for both counts: an exhaustive grep of one tool's own
complete, whole-campaign output, not a targeted or single-map search.

### SAV-624

- `tools/campaign -mode place -all` (777 lines, all 28 maps,
  `evidence/place-all.txt`) resolves a def-id or npc.reg-arm placement to its
  Data.bin `Humans[]` row through `className`, the same resolution `join`'s
  published `cullVerdict` performs. The earlier instrument printed only the raw
  `defID`/`class2` key, so a `Humans[54]`/`Humans[58]` search matched zero lines
  regardless of what the census held.
- `Humans[58]`/`NPC14_1` is authored three times on `10.alm` (`owner=2`) and
  three times on `20.alm` (`owner=5`), both in group 16, a non-Player group: the
  one `SAV-623`'s `T08` transfer re-owns on `20.alm`.
- `Humans[54]`/`NPC10_1` is authored three times on `41.alm`, directly under
  `owner=1` (`tools/campaign -mode units -map 41.alm`, records 36–38), the map's
  own Player slot.
- `50.alm`'s own `type-5 roster` line reports `slot 1 id=1 "Self" units=0`: the
  map the corpus's classKey-54 files share a `main_mission` number with authors
  nothing into the Player's slot. Those files are on `41.alm`, not `50.alm`
  (`SAV-627`).
- `20.alm`'s own slot-1 line is also `units=0`. classKey 58's population is
  never authored directly into a Player slot anywhere, consistent with its
  transfer origin.
- Authoring into the Player's own slot-1 row is not a general rule across the
  campaign. Of the 28 maps that declare one, 4 author 1 or more units into it
  directly (`41.alm` 3, `71.alm` 4, `150.alm` 1, `151.alm` 1) and the other 24
  author zero; `41.alm` is one of the four.

**Confidence.** High. An exhaustive whole-campaign census over one tool's own
complete, resolving output, and a direct read of three named maps' own
roster-declaration and unit-table lines.

### SAV-625

- Every one of the 7 distinct classKey-58 states
  (`evidence/classkey-58-54-distinct-states.txt`, joining EXP-0308's committed
  `campaign-values.tsv` and `party-lifecycle.txt`) is main mission 20, exactly
  three actors, typeWord `0x000a`, worn 6, and carries a classKey-201 actor at
  the immediately preceding runID. Adjacency is computed as
  `runID == min(classKey-58 runIDs) − 1`.
- These are the count, typeWord and bundled identity `PARTY-M20-030`'s
  single-save reading already names.
- Of the corpus's 15 raw main-mission-20 files (7 EN, 8 dated, 0 RU; not
  deduplicated, to test against the full population), 14 carry this signature.
  Exactly one, `2026-08-02/game9999.sav` (its own distinct SHA-256 state, not
  redundant with any other corpus file), does not.
- An unconditional source would produce 15/15: a direct Player-slot authored
  placement (already excluded by `SAV-624`) or an always-on transfer.
  `SAV-624`'s finding that classKey 58 is authored on `10.alm`/`20.alm` under a
  non-Player owner does not itself produce 15/15, because those units still
  reach the Player only through the same gated transfer.
- The 14/15 split is therefore inconsistent with an unconditional source and
  consistent with `PARTY-M20-030`'s `T08` trigger gate.

**Confidence.** High for the exhaustive count/signature match (every one of the
7 distinct states checked, not a sample) and for the 14/15 split itself, a
direct read of the corpus. Medium for reading that split as evidence of a
trigger gate specifically, rather than some other conditional cause EXP-0310
does not name: the split is consistent with a gated event and inconsistent with
an unconditional one, which is as far as a static count alone can discriminate.

### SAV-626

- All 7 distinct classKey-58 states are main mission 20 (`SAV-625`). `SAV-606`'s
  law gives `Mercenaries[]` at main mission 20 as `{1}` with zero exceptions in
  this same corpus. Type 14, classKey 58's own tavern type (`SAV-622`), is
  absent from the pool, so it cannot be hired at this state.
- Main-mission progress is monotone (`formats/sav/format.md`: "a lower requested
  main mission is rejected"). No earlier point in the same save's history could
  have supplied a type-14 hire either, and no accessible mission-20 file could
  represent "hired at 30+, still showing mission 20."
- `SAV-623`/`SAV-624` exclude a second scripted transfer or an authored
  placement anywhere in the campaign.
- `SAV-625`'s signature match and negative control leave the published `T08`
  transfer as the only mechanism this repository's own claims document that is
  consistent with the corpus's classKey-58 population.

**Confidence.** High. The pool exclusion chains from `SAV-606`'s High-graded,
exhaustively checked law and an already-published monotonicity fact, not from a
new reading. The elimination is over the three mechanisms this repository's own
claims currently document (hire, scripted transfer, authored placement), not
every conceivable mechanism.

## classKey 54 origin and mission-end loss

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-627 | classKey 54 occurs in three corpus files (two distinct states), exactly the corpus's main-mission-50 files, all selecting side-mission map `41.alm`, with a signature unlike classKey 58's. | High | ✔ promoted | [EXP-0310](../experiments/EXP-0310-classkey-corpus/) |
| SAV-628 | classKey 54's three actors are `41.alm`'s authored Player-slot placement, not a hire; hire is structurally excluded at every state the class occupies, as it is for classKey 58. | High | ✔ promoted | [EXP-0310](../experiments/EXP-0310-classkey-corpus/) |
| SAV-629 | One origin-blind mechanism destroys a classKey-58 or classKey-54 actor at mission end, whether hired, transferred or map-authored, so the 55-file corpus holds no post-mission file of either class. | High | ✔ promoted | [EXP-0310](../experiments/EXP-0310-classkey-corpus/) |

### SAV-627

- Population: three files, 1 EN, 2 dated and 0 RU, in two distinct states. The
  corpus's raw main-mission-50 file count is exactly three
  (`campaign-values.tsv`'s `main_mission` column, undeduplicated). All three
  carry this population and this `selected_mission` value; classKey 54 occurs in
  no other file.
- Signature, from `evidence/classkey-58-54-distinct-states.txt`, in both
  distinct states:
  - main mission 50, selected mission 41 (`+0x118`,
    `formats/sav/format.md:1499`);
  - typeWord `0x0003`, worn 7, exactly three actors, runIDs 88–90 (contiguous);
  - health 117 on every one of the six actor-occurrences across both states,
    with no variation, where classKey 58 has a four-value health range;
  - permanent unlocks `[14,6,13,4,7,3,2,9,12]`, never containing type 10;
  - hire flags all zero, and no classKey-201 actor in either file.
- The other four scenario-player actor counts in both states (Danath's party
  aside) match `41.alm`'s roster slot for slot: "Monsters" 16, "Kadagan" 8,
  "Nocturnal" 5, "Beasts" 10. They are read from `party/owner-coverage.txt`
  (`en/game0020.sav`) and from the `type-5 roster` line of
  `tools/campaign -mode place -map 41.alm` (SAV-624).

**Confidence.** High. A direct, exhaustive read of the corpus's committed tables
over its full main-mission-50 population: 3 of 3 files, not a sample.

### SAV-628

- Tavern type 10, classKey 54's type (SAV-622), appears in the permanent-unlock
  array of 0 of the 55 corpus rows, including both classKey-54 states
  (`[14,6,13,4,7,3,2,9,12]`, SAV-627).
- MERC-SHELF-002's shelf filter (`R1412`) keeps a `Mercenaries[]` element
  only when `R1606` finds it in that array. Type 10 therefore cannot pass
  the filter at either state, although `Mercenaries[]` names it
  (`{14,6,10,13}`).
- SAV-623/SAV-624 exclude a scripted transfer campaign-wide.
- The positive account is SAV-624's finding. `41.alm`, the map every classKey-54
  file's `selected_mission` field names (SAV-627), places exactly three
  `Humans[54]`/`NPC10_1` units under `owner=1`, the map's Player slot. That is
  the count, typeWord and worn signature the corpus files carry, on the map
  whose other four factions' actor counts those files' scenario-player roster
  reproduces exactly (SAV-627).
- No claim the repository publishes names a fourth mechanism for creating a
  fresh `Human` actor beyond the tavern spawn pipeline
  (MERC-CMD-007/MERC-HIRE-003) and the map's static authoring (SAV-623/SAV-624).
  Hire is structurally excluded, not merely unwitnessed.

**Confidence.** High. The exclusion chains from MERC-SHELF-002's exhaustively
checked shelf-filter mechanism and an exhaustive read of the corpus's
permanent-unlock array (0 of 55 rows contain type 10); it is not a new reading
of an Unknown. The positive account is a direct four-field structural match
(unit count, owner, Data.bin row, and the other four factions' actor counts) to
one exhaustively identified map, not corpus agreement alone. No hire-flag
transition was witnessed for type 10, unlike classKey 58's sibling case
(SAV-616). That would matter for a hire-consistency verdict; the verdict here is
that hire did not happen, so the absence is expected, not a gap.

### SAV-629

- Both classKeys' typeWords, `0x000a`=10 for 58 and `0x0003`=3 for 54, fail
  PARTY-ENDCULL-026's server-side survival test
  `0x21 <= word[actor+0x0e] < 0x40` (decimal 33–63).
- Neither carries PARTY-CULL-004's client-side survival flag (`+0x18c` bit 0).
  PARTY-FLAG-003's image-wide sweep (`EnumRefs disp:18c`, 232 hits, 0 orphan)
  names exactly one writer of that bit, the hero-builder `R0592`. Neither
  the tavern spawn pipeline (MERC-LEVEL-005's `new Human(0x1e8)`) nor a map's
  scripted-transfer, placement or directly-authored unit construction calls it.
- PARTY-MERC-007's fifteen-integer mercenary pool-count, whose array storage and
  element width `MERC-POOL-011` names, is the one channel that persists a hired
  type across a mission boundary. It does not apply to a scripted transfer's
  bundled identity: classKey 201/Sarindar is not a mercenary type at all, and
  the transferred classKey-58 actors' mercenary-type byte stays zero
  (PARTY-M20-032).
- No hired, transferred or directly-authored instance of either class therefore
  leaves an individually addressable trace once a mission ends. Origin is
  legible only from a mid-mission file.
- A within-corpus pair witnesses the boundary directly:
  - dated-safe `2026-08-02/game0009.sav`: Player ident `02cd0b20`, main mission
    20, `outcome=1`, `money=600`, `groupActors=5` (Danath, Sarindar and the
    three classKey-58 actors, `party-lifecycle.txt`);
  - `2026-08-02/game0010.sav`, one file later in the same session: the same
    Player ident, `outcome` and `money`, main mission 30, `groupActors=2`
    (Danath and Reniesta only).
  - The transferred squad and Sarindar are present in the first file and gone in
    the second; nothing else about the Player's identity changed.

**Confidence.** High. A citation cross-check of High-graded mechanisms
(PARTY-CULL-004, PARTY-ENDCULL-026, PARTY-FLAG-003, PARTY-MERC-007,
PARTY-M20-032) against the two classKeys' established typeWord values, plus one
direct within-corpus before/after read. No new disassembly was run for this
claim.

## Unit mover, route lists and order list

`SAV-637`–`SAV-645` are retired unused: the evidence supports seven claims, not
sixteen, and the reserved range was wider than what the two lists, the order
list, the class mechanics and the stat-block cross-references established.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-630 | The u16 elements of Unit's two embedded route lists (`+0x15c` static, `+0x178` dynamic) are `MOVE-ROUTE-004`'s node `+0x8` packed cell on both the save and load arm; most corpus lists are empty. | High / Medium | ✔ promoted | [EXP-0311](../experiments/EXP-0311-unit-mover-route/) |
| SAV-631 | Outside `Unit::Serialize`, six functions in the `L02478`-`L00612` AI/movement module touch the two route lists: two extraction routines, three runtime drivers and one search function. | High / Medium | ✔ promoted (amended) | [EXP-0311](../experiments/EXP-0311-unit-mover-route/) |
| SAV-632 | The `u16list` append, used identically by all three sites and not only in `Serialize`, grows the pool by 10 nodes only on an empty free list; teardown frees chain and pool and zeroes all five fields. | High | ✔ promoted | [EXP-0311](../experiments/EXP-0311-unit-mover-route/) |
| SAV-633 | The third embedded list, `*(*(Unit+0x158)+0x90)`, holds exactly 0 or 2 elements in the corpus, in one of two value pairs, consistent with `SAV-GRPPATROL-570`'s patrol-ring copy, not shown to be the only writer. | High / Medium / Unknown | ✔ promoted | [EXP-0311](../experiments/EXP-0311-unit-mover-route/) |
| SAV-634 | On load, all three list surfaces are rebuilt as ordinary runtime state through the class's append path; the `SAV-631` census finds no further post-load consumer of the two Unit-embedded lists. | Medium | ✔ promoted | [EXP-0311](../experiments/EXP-0311-unit-mover-route/) |
| SAV-635 | Four of `SAV-UNITFLD-049`'s five listed Unknowns, Unit `+0x8e`, `+0x90`, `+0xa2` and `+0xa4`, are already named at High confidence by other ledgers that do not cite it. | High | ✔ promoted (amended) | [EXP-0311](../experiments/EXP-0311-unit-mover-route/) |
| SAV-636 | `SAV-UNITFLD-049`'s fifth listed Unknown, Unit `+0x18`, is `Token+0x18`: `Token` is `Unit::Serialize`'s unadjusted first component, byte-mapped at High confidence by `SAV-TOKEN-034`. | High / Unknown | ✔ promoted (amended, partially retracted) | [EXP-0311](../experiments/EXP-0311-unit-mover-route/) |

### SAV-630

- Save arm: `u16list::Serialize` `R1546` walks the live node chain from
  `this+0x4` and pushes each node's `+0x8` before calling the element writer
  (`L08446` loads the node pointer from the frame local at -0x4; `L08447` adds 0x8 to form the element address; `L08448` pushes it;
  `L08449` calls `R1586`).
- Load arm: it reads the same two-byte element into a stack local
  (`L08450` calls `R1586`; `L08451` loads the element from the frame local at -0xc) and passes it to the
  append wrapper `R1587`, which stores it in the newly linked node's
  `+0x8` (`L08452` loads the 16-bit argument; `L08453` stores it at offset 0x8 of the node).
- Both arms archive node `+0x8`: the 12-byte node `MOVE-ROUTE-004` derived from
  the two route-extraction routines and names as the packed cell `(y<<8)|x`. An
  element is not an object or waypoint key.
- Corpus: `tools/savdoc -mode mover-route` reads this field inside `unit()`'s
  existing bounded extents, walking no new byte range, over the full
  3,448-record corpus in 16 save directories (`evidence/input-manifest.json`):
  - the static list is nonempty on 518 records (15.0%), with 1..13 elements;
    every count 1-11 and 13 occurs, 12 does not;
  - the dynamic list is nonempty on 177 (5.1%), with 1..5 elements;
  - all 2,371 decoded elements are plausible in-bounds cells, x ∈ [10, 110], y ∈
    [12, 132] (`evidence/mover-route-summary.tsv` rows `element_count`,
    `element_x_min`, `element_x_max`, `element_y_min`, `element_y_max`).
- A second, independent corpus check corroborates the element identity. All
  1,556 consecutive element pairs across both lists (1,146 static, 410 dynamic)
  are Chebyshev distance exactly 1, an adjacent map cell. The unit's saved cell,
  decoded from the unrelated Position object (`SAV-TOKENPOS-074`), lies within 2
  cells of the static list's first element on 511/518 records and within 1 cell
  of the dynamic list's first element on 173/177 (rows `static_pair_chebyshev1`,
  `dynamic_pair_chebyshev1`, `own_within2_of_static_first`,
  `own_within1_of_dynamic_first`). An object key, waypoint index or table key
  would not walk adjacent cells or anchor on a position decoded from a different
  structure.

**Confidence.** High for element identity. Both arms are disassembled to the
byte, the save side reached from the object it serializes and the load side from
the object it constructs, not inferred from a corpus stride. The alternative
`SAV-EMBED-039` left open, that an element indexes an object or a waypoint key
rather than a raw cell, is excluded: the node read and written is the node
`MOVE-ROUTE-004` independently derived as a route-search artifact, and its
`+0x8` is a 16-bit field, not a pointer. The adjacency and anchoring check
corroborates this but is not the basis for the High, since it cannot exclude
every alternative that coincidentally yields adjacent-looking values. Medium for
the corpus population reading: single-snapshot counts over the preserved saves,
not a controlled trace of a route in progress.

### SAV-631

- A whole-image `-disp 15c`/`-disp 178` sweep returns 87 and 33 raw hits
  (`evidence/disp-sweeps.txt`). All but a handful lie in code at
  `L08454`-`L08455` with no relation to Unit, AI or movement and are read as
  unrelated classes reusing the same small offset.
- `L08456` in `R1607` (a load of the dword at offset 0x15c of the object), the one hit above
  `L08457` and below the module boundary, is not a list accessor. It is a
  getter that returns `this+0x15c` and `this+0x160` through two out-pointers
  (`L08458..L08459`) after calling `R1608` on the same `this`
  (`L08460..L08461`). `R1608` opens with the same
  comparison of `[this+0x15c]` with zero, a null-check, that opens `R1609`, which treats
  `+0x15c` as a `malloc(0x100)` buffer pointer and `+0x160` as its element
  count, not a `Unit`. `R1607` is a further method of that non-`Unit`
  class.
- Six owning functions in `L02478`-`L00612` touch either field:
  - `R0437` (static) and `R0438` (dynamic), the route-extraction
    routines `MOVE-ROUTE-004`/`MOVE-PLANE-005` publish.
  - `R1543`, captured with a preceding step-cost helper `R1542` in
    one listing. On an empty dynamic list it seeds one synthetic node from the
    mover's cached tail cell `mover+0x06` through the class allocator
    `R0052` that `SAV-632` characterizes (`L08462..L08463`). On the
    ordinary path it pops the dynamic list's tail node, returns it to the free
    list and decrements count (`L08464..L08465`, the mirror of `R0052`'s
    insert), and calls the full teardown `R0439` when count reaches zero
    (`L08466`). It then walks the static list from its tail to find a matching
    cell and hands off to `R1610` (`L08467`, not read).
  - `R0043`. On one arm it defers to the reservation consumer
    `R0039`. On another it fully clears the static list, all five fields
    including the pool block, through `R1355`, guarded by a mover-status
    test at `mover+0x78`/`+0x7c`/`+0x9` (`L08468..L08469`). It separately
    seeds the dynamic list as `R1543` does (`L08470..L08471`).
  - `R1544`, a near-twin of the seed-on-empty snippet
    (`L08472..L08473`), plus a bounded five-node walk of the static list's
    tail chain and a branch into `R1356` when the dynamic list is empty.
  - `R0053`, which computes the static list's address at `L08041`
    (forming the address 0x15c bytes into the object) and is otherwise unread. Over 4KB of body precedes
    that instruction, so it is almost certainly the search driver, not a list
    accessor; it is named because the sweep found it, not because its body was
    traced.

**Confidence.** High for every cited instruction (disassembled in full,
`evidence/raw-listings.txt`, from the two whole-image sweeps in
`evidence/disp-sweeps.txt`) and for the six-function count being exhaustive over
the `L02478`-`L00612` module: every owning function the sweep returned there
is identified, and five of the six were read; of `R0053` only the
instruction at `L08041` was read. Medium, scope-limited, for completeness
beyond that module and beyond `R0053`'s body. An address computed in two
steps, carried in a field, or reached only from inside that function's ~4KB body
is invisible to this instrument, the blind spot `MOVE-PLANE-005` recorded for
its writer enumeration.

**Amended.** The count of owning functions is corrected
([`retracted.md`](retracted.md)). The former text said "Exactly six owning
functions", listed five, and then called `R0053` "a seventh site".
`evidence/disp-sweeps.txt` returns six owning functions for the two
displacements in the module outside `Unit::Serialize` (`R0210`): `R0053`,
`R0437`, `R1542` (the listing that holds `R1543`), `R0438`,
`R0043` and `R1544`. `R0053` is the sixth, the search function the
headline counts, not a seventh. The Confidence paragraph also said every owning
function the sweep returned was read; `R0053` was read only at
`L08041`. The six-function headline and the grades are unchanged.

### SAV-632

- `R0052(this, next, value-via-caller)`:
  - with a nonzero free-list head at `this+0x10` it unlinks and reuses one node
    (`L08474..L08475`);
  - with a zero head it grows the pool by exactly `this+0x18` nodes of 12 bytes
    through the block allocator `R0540`, `this+0x18` being constructed as
    the literal `0xa`=10 (`SAV-EMBED-039`/`SAV-WLIST-040`). It then walks the
    new block backward, linking each node onto the free list
    (`L08476..L08477`), and takes the last one;
  - either way it links the new node's `+0x0`/`+0x4` pointers from its two
    arguments, increments `this+0xc` (count) by exactly one, unconditionally
    (`L08478` loads the count at offset 0xc; `L06994` adds 1 to it;
    `L06995` stores it back at offset 0xc), and calls `R1611(1, node+0x8)` before
    returning the node.
- The append wrapper `R1587` (`Unit::Serialize`'s load arm) and the three
  runtime drivers `SAV-631` names all call this routine directly, so count and
  pool growth behave identically whichever caller inserts.
- `R0439(this)` is the converse. It walks any chain still reachable from
  `this+0x4`, calling `R1612(1, node+0x8)` on each, then unconditionally
  zeroes `this+0xc`/`+0x10`/`+0x8`/`+0x4`, releases the pool block through
  `R1355(*(this+0x14))`, and zeroes `this+0x14`. It passes the value
  stored at `this+0x14`, not its address (`L08479` loads the field at offset 0x14 of the object), where
  the allocator's addition of 0x14 to the object pointer at `L08480` computes and pushes the address
  itself. This is a full list-and-pool release, not a plain counter reset.
- `R1543` (`SAV-631`) calls it exactly when its tail-pop brings count to
  zero. `R0043`'s clear arm inlines the same four zero-stores and the
  `R1355` call instead of calling `R0439`.
- This closes the runtime mechanics `SAV-WLIST-040` left open.

**Confidence.** High. Every cited instruction is disassembled in full in
`evidence/raw-listings.txt`. The count increment and decrement and the pool
growth and release are read from the routines themselves, not inferred from a
corpus stride, extending `SAV-WLIST-040`'s Serialize-only reading of this class
to its complete runtime append and remove path.

### SAV-633

- `tools/savdoc -mode mover-route` decodes the field inside `unit()`'s existing
  `wlist(*(*(+0x158)+0x90))` bound for all 3,448 corpus records: 3,388 (98.3%)
  are empty and 60 (1.7%) carry exactly 2 elements. No record carries 1, 3 or
  any other count.
- The 60 take only two distinct ordered pairs, 30 records each: `64,21;53,19`
  and `67,12;54,13`.
- `SAV-GRPPATROL-570` publishes a writer for this field (`R0164..L08183`,
  `actor+0x158`>`order+0x90`). It prepends the AI's patrol path reversed and
  prepends the actor's current cell last. From a 1-element AI-owned source path
  it yields exactly 2 elements: the actor's cell first and the current waypoint
  (`order+0x02`) second, because `u16list::Serialize`'s save arm walks head to
  tail (`SAV-630`) and the head is the value the setter prepended most recently.
- A `-disp 90` sweep restricted to `L08481`-`L00612` returns 36 hit
  instructions in 12 distinct owning functions (`evidence/disp-sweeps.txt`); a
  whole-image sweep is unusably noisy on so common an offset. Two owners are
  attributable to this field by name: `R0164` and `Order::Serialize`
  `R0199` (`SAV-EMBED-039`). The other ten, `R0172`, `R0161`,
  `R0010`, `R0016`, `R0250`, `R0146`, `R0076`, `R0171`,
  `R0154` and `R0178`, are not distinguished by this sweep from unrelated
  classes' `+0x90` fields.
- The population searched is this 36-hit, 12-owner set inside the named range,
  not the whole image.

**Confidence.** High for the corpus shape: an exhaustive decode of the full
preserved population, internally cross-checked (30+30=60=`order_count_2`), and
for the `SAV-GRPPATROL-570` citation matching the observed 2-element shape.
Medium for completeness of the writer set: 2 of 12 candidate owning functions in
the swept range are positively identified. The claim does not assert that
`SAV-GRPPATROL-570` is the field's only writer, only that it is consistent with
the observed shape.

**Unknown.** What the other 10 owning functions access at `+0x90`, and whether
any of them writes this field.

### SAV-634

- `Unit::Serialize`'s load arm and `Order::Serialize`'s load arm
  (`R0199`, which destroys and reallocates the prior list before
  refilling it) both reach the class append path `R1587`/`R0052`
  (`SAV-630`/`SAV-632`) that every runtime writer uses. Nothing distinguishes a
  loaded node from one built by ordinary play.
- The owning-function set the whole-image `-disp 15c`/`-disp 178` sweep returns
  inside the `L02478`-`L00612` module (`SAV-631`) holds nothing beyond the
  two `Serialize` routines and the runtime drivers. In particular it holds no
  separate "first tick after load" consumer distinct from the step, seed and
  abort drivers a live unit also calls.
- The negative is bounded to that sweep and that module. An alias stored across
  the load boundary, or a consumer reached only through a virtual call this
  static sweep cannot see, is not excluded.

**Confidence.** Medium. The absence is read from the two whole-image sweeps
`SAV-631` scopes, not from a runtime trace of an actual load-then-tick boundary;
no dynamic instrumentation of the load path was run.

### SAV-635

- `ITEM-LOAD-005`/`HERO-SIGHT-007` name `actor+0x8e` as the actor's carried
  weight and `actor+0x90` as its derived load: own weight plus half the
  inventory container's running weight sum, `[actor+0x7c]+0x20`.
- `AI-SIGHT-092`/`HERO-SIGHT-007` name `actor+0xa4` as the `u16` sight radius in
  1/256-cell units, whose high byte is `actor+0xa5`.
- `HERO-REGEN-021` names `actor[0xa2]` as the health-regeneration accumulator's
  `% 100` remainder byte. That is a third naming: `formats/sav/format.md` had
  already closed `+0xa2/+0xa3` as the separately persisted regeneration
  remainder bytes (`SAV-REGENSTORE-529`, `SAV-REGENWIRE-532`), independently of
  both `HERO-REGEN-021` and `SAV-UNITFLD-049`.
- None of the four claims `ITEM-LOAD-005`, `HERO-SIGHT-007`, `AI-SIGHT-092` and
  `HERO-REGEN-021` cites `SAV-UNITFLD-049` back.
- For `+0x8e`, `+0x90` and `+0xa2`, `SAV-UNITFLD-049`'s dense enumeration names
  only the offset and grades each Unknown with no description. For `+0xa4` it
  quotes the field's value distribution via `UNIT-CTOR-004` ("sight `+0xa4 = 0`
  occurs in no record, the field taking exactly `1024 x {1, 1.25, 1.5, 2}` on
  `Unit`") without resolving the meaning or citing the resolution in
  `AI-SIGHT-092`/`HERO-SIGHT-007`.
- A reader relying on the framing "the unit fields `+0x8e`, `+0x90`, `+0xa4` and
  `+0x18` are promoted Unknown" would not find what any of them means without
  this cross-reference.
- The fifth listed Unknown, `+0x18`, is a different surface: `SAV-636`.

**Confidence.** High. Each cited field was independently disassembled to its
byte and graded High by its own experiment. This claim only cross-references,
and the four resolutions depend neither on each other nor on any new reading.

**Amended.** The count of non-citing claims is corrected
([`retracted.md`](retracted.md)). The card said "None of the three rows cited
above cites `SAV-UNITFLD-049` back" after naming four claims. The experiment
record names the same four, `ITEM-LOAD-005`, `HERO-SIGHT-007`, `AI-SIGHT-092`
and `HERO-REGEN-021`, as claims that do not cite it back. The four field
resolutions and the High grade are unchanged.

### SAV-636

- `SAV-TOKEN-034` states that `Token::Serialize` `R0950` "was reached
  from … Unit's first instruction after the frame", before any pointer
  adjustment. `Token`'s `this` and `Unit`'s `this` are therefore the same
  address, and `Token`-relative `+0x18`/`+0x1c` name the same bytes as
  `Unit`-relative `+0x18`/`+0x1c`.
- `SAV-TOKEN-034` and the Token table in `formats/sav/format.md` fix position
  and width, `u16` at stream bytes 23..24 and `u32` at 25..28, and cite no
  meaning for either row.
- `go run ./tools/claim -k "Token.{0,20}0x18"` and `-k "this\+0x1c"` return only
  `SAV-TOKEN-034`'s byte-map and one false-positive substring match
  (`SAV-HUMAN-043`'s unrelated `this+0x1cc`); no row they reach asserts what
  either field means or names a second reader.
- No whole-image census was run for either offset. `+0x18`/`+0x1c` are too
  common a small-struct displacement for an unrestricted sweep to be useful, the
  caution `SAV-633` applies to `+0x90`, and no code cluster like the AI/movement
  module that scoped `SAV-631`'s and `SAV-633`'s sweeps is known to own the
  `Token` lifecycle.

**Confidence.** High for the `Token+0x18` = `Unit+0x18` identity: `Token` is
`Unit::Serialize`'s literal first-called component, the disassembly shows no
adjustment, and the identity is not an inference from offset coincidence.
Unknown, explicitly bounded, for what `Token+0x18`/`+0x1c` mean and who else
reads or writes them. The search behind that is the two `-k` queries and the
existing text of `SAV-TOKEN-034`/`SAV-TOKENLOAD-095`, not a fresh disassembly
census of either offset.

**Unknown.** What `Token+0x18` and `Token+0x1c` mean, and which code other than
`Token::Serialize` reads or writes them.

**Amended.** The clause "no semantic reader, writer or consumer for it, or for
its neighbour `Token+0x1c`, is named by any claim this repository publishes,
this one included" is retracted ([`retracted.md`](retracted.md),
[EXP-0312](../experiments/EXP-0312-tail-and-registers/)). `ITEM-VALUE-115`
already named a `Token+0x1c` writer under `item.md`'s "Item+1c" wording, which
the `sav`-ledger-scoped `-k` search did not reach (`SAV-654`). The byte identity
and the byte-map stand.

## The stream tail, session registers and Token+0x18/+0x1c

`SAV-659`–`SAV-661` are retired unused: the evidence supports thirteen claims,
not sixteen, and the reserved range was wider than what the stream tail, the
session registers and `Token+0x18`/`+0x1c` established.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-646 | `world+0x118` is a heap-allocated, exclusively owned 400-byte raw block with its own construction, use and teardown: not an embedded array, not shared, and not the session object. | High / Medium | ✔ promoted | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |
| SAV-647 | Two of the 400-byte block's 100 dwords are debug-trace verbosity toggles: `+0x0` gates turn tracing and `+0x4` gates script tracing. | High | ✔ promoted (amended, partially retracted) | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |
| SAV-648 | `[L07886]` is a plain global int with exactly three whole-image touch points: a zero-initializer, a save-arm write after the `0xbadface1` marker, and a marker-gated load-arm read. | High | ✔ promoted | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |
| SAV-649 | Over the 45 parsed of the corpus's 46 distinct save streams, the 400-byte tail is zero in every save, but the stored `[L07886]` is non-zero in one, refuting always-zero and monotonic-counter readings. | Medium | ✔ promoted | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |
| SAV-650 | Session register slot 93 (`session+0xbea8`) has three literal-displacement write sites, and slots 90, 91 and 92 are the only other slots of the 100 written by literal displacement (`SAV-658`). | High / Medium / Unknown | ✔ promoted | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |
| SAV-651 | The check and pattern evaluators also write and read the 100-slot array at any slot, indexed by mission-authored data, unlike slot 93's fixed, script-independent write. | Medium | ✔ promoted (partially retracted) | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |
| SAV-652 | On LOAD, wholesale session restoration precedes the authored check binder `R0067`, whose opcode `0x10002` arm writes map-node+48 into a compiled register slot. | High / Unknown | ✔ promoted (partially retracted) | [EXP-0312](../experiments/EXP-0312-tail-and-registers/), [EXP-0316](../experiments/EXP-0316-session-residuals/) |
| SAV-653 | `Token::Serialize` (`R0950`) confirms the byte-map's `+0x18`/`+0x1c` widths, stream positions and helper pairs in both arms, and exposes no meaning for either field. | High | ✔ promoted | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |
| SAV-654 | `ITEM-VALUE-115` (`claims/item.md`) names a writer for `Token+0x1c` as "Item+1c", the byte `SAV-636` said no claim names a writer for; the two rows do not cross-reference. | High | ✔ promoted | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |
| SAV-655 | Over 3,170 position-verified Token-derived corpus records, `Token+0x18` takes only the values 0 and 2, split cleanly by class with five unexplained exceptions. | Medium | ✔ promoted | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |
| SAV-656 | `Token+0x1c` takes 90 distinct values over the same 3,170 records, with the sentinel `0xffffffff` on 34 of 70 Item records and no other class; consistent with, not proof of, a value semantic for equipment. | Medium | ✔ promoted | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |
| SAV-657 | `formats/sav/format.md`'s `world+0x118`/`this+0x14` phrases name `[L00285]`, `UNIT-GATE-012`'s server singleton for `+0x84`; neither format page states this or traces it for `world+0x118`. | High / Medium | ✔ promoted | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |
| SAV-658 | Slots 90, 91 and 92 of the session register array are also written by literal, non-indexed displacement; the `SAV-650` sweep finds these three and no others among the remaining 99 slots. | High / Medium / Unknown | ✔ promoted | [EXP-0312](../experiments/EXP-0312-tail-and-registers/) |

### SAV-646

- Writer: `R0967`, the world constructor (call-verified start; one direct
  caller, `L07364` in `R1304`). `UNIT-GATE-012` traces the same
  constructor/allocator pair for a different field, `+0x84`; that corroborates
  the function's identification but is not its source.
- It allocates exactly `0x190` bytes through `operator new`
  (`L08482` pushes 0x190; `L08483` calls `R0747`). When the allocation succeeds
  it runs the block's trivial constructor `R1613` (a repeated string store of `0x64`
  dwords from a zero value, no field-specific initialization), then stores the
  result at `this+0x118` (`L08036`, a store of the result at offset 0x118 of the object), the instruction
  sequence behind `SAV-TRAIL-026`'s `world+0x118` pointer.
- The generic (de)serializer `R1491`,
  `SerializeRawBlock(this, 0x190, ar)`, dispatches only on `ar+0x14`'s
  storing/loading direction to the raw block primitives `R0272`/`R0271`
  and exposes no field boundary.
- Teardown: `R1540` frees the pointer through a matching delete-style
  call (`L08484 CALL L08485`) and nulls `*(this+0x118)` (`L08039`). The
  same routine frees and nulls four sibling owned-singleton globals,
  `[L00004]` (session), `[L04624]`, `[L00380]` and `[L00240]`, with
  parallel construct/free/null lifecycles. The block and the session object are
  separate allocations with separate lifetimes, neither nested nor aliased.

**Confidence.** High for the allocation, the writer/reader chain and the
teardown free and null: every cited instruction is disassembled to the byte
(`evidence/q1-world118-and-L07886.txt`, `evidence/q2-session-slot93.txt`), and
the size (`0x190`=400) and the one write site to `this+0x118` are read from the
code, not inferred. Medium for "not shared with any other pointer": the one
write site was located and no second store of the returned pointer appears among
the routines disassembled, but no whole-image sweep for the allocated address
was run, so an alias reached only through code not disassembled would not be
found.

### SAV-647

- The two dwords are not free padding, save content or journal-like state.
- Two paired toggle sites, `L08486` for `+0x0` and `L08487` for `+0x4`, each
  read the current dword, branch on zero or nonzero, flip it, and log one of two
  paired confirmation strings read from `.data`, not `.rdata`:
  - `+0x0`: a 24-byte "on" string at `L08488` (this site's other branch,
    `L08489` pushing L08488; SHA-256 `ddaaf67eed84159d`…, cited by hash and
    length rather than verbatim) and `"Turn tracing turned off.\n"` at
    `L08490`;
  - `+0x4`: `"Script tracing turned on.\n"` at `L08491` and
    `"Script tracing turned off.\n"` at `L08492`.
- Four consumer sites gate a second pair of `Script:`-prefixed trace and error
  lines on the current value of the same two dwords before formatting them:
  `+0x0` at `L08493` and `L08494`, `+0x4` at `L08495` and `L08496`. The
  last two guard the literal strings
  `"Script: Trigger %d ( %d ifs, %d instants ).\n"` (`L08497`) and
  `"Script: Bad inst..."` (`L08498`).
- This answers the "what is this block" question `SAV-TRAIL-026`/`SAV-DOC-053`
  left at "loaded state, not free padding": at least these two fields are
  ordinary engine debug-option flags with matching diagnostic output, which
  excludes a per-save journal or gameplay-state reading for these two dwords.

**Confidence.** High. The toggle and consumer instructions, the string addresses
and their raw bytes are in `evidence/q1-world118-and-L07886.txt`. The paired
on/off strings at the addresses each toggle site pushes leave no other reading
for these two fields. The other 98 dwords are not characterized.

**Amended.** The clause "`-callers` finds 0 direct E8 callers for either toggle
routine's own containing function; they are reached indirectly and no caller is
asserted here" is retracted ([`retracted.md`](retracted.md),
[EXP-0315](../experiments/EXP-0315-world-head-trailer/)). The candidate's
`references.tsv` shows a direct `E8` call at `L02000` into `R0440`, the
function containing both toggle sites, and `EXP-0312`'s regen script never ran
`-callers` for it. The toggle and consumer instructions and the paired strings
stand.

### SAV-648

- It is a global, not a struct member. A whole-image dword-xref scan
  (`-xrefs L07886:L13159`) finds exactly 3 hits in any section
  (`evidence/q1-L07886-xrefs.txt`):
  - `L08499`/`L08500` in the document programme's store arm
    (`R0414`): loading `[L07886]`, pushing it and calling `R0685`, which
    writes the global's current value to the stream, unconditionally, directly
    after the `0xbadface1` marker;
  - `L08501` in the load arm (`R1541`), reached only when the stream's
    stored dword equals `0xbadface1`
    (`L08053` compares the frame local at -0x10 with 0xbadface1; `L08502` branches to `L07783` when not equal), which reads
    the stored value back into the global (the call to `R0686`);
  - `L08503` storing zero to `[L07886]` in `R0131`, the zero-initializer's
    first action. `R0131` has exactly one direct E8 caller, `L06672` in
    `R1536` (`evidence/q1-world118-and-L07886.txt`).
- No other code in the image reads or writes this address by direct or absolute
  addressing.

**Confidence.** High for the three-touch-point enumeration: an image-wide xref
scan for the literal dword, not a sample. A write reached only through
register-indirect addressing, never resolving to a literal `L07886` operand,
would not be found by this instrument; none is asserted.

### SAV-649

- Population: the corpus manifest (`evidence/input-manifest.json`) lists 99
  `.sav` files in 16 preserved-save directories from both EN/RU installs
  (byte-identical executables, `formats/sav/format.md:73`), 46 distinct by
  SHA-256.
- `tools/savdoc`'s load gate requires the header magic `Asg&`. One file,
  `EXP-0261-owner-runs/game9000.sav`, carries `Bsg&` and is not parsed by the
  record-level tools (`SAV-DOC` catalogue). The skip is a diagnosed, named
  reason: `evidence/census-summary.tsv`'s `manifest_unparsed_files`.
- Provenance (`evidence/q1-save-provenance.tsv`): of the 45 parsed streams, 39
  are original-written (produced by a lawful install without passing through
  this project's writer) and 6 are written by this project's SAV writer. The
  split comes from each directory's preserved `MANIFEST.md` or name for 15 of
  the 16 directories, and per file by name in the one further directory that
  mixes both. One directory is overridden: its name reads as an original
  acceptance run, but its one file is byte-identical to six files in unambiguous
  project-writer directories rather than to any observed original-engine resave
  shape. The unparsed 46th stream is also project-written by its directory's
  mixed-split filename convention, giving 39/46 original and 7/46 project
  overall.
- `tools/savdoc -mode tail-census` (`evidence/tail-census.tsv`) raw-scans every
  parseable save body for the `0xbadface1` marker, finds it exactly once in each
  of the 45 parsed streams, and censuses the 400 bytes after the stored
  `[L07886]` dword. All 45 have 0 non-zero bytes there
  (`evidence/census-summary.tsv` rows `tail_census_nonzero_tail_bytes_saves`,
  `tail_census_saves_total`). This extends the existing 4-save finding to that
  population with the same result.
- The stored dword is 0 in 44 streams and `520` in exactly one,
  `EXP-0261-owner-runs/game0005.sav`, an original-written stream. That
  directory's sequential saves read `0, 0, 0, 0, 0, 520, 0` across
  `game0000`–`game0006` (`evidence/tail-census.tsv`). The value returns to 0 in
  the next save after the 520, ruling out a monotonic or cumulative lifetime
  counter.
- The write that put 520 into the global before that save is not located;
  `SAV-648`'s three-touch-point census finds no fourth site.
- A figure about ROM1's behaviour should rest on the 39 original-written
  streams. The 6 project-written ones among the 45 corroborate that this
  project's writer reproduces the same tail and `+0x118`-region shape; they are
  not an independent ROM1 observation.

**Confidence.** Medium. A corpus census, not a controlled trace of the one save
that produced the non-zero value. The tail-zero population is the 45 parsed
preserved streams in the manifest (39 original-written, 6 project-written), not
every possible save, and not the one unparsed stream this instrument cannot
open.

**Unknown.** How a non-zero value first enters the global.

### SAV-650

- Slot 93 is the 94th of the 100 int32 slots at `session+0xbd34`.
- Sweep: `cursorsurf -slotsweep bd34:100` (`evidence/q2-slot-sweep-100.txt`) is
  decode-anchored. Unlike a byte-pattern or `-disp` search, it resolves every
  hit against the containing function's linear decode and accepts a match only
  when the decoder's rendered operand shows the exact hex value. Among all 100
  slots, only 90, 91, 92 and 93 have a literal (non-indexed, non-scaled) write
  anywhere in `.text`.
- Slot 0 also shows the array's base-address idiom, a `LEA`, two `ADD` and 61
  scaled-index accesses that establish or index the whole array rather than
  write one slot, plus one hit near `L08504` the sweep could not anchor to any
  decodable instruction start; that is the unresolved case `SAV-651` described
  as "near `L08505`".
- No other slot shows anything beyond zero hits, generic scaled-index access, a
  coincidental raw-byte match the decoder check rejects, or a bare
  numeric-immediate match unrelated to array addressing.
- Construction: the session sub-constructor `R1289` zero-fills the whole
  100-dword array with `REP STOSD`, then stores four slots by literal `MOV`:
  `90=0` at `L08506`, `91=0` at `L08507`, `92=1` at `L08508`, `93=0` at
  `L08509`. It is reached from `R1541` (`L08510 CALL R0077`),
  which calls `R1289` twice (`L08511`, `L08512`).
- Runtime writers of slot 93:
  - `R1614`, `L08513`, an increment of the field at offset 0xbea8. It has no direct near-call caller
    anywhere in `.text` and zero whole-image dword references of any kind
    (`cursorsurf -xrefs R1614`).
  - An inlined twin at `L08514`/`L08515`
    (a load of the 32-bit field at offset 0xbea8, an increment of it and a store back to the same field) inside the function
    starting at `R0076`. Its prologue is the compiler's SEH idiom
    (a load of the exception-list head from `FS:[0]`, a push of -1, the handler, and a store of the new frame pointer to `FS:[0]`), which `cursorsurf`'s
    automatic function map does not recognize and mis-anchors 6 bytes later, at
    `L08516`; the start was confirmed by hand from the SEH prologue shape.
- Both runtime sites begin an unconditional, unbranched three-step sequence: the
  increment and store, then `CALL R0433` (the check evaluator `TRIG-COND-003`
  names), then `CALL R0428` (the pattern evaluator `TRIG-CMP-006` names). The
  calls are at `L08517`/`L08518` in the standalone routine, whose whole
  eight-instruction body is this triple, and at `L08519`/`L01869` in the
  inlined twin, which continues with other work (`TRIG-EVAL-001`'s
  player→group→unit walk).
- `R0076` has exactly two direct `E8` callers
  (`cursorsurf -callers R0076`):
  - `L00163`, reached unconditionally once its containing routine (start
    `R0075`) passes a branch whose two paths converge before it;
  - `L00279` in `R0147`, gated on `server+0x4 mod 16 == 6`. That is the
    signed-modulo-16 idiom, at a different comparand, that gates
    `SESS-TICK-004`'s full-tick-counter increment at `mod 16 == 15` in the same
    function.

**Confidence.** High for the population: all three slot-93 write sites and the
90/91/92/93-only result come from a decode-anchored sweep of all 100 candidate
displacements, with every accepted hit's operand confirmed to show the target's
hex value, not merely contained in some instruction's byte span. High for the
runtime chaining: each site's three steps are consecutive instructions with no
intervening branch, so the slot-93 increment is confirmed, not inferred, to
precede both evaluators unconditionally every time either site runs. Medium for
the significance of `R0076`'s callers. `R0075`, its unconditional
caller, was disassembled in full: it computes elapsed time through repeated
calls through one function pointer and issues several formatted-string calls
before and after its `R0076` call. Its own callers and outer loop period
were not retraced; `TRIG-EVAL-001` states that period is the full tick the other
caller, `R0147`, gates on.

**Unknown.** Whether or how `R1614`, the copy with no located caller,
ever runs.

### SAV-651

- `R0433`'s prologue keeps session's `this` in a spare
  general-purpose register (the routine addresses its locals from the stack pointer, not
  from a frame register) before a 22-way jump table. Its slot accesses are
  `this + 4*index + 0xbd34` with the index supplied by the check's parameters
  (`evidence/q2-session-slot93.txt`), not a literal slot.
- The same scaled-index shape recurs in the dispatcher that hosts `SAV-647`'s
  log-gate sites (`R0262`, at `L08520`, beside the `session+0xb3ac`
  win-counter increment) and in the separate per-kind helper `L07742`, to
  which no direct caller, jump or absolute pointer is located in the image.
- A fourth context uses the shape at authored-check bind time, including the
  LOAD edge `SAV-711` establishes, rather than in ordinary evaluation:
  `L07782`, a store of the bound value at offset 0xbd34 plus 4 bytes times the index, in the binder `R0067`, gated on
  the check's opcode equalling `0x10002` (`L08521`…`L08522`). This is the
  mission-variable preset `TRIG-COND-003` cites at the same address.
- A whole-image `[base+0xbd34]` sweep, `-disp bd34`, returns 62 hits
  (`evidence/q2-disp-bd34-sweep.txt`), most of them these generic accesses.
  Because `[base+index*4+0xbd34]` still encodes the `0xbd34` literal, that sweep
  alone does not distinguish a literal single-slot access from a scaled-index
  access into any slot.
- `TRIG-STORE-002` and `MISSION-SLOT-008` (`claims/trigger.md`,
  `claims/mission.md`) establish that a check owns exactly one slot at its
  compiled-order subscript and that the shipped corpus's largest map compiles 64
  checks. On that reading no shipped map's check-owned index reaches slot 93, so
  no mission trigger condition in the analyzed corpus reads or writes it through
  the check mechanism.
- Whether a pattern's index assignment is bounded the same way is not
  established by those claims or by this experiment.

**Confidence.** Medium. The classification covers every containing routine
checked: `R0433`, the separately bounded `L07742`, `R0262` and
`R0067`. Whether a hardcoded literal write exists at one of the other 99
slots' distinct displacements is closed for the whole 100-slot population by a
decode-anchored sweep, not by this claim's `0xbd34`-displacement search: slots
90, 91 and 92 are such writes (`SAV-658`), and no other slot shows one
(`SAV-650`). The remaining Medium is for the reasoning that 64 checks stay below
slot 93.

**Amended.** The attribution that a call at `L07741` to `R1485` reaches the
nearby indexed register writers is retracted ([`retracted.md`](retracted.md),
[EXP-0316](../experiments/EXP-0316-session-residuals/)). `R1485` only repairs
the session self-pointer, storing `this` to `session+0xa9c0`, and its caller
`L07741` does not prove reach into the adjacent register writer (`SAV-710`,
`SAV-711`); `L07742` is a distinct body. Authored constants are written
instead by the earlier `R0067` LOAD binder. The evaluator and binder
classification stands.

### SAV-652

- `SAV-711` establishes the binder write; the binder's broader role is
  `TRIG-SAVE-008`.
- The exact thunk `R1485` only repairs `session+0xa9c0` (`SAV-710`).
  `L07742` is a distinct adjacent writer whose LOAD reach is not proved.
- Their address proximity establishes no general per-kind live-object rebuild.

**Confidence.** High for the corrected exact thunk and the selected binder
instructions.

**Unknown.** The complete LOAD path, and whether it reaches any per-kind helper.

**Amended.** The earlier claim that LOAD "re-derives specific slots' values from
other live objects' fields, not only reconstructing bookkeeping", through an
evidence chain from `L07741` to `R1485`, is retracted whole
([`retracted.md`](retracted.md),
[EXP-0316](../experiments/EXP-0316-session-residuals/)). That call chain does
not exist, `L07742` has no located caller, and the binder copies a compiled
map constant, not a live object's field. The corrected LOAD boundary above
replaces it.

### SAV-653

- Store arm: it reads `this+0x18` as a `u16` (`L08523` loads the 16-bit field at offset 0x18) and
  calls the `u16` store helper `L08299`, then reads `this+0x1c` as a `u32`
  (`L08524` loads the 32-bit field at offset 0x1c) and calls the generic `u32` store helper
  `R0685`. `Token::Serialize` uses the same helper pair for `this+0x0e` (u16)
  and `this+0x4`/`this+0x8`/`this`/`this+0x14` (u32).
- Load arm, the mirror: `L08525 CALL R0282` (`u16` load) into `this+0x18`,
  then `L08526 CALL R0686` (`u32` load) into `this+0x1c`.
- By their serialization shape both are purely mechanical fields, not a class
  key, pointer or count.
- `Token::Serialize` has 7 direct callers (`evidence/q3-token-serialize.txt`),
  matching its shared-base-class role across the `Token`-derived classes the
  byte-map table in `formats/sav/format.md` names.

**Confidence.** High. Every instruction is disassembled to the byte in both
arms; the width, offset and helper-pair identification leaves no other reading
of the serialization mechanics.

### SAV-654

- `savdoc`'s `item()` decoder reads `Item`'s "Value" field at `recordAt+25`, the
  record-relative stream position `Token::Serialize`'s byte-map fixes as
  `this+0x1c` (file bytes 25..28, `SAV-653`).
- `ITEM-VALUE-115` states that its writer "assigns Spells parameter 21 to
  Item+1c" from a Book's first Effect.
- `Item` inherits `Token`'s fields at a fixed, unadjusted offset, the
  single-inheritance layout `SAV-636` established for `Unit`, so `Item`'s
  `+0x1c` and `Token`'s `+0x1c` are the same byte.
- Basis: a `go run ./tools/claim -k "Token"` search over every ledger
  (`evidence/q3-claim-search-token.txt`) and a direct read of `ITEM-VALUE-115`
  (`evidence/q3-claim-item-value-115.txt`). `SAV-636`'s search was scoped to the
  `sav` ledger's "Token"/"this+0x1c" vocabulary and did not reach `item.md`'s
  "Item+1c" wording.

**Confidence.** High for the byte identity, a mechanical consequence of
`Token`'s fixed layout inside `Item` rather than a corpus inference, and for
`ITEM-VALUE-115` being ✔ promoted, independently authored content. This resolves
the writer question for `Item` only; it does not assert that `Token+0x1c` means
"value" for every subclass (`SAV-656`).

### SAV-655

- `tools/savdoc -mode token-unknowns` (`evidence/token-unknowns-full.tsv`)
  decodes `+0x18`/`+0x1c` for every position-verified record of the 10 classes
  the corpus contains (`evidence/census-summary.tsv`'s
  `token_unknowns_class_count`).
- Always 0: `Armor` (1077), `Effect` (101), `Item` (70), `Shield` (54),
  `SpellTransport` (1), `Weapon` (792).
- Always 2: `Building` (31), `Sack` (31).
- `Human` (516) and `Unit` (497): 2 on 1,008 of 1,013 records and 0 on the other
  5 (`evidence/census-summary.tsv`'s `token_unknowns_unk18_by_class_*` rows).
- The five exceptions sit in exactly two saves:
  - two `Unit` records in `en/game0009.sav`, sharing a `reference` value with 7
    other same-reference records that keep the typical value 2;
  - three `Human` records in `EXP-0261-owner-runs/game0005.sav`, sharing a
    different `reference` value with 1 other same-reference record that keeps 2.
- A shared `reference` alone therefore does not explain the split.

**Confidence.** Medium. A corpus population reading. The 0/2 split by class is
exhaustive over the decoded corpus, but no writer of `+0x18` outside
`Token::Serialize` (`SAV-653`) was located, so the value's semantic, what
separates the two groups of classes and the five exceptions, is not established,
only its shape.

**Unknown.** The condition that produces 0 instead of 2 for the five exceptions.

### SAV-656

- The values range 0..121319 with one sentinel, `0xffffffff`.
- Per class (`evidence/census-summary.tsv`'s
  `token_unknowns_unk1c_nonzero_by_class_*` rows, all ten classes):
  - 0 on every record: `Building` (all 31) and `SpellTransport` (its single
    record);
  - non-zero on every record: `Unit` 497/497, `Sack` 31/31, `Shield` 54/54,
    `Armor` 1077/1077, `Item` 70/70;
  - non-zero on some: `Weapon` 741/792, `Human` 393/516, `Effect` 65/101.
- The `0xffffffff` sentinel appears only on `Item` and nowhere else in the
  corpus (`evidence/token-unknowns-full.tsv`), on roughly half the Item records
  (34/70). This is consistent with the `ITEM-VALUE-115`/`SHOP-CONSUME-073`
  reading that an item's stored value is computed only once it reaches a shop
  shelf and stays at a sentinel otherwise.
- Only `Item`'s meaning for the field is independently published (`SAV-654`).
  The claim asserts nothing about what `Token+0x1c` means for `Unit`, `Human`,
  `Sack`, `Building` or `Effect`; the shared field could carry a different,
  subclass-specific meaning at the same byte in each class.

**Confidence.** Medium. A corpus population reading. The sentinel/class
correlation is exhaustive over the decoded corpus but does not by itself
discriminate a single semantic across all ten classes, only Item's published
one.

### SAV-657

- A `go run ./tools/claim -k "L00285"` search
  (`evidence/q1-claim-search-L00285.txt`) finds `UNIT-GATE-012`. Its
  `R0967`/`R1304` are the constructor/allocator pair `SAV-646`
  traces for `world+0x118`, and it calls `[L00285]` "the server singleton."
- `formats/sav/format.md`'s `+0x3f` table row and `formats/session/format.md`'s
  table cite `UNIT-GATE-012` for `+0x84`: `formats/sav/format.md:243` beside
  "difficulty (`world+0x84`)," and `formats/session/format.md:118` beside
  "server singleton". The identification is inferable from two citations of one
  claim, but neither page's prose states it, and `UNIT-GATE-012` never touches
  `+0x118` or the `this+0x14` sentence.
- This claim adds three things `UNIT-GATE-012` does not cover:
  - an explicit prose statement of the identification;
  - confirmation through different code (the trailer-block, log-gate and toggle
    sites, not the difficulty and campaign code `UNIT-GATE-012` traced). Every
    `world+0x118` dereference disassembled loads its base from `[L00285]`
    directly, for example
    `L00286` loading the singleton pointer from `[L00285]` and `L08527` loading the field at offset 0x118 of it
    (`evidence/q1-world118-and-L07886.txt`). Inside the document programme, the
    one cached local (the frame local at -0x90) is the local the load arm dereferences for
    the `+0x118` trailer thirty lines later
    (`L08528` loads the frame local at -0x90; `L08040` loads the field at offset 0x118 of that object);
  - the caution that `formats/session/format.md`'s table also labels a
    different, larger address "world" (`[L04624]`, `0xa4558` bytes, built by
    `R0118`), so `formats/sav/format.md`'s "world" must not be conflated
    with that page's "world" row.
- The dead-actor citation of the "`this+0x14` points to the world" sentence was
  wrong. `L08529`/`L08530` (loads of the frame local at -0x90 and then of the field at offset 0x14 of it) does
  not reach the dead-actor list: it continues to `L08531`, a load of the field at offset 0x4 of that object,
  one of four list-head fields on a different object.
- `this+0x14` is not `[L00285]` and not the world. the push of 0x174 at `L08532`
  sizes and the call to `R0967` at `L07364` constructs `[L00285]`'s `0x174`-byte
  object. Inside that constructor, the store at `L05907` of the constant L05906 into offset 0x14 of the object stores a
  literal address at `this+0x14`, not a pointer the constructor allocates
  (`evidence/q0-this14-object-L05906.txt`).
- `L05906` is a separate, statically constructed object with exactly three
  whole-image dword references (`cursorsurf -xrefs L05906:L13160`): a
  constructor call at `L08533`/`L08534` (`CALL R1615`), an
  `atexit`-registered destructor call at `L08535`/`L08536`
  (`CALL L08537`), and the `L05907` store.
- `R1615` zeroes exactly four dword fields, in the write order `+0x8`,
  `+0xc`, `+0x4`, `+0x0`. The programme reads each through the same `this+0x14`
  indirection:
  - `+0x0` at `L08538`/`L08539`, feeding a call to `R1590`;
  - `+0x4` at `L08530`/`L08531`, the pair `formats/sav/format.md` miscited
    as the dead-actor access, feeding a call to `R1122`;
  - `+0x8` at `L08540`/`L08541`, feeding a call to `R1592`;
  - `+0xc`, the dead-actor access, at `L08542`/`L08543`
    (a load of the field at offset 0x14 and then of the field at offset 0xc of the object it gives), followed by the vtable-slot-0
    indirect call `L08544`/`L08331`.
- `formats/sav/format.md`'s names for the other three fields, Building `+0x0`,
  SpellEffect `+0x4` and Sack `+0x8`, are not re-derived. This claim confirms
  the four-field structure, the read sites and the citation fix, not what
  `R1590`/`R1122`/`R1592` consume.

**Confidence.** High for the `[L00285]` identification of `world+0x118`: six
disassembled `[L00285]`-then-`+0x118` dereferences across four routines
converge without exception, independent of and in different code from
`UNIT-GATE-012`'s evidence. High for `this+0x14` being a distinct object at
`L05906`, its construction with three whole-image references, and the
four-field structure with the corrected `+0xc` dead-actor citation: every
address is disassembled, and the citation error shows in the same instructions
the earlier citation used, decoded one step further. Medium for naming the other
three fields Building, SpellEffect and Sack: those labels are inherited, not
re-derived from what `R1590`/`R1122`/`R1592` do. The claim does not
assert that `formats/sav/format.md` uses "world" consistently for `[L00285]`
everywhere; only `world+0x118` and the `this+0x14` structure were checked.

### SAV-658

- All three get the construction-time literal store `R1289` makes beside
  slot 93's (`SAV-650`): `90=0` at `L08506`, `91=0` at `L08507`, `92=1` at
  `L08508`, each a literal store of a value or immediate at offset `off` of the object.
- The sweep finds no further literal, scaled-index or immediate touch anywhere
  in `.text` for slots 91 or 92.
- Slot 90 has two further runtime sites, each pairing its increment with
  `session+0xbcb0`'s increment in the same block
  (`evidence/q2-session-slot93.txt`, `evidence/q2-slot-sweep-100.txt`):
  - `L08545`/`L08546` (loads of the fields at offsets 0xbcb0 and 0xbe9c;
    an increment of each; stores back to the same two fields), one case
    of `R0262`'s 34-way jump-table dispatch, which has three direct `E8`
    callers: `L06060` in `R0431`, `L06061` in `L08547`, and `L01870`
    in `R0428`, the pattern evaluator `SAV-651` names;
  - `L08548`/`L08549` (an increment and a store to the field at offset 0xbcb0 immediately before,
    then a load of the field at offset 0xbe9c, an increment and a store back to it, and a return cleaning 4 argument bytes), in a
    handler block at `L08550` that belongs to the function at `R1616` (no
    separate prologue), with exactly one direct `E8` caller, `L08551`, inside
    the check evaluator `R0433`.
- Slot 90's two runtime writers are thus reached by two structurally different
  mechanisms: one from the tick and pattern-dispatch cluster that holds slot
  93's unconditional site, one from the check evaluator's case dispatch. Slot
  93's two runtime writers are both reached from the tick-advance cluster
  (`SAV-650`).

**Confidence.** High for the population and the three write sites: the same
decode-anchored, operand-verified sweep as `SAV-650`, covering all 100 slots.
Medium for the completeness of slots 91 and 92 beyond `.text`: the sweep does
not see an indirect or vtable-dispatched write.

**Unknown.** Whether slot 91 or 92 is ever read or overwritten at runtime by a
mechanism outside this decode; it is not ruled out.

## Player settings block, Diary arrays, and container tail

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-662 | Across the full preserved corpus, `Player+0x30`'s 32-byte block holds only `AI-FORM-037`'s formation byte at `+0x1f`; the other 31 bytes are zero in all 410 records. | High | ✔ promoted | [EXP-0313](../experiments/EXP-0313-player-diary-container/) |
| SAV-663 | Three functions beyond `AI-FORM-037`'s setter share its `[this+0x30]`-relative `+0x1f` access and have no direct caller or stored-dword reference anywhere in the image. | Medium | ✔ promoted | [EXP-0313](../experiments/EXP-0313-player-diary-container/) |
| SAV-664 | `Player+0x2c` (u16) is `1 << (slot mod 16)` for the participant's registration slot, computed only when it owns the map (`Player+0x28==0`); the corpus shows only the value for slot 1. | High / Medium | ✔ promoted | [EXP-0313](../experiments/EXP-0313-player-diary-container/) |
| SAV-665 | `Player+0x50` (u32) is 0 in every preserved record, matching its constructor default; the bounded search run located no writer. | Medium | ✔ promoted | [EXP-0313](../experiments/EXP-0313-player-diary-container/) |

### SAV-662

- `AI-FORM-037` establishes the block's allocation, its constructor default
  (`R0202` zeroes the block, then writes `+0x1f=2`), and that
  `Player::Serialize` (`R0415`) hands it to `R0204` as a raw
  `CArchive::Write` (`R0271`)/`::Read` (`R0272`) call over a literal
  `0x20` bytes (`L00567`/`L00568`), gated only by `IsStoring` (`R1123`).
  Neither arm has a per-field dispatch.
- Call site in `Player::Serialize`: the continuation both arms reach after
  `SAV-PLAYER-028`'s seventeen-field run (`L08552..L08553`) calls an
  embedded object at `Player+0x24` (`R1340`, out of scope), then
  the call to `R0204` with `this` loaded from `[this+0x30]` (`L08554..L08160`),
  then dispatches `Diary` through a virtual call on `Player+0x40`'s vtable
  slot 8. This confirms the experiment brief's citation and explains why a
  direct-caller search for `Diary::Serialize` (`L08081`) finds zero E8
  callers.
- Corpus: `tools/savdoc -mode player-diary-container` captures the block
  verbatim in `player-block.tsv` for all 410 preserved `Player` records across
  45 distinct saves (`evidence/input-manifest.json`). `evidence/summary.tsv`'s
  `block32_byte_0x*` rows (`block32_bytes_ever_nonzero_count=1`) show `+0x1f` is
  the only byte of 32 ever nonzero (408 records at 2, 2 at 0); the remaining 31
  are 0 in every record, with no exception.
- Brief Q1 also asks for the routines that write and read each of the 31
  always-zero bytes outside `Serialize`. That is answered only for the bytes the
  image names outright:
  - the block constructor `R0202` (the allocation site `SAV-663` cites)
    writes `+0x8` and `+0x9` individually (`L08555`/`L08556`), redundant
    after its own `REP STOSD` zero-fill. That is the compiler pattern for a
    member initialized in its own right, not filler: those two bytes have a
    known writer and no known reader;
  - `+0x1e` has three readers that share the block's `[this+0x30]` indirection
    through a second object (`L08557`, `L08558`, `L08559`; each `RET 8`,
    16-byte aligned, padded with no-op bytes, zero direct E8 callers). Each uses the byte as
    a subscript into a stride-50 pairwise matrix (`base + 50*a + b`). No writer
    of `+0x1e` was located;
  - the remaining 28 of the 31 bytes were not searched individually for a writer
    or reader outside `Serialize`. The corpus-zero result is a population fact,
    not a reachability search, and does not stand in for one.

**Confidence.** High for the call-site location and the archive mechanism: both
arms are disassembled, and the mechanism admits no field dispatch to discover.
High for the corpus census, exhaustive over the full preserved population. That
census is a fact about the preserved saves, not proof that the other 31 bytes
are structurally inert: a byte written only by an authored surface absent from
every preserved save, as `AI-FORM-037` found for the map-corpus side of `+0x1f`,
would look identical here.

### SAV-663

- `R1617` (`R1617..L08560`) is a toggle: it reads `+0x1f`, computes
  its logical negation (a test of the byte for zero and a set-if-zero into the byte), writes it back, then re-reads
  the byte into the return value; three `+0x1f` instructions.
- `R1618` (`R1618..L08561`) is a plain getter, one instruction.
- Both are 16-byte aligned, follow the setter `R0080` across NOP padding,
  and use the setter's two-instruction pattern: a load of the field at `this+0x30`, then a
  byte access at offset 0x1f from it.
- `R1619` (`R1619..L08562`) reaches `+0x1f` through one more level
  of indirection: loads of the first stack argument and then of the byte at offset 0x8 of it;
  loads of the field at offset 0x38 and then, of that object, the field at offset 0x30; a store of the byte to offset 0x1f of the last object, one instruction
  touching `*(*(arg+0x38)+0x30)+0x1f`. It is also 16-byte aligned and padded with no-op bytes
  on both sides.
- `AI-FORM-037`'s image-wide `EnumRefs` search on the pattern `.*\+ 0x1f\].*`
  reports 8 hits in 8 owners and attributes the remainder not at a stack-relative `+0x1f` to
  the constructor default, the one setter and two group-move reads: four hits,
  none of them the five `+0x1f` instructions of these three functions.
- A full-image `-disp 1f` sweep, a second, complementary instrument that matches
  disp8 ModRM forms plus byte-register opcodes, returns 18 raw hits. Nine are
  this population: `AI-FORM-037`'s four (constructor default `L00561`, setter
  `L00563`, group-move reads `L00167`/`L00169`) and the three functions'
  five (`L08563`/`L08564`/`L08565` the toggle, `L08566` the getter,
  `L08567` `R1619`). The other nine are unrelated bases, frame-relative
  forms (stack-pointer or frame-pointer relative `+0x1f`) or non-byte-width forms, not examined
  further.
- `cursorsurf -callers R1617`, `-callers R1618` and `-callers R1619` each
  report zero direct E8 callers. `-xrefs` sweeps of `R1617:R0181` and
  `L13161:L13162` find no dword-literal reference to any of the three addresses
  anywhere in the image; a call through a stored function pointer would still
  show here as a literal.
- This does not contradict `AI-FORM-037`: its search covers what a structural
  reference enumeration surfaces, and all three functions take a generic pointer
  parameter, not a typed `Player*`. It narrows "the byte's whole surface is four
  instructions" to six instructions across four functions carrying the pattern
  (the setter one, the toggle three, the getter one, `R1619` one), three
  of the four unreachable by any means these tools can see.

**Confidence.** Medium. The structural fact (identical dereference pattern,
shared alignment, NOP-separated placement) is read directly from the
disassembly. The completeness claim, "unreferenced anywhere", is bounded by what
a direct-E8-caller sweep and an image-wide dword-literal sweep can see; a call
reached only through a computed or relocated address is not excluded.

### SAV-664

- `R0424` is called from the ALM-roster-to-session copy `R0423` at
  `L08105`, immediately before that routine's `Player+0x44` colour write at
  `L02129`.
- Slot assignment: it scans the existing roster (`L01800`/`L01801`-driven
  iteration), builds a 32-bit occupancy mask from each existing `Player+0x04`,
  finds the smallest slot 1..31 whose bit is clear, and stores it to the new
  `Player+0x04` (u16, `L07240`). Two branches are not exercised:
  - a fallback assigns from a monotonic counter at the roster object's `+0x20`
    when the scan finds nothing and `Player+0x04` remains 0
    (`L08568..L08569`);
  - the scan starts at candidate 16 instead of 1 when `Player+0x28==0` and a
    named-lookup call `L08570` on `Player+0x18` succeeds (`L08571`).
- Mask: a nonzero `Player+0x28` jumps to a store of zero to `[this+0x2c]` (`L08572`).
  Zero falls through to a signed mod-16 (the compiler's signed-modulo idiom,
  `L08573..L08574`) of `Player+0x4` and a left shift of 1 by that value, stored as
  `Player+0x2c` (`L08575`).
- Defaults: `Player+0x2c` is 0 in the constructor (`R0201` `L08576`).
  `Player+0x28`'s constructor default is 1 (`L08000`), later overwritten per
  record by the same ALM-copy routine from the roster's `+0x04`
  (`ALM-GRP-041`/`ALM-PLAYER-069`).
- Corpus, 410 records (`evidence/summary.tsv` rows `f_0x2c_value_*`,
  `cross_f44_f2c_*`): 314 are 0 and 96 are exactly 2, which is
  `1 << (1 mod 16)`. Every nonzero record has `Player+0x4==1` in effect, and the
  value correlates 1:1 with `Player+0x44==2` (`ALM-GRP-041`'s
  campaign-protagonist colour slot).
- No record exercises the 16-start or counter-fallback branches, and no nonzero
  value but 2 occurs.

**Confidence.** High for the mechanism, read in full at both write sites. Medium
for "the corpus never differs from slot 1": a population fact over 410 records
across 45 saves, not proof that the other branches are unreachable in play. The
two unexercised branches are read from the same disassembly, not inferred from
their absence.

### SAV-665

- `R0201` (`Player::Player`) sets `Player+0x50=0` at `L08577`.
- `evidence/summary.tsv`'s `f_0x50_value_0` row: all 410 preserved `Player`
  records, across 45 distinct saves and both EN/RU installs, read 0.
- No dedicated writer search (a `-disp 50` sweep) was run. `+0x50` is too common
  a small-struct displacement for an unrestricted image-wide sweep to attribute
  reliably, the caution `SAV-633`/`SAV-636` name for equally common offsets, and
  the experiment's budget went to the fields with an observed nonzero corpus
  value.

**Confidence.** Medium. The corpus fact is exhaustive and exact. The negative is
explicitly bounded: no writer search was run, which is not "a writer search
found none".

**Unknown.** Whether code the corpus does not exercise ever writes the field to
something other than 0.

## Player+0x58 and the Diary arrays

`SAV-673`–`SAV-677` are retired unused: the EXP-0313 evidence supports eleven
claims from its `SAV-662`–`SAV-677` range, not sixteen. `Player+0x34` is not
among them: `SAV-HERO-059` resolves it at High confidence (the field naming the
participant's own starting character, resolved through the identity map
`SAV-669` also uses), and the EXP-0313 corpus re-confirms that resolution
without a new fact, so it is a citation, not a sixteenth row. `Player+0x44`'s
meaning is likewise resolved (`ALM-GRP-041`, `PAL-SHADE-012`,
`ALM-PLAYER-069`); the only addition there is the `Player+0x28`/`Player+0x2c`
correlation folded into `SAV-664`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-666 | `Player+0x58` defaults to 95 and is overwritten to 0 or 50 on exactly the protagonist population with nonzero `Player+0x2c`, plus two zero-colour records; SAV-726 names a command writer. | Medium | ✔ promoted (amended) | [EXP-0313](../experiments/EXP-0313-player-diary-container/) |
| SAV-667 | A new `Diary` sizes both arrays once from `data.bin`'s Units table through `R1620`; every load resizes them from the stream's own count, and the two counts are equal across the corpus. | High | ✔ promoted (amended) | [EXP-0313](../experiments/EXP-0313-player-diary-container/), [EXP-0320](../experiments/EXP-0320-item-effect-classes/), [EXP-0327](../experiments/EXP-0327-diary-consumers/) |
| SAV-668 | Every Diary element pair in the corpus satisfies `word[i] = 1024 - dword[i]`; the traced mutator touches only the word array, SAV-842 locates the dword writer, and the relation is a finite corpus fact. | High / Medium / Unknown | ✔ promoted (amended, partially retracted, superseded) | [EXP-0313](../experiments/EXP-0313-player-diary-container/), [EXP-0327](../experiments/EXP-0327-diary-consumers/) |
| SAV-669 | `Diary+0x2c` resolves through the `SAV-PTRMAP-035` identity map, and every resolved reference in the corpus names the Diary's own enclosing Player record, never another object. | High | ✔ promoted (amended, partially retracted) | [EXP-0313](../experiments/EXP-0313-player-diary-container/), [EXP-0327](../experiments/EXP-0327-diary-consumers/) |

### SAV-666

`Player+0x58` is a u32. Constructor default 95/0x5f: `R0201`
`L08578`.

Corpus, `evidence/summary.tsv`'s `f_0x58_value_*` and `colour2_f58_*` rows,
410 records:

- 312 stay at 95: exactly the `Player+0x44 != 2` population.
- Of the remaining 98, 96 are the same 96 records `SAV-664` finds with
  `Player+0x2c` nonzero (`Player+0x44==2`): 77 at 0 and 19 at 50.
- The other 2 are `Player+0x44==0` records, both at 0.

Bounded writer search: a `-disp 58:L13163:L06644` sweep (the module range
covering `Player::Player`, the ALM-copy routine and the slot allocator) returns
15 accesses: 4 reads (`L08579`, `L08580`, `L08581`, `L03906`) and 11
writes.

- Five writes are inside `R0451` (`L04572` storing 0,
  `L08582`/`L08583` storing 5, `L05340` storing 7, `L08584` storing 0).
  The same routine reaches `+0x8e`/`+0x90`/`+0x92`, `ITEM-LOAD-005`'s actor
  weight/capacity/load fields, so its `this` is an actor, not a `Player`, and
  none of its five `+0x58` writes is a `Player` write.
- Five writes, one each in five distinct small routines (`L02814`,
  `L02815`, `R0185`, `R0672`, `L08585`), all store 0 and were not
  individually attributed to a class.
- Only the constructor's write (`L08578`, storing `0x5f`, inside
  `Player::Player`) is confirmed on a `Player`.
- Seven of the eleven writes store 0, one of the two override values, so the
  range is not empty of other writers. None of the other ten writes was
  confirmed to touch a `Player` rather than another small object at the same
  displacement, and the override routine was not located among them.

**Confidence.** Medium. The population correlation is exhaustive and exact:
every deviation from the constructor default lands in the named
protagonist-adjacent set, with no counterexample. The bounded search is an
explicitly bounded negative, not the current writer frontier after SAV-726: no
write attributable to a `Player` was identified among the fifteen hits besides
the constructor default, and the five unattributed small-routine writes, and any
writer outside the range, are not excluded.

**Amended.** The clause that the bounded search did not locate the producer is
narrowed by SAV-726, which locates command stores `L08586`/`L08587` outside
the searched interval. The corpus and bounded-search facts stand.

### SAV-667

Three constructors share vtable `L08066`:

- The no-arg `R1566` sets `+0x2c=0` (`L08588`), then calls
  `R1620`.
- The one-arg `R1565`, called from `Player::Player` (`L08333`) with
  the new `Player` as the sole argument, stores that argument verbatim into
  `+0x2c` (`L08589`), then makes the same call.
- The copy constructor `R1550` copies both arrays element for element
  (`L08590`, `L08591`), then copies `+0x2c` from its source
  (`L08592..L08593`).

Construction-time sizing, `R1620`:

- It reads a count `n` from `L02114` (passing L02114 as `this` to the call to `R1621`,
  `L08594..L08595`). `L02114` is `L02110 + 0x8c`, the Units table of
  `world.res:data/data.bin` in `DAT-OBJ-002`'s eleven-collection map;
  `ALM-CLS-038` as amended names the same table at the same offset.
- The getter `R1621` is three instructions on that `0x14`-byte
  collection object: loading the object pointer from the frame local at -0x4, then returning the field at offset 0x8.
- It calls `CDWordArray::SetSize(n,-1)` (`R1198`) and
  `CWordArray::SetSize(n,-1)` (`R1622`), then loops `i=0..n-1` writing
  `dword[i]=0` (`L08152`) and `word[i]=0x400`/1024 (`R0130`).
- This sizes storage but does not establish one common index domain: SAV-844
  shows actor+0c has class-dependent Units/Humans producers.
- No `Add`/`SetAtGrow` call exists in either constructor or in the one live
  mutator traced (`SAV-668`). No code reached grows the arrays past their
  current size, but the current size is not fixed at construction.

Load resizing: `CDWordArray::Serialize` (`L08153`) and
`CWordArray::Serialize` (`L08154`), the store/load routines `SAV-DIARY-042`
names, resize each array on load through the same `SetSize` entry, from a count
read off the stream (`ReadCount`, `L08596`), not the construction-time size:

- dword array: `L08597` calls `L08596`, `L08598` pushes the return value of that call, `L08599` calls `R1198`,
  then `Read` (`L08600` calls `R0272`).
- word array: `L08601` calls `L08596`, `L08602` pushes the return value of that call, `L08603` calls `R1622`,
  then `Read` (`L08604` calls `R0272`).

"Sized once" therefore holds only up to a save's first load.

Corpus, `tools/savdoc -mode player-diary-container`'s `diary.tsv`, 542 records
across 45 saves:

- `dword_count` is 119 in 540 (99.6%) and 0 in 2:
  `EXP-0261-owner-runs/game9001.sav` and `EXP-0261-owner-runs/game0000.sav`,
  which that directory's `MANIFEST.md` labels "generated candidate, wave 1" and
  "ROM1 resave of game9001". They are an empty `Diary` this project's writer
  produced and its unmodified original resave, both loaded and so resized to 0
  by the mechanism above; the zero is not a property of being mission-only.
- `evidence/summary.tsv`'s `diary_world_half_dword_count_*` rows: of 71
  town-half (`world_half=false`) records, 69 carry 119 elements, like every one
  of the 471 mission-half records; only the same 2 project-written records carry
  0.
- `dword_count == word_count` in all 542 (`evidence/summary.tsv`'s
  `diary_count_mismatch_records=0`).
- `world_half=false` is the between-mission/campaign-screen shape (no map
  loaded) and `world_half=true` the active-map/mission shape (`SAV-SHAPE-023`,
  `SAV-CITY-030`).

**Confidence.** High. The sizing mechanism is read on both constructors, the
shared construction-time sizing routine and the shared serializer's load arm.
The catalog is independently named at High confidence (`DAT-OBJ-002`,
`ALM-CLS-038`). The count equality, the fixed size 119 and the provenance of the
two exceptions are exhaustive over the preserved population.

**Amended.** EXP-0320 corrected the `diary_world_half_dword_count_*` gloss,
which had `world_half=false`/`world_half=true` labelled backwards
("mission-half" for the 71 false records, "town-half" for the 471 true ones).
EXP-0320 reproduced the shape reading on the witness file:
`tools/savdoc -mode postload` on `gameversions/saves/2026-08-02/game0010.sav`, sha256
`b4e5ceb73018d0c4656515ce643712fb11d5bfa57fe9065361de5ba8156bc0ac`, reads
`world_half=false`. The counts and the sizing mechanism are unchanged. SAV-844
(EXP-0327, [`retracted.md`](retracted.md)) narrows the inference "so an array
index is a Units-table row — a unit type": Units capacity sizes storage, but the
actor+0c producer selects Units for base Unit and Humans for Human. The
capacity, default and load-size mechanisms and the corpus counts stand.

### SAV-668

- Construction defaults (`SAV-667`) are `dword[i]=0`, `word[i]=1024`, which
  satisfy the relation from an element's birth.
- `R1564` takes an actor. It rejects only when vtable `+0x30` is nonzero
  and the unsigned `+0xc` index byte exceeds `0x3f` (SAV-844). It reads
  `word[index]` through the generic accessor `R0130` and, only if the value
  is `>0` (a floor guard), decrements it by 1 and stores it back
  (`L08605..L08606`). It never touches the dword array.
- Its sole direct caller is `R1484` (one E8 call, `L08607`), traced by
  SAV-842 and SAV-843 to the dword writer and a bounded attributed-teardown
  caller.
- Corpus, `evidence/summary.tsv`'s diary-element rows over all 64260 sampled
  elements (540 records times 119 elements; the `dword_count=0` records
  contribute none): `diary_elements_word_eq_1024_minus_dword=64260`,
  `diary_elements_word_ne_1024_minus_dword=0`. There is no exception.
- 280 elements have a nonzero dword value, in six pairs only:
  `(2,1022)x109, (6,1018)x73, (5,1019)x40, (4,1020)x29, (1,1023)x19, (3,1021)x10`;
  the rest hold the `(0,1024)` default.
- Only six of the 119 indices carry a nonzero dword anywhere in the 45-save
  corpus (`evidence/summary.tsv`'s `diary_dword_index_*_nonzero_records` rows):
  index 64 on 47 records, 72 on 16, 73 on 7, 88 on 56, 96 on 79, 100 on 75
  (`diary_dword_indices_ever_nonzero_count=6`, summing to the same 280).
- The subscript follows the actor's class-dependent Units/Humans definition
  ordinal (SAV-844). SAV-843 and SAV-844 identify the bounded event and the
  original table keys.

**Confidence.** High for the exact value relation (exhaustive over the corpus,
zero exceptions) and for the mutator mechanism (fully disassembled). Medium was
the grade of the former writer-search negative, which SAV-842 and SAV-843
supersede by locating the writer and a bounded attributed-teardown caller. The
six-index census is a finite historical population.

**Unknown.** Live event completeness, full event attribution and live
scheduling.

**Amended.** EXP-0327 narrows the index wording (SAV-842, SAV-844,
[`retracted.md`](retracted.md)). "Its +0xc index byte, bounded <=0x3f" is
replaced by a conditional bound: rejection needs vtable `+0x30` nonzero as well
as an index above 63, so a Unit index 64 is admitted and a Human index 64
rejected. "An element is a per-unit-type counter" is withdrawn: index identity
follows the actor's class-dependent definition producer. The formerly open
caller `R1484` is the dword writer. The complement relation is kept as a
finite corpus fact, not a permanent law.

### SAV-669

- `R1371` reads `[L00285]+0x88` through `R1488`, the same
  global-plus-offset `SAV-HERO-059` names for `Player+0x34`'s resolver
  `R1370`.
- Construction seeds the reference: `R1565` (`SAV-667`), called from
  `Player::Player` with `this` as the sole argument, stores the constructing
  `Player`'s own address into `+0x2c`. This literal self-reference is written
  into the save stream and resolved back to whichever object the loader most
  recently constructed at that address.
- LOAD separately overwrites it from the stream and remaps it at
  `L07737..L07738` (SAV-847).
- The `savdoc` resolver checks two identity-key domains. `SAV-PTRMAP-035`
  shows that `Token`-derived objects and `Player` records both mint entries into
  one shared map, so a resolver checking only `Token` keys under-resolves;
  `player-diary-container`'s resolver checks both.

Corpus, 542 records:

- 408 carry a nonzero `+0x2c` and all 408 resolve. A file-joined cross-check
  (`evidence/summary.tsv`'s `diary_ref_self_match=408`,
  `diary_ref_other_match=0`) confirms each names exactly its own owning Player
  record.
- The population splits by owner class, not by protagonist adjacency. 410
  records are a `Player`'s own `Diary` (`L07961`'s virtual dispatch on
  `Player+0x40`, `SAV-662`): 408 with a nonzero self-reference and 2 null, the
  same two project-written `EXP-0261-owner-runs/game9001.sav`/`game0000.sav`
  records `SAV-667` names.
- 132 are actor-owned, reached through `*(Humanoid+0x1e4)` (`SAV-EMBED-039`)
  while walking that `Player`'s group tree, and all 132 are null: the no-arg
  constructor's default (`R1566`), never resolved to anything. No
  `Humanoid`'s `Diary` is the target of any `+0x2c` reference in this corpus.
- Counts in `evidence/summary.tsv`: `diary_owner_class_Player=410`,
  `diary_owner_class_Humanoid=132`, `diary_owner_class_ref_nonzero_Player=408`,
  `diary_owner_class_ref_zero_Player=2`,
  `diary_owner_class_ref_zero_Humanoid=132`. `diary.tsv`'s `owner_class` column
  carries the split.

**Confidence.** High. The write site and the resolver are fully disassembled,
and the corpus check separates the live alternative, some other object class
versus a self-reference, with zero counterexamples over every resolved record in
the preserved population. The owner-class split narrows this: every one of the
410 Player-owned records is self-resolved or one of the two project-written
nulls `SAV-667` explains, and every one of the 132 actor-owned records is null
by constructor default, not by a failed resolution.

**Amended.** SAV-847 (EXP-0327, [`retracted.md`](retracted.md)) partially
retracts the clause "written at construction, not fixed up separately at load":
Diary Serialize stores the streamed key at `L07737` and calls `R1371` at
`L07738`, and the resolver overwrites `+0x2c` with the mapped pointer or zero.
Constructor seeding and the owner-reference census stand. `tools/savdoc`'s
exporter first tagged every `Diary` record by the top-level `Player` being
walked, whether the record was that `Player`'s own or an actor's; the correction
added `owner_class` and reproduces the split above. The sentence "concentrated
in the protagonist population (129 of 134)" measured that mis-attribution (the
campaign protagonist owns the most actors, so most actor-owned Diaries fell in
his subtree) and is withdrawn.

## Item container tail

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-670 | The container tail `u32 +0x1c`/`u32 +0x20` at `*(Sack+0x40)` and, when present, `*(Unit+0x7c)` is `ITEM-CONT-004`'s container class, sharing no fields with `Item+0x1c`/`Item+0x20`. | High | ✔ promoted | [EXP-0313](../experiments/EXP-0313-player-diary-container/) |
| SAV-671 | The container's `+0x1c` insertion index defaults to a sentinel that makes every ordinary insert append, and both `ITEM-CONT-004` non-default writers, pick-up and equip, occur in the corpus. | High / Medium | ✔ promoted | [EXP-0313](../experiments/EXP-0313-player-diary-container/) |
| SAV-672 | The container's `+0x20` weight sum defaults to 0, is copied by a transfer helper and zeroed by a reset pair with `+0x1c`, and in the corpus never nears `ITEM-LOAD-005`'s `0xfa00` threshold. | High | ✔ promoted (amended) | [EXP-0313](../experiments/EXP-0313-player-diary-container/), [EXP-0320](../experiments/EXP-0320-item-effect-classes/) |

### SAV-670

- `ITEM-CONT-004` (active, High) publishes the container's shape:
  `R0403` is the whole constructor (`CObList::CObList` on `this`,
  `+0x1c=0x2710`/10000, `+0x20=0`, object size `0x24` bytes, `callto:R0403`
  15 hits/14 owners). `+0x1c` is not a capacity but the index the next `Add`
  inserts at; the pick-up order (`L04608`) and the equip arm (`L04609`)
  rewrite it before an add.
- `ITEM-SAVE-014` separates the container from `Item` at the serializer: "the
  container's serializer is `R0952`," distinct from `Item::Serialize`
  (`R0949`). `SAV-CITYSTORE-516` names the container's `+0x1c`
  "insertion index" and `+0x20` "stored load."
- This claim adds the field-set separation from `Item`. `Item`'s constructor
  `R0884` writes a disjoint set
  (`+0x3c/0x40/0x42/0x44/0x45/0x46/0xc/0x4a/0x8/0x4c`,
  `experiments/EXP-0079-items-and-sacks/evidence/listings-item.md`), neither
  `+0x1c` nor `+0x20`.
- On `Item` the two offsets name unrelated fields. `Item+0x1c` is
  `ITEM-VALUE-115`'s value/spell-id scalar (a Book's first Effect's `+0x40`, or a
  Scroll's summed formula). `Item+0x20` is the counted `CObject*` Effect list of
  `SAV-MEMBER-036` and `ITEM-SAVE-014` (`R0951`), a variable-length list
  tested for emptiness by `ITEM-STARFLAG-096`. Neither is a plain `u32` scalar
  like the container's `+0x1c`/`+0x20`.
- `SAV-656`, the second field the brief names beside `ITEM-VALUE-115`, is in no
  ledger in this base (`go run ./tools/claim SAV-656` returns "not in any
  ledger"), so the brief's premise that two existing claims jointly name an
  `Item+0x1c`/`+0x20` pair does not hold in this base.
- A direct-E8-caller census (`cursorsurf -callers R0403`: 16 hits over 15
  distinct owners, `L08608` calling twice) gives one hit more than
  `ITEM-CONT-004`'s `EnumRefs callto:` census (15 hits/14 owners).

**Confidence.** High. Two constructors, the container's (published by
`ITEM-CONT-004`) and `Item`'s (read here), write disjoint field sets on two
objects identified by their own Serialize entry points (`R0952` versus
`R0949`). The live alternative, silently the same field pair, is excluded
by the constructors, not by the offsets' numeric coincidence.

**Unknown.** The single site on which the two caller censuses disagree was not
isolated; it is an open reconciliation, not a corrected number.

### SAV-671

- The default is `0x2710` (10000, `R0403` `L08609`, `ITEM-CONT-004`).
  `Add`'s insertion helper (`R0929`→`R0930`, `ITEM-CONT-004`)
  reads it as the target index. The generic insert primitive branches
  `index<count ? insert-before : AddTail`, so a sentinel above any realistic
  item count appends on every ordinary insert. This refines
  `SAV-CITYSTORE-516`'s "insertion index" naming and matches `ITEM-CONT-004`'s
  "not a capacity — it is the index the next `Add` inserts at."
- A transfer helper `R1623` copies `+0x1c` (and `+0x20`) verbatim
  between two container instances (`L08610..L08611`). A reset pair
  `R1140` zeroes both fields (`L08612`). Both lie in the sack/item
  module `L08613..L08614`.
- Pick-up: the ground-pickup write site is `ITEM-CONT-004`'s pick-up order
  (`L08615..L04608`), not a second writer. It sits inside the actor
  command-dispatch function spanning approximately `L08616..L08617`. It
  writes `[actor+0x7c]+0x1c` (the actor's own container, not the source Sack's)
  from a sign-extended 16-bit value read off a command-associated record, sets
  the same actor's `+0x50=2` and exits to a shared dispatch point, a mechanism
  `ITEM-CONT-004` did not publish.
- Equip: `L08618..L04616` reads the actor's container at `[actor+0x7c]`
  (`L08619`), writes the insertion index at `+0x1c` (`L04609`), then
  reaches the call to `R0930` (`L04616`) through a reloaded `[actor+0x7c]`. This is
  the `Add()` call of `ITEM-CONT-004`'s equip arm, with the same shape as the
  pick-up writer.
- Corpus, `evidence/summary.tsv`'s `container_tail_0x1c_*` rows, 3755 records:
  3667 (97.7%) hold the sentinel. All 88 exceptions are on `Human`-class owners
  (0 on `Sack`, 0 on plain `Unit`), with values 0, 2, 3, 5 or 8.

**Confidence.** High for the sentinel-forces-append mechanism (`ITEM-CONT-004`'s
reading, independently confirmed), for the transfer and reset call sites (fully
disassembled) and for the corpus population (exhaustive). Medium for the
pick-up writer's downstream consumption: the write site is `ITEM-CONT-004`'s
pick-up order and the `+0x50=2` latch and dispatch exit are read, but what
consumes `+0x50` afterward is not traced.

### SAV-672

- `+0x20` is `ITEM-LOAD-005`'s corrected reading of `HERO-SIGHT-007`,
  maintained by `ITEM-STACK-003`'s five weight-times-count sites, which update
  `Σ (s16)+0x4a × (u16)+0x42` into the sum
  (`L04595, L04596, L04597, L04598, L04599`, the multiply
  instructions).
- A bounded `-disp 20:L08613:L08614` sweep lands on the adjacent
  store-to-`+0x20` instructions in the same basic blocks and corroborates all
  five: `L08620/L07887`, `L08621/L08622`, `L08623/L07889`,
  `L08624/L07890` (verified by full disassembly at `L08625:L08626`),
  and `L08627/L07891` in `R1522`. `claims/retracted.md`'s
  `HERO-SIGHT-007` row names the same sites independently ("written at four
  sites inside the container class — `L07887`, `L07888`, `L07889` (adds)
  and `L07890`, `L07891` (subtracts)"). No site is added.
- The constructor default is 0 (`R0403` `L08628`), published by
  `ITEM-CONT-004` as `+0x20=0` and cited, not re-derived.
- A transfer helper `R1623` copies `+0x20` verbatim between two
  container instances (`L08629..L08630`), alongside `+0x1c` (`SAV-671`). A
  reset pair `R1140` zeroes both fields together (`L08631`). Both lie in
  the `L08613..L08614` sweep range.
- Corpus, `evidence/summary.tsv`'s `container_tail_0x20_max=224` over 3755
  records: the maximum, 224, is three orders of magnitude under
  `ITEM-LOAD-005`'s `0xfa00` (64000) branch threshold, so no preserved save
  exercises that branch's flat-32000 load substitution.
- Population per class and world half, `container_owner_class_world_half_*`
  over the 3755 records the two fields share: `Human` 69 town-half / 1138
  mission-half, `Sack` 307 mission-half (0 town-half), `Unit` 117 town-half /
  2124 mission-half. `SAV-671`'s 88 non-sentinel `+0x1c` exceptions all sit in
  the 1207-record `Human` population, none in `Sack` or plain `Unit`.
- `world_half=false` is the between-mission/campaign-screen shape (no map
  loaded, so no `Sack` record can exist there, consistent with 0 `Sack` records
  at `world_half=false`) and `world_half=true` the active-map/mission shape
  (`SAV-SHAPE-023`, `SAV-CITY-030`).

**Confidence.** High for the transfer and reset sites (each fully disassembled)
and for the corpus maximum and population (exhaustive). The claim adds
population and construction/lifecycle facts to an already-High mechanism; it
cites `ITEM-CONT-004` for the construction default and does not re-derive or
extend `ITEM-STACK-003`'s weight-times-count arithmetic.

**Amended.** EXP-0320 corrected the `container_owner_class_world_half_*` gloss,
which had `world_half=false`/`world_half=true` labelled backwards. At
`80670a95` it read "`Human` 69 mission-half / 1138 town-half, `Sack` 307
town-half (0 mission-half), `Unit` 117 mission-half / 2124 town-half", the same
inversion, in the same direction, as `SAV-667`'s
`diary_world_half_dword_count_*` gloss. EXP-0320 reproduced the shape reading
on the witness file: `tools/savdoc -mode postload` on
`gameversions/saves/2026-08-02/game0010.sav`, sha256
`b4e5ceb73018d0c4656515ce643712fb11d5bfa57fe9065361de5ba8156bc0ac`, reads
`world_half=false`. The counts and every other fact are unchanged.

## Guarded trailer-toggle command reach

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-694 | Command `46` parameter `80` reaches `R0440`, which behind gates including unsigned `Player+68` > `50` toggles trailer dwords 0 and 1 (subcommands `3` and `19`); no ordinary UI emitter is established. | High / Unknown | ✔ promoted (amended, branch candidate) | [EXP-0315](../experiments/EXP-0315-world-head-trailer/) |

### SAV-694

- Command `46` parameter `80` resolves its Player and calls `R0440`.
- The callee requires unsigned `Player+68` greater than `50`, a decisive
  subset of the route's conditions. `byte[cmd+4]==0` and global
  `[L00004]!=0` also gate it, and the Player lookup must return non-null.
- Subcommand `3` toggles trailer+0 and subcommand `19` toggles trailer+4, the
  first two dwords of the 400-byte trailer, through global `L00285` and
  world+118. Each maps zero to `1` and every
  nonzero value to `0`.
- Twelve local instruction vectors, gate bytes `50`/`51` only, stop before
  logging.
- The callee's subcommand `7` prints the engine's debug-command help.
  `functions.tsv` retains its string push at `L08632` of L08633, reading
  " <Alt-t>     Script tracing on/off\n", the engine's own name for subcommand
  `19`'s toggle.
- Neither an ordinary UI emitter nor a producer of the required gate byte is
  established.
- Closure evidence: the experiment's `RESULT.md` and `evidence/closure/`.

**Confidence.** High for the decoded selector tables, the static incoming call
and the local writes. "Unsigned" rests on the `jbe` mnemonic, not on the vector
population, which tests only gate bytes `50` and `51`. The subcommand `7` help
string narrows the activation Unknown but does not resolve it, since no key
handler is traced.

**Unknown.** Ordinary runtime activation, and the other `98` dwords.

**Amended.** This claim falsifies SAV-647's "0 direct E8 callers" sentence for the same callee
(`claims/retracted.md`); SAV-647's toggle sites, paired strings and consumer
sites otherwise stand.

## Primitive encoding boundaries and nonminimal-reader domain

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-758 | The primitive count and CString readers accept nonminimal successful-IO prefixes with no shortest-form rejection, and the CString writer threshold is `0xfffe`, not the generic `0xffff`. | High / Unknown | ✔ promoted | [EXP-0319](../experiments/EXP-0319-sav-generality/) |

### SAV-758

- Original `L08596` returns 1 from a u16 `0xffff`/u32=1 prefix, consuming 6
  bytes.
- `L08634` returns 1 from ff/u16=1, consuming 3 bytes, and from ff/u16
  `0xffff`/u32=1, consuming 7.
- No shortest-form rejection occurs in these selected reader branches.
- The CString writer threshold is `0xfffe`, unlike the generic count threshold
  `0xffff`: ff/`0xfffe` is a Unicode sentinel, not length 65534.
- Forty-seven assertions execute original branching instructions with only
  primitive archive IO cuts. SAV-WLIST-040 and SAV-FULLREAD-252 supply the
  broader grammar.
- No complete-SAV acceptance, allocation safety or normal writer population is
  implied.
- Closure evidence: the experiment's `RESULT.md` and `evidence/closure/`.

**Confidence.** High for the bounded primitive instruction paths and the
boundary and nonminimal vectors.

**Unknown.** Full reader, writer and runtime generality.

## Player+58 percentage producer and parameter-map correction

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-726 | Command `0x46` parameter 3 writes Player+0x58 through arm `L07740`: `0..2` stores `(2-x)*50`, `3..100` stores `x`, and any other value keeps the old value. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0317](../experiments/EXP-0317-campaign-player-residuals/) |

### SAV-726

- The receiver is the Player resolved by `R0079`; the input is the signed
  dword at `cmd+0x0e`.
- Both stores (`L08586`, `L08587`) and the no-store range (`x < 0` or
  `x > 100`) join at `L08635` and run the same actor walk unconditionally,
  dispatching each member's virtual `+0x50` derive at `L08636`.
- The original selector table assigns parameter 1, not 3, to the withdraw
  calls `L00508`/`L00509`.
- `UNIT-DERIVE-003`'s unit derive `R0836` and `HERO-MP-006`'s hero arm
  publish the same virtual slot consuming Player+0x58 through `actor+0x14` as
  the percentage that scales `manaMax` into the AI's mana floor. This claim adds
  the producer side of that field, not an independent typing of the consumer
  side.
- `SAV-666`'s bounded missed-writer surface is narrowed by this route, not
  closed: the five unattributed small-routine writes and any writer outside the
  old searched interval are not excluded.
- Player serialization is raw transport, not this command normalization.
- `SAV-666`'s corpus values are two of this law's three outputs: selector 1's
  50 and selector 2's 0 cover 98 of 410 records exactly, the constructor
  default 95 covers the remaining 312, and selector 0's 100 is absent from the
  corpus.
- Closure evidence: the experiment's `RESULT.md` and `evidence/closure/`.

**Confidence.** High for receiver resolution, the selector mapping and the 33
instruction vectors, reproduced value for value under an independent emulator
plus a wider synthetic sweep beyond the committed set.

**Unknown.** Whether any ordinary UI command emission reaches this arm; which
classes other than the unit and human families reachable through `Player+0x20`
change the `+0x50` derive target; whether selector 0's 100 output ever reaches a
shipped save; and `Player+0x50` itself, a distinct residual field, not the
derive slot, which was not located.

## LOAD self-pointer repair and authored register constants

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-710 | `Session+0xa9c0` is a saved self-pointer that LOAD repairs through thunk `R1485`; it is not the register-array base. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0316](../experiments/EXP-0316-session-residuals/) |
| SAV-711 | A resolution-rejected normal check does not advance the shared compiled-register cursor, so each rejection shifts every later check's and mission variable's slot down by one. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0316](../experiments/EXP-0316-session-residuals/) |

### SAV-710

- Serialize `R0270` transports the field inside the 2,508-byte raw span from
  `+0xa9bc`, at session-stream offsets 1852..1855.
- Document LOAD calls `R1485` at `L07741`. That complete two-instruction
  thunk writes `this` to `[this+0xa9c0]` and returns without touching `0xbd34`.
- Adjacent `L07742` is a separate indexed-register writer; its existence is
  not a call edge from the thunk.
- Closure evidence: the experiment's `RESULT.md` and `evidence/closure/`.

**Confidence.** High for the raw-span containment, the explicit document call,
the complete thunk and two state assertions.

**Unknown.** The transitive lifecycle, and whether LOAD reaches the adjacent
writer.

### SAV-711

- `R0067` keeps the cursor at the frame local at -0x9c (zeroed at `L08637`; `local_a0`
  in `MISSION-SLOT-008`'s decompilation).
- `L06631` comparing the frame local at -0x1ac with 0x1 and branching to `L06632` when not equal gates the increment at
  `L06633`. The rejecting parameter-lookup failure writes `0` into its own map
  entry instead.
- This settles `MISSION-SLOT-008`'s open question of whether a validation
  failure can shift the assignment: it can, downward, one slot per rejection.
- The opcode `0x10002` arm writes node+48 to `session[0xbd34+4*current_slot]`
  at `L07782` and advances the same slot at `L08638`. Opcode `0x10003`, a
  sibling opcode the binder's first pass names at `L08639`, bypasses both.
- World LOAD reaches this binder call (`L01841`) after the Sack read
  (`L07781`) and, when the global session pointer `[L00004]` was still null,
  after the session `Serialize` restore at `L07069`. Otherwise
  `L08640` comparing `[L00004]` with zero and `L08641` branching to `L08642` when not equal skips construction
  and restore. `TRIG-COND-003` and the EXP-0252 overturn of `TRIG-SAVE-008`
  (`claims/retracted.md`) publish the same three-address order.
- This is an authored-map constant, not a proved live-object rederive reached
  through `R1485`.
- Closure evidence: the experiment's `RESULT.md` and `evidence/closure/`.

**Confidence.** High for the selected static LOAD edge (including its
null-guarded reconstruction condition), the complete binder control flow, the
resolution-rejection guard and 72 local branch assertions.

**Unknown.** Complete LOAD execution, arbitrary indices, all callbacks, and
whether any shipped map's own build contains a rejection.

## Token recipient-mask semantics

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-678 | Token+0x18 is a 16-bit recipient publication mask in the located emission paths, both set and cleared there, not a class discriminator. | High / Medium / Unknown | ✔ promoted (branch candidate) | [EXP-0314](../experiments/EXP-0314-token-group-residuals/) |

### SAV-678

Helpers and reset:

- Complete helpers `L08643`/`R0676`/`L08644` respectively return the raw
  intersection with Player+0x2c, return 1 exactly when that intersection is
  zero, and OR the recipient's bits into Token+0x18.
- Reset `R1624` zeroes three fields, not one: word +0x18, dword +0x1c and
  dword +0x4, on a construction whose vtable's second slot is SAV-653's
  Token::Serialize.
- A second constructor of the same class, `R1134` (same vtable literal),
  copies +0x18 with the token's other fields rather than resetting it.

Clearing:

- Five instruction sites in four functions clear one recipient's bit with the
  complementary idiom: `L08645`/`L07383` in `R0064`; `L08646` in
  `R0655`; `L08647` in `R1625`; `L08648` in `R1626`.
- `R0655` clears, then calls the existing sender on the same recipient.
  The field is revoked and republished state, not a monotone accumulator, and a
  zero bit does not mean a recipient was never published.
- Two sites zero the whole field: `L08649` in `R1378`, paired with
  clearing the same +0x4c bit the sender tests, and `L08650` in
  `R0132`.

Senders, three functions plus one that inlines the same step:

- `R0059`: negative test, then setter (a dedup guard).
- `R0669`: the intersection test.
- `R0675`: a third caller of the negative test, branching on a virtual
  dispatch to either `R0059` or `R0562`.
- `R0562` implements the OR-publish step inline twice and ORs in 0xffff,
  every bit, when its recipient argument is null: the discriminator against a
  class flag or a Boolean.

The three helpers have exactly four call sites image-wide (`L08651`,
`L08652`, `L08653`, `L08654`: all CALL, zero JMP, zero absolute-dword
references), a complete census. The claim reuses the SAV-664 participant-mask
producer and SAV-653 wire transport. It does not establish a universal
fog-of-war meaning, a class discriminator, every clearing or inline site
image-wide, or the cause of every corpus exception. Closure evidence: the
experiment's `RESULT.md` and `evidence/closure/`.

**Confidence.** High for the three helper laws, the three-field reset, the
clear-then-republish mechanism at the four named clear sites, and the complete
four-site call census of the three helpers. Medium for the breadth of the
sender/clearer population: the third sender, the inline publisher and the
clearing sites were found by reading the functions the helper census and
existing references reached, not by an image-wide instruction-pattern sweep for
every `and`/`or` on a 16-bit [x+0x18] operand. The closest such sweep
(evidence/disp16-0x18.tsv, restricted to that exact 0x66-prefixed form) returns
54 anchored hits, 28 non-stack, each checked against Token's own object
identity.

**Unknown.** A universal fog-of-war meaning, a class discriminator, and the
cause of every corpus exception.

## Item Effects-list class membership and Token+0x0c

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-774 | `CArchive::ReadObject` (`L05928`) class-checks every typed reference, first occurrence or repeat, and admits any class derived from the requested one; for an item's Effects list that is `{Effect, Effect_DirectDamage}`. | High | ● active | [EXP-0320](../experiments/EXP-0320-item-effect-classes/) |
| SAV-775 | No located writer, and no member in a 98-file/5,134-item/234-member corpus census, puts a class other than `Effect` in an item's Effects list; `SAV-774` leaves `Effect_DirectDamage` as the only other class. | Medium | ● active (partially retracted) | [EXP-0320](../experiments/EXP-0320-item-effect-classes/) |
| SAV-776 | Both located item-Effect creation paths use the default constructor, and their own writer bodies contain no nonzero Token+0c stamp. | High / Medium / Unknown | ● active (amended, partially retracted) | [EXP-0320](../experiments/EXP-0320-item-effect-classes/), [EXP-0336](../experiments/EXP-0336-effect-attribution/) |

### SAV-774

Entry: the object is `this` and the first stack argument is `pClassRefRequested`. `L08655`, a push of that argument,
hands the requested class to `L08656` before the branch at `L08657`, so the
check has run before either of `ReadObject`'s arms is entered.

`L08656` is `ReadClass(pClassRefRequested, pSchema, pObTag)` (a return cleaning 0xc bytes;
`*pSchema` written at `L08658`, `*pObTag` at `L08659`;
`L08660`..`L08661` rejects a requested class whose `+0x8` schema word is
`0xffff`). The stream's 16-bit tag selects one of three paths:

- No class-tag bit (`L08662` testing the tag word, the frame local at -0x18, against the constant 0x80000000):
  `*pObTag` is written and null returned (`L08663`..`L08664`), the
  object-back-reference case `ReadObject` handles itself.
- `wTag == 0xffff` (`L08665` comparing the 16-bit tag at frame local -0xe with the constant 0xffff): first occurrence. The class
  name is read from the stream by the call to `L08666` at `L08667`, the schema is
  compared, then `L08668 JMP L08669`.
- Otherwise a class index resolves an already-seen class from the archive's
  load array (`L08670` loads the class pointer at 4 bytes times the class index in the load array) and falls into `L08669`.

Both class paths reach `L08669`: a comparison of the requested class with zero and `L08671` branching to `L08658` when zero,
the only branch that skips the check, taken only when the caller requested no
class; then `L08672` pushes the requested class, `L08673` takes the class the
stream named as the object, `L08674` calls `R0216`, and `L08675` pushes 0x6 and
calls `L08676` when that returns zero.

`R0216` is a derivation walk, not an identity test: it walks the class
descriptor up the `+0x10` base chain (`L08677` loads the base pointer at offset 0x10, `L08678` jumps back to
the entry), returning 1 on any ancestor match and 0 when the chain reaches null.
`ReadObject`'s object-back-reference arm applies the same walk to the
constructed object (`L08679 CALL R0214` = `GetRuntimeClass` through vtable
slot 0, then `R0216`) and reports the same error code 6
(`L08680`..`L08681`). The only unchecked returns in either function are a
null requested class and a null load-array entry (`L08682 JZ L08683`).

`Item::Serialize`'s read of its Effects list (`L08684`/`L08685 CALL R0951`)
reaches `L05928` through wrapper `L08128`
(`L08686` pushes L08129 and `L08687` calls `L05928`), which always requests a
non-null class, so no reference in that list can take the null-request bypass.

A byte-granular dword scan of every image section returns exactly one
non-`.text` dword holding `L08129`, at `L08688`, `Effect_DirectDamage`'s
`+0x10` base slot, and none holding `L08689`. The statically initialised
derived-class closure of `Effect` (descriptor `L08129`, name `L08690`, size
`0x48`, base `L08311` `Token`) is therefore exactly `Effect` and
`Effect_DirectDamage` (descriptor `L08689`, name `L08691`, size `0x60`),
consistent with `SAV-TOKENLOAD-093`'s base-chain enumeration of the same
descriptor set. Every other class name in that slot fails `R0216` and raises
archive error 6.

**Confidence.** High. Both functions are read from entry to terminating `RET`
(`L05928`..`L08692`, `L08656`..`L08693`), so their control flow is
enumerated, not sampled. The first excluded alternative, which this claim and
`formats/sav/format.md` once stated, is a reader that validates a typed
reference only against an already-resolved back-reference and constructs a
first occurrence by name alone. It is excluded by the two transfers that
converge on `L08669` before either class path can return (`L08668 JMP` from
the first-occurrence path, the `L08694` fall-through from the class-index
path), with `L08671 JZ` the sole bypass, gated on a null requested class that
`L08128` never passes. The second alternative, that the check tests class
identity, is excluded by `R0216`'s base-chain loop, which accepts any
ancestor. Scope: `L05928`/`L08656`, and descriptors as statically
initialised in the image. A reader that does not go through `L05928`, or a
descriptor built or patched at run time, is outside it and not claimed against.
`SAV-775` bounds what the corpus and the located writers produce at this one
list.

### SAV-775

Two located writers reach `Item::Serialize`'s list at `Item+0x20`:

- `ITEM-EFFGRAM-070`'s `Effects=` grammar (`R0999`, sole direct caller
  `R0992`) appends each `R1000` result. `R0999`'s only
  `+0xc`-adjacent instruction is a stack-local parameter read
  (a read at `L08695` of the routine's own second
  stack argument, at frame offset +0xc), not an object write.
- `R1031` is `SHOP-EFFALT-071`'s and `SHOP-MAGIC-007`'s magic-shop
  effect-generator dispatch. Its body re-derives `SHOP-MAGIC-007`'s
  "`2*ceiling-currentStoredPrice`" budget formula (`L08696` doubles the ceiling, then
  `L08697` subtracts the field at offset 0x1c of the item from it, the item's `+0x1c`). It is not a
  use-time writer: neither EXP-0320 nor `EXP-0226` places any of its callers at
  item use, only at the shop-stock route `SHOP-MAGIC-007`/`SHOP-EFFALT-071`
  name.

Field-level mechanism of the shop writer:

- `R1034` has exactly four direct callers image-wide, a complete census,
  all inside `R1031`.
- On one path `R1031` first calls `R1627` (its sole reference to it,
  and `R1627`'s only direct caller image-wide). `R1627` allocates 0x48
  bytes, constructs through the default `Effect` constructor (`R1049`),
  hardcodes kind `0x29`/41 and sets the operand from a table
  (`L06178`/`L08698`).
- Kind 41 is `castSpell`: the 50-entry kind-name table `MAGIC-EFFECT-015` cites
  at `L05192` stores one pointer per kind, index 41's slot at
  `L05192+41*4=L08699` holds `L08700`, and the bytes there read
  `castspell\0`. This is the node `SHOP-EFFCAST-065` says this generator
  appends.
- `SHOP-EFFCAST-065` does not say which field receives the node or what happens
  on a repeat kind. The target is the `Item+0x20` field `Item::Serialize`
  persists (`ITEM-SAVE-014`). `R1034` walks that list comparing each
  element's kind (`+0x3c`) with the new object's: a match updates that
  element's operand in place, no match appends the new object
  (`L05022`/`L08701`).

Both writers' construction sites were swept end to end for `[reg+0xc]` writes:
`R1000` over R1000-L08702, zero writes beyond the default constructor's
zero; `R1627` over R1627-L08703, one `+0xc` access, a stack-local read;
`R1034`+`R1031`, zero `+0xc` accesses of any kind.

Corpus census (`evidence/member.tsv`, `evidence/summary.tsv`):

- 16 roots: `gameversions/en`, `gameversions/ru` and 14 preserved save
  directories, the population `EXP-0313` used.
- 98 file instances (one `Bsg&`-magic file excluded, a documented non-original
  artifact), of which 95 complete a full document walk. The remaining 3 are one
  published digest, `SAV-DOC-053`/`SAV-EFFECTGRAPH-366`'s sole failing stream,
  which desynchronises only inside a later `SpellTransport` serializer after all
  `Player`/`Item` processing, so their item data is still captured.
- 5,134 items and 234 Effects-list member slots: every member's class is
  `Effect`, and all 234 carry Token `+0x0c`=0 (`SAV-776`).
- 57 of the 234 carry kind 41 (`0x29`), matching `R1627`'s hardcoded
  kind: the shop generator's merge routine fires in the preserved corpus, not
  only in static reachability.

Untraced callers: the shared list-append primitive is also reached through
intermediate thunk `L05022` (11 direct callers). 3 were disassembled and
build unrelated CString-based text, not an Effects list; the rest, including a
chain 13 callers deep through `R0451`, were not individually checked. A
third writer elsewhere in the image is not excluded.

Reader side: `SAV-CLASSSER-177`'s witnessed `SpellTransport` → `PointEffect` →
`Effect_DirectDamage` graph in `game0018.sav` places an `Effect_DirectDamage` in
a slot whose reader requests `Effect` (`SAV-CLASSSER-174`, the same `L08129`
descriptor `L08128` pushes for Item's list). That reference is accepted
because it passes the derivation check at `L08674`: `Effect_DirectDamage`'s
descriptor at `L08689` carries base pointer `L08129` at `L08688`. The
graph witnesses polymorphism through a checked subclass, not an unchecked
substitution. Any unlocated writer into an item's list is bound by the same
check, so the widest class set it could produce is
`{Effect, Effect_DirectDamage}`. `Effect_DirectDamage` is the class
`MAGIC-DMG-005`'s live damage-spell path constructs. Whether that path reaches
either located item-list writer is not established (`SAV-910`).

**Confidence.** Medium. The population is exact and the two located writers'
construction sites are read end to end. The negative, that no writer anywhere in
the image places a non-`Effect` class into this one list, rests on a caller
trace not carried to exhaustion (the 11-caller thunk `L05022`, the 13-deep
chain through `R0451`), and corpus agreement over 234 members does not by
itself exclude an unobserved writer. `SAV-774`'s reader-side check narrows such
a writer's output to one alternative class, `Effect_DirectDamage`; it does not
exclude a writer producing that one, so Medium and not High.

**Amended.** The clause that `SAV-776` places the live damage-spell path
"outside both located item-list writers' call graphs" is withdrawn
([`retracted.md`](retracted.md)). `SAV-910` withdraws the function identity at
`R1483` and the `L07735`→`R1067`→`R1483` caller chain that carried
that separation in `SAV-776`: `R1483` lies inside the six-byte instruction at
`L07736`, and no corrected whole-builder graph replaces the chain. The two
located writers, the 234-member corpus census, the `{Effect,
Effect_DirectDamage}` bound and the Medium grade do not rest on the withdrawn
clause and stand.

### SAV-776

- The grammar builder `R1000` allocates through `R1049`. Its complete
  `R1000..L08704` body writes kind, mode and operand but never Token+0c.
- The located magic-shop builder `R1627` uses the same constructor and has no
  own +0c store.
- The complete writer scans also excluded hidden string-move copies.
- The recorded 98-file/5,134-item/234-member corpus has Token+0c=0 on all 234
  item-owned Effects.
- Live spell-arm id stamps are a separate observed population.

**Confidence.** High for the two located writers' own stores and the
default-constructor relation, from the retained complete scans. Medium for the
bounded corpus.

**Unknown.** Writers outside the measured population, and complete ordinary
reach. No global absence claim is made.

**Amended.** EXP-0336 narrows the claim (SAV-910, [`retracted.md`](retracted.md)).
The former claimed function at `R1483` and the asserted
`L07735`→`R1067`→`R1483` caller chain are withdrawn: `R1483` lies
inside the six-byte instruction at `L07736`. The historical
whole-spell-builder caller separation is withdrawn with them, and the bounded
window supplies no corrected whole-builder graph. The two item writers'
own-store observations stand.

## Trailer block transfer, contents and the unit residual fields

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-790 | The 400-byte trailer block moves through one non-virtual call site for both directions; `R0272` reads and `R0271` writes. | High | ✔ promoted | [EXP-0321](../experiments/EXP-0321-sav-trailer-unit-fields/) |
| SAV-791 | Only trailer dwords 0 and 1 have a located consumer; dwords 2 through 99 have none within a complete search of the `[base+0x118]` displacement form. | Medium | ✔ promoted (amended) | [EXP-0321](../experiments/EXP-0321-sav-trailer-unit-fields/) |
| SAV-792 | `Unit+0x8e` is an incrementally maintained accumulator with a second writer, `R0451`, far better connected than the published derive, and `+0x90` is recomputed at each increment. | High | ✔ promoted | [EXP-0321](../experiments/EXP-0321-sav-trailer-unit-fields/) |
| SAV-793 | The carried-load arithmetic is 16-bit and its 64 000 threshold compare is signed; the consumer's encodings imply that a load large enough to wrap reads negative and cancels the speed penalty. | High / Medium | ✔ promoted | [EXP-0321](../experiments/EXP-0321-sav-trailer-unit-fields/) |
| SAV-794 | The original's load path reads `Unit+0x90` back and keeps it instead of recomputing it from the record's container, even when no recompute could have produced it. | High / Medium | ✔ promoted | [EXP-0321](../experiments/EXP-0321-sav-trailer-unit-fields/) |
| SAV-795 | In 31 streams no project-written input reaches, `Unit+0xa4` is wider than the published enumeration on both classes, and every `Unit` value is a whole number of cells. | Medium | ✔ promoted | [EXP-0321](../experiments/EXP-0321-sav-trailer-unit-fields/) |

### SAV-790

- `-xcallers R1491` returns exactly one direct E8/E9 site image-wide,
  `L08705` in `R1541`. The store arm reaches that same instruction by
  tail-jumping into the load arm's epilogue at
  `L08706`, a 5-byte relative jump to `L07783`, so `SAV-TRAIL-026`'s two arms share
  one transfer.
- An absolute-dword sweep over every section returns one hit, `L08707`, which
  its context bytes disqualify: the ten bytes around it encode
  a push of L08708 followed by a relative call, and the four "reference" bytes
  straddle the push immediate and the call opcode. No stored pointer to
  `R1491` exists, so no vtable holds it.
- `R1613` has the same shape: one caller `L08709` and zero dword
  references.
- Both archive primitives compute the fields at offsets 0x24 and 0x28 of the archive, their difference
  and a min against the requested count, then call memcpy `R0544`. The
  argument order separates them: at `L08710` the last value pushed is
  the first stack argument, the caller's buffer, which makes it memcpy's destination and the
  archive's buffer the source.
- `R1491` complements `ar+0x14` and tests bit 0, so MFC's set loading bit
  reaches `R0272`. Eighteen executed vectors over six mode values agree and
  give the transfer count as `0x190` at every one.
- The block is transferred unconditionally on load:
  `L08053` comparing the frame local at -0x10 with 0xbadface1 and `L08502` branching to `L07783` when not equal skips only
  `SAV-648`'s re-read of `[L07886]`, and a stream whose marker does not match
  still has 400 bytes consumed at that position.
- Evidence files: `evidence/q1-serializer-census.txt`,
  `evidence/q1-block-sites.txt`, `evidence/closure/vectors.tsv`.

**Confidence.** High. The two censuses are complete over their instrument's
form (direct E8/E9 in `.text`; absolute dwords in every section) and exclude the
two live alternatives by construction: a direction-specific second transfer
site would appear as a second E8/E9 hit, and a virtual dispatch would need a
stored pointer, whose one candidate lies byte for byte inside a `CALL rel32`
displacement. A call target computed into a register through arithmetic that
never forms the literal would not be found; none is asserted. The unconditional
load transfer is read from the branch target, not executed against a
marker-less stream.

### SAV-791

- The complete `[base+0x118]` displacement census over `.text` is 70 sites.
- Crossed with the complete absolute-dword reference census for the world
  singleton `L00285` (193 hits in 93 functions), 8 sites in 6 functions both
  use a non-stack base and sit in a function able to obtain the singleton by
  absolute addressing.
- Disassembly excludes two of the 8: `L08133` and `L08711` are
  `[this+0x118]` on a screen object whose `this` is the receiver, in functions
  that reference `L00285` for an unrelated purpose.
- The six that remain are exactly the trace-gate and toggle sites `SAV-647`
  names, and no seventh: `L08527` (index 0, in `R0076`), `L08712`
  (index 1, in `R0428`), `L00394` (index 1) and `L00395` (index 0, both
  in `R0262`), and `L08487` (index 1) and `L08486` (index 0, both in
  `R0440`).
- `SAV-647` names `L08486` as the `+0x0` toggle's read. The instruction
  before it, `L00397`, a load of `[L00285]`, is the singleton load and carries no
  `+0x118` displacement, so it is not in this census.
- The five further sites naming this block, in three functions, are lifecycle,
  not consumption: `L08036` (in `R0967`) stores the pointer,
  `L08037`/`L08038`/`L08039` (in `R1540`) free and null it, and
  `L08040` (in `R1541`) loads it for `SAV-790`'s transfer.
- This answers `SAV-TRAIL-026`'s "400 zero bytes is a shape, not a meaning" and
  `SAV-647`'s "the other 98 dwords are not characterized by this claim" as far
  as a static sweep reaches.
- Evidence files: `evidence/q1-disp118-census.txt`,
  `evidence/q1-blockuse-summary.tsv`, `evidence/q1-block-sites.txt`.

**Confidence.** Medium. The census is complete for the `[base+0x118]`
displacement form and the crossing is mechanical (`probe/blockuse.py` over the
two committed captures), but a consumer handed the block pointer as an argument,
or reaching it through an alias, uses no such displacement and is invisible. The
crossing criterion is also blind to a function that obtains the world as
`this`, which is how the constructor, the teardown and the document programme
reach it; those three are named from `SAV-646`'s reading, not found by the
crossing. No claim is made that the image contains no consumer, only that this
search found none.

**Amended.** The lifecycle-site count is corrected
([`retracted.md`](retracted.md)). The card said "The four further sites naming
this block are lifecycle" and listed five addresses. `evidence/q1-block-sites.txt`
and `evidence/q1-blockuse-summary.tsv` carry five `+0x118` sites in three
functions: the store in `R0967`, the test, load and null store in `R1540`,
and the transfer load in `R1541`. The six consumer sites, the dword 0/1
result and the Medium grade are unchanged.

### SAV-792

- The complete 0x66-prefixed access census is 5 anchored non-stack sites for
  `+0x8e` and 11 for `+0x90`, over three functions plus `Unit::Serialize`'s arm.
- `R0160`, `UNIT-CTOR-004`'s constructor, writes both to 0
  (`L08713`, `L08714`).
- `ITEM-LOAD-005`'s derive `R0280` has zero direct E8 callers image-wide
  and exactly two `.rdata` dword references, so it is reached only through a
  vtable slot: the `+0x50` slot `HERO-SIGHT-007` observes one indirect call to.
- `R0451` has 13 direct call sites in 8 containing functions
  (`R0430`, `R1536` ×3, `R0448`, `R1572`, `R1628`, `L08715` ×2,
  `L08716` ×2, `L08717` ×2). It takes a 16-bit delta and adds it to `+0x8e`
  at `L08718`/`L08719`/`L08720`. It then recomputes `+0x90` from `+0x8e`
  and `[actor+0x7c]+0x20` with the derive's three-branch arithmetic,
  instruction for instruction, bracketed before and after by the same
  `MOVSX`/`IDIV` load-over-capacity quotient.
- A third `+0x90` write inside the derive, `L08721`, sits outside the nine
  instructions `ITEM-LOAD-005` enumerates.
- `ITEM-LOAD-005` and `HERO-SIGHT-007` remain correct about the derive. Neither
  states that the derive is one of two implementations of the same arithmetic,
  or that the implementation with every direct call site is the other one.
- Evidence file: `evidence/q2-field-writers.txt`.

**Confidence.** High for the two censuses and the caller counts (`-disp16` is
decode-anchored and complete for the 0x66-prefixed form, and `-xcallers` is
complete for direct E8/E9) and for `R0451`'s body, disassembled whole.
The excluded alternative, that the derive is the only maintainer of either
field, needs `R0451`'s three writes not to exist; they are named
instructions with 13 call sites. A write of 32-bit or of byte width to either offset
is outside what `-disp16` reports, so "no other writer" holds only for the
16-bit form.

**Unknown.** What each of the 13 callers is doing.

### SAV-793

- `L08722` compares the field at offset 0x20 with 0xfa00 and is followed by a signed greater-or-equal branch, not an unsigned above-or-equal branch: the
  container's weight sum is compared as a signed dword.
- `L08723`, adding two 16-bit values, is a 16-bit add and wraps at 0x10000; only the low word
  of the 32-bit halved sum participates.
- The consumer reads the field back with `MOVSX` at `L08724` and branches with
  `JL` at `L08725`, so a stored value at or above 0x8000 is negative and the
  `load >= capacity` test fails.
- The same three properties appear again at `L08726`..`L08727` inside
  `SAV-792`'s `R0451`.
- Executed vectors separate this from the readings the published prose
  supports:
  - `own = 0xffff, sum = 224, container present` stores 111, where a 32-bit
    accumulate stores 65 647.
  - `own = 30, sum = -2` stores 29, where an unsigned threshold stores the flat
    32 000.
  - `own = 30, sum = 63999` stores 32 029 and `sum = 64000` stores 32 000,
    fixing the boundary as inclusive.
- `ITEM-LOAD-005`'s and `HERO-SIGHT-007`'s arithmetic is unchanged; this adds
  its integer width and signedness.
- No preserved record reaches the wrap: the largest `+0x90` in the corpus is
  457.
- Evidence files: `evidence/closure/vectors.tsv`,
  `evidence/q2-derive-and-load-arm.txt`.

**Confidence.** High for the widths, the signedness and the boundary: 110
executed vectors over own weights `{0, 1, 30, 0x7fff, 0xffff}` and sums
`{-2, 0, 1, 2, 3, 224, 63998, 63999, 64000, 64001, 0x7fffffff}`, with and
without a container, each predicting a different stored value under a 32-bit
accumulate or an unsigned compare. Medium for the behavioural consequence:
`MOVSX` and `JL` are read from their encodings and the wrap is executed, but no
vector runs the penalty branch and no preserved save reaches the wrapping
range, so the penalty cancellation rests on the two encodings, not on an
observation of the game.

### SAV-794

- `Unit::Serialize`'s load arm reads `+0x8e` at `L08728`/`L08729`, `+0x90`
  at `L08730`/`L08731` and `+0xa4` at `L00925`/`L00926` through the same
  archive helper, into the object.
- In `EXP-0261-owner-runs`, whose preserved manifest records which streams the
  original loaded and resaved, one `Human` carries `+0x8e = 178`,
  `+0x90 = 181` and an empty container in five streams: `game9003.sav` and
  `game9004.sav`, the two project-written inputs; `game0002.sav` and
  `game0004.sav`, the original's own city resaves of those two; and
  `game0005.sav`, the original's resave after the owner shopped and hired inside
  the same session.
- Under `SAV-793`'s arithmetic an empty container forces `+0x90 == +0x8e` at
  every recompute, so 181 is a value no recompute could produce. The original
  read it, played, and wrote it back unchanged, twice.
- Corpus-wide the relation holds on 1 352 of 1 420 records (1 345 of 1 411
  original-written), and `+0x90 >= +0x8e` on all 1 420. Of the 55 records whose
  container has contents, 38 agree; 51 records carry `+0x90 > +0x8e` with an
  empty container.
- A consumer must round-trip what the file supplies.
- Evidence files: `evidence/q3-summary.tsv`, `evidence/q3-unit-scalars.tsv`,
  `evidence/q2-derive-and-load-arm.txt`.

**Confidence.** High that the load does not recompute: the project-written
streams are the stimulus and the evidence is the original's own output. The
alternative, a recompute anywhere on the load path, predicts 178 in every
original resave and is refuted by two independent ones. Medium for what makes
the stored values disagree. Two writers of the field are enumerated for the
16-bit form (`SAV-792`) and neither has a third term, so an empty container
cannot produce 181. But a write of 32-bit or of byte width to the same offset is
outside `-disp16`'s reach, and a derive behind a trigger that a city session
does not fire predicts the same round trip as no derive at all. Neither reading
is excluded.

### SAV-795

- Population: the 31 preserved streams no project-written input reaches
  (`SAV-796`'s partition), 826 `Unit` and 448 `Human` records.
- On `Unit` the field has a zero low byte on all 826 and takes six values, not
  four: `1024:27 1280:39 1536:609 1792:16 2048:132 2304:3`. `SAV-UNITFLD-049`'s
  176-record population saw `1024 × {1, 1.25, 1.5, 2}`.
- Read in the unit `AI-SIGHT-092` and `TERR-FOG-081` publish, a `u16` in 1/256
  cell whose high byte `+0xa5` is the `Data.bin` Units column `scanRange`, the
  six values are sight radii of 4, 5, 6, 7, 8 and 9 whole cells, inside
  `TERR-FOG-081`'s shipped `scanRange` span of 4..12. The two values
  `SAV-UNITFLD-049` did not see are 7 cells (16 records) and 9 cells (3).
- The placed-actor mode is 6 cells (609 of 826), not the shipped Units table's
  mode of 7: a difference between two populations, not between two readings.
- On `Human` the field spans 1331..2037 in fourteen distinct values, with 19
  records outside `SAV-UNITFLD-049`'s published 1382..1658.
  `HERO-SIGHT-007`'s `ftol(((mind + reaction)/25 + 4) × 256)` reproduces the
  stored value on 442 of the 448.
- All six exceptions are three records duplicated across two streams
  (`en/game0020.sav` and `2026-08-15/game9999.sav` at offsets 6427, 8141 and
  9845), each storing 2037, the value for a stat sum of 99, against a stored
  `mind = 25`, `reaction = 49`, sum 74.
- Capacity is 300 on all 826 `Unit` records and `body × 10 + 1` on all 448
  `Human` records, extending `SAV-UNITFLD-049`'s 176/100 populations with the
  same result.
- Evidence files: `evidence/q3-summary.tsv`, `evidence/q3-unit-scalars.tsv`,
  `evidence/q3-authorship.tsv`.

**Confidence.** Medium: a corpus census over a stated population, not a
controlled trace. The two additional radii extend an enumeration taken from a
smaller population and contradict no instruction. The whole-cell reading is
`AI-SIGHT-092`'s and `TERR-FOG-081`'s at their own grades, not measured again.
The census adds that the low byte is zero on all 826 `Unit` records, which the
whole-cell `+0xa5` streamer predicts and the 1/256-cell `Human` derive does not;
the same census puts 442 of 448 `Human` records on values with a non-zero low
byte.

**Unknown.** Why six records store a sight value for a stat sum they do not
carry. A stale derive from before a stat change and a derive reading effective
rather than stored stats both fit these bytes, and nothing here separates them.

## Unit residual fields: the sight word and Token+0x18 in the corpus

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-796 | `SAV-UNITFLD-049`'s "sight `+0xa4 = 0` occurs in no record" holds on 1 274 original-authored records against the 276 it was measured on; its three zero records trace to project-written input, and the constructor leaves 1280. | Medium | ✔ promoted | [EXP-0321](../experiments/EXP-0321-sav-trailer-unit-fields/) |
| SAV-797 | Token+0x18 is 2 on 1 404 of 1 420 records and 0 on the other 16, no other value occurs, and the terminal stage does not decide which. | Medium | ✔ promoted | [EXP-0321](../experiments/EXP-0321-sav-trailer-unit-fields/) |

### SAV-796

- Population: the 46 distinct preserved streams (45 parsed), split three ways by
  their own directories' `MANIFEST.md`:
  - 31 original-authored: the original engine wrote them and no project-written
    input appears anywhere in their chain; 1 274 decoded `Unit`/`Human` records;
  - 8 original-written resaves of project-written candidates, all in
    `EXP-0261-owner-runs`, whose origin column names each generated candidate
    and each `ROM1 resave`; 137 records;
  - 7 project-written, one of which fails the header magic; 9 records.
- Over the 1 274 original-authored records `+0xa4 = 0` occurs 0 times.
- All three records carrying it in the whole 1 420-record census sit in
  `EXP-0261-owner-runs` and agree in every one of 24 scalar columns:
  `game9001.sav@249` and `game9006.sav@248`, generated candidates by the
  manifest's origin column, and `game0000.sav@249`, which the same column
  records as ROM1's resave of `game9001.sav`. That pair is the one
  `SAV-ORIGVALUE-399` publishes as a 1,527-byte projected candidate and the
  original's distinct 1,540-byte resave of it, carrying invented `Human` values.
  The original preserved that field; nothing here shows it chose it.
- The forbidden state is not the constructor's. `AI-SIGHT-092` reads
  `L00260`, a byte store of 0x5 at offset 0xa5 of the object, and
  `L00918`, a byte store of 0x0 at offset 0xa4 of the object, as the little-endian halves of the
  `u16` `0x0500`, and `R0160` is straight-line with no branch between
  them, so the constructor leaves 1280 at `+0xa4`. That value is present on 39
  of the 826 original-authored `Unit` records.
- `SAV-UNITFLD-049`'s sentence carries `UNIT-CTOR-004`'s byte-width phrasing of
  a word-width field, so the value it looks for is not the one the constructor
  writes.
- Of that constructor's fifteen `+0x84..+0xa5` stat immediates the disputed
  record carries five, `+0x84` 30, `+0x8c` 10, `+0x8e` 0, `+0x90` 0 and `+0x92`
  300, and contradicts ten. The ten include both immediates `SAV-UNITFLD-049`
  uses to fix the record alignment, 100 at `+0x98` and 50 at `+0x9e`. Those
  three records are the only three of the census's 516 `Human` records that fail
  that anchor; 907 of all 1 420 records fail it because `Unit` records do not
  carry the constructor's regeneration values.
- Dropping the 8 resave streams costs the `+0xa4` census exactly one distinct
  value, 0, and no other.
- Evidence files: `evidence/q3-authorship.tsv`, `evidence/q3-summary.tsv`,
  `evidence/q3-unit-scalars.tsv`, `evidence/q2-field-writers.txt`.

**Confidence.** Medium. A census over a population named by preserved manifests,
plus those manifests' own record of which streams the original loaded. It
excludes one reading of the zero, that the original produced it from its own
constructor: the constructor's word is 1280 and the corpus carries 1280. Both
readings of the three records named under Unknown leave the statement about the
original standing. The resave supports one further point, bounded to one chain:
the original read a `Human` whose stored sight word was 0 and wrote 0 back,
where `HERO-SIGHT-007`'s derive on that record's own `mind = 10`,
`reaction = 20` gives 1331, a value 2 original-authored records carry, so a
recompute at load would have been visible. That is `SAV-794`'s round trip at
`+0xa4`, not a second result.

**Unknown.** An original-authored stream outside these 46 carrying a zero.
Whether the three records hold a zero sight word or a scalar run the walker
places wrongly; the anchor that would settle it is the one they fail.

### SAV-797

- Over 45 preserved streams the cross-tabulation of stage against mask is
  `stage0:mask0=3 stage0:mask2=1282 stage1:mask2=16 stage3:mask2=9 stage4:mask2=84 stage5:mask0=13 stage5:mask2=13`.
- Of 26 records at stage 5, exactly half carry 0, and the same dead-unit
  population appears both ways: `en/game0002.sav` holds three stage-5 records
  with health −10014, −10017 and −10011 at mask 2, and `en/game0003.sav` holds
  the same three healths at mask 0. A zero mask is therefore not "removed from
  play".
- The three stage-0 records with mask 0 are all in
  `EXP-0261-owner-runs/game0005.sav`, which that directory's manifest records as
  the resave made after the owner hired three nameless `Human` actors. This is
  consistent with `SAV-678`'s publication reading, in which an unpublished
  recipient bit is ordinary state rather than an exception.
- This narrows `SAV-678`'s "the cause of every corpus exception" Unknown by
  rejecting one rival reading; it does not close it.
- Evidence files: `evidence/q3-summary.tsv`, `evidence/q3-unit-scalars.tsv`.

**Confidence.** Medium. A corpus census over a stated population. The stage/mask
cross-tab excludes the terminal-stage reading outright, since the same three
health values appear at both masks in two streams of the same campaign. The hire
reading is corroboration from a preserved manifest's own description, not an
observation of the hire, and no instruction is traced here: `SAV-678` owns the
mechanism.

**Unknown.** Whether every zero mask has the same cause.

## Group order values and the Group+0x20 list

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-GRPORDER-806 | Each of `grpAI+0x20`'s eight live values has a named, cited static producer; the preserved corpus's 567 Group records carry five of them, and `2`, `5` and `0x11` never occur. | High / Medium | ● active | [EXP-0322](../experiments/EXP-0322-group-order-producers/) |
| SAV-GRPLIST-807 | None of the shared list class's 18 external direct-call sites targets `Group+0x20`, the site `SAV-GRPAI-563` leaves unclassified; other writers and list readers remain Unknown. | High / Medium | ● active | [EXP-0322](../experiments/EXP-0322-group-order-producers/) |

### SAV-GRPORDER-806

- The value space and its arms are `AI-ORDER-010`'s own six-entry table plus its
  `0x11`/`0xff` arms. Producers by value:
  - `0`: the constructor and twelve command handlers (`AI-ORDER-010`,
    `AI-CMD-033`);
  - `1`: the player Guard command and the map-load guard setter (`AI-CMD-033`,
    `AI-AUTHOR-015`);
  - `2`: the script's Swarm sub-command, 6 shipped nodes over 20 maps
    (`AI-GROUPCMD-020`), plus two structurally distinct producers with no direct
    caller or stored absolute reference found for either (`AI-ORDER-031`'s
    amended orphan site, `AI-ORDER-294`);
  - `3`: the player Aggressive command, the script's Stand Ground sub-command
    (24 shipped nodes) and the map-load setter (`AI-CMD-033`, `AI-GROUPCMD-020`,
    `AI-AUTHOR-015`);
  - `4`: the player Move command and the script's Move sub-command (29 shipped
    nodes) (`AI-CMD-033`, `AI-GROUPCMD-020`);
  - `5`: the player `0x1a` command and the script's Swarm 2 sub-command (40
    shipped nodes) (`AI-CMD-033`, `AI-GROUPCMD-020`);
  - `0x11`: the script's Roam sub-command, 0 shipped nodes over 38 maps
    (`AI-GROUPCMD-020`, `AI-ROAM-025`);
  - `0xff`: the dispatcher's own empty-group force (`AI-ORDER-010`,
    `SAV-GRPDISPATCH-568`).
- `tools/savdoc -mode order-census` (EXP-0322) reads the byte at `AIOff+0x20`
  for every Group record over the sixteen directories `EXP-0321` names,
  sha256-deduplicated: 46 distinct-content streams. 1 fails the header-magic
  check before any parse (`game9000.sav`, project-written). 1 more parses but
  hits a walk-error the walker reports for any stream reaching it, a
  `SpellTransport` decode gap `SAV-CLASSSER-175`/`SAV-EFFECTGRAPH-366`
  characterize (`game0018.sav`, original-authored, offset 46969 of its decoded
  body). That leaves 44 streams and 567 Group records.
- The observed value set is `{0:26, 1:414, 3:33, 4:41, 0xff:53}`; `2`, `5` and
  `0x11` occur zero times. By authorship: original-authored 511 records,
  `0:23 1:382 3:27 4:27 0xff:52`; original-resave-of-project-input 47,
  `0:2 1:32 3:6 4:6 0xff:1`; project-written 9, `0:1 4:8`.
- `0x11`'s absence matches its own 0-shipped-node count. `5`'s absence is
  consistent with, not independently shown by,
  `AI-SWARM2-024`/`AI-CANDCOUNT-105`'s finding that its arm's ordinary path
  tail-calls directly into `4`'s arm whenever the candidate build is empty.
- Evidence file: `evidence/q4-order-census.tsv`.

**Confidence.** High for the census itself: the value read is the same field
`SAV-GRPDISPATCH-568` dispatches on, the population is exact and named, and the
instrument is reproducible. Medium overall: a single-snapshot corpus, not a
controlled trace, cannot show whether `2`, `5` or `0x11` ever occur in play and
never survive to a save point. The walker discards a failing stream's
already-parsed rows rather than emitting them, so `game0018.sav`'s 16
already-parsed Group records are reported as withheld
(`withheld_parsed_groups=16`) and excluded rather than counted zero.

**Unknown.** Why `2` is absent despite 6 shipped script nodes and no published
fallback; nothing this experiment finds explains it. Computed-call reachability
of value `2`'s two orphan producers. Any Group records of `game0018.sav` beyond
the failure point.

### SAV-GRPLIST-807

- `SAV-EMBED-039` places `Group+0x20` among five sites of one class (ctor
  `R1545`, vtable `L08172`, `Serialize` `R1546`). The class exposes
  exactly two mutation entry points every known instance shares: the append
  wrapper `R1587` and the grow/insert primitive `R0052`
  (`SAV-632`).
- A whole-image direct-call/tail-jump census of both
  (`evidence/q3-append-wrapper-callers.txt`: 2;
  `evidence/q3-grow-insert-callers.txt`: 17, one of which is `R1587`'s own
  internal call into `R0052`) leaves 18 distinct external sites. Each is
  classified by its true containing function, verified against that function's
  preceding `RET`+padding boundary and its own prologue, not taken from the
  census's raw label.
- Five of the 18 are reported "in" a preceding function that does not enclose
  them, the containing-function mis-attribution EXP-0322's own `AI-ORDER-031`
  amendment found once, here recurring five times:
  - `L08732`'s real container is `L08733`, not `R1331`; both are
    complete, independently `RET`-terminated functions
    (`evidence/q3-nongroup-singleton-sites.txt`);
  - `L08734`'s is `L08735`, not `L08736` (same file);
  - `L08737`'s is `R0163`, `AI-PATROL-013`'s patrol loader, not `L08738`,
    a mid-function `JMP` target with no prologue;
  - `L08739`'s is `R1543`, one of `SAV-631`'s five named Unit-embedded-list
    call sites, not `R1542`, a separate earlier function ending at its own
    clean `RET`;
  - `L08740`'s is `R1546`, `Serialize`'s load arm, not the destructor
    `L08741`, whose body (vtable install, `SAV-632`'s teardown, then
    `operator delete`) ends at its own `RET`+`INT3` padding before it
    (`evidence/q3-list-serialize-and-dtor.txt`).
  - `L08742`'s raw label, `R0161`, needed no correction.
- The 18 sites by target:
  - 1 reaches the patrol-path loader, `AI-PATROL-013`'s `R0163`
    (`Group+0x4c`). Its corpus population is also zero across the same 567
    records (`w3c_count` in `evidence/q4-order-census.tsv`), consistent with
    `AI-PATROL-013`'s finding that no shipped mission file reaches that loader.
  - 1 is `Serialize`'s load arm at `R1546` (`SAV-WLIST-040`), shared
    identically by every instance.
  - 4 build or rebuild the per-actor active ring at `order+0x90`:
    `AI-PATROL-018`'s `R0172` and `SAV-GRPPATROL-570`'s `R0164` by
    name, plus `R0161`.
  - 5 belong to the two Unit-embedded route lists: an exact reproduction, by an
    independent instrument, of `SAV-631`'s five named call sites, with none left
    over.
  - 2 calls, one apiece inside two near-identical routines at field offsets
    `+0x54098`/`+0x5409c` of an object not otherwise identified. Neither address
    nor either offset is named by any claim in any ledger
    (`evidence/q3-nongroup-singleton-sites.txt`,
    `evidence/q3-claim-search-nongroup-singleton.txt`).
  - 5 calls, all inside one function, `R0142`, append to a list based at that
    function's own `+0xa4538` (`evidence/q3-nongroup-search-site.txt`,
    `R0142`-`L08743`, whose return cleans 0x8 bytes; the outer ring loop's bound test against
    `actor+0xa5 + 1` is at `L00272`-`L00273`, inside this capture).
- `R0161`: `AI-PATROL-019` (corroborated by `AI-GUARD-012`) names it a live
  patrol-state/destination setter with a single, itself-uncalled caller, and
  cites `L00430` as writing "a literal 0" into "the ring it builds", beside a
  separate, differently shaped 12-byte-record list the same function builds a
  few instructions earlier (`L08744`-`L08745`, its own distinct realloc
  helper, unrelated to this class). Neither row names which structure `L00430`
  belongs to. Read through the function's end
  (`evidence/q3-second-ring-builder.txt`, `R0161`-`L08746`, whose return cleans 0x4 bytes), the
  preceding instructions `L08747`-`L08742` fetch the field at offset 0x90 of the object at offset 0x158 and
  call this class's grow/insert primitive on it with that pointer as `this`.
  So `L00430` is inside `order+0x90` itself, reached through the shared list
  class. This attributes `R0161`, which `SAV-633` names verbatim among its
  ten unattributed `+0x90` sweep owners.
- `R0142` is named, though not read, by three claims across three ledgers:
  `AI-SIGHT-006` (a "located, not read" reader of `actor+0xa5`), `MOVE-AREA-038`
  (one of two owners calling `R0144`, "not part of the search") and
  `TERR-PASS-051` (one of 13 owners of the bit-5 static-plane hash-table test).
  This read corroborates each: the function loads `actor+0xa5` as a ring-search
  radius (the loop bound compared at `L00272`-`L00273`), validates each
  candidate cell against the static plane (`R0143`) and `R0144`, and
  appends every passing candidate through this class's primitive.
- The pointers `R0142` builds at `+0x5402c`/`+0x540b4` are not a second field
  of that list: `+0x540b4` keys a lookup into the cell-record hash table near
  `world+0x540b8`, and `+0x5402c` receives the found record's copied body
  (`L08748`-`L08749`, 13 dwords) (`TERR-PASS-051`, `SAV-CELLLOAD-111`,
  `TERR-CELLREC-146`). They were read only far enough to exclude them from this
  class's fields.
- None of the 18 targets `Group+0x20`. Two places that do touch the field were
  read: the Group constructor's in-place construction (`L08750` adding 0x20 to the object pointer, then
  a call to `R1545`, `evidence/q1-group-ctor.txt`) and `Group::Serialize`'s
  dispatch through the vtable slot of the object at offset 0x20 (`L08751`, adding 0x20 to the object pointer,
  `evidence/q3-group-serialize-dispatch.txt`), identical to how it reaches every
  other instance.
- These two are not asserted to be the only ones. Both reach the field by
  an addition of 0x20 to the object pointer, a form the experiment's `-disp 0x20` sweep does not report, so
  no instrument here enumerates the field's touch sites and nothing excludes a
  reader of the list.
- Corpus corroboration: 0 of the 567 Group records `SAV-GRPORDER-806`'s census
  reaches carry a nonzero `Group+0x20` element count.

**Confidence.** High for the exhaustive direct-call census of both named entry
points and for the corpus corroboration, each exact and reproducible. It
excludes an incomplete or mis-labelled census as the reason no site targets
`Group+0x20`: every one of the 18 raw labels was re-verified against its own
function boundary, which is what corrected 5 of the 18. Medium overall: a bulk
copy, an aliased pointer or a third mutation path that bypasses both entry
points is excluded by neither instrument, and a computed call reaching either
entry point is invisible to the direct `E8`/`E9` census, the class of blind spot
`SAV-634` names for its own module-scoped negative result.

**Unknown.** Other writers of `Group+0x20` and readers of its list. The object
that owns `+0x54098`/`+0x5409c`.

## Serialized turn state

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-TURNLOAD-822 | The local Mover serializer reads the complete 180-byte block, including the turn fields; the first post-LOAD driver is still unproved. | High / Unknown | ✔ promoted | [EXP-0324](../experiments/EXP-0324-turn-continuation/) |

### SAV-TURNLOAD-822

- Unit serializer `R0210` loads the pointer at `Unit+0x154` at `L08137` and
  calls `R1348` before its own store/load branch.
- Complete `R1348..L08752` passes the Mover address and size `0xb4` to
  archive-write `R0271` or archive-read `R0272` according to the archive
  mode, then returns. It has no field-specific normalization.
- The fields `MOVE-TURN-044` names, current/desired bytes `+0`/`+1`,
  RotationSpeed byte `+a`, counter byte `+9d`, active DWORD `+a0` and estimate
  byte `+a4`, are therefore within the raw read, not values this serializer
  derives from a route.
- The selected Unit load arm `L08753..R1482` contains no further direct
  Mover access. Its unresolved virtual call at `L08754` (actor vtable
  `+0x30`), embedded serializers and later lifecycle prevent extending this
  local finding to whole-LOAD preservation or exact first-tick scheduling.
- Evidence: the unit-prefix, mover-serialize and complete unit-load-arm
  listings.

**Confidence.** High for the local address, length, mode branches and absence of
direct normalization in the complete Mover serializer. Unknown for whole-LOAD
and first-driver preservation. Normal archive return is assumed; short-read and
exception effects are not measured.

## World LOAD registration and first mover boundary

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-946 | Group LOAD has distinct embedded-list, separate-state and actor archive receivers before membership insertion. | High / Unknown | ✔ promoted | [EXP-0341](../experiments/EXP-0341-human-load-first-use/) |
| SAV-947 | The selected late world LOAD callbacks include record repair and a manager self-pointer store; no-world LOAD bypasses this suffix. | High / Unknown | ✔ promoted | [EXP-0341](../experiments/EXP-0341-human-load-first-use/) |
| SAV-948 | Resume setup has a conditional opcode 4 route to actor projection before ordinary actor subticks. | High / Unknown | ✔ promoted | [EXP-0341](../experiments/EXP-0341-human-load-first-use/) |
| SAV-949 | The first computed Human use after LOAD remains a path-conditioned frontier over a bounded static selection. | Medium / Unknown | ✔ promoted | [EXP-0341](../experiments/EXP-0341-human-load-first-use/) |
| SAV-LOADREG-878 | World LOAD registers the actor value rebuilt from Player groups through an append helper, not the creator's ID-assignment helper. | High | ✔ promoted | [EXP-0330](../experiments/EXP-0330-first-mover-event/) |
| SAV-LOADHOOK-879 | The selected world post-load actor callback uses `+24`, separately from the derive slot `+50`. | High | ✔ promoted | [EXP-0330](../experiments/EXP-0330-first-mover-event/) |
| SAV-FIRSTMOVE-880 | The selected LOAD-to-tick chain does not establish an unconditional first mover event. | High / Unknown | ✔ promoted | [EXP-0330](../experiments/EXP-0330-first-mover-event/) |

### SAV-946

- Constructor `R0149` passes `Group+20` to `R1545`, whose `L08172` vptr
  makes `L08755` select `R1546` at slot `+8`.
- The next call, `L08756`, uses `Group+3c` and can perform separate
  cleanup/list callbacks.
- The actor loop's `L08346` writes `L05928`'s object through its second
  output argument and returns the archive pointer.
- The reader's new-object arm registers the created receiver at `L08079`
  before calling that receiver's `+8` at `L08080`; its existing-object arm
  returns an indexed object instead.
- The Group caller reloads the actor from that output before insertion
  `R0152`. The insertion can remove old `actor+70` membership before
  appending the same actor, setting `actor+70` to the current Group and copying
  `actor+14` to `Group+44`.
- Evidence: selected Group/archive bodies, constructor slots and instruction
  audit.

**Confidence.** High for named receivers, branches and instruction order. This
does not establish transitive no-derive behavior. Unknown for archive aliases,
reentrant serializers, factories/removal helpers, arbitrary pointer aliases and
the first computed Human use.

### SAV-947

- After the actor hook, `L08757` invokes the `world+0c` receiver's `+4`.
- `L08758` calls `R1597`, which copies thirteen dwords into a stack record,
  calls `R1598` on that record and copies it back. Ten record dwords at
  `+4..+28` skip zero and are replaced only when `R1488` succeeds.
- At `L07741`, `R1485` only stores its manager receiver at `manager+a9c0`.
- At `L08759`, constructor `L05938`'s `L08760+0` resolves to `R1126`,
  which calls each separate manager-list entry's `+24`.
- The no-world arm skips these callbacks, clears `actor+40`/`+44`/`+5c` and sets
  the server run gate `+2c` to zero.
- Record and manager offsets are not Human offsets.
- Evidence: world suffix, copied-record repair, manager constructor/table and
  no-world branch.

**Confidence.** High for the selected branch/call order, stack-record
provenance, local success-only stores and resolved manager slot. Unknown for
`world+0c` and separate entry classes/callback effects, lookup/allocation
aliases and absolute first Human computation.

### SAV-948

- Before optional bootstrap `L03780`, `L03843` calls `R0545`. It sends
  byte opcode 4, WORD ID from the `interface+9b4` object's `+4`, zero
  destination and the resume argument in packet DWORD `+0a` through client
  manager `L00625`.
- Conditional on delivery unchanged to the ordinary server dispatcher with
  packet `+4` zero, the opcode 4 switch selects `L06671`; `L00171` matches
  packet WORD `+5` against Player WORD `+4`.
- A miss skips setup; a hit passes that Player and `(packet+0a==0)` to
  `R0131`.
- The later `L03739` projection call receives that Player's then-current `+34`
  actor, the same Player and mask-1. For a Human meeting `SAV-HUMPROJECT-461`'s
  mask conditions this reaches its computed damage-byte projection.
- The ordinary queue prefix precedes world-object and actor subticks.
- The selected projection is not claimed to be the first read inside setup or
  after LOAD.
- Evidence: frontend helper, transport/queue cuts, opcode 4 switch and
  receiver-specific setup call; `SAV-HUMPROJECT-461`.

**Confidence.** High for local packet fields, the two-stage switch, the lookup
and call arguments. Unknown for transport delivery, earlier setup/notification
callbacks, continued identity of `Player+34` with a particular restored Human,
earlier actors/packets and absolute first-read order.

### SAV-949

- The expanded 38-body selection distinguishes Group/archive receivers, late
  manager repairs and the conditional resume projection path.
- Before the selected ordinary Human subtick it retains `world+0c`/entry
  callbacks, archive cleanup, frontend transport, phase 6 Group actor work,
  phase 12 projection/full ticks, queued packets, `world+2c` entry `+18`
  callbacks and earlier actors.
- World-present and no-world paths have different callback/run gates.
- The known conditional Effect-before-order relation does not order these
  earlier cuts.
- Neither a universal derive-before-read guarantee nor safe replacement fields
  follow.
- Evidence: explicit call frontier and remaining native discriminator. Also
  cites `SAV-HUMLOAD-445`, `SAV-HUMRESUME-460`, `SAV-HUMFIRST-465`,
  `SAV-LOADREG-878`, `SAV-LOADHOOK-879`, `SAV-FIRSTMOVE-880`.

**Confidence.** Medium for this bounded static frontier. Unknown for absolute
first computation, values at that read, native skipped ticks and
alias/reentrant/exception ordering.

**Unknown.** A same-Human ordered trace through archive completion, repairs and
first computation or explicit bypass is missing.

### SAV-LOADREG-878

- Player `R0415` calls the group-list serializer on `Player+24`. On LOAD it
  then iterates each group's actor list, passes each actor to `R0032` on
  `Player+20` at `L01794` and stores the Player in that `actor+14`.
- In the world-present arm of `R0414`, the loop at `L06795..L08761`
  traverses `Player+20` and passes the same iterator result to
  `R0032([L00240],actor)` at `L01793` when `actor+4c` mask `0x08` is
  clear.
- Complete `R0032` delegates to `R0292` on `manager+4`, then `R0412`
  appends a node containing that exact actor pointer.
- Iterator `R1334` advances the node cursor before returning `node+8`.
- Creator `R0411` also calls `L08762` and `R0936` and writes low16 of
  the latter return to `actor+4`. Those calls and that store are absent from the
  restored insertion body.
- Evidence: world/Player bodies, insertion/iterator bodies and raw reference
  inventory.

**Confidence.** High for the selected receiver/list operations and the direct
distinction, confirmed in both matching images. No claim of arbitrary alias
safety, whole-LOAD byte preservation or global absence of an ID/derive writer
follows.

**Unknown.** Node allocation/initialization helpers remain transitive
boundaries. Actor creation and group archive callbacks are not fully closed
here.

### SAV-LOADHOOK-879

- World-present `R0414` calls the global actor manager's `+4` at `L08763`
  after tick-list registration and later world reconstruction.
- Constructor `R0120` installs `L06791`, whose `+4` is `R1599`; that
  body passes each iterator actor as `this` to `actor+24` at `L08764`.
- Unit `+24` selects `R1482`; Humanoid/Human `+24` select `R1569` and
  delegate to `R1482`.
- The hook reaches the allocated mover repair `L08765` only when actor BYTE
  `+13c` is zero. The repair's only direct mover store is DWORD `+7c` after
  successful lookup; its body has no direct `+a` access.
- These complete post-load wrappers do not directly invoke `actor+50`; the
  Human/Humanoid producer is in the separate `R0280` slot.
- Evidence: selected class tables, complete post-load bodies and mover repair.

**Confidence.** High for the original manager/actor slots, constructor
assignments, exact receivers, stage branch and bounded direct-store statement.
This proves neither no derive during all LOAD nor that the transported byte
survives until its first consumer.

**Unknown.** Token/reference/order helpers, lookup callbacks, other manager
post-load calls and arbitrary aliasing remain transitive boundaries.

### SAV-FIRSTMOVE-880

- Frontend `R1284` calls `R1305`, which invokes world `R0414` at
  `L08404`. After normal return the frontend's nonzero `R1578` result
  selects `R0099(1)`. The resume argument skips its fresh-world constructor
  branch.
- Frontend `+6bc!=3` and `+6b8!=0` admit `L03780 -> R0147`, whose server
  `+2c` gate must also pass.
- Before ordinary `R0193`, old phase 6 can call `R0076`, and old phase 12
  calls actor full-tick `R0860` plus other world/player work.
- Ordinary `R0193` itself calls `R0191` and `R0426` before `R0427`.
- World post-load callbacks, archive cleanup, intervening frontend calls, those
  earlier server branches and attached Effect dispatch remain before the
  selected order consumer.
- The bounded search stops at 51 selected windows and the fourth outward
  actor-dispatch caller.
- Evidence: LOAD/resume/server bodies, explicit caller paths and retained
  unresolved frontiers.

**Confidence.** High for the named ordinary edges and branch order. Unknown for
the absolute first allocated-mover producer/read, the value at that read,
skipped-event cases and native resume timing.

**Unknown.** An earlier callback and a branch that skips this consumer remain
live alternatives; absence of a first-event witness is not evidence that no
event occurs.

## SpellTransport delivery counter

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-CASTCONT-1006 | SpellTransport's saved u16 `+4c` is the remaining delivery counter consumed by `R0634`. | High / Unknown | ● active | [EXP-0362](../experiments/EXP-0362-cast-delivery/) |

### SAV-CASTCONT-1006

- Serializer `R1602` writes `child+44`, `child+48`, then that word; its
  matching loading arm restores them.
- The transport's own post-load `R1115` performs Position/child repair calls
  and no counter reset.
- Conditional execution of the real tail preserves 0, 1, 2, 4, 7 and 10 through
  store/load and local post-load, then hands off after max(1,N) tick calls.
- The independent exact parser closes natural `game0018`'s 73436-byte document
  and reads counter 4 at decoded 47163 on the known
  SpellTransport/PointEffect/DirectDamage graph.

**Confidence.** High for the wire-to-counter bridge, the exact natural record
and the conditional transport behavior. Unknown for full native SAVE/reload,
transitive repair and first-frame scheduler order.

## Saved effect target and caster fields

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1010 | PointEffect's own `+44` is the target Unit reference set at cast-time construction; within the population searched, only `Serialize`'s load arm repairs it for a PointEffect, and the post-load hook never touches it. | High / Unknown | ● active (partially retracted) | [EXP-0371](../experiments/EXP-0371-saved-effect-identity/), [EXP-0388](../experiments/EXP-0388-effect-construction/) |
| SAV-1011 | SpellEffect's caster `+0x3c` is zero-written by both base constructors and absent from every concrete `Serialize` body; the LOAD-path factory that picks the constructor was not identified. | High / Unknown | ● active | [EXP-0371](../experiments/EXP-0371-saved-effect-identity/) |

### SAV-1010

- Live constructor `R1046` stores its second argument into `this+44` at
  `L05870`, unconditionally, immediately before dereferencing that same
  value's `+0x10` field to copy the target Position into its own Position
  (`L05872` at `L08023`).
- This preserves the target-Unit typing `MAGIC-SHAPE-008` assigns the
  constructor's second argument. The claim adds the exact store site and its
  unconditional order, not the typing.
- `Serialize` `L08388`'s load arm (`L08766..L08767`) reads the archived
  raw dword into a stack local, stores it at `this+44` (`L08768`), then passes
  `&this+44` to `R1108`. That body (`R1108..L08769`) looks the pre-store
  value up in the global identity map at `[L00285]+0x88` and, on lookup
  helper `R1488`'s boolean return, overwrites `+44` with the live pointer or
  with literal zero.
- PointEffect's separate post-load hook `L05904` (vtable slot `+0x24`) does
  not reference `+44` anywhere in its body. Within `Serialize`, the repair
  happens exactly once, inside its load arm, before the later Phase-2 manager
  walk (`SAV-TOKENLOAD-094`) runs any object's post-load hook.
- `EnumRefs callto:R1108` returns exactly 3 static call sites and 2 distinct
  owners: `L08388` and `R1482`, called twice (`L08770`, `L08130`). The
  census comes from the reused `-noanalysis` project's reference database, which
  cannot see a CALL inside code that project never disassembled. `R1482` is
  by vtable position (Unit's own `+0x24` slot) a plausible Unit post-load hook;
  it was not read.
- What a null `+44` does to a subsequently ticked object is the separate result
  `MAGIC-TARGETID-182`/`MAGIC-TICKGATE-183`.

**Confidence.** High for the constructor store/dereference order, `R1108`'s
hit/miss branch, the post-load hook's absence of a `+44` reference, and
`L08388` being the one of the search's 2 owners confirmed to target
PointEffect's own field. Each was read from the class's own disassembly against
the digest-verified programme.

**Unknown.** Whether a genuine identity-map miss is reachable by any admitted
native save/load population. What field(s) `R1482` repairs, and whether it
ever touches a PointEffect's own `+44`; it was not traced.

**Amended.** The constructor's call to `L05872` at `L08023` was first
glossed as registering the object on the target, "a Unit-shaped registration".
`SAV-1068` (EXP-0388) retracts that gloss ([`retracted.md`](retracted.md)):
`L05871..L08023` passes the target Position to assignment on the effect
Position, and `L05872` copies `+00..+05/+08..+0b` and performs no registration
(`SAV-TOKENPOS-074`). The `+44` target store, the immediate dereference and the
LOAD repair stand.

### SAV-1011

- There are two base `SpellEffect` constructors, not one. Both run at cast time,
  before any archive read:
  - `L08020`, called directly by PointEffect's cast-time constructor
    `R1046` at `L08771`, zero-writes `+0x3c` at `L08772`;
  - `L05979`, called by AreaEffect's cast-time constructor `R0652` at
    `L08773` and by SpellTransport's cast-time constructor `R0633` at
    `L05978`, independently zero-writes `+0x3c` at `L08774`.
- `SpellEffect::Serialize` `R1600` (`SAV-CLASSSER-173`'s 39-byte wire extent)
  stores and loads only `+0x40` and `+0x41`. PointEffect's `L08388`,
  AreaEffect's `R1601` and SpellTransport's `R1602` each call `R1600`
  first and add no `+0x3c` access of their own. `Serialize` performs no
  construction of its own, so if LOAD's object-creation path invokes one of
  these two constructors, the reconstructed object's `+0x3c` is zero going into
  `Serialize`, which does not change it.
- `R1629`, PointEffect's zero-target constructor variant (vtable
  `L05356`, `+0x44`/`+0x48` explicitly zeroed, calling `L08020`), is a
  structural candidate for that LOAD-path constructor. The archive's
  class-dispatch/factory table that would confirm it was not traced.
- Scoped to the traced LOAD path, a PointEffect's caster is null immediately
  after LOAD, whether its target-identity repair (`SAV-1010`) hits or misses.
  `MAGIC-ATTRGATE-118`'s delayed-kill-credit branch, gated on both a recorded
  caster and `+0x41`, therefore cannot pass its caster half for any PointEffect
  reloaded while still in flight, on that path.
- This does not extend to the object's whole post-load lifetime:
  - `R0003` (Spell::Apply, the function `MAGIC-TICKGATE-183` reads for
    registration) writes a newly cast-time-constructed effect's `+0x3c` a second
    time, after its zero-writing constructor returns, at `L08775` (PointEffect
    branch) and `L08776` (AreaEffect branch). Each value comes from
    `R0416`'s `[this+0x3c]==0` test applied to the function's caster
    argument: that argument's pointer when the test is false, 0 when true.
  - `MAGIC-AREATICK-036` documents a third `+0x3c` writer, `R0641`, the
    shared-list Tick driver `MAGIC-TICKGATE-183` proves also walks a reloaded
    PointEffect. It clears `+0x3c` to null when the referenced caster's `+0x14`
    is 0.
- No writer census beyond these four sites was run.

**Confidence.** High for both base constructors' unconditional zero-write and
for the absence of `+0x3c` from all three concrete `Serialize` bodies. Unknown
for the three points below.

**Unknown.** The archive's LOAD-path create/factory hook, and so which
constructor LOAD calls. Whether a PointEffect ever survives SAVE/LOAD while
still in flight; no in-flight-population census was run. Whether any post-load
restorer sets, rather than clears, `+0x3c` on a reloaded object outside the
sites named here; no full writer census was run, so this is not ruled out.

## Save label region: producer identity and a corpus witness

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-SAVELABEL-1016 | At the one SAVE-dialog construction site read, the label buffer's offset 0 is reset to an empty string just before construction; other callers of the constructor were not enumerated. | High | ● active | [EXP-0372](../experiments/EXP-0372-save-label-encoding/) |
| SAV-SAVELABEL-1017 | The SAVE dialog's `dialog+0x68` staging area is `0x200` bytes wide, but only its first `0x100` bytes are the on-disk label; the second half holds a separate per-item string. | High | ● active | [EXP-0372](../experiments/EXP-0372-save-label-encoding/) |
| SAV-SAVELABEL-1018 | A lawfully produced save in the preserved owner corpus, `game0006.sav`, carries the non-printable byte `0xE0` as live label content before the terminator. | High / Medium / Unknown | ● active | [EXP-0372](../experiments/EXP-0372-save-label-encoding/) |

### SAV-SAVELABEL-1016

- Populate/construct routine `R0701` executes the formation of the address 0x134 bytes into the object and
  a byte store of zero at that address at `L08777`/`L08778`, then passes that same
  address as the SAVE dialog constructor's 7th and final argument at
  `L08779`, immediately before `CALL R1245` at `L08780`.
- `R1245` stores that argument, an address, unchanged into dialog `+0x68`
  (`L08781`/`L08782`). `dialog+0x68` is a pointer field holding the address
  of `application+0x134`, not itself that location; every copy made through it
  writes into that one application-owned buffer.
- For the SAVE path this re-derives, address for address, the identity
  `SAV-LABELTAIL-236` (`EXP-0250`) derived by the same method for both SAVE and
  LOAD; it is not a new discovery.
- It rules out stale-buffer residue at label offset 0 specifically, at this one
  construction site, for a freshly opened SAVE dialog.

**Confidence.** High for the one traced construction site: the zero-store and
the pointer-field identity are direct instruction reads, cross-checked against
`SAV-LABELTAIL-236`'s independent derivation of the same addresses.

**Unknown.** Whether `R1245` has other callers that skip this reset was not
checked; the claim covers this one site, not every SAVE dialog construction.

### SAV-SAVELABEL-1017

- Command dispatcher `R0738` (cmd `0x444`) makes two calls to the same
  copy primitive `R1604`:
  - at `L08783`, into `dialog+0x68`, from a retrieved list-item text (the
    label);
  - at `L08784`, into `dialog+0x68+0x100` (`L08785`: an addition of 0x100 to the base), from
    a `CStringArray`-shaped field at `dialog+0x6c`, indexed by `dialog+0x80`,
    converted `CString`-to-`LPCTSTR`.
- The SAVE writer's verbatim `0x100`-byte disk emission (`SAV-EMB-004`,
  re-derived at `L08786`..`L08787`) reads only the first half. The second
  string is not an on-disk sibling of the label.

**Confidence.** High. A direct instruction read of both copy-primitive call
sites and their distinct sources, and of the writer's fixed `0x100` length
argument.

**Unknown.** The second string's semantic content; nothing read here establishes
it, and it is not shown to be a file path.

### SAV-SAVELABEL-1018

- `gameversions/saves/2026-08-02/game0006.sav` (29,345 bytes, well-formed
  `Asg&`-magic header) has label bytes `e0 66 69 6e 69 73 68 65 64 00` then 246
  zero bytes, read directly at the file's own `u32@0x04` offset (`SAV-EMB-004`'s
  offset convention).
- `0xE0` is outside 7-bit printable ASCII and precedes valid ASCII "finished".
- The owner's testimony naming "finished" as an intended save label is recorded
  with the corpus (`gameversions/saves/2026-08-02/MANIFEST.md`), not with this
  file. The manifest's mtime table places `game0006.sav` at 2026-08-01 22:41,
  before the 2026-08-02 00:20..00:24 session the quoted testimony describes, and
  lists this file "kept for contrast" rather than as one of that session's
  saves.
- It is the only file in this 12-file corpus whose label text is "finished", so
  the identification is the best available reading, not a confirmed same-session
  witness.
- The sibling save `game0009.sav` (label `666`) matches the owner's other named
  intent with no leading anomaly, so `game0006`'s leading byte is the exception
  in this small corpus rather than the rule.
- `SAV-SAVELABEL-1016`'s offset-0 zero-init rules out stale residue at this
  exact position for the one SAVE dialog construction site traced.

**Confidence.** High for the raw observation: the byte sequence, its position
and the file's well-formedness, all directly read and reproducible. Medium for
"committed content, not offset-0 residue": the zero-init fact discriminates only
at that one position and one traced construction site, no other position's
origin was checked and no live process was run. Unknown for the two points
below.

**Unknown.** The byte's exact producing mechanism, and what a LOAD/redraw or
fresh re-save of this file does with it. The experiment traced no process launch
and no confirmed document-loader body.

## Town-shape save: progress scalars, roster ranges, and the write and load paths

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1026 | An original town save holds campaign progress in two scalars, `campaign +0x04` and one byte per `Player` at `+0x3c`; no counted collection in the file is sized or indexed by the mission space. | High / Medium | ● active | [EXP-0375](../experiments/EXP-0375-sav-town-roster/) |
| SAV-1027 | The two progress-bearing fields, `Player+0x3c` and `campaign +0x04`, are written and read by two serializers in two halves of the file through different archive primitives; a consumer implementing one restores half. | High | ● active | [EXP-0375](../experiments/EXP-0375-sav-town-roster/) |
| SAV-1028 | A town save's roster is measured end to end with no roster count, and the array `Humanoid::Serialize` emits twelve of its references from is thirteen dwords wide, not twelve. | High / Medium | ● active | [EXP-0375](../experiments/EXP-0375-sav-town-roster/) |
| SAV-1029 | Of two skipped-index-0 save arrays, the equipment slot is never filled (`Armor::Equip` refuses 0 at `L04755`) but the `Spellbook`'s is writable; the empty-group prune loses nothing. | High / Medium / Unknown | ● active | [EXP-0375](../experiments/EXP-0375-sav-town-roster/) |
| SAV-1030 | The document load arm `R0414` never refuses: it has one exit instruction, and every conditional branch after the `IsStoring` split is a forward skip inside the arm. | High / Unknown | ● active | [EXP-0375](../experiments/EXP-0375-sav-town-roster/) |

### SAV-1026

- `campaign +0x04`, the current main mission, sits in the uncompressed tail; the
  `Player` byte at `+0x3c` sits in the compressed document.
- Measured on the two preserved streams the original engine authored end to end
  with no project-written input in their history. `savdoc -mode authorship`
  classifies each directory from its manifest; 31 of the 77 distinct streams it
  deduplicates by SHA-256 are `original-authored`.
  - `gameversions/saves/2026-08-02/game0010.sav`, 3,544 B, sha256 `b4e5ceb7…`;
  - `gameversions/saves/2026-08-15/game0010.sav` ≡
    `gameversions/en/game0010.sav`, 3,223 B, sha256 `89cfca4c…`.
- File-absolute regions, first then second: header `0..20` / `0..20`; word-coded
  document `20..2219` / `20..1940`; label block `2219..2475` / `1940..2196`;
  embedded `&YA1` store `2475..3234` / `2196..2955` (22 records, pool 27, both);
  uncompressed campaign record `3234..3544` / `2955..3223`. The campaign record
  is consumed field by field, 310 of 310 and 268 of 268 bytes: 36 fields for the
  first, 27 for the second, the difference being one child mission record.
- Decoded stream, 6,110 B / 5,592 B: head `0..75`; `Player` list `75..5696` /
  `75..5178` holding one `Player`; dead-actor list `5696..5700` / `5178..5182`
  with count 0; the world-half shape byte at decoded offset 5700 / 5182, equal
  to 0; then the tail to the `0xbadface1` trailer and the programme end at the
  body length, with no residue. That offset is a measurement; `SAV-CITY-030`
  could only infer it.
- Both saves stand at `campaign base +04 mission = 30`, the first with a child
  record at mission 31.
- Every variable-length collection in either save carries its own count, and the
  census reads each count off the stream rather than requesting it: 44 counted
  collections in the first save, 39 in the second, 323 over the twelve
  town-shape streams, each row naming its owner, the offset of its count, its
  element width and its end.
- The largest counts:
  - 119, six times per save: the `CDWordArray` at `Diary+0x04` and the
    `CWordArray` at `Diary+0x18` of each of the three `Diary` records a town
    save holds (one per actor plus the `Player`'s, `SAV-PLDIARY-054`). 119 is
    the `data.bin` Units table's own row count, and the subscript is the actor's
    class-dependent Units/Humans definition ordinal, written at `L08150` and
    `L08151`, not a mission (`SAV-667`, `SAV-844`). `SAV-668` adds that only
    six of the 119 indices are ever nonzero across 45 preserved saves: 64, 72,
    73, 88, 96 and 100.
  - 19, the `Spellbook` at `*(Unit+0x140)`, whose index space is the spell id
    1..28 (`MAGIC-SPELL-001`, `MAGIC-BOOK-002`).
  - 12, the carried container at `*(Unit+0x7c)` (`SAV-CARRY-050`).
  - Every remaining count in either save is 5 or less. The campaign record's
    largest is 15, the mercenary-type arrays.
- Every remaining field whose byte length is not an element count is a fixed raw
  span whose length is a literal inside the emitting helper:
  - `R1594` pushes `0x18` on both of its arms (`L08788` store,
    `L08789` load) and carries both `Unit+0xa6` and `Unit+0x114`;
  - `R1630` pushes `0x16` (`L08790` / `L08791`) for `Unit+0xbe`;
  - `R1570` pushes `0x40` (`L08792` / `L08793`) for `Unit+0xd4`.
- `SAV-UNITPROG-156` publishes those four widths together with `*(Unit+0x154)`
  180 and `*(Unit+0x158)` 148. `Humanoid+0x1cc` 24 and `Player+0x30` 32 complete
  the set; `SAV-662` measures the last as 31 always-zero bytes and one formation
  byte over 410 records.
- This carries `SAV-CAMPAIGN-077`'s record-scoped negative to file scope for
  this shape. The whole-stream closure the byte map is built on is `SAV-DOC-053`
  and `SAV-FULLREAD-252`, not a result of this claim.
- Population of the outer census: 130 preserved `.sav` file instances across the
  two install roots and `gameversions/saves`, of which 129 attribute with zero
  residue and one is not an `Asg&` save. Deduplicated by SHA-256 over the nine
  directories the tool's own table classifies `original`, 44 distinct streams
  are decoded and 43 complete the object walk. 12 of those 43 carry the
  world-half-absent shape byte: 2 `original-authored` (the two measured here), 5
  `original-resave-of-project-input` and 5 `project-written`; the last ten are
  not ROM1 evidence for content.
- The 44th stream, `gameversions/saves/2026-08-15/game0018.sav` (39,141 B on
  disk, 73,436 B decoded, sha256 `1e2eb21f…`, `original-authored`), passes the
  outer framing but halts the object walk at decoded offset 46,969 with
  `no programme for class "SpellTransport"`. It is censused, not classified: its
  shape byte is not trustworthy, and no range, collection or negative in this
  claim covers it.
- Two save directories carry no manifest and are reported
  `unclassified-directory` by the tool's own table rather than measured.
- Evidence files: `evidence/outer-framing.tsv`, `evidence/town-ranges.tsv`,
  `evidence/campaign-fields.tsv`, `evidence/authorship.tsv`,
  `evidence/disasm-stat-spans.txt`. Also cites `SAV-CAMPAIGN-077`,
  `SAV-CAMPAIGN-076`, `SAV-DOC-053`, `SAV-FULLREAD-252`, `SAV-CITY-030`,
  `SAV-SHAPE-023`, `SAV-UNITPROG-156`, `SAV-PLDIARY-054`, `SAV-667`, `SAV-668`,
  `SAV-844`, `SAV-662`, `SAV-CARRY-050`, `MAGIC-SPELL-001`, `MAGIC-BOOK-002`.

**Confidence.** High for the measured ranges, the shape-byte offset and the
collection census: two independently decoded streams, each tiling with zero
residue, and every count read off the stream. High that no counted collection in
a town save is sized or indexed by the mission space. The live alternative this
rules out is a mission-indexed bitmap or counted list held outside the campaign
record, the reading `SAV-CAMPAIGN-077` could not exclude because it read one
record. It is excluded by exhaustion plus one cited index domain per collection:
the census leaves no counted collection unnamed, the only one large enough to
span a campaign is the `Diary` pair, and its size comes from the Units table
while its subscript comes from the actor's own definition ordinal at two named
producer instructions. Medium for the stronger reading that no completed-mission
information at all is present; that is this claim's Medium half, not its
headline. A short list can carry mission ids as values without being indexed by
them, and this census measures sizes and subscripts, not element semantics:
`Unit+0x15c` (5 elements in the first save), `Unit+0x178` (3), `Group+0x20`,
`*(Group+0x3c)+0x4c` and `*(*(Unit+0x158)+0x90)` are all counted here, and none
of their element meanings is established. The interiors of the seven fixed raw
spans above were not decomposed either, and a per-actor span could in principle
carry per-mission bits; `Player+0x30` is the one such span excluded by
measurement.

### SAV-1027

- `Player+0x3c` is emitted inside `Player::Serialize` `R0415` by
  `L08794` loading the object pointer from the frame local at -0x30, `L06702` loading the byte at offset 0x3c of it,
  `L08294` pushing that byte and `L08795` calling `L08298`, the archive's one-byte
  write. It is restored by `L08796` loading the object pointer from the frame local at -0x30,
  `L06703` adding 0x3c to it, `L08797` pushing the sum and `L08798` calling `R1581`,
  the one-byte read.
- `SAV-FLAG-027` names `L08294` as the store and `MISSION-VICTORY-034` names
  `L06702`; those are the PUSH and the MOV of one three-instruction idiom. No
  ledger names the instruction that emits it: `claim -k "L08795"` over the
  base commit's claims tree returns no row.
- `Player+0x3d` follows immediately, through the same pair of primitives, at
  `L08799`/`L08800` and `L08801`/`L08802`.
- `campaign +0x04`, the current main mission (`SAV-CAMPAIGN-077`), is written
  neither by `Player::Serialize` nor by `ar <<` at all.
  - It is written at `L08803`, an indirect call through the slot loaded from `[archive vtable + 0x40]` at `L08804`, inside `R1631`, where the pushed arguments
    are the address `this+0x4` and the literal `0x4`: a raw span write over the
    field's own address.
  - It is read at `L08805`, an indirect call through the slot loaded from
    `[archive vtable + 0x3c]` at `L08806`, inside `R1632`, with the same pointer and length.
- `SAV-CAMPPROG-071` publishes the wire programme these two routines emit.
  Neither routine address appears in any ledger: `claim -k "R1631"` and
  `claim -k "R1632"` each return no row, as does `claim -k "campaign\+0x04"`.
- The EN and RU images are byte-identical at 1,977,344 bytes, so each
  instruction here is one observation, not two.
- Evidence files: `evidence/disasm-serializers.txt`,
  `evidence/disasm-campaign-base.txt`, `evidence/disasm-stat-spans.txt`,
  `evidence/corpus-searches.txt`, `evidence/digest.txt`. Also cites
  `SAV-FLAG-027`, `MISSION-VICTORY-034`, `SAV-PLAYER-028`, `SAV-CAMPPROG-071`,
  `SAV-CAMPAIGN-077`.

**Confidence.** High. The rival is that one routine pair writes and reads the
whole progress record, which a consumer could implement once. It is ruled out by
the two routines' disjoint bodies, both read end to end from the same image
listing: `R0415` never touches `campaign +0x04`, `R1631` never
touches a `Player`, and the two records land in different halves of the file,
one inside the word-coded document and one in the uncompressed tail. The vtable
identification is bounded: `L08803` and `L08805` are indirect calls and the
archive vtable's contents were not resolved, so "raw write" and "raw read" are
role names, not resolved symbols. They rest on symmetry plus a direct-call twin:
the store routine takes slot `+0x40` and the load routine slot `+0x3c` with
identical pointer and length arguments, and the three span helpers read whole
(`R1594`, `R1630`, `R1570`) branch on the same `IsStoring`
test and reach `R0271` on the storing arm and `R0272` on the other with
the same argument shape. The scope is bounded too: "progress-bearing" means the
two fields the field-by-field enumeration of a town save identifies
(`SAV-1026`), together with the fields the corpus already names as progress.

**Unknown.** A reading in which the two slots are some other symmetric pair is
not excluded. A third progress field elsewhere in the image would not have been
found by this search, which covered the two measured saves' attributed regions
and the four named routines.

### SAV-1028

- Decoded-stream ranges, `gameversions/saves/2026-08-02/game0010.sav`: head
  `0..75`; `Player` list `75..5696` with one `Player` record at `87..5696`; that
  `Player`'s fixed field run `94..145`, 51 B, with `+0x3c` at 119 and `+0x3d` at
  120; two group records at `149..3041` and `3041..4942`; two `Human` records at
  `248..3029` and `3131..4930`; `Player+0x30`'s raw 32 at `4942..4974`; the
  `Player`'s own `Diary` at `4974..5696`.
- There is no roster object and no party count anywhere between the list head
  and the first `Human`: membership is the nesting itself (`PARTY-ROSTER-002`).
- Per-`Human` record, each variable-length part with the delimiter that bounds
  it and its measured element count in the second `Human` of that save: `Token`
  37 B; `Unit+0x20` effect list, u32 count, 0; `Unit+0x15c` u16 list, u32 count,
  0; `Unit+0x178` u16 list, u32 count, 0; raw `Unit+0xa6` 24, `Unit+0xbe` 22,
  `Unit+0x114` 24, `Unit+0xd4` 64; `*(Unit+0x154)` 180; `*(Unit+0x158)` 148;
  `*(*(Unit+0x158)+0x90)` u16 list, count 0; a 19-byte scalar run; `Unit+0x74`
  one object reference; `Unit+0x78` one reference; `Unit+0x80` CString, 9 B; a
  55-byte scalar run; `Unit+0x68` one reference; `Unit+0x7c` presence byte then
  the carried container, 0 elements here and 12 in the first `Human`;
  `Unit+0x140` presence byte then the `Spellbook`, 80 B and count 19 here,
  absent in the first `Human`; a 17-byte scalar tail; raw `Humanoid+0x1cc` 24 B;
  then thirteen object references (`SAV-HUMAN-043`).
- The six raw widths and their order are `SAV-UNITPROG-156`'s, re-measured on
  the wire rather than re-derived.
- The array the first twelve references come from holds thirteen dwords. The
  constructor `R0876` zero-fills the dword at offset 0x198 plus 4 bytes times the slot index at `L08807`
  under `L08808`, a 32-bit store of zero to the frame local at -0x4, and
  `L08809`, a comparison of that 32-bit local with 0xd, so `i = 0..12`. The same constructor
  then zero-fills six dwords at `+0x1cc` (`L08810`, bound `0x6`), the 24-byte
  raw span above.
- The destructor `R0982` walks `i = 1..12` (`L08811`, `L08812`) and
  deletes each non-null element, so the index the serializer skips is the index
  the destructor skips. Element 0 is allocated storage that the load arm does
  not write, the store arm does not emit, and nothing frees.
- The thirteenth reference is `Humanoid+0x1e4`, outside the array, zeroed
  separately at `L08813`.
- Evidence files: `evidence/town-ranges.tsv`, `evidence/disasm-equip-slots.txt`,
  `evidence/disasm-serializers.txt`, `evidence/refs-skipped-slots.txt`. Also
  cites `PARTY-ROSTER-002`, `SAV-HUMAN-043`, `SAV-UNITPROG-156`,
  `SAV-CARRY-050`, `SAV-SPELLBK-041`, `ITEM-EQUIP-006`, `SAV-DOC-053`,
  `SAV-FULLREAD-252`.

**Confidence.** High for the byte ranges and the per-field delimiters. Every
offset is a measured pair whose spans tile the record with no residue, in two
independently decoded streams, the two of the twelve town-shape streams decoded
that are `original-authored`; the per-field list is read from the second `Human`
record of `gameversions/saves/2026-08-02/game0010.sav`. The same programme
closes `SAV-FULLREAD-252`'s preserved population, 55 paths deduplicated to 31
distinct digests. It does not close
`gameversions/saves/2026-08-15/game0018.sav`, `original-authored` and preserved,
where the walk halts at decoded offset 46,969 with
`no programme for class "SpellTransport"`, the stream `SAV-1026` names as
uncovered. The live alternative ruled out is a flat roster, one counted array of
character records, which predicts a count between the list head and the first
character record; the measured stream has a `Player` record and two group
records there and no such count. High for the array being thirteen dwords wide:
the constructor's loop bound is the literal `0xd` against a zero-based index, in
the same routine that zeroes `+0x1cc` with bound `0x6` and `+0x1e4` singly, so
the three widths come from one instrument. Medium for reading `Unit+0x74`,
`Unit+0x78` and `Unit+0x68` by those structure offsets: the names are the
serializer's own operands, not an independent consumer's.

**Unknown.** What array indices 1 and 2 hold. `ITEM-EQUIP-006` carries
`actor+0x19c` and `actor+0x1a0` as Unknown, and they remain so.

### SAV-1029

- The empty-group prune is not a loss.
  - `SAV-ROSTER-024` establishes the file side: the save-file writer
    `R0804` removes from a qualifying player's group collection every
    group whose count is 0 (`L08278`, `R0197`, `R1582`), so an
    empty group never reaches the file.
  - The memory side is the removal's operands, which no ledger names
    (`claim -k "L08814"` over the base commit's claims tree returns no row).
    `L08815` / `L08816`, a push of the group pointer, pushes the group; the receiver is reloaded
    from the player at `L08817`, loading the player pointer from the frame local at -0xa4, and
    `L08814`, loading the field at offset 0x24 of it, the player's own live group collection, the
    object the iterator at `L08818` was constructed over (`L08819`).
  - The removal mutates live memory, not a serialization copy, so after SAVE
    file and memory agree: memory lost the group too. This is the writer cleanup
    `SAV-CITYSTORE-516` left outside its alias proof.
- `Humanoid+0x198` index 0 is not a loss.
  - Whole-image census `EnumRefs disp:198`: 78 hits, 41 distinct owners, 0 in
    orphan or undisassembled code. The indexed form `[base + index*4 + 0x198]`
    is 18 instructions in 10 owners: `L08820`; `L04614`; `L08821`,
    `L08822`, `L08823`, `L08824`; `L08807`; `L08825`, `L08826`,
    `L08827`; `L04746`; `L08828`; `L04718`, `L08829`, `L04720`;
    `L08830`; `L08831`, `L08832`.
  - Exactly two can leave a non-null element: `L04720` in `Armor::Equip`
    `R0848`, and `L08832` in the serializer's load arm inside the
    `1..12` loop.
  - `R0848`'s first test refuses index 0 outright:
    `L04710` loading the byte at offset 0x50, `L08833` testing it, and
    `L04755` branching to `L08834` when nonzero, the fall-through printing at `L08835` and
    returning at `L08836` (`SAV-CARRY-050`).
  - Every other indexed write stores zero: `L08807`, `L08827`, `L08830`.
    Index 0 is never filled, so skipping it drops nothing.
- The `Spellbook`'s index 0 is a loss.
  - `SAV-SPELLBK-041` publishes the skip and both mirrored arms, and
    `MAGIC-BOOK-002` the setter's grow-and-store behaviour; this claim adds that
    the skipped slot is reachable.
  - `R0464`, read whole from `R0464` to `L08837`, contains no test
    of the id beyond `L08838`, a comparison of the id with the first stack argument, against the current size: a
    short array is grown to `id+1` through `R1588` at `L05048`, and the
    element is stored at `L05049` unconditionally.
  - The load arm's `SetSize(n,-1)` reaches the same `R1588`, whose
    fresh-array arm allocates at `L08839` and then calls `R1633` at
    `L08840`. That routine's first act is a three-argument call at `L08841`
    taking the buffer, the literal `0` and `n*4`, a zero fill of the whole new
    span, so element 0 after a load is null rather than residue.
    `R1633`'s following per-element loop at `L08842` was not
    identified.
  - Census `EnumRefs callto:R0464`: 33 hits, 7 distinct owners, 0 in orphan
    code. 25 of the 33 push a literal id, all in the spawn routine
    `R0184` and all in {2, 3, 8, 0xb, 0xd, 0xe, 0x11, 0x14, 0x1b}. One
    more literal, `0xd`, is pushed outside it at `L08843` in `R1634`
    (`claim -k "L08843"` returns no row).
  - The remaining 7 push a non-literal index: `L08844`, also in
    `R0184`, whose data-driven arm tests the value first
    (`L08845` comparing the 32-bit local at frame offset -0x1c with zero and `L08846` branching to `L08847` when less or equal skips
    the whole construct-and-store when the table value is 0 or less), and
    `L08848`, `L08849`, `L08850`, `L08851`, `L08852` and `L08853`.
  - Only `L08853`, the `teachSpell` effect arm inside `R0843`, was
    traced to its index's origin. Its two preceding guards are a no-`Spellbook`
    skip (`L05050` / `L08854` / `L08855 JZ`) and an already-present skip
    through `R0017` (`L08856` / `L08857` / `L08858 JNZ`).
    `R0017` tests only the index against the array size and the slot
    against null, so neither guard rejects 0.
  - The pushed index (the frame local at -0x10) is the product of the effect's own `+0x40` and
    the routine's second argument (`L08859` or `L08860`), so an index of 0
    is produced whenever either factor is 0; it is not a truncation artefact.
  - The same arm narrows that product to a byte for the `Spell` object's own id
    field (`L08861`). `R0463` stores it at `L08862` and returns the
    object unconditionally at `L08863`. The id resolver `R0624` takes
    its zero arm at `L08864`/`L08865`, emits the `Invalid spell #0` message
    `MAGIC-SPELL-001` quotes, and reaches the epilogue by
    `L08866`, a jump to `L08867`, without aborting.
  - A teachSpell effect applied with a zero factor therefore places a live
    `Spell` at element 0, and the next SAVE writes `n = 1` with zero references.
- Corpus population: the 12 town-shaped streams decoded, of which 2 are
  `original-authored`. The four `Spellbook` records in that population all carry
  count 19, so the branch is named from the image and not witnessed in the
  corpus.
- Unresolved discrepancy: `MAGIC-BOOK-002`'s Unknown clause says 32 of the 33
  `SetAt` sites are in `R0184` and `R0656`. This census agrees on
  the total, 33, and on the owner set, but attributes 26 to `R0184` and 1
  to `R0656`: 27, not 32. `evidence/refs-spellbook.txt` lists every site
  with its owner; `magic.md` is not corrected by EXP-0375.
- Evidence files: `evidence/disasm-roster.txt`,
  `evidence/disasm-equip-slots.txt`, `evidence/disasm-spellbook.txt`,
  `evidence/disasm-spellbook-producers.txt`, `evidence/refs-skipped-slots.txt`,
  `evidence/refs-spellbook.txt`, `evidence/town-ranges.tsv`,
  `evidence/corpus-searches.txt`. Also cites `SAV-ROSTER-024`,
  `SAV-CITYSTORE-516`, `SAV-SPELLBK-041`, `SAV-CARRY-050`, `MAGIC-BOOK-002`,
  `MAGIC-SPELL-001`.

**Confidence.** High that the image contains this asymmetry. The live
alternative, that both skipped slots are unreachable storage and neither skip
can lose anything, is the status `Humanoid+0x198[0]` does have. The two bodies
discriminate: the equipment writer refuses index 0 at a named instruction
between entry and store, and the spellbook writer has no instruction of that
kind anywhere between entry and store. High for both non-losses: each rests on
the named instructions above (the reloaded live receiver, the refusal) plus a
whole-image displacement census with its own totals. High that nothing on the
traced path filters a zero index, the headline's claim. The live alternative,
some test between the product and the store that rejects 0 as `Armor::Equip`
does, is excluded by enumeration: the index local (the frame local at -0x10) is assigned
exactly twice in `R0843`, at `L08868` and `L08869`, both before the
effect-kind switch at `L04017`, so the value stored at `L08853` is the entry
product unchanged (a count of 32-bit stores to the frame local at -0x10 over the
function's section of `evidence/disasm-spellbook.txt` returns 2), and the arm's
three tests are the two guards above and a failed-allocation test (`L08870`).
Medium for reachability in the running game, which the headline does not claim:
six of the seven non-literal call sites are untraced; `L08848`, `L08849`,
`L08850`, `L08851` and `L08852` push locals not followed, and `L08844`
was followed only as far as its own `JLE` guard. Unknown whether any shipped
effect ever supplies a zero factor: no effect-table census was run, and EXP-0375
read the image and the preserved save corpus only, never the shipped data
tables. The census is bounded: `EnumRefs disp:198` finds instructions whose
encoded displacement is `0x198`, and a write through an address computed into a
register beforehand carries no such displacement and would not appear.

### SAV-1030

- `L05932`, a call to `R1123`, and `L08871`, a branch to `L05933` when zero, select the load arm.
  From `L05933` to the function's end the listing holds 24 conditional
  branches and one `RET`, at `L08872`, so there is no error exit for a
  validation failure to reach.
- What the arm does with an empty or semantically wrong record:
  - the `world+0x84` difficulty is substituted:
    `L08284` comparing the 32-bit local at frame offset -0x14 with 0x1 and branching to `L08873` when less, and
    `L08285` comparing it with 0x3 and branching to `L08873` when greater, gate the store at `L08429`, so an
    out-of-range value silently keeps the constructor's (`SAV-HEAD-025`, whose
    clamp clause this confirms from the same instructions);
  - the player list is read at `L08275` / `L08155`, a call to `R1372`, with no
    count test and no return test, so a count of 0 yields an empty roster and
    the arm continues;
  - `Player+0x3c` is restored raw at `L06703`/`L08798` with no test against
    its own `{0,1,2}` domain (`SAV-FLAG-027`, `SAV-1027`);
  - a trailer that is not `0xbadface1` merely skips the following read:
    `L08053` comparing the 32-bit local at frame offset -0x10 with 0xbadface1 and `L08502` branching to `L07783` when not equal
    jump over the push of L07886 at `L08874` into the same common tail, which is not
    a rejection (`SAV-648`, `SAV-FULLREAD-252`).
- On the world-half-absent arm reached at `L08875`/`L08876` the loader then
  walks every restored actor and zeroes `+0x40`, `+0x44` and `+0x5c`
  (`L08276`, `L08877`, `L08277`; `SAV-ROSTER-024`, `PARTY-LOSS-006`).
- "Malformed" here means semantically wrong: an empty count, an out-of-range
  enum, an unrecognised trailer. It does not cover a truncated stream.
- `SAV-FULLREAD-252` carries malformed-input acceptance as Unknown; this claim
  answers it for one arm and one meaning of malformed.
- Evidence files: `evidence/disasm-serializers.txt`,
  `evidence/disasm-loaders.txt`. Also cites `SAV-HEAD-025`, `SAV-DOC-053`,
  `SAV-FULLREAD-252`, `SAV-ROSTER-024`, `SAV-FLAG-027`, `SAV-648`,
  `PARTY-LOSS-006`.

**Confidence.** High that the arm neither refuses nor validates beyond the one
substitution. The live alternative, that a malformed record is detected and
rejected somewhere the single-field readings had not reached, is ruled out by
enumeration rather than sampling: every conditional branch in the arm was
listed, every target is forward and inside the arm, and the function has one
exit. Truncation is outside the scope: the archive primitives `R0686` and
`R1581` were not read in EXP-0375, MFC's own short-read behaviour was not
established, and a throw from inside a primitive would not appear as a branch in
this listing. Unknown for downstream faults, below.

**Unknown.** Whether the load arm's permissiveness becomes a fault downstream.
The experiment's discriminator required naming the consumer instruction that
dereferences an emptied record without a guard, and no such instruction was
named. The search that bounds this Unknown is `R0414` itself plus the
four loader routines exported with it, `R1635`'s enclosing body,
`R1325`, `R1530` and `R1490`; no consumer of the player
list or of a restored `Spellbook` outside them was read.

## Town-shape save: roster record stat spans

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1031 | A roster record carries the 24-byte stat layout twice, the 22-byte layout once and the 64-byte modifier block, each through one direction-symmetric helper; its scalar words are stored post-fold. | High / Medium / Unknown | ● active | [EXP-0375](../experiments/EXP-0375-sav-town-roster/) |

### SAV-1031

- Before the `IsStoring` split at `L06857`/`L08878`, `Unit::Serialize`
  `R0210` reaches four fixed spans by adding a literal offset to the object pointer:
  - `L08879` → the call to `R1594` at `L08376` from `+0xa6`;
  - `L08880` → the call to `R1630` at `L08881` from `+0xbe`;
  - `L08882` → the call to `R1594` at `L08883`, the same helper, from `+0x114`;
  - `L08884` → the call to `R1570` at `L08885` from `+0xd4`.
- Each helper, read whole, is its own write site and read site. It tests
  `IsStoring` (`R1123`), then pushes one literal length and calls the
  archive's raw span write `R0271`, or pushes the identical literal and calls
  the raw span read `R0272`: `R1594` pushes `0x18` at `L08788` and
  `L08789`, `R1630` pushes `0x16` at `L08790` and `L08791`, `R1570`
  pushes `0x40` at `L08792` and `L08793`.
- The 24/22/24/64 widths `SAV-UNITPROG-156` and `SAV-UNITLEN-045` publish are
  therefore literals inside three routines: store and load cannot disagree, and
  nothing is recomputed on either arm.
- `EnumRefs callto:R1594` is 4 hits in 3 owners, exactly two of them these
  two sites in `R0210`; `callto:R1630` is 5 hits in 5 owners;
  `callto:R1570` is 1 hit in 1 owner.
- `SAV-UNITPROG-156` leaves the meanings inside the raw blocks Unknown. The
  offsets name them against `HERO-MOD-016` and `UNIT-CTOR-004`:
  - `+0xa6` is the live copy and `+0x114` the base copy of
    `[toHit u16][6 skill u16][dmgBase u8][dmgSpread u8][active u8][5 u8]`;
  - `+0xbe` is the live copy of
    `[defence u16][absorption u16][6 protection u16][6 damage-kind u8]`;
  - `+0xd4..+0x113` is the modifier block, holding both layouts' modifiers
    (`+0xe6`, `+0xfe`) and the five scalar modifiers.
- Byte-level corroboration needs no constructor: one helper carries `+0xa6` and
  `+0x114`, so the file holds two records of one 24-byte layout 0x6e apart,
  which a live/base pair predicts and a single live copy does not.
- The scalars are stored folded. The store arm emits fourteen consecutive u16
  from `+0x84` to `+0x9e` (`L08886` … `L08887`, each through `R1583`),
  then bytes `+0xa2` and `+0xa3` (`L08888`, `L08889`), then `+0xa0`
  (`L08890`) and `+0xa4` (`L08891`, through `L08299`). `HERO-MOD-016`'s
  fold writes its five scalar adds into `+0x8c`, `+0x92`, `+0x96`, `+0x9c` and
  `+0xa4`: four inside the fourteen-word run and the fifth the last word
  emitted. What reaches the file at those five offsets is the post-fold value.
- `SAV-CITYSTORE-516` establishes that the store arm writes these spans and
  words from their existing values without recomputing them, and
  `SAV-HUMLOAD-445` that the load bodies restore the modifier block rather than
  rebuilding it. This claim names which restored bytes are which.
- Consequence for a consumer:
  - For the 24-byte layout both forms are in the file, and neither has to be
    derived.
  - For the 22-byte layout and the five folded scalars the file carries live
    and modifier but no base. A base is recoverable only by subtraction and
    only as far as the fold is additive. `HERO-ARMOUR-018` as amended makes all
    fourteen fields of the 22-byte layout additive. `HERO-MOD-016` records that
    `R0846`'s last field is assigned rather than added, so that one field
    of the 24-byte layout is not recoverable by subtraction at all and is
    available only because `+0x114` is in the file.
  - `UNIT-GATE-013`'s spawn adjustment writes `actor+0xa6 += 50` and
    `actor+0xbe += 50` and does not touch `+0x114`, so live minus modifier is
    not in general the base, even for the 24-byte layout.
- Prior art, searched as the subject over the base commit's claims tree:
  `claim -k "actor\+0x114"` returns exactly two rows, `HERO-MOD-016` and
  `UNIT-CTOR-004`, both read whole; `claim -k "0x114"` returns 23, of which
  `SAV-598` and `SAV-599` concern the campaign record's own `+0x114`, not an
  actor's. `claim -k "R1594"` returns one row (`SAV-HEROSKILL-064`),
  `claim -k "R1570"` one (`SAV-REGENWIRE-532`), and `claim -k "R1630"` no
  row.
- Other writers of the base copy are not bounded. `EnumRefs disp:114` over the
  whole image returns 49 instructions in 29 owners, none between `L02478` and
  `L06646`, the span containing every actor-family routine read. The
  instrument finds only an encoded displacement of `0x114`, and this
  serializer's access is an addition of 0x114 to the object pointer followed by an indirect span copy,
  which carries no such displacement, so no bound on who writes `+0x114` is
  claimed.
- Evidence files: `evidence/disasm-serializers.txt`,
  `evidence/disasm-stat-spans.txt`, `evidence/refs-stat-spans.txt`,
  `evidence/town-ranges.tsv`, `evidence/corpus-searches.txt`. Also cites
  `SAV-UNITPROG-156`, `SAV-CITYSTORE-516`, `SAV-HUMLOAD-445`,
  `SAV-REGENWIRE-532`, `SAV-UNITLEN-045`, `HERO-MOD-016`, `HERO-ARMOUR-018`,
  `UNIT-CTOR-004`, `UNIT-GATE-013`.

**Confidence.** High for the mechanism: three direction-symmetric helpers with
one literal length each, two of the four spans sharing one helper, and the
fourteen-word scalar run. The live alternative is that a roster record stores
base values and the engine refolds them at load, a shape a consumer would have
to implement. Instructions rule it out rather than inference: each helper's two
arms push the same literal and differ only in the primitive they call, so
neither arm computes anything; `HERO-MOD-016`'s fold destinations are four of
the fourteen serialized words plus `+0xa4`, so the serialized scalars are
post-fold; and `SAV-HUMLOAD-445` finds no modifier rebuild in the traced load
bodies. Medium for the live/base/modifier naming itself. That identification is
`HERO-MOD-016`'s and `UNIT-CTOR-004`'s; this claim corroborates it only by the
shared helper and the shared width, and the fold `R0840` and the five
actor constructors were not read. Nothing in this listing excludes a reading in
which `+0x114` is a second live copy rather than a base.

**Unknown.** What the individual bytes of the 22-byte and 64-byte spans mean
beyond the field lists the cited claims give, and whether any post-load
consumer refolds.

## Effect_DirectDamage's own 24-byte raw span: LOAD constructor, shared block identity, one corpus witness

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1032 | Effect_DirectDamage's descriptor creator field `R1636` only wraps the constructor body a live cast calls, so every located construction initializes the 24-byte span identically; a LOAD call to it is not traced. | High / Medium / Unknown | ● active | [EXP-0377](../experiments/EXP-0377-damage-span/) |
| SAV-1033 | Effect_DirectDamage's 24-byte span at `+0x48` is the known `actor+0xa6` combat-block layout; within the traced call graphs only `+0x13`/`+0x14`/`+0x15` have a located writer for this class. | Medium | ● active | [EXP-0377](../experiments/EXP-0377-damage-span/) |
| SAV-1034 | LOAD does not re-run the school setters: a loaded Effect_DirectDamage's span is the 24 bytes on disk, and the one corpus-witnessed record (n=1) is consistent with the writer and reader findings. | Medium | ● active | [EXP-0377](../experiments/EXP-0377-damage-span/) |

### SAV-1032

- `SAV-TOKENLOAD-093`'s descriptor table names `R1636` as this class's
  creator field without disassembling it. `R1636`
  (`R1636`..`L08892`, `evidence/functions.txt`) is 30 instructions, most
  of them prologue, an SEH unwind frame and epilogue around `new(0x60)`
  (`L08893`/`L08894`), a null check, and a call to `R1042` on success
  (`L08895`). It makes no field write of its own.
- `R1042`, the class's constructor body, has exactly three call sites
  image-wide, a complete and orphan-free population (`EnumRefs callto:R1042`:
  3 hits, 2 distinct owners, 0 orphan): `R0003` (`MAGIC-DMG-005`'s
  spell-apply dispatcher) at `L08896` and `L08897`, and `R1636` at
  `L08895`.
- Every located construction of the class, whichever site runs, is therefore
  built by one body: base `Effect` construction (a call to `R1049`);
  adding 0x48 to the object pointer and calling `R1044`, which zero-initializes the span through
  `SAV-HUMGAPS-449`'s block initializer with the pointer already offset to the
  span's start; vtable `L02980`; and `+0x44 = 0`, the caster, overwritten only
  on the cast-time path (`MAGIC-DMG-005`).
- `callto:R1636` returns 0 hits. The only connection between `R1636` and
  LOAD is `SAV-TOKENLOAD-093`'s descriptor-table record naming it as this
  class's creator field: a recorded field, not a traced invocation. `SAV-774`
  reads `CArchive::ReadObject`'s class check and the statically initialised
  descriptor set but traces no creator-field invocation.
- This claim settles only what `R1636` does if the archive invokes it: its
  single call target, identical to a live cast's.
- `go run ./tools/claim -k "R1636"` over the base commit's claims tree
  returns no row.
- Evidence files: `evidence/functions.txt`, `evidence/refs.txt`,
  `evidence/program-digest.tsv`. Also cites `SAV-TOKENLOAD-093`, `SAV-1011`,
  `SAV-774`, `SAV-HUMGAPS-449`, `MAGIC-DMG-005`.

**Confidence.** High for the two functions as disassembled from the current
Ghidra project. The live alternative is a constructor variant that skips the
span's zero-init, or seeds it with different defaults, depending on which of
the two located call sites runs. It is excluded because `R1636` carries no
code beyond the allocation and the one call, and `R1042`'s caller census is
complete and orphan-free for the analysed image. This is not a statement that
LOAD reaches `R1636`. Medium for the LOAD link, inherited from
`SAV-TOKENLOAD-093`'s descriptor-field record alone: narrowed by naming the
field, not closed.

**Unknown.** Whether the archive's LOAD-path dispatch calls `R1636`. For the
sibling SpellEffect family, `SAV-1011` leaves the same question Unknown: "this
experiment did not identify the archive's own create/factory hook that
allocates a class instance during LOAD". That gap is not closed here.

### SAV-1033

- The block's type is already identified, and is cited here:
  - `MAGIC-DMG-005` identifies the span (`+0x48`, `SAV-CLASSSER-173`) at High
    as "a `0x16`-byte combat block of the `actor+0xa6` layout";
  - `SAV-HUMGAPS-449` and `UNIT-CTOR-004` name that block generically;
  - `SAV-1031` names it field by field at its other embedding
    (`actor+0xa6`/`+0x114`:
    `[toHit u16][6 skill u16][dmgBase u8][dmgSpread u8][active u8][5 u8]`),
    corroborating `HERO-MOD-016`;
  - `HERO-DAMAGE-022`, `UNIT-STRUCTDAMAGE-064` and `MAGIC-RESIST-006` read it
    field for field through the same two resolvers.
- Dispatch chain: Effect_DirectDamage's vtable slot `+0x3c` is `R0564`
  (`UNIT-AREADIRECT-072`). Read whole, it pushes `this+0x44` (the caster) and
  `this+0x48` (the span's start) as the generic resolvers' `(A, attacker)` pair:
  - to `target->vtable[+0x4c]` on a Unit target (`L05510`), matching
    `HERO-DAMAGE-022`'s `R0265` instrument; `callto:R0265` reproduces
    as 3 hits, all `.rdata` slots, 0 direct callers;
  - directly to `R0653` on a Building target (`L08898`,
    `UNIT-STRUCTDAMAGE-064`).
- Span positions those resolvers read, ten of twenty-four in total:
  - `HERO-DAMAGE-022`: `+0x00` (i16 to-hit base), `+0x0e`/`+0x0f` (physical
    base/spread pair), `+0x10` (skill-slot resistance index), `+0x15` (copied
    to `target+0x48` when the attacker is a mage);
  - `MAGIC-RESIST-006`: `+0x11`/`+0x12` (a second, not-hit-gated component)
    and `+0x13`/`+0x14` (the elemental pair, admitted unconditionally because
    "an Effect_DirectDamage's block carries no physical pair");
  - `UNIT-STRUCTDAMAGE-064`: an explicit negative on
    `+0x0e`/`+0x0f`/to-hit/defence/absorption/protection/`+0x15` for the
    Building path.
  - No reader is located, at either embedding, for `+0x02..+0x0d`,
    `HERO-MOD-016`'s six skill `u16`.
- Writers for this class: `R0003`'s two construction sites and the five
  school setters (`R1043`/`e90`/`ec0`/`ef0`/`f20`;
  `EnumRefs callto:R1043/e90/ec0/ef0/f20`: 2/1/1/1/1 hits, each sole owner
  `R0003`) write only `+0x13`/`+0x14`/`+0x15` (`MAGIC-DMG-005`'s
  damageMin/spread/school). A read around both `R1042` call sites in
  `R0003` (1743 disassembled instruction lines,
  `evidence/functions3.txt`) finds no other literal store into `+0x48..+0x5f`.
  The two resolver bodies read nothing at `+0x01..+0x0d`, `+0x16` or `+0x17`.
- For this class, therefore:
  - `+0x00`/`+0x0e..+0x12` are structurally present but permanently zero: dead
    reads of the constructor's zero-init, never the
    toHit/skill/dmgBase/dmgSpread/active values that name gives them at
    `actor+0xa6`.
  - `+0x01..+0x0d`/`+0x16`/`+0x17` have no located writer or reader.
    `SAV-HUMGAPS-449`'s "final two unchanged" applies to `+0x16`/`+0x17` in
    this embedding by the same initializer, confirmed rather than assumed.
  - Both positions Q3 scopes outside the span resolve to the target. `+0x10`
    selects one of `target+0xc4..0xcc`/`target+0xce+i` (`HERO-DAMAGE-022`),
    always index 0 for this class since nothing sets it. `+0x15` selects
    `target+0xc4/c6/c8/ca/cc` by the jump table at `L04134`
    (`MAGIC-RESIST-006`); it is the one populated field whose meaning is
    defined entirely by the target's resistance array.
- Evidence files: `evidence/functions.txt`, `evidence/functions2.txt`,
  `evidence/functions3.txt`, `evidence/refs.txt`, `evidence/refs2.txt`,
  `evidence/refs3.txt`. Also cites `SAV-CLASSSER-173`, `SAV-HUMGAPS-449`,
  `UNIT-CTOR-004`, `SAV-1031`, `HERO-DAMAGE-022`, `UNIT-STRUCTDAMAGE-064`,
  `MAGIC-RESIST-006`, `MAGIC-DMG-005`, `UNIT-AREADIRECT-072`, `HERO-MOD-016`.

**Confidence.** Medium. The cited clauses are High in their own claims: the
generic resolver mechanics, and the shared-type identification, which is
`MAGIC-DMG-005`'s High finding, cited rather than re-derived. Medium for this
claim's own contribution, the dispatch-chain address confirmation
(`R0564` feeding this span to those resolvers) and the class-specific
writer and reader census. It is bounded to the call graphs
`callto:R1042`/`callto:R1043`-`f20`/`callto:R0564`/`callto:R0265`/`callto:R0653`
reached, each complete and orphan-free for its own target address, plus one
instruction-level read of `R0003`. It is not an image-wide census of
every place that might hold an `Effect_DirectDamage*` and write its span
through an unenumerated path; that live alternative, an unlocated writer or
reader outside the traced graph, is not excluded.

**Unknown.** `evidence/refs.txt`'s two vtable dumps show `L02980` differing
from base `Effect`'s `L02981` in exactly four of twenty slots: `+0x00`,
`+0x04`, `+0x08` (`SAV-1034`) and `+0x3c` (traced here). The other two,
`R1637` and `R1638`, were not traced.

### SAV-1034

- `Effect_DirectDamage::Serialize` (`L08389`) calls base `Effect::Serialize`
  (`L08329`), then adding 0x48 to the object pointer and calling `R1594(archive, blockPtr, 0x18)`.
  That generic raw-span helper tests `IsStoring` at `R1123`, then calls
  archive-write `R0271` or archive-read `R0272`, with no field-level logic
  in either arm: it copies the full 24 bytes verbatim in both directions.
- `SAV-1031` reads the same helper for `Unit::Serialize`'s two calls and
  reports its caller census as 4 hits/3 owners. `EnumRefs callto:R1594`,
  re-run independently, gives the same 4/3: `R0210` (`SAV-1031`'s
  `Unit::Serialize`) supplies two, `Effect_DirectDamage::Serialize` `L08389`
  one (at `L08899`), and the fourth, `R1639` (`L08900`), was not
  identified.
- The LOAD arm of `Effect_DirectDamage::Serialize`, the `IsStoring`-false branch
  of `R1594`, is the entire LOAD-time span operation traced. It never
  calls `R0003` or any of the five school setters (`L08389` and
  `R1594` read whole, `evidence/functions.txt`). This holds whichever of
  `SAV-1032`'s two located construction sites built the object, and whether or
  not the archive's LOAD dispatch reaches `R1636` (`SAV-1032`'s LOAD link is
  Medium, not High).
- On this traced path a loaded effect's `+0x13`/`+0x14`/`+0x15` are the bytes
  captured at save time, frozen against any later change to the spell's damage
  columns, and its dead or uninitialized positions persist unexamined for the
  object's remaining lifetime.
- Corpus check, n=1. `SAV-CLASSSER-177`'s full population is 55 paths / 31
  SHA-distinct saves, with exactly one witnessed `Effect_DirectDamage` record,
  in `game0018.sav` (raw sha256 `1e2eb21f…`, byte for byte the live
  `gameversions/en/game0018.sav`). The probe
  `experiments/EXP-0377-damage-span/probe` decodes that file (73,436-byte body,
  matching `SAV-CLASSSER-177`) and reads `decoded_body[47133:47157]`:
  - `+0x00`/`+0x0e..+0x12` are all `0x00`, agreeing with "dead, never written";
  - `+0x13`/`+0x14`/`+0x15` are `0x04`/`0x04`/`0x01`, a damage 4, spread 4,
    school-1/Fire triple: a plausible, in-range instance of `MAGIC-DMG-005`'s
    field identification, not checked against a specific `Data.bin` spell row;
  - `+0x16`/`+0x17` are `0xc6`/`0x02`, nonzero: the observed instance of
    `SAV-HUMGAPS-449`'s "final two unchanged" for this class, save-time heap
    residue passed through raw archive I/O rather than an authored value.
- Evidence files: `evidence/functions.txt`, `evidence/refs.txt`,
  `evidence/game0018-span.tsv`. Also cites `SAV-1031`, `SAV-1032`,
  `SAV-CLASSSER-173`, `SAV-CLASSSER-177`, `SAV-HUMGAPS-449`, `MAGIC-DMG-005`.

**Confidence.** Medium for the claim as a whole, because it reports the
mechanism and the single corpus instance together as one finding. The
raw-transfer mechanism alone, read at instruction level, is High. The corpus
half is n=1 over `SAV-CLASSSER-177`'s population: one witnessed record agreeing
with a prediction is not a census, and a second saved Effect_DirectDamage could
show a different residue pattern at `+0x16`/`+0x17` without contradicting
anything here.

**Unknown.** `R1639`, the fourth owner of the shared helper, was not
traced and is not claimed to be unrelated.

## SpellTransport references: corpus witnesses, archive writes and post-load hooks

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1037 | Over 129 save paths / 76 SHA-distinct, more than double `SAV-CLASSSER-177`'s corpus, `SpellTransport`, `PointEffect` and `Effect_DirectDamage` keep one witness digest each; `AreaEffect` gains one. | Medium | ● active (amended) | [EXP-0378](../experiments/EXP-0378-transport-scheduling/) |
| SAV-1042 | At LOAD, `SpellTransport`'s `+0x44`/`+0x48` are bound to fresh `CArchive::ReadObject` references; its post-load hook rebinds Position and dispatches each child's hook, writing neither field. | High / Unknown | ● active | [EXP-0380](../experiments/EXP-0380-transport-rebind/) |
| SAV-1043 | Alternative 4, that no writer path emits a `SpellTransport` record, fails for the searched population: re-running `SAV-1037`'s census over its 129 paths / 76 digests reproduces its witness count. | High / Medium / Unknown | ● active (amended) | [EXP-0380](../experiments/EXP-0380-transport-rebind/) |
| SAV-1047 | `CArchive::WriteObject` writes `u16 0` for null, a compact back-reference for a known pointer, and on first occurrence the class then the object's vtable `+0x8`; objects and classes share one index space. | High / Unknown | ● active | [EXP-0381](../experiments/EXP-0381-transport-store/) |
| SAV-1048 | `SpellTransport::Serialize`'s STORE arm calls `CArchive::WriteObject` on `+0x44` and `+0x48` with no null guard; the base `Serialize` call runs once before the branch, and the class chain passes the container check. | High / Unknown | ● active | [EXP-0381](../experiments/EXP-0381-transport-store/) |
| SAV-1050 | Neither `PointEffect`'s nor `AreaEffect`'s post-load hook calls the shared-container registrar or append primitive; like `SpellTransport`'s (`SAV-1042`), each only forwards to its child's repair. | High | ● active | [EXP-0382](../experiments/EXP-0382-shared-effect-owner/) |

### SAV-1037

- The search is `SAV-CLASSSER-177`'s exact-schema1-record and raw-name search,
  `tools/savclassser`'s algorithm reproduced rather than imported in
  `experiments/EXP-0378-transport-scheduling/probe`, because that tool is
  `package main`.
- Population: `gameversions/en`, `gameversions/ru` and the full current
  `gameversions/saves` tree, 129 `Asg&`-decoding save paths, 76 SHA-distinct
  (`evidence/class-record-census.tsv`), against `SAV-CLASSSER-177`'s 55 paths /
  31 SHA-distinct. This is the whole scanned population, not a
  ROM1-organic-play subset.
- At least 12 of the 129 paths are project-produced fixtures rather than owner
  play, by their own documentation or naming:
  - `gameversions/saves/2026-08-27/EXP-0261-owner-runs/MANIFEST.md` labels
    `game9001.sav`/`game9002.sav`/`game9003.sav`/`game9004.sav`/`game9006.sav`
    "generated candidate";
  - one further directory under that same date and two under `2026-08-30`
    hold paths literally named `generated`, `generated-repeat`, `en-generated`
    or `ru-generated`, undocumented by any `MANIFEST.md` but self-identified by
    path: `2026-08-27/story-1073-owner-witness/{generated,generated-repeat}/game0000.sav`,
    `2026-08-30/SAV-work-3173785a/{en,ru}-generated/game0000.sav`,
    `2026-08-30/SAV-work-d7606aa1/{en,ru}-generated/game0000.sav`;
  - `2026-08-30/story-1073-original-acceptance/game0000.sav` shares the digest
    of that generated family, though its directory name does not say so.
  - None of the twelve carries an exact-schema1-record or raw-name hit in any of
    the four classes graded here (re-derived from
    `evidence/class-record-census.tsv`). This characterises a subset inside the
    129/76 figures, not an additional exclusion, and is not an exhaustive
    provenance census of the population.
- `SpellTransport`, `PointEffect` and `Effect_DirectDamage` each show 3 paths /
  1 distinct digest: the single digest `SAV-CLASSSER-177`, `SAV-EFFECTGRAPH-366`
  and `SAV-CASTCONT-1006` cite for `gameversions/en/game0018.sav`
  (`1e2eb21f…`). The other two paths are byte-identical archive copies of that
  file, not additional saves.
- `gameversions/saves/2026-08-27/EXP-0261-owner-runs/game9000.sav` is excluded
  from the 129/76 figures on one format-level ground: its header magic is
  `Bsg&`, not `Asg&`, so the census algorithm does not read it
  (`evidence/skipped-files.tsv`). That directory's `MANIFEST.md`, preserved
  beside the corpus under `gameversions/` and tracked by no repository,
  separately records it as "generated candidate (EXP-0257)", consistent with
  the twelve above but not the ground of the exclusion.
- `AreaEffect`, with zero witnesses in `SAV-CLASSSER-177`'s corpus, shows one:
  `gameversions/saves/2027-09-07/game0125.sav` (digest `3a055c8d…`).
- `pipeline/LOG.md:3842` names that file individually: label "5m city A saved",
  30,826 bytes, counters 1791/112, 37 actors, verified
  `savtool verify identical`, descending from "rung 1", which the same entry
  defines as "the original's own bytes with the purse injected". The file is
  ROM1's own output, itemized by the log rather than documented only at
  directory level, but written from a document this project modified before
  ROM1 loaded it, not a plain unmodified owner-play session.
- Evidence files: `evidence/class-record-census.tsv`,
  `evidence/class-record-summary.tsv`, `evidence/skipped-files.tsv`,
  `evidence/input-manifest.json`. Also cites `SAV-CLASSSER-177`,
  `SAV-EFFECTGRAPH-366`, `SAV-CASTCONT-1006`.

**Confidence.** Medium, the grade `SAV-CLASSSER-177` carries, for the same
reason: a bounded corpus search that discriminates the
no-second-`SpellTransport`-witness hypothesis (it could have found one and did
not) but cannot establish future reach. The graded finding is the
`SpellTransport`/`PointEffect`/`Effect_DirectDamage` count, over the full
scanned population including its project-produced fixtures, none of which
contributes a hit. The `AreaEffect` witness is a corpus fact for a future
experiment and is not graded: neither its fields nor its graph position were
decoded, and its source document carries a project-injected value upstream of
the ROM1 write that produced it.

**Unknown.** The `AreaEffect` record in `game0125.sav` was not decoded field by
field, and its graph position was not determined: a direct cast, or a nested
reference reachable through `SpellTransport+0x48`, which `formats/sav/objects.md`
documents as an `AreaEffect` reference. Whether that file's provenance is
admissible as corpus evidence on the same footing as an untouched owner save is
not resolved.

**Amended.** The fixture-directory count is corrected
([`retracted.md`](retracted.md)). The card said "two further directories under
that same date and one under `2026-08-30`". The paths it lists, and the census
rows in `evidence/class-record-census.tsv`, are one directory under
`2026-08-27` (`story-1073-owner-witness`) and two under `2026-08-30`
(`SAV-work-3173785a`, `SAV-work-d7606aa1`). The 12-fixture total, the 129/76
figures and the witness counts are unchanged.

### SAV-1042

- `SpellTransport::Serialize` (`R1602`) branches its `IsLoading` arm at
  `L08901`/`L08902`. The LOAD arm (`L08902`-`L08903`) computes
  `this+0x44` as an address (`L08904` adds 0x44 to the object pointer) and calls
  `R1124(ar, &this->+0x44)`, and `this+0x48` the same way
  (`L08905` adds 0x48 to the object pointer, `L08906(ar, &this->+0x48)`).
- Both wrappers share one 10-instruction shape: push a literal class
  descriptor (`L05929` for `+0x44`, `L08907` for `+0x48`), call
  `CArchive::ReadObject` (`L05928`; `SAV-774`: "class-checks every typed
  reference it resolves") with the archive as `this`, then store the returned pointer
  directly into the field address the caller supplied.
- `L05929` is the descriptor the container-level load loop (`MAGIC-202`)
  pushes for each element it resolves generically: one typed-reference
  mechanism serves the container's per-element resolution and this class's two
  named fields.
- The post-load hook (`R1115`, `SAV-TOKENLOAD-093`'s slot-`0x24` target for
  this class) first calls the common Position-rebind hook
  (the call to `R1560` at `L08908`). Then, for each of `+0x44`/`+0x48`
  independently, it null-guards the field (`L08909`/`L08910`) and, only if
  non-null, reads it and dispatches the child's vtable `+0x24` slot
  (`L08911`/`L08912`). Its complete 27-instruction body contains no store
  to either field.
- The "Position/child repair calls" `SAV-CASTCONT-1006` names, and the
  canonical page's "repairs Position and child references", are therefore
  Position on the transport itself plus a forwarded call letting each child
  repair itself. The transport's `+0x44`/`+0x48` are set once, in `Serialize`'s
  LOAD arm, and never rewritten afterward on this path.
- Evidence files: `evidence/transport-serialize-R1602-body.txt`,
  `evidence/transport-postload-R1115-body.txt`,
  `evidence/readobject-wrapper-R1124-body.txt`,
  `evidence/readobject-wrapper-L08906-body.txt`. Also cites
  `SAV-CASTCONT-1006`, `SAV-774`, `SAV-TOKENLOAD-093`, `MAGIC-202`.

**Confidence.** High for both facts, each read as a complete instruction
sequence and cross-checked EN/RU (`R1602`, `R1115`, `R1124`,
`L08906` all byte-identical between images). This adds to
`SAV-CASTCONT-1006` and does not correct it: that claim's wording ("performs
Position/child repair calls") already reads as dispatch. New is that the hook
writes neither field, which bears on that claim's declared "transitive repair …
Unknown". No retraction is owed.

**Unknown.** What either child's own `+0x24` slot does with the dispatch. The
experiment stops at the moment of binding, per its stopping rule.

### SAV-1043

- The unmodified probe `experiments/EXP-0378-transport-scheduling/probe`
  (`SAV-1037`'s tool), run over `gameversions/en`, `gameversions/ru` and the
  current `gameversions/saves` tree, scans 129 `Asg&`-decoding save paths, 76
  SHA-distinct: the same population `SAV-1037` reports, not a larger one.
- `SpellTransport`: 3 file-path occurrences, 1 distinct digest, 3
  exact-schema1-record hits, 3 raw-name hits
  (`evidence/class-record-summary.tsv`), matching `SAV-1037` exactly.
- The one witnessed record is in `gameversions/en/game0018.sav`,
  `SAV-CASTCONT-1006`'s natural record. `MAGIC-202` reads the mechanism that
  turns such a record into a live object: the container's per-element LOAD loop
  resolves each element it holds through descriptor `L05929` and
  `CArchive::ReadObject`, the same wrapper and descriptor `SAV-1042` reads for
  `SpellTransport`'s `+0x44` child. That the witnessed record is one of that
  loop's elements is not read here.
- Evidence files: `evidence/class-record-census.tsv`,
  `evidence/class-record-summary.tsv`, `evidence/skipped-files.tsv`,
  `evidence/input-manifest.json`, `evidence/program-digest.tsv`. Also cites
  `SAV-1037`, `MAGIC-202`, `SAV-CASTCONT-1006`.

**Confidence.** Medium, the grade `SAV-1037` carries, for the same reason: a
bounded, reproducible corpus search that discriminates "no second witness
exists" from "one was found", not a claim about files outside the searched
population. This claim adds no file to that population; it reproduces the
existing count independently as the experiment's test of Alternative 4 rather
than assuming `SAV-1037`'s result unchanged. High only for the descriptor
identity read: the container's per-element load loop and
`SpellTransport::Serialize`'s `+0x44` push the same descriptor `L05929` into
the same `CArchive::ReadObject` wrapper.

**Unknown.** Whether the witnessed record is itself an element of that
container. Neither this experiment nor any claim it cites reads the class
descriptor at `L05929`, `SpellTransport`'s runtime-class chain, or its
base-class `Serialize` call at `L08913`, so that a `SpellTransport` passes
that loop's class check is inference here, not a read fact.

**Amended.** `SAV-1048` reads that chain and that call site: the base-class
`Serialize` call is at `L08914`, not the approximate `L08913` cited here.
The class-identity half of this Unknown is settled there (base chain
`L08915`→`L05929`, one hop). That corrects what this claim could conclude
and is not a retraction: `L08913` was cited as an unread call, and the exact
call address is `L08914`.

### SAV-1047

- `R1129` takes the archive as `this` and one stack argument `pOb`. `SAV-982` names
  `R1129` as the target of `R1110` (`R1110`), and its first-object
  branch, in one clause without reading the body.
- Null: `pOb` is tested first (`L08916`-`L08917`). A null writes a bare
  `u16` value of `0` through `L08299` (`SAV-PLAYER-028`'s `u16` writer),
  with no class information at all (`L08918`-`L08919`).
- Known pointer: `L08920`-`L08921` looks `pOb` up in the map at `ar+0x34`
  through `R1640`, a shared index map keyed by object pointer here and by
  class-descriptor pointer inside `CArchive::WriteClass`. The slot's contents
  (loaded as one dword) are `0` for a pointer never indexed (a branch to `L08922` when equal, the
  new-object branch), else the index already assigned to this pointer, written
  back as a compact reference with no bit set: a plain `u16` for an index below
  `0x7fff` (`L08923`-`L08924`), or `u16 0x7fff` followed by a raw `u32`
  (`L08925`-`L08926`, through `L08300`, `SAV-PLAYER-028`'s `u32`
  writer).
- `WriteClass`'s identical branch pair (`L08927`, `L08928`) sets bit
  `0x8000` or `0x80000000` into the same value before the same two writer calls.
  That is the one instruction-level difference between an object index and a
  class index on the wire, and the discriminator `CArchive::ReadClass`'s
  normalization (`L08929`-`L08930`, read by `SAV-774`) tests uniformly at
  bit 31.
- First occurrence (`L08922`-`L08931`): `WriteObject` dispatches vtable
  slot `0` (`GetRuntimeClass`, `L08932`), calls `CArchive::WriteClass`
  (`L08933`) with the result (`L08934`), calls `L08935`, looks `pOb`
  up a second time (`L08936`-`L08937`), registers it, then dispatches the
  object's `Serialize` through vtable slot `+0x8` with `ar` as its one argument
  (`L08938`-`L08939`).
  - Registration reads the running counter `ar+0x30` at `L08940`, stores it
    into the map slot at `L08941` and post-increments it at `L08942`;
    `L08943`, between the read and the store, is a push staging `ar` for
    the dispatch. `SAV-946` reads the same shape on the LOAD side at
    `L08079`/`L08080` ("calling that receiver+8 at `L08080`"), the slot
    `CArchive::ReadObject`'s new-object arm dispatches.
  - `L08935` is called at the identical point by `WriteObject` and
    `WriteClass`, and by the shared lazy-init call `L08944` before its own
    registration steps; it was read only to confirm this shared call shape.
- `WriteClass` asserts the class's schema word (descriptor `+0x8`) is not
  `0xffff` (`L08945`-`L08946`) and performs the same map lookup keyed by the
  descriptor. On first occurrence it writes `u16 0xffff` (`L08947`-`L08948`)
  and calls `R1641` (the class descriptor as `this`, one stack argument `ar`,
  `L08949`-`L08950`), which writes the schema word (`descriptor+0x8`) as a
  `u16` (`L08951`-`L08952`), the class name's length as a `u16`
  (`L08953`-`L08954`), then the name bytes through `R0271`
  (`L08955`), `SAV-PLAYER-028`'s length-prefixed raw writer.
- It then registers the descriptor in the same index structure, not a separate
  one: the same counter `ar+0x30` (`L08956`), the map at `ar+0x34`
  (`L08957`), the post-increment (`L08958`), the shape of `WriteObject`'s
  `L08940`/`L08941`/`L08942`. On LOAD, `ReadClass` (`L08959`-`L08960`)
  and `ReadObject` (`L08961`-`L08079`) read and advance `ar+0x30` the same
  way for classes and objects.
- Objects and classes therefore share one index space. A class is written
  inline (schema word plus name bytes) only on its first registration, by either
  path, anywhere in the archive. Every later reference to it, from any object,
  is the compact `0x8000 | index` (or `0x7fff`-escaped `0x80000000 | index`)
  back-reference, never a second name write.
- `SAV-774` names `L08682` among "the only unchecked returns in either
  function" without explaining why that return is always safe. The lazy-init
  call `L08944` sets `ar+0x30` to `1` on its first-touch allocation branch,
  reached only while `[esi+0x34]` is still null (`L08962`, and independently
  `L08963` for the load-side allocation shape), and never resets it on a later
  call. A null write (`L08299` with argument `0`) is indistinguishable on the
  wire from a small-index reference of value `0`, and index `0` is a load-array
  slot that no registration in the five bodies read whole (`R1129`,
  `L08933`, `L05928`, `L08656`, `L08944`), on either side, assigns a
  real pointer into. It is zero-filled once at allocation
  (`L08964`-`L08965`) and reads back as `0` on every path traced.
- Every `WriteObject` call writes something: no branch skips a child, and a
  null child is an explicit, first-checked, symmetrically encoded case rather
  than an omission.
- Prior art: `go run ./tools/claim -k` on `R1129` returns
  `SAV-982`/`SAV-CITYSTORE-516`/`SAV-PRODPOP-205`, none of which reads the body
  of `R1129`; `L08933`, `R1641` and `R1640` (`L08933`, `R1641`,
  `R1640`) each return no row. `L08944` is cited by no earlier claim.
- Evidence files: `evidence/writeobject-core-R1129-body.txt`,
  `evidence/writeclass-L08933-body.txt`,
  `evidence/writeclassname-schema-R1641-body.txt`,
  `evidence/objectclass-maplookup-R1640-body.txt`,
  `evidence/lazyinit-L08944-body.txt`, `evidence/readobject-L05928-body.txt`,
  `evidence/readclass-L08656-body.txt`, `evidence/unresolved.json`,
  `evidence/program-digest.tsv`. Also cites `SAV-982`, `SAV-774`, `SAV-907`,
  `SAV-984`, `SAV-946`, `SAV-PLAYER-028`.

**Confidence.** High for the four write-side bodies read whole (`R1129`,
`L08933`, `R1641`, `R1640`; 52/54/21/33 instructions, zero unresolved
per `evidence/unresolved.json`), and for the bit encoding: the `OR` setting
`0x8000`/`0x80000000` is present in `WriteClass` and absent in `WriteObject`,
read, not inferred. EN/RU agreement is asserted by `disasm.py`'s
`_assert_equal` check, not read from `evidence/unresolved.json` (a
per-function, EN-only unresolved-branch list); `evidence/program-digest.tsv`
records the identical SHA-256 for both preserved roots, so the EN/RU comparison
adds no confirming power beyond the single image read. Beyond `SAV-982`'s one
clause ("first-object branch queries slot00 and calls slot08"), this is the
first complete read of the function: the null branch, the back-reference branch
and `WriteClass`'s schema and name body. The decode side is published at High
and cited, not re-derived: `SAV-774` reads `ReadClass`'s bit-31 normalization;
`SAV-907`, `SAV-984` and `SAV-946` read `ReadObject`'s back-reference arm
("returns the existing object without rerunning its factory/serializer",
`SAV-907`) and its new-object registration ("registers the created receiver...
before calling that receiver+8", `SAV-946`). The index-0 argument is a bounded
finding, not a census: no other code that writes `ar+0x30`/`ar+0x34`, or calls
`L08966`, was searched for, so a writer elsewhere in the binary or a
mid-archive reset is not excluded.

**Unknown.**

- Whether index `0` reading back as `0` is guaranteed by the allocator itself
  (`L08967`, not read) rather than observed across the paths traced.
- Whether the index structure at `ar+0x34` is exactly a `CMapPtrToPtr` when
  storing and a growable pointer array when loading, as the two differently
  sized lazy allocations (`0x1c`/`0x14` bytes) inside `L08944` suggest, and
  what `L08935` does. Neither was read past the point needed to locate the
  counter's starting value and the shared call shape.
- The schema-mismatch error path (`L08968`-`L08969`) is read as present but
  not exercised.

### SAV-1048

- `R1602`'s prologue calls `R1600` (`SpellEffect::Serialize`,
  which `SAV-CLASSSER-172`/`SAV-CLASSSER-173` name as this class's base) at
  `L08914`. That is eight bytes past `SAV-1043`'s approximate `L08913`,
  which falls inside the preceding 32-bit store to the frame local at -4 and names no
  instruction; three instruction starts separate the two addresses
  (`L08970`, `L08971`, `L08972`), then `L08914` itself. No retraction is
  owed: `SAV-1043` named that address as unread, not as a specific instruction.
- The call precedes the `IsLoading` test (`L08973`-`L08901`) entirely, so
  it is one shared call reached by both arms, not a separate call inside each.
- The STORE arm (`L08974`-`L08975`) reads `this+0x44` and `this+0x48`
  directly (`L08976`, `L08977`), with no test against zero anywhere in the
  arm, and hands each to `R1110` (`SAV-1047`'s wrapper for
  `CArchive::WriteObject`) before writing the `u16` counter at `+0x4c` through
  `R1583` (`L08978`-`L08979`).
- That field order and typing is `SAV-CLASSSER-175`'s, at High ("references
  `*(+0x44)` and `*(+0x48)`, then `u16 +0x4c`"). This claim adds the call target
  (`R1110`/`R1129`) and the absence of a null guard: whichever of
  `+0x44`/`+0x48` is null at STORE time, `WriteObject` still runs and still
  writes something (`SAV-1047`'s null branch).
- `SpellTransport`'s `CRuntimeClass` descriptor (`L08915`, reached from its
  `GetRuntimeClass` thunk `R1642`), read as bytes, gives base pointer
  `L05929`, whose name resolves to `SpellEffect`: the descriptor `MAGIC-202`'s
  container-level per-element load loop pushes into `CArchive::ReadObject` for
  every element it resolves. `ReadObject`'s class check (`SAV-774`'s
  `R0216` ancestor walk, which "accepts any ancestor match" by walking the
  descriptor `+0x10` base chain to a match or to null) therefore accepts a
  `SpellTransport` on its first hop. This settles, by class identity,
  `SAV-1043`'s Unknown: "that a `SpellTransport` passes that loop's own class
  check is inference, not a read fact".
- Cross-checks by other methods, not new findings:
  - `SAV-CLASSSER-172`/`SAV-CLASSSER-175` establish the same hierarchy
    ("`SpellTransport`... is `SpellEffect`") by tracing the base-class
    `Serialize` call; the base-pointer field walk here agrees.
  - The vtable slot `+0x8` (raw dword at `L05917+0x8`,
    `evidence/vtable-slots.json`) is `R1602`, the function read here.
    `SAV-CLASSSER-172` publishes the same fact ("their slot-2 targets are
    respectively... `SpellTransport` `R1602`") from creator and constructor
    tracing.
  - `EXP-0362`'s dynamic capture (`measure.py`'s `trial()`, cited by
    `SAV-CASTCONT-1006`, which stubs `R1110`/`R1583` and runs the real
    STORE arm under emulation) observed the same call sequence and order; it
    was not rerun here.
- Evidence files: `evidence/transport-serialize-R1602-body.txt`,
  `evidence/vtable-slots.json`, `evidence/descriptor-chains.json`,
  `evidence/writeobject-wrapper-R1110-body.txt`. Also cites `SAV-1042`,
  `SAV-1043`, `SAV-1047`, `SAV-CASTCONT-1006`, `SAV-CLASSSER-172`,
  `SAV-CLASSSER-175`, `MAGIC-209`, `MAGIC-202`, `SAV-774`.

**Confidence.** High for the STORE arm's complete instruction sequence (50
instructions across both arms of `R1602`, zero unresolved, EN/RU-identical),
the corrected base-call-site address and its unconditional, shared placement,
and the descriptor-chain read settling class-identity eligibility. The EN and
RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement on this
code read is not independent corroboration. Against the brief's four
alternatives, over the population traced (not a corpus-wide census;
`SAV-1043`'s census is cited unchanged, not repeated):

- B3 is excluded for this path: the STORE arm unconditionally calls
  `WriteObject` on both fields, so nothing is written "from elsewhere" at LOAD
  because nothing is withheld at STORE.
- A separate resolution pass, as B2 states it, is excluded for this path:
  `SAV-1047`'s back-reference index is resolved synchronously, by the same map
  lookup, inside the one linear `WriteObject`-driven traversal. No second pass
  over an already-written object is read in any body this experiment or
  `SAV-1047` traces.
- B1 holds conditionally, not unconditionally: `WriteObject` writes a child's
  class and fields inline on the first occurrence of that object pointer
  anywhere in the archive, and a compact index on every later occurrence.
- B4 is not re-tested by a fresh file census: `SAV-1043` ran that census over
  the current corpus, and an unchanged population would reproduce it. This
  claim's test of B4 is the vtable-dispatch and descriptor-chain reads, either
  of which could have failed (a different function at that slot; a base chain
  not reaching `SpellEffect`) and did not.

**Unknown.**

- Whether either child field is that first occurrence at STORE time. The
  container's STORE arm (`MAGIC-209`), for example, may enumerate the same
  object independently as a top-level element before or after the transport's
  turn. `SAV-CLASSSER-177`'s corpus finding (three nested class bodies at
  distinct byte ranges in the one witnessed record) is consistent with all
  three being first occurrences in that file, but it is not a reading of the
  dedup map's state and does not settle this.
- Whether the one witnessed corpus record (`SAV-1043`, `SAV-CASTCONT-1006`) is
  reached through that container: the separate Unknown those claims carry.

### SAV-1050

- `PointEffect`'s post-load hook (vtable `+0x24`, `L05904`, 18
  instructions, `evidence/pointeffect-postload-L05904-body.txt`) and
  `AreaEffect`'s (`R1116`, 18 instructions,
  `evidence/areaeffect-postload-R1116-body.txt`) were each read whole,
  EN/RU-identical, with zero unresolved branches.
- Both first call the common hook `R1560` (the Position-rebind hook
  `SAV-1042` names), then null-guard their class-specific field and, only when
  non-null, dispatch that field's vtable `+0x24` slot once. The field is
  `PointEffect`'s `+0x48` and `AreaEffect`'s `+0x44`, the "inner Effect" field
  `MAGIC-213` reads, not the shared-container sibling field `SpellTransport`'s
  post-load hook touches. The shape is the one `SAV-1042` reads for
  `SpellTransport`'s hook, applied to a different field.
- Neither body contains an instruction referencing `R1113` (the registrar
  wrapper `MAGIC-187`/`MAGIC-197` read) or `R0647` (the generic append
  primitive `MAGIC-187` reads). This is a within-body negative over each hook's
  complete 18-instruction disassembly, not a fresh image-wide census.
  `MAGIC-187`'s image-wide census of `R1113`'s callers (4 sites, 2 owners,
  neither `L05904` nor `R1116`) already excludes either hook by full
  enumeration.
- For the two `SpellEffect`-lineage classes `SAV-1042` did not cover, this
  answers sub-question 4 of the preregistration: no post-load hook in this
  family disambiguates, adjusts or re-registers ownership of a
  doubly-referenced child; each forwards a repair call to whatever it already
  holds.
- With `SAV-1042` (`SpellTransport`) and `MAGIC-202` (the container-level LOAD
  arm, which resolves membership once, through the archive's typed-reference
  mechanism, before any post-load hook runs, per `SAV-TOKENLOAD-094`'s Phase-2
  ordering), the post-load picture across all three classes is one shape:
  LOAD-time field binding and LOAD-time container membership are each set once,
  and no post-load hook this experiment or `SAV-1042` reads revisits which
  object owns which.
- Evidence files: `evidence/pointeffect-postload-L05904-body.txt`,
  `evidence/areaeffect-postload-R1116-body.txt`, `evidence/vtable-slots.json`,
  `evidence/unresolved.json`. Also cites `SAV-1042`, `MAGIC-187`, `MAGIC-202`,
  `SAV-TOKENLOAD-094`, `MAGIC-213`.

**Confidence.** High for both hook bodies as complete instruction sequences,
EN/RU-identical, and for the absence of any registrar or append reference within
them: the basis `SAV-1042` uses for `SpellTransport`'s hook. This closes the
experiment's preregistered scope item ("`PointEffect`'s and `AreaEffect`'s own
post-load hooks... for whether either calls the registrar or the append
primitive"). Scope: this is not a claim that no other code path re-registers a
child post-load, only that these two 18-instruction hooks do not.
`MAGIC-202`'s broader finding, that `R1126`, the one insertion pass
proven to run after the archive completes, "performs no insertion...
dispatches, it does not register", is consistent with this and is cited, not
re-derived.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

## SpellEffect caster `+0x3c` through Serialize and LOAD construction

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1054 | `Token::Serialize` (`R0950`), the base every concrete `Serialize` in the `SpellEffect` lineage calls first, never touches `+0x3c` in its STORE or LOAD arm. | High | ● active | [EXP-0383](../experiments/EXP-0383-effect-token-reference/) |
| SAV-1055 | The LOAD factory `SAV-1011` left open is `CArchive::ReadObject` (`L05928`) calling the class descriptor's `+0xC` creator; for all four `SpellEffect`-lineage classes it reaches `L08020`, which zero-writes `+0x3c`. | High | ● active | [EXP-0383](../experiments/EXP-0383-effect-token-reference/) |
| SAV-1056 | A raw `E8 rel32` census finds 5 direct callers of `L08020` and 3 of `L05979`, all construction paths, one previously uncited; no new `+0x3c` write, and none in the two attribution tails. | High / Medium / Unknown | ● active | [EXP-0383](../experiments/EXP-0383-effect-token-reference/) |

### SAV-1054

- `R0950` is 121 instructions, read whole, with zero unresolved
  branches, EN/RU-identical (`evidence/token-serialize-R0950-body.txt`). An
  `IsLoading` test at `L08980`/`L08981` splits STORE
  (`L08982`-`L08983`) from LOAD (`L08984`-`L08985`).
- STORE writes to the wire from `this`'s `+0x4` (dword, `L08986`), `+0xc`
  (byte, `L08987`), `+0xe` (word, `L08988`), `+0x8` (dword, `L08989`),
  `+0x18` (word, `L08990`), `+0x1c` (dword, `L08991`), an unconditional
  dword at `L08992` (the base extent itself, no `this` offset add before the
  call) and `+0x14` (dword, `L08993`).
- A call at `L08994` forwards to `R0530` with no offset add: the
  base-of-all-bases call every path takes first, read whole, one instruction,
  `ret 4` (`evidence/token-basector-R0530-body.txt`). One to `R1494`
  (`L08995`) dispatches through `this->+0x10` as a receiver, not a raw field
  write.
- LOAD mirrors the same offset set in the same order (`+0x4` at `L08996`,
  `+0xc` at `L08997`, `+0xe` at `L08998`, `+0x8` at `L08999`, `+0x18` at
  `L08525`, `+0x1c` at `L08526`). It then reads two further raw dwords into
  a shared local slot, around a call (`L08312`) that registers `this` under
  the first loaded dword; it writes the second into `this->+0x14` (`L09000`)
  and passes `&this->+0x14` to one further callee (`L08196`), unread past
  confirming it receives that address.
- No instruction in either arm adds `0x3c` to a base register or dereferences an
  operand at that displacement; a grep of the disassembly text finds zero
  matches.
- `SAV-CLASSSER-173`/`-174`/`-175` read the four concrete `Serialize` bodies
  (`SpellEffect::Serialize`, `PointEffect::Serialize`, `AreaEffect::Serialize`,
  `SpellTransport::Serialize`) as touching no `+0x3c` either. This claim covers
  the shared base every one of them calls first, the one level `SAV-1011`'s
  three concrete-`Serialize` reads did not check.
- `SAV-653` (promoted, `EXP-0312`) reads this body's STORE and LOAD arms for the
  fields it touches and names the same offsets, addresses and helper pair:
  `+0x18` u16 via `L08299`/`R0282`, `+0x1c` u32 via `R0685`/`R0686`,
  `+0xe` u16 via the same u16 pair, `this`/`+0x4`/`+0x8`/`+0x14` u32 via
  `R0685`/`R0686`. That corroborates the offset set; the absence finding
  does not depend on it and rests on the exhaustive read and the grep, not on
  that claim's narrower field-width scope.
- Evidence files: `evidence/token-serialize-R0950-body.txt`,
  `evidence/token-basector-R0530-body.txt`, `evidence/unresolved.json`,
  `evidence/program-digest.tsv`. Also cites `SAV-1011`, `SAV-CLASSSER-173`,
  `SAV-CLASSSER-174`, `SAV-CLASSSER-175`, `SAV-653`.

**Confidence.** High for what the base and its two arms touch: a complete,
zero-unresolved-branch instruction-level read, cross-checked by direct text
search and, for the offsets that are written, by `SAV-653`'s independent full
read, which names the same offsets except `+0xc`. With `SAV-1011`'s absence at
the four leaf `Serialize` bodies, this rules out A1, A2 and A4 on their shared
serialization premise: no level of the inheritance chain, base or leaf, writes
`+0x3c` to or reads it from an archive, so no resolvable reference for LOAD to
restore exists anywhere in the traced `Serialize` call graph, and no consumer
can be reading a value `Serialize` ever put there. This supports A3's
serialization clause (not serialized at all).

**Unknown.** A3's further clause, that a non-zero on-disk value at this offset
would mean wrong offset attribution, is untested: no `.sav` corpus byte was
read, per the experiment's excluded sources.

### SAV-1055

- `CArchive::ReadObject` (`L05928`, 71 instructions, whole, zero unresolved,
  `evidence/archive-readobject-L05928-body.txt`): its new-object arm
  (`L08657 jne L09001`) calls `R1643(descriptor)` at `L09002`, with
  the class descriptor as `this`, resolved by `L08656` (re-read at the call site
  only, not re-disassembled). It registers the returned object in the archive's
  index map (`this+0x30`/`this+0x34`, `L08961`-`L09003`), then dispatches
  the object's vtable `+8` slot (`Serialize`) at `L08080`.
- `R1643` (24 instructions, whole, zero unresolved,
  `evidence/descriptor-createobject-R1643-body.txt`): a comparison of the descriptor's `+0xc` field with zero and
  a branch to `L09004` when equal skips construction when the descriptor's `+0xc` field is null;
  otherwise an indirect call at `L09005` through that field, not a
  literal address.
- `SpellEffect`'s `CRuntimeClass` descriptor (`L05929`), read as raw dwords
  rather than disassembly (`evidence/descriptor-dwords.json`):
  `+0x0=L09006` (name), `+0x4=0x44` (object size, 68), `+0x8=1` (schema),
  `+0xc=L09007`, `+0x10=L08311` (base descriptor), `+0x14=0`. `+0xc`
  confirms `class-lifecycle-map.txt`'s identification of the "create" field from
  the raw layout rather than from that table's column order.
- The other three descriptors, read the same way (`PointEffect` `L09008`,
  `AreaEffect` `L08907`, `SpellTransport` `L08915`, same evidence file):
  each `+0xc` matches the creator address read below (`L09009`,
  `L09010`, `L09011`), and each `+0x10` points to `SpellEffect`'s
  descriptor as the base. This confirms `class-lifecycle-map.txt`'s
  class/creator-field pairing for all four classes from the raw layout.
- Each creator, read to its end, whole, zero unresolved:
  - `SpellEffect` (`L09007`, 30 instructions) allocates `0x44` bytes via
    `R1147` and, on success, calls `L08020` with the object as `this` directly
    (`L09012`). `SpellEffect` is the lineage's base, so its creator reaches
    `L08020` with no further wrapper.
  - `PointEffect` (`L09009`, 30 instructions) allocates `0x4c` bytes and
    calls `R1629` (`L09013`). `R1629` (16 instructions,
    `evidence/pointeffect-zerotarget-ctor-R1629-body.txt`) calls `L08020`
    (`L09014`), installs vtable `L05356`, and zero-writes `+0x48` then
    `+0x44`. This confirms `SAV-1011`'s named structural candidate from the
    dispatch itself, not from shape alone.
  - `AreaEffect` (`L09010`, 30 instructions) allocates `0x50` bytes and calls
    `R1644` (`L09015`). `R1644` (18 instructions,
    `evidence/areaeffect-zerotarget-ctor-R1644-body.txt`, named by no earlier
    claim) calls `L08020` (`L09016`), installs vtable `L05352`
    (`AreaEffect`'s, per `MAGIC-213`), and zero-writes `+0x44`, `+8` and the byte
    `+0x48`.
  - `SpellTransport` (`L09011`, 30 instructions) allocates `0x50` bytes and
    calls `R1645` (`L09017`). `R1645` (16 instructions,
    `evidence/spelltransport-zerotarget-ctor-R1645-body.txt`, named by no
    earlier claim) calls `L08020` (`L09018`), installs vtable `L05917`
    (`SpellTransport`'s, per `MAGIC-214`/`MAGIC-215`), and zero-writes `+0x44`
    then `+0x48`.
- Every archive-driven LOAD path of the four classes (the creator field, then,
  for the three derived classes, one class-specific zero-target constructor)
  converges on base constructor `L08020`, which `SAV-1011` reads zero-writing
  `+0x3c` at `L08772`.
- `SAV-1011`'s second base constructor, `L05979`, is reached by none of the
  four creator-field chains. Per `SAV-1011`'s citations it is reached only by
  `AreaEffect`'s and `SpellTransport`'s cast-time constructors (`R0652` at
  `L08773`, `R0633` at `L05978`); `SAV-1056` adds a third, previously
  uncited caller.
- Sub-question 2: nothing repairs `+0x3c` on LOAD. `PointEffect`'s cast-time
  constructor also uses `L08020` (`L08771`, named by `SAV-1011`, reproduced
  by `SAV-1056`'s census). Both base constructors zero-write `+0x3c`
  (`L08772`, `L08774`), and `SAV-1054` shows `Serialize` never touches the
  field afterward.
- Post-load hooks: `PointEffect`'s and `AreaEffect`'s are read as not touching
  `+0x3c`
  (`experiments/EXP-0382-shared-effect-owner/evidence/pointeffect-postload-L05904-body.txt`,
  `.../areaeffect-postload-R1116-body.txt`, cited, not re-derived).
  `SpellTransport`'s (`R1115`, `SAV-1042`, re-read here,
  `evidence/spelltransport-postload-R1115-body.txt`) has no `+0x3c`
  displacement in its 27-instruction body either.
- Evidence files: `evidence/archive-readobject-L05928-body.txt`,
  `evidence/descriptor-createobject-R1643-body.txt`,
  `evidence/descriptor-dwords.json`,
  `evidence/spelleffect-creator-L09007-body.txt`,
  `evidence/pointeffect-creator-L09009-body.txt`,
  `evidence/pointeffect-zerotarget-ctor-R1629-body.txt`,
  `evidence/areaeffect-creator-L09010-body.txt`,
  `evidence/areaeffect-zerotarget-ctor-R1644-body.txt`,
  `evidence/spelltransport-creator-L09011-body.txt`,
  `evidence/spelltransport-zerotarget-ctor-R1645-body.txt`,
  `evidence/spelltransport-postload-R1115-body.txt`,
  `evidence/unresolved.json`, `evidence/program-digest.tsv`. Also cites
  `SAV-1011`, `SAV-1042`, `SAV-1054`, `SAV-1056`, `SAV-TOKENLOAD-093`,
  `MAGIC-213`, `MAGIC-214`, `MAGIC-215`.

**Confidence.** High for the complete dispatch chain from `ReadObject` to each
lineage class's base-constructor call: every link read whole with zero
unresolved branches and confirmed against a raw descriptor-dword read rather
than an existing table's column order. This closes `SAV-1011`'s Unknown ("this
experiment did not identify the archive's own create/factory hook") for all
four classes, not only the one `R1629` was a candidate for. It excludes
the live alternative A1, that some unlocated LOAD-path code restores a
resolvable reference, for the mechanism traced: the archive's object
construction, from `ReadObject` down to the base constructor, is a closed loop
with no branch left unaccounted for reaching any other target. It does not
exclude a restorer entirely outside this call chain, such as a later post-load
hook; the three post-load hooks above are the ones read.

### SAV-1056

- Instrument: `call_sites_to()` in
  `experiments/EXP-0383-effect-token-reference/disasm.py`, the raw `E8 rel32`
  scan with the documented blind spot that `EXP-0381`/`EXP-0382` use,
  reimplemented keyed to one fixed target (`evidence/call-site-census.json`).
  `SAV-1011` itself names 1 and 2 callers of the two base constructors.
- `L08020`, 5 sites, all accounted for, one already published and four
  confirmed by `SAV-1055`'s four archive creator-field chains:
  - `L09012`, `SpellEffect`'s creator field, direct;
  - `L09014`, in `R1629`, `PointEffect`'s LOAD-path constructor;
  - `L08771`, `PointEffect`'s cast-time constructor `R1046`, named by
    `SAV-1011`;
  - `L09016`, in `R1644`, `AreaEffect`'s LOAD-path constructor;
  - `L09018`, in `R1645`, `SpellTransport`'s LOAD-path constructor.
- `L05979`, 3 sites:
  - `L08773`, `AreaEffect`'s cast-time constructor `R0652`, named by
    `SAV-1011`;
  - `L05978`, `SpellTransport`'s cast-time constructor `R0633`, named by
    `SAV-1011` and `MAGIC-215`;
  - `L09019`, cited by no earlier claim (`go run ./tools/claim -k` on
    `R1646`, `R1646` and `L09019` each returns no row).
- `L09019` is inside `R1646` (41 instructions, whole, zero
  unresolved, `evidence/spelltransport-altctor-R1646-body.txt`), a third
  argument-taking constructor shape distinct from `MAGIC-215`'s `R0633`.
  It calls `L05979` (`L09019`) with one stack argument, installs vtable
  `L05917` (`SpellTransport`'s), then writes `+0x44 = 0` (`L09020`) and
  `+0x48 = [ebp+8]` (`L09021`): the mirror image of `R0633`, which
  writes `+0x44` from its argument and always zero-writes `+0x48`. It writes no
  `+0x3c` past the `L05979` call.
- The same function re-reads the just-zeroed `+0x44` at `L09022` and
  dereferences it at `L09023` (`[eax+0x10]`, with `eax` necessarily `0`,
  since nothing writes `+0x44` between the zero-write and the re-read). That is
  unrelated to `+0x3c`, out of scope and not characterized further.
- A whole-image reachability scan for `R1646` finds no static reference
  of any of three kinds: `E8` call, `E9` jump, or an unaligned raw
  little-endian dword equal to its entry address `R1646`
  (`evidence/reference-census.json`). With the null dereference above, this is
  consistent with dead code.
- Neither attribution tail read (`R0266`, `R0267`, the subjects
  of `MAGIC-217`) writes `this->+0x3c` anywhere in its body, confirmed by
  reading every instruction each body contains, not only its
  `0x3c`-displacement ones.
- `Spell::Apply`'s two writes (`L08775`, `L08776`) and the Tick driver's
  conditional clear (`R0641`, `MAGIC-AREATICK-036`) remain the only
  writers this citation graph names past construction time. No writer beyond
  these four sites (the two base constructors, `Spell::Apply`, the Tick driver)
  was found within the population this experiment and its cited claims have
  read; no new distinct write instruction was found.
- Evidence files: `evidence/call-site-census.json`,
  `evidence/spelltransport-altctor-R1646-body.txt`,
  `evidence/pointeffect-tick-attribution-R0266-body.txt`,
  `evidence/areaeffect-tick-attribution-R0267-body.txt`,
  `evidence/reference-census.json`, `evidence/unresolved.json`,
  `evidence/program-digest.tsv`. Also cites `SAV-1011`, `SAV-1055`,
  `MAGIC-215`, `MAGIC-216`, `MAGIC-AREATICK-036`, `MAGIC-217`.

**Confidence.** High for the census counts and the site list: a reproducible
direct-call scan, cross-checked against `SAV-1011`'s three published sites, all
three reproduced exactly before the technique is trusted for the two new sites.
Medium for the complete caller population of the two base constructors: a raw
`E8 rel32` scan cannot see an indirect or vtable-dispatched caller, the
documented blind spot `MAGIC-216` applies to a different target in this class
family. Every constructor call read anywhere in this family (eight, across
`SAV-1055` and this claim) is a direct `E8` call, so there is no positive
evidence of an indirect constructor caller; that is an absence within a small,
non-exhaustive sample, not an exclusion.

**Unknown.**

- Whether the alternate construction path `R1646` is ever reached: the
  reachability scan cannot see a computed or table-driven call.
- The complete writer population of `+0x3c`. This census covers direct callers
  of the two zero-writing base constructors, so it cannot see a writer reached
  any other way, and `Spell::Apply`'s two writes lie outside it and were not
  found by it. No image-wide write-site census of `+0x3c` was run; that
  population is not bounded by this claim.

## Token `+0x14` on the SpellEffect and Effect lineages

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1058 | Within the six SpellEffect- and Effect-lineage classes' construction, destruction and post-load bodies, `+0x14` is written only by Token's four constructors and by `Token::Serialize` LOAD. | High / Medium | ● active (partially retracted) | [EXP-0384](../experiments/EXP-0384-effect-token-field14/), [EXP-0388](../experiments/EXP-0388-effect-construction/) |
| SAV-1059 | Over 70 admitted lawful saves, nonzero `Token+0x14` occurs only on `Effect`: 18 of 192 `Effect` instances, all in one original save, each pointing at the `Human` whose body contains that `Effect`. | High / Medium | ● active | [EXP-0384](../experiments/EXP-0384-effect-token-field14/) |
| SAV-1060 | Two located readers of the six classes' own `+0x14` exist, `Token::Serialize` STORE and Token's copy constructor; the PointEffect and AreaEffect tick tails read `+0x14` only on other objects. | High / Medium / Unknown | ● active | [EXP-0384](../experiments/EXP-0384-effect-token-field14/) |

### SAV-1058

The six classes are `SpellEffect`, `PointEffect`, `AreaEffect`,
`SpellTransport`, `Effect` and `Effect_DirectDamage`. Every construction path
traced for them converges on `Token`'s own four constructors, the true `Token`
base constructors. No other instruction in the six classes' own creator,
constructor, destructor or post-load-hook bodies writes `+0x14` directly;
`Effect`'s and `Effect_DirectDamage`'s copy-shaped constructors reach it only by
calling one of the four (`ITEM-EFFOBJ-072`, `ITEM-CLASS-001`). The only other
writer of these six classes' `+0x14` outside the shared `Token` layer is
`Token::Serialize`'s LOAD arm (`SAV-1054`). The eighteen bodies named below with
an `evidence/*-body.txt` file are fresh whole-body reads, each with zero
unresolved (`evidence/instruction-counts.json`, `evidence/unresolved.json`,
`evidence/program-digest.tsv`).

`Token` constructors:

- Zero-constructors `R0923` (50 instructions,
  `evidence/token-basector1-R0923-body.txt`) and `R1009` (49
  instructions, `evidence/token-basector2-R1009-body.txt`) are near-identical.
  Both install vtable `L05965` (`L09024`/`L09025`), the vtable the shared
  `Token`-base real destructor `R1130` installs (`EXP-0382`). Both then
  write a literal `[this+0x14]=0` unconditionally (`L06039`/`L09026`),
  before allocating a 12-byte payload (`R0747`) and conditionally resolving
  it: `R0923`'s own arm through a global at `L04624`, `R1009`'s own arm
  through the caller's first argument (`ret 4`, one stack argument popped). That
  is a resolution-strategy difference, not a default/copy pair.
- `SAV-1011`'s two "`SpellEffect` base constructors" (`L08020`/`L05979`)
  call `R0923` and `R1009` respectively (`L08021`/`L08022`), not a
  single common first zero-constructor; both reached variants zero `+0x14`.
- Copy constructor `R1134` (76 instructions,
  `evidence/token-copyctor-R1134-body.txt`, `ret 4`) installs `L05965` and
  copies `+0x14` from its single argument (a load of the argument's `+0x14` at `L09027`,
  a store to the new object's `+0x14` at `L09028`). This matches `ITEM-EFFOBJ-072`'s "Copy
  `R1001` ... preserves the Token base": `Effect`'s copy constructor
  calls it.
- Owner constructor `R1135` (50 instructions,
  `evidence/token-ownerctor-R1135-body.txt`, `ret 8`) installs `L05965` and
  writes `+0x14` from its second argument (a load of the argument at frame slot `+0xc` and
  a store to the new object's `+0x14` at `L09029`/`L09030`).
- `ITEM-CLASS-001` independently reads a further caller of the first
  zero-constructor, `Item::Item` (`R0884`), calling `R0923` at `L06012`,
  matching the call-site census to the byte.
- A raw `E8 rel32` call-site census (`evidence/call-site-census.json`) finds 13
  direct callers of `R0923`, 7 of `R1009`, 2 of `R1134` and 1 of
  `R1135`. Of `R0923`'s 13, `L08021` is `SAV-1011`'s base-constructor
  site, `L09031` is in `Effect`'s default constructor and `L09032` in
  `Effect`'s second constructor `R1047`, not in `R1049`, whose
  body (`evidence/effect-defaultctor-R1049-body.txt`) ends at `L09033`, one
  byte before `R1047`.

`Effect`'s chain:

- factory `R1559` (30 instructions,
  `evidence/effect-creator-R1559-body.txt`) allocates 72 bytes, then calls
  default constructor `R1049` (`L09034`);
- `R1049` (24 instructions) calls `R0923` first (`L09031`), then
  installs vtable `L02981` and zero-writes
  `+0xe`/`+0x3c`/`+0x3d`/`+0x40`/`+0xc`/`+0x44`;
- second constructor `R1047` (71 instructions,
  `evidence/effect-ctor2-R1047-body.txt`) calls `R0923` at `L09032`,
  installs `L02981` and takes an incoming parameter at `[ebp+8]`;
- copy constructor `R1001` (34 instructions,
  `evidence/effect-copyctor-R1001-body.txt`; `ITEM-EFFOBJ-072`'s subject, read
  here for `+0x14`) calls `R1134` at `L09035` and installs `L02981`;
- `GetRuntimeClass`-shaped vtable `+0x0` slot `R1647` (8 instructions,
  `evidence/effect-getter-R1647-body.txt`) only returns the class descriptor
  address, touching no field;
- real destructor `R1648` (9 instructions,
  `evidence/effect-realdtor-R1648-body.txt`) is a bare forward to
  `R1130`, which `EXP-0382` read whole and found touching no `+0x14`;
- scalar destructor `R1649` (17 instructions,
  `evidence/effect-scalardtor-R1649-body.txt`) is the standard
  scalar-deleting-destructor thunk, with no field touch of its own.

`Effect_DirectDamage`'s chain:

- factory `R1636` (30 instructions,
  `evidence/effdd-creator-R1636-body.txt`) allocates 96 bytes, then calls
  `R1042`;
- `R1042` (26 instructions, `evidence/effdd-defaultctor-R1042-body.txt`)
  calls `Effect`'s default constructor `R1049` first (`L09036`, reaching
  `R0923` transitively), constructs its 24-byte raw span (`this+0x48`,
  `R1044`, `SAV-1033`'s subject), installs vtable `L02980` and zero-writes
  `+0x44`;
- second constructor `R1650` (32 instructions,
  `evidence/effdd-ctor2-R1650-body.txt`) calls `R1049` at `L09037` and
  installs `L02980`;
- copy constructor `R1068` (40 instructions,
  `evidence/effdd-copyctor-R1068-body.txt`) calls `R1001` at `L09038`
  (reaching `R1134` transitively) and installs `L02980`;
- `GetRuntimeClass` slot `R1637` (8 instructions,
  `evidence/effdd-getter-R1637-body.txt`) has `Effect`'s shape;
- real destructor `R1069` (9 instructions,
  `evidence/effdd-realdtor-R1069-body.txt`) forwards to `R1648`
  (`L09039`), which forwards to `R1130`;
- scalar destructor `R1638` (17 instructions,
  `evidence/effdd-scalardtor-R1638-body.txt`) is the same thunk, forwarding to
  `R1069`.

`R1637` and `R1638` are the two vtable slots `SAV-1033`'s closing
paragraph left untraced. These reads close the gap `EXP-0382`'s four-class
destructor census and `EXP-0383`'s four-class creator/zero-target-constructor
census (both over `PointEffect`, `AreaEffect`, `SpellTransport` and
`SpellEffect`'s base) leave for `Effect` and `Effect_DirectDamage`, and add the
two zero-constructors neither reads and the copy and owner constructors. Also
cites `SAV-CLASSSER-173`.

**Confidence.** High for the exhaustive absence within the population read:
every construction, destruction and post-load-hook body of the six classes (the
eighteen fresh reads plus the published whole-body reads of `SAV-1011`,
`SAV-1055`, `EXP-0382` and `EXP-0235`, none re-derived) is zero-unresolved and
instruction-accounted for `+0x14`, and the four `Token`-level constructors, two
zero and two argument-derived, are the sole literal writers found in it. Medium
for the image-wide writer population: most of the census's further callers (11
of `R0923`'s 13, 5 of `R1009`'s 7) reach classes outside the six, traced
only for `MAGIC-219`'s purpose and not attributed individually here (`Item`,
`Building` and `Unit` among them; `ITEM-CLASS-001`). The claim is not that the
four are the only place in the image `+0x14` is written. This excludes A1
(always zero) as stated: zero is one of four constructed defaults, and producers
exist (`Token::Serialize`'s LOAD arm, `SAV-1054`, and the copy and owner
constructors write it nonzero from an argument). It narrows A2/A3/A4 to what
`SAV-1059` finds in the corpus, since no class-specific literal producer beyond
the shared four-constructor layer exists to derive a class-specific meaning
from.

**Amended.** The clause that both `SpellEffect` bases `L08020/L05979` "are
themselves callers of the first zero-constructor" is partially retracted by
`SAV-1068` (EXP-0388; [`retracted.md`](retracted.md)): `L08022` calls
`R1009`, which builds Position from its argument. The call identities above
replace it. Both Token constructors still zero `+14`; the `+0x14` stores and the
four-constructor inventory stand.

### SAV-1059

- Population: `tools/savfull` (EXP-0251) re-run over every `.sav` path below
  `gameversions/` (`experiments/EXP-0384-effect-token-field14/corpus.sh`). The
  raw walk parses 76 distinct-SHA256 saves and fails exactly one,
  `saves/2026-08-27/EXP-0261-owner-runs/game9000.sav`, with the error
  `header magic "Bsg&" is not Asg&`. That file is an Againrom-writer candidate,
  labeled "generated candidate" by its directory's own `MANIFEST.md`, and is
  excluded because it never parses.
- An explicit provenance filter removes six further paths, all named in
  `evidence/population.txt`: `game9001-9004,9006.sav` from the same
  `MANIFEST.md`-labeled directory, and one `generated-repeat/game0000.sav` path;
  `pipeline/archive/LOG-2026-08-25-to-09-07.md` confirms that a sibling
  `generated/game0000.sav` in that directory is Againrom's own writer output. 70
  saves are admitted.
- Within the 70, two directories are a named sub-population of ROM1-written
  resaves of an Againrom-produced or Againrom-modified document rather than
  unmodified originals: `saves/2026-08-27/EXP-0261-owner-runs` (8 admitted,
  after the six-path exclusion) and `saves/2027-09-07` (30 admitted), 38 saves
  together. The 2026-09-07 entries of
  `pipeline/archive/LOG-2026-08-25-to-09-07.md` document them as world/city
  saves ROM1 wrote after loading a document produced or modified by this
  project's writer, container, archive-writer or compressor ("rung 1", the
  original's bytes with a purse value this project injected; "rung 3", the same
  content carried through this project's container, archive writer and
  compressor).
- The remaining 32 admitted saves, including `game0076.sav`, carry no such note.
  Under this repository's `gameversions/saves/<date>/` convention for owner
  saves (`AGENTS.md`) they are treated as ordinary owner-played saves; no
  per-file confirmation is traced beyond the two citations made for
  `game0076.sav` and `game0125.sav`.
- `identity-relations.tsv` filtered to `kind=="owner"` on the six target classes
  (`evidence/six-class-owner-references.tsv`,
  `evidence/six-class-status-summary.tsv`). `owner` is this repository's name
  for the `Token+0x14` relation, from `tools/savfull`'s shared `w.token(p)`
  parser, cited as a naming convention only, not as evidence of meaning. Counts:
  - `Effect`: 174 `zero` + 18 `resolved`;
  - `SpellTransport`, `PointEffect`, `AreaEffect`, `Effect_DirectDamage`: 1
    `zero` each;
  - `SpellEffect`: no row of any kind, not merely no `owner`-kind row,
    consistent with, but not re-deriving, `EXP-0246`'s older census.
- The four single-instance zeros come from `en/game0018.sav` (`SpellTransport`,
  `PointEffect`, `Effect_DirectDamage`; the file `SAV-1034`'s `n=1` witness
  reads) and `saves/2027-09-07/game0125.sav` (`AreaEffect`). `game0125.sav` is
  in the 38-save sub-population, not an original: the 2026-09-07 entry of
  `pipeline/archive/LOG-2026-08-25-to-09-07.md` names it (label "5m city A
  saved", 30,826 bytes) as descending from "rung 1", and `pipeline/LOG.md:80`
  independently lists it as "world save from the transplant rung", a ROM1 resave
  of a document this project produced.
- All 18 `resolved` `Effect` rows are in `saves/2026-09-09/game0076.sav`, sha256
  `ca6f2980…`. The digest matches `pipeline/LOG.md`'s recorded digest for that
  file, and `pipeline/LOG.md:1438` notes "Original game0076 is preserved from
  the owner's separate GOG installation". Every one resolves to
  `target_class=Human` (`evidence/six-class-resolved-targets.tsv`).
- A containment check
  (`experiments/EXP-0384-effect-token-field14/containment.py`,
  `evidence/effect-owner-containment.tsv`) compares each reference's file offset
  with the surrounding `Human` record's end-bounded archive span
  (`ref_at`..`end` on the archive event that decoded that `Human`, joined by
  `object_index`, not the span between one `Human` definition and the next). 18
  of 18 resolve inside that span, with 0 join misses.
- The same check records each reference's immediate container: the minimum-width
  decoded archive span of any class, excluding the referencing `Effect` record's
  own span, that holds the reference's offset. In all 18 it is an `Armor`,
  `Weapon` or `Shield` record nested inside the `Human`'s span, never the
  `Human` record directly: 13 `Armor`, 4 `Weapon`, 1 `Shield` (the
  `immediate_container_class` column). For example, `game0076.sav`'s `Effect` at
  file offset 1337 has `immediate_container_class=Armor` and resolves to the
  `Human` span 238..1621.
- So the referenced object is the `Human` whose own body embeds the `Effect`,
  reached through an equipped `Armor`/`Weapon`/`Shield` record: a
  self-referential owning-actor back-pointer, not a `Player` and not the
  `SpellEffect+0x3c` caster. The claim does not enumerate a generic equipment or
  status-effect list.
- Further evidence files: `evidence/image-summary.tsv`,
  `evidence/savfull-exit-note.txt`. Also cites `SAV-CLASSSER-177`,
  `SAV-TOKEN-034`.

**Confidence.** High for the corpus mechanics: `tools/savfull`'s
identity-resolution method is published (`SAV-PTRMAP-035`, `SAV-ARCHREL-253`),
the population count and every exclusion are named and reproducible
(`evidence/population.txt`), and the containment check is a direct file-offset
comparison against each `Human`'s true archive span over all 18 resolved
instances with 0 join misses. Medium for what the corpus establishes about
meaning. `n=1` for four of the six classes is consistent with A1 (always zero)
for those four but does not exclude a producer this corpus never exercises, the
bound `SAV-1034` states for `Effect_DirectDamage`'s span; one of the four,
`AreaEffect`'s zero, comes from a ROM1 resave of a produced document rather than
an unmodified original. For `Effect`, the corpus excludes A1 (nonzero is
observed) and A2 as stated (the target is a `Human`/actor, never shown to be
that `Human`'s own `Player`) at High within this one save, itself an
independently confirmed original.

**Unknown.** Whether every `Effect` in every save behaves this way, or only
effects reachable through this save's particular equipped-item shape. The
finding is bounded to `n=18` in `n=1` save; a second original save with resolved
`Effect` references is needed to state it as a population-wide `Effect` rule.

### SAV-1060

- The two located readers of the six classes' own `+0x14`: `Token::Serialize`'s
  STORE arm (`SAV-1054`), and `Token`'s copy constructor `R1134`, which
  reads a source object's `+0x14` at `L09027` to copy it into the object under
  construction (`SAV-1058`).
- No destructor, post-load hook, or either published per-tick attribution tail
  (`PointEffect`, `AreaEffect`) reads `this+0x14` where `this` is an object of
  one of the six classes.
  - Destructors: `EXP-0382`'s four-class real/scalar-destructor census
    (`PointEffect`, `AreaEffect`, `SpellTransport`, `SpellEffect`'s own base,
    plus the shared `Token`-base real destructor `R1130`) and the
    `Effect`/`Effect_DirectDamage` destructor reads (`SAV-1058`) contain no
    `+0x14` displacement.
  - Post-load hooks: `EXP-0235`'s six whole-body reads (`R1115`, `L05904`,
    `R1116`, `R1482`, `R1569`, `R1560`; the shared hook covering
    `Token`, `SpellEffect`, `Effect`, `Effect_DirectDamage` and ten further
    classes) contain no `+0x14` displacement.
- The two tick tails read `+0x14` on other, separately named objects reached by
  first dereferencing a field: `PointEffect`'s tail five times, `AreaEffect`'s
  three. Each has two caster reads (`SAV-1011`'s `+0x3c`, the
  `MAGIC-188`/`MAGIC-217` liveness gate) and one or two target reads
  (`MAGIC-TARGETID-182`'s `+0x44`-derived object). `PointEffect`'s tail alone
  has one further read, on the object the attached `Effect` payload's `+0x44`
  points at. The reads are grep-verified against `EXP-0383`'s committed bodies
  `evidence/pointeffect-tick-attribution-R0266-body.txt` and
  `evidence/areaeffect-tick-attribution-R0267-body.txt`, cited not re-derived.
- `R0266` (`PointEffect`), five `+0x14`-displaced reads:
  - `L09040`: `(this+0x48)`'s `+0x44` target's `+0x14`, a cleanup of the
    attached `Effect` payload's reference, which clears that reference to null
    when this read is zero. `MAGIC-TARGETID-182` names the field it clears,
    `this+0x48`'s own `+0x44`, but not this `+0x14` test.
  - `L05864`: the target's `+0x14`, guarded by the null check at `L05863`
    immediately before it. This, not `L09041`, is `MAGIC-TARGETID-182`'s "a
    guarded read".
  - `L06000` and `L06002`: the caster's `+0x14`, `MAGIC-217`'s two liveness
    checks.
  - `L09041`: a second, independent read of the target's `+0x14`, immediately
    before the final notify call, guarded only by the target pointer's null
    check at `L09042`; no earlier claim counted it.
- `R0267` (`AreaEffect`), three: `L06001` and `L06003` (caster, the
  same two liveness checks) and `L09043` (a single target read).
  `AreaEffect`'s tail has no equivalent of the `L09040` cleanup read.
- The target reads are `UNIT-OWNER-009`'s "owning Player" field.
  `UNIT-OWNER-009` publishes `actor+0x14` as "the owning Player" for a
  `Unit`/`Human` object generally, and every target-classed read here traces to
  such an object (`MAGIC-TARGETID-182`'s "target"). They are the target unit's
  owner-reference gate, the same field with the same established meaning, not a
  new field on a new class and not a new reader of one of the six classes. This
  names what `MAGIC-TARGETID-182`'s text leaves unnamed.
- The `L09040` read's object, the attached `Effect` payload's `+0x44` target,
  is named by no existing claim, and this claim does not identify its class.
- LOAD's miss zero from `R1371`: `token.md`'s "Identity map" section states
  that the null-on-miss outcome is written into the field itself
  (`SAV-TOKEN-034`, `SAV-PTRMAP-035`). No class-specific consumer of that zero
  beyond the shared `Token::Serialize` LOAD arm was found for any of the six
  classes.
- Further evidence files: `evidence/effect-realdtor-R1648-body.txt`,
  `evidence/effect-scalardtor-R1649-body.txt`,
  `evidence/effdd-realdtor-R1069-body.txt`,
  `evidence/effdd-scalardtor-R1638-body.txt`,
  `evidence/token-copyctor-R1134-body.txt`.

**Confidence.** High for the read census within the population read: every
destructor, post-load hook and tick-tail body for the six classes is
zero-unresolved and grep-verified for `+0x14`, both tick tails' `+0x14`
displacements are completely enumerated (five and three), and the copy
constructor's source read is a located reader. Medium for completeness beyond
that population: `Spell::Apply` and other cast-path functions past what
`SAV-1054`/`SAV-1055`/`SAV-1056` read are not re-walked, so an unlocated reader
elsewhere in the image is not excluded, only absent from the bodies this
experiment and its cited claims have read.

**Unknown.** The class of the `L09040` read's object; it is not traced.

## Order block raw copy and a creature's spellbook slots

Claim ids allocated to `EXP-0385`, `EXP-0386` and `EXP-0387` and returned
unused. A returned range stays retired, and no id in it is ever reissued.

- `EXP-0385` was allocated ids `1062`..`1063` of `claims/sav.md` and spent none,
  so `SAV-1062`..`SAV-1063` (2 ids) are returned unused. Its dying-population
  and crossing-window findings are published in `claims/hero.md`
  (`HERO-DYINGTICK-145`, `HERO-CROSSHOLD-146`) and `claims/ai.md` (`AI-332`),
  citing `formats/sav/actors.md` where the SAV corpus is the evidence.
  `formats/sav/actors.md` already cites non-`SAV`-prefixed IDs, so no SAV-stem
  row was needed.
- `EXP-0386` was allocated ids `1064`..`1065` and spent none, so
  `SAV-1064`..`SAV-1065` (2 ids) are returned unused. Its Part 2 guard-post and
  SAV-coverage question (the brief's Q7) is answered by citation:
  `AI-GRPGUARD-074`, `AI-POST-095`, `AI-POST-097`, `AI-LOAD-099`,
  `AI-GUARD-021`, `AI-PATROL-018`, `AI-SIGHT-006` and `SAV-PATROLCURSOR-571`
  give the complete field, routine and tick account, and the experiment's read
  set adds nothing to it (`AI-337`, `claims/ai.md`).
- `EXP-0387` was allocated ids `1066`..`1067` and spent `SAV-1066`, 1 id;
  `SAV-1067` is returned unused. The Part 4 question, whether a SAV carries a
  creature's spellbook, needed one row: the raw-copy structure and its
  practical-equivalence limitation are one finding. The next id after this range
  is `SAV-1068`.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1066 | `Order::Serialize` copies order bytes `0x00`-`0x93` raw and unconditionally on STORE and LOAD, so an original SAV carries a creature's three spellbook slot pairs at `ord+0x78`..`+0x8c`. | High / Medium | ● active | [EXP-0387](../experiments/EXP-0387-creature-spellbook/) |

### SAV-1066

- `Order::Serialize` (`R0199`), fresh whole-body read
  (`evidence/functions.txt`, 78 instructions): `CArchive::IsStoring`
  (`R1123`) branches STORE, `CArchive::Write(this, 0x94)` (`R0271`
  at `L00550`), against LOAD, `CArchive::Read(this, 0x94)` (`R0272` at
  `L00551`). Both are byte-count-only calls with no field enumeration in
  either arm, which matches `SAV-PATROLCURSOR-571`'s "148 raw bytes" at the
  instruction level.
- `0x94 = 148` covers offsets `0x00`-`0x93` inclusive, which contains all six
  spell-slot dwords (`0x78`, `0x7c`, `0x80`, `0x84`, `0x88`, `0x8c`, each
  `< 0x94`): the three class-spellbook slot pairs at `ord+0x78`..`+0x8c`
  (`AI-341`, `UNIT-SPELL-007`). An original SAV carries them; LOAD does not
  rebuild them from class.
- The only field in that range the routine treats differently from a raw copy is
  `+0x90`. The STORE arm only reads the existing `u16list` object and calls its
  `vt+8` Serialize with the archive (`L09044`/`L09045` load the object,
  `L09046` calls `vt+8`); it builds nothing. Only the LOAD arm rebuilds
  `+0x90`: it deletes the old object, allocates `0x1c` bytes (`L09047`),
  constructs a fresh one via `R1545(0xa)` (`L09048`) and stores it at
  `+0x90` (`L09049`), the "replaces order+90 with a fresh counted-word list"
  behaviour `SAV-PATROLCURSOR-571` names for the LOAD arm. The six spell-slot
  dwords are not among the fields either arm treats specially.
- This is a structural fact about the routine's body, not a trace of the wider
  LOAD call sequence. The raw copy is unconditional and would overwrite any
  earlier class-derived value (via `R0184`, `UNIT-SPELL-007`) with the
  archive's own bytes regardless of ordering, so the ordering does not change
  the outcome.
- `AI-341`'s whole-image census and base-register disambiguation covers all six
  slot-pair displacements plus eighteen narrower single-byte ones
  (`evidence/refs.txt`): 203 raw AI-module hits, each traced to a structure
  other than the order block, chiefly the actor's mover, which carries fields at
  the same `+0x7c`/`+0x80` offsets. Within the AI module it covers, it finds no
  writer of the six dwords other than `R0184`'s spawn-time class setup
  and this routine's raw LOAD copy.
- So on the current image, a value LOAD carries from an original SAV and a value
  a hypothetical class re-derivation would produce are byte-identical whenever
  nothing changed the slots between spawn and save. "Carried, not rebuilt" is an
  instruction-proven property of the SAV mechanism, but EXP-0387 found no
  runtime scenario in which a change between spawn and save would be lost by
  carrying the SAV's bytes rather than re-deriving them.

**Confidence.** High for the raw-copy structure and its `0x94`-byte span
containing the six offsets: a complete-body read, zero unresolved, matching
`SAV-PATROLCURSOR-571`'s published figure. Medium for "carried, not rebuilt"
mattering in practice: it is bounded by the whole-image literal-displacement
census `AI-341` runs, which cannot see a writer reached through a pointer or
another register-computed access.

**Unknown.** Whether unit construction, which gives a fresh actor's `Order` its
class-derived defaults, runs before or after this routine's LOAD arm in the
wider load sequence; it was not traced.

## Cast construction field sources

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1068 | The selected PointEffect and SpellTransport cast constructors leave Token `+08` and `+0c` without a direct store; their Position sources differ. | High / Unknown | ✔ promoted | [EXP-0388](../experiments/EXP-0388-effect-construction/) |
| SAV-1069 | The two selected cast paths initialize common SpellEffect bytes `+40=0` and `+41=1` with literal stores. | High / Medium / Unknown | ✔ promoted | [EXP-0388](../experiments/EXP-0388-effect-construction/) |

### SAV-1068

- PointEffect `R1046 -> L08020 -> R0923` reaches the default Token path.
  SpellTransport `R0633 -> L05979 -> R1009` reaches the
  Position-argument path (`L08021` versus `L08022`). Both reached variants
  still zero `+14`.
- Both call `SAV-678`'s reset `R1624`, writing `+04/+18/+1c=0`. The shared
  SpellEffect bases write `+0e=0`.
- Object-base `L08171` writes only its vtable, and the embedded-list chain
  `L09050 -> L09051` operates at Token `+20..+3b`, so neither supplies
  hidden stores to `+08/+0c`.
- PointEffect then copies target Position through `L05872` at `L08023`: this
  is Position assignment, not target registration.
- SpellTransport copies its Position argument through `R1010` at `L07895`;
  its native caller supplies caster Position.
- `SAV-TOKENPOS-074` establishes that both copies transfer `+00..+05` and
  `+08..+0b` and omit destination `+06/+07`.
- Token's saved-address key is `this`, not an initialized member.
- The 18-body, two-window probe retains both opposing synthetic fills in Token
  `+08`, constructor `+0c` and destination Position `+06/+07`; `MAGIC-225`
  distinguishes the caller's later writes.
- This claim corrects `SAV-1058`'s first-zero-constructor clause, now partially
  retracted, and the target-registration gloss in `SAV-1010` and
  `MAGIC-TICKGATE-183`, without changing their target store/dereference
  findings.
- Evidence files: `evidence/body-index.tsv`, `evidence/fresh-bodies.txt`,
  `evidence/vectors.tsv`, `evidence/write-events.tsv`,
  `evidence/call-boundaries.tsv`. Also cites `SAV-1054`.

**Confidence.** High for the direct store and no-store facts within the
539-instruction population and for the corrected call identities, verified
against the lawful image and 12 isolated runs.

**Unknown.** Native allocation values, transitive effects of controlled external
calls, arbitrary aliases and later first-SAVE values.

### SAV-1069

- PointEffect's base `L08020` writes them at `L09052/L09053`;
  SpellTransport's base `L05979` writes them at `L09054/L09055`.
- These are the two byte sources serialized after the Token head
  (`SAV-CLASSSER-173`).
- Both values survive the selected constructors in the isolated instruction
  probe.
- They do not form a universal first-SAVE invariant: PointEffect's direct caller
  already replaces `+41` from the cached Spell flag (polarity established by
  `MAGIC-ATTRGATE-118`), while the selected transport caller leaves its own
  common bytes unchanged.
- Evidence files: `evidence/body-index.tsv`, `evidence/vectors.tsv`,
  `evidence/write-events.tsv`, `evidence/call-boundaries.tsv`. Also cites
  `MAGIC-225`.

**Confidence.** High for the four literal stores and the controlled constructor
outputs. Medium for full native admission coverage: the probe controls
allocation, distance and node allocation rather than tracing their internals.

**Unknown.** All later writers before first SAVE.

## World-head transition frontier

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1070 | Selected transition leaves preserve the eleven head words locally; their callers do not establish interval preservation. | High / Unknown | ● active | [EXP-0389](../experiments/EXP-0389-world-head-transitions/) |
| SAV-1071 | The bounded transition-root expansion leaves the ordinary return's same-world alias and callback effects open, after pruning campaign-excluded and overlay controls. | High / Unknown | ● active | [EXP-0389](../experiments/EXP-0389-world-head-transitions/) |

### SAV-1070

- Entry `L09056` passes `world+30` to `L09057`, whose complete
  seven-instruction body has no call or object store.
- Through `L09058`, `L09059` passes `world+170` to the nine-instruction
  read-only leaf `L09060`; its returned pointer later escapes to a file
  helper.
- Return `L09061` passes the pointer value M loaded from `world+14` to
  `L09062`; local clears cover M+[0,c), not head offsets. The retained
  constructor sets `world+14=L05906`; later arbitrary aliases are unproved.
- SAVE wrapper `L09063/L08401` transfers current global `L00285` to
  `R0804` without an intervening call; the outer writer saves and reloads
  that same argument for `R0414` at `L08403`.
- No cross-transition receiver replacement or universal preservation is
  established.
- Evidence files: `anchors.tsv`, `receivers.tsv`, `evidence/functions.txt`,
  `evidence/decoder-audit.tsv`. Also cites `SAV-WHEADINIT-520`,
  `SAV-WHEADLIMIT-525`.

**Confidence.** High for the named instruction-local effects and receiver
transfers, checked in both disassemblies.

**Unknown.** Native reachability, effects of external callees, and complete
transition-to-SAVE identity or values.

### SAV-1071

- Twelve new bodies plus three retained controls contain 2,067 instructions.
  Their syntactic surface before the chosen SAVE cutoffs holds 199 call sites,
  172 of them unexpanded here. Both counts include conditional controls and are
  not ordinary-route reachability counts.
- Return's `world+6c -> L09064` forwards that alias with arguments 0/-1 to the
  known word-array sizing entry `R1622` at `L09065` (`SAV-667`); its
  complete zero-size footprint and external effects are unexpanded.
- Return callback `L09066` has a located receiver but no proved vptr producer
  or target.
- Separately, conditional overlay sites `L09067/L09068 -> L09069` forward
  `world+48` to `L09070` at `L09071`, then write argument 2 through the returned
  pointer at `L09072`. This unclosed destination is not a shipped-campaign
  mutation witness under `TRIG-INI-012`/`ITEM-SPAWN-026`.
- The overlay item-element store/callback at `L09073/L09074` is likewise
  conditional, and `ITEM-SPAWN-027` already excludes `R0945` on campaign
  entry.
- No selected-head store is established in this subset. Preservation and
  intervening mutation both still fit the admitted ordinary-route evidence; no
  all-path first-SAVE vector or `+140/+144` semantics follow.
- Evidence files: `expansion.tsv`, `receivers.tsv`, `evidence/scoped-calls.tsv`,
  `evidence/summary.json`. Also cites `SAV-WHEADLIMIT-525`.

**Confidence.** High for the finite population and the named pointer and call
frontier. Existing route exclusions are reused, not republished as discoveries;
corpus agreement does not choose a model.

**Unknown.** Transitive mutation, ordinary nonzero producers, whole-interval
preservation and native first-SAVE values.

## Saved-world actor population

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1074 | The selected world-present LOAD builds the map builder's actor-key index from existing actors; it does not establish an added actor from an absent ALM placement. | High / Medium / Unknown | ● active | [EXP-0390](../experiments/EXP-0390-saved-actor-population/) |

### SAV-1074

- Document `L01841` calls shared builder `R0067` on `world+44` with map and
  literal 0.
- Three passes store iterator actor pointers under keys from `actor+8`:
  - the global actor manager (`L09075/L09076`);
  - the dead-manager collection (`L09077/L09078`);
  - each Player's `+20` collection (`L09079/L09080`).
- The dead receiver follows the retained constructor's `world+68 = world+14`
  alias, since builder `+24` is world `+68`; arbitrary intervening alias
  mutation is not excluded.
- Those loops contain no local health/stage filter and no direct
  placement-spawner call.
- Numeric reference helper `R0507` reads that index. Its missing-key
  association path reaches the unexpanded value initializer `L09081` at
  `L09082`, so not even a zero miss result is claimed.
- Fresh entry separately reaches `R0151` at `L00958`; the ordinary resume
  argument bypasses the fresh constructor.
- The selected authored `0x10003` control counts existing Group actors, requests
  removal/destruction and writes counted-array records, but has no witness in
  the retained 38-map vocabulary census. Its external helpers and callbacks
  remain open.
- Twelve new bodies and eight retained controls do not close archive factories,
  arbitrary callbacks or the pre-tick boundary. No all-path archive exclusivity
  or additive ALM population rule follows.
- Evidence files: `anchors.tsv`, `retained-anchors.tsv`, `expansion.tsv`,
  `routes.tsv`, `evidence/functions.txt`, `evidence/decoder-audit.tsv`. Also
  cites `SAV-GRPLOAD-560`, `SAV-946`, `SAV-LOADREG-878`, `SAV-DEADLOAD-125`,
  `SAV-POSTLOAD-220` (corpus clause partially retracted), `SAV-949`.

**Confidence.** High for the named local instructions and pointer sources.
Medium for the bounded route account. The control census is reused historical
evidence, not a new root-wide survey or native witness.

**Unknown.** Missing-key initialization, transitive population changes, actor
identity outside collected memberships and the first ordinary tick.

## Session timer divisor

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1078 | The retained order-routine fault site divides by the raw session dword at +28; its local unsigned arithmetic requires that dword to be nonzero. | High / Medium / Unknown | ● active | [EXP-0391](../experiments/EXP-0391-owner-crash-divisor/) |
| SAV-1079 | The retained session initializer queries a frequency into +18 and derives +28 through a helper; 26 preserved original session files hold 10,000,000 and 10,000 respectively. | High / Medium / Unknown | ● active | [EXP-0391](../experiments/EXP-0391-owner-crash-divisor/) |

### SAV-1078

- Entry `L09083/L09084` fixes the frame base at receiver+8. At `L09085` the divisor is loaded
  from frame offset +20; `L09086/L09087/L09088` take the low counter difference as the dividend, clear the high half and execute unsigned division by the divisor.
- For every nonzero u32 divisor the quotient fits u32; a zero divisor faults for
  every dividend.
- `SAV-SESS-031`'s length-48 raw span makes this session-wire offset 1432,
  transferred by both SAVE and LOAD.
- Four retained event records join into two error/WER pairs with exception
  c0000094, module offset 001317b9 and timestamp 35c8443a, but contain neither
  SAV identity nor registers.
- All four selected Againrom compatibility files carry an all-zero head and
  therefore the same wire violation; they supply no original semantics.
- Evidence files: `evidence/retained-anchors.tsv`, `evidence/events.tsv`,
  `evidence/session-corpus.tsv`, `evidence/arithmetic-vectors.tsv`. Also cites
  `AI-PROGRESS-034` (switch-liveness clause partially retracted),
  `AI-ROUTE-045`.

**Confidence.** High for the selected retained-text matches, the raw-span
mapping and the mathematical implication. Medium for attributing the recorded
crashes to those files. No fresh decoding, machine emulation or native LOAD
trace occurred.

**Unknown.** Current image identity, native receiver/register values, the opened
SAV and intervening writers.

### SAV-1079

- Constructor `R0077` calls `R0124` with its receiver. The retained
  initializer passes receiver+18 to import L09089, identified as
  QueryPerformanceFrequency by the existing import table.
- At `L09090..L09091` it passes that u64 and u64 1000 to helper L09092;
  `L09093` stores the low return at receiver+28.
- The helper's unsigned-division interpretation fits its arguments and the
  measured population, but its body was not checked.
- The exact SAV reader measures 23 preserved EN files (22 with a session and one
  without) and four preserved RU files. Every one of the 26 session-bearing
  files carries the pair at wire offsets 1416 and 1432.
- Evidence files: `evidence/source-manifest.tsv`,
  `evidence/retained-anchors.tsv`, `evidence/session-corpus.tsv`. Also cites
  `SAV-SESS-031`.

**Confidence.** High for the local retained call/store facts and the bounded
numeric population. Medium for the helper's quotient semantics. This is no
universal Windows frequency, cross-host save rule or all-path constructor
invariant.

**Unknown.** Later writes, other timer-head fields and native correction
acceptance.

## Mercenary pool persistence

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1084 | A SAVE carries the mercenary pool as one u32 count, read from the working array's `m_nSize`, then both arrays' payloads under that count; the pristine array's `m_nSize` is never read. | High | ● active | [EXP-0395](../experiments/EXP-0395-mercenary-pool-persistence/) |
| SAV-1085 | LOAD establishes the mercenary pool from the document's own bytes, not from `[General] MercenaryCount`, a constructor default or the record reset. | High / Unknown | ● active | [EXP-0395](../experiments/EXP-0395-mercenary-pool-persistence/) |
| SAV-1086 | Across 119 admitted saves the serialized mercenary element count is 15 without exception, and the two arrays are equal in every file but one. | High / Medium | ● active | [EXP-0395](../experiments/EXP-0395-mercenary-pool-persistence/) |

### SAV-1084

- In `R1294`, `L09094`, a load of the dword at offset 0x64 of the record, fetches the
  working array's size, `L09095`…`L09096` write it as four bytes, and
  `L09097`, a test of that size, with a branch to `L09098` when zero, skips both payloads when it is
  zero.
- The first payload is set by `L09099`, doubling the size to a byte count, and
  `L09100`, loading the dword at offset 0x60 of the record. The second re-reads the same count
  from the stack slot, `L09101`, a load of the dword at stack offset +0x10, and writes
  `L09102`, a load of the dword at offset 0x74 of the record, under
  `L09103`, doubling the count to a byte count.
- Through its mercenary section the writer body references exactly three of the
  four pool fields, `+0x64`, `+0x60` and `+0x74`, and `+0x78` nowhere.
- A consumer therefore cannot express two different lengths for the two arrays.
  A grammar with a second count between the payloads is one this writer cannot
  produce, and a reader that expects one reads a shape that never occurs.
- This gives `SAV-CAMPPROG-071`'s wire programme item 2 its instruction-level
  source for each field. The single shared count read once at `L09094` and the
  absent `+0x78` are what this claim adds.
- Evidence files: `evidence/d-pool-release.txt`, `evidence/corpus-searches.txt`.
  Also cites `SAV-CAMPTAIL-070`, `SAV-CAMPAIGN-083`, `MERC-POOL-011`.

**Confidence.** High. Reading the writer body through its mercenary section at
instruction level, rather than inferring from the file, excludes both live
alternatives the question named: that the pool is not serialized at all, and
that the two arrays carry independent counts.

### SAV-1085

- `R0434` reads a u32 count at `L09104`…`L09105` and calls `SetSize`
  on both arrays with it: `L09106`, forming the address 0x5c bytes into the record, with
  the call to `R1622` at `L08024`, then `L09107`, forming the address 0x70 bytes into the record, with
  the call to `R1622` at `L09108`.
- Unless `L09109`, a test of the count, with a branch to `L09110` when zero takes the zero branch, it
  raw-reads `2·count` bytes into each: `L09111`, loading the dword at offset 0x60 of the record,
  with `L09112`, doubling the count, and `L08025`, an indirect call through the held function pointer, then
  `L09113`, loading the dword at offset 0x74 of the record, with
  `L09114`, doubling the count, and `L09115`, an indirect call through the same pointer.
- There is no default to fall back to. The campaign-record constructor
  `R0677` builds both arrays empty, `L09116`, forming the address 0x5c bytes into the record, with
  the call to `R1549` at `L09117` and `L09118`, forming the address 0x70 bytes into the record, with
  the call to `R1549` at `L09119`, and that constructor zeroes `m_pData`, `m_nSize`,
  `m_nMaxSize` and `m_nGrowBy`.
- Neither campaign LOAD driver, `R0098` or `R0099`, calls the
  record reset `R0757`, the constructor, or a `MercenaryCount` read. The
  reset's only three call sites (`L09120`, `L08431`, `L03810`) are all
  new-campaign initialisation.
- At the two drivers' five `R1430` sites the section argument is always a
  literal: `L09121` "Objects" at four, `L09122` "Projectiles" at `L09123`.
  The key is a literal at only three: `L09124` "Selection" at `L09125` and
  `L09126`, `L09127` "IDs" at `L09123`. At `L09128` and `L09129` the
  key is a stack buffer formatted at run time by a call to `R1196` from the
  format string at `L09130`, whose text was not printed.
- The run-time keys do not reopen the question: pushes of `L09131` "General",
  `L09132` "MercenaryCount" and `L09133` "Scenario\\Scenario.reg"
  occur zero times in either exported driver body, so no key formed at run time
  is read out of the section the pool lives in.
- The consequence a consumer must implement: a loaded working array survives the
  load unchanged, including one that disagrees with the registry. A serialized
  count of 0 loads as two arrays whose `m_pData` is NULL, because `R1622`
  with size 0 frees the buffer and zeroes the three fields
  (`L09134`…`L09135`), and nothing downstream repopulates them.
- Evidence files: `evidence/d-pool-release.txt`,
  `evidence/d-drivers-release.txt`, `evidence/d-reset-callers-release.txt`,
  `evidence/strings-release.txt`, `evidence/callers2-release.txt`,
  `evidence/corpus-searches.txt`. Also cites `SAV-CAMPTAIL-070`,
  `SAV-CAMPPROG-071`, `REG-SCN-067`, `SESS-START-034`, `MERC-POOL-012`.

**Confidence.** High. It excludes each of the three alternatives the question
named, and says by what. A recomputation from the registry: both drivers were
read for such a read, and every registry section they push is a printed literal
and a different section from the pool's. A constructor default: the
constructor's two calls are the empty-array constructor. A post-load reset: the
reset has exactly three call sites and none is downstream of either driver.

**Unknown.** No prediction this claim supports was run against the original.

### SAV-1086

- Population: every `.sav` below `gameversions/en`, `gameversions/ru`,
  `gameversions/rom1-demo` and `gameversions/saves`, minus `SAV-1059`'s two
  provenance rules: any `generated` path, and the six `MANIFEST.md`-labelled
  generated candidates in `saves/2026-08-27/EXP-0261-owner-runs`. This removes
  12 paths and admits 119, over 72 distinct SHA-256 and 12 directories, spanning
  main missions 10, 20, 30, 40, 50 and 110.
- Every admitted file decodes with element count 15.
- The population holds exactly three distinct working/pristine/flags triples:
  - 105 records over 58 distinct files and all 12 directories carry both arrays
    equal to `1,1,1,1,1,4,4,3,3,3,0,3,4,3,0` with no hire flag set;
  - 13 records over 2 directories carry the same two arrays with the type-14
    hire flag set;
  - exactly one file, `saves/2027-09-07/game0034.sav` at mission 40, carries
    `working[14] = 2` against `pristine[14] = 3` with the type-6 flag set.
- The repeated array is `[General] MercenaryCount` as `REG-SCN-067` reads it,
  which is what the reset's copy loop produces. The single departure has the
  shape the mission-end merge's hired arm produces.
- `install-en` (23 paths), `install-ru` (4) and `install-demo` (1) carry only
  the first triple.
- Evidence files: `evidence/population.txt`, `evidence/pool-corpus.tsv`,
  `evidence/pool-corpus-summary.txt`. Also cites `MERC-DEATH-006`, `SAV-1084`,
  `SAV-1085`.

**Confidence.** Medium. Corpus agreement is the whole instrument: 119 agreeing
files do not exclude a producer this corpus never exercises, and both
directories holding a nonzero hire flag are inside `SAV-1059`'s named
sub-population of ROM1 resaves of an Againrom-produced or Againrom-modified
document rather than unmodified originals. High only for the element count,
which is 15 in every admitted file with no other value observed. The 72 distinct
SHA-256 here sit against `SAV-1059`'s 70 admitted for a walk of the same tree
under that claim's narrower filter; the two-file difference is not traced, and
nothing in this claim turns on it.

## Generated mission documents: LOAD terminations

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1087 | Three recorded LOADs of three generated mission documents terminate at one instruction, after deserialization, on the resume route `SAV-948` names. | High / Medium / Unknown | ● active | [EXP-0396](../experiments/EXP-0396-generated-mission-load-crash/) |
| SAV-1088 | Armor, Shield and Weapon bind their definition entry with no range check, and the Weapon compact-record arm reads that entry's column 15 with no guard. | High / Medium / Unknown | ● active | [EXP-0396](../experiments/EXP-0396-generated-mission-load-crash/) |
| SAV-1089 | In game9204 the `L08027` termination is driven by one byte, the definition row 0 of the Weapon in a Human's `Unit+0x74` slot; none of 119 admitted original saves has a Weapon on row 0. | High / Medium / Unknown | ● active | [EXP-0396](../experiments/EXP-0396-generated-mission-load-crash/) |
| SAV-1091 | In game9220 the `L04760` termination is driven by one byte, the definition row 0 of Armor record 1133 in Human 259's equipment array; setting it to 15 lets LOAD complete. | High / Medium | ● active | [EXP-0397](../experiments/EXP-0397-load-crash-ladder/) |
| SAV-1092 | A mover block whose RotationSpeed byte `+0x0a` is 0 terminates the original with an integer divide by zero at `L09136` on the first turn of an ordinary move. | High / Unknown | ● active | [EXP-0397](../experiments/EXP-0397-load-crash-ladder/) |

### SAV-1087

- Application Error events 116552, 116554 and 116629 (PIDs 41340, 58168, 43768,
  image `942e9b72…7d03`) each carry exception `c0000005` at fault offset
  `0x0010e47e`.
- Their LocalDumps minidumps agree: EIP `L08027`, a load through the object pointer, a read of
  address `0x3c`; EAX `0x3c`, ECX `0`, EDX `0xf`, EBP `001ae558`; and the same
  nine return addresses `L09137`, `L09138`, `L09139`, `L09140`,
  `L09141`, `L09142`, `L09143`, `L09144`, `L09145`.
- Decoded call sites: `L03780`→`R0147`, `L01863`→`R0193`,
  `L09146`→`R0191`, `L01865`→`R0061`, opcode4 arm
  `L06672`→`R0131`, `L03739`→`R0059`, `L03240`→`R0669`, the
  `actor+0x74` arm `L04271`→`R0883`, and
  `L09147` (an indirect call through the slot at `+0x54`)→`R1033`, which is Weapon vtable
  `L04583 +0x54`. No frame is a `Serialize` routine.
- Each dump's captured memory holds exactly one UTF-16 save basename: game9200,
  game9202 and game9204 respectively.

**Confidence.** High for the fault site, registers and chain. Three dumps
sharing exception, EIP, non-heap registers and all nine frames exclude runs
faulting differently; every call site is decoded and none is a load arm, which
excludes a fault inside the deserializer. Medium for joining each process to its
document, which rests on memory residue, not a recorded file open.

**Unknown.** Any run outside the window 2026-09-22T20:00Z..2026-09-23T08:00Z.

### SAV-1088

- The load arms end `L04590`, passing L04589 (Armor) as `this`,
  `L04592`, passing L04591 (Shield) and `L04594`, passing L04593 (Weapon),
  each after a load of the byte at `[item+0x0c]` and followed by a call to `R0926` and a store to
  `item+0x3c`.
- `R0926` forwards to `L08026`, which computes `[coll+4] + idx*0x3c` and
  reads no size.
- Only class Item (`L04588`) calls `R0925` (returns `[coll+8]`), compares
  with `JLE`, then uses `L09148`.
- `R1033` (Weapon `vt+0x54`) executes `L09149`, pushing 0xf,
  `this = [item+0x3c] + 8`, a call to `R0476` (→`L09150`, `[this+4] + idx*4`) and
  `L08027`, loading through the returned pointer. An entry whose parameter array has a null data
  pointer faults there reading `0x3c`.
- Among Weapons rows 0..27 (runtime size 28, read from three dumps at
  `L04593`), row 0 is never serialized (`DAT-GRAM-003`) and row 23 `rem` has
  no parameter array; a row ≥ 28 indexes past the array.
- Also cites `ITEM-DEF-002`.

**Confidence.** High for the four arms, the three helpers and the faulting
routine, each read whole; this excludes a size compare inside the
Armor/Shield/Weapon lookup, which `L08026` does not contain. Medium that row
0's parameter array is empty, from `DAT-GRAM-003` plus the zero `this` in the dumps, not a
read of the entry.

**Unknown.** Whether another routine rejects an out-of-range row between LOAD
and first use, and what rows ≥ 28 resolve to.

### SAV-1089

- Setting the byte to 13 moves the fault; re-encoding alone does not.
- Decoding: `savdoc` with position-verified Token rows (row byte at Token head
  `+16`, `SAV-TOKEN-034`; `Unit+0x74` per `ITEM-EQUIP-006`).
  - game9204 (`b637aeaa…`) and game9200 (`3b0244c0…`) each carry 14 Weapons.
    Weapon record 805 on row 0 is the `Unit+0x74` object of Human record 259,
    with its row byte at decoded offset 821.
  - game9202 (`b12bc67b…`) carries 12 Weapons; record 804 is held by Human 258
    at offset 820.
  - Every other candidate Weapon names row 6, 7, 9, 12, 15 or 20, all with
    column 15 present.
- Owner run of the original EN install, one session:
  - control game9221 (`97fa465c…`, game9204 re-encoded with literal packets, no
    decoded byte changed) terminated at fault offset `0x0010e47e` (event 116636,
    dump rom.exe.3216, which names game9221);
  - game9220 (`9ad8caf4…`, decoded byte 821 changed 0→13, nothing else)
    terminated twice at `0x0010cb38` (events 116634 and 116638). Dump
    rom.exe.54776 names game9220; dump rom.exe.42792 names no save and is
    game9220's first attempt by owner statement.
- Both game9220 dumps share EIP, EAX `0x3c`, ECX `0`, EDX `0xf` and all nine
  frames. Frame 01 returns to `L09151`, after
  `L08824`, a load of the dword at offset 0x198 plus 4 bytes times the slot index, the Humanoid equipment array, instead of
  the `Unit+0x74` arm's `L09138`.
- `L04760` is `R0984` (Armor `vt+0x54`), with the same pattern of pushing 0xf,
  calling `R0476` and loading through the returned pointer. Each candidate also carries an
  Armor on row 0 (record 1133 in game9200 and game9204, 1131 in game9202). That
  next fault is outside this claim.
- Corpus: 119 admitted original saves (EN 23, RU 4, owner saves 92; generated
  paths, `SAV-1059`'s six candidates and the three digests excluded). The traces
  hold 2471 Weapon bodies. Their decoded row byte is in 2–26, none 0 and none ≥
  28, and it equals the Token `class_key` on all 2357 with a verified row. No
  original Armor or Shield Token row carries row 0.
- Of 1460 Human `Unit+0x74` steps, 429 are null and 1031 open a new Weapon (tag
  `ffff` or `0x8000|n`). The 1012 with a verified row fall on rows 20, 9, 7, 6,
  21, 13, 18, 3, 14, 19, 22 and 5.

**Confidence.** High for game9204's row byte as the driver of the `L08027`
fault. The pair excludes an environment cause, since the control, with
game9204's decoded body, faulted there in the same session, and a structural
misalignment, since a one-byte change inside an unchanged structure moved the
fault. Medium for game9200 and game9202, whose own documents were not
substituted and whose dump-to-file join rests on memory residue (`SAV-1087`).

**Unknown.** Which Human is `Player+0x34` at runtime, so which slots the resume
route reaches beyond the observed two.

### SAV-1091

- The Armor is array index 7 of Human 259's equipment array `actor+0x198+4i`.
- Events 116634 and 116638 and dumps rom.exe.42792 and rom.exe.54776 carry
  `c0000005` at `L04760`, a load through the returned pointer, in `R0984`, the dword at Armor
  vtable `L04581 +0x54`: pushing 0xf, `this = [item+0x3c] + 8`, calling `R0476`,
  the Weapon routine's shape, with EAX `0x3c`, ECX `0`.
- Frame 01 returns to `L09151` in the third arm of walker `R0669`, whose
  loop `i = 3..12` over `[actor + i*4 + 0x198]` stood at `[ebp-0x20]` = 7 in
  both dumps.
- In game9204 the only Armor, Shield or Weapon rows 0 are Weapon 805 (changed by
  `SAV-1089`) and Armor 1133 (decoded offset 1149).
- game9230 (`ff84e473…fe2b`) changed only offsets 821 and 1149 (0 → 13, 0 → 15).
  The owner's run of the original EN game reached the map, and the only event in
  its window is `c0000094` at `0014a2bf` (`SAV-1092`).
- Corpus: row 0 occurs on 0 of 3291 Armor records in the 119 admitted original
  files (byte-identical copies counted once per file).
- Also cites `SAV-1088`.

**Confidence.** High for the byte as the driver: game9230 differs from game9220
only in it, which excludes another Armor or field (R-B), and the same session
reproduced the document-driven `L08027` fault with game9221, which excludes
the environment. Medium that the actor at index 7 is Human 259: the dump's heap
pages are absent, and the join rests on the document holding the only row-0
Armor at that index.

### SAV-1092

- The generated document game9230 carries 0 in all 36 mover blocks.
- Event 116640 and dump rom.exe.8248 (naming game9230.sav): `c0000094` at
  `L09136`, a signed division by the divisor, which is `0` in the dump context (ECX `0`, ESI `0x40`, arc 64).
- `L09152` loads the mover from offset 0x154 of the object, and the divisor is `[mover+0xa]`.
- Return addresses `L09153`, `L07004`, `L09154`, `L04570` follow calls
  to the turn caller `R0056`, the step routine `R0054`, the move routine
  `R0178` and the order-progress routine `R0016`.
- Over 4044 mover blocks in the 119 admitted original files (byte-identical
  copies counted once per file) the byte is 0 in one, a whole-zero block of an
  original resave of a generated document. Humans carry 15..22 and Units 8..22.
- game9231 (`7a0af741…c597`) set it to 17 on 17 Humans and 16 on 19 Units and
  changed nothing else. Its session and the resave session, with player-ordered
  movement through the quest, did not terminate.
- Also cites `MOVE-TURN-044`, `SAV-TURNLOAD-822`.

**Confidence.** High for the byte as the driver: game9231 differs from game9230
only in the 36 bytes and ran with movement and no fault, which excludes an actor
with no document mover and a later derive rewriting the byte from a zero speed
word.

**Unknown.** Which actor faulted; the heap pages are absent.

## Generated mission-10 documents: state the original needs after LOAD

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1093 | Mission-10 buildings are drawn after LOAD when the document gives each Building publication mask 2 at Token+0x18, its type in the type word at Token+17, and the originals' 1,769 border and footprint block rows. | Medium / Unknown | ● active | [EXP-0397](../experiments/EXP-0397-load-crash-ladder/) |
| SAV-1094 | A mission-10 document whose campaign `+0x110` (`AutoGetMission`) is 0 leaves the party on the world map after the quest; with 20 the second mission starts. | Medium / Unknown | ● active | [EXP-0397](../experiments/EXP-0397-load-crash-ladder/) |
| SAV-1095 | Hero control on the first LOAD is decided by the document, not by the process: original-written resave game0003 and game9233 give control in a freshly started game, game9232 does not. | Medium / Unknown | ● active | [EXP-0397](../experiments/EXP-0397-load-crash-ladder/) |
| SAV-1096 | A generated mission-10 document with zero AI start state lets the witch fight at mission start; with the AI actor and group start fields set to the mission-start original's values and diplomacy template {1,1,0,0}, she does not. | Medium / Unknown | ● active | [EXP-0397](../experiments/EXP-0397-load-crash-ladder/) |
| SAV-1097 | The four block rows `20/20` (static/dynamic) at the sack cells drive two symptoms of a generated mission-10 document: trigger T05 firing at LOAD and the pick-up click being refused. | High / Medium / Unknown | ● active | [EXP-0397](../experiments/EXP-0397-load-crash-ladder/) |
| SAV-1098 | Unit collision after LOAD needs the block rows at unit cells: the generated mission-10 document has none, and adding a row for every cell record by the original rule gives the hero collision. | High / Medium / Unknown | ● active | [EXP-0397](../experiments/EXP-0397-load-crash-ladder/) |

### SAV-1093

- Every original mission-10 save carries the 1,769 border and footprint block
  rows. The generator wrote mask 0, the type at Token+16 and 6,400 zero-based
  rows.
- game9231 and its three original resaves (game0000..0002, written by the
  original's SAVE) left structures missing. The resaves repaired Token+0x18 to 2
  on 36 Humans and Units and 4 Sacks but kept 0 on all 18 Buildings. Every
  verified original Building token carries 2 and a nonzero type word.
- The 6,400-row array (dynamic = static ∈ {0,1,3}) overwrites the ingested
  border `0x1f` and every footprint. All 33 original 10.alm files (14 distinct
  documents) carry one identical set of 1,657 border rows (`1f/1f`,
  static/dynamic) and 112 footprint rows (`25/25`).
- game9232 (`9a51624e…bdca`) changed these fields, together with the control,
  passability and progression groups; the owner saw the buildings drawn.
- Also cites `SAV-678`, `SAV-BLOCK-011`, `TERR-PASS-053`.

**Confidence.** Medium. The three structure changes and the other groups were
made together in one document, so the run does not say which field drives the
drawing, and the mask reading is a corpus contrast.

**Unknown.** Whether water became impassable: the passability group (mover mask
`+0x05` 0 → 65/68, the same block rows) was never reported by an owner run.

### SAV-1094

- In the game9231 session and its resaves (field 0 throughout, passed through by
  the original's SAVE) the quest completed and the world map opened with nothing
  further.
- `+0x110` is 20 on 9 of 9 original EN mission-10 saves.
- game9233 (`beb11604…3d38`), carrying the round-3 change 0 → 20, completed the
  quest, opened the world map and started mission 20 (owner run).
- Also cites `REG-SCN-063`.

**Confidence.** Medium. game9233 differs from game9231 in many groups, so the
run supports the field without isolating it; the mechanism is `REG-SCN-063`'s
reading, not traced here.

**Unknown.** What mission 0 resolves to.

### SAV-1095

- game9232 (`9a51624e…bdca`) gave no control on the first LOAD from the menu and
  gave control after an original SAVE and re-LOAD in the same session.
- game0003 (`b1c2f862…648e`), that SAVE, loaded from the menu of a fresh
  process, gave control.
- game9233 set the human participant's subtree to game0003's values and gave
  control on the first LOAD. The group:
  - hero Token sub-cell `+4/+5` 37,37 → 128,128;
  - the hero's order object pending move (order 1, progress 3) cleared;
  - the `+0x178` word list cut from 3 entries to 0;
  - the 462-byte and 19-byte blocks;
  - Player body `+57` `0x32` → 0;
  - group AI `+0x48` 0 → 1;
  - session wire 1400..1415.
- Token+0x18 = 2 on every Human and Unit (game9232) was not sufficient.
- No single one of these bytes is out of population against the original corpus.
- Also cites `SAV-678`, `SAV-OWNER-048`.

**Confidence.** Medium for the group as the driver: two documents with the group
give control from a fresh process and the document without it does not, which
excludes a process-only cause.

**Unknown.** Which byte of the group drives it; no run changed one byte alone.

### SAV-1096

- The fix set: each AI actor's post, 19-byte block, order and mover bytes and
  each group's AI block set to the mission-start original's values, and the
  diplomacy template {1,1,0,0}.
- In game9232 the 35 AI actors carry post 0,0 (order object `+0x00/+0x01`),
  19-byte block `+4` = 0, order `+0x14`/`+0x71` = 0, mover `+0x08/+0x09` = 0,0,
  `+0x72` and `+0x82/+0x83` = 0. The 18 AI groups carry centre and guard radius
  0 (`+0x28/+0x29`, `+0x2d`). Session wire 1848..1851 is 0.
- Every original mission-10 document sets them; the three mission-start saves
  carry each actor's own cell as post and 5,255 at mover `+0x08/+0x09`.
- Owner runs:
  - the game0003 state, which carries these zeros (the original's SAVE of the
    game9232 session, re-LOADed in that session and loaded fresh as game9234),
    shows the witch attacking a bee, the bee not answering, and clubmen killing
    the witch;
  - game9233 (the fields copied from RU `game9999.sav`, 837 bytes) and the
    original mission-10 start show no fight.
- Also cites `AI-POST-042`, `AI-GRPGUARD-074`, `AI-DIPLO-083`.

**Confidence.** Medium. The whole set was copied at once and the mechanism is
not traced; the witch is identified by class key 200 and cell 36,51 only.

**Unknown.** Which field drives it.

### SAV-1097

- 10.alm trigger T05 is `IF Get sack(38,64) == 0 THEN Send message 17`; mission
  10's `event17.txt` is one part of 78 characters, the length of the owner's
  quoted text.
- The check calls `R0955`, which rejects a cell whose `map+0x10000` byte
  lacks bit 0x20 (`TRIG-SACK-022`). The pick-up order 0x21 calls the same lookup
  and refuses before any walk (`ITEM-PICK-016`). Row byte 0 is that plane
  (`SAV-BLOCK-011`).
- In all 33 original 10.alm files (14 distinct documents) the 99 sack-record
  cells (45 distinct document-cells) carry `20/20`.
- The generator's output (game9200, game9202, game9204) carries `00/00` there;
  game9232 and game9233 carry no row; game0003, the original's resave of the
  game9232 session, carries no row there. So the running original is read as not
  setting the bit for a loaded cell record.
- game9236 (`4d997252…84a3`) is game9233 plus exactly these four rows. In the
  owner's run of the original EN game, message 17 did not appear at LOAD, a
  click on a sack walked the hero to it, and the hero still passed through
  units.
- The owner later reported, for the round-5 documents without separating
  game9236 from game9237, that a sack click walked the hero to the sack and
  picked it up, and message 17 appeared when the sack at 38,64 was picked up.

**Confidence.** High for the four rows as the driver of both symptoms: game9236
changed only them, which excludes the sack object ids (Token+12, reversed
against the RU start) and the session slots and latch 0, all unchanged. Medium
for the route through `R0955`'s bit test, which is read from the
published rows and not traced in the run; message 17 appearing on the pick-up at
38,64 agrees with T05 now finding the sack. Medium that the original's LOAD does
not set static 0x20 for a loaded cell record: one resave contrast (game0003),
with the cell-table load arm not traced.

**Unknown.** Why a pick-up succeeds without the bit when the hero already stands
on the sack's cell and the pick-up is issued quickly (owner observation; walking
over a sack does not pick it up). Order 0x21 (`R0955`), session source 3
(`R0446`, `ITEM-CMD-007`) and the tick arm at `L00009` all reach the
bit-tested lookup; the command handler `R0430`, the other caller of the
looting routine `R0448` (`ITEM-PICK-009`), was not read.

### SAV-1098

- `TERR-PASS-051`: a cell blocks a mover when its byte and the mover mask share
  a bit. The ground mover masks 0x41 and 0x44 carry 0x40; the flier mask 0x82
  does not.
- RU `game9999.sav` carries dynamic bit 0x40 on 36 of 36 actor cells; game9233
  and the generator's output carry it on none.
- Over the 33 original 10.alm files (14 distinct documents), counting each file,
  2,038 record cells follow `static = payload+1 | 0x20`,
  `dynamic = static | 0x40` with a ground occupant, 3,630 carry footprint rows
  `25/25`, and 298 structure-only records break the rule.
- game9237 (`c0c061c3…6e9e`) is game9233 plus the 72 missing rows by that rule
  and two cell records moved to the actor standing on them (Human 99, 63,20 →
  64,21; Human 181, 66,13 → 67,12). Its rows equal RU `game9999.sav`'s except at
  the hero's cell and one empty record at 63,20.
- In the owner's run units blocked the hero, control worked, a sack click walked
  the hero to the sack, and nothing terminated. The owner later reported, for
  the round-5 documents without separating game9236 from game9237, that a sack
  click walked the hero to the sack and picked it up, and message 17 appeared
  when the sack at 38,64 was picked up.
- game9236, with only the sack rows, kept no collision.
- Also cites `SAV-TERRKEY-056`, `TERR-CELLREC-146`.

**Confidence.** High that the game9237 change set, not the hero's own bytes,
restores collision: game9237 left the hero's mover `+0x08/+0x09` (0,0 against
the originals' 5) unchanged. Against game9236 that set is 68 rows plus 5
cell-record bytes of the two moved records. Medium that the rows and not the two
record moves carry it, and for dynamic bit 0x40 as the specific bit: the rows
also set static 0x20 on those cells, and no run separated either. Medium that
the original's LOAD does not set dynamic 0x40 for a loaded cell record: game0003
carries 0x40 at 11 of 36 actor cells, all actors the AI moved in that session.

**Unknown.** Whether the missing rows also explain the bee not answering the
witch in game0003 (`AI-336` tests static 0x20); no round document could show it.

## Generated mission-10 document LOAD state

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1099 | Campaign `+0x120`, the first-MapPoint flag, is set on 22 of 54 organically numbered original mission saves in the dated corpus, and the selected mission alone does not predict it. | High / Medium / Unknown | ● active | [EXP-0398](../experiments/EXP-0398-generated-mission-load-state/) |
| SAV-1100 | In an owner run on three substitution documents, Player 0's body byte and its AI group's block are each alone insufficient for hero control on the first LOAD; the hero/session subgroup of about 25 bytes suffices. | Medium | ● active | [EXP-0398](../experiments/EXP-0398-generated-mission-load-state/) |
| SAV-1101 | In an owner run on two substitution documents, the witch's own 16-byte personal AI-start record, not her owner-1 AI group's 12-byte block, is what stops her mission-start fight. | High | ● active | [EXP-0398](../experiments/EXP-0398-generated-mission-load-state/) |
| SAV-1102 | Bit 5 (`0x20`) is in none of the three constructor-set movement masks (`0x41`, `0x44`, `0x82`), and dynamic bit 6 composes both ground masks; an indirect effect of static bit 5 on collision is not excluded. | High | ● active | [EXP-0398](../experiments/EXP-0398-generated-mission-load-state/) |
| SAV-1103 | Ingest blocks water by a hardcoded terrain-class-8 test, not the registry `Pass*` values, but LOAD writes saved block rows raw over both planes, so a saved row on a water cell could override that block. | High / Medium / Unknown | ● active | [EXP-0398](../experiments/EXP-0398-generated-mission-load-state/) |
| SAV-1104 | A mission-10 document's `AutoGetMission = 0` takes the `!= -1` fork and calls the routing loader with request 0, the documented no-op; what `R1411` does with that zero return is unread. | Medium | ● active | [EXP-0398](../experiments/EXP-0398-generated-mission-load-state/) |
| SAV-1105 | On a normal mission LOAD the Player/client list is filled by a pure archive copy-in with no selection step in the traced LOAD bodies; what picks the restored Player for the first mission command is Unknown. | Medium | ● active | [EXP-0398](../experiments/EXP-0398-generated-mission-load-state/) |
| SAV-1106 | On a fresh-process mission LOAD the 100-slot register array is restored wholesale, then the binder writes one slot per opcode-`0x10002` check; an in-session reLOAD skips the restore. | Medium / Unknown | ● active | [EXP-0398](../experiments/EXP-0398-generated-mission-load-state/) |
| SAV-1107 | `Group+0x20` is a word list of the shared class `SAV-EMBED-039` names; a direct-call census of its mutation entry points finds no site targeting it, 0 of 567 corpus Group records carry a nonzero count, and element meaning is Unknown. | High / Medium / Unknown | ● active | [EXP-0398](../experiments/EXP-0398-generated-mission-load-state/) |

### SAV-1099

- `SAV-CAMPAIGN-076` names campaign `+0x120` the first-MapPoint flag.
- `census.sh` runs `tools/savcampaign -state` over every
  `gameversions/saves/<date>/game0*.sav` one directory level below the corpus
  root: 54 files across seven dated snapshots. Every experiment or story
  subdirectory is excluded as synthetic test material, and `game9999.sav` as
  the tavern-balance save rather than a mission save.
- 22 files carry `first_point = 1`, spanning five of the seven dates
  (`evidence/q1-census.tsv`).
- `2027-09-07_game0005.sav` and `2027-09-07_game0029.sav` share
  `selected_mission = main_mission = 30` with opposite flag values, 0 and 1, so
  `selected_mission` alone does not predict the flag.

**Confidence.** High for field identity and order: that is
`SAV-CAMPAIGN-076`/`EXP-0232`'s complete-serializer reading, not this
experiment's finding. Medium for the census and the selected-mission
independence: a bounded, dated snapshot of 54 files, reproduced once into a
second scratch directory byte-identical, but not a byte-reproducible generator,
since the corpus grows.

**Unknown.** The writer and the reader of `+0x120`: no existing claim traces
the instructions at that offset, and tracing them needs a Ghidra project built
from the owned `rom.exe`, which this experiment did not build.

### SAV-1100

- The run narrows `SAV-1095`'s control-group Unknown, 27 bytes plus one cut, to
  the hero/session subgroup of about 25 bytes.
- game9243: game9204 plus every established fix plus only Player 0 body `+57`
  (`d6b76ed3…ae58f`).
- game9244: game9204 plus every established fix plus only the group's own AI
  block `+0x48` (`1c09a247…7a053`).
- game9243 and game9244 each left the hero unselectable on the first LOAD from
  the menu. The hero became selectable and answered a move click only after an
  in-session original SAVE and re-LOAD. This reproduces game9232's pre-fix
  fingerprint (`SAV-1095`).
- game9245: game9204 plus every established fix plus every other control-group
  byte (hero Token head, `+0x178` block, 19-byte block, word-list cut, session
  wire `+0x08/+0x10`), but not `+57` or `+0x48` (`642a313b…9275a97e`). It gave
  the hero control on the first LOAD: the owner reported "hero visible", the
  owner's recurring phrase for the controllable state in earlier rounds.
- Exactly one of the three isolated groups produced control. That was the
  pre-registered outcome that narrows the driver to that group and excludes the
  other two individually.
- Evidence file: `evidence/owner-testimony.md`.

**Confidence.** Medium. A clean three-way partition of `SAV-1095`'s byte set
shows the hero/session subgroup sufficient and the other two individually
insufficient. Control on game9245 is read from the owner's "hero visible"
report by the same convention as earlier rounds, not from a separately
confirmed selection or move-click observation; no run reported a successful
click on game9245.

**Unknown.** Which byte inside the ~25-byte subgroup drives control: no run
isolated a single byte, so `SAV-1095`'s Unknown is narrowed, not closed.

### SAV-1101

- game9246: game9237 with only the witch's own personal AI-start record
  reverted to its pre-round-4 value, the group block left fixed
  (`ffa970ec…4fb1de0`). The witch fought.
- game9247: game9237 with only the group block reverted, the personal record
  left fixed (`17410ab4…10fd757c1`). She did not fight.
- Breaking the personal record alone, with the group fixed, reproduces the
  fight; breaking the group alone, with the personal record fixed, does not.
  The fight's absence tracks the personal record's state in both directions,
  independent of the group block's state.
- This closes `SAV-1096`'s Unknown.
- The owner kit's pre-registered refutation table paired this observed
  combination with the opposite conclusion, naming the group block as driver.
  Re-deriving which record group each candidate reverts corrects this to the
  personal record, and the source table is corrected to match: a
  pre-registration error, not a reinterpretation of new evidence
  (`evidence/owner-testimony.md`).

**Confidence.** High. Each candidate reverts exactly one of the two record
groups, and the fight tracks the personal record's state in both directions
with the group held at the opposite state each time: a discriminating result
between the two named alternatives, not corpus agreement. The actor's identity
(a Human at cell 36,51, owner ordinal 1) is `SAV-1096`'s identification, not
independently reconfirmed here.

### SAV-1102

- `TERR-PASS-051`: the constructor `R0206` leaves the default mask
  `0x41`, and `R0471` maps movement domain 1/2/3 to `0x41`/`0x44`/`0x82`.
  None of the three contains bit 5; bit 6 composes `0x41` and `0x44`.
- A LOADed mover's mask byte (`mover+5`) is restored raw from the document
  (`SAV-ACTORINPUT-547`), not re-derived by `R0471` on the LOAD path. The
  fact therefore covers the constructor's default population, not the byte a
  loaded document carries.
- Bit 5 gates the cell's hash-table record: `R0035` (`AI-336`,
  `TERR-CELLREC-146`) returns null when static bit 5 is 0.
  `TERR-STRUCT-069`'s `R0453` recomputes dynamic bit 6/7 from that
  record's occupant slots when a record exists.
- `TERR-PASS-051` also documents one of 13
  sites testing the byte at `[cell+0x10000]` against 0x20 inside `R0053`, the
  route-search driver that `MOVE-AREA-038` reads as containing no plane-bit test
  by immediate. The two readings of that function are not reconciled here.
- This narrow fact does not by itself answer Q2c.

**Confidence.** High only for the narrow fact: no constructor-set mask contains
bit 5, and bit 6 composes both ground masks. Q2c's empirical question, whether
dynamic `0x40` alone restores collision, stays at `SAV-1098`'s Medium; this
claim does not answer it, and the two moved cell records are not addressed
here.

**Unknown.** Whether static bit 5 affects collision indirectly: a route from
bit 5 to the movement-blocking bit through the record's existence is not
excluded by this experiment.

### SAV-1103

- `ALM-TERR-043`: water impassability at ingest is hardcoded (terrain class 8,
  raw water-bit test) and never derived from the registry's ten `Pass*` scalars
  (0 of 10 key strings in either binary).
- `TERR-PASS-053`, `SAV-LOAD-057`: LOAD constructs terrain from the `.alm`
  first, then unconditionally applies every saved block row's static/dynamic
  byte raw over both planes at `R1360`. A saved row covering a water cell
  with static bit 0 clear would make that cell passable at runtime, overriding
  the ingest-time block. `SAV-1093` shows generated rows overwriting the
  ingested border the same way.
- The kit candidates game9243-game9245 of this experiment's owner run carried
  substituted block rows (`--blocks-file rows9237.txt`), which were not checked
  against the coastline cells the owner tested. Water blocking in those
  documents is therefore not shown to be independent of the substitutions.
- `TERR-WATERBOUND-164` is an ingest-time classifier (`L09155`/`L09156`,
  reached from `R0278`'s ingest tail per `SAV-LOAD-057`), not a runtime
  one.
- Owner run: on game9245 (`642a313b…9275a97e`) the hero did not enter water at
  the coastline. Water was not tested on game9246 or game9247
  (`evidence/owner-testimony.md`).

**Confidence.** High only for "ingest sets water blocking from a hardcoded
test, and the registry `Pass*` values play no part" (`ALM-TERR-043`). Medium for
the game9245 owner observation: one corroborating instance, not a check of
whether that document's substituted rows cover a water cell.

**Unknown.** Whether a saved row overriding a water cell's block bit occurs in
any shipped or generated document. `SAV-BLOCK-011`'s residual ("the 9 cells
where the save clears bit 0 the map sets") was not cross-checked against
`10.alm`'s water cells, and this experiment's kit files were not checked
either; water was not tested at all on game9246/game9247.

### SAV-1104

- `REG-SCN-063` traces the fork. `R1297` writes campaign `+0x110`
  (`AutoGetMission`), registry default `-1`. Its sole reader `R1411`
  branches:
  - `!= -1`: call `R0755(that mission)`, latch the announce, travel to
    the new mission's map object;
  - `== -1`: return 0 and send the party to the city, global-map object 0.
- `SAV-CAMPAIGN-078` traces `R0755(request)`: a `request` lower than the
  current main-progress record returns 0 immediately, a silent no-op that
  writes nothing.
- `AutoGetMission = 0` takes the `!= -1` arm (`0 != -1`) and calls
  `R0755(0)`, the documented no-op against any nonzero current record.
  That much composes from the two cited claims.
- The owner's observation for `SAV-1094`'s game9233, "the world map opened with
  nothing further", does not distinguish travel to the current mission's object
  from the city/object-0 arm: both leave the party on the world map with
  nothing new started.
- This narrows `SAV-1094`'s Unknown but does not close it.

**Confidence.** High for the routing fork and the no-op call, each on its cited
claim. Medium for this claim's composed reading: the step it needs, what
`R1411` does with a zero return, is an inference, not a read. The
city/object-0 arm is the live alternative this experiment does not exclude.

**Unknown.** Neither cited claim reads how `R1411` handles a zero return
from `R0755`: whether it tests the return value (for example routing to
the city/object-0 arm the `== -1` branch uses) or unconditionally falls through
to the announce-latch/travel step using the unchanged current record.

### SAV-1105

- `SAV-PLAYERIDENT-830` traces the plain Player-manager LOAD path: document
  LOAD calls `L08155`; `R1558` reads each Player reference through
  `CArchive::ReadObject` (`L08109`), then appends the returned object through
  `R0417`. "No active-player lookup selects this copy's destination."
- `SAV-919` corroborates from the tavern-key side: the same manager/list LOAD
  path (`R1372 -> R1558 -> L08109 -> R0417`) "contains no direct
  call" to the non-LOAD roster-registration routine `R0424`.
- `AI-FORMACTIVE-317` is the one documented mechanism that selects an active
  Player by name match: an existing-Player search (`L00181`-`L00182`) that
  retains the first accepted `CString` name match. It scopes itself to the
  return-to-game/client-description route, not a plain single-player mission
  LOAD, and grades its "composition after mission LOAD" reading only Medium.

**Confidence.** Medium. Two independently derived High readings of the same
LOAD path agree on "no selection inside the traced bodies". No claim traces a
call from the mission-command dispatcher, the routine issuing the first
post-LOAD command, to the name-match mechanism or to any equivalent specific
to the plain single-player LOAD route.

**Unknown.** What picks "the" restored Player for the first mission command.
Whether it is simply the first (or only) list entry by convention, rather than
the result of any search, is untested by any claim in hand.

### SAV-1106

- A fresh process is one whose global session pointer `[L00004]` is still
  null.
- `SAV-710`: the session self-pointer `session+0xa9c0` is repaired, not
  derived, by a two-instruction thunk (`R1485`) called from `L07741`.
- `SAV-711`: document LOAD reaches the session `Serialize` restore at
  `L07069` only while `[L00004]` is null; `L08640`/`L08641` skip
  construction and restore otherwise. An in-session resave-and-reLOAD, such as
  the route Q2a/Q2b's owner run used for control, takes the skip.
- `SAV-711` also: after the Sack read (`L07781`), LOAD reaches the
  authored-check binder `R0067`'s opcode-`0x10002` arm at `L01841`.
  Once per such check it writes one compiled map constant into the slot at the
  binder's advancing cursor and advances the cursor.
- `SAV-651` (active): the same 100-slot array (`session+0xbd34`) is read and
  written generically by the check evaluator `R0433` and by the pattern
  evaluator, at a slot index supplied by the check's or pattern's compiled
  data. Its attribution of the `R1485` call path to a per-kind rebuild is
  withdrawn; the slot-index and generic-evaluator material stands.
- `TRIG-STORE-002` and `MISSION-SLOT-008` bound a check's compiled subscript:
  on the shipped corpus's largest map (64 compiled checks), no check mechanism
  reaches slot 93.
- `SAV-652`'s original re-derivation claim, that specific slots' values are
  re-derived from other live objects' fields, is retracted in full: the cited
  `R1485` call chain reaches no per-kind field-copying helper, and no
  re-derivation step is established anywhere on the traced LOAD path. This
  claim relies on the corrected, active `SAV-652` instead: only the LOAD
  boundary, the wholesale restore preceding the binder's opcode-`0x10002`
  write.
- `SAV-650`: slot 93 is a per-tick counter incremented immediately before each
  evaluator call, a runtime cadence, not a LOAD-time population step.

**Confidence.** Medium. The cited claims compose, without new re-derivation,
into one LOAD-time state description.

**Unknown.** Which slot or slots `10.alm`'s trigger T05 reads, and whether
T05's check data uses opcode `0x10002`: no claim traces T05's compiled slot
index or opcode. `SAV-711`'s stated scope also remains Unknown: complete LOAD
execution, arbitrary indices, all callbacks, and whether any shipped map's
build actually contains a rejection.

### SAV-1107

- `SAV-EMBED-039` places `Group+0x20` as one of five sites sharing one class
  (constructor `R1545`, vtable `L08172`, `Serialize` `R1546`): the same
  class as `Unit+0x15c`, `Unit+0x178`, the order-embedded active ring and
  `Group+0x4c`.
- `SAV-632` traces the class's shared append primitive (`R0052`) and
  teardown (`R0439`).
- `SAV-GRPLIST-807`, the direct producer/consumer census, finds 18 distinct
  external call sites on the class's two mutation entry points, none targeting
  `Group+0x20`: the patrol-path loader targets `Group+0x4c`; four sites target
  the order-embedded active ring; `Serialize`'s own load arm; two unidentified
  singleton fields; five sites target an unrelated sight-search ring.
- Only the Group constructor's in-place construction and `Group::Serialize`'s
  vtable dispatch touch `Group+0x20` directly: construction and archive
  round-trip, not population.
- 0 of 567 Group records (`SAV-GRPORDER-806`'s census) carry a nonzero element
  count there.
- The word list Q2a's kit touches is a different field of the same class,
  `Unit+0x178` on the hero's record (`SAV-1095`, `SAV-1100`), populated with
  real entries and cut as part of the control group. `Group+0x20` is not that
  list and stays at zero across the observed corpus.

**Confidence.** High for the census completeness and the corpus
corroboration, the component mechanisms. Medium overall, since the null result
composes two instruments, each with a named residual scope.

**Unknown.** Element meaning, outright, not merely unread: `SAV-EMBED-039`
names this Unknown, and the corpus has no populated instance of this field to
read semantics from. A computed-call or bulk-copy mutation path is excluded by
neither of `SAV-GRPLIST-807`'s two instruments.

## Tavern roster inputs in a SAV

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1111 | The mission-130 owner SAV names what the tavern offers but not the order of its mercenary cells. | High / Medium | ● active | [EXP-0400](../experiments/EXP-0400-tavern-roster-presentation/) |
| SAV-1112 | The persisted `InnNPC` order is the talk-only cells' order. | High | ● active | [EXP-0400](../experiments/EXP-0400-tavern-roster-presentation/) |

### SAV-1111

- Campaign record (`tools/savtown`): `Mercenaries`
  `14,6,10,13,4,8,7,3,1,9,2,12,5`, unlocks `1..10,12,13,14`, working and
  pristine pools `1,1,1,1,1,4,4,3,3,3,0,3,4,3,0` and `InnNPC` `[25]`.
- Actor walk (`tools/savunit`): five `Human` records, the party, and no
  mercenary stock object.
- The 13 eligible types and one candidate fix the 14 cells. Their mercenary
  order comes from the ids of stock units the original builds after load
  (`TAVERN-ORDER-015`), which this file does not store.
- Evidence files: `evidence/sav-fields.tsv`, `evidence/sav-actors.txt`. Also
  cites `SAV-CAMPAIGN-076`.

**Confidence.** High for the decoded fields and the absent stock objects.
Medium that no SAV field influences the stock ids: the id allocator's state
after load was not traced.

### SAV-1112

- The inn view appends `R0741(InnNPC[j])` for `j` in array order
  (`L08032`..`L09157`) and paints element `j` at roster position
  `mercCount + j`.
- The talk handler reads `InnMission` and `InnNPC` at `sel − mercCount`
  (`REG-118`'s subscript with its bound corrected by `TOWN-468`).
- Removing an accepted pair (`SAV-CAMPAIGN-081`) moves every later talk cell
  back one position.
- Evidence file: `evidence/listing.txt`. Also cites `SAV-CAMPAIGN-076`.

**Confidence.** High: every step is a named instruction. Rival excluded: an
order sorted by npc id or mission, for which no sort exists between the array
read and the append.

## Original resave of a between-missions SAV

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1113 | In one original resave, the Group AI `+0x4c` and Order `+0x90` dwords hold raw heap-pointer bits. LOAD overwrites them without reading them, so 0 and any other value load alike. | High | ● active | [EXP-0401](../experiments/EXP-0401-town-resave-fields/) |
| SAV-1114 | Item Token `+0x08` is a pending pickup-announcement flag. `R0883` publishes it as record bit `0x40` and clears it, and the client posts a "Picked up" line for it on the carried-container arm. | High / Medium | ● active | [EXP-0401](../experiments/EXP-0401-town-resave-fields/) |
| SAV-1115 | In one original resave, 35 Item Token `+0x08 = 1` words came back 0. Token LOAD and SAVE copy the word literally, so a publication between LOAD and SAVE cleared them. | High / Medium | ● active | [EXP-0401](../experiments/EXP-0401-town-resave-fields/) |
| SAV-1116 | For the four Humans of one mission-140 SAV pair, the derive speed formula reproduces the saved speed word and mover byte. The pair cannot show whether the original re-derives after LOAD. | High / Medium / Unknown | ● active | [EXP-0401](../experiments/EXP-0401-town-resave-fields/) |
| SAV-1118 | After LOAD, session entry sends every tick-list actor with a full mask. The client state arm creates a drawable for an unknown id whose message passes the owner, class and `0x4001` tests, at the message position. | High / Medium / Unknown | ● active | [EXP-0402](../experiments/EXP-0402-loaded-actor-drawing/) |
| SAV-1119 | The actor-state sender names the Token Position only for a centred sub-cell, else the mover cell `+06`. An actor mid-step at the entry send gets its drawable created at the mover cell; the simulation keeps its Position. | High / Medium / Unknown | ● active | [EXP-0402](../experiments/EXP-0402-loaded-actor-drawing/) |
| SAV-1120 | In `game9271.sav` troll 57 is mid-step with mover cell 0,0: unless a tick moves it before the entry send, its drawable is created at 0,0. Troll 168 is centred; its drawable is created at 62,62. | Medium / Unknown | ● active | [EXP-0402](../experiments/EXP-0402-loaded-actor-drawing/) |
| SAV-1121 | Over 76 distinct original-written documents (100 admitted files) of the owner save tree, 69 mid-step actor records all have a mover cell at most one cell from their Position. | Medium / Unknown | ● active | [EXP-0402](../experiments/EXP-0402-loaded-actor-drawing/) |

### SAV-1113

- Input population: one Againrom-written between-missions SAV at main mission
  140 and the owner's original resave of it. Each holds one Player and four
  Humans.
- Body `0xdd` is byte `0x4c` of the first Group's AI raw80. The Player name
  starts at `0x51`, the Group count sits at `0x8b`, the Group list Count at
  `0x8f` and the raw80 at `0x91`. The field is AI `+0x4c`, the pointer to the
  AI word list.
- Human record start `+0x1f7` is byte `0x90` of the Order raw148, which begins
  at record `+0x167` when the Effect list and both route lists are empty. They
  are empty in all four Humans. The field is Order `+0x90`, the pointer to the
  order word list (`SAV-EMBED-039`).
- STORE writes each pointer inside its raw block (`L00550`, `L09158`). It
  then serializes the list through that pointer.
- LOAD deletes the constructor's list and reads the raw block. It then
  allocates a fresh `0x1c`-byte list (`L09047`, `L09159`), constructs it
  when the allocation is nonnull (`L09048`, `L09160`) and stores the result
  at `+0x90` or `+0x4c` (`L09049`, `L08203`) before loading the elements. The saved bits are never dereferenced (`SAV-1066`,
  `SAV-GRPAI-563`).
- The resave values are `0x0304e4b0`, `0x0304e6b0`, `0x030114c0`, `0x03051360`
  and `0x03052960`. Each occurs once in its body. None equals any of the
  body's 120 identity keys: 119 Token heads and the Player. They are the
  addresses of the lists the original held at SAVE.
- Evidence files: `evidence/q2-pointer-words.tsv`,
  `evidence/diff-attribution.tsv`, `evidence/listings.txt`.

**Confidence.** High. Two alternatives are excluded. A saved reference is
excluded because both LOAD arms overwrite the word before any read and no
value matches an identity key. A computed scalar is excluded because STORE
dispatches the list's `vt+8` through the same word (`L09044`..`L09046`,
`L09161`..`L09162`) and LOAD stores the new list's address there. The weakest input is the record offsets,
which rest on the savdoc programme walk.

### SAV-1114

- `R0883` sets compact-record byte `+4` bit `0x40` when Item Token `+0x08`
  is nonzero. It then stores `+0x08 = 0` (`L04873`..`L04874`). It does this
  for every record it writes. Its 10 direct call sites lie in `R0669`,
  `R0670`, `L09163` and `L09164`.
- The located setters are the pickup stamp `L04886` (`= 1`) and the merge OR
  `L04883` (`ITEM-GROUNDMOVE-130`, `ITEM-MERGE-129`).
- Producer `R0670` writes `msg+0x9 = 0x76` and `msg+0xc = 2`, with the count
  at `msg+0xd`. The client loop that reads records from `msg+0x13` (`L09165`)
  copies byte `+4` to display `+8` (`L09166`). On every record, whether or not
  the entry was appended (`L09167`, `L09168`, `L04822` all reach
  `L04823`), it tests `+8 & 0x40` (`L04949`) and clears the bit
  (`L04951`). It then posts
  global string 85 and the item name to the message line for `0xbb8` ms
  (`L04961`..`L04948`). When the count exceeds 1 it also posts strings 86
  and 87 with the count.
- EN string 85 is `Picked up`. RU string 85 is `Вы подняли:`.
- The twelve-slot equipment arm (`HERO-APPEAR-048`) carries the bit, and the
  search below found no message test for it.
- EXP-0256's anchor labels `wire-flags-equipment` (`L09166`) and
  `wire-flags-inventory` (`L09169`) are swapped: `R0670` sends subtype 2,
  which the client dispatch `L09170` routes to `L09171`.
- `R0883` also calls the item's virtual `+0x54` at `L09147`, which sets bit `0x20` of the same byte for the star display (`ITEM-STARFLAG-096`), and sets or clears bit `0x80` from `item+0x14` (`L09172`..`L09173`). The two claims name different bits of byte `+4`, written by one function, and do not conflict.
- Evidence file: `evidence/listings.txt`. Also cites `ITEM-STARFLAG-096`.

**Confidence.** High for the named writer, the clear, the client test and the
string selection. Each is an instruction in the listing. The clear-and-post
sequence at `L09174`..`L04948` excludes the reading that bit `0x40` is
only a display flag. Medium for the absence of other writers and readers of
Item `+0x08`. The search covered three populations:

- the 11 direct calls of container iterator `R1651`;
- the 10 direct calls of `R0883`;
- client instructions below `L06029` that test or mask `0x40` within four
  instructions of a `[reg+8]` operand, which gave one hit.

**Unknown.** Computed receivers, other iterators and the dword-width tests lie
outside that search. Writers outside it also exist: Item constructors
`R0884`/`R0972` store `+8 = 0` (`L09175`, `L09176`), and the Item
clone slot `vt+0x44` reaches Token copy constructor `R1134`, which copies
`+8` (`L09177`..`L09178`).

### SAV-1115

- The input carries `+0x08 = 1` on 35 Item records. 33 sit in the four Humans'
  `+7c` containers (15 Armor, 18 Weapon). 2 are equipped Armor.
- The resave carries 0 on all 35. The other 84 Token heads are 0 in both
  files.
- `R0950` writes `+0x08` literally on STORE (`L09179`) and reads it
  literally into `+0x08` on LOAD (`L08999`).
- The only located writer of 0 is `R0883` (`SAV-1114`). Actor entry
  `R0059` reaches it for carried and equipped Items subject to class,
  recipient ownership and type gates (`SAV-POSTLOAD-222`). All four Humans of
  this pair pass them.
- Each Human's Token `+18` publication mask goes from 0 to 2, so a send path
  published the four Humans to a recipient during the owner's session. This
  does not name the clearing call.
- Prediction: entering play from the input in the original posts one
  "Picked up" line for each flagged carried Item, up to 33, and none for the
  two equipped Items. Entering play from the resave posts none. The first
  half is refuted if no such line appears.
- Evidence files: `evidence/q1-item-token08.tsv`,
  `evidence/q1-item-token08-summary.tsv`, `evidence/listings.txt`.

**Confidence.** High for the 35-record measurement and for the literal LOAD
and SAVE arms. Medium that a `R0883` publication cleared the words. It is
the only zeroing writer in `SAV-1114`'s bounded search, but the resave cannot
separate session entry from a later publication.

**Unknown.** The message prediction has not been observed.

### SAV-1116

- Speed is stored twice: as Unit `+0x8c` (row 17, fifth u16) and as the mover
  byte `*(+0x154)+0x0a` inside raw180.
- Derive `R0280` (`vt+0x50`) computes speed in six steps:
  1. Cap Reaction at `(i8)mod[+0xd5] + 50`.
  2. Set speed to Reaction when Reaction is below 12, otherwise to
     `Reaction/5 + 12`.
  3. Add 10 for type word `0x13` or `0x15`.
  4. Set load to `+0x8e` plus half of container `+0x20`, or to `0x7d00` when
     container `+0x20` is at least `0xfa00` (`L04611`), and capacity to capped
     Body `x10 + 1`. When load is at least capacity, subtract `load/capacity`,
     with a floor of 6.
  5. Add `i16 mod[+0xd8]` (`R0840`). When speed is then negative, zero
     `mod[+0xd8]`, not speed. After that, add `mod[+0xda]` to capacity `+0x92`.
  6. Copy the low byte to the mover (`L03931`).
- From each file's own fields the formula gives Danath 20 (`25/5+12-1+4`),
  Naira 20, Brian 18 (`39/5+12-1`) and Fergard 16. The saved `+0x8c` and the
  saved mover byte equal these in both files. So do the saved load and
  capacity words.
- The traced Human LOAD bodies do not call derive (`SAV-HUMLOAD-445`,
  `SAV-LOADHOOK-879`).
- Derive has no direct call and two raw vtable-slot occurrences.
- None of the 49 indirect-call sites through `[e?x+0x50]` lies in the body of the entry join
  `R0131` or of actor entry `R0059`. Their callees were not searched.
- SAVE writes the live `+0x8c` word and the live raw180 mover byte. On either
  model the four Humans hold 20, 20, 18 and 16 after LOAD.
- Evidence files: `evidence/q3-speed.tsv`, `evidence/listings.txt`. Also cites
  `HERO-SPEED-008`, `ITEM-LOAD-005` and `MOVE-RATE-053`.

**Confidence.** High for the per-Human values. The formula is read at
instruction level and every input comes from the file. Medium that the traced
Human LOAD bodies and the bodies of `R0131` and `R0059` contain no derive
call. The search covers direct calls and the `e?x+0x50` form, attributes each
site by a backward frame scan, and does not follow callees.

**Unknown.** Whether the original re-derives after LOAD. The join reaches
derive through `R0065` on actors it constructs (`R0497` → `R0656` →
`L04444`; `R1562` → `L08131`); no search established whether any
session-entry path derives a loaded Human. Town entry was not identified, and
any later derive trigger is also open.


### SAV-1118

- Session-entry join `R0131` has one caller, `R0061` (`L06672`). It
  calls `L03114` (`L03115`). `L03114` calls sender `R0059` with mask
  `-1` for the Player's `+34` actor, for every element of the tick list
  `[L00240]+4` (`L03788..L03789`) and for every element of the dead
  list `*(world+0x14)+0xc` whose stage `+13c` is below 5 (`L03790`).
- The sender reduces the mask in four places:
  - it clears `0x1000` when the class is not `0x1a`, `0x1b` or `0x45`
    (`L09180`);
  - it keeps `0x507b` for a non-humanoid actor not owned by the recipient
    (`L03731`);
  - it keeps `0x50fb` for a humanoid not owned by the recipient whose type
    `+0e` lies outside `[0x21,0x3f]` (`L03732`);
  - it clears `0x2` when `+9c` is 0 (`L09181`).
  Bits `0x4000`, `0x1` and `0x20` survive each reduction. The packet is emitted
  when `mask & 0xff7fff7f` is nonzero (`L03238`). Token `+18`
  (`R0676`/`L08644`) gates only the Effect-list and `+4c` bit `0x08`
  extras after the packet, not the packet itself.
- Client arm `L02530` (opcodes `0x6c`, `0x6e`, `0x6f`, `0x70`) drops the
  message when the owner byte exceeds view `+9a8` or names a null view
  `+9a4` slot (`L09182`, `L09183`), or when a non-hero class has no table
  entry (`L09184..L09185`). It looks the runtime id up in view `+9b8`.
  For an absent id it drops the message unless mask bits `0x4001` are both set
  (`L09186`), then allocates a `0x1b0`-byte drawable (`L09187` or
  `L09188`) and inserts it. It stores the message position at drawable `+8`
  and `+c` (`L09189`, `L09190`).
- The LOAD route `R0099` runs the optional pre-entry tick (`L03780`) and
  then pumps the client dispatcher (`L03785`) until the session starts.
- Evidence files: `evidence/static/d-server.txt`,
  `evidence/static/callers.txt`, `evidence/static/r-client-state.txt`,
  `evidence/static/r-load-pump.txt`. Also cites `SAV-POSTLOAD-220`,
  `SAV-POSTLOAD-221`, `SAV-678`, `MOVE-TICK-013`, `ANIM-MSG-005`.

**Confidence.** High for the named instructions and the one-caller join. Medium
that the client state arm is the only route that creates an actor drawable.

**Unknown.** Whether the pre-entry tick runs for a given document, and how many
server ticks separate LOAD from the entry send. Other drawable allocation
sites exist and none was shown to run at LOAD: `0x1b0` allocations with the
state arm's constructor `R0593` at `L09191`, `L09192` (it inserts into
view `+9b8` under fixed id 1) and `L09193`, and the dispatcher arms `0x82`
(Building objects in the same `+9b8` map) and `0x86`, `0x8b`, `0x8c`
(`0x14c` objects).

### SAV-1119

- In `R0059`, mask bit `0x20` sends the Token Position through getters
  `R0165`/`R0166` (`cell<<8 | sub`) only when `R0040` returns 1
  (`L09194`, `L09195`). `R0040` returns 1 exactly when Position
  sub-cell `+4` and `+5` are both `0x80`.
- Otherwise the sender reads mover word `*(actor+154)+06`: its low byte and
  its high byte each become `(cell<<8)+0x80` (`L09196..L09197`). Mover
  `+06` is the cached route cell (`MOVE-CLAIM-007`).
- The client arm stores those two values at drawable `+8/+c`. A sweep of the
  dispatcher's own body `R0509..L02524` (7,325 instructions) finds 38
  instructions whose destination is `[reg + 0x8]` or `[reg + 0xc]`
  (`evidence/static/dispatcher-disp8-c-writes.txt`). Two write the `0x1b0`
  drawable, at `L09189` and `L09190`. Two write Building objects of arm
  `0x82` (`L09198`, `L09199`), six write the `0x14c` objects of arms
  `0x86`, `0x8b` and `0x8c`, and the rest write Item entries, hash-map nodes
  and counts and one view field. Move arm `0x6b` (`L09200`) looks the id
  up, drops an unknown id, and writes action, facing and duration with a
  per-direction delta; it writes no absolute position.
- Arrival routine `R0039` re-centres Position; its mask `-1` send
  (`L09201`) lies in the arm gated by cell type `0x1a` (`L00160`), not on
  every arrival.
- Of 51 direct sender calls (`evidence/static/r-send-sites.txt`), those with
  an immediate mask that includes bit `0x20` push `-1`, `0x20` or
  `0xbf7fff7f`. They include entry, joins, spawns, trigger instants 16 and 17
  (`R0122`, `R0123`) and the `0x1a` arrival arm. Several sites push a
  register whose value was not traced.
- Consequence: an actor that is still mid-step when the entry send reads it
  gets its drawable at the mover cell, while the simulation keeps it at its
  Position. For a restored actor that is the saved mover cell only if no tick
  before the entry send moved it. The pre-entry tick runs when campaign
  `+0x6b8` is nonzero, and the number of ticks before the send is
  `SAV-1118`'s Unknown. Whether the client shows the drawable was not read.
- Evidence files: `evidence/static/d-server.txt`,
  `evidence/static/r-client-state.txt`, `evidence/static/r-client-move.txt`,
  `evidence/static/r-send-sites.txt`,
  `evidence/static/dispatcher-disp8-c-writes.txt`. Also cites
  `AI-CENTRE-101`, `MOVE-CLAIM-007`, `SAV-POSTLOAD-221`.

**Confidence.** High for the Position/mover branch. It excludes the rival that
the sender always names the Position: the branch at `L09195` selects the
mover arm on a non-centred sub-cell. Medium for reading the restored bytes as
the state at the entry send, because the pre-entry tick is not excluded.
Medium for "stays misplaced until a later position send": the census of later
senders reads the pushed mask at each direct call site, not the reach of each
caller, and the `+8/+c` sweep covers direct writes in the dispatcher's own
body only; its callees and the per-frame animation driver were not searched.

**Unknown.** Whether fog of war or line of sight hides the drawable; which
later event first corrects the drawable in play.

### SAV-1120

- `game9271.sav` (this project's document, mission 141): map unit 57, runtime
  id 37, class 68, has Position cell 60,59, sub-cell `0x48,0x80`, and mover
  word `+06` = 0. If the entry send reads these bytes, it names cell 0,0 and
  its drawable is created at map cell 0,0. Map unit 168, runtime id 68, class
  68, has sub-cell `0x80,0x80` and its drawable is created at its Position
  62,62.
- The same file has two more such Units: map unit 30, class 80, cell 56,15,
  sub-cell `0xee,0xee`, and map unit 75, class 64, cell 16,23, sub-cell
  `0xee,0xee`. All three have mover word 0 and AI state `+50` = `0x0a`.
  `dragontooth.sav` has the same three records with AI state `0x0b`. Of the
  72 Unit/Humanoid/Human records in each file (62 Unit, 10 Human), these three
  are the only non-centred ones.
- In the original resave `game0001.sav`, both trolls are centred: unit 57 at
  60,61 with mover word 60,61 and unit 168 at 62,62. Both drawables are
  created at their Positions. `game0002.sav` agrees. `dragontoothresaved.sav`
  has one non-centred actor, one cell from its mover cell.
- The owner's observation after LOAD of `game9271.sav` fits unit 168 as the
  one troll seen, standing in guard state `0x0b` with `+54` 0, and unit 57,
  patrol state `0x0a` walking a nine-cell route, as the unseen attacker.
- Evidence files: `evidence/trolls.tsv`, `evidence/inputs-actors.tsv`,
  `evidence/census/noncentred.tsv`. Also cites `AI-STATE-011`.

**Confidence.** Medium. The file bytes and the rule of `SAV-1119` give the
predicted drawable cells exactly, and one owner run agrees. No run changed one
byte group alone, and the pre-entry tick could move unit 57 before the send.

**Unknown.** Whether unit 57 is still non-centred at the entry send; whether its
drawable is shown or later corrected in play.

### SAV-1121

- Population: every `.sav` under the owner save tree, 113 files
  (`evidence/q4-population.tsv`). 13 are excluded by path as project-written
  (a `generated` or `engine-from` path) or as the `EXP-0261` directory's six
  MANIFEST-labelled generated candidates. One more,
  `story-1073-original-acceptance/game0000.sav` (`bbee204a…`), is excluded by
  content: its bytes equal six excluded project-written files. Of the 99
  remaining files, 23 are byte-identical copies of earlier files and are
  counted once. That leaves 76 distinct original-written documents; some come
  from sessions seeded by a project document (`seeded`).
- The 76 documents decode completely and hold 2,779 actor records. 69 have a
  non-centred sub-cell. Their mover cell equals the Position cell for 33 and
  is one cell away for 36. None is farther.
- Evidence files: `evidence/q4-population.tsv`, `evidence/census/summary.tsv`,
  `evidence/census/noncentred.tsv`.

**Confidence.** Medium. A finite-corpus census over the named population.

**Unknown.** Whether an original save can hold a non-centred actor with a far
mover cell, for example during a teleport or a trigger return.

## Original patrol continuation endpoints

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1123 | Two owner-reported original EN continuations of mission-141 patrol probes save unit 57 idle at (60,61) after their counters advance, while units 30 and 75 retain active routes; the first dispatch remains Unknown. | Medium | ● active | [EXP-0403](../experiments/EXP-0403-patrol-dispatch/) |

### SAV-1123

- The four identified documents are one control re-encoded for original EN,
  a copy whose decoded body differs only at unit 57's order+04 low byte
  (`1` to `2`), and two owner-reported original-English resaves. Their full
  SHA-256 values, body sizes and local offsets are in the experiment's
  `evidence/inputs.tsv` and `evidence/latch-control.tsv`.
- Both outputs save map unit 57, runtime id 37, type word 68, at cell 60,61
  with state+50 `0x0a`, motion+54 zero, no dynamic route, order+08 zero,
  order+04 one, waypoint and destination 63,59, and unchanged deadline+138
  1097. Its Group has one member, selector+1c 12, AI order+20 zero,
  activity+45 one and override+48 one. It belongs to Player 2 and its
  saved owner key matches that Group's +44 reference.
- Units 30 and 75 also belong to Player 2. They change position and current
  route between the input and each output, retain motion+54 one and have
  nonempty routes. The first document-head dword advances 1219 to 1816 and
  2256: 597 and 1037 increments. Numerically, 38 and 65 intermediate values
  have residue 6 modulo 16, the AI slot in `AI-TICK-008`. This count assumes
  traversal through each intervening value; it is not a record of calls.
- Both outputs save order+04 as one. If each output descends from a different
  input probe as the owner reported, the input marked two changed to one
  regardless of the still unknown slot assignment. The walker has a known
  write of one (`SAV-PATROLCURSOR-571`), but this endpoint does not identify
  which writer ran or when.

**Confidence.** High for the decoded bytes of the four SHA-pinned documents
and the local original-instruction path checked against `rom.exe` SHA-256
`942e9b72…d367d03`. Medium for original-runtime attribution: the two output
files were supplied as original-game resaves, without a native dispatch trace
or a confirmed input-to-slot mapping. Saved endpoint agreement cannot select
one of the intervening runtime paths.

**Unknown.** Whether first Group dispatch reached unit 57, what guard left in
order+08, whether a route was admitted and later cancelled, and whether the
troll moved transiently between LOAD and SAVE. A first-phase native trace at
the Group gate, walker, guard return and route admission would discriminate
these alternatives. Continued movement of units 30 and 75 does not itself
prove that the Group dispatcher ran: a previously admitted route can advance
without a fresh Group call.
## Non-hero card after LOAD

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1125 | The entry sender selects card fields by ownership and class; a non-owned Human of type 9 or 24 receives health/mana but no primary or combat values in its reduced state message. | High | ✔ promoted | [EXP-0404](../experiments/EXP-0404-enemy-card-load/) |
| SAV-1126 | The selected type-9/24 client creation path leaves omitted primary values unchanged; the numeric card reads drawable bytes, with two independent five-byte extended-row loops. | High / Medium / Unknown | ✔ promoted | [EXP-0404](../experiments/EXP-0404-enemy-card-load/) |
| SAV-1127 | Five candidate Human pairs in one external mission-151 SAV and its reported original resave retain primary values that differ from the owner card screenshots; the two saves are 546 full-clock units apart. | Medium | ✔ promoted | [EXP-0404](../experiments/EXP-0404-enemy-card-load/) |
| SAV-1128 | Native mission151 shares primary words and attack with the external document on 217 authored ids; observed absorption0/10 does not invalidate the repeatable200 setter, and a conditional Human derive can overwrite it. | High / Medium / Unknown | ✔ promoted | [EXP-0404](../experiments/EXP-0404-enemy-card-load/) |

### SAV-1125

`SAV-1118` identifies entry's initial mask `-1`. Sender `R0059` reduces
it when actor owner differs from recipient. Non-humanoids retain `507b`;
their additional class test removes `1000` except for types 26, 27 and 69.
Humanoids retain the full mask for types 33..63 and `50fb` otherwise.
Equal-owner actors bypass these reductions. `+9c==0` separately removes
the current-mana bit. All masks here are hexadecimal.

Primary words `+84/+86/+88/+8a` enter the packet as low bytes only under
`20000000`. Client `R0509` writes them to `+138/+13b/+139/+13a`.
Combat packet masks are `100` attack `+a6`, `200` damage sums, `400`
defence `+be`, `800` armour `+c0`. Maxima use `1000`, sight `2000`,
speed `200000`. Protection masks `10000..100000` use the low bytes of
`+c4/+c6/+c8/+ca/+cc`. Skill masks `1000000..10000000` use low bytes
of `+aa/+ac/+ae/+b0/+b2`, not the actor's damage-resistance bytes.

Unit serialization stores and restores the primary words. Its definition
binding suffix (`SAV-ACTORBIND-544`) is separate from these values. The
selected SAVE scalar writes read actor state, not the client drawable.

**Confidence.** High for these named local branches and transfers. The
probe executes 44 ownership/class/type mask vectors from the original
instructions and validates its three helper contracts against their bodies.
This is not an independent runtime or all-message census.

**Unknown.** Later mutation before sending; complete client-writer lifetime;
arbitrary saved-class validity; whether an exact actor reaches the first
displayed card without an intervening writer.

### SAV-1126

Client state creation calls prototype fill `R0590` at `L03117`, then
appearance fill `R0551`. For types 9/24 the former only sets class/face
and flag state; it takes no prototype-stat copy. Its flags lack bit 0, so
the latter returns through its entry gate. The subsequent message writes
only fields named by its mask. Twelve synthetic vectors run the original
prototype and primary-copy slices: mask `50fb` retains each seed
`00/5a/c3`; control mask `20000000` copies payload 49/35/28/50 into
drawable byte order 49/28/50/35.

`R0877` draws Body/Agility/Mind/Spirit from `+138/+13b/+139/+13a`.
Primary rows require disclosure level 5; the `L01661` global can force
level 7. Ordinary attack reads signed `+146`, defence signed `+144`, armour
byte `+142`, and damage bytes `+148/+149` as min and min+spread. Sight
word `+102` uses 1/256 scaling, speed word `+104` prints as an integer.
The first extended loop walks five bytes `+14b..+14f` with text-pointer
offsets `78..88`, or `90..a0` for types 23/24. The second walks five bytes
`+150..+154` with captions at `a4..b4`. These are the numeric card loops.
`UNIT-PANEL-011`'s indexed read `L09202` instead belongs to the separate
type-23/24 spell-mask loop; it does not define these five-row card indexes.

Constructors `R1652`/`R0593` directly zero health/mana, sight and
speed. Their bodies contain no direct store into primary `+138..+13b` or
combat/extended `+142..+154`. This is a two-body direct-store observation,
not an all-callee or allocation-lifetime absence proof.

**Confidence.** High for the local mask/retention discriminator and the
positive card readers. Medium that stale drawable contents explain the
owner symptom. The live alternative is an untraced intervening writer.
Unknown for the first-frame contents and the full initialization history.

### SAV-1127

Inputs are external `game2320-input.sav` (SHA-256 `1ee3466b514ed90e69e05ad01eeceb49a20f2712cac0b64c7e88d1c06479b400`)
and reported original `game0002-original.sav` (`aaac9a89705250144acb8933ef93e52b9a02f83abc9c21e65c2917836e6d9db3`).
Research `savdoc` walks 222 actor bodies in each. Both name mission 151;
full-clock values are 8729 and 9275. The comparison is not an immediate
LOAD/SAVE pair and the external input is not an original-native oracle.

Type-9 Axeman candidates, runtime ids 98/99 and authored ids 252/253,
both retain primary 49/35/28/50, health 607/607, attack 308 and defence
216. The owner Axeman images show 196/3/94/3, health 607/607, attack
772 and defence 1272. Type-24 Mage candidates, ids 96/97 and authored
250/251, retain 17/40/46/50 and health 117/117, mana 630/630;
candidate 139/589 retains 17/24/42/50 with those same maxima. The Mage
image shows 5/0/0/0 and health/mana 117/117, 630/630.

**Confidence.** Medium for this finite comparison. Names are checked against
the repository's installed-text census. The screenshot contains no actor
id, so the five matches are candidates. Changing position, intervening
ticks and omitted client fields remain separate variables. Correct maxima
do not establish that the displayed primaries came from the actor.

**Unknown.** Exact screenshot actor and capture time; which writer supplied
each displayed abnormal byte; native baseline cards before the external
input; whether the resave was reloaded before the screenshots.

### SAV-1128

The two SAV-1127 documents plus native `game0003-original-native151.sav`
(SHA-256 `54a1d77a07da694f816307f0780a481465b7c54ffadc8d1068850d1336406414`)
and supplemental entry `game9999-original-entry151.sav`
(`d776f109a9d221999852b7d9d61efb315ea09f933336480e3d472edc472db846`)
are mission 151, difficulty2. Full clocks are 8729,9275,16,0. Exact research
walks reach EOF and tile every actor subtree with named member spans.
Actor counts are 222,222,242,242. Each file has five zero-id actors, which
are not cross-file joined; all nonzero map ids are unique. The four-file
intersection has 217 ids. Runtime ids and addresses do not define this join.

On those217 ids the external/native primary words, attack, maximum health,
class, type, definition index and owner-player index all match. Other
differences include 208 external-zero/native-nonzero Token+1c values,
217 actor+136 values 0/1,113 defence differences,62 speed differences and
99 sight differences. All110 Human+1e4 references name Diary bodies in
external/resave and are null in native/entry. Unequal clocks, party history,
effects and gameplay prevent treating this corpus as a same-state oracle.

For ids 250/251/252/253, absorption is 200 in external/resave and 0/0/10/10
in native/entry. Lawful mission 151 placement bytes+31/+32 are zero on
these four. Actions 78..81 instead author opcode 34, selector 16, value 200;
repeatable trigger 33 compares condition 1 to itself. Original `L09203`
writes actor+c0 as a word. The preceding repeatable triggers 29..32 enqueue
temporary casts 5/10/16/22 at those targets. The effect-dispatch common tail
`L06977` calls target virtual+50; both Human slots select `R0280`,
which zeros actor+be..+d3 and folds modifiers. Their absorption modifier
word is 0/10 in all four documents. This gives a conditional overwrite path,
not a universal SAVE value. Native clock 16 has latch 33=1 and four attached
effects per selected target, so a never-fired trigger is not its explanation.

**Confidence.** High for the exact actor-span tiling, reference resolution
and selected instruction relations. Medium for the finite corpus comparison
and the delayed-cast/derive explanation. The full member/byte matrix and
all-record extent/hash inventory are reproducible; arbitrary item/effect
identity across documents is not proven.

**Unknown.** Last writer before each SAVE, original callback cadence at the
sampled frame, cause of the other persistent constants, and whether any of
these actor differences supplied the displayed abnormal client bytes.
A runtime trace reaching the 200 setter then the documented Human derive
with unchanged modifiers, but retaining 200 afterward, refutes the proposed
local overwrite. No original-runtime trace was run.

## Projectile construction and lifetime

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1129 | A physical ranged attack builds its CProjectile in unit virtual slot +0x58 (`R0603`) at the action-7 ShootDelay tick; the 0x72 arm and melee build none, and a cast-diverted attacker's shot builds none through slot +0x58. | High / Medium | ● active | [EXP-0428](../experiments/EXP-0428-fresh-projectile/) |
| SAV-1130 | At the unit shot, the 16 stored projectile leaves take values from shooter, target and class, or stay zero; no leaf records the shooter, the weapon or an item. | High / Medium | ● active | [EXP-0428](../experiments/EXP-0428-fresh-projectile/) |
| SAV-1131 | The five live projectile creators take their id from one 16-bit counter at client world +0xa0c and insert into a 17-bucket id hash; the loader takes it from the document. Saved IDs follow bucket order, newest first within a bucket. | High / Medium | ● active | [EXP-0428](../experiments/EXP-0428-fresh-projectile/) |
| SAV-1132 | A projectile of pictures 0..12 never applies damage (pictures 13+ are unread): the simulation actor applies it at its own countdown expiry against the then-current target and range, so a shot and its damage are independent events. | High / Medium | ● active | [EXP-0428](../experiments/EXP-0428-fresh-projectile/) |
| SAV-1133 | A projectile in flight is removed only by its own countdown ending or by a collect-all clear; no removal path keys on shooter or target, and a missing target leaves it flying to its last aim point. | Medium / Unknown | ● active | [EXP-0428](../experiments/EXP-0428-fresh-projectile/) |
| SAV-1134 | The sole corpus projectile Prj266 fits the construction rules for its leaves except picture 10, which is a mage or spell picture, so no physical-arrow record was found in the scanned documents. | Medium | ● active | [EXP-0428](../experiments/EXP-0428-fresh-projectile/) |

### SAV-1129

- The client class `CProjectile` (size `0x14c`, vtable `L02587`) has one
  constructor, `R0609`, which calls the drawable base `R1652`.
- Constructor calls: `R0603` (unit slot +0x58, installed for CUnit at
  `L09204` and CAirUnit at `L09205`), `R0620` (unit slot +0x5c, the
  cast spawner), client dispatcher arms for opcodes 0x86, 0x8b and 0x8c
  (`L03002`, `L03003`, `L03004`), and the SAV loader (`L03007`).
  The linear sweep of `CALL` targets finds these 6 sites; the raw dword scan
  of the constructor address finds 0, because a `CALL rel32` carries no
  absolute dword, so the census is one instrument.
- Ranged attack message 0x72 (arm `L09206`) holds no constructor call. If
  the drawable's run counter (+0xa0) is zero it sets action 7, actiondir,
  actionphase 0 and actiontarget, and the run length to the class attack
  timeline. If the counter is non-zero it logs "Overriding ... by 'Shoot'"
  and writes nothing (the branch at `L09207`, listed to `L09208`), so that
  shot builds no projectile.
- Unit action driver `R0548`, action 7 (arm `L02711`), calls slot +0x58
  when actionphase equals class ShootDelay (+0xfc) and class Projectile
  (+0xd4) is non-zero.
- The simulation sends 0x72 when actor reach (+0x12c) exceeds 1, else 0x71
  (`R0549`, builder `R1064`). `tools/shotclass` over Data.bin Units and
  Humans on EN finds 50 rows that satisfy reach above 1, no weapon-spell
  divert and ShootDelay below charge plus relax: Units 18 (Catapult,
  Ballista, 4 Goblin_Sling, 4 Orc_Bow, 4 Bat_Sonic, 4 Dragon), Humans 32.
  The classes are Human Archer, Crossbowman, Orc Archer, Goblin slinger,
  Catapult 1 and 2, Sonic Bat and Dragon (class Projectile 1..7 and 12).
  Mage-staff Human rows (46) are cast-diverted (`MAGIC-AUTOCAST-020`): the
  weapon-spell attack goes to the cast spawner `R0620` (slot +0x5c), which
  builds a record (picture 10 in Prj266, `SAV-1134`), not through slot +0x58.
  132 Human rows have reach 1. RU `Data.bin` Units rows are byte-identical.
- Catapult and Ballista have no mana, so they take the strike path with
  pictures 5 and 6 plus their Fire_Ball weapon-spell rider.

**Confidence.** High for the routine, the 6 constructor call sites, the 0x72
arm and the row census, read end to end on EN. Medium for "only" at the
action-7 tick: the indirect callers through `[reg+0x58]` were not enumerated; the
evidence is the one site `L02716` and the two vtable slots (`L09204`,
`L09205`). Medium that the classes listed are the ones that reach it;
which class a hero's bow maps to is inference, not read. RU `rom.exe` is byte-identical to EN
(SHA-256 `942e9b72...7d03`). Medium for how often a 0x72 arrives while the
drawable is running: charge plus relax plus 2 exceeds each attack timeline
for the shipped rows, so an idle drawable is expected, but no run was counted.

**Unknown.** A structure as shooter. The object the Fire_Ball rider of the
siege rows constructs, which was not read. Whether the dropped-0x72 case
occurs in play. What triggers client arms 0x86, 0x8b and 0x8c: the simulation
stores those opcodes at `L05301`, `L03021`, `L05302`, `L03022`,
`L03023`, `L05303`, `L03055` and `L05304`, and whether any of those
senders runs for a ranged attack was not traced. The 0x72 path is answered;
the arms' senders are open.

### SAV-1130

All offsets are CProjectile bytes. Constructor base `R1652` zeroes x, y,
z, dir, phase, lastaction, action, actiondir, actionx/y/z, actionphase,
actionsegments and actionspell, and sets +0x80 to -1; it does not
initialise +0x86. Derived-field refresh `R0614` writes none of the 16.
Unit shot `R0603` then sets:

| Leaf | Offset | Value at construction |
|---|---|---|
| x, y | +8, +0xc | shooter x (y) + (ShootOffset[(dir-8) & 0xe] - class Center) * 8 |
| z, dir, phase, lastaction, actiondir, actionspell | +0x10, +0x6c, +0x70, +0x74, +0x85, +0xa4 | zero |
| picture | +0x20 | class Projectile (+0xd4) |
| action | +0x84 | 1 |
| actiontarget | +0x86 | shooter's target id (+0x86); no target in the unit hash builds nothing |
| actionx, actiony, actionz | +0x88..+0x90 | zero |
| actionphase | +0x94 | 0 |
| actionsegments | +0xa0 | truncated distance between shooter and target (+8, +0xc), divided by 200 |

- The driver `R0558` overwrites z, dir, phase, lastaction, actiondir and
  actionx/y/z on its first call. The shot is built during the unit pass of
  a tick, so that call is in the same tick.
- ShootOffset is the CArray data pointer at class +0xec, indexed without a
  bounds check. The cast spawner `R0620` also sets dir and actiondir at
  creation and takes picture from the shooter's actionspell. Neither spawner
  sets actionspell.
- Non-leaf fields: +4 id, +0x14 terrain pointer copied from the shooter,
  +0x28 and +0x2c position snapshots.

**Confidence.** High for the field writes and for the zero or unset fields,
read from the constructor, base and shot listings. Medium for the scale: the
256-per-cell fixed point and 8 per pixel are inferred from the saved leaf
values (`SAV-1134`) and `UNIT-STRUCTDELIVERY-065`, not from a constant in
the routine.

**Unknown.** The registry rows' ShootOffset values per class, which this
experiment did not decode.

### SAV-1131

- Live creators `R0603`, `R0620` and the dispatcher arms 0x86, 0x8b and
  0x8c, and the loader (`L03007`), reach the store at client world +0x9d4. Buckets are at
  +0x9d8, bucket count at +0x9dc (constructor `L09209` sets 17), count at
  +0x9e0, node free list at +0x9e4.
- Each live creator sets id = word at world +0xa0c, stores FreeIndex = id + 1 as
  a 16-bit word, and inserts a node (+0 next, +4 bucket, +8 id, +0xc
  object). Bucket is (id >> 4) mod 17; insertion is at the bucket head. An
  existing node with the same id keeps its place and has its object pointer
  replaced.
- The serializer (`L09210`) walks buckets in order and appends ids, so
  IDs order is ascending bucket, head first. The loader (`L03007`)
  recreates each record and sets FreeIndex from the document (the loader
  does not take its id from the counter). Listings: `L09209` (bucket count
  0x11) and the serializer range `L09210` are in `evidence/`.
- Corpus: 29 distinct world-state documents (`SAV-PROJCORP-430`) hold
  FreeIndex values 0 (7), 1 (3), 8 (6), 31, 33, 40 (2), 171, 180, 239, 252,
  258, 267, 288 (3); 28 of them have empty IDs. Instrument: EXP-0263
  `state-census.tsv`.

**Confidence.** High for the counter, the hash layout and the insertion
order. Medium for the corpus reading that FreeIndex is a monotone
construction counter rather than max(IDs) + 1: 28 empty-ID documents with a
non-zero FreeIndex exclude the latter for those documents, but they do not
show what a saved FreeIndex does after a load.

**Unknown.** Whether a world reset restores the counter to 0 (the seven
zero documents are consistent with it); the largest id a live world holds.

### SAV-1132

- The simulation swing start `R0245` (one caller, `L05237`, at phase
  0) takes d, the footprint-aware cell distance (`R0247`, minimum 1).
  For d > 1 it adds extra = (d * 256 + 128) / 200 to the countdown at actor
  +0x6c, which starts at charge. It then sends 0x72 or 0x71 through
  `R0549`; the 0x72 builder `R1064` writes attacker id and target id.
- At phase-5 expiry `R0001` re-reads the actor's current target
  (actor +0x5c). For a unit target it calls `R0246`, which returns with
  no damage when the target or attacker is null or unbound, when attacker HP
  is not above 0, or when reach is below `R0247` at that moment. Else it
  computes damage through target virtual +0x4c and subtracts it from HP,
  including from a target already at or below 0, and sends the 0x73 message
  only when HP was above 0 before or is above -10 after.
- The projectile driver `R0558`, for pictures 0..12, has no damage arm,
  no message send and no unit write: those pictures take the default arm
  (move, animation phase); pictures 10 and 12 also push smoke points. These
  cover every physical shot (1..7, 12) and the mage picture 10. Pictures 13
  and above dispatch through the jump table `L03000`; its first arm
  (`L05684`) writes the target unit and was not read as damage or
  otherwise.
- A target that moved beyond reach before expiry takes no damage, and a
  projectile built earlier still flies to it. A target that died takes a
  further HP subtraction and no message below -10.

**Confidence.** High for the call chain, the countdown and the damage gates,
and for the absence of any damage write in the driver for pictures 0..12.
The swing-start and expiry callers (`L05237`, `L00006`) and the countdown
are listed in `evidence/disasm-actor-tick-L05238.txt`. Medium for extra
approximating flight time: it is the distance the shot divides by the same
200 for segments (`SAV-1130`), and the shot and damage ticks align only if
client and simulation tick counts are equal, which is an inference. For
d = 1 an archer's damage tick (charge 20) precedes its release tick
(ShootDelay 21).

**Unknown.** Whether actor +0x5c is cleared when its target is removed, and
the weapon-spell rider's own timing. What the driver's jump-table arms for
pictures 13 and above do, including the unit write of the first arm.

### SAV-1133

- The world tick (`R0334`..`R0234`) calls driver slot +0x3c on every
  store node. A zero return is collected; a second loop looks each id up,
  unlinks the node and calls the destructor (slot +4). The driver returns 0
  when actionsegments is 0 and otherwise decrements it and returns 1. A shot
  built with segments N therefore lives N + 1 ticks (`SAV-1130`).
- While action is 1 and actiontarget is non-zero, the driver looks the
  target up in the unit hash (`L05558`). If found it overwrites actionx/y/z
  from the target and recomputes actiondir. If not found, or if actiontarget
  is 0, it keeps the existing actionx/y/z.
- A record stores no shooter reference (the 16 leaves, +4, +0x14, +0x28,
  +0x2c). All 23 sites that name world +0x9d4 were classified. Removal
  sites are the driver return, the unconditional collect-all-and-destroy
  functions `R0809` (caller `L07290`) and `R1301` (11 callers), the
  destructor loop in `L09211` (caller `L09212`) and unwind code.

**Confidence.** Medium: the census is a displacement and immediate sweep plus
a raw dword scan and agrees at 23 and 23, but a pointer reached by
arithmetic from another world field is invisible to both. High that the
driver and tick read as stated.

**Unknown.** What the 11 callers of `R1301` are, whether any runs
mid-mission, and what a shooter's death does to simulation damage already
counting down (`SAV-1132`).

### SAV-1134

Prj266 in `game0018.sav` holds action 1, actiondir 5, actionphase 2,
actionsegments 3, actionspell 0, actiontarget 157, actionx 19840, actiony
28288, actionz 0, dir 5, lastaction 1, phase 1, picture 10, x 19278, y
27985, z 0; FreeIndex is 267 and IDs is [266]. Instrument: a state dump by
`tools/savauthor` `TestProjectileLeafDump`.

- actionx and actiony equal 77 * 256 + 128 and 110 * 256 + 128, the centre
  of cell (77, 110). actionphase plus actionsegments is 5: each driver call
  adds 1 to the first and subtracts 1 from the second, so the record was
  built with 5 segments and has run 2 driver calls. Action 1, lastaction 1
  and actionspell 0 fit `SAV-1130`.
- Picture 10 is not an archer or crossbow value (class Projectile 1, 2). It
  is the Human Mage class 24 Projectile and the fire_arrow cast picture
  (2 * 1 + 8). Its leaves fit either constructor, so Prj266 does not
  witness a physical arrow shot. No stored record with picture 1..7 exists
  in the scanned documents.
- Population: `TestProjectileLeafDump` over 146 `.sav` files under
  `gameversions` parsed 104; the only nonempty Projectiles store is Prj266,
  in 3 byte-equal copies of `game0018.sav`. 42 files did not parse in this
  instrument, so the result is bounded to the 104.

**Confidence.** Medium: one record, two compatible constructors, 104 parsed
documents.

**Unknown.** A stored arrow record (picture 1..7) and its saved leaf values;
a cast record's actionspell, which would separate the two spawners.

## Hired siege Unit creation and reuse

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1135 | In EN `rom.exe` 942e9b72, the tavern spawn `R1142`, run by server command 0x38 from the hire vector, creates a fresh siege Unit (types 1, 2) as a party actor; a direct-call sweep finds one caller. | High / Medium | ● active | [EXP-0429](../experiments/EXP-0429-siege-entry/) |
| SAV-1136 | In EN `rom.exe` 942e9b72, `R1142` never reads existing party actors, so it never reuses a Unit; the only removal of prior hires read is command 0x37's `R1555`, which removes party actors whose type byte `+0x14c` is nonzero. | High / Medium | ● active | [EXP-0429](../experiments/EXP-0429-siege-entry/) |
| SAV-1137 | In EN `rom.exe` 942e9b72, a siege Unit's type is stored (dword `+0x148`) but its load arm zeroes it, so a loaded Unit is exempt from `R1555` and a fresh Unit is created beside it; a loaded NPC-named Human keeps its type and is replaced. | High / Medium | ● active | [EXP-0429](../experiments/EXP-0429-siege-entry/) |
| SAV-1138 | In EN `rom.exe` 942e9b72, opcode 0x38 has one literal store (in `R1653`, called from inn handler `R0786`) and 0x37 one (in town roster refresh `R0645`); no city-load routine read calls either. | Medium / Unknown | ● active | [EXP-0429](../experiments/EXP-0429-siege-entry/) |
| SAV-1139 | In EN `rom.exe` 942e9b72, a saved party siege Unit carries no per-actor link to a hire record in two compared saves: its type dword, one-actor Player group and class are the party-side fields, beside the campaign hire flag and pool arrays. | Medium | ● active | [EXP-0429](../experiments/EXP-0429-siege-entry/) |
| SAV-1140 | In five saves (EN `rom.exe` 942e9b72), a kit city save with one siege Unit, entered into mission 120, saved two Units and the same money; the original city resaves with the same party hold one Unit with type dword 2. | Medium | ● active | [EXP-0429](../experiments/EXP-0429-siege-entry/) |

### SAV-1135

- `R1142` is the body of the server command 0x38 handler (`L08144`), its
  only call site (linear sweep and raw dword scan). It takes a `CWordArray`
  `working[t] x hired[t]` and, for each type `t` with a nonzero entry, builds
  that many fresh actors. Types 1 and 2 are `Unit` (size `0x198`) named from the
  table at `L09213` ("Catapult", "Ballista") through `R0501`; types 3 and up
  are `Human` through `R0497`.
- Each actor gets a fresh id at `+4`, the type byte at `+0x14c`
  (`L09214`), the player at `+0x14`, an entry in the player flat list
  (`[player+0x20]`, `R0032`) and a new one-actor group (`new 0x48`,
  `R0149`) appended to `[player+0x24]`.
- The vector is built on the client by `R1654` (one caller, `L09215`) in
  `R1653`.
- `R1555` (command 0x37) also builds 15 prototype objects (types 1..2
  `Unit`, 3..15 `Human`) in `[tavern+0x6c]`. They are registered in the world
  registry `L00522`; no prototype is in a Player group of the five decoded
  saves.
- No call to `R1653`, `R0645` or `R1142` appears in the read range
  `R0099..L09145` of the mission starter (the function runs past
  `L06741`). The whole-image direct-call sweep finds exactly one caller of
  each (`L08144`, `L09216`, `L09217`); it excludes direct calls only, not
  calls through a virtual slot. The placement walk `R0065` reads the party
  from `[player+0x20]`.
- In `game0000.sav` (mission 120) the hero Player holds a Unit with run id 145
  and type dword 2 that is absent from `game9607.sav`.

**Confidence.** Reading is EN `rom.exe` 942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03; the RU image was not compared.
High that `R1142` builds a party siege Unit from the vector.
Medium that it is the only such creator in the binary: `R0501` has 8 call
sites, two in the tavern routines (`L09218`, `L09219`); six others were not
read. Medium that command 0x38 produced run id 145, which is inferred from the
saved state, not traced.

**Unknown.** The six other `R0501` callers (`L09220`, `L09221`,
`L02317`, `L09222`, `L03125`, `L09223`). Whether the original sent
command 0x38 during the run that wrote `game0000.sav`.

### SAV-1136

- `R1142` walks the hire vector and constructs; its range contains no read
  of the player flat list or group list before the constructor calls and no id
  comparison against an existing actor.
- `R1555` first credits `[tavern+0x9c]` to the player (`R0449`), then walks
  the player flat list and removes each actor whose byte `+0x14c` is nonzero
  (frees the id `R0868`, unlinks from the group `R1140`, removes from the
  list `L09224`), then zeroes `[tavern+0x9c]`.
- Nothing else in the read range removes or matches a party actor.

**Confidence.** Reading is EN `rom.exe` 942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03; the RU image was not compared.
High that neither routine reads a link between an actor and a
hire record. Medium that the pair `R1555` then `R1142` is the whole path;
the senders are `SAV-1138`.

**Unknown.** A reconcile step outside these routines was not searched for.

### SAV-1137

- `Unit::Serialize` (`R0210`) stores dword `+0x148` from the type byte
  `+0x14c` (`L09225`). Its load arm reads the dword into `+0x14c`
  (`L09226`) and then calls virtual slot `+0x30`.
- The Unit vtable `L00001` slot `+0x30` is `R0182`, which returns 0. The
  load arm then sets `+0x3c` from the Units table and writes 0 to `+0x14c`
  (`L09227`).
- The Human vtable `L00003` slot `+0x30` is `R0183`, which returns 1, so
  the type is kept. `Human::Serialize` (`L08212`) zeroes `+0x14c` after the
  load (`L09228`, `L09229`) when the template class name does not contain
  "NPC" (string `L09230`).
- In `Unit.scalar-run-2` the dword is bytes 47..50. `game9607.sav` object 189
  (Unit run id 14) holds 2. In `game0000.sav` the Unit with run id 14 holds 0,
  and the added Unit with run id 145 holds 2; the typed Humans of the kit save
  (run ids 3..13) are gone and replaced by Humans with run ids 147..168.
- The escape explanation is read from `game0000.sav` only. The three original
  saves `game0009.sav`, `game0010.sav` and `game0014.sav` hold a Human named
  "Ballista" (run id 14, type dword 0) and no Unit with run id 14; the
  difference from `game0000.sav` is unexplained.
- Census bound only: 2590 Units in 58 parsed corpus saves have type dword 0
  (`evidence/save-unit-type-census.txt`; one file with another header magic
  skipped). The population is not shown to hold a loaded siege Unit and the five
  case saves are outside it, so it does not test the load arm.

**Confidence.** Reading is EN `rom.exe` 942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03; the RU image was not compared.
High for the code reading. Medium for the explanation of the
two Units in `game0000.sav`: it fits every field read, but the order of
commands 0x37 and 0x38 was not traced. By the same load arm an original-written
Unit would load with type 0; that reload has not been read.

**Unknown.** A Unit saved with type 2 and then resaved by the original (no such
pair was found); why the three original saves hold a Human "Ballista".

### SAV-1138

- A literal store of opcode 0x37 (immediate to byte `[reg+9]`, linear sweep of
  `.text`) is found only at `L09231`, in `R0645` (the town roster
  refresh, called at `L09216` from `R1318`, which the campaign state
  machine calls at `L09232`). The only such store of 0x38 is at `L09233`, in
  `R1653`, whose sole caller is `L09217` in `R0786`.
- `R0786` is slot `0x84` of the inn view vtable (`L09234`); commands
  0x445 and 0x446 reach it through `R0716`, which calls vtable `+0x84`.
- The sweep finds no call to `R1653` or `R0645` other than those; the
  SAV loader is not a caller. `SAV-615` records hired Humans keeping the same
  idents from a city save into a mid-mission save.
- Immediates 0x37 and 0x38 are also pushed at 21 other sites
  (`evidence/imm-store-37-38.txt`); 5 sampled in review are string ids, sizes or call arguments, and 16 were not cleared.

**Confidence.** Reading is EN `rom.exe` 942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03; the RU image was not compared.
Medium: the sweep is static and covers immediate stores only: register-sourced stores and messages prefilled in a data object were not searched; the inn close handler is reached
by a virtual slot whose callers were not enumerated, and the 21 pushes were not
read.

**Unknown.** Whether entering a mission without opening the tavern sends 0x38;
all callers of vtable slot `+0x84`.

### SAV-1139

- Party Unit fields in `game9607.sav` object 189 and `game0014.sav` object 26:
  type dword 2 (bytes 47..50), companion dword `+0x144` 0 (bytes 51..54), class
  key row 27, type word 0x1b, one-actor Player group, owner reference to the
  hero Player.
- The campaign record is identical in all five saves: hire flags
  `[0,1,0,0,0,1,0,0,0,0,0,0,1,1,0]` (types 2, 6, 13, 14), pool arrays
  `[1,1,1,1,1,4,4,3,3,3,0,3,4,3,0]`. The siege type has flag 1 and pool 1.
- Between the two saves the Unit records differ in: position bytes, run id
  (14 and 97), identity key, owner pointers, the last 2 bytes of `raw-a6` and
  `raw-114`, bytes 0..1 of `raw-pointed-154`, bytes 113 and 144..147 of
  `raw-pointed-158`, plus byte 18 of `scalar-run-1`. The type dword, the
  companion dword, class key and type word are equal
  (`evidence/save-unit-diffs.txt`).

**Confidence.** Reading is EN `rom.exe` 942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03; the RU image was not compared.
Medium: an absence of a link over two Unit records and one
campaign record; the meaning of the differing raw bytes is unread.

**Unknown.** What the differing `raw-pointed-154` and `raw-pointed-158` bytes
are; whether any is a runtime pointer.

### SAV-1140

- `game9607.sav` (a kit-written city save): one siege Unit, Humans typed 6, 13
  and 14 in counts 4, 4 and 3, money 2711000.
- `game0000.sav` (original mission 120, after loading it): hero Player holds two
  Units, run ids 14 and 145, type dwords 0 and 2, in separate one-actor groups;
  money 2711000. Fresh Humans have run ids 147..168.
- `game0014.sav` (original city resave): one Unit, run id 97, type dword 2,
  money 2711000. `game0009.sav`: one Unit, run id 97, money 2422000, 289000
  lower. `game0010.sav` (original mission save): one Unit, run id 144, type
  dword 2, money 2422000 (`evidence/save-player-money.txt`).
- Hypothesis for equal money in the kit save and `game0000.sav`: a refund by
  command 0x37 followed by an equal charge by command 0x38 (`SAV-1136`). It
  needs `[tavern+0x9c]` populated when 0x37 runs; `R1142` zeroes it at entry
  and `R1555` at exit, and no save holds a Tavern object, so the precondition
  is unread. The alternative is that no charge ran. The cause of the 289000
  difference in `game0009.sav` was not read.

**Confidence.** Reading is EN `rom.exe` 942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03; the RU image was not compared.
Medium: five saves; the money mechanism is a hypothesis.

**Unknown.** The original runs that produced `game0009.sav`, `game0010.sav` and
`game0014.sav`; the cause of the 289000 difference.

## Stored physical shot and burst records

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1141 | In three original EN mission-150 saves, the records of pictures 1, 2 and 5 carry ActionTarget 340, the runtime id of one Human whose saved position equals their ActionX and ActionY; no shooter is stored. | High / Medium | ● active | [EXP-0430](../experiments/EXP-0430-physical-projectiles/) |
| SAV-1142 | In those saves the rock's and the bolt's x, y equal the shooter's position plus 8 times (class ShootOffset pair minus Center); segments are the shooter-target distance divided by 200, a divisor taken from SAV-1130. | High | ● active | [EXP-0430](../experiments/EXP-0430-physical-projectiles/) |
| SAV-1143 | In those saves each shooter's simulation countdown expires on the tick its projectile is collected: damage tick minus creation tick equals the shot's total segments (6, 7, 7) for an archer, a catapult and a crossbowman. | Medium | ● active (amended) | [EXP-0430](../experiments/EXP-0430-physical-projectiles/) |
| SAV-1144 | The picture-13 record has the leaves client arm 0x86 writes: ActionTarget 0, x and y equal to ActionX and ActionY at a cell centre, actionphase starting at -1, and 22 calls in all, the constant `R0635` sends for spell 2. | High / Medium | ● active | [EXP-0430](../experiments/EXP-0430-physical-projectiles/) |
| SAV-1145 | The picture-13 Fire_Ball burst at the catapult's target cell is created 4 ticks after the catapult's damage tick; the rider path leads to it, but a cast by another actor is not excluded, so the siege rider stays Medium. | Medium / Unknown | ● active (amended) | [EXP-0430](../experiments/EXP-0430-physical-projectiles/) |
| SAV-1146 | In the same three saves FreeIndex (28, 29, 31) equals max(IDs) + 1 and ids 27 and 28 were free when 29 and 30 were built, so one shared counter allocates and no id is reused; the pair 29, 30 is saved newest first. | High / Medium | ● active | [EXP-0430](../experiments/EXP-0430-physical-projectiles/) |
| SAV-1147 | A unit shot starts with actionphase 0, so its built segment count is actionphase plus actionsegments; a record built by client arm 0x86 or 0x8b starts at -1, so that count is actionphase plus actionsegments plus 1. | High | ● active | [EXP-0430](../experiments/EXP-0430-physical-projectiles/) |
| SAV-1149 | In `game0024` the catapult shot starts a SpellTransport countdown of 1448/384 = 3 ticks before the Fire_Ball blast; the observed 4-tick gap to the burst record is that countdown plus one tick whose cause was not read. | Medium / Unknown | ● active (amended) | [EXP-0432](../experiments/EXP-0432-projectile-pictures/) |
| SAV-1150 | A SAV does not carry tile-word bit 0x2000: the Fog store writes bit 15 only and no other tile-plane store was found; the clear at `L02344` is the view-plane parser's loop, so that a load starts unmarked is Medium. | High / Medium | ● active | [EXP-0432](../experiments/EXP-0432-projectile-pictures/) |
| SAV-1151 | No instruction of the SAV store or load routines addresses the view hash at +0xa7c; 13 instructions in 8 other routines do (the driver and view code), so the hash is not saved. | Medium | ● active | [EXP-0432](../experiments/EXP-0432-projectile-pictures/) |
| SAV-1152 | Of the 121 `.sav` files under the saves root, 80 have a Projectiles store and 5 of those hold a record (6 records, pictures 1, 2, 5, 10, 13); none holds picture 14 or more. | Medium | ● active | [EXP-0432](../experiments/EXP-0432-projectile-pictures/) |
| SAV-1153 | A unit shot's first driver call is on its creation tick, and the three saved shots fit creation = swing start (inferred) + ShootDelay if the swing message is delivered in the start tick's own timer step. | High / Medium | ● active (amended) | [EXP-0441](../experiments/EXP-0441-shot-timing/) |
| SAV-1154 | The code chain puts the Fire_Ball burst record 3 ticks after the rider tick if the transport is last in the effect list at its fire tick, else 2; the saved catapult case shows 4, so one to two ticks stay unexplained. | Medium / Unknown | ● active (amended) | [EXP-0441](../experiments/EXP-0441-shot-timing/) |
| SAV-1155 | For a target whose size byte is 1 or less, the Fire_Ball transport's target point is the target's cell and fine bytes; its countdown is the truncated Euclidean distance of those words divided by spell parameter 7. | High / Medium | ● active (amended) | [EXP-0441](../experiments/EXP-0441-shot-timing/) |
| SAV-1156 | No saved SpellTransport exists at ticks 3098 to 3101 of the mission-150 battle, so the countdown and the effect list at the fire tick are not observed; whether the transport is last in the list stays Unknown. | Medium / Unknown | ● active | [EXP-0453](../experiments/EXP-0453-shot-burst/) |
| SAV-1157 | A unit shot's offset (damage tick minus creation tick) is charge - ShootDelay + floor((d * 256 + 128) / 200) for d above 1, else charge - ShootDelay; archer, crossbowman and catapult fit saved shots, the other five classes are unobserved. | Medium / Unknown | ● active | [EXP-0453](../experiments/EXP-0453-shot-burst/) |
| SAV-1159 | The terminal route and the new-game arm clear the campaign record identically and differ only in frame state: new game sets start-mode words; the terminal route clears the client map and score cache. | High / Medium / Unknown | ● active | [EXP-0461](../experiments/EXP-0461-ending-remainder/EXP-0461.md); FAME-033 |

### SAV-1141

- Inputs: `game0022.sav` (head counter 3087), `game0023.sav` (3095) and `game0024.sav` (3109), one battle on mission 150, saved by the original EN game. Each holds 143 Unit and 47 Human actors.
- Records: Prj27 picture 1 (`game0022`), Prj28 picture 5 (`game0023`), Prj30 picture 2 and Prj29 picture 13 (`game0024`). Prj27, Prj28 and Prj30 have ActionTarget 340. Prj29 has ActionTarget 0.
- Runtime id 340 is a Human (type word 19) at x 6016, y 24448, which is cell (23, 95), sub-cell (128, 128). It is at the same position in all three saves. Its health is 152, 110 and 25.
- ActionX and ActionY are 6016 and 24448 in all four records. Z and ActionZ are 0.
- Exactly three actors target 340 in each save, with Unit and Human actors joined by the saved-address key at `+0x5c`: Human 327 (type word 14, class Human Archer, Projectile 1), Human 331 (type word 15, class Human CrossBowMan, Projectile 2) and Unit 47 (type word 26, class Catapult 1, Projectile 5).
- Each record's picture equals the Projectile of exactly one of them. The record names its target; the shooter is the actor with that target and class Projectile (`SAV-1130` stores no shooter).

**Confidence.** High for the target id, the position equality and the three-actor census over the 190 actors of each save. Medium for the shooter: it follows from uniqueness over those actors and from `SAV-1142`, not from a stored field. One battle, three saves, one target.

**Unknown.** The weapon or item used: no leaf stores it.

### SAV-1142

- Class constants come from `units.reg` (`evidence/unit-classes.txt`): Human Archer Center (64, 78), ShootOffset pair 4 (65, 30); Human CrossBowMan Center (64, 78), pair 4 (68, 42); Catapult 1 Center (64, 64), pair 5 (82, 31).
- The unit shot (`R0603`, `SAV-1130`) writes the start from the shooter's x, y and the pair `ShootOffset[(dir-8)&0xe]` minus Center, times 8.
- Prj28 (picture 5, 5 calls run, 2 segments left): over start offsets of -60..60 steps of 8 around Unit 47, one start reproduces the saved x, y under the flight rule of `ANIM-101`. That start is (+18, -33) steps, the pair (82, 31) minus (64, 64) (`evidence/start-search.txt`, `startsearch.py`). The projectile's saved dir 2 would select index 10; no shooter facing is stored, so the index is not tested.
- Prj30 (picture 2): the one start is (+4, -36) steps, the pair (68, 42) minus (64, 78). The projectile's saved dir 0 would select index 8; not tested.
- Prj27 (picture 1) ends on its target, so its start is not constrained. Its start is (5512, 25344) under pair index 8 taken from the saved dir; the model ends on the saved leaves for any start, so this shot does not test the start formula.
- Segments: integer square root of the squared distance from the shooter's saved position to (6016, 24448), divided by 200: 1378 gives 6, 1448 gives 7, 1557 gives 7. Each equals actionphase plus actionsegments of its record. The data alone admit any divisor from 197 to 206; 200 comes from the instructions read in `SAV-1130`, which the data confirm.

**Confidence.** High: the start formula is tested on two of three shots (rock and bolt), each a unique solution of the committed search that lands on a table entry; the three segment counts match the code-derived divisor.

**Unknown.** ShootOffset pairs the saves do not select; classes other than 14, 15 and 26.

### SAV-1143

- Method: a record with phase 0 start has run `actionphase` driver calls; the first call is on its creation tick; the call that finds segments 0 collects it. A shooter in sub-phase 5 (`+0x58`) with countdown n (`+0x6c`) applies on counter + n (`MAGIC-CASTTICK-030`). The head counter dates each save.
- Prj27: created 3082, collected 3088; Human 327 applies at 3087 + 1 = 3088.
- Prj28: created 3091, collected 3098; Unit 47 applies at 3095 + 3 = 3098.
- Prj30: created 3107, collected 3114; Human 331 applies at 3109 + 5 = 3114.
- The three shooters have different charge, ShootDelay and distance; in each, creation-to-damage equals the record's total segments (`evidence/timeline.txt`). This is observation at one distance per class.

**Confidence.** Medium: three exact matches, but the creation tick assumes the client driver runs once per head-counter tick, which `SAV-1132` records as an inference. One distance per class.

**Unknown.** Whether the match holds at other distances and at d = 1, where `SAV-1132` finds the damage tick before the release tick.

**Amended.** The assumption that the client driver runs once per tick and that a record's first call is on its creation tick holds for unit shots and for dispatcher-built records (`ANIM-113`, `SAV-1153`); the grade stays Medium for the one distance per class.

### SAV-1144

- Prj29 (`game0024`): picture 13, x 6016, y 24448, ActionX 6016, ActionY 24448, ActionTarget 0, actionphase 7, actionsegments 14, phase 3, lastaction 1, action 1, dir and actiondir 0, actionspell 0.
- Client arm 0x86 (`L03098`, `evidence/disasm-client-arm-86.txt`) builds a record with picture `msg+0xc` (odd only), x and y as cell `msg+0xd`, `msg+0xe` times 256 plus 128, z 0, ActionX/Y/Z copied from x, y, z, ActionTarget 0, actionphase -1 and actionsegments `msg+0xf`. It takes its id from the counter at world +0xa0c.
- Total life: 7 + 14 + 1 = 22 driver calls, which is 0x16.
- `R0635` writes `msg+0xf = 0x16` with opcode 0x86 when the effect's `+0x0c` (spell id) is 2, and picture `effect+0x0e`; `MAGIC-PIC-026` gives `2*2+9 = 13` for spell 2. The other 0x86 sender `R0638` passes a caller byte (16 or 18, `MAGIC-BURSTLIFE-034`), and `R0617` writes even pictures.
- Arm 0x8c builds picture 36 with 13 segments. Arm 0x8b is fed by the even-picture cast senders (`MAGIC-DELIVER-035`).
- Picture 13 is projectiles.reg ID 13, `fireexpl`, 11 phases, no rotation, 128 by 128 (`evidence/projectiles-reg-rows.txt`).

**Confidence.** High for the leaf agreement with arm 0x86 and the 22-call life. Medium that `R0635` sent it: 22 and picture 13 match no other sender read, but no message was observed.

**Unknown.** Whether a sender not read writes 0x16.

### SAV-1145

- Timeline (`evidence/timeline.txt`): Prj28 (catapult rock) is collected and Unit 47 applies at tick 3098. Prj29 is created at 3102: 8 driver calls by 3109 from a start of -1. Prj30 is created at 3107.
- Ids are consecutive: arrow 27, rock 28, burst 29, bolt 30 (`SAV-1146`). The burst sits at the cell centre of target 340, the cell the rock flew to.
- Rider path (`evidence/disasm-sim-damage-and-rider-R0246.txt`): after damage, `R0246` copies the weapon spell reference (`actor+0x74`) and calls `R0002`, then `R0003`. Spell 2 reaches the default arm `L03046`, where Distribution (row parameter 8) other than 1 builds a 0x50-byte area effect with `+0x0c` the spell id and `+0x0e` `2*id+9`. Its tick `R0629` reaches `R1086` and `R1067` when its stage byte is 0; `R1067` calls `R0635(effect, 1)`.
- The Catapult row's weapon is `Boulder Thrower{castSpell=Fire_Ball:70}` (`EXP-0428` `shot-classes.csv`).

**Confidence.** Medium for the rider as the source: path, adjacency and place agree, but the branch conditions at `R0631` and `R0632` were not read. Unknown for the 4-tick gap.

**Unknown.** Why the burst follows the damage tick by 4. Whether `R0637` at `L05631` sends a second 0x86 for spell 2 when the area effect ends (`R0635` branches on spell id, not on its flag). Whether a Fire_Ball cast by a mage at the same cell builds the same record. Next question: a save held between the damage tick and the burst, plus the area effect's first-tick delay and lifetime (`+0x4c`).

**Amended.** The third bullet names the cloud arm (`R1086`, `R1067`) for Fire_Ball; the effect's mode word is 0 and its tick takes the blast arm `R0630` (`ANIM-111`). The rider is not the only source: one creator serves every actor (`ANIM-111`). The gap is decomposed in `SAV-1149`. `R0637` is a cloud-only arm and sends no second 0x86.

### SAV-1146

- Saves: `game0022` FreeIndex 28, IDs [27]; `game0023` FreeIndex 29, IDs [28]; `game0024` FreeIndex 31, IDs [30, 29].
- Builders: Prj27, Prj28 and Prj30 are unit shots; Prj29 is an arm-0x86 build. All take the id from the same word at world +0xa0c and set FreeIndex to id + 1 (`SAV-1131`; `L03076`, `L09235`).
- Ids 27 and 28 were collected (3088, 3098) before 29 (3102) and 30 (3107) were built. A lowest-free-id allocator would have reused them.
- Ids 29 and 30 are both in bucket (id >> 4) mod 17 = 1. Insertion at the bucket head gives [30, 29], as saved. Ascending id order, or insertion order, would give [29, 30].
- Every FreeIndex here equals max(IDs) + 1.

**Confidence.** High for the shared counter and the exclusion of id reuse. Medium for FreeIndex as last-built id plus 1: it is observed only as equal to max(IDs) + 1, which all three saves satisfy. Head-first order rests on the serializer reading in `SAV-1131`; this data adds one same-bucket pair, which head insertion, newest-first and descending-id order all fit.

**Unknown.** Order across buckets: all ids here share bucket 1.

### SAV-1147

- Unit shot `R0603` sets `actionphase = 0` (`L09236`). Arms 0x86 and 0x8b set `actionphase = 0xffffffff` (`L02568`, `L02569`).
- Each driver call adds 1 to actionphase and subtracts 1 from actionsegments (`ANIM-101`), so their sum is constant.
- Prj27, Prj28 and Prj30 sum to 6, 7, 7, which equal the segments `SAV-1142` computes. Prj29 sums to 21, one less than its 22 calls (`SAV-1144`).
- Prj266 of `SAV-1134` (phase 2, segments 3, picture 10) sums to 5. The cast spawner starts at actionphase 0 (`MAGIC-CASTSPAWN-033`), so it was built with 5 segments; an arm build would have built 6.

**Confidence.** High: the initial stores are instructions. The three unit shots (Prj27, Prj28, Prj30) are the only independent agreeing observations; Prj29 is checked against its own 22-call constant and the Prj266 build count is inferred.

### SAV-1149

- Shot: Unit 47 (catapult) at (4992, 25472), target Human 340 at (6016, 24448), both from the saved `game0022` actor records (`evidence/transport-countdown.txt`). `R1144` returns the truncated square root of the summed squared word differences: 1448. Fire_Ball speed is 384 (`ANIM-111`). The transport's countdown is the quotient, 3 (`R0633`, IDIV at `L09237`).
- The countdown is 3 after construction and the transport's third tick fires it (`R0634`); the inner effect is then added to the effect list and its blast follows.
- Order inside one simulation sub-tick: message drain `R0191`, then `R0426`, which makes a virtual call (vtable +0x18) for each element of one list and then runs the effect walk `R0641`, then the queue flush `R0611` (`evidence/disasm-sim-subtick-R0193.txt`). Whether that list holds the actors of `R0037`, and whether a node appended during the effect walk is visited in the same walk, were not shown.
- Observed (`SAV-1145`): damage tick 3098, burst record created at 3102. If the transport's first tick is the tick after creation and the blast is sent on the tick the transport fires, the blast is sent on 3101 and a drain on the next tick creates the record at 3102.

**Confidence.** Medium: the countdown value is arithmetic on saved positions and instruction constants, and it fills three of the four ticks. Unknown for the fourth: the walk order and the client's delivery latency were not read.

**Unknown.** Whether the inner effect first ticks the tick after the transport fires: if so, the observed gap needs no leftover tick. A save held at ticks 3098 to 3101 of the mission-150 catapult battle would show the transport's `+0x4c` and decide which order holds.

**Amended.** The third bullet's open points are read in `EXP-0441`: the list of `R0426` is the actor list and the actor pass precedes the effect walk (`ANIM-113`), a node appended during the last visit of a walk waits for the next walk (`ANIM-114`), and the dispatcher that builds the burst record runs in the sending tick's timer step, so the burst record is built in the sending tick's timer step if the flush-to-queue path is same-tick, which was not traced; the fourth bullet's "drain on the next tick" stays a candidate. The code chain gives the record 2 or 3 ticks after the rider tick (`SAV-1154`).

### SAV-1150

- Bit 0x2000 of the tile word is set by the client mask arm for wall_of_fire (`L05664`, `MAGIC-WALLFIRE-058`) and by the Fire_Ball burst's driver arm (`ANIM-103`); it is tested at `L05660`.
- The SAV Fog store writes `tile & 0x8000` only and the load ORs the state back (`SAV-FOG-061`). So the Fog section cannot carry bit 0x2000. No other store of the tile plane into a SAV was found.
- Clears of the bit (`evidence/imm-clear-mask-dfff-e000.txt`, immediates 0xdfff, 0xe000 and 0xffffdfff over `.text`): `L02344 AND word [edi+ecx*2], 0xdfff` in the loop that ends the light landscape parser (`L02343` to `L09238`, `evidence/disasm-alm-tile-clear-L02343.txt`), once per cell over width times height; and `L09239` in the client 0x9b arm, which clears and rewrites the bit from the server's static plane bit 0x10 (`R1655`, `R1656`; `evidence/disasm-resync-R1655.txt`). The other two hits (`L09240`, `push 0xe000` at `L09241`) are not tile words.
- `ALM-TILEVIEW-122` scopes the parser's loop to the view plane and leaves later ownership and ordering Unknown. Whether that plane is the one the Fog store reads (`L07964`), and whether the clear runs before a SAV's state is applied, was not shown.
- Bit 13 is also ALM-authored on 7 464 cells over 38 maps (`TERR-TILE-044`), so the clear is not by itself evidence about a burst mark.
- `evidence/imm-2000-tile-bit.txt` lists the 23 `.text` hits of the immediate 0x2000. The setters and tests read here are `R1656`, `R0483` and the driver arm; `L09242` and `L09243` act on window flag words. The other hits were not classified.
- The player-join routine `R0131` sends the 0x9b resync (`L09244`).
- Method step 3 (compare saved terrain words around the burst cell) did not apply: a SAV carries no tile plane except the Fog runs.

**Confidence.** High that the Fog runs hold bit 15 only. Medium that no other producer writes the plane into a SAV, the scope `SAV-FOG-061` already states, and Medium that a load starts without the mark: that rests on the view-plane clear and the 0x9b resync, whose ownership against the Fog plane is open.

**Unknown.** Whether a SAV taken between the burst's call 4 and call 8 resumes with the cells unmarked: no corpus save was taken in that window. A save held after the 0x86 burst's call 8 (tick 3103 or later in the mission-150 battle) would test it directly.

### SAV-1151

- Instructions that address displacement 0xa7c: `lea` at `L09245` and `L09246` in the projectile driver `L09247`; `add reg, 0xa7c` at `L09248`, `L09249`, `L09250`, `L09251`, `L09252`, `L09253`, `L09254`, `L09255`, `L05663`, `L09256` and `L09257` (`evidence/disp-a7c-hash.txt`, `evidence/imm-a7c-hash.txt`).
- Their routines start at `R0391`, `L09211`, `R1657`, `R0379`, `L09258`, `R0483`, `L09247` and `L09259` by prologue. None of the 13 addresses lies between `L09260` and `R1635`, the range of the SAV store `R0084`, the load `R0099` and `R0414` (`SAV-FOG-061`, `SAV-DOC-053`).
- The hash holds keys `(y << 8) | x`. The driver arm 0x0d writes it at the burst's call 4 (`ANIM-103`) and `L05663` writes it for wall_of_fire (`MAGIC-WALLFIRE-058`).

**Confidence.** Medium: absence of a serializer reference by displacement. The scan sees memory operands and `add reg, imm`; a block copy of a containing object, or a displacement formed from two additions, is not seen.

**Unknown.** The hash's consumer and what a loaded game does with it. A save held after the burst's call 4 and before call 8 would hold its content if it were serialized.

### SAV-1152

- `corpusscan.py` walks every `.sav` under the saves root with the `&YA1` reader of `EXP-0430`: 121 files, 80 with a Projectiles store, 40 loaded as `&YA1` documents without a Projectiles section (1 527 to 5 674 bytes, no world section) and 1 reader failure; 75 of the 80 have an empty store and 5 have a record (`evidence/save-corpus-scan.txt`).
- The 5 files hold 6 records with pictures 1, 2, 5, 10, 10 and 13; two of the files are byte duplicates. The three `game0022` to `game0024` saves hold pictures 1, 5, 2 and 13.
- No record has picture 14 or more.

**Confidence.** Medium: a bounded negative over the 80 files that have a Projectiles store, from one corpus; the 40 files without the section hold no world, and the one reader failure is not counted.

**Unknown.** A mission-150 save held while a cast of Wall of Fire, Fire Sacrifice, Acid Stream or Meteor Storm is in flight, or while its burst runs, would hold a picture of 14 or more.

### SAV-1153

- Swing start `R0245` calls `R0549` (`L02725`), which sends the swing message in the actor pass of the start tick S. The flush `R0611` closes the sub-tick; the dispatcher of the same timer step delivers the message, and the sweep of `R0334` in that step runs the action driver `R0548` with phase 0 (`ANIM-113`). Its action-7 arm builds the shot when the phase equals ShootDelay, so the shot is built on tick S + ShootDelay.
- The shot is stored in `+0x9d4` during the actor sweep, and `+0x9d4` is swept later in the same pass, after its count is read (`L09261`), so the shot's first driver call is on its creation tick (`ANIM-113`).
- The start tick S is not read from a save; it is inferred from the countdown below. Premises read in code: the swing start `R0245` stores `(d * 256 + 128) / 200` as an output only when `d` (from `R0247`) is greater than 1 (`L09262` to `L09263`); the actor tick calls it at `L05237`, then sets sub-phase 5 with the countdown `+0x6c` = charge (`+0x134`) plus that output (`L09264` to `L04380`, `evidence/disasm-swing-countdown-L13164.txt`).
- Damage: the shooter's countdown after the swing is charge plus that term. For the three saved shots (`SAV-1143`) damage tick minus creation tick equals charge - ShootDelay plus that term: archer 20 - 21 + 7 = 6 (d = 5), crossbowman 8 - 9 + 8 = 7 (d = 6), catapult 2 - 0 + 5 = 7 (`evidence/shot-class-columns.txt`, `EXP-0428` `shot-classes.csv`). The three fits compare the first driver call with the damage tick; a one-tick delivery delay would shift all three by one, so they support the offset only under same-step delivery (`ANIM-113`, Medium).

**Confidence.** High for the sweep visiting the shot in its creation tick (instructions). Medium for the offset: S is inferred, the creation and damage ticks come from the promoted `SAV-1143` timeline (no save was re-read), one distance per class, and delivery in the start tick's own timer step is untraced.

**Unknown.** The offset at other distances, and for a swing whose message the dispatcher defers.

**Amended.** `EXP-0453`: the offset at other distances and classes is tabulated (`SAV-1157`, where the catapult's distance is 4); same-step delivery is read (`ANIM-121`).

### SAV-1154

- Chain (`ANIM-113`, `ANIM-114`, `ANIM-115`): the rider runs in the actor pass of tick R (`SAV-1145`: 3098). The transport is appended then and the effect walk of the same sub-tick visits it, so its third visit and its fire are on R + 2. The inner effect is appended during that walk. If the transport is the last list node at its fire tick, the inner effect first ticks on R + 3, sends 0x86 on R + 3, and the record exists on R + 3 (given same-step delivery). If another effect was appended after the transport before its fire tick, the walk reaches the inner effect in the fire tick and the record exists on R + 2.
- Observed (`SAV-1145`, `game0024`): the burst record is created at 3102, R + 4: one tick off the first branch and two off the second.
- The earlier reading (`SAV-1149`) put the blast on R + 3 and the record one drain later. Whether the flush-to-dispatcher path delivers in the sending timer step was not traced (`ANIM-113`), so a one-tick delivery delay stays a candidate for the leftover tick.

**Confidence.** Medium: each link is an instruction. Unknown for the fourth tick: neither the transport's first visit, the rider's tick dating (damage tick = counter + countdown) nor the creation formula of arm 0x86 records (phase starts at -1, `SAV-1147`) was shown to hold it.

**Unknown.** Candidates for the leftover tick: a one-tick delay on the untraced path from the flush `R0611` to the queue `R0509` drains; the transport's first visit on R + 1; the rider tick dating. The effect list's contents at the fire tick are not read. A save held at ticks 3098 to 3101 of the mission-150 catapult battle would show the transport's `+0x4c` and decide the second candidate.

**Amended.** `EXP-0453`: the delivery delay from the flush to the dispatcher is 0 on the local path (`ANIM-121`), which excludes the first candidate on that path and leaves it open for a non-local path; the effect list at the fire tick is not read (`SAV-1156`); the extra tick is still unexplained.

### SAV-1155

- Transport constructor `R0633`: inner effect at `+0x44`, countdown `+0x4c` = `R1144(caster position, effect +0x10)` divided by the signed speed word (`IDIV` at `L09237`). A second constructor, `R1646`, keeps the inner effect at `+0x48` and does the same division.
- `R1144` reads the words of cell * 256 + fine for both positions (`R0165`, `R0166`) and returns the truncated square root of the summed squares (`evidence/disasm-distance-R1144.txt`). The speed is spell parameter 7 (`ANIM-111`).
- In `R0003`, for a non-null target the local position is filled by the copy `R1010` (`L09265`; it copies bytes 0 to 5 and 8 to 11 of the target's position structure, `evidence/disasm-position-copy-R1010.txt`) and `L05872`. When the target's virtual slot +0x1c byte is greater than 1 (`L06145` to `L09266`), `L09267` to `L06146` rewrite it through `L06144` with the cell word plus `(size - 1) * 128`, minus 1, per axis (`R0863`, `R0862`). When the target is null (`L06147`) the point comes from `R0289`. The first two branches are read; the third is not characterised.
- The local position reaches the inner effect through `R0652` (`L06148` to `L03047`), `L05979` and `R1009`, which copy it with `R1010` into the effect's `+0x10` (`L07895`, `L02083`); the transport constructor reads that field and the caster position (`L09268`). `R1144` reads cell * 256 + fine from both (`R0165`, `R0166`), so for a size-1 target the fine bytes enter the distance.
- The saved case of `SAV-1149` has both endpoints at cell centres (fine 128), so it separates neither fine position nor cell centre, and says nothing about a target of size greater than 1.

**Confidence.** High for the metric, and for the copy of the fine bytes in the size-1-or-less branch (instructions). Medium for the divisor's source, which `ANIM-111` states. Unknown for the null-target branch.

**Unknown.** The multi-cell target branch (`L06144`) and the null-target branch (`R0289`) in effect; a shot at a target between cell centres would show the countdown, and no such save was read.

**Amended.** `EXP-0453`: the multi-cell and null-target branches are read (`MAGIC-245`).

### SAV-1156

- Corpus: 120 `.sav` files (the 121 of `SAV-1152` minus one engine-written save, excluded as not ROM1 evidence), 119 readable. The raw class name SpellTransport occurs in `game0018.sav` (head 3452, mission 40, a PointEffect transport) and its duplicate; AreaEffect in 2 files (heads 63, 1791). No effect class name occurs in `game0022`, `game0023` or `game0024` (heads 3087, 3095, 3109), the saves of the mission-150 battle (`evidence/corpus-effect-classes.txt`).
- In `game0018` the transport's last word (`+0x4c`) is 4; its age is not saved, so it cannot date a transport.
- Chain: the transport's countdown is distance divided by parameter 7 (`SAV-1155`). Each visit (`R0634`) decrements `+0x4c` and fires at 0 or below, appending the inner effect and setting the reap flag (`evidence/disasm-transport-fire-R0634.txt`, `evidence/vtables-L05917-L05352.txt`). With the rider at R = 3098 and a countdown of 3, the visits would read 2 at R, 1 at R + 1 and fire at R + 2, if the first visit is in the creation tick.
- Effect list at the fire tick: the append goes to the tail (`ANIM-114`), so a node appended after the transport and before R + 2 = 3100 makes the inner effect tick in the fire tick. The saves nearest that span are at 3095 and 3109, none inside 3096 to 3100, so the effects appended in those ticks (by any caster or another siege shot) are not known and whether the transport is last stays Unknown.
- The saved record is at 3102 (R + 4; `SAV-1145`). The delivery delay is 0 on the local path (`ANIM-121`), which excludes that candidate for the extra tick on that path. Remaining candidates: the transport's first visit, an unread element of the chain, and a non-local path.

**Confidence.** Medium for the countdown arithmetic and the visit order (instructions). The countdown value at 3098 to 3101 and the list order at the fire tick are not observed.

**Unknown.** The countdown at ticks 3098 to 3101; the effect list contents at the fire tick; the extra tick of `SAV-1154`. A save at one of those ticks would decide them.

### SAV-1157

- Swing start `R0245` stores `(d * 256 + 128) / 200` as the countdown extra only when d (from `R0247`) is above 1; `R0247` returns max(|dx|, |dy|) over cell * 256 + fine words, minus (size1 + size2) * 128 - 256, as 1 when at most 0x180 and else (value + 0x40) >> 8 (`evidence/disasm-swing-start-R0245.txt`).
- Actor sub-phase 5 loads `+0x6c` = charge (`+0x134`) plus the extra and the damage applies when it reaches 0. In action arm 7 (`L02711`) the phase compare against ShootDelay precedes the increment, so the shot is built at the start tick plus ShootDelay (`ANIM-113`, `SAV-1153`; `evidence/disasm-driver-arm7-L02693.txt`, `evidence/disasm-actor-tick-L13165.txt`).
- Offset = charge - ShootDelay + term(d). `evidence/shot-offset-table.txt` lists d = 1 to 12 for the 8 distinct (Projectile, charge, ShootDelay) classes of the 50 shot rows in `EXP-0428` `shot-classes.csv`: for example Human Archer 2, 3, 4, 6 at d = 2 to 5, Catapult 2 at d = 1 and 9 at d = 5. The saved shots fit: archer d = 5 gives 6, crossbowman d = 6 gives 7, catapult d = 4 gives 7 (the term 5 of `SAV-1153` is term(4); the catapult's d is inferred from that term, no saved position was re-read). The other five classes of the table are the same code applied to rows no save shows.
- For a class with ShootDelay above charge the offset at d = 1 is negative (Human Archer -1): the damage tick would precede the creation tick. That case was not observed.

**Confidence.** Medium: code plus three saved distances, one per class; classes other than archer, crossbowman and catapult are not observed.

**Unknown.** The offset of any other class in a save; the d = 1 case of a class with ShootDelay above charge; swings whose message the dispatcher defers.

## Actor damage-kind bytes and their writers

The two `HERO-` cards in this topic belong to the Heroes ledger and move to `hero.md` when that ledger is rewritten as cards.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| HERO-MODDK-161 | `R1562` is the only routine storing a constant into actor `+0x10e..+0x113` (all six bytes and six protections to 100); the other writers zero, copy through the archive, or add zero from shipped Armor and Shield. | Medium | ● active | [EXP-0433](../experiments/EXP-0433-effect-arms/) |
| HERO-DKIDX-162 | A unit-resolver strike reduces damage by the target's byte at `+0xce` plus the attacker's own active slot (attacker `+0xb6`), so slots 1..5 read `+0xcf..+0xd3`. | High / Unknown | ● active | [EXP-0433](../experiments/EXP-0433-effect-arms/) |

### SAV-1159

Instrument: as FAME-030; `evidence/reset-fields.tsv`, listings `terminal_reset`, `new_game_arm_head` (`L08085..L09269`) and `new_game_arm_body`, and the new-game store and call assertions in `evidence/checks.tsv`.

The terminal route `R0816` and the new-game arm differ only in frame state. New game only: the frame state word at `+3e0` is zeroed, `+6b8` is set to 1, `+3c4` is set to 1, `+6bc` to 2 and flag bits at `+494` and a string at `+430` are set through `R1382`, `R1658(1)` and `R1659` run, and `+65c` is set to 1. Terminal only: the client map array at `[frame+d0]+9b8` is cleared, the score cache at `[frame+d0]+3f54` is deleted and zeroed (`L03726`), a call to `R0679` runs on a session object with a string that was not identified, the root is rebuilt through `R1660`, the menu panel is added, bit 0x80 is set in `+3dc` and `+41c` is zeroed. The campaign record is cleared identically (FAME-033).

**Confidence.** High that the record is cleared identically; Medium for the frame-level lists, which are a read of the two arms without a run. Unknown: what `R0679` does with the unidentified string.

### HERO-MODDK-161

- Sweep: every byte, word, plus dword store with a register-based
  displacement in `0x10b..0x113` over the whole `.text` of `rom.exe` (capstone
  linear sweep with data resynchronisation, 47 hits; EN and RU images are
  byte-identical). One hit lies in a routine of the actor families
  (`L00001`, `L00002`, `L00003` constructors): `L09270`, a byte store of 0x64 at
`+0x10e`, in `R1562`. Of the other 46, two are the
  dispatcher's word stores at `+0x10c` (`ITEM-EFFKEY-147` kind 25) and 44 are
  dword stores at `+0x10c` or `+0x110` in 32 functions, none of which stores
  an actor-family vtable constant (the functions, their callers and vtable
  references are in `ghidra-owner-classes.txt`).
- `R1562` loops `i = 0..5` writing the u16 `100` at `actor+0x102+2i`
  and the byte `100` at `actor+0x10e+i`, then calls `vt+0x50`. Its callers
  are 4 sites in 3 owners. `R0430` (sole caller `R0061`) calls
  it at `L09271` for the `self` target and at `L09272` for each `army`
  actor, in the arm taken when the order text matches the string `+god`
  (`push L09273` at `L09274`, beside the strings `#modify `, `self`,
  `army`, `+spell `). `R0823` (`L09275`) and `R0065`
  (`L09276`) call it on a new actor when a `+0x12c` or `+0x134` dword of
  their source record is nonzero (`+0x134` only for `R0065`).
- Zero fill: five actor constructors call `R0844`, which clears the
  whole 0x40-byte modifier object.
- Archive: `R1570` copies the 0x40-byte modifier object through the
  archive, so a nonzero value can also arrive from a loaded SAV in the direction chosen by a test on `R1123`; its one direct
  caller is `R0210`.
- Folds: `R0845` and `R0847` add or subtract a 0x16-byte block
  into the modifier (`this = actor+0xfe`) and the live block (`actor+0xbe`).
  Their callers are Armor and Shield equip and remove (`R0848`,
  `R0851`, `R0849`, `R0852`), whose shipped blocks carry
  zero in the damage-kind bytes (`ITEM-ARMFOLD-033`; a block loaded from a SAV
  is Unknown), and `R0875`, reached
  from the fold `R0840` (modifier into live).
- Weapon equip `R0850` stores modifier `+0xe6`, `+0xf4`, `+0xf5`,
  `+0xf9..+0xfb`, `+0xfe` and the active byte `+0xb6`; neither it nor unequip
  `R0853` has a store in the block. No effect arm does
  (`ITEM-EFFARM-146`).
- Corpus: of 1,360 walked actor records in 90 save files (the two install
  save directories and 15 dated directories), 14 carry a nonzero modifier
  damage-kind byte, in 7 files. Three of those files (`game0018..0020`) in
  the install directory are byte-identical copies of dated-directory files, so
  the 14 records are 8 records in 4 distinct files, all from one dated series.
  All 14 are Humans with all six modifier damage-kind bytes
  and all six modifier protection words equal to 100 and live damage-kind
  bytes equal to 100, the output pattern of `R1562`. A further 9 Unit records carry nonzero live damage-kind bytes with a
  zero modifier.

**Confidence.** Medium. The direct-store sweep and every named routine are
exact, but the sweep cannot see bulk copies with a register-held size,
`REP MOVS`, or pointer arithmetic that never encodes the displacement. The
`memcpy` size sweep covers `L02478..L08481` only, and 35 of its 44 sites
print a register as the size operand, which the file does not resolve. The
address-form sweep (`sweep-addr.txt`) lists 80 pointer-forming
instructions, of which the actor-class ones are the constructors, the spawn
routine `R0184`, the derive `R0280`, the archive routine and
the Armor and Shield routines above.

**Unknown.** Which routine wrote the 14 saved records: the `#modify` order is
consistent with them and no other path found here produces that pattern, but
no save records its origin. What sets the placed-object dwords `+0x12c` and
`+0x134`. The writer of the 9 Unit records' live bytes (the Unit definition
load path was not read). Spell and potion writers beyond the Effect
dispatcher, taken from `MAGIC-ATTACH-016` and not re-read here.

### HERO-DKIDX-162

- `R0265` is the actor `vt+0x4c` strike resolver: `this` = target,
  argument 1 `A` = a 0x16-byte attack block, argument 2 = attacker. After the
  absorption clamp it reads `A+0x10` (`L09277` loads the byte at offset 0x10) and the
  target byte `[this+A+0x10+0xce]` (`L04147` loads the byte at that address), both
  zero-extended. A nonzero byte `v` gives `damage = ftol(damage * (100 - v) /
  100 + 0.75)` (`L09278`..`L09279`); zero skips the step.
- The melee strike `R0246` passes `A = attacker+0xa6` (`L04062`,
  `L02047`), so `A+0x10` is `attacker+0xb6`, the live attack block's active
  byte. The index is the attacker's active slot; the byte read is the
  target's.
- Active byte writers (sweep `store 0xb6`, 3 hits): Weapon equip
  `R0850` stores the `Weapons` attackType when it is below 10
  (`L04008`) and `0` when it is 10 or more (`L09280`); unequip
  `R0853` stores `0` (`L04199`). The attack-block fold
  `R0846` skips the byte (`HERO-MOD-016`).
- The `Weapons` attackType column, 27 rows per root, equal in EN and RU:
  1 for Dagger, the four swords and Plasma Sword; 2 for Axe and Two Handed
  Axe; 3 for BareHands, the clubs, Mace, Morning Star, Pick Hammer, War
  Hammer and both staves; 4 for Pike, Halberd and Lance; 5 for the bows,
  Crossbow, Sonic Beam and Boulder Thrower; 11 for Flame Thrower; -1 for the
  row named `rem`.

| active byte | target byte read | modifier byte |
|---|---|---|
| 0 (nothing wielded, or an elemental weapon) | `+0xce` | `+0x10e` |
| 1 blade | `+0xcf` | `+0x10f` |
| 2 axe | `+0xd0` | `+0x110` |
| 3 bludgeon | `+0xd1` | `+0x111` |
| 4 pike | `+0xd2` | `+0x112` |
| 5 shooting | `+0xd3` | `+0x113` |

- A value of 100 gives damage 0 after the `0.75` term truncates. Above 100
  the factor is negative (`HERO-FOLD-033` states the unclamped 8-bit adder);
  the resolver's own clamp at 0 is stated in `HERO-DAMAGE-022`.
- Live `+0xce..+0xd3` hold the fold of the modifier bytes (`HERO-FOLD-033`);
  only `R1562` and the archive set a modifier damage-kind byte
  (`HERO-MODDK-161`), so for a Human target these six bytes are zero unless
  one of those wrote them.

**Confidence.** High for the resolver's two reads and the melee caller's
argument, and for the active-byte writers. Unknown for the callers noted below.

**Unknown.** The other `vt+0x4c` call sites that can reach an actor (the
effect-damage path passes `effect+0x48`, `L05510`) were not read, so their
`A+0x10` is not established. Active byte 0 is a target of elemental
weapons (attackType 11); whether such an attack reaches this resolver was
not read. Active byte `0xff` (attackType -1) would index `target+0x1cd`;
that row's reachability is not shown.

## Open questions

- Whether a Building or Sack Position terrain key is used, replaced or
  discarded before the mandatory session-entry sender. EXP-0249 proves that the
  Sack arm emits cell/sub-cell full X/Y, while the Building arm reads both and
  emits only cell X/Y; neither reads Position `+0x08`. It also finds the earlier
  optional ordinary tick whose computed actors and events prevent an
  absolute-first static classification (`SAV-POSTLOAD-220`, `221`). The city
  arm has no loaded Building/Sack; retained-actor replacement is conditional on
  the entry boolean, saved Player latch and seating state. Actor entry consumes
  only Items admitted by its recipient/type/ownership gates and eligible
  Effects. Sack pickup moves quantity-one Items into an existing actor
  container, can deep-copy Effects during a stack split and then deletes the
  source container (`SAV-POSTLOAD-222`, `223`); only their absolute ordering
  against the pre-entry tick remains open. Discriminator: a safely witnessed
  cross-process load with the pre-entry tick controlled, followed by position
  use and a fresh save, separates raw-key survival from an earlier alias
  writer, repair or destructor.
- The members inside the 4374-byte session region past `SAV-SESS-031`.
- The contents of `Group+0x4c`'s and `Group+0x20`'s two `u16` lists, unopened
  by this repository (`Group+0x20`: see `SAV-1107`).
- Which live events reach the `Diary` dword writer. `SAV-842` locates it,
  `R1484`, and `SAV-843` places its one located call in attributed actor
  teardown; whether every kill reaches that path, whether its source is always
  the lethal one, and which event runs first after LOAD are open. `SAV-668`'s
  per-unit-type reading is withdrawn: the index follows the actor's
  class-dependent Units or Humans ordinal (`SAV-844`), and the meaning of each
  unobserved lower index is Unknown. The arrays' sizing, the corpus
  `word[i] = 1024 - dword[i]` relation (a finite corpus fact, `SAV-842`) and
  `Diary+0x2c` are closed by `SAV-667`–`SAV-669`.
- The 10 of 12 unidentified `+0x90` owning functions of `SAV-633`. The three
  Unit/order-relative `u16` lists' element identity, corpus shape and bounded
  runtime-accessor census are closed by `SAV-630`–`SAV-634`.
- For `SAV-636`'s `Token+0x18`/`+0x1c`: whether the operational recipient-mask
  meaning and clear-then-republish lifecycle of `+0x18` (`SAV-678`) generalize
  to every class, `+0x1c`'s meaning for every class but `Item` (`SAV-654`), and
  the five `+0x18` exceptions `SAV-655` could not explain. Serialization
  mechanics (`SAV-653`) and the corpus population of both fields (`SAV-655`,
  `SAV-656`) are closed.
- The instruction at `L08504`. The decode-anchored sweep of all 100 slot
  displacements (`SAV-650`, `SAV-658`) finds literal writes at exactly 90, 91,
  92 and 93, which closes the completeness question `SAV-651` left open. One
  raw match near `L08504`, inside a routine the sweep's function map anchors
  to `L08505`, could not be linearly decoded to a landing instruction in two
  attempts. The first byte of `L08505`, `0x9a`, is a far-call opcode outside
  this decoder's covered subset, which suggests the anchor may itself be a false
  function start rather than the decoder's only gap. What instruction is at
  `L08504` is Unknown.
- A save of a 256×256 map. The window `[0x807, 0x807+0xe5e7)` was only
  exercised by an 80×80 map, where nothing lies past its high end. On a
  256-wide plane rows 238..255 fall outside it entirely, so a save of
  `Islands`/`Tomb`/`scn:140` is the one measurement that tests the bound the
  other way; it costs one play session, not an experiment. It would also test
  `SAV-ID-015`'s ordinal law on a second map.
- Opcodes `0x00`, `0x7f`, `0x80`, `0x81` (`SAV-PACK-007`): a census finds 0
  occurrences over 12 516 positions, still only an absence. Settled by reading
  the codec in `rom.exe`.

## Entry cost argument source and mover stores

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1161 | The selected entry caller passes Position cell-X/Y low bytes through locals to the domain-1 cost read; its five explicit mover stores follow that call on every local normal-return CFG path. | High / Unknown | ● active | [EXP-0466](../experiments/EXP-0466-entry-cost-order/) |

### SAV-1161

- Caller `R0050` reads actor `+10` separately for cell-X/Y getters
  `R0299/R0300`. They load Position `+00/+01` as a byte. The caller caches
  these low bytes in frame `-08/-0c`, pushes Y, X and actor, and calls
  `R1345` at `L07074`, whose native return boundary is `L09281`.
- After saving two callee-saved registers, the callee's stack `+0c/+10/+14` holds actor/X/Y.
  On domain 1 it masks X/Y to eight bits and reads
  `world + (Y << 8) + X`. Domains 2/3 return the actor speed word; domain 0 or
  above 3 returns zero without that inline read.
- Explicit caller stores after normal return are mover word `+72` at
  `L07161`, then bytes `+82/+83/+84/+85` at
  `L02490/L02491/L02488/L09282`. The first two receive the caller-local
  cell bytes; the latter receive low full-coordinate getter bytes, Position
  `+04/+05`, sampled after the call. Each mover destination reloads actor `+154`.
- The complete caller CFG places the cost-call node before every path to these
  five explicit mover stores. Its other 11 explicit memory stores are caller
  frame locals. Implicit stack stores and callee effects are outside this count.

**Confidence.** High for these selected native instruction, receiver and
normal-return CFG facts: eight complete bodies and one incoming-call slice,
187 instructions per image, complete contiguous decode, raw relative targets,
local branch boundaries and a call-dominance check. EN/RU are one code image.
The 129 caller instructions match retained EXP-0292 instructions on each root.
No order or argument-value inference is taken from a synthetic callback cut.

**Unknown.** Concrete virtual callback targets/effects, actor/Position/mover
pointer stability, exceptional paths, native reachability and the cost byte's
writer/value history. Separate Position reads are not an atomic snapshot.
A fault before normal return precedes these explicit caller stores; it does not
exclude callback stores or any earlier writer.

**Evidence.** [EXP-0466](../experiments/EXP-0466-entry-cost-order/EXP-0466.md),
`evidence/instructions.tsv`, `evidence/anchors.tsv`, `evidence/edges.tsv`,
`evidence/ordering.tsv`, `evidence/caller-stores.tsv`, `evidence/callee-paths.tsv`,
`evidence/retained-comparison.tsv`, `evidence/spans.tsv`, `evidence/provenance.json`.

## Destroyed Building records

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1163 | A structure destroyed in play stays in the SAV as an ordinary 77-byte Building record with word `+0x42` equal to 0; in the 8 observed files no record is removed. | High / Unknown | ● active | [EXP-0480](../experiments/EXP-0480-structures/) |
| SAV-1164 | The destroyed Bee House's one footprint cell keeps its block rows and its cell-record Building key, equal to those of the intact structure in three files of the same signature. | Medium | ● active | [EXP-0480](../experiments/EXP-0480-structures/) |
| SAV-1165 | SAV marks destruction only by HP 0 with a positive maximum, which an authored-zero placement shares; in the 7,232 records and the one destroyed structure's record and cell no ruin flag or negative HP appears. | High / Unknown | ● active | [EXP-0480](../experiments/EXP-0480-structures/) |
| SAV-1166 | The introduction of a Building to a recipient reaches the same `0x82` arm as the HP notification, so a restored HP word is what that arm publishes as drawable health. | High / Unknown | ● active | [EXP-0480](../experiments/EXP-0480-structures/) |

### SAV-1163

- Population and join are in `TERR-STRUCT-212`. In the 199 files whose map resolves, the Building-family
  record count equals the authored type-4 count of that map (45 of 45 on all 15 files of mission 131).
  180 of them also match record for record on cell and kind; the other 19 are generated-document files.
- The Bee House of mission 131 (kind 24, cell (131,84)) is HP 0 with maximum 100 in 8 files
  (`oldsaves5` 0011 to 0017 and `oldsaves6` 0002, as byte-identical copies elsewhere) and in each of them its
  record is present with the same Token mask (`0002`), type word, `+0x44` (100), `+0x46`, `+0x48`, `+0x60`,
  `+0x61`, `+0x64` and `+0x68` as the same structure in three original files where HP is 100 (`oldsaves5` 0008, 0009 and 9999).
- Only word `+0x42` differs. `Building::Serialize` writes HP as an ordinary word (`SAV-BLDG-037`). That LOAD
  restores the word, and no separate destroyed state, is an inference from that serializer; no original LOAD
  followed by SAVE of an HP 0 record was made.
- An engine-lineage file written from `oldsaves5` 0017 (`exp-engine-lineage/game1017-engine-from-0017`,
  hash `1c7cb03aecac`) holds the Bee House at HP 100, and the original's resave of it (`5d398a6a8f80`) keeps
  HP 100 with rows `21/21`. These are engine-lineage files, not original evidence for LOAD. The Building identity
  keys repeat across some destroyed files and differ in others, so some may be load-and-resave products; that
  round trip was not tested.
- The 8 files are distinct hashes of one session lineage, not independent observations. Bound: the record persists in every file that shows destruction. A file written after a later removal
  would show a count below the authored count; none does in this population.

**Confidence.** High for the census over these files. Unknown: the producer of every file and whether
the original removes a structure at a later time that no file reaches (`UNIT-STRUCTBOUND-081`).

**Evidence.** [EXP-0480](../experiments/EXP-0480-structures/), `evidence/census/`,
`evidence/save-structure-join.tsv`, `evidence/nonpositive-hp.tsv`.

### SAV-1164

- The Bee House at (131,84) is 1x1. In each of the 8 destroyed files its cell record has Building key
  at payload `+0c` equal to the record's identity key, payload `+00,+01,+02` = `0b 00 00`, and block rows
  static/dynamic `25/25`.
- The three original files with HP 100 carry the same payload and rows. Four other files with HP 100
  (two under an engine-lineage experiment directory, `oldsaves5` 0020 and a crash-investigation file) carry
  rows `21/21`, payload `0b 05 00` and `+0x48` = 0; their producer chain is not established here.
- The Switch at (97,91), HP 0 in 6 files, has `25/25` and `08 00 00` as in its 5 intact files.
- Cells of damaged houses (HP 58 to 67 of 100) carry `35/35`; the destroyed house does not
  (`evidence/changed-in-play-footprint.tsv`). The meaning of that bit is not established.

**Confidence.** Medium: one 1x1 structure in one session; the block rows are the observed state at save
time, not evidence of a rule.

**Evidence.** [EXP-0480](../experiments/EXP-0480-structures/), `evidence/changed-in-play-footprint.tsv`,
`evidence/census/`.

### SAV-1165

- 149 records have HP not positive with a positive maximum: 134 authored-zero (`TERR-STRUCT-211`) and 15
  with authored word nonzero (`TERR-STRUCT-212`).
- All 149 are exactly 0; of 7,232 records none is negative. The 149 are 134 authored zeros plus one in-play
  event (15 records of three structures in one session lineage).
- The comparison covers one 77-byte record against its siblings and one cell record. No whole-document
  search for another destruction marker was run.
- Distinguishing authored zero from destruction requires the map's type-4 word, which SAV does not
  carry.

**Confidence.** High for the census. Unknown: a destruction marker elsewhere in the document, and the writer.
One event cannot tell the unclamped physical writer (`L09283`) from the clamped consumer (`L09284`); a
physical strike can store a negative word (`UNIT-STRUCTDELIVERY-065`) and no file here shows one.

**Evidence.** [EXP-0480](../experiments/EXP-0480-structures/), `evidence/nonpositive-hp.tsv`.

### SAV-1166

- `R0675` (callers `R0562`, `R1661`, `R0811`) calls `R0562` with tag 0 for an object whose
  virtual `+0x2c` is false, as the Building is (`UNIT-STRUCTDAMAGE-064`).
- `R0675` calls `R0562` only after `L09285` and `R0676` succeed.
- `R0562` then takes the Building arm: virtual `+0x34` true, and a nonzero object word `+0x0e`
  (`L09286..L09287`), which matches the zero-type-word case of `SAV-1093`. Message word `+0x12` is word
  `+0x42` when `+0x44` is nonzero, else 1000 (`L09288..L09289`). The HP notification (`R0561`, five
  call sites in four owners) reaches the same arm with tag `0x73` (`UNIT-STRUCTCONT-078`).
- Chain: saved HP word -> published health word -> client ruin gate (`TERR-STRUCT-210`).

**Confidence.** High for the local reads. Unknown: whether LOAD or mission start invokes an
introduction, the timing of that call, and the callers' own call trees.

**Evidence.** [EXP-0480](../experiments/EXP-0480-structures/), `evidence/static-facts.tsv`.

## Dead actor state word

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1169 | Every tick with signed health at or below 0 and `+0x54` not `0x10` stores 0 into `actor+0x50` (`L04557`); teardown stores `0x10` (`L00607`). No other store is found in the seven direct-call closures read. | High / Medium | ● active | [EXP-0490](../experiments/EXP-0490-dead-actor-state/) |
| SAV-1170 | Of 101 distinct-digest saves, dying actors store `+50` 0 (67 of 68 records) and ungrouped dead-list actors `0x10` (536 of 537); original-install files hold 21 of the 71 and 240 of the 652 raw records. | High / Medium | ● active | [EXP-0490](../experiments/EXP-0490-dead-actor-state/) |
| SAV-1171 | Unit LOAD reads `+0x50` as four raw bytes with no branch and the post-load rebind `R1482` and its direct-call closure hold no `0x50` operand; a grouped dead actor with `+0x54` not `0x10` is ticked and gets 0. | High / Medium | ● active | [EXP-0490](../experiments/EXP-0490-dead-actor-state/) |

### SAV-1169

- Per-tick routine `R0037`: `L05756`/`L09290` return before everything else when `+0x54` is `0x10`. Past
  that, the branch test `L04385 JG` skips the branch for health above 0. For signed health at or below 0,
  `L04557`, a store of zero at offset 0x50 of the object, and `L04558`, a store of zero at offset 0x54 of the object, run before the stage test at `L04559`
  (`HERO-DYINGTICK-145`). The branch holds three stores to `+0x50`, `+0x54`, `+0x58`: those two and
  `L01710 +0x54 = 0x10`. The routine has no direct caller; three vtables hold it (dwords at `L03030`,
  `L03031`, `L03032`, that is slot 20 of the table at `L09291` and slot 6 of the tables at `L00002` and
  `L00003`); which classes use which table is not read here.
- Teardown `R0208` stores `+0x54 = 0x10` at `L00609` and `+0x50 = 0x10` at `L00607` as its first two
  stores. A raw caller census (`callers.txt`: `E8`/`E9` rel32 sites plus dwords in every section) finds one site,
  `L01924`, reached when the dying countdown `+0x6c` is not positive and signed health, after the `+0x4a > 1`
  clamp to -1000, is at or below -10 (`HERO-DEATH-026`).
- Seven direct-call closures, each with every store-form operand of displacement `0x50` classified: teardown (439
  functions) holds one actor store, `L00607`; decay machine `R0867` (273), dead-list loop `R0860` (307),
  reap sweep `R0427` (262), post-load rebind `R1482` (10) and dead-manager walk `R1599` (6) hold none; cast
  interrupt `R0014` (76) and the dying-time getters hold none. The other hits are a packet counter, a stack slot,
  a constant-pointer initialiser, library same-value stores and virtual calls through slot `0x50`.
- Open set. `image-store-sweep.tsv` (linear sweep of `.text`, 260 store forms, owners by nearest function start)
  holds 185 store forms in the order and state setters `R0008..L09292`, 181 of them in owners that reference
  neither `+0x94` nor `+0x13c` (`setter-summary.txt`). Three store the literal `0xb` (`L09293`, `L00377`,
  `L09294`), beside the construction value at `L00324`. None is excluded for a still-grouped dying actor. A
  setter that reaches such an actor after its tick would leave the exception shape of `SAV-1170` (dying, `+54` = 0,
  `+50` = `0xb`). The sweep misses a store through a computed pointer and a block copy.
- Before death the word holds what the AI left. The strike `R0246` has no `+0x50` operand in its own body; its
  564-function closure holds 12 further actor-looking stores (`R0044`, `R0063` and stores of byte or of word width in
  `R0888..L04087`) reached through the effect applier `R0003`, whose arms are unread here. Two others in that
  closure are classified: `R0185` is a constructor body (callers `L00495`..`L00503`) and `R0160` a stat
  initialiser (caller `L06971`); neither acts on a live actor.
- Population of the pre-death word: 5,394 living records of 130 saves hold `0:269, 1:113, 3:11, 0xa:44, 0xb:4439,
  0xc:518`. In the owner's pair the victim holds 1 before and 0 after.
- Revive `R0672` calls `R0044` at `L09295`, which stores `0xc` at `L00119` under conditions not read here.

**Confidence.** High for the two writes, their conditions and the one-caller teardown (whole bodies, raw bytes, raw
caller census). Medium for exclusivity: the closures follow direct calls only (indirect calls are counted per root in
`closure-summary.tsv`) and the setter set above is not excluded.

**Unknown.** Whether an effect applier arm or an order setter writes a nonzero word to the victim in the sub-tick of
the kill; the 12 strike-closure stores' targets; the revive store's conditions; spell damage.

**Evidence.** [EXP-0490](../experiments/EXP-0490-dead-actor-state/), `evidence/dying-branch.txt`,
`evidence/closure-stores.tsv`, `evidence/closure-summary.tsv`, `evidence/callers.txt`,
`evidence/image-store-sweep.tsv`, `evidence/setter-summary.txt`, `evidence/bodies/`, `evidence/corpus-census.txt`.

### SAV-1170

- Population: 131 `.sav` files below the seat saves root; 130 parsed by the exact reader (one fails,
  `2026-08-27/EXP-0261-owner-runs/game9000.sav`); 101 distinct SHA-256 digests among the 130, so 29 parsed files are
  byte copies.
- Producer is classified from the path only (`corpus-census.txt`): `manifest-original` (a directory whose
  `MANIFEST.md` names the original install), `install-copy` (the `EXP-0278` copy of an install tree), `named-original`
  (path says "original", or the `2026-10-06` directory), `engine-named`, `unrecorded` (no producer record).
- Dying actors (stage above 0, health at or below 0, in a group, not in the dead list): `+50` = 0 and `+54` = 0 in 71
  records over 19 files, 67 records in distinct-digest files (Human 65, Unit 6 of the 71). Producer of the 71:
  `manifest-original` 16, `install-copy` 4, `named-original` 1, `unrecorded` 50. One further record (Human) holds
  `0xb` with `+54` = 0 in `game9998` of a crash-investigation directory, `unrecorded`. No record in the census has
  health at or below 0 with stage 0 outside the dead list.
- Dead-list actors in no group: `+50` = `+54` = `0x10` in 652 records at stages 1 to 5 (1, 1, 26, 383, 241), 536
  records in distinct-digest files. Producer of the 652: `manifest-original` 117, `install-copy` 116,
  `named-original` 7, `engine-named` 2, `unrecorded` 410. One record at stage 5 holds `0xb` with `+54` = 0
  (`oldsaves7`, `unrecorded`).
- Five dead-list actors that are also in a group hold `0xb` in `game1017-engine-from-0017`; `game0001-original-resave1017`
  holds 0 for all five with health 0 changed to -1 (`SAV-1171`).
- Owner pair, "boltcoming" to "dying" (112 records matched by identity key, 20 changed): the victim (obj 297 of
  "dying") changes `+50` 1 to 0, stage 0 to 1, health 156 to -11, `+6c` 0 to 7, and no other record changes `+50`.
  The one dead-list actor (stage 4, health -284 to -286) holds `0x10`/`0x10` in both.

**Confidence.** High for the record counts as occurrences in these files, and for the pair diff. Medium for the
original's output: producers are path-classified, most records sit in `unrecorded` files, successive saves of one
session repeat records, and the exceptions show non-original producers exist in the population.

**Unknown.** The producer of the three exception records; a save written in the sub-tick of an order to a dying
actor; any original-written dying record with a nonzero word (none among the 21 original-install records).

**Evidence.** [EXP-0490](../experiments/EXP-0490-dead-actor-state/), `evidence/pair-records.txt`,
`evidence/pair-diff.txt`, `evidence/corpus-census.txt`, `evidence/saves-manifest.tsv`.

### SAV-1171

- `Unit::Serialize` `R0210`: the store arm writes `+0x50`, `+0x54`, `+0x58` with the archive writer `R0271`
  and literal length 4 (`L09296..L09297`); the load arm reads them with `R0272` (`L09298..L09299`).
  Neither span holds a branch instruction.
- The post-load rebind `R1482` and its 10-function direct-call closure, and the dead-manager walk `R1599`
  and its 6, hold no operand of displacement `0x50` in any form (`any_operand_disp50` in `closure-summary.tsv`).
  Reads through an indirect call or a computed pointer are not covered.
- `Player::Serialize` LOAD (`L07331`..`L09300`) appends every `Group.actors` member to `[player+0x20]` with no
  filter. World LOAD (`L06795..L09301`) appends every actor of each `[player+0x20]` to the tick list at
  `[L00240]` unless `+0x4c` bit 3 is set; the loop tests nothing else. A grouped actor is therefore on the tick
  list unless that bit is set.
- The tick routine returns at once for `+0x54 = 0x10`. Otherwise a health at or below 0 rewrites `+0x50` to 0
  (`SAV-1169`). A grouped dead actor loaded with `+0x54` = `0x10` keeps its word; no such record exists in the
  population, so that case is Unknown.
- An actor in the dead list and in no group is never on the tick list; its word is not read by the tick, the decay
  machine or the session-entry sender (`R0059` has no `0x50` operand).
- Load witness: one pair, `game1017-engine-from-0017` to `game0001-original-resave1017` (lineage by file name): five
  grouped dead-list actors, `0xb` in, 0 out, health 0 to -1 (the decay machine acts on the dead list). The values 0
  and `0x10` are present in original-written saves (`SAV-1170`); that is an output, not a witnessed load. The owner
  reports that the original loaded "dying" normally; no committed file derives from it.

**Confidence.** High for the arm, the post-load closure, the player-list append and the tick-list loop. Medium for
"every value loads": one witnessed value (`0xb`, five records, one lineage); the rest rests on the absence of a branch.

**Unknown.** Native behaviour for other values; a reader of the word on a dead-list-only actor outside the routines
read here (no image-wide read census was run); the grouped dead-actor case with `+0x54` = `0x10`.

**Evidence.** [EXP-0490](../experiments/EXP-0490-dead-actor-state/), `evidence/load-arm.txt`,
`evidence/world-tick-rebuild.txt`, `evidence/closure-summary.tsv`, `evidence/pair-records.txt`.

## Enemy card knowledge carrier

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1172 | The only serialized input found to the monster table nibble is the Diary in the Player tail (119 dword counts, 119 word remainders); no SAV record naming the client table at `R+0x3f58` was found, so the original rebuilds it from opcode 186. | Medium | ● active (amended) | [EXP-0491](../experiments/EXP-0491-enemy-card-knowledge/) |
| SAV-1173 | Mission 151 owner saves carry all-zero Diaries, so the level 5 to 7 cards the owner reported on enemy Humans did not come from a Diary count; the source is Unknown. | High / Unknown | ● active | [EXP-0491](../experiments/EXP-0491-enemy-card-knowledge/) |

### SAV-1172

- Carrier: the Player record tail holds the Diary as a dword array of kill counts per Units/Humans row (119 elements) and a word array of remainders (1024 minus count), with a trailing reference restored at LOAD (`SAV-847`). The count is capped at 17, seen in two saves with byte-identical Diaries (`game0002-bigsack`, `game0017-victory`); the remainder relation is `SAV-668`'s.
- Not carried, as far as searched: the client table at `R+0x3f58` and `+0x3f5c`. Searched population: `tools/savdoc`, `formats/sav` and `formats/rom2-sav`; the only hit is the client-array statement of `SAV-995`. The frontend object is constructed empty (`R0391`); the image's serializers were not enumerated. After LOAD the original therefore holds no monster level until opcode 186 arrives.
- The Diary-to-table path is opcode 186 (`L08064` to `L08062`); the first send after LOAD is gated on a change of `count>>1` and was not read to its trigger.

**Confidence.** High that the Diary is saved. Medium that the table is not: a bounded decoder search, not an image-wide serializer census. Medium for the post-LOAD statement: no run was made.

**Amended.** `Player::Serialize` (`R0415`) writes no `Player+0x68`, and a Player rebuilt at LOAD holds the constructor value 0 by that route (`MENU-103`, Medium: the other archive-side writers were sampled); the byte that selects the all-15 send is then not restored by LOAD.

**Unknown.** The first refresh and delivery timing after LOAD.

### SAV-1173

- Population: four owner saves of mission 151 (`game0002-original`, `game0003-original-native151`, `game2320-input`, `game9999-original-entry151`). `game0002-original` and `game2320-input` hold 115 Diaries each, the other two 5; 20 are Player-class and 220 Humanoid-class. All 240 are zero (`evidence/diary-census-m151.tsv`).
- The owner's cards of enemy Humans of types 9 and 24 showed levels 5 to 7. A Human reads no table nibble (`UNIT-146`), so the source was not separated among the three level inputs read in `R0877` (owner index 0, row bit 8 for the enemy entry, global `L01661` set by `#show map`) and an untraced writer (`SAV-1126`). A zero Diary excludes no Human input.
- Owner-reported, not reproduced by this experiment: in the original `#show map` also makes every card show its values. This does not settle which level input it uses and raises no grade.
- `SAV-1125` already fixes that the numbers of such a Human come from packets; this claim concerns only the level that gates their display.

**Confidence.** High that the Diary counts are zero in these saves; the observation is owner-reported with no stated source; Unknown for which input applied.

**Unknown.** Which of the three sources applied; no runtime observation was made.

## Ranged removal archive frontier

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1180 | Ranged Weapon removal writes members of Unit's raw attack and Modifier spans; their following actual SAVE vector remains Unknown. | High / Unknown | ✔ promoted | [EXP-0496](../experiments/EXP-0496-ranged-unequip-modifier/) |

### SAV-1180

Unit Serialize `R0210` supplies actor `+a6` to `R1594` with literal
length 24 and actor `+d4` to `R1570` with literal length 64. The storing
arms pass the current buffer addresses to raw archive Write. Human's
ordinary serializer reaches this Unit programme through Humanoid.

ITEM-161's selector store `+b6` is live-attack member 16. Its word `+e6`
is Modifier members 18..19, and bytes `+f9/+fa/+fb` are members 37/38/39.
Physical-kind 5 stores `+f4/+f5` are members 32/33, and `+fe` is members 42..43.
These writes are part of the raw serialized buffers; the archive helper does
not reconstruct them from the equipment references. Four isolated native
helper controls with distinct nonzero synthetic payloads preserve the
existing source address, payload and lengths at Write/Read call boundaries.
IsStoring and actual archive transfer are explicit cuts; no file is written.

The promoted constructor controls SAV-HUMGAPS-449 and SAV-HUMNEW-505 identify
separate live `+a6`, Base `+114` and Modifier `+d4` buffers. The shared attack
initializer writes 22 of 24 bytes; Modifier uses its separate memset 64. They
establish neither a safe default nor preservation until this archive.

The intervening frontier includes equipment Effects, actor derive, a
quotient-changing load refresh's additional derive, optional Spell deletion,
command22 container/load/notification events, elapsed runtime and earlier
nested serialization. Selected Unit derive/load and base-Effect paths run
natively in the controls. Human derive and Spell deletion are explicit cuts;
their complete transitive effects are unresolved. A synthetic callback model
can change the modifier again while satisfying every local removal byte.

**Confidence.** High for local member joins and raw-span call operands.
Complete PE-mapped instruction listings and four source-transfer controls
exclude a locally synthesized/default archive vector at those helper
boundaries. They do not execute archive buffering or a native SAVE.

**Unknown.** The first-archive vector after ordinary ranged removal, full
callback/lifetime closure, actual buffer/file correspondence and runtime
occurrence. No first-archive value is supplied.

## Unit-shot origin, projectile counter, shooter death and trail

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1188 | The unit shot starts at the shooter's current x/y plus 8 × (ShootOffset pair − Center), the pair chosen by ((dir − 8) & 14) / 2 from eight XY pairs; the cast producer reads the same class data from the cached centre. | High / Medium | ✔ promoted (amended) | [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/); [EXP-0508](../experiments/EXP-0508-shot-remainder/) |
| SAV-1189 | Eight instructions store the projectile counter at client world `+0xa0c`: zero in the world constructor, old + 1 at six record insertions and the low u16 of FreeIndex at LOAD; no other writer was found in the decoded population. | High / Medium | ✔ promoted (amended) | [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/); [EXP-0508](../experiments/EXP-0508-shot-remainder/) |
| SAV-1190 | No found writer resets the projectile counter at a mission boundary: its one zero store is in the client-world constructor, reached only from virtual method `R0326`, and the collect-all clears do not write it. | Medium | ✔ promoted (amended) | [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/); [EXP-0508](../experiments/EXP-0508-shot-remainder/) |
| SAV-1191 | LOAD sets the projectile counter to the low u16 of `Projectiles/FreeIndex`, 0 when absent, before it inserts the saved records under their document IDs; SAVE writes the same counter. | High | ✔ promoted | [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/) |
| SAV-1192 | A shooter's pending attack damage, ranged included, is never applied once an actor tick sees its HP at or below 0: that tick overwrites the countdown and both damage evaluators refuse an attacker at HP 0 or less. | High | ✔ promoted | [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/) |
| SAV-1193 | SAVE writes no projectile trail point and LOAD builds each record with an empty trail; the driver's trail arm then refills it for pictures 10 and 12, and other trail writers were not searched. | High | ✔ promoted | [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/) |
| SAV-1194 | The leave-map routine `R0121` takes the actor off the list the walker `R0427` ticks, without writing its countdown or state; other callers of the tick slot were not searched. | High / Medium | ● active | [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/) |

### SAV-1188

- Unit shot `R0603`, window `L13450..L13451`, with the shooter in edi
  and the record in esi: the class is `[L02113][shooter+0x20]`; i =
  `(shooter dir +0x6c − 8) & 0xe` is a dword index into the ShootOffset data
  `[class+0xec]`.
- `x = shooter+0x08 + 8 × (ShootOffset[i] − class+0x34)`,
  `y = shooter+0x0c + 8 × (ShootOffset[i+1] − class+0x38)`; both are copied to
  `+0x28/+0x2c`. The window has no test of the array count `+0xf0` and no
  other fallback.
- Loader keys: `CenterX` (`L13452`) goes to class `+0x34` (`L13453`),
  `CenterY` (`L13454`) to `+0x38` (`L13455`), `ShootOffset` (`L13456`)
  to the CArray at `+0xe8` (`L13457`), data `+0xec`, count `+0xf0`; an empty
  array takes the parent class's array (`L13458..L13459`).
- Pair k = i / 2 serves facings 8+2k and 9+2k (mod 16): S/SW/W/NW/N/NE/E/SE
  use pairs 0..7. EN `units/units.reg` (34 classes): Human Archer 14, Human
  CrossBowMan 15, Catapult 1 26 and Catapult 2 27 each hold 16 dwords. The
  original window, run on synthetic memory for those four classes and all 16
  facings (64 runs), gives the formula's x/y and the same `+0x28/+0x2c`.
- The cast producer `R0620` uses the same i, array and Center but adds the
  caster's cached `+0x58/+0x5c` and has the empty-array and picture-60
  fallback (`MAGIC-261`). The two routines differ in base and fallback.

**Confidence.** High for the instructions, the key mapping and the 64 runs.
Medium for 8 fine units per map pixel, the scale `SAV-1130` already grades.

**Unknown.** The first drawn position after the shot is built; a customised
registry with fewer than 16 dwords (the window reads without a bound).

**Amended.** `SAV-1196` narrows the inheritance clause: an empty array
re-reads the key from the parent's own section, one level only, rather than
taking the parent's resolved array. `SAV-1196` gives the start window with
no array.

**Evidence.** [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/),
`evidence/disasm-direction-start.txt`, `evidence/class-key-names.txt`,
`evidence/direction/unit-shot-origin.tsv`, `evidence/direction/registry-input.tsv`.

### SAV-1189

| Instruction | Routine | Value |
|---|---|---|
| `L13460` | client-world constructor `R0391` | 0 |
| `L13461` | client arm 0x86 | old + 1 |
| `L13462` | client arm 0x8b | old + 1 |
| `L13463` | client arm 0x8c | old + 1 |
| `L13464` | unit shot `R0603` | old + 1 |
| `L13465` | cast producer `R0620`, as `[world+0x9d4]+0x38` | old + 1 |
| `L13466` | cast producer, second record | old + 1 |
| `L13467` | mission loader, SAV document arm | low u16 of FreeIndex |

- Each store is 16 bits, so the increments wrap at 65536. Each insertion
  takes the old value as the record ID (`SAV-1131`).
- Population: a linear `.text` sweep finds 17 instructions with displacement
  `+0xa0c` (these eight stores, nine loads); a raw dword scan of every section
  finds 19, the two extra inside instructions at `L13468` and `L13469`
  whose bytes contain it. No decoded memory operand starting at
  `+0x9fd..+0xa0d` other than `+0xa0c` overlaps the word. The store base
  `+0x9d4` has 23 decoded touches and 23 raw hits; a non-stack 16-bit MOV to
  `+0x38` occurs once in `.text`, at `L13465`.

**Confidence.** High for each listed store and value. Medium that the list is
complete: a store through a pointer computed from another field, a block copy
or a width other than a 16-bit MOV at `+0x38` would not be seen.

**Unknown.** A writer outside the decoded population. A native watchpoint on
the two bytes would settle it.

**Amended.** `SAV-1200` widens the search population with derived-base,
callee and bulk-write searches; the writer list is unchanged and its
completeness stays Medium.

**Evidence.** [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/),
`evidence/identity/counter-census.txt`, `evidence/identity/counter-ranges.txt`.

### SAV-1190

- The constructor `R0391` has one call site, `L06468`, in `R0315`,
  which stores the new world at campaign `+0xd0`. `R0315` has one call site,
  `L13470`, in `R0326`, a virtual method referenced only by the `.rdata`
  dword `L01309`. The committed listing holds its prologue only.
  Neither constructor address occurs as a raw dword.
- The mission loader `R0099` reads the existing world from campaign `+0xd0`
  (`L13471`). Its ordinary entry sends message 4 (`L03843`, `R0545`).
  The FreeIndex read is reached only when campaign `+0x6bc` is 2 and the
  loader's document argument is non-zero (`L13472..L13473`), the arm that
  reads the saved `CurrentState`, `GameOptions` and `Projectiles` keys; no
  branch of the loader enters `L13474..L13475` from outside it.
- The collect-all routines `R0809` (706 instructions) and `R1301` (652)
  destroy and unlink every record. Neither has a `+0xa0c` operand, a non-stack
  16-bit store to `+0x38` or a MOVS/STOS instruction.
- So, within the found writers, the counter keeps counting across collect-all
  clears and mission entries of one process; a SAV LOAD replaces it
  (`SAV-1191`).

**Confidence.** Medium. The call chain and bodies are direct, but the claim
rests on `SAV-1189`'s writer list, and the callees of the two clear routines
and a whole-object copy of the world were not excluded.

**Unknown.** Whether `R0326` runs more than once per process. A watchpoint
across a mission change would settle the reset question.

**Amended.** `SAV-1201` identifies `R0326` as the application object's
InitInstance slot and reads the mission loader's three callers; the grade
stays Medium.

**Evidence.** [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/),
`evidence/identity/counter-census.txt`, `evidence/identity/counter-ranges.txt`,
`evidence/identity/counter-reset-census.txt`, `evidence/vtables.txt`.

### SAV-1191

- LOAD (`L13475..L13467`): `Projectiles/FreeIndex` is read with default 0
  (`L13476`) and its low 16 bits are stored at world `+0xa0c` (`L13467`)
  before the `IDs` vector is read and the records are built
  (`L13477..L13478`). Each record takes its ID from the vector
  (`L13479`, `L13480`); the counter is not read or derived from the IDs.
- SAVE (`L13481..L13482`) writes the 16-bit counter, zero-extended, as
  `Projectiles/FreeIndex`. A save and a load therefore return the counter
  unchanged, and the next live shot takes ID FreeIndex.
- The owner's EN mission-150 saves game0022, game0023 and game0024 hold
  FreeIndex 28, 29 and 31 with records Prj27, Prj28 and Prj29/Prj30.

**Confidence.** High: direct listings of both arms.

**Unknown.** A native LOAD followed by a new shot was not run.

**Evidence.** [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/),
`evidence/disasm-save-load.txt`, `evidence/save-observations.txt`.

### SAV-1192

- Actor tick `R0037` (vtable slot `+0x18` of `L00001`, `L00002` and
  `L00003`) tests the signed HP word `+0x94` (`R0864`). At or below 0 it
  sets `+0x50` and `+0x54` to 0. On the first such tick (`+0x13c` was 0) it sets
  `+0x13c`, halves `+0xbe`, stores the dying time from virtual slot `+0x6c`
  minus 1 into the countdown byte `+0x6c` (`L04377`) and returns. Later ticks
  decrement that byte and return; at its end an HP at or below −10 sets state
  `0x10` (`L13483..L01924`).
- The pending attack is the countdown `+0x6c` with sub-state `+0x58` 5; only
  its arm `L13484` calls the resolver `R0001` at zero (`SAV-1132`). A dead
  tick returns before the dispatch that reaches that arm.
- Both evaluators also refuse a dead attacker: the unit evaluator `R0246`
  jumps to its exit when attacker HP is not above 0 (`L13485..L13486`),
  before the Fire_Ball rider call `L03045` (`ANIM-115`); the Building
  evaluator `R0563` does the same (`L13487..L13488`).
- The active-list walker `R0427`, called on `[L00240]` (`L13489`),
  skips the tick for state `0x10` and moves the actor to the dead list.

**Confidence.** High for this control flow: complete listings of the tick's
death branch, both evaluators and the walker. The visual record is not
touched by this path; its survival is `SAV-1133`'s Medium.

**Unknown.** An HP that rises above 0 before any tick sees it at or below 0.

**Evidence.** [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/),
`evidence/disasm-shooter-death.txt`, `evidence/vtables.txt`.

### SAV-1193

- The writer's record loop `L13490..L13491` writes the sixteen leaves and
  reads no offset at or above `+0xa8` of the record, so not the trail at
  `+0x138..+0x140`.
- The loader builds each record with the constructor `R0609`, whose CArray
  constructor sets data and count to 0 (`ANIM-140`). The leaf arm
  `L13477..L13478` and the post-load helper `R0614..L13492` write no
  trail field; the helper's two virtual calls are `+0x20` and `+0x24`, which for
  the projectile vtable `L02587` return 1 (`L03099`, `L03100`).
- The driver's trail arm appends for pictures 10 and 12 only (`ANIM-140`);
  whether another writer gives a trail to other pictures was not searched. The four owner records have 16 numeric leaves
  each and pictures 1, 5, 2 and 13.

**Confidence.** High for the omission and the empty trail after these direct
arms. Which pictures can hold a trail before SAVE carries `ANIM-140`'s bound.

**Unknown.** The first drawn frame after LOAD. A native SAVE and LOAD of a
picture-10 or 12 shot in flight would show the trail loss and regrowth.

**Evidence.** [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/),
`evidence/disasm-save-load.txt`, `evidence/disasm-smoke.txt`, `evidence/vtables.txt`,
`evidence/save-observations.txt`.

### SAV-1194

- `R0121..R0122`: when `+0x4c` bit 3 is clear it releases the actor's
  occupancy (`L07141`), sets bit 3, removes the actor from the list at
  `[L00240]+4` (`L12780`) and sends message 0x74 twice. It writes neither
  `+0x6c` nor `+0x54` nor `+0x58`.
- That list is the one the walker `R0427` ticks (`SAV-1192`), so that
  walker does not advance the countdown while the actor is off it. Other
  callers of tick slot `+0x18` were not searched.
- The two return routines `R0122..R0123` and `R0123..R0671`
  re-insert the actor through `R0032` on `[L00240]` (`L13493`,
  `L13494`) and clear bit 3 of `+0x4c`. Their own instructions write no
  other actor field; the callees `R1145`, `R0032`, `R0811`,
  `R0059` and `R0623` were not read.

**Confidence.** High for the leave-map and return routines' own instructions.
Medium that the countdown stops off the map: it rests on `R0427` being
the only tick caller. Medium for what a full leave and return does to the
pending hit: the routines' callers and callees were not read.

**Unknown.** Another caller of tick slot `+0x18`; a census of indirect calls
through `+0x18` on actor pointers or a native watch on `+0x6c` would settle
it. The pending hit after a leave-and-return; what actor destruction does to
a pending hit (no destructor census).

**Evidence.** [EXP-0505](../experiments/EXP-0505-unit-shot-leaves/),
`evidence/disasm-shooter-death.txt`.

## Unit-shot admission, empty shoot offsets and the projectile counter

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1196 | With no ShootOffset dwords after the one-level parent re-read, the data pointer is 0 and `L13646` reads address 0x00..0x38, in the Win32 null region, after the record is built and before x is stored. | High / Unknown | ✔ promoted | [EXP-0508](../experiments/EXP-0508-shot-remainder/) |
| SAV-1197 | A ranged attack whose target word names a structure builds an ordinary unit-shot record: the structure is in the same client hash, `R0603` measures segments from its `+08/+0c` and the driver homes on its `+0x58/+0x5c`. | High / Medium | ✔ promoted | [EXP-0508](../experiments/EXP-0508-shot-remainder/) |
| SAV-1198 | A shooter whose class Projectile is 13 or above is admitted like any other: the release test checks only Projectile non-zero and `R0603` copies the value into the record's picture with no bound. | High | ✔ promoted | [EXP-0508](../experiments/EXP-0508-shot-remainder/) |
| SAV-1199 | Client arms 0x8b and 0x8c build records directly, not through `R0603`; the simulation sends 0x8b from the spell senders `R0617`/`R0618` when the source has no client ID and 0x8c from `R0619`, in actor states 0xd and 0xe. | High | ✔ promoted | [EXP-0508](../experiments/EXP-0508-shot-remainder/) |
| SAV-1200 | Four searches over a 554,699-instruction linear `.text` decode find the same eight writers of the projectile counter as `SAV-1189`; 1,909 bulk-write sites with an unresolved destination or length remain unsearched. | High / Medium | ✔ promoted | [EXP-0508](../experiments/EXP-0508-shot-remainder/) |
| SAV-1201 | The client world is built once, from the application object's InitInstance slot; the mission loader's two non-LOAD callers pass flag 0, skipping the FreeIndex arm, so no found writer resets the counter at a mission start or end. | Medium | ✔ promoted | [EXP-0508](../experiments/EXP-0508-shot-remainder/) |

### SAV-1196

- The class constructor `L13647` builds the ShootOffset CArray at `+0xe8`
  through `L05586`, which sets data `+0xec` and count `+0xf0` to 0.
- The array reader `R1061` returns 0 without touching the array when the
  key is absent. A present key with no values calls SetSize(0) (`R1198`),
  which frees the data and stores 0.
- Inheritance `L13457..L13459`: when the own count is 0 and the parent
  index is not −1, the loader re-reads the `ShootOffset` key from the parent's
  own `Unit%d` section, whose path `L13648..L13649` builds from the
  parent's section index. It does not copy the parent's resolved array, so a
  class two levels below the class that holds the key gets none. Projectile
  `+0xd4` instead defaults to the parent's resolved value
  (`L13650..L13651`).
- Unit shot `R0603`: the record is allocated and constructed
  (`L13652..L03005`) before the window `L13450..L13451`. At
  `L13258` it loads data `+0xec`, and `L13646` reads `[data + 4i]` with
  i = (dir − 8) & 14.
- The original window, run in Unicorn with page 0 unmapped and data 0 for all
  16 facings, stops at `L13646` on the emulator's unmapped-read hook,
  reading 0x00..0x38, and stores no x. The run adds no discrimination beyond
  the listing.
- That the original faults rests on a platform premise: the lowest pages of a
  Win32 process are never mapped, so a read of 0x00..0x38 raises an access
  violation. The Unicorn run models that premise by leaving page 0 unmapped.
- EN and RU `units/units.reg` (34 sections, same SHA-256): every class with a
  resolved Projectile (14, 15, 24, 26, 27, 65, 70, 71, 79) holds its own 16
  ShootOffset dwords, so the case needs a customised registry.

**Confidence.** High for the loader, the inheritance rule and the faulting
read, given the Win32 null-region premise. Unknown what follows the access
violation.

**Unknown.** Whether a structured exception handler catches the fault and
what the record and the shooter then hold. A run of the original with such a
registry would settle it. A non-empty array shorter than 16 dwords stays
`SAV-1188`'s Unknown.

**Evidence.** [EXP-0508](../experiments/EXP-0508-shot-remainder/),
`evidence/q1-listings.txt`, `evidence/q1-null-array-run.tsv`,
`evidence/q1-units-registry.tsv`.

### SAV-1197

- Client arm 0x82 (`L04899`) inserts a structure object (constructors
  `L03225`, `R1842` over the drawable base `R1652`) into the client
  hash at world `+0x9b8`, keyed by message `+0xa`. The simulation sender
  `L13653..L13654` writes the building's `+4` there.
- The 0x72 builder `R1064` writes the actor's target `+0x5c` word `+4` at
  message `+0xe` with no type test; `R0549` chooses 0x72 when the reach is
  above 1.
- Client arm 0x72 (`L09206`) looks up the shooter in `+0x9b8`; a missing
  shooter logs "Invalid unit #%d. Command Ranged Attack." and builds nothing.
  It sets the shooter's action to 7 (`L05332`) and copies message `+0xe` to
  its target word `+0x86` (`L13655`).
- The unit shot `R0603` looks up `+0x86` in the same hash
  (`L13656..L13657`). A structure is found like a unit; a missing ID
  builds nothing (`L13658`). Segments are the integer distance from the
  shooter's `+08/+0c` to the found object's `+08/+0c` divided by 200
  (`L13659..L13660`); the picture and the other leaves are those of
  `SAV-1130`.
- The driver homes on the found object's `+0x58/+0x5c` (`ANIM-139`), which
  for a structure the post-load helper `R0614` derives as for any
  drawable.

**Confidence.** High for the client path: complete listings. Medium that the
simulation sends 0x72 with a structure target during play: the builder takes
any target word, but the actor states that set a structure as the target were
not read.

**Unknown.** Whether a structure's `+08/+0c` and `+0x58/+0x5c` are the same
anchor as a unit's; a run against a structure would show the drawn path.

**Evidence.** [EXP-0508](../experiments/EXP-0508-shot-remainder/),
`evidence/q2-listings.txt`.

### SAV-1198

- The action-7 release test `L02711..L02716` calls the unit's slot
  `+0x58` when the phase `+0x94` equals the class's `+0xfc` and
  the class Projectile `+0xd4` is non-zero. No other test applies. Slot
  `+0x58` of the unit vtables `L02468` and `L02585` is the unit shot
  `R0603`.
- `L13661/L13662` copies Projectile into the record's picture `+0x20`
  with no bound. The record takes the `SAV-1130` leaves and the
  `SAV-1196`/`SAV-1188` start point.
- After that, the driver arm for the picture (`ANIM-143`) and the draw arm
  (`ANIM-145`) govern the record; a picture with no defined projectile row is
  refused by the draw (`ANIM-PROJ-026`).
- No shipped class has a resolved Projectile of 13 or above (EN and RU
  `units/units.reg`: 1, 2, 3, 4, 5, 6, 7, 10, 12).

**Confidence.** High: complete listings of the test and the copy.

**Evidence.** [EXP-0508](../experiments/EXP-0508-shot-remainder/),
`evidence/q1-listings.txt`, `evidence/q1-units-registry.tsv`,
`evidence/q2-listings.txt`.

### SAV-1199

- Simulation side: `R0268` is called only at `L01994` (actor state 0xd,
  a target) and `L01995` (state 0xe, a cell) in the actor tick. State 3 with
  a weapon-spell item (`L13663..L13664`) switches to 0xd instead of the
  swing `R0245`, which sends 0x72.
- The senders `R0617` and `R0618` build message 0x86 and rewrite the
  opcode to 0x8b when the source's word `+0xe` is 0 (`L05307..L03021`),
  the source having no client ID (`MAGIC-DELIVER-035`). The picture byte
  `+0xc` is 2 × spell + 8 (`L05310`). With a target argument the message
  carries the target's `+4` at `+0xd` and a word at `+0xf`. `R0619` sends
  0x8c.
- With a client ID the message stays 0x86 with an even picture. Client arm
  0x86 builds a record only for an odd picture (`L05315..L05316`); the
  even path `L03013..L05320` looks up the source unit, sets its action
  `+0x84` to 8 and its target fields, so the cast producer builds the record
  (`ANIM-144`).
- Client arm 0x8b (`L05320..L13638`) constructs the record with
  `R0609`: picture = message `+0xc`; x/y = cell × 256 + 128; the target is
  message `+0xd` when the registry slot's `+0x2c` is non-zero, otherwise the
  aim cell; segments = message `+0xf`; actionphase −1; action 1; ID = counter,
  then counter + 1, inserted through `L13665`.
- Client arm 0x8c (`L05321..L13639`) builds a picture-36 record the same
  way.

**Confidence.** High: complete listings of the two call sites, the senders
and the three client arms.

**Unknown.** Which shipped sources have no client ID in play.

**Evidence.** [EXP-0508](../experiments/EXP-0508-shot-remainder/),
`evidence/q2-listings.txt`.

### SAV-1200

| Search | Population | Found |
|---|---|---|
| 1 | memory operands of any base whose bytes overlap `+0xa0c..+0xa0d` | 17 operands, 7 writes |
| 2 | `reg := base + K`, K 0x10..0xa0d, then `[reg + d]` overlapping the word within 60 instructions | 10,868 derivation sites, 1 hit: `L13465` (K 0x9d4, d 0x38) |
| 3 | direct callees reached with ecx = base + K from search 2, intra-procedural bodies | 1,537 pairs, 0 hits |
| 4 | REP STOS/MOVS and calls of memset `L02418` or memmove `R0544`, destination and length resolved within 25 instructions | 2,111 sites: 195 resolved and not covering, 7 covering, 1,909 unresolved |

- The eight writers are `SAV-1189`'s: `L13460`, `L13461`, `L13462`,
  `L13463`, `L13464`, `L13465`, `L13466`, `L13467`.
- None of the 7 covering bulk writes targets the client world (0x3f78 bytes):
  `L13666` clears 614,400 bytes, larger than the world; `L13667` is in
  `L13668` (four direct callers `L13669`, `L13670`, `L13671`,
  `L13672`), a display-mode object: it releases an interface held at
  `+0x838`, copies one 0x13c-byte entry of a table at `+0x9a4`, and its failure
  string is "Could not fit the rendering surfaces in video memory for this
  hardware device."; `L13673` is in the AI
  session constructor; `L13674` clears a 0x271-dword table; `L13675` writes
  `[esi+0xa50]+0x92ef4` (from `L13676`); `L13677` has no reference;
  `L13678` writes the stack.
- Population: Capstone linear decode of `.text` with skipdata; EN and RU
  `rom.exe` have the same SHA-256.

**Confidence.** High for each search over its population. Medium that the
list is complete: the 1,909 unresolved bulk sites, indirect callees and
stores through a pointer read from memory are not excluded.

**Unknown.** A writer in the unresolved population. A native watchpoint on
the two bytes across a mission change would settle it.

**Evidence.** [EXP-0508](../experiments/EXP-0508-shot-remainder/),
`evidence/q3-counter-census.txt`, `evidence/q3-listings.txt`.

### SAV-1201

- The world constructor `R0391` has one call site, `L06468` in
  `R0315`, whose one call site is `L13470` in `R0326`. `R0326` is
  slot `+0x58` of vtable `L01179`, which the constructor `L01312` installs
  in the global application object `L13679`. Slot `+0x58` is the MFC
  InitInstance slot: the startup routine calls `+0x8c`, then `+0x58`, then
  `+0x5c` or `+0x70` (`L13680..L13681`).
- The mission loader `R0099` reaches the FreeIndex arm only when campaign
  `+0x6bc` is 2 and its flag argument is non-zero (`L13682`,
  `L13472..L13473`). It has three callers:
  - `L03842` and `L07366` push edi, which is 0 from `L11306`; the other
    edi writes in `R0701..L13683` do not reach them by direct flow,
    indirect jumps not followed;
  - `L06711` pushes 1 in `R1284`, which sets `+0x6bc` to 2 and is called
    at `L13684` after the collect-all `R1301` (`L13685`). This is the
    LOAD path (`SAV-1191`).
- `L13686..L03821` copies "game9999.sav" to campaign `+0x234` and calls
  the SAVE routine `R0084`. Its intra-procedural body (565 instructions)
  has one counter operand, the read `L13687`.
- So, within the `SAV-1200` writers: the counter is 0 once per process, counts
  across mission entries and collect-all clears, and is replaced only by a
  LOAD.

**Confidence.** Medium. The call chain is direct, but that InitInstance runs
once rests on the MFC startup shape, and the reset conclusion rests on
`SAV-1200`'s Medium completeness.

**Unknown.** Another indirect call of slot `+0x58` on the application object;
a native watchpoint across a mission change would settle both.

**Evidence.** [EXP-0508](../experiments/EXP-0508-shot-remainder/),
`evidence/q3-listings.txt`, `evidence/q3-counter-census.txt`.

## Saved form of records in flight

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1202 | SAVE writes the 16 Prj leaves of every record with no filter; LOAD reads the 16, rebinds `+0x14`, `+0xe0`, the ID and the map node, and loses the word list `+0xa8`, the point array, the trail and the cached x/y `+0x28/+0x2c`. | High | ✔ promoted | [EXP-0509](../experiments/EXP-0509-flight-records/) |
| SAV-1203 | A loaded record's derived cell is (0,0) until `R0614` next runs after a driver call, so a loaded picture-13 record at saved actionphase 7 marks the 3 × 3 around map cell (0,0), four of them words before the tile plane. | Medium | ✔ promoted | [EXP-0509](../experiments/EXP-0509-flight-records/) |
| SAV-1204 | A cast record with 0 segments takes a counter ID and dies in the 0x401 pass that built it, so no SAV holds it; the 133-file corpus holds 8 distinct records, pictures 1, 2, 5, 10 and 13. | High / Medium | ✔ promoted | [EXP-0509](../experiments/EXP-0509-flight-records/) |

### SAV-1202

- SAVE Prj loop `L13430..L08366` walks the record map `+0x9d4` and writes
  x, y, z, picture, dir, phase, lastaction, action, actiondir, actiontarget,
  actionx, actiony, actionz, actionphase, actionsegments and actionspell for
  each record; no picture, segment or action test skips one.
- LOAD loop `L13477..L13432` constructs a record with `R0609`, reads
  the 16 leaves into `+08, +0c, +10, +20, +6c, +70, +74, +84, +85, +86, +88,
  +8c, +90, +94, +a0, +a4` with the constructed values as defaults, sets
  `+0xe0` (view) and `+0x14` (view `+0x9b4`), takes the ID from IDs, links
  the node and calls `R0614` (`L13758`).
- Not in the leaves and not rebuilt by LOAD: the word list `+0xa8` (data
  `+0xac`, count `+0xb0`), the point array `+0x110` (`MAGIC-274`), the trail
  `+0x138` (`SAV-1193`), and the cached x/y `+0x28/+0x2c`, which the base
  constructor sets to 0 (`L13759`, `L13760`). Outside the record: the
  picture-13 entries in hash view `+0xa7c`, tile bit 0x2000 (`SAV-1150`) and
  the `+0x124` marks of pictures 20 and 30 on their target.
- Rebuilt later by the driver: the point array of picture 34 on each call
  whose target is found; for picture 36 the array is cleared and rebuilt from
  the empty word list, so a loaded 36 draws no segment.

**Confidence.** High: both loops and the constructors are complete
listings.

**Evidence.** [EXP-0509](../experiments/EXP-0509-flight-records/),
`evidence/q5-listings.txt`, `evidence/q2-listings.txt`.

### SAV-1203

- `R0614` derives the cell `+0x34/+0x38` (cached x/y >> 8) and the screen
  fields `+0x3c..+0x68` from `+0x28/+0x2c`. The producers store x/y into
  `+0x28/+0x2c` before calling it; LOAD does not, so LOAD's call derives from
  (0,0). The view rebuild `R1657` (from the view frame `R0379`) calls
  `R0614` for every record and then slot `+0x40` (`R2160`), which does
  not write `+0x28`.
- The driver writes `+0x28/+0x2c` := x/y only in its tail (`L13400`,
  `L13401`), after the picture switch, and returns. It does not call
  `R0614` (none of that routine's 12 rel32 callers in
  `evidence/tables.txt` lies in `R0558..L13621`) and stores nothing at
  `+0x34/+0x38`. So the derived cell stays (0,0) after LOAD until
  `R0614` next runs after a driver tail, for example in the view rebuild.
  The first call after LOAD always runs its arm with cell (0,0); the second
  does too when no rebuild ran between the two calls.
- Arm 13 reads `+0x34/+0x38` at actionphase 4 (hash registration) and 8
  (burn), after the increment. A record saved at actionphase 3 or 7 reaches
  that value on its first call after LOAD; one saved at 2 or 6 reaches it on
  its second call, displaced when no rebuild intervened.
- Unicorn run of the original window `L03001..L13761`, map 64 × 64,
  zero planes: derived cell (0,0) writes 0x2000 at tile indices −65, −1, 63,
  −64, 0, 64, −63, 1, 65; control cell (10,12) writes the 3 × 3 at 713..843.
  For width w the indices are −w−1, −w, −w+1 and −1 before the plane, and 0,
  1, w−1, w, w+1 in it, w−1 being cell (w−1, 0). A word is skipped only when
  (word >> 6) & 0x7f is 8 to 11 and the terrain byte at the same index is 0.
- Corpus: `game0024` holds picture 13 at saved actionphase 7, so its first
  call after LOAD is the burn call.

**Confidence.** Medium: the loader, the derivation and the driver tail are
listed and the arithmetic is run, but writers of `+0x28/+0x2c` outside the
loader, the producers and the driver were not enumerated over `.text`.

**Unknown.** The order of the view frame `R0379`, and so of the
rebuild, against the 0x401 pass after LOAD, which decides the actionphase 2
and 6 case. What the four words before the tile plane hold and what the OR
does to them; whether a frame drawn between LOAD and the first call shows
the record at the (0,0) screen position. A native watchpoint on the tile
plane after loading `game0024` settles the first and the last.

**Evidence.** [EXP-0509](../experiments/EXP-0509-flight-records/),
`evidence/q5-listings.txt`, `evidence/q5-arm13-run.txt`,
`evidence/q4-listings.txt`, `evidence/corpus-projectiles.txt`.

### SAV-1204

- The cast producer and arm 0x8b take the record ID from the counter at
  view `+0xa0c` before the first driver call; a record with 0 segments
  returns 0 on that call and is destroyed in that sweep. The counter is not
  given back.
- Cast producer: it runs in the actor sweep of the 0x401 pass, which comes
  before the record sweep (`ANIM-113`, `MAGIC-261`), so the record dies in
  the pass that built it (`ANIM-147`) and no SAVE can fall between.
- Arm 0x8b: in the timer entries `R0260` and `R0454` the client
  dispatcher runs before the 0x401 pass (`ANIM-113`), so the record dies in
  the same timer step. The other timer entries and the other callers of
  `R0509` were not read.
- Corpus scan, owner saves read in place, one engine-written file excluded:
  133 files, 90 with a Projectiles store, 1 reader failure, 7 files with a
  record, 6 distinct by hash, 8 distinct records. Pictures: 1 (one), 2
  (four), 5 (one), 10 (one), 13 (one). Every held record has non-zero
  segments except `game0022`'s picture 1, saved with 0 segments.
- Raw class names in the same files: SpellTransport 3, AreaEffect 5,
  PointEffect 3.

**Confidence.** High for the cast-producer clause (listings and the pass
order of `ANIM-113` and `MAGIC-261`). Medium for arm 0x8b outside the two
timer entries `ANIM-113` reads. The corpus counts are measurements of the
present corpus.

**Unknown.** Whether `game0022`'s picture-1 record, saved with 0 segments,
is a unit shot between its last move and its removal call, as for
`ANIM-150`.

**Evidence.** [EXP-0509](../experiments/EXP-0509-flight-records/),
`evidence/corpus-projectiles.txt`, `evidence/q1-listings.txt`.

## Stored counter pair at mission start

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1210 | In 28 saves of five owner directories the head's full tick equals floor((sub-tick + 1) / 16) (one file is 15 modulo 16); the restart slot `game9999.sav` holds sub-tick 1, full tick 0 in 5 of 5, and the two mission-0 saves hold 0, 0. | High | ● active | [EXP-0521](../experiments/EXP-0521-session-clock/) |

### SAV-1210

- Population: every `.sav` in five dated directories of the owner save root (`2026-08-02`, `-08-12`,
  `-08-14`, `-08-15`, `-08-24`), 28 files, read in place with `tools/savflag`; the head's first two
  dwords are `world+0x04` and `world+0x00` (`SAV-HEAD-025`). The five are the fixed list in the
  experiment's `regen.sh`; no selection rule is recorded. The save root holds later dated
  directories, which were not read. The census output does not record the files' edition.
- floor((sub-tick + 1) / 16) is the full-tick count the stepper reaches from (0, 0) after that many
  sub-ticks (`SESS-086`); it differs from `SAV-HEAD-025`'s `sub-tick >> 4` only when the sub-tick
  is 15 modulo 16. One file of the 28, `2026-08-02/game0003.sav` (sub-tick 9343, 15 modulo 16, full
  tick 584), is such a file: it is the file that separates the two formulas and fits only
  `(sub + 1) >> 4` (`sub >> 4` is 583). The other 27 fit both.
- `game9999.sav`, the slot whose default label is "Restart last mission": (1, 0) for missions 20,
  10, 10, 41 and 30. Sub-ticks over the other 21 mission saves range from 40 to 15 017.
- `game0010.sav` in `2026-08-02` and `2026-08-15`, mission number 0: (0, 0).

**Confidence.** High for the measurement over these 28 files. The edition of the files is not recorded by the census, and
no file of the later dated directories was read.

**Evidence.** [EXP-0521](../experiments/EXP-0521-session-clock/), `evidence/c01-save-counters.txt`.

## Mover records in the restart slot

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1213 | In 9 distinct original `game9999.sav`, 421 spawner-placed Unit records hold exactly the Mover placement predicts; 22 carried player-1 Humans hold the join walk's constructor Mover with rate = their speed's low byte. | High / Medium | ✔ promoted (amended) | [EXP-0523](../experiments/EXP-0523-start-motion/) |
| SAV-1214 | One restart slot holds a hero whose static (13) and dynamic (4) embedded route lists lie 42..50 cells from its saved cell; every other record of the 443 holds two empty lists. | High / Medium | ✔ promoted | [EXP-0523](../experiments/EXP-0523-start-motion/) |

### SAV-1213

- Population: every `game9999.sav` under the owner save root, 10 files in 10 directories, 9 distinct by SHA-256 (one directory holds a byte copy of another's file). `tools/savdoc -mode mover-start` decodes all 443 Unit records (443 walks complete) and reads the 180-byte block `*(Unit+0x154)` that `unit()` already bounds, the route-list counts, order bytes `+0x08`/`+0x09` and the speed word `+0x8c` (`evidence/c01/units.tsv`, `evidence/c01/offsets.tsv`, `evidence/c03-census-summary.txt`). The restart slot is saved after the mission start's one sub-tick (`SESS-086`); `SAV-1210` reads sub-tick 1, full tick 0 in the five files of `2026-08-02`, `-08-12`, `-08-14`, `-08-15` and `-08-24`. The counters of the other four distinct files (`2026-08-27`, `2026-09-27`, its `oldsaves7` and `2027-09-07`) were not read.
- Spawner-pattern, 421 records (312 Unit, 109 Human; players 1..6), Mover `+0x82..+0x84` not all zero:
  - `+0x00` = `+0x01` on 421, values 0x40..0xBF (Unit 64..191, Human 64..190), none at 0xC0;
  - `+0x05` 0x41 on 318, 0x44 on 73, 0x82 on 30; `+0x08` 5, `+0x09` 0xff, word `+0x06` 0 on all;
  - `+0x0a` equals the low byte of `+0x8c` on 109 of 109 Humans and on 68 of 312 Units;
  - `+0x82`/`+0x83` equal the saved cell and `+0x84`/`+0x85` the saved sub-cell 0x80, 0x80 on 421;
  - word `+0x72` equals `+0x8c` on 103 of 103 records with mask 0x44 or 0x82; on the 318 with mask 0x41 an integer cost byte k with `(speed * 8) // k = +0x72` exists for every record, k from 6 to 15;
  - every other Mover byte is 0 on 421; both route lists are empty, order bytes `+0x08`/`+0x09` are 0, no Effect is attached.
- Walk-pattern, 22 records, all player 1 Humans (9 of them `*(Player+0x34)`): `+0x00`/`+0x01` 0, `+0x05` 0x41, `+0x08` 5, `+0x09` 0xff, `+0x72` and `+0x82..+0x85` 0, every other byte 0, order bytes 0; `+0x0a` equals the low byte of `+0x8c` on 22 of 22, 15..20, which differs from the constructor's 0x10 on 16 of them.
- Fit: the spawner-pattern records match `MOVE-119`, `MOVE-120` and `MOVE-121` field for field: facing from the constructor draw, mask by domain, the Human rate from derive, the five occupy fields, and the constructor's zero elsewhere. The walk-pattern records match `MOVE-122`'s replacement except byte `+0x0a`.
- Contradiction with the walk alone: byte `+0x0a` of the 22 carried records holds the derived rate, so a producer wrote it after `L14157` and before the save. The Human derive store `L03932` is the producer of that value named in `MOVE-RATE-053`; its call site on this path is not identified.
- Within this population no write between placement and the save left a spawner-placed Mover byte different from its placement value.

**Confidence.** High for the measurement over these 9 files. Medium for the inference that nothing changed a spawner-placed Mover between placement and the save, over the five files with a recorded sub-tick 1 (`SAV-1210`): a write of the same value is invisible. Unknown for the same inference over the other four files, whose sub-tick is not recorded, and for same-value writes. The 3 player-1 spawner-pattern records are all in the `2026-08-15` file. Under independent uniform `rand()` values about 3.3 of 421 facings would be 0xC0 (`MOVE-119`); none is, an outcome of probability about 0.04.

**Unknown.** The mission number, edition and sub-tick of the four files outside `SAV-1210` are not recorded by this census; the edition of the other five is not recorded either. Which call runs Human derive on the carried members before the save.

**Amended.** `SAV-1221` records sub-tick 1, full tick 0 in all 9 distinct files, the four outside `SAV-1210` included, and `SAV-1225` lists their missions (30, 141, 141 and 30). The inference over those four files (nothing changed a spawner-placed Mover between placement and the save) moves from Unknown to Medium, the grade of the five files `SAV-1210` read; same-value writes stay invisible. Their edition stays unrecorded.

### SAV-1214

- The census (`SAV-1213`) finds one record with non-empty embedded lists: the `*(Player+0x34)` hero of the `2026-08-02` restart slot, a mission-20 file in `SAV-1210`, with 13 static (`Unit+0x15c`) and 4 dynamic (`Unit+0x178`) elements. The other 8 heroes, 13 carried companions and 421 spawner placements hold 0 and 0.
- `tools/savdoc -mode mover-route` over that directory (`evidence/c03-census-summary.txt`, distances only): every consecutive element pair is Chebyshev distance 1, and the elements lie 42..50 cells from the hero's saved cell (static) and 42..45 (dynamic).
- Reading: these lists are a route the hero held before the walk placed it, carried across the mission edge. The join walk does not access the lists (`MOVE-122`) and the end-of-mission cull resets other actor fields only (`PARTY-ENDCULL-026`).

**Confidence.** High for the counts and distances. Medium for the reading that the lists come from the previous mission: the route's cells were not compared with the previous map and no other restart slot carries a non-empty list.

## Group and order blocks in the restart slot

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1218 | In 9 distinct owner-root `game9999.sav`, 133 map Groups at order 1 hold the guard setter's AI block (centroid and spread recomputed; `+0x2b`, `+0x2c` within bounds); 19 add `+0x48` 1, and no other AI byte is nonzero. | High / Medium | ✔ promoted | [EXP-0524](../experiments/EXP-0524-group-start/) |
| SAV-1219 | In the same files all 13 Player-1 Groups hold order 3; the one map Group holds only `+0x20` 3 and `+0x45` 1, while 12 carried Groups hold AI bytes no store read here writes, from about four independent lineages, and 4 hold 3 or 4 members. | High / Medium | ✔ promoted | [EXP-0524](../experiments/EXP-0524-group-start/) |
| SAV-1220 | In the same files all 443 members hold order `+0x00` = their saved cell word, `+0x08` = `+0x09` = 0 and an empty `+0x90` list; their other nonzero order bytes are only `+0x14`, `+0x71` and, on Units, `+0x40`, `+0x44` and one spell slot. | High / Medium | ✔ promoted | [EXP-0524](../experiments/EXP-0524-group-start/) |
| SAV-1221 | All 10 `game9999.sav` under the owner save root, 9 distinct by SHA-256, hold sub-tick 1 and full tick 0, the four distinct files `SAV-1210` did not read among them. | High | ✔ promoted | [EXP-0524](../experiments/EXP-0524-group-start/) |

### SAV-1218

- Population: every `game9999.sav` under the owner save root, 10 files, 9 distinct (`SAV-1221`), read in place with `tools/savdoc -mode group-start` (`evidence/c01`, summary `evidence/c02-census-summary.txt`). The census reads the 80 AI bytes that `group()` bounds and masks the list pointer `+0x4c`..`+0x4f`; it recomputes three fields from the members' saved cell and sub-cell (the member sight `actor+0xa5` is not read, so `+0x2b` is not recomputed and `+0x2c` is only bounded): fine = cell * 256 + sub-cell, the per-axis mean truncated, `+0x28` = the means' high bytes, `+0x24` = (y << 16) + x, and `+0x2a` = the largest Chebyshev cell distance to that cell.
- 146 Groups (443 members): 133 map Groups of Players 2 and above at order 1, 1 map Group of Player 1 and 12 carried Groups of Player 1 (`SAV-1219`). A Group is a map Group when every member holds the spawner Mover pattern of `SAV-1213` and carried when every member holds the walk pattern.
- On 133 of 133 guard Groups: `+0x24`/`+0x28` equal the recomputed centroid; `+0x2a` equals the recomputed spread (the Chebyshev rule is empirical: `R0167` was not read); max(`+0x2a`, `+0x2b`) <= `+0x2c` <= `+0x2a` + `+0x2b`; word `+0x00` equals word `+0x28`; `+0x2d` = `+0x38` = max(`+0x2c`, 8), with 47 Groups floored to 8; `+0x39` 0; `+0x45` 1; `+0x48` 1 on 19 and 0 on 114. `+0x2b` takes 4, 5, 6, 7, 8 and 12. Every other AI byte is 0, both word lists are empty, `Group+0x40` is 0 and `Group+0x44` nonzero. Member counts run from 1 (51 Groups) to 15.
- Fit: `AI-450` (setter and centroid), `AI-449` (`+0x45` 1 from the constructor), `AI-452` (`+0x48`), `AI-454` (no dispatch or rebuild before the save).
- SAVE writes the 80 bytes raw from the live block, then the `+0x4c` list's values; `+0x4c` itself holds the live allocation address, which LOAD replaces (`SAV-GRPSAVENEXT-572`, `SAV-1113`). `Group+0x1c` is written literally (`SAV-GRPFIRSTSAVE-579`). What SAVE writes for `+0x40`/`+0x44` is not re-read here; their remap through the identity map is LOAD's (`SAV-GRPLOAD-560`).

**Confidence.** **High** for the measurement over these 9 files and for the fit of `+0x24`, `+0x28`, `+0x00`, `+0x2a` and the derived fields. **Medium** for `+0x2b` (not recomputed), `+0x2c` (bounded only) and the `+0x2a` rule (empirical). **Medium** that no routine wrote these fields between the setter and the save: a write of the same value is invisible to the census. Authorship: 5 files original-authored, 1 original resave of a project input, 3 in directories with no recorded class (`SAV-1221`).

**Unknown.** The edition of each file; the census does not record it.

### SAV-1219

- Player 1 holds 13 Groups over the 9 files, all at order 3. Population and instrument as `SAV-1218`.
- The map Group (file `e0515abeb683`, 3 Humans of the spawner pattern): every AI byte 0 except `+0x20` 3 and `+0x45` 1, both lists empty. This is `AI-449`'s constructor block with `AI-451`'s order.
- The 12 carried Groups (22 members, 9 holding `*(Player+0x34)`): word `+0x00`, `+0x2d`, `+0x38` and `+0x39` 0 on 12. `+0x48` 1 on 11. `+0x45` 1 on 10 and 0 on 2 (files `8da6bec86028` and `29aea649c788`). On 7 Groups `+0x24`/`+0x28` hold a cell that is not the members' recomputed centroid, with `+0x2b` 6 and `+0x2c` 6 or 7, and word `+0x0a` nonzero; on 6 of them `+0x44` is 0x10 or 0x13. The other 5 hold 0 in those fields. Member counts: 1 on 8 Groups, 3 on 2, 4 on 2. `Group+0x1c` is nonzero on 7 and 0 on 5; `Group+0x40` 0 on 12.
- No store read here writes a carried Group's `+0x0a`, `+0x24`..`+0x2c` or `+0x44`, and none writes `+0x45` 0: the Stand Ground setter stores only `+0x20` (`AI-451`), the constructor writes `+0x45` 1 (`AI-449`) and the binder only `+0x48` (`AI-452`). These are not the guard setter's output either: it writes `+0x00` = `+0x28` and `+0x2d` = `+0x38` of at least 8 (`AI-450`), and the 7 stale Groups hold 0 there. The producer of the stale bytes is Unknown. They fit values left from before this mission start (the reuse of the surviving Player and its Groups, `PARTY-PERSIST-028`, and the walk's use of each member's existing `actor+0x70`, `AI-449`); a write earlier in this same mission start, at a position the member later left, is not excluded.
- Independence: the 7 stale Groups are not 7 observations. `d954bd394473` Groups 1 and 2, `349726ca4569` Groups 1 and 2 and `d8e6f1d21a7d` Group 1 share centroid, `+0x0a`/`+0x0b` and `+0x44` values (`+0x24` 192942208 and 209719424), so about four lineages remain with `5822c37e8fa5` and `e0515abeb683`. The 5 other carried Groups (`f7bc275c0eca`, `48a2a2669b7d`, `d8e6f1d21a7d` Group 2, `8da6bec86028`, `29aea649c788`) are the constructor + Stand Ground + `+0x48` shape and discriminate nothing between a carried and a fresh Group.
- Provenance: both Groups with `+0x45` 0 (`8da6bec86028`, `29aea649c788`) come from files in directories with no recorded authorship class (`2026-09-27`, `oldsaves7`); `349726ca4569`, a stale-value file, is the third such file (`2027-09-07`).
- SAVE writes these Groups as `SAV-1218` states.

**Confidence.** **High** for the measurement over 9 files and 12 carried Groups. **Medium** for the reading that a carried member keeps the reused Player's Group, and **Unknown** for where the stale bytes were written: the carry arm `L07346` (with its Group build `L14238`), the routines that carry the Player across the edge and the AI slot of the previous mission were not read here, and the 7 stale Groups form about four independent lineages.

**Unknown.** Which routine writes `+0x24`..`+0x2c` and `+0x0a` without `+0x2d` (before the mission start or earlier in this one). Whether a carried Group's `+0x48` 1 was set by this mission's binder through a hero reference or carried. Which routine left `+0x45` 0 on two Groups; `AI-ACTIVITY-324`'s rebuild is one producer of that value.

### SAV-1220

- Population and instrument as `SAV-1218`; the census reads the 148 order bytes that `unit()` bounds and masks the list pointer `+0x90`..`+0x93`.
- On 443 of 443 members: word `+0x00` equals the record's saved cell word, `+0x08` 0, `+0x09` 0, and the `+0x90` list count 0.
- 312 map Units (order 1): `+0x14` 1, 4, 5 or 8; `+0x71` 0 (9), 1 (194), 2 (51), 3 (36), 4 (20) or 6 (2); `+0x40` and `+0x44` one value pair per type word over 13 type words (`+0x44` 3, 4, 5, 8, 0xa or 0x3f; `+0x40` 0xf, 0x1e or 0x3c); on 24 Units one spell slot, `+0x78` a spell id and `+0x84` 3270 = 10 * 0x147.
- 106 map Humans (order 1), 3 map Humans of Player 1 and 22 carried Humans (order 3): `+0x14` 1, 4, 5 or 6, `+0x71` 1, no other nonzero byte.
- Fit: the constructor's zero block with `+0x71` 1 (`AI-449`); the post and `+0x08` from the stance setters (`AI-450`, `AI-451`); `+0x40`, `+0x44` and `+0x71` from the Units table stream (`AI-WITHDRAW-027`), which leaves a Human's `+0x71` at the constructor's 1; the spell slot from the spawn setup (`AI-341`; `w04` holds its only two order-block stores, `L12607` and `L12609`).
- SAVE writes the 148 bytes raw from the live block, then the `+0x90` list's values; `+0x90` holds the live allocation address (`SAV-GRPSAVENEXT-572`, `SAV-GRPPATROL-570`).

**Confidence.** **High** for the measurement. **Medium** that no routine wrote these fields between the setters and the save (same-value writes are invisible) and that `+0x14` is `actor+0x12c`: the census does not read a saved reach to compare.

### SAV-1221

- `evidence/c01/files.tsv`: the head's first two dwords (`SAV-HEAD-025`) are (1, 0) in all 10 files: `2026-08-02`, `-08-12`, `-08-14`, `-08-15`, `-08-24`, `-08-27/EXP-0261-owner-runs`, `-08-30/EXP-0278-human-runtime-en` (a byte copy of `2026-08-24`'s file), `2026-09-27`, `2026-09-27/oldsaves7` and `2027-09-07`.
- `tools/savdoc`'s authorship classes: 5 original-authored, 1 original resave of project input (`2026-08-27`), 3 in directories without a recorded class (`2026-09-27`, its `oldsaves7`, `2027-09-07`). The decoded walk completes on all 10. Both carried Groups with `+0x45` 0 come from the first two unclassified files.
- With `SESS-086`, every restart-slot save of this root was written after the mission start's one sub-tick and before a second.

**Confidence.** **High** for the measurement over these files. The edition of each file is not recorded.

## Cell planes in the restart slot

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| SAV-1223 | In 9 distinct `game9999.sav` (5 maps), the 1,653 record rows and 16,449 constant border rows equal the rows the model builds from the map ingest, saved Buildings, actors and Sacks, and type-9 casters. | High / Medium | ✔ promoted | [EXP-0525](../experiments/EXP-0525-start-planes/) |
| SAV-1224 | In the same 9 files the 1,653 cell records are the set the saved objects predict; baselines, occupant slots and the +0x2c..+0x33 tail match, and every cost baseline holds the map.reg value. | High / Medium | ✔ promoted | [EXP-0525](../experiments/EXP-0525-start-planes/) |
| SAV-1225 | The four distinct restart-slot files `SAV-1210` did not read are missions 30, 141, 141 and 30; each is at sub-tick 1, full tick 0 (`SAV-1221`). | High | ✔ promoted | [EXP-0525](../experiments/EXP-0525-start-planes/) |
| SAV-1226 | Every saved row whose static byte lacks a low bit the ingest sets is a structure cell opened by the &0xfa arm: bit 0 on 64 rows and bit 2 on 15 rows over 9 files, 9 bit-0 rows per mission-10 file. | High / Medium | ✔ promoted | [EXP-0525](../experiments/EXP-0525-start-planes/) |

### SAV-1223

- Population: every `game9999.sav` under the owner save root, 10 files, 9 distinct by SHA-256 (the `2026-08-30` file is a byte copy of `2026-08-24`'s); maps 20, 10 (2 files), 41, 30 (3 files) and 141 (2 files), so the 9 files cover 5 distinct maps. Authorship by `SAV-1221`'s classes: 5 files original-authored, 1 original resave of a project input, 3 in directories with no recorded class. All 9 decode completely with `tools/savdoc -mode start-planes` (`evidence/export-list.txt`).
- Model (`tools/startplanes`): the map's ingest (`TERR-PASS-049`, the `tools/terrpass` transcription, map.reg costs), then the saved Buildings' footprints through the recompute rule of `TERR-CELLREC-146` and `TERR-STRUCT-071`, the 443 live actors over their n by n footprints (slot by Mover mask: 0x41 and 0x44 ground, 0x82 air), the type-9 cell casters (`TERR-229`) and the Sacks. Predicted rows: window cells `[0x807, 0xedee)` with dynamic above 0x0f (`TERR-PASS-053`).
- Result (`evidence/census/summary.tsv`): rows saved 18,102, predicted 18,102, equal in both bytes 18,102; rows differing, missing or extra 0 in every file.
- Composition: 16,449 border rows (0x1f static, 0x1f dynamic, a constant on the 8-cell border) and 1,653 record rows; every non-border row is a cell record. A constant matches the border rows, so the discriminating rows are the 1,653 record rows (9.1 percent of 18,102).
- Input origin. Independent of the save: the border, the ingest base bytes (map and `map.reg`, including the cost), the type-9 casters (`TERR-229`, read from the map) and the anchors of the type-4 and ground type-8 records, which equal the saved Building anchors (220 of 220) and Sack cells (37 of 37). Taken from the saved file: each Building's size, `Present` and `Pass` masks; each actor's cell, Mover mask and key; the Sack cells; and the actor footprint side (read from the saved cell table's slots). The census compares the save with a model built from the save's own objects for the record set, bit 5, bits 6 and 7 and the Building slot, and the `Present` and `Pass` masks that decide the `|5` outcome (762 and 876 rows) and the `&0xfa` outcome (64 and 15 rows) are not compared with install data.
- Bit attribution of the saved rows (`evidence/census/bits.tsv`): bit 5 by record source: structure 1,113, live actor 487, structure and actor 8, Sack 37, type-9 caster 4, structure and caster 4. Dynamic bit 6 on 449 ground-actor cells and bit 7 on 46 air-actor cells, none elsewhere. Bits 0 and 2 off the border come from the ingest (187 and 38) or the structure `|5` arm (762 and 876). Bits 1, 3 and 4 occur only on border rows. No saved bit lacks a model source.
- Map-file checks (`evidence/census/crosscheck.tsv`): the type-4 anchors equal the saved Building anchors (220 of 220) and the ground type-8 anchors equal the saved Sack cells (37 of 37).
- EN and RU: one `rom.exe` (SHA-256 `942e9b72…7d03`), the five maps byte-identical on both roots, the map.reg Cost values identical, and the census output identical row for row.

**Confidence.** High for the measurement over these 9 files (5 maps) and for the inputs independent of the save: the border, the ingest base bytes and the type-9 casters. Medium for the rest, which is the save's consistency with its own objects: given the saved Buildings, actors and Sacks, whose masks and footprints are not compared with install data, the model leaves no row, record or bit unexplained (0 found). That excludes an unexplained record or bit; it does not separate the recompute and placement rule from a different rule that fits the same saved objects. The actor footprint side comes from the saved slots (`SAV-1224`), so bit 6 and 7 placement checks the recompute rule, not the footprint size.

**Unknown.** The edition each file was written in; both roots give the same inputs.

### SAV-1224

- The 9 files (5 maps) hold 1,653 cell records (174 to 197 per file). The model predicts the same 1,653 cells, none more or fewer (`evidence/census/records.tsv`). The record set is derived from the saved Buildings, actors and Sacks plus the map's type-9 casters (`SAV-1223`'s input origin), so its equality is the save's consistency with its own objects.
- Per record, against the model: static baseline `+0x01` equal 1,653; cost baseline `+0x00` equal to the map.reg ingest cost 1,653, and on the 376 records where rom.exe's default costs differ, equal to the map.reg value 376 and to the default 0; layer count `+0x02` and all six layer slots 0 on 1,653; ground slot equal on 449, air slot on 46, Sack slot on 37, Building slot set exactly on the 1,125 predicted cells; tail `+0x2c..+0x33` equal to the type-9 prediction on 1,653, nonzero on 8.
- Footprints: every live actor's key fills an n by n block of slots anchored at its saved cell: n = 1 for 429 actors, 2 for 12, 3 for 2. No actor lies in the dead list.

**Confidence.** High for the baselines, the cost baselines and the tails, which are rebuilt from the map and `map.reg` (the cost baselines discriminate on 376 records, `TERR-230`). Medium for the record set, the occupant slots and the Building slot, which are derived from the saved objects.

### SAV-1225

- The sub-tick 1, full tick 0 of these files is `SAV-1221`'s measurement (10 files, 9 distinct; `evidence/census/summary.tsv`, columns `head_dword0`, `head_dword1`, agree). This claim adds the mission of each file.
- The four files outside `SAV-1210`'s five directories: `2026-08-27/EXP-0261-owner-runs` (mission 30), `2026-09-27` and its `oldsaves7` (mission 141) and `2027-09-07` (mission 30).

**Confidence.** High for the measurement.

### SAV-1226

- A saved row clears a low bit the ingest sets only on a structure footprint cell whose Passability bit is clear, which the `&0xfa` arm opens (`TERR-STRUCT-071`): bit 0 on 3, 9, 9, 3, 10, 5, 10, 10 and 5 rows and bit 2 on 0, 0, 0, 1, 2, 4, 2, 2 and 4 rows of the 9 files (`evidence/census/bits.tsv`). No other cleared low bit occurs.
- Each such cell also has its cost replaced by `CostCracked` in the model (`TERR-230`); the save carries the opened bytes, which LOAD applies over a fresh ingest (`SAV-LOAD-057`).
- On scn:10.alm both mission-10 files clear bit 0 on 9 cells, the count `SAV-BLOCK-011` left as its Unknown.

**Confidence.** High for this population. Medium that the 9 cells are the same 9 `SAV-BLOCK-011` measured: its four mission-10 saves are not identified with these two files by hash here.
