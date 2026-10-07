# DIALOGUE — the window a mission's script speaks through, and its lifecycle

Claims about the presentation side of `claims/mission.md` and
`claims/trigger.md`. Those say what the script decides and which resource its
decision names; these say what object draws it, everything that can enter that
path, what happens when the resource is absent, what ends the display and what
the window measures. The glyph rule, how a text byte becomes a pixel, and the
font atlases are EXP-0097's area; this ledger cites it rather than restating it.
Spec: [`formats/dialogue/format.md`](../formats/dialogue/format.md). Format of
this file: [registry.md](registry.md).

## Terms

- A root is the install a measurement was read from; every measurement names
  its root. `rom.exe` is byte-identical in the EN and RU roots
  (`sha256 942e9b72…d367d03`, 1 977 344 B), so every address is a fact about
  both, and only resources can differ by language.

## Window class, announcement path and lifecycle

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DLG-WIN-001 | One `0x84`-byte panel class with one constructor draws every line of authored dialogue, and its six calling routines are the complete set of entries into it. | High | ● active (amended) | [EXP-0098](../experiments/EXP-0098-mission-text/) |
| DLG-PATH-002 | Script and outcome announcements share one transport; the failure close is selected by its own stored panel pointer. | High | ● active (partially retracted, amended) | [EXP-0098](../experiments/EXP-0098-mission-text/), [EXP-0101](../experiments/EXP-0101-mission-end/), [EXP-0274](../experiments/EXP-0274-defeat-modes/) |
| DLG-ABSENT-003 | A named event text that does not ship produces nothing: no window, no fallback, no fault, no blocked state; this answers `MISSION-TEXT-005`'s Unknown. | High | ● active | [EXP-0098](../experiments/EXP-0098-mission-text/) |
| DLG-EMPTY-004 | The window's `"Nothing to say"` fallback is reached only by an existing file that yields no part 1, and no shipped event file in the measured corpus reaches it. | High | ● active (amended) | [EXP-0098](../experiments/EXP-0098-mission-text/), [EXP-0417](../experiments/EXP-0417-offer-first-part/) |
| DLG-LIFE-005 | Only input ends the display, nothing queues, and an announcement that arrives while any dialog is open is discarded. | High / Medium | ● active (amended, partially retracted) | [EXP-0098](../experiments/EXP-0098-mission-text/), [EXP-0108](../experiments/EXP-0108-panel-modality/) |
| DLG-READ-006 | The event text is read at fire time, not at map load, and it is read twice per announcement. | High | ● active | [EXP-0098](../experiments/EXP-0098-mission-text/) |

### DLG-WIN-001

- `R0695(name)` allocates `0x84` bytes (`L03299`, the size argument 0x84) and
  constructs the panel with `R0696(id=9, 30, 120, 610, 360)`. Rects are
  `{left,top,right,bottom}`, fixed by `R0697` returning the field `+0xc` of
  one rect less the field `+0x4` of the other, `bottom − top`.
- It then adds:
  - a portrait (`R0698`, id 12, 30,54–118,168), only when `panel+0x7c`;
  - a text control (`R0699`, id 10, at 128,36–428,172 with the portrait
    and 48,36–428,172 without, its base constructor shrinking the height to
    135, `DLG-RECT-037`), whose initial content is the literal
    `"Nothing to say"` at `L03300`/`L03301`;
  - always a button (`R0700`, id 11, 200,172–280,198) whose label is
    string-table entry 77 and whose command is `0x46f`.
- Then `R0361` shows it.
- `EnumRefs callto:R0696` returns 1 hit, 1 owner, 0 in orphan code:
  `R0695` is the class's only constructor.
- `EnumRefs callto:R0695` returns 7 hits, 6 owners, 0 in orphan: the mission
  event text (`R0701`), the inn's NPCs (`R0702`, ×2), the
  mercenary hall (`R0703`, `R0704`), the shop keeper
  (`R0705`) and the training hall (`R0706`).
- Each of the six formats a name under a `main.res` text prefix. The two
  mercenary-hall owners share `text/inn/mercenary/`, so the six owners name
  five families, and all five ship (`evidence/family-en.txt`: 318/22/14/4/33
  nodes).
- The child rects are panel-relative, and the drawn layout settles it. Read as
  screen coordinates, the portrait's top edge (54) would sit 66 px above the
  panel's own (120). Read as panel-relative, all four children fall inside the
  580×240 constructor rectangle and inside the 488×232 rectangle drawn
  (`DLG-PANEL-035`) with nothing overhanging: 118 < 488, 428 < 488, 172 < 232,
  198 < 232. Four independent literals would have to coincide for that.

**Confidence.** High for the geometry, the two layouts, the button command and
the string-table index: each is a named `PUSH` in one routine read whole, and an
instruction, not a fit, settles the rect convention. High for the entry
enumeration with its instrument stated: `callto:` on the repaired table, which
also scans `.rdata` dwords holding the address, 0 orphan on both sweeps. Its
blind spot is a construction reached only through a vtable slot; the class
census `EnumRefs range:R1228:L13136` (13 functions, 1468 function bytes,
0 orphan bytes, 0 orphan runs) closes that hole for this class.

**Amended.** The family bullet read "Each of the six formats a name in its own
family, and all six families ship", beside five node counts.
`evidence/family-en.txt` gives `R0703` (`inn\mercenary\npc%02d`) and
`R0704` (`inn\mercenary\npc35`) the same prefix,
`text/inn/mercenary/` (14 nodes): six owners, five families. The entry
enumeration and the node counts stand.

A second correction (EXP-0410): the overhang bullet read "all four children
fall inside the 580×240 panel with nothing overhanging: 118 < 580, 428 < 580,
172 < 240, 198 < 240", and the text control bullet gave no height for the live
rectangle. 580×240 is the constructor rectangle, whose size `R0707`
replaces with 488×232 when the panel is drawn (`DLG-PANEL-035`); the argument
holds on both sizes. The text control's base constructor shrinks its height to
135 (`DLG-RECT-037`). The child literals and the entry enumeration stand.

### DLG-PATH-002

- `R0188` stamps static packet `L02444` with `0xb6`; `R0623` stamps the
  supplied outcome opcode; both use `R0217`.
- Client `R0509` selects `opcode-3` through index `L02523[188]` and table
  `L02524[46]`: `0xb6 -> 0x433` with its number, `0xb4 -> 0x433/0xff`,
  `0xb5 -> 0x430`.
- The failure sentinel routes to `0x431`: constructor `L03302`, panel stored
  at `frontend+0x110`, title 141. The vtable slot `L03303` contains
  `L03304`, not `R0708`.
- The base close stores the result and posts `0x44c`. `R0709` matches
  `+0x110` and maps `0x445 -> 0x41e` (teardown/menu), otherwise `0x418` (load
  selection).
- The separate win programme may post `0x41d` under its fast-transition
  predicate or build its own panel.

**Confidence.** High for the named instruction paths and the discriminating
branch and vtable counterexamples. No original runtime timing or GUI
observation is claimed.

**Amended.** `MISSION-DEFEAT-046` (EXP-0274) refutes the old dialogue-sentinel
close attribution and the later added neighbouring-class `0x41d` hop.
`retracted.md` records the withdrawn text: the failure close tested
`frontend+0x414 == 0xff`, and its class handler `R0708` posted `0x41d`
upstream. The actual failure close matches stored panel pointer `+0x110` and
result `panel+0x60`, and its handler is `L03304`. The shared transport
stands.

### DLG-ABSENT-003

- In `R0701`'s `0x433` arm the path is built, the reader is constructed
  on the stack, and `L03305` calls R0535 with both optional
  arguments zero (`L03306` supplies the same zero twice): flags 0 and no
  error-out pointer, so the routine's own error code 2 is discarded.
- `R0535` returns 1 only when `R0536(name)` resolves and 0
  otherwise. `L03307`/`L03308` compare the return value with zero and jump to `L03309` when equal, which skips the
  call to the window builder, and the arm returns 1 as if it had worked.
- The only durable effect is `campaign+0x414`, written before the test at
  `L03310`.
- Corpus, EN root: of 242 numbers the 28 campaign scripts raise, 223 name a
  shipped file and 19 are silent. RU raises the identical 242 (0 differing
  maps) and silences 16 (`evidence/join.txt`).
- This reproduces `MISSION-TEXT-005`'s 223/242 from two independent
  instruments, the raised set from `scenario.res` and the shipped set from
  `main.res`, and makes the figure a behavioural statement rather than a
  discrepancy.

**Confidence.** High. The guard, its two zero arguments and the skip are named
instructions in one routine read whole. The alternative readings, a fallback
resource, an empty window or a logged failure, each predict an instruction that
is not there, and the routine has no other exit between the open and the
builder.

### DLG-EMPTY-004

- The window's "on show" slot `vt+0x80` is `R0710`, which calls the
  pager `R0711` once and discards its return.
- The pager returns 0 when `R0712` cannot find a tag containing
  `part=<n>`. On that path nothing sets the text control, so the window opens
  on the constructor's own literal `"Nothing to say"`. The next pager call
  increments to part 2. It closes through command `0x46f` only if that lookup
  fails; an accepted part 2 replaces the literal and continues.
- Corpus: 0 of 225 EN and 0 of 228 RU event files lack a `part=1` tag, and 0
  files on either root have a gap in their part sequence before their own
  maximum (`evidence/markup-en.txt`, `markup-ru.txt`).
- The string is unreachable from shipped mission data, so a consumer that omits
  it diverges on nothing the original shows. It is reachable from an authored
  file, which is why it is published.

**Confidence.** High for the mechanism: the ignored return and the two literals
are named instructions, and the pager's zero return is its own `if`. High for
the corpus half: exhaustive over every event file of both roots, and the
measurement could have failed, since one file with a part-1 typo would show it.

**Amended.** DIALOGUE-049 corrects the former automatic button-close clause:
failure of part 1 does not preclude part 2. The corpus clause names the measured
event files, not every text family. The default text and ignored show return
stand. Native presentation of an authored missing-first-part file is Unknown.

### DLG-LIFE-005

- Three inputs reach the same command `0x46f`: a click on the button, Enter
  (`0x0d`) and Escape (`0x1b`). The button's own key handler `R0713`
  answers Enter by posting `0x46f`; the panel's key slot `vt+0x6c`
  `R0714` answers Escape by self-sending it, and has the same arm for
  Enter, reached only when no child takes the key (`DLG-KEYS-040`).
- `R0715` answers it by calling the pager. Non-zero: redraw children 10
  and 12 and stay open. Zero: post `0x445`, which `R0716` turns into
  `vt+0x84` (free the text buffer) plus `0x44c`, whose window arm is
  `L03311` calls R0709.
- There is no timer: the class census over `R1228..L13136` finds 13
  functions, 0 orphan bytes, and none references `0x113`.
- One instruction settles overlap. `R0361` sets bit 3 of
  `campaign+0x3dc` when it shows a panel, and the `0x433` arm tests
  bit 0x8 of the byte at `+0x3dc` (`L03312`) and returns without building
  anything when it is set.
- Announcements are therefore dropped, not deferred: a consumer that queues
  them shows text the original never showed. The drop at `L03312` does not
  depend on where the clears of the bit live.

**Confidence.** High for the three inputs, the pager loop, the absence of a
timer and the drop: each is a named instruction. Medium for "no other route
closes it": the container dispatcher `R0390` and the base key handler
`R0717` were read and add none, but the panel's key-release and
system-key slots and the senders of `0x445` outside the class were not
(`DLG-KEYS-040`). Medium for which handler answers Enter, which rests on the
dispatcher's child order.

**Amended.** EXP-0108 refutes the clear enumeration in both halves, and
`retracted.md` withdraws it. The withdrawn sentence: "The bit is cleared by
nine instructions, all nine inside `R0709`" (`EnumRefs imm:fffffff7`:
19 hits / 10 owners / 0 in orphan), with the grade "the gate's set and its nine
clears are enumerated with the instrument and its orphan count stated".
`imm:fffffff7` is a dword-immediate sweep and is blind by construction to the
byte-width form of the same operation. It missed `L03313` and
`L03314`, each masking the byte with 0xf7, inside `R0709`, so the count is at least twelve,
and `L03315`, a byte mask with 0xf7 inside `R0701`, between
`L03316` (a load of the 32-bit field `+0x3dc`) and `L03317` (the store back to `+0x3dc`), so not all
clears are in `R0709`. That arm also sets the bit itself
(`L03318` sets bit 0x8) without going through `R0361`: a shutdown
bracket that wipes the screen through `R0718` at level 0, calls
`R0719` and posts `WM_CLOSE`. The miss is what `AGENTS.md`'s
enumeration rule 4 guards against: reduce by the field's own width and read
every hit of byte or 16-bit width. The three inputs, the absent timer, the drop at
`L03312` and "dropped, not deferred" stand.

A second correction (EXP-0410): bullet 1 read "through `vt+0x6c`
`R0714`, the keys `0x0d` (RETURN) and `0x1b` (ESCAPE), both of which
self-send `0x46f`". The container dispatcher offers a key to the panel's
children before `R0714`, and the button's key handler takes Enter first,
so `R0714` is the route for Escape and only the fallback for Enter. The
command is the same, so the three inputs, the pager loop and the drop stand.
The Confidence paragraph's Medium clause read "`R0716`'s default arm
`R0390` and the base key handler's siblings were not read line by
line"; both were read.

### DLG-READ-006

- `R0701` opens the resource once only to decide whether to proceed
  (`DLG-ABSENT-003`) and discards the payload.
- `R0695` then passes the relative name (`L03319` reads the stack slot `+0x14`,
  the string before the `"main\text\"` and `".txt"` concatenations) to the
  constructor, which stores it at `panel+0x6c`.
- `R0720` rebuilds the same path from it
  (`L03320` passes L03321 `"main\text\"`, `L03322` passes L03323 `".txt"`),
  opens it again, and copies the whole payload into `panel+0x74`
  (`R0721` size → `R0722` alloc → `R0533` read → NUL).
- Nothing loads a mission's event files at map load: the only openers of this
  name family are those two, and both are inside the fire-time path.
- `vt+0x84` frees `panel+0x74` on close, so the text is resident only while the
  window is.

**Confidence.** High. Both opens are named instructions in two routines read
whole, and the relative-name hand-off is the instruction that makes the second
one possible. The "at load" rival predicts an opener in the map-load path, and
`R0701`'s arm is reached only by a posted `0x433`.

## Content model, layout and language roots

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DLG-MARKUP-007 | The content model is a tag scan, not a grammar; its vocabulary is fifteen lowercase literals, four of which occur in no event file on either root, and `iamfighter` and `fighter` in no EN one. | High / Medium | ● active (amended) | [EXP-0098](../experiments/EXP-0098-mission-text/) |
| DLG-FACE-008 | The whole file decides once whether the window has a portrait, each part decides which portrait it shows, and the two tests differ. | High / Unknown | ● active | [EXP-0098](../experiments/EXP-0098-mission-text/) |
| DLG-WRAP-009 | The text is wrapped and measured against its own font, and the window shows at most 7 lines of 17 px: the control scrolls by key only, and no shipped block has more than 7 lines. | High / Medium | ● active (amended, partially retracted) | [EXP-0098](../experiments/EXP-0098-mission-text/), amended [EXP-0410](../experiments/EXP-0410-dialogue-layout/) |
| DLG-LANG-010 | Only resources depend on the language root, and the roots differ in three mission events: EN silences three announcements that RU speaks. | High / Medium | ● active | [EXP-0098](../experiments/EXP-0098-mission-text/) |
| DLG-READER-011 | Before its EXP-0101 repair, this repository's shared `.res` reader rejected the RU `MAIN.RES`, so every tool using it was blind to the RU half of `main.res`. | High | ● active (amended) | [EXP-0098](../experiments/EXP-0098-mission-text/), [EXP-0101](../experiments/EXP-0101-mission-end/) |

### DLG-MARKUP-007

- `R0712(part, …)` walks the payload for `<` … `>`, lowercases each tag
  body (`R0723` → `R0724`) and substring-searches it.
- The literals, in the order the routine tests them: `part=%d`, `npc=`,
  `iamfemale`, `iammale`, `iammage`, `iamfighter`, `npcalive=`, `npcdead=`,
  `female`, `male`, `mage`, `fighter`, `sound=`, `tune=`, `tips=`.
- The part match is `CString::Find("part=%d")`, a substring test: a tag reading
  `part=10` also satisfies a search for part 1, and the first matching tag in
  file order wins.
- Corpus, both roots: 225 EN / 228 RU event files, 513 / 518 parts, max part
  13, and 0 parts on either root whose number is a prefix of an earlier part
  number in the same file. The hazard exists and the shipped data never trips
  it.
- Never used in either root's event files: `npcalive=`, `npcdead=`,
  `sound=`, `tune=`. `iamfighter` and `fighter` each count once in RU's and
  never in EN's; `fighter` is matched only through `iamfighter`
  (`evidence/markup-*.txt`, `DLG-TAGARM-027`).

**Confidence.** High for the vocabulary and the scan: one routine's own string
operands, read with `StrDump fnstr:`, and the lowercasing is a named call.
Medium for what each conditional literal does: only `part=`, `npc=` and
`tips=` were followed to their consumers, and the four `iam*` arms and the four
npc-flag arms were read as tests, not as effects.

**Amended.** The headline read "five of them are dead in the shipped corpus".
Its own list is four literals unused on both roots plus `iamfighter` and
`fighter`, unused on EN only (`evidence/markup-en.txt` 0, `markup-ru.txt` 1
each), over the 225 EN / 228 RU event files. `DLG-NPCTAG-019` bounds the
`sound=` zero to that corpus: on EN `sound=` occurs 7 times, all in `text/inn`.
It also finds one `iamfigter` misspelling per root, which no arm matches.
`DLG-TAGARM-027` reads the four `iam*` arms and the four speaker (npc-flag) arms
to their effects, and `DLG-SOUND-028` follows `sound=` to the pager; neither
follows `npcalive=`, `npcdead=` or `tune=`, and for those three the Medium
clause stands.

### DLG-FACE-008

- The constructor's `R0720` lowercases the entire payload and sets
  `panel+0x7c = (Find("npc") >= 0)` (`L03324` passes L03325 `"npc"`).
  `R0695` reads that flag to choose between the two layouts.
- Inside `R0712` the same field is rewritten from the current part's own
  tag, `panel+0x7c = (Find("npc=") >= 0)`, and `R0711` uses it to decide
  whether to refresh child 12.
- A file in which only one part names a speaker therefore gets the portrait
  pane for all of its parts, and the parts without `npc=` leave the previous
  face standing.
- The first test has no `=`: prose containing the letters `npc` would open the
  portrait layout on a file with no speaker at all.
- Corpus: on both roots the two tests agree on every file, 225/225 EN and
  228/228 RU, so the shipped corpus cannot discriminate them, and the
  separation is established from the instructions alone.

**Confidence.** High for the two tests and their different consumers: three
named instructions in two routines.

**Unknown.** Whether the disagreement is ever reachable in shipped data. By the
census above it is not, and no shipped file exercises the no-portrait layout
through this entry point.

### DLG-WRAP-009

- `R0725(ctrl, text)` calls `R0726(ctrl+8, text)`, which splits
  the string and calls `R0727(rect, line)` per piece into a single
  static accumulator at the global `L03326` (reset by `R0728(0,-1)`), then
  copies it into the control's own line array at `ctrl+0x64`.
- It then computes
  `ctrl+0x8c = min( Height(ctrl+8) / ctrl+0x60 , GetSize(ctrl+0x64) )`
  (`L03327`…`L03328`), with `R0697` = `bottom − top` and
  `R0729` = `[this+8]`, the array size.
- The line pitch `ctrl+0x60` defaults to font height + 2 (`R0699`:
  `param_10 == 0 → param_1[0x17] + 2`). `R0695` passes 0, so the
  dialogue window always takes the default.
- The control can scroll: `vt+0x80` `R0730` clamps a top line into
  `[0, size − visible]`, redraws, drives a scrollbar child and notifies with
  command `0x46d`. `R0695` gives it no scrollbar child, and no arm of
  the panel class answers `0x46d`. The control's own key handler `R0731`
  moves the top line on PageUp, PageDown, Up and Down while it holds focus
  (`DLG-KEYS-040`); the base key handler `R0717` binds Tab and the four
  arrows to focus movement.
