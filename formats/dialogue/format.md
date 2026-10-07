<a id="dialogue--the-window-a-missions-script-speaks-through--specification-partial"></a>

# Dialogue windows and announcements

Script and outcome messages select a dialogue or outcome panel through the
client dispatcher. Text resources, speaker flags and current panel state
control the displayed content. Missing mission-event text is silent; an announcement
arriving while a dialogue is open is discarded. — DLG-PATH-002 (amended),
DLG-MSGNUM-025

The four NPC-flag conditional effects, complete speech-name construction,
whether the text control holds focus for its key scroll (no shipped block needs
it) and source of the valid-mission list remain Unknown. Glyph conversion is
defined in [TEXT](../text/format.md).

<a id="the-one-thing-a-consumer-must-not-get-wrong"></a>

## State and identity

An absent mission-event text resource returns without fallback, retry or error. If the
dialogue-open flag is set, the arriving announcement is dropped.
— DLG-PATH-002 (amended)

<a id="the-path"></a>

## Message route

```
script instant 2          R0188   opcode 0xb6 into the static message at L02444
mission outcome           R0623   opcode = argument (0xb4 lose, 0xb5 win), SAME object
      |                   R0217   broadcast to every recipient at this+0x18b8
      v
client dispatcher         R0509   pool L00625; opcode-3 bounded 0xbb;
                                         index L02523[188] -> jump L02524[46]
      0xb6 -> PostMessage(hwnd, 0x433, number, 0)
      0xb4 -> PostMessage(hwnd, 0x433, 0xff,   0)      <- the sentinel
      0xb5 -> PostMessage(hwnd, 0x430, 0,      0)
      v
session window            R0701   0x416..0x485 via L12997[112] -> L12998[58]
      0x433 arm:  campaign+0x414 = wParam            (ALWAYS, sentinel included)
                  wParam == 0xff      -> post 0x431  (the lose panel)
                  campaign+0x3dc & 8  -> return      (a dialog is open: DROP)
                  open main\text\battle\m<mission>\event<NN>.txt
                  absent              -> return      (silence)
                  present             -> R0695(relative name)
```

### Where `<mission>` and `<NN>` come from

`<NN>` is the script parameter, carried as a **dword** the whole way — written
`L03532` stores the 32-bit number at `msg+0xa` and `L03533` reads it back
— and stored in `campaign+0x414`. `<mission>` is **`campaign+0x660`** and never comes from the
packet. That one field also formats `%d.alm`, `main\text\battle\m%d\briefing.txt` and `m%d`,
so the text directory number is the map file number by construction. It is the `+0x118` of the
record embedded at `campaign+0x548`, written only through that record's own setters; nothing in
the image writes displacement `0x660`.

**Reserved value.** `<NN> = 255` is the mission-lost sentinel and reaches the lose panel, not a
text window. `<NN>` is otherwise unconstrained; installed text identifiers use 0..25.

**Mode.** The `0x433` arm carries no `campaign+0x6bc` test, so it runs in any session. Only the
map load is mode-gated: `campaign+0x6bc == 2` builds `<n>.alm` from `campaign+0x660`, and any
other value loads the map named by the CString at `campaign+0x6b4`.

<a id="the-window"></a>

## Panel state

### Backdrop show, clipping and close

The reached mission-event and five town entry bodies call the same dialogue
builder. Its tail calls `R0361`, which sets state mask 0x8 for the ordinary
dialogue pointer, attaches the panel and invokes slots `+0x78` and `+0x80`.
The first part's return does not select the backdrop operation. Show then
requests one in-place level-3 lookup remap, unlocks, presents and calls the
panel's paint slot. Town and mission state do not select separate arithmetic.
— DIALOGUE-054

The requested rectangle is `[0, screenRight) x [0, screenBottom)`, intersected
with the incoming clip. The pixel address is `surface + y * bytePitch + 2*x`.
Neither the request nor the remap establishes that the native clip covers the
whole screen. — DIALOGUE-056

For full-table RGB565 and RGB555, each packed channel becomes
`floor(channel * 13 / 16)`. The reduced path indexes by `pixel >> 3`; for these
layouts it replaces the blue channel with `(blue & ~7) + 4` before that law.
Black therefore maps to pixel 3 in the reduced path. RGB555's bit 15 lies
outside the tested valid population. DLG-DIM-013's unqualified arithmetic and
unchanged-black consequence are partially retracted. — DIALOGUE-055

