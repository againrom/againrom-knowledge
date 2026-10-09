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
the receivers after 0x442/0x42d. Full Kaarg room destinations and schedules
need further derived painters and scheduling callers. Authorized original observation would
settle live pixels, device conversion and elapsed-time appearance.
Exhaustive campaign order needs the availability/event graph or an
observed route (`R2-ENGINE-238`, `R2-ENGINE-234`, `R2-ENGINE-237`,
`R2-ENGINE-235`, `R2-SESSION-115`; R2-ENGINE-238 is amended).

## First town ID 1

### First-town art

The selected first-town art has 339 keys per locale and equal EN/RU
payloads. Its 640x480 town mask contains 150 original index values;
39816 pixels use the nine recognized byte codes and 267384 use default
codes, including 266003 zero cells. The other nonzero default values
occupy 1381 cells. R2-ASSET-057.

The generic square has 36 numbered BMP animation pictures and 41 .16a
resources containing 1180 frames. These populations do not establish
native use. R2-ASSET-058. Selected .16a frames fit the native WORD-run
commands. BMP opaque copying and zero-WORD-key copying use separate
methods. Normal-palette RGB composition remains Medium because display
quantization, reduced-mode state and live frames are unobserved.
R2-ASSET-059.

The generic inn has 154 art keys. Generic shop graphics have 47 BMPs and
its movie subtree has 54 BMPs. ShopFrame.256 has a 316x303 header, but its
palette 8 pixel stream and 1743-byte tail remain Unknown. R2-ASSET-060.

## Town 1 square

Town ID 1 binds the generic native view through member+0x110. EN constructor `L2.00559` installs vtable `L2.00562` and calls initializer `L2.00709`; RU uses `L2.00700`. EN enter/painter/loader are `L2.00710`/`L2.00711`/`L2.00565`. The view is 640x480, centered in 640x480,800x600 or 1024x768 at (0,0),(80,60),(192,144). The shared initializer selects the dimensions from its force word, command-line flags and retained resolution buffer (`R2-ENGINE-239`, High).

Town 1 music is `music\b14.wav`, archive `music.res` key `b14.wav`. Playback is conditional on the music-enable word. The square loader has 19 graphics-interface literal keys. The populations below use paths relative to `graphics.res`. The school highlight and fighter/mage sprites load but have no use in the complete selected painter or advance bodies (`R2-ENGINE-240`, High).

| Resource below interface/ | Loaded population |
|---|---|
| `town/{townmain,townmask,Tavern_l,Trener_l,Shop_l,Town_add}.bmp` | six fixed files |
| `town/sign/V%.2d.bmp`, `town/door/T%.2d.bmp`, `town/stars/S%.2d.bmp`, `town/fluger/F%.2d.bmp` | indices 00..09,00..08,00..08,00..07 |
| `townbirds/{tavern,fighter,mage,shopie,Guards}/sprites.16a` | five fixed files |
| `townbirds/Birds%d/sprites.16a` | variants 1..9 |
| `townbirds/HORSE%d/A%d/sprites.16a` | selected variant 1..5, actions 1..3 |
| `townbirds/BABA%d/A%d/sprites.16a` | selected variant 1..4, actions 1..2 |
| `townbirds/DERVISH%d/sprites.16a` | selected variant 1..4 unequal to selected BABA variant |

The active painter follows this order. Destinations add to the view origin. Town_add uses a converted-zero-pixel skip receiver; other selected BMP layers use the ordinary copy receiver. Device conversion does not establish a fixed RGB565 display (`R2-ENGINE-240`, High).

| Order | Layer | Destination | Condition |
|---|---|---|---|
| 1 | townmain | (0,0) | loaded |
| 2 | selected birds | (0,0) | active group and incomplete frame |
| 3 | Town_add | (0,0) | active bird-paint branch |
| 4 | Shop_l / Tavern_l | (264,264) / (144,332) | selector 1 /2 |
| 5 | tavern sprite | (124,312) | retained frame |
| 6 | sign/V | (360,232) | retained bitmap |
| 7 | gate/T | (180,148) | retained bitmap |
| 8 | stars/S | (340,288) | bitmap non-null |
| 9 | shopie sprite | (276,296) | retained frame |
| 10 | weather vane/F | (308,64) | retained bitmap |
| 11 | guards sprite | (184,158) | retained frame |
| 12 | HORSE | selected position | episode frame, otherwise 0 |
| 13 | BABA | selected position | episode frame, otherwise 0 |
| 14 | DERVISH | selected position | retained frame |
| 15 | shared children | delegated | after square layers |

