# Repeated interface widgets

The selected push buttons, sliders, scroll bars and lists are executable
classes with shared painters. Resource art supplies frame pixels; constructor
arguments and local handlers supply geometry and interaction. The contracts
below identify selected callers and leave native delivery and a complete
widget census open. — MENU-115, MENU-117, MENU-121, MENU-129

## Push buttons

The in-play menu wrapper `R1167` and dialogue pager use generic button
painter `R0768`. It draws a procedural bevel, a centered label and a
label shadow; there is no button sprite or interior fill. Bevel RGB requests
are light `(41,69,63)` and dark `(7,12,9)`. A pressed latch while inside
swaps the bevel and changes shadow distance from 2 to 4. The label stays
at the same anchor. Disabled painting applies level-3 remapping afterward.
— MENU-115

For half-open rectangle `(L,T,R,B)`, the label anchor is
`(L+trunc((R-1-L)/2)+1,T+trunc((B-1-T)/2))`, with anchor flags 10.
Menu rows are `(40,40+30*(n-1),width-48,40+30*n)` relative to the panel.
The menu's zero-color argument selects an initial grey label ramp, or gold
when enabled and hovered or focused. Focus has no separate outline here.
The dialogue pager's nonzero-color argument selects gold idle and brown
hover. Initial requested RGB is grey `(210,210,210)`, gold `(185,159,73)`
and brown `(150,90,0)`; runtime cell replacement and native quantization
remain Unknown. — MENU-115, MENU-ART-014, DIALOGUE-065

The menu button class shares the painter and the down latch but not the
release, key and character handlers. In both classes enabled down with a
parent latches and captures, and a prior latch followed by up inside
activates; outside up clears without activation. Right/bottom are
exclusive. The generic button sends its command field as the message; the
menu button sends `0x47a` carrying the command field. The menu release
handler does not recheck parent or enabled state. The menu character
handler accepts Enter (byte 13) and the lowercased caption accelerator
when enabled with a parent; its key handler declines every key. In the
generic button Enter arrives through the key handler, and parent
focus-first dispatch lets focused Cancel consume Enter before default OK.
Space has no separate local button arm. Native delivery of Enter as a
character, capture replacement and disable/deletion between events remain
Unknown. — MENU-116, DIALOGUE-045, MENU-KEY-013, MENU-086, MENU-093

## Sliders and scroll bars

Constructor `R1185`, table `L06344` and painter `R2140` are shared.
The painter reads `graphics.res::interface/scrlbars.256`. All 26 indexed
frames are 24x24. Each selected part requests shadow at offset `(4,4)`,
mode 4, then body mode 0. Disabled presentation requests level-3 remapping.
— MENU-117, MENU-129

| Orientation | First cap | Track | Last cap | Thumb | Cap state variants |
|---|---:|---:|---:|---:|---|
| Horizontal | 0 | 7 | 8 | 10 | 3 / 11 |
| Vertical, width less than height | 18 | 19 | 20 | 22 | 21 / 23 |

Mouse move writes the endcap state fields from pointer membership. These
variants are not established as press-only art. The thumb is fixed art,
without a visible/total-row size fraction in this painter. — MENU-117

For horizontal width `W`, height `H`, left `L`, top `T`, maximum `N` and
position `p`, knob hit-left is
`K=L+H-4+trunc(p*(W-2H-4)/N)` for nonzero `N`, otherwise `L+H-4`.
Hit width is 16, although the artwork is 24 wide: body is requested at
`(K+1,T)`, shadow at `(K+5,T+4)`. Pointer mapping is
`clamp(trunc(N*(x-L-H-2)/(W-2H-4)),0,N)`.
Track down and drag use that mapping. Endcaps and focused Left/Right
step `max(trunc(N/16),1)`, so speed `N=8` steps one level. Changes post
`0x46d`; release posts `0x473` and clears capture. Volume attenuation
conversion and speed-at-OK application are separate owner contracts.
— MENU-118, MENU-074, MENU-076

The sound configuration initializer sets all three volume maxima to 5000
and initial attenuation to -700. Their default arrow/endcap step is therefore
312 position units; it is not a fixed attenuation step under the quadratic
conversion. Later arbitrary range writers remain outside the selected read.
— MENU-118

The selected settings builders request these 24-pixel-high rectangles.
They are constructor geometry, with the existing owner confidence limits.
— MENU-073, MENU-075

| Owner / slider | Requested rectangle |
|---|---|
| Game speed | `(40,84,232,108)` |
| Music volume | `(258,190,W-40,214)` |
| Effects volume | `(258,240,W-40,264)` |
| Speech volume | `(258,290,W-40,314)` |