A second explicit show invocation repeats the remap even if state mask 0x8 is
already set. Command `0x46f` repaints children on continuation or reaches the
base close-message path. The default close clears mask 0x8 and requests parent
paint, except when the remaining state is 1 and `map+0x80` is zero. These
selected handlers have no direct backdrop remap or saved-surface restore.
— DIALOGUE-057

Native pixels, configured table mode, incoming clip, repaint completion and
presentation cadence remain Unknown. Child and parent painters were stubbed
in transition replay; their native output is not established by that replay.
— DIALOGUE-054, DIALOGUE-055, DIALOGUE-056, DIALOGUE-057

The failure panel is **not** the following dialogue class. Its constructor `L03302` installs
vtable `L06593`; slot `+0x48` is `L03304`, forwarding the base result/close mechanism.
It offers Exit to Main Menu (`0x445`) and Load Game (`0x446`). Base close posts `0x44c`; the
frontend matches the stored failure-panel pointer `+0x110`, then chooses `0x41e` teardown/menu
or `0x418` save selection. There is no failure-panel `R0708 -> 0x41d` hop: that method is
in a neighboring class (`DLG-PATH-002` (amended), `MISSION-DEFEAT-046`). Both outcome panels set bit 8;
the campaign idle gate pauses stepping while shown (`SESS-DEFEAT-065`).

`R0695` is the class's **only** constructor. Rects are `{left, top, right, bottom}`, and the
child rectangles are relative to the panel.

```
panel     id  9   ( 30,120)-(610,360)   0x84 bytes, vtable L03624; only the size is used
portrait  id 12   ( 30, 54)-(118,168)   only when panel+0x7c
text      id 10   (128, 36)-(428,172)   with a portrait
          id 10   ( 48, 36)-(428,172)   without one
button    id 11   (200,172)-(280,198)   label = main.txt line 77, command 0x46f
```

Six calling routines, at seven call sites, name five resource families, and all five ship:
mission events (`battle\m%d\event%02d`), the inn's NPCs (`inn\NPC\npc%02dm%d`, two call sites),
the mercenary hall (two routines, `inn\mercenary\npc%02d` and `…\npc35`, sharing one family), the
shop (`shop\npc31m%d`) and the training hall (`training\npc34m%d`). — `DLG-WIN-001`, its
six-families clause amended to five

All six routines use the same builder, because it takes only the name and shows the
panel itself. Every shipped node (268 EN, 278 RU) contains `npc`, so every shipped dialogue takes
the portrait layout. — `DLG-RECT-037`

### Offer handoff and failed first part

The reached shop and training show tails require a positive offer-array count.
They read its head mission word and pass `shop\npc31m<word>` or
`training\npc34m<word>` to the builder. After normal return they register the
word and remove array element 0, count 1, without testing the builder result.
Shop array header/data/count are campaign `+630/+634/+638`; training uses
`+644/+648/+64c`. Earlier room initialization, constructor/resource exceptions
and downstream registration effects are outside this local order. — DIALOGUE-048

The panel starts with part `+68=0`, voice `+70=0`, tips `+80=0`, payload at
`+74` and allocation length plus one at `+78`. Its show slot increments to
part 1, ignores a failed lookup and still calls base show. No successful parse
means no text replacement: child 10 retains `"Nothing to say"`. The next pager
call tries part 2, which can succeed and continue. Command `0x46f` closes only
when its own next lookup fails. Native visibility and input delivery remain
Unknown. — DIALOGUE-049, DLG-EMPTY-004 (amended)

A rejected candidate resumes the same-part scan, rather than advancing the
part. Its `npc=` test can already have changed portrait flag `+7c`. After the
conditions, `tips=` can store `+80` before a missing header LF forces return 0.
Those stores have no rollback in the parser body. A later failed `0x46f`
lookup sends `0x45b` when tips is nonzero, then `0x445`; receiver presentation
is Unknown. — DIALOGUE-050

Each preserved EN/RU MAIN.RES tree contains four shop offer nodes and three
training offer nodes. All 14 contain one unconditional part-1 candidate.
The original parser accepts those first parts across 448 supplied actor/gate
vectors under declared primitive hooks. No selected stock node exercises the
absent or rejected first-part controls. Native overrides, actor reachability
and renderer/input effects remain Unknown. — DIALOGUE-051

