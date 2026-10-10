# VIDEO — music, sound effects and cutscenes

Level 2 public functional ledger. This edition preserves the functional
conclusion, confidence, status and private evidence identity while omitting
instruction listings, executable-address inventories and reconstructable
shipped-content tables. The private research snapshot named by
`SOURCE.md` retains the complete evidence. Format of this file:
[registry.md](registry.md). IDs are permanent.

## Music

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-MUSIC-001 | The two preserved ROM1 roots carry the same 110,632,356-byte music archive, exposing exactly 21 file nodes whose matching payloads are byte-identical 21/21. | Medium | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |
| VIDEO-MUSIC-002 | Every examined shipped music member is a complete ordinary PCM RIFF/WAVE with two channels, 22,050 Hz, 16-bit samples and no observed extra loop/metadata chunks. | Medium | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |
| VIDEO-MUSIC-003 | Shipped track duration follows PCM frame count and reproduces as `(payloadSize − 44) / 4 / 22050`; the shipped durations span 12.982404 s to 88.693016 s. | Medium | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |
| VIDEO-MUSIC-004 | Music reachability depends on the ordinary resource resolver and process/environment layout; presence of an archive somewhere in the install tree alone does not prove it is mounted. | High / Medium | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |
| VIDEO-MUSIC-005 | The recovered music requests use fixed candidate names that join to the shipped music namespace; the recovered request population accounts for the observed candidate families. | High / Medium | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |
| VIDEO-MUSIC-006 | Recovered music requests replace candidate **lists**, not isolated one-shot tracks. | High / Medium | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |
| VIDEO-MUSIC-007 | With an initialized player buffer, ordinary mode chooses a randomized candidate order and advances through it at stream end; fixed-source mode retains/reloads the selected candidate. | High | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |
| VIDEO-MUSIC-008 | The researched school-state bit changes the order of the two school candidates, not which of the two is permitted; ordinary playback still uses the common randomized/list progression rules. | High | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |
| VIDEO-MUSIC-009 | Text markup includes a numeric music-candidate selector that enters persistent fixed-source mode. | High / Medium | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |
| VIDEO-MUSIC-010 | Request availability, playback enable and music volume are distinct controls. | High / Unknown | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |
| VIDEO-MUSIC-011 | Music uses a streaming audio buffer. | High | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |
| VIDEO-MUSIC-012 | Archive capacity and request reachability are separate: adding a valid music payload does not by itself create a request for it. | High / Medium | ✔ promoted | [EXP-0229](../experiments/EXP-0229-music-runtime/) |

### VIDEO-MUSIC-001

- The archive's SHA-256 is
  `61b7fcd4c4515525725fa6bd45ab4bd7b84453a6a36d36639404ba10fc6dbe86` on both
  roots.
- The 21 matching payloads are byte-identical, not only name- or
  metadata-equal.
- The exact member-name inventory is corpus evidence rather than part of the
  public format grammar.

**Confidence.** Medium — an exhaustive two-root corpus result; agreement alone
does not discriminate against a different lawful release.

### VIDEO-MUSIC-002

**Confidence.** Medium — exhaustive for the examined shipped population, not a
universal WAVE restriction.

### VIDEO-MUSIC-003

No separate authored loop point was found in the examined WAVE payloads: the
absence of `smpl` and of every other non-`fmt `/`data` chunk means no shipped
WAVE supplies one.

**Confidence.** Medium — exhaustive over the shipped corpus, with the
arithmetic fixed by each payload's own header.

### VIDEO-MUSIC-004

**Confidence.** High for the resolver condition / Medium for the
preserved-layout observation.

### VIDEO-MUSIC-005

Exact literal/reference tables are private evidence.

**Confidence.** High for the recovered literal/reference population / Medium
for UI-surface interpretation.

### VIDEO-MUSIC-006

One family builds the larger mission/background list; the others build smaller
context lists. Computed indirect request paths remain possible.

**Confidence.** High within the recovered direct-call population / Medium for
named-surface interpretation.

### VIDEO-MUSIC-007

A one-entry ordinary list therefore repeats.

**Confidence.** High — competing ordinary/fixed-source models are
distinguished by the reached state machine.

### VIDEO-MUSIC-008

**Confidence.** High.

### VIDEO-MUSIC-009

No use was found in the recovered pager corpus; additional unrecovered content
remains possible.

**Confidence.** High for the parser-to-player mechanism / Medium for bounded
shipped absence.

### VIDEO-MUSIC-010

Music volume is independent of SFX and speech volume; one initializer path
remains Unknown.

**Confidence.** High for the distinct controls and ordering / Unknown for the
unresolved initializer.

### VIDEO-MUSIC-011

Stop/reload behavior depends on ordinary versus fixed-source state, and
replacement/destruction pass through that state-aware cleanup path. No stored
absolute pointer to that conditional routine exists in the parsed sections.

**Confidence.** High within the direct-call scope — both state arms, the
import, the call arguments, the refill worker and the destructor are recorded;
a computed indirect caller remains possible.

### VIDEO-MUSIC-012

Recovered live lists are fixed-name lists; runtime-computed names and targets
remain open.

**Confidence.** High for the recovered list-building and resolver mechanisms /
Medium for the negative population boundary, which does not exclude
runtime-computed names, targets or another representation.

## Sound effects

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-SFX-013 | The common SFX playback boundary receives an already selected sample plus volume/attenuation, pan, play/loop state, priority/category and optional frequency. | High | ✔ promoted | [EXP-0230](../experiments/EXP-0230-non-music-sfx-events/) |
| VIDEO-SFX-014 | The static research found a large bounded population of direct callers to the common SFX terminal, with some receiver sources classified and others unresolved. | High / Medium | ✔ promoted | [EXP-0230](../experiments/EXP-0230-non-music-sfx-events/) |
| VIDEO-SFX-015 | Registry capacity is data-sized while event reachability is supplied by compiled selectors, inherited class values, formulas or direct filenames. | High / Medium | ✔ promoted | [EXP-0230](../experiments/EXP-0230-non-music-sfx-events/) |
| VIDEO-SFX-016 | Recovered interface actions select fixed SFX registry slots before entering the common playback boundary. | High / Unknown | ✔ promoted | [EXP-0230](../experiments/EXP-0230-non-music-sfx-events/) |
| VIDEO-SFX-017 | Ambient sound selection is driven by visible terrain/object state plus a scheduled deadline branch. | High / Medium | ✔ promoted | [EXP-0230](../experiments/EXP-0230-non-music-sfx-events/) |
| VIDEO-SFX-018 | Unit swing, cast and hurt sounds do not share one universal five-element source: different actions use inherited class entries, formulas and a voice-bank route with throttling/gates. | High / Medium | ✔ promoted | [EXP-0230](../experiments/EXP-0230-non-music-sfx-events/) |
| VIDEO-SFX-019 | Spell SFX selection combines formula-driven choices with a delayed fixed projectile choice; not every spell is proved to execute every candidate arm. | High / Medium / Unknown | ✔ promoted | [EXP-0230](../experiments/EXP-0230-non-music-sfx-events/) |
| VIDEO-SFX-020 | The two preserved roots agree on the sparse SFX registry and on payloads reached by the recovered registry join, while the complete SFX archives are not byte-identical. | Medium | ✔ promoted | [EXP-0230](../experiments/EXP-0230-non-music-sfx-events/) |
| VIDEO-SFX-021 | A separate direct-filename SFX route exists, so a registry-only sound model is false. | High / Medium | ✔ promoted | [EXP-0230](../experiments/EXP-0230-non-music-sfx-events/) |

### VIDEO-SFX-013

Slot selection occurs upstream.

**Confidence.** High — the alternative model in which a playback argument is
itself the registry slot is excluded.

### VIDEO-SFX-014

This is not claimed as a universal runtime census: fully computed targets,
aliases derived after a bulk copy, and registry-derived receivers passed
through members or parameters all remain possible.

**Confidence.** High for the measured static populations / Medium for their
union as a runtime bound; the unresolved-receiver complement is explicitly
unresolved, not a demonstrated non-registry source.

### VIDEO-SFX-015

Adding a registry row can change an existing selector but does not itself
create a new event.

**Confidence.** High for loader/selector forms / Medium outside recovered
forms.

### VIDEO-SFX-016

Exact shipped slot-to-resource names are corpus/content evidence and are
omitted from this public edition; one selector family's common gate, update
ordering and visible transition labels remain Unknown, because only its
selector tail was recovered.

**Confidence.** High for each recovered selector and for the gates and
orderings stated outside that family / Unknown for that family's common gate,
update ordering and physical labels; none of it is inferred from the registry
path.

### VIDEO-SFX-017

One potential branch is unreachable in the examined shipped object corpus but
is not structurally impossible.

**Confidence.** High for control flow/selectors/timing / Medium for shipped
non-reachability.

### VIDEO-SFX-018

**Confidence.** High for the reached selector/gate logic / Medium for the
exhaustive shipped class-table join.

### VIDEO-SFX-019

Ordering of several retained client tails remains Unknown.

**Confidence.** High for the recovered selectors/gates / Medium for corpus
joins / Unknown for the stated ordering gaps.

### VIDEO-SFX-020

Exact missing-slot and filename inventories are private corpus evidence.

**Confidence.** Medium — exhaustive for two preserved corpora only.

### VIDEO-SFX-021

Filename spelling alone does not identify which producer route owns an event,
and computed names remain open.

**Confidence.** High for the recovered literal/reference population / Medium
for bounded corpus join.

## Cutscenes

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-029 | Startup selects one of two cutscene namespaces from startup/environment media-speed state. | High | ✔ promoted | [EXP-0264](../experiments/EXP-0264-cutscene-smacker-abi/) |
| VIDEO-030 | When the cutscene gate is active, a mission request probes numbered `.smk` members in increasing two-digit order within the selected namespace. | High / Medium | ✔ promoted | [EXP-0264](../experiments/EXP-0264-cutscene-smacker-abi/) |
| VIDEO-031 | The sidecar filename is derived by replacing the movie's **last** extension with `.reg`, preserving earlier dots. | High / Medium / Unknown | ✔ promoted (amended, partially retracted) | [EXP-0264](../experiments/EXP-0264-cutscene-smacker-abi/) |
| VIDEO-032 | Cutscene sidecars provide initial position plus frame-bounded fade and pan records. | High / Medium | ✔ promoted | [EXP-0264](../experiments/EXP-0264-cutscene-smacker-abi/) |
| VIDEO-033 | ROM1 links a bundled Smacker decoder through a fixed imported API surface. | High / Unknown | ✔ promoted | [EXP-0264](../experiments/EXP-0264-cutscene-smacker-abi/) |
| VIDEO-034 | Game-side cutscene setup selects the movie resource, opens decoder state, obtains dimensions, configures a destination/blitter path and enables decoder sound after successful setup. | High / Unknown | ✔ promoted | [EXP-0264](../experiments/EXP-0264-cutscene-smacker-abi/) |
| VIDEO-035 | A game-side frame step waits for readiness/focus, updates sidecar fade/pan state, decodes/presents the frame, checks termination state and advances. | High / Unknown | ✔ promoted | [EXP-0264](../experiments/EXP-0264-cutscene-smacker-abi/) |
| VIDEO-036 | Ordinary key/system-key down, mouse-button down, close and quit messages stop the current numbered cutscene scan. | High / Unknown | ✔ promoted | [EXP-0264](../experiments/EXP-0264-cutscene-smacker-abi/) |

### VIDEO-029

Both archive families may be attempted, but the numbered mission route uses one
selected namespace and has no general fallback to the other.

**Confidence.** High for the recovered branch logic / runtime mount success
outside scope.

### VIDEO-030

Missing/open-failure continuation and user-stop termination are distinct; the
gate producer and visible failure transitions remain Unknown.

**Confidence.** High for gate/format/bounds/branch behavior / Medium for
examined archive population.

### VIDEO-031

The sidecar is loaded before the original movie resource is opened. Exact
shipped sidecar hashes/values are private evidence.

**Confidence.** High for filename derivation and load order / Medium for corpus
equality / Unknown for missing-sidecar presentation.

**Amended.** The former first-dot reading is refuted: EXP-0264's correction in
[`retracted.md`](retracted.md) replaces it with the last-dot rule stated
above. The sidecar-before-open order and the corpus equality stand.

### VIDEO-032

Fade state affects palette presentation; pan state changes presentation
coordinates across the reached frame interval.

**Confidence.** High for reader-to-consumer data flow / Medium for two-root
sidecar census.

### VIDEO-033

The public edition records only this game-side dependency boundary;
export-body/internal decoder details require separate third-party review.

**Confidence.** High within the parsed import/export boundary / Unknown for
computed dynamic invocation.

### VIDEO-034

Internal decoder structure offsets are intentionally omitted.

**Confidence.** High for the reached ROM1-side call order / decoder-internal
representation remains Unknown; decoder internals are also outside the public
scope.

### VIDEO-035

Exact callback meaning, pixels, audio latency and physical cadence remain
Unknown.

**Confidence.** High for recovered ROM1-side ordering / Unknown for physical
presentation and opaque decoder/host semantics.

### VIDEO-036

Natural completion and some construction/open failures follow the continuation
path.

**Confidence.** High for message/return branches / Unknown for physical error
presentation.

