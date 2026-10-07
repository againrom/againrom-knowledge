<a id="town-reactions-and-tavern-interior-animation"></a>

# Town rooms, reactions and animation

Town pointer handlers and paint routines maintain click targets, entrance
state, ambient episodes, tavern selection and school-training pictures. These
contracts are conditional on the named handlers being reached. Physical-pointer
delivery and audio-device results remain Unknown.

## World-map Return input

The world-map key-down and character slots call the same progress helper
without testing the key value; each returns handled. Return is one admitted
value. The world-map key-up slot returns unhandled and changes no travel
field. These are local message-slot contracts, not a physical keyboard
observation. (`TOWN-485`)

At zero route progress, the helper writes nothing. At nonzero progress it
assigns route-count-plus-one to progress and Cross-frame-count-plus-one to
Cross. It reads neither selected mission nor destination or current position,
and it neither builds a route nor posts arrival. Even when the route has been
fully revealed, a remaining Cross animation takes this same assignment.
Completion and destination entry belong to later paint. Outward/homeward
coordinates do not change the helper's local rule. Native admission of the
supplied idle and ready-counter states remains Unknown. (`TOWN-486`)

The inspected campaign Return key-down and character handlers forward to the
root with the third parameter replaced by zero. The root offers keyboard
messages to focus first. A nonzero result stops the route; zero permits the
ordered child pass and then the root's own slot. Mouse capture does not
select that keyboard recipient. Live focus, child identity, physical
main/keypad Return translation, system keys and the campaign key-up route
remain Unknown. The instruction-bound input does not establish native paint
cadence or a completed transition. (`TOWN-487`)

## Tavern entry ordering boundary

The located campaign tavern helper calls the command 37 sender before its
computed view activation. On the reached server command 37 arm, Player+04
selects the Player and a miss bypasses stock construction. A match first
changes Player+38 for the refund and requests its event. That local ordering
does not establish the earliest operation after a city click: earlier UI
callbacks, transport, actual command category and active client key remain
Unknown, as does the instruction behind the reported N3 tavern failure.
(SAV-918)

## Selector and labels

Message `200h` can reach town vtable `L11548+4c`, wrapper `L11549`, then
`L11546`. Each delivered update samples the mask and writes `town+b4`;
there is no enter-only comparison. Mask bytes map to selectors
`80h→2`, `90h→1`, `a0h→8`, `b0h→16`, `c0h→4`; other values map to `-1`.
Selectors 1, 2 and 4 draw Shop, Tavern and School label pictures respectively.
A delivered blank-mask update removes those labels but does not reset active
shop/tavern/school frames or enable bits. Focus loss is not that proved update.
(`TOWN-399`)

## Click targets

A left click in the town view reaches slot `+54`, `R0704`, which samples the
picture `townmask.bmp` and tests nothing else: no drawn figure has a rectangle,
sprite bound or test of its own there. The right button and the other click
slots are default stubs. Five mask values act: `80h` tavern, `90h` shop, `a0h`
gate, `b0h` statue, `c0h` school. Every path returns 1, and a click on any
other pixel does nothing in that routine. The mask is 640x480 at 8 bits. Each
value forms one large connected region, and `80h`, `90h` and `a0h` also hold
single stray pixels elsewhere. While the tips option is on, its initial value,
the tip popup over (328,0)-(640,200) is offered a click first, with its text
area, Close button and checkbox. The list/body passes left-down below; the
checkbox consumes it and writes the tips flag, while Close consumes it and
posts `45ah` on release inside. The list consumes left-up; the checkbox passes
it below. Delivery of Close's command to the room depends on campaign child
`10h`. The popup message slot returns 0 for `100h`, `445h` and `446h` and
forwards the rest; this corrects TOWN-186's former forwarding clause.
(`TOWN-185`, `TOWN-186`, `TOWN-207`, `TOWN-473`, `TOWN-480`)