### Drawn geometry

The panel is not drawn at its constructor rectangle. The base constructor snaps the size and
centres it on the screen, which is 640x480, 800x600 or 1024x768 by the option string:

```
W' = ((W - 8) / 96) * 96 + 8                 580 -> 488
H' = ((((H - 104) + a) >> 6) << 6) + 104     240 -> 232      a = 63 when H - 104 < 0, else 0
left = (screenW - W') >> 1                   top = (screenH - H') >> 1
```

The 488 x 232 panel is drawn at (76,124), (156,184) and (268,268) at the three screen sizes. It has
no bitmap background. Its frame is nine pieces of `interface\lm.256`, tiled inside the rectangle
less 8 px at right and bottom (480 x 224 at 640x480). — `DLG-PANEL-035`, `DLG-WIN-001` (its
overhang bounds amended to the drawn size)

```
frames    0 96x64   1 48x48   2 96x48   3 48x48   4 48x64   5 48x64   6 48x48   7 96x48   8 48x48
corners   frame 1 at (L,T)   3 at (R-48,T)   6 at (L,B-48)   8 at (R-48,B-48)
edges     frames 2 (top) and 7 (bottom): 4 tiles at x = L+48+96i
          frames 4 (left) and 5 (right): 2 tiles at y = T+48+64j
interior  frame 0: 4 x 2 tiles from (L+48,T+48)
shadow    frames 3, 5, 6, 7 and 8 at (+8,+8), blit mode 6, before the body
```

The selected painter requests nine shadow sprites before 24 body sprites at
each supplied screen geometry. The selected sprite wrapper adds no record
origin. The selected forward normal path writes the palette word chosen by
each literal, including palette word zero; skip runs leave the destination
unchanged. Sixteen private literal/palette/clip controls exclude blending
and palette-zero keying on that path. Forward mode 6 instead remaps existing destination pixels;
literal byte values are discarded. Full RGB565/RGB555 tables apply
`floor(channel * 10 / 16)`. Reduced tables replace the blue channel with
`(blue & ~7) + 4` before that law. Shadow requests do not fill every pixel of
the 8 px band. DLG-PANEL-035's former second-draw order and band-fill wording
are amended. Native clip/table selection, reverse decoding and final frame
pixels remain Unknown. — DIALOGUE-063, DLG-PANEL-035 (amended)

The rectangles at 640x480. At 800x600 they move by (80,60) and at 1024x768 by (192,144).
— `DLG-PORTRAIT-036`, `DLG-RECT-037`

```
panel drawn     ( 76,124)-(564,356)   488 x 232
portrait child  (106,178)-(194,292)   the 88 x 108 surface is blitted at (106,178), keyed on colour 0
black fill      (114,185)-(186,279)   72 x 94, drawn before the surface
picture window  (114,185)-(186,277)   72 x 92; 72 x 96 to y = 281 when the second portrait word is not -1
text control    (204,160)-(504,295)   300 x 135 with a portrait
                (124,160)-(504,295)   380 x 135 without one
button          (276,296)-(356,322)   80 x 26
```

<a id="the-speakers-figure"></a>

## Speaker figure

Child 12's picture is built by `R0740(npcId)` and comes in two forms, chosen by the
speaker's own flags word. `DLG-FIGURE-020`, `DLG-FIGURE-021`, `DLG-SPEAKER-022`.

```
speaker = R0742(npcId)              a live actor matching the npc<n> section's
                                           Flags predicate, or 0
          ?: R0743(npcId)           else a drawable synthesised from npc.reg

surfaces  R0744(0x58,0x6c)  =  88 x 108   returned to child 12
          R0744(0xa0,0xf0)  = 160 x 240   composition canvas

speaker+0x18c & 0x11 == 0   flat: graphics\infowindow\<InfoPicture><face>.bmp into the canvas
speaker+0x18c & 0x11 != 0   figure: speaker->vt+0x80(0, canvas, 0)

blit      window of the canvas -> (8,7) of the 88 x 108 surface
          second portrait word not -1:  (x0, 0x90 - PortraitY1)-(x0 + 72, 0xf0 - PortraitY1)   72 x 96
          second portrait word -1:      (0x24,0x8c)-(0x6c,0xe8), the key-absent default            72 x 92
```

