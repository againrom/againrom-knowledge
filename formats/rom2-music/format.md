# ROM2 music

ROM2 plays music from one archive, `music.res`, through one streaming
player. Screens request fixed keys; a mission requests a 17-entry list
and a map's type-12 music areas choose entries within it. These contracts
cover the EN and RU `allods2.exe` client; runtime audio was not observed.

## Archive

EN `music.res` and RU `MUSIC.RES` are identical. They hold 21 keys at the
archive root: `b00.wav`..`b16.wav`, `chrgen.wav`, `credit.wav`,
`map.wav` and `menu.wav`. Every entry is a 44-byte RIFF WAVE header and
one data chunk: PCM format 1, 2 channels, 22050 Hz, 16 bits, block align
4. Durations run from 38.973 s (`map.wav`) to 137.989 s (`b13.wav`).
— R2-ASSET-083

The client requests keys with the prefix `music\`; the mission list uses
upper-case `B00`..`B16`. Key lookup is therefore inferred to ignore case.
— R2-ENGINE-335

## Names and Music.ini

`main.res` `text/tunes.txt` maps each key to a display name, one
`key=name` line per key. The sound panel shows a list entry's name by
dropping the `music\` prefix and looking the rest up in this map. The
file names `credits.wav`, not the archive key `credit.wav`, so the
credits entry has no name. — R2-ASSET-083, R2-ENGINE-340

Root `Music.ini` has `[Global]` with `Count=17` and `Melody1`..`Melody17`,
the EN names of `b00`..`b16`. No game executable or library names the
file; the map editor does. No game image contains `music.ini` or the key
token `melody` in ASCII, ignoring case; UTF-16 and fully composed names
were not searched.
— R2-ASSET-083

## Screen keys

| Screen | Message | Key | Same key already set |
|---|---|---|---|
| Main menu | 0x421 | `menu.wav` | resume |
| Pre-create (New Game) | 0x425 | `chrgen.wav` | resume |
| Town 1 / 2 / 3 | 0x42e, 0x468 | `b14.wav` / `b16.wav` / `b15.wav` | resume |
| World map | 0x42d, 0x41d | `map.wav` | restart |
| Credits | 0x428 | `credit.wav` | restart |
| Mission | mission entry | list `B00`..`B16` | set again; see Mission music |

"Resume" starts the paused player without changing the list. "Restart"
sets the list again and plays an entry from its first byte. Each request
does nothing while music is unavailable. — R2-ENGINE-335, R2-ENGINE-336

## Mission music

1. Mission entry stops the player, sets the list `B00`..`B16`, opens
   entry rand() mod 17 and clears the area pick to -1.
2. If the hero exists, area select runs at once, before start, and then
   every 16 mission ticks on the hero position.
3. Area select takes the nearest type-12 area whose radius·256 exceeds
   the distance to (x·256, y·256); the head record at (0,0) is the
   default when no area holds the hero. Areas with all themes -1 are
   ignored.
4. The pick is theme[(rand()·4)>>15], drawn again while it is -1. It is
   a list index: theme t plays `b<t>.wav`.
5. A pick replaces the random entry before start, and later picks take
   effect at the end of the current entry.

Nothing but a new list set clears the pick. On a campaign map, whose
head record is admitted, every area select leaves a pick, so every
mission entry is a theme of the area holding the hero or of the head;
the random entry plays only when no hero exists at mission entry. On a
map without an admitted head the list plays in order from its random
start until the hero first enters an area. Closing certain in-mission
views sets the list again. EN head preservation is established below;
clearing between missions remains Medium because the full cleanup
contexts are unread.
— R2-ENGINE-336

The EN pointer-array helpers preserve the admitted head before appending
area pointers. The EN initialization stores counter zero; the tick
increments before its mask test, so without another writer the first
periodic check is tick 16. The counter store's containing body is unread.
In mode 2 the position-object selector tests actor word +0x1dc == 0x15;
in any other mode it compares actor +0x14 with view+0x9cc and requires
flag 1 at +0x1b8. The source-container callee `L2.05014` has an unread
body. The position object's hero identity and unit, and clearing on
every transition, remain Medium. The two cleanup loops
are listed, but their complete event contexts are unread.
— R2-ENGINE-834

The six close-tested windows are EN's generic/druid/Kaarg shops and
inns. Their close branches rebuild the mission list only outside
campaign mode 2. The entry selectors have a branch for modes other than 2;
no open was observed. Continuation of the square's b14/b16/b15 inside
those rooms is Medium: their selected
entry bodies make no direct music request, but their virtual descendants
were not fully traced. Other room routes remain Unknown.
— R2-ENGINE-835, R2-ENGINE-838

The type-12 record shape and the shipped areas are in
[ALM type 12](../rom2-alm/format.md#music-areas-type-12).
— R2-ASSET-084

## Player

The player streams the open entry into a looping DirectSound buffer. At
an entry's end it opens the area pick when there is one, else the held
entry, else the next entry of the list. A one-key screen therefore
repeats its key; the mission list advances in order only while the pick
is -1. Stop pauses the buffer at its position; start resumes it.
— R2-ENGINE-337

Selected paths stop music for the startup logos, movies, mission entry
and mission end; a later screen request resumes it. EN's 0x451 arm stops
only when its readiness helper returns nonzero and application+0x404 is
zero. The EN decoded-operand sweep of `.text` finds seven `push 0x451`
sites: `L2.05019`, `L2.05020`, `L2.05021`, `L2.05022`, `L2.05023`,
`L2.05024` and `L2.05025`; the last two are view-close routes.
RU was not swept; ids held in
data or built in registers were not searched. The linear sweep can lose
instruction alignment after data. The first site posts to a window whose
forwarding to this arm is unread. Complete event identities remain Medium.
The 0x41d, 0x45c and 0x45d arms reach mission end under different mode
and app+0x404 bit-0 tests. The 0x41d failure branch posts menu 0x421 in
mode 2; modes 1/3 post 0x459 and mode 0 posts 0x452. Mission end calls
stop only when the music-available word is nonzero. These messages are
not interchangeable pause commands.
A stream read failure closes the application. — R2-ENGINE-338,
R2-ENGINE-836

A dialogue tag `tune=N` cross-fades to list entry N and sets a hold flag
that only the random flag clears. While it is set, each of the stops
above restarts the held entry instead of stopping, and a track end
repeats the held entry only while the area pick is -1. No shipped file
uses the tag. — R2-ENGINE-338, R2-ENGINE-341

## Fade and player state

The dialogue tune and EN state-1 music-off paths call fade(2000,8000).
The first argument caps an effective time T in milliseconds; the second
sets the positive magnitude F of the negative volume threshold -F.
The buffer uses hundredths of a dB under its DirectSound identification.
T is capped by source and buffered-time calculations too; the exact
remaining-source interpretation is Medium because its virtual getters
were not fully read. — R2-ENGINE-839

The timer interval is
`delta=min(100,(floor(bufferBytes/bytesPerSecond)*1000)>>4)` ms.
The player timer callback reads timeGetTime. In EN, the head accepts a
callback when now >= due and advances due by delta once, and the tail
applies the step and stop test. The code between them
(`L2.05031..L2.05032`) and the RU head are unread, so a step on every due
callback is Medium. The step uses the stored delta.
With captured initial volume V0, positive T and F and
ordinary inputs without product overflow, each step computes:

```text
q = floor(F*delta/T)
next(v) = v - trunc_toward_zero(
  q*(F+V0)/F * (0.3*(v+F/4)^4/(F/2)^4+0.9))