The initializer attaches child 0x445. Entry can attach child 0x467 with native rectangle (328,0,640,200), extent 312x200. Its view/screen coordinate reference, live purpose and visibility are Unknown. The generic own-overlay method is empty; shared child dispatch remains. No absence of other shell text, status or buttons is claimed (`R2-ENGINE-239`).

The selected EN delayed-tip receiver calls the child at the pointer and its tip slot. Text appears when its elapsed accumulator crosses 500 ms while below 25500 ms and the native admission fields allow it. For pointer (x,y), maximum line width w and line count n, the initial box is (x,y-5-14n,x+w+11,y); right/top overflow is shifted inside maintained screen bounds. Text begins at (left+5,top+4+14i). Live admission and other shell visibility remain Unknown (`R2-ENGINE-239`, High for the selected native receiver).

### Town 1 mask and actions

The 640x480 mask has 307200 cells and 150 distinct raw indices. The sampler subtracts view origin and reads byte[x+640*y] through a finite selector table. Every present positive selector is below; all other present indices total 267384 cells and return -1 (`R2-ENGINE-241`, High).

| Byte | Selector | Pixels | Tip | Click |
|---|---|---|---|---|
| 32 | 0x400 | 17 | none | no action |
| 64 | 0x800 | 18 | none | no action |
| 80 | 0x1000 | 4 | none | no action |
| 96 | 0x200 | 6 | none | no action |
| 128 | 2 | 5764 | inn | leave-and-remove request 0x445, post 0x42b |
| 144 | 1 | 5771 | shop | leave-and-remove request 0x445, post 0x42a |
| 160 | 8 | 7911 | gates | flag 0x301 true: leave-and-remove request 0x445, 0x442(1,0), 0x42d; false:`plagatguard` |
| 176 | 0x10 | 5592 | main menu | post 0x41f |
| 192 | 4 | 14733 | school | no room or other click action |

Only shop/inn draw separate hover highlights. Shop starts its enter sound under a latch and can start shopie when `rand()%100>95`. School takes an explicit hover no-op after common selector/guard-direction writes. Other ordinary selectors OR their value into animation flags; inn activates tavern, menu activates stars,0x200 can set BABA and 0x400 sets DERVISH. The complete selected advance has no receiver for 0x800/0x1000. Gate hover controls guard direction through flag 0x301. Tip indices 233/236/234/237/235 identify shop/inn/school/gates/menu. Posted gate navigation acceptance and actual flag value remain Unknown (`R2-ENGINE-241`, High for finite native dispatch).

The remaining EN input slots +0x48/+0x4c/+0x50/+0x58..+0x70 contain no direct selector+0xb4 read that opens a room. Motion delegates to the measured hover; the key body accepts Enter/Escape and forwards other keys. The school and four figure click result is bounded to these input bodies and the finite click table (`R2-ENGINE-241`, amended).

### Town 1 animation

The painter admits one step when unsigned elapsed time exceeds 67 ms, sets a new baseline and performs no multi-frame catch-up loop. Actor scheduling runs every paint. R(n) = floor(rand()*n/32767)%n yields 0..n-1; uniformity and independence are unestablished. Live redraw cadence and sound playback are Unknown (`R2-ENGINE-242`, High for the bounded native mechanisms).