## Cutscene decoder boundary

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-045 | The original decoder integration can consume a game-provided resource source rather than requiring a standalone filename. | High | ✔ promoted (amended, partially retracted) | [EXP-0266](../experiments/EXP-0266-smacker-buffer-contract/) |
| VIDEO-046 | For the borrowed-resource route, the source's current origin determines where movie decoding begins. | High / Unknown | ✔ promoted | [EXP-0266](../experiments/EXP-0266-smacker-buffer-contract/) |
| VIDEO-047 | Decoder setup configures a borrowed game destination and frame decode writes into that destination. | High / Unknown | ✔ promoted | [EXP-0266](../experiments/EXP-0266-smacker-buffer-contract/) |
| VIDEO-048 | The researched decoder supports indexed palette output and additional packed-color modes; ROM1's reached caller uses indexed output. | High / Medium / Unknown | ✔ promoted | [EXP-0266](../experiments/EXP-0266-smacker-buffer-contract/) |
| VIDEO-049 | Frame decode and frame advance are separate operations. | High / Unknown | ✔ promoted | [EXP-0266](../experiments/EXP-0266-smacker-buffer-contract/) |
| VIDEO-050 | Decoder readiness/wait behavior has clock- and sound-progress-dependent paths rather than being a single fixed sleep. | High / Unknown | ✔ promoted | [EXP-0266](../experiments/EXP-0266-smacker-buffer-contract/) |
| VIDEO-051 | Decoder close respects ownership of a borrowed game resource source and configured destination; separate game/buffer cleanup owns other resources. | High / Unknown | ✔ promoted | [EXP-0266](../experiments/EXP-0266-smacker-buffer-contract/) |

### VIDEO-045

The public claim is limited to that ownership/input distinction;
bundled-decoder ABI flag values and instruction evidence are private. These
branches do not establish a general malformed-input contract.

**Confidence.** High — the recovered callee body, restored import-name map and
read-helper slice distinguish text, handle and caller-state alternatives, and
native nonzero-offset handle input matches filename input.

**Amended.** The allocation-selector clause is refuted: EXP-0266's correction
in [`retracted.md`](retracted.md) separates the two selectors that clause had
joined. ROM1's borrowed-handle route stands.

### VIDEO-046

Decoder buffering policy is distinct from a compressed-payload length.

**Confidence.** High for the bounded discriminator / I/O-failure and broader
ownership paths Unknown.

### VIDEO-047

ROM1 uses the ordinary indexed-output path; exact decoder descriptor
offsets/mode flags are intentionally omitted.

**Confidence.** High for reached setup/decode behavior / unsupported output
modes and extents Unknown.

### VIDEO-048

Palette changes are exposed to the game-side presentation path. Exact
decoder-internal tables/flags are private third-party evidence.

**Confidence.** High for the reached mode distinction, the retained layout
measurements and the bounded differential sample / Medium for EN/RU sample
parity / Unknown for the isolated third flag value, for all palette-change
sequences and for other video populations.

### VIDEO-049

Decoder-local wrapping behavior does not itself define ROM1's cutscene end
condition; the game-side caller owns the stopping bound.

**Confidence.** High for the retained increment, wrap and latch branches and
for the observed no-advance results / Unknown for native terminal/ring-frame
behaviour and for complete drop-helper semantics.

### VIDEO-050

The public claim does not expose decoder state offsets or callback internals.
The wait helper's own latency is outside the retained proof, so no nonblocking
or wall-clock playback guarantee follows.

**Confidence.** High for the bounded clock arithmetic, branches and return
roles / Unknown for audible timing, wraparound behaviour, sound-callback
semantics and observed ROM1 cadence; no native wait call was made.

### VIDEO-051

Exact allocator flags, structure sizes and internal free lists are private
third-party evidence. In the bounded native sample the borrowed source remains
seekable and the destination bytes remain readable and unchanged after decoder
close.

**Confidence.** High for the bounded ownership and survival checks, which the
native survival results corroborate on the named sample / Unknown for decoding
after destination release, for concurrent calls and for all error-path
lifetimes.

## Option controls and consumers

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-SFX-053 | The recovered Sound Options Acknowledgments checkbox imports the stored Acknowledgement value. | High / Medium | ✔ promoted | [EXP-0364](../experiments/EXP-0364-acknowledgment/) |
| VIDEO-SFX-054 | In the recovered selected-unit response path, a zero Acknowledgement value rejects the response chooser before candidate collection; nonzero values pass that gate. | High / Medium / Unknown | ✔ promoted | [EXP-0364](../experiments/EXP-0364-acknowledgment/) |
| VIDEO-MUSIC-056 | With a playback buffer present, the recovered Random Order setter selects ordinary mode and initializes the candidate permutation0..n-1. | High / Medium / Unknown | ✔ promoted | [EXP-0366](../experiments/EXP-0366-option-consumers/), derived instructions, controlled measurements and two branch mutations |
| VIDEO-OPTIONS-057 | The recovered Sound Options list walks the player current candidate bank and looks up normalized names after their six-character resource prefix in the dictionary loaded from main/text/tunes.txt. | High / Medium / Unknown | ✔ promoted | [EXP-0366](../experiments/EXP-0366-option-consumers/), derived instructions, controlled measurements and two branch mutations |

### VIDEO-SFX-053

A separate dialog message exports the selected control value to that cell
before forwarding the close message. Changing the selected control value alone
does not update the cell in the isolated probe; the close message alone does
not export it. This does not establish every control transaction or native
registry persistence.

**Confidence.** High for the joined copy methods and discriminated message
arms; Medium for their original dialog interpretation.

### VIDEO-SFX-054

The tested caller sends its command before choosing a response, so disabling
the option leaves that command-service event intact. The concrete unit response
method separately refuses elapsed time below 3000 milliseconds since its
preceding response timestamp. Conditional original-instruction execution
discriminates zero/nonzero, the timing boundary and two mutations; it does not
establish all voice-request families or audible playback.

**Confidence.** High for the reached instruction conditions and request
boundary; Medium for the selected original command-path interpretation.
Dialogue, damage/death audio and unexamined callers remain Unknown.

### VIDEO-MUSIC-056

Disabled retains that sequential order. Enabled performs n swaps, each using
two CRT rand()%n indices. It does not select fixed-source mode. An absent buffer
leaves the measured mode, flag and order unchanged. The Sound Options checkbox
writes its configuration value and invokes this setter immediately.

**Confidence.** High for the original setter, control arm and controlled draws;
Medium for dialog interpretation. Native RNG seed, distribution, complete
context activation and listening remain Unknown.

### VIDEO-OPTIONS-057

Selecting a list row only changes the selected index. Play enables music,
reloads only when the selected index differs, then applies gain and starts.
Stop disables music and invokes the state-specific stop or transition method.
These controls operate separately from channel volume.

**Confidence.** High for index walk and measured selection/Play/Stop event
arms; Medium for the title dictionary and dialog interpretation. CString
normalization is a named substitute; Unicode, missing-title behavior, native
list population, fade audibility and registry lifetime remain Unknown.

## Character-generator sounds

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-SFX-058 | On the pre-create page each left press restarts a difficulty button's `level1..3.wav` or a hero button's `char.wav`; Back, Escape, and OK or Enter with a non-empty name request `ok.wav` just before the page's close stops it. | High / Unknown | ✔ promoted | [EXP-0407](../experiments/EXP-0407-chargen-sounds/) |
| VIDEO-SFX-059 | On the detailed page each applied statistic step restarts `+_-.wav`, also on double-click and held-button repeat; a skill press requests its class's member for that skill slot unless it plays, as the school room does. | High / Medium / Unknown | ✔ promoted (amended) | [EXP-0407](../experiments/EXP-0407-chargen-sounds/) |
| VIDEO-SFX-060 | Every `chrgen` request found passes SFX volume, pan 0, no loop and priority 128; at most one instance of each member plays, and members share 16 channels, displacing only a lower-priority sound when all are busy. | High / Medium / Unknown | ✔ promoted | [EXP-0407](../experiments/EXP-0407-chargen-sounds/) |

### VIDEO-SFX-058

- A press on difficulty button 0, 1 or 2 (`UNIT-GATE-014`) stops a playing
  `level1.wav`, `level2.wav` or `level3.wav`, rewinds it and requests it again,
  whether or not that button was already selected.
- A press on any of the four hero buttons (`SESS-HERO-013`) does the same with
  `char.wav`. A comparison with the current hero guards only the name-field
  update; the request follows on both branches.
- OK and Enter with a non-empty name (`TEXT-CHARGEN-028`), and the amulet Back
  control and Escape (`HERO-CHARGEN-083`), each request their own `ok.wav`
  sample unless it plays and then send the page its continue or Back message.
  With an empty name, OK and Enter request nothing and send nothing.
  The open page closes inside that call: the close stops the new instance and
  deletes the page's samples before the press or key routine returns. Escape
  reaches the page through the frame's key forwarding after the
  character-generation transition (`MENU-ESC-010`, `TOWN-373`).
- Motion, release, the right button and a double-click's second click request
  no `chrgen` member, and the page's open routine makes no playback call.
  Typing the name requests non-`chrgen` letter sounds.

**Confidence.** High for the controls, events, arguments and replay rules:
every arm is read with its switch table, and the only selection comparison
guards other work. High that the close stops `ok.wav` within the same call on a
page whose open ran; the open sets the page's open flag on every path.

**Unknown.** Whether any of that `ok.wav` instance is audible. A press reaching
the page while its open flag is clear was not searched for.

### VIDEO-SFX-059

- An increase applies only when its cost fits the free points and the value is
  below 45; a decrease only when the value is above 15. Each applied step stops
  and rewinds a playing `+_-.wav` and requests it; a refused step is silent.
- The statistic panel treats a double-click's second click and the held-button
  repeat as presses. While the left button stays down, the first cursor tick
  more than 150 ms after the last mouse message posts the repeat, and each
  later tick more than 66 ms after the tick that posted the previous repeat
  posts it again; any mouse message restores the 150 ms delay. The other
  panels and pages ignore both.
- A skill press selects index 0..4 and requests the member loaded at that index
  unless it plays: `fsword`, `faxe`, `fclub`, `fpike`, `fbow`, or with the
  hero's class bit set `mfire`, `mwater`, `mair`, `mearth`, `mastral`. It stores
  index + 1 as the draft skill (`SAV-934`); no comparison with the previous
  selection guards the request.
- The school room requests the same ten members on a left press in a skill
  column unless playing, also when the press deselects the column. Its member
  follows the stored skill slot (`TOWN-GENERAL-106`) as on the detailed page:
  1 sword or fire, 2 axe or water, 3 club or air, 4 pike or earth, 5 bow or
  astral.
- Motion and release request no `chrgen` member on the page, its panels or the
  school room.

**Confidence.** High for the step conditions, event routing, index-to-member
and slot-to-member mapping and replay rules. Medium for the on-screen position
of each skill index and school column: the hit masks were not rendered.

**Unknown.** The cursor tick rate and the system double-click interval, which
set the repeat cadence and decide when a second click counts as a
double-click. The meaning of the school room field that blocks its press when
non-zero.

**Amended.** `MENU-143` settles whose interval it is: the operating system's
double-click time and rectangle, with no game-set value; the value on a given
machine stays Unknown. `MENU-146` adds that a held button after a double click
posts no repeat until the next press.

### VIDEO-SFX-060

- The request goes through the common SFX boundary (`VIDEO-SFX-013`). It loads
  the sample on first use, takes the first of its duplicate buffers that is not
  playing and needs a channel: the first idle one of 16, or, when all 16 play,
  the one holding the lowest priority below 128, which is stopped. Otherwise
  the request is dropped, so a `chrgen` request never displaces another at 128.
- Each `chrgen` caller either restarts or skips a playing instance of its own
  member, so at most one instance of each member plays; different members
  overlap.
- The SFX volume option only sets the volume value; no `chrgen` request tests
  it or an enable flag. A static initializer sets it to −700 at process start;
  a settings load replaces it with the registry value `SoundSfxPos` when
  `HKLM\SOFTWARE\1C\Allods` opens; the option's control stores −(p − m)²/m,
  truncated, for control position p and control value m.
- `ok.wav` is also requested by a press on any of the main menu's eight buttons
  and on the Hall of Fame OK, each skipping a playing instance.
- The searched population is every reference to the 16 literals and every
  playback call in the code of the five owning surfaces; 80 playback calls
  elsewhere were not attributed.

**Confidence.** High for the arguments, the buffer and channel rules and the
channel count, set once from a constant 16. High for the static initializer
and the registry load of the volume value, each read whole. Medium that no
other code requests these members: a sample pointer copied out of an unswept
field is not excluded.

**Unknown.** The runtime volume value: whether the registry holds
`SoundSfxPos`, and any other write through the settings object's address.
Mixing, device behaviour and audibility.

## Music at mission loss and campaign LOAD