- The constructor rectangle is 300×136 px with the portrait and 380×136
  without, and the control's base constructor shrinks its height to 135
  (`DLG-RECT-037`), so with the pitch of 17 the window shows at most 7 lines.
  The longest shipped part body is 230 characters (EN) / 211 (RU), and no
  shipped block wraps to more than 7 lines (`DLG-LINE-038`).
- The per-glyph advance that turns 300 px into a character count is
  EXP-0097's, and no figure here depends on it.

**Confidence.** High for the wrap, the clamp formula, the default pitch and
the control's key scroll: named instructions plus two one-line accessors read
whole. Medium for "no shipped block has more than 7 lines" (`DLG-LINE-038`, a
transcription of the wrapper, not an emulation) and for whether the text
control holds focus when the panel opens, which decides whether its key scroll
is live (`DLG-KEYS-040`).

**Amended.** The headline read "this window clamps the line count rather than
scrolling". The fifth bullet read "The rect is 300×136 px with the portrait and
380×136 without", with no live height, and the fourth ended "The base key
handler `R0717` binds Tab and the four arrows to focus movement, not
scrolling". The Medium clause read "text past the clamp is unreachable in this
window": `R0390`, the panel's default command arm, was not read, and
whether `panel+0x38`'s navigation links are ever filled was not established.
EXP-0410 corrects three points. The live rectangle is 135 px high, so the cap
is `floor(135 / 17)` = 7 lines and not 8. The text control has its own key
handler, so text past the clamp is reachable by key in principle. `R0390`
was read: it offers a key to the focused child, then to every child in order,
then to the panel's own slot. The wrap, the clamp formula, the default pitch
and the absence of a scrollbar child stand.

### DLG-LANG-010

- `rom.exe` is byte-identical across both preserved roots and the GOG install
  (`sha256 942e9b72…`, 1 977 344 B), so every geometry, index and format string
  in this ledger is shared.
- The two language-carrying inputs are
  `main.res::text/battle/m<n>/event<NN>.txt` and the indexed table
  `main.res::text/main.txt`.
- The compiled indices survive the change of root: 274 lines on both, with
  entry 77 (the button), 140 (the win panel) and 141 (the lose panel) present
  on both. The EN lines measure 2, 17 and 14 bytes and are ASCII; the RU lines
  7, 16 and 16 bytes and are non-ASCII (`evidence/strtab-*.txt`).
- The event corpora differ: EN 225 files, RU 228. The three RU-only files are
  m100/event09, m130/event07 and m150/event10. All three are raised by scripts
  that are identical on both roots (28/28 maps, 0 differing raised sets), so
  those three announcements are silent in English and spoken in Russian.
- `speech.res` carries no `battle` subtree on either root (0 of 355 EN / 349 RU
  nodes), so the `speech\…\….wav` the pager composes never resolves for a
  mission event.

**Confidence.** High for the identical image and the three-file delta: a byte
hash, and two exhaustive node censuses whose per-mission diff is three lines.
High for the string table's index stability: the indices are compiled
constants and both roots have the same line count, which a compiled index
requires and a differing count would have refuted. Medium for the speech
clause: the `%s` of `"npc%02de%sp%d"` was not pinned, so the absence rests on
there being no `battle` node at all rather than on a composed name missing.

### DLG-READER-011

- As it stood before the EXP-0101 commit, `internal/rom.OpenArchive` required
  `(len(file) − registryOffset) % 32 == 0` (`internal/rom/res.go:38`).
- The RU root's `MAIN.RES` is 4 922 540 bytes with registry offset 4 906 389.
  The remainder is 504 records of 32 bytes plus 23 trailing bytes, so the
  reader rejected it outright with `bad registry offset 4906389`.
- The registry itself is well formed and parses cleanly: 504 nodes against
  EN's 494, the count `RES-*`'s own EN↔RU compare already published.
- This is a fact about the instrument, not about the game (`METHODOLOGY.md` →
  *facts about the instrument*). It is recorded because it silently removes a
  whole root from any measurement: `tools/campaign -mode msgs` returned
  `open: …\main.res: bad registry offset` on RU
  (`evidence/raised-vs-shipped-ru.txt`), and any lane that measures "the RU
  text" through `OpenArchive` gets an error it may read as absence.
- EXP-0098 reproduced the parse locally (`tools/misstext` `openRes`) rather
  than patch a file its sibling lanes share.

**Confidence.** High. The constant is one line of committed source, the
arithmetic is exact, and the tolerant re-parse yields a node count that agrees
with an independently published EN↔RU figure. `internal/rom/res_test.go`
verifies the repair; its synthetic cases (the repository's own bytes) pin the
count-from-`@0x14` law and the residue tolerance without the install.

**Amended.** EXP-0101 repaired the reader on 2026-08-02.
`internal/rom.OpenArchive` now implements `RES-ACCEPT-031` directly: magic,
registry at `@0x10`, exactly `@0x14` records of 32 bytes, trailing bytes
ignored (`Archive.Residue` reports them). It opens 9/9 EN and 8/8 RU archives,
RU `MAIN.RES` as 504 nodes / 462 files, which is `RES-GEOM-028`'s own figure.
The measurement above describes the instrument before that commit; it is no
longer true of the tree, and `tools/campaign -mode msgs` on RU now returns
output rather than `open: … bad registry offset`.

## Modality while a panel is up

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DLG-STOP-012 | The world stops while a dialogue panel is displayed, and one instruction in the whole image decides it. | High | ● active | [EXP-0108](../experiments/EXP-0108-panel-modality/) |
| DLG-DIM-013 | Panel show requests one destructive level-3 shroud remap of the screen rectangle; per-source-channel 13/16 arithmetic and strict darkening apply to the full table path, not the reduced path. | High | ● active (amended, partially retracted) | [EXP-0108](../experiments/EXP-0108-panel-modality/), [EXP-0418](../experiments/EXP-0418-dialogue-backdrop/) |
| DLG-CLOCK-014 | The time a dialogue panel stops is discarded, not caught up, and the close's guard names the only state the engine expects the panel to interrupt. | High | ● active | [EXP-0108](../experiments/EXP-0108-panel-modality/) |
| DLG-DRAW-015 | Nothing else is drawn differently while a panel is up: seven instructions in six routines read the panel bit, and none of them is in a draw path. | High / Medium | ● active (amended) | [EXP-0108](../experiments/EXP-0108-panel-modality/) |
| DLG-ENTRY-016 | All six entry points and both mission-outcome panels take the identical mechanism; they differ in the state each runs from, plus one close arm. | High / Medium | ● active | [EXP-0108](../experiments/EXP-0108-panel-modality/) |
| DLG-MODAL-017 | A second, stronger modality, a nested message loop, exists in the image; no dialogue entry uses it, and it makes harmless the one panel the halt gate cannot see. | High | ● active | [EXP-0108](../experiments/EXP-0108-panel-modality/) |

### DLG-STOP-012

- `R0361` is `DLG-WIN-001`'s "then `R0361` shows it" and the only
  routine that puts a panel on the session window (`callto:R0361` = 19 hits /
  6 owners / 0 in orphan). It sets bit 3 of `campaign+0x3dc` for every panel
  it is handed (`L03329` sets bit 0x8, `L03330` stores the word back to `campaign+0x3dc`).
- The MFC idle handler loads that word at `L03331` and, once the phase is the
  campaign's, tests it: `L03332`/`L03333`/`L03334` compare the phase with 0x2, test the word against 0x4008 and jump to `L03335` when the result is non-zero. `L03335` sets the keep-idling flag and jumps to
  the handler's common tail, running no pacer arm at all.
- `EnumRefs imm:4008` is 1 hit / 1 owner / 0 in orphan: that test exists once
  in the image.
- Against every other branch of the same routine, the halted branch omits:
  - the sub-tick `R0147` (`callto:R0147` = 4 hits / 4 owners / 0 in
    orphan: `R0260`, `R0454`, `R0455`, `R0099`;
    the branch calls none of them);
  - `ANIM-CLOCK-001`'s `0x401` presentation tick;
  - the `0x402` post;
  - `R0509` on the map view;
  - `R0192`'s command drain (`SESS-PAUSE-020`), which makes the halt
    stricter than the pause.
- All five branches then converge on `L03336` and do identical work there,
  so nothing in the tail can be what a panel changes.
- The rival that would make this cosmetic is a timer, and the import table
  refutes it: `SetTimer`, `KillTimer`, `timeSetEvent` and `timeKillEvent` are
  absent from rom.exe's entire import directory (`SESS-TIMER-022`). No
  callback exists that could advance `server+0x04` behind the panel; its only
  incrementer is inside `R0193` (`SESS-TICK-004`; `callto:R0193` =
  2 hits, both in the server family).
- This evidence also corrects `DLG-LIFE-005` (filed in `retracted.md`). That
  claim's "the bit is cleared by nine instructions, all nine inside
  `R0709`" is wrong in both halves, because its instrument
  `imm:fffffff7` is blind to the byte-width form: `L03314`, a byte mask with 0xf7, the
  clear the dialogue panel itself takes, is one it missed, and `L03315` is in
  `R0701`.
- Bit 3 is not exclusively `R0361`'s to set: `L03318`, a bit-set of 0x8, sets
  it directly in a shutdown bracket that darkens through `R0718` at
  level 0 and posts `WM_CLOSE`. Neither correction touches this claim, whose
  subject is the gate rather than the writer set.

**Confidence.** High. Every step is a named instruction in one of three
routines read whole. The gate is an image-wide enumeration returning 1 hit with
0 orphan, and the tick's issuer set another returning 4 with 0. An import-table
absence, the instrument this repository treats as immune, closes the timer
rival.

### DLG-DIM-013

- Inside `R0361`, between a DirectDraw Lock (`R0346` →
  `vt+0x64`, with `L01510` storing 0x6c to the 32-bit global `L01509`, the
  `DDSURFACEDESC` size) and Unlock (`R0312` → `vt+0x80`), the routine
  calls the remap with the whole surface: `L03337`..`L03338` pass, in push order, 0x3,
  `[L01259]`, `[L00618]`, 0x0 and 0x0, then call R0732.
- `L03339` is a RECT (`R0364` hands it to `CopyRect`, IAT
  `L01508`), so those two globals are its right and bottom.
- `R0732(x0,y0,x1,y1,L)` clips x to `[L01503]`/`[L01505]` and y
  to `[L01504]`/`[L01506]`, and per pixel does
  `dst = LUT[L·[L03340] + (dst>>3 or dst)]` through `[L03341]`
  (`L03342`, `L03343`, `L03344`). The
  `L03345` split on the 32-bit global `L03346` being zero or not chooses the 13- or 16-bit
  index. This is the table, stride global and pair of index forms
  `TERR-FOG-037` describes for the shroud blitters.
- `TERR-FOG-084` fixes its law as `out = (in × (16 − L)) >> 4`, so L = 3 is a
  per-channel gain of 13/16 = 0.8125.
- Re-executed over the whole 16-bit pixel space
  (`tools/panelmodal -mode shade`): 65 535 of 65 536 pixels get strictly
  darker, 1 (black) is unchanged, and none is brightened. Mean channel level
  0.5000 → 0.3937 in RGB565 and → 0.3911 in RGB555.
  - The selected show body performs one direct remap into the locked
    framebuffer before the panel's own `vt+0x34` draw. This establishes one
    direct call per show invocation, not native backdrop persistence or
    cadence. Painter side effects and other native repaint routes are Unknown.
- `callto:R0732` = 11 hits / 11 owners / 0 in orphan. The other ten are in
  the code-region, code-region, code-region and code-region control
  families. Five also push level 3 (`L03347`, `L03348`, `L03349`,
  `L03350`, `L03351`), one pushes 10 (`L03352`), one 8 (`L03353`),
  and three pass the level in a register (`L03354`, `L03355`, `L03356`).
- `R0361` is the only one of the eleven that passes the screen rect:
  `refto:L00618` (65 hits / 47 owners / 0 in orphan) contains `L03357` and
  none of the other ten owners. None of `R0361`'s five sibling show
  routines calls it at all.

**Confidence.** High for the call, its five arguments, the Lock/Unlock pair and
the once-ness: each is a named instruction in two routines read whole, and the
whole-screen extent is a property of the argument list, not an observation that
nothing smaller was seen. The numeric result is no stronger than
`TERR-FOG-084`, which is High: the probe re-derives nothing and applies that
law. It first reproduces that experiment's own two anchors (L = 16 all-zero,
L = 8 `== (px>>1) & mask`, 65536/65536 in both layouts), so the law applied is
visibly the one pinned.

**Amended.** DIALOGUE-055 narrows the arithmetic and strict-darkening clauses
to the full table path. Original reduced-table construction samples the low
channel at 4, 12, 20 and 28 before the gain; input black maps to pixel 3 at
level 3. DIALOGUE-056 distinguishes the requested screen rectangle from the
incoming clip. DIALOGUE-057 distinguishes one remap per show invocation from
  idempotence: a second explicit show compounds it. The prior formula-only
  probe did not execute either original routine. The call and arguments stand;
  the universal arithmetic and unchanged-black consequence are partially
  retracted in `retracted.md`. DIALOGUE-057 also narrows the former non-per-frame
  and persistence wording to the direct selected handlers. DLG-STOP-012's idle
  pacer gate does not establish painter outputs or every repaint route.
  Native backdrop retention, event order and cadence remain Unknown; the
  former categorical persistence clause is partially retracted there.

### DLG-CLOCK-014

- The dialogue panel is stored in no `campaign+…` slot: `R0695`
  allocates it, hands it to the shared panel constructor at `R0361` (call at `L03358`), and stores it nowhere.
  Its close therefore runs `R0709`'s default arm, after all 25
  stored-panel comparisons fail.
- The default arm reads the screen-state word `campaign+0x3dc` (`L03359`),
  clears its bit 3 (`L03314`) and writes it back (`L03360`). It then
  compares the cleared word with the constant 1 (set at `L03361`) and, when
  they differ, skips the restart (`L03362`). When they are equal it also
  requires the pacer phase `campaign+0x6bc` to be 2 (`L03363`). If both
  hold it calls the import `timeGetTime` (the indirect call at `L03364`,
  import slot `L00849`), stores the result in `campaign+0x3fc`
  (`L03365`), and sets `campaign+0x3e4` to 0 (`L03366`) and
  `campaign+0x400` to 0 (`L03367`).
- `campaign+0x3e4` is the pacer phase. `R0454` rebases its deadline base
  `campaign+0x3ec` to `timeGetTime()` only at phase 0: the phase test is at `L03368`, the clock call at `L02632`
  and the store into `campaign+0x3ec` at `L03369`.
- Zeroing the phase makes the first iteration after the close rebase and issue
  one sub-tick. An untouched phase would run the catch-up loop (the branch at `L02630`) until
  the phase counter, masked to 4 bits, wrapped back to 0: up to fifteen sub-ticks fired
  back to back (`SESS-CLOCK-021`).
- The equality test against 1 makes this discriminating rather than incidental: the
  restart happens only when clearing bit 3 leaves the word at exactly 1,
  `SESS-SCREEN-003`'s map-session-only state, and only in phase 2. A panel
  closed over the town (`campaign+0x3dc == 0`, `SHOP-TOWN-023`) takes the same
  arm, clears the same bit and restarts nothing.
- The close tail then redraws the frame window
  (the indirect call through slot `+0x34` of the object at `campaign+0xcc`, at `L03370`), which repaints
  `DLG-DIM-013`'s darkened pixels. It is skipped only when the word is 1 and
  `mapview+0x80` is zero; then the resumed pacer's next frame repaints them.

**Confidence.** High. One routine read whole, every step asserted against the bytes.
The raw bytes at `L03359` (a 14-byte run of load, mask, branch and compare instructions),
carry the jump-on-zero with an 8-bit displacement of 0x3b that fixes the no-op exit at `L03371`
independently of any disassembler.

### DLG-DRAW-015

- No lighting change, no palette change and no per-frame dimming.
- `EnumRefs disp:3dc`, whole image: 242 hits / 80 owners / 0 in orphan or
  undisassembled code. Every hit that could involve bit 3 was followed to the
  instruction that tests it, a step a `disp:` sweep cannot take.
- The one mask carried in a register, the byte test of `+0x3dc` at `L03372`, is
  bit 0: `L03373` loads 1 into that register at the head of `R0232`, never reassigned
  before it. The four byte-wide loads `L03374`, `L03375`, `L03376`,
  `L03377` test `0x3`, `0x1`, `0x2`, `0x2` respectively.
- That leaves seven readers of bit 3, in six routines:
  - `L03333` (the test against 0x4008, the halt gate; `imm:4008` = 1 hit
    image-wide);
  - `L03312` (the announcement drop, `DLG-LIFE-005`);
  - `L03378` (the close arm, `DLG-CLOCK-014`);
  - four early-outs in three routines:
    `L03379` (a byte test of `+0x3dc` against 0xa) in `R0733`, `L03380`
    and `L03381` in `R0734`, and `L03382` in `R0735`. Each
    returns or jumps past its body when the bit is set.
- None of the seven lies in the terrain/sprite/shroud blit family
  (code-region) or the map-view draw family
  (code-region) that `TERR-*` establishes.
- The two composite masks that read the word inside the code-region family do
  not contain bit 3: `L03383 TEST …,0x226` is bits 1, 2, 5, 9 and
  `L03384 TEST …,0x627` is bits 0, 1, 2, 5, 9, 10.
- A consumer needs no second render path for "a panel is up", only
  `DLG-DIM-013`'s one-shot remap.

**Confidence.** High for the reader enumeration with its instrument and orphan
count stated, for the register-carried mask being bit 0, and for the two mask
arithmetics. Its blind spot: a wholesale block copy (string move or `memcpy`) of the campaign
object carries no displacement, and no `disp:` sweep can see it. Medium that
the four early-outs are input paths rather than draw paths: each was read at
its guard and the two instructions after it, not end to end. `R0733`
and `R0734` read the cursor globals `[L01257]`/`[L01258]` that
the pacer's own edge-scroll block reads at `L03385`…`L03386`, and
`R0735`'s not-taken path calls the panel command router `R0716`.

**Amended.** The headline read "six instructions read the panel bit" and the
card "six readers of bit 3" and "three early-outs", against seven listed
addresses. `evidence/enum-en.txt` and section 11 of
`evidence/rom-panel-excerpt.md` show seven bit-3 tests in six routines;
`R0734` holds two (`L03380`, `L03381`), so the early-outs are four
instructions in three routines. The draw-path negative holds for all seven
addresses.

### DLG-ENTRY-016

- `callto:R0695` = 7 hits / 6 owners / 0 in orphan, reproducing
  `DLG-WIN-001`'s six. All seven reach the same `L03358` call of R0361
  inside the constructor, so the bit, the darkening and the gate are the same
  at every entry, and whether the entry points differ reduces to whether the
  state differs.
- The gate needs three conditions at once:
  - `campaign+0x3dc & 1`, a map session (the bit-0 test at `L03387` diverts first);
  - `campaign+0x6bc == 2`, the campaign;
  - `campaign+0x6b8 != 0`, this process owns the simulation (`L03388`
    diverts a pure network client to `R0640` before the mask is ever
    tested).
- The mission-event entry, `R0701`'s `0x433` arm, is the one that runs
  with a map session on screen.
- The inn, the two mercenary-hall entries, the shop keeper and the training
  hall run from the town, which `SHOP-TOWN-023` fixes at
  `campaign+0x3dc == 0`. There `R0361` still darkens the screen and sets
  bit 3, and the halt and the clock restart are no-ops: a mechanism that
  executes and changes nothing observable.
- `DLG-PATH-002`'s two outcome panels use the same routine:
  `L03389` calls R0361 for the win panel (string-table entry 140, stored
  at `campaign+0x114`) and `L03390` for the lose panel (entry 141,
  `campaign+0x110`). Both darken and both halt.
- They diverge only on close. `R0709` dispatches on 25 stored panel
  pointers. `+0x110` has its own arm (`L03391`: clears bit 3, posts `0x41e`
  or `0x418`, and does not restart the clock). `+0x114` appears in none of the
  25 and falls through to the default arm with the dialogue panel.

**Confidence.** High for the shared path, the two outcome arms and the 25-slot
list. `R0709` was read whole and every comparison form in it collected,
not only the common compare against a frame-local slot: two of the twenty-five are
a plain register-to-register compare after a separate load, and reading only the common form
is how a slot list loses a member. Medium for the five town entries: this claim
fixes the state the gate requires and takes the town's state from
`SHOP-TOWN-023`. The five callers' own reachability was not read, so "no town
dialogue can ever run with bit 0 set" is not claimed.

### DLG-MODAL-017

