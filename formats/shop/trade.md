# Trading and shop interface

[Reference](format.md)

## Campaign offer handoff

The reached shop show tail tests campaign offer count `+638`, reads the head
word through array header/data `+630/+634`, and builds `shop\npc31m<word>`.
After normal builder return it registers the word and removes array element 0,
count 1, without testing the return value. A failed first text part does not
supply a rejection result to this caller. Earlier room initialization, resource
exceptions, native entry and downstream registration effects remain Unknown.
(DIALOGUE-048)

## Trading

A customer opens the shop through `Shop::Open`, which creates a `CMultiShopInstance` (max 250
live). Items move between the shelves and the instance's tray at `+0x78`; the tray element's
`+0x14` is its current owner — `0` = the shop, otherwise the customer's `Player`.

```
buy   (R1628): for each tray element with owner == 0:
                        total = (u16)elem+0x42 * elem+0x1c
                        if (Player+0x38 < total) STOP THE WHOLE COMMIT
                        Player+0x38 -= total;  owner = Player;  -> customer+0x7c

sell  (R1572): for each tray element with owner != 0 and elem+0x1c != 0:
                        Player+0x38 += ftol(0.5 * ((u16)elem+0x42 * elem+0x1c) + 0.5)
                        owner = 0;  -> back onto a shelf
```

So **the shop pays exactly half of what it charges**, rounded half-up, and an unaffordable
item aborts the rest of the purchase rather than being skipped — **but see *The counter* below:
the shipped client never sends a basket that can reach the abort** (`SHOP-BUY-009`,
`SHOP-TRAY-026`).

The `ceil` is applied to `quantity × price` as a whole, not per unit. Selling `n` odd-priced
items as one stack pays `ceil(n·p/2)`; selling them one at a time pays `n·(p+1)/2`, up to
`⌊n/2⌋` more. **A consumer that halves per unit will disagree with the engine by up to one coin
per pair, on odd prices only** (`SHOP-SELL-010`).

The unit price `item+0x1c` is a property of the item. Neither commit path reads anything of
the customer except the money at `Player+0x38`, and the generator never sees a customer — so
two customers at one shop pay the same (`SHOP-PRICE-011`).

### Where the price comes from, and every rounding

`__ftol` (`R0279`) rounds **toward zero**. The named base, non-cast,
cast and sell paths have these rounding boundaries. The former complete
site-count clause of SHOP-ROUND-017 is retracted (`SHOP-102`,
`SHOP-103`, `SHOP-ROUND-017`):

- Armor (`R0960`), Shield (`R0961`) and Weapon (`R0962`) all set `item+0x1c = ftol(base * material.+0x30 * tier.+0x30 + 0.5)`, which rounds half up. `base` is parameter slot 2 of the item's own `Data.bin` entry; `material` is `Materials[item+0x46]` (table `L02110`) and `tier` is the entry `[item+0x45]` (table `L04635`).
- Enchantment addend: `R1720(n) = ftol((1.5^(n/70.0) + 1.0) * n * 50.0)`, which truncates. `item+0x1c` grows by that amount and is then clamped to 9 999 999.
- Cast addend: `R1780 = ftol(10*S*2^log_1.2(1+power/30))`, which truncates.
- Sell payout: `R1572 = ftol(0.5 * (qty * price) + 0.5)`, which rounds half up.
- Buy: `R1628` uses integer multiplication only, no floating point at all.

Weapon snapshots write stored `Item+0x1c` as dword tag 1. The selected quote
accessor reads tag 1. For quantity 1 and stored price `1 <= P <= 9,999,999`, the
selected quote and commit bodies calculate buy `P` and sell `ceil(P/2)`.
Neither commit body has a staff-specific price branch. The wire-record to
client-object ingestion and native displayed quote remain Unknown. The
odd-price quantity-2 controls retain the per-unit display versus per-stack
payout difference described above. Arbitrary loaded prices and overflowing
intermediates are outside this agreement claim (`SHOP-104`).

These consumers use the stored price. The named Token LOAD path restores it
from the archived dword, and the direct Item/Weapon archive bodies do not
reconcile it with the loaded effects. The reached Weapon price method
recomputes unconditionally when called; text effect attachment explicitly
calls it. The named paths establish no price-valid flag. A serialized price
can remain authoritative at the direct snapshot/merchant consumers even
when fresh recalculation would differ. Transitive archive callbacks and
post-load gameplay remain Unknown (`SHOP-105`).

