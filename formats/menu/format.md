<a id="menu--menu-surface-composition-contracts--specification"></a>

# Menu surfaces and input

The main menu composes bitmap overlays over a full-screen background and uses
a pixel hit mask. In-play Esc menus use text rows and a sprite nine-patch.
The command panel and Drop Gold modal have separate state and input rules.
BMP and sprite codecs are defined in their respective references.
— MENU-ASSET-001, MENU-ASSET-002, MENU-MASK-003, MENU-STATE-007,
MENU-ESC-010 (partially retracted), MENU-COMBAT-017

## Asset set (18 files under `graphics/mainmenu/`)

| Entry | Dims | Bpp | Role |
|-------|------|-----|------|
| `menu_.bmp` | 640×480 | 24 | base brooch background (drawn first) |
| `menumask.bmp` | 640×480 | 8 | hit mask — palette **indices**, not colour |
| `button1..8.bmp` | see below | 24 | per-button **hover/highlight** overlay |
| `button1..8p.bmp` | see below | 24 | per-button **pressed** overlay |

The engine (`rom.exe` `R1151`) loads them as
`main\graphics\MainMenu\{menu_.bmp, MenuMask.bmp, button%d.bmp, button%dp.bmp}`, with the
button loop running `n = 1..8` — so there are exactly **8 buttons**, each with a hover and
a pressed variant (MENU-ASSET-001/002). The mask is loaded by a distinct 8-bit loader
(`R1152`) that keeps raw indices; the colour bitmaps by `R1176`.

## Hit mask (`menumask.bmp`)

- 640×480, 8bpp. The palette is the **identity grayscale ramp** `palette[i]=(i,i,i)`, so
  the palette is decorative — the **8-bit index value** is the semantic (MENU-MASK-003).
- The hit-test `R1148` reads the index at the cursor as
  `idx = maskBuf[(y − top)·640 + (x − left)]` (menu origin `(left,top) = (0,0)`), then:

  | index | button | index | button |
  |------:|:------:|------:|:------:|
  | 0x80 | 1 | 0xc0 | 5 |
  | 0x90 | 2 | 0xd0 | 6 |
  | 0xa0 | 3 | 0xe0 | 7 |
  | 0xb0 | 4 | 0xf0 | 8 |

  Asset number = `idx/16 − 7`. Index 0 (background) and **every** other value — including
  the anti-alias edge ramp `0x10,0x12,…,0x1e` (= hot ≫ 3) and stray pixels — select **no
  button** (switch default) (MENU-MASK-004). Of the 43 index values present in the shipped
  mask, only these 8 are "hot".

## Overlay placement (static `rom.exe` tables)

Two contiguous tables, 8 entries × 16 B = `{x, y, w, h}` (int32 LE):

- **normal/hover** — the global at `L06179` (file offset `0x198ed8`)
- **pressed** — the global at `L06180` (immediately after, `+0x80`)

`R1153` places button *i* at screen rect `(x, y) … (x+w, y+h)`, offset by the menu
origin (0,0). Overlays draw **1:1** — each `(w,h)` equals the BMP's pixel dimensions,
each normal rectangle brackets its button's mask region
(MENU-GEOM-005/006).

| btn | hover (x,y,w,h) | pressed (x,y,w,h) |
|----:|-----------------|-------------------|
| 1 | 112, 64, 212, 136 | 116, 64, 208, 138 |
| 2 | 84, 88, 236, 148 | 88, 88, 236, 152 |
| 3 | 84, 236, 236, 152 | 88, 236, 236, 152 |
| 4 | 112, 276, 212, 136 | 116, 272, 208, 140 |
| 5 | 320, 60, 212, 140 | 320, 64, 212, 140 |
| 6 | 324, 88, 236, 148 | 324, 88, 232, 152 |
| 7 | 324, 236, 236, 152 | 324, 236, 236, 152 |
| 8 | 324, 276, 208, 136 | 320, 272, 212, 140 |