`vt+0x80` is `R0745`, the same figure compositor the world view uses. It pops 12 bytes on return;
its three parameters are the picture surface, an optional stencil surface and an optional third
surface, and none of them selects layers. The dialogue passes a null stencil, so the click-map
pass does not run and the colour pass draws every occupied slot, the head slot included.

The 88 x 108 surface is built in four layers: cleared; `interface\t_back.bmp` (160 x 240) copied
opaque through the window to (8,7); the canvas copied through the same window, keyed;
`interface\t_border.256` frame 0 (88 x 108) drawn at (0,0). The border's opening has the bounding
box 72 x 92 at (8,7), with 118 frame pixels inside it, and 286 of the 288 pixels in the four extra
rows of the 72 x 96 window lie under the frame. The child fills its 72 x 94 region black before it
blits the surface. Which rows of the 240-row canvas the window frames stays Unknown.
— `DLG-PORTRAIT-036`, `DLG-FIGURE-020` (its window height amended)

The selected portrait suffix reverses the background before its opaque copy,
copies the canvas through the keyed path, reverses the background back and
requests the border last. Both copy requests use destination (8,7) and the
same selected window. The third/fourth metadata words do not change the four
supplied suffix controls. The original bitmap reversal reverses physical
rows. — DIALOGUE-064

In the selected opaque/keyed primitives, destination y advances while
physical source rows descend from `H - 1 - sourceTop`. Opaque copy writes
zero; keyed copy skips 16-bit zero and copies nonzero words unchanged. It
does not blend. These marker/zero/clip controls do not establish upstream
canvas contents or final native orientation and exposed rows.
— DIALOGUE-064, DLG-PORTRAIT-036 (amended)

A synthesised speaker's twelve visible-equipment slots are zero, so its figure is the face sheet
alone. A live speaker's figure carries whatever it is wearing at that moment.

The synthesiser `R0743` starts the flags word at `0x48` and reads a subset of the section's
`Flags` tokens: `Hero` ORs 1, `Human` `0x10`, `Female` 4 and `Mage` 2; `MySex` and `MyClass` OR the
primary actor's own 4 and 2 (sex and class), and `!MySex` and `!MyClass` OR their inverse. `Me`,
`!Me`, `!Mage`, `!Female`, `!Hero`, `!Human`, `Platoon` and `Face` have no arm. A section with
`Start` takes its face from the `Face` key of the `npc.reg` section chosen by `bits & 6`
(`MaleFighter` 0, `MaleMage` 2, `FemaleFighter` 4, `FemaleMage` 6, default 1); any other section
takes its own `Face`. The figure is then `graphics\equipment\<directory>\<face>.256`, the directory
chosen by `bits & 6` (`mfighter`, `mmage`, `ffighter`, `fmage`), and a result whose `bits & 3` is 1
also draws the hero back layer. When no live actor passes their terms, `npc21` to `npc24` draw an
archetype chosen by the tokens and the primary's sex and class (`npc21` always `MaleFighter`).
Whether a live hero passes the terms for a given party, so that the synthesiser is never reached, is
Unknown. No party hero, no stock mercenary and no placement of the chapter-140 map (`scn:140.alm`)
passes the terms of `npc62`. If the tavern's client holds no other actor that passes, the dialogue
draws the synthesised `fmage\4.256` for it. The client's actor map was not enumerated; actors that
do pass exist in `scn:141.alm` and, on the EN root, in `Beast.ALM`, `Cross.ALM` and `Horror.alm`.
— `DLG-SYNTH-042`, `DLG-SPEAKER-041`

A person the client creates from a state message with a type id below `0x1a` takes its face from the
low seven bits of the message's face byte, and bit 7 is its Female bit. For a person the npc arm of
a placement creates in zero mode, the constructor takes the byte from the Humans row's `face` column
and ORs the row's gender column into bit 7; a row whose gender column is -1 keeps the default,
female. A face of 0 names a sheet that no root ships, and drawing it aborts the program.
— `DLG-FACEBYTE-043`

## Lifecycle