Ten painted figures overlap a region: tavern door and label (`80h`),
shopkeeper and label (`90h`), gate door and guards (`a0h`), statue star
(`b0h`), school fighter, mage and label (`c0h`). Horses 2 and 4 reach `80h` and
`c0h` with 7 and at most 167 pixels. Sign, weathervane, birds,
horses 1, 3 and 5, baba and dervish overlap none. Outside the tip popup, a
figure is clickable where it lies on a region and inert elsewhere. A region is
larger than the figures on it, so it is not their outline. (`TOWN-474`)

Each region posts one message: tavern `42bh`, shop `42ah`, school `42ch`, gate
`42dh`, statue `41fh`. Tavern and school open only while campaign state `+3dc`
is 0, and the shop while it is 0 or is 1 with phase word `+6bc` at most 1.
The gate and statue have no `+3dc` condition. The gate first asks the offer
helper: with nothing on offer it opens the dialogue `inn\mercenary\npc35` and
posts nothing, otherwise it posts. The statue opens the town menu. The four
room arms send `445h` first, which leaves the town view. The shop and the
school speak on entry when the campaign holds an offer for them. No separate
mercenary hall is reached from the town. (`TOWN-475`)

The reached school show tail tests offer count `campaign+64c`, reads the head
word through array header/data `+644/+648`, and builds
`training\npc34m<word>`. After normal builder return it registers that word
and removes array element 0, count 1, without testing the builder result.
The shop tail has the same order using `+630/+634/+638` and
`shop\npc31m<word>`. Earlier room initialization, resource exceptions and
native offer-array contents remain Unknown. (DIALOGUE-048)

Inside the shop the merchant panel tests four shelf rectangles and two prompt
rectangles, and the painted merchant is not tested in that panel. The school
view tests the mage panel in one state and the fighter panel in another through
their class masks. A click on one of the five skill icons of the shown panel
changes that icon's state, plays its sound and selects its skill. The school
slots read do not test the trainers, the diamond or the column picture; the
remaining constructor children have bounded figure routes with the tip popup
absent. The merchant passes left-down: it is inert only with the prompt arm
disabled or a proved prompt miss. An enabled prompt hit clears bit `80h` and
returns the held item, still passing down; live prompt state and geometry are
Unknown. Merchant left-up is consumed and acts only while an item is carried.
The school view consumes figure presses and acts only inside a class panel.
School left-up returns 0 except for the fighter's 608-pixel button-child overlap
at x 464..479, y 200..237, which returns 1 without a button action. These are
painter rectangles and assume no capture child. (`TOWN-478`, `TOWN-479`)

With the shop tip popup shown, its checkbox (204,258)-(352,274) intersects the
merchant figure and writes the tips flag on down. The base figure no-action
result excludes this known overlay action. (`TOWN-479`, `TOWN-480`)

## Entrance state

| entrance | arm and retained state | advance and end | conditional sound request |
|---|---|---|---|
| Shop | Selector 1: `rand()%100 > 95` sets bit 1 if clear. Sprite `+1d8`, frame `+1e0`, shipped count 30. | One increment per admitted hub; equality with sprite count resets frame 0 and clears bit 1. Leave does not stop the active cycle. | Sound latch `+ac=0` permits `SFX/Town/Shop/Enter.wav`; clears school latch. (`TOWN-400`) |
| Tavern | Selector 2 ORs bit 2 on every delivered update. Sprite `+160`, frame `+168`, count 10. | One increment per admitted hub; count equality resets frame 0 and clears bit 2. Another update can rearm. | Frame 0 before increment permits `SFX/Town/Point.wav` and cancellation calls for shop/school sounds. (`TOWN-401`) |
| School exterior | Selector 4 ORs bit 4 without resetting fighter/mage frames `+1cc/+1d4` or static directions `L11559/L11560`. Both sprites have 11 frames. | At endpoint 0, remainder 96..99 starts +1; at 10 it starts -1, otherwise endpoint wait. Return to 0 with -1 clears shared bit 4. The hub still calls the second helper that tick; either clear can pause both thereafter. | Latch `+b0=0` permits `SFX/Town/School/Point.wav`; clears shop latch. Leave does not itself reverse the figures. (`TOWN-402`) |
| Gate | The hub always calls the gate helper. If availability helper returns -1, select `door/T08.bmp` and return. Otherwise poll current pointer mask every update. | Inside selector 8 decrements `+1a0` toward 0, outside increments toward 8; `+19c` selects `door/T00..T08.bmp`. Re-entry reverses existing progress. | Latch `+1a4` changes request `GateUp.wav`/`GateDn.wav`. (`TOWN-403`) |