Terms. The player is the object at `frame+0xc8`. The music gate is `L12817`, the same gate
`VIDEO-MUSIC-006` reads. The stop is `R1233`, the replace is `R2050` and the start is
`R2086`. The mask is `frame+0x3dc`. Window messages are posted through one import slot,
`[L03535]`, so every hop below is queued, not nested.

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-MUSIC-061 | No arm read on the mission-lost panel's display path and no direct call from it reaches a music routine (126, 162 and 62 computed-call functions unread); its one sound call is a fixed SFX. | High / Medium | ● active | [EXP-0426](../experiments/EXP-0426-loss-load-music/) |
| VIDEO-MUSIC-062 | Exit to Main Menu stops the player through the teardown, then requests the one-entry menu list from the `0x421` arm; both happen after the choice, not at display. | High / Medium | ● active | [EXP-0426](../experiments/EXP-0426-loss-load-music/) |
| VIDEO-MUSIC-063 | In the arms read, Load Game from the failure panel makes no music call when chosen or when the save dialog opens; the stop comes at selection, and cancelling takes the Exit route. | High / Medium | ● active (amended) | [EXP-0426](../experiments/EXP-0426-loss-load-music/) |
| VIDEO-MUSIC-064 | LOAD of a mission save stops the player, then, if `R0099` reaches `L12818` with the gate set, requests the twelve-entry `B00`..`B11` list and starts it. | High / Medium | ● active | [EXP-0426](../experiments/EXP-0426-loss-load-music/) |
| VIDEO-MUSIC-065 | LOAD of a town save makes no music request in the load routine; it posts `0x42e`, whose arm requests the one-entry Town list. | High / Medium | ● active | [EXP-0426](../experiments/EXP-0426-loss-load-music/) |
| VIDEO-MUSIC-066 | The menu and Town owners skip the replace when the player's list already starts with their entry; the `B00`..`B11` owner never compares, so a mission LOAD that reaches it re-requests. | High / Medium | ● active | [EXP-0426](../experiments/EXP-0426-loss-load-music/) |

### VIDEO-MUSIC-061

- Order, EN and RU executables byte-identical: the client `0xb4` arm at `L06592`
  posts `0x433` with wParam `0xff`; the `0x433` arm at `L03537` stores
  `frame+0x414 = 0xff` and posts `0x431`; the `0x431` arm at `L12819`, taken only while
  latch `frame+0x3c0` is zero, sets the latch, constructs `L03302`, stores the panel at
  `frame+0x110`, shows it with `R0361` and then calls the SFX terminal `R0386`.
- `R0361` sets mask bit 8 (bit `0x8000` for a panel stored at `frame+0x3a8`), remaps
  the screen and calls virtual slots of the panel. It calls no music routine.
- The SFX call carries a fixed registry slot, 16, in `EXP-0230` (event `mission-failed`,
  `VIDEO-SFX-013`, `VIDEO-SFX-014`). It does not call the stop, the replace or the start.
- Direct-call closures from the reporter `R0132` (686 functions), the client
  dispatcher `R0509` (574), the panel constructor (59), the show routine (94) and the
  panel's vtable slots `+0x48`, `+0x78`, `+0x80`, `+0x84` reach none of the twenty-four
  music routines marked in `evidence/closures.tsv`.
- The stop, the replace, the start and fourteen owner or teardown entries have no stored
  address in any section of either executable and no `E9` transfer; each has only direct
  `E8` references (`evidence/scan-en.tsv`, `evidence/scan-ru.tsv`, identical).
- Bound: no direct-call path from the reporter, the client dispatcher, the panel constructor, the show routine or the panel's vtable slots reaches a music routine, and none of the arms read calls one. The computed-call functions unread are 126 in the reporter closure, 162 in the `0xb4` handler closure and 62 in the SFX terminal closure.
- The player keeps its previous list and state at display. Where a track ends during the
  panel, `VIDEO-MUSIC-007`'s end-of-stream advance applies; it is not a loss request.

**Confidence.** High for the arm-by-arm reads and the three `E8` posts. Medium that no
request exists between the reporter and the panel: the closures hold 126 functions with
a computed call that were not read, and the first `0xb4` hop through the session queue is
inherited from `MISSION-DEFEAT-046`.

**Unknown.** Whether the stream worker keeps running while the panel is up. Which sample
SFX slot 16 plays and whether it is audible over the music.

### VIDEO-MUSIC-062

- The panel result `0x445` makes the `+0x110` close arm post `0x41e`. The `0x41e` arm at
  `L07930` calls teardown `R1283`, then posts `0x421` (`MISSION-DEFEAT-046`).
- `R1283` calls `R1302` only when mask bit 0 is set. `R1302` calls the stop on
  the player at `L12820` when the gate is non-zero, then clears mask bit 0.
- The stop takes no argument. With a buffer at `player+0x9c` and a non-empty list at
  `player+0x10`, ordinary mode (`player+0x18 == 0`)
  stops the buffer through the DirectSound buffer's stop slot and sets `player+0xc = 1`;
  fixed-source mode reloads the selected candidate instead (`VIDEO-MUSIC-007`).
- The `0x421` arm at `L08085` calls the menu owner `R0816` only when the mask is zero.
  The owner requests a one-entry list, `music\menu.wav`, with the replace at `L12821` and
  starts at `L12822`, when the gate is non-zero.
- The menu owner skips the replace when the player's list is non-empty and its first entry
  equals `music\menu.wav` (`VIDEO-MUSIC-066`); the start still runs.

**Confidence.** High for the call order and the gates, and for the read fact that the
`+0x110` close arm clears mask bit 8 before it posts `0x41e` (`evidence/listing.txt`). Medium for the
remaining mask bits: that `frame+0x3dc` is exactly bit 0 on the mission screen before the panel (and so is
zero after teardown) rests on the mission-screen state in `MISSION-STOP-016`, not on a
witnessed value.

**Unknown.** Any other bit set in the mask at the failure, which would skip the menu request.
The audible gap between the stop and the menu start.

### VIDEO-MUSIC-063

- Result `0x446` makes the `+0x110` close arm post `0x418`. The `0x418` arm at `L12823`
  builds the save-selection dialog, stores it at `frame+0x128` and shows it with `R0361`.
  It calls no music routine in the arm read (`evidence/listing.txt`).
- A selected entry makes the `+0x128` close arm post `0x419`. The `0x419` arm at
  `L06707` calls `R1283` only when the mask equals 1, which stops the player as in
  `VIDEO-MUSIC-062`, then calls `R1303`, and for mask zero sends command `0x445` to
  `frame+0x100` and calls the load path `R1284`. The mask is zero after a taken teardown.
- A cancelled dialog (result other than `0x445`) posts `0x41e` when the session pointer
  `L00285` is non-zero and `frame+0x414 == 0xff`, which the failure arm stored. It then
  follows `VIDEO-MUSIC-062` exactly.

**Confidence.** High for the arm reads and the exclusive close chain. Medium that the mask
equals 1 when the dialog closes (see `VIDEO-MUSIC-062`).

**Unknown.** The result code of a dialog dismissed by other means than the two buttons.

**Amended.** The Unknown clause on other dismissals is answered by `MISSION-067` and `VIDEO-078`.

### VIDEO-MUSIC-064

- The load routine `R1284` reads `CurrentState`/`InBattle` with default 1 (`SAV-914`).
  A non-zero value calls `R0099` with argument 1 at `L06711`.
- `R0099`, when the gate is non-zero, calls the stop on the player at `L12824`
  before it rebuilds the session. Between the stop and `L12818` a wait loop
  (`L12825..L12826`) has three exits that return zero without a request: a pumped message
  of id `0x12` (`L12827`), a 60,000 ms timeout (`L12828`, to `L06739`) and a zero
  result of `R0509(0x64)` (`L12829`). Only if the routine reaches `L12818`, and the
  phase is not 3, does it call the list owner `R2087`, which itself requires the gate.
- `R2087`, when the gate is non-zero, builds one list of twelve strings,
  `music\B00.wav` through `music\B11.wav` in that order, replaces the player's list at
  `L12830` and starts at `L12831`. Argument 1 suppresses the map load and the `0x442`
  post (`SESS-START-036`).
- Direct closure from `R1284` reaches the stop, the replace and the start only through
  `R0099` and `R2087` (`evidence/closures.tsv`). Nothing posted by this branch
  requests music.
- The replace is the player's usual one: a random initial candidate and ordinary progression
  (`VIDEO-MUSIC-007`).

**Confidence.** High for the call order, the three exits and the arguments. A failed mission-start
exit therefore leaves the player stopped with no replacement. Medium for the closure's absence
clause: 400 functions of the load closure hold a computed call that was not read.

**Unknown.** The document loader's own effect on a running player beyond its direct
closure, which reached no music routine.

### VIDEO-MUSIC-065

- A zero `InBattle` value takes the town branch of `R1284`: `L12832`, `R0192`,
  `R0509(100)`, `R0098`, `R1670`, `R0192`, then post `0x42e`
  (`SHOP-TOWN-022`). Their direct closures reach no music routine.
- The `0x42e` arm at `L12833` calls `R1320`, which is a surface transition (`TOWN-372`) and, when the
  gate is non-zero, requests the one-entry list `music\Town.wav` with the replace at
  `L12834` and starts at `L12835`.
- The `0x42e` arm runs from the message queue after the load routine returns. The difference
  from `VIDEO-MUSIC-064` is therefore a post, not a synchronous call.
- The Town owner skips the replace when the player's list is non-empty and starts with
  `music\Town.wav` (`VIDEO-MUSIC-066`).
- No stop is issued before the replace on this branch except the one the replace routine
  itself runs first (`VIDEO-MUSIC-006`); with a skipped replace the running track continues.

**Confidence.** High for the order and arguments. Medium for the absence of any other
request: 234 functions of the `R0192` closure alone hold a computed call that was not read.

**Unknown.** Whether a later message in a town session replaces the list again before the
player hears the Town list.

### VIDEO-MUSIC-066

- The menu owner `R0816` and the Town owner `R1320` call string comparison `R0751`
  on the first entry of the player's list against their own literal and request the replace
  only when the list is empty or the strings differ.
- Owner `R2087` does no comparison and replaces unconditionally whenever the gate is non-zero.
- Consequences: a mission LOAD that reaches `L12818` with the gate set re-requests the list
  even when the `B` list is already playing; a town LOAD while the Town list plays keeps the
  running track, since that branch has no preceding stop. On Exit the teardown stop runs first
  (`VIDEO-MUSIC-062`), so the skipped menu replace avoids a list rebuild, not the stop.
- The comparison reads the first array element. Whether the random initial selection
  (`VIDEO-MUSIC-007`) reorders that array is not read, which is immaterial for the two
  one-entry lists these owners build.

**Confidence.** High for the two comparisons and the unconditional `B` request. Medium for
the equality semantics of `R0751`, read as a zero-on-equal byte comparison with a locale branch
when `[L03507]` is non-zero.

**Unknown.** The comparison's case handling under that locale branch, and the other six
list owners' behaviour, which this experiment did not read.

## Map command voices

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-067 | Ten builder tails and the selection tail call the speaker chooser, then one voice reader: move, attack, swarm, patrol, town `vt+0x6c`; guard, stand ground, defend `+0x70`; retreat `+0x74`; pickup `+0x7c`; selection `+0x78`; cast none. | High / Medium / Unknown | ● active | [EXP-0445](../experiments/EXP-0445-command-voices/) |
| VIDEO-068 | The speaker is one random member of the highest non-empty tier: hero-shaped (`+0x18c` bit 0x1), else armed human, else unarmed human; a member needs `+0x7c` set and health `+0xfc` above 0; no owner, distance or visibility field is read. | High / Medium / Unknown | ● active | [EXP-0445](../experiments/EXP-0445-command-voices/) |
| VIDEO-069 | The selection reply is `vt+0x78` (`select1` or `select2`, 2000 ms), played by the chooser speaker when Shift is up and the summary bits allow it; it shares the stamp `+0x190` with every other voice. | High / Medium / Unknown | ● active | [EXP-0445](../experiments/EXP-0445-command-voices/) |
| VIDEO-070 | Guard, Stand Ground and the defend order play the speaker bank's `defend` slot (`+0x1c`), Retreat plays `retreat` (`+0x18`) and a pickup order plays `idle` (`+0x20`), each behind the shared 3000 ms stamp. | High / Medium | ● active | [EXP-0445](../experiments/EXP-0445-command-voices/) |

### VIDEO-067

```
gesture (builder)            order  chooser call  reader   bank slot
move       (R0215)    0x16   L12836      +0x6c    command1..3 / defend
attack     (R0213)    0x19   L12837      +0x6c    command1..3 / defend
swarm      (R0235)    0x1a   L12838      +0x6c    command1..3 / defend
patrol     (R0237)    0x1d   L12839      +0x6c    command1..3 / defend
town       (R2084)    0x24   L12840      +0x6c    command1..3 / defend
guard      (R0223)    0x17   L12841      +0x70    defend
stand grd  (R0224)    0x18   L12842      +0x70    defend
defend     (R0236)    0x1b   L12843      +0x70    defend
retreat    (R0102)    0x14   L12844      +0x74    retreat
pickup     (R2088)    0x21   L12845      +0x7c    idle
selection  (R0212)    -      L12846      +0x78    select1 / select2
cast       (R0090, R0091)  0x1f 0x26 0x25 0x1e   none
```