### Which shelf an item returns to

`R1717`, given the item's class code at `item+0x44`:

| condition | shelf | code |
|---|---|---|
| `item+0x44 ∈ {3, 4, 5}` | 3 | 3 |
| `item->vt+0x50()` is false | 2 | 4 |
| `item+0x44 == 2` | 1 | 2 |
| `item+0x44 == 1` | 0 | 1 |
| otherwise | 2 | 4 — with the log line `"Item of strange type is returned to shop - placed to Magic Items"` |

### The counter — the table, the two buttons, the move

The shop screen is one view (vtable `L03649`) with five children, each of a class built at
exactly one site image-wide. Three are item grids; each computes `cols`/`rows` from its own
rectangle at **80 px per cell**, allocates `cols·rows` CRects, and shares one hit test that walks
exactly those rectangles (`SHOP-TRAY-024`):

```
child   rect (l,t,r,b)   cols x rows   container code (vt+0xa8)   what it is
+0x68   0,0,164,303        2 x 3 = 6   parent+0x132 + 5 (5..8)    the shop's four shelves, scrolling
+0x70   0,303,480,390      5 x 1 = 5   4                          THE TABLE  (graphics\interface\ShopTable.bmp)
+0x6c   0,390,480,480      5 x 1 = 5   2                          the party backpack, scrolling
```

The other two children and the whole layout are `SHOP-SCREEN-030`…`SHOP-SCREEN-039`,
whose claim rows carry every number with its citing address:

```
child   rect (l,t,r,b)      id     background art          measured
+0x68     0,  0,164,303   1002   ShopInv.bmp             164x303   cells (1+80c, 31+80r) 80x80, 2x3
+0x74   164,  0,480,303   1005   ShopFrame.256           316x303   the merchant's room
+0x70     0,303,480,390   1003   ShopTable.bmp           472x87    cells (32+80c, 303)   80x80, 5x1
+0x6c     0,390,480,480   1001   (none)                            cells (32+80c, 395)   80x80, 5x1
+0x78   464,  0,640,238   1006   ShopMenu.bmp            176x238   the four command buttons
+0x7c   480,238,640,480      7   (composed at run time)            THE CHARACTER PANEL, borrowed
```

**Every rect above is view-relative, and the view's own rect is
`((screenW-640)/2, (screenH-480)/2, screenW-(screenW-640)/2, screenH-(screenH-480)/2)` — four
globals with one writer each in `R0341`. The shop screen is a fixed 640x480 panel centred
on the display, so at the shipped 640x480 default the origin is `(0,0)` and every rect above is
also a screen coordinate** (`SHOP-VIEW-044`). Activation clears the whole screen to colour 0 first.

`ShopTable.bmp` is 472 px wide and the table's blit asks for 480 (`SHOP-SCREEN-032`).

The shelf grid carries two 72x32 arrow rects, `(46,0,118,32)` and `(46,271,118,303)`, each
scrolling by one row of two cells (`SHOP-SCREEN-033`). The four shelves are chosen by clicking the
room picture. The backpack strip has its own two 32 px edge rects with two arrows: the left is drawn
only while the scroll base is above 0, the right only while the element count minus the base exceeds
5, and at a bound no arrow bitmap is blitted (`SHOP-110`). A click steps the base by one cell,
left when the base is not 0 and right when the count exceeds base + 5, so a strip holding 5 or fewer
elements never scrolls (`SHOP-111`). A click inside an edge rect is consumed at every base, and the
cursor over it does not depend on the bound (`SHOP-112`). **The merchant panel holds two four-rect arrays and they are different**: the hit
loop reads `panel+0x60 + 0x10*i` and the draw loop reads `panel+0xa0 + 0x10*i`, and the art loader
names the folder as `4 - i` (`SHOP-SHELF-047`, correcting partially retracted `SHOP-SCREEN-034`;
its loop bound and rect stride stand):

```
i   hit rect                draw rect               size     folder
0   354,110,459,295         353,108,433,220         80x112   shopanim\04
1   169,110,274,295         197,108,277,220         80x112   shopanim\03
2   314,  5,454,105         313, 20,445,108        132x88    shopanim\02
3   172,  5,314,105         201, 20,313,108        112x88    shopanim\01
```

