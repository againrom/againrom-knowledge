# TEXT — a string byte to a glyph on screen

Claims about the path a text byte travels: what the engine does to it between the shipped file
and the glyph blit, which atlas record it lands on, and what governs the whole path. Not a file
format. The pieces it sits on belong elsewhere and are not restated here:
[`claims/spr16a.md`](spr16a.md) owns the atlases as *files* (their pixel grammar, their record
count, the `.dat` sidecar, what the high half holds), while [`claims/res.md`](res.md) owns the
container the strings and atlases are read from. Spec:
[`formats/text/format.md`](../formats/text/format.md). Format of this file:
[registry.md](registry.md). IDs are permanent.

**The one thing to carry away.** The engine has a **language selector** and a **code-page
converter**, and both had been read as absent. `SPR16A-TXT-023` published "no code-page pass" and
therefore "how the RU game displays Russian text is open"; the pass exists, is called on every
byte of every string the engine draws, and is the reason the atlases' Cyrillic sits where it
does. The refuted clauses are in [`retracted.md`](retracted.md).

## Code page, selector and atlas index

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-CONV-001 | Every string byte the engine draws passes through a code-page converter first — `R0793` — and it is the missing half of the index rule. | High | ● active (partially retracted) | [EXP-0097](../experiments/EXP-0097-ru-text/), amended [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |
| TEXT-LANG-002 | The engine is language-conditional, the switch is one dword, and the game names its own language in its own data. | High | ● active | [EXP-0097](../experiments/EXP-0097-ru-text/) |
| TEXT-INDEX-003 | The atlas subscript is `record = (byte)(conv(b) − 0x20)`, and it is unbounded. | High / Medium | ● active | [EXP-0097](../experiments/EXP-0097-ru-text/) |
| TEXT-FIT-004 | The shipped RU text and the shipped atlases fit the converter exactly, in both directions, and neither fits without it. | High | ● active (superseded) | [EXP-0097](../experiments/EXP-0097-ru-text/), remeasured [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |
| TEXT-IN-005 | A second converter runs the other way, and it is the input path: `R1517` maps CP1251's Cyrillic onto the engine's own CP866. | High / Medium | ● active | [EXP-0097](../experiments/EXP-0097-ru-text/) |
| TEXT-LOWER-006 | The case fold is language-conditional too: `R1173` is a CP866-aware `tolower`. | High / Unknown | ● active | [EXP-0097](../experiments/EXP-0097-ru-text/) |
| TEXT-DOM-010 | The drawn-byte converter is NOT one-to-one, and on the Russian selector the failure is exactly 64 collision pairs. | High | ● active | [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |
| TEXT-ALIAS-011 | A colliding byte is not an error: it silently draws the other byte's glyph, in bounds, with no clamp and no substitute. | High / Medium | ● active | [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |
| TEXT-SEL0-012 | The non-injectivity belongs to selector 1 alone; on every other selector the pass is the identity on all 256 values. | High | ● active | [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |
| TEXT-FIT2-013 | The shipped Russian text lies entirely inside the source set, measured on a corpus 46× larger than the one previously reported — and the previous population was selected by a filter that leans toward the answer. | High / Medium | ● active | [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |
| TEXT-CAP-018 | The limit the text path implies, stated for the customisation seam: 160 distinguishable glyphs on the Russian selector, 224 on any other, and the ceiling is the byte, not the atlas. | High | ● active | [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |

### TEXT-CONV-001

The routine is nine tests and two adds: at `R0793` it loads the selector global `[L06203]`,
compares it with 1 and loads the argument byte from the stack — on any selector but 1 the argument is returned
untouched — then `L11045` tests the byte against the range `0x80..0xAF` and adds 0x30
(**`0x80..0xAF` → `0xB0..0xDF`**), and `L07934` tests it against `0xE0..0xEF` and adds 0x10
(**`0xE0..0xEF` → `0xF0..0xFF`**). Every other byte is returned unchanged, and
neither add can carry out of the low byte (`0xAF+0x30 = 0xDF`, `0xEF+0x10 = 0xFF`), so the returned byte is
always clean. `SPR16A-FONT-018` read the atlas subscript as the subtraction of 0x20 from the string byte; the
instruction two before it is the call to `R0793` at `L07922`, and the subtraction is applied to that call's
**result** (`L11046`…`L07943`: the call, the load of the frame argument, the copy of the result and the subtraction of 0x20).
~~The map is **injective** on `0x80..0xFF` — the two moved blocks land on `0xB0..0xDF` and
`0xF0..0xFF`, which no unmoved byte occupies (`evidence/converter-properties.txt`) — so no two
source bytes can collide on one glyph~~ — **REFUTED by `TEXT-DOM-010`.** `0xB0..0xDF` and
`0xF0..0xFF` *are* unmoved bytes and *do* occupy those cells: nothing removes them from the
fall-through path, so the map has 64 collision pairs and image 192 of 256. Two source bytes collide
on one glyph 64 times over, and what is drawn is `TEXT-ALIAS-011`. **Everything else in this row
stands** — the routine, the two adds, the block boundaries, the `CALL` two instructions before the
`SUB`, and the byte-cleanliness of `AL` were all re-read from both installs by EXP-0103 and are
unchanged

**Confidence.** **High** for the routine and the map, which EXP-0103 re-derived byte-for-byte / the
injectivity clause carried **High** and was wrong; see [`retracted.md`](retracted.md). The clause's
own stated reason named its counterexample, which is why it survived review
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Amended.** The injectivity clause struck through above is refuted by `TEXT-DOM-010` (EXP-0103;
[`retracted.md`](retracted.md), REFUTED). The routine, the two adds, the block boundaries, the
`CALL` before the `SUB` and the byte-cleanliness of `AL` stand.

### TEXT-LANG-002

`[L06203]` has exactly **one writer** image-wide and three readers (`EnumRefs refto:L06203` on
the repaired table: 4 hits, 4 owners, **0 in orphan or undisassembled code**; the companion
`imm:L06203` returns 0, so no site reaches it as an immediate). The writer is `R0616`, and it
computes the value from a **resource**: it opens `main\id` (the literal at `L11047`, `L11048`),
reads it into a stack buffer, takes that buffer's **last character** — `L11049` (a repeated byte scan for the terminator) /
the negate and decrement give the length, then `L11050` loads the signed byte at the buffer index `length − 1`,
which is `buf[strlen-1]` — and `L11051` subtracts 0x30 and `L11052` stores the result in `[L06203]`,
i.e. **the trailing ASCII digit as a number**. Measured, per root: EN `main\id` is 9 bytes
`656e676c6973682030` = `"english 0"` → selector **0**; RU is 9 bytes `7275737369616e2031` =
`"russian 1"` → selector **1**. `R0616` has one caller (`R0326` at `L11053`).
**`rom.exe` is byte-identical between the two roots** — sha256 `942e9b72…d367d03` on both — so the
entire language difference in the executable's behaviour is this one resource; a consumer that ships
one binary and two data sets is doing what the engine does

**Confidence.** **High.** The writer enumeration names its instrument and that instrument's blind
spot is not live here: `refto:` sees reads and writes through the reference manager, `imm:` covers
the address-as-immediate route and returns 0, and a whole-image `range:` census of the two text
modules shows **0 orphan bytes**. The two selector values are file measurements with the bytes
quoted, and the executable identity is our own hash of both roots

### TEXT-INDEX-003

The subtraction is byte-wide (the subtraction of 0x20 at `L07943`; the same subtraction at `L10047` and
`L11054`/`L11055`), and the result is widened with a mask of 0xff before it indexes (`L11056`,
`L11057`, `L11058`, `L11059`). The three accessors it reaches all subscript the frame-pointer
table with **no compare against the record count**: `R1250` (`vt+0x34`, the glyph blit) is
the load of the record table pointer at offset 0xc (`R1250`), the load of the glyph argument, and the indexed 4-byte-entry load at `L11060`, and
`R1787` (`vt+0x20`, width → frame `+0x0`) and `R0817` (`vt+0x24`, height → frame
`+0x4`) have the identical shape. So **a byte below `0x20` reads out of bounds**: `b − 0x20` wraps
to `0xE0..0xFF` = records **224..255**, past the last of the 224 (`SPR16A-FONT-015`), and the table
holds exactly `count` pointers (`SPR16A-RDR-017`) — the loaded pointer is whatever follows it, and
it is dereferenced. There is **no substitute glyph, no clamp and no skip**; the only byte the draw
loop special-cases is record **0** (space), which draws nothing and adds `GetHeight(0) >> 1`
(`L11061` tests the glyph for zero and branches to `L11062` on zero; the virtual call through slot offset 0x24 and the arithmetic shift right by 1 follow) before falling into the
shared advance. Since the loop is `strlen`-bounded a NUL never arrives, but **a newline or tab does
if a caller passes one** — so line splitting is the caller's obligation, not the text routine's

**Confidence.** **High** for the index rule and for the absence of a bound (six quoted instructions
across three accessors, all three read whole and each 13 bytes long, plus the two `SUB`/`AND` pairs)
/ **Medium** for the *consequence* of an out-of-range record: the load and dereference are read, but
no shipped string was found that reaches the draw routines carrying a control byte, so the fault is
predicted from the instructions rather than witnessed. Whether any caller splits lines before
drawing is EXP-0098's area and is not answered here

### TEXT-FIT-004

Domain side: over the RU root's text surfaces — 30 nodes, **1 952 high bytes** — every high byte
lies in `0x80..0xAF` or `0xE0..0xEF`, the two blocks `TEXT-CONV-001` moves, and **0 lie anywhere
else** (`0xB0..0xDF`: 0; `0xF0..0xFF`: 0 on `MAIN.RES` and `patch.res`, whose 1 895 bytes are the
localised text). Image side: the converter's image is the 64 cells `0xB0..0xDF ∪ 0xF0..0xFF`, and
those cells carry ink on **64/64** in `font1`, `font2`, `font4` and `font5`; in the RU root's
`font2` — the one node the RU release replaced (`SPR16A-FONT-022`) — the high-half ink is
**exactly** that image, `0` cells of ink outside it, while the EN `font2` has **35** (the CP437
accent records the RU build blanked). Closure: under the converter **1 952 / 1 952** RU high bytes
land on a record that has ink, on all four 224-record atlases — **100.00 %**. Under the identity
rule the same bytes score **41.39 %** on font1/font4/font5 and **2.77 %** on font2, whose 54 on-ink
hits are all residue from `world.res`, so the localised text alone scores **0 of 1 895** there: the
RU release's own smallest font would draw **nothing at all** for every Russian byte. Population and
its blind spot are stated in the probe: text nodes are `.txt`/`.ini`/`.lst`/extensionless payloads
plus the embedded string tables of `.reg` and `data.bin`; **a display string embedded in any other
node type is not counted**. ~~30 nodes, 1 952 high bytes~~ — **the figure and the population are
SUPERSEDED by `TEXT-FIT2-013`**, which measures 419 nodes and 87 293 high bytes on the same root and
reproduces the 100 % fit exactly. Both of this row's filters lean toward the answer: the node gate
requires ASCII majority, which a majority-Cyrillic file fails by construction and which admitted 39
of 419 RU text nodes, and the run filter counts a byte as a letter only when it already lies in the
two blocks. **The conclusion is unchanged and now rests on 46× the bytes**

**Confidence.** **High.** This is corpus evidence, which caps at Medium on its own — what lifts it
is that the two sides were measured independently and the fit is exact and two-sided: the code's
*image* (read from instructions, before the census ran) equals the atlas's *ink set*, and the
corpus's *alphabet* equals the code's *domain*, with zero exceptions on either. The rival — that
strings and atlases agreed all along and no converter is needed — is refuted by the same numbers at
0/1 895 on font2. The figures carried **High** when believed and were measured through a biased
population; see [`retracted.md`](retracted.md)

**Amended.** The figure and the population struck through above are superseded by `TEXT-FIT2-013`
(EXP-0103; [`retracted.md`](retracted.md), SUPERSEDED). The conclusion stands.

### TEXT-IN-005

the test at `L11063` (a byte below `0x80`) returns any ASCII byte **untouched** — this arm calls no CRT routine
at all — then `L11064` loads the selector global `[L06203]` and decrements it, and a non-zero result returns unchanged on any other
selector. On selector 1: `L11065` tests the byte against `0xC0..0xEF` and `L11066` adds 0xC0, and `L11067` tests it against `0xF0..0xFF` and
`L11068` adds 0xF0. Adding 0xC0 is
`− 0x40` and adding 0xF0 is `− 0x10` in byte arithmetic, so the map is `0xC0..0xEF → 0x80..0xAF`
and `0xF0..0xFF → 0xE0..0xEF` — exactly the inverse of the block layout the display side assumes,
from the encoding a Windows edit control hands over. **6 call sites, 6 distinct owners, 0 orphan**
(`EnumRefs callto:R1517`). This is the reason the RU root's `README.TXT` reads as CP1251 while
its data strings read as CP866 (`SPR16A-TXT-023` measured both and could not join them): they are
two different sides of the same boundary

**Confidence.** **High** for the map (every branch and both adds are quoted from an 18-instruction
routine read whole) and for the call-site count (instrument named, 0 orphan) / **Medium** for
calling it *the input path*: the six owners were enumerated but not read, so which surface each
serves — typed hero name, chat, a file read — is not established here

### TEXT-LOWER-006

`R1173` loads the selector global `[L06203]`, decrements it and jumps to `L11069` on zero, selecting the arm; **every other selector
falls through to the CRT `tolower`** at `R1174` (`L11070`), so the English build's folding is
the C library's and nothing else. On selector 1: `L11071` tests the byte against `0x80..0x8F` and `L11072` adds 0x20, which
folds `0x80..0x8F → 0xA0..0xAF`, and `L11073` tests it against `0x90..0x9F` and
`L11074` adds 0x50, which folds `0x90..0x9F → 0xE0..0xEF`; anything else
reaches the CRT call at `L11075`. Those are exactly CP866's two uppercase blocks onto its two
lowercase blocks — the same block geometry `TEXT-CONV-001` moves, which is what makes the three
routines one system rather than three coincidences. **3 call sites, 3 distinct owners, 0 orphan**
(`EnumRefs callto:R1173`); one owner, `R1852`, calls **both** this and `TEXT-IN-005`'s
routine

**Confidence.** **High** for the fold and the CRT fallthrough (both arms quoted; the fallthrough is
a `CALL` to a named address on two of the three exits) and for the call-site count (instrument
named, 0 orphan) / **Unknown** what the three callers do with it — no case-insensitive comparison
was traced to a surface. Note this fold is **not** the archive path fold: `RES-CASE-036`'s
`R1853` is a different routine reaching the CRT `tolower` directly, and it is not
selector-gated

### TEXT-DOM-010

`R0793` is 42 bytes read whole out of both roots' `rom.exe` by virtual address
(`evidence/converter-domain.txt`; the two dumps are identical and the probe checks its own
transcription against them). The first arm's image is `0xB0..0xDF`, and **nothing removes
`0xB0..0xDF` from the fall-through path**: the above-branch at `L07933` sends every byte above `0xAF` to
`L07934`, and the carry branch at `L07935` sends everything below `0xE0` straight to the plain return (popping 4 argument bytes) at
`L07936`. The same shape returns `0xF0..0xFF` unchanged via the above-branch at `L07937`, onto the second arm's
image. So under selector 1 the map has **image 192 of 256**, **64 output values with two preimages
each** — `0xB0` from `0x80` and `0xB0` … `0xDF` from `0xAF` and `0xDF` (48 pairs), `0xF0` from
`0xE0` and `0xF0` … `0xFF` from `0xEF` and `0xFF` (16 pairs) — and it is **not injective on
`0x80..0xFF`**. The routine has no third test and no default arm; the collision is structural, not a
data accident. There are 2^64 maximal one-to-one subsets of size 192, so **the code alone does not
name a domain**: two are natural, the source set `0x00..0xAF` u `0xE0..0xEF` and the map's own image
`0x00..0x7F` u `0xB0..0xDF` u `0xF0..0xFF`, and only the shipped text picks between them
(`TEXT-FIT2-013`)

**Confidence.** **High.** Every branch is a quoted instruction from one 42-byte routine read end to
end and verified byte-for-byte against both installs, and the map was then enumerated over the
**whole** 0..255 range rather than sampled, for both selector values. This **refutes**
`TEXT-CONV-001`'s clause "the map is injective on `0x80..0xFF`", which carried High; that clause's
own stated reason — "the two moved blocks land on `0xB0..0xDF` and `0xF0..0xFF`, which no unmoved
byte occupies" — names its counterexample, since those two ranges *are* the unmoved bytes occupying
them. See [`retracted.md`](retracted.md)

### TEXT-ALIAS-011

Under selector 1, byte `b` in `0xB0..0xDF` selects record `b − 0x20` = 144..191 — **the same
record** as byte `b − 0x30` — and a byte `b` in `0xF0..0xFF` selects record 208..223, the same as
`b − 0x10` (`evidence/conv-table.csv` gives the aliasing partner for all 256 values). Every one of
those records is **inside** a 224-record atlas, so unlike `TEXT-INDEX-003`'s sub-`0x20` case the
subscript does not even leave the frame table: the blit is well-formed and a real glyph appears.
`R0793` has two `RET`s and no default branch, and the three accessors subscript with no
compare — `L11076`, `L11077` and `L11060` are all loads from a base plus index times 4, each routine read
whole. Per `SPR16A-FONT-020` the records concerned hold Cyrillic, so the visible consequence is that
the RU build cannot show the 64 characters those source bytes name in their own page: each draws a
Russian letter instead

**Confidence.** **High** for the aliasing and for the absence of any error arm — the record
arithmetic is the byte-wide `SUB`/`AND` pair already published in `TEXT-INDEX-003`, re-read here,
and all three accessors were read end to end rather than decompiled / **Medium** for *which letter*
appears, which rests on `SPR16A-FONT-020`'s reading of the atlas rather than on anything measured
this round

### TEXT-SEL0-012

the selector compare with 1 at `L11078` and the unequal branch to `L07936` at `L11079` reach a bare return (popping 4 argument bytes) with the argument byte untouched, so
the returned byte is the argument. Enumerated over the whole range at selector 0: **0 bytes moved,
image 256, 0 collisions** (`evidence/converter-domain.txt`). Two consequences a consumer must carry.
First, the two selectors do not merely differ in *which* glyph a high byte draws — they differ in
whether the mapping is **information-preserving at all**, so a reimplementation cannot model the
pass as one table with a language parameter unless that table is allowed to be non-injective in one
column. Second, the reachable record set differs: at selector 0 all 224 records of a 224-record
atlas are selectable, at selector 1 only 160 (`TEXT-CAP-018`). The upper 24 bits of the returned
`EAX` are the selector's on this arm, and are clean only because the shipped selectors are 0 and 1;
every caller reads `AL`, so it does not matter, but a consumer returning a full word would inherit a
bug the original does not have

**Confidence.** **High.** Two quoted instructions and a bare `RET`, plus an exhaustive 256-value
enumeration at that selector. The alternative this rules out is the one a reader of `TEXT-CONV-001`
would form — that the converter is a code-page table applied always — and it is ruled out by the
`JNZ` target being the function's own exit

### TEXT-FIT2-013

Population, stated so it can be attacked: every node of every archive on the root whose
**extension** is `txt`/`ini`/`lst`/none and which passes a **code-page-neutral** gate (under 2 %
control bytes; it counts control bytes, never letters). RU: **419 nodes, 146 709 bytes, 87 293 bytes
at or above `0x80`, of which 0 lie in `0xB0..0xDF` and 0 in `0xF0..0xFF`** — 100.0000 % inside
`0x80..0xAF` u `0xE0..0xEF`. EN control, same probe: 413 nodes, 132 727 bytes, **one** high byte in
the entire corpus (`0xE5`). Widening to every NUL-terminated run in every non-sampled node adds only
`.alm` and `.bin` residue, and that residue is **root-invariant** — 1 002 bytes outside the set on
RU against 1 006 on EN, 111 against 111 — so it is binary the run heuristic mistakes for text, not
localised text. **Why `TEXT-FIT-004`'s 1 952 over 30 nodes does not reproduce:** its node gate
requires ASCII majority, which a single-byte Cyrillic file fails by construction, admitting **39 of
419** RU text nodes; and its run filter credits a byte as a letter only when it already lies in the
two blocks. Run over the same nodes in one pass, that filter reports 99.78 % fit where the neutral
one reports 98.75 %, and on EN 52.49 % against 26.99 % — it roughly doubles the apparent fit by
declining to look outside the blocks (`evidence/corpus-en.txt`, `evidence/corpus-ru.txt`)

**Confidence.** **High** for the count itself — it is exhaustive over a mechanically defined
population, the search for a counterexample was neutral over the whole high range by construction,
and zero were found / **Medium for the consequence** that the RU release never relies on a colliding
byte, and the cap is not negotiable: this is corpus evidence, and the *population* is a filter
choice that no reading of the engine forces. What would lift it is a decode of the nodes that carry
display strings, which this round did not do

### TEXT-CAP-018

The subscript is byte-wide (the subtraction of 0x20, then a mask of 0xff), so a font can address at most **256**
records; the shipped atlases provide **224**, and `font3` 64. Enumerated over all 256 inputs: at
selector 0 every one of the 224 records is selectable and 32 bytes (`0x00..0x1F`) read past the end;
at selector 1 only **160** are selectable — 96 from `0x20..0x7F`, 48 from the first arm's image
`0xB0..0xDF`, 16 from the second's `0xF0..0xFF` — and **64 records are unreachable dead weight**,
exactly the character codes `0x80..0xAF` and `0xE0..0xEF`. So a Russian string can name 160 distinct
glyphs, not 224, and the missing 64 are not missing art but unaddressable slots. Three separate
things would each have to change to lift it, and a consumer should know which: adding records past
224 needs only the atlas and its `.dat`, because the readers validate no frame header at all
(`SPR16A-RDR-017`); reaching records 224..255 needs the byte-wide subscript widened, which changes
code and not files; reaching the 64 dead records at selector 1 needs a third test inside the
converter, which changes the meaning of every shipped Russian byte and so **cannot** be done without
rewriting the text. That last one is the seam: the code-page pass is where language support is
bolted on, and it is not extensible in place

**Confidence.** **High.** The arithmetic is an exhaustive enumeration over all 256 values against
record counts taken from each shipped file's own trailer, with no free parameter, and the three
"what would have to change" clauses each name the instruction or the file that carries the limit.
What this does **not** establish is whether anything downstream — a save file, a network path, a
`.dat` consumer — assumes 224 rather than 160; no such sweep was run

## Text API, atlases and tilde markup

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-API-007 | The text API is four routines on one class, and all three that touch a string byte convert. | High / Medium | ● active | [EXP-0097](../experiments/EXP-0097-ru-text/) |
| TEXT-FONT3-008 | One shipped atlas cannot render Russian at all, and it is a structural fact rather than a gap in its art. | High / Medium | ● active | [EXP-0097](../experiments/EXP-0097-ru-text/) |
| TEXT-TILDE-009 | `~` is markup, not a character, and a consumer that measures text must know it. | High / Unknown | ● active (amended, partially retracted, superseded) | [EXP-0097](../experiments/EXP-0097-ru-text/), amended [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |
| TEXT-FONT-015 | Every value the converter can produce reaches a record that has ink, on every 224-record atlas, on both roots — so "converts correctly, then draws nothing" does not happen. | High / Medium | ● active | [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |
| TEXT-FONT2-016 | `font2` is the only atlas the roots do not share, and the RU build blanked its entire Latin high half — which is free under selector 1 and destroys text under selector 0. | High / Medium | ● active | [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |
| TEXT-TILDE2-017 | The two byte loops do not test `~` at the same point: the measurer tests the raw byte, the draw tests the converted one. The conclusion survives; the published reason does not. | High | ● active | [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |
| TEXT-079 | A lone `~` in the font1-3 draw underlines the next glyph and takes no pen advance: a line from x, that glyph's advance-table width long, one row below frame 0, in ramp entry 15. `~~` is one literal glyph; the font4 draw has no tilde arm. | High / Unknown | ● active | [EXP-0410](../experiments/EXP-0410-dialogue-layout/) |

### TEXT-API-007

The font object is 0x10 bytes — `+0x0` vtable `L11080`, `+0x4` the glyph sprite, `+0x8` the `.dat`
advance table, `+0xc` the letter spacing (`SPR16A-FONT-018`) — and its method table has 12 slots of
which `+0x14` is `R0767` (draw) and `+0x18` is a two-instruction getter. Beside the vtable
sit `R1854` (a second draw, taking a length rather than a flag word), `R0766` (measure
a string's width) and `R0727` (a wrap/layout pass that owns no byte loop of its own — it
calls the measurer five times and never indexes an atlas). The converter's callers are exactly the
first three: **5 call sites, 3 distinct owners, 0 orphan** (`EnumRefs callto:R0793` — `L07941`,
`L11081` in the measurer; `L07922`, `L10051` in the draw; `L10052` in the second draw). Four
font objects are constructed, all in `R1788`, from whole string literals
`graphics\font1\font1` … `graphics\font4\font4` (`L11082`, `L11083`, `L11084`, `L11085`)
with spacing `2`, into four globals: **font1 `[L03615]` 179 refs / 48 owners, font2
`[L02677]` 135 / 22, font3 `[L10053]` 4 / 4, font4 `[L06186]` 89 / 20**, all with 0
orphan hits

**Confidence.** **High** for the converter's call-site set *within the text module*: a `range:`
census of `L11086..L11087` and `L11088..L11089` reports **0 orphan bytes and 0
undisassembled bytes that are code** (17 and 68 bytes of padding), so no hidden routine sits between
these and the enumeration cannot have dropped one there / **Medium** for "every text path in the
image converts", and the blind spot is named and is real: the glyph blit is reached as `vt+0x34`, so
a caller that computed a subscript itself and dispatched virtually would appear in **no** `callto:`
sweep. None was found; none was ruled out / **Medium** for the per-font reference counts as a
measure of *surfaces* — they count references, and the owners were not read

### TEXT-FONT3-008

`font3.16` has **64** records, not 224 (`SPR16A-FONT-015`), so under `record = char − 0x20` it
covers characters `0x20..0x5F` only — no lowercase ASCII and **none** of the converter's image:
`0 of 64` of the cells `0xB0..0xDF ∪ 0xF0..0xFF` exist in it, on both roots. Its ink is narrower
still: **11 records** carry any, at characters `0x30..0x39` and `0x46` — the ten digits and one
letter — with `53 of 64` advances zero. It is constructed like the others and read from three sites
(`R1855`, `R1856`, `R1059`). So the area's question has no single answer: any
surface drawn with `font3` is numeric whatever the selector says, and a consumer must not treat the
four atlases as interchangeable

**Confidence.** **High** for the record count, the covered character range and the ink census (the
count is the file's own trailer under an already-closed walk, and the range follows from the index
rule with no free parameter; the ink figures are a decode of every record on both roots) /
**Medium** for which *surface* `font3` serves — its three readers were enumerated, not read

### TEXT-TILDE-009

~~Both byte loops test it, at the same point and in the same way, on the value **after**
conversion~~ — **SUPERSEDED by `TEXT-TILDE2-017`**: the draw tests the **converted** byte
(`L07938`, a compare of the converted byte with 0x5e, after the call at `L07922`), the measurer tests the **raw** one
(`L07939`, a compare of the byte with 0x7e, nine bytes *before* its call at `L07941`). The two agree on all 256
inputs anyway — nothing but `0x7e` converts to `0x7e` — so the rule below is unchanged, but a
consumer must not convert before testing in the measurer. The two addresses this row already gave
are correct; the sentence joining them was not: the compare at `L07938` (with 0x5e) in the draw and
the compare at `L07939` (with 0x7e) in the measurer (`0x5e` is `0x7e − 0x20`, the same char). **Doubled**, it is
the literal glyph: the draw takes the equality compare of the two bytes at `L11090` (equal jumps to `L11091`) into the ordinary path and
then skips its partner (`L11092` compares the next byte with the same value and, when equal, `L11093` increments the stack index at frame offset `+0x2c`), and the measurer does the same at `L11094`. **Single**, it
draws a **rule** and occupies no width: the draw calls `R0769` — a Bresenham line routine — with
the colour taken from `word ptr [arg+0x1e]` (`L11095`), from `(x, y + GetHeight(0))` to
`(x + dat[j], y + GetHeight(0))` (`L07707`…`L07708`), `j` the glyph index of the byte after
the `~` (`TEXT-079`), and then **jumps past the advance
block** (`L11096 JMP L11097`), so `x` does not move; the measurer's single-`~` arm likewise
reaches `L11098` without adding anything to its running total. The two routines therefore agree,
which is what makes this a rule rather than a quirk of one of them

**Confidence.** **High** (both arms of both routines are quoted, and the two agree — a consumer's
width calculation and its draw would desynchronise if either were read wrong) / **Unknown** what the
rule is *for*, and which shipped strings use it: the five dialogue families and string 77 hold no
`~` (`TEXT-079`), and no other text was searched. The "after conversion"
clause carried **High** and was wrong for the measurer; see [`retracted.md`](retracted.md)

**Amended.** The clause that both byte loops test `~` at the same point and after conversion is
superseded by `TEXT-TILDE2-017` (EXP-0103; [`retracted.md`](retracted.md), SUPERSEDED): the measurer
tests the raw byte. The markup rule stands.

A second correction (EXP-0410): the single-`~` sentence read `to (x + dat[0x5e], y + GetHeight(0))`,
and the Unknown read "the argument struct whose `+0x1e` supplies the colour was not identified" and
"no census of `~` in the shipped text was run". The line ends at `x + dat[j]`, `j` the recoded byte
after the `~`, and the colour word is entry 15 of the fifth argument, the ink ramp (`TEXT-079`). The
dialogue families and string 77 were counted: no `~`. The rule, the doubled case, the no-advance
jump and the measurer's arm stand.

### TEXT-FONT-015

Record counts taken from each file's own trailer under a walk that is exact: `font1` 224, `font2`
224, `font3` 64, `font4` 224, `font5` 224, identically on both roots, with 0 bytes of slack between
the last record and the trailer on `font3`, `font4` and `font5`. Decoding **every** record of every
sheet on both roots (`evidence/font-records-en.csv`, `evidence/font-records-ru.csv`, 960 rows per
root): of the 64 character codes `0xB0..0xDF` u `0xF0..0xFF` that are the converter's image, **64 of
64 carry ink** in `font1`, `font2`, `font4` and `font5` on **both** installs. Consequently, under
selector 1 the only byte in `0x20..0xFF` that selects an inkless record on those four atlases is
`0x20` itself — which the draw loop special-cases before it ever blits (`L11061`, a zero test with a branch to `L11062`). The 32 bytes `0x00..0x1F` still read past the last record, which is `TEXT-INDEX-003`
measured per font rather than predicted: records 224..255 against a count of 224. `font3` is the
exception and is already published (`TEXT-FONT3-008`): 64 records, 11 with ink, and 0 of the image's
64 cells exist in it at all

**Confidence.** **High** for the record counts and the ink census — the counts are each file's own
trailer under the reader's own predicate (`SPR16A-RDR-017`), the walk is exact, and the census is a
decode of every record on both roots rather than a sample / **Medium** for the decoder: the `.16`
and `.16a` grammars are this repo's published claims (`SPR16A-FONT-013`, `SPR16A-RLE-002`) re-run
here, not re-derived, so an error in them would move these figures. "Ink" is counted as any pixel
the decoder set, which is a decode property, not a rendering

### TEXT-FONT2-016

`font2.16` is 13 892 bytes on EN and 9 436 on RU (sha256 `c262119b…cb0c071b` against
`c5108777…ed66cdef`); `font1`, `font3`, `font4` and `font5` are byte-identical across the installs.
Both `font2` files declare 224 records and walk exactly. EN carries ink on **194** records, RU on
**159**; the 35-record difference lies **entirely** in character codes `0x80..0xAF`, and those are
precisely the records selector 1 can never select, because no byte converts into `0x80..0xAF`
(`TEXT-DOM-010`). So the RU release deleted exactly the art its own selector cannot reach, and lost
nothing. The consequence runs the other way and is a real trap for a consumer: under **selector 0**
the RU root's `font2` selects an inkless record for **65** byte values — `0x20` plus the whole of
`0x80..0xAF` and `0xE0..0xEF` — against 30 for the EN file. A build that ships RU data with the EN
selector, or that omits the selector entirely, loses text on that font and only on that font,
silently and with no fault

**Confidence.** **High** for the counts and for the localisation of the difference: two files, both
walked exactly, every record decoded on both roots, and the blanked set intersected against the
converter's reachable set by enumeration rather than by inspection / **Medium** for the selector-0
consequence, which is predicted from the index rule and the ink census and was not witnessed
running. This overturns nothing in `SPR16A-FONT-022`; it adds *which* records and *why* it cost the
RU build nothing

### TEXT-TILDE2-017

In the measurer `R0766` the test is the compare at `L07939` with 0x7e, where the tested byte was loaded three bytes
earlier by `L07940` (a load from the string at the index) — the string byte — and the call to `R0793` does not run
until `L07941`, **nine bytes later**, after the branch has already been taken. In the draw
`R0767` the test is the compare at `L07938` with 0x5e on the converted byte, which is `L07942` (a copy of the result) and
`L07943` (the subtraction of 0x20) applied to the result of the call to `R0793` at `L07922`. `TEXT-TILDE-009` says
both test the converted value "at the same point and in the same way"; they do not. The two
nonetheless agree on all 256 inputs, and the reason is `TEXT-DOM-010`: `conv` is the identity below
`0x80` and sends everything at or above it to `0xB0` or higher, so `conv(b)` equals `0x7e` if and
only if `b` equals `0x7e`. The rule `TEXT-TILDE-009` published is therefore correct and stays; what
changes is where a consumer must put the test. A reimplementation following that row's description
literally will convert before testing in the measurer — harmless today, and it stops being harmless
the moment the converter is touched, which is what `TEXT-CAP-018` is about

**Confidence.** **High.** Four instruction addresses, all inside routines this experiment dumped and
decoded from both installs, and the agreement is proved by the enumeration rather than assumed. The
rival — that the addresses are one routine read twice, or that a second `CMP` exists after the
`CALL` — is killed by the measurer's own listing, in which the only `CMP` against `0x7e` before
`L11094` is the one at `L07939`
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### TEXT-079

- Each loop step of `R0767` recodes the byte at `i` and the byte at `i + 1`
  (`R0793`) and subtracts 0x20 from both, so `~` becomes 0x5e. The tilde arm
  (`L07938`..`L11096`) runs when byte `i` is 0x5e while byte `i + 1` is not (the equality compare of the two bytes at `L11090`). When both are 0x5e the ordinary glyph path draws one `~` glyph and the loop
  skips the partner (`L11092`..`L11093`): the doubled case of `TEXT-TILDE-009`.
- The arm calls the line routine `R0769(x1, y1, x2, y2, colour)`, which plots both
  endpoints, with
  - x1 = x, the pen position;
  - x2 = x + `dat[j]`, where `j` is the recoded index of byte `i + 1` (the stack slot written
    at `L07709` and read at `L07710`) and `dat` is the font's advance table, without the
    letter spacing;
  - y1 = y2 = y + `vt+0x24(0)`, the row below frame 0's height, which is 15 in font1
    (`TEXT-078`);
  - colour = the 16-bit word at `ramp + 0x1e` (`L11095`): entry 15 of the fifth argument,
    the top level of the ramp that inks the text.
- The arm then jumps past the advance block (`L11096 JMP L11097`), so the pen stays at x
  and the next step draws the next glyph there. The line is `dat[j] + 1` pixels long and lies
  under that glyph.
- Through the wrapper `R0571` both calls draw the line, the first `d` pixels right and
  down in the shadow ramp's entry 15 (`TEXT-078`).
- A tilde in the last byte of a string reads its partner from the terminating NUL, index 0xe0
  after the recode, which is one dword past the end of font1's 224-entry table
  (`evidence/metrics.tsv`). The line length then depends on what follows the table in memory.
- The font4 draw `R1854` (`TEXT-067`) has no tilde arm. Its loop recodes each byte,
  subtracts 0x20 and either draws the glyph or advances for a space, so it draws `~` as glyph
  0x5e and advances by that entry, while the measurer skips a lone `~` (`TEXT-TILDE2-017`). A
  font4 string that holds a tilde is measured narrower than it is drawn. No font4 string was
  searched for one.
- Census: the five dialogue families of both roots contain no `~` byte (`DLG-LINE-038`), and
  the button label, string 77, contains none (`DLG-BUTTON-039`). No other text was searched.

**Confidence.** High for the arm's operands and the doubled case: `R0767` and
`R0769` are read whole and each operand is traced to its stack slot, which settles the
extent as the next glyph's advance and not the tilde glyph's own entry, 0x5e. High for the
missing arm in `R1854`, read whole, which holds no compare against 0x5e.

**Unknown.** The read past the end of the advance table for a final tilde, what lies after the
table, and whether any shipped string outside the searched population reaches the arm.

## Draw routines

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-065 | `DrawText`'s font4 (`.16a`-class) callee unconditionally null-dereferences the argument `DrawText` passes as a literal zero, so `DrawText` cannot draw font4 text; only font1-3 (`.16`-class) pass through it. | High | ● active | [EXP-0399](../experiments/EXP-0399-text-presentation/EXP-0399.md) |
| TEXT-066 | Font4's constructor builds its `.16a` sprite and its own shading table once, at font-load time — not per `DrawText` call. | High | ● active | [EXP-0399](../experiments/EXP-0399-text-presentation/EXP-0399.md) |
| TEXT-067 | The second draw routine's vtable slot is a bare no-op for font1-3 (`.16`-class) and the published alpha compositor for font4 (`.16a`-class), which makes it the confirmed font4 draw path. | High | ● active | [EXP-0399](../experiments/EXP-0399-text-presentation/EXP-0399.md) |
| TEXT-068 | Neither text draw routine issues an outline/shadow/second-offset pass around a glyph; the tilde markup case is a substitute pixel-producing call, not an additional one. | High | ● active | [EXP-0399](../experiments/EXP-0399-text-presentation/EXP-0399.md) |
| TEXT-069 | No surviving claim in the corpus asserts anti-aliasing, supersampling or sub-pixel filtering anywhere in the sprite/text pixel path, and neither text draw routine's own body contains filtering logic. | Medium | ● active | [EXP-0399](../experiments/EXP-0399-text-presentation/EXP-0399.md) |
| TEXT-070 | Colour reaches a font4 pixel through a table-pointer selector, not a level shift: the shift-left-by-9 level addressing belongs to `R1547`, not to font4's real receiver `R1785`. | High / Medium / Unknown | ● active | [EXP-0399](../experiments/EXP-0399-text-presentation/EXP-0399.md) |
| TEXT-071 | `R1547` unconditionally dereferences the argument `DrawText` supplies as a literal zero, on both branches, before any pixel work: a null-pointer-plus-8 fault, not a benign argument miscount. | High / Unknown | ● active | [EXP-0399](../experiments/EXP-0399-text-presentation/EXP-0399.md) |
| TEXT-078 | The font's two draw routines read flag values 1, 2, 4 and 8 as anchors: 1 and 2 move x left by the string's width and by half of it, 4 and 8 move y up by frame 0's width, not its height, and by half of it. | High / Unknown | ● active | [EXP-0410](../experiments/EXP-0410-dialogue-layout/) |

### TEXT-065

`DrawText`'s glyph-draw call is a polymorphic vtable call, but its font4 (`.16a`-class) callee
unconditionally null-dereferences the exact argument `DrawText` passes as a literal zero, before any
pixel work — `DrawText` cannot be the routine that draws font4 text; only font1-3 (`.16`-class) can
pass through it without crashing.

`R0767` at `L11099` issues the virtual call through slot offset 0x34 after pushing exactly five stack
dwords, the first one pushed (which becomes the callee's own arg5) a literal 0 pushed at
`L11100`. For font1-3 (`.16`-class, vtable `L10048`) this resolves to
`R1250`→`R1251`, the already-published opaque byte blitter (`SPR16A-FONT-013`,
`SPR16A-072`); `R1250`'s full body never reads its own arg5 at all (only arg1/this, arg3 and
arg4 feed the six values it forwards to `R1251`) — the constant `DrawText` pushes there is
simply inert for this class. Font4 alone is a `.16a`-class object (vtable `L07521`, confirmed
directly: its constructor `R1790` calls `R1771` at `L11101`, which sets the
sprite's own vptr to `L07521`), so the same instruction would instead resolve to `R1547`
— but `R1547` dereferences arg5 unconditionally on both of its branches (the load of the argument at frame offset 0x18
then the load of its field at offset 0x8 at `L11102`/`L11103` and `L11104`/`L11105`) before any
pixel-producing work, reading it as a shadeObj pointer (`TERR-LIGHT-059`'s own name for this
argument, in its ordinary, working caller). Because `DrawText` supplies arg5 as the constant `0`,
this dereference reads `[0x00000008]` on every call — an access violation under ordinary Win32
memory protection (the low 64KB is never mapped), unconditionally, on both branches, before a single
pixel is touched. `DrawText` therefore cannot be the routine that draws font4 glyphs in working
shipped code: the two draw routines split by sprite class, not merely by convention — `DrawText`
draws `.16`-class fonts (font1-3) only, and font4 is drawn, where it is drawn at all, through the
second draw routine instead (`R1854`, `TEXT-067`)

**Confidence.** **High.** Both halves of the crash argument are direct instruction reads (the
constant push, and the unconditional dereference on both of `R1547`'s branches,
`evidence/d-text.txt`/`evidence/d-R1547.txt`), and the consequence follows from ordinary Win32
memory protection, not a probabilistic model. This experiment did not observe a crash at runtime —
no game state was run — but a deterministic access violation on every call needs only the two reads
above to state, not a runtime witness. `TEXT-071` documents the argument-count asymmetry this same
reading also shows; this row states its practical consequence

### TEXT-066

`R1790`, the function `SPR16A-FONT-018` already names as font4's sole construction site,
calls `R1771` (the `.16a` constructor) at `L11101`, then at `L10035` calls
`R0919` with `(count=0x10, mode=4, tint=0)` — `BuildShade(16,4,0)`, the exact signature
`PAL-MODE4-010`/`SPR16A-PIX-011` already read as the `.16a` shading-LUT construction arm (no
day/night tint reference on this arm). Font1-3's construction site (`R1789`) reaches neither
call; those fonts' sprite objects carry a NULL palette pointer (`SPR16A-FONT-013`)

**Confidence.** **High** for the instruction read: both calls and their literal pushed arguments are
quoted in `evidence/d-L10035-owner.txt`. The interpretation of mode 4/tint 0 as "the `.16a` arm" is
inherited from `PAL-MODE4-010`

### TEXT-067

The second draw routine's class split is now settled at both ends: font1-3 (`.16`-class) reach a
bare no-op through this routine's vtable slot — it draws literally nothing for those fonts — while
font4 (`.16a`-class) reaches the published alpha compositor, argument-count-consistent, and (per
`TEXT-065`) is this experiment's confirmed font4 draw path, not merely its best candidate.

`R1854`'s glyph write is the virtual call through slot offset 0x18 (vt+0x18), not `+0x34`, after pushing
exactly five stack dwords (the same shape as `DrawText`'s own call). For the `.16` class this slot
is `R1857`, whose entire disassembled body (read for this correction,
`evidence/d-R1857.txt`) is one instruction at `R1857`, a return popping 0x14 argument bytes — a bare return popping its five
stack args and doing nothing else; font1-3 draw no pixels at all through this routine. For the
`.16a` class the same slot is `R1785`, which `SPR16A-078` (promoted) already reads selecting
`R1518` (forward) or `R1784` (reversed) — the exact compositor `SPR16A-ALPHA-025`
names for the `.16a` alpha blend; its own body (already disassembled and committed by `EXP-0356`,
`experiments/EXP-0356-sprite-decoder-traces/evidence/receivers.txt`, read here, not regenerated)
ends with a return popping 0x14 argument bytes on both branches, popping exactly the five dwords `R1854` pushes. Given
`TEXT-065`'s finding that `DrawText`'s font4 dispatch faults unconditionally, this routine is no
longer merely argument-count-consistent among two candidates — it is the only one of the two draw
routines that can produce font4 pixels in working code

**Confidence.** **High** for both halves. The `.16` half is now a direct, complete disassembly of
the entire callee body (one instruction), not an absence/Unknown. The `.16a` half inherits
`SPR16A-078`'s own instruction read, this experiment's own read of `R1854`'s dispatch
instruction, and the cross-check of the 0x14-byte return from `EXP-0356`'s already-committed evidence

### TEXT-068

Full bodies of `R0767` and `R1854` (`evidence/d-text.txt`) show exactly one
pixel-producing call per ordinary per-character loop iteration, one destination coordinate pair per
call. The tilde branch is a different call in the same slot, not a second call alongside it: for a
single `~`, the draw routine's per-character glyph call is replaced by one call to `R0769` (a
Bresenham line routine, already published — `TEXT-TILDE-009`), drawing a horizontal rule instead of
a glyph, and the advance step is skipped for that character — still exactly one pixel-producing call
per character, never two. A caller wanting a drop-shadow effect needs two separate `DrawText` calls
at different colours/offsets; neither draw routine provides a second pass on top of its own single
call, ordinary or tilde

**Confidence.** **High** — a full read of both routines' complete bodies settles a structural
absence rather than bounding a search, and `TEXT-TILDE-009` already reads the tilde branch as
instead-of, not in-addition-to (it jumps past the advance block rather than falling through to it)

### TEXT-069

Searched: `supersampl` (0 rows), `anti-alias` (one claim ID, listed twice — once in its ledger, once
in `retracted.md`, as every retracted row is), `sub-pixel` (the same claim). The one directly
on-point claim, `SPR16A-PIX-009` ("the low byte is only anti-alias/sub-pixel coverage"), is
retracted; its own overturn argues for the destination-blend model `TEXT-065` also relies on, not a
coverage/AA model. What can look like smoothing is (a) the pre-graded intensity levels already baked
into the `.16` atlas at author time (`SPR16A-FONT-013`), and (b) font4's destination blend
(`TEXT-065`), a general sprite-compositing mechanism reused for text, not a text-specific AA
algorithm

**Confidence.** **Medium** — the absence is bounded to the three search terms named and the two draw
routines' own bodies read in full, not a whole-image search for every possible filtering
implementation

### TEXT-070

Colour/shading reaches a font4 pixel through a table-pointer selector, not a level shift — the
shift-left-by-9 level addressing this experiment previously attributed to font4's draw-time argument
belongs to `R1547`, which `TEXT-065` now shows is not font4's actual draw path; font4's real
receiver (`R1785`) uses a different mechanism entirely.

Font1-3: unchanged — a caller-selected 16-entry ramp from a small fixed set (thirteen ramps,
`SPR16A-FONT-013`), built by `R0572` — whose only static caller image-wide is `R1242`
(`evidence/callers.txt`), not either draw routine or the font constructor, so ramp construction is a
subsystem separate from drawing. Font4, drawn via `R1854`→`R1785` (`TEXT-067`): its
own fourth argument is a table-pointer override — a zero test, a branch on zero and a copy takes the
caller-supplied pointer if it is non-zero, otherwise (the zero-branch target) a load of `+0x1c` falls
back to a pointer stored on the sprite object itself, `this+0x1c`. There is no `SHL` or comparable
shift anywhere in `R1785`'s body — the level-shaped, bit-shifted table-row addressing this
experiment previously described belongs to `R1547`, a function `TEXT-065` now shows
`DrawText` cannot safely reach for a font4 object, and no known caller in this experiment's evidence
reaches it with a font object either (its established caller is `TERR-LIGHT-059`'s ordinary
lit-sprite path, unrelated to text) — so that level-shift description does not describe font4's
actual draw-time colour mechanism. What `this+0x1c` holds is not established here: it is a distinct
field from the `this+0x14` shading table `R0919`/`TEXT-066` builds (the address of `+0x14` taken in
`R0919`, not `+0x1c`) — whether font4's default draw-time table is the `BuildShade` output or
a different field is Unknown. No claim in the corpus describes a distinct low-resolution text/UI
plane later stretched, and neither draw routine nor `R1250`/`R1785` computes a second
scale factor for x/y before handing coordinates onward — text composes directly into the
framebuffer's own channel-mask format, dispatch-layer stride only

**Confidence.** **High** for the ramp-vs-table-pointer split, the sole-caller census, and
`R1785`'s own argument-4 mechanism (each a direct, complete disassembly read). Explicit
**Unknown**, newly scoped: what `this+0x1c` holds and whether it derives from `BuildShade`'s
`this+0x14` output. **Medium** retained for "no separate scaling stage" (the pixel writers
themselves were only partly read)

### TEXT-071

The argument-count mismatch this experiment first found between `DrawText`'s font4 dispatch and its
resolved callee is not a benign miscount: `R1547` unconditionally dereferences the exact
argument `DrawText` supplies as a literal zero, on both of its branches, before any pixel work — an
unconditional null-pointer-plus-8 fault, not merely a stack-argument-count anomaly.

`DrawText` (`R0767`) pushes exactly five stack dwords before the virtual call through slot offset 0x34 at
`L11099`, the first pushed (its own arg5) a literal `0` (`L11100`). For font1-3 the resolved
callee `R1250` never reads arg5 at all — harmless. For font4 the resolved callee
`R1547` reads arg5 as a shadeObj pointer and dereferences `arg5+8` unconditionally on both
branches (`L11102`/`L11103`, `L11104`/`L11105`) before any pixel-producing work —
with arg5 forced to `0` by `DrawText`, this reads `[0x00000008]`, an access violation under ordinary
Win32 memory protection, on every call, both branches. The second draw routine `R1854` also
pushes exactly five stack dwords before the virtual call through slot offset 0x18 at `L11106`; for font4 the
resolved callee `R1785` (already disassembled and committed by `EXP-0356`) ends with a return popping 0x14 argument bytes on
both branches, popping exactly the five pushed, and its own fourth argument is tested for zero
before use (a zero test and branch) rather than dereferenced unconditionally like `R1547`'s arg5
— safe against the same kind of caller-supplied zero. So `DrawText`'s font4 dispatch is not merely
argument-count-inconsistent, it is a deterministic fault, while the second draw routine's font4
dispatch is both argument-count-consistent and unconditional-dereference-free. This experiment now
states positively that `DrawText` draws `.16`-class fonts only and `R1854` draws the
`.16a`-class font (font4); the open question narrows from "which of the two routines draws font4" to
"which callers/screens invoke `R1854` with font4 active" — no caller-side census of
`R1854` was run

**Confidence.** **High** for the argument-5 dereference and its consequence, and for the resolved
routine split (`DrawText` draws `.16` only, `R1854` draws `.16a` only) — each rests on direct
instruction reads cross-checked against Win32 memory-protection semantics, not a probabilistic
count. Explicit **Unknown**, narrowed: which callers/screens invoke `R1854` with font4 active
— that needs a caller-side census this experiment did not run

### TEXT-078

- `R0767`, the draw slot `vt+0x14` (`TEXT-API-007`), takes x, y, the string, a flag word
  and a ramp (a return popping 0x14 argument bytes). It tests the flag byte at `L06509`, `L11107`, `L11108` and
  `L11109`, one value each, and reads no other bit. The second draw `R1854`
  (`TEXT-067`) repeats the four tests at `L11110`, `L11111`, `L11112` and `L11113` with
  the same operations.
- The tests are independent, so the effects add.

| Flag value | Effect on the origin |
|---|---|
| 1 | x = x - width of the string (`R0766`) |
| 2 | x = x - (width >> 1) |
| 4 | y = y - `vt+0x20(0)` of the glyph sprite |
| 8 | y = y - (`vt+0x20(0)` >> 1) |

- Slot `+0x20` of both sprite class tables, `L10048` (font1-3) and `L07521` (font4),
  read as raw dwords, is `R1787`. It returns the first dword of frame 0's record, its
  width. Slot `+0x24` is `R0817`, which returns the second dword, its height.
  `R0767` uses the height getter for the advance of a space (half of it) and for the
  underline row (`TEXT-079`), so the two vertical anchors take the width where the height would
  be expected.
- Font1 frame 0 measures 16 wide and 15 high on both roots (`evidence/metrics.tsv`). Value 4
  subtracts 16 and value 8 subtracts 8; the height would give 15 and 7.
- Value 10, the anchor of the dialogue button label (`DLG-BUTTON-039`), centres the string on x
  by its measured width and puts the top of a font1 string's cell at y - 8, so the cell spans
  rows y - 8 to y + 6.
- `R0571(x, y, string, anchor, ramp, d)` calls the font's draw slot twice with the same
  anchor: at (x + d, y + d) with the shadow ramp that the font's slot `+0x18` returns
  (`R1278`, `L03613`), then at (x, y) with `ramp`. Each call applies the anchors
  itself, so the shadow keeps the ink's alignment. `TEXT-068` stands: the offset pass is the
  caller's second call, not a pass inside the routine.

**Confidence.** High. Both routines are read whole, each test and both getters are named
instructions, and the class-table slots are raw dwords. The font1 frame 0 size is a decode of
`font1.16` on both roots.

**Unknown.** Which callers pass which flag values was not enumerated; the dialogue's own values,
0 for the text lines (`DLG-LINE-038`) and 10 for the button label, are the only ones established
here.

## Root corpora

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-ROOT-014 | On the text path the two installs are not variants of one corpus — they are two corpora, and the executable is not the difference. | High / Medium | ● active | [EXP-0103](../experiments/EXP-0103-ru-text-domain/) |
| TEXT-ITEMNAME-019 | Item names are a text corpus like any other, and `TEXT-ROOT-014`'s finding holds on them exactly: the two roots ship two item-name corpora and one key file. | High | ● active | **[EXP-0142](../experiments/EXP-0142-item-names/)** |
| TEXT-PATTERN-020 | The Russian root ships two nodes that look exactly like a name-composition table, and `rom.exe` names neither, so the composition model they support cannot be the live one. | High / Medium | ● active | **[EXP-0142](../experiments/EXP-0142-item-names/)** |
| TEXT-BATTLEROOT-063 | An owner-supplied pre-release root's battle-event corpus is confined to three missions; where it shares a scene, speakers differ only by same-length substitution, while EN and RU differ in all three ways. | High / Medium | ● active | [EXP-0393](../experiments/EXP-0393-inn-text-closure/EXP-0393.md); extends TEXT-ROOT-014 |

### TEXT-ROOT-014

`rom.exe` is byte-identical, sha256
`942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03` on both, verified here by hashing
both files rather than taken from `TEXT-LANG-002`. Over the union of nodes with a text extension —
**no heuristic in this measurement at all**, membership is the extension and the verdict is a
payload SHA-256 — there are **437** nodes: **422 differ in bytes, 12 exist only on RU, 1 only on EN,
and 2 are identical.** The two identical ones are `main.res/text/cutpaths.txt` and
`main.res/text/heropicture.txt`, both configuration lists rather than prose; the EN-only one is
`graphics.res/version.txt`. The 12 RU-only nodes are three `text/battle/mNNN/eventNN.txt`, seven
`text/inn/npc/npc*.txt`, `text/material.txt` and `text/pattern.txt`. Over **all** nodes: 3 151
identical, 898 differing, 18 EN-only, 14 RU-only. This is the text-path counterpart of `AI-ROOT-049`
and it is a stronger result in the same direction: where the map path differed in a few files, the
text path shares essentially nothing. A consumer cannot treat one root's text as a translation of
the other's, and cannot key a mission's text off a node list taken from either root alone

**Confidence.** **High** for every count: the node sets come from the container walk already closed
in `claims/res.md`, membership is by extension, and every verdict is a SHA-256 of the payload on
each root (`evidence/diff-nodes.csv` carries both hashes for all 932 rows) / **Medium** for "reaches
something a player sees": the RU-only nodes sit under `text/battle/` and `text/inn/npc/` beside
nodes that plainly do, but no reader was traced to any of them this round

### TEXT-ITEMNAME-019

`main.res:text/itemname.txt` is 7 703 bytes and 416 lines on the EN root and 9 582 bytes and 416
lines on the RU root, and **0 of the 416 lines are identical** while `text/itemname.bin`, the 416
`u16` keys that pair them, is **byte-identical** (sha256
`adb09ccc7a210ead12812901f847e299747138030cdd085addb1a1895a7c1e24`). The RU corpus carries 8 035
high bytes over its 416 lines and the EN corpus **0**; the longest RU line is 44 bytes against 29
EN. Every one of those high bytes therefore reaches the screen through `TEXT-CONV-001`'s converter
under selector 1, which `TEXT-LANG-002` derives from `main\id` and not from the executable, and
`TEXT-CAP-018`'s 160-of-224 ceiling applies to item names unchanged. The store the display actually
reads is `ITEM-DISPNAME-036`'s map; `world.res:data/data.bin`'s own item row names are **identical**
on the two roots and are not the localisation seam (`DAT-ITEMNAME-010`)

**Confidence.** High: two node payloads compared byte for byte on both roots and a line-by-line
comparison of a complete 416-line corpus; the onward clause about the converter is inherited from
the TEXT ledger rather than re-derived here

### TEXT-PATTERN-020

`MAIN.RES` carries `text/pattern.txt` (14 144 B) and `text/material.txt` (143 B) which the EN
`main.res` does not. `pattern.txt` is an `English=English` table whose left column is the
composed-name form (`Bronze Plate Bracers=Bronze Plate Bracers`) and `material.txt` is the 15-line
English material list with `Wood` written `Wooden`, which is the very substitution the Armor and
Shield constructors' `R1858` performs. Together they are a complete composition worksheet.
The refutation is a **raw byte scan of the whole image**, not an xref sweep: the byte strings
`pattern.txt` and `material.txt` occur **nowhere** in `rom.exe`, which is byte-identical on the two
roots, so no build of the shipped executable can open either. They are localisation tooling left in
the archive

**Confidence.** High for the absence, the instrument being a raw scan of the whole file for the
literal rather than any sweep that can be blind to orphan code, a vtable or a static initialiser /
Medium for `pattern.txt` being a worksheet rather than data for some other consumer: nothing was
found that reads it, which is not the same as establishing what it was for

### TEXT-BATTLEROOT-063

A fourth, owner-supplied pre-release data root's battle-event corpus is not a subset of EN's or RU's
shape — it is confined to three missions end to end — and where a scene is shared with that root its
speaker assignment always differs by same-length substitution, never by reorder or length change; EN
and RU disagree with each other in all three ways.

Reproducing `TEXT-ROOT-014` from an independent registry+resource walk: EN 225, RU 228, RU-only
exactly `m100/event09`, `m130/event07`, `m150/event10` (matching that row's published count, now
named). The pre-release root's 23 files sit entirely inside missions 41, 51 and 91; every other
mission's battle events (all 225/228 of them) are absent from it, and three of its own 23
(`m41/event10`, `m41/event11`, `m41/event12`) exist on neither preserved install. Of 231 total rows
in the three-root union, 18 show a different `<npc=..>` tag sequence between roots that both ship
the file. Split by which roots are being compared: the 5 rows where the pre-release root disagrees
with EN or RU are all same-length substitutions, never a reorder or a length change (e.g.
`m91/event05`: EN/RU `25,23,21`, pre-release `21,22,21`). The remaining 13 rows are EN-vs-RU
disagreements, where the pre-release root ships nothing to compare, and there all three kinds occur:
3 are pure reorders of the same speaker multiset (e.g. `m10/event02`: EN `51,51,21`, RU `21,51,51`),
5 are same-length substitutions (e.g. `m131/event10`: EN `62,62,62,21,62` vs. RU `74,74,74,21,74`),
and 5 are outright length changes — RU's own sequence has a different element count from EN's, not
merely a different order or value (e.g. `m101/event06`: EN 1 element `21`, RU 2 elements `44,21`;
`m90/event06`: EN 3 elements, RU 4; `m70/event02`: EN 12 elements, RU 14; `m100/event05`: EN 10, RU
9 — the sole case in either direction where EN is longer)

**Confidence.** High for the file-set counts and the RU-only three-file reproduction of
`TEXT-ROOT-014` (both are closed-population node censuses) / High for the
same-length-substitution-only pattern on the pre-release comparisons and for the
reorder/substitution/length-change split on the EN-vs-RU comparisons — this is the `<npc=..>` markup
itself (`DLG-MARKUP-007`'s own vocabulary), directly parsed, not inferred, and the classification
(reorder = same multiset; substitution = same length, different multiset; length change = different
element count) is a mechanical comparison, not a reading / **Medium** for what a substituted,
reordered or lengthened speaker sequence means in play: no scene was traced to a consumer, only its
markup was compared

## String tables

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-STRTAB-023 | Sixteen text files are loaded into one shared line-pointer array, in a fixed order, and that order *is* the index space. | High | ● active (partially retracted) | [EXP-0143](../experiments/EXP-0143-actor-names/), amended [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-NAMETAB-026 | Which of the sixteen tables localise, measured line by line on both roots — and the answer is *almost all of them*, with the exceptions naming themselves. | High / Medium | ● active (amended) | [EXP-0143](../experiments/EXP-0143-actor-names/), amended [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-STRTAB-030 | `R1773(this,id)` is a two-field array-index accessor that returns an address, not a Win32 `STRINGTABLE` lookup, and no STRINGTABLE resource is read on the tip popup's caption path. | High | ● active | [EXP-0198](../experiments/EXP-0198-tip-popup-presentation/) |
| TEXT-STRTAB-031 | A caller census of the index accessor corroborates `TEXT-STRTAB-023`'s one-shared-array reading; a 2-site gap between two scan instruments resolves to two further thiscall thunks on the same object. | Medium | ● active | [EXP-0198](../experiments/EXP-0198-tip-popup-presentation/) |
| TEXT-UI-032 | The sixteen table paths and bases are shared, but the complete line-array length is root-specific: EN 1,568, RU 1,527. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-033 | The selected global and local string accessors have no unclassified static consumer or stored callback in the owned image. | Medium | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |

### TEXT-STRTAB-023

`R0661(this, node)` reads the node's whole payload into one heap buffer, then walks it: it
appends the current line's start pointer to a growing pointer array at **`[L04369]`** (count
`[L11114]`, capacity `[L11115]`, grow-by `[L11116]`, doubling through
`R0747`/`R0525`), scans forward to the next `CR`, overwrites that byte with a NUL **in
place**, and advances the cursor by **2**; the loop ends when the cursor passes the buffer end, and
the table object keeps how many lines it added at `+0x08` and the index it started at at `+0x0c`.
Nothing validates a line, nothing counts them in advance, and no code-page pass runs at load — the
bytes are used exactly as shipped. `EnumRefs callto:R0661` returns **16 hits / 2 owners / 0
orphan**: fifteen consecutive calls in `R0326` at `L11117`…`L11118` and one in
`R1551` at `L11119`, each with its own `this` and its own path literal, read off the `PUSH`
of every site — in order, `main\text\` `main.txt`, `heropicture.txt`, `stats.txt`, `spells.txt`,
`spell.txt`, `dialogs.txt`, `unitname.txt`, `building.txt`, `itemname.txt`, `sites.txt`,
`npcnames.txt`, `cutscene.txt`, `cutpaths.txt`, `tunes.txt`, then `patch.res::patch.txt`, then
`main\text\credits.txt`. Two indexing styles coexist and a consumer must not confuse them:
`R0668(table, i)` adds the table's own `+0x0c` base, so a table's lines are numbered **from 0
within that table**; a direct load of the global `[L04369]` and of its entry `4·k` uses the **global**
subscript, and `main.txt` is loaded first so its own line numbers are the global ones. Measured
separately on both roots, the bases run `main.txt` 0, `heropicture` 274, `stats` 300, `spells` 350,
`spell` 374, `dialogs` 402, `unitname` 568, `building` 649, `itemname` 715, `sites` 1131, `npcnames`
1150, `cutscene` 1245, `cutpaths` 1259, `tunes` 1273, `patch` 1294, `credits` 1361. ~~The sixteen
files hold 1,568 lines on both roots.~~ **EXP-0228 refutes that total alone:** `credits.txt` has 207
EN lines and 166 RU lines, so the arrays hold **1,568 EN / 1,527 RU**; every earlier table count and
base stands (`TEXT-UI-032`). Three limits a consumer inherits: a file **must** end every line with
`CR` and the byte after it is skipped whatever it is (the scan has no end check, so an unterminated
last line runs off the buffer); the index space is **positional**, so inserting a line renumbers
everything after it *in every later file*; and no reader bounds-checks a subscript

**Confidence.** **High.** The loader is one routine read whole, the call-site count names its
instrument with 0 orphan, every path literal is read off its own `PUSH`, and the load order fixes
each root's bases arithmetically from its own line counts. The old equal-total clause was High and
wrong; it is recorded in [`retracted.md`](retracted.md). `main.txt` at 0 remains a read of the call
order

**Amended.** EXP-0228 refutes the both-root total of 1,568 lines alone (`TEXT-UI-032`;
[`retracted.md`](retracted.md), REFUTED). The loader, the file order, the bases, the array structure
and the two index forms stand.

### TEXT-NAMETAB-026

Per file, lines that are byte-identical `en` to `ru` (`evidence/string-tables.txt`, corrected by
EXP-0228): `spells` 0/24, `spell` 0/28, `building` 0/66, `itemname` 0/416, `sites` 0/19, `tunes`
0/21 — **completely** translated; `stats` 4/50, `npcnames` 16/95, `dialogs` 19/166, `main` 17/274,
`unitname` 42/81, `patch` 46/67 — translated except for structural, empty and format-only lines; and
two that are **not localised at all**, `heropicture` 26/26 and `cutpaths` 14/14, both of which hold
resource paths rather than prose (`HERO-APPEAR-052` already reads the first as a path list).
`credits` has unequal lengths: **15 equal positions of 166 paired, plus 41 EN-only positions**
(`TEXT-UI-032`), replacing the ambiguous old `15/207` ratio. The discriminator is not the ratio but
*what* stays identical: in `unitname` the 42 identical lines are exactly the 42 empty ones and all
39 that carry a name differ (`UNIT-NAMETAB-041`), and in `main.txt` the 17 identical lines are seven
structural ones plus the ten hall-of-fame names (`FAME-DEFAULT-008`). So a table that mixes
translated and untranslated lines is doing so deliberately, per line, not by omission. G1
consequence, stated plainly: **every word the program chooses about a person or a thing is authored
in these files and moves with the root** — the one naming surface that does not is the executable's
own default-hero literals (`SESS-DEFNAME-032`)

**Confidence.** **High** for the census — sixteen files, both roots, complete, compared positionally
on the loader's own split with unequal tails reported separately / **Medium** for calling
`heropicture` and `cutpaths` "paths rather than prose", which is read off their contents and one
prior claim, not off a consumer

**Amended.** EXP-0228 corrected the `credits` count: 15 equal positions of 166 paired plus 41
EN-only positions (`TEXT-UI-032`) replace the old `15/207` ratio.

### TEXT-STRTAB-030

Read whole (`evidence/disasm-strtab-loader-R1773.txt`, a return popping 4 argument bytes): the routine loads the object's own `+4` field (`this` is the saved receiver at frame offset −4), loads the sole stack argument `id`,
computes `+4field + id*4` by an address calculation and returns that address; the routine never
dereferences it, and every caller dereferences the returned address itself (a 32-bit load through it). For
the fixed table object `this=L09841` that the popup's two captions use, `+4` reads
`[L04369]`, the address `TEXT-STRTAB-023` already publishes as the base of the one shared
line-pointer array that `R0661` fills from sixteen text files. `TOWN-185` (`EXP-0197`)
recorded that resolving the two captions would need the Win32 `STRINGTABLE` format decoded, and that
this experiment did not build it; no STRINGTABLE resource exists on this call path. The id is a
plain subscript into the array `TEXT-STRTAB-023` already documents, and `TOWN-206` (`EXP-0198`)
resolves both captions through it, on both roots

**Confidence.** **High.** The routine is four instructions, read in full, with an unambiguous
address-calculation-only (no-dereference) return shape and a return popping 4 argument bytes matching its one stack argument. The alias
of `this+4` to `TEXT-STRTAB-023`'s own array base at `[L04369]` is a direct address match on the
fixed global `this=L09841`, not an inference

### TEXT-STRTAB-031

A caller census of the index accessor independently corroborates `TEXT-STRTAB-023`'s
one-shared-array reading, and a 2-site gap between two scan instruments resolves to two further
thiscall thunks on the same object, not a defect in either scan.

Ghidra's own reference index (`ScanField sym:R1773`,
`evidence/scan-R1773-callers/SCAN_SUMMARY.md`) finds **21** call sites to `R1773` through
`this=L09841`. A text-match scan for the immediate operand `L09841` (`ScanField imm:L09841`,
`evidence/scan-strtab-imm-L09841/SCAN_SUMMARY.md`) finds **23** sites setting up that same `this`.
The 2-site difference is `R1859` and `L11120`, each a five-byte load of the constant `L09841` into the receiver immediately
followed by an unconditional `JMP` (`evidence/disasm-strtab-candidate-loaders.txt`,
`evidence/disasm-strtab-thunk-L11120.txt`), to `L11121` and `L11122` respectively, neither
of which is `R1773` (`R1773`). `R1859` is called directly as `R1859`;
`L11120` is registered through `atexit`-shaped code at `L11123` (the push of `L11120` then
the call to `R1860`, a routine that converts its own argument call's return into `0` on success or
`-1` on failure). Both extra sites are thiscall setups for the same object reaching a method other
than the index accessor through a short forwarding thunk. This corroborates, from an instrument
independent of `TEXT-STRTAB-023`'s own `EnumRefs callto:R0661` load-site enumeration, that the
object at `L09841` is a small, fixed-shape struct reached through a handful of named accessor
and lifecycle routines rather than reconstructed ad hoc at each call site

**Confidence.** **Medium.** Corpus agreement between two scans is capped at Medium by this ledger's
own confidence rule. The 2-site discrepancy was traced to two disassembled, byte-confirmed thunks
rather than left unexplained, which is evidence for that resolution; the underlying shared-array
reading is `TEXT-STRTAB-023`'s own High-confidence claim and is not re-graded here

### TEXT-UI-032

The first fifteen tables have equal counts and bases in both roots. `credits.txt`, loaded last at
base 1361, has 207 EN lines and 166 RU lines. Of the 166 paired positions, 15 are byte-identical and
151 differ; EN has 41 additional trailing positions. Every EXP-0228 requested consumer addresses an
earlier table, so the correction changes no selected index

**Confidence.** **High.** Both complete archives were parsed independently with the loader's CR
split; hashes and every row are preserved in `input-manifest.json`, `table-summary.tsv`, and
`strings.tsv`

### TEXT-UI-033

The repaired 8,005-function project finds global accessor `R1773`: 21 direct calls / 8 owners /
0 orphan; local accessor `R0668`: 302 / 57 / 0; line-array address `L04369`: 230 references
/ 47 owners / 0 orphan. The table object's 23 immediate references resolve to the 21 accessor calls
plus the lifecycle thunks at `R1859` and `L11120`. A whole-file dword scan finds zero
stored pointers to either accessor, and neither occurs in an `.rdata` slot

**Confidence.** **Medium.** The named call, reference, orphan, pointer, callback, initializer, and
vtable representations are complete. A runtime-computed target or bulk copy that never stores either
accessor address remains outside the negative

## UI label consumers

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-UI-034 | The shared character/unit panel selects its displayed actor name as `unitname.txt[typeID]`. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-035 | The same panel's persistent statistic captions come from fixed global `main.txt` slots, not `stats.txt`. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-036 | The character/unit panel supplies its own numeric grammar and gates caption groups by visibility level. | High / Medium | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-037 | `stats.txt` has exactly three accessor calls, all in the item-description formatter, and none in the character/unit panel. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-038 | The item-description formatter uses ten executable-authored format literals around its selected `stats.txt` labels and values. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-039 | The shop hover consumer maps four stock rectangles to global slots 62..65 and the shopkeeper rectangle to slot 61. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-040 | The shop identify modal draws slots 79 and 80 as two raw lines, so the `%d` in slot 80 is literal on this consumer. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-041 | The shop control consumer chooses slots 72, 70, 71, and 73 for Undo, Buy, Sell, and Exit. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-042 | The client opcode `0xbe` save-acknowledgement arm selects global slot 203 for a town notice. | High / Medium | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-043 | Mission outcome panels select global slots 140 and 141: success is `Mission Completed` / `Миссия выполнена`, failure is `Mission Failed` / `Миссия провалена`. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-044 | The mission Esc menu initializes its visible controls from local `dialogs.txt` slots 34, 35 or 76, 36, 37, 38, 39, and 40. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-045 | The town Esc menu uses local `dialogs.txt` slots 34, 35, 37, 77, and 40. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-046 | The world map uses three localized index forms. | High / Medium | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |
| TEXT-UI-047 | The selection status uses two fixed lines for root-specific grammar. | High | ● active | [EXP-0228](../experiments/EXP-0228-ui-text-indices/) |

### TEXT-UI-034

In `R0877`, `[actor+0x20]` supplies the local index, object `L11124` selects
`unitname.txt`, and `L11125` calls the local accessor before the draw. All 81 positions are
stable across roots; every non-empty name is localized

**Confidence.** **High.** The index load, object identity, accessor call, and draw are one retained
instruction path, and both roots were extracted positionally
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration. The agreement between the per-root
data stands as corroboration: the `unitname.txt` strings are separate files in each root.

### TEXT-UI-035

Slots 15..18 are the primary statistics; 19..20 Health/Mana; 21..22 Sight/Speed; 23..26 damage,
absorption, attack, defence; 27..28 group headings; 30..34 weapon skills; 36..40 magic schools;
41..45 resistances; 35 Weight; 46 XP; 190..191 the conditional spellcaster and armour-piercing
flags. The Russian root supplies the same indices with Russian bytes

**Confidence.** **High.** `R0877` was retained through its terminator and every selected slot
is independently reproduced from both archives

### TEXT-UI-036

Primary statistics require level above 4; Sight/Speed above 1; attack/damage above 2;
defence/absorption above 3; skills/resistances above 6. Positive pools gate Health/Mana, and actor
kind selects the skill family. The consumer uses `%d`, `%d/%d`, `%d.%d`, and `#%s: %d-%d`; none is a
line in `stats.txt`

**Confidence.** **High** for the comparisons and format operands, read from one complete retained
consumer, byte-checked in `format-literals.tsv` / **Medium** for calling the controlling value
“visibility level”; its provenance was not decoded in this experiment

### TEXT-UI-037

Object `L04315` reaches `R0668` at `L11126`, `L11127`, and `L11128` inside
`R0964`; the local index comes from the item's encoded tag. The other local-accessor calls in
that function use different table objects

**Confidence.** **High.** The repaired-project immediate/object population and each call site's
receiver distinguish the three retained calls from same-function calls to other tables

### TEXT-UI-038

The forms are `#%s %+d`, `#%s %d`, `-%d`, `#%s: %d`, `#%s: %5.1f`, `#%s: +%d`, `#%s: +%d%%`,
`#%s %s%s`, ` %s %s%s`, and `#%s: %d-%d`. Thus the table supplies statistic names while code
supplies sign, punctuation, numeric precision, range, and suffix composition

**Confidence.** **High.** Every literal's address and raw bytes are in `format-literals.tsv`, each
push is inside the retained `R0964` range, and all 50 local labels are reproduced for EN and
RU
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

### TEXT-UI-039

The pairs are Armor/Броня, Weapons/Оружие, Magic items/Магические предметы, Scrolls, books &
potions/Свитки, книги и пузырьки, and Shopkeeper/Продавец. `R1861` computes `62+i` for the
four groups and uses fixed 61 for the shopkeeper

**Confidence.** **High.** Both call sites and their immediate/register arguments are in the complete
global-accessor population; the five slots are reproduced from both roots
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration. The agreement between the per-root
data stands as corroboration: the group-name strings are separate files in each root.

### TEXT-UI-040

`R1539` looks up `Do you want to identify this item` and `for %d gold coins?` and draws each
centered; no format call or price operand occurs between lookup and draw. Both roots ship the same
English bytes for both positions

**Confidence.** **High** for this consumer's lookup-to-draw path and both-root bytes. No claim is
made that another, unlocated consumer cannot format the same line

### TEXT-UI-041

`R1027` uses each slot twice while constructing the four controls; EN/RU pairs are
Undo/Отмена, Buy/Купить, Sell/Продать, Exit/Выход. Existing presentation evidence determines where a
separate or composed number is drawn; this row fixes the table selection

**Confidence.** **High.** All eight calls belong to one enumerated global-accessor owner and use
four fixed arguments reproduced from both roots

### TEXT-UI-042

At `L11129` it pushes 203, calls the global accessor at `L11130`, and passes
`Your character is saved` / `Ваш персонаж сохранен` to the notice routine with numeric argument
`0x7530`. The argument's unit is not established here

**Confidence.** **High** for the opcode arm, index, bytes, and numeric operand, all in one retained
path / **Medium** for the human label “save acknowledgement”, based on the arm's surrounding
save-data handling rather than a named symbol

### TEXT-UI-043

The message-handler arms construct the two panels from those fixed positions; neither phrase is
mission-script dialogue

**Confidence.** **High.** Both construction arms and fixed array reads are retained, and both roots
reproduce the selected bytes

### TEXT-UI-044

The first branch offers Save Game and Load Game; the alternate mode offers Diplomacy in their place.
The remaining entries are Game Options, Sound Options, Quest Objectives, End Quest, and Return to
Game. Russian uses the same local positions, and the embedded tilde remains hotkey markup

**Confidence.** **High.** `R0362` was retained through all eight constructors and each
receiver is the dialogs table object; both roots preserve the index sequence

### TEXT-UI-045

They are Save Game, Load Game, Sound Options, Abort Game, and Return to Game, with localized Russian
bytes at the same positions. This is a distinct five-control constructor, not a condition on the
mission menu

**Confidence.** **High.** `R0363` supplies five fixed indices to the dialogs accessor in one
retained constructor; the extracted roots agree on positions

### TEXT-UI-046

A non-zero mission payment selects global slot 262 and formats `%s: %d c`; the return entry
concatenates local `sites.txt[0]` then global slot 261; hit-tested marker `i` returns `sites.txt[i]`
only after its availability test. The fixed pairs are Payment/Оплата, Plagat/Плагат, and Return to
town/Вернуться в город; the sites table has 19 localized positions in both roots

**Confidence.** **High** for the indices, format operands, order, gate, and corpus / **Medium** for
the visible spacing of the return composition, because this experiment did not decode the string
helper's separator behaviour

### TEXT-UI-047

Zero selected draws slots 47 and 48: EN `No units` + `selected`, RU `Персонаж` + `не выбран`. Two or
more draws slots 49 and 50 plus the count: EN `Units` + `selected:`, RU `Выбрано` + `персонажей:`.
Exactly one selected actor takes the full unit-panel path instead

**Confidence.** **High.** The count branches and four direct global-array reads are inside one
retained routine, and the line pairs are reproduced byte for byte from both roots
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration. The agreement between the per-root
data stands as corroboration: the line-pair strings are separate files in each root.

## Hover help

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-HOVER-048 | Hover help crosses a shared 500 ms cursor-idle threshold, rather than being requested by each screen on every paint. | High | ● active | [EXP-0361](../experiments/EXP-0361-hover-tooltips/) |
| TEXT-HOVERSET-049 | The inspected base-widget family has 96 static tables and 28 distinct hover getters. | High / Medium | ● active | [EXP-0361](../experiments/EXP-0361-hover-tooltips/) |
| TEXT-HOVERCHAR-050 | Character labels and generator icons return shipped explanatory prose. | High | ● active | [EXP-0361](../experiments/EXP-0361-hover-tooltips/) |
| TEXT-HOVERROOM-051 | The same delayed route serves room hotspots, navigation and inventory areas. | High | ● active | [EXP-0361](../experiments/EXP-0361-hover-tooltips/) |
| TEXT-HOVERTEXT-052 | Hover content is a mixture of installed prose and formatted live information, not a single ready-made string per icon. | High | ● active | [EXP-0361](../experiments/EXP-0361-hover-tooltips/) |
| TEXT-HOVERPAINT-053 | Normal hover help preserves source-authored line breaks and has its own layout, separate from introductory tips. | High | ● active | [EXP-0361](../experiments/EXP-0361-hover-tooltips/) |
| TEXT-080 | Spellbook hover getter `R1207` joins a title and mana line with damage, range and duration lines and at most one caption of `main[182..187,217]`; the last caption with a value replaces earlier ones. | High | ● active | [EXP-0435](../experiments/EXP-0435-tooltip-sources/) |
| TEXT-081 | Spellbook values come from fold `R0082` over the selected actors: per cell, sums or minima and maxima of spell-record fields; captions follow level formulas for ten spell ids; in the spellbook, ids 17, 27 and 28 never fill a caption. | High / Medium | ● active (amended) | [EXP-0435](../experiments/EXP-0435-tooltip-sources/) |
| TEXT-082 | Generator attribute hover `R1211` tests five rectangles per attribute row: `main[155+i]`, `main[273]`, `%s = %d`, then `%+d` of the next-point cost and of the refund, the last two comma-grouped. | High / Medium | ● active (amended) | [EXP-0435](../experiments/EXP-0435-tooltip-sources/) |
| TEXT-083 | Map-list hover `L11131` chooses by cursor x alone: the row's description left of 300 px, then `dialogs.txt[134]`, `[135]`, `[136]` per column; no map field selects 135 or 136. | High | ● active | [EXP-0435](../experiments/EXP-0435-tooltip-sources/) |
| TEXT-084 | The inherited hover hint at `+0x3c` has writers found by the traced pattern in three base constructors and setter `L11132`; 252 vslot-6 call sites exist, of which 8 pass a `dialogs.txt` lookup. | High / Medium | ● active (amended) | [EXP-0435](../experiments/EXP-0435-tooltip-sources/) |
| TEXT-085 | The 172 hint-forwarding constructor call sites pass null (63), a static address (54) or a text-table entry (55, in `dialogs.txt` and `patch.txt`); 17 final classes inherit the getter. | High / Medium | ● active (amended) | [EXP-0435](../experiments/EXP-0435-tooltip-sources/) |
| TEXT-096 | Spell record byte +9 is Max Range plus a power term, bytes +0xe and +0xf are the damage minimum and spread scaled by power/30 + 1, while word +0x10 is a duration in sixteenths; `R0625` fills all four from Data.bin spell columns. | High | ● active | [EXP-0454](../experiments/EXP-0454-hover-help/) |
| TEXT-097 | In fold `R0082` an actor with +0x7c zero is skipped before it is counted; with +0x3dc not 1 and bit 2 clear a counted actor sets flags 8 and skips the cell fold; the record level wraps as a byte (sum below 30, or 286 and up). | High | ● active | [EXP-0454](../experiments/EXP-0454-hover-help/) |
| TEXT-098 | Routine `R0807` groups a decimal string in threes from the right with commas and keeps a leading sign with the first group: `-123` is unchanged, `-1234` becomes `-1,234` and `+1234567` becomes `+1,234,567`. | Medium | ● active | [EXP-0454](../experiments/EXP-0454-hover-help/) |
| TEXT-105 | Constructor `L11133` (class `L11134`, five uses, `dialogs.txt` 111..115) passes its hint to itself and two child controls but gives the third an empty one, and sets flag 0x20, which empties the getter's answer, on one child. | High / Medium | ● active | [EXP-0489](../experiments/EXP-0489-troll-hints/) |
| TEXT-106 | The hover hit test `R0791` and its 11 other callee routines read no `+0x18` flag word and make one indirect call, so no flag takes part in choosing the control; only the chosen control's getter can answer empty. | High / Medium | ● active | [EXP-0489](../experiments/EXP-0489-troll-hints/) |

### TEXT-HOVER-048

`R0370`, called on controller `L01181` by `L11135`, samples imported `timeGetTime` at
`L11136`; `+58` accumulates elapsed time using `+5c`. `L11137..L11138` requires signed new
elapsed >=500 and <25500, and unsigned old elapsed <500. With `+9c ==0`, `+a0 !=0` and a root at
frame `+cc`, `L11139` calls `R0791`, then `L11140` calls the result's vslot `+14`, copies the
text and raises nonempty help. The normal pointer-change route `L11141` clears elapsed/visibility
and refreshes the clock; the polling branch excludes a held right button. Elapsed >=25500 hides an
existing box. `R1202` hides without clearing elapsed; frame mouse-button cases201..206 call it,
while `L11142` enters a marquee and hides help. A failed crossing is not retried merely because
the pointer stays still

**Confidence.** **High for the identified static arithmetic and branches.** The
complete181-instruction poll has no unresolved jump and is checked by two decoders. Native polling
cadence, all focus paths, the outer rendering gate and arbitrary intervening field writers are not
claimed

### TEXT-HOVERSET-049

Aligned `.rdata`/`.data` tables matching vslots08/0c/10=`R0530/R0531/L06492` and carrying30
code pointers yield21 specialized text-return bodies, six local zero-return bodies and inherited
getter `L06345` in 69 tables. The inherited getter returns the owned CString at `+3c` unless
flag20 is set; vslot18 (`L11132`) copies the supplied hint. `R0791` searches children
recursively in reverse stored order, then the current widget's screen rectangle. `R0370` asks
exactly that widget and has no parent-text retry after an empty answer

**Confidence.** **High for the measured signature population and local bodies; Medium as a UI
inventory.** Tables with different base methods, computed replacement tables and all
setter/constructor bindings for inherited hints were not exhaustively recovered. A null getter does
not prove screen-wide absence. Source-level names are not original symbols

### TEXT-HOVERCHAR-050

Shared card helper `R0916`, reached from generator, tavern and mission character getters, selects
`main.txt[155..158]` for primary attributes,159/160 for health/mana,161/162 for
damage/attack,163/164 for armour/defence,165..170 for
weight/sight/speed/skills/resistance/experience,171..175 or 176..180 for weapon/magic skills and 188
for monster weapon resistance. Its disclosure level, actor kind, ownership and rectangle gates
remain part of the result. Generator getters `R1211`, `L06320`, `R1516` and `R1154` also
bind free points273, skill pictures171..180, difficulty/template/continue/back247..255 and
name-entry256. Attribute +/- and value rectangles additionally format current values instead of
returning only prose

**Confidence.** **High for the reached getters, source indices and bounded control branches.** Both
original roots contain the same indexed populations. Native reachability in every transient state
and every inherited button hint remain unproved

### TEXT-HOVERROOM-051

`R0739` binds town mask1/4/2/8/10hex to main233/234/236/237/235. `L11143` binds school
weapon/magic icons to 171..180 through helper `R1862`'s visual-slot permutation1,2,4,3,5.
`R1861` binds shopkeeper61 and stock groups62..65. Shop shelf/backpack/table getters
`R1020/R1021/R1022` and inventory bar `R0990` return scrolling/background labels54..60,
or call the item formatter for an occupied cell; inventory gold uses 74. Character getter `R0734`
selects book/backpack/doll/menu labels8..14 and hero/portrait arrows52..53 or 121..122, then
card/item descriptions. Tavern getter `R1018` likewise splits candidate card and worn-item mask
regions. Each getter retains its active, held-item, state-mask, selection and geometric conditions

**Confidence.** **High for local dispatch, text sources and source geometry.** The table does not
turn every visible room button into a hover target, and does not establish every full-screen
occlusion/focus state

### TEXT-HOVERTEXT-052

Spellbook getter `R1207` reads names from `spells.txt` via objectL11144, formats labels
main117/118/123/124/182..187/217 with present live fields, and requires its availability bit. Card
helper `R0916` builds a monster spell list from main192 plus `spell.txt` objectL11145.
Shop/equipment getters call the existing item formatter `R0964`; world-map getter `R1863`
reads `sites.txt` only for a discovered or offered marker. Map-list getter `L11131` returns
row-owned text or `dialogs.txt[134..136]`. Paired RES measurements reproduce eight relevant tables
with equal EN/RU counts: main274, stats50, spells24, spell28, dialogs166, unitname81, itemname416,
sites19. The two independent readers agree on all1742 selected line lengths/hashes

**Confidence.** **High for the specific getters, table bindings and measured inputs.** Exact
complete item-formatter semantics remain with their existing claims; arbitrary runtime values, all
inherited hints and a claim that every line is localized are excluded

### TEXT-HOVERPAINT-053

`L11146` splits the owned text on byte23hex (`#`), counts lines and measures their maximum width.
`L11147` places a width=maxWidth+11, height=14*lines+5 rectangle above the stored cursor, shifting
left for right overflow and down at the top bound. `R0371` draws split lines through `R0571`
using font globalL02677 (`font2`, TEXT-API-007), at insets5/4 and pitch14. No word-wrap pass
occurs in these three complete routines. Introductory tip constructor `R1261` and its TipsMode
gate are a different route; the hover controller reads its own `+a0` field rather than that
preference

**Confidence.** **High for the complete normal layout routines and separate gates.** Native pixels,
all redraw/focus paths and oversized/corrupt authored lines remain unproved

### TEXT-080

Getter `R1207` converts the cursor (`L01257`, `L01258`) with `L00202` to a cell index, which is -1
when x < 0, y < 0, y >= 75 or x > 0x1c7. It reads the object at the singleton view's `+0xd0` and returns null
unless the cell is not negative and bit `cell` of that object's dword `+0x148` is set. Five CStrings are built; each `R0567` call is
MFC-style `FormatV` (`L11148`: `GetBuffer`, `vsprintf`, `ReleaseBuffer(-1)`) and replaces its destination.

| Slot | Content | Template |
|---|---|---|
| A | `spells.txt[cell]` (object `L11144`), `main[117]`, dword `+0x2cc+4*cell` | `%s#%s: %d` |
| B | `main[118]`, pair `+0x14c`/`+0x1ac` | `#%s: %d` or `#%s: %d-%d` |
| C | `main[123]`, pair `+0x20c`/`+0x26c` | same |
| D | `main[124]`, pair `+0x38c`/`+0x3ec`, each value times 0.0625 | `#%s: %5.1f` or `#%s: %5.1f-%5.1f` |
| E | one caption block, below | below |

Slots B..D are skipped when the pair's max is 0; an equal pair formats one value, otherwise the range. The
seven caption blocks run in code order 182, 183, 184, 185, 187, 186, 217 and all write slot E, so the last
block with a value wins. A block runs when its max exceeds -65535 (182, 184) or is nonzero (the rest). Each
caption's label is `main[N]` for its own number. The pairs and templates, equal value then range, are:

| Caption | Pair (min/max) | Templates |
|---|---|---|
| 182 | `+0x44c`/`+0x4ac` | `#%s: %d`, `#%s: %d...%d` |
| 183 | `+0x50c`/`+0x56c` | `#%s: +%d`, `#%s: +%d...+%d` |
| 184 | `+0x5cc`/`+0x62c` | `#%s: %d`, `#%s: %d...%d` |
| 185 | `+0x68c`/`+0x6ec` | `#%s: +%d%%`, `#%s: +%d...+%d%%` |
| 187 | `+0x74c`/`+0x7ac` | as 185 |
| 186 | `+0x80c`/`+0x86c` | `#%s: %d`, `#%s: %d-%d` |
| 217 | `+0x8cc`/`+0x92c` | as 186 |

The result is `%s%s%s%s%s` of A, B, C, D, E into static buffer `L02754`, so the lines are `#`-separated
(TEXT-HOVERPAINT-053). The two roots run the same executable bytes (SHA-256 `942e9b72...7d03`).

**Confidence.** **High.** The routine `R1207..L11149` is read whole: every Format call, its destination
frame slot, template address and operand displacement is in the evidence, and the instruction facts listed
there are checked against the decoded listing. The single-caption rule follows from all seven blocks writing
one slot before one join.

**Unknown.** Which object sits at the singleton's `+0xd0` at run time is bounded by TEXT-081; the cell
geometry behind `L00202` beyond the bounds above was not decoded.

### TEXT-081

Fold `R0082` (30 call sites) clears, on its `this`, the unit count (`+0x140`), flags (`+0x144`) and spell
mask (`+0x148`), then initialises per-cell arrays of 24 dwords: sums `+0x14c`, `+0x1ac` to 0; minima
`+0x20c`, `+0x2cc`, `+0x38c`, `+0x44c`, `+0x50c`, `+0x5cc`, `+0x68c`, `+0x74c`, `+0x80c`, `+0x8cc` to 0xffff;
maxima `+0x26c`, `+0x32c`, `+0x3ec`, `+0x56c`, `+0x6ec`, `+0x7ac`, `+0x86c`, `+0x92c` to 0; maxima `+0x4ac`
and `+0x62c` to 0xffff0001. It walks the selection map at `+0x9b8` when `+0x9c4` is nonzero.

An actor with a zero `+0x7c` is skipped silently. Otherwise the unit count `+0x140` is incremented and the
mask `+0x148` ORs the actor's `+0x18` (`L11150..L11151`). Then, if the singleton view's `+0x3dc` is not 1
and lacks bit 2, flags become 8 and only the per-cell fold is skipped. For such an actor the getter's mask test
can pass while the mana dword stays at 0xffff, so no value lines show. The actor's class string must equal
`CUnit` for the spell loop; its `+0x20` must be 0x17 or 0x18 and its `+0x18` nonzero. For each set bit `b` of
`+0x18` (0..23) the spell id is the byte at `L00199+4*b`, and `R0463(id)` builds a record. The skill byte
is the actor's byte at `+0x14a` plus the index `[[record+4]+0xc]+8`, the stat byte is at `+0x139`,
`L11152(skill, stat)` fills the record, and level `L` is skill plus stat minus 30, clamped to 0..100.

Per cell: record byte `+0xe` is summed into `+0x14c`, bytes `+0xe` and `+0xf` into `+0x1ac`; the word at
record `+0xc` is stored, not folded, into `+0x2cc` and `+0x32c`, so the last actor visited sets the mana value
shown; min and max of record byte `+0x9` and of word `+0x10` fill the 123 and 124 pairs. The seven accessors
take `L` and return a value only for their spell ids, else 0:

| Caption | Accessor | Spell ids | Value for level L |
|---|---|---|---|
| 182 | `L11153` | 24; 7, 28 | L/15+1; -(L/15+1) |
| 183 | `L11154` | 5, 16, 10, 22 | L/2 |
| 184 | `L11155` | 12, 17 | -(L/30+1) |
| 185 | `L11156` | 23 | 4L/5+20 |
| 187 | `L11157` | 27 | 4L/5+20 |
| 186 | `L11158` | 14 | min(L/20+2, 7) |
| 217 | `L11159` | 18 | L/10+3 |

The id table holds ids 1..10, 12..16 and 18..26 in 24 cells; ids 17, 27 and 28 are not in it. Caption 187
therefore cannot receive a value from this fold, and 182 and 184 receive none for ids 28 and 17. Cells
carrying captions: 182 cells 6 and 13; 183 cells 4, 7, 16, 19; 184 cell 11; 185 cell 5; 186 cell 9;
217 cell 23.

**Confidence.** **High** for the arrays, accessors, formulas, id table and cell mapping, read at instruction
level. **Medium** that the fold's `this` is the object the getter reads at the singleton's `+0xd0`: both use
the same field displacements, but the 30 callers were not all traced. Which spell each id names was not
re-derived.

**Unknown.** What `+0x20` values 0x17 and 0x18 and the `+0x3dc` bits name; the callers' firing times; whether
any state yields an actor counted and masked that fails the `+0x3dc` test; the meaning of record bytes `+0x9`,
`+0xe`, `+0xf` plus word `+0x10` beyond their use here. A second consumer of the accessors exists: twelve
calls at `L11160..L11161`, in the item formatter family (TEXT-HOVERTEXT-052), call six of the seven
accessors, including `L11157`, and format `main[182..187]`; what it shows for ids outside the book was not
read. Stores to `+0x74c` and `+0x7ac` occur only at `L11162`, `L11163`, `L11164` and `L11165`, all in
the fold (store-pattern scan of the code map).

**Amended.** `TEXT-097` corrects two clauses: in the `+0x3dc` arm the flags are assigned 8 (not ORed) and the class-name tests are skipped as well as the per-cell fold; the level of the caption formulas is the signed 0..100 clamp, while the record fill wraps below skill plus stat 30. `TEXT-096` states what record bytes `+0x9`, `+0xe`, `+0xf` plus word `+0x10` hold.

### TEXT-082

Getter `R1211` returns null unless `[[view+0x5c]+0x104]` is nonzero. For attribute row `i` = 0..3 it
point-tests five rectangles, advancing a rectangle-table pointer by 0x30 and a second pointer by 0x10 per
row, and returns at the first hit; after row 3 it returns null.

| Rectangle | Returned text |
|---|---|
| 1 | `main[155+i]` |
| 2 | `main[273]` |
| 3 | `%s = %d`: the label CString `[view+0x1c0][i]` (set from `main[15..18]`), then the value dword `view+0x1d0+4*i` |
| 4 | `%+d` of -(T(v+1)-T(v)), v the same value dword, from `R0828` |
| 5 | `%+d` of T(v)-T(v-1), from `R0829` |

`T` is `R0824` (HERO-COST-002). Both signed texts go to static `L11166` and then through `R0807`,
which keeps the string when its length is 3 or less, or 4 when the first character is `-` or `+`, and
otherwise loops, taking the right three characters and joining them to the rest with `,`. Rectangles 1, 2
and 3 return without grouping. The EN `main[155..158]` lines carry 2, 5, 3 and 4 `#` separators;
`main[15..18]` and `main[273]` carry none.

**Confidence.** **High** for the rectangle order, indices, templates, values read and the cost functions,
read in the complete getter and callees. **Medium** for the grouping output shape: the loop and constants
are read, but no output string was produced.

**Unknown.** The rectangles' pixel positions were not measured; the owner of `[view+0x5c]+0x104` was not
named.

**Amended.** `MENU-066` gives the rectangle positions, shows that the second rectangle is one record shared by the four rows, and names the owner of `[view+0x5c]+0x104`; `TEXT-098` states the grouped output. The five tests per row stand.

### TEXT-083

Getter `L11131` reads the control's screen left `L` and cursor x `L01257`. If x < L+0x12c it takes the row
under the cursor, `[control+0x84]` plus the row from `L11167(y)`, and returns its record's CString at `+8`
(the description, below) when the vslot `+0x78` validity call accepts the row, else null. If x < L+0x186 it
returns `dialogs.txt[134]`; if x < L+0x1a4, `dialogs.txt[135]`; otherwise `[136]`. No record field enters
the test.

The row text `%s#%dx%d#%d#%d` (at `L11168`) takes the record's name (`+4`), `+0x14` minus 16, `+0x18` minus
16, `+0xc` and `+0x10`. The draw routine's column advances (0x1e plus 0x10e, 0x5a, 0x1e) equal the getter's
thresholds, so 134, 135 and 136 caption the size, `+0xc` and `+0x10` columns. The loader `R0479` fills
the record from a `.alm`: it compares the tag with `M7R`, then width and height go to `+0x14` and `+0x18`,
a 40-byte seek follows, then the 64-byte name (payload `+0x30`) to `+4`, two dwords (payload `+0x70`, `+0x74`)
to `+0xc` and `+0x10`, and a 512-byte block from payload `+0x78` whose line feeds become `#` before its leading
C string is stored at `+8`. The reads run only when the `M7R` compare matches and a header dword is at least
2 (`L11169..L11170`). The loader has one return-1 path, which tests `+0xc` > 1, and three return-0 paths;
`L11171` drops a map whose loader returns 0.

The EN `dialogs.txt[134..136]` lines are 11, 26 and 20 bytes, RU 12, 32 and 28, none with a `#`.

**Confidence.** **High.** The getter and loader are read whole, the thresholds equal the drawn column advances,
and both roots share the executable. The column-to-field pairing follows draw order and arguments, not a
label.

**Unknown.** What values
`+0x74` takes beyond ALM-META-026 is outside this read.

### TEXT-084

Three base constructors, `R1219`, `R1218` (three stack parameters) and `R0773` (six), construct the
CString at `+0x3c` and set flags `+0x18` to 1. The latter two call `CString::operator=(LPCSTR)` when their
text parameter (the third, respectively sixth) is nonzero; `R1219` assigns the static `L11172`.
Setter `L11132` (vslot 6 in all 96 widget tables) assigns its argument to `+0x3c`; getter `L06345`
returns it unless flag 0x20 is set (TEXT-HOVERSET-049). The instrument's pattern finds ten CString
construct, assign or copy calls on a `+0x3c` receiver: six in the three base constructors, one in the setter
and three in other functions (`L11173`, `L06373`, `L06371`) whose receiver class was not resolved.

The call `[reg+0x18]` occurs at 252 sites image-wide with unresolved receivers. Eight sit within one routine,
`L06329`: four pairs, each passing `dialogs.txt[23]` and then `[75]` to vslot 6. No other site has a
text lookup within the 14 preceding instructions.

**Confidence.** **High** for the constructors, setter, getter and the listed sites, read from the code map
(11 788 entry points, 454 310 instructions). **Medium** for the absence of other writers: the 252 receivers
and three other functions were not classified.

**Unknown.** What the 244 other vslot-6 receivers are and what text they pass; whether a lookup stored in a
local before the call is missed by the 14-instruction window.

**Amended.** `MENU-068` names the controls behind the eight setter sites, and `MENU-069` the population of vslot-6 receivers: a backward trace resolves 9 of the 252: 8 in reached classes and 1 constructing a class with no widget table. The function the instrument names `L11173` has its entry at `L06369` and constructs class `L06370`, outside the 96 tables; `L06371` constructs class `L06372`.

### TEXT-085

From the two base constructors with a text parameter, the parameter was traced backwards through 24
forwarding constructors to 172 call sites in 72 functions that call 23 distinct constructors, by the
argument's source:

| Source | Sites |
|---|---|
| null | 63 |
| address of an uninitialised data cell | 51 |
| address of a literal in initialised data | 3 |
| `dialogs.txt` lookup | 52 |
| `patch.txt` lookup | 3 |

The 51 data cells are single addresses at four-byte strides, 50 in `L06391..L06392` and one at `L06393`;
no instruction other than the argument push references them, so they read as empty strings. The three
literals (`L11174`, `L11175`, `L11176`; 18, 21, 20 bytes, not read from a text resource) are passed
by `R1190`. The 52 `dialogs.txt` sites use 51 distinct indices; the three `patch.txt` sites (object
`L06351`, from `patch.res`, 67 lines) use 52, 53 and 54, all in `R1217`. One site (`L06348`) resolves
only by hand, to `dialogs.txt[117]`, because overlapping decodes hide its push. One site, `L11177`, builds
the map list with `dialogs.txt[118]`, and its class `L11178` has its own getter `L11131` (TEXT-083).
The other 171 sites construct 17 final classes, all carrying getter `L06345`; 29 of the 171 sites construct
the base class itself with null. These are 17 of the 69 inherited-getter tables; the other 52 are not reached
by this trace.

**Confidence.** **High** for the counts under the stated trace: the code map, table reader and operand facts
are reproduced from both roots, the two roots share the executable, and the table lines exist in both.
**Medium** that the uninitialised cells hold empty strings, since a pointer-based writer is not excluded.

**Unknown.** The population of the 52 tables not reached; which constructed controls are ever shown; whether
the literal-text controls are reachable in play.

**Amended.** `MENU-068` names the controls behind the `patch.txt` 52..54 and `dialogs.txt[117]` sites, and `MENU-069` classifies the 52 tables this trace does not reach.

### TEXT-096

`R0625(power)` runs on the record built by `R0463`/`R0624` (MAGIC-SPELL-001) and takes the power as a byte (`ebp+8 & 0xff`). Parameter `p` is the Data.bin spell column at schema title `p+1`: 6 is Max Range, 11 Area Effect Duaration, 14 Spell Duration, 16 damageMin, 17 damageMax, 18 Defensive.

- `+9`: low byte of parameter 6 (`L03041`). Id `0x1a` then adds `power / 3` (`L05074`); any other id adds `power / 30` when `+9` is nonzero (`L05073`), so a Max Range of 0 stays 0.
- `f = power / 30 + 1` as a double (`L11179`, `L11180`; constants `[L05069]` = 30.0 and `[L04115]` = 1.0).
- `+0xe`: parameter 16 above 0 gives the low byte of `ftol(p16 * f)` (`L11181`), else 0 (`L05188`). `+0xf`: parameter 17 above 0 gives the low byte of `ftol(p17 * f - byte(+0xe))` (`L11182`), else 0 (`L05190`). `+0xe` is the minimum and `+0xe + +0xf` the maximum the fold sums (TEXT-081).
- `+0x10` (word), in sixteenths of the Spell Duration unit: parameter 14 above 0 and id `0xf`: `1.05^power * 3.0 * 16.0` when below 65000.0, else 65000 (`L11183`..`L11184`; the power helper is `pow` with the base first and `R0279` the truncating conversion, HERO-COST-002; constants `[L05234]` = 3.0, `[L05072]` = 16.0, `[L11185]` = 65000.0). Parameter 14 above 0, other ids: `ftol(1.025^power * p14 * 16.0)` (`L11186`..`L11187`). Otherwise parameter 11 above 0: `(p11 << 4) + ((power << 4) / 10)` in integers (`L11188`..`L11189`). Otherwise 0 (`L11190`). The hover prints it times 0.0625 (TEXT-080).
- `R0904` (the cast path) clamps the signed power to 0..100 before the call; `L11152` (the fold's path) does not (TEXT-097).

Data.bin spell rows, parsed with `tools/placedb` from both roots (the 28-row tables are identical): the base of `+9` is nonzero for 26 ids and 0 for ids 4 and 18. `+0xe` and `+0xf` are nonzero at power 0 for ids 1, 2, 3, 6, 9, 11, 13, 14 and 21 and 0 for the other 19; parameter 17 exceeds parameter 16 in all nine. `+0x10` takes the parameter 14 path for ids 5, 8, 10, 15, 16, 18, 20, 22, 23, 24, 27 and 28, the parameter 11 path for ids 3, 7, 12, 17, 19 and 21 and 0 for ids 1, 2, 4, 6, 9, 11, 13, 14, 25 and 26.

The fold `R0082` reads `+9` plus word `+0x10` for the Range and Duration pairs (TEXT-081); `+0xe` and `+0xf` feed the damage sums. The hover value for a spell is therefore the per-cast value at the actor's own power, not the Data.bin column.

**Confidence.** **High** for every branch, formula and constant: the whole of `R0463`, `R0624`, `R0904`, `L11152` and `R0625` is listed and the sites are reproduced from both roots, which share one executable. The row populations are exact applications of the branch conditions to the parsed table.

**Unknown.** The per-id parameter values behind the two `+0x10` paths are in the private evidence only. Whether `+0x10` has a reader outside the spellbook fold in a routine that takes no `Spell*`.

### TEXT-097

The session object (the call to `R0347`, then its vslot `+0x7c`) is read once at `L11191` into a frame local; the +0x3dc test uses that one value for every actor of a fold call.

- Actor `+0x7c` equal to 0 (`L11192`, `L11193`, `L11194`): jump to `L01338`, back to the loop head. The actor is not counted (`+0x140`), not ORed into the mask (`+0x148`), not stored as the first actor (`+0x138`, `+0x13c`), sets no flag bit and runs no spell loop.
- Otherwise the actor is counted (`L01333`), its `+0x18` is ORed into the mask (`L11195`) and the first such actor is stored (`L01339`).
- Session `+0x3dc` not equal to 1 and `(+0x3dc & 2) == 0` (`L11196`..`L11197`): jump to `L11198`, which assigns flags `+0x144` = 8 (`L00702`) and returns to the loop. For that actor the class-name tests that set flag bits 1, 2, 0x20 and 0x200 and the `+0x18c & 1` bit 8 do not run, nor does the per-cell fold. The mana dword and cell arrays keep their initial values.
- Tail: flag 4 is ORed when the first actor's `+0x14` differs from `[view+0x9b4]` (`L00706`..`L00707`). A count of 0 or `flags & 0x24` posts message `0x40a` (`L11199`); otherwise `0x409` is posted with the `R0093` value (`L11200`).
- Record fill versus caption level: the fold calls `L11152(skill, stat)` for the record (`L11201`), which adds in byte arithmetic (`L11202`..`L11203`: `skill + (stat - 30)` kept in a byte) before the compare with 100 (`L11204`). A sum below 30 wraps to 226..255 and is clamped to 100; a sum of 286 or more keeps only its low eight bits, so the level is (sum - 30) modulo 256 before the clamp, for example sum - 286 for 286 to 386; the compare with 0 at `L11205` tests an unsigned byte and never fires. The caption accessors take a separate 32-bit signed level (`L11206`..`L11207`: `skill + stat - 30` clamped to 0..100). The cast path `R0904` clamps signed.

**Confidence.** **High**: every cited address is an instruction row reproduced from both roots, and the session value is stored once before the loop.
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration.

**Unknown.** What actor `+0x7c` and the `+0x3dc` bits mean. Whether any shipped character has skill plus Mind below 30 (record power 100) or at 286 or more (record power wrapped). How the single-actor branch at `L11208` (count 1 and flag 8) uses the first actor.

### TEXT-098

The routine takes a `CString` by reference and replaces it. With `S` a copy of the argument and `R` empty, it loops while the length of `S` is above 3 and not (length 4 with first character `-` or `+`) (`L11209`, `L11210`, `L11211`, `L11212`): `R = ',' + Right(S, 3) + R` (`L11213`, `L11214`, `L11215`), then `S = Left(S, len - 3)` (`L11216`, `L11217`); the result is `S + R` (`L11218`). The helpers are identified by their shapes: `R1777` Right, `R1257` Left, `R1778` character plus string, `R1776` string plus string, `R0906` assign.

A model of the loop (`group.py`) equals thousands grouping with the sign kept in front of the first group for every `%+d` and `%d` text from -1000000 to 1000000 (0 mismatches over 4000002 strings), and the table in the evidence lists `-5`, `+999`, `-999`, `1000`, `-1000`, `-1234`, `-100000` and `+1234567`. The two signed numbers of the attribute hover (TEXT-082) pass through it.

**Confidence.** **High** for the loop shape and the exit test, read in the whole routine. **Medium** for the string helpers' identities and so for the output text: the model transcribes the listing and no output was produced by running the original.

### TEXT-105

- Constructor `L11133` takes the hint as its sixth parameter (`[ebp+0x1c]`). It passes it to base constructor `R0773` for the control's own `+0x3c` (`L11219`..`L11220`), to child constructor `L11221` (`L11222`..`L11223`; class `L06366`, base `L06347`, a text edit box per `MENU-068`; stored at `+0x60`) and to child constructor `L11224` (`L11225`..`L11226`; class `L06365`, base `R0762`; stored at `+0x5c`). The third child, `L11227` (class `L11228`, base `R0700`, a button), receives the zero cell `L06393` (`L11229`), so its hint is empty. All four classes carry getter `L06345`.
- Five call sites, all in builder `L06385`, pass `dialogs.txt` 111, 112, 113, 114 and 115 (`L11230`, `L11231`, `L11232`, `L11233`, `L11234`).
- Flag 0x20 is set on the child at `+0x5c` by the constructor (`L11235`..`L11236`; vslot 7 `R1864` with mask 0x20, value 1). Three methods write the same flag on the same child: `L11237` (`L11238`) and `L11239` (`L11240`) set it, and the method containing `L11241` (`L11242`) clears it. Each is paired with a vslot 10 call (`R0777`, which writes flag 4 of the child) on that child: value 0 with each setting (`L11243`, `L11244`) and value 1 with the clearing (`L11245`).
- Getter `L06345` returns the CString unless flag 0x20 is set (`TEXT-HOVERSET-049`). The `+0x5c` child therefore answers an empty hint whenever flag 0x20 is set, and the poll raises nothing for empty text (`TEXT-HOVER-048`). There is no ancestor retry, so a point over the third child shows nothing although the control has a hint.
- Among the 24 forwarding constructors of `TEXT-085` and the two base constructors, only `L11133` calls the allocator `R1147`.

**Confidence.** High for the constructor, the five call sites, the arguments and the flag writes: each is a cited instruction in a body read whole (168, 44, 60 and 63 instructions). Medium for the child roles beyond constructor identity and for the meaning of the pairing with flag 4, which is only written through vslot 10 here.

**Unknown.** The rectangles of the three children against the control's own, which decide where each hint can show. What flag 4 means. Whether the dialog built by `L06385` is reachable in single-player play. The entry of the method containing `L11241`, which is not a code-map function start.

### TEXT-106

- The routines are `R0791`, `R1865`, `L11246`, `R1866`, `R1867`, `L03643`, `R1212`, `R1224`, `R1213`, `R1868`, `R1869` and `L11247` (65, 8, 14, 8, 9, 14, 30, 23, 17, 9, 11 and 11 instructions, no unresolved branch). No instruction reads displacement `+0x18` through a register other than `esp` or `ebp`. The one indirect call is `L11248`, through import slot `L01178`; it passes a point and a rectangle pointer (`L11249`..`L11248`) and its answer is tested for zero at `L11250`.
- The routine reads the child count from the list at `this+0x1c` (`L11251`, `R1868`), starts at count-1 (`L11252`) and decrements (`L11253`), so the last stored child is tried first. It fetches each child through `R1869` (`L11254`), recurses (`L11255`) and returns a nonzero answer (`L11256`..`L11257`). With no child hit it builds the control's screen rectangle from `this+8` (`L11258`..`L11259`) and tests the point against it (`L11260`): a hit returns this control, else 0.
- Callers are the poll at `L11139`, the routine itself (`L11255`) and `L11261` in `R0790`, the tavern click handler of `TOWN-013`. The poll makes one vslot 5 call, to the returned control (`L11140`); it has no backward branch.
- A control whose flag 4 is clear can therefore be returned when the point lies in its rectangle, and `TEXT-105` shows its hint is then removed only by flag 0x20 in the getter.

**Confidence.** High for the enumeration: 12 bodies read whole and listed in the evidence. Medium for the conclusion that no flag takes part: it covers these routines only, the poll's own gate is `TEXT-HOVER-048`'s, and a flag word passed as a parameter would not show as a `+0x18` read.

**Unknown.** The meaning of flag 4 and whether any caller filters the returned control after the poll.

## Character-generation labels

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-CHARGEN-027 | The final character generator has a mixed resource-to-control contract: three text-overlay navigation buttons, one localised caption bitmap, and two five-entry pictorial skill banks. | High | ● active | [EXP-0157](../experiments/EXP-0157-chargen-labels/) |
| TEXT-CHARGEN-028 | The earlier precreation screen binds the name prompt directly, but sex and class are one four-way pictorial choice with no separate text labels. | High | ● active | [EXP-0157](../experiments/EXP-0157-chargen-labels/) |
| TEXT-CHARGEN-029 | Character-generation localisation selects content by root, not label index, and none of the requested static labels is executable-authored. | High | ● active | [EXP-0157](../experiments/EXP-0157-chargen-labels/) |

### TEXT-CHARGEN-027

`R1870` constructs the statistic, card, navigation and skill children and initialises them at
`L11262..L11263`. Navigation construction reads global string slots **238, 239 and 260** at
`L11264..L11265`; the measured EN/RU pairs are `Accept` / `Принять`, `Reset` / `Сбросить`, and
`Back` / `Назад`. Click indices 0/1/2 call continue, reset and back at `L11266`, `L11267`,
`L11268`. The statistic panel instead loads `main.res::graphics\chrgen\leftup.bmp` at `L11269`,
blits it, then draws four formatted statistic values and the remaining-points number. The 160×238
bitmap itself carries EN `Body`, `Agility`, `Mind`, `Spirit`, `Free pool` or RU `Сила`, `Ловкость`,
`Разум`, `Дух`, `Свободные очки`; its hashes differ by root. Slots 15..18 are metadata and slots
155..158 plus 273 are hover prose, not the persistent caption draw. `R0133` constructs
exactly five fighter images (`sword`, `axe`, `Mace`, `Pike`, `Bow`) or five mage images (`fire`,
`water`, `air`, `earth`, `astral`); `R1871` draws images with no font call. Hover index `i`
alone returns slots **171..175** or **176..180** at `L11270..L11271`

**Confidence.** **High for the static draw contract.** The executable binding and event paths
discriminate an overlaid-text model from a bitmap/picture model, while the complete both-root
resource comparison and decoded bitmap provide an independent content witness. No runtime frame is
claimed; transient tooltip reachability remains unwitnessed

### TEXT-CHARGEN-028

the push of 0x464 at `L11272` and the call to `R1481` at `L07730` construct the name-entry child and `L11273`
stores it at owner `+0x1bc`, closing `TEXT-NAMEIN-024`'s Medium surface identification. The draw
reads `[L04369]+0x1f4`, global slot **125**, at `L11274..L11275`: EN `Character name:`, RU
`Имя персонажа:`. The same child's tooltip getter returns slot **256**, EN
`Type character name here`, RU `Введите здесь имя персонажа`. `R1672` loads four
`PreCreate\Heroes` image families (`mf`, `mm`, `ff`, `fm`) and `R1474` draws their states
without a class- or sex-string read. Hit testing produces one value 0..3; acceptance passes the
selected field to `R0591`, which masks `0xc0` and dispatches four character templates. The
earlier confirmation is a byte-identical `PreCreate\ButtonOk.bmp` with baked `OK`; a non-empty name
lets it emit continue message `0x445`

**Confidence.** **High for construction, draw and four-way dispatch.** A model with two labelled
selectors predicts two label sources or text draws and is refuted by the complete control
constructor and draw. The semantic gloss of what the `OK` confirms is not encoded and remains
interpretation

### TEXT-CHARGEN-029

Both roots use the same global string slots and resource paths in the same byte-identical
executable. Their active `main.res` supplies different bytes at every cited label/help slot and
different pixels in `graphics\chrgen\leftup.bmp`; the pictorial class/sex and skill resources and
`ButtonOk.bmp` are byte-identical. `R0616` opens `main\id`, reads its last character and
stores digit-minus-`0` at `[L06203]`: `english 0` → 0, `russian 1` → 1. That value gates byte
conversion (`TEXT-LANG-002`, with `TEXT-CONV-001`'s injectivity clause explicitly retracted by
`TEXT-DOM-010`); it does not branch any cited consumer to another index. Static wording is therefore
table bytes or bitmap pixels. Typed names and numbers are generated values, not labels. The seams
differ: wording is positional data; persistent statistic captions require editing a fixed 160×238
image; replacing one of five skill images is data-only, while adding a sixth choice, a third class
bank or separate sex/class controls changes the fixed consumer loops and executable behaviour

**Confidence.** **High for the source and selector mechanism.** Same-index/different-root bytes,
same-path/different-root pixels, identical executable and direct consumers discriminate the
bilingual-table and code-literal rivals. The customisation classification follows the named fixed
loops and files; no save or command format implication was examined

## Character-name entry

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-NAMEIN-024 | The typed-name field is a text-entry class with a hard ten-character cap applied *before* conversion — and `TEXT-IN-005`'s six owners are now all read. | High / Medium | ● active (partially retracted) | [EXP-0143](../experiments/EXP-0143-actor-names/), amended [EXP-0406](../experiments/EXP-0406-precreate-name/EXP-0406.md) |
| TEXT-COLL-025 | A typed name *can* reach a colliding byte at runtime: of 192 reachable glyph records, 48 collide, and 16 of those three ways. | High / Unknown | ● active | [EXP-0143](../experiments/EXP-0143-actor-names/) |
| TEXT-073 | Both traced openings of the pre-create screen load the name field from `main+0x430` after replacing `Unnamed` or `npcnames.txt` entries 20..23 with entry 20, so a new campaign without `-name` opens at EN `Danath`, RU `Данас`. | High / Unknown | ● active | [EXP-0406](../experiments/EXP-0406-precreate-name/EXP-0406.md) |
| TEXT-074 | A left-button press on hero slot 0, 1, 2 or 3 writes `npcnames.txt` entry 20, 21, 23 or 22 into the name field only when the slot differs from the last pressed one and the text is `Unnamed` or one of entries 20..23. | High / Medium / Unknown | ● active | [EXP-0406](../experiments/EXP-0406-precreate-name/EXP-0406.md) |
| TEXT-075 | The name field only appends: a converted byte at or above `0x20` goes to the end while the text is under 10 bytes, key 8 removes the last byte, and nothing selects or replaces, so the first accepted keystroke extends the seeded name. | High / Unknown | ● active | [EXP-0406](../experiments/EXP-0406-precreate-name/EXP-0406.md) |
| TEXT-076 | The caret is `\|` (`0x7c`) appended to the drawn copy of the name, in the name's font and colour; it flips at the first paint more than 500 ms after the last flip, character message under the cap or construction, with no focus test. | High / Unknown | ● active | [EXP-0406](../experiments/EXP-0406-precreate-name/EXP-0406.md) |
| TEXT-077 | The prompt `main.txt[125]` and the name are left-aligned font4 draws at origin + (224,305) and + (224,321); full-level inks are RGB(65,47,20) and RGB(101,39,61) on the normal ramp, (57,41,17) and (89,34,54) on the low-memory `/18` ramp. | High / Unknown | ● active | [EXP-0406](../experiments/EXP-0406-precreate-name/EXP-0406.md) |

### TEXT-NAMEIN-024

The class's vtable is ~~**`L07728`**~~ **`L07729`** (30 slots, `+0x00..+0x74`);
~~slot 23 (`+0x5c`)~~ slot 29 (`+0x74`) is
`R1480`, eighteen instructions read whole: `L11276` loads the string pointer at `+0x60` (the MFC
`CString`'s `m_pchData`), `L11277` compares the 32-bit length field at offset −0x8 with 0xa (its `nDataLength`) and exits on greater-or-equal →
**return 0, the keystroke is discarded**; otherwise `L11278` calls `R1517` — the input converter
— and `L11279` calls `R1872`, which sets a caret-blink pair (`+0x70 = 1`,
`+0x74 =` ~~`GetTickCount`~~ `timeGetTime`), tests the byte against 0x20 with a carry exit so **every byte below `0x20` is dropped**, and
appends through `CString::operator+=` at `L11280`. `R1873` is backspace — `Left(len-1)` —
dispatched from `R1874` on key **8** and no other. `R1154` on the same class returns
`stringTable[256]` when its owner's `+0x1e0` is non-zero, and index 256 of `TEXT-STRTAB-023`'s table
is the character-name hint. So the cap is on the **stored byte length**, and because the converter
is one byte in, one byte out, ten characters and ten bytes are the same limit in either language.
**The six owners of `R1517`** (`EnumRefs callto:R1517`: 6 hits, 6 owners, 0 orphan), each
read: `R0780` and `R1852` convert a keystroke only to compare it with `this+0x74`
(`R1852` also accepts `0x0d`) — hotkey matchers that store nothing; `R1223` converts
whole strings out of an array of **0x104**-byte records into a list control — a file browser, so one
owner really is a file read; `R1267` appends into a width-measured, wrapping multi-line
buffer; `R1258` appends any byte `>= 0x20` with **no** length cap; and `R1480` is the
capped field. That closes the Medium `TEXT-IN-005` published about itself

**Confidence.** **High** for the cap, the drop-below-`0x20` rule, the backspace key and the six
owners' shapes (every one is a cited instruction from a routine read end to end, and the enumeration
names its instrument with 0 orphan) / **Medium** that this class is the *character-name* field: the
tie is `R1154`'s tooltip index plus the corpus fact that all ten shipped hall-of-fame names
are nine characters or fewer, which is consistent with a cap of ten but does not discriminate it
from any larger cap. A screen that instantiates the class was not traced

**Amended.** EXP-0406 (`TEXT-075`, `TEXT-076`; [`retracted.md`](retracted.md), PARTIALLY
RETRACTED) corrects three words of the mechanism. The 30-slot table starts at `L07729`
(`EnumRefs vt:L07729:30`); `L07728` is its slot 6, so the entry `L11281` that holds
`R1480` is slot 29 (`+0x74`), the character slot of `TOWN-211`'s message map. The caret
stamp calls `[L00849]`, which the image's import table binds to WINMM `timeGetTime`;
`GetTickCount` is `[L11282]`, which neither `R1875` nor `R1872` calls. The cap,
the drop below `0x20`, the backspace key, the tooltip index and the six owners stand. The Medium
surface identification was closed by `TEXT-CHARGEN-028`'s `L07730` construction; `TEXT-075`
reads that construction against this vtable.

### TEXT-COLL-025

Composing the input map (`TEXT-IN-005`) with the display map (`TEXT-DOM-010`) over what a keystroke
can store under selector 1: `0x00..0x7F` passes both untouched; `0x80..0xBF` is stored unchanged;
`0xC0..0xEF` is stored as `0x80..0xAF`; `0xF0..0xFF` is stored as `0xE0..0xEF`. Stored `0xC0..0xDF`
and `0xF0..0xFF` are therefore **unreachable** through the keyboard. At the draw, stored
`0x80..0xAF` becomes records 144..191, stored `0xB0..0xBF` becomes records 144..159, stored
`0xE0..0xEF` becomes records 208..223. Enumerated over all 256 typed values: **192** distinct
records are reachable and **48** of them collide. Records **160..191** take two typed preimages
each, `0x90..0xAF` and `0xD0..0xEF`; records **144..159** take **three**, `0x80..0x8F`, `0xB0..0xBF`
and `0xC0..0xCF` — the first and third colliding at the *input* converter and the first and second
at the *display* converter, and the third is the CP1251 block carrying the first sixteen Cyrillic
capitals. So 64 typed values are lost, the same 64 `TEXT-ALIAS-011` counts, but spread over 48
records rather than 64 pairs, because the input converter folds a second pair on top of the display
converter's. Two different names can be drawn identically, with no clamp and no substitute. The two
cross-language cases fall out of the same table: **a Latin name on the Russian install is safe**
(every byte below `0x80`, both maps identity), and **a Cyrillic name on the English install** takes
the non-1 selector where both maps are the identity, so a CP1251 byte `0xC0..0xFF` is stored raw and
drawn as record `0xA0..0xDF` — inside the 224-record atlas, so a glyph appears and nothing faults.
This closes the open item `docs/status/text.md` carried, positively

**Confidence.** **High** for the composition and the reachable set (both maps are already published
at High from routines read end to end, and the composition is an exhaustive enumeration over 256
inputs with no free parameter, taken through the one storage path `TEXT-NAMEIN-024` reads) /
**Unknown** which *glyph* each colliding record draws, which is `SPR16A-FONT-020`'s question, and
therefore whether a player would notice. The prediction that would settle it, and what refutes it,
is in the experiment

### TEXT-073

The pre-create screen is `main+0x360`, constructed at `L07694` by `R1473`, which stores
vtable `L07691` (`L07692`) and runs the create routine `R0597` (`L11283`). Its opener
`R0909` has two callers (`EnumRefs callto:R0909`: 2 hits, 2 owners, 0 orphan): `L11284`
in the `0x425` new-campaign arm of `R0701`, and `L11285` in the result router
`R0709`, reached from its detailed-generation arm (`L11286`) and by fall-through from its
`+0x35c` arm when that result is not `0x446` and `main+0x490` bit 2 is clear
(`L11287..L11288`). The opener sets `pre+0x1c4` to a copy of the string at
`main+0x430` (`L11289..L11290`), `pre+0x1d0` to `main+0x65c` (`L11291`) and `pre+0x1cc` to
`main+0x494 >> 6` (`L07690`), then calls `vt+0x80` (`L11292`), the enter routine
`R1475`.

The enter clears both button groups, sets `pre+0x1cc = 0` (`L04519`), lights hero element 0
(`L11293`) and lights difficulty element `pre+0x1d0` (`L11294`). It compares `pre+0x1c4` byte
for byte with `npcnames.txt` entries 20, 21, 22 and 23 and with `Unnamed` (`L11295`); any match
replaces `pre+0x1c4` with entry 20 (`L11296..L11297`). The entries come from `R0668` on
the table object `L11298`, which `L11299..L11300` loads from `main\text\npcnames.txt`; the
index counts CR-terminated lines from 0 (`TEXT-STRTAB-023`). The enter then sets
`pre+0x1d8 = pre+0x1cc` (`L11301`), copies `pre+0x1c4` into the name field through `R1876`
(`L11302`) and sets `pre+0x1e0 = 1` (`L11303`).

On a new campaign `main+0x430` is written by `R1382` (`L09568`) and again by
`R1553` on `main+0x420` with three zero arguments (`L11304..L11305`; `EDI` is zeroed at
`R0701`'s entry, `L11306`, in EXP-0208's committed export, and the path to the switch
does not write it). Both copy the bytes after `-name` in the string `SESS-HERO-013`
reads as the command line, at most 31 bytes, cut at the first space (`L11307..L11308`,
`L11309..L11310`), or `Unnamed` when `-name` is absent (`L11311..L11312`,
`L11313..L11314`). Without `-name` the first opening therefore shows entry 20: EN `Danath`
(`44 61 6e 61 74 68`), RU `Данас` (`84 a0 ad a0 e1`). With `-nameX` it shows `X`; `-name X` stores
an empty string, which the enter keeps.

The detailed screen's Back reopens through `L11285` with `main+0x430` holding the text the
pre-create forwarded (`L11315..L11316`). The enter keeps a typed name, turns any of the five
default texts into entry 20, keeps the difficulty and resets the hero to slot 0, whatever
`pre+0x1cc` the opener copied.

`Master Oberic` (`L11317`) has one reference, `L11318` in `R0597`, which writes it into
`pre+0x1c4`. Both openers overwrite `pre+0x1c4` before the enter, so no traced path displays it.

**Confidence.** **High** for the two traced openings, the replacement rule, the seeding and the
per-root bytes: each step is a cited instruction in a routine read end to end, the caller census of
`R0909` is complete for direct calls, and the entry bytes come from both roots'
`main.res::text/npcnames.txt`.

**Unknown.** What object `R1877` returns: this experiment reads only its `+4`, `+0x70` use
and takes the command-line identity from `SESS-HERO-013`. Whether a path outside `R0909`
dispatches the pre-create's `vt+0x80`: no census of `+0x80` virtual calls was run. The `0x425` arm
posts `0x42f` instead of opening when `main+0x6bc` is 3 (`L11319`); `R1382` stores 2 there
(`L09316`), and the three calls between (`L08431`, `L11320`, `L09569`) were not read. What
`main+0x430` holds when the `+0x35c` route opens the screen was not traced.

### TEXT-074

The pre-create's left-button slot `vt+0x54` is `R1878` (`TOWN-211` maps message `0x201` to
`+0x54`). It calls `R1879(x, y, 1)` at `L11321`, which reads the region code under the
point from `Mask.bmp` (`R1880`) and, for the four hero codes, lights that element and stores
its slot in `pre+0x1cc`: `0x50` → 0 (`L09349`), `0x64` → 1 (`L09350`), `0x78` → 2 (`L09351`),
`0x8c` → 3 (`L09352`). The click routine then switches on the same code (`L11322`). Each hero
arm tests the last pressed slot `pre+0x1d8`, reads the field text (`R1881`) and writes an
`npcnames.txt` entry through `R1876` only when the text equals entry 20, 21, 22 or 23 or
`Unnamed`:

- `0x50`, slot 0, male fighter: skipped when `pre+0x1d8` is 0 (`L11323`); writes entry 20
  (`L11324..L11325`); stores 0 (`L11326`).
- `0x64`, slot 1, female fighter: skipped when it is 1 (`L11327`); writes entry 21
  (`L11328..L11329`); stores 1 (`L11330`, `L11331`).
- `0x78`, slot 2, female mage: skipped when it is 2 (`L11332`); writes entry 23
  (`L11333..L11334`); stores 2 (`L11335`).
- `0x8c`, slot 3, male mage: skipped when it is 3 (`L11336`); writes entry 22
  (`L11337..L11338`); stores 3 (`L11339`).

The store runs whether or not the text was written. `EnumRefs disp:1d8` finds 33 hits in 20
owners; the hits in pre-create routines are these arms and the enter, which stores the entered
slot 0 there (`L11301`), so `pre+0x1d8` holds the last pressed slot. The other 18 owners were not
traced to a pre-create pointer. The field text has five writers (`EnumRefs callto:R1876`: 5 hits,
2 owners, 0 orphan): the enter and these four arms. A typed name survives
every press. The slot meanings are `SESS-HERO-013`'s. From the first opening's text the four
presses leave: slot 0 EN `Danath`, RU `Данас` (no write); slot 1 `Naira`, `Найра`; slot 2
`Reniesta`, `Рениеста`; slot 3 `Fergard`, `Фергард`. The probe reads each entry's bytes on both
roots.

The pointer-move slot `vt+0x4c`, `R1882`, calls `R1879(x, y, flags & 1)` at
`L11340`. A move with bit 0 set over a portrait therefore stores its slot in `pre+0x1cc` without
entering a name arm and without updating `pre+0x1d8`.

**Confidence.** **High** for the four arms, their tests, the written entries and the writer census:
every branch is a cited instruction and the switch table is decoded from the image. **Medium** for
the pointer-move clause: it rests on `TOWN-211`'s `0x200` → `+0x4c` slot and on reading bit 0 of
the move's flags as the Win32 left-button bit, and no drag was observed.

**Unknown.** Whether any of the 18 `disp:1d8` owners outside the pre-create routines writes
`pre+0x1d8` through a pre-create pointer.

### TEXT-075

The name field is the class whose 30-slot vtable is `L07729` (`EnumRefs vt:L07729:30`: no slot
without a function). Three instructions store that address (`imm:L07729`, `refto:L07729`: 3 hits,
3 owners, 0 orphan): the constructors `R1883` (`L11341`) and `R1481` (`L11342`)
and the destructor `R1884` (`L11343`). `R1883` has no caller. `R1481` has
one, `L07730` in the pre-create's create routine, which builds the field with id `0x464` and rect
(224,310)-(362,337) (`L11344..L11272`) and stores it at `pre+0x1bc` (`L11273`). The text is
the MFC `CString` at `+0x60`.

The class's own slots, each read whole:

- `+0x74`, `R1480` (`TOWN-211`: message `0x102`): returns 0 when the length at `[+0x60]-8`
  is 10 or more (`L11277`); otherwise converts the byte with `R1517` (`TEXT-IN-005`) and
  calls `R1872`, which drops a byte below `0x20` (`L11345`) and appends any other with
  `CString::operator+=` (`L11346`).
- `+0x6c`, `R1874` (message `0x100`): on key 8 calls `R1873`, which replaces the text
  with its `Left(length - 1)` (`L11347..L11348`); every other key returns 0.
- `+0x54`, `R1472`: when the point is inside the field, calls `R1885` on the owner
  with the field (`L11349`), which gives the field the owner's `+0x38` focus; the text is
  untouched.
- `+0x4c`, `R1886`: sets the draw table `+0x64` to `+0x68` when the point is inside the
  field and to `+0x6c` when it is not; `R1875` stores the same table in both.
- `+0x2c` paints (`TEXT-076`); `+0x14` returns tooltip slot 256 (`TEXT-NAMEIN-024`).

The inherited slots `+0x30`, `+0x50`, `+0x58..+0x68` and `+0x70` return without a store
(`R1887`, `R1888`, `R1812..R1889`, `R1890`, `R1891`).
The only other writer found is the setter `R1876` (`TEXT-074`). No routine selects
text, moves an insertion point or replaces the text on a keystroke, so the first accepted keystroke
extends the seeded name. From each seed, typing stops at 10 bytes: EN `Danath` takes 4 more bytes,
RU `Данас` 5; `Naira` and `Найра` 5; `Fergard` and `Фергард` 3; `Reniesta` and `Рениеста` 2. The
cap limits typing only: the setter and the `-name` seed (`TEXT-073`) are not capped.

**Confidence.** **High** for the class identity, the construction census and the edit model: every
class slot is read end to end and every reference census names its instrument with 0 orphan.

**Unknown.** Whether a keystroke reaches the field before the field is pressed. By `TOWN-211` and
`TOWN-139`, a keyboard message goes to the owner's focused child `+0x38` and, unconsumed, to every
child's `vt+0x48`, which for this class ends in the slots above; this experiment did not re-read
that walk, what sets the owner's `+0x38` at opening, or how the frame delivers keyboard messages to
the pre-create. No `disp:` sweep of `+0x60` ran, so a write through a field pointer outside the
class is not excluded.

### TEXT-076

`R1892` is the field's paint slot `+0x2c` (`callto:R1892`: the vtable entry `L11350`
only). It reads `timeGetTime` (`L11351`, import slot `L00849`) and, when the caret flag
`+0x70` is set, draws a copy of the text with `|` (`0x7c`) appended by MFC's
`operator+(CString, char)` (`L11352..L11353`); otherwise it draws the text alone
(`L11354..L11355`). The draw is one font4 `vt+0x14` call with the field's table `+0x64`
(`L11356..L11357`, `TEXT-077`), so the caret is a glyph of the drawn string: it starts at the
text's advance sum and takes the text's colour. Font4 record 92 (`0x7c - 0x20`) inks columns 0..2
of rows 0..14, 45 pixels with 13 at level 15, and advances 4. The advance sums of the seeds are EN
`Danath` 54, `Naira` 42, `Fergard` 59, `Reniesta` 63 and RU `Данас` 47, `Найра` 50, `Фергард` 66,
`Рениеста` 72 pixels (`font4.dat` plus spacing 2, `SPR16A-FONT-018`).

After the draw, when `now - [+0x74]` exceeds 500 (`L11358..L11359`, a compare with 0x1f4 then
an exit on below-or-equal), the paint stores `now` in `+0x74` and flips `+0x70` (`L11360..L11361`). The pair is set
to (1, now) at construction by `R1875` (`L11362`, `L11363`; `callto:R1875`: the two
constructors) and by `R1872` before its `0x20` test (`L11364`, `L11365`;
`callto:R1872`: `L11279` only). Every character message under the cap, a dropped control byte
included, therefore shows the caret and restarts its phase; a character message at the cap, the key
8 handler and the setter do not. The paint reads no focus state: its fields are the owner's
`+0x08` and `+0x0c` and the field's `+0x08`, `+0x14`, `+0x60`, `+0x64`, `+0x70` and `+0x74`, so
the caret blinks whether or not the field holds focus.

**Confidence.** **High** for the caret glyph, its position and colour, the 500 ms rule and the
reset sites: the paint and both reset routines are read whole, and the census of each reset routine
is complete.

**Unknown.** How often the pre-create repaints. The flip happens only inside a paint, so each
visible phase lasts 500 ms plus the wait for the next paint.

### TEXT-077

The pre-create paint `R1474` (`vt+0x2c`) draws the prompt at `L11366..L11275`: string
`[[L04369]+0x1f4]`, global slot 125 (`TEXT-CHARGEN-028`), at x = field `+0x08` + screen `+0x08`
and y = field `+0x0c` + screen `+0x0c` - 5, flags 0, table `[[L11367]+8]`, through
`[L06186]`'s `vt+0x14`. `[L06186]` is font4 (`TEXT-API-007`), and its vtable `L11368`
holds `R1854` at `+0x14`. The field paint (`TEXT-076`) draws the name through the same
routine at x = field `+0x08` + screen `+0x08` and y = field `+0x14` + screen `+0x0c` - 16, where 16
is what the glyph sheet's `vt+0x24(0)` returns for font4's 16x16 record 0, with flags 0 and table
`+0x64` = `[[L11369]+8]` (`R1875`). With the field rect (224,310)-(362,337) (`TEXT-075`)
the prompt's cells start at screen origin + (224,305) and the name's at + (224,321).
`R1854` moves x only for flag values 1 (right) and 2 (centre) and y only for 4 and 8, so
flags 0 is left and top alignment and no text width enters either position. The screen origin is
`[L09808]`, `[L09809]` (`L11370..L07694`, the pre-create's construction), which is
`((W-640)/2, (H-480)/2)` (`SHOP-VIEW-044`): (0,0) at 640x480 and (192,144) at 1024x768, the two
sizes `L11371..L11372` selects.

Every inked literal word of `font4.16a` indexes palette entry 255 except four words at 254 in
record 6 (11,261 words and 4; 193 of 224 records carry ink), and the file is byte-identical on both
roots. The table pointer passed as the fourth argument selects the source table (`SPR16A-078`).
`R0572` fills two 256-entry palettes and builds each ramp with
`R1107(palette, 16, 4, 0)`: `L11373` bytes 0, 1, 2 = (20, 47, 65)·i/255 (`L11374`,
`L11375`, `L11376`) into `[L11367]` (`L11377`), and `L11378` bytes 0, 1, 2 =
(61, 39, 101)·i/255 (`L11379`, `L11380`, `L11381`) into `[L11369]` (`L11382`). Mode 4
(`L10380`) scales every entry by level/16 for levels 1..16, or by level/18 when `[L03346]` is
non-zero (`PAL-MODE4-010`), and packs the three bytes in the order `SPR16A-PAL-008` reads as
`[B,G,R,0]`. A level-15 word reads the table row built with level 16 (`SPR16A-ALPHA-025`). With the
`/16` table (`[L03346]` = 0) that row is `pal · 16/16` and the normal-memory blit adds no
destination term, so the prompt's opaque ink is RGB(65,47,20) and the name's RGB(101,39,61) before
packing; lower levels blend these with the background. `[L03346]`'s only writer sets it when
`GlobalMemoryStatus` reports less than 24,000,000 bytes of physical memory (`PAL-MODE4-010`). A
level-15 word then reads RGB(57,41,17) and RGB(89,34,54) from the `/18` table before packing, and
that branch's blit reads destination row `L` rather than `1+L` (`SPR16A-080`), so it also adds row
15, a sixteenth of the quantized background (`SPR16A-ALPHA-025`). The strings differ by root, EN
`Character name:` with 127 pixels of advance and RU `Имя персонажа:` with 132; placement, font and
colours do not, because the executable and both font4 files are byte-identical across the roots.

**Confidence.** **High** for placement, alignment, font, the two tables, the palette values and
both ramps' level-15 entries: each is a cited instruction or a decoded byte, and the channel order
follows `SPR16A-PAL-008` through the mode-4 arm that also serves file palettes.

**Unknown.** The pixel a display shows: packing uses the framebuffer's channel widths, which are
runtime state. Whether a native original runs with `[L03346]` set, which `PAL-MODE4-010` leaves
Unknown.

## Save-label chooser

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-SAVELABEL-054 | The established byte-indexed code-page converter `R0793` is not called by any of the nine SAVE/LOAD dialog functions, directly or through its only wrapper. | Medium | ● active | [EXP-0372](../experiments/EXP-0372-save-label-encoding/EXP-0372.md) |
| TEXT-SAVELABEL-055 | Withdrawn: the capped name class was said to be constructed nowhere by a literal vtable store; its vtable is `L07729`, three instructions store it, and `TEXT-075` finds its one construction at `L07730`. | — | ✖ retracted | [EXP-0372](../experiments/EXP-0372-save-label-encoding/EXP-0372.md), retracted by [EXP-0406](../experiments/EXP-0406-precreate-name/EXP-0406.md) |
| TEXT-SAVELABEL-057 | The SAVE label's copy primitive `R1604`/`R1893` (`SAV-LABELTAIL-236`'s NUL-bounded copy into application storage) filters no byte value other than `0x00` within the disassembled range read. | High | ● active | [EXP-0372](../experiments/EXP-0372-save-label-encoding/EXP-0372.md) |
| TEXT-SAVELABEL-058 | Neither catalogued text-rendering mechanism is reached by any of the nine SAVE/LOAD dialog functions, extending `TEXT-SAVELABEL-054` to the draw functions, the glyph blit and the GDI text-out imports. | High / Medium | ● active | [EXP-0376](../experiments/EXP-0376-save-label-draw/EXP-0376.md) |
| TEXT-SAVELABEL-059 | The four MFC `CDC` GDI-wrapper stubs sit at matching offsets in the `CDC`, `CClientDC`, `CWindowDC` and `CPaintDC` tables; `imm:`/`disp:`/`refto:` find no reference to `L11383`/`L11384`/`L11385`/`L11386`, slot 11 of each. | High / Unknown | ● active (partially retracted) | [EXP-0376](../experiments/EXP-0376-save-label-draw/EXP-0376.md), partially retracted by [EXP-0406](../experiments/EXP-0406-precreate-name/EXP-0406.md) |
| TEXT-SAVELABEL-060 | `R0375`, the child-vector-by-id accessor `TOWN-354` reads, sits immediately before the copy `SAV-SAVELABEL-1017` traced into `dialog+0x68`, and it is a system-wide utility, not SAVE/LOAD-specific. | High / Medium | ● active | [EXP-0376](../experiments/EXP-0376-save-label-draw/EXP-0376.md) |
| TEXT-SAVELABEL-061 | By elimination among the two catalogued text-rendering mechanisms and their sole GDI-wrapper path, the SAVE/LOAD chooser's label pixels are consistent with native list-control painting outside `rom.exe`'s code. | Medium / Unknown | ● active (amended) | [EXP-0376](../experiments/EXP-0376-save-label-draw/EXP-0376.md), amended [EXP-0406](../experiments/EXP-0406-precreate-name/EXP-0406.md) |

### TEXT-SAVELABEL-054

`callto:R0793` finds 3 direct callers; through its only wrapper `R0766`
(`callto:R0766`), 14 more callers, 15 distinct owning functions total (`R0767` and
`R1854` call both the converter directly and the wrapper, so they are counted once, not
twice). None of the nine SAVE/LOAD dialog functions this experiment and `SAV-LABELTAIL-236` traced —
`R1246`, `R1220`, `R0737`, `R1245`, `R1221`, `R0738`, `R0701`, `R1284`,
`R0084` — is among the 15. `callto:` also scans every `.rdata` dword for a literal match to
either function's address; none is found, so vtable-dispatched entry to either is excluded

**Confidence.** **Medium.** A bounded, reproducible, two-levels-deep reverse-direct-call search
(direct callers of the converter, plus callers of its one known wrapper) over a named function
population, with vtable-dispatched entry ruled out by the same scan. It does not establish what
mechanism, if any, draws the label instead — no paint/`WM_DRAWITEM` call chain was traced, and a
caller reached only through a third level, or an indirect call through a pointer not stored in
`.rdata`, would not appear

### TEXT-SAVELABEL-055

- What stands is the search. No disassembled instruction references `L07728` in the
  1,977,344-byte EN `rom.exe` (`imm`, `disp`, `refto` and `text` modes; EXP-0406 repeats `imm` and
  `refto` with 0 hits), and the same modes find the SAVE dialog's vtable `L06376` at
  `L11387`.
- What fell is the address. `L07728` is slot 6 of the class's 30-slot vtable, which starts at
  `L07729`. That address is stored by the literal-vtable-store idiom in both constructors
  (`L11341`, `L11342`) and in the destructor (`L11343`), and the constructor `R1481`
  is called once, at `L07730`, where the pre-create screen builds its name field (`TEXT-075`).
- The class is not the save-label producer for that reason instead: its one direct construction
  belongs to the pre-create screen (`callto:R1481` 1 hit, `callto:R1883` 0 hits).

**Confidence.** The High was earned for the absence of `L07728` and stated about the class; it
stands only for that address. See [`retracted.md`](retracted.md).

**Amended.** Retracted as a whole by EXP-0406 (`TEXT-075`, [`retracted.md`](retracted.md)). The
withdrawn headline read: "`TEXT-NAMEIN-024`'s capped text-entry class is not constructed anywhere
in the image by the literal-vtable-store idiom every other read constructor in this codebase uses."
Its Medium clause, that a class built by another means would escape the search, was not the gap;
the searched address was.

### TEXT-SAVELABEL-057

The SAVE label's own copy primitive, `R1604`/`R1893` (the function `SAV-LABELTAIL-236`
already named as the NUL-bounded copy into application storage), filters no byte value other than
`0x00` within the disassembled range read.

The byte-wise fallback loop has exactly one value comparison, the zero test of the loaded byte with a zero branch, at
`L11388..L11389`, testing only for the terminator; every byte `0x01..0xFF` it copies is stored
unchanged. Its aligned-dword path uses a separate mechanism, the CRT's standard
`0x7efefeff`/`0x81010100` zero-byte detector plus four byte selectors
(a test of each of the four bytes of the loaded dword: the low byte, the next byte, and the masks 0xff0000 and 0xff000000) — five value-bearing tests,
not one, each locating the terminating `0x00` rather than filtering byte content; this detector's
full shape is visible within the committed evidence in the adjacent routine at `L11390..L11391`,
syntactically identical CRT machinery not itself on this call's executed path. The committed
disassembly range (`R1604..L11392`) stops before the primitive's own executed instance of this
detector (beginning `L11393`) completes; the function's own tail through its `RET` was not read
here. This is the opposite of `TEXT-NAMEIN-024`'s capped field, which drops every byte below `0x20`
and caps stored length at ten — the SAVE label path has neither rule in the range read. A live
corpus witness is `SAV-SAVELABEL-1018`

**Confidence.** **High** for the instruction read within the captured range: every test located is a
terminator/zero-byte-position test, none a value filter, and the byte-wise loop's single
zero test of the loaded byte is the only comparison gating what gets stored. The primitive's own tail beyond
`L11392` was not read

### TEXT-SAVELABEL-058

Neither of this codebase's two catalogued text-rendering mechanisms is reached by any of the nine
SAVE/LOAD dialog functions, extending `TEXT-SAVELABEL-054`'s converter-only negative to the draw
functions, the glyph blit, and the whole GDI text-out import surface.

`callto:R0767`/`callto:R1854`/`callto:R1250` (the font-atlas draw functions and glyph blit
already established for other text surfaces) find, in the whole 1,977,344-byte image, exactly one
reference each — the function's own vtable slot on its owning font object's table
(`L11394`/`L11395`/`L11396`), never a literal call site; the same `.rdata` scan that catches
vtable-slot stores finds no second reference. `refto:` (validated against a known-good control
before use: `callto:L11397` finds 0 hits for `GetDC`, whose caller `R1894` is already named
in `EXP-0355`'s own committed evidence, showing `callto:` alone would silently miss every import
call — `refto:` correctly finds it) for the six GDI-text imports this image actually links
(`ExtTextOutA`, `TextOutA`, `TabbedTextOutA`, `DrawTextA`, `GetTextExtentPoint32A`,
`GetTextExtentPointA`) finds callers in seven functions system-wide — four MFC `CDC` wrapper stubs
(`R1895`/`R1896`/`R1897`/`R1898`), `CDC::FillSolidRect` (`L11398`, itself implemented
with `ExtTextOutA`, an MFC idiom not this codebase's own), and two further unidentified functions —
`R1899` (reached only from `R1900`, four literal sites) and `R1901` (the sole
caller of `GetTextExtentPointA`; its own callers were not searched by this experiment) — none of the
nine SAVE/LOAD dialog functions among them, and neither `R1900`'s own callers nor
`R1901`'s were searched, so an unsearched path from the nine reaching either through a
further level is unfound, not excluded. Within the nine functions, the only references to any of the
four known font-object globals (`font1`..`font4`, `L03615`/`L02677`/`L10053`/`L06186`) are
eighteen occurrences (six in `R1220`, seven in `R1221`, five in `R0701`) of one
repeated idiom — load the font1 pointer, then a compile-time-constant caption selector (a pushed
index `0x97`/`0x1d`/`0x90` into the caption-lookup accessor `R0668`, or, for `R0701`,
a fixed `+0x34c` displacement off a second string-table global) into a control-construction helper —
a fixed-caption/warning control, not a per-item draw call; none reads a loop variable, list index,
or the label buffer. The remaining two of the nine dialog functions, `R1246` and
`R1245`, carry no reference to any of the four font-object globals at all
(`evidence/disasm-ctors.txt`)

**Confidence.** **High** for the search itself (six modes, reproducible, whole-image scope).
**Medium** for "this rules out the two catalogued mechanisms": `callto:`/`refto:` cannot see a
virtual call dispatched through a font-object pointer without a literal reference to the callee's
own address (indirect calls through offset +0x14 were not enumerated system-wide), so an
uncatalogued caller reached only that way would not appear; the font-global check is complete for
the four *known* font objects only

### TEXT-SAVELABEL-059

The four MFC `CDC` GDI-wrapper stubs this image links (`R1895`/`R1896`/`R1897`/`R1898`,
wrapping `TextOutA`/`ExtTextOutA`/`TabbedTextOutA`/`DrawTextA`) sit at matching offsets across four
near-identical vtables (`L11383`/`L11384`/`L11385`/`L11386`), identical over the 20 slots
read (`vt:` dumped 20 slots per base, spacing `0x80` between bases; 20 is what was read, not a
measured total table length)~~, and none of those four vtables is constructed anywhere in the
image — extending `TEXT-SAVELABEL-055`'s vtable-never-constructed pattern from the character-name
class to this GDI-wrapper family~~.

`callto:` for three of the four wrappers (`R1896`/`R1897`/`R1898`) finds exactly four
`RDATA-SLOT` entries each (one per vtable, at matching offsets) and no literal caller; the fourth,
`R1895` (the `TextOutA` wrapper), was not reverse-searched by `callto:` in this experiment, but
the same four `vt:` dumps place it at slot `+0x38` in all four vtables. `imm:`/`disp:`/`refto:` for
all four ~~vtable base~~ addresses find 0 hits each across the whole image, the same three-mode
idiom `TEXT-SAVELABEL-055` used for its own control case. `refto:L11398` (`CDC::FillSolidRect`, the one
function in this family with a real body) also finds 0 callers anywhere

**Confidence.** **High** for the search (three modes times four addresses, one reproducible
process, zero hits; the `vt:` slot pattern shown for all four wrappers, and `callto:` literal-caller
absence measured for three of the four; `R1895`'s own literal-caller count was not separately
measured). ~~**Medium** that this makes the family unreachable: the same caveat
`TEXT-SAVELABEL-055` already named applies unchanged — an immediate in undisassembled bytes, or a
class constructed some other way, would not appear~~

**Unknown.** Whether a constructed `CDC`, `CWindowDC` or `CPaintDC` object reaches one of the four
text wrappers on a save-label path.

**Amended.** EXP-0406 ([`retracted.md`](retracted.md), PARTIALLY RETRACTED) withdraws the
never-constructed clause and its Medium. The four searched addresses are slot 11 (`+0x2c`) of
tables that start `0x2c` lower, at `L07723`, `L07724`, `L07725` and `L07726`
(`vt:` 31 slots each). The dword before each start is an RTTI locator whose type descriptor names
`.?AVCDC@@`, `.?AVCClientDC@@`, `.?AVCWindowDC@@` and `.?AVCPaintDC@@`, and the four wrappers are
slots 25..28 (`+0x64..+0x70`) of every table. Each start is stored by its constructor (`L11399`,
`L11400`, `L11401`, `L11402`) and its destructor (`L11403`, `L11404`, `L11405`,
`L11406`). `callto:` finds six calls to the `CDC` constructor `L11407`, three of them from the
other three constructors; three to the `CWindowDC` constructor `L11408`; two to the `CPaintDC`
constructor `L11409`; and none to the `CClientDC` constructor `L11410`. The `CPaintDC` call at
`L07727` is in `R1479`. `callto:R1479` finds only the `.rdata` dword `L11411`, the
last of six dwords at `L11412` that read `0x0f`, 0, 0, 0, `0x0c`, `R1479`: the layout of
an MFC message-map entry binding message `0x0f` (`WM_PAINT`) to that routine.
`TEXT-SAVELABEL-055`, the pattern this card extended, is retracted. The zero-hit search stands for
the four interior addresses only.

### TEXT-SAVELABEL-060

`R0375` — the child-vector-by-id accessor `TOWN-354` already reads (walks the child vector at
`this+0x1c`, returns the child whose `+0x4` equals the passed id) — sits immediately before the copy
`SAV-SAVELABEL-1017` already traced into `dialog+0x68` at the SAVE dispatcher's rename-commit
handler. This experiment adds that the same accessor is a system-wide utility, not
SAVE/LOAD-dialog-specific.

`callto:R0375` finds 189 call sites across 54 distinct owning functions image-wide,
`R0738` among them (three sites: `L11413`/`L11414`/`L11415`, using tags `4`/`1`/`3`);
at the `L11414` site the returned child's own vtable slot `+0x3c` is then called indirectly
(an indirect call through offset +0x3c) before the already-published copy into `dialog+0x68` at `L08783`
(that copy's fixed-capacity source is `SAV-LABELTAIL-236`'s 256-byte stack buffer, not re-measured
here). The returned child's concrete class/vtable identity — narrower than `TOWN-354`'s own scope,
which already places it in the dialog's own child vector — and the `+0x3c` method's own body were
not located in this experiment

**Confidence.** **High** for the accessor's body, established by `TOWN-354`. **Medium** for the
genericness count measured here (a named, reproducible, whole-image `callto:` search) and for "this
accessor feeds the already-published copy": the call ordering and tag values are a direct
instruction read, but the returned child's concrete class and the `+0x3c` method's own body were not
traced, so whether it reads back already-rendered widget state or something else remains open

### TEXT-SAVELABEL-061

By elimination among this codebase's two catalogued text-rendering mechanisms (`TEXT-SAVELABEL-058`)
and their sole GDI-wrapper access path (`TEXT-SAVELABEL-059`), the SAVE/LOAD chooser's per-item
label pixels are consistent with native Win32/MFC list-control default painting occurring entirely
outside `rom.exe`'s own code — an inference by elimination among named candidates, not a positive
trace of an external paint call, and it does not exclude an uncatalogued in-game draw routine
reached only by virtual dispatch this search cannot enumerate.

Because no draw primitive inside this image is shown reached, no byte-to-glyph table is shown
indexed for this field either, so `TEXT-SAVELABEL-054`'s open item narrows rather than closes: EN
and RU cannot differ in a mechanism this image does not exhibit for this field, because the two
lawful executables are the same file (independently re-verified in this experiment, not only cited
from `EXP-0372`) — any difference would have to come from outside `rom.exe`, unreachable by static
analysis. What the draw path does with a byte outside 7-bit printable ASCII, and what bounds a
label's drawn (as opposed to retrieved) length in a chooser row, are Unreachable by this experiment:
both require observing the native control's own paint step, which needs a running original

**Confidence.** **Medium** for the elimination and for "EN/RU cannot differ here" (both rest on
`TEXT-SAVELABEL-058`/`-059`'s own Medium bounds). **Unreachable**, not merely Unknown, for the
non-ASCII draw handling and the visible length bound: both need a running original, excluded by the
static-analysis-only boundary

**Amended.** EXP-0406 partially retracts `TEXT-SAVELABEL-059`: the `CDC`, `CWindowDC` and
`CPaintDC` classes whose tables hold the GDI text wrappers are constructed in the image, so the
wrappers are not unreachable by construction. The GDI-wrapper leg of this elimination rests on
`TEXT-SAVELABEL-058`'s traced-path negative alone: no traced SAVE/LOAD chooser path reaches the
wrappers or the text-out imports.

## Help text

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-086 | `main\text\help.txt` is read once at startup as one NUL-terminated string into the global string object at `L06232`, outside the sixteen-file line table, with no code-page pass at load. | High | ● active | [EXP-0437](../experiments/EXP-0437-f1-help/) |
| TEXT-087 | Both roots' `help.txt` are CRLF-terminated paragraph lines with no other control byte and no NUL; EN is 965 bytes of 7-bit text in 35 pieces, RU 1313 bytes with 837 high bytes of 34 values in 39 pieces. | High | ● active | [EXP-0437](../experiments/EXP-0437-f1-help/) |
| TEXT-088 | Help is wrapped at 408 px by the dialogue splitter and wrapper, overflows the 12-line body on both roots, and is rewrapped at 382 px beside a scroll bar; it is not paged. | High / Medium | ● active (amended) | [EXP-0437](../experiments/EXP-0437-f1-help/) |
| TEXT-089 | Each root's `help.txt` holds exactly one doubled tilde and no lone tilde, so help draws one literal `~` glyph and never reaches the underline arm; the OK label holds no tilde. | High | ● active | [EXP-0437](../experiments/EXP-0437-f1-help/) |

### TEXT-086

- The startup routine `R0326` pushes `L06232` and the literal at `L11416` (`main\text\help.txt`, between the `main\text\heropicture.txt` and `main\text\main.txt` literals) at `L11417`/`L06250` and calls `R1293` (`L11418`). The sixteen-file table loads that follow (the table receiver, the path push and the call to `R0661`, `TEXT-STRTAB-023`) do not include it.
- `R1293` opens the node (`R0535`), takes its size (`R0721`), allocates size plus 1 (`R0722`), reads the whole payload (`R0533`), stores a NUL at `buffer[size]` and copies the buffer into the string object (`R1774`, `R0525`). Nothing splits the text into lines, and no `R0793` call lies in the routine.
- `L06232` has four references image-wide (`EnumRefs refto:L06232`, 4 hits, 4 owners, 1 in orphan code): the load push, two stubs that load the receiver `L06232` at `L11419` and `L11420` (the object's destructor call), and the single read at `L06231`, the help case of `MENU-051`. Help is the only consumer.
- The converter acts at draw time (`TEXT-CONV-001`), so the stored bytes are the file's bytes.
- Each root's `main.res` holds exactly one entry named `text/help.txt` (`tools/helptxt`).

**Confidence.** High: the load routine is read whole, the reference enumeration names its instrument and returns one reader, and the file population is a listing of both roots' archives.

### TEXT-087

- Instrument: `tools/helptxt` over each root's `main.res`, which prints counts, never text. Both files end with `CRLF`; bare CR 0, bare LF 0, NUL 0, other control bytes 0.
- EN: 965 bytes, 34 `CRLF`, 35 pieces on splitting at `CRLF`, 131 space bytes, no byte at or above `0x80`. The two empty pieces are the last two (a blank line before the final `CRLF`). Longest piece 60 bytes.
- RU: 1313 bytes, 38 `CRLF`, 39 pieces, 157 space bytes, 837 bytes at or above `0x80` taking 34 distinct values, and the first byte is above `0x80`. The two empty pieces are piece 1 (a blank line after the first line) and the last (the file ends after one `CRLF`). Longest piece 61 bytes.
- The files therefore share terminator, container and markup, and differ in language bytes, blank-line placement (EN trailing, RU after the first line) and line counts. All 837 RU high bytes lie in the converter's source set for the Russian selector, `0x80..0xAF` and `0xE0..0xEF` (`TEXT-DOM-010`, `TEXT-FIT2-013`): 0 occurrences fall outside it.
- `help.txt` carries no section, page or escape marker other than the doubled tilde of `TEXT-089`.

**Confidence.** High for the census: each figure is a direct count over the archive entry of each root, with the output committed. The meaning of a byte (which glyph a high byte draws on the Russian selector) is `TEXT-CONV-001`'s, not re-derived.

### TEXT-088

- The body control of `MENU-052` gives the text to `R0726`: the splitter `R0792` cuts at `CRLF` and trims left, the wrapper `R0727` fits words to the width with the measure `R0766` over font 1's `.dat` advances plus spacing 2 (`DLG-LINE-038`, `TEXT-079`).
- Instrument: a transcription of that splitter, wrapper and measure in `tools/helptxt`, over each root's `font1.dat` and `help.txt`, no emulation. Wrap width 408 (body rectangle width). EN: 33 non-empty pieces wrap to 35 lines; RU: 37 pieces wrap to 47 lines. No wrapper stall on either.
- `R1184` compares the line count times the line height with the body height 211 (`MENU-052`): 35 lines against 211 EN, 47 against 211 RU, a scroll bar on both. The width then drops by `0x1a` to 382 and the text is rewrapped: EN 35 lines (widest 376 px), RU 48 lines (widest 365 px). The comparison does not change with the line height read (15 or 17): 35 lines times 15 is 525.
- Arithmetic only: with 12 visible lines, lines minus visible would be 23 (EN) and 36 (RU). No routine read sets the scroll bar's range, so that formula is an assumption.
- Each paragraph's first line is indented 10 px and every line except a paragraph's last and the final line is justified (`DLG-LINE-038`), with a 1 px shadow. No key or marker pages the text.

**Confidence.** High for the pipeline and the scroll decision (routines read whole, the decision holds for any line height read). Medium for the exact line counts and the 12-line window: the wrapper is transcribed from its listing and checked against the dialogue rules, not emulated, and the body height is derived arithmetic (`MENU-052`).

**Unknown.** Native floating-point justification (`DLG-LINE-038`'s Unknown applies). The scroll bar's range and initial position are read in `MENU-079`.

**Amended.** The scroll bar's range is read: `R0730` passes lines minus visible plus 1 to the bar (`MENU-079`), so the top position runs 0..23 for EN and 0..36 for RU, starting at 0. The 23 and 36 above are no longer an unread formula.

### TEXT-089

- Instrument: `tools/helptxt` byte census. EN `help.txt` holds two `~` bytes at offsets 103 and 104; RU holds two at offsets 95 and 96. Each pair is adjacent, so the file holds one doubled tilde and no lone tilde.
- By `TEXT-079` a doubled tilde draws one literal `~` glyph and the loop skips the partner, and the underline arm runs only for a lone tilde. Help never reaches the arm. The measurer counts a doubled tilde once, so the wrap widths of `TEXT-088` treat the pair as one glyph.
- The first line of `dialogs.txt`, the OK label, holds no tilde in either root (length 2 EN, 7 RU).
- This extends the tilde census of `TEXT-079`, which covered the five dialogue families and one button label: it adds `help.txt` and the OK label for both roots.

**Confidence.** High: a direct byte count over the file of each root. The glyph drawn follows `TEXT-079`'s High clause.

## Spellbook item card

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-090 | For an item whose stream carries kind 42, the formatter appends a spell fragment to the name line, ` of <spell>` EN and ` с заклинанием <spell>` RU, and adds no prefix to the buffer it composes. | High / Medium | ● active | [EXP-0438](../experiments/EXP-0438-resolver-regen-card/) |
| TEXT-091 | The spell name in that fragment is line `spell id - 1` of `main\text\spell.txt` (28 lines on both roots), inserted unchanged; EN and RU differ only in text-file content, since `rom.exe` is identical and holds the one format literal. | High / Medium | ● active | [EXP-0438](../experiments/EXP-0438-resolver-regen-card/) |
| TEXT-092 | The same formatter gives a non-book item a new `#` line `casts <spell>` for effect kind 41 and a `#Magic:` header from marker `0x33`; a book (class `0xe00`) takes the ` of <spell>` form for kind 41 and no header. | High / Medium | ● active | [EXP-0438](../experiments/EXP-0438-resolver-regen-card/) |

### TEXT-090

- `R0964` sets the buffer at `L04995` to the empty string
  (`L11421`). It looks the item's code word `item+0x06` up in the name map
  `L04790` (`L11422`..`L11423`, `ITEM-DISPNAME-036`), copies the found
  name into the buffer (`L11424`..`L11425`), then loops over the attribute
  stream once per entry, the count being the byte at `item+0x09`
  (`L04992`..`L11426`, `L04924`..`L04994`). Each pass reads a tag byte,
  forms `tag - 1`, indexes the byte table `L04646` (limit `0x32`,
  `L11427`) and jumps through the dword table `L04647` (`L11428`).
- Tag 42 (`teachSpell`, `ITEM-EFFKEY-147`) is table entry 4, `L11429`. The
  arm reads one stream byte, the spell id; loads `main[90]` and `main[91]`
  from the shared string array `[L04369]` at `+0x168` and `+0x16c`; takes
  the spell name as line `id - 1` of the table object `L11145` through
  `R0668` (`L11430`..`L11431`); and calls the formatter
  `R1196` with the literal `L11432`, ` %s %s%s`
  (`L11433`, `L11434`). The arguments are `main[90]`, the spell name,
  `main[91]`. The result is appended to the end of the buffer
  (`L11435`..`L04996`). The literal has a leading space and no `#`, and the
  name line carries no text before the name from this arm, so the fragment
  continues the name line.
- Strings (`tools/spellcard`, both roots, 274 `main.txt` lines): EN `main[90]`
  is `of` and RU `main[90]` is `с заклинанием`; `main[91]` is empty on both. The
  five book codes `0xe13`..`0xe17` name `Book` in EN and `Книга` in RU. An
  item named Book that carries kind 42 therefore reads `Book of <spell>` in EN
  and `Книга с заклинанием <spell>` in RU.
- The formatter has four call sites: `L04927`, `L04925`, `L04926` and
  `L04928` (`EnumRefs callto:R0964`, 0 orphan hits). The hover painter
  splits the buffer at `#` (`TEXT-HOVERPAINT-053`). The other three painters
  were not read.

**Confidence.** High for the composition: the arm, its pushes, the literal and
the append are cited instructions read whole, and the name copy precedes the
loop. Medium for what a player sees, because only the hover painter's
treatment of `#` is established and that painter was not traced for this card.

**Unknown.** Whether a non-hover painter inserts a title or a break around the
buffer; the three other callers were not read. Which shipped item rows carry
kind 41 or kind 42: no evidence file lists the stream of any row, so which kind
a shipped book uses, and so whether a shipped book card shows this fragment, is
open.

### TEXT-091

- The spell id byte is one-based: the arm decrements it (`L11430`)
  before the table call. The table object `L11145` is
  `main\text\spell.txt` (`L11436`; `L11437` is `spells.txt`,
  `L03191` `stats.txt`, `L11438` `main.txt`; `TEXT-STRTAB-023`). The
  accessor `R0668(table, i)` returns that table's line `i`
  (`UNIT-NAME-039`).
- Both roots carry 28 `spell.txt` lines; row lengths are 4..21 characters EN
  and 4..23 RU, one line per spell and no per-case column. The formatter has
  no second index, so the RU row is shown in the form it is stored in, after
  `с заклинанием`.
- `rom.exe` has one SHA-256 on both roots (`inputs.tsv`), so the format
  literals `L04998`..`L04649` are identical; the only EN/RU difference
  in the card is the content of the text files.

**Confidence.** High for the row selection and for the absence of a case-form
choice. Medium for the grammatical form of any one RU row: the rows were
counted and hashed, not read for inflection, and the card shows whatever form
a row holds.

### TEXT-092

- Tag 41 (`castSpell`) is table entry 3, `L11439`. It reads the spell id and
  tests the item class `item+0x06 & 0xf00` against `0xe00`
  (`L11440`..`L11441`). For class `0xe00` it takes the ` of` form of
  `TEXT-090` (`main[90]`, `main[91]`, literal `L11432`;
  `L11442`..`L11443`). For any other class it loads `main[92]` and
  `main[93]` (`+0x170`, `+0x174`) and the literal `L11444`, `#%s %s%s`
  (`L11445`..`L11446`), which starts a new line with the verb word. EN
  `main[92]` is `casts` and RU `main[92]` is `с заклинанием`; `main[93]` is
  empty on both.
- The marker tag `0x33` (51) is table entry 7, `L11447`. For class `0xe00`
  it skips to the loop tail (the zero branch to `L04924` at `L11448`). For any other class it
  appends the literal `L11449`, `#`, then `main[189]` (`+0x2f4`,
  `L11450`): EN `Magic:`, RU `магия:`. The writer emits the marker first in
  every stream built by `R0998` (`ITEM-152`), so a non-book card
  carries `#Magic:` before its spell lines.
- The other literals the arms use are `L04998` `#%s %+d`, `L11451`
  `#%s: %d` and `L04649` `#%s %d`; they are not part of this question.

**Confidence.** High for the arms and literals, read whole. Medium for the
on-screen result, for the reason in `TEXT-090`.

**Unknown.** Which other effects a Book row carries, and so whether a shipped
book card has lines beyond the fragment.

## Notice and pause lines

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-094 | `main.txt[119]` is the Pause panel's body text: one load of slot 119 exists among the 30 hits for displacement or immediate `0x1dc`, in the Pause arm; it is a modal panel, not a message-line line. | High / Medium | ● active | [EXP-0446](../experiments/EXP-0446-toggle-notice/) |
| TEXT-095 | Slots 94..107 and 218..220 are read by the seven setting arms and slots 108..116 by the speed arms; two other routines have scaled-index loads at those displacements, not traced to their base. | Medium | ● active | [EXP-0446](../experiments/EXP-0446-toggle-notice/) |

### TEXT-094

- Population: `EnumRefs disp:1dc` (25 hits, 20 owners) and `imm:1dc` (5 hits, 5 owners) over the EN `rom.exe`, 0 hits in orphan or undisassembled code, and `refto:L04369` (230 hits, 47 owners, `TEXT-STRTAB-023`). A function can read slot 119 from the shared array only if it is one of those 47 owners; three of the 30 hit owners are (`R1210`, `R0231`, `R1902`).
- `R1210` stores the constant `0x22` into the field at `+0x1dc` (`L11452`), a field write. `R1902` takes the address of `+0x1dc` (`L11453`), the address of a field of its own object. `R0231` loads `[[L04369]+0x1dc]` at `L11454`, inside the Pause arm: it allocates `0x78` bytes, calls `R1179` with rectangle (0x20, 0x30, 0x260, 0x1b0) and shows the panel through `R0361`, the panel `MENU-052` describes. The arm requires `campaign+0x6bc == 2` and `campaign+0x3dc == 1` (`L11455`, `L11456`).
- No toggle or speed arm loads slot 119, and none draws on the Pause panel: the two use unrelated surfaces (`MENU-059`).

**Confidence.** High for the Pause load and for the surface (named instructions). Medium for "the only load": the owner rule excludes the other 27 hits but misses a pointer copied out of the array before use and a load through an index register with another displacement.

**Unknown.** Loads of slot 119 through a copy of the array pointer or a computed index.

### TEXT-095

- Indexed loads at the notice displacements, from `EnumRefs disp:` at 0x178..0x1d0 step 4 and 0x368..0x370, in owners of `refto:L04369`: `R0819` at `L11457`, `L11458`, `L11459`, `L11460`, `L08820`, `L11461` and `L11462` (the seven setting arms), and `R0231` at `L11463` and `L11464` (the speed arms).
- Four further `[base + index*4 + disp]` loads exist at those displacements: `L11465` in `R1207` (`0x1ac`) and `L11466`, `L11467`, `L11468` in `R1211` (`0x1d0`). Those routines are the spellbook hover getter and the generator attribute getter, whose array loads `TEXT-080` and `TEXT-082` read at other displacements (`main[182..187,217]`, `main[155+i]`). The register that holds the base of these four loads was not traced.
- The remaining owner hits at these displacements are stores or fixed-field accesses (for example in `R1474`, `R1902`, `R1210`); they were classified by operand shape, not each traced.

**Confidence.** Medium: the nine arm loads are named instructions; the exclusion of the other hits rests on the owner rule, operand shape and two earlier claims.

**Unknown.** The base register of the four indexed loads in `R1207` and `R1211`, and any array read with a computed index outside the swept displacements.


## Settings dialog strings

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-099 | Game Options reads 28 `dialogs.txt` rows and `patch.txt` rows 52 to 54, and Sound Options reads 20 `dialogs.txt` rows; every cited row exists on both roots. | High / Medium | ● active | [EXP-0457](../experiments/EXP-0457-settings/) |
| TEXT-100 | `tunes.txt` has 21 lowercase `name.wav` keys with titles, identical on both roots; `cutscene.txt` and `cutpaths.txt` have 14 rows each, and the `cutpaths.txt` rows are identical on both roots. | High / Medium | ● active | [EXP-0457](../experiments/EXP-0457-settings/) |

### TEXT-099

Game Options (MENU-073) reads these `dialogs.txt` local rows: 0, 1, 50 to 59, 60 to 67, 78, 150, 156, 159 to 162 and 164 (28 distinct rows), and `patch.txt` rows 52, 53, 54. Sound Options (MENU-075) reads rows 0, 7 to 21, 23, 75, 143 and 165 (20 distinct rows). Row 0 is the OK caption and row 1 the Cancel caption. The rows come from the text accessor `R0668` on the dialogs object `L06197` and the patch object `L06351`.

Both roots hold 166 `dialogs.txt` rows and 67 `patch.txt` rows, and every cited slot is present in both. Per-slot byte lengths differ by language (for example row 0 is 2 bytes on EN and 7 on RU), so the controls are sized by rectangle, not by text. Slot plus byte-length rows are in `textrows.txt` of the experiment evidence.

**Confidence.** **High** for the slot indices and the presence on both roots (the instruction pushes and the string tables of both roots were compared by index and length). **Medium** for the hint-versus-caption order of each pair, which follows the constructor-argument order of the control classes (MENU-068).
The EN and RU `rom.exe` are one image (SHA-256 `942e9b72…`), so their agreement
on this code read is not independent corroboration. The agreement between the per-root
data stands as corroboration: the string tables are separate files in each root.

**Unknown.** The text of the rows is not compared between languages beyond length; several hint and caption pairs have equal lengths on each root and may be the same string.

### TEXT-100

`main/text/tunes.txt` holds one `key=title` row per music track: 21 rows on each root, the keys lowercase and ending `.wav`, identical on EN and RU, every row with a nonempty title. The Sound Options list resolves each candidate name of the player's bank, without its six-character `music\` prefix, to a title through a dictionary built from this table (VIDEO-OPTIONS-057, VIDEO-076). `cutscene.txt` holds the 14 cutscene titles shown in the cutscene list and `cutpaths.txt` the 14 directory names of the movie parts; the latter is not localised (identical rows on both roots, lengths 3 to 7 bytes) and the former has 14 nonempty rows on both. Both are loaded at start (`L11469`, `L11470`) into objects `L11471` and `L11472` (VIDEO-075). The three tables sit in global string order after `npcnames.txt`, at global bases 1245, 1259 and 1273.

**Confidence.** **High** for the row counts, key identity and the load calls. **Medium** that the key lookup is case-folded, which rests on the identity of the lowercase routine and was not traced.

**Unknown.** Behaviour for a candidate with no dictionary row; whether the lookup compares case-sensitively.


## Typed byte to font record

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| TEXT-108 | On the Russian selector a typed `CP1251` byte reaches record `byte − 0x30` over `0xC0..0xEF` and `byte − 0x20` over `0xF0..0xFF`; the offset is not one constant, and font3's 64 records hold none of the 64 bytes. | High | ● active | [EXP-0488](../experiments/EXP-0488-chargen-atlas/) |
| TEXT-109 | The records the Russian rule reaches draw the typed letter in font1 and font2 (17 of 18 homoglyph pairs pixel-identical, against 4 of 18 at `byte − 0x20`); font4 and font5 follow the same arrangement on advance evidence only. | High / Medium | ● active | [EXP-0488](../experiments/EXP-0488-chargen-atlas/) |
| TEXT-110 | On the English selector a typed byte reaches record `byte − 0x20`; of the 64 `CP1251` bytes `0xC0..0xFF`, 16 draw their own letter, 32 another Cyrillic letter, 16 reach records 192..207 (blank except `ß`). | High / Medium | ● active | [EXP-0488](../experiments/EXP-0488-chargen-atlas/) |

### TEXT-108

- Input conversion `R1517` (`R1517..L11473`, one image on both roots) returns a byte below `0x80` unchanged. At selector 1 (`[L06203]`, `TEXT-LANG-002`) it adds `0xC0` to `0xC0..0xEF` (`−0x40`) and `0xF0` to `0xF0..0xFF` (`−0x10`). The display conversion `R0793` adds `0x30` to `0x80..0xAF` and `0x10` to `0xE0..0xEF` (`TEXT-CONV-001`). The record is the display result minus `0x20` (`L07943`).
- Composed for a typed byte `t` at selector 1: `0xC0..0xEF` is stored `t − 0x40`, drawn `t − 0x10`, record `t − 0x30` (144..191). `0xF0..0xFF` is stored `t − 0x10`, drawn `t`, record `t − 0x20` (208..223). `evidence/cp1251-map-ru.tsv` lists all 128 typed bytes `0x80..0xFF` per font with the stored byte, the drawn byte and the record.
- The input conversion covers only bytes typed through the six owners `TEXT-NAMEIN-024` counts. A shipped string byte reaches the display conversion alone: stored `0x80..0xAF` gives record `byte + 0x10`, `0xE0..0xEF` gives `byte − 0x10`, other bytes `byte − 0x20`.
- Fonts 1, 2, 4 and 5 have 224 records (`font3.16` has 64: trailer `0x40`, 1073 bytes), so all 64 typed Cyrillic bytes reach a record inside those four tables. In font3 all 64 reach records 144..223, outside its table, with no clamp (`TEXT-INDEX-003`).
- Over the 18 Latin and Cyrillic homoglyph pairs of `SPR16A-FONT-020`, the best single constant `k` in `record = t − k` over `k = 0..255`, counting records 96..223, is `0x30` with 13 of 18 pixel-identical in font1 and font2. The converters as read give 17 of 18 (`TEXT-109`).

**Confidence.** **High** for the arithmetic: both converters are read whole on both roots and the table is enumerated over all 256 typed values.

**Unknown.** The code page of the byte the keyboard layer hands to the input conversion. The image names none. The two input ranges equal the Cyrillic positions of `CP1251`, and the stored result equals the positions of `CP866`, so `CP1251` is a naming convention for the measured ranges. Whether a given Windows installation delivers `CP1251` bytes to the field was not observed.

### TEXT-109

- Method: for each of the 18 pairs (A/А through x/х), the record reached by the typed Cyrillic letter under a rule is compared cell by cell with the Latin letter's record in the same font. `evidence/cp1251-test-ru.txt` lists every pair.
- font1: the converters as read give 17 of 18 pixel-identical (`М` differs by 2 cells), `byte − 0x20` gives 4 of 18, `byte − 0x30` 13 of 18, `byte − 0x10` 0 of 14, `byte − 0x40` 0 of 18. font2: 17 of 18 (`М` differs by 25 cells), 4, 13, 6 of 14 and 0 of 18.
- font4 and font5 draw Cyrillic in a different style from Latin: 5 of 18 and 0 of 18 under the converters, 3 of 18 and 0 of 18 under `byte − 0x20`. The pair test does not discriminate there. The `.dat` advance of the 64 Cyrillic records correlates with font1's at the same record (Pearson r 0.822 font4, 0.593 font5) more than at any of the 63 other cyclic alignments (largest 0.553 and 0.336). The Latin control over 52 records gives 0.833 and 0.724. For font2 the same-record r is 0.370 and the shift-by-32 alignment, which swaps the capitals with the lowercase, gives the same 0.370, so the advance test does not discriminate there and font2 rests on the pixel pairs.
- The arrangement is the same on both roots. The 64 Cyrillic records (144..191, 208..223) have identical decoded glyph hashes and advances in all five fonts on EN and RU. RU `font2` differs from EN `font2` at 35 records, 96..122, 128..133, 193 and 207, each blank with advance 0 in RU; none is a Cyrillic record. The RU `font2.dat` therefore differs from the EN file.
- `font5` is loaded by nothing in the install (`SPR16A-FONT-018`); its row is a property of the file.

**Confidence.** **High** for font1 and font2: 17 pixel-identical pairs at the converter records against 4 at the unconverted records, with the capitals in alphabet order over `0xB0..0xCF` (`SPR16A-FONT-020`). **Medium** for font4 and font5: the advance profile of the 64 records matches font1's at the control's strength, but a correlation does not name a letter. The pair identity of `М` is not exact in font1 or font2.

**Unknown.** The letter drawn at each record in font4 and font5 by pixel evidence. Which text surface draws with font3, which holds no Cyrillic record.

### TEXT-110

- Selector 0, the English `main/id` (`english 0`): both conversions return their argument, so a typed byte `t` reaches record `t − 0x20`. A typed `CP1251` byte `0xC0..0xCF` reaches records 160..175, which hold Р..Я. `0xD0..0xDF` reaches 176..191 (а..п). `0xE0..0xEF` reaches 192..207. `0xF0..0xFF` reaches 208..223 (р..я).
- Counts over the 64 typed bytes `0xC0..0xFF`: 16 draw their own letter (`0xF0..0xFF`), 32 draw another Cyrillic letter (А..п draw Р..п of the arrangement, shifted by 16 letters), and `0xE0..0xEF` reach the blank region. In font1, font4 and font5 records 192..207 are blank except 193 (`ß`): 15 blank and one `ß`. In EN `font2` record 207 also carries a glyph, so 14 blank, `ß` and one other glyph.
- Reading `0xC0..0xFF` as the 64 Latin-1 characters `À..ÿ`, none reaches a record whose arrangement label is that character: 0 own, 15 blank, 49 another label (`evidence/cp1251-test-en.txt`).

**Confidence.** **High** for the records and the blank states, which are measured. **Medium** for the letter named at each record in font4 and font5 (`TEXT-109`).

**Unknown.** Whether any shipped EN text holds a byte `0xC0..0xFF`: the census of `TEXT-ALIAS-011`'s evidence found 58 high bytes in two EN text nodes, which this experiment did not re-run.