- **Chooser callers.** `R2089` has exactly eleven direct callers, the eleven rows above with a call (a scan for the near-call opcode byte and the decoded sweep agree, 0 other sites; `evidence/xref-voice.txt`). Each builder stores its order byte at offset 9 of the request object (`evidence/opcode-stores.txt`), sends the order (`R0217`) and then calls the chooser, so a refused reply never holds back the order. The reader of each tail is the indirect call through a vtable slot after the chooser's null test.
- **Cast.** The two cast builders hold no chooser call; their only indirect calls are `vt+0x7c` of the object `R0347` returns (`evidence/calls-cast-builders.txt`).
- **Input surfaces.** The map click `R0211` reaches the builders by cursor (`AI-CLICK-050`); the minimap handler `R0233` reaches move, attack, swarm, defend and patrol (`AI-MINIMAP-062`); the command panel `R0097`, which has 14 direct callers, reaches guard, stand ground and retreat through the jump table `L00645` for buttons 3, 7 and 8 (`AI-PANEL-123`). The three panel builders have no other caller.
- **Draw-state gate.** Move and swarm, and only these, skip the reader when the chosen unit's `+0x74` equals 1 (`L12847`, `L12848`); `ANIM-STATE-002` reads draw state 1 as the move action. The chooser has already picked the unit, so a moving speaker silences the reply and no second unit is tried.
- **Executed.** Each tail, started at its order push and stopped before its epilogue, with one selected unit in each of the hero, hero-shaped mage, mercenary and peasant banks: one command-service call, then the slot of the table. Acknowledgement 0 removes every reply and keeps the command; a clock of 2500 ms (stamp 0) removes every reply; `+0x74 = 1` removes move and swarm only.
- **Reader census.** `evidence/vcalls-census.txt` lists the 340 indirect calls whose displacement is `0x6c`, `0x70`, `0x74`, `0x78` or `0x7c`. 224 follow a call to `R0347` (222 of them `vt+0x7c` of the object it returns). 11 are the tails above. 94 have a `PUSH` within the six preceding instructions; the five routines take no stack argument (`RET` without a count). The remaining 11 (`evidence/vcalls-classification.txt`) include four user-interface routines that compare or store the returned value (`L12849`, `L12850`, `L12851`, `L12852`), the routine `L12853` that stores the result, and `L12854`, whose receiver is a child-control lookup.

**Confidence.** **High** for the eleven chooser callers, their orders, readers and slots: decoded whole and executed on both roots with a sentinel per slot. **Medium** that no other path reaches the five slots on a unit: the census sorts the other sites by window and by use of the result, not by receiver proof. **Unknown** the receivers of `L12855` and of the four sites in `0x00573..0x00583` (`L12856`, `L12857`, `L12858`, `L12859`), which were not read.

### VIDEO-068

- **Gates.** `R2089` returns 0 unless the session's `+0x6bc` is 2 (`L12860`, the campaign; `DLG-ENTRY-016`) and the Acknowledgement cell is non-zero (`L12861`; `VIDEO-SFX-054`).
- **Population.** It walks the selection map `view+0x9b8` (`AI-SELECT-065`) from the first bucket in bucket-chain order. For each member `unit = value`:
  - `[unit+0x7c]` must be non-zero (`L12862`);
  - `MOVSX word [unit+0xfc]` must be above 0 (`L12863`, `L12864`), so health 0 and every negative value are excluded;
  - `+0x18c & 1` appends the unit to list A (`L12865`..`L12866`);
  - otherwise, while A is empty, `+0x18c & 0x10` appends it to list B when `+0x15c` is non-zero (`L12867`..`L12868`), else to list C while B is empty (`L12869`..`L12870`);
  - a member with neither bit is dropped.
- **Pick.** After the walk: list A if non-empty (`L12871`), else list B (`L12872`), else list C (`L12873`), else 0. The index is `rand() * n / 0x7fff` (`L12874`..`L12875`, repeated at `L12876` and `L12877`), near-uniform (counts differ by at most two of the 32768 draws; draw 32767 selects the slot one past the last member). Without a hero-shaped unit the order is armed humans, then unarmed humans. A and B or C members speak from the bank `HERO-APPEAR-055` gives them: the hero or mage bank for A, the mercenary bank for B, the peasant bank for C, the mage bank for a B or C drawable with bit 0x2.
- **Fields read.** The only unit displacements the chooser reads, over all 20 executed cases and in the listing, are `+0x7c`, `+0xfc`, `+0x15c` and `+0x18c`; it reads no owner, player, position, distance or visibility field and writes nothing (`evidence/probe.json`, `access`). The stamp is not read, so a unit on cooldown can be picked and then stay silent; no second candidate is chosen.
- **Executed.** Hero before or after an armed and an unarmed human: the hero for every draw. Armed before or after unarmed: the armed one. A monster alone (neither bit): no speaker and no `rand()` call. A mage bit with bit 0x10 and no hero bit: tier B or C by `+0x15c`. Three heroes: draws 0, 10922 | 10923, 21845 | 21846 and 32766 give members 0, 0 | 1, 1 | 2 and 2. Health 0 and -10 excluded, 1 included. `+0x7c = 0` excluded. Empty map, Acknowledgement 0 and session mode 1: no speaker.
- **Ownership.** The chooser does not test it. Upstream, the command panel is disabled for a selection whose primary object another player owns (summary bit `0x4`, `AI-PANEL-061`) and the hover cursor gives select or default when that bit is set (`AI-CURSOR-226`); a click selection can hold a foreign unit (`AI-SELECT-122`).

**Confidence.** **High** for the gates, per-member tests, tier order and index formula: the routine is read whole and executed on both roots over 20 discriminating cases, including both list orders. **Medium** for the meaning of the bits and the flag: `+0x18c` bit 0x1 is the hero-shaped drawable (`HERO-APPEAR-041`), bit 0x10 follows a wire class below 0x1a (`ANIM-096`), and `+0x7c` as "selected" is inferred (`AI-SELECT-065`). **Unknown** the content of the list slot one past the last element, which a draw of 32767 (1 in 32768) selects; the executed stub returns 0 there. **Unknown** whether any order can be issued from a foreign selection, and whether a structure can be a member that reaches the chooser.

### VIDEO-069

- **Caller.** The only call of `vt+0x78` on a unit is `L12878`, after the chooser call `L12846`, at the end of `R0212`. Its three callers (`L12879`, `L12880`, `L12881`) are in the map-click handler `R0211`: the click with an empty selection, the select cursor arm and the drag-rectangle branch (`AI-CLICK-050`). The key and digit group selections are not callers.
- **Gates.** After the selection summary is rebuilt (`L12882`), the tail tests: the Shift latch `[L00670]` clear (`L12883`, `AI-KEYMOD-059`), the selection count `view+0x140` non-zero (`L12884`), `view+0x144 & 1` (`L12885`) and `view+0x144 & 4` clear (`L12886`). The chooser's own gates of `VIDEO-068` follow.
- **Executed.** All 16 combinations of Shift 0 or 1, count 0 or 1 and flags 0, 1, 4 and 5: only Shift 0, count 1, flags 1 plays (`select1` for a hero at draw 0).
- **Sharing.** The reader is `R0579` (`ANIM-119`): it uses the chooser's speaker and the stamp `unit+0x190` of every other voice, and its threshold is 2000 ms. Executed on one unit: a move reply at 5000 ms then a selection reply at 6999 ms is silent, and at 7000 ms plays; a selection reply at 5000 ms then a move reply at 7999 ms is silent, and at 8000 ms plays; a move reply then a retreat reply is silent at 2999 ms and plays at 3000 ms.

**Confidence.** **High** for the gate matrix and the shared stamp: the original tail and readers executed on both roots. **Medium** for the bit meanings (`+0x144` bit 0x1 follows the class-name test "CUnit", bit 0x4 is ownership of the primary object; `AI-PANEL-061`, which does not establish the other bits) and for the absence of other callers (the eleven-caller census of `VIDEO-067`). **Unknown** which input branches of the routine reach the tail: it was executed from its first gate.

### VIDEO-070

- **Slots.** Guard (`0x17`) and Stand Ground (`0x18`) are panel builders; the defend order (`0x1b`) is a map and minimap builder. All three call `vt+0x70`, `R0580`, which reads `[bank+0x1c]`, `defend.wav` (`ANIM-094`). Retreat (`0x14`, panel) calls `vt+0x74`, `[bank+0x18]`, `retreat.wav`. A pickup order (`0x21`, map click) calls `vt+0x7c`, `[bank+0x20]`, `idle.wav`. None draws from `rand()`.
- **Banks.** The bank is the speaker's (`HERO-APPEAR-055`): `mf_hero` or `ff_hero`, `m_mage` or `f_mage`, `mf_merc` or `ff_merc`, `m_peasant` or `f_peasant`. Executed with one selected unit in a hero, hero-shaped mage, mercenary and peasant bank: guard, stand ground and defend decode to `defend` of that bank, retreat to `retreat`, pickup to `idle`.
- **Stamp.** The three share the 3000 ms stamp with every other voice, so after a reply of any gesture that unit's next command, defend, retreat or idle reply is refused for 3000 ms and its next selection reply for 2000 ms.
- **Other stance changes.** A stance or retreat set by anything other than these builders reaches none of the five slots on a unit within the census of `VIDEO-067`.

**Confidence.** **High** for the slot of each gesture and the bank rule: decoded and executed on both roots. **Medium** for the last bullet, which depends on the census.

## Cutscene presentation

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-071 | The player copies the sidecar start to the blit source origin once after open, arms each fade at its start frame, scales the decoder palette per frame, and adds each pan step to the source origin after the frame advance. | High / Medium / Unknown | ● active (amended, partially retracted) | [EXP-0447](../experiments/EXP-0447-cutscene-presentation/), [EXP-0460](../experiments/EXP-0460-cutscene-sound/) |
| VIDEO-072 | A movie is drawn doubled when twice its width fits the 640 wide output region or twice its height fits the 360 high region; a 480 high movie then takes its own size as the blit region. Display size does not enter. | High / Medium / Unknown | ● active (amended) | [EXP-0447](../experiments/EXP-0447-cutscene-presentation/), [EXP-0460](../experiments/EXP-0460-cutscene-sound/) |
| VIDEO-073 | Frame pacing is the decoder wait: the sound playback position when a track is open and on, else a timer in 10 microsecond units from the header interval. The player adds no floor or ceiling to the interval. | High / Medium / Unknown | ● active (amended) | [EXP-0447](../experiments/EXP-0447-cutscene-presentation/), [EXP-0460](../experiments/EXP-0460-cutscene-sound/) |
| VIDEO-074 | Both roots hold 33 movies each; every header has flags 0 and one 16-bit 22050 Hz audio track with the compression flag set, mono in 15 or 16 movies and stereo in the rest. Intervals: -4000, -6666, -6673, -8333. | High | ● active | [EXP-0447](../experiments/EXP-0447-cutscene-presentation/) |
| VIDEO-081 | With a track open and on, the frame wait compares a DirectSound byte position, wall-clock interpolated, with a byte target; no time or count bounds it. Without an opened track the timer paces (read, not executed; no sound output opened). | High / Medium / Unknown | ● active | [EXP-0460](../experiments/EXP-0460-cutscene-sound/) |
| VIDEO-082 | At the final frame the player blits, then closes the movie before it returns: SmackClose stops and releases each sound buffer, so queued sound is cut rather than drained (read, not executed; no sound output opened), before the next screen. | High / Medium / Unknown | ● active | [EXP-0460](../experiments/EXP-0460-cutscene-sound/) |
| VIDEO-083 | Read, not executed: a missing sidecar raises an exception from the registry open, a negative count reads no record, and equal start and end frames leave a fade active and a pan running to the end. No shipped sidecar does any. | High / Medium / Unknown | ● active | [EXP-0460](../experiments/EXP-0460-cutscene-sound/) |
| VIDEO-084 | The fade step scales the current frame's palette, set at open or by SmackNextFrame; the indexed blitter does not clip its source on synthetic memory, 16-bit modes matched in range, and the open compares nothing against the interval. | High / Unknown | ● active | [EXP-0460](../experiments/EXP-0460-cutscene-sound/) |

### VIDEO-071

Frame step `R2090`, called by the player `R0719` with argument 1, runs in this order:

1. Wait: the loop `R2091` repeats `SmackWait` until it returns 0.
2. `SmackBufferFocused` on the buffer; a zero return ends the step with 1, before any state below changes.
3. Fade arm `R2092`, then pan state `R2093`.
4. Fade apply `R2094` on the decoder palette (`handle+0x6c`). With a fade active the scaled palette goes to `SmackBlitSetPalette`; otherwise the raw palette goes there when `handle+0x68` is non-zero.
5. `SmackDoFrame`, display surface lock (`vt+0x64`; a failed lock at `L12887` ends the movie, return 0), `SmackBlit`, unlock (`vt+0x80`).
6. If the decoder frame counter `handle+0x374` equals frames minus 1 the step returns 0 and no advance runs. Otherwise `SmackNextFrame`, then pan add `R2095`, return 1.