| Element | Frames | Admission and advance |
|---|---|---|
| tavern | 10 sprite frames | inn bit 2; increment, at 10 reset 0/clear bit; old 0 requests Point.wav |
| sign | V00..09 | R(100)>94 sets bit 0x40; increment, at 10 reset 0/clear bit; old 0 requests Flag.wav |
| stars | S00..08 | menu bit 0x10; increment and display below 9; later calls set null and increment persistent blank counter; its tenth such call resets index 0/counter 0 but remains null; old 0 requests Stars.wav |
| shopie | 30 sprite frames | special shop bit 1; increment, at 30 reset 0/clear bit |
| weather vane | F00..07 | R(100)>97 sets bit 0x20; increment, at 8 reset 0/clear bit; old 0 requests Flugel.wav |

These five rules are native episode state, not a promised live frame rate. Entry requests Crowd.wav through its loop-start helper. The lazy step baseline and stars blank counter persist across entry (`R2-ENGINE-242`).

Gates use T00..08 at (180,148). Flag 0x301 false selects T08 immediately. The visual "closed" label for T08 is unmeasured. With the flag true, pointer gate 8 decrements toward 0; away increments toward 8, with endpoint clamps. Direction changes request GateUp/GateDn. Guards use eight frames at (184,158), begin at 7 and add hover direction. Common hover sets 1; gate hover with flag false sets -1. Endpoints 0/7 stop motion and release the guard sound handle; direction transitions request Guard1/Guard2. Actual flag value and accepted navigation remain Unknown (`R2-ENGINE-243`, amended, High for selected native cursor rules).

| Variant | HORSE position | BABA position | DERVISH position |
|---|---|---|---|
| 1 | (104,404) | (216,364) | (224,364) |
| 2 | (104,404) | (308,424) | (324,424) |
| 3 | (256,344) | (384,424) | (392,420) |
| 4 | (448,400) | (580,384) | (592,388) |
| 5 | (140,400) | no variant | no variant |

The loader chooses one HORSE with R(5), one BABA with R(4), and a DERVISH with R(4) unequal to BABA. HORSE A1/A2/A3 have 15 frames each; BABA A1/A2 have 31/32; DERVISH has 30. Initial waits are 2000..3999 ms for BABA/HORSE; later waits are 2000..6999 ms. Strict elapsed>wait chooses an action and resets frame 0. Active frames update their baseline; completion restores idle frame -1, drawn as 0. Elapsed start arms do not test the active flag. DERVISH enters with bit 0x400 and advances modulo 30 (`R2-ENGINE-244`, High).

HORSE requests Horse2 at A1 frame 14 and A2 frames 8/14, Horse1 at A3 frame 1 and Horse2 at A3 frame 14. Horse3 loads but has no request in the complete selected painter. Birds1..9 each have 57 frames; after 1000..2999 ms an inactive group chooses one of three groups of three and a prefix of one/two/three birds. Frames advance once per admitted step. They draw at view origin, then Town_add occludes them. All selected completions clear bit 0x80; bird painting updates the next-wait baseline. One bird requests Birds1, otherwise Birds2. Actual random choices, actor state and audible playback remain Unknown (`R2-ENGINE-244`, High for the finite native populations and schedules).

## Town 1 rooms

Campaign town ID 1 selects the generic inn and shop in both preserved clients. Their selected EN constructors share page or child bases with Kaarg. Shared inn receivers supply buttons, talk and exit. Generic table `L2.00739` uses enter +0x80 `L2.00552`; Kaarg table `L2.00741` replaces it with `L2.00613`. Roster handling belongs to each page's enter body. Per-town center/art methods supply backgrounds, overlays and animation episodes. Detailed geometry below is EN-only. Complete layout equality and differences limited to art remain unproven (`R2-ENGINE-245`, High).

The inn page builds left (0,0,160,480), right-button (480,0,640,238) and center (160,0,480,480) panels. It also attaches the campaign hero panel at (640-panelWidth,0). Its primary graphics are `interface/inn/{LeftStats,LeftPicture,ButtonsArea,CenterArea,manback,ManBackTalk,LUOver,LDOver,RUOver}.bmp` and on/off art for three buttons. The center painter draws background, candle, cauldron, conditional tender, seated actors, edge overlays, a conditional lower-right image and shared children, in that order. The lower-right key and complete actor portrait/stat drawing remain Unknown (`R2-ENGINE-246`, High / Unknown).

