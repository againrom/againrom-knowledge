# MISSION — starting a campaign mission and ending it in a win

Claims about what a campaign mission is end to end: where the player's units
come from, what puts them on a cell, which shipped file each placed person is
read out of, what carries the dialogue, and which engine parts a consumer must
build before a mission can start and finish. [`alm.md`](alm.md) says what a map
stores; [`trigger.md`](trigger.md) says what the evaluator does with the compiled
script. Spec: [`formats/mission/format.md`](../formats/mission/format.md).
Format of this file: [registry.md](registry.md).

## Map message line

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-MSGLINE-056 | In a campaign session the message line draws its oldest line at (8, 8) and each next one 17 px lower, left-aligned in font1 with a 1-px dark shadow, in the colour its post chose; the text width moves nothing. | High | ● active | [EXP-0408](../experiments/EXP-0408-price-message-line/) |
| MISSION-MSGLINE-057 | In a campaign session the line keeps at most ⌊⌊H/17⌋/2⌋ lines for a map view H pixels high, 14, 16 or 22 at 640, 800 or 1024 wide; a post appends at the bottom, drops the oldest past that, and lines expire oldest first. | High / Medium | ● active | [EXP-0408](../experiments/EXP-0408-price-message-line/) |
| MISSION-MSGPOST-058 | Forty-five call sites post to the map message line: 22 only under the `-trace` switch, and 23 for key toggles, game speed, skill increases, pickups, gold, chat, player and diplomacy notices, a save notice and phase-3 console lines. | High | ● active | [EXP-0408](../experiments/EXP-0408-price-message-line/) |

### MISSION-MSGLINE-056

- The line is the list object at `view+0xa10` of the map view
  (`SESS-VIEW-028`): vtable `L06501`, constructed once by `R1275`
  (1 call, `L06502` in `R0391`). It holds a text array at `+0x04`
  (count `+0x0c`), a colour-ramp array at `+0x18`, a lifetime array at
  `+0x2c`, the last tick time `+0x40`, the elapsed time `+0x44`, the capacity
  `+0x48`, a rectangle at `+0x4c` and a redraw flag `+0x5c`.
- `R1276` draws it. It has 2 calls, 0 in orphan code: the view's
  message `0x402` arm through `R0379` (`L06503`), and `L06504` in
  `R1277` on an object read from `+0x68`.
- For `i = 0 .. count−1` it calls `R0571(x, y + off, text[i], 0,
  ramp[i], 1)` and then adds the font's frame-0 height plus 2 to `off`
  (`L06505`..`L06506`). While `campaign+0x6bc != 3`, which includes the
  campaign session, phase 2 (`SESS-PHASE-002`), `x = 8`, `y = 8` and the
  font is font1 `[L03615]`; with 3, `x = 0`, `y = 220` and font2
  `[L02677]` (`L06507`..`L06508`). Font1 cells are 15 pixels high
  and font2 cells 10 (`SPR16A-FONT-015`, `SPR16A-FONT-018`), so the pitch is
  17 or 12.
- Flag 0 leaves `x` alone in `R0767`: bit 0 subtracts the text width
  and bit 1 half of it (`L06509`..`L06510`). Every line starts at the
  same `x`, whatever its width.
- `R0571` draws the text twice: at `(x+1, y+1)` with the ramp that
  the font's `vt+0x18` returns, `L03613` (`R1278`), then at `(x, y)`
  with the line's ramp (`L06511`..`L06512`). `R0767` hands the
  ramp to the glyph blit.
- `R0572` fills each ramp once as 16 words, entry `k` packing the
  channels as `TERR-LIGHT-019` does: `L03614` white, `17k` per channel
  (`L06513`); `L03668` grey, `14k` (`L06514`); `L03616`,
  `(185k/15, 159k/15, 73k/15)` (`L06515`); `L03613`, 8 per channel in
  every entry (`L06516`..`L06517`). A player's ramp is
  `L02665 + 32·n` (`ANIM-NUM-020`).
- A line is drawn as posted. A carriage return in it would split it at
  `L06518`, but both post routines replace every carriage return with a
  space (`L06519`, `L06520`).

**Confidence.** High: the origin, the pitch, the alignment flag, the shadow
and the ramps are named instructions in complete routine listings, and the
cell heights are measured atlases.

**Unknown.** Which object the second draw call at `L06504` reaches, and on
which screen. Phase 3 is named only by its writers (`SESS-PHASE-002`); its
paced idle arm draws no map view (`SESS-IDLE-007`).

### MISSION-MSGLINE-057

- `R1279(rect)` copies the rectangle to `+0x4c` and sets the capacity
  `+0x48` (`R1279`..`L06521`). While `campaign+0x6bc != 3` it is
  `((bottom − top) / (h1 + 2)) / 2`, signed division, with `h1` font1's
  frame-0 height; with 3 it is `(screenH − 480) / (h2 + 2) + 14`, with `h2`
  font2's.
- It has 2 calls. `R1280` passes the view rectangle `view+0x8`
  (`L06522`) after snapping its bottom to whole 32-pixel rows (`L06523`);
  `R1280` has one call, `L06524` in the view's constructor.
  `L06525` in `R0099` resizes it in phase 3 from the screen
  rectangle `L03339` less 72 at the bottom.
- The map view is 480, 576 or 768 pixels high at 640×480, 800×600 and
  1024×768 (`SESS-VIEW-028`): ⌊480/17⌋ = 28 → 14, ⌊576/17⌋ = 33 → 16,
  ⌊768/17⌋ = 45 → 22. Phase 3 gives 14, 24 and 38. The panel-driven row
  writer `L06526` (`SESS-VIEW-028`) does not resize the list.
- `R0588(text, ramp, lifetime)` sets the redraw flag, word-wraps the
  text with `R0726` against the list rectangle, 480, 640 or 864 wide
  in a map session, and appends every piece with the same ramp and lifetime
  (`L06527`..`L06528`). When the count is then 1 it restarts the clock:
  `+0x40 = timeGetTime()`, `+0x44 = 0` (`L06529`..`L06530`). While the
  count exceeds the capacity it removes entry 0 from all three arrays
  (`L06531`..`L06532`) and leaves `+0x44` unchanged.
- `R1193` is the same post with font1 always. It drops a one-piece post
  equal to the newest line (`strcmp` at `L06268`, `L06533`..`L06269`).
- `R1281`, called once, from the view's message `0x401` arm
  (`L06534`), adds `now − last` to `+0x44`. When `+0x44` exceeds the
  lifetime of entry 0 (unsigned, `L06535`) it zeroes `+0x44` and removes
  entry 0: one line per call (`R1281`..`L06536`).
- Derived from those instructions: lifetimes run one after another, each from
  the call that removed the line above it; an overflow removal passes the time
  already counted to the new top line; and a post of two or more pieces into an
  empty list does not restart the clock, so its first piece is measured from the
  last tick before the list emptied.

**Confidence.** High for the capacity formula, the three sizes, the append
order, the overflow rule and the one-per-call expiry: named instructions and
the view heights of `SESS-VIEW-028`. Medium for the derived timing clauses:
they follow from the instructions, but the interval between `0x401` messages
was not measured.

**Unknown.** The interval between `0x401` messages, and the value `+0x40`
holds before the first post. The phase when the view is constructed, and
whether a campaign session can follow the phase-3 resize, which no call
undoes.

### MISSION-MSGPOST-058

- Census (`evidence/refs.txt`, `evidence/callsites.tsv`): `R0588` has
  44 direct calls from 3 owners and `R1193` 1, 0 in orphan code. A
  `rel32` scan of every executable byte finds the same 44 and 1 calls and no
  jump.
- The switch: `R0326` searches the command line for `-trace`
  (`L06537`) with `R1282` and stores 1 in `[L04662]` on a match
  (`L06538`..`L06539`). `R1282` returns the address of the match:
  the `-session"` parse adds 9 to its result (`L06540`..`L06541`). The
  flag lies in the zero-filled tail of `.data` and has 1 write and 22 reads,
  all in `R0509`, 0 in orphan code. Each read's `JZ` skips one post.
- The 22 posts under `-trace`: message `0x92` subtype 1, string 129 (`No way`),
  white, 3000 ms (`L06542`); and 21 grey diagnostics formatted from the
  literals `L06543`..`L06544` (`evidence/literals.tsv`): 30 000 ms at
  `L06545`, `L06546`, `L06547`, `L06548`, `L06549` and `L06550`;
  10 000 ms at `L06551` (`ITEM-PICT-049`); 5000 ms at `L06552`,
  `L06553`, `L04668` (`ITEM-APPEAR-024`), `L06554`, `L06555`,
  `L06556`, `L06557`, `L06558`, `L06559`, `L06560`, `L06561`,
  `L06562`, `L06563` and `L06564`.
- The 23 posts without the switch:
  1. Seven toggles in `R0819`, each under the Ctrl latch
     `[L00627]` (`AI-KEYMOD-059`), grey, 2000 ms, with the new state's
     string: vkey `0x57` strings 94..96 (`L06565`), `0x46` 97..99
     (`L06566`), `0x48` 100..101 (`L06567`), `0x55` 218..220 (`L06568`),
     `0x4c` 102..103 (`L06569`, `ANIM-NUM-020`), `0x4e` 104..105
     (`L06570`) and `0x4f` 106..107 (`L06571`).
  2. Game speed: vkeys `0x6b` and `0x6d` without the latch, while
     `campaign+0x6bc == 2` and `campaign+0x3dc == 1`, post string 108 plus the
     new speed through `R1193`, grey, 2000 ms (`L06262`..`L06271`).
  3. Skill increase: message `0x92` subtype 2 posts string `129 + k`, or
     `134 + k` when `[[view+0x3f54]+0x18c] & 2`, with `k = msg+0xe`, white,
     3000 ms (`L06572`, `L06573`). Strings 130..134 are the five weapon
     skills and 135..139 the five magic spheres.
  4. Message `0x92` subtypes 3..7: `"%s %s %s"` from strings 204..209 and
     221..226 around a player name, white, 5000 ms (`L06574`, `L06575`,
     `L06576`, `L06577`, `L06578`).
  5. The pickup (`L04948`) and gold (`L04965`) lines of
     `ITEM-PICKTEXT-145`, white, 3000 ms.
  6. Chat, message `0x91`: a sender id of 0, past the player count or naming
     no player posts the text grey for 10 000 ms (`L06579`); a known sender
     posts `name: text` in its player ramp for 10 000 ms (`L06580`), unless
     mask 4 of that player's word in `[[view+0x9b4]+0x38]` is set.
  7. Diplomacy, message `0xb9`: `"%s %s %s"` from strings 142..149, white,
     3000 ms, each followed by a sound call (`L06581`, `L06582`).
  8. Message `0xbe`: string 203 (`Your character is saved`), grey, 30 000 ms
     (`L06583`).
  9. Frame message `0x471`, only in phase 3: `n|text` posts `text` in player
     ramp `n & 0xf`, and a text without `|` posts grey, both 30 000 ms
     (`L06584`).
- Vkey letters are the Windows naming, as in `AI-KEYMOD-059`: `0x57` W, `0x46`
  F, `0x48` H, `0x55` U, `0x4c` L, `0x4e` N, `0x4f` O, `0x6b` and `0x6d` the
  keypad plus and minus.

**Confidence.** High: both censuses have 0 orphan hits and agree with the
`rel32` scan, the flag has one write, and every gate, string, colour and
lifetime is a named instruction.

**Unknown.** Callers through a code pointer outside `.rdata`. Which unit
`view+0x3f54` holds, and so whose flag selects the magic strings. Which
producers send message `0x92` subtypes 3..7, message `0xbe` and frame message
`0x471` in a single-player session.

## Start scatter

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-SCATTER-053 | Mission entry seats each unit around the drop cell at radius trunc(max(5, sqrt(n)+4)) of the flat unit count, passed as its low byte; the hero first tries radius 0. | High / Unknown | ✔ promoted | [EXP-0365](../experiments/EXP-0365-start-scatter/) |

### MISSION-SCATTER-053

- Mission-entry range L06585..L06586 computes trunc(max(5, sqrt(n)+4)) from
  the flat unit-list count n. The referenced doubles L03914/L03956 are
  4.0/5.0.
- Caller L06587 passes the computed radius low byte, r=R&255. A controlled
  n=63504 computes R=256 but calls placement with r=0.
- Primitive R1145 makes at most floor(r*r/2)+2 random attempts. Each attempt
  draws y before x through the inclusive R0861, and each offset equals
  -floor(r/2)+draw(0..r). At odd r=5 the sampled offsets are -2..+3; the
  fallback instead scans -2..+2, x outer and y inner.
- Caller L06588..L06589 tries the hero at radius 0 and keeps a success. A
  hero failure and every nonhero enter the computed-radius call.
- A primitive refusal returns 0 without the occupancy commit.
- Controls: seventeen population controls and fifteen placement controls per
  identical EN/RU executable distinguish the expression, endpoints, axis order,
  fallback and hero retry. Mutations change the additive constant and the
  inclusive retry branch.
- Evidence files: `evidence/instructions.tsv`, `evidence/measurements.json`.

**Confidence.** High for the selected original instruction ranges and the
controlled service inputs.

**Unknown.** Native RNG seeding, the full collision/footprint rules, the native
scheduler and the subsequent off-map actor lifetime. Native reachability of the
n=63504 count.

## Defeat branches and the failure panel

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-DEFEAT-045 | Primary fall, scripted outcome and automatic recovery are separate ordered branches of reporter `R0132`; campaign maps load mode zero and cannot take the recovery branch. | High | ● active | [EXP-0274](../experiments/EXP-0274-defeat-modes/) |
| MISSION-DEFEAT-046 | The actual failure panel offers Exit to Main Menu and Load Game, and neither choice is an automatic repair command. | High | ● active | [EXP-0274](../experiments/EXP-0274-defeat-modes/) |

### MISSION-DEFEAT-045

- Reporter `R0132` first requires player `+0x3d != 0` and `+0x28 == 0`.
- For unsigned latch byte `+0x3c >= 2`, only actual mode `server+0x0c != 0`,
  entry-active byte `player+0x3f != 0` and signed primary HP `< -53` reach the
  `R0823` repair at `L03850`, then `R0065` placement at `L06590`, then
  latch zero. Campaign maps load mode zero and cannot take this path. This arm
  dereferences `player+0x34` and has no null-primary guard.