Layout: two columns of 4 (buttons 1–4 left, 5–8 right), each column top→bottom.

## Interaction / draw state

- Object fields (menu `this`): `+0x6c` hover-overlay ptr array, `+0x80` pressed-overlay
  ptr array, `+0x94` hover placement RECTs, `+0xa8` pressed placement RECTs, `+0xb8`
  background, `+0xbc` mask, `+0xd8` hovered index, `+0xdc` shown index, `+0xe0`
  pressed-latch index, `+0xe4` per-button disable bitfield.
- Selection (`R1148`): while not pressed, the **hover** overlay of the hovered
  button is composited over the base at its normal rect; on mouse-down over the latched
  button, its **pressed** overlay is composited at the pressed rect. A set bit in the
  `+0xe4` disable field suppresses the overlay for that button (MENU-STATE-007).
- Click (`R0820`): posts a per-button command message; button 8 → WM_CLOSE (quit).
  The other message ids are not decoded (Medium).

## The in-mission command panel (`MENU-COMBAT-018`, `MENU-COMBAT-019`)

The mission root owns a map viewport and a fixed 160-pixel right column. The right column owns,
in draw and hit-test order, the minimap, command panel, character panel and a lower filler:

| surface | local rectangle | 640x480 screen rectangle | 800x600 | 1024x768 |
|---|---|---|---|---|
| minimap | `(0,0,160,158)` | `(480,0,640,158)` | `(640,0,800,158)` | `(864,0,1024,158)` |
| command | `(0,158,160,238)` | `(480,158,640,238)` | `(640,158,800,238)` | `(864,158,1024,238)` |
| character | `(0,238,160,480)` | `(480,238,640,480)` | `(640,238,800,480)` | `(864,238,1024,480)` |
| filler | `(0,480,160,H)` | empty | `(640,480,800,600)` | `(864,480,1024,768)` |

The command panel has no child controls. Its eight hit cells are 34x34, row-major, at
`left=8+34*(i&3)`, `top=7+34*(i>>2)`. Right and bottom are exclusive, as in `PtInRect`.
Inventory, spellbook, portrait and Esc controls are in the character-panel sibling and must not
be attached to this object. That sibling is not display-only: an admitted inventory-item release
can arm Cast mode 5, and its transfer branches can produce item orders `0x22/0x32`
(MENU-COMBAT-017, AI-PANEL-123). The measured inventory-open route appends the inventory grid to the map view, correcting the partially retracted parent clauses of MENU-COMBAT-017 and SESS-INPUT-037 (MENU-088). The grid also transfers the selected item on
physical left-up through vslot `+0xa4`. Destination code 2 accepts a valid hit as an insertion index
and uses merge/insert-before-gold/append fallback for `-1` or an out-of-range hit.
The grid also owns two left-down edge strips that scroll its viewport, and left-button move starts
an item or money drag before that release. Its three right-button edges are consumed no-ops. None
of these inventory controls is physically part of the command-panel object.

| i | rect | tooltip / key | action |
|---:|---|---|---|
| 0 | `(8,7)-(42,41)` | Attack / A | arm mode 1 |
| 1 | `(42,7)-(76,41)` | Move / M | arm mode 2 |
| 2 | `(76,7)-(110,41)` | Guard / G | immediate order `0x17` |
| 3 | `(110,7)-(144,41)` | Defend / D | arm mode 4 |
| 4 | `(8,41)-(42,75)` | Cast / C | arm mode 5; suppressed while cast UI is open |
| 5 | `(42,41)-(76,75)` | Swarm / S | arm mode 6 |
| 6 | `(76,41)-(110,75)` | Stand Ground / T | immediate order `0x18` |
| 7 | `(110,41)-(144,75)` | Retreat / R | immediate order `0x14` |

Tooltip source is `main.res::text/main.txt[i]`. Russian labels, in the same order, are
`Атаковать`, `Идти`, `Охранять`, `Защищать`, `Колдовать`,
`Идти в боевой готовности`, `Держать позицию`, `Отступить`; the accelerator letters remain
Latin. A tooltip appears only while the panel and cell are enabled, no special cursor or child
dialog is active, and screen-state bits `0x0a` are clear (MENU-COMBAT-019).