```
show      vt+0x80 -> advance one part; the RETURN IS IGNORED, so a file with no
                     part 1 opens on the literal "Nothing to say"
input     click on the button (release inside it), Enter (0x0d) or Escape (0x1b) -> command 0x46f
            Enter    the button's key handler posts 0x46f
            Escape   the panel's key slot sends 0x46f
            Space    no dialogue action
          -> another part remains: set the text, refresh the face, STAY OPEN
          -> none remains: post 0x45b with panel+0x80 when a tips= tag set it
                           post 0x445 -> vt+0x84 (free the text) -> post 0x44c
                                      -> R0709 clears campaign+0x3dc bit 3
timer     none. 13 functions in the class, not one references 0x113.
overlap   impossible: R0361 sets bit 3 on show and the 0x433 arm drops on it.
mouse     the panel is the root's capture object while it is up
```

A dialogue of N parts takes N presses of any of the three inputs, and the last one closes it. The
panel has no accept or decline state: on the last page of a quest offer Enter, Escape and the click
close it identically, and the quest is registered when the window opens (the shop and the training
hall right after the constructor returns, the inn by queueing it for commit when the player
leaves), not when it closes. A `tips=` tag stores the number after it in `panel+0x80`; 12 EN and
12 RU mission event blocks carry one and no inn, mercenary, shop or training block does.
That the button's key handler rather than the panel's key slot answers Enter is Medium: it rests
on the dispatcher's child order, and both send the same command. Whether siblings of the panel
under the root answer Space, the key-release and system-key routes and what the `0x45b` arm
shows are Unknown.
— `DLG-KEYS-040`, `DLG-LIFE-005` (amended for the routing of Enter; its clear enumeration partially
retracted)

### Captured mouse input

The captured panel gives mouse messages to its own children. A left press
outside every child returns 0 through the panel's own down slot. The root
dispatcher then runs its own mouse slot, which returns 0; it does not retry
hit tests on siblings behind the panel. The panel makes no call into its
parent to forward the press and emits no pager or close command on that path.
This composition assumes the shown campaign root and dialogue vtables; native
event ordering and capture changes by outside callers remain Unknown.
— DIALOGUE-044

The button press sets its pressed field and requests capture in the panel.
Release with that field set clears it, restores capture and tests the release
point against the button rectangle. Inside posts `0x46f`; outside posts no
command. Release without a press also posts no command. The hit test calls
`PtInRect`. The panel's own release, double-click
and right-button slots return 0. — DIALOGUE-045, DIALOGUE-044

The text child's left-down returns 0; its up and double click send `0x472` and
`0x444` to the panel, not `0x46f`. Those messages do not enter the panel's close
arm. Its drag route can scroll with mouse flag 1 and its own scroll state set.
No native input or window-timing result is established. — DIALOGUE-044

## Content

The file is scanned, not parsed: find `<` … `>`, lowercase the body, substring-search it.
Fifteen literals, in test order:

```
part=%d  npc=  iamfemale  iammale  iammage  iamfighter  npcalive=  npcdead=
female  male  mage  fighter  sound=  tune=  tips=
```

`part=%d` is a **substring** test, so `part=10` satisfies a search for part 1 — the first
matching tag in file order wins. Shipped data never trips it (0 of 513 EN / 518 RU parts).
`npcalive=`, `npcdead=`, `sound=` and `tune=` occur in no `text/battle` event file on either
root, and `iamfighter` and `fighter` in none on EN. The `sound=` zero is scoped to those event
files: on EN `sound=` occurs 7 times, all in `text/inn`, while `npcalive=`, `npcdead=` and
`tune=` occur in none of the four tag-carrying families. — `DLG-MARKUP-007` (its `sound=` zero
narrowed to event files), `DLG-NPCTAG-019` (its reason for the RU raw-file census partially
retracted; the counts stand)

Two different `npc` tests decide the face: the constructor's `Find("npc")` over the **whole
lowercased file** picks the layout once, and the per-part `Find("npc=")` refreshes the portrait.

### Resource bytes delivered to the text control

The selected accepted-part parser starts the body after the first LF
following the accepted header. It scans forward to the next `<` or NUL.
At `<` it scans backward until CR and uses that CR as the exclusive end;
at payload end it uses NUL. The backward scan has no body-start bound.
The reached pager saves that end byte, writes NUL there, passes the body
pointer to the text control and restores the byte. The text control passes
the same raw bytes to wrapping. Leading, trailing and interior bytes are
not trimmed at this resource boundary. — DIALOGUE-068