- Below latch 2, a null primary or a nonzero primary stage `+0x13c` writes latch
  2, sends `0xb4` only when joined byte `+0x3e != 0` and actual mode is zero,
  then calls `L06591`, which sets HP to -50 for every actor in the flat
  ownership index.
- Otherwise script loss counter `+0xb3b4 == 1` wins over success `+0xb3ac == 1`.
  These arms have no mode test.
- Network-mode primary fall therefore suppresses this failure packet, not every
  possible scripted failure.
- Latch 2 cannot reach either counter arm. Raising the loss count beyond 1
  cannot by itself produce a later win.
- Primary means `player+0x34`, not any companion.

**Confidence.** High for the ordered gates, the counterexamples and the caller
linkage. The evidence includes no wall-clock recovery delay, no null-primary
recovery safety and no original runtime witness.

### MISSION-DEFEAT-046

- `0xb4 -> client L06592 -> window 0x433/0xff -> 0x431` constructs `L03302`,
  stores the panel at `frontend+0x110` and shows it via `R0361`.
- Constructor vtable `L06593` has `+0x48 = L03304`, a forwarder to
  `R0716`. That base accepts `0x445/0x446`, stores the result through slot
  `+0x84 = R1228`, and posts `0x44c`.
- Close `R0709` matches the actual `+0x110` pointer. Result
  `0x445 -> 0x41e -> R1283` teardown -> `0x421` -> `R0816` main menu when
  the screen word is zero. Any other result -> `0x418`, the save-selection
  panel.
- A later successful selection `0x419` takes the campaign load path `R1284`,
  not the dead-primary repair helper.
- Draw/init slot `+0x78 = L06594` uses title main[141], dialogs[44] for
  `0x445` and dialogs[35] for `0x446`. EN labels are `~Exit to Main Menu` and
  `~Load Game`; RU labels are `Выйти в ~главное меню` and `~Восстановить игру`.
- Load is disabled when `R1169` finds no `game*.sav`.
- The neighboring vtable's handler `R0708` does post `0x41d`, but it is not
  this panel's.

**Confidence.** High for the raw PE vtable/message tables, the localized text
identities and the complete close-selector paths. This is static dispatch, not a
witnessed GUI or a successful load.

## Party entry and the drop cell

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-START-001 | `R0065` puts the player's own unit list on the map and never reads the type-6 array; a campaign map usually places nobody for the player, but 5 of the 28 do. | High | ● active (amended, partially retracted, superseded) | [EXP-0096](../experiments/EXP-0096-campaign-mission/) |
| MISSION-DROP-002 | The map's whole contribution to the start is one packed cell, and the engine picks from the drop array at random rather than taking the first entry. | High | ● active (amended, superseded) | [EXP-0096](../experiments/EXP-0096-campaign-mission/) |

### MISSION-START-001

- `R0065`, the routine that puts the player on the map, never reads the
  type-6 array. It iterates the player's own unit list at `player+0x20`
  (list-iteration calls to `R0028` at `L06595` and `R0029` at `L06596`) and calls the
  placement routine `R1145` on each.
- The hero (`player+0x34`, tested at `L06597`) goes to the exact drop cell
  with radius `0`. Everything else goes to the same cell with a computed radius
  `ftol(max(5.0, sqrt(nUnits) + [L03914]))` (`L06598`…`L06586`). The
  routine logs `"Error - can't place hero from previous mission on map."`
  (`L06599`) when both attempts fail.