- **Reader.** `R1399` reads, with default 0 for every key (strings `L12888` to `L12889`): the `Common` keys `startx`, `starty`, `nFadings`, `nPanaramings` into player `+0x2c`, `+0x30`, `+0x14`, `+0x20`. It allocates `count * 16` zeroed bytes per family and fills record `i` from section `Fading<i+1>` or `Panaraming<i+1>`: a dword start frame at `+0`, a dword end frame at `+4`, then two floats (`startfade`, `endfade`) or two dwords (`stepx`, `stepy`). The section and key schema is `REG-CUT-053`.
- **Blit arguments.** `L12890` to `L12891` pushes the `SmackBlit` parameters from the player: source x `+0x348`, source y `+0x34c`, source width `+0x350`, source height `+0x354`, destination x `+0x358`, destination y `+0x35c`. The executed blit probe of `VIDEO-072` shows the source rectangle is read at the source coordinates and the destination origin is where it lands, which separates source from destination.
- **Start.** After `R2096` succeeds, `L12892` copies `+0x2c` and `+0x30` to the source origin `+0x348` and `+0x34c`. The destination `+0x358` and `+0x35c` is set by the player before the loader, and no sidecar value reaches it. The copy follows the doubling arm of `VIDEO-072`, so a start is never halved. All 33 shipped sidecars hold `startx = starty = 0` (`REG-CUT-053`).
- **Fade.** The arm tests that the record index `+0x18` is below `+0x14` and that the record start equals the frame counter. It sets active `+0x34`, advances the index, stores remaining `+0x344 = end - start`, factor `+0x33c = startfade` and delta `+0x340 = (endfade - startfade) / (end - start)`. Each apply adds delta to the factor, multiplies all 768 palette bytes by it with truncation (`__ftol`, no clamp) into `+0x38`, and decrements remaining; at 0 the fade is inactive. A fade from frame `s` to `e` is applied on frames `s` to `e - 1`: frame `s` already shows `startfade + delta` and frame `e - 1` shows `endfade`, within float32 accumulation and with the bytes truncated (a fade-in ending at 1.0 can give 254 for 255).
- **Pan.** Per step, with the record index `+0x24` below `+0x20`: a record start equal to the counter loads `+0x360` and `+0x364` with the record's step, and a record end equal to the counter zeroes both and advances the index. After each frame advance `R2095` adds the two steps to `+0x348` and `+0x34c`, with no clamp. The first blit with a moved origin is frame `s + 1`; frame `e` is drawn at the start origin plus `(e - s)` steps.
- **Order.** Both families consume records strictly in index order. A fade whose start frame is already past, or a pan whose end frame is already past, never matches, and no later record of that family runs.
- **Executed.** None of these ROM routines was executed. The blit and wait calls they make were executed for `VIDEO-072` and `VIDEO-073`.

**Confidence.** **High** for the reader layout, the consumer of each value, the frame step order and the arithmetic: whole routines read on the one `rom.exe` both roots hold, with the key strings read from its image. The live alternatives (gamma factor, blit scale, destination offset, rectangle) are excluded because the fade writes only the palette and the pan only the source origin. **High** that the palette scaled at step 4 is the current frame's (`VIDEO-084`). **Unknown** nothing further here that `VIDEO-083` and `VIDEO-084` do not bound.

**Amended.** The clause that the palette scaled at step 4 is the previous decode's is withdrawn (`retracted.md`); `VIDEO-084` records the current frame's palette. The reader's result for a missing sidecar, a negative count and equal start and end frames is in `VIDEO-083`, and the clip of an out-of-range source origin in `VIDEO-084`.

### VIDEO-072

- **Region.** The player calls the loader with width 0x280 and height 0x168 (`L12893`), stored in `+0x350` and `+0x354`, and the destination origin `(L09808, L09809 + 0x3c)`. After the loader returns, a movie of height 0x1e0 replaces the region with its own size (`R2097`, `L12894` to `L12895`) and the origin with `(L09808, L09809)`. The doubling decision below runs inside the loader, before that replacement, so it uses 640x360 for every movie.
- **Decision.** `R2096` compares unsigned. If `2 * width > region width` and `2 * height > region height` it calls `SmackBlitOpen(mode | 1)` and sets doubled `+0x368 = 0`. Otherwise it calls `SmackBlitOpen(mode | 2)`, shifts source x, y, width and height right by 1, halves the destination origin toward zero, and sets `+0x368 = 1`; close `R2098` shifts them back. `mode` comes from `R2099`: 0 for an 8-bit surface, `0x80000000` for 5-5-5, `0xc0000000` for 5-6-5, else the message `Unsupported pixel format.`.
- **Inputs.** Movie width and height, and the two region constants. The surface pixel format enters only through `mode`, and no window size or movie flag is read. The display size words `-640`, `-800` and `-1024` set only the origin, `(width - 640) / 2` and `(height - 480) / 2` (`R0341`).
- **Blit executed.** `tools/smackblit` calls the installed decoder's `SmackBlitOpen` and `SmackBlit` on synthetic indexed memory, source rectangle 5x3 at (3,1), destination (7,2). Mode 1 writes 15 cells, a plain copy. Mode 2 writes 60 cells, each source pixel as a 2x2 block, and doubles the destination origin to (14,4), so the halving above is undone by the blitter.
- **Shipped population.** Census of `VIDEO-074`. Doubled by the rule: the `320x180` movies, 4 on EN (`INTRO/04` and `M150/01`, each in both containers) and 2 on RU (`INTRO/04`), drawn as 640x360. Not doubled: 22 `640x360` and 4 `800x360` on EN, 24 and 4 on RU, and 3 `640x480` logos on each root. An `800x360` movie shows a 640 wide window of its source.

**Confidence.** **High** for the decision, the constants, the halving arithmetic and the indexed blit kinds: whole routines read and the blitter executed. **Medium** that no other call of the loader sets a different region: `R1399` has the single call site `L12893`, but a computed call is outside the read. **Unknown** whether ROM1 ever runs in a 16-bit surface mode.

**Amended.** The blit output for the two 16-bit modes and the effect of the halving on an odd destination origin are in `VIDEO-084`.

### VIDEO-073

- **Player.** Step argument 1 selects the blocking wait; the only pacing in the player is the `R2091` loop. All pending messages are drained before each step (`L12896`) and none during the wait; `WM_CLOSE`, `WM_QUIT`, key down, system key down and left or right button down (`0x10`, `0x12`, `0x100`, `0x104`, `0x201`, `0x204`) end the movie at `L12897`. The constructor calls `SmackSoundUseDirectSound` (`L12898`), which installs the readback `D00003`. The player opens with track flags `0xff000` (bit `0x80`, the frame-rate override, is clear), so the header interval is used, and it turns sound on after open (`SmackSoundOnOff(h, 1)`, `L12899`).
- **Interval.** Open reads the signed interval `v` at header `+0x10` (`D00004`, `D00005`; the open and `SmackWait` bodies `D00006` to `D00007` are the `VIDEO-050` listing). For `v >= 0` the state value `+0x424` is `v * 100` modulo 2^32; for `v < 0` it is `-v`. The unit is 10 microseconds, so `-6666` is 66.66 ms and a positive value is milliseconds. The clock is `timeGetTime * 100`.
- **Timer wait.** Used when no track opened (`+0x444 = -1`). The deadline `+0x420` starts at -1: the first poll sets it to now plus the interval and returns not ready. After that a poll with now below the deadline is not ready. Otherwise it is ready and the deadline becomes: now plus interval if now is within one interval after the deadline; deadline plus interval if now is up to one further interval late; now plus interval if later. A non-zero interval sets the latch `+0x41c`, which keeps later polls ready until `SmackNextFrame` clears it.
- **Sound wait.** Used when a track opened and sound is on. Each `SmackDoFrame` adds the bytes per frame to the track targets and arms `+0x448` (`D00008`). The wait is not ready while target `(+0x20 >> 10) + +0x24` exceeds the playback position plus 8; becoming ready clears `+0x448`, and while it is clear polls are ready. The position readback `[D00009]` is `D00003` for the DirectSound backend, interpolated with `timeGetTime`. The first track that opens is the clock.
- **Audio ahead or behind.** Behind: the wait holds the frame until the position reaches the target, with no timeout. Ahead: the wait returns ready without sleeping, and `SmackDoFrame` skips the decode (`D00010` returns 1 once the position passes the second target `+0x28`) unless open flag `0x400` is set. ROM1 does not set it and ignores the return, so it blits the unchanged buffer and advances. Sound off (`+0x44c` non-zero) makes every poll ready, with no pacing at all.
- **No audio.** A movie whose track word lacks bit 30 or has rate 0 opens no track and uses the timer. Executed: `tools/smackpace` polled `SmackWait` on header-patched scratch copies of one shipped movie, opened with no track flag. The shipped interval -6666 became 6666 units and the first ready poll of frame 1 came at 66 ms.
- **Unfocused.** `SmackBufferFocused` is 0 unless the buffer's window is the foreground window, or, for a direct-draw buffer, its focus word `+0x460` is non-zero. The step then returns before decode or advance, so the movie holds its frame and the loop spins on ready polls.
- **Degenerate intervals.** Frame 1, first ready poll, cap 450 ms. Header 0 gives 0 units, ready at once, and the latch is never set. Headers -1 and 1 give 1 and 100 units, ready within 5 ms. Header 100 gives 100 ms, ready at 99 to 100 ms. Header -100000 gives 1 s, not ready within the cap. Header 2147483647 gives 4294967196 units, a deadline 100 units in the past, ready at once. Header 42949673 gives 4 units, ready at once; a positive value of 42949673 or more wraps. Header -2147483648 negates to 2147483648 units, and whether it waits depends on the boot-time phase of `timeGetTime * 100`, so it is not in the reproduced evidence. With sound open, interval 0 arms nothing: no wait and no skip. The player reads no interval and applies no floor or ceiling.

**Confidence.** **High** for the interval derivation, the timer wait arithmetic and the absence of a floor in the player: read in full, and executed on the installed decoder through `timeGetTime`, timer path only. **Medium** for the sound wait, the skip rule and the unfocused behaviour: read in full, but no sound output was opened, so no sound position was observed. **Unknown** the audible timing and the cursor's real advance rate, and the effect of losing focus while sound plays.

**Amended.** The sound position, its bounds and the result with no audio device are in `VIDEO-081`; what the final-frame return does to queued sound is in `VIDEO-082`.

### VIDEO-074

- **Population.** Every `.res` and `.lm` container on each root plus loose `.smk` files (none found): 12 EN and 11 RU archives opened, 33 `.smk` nodes per root, 18 in `VIDEO4.RES` and 15 in `VIDEO8.RES`. No node of another name carries the `SMK2` magic. `rom.exe` also names `video4\rom.smk` (`L12900`, `L12901`); no container node or loose file has that name, and what the call does without it is Unknown. The header fields are those the decoder's open routine reads: `+0x14` flags, `+0x48 + 4 i` audio word of track `i`.
- **Audio word.** Bit 31 compressed, bit 30 present, bit 29 16-bit, bit 28 stereo, low 24 bits the rate. Track 0 of every movie is `e0005622` (flag set, 16-bit, mono, 22050 Hz) or `f0005622` (the same, stereo); bits 27 to 24 are clear in 66 of 66 and the decoder reads only bits 31 to 28 and the rate. Tracks 1 to 6 hold `00000000` in all 66 headers. The decoder routes bit 31 to the compressed-audio routine `D00011` (`D00012`, `D00013`) and copies raw otherwise. The stored coding is not decoded here; the routine reads a bit-tree stream, consistent with Smacker audio, which is inference.
- **By container.** `VIDEO8.RES` holds 15 stereo movies on each root. `VIDEO4.RES` holds 15 mono and 3 stereo on EN, 16 mono and 2 stereo on RU. Its stereo movies are the logos `LOGOS/1c`, `buka` and `nival` on EN, and `buka` and `nival` on RU; RU `LOGOS/1c` is mono.
- **Header.** Flags are 0 in 66 of 66 (no ring, no y-scale). Intervals: `-6666` in 30 EN and 29 RU, `-6673` (`M120/01`, both containers) in 2 on each root, `-4000` (`LOGOS/buka`) in 1 on each root, `-8333` (RU `LOGOS/1c`) in 1. Frames run 75 to 1309 and durations 5.0 to 87.3 s. No interval is zero or positive.

**Confidence.** **High** for the census: every node of the named archives, parsed by `tools/smkcensus`. It does not cover movie files outside `Allods/VIDEO4.RES` and `Allods/VIDEO8.RES`, or the audio payload.

### VIDEO-081