The primary inn destinations add to the page origin (`R2-ENGINE-246`, High for selected EN operands):

| Layer | Destination | Condition |
|---|---|---|
| LeftStats / LeftPicture | (0,0) / (0,238) | left painter |
| ButtonsArea | (480,0) | page+0x110 nonzero |
| CenterArea | (160,0) | center painter |
| candle / cauldron | (160,48) / (420,160) | retained frames |
| tender | (240,152) | retained episode |
| seat back then actor | populated seat origin | hire entries before talk entries |
| LUOver / LDOver / RUOver | (160,0) / (160,238) / (464,0) | after seated actors |
| lower-right shared art | (464,238), crop 16x242 | hero-panel state; key Unknown |
| shared children | delegated | after center layers |

Inn button rectangles are (484,44,624,90), (484,91,624,137) and (484,138,624,184). A selected talk actor gives a blank first caption, Talk and Exit. A selected hire entry gives Hire or Fire according to option bit 31. Matching press/release selects the local action; Talk formats actor/topic text, and Exit starts room removal (`R2-ENGINE-246`, High).

The generic inn enter `L2.00552` consumes DLL options in order, with kinds 1/2 assigned to hire entries and other kinds to talk entries. The conditional stage-10 speakers are NPC 207/topic 9, NPC 2108/topic 8 and NPC 517/topic 10. Three rows of six 48x64 seats use corners `(176+48*c,480-64*(r+1),224+48*c,480-64*r)`. Only populated seats participate in the hit receiver. Selection updates the actor index and buttons; talk formats `npc%dtalk%d` from the first matching copied option and calls DLL TalkTo. Complete fresh actor artwork and state-gated roster continuation remain Unknown (`R2-ENGINE-247`, High / Medium / Unknown).

The generic inn loads 10 candle pictures, 21 cauldron pictures, 24 breath pictures and 40 drink pictures at (160,48), (420,160), (240,152) and (240,152). Candle/cauldron advance once after elapsed time exceeds 100 ms; active tender motion advances once after 83 ms. Its odd/even retained delay selects drink/breath; reset delay is 3000..5047 ms from the bounded 15-bit source, without a probability claim. Drink runs forward then reverse; breath returns to idle after forward completion. Only the selected seated actor advances after 125 ms. The wrap helper uses native field+0x8-1, whose relation to loaded-list counts is Unknown; complete visible cadence and audible playback are unobserved (`R2-ENGINE-248`, High / Medium / Unknown).

Inn animation archive keys use `interface/inn/candle/t0000.bmp`..`t0009.bmp`, `cauldron/t0000.bmp`..`t0020.bmp`, `tender/breath/br0001.bmp`..`br0024.bmp` and `tender/drink/dr0001.bmp`..`dr0040.bmp`. Actor art selects `interface/inn/Unit%d/sprites.16a` or HeroMage/HeroFighter. The missing-controlled-actor fallback selects Unit1; its use by every fresh actor remains Unknown (`R2-ENGINE-248`, `R2-ENGINE-247`).

The shop builds inventory (0,303,480,390), hero inventory (0,390,480,480), item information (0,0,164,303), main art (164,0,480,303) and buttons (464,0,640,238). Its selected painter draws ShopFrame at (164,0), ShopMain at (169,8), enabled category layers 0..3 at (353,108), (197,108), (313,20), (201,20) and the keeper at (277,112), then shared children. Support art includes myitem/shopitem, cost and inventory-background families. State-supplied items populate the panels; complete initial stock, item-price text and episode state remain Unknown (`R2-ENGINE-249`, High / Medium / Unknown).

Selected shop archive keys are `graphics.res:interface/ShopFrame.256`, `interface/shopanim/ShopMain.bmp` and `interface/shopanim/01..04/1..11.bmp`; keeper pictures use `movies.res:shopanim/{Pose2-3,Yes,No}/`. Support keys are `interface/{myitem,shopitem}.256`, `costs1..7.bmp`, `costm1..7.bmp` and `backinv{g,b,s}.bmp`. Buttons append the per-town prefix and ShopButton1..4.bmp under interface/. The generic empty prefix remains a Medium inference; Kaarg supplies its own prefix (`R2-ENGINE-249`).

