# Mission-load bounds

[Reference](format.md)

The reader records errors, while the mission caller gates on a pointer.
A null Altitudes pointer takes the caller's report-and-return-1 branch.
Missing Tiles sets reader status 5; with nonnull Altitudes the caller still
enters world build. No native malformed-map run establishes a box, menu
return or fault. — ALM-OUTCOME-323; ALM-REQ-055 (acceptance scope narrowed),
ALM-REQ-056 (amended)

## Counts and lengths

| Input | Operation in the selected game bodies | Bound of the conclusion |
|---|---|---|
| Header count/version | Unsigned count >=3, version <=1001; unsigned record loop | Header gates, not complete acceptance |
| Type0 W/H | Stored directly; type1 allocation/read requests low32(W * H * 2) | Overflow and geometric validation not found locally |
| Type2/type3 lengths | Type2 payloadSize drives both; type3's own size is not used for its request | A preceding type2 value is a dependency |
| Type4/5/6/8 metadata counts | Signed loops | Zero/high-bit values skip iterations; surrounding work can still execute |
| Type7 three counts | Signed loops; third count also reaches resize and count *184 allocation arithmetic | Growth increment 4..1024 is not a maximum count |
| Type8 element/type9 outer and element counts | Signed loops | Raw u32 width is not a 2^32-1 execution limit |
| Present type5 with nonpositive count | Post-loop access to roster slot0 still occurs | Different from an absent type5 default |

These local branches are established; missing-check findings are Medium.
The selected bodies are the whole reader, header helpers, trigger resize and
listed accessors; zero section scans were run. Stream, allocation and growth
callees can add behavior. — ALM-LOAD-320, ALM-META-024 (amended);
the execution-limit clauses of ALM-LIM-067 are partially retracted.

The selected grid ingest narrows source and destination indices to low16,
uses full signed dimensions for its loop comparisons and forms edge bytes
from low8(W-9), low8(H-9). These operations do not establish a safe maximum
dimension or termination beyond the narrowed counters. — ALM-GRID-325

## Placement decisions

| Input/path | Local decision | Remaining boundary |
|---|---|---|
| Type6 wire +0x18 link | Decrement modulo 2^32; reject index -1 or signed index >= signed loot count | Memory layout and actual fault are Unknown |
| Type8 element +0x06 link | Zero means none; otherwise decrement and use backing +4*index modulo 2^32 | No count check found in selected reader/consumer/accessor bodies |
| Building kind | Low8 kind zero or above count-1 skips definition setup | Constructor/walker still have continuation paths; later validity Unknown |
| Structure owner miss | Report and continue registration | Completed world/display outcome depends on callees |
| Unit owner miss | Report and skip before construction | Later mission completion is not proven |
| Unit row lookup miss | Returns 0; walker forwards that row | Row0 contents and constructors remain prerequisites |
| Unit actor-id zero or placement-result zero | Report, delete and skip | Message receiver and admission predicate remain unread |
| Unit/structure/ground-loot coordinates | Shift by 8 where specified, then low-byte coordinates | Truncation is not a geometric clamp |
| Unit group | Compare dword equality; request group on no match | Allocation and arbitrary group limits remain unresolved |
| Item class0/15 | Factory returns null; consumer skips the element | Other elements continue locally |
| Nonempty stock with missing actor | Local path dereferences the lookup result | Native fault not observed; lookup helper remains a boundary |
| Type9 missing actor/building target | Report and skip | Later load behavior remains conditional |
| Type9 cell arm | Reads tail entries0/1 without a local count test | Short-tail consequences Unknown |

— ALM-LINK-321, ALM-LINK-322, ALM-STRUCT-324, UNIT-LOAD-159,
ITEM-LOAD-180. The old ALM-UNIT-040 link predicate is partially retracted.
ALM-127, ALM-CLS-036, ALM-OWN-039, ALM-EFFLINK-072 and ITEM-OWNED-028
retain their other clauses with these local outcome boundaries.

For a positive loot count n, the unit-link gate admits exactly
[1,n] and [0x80000001,0xffffffff]. Wire 0xffffffff addresses backing-8;
0x80000001 aliases slot0 after multiplication wraps. Admitted high-bit values
therefore do not all imply an out-of-array address. — ALM-LINK-321

## Reporting and unresolved references

The message sender copies text and submits message0x471 to a window.
The load reporter selects that sender or a separate unread logger from
session +0x0c. Program text alone does not establish a dialog. The message
handler, upstream return-to-menu behavior and actual faults remain Unknown.
The helper Unknowns in HERO-163 and TRIG-HEROFAIL-078 are narrowed to those
receivers. Reading the item factory's final return closes only that epilogue
Unknown in ITEM-MAGROW-177. — ALM-OUTCOME-323, ITEM-LOAD-180

Arbitrary numeric text/script reference bounds remain Unknown. The reader's
type7 counts do not establish reference semantics; the AI builder and script
binder bodies were not selected. — ALM-SCRIPT-327

## Installed ranges and margins

The population is all loose ALM files directly under each EN/RU root and all
ALM members of scenario.res, without deduplication. All 38 EN and 34 RU maps
have version990. These are shipped-file measurements, not editor-session
or native-acceptance observations. — ALM-CORP-326

| Value | EN min..max | RU min..max | Applicable comparison |
|---|---:|---:|---|
| W and H | 72..256 | 72..256 | Local multiplication and index narrowing; no proven safe dimension cap |
| Structure kind | 1..66 | 1..66 | Low8 versus runtime count-1; table capacity not reacquired |
| Structure owner | 1..8 | 1..8 | All within each map's roster; count-minus-owner 0..7 |
| Structure x/y cells | 8..246 / 8..246 | 8..245 / 8..246 | Minimum distance to geometric upper edge 9/9 EN, 10/9 RU |
| Unit class/secondary | 1..80 / 0..69 | 1..80 / 0..69 | Signed class band and low-byte lookup differ |
| Unit definition | 0..1086 | 0..1085 | Lookup key; numerical maximum is not a row-count bound |
| Unit owner/group | 1..9 / 0..436 | 1..9 / 0..436 | Owner count difference 0..7; no group ceiling established |
| Unit x/y cells | 8..246 / 8..247 | 8..246 / 8..247 | Minimum distance to geometric upper edge 9/6 |
| Unit loot link | 0..22 | 0..22 | Per-map declared-count minus wire link 0..23; zero denotes none |
| Loot element count | 0..7 | 0..7 | Signed loop; shipped maximum is not a native cap |
| Raw item code | 271..64635 | 271..64635 | Low16 packed class/material/shape/row |
| Item class/row | 1..14 / 1..49 | 1..14 / 1..49 | Class width4; row width5 except class14 width8 |
| Item caster link | 0..37 | 0..37 | Decoded-count minus link 0..122 EN, 0..38 RU; zero denotes none |
| Script/text references | Unknown | Unknown | Unavailable in the preregistered corpus grammar |

No selected framing/truncation counter is nonzero. RU has one map without
types4/5/6/8/9. Zero outside-rectangle/roster counters describe this population
only. The corpus output label type9_target_word measures raw u32 at +0x0e,
combining two wire words; it does not establish a semantic target range.
— ALM-CORP-326
