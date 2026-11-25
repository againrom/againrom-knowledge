<a id="rom2-character-file-a2c"></a>

# ROM2 character file (`.a2c`)

Every statement about a loader, writer or field below is read in EN `allods2.exe`; RU `allods2.exe`
and `a2server.exe` were not read.

A ROM2 character is a standalone file named `<decimal id0><decimal id1>.a2c`.
It is a magic dword and six sections, each with a 16-byte header, a per-section
scramble and a doubling checksum. Section A is the identity, name and unit
choice; B the client's copy of the Player kill-grade block; C the nine client
hotkey entries; D the worn items; E the stats, money, spells and skill
experience; F the bag (names Medium). — R2-SESSION-022, R2-SESSION-416, R2-SESSION-514,
R2-SESSION-515, R2-SESSION-516, R2-SESSION-517

## File

```text
u32 0x04507989
A  tag 0xAAAAAAAA  52 bytes
B  tag 0x55555555  byte-run coded, decodes to 2,560 bytes
C  tag 0x40A40A40  nine hotkey entries
D  tag 0xDE0DE0DE  worn-item list (name Medium)
E  tag 0x41392521  0x34 field bytes plus pad bytes
F  tag 0x3A5A3A5A  bag item list (name Medium)
```

The name is the two identity dwords from section A in decimal; the
character-creation routine fills them from a performance counter value.
— R2-SESSION-022

## Reader and writer

