<a id="text--public-functional-specification"></a>

# Text bytes, tables and input

ROM1 stores one-byte strings. The active language selector controls input and
display conversion; font indexing follows conversion. Positional string
tables retain source byte order. Glyph framing and advances are defined in
[SPR16A](../spr16a/format.md). — TEXT-CONV-001 (partially retracted),
TEXT-LANG-002, TEXT-STRTAB-023, SPR16A-FONT-018

## Encoding model

ROM1 uses a one-byte text pipeline. The active language selector chooses whether the
Russian conversion rules are applied.

For display in the Russian mode, the byte transform is:

```text
0x80..0xAF -> byte + 0x30
0xE0..0xEF -> byte + 0x10
otherwise  -> unchanged
```

In the other selector state the display transform is the identity. The transformed byte
selects a font record by:

```text
record = uint8(transformedByte - 0x20)
```

This is a byte operation, not Unicode decoding. A compatible implementation may expose
Unicode internally, but conversion to/from the original resources must preserve the
original one-byte semantics. — `TEXT-CONV-001` (partially retracted), `TEXT-LANG-002`,
`TEXT-INDEX-003`, `TEXT-SEL0-012`

## Input conversion

In Russian mode the keyboard/input conversion into stored bytes is:

```text
< 0x80      -> unchanged
0xC0..0xEF  -> byte - 0x40
0xF0..0xFF  -> byte - 0x10
otherwise   -> unchanged
```

The game's Russian lowercase helper additionally maps the two Cyrillic uppercase ranges
before falling back to the ordinary single-byte lowercase operation:

```text
0x80..0x8F -> byte + 0x20
0x90..0x9F -> byte + 0x50
otherwise  -> ordinary single-byte lowercase
```

These transforms are byte functions; no multibyte encoding is involved.

The display transform is not injective in Russian mode: distinct stored bytes can map to
the same font record. A reimplementation should therefore keep stored text bytes and
rendered-glyph identity conceptually separate rather than normalizing the source data. —
`TEXT-DOM-010`

## Font indexing and bounds

The original font access path indexes records directly after the `-0x20` transform and
does not provide a general substitute-glyph/clamp rule. Therefore callers are
responsible for splitting control characters and supplying bytes valid for the chosen
font.

A compatible implementation may add defensive bounds checks for safety, but that is an
implementation divergence and must not be mistaken for a property of the original file
format.

Font record counts, cell geometry and sprite framing are documented in the SPR16A
specification rather than duplicated here. A font is two resources — the glyph sprite
set and a separate per-record advance table — so the drawn cell width is not the advance
a text measurer must add. — `SPR16A-FONT-018`

### Typed byte to record

A byte typed through a text field passes the input conversion and then the display conversion.
At the Russian selector a typed `CP1251` byte `0xC0..0xEF` reaches record `byte − 0x30`
(records 144..191) and `0xF0..0xFF` reaches record `byte − 0x20` (208..223), so the offset is
not one constant. A shipped string byte passes the display conversion alone. The 224-record
atlases carry the records and `font3`, with 64 records, holds none of the 64 Cyrillic bytes.
The letter at those records is the typed letter in font1 and font2, and by advance evidence
only in font4 and font5. — `TEXT-108`, `TEXT-109`

At the English selector a typed byte reaches record `byte − 0x20`. Of the 64 `CP1251` bytes
`0xC0..0xFF`, 16 draw their own letter, 32 draw another Cyrillic letter and 16 reach records
192..207, which are blank except `ß` (and one more glyph in EN `font2`). — `TEXT-110`

## Markup byte

The tilde byte has markup semantics in the font1-3 draw path:

- a doubled tilde represents one literal tilde glyph;
- a lone tilde draws no glyph and takes no pen advance. It underlines the next glyph: a line
  from the pen x to x plus that glyph's advance-table entry, on row y plus the height of the
  font's frame 0 (the row under the glyph cell), in entry 15 of the ink ramp, with both end
  pixels drawn;
- the measurer skips a lone tilde and counts a doubled one once.