Each draw rect is exactly the size of its own eleven loaded frame files. Opening a shelf resets the rack
to its first item and stores the shelf index in `view+0x132`, which is `0x64` while none is open
(partially retracted `SHOP-SCREEN-034`; its scroll reset, shelf-index store and initial value stand).

No text drawn by the shop code `L09774..L14000` names the selected shelf. That code draws 18
text runs — the grid's quantity and price, the button panel's captions and numbers, and the
identify modal's four lines — and none reads `view+0x132` (`SHOP-124`). Three paths outside that
range were not read and stay open: the shop view's inherited paint `R1311`, the tip popup's
paint and hover getter, and the borrowed character panel's paint. Within the code read, a shelf is named only by
hover help: the
merchant panel's hover getter returns `main.txt` slot `62 + i` for the hit rect `i` under the
pointer, open or not, and slot 61 for the shopkeeper rect `(274,110,354,295)`; the shared hover box
draws it above the pointer in font2 (`SHOP-125`):

```
i   hit rect           main.txt   EN                         RU
0   354,110,459,295    62         Armor                      Броня
1   169,110,274,295    63         Weapons                    Оружие
2   314,  5,454,105    64         Magic items                Магические предметы
3   172,  5,314,105    65         Scrolls, books & potions   Свитки, книги и пузырьки
```

The identify modal draws slots 79 and 80 and the executable literals `Yes` and `No`, so the RU
install shows the English `Yes` and `No` there (`SHOP-126`).

The merchant himself is drawn at `(277,112)`, 76x176, from
`movies\shopanim\Pose2-3\1.bmp` — two immediate offsets
`(+113, +112)` from the merchant panel's own top-left, not a rect. His Yes and No animations
replace him at the same anchor (`SHOP-MERCHANT-046`).

#### Shop-interior progression

Field offsets and mode masks in this section are hexadecimal; counts and
milliseconds are decimal.

Native visual timing remains Unknown. Merchant-panel paint `R1539` is vtable
`L09658+2c`. When `view+148` is nonzero and unsigned
`timeGetTime()-L09660 >= 100`, it calls the sound hook, stores a new clock
sample, advances enabled racks, then advances at most one merchant mode. Drawing
follows that block even if the block was skipped. No elapsed-time catch-up loop
is present. Thus the local gate is 100 ms, not guaranteed visible10 fps. The
initial clock value is `now-100`; the two paint timestamps are static globals,
not one clock per series (`SHOP-ANIMATION-081`).

| series | start and normal progression | termination or replacement |
|---|---|---|
| Rack `i=0..3`, folder `4-i` | Selection loads files 1..11, sets bit`1<<i`, resets index 0. Eligible paint increments; reaching 9 rewrites the index to 3. The steady loop uses indices 3..8, files 4..9. | Switching sets the old rack index 9. A paint before the next eligible advance may draw file10. Advance 9→10 disables/frees it before drawing; file11 is not a normal draw. |
| Base merchant | `Pose2-3/1.bmp`, `panel+128`; drawn when no merchant mode bit is set. | Replaced by the highest-priority enabled mode below. |
| Idle, bit 10 | One bitmap at+12c is replaced per step; index+254 increments modulo30 and loads `Pose2-3/(index+1).bmp`. From 0, files 2..28 are drawn. | At index 28, file29 is loaded, the bit clears and the index resets before the draw. |
| Yes, bit 20 | Array+1e0 has base pose in slot 0 and Yes files 2..12 in slots 1..11. | Advance to 12 clears the bit, resets index 0 and frees the array before the draw. |
| No, bit 40 | Array+210 has base pose in slot 0 and No files 2..12 in slots 1..11. | Same twelve-step boundary and cleanup as Yes. |

The first rack/base reaction frame can be skipped if the first paint already
advances it. Selecting the current shelf does nothing. Entry selects rack 0,
folder 04; clicking a different rack also arms Yes. The four rack folders each
contain 11 files; Pose2-3 contains 29, Yes13 and No12, with equal corresponding
file hashes across roots. Yes/1, Yes/13 and No/1 are not loaded by these reaction
loaders. Counts describe inventory, not frames guaranteed visible
(`SHOP-ANIMATION-082`, `SHOP-ANIMATION-083`).

