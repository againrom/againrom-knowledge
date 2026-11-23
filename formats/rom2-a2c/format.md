<a id="rom2-character-file-a2c"></a>

# ROM2 character file (`.a2c`)

A ROM2 character is a standalone file named `<decimal id0><decimal id1>.a2c`.
It is a magic dword and six sections, each with a 16-byte header, a per-section
scramble and a doubling checksum. Section meanings are Unknown except the
identity pair, the name string and section B, the client's copy of the Player
kill-grade block (R2-SESSION-416). — R2-SESSION-022, R2-SESSION-416

## File

```text
u32 0x04507989
A  tag 0xAAAAAAAA  52 bytes
B  tag 0x55555555  byte-run coded, decodes to 2,560 bytes
C  tag 0x40A40A40  raw block
D  tag 0xDE0DE0DE  item list
E  tag 0x41392521  0x34 field bytes plus pad bytes
F  tag 0x3A5A3A5A  item list
```

The writer rewrites an existing file, replacing the sections its caller
supplies and omitting a section with no data. The name is the two identity
dwords from section A in decimal; the character-creation routine fills them
from a performance counter value. — R2-SESSION-022

## Section header

| Offset | Width | Field |
|---:|---:|---|
| 0 | 4 | tag |
| 4 | 4 | length of the stored body (E: its extent is `0x34` plus the pad count; the header length was 59 against an extent of 60 in one file) |
| 8 | 2 | zero |
| 10 | 2 | key, 15 bits |
| 12 | 4 | checksum |

## Scramble and checksum

For every section except E, each stored byte `i` is XORed with bits 16 to 23
of a 32-bit register that starts as `key | key << 16`, shifts left one bit per
byte and has `key` ORed back in after every 16th byte. The checksum is
`c = c * 2 + byte`, 32-bit, over the plain bytes. Only the last 32 plain
bytes affect the result, so a verifying checksum excludes error only there.
The scramble is supported by structure (B size, the id pair and file name, D
framing) plus that tail checksum. — R2-SESSION-022

## Sections

- **A**: `u32 id0`, `u32 id1`, a zero `u32`, a NUL-terminated string in
  `[0x0c, 0x2c)`, eight more bytes.
- **B**: `u32 outBytes`, then packets. Opcode below 0x80 copies that many
  literal bytes; 0x80 and above repeats the next byte `n & 0x7f` times.
- **C**: raw block, 72 bytes in both measured files, holding small nonzero
  values (0 to 19); its stored checksum is 0 and does not discriminate.
- **D, F**: a 9-byte header with the payload length as `u16` at offset 7,
  then records of `u16`, `u8` and a body. A zero `u16` is an empty record.
  Record framing was replayed on three non-empty records only.
- **E**: 16 fields at offsets 0, 4, 8, 12, 16 (`u32`), 20, 21, 22, 23 (`u8`),
  24, 28, 32, 36, 40, 44, 48 (`u32`). A pad byte precedes a field when its key
  bit is set. The first nine fields use bits 0 to 8, offset 24 uses bit 14,
  offset 28 uses bit 13 and offsets 32 to 48 use bits 9 to 13, so bit 13 pads two
  fields. Each field has a transform chain with constants and earlier fields;
  the checksum covers the decoded 0x34 bytes.

A `BlCh` (`0x68436c42`) first dword marks a different variant whose flags at
offset 0x4c drive a cleanup scan; its writer was not found. — R2-SESSION-022

## Relation to the save

The save's Player body carries a 2,560-byte raw block of the same size as
section B. The identity is inference (Medium for the object identity). — R2-SESSION-020

Section B is the client's copy of that block (client +0x3f70). Both receive arms that
update the copy (0xba: the whole run-coded block; 0xb3: one byte) rewrite section B alone
unless application +0x5d8 is 2, and character creation clears the copy. Whether loading
a character reads B back into the copy is Unknown. — R2-SESSION-416, R2-ENGINE-1271

When a client uploads a character file (message 0xbe) for a Player without a hero, the EN server
copies the buffer the reader returns in its fifth out-parameter, taken to be section B, into
Player +0x44 and sends 0xba, which rewrites the client copy (Medium that the copied buffer is
section B). — R2-ENGINE-1711