- Its second half is the only place a map could contribute a starting unit, and
  it is dead on shipped data. It looks up a node named `"Humans"`
  (the string argument `L06600` pushed at `L06601`, then the call to `R0475` at `L06602` on `server+0x44`,
  `MISSION-DROP-002`'s map sub-object) and builds one unit per child from a
  `name#id` string. No shipped map has such a node: a label sweep over all 1477
  nodes of all 38 maps finds two containing "human", both ordinary leaf nodes
  bound to triggers (`evidence/audit.txt` §1).
  That entry point is the `World\Mission\<n>.ini` overlay's (`TRIG-INI-012`),
  which the install does not ship.
- Over all 28 campaign maps of both roots, roster slot 1 owns 19 of 2333 type-6
  records: 3 on `41.alm`, 4 on `71.alm`, 10 on `120.alm`, 1 each on `150.alm`
  and `151.alm`, and 0 on the other 23
  (`EXP-0100/evidence/slot1-census-{en,ru}.txt`, identical figures on both
  roots). The owner key is the type-5 `+0x00` id of `ALM-GRP-041`; it leaves 422
  of 2333 records naming no slot, against the rival key `+0x04`'s 2324.
  `10.alm`'s 0 of 35 is real but is not the corpus.
- The type-6 spawner `R0151` installs each record into its owner's own
  containers (`L06603`, `L00295`, `PARTY-WRITE-011`), and the placement walk
  then runs over that container. On those five maps the drop cell overwrites the
  map-authored positions of the player's own units.

**Confidence.** High for the positive half: the walk's instruction sequence is
read whole, the `"Humans"` node's corpus absence has its denominator, and
`R0065`'s three callers are enumerated (`EXP-0100`).

**Amended.** EXP-0100 narrowed the headline. The absence clause, that a campaign
map places nobody for the player, is retracted, and so is the discriminator
built on it: "the alternative predicts a nonzero slot-1 count and is excluded
outright" was an argument from a one-map sample, and the corpus has five nonzero
maps. [`retracted.md`](retracted.md) records the clause as superseded by the
slot-1 census above. The `"Humans"` lookup's base read "on `mapObj+0x44`".
`R0065`'s `this` is the server singleton (EXP-0100, `MISSION-DROP-002`),
so the lookup runs on `server+0x44`, the map sub-object; its instructions and
behaviour are unchanged. The positive half is otherwise unamended.

### MISSION-DROP-002

- `R0065`'s `this` is the server singleton `[L00285]`, not a map
  object. Its caller `R0131` passes its own `this` (`L06604`); that
  `this` carries the mission number at `+0x80` and `SESS-LOAD-009`'s three
  sub-objects at `+0x14`, and it is reached with `ECX = [L00285]` at
  `L06605`.
- The gate is `server+0x0c` and the array is `server+0x6c`, the same memory as
  `(server+0x44)+0x28`, since `server+0x44` is the map sub-object `R0128`
  is invoked on.
- Before the array is consulted, `player+0x60` is a per-player override, gated
  on `server+0x0c` being nonzero (`L06606`…`L06607`); when it is taken the
  array is not read. `server+0x0c` is `arg0 < 2` from the server's constructor
  (`L06608`/`L06609`/`L06610`), and the campaign start passes 2. On the
  campaign path `player+0x60` is therefore never read, and the drop always comes
  from the array or the random fallback (`PARTY-PERSIST-014`,
  `MISSION-DROP-011`).
- The index is `R0861(GetSize() - 1)` (`L06611`…`L06612`), a uniform
  random pick over the array, not `[0]`. On the shipped corpus every map carries
  exactly one word, so the pick is degenerate; an authored map with two would
  diverge from any consumer that hard-codes the first.
- The zero-coordinate fallback is `R0861(0x46) + 0x1e` on each axis
  independently. `R0861(n)` is `(rand() * (n+1)) / 32768`
  (`L06613`…`L06614`: the product is sign-extended to 64 bits, masked with 0x7fff and shifted right by 15) and is therefore
  inclusive of `n`. The range is 30..100, not the `rand() % 70 + 30` that
  `TRIG-DROP-013` published, which is 30..99 and differently distributed. This
  corrects two clauses of `TRIG-DROP-013`.
- Corpus: 38 of 38 shipped maps carry exactly one `0x10002` instant node,
  campaign and loose alike (`evidence/drop-locations.txt`). On 9 of the 10 loose
  root maps it is the map's only script node.

**Confidence.** High. The random index, the override and the RNG's inclusive
bound are each named instructions in two routines read at instruction level; the
38/38 census is exhaustive over the corpus. The amendment touches none of it,
because a displacement's number is what those instructions carry.

**Amended.** EXP-0100: every `mapObj+…` in this claim is `server+…`. The gate
first written as `mapObj+0xc != 0` is `server+0x0c`, identified by its
constructor rather than only located. The behaviour is unchanged; the base name
was not. [`retracted.md`](retracted.md) records the old base name as superseded.

## Win and lose authoring

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-WIN-003 | Every campaign map is won by exactly one authored action, and none of the 10 loose maps outside the campaign can be won at all. | High | ● active (amended, partially retracted) | [EXP-0096](../experiments/EXP-0096-campaign-mission/) |
| MISSION-VIP-004 | A trigger can exist only for the side effect of evaluating its conditions; that is how a protect-this-unit objective is authored, and authorship, not reference, arms it. | Medium | ● active (amended, partially retracted) | [EXP-0096](../experiments/EXP-0096-campaign-mission/) |

### MISSION-WIN-003

- Over the 28 maps in `scenario.res`, the count of instant-4
  (`Force "Mission Complete" state`) nodes is 1, on 28 of 28. Over the 10 loose
  maps at the install root it is 0, on 10 of 10 (`evidence/audit.txt` §5).
- The trigger that carries it is `once = 1` on 28/28 and takes one to three
  condition pairs.
- The action is not unique to one trigger: on `81.alm` four distinct triggers
  reference the same win node and on `130.alm` two do. A consumer binds by node
  id and does not assume a one-to-one map.
- Instant 5 (lose) is authored on 15 of 28 maps, identically on both roots
  (`MISSION-TYP-028`). Check 18 is authored on 7 of 28 EN maps and 6 of 28 RU
  maps (`MISSION-LOSE-025`).
- This is the corpus half of `TRIG-END-009`'s "there is no evaluator for
  victory". A map with two win actions, or a campaign map with none, would have
  refuted the one-authored-action reading.

**Confidence.** High. The enumeration is exhaustive over both corpora with the
instrument named: a whole-payload type-7 walk that tiles every map exactly, so
no node can be missed. The loose-map half is the discriminating case: it
predicts that a skirmish map never ends in victory, which any instant 4 there
would have refuted.

**Amended.** The lose-arm figures read "Instant 5 (lose) appears on 14 of 28 and
check 18 on 8 of 28". [EXP-0161](../experiments/EXP-0161-mission-30/) replaces
them: `MISSION-TYP-028` counts 15 instant-5 maps and `MISSION-LOSE-025` counts 7
EN and 6 RU check-18 maps, the difference being `100.alm`. The same two EXP-0096
figures are already corrected in `MISSION-TYP-010` and `MISSION-VIP-004`
([`retracted.md`](retracted.md)). The win-action census stands.

### MISSION-VIP-004

- `10.alm`'s trigger 11 compares two check-18 nodes with `!=` and has no action
  at all.
- Check 18 writes no slot (`TRIG-COND-003`), so both operands read cells nothing
  ever writes and the pattern is permanently false. Its entire effect is check
  18's own side effect, `session+0xb3b4++`, a loss, when the named unit is dead.
- Being authored arms it, not being referenced. `R0067`'s second pass
  walks every check node and gives each one the next slot in list order whether
  a trigger names it or not (`R0067` decompile, `local_a0` incremented in
  both the runtime arm and the `0x10002` arm). `TRIG-EVAL-001`'s pass evaluates
  every check once per full tick.
- `10.alm` carries one further unreferenced check node, a distance test, which
  is likewise evaluated harmlessly.
- Corpus: 10 check-18 nodes over 7 of the 28 campaign maps.

**Confidence.** Medium. The compile order and the evaluator are read at
instruction level, and the check-18 arm is `TRIG-COND-003`'s. Not discriminated:
whether the compiled check list is the whole node list or only the subset that
passes the builder's per-node validation. The builder skips its slot increment
on a validation failure, and nothing in the shipped corpus separates the two
readings.

**Amended.** The corpus clause's map count was 8;
[EXP-0161](../experiments/EXP-0161-mission-30/) (`MISSION-LOSE-025`) corrects it
to 7 ([`retracted.md`](retracted.md)). The node count 10 stands.

## Dialogue files

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-TEXT-005 | Mission dialogue is not in the map: `Send message` takes a number and the engine opens `main.res::text/battle/m<mission>/event<NN>.txt`. | High / Medium | ● active | [EXP-0096](../experiments/EXP-0096-campaign-mission/) |

### MISSION-TEXT-005

- `R0701` formats `"battle\m%d\event%02d"` (`L03546`) from the
  mission number and the message number, prefixes `"main\text\"` (`L06615`)
  and appends `".txt"` (`L03430`), in three calls at `L06616`, `L06617`,
  `L06618`. It then opens the result at `L03305`.
- The same string block holds `"main\text\battle\m%d\briefing.txt"`,
  `briefmap.txt`, `title.txt` and `tips%02d.txt`. `EnumRefs imm:` gives each
  literal 1 hit, 1 owner, 0 in orphan code.
- The message numbers a map's script raises should be exactly the event files
  shipped for that mission. Over the 28 campaign maps, 223 of 242 raised numbers
  name a shipped file; 19 raised numbers on 7 maps ship no file, and 2 shipped
  files are never raised. `10.alm` is exact both ways, 11 of 11
  (`evidence/message-files.txt`).
- The failure path of `R0701` (EXP-0098) does nothing at all: no window,
  no fallback, no message, no blocked state (`DLG-ABSENT-003`). The 19 unshipped
  raises are 19 silent no-ops, not a discrepancy. Three of them are EN-only: RU
  ships `m100/event09`, `m130/event07` and `m150/event10` from identical scripts
  (`DLG-LANG-010`).
- The files are markup with `<NPC=n,…>` tags whose `n` is an `npc.reg` section
  number, the id space `MISSION-ARM-006`'s npc arm uses.

**Confidence.** High for the path construction: the format string, both
concatenations and the open are named instructions, and the literal's reference
set is enumerated with its instrument's blind spot named. Medium for the
number→file mapping: 223/242 is corpus agreement with 21 exceptions, and the 21
exceptions are still unexplained as authoring.

## Placement arms and the definition lookup

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-ARM-006 | A type-6 placement resolves down one of four arms; the class key, not a flag, is the outer discriminator, and the NPC arm is tested before the definition id. | High | ● active | [EXP-0096](../experiments/EXP-0096-campaign-mission/) |
| MISSION-DEF-007 | The Humans definition lookup searches backwards, never tests index 0 and returns 0, a valid index, on a miss; `DataBinID == 26` is a sentinel, not an id. | High / Medium | ● active (amended) | [EXP-0096](../experiments/EXP-0096-campaign-mission/) |

### MISSION-ARM-006

- `R0151`: `CMP local_34[2],0x1a` splits the record set first.
  `+0x08 >= 0x1a` builds a `Unit` (`new(0x198)`, `R0501`) out of the
  Units collection of `world.res:data/data.bin`, keyed on parameters
  `0x1d`/`0x1e`. Everything below builds a `Human` (`new(0x1e8)`,
  `R0497`) out of the Humans collection.
- Inside the humans band the tests nest
  `if ((flags & 1) == 0) { if (defID == 0 || defID == 0xcdcdcdcd) typeID-arm else definitionId-arm } else npc-arm`.
  A record carrying both the NPC flag and a nonzero `+0x10` takes the NPC arm,
  and its own `+0x10` is never read.
- The NPC arm reaches a definition through `scenario.res::npc.reg`:
  `R1285(L02112, +0x0a)` fetches the npc record and `R0918`
  returns its `DataBinID` at `npc+0x14`.
- Corpus, all 8094 type-6 records of all 38 maps: 6672 Units, 1405
  Humans-by-definition-id, 15 Humans-by-npc, 2 Humans-by-typeID
  (`evidence/placement-arms.txt`).
- This makes `ALM-CLS-038`'s "two overriding paths" precise: they are ordered,
  and the order decides two of `10.alm`'s three script-named units.

**Confidence.** High. The branch order is the routine's own nesting read at
instruction level, and it discriminates against the standing reading directly.
The two records it disagrees about are `10.alm`'s escortee and its mage; the npc
arm's answers are the two people the shipped dialogue names, while `+0x10`'s
answers are generic.

### MISSION-DEF-007

- `R0498(id)` walks the collection at `defs+0xa0` from `GetCount()-1`
  downward (`L06619`…`L06620`). The loop stops at `> 0`, so entry 0 is never
  compared (`L06621 JLE`).
- It compares parameter `0x18`, the Humans `serverID` column, against the
  argument (the parameter number 0x18 is pushed at `L06622` and the comparison is at `L06623`) and falls out with a zero return value
  (`L06624`) on a miss.
- A consumer that treats the return as "index or error" silently places entry 0
  for every unresolved id.
- `R0918` reads `npc+0x14` and returns it unless it equals `0x1a`. In
  that case it composes the template from the npc's own `Flags` tokens: `"Mage"`
  `L06625`, `"Female"` `L06626`, `"MySex"` `L06627`, `"MyClass"`
  `L06628`, each tested by `R1286` (`L06629`…`L06630`).
- This answers what `REG-NPC-058` left Unknown. The ten `DataBinID` values it
  measured as naming nothing in `Templates.ini` include four `26`s, and 26 is
  the sentinel. `[npc21]`, `Flags="Hero,Me,Start"`, is the player's own record.

**Confidence.** High for the lookup's direction, its skipped index and its miss
value: all four are named instructions in a routine read whole. Medium for the
sentinel's meaning: the `== 0x1a` test, the four token literals and their tester
are read, but the arithmetic that turns the tokens into an id is not, and the
remaining six unresolved `DataBinID` entries, `npc25`…`npc30`, are still
unexplained. They carry the five values `42..46`, with `46` twice.

**Amended.** The Confidence paragraph read "the remaining six unresolved
`DataBinID` values `42..46`", which is five integers. EXP-0043's
`evidence/values-scenario_npc.csv` (the table behind `REG-NPC-058`) lists the
six entries: `npc25` 42, `npc26` 43, `npc27` 44, `npc28` 45, `npc29` 46 and
`npc30` 46. The count of six entries and the range `42..46` stand.

## Check slots and mission variables

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-SLOT-008 | A mission variable and a compiled check result share one array, and the shipped corpus already contains a collision, on `60.alm`. | Medium | ● active | [EXP-0096](../experiments/EXP-0096-campaign-mission/) |

### MISSION-SLOT-008

- `R0067`'s second pass assigns `local_a0`, reset to 0 at the head of the
  pass, to each check node in list order. It increments it for both a runtime
  check and a `0x10002` constant (`R0067` decompile: the two increments and
  the preset write `slot[local_a0] = value[0]`).
- `TRIG-COND-003`'s variable read (check 19) and `TRIG-ACT-004`'s variable
  writes (instants 3 and 8) address `session[0xbd34 + p0*4]` with `p0` the
  authored number.
- The two id spaces are therefore the same 100-int array. Any authored variable
  index below a map's check-node count names a cell a check overwrites every
  full tick.
- Corpus, 17 maps use variables. `60.alm` uses indices 32 and 33 while carrying
  64 check nodes, both inside the compiled range; every other map's indices sit
  above its own count (`evidence/audit.txt` §2).
- Shipped maxima: 64 check nodes of 100, 37 triggers of 1000.
- `SAV-711`: a resolution-rejected normal check does not advance the shared
  cursor `R0067` keeps at `[EBP-0x9c]` (`local_a0` above).
  The compare at `L06631` of the local flag at frame offset -0x1ac with 1, branching to `L06632` when unequal, gates the increment at
  `L06633`, so a validation failure shifts every later assignment down by one
  slot.

**Confidence.** Medium. The slot law is read from the builder's own counter, and
the register file is `TRIG-STORE-002`'s. `SAV-711` settles the question this
claim left open and confirms, rather than tightens, the general boundary's
status as an upper bound: a map's check-node count remains the ceiling, but its
highest used slot can sit below it whenever its own build rejected a check. How
many shipped maps' builds did so was not remeasured. The collision remains
established for `60.alm` on the compiled-order reading.

## `10.alm`, the first mission

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-M10-009 | `10.alm`, the campaign's first mission, decodes end to end into 13 triggers, and its win is a two-step escort gated on a kill. | High | ● active | [EXP-0096](../experiments/EXP-0096-campaign-mission/) |
| MISSION-TYP-010 | `10.alm` is structurally typical of the campaign and exceptional in one respect: its win depends on a variable set by walking a placed unit to a coordinate. | Medium | ● active (amended, partially retracted) | [EXP-0096](../experiments/EXP-0096-campaign-mission/) |

### MISSION-M10-009

- 80×80, 28 instant nodes, 16 check nodes (2 of them constants, `0` and `3`), 13
  triggers, 35 placements, 18 structures, 5 roster slots. The type-7 payload
  tiles exactly.
- The win chain is three triggers:
  - Trigger 2 fires when the hero is within 3 of `(36,51)` and the two-member
    group 1 has zero live units. It raises message 1 and hands group 2 to player
    1: the escortee joins the player.
  - Trigger 3 fires when that unit is within 3 of `(56,21)`. It raises message
    2, hands the unit back to player 2, empties its inventory into the hero
    (instant 28) and does `slot[50]++`.
  - Trigger 4 fires when the hero is within 3 of `(66,16)` and `slot[50] != 0`.
    It raises message 3 and wins.
- Trigger 1 carries the drop location `(17,66)`. Trigger 0 is `FALSE == FALSE`,
  the editor's "always" idiom, and starts two patrols. Triggers 5..10 and 12 are
  message-only. Trigger 11 is the VIP pair of `MISSION-VIP-004`.
- The two script-named units resolve through `npc.reg`: `[npc51]`
  `Flags="Human,Female,!Mage,Face"` `DataBinID=509` and `[npc52]`
  `Flags="Human,Mage,!Female,Face"` `DataBinID=510`, which `Templates.ini`
  labels `M10_Witch` and `M10_Merchant`.
- Full rendering, with the map-authored label strings withheld:
  `evidence/trigger-10.txt`.

**Confidence.** High. The whole type-7 payload is consumed exactly, every node
and every trigger is rendered against the install's own
`Description Checks.ini`/`Description Instants.ini` vocabularies, and every
parameter resolves. The win chain is a complete reading of one map, not a
sample.

### MISSION-TYP-010

- Typical: exactly one win action and exactly one drop location, like all 28; a
  `once=1` win trigger, like all 28; message-per-objective as the dominant
  action, with instant 2 at 254 of the 797 shipped instant nodes;
  group-count-is-zero (check 1, 96 nodes) and distance (checks 7 and 15, 184
  nodes) as the dominant tests.
- Exceptional: it is the only campaign map whose win depends on a variable set
  by walking a placed unit to a coordinate rather than by killing something or
  reaching somewhere.
- It carries no instant 5, which 15 of 28 maps do.
- The vocabulary the campaign uses is far narrower than the vocabulary the
  editor declares: 26 of 34 instant arms and 17 of 22 check arms appear at all.
  The tail is long: 9 instant opcodes appear fewer than 10 times each
  (`evidence/audit.txt` §4).

**Confidence.** Medium. This is corpus agreement over an exhaustive walk of both
corpora. Nothing here discriminates against a rival model of what "typical"
means, and the exceptional clause is an observation about 28 maps, not a
statement about the engine.

**Amended.** The instant-5 map count was 14 of 28;
[EXP-0161](../experiments/EXP-0161-mission-30/) (`MISSION-TYP-028`) corrects it
to 15. The 14 is in [`retracted.md`](retracted.md).

## Drop override and seating

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-DROP-011 | Exactly one instruction writes `player+0x60`, in session opcode `0x32`, and its value is the current cell of one of the player's own units; the campaign path never reads it. | High / Medium | ● active (amended) | [EXP-0100](../experiments/EXP-0100-party-origin/) |
| MISSION-SEAT-012 | Seating a unit is a bounded random-then-raster search around the drop cell; it can fail per unit, leaving the surplus off the map, and the mission starts anyway. | High / Medium | ● active (amended, partially retracted) | [EXP-0100](../experiments/EXP-0100-party-origin/) |

### MISSION-DROP-011

- This is the open item `formats/mission/format.md` names.
- Instrument: `EnumRefs disp:60`, whole image, 602 hits over 318 owners, 2 in
  orphan code: the address computation at stack offset +0x60 at `L06634` and the load from offset 0x60 of an object at `L06635`,
  both in the code-region region, far from any `Player` method, one of them a
  stack displacement.
- Reduced by the field's own `u16` width, per the rule EXP-0066 set: 42 of the
  602 are byte-wide or 16-bit-wide, and five are 16-bit-wide. They are three reads in
  `R0065` (`L06636`, `L06637`, `L06638`), the constructor's zero
  (`L06639`, `R0201`) and the 16-bit store to offset 0x60 at `L06640` in
  `R0061`.
- Not one of the 37 byte-wide hits is on a `Player`. The only two that store
  (`L06641`/`L06642`, `R1138`) write an object whose `+0x54` was set
  to `0xe` two instructions earlier and whose `+0x61` is the next field; a
  `Player`'s `+0x64` is a `CString`.
- No dword hit at displacement `0x60` lies in `L06643..L06644` or
  `L06645..L06646`, the ranges the `Player` methods occupy. A dword store
  there would have to span `+0x60` and `+0x64`.
- The writing arm (`SESS-CMD-016`'s `0x32`, at `L06647`) resolves an actor
  from the command's `(slot, unit id)` pair and reads that actor's own cell as
  two bytes (`R0299`/`R0300`). It packs them `x | (y << 8)`, the
  packing `R0065` unpacks, and stores the word on `[actor+0x14]`, the
  actor's owning `Player` (`L06648`…`L06640`).
- The override is "start the next map where this unit is standing now" where it
  is read at all. `R0065` reads it only when `server+0x0c` is nonzero,
  and the campaign start leaves that field 0, so on the campaign path the
  override is never read (`MISSION-DROP-002`). `SESS-CMD-016` graded `0x32`
  Medium for want of a name; this is the name.

**Confidence.** High that the writer is unique: a whole-image displacement
sweep, the width filter stated with both halves, its two orphan hits read and
excluded by address, and the dword band checked inside the owning class's own
ranges. High for what the arm computes: nine consecutive named instructions. High
that the campaign path never reads it: the gate at `L06606`…`L06607` and
the campaign start's argument 2 are instruction-level (`MISSION-DROP-002`).
Medium that the meaning is "enter the next mission here" where `server+0x0c` is
nonzero: the packing and the destination are read, the command's provenance is
not, and `SESS-CMD-015` already says nothing shows what fills the command pool.
This sweep's own blind spot: a `memcpy`/`REP MOVSD` over a `Player` carries no
displacement at all.

**Amended.** EXP-0155 (`TRIG-CAST-033`) identifies the object of the two byte
stores in `R1138` as the temporary casting actor for trigger instant 21,
not a message record ([`retracted.md`](retracted.md)). The writer enumeration
and the `Player+0x60` conclusion are untouched. `MISSION-DROP-002` (EXP-0100)
bounds the override's reach: `R0065` reads `player+0x60` only when
`server+0x0c` is nonzero (`L06606`…`L06607`), that field is `arg0 < 2` from
the server's constructor, and the campaign start passes 2. The clause "start the
next map where this unit is standing now" was stated without that gate; it holds
only where `server+0x0c` is nonzero, and on the campaign path the drop comes from
the array or the random fallback.

### MISSION-SEAT-012

- `R1145(x, y, r)` has two phases.
  - Phase one: `tries = r·r/2 + 1`, `half = r/2`, then up to `tries + 1` random
    attempts at `(x − half + rand(r), y − half + rand(r))`, each tested by
    `R1287` on the world (`L06649`…`L06650`).
  - Phase two, only if `r > 0`: an exhaustive raster of the whole square
    `[x ± half] × [y ± half]`, x outer and y inner; the first free cell wins
    (`L06651`…`L06652`).
- If neither phase finds a cell the routine returns 0 without committing, and
  the caller logs `"Error - can't place hero from previous mission on map."`
  (`L06599`) for any member, although the string names only the hero. On
  success `R0050` commits the actor to the world (`L06653`).
- Order and radius come from the walk: the flat index in list order. The hero
  (`player+0x34`) is tried first at the exact cell with `r = 0`, where the box
  is one cell and both attempts are that cell, since `R0861(0)` is 0. A
  hero seated there keeps that cell. A hero that fails there, and every other
  member, is tried at `r = ftol(max(5.0, sqrt(n) + [L03914]))` with `n` the
  flat index's own count (`L06654`…`L06586`, `L06597`…`L06655`;
  `MISSION-SCATTER-053`).
- The drop cell applies to all members, not one, and the box grows with the
  party. The seat is a search over a bounded box, so a party larger than the
  box's free cells leaves the surplus in the collections and off the map, with a
  log line and no other consequence. Neither the walk nor `R0131` tests
  the count against anything.

**Confidence.** High for both phases, the failure return and the radius: three
routines read at instruction level, each bound named by its own instruction.
Medium that a failed seat has no further consequence: the walk's per-actor tail
runs unconditionally afterwards and nothing there re-tests the return, but the
actor's later behaviour off-map was not followed.

**Amended.** The order clause read "Every member is then tried at
`r = ftol(max(5.0, sqrt(n) + [L03914]))`". `MISSION-SCATTER-053` (EXP-0365,
`evidence/instructions.tsv`) replaces it: a nonzero return from the hero's
radius-0 call at `L06656` skips the computed-radius call (`L06657`,
the branch to `L06589` at `L06658`). A seated hero keeps radius 0; only a failed hero and
every nonhero enter the computed-radius call. Both phases, the radius expression
and the failure return stand.