Guard animation is separate from the door: eight-frame sprite `+e4`, frame
`+e8`, direction `+ec`. Pointer updates over an unavailable gate set -1;
other recovered arms set +1. Its unconditional hub helper clamps at endpoints,
stops on overshoot, and uses `Guard1.wav`/`Guard2.wav` on direction-latch changes.
The same hub can request then release a guard sound. (`TOWN-403`)

In the town routines read, the guards start moving only when a pointer message
is delivered. The hub then moves the frame one step per admitted hub, at most
one hub per 68 ms. With the pointer on the gate and nothing on offer the sweep
runs frames 6 to 0 over seven hubs and clamps on the eighth. Any other message
sets +1, which reverses a sweep in progress from the current frame. Entry sets
frame 7, step 0 and latch 0 and reloads the sheet, so the guards start every
visit at rest. Leaving the view frees the sheet and sounds and writes none of
those fields. The town routines hold no timer, loop or return trigger for the
guards. (`TOWN-476`, `TOWN-477`)

All sound requests above are conditional calls to `R0386`, not guarantees
of audible output. Null pointers, status lookup, audio availability and indirect
buffer operations can prevent playback. Latch behavior is separate from the
animation frame. (`TOWN-399`, `TOWN-400`, `TOWN-401`, `TOWN-402`, `TOWN-403`)

## Room return, gate latches and dialogue

Town enter calls a synchronous repaint before queued pointer delivery. When
the view's `+18` bit `20h` is clear and more than 67 ms passed, that repaint
admits the hub with guard step 0. The first process paint only primes the
clock. Re-entry flags and native timing remain unmeasured. The room-close and
enter routines read make no window/capture/cursor/clip call; the executable's
cursor pump posts `400h`. This does not decide OS-generated pointer messages
after a room closes. (`TOWN-481`)

The 96 guard-displacement stores outside the town/merchant clusters belong to
33 owners. None is reached by the direct-call closure of 62 town roots and
567 functions; ten stores in six owners remain unattributed. Frames 0 and 1
of the eight-frame guard sheet show crossed halberds. In the measured closing
sweep a request starts at tick 1 and reaches frame 0 at tick 7. Guard1's
418.14-ms sample can remain active there only with mean intervals below
69.69 ms; Guard2's 356.96-ms sample cannot under the strict >67-ms hub gate.
The fresh-enter latch suppresses Guard1's first closing request. Actual audio
output and native hub intervals remain Unknown. (`TOWN-482`)

The gate offer helper returns -1 exactly when the main latch and all child
latches are zero. Hearing does not latch; acceptance does. Loading the next
main record clears its latch, ages children and deletes them at age 2, so a
retained latched child can still be offered. The last-main-win arm posts
`428h` and skips that town-entry route; its later effect and a loaded last-mission
town state are not established. (`TOWN-483`)

The unavailable gate opens the non-modal npc35 dialogue with one pager button.
Its voice starts when the panel is shown and stops at the next pager call,
including the last-page press. The town's `402h` repaint skips while the whole
campaign `+3dc` dword is nonzero; other paint sources are not closed. Outside
the button, the model passes 305120 of 307200 pixels to the town click slot.
A missing node throws to the default exception handler without constructing
a panel; the displayed error text remains Unknown. (`TOWN-484`)

## Clock and independent ambience

Paint admits a single hub when unsigned elapsed time is strictly greater than
67 ms. No catch-up loop exists in that block. Sign and fluger are independent
random triggers, bits `40h`/`20h`, with 10/8 frames and conditional Flag/Flugel
sounds. They are not shop entrance reactions. Birds and the other town wildlife
retain separate arming mechanisms. Stars belong to selector 16, outside the
four entrance selectors. (`TOWN-404`)

The hub has nine non-bird tested flag bits, invoking ten gated helpers because
bit 4 calls two. Gate and guard helpers are unconditional. This corrects only
the eight-count clause of `TOWN-158` (partially retracted). (`TOWN-405`)