The EN reader accepts the sections in any order. It returns 0 on a wrong magic,
an unknown tag or a checksum mismatch, and 1 at the end of the file; a section
the file lacks is returned as null, and a repeated tag replaces the earlier one.
Each section lands in one static buffer. B is run-decoded with no bound on its
decoded length. The reader tests no stream read result and compares no section
length with its static buffer (A, C and B's decoded length are all unbounded).
— R2-SESSION-512

The EN writer first reads an existing non-empty file; if that read fails it
writes nothing. It then replaces each section its caller passes, keeps the
others, and needs an A from one of the two. It writes A, B, C, D, E and F in
that order, each only when present, with a fresh random 15-bit key per section
and random E pad bytes. Opening for writing happens before the A test, so a call
with no A on a missing file may leave an empty file (Medium). — R2-SESSION-513

## Section header

| Offset | Width | Field |
|---:|---:|---|
| 0 | 4 | tag |
| 4 | 4 | length of the stored body; for E the writer stores `0x34 +` the key's bit count, one less than the bytes written when key bit 13 is set, and the reader ignores it |
| 8 | 2 | zero |
| 10 | 2 | key, 15 bits |
| 12 | 4 | checksum |

— R2-SESSION-022, R2-SESSION-513

## Scramble and checksum

For every section except E, each stored byte `i` is XORed with bits 16 to 23
of a 32-bit register that starts as `key | key << 16`, shifts left one bit per
byte and has `key` ORed back in after every 16th byte. The checksum is
`c = c * 2 + byte`, 32-bit, over the plain bytes; for B the writer takes it over
the run-coded bytes. Only the last 32 bytes affect the result, so a verifying
checksum excludes error only there. — R2-SESSION-022, R2-SESSION-512,
R2-SESSION-513

## Section A

| Offset | Width | Field | Claim |
|---:|---:|---|---|
| 0 | 8 | identity pair `id0`, `id1`; the file name | R2-SESSION-022 |
| 8 | 4 | list key: the client list keeps a file whose value matches the application's; creation stores the application value; the server writer stores 0 | R2-ENGINE-1840, R2-SESSION-514 |
| 0x0c | 32 | name, NUL-terminated; the client splits it at a bar character into two strings | R2-SESSION-514 |
| 0x2c | 1 | flag; bits 6 and 7 pick the unit row `Start_MF` (0), `Start_MM` (0x40), `Start_FF` (0x80), `Start_FM` (0xc0) in the client loader and the creation template; the server loader's mapping was not read (Medium) | R2-SESSION-514 |
| 0x2d | 1 | unit byte +0x4b | R2-SESSION-514 |
| 0x2e | 1 | unit byte +0x23c | R2-SESSION-514 |
| 0x2f | 1 | client field; the server writer stores 4; the Esc write clears bit 0 | R2-SESSION-514 |
| 0x30 | 1 | client field; the rename stores N of a `:N` name suffix (1..16, else 0); the server writer stores owner byte +0xa44 | R2-SESSION-514 |
| 0x31..0x33 | 3 | not read by the client or server loader bodies read; Unknown | R2-SESSION-514 |

The server loader reads only +0x2c, +0x2d and +0x2e of A (Medium). The uses of
+0x2f and +0x30 after a client load are Unknown. — R2-SESSION-514

## Section B

`u32 outBytes`, then packets. Opcode below 0x80 copies that many literal bytes;
0x80 and above repeats the next byte `n & 0x7f` times. — R2-SESSION-022

## Section C

Nine entries, each `u16 kind`, `u16 value`, `u32 n` and `n` bytes: the client's
nine hotkey entries, written by the client and read back by the client loader.
Nine entries with `n` = 0 make the 72 bytes both measured files hold. The
server loader ignores C. Medium that the entries are the hotkey entries at every
load. — R2-SESSION-515

## Sections D and F

A 9-byte header with the payload length as `u16` at offset 7, then records:

| Part | Width | Field | Claim |
|---|---:|---|---|
| item word | 2 | item word +0x40; zero is an empty item; `(word >> 8) & 0xf` is the class that picks the item constructor | R2-SESSION-516 |
| flags | 1 | bit 0x80: low six bits are the count; else bit 0x20: count 1, low four bits are an effect-unit count n; else a `u16` count follows | R2-SESSION-516 |
| skipped | 4 | only with bit 0x20 | R2-SESSION-516 |
| effects | 2 per unit | type byte and value byte; types 0x29 and 0x2c..0x30 take two units and two more bytes | R2-SESSION-516 |

D holds 12 records (unit +0x74, +0x78 and slots 3..12 on the server writer), each
passed to unit slot +0x44 by both loaders; that they are the worn items, and that
F is the bag, is Medium. Only the server loader reads F. Both measured files hold 12 D records and no F
record. — R2-SESSION-516

## Section E

16 fields; a pad byte precedes a field when its key bit is set. The reader's
chain undoes the stored form in this order, each step using the already decoded
earlier fields; the writer applies the exact inverse. The checksum covers the
decoded 0x34 bytes.

| Offset | Width | Pad bit | Decode | Holds | Claim |
|---:|---:|---:|---|---|---|
| 0 | 4 | 0 | `^ 0x1529251` | Player +0xa48 | R2-SESSION-517 |
| 4 | 4 | 1 | `+ f0*5 + 0x13141516` | Player +0xa4c | R2-SESSION-517 |
| 8 | 4 | 2 | `+ f1*7 + 0xabcdef` | Player +0xa58 | R2-SESSION-517 |
| 12 | 4 | 3 | `^ 0x17ff12aa` | Player +0xa54 | R2-SESSION-517 |
| 16 | 4 | 4 | `+ f0*3 - 0x21524542` | Player +0x3c, money (Medium); 1000 in a new file only if that file is a new character's first (Medium) | R2-SESSION-517, R2-ENGINE-1853 |
| 20..23 | 1 each | 5..8 | b5 `+ low byte of (f4*0x11 + f0*0x13)`; b6 `+ b5*3`; b7 `+ b5 + b6*5`; b8 `+ b5*7 + b7*9` | writer: unit word less its bonus byte (+0xd4..+0xd7 on the server writer); loaders: the word; unit words +0x84..+0x8a, the four base stats (names Medium) | R2-SESSION-517 |
| 24 | 4 | 14 | `- 0x10121974` | spell mask: bit s for each filled book slot s >= 1 | R2-SESSION-517 |
| 28 | 4 | 13 | none | byte +0x78 of unit +0x1c4 | R2-SESSION-517 |
| 32..48 | 4 each | 9..13 | g0 `^ 0xdadedade`; g(i) `- g(i-1)*0x771` | unit +0x240..+0x250, skill experience | R2-SESSION-517 |

The client loader reads every field except +24 and +28. The meanings of Player
+0xa48..+0xa58 and of the unit +0x1c4 byte are Unknown. — R2-SESSION-517

## Loaders

Section use read in the two loader bodies:

| Section | EN client loader | EN server loader |
|---|---|---|
| A | required | required |
| B | never read | copied into Player +0x44 when present |
| C | read when present | never read |
| D | read when present | required |
| E | required | required |
| F | never read | read into the bag when present |

In the EN client loader body the local for B is only passed to the reader. A
read of the static buffer `L2.08755` by other code was not searched; that no
other path fills client +0x3f70 from a file is Medium (census of displacement
and immediate forms). The 0xba and 0xb3 receipts change the copy. — R2-ENGINE-1852, R2-SESSION-416, R2-ENGINE-1271

## Client writes

Each write passes only the sections below; an existing file keeps the others.
Events are Medium: the bodies and their section sets are read; which player
step reaches them, and the reading of the arguments of `L2.08764` and
`L2.06472` as key codes, are not.

| Event | Sections passed | Claim |
|---|---|---|
| result of the app +0x36c screen other than 0x446 (Accept 0x445 by R2-ENGINE-283 on the New Game route; the multiplayer route to this screen is Medium, R2-ENGINE-1840), mode not 2 | A, B, C, D, E | R2-ENGINE-1854 |
| rename | A | R2-ENGINE-1854 |
| key 0x1b of `L2.06472` when app +0x404 is 1 and the mode is not 2 | A, E | R2-ENGINE-1854 |
| hotkey assignment | C | R2-ENGINE-1854 |
| 0xb3 or 0xba receipt outside mode 2 | B | R2-SESSION-416 |
| two pump arms of the application object (client by inference, Medium; message types Unknown) | D; F | R2-ENGINE-1854 |

Neither the creation routine nor its template contains a call of the writer or
of any routine that calls it; their callees were not followed, so that creation
writes no file is Medium. A new character's first file is the Detail-screen
write, holding A, B, C, D and E and no F (Medium). The
D bytes it carries come from a shared static buffer and are Unknown.
— R2-ENGINE-1853

A `BlCh` (`0x68436c42`) first dword marks a different variant whose flags at
offset 0x4c drive a cleanup scan; its writer was not found. — R2-SESSION-022

## Multiplayer list and server copy

The EN multiplayer character list reads every `*.a2c` of the current directory, checks the magic
and the checksum, takes the name from section A +0xc and keeps the file only when A +8 equals an
application value (+0x3d8, source Unknown; an A +8 of 0x529c0291 or 0xbaadf00d is first replaced
by +0x3d8). It orders the kept files by file time and appends one text entry; OK on that entry
runs the creation routine, which creates a new character in memory and writes no file (callees unread, Medium).
Medium that the entry means "new character". — R2-ENGINE-1840, R2-ENGINE-1853

The EN in-process server character writer runs only when Session +0x74 is nonzero; the map build
stores `0x84 > 1` there, the server init stores `mode < 2`, other writers are unclassified, and
whether a loose root map reaches the map build is Unknown (Medium). It writes to
`<base>chr\<owner +0xa70>\<id0><id1>.a2c` when the owner's +0xa70 string is non-empty, else to
`<id0><id1>.a2c`, through the same file writer. It passes sections A, B (owner +0x44), D, E and
F and no C, so an existing file keeps its C; the A and E fields are in the tables above. Medium that `<base>` is ChrBase. Which
events call the writer, besides unit death, is Unknown. — R2-ENGINE-1841, R2-SESSION-514,
R2-SESSION-516, R2-SESSION-517

## Relation to the save

The save's Player body carries a 2,560-byte raw block of the same size as
section B. The identity is inference (Medium for the object identity). — R2-SESSION-020

Section B is the client's copy of that block (client +0x3f70). Both receive arms that
update the copy (0xba: the whole run-coded block; 0xb3: one byte) rewrite section B alone
unless application +0x5d8 is 2, and character creation clears the copy. — R2-SESSION-416,
R2-ENGINE-1271

When a client uploads a character file (message 0xbe) for a Player without a hero, the EN server
copies the file's section B into Player +0x44 and sends 0xba, which rewrites the client copy.
— R2-ENGINE-1711, R2-SESSION-512
