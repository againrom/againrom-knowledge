# Claim registry — MENU (menu surfaces: the main menu, the documents panel, the in-play Esc menus)

Level 2 ledger. Index: [registry.md](registry.md) · spec: [`formats/menu/format.md`](../formats/menu/format.md). Legend: `✔ promoted` (also in the spec) · `● active` · `✖ retracted`. IDs are permanent.

This ledger covers the **composition contract** of a menu surface — which archive nodes form it,
what its hit regions mean, and where the engine places each element. It began with the main menu
(`MENU-ASSET-001`…`MENU-STATE-007`) and was extended on 2026-08-13 to a second full-screen bitmap
surface, the campaign documents panel (`MENU-DOC-009`), which is built the same way out of
`graphics.res` rather than `main.res`. On 2026-08-14 it was extended again, to the two in-play
menus Esc raises (`MENU-ESC-010`…`MENU-INPUT-016`). Those two are not bitmap surfaces: they carry
no art of their own and are drawn as a sprite nine-patch over a darkened frame, so the ledger's
subject is now the menu surface rather than the full-screen bitmap. The BMP *encoding* itself is
standard Windows BMP (out of scope to re-derive).

| ID | Claim | Confidence | Status | Evidence |
|----|-------|-----------|--------|----------|
| MENU-COMBAT-017 | The mission right column is one fixed 160-pixel container with four ordered children; the command panel is the second child and not the minimap or character panel. | High | ● active (partially retracted) | [EXP-0190](../experiments/EXP-0190-combat-controls/), [EXP-0476](../experiments/EXP-0476-backspace-descendants/) |
| MENU-COMBAT-018 | The eight command cells are 34x34 at panel-local `(8+34c,7+34r)`, row-major, and all visible states come from four full-panel BMPs. | High | ● active | [EXP-0190](../experiments/EXP-0190-combat-controls/) |
| MENU-COMBAT-019 | The eight cells and tooltip indices are one exact table: | High | ● active | [EXP-0190](../experiments/EXP-0190-combat-controls/) |
| MENU-ASSET-001 | The main menu is the 18-file `graphics/mainmenu/` subtree of `main.res`: a 640x480 24bpp base, a 640x480 8bpp hit mask, and two 24bpp overlay sets for buttons 1..8. | High | ✔ promoted | [EXP-0026](../experiments/EXP-0026-mainmenu-assets/) |
| MENU-ASSET-002 | There are **exactly 8 buttons**. | High | ✔ promoted | [EXP-0026](../experiments/EXP-0026-mainmenu-assets/) |
| MENU-MASK-003 | `menumask.bmp`'s 256-entry palette is the **identity grayscale ramp** `palette[i]=(i,i,i)` (all 256 entries) — it carries no colour meaning; the raw **8-bit index** is the semantic marker. | High | ✔ promoted | [EXP-0027](../experiments/EXP-0027-menu-hitmask/) |
| MENU-MASK-004 | The hit-test `R1148` reads the mask byte under the cursor and maps eight index values, 0x80 to 0xf0 in steps of 0x10, to buttons 1..8; every other index is no button. | High | ✔ promoted | [EXP-0027](../experiments/EXP-0027-menu-hitmask/) |
| MENU-GEOM-005 | Overlay placement is **two static `rom.exe` tables**: normal/hover at the global `L06179` (VA; file `0x198ed8`) and pressed at the global `L06180`, contiguous, 8 entries × 16 B = `{x,y,w,h}` int32 LE. | High / Medium | ✔ promoted | [EXP-0028](../experiments/EXP-0028-menu-overlay-geometry/) |
| MENU-GEOM-006 | Overlays draw **1:1** (no scale): each table entry's `(w,h)` equals the corresponding BMP's pixel dimensions — **16/16** exact (normal + pressed). | High | ✔ promoted | [EXP-0028](../experiments/EXP-0028-menu-overlay-geometry/) |
| MENU-STATE-007 | `button%d.bmp` is the **hover/highlight** overlay, `button%dp.bmp` the **pressed** overlay. | Medium | ● active | [EXP-0027](../experiments/EXP-0027-menu-hitmask/) |
| MENU-STRTAB-008 | Which surfaces the one global string index space names, from the sites read this round — the map a consumer needs before it can renumber anything. | Medium | ● active | [EXP-0143](../experiments/EXP-0143-actor-names/) |
| MENU-DOC-009 | The documents panel, whole: eleven bitmaps, three hit rectangles that equal their bitmaps, 21-line pages, and one entry point. | High / Medium | ● active | [EXP-0151](../experiments/EXP-0151-mission-documents/) |
| MENU-ESC-010 | Esc during play raises two different surfaces, one per UI state, and both are panels of one family rather than the main menu. | High | ✔ promoted (partially retracted) | [EXP-0162](../experiments/EXP-0162-ingame-menu/) |
| MENU-ITEM-011 | The mission menu constructs eight entries and shows seven; two of them are exclusive. | High / Medium | ✔ promoted | [EXP-0162](../experiments/EXP-0162-ingame-menu/) |
| MENU-ITEM-012 | The town menu is a different class with five entries, and one of them is a different action with a different word. | High / Unknown | ✔ promoted | [EXP-0162](../experiments/EXP-0162-ingame-menu/) |
| MENU-KEY-013 | The effective accelerator is the letter after `~` in the label, not the constructor's immediate, and the difference is visible in the shipped English text. | High | ✔ promoted | [EXP-0162](../experiments/EXP-0162-ingame-menu/) |
| MENU-ART-014 | Both Esc menus load no art of their own; the frame is a nine-patch out of one sprite bank, `graphics\interface\lm.256`, plus an 8-px drop shadow. | High / Unknown | ✔ promoted | [EXP-0162](../experiments/EXP-0162-ingame-menu/) |
| MENU-STOP-015 | Both Esc menus stop the world, and the mechanism is the idle-handler gate, not `MISSION-STOP-016`'s `server+0x2c`. | High | ● active | [EXP-0162](../experiments/EXP-0162-ingame-menu/) |
| MENU-INPUT-016 | Esc opens and Esc closes; the panel takes keyboard focus among its own rows and, as the root's capture object, receives every mouse message while it is up. | High / Unknown | ● active (partially retracted) | [EXP-0162](../experiments/EXP-0162-ingame-menu/) |
| MENU-CURSOR-046 | The main-menu transition is the one surface transition that ends on `select` rather than `default`. | High / Medium | ● active | [EXP-0214](../experiments/EXP-0214-cursor-surfaces/) |

### MENU-COMBAT-017

**The mission right column is one fixed 160-pixel container with four ordered children; the command panel is the second child and not the minimap or character panel.** `R0315` constructs root `campaign+0xcc`, map `+d0`, right column `+d4`, minimap `+d8`, command panel `+dc`, character panel `+e0` and filler `+e4`; `R1149` appends `d8,dc,e0,e4` under `d4`, then `d0,d4` under the root. Command-panel screen rectangles are `(480,158)-(640,238)`, `(640,158)-(800,238)` and `(864,158)-(1024,238)` at the three shipped resolutions. The command-panel class is `R1150`, vtable `L01177`, id 6, 160x80, with no children. Inventory, spellbook and Esc hit regions belong to the separate character-panel sibling. The inventory-grid double-click can enter the same item Cast/transfer vmethod as character-panel left-up. The measured inventory-open route appends that grid to the map view, not the character panel (MENU-088). Its physical left-up separately transfers the selected item through grid vslot `+0xa4` to destination code 2 and emits `0x22/0x32`, including when the hit-test returns `-1`. A gold-sentinel cell with a positive purse instead attaches the always-created Drop Gold modal at `campaign+0x10c`, rect `(100,H-200)-(396,H-32)`: edit id `0x989685`, action `0x989681`/`0x445`, cancel `0x989682`/`0x446`; Enter/Esc alias those buttons

**Confidence.** High (construction arguments, append order, inventory/dialog/button vtables, physical handlers and all child constructors are read end to end; the rectangles are exact evaluation of hard-coded arguments)

**Amended.** The former nested-owner clause is partially retracted. MENU-088 establishes the map-view parent on the measured opening route; geometry, item/money handlers and command-panel ownership stand.

### MENU-COMBAT-018

**The eight command cells are 34x34 at panel-local `(8+34c,7+34r)`, row-major, and all visible states come from four full-panel BMPs.** Inactive draws `HeadsR.bmp`; active draws `CommandBarR.bmp`, then each disabled cell's same source rectangle from `CommandEmpR.bmp`, then the selected cell from `CommandDnR.bmp`. There is no hover bitmap, separate border or separate icon, and no child button object. All four assets are 160x80x24 uncompressed BMPs, byte-identical across EN/RU. Active, enable-mask, selected-index and dirty state are panel fields `+0x68,+0x64,+0x60,+0x6c`. A margin or disabled left-down clears only the selected overlay and sends index -1, leaving a previously armed map mode intact

**Confidence.** High (one draw routine and one left-down handler read whole, plus BMP-header and hash measurements on both roots)

### MENU-COMBAT-019

**The eight cells and tooltip indices are one exact table:** Attack/A, Move/M, Guard/G, Defend/D, Cast/C, Swarm/S, Stand Ground/T, Retreat/R. Cell 0..7 returns `main.txt[0..7]` only when active, enabled, no modal/special cursor is present and frame-state bits `0x0a` are clear. The RU file carries the corresponding CP866 strings `Атаковать`, `Идти`, `Охранять`, `Защищать`, `Колдовать`, `Идти в боевой готовности`, `Держать позицию`, `Отступить`, with the same Latin accelerators. Left down acts; double-click aliases it; left-drag re-enters it on every delivered move; right up cancels through map message `0x405`; the other right edges are no-ops. **G2:** cell geometry and index order are `.text`; changing them requires executable edits. Text and the four full-panel images are data, so changing a label or any embedded icon changes only its shipped resource bytes

**Confidence.** High for the event routes, predicates, indices and EN/RU bytes (handlers read whole and both resource roots measured)

### MENU-ASSET-001

The main menu is the **18-file** `graphics/mainmenu/` subtree of `main.res`: `menu_.bmp` (640×480, 24bpp base brooch background), `menumask.bmp` (640×480, **8bpp** hit mask), `button1..8.bmp` (24bpp hover overlays) and `button1..8p.bmp` (24bpp pressed overlays). All 18 sha256 distinct; overlay dims per `inventory.csv`. `rom.exe` requests them as `main\graphics\MainMenu\{menu_.bmp, MenuMask.bmp, button%d.bmp, button%dp.bmp}` (referenced only by the loader `R1151`)

**Confidence.** High (the asset set is not inferred from the directory listing but from the **only** referrer of those path strings; "18 files" is then a listing)

### MENU-ASSET-002

There are **exactly 8 buttons**. The loader `R1151` `sprintf`s `button%d.bmp`/`button%dp.bmp` in a loop `n = 1..8` (`do{…}while(n<8)`, start 1), storing hover overlays into the array at `this+0x6c` and pressed overlays at `this+0x80`, mask at `this+0xbc`, background at `this+0xb8`

**Confidence.** High (the count is the loop bound itself — `do{…}while(n<8)` from 1 — not a count of files that happen to exist)

### MENU-MASK-003

`menumask.bmp`'s 256-entry palette is the **identity grayscale ramp** `palette[i]=(i,i,i)` (all 256 entries) — it carries no colour meaning; the raw **8-bit index** is the semantic marker. The mask is loaded by `R1152` as raw indices (buffer at maskObj+0x10, `w·h` bytes, never rendered to colour)

**Confidence.** High (the identity ramp is a measurement over all 256 entries; "the index is the semantic marker" is carried by the loader, which keeps the buffer as raw indices — an absence of any colour conversion on that path)

### MENU-MASK-004

The hit-test `R1148` reads `idx = maskBuf[(y−top)·640 + (x−left)]` (640 = mask width) and maps eight index values to buttons: **0x80→button1, 0x90→button2, 0xa0→button3, 0xb0→button4, 0xc0→button5, 0xd0→button6, 0xe0→button7, 0xf0→button8** (asset number = idx/16 − 7). Index 0 (background) and every other value (incl. the edge ramp 0x10..0x1e = hot>>3 and stray AA pixels) fall to the switch default = **no button**. Verified: only these 8 of the 43 present indices are "hot", and each region is bracketed by its button's placement rect 8/8

**Confidence.** High (the eight mappings are `switch` cases read off the hit-test, and the *default* — every other index, including the edge ramp — is read too, which is what makes "no button" a fact rather than an assumption)

### MENU-GEOM-005

Overlay placement is **two static `rom.exe` tables**: normal/hover at the global `L06179` (VA; file `0x198ed8`) and pressed at the global `L06180`, contiguous, 8 entries × 16 B = `{x,y,w,h}` int32 LE. `R1153` `SetRect`s each button at `(x,y)`→`(x+w,y+h)` offset by the menu origin. Origin = (0,0) (full-screen 640×480 mask; hit-test indexes it from (0,0)). Values listed in `placement-tables.csv` / `rom-layout-excerpt.md`

**Confidence.** High (static tables read at named addresses and the `SetRect` that consumes them) / Medium (the origin `(0,0)` — deduced from the mask being full-screen and indexed from `(0,0)`, not from an instruction that sets it)

### MENU-GEOM-006

Overlays draw **1:1** (no scale): each table entry's `(w,h)` equals the corresponding BMP's pixel dimensions — **16/16** exact (normal + pressed). Each normal rect **brackets** its button's mask region — **8/8**. These two discriminating checks jointly bind BMP dims ↔ hit-test index ↔ placement table into one consistent geometry

**Confidence.** High (16/16 exact and 8/8 bracketing across three independently-derived sources — dims from the files, indices from the mask, rects from the binary; a wrong pairing anywhere breaks one of the two. This row is the shape the confidence scale asks for and states its own discriminator without prompting)

### MENU-STATE-007

`button%d.bmp` is the **hover/highlight** overlay, `button%dp.bmp` the **pressed** overlay. In `R1148`: not-pressed & hovering button i → draw `this+0x6c[i]` at normal rect `this+0x94[i]`; mouse-down while still hovering the latched button i → draw `this+0x80[i]` at pressed rect `this+0xa8[i]`. `this+0xe4` is a per-button **disable bitfield** — bit i set suppresses the overlay. The click dispatcher `R0820` posts a per-button command message (button 8 → WM_CLOSE 0x10); message-id meanings beyond WM_CLOSE are undecoded

**Confidence.** Medium

### MENU-STRTAB-008

**Which surfaces the one global string index space names, from the sites read this round — the map a consumer needs before it can renumber anything.** Every entry loads the string-table pointer at `L04369` and then subscripts it by `4 × index` (`TEXT-STRTAB-023`). Index **77** is the dialogue panel's button (`L06181`, `R0695`). Indices **140/141** are the mission-outcome panels (`L06182`, `L06183`). Indices **47..50** are the multi-selection panel, `R0389` reading `+0xbc/+0xc0/+0xc4/+0xc8` — a selected-unit **count**, not names — and the same routine reads `class+0xd8` (`InfoPicture`) off the `units.reg` class array. Index **256** is the text-entry field's tooltip, returned by `R1154` when its owner's `+0x1e0` is non-zero, and that class is `TEXT-NAMEIN-024`'s ten-character name field. The in-mission message layer `R0509` resolves **26** distinct indices — 85..89, 129, 142..149, 204..209, 221..226 — a contiguous-block pattern that matches item-pickup, alliance, join/rejoin and cheat messages, all of which concatenate a **player** name. `EnumRefs refto:L04369` returns 230 hits over **47** owners, so this map covers a minority of the consumers: the remaining sites fold the index into a memory displacement and were not resolved to an index this round. The *other* fifteen tables are not reached this way at all — they are read through `R0668(table, i)` with a table-local index, 302 hits over 57 owners — and the two surfaces this round followed there are the unit information panel reading `unitname.txt` (`UNIT-NAME-039`) and the character screen reading `npcnames.txt` at `R1155` `L06184`/`L06185` under an index built from the mage bit `+0x18c & 4`

**Confidence.** **Medium.** Each named index is a cited instruction and each agrees with the EN text at that line of `text/main.txt`, which is corroboration across two artefacts rather than a read of the drawing call; and the coverage is explicitly partial — 5 of 47 owners mapped. The three name-validation messages at 193/194/195 have **no** located consumer

### MENU-DOC-009

**The documents panel, whole: eleven bitmaps, three hit rectangles that equal their bitmaps, 21-line pages, and one entry point.** `R1156` is the only referrer of any of the six document art path literals and loads them all: `+0x6c` = `graphics\interface\Docs\sheet.bmp`, a 3-element array at `+0x74` = `Arrows\{00_l,01_l,11_l}`, a 3-element array at `+0x88` = `Arrows\{00_r,01_r,11_r}`, a 4-element array at `+0x9c` = `OK\{Ok_off,Ok_on,Ok_l_off,Ok_l_on}`. Geometry is `R1157`, in screen pixels: left arrow `(0,200,56,240)`, right arrow `(576,200,636,240)`, OK `(560,416,604,448)` — each **equal to its bitmap's pixel size**, 56×40 / 60×40 / 44×32, measured from the BMP headers on both roots, which checks the two readings against each other. `sheet.bmp` is 640×480×24 and `1.bmp` 464×344×24. `R1158` draws sheet, left arrow, right arrow, OK, then the current document at the panel origin. A picture is blitted at the document RECT's top-left `(92,72)` at natural size, **8 px wider than the 456-px RECT**; text is drawn by `R1159` with `ECX = [L06186]`, `TEXT-API-007`'s font4, from line `+0x38` to `+0x38+0x15` with the pitch taken from the glyph sprite because the caller passes 0. `R1160` selects array element 1 on hover and 2 on press (3 for OK) and 0 on leave, so the arrow file names read as `00` idle / `01` hover / `11` pressed; `Ok_on` is selected by no arm read here. Paging is `R1161`/`R1162`, a step of **`0x15` = 21 lines** on `+0x38` bounded by the line count at `+0x18`, and `R1163`/`R1164` change **document** only when the page step returns 0, so one arrow pair walks pages first and documents second. OK posts `0x445`. The panel is 248 bytes, window id `0x4ba`, constructed at `L06187` into `app+0x3a4`; `R1165` is its only shower, called once, from the campaign state machine `R0701` at `L06188` under `campaign+0x3dc == 1`, on the arm of **message `0x463`** — one of 112, read out of the PE's own index and jump tables on both roots

**Confidence.** High for the asset set (the only referrer of the literals), the geometry (immediates, cross-checked against BMP headers on both roots) and the page step / High for "one message raises it": the switch is read from the image's own tables rather than a listing, and both roots agree / **Medium** for "one entry point in the image": a control with the numerically equal id `0x463` exists in `R1166`'s own id space and whether its notification can reach this handler was not ruled out

### MENU-ESC-010

**Esc during play raises two different surfaces, one per UI state, and both are panels of one family rather than the main menu.** `R0231` is the frame window's `WM_KEYDOWN` handler — read off the MFC message map at `L06189`, whose six dwords are `{0x100, 0, 0, 0, 0x10, R0231}` and which reproduces on both roots. Its `VK_ESCAPE` arm loads the UI state word (`L03621`) and branches: `== 1` posts **`0x416`** (compare at `L06190`, message constant pushed at `L06191`), `== 0` **and** the `campaign+0x3b4` CString empty posts **`0x41f`** (state compared with zero at `L06192`, the CString length field at offset -8 compared with zero at `L06193`, message constant pushed at `L06194`), anything else falls to `L03620`, which forwards `WM_KEYDOWN` to the root container `campaign+0xcc`. `SESS-SCREEN-003` fixes `campaign+0x3dc == 1` as a map session on screen and `SHOP-TOWN-023` fixes `== 0` as the town. `R0701` case `0x416` re-checks `== 1` and builds `R0362` at `(100,60)-(440,400)` = **340×340** (the four rectangle arguments 0x190, 0x1b8, 0x3c and 0x64 are pushed at `L01477`); case `0x41f` is ungated and builds `R0363` at `(100,100)-(440,340)` = **340×240** (the four rectangle arguments 0x154, 0x1b8, 0x64 and 0x64 are pushed at `L01479`). Both then take the shared tail at `L01471` that calls `R0361`. The same `0x416` is posted by the ~~command panel's own~~ *character-panel sibling's* 32×32 button at panel-local `(0x7e,0xce)-(0x9e,0xee)` (`L06195`, in `R0332`, gated `campaign+0x3dc & 1`), whose tooltip `R0734` resolves to global string index 14 — `text/main.txt` line 15, `Main Menu <ESC>` on the EN root and the same line in Russian on the RU root. Neither path reaches the main-menu window: `callto:L03833` is **1 hit / 1 owner / 0 orphan** (`L06196`), and the routine that shows that window, `R0816`, first frees the session's object lists and plays `music\menu.wav` — it is the exit path, not a pause