Paragraph splitting and wrapping are later operations. CRLF pieces and
their remaining suffix receive TrimLeft. Each wrapping piece receives
TrimLeft then TrimRight before fitting; later slices receive TrimLeft.
Original trim instructions preserve interior bytes in the finite controls
under supplied SPACE/TAB/CR/LF/VT/FF classification and single-byte stepping.
Native locale, multibyte stepping and copy-on-write behavior remain Unknown.
— DIALOGUE-069

The explicitly supplied installed-body projection truncates payload at NUL,
selects every part-bearing tag without testing acceptance and rejects missing
header LF or a missing CR within the body before a next tag. It matches
the original accepted-tail intervals and original-trim line arrays on all
688 EN and 732 RU candidate bodies under declared services. Rejected or
unselected tags and malformed backward ranges distinguish that projection
from a reached original caller. Native accepted-tag population and string
configuration remain Unknown. — DIALOGUE-070, DIALOGUE-062 (amended)

### The eight sex-and-class conditionals

Each **abandons the tag and resumes the scan** when the tag and the bit disagree; the scan is
one forward pass, so another tag declaring the same part can still match. When no tag is
accepted, the parser returns 0; show ignores it, while command `0x46f` closes on that failure
(DIALOGUE-049).
`iam*` tests the player's own hero;
the bare four test the speaker, and are skipped entirely unless the npc section has a `Start`
key and a speaker resolves.

```
actor+0x18c bit 0x4 = female        bit 0x2 = spellcaster
iamfemale   player  reject if clear    female   speaker  reject if clear
iammale     player  reject if set      male     speaker  reject if set  (and Find("female") == -1)
iammage     player  reject if clear    mage     speaker  reject if clear
iamfighter  player  reject if set      fighter  speaker  reject if set
```

**The literals nest and an arm that does not reject falls through.** `Find` is a substring test,
so `iamfemale` also satisfies `female`, `iammale` satisfies `male`, `iammage` satisfies `mage`
and `iamfighter` satisfies `fighter`. `iamfemale` also contains `male`, but the `male` arm
re-tests `Find("female") == -1` and skips its speaker test for any body containing `female`, so
an `iamfemale` body reaches the `female` arm's test only. That is the only guard. A part tagged
for a female player therefore also requires a female speaker, and an `iamfemale`/`iammale` pair
read to a female player by a male speaker matches neither tag. Nine bodies on each root carry
`iamfemale` and nine carry `iammale`. — `DLG-TAGARM-027`, its `iamfemale`-satisfies-`male`
clause amended

### Sound

