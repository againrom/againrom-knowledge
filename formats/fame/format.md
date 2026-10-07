<a id="fame-hall-of-fame-famehalldat--specification-core"></a>

# FAME hall of fame (`famehall.dat`)

The file stores a count followed by variable-length names and three words per
record. The ordinary reader preserves stored order and both opaque tail
words. Insertion and default seeding are separate operations. The tails are
not padding or required-zero validation fields. — FAME-READER-009,
FAME-INSERT-011, FAME-TAILS-013, FAME-SEED-015

## Wire layout

All integer fields occupy four little-endian bytes. The file has no magic,
version, archive wrapper or compression. — FAME-HDR-001, FAME-REC-002,
FAME-WRITE-007

```text
[count:32]
repeat count times:
    [nameSpan:32][name bytes:nameSpan][score:32][tail1:32][tail2:32]
```

The writer emits `4 + sum(16 + nameSpan)` bytes. Exact end-of-file after the
records is a writer and shipped-file property; the reader does not check for
trailing bytes. — FAME-REC-002, FAME-READER-009

| Field | Ordinary writer / producer | Reader / consumer | Claims |
|---|---|---|---|
| `count` | Current array count; not a constant ten | Resizes the record array and controls the record loop; zero clears it | FAME-HDR-001, FAME-READER-009 |
| `nameSpan` | CString stored byte length plus one | Number of bytes requested from the file; not a terminator-validation rule | FAME-STRING-010 |
| `name bytes` | CString bytes including the final NUL | Builds a CString by scanning from the local buffer to the first NUL; that string supplies display text | FAME-STRING-010, FAME-DISPLAY-012 |
| `score` | First record word; computed or seeded by the two located producers | Transferred unchanged; insertion compares signed 32-bit values and display formats with `%d` | FAME-INSERT-011, FAME-DISPLAY-012, FAME-PRODUCER-014, FAME-SEED-015 |
| `tail1` | Zero in both located producers | Read, copied and written unchanged; no direct read in the traced display body | FAME-TAILS-013 |
| `tail2` | Zero in both located producers | Same independent four-byte transfer as `tail1` | FAME-TAILS-013 |

In memory, each record is 16 bytes: a CString data pointer at `+0`, followed
by the three words at `+4/+8/+c`. The containing object has its array pointer
at `+134`, count at `+138`, and insertion limit at `+12c`. Its constructor sets
that limit to ten and initializes an empty array. — FAME-REC-002,
FAME-READER-009, FAME-INSERT-011

## Loading, strings and writing

The startup path opens `famehall.dat` for reading. An open failure or a file of
zero bytes takes the default-seeding path. A nonempty file goes to the reader:
a four-byte zero count produces an empty table, not seeded defaults. The
writer uses raw file writes; the previously traced exit handler opens the
file with create/write mode. — FAME-SEED-015, FAME-WRITE-007

The reader neither sorts the records nor clamps their count to ten. It reads
all three words without score-range or zero-tail tests. A loaded table can
therefore retain ties, nonzero tails and a different order until another
operation changes it. Signed loop/allocation arithmetic means that absence of
a guard is not a promise that every 32-bit count succeeds. — FAME-READER-009,
FAME-INSERT-011, FAME-BOUNDARY-016

The name buffer occupies 1,024 bytes. The reader passes the supplied prefix
directly to `Read`, then passes the buffer to a char-string constructor that
calls `lstrlenA`. There is no local ASCII, positive-length, length-bound or
final-NUL check. A well-formed ordinary name has one terminating NUL inside
that buffer. A 1,024-byte span with 1,023 non-NUL bytes followed by NUL fits the
observed buffer; it is not an enforced maximum-name validation rule.
— FAME-STRING-010

An early NUL shortens the in-memory string even though the full prefixed span
has been consumed; the writer then emits the shorter string. Non-ASCII byte
values are not rejected here. Zero-length or unterminated spans can reuse
previous buffer content; these instruction examples establish missing
validation, not a portable malformed-name contract. The effective native
encoding, upstream name-entry limits and rendering remain separate questions.
— FAME-STRING-010, FAME-BOUNDARY-016

## Insertion and display