Idle is armed only when none of bits 10/20/40 is set and elapsed time since
`L09661` reaches a freshly computed `5000+1000*(rand()%5)` threshold. The
threshold is recomputed at every eligible pass, including passes whose mode
bits prevent arming. It is not a single uniformly sampled 5–9 second delay.
Completion of each merchant mode refreshes the idle timestamp
(`SHOP-ANIMATION-081`).

Both advance and draw prioritize idle, then Yes, then No. Sell-with-items and
affordable buy load Yes and OR20; unaffordable buy loads No and OR40. These
triggers and the rack-click Yes arm do not reset the shared index or clear other
mode bits. Repeated input therefore is not a proved fresh restart. The two
parameterized constructors zero the index. Exit clears flags and frees static
and animated resources; entry reloads resources and resets rack selection.
Neither named transition resets the merchant index or static timestamps.
Object reuse, inherited transition effects and actual mid-animation re-entry
remain Unknown (`SHOP-ANIMATION-084`).

The sound hook runs before frame advancement. While idle it requests view+b4
at index 1, view+94 at 10/14/18/22, and view+b8 at 24. With no merchant mode it can
request view+b0 after its separate redrawn 30-second-base threshold,
`L09665(view+70+20b0)==0`, and non-null/guard tests. The table-query helper's
meaning is unexpanded. Shelf selection requests+a0; sell+a8;
affordable buy+a4; refusal+90; entry+ac and conditionally+b0. The common request
helper requires a non-null sample and guard result 0. Its priority128 argument
is not a sound id. Sample identity, audible overlap, presentation timing and
the unexpanded entry+bc helper remain Unknown
(`SHOP-ANIMATION-085`, `VIDEO-SFX-013`).

The button panel's four rects, top to bottom, are `(494,15,614,67)` clear the table with the purse
printed on it, `(483,67,623,113)` **buy** with the buy total, `(483,114,623,160)` **sell** with the
sell total, and `(494,160,614,212)` clear and leave with the sum (`SHOP-SCREEN-035`).

One cell draws a background chosen from five (`backinvs.bmp` for a shelf item that is both
affordable and usable, `backinv.bmp` for a usable one elsewhere, `backinvg.bmp` otherwise), then the
item icon, then the quantity as `"%d"` at `(cell.left+10, cell.bottom-15)` when it exceeds 1, then a
price plaque right-aligned to the cell's right edge. The plaque is `costm(d+1).bmp` on the player's
side and `costs(d+1).bmp` on the shop's, `d = floor(log10(price))` clamped to 6, and the number on
it is `ceil(price/2)` on the player's side and `price` on the shop's (`SHOP-SCREEN-036`,
`SHOP-SCREEN-037`). The grids draw no item name and no characteristics panel
(`SHOP-SCREEN-039`).

The plaque's text is the number alone, with no label on either root: `"%d"`, grouped in threes
with commas (`1,250`, `9,999,999`, `-1,250`; `TOWN-469`), its right end 6 pixels inside the
cell's right edge and its top 1 pixel below the cell's top, in font2 with a 1-pixel shadow. Its
colour ramp `L03616` holds `(185k/15, 159k/15, 73k/15)` in entry `k` (`MISSION-MSGLINE-056`).
Every element except the money cell and a quantity of 0 draws it, whatever the price's sign. For
a price of 0 or less `log10` has no value, and the x87 integer-indefinite result selects the first
plaque; that is derived, Medium. Across every reader found, no item-information text composes a
price, and these grids alone draw an item's own price (`ITEM-PRICETAG-144`).

Of the 134 resources the screen names, 132 are byte-identical on the EN and RU installs; the two
that differ are `main\text\tips\shop1.txt` and `shop2.txt` (`SHOP-SCREEN-038`).
Both are shown in **one** widget of a different class (vtable `L06362`, ctor
`R1261(id, l, t, r, b, text)`), id 1011, at `(0,162,312,298)` inside the **merchant panel** —
view-relative `(164,162,476,298)`. It exists only when `[L03631]` is non-zero, and `shop2.txt`
replaces `shop1.txt` at most once per visit, behind the latch `view+0x8c` (`SHOP-TIP-045`).

The money cell — an element whose `+0x6` is `0xffff` — draws `backinv.bmp` and then
`graphics\interface\money\money.16a`, a single 80x80 `.16a` frame that
composites over it (`SHOP-MONEY-048`).

