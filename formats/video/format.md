<a id="audio-and-video--public-functional-specification"></a><a id="cutscenes"></a>

# Audio and video resource contract

ROM1 selects music, sound effects and cutscenes through the resource resolver.
Playback state belongs to the game. Music and sound payloads use standard
audio formats; numbered movie payloads use Smacker. The bundled decoder's
internal ABI and implementation are outside this reference.

<a id="scope"></a>

## Resource families

| Surface | Resource | Game state |
|---|---|---|
| Music | `MUSIC.RES`, archive identity `music` | Candidate list, shuffled order, fixed-source mode, enabled flag and volume |
| Sound effects | Sparse numeric entries in `SFX.RES`, plus direct paths | Sample selector, volume, pan, loop/play state, priority and optional frequency |
| Cutscenes | `VIDEO4.RES` or `VIDEO8.RES`, numbered `.smk` members and same-basename `.reg` | Selected namespace, mission identity, scan index, fade/pan and stop state |

— VIDEO-MUSIC-001, VIDEO-MUSIC-002, VIDEO-MUSIC-003, VIDEO-MUSIC-004,
VIDEO-MUSIC-005, VIDEO-MUSIC-006, VIDEO-MUSIC-007, VIDEO-MUSIC-008,
VIDEO-MUSIC-009, VIDEO-MUSIC-010, VIDEO-MUSIC-011, VIDEO-MUSIC-012,
VIDEO-SFX-013, VIDEO-SFX-014, VIDEO-029, VIDEO-030

<a id="shipped-payload-shape"></a><a id="namespace-and-selection"></a>

## Music

Installed music members are stereo PCM RIFF/WAVE, 22050 Hz and 16 bits per
sample. Their payloads have no identified loop metadata. Looping and track
succession are playback state. — VIDEO-MUSIC-001, VIDEO-MUSIC-002

The program has fixed candidate lists for UI/game contexts and a separate
mission/background family. Ordinary playback advances through its configured
candidate order: a one-member list repeats, and a multi-member list advances
through its order. Random Order off uses the sequential list; on performs
one pairwise swap per candidate using two CRT random indices. Fixed-source
mode can retain one candidate and is separate from that checkbox.
The context-list assignment to named surface transitions retains Medium
confidence. — VIDEO-MUSIC-003, VIDEO-MUSIC-004, VIDEO-MUSIC-005,
VIDEO-MUSIC-006, VIDEO-MUSIC-056, TOWN-372

Archive mounting follows the resource resolver and working-directory rules.
A missing music archive is a recoverable resource condition. Enable/disable
and volume are independent settings. The player uses a streaming buffer.
— VIDEO-MUSIC-007, VIDEO-MUSIC-008, VIDEO-MUSIC-009, VIDEO-MUSIC-010,
VIDEO-MUSIC-011, VIDEO-MUSIC-012

<a id="public-receiver-contract"></a><a id="sample-selection"></a>

## Sound effects

The common play boundary receives attenuation/volume, pan, repeat state,
priority low byte and optional playback frequency. Sample selection and
volume category precede that call. — VIDEO-SFX-013, VIDEO-SFX-014,
VIDEO-SFX-015, VIDEO-SFX-085

Selectors include registry-backed numbers, unit class/action values,
spell/projectile values, ambient state and literal UI/voice paths. The
registry is sparse; a valid selector can have no installed sample. Ambient
effects use visible terrain/object state and scheduled replay.
— VIDEO-SFX-016, VIDEO-SFX-017, VIDEO-SFX-018, VIDEO-SFX-019,
VIDEO-SFX-020, VIDEO-SFX-021

## Shared SFX delivery

The [per-caller request recipes](sfx-requests.md) retain all 112 reviewed
terminal tuples and the 11 shop wrapper field/path bindings, with source kinds,
original navigation identifiers and explicit absent/computed-source boundaries.
— VIDEO-SFX-085, VIDEO-SFX-086, VIDEO-SFX-087

The companion also defines the positional fine-world/view-centre contract
for all 13 recovered helper calls and the separate integer-cell ambient
accumulator. Positional attenuation uses exp; the former log10 label in
ANIM-SND-022 is partially retracted. Native x87 precision and audible output
remain Unknown. — VIDEO-SFX-088, VIDEO-SFX-089, VIDEO-SFX-090,
VIDEO-SFX-091, ANIM-SND-022