### Bird episode

Bird state is split between the view and process statics. The view owns clock
`+b8`, group `+bc`, count `+c0`, a nine-pointer array `+c8`, three progress
words `+d8/+dc/+e0`, and active bit `80h`. Process statics `L11525/bc/c0`
own the once latch, admitted-hub clock and next episode delay. Paint arms only
when strict elapsed exceeds the retained 1000..2999-ms delay and bit `80h` is
clear. Entry clears the view flags and resets the view clock; the next arm
zeros all progress, while no direct exact-address static reset is known.
(`TOWN-415`, `TOWN-417`)

The loader binds nine `TownBirds/Birds1..9/sprites.16a` sheets, each 57
frames. An arm selects group 0..2 and count 1..3; composition uses
`array[group*3+i]` for `i<count`. Every admitted hub increments all three
progress words once while active. Paint draws each selected nonterminal bird
through the class-specific `.16a` frame-draw slot+18, then draws keyed
`Town_add.bmp`, including on
the final active paint. It clears bit `80h` only after every selected sprite is
terminal. Count 1 conditionally requests `Birds1.wav`; count 2/3 requests
`Birds2.wav`, both repeat0. (`TOWN-415`, `TOWN-416`)

### Horse, baba and dervish

The exterior loader selects one of five horse positions, one of four baba
positions and a different one of four dervish positions. It binds the chosen
horse's A1..A3 sheets at `+128`, baba's A1..A2 at `+fc`, and the selected
dervish sheet at `+150`. All installed horse sheets have 15 frames, baba A1
has 31 and A2 has 32, and dervish has 30. (`TOWN-439`)

| family | entry and activation | admitted hub progression | terminal / idle |
|---|---|---|---|
| Baba | selects A1, current -1 and delay2000..3999 ms; a strict elapsed paint test rerolls A1/A2 and delay2000..6999, then sets bit `200h` | one forward frame through the selected sheet | count 31/32 clears the bit and writes -1; retained sheet draws frame 0 until another delay |
| Horse | selects A1, current -1 and its own delay2000..3999 ms; a separate strict test rerolls A1/A2/A3 and delay2000..6999, then sets bit `100h` | one forward frame through a 15-frame sheet | count 15 clears the bit and writes -1; retained sheet draws frame 0 until another delay |
| Dervish | current0 and bit `400h` are set on entry | `(current+1)%30` | wraps without clearing; a null sheet writes -1 but retains the bit |

Every eligible own-paint runs the two delay tests; the admitted hub gives each
active family at most one step and has no catch-up loop. Horse selector `+144`
is retained but dormant at entry current -1, then overwritten by every horse
arm before its frame-specific sound gates. Baba and horse therefore combine
entry selection with independent paint-time reselection; dervish is
entry-active and continuously cyclic. (`TOWN-440`, `TOWN-441`, `TOWN-442`,
`TOWN-443`, `TOWN-444`)

The sound loader binds Horse2/Horse3/Horse1 to `+a0/+a4/+a8`. Non-null and
status 0 gates request Horse1 at A3 frame 1 and Horse2 at A1 frame 14, A2
frames 8/14 and A3 frame 14, repeat0 and priority128. No direct Horse3 request
receiver or baba/dervish-specific request is present in the bounded town
closure; indirect/computed and runtime audio results remain open.
(`TOWN-445`)

Leave and destruction release the three family assets and horse sound
objects; re-entry reloads positions/assets, restores horse and baba idle A1
state and arms dervish at 0. The families share only flag word `+208` and the
process-static admitted-hub clock. Exhaustive mask-selector outputs are
-1/1/2/4/8/16, so its default OR does not directly arm bits
`100h/200h/400h`; horse/baba delay arms and dervish entry arm remain separate.
(`TOWN-446`, `TOWN-447`)