`sound=` is an out-parameter, not an effect: the scan takes the value to the next `"` or `;` and
hands it back. The pager plays it as `speech\` + the resource path's directory prefix + the
value + `.wav`. When the tag is absent the pager composes the name itself, and **which format it
uses is decided by the resource's leaf name**: a leaf beginning `event` gives `npc%02de%sp%d`,
any other leaf gives `%sp%d`. That is the only family test in the whole surface that reads the
file name rather than the directory.

<a id="measurement"></a>

## Text measurement

The text control wraps into a line array at the rectangle's full width, 300 px with a portrait
and 380 without, and computes

```
ctrl+0x8c = min( (bottom - top) / pitch , lineCount )        pitch = fontHeight + 2
```

The control's base constructor first cuts its height to `n * (h + 4) + 2` with
`n = floor(136 / (h + 4))`. With font 1 (`h` = 15) the height is 135 and the pitch 17, so the
window shows at most 7 lines. Original wrapping instructions with declared
string/array services process the selected five families' 688 EN blocks into
2542 lines and 732 RU blocks into 2650 lines. Each maximum is 7. The installed
candidate population includes conditional parts; it does not establish native
reachability or a universal maximum. — DIALOGUE-062
The extraction and low-level string services are explicitly supplied;
their bounded original accepted-tail and trim comparison is DIALOGUE-070.

The control has a scroll setter (`vt+0x80`, notify `0x46d`), and this window builds it no
scrollbar and answers no `0x46d`. The control's own key handler moves the top line on PageUp,
PageDown, Up and Down while it holds focus. The 300 px width is not a fixed character count:
wrapping calls the text measurer (`TEXT-API-007`), whose per-glyph advance is
`.dat[glyph] + spacing` (`SPR16A-FONT-018`). The visible-line formula above uses the resulting
`lineCount`, not a fixed pixel-to-character conversion. — `DLG-WRAP-009` (amended, partially
retracted), `DLG-RECT-037`, `DLG-LINE-038`, `DLG-KEYS-040`

### Fitting and breaks

The selected wrapper uses strict width comparisons. With synthetic advances
8 and spacing 2, A measures 10, A-space 27 and A-space-B 37. At exact width
10 a lone A repeats unchanged remainder; at width 37 A-space-B becomes two
lines. At exact prefix width 27 it repeats unchanged remainder. An over-wide
first word remains whole rather than splitting by glyph. — DIALOGUE-060

CRLF separates paragraph pieces. LF alone remains data; bare CR and trailing
bare CR repeat unchanged remainder under the declared services. Leading
CRLF yields an empty first line; empty input yields no line and spaces alone
yield one empty line. This finite result does not prove native hangs or
universal malformed-input termination. — DIALOGUE-060

## Line placement

Ordinary fitting nonempty pieces keep a trailing space and their last line
ends in CR. Empty and over-wide controls have exceptions; DLG-LINE-038's
former unconditional marker rule is narrowed. — DIALOGUE-060, DLG-LINE-038 (amended)

For line `i` of `n`, counted from 0 over the whole array, with `first` the control's top line (0
unless scrolled):

```
paragraphFirst  i == 0, or line i-1 ends in CR
justified       i != n-1, and line i does not end in CR
p               10 when paragraphFirst (font 1 .dat dword 32), else 0
x0 = rect.left + p              y = rect.top + 17 * (i - first)
text            the line with its CR removed
justified       W = rect.width - p
                words = the text cut at blanks: trim the right end once, trim the left end
                        before each word, so runs of blanks collapse
                one word: drawn at x0
                gap = (W - sum of the word widths) / (words - 1), floating point
                gap stored as double; xacc = x0 stored as double
                word k drawn at trunc(xacc)
                integer width_k + stored xacc, then + stored gap in x87
                result stored as double after each word
otherwise       the whole text drawn at (x0, y), left aligned
```

A paragraph's lines are justified except its last, which stays left aligned, and a line before an
authored break is such a last line. Nothing is centred. Every draw is a shadow at (x + 1, y + 1)
in a flat (8,8,8) ramp followed by the ink at (x, y), level 15 of the white ramp, (255,255,255),
with anchor 0. At 640x480 with a portrait the line tops are 160, 177, 194, 211, 228, 245 and 262,
the left edge is 204 (214 on a paragraph's first line) and the justify width 300 (290 on that
line). — `DLG-LINE-038`, `TEXT-078`

The gap and accumulator spill as doubles. Each word performs two x87
additions before the next double store; the accumulator is not retained
between words. Integer conversion temporarily selects truncation and restores
the incoming control word. Supplied nearest-even PC24/PC53/PC64 controls
confirm this sequence. PC64 can differ by a pixel from per-add double on
discriminating synthetic widths. — DIALOGUE-061

Over the selected corpus's 11415 EN and 9550 RU justified word positions,
nearest-even PC53/PC64 spilled models agree with per-add double. Retained
PC64 differs at 634 EN and 332 RU positions. These are Medium model results
over original-hooked wrapped lines. Native precision/rounding and visible
word coordinates remain Unknown; the corpus does not select native FPU state.
— DIALOGUE-062, DLG-LINE-038 (amended)

## Button

The button is panel-relative (200,172)-(280,198), font 1, command `0x46f`, with the label of
`main.txt` line 77: "Ok" in EN, 22 px wide, and "Принять" in RU, 70 px, with no accelerator. It
has no art and no fill: its painter requests parent repaint, then the label
and a two-colour bevel. It requests level-3 remapping after painting when
flag 1 is clear; the dialogue button is built with it set. Full/reduced
level-3 laws are stated above. Native parent-paint effects remain Unknown.
— DLG-BUTTON-039 (amended), DIALOGUE-065, DIALOGUE-055

```
light RGB (41,69,63)   dark RGB (7,12,9)   each reduced to the screen format;  R' = R-1, B' = B-1
colour 1  vertical x = R', y = T+2..B'-2         vertical x = R'-1, y = T+1..B'-1
          horizontal y = B', x = L+2..R'-2       horizontal y = B'-1, x = L+1..R'-1
          pixel (R'-2,B'-2)