The 112 recovered direct requests retain distinct metadata: 79 priorities
are 128, 17 are 220, three are 100 and 13 depend on attenuation before the
volume setting. The latter use the low byte of
`trunc((10000-abs(D))/100)`; with D from -10000 through 0 this yields 0..100.
The effects/speech volume setting does not change that priority on those
paths. Seven sites request repeat: river, firewall, inn water, shop ambience,
crowd and school Fight1/Fight2. The remaining 105 request no repeat.
Twelve sites use speech gain and 100 use effects gain; all supply frequency
zero. Dynamic sample identity remains a selector/path rule. The finite
direct-call census is not closure over computed or copied pointers.
— VIDEO-SFX-085

Filename prefix does not determine gain. Shop `start.wav` uses speech gain;
school Command1..3 and Fight1/Fight2 use effects gain. Hurt, fall and bank
replies retain their established sample selectors and use the metadata of
their dispatcher paths. — VIDEO-SFX-085

A request first needs a loaded sample and a duplicate buffer that is not
playing. If all that sample's buffers play, it returns before competing for
a global channel. The named startup supplies 16 for both array sizes.
The global allocator takes the first idle channel or, with all busy, the
first lowest-priority channel strictly below the request. Equality cannot
evict. No eligible buffer/channel ends this request; later caller retries
are separate. — VIDEO-SFX-086

Priority eviction stops and rewinds a buffer, clears its channel and retains
the sample's allocated buffers. Sample destruction instead stops every
playing duplicate, detaches matching channel buffer pointers, releases the
duplicates and frees their array. An empty channel's other fields are
normalized when it is reused. Repeat zero requests ordinary Play; nonzero
requests its loop flag. Playback-call return values in the inspected service
are not checked, so the contract states requests rather than native success.
— VIDEO-SFX-086

The two inspected global cleanup loops Stop their non-null channel buffers
and clear all channel fields without rewinding. They then unload completion
and failure notification samples. Other branch reachability remains bounded
by the named cleanup slices. — VIDEO-SFX-086

Music owns its streaming buffer and uses the music enable/volume path.
Movie audio belongs to the decoder backend and its open/close controls.
These named routes bypass the shared SFX channel allocator. Separate
ownership does not establish independent physical device resources.
Native mixing, latency, audible restart cadence, device failures and
computed routes remain Unknown. — VIDEO-SFX-087

## Character-generator sounds

The pre-create page restarts `level1..3.wav` on each press on a difficulty
button and `char.wav` on each press on a hero button, selected or not. OK or
Enter with a non-empty name, the Back amulet and Escape request `ok.wav` unless
it plays and then close the page; with an empty name OK and Enter do nothing.
The close stops `ok.wav` within the same call. — VIDEO-SFX-058

The detailed page restarts `+_-.wav` for each applied statistic step,
including a double-click's second click and the held-button repeat; a refused
step is silent. A skill press requests its class's member for that skill slot
unless it plays: slot 1 sword or fire, 2 axe or water, 3 club or air, 4 pike
or earth, 5 bow or astral. The school room requests the same members by slot.
— VIDEO-SFX-059

Every request uses the SFX volume, pan 0, no loop and priority 128. At most one
instance of each member plays; members share 16 channels and displace only a
lower-priority sound when all are busy. The volume option never suppresses a
request. The main menu buttons and the Hall of Fame OK also request `ok.wav`.
Motion, release, the right button and page opening play nothing. Audibility of
the closing `ok.wav`, the repeat cadence and requests from outside the searched
code remain Unknown. — VIDEO-SFX-058, VIDEO-SFX-059, VIDEO-SFX-060

<a id="resource-selection"></a><a id="sidecar"></a><a id="user-stop-behaviour"></a><a id="decoder-boundary"></a>

## Cutscene sequence

1. Select the low/high video namespace from startup/environment state.
   No general fallback to the other namespace follows selection.
2. Form a numbered movie path from mission identity and an increasing
   two-digit index. Missing or unopenable members can advance the scan.
3. Replace the movie extension with `.reg` to obtain its sidecar. Apply its
   initial position, frame-bounded fades and frame-bounded pans.
4. Supply the resource-backed movie to the decoder. Draw decoded frames into
   the game-owned destination and track palette/frame progress.
5. Stop the complete numbered scan on ordinary key-down, system-key-down,
   mouse-button-down, close or quit events. Natural completion and some
   open/construction failures instead continue the scan.

— VIDEO-029, VIDEO-030, VIDEO-031 (its former first-dot reading is
retracted), VIDEO-032, VIDEO-033, VIDEO-034, VIDEO-035, VIDEO-036

The decoder boundary needs resource input, frame dimensions/count, palette
changes, frame progression/timing and destination ownership. Sidecar fade/pan
state and user interruption operate around this boundary. The sidecar is a
[REG store](../reg/format.md); its records are not part of the SMK bitstream.
— VIDEO-045 (its allocation-selector clause is retracted), VIDEO-046,
VIDEO-047, VIDEO-048, VIDEO-049, VIDEO-050, VIDEO-051