### The character panel, `+0x7c`

The sixth region is **not built by the shop**. It is `campaign+0xe0`, the third of four widgets the
main-frame builder stacks in the mission screen's 160-pixel right strip, at `(0,238,160,480)`. On
activation the shop takes it from `campaign+0xd4`, offsets its rect by the literal `640 - width`,
re-sets it and adopts it; on deactivation it offsets back and hands it over. One object, two parents
(`SHOP-FIGURE-041`).

It draws the shown character **composed in that character's own equipment**: virtual slot `+0x80` of
the member's drawable — the world figure compositor, which has no layer selector, so the head is
included — into a 160x240 canvas blitted whole at `(panel.left, panel.top + 2)`. The composition is
cached on a string key at `panel+0x7c` and redone when the key changes or when `member+0x18c` bit 3
is set. When `member+0x18c & 0x11` is zero the same routine draws a flat picture instead
(`SHOP-FIGURE-042`).

The **party picker** is the panel's own, not the shop's. Two 32x32 rects, panel-relative
`(1,205,33,237)` and `(119,205,151,237)` — `(481,443,513,475)` and `(599,443,631,475)` at 640x480 —
post messages `0x414` and `0x415`. The shop view's arms step `view+0x130`, a `u16` index into the
roster `CArray` at `view+0x108`, with wrap-around in both directions, and then in the same call:

```
view+0x6c . vt+0x90(member+0xc8)        bind the item strip's list
[view+0x6c] + 0x90 = member+0xdc        a POINTER to that member's own scroll base
R1767 ; R1768             rebuild and repaint the strip
member+0x18c |= 8                       recompose the figure
view+0x14c   |= 9                       invalidate
```

So one step changes the figure and the strip together, and **the strip at `(0,390,480,480)` is the
shown member's own container rather than a party-wide backpack** (partially retracted
`SHOP-PICKER-043`; rectangles, messages and rebinding mechanics stand, while the sound-id clause
is withdrawn).