The draw is four 160x80, 24-bpp BMPs. Inactive: `HeadsR.bmp`. Active: draw `CommandBarR.bmp`,
copy the same 34x34 cell from `CommandEmpR.bmp` over every disabled cell, then copy the selected
cell from `CommandDnR.bmp`. Icons and borders are embedded in these full-panel images; there is
no separate hover image. Left down acts immediately, double-click aliases it, and left-drag
repeats it as the pointer crosses enabled cells. Right up sends map cancel `0x405`; the other
right edges are no-ops (MENU-COMBAT-018/019).

**Customisation boundary.** Geometry, cell order and action binding are executable constants.
Labels and every visible pixel are resource data. Enlarging the panel or adding a ninth cell
requires code and new art; replacing a label or embedded icon changes only resource bytes.

## The Drop Gold modal

The mission root always constructs a 296x168 dialog at `(100,height-200)-(396,height-32)` (height is the screen height) and stores it at
`campaign+0x10c`. It is attached only when physical `WM_LBUTTONDBLCLK` on the inventory
grid hits the sentinel gold cell (`item+6 == 0xffff`) while the local purse is positive.

Its children, in construction order, are an amount edit id `0x989685` at local
`(30,65)-(266,85)`, action button id `0x989681` at `(68,118)-(138,138)` posting `0x445`, cancel
id `0x989682` at `(158,118)-(228,138)` posting `0x446`, and two text rows at
`(20,20)-(276,40)` and `(20,40)-(276,60)`. Button labels are
`dialogs.txt[0..1]`: EN `Ok` / `Cancel`, RU `Принять` / `Отменить`; auxiliary strings
`[46..49]` supply its auxiliary strings. The modal input contract is `AI-PANEL-123`: buttons arm
on left down and post on an inside left up.
The modal own Enter aliases action; Esc aliases cancel. A focused cancel button
consumes Enter as cancel before the modal own slot (MENU-093).

Action parses the edit, resolves the entered amount, subtracts it from the purse and emits opcode
`0x23` at the current map cell before closing. Cancel closes without an order. Its screen placement
and child geometry are executable constants; all six strings are localised `main.res` data. The generic dialog-frame assets are outside this command-panel contract.

Every measured modal open sets text to `0`. The fresh edit constructor sets
caret and selection fields to zero; its open-time text setter does not reset
them. Opening tests the client purse, while selection/ownership admission
occurs later (MENU-090). MENU-088's explicit selection-reset clause is
partially retracted by MENU-091; its parent/focus and clear-suppression facts stand.

The generic edit maps Left/Right/Home/End and Backspace/Delete. Shift changes
selection endpoints. Backspace edits before the caret without a local endpoint
store; Delete first removes a positive selection, otherwise edits at the caret.
The admitted character/insertion bodies contain no digit-only test. Tab uses
generic focus traversal, reversed with Shift (MENU-091, MENU-093).

Action obtains text through a formatting call, uses `%f` and converts its
float to an integer without checking parse success. Empty, malformed,
percent-format, locale and out-of-range results remain Unknown (MENU-092).
For a matched positive gold source with no arithmetic overflow, selection
below its quantity creates a copy holding the request and leaves the remainder;
selection at least its quantity creates a copy holding the old quantity and
retains a zero-quantity sentinel. Zero/negative requests are not rejected by
that helper. It stores campaign selected object/index/source, then modal action
debits client Player `+0c`, enqueues `0x23`, clears selection/cursor fields and
deletes the temporary object. A refused resolver has no null-result guard in
the action body. Cancel bypasses those operations (MENU-093).

