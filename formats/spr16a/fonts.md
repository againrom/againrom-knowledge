# Font atlases and glyphs

[Reference](format.md)

## Font atlases — five of them (`SPR16A-FONT-007`, `SPR16A-FONT-013…SPR16A-FONT-015`, `SPR16A-FONT-020…SPR16A-FONT-022`)

The install ships **five** font atlases in two container forms (SPR16A-FONT-015), and
`rom.exe` loads **four** of them (SPR16A-FONT-018):

| file | form | glyphs | cell | advance range | notes |
|---|---|---:|---|---|---|
| `font1/font1.16` | byte-control | 224 | 16×15 | 0..14 | + an appended section chain |
| `font2/font2.16` | byte-control | 224 | 8×10 | 0..7 | + appended sections; **the one node the RU release replaces** |
| `font3/font3.16` | byte-control | 64 | 8×6 | 0..3 | 53 of the 64 records are empty |
| `font4/font4.16a` | **ordinary `.16a`** | 224 | 16×16 | 0..15 | decoded by the `.16a` path above |
| `font5/font5.16a` | **ordinary `.16a`** | 224 | 24×24 | 0..20 | **loaded by nothing in the install** |

Glyph record `k` is selected by display byte `32+k`. The four 224-record
atlases cover bytes 32..255; `font3` covers 32..95. Record 0 is a blank
space glyph. Use the separate advance table for spacing.
— SPR16A-FONT-007, SPR16A-FONT-018

### The high half — CP437 letters plus Cyrillic in a hybrid arrangement (SPR16A-FONT-020)

Records 96..223 of every 224-record atlas hold, under `char = record + 32`:

```
0x80..0x9A   CP437's accented Latin (Ç ü é … Ö Ü); its ¢£¥₧ƒ dropped (blank)
0xA0..0xA5   á í ó ú ñ Ñ; CP437's ª..» dropped
0xB0..0xCF   А..Я        — the box-drawing region carries the Cyrillic uppercase
0xD0..0xDF   а..п
0xE0..0xEF   blank except ß at 0xE1
0xF0..0xFF   р..я
```

The atlas arrangement is not CP866, CP1251 or KOI8-R. The RU `font2.16`
blanks its accent records at bytes `0x80..0x9a`, `0xa0..0xa5`, `0xe1` and
`0xef`; the other named atlases retain their EN forms. — SPR16A-FONT-020,
SPR16A-FONT-022