The painter draws horse, baba and dervish last, in that order, after the main
picture, the bird layer, the door labels and nine static sprite fields. The
direct blit census of the town range is 16 painter sites and 2 bird-layer
sites. (`TOWN-489`) Every frame of a sheet has one size, and a sheet is drawn
at a view-relative table position with no per-frame offset. All 27 canvases lie
inside 640x480 from the view origin. Horse bounding boxes overlap baba in 5 and
dervish in 3 of 20 position pairs; baba and dervish never overlap at a
reachable pair. (`TOWN-490`) The view's `+8` and `+0xc` are left and top, set
once by the constructor from screen globals and zero on the default 640x480
configuration; the first pair of each position table is x, the second y. The
origin is constant after construction in the town class table, a bounded
result. (`TOWN-503`) The entry roll is a scaled quotient followed by a modulo
or mask, not a plain modulo. (`TOWN-004`, `TOWN-505`)

A horse or baba episode arms when strict elapsed time since its clock exceeds
its delay. Every hub step refreshes the clock, so the wait runs from the
previous episode's last step. The entry delay is 2000..3999 ms and later delays
are 2000..6999 ms; a pause longer than the delay restarts a running episode.
(`TOWN-491`) A family changes frame at most once per admitted hub, at least 68
ms apart. Horse frames 1..14 take at least 14 hub intervals, baba A1 and A2 at
least 30 and 31, and one dervish revolution at least 30. Delivered paint
cadence is Unknown. (`TOWN-492`)

Horse1, Horse2, Horse3 and Crowd are 22050 Hz mono 16-bit samples of 538.8,
44.7, 29.2 and 3664.6 ms. Horse requests use repeat 0 and priority 128; the
crowd request uses repeat 1. (`TOWN-493`) A horse request is tested on every
eligible paint while its frame gate holds and passes only when no buffer of the
sample plays, so Horse2 can start again inside one frame dwell. (`TOWN-494`)
No town-range code other than sound cleanup (called from entry reload, leave
and destruction) stops these buffers; the shared channel allocator can stop a
lower-priority playing sound. (`TOWN-495`) The horse volume is the effects
volume of the sound configuration, written by its initializer, the registry
load and the options slider. A priority-128 request takes an idle channel or
stops a strictly lower-priority sound and is otherwise dropped. (`TOWN-504`)

No `srand` call lies in the town range; three direct sites seed the process
random generator from the clock in sound initialisation, the scenario
constructors and the AI manager constructor, and the direct-call closures of
entry, painter and click handler reach none of them. That the town does not
reseed is Medium: computed calls were not followed. Town entry draws at least
five times in a fixed order. (`TOWN-505`) A sheet that
fails to open reports a fatal error naming its path and aborts the process; the
loader never skips or stubs it. (`TOWN-506`)

### Statue star and crowd

The star loader binds nine 64×44 `Town/stars/S00..S08.bmp` pictures at `+1ac`,
initializes current `+1c0` through `-1→0`, and stores S00 in selected pointer
`+1bc`. That pointer draws at view-relative `(340,288)` whenever non-null.
Selector 16 ORs bit `10h`; the admitted hub then advances S01..S08. A step
beginning at current0 conditionally requests `Stars.wav`, repeat0. The step
reaching current9 hides the pointer and clears the bit; it does not wrap.
(`TOWN-418`)

Each terminal star call increments process-static `L11528`. On the tenth it
resets current and the static to 0, but still ends hidden; a subsequent selector
arm begins the visible sequence again. Entry independently restores S00 and
does not directly clear that exact static. Thus the star is rearm-driven, not
an autonomous cycle. Alias/bulk reset and physical pointer delivery remain
Unknown. (`TOWN-419`)

Crowd has no reached paint/progression owner in the bounded 39-function town
range. Entry loads `Crowd.wav` at `+74` and conditionally requests repeat1;
reload, leave and destruction reach detach/release/zero cleanup. This is a
bounded sound-route result, not a whole-image absence of a visual crowd.
Bird and star share only the admitted hub after separate activation; crowd
bypasses it on the measured route. (`TOWN-420`) No node path of either root
contains the stem `crowd` except the sound, and the `TownBirds` directory holds
41 nodes in nine groups; a crowd under another name is not excluded.
(`TOWN-496`) Every direct call through slots `0x18` and `0x38` in the town
range is one of 18 sites, and no other town-range routine on the painter route
calls those slots; the hook's child walk paints the tip popup child
(`TOWN-208`), which was not examined for a crowd. A crowd drawn inside a named
picture or through another slot is not excluded. (`TOWN-497`) The 462
undisassembled bytes of the town range are padding and four switch tables, and
the range's child widgets and helpers hold no sprite draw; crowd imagery inside
the named pictures was not examined. (`TOWN-507`)