- `R0736(panel)` calls `R0361` and then runs a nested
  `GetMessageA` loop: `L03392` loads the import `[L03393]`
  (`GetMessageA`), `L03394` calls through it,
  `L03395` compares the stack slot `+0x14` with 0x44c, `TranslateMessage`
  `[L03396]` and `DispatchMessageA` `[L03397]`. It loops until the
  close message `0x44c` arrives, re-posts it and returns.
- The loop never returns to `CWinThread`'s idle processing, so `R0560`
  does not run at all: nothing is paced, gated or repainted, whatever
  `campaign+0x3dc` and `+0x6bc` hold.
- `callto:R0736` = 18 hits / 11 owners / 0 in orphan, and none of
  `DLG-WIN-001`'s six entry points is among the owners. Four of the eighteen
  are inside `R0701` itself (`L03398`, `L03399`, `L03400`,
  `L03401`), in arms other than the `0x433` one. That routine therefore
  appears in both call lists with a different meaning in each, and the two
  lists must be separated by call site rather than by owner.
- Two of the eleven owners, `R0737` (`L03402`) and `R0738`
  (`L03403`), store their panel at `campaign+0x3a8` immediately before the
  call. That is the single pointer `R0361` special-cases:
  `L03404` reads the field `+0x3a8` of the campaign, `L03405` compares it with the panel, and
  `L03406` sets 0x80 in the word's second byte, bit 15, which the `0x4008` mask does not contain.
- The image therefore contains a panel whose show sets a state bit the idle
  gate never tests, a halt that cannot fire. It is not a defect only because
  that class's modality comes from the message loop instead: the
  `DLG-EMPTY-004` shape, twice over.
- `disp:3a8` = 10 hits / 6 owners / 0 in orphan: those two stores, the
  constructor's initialiser `L03407`, one reader in `R0739`, and one
  clearing arm, `L03408`, which masks the second byte with 0x7f.

**Confidence.** High. The loop, its terminating message, the enumeration with
its orphan count, the `+0x3a8` special case and its clearing arm are each named
instructions.

## The npc tag, the speaker and its figure

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DLG-NPCTAG-018 | The number in `<npc=N>` is the `npc<n>` section id of `scenario.res::npc.reg`, established by a subscript path, not by two shipped files agreeing. | High / Medium | ● active | [EXP-0141](../experiments/EXP-0141-npc-fields-and-tag/) |
| DLG-NPCTAG-019 | On EN the `<npc=` tag occurs in four `main.res` text families, which corrects two shipped-usage points of `DLG-MARKUP-007`. | High / Medium | ● active (amended, partially retracted) | [EXP-0141](../experiments/EXP-0141-npc-fields-and-tag/) |
| DLG-FIGURE-020 | The speaker's figure is the world figure compositor's output, called on the speaker's own drawable with the stencil surface null. | High | ● active (amended) | [EXP-0166](../experiments/EXP-0166-dialogue-dress/) |
| DLG-FIGURE-021 | No layer is excluded at the dialogue site, and the head slot is equipment slot 6, drawn in both halves of the compositor. | High | ● active | [EXP-0166](../experiments/EXP-0166-dialogue-dress/) |
| DLG-SPEAKER-022 | The speaker is a live actor when one matches the section, and otherwise a synthesised drawable with twelve empty equipment slots. | High / Unknown | ● active | [EXP-0166](../experiments/EXP-0166-dialogue-dress/) |
| DLG-SPEAKER-023 | The live-actor predicate is seventeen `Flags` tokens plus two record comparisons and a state gate, and most shipped speakers are on the composed-figure arm. | High / Medium | ● active | [EXP-0166](../experiments/EXP-0166-dialogue-dress/) |
| DLG-DRESS-024 | A live named speaker is drawn in its spawn outfit; mission 40's `npc25` joins its section, its `Data.bin` template and its map label in three independent directions. | High / Medium / Unknown | ● active (amended, partially retracted) | [EXP-0166](../experiments/EXP-0166-dialogue-dress/), [EXP-0192](../experiments/EXP-0192-mission20-party-boundary/) |
| DLG-SPEAKER-041 | Among the party's heroes, the stock mercenaries and the placements of `scn:140.alm`, no actor answers `npc62` (Human, Mage, Female, Face 4), so the chapter-140 tavern's dialogue draws the synthesised, unequipped face sheet `fmage\4.256`. | High / Medium / Unknown | ● active | [EXP-0413](../experiments/EXP-0413-speaker-figures/) |
| DLG-SYNTH-042 | When no live actor answers `npc21` to `npc24`, the synthesiser builds a Hero drawable whose face is the `Face` key of the `npc.reg` archetype chosen by its sex and class bits; the tokens `Me` and `!Mage` have no arm. | High / Unknown | ● active | [EXP-0413](../experiments/EXP-0413-speaker-figures/) |
| DLG-FACEBYTE-043 | For an npc-arm Humans actor in zero mode, `R0656`'s tail ORs the gender local into bit 7 of `actor+0x4b` after the streamer's face store; a stored face of 0 names sheet 0, which neither root ships, and the load aborts. | High / Medium / Unknown | ● active | [EXP-0413](../experiments/EXP-0413-speaker-figures/) |

### DLG-NPCTAG-018

- `R0712` calls the registry loader `R0499` on entry
  (`L03409`), scans `<`…`>`, lower-cases the tag body (`R0723`,
  `DLG-MARKUP-007`), and requires
  `Find(sprintf("part=%d", caller's part)) >= 0`. Then `Find("npc=")` →
  `L03410 CString::Mid(idx+4, 5)` → `L03411 sscanf(mid, "%d", *(frame slot +0xc))`.
- In the caller `R0711` that out-parameter is the frame local at −0x24, and it
  reaches the npc array unmodified: `L03412` loads it and passes it as an argument to the call at
  `L03413` of R0740. The argument passes through
  `R0740` (its first argument) to `R0741` to `R0742` /
  `R0743` with no arithmetic in any frame, ending at
  `L03414` loads the array base from the global `L03415` and
  `L03416` reads its 32-bit element at the index, and at `L03417`…`L03418`.
  That is the array whose element *i* the loader fills from section `npc<i>`
  (`REG-NPC-088`).
- The same local also feeds `sprintf("npc%02de%sp%d")` at `L03419`, the
  speech leaf, so the tag's number names a file as well as a section. A second,
  gated call at `L03420` is reached only for `21 <= N <= 24` (`REG-NPC-090`).
- `R0742` uses the record to decide whether an already-live actor is
  this npc: it compares `record+0xc` against `actor+0x24` under the `Flags`
  token `Face` (`L03421`/`L03422`) and `record+0x10` against `actor+0x20`
  under the token `Picture` (`L03423`/`L03424`).
- Corpus, both roots: 59 distinct tag values, the same 59 on each root, and
  59/59 name an `npc<n>` section that has a record; 46 records no tag names.
  That census is corpus agreement and is not the reason for the grade.

G2: the tag's domain is the array bound, 0..255, and the shipped ids reach
132. An author may add a section and tag it with no code change, but a tag
naming a slot with no section dereferences a null array element in
`R0743`.

**Confidence.** High for the identity: an unbroken register-to-subscript path,
every hop a named instruction, and no instruction between the `sscanf` and the
the 32-bit element read at `L03416` alters the value. Medium for the 59/59 census.

### DLG-NPCTAG-019

- Walking `main.res` node by node on EN, `<npc=` occurs 718 times over 268
  `.txt` nodes in four families: `text/battle` 225 nodes / 535 tags,
  `text/inn` 36 / 135, `text/shop` 4 / 8, `text/training` 33 / 40
  (`evidence/tag-families-en.csv`).
- RU figures come from a byte census of the raw file, not a node walk. RU
  carries 762 tags with the identical 59 distinct values.
- Correction 1: `DLG-MARKUP-007` records `sound=` as "never used on either
  root". That census was `text/battle`, its own 225 EN event files. `sound=`
  occurs 7 times on EN, all in `text/inn` (as `sound="npcNmNpa"`,
  `sound="imfNmNpN"` and siblings). The clause is true of its corpus and false
  of the file: three node families and 183 tags lay outside it.
- `npcalive=`, `npcdead=` and `tune=` remain 0 across all four families, so
  those three are dead in shipped EN data by a wider census.
- Correction 2: EN and RU each ship one `iamfigter`, a misspelling of the
  parser's literal `iamfighter`; `iamfigter` occurs in no `rom.exe` string.
  `CString::Find("iamfighter")` cannot match it, so that arm is unreachable for
  that part on either root. `iamfighter` itself is used once on RU and never
  on EN, as `DLG-MARKUP-007` reports.

G2: the vocabulary is fixed in the image. An author may use any of the fifteen
literals in any of the four families, and the three dead literals are usable
today with no code change.

**Confidence.** High for the per-family counts and the two shipped misuses: an
exhaustive node walk of one root and a byte census of both. Medium that no
fifth family carries the tag on RU: the RU count matches the raw-file grep
exactly, but RU was not walked node by node.