A text measurer must mirror the draw path's markup handling or widths can diverge. The font4
draw has no tilde arm: it draws the byte as glyph 0x5e, so a font4 string that holds a tilde is
measured narrower than it is drawn. A tilde in the last byte of a string reads the advance table
one entry past its end. No shipped dialogue text or dialogue button label holds a tilde; each root's help text holds one
doubled tilde and no lone tilde (`TEXT-089`); other text was not searched. — `TEXT-TILDE-009` (underline extent partially retracted),
`TEXT-TILDE2-017`, `TEXT-079`

## Draw anchors

The font draw routines take a flag word of which only the values 1, 2, 4 and 8 are read. The
effects add:

| value | effect on the drawing origin |
|---|---|
| 1 | x moves left by the width of the string |
| 2 | x moves left by half of that width |
| 4 | y moves up by the width of the font's frame 0 |
| 8 | y moves up by half of that width |

The vertical anchors use the width of frame 0 where its height would be expected. Frame 0 of
font 1 is 16 wide and 15 high, so value 8 moves y up by 8. Value 10, the two centring values,
centres a string on x and puts the top of its cell at y - 8. The callers' other values are not
enumerated. — `TEXT-078`

## Resource string tables

[Pointer-hover help](hover.md) describes the common delayed display route,
control-specific sources, line layout and the bounds of the control inventory,
including the composed spellbook caption, the character-generation attribute
rows, the map-list columns and the binding of inherited hints.
— TEXT-HOVER-048, TEXT-HOVERSET-049, TEXT-HOVERCHAR-050,
TEXT-HOVERROOM-051, TEXT-HOVERTEXT-052, TEXT-HOVERPAINT-053, TEXT-080,
TEXT-081, TEXT-082, TEXT-083, TEXT-084, TEXT-085, TEXT-096, TEXT-097, TEXT-098,
TEXT-105, TEXT-106, MENU-066, MENU-067, MENU-068, MENU-069, MENU-096

The game loads multiple CRLF-delimited text resources into positional tables. Entries
are addressed by table-local or global numeric indices depending on the consumer.
Loading is byte-preserving; conversion happens when text is displayed or entered, not
when the resource file is parsed. — `TEXT-STRTAB-023` (both-root line total partially
retracted; loader and table structure retained)

Table rules:

- table ordering is semantically significant;
- inserting/removing a line in a positional table can renumber later entries;
- consumers may apply their own numeric formatting and visibility rules after selecting
  a string entry;
- path-like tables and UI-prose tables use the same basic storage mechanism but should
  not be assumed interchangeable;
- a consumer must not infer a language-independent semantic key from the displayed
  wording alone.

## Character-name entry

The original character-name control stores a bounded one-byte string and applies the
input conversion before appending accepted bytes. Backspace is handled as editing rather
than a stored character; control bytes below the printable range are rejected in the
reached input path. — `TEXT-NAMEIN-024` (vtable address, slot index and timer clauses
partially retracted)

The pre-create screen's name field opens at `npcnames.txt` entry 20 (table-local,
counted from 0): EN `Danath`, RU `Данас`. On every traced opening the screen copies its
stored name into the field after replacing `Unnamed` or any of entries 20..23 with entry
20; a new campaign stores `Unnamed` unless the command line carries `-name`. A return
from the detailed page keeps a typed name, turns a default name into entry 20 and resets
the hero choice to slot 0. — `TEXT-073`

A left-button press on a hero writes that hero's entry only when the field holds
`Unnamed` or one of entries 20..23 and the pressed slot differs from the last pressed
one: slot 0 (male fighter) entry 20, slot 1 (female fighter) entry 21, slot 2 (female
mage) entry 23, slot 3 (male mage) entry 22. A typed name is never replaced.
— `TEXT-074`

The text each press leaves when made from the first opening (`TEXT-074`):

| hero choice | EN | RU |
|---|---|---|
| slot 0, male fighter | `Danath` | `Данас` |
| slot 1, female fighter | `Naira` | `Найра` |
| slot 2, female mage | `Reniesta` | `Рениеста` |
| slot 3, male mage | `Fergard` | `Фергард` |

Typing only appends: a converted byte at or above `0x20` is added at the end while the
text is shorter than 10 bytes, and key 8 removes the last byte. There is no selection,
insertion point or replace-on-first-key, so the first keystroke extends the seeded name.
The cap limits typing only. — `TEXT-075`