For mouse-family routing, a non-null `control+34` handler bypasses the child
broadcast, including after a zero return. Zero still permits the final vtable
slot decision. With a null handler, the broadcast runs. The keyboard `+38`
path differs: a zero handler result permits a child broadcast. This corrects
the identical-fallback clause of `TOWN-211` (partially retracted). (`TOWN-406`)

## Limits

Complete residence behavior depends on event delivery. No static conclusion
here proves that a stationary physical pointer receives no updates. Frame and
latch persistence through focus changes, screen reuse, untaken indirect calls,
or changing gate availability is not closed. School interiors are outside this
contract. (`TOWN-399`, `TOWN-402`, `TOWN-403`)

The room figure and popup routes assume the enumerated constructor/enter
children and no capture child at the press. The merchant's live prompt
rectangles and children added by other paths remain Unknown. The guard-store
closure leaves ten stores unattributed and does not exclude code-pointer or
bulk-copy paths. Guard audibility is conditional on timing and a successful
audio request; OS-side pointer messages remain Unknown.
(`TOWN-479`, `TOWN-480`, `TOWN-481`, `TOWN-482`)

## Tavern interior draw and clocks

The central child owns four interior sequences. Its painter first draws
`CenterArea.bmp`, then candle and cauldron, then a mode-selected tender episode.
Coordinates below are relative to the parent origin. Each sequence stores an
index at `+18` and a cached picture pointer at `+14`; drawing uses the pointer,
not a fresh index lookup. Loaded pointers must be non-null. (`TOWN-407`)

| series | child offset | loaded files | position | progression when the central painter runs |
|---|---|---|---|---|
| Candle | `1dc` | `candle/t0000..t0009.bmp` | `(160,48)` | Draw current, then share the strict >100-ms gate with cauldron; `(index+1) % (count-1)` visits 0..8. (`TOWN-408`) |
| Cauldron | `20c` | `cauldron/t0000..t0020.bmp` | `(420,160)` | The same admitted step visits 0..19. Neither series selects its last loaded entry through this cycle. (`TOWN-408`) |
| Tender breath | `23c` | `tender/breath/br0001..br0024.bmp` | `(240,152)` | Mode2 draws before one strict >83-ms forward step. Completion clears mode/index but retains the last cached picture. (`TOWN-410`) |
| Tender drink | `26c` | `tender/drink/dr0001..dr0040.bmp` | `(240,152)` | Mode1 draws before a bounded forward/reverse step. The top step immediately reverses to index 38; reverse completion at 0 disables the episode. (`TOWN-411`) |

Tender delay is `3000 + rand()/16`, range 3000..5047 ms from the measured
15-bit random result. Strict elapsed >delay arms mode 1/direction 1 for odd
delay, or mode 2 for even delay. That test does not require idle mode. The
>83-ms tender gate resets the same tender timestamp after one step; completion
chooses the next delay. Mode0 draws neither tender series. These are conditional
gates, not measured frame rates or uniform random probabilities. (`TOWN-409`)

Two consequences matter. A long paint gap can rearm a descending drink toward
ascent. A completed breath retains cached `br0024.bmp` while its index is 0;
the next episode first draws that cached last picture and then selects
`br0002.bmp`, unless a reload or another writer intervenes. No catch-up loop
is present in either measured animation gate. (`TOWN-408`, `TOWN-410`, `TOWN-411`)

## Tavern roster grid and number grouping

The roster child fills 18 rects of 48x64 in its constructor, bottom row
first. List position `i` is column `i%6` and row `⌊i/6⌋` counted up from the
bottom: `x = 176 + 48·(i%6)`, `y = 480 − 64·(⌊i/6⌋+1)`. Paint and hit test
address the rect at `0x60 + 0x10·i`. The ÷3 slot formula `TOWN-065` published
is retracted. — TOWN-467