**Amended.** EXP-0171 corrects the literal-run enumeration, and `retracted.md`
withdraws it. The withdrawn sentence: "The parser's `.data` literal run is
twenty-four strings at `L03321`…`L03425`, of which fifteen are the tag
vocabulary and the rest are `main\text\`, `%d` ×4, `.txt`, `.res` and the
format `npc%02de%sp%d`." Read through the PE section table,
`[L03321, L03426)` holds twenty-five non-empty strings plus a lone `0x01`
byte at `L03427`. `%d` occurs five times, not four. `npc`, `Start` and a
second `female` at `L03428` are in the run and were not listed. `.txt` and
`npc%02de%sp%d` are not in it: the format string is at `L03429`, below the
run, and `.txt` is at `L03430`. `Start` is the omission that matters: it is
the `npc.reg` key read at `L03431` that gates all four speaker-conditional
arms (`DLG-TAGARM-027`). The per-family tag counts, the `sound=` correction and
the two shipped misspellings stand. `retracted.md` also withdraws the reason
given for the RU raw-file census: "because the RU `MAIN.RES` cannot be opened
by this repository's reader (`DLG-READER-011`)". `DLG-READER-011` records
EXP-0101's repair of that reader, which precedes EXP-0141 and opens RU
`MAIN.RES` as 504 nodes. The RU figures stand as a raw-file census, and so
does the Medium grade for RU carrying no fifth family.

### DLG-FIGURE-020

- `R0740(npcId)` (`UNIT-PICT-036`) has one caller, `R0711`,
  which fetches child 12 at `L03432` and hands it the return value at
  `L03413`/`L03433`.
- It allocates two surfaces through `R0744(w,h)`: `(0x58, 0x6c)` =
  88 x 108, the one the panel receives, and `(0xa0, 0xf0)` = 160 x 240, a
  scratch canvas.
- It branches on the flag word `+0x18c` of the actor held in the frame local −0x14, masked with
  0x11 (`L03434`..`L03435`), jumping to `L03436` when non-zero. Clear takes
  `UNIT-PICT-036`'s flat-BMP path. Set takes `L03436`..`L03437`, a virtual call of the slot at
  offset 0x80 with the arguments 0x0, the frame local −0x34 and 0x0, on the
  speaker.
- Slot `+0x80` of both drawable vtables holds `R0745`, read as raw
  dwords `L03438 -> R0745` and `L03439 -> R0745`: the world
  figure compositor of `HERO-APPEAR-051` and `UNIT-FIGURE-032`.
- That routine pops 12 bytes on return, and its three parameters are read at their
  consumers:
  - arg2 (stack offset +0x8a0) is the picture surface, cleared at entry;
  - arg3 (stack offset +0x8a4) is `HERO-FIGURE-057`'s stencil surface, skipped whole
    when zero (`L03440`/`L03441`, a compare against zero held in another register, jumping on equal);
  - arg1 (stack offset +0x89c) is a third surface reached only through
    the same compare-and-jump-on-equal test at `L03442`.
- The dialogue passes `(0, canvas, 0)`, so the colour pass runs and the
  click-map pass does not.
- The panel then blits a 72 x 96 window of the 160 x 240 canvas, or a 72 x 92
  window on the default arm, to `(8,7)` of the 88 x 108 surface
  (`REG-NPC-089`, `DLG-PORTRAIT-036`).
- Every equipment-tree `.256` on both roots is a single 160 x 240 frame: 59 face
  sheets, 780 of 784 `primary/` and 144 `secondary/` parsed, and 4 `primary/`
  nodes too short to carry one. A layer therefore covers the canvas rather than
  being placed on it.

**Confidence.** High. Each load-bearing run was byte-verified against the raw
image through the PE section table on three roots, 0 mismatches, and both
vtable slots were read as raw dwords at fixed addresses, which consults no call
graph.

**Amended.** The blit bullet read "to `(8,7)` of the 88 x 108 canvas". In this
card the canvas is the 160 x 240 scratch surface; the 88 x 108 target is the
surface `R0744(0x58, 0x6c)` allocates and the panel receives
(`evidence/listing-dialogue-portrait.txt`). The blit geometry stands.

A second correction (EXP-0410): the blit bullet read "a fixed 72 x 96 window".
The window is 72 x 96 when the second word of the record's portrait quadruple
is not -1 and 72 x 92 when it is -1 (`DLG-PORTRAIT-036`). The canvas, the
target surface and the destination `(8,7)` stand.

### DLG-FIGURE-021

- `R0745` has no parameter that selects layers: a 12-byte callee-popped frame, three
  parameters, all three surfaces (`DLG-FIGURE-020`).
- The layer set comes only from `drawable+0x15c + 4i` for `i = 0..11`
  (`L03443`, `L03444`; an empty slot skipped at `L03445`) and from
  `drawable+0x18c`. The dialogue's layer set is therefore every other caller's,
  `HERO-FIGURE-059`'s two orders included.
- Equipment slot 6 is the head slot, fixed by two instruments that read none of
  each other's evidence:
  - the `Armors` collection's `Slot` column (`ITEM-ARMSLOT-031`), whose value 6
    selects exactly `Hat`, `Cap`, `Low Hat`, `Helm`, `Soft Helm`,
    `Chain Helm`, `Full Helm` and `Plate Helm` and none of the other 22 rows on
    either root;
  - the compositor, which stamps `primary[5]` with the byte `6` at `L03446`
    and `L03447`. The stencil byte is the equipment slot number
    (`HERO-FIGURE-058`), and drawable index `k` is slot `k+1`
    (`HERO-APPEAR-050`).
- `primary[]` is based at stack offset +0x28 (`L03448`), so stack offset +0x3c is
  `primary[5]`. It is colour-blitted in both halves:
  `L03449`…`L03450` call the virtual slot at offset 0x18, non-mage, and
  `L03451`…`L03452` mage.

G2: an author may put any item in slot 6 through that item's own `Armors` row.
The head layer's presence in the dialogue figure is not separately
controllable, and suppressing it is a code change.

**Confidence.** High. The parameter count is the routine's own `RET`
immediate, and the pushes at the call site account for all three. The four
head-slot runs are byte-verified on three roots. The slot identification joins
a data column to an image immediate through the index rather than through a
name.

### DLG-SPEAKER-022

- `R0741(npcId)` calls R0742 (`L03453`) and, only when that
  returns 0, R0743 (`L03454`).
- `R0742` returns a live client actor (`DLG-SPEAKER-023`).
- `R0743` allocates `0x1b0` bytes and constructs at `R0593`
  (`REG-NPC-088`), whose first act (`L03455`..`L03456`) is a fill of 0xc zero dwords
  from `+0x15c`, clearing the twelve visible-equipment slots.
- Nothing on the synthesiser's path writes them. Two searches:
  - a byte scan of the whole extent `[R0743..R0742)`, 1107 bytes on each
    of three roots, finds 0 occurrences of the dword `0x0000015c`, the encoding
    any dword displacement of that field must carry;
  - `HERO-APPEAR-054` independently enumerates the array's writers as the four
    message stores in `R0509` plus the constructor and the destructor,
    none of which is on this path.
- A synthesised speaker contributes no equipment layer, and its figure is the
  face sheet `graphics\equipment\<figure>\<face>.256` alone (`HERO-DOLL-078`).
- A live speaker's figure carries whatever its twelve slots hold at that
  moment. Those slots change only by a re-send (`HERO-APPEAR-054`), so
  equipping an actor changes what its dialogue figure wears.

**Confidence.** High for the two arms and for the empty array. They are named
instructions, and the negative rests on a published writer enumeration and not
on the extent scan alone, which cannot see a wholesale copy of the containing
object.

**Unknown.** Which arm a given shipped dialogue takes at run time, which is
per-mission state.

### DLG-SPEAKER-023

- `R0742(npcId)` walks the client actor list at `this+0x9b8` and filters
  by runtime class against `L00620` (`L03457` calls R0214 and
  `L03458` branches on equal).
- It combines one term per token present in the section's `Flags`, in this
  order, with each `PUSH`'s address: `Me` `L03459`, `Mage` `L03460`,
  `Female` `L03461`, `Hero` `L03462`, `Human` `L03463`, `MySex`
  `L03464`, `MyClass` `L03465`, `!Me` `L03466`, `!Mage` `L03467`,
  `!Female` `L03468`, `!Hero` `L03469`, `!Human` `L03470`, `!MySex`
  `L03471`, `!MyClass` `L03472`, `Platoon` `L03473`, `Face` `L03474`,
  `Picture` `L03475`.
- The bits of `+0x18c`: bit 0 `Hero`, bit 1 `Mage`, bit 2 `Female`, bit 4
  `Human` and bit 5 `Me`. Bit 5 is read at `L03476` (a mask with 0x20) and tested
  at `L03477` (the frame local −0x38 against zero). `MySex` and `MyClass` compare
  bits 2 and 1 against the player's own drawable at `[this+0x3f54]+0x18c`.
- `Face` compares `record+0xc` against `actor+0x24` (`L03421`/`L03422`),
  and `Picture` compares `record+0x10` against `actor+0x20`
  (`L03423`/`L03424`).
- `L03478`..`L03479` compare the byte at `+0x15a` with 0x1; a following signed less-or-equal branch
  rejects any candidate above 1, and the first survivor is returned at
  `L03480`.
- Corpus, all three roots identically: of 105 `npc<n>` sections carrying a
  `Flags` list, 57 carry `Hero` or `Human` and take the composed-figure arm and
  48 take the flat arm. Of the 36 distinct `<npc=N>` values in
  `main.res::text/battle/m*`, the same 36 on EN and RU, 29 are composed-figure
  and 7 flat.

**Confidence.** High for the token list and the bits: a complete read of one
routine's string operands, each with its own `PUSH` address, and each bit read
at its own `AND`. Medium for the two censuses, which are corpus agreement.

### DLG-DRESS-024

- Mission 40's sole npc-arm placement is `npc25`, the same section its
  dialogue tags. Its `Flags` are `Hero,Face,!Female,!Mage`, and `DataBinID=42`
  resolves `Humans[42] PC_Paladin`, matching the map's `Paladin` label.
- Seven cells equip slots 1, 6, 7, 8, 9, 10 and 12, including a plate helm.
- Across both roots, 166 Humans rows name equipment: 156 fill slot 1, 46 slot 2,
  36 slot 4, 39 slot 5, 110 slot 6, 148 slot 7, 81 slot 8, 64 slot 9, 105
  slot 10 and 123 slot 12.
- A live named speaker is drawn in its spawn outfit.
- `npc25` keeps that outfit across the boundary by constructor mode
  (`PARTY-M20-031`): its exact `Hero` flag makes the npc arm request the
  player-character typeID overwrite, after which the culls leave the worn
  slots untouched.
- This is not universal to Humans: mission 20 transfers four dressed Humans
  that retain out-of-band table typeIDs and are removed (`PARTY-M20-031`).

G2: dress remains template data plus later equipment changes; persistence is a
separate classification decision.

**Confidence.** High for the mission-40 join and for the constructor-mode
reason `npc25` keeps its outfit, which rests on `PARTY-M20-031`. Medium for
the 166-row census.

**Unknown.** How the other 28 composed-arm tags resolve at run time.

**Amended.** EXP-0192 (`PARTY-M20-031`) narrows the claim, and `retracted.md`
withdraws the universal cross-boundary sentence: "A speaker who joins the party
keeps that outfit across a mission boundary", and "`PARTY-BAND-027` puts every
Humans-arm actor inside the surviving typeID band". Outfit persistence is
conditional on actor survival, not a property of the Humans arm: `npc25`
survives because it is flagged `Hero`, and mission 20's four dressed Humans
join and are then removed. The mission-40 outfit and all equipment counts
stand. The Confidence paragraph graded "the corrected constructor reason", a
label the card did not carry. The graded clause is the `npc25` bullet's
`Hero`-flag reason, which `PARTY-M20-031` names constructor mode; the bullet
and the Confidence paragraph now name it.

### DLG-SPEAKER-041

- `npc62` reads the same on both roots: `Flags` `Human,Mage,Female,Face`, `Face` 4, no `Picture`, no portrait quadruple, no `DataBinID`.
- `R0741` returns the first client actor that passes every term of `R0742` (`DLG-SPEAKER-023`) and otherwise builds the synthesised drawable (`DLG-SPEAKER-022`). For `npc62` the terms read actor bits `0x10`, `2` and `4` of `+0x18c` and the face at `+0x24`. The list is the client's actor map at `this+0x9b8`. The inn view's walk `R0746` reads the same client fields `+0x9c4`, `+0x9c0` and `+0x9bc` (`L03481`, `L03482`, `L03483`, in the listing `TAVERN-ORDER-015` cites).
- The searched population is the party's hero drawables and the stock units the tavern builds for the stage's level (`TAVERN-ORDER-015`), and the placements of `scn:140.alm`, added in case the client keeps the mission map's actors. No actor map of the tavern's client was enumerated, and that the client holds the placements of `scn:140.alm` is assumed from the chapter number, not established.
  - `R0590` classes a hero drawable by its type id in `[0x20,0x40)` (`L03484`..`L03485`): bits `(old & 0x80) | 9`, plus Female and Mage from `typeID - 0x21`. Bit `0x10` is never set, so the `Human` term rejects every party hero.
  - The stock units are the 52 shelf rows of the `Data.bin` Humans table (rows 47 to 98: 13 types at four levels). Twelve are mages, `NPC03`, `NPC04` and `NPC05` at each level, with faces 2, 2 and 1 at every level; `NPC04` is the female one. The `Face` 4 term rejects all twelve.
  - `scn:140.alm` has 171 type-6 records on both roots: 170 Unit-arm records, one npc-arm record and no type-key or definition-ID record. No Unit-arm record has a type word in the hero range `[0x20,0x40)` (0 of 6672 EN, 0 of 3366 RU), so its drawable keeps the flags word at 0 (`L03486`, `L03487`; `UNIT-PICT-035`) and the `Human` term rejects it. The npc-arm record is `npc26` in hero mode, row 43 `PC_Elf` (face 4, male): the `Human` term rejects a hero drawable.
- The tavern grid's cell and left panel for the same record are separate presentations (`TAVERN-TALKPIC-016`, `TAVERN-TALKSTATS-017`) and are not read here.
- Over the 215 Humans rows of each root, the terms human band, mage type id 23 or 24, female by the gender column or its default, and face column 4 pass one row on each root: 208 `M141_MageTooth`, server id 517. Eight rows differ between the roots, none a mage row: rows 54, 67, 80, 93, 159, 162, 165 and 168 have type id 3 on EN and 4 on RU, and the last four also face 1 on EN and 29 on RU. It is placed only by the definition-ID arm of the spawner (id `0x205`, secondary word zero): `scn:141.alm` record 0 on both roots, and `Beast.ALM` (5 records) and `Cross.ALM` (1) beside the EN executable.
- Placement census: 38 EN maps with 8094 type-6 records (unit 6672, npc 15, type-key 2, definition-ID 1405) and 34 RU maps with 3991 (3366, 15, 1, 609); each definition id resolves to exactly one row. Each record was read as the actor its arm builds: the row's class bits (the type-key arm's row is the first Humans row with that type id, `ALM-CLS-052`), then the arm's own store of face (the secondary word's low byte) and of the female bit (flags bit 2), which the type-key arm and the definition-ID arm with a nonzero secondary word make (`L03488`..`L03202`).
  - The actors that pass the four terms are 57 on EN and 1 on RU: row 208 (7 EN, 1 RU) and, on EN only, 50 records of `Horror.alm` whose mage rows take face 4 and the female flag from the placement. The same counts result when the record's own type word stands for the actor's type id. The type-key records (2 EN, 1 RU) and the zero-mode npc records (11 per root) build none. None of them lies in `scn:140.alm`.
- The synthesised drawable starts `+0x18c` at `0x48` (`L03489`); `Human` ORs `0x10` (`L03490`), `Female` `4` (`L03491`) and `Mage` `2` (`L03492`), giving `0x5e`. The record has no `Start`, so the face is `record+0xc`, 4 (`L03493`..`L03494`).
- `R0740` takes the composed-figure arm because `0x5e & 0x11` is nonzero. `R0745` selects the directory by `0x5e & 6` = 6, `graphics\equipment\fmage\`, and names the face sheet `graphics\equipment\fmage\4.256` (`%s%d.256`). The archive holds `fmage` sheets 1 to 5 on both roots.
- Layers: no equipment (`DLG-SPEAKER-022`); no hero back layer, since `0x5e & 3` is 2 and the gate wants 1 (`HERO-FIGURE-144`); a horse layer only when `[+0x20]` lies in `0x11..0x15` (`HERO-DOLL-078`). With no portrait quadruple the picture window is the 72x92 default over the `t_back` and `t_border` surface (`DLG-PORTRAIT-036`).

**Confidence.** High for the resolver, the filter terms, the synthesiser's bits and face, the class setter's bits and the sheet named: named instructions and installed data. The executables are one program, so RU adds no confirmation of a code fact; the data facts were measured per root. Medium that no actor answers at chapter 140: the population is the party's hero drawables, the 52 shelf rows and the 171 placements of `scn:140.alm`, each read as the actor its arm builds; the Unit-arm result rests on `UNIT-PICT-035`. It excludes an actor entered by a runtime spawn, a trigger or a path this experiment did not read, and no save's actor map was enumerated. The passing placements of `scn:141.alm` and of the EN root maps reach the tavern's client only if that client holds another map's actors; that was not examined.

**Unknown.** The `+0x20` word of the synthesised drawable. `R0743` writes it only under a `Picture` token, the constructor does not write it, and the object comes from `R0747` and the runtime allocator (`R0722`, then `R0748`, not read). The horse layer draws only if that word lies in `0x11..0x15`. Which rows of the canvas the window frames, and the drawn pixels (`DLG-PORTRAIT-036`).

### DLG-SYNTH-042

- `npc21` to `npc24` read the same on both roots: `Flags` `Hero,Me,Start`, `Hero,Mage,!MySex,Start`, `Hero,!Mage,!MySex,Start` and `Hero,!MyClass,MySex,Start`. None carries `Face`, `Picture` or a portrait quadruple; all four carry `DataBinID` 26.
- A live actor answers first when one passes the terms (`DLG-SPEAKER-023`). `Me` is bit `0x20`, which the client sets on the actor it makes primary (`L03495`..`L03216`). The synthesiser runs only when no actor passes.
- `R0743` starts `+0x18c` at `0x48` and sets its bits from these tokens only (`L03496`..`L03497`): `Hero` ORs 1, `Human` `0x10`, `Female` 4, `Mage` 2; `MySex` and `MyClass` OR the primary actor's own 4 and 2 from `[this+0x3f54]+0x18c`; `!MySex` and `!MyClass` OR the inverted values. `Me`, `!Me`, `!Mage`, `!Female`, `!Hero`, `!Human`, `Platoon` and `Face` have no arm. `Picture` sets `+0x20` (`L03498`).
- `Start` reads `scenario\npc.reg`, takes the section chosen by `bits & 6` (`MaleFighter` 0, `MaleMage` 2, `FemaleFighter` 4, `FemaleMage` 6) and stores that section's `Face` key, default 1, as the face (`L03499`..`L03500`). A record without `Start` takes its own `Face` (`L03493`). The four keys read 5, 3, 1 and 1 in that order on both roots.
- Bits and figure by the primary's sex and class, identical on both roots:
  - `npc21`: `0x49` for every primary; `MaleFighter`, face 5, `mfighter\5.256`.
  - `npc22`: male primary `0x4f`, `FemaleMage`, face 1, `fmage\1.256`; female primary `0x4b`, `MaleMage`, face 3, `mmage\3.256`.
  - `npc23`: male primary `0x4d`, `FemaleFighter`, face 1, `ffighter\1.256`; female primary `0x49`, `MaleFighter`, face 5, `mfighter\5.256`.
  - `npc24`: male fighter `0x4b` `MaleMage` 3; male mage `0x49` `MaleFighter` 5; female fighter `0x4f` `FemaleMage` 1; female mage `0x4d` `FemaleFighter` 1.
  - The 16 results name four sheets, `mfighter\5.256`, `mmage\3.256`, `ffighter\1.256` and `fmage\1.256`; all four ship on both roots.
- Every result has bit 0 set, so `R0740` takes the composed-figure arm, with twelve empty equipment slots (`DLG-SPEAKER-022`). The compositor adds the hero back layer, `backm` or `backf` by bit 2, to every result whose `bits & 3` is 1: the results `0x49` and `0x4d`. The results `0x4b` and `0x4f` draw the face sheet alone (`HERO-FIGURE-144`). The `[+0x20]` word that gates the horse layer is unwritten here as in `DLG-SPEAKER-041`.
- The synthesised figure is an archetype selected by the record's tokens and the primary's sex and class. The synthesiser reads no party list and no key names a fixed picture per record. Only `npc21` is independent of the primary. A live actor that passes the terms is drawn instead (`DLG-SPEAKER-023`); the hero arm sets bit 0 and the client sets `Me` on the primary, which are the two terms of `npc21`.
- The filter's and the synthesiser's `MySex` and `MyClass` terms read the primary through `[this+0x3f54]` with no null test (`L03501`, `L03502`).

**Confidence.** High for the token arms, the archetype mapping, the `Face` keys and the table: one routine read whole, the registry of both roots read whole, each sheet looked up in the installed archive.

**Unknown.** Which arm a dialogue takes for a given party, that is, whether a live hero passes the terms. Whether a dialogue can run with a null primary. Which rows of the canvas the window frames (`DLG-PORTRAIT-036`).

### DLG-FACEBYTE-043

- Case: an npc-arm placement (`R0151`, flags bit 0) whose record lacks `Hero`. `R0749` returns whether the record holds the `Hero` token and the arm passes it to `R0497` as the constructor mode; mode 0 is zero mode. The arm joins the record's `DataBinID` to the Humans row whose server id, slot 24, equals it (`R0498`, a scan from the last row down that never tests element 0) and passes the row's name. The arm stores nothing to `+0x4b`. The type-key arm (`L03488`..`L03503`) and the definition-ID arm (`L03504`..`L03202`) do.
- `R0656` resolves the row again by name: it scans the table from element 1 and takes the first row whose name equals the given one (`L03505`..`L03506`; `R0750` reaches the runtime's exact string compare `R0751`, byte-wise or lead-byte aware by the flag at `L03507`). The 210 non-empty names of each root are unique, also ignoring case, so both routes reach one row. The 5 rows without a type id have empty names.
- Stores to `actor+0x4b` on the constructor path, in order:
  1. `L03508` stores the byte 1 at `actor+0x4b` in `R0185`: the default.
  2. `R0657` passes `&actor+0x4b` (`L03509` forms the address by adding 0x4b to the actor base) to the byte slot reader `R0281` for Humans slot 17, `face`. The reader stores only when the getter `R0752` does not return -1 (`L03510`, `L00927`), so a column of -1 leaves the default.
  3. `L03511` stores a byte at `actor+0x4b` in `R0656`: the suffix face, when it is above 0. A name with a `.` sets the gender local to 1 when the character after the dot is `f` and to 0 otherwise (`L03512`..`L03513`), and the digits after that character give the suffix face (`L03514`).
  4. `L03120`..`L03121`, the zero-mode tail: the gender local is shifted left by 7, merged into the face byte, and the byte is stored at `actor+0x4b`. Only the low bit of the local reaches bit 7.
- The gender local defaults to 1 (`L03515`), takes the name suffix when there is one, and then takes Humans slot 18, `gender ( is female? )`, unless that value is -1 (`L03516`..`L03517`). A non-zero mode, a `Hero` record, sets the type id to the local plus `0x21` or `0x23` and skips the tail (`L03518`..`L03519`).
- Input: the gender local. Of 215 Humans rows per root, 184 carry a gender value and 31 default to 1, 26 of them in the human band; the counts are equal on both roots, and no row name holds a `.`.
- Client: the first state message carries the byte (`PAL-FACE-005` identifies the byte and traces its carriage). `R0675` calls `R0059` with mask -1 (`L03520`, `L03521`) and `R0059` appends the byte under mask bit `0x4000` (`L03522`..`L03523`). The creation path of `R0509` passes it to `R0590`: for a type id below `0x1a` bit 7 becomes the Female bit `4`, the face is `byte & 0x7f`, and type ids 23 and 24 add Mage (`L03170`..`L03524`). The update path stores the raw byte to `+0x24` (`L03525`).
- Placements: 15 npc-arm placements per root, 4 in hero mode and 11 in zero mode. Each zero-mode placement resolves to one row. The sex term of the record agrees with the modelled bit 7 in 11 of 11 (one `Female`, ten `!Female`), and 10 of 11 pass their own record's filter against the actor built. `npc55` does not: row `F_BrigandLeader3` has face 25 and the record `Face` 21.
- Face 0: `R0745` formats `%s%d.256` with `[drawable+0x24]` (`L03526`..`L03527`) and no instruction remaps 0. No figure directory holds sheet 0: `mfighter` 1 to 31, `mmage` 1 to 13, `ffighter` 1 to 10, `fmage` 1 to 5, both roots. The sprite constructor `R0753` calls `R0535`, which returns 0 for a missing node, then formats `FATAL ERROR: can't load ` plus the path and calls `R0754` (`L03528`..`L03529`), the abort path of `TAVERN-TALKPIC-016`. A figure drawn for a face of 0 aborts the program.
- No shipped source stores face 0: the 210 human-band rows per root have effective faces 1 to 31, the 5 rows without a type id have none, and the type-key and definition-ID placements store no low byte 0 (2 and 1 type-key, 433 and 17 definition-ID records with a nonzero secondary word, EN and RU).

**Confidence.** High for the writer chain, its input and the face-0 code path: named instructions in routines read whole, and the data results per root. Medium that no other store follows the streamer's on this path. Instrument: `EnumRefs` `disp:4b`, `disp:4a`, `disp:49`, `disp:48` and `imm:4b` on the repaired project, reduced to memory-destination writes that start at `0x4b` (byte), `0x4a` (word) or `0x48` (dword, qword) and to additions of the constant 0x4b to a base, stack bases set aside: 79 hits in 62 owners. Four owners are read whole (`R0151`, `R0185`, `R0656`, `R0657`); 58 are classed only by constructor vtable constants, virtual slots and callers: 21 constructors of other classes, 8 virtual methods, 29 free functions. Blind spots: bulk copies (`REP MOVS`, `memcpy`), pointer tables, free functions handed the actor, ten non-stack `LEA` forms at `+0x48`, and `R0067`, which the placement driver calls after the spawner (`L00959`) and which stores 1 at `+0x48` through `[Z+0x3c]` in eight places: `SAV-GRPAI-563` types two of them (`L03530`, `L03531`) as `AI+0x48` of a group's AI block, and the other six were not typed.

**Unknown.** Which caller of `R0675` carries a placement's first message. Whether any update message carries mask bit `0x4000` for a person: the update path would store bit 7 into the face. The face of a person stored by a save, a trigger or a mod.

## Message number, mission number, tag arms and sound

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DLG-MSGNUM-025 | The message number is a dword from the script parameter to the `%02d` field, and the value 255 is byte-identical to the mission-lost sentinel. | High | ● active | [EXP-0171](../experiments/EXP-0171-script-message/) |
| DLG-MISSION-026 | `campaign+0x660` is the mission number and the map file number at once, and its writers are inside the record embedded at `campaign+0x548`, not at that displacement. | High / Medium | ● active | [EXP-0171](../experiments/EXP-0171-script-message/) |
| DLG-TAGARM-027 | The eight conditional markup arms test the player's hero or the speaker, and substring nesting makes four of them fire twice. | High / Medium | ● active (amended) | [EXP-0171](../experiments/EXP-0171-script-message/) |
| DLG-SOUND-028 | `sound=` names a wave and plays nothing; the pager plays it, and the name it composes when the tag is absent depends on the resource's own leaf name. | High / Medium | ● active | [EXP-0171](../experiments/EXP-0171-script-message/) |

### DLG-MSGNUM-025

- The sender writes it whole: `L03532` stores the 32-bit number at `msg+0xa`.
- The client dispatcher's `0xb6` arm reads it back whole at
  `L03533` (the 32-bit field at `msg+0xa`) and posts at `L03534` the message 0x433
  through `[L03535]`, with wParam the number and lParam 0.
- The handler `R0701` reads wParam from the stack slot `+0xb4`
  (`L03536`), compares it with 0xff (`L03537`),
  and stores it at `campaign+0x414` (`L03310`) before the branch and
  on both paths. The branch at `L03538` jumps to `L03312` when not equal, so the equal case goes to
  `L03539`, which supplies 0x431, the mission-lost panel.
- The dispatcher's own lose arm posts that value: `L03540` supplies 0xff, then
  `L03541` supplies 0x433.
- A script raising message number 255 is therefore indistinguishable from a
  lost mission at the first instruction that inspects it. `campaign+0x414`, the
  field read at `L03542` for a comparison with 0xff on the lose chain
  (`MISSION-PATH-015`), is the same field an ordinary announcement's number
  lands in.
- Past the sentinel the arm gates once on an open dialog
  (`L03312` tests bit 0x8 of the byte at `campaign+0x3dc`, `DLG-LIFE-005`), formats the path
  from `L03543` and `L03544`, opens at `L03305`, and skips the panel on a
  null result (`L03307`/`L03308`: a comparison with zero and a jump on equal to `L03309`,
  `DLG-ABSENT-003`).
- No `campaign+0x6bc` test occurs between `L03537` and `L03545`, so the
  event-text path is not gated on game mode.

G2: message numbers 0..254 and 256 upward are free; 255 is reserved by the
image. Lifting that costs a code change but no shipped bytes, since the corpus
maximum is 25.

**Confidence.** High. Three routines were read at instruction level, both
accesses to `packet+0x0a` are dword on their own operands, and the sentinel is
the same immediate at both ends. The mode clause is scoped to the two address
ranges named and says nothing about the dispatcher.

### DLG-MISSION-026

- Four format strings take it, each referenced once (`EnumRefs imm:`, 1 hit /
  1 owner / 0 orphan each):
  - `battle\m%d\event%02d` at `L03546`, read at `L03544`;
  - `%d.alm` at `L03547`, read at `L03548`;
  - `main\text\battle\m%d\briefing.txt` at `L03549`, read at `L03550`;
  - `m%d` at `L03551`, read at `L03552`.
- The map-file build is gated at `L03553` on the 32-bit field `campaign+0x6bc` being 0x2. The
  other arm takes the map-name CString at `campaign+0x6b4`, which the `-map`
  command-line arm writes at `L03554`.
- A campaign mission is loaded by number and any other game by name, so the
  event-text directory number is the map file number by construction rather
  than by convention.
- Corpus, both roots: the 28 `text/battle/m<N>` directories are the 28 campaign
  map numbers, in both directions.
- The field has no writer at its own displacement. `EnumRefs disp:660` returns
  7 hits / 3 owners / 0 in orphan and every hit is a read; `imm:660` returns 0;
  `disp:658`, `disp:668`, `disp:66c` and `disp:670` through `disp:67c` are
  empty or foreign.
- It is the record's `+0x118`, and `0x548 + 0x118 = 0x660`. The writers are
  `L03555` and `L03556` in `R0755`, `L03557` in `R0756`,
  `L03558` in the record constructor `R0757` with a zero value in the scratch register, and
  `L03559` in `R0677`.
- `R0755` refuses a number above the record's `+0x4` and reads `-1` as
  the value at `+0x4`. `R0756` accepts only a number that some entry of
  the `0x4c`-byte list at `+0x4c` of the object with count `+0x50` carries at its
  own `+0x4`.
- `EnumRefs callto:R0755`: 6 hits, 6 owners, 0 in orphan. The campaign's own
  start passes ten, at `L03560` (the argument 0xa), with the receiver `campaign+0x548`.

G2: the mission number set is data, not a compiled table; no run of 10, 20, 30
exists anywhere in the image as dwords or as words. A new mission costs a list
entry and an `<n>.alm`, and its text directory has to carry the same number.

