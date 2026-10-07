# Claim registry — FAME (famehall.dat)

Level 2 ledger. Index: [registry.md](registry.md) · spec:
[`formats/fame/format.md`](../formats/fame/format.md).
IDs are permanent. `✔ promoted` means the claim is reflected in the spec;
`● active (amended)` and `● partially retracted` require reading the correction
scope in [retracted.md](retracted.md).

The shipped EN/RU copies contain one distinct table. FAME-HDR-001 through
FAME-DEFAULT-008 retain that sample and the prior writer witness.
FAME-READER-009 through FAME-BOUNDARY-016 supply the current bounded lifecycle:
original static instructions, checked with synthetic instruction vectors under
explicit environment hooks. The older validator's six rejected mutants are
facts about that validator, not original-reader rejection evidence. Unknowns
inside the earlier writer-only row are superseded only to the extent of the
new field-specific rows; no global tail semantics or native score session is
claimed.

FAME-021 through FAME-025 add static upstream producers, conditional source
identity and a bounded startup precision request. Their new measurements decode
files only; no original instruction or process executes. Earlier upstream
Unknowns are superseded only where these rows supply an explicit receiver path.

## Hall of fame record and score lifecycle

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| FAME-HDR-001 | A four-byte little-endian record count opens the file, followed immediately by records; no magic or separate version field. | High | ✔ promoted | [EXP-0012](../experiments/EXP-0012-famehall/), [EXP-0086](../experiments/EXP-0086-saves/), [EXP-0331 reader/writer](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md) |
| FAME-REC-002 | Record framing is `[u32 nameSpan][nameSpan bytes][three 32-bit words]`, size 16+nameSpan; in memory it is a 16-byte record with CString pointer+0 followed by the three words at +4/+8/+c. | High | ● active (amended) | [EXP-0012](../experiments/EXP-0012-famehall/), [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), FAME-READER-009 |
| FAME-NAME-003 | All ten shipped names are NUL-terminated ASCII with prefix=strlen+1. | High | ● active (partially retracted) | [EXP-0012](../experiments/EXP-0012-famehall/), [EXP-0086](../experiments/EXP-0086-saves/), [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), FAME-STRING-010 |
| FAME-SCORE-004 | The first 32-bit word is the record score used for insertion and decimal display. | High | ● active (partially retracted) | [EXP-0012](../experiments/EXP-0012-famehall/), [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), FAME-INSERT-011, FAME-DISPLAY-012 |
| FAME-UNK-005 | The two trailing words contribute 80 zero bytes in the shipped ten-record table. | High / Unknown | ● active (amended) | [EXP-0012](../experiments/EXP-0012-famehall/), [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), FAME-TAILS-013 |
| FAME-DEFAULT-006 | The sample is consistent with the shipped **default** table (10 seeded names, clean descending ladder ending at a `0` entry, all-zero trailing fields) — i.e. | Low | ● active | [EXP-0012](../experiments/EXP-0012-famehall/) |
| FAME-WRITE-007 | **`famehall.dat` is written by a raw `CFile`, on `WM_DESTROY`, and shares nothing with the save path.** `R0801(this, CFile*)` loads `[[file]+0x40]` — the `CFile::Write` slot — into `EBP` once and calls it: first `obj+0x138` as 4 ... | High / Unknown | ● active | [EXP-0086](../experiments/EXP-0086-saves/) |
| FAME-DEFAULT-008 | EN/RU famehall.dat are byte-identical 228-byte files with count 10 and nameSpan=strlen+1 on all ten records. | High | ● active (amended) | [EXP-0143](../experiments/EXP-0143-actor-names/), [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), FAME-SEED-015, FAME-DISPLAY-012 |
| FAME-READER-009 | Complete L03685 reads count, resizes the object+134/+138 array, then reads each prefixed name and all three words through CFile slot+3c. | High / Medium / Unknown | ✔ promoted | [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), evidence/reader.txt, writer.txt, array-size.txt, instruction-vectors.json |
| FAME-STRING-010 | Reader L03686..L03687 passes nameSpan unchanged to Read into a 1024-byte stack buffer, then calls R0473; that helper calls imported lstrlenA before CString construction. | High / Medium / Unknown | ✔ promoted | [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), evidence/reader.txt, string-from-char.txt, writer.txt, local-facts.json, instruction-vectors.json |
| FAME-INSERT-011 | Complete L03688 inserts before the first stored score<=new score using signed JGE at L03689; otherwise appends. | High / Medium | ✔ promoted | [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), evidence/insert.txt, record-copy.txt, campaign-constructor.txt, array-size.txt, instruction-vectors.json |
| FAME-DISPLAY-012 | Complete initializer L03690 obtains campaign via frame+548 and copies campaign count+138 to display+74. | High / Unknown | ✔ promoted | [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), evidence/display-init.txt, display.txt, references.json, local-facts.json |
| FAME-TAILS-013 | The two independently streamed words at record+8/+c are opaque carried values in the traced ordinary lifecycle. | High / Medium / Unknown | ✔ promoted | [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), evidence/reader.txt, writer.txt, record-copy.txt, score-producer.txt, default-producer.txt, display.txt, instruction-vectors.json |
| FAME-PRODUCER-014 | Complete R0802 gets source from frame+d0 then server+3f54; if nonnull, copies the character buffer at source+e4 into the record's CString field at +0. | High / Medium / Unknown | ✔ promoted | [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), evidence/score-producer.txt, score-caller-slice.txt, float-to-integer.txt, local-facts.json, instruction-vectors.json |
| FAME-SEED-015 | Startup slice L03691..L03692 takes fallback L03693 when file-open returns 0 or file length is 0; nonempty input reaches L03685. | High / Medium | ✔ promoted | [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), evidence/startup-caller-slice.txt, file-length.txt, default-producer.txt, random.txt, instruction-vectors.json |
| FAME-BOUNDARY-016 | The measured population is one executable identity from EN/RU, one distinct shipped FAME table, eight complete lifecycle bodies, three caller slices and 12 helper entries. | High / Unknown | ✔ promoted | [EXP-0331](../experiments/EXP-0331-fame-record-lifecycle/RESULT.md), evidence/inputs.json, references.json, local-facts.json, probe.py, instruction-vectors.json |
| FAME-021 | Campaign+124 (A) is zeroed by constructorR0677 and resetR0757. | High / Unknown | ✔ promoted | [EXP-0359](../experiments/EXP-0359-fame-score-sources/EXP-0359.md), selected operations/time_accumulation and campaign_save/load_scalars; SESS-TICK-004, SAV-890, SAV-892 |
| FAME-022 | Campaign+128 (B) is a received hostile corpse-transition counter on the selected client arm. | High / Unknown | ✔ promoted | [EXP-0359](../experiments/EXP-0359-fame-score-sources/EXP-0359.md), selected operations/hostile_transition_increment, tables.json, stage-model.tsv; ANIM-DEATH-007, HERO-REVIVE-068, UNIT-VPLAYER-022, SAV-598, SAV-892 |
| FAME-023 | On actor-state projectionR0059, effective mask4 copies simulation actor+130 as an unconverted dword through appendL03694. | High / Unknown | ✔ promoted | [EXP-0359](../experiments/EXP-0359-fame-score-sources/EXP-0359.md), selected operations/projection_xp_stage, append_dword, packet_payload_pack, client_xp_and_stage; HERO-XP-077, SAV-HEROXP-063, ANIM-MSG-005 |
| FAME-024 | FAME source is the cached client drawable at [[frame+d0]+3f54]. | High / Medium / Unknown | ✔ promoted | [EXP-0359](../experiments/EXP-0359-fame-score-sources/EXP-0359.md), selected operations/source_cache, preview_cache, preview_file_xp, preview_xp_zero, resume_projection, score_inputs; FAME-PRODUCER-014, SAV-890, SAV-948, SAV-HUMPROJECT-461 |
| FAME-025 | The PE entryL03695 reaches callR0803 in its selected prefix; that dispatcher conditionally calls the stored pointerL03696, whose image value isL03697. | High / Unknown | ✔ promoted | [EXP-0359](../experiments/EXP-0359-fame-score-sources/EXP-0359.md), selected operations/startup_initializer_call, precision_request, control_merge, control_precision_mapping, integer_conversion; FAME-PRODUCER-014 |