**Confidence.** High for the handler identity (the message map is data, not a decompiler rendering, and both roots' bytes agree), the two branch conditions, both rects (immediates verified through the PE section table on both roots, `evidence/anchors.txt`) and the single main-menu construction (enumeration instrument named, 0 orphan) / the *meaning* of `campaign+0x3dc == 1` and `== 0` is taken from `SESS-SCREEN-003` and `SHOP-TOWN-023`, not re-derived here

**Amended.** The owning-panel noun is partially retracted; the button rectangle, tooltip and posted message stand. The correction is in `retracted.md` and `MENU-COMBAT-017`.

### MENU-ITEM-011

**The mission menu constructs eight entries and shows seven; two of them are exclusive.** `R0362` builds each row as `R1167(id, R0668(index), font [L03615], 0, message, accelerator, colourPtr)` and stacks it with `R1168(button, 0x1e)`. In screen order, with `text/dialogs.txt` line and posted message: `~Save Game` `0x41a`; then either `~Load Game` `0x418` when `campaign+0x6bc == 2` **or** `Diplomacy` `0x43c` otherwise — the two write the same local and only one `R1168` follows; then `Game ~Options` `0x41b`, `Sou~nd Options` `0x422`, `~Quest Objectives` `0x420`, `~End Quest` `0x41c`, `~Return to Game` `0x446`. Label indices are `0x22 0x23 0x4c 0x24 0x25 0x26 0x27 0x28` on the descriptor `L06197`, which the loader at `L06198`/`L06199` fills from `main\text\dialogs.txt` (path string `L06200`, descriptor object `L06197`), and `R0668(table,i)` resolves as `[L04369][table->+0xc + i]`. Four rows carry a disable: `~Save Game` when `campaign+0x6bc` is 0 or 1; `~Load Game` when `R1169()` is 0, and that routine is a `FindFirstFileA` on `<game dir>\game*.sav`; `Sou~nd Options` when `R1170()` returns 5, its `this+0x9c == 0` arm; `~Quest Objectives` when `campaign+0x6bc != 2`. The **first argument** to `R1167` is a control id, **not** the row: it runs 1,2,3,4,5,6,8 with 7 taken by `Diplomacy`, while the row order is the `R1168` call order above. `0x446` is the family-wide close, not a distinct action: `R0716` turns `0x445`/`0x446` into the standard `0x44c` teardown, and 39 owners in the image post it

**Confidence.** High for the entry list, its order, the messages, the label file binding and the exclusive pair (one routine read whole, the loader binding two adjacent instructions, and the labels re-read out of both roots' `main.res` where the RU strings are the same menu in different bytes) / **Medium** for the four disable rules: each is one cited call whose predicate was read, but whether any other site re-enables a row later was not swept

### MENU-ITEM-012

**The town menu is a different class with five entries, and one of them is a different action with a different word.** `R0363`, vtable `L06201` against the mission menu's `L06202`, same button and row helpers. In screen order: `~Save Game` `0x41a`, `~Load Game` `0x418`, `Sou~nd Options` `0x422`, **`Abort Game`** (label index `0x4d`) `0x41c`, `~Return to Game` `0x446`. No `Game Options`, no `Quest Objectives`, no `Diplomacy`, and Save precedes Load where the constructor's control ids are 2 and 1 — another instance of `MENU-ITEM-011`'s id-is-not-row point. A search of the constructor for the disable call `(*button)vt+0x1c(1,0)` returns **0**: every row is built enabled. `0x41c` is the same message the mission menu's `~End Quest` posts; `R0701`'s `0x41c` arm branches on `campaign+0x3dc`, and the two surfaces take **different** branches of it: `== 1` raises the five-row confirmation `R1171` (`Change Map` or `~Victory!` at `0x41d`, `~Exit to Main Menu` and `Exit to ~Windows` both at `0x41e`, `~Return to Game` at `0x446`), `== 0` raises the three-row `R1172` (the same two exit rows and `~Return to Game`), both at `(100,100)-(440,340)`. So one message, two panels, and the mission row's own word for it is `~End Quest` while the town's is `Abort Game`

**Confidence.** High for the entry list, its order, the messages, the absence of disables and the two confirmation panels' own rows (three routines read whole, labels re-read on both roots) / **Unknown** for what the two exit rows do: both post `0x41e` and differ only in control id, and how the exit target is distinguished was not established

### MENU-KEY-013

**The effective accelerator is the letter after `~` in the label, not the constructor's immediate, and the difference is visible in the shipped English text.** `R0700` stores the passed accelerator, then walks the label from index 1: `~~` is an escape and skips two, otherwise the first character preceded by a single `~` replaces it via `R1173(ch) & 0xff`. `R1173` lowercases — `R1174` normally, and when `[L06203] == 1` it lowercases CP866 Cyrillic instead (`+0x20` over `0x80..0x8f`, `+0x50` over `0x90..0x9f`). Consequences measured on both roots (`evidence/labels.txt`): on the EN root eleven of thirteen labels carry a `~` and take it, and the two that do not — `Diplomacy` and `Abort Game` — keep the constructor's `D` and `E`; on the RU root **all thirteen** carry a `~`, including the rows those two correspond to, so those two entries have a language-dependent accelerator while the code has not changed. The RU labels also place the `~` away from the first letter where the first letters would collide — mission rows 5, 6 and 7 mark bytes `0xa0`, `0xaa` and `0xad` at label offsets 1, 2 and 3 — which the immediate model cannot produce at all. One residue: mission row 5 passes `0x4d` = `M` and its EN label is `~Quest Objectives`, so the shipped accelerator is `Q` and the immediate is inert. A reimplementation that hard-codes the accelerators reproduces the English release and diverges on the Russian one

**Confidence.** High. The scan is one routine read whole, and the discriminator is not corpus agreement: the two roots carry disjoint label bytes and disagree about which entries carry a `~` at all, which no single-root reading could have separated from *the immediate is the accelerator*

### MENU-ART-014

**Both Esc menus load no art of their own; the frame is a nine-patch out of one sprite bank, `graphics\interface\lm.256`, plus an 8-px drop shadow.** The family's art-load slot `vt+0x78` is `R1175`, whose whole body is `RET`. Painting is `R0760` through the global `[L03604]`, which has exactly **one writer and one reader**: `refto:L03604` = 29 hits / 2 owners / 0 orphan, the write at `L06204` (the constructed object stored into `L03604`) in the interface loader `R1057` and 28 reads all inside `R0760`. The store pairs with the construct that precedes it, the path string `L06205` (`graphics\interface\lm.256`) pushed at `L06206` into the `.256` constructor `R0753` (its two neighbours use the `.bmp` constructor `R1176`). Frames 0..8 are the nine-patch and their measured sizes match the placement arithmetic exactly: centre `96×64`, corners `48×48`, top and bottom edges `96×48`, left and right edges `48×64`, against an inset of `0x30` = 48, a horizontal step of `0x60` = 96 and a vertical step of `0x40` = 64. Frames 9..17 are a second nine-patch of the same shape at 32/48 px and are not referenced by `R0760`. The file is 39 943 bytes with the same sha256 on both roots. Two passes: `vt+0x1c` draws frames 3,6,8,7,5 — the right column and bottom row only — over a rect offset `+8,+8`, then `vt+0x18` draws all nine at the panel rect shrunk by 8 on right and bottom. Rows are laid out by `R1168`: `panel+0x70 = panel+0x78`, `panel+0x78 = top + 0x1e`, then `button->SetRect(&panel+0x6c)` and `panel->AddChild(button)`, with `R1177` forcing `panel+0x6c = 0x28` and `panel+0x74 = width − 0x30` and seeding top/bottom from `(0,0,0xf0,0x28)` — so each row is panel-local `(40, 40+30(n−1), width−48, 40+30n)`

**Confidence.** High for the bank's identity, its exclusivity, the frame sizes and the row layout (the enumeration names its instrument and returns one writer with 0 orphan; the frame sizes are read from the `.256` structure on both roots and agree independently with the code's immediates; the row helper is nine instructions read raw) / **Unknown** for how the tiling covers a frame width that is not `96 + 96k`: the counts are `(w−96)/96` and `(h−96)/64` with truncating division (`L06207`: signed division of the extent minus 0x60 by the tile size 0x60), which for this panel's 332×332 frame leaves a 44-px band inside the right edge and one inside the bottom edge that no placed tile accounts for. EXP-0162 states this as a prediction with its refutation rather than as a claim

### MENU-STOP-015

**Both Esc menus stop the world, and the mechanism is the idle-handler gate, not `MISSION-STOP-016`'s `server+0x2c`.** Both `R0701` arms end at `R0361`, which sets bit 3 of `campaign+0x3dc` (the OR with 0x8 at `L03329`) — or bit 15 for the one panel stored at `campaign+0x3a8`, which is not either of these — and the MFC idle handler tests it once in the whole image: `L03333` (a test of the word against the mask 0x4008), `EnumRefs imm:4008` = **1 hit / 1 owner / 0 orphan**, verified byte-for-byte on both roots. `DLG-STOP-012` measured that branch: it runs no pacer arm at all, which is stricter than the pause, and the timer rival is closed by an import-table absence (`SESS-TIMER-022`). Neither Esc path writes, reads or reaches `server+0x2c`; `MISSION-STOP-016`'s two writers are `L06208` and `L06209`, in the session-start and mission-teardown routines, and neither is on this path. Neither panel is stored in a `campaign+…` slot, so its close takes `R0709`'s default arm, and `DLG-CLOCK-014` fixes what that does: clear bit 3, and if the word is then exactly 1 and the pacer phase is 2, zero the phase so the stopped time is **discarded** rather than caught up. In the town the word returns to 0 and the restart correctly does nothing. Compositing is `DLG-DIM-013` unchanged: between a DirectDraw Lock and Unlock, the level argument 3 pushed at `L03337` and the call to `R0732` at `L03338` remap **the whole screen** in place through the shroud table `[L03341]` at level 3, a per-channel gain of 13/16 by `TERR-FOG-084`'s law — one destructive pass, not a per-frame blend and not a translucent draw, and it survives only because this gate stops everything that would repaint

**Confidence.** High for the gate, its single-hit enumeration, the shade level and the absence of `server+0x2c` from this path (each a named instruction, the immediates re-read through the PE section table on both roots) / the *consequences* of the gate and of the close arm are `DLG-STOP-012` and `DLG-CLOCK-014`, both High, applied here rather than re-derived; this row establishes that the Esc menus take that path, not that the path behaves as those rows say

### MENU-INPUT-016

**Esc opens and Esc closes; the panel takes keyboard focus among its own rows and, as the root's capture object, receives every mouse message while it is up.** With the panel up `campaign+0x3dc` is 9, so neither arm of `R0231`'s `VK_ESCAPE` case matches and the key is forwarded to the root container, reaching the panel's key slot `R1178` — `0x26`/`0x28` move the focused row, everything else falls to `R0818`, which on `0x1b` posts `0x446`, the same message `~Return to Game` posts. The same word being 9 also makes the character panel's Main Menu button inert while the panel is up, because `R0701`'s `0x416` arm requires exactly 1. `R0361` adds the panel as the **last child of `campaign+0xcc`** (`L06210` loads that parent from offset 0xcc), the same parent the map view and the side panels have, and `R0385` only appends and sets the parent pointer. `R0390`, the base dispatcher, sends a mouse message to the capture object at `this+0x34` when one is set and otherwise walks the children, delivering to the first whose rect contains the point and stopping there; ~~`R0775`, the panel's show, sets `this+0x5c = 1` and calls `R0787(1,0)`, which moves focus among the panel's **own** children and sets no capture. A click outside the panel is therefore delivered to the window under the cursor~~ *`R0775`, the panel's show, sets `this+0x5c = 1`, calls `vt+0x24(1)` (`L06211`), which is `R0776` in both class tables and stores the panel in `root+0x34`, then `vt+0x28(1)` and `R0787(1,0)`. `R0390` gives a mouse message to `root+0x34` before any child, the root's own mouse slots are stubs, and the panel's children do not contain a point outside it, so a click outside the panel reaches no other child of the root.*

**Confidence.** High for the open/close path and for the routing rule (three routines read whole, and the forwarding target `L03620` is fixed by the `JNZ rel32` in the same listing) / **Unknown** for ~~what a click outside then *does*: the map view's and side panels' own handlers were not read this round, so this row claims delivery and not effect. EXP-0162 carries that as an open question~~ *what the panel's own handlers do beyond the slots read*

**Amended.** The mouse-capture clause of the headline and the last two body sentences are partially retracted; the open and close path, key forwarding, the `campaign+0x3dc` value 9, the append as last child of `campaign+0xcc` and the dispatcher rule stand. The correction is in `retracted.md`.

### MENU-CURSOR-046

**The main-menu transition is the one surface transition that ends on `select` rather than `default`.** `R0816` is identified as the main-menu routine by its own push of the string `L06212` at `L06213`, which is `music\menu.wav`; the push site was found by a whole-image dword scan of the string block rather than by a literal heuristic. It sets `campaign+0x3dc` bit 7 (an OR with 0x80 at `L06214`, `TOWN-373`), loads `wait` (slot 26) at `L06215` and calls the set-cursor adapter at `L06216`, and at its tail loads **`select`** (slot 5, `L06217`) at `L06218` and calls the adapter at `L06219`. The other eleven routines of the same family end on `default` (slot 0) instead (`TOWN-372`). The same routine also pushes `World\Data\` at `L06220`, `-serverid` at `L06221`, `.srv` at `L06222` and `default.srv` at `L06223`, so its range extends to at least `L06223`

**Confidence.** High for the two cursor sets and the bit, all named instructions in the complete 69-caller population (`AI-CURSOR-175`), and for the slot identities, which come from `SPR16A-CURSOR-067`. **Medium** for the surface identification: it rests on the routine pushing the menu music path, and the alternative that `music\menu.wav` is loaded by some other screen is not excluded by the literal alone


`EXP-0214` was allocated `MENU-CURSOR-046`..`MENU-CURSOR-050` (5 ids) and spent
`MENU-CURSOR-046`, 1 id. **`MENU-CURSOR-047`..`MENU-CURSOR-050` (4 ids) are returned
unused**, none ever reissued. The next free `menu.md` id is therefore `MENU-CURSOR-051`.


## Help panel and Cast key

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MENU-051 | The campaign frame's key handler posts help message `0x434` on F1 in any state; the frame builds the help panel only when `campaign+0x3dc` equals exactly 1, so F1 is ignored in the town and over any panel that sets an overlay bit. | High | ● active (amended) | [EXP-0437](../experiments/EXP-0437-f1-help/) |
| MENU-052 | Help is the Pause panel's class and constructor call with the help text in place of the pause line: id 1, rectangle 576x384 snapped to 488x360 and centred, one OK button, no title, scrolling body in font 1. | High / Medium | ● active (amended) | [EXP-0437](../experiments/EXP-0437-f1-help/) |
| MENU-053 | Help stops the world through the Esc menus' idle gate and closes on the OK button or Esc; F1 over help is ignored; the scroll keys act while the text control holds focus (`MENU-078`). | High / Medium | ● active (amended) | [EXP-0437](../experiments/EXP-0437-f1-help/) |
| MENU-054 | The C key reaches the Cast arm only in a map session with no text entry open, a nonzero selection count and `view+0x144 & 0x24` clear; otherwise it falls to a second dispatch whose C entry returns 0. It reaches nothing in the town. | High | ● active (amended) | [EXP-0437](../experiments/EXP-0437-f1-help/) |
| MENU-055 | C arms Cast mode 5 only when a selected object set `view+0x144` bit `0x200`; for a nonempty selection with the bit clear the key is consumed with no message, sound call or state change. An armed Cast posts `0x408` and chooses no spell. | High / Medium / Unknown | ● active (amended) | [EXP-0437](../experiments/EXP-0437-f1-help/) |
| MENU-056 | F1 in the campaign frame reaches no accelerator, help file or WinHelp call; its only path is the `text/help.txt` panel. The library `ID_HELP` machinery is in the image and untraced. Bounded to the population named. | Medium | ● active (amended) | [EXP-0437](../experiments/EXP-0437-f1-help/) |

### MENU-051

- The F1 arm of the frame's `WM_KEYDOWN` handler `R0231` is a `PostMessageA` call (`L06224`, import slot `L03535`) of message `0x434` (`L06225`) to the frame window handle at frame+0x1c (loaded at `L06226`), whose two remaining arguments are the same value, then a jump to `L06227`. No compare on `campaign+0x3dc` or `campaign+0x6bc` precedes it. Pause (key `0x13`) tests `+0x6bc == 2` and `+0x3dc == 1` before it acts, and F2 tests `+0x6bc == 2` and `+0x3dc` in {0, 1}, so the absence is a property of this arm.
- The frame procedure `R0701` case `0x434` is guarded at `L06228` by a compare of the word at offset 0x3dc with the constant 1 (set at `L06229`), branching to `L06230` when they differ. A different word leaves the case without building a panel or posting a message.
- Accepted: word 1, the map-session-only state of `SESS-SCREEN-003`. Ignored: word 0, the town (`SHOP-TOWN-023`), and every word with an overlay bit, which includes the Esc menu and Pause panel (exactly 9, `MENU-STOP-015`), any dialogue panel (`DLG-CLOCK-014`) and the help panel itself (`MENU-053`).
- The identity of the F1 key with this arm (the virtual-key compare and its dispatch) rests on `AI-KEY-125` and its `keyboard.tsv`; the committed listing of `R0231` holds the arm.
- `rom.exe` is byte-identical on both roots (SHA-256 `942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`), so EN and RU share the rule.

**Confidence.** High for both arms: each is a named instruction and the compare operand is fixed by the listing's raw bytes. The live alternatives named by the question (accepted in every state, accepted only in a map session, accepted in a map session and the town) are excluded by the missing compare in the key arm and the exact-equality compare in the frame case.

**Unknown.** Every value `campaign+0x3dc` can take: `SESS-SCREEN-003` leaves bits 1, 2, 9, 10 and 13 individually Unknown, so the ignored set is stated by the one accepted value. Frame states in which `R0231` is not the key handler (the main menu surface) were not examined. Whether the Drop Gold modal (`MENU-COMBAT-017`), the chat text entry (`campaign+0xc0`) and the spellbook popup change the word: none is shown to, so F1 may open help over any of them.

**Amended.** The Unknown on the Drop Gold modal, the chat entry and the spellbook popup is answered by `MENU-080`: the modal sets bit 3 and blocks F1 (High); the entry and the popup were found to leave the word at 1, so F1 should open help over them (Medium: absence over the named routines and the displacement census, indirect and register-indexed stores and bulk copies not excluded). The save and load chooser panels also set bit 3, so F1 is ignored there. The accepted and ignored sets above stand.

### MENU-052

- Construction: from `L06231` the global `L06232` is the receiver, the arguments are `0`, `0`, the text pointer, `0x1b0`, `0x260`, `0x30`, `0x20` and `1`, and the call goes to `R1179`. The Pause arm of `R0231` makes the same call with `[[L04369]+0x1dc]` as the text (line 119 of the shared string index space, `TEXT-STRTAB-023`). Both panels then go through `R0361` (`MENU-STOP-015`). The class vtable is `L06233`, over the modal panel base `R1180`, which stores the text pointer, a null second string and button case 0.
- Size and position: `R0707` replaces the rectangle with `w = ((cw-8)/0x60)*0x60+8`, `h = ((ch-0x68)/0x40)*0x40+0x68` and centres it (`DLG-PANEL-035`). The constructor rectangle `(left, top, right, bottom)` is `(0x20, 0x30, 0x260, 0x1b0)`, 576x384, which snaps to 488x360. Origins are (76,60) at 640x480, (156,120) at 800x600 and (268,204) at 1024x768.
- Build routine `R1181`: no title or footer string, one OK button of case 0 (control id 4, message `0x445`) whose label is the first line of `dialogs.txt` read through `R0668(0)`, and a body created through vtable slot `+0x88`, which builds `R1182` (a text control, vtable `L06234`) over body rectangle `(40,56,448,272)` in font 1 `[L03615]` with colour pointer `[L06235]`.
- Text control: the base constructor `R1183` shrinks the bottom to whole rows of font height plus 4 (11 rows, bottom 267, height 211); `R1182` then sets the line pitch to font height plus 2 (17 for font 1's 15) and the visible count to height divided by pitch (12). `R1184` adds a scroll bar (width 24, control id `0xdf23`, class `R1185`) and narrows the text width by `0x1a` when the wrapped line count times the line height exceeds the body height, then rewraps at the narrower width. `R0763` paints through `R0764` (`DLG-LINE-038`).
- `text/help.txt` gives a scroll bar on both roots (`TEXT-088`).

**Confidence.** High for the construction call, the shared class, the rectangle snap and the widget sequence (each routine read, operands traced to immediates). Medium for the body height 211 and the 12-line figure: they combine immediates read from `R1183`, `R1182` and `R1184` by arithmetic, and no frame was observed.

**Unknown.** Native pixels. The colour pointer is read in `MENU-079`.

**Amended.** The colour `[L06235]` points to is read in `MENU-079`: ramp `L03668`, grey with 14 per channel per entry, over a shadow from ramp `L03613`. The body, scroll bar and OK rectangles are given there. The construction, the snap and the widget sequence above stand.

### MENU-053

- Stop: the help panel is shown by `R0361`, which ORs `8` into `campaign+0x3dc` (`MENU-STOP-015`), darkens the screen once (`DLG-DIM-013`) and is stored in no `campaign+…` slot, so its close takes `R0709`'s default arm (`DLG-CLOCK-014`): bit 3 is cleared and the stopped time is discarded. The word while help is open is 9, so F1 over help meets `MENU-051`'s compare and is ignored.
- Close: the panel's message handler `R0716` treats `0x445` and `0x446` alike, hiding the panel through vtable slot `+0x84` and posting `0x44c`. The panel's key handler `R0818` turns Esc (`0x1b`) into `0x446`. With help open the frame handler's Esc arm does not open the Esc menu: the word is neither 1 nor 0, so the arm falls through to the forwarder that delivers the key to the root and so to the panel (`MENU-INPUT-016`). The OK button's key handler `R0713` posts its own message on Enter (`0xd`).
- Focus and navigation: `R1181` calls vtable slot `+0x1c(0x10, 1)` on the OK button; `R0717` moves focus by Tab and the four arrow keys over the neighbour links at `+0x48..+0x54`.
- Scroll: the scroll bar's key handler `R0731` acts on Page Up, Page Down, Up and Down (`0x21`, `0x22`, `0x26`, `0x28`) by calling its set-position slot `+0x80` with a one-line or one-page step and then posts `0x46d`; it acts only when its own flag test `vtable+0x20(4)` succeeds and otherwise defers to the base handler. The frame handler forwards those four keys to the root.
- No key pages the text and no markup paginates it in the routines read: the body scrolls through the scroll bar.

**Confidence.** High for the stop mechanism, the close messages and the Esc routing (instruction reads, callee behaviour cited from the named claims). Medium for the scroll key set and its one-line or one-page steps (the handler's four arms are read whole, the step operands from `R0781`, `R0782`, `R0783` and `R0784`).

**Unknown.** Resolved in `MENU-078`: the flag test belongs to the text control, which holds the initial focus; mouse handling and the wheel search are read there.

**Amended.** The scroll key arms belong to the text control, not the scroll bar: `R0731` is reached through the text control's `vt+0x6c` `R1186`; the scroll bar class has its own key routine `R1187`, which acts on Left and Right only (`claims/retracted.md`). The focus rule is resolved in `MENU-078`: the text control holds the initial focus, so the four keys scroll at once, and they do nothing while the OK button holds focus. The scroll bar's mouse handling is read there and there is no wheel route. The stop mechanism, the close messages and the Esc routing stand.

### MENU-054

- The map view's key handler `R0819` first requires `campaign+0x3dc == 1` (`L06236`) and, when `campaign+0xc0` is nonzero (the chat text entry of `AI-KEY-125`), forwards the key to that object and returns. Digits take a separate arm.
- The letter group is dispatched through a byte table at `L06237` and a target table at `L06238` (an indirect jump through that 4-byte-entry table). Its guard: at `L06239` the word at offset 0x3dc must equal 1; at `L06240` the selection count at offset 0x140 must be nonzero (zero skips the group); at `L06241`/`L06242` the word at offset 0x144 is masked with 0x24 and a nonzero result jumps to `L06243`. Key `C` (`0x43`, index 2 from base `0x41`) maps to table entry 1, the block at `L06244`.
- A zero selection count (`JZ` at `L06245`) or a set `0x24` bit (`JNZ` at `L06246`) jumps to `L06243`, which starts a second dispatch: key minus `0x20` indexes the byte table at `L02680` and the target table at `L06247`. For C the index is `0x23`, the byte is 13 and entry 13 is `L06248`, which jumps to the epilogue with `EAX = 0`. Only the block at `L06244` ends with `EAX = 1`. When the text entry is open the handler forwards the key and returns 0 as well.
- That block calls `R0317` (does a spellbook popup exist) and jumps to `L06249` when the result is nonzero, otherwise calls `R0097` with argument 5; it then sets the arm's return value to 1. A second C press while the popup exists is consumed with no further call.
- Bit `0x4` of `view+0x144` is ownership by another player and bit `0x20` the structure flag (`AI-PANEL-061`). The town has no map view and no state word of 1, so the key reaches nothing there.
- The other Cast controls route to the same arm and are cited, not re-derived: the command-panel Cast cell (`MENU-COMBAT-018`, `MENU-COMBAT-019`), the quick-spell keys F5 to F8 (`AI-QUICKINVOKE-279`, `AI-KEY-125`) and the item Cast action through `R0097(0x0a)` (`AI-PANEL-123`).

**Confidence.** High: `R0819`'s compares are read in the listing, the entry for `C` is read from the PE's table bytes (byte table `00 08 01 02 …`, `C` at index 2 yields 1, target table entry 1 is `L06244`; the second dispatch's tables are committed beside it), and `R0317` and `R0097` are read whole.

**Unknown.** What the caller does with the returned 0 (whether the key is passed on and who consumes it). How the selection count and ownership bit behave for selections that mix owners: `AI-SPELLCAP-288` bounds the capability side only.

**Amended.** `MENU-063` answers the Unknown about the caller of the returned 0: the frame handler never reads it and always calls the MFC default, so 0 only lets the root offer the key to its other children and its own key routine. The mixed-owner selection question stays Unknown.

### MENU-055

- `R0097(5)` computes the enable mask with `R0093`: `view+0x144 & 4` returns 0, otherwise `0xef`, plus `0x10` when `view+0x144 & 0x200` (`AI-PANEL-060`). Cast is mode 5, tested as `mask & (1 << 4)`, so it passes only when `0x200` is set. `0x200` is set when any accepted selected object is a `CUnit` whose `+0x20` is `0x17` or `0x18` and whose `+0x18` is nonzero (`AI-SPELLCAP-288`).
- Refused: for a nonempty selection with `view+0x144 & 0x24` clear and the bit clear (no spell-capable object, that is no valid caster), `R0097` does nothing after the mask test. Before the test it only calls `AfxGetThread` and `R0093` and, when the stored mode is already 5 and the request is not, closes the popup (not reached for a request of 5). It posts no message line, plays no sound and leaves `view+0x99c` unchanged. `R0819` still returns 1, so the key is consumed. An empty selection or a set `0x24` bit does not reach this routine (`MENU-054`).
- Armed: with the bit set, `view+0x99c` becomes 5 and the command panel receives `0x40d` with index 4. If no spellbook popup exists, `R0383` runs: it sizes the popup region (`R0388`), posts `0x408` to the right-column container `campaign+0xd4` and calls `R0387`, which records the popup's row count and sets `view+0x74 = 1`. No spell is chosen here: the order built on the next map click reads the stored current spell or item (`AI-SPELLGUARD-289`, `AI-PANEL-123`), and a negative result emits no cast.
- The no-spell case is not a separate refusal: the capability bit tests the selection, not the book.

**Confidence.** High for the mask test, the refusal path having no message or sound call and the armed-mode writes (`R0093` and `R0097` read whole). Medium for the spellbook opening through `0x408`: the post is read but the receiver of `0x408` was not traced.

**Unknown.** Whether anything preselects a spell after C when none was chosen earlier in the session (the initial value of the controller's current spell was not traced), what `0x408` shows in the right column, and whether `R0386`, called inside `R0383`, plays a sound.

**Amended.** `MENU-064` and `MENU-065` answer the Unknown: `0x408` is a refresh posted to the right-column container (three child panels write one field each) and shows no popup; the popup is `campaign+0xec`, and no code in the C path writes its current spell. Whether `R0386` plays a sound stays Unknown.

### MENU-056

- Population searched: the PE resource directory of `rom.exe` (type ids 1, 3, 5, 12 and 14 only: no accelerator table, type 9); the image bytes for `.htm`, `.chm` and `winhlp` (no hit) and for `.hlp` and `winhelp` (one hit each, at file offsets `0x0019d52c` and `0x001cd86c`); and the references to every symbol or string containing `Help` (`EnumRefs sym:Help`).
- The path literal `main\text\help.txt` has one reference, the startup routine `R0326` at `L06250` (`TEXT-086`). The `WinHelpA` import has two call sites: `R1188`, the frame-destroy routine, which calls `WinHelpA(hwnd, 0, 2, 0)` (null file, command 2) and is reached from `R0805` at `L06251`; and `R1189`, a library virtual method held in 27 vtable slots with no direct caller.
- Other strings containing `help` belong to other surfaces: three `*_help_context` literals (`R1190`) and `<Alt-h> This help` (`R0440`). Neither is on the F1 path; their surfaces were not examined.
- Each root's `main.res` holds exactly one entry named `text/help.txt`.

**Confidence.** Medium: absence over the named population only. The library's standard `ID_HELP` command (`0xE145`) has dword hits at file offsets `0x0019c650` and `0x0019c654` and in code immediates; its producers and the virtual `WinHelp` method's callers were not traced, so a help route through an MFC modal-dialog hook is not excluded.

**Unknown.** Whether F1 inside a native modal dialog (the save and load chooser) reaches the library's `ID_HELP` handling, and what the `.hlp` literal at `0x0019d52c` is consumed by.

**Amended.** The save and load chooser is not a Windows common dialog: the image holds no `GetOpenFileNameA` or `GetSaveFileNameA` name, and its panels go through the opener that sets bit 3 of `campaign+0x3dc`, so the frame ignores F1 there (`MENU-080`). Whether a library `ID_HELP` route exists for a native child window of those panels stays Unknown; the `.hlp` literal's consumer is not read.

## Settings shortcut and speed notices

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MENU-057 | In a map session Ctrl+W, F, H, U, L, N and O each change one setting and post its new state's line, slot = base + state value, bases 94, 97, 100, 218, 102, 104, 106; W is retreat, F formation, L flying damage, O smoothing. | High | ● active | [EXP-0446](../experiments/EXP-0446-toggle-notice/) |
| MENU-058 | Numpad plus and minus without Ctrl, in a campaign session on the map screen, step the speed index by one, clamped to 0..8, and post slot 108 plus the clamped index through the duplicate-dropping post; Ctrl variants post nothing. | High | ● active | [EXP-0446](../experiments/EXP-0446-toggle-notice/) |
| MENU-059 | These notices are grey lines of 2000 ms appended below the existing lines of the map message line, the oldest removed past the capacity. Toggles are gated by the map screen word, the text entry and the Ctrl latch; speed also needs phase 2. | High / Medium | ● active | [EXP-0446](../experiments/EXP-0446-toggle-notice/) |

### MENU-057

- Dispatch: `R0819` is the map view's key-down handler (vtable slot at `L06252`). A key outside its first jump table reaches `L06253`. With a nonempty selection and `view+0x144 & 0x24` clear, letters `A..T` go through the letter table at `L06238`, where `F`, `H`, `L`, `N` and `O` reach `L06243`; `W` lies outside that table and also reaches `L06243`. A zero selection or a set `0x24` bit jumps straight to `L06243`. The second dispatch there (key minus `0x20`, byte table `L02680`, target table `L06247`, all read from the PE by `tools/togglenotice`) sends vkeys `0x57`, `0x46`, `0x48`, `0x55`, `0x4c`, `0x4e` and `0x4f` to seven distinct arms. Selection state does not change which arm runs.
- Each arm opens by comparing the global `L00627` with zero (the Ctrl latch, `AI-KEYMOD-059`) and skips everything when it is clear.

  | vkey | arm | state | change | slot | other call |
  |---|---|---|---|---|---|
  | `0x57` W | `L06254` | `view+0xa9c` | +1 mod 3 | 94 + state | `R1191(state)` before the post |
  | `0x46` F | `L06255` | `view+0xa98` | +1 mod 3 | 97 + state | `R0078(state)` before the post |
  | `0x48` H | `L06256` | `view+0xaa0` | 0 or 1 | 100 + state | none |
  | `0x55` U | `L06257` | `[L06258]` | +1 mod 3 | 218 + state | `R1192(state)` before the post |
  | `0x4c` L | `L02679` | `view+0xaa4` | 0 or 1 | 102 + state | none |
  | `0x4e` N | `L06259` | `[L06260]` | new = (old == 0) | 104 + state | `R0532(1)` after the post |
  | `0x4f` O | `L06261` | `[L04367]` | new = (old == 0) | 106 + state | none |

- Index: the post pushes `[[L04369] + 4 * (base + state)]`, a line of the shared line array (`TEXT-STRTAB-023`), as `[array + state*4 + disp]` with `disp` `0x178`, `0x184`, `0x190`, `0x368`, `0x198`, `0x1a0` and `0x1a8`. The state read is the value just stored, so the line names the new state. Three-state settings run 0, 1, 2, 0. The arms hold no table of state names.
- `R1191`, `R0078` and `R1192` build record type `0x46` with subtype 1, 2 or 3 and the state at `+0xe`, with the word at `[view+0x9b4]+4` as player, and hand it to `R0217` on object `L00625`. The post is issued whatever that call does.
- The `keyboard.tsv` rows of `AI-KEY-125` for Ctrl+F, Ctrl+L, Ctrl+O and Ctrl+W name F retreat, W formation, L smoothing and O flying damage. The arms read the other way round. The arms agree with `ANIM-NUM-020`: `view+0xaa4` is flying damage and L toggles it. The H, N and U rows agree, and `MISSION-MSGPOST-058` and `ANIM-NUM-020` already carried the correct map; the cause of the `keyboard.tsv` error is not identified. `TERR-LIGHT-108` reads `[L06260]` as the ShowTimeFlow option, which agrees with N as day/night.

**Confidence.** High: the vkey-to-arm mapping is read from the PE's index and target bytes (`key-arms.tsv`, with a test), and every arm, state field, displacement and call is a named instruction in the complete listing of `R0819` (624 instructions, 0 not disassembled). The setting names come from the text each slot holds, not from the state fields' other readers, which were not traced.

**Unknown.** What the receiver does with record `0x46` subtypes 1..3, and what `R0532(1)` redraws.

### MENU-058

- Arms: the plus key (`0x6b`) at `L06262` and the minus key (`0x6d`) at `L06263` of `R0231`, the frame's key-down handler. With the Ctrl latch set the arm writes `campaign+0x40c` only (plus: 1; minus: 0 with a fresh stamp and cleared counters, `AI-KEY-125`) and posts nothing.
- Without Ctrl, `campaign+0x6bc == 2` (`L06264`, `L06265`) and `campaign+0x3dc == 1` (`L06266`, `L06267`) are both required, else the key ends there. Then `R0559(campaign+0x3f4 ± 1)` runs.
- `R0559` clamps its argument to 0..8 with a signed compare, stores it in `+0x3f4` and sets the rate from the ladder 8, 10, 12, 14, 16, 20, 24, 28, 32 (jump table `L02620`).
- The post reads `+0x3f4` again, so at either end of the range the index is the clamped one. It pushes `[[L04369] + 4 * (108 + index)]` (displacement `0x1b0`), ramp `L03668` and lifetime `0x7d0`, and calls `R1193` on `[campaign+0xd0] + 0xa10`, the map view's list (`SESS-VIEW-028`).
- `R1193` drops a one-piece post equal to the newest line (`MISSION-MSGLINE-057`; compare re-read at `L06268`..`L06269`). A second press at the end of the range, while the same line is still the newest, adds nothing; once the line has expired it is posted again.
- Slots 108..116 are the nine speeds in index order.

**Confidence.** High: both arms, the gates, the clamp and the index are named instructions in complete listings (`r-gates.txt`, `d-fns.txt`).

- Text entry: `R0231` from `R0231` to `L06270` and the arms `L06262`..`L06271` hold no `campaign+0xc0` test. `AI-KEY-125` states that an open text entry consumes map shortcuts; no test for it was found in these bounds.

**Unknown.** Whether an earlier handler in the frame's message map consumes the key while a text entry is open.

### MENU-059

- Surface: the seven toggle arms call `R0588` and the speed arms `R1193` on `view+0xa10`, the list `MISSION-MSGLINE-056` describes: in a campaign session drawn from (8, 8) in font 1 with a 1-pixel shadow, 17 pixels per line. The ramp `L03668` is the grey one and the lifetime `0x7d0` is 2000 ms. The line is not a modal panel and has no rectangle of its own.
- Several notices: a post appends below the existing lines, a post past the capacity (14, 16 or 22 lines at 640, 800 or 1024 wide) removes the oldest, and lines expire oldest first, one per tick message, each lifetime counted from the removal of the line above it (`MISSION-MSGLINE-057`). Two toggles in a row therefore show two lines at once. `R0588` never drops a duplicate; only the speed post does.
- Toggle gate: `campaign+0x3dc == 1` (`L06236`, `L06239`), the text entry `campaign+0xc0` closed (`L06272`; an open one takes the key) and the Ctrl latch. On these arms `R0819` reads no `campaign+0x6bc`, player count, selection, ownership bit or option: the 624 instructions hold no `0x6bc` operand.
- Speed gate: phase `campaign+0x6bc == 2`, word `+0x3dc == 1` and no Ctrl (`MENU-058`).
- The frame arm `L06273`..`L06274` forwards letter keys `A..Z` to `campaign+0xcc` with message `0x100`; where that route ends was not read, so the absence of other gates is bounded to `R0819` and the two speed arms.
- In phase 3 the same list is drawn at (0, 220) in font 2 (`MISSION-MSGLINE-056`); the toggle arms add no phase test.

**Confidence.** High for the surface, the lifetime, the ramp and the gates read from the listing. Medium for the dwell as wall-clock time and for the stacking timing: they inherit `MISSION-MSGLINE-057`'s Medium, since the interval of the `0x401` tick message was not measured.

**Unknown.** The `0x401` interval. Whether a phase-3 or networked session delivers these keys to `R0819` at run time: no original was run.

## F12, Backspace, Alt band and the Cast popup

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MENU-060 | In a map session F12 flips the static dword `L06275`; while it is set `R0379` draws a box and a `%3.1f fps` line. No other instruction references the address. | High / Medium / Unknown | ● active (amended) | [EXP-0450](../experiments/EXP-0450-keyboard-remainder/) |
| MENU-061 | The map view's own Backspace arm requires screen word 1 and closed chat, clears the message line's text, colour and lifetime arrays through SetSize-like calls, and returns 0. | High / Medium | ● active (amended) | [EXP-0450](../experiments/EXP-0450-keyboard-remainder/), [EXP-0476](../experiments/EXP-0476-backspace-descendants/) |
| MENU-062 | In a map session Alt plus a letter B..Y except S broadcasts a type `0x46` record with sub-selector `0x80` and index letter minus `A`; Alt+S is the screenshot. No local effect found in the handler. | High / Unknown | ● active (amended) | [EXP-0450](../experiments/EXP-0450-keyboard-remainder/) |
| MENU-063 | The frame discards the root's key return and always calls its library default; within the GUI, zero continues to later children and the own slot, while a nonzero descendant return can suppress the map shortcut. | High / Unknown | ● active (amended) | [EXP-0450](../experiments/EXP-0450-keyboard-remainder/), [EXP-0476](../experiments/EXP-0476-backspace-descendants/) |
| MENU-064 | The Cast arm's `0x408` goes to the container `campaign+0xd4` and shows no popup; three child panels (Medium that they are all) write one field each. The spellbook popup is `campaign+0xec`, shown by `R0383`. | High / Medium / Unknown | ● active | [EXP-0450](../experiments/EXP-0450-keyboard-remainder/) |
| MENU-065 | In the spell-bar class the current slot `[+0x60]` is written by constructors, the `0x411`/`0x417` arms, the click routine and `-1` stores in `R1194`/`R1195`; `0x417` is pushed only by F5..F8; `R0097(5)` is unread. | High / Medium / Unknown | ● active | [EXP-0450](../experiments/EXP-0450-keyboard-remainder/) |

### MENU-060

- Route: the frame key handler `R0231` sends F12 (`0x7b`, byte `0x12` at `L00671 + 0x73`, entry 18 of the 21-dword table at `L00672`) to `L03620`, which forwards message `0x100` to the root `campaign+0xcc` through `vt+0x48`. The map view's key handler `R0819` reaches its F12 arm only after `campaign+0x3dc == 1` (`L06236`) and with the chat entry `campaign+0xc0` closed (`L06272`; an open entry takes the key, `MENU-054`).
- Toggle: at `L06276` the global `L06275` becomes 1 when it was zero and 0 otherwise, then the arm's return value is set to 1. The cell lies past the raw `.data` bytes (image `.bss`), so it starts 0: the readout is off at load.
- Readers: `EnumRefs refto:L06275` gives 4 hits, 2 owners, 0 orphan: the toggle (`L06277`, `L06278`) and two reads in `R0379`, the map view's `0x402` frame routine of `MISSION-MSGLINE-056` (`L06279`, `L06280`). Among direct references the F12 arm is the only writer (a computed or indirect access is not excluded). The cell is in no save record found in this population.
- Effect (listing `rom-fps-draw.txt`): at `L06279` a set cell runs `R0718` with five arguments (`[view+0x10]-0x78`, `0`, `[view+0x10]-0x1e`, `0x18`, a colour built from the pixel-format shifts at `L06281..L06282`), then `R1196` (a `sprintf`) with the format `%3.1f fps` (`L06283`) and the double at `view+0xd0`, then `R0571` in font1 `[L03615]`. At `L06280` a set cell also calls `R0351` on a local rectangle. A clear cell skips both.
- The double: `view+0xd0` is stored by `FSTP` after `FILD` / `FDIVP` at `L06284`..`L06285`, `view+0xc8` is reduced by `0x3e8` and `view+0xc4` cleared (`L06286`..`L06287`): a recomputation once per 1000 view-tick units; the unit of `view+0xc8` and `view+0xc4` is unread, so wall-clock time is not claimed.
- Rejected: a gameplay switch (no direct reader outside the draw routine), a store with no reader, a grid or overlay (the only draw found is the number).

**Confidence.** High for the route, the toggle, the reader set and the format string (arm and tables read whole, string dumped from its own address). Medium for the box geometry (`[view+0x10]` is taken as the view width, not confirmed) and for the value being a frame rate: the label says so, the numerator and denominator of the divide were not read.

**Unknown.** Where the text lands inside the box, the unit of `view+0xc8` and `view+0xc4`, and the `R0351` rectangle's role. Whether a native run draws the box: no original was run.

**Amended.** The box and the text position, the width, the numerator and the unit are read in `MENU-081`: `[view+0x10]` is the view's right edge, equal to the width with left 0; the box is 90x24 at (width - 120, 0); the text is right-aligned at width - 38 and y 0; the rate is draw calls times 1000.0 over elapsed `timeGetTime` milliseconds, in frames per second. `R0351` presents the box rectangle. The route, the toggle and the reader set stand.

### MENU-061

- Route: the frame handler's byte table sends Backspace (`0x08`, index 0) to `L03620`, the same forward as F12. In `R0819` the arm at `L06288` calls `R1197` on the address 0xa10 past the saved object pointer in the stack frame and then jumps to `L06248`, the exit of `MENU-054` with return value 0. Gate: `campaign+0x3dc == 1` and a closed chat entry.
- Routine: `R1197` pushes `-1`, `0` and calls `R0728` on `this+0x04`, `L06289` on `this+0x18` and `R1198` on `this+0x2c`: the arguments of an MFC `CArray::SetSize(0, -1)` (count 0, default grow) on each array. It writes no other field: the redraw flag `+0x5c`, the tick times `+0x40`, `+0x44` and the capacity `+0x48` are untouched.
- Contents: the three arrays are the message line's text (`+0x04`), colour ramp (`+0x18`) and lifetime (`+0x2c`) of `MISSION-MSGLINE-056`, appended by `R0588` and `R1193` (`MENU-059`). The next draw loops over count 0 and shows no line.
- Other callers (`EnumRefs callto:R1197`: 3 hits, 3 owners, 0 orphan): the message line's destructor `R1199` (`L06290`) and `R1149` (`L06291`, not read).
- Rejected: the chat text entry, the selection set, the order queue and the session record queue: the routine's only operand is `view+0xa10`.

**Confidence.** High for the route, the three operands and the absence of other writes (routine read whole). Medium for the three callees being `SetSize` of the three array classes: their bodies were not read, the argument pair and the layout of `MISSION-MSGLINE-056` carry the identification.

**Unknown.** Further live-session descendant effects outside the mapped routes remain Unknown (MENU-089).

**Amended.** MENU-086 establishes entry order, root-focus clearing and repeat offers. MENU-087 and MENU-088 establish text descendants and transient focus that can consume Backspace or suppress the clear. The clear-arm operands and return stand; the bounded ordinary own-handler result is MENU-082.

### MENU-062

- Route: `R0232` (the frame's `WM_SYSKEYDOWN` handler) tests the Alt context bit `lParam & 0x2000`, then in order: Alt+digit and Alt+numpad digit forward to the root (`AI-KEY-125`), `0x53` calls the screenshot routine `R1200`, and a key in `0x42..0x59` with `campaign+0x3dc` bit 0 set calls `R1201(key - 0x41)` (`L06292`..`L06293`). Alt+A and Alt+Z fall out of the range test. After every path the handler calls the MFC default `L01210`. `EnumRefs callto:R1201`: 1 hit.
- Record: `R1201` fills the static record `L02444`: `+9 = 0x46`, `+5` = the word at `[[view+0x9b4]+4]` (the local player's id), `+7 = 0`, `+0xa = 0x80` (a dword), `+0xe` = the index; then `R0217` on `L00625` sends it. A zero word at `+7` is the broadcast addressee of `SESS-CMD-015`.
- Population: 23 keys (B..Y except S) post the same record shape; the index is the only difference. Receiver and effect: `AI-378`.
- Sender-side effect: none found in the handler, which does not touch the state word, the latches or the message line. `R0217` and the trailing call `R1202` (`L06294`) are unread.

**Confidence.** High for the gate, the record fields and the single call site (handler and builder read whole, tables dumped).

**Unknown.** Whether a single-player session loops the record back to `R0061`: the path from `R0217` to the drain loop was not re-read here (`SESS-CMD-015`).

**Amended.** Alt+Backspace and Alt+F12 are not in the band and reach neither the Backspace nor the F12 arm: both arrive as `WM_SYSKEYDOWN`, which `R0232` leaves to the MFC default (`MENU-082`).

### MENU-063

- Frame: `R0231` forwards the key by calling the virtual slot at offset 0x48 of `campaign+0xcc` with the key and message `0x100` (`L03620`..`L06274`) and never reads the call's return value: the next code at `L06227` passes the object on to `L01210` (the MFC default). It returns popping 0xc argument bytes and sets no return value. Every arm of its table ends at `L06227`. A C key (`0x43`) takes byte `0x14` at `L00671 + 0x3b`, entry 20, `L06273`, which forwards letters `0x41..0x5a`: the letter route `MENU-059` left unread.
- Dispatcher: `R0390` (root `vt+0x48`) for `0x100..0x102` calls the focus child `[this+0x38]` first when set, then the child walk `R0642`, then, when both returned 0, the object's own `vt+0x6c`, `+0x70` or `+0x74` (`L06295`..`L06296`); a nonzero result is returned at once. `R0642` stops on a nonzero child result and, for non-mouse messages, continues to the next child on 0 (`L06297`..`L06298`).
- Map view: `R0819` is the own `vt+0x6c` of the map view `campaign+0xd0`. Its return 0 reaches the root's walk as that child's result, so the root offers the key to its remaining children and then to its own routine before returning 0 to the frame, which discards it. For C this is the fall-through of `MENU-054`: no selection, a set `0x24` bit, a state word other than 1 or an open chat entry.
- Character stub: the mapped root's own `vt+0x74` is `R1203`. Its complete five-byte body zeroes the return value and returns popping 4 argument bytes; it returns 0 without a memory write.
- The frame drops the GUI return. Inside the GUI, the chat text and focused modal edit have concrete Backspace effects (MENU-087, MENU-088); their nonzero return can stop traversal before the map clear.

**Confidence.** High for the frame/default route, the GUI dispatch order and the measured root character stub. Global later-session routing remains Unknown.

**Unknown.** Arbitrary later focus/list contents, pre-dispatch hooks and native translated-character delivery (MENU-089). This amendment does not close every descendant's C-key behavior.

**Amended.** MENU-086 establishes insertion order, an initially null root focus and possible repeat offers. MENU-087 and MENU-088 establish chat/map focus and the Drop Gold root focus. MENU-082's ordinary own-slot results stand; they do not cover all later descendants.

### MENU-064

- Posting: `R0383` loads `campaign+0xec`, sizes it with `R0388` (a `R0375` child-3 test picks the height variant, and the overlay bit `0x2` adds a fixed `0x131,0x1e0,0x186` placement), attaches it with `R0385`, calls the sound routine `R0386`, then calls the virtual slot at offset 0x48 of `campaign+0xd4` with the arguments 0, 0 and 0x408 (`L01646`..`L01647`) and `R0387`. The close routine `R0335` posts the same `0x408` (`L01644`).
- Receivers: `R0390` on the container walks its children on 0. Three child routines act on `0x408` (jump entries of message minus `0x402`, read from the PE): `R1204` sets `[this+0x6c] = -100` (`L06299`); `R1205` sets `[this+0x60] = 1` (`L06300`); `R0350` sets `[this+0x5c] = 1` (`L06301`). `R1206` has no `0x408` entry (`L06302`, its default).
- Popup: `campaign+0xec` is built once by `R0083` (`R0315`, `L06303`); it paints 24 cells from bits of `view+0x148` (`L05259`, `L06304`), labels slots `F%d` (`L06305`) and highlights the cell equal to `[this+0x60]` (`R1059`); `R1207` is its tooltip.
- `0x408` is a general refresh: 29 immediates in 20 routines (`EnumRefs imm:408`, `rom-enum-imm.txt`); the Cast open and close are two of them.
- Rejected: a `0x408` handler that creates or shows the popup (the three handlers write one field each).

**Confidence.** High for the posting chain, the three field writes and the popup's identity as the spell bar (listings read whole; tables read from the PE). Medium for the container's children being exactly these routines: the child slots were taken from the construction in `R0315` read earlier, and are not committed beside this card.

**Unknown.** What `+0x6c`, `+0x60` and `+0x5c` of the three panels mean (cursor, mode or redraw state). Whether `R0386` plays a sound: it tests `[L02703]`, walks an object list at `this+0x10`, and its arguments are `[L02733]`, `0`, `0`, `0xdc`; the buffer chosen was not followed.

### MENU-065

- Writers of `[this+0x60]` in the spell-bar routines `R1208`..`R1195` (`EnumRefs disp:60`, `rom-enum-disp.txt`; a store through a pointer to `campaign+0xec` from another routine is outside this search): both constructors `R1208` and `R0083` store `-1` (`L06306`, `L06307`); `R0096` stores `-1` for message `0x411` (`L06308`) and a slot value for `0x417` (`L06309`); `R1209`, the click routine, stores the hit cell when its bit in `view+0x148` is set (`L06310`); `R1194` and `R1195` store `-1` (`R1194`, `R1195`).
- Senders: `EnumRefs imm:411` finds one push, `R0315` at `L06311` (construction); `imm:417` finds four, the F5..F8 arms of `R0819` (`L06312`, `L06313`, `L06314`, `L06315`; census `rom-enum-imm.txt`). The C path's own routines `R0819` (arm `L06244`), `R0383` and `R0387` contain neither. `R0097(5)` is not re-read in this experiment (`MENU-055` read it).
- `R0387` writes `view+0x68`, clamps `view+0x60` (the view's field, not the popup's) and sets `view+0x74 = 1`.
- So C leaves the popup's current spell as it was: `-1` until a quick-slot key, a click or an earlier assignment sets it (`AI-SPELLGUARD-289`, `AI-PANEL-123`).

**Confidence.** High that the C path's routines read here select nothing (bounded to immediate pushes of `0x411` and `0x417` and to displacement stores to `+0x60` in the class's routines; `R0097` is not re-read). Medium that no computed message id reaches `R0096`.

**Unknown.** Whether closing the popup resets the current spell: `R0335` contains no `0x411` push. The message arms of `R1194` and `R1195`. Whether the item-Cast route of `AI-PANEL-123` assigns a slot before the popup shows.

## Hover help bindings and rectangles

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MENU-066 | Attribute-row hover rectangles are panel-relative: value x 82..102, cost 107..127, refund 132..152, y 54+32i..74+32i for rows i 0..3; one shared pool rectangle (46,181,123,203); the gate dword is 1 only while the owner screen runs. | High / Medium | ● active | [EXP-0454](../experiments/EXP-0454-hover-help/) |
| MENU-067 | In a map-list record `+0x14` and `+0x18` are the map file's first two header dwords, its width and height; the list row prints each minus 16 and the `.alm` loader seeks 0x28 bytes after reading them. | High / Medium | ● active | [EXP-0454](../experiments/EXP-0454-hover-help/) |
| MENU-068 | `dialogs.txt` 23 and 75 reach four Sound Options controls through the vslot 6 setter, `dialogs.txt` 117 reaches a text edit box through its constructor, and `patch.txt` 52..54 reach three Game Options controls as hint and caption. | High / Medium | ● active | [EXP-0454](../experiments/EXP-0454-hover-help/) |
| MENU-069 | Of the 52 inherited-getter widget tables not reached by the traced hint sites, 46 are built with a null hint text and 6 differ; of the 9 of 252 vslot-6 receivers that resolve, none is one of the 52. | High / Unknown | ● active | [EXP-0454](../experiments/EXP-0454-hover-help/) |
| MENU-096 | The 55 constructor sites that read a text table (52 `dialogs.txt` entries at 51 indices, 3 `patch.txt`) sit in 14 builder routines; of 172 hint-forwarding constructor sites 114 pass no hint at construction and 58 pass text. | High / Medium | ● active | [EXP-0489](../experiments/EXP-0489-troll-hints/) |

### MENU-066

The initialiser `R1210` writes the rectangles into the panel (vtable `L06316`) in panel coordinates; the getter `R1211` copies each, offsets it by the panel origin (`R1212`: the panel's `(+8, +0xc)` plus the same pair of every ancestor on the `+0x30` chain, added by `R1213`) and tests it with `PtInRect`, which is half-open: left <= x < right and top <= y < bottom. Tests run per row in the order label, pool, value, cost, refund (TEXT-082): five tests per row over 17 distinct rectangles (4 label, 12 value, cost and refund, 1 pool).

| Rectangle | Panel offset | left | top | right | bottom |
|---|---|---|---|---|---|
| value, row i | `+0xb4 + 0x30 i` | 82 | 54 + 32 i | 102 | 74 + 32 i |
| cost, row i | `+0xc4 + 0x30 i` | 107 | 54 + 32 i | 127 | 74 + 32 i |
| refund, row i | `+0xd4 + 0x30 i` | 132 | 54 + 32 i | 152 | 74 + 32 i |
| label, row i | `+0x64 + 0x10 i` | 16 | 57 + 33 i | 16 + W | 57 + 33 i + H |
| pool | `+0xa4` | 46 | 181 | 123 | 203 |

Rows i are 0..3. `W` is the glyph-advance sum of the row label on the font object `[L06186]` (`R0766`) and `H` the vslot `+0x24` result of that object, both runtime values. The pool rectangle is one record tested on every row, so TEXT-082's second rectangle is the same on all rows; its y range lies below the value, cost and refund rectangles (last bottom 170).

The gate `[[view+0x5c]+0x104]` is a dword of the panel's owner object `A`, which the panel constructor `R1214` receives as its last stack argument from `A`'s build routine `R1215`. Within `A`'s methods three stores write it: 0 in the build routine (`L06317`), 1 at the end of the start routine (`L06318`), 0 in the teardown routine (`L06319`). Readers through `[this+0x5c]` are `R1211`, `R1216` and `L06320`. By these stores the hover answers only between the start and teardown of that screen.

**Confidence.** **High** for the rectangle immediates, strides, order and the three stores (instruction rows reproduced from both roots). **Medium** for the owner role and for "only between start and teardown": a whole-image scan for operands `0x104` (`scan104.txt`, 131 rows, displacement or immediate) was matched to `A`'s methods by owner routine, and stores in routines of other classes were taken to be other objects, not traced.

**Unknown.** `W` and `H` in pixels (font data not read), and so whether a label rectangle (bottom 156 + H for the last row) reaches the pool's top 181. The screen `A` implements, and whether anything besides start and teardown changes the dword. The panel origin's absolute value.

### MENU-067

The loader `R0479` reads 4 bytes into record `+0x14` (`L06321`, `L06322`) and 4 bytes into `+0x18` (`L06323`, `L06324`), then seeks 0x28 (`L06325`). These are the type-0 payload `+0` and `+4` dwords that ALM-META-091 stores as `M+0` and `M+4`. The row text routine pushes `+0x18 - 0x10` and `+0x14 - 0x10` (`L06326`, `L06327`) for the `%dx%d` columns of TEXT-083, so the shown size is the raw size minus 16. The record constructor `L06328` initialises only the three CStrings at `+0`, `+4` and `+8` (allocation 0x1c); the loader fills the numbers. The shown size equals the playable rectangle `(8, 8, W-9, H-9)` of TERR-SIGHT-116, which has width `W - 16`.

**Confidence.** **High** that the two fields are the raw dwords and that the display subtracts 16 (named instructions). **Medium** that the 16 is the 8-cell border on each side: the match with TERR-SIGHT-116 is arithmetic, and no instruction ties the two.

**Unknown.** Whether the first two header dwords are width and height in every shipped map (the loader does not name them).

### MENU-068

- `dialogs.txt` 23 and 75, builder `L06329` (the Sound Options dialog, VIDEO-SFX-053): four setter pairs, each a call of the virtual slot at offset 0x18 with argument 0x17 followed by a call of the same slot with argument 0x4b (`L06330`/`L06331`, `L06332`/`L06333`, `L06334`/`L06335`, `L06336`/`L06337`). The receivers, by the construction before each pair: a control of constructor `L06338` (`L06339`, class `L06340`), two of `R0700` (`L06341`, `L06342`; class `L03644`) and one of `R1185` (`L06343`; class `L06344`). All three classes carry the inherited getter `L06345` and their constructors take a hint at construction (`dialogs.txt` 10, 13, 15 and 19 respectively for the four controls, the first-pushed argument), so the setter replaces it later.
- `dialogs.txt` 117: pushed at `L06346` as the last argument of constructor `L06347` (`L06348`, class `L06349`, a text edit box), in the routine at `L06350`. It is a constructor hint.
- `patch.txt` 52, 53 and 54 (object `L06351`), builder `R1217` (the Game Options dialog, TOWN-OPTIONS-457): each is pushed as the last constructor argument of a `L06338` control (`L06352`/`L06353`, `L06354`/`L06355`, `L06356`/`L06357`) and again to caption setter `L06358` of the same control (`L06359`..`L06360` and the two following ones), which forwards to `this+0x64`. The same dialog also reads `dialogs.txt` 52, 53 and 54 for other controls.

**Confidence.** **High** for the push, constructor and setter addresses. **Medium** that the setter receivers are exactly those four controls: they come from the nearest constructor call before each pair, not a data-flow proof; **Medium** that the last constructor argument of `L06338`, `R0700` and `L06347` is the hint text, which follows the TEXT-085 trace.

**Unknown.** The other 243 vslot-6 receivers (a backward trace resolves 9 of 252; one resolved site is a constructor, `R0028`, with no widget table). What `L06358`'s `this+0x64` object draws.

### MENU-069

The `L06361`..`L06362` population is the 69 widget tables whose `+0x14` getter is `L06345` (TEXT-HOVERSET-049); TEXT-085 reaches 17 by the hint-forwarding sites and the other 52 are classified here from the constructor that stores each table's vtable address, followed through its base-constructor chain to the text parameter of `R0773`, `R1218` or `R1219`.

| Traced text of the 52 | Tables |
|---|---|
| null constant | 46 |
| no constructor in the code map | 2 (`L06363`, `L06364`) |
| forwards a parameter, base of a reached class | 2 (`L06365`, `L06366`) |
| static address plus a parameter | 1 (`L06367`, base of two reached classes) |
| null and static in different constructors | 1 (`L06368`) |

Null text means the base constructor assigns no hint and the hint is empty until a setter writes it; hover raises only a nonempty string (TEXT-HOVER-048). By direct callers of the storing constructors, 44 of the 52 have a builder caller, 5 only a derived constructor and 3 none in the code map. Of the 9 vslot-6 sites whose receiver resolves (of 252), 8 target reached classes (MENU-068) and 1 is a constructor with no widget table (`R0028`); none is one of the 52. The other 243 are unresolved.

**Confidence.** **High** for the table classification under the stated trace and for the 9 resolved receivers. **Unknown** whether any of the 52 shows a hint: a hint can arrive through any of the 243 unresolved vslot-6 receivers or a derived constructor passing text the trace read as null.

**Unknown.** Which of the 52 controls appear in play, and which of the 243 unresolved vslot-6 sites are widget setters. The three `+0x3c` writers outside the base constructors: `L06369` constructs class `L06370` (not one of the 96 tables) and `L06371` class `L06372`; the receiver class of the one in `L06373` is unresolved.

### MENU-096

The sites are those `TEXT-085` traces; the eight setter calls of `MENU-068` are additional. Where a class is named, the builder is its vtable slot `0x78` function; `L06329` is the Sound Options dialog and `R1217` the Game Options dialog (`MENU-068`).

| Builder | Dialog class vtable | Text sites | Sources |
|---|---|---|---|
| `L06329` | `L06374` | 9 | `dialogs.txt` 8, 10, 165, 11, 13, 15, 19, 20, 21 |
| `R1220` | `L06375` | 4 | `dialogs.txt` 25, 26, 158, 27 |
| `R1221` | `L06376` | 5 | `dialogs.txt` 30, 31, 32, 158, 33 |
| `R1222` | `L01485` | 2 | `dialogs.txt` 46, 47 |
| `R1217` | `L06377` | 8 | `dialogs.txt` 51, 54, 52, 56, 78; `patch.txt` 52, 53, 54 |
| `L06378` | not found | 4 | `dialogs.txt` 164, 156, 59, 64 |
| `L06379` | `L06380` | 2 | `dialogs.txt` 69, 71 |
| `R1223` | `L06381` | 1 | `dialogs.txt` 81 |
| `L06382` | not found | 2 | `dialogs.txt` 82, 84 |
| `L06383` | `L06384` | 5 | `dialogs.txt` 90, 91, 93, 126, 127 |
| `L06385` | `L06386` | 7 | `dialogs.txt` 111, 112, 113, 114, 115, 128, 129 |
| `L06387` | not found | 1 | `dialogs.txt` 131 |
| `L06350` | `L06388` | 1 | `dialogs.txt` 117 |
| `L06389` | `L06390` | 4 | `dialogs.txt` 118, 132, 133, 137 |

The other 117 sites hold no text-table hint: 63 pass null, 51 pass a zero cell (50 at `L06391`..`L06392`, one at `L06393`) and 3 pass the executable literals of `R1190` (`TEXT-085`). The 114 null and zero-cell sites pass no hint at construction; a later setter can still write one (`MENU-068`, `MENU-069`). The text shown for a table entry is the installed line at that index in each locale; line lengths and separator counts for both roots are in the evidence, the lines are not.

**Confidence.** High for the builders, indices and counts, reproduced from both roots over one executable. Medium that a site's text is that control's hint: it follows the argument trace of `TEXT-085` and not a data-flow proof, and the last constructor argument of each base constructor is taken as the hint as in `MENU-068`.

**Unknown.** The class tables of the three builders marked not found. Which of these dialogs are reachable in single-player play. The 243 unresolved vslot 6 receivers of `MENU-069`.

## Mission-screen readout widget

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MENU-070 | Widget 8 shows the hovered object when a hover id resolves and no item is held, else the single selected object, and draws its readout only for a `CStructure` (or a class derived from it); other classes go to `R0877`. | High / Medium / Unknown | ● active | [EXP-0455](../experiments/EXP-0455-mission-screen/) |
| MENU-071 | Widget 8's three structure lines are the `building.txt` text line of class id minus one, string 19 (`Health`) and `%d/%d`, in font `[L02677]` at right edge minus 88; the first two use ramp `L03616`, the ratio `L06394`. | High / Medium / Unknown | ● active | [EXP-0455](../experiments/EXP-0455-mission-screen/) |
| MENU-072 | The ratio is the drawable's current health `+0xfc` over `+0x100`: the first is message word `+0x12`, the second a signed word of a class table, replaced by 1000 when zero; neither store reads the object's class record. | High / Medium / Unknown | ● active | [EXP-0455](../experiments/EXP-0455-mission-screen/) |

### MENU-070

- `R0349` (`R0349`..`L01381`, `evidence/disasm-widget8-paint-R0349.txt`) draws only when `[sess+0x3dc] == 1` (`L01386`); the height gate and the resolution limit are `TOWN-093`'s.
- Hover: when the word `view+0x990` is non-zero (`L06395`) and the session dword `+0x3cc` is zero (`L06396`), the routine looks the id up in the object map `view+0x9b8` (`L06397`..`L06398`): bucket `(id >> 4) mod [view+0x9c0]` in the array `[view+0x9bc]`, chain compared on the word at node `+8`, value at node `+0xc` (the layout is `AI-SELECT-065`'s). A found value is taken at `L06399`.
- Selected: without a found hover value the routine tests `view+0x140 == 1` (`L06400`) and takes `view+0x138`. The hovered object therefore wins whenever it resolves; the selection is read only when no hover id resolves or an item is held.
- Class: the object is tested with `R0214` against the runtime-class record `L06401` (name `CStructure`, `evidence/dwords-crtc-gameobject-family-L10352.txt`). A match selects the structure arm; a miss passes the object and the screen rect to `R0877` (`L06402`). The records of `CBridge`, `CVerticalWoodenBridge` and `CHorisontalWoodenBridge` name `CStructure` or `CBridge` as base, so they take the structure arm; `CBackPack`, `CUnit` and `CProjectile` name `CGameObject` and `CAirUnit` names `CUnit`; none of the four derives from `CStructure`, so none takes the structure arm.

**Confidence.** High for the gate, the two object sources, their precedence and the class record (all instructions read). Medium that `IsKindOf` on this record admits exactly the four structure classes: the base chain is read from the five records, not from `R0214`'s body. Population: EN `rom.exe`; the RU image has the same SHA-256 (`evidence/image-hashes.txt`).

**Unknown.** What `R0877` draws and for which classes; it is read only to its call here.

### MENU-071

- The screen rect comes from `R1224` applied to the widget rect at `esi+8` (`L06403`); right and top below are its fields `+8` and `+4` of the output (`[esp+0x28]`, `[esp+0x24]`).
- Each line is one call of `R0571(x, y, text, 2, ramp, 1)` on the font object `[L02677]`: flag 2 and shadow offset 1 (`L06404`, `L06405`, `L06406`).
- Line 1: x = right − 0x58, y = top + 0x1c, text `R0668(L06407, objectClassId − 1)`: `[L04369][[L06407+0xc] + id − 1]`. The object `L06407` is the `building.txt` text object (`evidence/disasm-building-text-load-L13153.txt`); with base 649 (`TEXT-STRTAB-023`) EN gives `Shop` for ids 34 and 35, `Grave` for 39 and 40 and `Church` for 66, RU `Лавка`, `Могила` and `Ратуша`. The six rows are in `evidence/uitext-targets.txt`, read from both installs.
- Line 2: x = right − 0x58, y = top + 0x2c, text `[L04369][0x4c / 4]`, global string 19: `Health` (EN), `Жизнь` (RU), `main.txt` line 19.
- Line 3: x = right − 0x58, y = top + 0x36, `sprintf("%d/%d", (i16)obj+0xfc, (i16)obj+0x100)` with format `L04526`.
- Ramps (`evidence/ramps.txt`, `R0572` run on the emulator over synthetic memory, 8-bit channels at shift 0): line 1 and line 2 use `L03616`, whose entry 15 is (185,159,73); line 3 uses `L06394`, entry 15 (107,154,120). The shadow pass of `R0571` uses the flat ramp `L03613` (`MISSION-MSGLINE-056`).
- The routine contains no owner text and no player-name read, so `TOWN-093`'s "name/owner label" is a name and the word `Health`.

**Confidence.** High for the arguments, offsets, text objects and ramp values. Medium for which ramp entry draws the glyph body: `R0571` was not re-read, the entry-15 reading follows `MISSION-MSGLINE-056`. Medium for the base 649, carried from `TEXT-STRTAB-023` rather than measured here.

**Unknown.** The font object's glyph cell height and whether the three lines overlap at 1024x768; the colours after the display's quantisation.


### MENU-072

- `+0xfc` is written at `L06408` from message word `+0x12` in the client arm `0x82` (`evidence/disasm-client-arm-0x82-hp-store-L13154.txt`); `TERR-STRUCT-102` already reads it as the current health.
- `+0x100` is written at `L06409` from the signed word at byte `+0xc` of the array behind `[L06410] + messageByte+0x11 × 0x1c + 0xc`, index 3, and replaced by `0x3e8` (1000) when that word is zero (`L06411`..`L06412`). Neither value is read from the object class record `[L01425][obj+0x20]`.
- The simulation side of message `0x82` is `UNIT-STRUCTCONT-078`: it sends the object's HP word, or the literal 1000 when its own maximum is zero.
- `Indestructible` classes 15, 16, 34, 35, 39, 40 and 66 are all `Usable` (`EXP-0037`'s `structures.reg` rows). A readout for them shows the message value over the table word or 1000.

**Confidence.** High for the two stores, their sources and the 1000 fallback (instructions read). Medium that table index 3 is the class maximum health: the table identity at `L06410` is not named in this experiment.

**Unknown.** The value the table holds for classes 34, 35, 39, 40 and 66 and so the exact readout (`1000/1000` if both sources fall back); the table was not read from the install here. The RU `structures.reg` rows and the RU copy of that table were not compared; RU coverage of this claim is the image hash (the code) and, for `MENU-071`, the RU text rows.


## Settings dialogs

Both dialogs are modal children of the campaign window. Rectangles are dialog-local pixels at the snapped width `W = 488`, read by hand from the constructor argument pushes of each control in its builder listing; they are not post-layout rectangles. `d<N>` and `p<N>` name local rows of `dialogs.txt` and `patch.txt` (TEXT-099). The options object at `L06413`, its stored values and their persistence are in VIDEO-077; the sound configuration at `L03126` is in MENU-076 and VIDEO-076.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MENU-073 | The Game Options dialog holds 19 controls: 5 labels, a 0..8 speed slider, 8 checkboxes, 3 radio groups, OK and Cancel; each non-label control id imports and exports one stored option. | High / Medium | ● active | [EXP-0457](../experiments/EXP-0457-settings/) |
| MENU-074 | Game Options writes its options only on OK, in a fixed order with Animation 0 forcing Lighting 0, then sends three party commands; Cancel and Esc forward the close with no export; the dialog's own code restores nothing. | High / Medium | ● active | [EXP-0457](../experiments/EXP-0457-settings/) |
| MENU-075 | The Sound Options dialog holds 15 controls (track list with scrollbar, Play, Stop, Random Order, Acknowledgments, three volume sliders, OK, 5 labels) and has no Cancel button. | High / Medium | ● active | [EXP-0457](../experiments/EXP-0457-settings/) |
| MENU-076 | In Sound Options every control but Acknowledgments applies at the click or slider move; Acknowledgments is exported by OK only; OK and Esc both close and neither restores the earlier values. | High / Medium | ● active | [EXP-0457](../experiments/EXP-0457-settings/) |
| MENU-077 | Game Options, Sound Options and the cutscene list share the dialog base: size snaps to `96k+8` by `64k+104`, centred, drawn as the `lm.256` nine-piece frame with an 8 px shadow; title and buttons are builder controls. | High / Medium | ● active | [EXP-0457](../experiments/EXP-0457-settings/) |

### MENU-073

Class vtable `L06377`, constructor `R1225`, builder (vtable `+0x78`) `R1217`, message handler (vtable `+0x48`) `R1226`. The campaign arm `L06414` creates it for message `0x41b`, only when campaign `+0x3dc` equals 1 (a map session). The Esc menu button of id 3 posts that message (`L06415`). The town Esc menu `R0363` posts neither Game Options nor Sound Options (MENU-ESC-010 covers the menus). Constructor arguments are `(1, 40, 0, 600, 480, 0)`, a 560 by 480 rectangle that the base snaps to 488 by 424 at (76,28) on a 640 by 480 screen (MENU-077).

All controls are built with `(id, left, top, right, bottom, ...)`. Text is `d<N>` hint then caption unless noted.

| id | kind | rect | text | bound storage |
|---|---|---|---|---|
| -1 | title label, centred | 40,20,W-40,40 | d150 | none |
| 1 | label | 40,56,232,80 | d50 | none |
| 2 | slider, max 8 | 40,84,232,108 | hint d51 | campaign `+0x3f4` (speed level) |
| 0xc | checkbox | 40,120,232,144 | d54, d55 | `L06260` |
| 3 | checkbox | 40,148,232,172 | d52, d53 | `L04367` |
| 0x1f | checkbox | 40,176,232,200 | p52 | `L06416` |
| 0x20 | checkbox | 40,204,232,228 | p53 | `L06417` |
| 0x21 | checkbox | 40,232,232,256 | p54 | `L05659` |
| 4 | checkbox | 256,56,472,80 | d56, d57 | `[L06418]` |
| 5 | checkbox | 256,84,472,108 | caption d78 | `[L06419]` |
| 0xd | checkbox | 256,112,472,136 | caption d156 | `L03631` |
| 0xf | label | 256,152,448,176 | d159 | none |
| 0xe | 3-item radio | 256,180,448,252 | items d160..d162, hint d164 | `L06258` |
| 6 | label | 40,264,232,288 | d58 | none |
| 7 | 3-item radio | 40,288,232,360 | items d60..d62, hint d59 | `[L06420]` |
| 8 | label | 256,264,424,288 | d63 | none |
| 9 | 3-item radio | 256,288,424,360 | items d65..d67, hint d64 | `[L06421]` |
| 0xa | OK button, message `0x445` | W/7,368,3W/7,392 | d0 | none |
| 0xb | Cancel button, message `0x446` | 4W/7,368,6W/7,392 | d1 | none |

At `W = 488` the buttons span x 69..209 and 278..418 (integer division). The values bound by the pointer cells `[L06420]`, `[L06421]`, `[L06418]` and `[L06419]` live in the party object `[campaign+0xd0]` at `+0xa98`, `+0xa9c`, `+0xaa0` and `+0xaa4` (options constructor `R1227`, `L06422`..`L06423`); the meanings are formation mode, retreat (wimpy) mode, show-all-hit-points and flying damage numbers. Id 0xe is the autohealing mode (TOWN-AUTOHEAL-458). Ids 0x1f, 0x20 and 0x21 are Shadows, Dynamic lighting and Object animations (TOWN-GRAPHICS-459); id 3 is Smoothing (TOWN-SMOOTH-460); id 0xc is the Day/Night change flag; id 0xd is TipsMode (TOWN-186).

A checkbox control (vtable `L06340`) holds its value or bitmask at `+0x80`. Import (vtable `+0x44`, `L06424`) copies the bound dword into `+0x80` when the dialog is built; export (vtable `+0x3c`, `L06425`) writes `+0x80` back. A left click (`L06426`) toggles the bit and posts `0x46d` to the dialog (`L06427`). None of these stores a pointer to the bound dword; the class code is in `rom-control-classes.txt`. The speed slider (vtable `L06344`; import `L06428`, export `L06429`) holds `{position, max}` at `+0xc4`, `+0xc8` and has flag 1 (enabled) cleared when campaign `+0x6bc` is 0 or 1 (`L06430`..`L06431`; SESS-PHASE-002 names phases 2 and 3).

**Confidence.** **High** for the control ids, kinds, slot indices and bound addresses (instructions read; the EN and RU roots run one `rom.exe`). **Medium** for the pixel rectangles: they were read by hand from the constructor argument pushes of the builder in the committed listing, not computed, and a control constructor or the base could adjust a rectangle after construction.

**Unknown.** The meaning of campaign `+0x6bc` values 0 and 1. Whether the Esc menu is the only poster of `0x41b`: the arm's posters were not swept image-wide.

### MENU-074

The handler `R1226` acts only on message `0x445` (OK). Every message, including `0x445` after the exports, is then forwarded to the base handler `R0716`, which for `0x445` and `0x446` closes the dialog with the message as result code (`R1228`).

On OK the handler exports, in this order:

1. Slider id 2: `{position, max}`, then the campaign speed setter `R0559(position)`. It clamps to 0..8, stores campaign `+0x3f4`, sets the frame interval `+0x3f0` to 1000 divided by a rate that rises from 8 at level 0 to 32 at level 8, and re-reads the clock into two timer fields.
2. Id 0xd to `L03631`, id 0xc to `L06260`, id 3 to `L04367`.
3. Ids 0x1f, 0x20, 0x21 to `L06416`, `L06417`, `L05659`. If Animation (`L05659`) is then 0, Lighting (`L06417`) is set to 0 (`L06432`, `L06433`).
4. Ids 4, 5, 7, 9 through the party pointer cells to the party object, then id 0xe to `L06258`.
5. Three party commands through `[campaign+0xd0]`: `R0078(formation % 3)`, `R1191(retreat % 3)` and `R1192(autocasting % 3)`, each a type-0x46 record posted to `R0217` with subtypes 2, 1 and 3 (MENU-057).

The commands are sent on every OK, whether or not the value changed. A checkbox or radio click only changes the control's own copy: the handler does not act on `0x46d`, so no option is written before OK. Cancel (button message `0x446`) and Esc (`R0818` sends `0x446` to vtable `+0x48`) reach the base handler without any export, so every stored option keeps the value it had when the dialog opened and no restore code exists. A speed change takes effect inside the OK arm; the Smoothing, Shadows, Lighting and Animation flags take effect at their consumers (TOWN-GRAPHICS-459, TOWN-SMOOTH-460).

**Confidence.** **High** for the export order, the Animation to Lighting coupling, the clamp and the absence of an export on `0x446` (instructions read; TOWN-OPTIONS-457 bounds the same arm by measurement). **Medium** that nothing outside the dialog restores a value on Cancel: the dialog's own code does not, the `0x44c` receiver `R0709` was not read, and no other writer was swept.

**Unknown.** What `0x44c`, sent by the base close routine after the dialog ends, does with the result code (`R0709` not read; it compares the result against `0x445` in at least one arm). Which key a control uses for Enter: the base key handler `R0717` does not handle it.

### MENU-075

Class vtable `L06374`, constructor `R1229` (the extra field `+0x70` holds the sound configuration `L03126`), builder `L06329`, handler `L03131`. The campaign arm `L06434` creates it for message `0x422` when campaign `+0x3dc` equals 1, or equals 0 with `[L00285]` nonzero. Posting sites of `0x422` are `L06435` (the mission Esc menu button of id 4), `L06436`, `L06437`, `L06438`, `L06439`, `L06440` and `L06441`. Constructor arguments are `(1, 100, 30, 640, 450, 0)`, a 540 by 420 rectangle snapped to 488 by 360 at (76,60) (MENU-077). The dialog keeps the music player (campaign `+0xc8`) at `+0x68`, the selected row at `+0x6c` (initially 0) and the configuration at `+0x70`.

| id | kind | rect | text | bound value |
|---|---|---|---|---|
| 0x22b | title label, centred | 40,20,W-40,45 | d7 | none |
| 0x22d | label (Tracks) | 40,60,W-40,78 | d143 | none |
| 3 | list box | 40,80,W-64,170 | hint d11 | selected row to dialog `+0x6c` |
| 0xa | scrollbar on the list, 24 wide | right of the list | none | none |
| 2 | checkbox (Random Order) | 40,190,252,214 | hint d10, caption d9 | `config+0` |
| 0x28 | checkbox (Acknowledgments) | 40,223,252,247 | caption d165 | `L06442` |
| 4 | Play button, message `0x476` | 40,256,140,280 | d12, d13 | none |
| 5 | Stop button, message `0x477` | 150,256,252,280 | d14, d15 | none |
| 1 | OK button, message `0x475` | 40,290,252,314 | d0, hint d8 | none |
| 0x1a | label (music volume) | 258,175,W-40,190 | d16 | none |
| 6 | slider | 258,190,W-40,214 | hint d19 | `config+8`, range `config+0xc` |
| 0x1b | label (sound effects volume) | 258,224,W-40,239 | d17 | none |
| 7 | slider | 258,240,W-40,264 | hint d20 | `config+0x10`, range `config+0x14` |
| 0x1c | label (speech volume) | 258,275,W-40,290 | d18 | none |
| 8 | slider | 258,290,W-40,314 | hint d21 | `config+0x18`, range `config+0x1c` |

The list shows the titles of the player's current candidate bank (VIDEO-OPTIONS-057, VIDEO-076). Ids 2, 4, 5 and 6 get flag 1 (enabled) cleared with hint d23 when `config+0x20` is 0 (`-nomusic`, VIDEO-077), and hint d75 when the candidate bank is empty. The builder ends by setting player `+0x68` to 1 (`L06443`). There is no Cancel button.

**Confidence.** **High** for ids, kinds, text slots, bound values and the gate rules. **Medium** for the pixel rectangles (constructor arguments, as in MENU-073) and for the scrollbar geometry, which is read from one constructor call.

**Unknown.** The consumer of player `+0x68` (set to 1 at build, 0 only by the OK arm). The poster identity of `L06436`.

### MENU-076

Handler `L03131` indexes messages `0x466..0x477` through the byte table at `L06444`; messages outside that range, including `0x445` and `0x446`, go to the base handler `R0716` (`L06445`).

| message | effect |
|---|---|
| `0x46d`, id 2 | store `config+0` and call the player's shuffle setter `R1230(value)` (VIDEO-MUSIC-056) |
| `0x46d`, id 3 | store the selected row at dialog `+0x6c` only |
| `0x46d`, id 6 | volume `= R0659(position, config+0xc)`; if it differs from the player's gain, store `config+8` and, unless the player state is 2, apply it as gain through `R1231` |
| `0x46d`, id 7 / 8 | volume to `L02733` / `L02739`; id 7 also posts `0x484` to `[L06446]` |
| `0x46d`, ids 4, 5, 0x28 | none |
| `0x473` (slider release), wParam 7 / 8 | play a test sound at `L02733` / `L02739` through `R0386` |
| `0x476` Play | `[L06447] = 1`; if the selected row equals the player's current row, set gain and start; otherwise stop, select the row (`R1232`), set gain and start |
| `0x477` Stop | player state 2: stop (`R1233`); state 1: fade (`L06448(0x7d0, 0x1f40)`); then `[L06447] = 0` |
| `0x475` OK | player `+0x68 = 0`, export Acknowledgments to `L06442`, forward `0x445` (close) |

The gain setter `R1231` clamps to -10000..0. Slider positions map to volumes by `volume = trunc(-max * ((position - max) / max)^2)`; the inverse, used when the dialog is built, is `max - trunc(max * sqrt(-volume / max))` (`L06449`, `R0659`; truncation by `R0279`).

Volumes, Random Order and the row selection are therefore applied at the control event, and Esc, which has no export, leaves them as moved. Acknowledgments is the only control written by OK alone, so Esc discards its change. Esc and OK both close; only OK clears player `+0x68`.

**Confidence.** **High** for the per-message arms, the slider formula and the closing paths (instructions read). **Medium** that the Play and Stop calls reach the music player as VIDEO-OPTIONS-057 states: the receivers are the dialog's player field, not an independent trace.

**Unknown.** Whether Esc leaving player `+0x68` at 1 changes player behaviour. How registry persistence of the volumes is triggered after a slider move (VIDEO-077 names the save routine only).

### MENU-077

Base constructor `R0758` calls the panel constructor `R0759` (vtable `L03623`) and then the snap routine `R0707`: width `W' = ((W - 8) / 96) * 96 + 8`, height `H' = ((H - 104) / 64) * 64 + 104` with truncation toward zero, centred on the screen words `[L00618]` and `[L01259]`. At 640 by 480:

| dialog | argument rectangle | frame size | origin |
|---|---|---|---|
| Game Options | 560 by 480 | 488 by 424 | (76,28) |
| Sound Options | 540 by 420 | 488 by 360 | (76,60) |
| cutscene list (VIDEO-075) | 440 by 420 | 392 by 360 | (124,60) |

The painter `R0760` draws the frame from the nine-piece `interface/lm.256` sprite (DLG-PANEL-035) with an 8 px shadow band; the body is `(W' - 8)` by `(H' - 8)`, tiles are 96 by 64, and the background-bitmap argument is 0 in all three. The frame draws no title and no button: the title label, OK and Cancel are controls added by each builder at the coordinates in MENU-073 and MENU-075. Esc is handled by the base (`R0818`, message `0x446`), so every dialog of the base closes on Esc.

**Confidence.** **High** for the snap arithmetic, the painter and the Esc path; **Medium** for the 640 by 480 sizes, which assume the screen words hold 640 and 480 (other screen modes change the origin but not the snapped size).

**Unknown.** Which screen-mode values `[L00618]` and `[L01259]` take in the 800 by 600 and 1024 by 768 modes.

## Help panel keys, layout and colour, the F1 gate, the fps readout and the key remainder

The help panel is the Pause panel's class (`MENU-052`). Rectangles are panel-local pixels at the snapped 488x360 size; screen positions add the panel origin (76,60) at 640x480, (156,120) at 800x600 and (268,204) at 1024x768. Every fact is a static read of `rom.exe`, one image on both roots.

| ID | Public functional claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MENU-078 | In help, the arrow and page keys scroll the text only while the text control holds focus, which it does on opening; with the OK button focused they do nothing. Enter closes from either focus; no wheel route found. | High / Medium | ● active | [EXP-0465](../experiments/EXP-0465-keyboard-help-remainder/) |
| MENU-079 | Help's text control is 382x211 at (40,56) with 12 lines; its scroll bar is 24x211 at (422,56), positions 0..lines minus 12 (range argument lines minus 11); the OK button is 96x24 at (196,300); text is grey ramp entry 15 with a 1 px shadow. | High / Medium | ● active | [EXP-0465](../experiments/EXP-0465-keyboard-help-remainder/) |
| MENU-080 | The Drop Gold modal and the save and load chooser panels set bit 3 of `campaign+0x3dc` and so block F1; the chat entry and the spellbook popup were found to leave it at 1 (Medium), so F1 should open help over them. | High / Medium | ● active | [EXP-0465](../experiments/EXP-0465-keyboard-help-remainder/) |
| MENU-081 | The F12 readout is a 90x24 box at (width-120, 0) with `%3.1f fps` right-aligned at width-38 in font 1; the value is draw calls x 1000.0 over elapsed `timeGetTime` ms, recomputed above 1000 ms. | High / Medium | ● active (amended) | [EXP-0465](../experiments/EXP-0465-keyboard-help-remainder/), [EXP-0472](../experiments/EXP-0472-fps-readout/) |
| MENU-082 | The mapped ordinary map objects' own Backspace slots add no effect after the clear; Alt+Backspace and Alt+F12 take the system-key default and reach neither own shortcut arm. | High | ● active (amended) | [EXP-0465](../experiments/EXP-0465-keyboard-help-remainder/), [EXP-0476](../experiments/EXP-0476-backspace-descendants/) |
| MENU-083 | The map-view constructor initializes the rate to double 0.0; the draw path primes time on its first call and samples above 1000 ms even when F12 is hidden. | High / Medium | ● active | [EXP-0472](../experiments/EXP-0472-fps-readout/) |
| MENU-084 | F12 uses font1's installed `.16` glyphs and `.dat` advances from `graphics.res`; both complete nodes match between the measured EN/RU roots. | High | ● active | [EXP-0472](../experiments/EXP-0472-fps-readout/) |
| MENU-085 | F12 packs box and shadow from component 8 and ink from a 16-entry 17*i ramp using surface-derived channel shifts; RGB565 and RGB555 give conditional dark words 0x0841 and 0x0421. | High / Unknown | ● active | [EXP-0472](../experiments/EXP-0472-fps-readout/) |

### MENU-078

- Child order and flags: the panel's children are the text control, the scroll bar and the OK button, appended in that order by `R0385`. The text control's base constructor ORs flag `2` (tab stop) into `+0x18`, giving 3; the scroll bar constructor leaves 1; the OK button holds 1 and 2. The panel show routine `R0775` runs `vt+0x24(1)` and `vt+0x28(1)` on itself, then `R0787(1, 0)` (Tab forward), which selects the first child whose flags & 3 is 3: the text control. The OK button takes focus only after a Tab.
- Key route: the panel's own key routine `R0818` turns Esc into `0x446` and passes other keys to `R0717`: Tab moves focus (backward while the shift latch `L00670` is set), and the arrows follow the neighbour links `+0x48` (Up), `+0x4c` (Down), `+0x50` (Left), `+0x54` (Right). `R1234` links the body and the OK button, so Up from the button moves focus to the text control without scrolling; the other arrows do nothing.
- Scroll keys: the text control's own key routine `vt+0x6c` is `R1186`, which calls `R0731`. It acts only when `vt+0x20(4)` (the focus flag) is set; otherwise it defers to the base stub, which returns 0. Page Up (`0x21`), Page Down (`0x22`), Up (`0x26`) and Down (`0x28`) call `R0781`, `R0782`, `R0783` and `R0784`, post `0x46d` to the parent and return 1. End, Home, Left and Right (`0x23`, `0x24`, `0x25`, `0x27`) do nothing. The scroll bar's own key routine `R1187` acts only on Left and Right when its width exceeds its height, so it does nothing on the vertical help bar.
- Steps: the control starts with top line `+0x84` = 0 and current line `+0x88` = -1. The set-position routine `R0730` clamps to 0..(lines - visible) and stores the result in both fields, so after any scroll command they are equal. Up (`R0783`) does nothing while `+0x88` is 0 or below, so the first Up does nothing. Down sets position `+0x88` + 1, so the first Down sets position 0 and moves nothing visible. Page Up (`R0781`) sets top minus visible rows when `+0x88` equals `+0x84` and otherwise sets position `+0x84`, so the first Page Up (-1 against 0) resyncs to 0 and moves nothing. Page Down (`R0782`) sets top plus visible rows minus 1 (11 lines for 12 visible) unless `+0x88` already equals that value, in which case it sets top plus twice visible minus 1; with `+0x88` equal to `+0x84` after every command, the second branch cannot fire while 12 rows are visible. Positions run 0..lines minus 12 (range argument lines minus 11).
- Enter: the OK button's key routine `R0713` acts on Enter (`0xd`) when its parent is set and flag 1 is set; it needs no focus and posts `0x445`, which closes help. The dispatcher `R0390` offers a key to the focus child, then to every child until one returns nonzero, so Enter reaches the button from either focus.
- Mouse: the scroll bar class is `R1185` (vtable `L06344`); its `vt+0x54` `R1235` is the left-button-down handler, `vt+0x4c` `R1236` mouse move and `vt+0x58` `R1237` left-button-up. On a vertical bar (width below height) the press posts `0x469` (line up) within one bar width of the top, `0x46a` (line down) within the bar width minus 4 of the bottom, `0x46b` (page up) above the thumb and `0x46c` (page down) below it; a press on the thumb starts a drag. A drag posts `0x468` with position (range - 1) x (y - top - 0x18) / (height - 3 x (width - 4)), clamped. Left-button-up posts `0x473` and ends the drag. The text control's message routine `R1238` turns `0x468`..`0x46c` from control id `0xdf23` into set-position, line up, line down, page up and page down.
- Wheel: a plain 4-byte search of the file finds seven `0x20a` hits (file offsets in the raw-search listing): the library scroll-view class's message-map entry at `L06450` (a dword followed by three zero dwords; the class map `R1239` is referenced only by a vtable whose constructor has no caller); two instruction-operand coincidences (`L06451`, inside a store of 2 to `+0xa`, and `L06452`, inside a jump displacement); one data-section coincidence (`L06453`); one hit in resource data; and two library pushes of `0x20a` for `SendMessageA` at `L00212` and `L00213` in `R1240`, whose callers and reachability were not traced. The frame procedure `R0701` has no `0x20a` arm and forwards to the root only `0x101`, `0x200`..`0x206` and `0x400` (`MENU-082`); no game message map holds `0x20a`; `R0390` routes only `0x100`..`0x102`, `0x200`..`0x206` and `0x400`. `AI-KEY-125`'s "no wheel route" is narrowed to this population.

**Confidence.** High for the key arms, the focus gate, the Enter route and the scroll bar's message codes (each routine read whole). Medium for the initial focus (the Tab selection and the flag values are read; no run observed the focus), for the step sizes (instruction reads of the four step routines) and for the drag formula (arithmetic read from the listing). The wheel result is Medium: no route found over the game message maps, the frame procedure and the dispatcher's message set, with the two library send sites untraced.

**Unknown.** What a native run draws for the thumb before the first scroll command: the bar's position and range fields are 0 until the first set-position (`MENU-079`). Whether a window message outside the dispatcher's set reaches the scroll bar, and whether `R1240` is reachable from a game window.

### MENU-079

- Geometry: `R0361` adds the panel; `R1181` builds it. The text control rectangle is (40,56)-(422,267), 382x211, after `R1184` narrows the right edge by `0x1a` from 448 and adds the scroll bar `R1185(0xdf23, 422, 56, 446, 267)`, 24x211. The visible count `+0x8c` is height 211 divided by pitch 17: 12. The OK rectangle is computed at `L06454`..`L06455` as left W/2 - 0x30 = 196, top H - 0x3c = 300, right W/2 + 0x30 = 292, bottom H - 0x24 = 324 with W = 488 and H = 360: 96x24. On screen at 640x480 the OK button is (272,360)-(368,384) and the scroll bar (498,116)-(522,327); at 800x600 the OK button is (352,420)-(448,444); at 1024x768 it is (464,504)-(560,528).
- Scroll range and position: `R0730` calls the scroll bar's `R1241(position, lines - visible + 1)`, the only routine reached from the text control that sets the bar's range or position. Positions run 0..lines - visible. With the wrapped line counts of `TEXT-088` that is 0..23 for EN (35 lines) and 0..36 for RU (48 lines). The text control starts at top line 0. The scroll bar constructor zeroes its position `+0xc4` and range `+0xc8`; they change at the first set-position.
- Ink: the control is created with the colour pointer `[L06235]`. The cell's initial `.data` value is `L03668`; a plain 4-byte search finds 75 hits in code, every one a load form (`8b /r` or `a1`), and none is a store form (`89`, `a3`, `c7 05`); indirect writes are not excluded. `R0572`, called once by `R1242`, fills 16-entry ramps; `L03668` is grey with a step of 14 per channel, so entry 15 is 210, and `L03614` is white with a step of 17. The glyph body takes ramp entry 15 (`TEXT-079`, `MISSION-MSGLINE-056`), narrowed to the pixel format by the shifts at `L06281`..`L06282`.
- Shadow: `R0571` draws the shadow first at (x + 1, y + 1) with the font's own ramp, `L03613`, which holds 8 per channel in every entry, then the ink at (x, y). Justified lines (`R0765`) call it per word; other lines call it from `R0764`.

**Confidence.** High for the rectangles (constructor immediates and the snap arithmetic, `MENU-052`), the range formula and the shadow offset and ramp. Medium for the displayed colour: the ramp selection of entry 15 comes from `TEXT-079`, and the narrowed pixel value depends on the format words.

**Unknown.** Whether an indirect write replaces `[L06235]` after load. The pixel value a native run draws.

### MENU-080

- Opener: the inventory-grid routine `R1243` ORs 8 into the word and stores it (`L06456`..`L06457`, the store to offset 0x3dc) before adding the Drop Gold modal `campaign+0x10c` (class vtable `L01485`) to the root and showing it. Its close arm in `R0709` is a compare of the closing object with `campaign+0x10c` at `L06458`, with the bit-3 clear (AND with 0xf7) at `L03313`. The word is 9 while the modal is open, so the F1 case's compare (`MENU-051`) fails.
- Chat entry: `R1244`, called once from `L06459` in `R0819`, attaches `campaign+0x124` under the map view, shows it, and writes `campaign+0xc0` = 1. It stores nothing to `+0x3dc`, and the entry's close arm at `L06460`..`L06461` stores nothing to it. The word stays 1, so F1 opens help over the open entry.
- Spellbook popup: `R0383` (the object `campaign+0xec`) reads the word (a test of bit 1 at `L06462`, choosing the attach parent) and writes none; no store to the displacement `0x3dc` lies in `L06463`..`L06462` or `L06464`..`L06465`, and `R0709` has no compare on `+0xec`. The word stays 1, so F1 opens help over the popup.
- Chooser: the frame key handler posts `0x41a` for F2 and, for F3 while `+0x6bc` is 2, `0x418` (`L06466`, `L06467`). The frame cases for `0x41a` and `0x418` construct the panels `R1245` (the SAVE dialog constructor of `SAV-SAVELABEL-1016`) and `R1246`, over the modal panel base constructor `R0758`, and add them through `R0361`, which ORs 8 into the word. F1 in a chooser therefore reaches the frame case `0x434`, which ignores it.
- Not a common dialog: the image holds no `GetOpenFileNameA` or `GetSaveFileNameA` name and imports one comdlg32 function, `GetFileTitleA`. Census of the word: `EnumRefs disp:3dc`, about 56 store sites in about 33 routines; the store forms of the openers and closers named above are read.

**Confidence.** High for the Drop Gold bit and the chooser panels' use of the common opener (instructions named). Medium for the chat entry and the spellbook popup leaving the word unchanged: absence over the named routines and the displacement census, with indirect and register-indexed stores and bulk copies not excluded.

**Unknown.** Whether the chooser panels embed a native child window that takes F1 before the frame; the library `ID_HELP` entry (`0xE145`, `R1247`) has no path found from a game key. The main menu's own chooser is not read.

### MENU-081

- Geometry: `R0391` builds the map view through the base constructor with rectangle (0, 0, `[L00618]` - 0xa0, `[L01259]`) (`L06468`); its stored rectangle at `view+8`..`view+0x14` is (left, top, right, bottom). `[view+0x10]` is the right edge, so with left 0 it is the width, 480 at 640x480. The root `campaign+0xcc` is built with rectangle (0, 0, `[L00618]`, `[L01259]`) (`L06469`..`L06470`), so its left and top are 0 and screen and view coordinates agree.
- Box: `R0379` calls `R0718(right - 0x78, 0, right - 0x1e, 0x18, colour)`, a clipped solid fill of 16-bit pixels with arguments (x0, y0, x1, y1, colour): 90x24 from (width - 120, 0) to (width - 30, 24). The colour is the value 8 narrowed per channel by the format shifts at `L06281`..`L06282` and `L06471`.
- Text: `R1196` formats `%3.1f fps` with the double at `view+0xd0`; `R0571` draws it in font 1 `[L03615]` with the white ramp `L03614`, flag 1 and shadow offset 1, at x = right - 0x26 and y = top. Flag 1 of `R0767` right-aligns, so the text's right edge is at width - 38, 8 pixels inside the box's right edge, and its top is at y = 0 inside the 24-pixel box (font 1 is 15 pixels tall, `MENU-052`).
- Rate: after the screen draw `R0379` adds 1 to `view+0xc4` (`L06472`..`L06473`) on each call with the state word 1, before the F12 test. It reads `timeGetTime` (import `L00849`) and adds the difference from the previous reading `view+0xcc` to `view+0xc8` (`L06474`..`L06475`). When `view+0xc8` exceeds `0x3e8` (unsigned, `L06476`) it stores `view+0xc4` x 1000.0 (`L06477`) divided by `view+0xc8` at `view+0xd0`, reduces `view+0xc8` by 1000 and clears `view+0xc4`. The numerator is draw calls, the denominator elapsed milliseconds, the unit frames per second.
- Present: a set flag also calls `R0351` on a local rectangle built from the constants (360, 0, 450, 24) at `L06478`..`L06479`; they equal the box only at width 480. `R0351` clips to the screen words and drives DirectDraw surface calls (it tests a surface-lost result); it presents the rectangle.

**Confidence.** High for the arithmetic, the units, the box rectangle and the text position (immediates and the right-align flag read). Medium that `R0351` presents the rectangle (the first half of its body read) and for 480 as the width, which assumes `[L00618]` holds 640.

**Unknown.** The native active pixel format and displayed pixels. What is presented at widths other than 480, where the constant present rectangle no longer equals the box.

**Amended.** `MENU-083` reads the explicit constructor stores of double 0.0 and the first-call timer priming. Sampling precedes the F12 test; enabling it need not show the constructor value. `MENU-084` resolves font 1 to the installed `.16`/`.dat` pair. `MENU-085` reads mask-derived colour packing and conditional RGB565/RGB555 words; no native pixel was observed.

### MENU-082

- Clear caller: `R1149` (called once, from `R0099` at `L06480`) detaches the container's children, re-adds `campaign+0xd8`, `+0xdc`, `+0xe0` and `+0xe4` to the container `campaign+0xd4`, calls `R1197` on `view+0xa10` (`L06291`), adds the map view `campaign+0xd0` and the container to the root `campaign+0xcc`, and sets the word to 1 (`L06481`). The other caller is the message line's destructor `R1199` (`MENU-061`).
- Root children: the map view's own key routine `R0819` ends the Backspace arm at `L06248` with `EAX` = 0. The container's vtable `L06482` and its four children (`L01176`, `L01177`, `L00722`, `L06483`) each override `vt+0x48` with a routine that calls `R0390` first and then tests only message numbers from `0x402` up (and `0x200` in one), none of `0x100`..`0x102`. Their `vt+0x6c` is the stub `R1248` (returns 0), except `campaign+0xe0` (`R1249`), which acts only on Tab (`0x9`). The root's own `vt+0x6c` is the same stub. These mapped ordinary own slots add no Backspace effect after the clear. Population read: the root, the map view, the container and its four children; the descendants of those children were not read.
- Alt keys: the frame message map (read whole, `L06484`.., entries of six dwords) sends `WM_KEYDOWN` (`0x100`) to `R0231` and `WM_SYSKEYDOWN` (`0x104`) to `R0232`; it holds no entry for `WM_SYSCHAR` (`0x106`). The frame procedure `R0701` (read whole) forwards to the root only `0x101`, `0x200`..`0x206` and `0x400` (`L06485`..`L06486`); `0x100` reaches its tail unforwarded, because `R0231` forwards the key itself, and `0x104` and `0x105` take its default arm. `R0232` tests the Alt context bit, then handles `0x31`..`0x39`, `0x60`..`0x69`, `0x53` and `0x42`..`0x59`; Backspace (`0x08`) and F12 (`0x7b`) match none and reach only the MFC default `L01210`. The Backspace and F12 arms (`MENU-060`, `MENU-061`) sit behind `WM_KEYDOWN`, so Alt+Backspace and Alt+F12 reach neither. Population: the frame message map, `R0701`, `R0232` and the arms of `R0231`.

**Confidence.** High for the named ordinary own slots, clear callers and Alt-key population. This is an own-handler bound, not a global descendant result.

**Unknown.** Further later-session descendant effects, retained map-view lists, native hooks and the library system-key default (MENU-089).

**Amended.** MENU-086 identifies root focus immediately after map entry; MENU-087 and MENU-088 identify reachable chat and modal focus plus concrete descendant Backspace effects. The ordinary own-slot and Alt-band facts stand.

### MENU-083

The map-view constructor `R0391` explicitly clears the draw count
`view+0xc4`, elapsed milliseconds `+0xc8`, previous timestamp `+0xcc` and both
dwords of the double at `+0xd0/+0xd4`. The double therefore starts at positive
0.0, independent of allocation contents. `R0315` calls this constructor
and stores its return as `campaign+0xd0`.

The bounded sampling block of `R0379` increments the draw count, reads
`timeGetTime`, and replaces a zero previous timestamp with that reading before
subtracting it. The first call contributes zero elapsed time and one draw.
An unsigned `elapsed <= 1000` skips sampling. Above 1000 it stores
`draw_count * 1000.0 / elapsed_ms`, subtracts 1000 once and clears the count.
This block precedes the F12 flag test. The F12 toggle arm changes only the
flag, so it does not start or reset this sampling block.

Before the first sample, if no other path has changed the constructor state,
`%3.1f fps` formats the stored zero as `0.0 fps`. A readout enabled after
sampling can begin with the already sampled value.

**Confidence.** High for the explicit initialization and the named sampling
and toggle blocks. Raw stores, unsigned branch and block order discriminate
uninitialized memory, elapsed time since system zero, a `>=1000` threshold
and an F12-driven sample reset. Medium for that zero surviving an entire
startup: the constructor, its two own helpers, its caller slice and the map
entry body were read, not every called or indirect mutation path.

**Unknown.** The first visible rate in a native run and complete startup
mutation coverage. No native frame was observed.

### MENU-084

The readout's font pointer is loaded from `L03615`. Its constructor call
passes `graphics\font1\font1` and letter spacing 2 to the `.16` font class,
whose loader appends `.16` and `.dat` and stores them at font `+4` and `+8`.
The font draw slot selects `R0767`; its glyph-ramp slot selects
`R1250` and `R1251`.

Under the archive identity rule `RES-IDENT-034`, the installed resources are
`graphics.res::font1/font1.16` and `font1/font1.dat`. On both measured EN/RU
roots the glyph node contains 32932 bytes, 224 indexed glyphs and a 16x15
frame 0; the advance node contains 896 bytes, 224 dwords. Complete node hashes
are identical across roots, discriminating a per-edition font1 payload.
The proportional advance remains that of `SPR16A-FONT-018`. The ordinary
frame-table limit is retained in partially retracted `SPR16A-FONT-014`;
it establishes no global tail-unreachability claim.

**Confidence.** High for static font selection and installed metadata. The
constructor argument, store and draw load identify font1 independently of
its appearance; resource names and hashes corroborate that result.

**Unknown.** Native CWD/resource selection and indirect replacement of the
font pointer. Neither input root contains `update.lst`; the override
mechanism of `RES-MASK-035` is outside the measured resource selection.

### MENU-085

The F12 box packs the component value 8 per channel at draw time.
`R0572` constructs all 16 shadow-ramp entries with the same component
8 and ink-ramp entry `i` with component `17*i`, `i=0..15`.
`R0571` draws shadow first at (x+1,y+1), then ink at (x,y).
The glyph literal writer indexes the supplied ramp with four-bit samples;
full white is entry 15 rather than every literal pixel's value.

`R1242` obtains a 32-byte format structure from the surface, narrows
the masks at offsets `+0x10/+0x14/+0x18` to words, and derives each channel's
shift from its lowest set bit and span width from highest minus lowest plus
one. The channel identities are established by `TERR-LIGHT-124`.
For nonzero contiguous masks with widths at most 8, component `c` contributes
`(c >> (8-width)) << shift`; contributions are ORed and stored as 16-bit
words. The shifts truncate. Format setup rebuilds the ramps.

| Supplied R/G/B masks | Box and every shadow entry | Ink entry 15 |
|---|---|---|
| RGB565 `f800/07e0/001f` | `0x0841` | `0xffff` |
| RGB555 `7c00/03e0/001f` | `0x0421` | `0x7fff` |

**Confidence.** High for the static formula, ramp selection and conditional
words. Instruction-level mask queries, bit scans, shifts and ramp-index
loads distinguish a fixed packed colour, rounding and solid-white glyphs.
The examples are arithmetic derived from the read blocks with supplied
masks; no original instruction was executed for these measurements.

**Unknown.** The native returned masks and active layout, output pixels and
display conversion. Sparse, zero or wider channel masks are outside the
numeric examples. Packed words do not establish displayed brightness.

## Backspace descendants and focus

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MENU-086 | Ordinary GUI keys go to focus, then children in insertion order, then the own slot; a zero focus return permits re-offer through the list, whose key walk has no visibility/enabled filter. | High | ✔ promoted (branch candidate) | [EXP-0476](../experiments/EXP-0476-backspace-descendants/) |
| MENU-087 | The opened chat's text child consumes Backspace with return 1, requests the current byte-string prefix of length minus one, trims a leading classified run and restores the last stored row when the result is empty. | High / Medium | ✔ promoted (branch candidate) | [EXP-0476](../experiments/EXP-0476-backspace-descendants/) |
| MENU-088 | The measured inventory-open route appends the grid to the map view; its gold modal takes root focus, and its edit's Backspace deletes before a positive caret or returns 0 at caret zero while still notifying its parent. | High | ✔ promoted (amended, partially retracted, branch candidate) | [EXP-0476](../experiments/EXP-0476-backspace-descendants/), [EXP-0477](../experiments/EXP-0477-inventory-gold/) |
| MENU-089 | The measured common-GUI signature has 96 tables and 42 keydown targets; global map-session Backspace effects remain Unknown beyond the mapped construction, handler and focus routes. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0476](../experiments/EXP-0476-backspace-descendants/) |

### MENU-086

`R0390` for `0x100..0x102` calls `[this+0x38]` through `vt+0x48`,
then `R0642`, then the own keydown/up/char slots `+0x6c/+0x70/+0x74`.
It returns the first nonzero result. The child walk increments index from
zero; it does not exclude focus from the list or test visibility/enabled
flags for ordinary keys. Its capture/point gates apply to mouse messages.
Individual handlers can have their own flag gates.

`R1149` detaches the root and right-column lists, clears their
focus/saved-focus through `R1252`, then appends map view and column
to root and minimap, command, character, filler to column in that order.
Root focus is null immediately after this routine. `R0385` sets
the appended child's parent and does not set focus.

The common focus setter `R0777` toggles flag `4`, saves the immediate
parent's `+38` at `+44`, and stores itself at parent `+38`. Clearing restores
the saved pointer when it is still the current focus. It does not propagate
focus to ancestors. `R0787` tests child flags `1|2` to select focus;
these tests do not filter the key walk. Generic dialog open
`R0775` focuses the dialog and its first eligible child.

**Confidence.** High within these whole-body raw dispatcher, append,
clear and focus routines. The distinct branch predicates and focus stores
exclude reverse traversal, automatic ancestor focus and focus-child exclusion.

**Unknown.** A later arbitrary session's focus/list contents, retained
map-view children on re-entry, native hooks and translated character delivery.

### MENU-087

`R1244` appends campaign `+124` to the map view and opens chat.
`R1253` constructs and appends exactly one text child. Its flags
allow generic focus selection. `R1254` sets campaign `+c0` and
opens through `R0775`: map-view focus becomes chat, chat focus its
text child. The setter does not change root focus on this route.

The raw key switch of `R1255` maps `8` to `L06487`, calls
`R1256`, and returns `1`, including an already empty entry. The
routine replaces current string `+84` with its prefix of length minus one;
`R1257` clamps negative length to zero. It calls `R0799`,
which advances past leading bytes accepted by `R0800` and compacts
the retained suffix. If the current result is empty and row count `+78`
is positive, it restores the last array string at `+70`, removes that
array entry and repeats the trim. It does not submit chat text through
the Enter arm.

A child return of `1` stops GUI traversal before the map's own keydown.
The map's chat-open branch also forwards to the chat panel and returns
before the clear, if it is reached. The text char slot `R1258`
returns `0` for char `8`; no second deletion exists in that char handler.

**Confidence.** High for the constructors, pointer stores, native switch,
prefix/restore operations, returns and clear suppression. Medium for the
classifier's bit-8 locale path being whitespace. Therefore a universal
one-displayed-character deletion is not claimed.

**Unknown.** Native display, locale/classification state, and delivery of a
translated `WM_CHAR(8)` beyond the read GUI handler.

### MENU-088

The map-view dispatcher `R0333` calls `R0382` with its own
object. At `L06488` the latter appends inventory grid campaign `+e8` to
that map view. This measured route refutes the character-panel parent
clause of SESS-INPUT-037 and the nested-owner clause of MENU-COMBAT-017.
The grid's own keydown is the zero stub. Its dispatcher `R1259`
adds a `0x401` arm and uses the base keyboard route for `0x100`.
The popup campaign `+ec` also has a zero own keydown; `R0383`
attaches it to campaign `+f0` with screen bit `2`, or the map view otherwise.
These own-slot facts do not close arbitrary descendant mutations.

The gold-sentinel positive-purse branch of `R1243` sets screen bit
`8`, appends Drop Gold campaign `+10c` to root, and invokes its generic open.
It becomes root focus; first eligible child is edit `0x989685`, whose
constructor sets flags `1|2`. The full `R1222` builder appends
exactly edit, action button, cancel button and two labels in that order.

Generic edit keydown `R1260` requires focus flag `4`, a parent and
its own flag `1`; failure returns `0` before updating state. Admission
writes `timeGetTime` to `+78`. Its raw key-8 table target `L06489`
decrements positive caret `+70`, replaces text with `Left(newCaret) +
Right(length-newCaret-1)`,
redraws and returns `1`. It does not call the selection-range deletion
helper of the Delete arm. At caret zero it leaves text/caret unchanged
and returns `0`. Both admitted outcomes dispatch parent message `0x46d`
with edit id and zero. A zero return permits a later offer through the
child list. The Drop Gold own slot falls through to zero for Backspace;
its buttons answer Enter only and its labels use the zero stub.

The modal's screen bit independently defeats the map's exact word-equals-1
clear gate. Backspace does not select the modal's `0x445` action branch.
If generic edit char `8` is independently delivered and admitted, it
inserts nothing but updates time, redraws, notifies `0x46d` and returns `1`.

**Confidence.** High for the measured parent, constructor population,
focus stores, raw switch, deletion operands, gates and returns. The raw
string routines distinguish deletion before caret from selection deletion.

**Unknown.** Runtime offer multiplicity, native translated-character
delivery, and attachment/focus routes outside the measured ones.

**Amended.** MENU-091 (EXP-0477) partially retracts the explicit selection-reset
clause. The whole Backspace arm contains no `+68/+6c` store, and its timer
helper writes only `+78/+74`. The other bounded facts and grades stand.

### MENU-089

The identical 1977344-byte lawful EN/RU images have SHA256
`942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`.
The owned Ghidra program digest matches them. Dry repair owns all 6204
distinct `.rdata` code-pointer targets, with zero remaining repairs.

A raw aligned scan of actual `.rdata`, `L06490..L06491`, selects
signature `+8=R0530`, `+c=R0531`, `+10=L06492`: 96 tables,
42 distinct own keydown targets, all 96 focus slots `+28=R0777`.
The bounded map/chat/Drop Gold study enumerates 15 complete selected
table surfaces and both chat's one-child and Drop Gold's five-child
builders. Other signature classes' relation to a live map is not established.

`disp:38` has 889 hits/367 owners/one orphan; `imm:38` has
1039 hits/312 owners/no orphan. Width-overlap and direct/table caller
sweeps are retained. The one orphan at `L06493` is in game-state context;
this does not classify the hundreds of other owners as GUI focus routes.
The full-file pointer scan finds the same 96 focus-setter table addresses
and no stored address for append, focus selection or focus cycling.
Computed/register pointers, bulk copies, native child windows and hooks
are outside what those searches exclude.

**Confidence.** High for this explicit signature/table and named-builder
population. Unknown for global live-session closure. A class with an
overridden signature, an arbitrary later mutation or an unclosed reachable
focus route can fit the same bytes.

**Unknown.** Further Backspace actions, including a zero-return effect
after the clear, across all possible map-session descendants. The frame can
append child id `0x10` through `R1261`; its full focus/lifetime route
is not closed. The alternate `+ec` parent and retained map-view list also
prevent a universal no-action result. Native hooks/message-pump behavior
and actual display remain unwitnessed.

## Inventory gold input and transaction

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MENU-090 | The measured gold double-click opens on a positive client purse; every open sets edit text to 0, while the later selection resolver separately tests selection and current-actor ownership. | High | ✔ promoted (branch candidate) | [EXP-0477](../experiments/EXP-0477-inventory-gold/) |
| MENU-091 | The measured generic edit handles caret and selection keys; its Backspace arm changes text and caret but contains no selection-endpoint store, and its admitted character route has no digit-only filter. | High | ✔ promoted (branch candidate) | [EXP-0477](../experiments/EXP-0477-inventory-gold/) |
| MENU-092 | Drop Gold action formats the edit text, parses through %f and converts the float to an integer without checking parse success; malformed, locale-dependent and out-of-range results remain Unknown. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0477](../experiments/EXP-0477-inventory-gold/) |
| MENU-093 | On the measured successful gold-selection path, modal action changes client object/purse state and queues 0x23; cancel bypasses that transaction, and a focused cancel button consumes Enter as cancel. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0477](../experiments/EXP-0477-inventory-gold/) |
| MENU-094 | Server 0x23 independently requires a Player, first actor and positive affordable amount, debits Player+38 before Sack creation/merge, and ignores that wrapper's zero return; pickup credits Sack+3c to the looter's Player. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0477](../experiments/EXP-0477-inventory-gold/) |
| MENU-095 | The named Player and Sack serializers persist current server gold at +38 and +3c; LOAD restores the same scalar sources later used by gold and pickup consumers, but native round-trip and queue-time SAVE results remain Unknown. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0477](../experiments/EXP-0477-inventory-gold/) |

### MENU-090

Inventory grid vtable `L06494+5c` selects `R1243`. It requires live
content, hit index other than `-1`, sentinel `word[item+6]==0xffff` and positive
client Player `+0c`. It sets screen bit `8`, attaches campaign `+10c` to root,
stores hit index in modal `+68`, opens it and clears campaign `+41c`.

`R1262` sets text to literal `0` on each open. Its edit setter
`R1263` replaces the CString at `+5c`; it contains no caret/selection
reset. The fresh constructor initializes `+68/+6c/+70=0`. MENU-086 and
MENU-088 supply the immediate-parent focus route. Selection/ownership gates
occur later in `R1264`, not in this opening body.

**Confidence.** High for the whole opening, setter and constructor bodies and
their raw slots. Separate predicates distinguish opening from final admission.

**Unknown.** Native reopening display and current caret/selection/focus state.

### MENU-091

The raw `R1260` table binds Left/Right/Home/End/Backspace/Delete to
`L06495/L06496/L06497/L06498/L06489/L06499`. Admission requires
focus flag `4`, a parent and flag `1`. Left/Right move one byte within the
tested bounds; Home/End move to zero/length. The Shift-latch branches adjust
the selection endpoints. Without Shift those movement arms copy selection
start `+68` to end `+6c`; caret `+70` is a separate store.

The whole Backspace arm tests `+70!=0`, decrements it and replaces the string
with `Left(caret)+Right(length-caret-1)`; it contains no `+68/+6c` store.
The called `R1265` writes only time `+78` and flag `+74`. Delete first
calls selection deletion when `end-start>0`; otherwise it removes at the
caret. The character route `R1266 -> R1267` admits codes at
least `0x20`, replaces a positive selection and inserts a converted byte.
Neither complete body tests for digits. These are local handler statements,
not whole-image absence claims.

**Confidence.** High for the raw switch, whole local bodies and distinct
Backspace/Delete operations. The complete Backspace branch and timer helper
exclude their formerly stated explicit selection reset.

**Unknown.** Native translated-character delivery, byte/display character
correspondence, reopened state and effects outside these named bodies.

### MENU-092

The action's edit getter `R1268` obtains the CString pointer and calls
the formatter with that text as its format argument. `R1269` next
calls the scan wrapper `R1270` with literal `%f`, then loads its local
float and calls `_ftol` at `R0279`. It neither initializes that local float
before the scan nor tests the scan return before converting it.

For ordinary successfully parsed finite text whose integer conversion is in
range, the converted request reaches the grid resolver. The getter's format
operation means percent text is not a demonstrated literal copy. No clean
empty/non-number/percent/NaN/infinity/overflow rejection is established.

**Confidence.** High for argument order, literal identity, raw calls and the
missing local initialization/return branch. Unknown for failed and
environment-dependent conversion results; no native parser corpus was run.

### MENU-093

`R1264` requires no campaign `+3cc` selection, an index below count
and equality between current actor owner pointer and active client Player.
It calls `R1271` through grid `+84`, then writes selected object,
index and source at campaign `+3cc/+3d0/+3d4` and refreshes the grid.
The refresh aliases grid `+84/+90` to actor
`+c8/+dc`, clears its cell-pointer array and clamps the actor scroll field
through `R1272`; it contains no source-quantity store.

For a matched gold sentinel with positive source quantity and no arithmetic
overflow, a request below the source quantity leaves the remainder and
creates a copied object with the request. A request at least the quantity
copies the old quantity and retains the sentinel with quantity zero. Zero
and negative requests have no rejection branch here. Fourteen distinct
small-integer substitutions are derived controls for these two branches.

The successful `0x445` body debits client Player `+0c` by selected `+10`,
queues `0x23` with active client Player ID and current map cell, clears
campaign `+3cc/+3d8=0` and `+3d0/+3d4=-1`, and deletes the cursor and
temporary selected object.
There is no null-result guard before its selected-quantity dereference.
`0x446` bypasses this body. A focused cancel button posts its own `0x446`
on Enter before the modal own-slot Enter alias for `0x445`; Tab uses common
focus traversal, reversed by the Shift latch. Generic close clears open
state, capture and focus and stores its result; the dialog dispatcher then
posts `0x44c` panel removal.

Grid move independently requests 1000 gold or the Shift-selected source
quantity and debits the client purse when its money cursor is built.

**Confidence.** High for the local predicates, pointer/scalar stores and
message convergence. Unknown for exceptional pointers, native transaction
reach and complete drag cancellation/refund chronology.

### MENU-094

The session opcode table binds `0x23` to `L06500`. It resolves packet
Player ID, obtains that Player's first actor, requires amount `>0` and
`<=Player+38`, then subtracts amount from `+38`. Independent column/row
distance tests allow `2`; a farther request uses the first actor's cell.
It calls `R0472(position,null,amount)` and ignores its return.

The wrapper creates a Sack with a fresh empty container when none exists,
or adds to existing Sack `+3c`; it recomputes Sack `+1c`. Registration refusal
deletes the new Sack and returns zero after the command's Player debit.
The local command arm contains no refund call on that return.

Pickup `R0448` credits positive Sack `+3c` to the looter's owner Player
through `R0449`, which writes Player `+38` and sends its new purse
value. It transfers Items and deletes the Sack. The selected client money
object and server Sack are different objects with different field layouts.

**Confidence.** High for the raw opcode table, complete command arm and named
wrapper/pickup bodies. These predicates exclude unconditional server debit
and selection of an arbitrary current client actor as the server's actor.

**Unknown.** Whole-runtime registration-refusal reach, external rollback,
client reconciliation and exceptional allocation/pointer paths.

### MENU-095

The four-byte archive primitives are established by SAV-PLAYER-028.
`R0415` writes current Player `+38` as a four-byte scalar through
`R1273` XOR `0x5c073f4d`; its LOAD arm reads into `+38` and applies
the same transform. `R1274` calls Token serialization, writes/reads
Sack `+3c` as four bytes and serializes its existing container at `+40`.
MENU-094's drop and pickup bodies consume those same server fields.

The modal's client Player `+0c`, temporary selected object's `+10` and queued
command amount are separate sources. The named server scalar operations
do not substitute the client preview value or reconstruct Sack gold from
an item count. This is a local serializer join, not a complete client
archive/call-graph absence claim.

**Confidence.** High for scalar source, width, symmetric transform and direct
consumer join. Unknown for native SAVE/LOAD/pickup continuation and for the
state a SAVE observes between client debit and command execution.

## Typed chat commands, the privilege byte and host-console input

Population: the `rom.exe` image (one image for the EN and RU installs), the chat opcode `0x91` arm of `R0061`, the typed-command parser `R0430` and its 24-literal block at `L13195`. All command words are ASCII bytes in the image, so EN and RU accept the same words. Instrument: `tools/ghidra` `DisasmFn`, `DisasmRange` and `EnumRefs`; listings in `experiments/EXP-0492-cheat-commands/evidence/listings/`. Nothing was run.

| ID | Claim | Confidence | Status | Evidence |
|---|---|---|---|---|
| MENU-099 | The chat line reaches one typed-command parser, `R0430`, through one call (`L13196`) taken when the first character is `#`; its chain tests 13 literals (12 commands) in a fixed order with a case-sensitive prefix match. | High | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-100 | The parser's first test is `server+0x0c != 0`, which returns silently, so no `#` command, `#Chicken` included, acts on a map whose participant flag is set; the notice and privilege tests below come after it. | High / Medium | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-101 | Nine arms of the 12 commands test the privilege byte through `R0444` (10 direct call sites, all in the parser); `#modify`, `#event` and `#Chicken` call no such test, so `#modify +knowledge` is not privilege-gated. | High | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-102 | `#Chicken` is a prefix match that sets the sender's privilege byte to `0xff`, logs one line naming the Player and sends notice 5; it has no gate besides MENU-100. | High | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-103 | The privilege byte `Player+0x68` is 0 from the constructor, 0xff after `#Chicken`, 0 after `#kill cheaters`, and is not written by `Player::Serialize`, so SAVE does not carry it and a LOAD restores 0 by the constructor (Medium). | High / Medium / Unknown | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-104 | `#create [N ]<name>` credits N gold for the name `Gold` and otherwise adds a named item of count N to the hero's inventory, under the participant flag, the single-player flag, the privilege byte and a hero test. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-105 | `#modify <self\|army>` has four arms with no privilege test: `+god` (self or whole army), `+spell <id>` and `+spells` (self only) and `+knowledge` (either); each ends in notice 7. | High | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-106 | `#summon [hero ]<name>` spawns N creatures or one hero by type name at the caller's hero, under the privilege byte, owned by the caller. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-107 | The four kill commands require the privilege byte and set every actor of the target Players to a health word of `0xffce` (death on the next tick is Medium); `#kill cheaters` also clears the targets' privilege byte and sends no notice. | High / Medium | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-108 | `#pickup all` moves every Sack on the server's list into the caller's hero through the ordinary pickup routine, crediting its gold and inventory, under the privilege byte. | High | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-109 | `#show map`, `#hide map` and `#victory` send opcode `0xaa` with selector 1, 0 and 2 to the caller; the client reveals the whole map and sets a runtime flag, clears the flag, or runs the mission-win arm; no server state is written. | High / Medium | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-110 | `#event <n>` has no privilege test and sends packet `0xb6` with n to the caller, which opens the mission event text panel n; it sends no notice. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-111 | Three broadcast notices (`0x92` subtypes 5, 6 and 7) tell every client that a Player became a cheater, used a command ineffectively, or used a command successfully. | High | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-112 | The phase-3 host console panel accepts two further `#` commands with no privilege test: `disconnect <id>` and `curse <id>`. | High / Unknown | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-113 | Outside the parser and the Alt console, no further cheat-like key or text input was found in four bounded populations; 28 command-line switch literals are named and only `-trace` was read. | Medium / Unknown | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |
| MENU-114 | Each command's state change reaches the SAV only through the ordinary records it edits; the three non-saved effects are the privilege byte, the map reveal flag and the `+knowledge` Diary resend. | High / Medium | ✔ promoted (branch candidate) | [EXP-0492](../experiments/EXP-0492-cheat-commands/) |

### MENU-099

The chat producer `R2123(text, recipient)` builds a record with `+9 = 0x91` and the text at `+0xf`; it is called by the Enter handler `R2124` and by the frame handler `R0709`. The server arm `L09396` of `R0061` formats `"%d|%s: %s"` for display and, when the first character is `#` (`L13197`), calls `R0430(msg, Player)` at `L13196`. That is the only caller of `R0430` (`rom-enum.txt`, `callto:R0430`: 1 hit).

The parser tests each literal with `CString::Find` (`R0996`, which calls `strstr`) and takes a match only at index 0, so every test is a case-sensitive prefix test: `#Chicken123` matches `#Chicken`. The chain order, by address of the test: `#create ` `L13198`, `#modify ` `L13199`, `#summon ` `L13200`, `#killall` `L13201` or `#kill all` `L13202`, `#kill cheaters` `L13203`, `#kill ` `L13204`, `#pickup all` `L13205`, `#show map` `L13206`, `#hide map` `L13207`, `#victory` `L13208`, `#event ` `L13209`, and `#Chicken` `L13210` as the last test. An earlier test shadows a later one, so `#kill cheaters` must precede `#kill `, which it does.

The literal block at `L13195` holds 24 strings: the 13 command literals above (12 commands: `#killall` and `#kill all` share one arm), the argument words `Gold`, `self`, `army`, `+god`, `+spell `, `+spells`, `+knowledge`, `hero`, the log strings `All sacks picked up`, ` enable cheating.` and `Player `. The 89 call sites of `R0996` in 30 owner routines are enumerated in `rom-enum.txt` (`callto:R0996`, 0 in orphan code). Their classification from string operands, as CSV, INI, registry or dialogue-script parsers rather than chat input, is a lane note not in the committed evidence (Medium).

**Confidence.** High: the parser is read whole and the literal block is dumped (`rom-tables.txt`). Medium for the classification of the other 29 owners of `R0996`.

**Unknown.** Whether another code path posts a `0x91` record that begins with `#` without the producer above: the opcode `0x91` senders enumerated here are the chat producer and the console reply `R0443`, which writes `+7 = 0` (`AI-398`). The `:` arm of the same chat arm (`L13211` to `L13212`, `rom-chat-arm.txt`) relays by team bit or plainly and parses nothing.

### MENU-100

`R0430` is a server method. Its first instruction sequence tests `[this+0xc]` and returns when it is nonzero. `server+0x0c` is "more than one participant": `L02087` stores `(map+0xd4 > 1)`, and the server constructor stores `arg0 < 2`. The single-player flag `server+0x14c` is recomputed as `(+0x0c == 0)` by its two writers (`SHOP-ENTRY-016`). `#create` additionally tests `server+0x14c`.

A map whose participant count `map+0xd4` exceeds 1 therefore ignores every typed command, `#Chicken` included. The lane read `map+0xd4` as 1 on all 56 embedded campaign maps and as 4, 8, 12 or 16 on 15 of 16 loose maps; that census is a lane note and is not in the committed evidence.

**Confidence.** High for the gate and the stores. Medium for which shipped maps pass it, because the per-map values are not committed.

**Unknown.** Whether a map can be played with `map+0xd4 > 1` in a session whose server constructor argument is below 2.

### MENU-101

`R0444` is the test `Player+0x68 > 0x32` (unsigned byte). `EnumRefs callto:R0444` (direct calls; `rom-enum.txt`: 10 hits, 1 owner, 0 in orphan code) finds exactly 10 call sites, all inside `R0430`; an inline `+0x68 > 0x32` test is excluded only through the `disp:68` census (`MENU-103` bounds it): `L13213` (`#create`), `L13214` (`#summon`), `L13215` (`#killall` and `#kill all`), `L13216` (`#kill cheaters`, the sender), `L13217` (`#kill cheaters`, each other Player in its loop) and `L13218` (`#kill <name>`), `L13219` (`#pickup all`), `L13220` (`#show map`), `L13221` (`#hide map`) and `L13222` (`#victory`). A refused command sends notice 6 (MENU-111).

`#modify` (every arm), `#event` and `#Chicken` have no call to it. The `#modify +knowledge` arm re-sends the Diary; it tests the separate threshold `Player+0x68 > 10` (`R2125`, sole caller `L13223` in `L08064`) only to choose the count it sends: `0xffff` for every nibble when the byte exceeds 10, the Player's real counts otherwise (`UNIT-147`, `SAV-845`). This answers the `UNIT-147` Unknown on whether `#modify +knowledge` is gated: it is not.

**Confidence.** High: the 10 sites come from the call enumeration and each arm is read whole.

### MENU-102

After the earlier tests fail, `L13210` tests the line against `#Chicken` (`L02012`). A match does three things and has no other gate: it formats `"Player <name> enable cheating."` and posts it as a frame message (`R1347`), stores `0xff` in `Player+0x68` through `L13224` calling `R0445`, and sends notice 5 (MENU-111). Since the test is a prefix match, any line that begins with `#Chicken` qualifies.

**Confidence.** High; this replaces the Medium prefix-versus-whole-line statement of `AI-378`.

### MENU-103

`Player+0x68` is a byte. Byte-width accesses in the image: the constructor store `L02014` (value 0, `R0201`, vtable `L07944`), the test at `L13225`, the setter at `L13226`, the `>10` test at `L13227` and the console test at `L02001`. The setter `R0445` has two call sites, both in the parser: `0xff` at `L13224` and `0` at `L02013` for each other Player whose byte exceeds `0x32` (MENU-107).

`Player::Serialize` (`R0415`) reads and writes no `+0x68` (`rom-player.txt`). A Player rebuilt from the archive is made by the class factory `R0203`, which calls the constructor, so after LOAD the byte is 0 by that route.

Bounds. "Byte-width accesses" means the `disp:68` enumeration (`[reg+0x68]` operand form, `rom-enum-disp.txt`: 750 hits, 325 owners) plus the call enumeration of the setter; it excludes other operand forms, address-of forms and bulk copies. "Reads and writes no `+0x68`" is true of the body of `R0415`; its base call `R0530` and the callees that take the archive were not read. Wider-width stores at displacement `0x68` belong to many classes; those sampled are other objects and the classification of all of them is partial.

Lifetime across mission entry, mission end and the town: the Player object persists in memory across the campaign path (`PARTY-PERSIST-014`, `-028`), which suggests the byte survives them in one process, but the five other constructor callers (`L13228`, `L13229`, `L13230`, `L13231`, `L13232`) were not read for a mission-entry rebuild. Whether the town has a chat entry is Unknown: the construction site of the chat edit control was not read.

**Confidence.** High for the constructor value, the two writers and the absence from `Player::Serialize`. Medium for the LOAD value, because the other archive-side writers of the byte are classified only by sampling. Unknown for survival across mission entry, mission end and town entry within one process.

### MENU-104

The arm tests, in order: `server+0x0c == 0` (MENU-100), `[L00285]+0x14c != 0`, the privilege byte (refusal: notice 6), and, when the hero actor's byte `[Player+0x34]+0x13c` is above 0, sends notice 6 and returns (`L13233` to `L13234`). The count is the token before the first space: `L13235` parses it with `atoi`; a positive value is the count and the rest of the line is the name, otherwise the count is 1 and the whole text is the name.

For the name `Gold` the arm calls `R0449(Player, count, 0)`, which adds to `Player+0x38` and sends opcode `0x67` with the new balance, then notice 7. For any other name it calls the item factory `R0947(L02110, name)`; a null result, or a failure of the validity test `L09852`, gives notice 6. Success stores the count at `item+0x42`, adds the item to the hero's inventory (`Player+0x34`, container `+0x7c`, `R0929`), runs `R0451`, projects the actor with `R0059` mask `-1` and sends notice 7.

**Confidence.** High for the gates, the count syntax and the gold arm. Unknown for the item-name grammar of `R0947`: its families go through sub-databases at `L04589`, `L04593`, `L04591` and `L04587`, and the name set it accepts was not decoded.

### MENU-105

`#modify ` is followed by `self` or `army` and then one of the `+` words, tested in this order: `+god`, `+spell `, `+spells`, `+knowledge`. The arms:

- `+god`: `R1562` on the hero (`self`) or on every actor of `[Player+0x20]` (`army`). It writes the six modifier protection words and six damage-kind bytes to 100 and recomputes (`HERO-MODDK-161`), then projects with mask `0xbf7fff7f`.
- `+spell <id>`: `self` only; needs the hero's spellbook `actor+0x140` and `0 < id < size(L05046)`. It builds `Spell(id)` (`R0463`) and inserts it (`R0464`). An out-of-range id still projects and sends notice 7.
- `+spells`: inserts ids 1 to 28.
- `+knowledge`: `L08064` re-sends the Diary (MENU-101, MENU-114); `self` and `army` behave alike.

None tests the privilege byte; each ends with notice 7.

**Confidence.** High for the arms as read whole. The 28 of `+spells` is the loop bound read from the code, not the shipped spell count.

### MENU-106

After the privilege test and a non-null `Player+0x34`, `#summon ` takes an optional `hero ` word. With `hero` the count is forced to 1 and the hero flag is set. Otherwise the token before the first space is the count as in MENU-104. The arm loops `R0066(server, &name, actor, heroflag)`, which tries `R0501(name)` and, when the type word is 0, falls back to `R0497(name, flag, 0)`. The owner is the owner of the actor passed in, which is the caller's hero.

**Confidence.** High for the gates and the loop. Unknown for the accepted name set and for the placement cell of each spawn: `R0501` and `R0497` were not read.

### MENU-107

All four arms test the privilege byte first (MENU-101). The kill is the helper `L06591`, which writes the health word `+0x94 = 0xffce` (-50) on every actor of `[Player+0x20]`, which is a health at or below the death threshold of the tick routine; the death arm was not read here, so "dies on the next tick" is Medium. Whether `#killall` includes the caller is Unknown: the matrix test on `[caller][caller]` bit 0 does not exclude it.

- `#killall` and `#kill all` (prefix match, so `#killall` also takes any suffix): every Player whose relation bit 0 toward the caller is set in the matrix `[L00004]+0xa9c4 + 0x32*P.id + caller.id`; then notice 7.
- `#kill cheaters`: every other Player whose privilege byte exceeds `0x32` has the byte set to 0 (`L02013`) and its actors killed. No notice.
- `#kill <name>`: the Player found by name through `L13236` (`R2126`, `R1387`, `R0751`); its actors are killed; notice 7.

The name compare `R0751` is a byte compare (a multibyte-aware compare when the locale flag `[L03507]` is set); it is not a case-insensitive compare, in contrast to `R0911` (`_stricmp`).

**Confidence.** High for the arms and the helper. Medium for the name compare being case-sensitive: the locale-flag branch (lead-byte table `L13237`) was read only to its first lines.

### MENU-108

After the privilege test and a non-null `Player+0x34`, the arm walks every Sack in the list at `[server+0x14]+8`. Per Sack it removes it from the grid (`R0447`) and from the list (`L02016`) and calls `R0448` with the caller's hero. That is the ordinary pickup (`MENU-094`): credit `Sack+0x3c` gold through `R0449`, pour the container into the hero's inventory (`R0450`), destroy the Sack and project with mask `0xa08000`. Then notice 7 and the frame log line `All sacks picked up`.

**Confidence.** High for the arm and the pickup routine.

### MENU-109

All three arms test the privilege byte, then send opcode `0xaa` through the session with the caller as addressee and the selector at `+0xa`: 1 for `#show map`, 0 for `#hide map`, 2 for `#victory`. The first two also send notice 7; `#victory` sends none. No server state is written.

The client arm is `L10592` (byte table `L02523`, dword table `L02524`, `rom-tables.txt`):

- Selector 0 stores `[L01661] = 0`.
- Selector 1 stores `[L01661] = 1` and ORs `0xc000` into every word of the view's map array (`[view+0x80]`, data at `+0xc`, width `+4`, height `+8`). The flag freezes the periodic fog clear (`TERR-FOG-085`) and forces the enemy-card level to 7 (`UNIT-146`).
- Selector 2 posts frame message `0x430`, the win arm that packet `0xb5` also reaches: it sets `campaign+0x3bc = 1` and either posts `0x41d` (side missions) or builds the Victory and Continue panel.

**Confidence.** High for the sends and the three client arms. Medium for the effect of `0x430` on any SAV-carried flag: the arm was read to `campaign+0x3bc` only.

### MENU-110

`#event ` takes the rest of the line as a number n and calls `R0188(mgr, Player, n, 0)`, which sends packet `0xb6` (client message `0x433`, event text n; `DLG-PATH-002`) to the sender. There is no privilege test, no notice and no check that n names an event of the loaded map.

**Confidence.** High for the send. Unknown for the display of an n without text, which depends on the client panel.

### MENU-111

`R1343(code, playerword, 0)` calls `R0623(0x92, code, playerword, 0)`: a broadcast record with subtype = code and the Player index at `+0xe`. The client arm in `R0509` (jump table `L13238`) formats a three-part line from the string table entries 221 to 226:

- Subtype 5 (entries 221, 222): the Player has decided to become a cheater. Sent by `#Chicken`.
- Subtype 6 (223, 224): the Player used cheats ineffectively. Sent by a refused privilege test and by a failed `#create`.
- Subtype 7 (225, 226): the Player used cheats successfully. Sent by the arms that end in notice 7.

**Confidence.** High for the send sites and the subtype arms.

### MENU-112

`R2127` is a key handler of the class built by `R2128`, constructed by `R0099` at `L13239` only when `campaign+0x6bc == 3`, the phase-3 host console panel. On Enter, text that begins with `#` is lowercased (`R0723`) and passed to `R2129`; other text is sent as an ordinary `0x91` chat record to `L00522`.

- `disconnect <id>`: finds the connection by id (`L08057`) and the Player (`L00171`), then `L12801` disconnects the client.
- `curse <id>`: for a Player with a hero, sets, on the hero actor reached through `Player+0x34`, `+0x130 = 0` (total experience), the stat words `+0x84`, `+0x86` and `+0x88` to 10 and `+0x8a` to 1, runs `vt+0x50`, projects with `R0059(hero, 0, -1)` and `L12802(mgr, Player)`.

There is no privilege test. These commands act on remote Players in a multiplayer host session and are lowercase only, because the lowercasing precedes the compare.

**Confidence.** High for the dispatch, the two literals and the `curse` stores. Unknown for what `L12801` and `L12802` do beyond their names and for the meaning of the four stat words (`HERO-STAT-001` names the offsets).

### MENU-113

Searched: (a) the 24-literal command block and the 89 `R0996` call sites; (b) the command-line switch strings in `.data` and `.rdata`; (c) the mission key table (`AI-KEY-125`, `keyboard.tsv`) and the Alt console; (d) the two `#` consumers MENU-099 and MENU-112. Found no further typed command or cheat key.

The 28 switch literals are `-saveonserver`, `-internetserver`, `-aslfile`, `-startserver`, `-window`, `-safevideo`, `-detail0`, `-detail1`, `-detail2`, `-noanimation`, `-noshadows`, `-nodynamiclighting`, `-trace`, `-nomusic`, `-protocol0` to `-protocol4`, `-protocol`, `-serverid`, `-name`, `-mage`, `-female`, `-waitforever`, `-emulation`, `-systemmemory`, `-map"`, `-session"` and `-ip"`. `-trace` is read at `L06538` into `[L04662]` and gates 22 diagnostic message-line posts (`MISSION-MSGPOST-058`). The effects of the others were not read.

**Confidence.** Medium: a bounded string and key search, not a read of every input route.

**Unknown.** The effects of the 27 other switches; registry and INI values that unlock behaviour; any input handled by data (a script command) rather than code.

### MENU-114

By command, what SAV carries (`SAV-1172`, `HERO-MODDK-161`, `SAV-PLAYER-028`):

- `#create` gold and items, `#pickup all`, `#summon`, kills, `#modify +spell` and `+spells`: ordinary saved state (`Player+0x38`, inventory, actors, Spellbook).
- `#modify +god`: the modifier damage-kind bytes the actor archive carries.
- `#modify +knowledge`: the Diary counts are unchanged and saved as before; the effect is a client table resend only.
- `#show map` and `#hide map`: `[L01661]` is a global outside the SAV; the revealed map words are client state.
- `#victory`: writes none of the SAV-carried state itself. The `Player+0x3c` win latch is written by the reporter `R0132` (`SAV-FLAG-027`); the effect of the `0x430` arm on it was not read.
- `#event`, the notices and `#Chicken`: no SAV state.

A LOAD restores the saved records and, by the constructor route, leaves the privilege byte at 0 (MENU-103, Medium: the other archive-side writers were sampled), so a cheat effect that lives in a saved record survives SAVE and LOAD while the privilege to repeat it does not.

**Confidence.** High for the serializer joins named. Medium for the three non-saved effects, which rest on the absence of a store in the serializers read. Unknown for the win latch.