## Mission end: counters, latch, close and stepping

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-END-013 | The reporter's outcome equality tests sit behind an earlier latch and mode branch; they are not symmetric outcome-change notifications. | High | ● active (partially retracted, amended) | [EXP-0101](../experiments/EXP-0101-mission-end/), [EXP-0274](../experiments/EXP-0274-defeat-modes/) |
| MISSION-LATCH-014 | The outcome counters only rise, nothing clears them inside a mission, and check 18 increments on every evaluation, so "outcome > 0" is the wrong predicate. | High / Medium | ● active | [EXP-0101](../experiments/EXP-0101-mission-end/) |
| MISSION-PATH-015 | The failure panel closes to the main menu or to save selection through the base close handler, not through its own `0x41d` post. | High | ● active (partially retracted, amended) | [EXP-0101](../experiments/EXP-0101-mission-end/), [EXP-0274](../experiments/EXP-0274-defeat-modes/) |
| MISSION-STOP-016 | Panel pause and session teardown are different gates: campaign idle skips the stepper while an outcome panel is shown, although the run dword stays 1. | High | ● active (partially retracted, amended) | [EXP-0101](../experiments/EXP-0101-mission-end/), [EXP-0274](../experiments/EXP-0274-defeat-modes/) |

### MISSION-END-013

- `R0132` requires `player+0x3d != 0` and `player+0x28 == 0`.
- Latch byte `+0x3c >= 2` enters a separate recovery arm and cannot reach the
  counters in that invocation. Actual `server+0x0c == 0` skips recovery without
  clearing the latch. Nonzero mode additionally requires `player+0x3f != 0` and
  signed primary HP below -53.
- Below latch 2, an absent primary or a nonzero primary stage takes precedence.
  Otherwise loss `session+0xb3b4 == 1` precedes win `+0xb3ac == 1`, with
  announcements `0xb4`/`0xb5` and latches 2/1. Neither counter arm tests mode.
- Win-to-loss is locally admissible.
- `MISSION-DEFEAT-045` and `SESS-DEFEAT-064` carry the complete caller and mode
  scope.

**Confidence.** High for the named instruction paths and the discriminating
branch/vtable counterexamples. No original runtime timing or GUI observation is
claimed.

**Amended.** EXP-0274 refutes two clauses ([`retracted.md`](retracted.md)):
loss-to-win merely because the loss count passes 1, and unqualified automatic
latch clearing. Latch at least 2 is an earlier branch that cannot reach the
counters; actual mode zero skips replacement; nonzero mode needs entry-active
and HP below -53 before reset; incrementing the loss count beyond 1 is
insufficient.

### MISSION-LATCH-014