### FAME-HDR-001

A four-byte little-endian record count opens the file, followed immediately by records; no magic or separate version field. The shipped table has count 10. Writer R0801 writes object+138 and loops over that variable; reader L03685 consumes the same header. Ten is not a constant file-count requirement.

**Confidence.** High for framing and the variable count; sample count is a whole-file measurement

### FAME-REC-002

Record framing is `[u32 nameSpan][nameSpan bytes][three 32-bit words]`, size 16+nameSpan; in memory it is a 16-byte record with CString pointer+0 followed by the three words at +4/+8/+c. Reader and writer transfer these fields in this order. The shipped ten records tile 228 bytes and the writer emits only its record stream. Exact EOF is not a reader acceptance check: trailing input is ignored after count records. The original six rejected mutants tested the repository validator only.

**Confidence.** High for reader/writer framing and finite sample arithmetic; no malformed-I/O guarantee

**Amended.** The complete reader consumes the declared records and returns without comparing its cursor to file length. A synthetic trailing suffix remains unread. The writer's framing, exact shipped-file tiling and the old validator's finite mutant results stand.

### FAME-NAME-003

All ten shipped names are NUL-terminated ASCII with prefix=strlen+1. The writer loads CString nDataLength from [data-8], adds 1 and writes that many bytes; that is its arithmetic, not a reader validation rule. The reader consumes the prefixed span and then scans to the first NUL. ASCII-only and prefix=strlen+1 are withdrawn as universal input constraints.

**Confidence.** High for the sample, writer length arithmetic and corrected reader distinction

**Amended.** The writer uses CString stored length plus one. The reader consumes the prefixed span, then separately calls a NUL-scanning constructor; it does not check ASCII, a final NUL or equality with the scanned length. Early NUL and non-ASCII synthetic controls distinguish those rules. Shipped-name observations stand.

### FAME-SCORE-004

The first 32-bit word is the record score used for insertion and decimal display. The shipped scores descend strictly from 70006 to 0, with two above 65535. Strict descending order is a sample fact, not a format requirement: load preserves order and insertion allows ties, comparing signed 32-bit values. The width remains four bytes; unsigned ranking is not supported.

**Confidence.** High for width and the bounded insertion/display role; sample order is measured, not generalized

**Amended.** Load preserves stored order. Insertion uses signed comparison, puts a new equal score before an old equal score, and does not sort the existing table. Display uses the first word with `%d`. Four-byte width and the shipped table's strict order stand.

### FAME-UNK-005

The two trailing words contribute 80 zero bytes in the shipped ten-record table. Their broader meaning cannot be decoded from that sample. The earlier reserved/unknown label supplied no reserved-zero rule; the current ordinary transfer and producer roles are FAME-TAILS-013.

**Confidence.** High for the zero-byte measurement; Unknown for a semantic label beyond the traced lifecycle

**Amended.** Reader, copy and writer carry both words independently. Both located producers zero them; the complete display body reads neither directly. The sample measurement stands and no broader semantic label or universal zero rule is established.

### FAME-DEFAULT-006

The sample is consistent with the shipped **default** table (10 seeded names, clean descending ladder ending at a `0` entry, all-zero trailing fields) — i.e. not yet modified by play. Context, not a structural claim

**Confidence.** Low

### FAME-WRITE-007

