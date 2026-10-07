<a id="spr16a-16a--16-sprite--specification"></a>

# SPR16A sprites and `.16` fonts

`.16a` uses the SPR256 frame container with u16-control RLE. `.16` uses a
separate loader plus byte-control glyph data. Its complete u32 trailer is
stored unchanged; table-allocation arithmetic and signed indexing predicates
use that word differently. — SPR16A-STRUCT-001, SPR16A-FONT-013,
SPR16A-070, SPR16A-071

A `.16a` literal selects a source shade-table entry and a destination blend
level. A `.16` literal contains two four-bit text intensities. Neither is a
literal framebuffer word. — SPR16A-PIX-011, SPR16A-FONT-013

The `.16a` source offset is the complete raw u16; the normal and low-memory
destination-table origins differ. Reversed clipping also depends on run
partition. The selected decoder contract is in [encoding](encoding.md).
— SPR16A-080, SPR16A-081

## Wire layout

| Part | .16a | .16 |
|---|---|---|
| Leading palette | 1024 bytes only when trailer bit 31 is set | None |
| Frame header | Three u32: width, height, encoded byte count | Same widths |
| Frame data | u16 controls; 14-bit counts; literal u16 lookup offsets | Byte controls; 6-bit counts; paired 4-bit intensities |
| Final u32 | Low 31-bit frame count and bit 31 palette flag | Raw stored word; allocation wraps at 32 bits and indexing uses signed comparisons |

The standard size is `paletteBytes+sum(12+dataSize)+4`. All multibyte
fields are little-endian. The original frame-indexing blocks do not check
frame extents. A `.16` word enters that block only when positive as a signed
32-bit value, conditional on the allocation call returning normally.
— SPR16A-STRUCT-001, SPR16A-TRLR-012 (partially retracted),
SPR16A-RDR-017 (partially retracted), SPR16A-FONT-013, SPR16A-071

The separate `.16` and `.16a` entry identities are verified. Their retained
reference enumeration covers the named repaired, disassembled project;
bytes never disassembled and other entry paths remain outside that result.
An exhaustive loader inventory remains Unknown. — SPR16A-RDR-017 (partially retracted)

## Read and write order

Select the codec from the resource contract and read the trailer. `.16`
copies the full resource from offset 0, keeps no palette, and supplies the
selected frame's dimensions and data pointer to its decoder; the trailer
does not select a decoding mode. Its allocation request and indexing rules
are detailed in [the `.16` trailer reference](fonts.md#standalone-16-trailer).
— SPR16A-070, SPR16A-071, SPR16A-072

The named positive-count fonts add no frame-table pointers for their tails.
The loader still reads the full resource, and its selected copy constructor
copies the full stored buffer length if invoked. Other runtime uses of
those tails remain Unknown. — SPR16A-FONT-014 (partially retracted), SPR16A-072

Decode .16a word
RLE through its palette/blend tables; decode .16 byte RLE through the
caller's text ramp and paired .dat advance table. To write the standard
form, emit the optional palette, frame headers and encoded data, then the
count/flag trailer. Retained overwrite tails have no general writer.
— SPR16A-RLE-002, SPR16A-FONT-014 (partially retracted), SPR16A-FONT-018

No original `.16` producer is established by the bounded game/editor
search. The emission description states the ordinary record layout;
original producer behavior remains Unknown. — SPR16A-073

## Reference map

| Reference | Contents |
|---|---|
| <a id="structure-16a-standard-form--spr16a-struct-001-542542"></a><a id="frame-pixel-data--16-bit-word-rle-spr16a-rle-002-spr16a-rle-003-27272727-frames"></a><a id="defensible-limits--a-decision-not-a-measurement"></a><a id="open-own-experiments"></a><a id="closed"></a><a id="sprite-container"></a><a id="word-rle"></a><a id="literal-word-colour--the-display-model-spr16a-pix-011"></a><a id="structural-bounds"></a><a id="unknowns"></a><a id="read-and-write-sequence"></a> [Sprite encoding and pixels](encoding.md) | Container, word RLE, blend, bounds and emission |
| <a id="font-atlases--five-of-them-spr16a-font-007-spr16a-font-013spr16a-font-015-spr16a-font-020spr16a-font-022"></a><a id="the-high-half--cp437-letters-plus-cyrillic-in-a-hybrid-arrangement-spr16a-font-020"></a><a id="a-font-is-two-nodes--fontndat-is-the-advance-table-spr16a-font-018"></a><a id="16-glyph-pixel-grammar-spr16a-font-013"></a><a id="the-16-section-chain-spr16a-font-014-spr16a-font-021"></a> [Font atlases and glyphs](fonts.md) | Byte glyph RLE, text slots and advance tables |
| <a id="cursor-sheets-spr16a-cursor-046"></a><a id="the-item-icon-subtree-spr16a-icon-027"></a><a id="projectile-sheets"></a><a id="spell-art-which-sheets-one-cast-can-reach"></a><a id="which-consumer-draws-each-spell-reachable-sheet-spr16a-part-030"></a><a id="frame-demand-of-the-mark-path"></a> [Cursor, item and spell resources](resources.md) | Cursor, inventory, projectile and spell sheet consumers |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.