For the vertical bar, thumb travel is
`q=trunc(p*(H-3W+8)/(N-1))` for `N>=2`, otherwise zero.
Body is `(L,T+W+q-4)` and shadow `(L+4,T+W+q)`.
Arrows request `0x469/0x46a`; track above/below requests `0x46b/0x46c`.
The nondegenerate drag request `0x468` carries
`clamp(trunc((N-1)*(y-T-24)/(H-3*(W-4))),0,N-1)`.
Drawing and drag use different denominators. Vertical key arms belong
to the associated list/text control. A held-left move can be offered
to down again; native repeat cadence is Unknown. — MENU-119

Degenerate geometry, negative ranges, zero-range first paint, native drag
delivery and wheel routing outside the named help dispatcher population
remain Unknown. The bounded help wheel result does not establish global
wheel absence. — MENU-118, MENU-119, MENU-078

## Scrollable lists

SAVE/LOAD, Sound Options and cutscene builders construct shared list
table `L13377` through `R0762` or `R1183`, and shared bar class
`R1185`. These are three positive caller links, not a count of every
list in the image. — MENU-121

| Owner | List id | Initial list rectangle | Bar id |
|---|---:|---|---:|
| SAVE/LOAD | 3 | `(40,128,W-64,272)` | 10 |
| Sound Options | 3 | `(40,80,W-64,170)` | 10 |
| Cutscene chooser | 2 | builder rectangle, right reduced by 24 | `0x29b` |

Each passes font1. Pitch is font height plus 4; the list bottom becomes
`top+floor(height/pitch)*pitch+2`. Its attached bar is 24 wide and uses
that adjusted bottom. — MENU-120, MENU-121

Selection is clamped to the item count and brought into view. The bar
position/count pair is `(selection,count)`, rather than `(top,count-visible)`.
Negative selection resets top to zero and retains the negative index.
Focused Up/Down select one row. Page Up first selects the top visible row,
then targets `top-visible`; Page Down first selects the last visible row,
then targets `top+2*visible-1`. Home/End/Left/Right have no action in this
selected key table. Down selects and posts `0x46d`, up `0x472`, and
double-click `0x444`. Help uses a derived text control with a different
top-line setter/range; its contract is separate. — MENU-120, MENU-078, MENU-079

The saved header title is appended to the list string bank. Paint slot
`+2c`, `R2146`, calls row slot `+7c`, `R2147`, which retrieves the
string and draws it through `R0571`. Font1 slot `+14`, `R0767`,
reaches byte converter `R0793`. This establishes an in-image renderer;
native glyph pixels and malformed-label clipping remain Unknown.
— MENU-122

## Radio buttons and check boxes

The common art is `graphics.res::interface/radiob.256`. Normal radio and
checkbox frames are 24x24; the tip variant is 16x16. Coordinates below
are draw requests relative to absolute rectangle `(L,T,R,B)`, with row `i`.
Final sprite-wrapper placement and runtime label ramp colors remain Unknown.
— MENU-123, MENU-124, MENU-125, MENU-129

| Kind / painter | Clear / selected frame | Pitch | Body | Shadow | Label |
|---|---|---:|---|---|---|
| Radio `R2150` | 0 / 1 | 24 | `(L+1,T+24*i)` | `(L+5,T+24*i+4)` | `(L+30,T+24*i+5)` |
| Checkbox `R2153` | 2 / 3 | 24 | `(L+1,T+24*i)` | `(L+5,T+24*i+4)` | `(L+30,T+24*i+5)` |
| Tip checkbox `R2155` | 4 / 5 | 16 | `(L+1,T+16*i)` | `(L+5,T+16*i+4)` | `(L+22,T+16*i+3)` |

Body mode is 0; shadow mode 4. Disabled groups apply level-3 destination
remapping. Radio selection plus focus changes the selected row's label
ramp; standard checkbox focus changes the whole group's label ramp.
The tip painter tests each row under the pointer and selects its own
inside/outside ramps. No additional pressed sprite is selected in these
three local painters. — MENU-123, MENU-124, MENU-125

Radio down selects and sends `0x46d`; focused Up/Down change the local
selection and repaint without a local notification, while Space copies
the keyboard row and notifies. Standard checkbox down and focused Space
toggle a bit and notify; Up/Down move only its keyboard row. Release is
a no-op. Tip checkboxes inherit those inputs with a 16-pixel row selector.
Parent persistence after radio key-only changes remains Unknown.
— MENU-123, MENU-124, MENU-125

## Text input

Save Game constructs generic edit table `L06349`, painter `R2157`,
requested rectangle `(40,68,W-40,92)`, id 1, font1 and its supplied ramp.
Gold uses the same edit class. Character-generation name entry is a
different class with a glyph caret and ten-byte limit. — MENU-126,
MENU-088, MENU-090, TEXT-075, TEXT-076, TEXT-NAMEIN-024
MENU-088's explicit Backspace selection-reset clause is retracted.
TEXT-NAMEIN-024's former vtable base, slot index and caret-clock clauses
are partially retracted; the name-field limit and corrected class identity stand.