Server `0x23` independently resolves the packet Player and its first actor,
requires a positive affordable amount, debits Player `+38` and creates or adds
to Sack `+3c`, using a per-axis distance limit of 2 and an actor-cell fallback.
Its local arm ignores the Sack wrapper's zero return after that debit. Pickup
credits positive Sack gold to the looter's owner Player. Runtime refusal/refund
and client reconciliation remain Unknown (MENU-094).

The Player/Sack serializer bodies write current server `+38/+3c`; Player gold
uses a reversible XOR and Sack gold a direct four-byte scalar. LOAD restores
those scalar sources for later consumers. Native round-trip, post-LOAD pickup
and SAVE-before-queue-execution remain Unknown (MENU-095).

## The in-play Esc menus (`MENU-ESC-010` (partially retracted)…`MENU-INPUT-016` (partially retracted))

Two surfaces, not one. `VK_ESCAPE` reaches the frame window's `WM_KEYDOWN` handler
`R0231`, which reads the UI state word `campaign+0x3dc`:

| state | meaning | posts | constructor | screen rect | rows |
|---|---|---|---|---|---|
| `== 1` | map session on screen | `0x416` | `R0362` | `(100,60)-(440,400)`, 340×340 | 7 of 8 built |
| `== 0` | town | `0x41f` | `R0363` | `(100,100)-(440,340)`, 340×240 | 5 |
| other | — | nothing | — | key forwarded to `campaign+0xcc` | — |

The town arm additionally requires the `campaign+0x3b4` CString to be empty. The
mission surface is also raised by a 32×32 button on the character-panel sibling at panel-local
`(0x7e,0xce)-(0x9e,0xee)`, whose tooltip is global string index 14 (MENU-ESC-010, partially
retracted: this row is the owning-panel correction).

Entries, in screen order. Labels come from `main.res::text/dialogs.txt` through the
descriptor at `L06197`; `R0668(table,i)` resolves to
`[L04369][table->+0xc + i]` (MENU-ITEM-011, MENU-ITEM-012):

| row | mission label (index) | msg | town label (index) | msg |
|----:|---|---|---|---|
| 1 | `~Save Game` (0x22) | `0x41a` | `~Save Game` (0x22) | `0x41a` |
| 2 | `~Load Game` (0x23) **or** `Diplomacy` (0x4c) | `0x418` / `0x43c` | `~Load Game` (0x23) | `0x418` |
| 3 | `Game ~Options` (0x24) | `0x41b` | `Sou~nd Options` (0x25) | `0x422` |
| 4 | `Sou~nd Options` (0x25) | `0x422` | `Abort Game` (0x4d) | `0x41c` |
| 5 | `~Quest Objectives` (0x26) | `0x420` | `~Return to Game` (0x28) | `0x446` |
| 6 | `~End Quest` (0x27) | `0x41c` | | |
| 7 | `~Return to Game` (0x28) | `0x446` | | |

Row 2 of the mission menu is `~Load Game` when `campaign+0x6bc == 2` and `Diplomacy`
otherwise; only one of the two is built. `0x446` is the family-wide close message,
not a distinct action. The first argument of the button constructor is a **control
id, not a row index**.

Disabled at construction — mission menu only; every town row is built enabled:

| row | disabled when |
|---|---|
| `~Save Game` | `campaign+0x6bc` is 0 or 1 |
| `~Load Game` | no file matches `game*.sav` in the game directory (`R1169`) |
| `Sou~nd Options` | `R1170()` returns 5, its `this+0x9c == 0` arm |
| `~Quest Objectives` | `campaign+0x6bc != 2` |

**Accelerator.** The constructor passes an immediate, then `R0700` overwrites
it with the character following a single `~` in the label, lowercased by
`R1173` (CP866-aware). `~~` is an escape. Labels without a `~` keep the
immediate, which on the English release is `Diplomacy` → `D` and `Abort Game` → `E`;
on the Russian release both labels carry a `~` and the accelerator is a different,
language-dependent letter. Hard-coding the accelerators reproduces one release only
(MENU-KEY-013).

