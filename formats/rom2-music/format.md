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
views sets the list again. Whether the area array keeps the head and is
cleared between missions rests on unread helpers (Medium).
— R2-ENGINE-336

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

Music stops for the startup logos, every movie, mission entry, mission
end and message 0x451, and resumes only when a later screen requests it.
A stream read failure closes the application. — R2-ENGINE-338

A dialogue tag `tune=N` cross-fades to list entry N and sets a hold flag
that only the random flag clears. While it is set, each of the stops
above restarts the held entry instead of stopping, and a track end
repeats the held entry only while the area pick is -1. No shipped file
uses the tag. — R2-ENGINE-338, R2-ENGINE-341

## Settings

| Registry value | Meaning | Default |
|---|---|---|
| SoundRandom | random flag; clears a held entry | 0 |
| SoundMusPos | music volume, hundredths of a dB | -700 |
| SoundSfxPos | effects volume | -700 |
| SoundSpeechPos | speech volume | -700 |
| MusicEnabled | music on | 1 |

The values live under the client's HKLM key, are read once at startup
and are written at exit and after a movie. After `-nomusic` or a failed
`music.res` mount music is unavailable: screen requests, the mission list
and the mission tick are ignored. The sound panel's music-on command does
not test this; its effect then is Unknown. MusicEnabled gates start only.
— R2-ENGINE-339

The sound panel's music-on command resumes or restarts the chosen melody;
music-off stops or fades the player. The panel also sets the chosen
melody, the random flag and the three volumes, clamped to -10000..0.
— R2-ENGINE-340

## Unknowns

Audible output, the identities of the in-mission views that restart the
list, the sender of message 0x451, the panel layout, and the RU
tune-name build remain Unknown. — R2-ENGINE-336, R2-ENGINE-338,
R2-ENGINE-340