**`famehall.dat` is written by a raw `CFile`, on `WM_DESTROY`, and shares nothing with the save path.** `R0801(this, CFile*)` loads `[[file]+0x40]` — the `CFile::Write` slot — into `EBP` once and calls it: first `obj+0x138` as 4 bytes (the count), then per record `strlen+1` as 4 bytes, the string that many bytes, and `+0x04`/`+0x08`/`+0x0c` as 4 bytes each, striding `0x10` through the array at `obj+0x134`. **No `Asg&` header, no `CArchive`, no class record, no tag, no compression** — every one of which the save writer `R0804` uses. Its **only** caller is `R0805` (`EnumRefs callto:R0801`: 1 hit, 1 owner, 0 orphan), which opens the file `CFile::Open("famehall.dat", 0x1001, NULL)` = `modeCreate\|modeWrite` and is itself reached only through `L03698` (`EnumRefs callto:R0805`: 1 hit, 0 owners, 0 orphan) — a `pfn` slot in the six-dword message-map entry array at `L03699` whose `nMessage` dword is **2**, `WM_DESTROY`. So “written in the same minute as the last save” is the application exiting, not a shared writer. **And the round trip is byte-exact, twice**: the game rewrote the file at 00:24 and again at 00:40 with fresh mtimes, and both results hash `1845f772…1cbe9d`, **identical to the pristine EN root and to the pristine RU root** (the 00:40 copy was compared with `cmp` — 0 differing offsets over 228 bytes — and then discarded, per the corpus manifest's own addendum). That addendum asks what the game did at 00:40 that rewrites the hall of fame without changing it; this row answers it — **it exited** — a read‑modify‑write reproducing its own input, which is the test a single shipped sample could never give

**Confidence.** **High** for the writer, the mode, the trigger and the record arithmetic (one routine read whole; the `CFile::Write` slot is loaded once and used for every field; both caller sets are `EnumRefs callto:` with 0 orphan; the message-map record is read out of `.rdata`) / **High** that the shipped bytes survive a round trip (three sha256 over three copies, one of them written by the game) / **Unknown** what the three per-record dwords mean beyond `FAME-SCORE-004`'s width — the writer names nothing, and nothing in this session changed the table

### FAME-DEFAULT-008

EN/RU famehall.dat are byte-identical 228-byte files with count 10 and nameSpan=strlen+1 on all ten records. Their names equal UI string-table indices 263..272, unchanged between those localized tables. These are two copies of one distinct content sample. The former choice between file seeding and independent UI-table drawing is now resolved on the located paths: fallback L03693 populates records from those indices, while display R0806 reads record names. This does not establish the exact RNG state or history of the shipped table.

**Confidence.** High for whole-file identity, the retained source-table correspondence and the bounded source/display resolution

**Amended.** The fallback body constructs records from UI indices 263..272; the complete display body takes names from those records. Missing/zero-byte files select fallback, while a nonempty zero-count file loads an empty table. No exact shipped RNG state or native reseed witness is claimed.

### FAME-READER-009

Complete L03685 reads count, resizes the object+134/+138 array, then reads each prefixed name and all three words through CFile slot+3c. Count 0 clears the array. No local sorting, ten-record clamp, score/tail validation or EOF check occurs. Signed loop/allocation arithmetic and unchecked reads prevent interpreting this as acceptance of every 32-bit count. Thirteen records, unsorted signed extremes/ties and nonzero tails survive the bounded instruction reader/writer vectors unchanged; extra suffix bytes remain unread.

**Confidence.** High for complete local control/data flow; Medium for conditional instruction execution as evidence beyond that static body; exceptional inputs Unknown

### FAME-STRING-010

Reader L03686..L03687 passes nameSpan unchanged to Read into a 1024-byte stack buffer, then calls R0473; that helper calls imported lstrlenA before CString construction. No local positive-length, maximum, ASCII or terminator check exists. Early NUL normalizes the next write; non-ASCII bytes are copied. A 1024-byte NUL-terminated span fits. Synthetic unterminated/zero-length second records reuse prior buffer bytes, establishing missing validation under the probe environment rather than universal malformed-file behavior. Writer length is CString nDataLength+1.

**Confidence.** High for prefix-versus-scan and absent local checks; Medium for buffer-reuse instruction examples; native encoding, upstream name limits and failures Unknown

### FAME-INSERT-011

Complete L03688 inserts before the first stored score<=new score using signed JGE at L03689; otherwise appends. Ties put the new record first. Copies/shifts preserve the complete 16-byte record, and the final comparison trims to object+12c (constructor default 10). There is no name deduplication or full sorting of old rows. The pointer/record-copy and comparison mechanics are read; vector outcomes assume explicit allocator/string/copy hooks and normal return from untraced helper L03700 on the newly appended empty slot before memmove.

**Confidence.** High for comparison, direct copy stores and limit branches; Medium for complete successful insertion across the untraced helper boundary

### FAME-DISPLAY-012

Complete initializer L03690 obtains campaign via frame+548 and copies campaign count+138 to display+74. Complete body R0806 traverses 16-byte records in stored order, formats one-based rank with `%d.`, passes record name+0 to drawing, and formats record+4 with `%d` before helper R0807. It reads neither record+8 nor+c directly. Stored body pointers occur at L03701/L03702. The final number helper, draw callbacks and native pixels were not traced or executed.

**Confidence.** High for initializer/body field uses and literal formats; downstream text transformation and full UI lifecycle Unknown

### FAME-TAILS-013

The two independently streamed words at record+8/+c are opaque carried values in the traced ordinary lifecycle. Reader L03685, writer R0801 and record-copy L03703 preserve each; both found producers R0802/L03693 initialize each to 0 and do not change either before insertion. Complete display R0806 reads neither directly. Nonzero synthetic values survive read/write and surviving insertion records under the documented environment hooks. These are not padding, a proved reserved-zero constraint, or proved globally unused storage. No semantic level/mission/difficulty/time interpretation or other nonzero producer is established.

**Confidence.** High for the named bodies and independent transfer/zero stores; Medium for hooked insertion survival; Unknown for other consumers, native nonzero lifecycle and broader purpose

### FAME-PRODUCER-014

Complete R0802 gets source from frame+d0 then server+3f54; if nonnull, copies the character buffer at source+e4 into the record's CString field at +0. With signed A=campaign+124, B=campaign+128, C=source+108, its x87 operations compute C/(A*stored 10.0)*B if A!=0, else C*stored_binary64(2e-6)*B. Helper R0279 selects truncation for signed 64-bit FISTP; the low dword of the result becomes the first word. Both remaining words retain 0. The sole located direct call is L03704, inside the retained caller slice; broader scheduling and upstream meanings are not resolved. Five instruction examples declare FPCW 037f; the A 0/B3/C1000000 example gives 5, so exact-rational rounding or native initial FPU state must not be inferred.

**Confidence.** High for immediate source offsets, local arithmetic, conversion and zero stores; Medium for conditional numeric examples; upstream meaning, overflow/exception and native FPU state Unknown

### FAME-SEED-015

Startup slice L03691..L03692 takes fallback L03693 when file-open returns 0 or file length is 0; nonempty input reaches L03685. Complete seed body uses UI indices 263..271 with base 70000 decreasing by 7000, plus a random-derived signed remainder modulo 5000; index 272 gets score 0. All ten names are copied into records, both trailing words remain 0, and all records use L03688. The random helper returns 0..32767; exact shipped RNG state is not established. Synthetic startup vectors select ten seeds for missing/zero-byte files and zero records for a four-byte zero header.

**Confidence.** High for branch predicates, source indices and producer stores; Medium for startup with file/allocator environment hooks; no native deletion/reseed witness

### FAME-BOUNDARY-016

The measured population is one executable identity from EN/RU, one distinct shipped FAME table, eight complete lifecycle bodies, three caller slices and 12 helper entries. Whole-text E8/E9 byte-form and all-section literal-DWORD censuses find three insert call sites in two producer bodies, but do not exclude computed calls or aliases. The 26 instruction vectors execute original local bytes with memory-only hooks, including untracedL03700, storage/lifetime, file, allocator and memory-copy boundaries. They establish conditional discriminators, not a native process, OS I/O, UI, exception or universal malformed-input contract. No owner artifact or original process was required.

**Confidence.** High for the measured population and method boundary; unresolved behavior explicitly Unknown

### FAME-021

Campaign+124 (A) is zeroed by constructorR0677 and resetR0757. Frontend message425 calls that reset with frame+548. In the admitted message41d completion arm (mode other than 0/1/3, frame+3bc nonzero), L03705..L03706 adds signed trunc(N/16), N=[[L00285]+4], to the current A with ordinary 32-bit addition. SESS-TICK-004 identifies N as simulation sub-ticks; this is accumulation in groups of 16 sub-ticks, not proved wall-clock seconds or exact equality to the separate full-tick counter. The addition precedes SAV-890's final-score predicate/call. Campaign SAVE/LOAD transfer A unchanged as four bytes; no local clamp or positive-value test constrains the loaded word or this addition.

**Confidence.** High for named stores, signed division and admitted ordering; global writer closure, mission-entry clock baseline and native elapsed-time/session totals Unknown

### FAME-022

Campaign+128 (B) is a received hostile corpse-transition counter on the selected client arm. DispatcherR0509 saves drawable+15a before mask8 applies the incoming stage. The tableL02543 sends new stages2/3/4 toL02767; only old unsigned stage<2 and relation word mask1 admit add1/store atL03707, through frame+548. The relation is local CPlayer+38 indexed by drawable owner CPlayer+4 (UNIT-VPLAYER-022). Stage1 is the death/fall arm; stage5 has another arm (ANIM-DEATH-007). This increment checks no attacker, killer or prior-ID ledger. Constructor/reset zero B; LOAD restores its raw word. CUnit constructorR0593 sets stage0, so a newly admitted CUnit receiving2/3/4 can satisfy the arm; repeated 2-to-3 does not. A later reset below 2 can permit another increment. No local increment clamp exists.

**Confidence.** High for the conditional local transition and counter stores; Unknown for native reconstruction ordering, initial-corpse double counting, global totals and a once-per-death interpretation

### FAME-023

On actor-state projectionR0059, effective mask4 copies simulation actor+130 as an unconverted dword through appendL03694. The state-packet constructor binds vtableL03708 slot8 to packerL03709, which selects6e/6f/70/6c by mask width and passes the current payload bytes to transport packing. ClientR0509's common state arm decodes that flag4 dword toL03710 and stores it at drawable+108 (L03711); without flag4 the copy is skipped. This raw transfer has no established Human-class prerequisite. On the Human/Humanoid source path established by HERO-XP-077 and SAV-HEROXP-063, the field carries the independently stored six-slot experience aggregate; the Human LOAD restores its raw signed aggregate without local recomputation from saved skill slots. Its six-slot meaning for other actor classes or an unproved final cached source remains Unknown. Transport success, object continuity and final delivery order are conditional.

**Confidence.** High for the original producer/packet/client field copies and the qualified Human/Humanoid experience meaning; other-class six-slot interpretation, final cached class and native latest-value delivery Unknown

### FAME-024

FAME source is the cached client drawable at [[frame+d0]+3f54]. State reception stores the newly constructed drawable there when its map count+9c4 equals 1 (L03712..L03713); the assignment has no ownership/selection predicate. Preview paths also allocate CUnit and assign this cache; character-file reading transfers16 bytes at source+fc, including+108, while temporary deriveR0808 explicitly zeros source+108. These are distinct source producers, not a universal fresh-game C=0 guarantee. Resume setup requests all-state projection of the then-current Player+34 (SAV-948), but initial packet ordering/cache identity and the last delivered experience/corpse update before finalR0802 are unproved. The final call uses current cached fields after the completion-time addition; an intervening source virtual call inR0809 remains a side-effect boundary.

**Confidence.** High for selected cache assignments and local final use; Medium for the conditional cached-actor lifecycle; native main-hero identity, preview replacement and latest-value timing Unknown

### FAME-025

The PE entryL03695 reaches callR0803 in its selected prefix; that dispatcher conditionally calls the stored pointerL03696, whose image value isL03697. After returning helpers, that initializer callsL03714, which requests value10000 under mask30000 throughL03715/L03716. The conversion pair maps this request to x87 precision bits0200 (53-bit) while preserving incoming rounding bits. This is a conditional local startup precision request, not a measured full control word or proof it survives to FAME. Separate FNINITL03717 and other control setters have no established path to the score call here. The score still follows FAME-PRODUCER-014's signed A/B/C loads and x87 order;R0279 temporarily forces truncation for signed64 FISTP and restores prior control, with low EAX stored as score. No native rounding, exception/overflow outcome or safe gameplay maximum is established.

**Confidence.** High for selected startup/control mapping and conversion instructions; native initial/final FPCW, exceptional conversions and complete earned-score session Unknown


## Earned score source and campaign ending

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| FAME-026 | The FAME source is one cached drawable: the one registered when the client map was empty, not a chosen hero. Resume: the hero (Medium); new game: Unknown (FAME-034). Experience arrives only by mask-4 packets that pass the owner filter. | High / Medium / Unknown | ● active (amended) | [EXP-0448](../experiments/EXP-0448-score-ending/EXP-0448.md); FAME-034 |
| FAME-027 | No game-code instruction writes the x87 control word, but game code calls C runtime entries that do. The word at the score call is Unknown; on a 33 x 33 x 1199 grid 53-bit and 64-bit precision differ on 1471 points. | High / Unknown | ✔ promoted | [EXP-0448](../experiments/EXP-0448-score-ending/EXP-0448.md) |
| FAME-028 | At five of 11 world-step call sites the client pump runs adjacent to the step. Resume projects Player+34, the unfiltered global list and world-list actors with actor+13c below 5, with no stage baseline. | High / Medium / Unknown | ✔ promoted | [EXP-0448](../experiments/EXP-0448-score-ending/EXP-0448.md) |
| FAME-029 | Credits end on a key, a left-button press or the scroll end; the hall ends on its button (FAME-031). Message 442 is the only SAVE caller not excluded by SAV-972; its frame child's mouse-up posts none (FAME-032). | High / Medium | ● active (partially retracted) | [EXP-0448](../experiments/EXP-0448-score-ending/EXP-0448.md); SAV-972 |
| FAME-030 | The credits roll one pixel per draw call that finds 23 ms or more since the last step; a run is 480 pixels plus lines times the font1 line height (15 in EN and RU). Wall time above the bound is Unknown. | High / Medium / Unknown | ● active | [EXP-0461](../experiments/EXP-0461-ending-remainder/EXP-0461.md) |
| FAME-031 | Credits end on any key-down or a left-button press. The hall ends on Enter via its button or a left-button release in its rectangle. The hall's key slot, button and base handler do nothing on Escape (Medium). | High / Medium | ● active | [EXP-0461](../experiments/EXP-0461-ending-remainder/EXP-0461.md) |
| FAME-032 | The frame child with widget id 442 is the main-menu panel; its mouse-up slot posts no 442. Message 442 has one literal post site, after a mode-2 load handler. Modal mouse capture is Medium. | High / Medium | ● active | [EXP-0461](../experiments/EXP-0461-ending-remainder/EXP-0461.md) |
| FAME-033 | The terminal and new-game routes clear the campaign record through the same routine, which itself sets step 10, so the fields they reset and load are identical. | High / Medium | ● active | [EXP-0461](../experiments/EXP-0461-ending-remainder/EXP-0461.md) |
| FAME-034 | The announcer skips players whose human flag is zero; no explicit byte store sets it outside participant entry. On a new game a candidate chain registers a drawable before entry; the first drawable is Unknown. | High / Unknown | ● active | [EXP-0461](../experiments/EXP-0461-ending-remainder/EXP-0461.md) |

### FAME-026

Instrument: capstone decode of the matched EN/RU `rom.exe` (SHA-256 in `evidence/inputs.tsv`), listings in `evidence/listings.txt`, address assertions in `evidence/checks.tsv`.

Cache assignment. The cached drawable is the one registered when the map held one element, which is the first registration on an empty map. Receiver `R0509` stores the drawable it just built or found into `[frame+d0]+3f54` at `L03713`. The store runs only in the new-object arm, after the drawable is registered, when the map element count at `frame+9c4` equals 1 (`L03718`, `L03719`). The count is the element count of the id-keyed map at `frame+9b8`. The store has no class, owner or hero test. The score producer `R0802` reads only this pointer (`L03720`) and loads `source+108` (`L03721`); it reads neither `Player+34` nor any other drawable. The score therefore uses one actor's experience aggregate, not a party sum.

Writers of the cache word, from the census of all decoded `.text` (`evidence/cache-field-writers.tsv`, 9 stores): zero at `L03722`, `L03723`, `L03724`; the receiver store at `L03713`; zero at `L03725` after the map is cleared; zero at `L03726` in the terminal reset route (FAME-029); and three preview-object stores at `L03727`, `L03728`, `L03729` (FAME-024).

Experience delivery. `source+108` is written only at `L03711`, from the flag-4 dword of a state packet. The sender `R0059` builds that dword from `actor+130` (`L03730`). Its owner filter removes bit 4 by `and 0x507b` (`L03731`) when the actor's owner `actor+14` differs from the recipient and the actor's virtual slot `+30` returns 0. When that slot returns nonzero, it removes bit 4 by `and 0x50fb` (`L03732`) unless the kind word `actor+0e` lies in `0x21..0x3f` (`L03733`, `L03734`). An actor owned by the recipient is never filtered. The gain routine `R0810` returns without effect unless `actor+0e` lies in `0x21..0x3f` (`L03735`, `L03736`) and, after updating `+130`, sends mask 4 (`L03737`) or `0x1f001304` on a level change (`L03738`) with recipient 0, which `R0059` expands to every eligible player.

Which actor is first. Entry routine `R0131` (mission entry and resume, SAV-948) projects `Player+34` with mask -1 at `L03739`, then calls `L03114` (`L03115`), which projects `Player+34` again before every other actor (`L03740`, FAME-028). `Player+34` is the participant's own named hero (SAV-HERO-059). The actor-add announcer `R0811` also sends mask -1 to human players (`L03521`, six callers in `evidence/calls.tsv`); the announcer is traced in FAME-034: it emits nothing to the participant before entry, but on the new-game path a drawable is registered before entry, and the candidate producer chain is in FAME-034.

**Amended.** The announcer route is traced (FAME-034): it emits nothing to the participant before entry. On a new game the client waits for a non-null cache before it sends entry, so the first drawable there is not shown to be the hero's. The candidate producer chain is in FAME-034; its projection arguments were not read, so the identity stays Unknown. The cache assignment, writer census, sender filter and gain gate stand.

**Confidence.** High for the cache assignment and its predicate, the full writer census, the single reader and field, the sender's filter masks and the gain gate, each read at instruction level. Medium for the hero being the cached drawable on resume: it follows from the order inside two routines, with no run. On the new-game path it is Unknown (FAME-034). Unknown: whether the local transport delivers every packet in send order, which route registers the first drawable on the new-game path, and the meaning of the virtual slot `+30`.

### FAME-027

Instrument: capstone linear sweep of `.text` (no control-flow recovery); control-word instruction census `evidence/control-census.tsv` and `evidence/control-by-routine.tsv`; raw-byte cross-check `evidence/control-rawbytes.tsv`; caller tables `evidence/calls.tsv` and `evidence/dword-refs.tsv`; imports `evidence/imports.tsv`; arithmetic model `evidence/arith-probe.tsv`.

Inventory. The sweep decodes 71 x87 control or environment instructions (control-word and environment load, store, save, restore and initialise forms, plus the SSE and XSAVE control-register loads). The sweep attributes each to the preceding standard frame-setup sequence (column `prologue_attributed`); frameless library routines have no such prologue, so the grouping in `control-by-routine.tsv` is attribution and not routine boundaries. One, at `L03741`, is a jump-table dword decoded as an instruction (`evidence/data-words.tsv`). The other 70 lie at or above `R0544`, the C runtime range. The raw-byte scan of `.text` for the memory-operand control-word, environment-restore, SSE control-register and x87-initialise byte patterns finds 97 matches, 45 below `R0544`: 44 lie inside other instructions and one is the jump-table data above. No game-code instruction below `R0544` writes the control word.

Game code calls runtime entries that do. `R0812` (called from `L03742` and `L03743`) reaches `fldcw [L03744]` at `L03745`. `L03746` (called from `L03747`) and `L03748` (12 sites including `L03749` and `L03750`) reach the reloads of the same constant at `L03751` and `L03752`. The constant at `L03744` is `027f`: precision 53-bit, nearest rounding, all exceptions masked. Four of the 20 call sites of the merge helper `L03753` request `0x133f` (64-bit precision, nearest rounding): `L03754`, `L03755`, `L03756` and `L03757`, each preceded by `push 0x133f` (`push_133f` rows). The other 16 push a saved word; whether they restore it was not audited.

Persistent setter. `L03716` merges a request into the hardware word as `(value & mask) | (old & ~mask)` (`L03758`..`L03759`). Its only direct caller is wrapper `L03715`, whose only caller is `L03714`, which passes the mask 0x30000 and the value 0x10000 (`L03714`, `L03760`). `L03714` is called from the startup initializer `L03697` (`L03761`) and from the x87-initialise entry `L03762` (`L03763`). Entry `L03762` has no direct call and no dword reference in any section. Value 0x10000 under mask 0x30000 requests precision 53-bit (FAME-025). The mask leaves rounding and exception bits as inherited. Conversion helper `R0279` saves the word (`L03764`), sets the rounding field (bits 10-11) to 0b11 by OR-ing 0xc into the control word's high byte (`L03765`), runs the integer-store conversion to a 64-bit slot and reloads the saved word (`L03766`). The producer stores the low dword.

Probe. The model rounds each x87 result to a 53-bit or 64-bit significand in the order of `L03767`..`L03768` (FILD B, FILD C, A times 10.0, divide, multiply by B; or C times binary64 2e-6 times B when A is 0) and truncates. The grid is A 0..32, B 0..32 and the 1199-value C list `range(0,3001,3)` plus `range(3000,200001,997)` (3000 appears twice): 1,305,711 points. On this grid 53-bit and 64-bit precision give different scores on 1471 points, 53-bit differs from the exact rational on 648 and 64-bit on 833. Example A=1, B=15, C=42: exact 63, 53-bit 63, 64-bit 62. The real domain of A, B and C in play is not established, so 1471 is a count on this grid and not a frequency.

**Confidence.** High for the instruction-level absence of control-word writes in game code below `R0544`, the game-to-runtime calls listed, the setter's merge formula and call graph, and the conversion helper. Unknown for the control word at the score call: the runtime entries above set `027f` or `0x133f`, the 16 saved-word callers were not audited for restoring it, and the model assumes ideal x87 rounding without exponent effects. Unknown: the control word at process start (the loader and the 14 imported DLLs listed in `evidence/imports.tsv`, ADVAPI32, COMCTL32, DDRAW, DSOUND, GDI32, KERNEL32, SHELL32, USER32, WINMM, WINSPOOL, WSOCK32, comdlg32, ole32 and smackw32, run code outside this image), rounding and exception bits at the call, and any entry into `L03762`.

### FAME-028

Instrument: capstone decode; listings of the pump, step and resume bodies; `evidence/calls.tsv`.

Order. The client pump `R0509` polls the queue through `R0513` in a loop (`L03769`, exit at `L02956`/`L02947`). A packet whose opcode byte equals the argument ends the call with 1 (`L02946`); with an empty queue the call returns 0 when the argument is 0 and bit 0 of `frame+3dc` is clear (`L03770`..`L03771`), and otherwise polls again (`L02950`). The world-step routines are called directly from 11 sites (`R0192` at `L03772`, `L03773`, `L03774`, `L03775`, `L03776`, `L03777`, `L03778`; `R0147` at `L03062`, `L02626`, `L03779`, `L03780`). The pump is called right after the step at five of them (step, then pump): `L03773`/`L03781`, `L03774`/`L02962`, `L03776`/`L03782` (world step `R0192`) and `L03062`/`L03063`, `L02626`/`L02960` (world step `R0147`). The step body `R0193` increments the world tick, runs actor subticks and ends with the queue flush `R0611`; actor state packets, including stage mask 8, are produced inside it. The client reads the old stage at `L02542` before applying a mask-8 stage at `L02541`, so a stage packet produced by step N is counted by the pump that follows step N. The other pump sites `L03783`, `L03784`, `L03785`, `L03786` and `L03787` follow a timer wait or poll loop, not an adjacent step.

Loaded corpses. Resume `L03114` projects `Player+34` first (`L03740`). It then walks the global actor list `[L00240]` (MOVE-TICK-013) and projects each actor with mask -1 and no stage test (`L03788`, `L03789`). It then walks the world list at `[[L00285]+14]+c` and projects each actor whose `actor+13c` is below 5 (`L03790`, `L03791`). The mask carries the stage (bit 8 survives both filters). A drawable built by `R0593` starts at stage 0 (FAME-022), so a hostile corpse at stage 2, 3 or 4 meets the increment arm; the arm sends 0 to 2/3/4 to `L02767` with old stage below 2. No instruction of the resume bodies stores to the campaign counter (`L03114`..`L03792` listing), and the campaign LOAD restores the saved raw word (FAME-022). So loaded corpses are neither skipped nor baselined by these bodies.

**Confidence.** High for the call order at the five sites, the pump's loop and exit, the resume projection loops and their guards, and the absence of a counter store in them. Medium for the counting consequence: it needs delivery to an empty client map, the hostile relation word and a saved counter that already includes the corpse. Unknown: whether the local transport returns a packet to the pump of the step that sent it, how many stage 5 actors stay on the actor list, native behaviour.

### FAME-029

Instrument: capstone decode; listings `q4_*`; `evidence/vtables.tsv`; `evidence/calls.tsv` (sites that push a message id).

Credits (SAV-970 screen). Activation `L03793` sets the scroll offset at `+94` to 0x1e0 (`L03794`). The draw routine `L03795` runs a step only when `timeGetTime` has advanced by 23 ms or more since its stored value (`L03796`); a step decrements the offset, and the first visible line index at `+98` becomes the absolute offset divided by the line height once the offset is negative (`L03797`..`L03798`). When `+98` reaches or passes the clamped last line index it sends 445 to itself through `L03799` (`L03800`, `L03801`). Slots 27 (`L03802`, one argument, then the base handler `R0717`) and 21 (`L03803`, three arguments) of the credits vtable also call `L03799` (`evidence/vtables.tsv`); they are the key-down and left-button-press slots (FAME-031), the same slots 27 and 21 in which the hall-of-fame class (slots `R0813`, `R0814`) carries its input handling. Message 445 reaches `R0716`, which with `modal+5c` nonzero posts 44c (SAV-971). The run time is a function of the line count of the credits text and the line height; both are now read (FAME-030).

Hall of fame (FAME-DISPLAY-012 screen). Routine `L03804` (slot 30 of the hall vtable at base `L03805`, `evidence/vtables.tsv`) builds a button with command 0x445 and stores the rectangle `0x230,0x1a0,0x25c,0x1c0` at `+a8`. Mouse-down inside the rectangle records a pressed state (`R0814`); mouse-up inside it calls `L03806`, which sends 445 (`R0815`, `L03807`, `L03808`). Key-down slot `R0813` calls only the base handler `R0717`, which acts only on Tab and the four arrows (FAME-031; the Enter and Escape clause is withdrawn). The screen's own methods (`L03804`..`L03809`) contain no `timeGetTime` call.

Terminal route. After the close chain SAV-971 describes, `R0816` calls campaign reset `R0757` (`L03810`), clears the drawable map array (`L03811`..`L03812`) and zeroes the score cache word (`L03726`).

Producers of 41d (`evidence/calls.tsv`, `push_41d`): `L03813`, `L03814`, `L03815` (dialog or panel constructors binding command 41d to a control), `L03816` (after a 446 modal close), `L03817` (frame code when the selected mission number is a multiple of 10), `L03818`, `L03819`.

SAVE. Writer `R0084` has two callers: `L03820` (SAV-972) and `L03821` in the arm for message 442. That arm copies text string 0x37 and the literal at `L03822` into `frame+134` and `frame+234` first (`L03823`..`L03824`), so message 442 saves from copied fixed text, with no test of `frame+3dc`; the content of that text was not read. Message 442 has one post site, the PostMessage at `L03825` after a mode-2 load handler; the frame child constructed at `L03826` carries 442 only as its widget id (FAME-032).

**Amended.** The base handler maps only Tab and the four arrows, not Enter and Escape, and the frame child with id 442 is not a second producer of 442 (`claims/retracted.md`). Slots 21 and 27 are the left-button press and key-down slots, and the credits duration is read in FAME-030. The scroll gate, end test, 445 sends, terminal clear and SAVE callers stand.

**Confidence.** High for the 23 ms scroll gate, the end test, the 445 sends, the vtable slot contents, the terminal clear and the two SAVE callers. Medium for the order credits then hall of fame then reset (SAV-971 scope). Unknown: which of the seven 41d producers is the Victory acknowledgment and base-class timers. The credits duration, the hall's keys and the 442 child are settled in FAME-030..FAME-032.

### FAME-030

Instrument: capstone decode of the matched EN/RU `rom.exe` (SHA-256 in `evidence/inputs.tsv`), listings in `evidence/listings.txt`, address assertions in `evidence/checks.tsv`; resource facts in `evidence/credits-facts.tsv` from a private reader of the `&YA1` archives.

Step. The draw routine `L03795` returns without drawing or stepping unless `timeGetTime` has advanced by 23 ms or more since the stored value (`L03796`, `L03827`). Each qualifying call decrements the scroll offset by one (`L03797`), so one step is one pixel. Activation `L03793` sets the offset to 0x1e0 (`L03794`). The image imports only `timeGetTime` from the multimedia library and has no timer-resolution request, so the draw-call cadence and timer granularity are Unknown.

Line pitch. Line y is the rectangle top plus the offset plus the line index times the line height (`L03828`). The height is the value `[[L03615]+4]` returned by slot 9 `R0817`. The object at `L03615` is the font1 loader object, whose path is `graphics\font1\font1` with the `.16` and `.dat` suffixes (`L03829`). The `.16` file in both installs has the same size, hash prefix and record count, and every record is 16 wide and 15 high (`evidence/credits-facts.tsv`). The line pitch is 15 pixels in EN and RU.

End. The end test `L03800` closes the screen through `L03799` when the first visible line index reaches the line count at `+7c`. The first visible index is the absolute offset divided by the height once the offset is negative. The line count is the number of CR-terminated segments of `main\text\credits.txt` read by `R0661`. A run is therefore 480 plus the line count times 15 steps. The EN file has 207 lines and the RU file 166; the first step needs no wait, so the minimum run is the step count less one times 23 ms: 82432 ms in EN and 68287 ms in RU. The text file is a resource of each install, not a claim of this ledger.

**Confidence.** High for the gate, the one-pixel step, the start offset, the end test, and the binding of the height read to slot 9 `R0817` (`evidence/checks.tsv`, `evidence/vtables.tsv`). Medium for the line height value, the line counts and the file size and hash, which come from a private reader of the shipped archives and are recorded by hand in `evidence/credits-facts.tsv`; the reader is not committed. Unknown: timer granularity, the cadence of draw calls, and so the wall time above the lower bound.

### FAME-031

Instrument: as FAME-030; `evidence/base-key-dispatch.tsv` prints the base handler's byte table and jump table.

Credits. The widget event slot `R0390` routes keys 0x100..0x102 to the focus child, then to the child list, then to the widget's own slots, and routes mouse messages to the captured child when one is set, otherwise by hit test. Credits slot 27 `L03802` is the key-down slot (`0x100` dispatches to `+6c`) and calls `L03799` for any key code, with no virtual-key test, before the base handler. Credits slot 21 `L03803` is the left-button-press slot (`0x201` dispatches to `+54`) and also calls `L03799`. The right-button, double-click and button-up slots are stubs that return zero. `L03799` sends 445 to the widget's own event slot directly.

Base key handler. `R0717` subtracts 9 from the key code and indexes a byte table (`L03830`) and a jump table (`L03831`). Five key codes have an action: Tab calls the focus cycle `R0787`, Left, Up, Right and Down send to the focus child's slots `+50`, `+48`, `+54` and `+4c`. Enter and Escape reach the default arm, which does nothing. The earlier reading that the base handler maps Enter and Escape is withdrawn (`claims/retracted.md`).

Hall of fame. The hall builds a button through `L03804` with command 445 and a zero-size rectangle, so it is never hit by the mouse. The button's key-down `R0713` acts only on Enter, when its parent is set and flag bit 1 is on; it posts 445 to the main window. The frame window procedure forwards 445 to the root, where the hall's modal event `R0716` closes it. The hall also ends on a left-button release inside the rectangle `0x230,0x1a0,0x25c,0x1c0` (`R0815`, `L03806`). Its char slot `L03832` returns zero. The hall's key slot, its button and the base handler do nothing on Escape; other root-level handlers were not enumerated, and Escape is tested at about 22 sites in `.text`, some of which act (for example `R0818` and `R0714`).

**Confidence.** High for the credits input slots, the base handler's table and the mouse rectangle (`evidence/listings.txt`, `evidence/base-key-dispatch.tsv`). Medium for the hall button's Enter path: the button key-down `R0713` is gated on `[this+0x30]` and slot `+0x20`, and its routing through `R0390` and `R0642` and the constructor argument order `R0700` are not in the listings. Medium that no other root child consumes Escape or another key during the hall: the map-view key handler `R0819` acts only when the frame state word is 1, and the hall and credits set other bits, but the full root child list during the hall was not enumerated.

### FAME-032

Instrument: as FAME-030; `evidence/calls.tsv` (`push_442`).

The frame constructor `R0315` creates the child at `+f8` with class `L03833` and id 442 (`L03826`). Its mouse-up `R0820` posts one of the main-menu commands (425, 426, 43b, 428, 418, 422, 429 or 10) by hotspot and disabled bits; the id 442 is its widget id and is never posted by it. The immediate 0x442 occurs only at the constructor push and at the PostMessage `L03825` after the mode-2 load handler. The terminal route adds the panel to the root at `L03834` after the root rebuild, and the hall entered from the menu hotspot leaves it in the root. While a modal widget is active it is the root's captured child (`R0776`), so a mouse message reaches the modal and not the panel. A key reaches the panel only when the modal returns zero; the hall's Enter is consumed first.

**Confidence.** High for the class, the id, the literal-push census of 0x442 and that mouse-up slot `R0820` posts none; the panel's other slots were not read, and computed ids are not excluded, so "never posts" is Medium. Medium for the capture rule `R0776` (not in the listings) and for the panel never receiving input during credits or the hall, because the sibling key order and the modal's return values were read, not run.

### FAME-033

Instrument: as FAME-030; `evidence/campaign-stores.tsv` lists the stores of the constructor, reset and step routines; `evidence/reset-fields.tsv` tabulates them by path.

Both routes call `R0757` on the record at `frame+548` (`L03835`, `L03810` terminal; `L03836` new game) and the reset itself ends with `R0755(10)`; the new-game arm's own later call `R0755(10)` finds the step already 10 and only rewrites `+118` with itself, a no-op. The reset reloads the scenario registry sizes, sets the 15-entry arrays of words and of dwords, clears the six dword vectors and the two record vectors, sets the auto-get mission to -1, zeros the step and loader fields, and finishes with `R0755(10)`. That step sets the step field and the last-loaded field to 10 and loads the Mission10 section: payment, shop price bounds, enable-mercenary, the document, mercenary, inn, shop and trade-centre lists, and the hero add field. The record constructor `R0677` sets a field at `+12c` to 10 that neither route writes, and a vector at `+ac` cleared by the reset that no loader store writes. The frame init at `L03837` also calls the reset. Differences between the two routes lie in frame state: see SAV-1159.

**Confidence.** High for the shared routine, the call sites and the call order (`evidence/checks.tsv`, `evidence/listings.txt`). Medium for the field-level lists, which are tabulated by hand in `evidence/reset-fields.tsv` and partly inferred; `evidence/campaign-stores.tsv` holds the decoded stores. The field meanings were not named beyond the registry keys the loader reads.

### FAME-034

Instrument: as FAME-030; `evidence/human-flag-stores.tsv` is a census of byte stores to a `+3f` displacement in decoded `.text`.

The announcer `R0811` loops the players and calls `R0675`, which tests the human flag first (`L03838`) and returns for a zero flag. Over the explicit byte stores to a `+0x3f` displacement in the linear `.text` sweep (`evidence/human-flag-stores.tsv`, six stores), the only nonzero store is `L03839` inside the participant entry routine `R0131`; the other five write zero. Indirect and computed writes are not excluded. The entry routine's calls between that store and its hero projection `L03739` include no actor-add path. The announcer emits nothing to the participant before entry and nothing before the hero projection once the flag is set. The projector `R0059` with recipient 0 also broadcasts to a player when the actor's owner equals that player (`L03840`), which is a separate route that does not pass the flag.

On the new-game path the dispatcher calls the client hero-create routine `R0821` at `L03841` and the start routine `R0099` at `L03842`; the start routine sends entry through `R0545` (`L03843`). `R0821` returns success only when the score cache at `[frame+d0]+3f54` is non-null (`L03844`) and fails at its 15 s timeout, so on that path a drawable is registered before entry. The candidate producer chain is: `R0821` sends packet 0x48 through `R0822` (`L03845`); the handler `L03846` calls `R0823` (`L03847`); that routine calls the projector `R0059` at `L03848`; with recipient 0 the projector reaches the player through the owner match at `L03840`, which does not test the human flag. The projection arguments at `L03848` were not read, so which drawable is first stays Unknown.

**Confidence.** High for the announcer's flag test, the flag store census and the entry order. Unknown: which drawable that chain registers first on the new-game path (the chain is a candidate, not a read-through). On resume the hero projection stays the first registration by the argument above (Medium).