```

A negative base with exponent 4.0 follows the listed even-integer branch
of the pow entry; its half-constant is unread, so the fourth-power
interpretation depends on that constant. The timer tail calls stop when
the volume before a step is less than -F; equality still attempts a step.
After a crossing, stop occurs when the tail is next reached; composition
through the next due callback is Medium. No final write forces exactly
-F. Setup, step and tail arithmetic in EN/RU remain High; actual callback
cadence and audible duration remain Unknown. Zero-T handling, x87 edge
rounding and the native V0 = -8000 case remain Unknown: the model gives
zero step forever at that volume, so the stop test never holds.
— R2-ENGINE-839

States 1 and 2 are computed results: with a buffer and open source,
unpaused music returns 1 with fade clear or 2 with fade set. Paused wins
and returns 3. Start clears paused; fade sets its flag; ordinary stop
sets paused and leaves fade set. Held stop instead clears fade,
reselects the held entry from the beginning and restores configured
volume. These are positive writers, not a complete alias census.
— R2-ENGINE-840

List set with hold first reselects the old list's held entry through
stop, then selects a random entry of the replacement list. Hold and
the old held index remain, and the area pick becomes -1. In EN, the final
selection resets the buffer cursor to zero without a Play call;
immediate audio depends on the prior buffer state. Empty lists and an
out-of-range held index remain Unknown. — R2-ENGINE-840

## Town comparison and LOAD

EN's town music body has only IDs 1/2/3 selecting b14/b16/b15. Other
IDs append no key to its temporary empty list; a default track or silence
is not established. Its same-key comparison is byte-sensitive and
compares element zero. This does not establish archive lookup casing.
— R2-ENGINE-835

The EN LOAD entry writes application mode 2 and routes to mission
entry(1) or town 0x42e. The mission entry's mode-3 bypass is not a LOAD
flag; selected mode-3 writers initialize connection setup. If intervening
work preserves mode 2 and succeeds, a mission LOAD rebuilds B00..B16
and picks from the restored position, while a town LOAD requests its
town key with the same-key resume rule. Both composed playback results
are Medium: intervening initialization, actor restoration and town
transit were not fully traced. Cold LOAD audio/cursor remains Unknown.
— R2-ENGINE-837

## Settings

| Registry value | Meaning | Default |
|---|---|---|
| SoundRandom | random flag; clears a held entry | 0 |
| SoundMusPos | music volume, hundredths of a dB | -700 |
| SoundSfxPos | effects volume | -700 |
| SoundSpeechPos | speech volume | -700 |
| MusicEnabled | music on | 1 |
| Acknowledgement | unit acknowledgement voices on | 1 |

The values live under the client's HKLM key, are read once at startup
and are written at exit and after a movie. After `-nomusic` or a failed
`music.res` mount music is unavailable: screen requests, the mission list
and the mission tick are ignored. The sound panel's music-on command does
not test this; its effect then is Unknown. MusicEnabled gates start only.
— R2-ENGINE-339

Acknowledgement belongs to the options object of R2-ENGINE-311's load and
store, not the settings object, and is the Sound Options checkbox (Medium:
control 0x28, the only checkbox in the windows read). With it 0, or outside
application mode word 2, no unit voices a selection or an order at the 13
call sites of the speaker picker, which returns 0 under those tests. A
gesture slot called by other code (Medium: none found) is not gated by it.
The `P` key sends pick-up through a sender that calls the picker. Complete
EN/RU digit-group bodies contain no direct picker/click call; silence
through every indirect descendant remains Medium. In EN, rectangle-tested
portrait-panel branches send 0x40e/0x40f to a view whose handlers can call
the ordinary click routine while a command collection is nonempty. This
composed route is Medium, not every portrait click; full RU parity is open.
Gesture voices play at SoundSpeechPos. — R2-ENGINE-720, R2-ENGINE-806,
R2-ENGINE-869

The unit slot that plays the easy, hard and dead files has five callers in
the client message pump (Medium: some call encodings were not scanned), not
through the picker, and no test of Acknowledgement or the mode word guards
the calls. The slot plays the voice subject to its 1500 ms gate. The arm for
message types 0x6c, 0x6e, 0x6f and 0x70 calls it with `dead` when a unit's
health falls from above 0 to -10..0 and with `hard` while it stays there
(Medium that all four types reach the calls). The read state-2..4 and
state-5 arms contain no direct slot call; absence through their descendants
remains Medium. A
health update calls it with the unit type's sound when the health is
unchanged, makes no slot +0x74 call at -10 or below, and calls `hard` below
half health and `easy` otherwise. These tests use health before the update.
R2-ENGINE-804 is partially retracted only for its unsupported label of
`[L2.00450]` as another sound call: its target remains Unknown. The RU
application +0x3f4/+0x3bc/+0x3c8 condition is after the state-1 voices.
In RU, with mode != 2, the current-unit match, new state 4 and old state < 4, text 0x44 is inserted when any of application +0x3f4/+0x3bc/+0x3c8 is zero; all three nonzero skip it. Text 0x45 follows under the same preceding conditions, independently of those three field tests.
The field meanings remain Unknown; this is not a proven hurt-voice gate.
— R2-ENGINE-804, R2-ENGINE-721, R2-ENGINE-863

Both wrappers supply 0x73. In EN, the read builder writes unit +0x94 to payload +0xc and the ID to +0xa; EN serialization selects 0x6e/0x6f/0x70/0x6c for one/two/three/four mask bytes. The corresponding RU helper/table and builder relationships remain Medium, with those bodies unread.
Selected EN 0x73 senders follow subtraction from unit +0x94;
complete game-event identities remain Medium/Unknown. Selected EN voice
word writers include zero, copies, packet/server transfer, supplied
fallback NPC and server template reductions. They do not establish the
word of every creature or hire; RU server-template event parity is open.
— R2-ENGINE-863, R2-ENGINE-864

A sound object returns without playing when no sound
device is open, its file open fails, its load is incomplete, it has no
buffers or every buffer is playing; a failed load leaves the object unloaded,
so each play retries it. A gesture whose file is absent is therefore silent
in both roots (High for the exits; that nothing is logged is Medium, since the
callees are unread; Medium for failures inside the WAV parse). — R2-ENGINE-913

Composed (Medium): hires carry their ID 1..7 as the voice word. EN hires 1..6 play
only easy, hard and dead, and the EN catapult plays no voice. RU hires 1..6
also play select, attack, move, swarm, retreat and defend, never pickup; the
RU catapult plays only its generic voice's easy, hard and dead. The weakest
links are the hire creator and the catapult's generic voice, which bears on the
voice name only, not on the counts. The file counts come from the complete
`characters/` key lists of `sfx.res` and the `speech.res` statement; other
archives and the mounted lookup were not searched, and no `.wav` lies loose in
the install trees (not the save folders). The EN "none" is Medium for that
reason. Hires 8..20 are Unknown. — R2-ENGINE-914

## Shop speech

The known EN/RU speech grid has page getter 4, so only the middle shared
left-release +0xa4 call dispatches its speech method. Read EN stock and
page-2 grids use other methods. Full admission and receiver coverage are
Medium. The request requires mode 2, source page 5..8 and draw(100)<30;
it is not conditioned on an established successful purchase.
— R2-ENGINE-861

Pointer comparisons select `shop`, `shop_druid` or `shop_kaarg`. The
first recognized tag in 2..11,21..25,27..31,33..37 selects
`speech\<shop>\effects\<tag:02>.wav`; otherwise a nonzero last tag-42
value selects `books\<value:02>.wav`. Otherwise it forms
`speech\<shop>\s<kind:02>i<subtype:02>p<p>.wav`, with
kind `(u16(item+6)>>8)&15`, subtype `u16(item+6)&31`, p=draw(3)+1.
The draw uses `floor(r*N/32767)%N` for these small positive N, so the
core endpoint r=32767 wraps to zero. Core seeding and complete live-item
decoding remain open. The filename literals and three shop stems are read in EN; the same RU strings are Medium inference from the corresponding code shape, not independently read literals. High covers numeric operands and
attribute precedence in both roots. — R2-ENGINE-862

Each root's 491 itemname-record projection forms 852 distinct normalized
names, 283 present and 569 absent from its speech key set. This is not
the shop-admitted/mutated population and proves neither universal key
coverage nor 569 actual failed requests. — R2-ASSET-212

## Speaker list endpoint and sound availability

The EN/RU picker can choose index count when rand15=32767. The read grow bodies request no extra logical element for a sentinel. The initial EN zero-fill covers the supplied logical range. No initialized sentinel is established (Medium: the EN initialization tail and RU L2.05133 are unread). data[count] remains Unknown.
The accessor has no bounds check. The following DWORD, speaker pointer
and native outcome remain Unknown. Selected EN speaker-field masks
preserve 0x1/0x10, and extra +0x78 stores include an object pointer;
neither field sweep establishes every unit owner or alias writer.
— R2-ENGINE-867, R2-ENGINE-868

Sound load/play use an object pointer cleared on listed EN/RU startup
failures and release. The EN factory thunk reaches DSOUND.dll ordinal 1;
the external body and complete pointer lifetime remain open. Zero there
means no usable pointer on these paths, not an options-checkbox value.
— R2-ENGINE-866

The sound panel's music-on command resumes or restarts the chosen melody;
the EN music-off arm stops or fades the player. The panel also sets the
chosen melody, the random flag and the three volumes, clamped to -10000..0.
— R2-ENGINE-340

## Unknowns

Audible output, the complete 0x451 sender events, complete indirect
screen inventory, panel layout and RU tune-name build remain Unknown.
Both client byte scans find six/eight/ten rel32 set-list/start/stop
candidates. This is a byte census with Medium completeness; computed
pointers and other jump routes were not closed.
— R2-ENGINE-338, R2-ENGINE-340, R2-ENGINE-836, R2-ENGINE-838