- **Backend.** The player constructor `R2100` calls `SmackSoundUseDirectSound(0)` (`L12898`) and ignores its result. The export (`D00014`) loads `DSOUND.DLL` and installs the DirectSound routines: the service `D00015` in `[D00016]`, the track open `D00017` in `[D00018]`, the track close `D00019` in `[D00020]` and the position readback `D00003` in `[D00009]`. It returns 0 and installs nothing when `LoadLibrary` returns a value below 0x20 (`D00021`). `SmackWait` calls the service on a poll when the open-track count `[D00022]` is non-zero (`D00023` to `D00024`).
- **Timer feed.** When the first track links into the device list (`D00025` to `D00026`) and the hook word `[D00027]` is zero, the track open calls `timeSetEvent(75, 25, D00028, 0, periodic)` (`D00029`). The callback `D00028` runs the service `D00015`. `D00030` is a second caller of the service, installed on the other branch through hook pointers `[D00031]` to `[D00032]`, which were not read. The service lets one entrant through at a time (`[D00033]`, `D00034` to `D00035`) and at most every 10 ms (`D00036` to `D00037`). So the buffer is fed and the cursor read from a timer thread that does not depend on the wait polls. Read, not executed.
- **Position.** `D00003` runs the service, then derives played bytes: the track's written total `+0x70` minus the queued bytes (`+0x44` buffer size minus free bytes `+0x74`); with no buffer (`+0x60` zero) it is the written total. The service raises `+0x74` by the advance of the DirectSound play cursor (vtable `+0x10`, previous cursor `+0x78`, `D00038` to `D00039`). The unit is bytes of decoded 16-bit audio at the output rate. A negative result is clamped to 0.
- **Interpolation.** The readback then reads `timeGetTime`. If the count equals the previous count `+0x84` and is not zero, the result is the previous count plus elapsed milliseconds times the rate field `+0x14`, divided by 1000, capped at the written total; otherwise the count and the time are stored (`+0x84`, `+0x88`). The result is never below the stored maximum `+0x8c`, which it updates. The wait compares this value plus 8 with the byte target of `VIDEO-073`.
- **No bound.** The ROM loop `R2091` (`R2091` to `L12902`) calls `SmackWait` until it returns 0 and holds no counter, clock test or sleep. No message is pumped inside it, so a held wait leaves the window unresponsive and no key or click ends it. `SmackWait` (`D00040`) has no time test: a not-ready poll calls the stream read-ahead `D00041` and returns 1. Its ready exits other than the target being met are state exits: the arm flag `+0x448` is clear, sound is off `+0x44c`, there is no clock track, or the stream-error flag `+0x3a8` (set at `D00042` and `D00043`) is non-zero. None tests elapsed time or a count.
- **Stalled position.** The service never tests the return of `GetCurrentPosition` (`D00044` to `D00045`; on failure the stack slot keeps the `timeGetTime` value stored at `D00046`) and skips the buffer fill when `Lock` fails (`D00047`). Silence written into gaps is not added to `+0x70` or `+0x74` (`D00048` to `D00049`), so silence never advances the position. When the cursor stops, the interpolation reaches only the written total, which stops growing once the buffer is full or the supply ends. The timer feed does not change this: it moves data into the buffer, and the buffer is bounded. A target above that total is never met. After the supply ends and the queue drains, the service stops the buffer and sets free bytes to the buffer size (`D00050` to `D00051`) and the readback returns the written total.
- **Track shorter than the video.** By the same arithmetic a track whose decoded bytes fall below the target holds the wait for good.
- **No device, sound off, track failure.** `SmackOpen` allocates a track for each requested track bit and calls the backend open (`D00052` to `D00053`, `D00054`); on failure it frees the track and leaves the clock track `+0x444` at -1 (`D00055` to `D00056`, `D00057`). The DirectSound track open `D00017` returns 0 when device init `D00058` or buffer creation `D00059` fails (`D00060`, `D00061`), without a call to another backend. With `+0x444` at -1, `SmackWait` takes the timer path of `VIDEO-073` (`D00062`) and `SmackSoundOnOff(h, 1)` returns at once (`D00063`). Among the 16 Smack import slots of `rom.exe` (`L12903` to `L12904`, `evidence/static/rom-refs.txt`), `SmackSoundOnOff` has two call sites, `(h, 1)` after open (`L12899`) and `(NULL, 0)` in the close (`L12905`), which returns 0 at `D00064`; `SmackVolumePan` is not imported. Use of `GetProcAddress` to reach other decoder exports was not searched.
- **Shipped population.** Track 0 of every movie: 66 of 66 end with 12 to 25 frames that hold no track-0 chunk (EN: 14 in 2, 15 in 30, 25 in 1; RU: 12 in 1, 14 in 2, 15 in 29, 25 in 1), so the audio is delivered about one second ahead of the last frame. The unpacked byte total of the chunks, less frames times bytes per second times the interval, lies between -82 and +2 bytes in 66 of 66. The final wait target is one frame of audio below that total, so the shorter-track stall above is not reached by any shipped movie, by this arithmetic.

**Confidence.** **High** for the position formula, the interpolation, the absence of any time or count bound in the ROM loop and `SmackWait`, and the shipped census: whole routines read on the one `rom.exe` and the one `smackw32.dll` both roots hold, census of every `.smk` node on EN and RU. **Medium** for the stalled-position, device-failure and track-open-failure outcomes: read in full, but no sound output was opened, so no DirectSound call ran. **Unknown** the audible timing and the real cursor advance rate, how the timer thread behaves during a held wait (whether it keeps feeding the buffer and reading the cursor while the main thread loops, and whether it survives a lost device), the hook branch of the timer start, the waveOut backend's behaviour (only its readback was read; the ROM never installs it), what substitutes for a `DSOUND.DLL` that does not load (the wait's track-open pointer then keeps its load-time value, which was not read), and whether the stream-error flag `+0x3a8` is ever set on a shipped movie.

### VIDEO-082

- **Order.** The step `R2090` blits the final frame (`L12891`), unlocks the surface, finds the decoder frame counter equal to frames minus 1 and returns 0 (`L12906` to `L12907`) before `SmackNextFrame`. The player `R0719` treats 0 as the end: it calls the close `R2101` (`L12908`, `L12909`), then the destructor `R2102`, then returns (1 for a natural end, 0 for the interrupt branch at `L12897`, which closes the same way). No wait runs after the final frame, so the movie closes at once after its last blit.
- **Close.** `R2098` calls `SmackBufferClose`, `SmackClose(handle)`, `SmackBlitClose`, then `SmackSoundOnOff(NULL, 0)`, which returns 0 without effect (`D00064`). The scratch state is cleared and the doubled origin restored.
- **Sound.** `SmackClose` (`D00065`) walks the seven track slots from `+0x428` and calls the backend track close `[D00020]` for each open one (`D00066` to `D00067`). The DirectSound close `D00019` unlinks the track and calls `D00068`, which calls `Stop` (vtable `+0x48`) and `Release` (vtable `+0x8`) on the buffer. No routine on this path waits for the queue. Queued sound is stopped and released, not drained or left playing. When the last track closes the device is shut down (`D00069`).
- **Next screen.** The close is complete inside the player before it returns, so the next screen starts after the sound has been stopped. The player has five call sites in three functions: `R2035` (three), `R1658` and `R0701` (`evidence/static/rom-refs.txt`).
- **Cut length.** The last wait before the final frame holds until the position reaches the target of frames minus 1 steps (`VIDEO-081`). The remainder queued at the stop is therefore about one frame of audio less the census deficit of `VIDEO-081`, which is at most one frame step: 0.04 to 0.084 s at the shipped intervals of `VIDEO-074`. This counts unplayed audio, whether queued in the buffer or not yet written, so it does not assume the buffer is fed only by wait polls. It is derived from the target arithmetic, not heard.

**Confidence.** **High** for the call order and the stop and release of the buffer: whole routines read, with the ROM call sites enumerated (`evidence/static/rom-refs.txt`). **Medium** that the cut sound is about one frame long and that no audible tail survives the release: the arithmetic is read from `SmackDoFrame` and the readback, and no sound output was opened. **Unknown** what the operating system does with a released buffer's last milliseconds, and the mixer's own latency.

### VIDEO-083

- **Reader.** `R1399` opens the sidecar through `R1452`, then reads `Common` keys and records as in `VIDEO-071`. The four functions of the chain (`R0719`, `R1399`, `R1452`, `R2096`) have no try block in their exception tables: `evidence/eh/eh-en.tsv` and `eh-ru.tsv` give the header of each (magic `0x19930520`, try-block count 0) for these four and for `R0701` (45 states) and `R1658`, two of the three functions holding the player's call sites. `R2035`, the third, has no `push -1; push handler` prologue in its first 32 bytes, hence no handler of that form.
- **Missing sidecar.** When the file open `[vt+0x28]` fails (`L12910`, `L12911`), `R1452` formats a message with the file name (`L12912`) and throws it through `RaiseException` (`R2103`, `L12913`), whose record template at `L12914` carries the C++ exception code `0xe06d7363`. Nothing in the chain catches it. Whether a frame above the direct callers catches it, or the process handler ends the game, is not read. No default value is substituted.
- **Negative count.** The count is read as a signed dword. The record allocation takes `count << 4` (`L12915`, `L12916`) and its result is tested for 0. The fill loop (`L12917`, `L12918`) and both consumers (`L12919`, `L12920`) compare signed, so a negative count reads no record and consumes none. Whether `R0747` returns 0 or raises for the oversized request is not read.
- **Equal start and end frames, fade.** The arm matches the start frame, stores remaining `end - start = 0` and divides by it (`L12921` to `L12922`), so the delta is infinite or not a number. The apply step then decrements remaining to -1 (`L12923` to `L12924`), which is never 0 again: the fade stays active to the last frame, scaling by the infinite or not-a-number factor through `__ftol` (`R0279`). The resulting palette bytes are not derived.
- **Equal start and end frames, pan.** The pan state tests the start frame first and returns (`L12925` to `L12926`); the end-frame test, which advances the record index, runs only when the start frame does not match. At equal frames the step loads at `s`, the end test never runs, the index never advances, and the pan runs from frame `s + 1` to the last frame; later records are never consumed.
- **Shipped population.** Every `.smk` node of both roots has a sidecar: 33 of 33 per root (`REG-CUT-053`), none unparsable. Per root the 33 sidecars hold 61 fade records and 4 pan records. Counted over those: 0 equal, reversed, overlapping or unordered start and end frames, 0 start frames at or past the frame count, 0 unread sections, 0 counts that differ from the sections present, 0 non-zero `startx` or `starty`. The 4 pan records are the two 800 wide movies of `VIDEO-072`, listed twice (once per container). Replaying the pan arithmetic over the movie's frames gives a source origin extent of 160, which equals the slack between the movie width and the 640 wide blit region, so the source rectangle ends exactly at the end of a row.

**Confidence.** **High** for the shipped population (census of every sidecar by `tools/sidecarcensus`, framing as in `REG-REC-032`) and for the pan arithmetic. **Medium** for the missing-sidecar, negative-count and equal-frame outcomes: read in full, not executed, and no shipped movie reaches them. **Unknown** the byte values of a fade by an infinite or not-a-number factor, the allocator's behaviour for the oversized request, and what lies above the direct callers.

### VIDEO-084

- **Palette at the fade step.** The fade step scales the 768 palette bytes at `handle+0x6c` (`VIDEO-071`). Executed on the installed decoder in a 386 child (`tools/smackstate`), no audio flag, no window: for every `.smk` node of `VIDEO4.RES` and `VIDEO8.RES` of both roots (33 per root), `SmackOpen` leaves a non-zero palette and sets `NewPalette` `+0x68` (66 of 66). `SmackDoFrame` changes the palette bytes 0 times in 17844 stepped frames. After `SmackNextFrame` the palette changes 432 times in all, in 27 of the 66 movies, and per movie the number of changes equals the number of advances that leave `+0x68` set. `evidence/state-*/palette-changes.tsv` lists each of the 432 changes with the step of the `SmackNextFrame` call, the frame counter after it (always step + 1) and the flag `+0x68` (always set): the palette for frame N is in place after step N-1's `SmackNextFrame` and before step N's fade step and `SmackDoFrame`. Inference: `SmackNextFrame` loads the next frame's palette chunk, so the palette at step 4 is the current frame's, the open palette for frame 0. The probe shows when the bytes change, not which frame's chunk they come from. The earlier reading that it is the previous decode's is withdrawn.
- **Blit clipping.** In the two indexed modes only (1 and 2), `SmackBlit` called on synthetic memory (`tools/smackedge`), with a source origin inside the source rows, past the right edge, left of the row, below the last row, above the first row and 40 columns right: all 6 cases write the full rectangle (15, 15, 15, 20, 15, 15 cells; 4 times as many doubled), each cell equal to the linear source address `row * pitch + column`, guard bytes included. In those modes the blitter does not clip the source rectangle: it reads whatever lies at the computed address.
- **16-bit modes.** The modes `0x80000001`, `0x80000002`, `0xc0000001` and `0xc0000002` on synthetic memory write the same cell counts and bounding boxes as the indexed modes (15 cells plain, 60 doubled, 5x3 source) with each cell equal to the learned output word of its source index, for destination origins (7,5) and (6,4). The 16-bit rows ran only the in-range source origin (3,1); no out-of-range origin was run in a 16-bit mode. The doubling mode places the block at twice the origin argument in both axes, odd or even. Whether `R2099` returns a 16-bit mode in practice depends on the surface the display setup creates, which was not read.
- **Odd origin.** The loader halves the destination origin toward zero (`L12927` to `L12928`) and the blitter doubles it, so an odd origin lands one cell left and up of the unhalved origin. The origin is `(L09808, L09809 + 0x3c)`, with the display words giving `(width - 640) / 2` and `(height - 480) / 2` (`VIDEO-072`): 0, 80, 192 and 0, 60, 144 for `-640`, `-800` and `-1024`. All are even, so no value of those words yields an odd origin.
- **Interval.** `SmackOpen` (`D00070` to `D00071`) checks the `SMK2` magic and the allocations, compares buffer sizes with `0x3000` and `0x2000` (`D00072`, `D00073`) and tests flags, and stores the interval by the rule of `VIDEO-073`. The open compares nothing against the interval. The ROM routines read here (`R2096` and the player step) read the handle fields width, height, frame count, frame counter, palette and `NewPalette`, never the interval. Executed in `EXP-0447` on header-patched copies: open succeeded for the shipped value and for eight patched values, 0, 1, -1, 100, -100000, 42949673, 2147483647 and -2147483648 (nine variants in all). No code path in the decoder's open or the player validates or bounds the interval.

**Confidence.** **High** for the palette timing on the 66 shipped movies in the ROM's call order (executed; the inference about the palette chunk is Medium), for the lack of source clipping in the two indexed modes and the in-range 16-bit outputs on synthetic memory (executed), and for the absence of any interval check in the open and the player (whole routines read, nine interval variants executed). **Unknown** whether `rom.exe` ever selects a 16-bit surface, and the blit's behaviour when the source address leaves mapped memory (the probe stayed inside its allocation).