- All four writers are increments, never stores:
  - instant arm 4, `L06659`, an in-place increment of the field at `+0xb3ac`;
  - instant arm 5, `L06660`, an in-place increment of the field at `+0xb3b4`;
  - check arm 18, `L06661`, an in-place increment of the field at `+0xb3b4` (`R1288`, gated on the
    unit's state being `0x10`);
  - a fourth in orphan code: `L06662` load, `L06663` increment, `L06664`
    store, of `+0xb3ac`, inside `ORPHAN[L06665..L06666]`.
- The only routine that clears either is `R1289` (`L06667`,
  `L06668`), whose callers are three constructors (`EnumRefs callto:R1289`: 3
  hits / 3 owners / 0 orphan). No mission-end, win or lose path resets them.
- Check 18's increment is a side effect of evaluating the check, and
  `TRIG-EVAL-001` evaluates every check once per full tick. One dead VIP drives
  the counter up without bound, and the reporter matches on exactly one tick.
- Consumer consequence: a reader that fires on `counter > 0` announces a loss
  every tick. One that lets two win actions execute in the same trigger pass
  steps `0 → 2` and announces nothing.
- Corpus (`evidence/outcome-census-en.txt`, `-ru.txt`): of 28 campaign maps, 8
  EN / 7 RU can drive lose past 1, and 2 can drive win past 1: `81.alm` with 4
  win-firing triggers, `130.alm` with 2.

**Confidence.** High for the writer set and the clearing set. Both enumerations
name their instrument: `disp:` sees orphan code and is what printed the fourth
writer, and `callto:` runs on the repaired table with 0 orphan. That writer's
increment carries no displacement of its own, the case `AGENTS.md` rule 5 names.
Medium for the fourth writer's role: its dispatcher was not located, its jump
table at `L00517` has 0 references by immediate and 0 by displacement
image-wide, and 163 of the 165 bytes before its first case are never
disassembled, so its reachability is open.

### MISSION-PATH-015

- `0xb4 -> 0x433/0xff -> 0x431` builds the panel at `frontend+0x110`.
- Constructor `L03302` installs vtable `L06593`; slot `+0x48` contains
  `L03304`, forwarding the base close handler.
- Handler `R0708` belongs to the neighboring class (slot `L06669`).
- The actual close dispatcher `R0709` matches `+0x110`, then posts `0x41e`
  for result `0x445`, otherwise `0x418`. Exit takes teardown
  `R1283 -> 0x421`; with screen word zero, `R0816` opens the main menu.
  Load opens save selection.
- The separate `0x41d` phase-2 zero-win-flag branch still targets `0x41e`, but
  it is not the actual failure panel's upstream hop.
- `MISSION-DEFEAT-046` is the corrected failure programme. Success Continue's
  control and result `0x446` are `MISSION-VICTORY-029`, its handler
  `MISSION-VICTORY-030`, its input programme `MISSION-VICTORY-031` and the
  reachability it leaves `MISSION-VICTORY-033`.

**Confidence.** High for the named instruction paths and the discriminating
branch/vtable counterexamples. No original runtime timing or GUI observation is
claimed.

**Amended.** EXP-0274 refutes the clause that the failure panel's class handler
`R0708` posts `0x41d` ([`retracted.md`](retracted.md)). The failure
constructor installs `L06593`, whose message slot is `L03304`; the base
close and the `+0x110` pointer arm choose `0x41e` or `0x418`; the borrowed
`R0708` is in the neighboring vtable `L06670`. The success cross-reference
read "`MISSION-VICTORY-034` owns success Continue". `MISSION-VICTORY-034` is
mission-completion restore and the latch reset; Continue's control, result and
handler are EXP-0233's `MISSION-VICTORY-029`…`031` and `MISSION-VICTORY-033`.

### MISSION-STOP-016

- Stepper `R0147` requires `server+0x2c != 0`. Start `R0512` sets it, and
  teardown `R0826` clears it and both clocks. Construction and `R0414`
  deserialization also write this field.
- Campaign idle `R0560` skips the stepper while phase `frontend+0x6bc == 2`
  and screen mask `+0x3dc & 0x4008 != 0`. Both outcome panels use `R0361`,
  which sets bit 8 for them. The world therefore does not keep stepping until
  dismissal merely because the run dword remains 1.
- Network authority/client paths do not inherit this campaign-only idle gate
  (`SESS-DEFEAT-065`).
- Teardown clears participant `+0x3d`. Retained human-type actors in
  `[0x21,0x3f]` have HP/MP restored and stage cleared without the XP-loss leaf.
  This teardown does not clear player outcome byte `+0x3c`.

**Confidence.** High for the named instruction paths and the discriminating
branch/vtable counterexamples. No original runtime timing or GUI observation is
claimed.

**Amended.** EXP-0274 refutes two clauses ([`retracted.md`](retracted.md)): that
nothing about the outcome stops the simulation, and the old exhaustive claim of
exactly two writers of `server+0x2c`. Campaign idle checks phase 2 and screen
`0x4008` before calling the stepper, and outcome panels set bit 8. Construction
and deserialization also write the run dword, and pause precedes teardown
without requiring a run-dword write.

## Outcome scripts across the two roots

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-ROOT-017 | `rom.exe` is identical in the two releases but the campaign script is not, so a mission-outcome measurement names its root or it is not made. | Medium | ● active (amended, partially retracted) | [EXP-0101](../experiments/EXP-0101-mission-end/) |
| MISSION-VIP-018 | A protect-this-unit objective can announce its loss only because no shipped map names one unit in two VIP nodes, and that is a corpus fact, not a rule. | Medium | ● active | [EXP-0101](../experiments/EXP-0101-mission-end/) |

### MISSION-ROOT-017

- `rom.exe` is identical: sha256
  `942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`, 1 977 344
  bytes, byte for byte in `gameversions\en\rom.exe` and
  `gameversions\ru\ROM.EXE`. Every instruction this experiment cites holds on
  both.
- The scripts are not. Comparing each campaign map's type-7 payload byte for
  byte, 26 of 28 are identical.
  - `140.alm` differs with every counted quantity equal, so its 42 816 bytes
    differ only where the census does not read: the 64-byte label and name
    fields.
  - `100.alm` differs in substance: EN 65 596 bytes with 1 VIP node and 2
    lose-firing triggers, against RU 63 636 bytes with 0 and 1.
- The loose maps differ too: EN ships 10 at the install root, RU ships 6.

**Confidence.** Medium. The census is exhaustive, byte-exact over both
corpora, with its instrument named: a whole-payload type-7 walk that tiles 38/38
on EN. It does not establish why `100.alm` was re-authored, and the `140.alm`
reading that only text differs is inferred from the counts agreeing, not from
decoding the label fields.

**Amended.** [EXP-0124](../experiments/EXP-0124-script-arms/) refutes the clause
"RU's `Horror.alm` is the one map of either corpus whose type-7 payload does not
tile exactly", and with it the instrument figure "33/34 on RU"
([`retracted.md`](retracted.md)). `ru/Horror.alm` has no type-7 record: it holds
four records, typeIds 0..3, and tiles exactly. The RU corpus is 28 campaign maps
plus the 5 loose files that carry a script, and an independent walk reports
33/33 with nothing unparsed. `ALM-CORP-060` publishes the file, two-byte diff
included. The `rom.exe` identity, the 26/28 identical scripts, `140.alm`,
`100.alm` and the 10-vs-6 loose-map counts stand.

### MISSION-VIP-018

- `MISSION-VIP-004` establishes what check 18 does; this claim measures what its
  repetition costs.
- k VIP nodes naming the same unit add k to `session+0xb3b4` in one trigger
  pass, and `MISSION-END-013`'s reporter matches only `== 1`. A map with two VIP
  nodes on one unit could never announce that unit's death.
- Census over both roots, grouping every check-18 node by its `Par0`; the
  install's own `Description Checks.ini` `[Group18]` gives `Name=VIP`,
  `Par0_Value=Target_Unit`. EN: 7 maps carry check-18 nodes
  (`10 20 31 91 100 110 150`), 10 nodes over 10 distinct units. RU: 6 of the
  same set (`100.alm` carries none there).
- The maximum number of nodes on any one unit is 1 on every map of both roots.
  The predicted failure is reachable in the format and absent from the shipped
  data.

**Confidence.** Medium. A census with its denominator can show that a shape does
not occur, never that it is impossible. The mechanism half rests on
`MISSION-END-013`'s `== 1` and `TRIG-EVAL-001`'s once-per-pass, both High, but
the boundary it draws for an authored map is an upper bound.

## Opening view

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-VIEW-019 | At mission start the engine brings the view to the hero through network message `0xab`, not through a store to the view, so a whole-image sweep on the view's own fields cannot find it. | High / Medium | ● active | [EXP-0118](../experiments/EXP-0118-map-view/) |
| MISSION-VIEW-020 | The view position is also persisted under `[View] X` / `[View] Y` and restored on the mission-entry routine itself, so the opening camera is not one decision. | Medium | ● active | [EXP-0118](../experiments/EXP-0118-map-view/) |

### MISSION-VIEW-019

The chain, all of it named instructions:

1. Session command opcode `0x4`, the byte at `cmd+0x9`, reaches the arm at
   `L06671`, which calls `R0131` (`L06672`, `callto:` 1 hit / 1
   owner). The opcode is resolved by brute force over `R0061`'s own index
   table: `L02521`-style dispatch at `L06673`…`L06674`, byte table
   `L00564`, arm table `L03265`, both roots identical on all 189
   entries.
2. `R0131` is `MISSION-DROP-002`'s routine: it runs `MISSION-START-001`'s
   placement walk `R0065` (`L06675`).
3. Immediately after the walk, gated on the player having a hero
   (the 32-bit compare with zero at `L06676`, `MISSION-START-001`'s
   `player+0x34`), it fills the static message object at `L02444`, the
   buffer `DLG-PATH-002` names. It writes opcode `0xab`
   (a byte store at message offset 0x9 at `L06677`), the player's slot at `+0x7`
   (`L06678`), and the hero's own cell at `+0xa` and `+0xe`, read by
   `R0299`/`R0300` (`L06679`, `L06680`), the pair
   `MISSION-DROP-011` names for reading an actor's cell as two bytes. It then
   dispatches the message (`L06681`).
4. On the client, `R0509`'s arm for opcode `0xab` and for no other value
   (byte table `L02523`, arm table `L02524`, 189 entries, both roots
   identical) computes `ScrollTo(pkt[+0xa] − visCols/2, pkt[+0xe] − visRows/2)`
   (`L06682`…`L06683`). It is gated on `campaign+0x3c8 == 0` (`L06684`);
   mission entry sets that field from its own argument at `L06685`. This is
   `SESS-VIEW-031`'s message `0x406` and the centring expression the Alt+digit
   group hotkey uses.

The opening origin is therefore `clamp(heroCell − viewportSpan/2)` under
`SESS-VIEW-030`'s band: at 640×480 that is `hero − (7,7)`, at 800×600
`hero − (10,9)`, at 1024×768 `hero − (13,12)`.

Only two instructions in the whole image write `0xab` into a message opcode
(`EnumRefs re:.*0xab$`, 2 hits / 2 owners / 0 orphan). The second, `L06686` in
`R1290`, builds the same message on the `server+0x00 % 5 == 0` sub-tick
from a unit's cell, and no instruction in that routine hands it to a dispatcher.
Its constructor chain `R0546`→`R1291` only zeroes fields and
installs vtables `L02445`/`L06687`, so there is no self-enqueue either,
and the object dies in the frame.

**Confidence.** High for the whole chain: four routines read at instruction
level, both dispatch tables resolved by brute force over their own index bytes
rather than transcribed, and every cited address re-read from both roots. Medium
for the ordering against `SESS-VIEW-029`'s `[View] X/Y` restore: both writers
exist on the mission-entry path, and which lands last was not established
because the command queue between `R0131` and the client arm was not
followed. Medium that `R1290`'s message is inert: the absence of a
dispatch is read from that routine and its two constructors, an absence over
three routines and not over the image.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### MISSION-VIEW-020

- `R0099` is `SESS-LOAD-009`'s only campaign caller of `R0512`
  (`L06688`) and the routine that clears the view's landscape pointer on the
  way in (`L06689`). It later reads both keys through the `.reg` int getter
  `R0452` and stores the results straight onto the view
  (`L06690`…`L06691`).
- `R0098` does the same (`L06692`…`L06693`), and `R0084`
  writes the pair back out (`L06694`…`L06695`).
- The same store carries `CurrentState/InBattle`, `Character/Name`,
  `GameOptions/*`, `Objects/Group%d`, `Inventory/IsOpen`, `SpellBook/IsOpen`,
  `Projectiles/*` and `Fog/*` (`StrDump fnstr:R0099`). It is the session's own
  persisted state, not a map or a registry key.
- The restore is a no-op when the key is absent: the default handed to the
  getter is the field's current value, and the getter returns it on a miss
  (`L06696`). The pair can only carry a previously saved origin forward.
- Two entries later the same routine restores the two panels (`L01648`,
  `L01649`), which re-enters `R0387`, recomputes `visRows` and
  re-clamps the vertical origin (`SESS-VIEW-028`).
- A consumer that models the opening camera as one decision will be wrong on one
  of the two paths.

**Confidence.** Medium. Every instruction in the chain is named and re-read on
both roots, but three inputs a High grade needs are not established: which file
backs the store, whether a shipped save carries a `[View]` section at all, and
where this restore falls relative to `MISSION-VIEW-019`'s packet. The weakest
input is the store's own contents, which is not a fact about the image.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

## Documents and payment

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-DOC-021 | A document's value is formatted straight into a resource path, and the document collection is campaign state that a save carries. | High / Medium | ● active | [EXP-0151](../experiments/EXP-0151-mission-documents/) |
| MISSION-MONEY-022 | A per-mission `Payment` is a completion grant, not an entry grant. | High / Medium | ✔ promoted | [EXP-0152](../experiments/EXP-0152-money-cycle/) |
| MISSION-DOC-023 | Mission 10's three text-document entries exist before hero creation and before the first playable tick. | High | ✔ promoted | [EXP-0154](../experiments/EXP-0154-campaign-start/) |

### MISSION-DOC-021

- `R1292` forks on the element's kind and builds the path with one
  `sprintf` each: kind 0 → `graphics\interface\Docs\%d.bmp` (`L06697`), kind
  1 → `main\text\Docs\%d.txt` (`L06698`). The first component names the
  archive.
- Both resolve exactly. `main.res` ships `text/docs/1.txt`…`4.txt` and
  `graphics.res` ships `interface/docs/1.bmp`, on both roots, and the registry's
  4 text and 1 picture values cover them 5/5, with no unused file and no
  unresolved value.
- A text document is read whole by `R1293` (open, read the node's whole
  length, append one NUL, assign to element `+0x0c`) and wrapped by
  `DLG-WRAP-009`'s `R0726(element+0x24, text)` into the string array at
  `+0x10`. A picture is loaded into `+0x08` with no layout.
- Lifetime is the whole campaign, and the record is in the save file.
  `R1294` writes the collection as `u32 count` then `count` pairs of
  dwords; `R1295`/`R1296` push exactly `(this,4)` and
  `(this+4,4)`, the value and the kind, and nothing else. It runs with
  `ECX = app+0x548`, the campaign record, called at `L06699` from
  `R0084`, whose string operands are the save tail store's own keys
  (`Fog/Data`, `Projectiles/*`, `View/*`, `Character/Name`).
- `R0434` reads the collection back from `R0098` and
  `R0099`, the two mission-entry routines (`MISSION-VIEW-020`).
- Measured on all 15 preserved saves: the run sits uncompressed at the end of
  the file. Exactly three (trailer, count) readings fit every file, and only
  `trailer=32 count=3 (1,1)(2,1)(3,1)` reproduces `[Mission10] AddTextDocument`;
  the other two are sub-runs of it. The 32-byte trailer's first dword is 10, 20
  or 30 across the corpus, the campaign's mission number.
- The save's tail key-value store carries no document leaf: 19 leaves on 14
  saves, 15 on the town-screen save.

**Confidence.** High for the two path rules and the 5/5 resolution: two literals
are read from the section bytes, and the resource sets are exhaustive on both
roots. High for the record's field set and its 8-byte stride: the serializer and
both element helpers are read whole. Medium for the save location: the run was
found by walking backwards from EOF, the 32-byte trailer between it and EOF is
unattributed, and no save from a campaign past mission 30 was available. A save
past mission 50 must read count=4 with (4,1) appended, and past mission 60
count=5 with (1,0).

### MISSION-MONEY-022

- `R1297` loads `[Mission<n>] Payment` into mission-record `+0x0c`. The
  mission-completion path passes that field to `R1298`, while mission
  entry has no corresponding read (`REG-SCN-067`).
- Both preserved roots carry 24 mission sections, 9 with non-zero `Payment`:
  `1000, 700, 3000, 5000, 70000, 40000, 120000, 350000, 700000`, total
  1,289,700.
- G2: `Payment` is one 32-bit REG scalar; widening it changes `scenario.reg`
  bytes and the mission-record layout.

**Confidence.** High for the key, the field, the timing and the two-root corpus.
Medium for interpreting the value as a purse credit: `R1298` was not read
through, so the evidence identifies timing and source but does not independently
prove sign, recipient or that the callee reaches `Player+0x38`.

### MISSION-DOC-023

- The new-campaign state machine constructs the session, resets the campaign
  record, and calls `R0755(10)`. Its loader appends
  `[Mission10] AddTextDocument = {1,2,3}` as `(1,1)`, `(2,1)`, `(3,1)` to the
  collection at campaign record `+0x144..+0x154`.
- The panel reads that collection. The separate access item is not implied by
  its contents.
- Both roots carry the same values.

**Confidence.** High: instruction ordering plus both-root registry reproduction.

## `30.alm` and its side mission `31.alm`

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-M30-024 | `30.alm` decodes end to end, byte-identical on both roots, and its win is one distance test with no precondition. | High | ● active (amended) | [EXP-0161](../experiments/EXP-0161-mission-30/) |
| MISSION-LOSE-025 | `30.alm` has no script-authored way to lose, and eleven of the 28 campaign maps, the same set on both roots, are in the same position. | High / Medium | ● active | [EXP-0161](../experiments/EXP-0161-mission-30/) |
| MISSION-CURE-026 | `30.alm` hands the hero a quest item and consumes it at the destination, but nothing in the map ever tests for it. | High / Medium | ● active (amended) | [EXP-0161](../experiments/EXP-0161-mission-30/) |
| MISSION-PAIR-027 | `31.alm` is a standalone side mission attached to mission 30's town phase, not a stage of mission 30, and the `N`/`N+1` pairing does not cover the campaign. | Medium | ● active | [EXP-0161](../experiments/EXP-0161-mission-30/) |
| MISSION-TYP-028 | `30.alm` is structurally typical of the campaign and exceptional in being unlosable; `MISSION-TYP-010`'s instant-5 map count is 15, not 14. | High / Medium | ● active | [EXP-0161](../experiments/EXP-0161-mission-30/) |

### MISSION-M30-024

- 80x80, format version 990, 10 records, container and type-7 payload both
  tiling exactly: 15 instant nodes, 10 check nodes, 8 triggers.
- The same bytes on both roots: sha256 `56a10d9d…` for the archive member on EN
  and RU alike, which is stronger than `MISSION-ROOT-017`'s payload-level
  census.
- `trig[0]` is dropped by `TRIG-BIND-010`'s zero-left rule, and its drop cell
  (15,69) reaches the start through the build-time drop table instead, leaving 7
  live triggers.
- `trig[7]` is `constant 0 <= constant 0`, always true, and runs instant 12 with
  `Unit=10001`, `Item=6`.
- `trig[6]` carries the map's only instant-4 node. Its only condition is check
  node opcode 6 (`Distance between units`, `TRIG-DIST-014`'s Chebyshev metric,
  `0xff` when either unit is dead) between unit 10001, the first hero ordinal
  (`TRIG-REC-011`), and unit 56, compared `<= 3`. Unit 56 is `npc.reg [npc53]`,
  `DataBinID=511`, placed at cell (65,15). The trigger's action list is message
  16, two instant-13 removals, then the win.
- `trig[1]` and `trig[3]` read the same check node, `c1` on group 8, both
  `== constant 0` and both `once=1`, so messages 7 and 8 fire together when that
  group dies. The group-5 check node the second trigger is named after is read
  by nothing.
- Three check nodes and four instant nodes are named by no trigger, including a
  `+1000` money grant to player 1. A fifth instant, the drop location, is named
  only by the trigger the builder drops and reaches the start through the drop
  table instead.
- One label disagrees with its own opcode, and the corpus says a label is not
  evidence: of 1467 nodes, 103 carry the catalogue name of their own opcode,
  1362 carry author free text, and 2 carry a catalogue name belonging to a
  different opcode.
- Rendering with the map-authored strings withheld: `evidence/script-30-31.txt`.

**Confidence.** High. The whole type-7 payload is consumed exactly on both
roots, every node and every trigger renders against the install's own
`Description Checks.ini` / `Description Instants.ini` vocabularies, and every
parameter resolves. The same walk reproduces `MISSION-M10-009`'s reading of
`10.alm` and ten separately published corpus counts before it is pointed at this
map.

**Amended.** `10001` is the primary character, not a roster position, and an ordinal above it names a role (`TRIG-HEROORD-075`, `TRIG-HEROTPL-076`); the wording "the first hero ordinal" in this card is positional.

### MISSION-LOSE-025

- The map authors one instant-5 node that no trigger names, and it authors no
  check-18 node at all.
- Both halves matter because the two arms are armed differently. An instant-5
  node fires only when a trigger's action list reaches it. A check-18 node is
  armed by being authored (`MISSION-VIP-004`): the builder gives every check
  node a slot and the evaluator runs every check once per full tick. The lose
  test is therefore referenced instant-5 nodes plus authored check-18 nodes, and
  for this map both are zero.
- Corpus (`evidence/outcome-census.txt`): 15 of 28 maps author an instant-5 node
  and 12 of those reach it. The check-18 nodes sit on 7 EN maps (10 nodes) and 6
  RU maps (9 nodes). The difference is `100.alm`, whose EN copy carries a
  check-18 node and whose RU copy does not, an independent reproduction of
  `MISSION-ROOT-017`'s "EN … with 1 VIP node … against RU … with 0".
- `MISSION-VIP-004`'s "10 check-18 nodes over 8 of the 28 campaign maps" is
  corrected to 7 here.
- The maps with neither are `30`, `41`, `51`, `61`, `80`, `90`, `120`, `121`,
  `131`, `140`, `141`: 11 of 28, and the set is identical on both roots.
- Every one of the 28 authors exactly one instant-4 node, and every one is
  reached.

**Confidence.** High for the census: it is exhaustive over both roots, and a
second independent instrument, EXP-0155's `tools/trigclosure`, reports the same
authored and referenced counts. Medium for "the mission cannot be lost", which
is scoped to the map's own script. This experiment did not search for an
engine-level loss condition such as the loss of the whole party, and nothing it
establishes rules out `MISSION-END-013`'s counter being reached from outside the
script.

### MISSION-CURE-026

- `trig[7]`, the always-true start trigger, runs instant 12 with `Unit=10001`
  and `Item=6`. By `TRIG-ADDITEM-027` that creates class 14 index 30 from
  `ITEM-CODE-029`'s factory and adds it to the first hero's own container,
  taking nothing from anywhere.
- The win trigger's action list then runs two instant-13 nodes,
  `resolve and destroy item` (`TRIG-ACT-004`), with the same code, against unit
  10001 and unit 10002. 10002 is the second hero ordinal, which exists because
  `[Mission30] AddHero=22` inserts a companion before the mission starts. The
  pair covers both possible carriers.
- The item is not a condition of anything. The editor vocabulary contains check
  17, `Item in inventory`, and the campaign authors 23 of them; `30.alm` authors
  none. Its ten check nodes are four group-population tests, two constants, two
  point-to-unit distances and two unit-to-unit distances, and no node in the map
  reads an inventory.
- The map's own type-8 loot is four records at cells (28,56), (32,31), (39,64)
  and (36,42), none of them the script's item and none at the destination.

**Confidence.** High for the mechanism: both arms are published claims, the node
parameters resolve exactly, and the absence of an inventory test is an
exhaustive enumeration of ten check nodes rather than a search. Medium for what
the object is: `ITEM-CODE-029` records the three constructor arguments as
Unknown, so beyond class 14 index 30 the only evidence for the object's identity
is the map author's own label text, and `evidence/label-audit.txt` shows author
label text is unreliable in this corpus.

**Amended.** `10002` is a mage of the sex opposite the primary's, not a second roster position, and `10001` is the primary character (`TRIG-HEROORD-075`, `TRIG-HEROTPL-076`); "the second hero ordinal" and "the first hero" in this card are positional wording.

### MISSION-PAIR-027

- `scenario.res` holds 28 maps: fifteen numbered in tens from 10 to 150, and
  thirteen `N+1` maps. `10.alm` and `20.alm` have no partner.
- `scenario.reg` `[Mission30]` carries `ShopMission=31`, which `REG-SCN-064`
  reads as the shop building's offer list, and `[Mission31]` carries
  `Payment=700`.
- The two maps share no node: a record-by-record and node-by-node comparison
  differs at every compared position (`evidence/diff-30-31.txt`). `31.alm`'s
  type-4 record is empty where `30.alm`'s is 560 bytes, and `31.alm` carries its
  own single instant-4 node, its own trigger-run instant-5 node and three
  check-18 nodes.
- Its win is a three-step rescue: three `check 15` (`TRIG-NEAREST-029`) arrivals
  at cells (17,52), (27,43) and (66,37), each handing one placed unit to player
  1 with instant 19 and incrementing mission variable 50 with instant 8, then a
  trigger on `variable 50 == 3`. Its lose is a check-18 trigger over the same
  three units.
- `main.res` ships `text/battle/m30/event{04,07,08,09,13,15,16,17}.txt` and
  `text/battle/m31/event{01..06}.txt`, matching each map's authored instant-2
  numbers exactly, including the two `30.alm` nodes no trigger reaches.

**Confidence.** Medium. The registry key, the map counts and the two scripts are
exact and exhaustive, but the reading of what the pairing means rests on
`REG-SCN-064`'s offer-list result rather than on anything this experiment
discriminates, and only 5 of the 13 side maps are named by a `ShopMission` or
`TCMission` key at all.

### MISSION-TYP-028

- Typical: exactly one instant-4 node and exactly one drop location, like all
  28; a `once=1` win trigger, like all 28; message-per-objective as the dominant
  action, 8 of its 15 instants; group-population and distance checks as its
  whole condition vocabulary.
- The map's entire executable surface is five instant arms and four check arms:
  instants 2, 4, 12, 13 and the build-time `0x10002`; checks 1, 6, 7 and the
  build-time `0x10002`. `TRIG-CLOSURE-037`'s matrix already grades every one of
  them `full`.
- Exceptional: it is one of 11 of 28 maps with no script-authored lose path
  (`MISSION-LOSE-025`), and one of only 5 maps authoring an instant-12 item
  creation.
- Corrected: `MISSION-TYP-010` says an instant-5 node is authored on 14 of 28
  maps. It is 15, identically on both roots, and two independent walks agree.

**Confidence.** High for the corrected count: an exhaustive census reproduced by
two instruments that share no code. Medium for the rest: corpus agreement over
an exhaustive walk of both corpora, which is what "typical" can be measured
against and no more.

## Success panel and Continue

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-VICTORY-029 | The ordinary campaign-success panel is a 384×180 compiled surface at `(128,200)-(512,380)` with two 288×24 buttons, Victory above Continue. | High | ● active | [EXP-0233](../experiments/EXP-0233-victory-continue/) |
| MISSION-VICTORY-030 | Victory and Continue are stored results, not state writers: Victory queues `0x41d` and tears the mission down, and Continue queues `0x446` without touching `campaign+0x3bc`. | High | ● active | [EXP-0233](../experiments/EXP-0233-victory-continue/) |
| MISSION-VICTORY-031 | The success panel's static input programme starts on Victory, but Escape is not a Continue click. | High / Medium | ● active | [EXP-0233](../experiments/EXP-0233-victory-continue/) |
| MISSION-VICTORY-032 | At construction, End Quest's first row is disabled exactly when `campaign phase == 2 && copied campaign+0x3bc == 0`. | High / Unknown | ● active | [EXP-0233](../experiments/EXP-0233-victory-continue/) |
| MISSION-VICTORY-033 | At construction, “Victory after Continue” is reachability, not a `completed && continued` predicate. | High / Unknown | ● active | [EXP-0233](../experiments/EXP-0233-victory-continue/) |
| MISSION-VICTORY-034 | Mission completion is restored through `Player+0x3c`, and mission entry, resumed entry included, explicitly resets the delayed-Victory latch. | High / Medium / Unknown | ● active | [EXP-0233](../experiments/EXP-0233-victory-continue/) |
| MISSION-VICTORY-035 | Failure and loose maps do not share the success Continue transition: failure offers Exit to Main Menu and Load Game, and loose maps author no success action. | High | ● active (partially retracted, amended) | [EXP-0233](../experiments/EXP-0233-victory-continue/), [EXP-0274](../experiments/EXP-0274-defeat-modes/) |
| MISSION-VICTORY-036 | The customisation boundary splits resource wording in `main.res` from compiled choices in `rom.exe` and from the traced load semantics of the delayed-Victory latch. | High / Unknown | ● active | [EXP-0233](../experiments/EXP-0233-victory-continue/) |

### MISSION-VICTORY-029

- The panel is 384×180 at `(128,200)-(512,380)`. Its two buttons are 288×24
  each: Victory at `(176,284)-(464,308)` and Continue at `(176,308)-(464,332)`.
- Its constructor is called once, from the `0x430` success arm, and its builder
  is the sole target of success-vtable slot `+0x78`.
- Victory is control 2, local dialogue string 43, result `0x41d`. Continue is
  control 3, local string 154, result `0x446`.
- The selected bytes read `Mission Completed`, `~Victory!`, `Continue` on EN and
  `Миссия выполнена`, `~Победа!`, `Продолжить` on RU.
- Both roots carry the same executable hash, so geometry and control programme
  are common.

**Confidence.** High for the constructed object, rectangles, order and selected
resource bytes: two distinct controls refute the one-control/shortcut rival, the
sole constructor and vtable populations are enumerated, and both resource roots
are read. Exact rendered pixels remain unwitnessed and are not part of the
claim.

### MISSION-VICTORY-030

- Button activation posts `0x47a` with `button+0x70`. Success handler
  `R1299` delegates to `R1300`, which closes the panel with
  `0x445` and queues the non-zero result.
- Victory therefore queues `0x41d`, reaches the already-published win-latch test
  and tears the mission down.
- Continue queues `0x446`, closes the panel and does not touch `campaign+0x3bc`.
- Generic close builds `0x44c` with the panel pointer for teardown; it does not
  mutate campaign outcome state.

**Confidence.** High. The two stored constants, the delegate, the shared handler
and the generic close arm are read whole, and neither handler appears among the
nine exact direct-displacement hits for `campaign+0x3bc`.

### MISSION-VICTORY-031

- Base controls start with flags 1, buttons add bit 1, and panel show walks
  children forward requiring both bits. The caption is skipped and Victory, the
  first button, is focused.
- A focused button accepts Return or its accelerator and posts its stored
  result.
- The tilde makes Victory's stored accelerator `V` on EN and the CP866 byte for
  `П` on RU. Continue has no tilde and keeps explicit Latin `C`.
- Escape invokes the panel handler directly with `0x446`, which takes generic
  close without queuing the Continue button's result. It leaves the same latch
  value by a different dispatch trace.

**Confidence.** High for child order, flag tests, stored accelerators and the
Escape branch. Medium that real Return and accelerator delivery produce those
actions: the generic focus/key chain is complete, but no running-original
keyboard result was witnessed.

### MISSION-VICTORY-032

- Phases 0, 1 and 3 construct Change Map; phase 2 constructs Victory; all have
  result `0x41d`.
- Immediately after adding that row, `L06700..L06701` calls its
  `vt+0x1c(1,0)` when the copied latch is zero and phase is 2; that virtual
  clears enabled bit 0.
- All three direct callers create a fresh confirmation and copy the current
  latch into it.
- Disabled paint applies the generic shade-level-3 overlay over the full control
  rectangle.

**Confidence.** High for the constructor-time predicate and the paint operation:
the bounded predicate arm, all three direct-call inputs, the setter and the
disabled paint arm are read.

**Unknown.** This establishes the initial row state only. Whether later code
re-enables the row: post-construction enable mutators were not enumerated. The
exact visible pixels: the disabled paint's palette pixels were not observed.

### MISSION-VICTORY-033

- Both success branches set `campaign+0x3bc = 1` before either immediate `0x41d`
  or panel construction. Continue and Escape leave it 1.
- Immediate Victory makes a later menu unreachable by tearing down. Dismissal
  leaves the completed mission open, so a later fresh End Quest constructor
  receives 1 and does not take its disable arm.
- There is no Continue bit in the bounded handlers or the constructor predicate.
  The smaller model, in which completion sets the constructor input and
  dismissal preserves reachability, explains every observed pre-save transition.

**Confidence.** High for the ordering, the bounded choice chain and the
constructor input.

**Unknown.** Later post-construction presentation, because its enable-mutator
population was not enumerated.

### MISSION-VICTORY-034

- `Player+0x3c` is stored at `L06702` and loaded at `L06703`, and the
  preserved corpus contains three value-1 saves.
- The exact nine-hit direct-displacement census for `campaign+0x3bc` contains no
  serializer access. It does not cover aliases, address arithmetic, whole-object
  copies or a different serialized field.
- `R0098` resets the latch at `L06704`, and resumed mission entry
  resets it at `L06705`.
- The load-game path has already set campaign phase 2 at `L06706`, so the
  traced End Quest constructor receives latch 0 and initially disables Victory
  even though the Player record still says complete.

**Confidence.** High for the Player field, the explicit resets, the load
ordering and the resulting constructor input. Medium for the player-visible
save-after-Continue presentation, because it was not witnessed and
post-construction enable mutators were not enumerated.

**Unknown.** Whether another serialized field represents Continue or delayed
permission.

### MISSION-VICTORY-035

- Failure selects `Mission Failed` / `Миссия провалена` and constructs Exit to
  Main Menu (`0x445`) plus Load Game (`0x446`), not Continue or Restart.
  `MISSION-DEFEAT-046` fixes both localized resource indices, the actual class
  handler and the close route.
- The loose-map negative control remains: `MISSION-WIN-003` finds no authored
  success action on all ten.
- Success announcement/Continue need not clear the server run dword, but the
  shown campaign panel blocks idle stepping via bit 8 (`SESS-DEFEAT-065`).
- Success Continue remains `MISSION-VICTORY-029` (control and result `0x446`),
  `MISSION-VICTORY-030` (handler), `MISSION-VICTORY-031` (input programme) and
  `MISSION-VICTORY-033` (the reachability it leaves).

**Confidence.** High for the named instruction paths and the discriminating
branch/vtable counterexamples. No original runtime timing or GUI observation is
claimed.

**Amended.** EXP-0274 refutes the failure label, handler and continued-stepping
clauses ([`retracted.md`](retracted.md)): the former Restart label, the
failure-handler `0x41d` hop, and unqualified continued stepping of the completed
world. The first label is Exit to Main Menu in EN and RU, the actual handler is
`L03304`, and campaign outcome panels pause idle stepping. The distinct
success Continue and the prior loose-map census stand. The success
cross-reference read "Success Continue remains `MISSION-VICTORY-034`".
`MISSION-VICTORY-034` is mission-completion restore and the latch reset;
Continue's control, result and handler are EXP-0233's
`MISSION-VICTORY-029`…`031` and `MISSION-VICTORY-033`.

### MISSION-VICTORY-036

- `main.res` supplies the EN/RU title and button bytes.
- `rom.exe` fixes the two-button count, order, rectangles, results, focus rules,
  constructor predicate and explicit load-time latch reset; no campaign map
  authors them.
- `Player+0x3c` carries the three-valued mission result, while the traced load
  path does not restore delayed permission into `campaign+0x3bc`. A consumer
  that restores that latch before the traced constructor changes this load
  behaviour and needs an explicit derivation or representation.

**Confidence.** High for the resource selections, compiled operands, Player
serializer and explicit reset behaviour.

**Unknown.** Whether another serialized field or a later route restores Continue
or delayed permission. The direct-displacement census cannot establish that
every original save lacks some other Continue or delayed-permission
representation or a later restoration route.

## Failure-panel exits and the LOAD entry

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-DEFEAT-059 | Cancelling the save-selection dialog opened from the failure panel behaves as Exit to Main Menu, because the failure arm leaves `frame+0x414 = 0xff` and the dialog's close arm posts `0x41e` for it. | High / Medium | ● active (amended) | [EXP-0426](../experiments/EXP-0426-loss-load-music/) |
| MISSION-LOAD-060 | The `0x419` LOAD arm tears the mission screen down first only when the mask equals 1; from a zero mask it skips the teardown and still sends `0x445` to `frame+0x100` and runs the load path. | High / Medium | ● active | [EXP-0426](../experiments/EXP-0426-loss-load-music/) |
| MISSION-LOAD-061 | When the load path's join call `R0912` returns zero, it posts `0x421`, which reaches the menu surface only for a zero mask; nothing in the load path itself stops the music on that exit. | High / Medium | ● active (amended) | [EXP-0426](../experiments/EXP-0426-loss-load-music/) |

### MISSION-DEFEAT-059

- The `0x433` arm at `L03537` stores its wParam in `frame+0x414`; the failure message
  carries `0xff` (`MISSION-DEFEAT-046`).
- The close arm of the dialog at `frame+0x128` clears mask bit 8, then branches on the
  dialog's result `+0x60`. Result `0x445` posts `0x419`. Any other result posts `0x421`
  when the session pointer `L00285` is zero, and `0x41e` when it is non-zero and
  `frame+0x414 == 0xff`; a non-zero session with another `frame+0x414` posts nothing.
- `0x41e` is the Exit arm: teardown, then `0x421` (`VIDEO-MUSIC-062`).

**Confidence.** High for the close arm's three branches. Medium that nothing rewrites
`frame+0x414` between the failure arm and the dialog's close, which was not swept.

**Unknown.** Which input produces a result other than `0x445` in the dialog (Cancel, Escape
or a window close).

**Amended.** The Unknown clause is answered by `MISSION-067`: Cancel, Escape and Load with no selection return `0x446`; a window close is not a dialog result.

### MISSION-LOAD-060

- Arm `L06707` calls `R1283` when `frame+0x3dc == 1`. `R1283` calls session teardown
  `R0826` when the phase is 1 or 2, the session exists and mask bit 0 is set, calls
  `R1301(1)` when `L06708` is non-zero and the phase is not 3, and calls `R1302` when
  bit 0 is set. `R1302` clears bit 0.
- The arm then calls `R1303`. For mask zero, which a taken teardown leaves, it calls
  `R1301(1)` and sends command `0x445` to the object at `frame+0x100`. It ends with
  `R1284`.
- From the main menu the mask is expected to be zero, so the teardown is skipped. From the
  mission screen or the failure flow it is expected to be 1 (`VIDEO-MUSIC-063`).
- `SAV-972` states that selected LOAD completion requests `0x419 -> R1284`; this claim adds
  the mask-dependent teardown in front of it.

**Confidence.** High for the arm's branches and `R1283`'s guards. Medium for the mask
values, which are inferred from `MISSION-STOP-016` and `VIDEO-MUSIC-062`, not read at runtime.

**Unknown.** The mask at LOAD from the in-game Esc menu.

### MISSION-LOAD-061

- `R1284` calls the join routine `R0912` at `L06709` after `R1304(2)` and the
  document call `R1305`. A zero result sets the message to
  `0x421` and posts it at `L06710`; a non-zero result branches on `InBattle`
  (`VIDEO-MUSIC-064`, `VIDEO-MUSIC-065`).
- The `0x421` arm calls `R0816` only when the mask is zero, so a failed join from the
  main menu requests the menu list and a failed join with a non-zero mask requests nothing.
- No stop or replace precedes the post inside `R1284`.

**Confidence.** High for the failure branch and its message. Medium for the mask clause.

**Unknown.** What a zero return means for the session object left behind, and whether the
original shows an error before the menu.

**Amended.** The Unknown clause on the zero join `R0912` stands: `R0099` is called at `L06711` only after the join succeeds. The amendment adds the later zero return of `R0099` in the LOAD path, which `MISSION-068` covers.

## Minimap paint, structure selection and cursor surfaces

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-063 | `[view+0x9bc]` in the minimap paint is the bucket array of the map view's object map; the paint calls every node's slot `+0x34`: structures, bridges, units and air units draw a one-colour rectangle, three classes nothing. | High / Medium / Unknown | ● active | [EXP-0455](../experiments/EXP-0455-mission-screen/) |
| MISSION-064 | The minimap's masked pixel loop works per fog block: four tile words with no bit in `0xc000` replace the block from a second pixel buffer, an OR of `0x8000` averages it with that buffer, any other value leaves it. | High / Medium / Unknown | ● active | [EXP-0455](../experiments/EXP-0455-mission-screen/) |
| MISSION-065 | A structure with `Indestructible` clear is selected by a plain click, never by a rectangle, through the ordinary `+0x7c` flag write (`R1306`); hit-mask bit `0x20` is read at four sites, all in `R0219`. | High / Medium / Unknown | ● active | [EXP-0455](../experiments/EXP-0455-mission-screen/) |
| MISSION-066 | No direct load applies cursor slot 25 `dice`; its one non-registry read is a default-or-`dice` guard, slot `+0x2c` of the view at `campaign+0x358`. Slot 26 `wait` is loaded in 12 campaign view entries, 11 of which also load default. | High / Medium / Unknown | ● active | [EXP-0455](../experiments/EXP-0455-mission-screen/) |

### MISSION-063

- `R0310` (`evidence/disasm-minimap-paint-R0310.txt`) reads the map view through `[this+0x5c]` (`L06712`). Its displacements `+0x80` (tile plane), `+0x9bc`, `+0x9c0`, `+0x9c4` and `+0xa70` match the view layout of `AI-SELECT-065`, `AI-CURSOR-231` and `AI-CURSOR-202`. The array is therefore the bucket array of the object map `view+0x9b8` (count `+0x9c0`), not a field of the widget.
- The walk at `L01170`..`L01171` is the `GetNextAssoc` bucket scan of that map: for every node it calls `[[node+0xc]]+0x34` with a destination point from `this+0x60`/`+0x64`/`+0x68` and the scale `this+0x68`, then loops back (branch to `L06713`). `AI-MINIMAP-156` took the first non-empty entry only and left the array's identity open.
- Slot `+0x34` per class (`evidence/dwords-vtables-slot34.txt`, `evidence/disasm-getruntimeclass-L13155.txt`): `R1307` for `CStructure`, `CBridge` and both wooden bridges; `R1308` for `CUnit` and `CAirUnit`; `R1309` (a bare return popping 0xc argument bytes, no code) for `CGameObject`, `CBackPack` and `CProjectile`.
- `R1307` ORs the fog words of the four cells at the object's position (object fields `+8` and `+0xc`, each shifted right by 8) and draws only when the OR masked with `0xc000` equals `0xc000`; `R1308` adds two conditions, byte `+0x15a` at most 1 and `+0x18c` bit `0x80` clear. Both call `R0718(x0, y0, x1, y1, colour)`, a clipped solid fill of a 16-bit rectangle in the surface `[L01168]`. The colour is the word at `+0x1148` of the owner's player record, found through `[L05816][[obj+0x14]+8]` and `[record+8]`.
- After the walk one of two closing arms draws the viewport rectangle through the same fill with a white colour built from the channel masks (`L01173`..`L01174`, `L01175`..`L01160`).

**Confidence.** High for the walk, the three slot targets, the conditions and the fill arguments (instructions read). Medium that `this+0x5c` is the `CMapView` itself: five displacements agree with that view's layout, but the store of `this+0x5c` is not read here. The fog reading of `0xc000` as visible is `TERR-FOG-082`'s, Medium.

**Unknown.** The size of a blip in map cells (the extent arguments come from the object's slots `+0x20` and `+0x24`, not read); whether `+0x1148` is the player's chosen colour.


### MISSION-064

- Per fog block (`L06714`..`L01170`) the loop ORs `& 0xc000` of four neighbouring tile words of `[view+0x80]`'s plane (`L06715`..`L06716`). Zero: `REP MOVSW` copies the block's rows from the second buffer into the surface `[L01168]` (`L06717`..`L06718`). `0x8000`: each destination word becomes `((src >> 1) & m) + ((dst >> 1) & m)` with `m` built from the channel bit counts (`[ebp-0x54]`, `L06719`..`L06720`), a 50 per cent average per channel (`L06721`..`L06722`). Other values, including `0xc000`: untouched.
- The second buffer is the pixel data of the object at global `L01161`, `[obj+0x10] + 8` (`L06723`); the paint blits that same object first (`L01162`, slot `+0x18`).
- `TOWN-091` calls the unit a masked 2x2 block; it is a fog block of side `[ebp-0x14]` pixels.
- The routine returns at once when `[sess+0x3dc]` bit `0x2` is set (`L06724`) and repaints only when `[this+0x6c]` differs from `[view+0xa70]` by at least 11 (`L06725`..`L06726`).

**Confidence.** High for the three branches and the arithmetic. Medium for the fog reading: this loop shows that `0x8000` is a distinct third case, and the meaning of the three cases rests on `TERR-FOG-082`'s state mapping.

**Unknown.** Which picture the object `L01161` holds; whether the second buffer is a shroud, so that the three cases draw unseen, explored and visible cells.


### MISSION-065

- Plain click and rectangle: `AI-SELECT-122` (a plain click selects a structure only when its class `+0x64` is zero). The rectangle path skips every structure: `R0212` tests each candidate with `R0214` against `L06401` and a hit leaves its accept flag unset (`L06727`..`L06728`, `evidence/disasm-box-select-structure-skip-L13156.txt`).
- The select method for `CStructure`, `CBridge`, both wooden bridges, `CGameObject`, `CBackPack` and `CProjectile` is `R1306` (vtable `+0x14`; the seven vtable slots that hold it are listed by a whole-image dword search over the EN image): it stores its argument in `+0x7c` and `1` in `+0x10c`. `CUnit` and `CAirUnit` bind another routine. The selection rebuild counts the objects with `+0x7c` set (`AI-CURSOR-202`), so a selected structure gives `view+0x140 == 1` and `view+0x138` its object, which is what widget 8 reads (`MENU-070`).
- Bit `0x20` of the hover hit mask is set for every hit `CStructure` (`AI-CURSOR-231`). `R0218` has six direct callers, all in `R0219`. The sites that test bit `0x20`: `L06729` and `L06730` (mask test against 0x23; no selection or a selection with summary bits `0x24`) choose slot `select` `[L06217]`; `L01526` (mask test against 0x20, under mask bit `0x4` and a zero gate) chooses `select` when set, else slot `attack` `[L00628]`; `L01529` (mask test against 0x23) belongs to the same cascade without bit `0x4`.
- `R0211`'s selection and order arms read `view+0x98c` and `+0x990`, not the mask (`evidence/disasm-click-dispatch-town-L13157.txt`).

**Confidence.** High for the box skip, the select method, the six callers and the four test sites (listings read). Medium that no other routine receives the mask: the callers are direct `E8` calls and an indirect call is outside the search (`AI-CURSOR-231`'s bound).

**Unknown.** What mask bit `0x4` means (`AI-CURSOR-231` lists it without a name); the click path that calls `vt+0x14` for a single structure click was not read to its call.

### MISSION-066

- Slot 25 `dice`, `L06731` (`evidence/scan-cursor-slots.txt`, direct displacement loads over the EN image): three references. The registry store (`L06732`) and the registry teardown (`L06733`, a `vt+4` delete) belong to the registry. The third, `L06734` in `R1310`, loads the slot's `+4` and compares it with the current cursor `[L00619]`: if the current cursor is neither slot 0 `default` nor `dice`, it calls `R0320` with the default slot, then `R1311`.
- `R1310` is slot `+0x2c` of vtable `R1312` (`evidence/dwords-vtable-R1312.txt`), whose constructors are `R1313` and `R1314`. `R1314` is called at `L06735` in the campaign screen constructor with id `0x456` and stored in `campaign+0x358` (`L06736`, `evidence/disasm-view-construction-L13158.txt`); `TOWN-140` names it as the fourth view of the gate `0x226` and `TOWN-242` as the final or detailed character-generation view.
- Slot 26 `wait`, `L06737`: 14 references, the registry's two and 12 loads, in `R1315`, `R1316`, `L06738`, `R1165`, `R1317`, `R0816`, `R1318`, `R1319`, `R1320`, `R0099`, `R1321` and `R0909`. Each load is followed by `R0320` early in the routine and by the build of one campaign view; `R1315`, `R1318`, `R1319` and `R1321` set `[sess+0x3dc]` bits `0x2`, `0x4`, `0x20` and `0x200` (`TOWN-140`). Eleven of the 12 also load slot 0 `default` later in the same routine (the call after it was read in four, see Confidence) (`evidence/scan-cursor-slots.txt`); `R0816` does not.

**Confidence.** The clauses carry their own grades. `dice`: High for the reference census (three direct displacement loads in the EN image), the compare in `R1310` and the construction chain; Medium that slot `+0x2c` is called as the view's cursor hook, since the slot is read from the table and its caller was not searched; Unknown for the applier. `wait`: High that 12 routines load the slot and that 11 also load slot 0; the call that follows the slot-0 load (`R0320`) is read in four of them (`R1316`, `L06738`, `R1165`, `R1321`), so "restores default" is established for those four only. Medium that all 12 are view entries: the first lines of each were read, the bodies of `R0816`, `R1320`, `R0099` and `R0909` only partly.

**Unknown.** Which instruction applies the `dice` cursor: no direct slot load does, so the one reader is a guard, not an applier, and the five set sites with no static slot load (`AI-CURSOR-175`) are not excluded. Whether any entry draws a wait picture of its own: only the cursor slot was searched.

## Load cancel routes and mission-start wait exits

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-067 | The save-selection dialog returns `0x446` for Cancel, Escape and Load with no selection, and `0x445` only for Load with a selection; Delete and list selection leave it open; a window close is not a dialog result. | High / Medium / Unknown | ● active | [EXP-0459](../experiments/EXP-0459-start-exits/) |
| MISSION-068 | A zero return of `R0099` is dropped by the `0x42f`, `0x467` and LOAD callers; only the `0x457` arm shows an error panel and posts `0x45c`; none of the three clears mask 1 itself, `0x45c` does in phases 0 and 1. | High / Medium | ● active | [EXP-0459](../experiments/EXP-0459-start-exits/) |
| MISSION-069 | On a single-player start the wait loop's three zero exits need a quit message, 60,000 ms of silence or a failed map open; a type-`0x0b` error posts `0x45c`, which ends the loop through `WM_CLOSE` or leaves it to the timeout. | High / Medium / Unknown | ● active | [EXP-0459](../experiments/EXP-0459-start-exits/) |

### MISSION-067

- The dialog (vtable `L06375`) has three buttons: Load (id 4, message `0x478`), Cancel (id 5, `0x479`) and Delete (id 6, `0x474`). Its handler `R0737` sends `0x444` and `0x478` to the base with `0x446` when the selection `+0x80` is negative and with `0x445` when it is not. `0x479` sends `0x446`. List selection (`0x46d`, wParam 3) stores `+0x80`. `0x474` deletes the selected save after a confirm panel and the dialog stays.
- The base `R0716` handles only `0x445` and `0x446`. When the open flag `+0x5c` is set it stores the message as the result `+0x60`, hides the dialog and sends `0x44c` to the frame; the frame's `0x44c` arm calls the close handler `R0709` (`MISSION-DEFEAT-059`).
- The key slot `R0818` maps VK `0x1b` to `0x446` through the dialog's message slot. Other keys reach `R0717`, which handles only Tab and the four arrows; Enter is not mapped there.
- A window close does not reach the dialog. The frame's `WM_CLOSE` entry calls `R1322`: teardown `R1283`, mask set to zero, then the MFC base close. No result is stored and `0x41e` and `0x421` are not posted.
- The dialog is a game UI object shown with `R0361`, not an operating-system window, so it has no title-bar close of its own.

**Confidence.** High for the handler arms, the base and the close handler's use of the result. Medium for Escape: the key slot is read, the container dispatch that calls it was read only in part. Unknown for any other input (mouse on a non-button area, list-control messages, a held key).

**Unknown.** Whether another message reaches `0x446` through the container's child dispatch.

### MISSION-068

- `R0099` returns zero at `L06739` (timeout, after storing `0x1005` in `L06740`) and at `L06741` (pumped `0x12`; zero `R0509(0x64)`). The last two leave `L06740` unwritten.
- The `0x457` arm (`L06742`; call at `L03842`, argument 0; `L06740` zeroed by `R0821` before the call when the phase is not 3) on zero shows the error panel with row `L06740 & 0xff` (`R0736`, a modal message loop that ends on `0x44c`), then posts `0x45c`. The row is main.txt line `192 + index` for index 10 or lower and a row of the patch text object above that. Index 0 is a stale label, index 4 the missing current map file (`0x1004`), 5 the map-switch packet wait (`0x1005`), 7, 8 and 9 the player-list, character-login and server-response waits. EN and RU have 11 non-empty rows each, with different lengths and digests (`evidence/errrows.txt`).
- The `0x45c` arm posts `WM_CLOSE` when the CString at `frame+0x3b4` is non-empty. Otherwise phase 0 and 1 run `R1301(1)`, `R1302` when mask bit 0 is set and `R1303`, then post `0x454`, `0x455` or `0x452` by `L06743`; phase 3 posts `0x453` or `0x451`; phase 2 does nothing.
- The `0x42f` arm (`L06744`, argument 0) and the `0x467` arm for a non-zero wParam call `R0099` and ignore the result: no panel, no post. The LOAD path (`R1284`, call `L06711`) ignores it and goes on to `L06745`, which stores `L06746 = 1`.
- After the timeout and the zero-dispatcher exits the rebuilt mission containers (`R1149`) stay on screen with mask 1, `frame+0x3c8` keeps the start argument, `world+0x80` is zero or not as the exit left it, and nothing posts `WM_QUIT` again. A `0x12` exit from the frame close runs `R1322`, `R1283` and `R1302`, which clear mask bit 0 and tear down (`MISSION-067`).

**Confidence.** High for the exits' return paths, the three callers' branches and the row selection. Medium for the surface and state left, which are read from the code and were not observed, and for the row meanings (index 4, 7, 8, 9): they were read from the installed text in a private run, and `evidence/errrows.txt` holds lengths and digests only.

**Unknown.** What the idle stepper (`R0260`, `R0454`), which runs while mask bit 0 is set, does with `world+0x80` zero.

### MISSION-069

- The single-player start is phase 2 of `frame+0x6bc`. Its callers of `R0099` are the `0x42f` arm (after the `0x358`-dialog OK) and the LOAD path; both drop the zero (`MISSION-068`). The `0x457` arm tests only phase 3 before its call. Its five posters are `L06747`, `L06748`, `L06749`, `L06750` and `L06751`; the guards of `L06747` (the `0x459` command-line arm) and `L06749` (`R1323`) were read, the phase of the other three was not, so the exclusion of phase 2 from the panel route is Medium.
- Pumped `0x12` needs a quit message: the frame close, another `AfxPostQuitMessage` caller in the MFC code, the game's own exit routine `L06752`, or a `WM_CLOSE` posted by the `0x45c` arm when the CString at `frame+0x3b4` is non-empty.
- The 60,000 ms timeout is tested only when the receive count `L06753(L00625)` is zero. Nothing in the loop reads installed data to shorten or lengthen it.
- A dispatcher zero is a type-`0x06` message whose world load leaves `world+0x80` zero (`0x1004`); the map is the file named by the session's map name under the `scenario\` prefix in phase 2. A type-`0x0b` message stores its code with `0x1000` set. With mask bit 0 set (as `R1149` leaves it) it sends `0x446` to the container and posts `0x45c`, then the loop continues and dispatches `0x45c`: in phase 2 the arm posts `WM_CLOSE` when `frame+0x3b4` is non-empty, which leads to the `0x12` exit, and does nothing when it is empty, which leaves the loop to the timeout. With the mask clear the arm drains the queue and returns zero. Which branch a shipped start takes is not established.
- The `0x1004` condition and the `frame+0x3b4` writers were not checked against the shipped mission list of either install. The EN and RU executables are byte-identical (`evidence/program-digest.tsv`).

**Confidence.** High for the conditions each exit tests and for the `0x0b` arm's mask branch. Medium that no shipped-data condition was found feeding them: the search covered the loop's own reads and the type-`0x06` path; the dispatcher's closure has 574 functions, 162 of them with computed calls that were not read. Unknown at runtime, since no run was made.

**Unknown.** The writers of `frame+0x3b4`, which decide the `0x45c` branch, and the container's handling of `0x446` at `frame+0xcc`. The sender of a type-`0x0b` message (no immediate store of `0x0b` at `msg+9` was found), the arms of `R0509` other than types `0x06`, `0x0b`, `0x64`, `0xb4`, `0xb5` and `0xb6`, and whether a shipped mission's map file is absent on any install.

## Mission 70 opening report

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-REPORT-070 | In the preserved EN/RU 70.alm, event 1 has one incoming trigger: start compares condition 1 with itself using equality and has once=1; its authored pairs contain no combat or proximity condition. | High | ● active | [EXP-0471](../experiments/EXP-0471-m70-start-report/) |

### MISSION-REPORT-070

Both `scenario.res::70.alm` payloads are 131772 bytes with SHA256
`c29e635bc43716bd3551bbe8d8eae036b483a0515edfc9359b9e1f76c3a6dcae`.
The complete selected type-7 payload is 39080 bytes: 24 actions, 23 conditions
and 9 triggers, with no residue and unique IDs in each node list.

One action has opcode 2 and first plain operand 1: ID24, array index 20.
Exactly one of the nine triggers names it: index 1, label `start`. Its pairs
are `(1,1),(0,0),(0,0)`, comparison words `(0,0,0)`, action order
`25,24,27,0`, and once word 1. Condition ID1 has opcode `0x10002`, plain
value 10 and no authored object references. Actions 25 and 27 precede and follow
the report; their effects are not part of this claim.

`TRIG-BIND-010` retains the nonzero first pair and omits the two zero pairs.
`TRIG-COND-003` binds the variable condition to a slot. `TRIG-CMP-006`'s code 0
compares the same slot with itself, so this pair is true whenever evaluated.
`TRIG-FIRE-007` still requires the once latch to admit evaluation. The
authored predicate therefore needs no battle state; this does not establish
when a native script pass or a client dialogue occurs.

**Confidence.** High for this bounded authored graph and conditional equality.
The stdlib archive/script parser tiles every record of the selected payload
and enumerates all 24 actions and all 9 incoming-reference candidates. The
absence clause concerns only the retained pairs of this one trigger. It does
not enumerate executable writers or the virtual filesystem. Published binder,
comparison and latch claims supply the compiled semantics; no engine code
or inferred entry time supplies an input.

**Unknown.** Native first-seconds presentation, script-pass admission,
current or saved latch state, effects of actions 25/27, client admission and
delivery order, archive overrides and the owner's save lineage. A reached
native trigger evaluation and report-open observation are the next proof.

## Original mission survey: world-map objects and the troll drop

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-MAP-073 | `[MissionObjects] Mission<n>` in `globalmap.reg` names `MapObject<value>`; five main missions (50, 60, 70, 100, 110) have a picture overlay on that object, and missions 90, 120 and 140 have none. | High | ● active | [EXP-0478](../experiments/EXP-0478-mission-scripts/) |
| MISSION-DROP-074 | A dying unit leaves `p39 + U[0,p40]` gold when `U[0,100] < p38` of its own Data.bin Units row, chosen by class and `class2`; Troll rows 84..87 pay 200+300 up to 266200+399300 at p38 90; only `actor+0xe > 0x40` pays. | High / Medium | ● active | [EXP-0478](../experiments/EXP-0478-mission-scripts/) |

### MISSION-MAP-073

`R1324` returns the 1-based `Mission<n>` value minus 1 and `R1325` reads `MapObject<index+1>`, so mission n uses `MapObject<value>`. Values: 10 to 4, 20 to 20, 30 to 10, 40 to 21, 50 to 19, 60 to 18, 70 to 3, 80 to 22, 90 to 15, 100 to 13, 110 to 14, 120 to 17, 130 to 11, 140 to 2, 150 to 2, 41 to 23, 51 to 24, 61 to 25, 71 to 35, 81 to 27, 91 to 28, 101 to 29, 111 to 30, 121 to 31, 141 to 32, 151 to 33, 31 to 34, 131 to 26. The five picture objects 3, 13, 14, 18, 19 are the targets of missions 70, 100, 110, 60 and 50, with the overlays `onmap03`, `onmap13`, `onmap14`, `onmap18`, `onmap19` (640x480, transparent). `GMap.bmp` is static and carries every named town; mission 70 adds the monastery picture, not a label. MapObject 2 serves both 140 and 150.

**Confidence.** **High** (decoded registry, two locales, decompiled value routine and marker builder).

### MISSION-DROP-074

Population: Data.bin Units rows 72, 74, 80, 82, 84..87 on EN, rows 84 and 86 on RU (`mission0478 -mode row`). `R0208` (actor death, entered for `actor+0xe > 0x40`, so class 64 Goblin never pays whatever its row holds) draws `R0861(100)` and, when it is below parameter 0x26, sets gold to parameter 0x27 plus `R0861(parameter 0x28)`. `R0861(n)` is `(rand*(n+1))/32768`, so inclusive. The comparison is strict, so the chance is `p38/101` for `p38` of 90. The Units collection has four rows per monster, keyed on parameters 0x1d (class) and 0x1e (tier) (`MISSION-ARM-006`); the type-6 `class2` field is the tier. The actor's own row is read (`HERO-KILL-027`).

Parameter indices are those of `mission0478 -mode row`; its column labels are not relied on.

| Row | Name | p0x1d | p0x1e | p38 | p39 | p40 | Gold |
|---|---|---|---|---|---|---|---|
| 84 | Troll | 68 | 1 | 90 | 200 | 300 | 200 to 500 |
| 85 | Troll.2 | 68 | 2 | 90 | 2200 | 3300 | 2200 to 5500 |
| 86 | Troll.3 | 68 | 3 | 90 | 24200 | 36300 | 24200 to 60500 |
| 87 | Troll.4 | 68 | 4 | 90 | 266200 | 399300 | 266200 to 665500 |
| 80 | Ogre | 66 | 1 | 90 | 50 | 250 | 50 to 300 |
| 82 | Ogre.3 | 66 | 3 | 90 | 6050 | 30250 | 6050 to 36300 |
| 72 | Orc_Sword | 80 | 1 | 80 | 30 | 90 | 30 to 120 |
| 74 | Orc_Sword.3 | 80 | 3 | 80 | 3630 | 10890 | 3630 to 14520 |

Mission 111 places the Troll (unit 119, class 68, `class2` 3), the Ogre (unit 53, class 66, `class2` 3) and the Orc at (87,36) (class 80, `class2` 3), so each uses its tier 3 row: unit 119 draws `24200 + U[0,36300]` with chance 90/101 and nothing with chance 11/101. The gold goes through `R0871` into the corpse sack. Items come from the actor's inventory, not these columns (`ITEM-DEATH-012`); the magic treasure columns 41..43 are not read in this routine.

**Confidence.** **High** for the draw, compare, range and the tier rows dumped (84 and 86 on both roots, the others on EN). **Medium** that rows 85 and 87 feed the draw identically (the draw reads the actor's own row; only rows 84 and 86 were dumped on both roots) and that the item and magic columns are unused (not searched outside the routine).

## Mission 100: the quest amulet

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MISSION-M100-079 | In 100.alm the amulet 0x0e23 starts as loot of unit 146; T10 takes it from 10001 at (78,132), the win pass from roles 10001..10005; a holder no role names keeps it, a kept one through the culls; town entry unsearched. | High / Medium / Unknown | ● active | [EXP-0511](../experiments/EXP-0511-m100-servant-amulet/) |

### MISSION-M100-079

`100.alm`, EN and RU. Every node is opcode 12 or 13 with item value 11, which is code `0x0e23` (`0xe18 + 11`), class 14, `MagicItems` row 35.

- Placement: unit 146 (owner 6 "Enemies", group 24, (120,112)) carries `[0x0e14, 0x0e23]` in its type-6 loot list. Action 22 (opcode 12, add the amulet to 10001) is in no trigger slot.
- T10 (once): the nearest Self unit is at most 3 from (78,132) (C25 <= C18), and C37, the item check for item 11 on 10001, equals TRUE. It runs message 7, action 21 (opcode 13 on 10001), and hands groups 15 and 20 to player 1.
- T3 (once): C14 == C7 (FALSE) and C43 (check opcode 4, "D - 4th health", on unit 245) > C1 (0). It runs message 8, unit 245 to player 1, Mission Complete (instant 9) and action 36 (opcode 13 on 10001). T16 has the same conditions and runs opcode 13 on 10002, 10003, 10004 and 10005 in the same pass (`TRIG-M100-097`).
- Opcode 13 reaches only the pack (`ITEM-169`). An unresolved role's slot runs action 2 instead (`TRIG-M100-096`).

At the end the amulet's fate follows its holder:

1. A holder the matching role resolves to loses one amulet from its pack. **High** that opcode 13 takes from that holder's pack; **Medium** that the amulet is there (`ITEM-169`'s container clause, and the amulet not consumed by use).
2. A mercenary or other holder no role resolves to keeps it in the mission. The culls destroy that holder with its pack unless it is a kept player character (`PARTY-CULL-004`, `PARTY-ENDCULL-026`). A kept player character holds the amulet through both culls (**Medium**, `ITEM-170`'s closure). Whether it reaches town with it is **Unknown**: the arm's continuation at `L08133`, the `0x428` handler and town entry were not searched (`ITEM-170`).
3. An amulet in the cast slot of a kept actor returns to the pack only under the conditions of `PARTY-M100-034` (castSpell effect, item byte `+0x44` not 2, state not 0xd or 0xe with `+0x136` = 0, matching spell id); otherwise it is deleted at the reset when `+0x136` is 0.
4. An amulet on the ground is not carried. This rests on the culls emptying the document's other objects (`PARTY-CULL-004`).

**Confidence.** **High** for the authored placement, triggers and nodes (decoded whole, both roots; the comparison codes read from the pass's table, `compare-table.tsv`). **Medium** for outcomes 2 to 4 within the culls, which depend on `ITEM-170`'s closure and on which roles resolve. **Unknown** for the amulet's fate after the culls.

**Unknown.** Which hero a role resolves to in a given roster (`TRIG-HEROORD-075`). No original save of mission 100 exists in the searched corpus of 134 files, so no outcome was observed.