Insertion walks the stored order and inserts before the first existing score
less than or equal to the new score, using signed comparisons. A new equal
score precedes the old equal score. If no such position exists, it appends.
The operation copies whole records and trims the array to its configured
limit after insertion. It does not deduplicate names or repair an already
unsorted table. The normal constructor's limit is ten; the load path has no
corresponding clamp. — FAME-INSERT-011

The display initializer copies the current record count. The display body
walks the array in its stored order, formats a one-based rank with `%d.`, passes
the record name to drawing, and formats the score with `%d` before a further
number-formatting helper. Neither trailing word is read directly by that
body. The helper's final text transformation and native rendered pixels are
outside this contract. — FAME-DISPLAY-012, FAME-TAILS-013

<a id="located-producers"></a>

## Score producers

The non-default producer takes its name from a char buffer at a live source
object's `+e4`. With signed words `A = campaign+124`, `B = campaign+128` and
`C = source+108`, it computes `C / (A * 10.0) * B` when A is nonzero, or
`C * stored_binary64(2e-6) * B` otherwise, using the observed x87 instruction
order. A helper truncates to a signed 64-bit integer and the producer stores
the low 32 bits as score. Both other words remain zero. The expression follows
the original operation order and does not promise exact rational rounding at
integer boundaries. — FAME-PRODUCER-014

| Input | Established producer and unit | Boundary |
|---|---|---|
| A | Constructor/reset zero it; the admitted mission-completion arm adds signed trunc(simulation sub-ticks/16) before the terminal score call | Groups of 16 sub-ticks; wall-clock seconds and the mission-entry clock baseline are not established |
| B | Constructor/reset zero it; client reception increments for a hostile drawable whose old stage is below 2 and whose new stage is 2,3 or 4 | A received corpse-stage transition, with no killer or once-per-ID test; stage1 death/fall and stage5 do not enter this increment |
| C | Effective state-packet mask4 copies the raw simulation actor+130 dword to drawable+108 | Six-slot experience meaning requires the established Human/Humanoid source path; final cached class and latest delivery are unproved |

Campaign SAVE/LOAD preserves A and B as raw four-byte words. The named addition
and increment have no local range clamp. A freshly constructed CUnit starts
at stage0, so a first received hostile corpse at stage2/3/4 can satisfy the B
increment; repeated 2-to-3 cannot. Resume projects every on-map actor with
no stage baseline, so a loaded hostile corpse at stage2 to 4 can add a count
when the saved B already includes it; this is Medium and native behaviour
remains Unknown. — FAME-021, FAME-022, FAME-028

C's raw projection does not establish a Human-class prerequisite. For the
Human/Humanoid source path, the copied field is the stored aggregate of six
experience slots, maintained by creation/gain/purchase/loss operations. Human
LOAD separately restores the aggregate and slot values. This score path does
not recompute them from levels or sum the party. The six-slot interpretation
for other actor classes or an unproved final cached source remains Unknown.
— FAME-023, HERO-XP-077, SAV-HEROXP-063

The source cache is at `[[frame+d0]+3f54]`. One receiver assignment selects a
new drawable when its map count is 1, without an ownership predicate. Preview
paths can also populate this cache, transfer a raw value into C or explicitly
zero C in a temporary derive. A full new-game or LOAD ordering that guarantees
the main hero and the latest experience at final scoring remains Unknown.
— FAME-024

A selected startup initializer requests x87 precision bits0200 (53-bit) and
preserves incoming rounding bits. Its full control word and survival to the
score call are unverified. The final conversion helper forces truncation only
for its signed64 integer conversion and then restores the prior control.
Native rounding, exceptional conversions and a safe upstream gameplay maximum
remain Unknown. — FAME-025

The fallback producer obtains names from UI string-table indices 263 through
272. For the first nine it starts a score base at 70000, subtracts 7000 per
row and adds a random-derived remainder modulo 5000. It assigns zero score
to the last row. Both trailing words stay zero and all ten records pass
through the same insertion routine. — FAME-DEFAULT-008,
FAME-SEED-015

The cached source is the drawable registered when the client map was empty,
not a chosen hero. Experience reaches it only through mask4 state packets that
the sender's owner filter passes. Mission entry projects the participant's own
hero first, so on resume the hero is the likely source (Medium). The actor-add
announcer at `R0811` skips every player whose human flag is zero and the
flag has no explicit byte store outside that entry, so the announcer emits
nothing before it. On a new game the client waits for the cache to be non-null
before it sends entry and fails at 15 s, so a drawable is registered earlier;
the candidate chain runs from the hero-create packet through the projector and
its owner match, and the first drawable there is Unknown. — FAME-026, FAME-034