**The table is one strip of five places, shared by both sides of the deal.** It does not scroll;
its display-list add refuses to append a sixth element; and the server-side container
(`instance+0x78`) has no capacity of its own. Which side an element is on is ownership, not
position — `elem+0x14` on the server, the container code on the screen — and it is drawn with
`myitem.256` + `costm1..7` (yours) versus `shopitem.256` + `costs1..7` (the shop's), seven price
plaques per side, one per digit of the `9 999 999` clamp.

Every move is command **`0x22`** — `R0238(srcCode, srcIdx, dstCode, dstPos, qty)`, packed at
`cmd+0x0c/0x0e/0x0d/0x10/0x12`. Container codes are shared with the inventory screen: `1`
equipment, `2` the party backpack, `4` the tray, `5..8` the four shelves; the shop arm requires both
codes in `4..8`. The shelf helper `R1718` retains its published split rule:
`stack -= (qty−1)`, detach one unit through the item's `vt+0x40`, give the detached object
`qty`. The list helper `R0095` splits only when count exceeds requested quantity.
Otherwise `L04880` branches to whole removal at `L04881`, bypassing the split virtual
and returning the same Item pointer; count one/request one follows that whole-item branch.
This correction does not remeasure the shelf helper or city-sale lifecycle. The destination-4
arm stamps `item+0x14 = Player` (partially retracted `SHOP-TRAY-025`; command layout,
shelf-specific split and ownership stamp stand, while its universal-split clause is withdrawn).

The screen keeps three running numbers over the table (`R1028`) and gates three commands on
them (`SHOP-TRAY-026`):

```
purse      = campaignScreen+0x9b4 -> +0x0c
buyTotal   = -SUM over elements whose code != 2 of  price * qty        (full price)
sellTotal  = +SUM over elements whose code == 2 of  ceil(price/2) * qty  (half, PER UNIT)

0x33 buy    only if buyTotal  != 0 and purse + buyTotal >= 0   (else the refusal animation)
0x34 sell   only if sellTotal != 0
0x35 clear the table
```

**Two consequences a consumer must reproduce.** (1) Because the client refuses an unaffordable
*total* and the server debits monotonically, the buy loop's abort is unreachable through the
shipped UI. (2) The screen halves **per unit** and the shop pays **per stack**, so on an odd price
the displayed sell total exceeds the coins received by `⌊qty/2⌋` (`SHOP-TRAY-027`) — the same delta
as splitting the stack. A consumer that makes them agree has removed a visible engine behaviour.

### A held item released over the character panel

The held object `session+0x3cc` carries two separate tags. Its origin code `session+0x3d4` and index
`session+0x3d0` are written by the pick-up beside the slot; the origin code is the source
container's `vt+0xa8`: `1` equipment, `2` backpack, `4` tray, shelf index + 5 for a shelf
(`SHOP-096`). The element's own `+0x18` is a stamp written when it enters a container: shelf index
+ 5 for a shelf element, `2` for a backpack element, `2` or `5` for a tray element by flag bit
0x80 at contents load (a tray insert keeps any stamp other than 1), `1` for new party equipment,
`0` from a constructor, and carried unchanged by the grid take `R1264` and the base take (`SHOP-097`). The
stamp and the origin code agree for a shelf or backpack object and differ for a tray object. A shelf
pick-up writes the stamp itself (shown shelf index + 5, on the held copy) and a shelf insert writes it from
the object's own shelf index; a tray pick-up writes none (`SHOP-114`).

The panel's button-up slot routes by the stamp alone (`SHOP-098`):

```
stamp 1, 2        equip arm  panel vt+0x7c(slot = item class nibble - 1)
stamp 5 .. 8      return arm R1781 on the shop view
any other value   no release arm; replay left it held (Medium)
every path        returned 1 in replay (Medium)
```

The return arm routes by the origin code, not the stamp (`SHOP-099`): code 1 calls the equip arm
with the origin slot, 2, 4 and 5..8 call the origin container's drop slot `vt+0xa4(origin index)`;
codes 1, 2, 4 and 5..8 set a dirty bit on `shop+0x14c` (0x8 panel, 0x1 backpack, 0x4 tray only when
`shop+0x84` is 0, 0x2 shelves), and every code clears the cursor slot. The routine's own
instructions contain no call to the move command. The backpack and shelf drops insert the held object
into the shown member's pack list or the shelf list (merging equal id and stamp), issue command `0x22` (always for the backpack drop; for the shelf drop when the origin code differs from the shown shelf's), clear the cursor and post `0x46d`; a held coin element adds its quantity to the purse on
the backpack drop; no member equipment slot is written (`SHOP-115`). The tray arm calls the tray's
`vt+0xa4`, which only requests the shop speech: the tray-origin object is stored in no list and no command
is issued; its pointer is dropped by the cursor clear (`SHOP-116`, consequence Medium). A tray object stamped 2 takes the equip arm with origin
code 4 and issues command `0x22` with source code 4 (acceptance Unknown).

For the shop's literal child rectangles, the shop view hands a release to the panel only when no
capture child is set and the point is in `(480,238)-(640,480)`; the six children are tested in order, the first containing the point
receives the message and the traversal stops. Focus does not affect a button-up (`SHOP-100`).

The equip arm (`SHOP-101`) rejects and drops the object through the session backpack grid unless
the panel host flag `+0x140` is 1, the member belongs to the host player, the usability test
passes, a member-state test passes for the first two slots, and an item attribute `0x2a` test passes.
On success it stores the element in the member slot, issues command `0x22` with the origin code and
index as source and container `1` position `slot+1` as destination, recomposes the figure and posts
`0x46d`. A stackable element equips one unit and returns the rest.

### Duplicates

Two items are "the same thing" iff `R0928` says so: equal item code, then **both** stackable →
yes without reading effects; one stackable → no; **neither** stackable → only if their ordered
effect lists at `+0x20` agree. Differing enchantment therefore separates equal-code non-stackable
items. Equal-code Potions are stackable even with effects and compare equal without examining them.

The 100/100/20/`rand(1..8)` random draws use plain `CObArray::Add`, so equal draws remain separate.
The six literal Potions and items *returned* to a shelf instead use the merging insert
`R1743`; the family's other merge routine, `R1744`, has 0 callers. A fresh shelf can
therefore hold duplicate random draws as separate objects — 200 draws from at most 367 admitted
triples — while literal-Potion additions and returned equal items merge (partially retracted
`SHOP-DUP-028`; random append, return-path merge and caller counts stand, while universal
enchantment separation is withdrawn).