Shop hit rectangles are Undo (494,15,614,67), Buy (483,67,623,113), Sell (483,114,623,160) and Exit (494,160,614,212). These are finite local targets, with distinct hit/art/text rectangles. They do not establish all transaction side effects (`R2-ENGINE-249`, High).

Entering the shop or inn sends square leave-and-remove request 0x445. It invokes leave +0x84, clears the active word, removes the optional child and releases square art and containers. Message 0x44c removes the square from the application stack. The square is not retained there while the room is open. Inn and shop Exit/Escape also remove their room through 0x445/0x44c. A separate conditional 0x42e post calls the town dispatcher and enters the square again: shop Exit/Escape require application+0x404=2, and inn Exit/Escape require +0x404=4 (`R2-ENGINE-250`, partially retracted, High for selected EN instructions).

Re-entry reloads square art and repeats HORSE/BABA/DERVISH variant and position selection; random draws may select the same variants again. It clears per-window animation flags, initializes tavern/sign/stars/shopie/vane frame selections, restores BABA/HORSE idle frames with new initial waits and starts DERVISH at frame 0 under bit 0x400. Gates initialize to cursor 8/T08 with a cleared direction latch. The guard returns to frame 7 with step 0 and a cleared sound-direction latch; the pointer selector resets to -1. Global painter clock state, bird delay and the stars blank counter are not cleared by the selected loader/enter/release (`R2-ENGINE-250`, partially retracted).

The composed visible return remains Medium. Application+0x404 writers and values, the subsequent no-0x42e outcome, complete input/queue admission and visible return timing remain Unknown (`R2-ENGINE-250`, partially retracted).

## Shared town class and overrides

The generic town constructor calls the native page base. Druid and Kaarg
call the generic town constructor, then assign different tables. Each square
has 44 slots from +0x00 through +0xac. Generic inherits 23 native page
slots, overrides 11 and adds 10; druid inherits 32 generic slots and
overrides 12; Kaarg inherits 31 and overrides 13. High applies to table
identity and constructor binding. Data/logic classification is bounded by
selected methods; universal behavior equivalence remains Medium
(`R2-ENGINE-251`, `R2-ENGINE-252`).

| Class | EN / RU table | Direct base | Per-town changed slots |
|---|---|---|---|
| generic, ID 1 | L2.00562 / L2.00700 | native page L2.00651 / L2.00701 | baseline town methods and ten extensions |
| druid, ID 3 | L2.00564 / L2.00702 | generic | 04,2c,54,80,88,8c,90,98,9c,a0,a4,a8 |
| Kaarg, ID 2 | L2.00563 / L2.00703 | generic | same set plus 14 |

| Slot | Role | Druid / Kaarg relation to generic |
|---|---|---|
| 04 | deleting destructor | same wrapper with different destruction target; deeper member ownership differs |
| 14 | tip | druid inherits; Kaarg extends selector-to-text data |
| 2c,54,80,98,9c,a8 | paint, click, entry, hover, schedule, advance | distinct native control flow |
| 88 | sound load | druid changes keys and local state; Kaarg also omits the pre-load virtual release |
| 8c,90,a4 | sound release, ambience start, art release | owned lists/fields and counts differ |
| a0 | art load | both per-town data and distinct loading programs |
| remaining slots | shared native callbacks | exact same-locale targets inherited |

Druid art release `L2.00652` and Kaarg `L2.00653` call a subset of generic
`L2.00654` cleanup primitives. Generic also calls `L2.00655`, `L2.00656`
and `L2.00657` (`R2-ENGINE-252`, Medium for the data/logic label).

The generic gate click tests Scenario variable 0x301. Kaarg advances
reversible gate frames. Druid starts selected people from hover and uses
bug/lizard sprite routes. These are positive town-specific programs.
The selected generic click, shared town input and menu paths do not open
a school; trainer art and a tooltip alone do not establish a school page.
Unexamined global window callbacks remain outside that bounded negative
(`R2-ENGINE-252`).