**Confidence.** High for the field identity, its four consumers and the mode
gate: each is a named instruction, both literals' reference sets are
enumerated with the instrument and its orphan count stated, and the
embedded-record arithmetic is exact. Medium for the writer set being complete:
a whole-record `memcpy` carries no displacement and would appear in neither
sweep (`docs/INSTRUMENT.md` rules 4 and 9).

### DLG-TAGARM-027

- This claim closes the Medium clause of `DLG-MARKUP-007`, which read these
  arms as tests and not as effects.
- The four `iam*` arms test the player's own hero, subject
  `+0xd0` of the frame local −0x14 dereferenced through `+0x3f54`. Each abandons the tag and
  resumes the scan when tag and bit disagree:
  - `L03561` `iamfemale` with the 0x4 mask test at `L03562` and a jump on non-zero, rejecting
    when the bit is clear;
  - `L03563` `iammale` with the 0x4 mask test at `L03564` and a jump on zero;
  - `L03565` `iammage` with the 0x2 mask test at `L03566` and a jump on non-zero;
  - `L03567` `iamfighter` with the 0x2 mask test at `L03568` and a jump on zero.
- Bit `0x4` is female and bit `0x2` is spellcaster on the actor's `+0x18c`,
  read off these four arms' own polarities.
- The four sex and class arms test the speaker instead, behind two gates: the
  npc section's `Start` key (`L03431` supplies L03569, then
  the test of the return value at `L03570` and a jump to `L03571` on zero) and a resolved speaker
  (`L03572` calls R0742, then `L03573` jumps to `L03571` on zero). Either failing
  skips all four. They are `L03574` `female` with the 0x4 mask test at `L03575` and
  a jump on non-zero, `L03576` `male`, `L03577` `mage` and `L03578` `fighter`, on the
  speaker's own `+0x18c`.
- The scan is one forward pass over the tags. An arm that disagrees jumps to
  `L03579`, which resumes at `L03580` while the cursor is not at the
  terminator and returns 0 when it is; the pager turns a 0 into `0x445`, which
  closes the window (`DLG-LIFE-005`). A suppressed tag is therefore not a
  suppressed part: another tag declaring the same part can still match, which
  is how the shipped `iamfemale`/`iammale` pairs are authored.
- An arm that does not disagree falls through to the next arm on the same tag
  body, and `CString::Find` is a substring test. So `iamfemale` also satisfies
  `female`, `iammale` satisfies `male`, `iammage` satisfies `mage`, and
  `iamfighter` satisfies `fighter`. `iamfemale` also contains `male`, but the
  `male` arm re-tests `L03581 Find("female") == -1` and
  `L03582`, a jump on non-zero to `L03577`, skips its speaker test for any body containing
  `female`. That is the one guard.
- A part tagged for a female player therefore additionally requires a female
  speaker, and an `iamfemale`/`iammale` pair read to a female player by a male
  speaker satisfies neither tag, so that part is absent and the window closes.
- Corpus, `text/battle`, both roots (`evidence/tag-bodies.txt`): 535 EN tag
  bodies in 225 files and 543 RU in 228. `iamfemale` 9 and `iammale` 9 on each
  root; `iammage` 2 EN and 3 RU; `iamfighter` 0 EN and 1 RU; `female` 20 EN and
  22 RU; `male` 40 EN and 44 RU, with every `iam*` body counted in those.
- `mage` 2 EN / 3 RU and `fighter` 0 EN / 1 RU are matched only through
  `iammage` and `iamfighter`, so neither literal is used standalone on either
  root. `npcalive=`, `npcdead=`, `sound=` and `tune=` are 0.
- The misspelling `iamfigter` (`DLG-NPCTAG-019`) appears once on each root, at
  `npc=52` part 1, and matches no literal.

**Confidence.** High for the eight arms and the two gates: one routine read at
instruction level, every test a named instruction. High for the census: an
exhaustive tag-body walk of both roots' `text/battle` nodes, whose 535 EN bodies
reproduce `DLG-NPCTAG-019`'s independently published figure. Medium that a
nesting-suppressed part is absent in play: the fall-through and the loop tail
are read from branch targets, not observed.

**Amended.** The nesting bullet read "`iamfemale` also satisfies `female` and
`male`", beside the `male` arm's own `Find("female") == -1` guard.
`evidence/markup-arms.txt` shows that guard's jump on non-zero at `L03582` to `L03577`
jumping past the `male` arm's speaker test whenever the body contains
`female`, so an `iamfemale` body reaches the `female` arm's test only. The
four double firings, the two gates and the census stand.

### DLG-SOUND-028

- The parser's arm at `L03571` finds `sound=`. When the tag is absent it
  writes the empty literal `L03583` to the caller's out-parameter
  (`L03584`). Otherwise it skips six characters and takes the value to the
  next quote when the first character is one (`L03585` compares with 0x22), or to
  the next semicolon when it is not (`L03586` supplies 0x3b), storing it with
  `CString::operator=`. It reaches no sound routine.
- Its only caller is the pager `R0711` (`EnumRefs callto:R0712`: 1 hit,
  1 owner, 0 in orphan).
- The pager takes the resource name at `panel+0x6c`, cuts it at the last
  backslash (`L03587` supplies 0x5c), lowercases the leaf and tests `Left(5)`
  against `event` (`L03588` supplies L03589, `L03590` supplies 5, compared at
  `L03591`).
- A leaf beginning `event` composes `npc%02de%sp%d` at `L03429`; any other
  leaf composes `%sp%d` at `L03592`. Both are discarded when the `sound=`
  value is non-empty: `L03593` re-tests it and `L03594 JZ` skips the
  composition otherwise.