**The shared message is not a shared action.** Both surfaces post `0x41c`, and
`R0701` branches on `campaign+0x3dc`: `== 1` raises the five-row confirmation
`R1171` (`Change Map` or `~Victory!` at `0x41d`, `~Exit to Main Menu` and
`Exit to ~Windows` both at `0x41e`, `~Return to Game`), `== 0` raises the three-row
`R1172` (the same two exit rows and `~Return to Game`), both at
`(100,100)-(440,340)` (MENU-ITEM-012).

**Geometry.** Each row is a panel-local rect `(40, 40+30(n−1), width−48, 40+30n)`:
252 px wide, 30 px tall, the first starting 40 px below the panel top
(`R1168` with step `0x1e`, `R1177` forcing left `0x28` and right
`width − 0x30`).

**Art.** Neither surface loads a bitmap: the family's art-load slot is `RET`. The
frame is a nine-patch drawn by `R0760` out of `graphics.res::interface/lm.256`
(one writer of `[L03604]`, one reader). Frames 0..8: centre 96×64, corners 48×48,
top and bottom edges 96×48, left and right edges 48×64. Inset 48, horizontal step 96,
vertical step 64. Two passes: a drop shadow of the right column and bottom row offset
`+8,+8`, then the full nine-patch at the panel rect shrunk by 8 on right and bottom.
Frames 9..17 are a second nine-patch at 32/48 px, unused by this routine
(MENU-ART-014).

**Compositing and modality** are not part of this contract and are specified by the
claims: the whole screen is darkened once, destructively, at shade level 3
(`MENU-STOP-015`, `DLG-DIM-013`); the world stops on the `0x4008` idle gate
(`MENU-STOP-015`, `DLG-STOP-012`); Esc also closes, and the panel is the root's capture
object while it is up (`MENU-INPUT-016`, partially retracted).

## F1 help and the Cast key (`MENU-051`…`MENU-056`)

F1 is posted by the campaign frame's key handler whatever the state word (the F1 key identity
is `AI-KEY-125`), and the frame builds help only when the state word `campaign+0x3dc` is
exactly 1: a map session with no overlay bit set. It is ignored in the town and over panels
that set an overlay bit: the Esc menus, Pause, dialogue panels, the Drop Gold modal, the save
and load chooser panels and help itself. The chat entry and the spellbook popup were found to leave
the word at 1, so help should open over them (Medium: absence over the named routines and a
displacement census; `MENU-051`, `MENU-080`).

Help is the Pause panel's class and constructor call with `text/help.txt` as the text: a
modal panel of constructor rectangle 576x384 snapped to 488x360 and centred, one OK button
labelled from the first line of `dialogs.txt`, no title, and a font-1 text control that adds
a 24-pixel scroll bar when the wrapped text is taller than the body (`MENU-052`; layout in
`TEXT-088`). Both shipped files overflow, so help always scrolls. It stops the world through
the Esc menus' idle gate and closes on OK, Esc or the same close messages `0x445` and
`0x446` (`MENU-053`, `MENU-STOP-015`). The text control holds the initial focus, so Up, Down, Page
Up and Page Down scroll by one line or about one page while it holds focus and do nothing while
the OK button holds focus; the current line starts at -1, so the first Down and the first Page Up
move nothing visible, and Page Down's second branch cannot fire with 12 visible rows; Tab moves focus between the two, and Enter closes help from either (`MENU-078`). The
text control is 382x211 at (40,56) with 12 visible lines, the scroll bar 24x211 at (422,56) with
positions 0..lines minus 12 (range argument lines minus 11), the OK button 96x24 at (196,300); the text is grey with a
one-pixel shadow (`MENU-079`). The scroll bar takes arrow-end clicks, track clicks and thumb
drags, and no wheel route was found (Medium; `MENU-078`). The key reaches no other help channel in the
searched population (`MENU-056`, bounded).