## Druid square composition

Druid art has 204 keys in its selected square subtree, 82 in its inn,
19 in its shop and 61 in the movie shop subtree. The four subtrees total
366 keys per locale; 364 payloads match. The two differing lizard BMPs
are not selected by the measured square loader, which loads the matching
sprite member. No global unused-resource claim follows
(`R2-ASSET-063`, High).

The background and indexed mask are 640x480. Shop/inn highlights are
152x164 and 152x96. Woman subsets contain 15/20/20 BMPs; man subsets
19/19/20, all numbered from 1. The native loader attempts 1..20 and
stops a subset on failed lookup. Bug and lizard sprite envelopes contain
183 and 42 frames at 640x268 and 100x120. Presence and native selection
are separate contracts (`R2-ASSET-065`, `R2-ENGINE-256`, High).

The active painter adds the view origin to these destinations. Define
q=trunc(s/2) for person subset s=0..2. Each episode indexes its retained
subset vector (`R2-ENGINE-253`, High).

| Order | Layer | Destination | Condition |
|---|---|---|---|
| 1 | townmain | (0,0) | active |
| 2 | shop highlight | (420,224) | hover 1 or woman subset 2 active |
| 3 | woman | (336+4q,244-20q) | active subset; idle hover 1 uses subset 2/file 1 at (340,224), otherwise subset 0/file 1 at (336,244) |
| 4 | inn highlight | (0,184) | hover 2 or man subset 2 active |
| 5 | man | (164-12q,200-16q) | active subset; idle hover 2 uses subset 2/file 1 at (152,184), otherwise subset 0/file 1 at (164,200) |
| 6 | lizard sprite | (0,300) | sequence value minus 1; otherwise frame 0 |
| 7 | bug sprite | (0,212) | nonnegative route/cursor, frame 61*route+cursor |
| 8 | own overlay, then children | delegated | own overlay target is empty |

The native sprite constructor binds slot +0x18 to L2.00641, forwarding
these zero-flag square calls to L2.00642. Selected ROM2 code establishes
the palette-bearing envelopes, word-run geometry and conditional full-table
memory-mode blend. All 450 bug/lizard frames across EN/RU fit that geometry.
Source-RGB composition is Medium: native 16-bit packing, alternative
memory mode, truncation, children and device presentation remain unobserved
(`R2-ASSET-066`, High / Medium).

### Druid hotspots

Both original masks have the same 307200 indexed cells: 262188 zero and
five nonzero regions. No other byte occurs; this includes byte 176 despite
the menu click arm. The inherited sampler uses view-relative mask pixels,
not figure rectangles (`R2-ASSET-064`, `R2-ENGINE-254`, High).

| Byte | Selector | Pixels | Native hover | Click |
|---|---|---|---|---|
| 144 | 1 | 24098 | shop highlight and woman subset 2 | 0x42a, druid shop |
| 128 | 2 | 3397 | inn highlight and man subset 2 | 0x42b, druid inn |
| 160 | 8 | 9952 | latched Dout sound request | 0x442(1,0), then 0x42d |
| 96 | 0x200 | 2144 | no episode start | druidinnkeeper%d dialogue |
| 80 | 0x1000 | 5421 | no episode start | druidshopkeeper%d dialogue |

Changed shop/inn admission starts subset 2 unless it is already running:
set cursor -1, set flag 1/2, then advance immediately to cursor 0 and
reset the family clock. Denter2/Denter1 and Dout have separate latches.
The arm stops and then replays the same family sound: +248 Ddruid1 for
selector 1 at `L2.00662`, +24c Ddruid2 for selector 2 at `L2.00663`.
Scheduler person starts stop/replay these keys at `L2.00664`/`L2.00665`
and `L2.00666`/`L2.00667`, respectively.
The inherited +0x4c also forwards hover outside the painter
(`R2-ENGINE-254`, High).