The five fonts share this arrangement on both roots: the 64 Cyrillic records (144..191,
208..223) have identical glyphs and advances in EN and RU. The RU `font2.16` and `font2.dat`
differ from the EN files only at 35 records (96..122, 128..133, 193, 207), each blank with
advance 0 in RU. The record that a typed byte reaches is stated in
[TEXT](../text/format.md#typed-byte-to-record). — SPR16A-FONT-020, SPR16A-FONT-022, TEXT-109

The no-conversion clause of SPR16A-TXT-023 is partially retracted. The
selector-dependent byte transform is specified by
[TEXT](../text/format.md#encoding-model); atlas indexing happens after that
transform. — SPR16A-TXT-023, TEXT-CONV-001 (partially retracted), TEXT-LANG-002

### A font is two nodes — `fontN.dat` is the advance table (SPR16A-FONT-018)

Each font object loads **`<base>.16`/`.16a` *and* `<base>.dat`**. The sidecar is one `u32`
per glyph (896 B = 224, or 256 B = 64 for `font3`) and text layout is

```
x += dat[glyph] + spacing          // spacing = 2, from the construction site
x += height(0)/2 + dat[0] + spacing   // for glyph 0, the space
```

The cell width is distinct from `dat[glyph]`. Use the advance value and
the space special case above for text layout. — SPR16A-FONT-018

### Standalone `.16` trailer

The final four bytes are read as one little-endian word directly into
`this+4`. The successful loader stores it unchanged. Two 2-bit masks (value AND 3)
elsewhere in the complete loader handle error-message string
copying before the trailer read; neither masks that word. — SPR16A-070

| Operation | Original `.16` expression |
|---|---|
| Stored word | Complete raw u32 |
| Resource allocation request | Full resource byte length |
| Pointer-table allocation request | 32-bit `raw << 2`, equivalent to `4*(raw & 0x3fffffff)` |
| Indexing entry | `int32(raw) > 0` |
| Indexing continuation | Signed 32-bit `index < raw` after increment |
| Cursor advance | `cursor + 12 + u32(cursor+8)`, in 32-bit address arithmetic |

The masking expression in the table is an arithmetic equivalence; the
instruction is a shift. Allocation calls precede indexing. Zero and
bit-31-set words skip indexing if the allocation call returns normally.
A positive raw value retains its complete logical bound even when upper
bits disappear from the allocation request. No native allocation result or
completed large-count traversal follows from these expressions.
— SPR16A-070, SPR16A-071

The loader rewinds to resource offset 0, copies the full resource including
the final trailer, and sets `this+0x20` to zero. The selected decoder adapter
loads `table[index]`, passes width, height, frame pointer plus 12 and the
caller's ramp, and does not read the stored word. Trailer bits do not select
a palette, payload origin or decoding mode on this path. — SPR16A-072

An original `.16` producer remains Unknown. A bounded search of the game
and EN editor suffix references, loader paths and selected class methods
identified no frame-record writer. Memory copying, diagnostic count output
and conversion to a system bitmap do not establish one. — SPR16A-073

### Text pixel composition — draw-time dispatch, two routines, not interchangeable

Which pixel-write rule a drawn glyph gets is decided at draw time by the
sprite object's own vtable, not stored per font file — but the two text draw
routines are not two equally-valid paths to the same result: one of them
crashes on a font4 sprite before producing a pixel, so each routine draws
exactly one sprite class in working code. — TEXT-065, TEXT-067, TEXT-071

- Font1-3 (`.16`-class) reach the opaque `.16` byte blitter above (no
  destination read) through the primary draw routine (`DrawText`). Through
  the second draw routine their equivalent vtable slot is a bare no-op (one
  instruction, `RET`) — font1-3 draw no pixels at all through that routine.
  — TEXT-065, TEXT-067
- Font4 (`.16a`-class) reaches a genuine blend against the destination pixel
  — the published `.16a` alpha compositor — but only through the **second**
  draw routine's own vtable slot (`vt+0x18` → `R1785` →
  `R1518`/`R1784`), not through `vt+0x34`/`vt+0x14`. — TEXT-067,
  SPR16A-FONT-094
- The primary draw routine's own font4 call site (`vt+0x34`) is not merely
  argument-count-short against its resolved receiver, it unconditionally
  dereferences a null pointer before any pixel work, on every call: the
  primary routine always supplies a literal zero in the exact stack slot that
  receiver reads as a pointer and dereferences. The two routines therefore
  split by sprite class as a matter of working code, not just as an
  observation: the primary routine draws `.16`-class fonts (font1-3) only,
  and font4 is drawn, where it is drawn at all, through the second routine.
  Which callers/screens actually invoke the second routine with font4 active
  remains open — no caller-side census was run. — TEXT-071

Neither draw routine writes a glyph more than once for an ordinary character
— there is no outline/shadow primitive inside the text-drawing code itself.
The one exception, drawing a `~` as a rule instead of a glyph, substitutes a
different pixel-producing call rather than adding a second one alongside the
glyph call. — TEXT-068

No surviving claim in this repository asserts anti-aliasing, supersampling or
sub-pixel filtering anywhere in the sprite/text pixel path, within the bounded
search this repository has run (three search terms, both draw routines' full
bodies); one claim that once did is retracted. Apparent smoothing is either
the pre-graded intensity levels already baked into the `.16` atlas, or font4's
destination blend above — not established as a filter applied at draw time
within that search. — TEXT-069

Colour reaches a pixel differently by font class. Font1-3 select among a
small, caller-chosen fixed set of ramps (above), built by a subsystem outside
both draw routines, not recomputed per call. Font4, drawn through the second
routine's receiver, takes a table-pointer argument: a non-zero caller value
overrides the table, a zero value falls back to a pointer stored on the
sprite object itself — not a level shift. (A level-shaped, bit-shifted
argument does exist elsewhere in this receiver family, in the code the
primary draw routine cannot safely reach for font4; it does not describe
font4's actual draw-time colour mechanism.) Neither path inserts a separate
text scaling stage; both compose directly into the active framebuffer's own
channel-mask format. — TEXT-066, TEXT-070

### `.16` glyph pixel grammar (SPR16A-FONT-013)

A `.16` glyph record has the same `[u32 w][u32 h][u32 dataSize][data]` shape, but `data` is a
**byte** control stream, not the `.16a` word stream:

```
control byte c:  op = c >> 6 ,  n = c & 0x3F
  op 00  literal   : the next n BYTES follow, each carrying TWO 4-bit pixels,
                     LOW nibble = the left-hand pixel. A ZERO HIGH nibble ENDS
                     the run and is PAD, not a pixel — the run is 2n-1 px long.
                     A zero LOW nibble is a WRITTEN pixel, value 0, not transparent.
  op 01  blank rows: emit n fully-transparent rows (at a row boundary)
  op 10  skip      : emit n transparent pixels in the current row
  op 11  SKIP too  : the dispatch tests 0x00 then 0x40, so 0x80 and 0xC0 share the
                     skip branch
```

- A complete glyph stream consumes its data and produces exactly `width*height`
  pixels/skips. Literal runs do not cross a row end. A final zero high nibble
  is padding; a zero low nibble is a written intensity value. — SPR16A-FONT-013
- Observed pixel values: `{4,5,6,7,8,9,11,13,15}`. **The 4-bit value is an intensity level of
  a text colour the caller chooses** (SPR16A-FONT-013). It is *not* a level into the `.16a`
  path's 16-level LUT — a `.16` object has no palette and never builds one — and *not* a
  palette index. The blitter indexes a **16-entry `u16` table passed as its sixth argument**
  (the 16-bit load indexed by twice the entry number, so the whole table is 32 bytes), and the engine builds thirteen of
  them, each entry `base * k / 15` packed to the active framebuffer format. Rendering the
  value as coverage is therefore not a display convention — it is what the format means:

```
pixel = ramp[v]        // ramp = 16 u16 entries chosen by the text's caller
                       // ramp[k] ~ baseColour * k / 15, packed to the framebuffer
```

- The write is **opaque** — unlike the `.16a` u16 blitter there is no destination read and no
  additive blend. A `.16` glyph replaces the pixels it covers.

### The `.16` section chain (SPR16A-FONT-014 partially retracted, SPR16A-FONT-021)

`font1`/`font2` do not end at their count trailer. Each continues into further
`[records][4-byte trailer]` sections, every trailer an identical copy of the first:

```
EN font1.16  224 recs 16x15 | e0000000 | (26 B) | 11 recs | e0000000 | (12 B) | 115 recs | e0000000
EN font2.16  224 recs  8x10 | e0000000 | (57 B) | 49 recs | e0000000
RU font2.16  224 recs  8x10 | e0000000 | (28 B) | e0000000          (font1.16 identical EN/RU)
font3.16      64 recs  8x 6 | 00000040
```

The named shipped files have positive counts and their indexed records
begin at the front. The signed count predicates are stated above. Bytes after
that frame list add no pointers to the selected frame table. The loader
nevertheless reads the full resource, and its selected copy constructor
copies the full stored byte length if invoked. Other runtime uses of the
tails, native invocation of that copy and visible rendering of the tails
remain Unknown. The named residual sections are older in-place overwrite
layers and cut record suffixes. They do not extend the indexed count or
define a new section grammar. — SPR16A-FONT-014 (partially retracted), SPR16A-072,
SPR16A-FONT-021, SPR16A-FONT-022, SPR16A-071