The C key arms Cast only in a map session with no text entry open, a nonzero selection
count and `view+0x144 & 0x24` clear; a second press while the spellbook popup exists is
consumed (`MENU-054`). With no selection or `view+0x144 & 0x24` set, the key falls to a second
dispatch that returns 0; the frame handler discards that 0 and always calls the MFC default (`MENU-054`, `MENU-063`). Cast mode is
armed only when a selected object set capability bit `0x200`; for a nonempty selection with
the bit clear the key is consumed with no message, sound call or state change, and an armed
Cast posts `0x408` to the right column and chooses no spell (`MENU-055`, `AI-SPELLCAP-288`). The
`0x408` shows no popup: it resets one field in three right-column panels, and the spellbook popup is
the separate object `campaign+0xec`, opened by the key's own routine and never given a current spell
by the C path (`MENU-064`, `MENU-065`).

## Settings shortcut and speed notices (`MENU-057`…`MENU-059`)

In a map session, Ctrl plus W, F, H, U, L, N or O changes one setting and appends one line
to the map message line: retreat mode (W, three states), formation (F, three), show health
(H, two), autoheal (U, three), flying damage (L, two), day/night changes (N, two) and
smoothing (O, two). The line is slot `base + new state` of the shared text array, with bases
94, 97, 100, 218, 102, 104 and 106; the arm reads no table of state names (`MENU-057`).
Retreat, formation and autoheal also send a type `0x46` record before the post. The numpad
speed keys without Ctrl, in a campaign session on the map screen, step the speed index by
one within 0..8 and post slot 108 plus the clamped index through the post that drops a
duplicate of the newest line; the Ctrl variants post nothing (`MENU-058`).

All of these lines are grey and live 2000 ms on the message line of `MISSION-MSGLINE-056`:
each is appended below the existing lines, the oldest is removed past the capacity of that
list, and the toggle posts never drop a duplicate. A toggle needs the map screen word equal
to 1, a closed text entry and the Ctrl latch, and the key handler tests no phase, player
count, selection or option; the speed step also needs phase 2 (`MENU-059`,
`MISSION-MSGLINE-057`). The frame forwards letter keys `A`..`Z` to the root as message `0x100` (`MENU-063`). The Pause panel's text is the separate modal-panel line, slot 119, and in the
searched population no message-line line (`TEXT-094`, Medium). Slots 94..116 and
218..220 have no other array reader in the searched population (`TEXT-095`, Medium).

## F12, Backspace and the Alt band (`MENU-060`...`MENU-065`, `AI-378`)

In a map session, with the state word 1 and the chat entry closed, F12 flips the static flag
`L06275`, which starts clear. While it is set the map view's frame routine draws a box in the
box at an offset from `[view+0x10]` and a `%3.1f fps` line from a double it recomputes
periodically; no other instruction references the flag (`MENU-060`). The box is 90x24 at
(width - 120, 0) and the text is right-aligned 38 pixels from the view's right edge at y 0; the
value is draw calls times 1000.0 over elapsed `timeGetTime` milliseconds, recomputed when more
than 1000 ms have accumulated (`MENU-081`). The constructor explicitly starts
the rate at double 0.0. The first draw primes its previous timestamp, and sampling
runs before the F12 flag test, so a later first activation can show an already
sampled value. The F12 toggle does not reset sampling (`MENU-083`). The selected
font is `graphics.res::font1/font1.16` with `font1.dat` advances; both complete
nodes match between the measured EN/RU roots (`MENU-084`).

Box and shadow use component 8 per channel; the glyph samples index a 16-entry
ink ramp with components `17*i`, `i=0..15`, and the shadow is offset (1,1).
Surface-derived masks determine the truncating channel shifts. With supplied
RGB565 masks the box and shadow word is `0x0841` and ink entry 15 `0xffff`;
RGB555 gives `0x0421` and `0x7fff`. Active native masks and displayed pixels
remain Unknown (`MENU-085`). The map view's own Backspace arm requires
screen word `1` and closed chat, clears the message line through three
SetSize-like calls on its text, colour and lifetime arrays, and returns `0`
(`MENU-061`). The mapped ordinary map objects' own Backspace slots add no
effect after that clear (`MENU-082`).