- The name that is used is built at `L03595`, `L03596` and `L03597` from
  `speech\` at `L03598`, the resource path's directory prefix, the value, and
  `.wav` at `L03599`.
- The mission-event family is therefore the one family selected by a file-name
  prefix rather than by its directory, and a `sound=` tag on it overrides the
  composed speech name rather than adding to it.
- Corpus, both roots: `sound=` and `tune=` occur 0 times in `text/battle`, so
  the override is unexercised by shipped data. `DLG-NPCTAG-019` found `sound=`
  7 times on EN, all in `text/inn`.

G2: authored audio on a mission announcement needs a `sound=` tag and a
matching `speech.res` node, with no code change and no change to shipped bytes.

**Confidence.** High for the arm, the caller enumeration and the family test:
one routine's own instructions plus a `callto:` sweep with its instrument and
orphan count stated. High for the corpus zero: an exhaustive tag-body walk of
both roots. Medium for the concatenation order, which is read from the operand
order of three CString helpers whose arity was inferred from their other call
sites rather than pinned.

## Inn mission-giver text and voice

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DLG-ZEROARM-029 | The inn's "nothing to offer" zero arm never advances `InnMission` whatever its shipped text says, and that text is not uniformly backward-only: it splits by data root. | High / Medium | ● active | [EXP-0393](../experiments/EXP-0393-inn-text-closure/) |
| DLG-INNVOICE-030 | Text and voice presence for the inn's mission-giver NPCs do not track each other in either direction, and differ per root independently of registry addressing. | High | ● active (amended) | [EXP-0393](../experiments/EXP-0393-inn-text-closure/) |

### DLG-ZEROARM-029

- This claim extends `REG-SCN-064`, which established the mechanism
  (`InnMission[i]==0` speaks a line keyed by the live main mission,
  `L03600`/`L03601`) without classifying its content.
- It reads every file the shipped corpus lets the zero arm reach: 9 of 13
  addressed combos ship.
- All 9 recap something the player has already done: "...we have obtained the
  Cloak..." (EN/RU Mission110), "My vow is fulfilled. The Envoy is avenged..."
  (EN/RU Mission130), "you had taught them a good lesson... even before we
  reached the city" (demo Mission50), "did you deal with the turtle?" (demo
  Mission90).
- EN and RU Mission50, a shared thirteen-part, mostly past-tense backstory,
  close with an explicit reciprocal proposal, "Help me to find his assassins
  and I will help you find the Cloak and Scroll", which reads as an offer in
  ordinary language.
- EN/RU Mission110 and Mission130 close with a forward-pointing suggestion
  short of a proposal ("Let's go check out the shop", "We should ask the
  shopkeeper").
- Only the pre-release root's three shipped instances (Mission50,
  Mission90 ×2) carry no forward-pointing content at all.
- Four more addressed zero-arm combos ship on no root read here (demo
  Mission60 ×2, Mission110, Mission130) and are not classified: no text exists
  to read.
- Only openings and closings are quoted, bounded to 140 runes per file; each
  file's own byte and tagged-part count is reported alongside, never its full
  text (`evidence/zeroarm-text.txt`).

**Confidence.** High for the mechanism reproduction: an independent walk of the
same registry reaches the same 22/22/24 addressed combos
`REG-SCN-064`/`REG-INNCLOSE-116` already report. Medium for the content
classification: it is a reading of 9 files against a recap/forward-pointing
distinction, not a traced consumer of any in-game flag, and none was found or
claimed to exist. It rules out a single uniform answer across the corpus: a
"never any forward content" reading is refuted by EN/RU Mission50, and a "reads
like an ordinary quest offer everywhere" reading is refuted by the 3 demo
instances and by the registry state itself, which never changes at any of the
9.

### DLG-INNVOICE-030

- Population: the union of all three roots' `text/inn/npc/*.txt` (32 distinct
  names) and the matching voice parts (149 distinct `inn/npc/*.wav` names), in
  `text/inn/npc/` and `speech.res::inn/npc/`.
- 31 of the 32 text paths are asymmetric across roots (`REG-INNCLOSE-116`'s
  closure counts). The one exception, `npc32m51.txt`, ships identically present
  on EN, RU and the pre-release root alike.
- Voice compounds the asymmetry. `npc25m150p-.wav` and `npc90m41p2.wav` ship on
  RU with no EN part of that number.
  `npc25m150pa.wav`/`npc25m150pb.wav`/`npc25m150pc.wav` and `npc30m151p6.wav`,
  `npc25m140p5.wav` ship on EN with no RU part of that number.
- Both directions occur inside files whose EN and RU text both ship, so text
  presence does not predict voice-part completeness.
- The pre-release root ships text and voice for the same four stems
  (`npc32m51`, `npc32m90`, `npc90m50`, `npc93m90`) and neither form for any of
  the other 28 `npc/` files. Two of those four (`npc32m51`, `npc32m90`) also
  carry demo-only voice parts under a distinct `mf_<stem>p<n>.wav` name absent
  from EN/RU entirely, alongside the `npc<stem>p<n>.wav` parts of the same stem.
- `DLG-WIN-001` already names this class, "the inn's NPCs (R0702, ×2)",
  as one of its six entries.

**Confidence.** High. Every figure is a direct per-root file-presence count
over a named, closed union of paths (`evidence/dialogue-presence.csv`). The
alternative "presence is 1:1" is refuted by seven named voice counterexamples
running in both directions plus the text-level asymmetry count, not by an
aggregate mismatch count alone.

**Amended.** The Confidence paragraph read "five named voice
counterexamples" against the seven `.wav` names the card lists.
`evidence/dialogue-presence.csv` carries all seven: two RU-only, five
EN-only.

## Panel geometry, line placement and input

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DLG-PANEL-035 | At 640x480 the panel is drawn 488x232 at (76,124): the 580x240 constructor rectangle is snapped and centred, then painted as a nine-piece `lm.256` frame with an 8 px shadow band. | High / Unknown | ✔ promoted (amended) | [EXP-0410](../experiments/EXP-0410-dialogue-layout/), [EXP-0419](../experiments/EXP-0419-dialogue-render/) |
| DLG-PORTRAIT-036 | The portrait pane is an 88x108 picture surface at panel offset (30,54): a 72x94 black fill, a 72x96 or 72x92 picture window at (8,7), and a border frame whose opening is 72x92. | High / Unknown | ✔ promoted (amended) | [EXP-0410](../experiments/EXP-0410-dialogue-layout/), [EXP-0419](../experiments/EXP-0419-dialogue-render/) |
| DLG-RECT-037 | All six opening routines (seven call sites) build the identical panel, because the constructor takes only a name; every shipped node takes the portrait layout, so the text rectangle is 300x135 at (204,160). | High | ● active | [EXP-0410](../experiments/EXP-0410-dialogue-layout/) |
| DLG-LINE-038 | Each wrapped line is drawn from the rectangle's left, 17 px below the previous one, justified by widening the word gaps except on a paragraph's last line and one-word lines; a paragraph's first line is indented 10 px. | High / Medium | ✔ promoted (amended) | [EXP-0410](../experiments/EXP-0410-dialogue-layout/), [EXP-0419](../experiments/EXP-0419-dialogue-render/) |
| DLG-BUTTON-039 | The button is drawn from lines and a font-1 label with no art and no fill: a two-colour bevel, the label centred at (316,308) in gold ink (brown on hover), and a shadow offset of 2 px (4 while pressed inside). | High / Unknown | ✔ promoted (amended) | [EXP-0410](../experiments/EXP-0410-dialogue-layout/), [EXP-0419](../experiments/EXP-0419-dialogue-render/) |
| DLG-KEYS-040 | Enter, Escape and a click on the button each raise command `0x46f`, which turns the page or, on the last page, closes the panel; Space has no dialogue action, and the panel has no accept or decline state. | High / Medium / Unknown | ● active (amended) | [EXP-0410](../experiments/EXP-0410-dialogue-layout/) |

### DLG-PANEL-035

- `R0695` passes the literals 30, 120, 610, 360 to `R0696`
  (`DLG-WIN-001`). Only the width and height of that rectangle, 580x240, are
  used: the base constructor `R0758` calls `R0707`, which
  replaces the size and centres it on the screen.
- `W' = ((W - 8) / 96) * 96 + 8` and
  `H' = ((((H - 104) + a) >> 6) << 6) + 104`, with `a` = 63 when `H - 104` is
  negative and 0 otherwise. Then `left = (screenW - W') >> 1` and
  `top = (screenH - H') >> 1`. 580x240 becomes 488x232.
- `R0341` sets the screen words `L00618` and `L01259` to
  640x480, 800x600 or 1024x768 (`L03602`, `L03603`) by the option strings
  `-640`, `-800` and `-1024` it names. The panel is drawn at (76,124),
  (156,184) and (268,268), 488x232 on all three.
- The background bitmap argument is 0 (`R0759` stores it at
  `panel+0x64`), so the panel has no bitmap background. The painter
  `R0760` (`vt+0x30`) draws a frame from `interface/lm.256`
  (global `L03604`) instead.
- It clamps the rectangle to the screen and shrinks its right and bottom by
  8, giving a 480x224 body. The `lm.256` frames measure 0: 96x64, 1: 48x48,
  2: 96x48, 3: 48x48, 4: 48x64, 5: 48x64, 6: 48x48, 7: 96x48, 8: 48x48, the
  same on both roots.
  - Corners: frame 1 at (L,T), frame 3 at (R-48,T), frame 6 at (L,B-48),
    frame 8 at (R-48,B-48).
  - Edges: `(bw - 96) / 96` = 4 tiles of frame 2 at `x = L+48+96i`, `y = T`,
    and of frame 7 at `y = B-48`; `(bh - 96) / 64` = 2 tiles of frame 4 at
    `x = L`, `y = T+48+64j`, and of frame 5 at `x = R-48`.
  - Interior: frame 0 in 4 by 2 tiles from (L+48,T+48).
  - The tile rectangles cover 96 + 96*4 = 480 by 96 + 64*2 = 224.
    This is geometric coverage, not opaque coverage: the selected forward
    sprite decoder preserves destination pixels at skip runs (DIALOGUE-063).
- Frames 3, 5, 6, 7 and 8 request shadows 8 px right and down through
  `vt+0x1c` (blit mode 6). The nine shadow requests precede all 24 body
  requests. Their mask remaps existing destination pixels rather than
  filling every pixel in the band (DIALOGUE-063).
- `R0361` darkens the whole screen behind the panel once
  (`DLG-DIM-013`). The root container `campaign+0xcc` spans the screen from
  (0,0) (`R0315`), so panel coordinates are screen coordinates.

**Confidence.** High for the size, the position and the tiling: `R0707`,
`R0760` and `R0341` are read whole, and every figure is
arithmetic on the nine frame sizes measured from both roots' `lm.256`. The
rival, a panel drawn at its constructor rectangle (30,120)-(610,360), is
refuted because `R0707` replaces all four coordinates.

**Unknown.** Native table mode, clipping, final pixels and parent repaint
effects. The selected sprite wrapper passes the requested position without
adding a record origin; selected forward mode 6 is decoded (DIALOGUE-063).

**Amended.** DIALOGUE-063 replaces the former second-draw order, unqualified
band-fill wording and opaque body-coverage conclusion. Tile rectangles span
the body, while stream skips preserve destination pixels. It closes the
selected-wrapper origin and shadow-primitive questions under supplied private
state, without claiming native pixels.

### DLG-PORTRAIT-036

- `R0695` adds the portrait child `R0698` (id 12, bitmap 0,
  mode 2) at panel-relative (30,54)-(118,168), 88x114, only when `panel+0x7c`
  is set (`DLG-FACE-008`). At 640x480 it is (106,178)-(194,292).
- Its paint `R0761` first fills (L+8,T+7)-(L+80,T+101), 72x94, with
  black (`R0718`), then blits the picture surface at (L,T) keyed on
  colour 0 (`vt+0x38`). The surface is 88x108, six rows shorter than the
  child rectangle.
- `R0740` builds the surface (`UNIT-PICT-036`, `DLG-FIGURE-020`), in
  this order: cleared; `interface/t_back.bmp` (160x240, 24 bpp) copied opaque
  through the source window to (8,7); the figure canvas (160x240) copied keyed
  through the same source window to (8,7); `interface/t_border.256` frame 0
  (88x108) drawn at (0,0). The canvas is freed afterwards.
- The source window depends on the second word of the record's portrait
  quadruple at `+0x20` (`PortraitY1`, `REG-NPC-089`), tested at `L03605`.
  Not -1: (x0, 0x90 - y1)-(x0 + 72, 0xf0 - y1), 72x96, with x0 the first
  word. -1: (0x24,0x8c)-(0x6c,0xe8), 72x92.
- The border frame, measured by walking its control stream under
  `SPR256-RLE-007`, `SPR256-RLE-020`, `SPR256-STRUCT-001` and
  `SPR256-TRLR-021` (identical on both roots): 2231 control bytes, 1626
  written and 7878 unwritten pixels. The unwritten region that contains
  (44,55) has 6506 pixels and the bounding box (8,7)-(80,99), 72x92, and does
  not fill that box. The frame writes 118 pixels inside the 72x92 window, 404
  inside the 72x96 window and 262 inside the black fill.
- The border is drawn last, so the picture the player sees is the part of the
  window that the frame leaves unwritten. At 640x480 the window is
  (114,185)-(186,277) for the -1 case and (114,185)-(186,281) otherwise; the
  four extra rows of the second lie under 286 border pixels.

**Confidence.** High for the child rectangle, the fill, the surface size, the
layer order and the two windows: each is an operand of a named instruction in
`R0761` or `R0740`, both read whole. High for the border
measurement under the stated instrument (a walk of the control stream, not a
render). Which arm a shipped record takes was not classified
(`DLG-SPEAKER-022`).

**Unknown.** Which rows are finally visible from the produced 240-row canvas
(`REG-NPC-091` grades the bottom-up reading Medium for the flat arm only).
Upstream canvas contents and native orientation remain outside the suffix
and primitive controls.

**Amended.** DIALOGUE-064 establishes the selected suffix's reversal/copy order
and windows. The keyed primitive copies nonzero source words unchanged and
skips zero; it descends physical source rows as destination y increases.
These local facts close the keyed-blend question without identifying final
native visible canvas rows.

### DLG-RECT-037

- `R0695(name)` is the only constructor (`DLG-WIN-001`) and takes the
  name alone, and it shows the panel itself (`R0361`), so no caller has
  a step between construction and display. Every rectangle below is a
  constructor literal offset by the panel origin (`DLG-PANEL-035`).
- The routines, their call sites, name formats and shipped nodes. A node is a
  `main.res` text node whose base name starts with the stem, and every one
  contains `npc`, the whole-file test of `DLG-FACE-008`.

| Routine | Call site | Name format | EN nodes | RU nodes |
|---|---|---|---|---|
| `R0701`, mission event | `L03606` | `battle\m%d\event%02d` | 225 | 228 |
| `R0702`, inn NPC | `L03607`, `L03608` | `inn\NPC\npc%02dm%d` | 22 | 29 |
| `R0703`, mercenary hall | `L03609` | `inn\mercenary\npc%02d` | 14 | 14 |
| `R0704`, mercenary hall | `L03610` | `inn\mercenary\npc35` | 14 | 14 |
| `R0705`, shop keeper | `L03611` | `shop\npc31m%d` | 4 | 4 |
| `R0706`, training hall | `L03612` | `training\npc34m%d` | 3 | 3 |

- The two mercenary routines share one folder, so the 14 nodes count once:
  268 EN and 278 RU nodes in all, none without `npc`
  (`evidence/strings.tsv`). Every shipped dialogue therefore takes the
  portrait layout; the no-portrait layout below is not reached by shipped
  data. `DLG-WIN-001`'s 318/22/14/4/33 count every node under each folder
  prefix, a different population.
- The rectangles at 640x480, as left, top, right, bottom:

| Item | Rectangle | Size | Source |
|---|---|---|---|
| panel constructor | 30,120,610,360 | 580x240 | literal; size only |
| panel drawn | 76,124,564,356 | 488x232 | `DLG-PANEL-035` |
| frame body | 76,124,556,348 | 480x224 | `DLG-PANEL-035` |
| portrait child | 106,178,194,292 | 88x114 | literal 30,54,118,168 |
| portrait surface blit | 106,178,194,286 | 88x108 | `DLG-PORTRAIT-036` |
| black fill | 114,185,186,279 | 72x94 | `DLG-PORTRAIT-036` |
| picture window, -1 / other | 114,185,186,277 / 114,185,186,281 | 72x92 / 72x96 | `DLG-PORTRAIT-036` |
| text control, constructor | 204,160,504,296 | 300x136 | literal 128,36,428,172 |
| text control, live | 204,160,504,295 | 300x135 | base constructor |
| text control, no portrait | 124,160,504,295 | 380x135 | literal 48,36,428,172 |
| button | 276,296,356,322 | 80x26 | literal 200,172,280,198 |
| button label cell, EN / RU | 305,300,327,315 / 281,300,351,315 | 22x15 / 70x15 | `DLG-BUTTON-039` |

- At 800x600 and 1024x768 every rectangle moves by (80,60) and (192,144),
  the shift of the panel origin (`evidence/layout.tsv`).
- The text control's base constructor `R0762` recomputes the height
  from the font: `n = floor(136 / (h + 4))` rows of `h + 4`, plus 2. With
  font 1 (`h` = 15) that is 7 rows and 7 * 19 + 2 = 135. `R0699` then
  sets the pitch to `h + 2` = 17, and the control shows
  `min(floor(135 / 17), line count)` = min(7, line count) lines.

**Confidence.** High. The literals are `PUSH` operands read in one routine;
the derived values are arithmetic on the snap, the frame sizes and the font-1
metrics measured on both roots (`evidence/metrics.tsv`). The stem counts are
directory counts of both roots' `main.res` over the five folders named, so
"every node has `npc`" is a census of those nodes and not a statement about a
node that does not ship.

### DLG-LINE-038

- `R0763` paints the text control through
  `R0764(font, rect, first, last, lines, ink, pitch)`: `first` is the
  control's top line (0 for every shipped block), `last` is `first` plus the
  visible count (`+0x8c`, at most 7), and `pitch` is 17. The wrapped lines
  come from `R0725` through `R0726` at the rectangle's full
  width, 300 with the portrait (`DLG-WRAP-009`). Ordinary fitting nonempty
  pieces keep a trailing space, and their last line ends in CR. The empty
  and over-wide controls have exceptions (DIALOGUE-060).
- For line `i`, counted from 0 over the whole array of `n` lines:
  - the line is a paragraph's first when `i = 0` or line `i - 1` ends in CR,
    and then `p` = 10, the font-1 `.dat` dword 32; otherwise `p` = 0;
  - it is justified when `i != n - 1` and it does not end in CR;
  - the drawn copy has its CR removed;
  - `x0 = rect.left + p` and `y = rect.top + 17 * (i - first)`.
- A justified line goes to `R0765(x0, y, W, line, ink)` with
  `W = rect.width - p`. It trims the right end, then repeatedly trims the left
  end and cuts a word at the first space, so runs of blanks collapse, and it
  sums the words' widths `m` (`R0766`). With two or more words
  `gap = (W - sum) / (words - 1)` in floating point.
  - Word `k` is drawn at `trunc(xacc)`, `xacc` starting at `x0` and becoming
    `(xacc + m_k) + gap` after each word, so the first word starts at `x0` and
    the last ends at about `x0 + W`. The natural space width is not used.
  - A one-word line is drawn at `x0`.
- Every other line is drawn whole at `(x0, y)`, left aligned: a paragraph's
  last line (which includes the line before an authored break) and the last
  array line. No line is centred.
- Every draw goes through `R0571(x, y, text, 0, ink, 1)`, which calls
  the font's draw slot `vt+0x14`, `R0767`, twice: first at
  `(x + 1, y + 1)` with the shadow ramp `L03613` (flat 8,8,8), then at
  `(x, y)` with the ink ramp `L03614`, whose level 15 is (255,255,255)
  (`MISSION-MSGLINE-056`). The first line's top edge is the rectangle top:
  no ascent offset is added.
- At 640x480 with the portrait, line tops are 160, 177, 194, 211, 228, 245
  and 262; the left is 204, or 214 on a paragraph's first line; the justify
  width is 300, or 290 on that line (`evidence/text.tsv`).
- Shipped population, both roots' five families, from a transcription of the
  splitter, the wrapper and the measure (`tools/dlglayout`, no emulation):
  EN 688 blocks and RU 732, one per tag containing `part=`. The 535 EN and
  543 RU mission event blocks equal the tag-body census of `DLG-TAGARM-027`.
  - Lines per block, 1 to 7: EN 48, 154, 172, 104, 76, 64, 70; RU 74, 171,
    145, 115, 87, 76, 64. No block exceeds 7 lines.
  - Authored breaks (a CRLF inside a block): EN 0 blocks; RU 12 (7 mission
    event, 4 inn, 1 shop).
  - No block has a one-word line that is not a paragraph's last, an unclosed
    line wider than 300, an empty text, a glyph outside `font1.dat`, or a
    wrapper stall, and no block contains a `~` byte, so the tilde markup
    (`TEXT-079`) never reaches a shipped dialogue.

**Confidence.** High for the placement rule: `R0764`, `R0765`
and `R0571` are read whole, and each operand is traced to its source.
The alternatives of one left x for every line and of centring are excluded by
the justify path and by the absence of any centring flag in the call. Medium
for the population figures: the wrapper is transcribed from its listing and
checked by hand traces, not emulated or observed.

**Unknown.** Native x87 precision and rounding mode. DIALOGUE-061 establishes
double stores between words and two x87 additions before each store.
DIALOGUE-062 measures zero coordinate differences between nearest-even
PC53/PC64 spilled models and per-add double over the stated 20965 word
positions; this corpus agreement does not select native floating-point state.

**Amended.** DIALOGUE-060 narrows the unconditional trailing-space/CR rule.
DIALOGUE-061 closes the retained-accumulator operation-order alternative;
DIALOGUE-062 reproduces the named census through original wrapping instructions
with declared services. The census remains Medium and does not prove native
reachability or visible words.

### DLG-BUTTON-039

- The button is `R0700`: id 11, panel-relative (200,172)-(280,198),
  80x26, font 1 (`[L03615]`), command `0x46f`, and the label is
  string-table entry 77 (`DLG-LANG-010`): "Ok" in EN, 22 px wide, and
  "Принять" in RU, 70 px wide. It has no `~` accelerator in either root.
- `R0768` paints it from primitives. There is no art and no interior
  fill: the parent repaints its frame under the rectangle first, then the
  label is drawn, then the bevel.
- The two colours are RGB (41,69,63), light, and RGB (7,12,9), dark, each
  channel reduced to the screen format's width and shift. With `R' = R - 1`
  and `B' = B - 1`, lines run through `R0769` (both ends drawn) and
  pixels through `R0770`:
  - colour 1: vertical `x = R'`, `y = T+2..B'-2`; vertical `x = R'-1`,
    `y = T+1..B'-1`; horizontal `y = B'`, `x = L+2..R'-2`; horizontal
    `y = B'-1`, `x = L+1..R'-1`; pixel (R'-2,B'-2);
  - colour 2: horizontal `y = T`, `x = L+2..R'-2`; vertical `x = L`,
    `y = T+2..B'-2`; pixel (L+1,T+1).
- Idle, colour 1 is dark and colour 2 light, which is a raised bevel with a
  one-pixel light top and left and a two-pixel dark bottom and right. Pressed,
  which needs the pressed flag (`+0x6c`, set by the mouse press) and the
  cursor inside the rectangle, colour 1 is light and colour 2 dark.
- The label is drawn through `R0571` at
  `(L + (R' - L) / 2 + 1, T + (B' - T) / 2)` = (316,308) at 640x480, with the
  anchor flags 0xa (`TEXT-078`), which put the cell at (305,300)-(327,315) in
  EN and (281,300)-(351,315) in RU. The shadow offset is 2 idle/outside and 4
  while pressed with cursor membership true.
- The ink ramp is `L03616` idle, level 15 = (185,159,73), and
  `L03617` while the hover flag (`+0x68`, set by the mouse move
  `R0771`) is set, level 15 = (150,90,0). The shadow ramp is
  `L03613`, flat (8,8,8). The ramp entries are from CPU emulation of the
  executable's own ramp builder `R0572` over synthetic memory
  (`evidence/ramps.tsv`); the original was not run.
- The paint ends by testing flag 1 of the control word `+0x18` (`vt+0x20`,
  `R0772`, an AND). With it clear the routine remaps the button's
  screen rectangle through `R0732` at level 3 (`L03347`), the shade of
  `DLG-DIM-013`. The word starts as 3: the base constructor stores 1
  (`R0773`, `L03618`) and the button constructor ORs 2 (`L03619`).
  `R0695` appends the button and shows the panel without another write,
  so the dialogue button is drawn undimmed. No routine read here clears flag 1.

**Confidence.** High for the geometry, the colour assignment, the label
anchor and the offsets: every operand is read in `R0768`, listed whole.
High for the ramp values under the stated instrument.

**Unknown.** Native screen-format selection, event delivery, parent repaint,
glyph pixels and cadence. The selected builder and bevel are now executed
under supplied RGB565/RGB555 fields (DIALOGUE-065).

**Amended.** DIALOGUE-065 closes quantization for the two supplied layouts.
Pressed presentation requires the pressed flag and cursor membership; the
hover field independently selects the ramp. Level-3 disabled remapping follows
painting; its full/reduced arithmetic is DIALOGUE-055.

### DLG-KEYS-040

- A key reaches the panel through the campaign window's key handler
  `R0231` (message-map record for `0x100`). Enter and Space take its
  forwarding arm `L03620`. Escape takes arm `L03621`, which posts
  `0x416` only when `campaign+0x3dc` is 1 and `0x41f` only when it is 0 with
  `campaign+0x3b4` empty; a panel on screen sets bit 3, so Escape forwards as
  well (`MENU-ESC-010`, `MENU-INPUT-016`). `R0774` forwards `WM_CHAR`
  the same way. The arm calls `vt+0x48` of the root container `campaign+0xcc`.
- The class's show slot `vt+0x80` (`R0710`) calls the base show
  `R0775` at `L03622`. That routine calls `vt+0x24(1)`, which is
  `R0776` in the base and dialogue class tables (`L03623`,
  `L03624`): it saves the root's capture object at `root+0x40` and stores
  the panel in `root+0x34`. `R0390` gives a mouse message
  (`0x200`..`0x206`, `0x400`) to that object before any child
  (`L03625`..`L03626`), so the panel receives the mouse while it is up.
- The root offers the key to its focus object, `root+0x38`, which is the
  panel after `R0775` (`vt+0x28(1)`, `R0777`). The panel's
  `vt+0x48` is `R0715`; a message other than `0x46f` goes through
  `R0716` to the container dispatcher `R0390`, which offers a
  key to the focused child `panel+0x38`, then to every child in order
  (portrait, text control, button), then to the panel's own key slot
  `R0714`.
- Enter (`0x0d`): the button's key handler `R0713` answers in the
  second step. It posts `0x46f` to the main window and returns 1, so the
  panel's own arm never sees the key.
- Escape (`0x1b`): no child takes it, and `R0714` calls
  `vt+0x48(0x46f)` on the panel directly and returns 1. That routine has the
  same arm for `0x0d`; every other key goes to the base handler
  `R0717`.
- A click on the button: the press (`R0778`) sets the pressed flag,
  the release (`R0779`) clears it and, with the point inside the
  rectangle, posts `0x46f`. The main window's default arm (`L03627`)
  forwards a posted `0x46f` to the root, and `R0715` answers it.
- Space (`0x20`) has no dialogue action. The text control takes only
  `0x21`, `0x22`, `0x26` and `0x28`; the button takes only `0x0d`;
  `R0714` takes `0x0d` and `0x1b`; the base handler's table covers
  `0x09` and `0x25`..`0x28`. The button's `WM_CHAR` handler `R0780`
  fires only on its accelerator `+0x74`, which is 0 for label 77.
- `0x46f` in `R0715` calls the pager `R0711`, which adds 1 to
  `panel+0x68` and looks up that part (`R0712`, `DLG-MARKUP-007`).
  - Found: the text of child 10 is replaced (`R0725`), children 10
    and 12 repaint, and the panel stays open. A dialogue of N parts takes N
    presses.
  - Not found, the last page or a suppressed part (`DLG-TAGARM-027`): when
    `panel+0x80` is nonzero the panel posts `0x45b` to the main window with
    that value, then calls `R0716(0x445)`, which hides the panel
    (`vt+0x84`) and posts `0x44c`, whose arm is `L03311` (a call of R0709)
    (`DLG-LIFE-005`).
- The `tips=` arm of the pager (`L03628`..`L03629`) reads the five
  characters after `tips=` in the part's tag with `%d` into `panel+0x80`.
  The `0x45b` arm of the main window (`L03630`) does nothing while the
  global `L03631` is 0.
  - Census over the five families' shipped blocks (`evidence/wrap.tsv`):
    `tips=` is in 12 EN and 12 RU mission event blocks, with the numbers 1 to
    7, and in no inn, mercenary, shop or training block on either root. No
    shipped town dialogue sets `panel+0x80`, so the `0x45b` post exists only
    at the close of some mission event dialogues.
- A key on the focused text control: `R0731` acts on `0x21`
  (`R0781`), `0x22` (`R0782`), `0x26` (`R0783`) and
  `0x28` (`R0784`) when focus flag 4 is set. Each moves the top line
  through `R0730`, which clamps it to `[0, lines - visible]`, then
  sends `0x46d` to the panel and returns 1; `0x23`, `0x24`, `0x25` and `0x27`
  return 0. With 7 or fewer lines the range is 0, and no shipped block has
  more (`DLG-LINE-038`).
- On the last page of a quest offer Enter, Escape and the click do the same
  close, so the alternative of Escape declining is refuted. The panel has
  three children and no accept state: its fields are the page `+0x68`, the
  name `+0x6c`, an object pointer `+0x70` that the pager deletes and replaces,
  the payload `+0x74` and its size `+0x78`, the portrait flag `+0x7c` and the
  tips value `+0x80`, written by the constructor, the loader `R0720`,
  the pager and the part lookup `R0712`, and by neither key handler.
  Registration does not wait for the close:
  the shop calls `R0785` at `L03632` and the training hall at
  `L03633`, right after the constructor returns, and the inn appends the
  mission to its queue at `L03607` for `R0786` to commit when the
  player leaves (`REG-SCN-062`, `REG-SCN-064`, `DLG-ZEROARM-029`).

**Confidence.** High for what each handler does with each key and for the
command flow: every arm is a named instruction in a routine read whole, and
the answer to the last-page question follows from one command. Medium that
the button rather than the panel answers Enter (it rests on the dispatcher's
child order, derived by reading and not observed, and both arms send the same
command), that the text control holds focus at open (its base constructor sets
flag 2, the portrait lacks it, and the show routine focuses the first child
with flags 1 and 2, `R0787`), and that the posted `0x46f` reaches the
panel before another child of the root.

**Unknown.** Whether siblings of the panel under the root react to Space; the
handlers for key release (`0x101`), system keys and the panel's own `0x101` and
`0x102` slots, which were not read; what a click outside the button does once
the panel holds the capture (the panel's mouse slots were not read); senders of
`0x445` and `0x446` outside this class, which were not enumerated; what the
`0x45b` arm shows and where `L03631` is set.

**Amended.** DIALOGUE-044 and DIALOGUE-045 establish the named captured mouse
slots and outside-button release path. Native event ordering remains Unknown.


## Original backdrop operation and transitions

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DIALOGUE-054 | The reached mission and five town dialogue entry bodies share show R0361, which requests one level-3 in-place backdrop remap after setup, regardless of first-part return or existing dialogue state bit. | High | ✔ promoted | [EXP-0418](../experiments/EXP-0418-dialogue-backdrop/) |
| DIALOGUE-055 | Original level-3 full-table RGB565/RGB555 replay applies floor(channel*13/16); reduced lookup samples blue at 4,12,20,28 first, so black maps to pixel 3 rather than remaining black. | High | ✔ promoted | [EXP-0418](../experiments/EXP-0418-dialogue-backdrop/) |
| DIALOGUE-056 | Backdrop remap R0732 intersects its half-open requested rectangle with the incoming signed clip and addresses 16-bit pixels by y*bytePitch+2*x; show requests the screen right/bottom extent. | High | ✔ promoted | [EXP-0418](../experiments/EXP-0418-dialogue-backdrop/) |
| DIALOGUE-057 | The selected dialogue advance and default-close handlers repaint children or request parent paint; neither directly repeats the backdrop remap nor restores a saved surface, while a second explicit show compounds the remap. | High | ✔ promoted | [EXP-0418](../experiments/EXP-0418-dialogue-backdrop/) |

### DIALOGUE-054

Seven builder calls in six reached bodies share `L03358 -> R0361`:
mission `L03606`; mercenary `L03609`, `L03611`; inn `L03607`,
`L03608`; shop `L03610`; training `L03612`. This is the selected
population, not a whole-image reachability claim.

Show stores state at `L03330` after setting bit 0x8 for an ordinary pointer. Equality
with `campaign+0x3a8` selects setting bit 0x8000 instead; that special pointer is not
the dialogue route. It attaches, calls slots `+0x78` and `+0x80`, locks and
calls `R0732(0,0,screenRight,screenBottom,3)` at `L03338`. Unlock and
presentation precede slot `+0x34` paint. Original table `L03624` supplies
the slots; `+0x80` calls the first-part routine and capture setup, with no
return test in show before the remap.

**Confidence.** High for this source sequence. Fresh repaired Ghidra 12.1.2
bytes match the owned image's initialized .text and .rdata sections. Direct
transfers and instruction starts are bound to installed bytes. Original show
replay with private surfaces distinguishes states 0,1,8,9 and parser returns
0/1; no branch suppresses its remap. Setup, OS and paint calls are explicit
stubs. Flag state alone therefore does not select a separate shading mode.

**Unknown.** Native reachability beyond the selected bodies, actual first
part contents, OS lock validity and presented output. EN/RU code identity is
one observation, not two independent confirmations.

### DIALOGUE-055

The instrument executes original `R0788` and `R0732`, with only memory
information and allocation stubbed for table construction. Returned memory
information field `+8` at 23999999 selects reduced stride 8192; 24000000
selects full stride 65536 (`L03634..L03635`). The later pixel split reads
`L03346` and chooses `pixel >> 3` or `pixel`.

Full RGB565 has 65536 valid inputs; RGB555 has 32768 with bit 15 clear. Level
3 replay agrees with per-channel truncation on every full-table input. The
reduced builder starts the low channel at 4 and increments by 8
(`L03636..L03637`, `L03638`); lookup groups eight source values. It
differs from exact gain on the source channels at 57344/65536 RGB565 and
28672/32768 RGB555 values. In either reduced layout inputs 0..3 become 3,
3 stays 3, and 4..7 also become 3. The pixel table gives the full counts.

**Confidence.** High for these packed layouts and parameterized modes. The
entire valid input populations execute original loops; a full-table-only
model fails the reduced-mode controls. Truncation, nearest rounding, a 13/17
gain and multiplying the packed word are distinct models. Channel 7 maps to
5, channel 31 to 25; black maps to 0 in full mode and 3 in reduced mode.

**Unknown.** The native memory-information return and active mode, other
channel layouts, RGB555 inputs with bit 15 set, malformed geometries and
perceived brightness after display conversion. Packed-value comparisons do
not measure display luminance. DLG-DIM-013's universal arithmetic is narrowed.

### DIALOGUE-056

Show reads right/bottom from `L00618/L01259` and passes left/top zero.
`R0732` clamps x to `L01503/L01505`. The y loop starts at requested
top, skips rows below `L01504` or at/above `L01506`, and stops before
requested bottom. Nonpositive x width or reversed/empty y produces no write.
Rows advance by `L01511` bytes; pixels use two bytes.

**Confidence.** High for the named instruction body. Twenty-four original
replays cover six rectangle controls in four layout/mode pairs. A 9x6
surface has eight padding bytes per row and a 64-byte trailing guard. Full,
interior, crossing, empty-x, reversed-y and outside requests all match the
half-open intersection. Every untouched pixel, padding byte and guard byte
is compared. A 9x6 show replay with incoming clip `(2,1)-(7,5)` changes
exactly 20 pixels despite requesting `(0,0)-(9,6)`. This excludes panel-only,
inclusive-end and width-based pitch alternatives in the tested geometry.

**Unknown.** Incoming clips on native town/mission routes and surface-loss
behavior. The primitive does not reset its clip. No claim covers arbitrary
coordinates, pitches or unsupported pixel indices.

### DIALOGUE-057

Original show replay calls the remap a second time for an explicit second
show while mask 0x8 is set. This is a control, not evidence that normal page
advance invokes show twice. At command `0x46f`, `R0715` calls `R0711`.
Nonzero selects paints for children 10 and 12; zero reaches `R0716`, whose
presence-flag arm invokes slot `+0x84` and posts `0x44c`.

Default close `L03359..L03639` clears mask 0x8, conditionally resets the
clock, destroys the panel, sends `0x408` and calls parent slot `+0x34` unless
remaining state is 1 with `map+0x80 == 0`. No direct remap or saved-surface
restore occurs in these selected bodies. The replay covers both parser
branches and both parent-paint conditions with explicit painter, parser,
message, clock and destructor stubs.

**Confidence.** High for the selected source handlers and replay controls.
The no-direct-remap clause is confined to their complete listings. It is not
a transitive or image-wide absence claim. The source and replay exclude
automatic remapping inside this command branch and idempotence inside show.

**Unknown.** Child/parent painter side effects beyond these handlers, native
event ordering, first visible frame, persistence between frames and final
closed-screen pixels. Repaint dispatch does not prove repaint completion.

## Open questions

- Which 96 rows of the 240-row composition the portrait pane frames
  (`DLG-FIGURE-020`). `R0740` blits a 72 x 96 window of the
  160 x 240 canvas, or 72 x 92 on the default arm, to `(8,7)` of the 88 x 108
  surface the panel shows, with source top `0x90 - PortraitY1` or the default
  `(0x24, 0x8c)-(0x6c, 0xe8)` (`REG-NPC-089`, `DLG-PORTRAIT-036`). Reading
  those canvas rows as picture rows needs the canvas's row order. `REG-NPC-091` grades the bottom-up reading Medium on the
  flat arm, from 33/33 renders judged by eye, and does not address the composed
  arm, whose filler is the sprite blitter rather than a DIB loader. Bottom-up
  frames the top of the figure; top-down frames the lower legs; the two are up
  to 122 rows apart. The measurement: whether the loader that fills the canvas
  from a BMP walks rows forwards or backwards, and whether `R0789`'s
  colour blit walks them the same way.
- Which arm a given shipped dialogue takes at run time (`DLG-SPEAKER-022`,
  `DLG-DRESS-024`). `DLG-SPEAKER-022` establishes both arms, and
  `DLG-DRESS-024` traces one section to a live placement. The other 28
  composed-figure tags depend on whether an ordinary map-placed human satisfies
  the section's predicate in that mission, which is per-map state. The
  measurement: for each tagging mission, resolve its `.alm` humans' `+0x18c`
  bits and face bytes against the tagged sections' `Flags` and `Face`.
## Captured mouse input

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DIALOGUE-044 | The captured dialogue panel routes mouse input to its children; an outside left press returns 0 without a pager or close command, and its root does not retry sibling hit tests. | High / Medium | ✔ promoted | [EXP-0415](../experiments/EXP-0415-dialogue-queue/) |
| DIALOGUE-045 | The dialogue button posts its command only for a release inside its rectangle after a press; a release outside clears the press and posts no command. | High | ✔ promoted | [EXP-0415](../experiments/EXP-0415-dialogue-queue/) |

### DIALOGUE-044

The constructor installs table `L03624`. Show `R0775` calls slot `+0x24`,
`R0776`, which stores the panel in its parent capture pointer `+0x34` and
saves the old pointer at `+0x40`. The campaign root is constructed by
`R0773` with table `L03640`.

`R0390` gives messages `0x200..0x206` and `0x400` to capture first. After
that call it jumps to `L03641`; a zero return runs the root's own mouse slot,
not `R0642`'s sibling loop. The root's down/up/double/right slots return 0.
This is return-value propagation to the parent dispatcher, not a panel call
into the parent or a second hit test behind the panel.

The dialogue's own slots `+0x4c/+0x50/+0x58/+0x5c/+0x60/+0x64/+0x68` return 0.
Its left-down slot `+0x54` is `R0790`. When panel `+0x40 == +0x34`, it clears
`+0x40`, recursively hit-tests children through `R0791`, and calls a found
child's `+0x54`. No hit or a hit on the panel itself reaches `L03642`, which
returns 0. A point outside all child rectangles is also excluded from the
preceding `R0642` child-message pass by `L03643`'s `PtInRect` call.

Text child down returns 0. Its up sends `0x472` and double click sends `0x444`
to the panel; neither is `0x46f`, and neither enters the panel's `0x445/0x446`
close arm. Its drag route scrolls only with mouse flag 1 and child `+0x90`
nonzero. Scroll helpers are reused from `DLG-KEYS-040`, not reread here.

**Confidence.** High for the named slots, branches, returned values and absence
of pager/close commands in these bodies. The instrument reads 50 selected
functions, raw PE vtable dwords and independent branch destinations; this is
not a whole-image absence claim. Medium for the composed campaign input route:
the capture/root dispatch is static and assumes the shown root/panel identity.

**Unknown.** Native input ordering, actual OS events, capture changes by outside
callers, window teardown during dispatch and panel types with another vtable.
The text scroll helper bodies and drawing/sound callees are outside this read.

### DIALOGUE-045

Button table `L03644` maps `+0x54` to `R0778` and `+0x58` to `R0779`.
Down requires a parent and enabled flag 1. With pressed field `+0x6c` zero it
calls `L03645(1)`. That routine sets `+0x6c = 1` and, when flag 8 is clear,
requests capture in the button's parent, the dialogue panel.

Up has the same parent/enabled prerequisite. With `+0x6c != 0` it calls
`L03645(0)`, clearing the flag and restoring capture when flag 8 was set.
Only then does `L03646` test the release point against the screen rectangle
through `L03643`; its false branch skips `L03647..L03648`, the command
post. No pressed flag means no post even for a release inside. The dialogue
constructor binds command `0x46f`; that command's pager/close behavior is
`DLG-KEYS-040`.

**Confidence.** High for the two handler bodies, flag stores, capture calls and
release hit-test branch. These are checked original instructions, not a native
click experiment. The rectangle's right and bottom edges use `PtInRect`.

**Unknown.** Native delivery after a press, arbitrary capture replacement,
disabling or deleting the button between events, and sound/paint side effects.

## Shop and training offer first part

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DIALOGUE-048 | Reached shop and training offer tails test their array count, name its head in the dialogue builder, then register and remove that head after normal return without testing the builder result. | High | ● active | [EXP-0417](../experiments/EXP-0417-offer-first-part/) |
| DIALOGUE-049 | On a constructed dialogue panel, a failed first-part lookup leaves the default text and still calls base show; the next pager call tries part 2 and can continue rather than close. | High | ● active | [EXP-0417](../experiments/EXP-0417-offer-first-part/) |
| DIALOGUE-050 | Dialogue tag rejection resumes the same-part scan; portrait and tips fields can be changed before a parser return of 0, with no rollback in the reached body. | High | ● active | [EXP-0417](../experiments/EXP-0417-offer-first-part/) |
| DIALOGUE-051 | Each EN/RU shop and training offer node in the measured MAIN.RES trees has one unconditional part-1 candidate; none of those 14 nodes represents an absent or conditionally rejected first part. | High / Medium | ● active | [EXP-0417](../experiments/EXP-0417-offer-first-part/) |

### DIALOGUE-048

Shop show `R0705`, table `L03649+80`, reaches an offer-count test through
`L03650` and `L03651`. Its array header is `campaign+630`, data pointer
`+634`, count `+638`. A positive count reads the head word, builds
`shop\npc31m<word>` and calls `R0695` at `L03611`. After normal return it
calls registration `R0785` at `L03632`, then `004570ad5(array,0,1)` at
`L03652`. There is no builder-result branch between these calls.

Training show `R0706`, table `L03653+80`, reaches the count test for
`campaign+64c` from `L03654`. Its array header is `+644`, data pointer `+648`. A positive count
reads the head word, builds `training\npc34m<word>` and calls the same builder
at `L03612`. Normal return reaches registration at `L03633`, then removal
of element 0 with count 1 at `L03655`, with no builder-result branch.

The eight original-tail controls vary count 0/1 and a synthetic builder return
0/1 for each caller. Count 0 emits none of the three calls. Count 1 emits
builder, registration and removal in that order for either return value.
Mission word 31 is supplied to both tails; training word 31 is a synthetic
control, not a shipped training node. Existing REG-SCN-064 (amended) and
REG-SCN-065 describe the registration consequences separately.

**Confidence.** High for these reached tails, their count/head reads and their
local order after a normal builder return. Raw listing rows and direct branch
targets are bound to the input executable. Original tail instructions run;
builder, registration and removal are explicit host boundaries in the probe.

**Unknown.** Earlier room initialization, native room entry, actual array
contents at entry, construction failure or resource exceptions, asynchronous
delivery and downstream registration/removal effects under this instrument.
An absent text part is not treated as a missing resource or a thrown exception.

### DIALOGUE-049

The panel constructor `R0696` initializes current part `+68`, voice `+70`
and tips `+80` to zero. Loader `R0720` stores payload at `+74`, allocates its
length plus one at `+78` and writes the final NUL. The builder `R0695` adds
text child 10 containing `"Nothing to say"`, button 11 sending `0x46f`, and
portrait child 12 when `+7c` is nonzero (DLG-WIN-001, DLG-RECT-037).

Builder handoff `R0361` calls virtual `+80`, panel slot `R0710`. That
slot calls pager `R0711` at `L03656` and calls base show `R0775` at
`L03622` without testing the pager return. The pager increments `+68` before
parser `R0712`. A zero parser return branches at `L03657` before the text,
voice, portrait and tune updates. Thus an existing readable payload yielding
no accepted part 1 leaves the constructed text and still reaches base show.

In original-code controls with either no part-1 tag or a conditionally rejected
part-1 tag, show advances to 1 and reaches only the hooked base-show boundary.
A following pager call advances to 2 and accepts the synthetic body
`B\r\n`, reaching text, voice and portrait boundaries. The voice filename
argument is `speech\shop\npc31m31p2.wav`. A payload with neither part reaches
return 0 on the second call instead. These discriminate continuation from an
automatic first-failure close.

Command `0x46f` in `R0715` calls the pager, tests its result, and on failure
sends campaign message `0x45b` when `+80` is nonzero, followed by base message
`0x445`. Failure of the first lookup is therefore different from failure of
the next lookup through this command. This corrects DLG-EMPTY-004's former
automatic button-close clause.

**Confidence.** High for the constructor stores, named call order, omitted
result test and executed synthetic continuation. Text, voice, portrait and
base-window callees are hooked boundaries, not observed output.

**Unknown.** Native visibility, text pixels, focus/input delivery, real voice
loading/playback, portrait construction and exceptions. The full builder and
earlier room code were read but not executed end to end. Allocation success is
an instrument premise; explicit `sound=` routing is outside the probe.

### DIALOGUE-050

Parser `R0712` searches tag substrings for `part=<current>`. A condition
failure branches to `L03579` and resumes the scan from `L03580`, so another
tag declaring the same part can succeed (DLG-TAGARM-027). The synthetic
same-part control rejects its first candidate and accepts its second.
`part=10` also satisfies the part-1 substring search; this is a separate
synthetic positive control, not a shipped malformed tag.

The loader's whole-payload `npc` test sets `+7c`; the candidate's `npc=` test
overwrites it at `L03658` before the player conditional gates. A rejected
candidate with no `npc=` clears the flag even when another part put `npc` in
the payload. A candidate rejected after `npc=31` retains the flag and output
speaker 31. Neither change is rolled back on the parser's zero-return path.

After the conditional gates, `tips=` stores the parsed integer at panel `+80`
before the header-LF scan. A synthetic accepted-condition tag with `tips=7`
and no following LF returns 0 but leaves `+80=7`. Show leaves the default text;
a following failed `0x46f` call reaches messages `0x45b` and `0x445`. Controls
with tips 0 and 7 distinguish the conditional extra message. This says nothing
about what the `0x45b` receiver presents.

**Confidence.** High for these stores, same-part retry and bounded original
parser/pager controls. CString, registry, actor lookup and sscanf operations
are declared host hooks. Start, resolution and sex/class bits are supplied;
the Start/resolution control is not an independent native actor observation.

**Unknown.** Native reachability of each supplied actor state, arbitrary tag
combinations, real message delivery, explicit `sound=` and receiver effects.
No universal absence of other writers or caller-side rollback is claimed.

### DIALOGUE-051

The input reader enumerates every MAIN.RES node in the preserved EN and RU
roots, selecting `text/shop/npc31m<digits>.txt` and
`text/training/npc34m<digits>.txt`. Each root has four shop nodes (31, 91, 110,
130) and three training nodes (61, 121, 131). All 14 selected payloads contain
one part-1 candidate, none carries a first-part conditional, and all contain
the whole-payload `npc` substring. The census does not cover other families,
other installs, loose overrides or custom resources.

The original parser accepts part 1 for all 448 vectors: 14 payloads times
four supplied hero bit patterns times four supplied speaker bit patterns times
two coupled Start/resolution gates. This confirms the selected first tags do
not depend on those controls; it does not enumerate actual actors or all
independent combinations of Start and resolution.

**Confidence.** High for the archive census, lengths, tag counts and digests.
Medium for the corpus parser projection under the declared hooks. Corpus
agreement does not prove a native panel was shown. EN/RU executable equality
is one code observation, not an independent replication.

**Unknown.** Native content overrides, offer-array reachability and renderer
or input causes of an apparently absent panel. These stock payloads cannot
distinguish native failure models; the synthetic controls establish the local
missing/rejected-part mechanism separately.

## Selected rendering instructions and private controls

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DIALOGUE-060 | Selected original wrapping uses strict fitting, keeps an over-wide first word whole and splits paragraphs on CRLF; exact-width and bare-CR controls can repeat an unchanged remainder under declared services. | High | ✔ promoted | [EXP-0419](../experiments/EXP-0419-dialogue-render/) |
| DIALOGUE-061 | Selected justification stores the gap and accumulator as doubles, adds integer width to the accumulator then gap before each spill, and truncates draw x with a restored control word; native FPU state is Unknown. | High | ✔ promoted | [EXP-0419](../experiments/EXP-0419-dialogue-render/) |
| DIALOGUE-062 | Over 688 EN and 732 RU selected blocks, original-hooked wrapping yields at most 7 lines; nearest-even PC53/PC64 spilled models agree with per-add double at 20965 justified word coordinates. | Medium | ✔ promoted (amended) | [EXP-0419](../experiments/EXP-0419-dialogue-render/), [EXP-0421](../experiments/EXP-0421-dialogue-input/) |
| DIALOGUE-063 | At each supplied panel geometry, the selected painter requests 9 shadows before 24 body sprites; forward mode 6 remaps destination pixels through its mask, without adding a sprite-record origin. | High | ✔ promoted | [EXP-0419](../experiments/EXP-0419-dialogue-render/) |
| DIALOGUE-064 | The selected portrait suffix places its 72x92 default or 72x96 metadata window at (8,7); original copy controls descend physical source rows, copying nonzero words unchanged and skipping zero on the keyed path. | High | ✔ promoted | [EXP-0419](../experiments/EXP-0419-dialogue-render/) |
| DIALOGUE-065 | With supplied RGB565/RGB555 and cursor state, the selected OK painter keeps anchor (316,308), selects hover ink independently, and swaps bevel colors and shadow 2 to 4 only when pressed and inside. | High | ✔ promoted | [EXP-0419](../experiments/EXP-0419-dialogue-render/) |

### DIALOGUE-060

The selected wrapper `R0727`, paragraph splitter `R0792` and combined
entry `R0726` run as original instructions in private memory. CString
construction, slicing, trimming and line-array storage are declared service
hooks. Width measurement `R0766`/`R0793` runs original instructions.
Allocation and C whitespace classification are supplied.

Fourteen controls use synthetic advances 8, spacing 2 and the original extra
space width 7. A measures 10, A-space 27 and A-space-B 37. A at width 10
and A-space-B at width 27 revisit the same unchanged remainder. The complete
37-wide A-space-B at width 37 becomes two lines, excluding inclusive fitting.
An A wider than width 9 stays whole without a trailing space or CR. The
first word AAAA at width 15 also stays whole rather than splitting by glyph.

CRLF separates pieces. LF alone remains data. Bare CR and trailing bare CR
repeat unchanged remainder. Leading CRLF produces an empty first line;
empty input produces no line and spaces alone produce one empty line.
Ordinary fitting pieces receive a trailing space and their final line CR.

**Confidence.** High for the selected comparisons and finite controls under
these services. Instruction starts and direct targets are bound to the
preserved image. The exact-width, over-wide and three newline controls
discriminate strict/inclusive fitting, word/glyph splitting and CRLF/any-newline
models. A repeated remainder is detected at the second equal loop entry,
not interpreted as a completed native hang.

**Unknown.** Native locale, allocator behavior, arbitrary malformed bytes,
pointer faults, termination outside this finite population and native pixels.

### DIALOGUE-061

In `R0765`, the integer residual width divided by the word-count minus one
is stored as a double at `L03659`. Initial integer x is stored as a double
at `L03660`. Each word loads its integer width, adds the stored accumulator
at `L03661`, adds the stored gap at `L03662` and stores the accumulator
as a double at `L03663`. The two additions precede this store.

Conversion `R0279` saves the incoming x87 control word, selects truncation,
converts to an integer and restores that word. Thus truncation of draw x is
separate from precision and rounding of the preceding arithmetic.

Twenty-four supplied width vectors run at nearest-even PC24, PC53 and PC64,
72 original-instruction controls. Every x equals the corresponding binary
rational instruction model and the incoming control word is preserved. For
x=214, width=290 and widths 10,20,30,40, spilled PC64 yields last x=463;
retaining the accumulator yields 464. For x=159, width=1469 and widths
61,204,61,52,33,69, spilled PC64 yields last x=1559; per-add double yields
1558. A negative-x vector gives the opposite one-pixel distinction.

**Confidence.** High for the selected operation sequence and bounded controls.
Original double stores exclude a retained accumulator. The distinguishing
positive and negative controls exclude per-add double as an unconditional
replacement for the two-addition PC64 sequence. Supplied width, draw and
string services do not establish native font or floating-point state.

**Unknown.** Native precision/rounding selection, exceptional values, arbitrary
integer overflow, visual word coordinates and native invocation cadence.

### DIALOGUE-062

The selected MAIN.RES families are mission event stems, inn NPC, mercenary
NPC, shop npc31m and training npc34m. The supplied extractor truncates at
the first NUL, selects each tag containing part= without an acceptance test,
starts after its following LF, and ends before the last CR preceding the next
tag or at payload end. Missing LF or a bounded CR raises an instrument error.
Conditional bodies are included as installed candidates. CString, array and
six-byte ASCII trimming services are supplied; the native parser/show/page
caller is not executed by this projection. Font-1 advances come from each original graphics
archive; first-record geometry is 16x15 and the advance table has 224 entries.
The font conversion selector is supplied as 0 for EN and 1 for RU, and the
font-height service is supplied as 15. Native configuration is not observed.

Original wrapping instructions with declared services yield EN 688 blocks,
2542 lines, maximum 7 and no multi-paragraph block. RU yields 732 blocks,
2650 lines, maximum 7 and 12 multi-paragraph blocks. The populations contain
1854 EN and 1906 RU justified lines, 11415 EN and 9550 RU word positions.

Binary rational nearest-even PC53/PC64 models with the original spill points
give zero coordinate differences from per-add double over those 20965
positions. A retained-PC64 accumulator differs at 634 EN and 332 RU
positions. These are arithmetic models over original-hooked wrapped lines,
not original justification execution on every corpus line. The 16 corpus
width vectors used in DIALOGUE-061 are additional finite controls.

**Confidence.** Medium for the corpus projection and model agreement. EN/RU
executable equality gives one code population. Agreement cannot select
native precision or prove the population's conditional reachability.

**Unknown.** Loose overrides, other installs, native reachability, native FPU
state and visible line/word positions. No universal maximum is claimed.

**Amended.** DIALOGUE-070 binds the candidate-body extraction to the original
accepted tail under explicit acceptance and string services. DIALOGUE-068
distinguishes the reached original caller from that projection. The former
"through the next tag" range shorthand is narrowed; all counts and arithmetic
model results stand.

### DIALOGUE-063

With the three existing 488x232 panel geometries supplied to coordinate
conversion, `R0760` requests 9 shadow sprites before 24 body sprites at
each screen size, 99 captured requests. Shadows use frames 3,5,6,7,8,
offset (+8,+8), level 6 and reverse argument 0. Body requests use level 0.
Parent coordinates, locks and clip services are explicit boundaries.

The original sprite vtable selects `R0794` for body and `R0795` for
shadow. The selected wrapper reads width/height and starts data at frame+12;
it passes the requested coordinates without adding a record origin.
Sixteen private forward-body controls execute `R0796`. RLE literals
select and write a supplied palette word, including palette value zero;
skip runs leave destination values unchanged. Source indices 0,1,127,255,
whole/clipped masks and padded buffers exclude destination blending and
palette-zero keying on this selected normal path.
The forward shadow selects `L03664` for reduced tables or `L03665` for
full tables. Stream literals mark covered pixels; their byte values are
discarded and existing destination values index the level-6 table.

Twenty-four private controls cover RGB565/RGB555, full/reduced selection,
literal bytes 1/127/255 and whole/clipped two-row masks. All pixels, 8-byte
row padding and trailing guards agree. The full-table law is
floor(channel*10/16). The reduced path substitutes the low blue-channel
representative (blue & ~7)+4, then applies that law. This excludes copying
the literal color and an unqualified full-channel law for reduced lookup.

**Confidence.** High for the selected request order, wrapper and forward
primitive under the stated formats and services. Fresh vtable/data, instruction
and transfer bindings connect caller, wrapper and decoder. Literal variation
and varied destination markers discriminate replacement from remapping.

**Unknown.** Native table/clip selection, reverse decoding, other layouts,
every shipped frame pixel, parent repaint and final native shadow pixels.

### DIALOGUE-064

The locally targeted suffix `L03666` runs with supplied actor, record,
canvas, background and destination state. Four record vectors choose
(36,140)-(108,232) when the second word is -1, otherwise
(x0,144-y1)-(x0+72,240-y1). Both opaque background and keyed canvas requests
place the selected window at (8,7). Varying the third/fourth words leaves
the selected requests unchanged. The background reverses before the opaque
copy and reverses back after the keyed canvas copy; border frame 0 is last.

Original bitmap reversal `L03667` reverses four three-pixel marker rows.
Sixteen copy controls run `R0797`/`R0798` over RGB565/RGB555, opaque/keyed,
whole/window/clipped/outside cases. Destination rows advance while physical
source rows descend from H-1-sourceTop. Opaque copy writes zero; keyed copy
skips 16-bit zero and copies every other supplied word unchanged. Pixels,
8-byte padding and guards agree. No blend occurs on this selected keyed path.

**Confidence.** High for the selected suffix, four metadata vectors and
finite primitive controls. Original vtable/data and instruction bindings
connect the requests to these primitives. Marker rows discriminate physical
address order; zeros discriminate opaque copy, keying and blending.

**Unknown.** Upstream actor/canvas production, final canvas orientation and
which rows are visibly exposed through the border. The suffix captures clear,
copy, border and cleanup services; it does not render their composition
end to end or prove native pixels.

### DIALOGUE-065

Original `R0768` paints the button with supplied screen rectangle
(276,296)-(356,322), cursor state and RGB565/RGB555 fields. Parent painting,
coordinate conversion, label drawing and disabled remapping are captured
services. Original line `R0769` and point `R0770` primitives write all
299 distinct bevel pixels exactly; padding and trailing guards remain intact.
The builder passes nonzero `L03668` through constructor `R0700` to +64,
selecting the dialogue ink arm. Constructor stores hover/press as zero and
sets flag word 3 through base `R0773`. These selected stores are read;
construction and parent attachment are not replayed end to end.

Six state vectors per format cover idle, hover, pressed-inside, pressed-outside,
pressed without hover and disabled-hover. Every label request uses (316,308)
and anchor 10. Hover +68 selects its ramp independently of press +6c. The
pressed flag plus supplied inside membership exchanges bevel colors and
changes shadow offset from 2 to 4; pressed outside retains offset 2.
Parent paint precedes label/bevel, and disabled level-3 remap follows painting.

The selected ramp-builder prefix `R0572` runs to its first allocation
boundary with supplied format fields. Level-15 text, idle and hover words
are RGB565 65535,48361,37568 and RGB555 32767,24169,18784. Flat shadow words
are 2113 and 1057. Light/dark bevel words are RGB565 10791/97 and RGB555
5383/33. This is conditional quantization, not native format detection.

**Confidence.** High for the selected painter, 12 state controls, bevel pixels
and supplied-format ramp outputs. The outside-cursor and no-hover controls
discriminate pressed-only and coupled-hover models. Whole-buffer comparisons
cover unchanged pixels as well as the bevel.

**Unknown.** Native format, hover/capture delivery, label glyph pixels, parent
frame repaint, disabled-remap pixels and cadence. DIALOGUE-055 separately
defines full/reduced level-3 arithmetic; this painter only captures its request.

## Resource input boundary and supplied projection

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DIALOGUE-068 | The selected accepted-part parser starts after header LF and ends at NUL or before the CR found backward from the next tag; the reached pager temporarily terminates that range and passes its bytes unchanged to the text and wrap entries. | High | ✔ promoted | [EXP-0421](../experiments/EXP-0421-dialogue-input/) |
| DIALOGUE-069 | Selected paragraph and wrap instructions trim boundary whitespace after resource delivery; original TrimLeft/TrimRight agree with the supplied six-byte ASCII services in 30 finite controls and preserve interior control bytes. | High | ✔ promoted | [EXP-0421](../experiments/EXP-0421-dialogue-input/) |
| DIALOGUE-070 | The explicitly supplied DIALOGUE-062 projection matches original accepted-tail ranges and original-trim wrapped lines on all 688 EN and 732 RU candidate bodies; native tag acceptance, locale and invocation remain unobserved. | Medium | ✔ promoted | [EXP-0421](../experiments/EXP-0421-dialogue-input/) |

### DIALOGUE-068

The selected parser is `R0712`, reached from pager `R0711` by show
`R0710`. Accepted header LF scanning is `L03669..L03670`.
The body start is the byte after that LF. The forward body scan stops at
`<` or NUL. At `<`, `L03671..L03672` scans backward until CR, without
a body-start test. The end pointer names that CR; at payload end it names
the terminating NUL. This is a pointer interval, not whitespace trimming.

The pager saves the end byte, writes NUL there, passes the start pointer
to text setter `R0725` at `L03673`, then restores the byte at
`L03674`. The original text setter passes that same argument to wrapper
entry `R0726` at `L03675`. The tag's lowercasing and CString services
do not copy or transform this body on the selected path.

Thirty-one finite controls enter full original show/parser/pager bodies with
declared string, registry and actor services. Twenty-nine reach both text
and wrap entries, preserve leading, trailing and interior supplied bytes,
and restore the payload. Missing header LF reaches no text boundary. With
no CR anywhere before the next tag, the instrument stops the backward scan
when it leaves the supplied payload. Two malformed controls find the header
CR before body start and pass data through the following tag; the bounded
Python extraction rejects them. These are private controls, not native hangs.

**Confidence.** High for these named pointer stores, transfers and finite
controls. Complete selected source spans are decoded from the lawful image;
internal branch targets bind to instruction starts. Header LF-only,
payload-end, adjacent-tag, rejected-tag and whitespace controls discriminate
delimiter, acceptance and preprocessing models. EN/RU images are identical
and count as one code population.

**Unknown.** Native acceptance and loose-resource loading, earlier caller
state, allocation failures, arbitrary malformed buffers, other callers and
visible pixels. A native memory/argument capture at these entries would
test the supplied-state premises without inferring them from this replay.

### DIALOGUE-069

Paragraph splitter `R0792` copies its input and splits CRLF pieces.
It calls TrimLeft on a CRLF piece and the remaining suffix at `L03676`
and `L03677`. Wrapper `R0727` copies each piece, then calls TrimLeft
at `L03678` and TrimRight at `L03679` before fitting. Later prefix and
remainder slices receive TrimLeft at `L03680` and `L03681`.
These operations occur after the raw text boundary of DIALOGUE-068.

Original TrimLeft `R0799` scans the leading classified characters and
moves the remaining bytes including NUL. TrimRight `L03682` scans forward,
remembers the current whitespace-run start, resets it after non-whitespace,
and terminates the final run. Both call classification `R0800` and
character stepping `L03683`; unique-buffer preparation is `L03684`.

The probe executes both trim bodies with unique private buffers, successful
memory movement, single-byte stepping and a supplied classifier for
SPACE/TAB/CR/LF/VT/FF. Thirty left, right and combined controls over bytes
1,9,10,11,12,13,28,32,127,160 agree byte-for-byte with the supplied services.
Their interior byte between A and B is retained. These services do not
establish the native classification or multibyte configuration.

**Confidence.** High for selected operation order and the finite comparison
under the named services. Separate leading, trailing and interior controls
exclude treating TrimLeft/TrimRight as whole-string normalization. Raw caller
delivery and later line/word construction are separate boundaries.

**Unknown.** Native locale tables, high-byte classification, multibyte stepping,
copy-on-write allocation, embedded-NUL CString behavior and arbitrary counts.
Native classification and character-step arguments/results on those cases
are the next discriminating observation.

### DIALOGUE-070

DIALOGUE-062 uses authored archive and body extraction, then original wrapping
and width instructions with supplied CString and array services. It does not
execute the resource parser, show or pager. Its body extractor selects every
part-bearing tag; substring and conditional acceptance are separate facts.
Its CString Left/Right counts are signed and clamped to the current length.
Its trimming removes the six bytes listed in DIALOGUE-069.

The complete independent MAIN.RES reader visits 494 EN and 504 RU nodes,
yielding 452 and 462 distinct file paths. The five selected families give
688 EN and 732 RU part-bearing candidate bodies. Each candidate enters
original `L03669` after a supplied accepted-header decision. All 1420
raw intervals equal the explicit Python projection. Replacing only the
supplied trim methods with original trim instructions under DIALOGUE-069's
services gives zero line-array differences: EN 2542 lines, maximum 7, no
multi-paragraph blocks; RU 2650 lines, maximum 7, 12 multi-paragraph blocks.
The projection reproduces 1854/1906 justified lines and 11415/9550 words.

Both runs use installed font-1 advances, spacing 2, width 300, height service
15 and conversion selector EN=0/RU=1. This comparison does not repeat or
extend DIALOGUE-062's arithmetic model results. Synthetic skipped and rejected
tags demonstrate that the first projected candidate need not be the first
body selected by a reached original call. Malformed backward-range controls
demonstrate another extraction difference outside the admitted stock bodies.

**Confidence.** Medium for the candidate projection and caller applicability.
Installed range and line equality is complete over this named population;
acceptance, font configuration and low-level string services remain supplied.
Equality cannot select native reachability, locale or every caller.

**Unknown.** Native accepted-tag population, overrides, other installed
families, native glyph conversion/height and arbitrary string edge cases.
Capture accepted header, body pointers and string-service configuration at a
reached native caller to discriminate these remaining premises.

## Mission 70 report identity

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DLG-REPORT-072 | The reported RU battle fragment has one match among 321 main.res text/battle/*.txt leaves: m70/event01.txt, Part1/NPC25; the preserved EN event has the same mission, event, part and NPC identity. | High | ● active | [EXP-0471](../experiments/EXP-0471-m70-start-report/) |

### DLG-REPORT-072

The RU locator projects the 321 named archive leaves as CP866 and searches
the owner's three fragment terms. One leaf matches. This declared display
projection is a locator, not a native encoding or parser claim.

The selected resource is `main.res::text/battle/m70/event01.txt`. RU has
110 bytes, SHA256
`4b0b7b602122f9ca3c59e857921e0b8d19331c98b18d4ac25088a0e8db303e4c`,
at `[107160,107270)`. EN has 120 bytes, SHA256
`1ad049e2c05105b2ec32ae63b87fa7ff73a696ad90ce1967214833be3ac91e54`,
at `[5871454,5871574)`. The sole part-bearing tag names Part1/NPC25.
The mission identity comes from directory 70 in both roots. The native titles
are `Monk Warriors` in EN and `Тайна монастыря` in RU.

Both `scenario.res::npc.reg` payloads have SHA256
`c91cff00c382617fb4867f734a176eb2b0358c6896e4e7b346384d3ceecfb223`.
Section npc25 has `Flags=Hero,Face,!Female,!Mage`, `Face=1` and
`DataBinID=42`. These authored keys do not identify a reached live actor.
The event's authored announcement is `MISSION-REPORT-070`.

**Confidence.** High for the named archive population, byte intervals and
cross-root resource identity. A complete main.res registry walk provides
the population; the selected resource bytes and tag spans are hashed by the
bounded probe. The search does not include patch overrides or other archives,
and does not prove which native tag is accepted.

**Unknown.** Native selected source, accepted tag, displayed actor/name,
first-seconds delivery and the reported save lineage. Capture the reached
resource and speaker at report-open to distinguish those premises.

## Voice binding and speaker identity in four missions

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| DLG-VOICE-074 | Mission 111 (The New Catapult) has no voice resource for any of its 14 event parts in EN or RU: `speech.res` holds no `battle\` path and no leaf `npc<NN>e<n>*` in any of seven archives; its inn text carries two voiced parts. | High | ● active | [EXP-0479](../experiments/EXP-0479-dialogue-art/) |
| DLG-SPEAKER-075 | Mission 110 event 02 has five parts: Part2, Part3 and Part5 (EN) speak as `npc67`, whose terms select the one placed unit 57 (Humans row `F_BrigandLeader3`); Parts 1 and 4 speak as the hero. RU Part5 is `npc25`. | High / Medium | ● active | [EXP-0479](../experiments/EXP-0479-dialogue-art/) |
| DLG-VOICE-076 | The portal-passage part of Skrakan's inn text (Part3, `npc25m150`) has 3 sentence marks in EN and 1 in RU, the RU part lacking the EN last sentence; whether the RU voice speaks more than its text is Unknown. | High / Unknown | ● active | [EXP-0479](../experiments/EXP-0479-dialogue-art/) |

### DLG-VOICE-074

- Population: `speech.res` of both roots (EN 355 nodes, RU 349), all in four directories `inn/mercenary`,
  `inn/npc`, `shop`, `training`. A walk of `sfx.res`, `main.res`, `graphics.res`, `world.res`,
  `scenario.res`, `movies.res` and `patch.res` finds no leaf `npc<digits>e<digit>*.wav` in either root.
- Mission 111 event texts `text/battle/m111/event01..08` hold 14 tagged parts in each root. Using the pager's
  event-leaf name `npc%02de%sp%d` (`DLG-SOUND-028`), the prefixes `npc21e`, `npc23e` and `npc25e` match
  0 nodes.
- Loss control: the inn family `text/inn/npc/<leaf>.txt` is bound to `inn/npc/<leaf>p<part>.wav`
  (or, when a part carries `sound=`, to `inn/npc/<value>.wav`; the value is the whole leaf). Of 114 EN
  and 152 RU tagged parts, 92 and 89 compose a name present in `speech.res`; 44 voice nodes in each
  root have no tagged part. So a missing voice is a missing node, not a failed composition.
- The inn text `npc02m111` (NPC2, 2 parts) is voiced in both roots (`npc02m111p1.wav`, `p2.wav`).
  Durations from the wave header: EN 10.32 s and 4.66 s, RU 12.10 s and 4.34 s.
- The 14 are tagged parts: event06 holds male and female variants of the same two parts.
- The binding is by name only: the part number and the leaf name, with `sound=` as an override.

**Confidence.** High for the searched population (seven archives per root, both roots). `Allods\*.RES`
(music and video archives) and loose files were not searched; the installs hold no loose `.wav`. Whether
the pager fails silently on an absent node was not run; `DLG-SOUND-028` names the composition.

### DLG-SPEAKER-075

- `text/battle/m110/event02.txt` has five parts. EN order of `<npc=N>`: 25, 67, 67, 25, 67. RU order:
  25, 67, 67, 25, 25.
- `npc.reg` sections (`DLG-NPCTAG-018`): npc25 `Flags=Hero,Face,!Female,!Mage`, `Face=1`; npc67
  `Flags=Human,!Mage,!Female,Face`, `Face=25`.
- `110.alm` has 39 type-6 placements and none uses the npc arm (flags bit 0). Applying the Flags terms
  to the Humans row reached by definition id (`MISSION-ARM-006`; Humans slots `0x10` typeID, `0x11`
  face, `0x12` gender) selects one unit of 39: unit 57, definition id to row `F_BrigandLeader3`,
  typeID 9, face 25, gender 0. No other placed unit has face 25.
- The announcement is one send-message action (id 15, trigger `Hi`) with condition
  `(26, 4, 28, 1, 0, 0)`. It names no unit.
- A part's speaker is therefore chosen by the part's own `npc` tag, resolved at show time
  (`DLG-SPEAKER-022`), not by the trigger.

**Confidence.** High for the tag order, section keys and the unique face match in the placed
population. Medium for the speaker being unit 57 at run time: the resolver's live-actor pick
among actors passing the terms was read from `DLG-SPEAKER-022`, not executed; death or
absence of unit 57 falls to the synthesised drawable.

**Unknown.** The EN/RU difference at Part5 (`npc67` against `npc25`), which one the authors
intended, and which actor name and portrait appear for RU Part5 at run time (not observed).

### DLG-VOICE-076

- File `text/inn/npc/npc25m150.txt`. EN has 12 parts, Part1 NPC25, Parts 2 to 12 NPC30 (Skrakan);
  RU has 10. The portal-passage line is Part3 in both roots.
- EN Part3: 199 chars, 3 sentence marks, voice `inn/npc/npc25m150p3.wav` 12.82 s (15.5 chars/s);
  silences of 0.2 s or more at 4.10 s, 6.24 s, 9.56 s.
- RU Part3: 127 chars, 1 sentence mark, voice `inn/npc/npc25m150p3.wav` 12.86 s (9.9 chars/s);
  silences at 3.00 s, 10.44 s, 12.50 s. RU Parts 3 to 9 run 9.9 to 12.1 chars/s against 12.8
  to 16.9 EN for Parts 2 to 9.
- Text comparison (counted): EN Part3 has 3 sentence marks and its last sentence says the Portal is a
  passage to another world; RU Part3 has 1 sentence mark and no clause corresponding to that last
  sentence. Both parts bind by `npc25m150p3` (`DLG-VOICE-074`); no tag differs.
- Rate context (RU Parts 4 to 9: 10.8 to 12.1 chars/s; all ten RU parts: 9.4 to 14.5): 127 characters at
  those rates take 8.8 to 13.5 s, a range that contains the 12.86 s voice. The 12.86 s is 1.1 to 2.4 s above
  the Parts 4 to 9 prediction. EN and RU Part3 durations are equal (12.82 s, 12.86 s), but RU voices are
  1.2 to 1.5 times as long as EN for Parts 4 to 9, so the equality is not evidence.

**Confidence.** High for the EN text and voice and for the RU text lacking the EN last sentence
(counted comparison of both `npc25m150` texts). Unknown whether the RU voice speaks anything beyond its
text: the excess over the RU rate is inside delivery variation and no transcript was made.

**Unknown.** The words of the RU voice tail; a listening check settles it.