colour 2  horizontal y = T, x = L+2..R'-2        vertical x = L, y = T+2..B'-2
          pixel (L+1,T+1)
idle      colour 1 dark, colour 2 light          shadow offset 2
pressed   colour 1 light, colour 2 dark          shadow offset 4     (pressed flag set, cursor inside)
label     centred at (L + (R'-L)/2 + 1, T + (B'-T)/2) = (316,308) at 640x480, anchor 0xa
ink       idle ramp level 15 = (185,159,73);  hover ramp level 15 = (150,90,0);  shadow flat (8,8,8)
```

For supplied RGB565 and RGB555 fields, original painter/line/point controls
write all 299 distinct bevel pixels exactly. Label anchor (316,308) stays
fixed. Hover selects its ramp independently; pressed presentation requires
both pressed state and cursor membership. Pressed outside uses shadow 2
and the idle bevel. Twelve finite state controls cover both layouts.
— DIALOGUE-065, DLG-BUTTON-039 (amended)

The original ramp-builder prefix gives these conditional packed values.
— DIALOGUE-065

| Layout | Light bevel | Dark bevel | Text level 15 | Idle level 15 | Hover level 15 | Flat shadow |
|---|---|---|---|---|---|---|
| RGB565 | 10791 | 97 | 65535 | 48361 | 37568 | 2113 |
| RGB555 | 5383 | 33 | 32767 | 24169 | 18784 | 1057 |

Native format selection, glyph pixels, parent repaint, disabled-remap pixels,
hover/capture delivery and cadence remain Unknown. Conditional quantization
closes DLG-BUTTON-039's former supplied-format question without claiming a
native display witness. — DIALOGUE-065, DLG-BUTTON-039 (amended), TEXT-078

## Mission 70 report identity

The reported RU fragment identifies `main.res::text/battle/m70/event01.txt`,
Part1/NPC25. The preserved EN event has the same mission/event/part/NPC
identity; directory 70 joins them despite different native mission titles.
The bounded locator covers 321 RU main.res battle-text leaves. Patch sources,
native accepted tags, live speaker and entry timing remain Unknown.
— DLG-REPORT-072

## Language

`rom.exe` is byte-identical in both roots. Everything above is compiled. The two
language-carrying inputs are the event files and `main.res::text/main.txt`, which has **274
lines on both roots** — the indices 77/140/141 are compiled constants and a root with a
different line count would break them. The event corpora differ by three files, all RU-only:
`m100/event09`, `m130/event07`, `m150/event10`.

A fourth, owner-supplied pre-release data root's event corpus is not a subset of either shape: it
carries files for exactly three missions (41, 51, 91), three of which (`m41/event10-12`) exist on
neither preserved install, and every other mission's events are entirely absent from it. Where a
scene is shared with that root its `<npc=..>` speaker sequence usually matches, and every
disagreement that does occur is a same-length substitution, never a reorder or a length change.
EN and RU disagree with each other more often and in all three ways: same-length substitution,
pure reorder of the same speaker multiset, and outright length change, where one root's tag
sequence has a different element count from the other's (`TEXT-BATTLEROOT-063`).

The inn's own dialogue entry (`R0702`, one of this page's six `R0695` callers)
shows the same shape one level down: the "nothing to offer" zero arm never advances `InnMission`
regardless of its text, but that text is not uniformly backward-only — it splits by root: the
pre-release root's three shipped instances are recap with no forward-pointing content, while
EN/RU's three each add some (one pair as an explicit reciprocal proposal) (`DLG-ZEROARM-029`); and
text/voice presence for its own NPC family does not track 1:1 across roots, in either direction
(`DLG-INNVOICE-030`).

A voice wave exists only where a `speech.res` node has the name the pager composes
(`DLG-SOUND-028`): mission event text has no such node in either root, so mission 111's fourteen tagged
event parts have no voice node while its inn text is voiced (`DLG-VOICE-074`). A part's speaker is its own
`<npc=N>` tag, resolved at show time to a placed actor passing the section's terms; by the resolver's
terms, mission 110 event 02 resolves `npc67` to the single placed brigand leader, and the roots differ at
the last part (`DLG-SPEAKER-075`). Whether a voice carries speech beyond its text is not established:
Skrakan's portal-passage inn part has one more sentence in the EN text than in the RU text, and the RU
voice's extra length is within delivery variation (`DLG-VOICE-076`).