The GUI offers ordinary keys to focus, then children in insertion order,
then the own slot. A nonzero return stops that route; zero permits another
offer to the focus object if it also appears in the child list. The key walk
has no visibility/enabled filter. Map entry clears root focus; later dialog
opening can set immediate-parent focus (`MENU-063`, `MENU-086`).

Opened chat focuses its text child. Backspace returns `1`, takes the current
byte-string prefix of length minus one, trims a leading classified run,
and restores the last stored row when the result is empty. The classifier's
whitespace interpretation is Medium. This route suppresses message-line
clear (`MENU-087`). Drop Gold becomes root focus. Its edit requires focus
flag `4`, a parent and flag `1`; Backspace deletes before a positive caret
and returns `1`, or returns `0` at caret zero. Both admitted outcomes update
the edit time and notify the parent with `0x46d`. The modal's screen bit `8`
independently suppresses the clear (`MENU-088`, whose explicit selection-reset
clause is partially retracted by MENU-091).

These are conditional native handler facts. The 96-table common-GUI
signature and 42 keydown targets do not close every live map-session
descendant/focus route; further actions remain Unknown (`MENU-089`).
Alt+Backspace and Alt+F12 arrive as system key messages, which the frame
procedure leaves to its default arm, and reach neither arm (`MENU-082`).

Alt plus a letter B..Y except S broadcasts one type `0x46` record with sub-selector `0x80` and the
index letter minus `A`; Alt+S is the screenshot (`MENU-062`). The record's receiver arm is a debug
console that acts only for a Player whose privilege byte exceeds `0x32`. A fresh Player holds 0, and the chat line
`#Chicken` sets it to `0xff` (`MENU-102`, `MENU-103`). D toggles turn tracing, T script tracing, Q the AI admission override
and prints its state, H prints a help, I the last-turn and average-turn AI statistics, U the
mission unit experience; the other 17 keys do nothing (`AI-378`).

The console tests no participant flag, and the byte can be raised only where `#Chicken` acts
(`AI-397`, Medium; a Player privileged earlier in the process is not excluded). A reply is a `0x91` record with addressee 0, which the send routine takes as all connections; the filter inside that loop is Unknown (`AI-398`).

Chat lines that begin with `#` reach one typed-command parser of 13 literals and 12 commands (`MENU-099`), which returns at once on
a map whose participant flag is set (`MENU-100`). The privilege byte is checked by nine arms of the 12 commands
and not by `#modify`, `#event` or `#Chicken` (`MENU-101`); `#Chicken` raises it to `0xff`
(`MENU-102`). The byte is not saved and a LOAD leaves it at 0 by the constructor route (`MENU-103`, Medium). The commands:

| Command | Effect | Claim |
|---|---|---|
| `#create [N ]<name>` | N gold for `Gold`, else N of a named item in the hero's inventory | `MENU-104` |
| `#modify self\|army +god` | modifier protection and damage-kind bytes to 100 | `MENU-105` |
| `#modify self +spell <id>`, `+spells` | spell id, spells 1 to 28 into the spellbook | `MENU-105` |
| `#modify self\|army +knowledge` | Diary resend to the client | `MENU-105` |
| `#summon [hero ]<name>` | N creatures or one hero | `MENU-106` |
| `#killall`, `#kill all`, `#kill cheaters`, `#kill <name>` | kill the actors of the targets | `MENU-107` |
| `#pickup all` | every Sack picked up by the hero | `MENU-108` |
| `#show map`, `#hide map`, `#victory` | opcode `0xaa` selector 1, 0, 2 | `MENU-109` |
| `#event <n>` | event text panel n | `MENU-110` |

Three broadcast notices report a cheater, an ineffective use and a successful use (`MENU-111`). The
phase-3 host console adds `disconnect <id>` and `curse <id>` (`MENU-112`). No other cheat input was
found in the searched populations (`MENU-113`); the SAV consequence of each state change is
`MENU-114`.