Mercenary cells come first and talk-only cells after them; the hit test covers
both halves and stores the selection. The bound below which a selection is a
mercenary is the mercenary count, not the `InnNPC` size. — TOWN-468

Nineteen call sites in nine routines group a decimal string with commas: the
character generator's remaining points and stat cost, the hall of fame score,
the tavern button values and mercenary price, one item-grid money cell, the
shop grid quantity and price, the shop button values and the school widget
values. Grouping inserts `,` before each three trailing digits and keeps a
leading `-` or `+` without a comma after it (`-1000` → `-1,000`, `-999` →
`-999`). The tavern pool count, the hall of fame rank and the generator's
`"%s = %d"` text are drawn ungrouped; the character/unit panel calls no
grouping. Whether a second grouping routine exists is Medium. — TOWN-469

`graphics.res` ships 28 inn sheets per root: `HeroFighter`, `HeroMage`,
`Unit1`..`Unit15`, `Unit29`, `Unit30`, `Unit32`, `Unit41`..`Unit44`,
`Unit52`, `Unit61`, `Unit62`, `Unit64`. The sheets above 15 are keyed by npc
id; that 29, 42, 43, 44 and 61 serve talk cells is Medium. — TOWN-470

## Tavern lifetime and sound boundary

Entry reloads all four arrays and resets each index and cached picture to the
first entry. Timestamps, mode and direction reside at static addresses, not in
those arrays. Exact-address references to eight named globals all belong to
the central painter; this does not exclude indirect or overlapping writes.
Leave passes all four objects to a cleanup helper whose body is not closed
here. Actual release, destruction and cross-visit state remain Unknown.
(`TOWN-412`)

Conditional requests are steam (>10000-ms separate clock), chair and
`Town/Shop/Breath.wav` (breath arm), drink (index 30 on every paint, either
direction), glotok (reverse completion), water (parent own-paint, repeat1),
and enter (entry, repeat0). Other named requests use repeat0. A non-null sound,
status-zero result and further audio-buffer gates are required. Request
sites do not prove audible playback. (`TOWN-413`)

The local painter is time-driven. Its complete invocation route and cadence,
indirect interaction effects and all alias writers are not closed. The party
pickers and read central input bodies contain no direct interior-state reset;
their complete callee effects remain Unknown. The selected-character preview
is outside this contract. (`TOWN-414`)

## School training presentation

The four `movies/training/{mage,fighter}/{m,tr}` BMP series belong to the
school-room object. The school loader calls the only direct family loaders:
mage `tr0000..tr0022` fills pointer array `+e4/+e8`, mage `m0001..m0011`
builds sequence `+164`, fighter `tr0000..tr0018` fills `+198/+19c`, and
fighter `m0001..m0009` builds sequence `+218`. Both shipped roots contain
these four gapless groups, 62 BMPs per root. (`TOWN-427`)

The room painter gives each class one lower-area anchor. The mage side is
room-relative `(0,200)` and fighter side `(320,200)`. On each side an active
`m` bit selects the sequence's cached picture; otherwise the current `tr`
picture is drawn at the same anchor. The modes replace one another; they are
not layered and are not a separate movie surface. (`TOWN-428`)

| family | activation | admitted progression | terminal behavior |
|---|---|---|---|
| Mage `tr` | entry/current-class transition, bit 1 | one modulo-23 step under strict elapsed `>83` ms | wrap to index 0 clears bit 1 |
| Fighter `tr` | entry/current-class transition, bit 4 | one modulo-19 step under the same gate | wrap to index 0 clears bit 4 |
| Mage `m` | idle side, strict elapsed `>3000+rand()/10`, bit 2 | forward indices 0..10; terminal cached 10 hold; reverse starts from index 8 to 7 | reverse terminal 0 clears bit 2 before the `m` draw |
| Fighter `m` | independent idle side, same delay shape, bit 8 | forward 0..8; terminal cached 8 hold; reverse 7..0 | reverse terminal 0 clears bit 8 before the `m` draw |

`tr` starts after a class change and participates in the pending-transition
input block until its counter returns to zero and the column reaches that
class endpoint. Mage index at least 5 can start column step+1; fighter index at
least 6 can start step-1. Those starts conditionally request
`SFX\Town\School\Rotate.wav`. (`TOWN-429`)

