# ROM2 town presentation

The named fresh campaign route enters generic town ID 1, then Kaarg town
ID 2. The second town uses Kaarg square, inn and shop art. Route confidence
is Medium; positive resource and square instruction contracts are High.
The scope is selected native paths in the preserved EN/RU clients
(`R2-SESSION-115`, `R2-ENGINE-231`).

## Town selection and resources

The current record's ID at +4 selects the view. EN dispatcher `L2.00266`
and RU `L2.00267` use these arms. Playback depends on the music-enable
word; the three music keys exist in both music archives
(`R2-ENGINE-231`, High).

| ID | EN / RU arm | Member | Art | Music key | EN / RU loader |
|---|---|---|---|---|---|
| 1 | `L2.00279` / `L2.00281` | +0x110 | generic | `music\b14.wav` | `L2.00565` / `L2.00566` |
| 2 | `L2.00264` / `L2.00265` | +0x118 | Kaarg | `music\b16.wav` | `L2.00567` / `L2.00568` |
| 3 | `L2.00280` / `L2.00282` | +0x114 | druid | `music\b15.wav` | `L2.00569` / `L2.00570` |

The loaders push the following graphics keys. Paths are relative to
`graphics.res`; native strings prefix them with `graphics\`. Brace groups
enumerate literal keys; percent expressions preserve filename formats.
Continuation paths in a cell retain its leading interface directory
(`R2-ENGINE-231`).

| Art | Background and mask | Highlights and additions | Animations |
|---|---|---|---|
| generic | `interface/town/{townmain,townmask}.bmp` | `interface/town/{Tavern_l,Trener_l,Shop_l,Town_add}.bmp` | `interface/town/sign/V%.2d.bmp`, `door/T%.2d.bmp`, `stars/S%.2d.bmp`, `fluger/F%.2d.bmp`; `interface/townbirds/{tavern,fighter,mage,shopie,Guards}/sprites.16a`, `Birds%d/sprites.16a`, `HORSE%d/A%d/sprites.16a`, `BABA%d/A%d/sprites.16a`, `DERVISH%d/sprites.16a` |
| druid | `interface/town_druid/{townmain,townmask}.bmp` | `interface/town_druid/{hili_tavern,hili_shop}.bmp` | `interface/town_druid/woman/a%d%04d.bmp`, `man/a%d%04d.bmp`, `bug/sprites.16a`, `Lizard/sprites.16a` |
| Kaarg | `interface/town_kaarg/{townmain,townmask}.bmp` | `interface/town_kaarg/{hili_tavern,hili_shop}.bmp` | `interface/town_kaarg/dervish/d%04d.bmp`, `guard/g%04d.bmp`, `girl1/g%d%03d.bmp`, `girl2/g%d%03d.bmp`, `maingates/m%04d.bmp` |

The named stage path admits ID 1 at stage 10. Ordinary departures from
missions 10 and 20 admit ID 2 at stage 30. Ordinary Leave50 adds ID 3;
nonzero bank775 restoration can append IDs 2 and 3 independently of stage.
Availability does not establish a player's later visit order
(`R2-SESSION-115`, Medium).

The declared town/room archive population has 1101 keys per locale, with
1099 matching payloads. The two differences are druid lizard BMPs. All
measured Kaarg square and room payloads match between EN and RU. This is
a bounded archive comparison, not an asset-use census (`R2-ASSET-049`, High).

## Square layers and rectangle

The Kaarg view is 640x480. The measured EN initializer centers it inside
640x480, 800x600 or 1024x768 screens, giving origins (0,0), (80,60) or
(192,144). A nonzero force word selects 640x480. Otherwise the initializer
tests command-line -800, -1024 and -640 in that order, then the RESOLUTION
registry buffer for -800 and -1024. The fallback is 640x480; a failed
registry query supplies -640. A stored registry value can therefore select
a larger screen without a command-line switch (`R2-ENGINE-238`, amended, High).
The painter destinations in the following table add to the view origin.

Active painter EN `L2.00571` / RU `L2.00572` draws these opaque bitmap
layers in order. A zero active word returns before painting. Visible
source BMPs are uncompressed 24-bpp; native device conversion does not
establish a fixed RGB565 display (`R2-ENGINE-232`, `R2-ASSET-051`, High).

| Order | Layer | Destination | Condition |
|---|---|---|---|
| 1 | `townmain.bmp` | (0,0) | active view |
| 2 | `hili_shop.bmp` or `hili_tavern.bmp` | (328,256) or (480,196) | selector 1 or 2 |
| 3 | girl1 | (216,284) | episode frame; otherwise first frame |
| 4 | girl2 | (260,284) | episode frame; otherwise first frame |
| 5 | guard | (140,152) or (184,156) | first position for active cursor 13..37; second otherwise |
| 6 | dervish | (416,328) | episode frame; otherwise first frame |
| 7 | main gates | (152,256) | retained gate cursor |
| 8 | own overlay, then children | delegated | own overlay target is empty; child paint remains |

Conditional entry creates child ID 0x467 with rectangle bounds
(left=328, top=0, right=640, bottom=200), passed through SetRect.
Its extent is 312x200. These operands are High; whether the coordinates
are relative to the view or screen, its live purpose and visibility are
Unknown. The tip getter supplies text
keys, but the final receiver, status line, buttons and their destinations
are not established. No absence of these elements is claimed
(`R2-ENGINE-238`, amended, High / Unknown).

## Mask, hotspots and clicks

Sampler EN `L2.00581` / RU `L2.00582` rejects a missing mask or an
out-of-view point, subtracts the view origin and samples a top-origin byte
at `x+640*y`. Sprite bounds do not determine the region. Both masks have
307200 cells: 217863 zero and 89337 in eight action regions. Byte 80 maps
to selector 0x1000 but has no pixel in either preserved mask; the remaining
247 byte values return -1 (`R2-ENGINE-233`, amended, `R2-ASSET-050`, High).

Hover stores a changed selector at +0xb4. Shop selector 1 starts
`sfx\town_kaarg\Kenter2.wav` under latch +0xac and stops the inn slot;
inn selector 2 starts `Kenter1.wav` under latch +0x250 and stops the shop
slot. Gate selector 8 stops both under latch +0x254. Selector -1 resets
and clears the latches. These arms do not write the animation flag word
at +0x208. Person selectors take a switch no-op arm. Among the named
selectors, only 0, 3, 16 and 0x1000 reach the OR into +0x208; the sampler
emits only 16 and 0x1000 from that group. Other arbitrary values can take
the default OR arm. The inherited EN slot +0x4c calls hover +0x98, so hover
is also reachable outside the painter's 100 ms arm (`R2-ENGINE-233`, amended, High).

| Byte | Selector | Pixels | Region | Hover art | Click |
|---|---|---|---|---|---|
| 144 | 1 | 3792 | shop | `hili_shop.bmp`, 40x76 | 0x42a: Kaarg shop |
| 128 | 2 | 18424 | inn | `hili_tavern.bmp`, 140x128 | 0x42b: Kaarg inn |
| 160 | 8 | 4415 | gates | gate frame series | 0x442(1,0), then 0x42d: navigation posting |
| 176 | 16 | 40667 | main menu | no separate highlight in selected painter | 0x41f: main-menu popup |
| 96 | 0x200 | 3180 | girl1 | no person-hover animation arm | `kaargwoman%d` dialogue |
| 192 | 4 | 3823 | girl2 | no person-hover animation arm | `kaargwoman%d` dialogue |
| 64 | 0x800 | 4751 | guard | no person-hover animation arm | `kaargguard%d` dialogue |
| 32 | 0x400 | 10285 | dervish | no person-hover animation arm | `kaargman%d` dialogue |

Click EN `L2.00590` / RU `L2.00591` uses Scenario variable 0x300 as the
dialogue suffix. Shop, inn and gate clicks prepare departure before
posting. Selector 16 opens a 440x340 main menu at (100,100), not a school.
Its native constructor binds save, load, sound, quest end, return and exit
labels. Gate acceptance after the posted messages is Unknown
(`R2-ENGINE-234`, High for posting and identified receivers).

Tip indices are 233/236/237/235 for shop/inn/gates/menu and
359/360/361/362 for girl1/girl2/guard/dervish. The last region is labelled
sleeping man in EN and beggar in RU; the mask is identical. Final text
drawing positions are Unknown (`R2-ENGINE-234`, `R2-ENGINE-238`, amended).

## Animation frames and clocks

Kaarg has 218 numbered BMPs in seven series, plus background, mask and two
building highlights. Its selected loader and painter contain no TownBirds
key or sprite draw. This absence is local (`R2-ASSET-051`, `R2-ENGINE-231`).

| Figure | Files under `interface/town_kaarg` | Count | Geometry |
|---|---|---|---|
| dervish | `dervish/d0000..d0029.bmp` | 30 | 112x148 |
| guard | `guard/g0000..g0054.bmp` | 55 | 72x68 or 76x68 |
| girl1 subset 1 / 2 | `girl1/g1000..g1030.bmp` / `g2001..g2030.bmp` | 31 / 30 | 44x108 |
| girl2 subset 1 / 2 | `girl2/g1000..g1030.bmp` / `g2001..g2030.bmp` | 31 / 30 | 68x104 |
| gates | `maingates/m0000..m0010.bmp` | 11 | 60x84 |

Painter time uses `timeGetTime`. Unsigned elapsed time strictly greater
than 100 ms admits one advance and sets the baseline to current time;
there is no catch-up loop. The episode scheduler runs on every active
paint, including paints without advancement (`R2-ENGINE-235`, High).

Define `R(n)=floor(rand()*n/32767)%n` for the measured 15-bit helper.
It yields 0..n-1; uniformity and independence are not established. A family
starts when unsigned elapsed time exceeds its retained wait, records
current time, sets its flag and resets its cursor. Girls choose one of two
subsets with `R(2)` (`R2-ENGINE-235`).
The elapsed-wait arms do not check the family flag; a slow redraw cadence
can restart an episode that is still running (`R2-ENGINE-235`, High).

| Family | Initial wait, ms | Later wait, ms | Flag |
|---|---|---|---|
| girl1 | 2000..3999 | 3500..8499 | 0x200 |
| girl2 | 2000..3999 | 3500..8499 | 4 |
| guard | 2000..3999 | 7500..12499 | 0x800 |
| dervish | 4000..4499 | 3700..4199 | 0x400 |

An admitted advance increments each active person's cursor by one.
Equality with the vector count clears its flag and restores idle state.
Cursor zero of a girl's second subset selects `g2001.bmp`. Hovering does
not start person episodes. Delivered redraw cadence and the baseline
before re-entry are Unknown (`R2-ENGINE-235`, `R2-ENGINE-233`, amended).

Audible background ambience is Unknown. The selected scheduler requests
one of slots +0x20c/+0x210/+0x214 after 2000+R(2000) ms with R(3), one
of +0x218..+0x224 after a separate 2000+R(2000) ms with R(4), +0x234
after more than 45000 ms, and +0x238..+0x24c at guard cursor states.
These are bounded request clocks; playback was not observed (`R2-ENGINE-235`).

Each admitted advance samples the pointer for gates. Selector 8 increments
the cursor toward 10; other selectors decrement it toward 0. Direction
changes request `SFX\Town_kaarg\Kdoor1.wav` or `Kdoor2.wav` in the selected
arms. Endpoint clamps operate independently of the gate flag. Highlighting
uses these frames rather than a separate bitmap (`R2-ENGINE-236`, High).

## Room art and shared bodies

Campaign ID 2 selects Kaarg shop and inn pages at controller members
+0x100 and +0x10c. People open dialogue, gates post navigation, and the
menu opens a popup (`R2-ENGINE-234`, `R2-ENGINE-237`).

| Room | Background and other art | Own animation art |
|---|---|---|
| inn | `interface/inn_kaarg/TavernMain.bmp`, 320x480; generic `interface/inn/{manback,ManBackTalk,LUOver,LDOver,RUOver}.bmp` | `interface/inn_kaarg/taverner/a1%04d.bmp` (1..15), `a2%04d.bmp` (1..25), `a3%04d.bmp` (1..3), `a4%04d.bmp` (0..6), `a5%04d.bmp` (0..28) |
| shop | `interface/shop_kaarg/ShopMain.bmp`, 288x288; `ShopFrame.256`, one 316x303 sprite envelope; `hili_{armor,magic,potion,weapon}.bmp` | six `fire/{dark,select}/{lite,burn,cicle}` families, each six 80x88 BMPs; `movies.res:shop_kaarg/a10000.bmp`, `a1%04d.bmp`, `a2%04d.bmp` keeper art, 72x84 |

Inn taverner geometry is 248x296 / 184x188 / 44x20 / 64x64 / 84x116
for its five series. The shop movie population has a1 indices 0000..0020
and a2 indices 0001..0020. No `inn_kaarg` subtree occurs in either complete
156-node movie listing; the measured inn animation art is in graphics.res
(`R2-ASSET-049`, `R2-ASSET-052`, High).

Shop keeper updater EN `L2.00621` / RU `L2.00622` increments a counter
modulo 20. Flags 0x10/0x20 choose a1/a2 with filename index counter+1;
otherwise it uses a10000. Counter zero clears episode flags outside the
low nibble (`R2-ENGINE-237`, High).

The Kaarg inn page and center call common inn bases. The Kaarg shop
center calls the common shop base. These calls and shared inn art establish
shared structure. Layout reuse is Medium; complete layout equality and
differences limited to art are unproven. Full room destinations, clocks,
triggers and the shop-frame pixel decoder remain Unknown
(`R2-ENGINE-237`, `R2-ASSET-052`).

## Unknown boundaries

Shell text, status and buttons need a native trace from the active Kaarg
tip getter to the final renderer and destinations. Gate acceptance needs
the receivers after 0x442/0x42d. Room layouts and cadence need derived
painters and scheduling callers. Authorized original observation would
settle live pixels, device conversion and elapsed-time appearance.
Exhaustive campaign order needs the availability/event graph or an
observed route (`R2-ENGINE-238`, `R2-ENGINE-234`, `R2-ENGINE-237`,
`R2-ENGINE-235`, `R2-SESSION-115`; R2-ENGINE-238 is amended).