## Settings, music lists and cutscene list

Registry values named here sit under `HKLM\SOFTWARE\1C\Allods` (TOWN-186). The dialogs that edit them are MENU-073 to MENU-076.

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-075 | The cutscene list (message `0x436`, class `L12929`) shows `cutscene.txt` rows 0..N, N the registry dword at `L12930` (floor 2); OK plays that row's `cutpaths.txt` directory. | High / Medium | ● active | [EXP-0457](../experiments/EXP-0457-settings/) |
| VIDEO-076 | Among the nine request owners of the EXP-0229 census, the 21 music tracks form nine literal lists, one 12-entry list shared by every mission; no direct reference copies the registry `SoundRandom` value to the player. | High / Medium | ● active | [EXP-0457](../experiments/EXP-0457-settings/) |
| VIDEO-077 | Shadows, Dynamic lighting, Object animations and Smoothing are 0/1 dwords defaulting to 1, cleared by command-line switches and replaced by a registry load that runs later; they persist with nine other option values. | High / Medium | ● active | [EXP-0457](../experiments/EXP-0457-settings/) |

### VIDEO-075

The campaign arm `L12931` handles message `0x436`; the message is posted at `L12932` by an arm whose own trigger was not traced. The arm does nothing when the byte `[L12933]` is 0. The byte's image initial value is 1, so the gate is open unless something clears it; the only store found by the 4-byte sweep is 1 at `L12934`, in the video-resource setup of the application start (`L07629`..`L12934`).

The arm fills a string array with `cutscene.txt` rows 0 to N inclusive, through accessor `R0668` on the table object `L11471` (loaded from that file at `L11469`). N is the dword `L12930`. It is read from the registry value named `Using VxD` (`L12935`), raised to 2 when lower (`L12936`, `L12937`), and written back at save (`L12938`). The arm then builds the dialog `L12939` (arguments `(0x11, 100, 30, 540, 450, ...)`, a 440 by 420 rectangle; MENU-077). The dialog's builder `L12940` adds a list box (`R1183`, id 2, initial row `[L12941]`) and a scrollbar (`R1185`, id `0x29b`).

On OK (`0x445`) the handler stores the selected row at `[L12941]` (`L12942`) and calls the cutscene routine `R1658` with it. That routine saves the settings (`L12943`), then forms the path `%s\%s\%02d.smk` from the video directory string `L07641`, the row's `cutpaths.txt` text (object `L11472`) and a part number from 1 to 99 (`L12944`, `L12945`). A part that cannot be opened is skipped; a part whose playback routine `R0719` returns 0 ends the sequence. When called with row -1, the routine builds `m%d` from the current mission number (campaign `+0x660`), compares it with each `cutpaths.txt` row (`R0751`) and, on a match at row k, plays row k and raises `L12930` to k when it is lower (`L12946`..`L12947`). **Medium:** N is then a high-water mark of the `cutpaths.txt` row of the mission movies played, with a floor of 2; the callers that pass -1 were not traced.

The two tables have 14 rows each on both roots (TEXT-100). The video directory string is `video4` or `video8`; the `-8x` and `-4x` tests (`L12948`..`L12949`) choose between them, and the rule was not traced further. Archive contents are VIDEO-074.

**Confidence.** **High** for the arm, the list source and bound, the floor, the persistence and the path format. **Medium** that the menu is the owner's media menu (it is the only list dialog that reads a movie table), and for the high-water-mark reading, because the callers that pass -1 were not traced.

**Unknown.** The poster of `0x436`. The caller set of `R1658`. What sets `[L12933]` besides `L12934`'s branch condition. Whether N can exceed 13: no clamp was found, and the scan raises it only to a matched row index.

### VIDEO-076

Request lists are literals in nine owner routines. The nine owners, their addresses and list sizes are taken from the EXP-0229 census (`request-sites.tsv`, one row per replace call), not recomputed by EXP-0457: the shop `R1315`, the map `R1317`, the menu `R0816`, the inn `R1318`, the school `R1319` (two tracks, order set by the school-state bit, VIDEO-MUSIC-008), the Town `R1320`, the second inn `L08084`, the character generator `R0909` and the mission owner `R2087`. The eight screen owners hold one or two tracks each; the mission owner holds the 12 tracks `B00`..`B11`. These nine lists contain 21 tracks, the number of files in `MUSIC.RES` (VIDEO-MUSIC-001). Every owner honours the `-nomusic` gate (`config+0x20`).

A mission therefore has one list, shared by all missions (VIDEO-MUSIC-064). The other selection path in the recovered code is the numeric music-candidate selector of text markup (VIDEO-MUSIC-009). The Sound Options list shows the player's current candidate bank (VIDEO-OPTIONS-057), so in a mission it shows the 12 mission tracks. Progression is VIDEO-MUSIC-007: ordinary mode plays a randomized order and advances at stream end, fixed-source mode repeats the chosen candidate. The player constructor sets ordinary mode (`+0x20 = 1`).

The stored `SoundRandom` value is `config+0` (default 0). Only four instructions in the image reference `L03126`: the registry load (`L03128`), the registry save (`L03129`), the static initialiser (`L03127`) and the Sound Options constructor call (`L03130`). None copies `config+0` to the player, so the value changes the player only when the Random Order checkbox is clicked (MENU-076); the checkbox is built unchecked from the default 0 while the player starts in ordinary mode.

**Confidence.** **High** for the nine owners and their list sizes as the EXP-0229 census states them, and for the four references of `L03126` (`sweeps.py`: a 4-byte scan of the section bytes, plus the call-site scans for `R1230`, regenerated by `regen.sh`). **Medium** that no other per-mission selection exists: the claim is bounded by the nine owners and the markup selector. **Medium** for the unchecked-checkbox observation, which assumes no computed access to `config+0` (none was found among the direct references).

**Unknown.** Case handling and the exact identity of the title lookup (`Mid(6)` then lowercase, then a dictionary at `L12950`). The consumer of player `+0x68` (MENU-075).

### VIDEO-077

The 13 option dwords are read by `R1983` and written by the mirrored save `R1984`, with `RegQueryValueEx` through the options object `L06413`: GameSpeed (campaign `+0x3f4`), FormationMode, WimpyMode, ShowAllHitPoints, ShowFlyingHP (party object), Smoothing `L04367`, ShowTimeFlow `L06260`, TipsMode `L03631`, Acknowledgement `L06442`, AutoCasting `L06258`, Shadows `L06416`, Lighting `L06417` and Animation `L05659`. The load has no range clamp and ignores a missing value, so the in-memory default stays.

Defaults: the options constructor `R1845` (static initialiser `R1846`) sets Smoothing, ShowTimeFlow, TipsMode, Acknowledgement and AutoCasting to 1; the data section holds 1 for Shadows, Lighting and Animation; the campaign constructor sets the speed level to 4 (`L09328`). A checkbox yields 0 or 1 (MENU-073); the consumers of the graphics flags are TOWN-GRAPHICS-459 and TOWN-SMOOTH-460.

The command line is searched for substrings after the options object is built (`L12951`..`L12952`). `-nodynamiclighting` clears Lighting, `-noshadows` clears Shadows, `-noanimation` clears Animation, `-detail2` clears Shadows, `-detail1` clears Shadows and Lighting, `-detail0` clears all three, and `-nomusic` clears `config+0x20`. The registry load runs later, at `L12222`, so a stored value replaces a switch. The strings `-window`, `-safevideo`, `-640`, `-800`, `-1024`, `-8x` and `-4x` also exist; only `-8x` and `-4x` were traced (VIDEO-075), so no resolution or windowing option is claimed.

The sound configuration `L03126` (constructor `R0658`, static initialiser `L03127`) holds SoundRandom `+0` (0), music, effects and speech volumes `+8`, `+0x10`, `+0x18` (default -700 each), their ranges `+0xc`, `+0x14`, `+0x1c` (5000 each), the music-available gate `+0x20` (1) and MusicEnabled `+0x24` (1; `L06447` is the playback enable). Its registry values are SoundRandom, SoundMusPos, SoundSfxPos, SoundSpeechPos and MusicEnabled (`L03133`, `R0660`). The slider position at the default volume is 3130 by the inverse in MENU-076. Other values of the key are `Using VxD` (VIDEO-075), `phonebooksize`, `comportsettings`, `lastprotocol`, `lastip` and `CD`.

The save runs in the campaign teardown (`L12953`) and before a cutscene (`L12943`), so a changed option reaches the registry at the next of those, not at the click.

**Confidence.** **High** for the name and storage pairing (the load pairs each name pointer with its destination), the defaults read from instructions and the switch effects. **Medium** that a registry value replaces a switch, which rests on the order of calls in one function, and for the save timing, whose two callers were read but not all paths into the teardown.

**Unknown.** Whether the data-section initial values of `L06416`, `L06417`, `L05659` are the only writers before the load. The effect of `-window`, `-safevideo`, `-640`, `-800` and `-1024`. The `L12954` path that also clears `config+0x20`.

## Music after a Load cancel and after a mission-start exit

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-078 | Leaving the failure-panel save dialog without loading makes no music call in the dialog; the Exit route stops the player at `R1302` and requests the menu list at `R0816`; a window close stops the player and requests no list. | High / Medium | ● active | [EXP-0459](../experiments/EXP-0459-start-exits/) |
| VIDEO-079 | After a zero return of `R0099`, no caller makes a music request in the arms read: the player stays stopped from the stop at `L12824`, and the `0x45c` arm in phase 2 does nothing when `frame+0x3b4` is empty. | Medium | ● active | [EXP-0459](../experiments/EXP-0459-start-exits/) |

### VIDEO-078

- Cancel, Escape and Load with no selection end the dialog with `0x446` (`MISSION-067`). In the failure flow the session exists and `frame+0x414` is `0xff`, so the close arm posts `0x41e`: teardown `R1283` (stop through `R1302` and `R1233`, gated by `L12817`), then `0x421`, whose arm calls the menu list owner `R0816` only for a zero mask (`VIDEO-MUSIC-062`, `MISSION-DEFEAT-059`).
- From a session-less main menu the result posts `0x421` directly: the menu list is requested without a stop.
- The window close calls `R1322`, which runs `R1283` and then the MFC close; its closure reaches only the stop (`evidence/closures.tsv`). No list is requested.
- No closure from the dialog handler, its base, the key slots, the modal panel loop or the panel constructor reaches a music routine.

**Confidence.** High for the arms and the stop. Medium for the menu list on the cancel route (mask value inferred, `VIDEO-MUSIC-062`) and for the closure's absence clause: the closures hold functions with computed calls that were not read (`evidence/closures.tsv`).

### VIDEO-079

- `R0099` stops the player at `L12824` when `L12817` is non-zero, before the wait loop. None of its zero exits requests a list (`VIDEO-MUSIC-064`).
- The `0x42f`, `0x467` and LOAD callers drop the zero and post nothing (`MISSION-068`). The `0x457` arm shows the panel and posts `0x45c`. In phase 2 that arm does nothing when the CString at `frame+0x3b4` is empty; when it is non-empty the arm posts `WM_CLOSE`, whose handler runs the stop `R1283` and closes the application (`MISSION-067`). In phases 0 and 1 it runs `R1302` and posts `0x454`, `0x455` or `0x452`; those arms were not read for music.
- Direct closures from the panel constructor, the modal loop, `R1301`, `R1303`, `R1149`, `R0509`, `R0912`, `R0514` and `L06708` reach no music routine.

**Confidence.** Medium: the closures hold functions with computed calls that were not read (`evidence/closures.tsv`), and the phase 0 and 1 follow-on arms were not read.

**Unknown.** Whether a later message after the zero exit requests a list, for example from the in-game menu while the mission screen stays up.

## Shared SFX delivery

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-SFX-085 | The 112 recovered ROM1 direct SFX requests use distinct sample selectors, priorities, repeat modes and effects/speech gains; 13 priorities depend on pre-setting attenuation. | High / Medium | ✔ promoted | [EXP-0473](../experiments/EXP-0473-sfx-delivery/) |
| VIDEO-SFX-086 | ROM1 admits a nonplaying sample buffer before choosing a global SFX channel; equal priorities cannot evict, and sample destruction releases buffers that priority eviction retains. | High | ✔ promoted | [EXP-0473](../experiments/EXP-0473-sfx-delivery/) |
| VIDEO-SFX-087 | The inspected ROM1 music streaming and movie decoder routes own playback separately from the shared SFX request allocator. | High | ✔ promoted | [EXP-0473](../experiments/EXP-0473-sfx-delivery/) |

### VIDEO-SFX-085

R0386 receives a sample object and volume, pan, repeat, priority
low byte and optional frequency. The current repaired direct-call population
and raw .text E8-to-target scan agree on 112 addresses in 74 owners.
The recipes bind every terminal site to these arguments and a sample
source: numeric registry, class/action, voice bank, owner field or dynamic
speech path. TOWN-347, TOWN-493, TOWN-504, VIDEO-SFX-060, ANIM-094,
ANIM-128 and ANIM-119 supply reused selectors and membership.

Among those sites, 79 priorities are 128, 17 are 220, three are 100 and
13 are trunc((10000-abs(D))/100)&255. D is attenuation before adding a
volume setting. The effects/speech setting therefore does not change the
priority on those 13 routes. With D from -10000 to 0 the priority is 0..100;
this conditional arithmetic is not a native saturation witness.