The two dialogue suffixes use Scenario variable 0x300. They call the
shared dialogue entry. Shop, inn and navigation prepare departure first;
menu 0x41f does not. The complete selected druid loader, painter and
advance contain no gate frame operation. Navigation posting and Dout
remain present. Acceptance after posting is Unknown. The inherited tip
has indices 233/236/237/235 for shop/inn/navigation/menu and no text
for the two keeper selectors (`R2-ENGINE-255`, High / Unknown).

### Druid episodes

One frame advance is admitted after more than 100 ms; the process baseline
is reset without catch-up. Scheduling runs every active paint. R(n) is
the bounded 15-bit helper defined above, without a uniformity claim
(`R2-ENGINE-256`, High).

| Family | Initial / later wait, ms | Start / completion |
|---|---|---|
| woman/man | 2000+R(2000) / 3500+R(5000) | strictly elapsed wait replaces clock/wait; start only when idle, select hover subset 2 or R(2), cursor 0; equality with vector count ends |
| bug | 7000+R(5000) / 10000+R(10000) | choose R(3), retain cursor, set 0x80; each admitted increment ends at 61 |
| lizard | no elapsed wait | when idle, R(15)+1 selects routes 1/2/3 for values 1/2/3, otherwise route 0; cursor 0, sentinel -1 ends |

Bug entry cursor -1 delays its first visible frame until an admitted
advance. Later starts retain cursor 0 after normal completion. The start
arm can reschedule an active bug route without resetting its cursor.
Lizard vectors have 8/11/25/47 values before -1, with repeated holds and
one-based sprite ordinals. The scheduler also requests bird/tree sounds
after separate 2000+R(2000) waits and wolf after more than 60000 ms.
Delivered timing and audible output remain Unknown (`R2-ENGINE-256`).

### Druid campaign availability

The positively identified type-2 town ID 3 record is added on ordinary Leave50.
Nonzero bank775 restoration adds towns 2 and 3 independently of bank768.
Type-2 town ID 3 departure preserves its availability. The measured EnterLocation
setter stores the supplied current pointer without an admission test;
the preceding client's accepted visit remains Unknown. The named fresh
town1-to-town2 path plus the later Leave50 addition is a conditional route,
not an exhaustive visit order (`R2-SESSION-123`, High / Medium).

## Druid rooms and shared children

Campaign selection binds inn pointers +0x104/+0x10c/+0x108 and shop
pointers +0xf8/+0x100/+0xfc for town IDs 1/2/3. The inn page/center
and shop page/center/header variants inherit native room bases. Shared
inn sides, shop lists, actions and inventory bases retain their positively
bound tables. Exact pointer reuse is High; a complete data/logic taxonomy
and whole room composition are Medium (`R2-ENGINE-257`).

| Classes | Callable slots | Data overrides against town 1 | Logic overrides | Equivalent cleanup overrides |
|---|---|---|---|---|
| inn pages 2/3 | 00..8c | 04,78,88,8c | 80,84 | none |
| inn centers 2/3 | 00..84 | 04,80,84 | 2c | none |
| shop pages 2/3 | 00..94 | 04,78,8c,90,94 | 80,88 | none |
| shop center 2 | 00..a8 | 04,14,7c,88 | 2c,54,78,80,84,8c,a4 | 90,94,98,9c,a0 |
| shop center 3 | 00..a8 | 04,14,78,7c,88 | 2c,54,80,84,8c,a4 | 90,94,98,9c,a0 |
| shop headers 2/3 | 00..b8 | 04,b4 | none | none |

These ten derived classes have 396 slots per locale: 326 inherited,
37 data, 23 logic and 10 equivalent overrides. Across both locales that
is 652 inherited, 74 data, 46 logic and 20 equivalent rows
(`R2-ENGINE-257`, Medium for semantic labels).

All ten derived +04 wrappers retain the base's 17 instructions and change
only the destructor call target. The data label covers the wrapper. Derived
destructor bodies were read only for the headers: druid `L2.00673`
forwards to `L2.00674`; other derived room destructor bodies remain
unclassified. Shared widget and shell +04 have the same deleting-destructor
wrapper shape; window +04 semantics remain Unknown. Shop-center +14
changes only the text-index base from 0x3e to 0x116 or 0x112, retaining
the base's four-category rectangle loop (`R2-ENGINE-257`, Medium).