Each idle armer requires both the `tr` and `m` bit for its own side to be
clear. A busy side refreshes its timestamp. With the accepted 15-bit random
result, `3000+rand()/10` spans 3000..6276 ms; the comparison is strict.
Arming sets direction 1 and sequence index -1. The shared paint gate admits at
most one step per side and has no elapsed-time catch-up. (`TOWN-430`)

At mage `m` ascent completion, the direction changes to reverse and samples
hold threshold `((rand()*20)/0x7fff)%20 + 20`, range 20..39. This is not
`20+rand()%20`, and no uniform distribution is established. The sequence
index is then set to 8 while the cached pointer remains terminal index 10. The
terminal stays selected during hold; the first reverse selection is index 7,
so return omits indices 9 and 8. Fighter samples the same threshold, retains
index 8 at its terminal and reverses normally to 7. An active same-side `tr`
forces an ascending `m` toward reverse rather than instantly removing it.
(`TOWN-431`, `TOWN-432`)

Entry rebuilds both sequences, resets the local flags/counters and primes the
two `tr` pointers at index 0. Leave and destruction call both family cleanup
routines. The idle clocks, random extras and hold counters are static state
and are not explicitly cleared by the read lifecycle bodies; object identity,
alias writers and cross-visit timing remain Unknown. (`TOWN-433`)

The room painter consumes the school object's fields. No direct `m`
arm/step/draw body requests a sound. Computed filenames and targets, other
school audio paths, delivered paint cadence, visible output and audible
results remain Unknown. — TOWN-434

## Tavern initial selection boundaries

An empty collected mercenary list sets the inn's selection to -1. Its direct
activation caption call and paint-time price refresh contain signed upper-bound
checks that do not exclude this negative index. The local condition is High;
preservation across intervening UI/resource calls is Medium. A nonempty NPC
list does not change the measured empty-mercenary selection store. (SAV-929)

Initial party selection is separate. Its helper returns the first primary
party index with bit 0x20, or -1, and activation then indexes the party array
without checking that sentinel. A live marked party entry is therefore an
additional local prerequisite. (SAV-930)

Native click continuation, resource lifetime, first paint and first-failure
attribution remain Unknown. — SAV-932

## Game Options consumers

OK exports the tested Game Options values; Cancel alone leaves them unchanged.
Smoothing, Shadows, Lighting and Animation have distinct flags. Committing
Animation off also turns Lighting off. The global AutoCasting control emits
an autohealing mode command, separate from selected-spell autocast.
— TOWN-OPTIONS-457

Autohealing modes No/Standard/Often (0/1/2) set mana-floor percentages100/50/0.
The actor floor is truncated maximum mana times percentage divided by 100.
The recovered heal request requires current mana strictly above the floor;
equality does not pass. Direct percentages3..100 are accepted by the command;
invalid values leave the prior percentage. Complete targeting and scheduler
rules remain separate from this threshold. — TOWN-AUTOHEAL-458,
SAV-PLAYER-028, HERO-MP-006

Shadows gates the named unit shadow pass without suppressing its body.
Lighting gates the named dynamic-light body independently of day/night.
Animation off freezes the tested scenery frame at 0. These are bounded
consumers; the full water, town, interface and moving-actor populations remain
Unknown. — TOWN-GRAPHICS-459

In the recovered backpack painter, Smoothing enables `spritesb.256` after
its `sprites.256` base. Opaque indexed boundary pixels half-mix destination and palette colours
in the original 16-bit surface, masking each shifted half before addition:
mask 0x7bef for RGB565 or 0x3def for RGB555. Clipped pixels remain untouched.
This is an extra sprite pass; native complete-screen equivalence and coverage
of every drawable family remain Unknown. The painter pairs the overlay draw
`vt+0x34` with the five-argument base draw `vt+0x18` (`SPR256-077`), so
`SPR256-OVL-014`'s `vt+0x14` pairing is narrowed; that claim's no-path-read
clause is superseded. — TOWN-SMOOTH-460, SPR256-OVL-014