The caret is the glyph `|` (`0x7c`) appended to the drawn string, so it stands at the
text's advance sum in the text's font and colour. It flips at the first paint more than
500 ms after the last flip; a character message under the cap, and the field's
construction, show it and restart that interval. The draw does not test focus.
— `TEXT-076`

The prompt, global string slot 125, and the name are left-aligned font4 draws. The
prompt's glyph cells start at screen origin + (224,305) and the name's at + (224,321),
where the origin centres the 640x480 screen on the display; no text width enters either
position. With the normal `/16` ramp the prompt's opaque ink is RGB(65,47,20) and the
name's RGB(101,39,61) before packing, on both roots. A memory check selects the `/18` ramp
when the machine reports less than 24,000,000 bytes of physical memory; that ramp gives
(57,41,17) and (89,34,54), and in that branch a level-15 word also adds a sixteenth of the
quantized background. Only the strings differ by root. — `TEXT-077`

Two different typed byte sequences can therefore become visually identical under the
Russian display transform: composing the input conversion with the display conversion
leaves a bounded set of reachable glyph records that more than one keystroke sequence
can reach. A compatible save/editor implementation should preserve the stored bytes
rather than replacing them solely from rendered appearance. — `TEXT-COLL-025`

## Save-label entry

The SAVE/LOAD chooser's list-item label is not the character-name control's bounded,
filtering class above. Its recognized producer path is a generic list-control item-text
copy that filters no byte value other than the NUL terminator, distinct from the
character-name entry's rejection of control bytes below the printable range.
— `TEXT-SAVELABEL-057`

`TEXT-SAVELABEL-055`'s exclusion of the character-name control is retracted: it searched
an address inside that class's vtable rather than the vtable itself. The control is still
not the save-label producer (`TEXT-SAVELABEL-057`), and its only direct construction is the
pre-create screen's name field (`TEXT-075`).

The byte-indexed display-conversion selector documented above for other text surfaces is
not called, directly or through its only wrapper, by any traced save-label chooser code
path. — `TEXT-SAVELABEL-054`

Neither the byte-indexed selector's own draw functions nor this image's GDI text-out
import surface is reached by any traced save-label chooser code path either; the MFC `CDC`
classes whose tables hold the GDI text wrappers are constructed, so that leg rests on the
traced chooser paths alone (`TEXT-SAVELABEL-059`'s never-constructed clause is retracted). By
elimination among these named mechanisms, the field's pixels are consistent with native
Win32/MFC list-control default painting outside this executable's own code — an
elimination among catalogued candidates, not a positive trace, and it does not exclude an
uncatalogued in-game draw routine reached only by virtual dispatch this search cannot
enumerate. Because the two lawful executables are the same file, EN and RU cannot differ
in a mechanism this image does not exhibit for this field. What the draw path does with a
byte outside 7-bit printable ASCII, and what bounds the field's *drawn* (as opposed to
retrieved) length, both require observing a running original and are not established by
static analysis. — `TEXT-SAVELABEL-058`, `TEXT-SAVELABEL-059`, `TEXT-SAVELABEL-060`,
`TEXT-SAVELABEL-061`

## Character-generation and UI labels

ROM1 mixes three presentation mechanisms:

1. positional string-table entries;
2. captions baked into bitmap resources;
3. pictorial controls with text used only for hover/help or other secondary surfaces.

A replacement engine must not assume every visible caption has a corresponding string
entry, or that every descriptive string is persistently drawn next to its control.
— `TEXT-CHARGEN-027`, `TEXT-CHARGEN-028`, `TEXT-CHARGEN-029`,
`TEXT-UI-032`, `TEXT-UI-033`, `TEXT-UI-034`, `TEXT-UI-035`, `TEXT-UI-036`, `TEXT-UI-037`,
`TEXT-UI-038`, `TEXT-UI-039`, `TEXT-UI-040`, `TEXT-UI-041`, `TEXT-UI-042`, `TEXT-UI-043`,
`TEXT-UI-044`, `TEXT-UI-045`, `TEXT-UI-046`, `TEXT-UI-047`