The frame handler never reads the map key handler's return: it forwards the key, then always
calls the MFC default, so a return of 0 only lets the root offer the key to its other children
(`MENU-063`).

<a id="open--not-established"></a>

## Game Options and Sound Options dialogs (`MENU-073`...`MENU-077`)

Both dialogs are modal children of the campaign window and share one dialog base (`MENU-077`).
The base snaps a requested rectangle to a width of `96k + 8` and a height of `64k + 104`, centres
it, and paints a nine-piece frame with an 8 px shadow; the title and the buttons are controls the
builder adds. Esc closes any dialog of the base with message `0x446`.

Game Options has 19 controls: five labels, a speed slider, eight checkboxes, three radio groups
and OK and Cancel (`MENU-073`). Each control keeps a local copy of its bound option. Only OK
writes them: the speed level, the flags, Smoothing, the graphics flags, the party display values
and the autohealing mode, in a fixed order, with Animation 0 forcing Lighting 0, and then three
party commands. Cancel and Esc write nothing; the dialog's own code restores nothing, and the receiver of the follow-up close message was not read (`MENU-074`).

Sound Options has 15 controls and no Cancel button (`MENU-075`). Random Order, the track
selection, the three volume sliders, Play and Stop act when the control changes. OK exports only
the Acknowledgments checkbox and closes; Esc closes without it and leaves the moved volumes in
place (`MENU-076`). The speech slider (id 8) sets `[L02739]`, the speech attenuation of the
sound configuration object (default -700, saved as `SoundSpeechPos`); the effects slider sets
`[L02733]`. Voices add the speech setting once to the distance term and effects add theirs; no
instruction combines them (`ANIM-128`).

## Unknowns

- The `+0xe4` disable bits' source (which buttons start disabled) and the non-WM_CLOSE
  command-message meanings — needs further `rom.exe`/runtime.
- The DIB loader's row order (top-down vs flip) is inferred from the 8/8 placement bracket
  ; the screen y here is top-down. Its native row-order behavior remains Unknown.
- How the Esc menus' nine-patch covers a frame width that is not `96 + 96k`. The tile
  counts are `(w−96)/96` and `(h−96)/64` with truncating division, which for the
  332×332 frame these panels have leaves a 44-px band inside the right edge and one
  inside the bottom edge unaccounted for by any placed tile. `MENU-ART-014` retains the unaccounted-for band as an unresolved draw-boundary prediction.
- What each entry's raised panel then does. `MENU-ITEM-011` and `MENU-ITEM-012` name the constructor and rect each
  message raises and stops there; the two exit rows of both `0x41c` confirmations post the same
  `0x41e` and differ only in control id, and how the exit target is distinguished is unread.
- Whether a click landing outside an Esc panel changes anything: the panel holds the root's mouse
  capture and the slots read drop the click (`MENU-INPUT-016`, partially retracted).
- The native pixel value of the help panel's grey text and of the readout box (`MENU-079`, `MENU-085`),
  what the three right-column panels' fields set by `0x408` mean and whether the Cast popup's
  sound call plays a sound (`MENU-064`).
- The first visible F12 rate after an entire native startup (`MENU-083`), whether a
  single-player session loops the Alt record back (`MENU-062`), the readers of the two console
  trace flags and where a client shows the console's chat lines (`AI-378`), arbitrary later
  root/descendant focus and further Backspace effects (`MENU-089`), and whether the save and load chooser panels embed a native
  child window that takes F1 (`MENU-080`).
- Whether a phase-3 or networked session delivers the settings keys, the interval of the tick
  message that expires message-line lines, and the receiver's use of the `0x46` records
  (`MENU-057`, `MENU-059`).
- Whether the Esc menu is the only poster of the Game Options message, the meaning of campaign
  `+0x6bc` values 0 and 1, and what the base close routine's follow-up message does with the
  result code (`MENU-073`, `MENU-074`); the consumer of the music player flag that the Sound
  Options dialog sets to 1 at build and clears only on OK (`MENU-075`, `MENU-076`).