## Cutscene presentation

Each frame step waits, checks that the output buffer is focused, arms and
applies the sidecar fade to the decoder palette, decodes, blits and advances.
The sidecar start and every pan step move the blit source origin, never the
destination; a fade scales the palette by a per-frame factor and reaches its
end value, within float accumulation, on the frame before its end frame. Records are consumed in index
order. — VIDEO-071

A movie is drawn doubled when twice its width fits the 640 wide output region
or twice its height fits the 360 high region; a 480 high movie takes its own
size as the blit region. The display size changes only the origin of the output.
The shipped doubled movies are the 320x180 ones. — VIDEO-072

The frame wait uses the sound playback position when an audio track is open and
on, and a timer in 10 microsecond units from the header interval otherwise.
Positive intervals are milliseconds, negative ones are the magnitude, zero
means no wait, and the player adds no floor or ceiling. — VIDEO-073

With a track open and on, the wait compares a DirectSound byte position,
interpolated by wall clock between cursor reads, with a byte target, and no time
or count bounds it. If no track opens, the timer paces the movie. Read, not
executed; no sound output was opened. — VIDEO-081

At the final frame the player blits and then closes the movie before it returns.
The close stops and releases each sound buffer, so queued sound is cut rather
than drained (read, not executed; no sound output was opened). — VIDEO-082

Read, not executed: a missing sidecar raises an exception, a negative record
count reads no record, and equal start and end frames leave a fade active and a
pan running to the last frame. No shipped sidecar does any of these. — VIDEO-083

The fade scales the current frame's palette. On synthetic memory the two indexed
blit modes do not clip their source rectangle and the 16-bit modes matched the
indexed ones for in-range origins. The decoder's open compares nothing against the
header interval. — VIDEO-084

Both roots hold 33 shipped movies with header flags 0, one 16-bit 22050 Hz
audio track with the compression flag set, mono or stereo, and intervals from -4000 to
-8333. — VIDEO-074

<a id="a-namespace-that-is-not-a-cutscene-surface"></a>

## School training pictures

The four school-training resource families belong to the school-room loader
and painter. They use fixed room-relative anchors; an active sequence
replaces the current still on the same side. These named paths do not pass
the frames to a movie/cutscene surface. — TOWN-427, TOWN-428

The classification covers the named literal/direct/stored-pointer and
field-to-draw paths. Runtime-computed filenames or targets elsewhere are not
excluded. Native visibility and audio remain unverified; the animated arms
make no direct sound request. — TOWN-434

<a id="unknown--bounded-areas"></a>

## Cutscene list, music lists and stored options (`VIDEO-075`...`VIDEO-077`)

The cutscene list is the only list dialog that reads a movie table. It shows the titles of
`cutscene.txt` rows 0 to N, where N is a registry dword with a floor of 2 that a played mission
movie is read to raise to its own `cutpaths.txt` row (the callers were not traced). OK plays the numbered parts of that row's directory
in order (`VIDEO-075`).

The 21 music tracks come from nine literal request lists in nine owner routines. The mission
owner's 12-entry list is shared by every mission; among those nine owners and the markup
selector no per-mission list exists. No direct reference copies the registry random-order value
to the player at start (`VIDEO-076`).

Shadows, Dynamic lighting, Object animations and Smoothing are 0 or 1, default 1, and persist in
the registry with nine other option values and five sound values. Command-line switches clear the
graphics flags before the registry load, which then replaces them (`VIDEO-077`).

## Unknowns

All OS/decoder error paths, arbitrary malformed media, hardware/driver timing,
runtime-computed sample/movie names and sidecar branches beyond the named
consumers remain unspecified. No container writer or new audio/movie codec
is defined by this behavioral contract. Standard payload production and
bundled decoder internals are separate subjects.

## Sound Options playback controls

The track list uses the current player candidate bank, not all archive members.
Titles come from the `main/text/tunes.txt` dictionary using normalized names
without their six-character resource prefix. Selecting a row changes the
selected index; Play separately enables music, reloads a different selection,
applies gain and starts. Stop disables music and invokes the current-state
stop or transition operation. Exact native title normalization and dictionary
miss behavior, fade audibility and registry lifetime remain Unknown.
— VIDEO-OPTIONS-057

Random Order changes the ordinary candidate permutation immediately. Disabled
uses 0..n-1; enabled performs n swaps with two CRT rand()%n indices each.
Without a playback buffer, the measured setter does not change mode, flag or
permutation. This is not a switch to fixed-source playback.
— VIDEO-MUSIC-056