No game-code instruction writes the x87 control word, but game code calls C
runtime entries that reload `027f` (53-bit, nearest) and four merge-helper
call sites request `0x133f` (64-bit, nearest) with unaudited restores. The
persistent setter requests 53-bit precision. The control word at the score
call is Unknown. On a grid of A 0..32, B 0..32 and 1199 values of C, 53-bit
against 64-bit precision changes the stored score on 1471 of 1,305,711 grid
points; the real domain is not established, so this is a grid count and not a
frequency. — FAME-027

The client packet pump runs right after the world step at five of the 11
world-step call sites, so a corpse stage change there is counted in the same
frame. Resume projects the participant's hero, the global actor list
unfiltered and world-list actors only when their stage word is below 5, with
no stage baseline. — FAME-028

The credits roll one pixel per draw call that finds 23 ms or more since the
last step, starting 480 pixels down, and end when the first visible line index
reaches the line count, so a run is 480 plus 15 per text line steps; the line
height is the first record of the `font1` `.16` file and is the same in EN and
RU, while the line count differs. The run time is at least the step count less
one times 23 ms (82432 ms EN, 68287 ms RU) and is otherwise Unknown. — FAME-030

Credits end on any key-down or a left-button press, or at the end of the
scroll. The hall of fame ends on Enter through its button (Medium), or on a
left-button release in its rectangle; the base key handler acts only on Tab
and the four arrow keys; the hall's key slot, its button and the base handler
do nothing on Escape, and other root-level handlers were not enumerated (Medium). The earlier FAME-029 clause that Enter and
Escape reach a base action is withdrawn. — FAME-029, FAME-031

The terminal route and the new-game arm clear the campaign record through the
same routine, which itself sets step 10; they differ only in frame state, and the
terminal route also clears the client map and the cached source. — FAME-033,
SAV-1159

SAVE has two callers; message442 is the one not excluded by SAV-972's save
route. The frame child with id 442 is the main-menu panel; its mouse-up slot
posts no 442, the only literal post site follows a mode-2 load handler, and
modal mouse capture keeps mouse input from the panel (Medium). The earlier FAME-029 clause
that this child is a second 442 producer is withdrawn. — FAME-029, FAME-032,
SAV-972

<a id="sample-facts-and-remaining-boundaries"></a>

## Unknowns

Which producer of message41d is the Victory acknowledgment, the credits
wall time (the timer granularity is Unknown), and whether the
local transport returns a packet to the pump of the step that sent it remain
Unknown. The FAME-029 duration and hall-key Unknowns are settled and amended
in FAME-030 and FAME-031. — FAME-026, FAME-028, FAME-029

Names are not restricted to ASCII and loaded scores need not be strictly
descending. Those earlier clauses of FAME-NAME-003 and FAME-SCORE-004 are
partially retracted. Installed content does not establish prior-play history.
— FAME-NAME-003, FAME-SCORE-004, FAME-DEFAULT-006, FAME-DEFAULT-008

Allocation failure, short reads, overflowing lengths/counts, exceptions and
string lifetime are outside the ordinary successful-transfer contract.
The effective name encoding, upstream limits and final display transformation
remain Unknown. Neither tail word has an established level, mission,
difficulty or time meaning; a full native nonzero-tail lifecycle and complete
score-production session remain unverified. — FAME-TAILS-013, FAME-BOUNDARY-016

## Read and write sequence

1. Read u32 count and process that many records in stored order.
2. Read u32 nameSpan and that many name bytes. The ordinary CString value
   ends at the first NUL; the complete span is still consumed.
3. Read score, tail1 and tail2 as independent four-byte words.
4. Retain order, ties and tails. Loading does not sort, trim or deduplicate.

To write, emit the current count, then each CString's byte length plus one,
its bytes and final NUL, and all three stored words. The output size is
`4 + sum(16 + nameSpan)`. Use zero tails only for the two specified record
producers; preserve them in ordinary record transfers. — FAME-REC-002,
FAME-WRITE-007, FAME-STRING-010, FAME-TAILS-013