Where a control does take its caption from the string tables, the reached constructors
read fixed global indices and copy the bytes into storage the control owns, rather than
borrowing the loader's pointers; the control's destructor releases that storage. The
pressed and unpressed paint branches then select presentation only, not a different
caption source, and a caption element can be replaced later by a selection writer while
the surrounding numeric strings are refreshed as separate paint arguments. A consumer
must therefore treat caption identity as element position in the control's own array,
not as the rendered wording. — `TOWN-383`, `TOWN-384`, `TOWN-385`, `TOWN-391`,
`TOWN-392`, `TOWN-393`

<a id="corpus-observations-versus-rules"></a>

## Help text

`main\text\help.txt` is read once at startup as one NUL-terminated string, outside the
sixteen-file line table, and no code-page pass runs at load (`TEXT-086`). Each root's file is
CRLF-terminated paragraph lines with no other control byte; EN is 7-bit, RU carries Cyrillic
high bytes, and the two differ in blank-line placement and line counts (`TEXT-087`). The
panel wraps it with the dialogue splitter and wrapper at the body width, rewraps beside the
scroll bar, and scrolls it; it is not paged (`TEXT-088`). Its one doubled tilde draws a
literal tilde glyph (`TEXT-089`).

## Settings dialog strings (`TEXT-099`, `TEXT-100`)

Game Options reads 28 `dialogs.txt` rows and `patch.txt` rows 52 to 54; Sound Options reads 20
`dialogs.txt` rows. Every row exists on both roots (`TEXT-099`). `tunes.txt` has 21 `name.wav`
keys with titles, identical keys on both roots; `cutscene.txt` and `cutpaths.txt` have 14 rows
each, and the `cutpaths.txt` rows are identical on both roots (`TEXT-100`).

## Unknown / bounded areas

The published rules do not establish:

- a Unicode encoding for the original resources;
- safe behaviour for arbitrary malformed byte values or out-of-range font indices;
- semantic equality of every EN/RU string-table entry;
- a universal UI labelling mechanism;
- coverage of strings embedded in every possible non-text resource type;
- whether the save-label chooser's byte-to-glyph mapping can differ between EN and RU for
  this field: the two lawful executables are byte-identical (`TEXT-SAVELABEL-061`), so any
  such difference would have to come from outside `rom.exe` — no traced chooser path
  reaches this image's own draw mechanisms (`TEXT-SAVELABEL-058`; `TEXT-SAVELABEL-059`
  partially retracted), so the mapping is a property of whatever paints the field, not
  settled by that elimination;
- which native control paints the save-label chooser's field (a catalogued in-image
  mechanism is excluded by elimination, `TEXT-SAVELABEL-058`/`-061`, but the specific
  native control is not identified), what it does with a byte outside 7-bit printable
  ASCII, and what bounds the field's drawn length — all three need a running original;
- whether a keystroke reaches the pre-create name field before the field is pressed
  (`TEXT-075`), and what the `-name` source string and any untraced opening of that
  screen supply (`TEXT-073`);
- the caret's visible blink period, which adds the repaint interval to 500 ms
  (`TEXT-076`);
- the displayed pixel of the prompt and name inks, which depends on the framebuffer's
  channel widths, and whether a native original runs with the `/18` ramp (`TEXT-077`);
- which anchor values callers other than the dialogue pass to the draw routines
  (`TEXT-078`), what a tilde in the last byte of a string reads past the end of the advance
  table, and whether any text outside the dialogue families holds a tilde (`TEXT-079`).

## Decode and encode sequence

1. Parse CRLF-delimited resource tables as bytes, retaining positional order.
2. Apply the active selector's display transform to each display byte.
3. Handle controls and markup on the appropriate text path, then index
   `uint8(transformedByte-0x20)` in the selected font.
4. Measure using that font's advance sidecar and the draw path's markup rule.

For keyboard input, apply the input transform before appending accepted bytes;
backspace edits the current string. Preserve stored bytes on save or resource
rewrite. Display conversion is not injective and therefore has no unique
inverse suitable for reconstructing original text. — TEXT-STRTAB-023 (line total
partially retracted), TEXT-INDEX-003, TEXT-DOM-010, TEXT-NAMEIN-024 (vtable clauses
partially retracted), TEXT-COLL-025