Seven sites request repeat 1: river, firewall, inn water, the shop loop
wrapper, crowd, and school Fight1/Fight2. The other 105 request repeat 0.
Twelve sites use speech gain and 100 use effects gain. Sixteen pass
positional pan and 96 pass zero. All 112 recovered setups pass frequency
zero. Filename prefix alone cannot classify gain: shop start.wav uses
speech gain, while school Command1..3 and Fight1/Fight2 use effects gain.

The formerly unparsed stack sites L12955 and L12956 both use priority
220. Their intermediate array lookup consumes one selector push before
the service. At the latter site the slot 4/1 stores fall through to slot 2.
The ordinary and loop wrappers R1701 and R2104 have ten and
one recovered encoded callers respectively. The loop caller supplies the
shop InShop.wav sample. Ninety-six literal loader-helper bindings in
14 owners support the owner-field paths. Dynamic selection remains a rule,
not a statement that every possible selected sample exists.

Hurt and fall use the established ANIM-094 selector; bank replies reuse
ANIM-119. Their metadata is extended without reopening service membership
or replacing the corrected animation timing claims.

**Confidence.** High for the finite direct-call population and instruction-bound
metadata. Medium for closure over all possible routes and runtime-built
receiver/path associations. Ghidra 12.1.2, the repaired original project and raw PE
Capstone 5.0.7 decoding retain unnamed/raw candidates rather than filtering
them by a function name. Encoded direct calls and literal stored words are
finite populations, not computed pointer or bulk-copy closure.

**Unknown.** Native audibility, latency, restart cadence, device failures,
malformed spatial arithmetic, live bank/registry selection, dynamic sample
existence and routes through computed pointers or copied pointer-bearing
structures remain unobserved.

### VIDEO-SFX-086

The named startup path calls R0576 with count 16. That count sizes
the global channel array and the buffer array of each successfully loaded
sample. R0386 requires a device, a loaded sample and its buffer array,
then scans for the first duplicate without the playing status bit. A sample
with all duplicates playing returns before R2047, regardless of its
requested priority. Checked file/open or initial buffer-creation failure
ends the request before channel competition.

R2047 selects the first empty or nonplaying global channel. When
all are busy it selects the first channel at the lowest priority strictly
below the request; equality cannot evict. No eligible channel ends the
request. Priority eviction calls R2046 to stop and rewind the victim,
then clears its channel buffer, sample pointer and priority. It retains the
sample's buffer ownership.

An admitted request records buffer/sample/priority, clamps volume only below
-10000 and requests volume, pan, optional nonzero frequency and playback.
R1919 maps repeat zero to Play flag 0 and nonzero to flag 1. The
inspected refusal branches store no pending request; future caller retries
are separate. Device-call results in this playback path are not checked.

With a device and a loaded sample, R2105 stops every playing duplicate
and clears each matching global channel's buffer pointer. Its priority/sample fields are left until
the allocator normalizes an empty channel. R2106 releases every
duplicate, frees the array and clears the sample's array/loaded fields.
Replacement R2045/R2107 and town cleanup R1913 first
stop/rewind the channel found by R1476, then destroy the whole sample.
Stopping the first returned channel is not the whole cleanup contract.

The decoded cleanup slices in R0809 and R1301 Stop every
non-null global channel buffer and clear buffer, sample and priority fields.
Those channel loops contain no rewind call. They then unload registry samples
14 and 16 through R2106. No claim that every branch reaches these
slices follows from this local reading.

**Confidence.** High for branch order, the named startup count and the
stop/clear/release consequences in the decoded routines. The sample scan
before the allocator excludes treating these as a single admission pool.
Separate release and retained-buffer paths exclude uniform cleanup.

**Unknown.** Native success, saturation behavior, allocation or duplicate
creation failures, driver HRESULTs, teardown after device loss and arbitrary
malformed state remain unobserved. Unchecked duplicate creation is not a
guarantee of graceful refusal. Eviction itself supplies no loop restart;
later requests depend on the caller.

### VIDEO-SFX-087

Music R2108 creates a buffer owned by the streaming player at +9c
through the global sound device. The inspected body uses neither the SFX
channel array nor R2047. R2086 starts its own buffer with
repeat flag 1 under the music enable setting. R1233 and the music
volume path manage that player's state independently. VIDEO-MUSIC-011
supplies the streaming contract.

Movie construction R2100 supplies argument zero to the decoder's
sound-backend setup. The inspected game-side open/close calls enable sound,
close the movie and disable decoder sound without calling R0386.
VIDEO-081 and VIDEO-082 supply the already established decoder buffer clock
and stop/release contract; no new decoder implementation conclusion is needed.

**Confidence.** High for these named original game-side routes and their
buffer/decoder ownership. This excludes assigning them a fabricated SFX
priority or assuming their media track loops use the sample allocator.

**Unknown.** Native coexistence, physical device competition, audible
transitions and computed bypass routes outside these named bodies remain
unobserved. Separate ownership does not prove independent hardware resources.

## SFX spatial inputs

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| VIDEO-SFX-088 | The reached ROM1 positional SFX helper derives listener centre from the map view, computes exponential distance attenuation and signed horizontal pan, and returns stored attenuation D. | High / Unknown | ● active | [EXP-0474](../experiments/EXP-0474-sfx-spatial-inputs/) |
| VIDEO-SFX-089 | All 13 recovered direct positional-helper calls pass drawable fine world x/y and their associated view; three picture callers centre message cells and picture 51 passes newly updated x/y. | High / Medium | ● active | [EXP-0474](../experiments/EXP-0474-sfx-spatial-inputs/) |
| VIDEO-SFX-090 | The reached ROM1 ambient owner quantizes cell distance into exponential weights, averages weighted pan by source count, takes maximum river/fire attenuation and uses bird/crow population gain. | High / Unknown | ● active | [EXP-0474](../experiments/EXP-0474-sfx-spatial-inputs/) |
| VIDEO-SFX-091 | The named ambient owner scans after origin movement or a strictly expired unsigned deadline; existing river/fire buffers are updated or stopped and rewound according to their matched populations. | High / Unknown | ● active | [EXP-0474](../experiments/EXP-0474-sfx-spatial-inputs/) |

### VIDEO-SFX-088

The thiscall helper R0574 receives fine x, fine y, outD and outPan. Its
only view reads are dwords origin+5c/+60 and span+64/+68. Listener fine x/y
are `(origin<<8)+(span<<7)` using 32-bit shifts/adds. Source minus listener
is converted by FILD before floating-point squaring. For ordinary finite,
nonoverflowing inputs, its source arithmetic is:

```
pan = clamp(trunc(dx*4000/(width*256)), -10000, 10000)
D = trunc(max(-10000, -(exp(sqrt(dx*dx+dy*dy)/256/8)-1)*100))
```

Float64 constants 4000,256,8,1,100,-10000,10000 are read at
L12957..L12958. Pan converts before its signed clamps; D lower-clamps
before conversion and has no separate upper clamp. The reached L12959
thunk reaches FLDL2E, FMULP, FRNDINT, F2XM1 and FSCALE, establishing exp.
The qword FISTP conversion at R0279 temporarily selects truncation and
restores the incoming control word. It returns the stored qword's low dword.
The helper stores D and returns it in EAX, then RET 16; no success boolean
is established. VIDEO-SFX-085's 13 priorities use that pre-setting D.

Ten fresh Unicorn 2.1.4 original-instruction controls at explicit CW027f
execute the original sqrt, exp and conversion bodies. With origin 8/span 15,
eight-cell horizontal radius gives D=-171, pan=2133; the vertical control
keeps D=-171 and gives pan=0. One cell gives -13/266; eighty cells clamps
to -10000/10000. These are synthetic controls, not native observations.

**Confidence.** High for this instruction-bound arithmetic, input reads,
constants, conversion placement and return value. Current original PE decoding
with Capstone 5.0.7 plus actual original-instruction execution excludes the
old log10 label, axis-only distance, actor listener and screen-position input
models for this helper. Only the arithmetic label in ANIM-SND-022 is retracted;
its hooks, selectors, throttle and existing confidence remain in force.

**Unknown.** Native entry x87 control/precision, libm exceptional/overflow
paths and native bit-identical results; malformed spans, integer overflow,
out-of-range conversion, native audibility and device success.

### VIDEO-SFX-089

The raw .text byte-by-byte E8 rel32 scan finds 13 candidates to R0574.
All 13 are decoded calls in the existing request population; none is discarded
for missing function attribution. Three view-dispatch calls at L12960,
L12961 and L12962 use the incoming view and new drawable +8/+c. Their
input message fields are byte +d/+e, byte +a/+b and packed word +e low/high
bytes respectively, each expanded to `(cell<<8)+128`. The message pointer
comes from R0513 on that named dispatcher path. Exact player event labels
are not inferred from these field offsets.

Nine class/action/hurt/voice calls at L12963,L12964,L12965,L12966,
L12967,L12968,L12969,L12970,L12971 read drawable +8/+c and its
view +e0. Picture-51 phase-8 call L12972 passes newly updated +8/+c;
its z update is separate and not passed. Every terminal association is in
the experiment's finite consumer table and VIDEO-SFX-085's request recipes.

SESS-VIEW-028/029 bind the reused cell-origin fields, viewport span division
by 32 and fine-coordinate division by 256. For odd span 15, the positional
listener is at origin+7.5 cells. This helper does not consume the screen
projection of a drawable or its z field.

**Confidence.** High for all 13 concrete argument and terminal associations;
original operand decoding resolves field widths and centering. Medium for
closure over all possible request routes. This is a finite direct-call census,
not an image-wide writer/alias or lifecycle census. Computed calls, copied
pointer-bearing structures and runtime message reach are not excluded.

**Unknown.** Exact live semantic trigger coverage of the three message arms,
intervening drawable position writes, dynamic receiver lifetime and indirect
helper reach beyond the decoded population.

### VIDEO-SFX-090

Ambient R0484 uses integer-cell listener
`originX+sar(width,1),originY+sar(height,1)`. It scans inclusive bounds
`max(origin-span,8)` through `min(origin+2*span,mapDimension-8)` on each axis.
Terrain word source is view+80 owner's +c; object byte source is its +14.
River kind `((word&0x1fff)>>6)` is 8..11. Firewall matches flag 8 of a
packed-cell entry under view+9f0. Due object bytes index definition+40 after
subtracting one: >=0 adds a bird, ==-2 a crow, other negatives neither.

For each matched cell, original integer products form squared distance before
FSQRT. It truncates that radius and then divides by 8 toward zero. With
`q=trunc(trunc(sqrt(int32(dx*dx+dy*dy)))/8)`, its ordinary-input weight is
`A=max(0,trunc(10000-(exp(q)-1)*100))`. Raw pan is signed
`clamp(trunc(dx*2000/width),-2000,2000)`; each source contributes
`trunc(int32(rawPan*A)/10000)`. Group pan is the truncating sum divided by
source count, including zero-A sources, not by the sum of A.

River/fire have independent counts, pan sums and maxima. Their term is
`max(A)-10000`. Bird/crow share count and pan sum; their term is
`min(trunc(count*1000/30),1000)-2000`. Individual bird/crow counts retain the
existing category selection. VIDEO-SFX-085 supplies effects setting,
priority 220, repeat and frequency values for the three terminal requests.

Fresh original-instruction ambient controls at CW027f give river term 0 at
seven cells and -172 at eight. Sources at horizontal distances 7 and 16 give
term 0 and pan 1402. The same bird/crow positions give term -1934 and pan
1402 for count 2. This distinguishes quantization, maximum from summed
attenuation, count averaging from weight normalization and population gain.
The positional eight-cell D=-171 uses different truncation placement.

**Confidence.** High for the named owner's instruction-bound arithmetic,
classification, counts and grouping. Source execution uses synthetic maps
and explicit services, independently of the positional controls.

**Unknown.** Native x87 precision/control and libm edges, overflow, malformed
dimensions/class tables, arbitrary map population, native audible mix and
device output. These formulas do not prescribe exceptional conversion behavior.

### VIDEO-SFX-091

Ambient R0484 enters its scan when origin+5c/+60 differs from cached +3f6c/+3f70
or unsigned deadline+3f74 < clock. Equality is not due. Only movement writes
the cached origin. Width/height change alone does not satisfy movement.
River/fire run on either predicate; bird/crow object scanning requires due.
Nonempty bird/crow stores `clock+trunc(rand()/2)+10000` after its request;
zero count leaves deadline unchanged.

River/fire without a playing channel and with nonzero count request their
sample with repeat 1, priority 220 and frequency 0. A playing channel with
nonzero count calls volume with `clamp(term+effectsSetting,-10000,0)` and
pan on its non-null buffer. With zero count it calls Stop then position 0 on
that buffer. These local paths do not destroy the sample or clear the channel.
They do not establish a future retry or restart policy.

Eighteen fresh ambient controls include strict deadline equality, movement
without due, due without movement, nonempty playing update, empty playing
Stop/rewind, zero-weight count, count-30 gain ceiling, extent-only change and
the upper/lower existing-buffer volume clamps. Device calls
are captured services; the original owner instructions decide all operations.

**Confidence.** High for these named predicates, writes and requested buffer
operations. No global caller/lifetime census is claimed.

**Unknown.** Native clock wrap cadence, later caller retries/restarts,
intervening world/view updates, device HRESULTs and actual audibility.