Every remaining derived slot has its direct base's exact target. The five
shop cleanup overrides have identical native instruction sequences after
normalizing local branch targets. Other changed addresses alone do not
prove changed behavior. Shared callback purposes that were not interpreted
remain Unknown (`R2-ENGINE-257`, `R2-ENGINE-260`).

The EN shared character panel is stored at application +0xe0 by constructor
L2.00696, with table L2.00699 and direct base L2.00671. Its prefix has
33 nonzero targets through +0x80, followed by zero +0x84 and the next
table at +0x88. It has 18 reused targets, 12 changed targets and three
extensions against the 30-target window prefix. The zero cell's role and
changed method purposes remain Unknown; no per-town panel class follows
from this binding (`R2-ENGINE-260`, High / Unknown).

### Druid inn

The inn draws `interface/inn_druid/TavernMain.bmp`, shared inn manback,
ManBackTalk, LUOver, LDOver and RUOver bitmaps, druid waterdrop and
taverner art. Let P be the page origin and C=(160,0) the center offset.
Background is P+C, water P+C+(168,144), taverner a1/a2 P+C+(40,128)
and static a30001 P+C+(104,152). Shared portraits, selection/quest
presentation, edges and children follow. The lower-right edge chooses
one of two global bitmaps whose identities remain Unknown
(`R2-ENGINE-258`, High for bounded draw rules).

| Animation | Installed files | Native rule |
|---|---|---|
| waterdrop | d0001..d0010, 116x96 | >100 ms; (cursor+1) mod (count-1), steady files 1..9 |
| taverner a1 | a10001..a10040, 128x196 | state 1, >100 ms, one-shot cursor through 39 |
| taverner a2 | a20001..a20030, 128x196 | state 2, >100 ms, one-shot cursor through 29 |
| static taverner | a30001, 32x24 | states 3/4 at a separate destination |

The initial wait is rand()/16+3200, or 3200..5247 ms on the selected
15-bit source. Idle expiry chooses (wait&3)+1, mapping 4 to 3. A1/a2
completion returns to idle. State 3 changes to 4 and stores a different
wait; the next admitted >100-ms state-4 tick returns to idle. Cursor
reset preserves the cached bitmap, so a restarted episode need not first
display file 1. Water uses a count-minus-one helper despite ten loaded
names. Tooltip, mouse and click methods are inherited
(`R2-ENGINE-258`, `R2-ASSET-065`, High).

### Druid shop

The center loads ShopFrame.256, ShopMain.bmp, elven.bmp and
hili_armor/magic/potion/elven.bmp under `interface/shop_druid`, plus
`movies.res:shop_druid/a10001.bmp` and a1/a2 keeper series. The header
loads druid ShopInv and four ShopArrow bitmaps. Shared inventory/action
widgets remain native classes (`R2-ENGINE-259`, `R2-ASSET-065`).

With C=(164,0) relative to page P, center order/destinations are frame
P+C+(0,0), background +(5,8), admitted elven +(93,172), selected armor
+(5,112), magic +(5,52) or potion +(121,36), keeper +(197,92), then
children. Campaign query 0x302 admits elven paint/click/selection only
on nonzero return; outside campaign the predicate admits it. Tooltip
still scans four rectangles. Original ScenarioGetVar reads bank770 for
0x302; ordinary Leave70 stores it to 1. This enables the gate without
an intervening writer. Other producers and restoration remain Unknown
(`R2-ENGINE-259`, High for these local rules).

Keeper ticks require at least 100 ms. Each tick samples a fresh
5000+1000*(rand()%5) idle threshold, then chooses episode flag 0x10/0x20
when idle and admitted. Update increments counter modulo 30 and uses
counter+1 as the a1/a2 filename. Zero clears the episode; default is
a10001. An initialized episode starts its updated sequence at file 2,
passes file 30 and returns to default file 1. Kaarg instead uses modulus
20/default a10000 and additional fire transitions. Shared shell/button
art, pending dialogue, audible timing and live pixels remain Unknown
(`R2-ENGINE-259`).