The generic edit asks its parent to repaint, then draws a procedural
bevel: top/left RGB `(8,8,8)`, bottom/right `(94,115,101)`. It loads
no frame art and has no own uniform background fill. Text starts at
`L+4`, vertically anchored at `T+trunc((B-T)/2)`, flag 8. — MENU-126

Selection requests level-12 remapping over the prefix-width interval,
from `T+2` to `B-2`, before text. It has no local focus gate. Focus plus
blink phase requests a white two-pixel caret at
`L+4+width(prefix(caret))`, over the same vertical interval. Message
`0x462` toggles phase when unsigned elapsed time exceeds 500 ms; reset
sets phase visible. Native visible blink interval includes delivery/repaint
cadence and remains Unknown. — MENU-126

Mouse down sets caret and selection endpoints by prefix-width hit testing;
held-left move changes selection. Left/Right/Home/End, Backspace/Delete
and Shift selection use the generic edit handlers. Native overflow clipping,
runtime glyph/ramp pixels, final Save rectangle and reopen/post-load state
remain Unknown. — MENU-126, MENU-091

## Window and dialog frames

Two concrete `lm.256` frame painters are identified in a bounded population
of 96 raw aligned static base-widget tables. This is not a global count
of every ROM1 window kind. Nonmatching or computed tables, slot `+2c`
decoration and supplied background-bitmap callers remain outside that count.
— MENU-127

| Family | Painter | Frame indices | Sizing | Selected owners |
|---|---|---|---|---|
| Normal decorative frame | `R0760` | 0..8 | normal-dialog snap below; Esc uses its own requested rectangle | Esc, dialogue, Save/Load, Game/Sound Options, cutscene chooser |
| Room-tip decorative frame | `R1975` | 9..17 | bypasses normal-dialog snap | four room-tip owners |

Normal dialogs snap `W'=trunc((W-8)/96)*96+8`,
`H'=trunc((H-104)/64)*64+104`, then center. At supplied 640x480, Game
Options is 488x424 at `(76,28)`, Sound Options 488x360 at `(76,60)` and
cutscene chooser 392x360 at `(124,60)`. These frame requests have 48-pixel
corners, 96x64 center tiles and an 8-pixel shadow band. — MENU-127,
MENU-077, MENU-ART-014, DLG-PANEL-035

Room-tip frames use 32-pixel corners, 48x32 center/horizontal tiles and
32x32 vertical tiles. Body right/bottom are reduced by 8. Interior counts
truncate `(bodyW-64)/48` and `(bodyH-64)/32`; right/bottom shadows are
requested first, offset `(8,8)`, mode 6. A nonzero normal-panel background
pointer instead paints the supplied bitmap, without local stretching.
The full owner map, bitmap-caller population and native cropping/remap
pixels remain Unknown. — MENU-127, MENU-129, TOWN-480

## Normal tooltips and hover boxes

Painter `R0371` requests green RGB `(36,44,39)` fill over
`(L+1,T+1,R-1,B-2)`, a two-color bevel `(160,120,50)` / `(80,60,24)`,
and 4x4 `Ball.bmp` corner copies at `(L,T)`, `(R-3,T)`, `(L,B-4)` and
`(R-3,B-4)`. These are requested colors/coordinates; native keying and
pixels remain Unknown. — MENU-128, MENU-129

Authored `#` breaks split lines. There is no local word wrap in the normal
split/layout/paint routines. Width is maximum measured line width plus 11,
height `14*lineCount+5`; font2 text uses inset `(5,4)` and pitch 14.
Placement starts above the pointer, shifts left for right overflow and
down for top overflow. No additional left/bottom clamp is present locally.
Oversized authored lines, other hover paths and runtime label ramp colors
remain Unknown. Stationary-pointer timing and hide transitions are separate.
— MENU-128, TEXT-HOVERPAINT-053, TEXT-HOVER-048

## Resource population and limits

The measured direct interface population excludes subdirectories: 87 EN
and 85 RU entries, with seven byte-identical `.256` banks and 54 indexed
frames per root. `lm.256` has 18, `radiob.256` 6, `scrlbars.256` 26;
`t_border.256` has one 88x108 frame. `Ball.bmp` is 4x4, 24-bit,
`t_back.bmp` 160x240, 24-bit. Indexed frames and appended data are counted
separately. A sprite header has no origin field; literal-pixel bounds are
not a widget origin. The metadata establishes this resource population,
not a complete widget/frame catalogue or native malformed-stream behavior.
— MENU-129
