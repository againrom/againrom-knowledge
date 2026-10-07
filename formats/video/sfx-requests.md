# Shared SFX request recipes

The 112 recovered direct SFX requests have distinct sample-source and playback
tuples in 74 original owners. The table publishes the reviewed selector or
receiver-field recipe for every terminal site. High applies to this finite
population and its instruction-bound metadata. Medium applies to closure over
all possible routes and runtime-built receiver/path associations.
— VIDEO-SFX-085

The input is the identical installed EN/RU ROM1 executable, 1,977,344 bytes,
SHA256 digest `942e9b72610eeba2f3d74930ee47cbc4476b85a942c14348f874ec7b7d367d03`.
The repaired Ghidra 12.1.2 direct-call census and the raw executable
Capstone 5.0.7 `E8 rel32` scan agree on the 112 service sites. This is static
request evidence; no native audio output was observed. The finite direct-call
and stored literal-word searches do not close computed pointers, copied
pointer-bearing structures or runtime-built paths. — VIDEO-SFX-085

## Recipe fields

`Original site` and `Original owner` are opaque executable navigation
identifiers. They are not semantic event names, selector values or runtime
types. `Source kind` restates the syntax of the reviewed source rule; the
rule itself retains its original receiver-field association and selector.
Original field offsets describe that association, not a required layout for
another implementation. — VIDEO-SFX-085

`setting` is the effects or speech volume setting named by `Volume category`.
`D+setting` adds that setting to pre-setting attenuation D. The distance-derived
priority is `trunc((10000-abs(D))/100)&255`, calculated before adding the volume
setting. With D from -10000 through 0 it yields 0..100. This conditional
arithmetic does not establish overflow or malformed spatial behavior.
`ambient_term+setting` and `positional` refer to the separate arithmetic and
input contracts below. Zero pan is explicit. All recovered frequency
arguments are zero, so these requests skip the service's frequency-change call.
— VIDEO-SFX-085

Repeat zero requests Play flag 0; nonzero requests flag 1. The service clamps
volume only below -10000 and does not check its volume, pan, frequency or Play
results. A recipe records a request rather than successful audible playback.
The 79 priority-128, 17 priority-220, three priority-100 and 13 distance-derived
sites include seven repeat-1 and 105 repeat-0 requests, 100 effects and 12
speech requests, and 16 positional and 96 zero-pan requests.
— VIDEO-SFX-085, VIDEO-SFX-086

## Positional input and arithmetic

The 13 distance-priority recipes use drawable fine world x/y, at 256 fine
units per cell, and their associated map view. The listener centre is
`originX*256+widthCells*128, originY*256+heightCells*128`; the original uses
32-bit shifts and adds before converting each source-minus-centre delta to
floating point. These are world/view relationships; projected screen x/y and
drawable z are not helper inputs. View origin is in cells, and the viewport
span comes from its rectangle dimensions divided by 32 toward zero.
SESS-VIEW-028/029 supply that frame mapping. — VIDEO-SFX-088, VIDEO-SFX-089

For ordinary finite inputs with valid width and representable integer
operations, the instruction-bound algebra is:

```
dx = sourceFineX - listenerFineX
dy = sourceFineY - listenerFineY
pan = clamp(trunc(dx * 4000 / (widthCells * 256)), -10000, 10000)
D = trunc(max(-10000, -(exp(sqrt(dx*dx + dy*dy) / 256 / 8) - 1) * 100))
```

The positional squared deltas are floating-point products. Pan truncates
before clamping; D lower-clamps before truncation. There is no additional D
upper clamp. The helper returns stored D, so a zero return at the centre is
an attenuation result. Priority still derives from D before adding effects
or speech setting under the terminal recipe. — VIDEO-SFX-088, VIDEO-SFX-085

The conversion temporarily selects rounding toward zero, stores a signed
qword and returns its low dword, then restores the prior control word. The
formula states the arithmetic for ordinary inputs; it is not a native FPU
bit-precision specification. Native x87 entry precision/control, libm edges,
invalid dimensions, overflow and out-of-range conversions remain Unknown.
The former log10 arithmetic label of ANIM-SND-022 is partially retracted:
the reached original path uses exp. — VIDEO-SFX-088, ANIM-SND-022

Ten original-instruction controls with explicit CW027f and artificial view
origin 8/span 15 give positional D=-171 and pan=2133 at eight horizontal
cells; the vertical control keeps D=-171 and pan=0. These are synthetic
execution results. Native playback and precision were not observed.
— VIDEO-SFX-088

The following field recipes identify the input for each concrete terminal.
Original sites remain opaque navigation identifiers. Message arm labels A/B/C
are bounded aliases; exact live player trigger coverage is Unknown.
— VIDEO-SFX-089

| Original terminal site | Semantic recipe | Source fine x/y | Associated view |
|---|---|---|---|
| `L13021` | Message picture A, registry[500+picture] | Message byte cells +d/+e, each `(cell<<8)+128` on the new drawable | Dispatching map view |
| `L13022` | Message picture B, registry[500+picture] | Message byte cells +a/+b, each `(cell<<8)+128` on the new drawable | Dispatching map view |
| `L13023` | Message object picture C, registry[500+object.picture] | Message packed word +e low/high byte cells, each centred by 128 | Dispatching map view |
| `L03156` | Class swing Sound[0] | Current drawable x/y | Drawable's map view |
| `L13024` | Unit action spell | Current drawable x/y | Drawable's map view |
| `L02732` | Hurt k=0, class Sound[1] | Current drawable x/y | Drawable's map view |
| `L02700` | Hurt/fall bank or class source | Current drawable x/y | Drawable's map view |
| `L13025` | Command/defend bank reply | Current drawable x/y | Drawable's map view |
| `L13026` | Selection bank reply | Current drawable x/y | Drawable's map view |
| `L13027` | Defend bank reply | Current drawable x/y | Drawable's map view |
| `L13028` | Retreat bank reply | Current drawable x/y | Drawable's map view |
| `L13029` | Idle bank reply | Current drawable x/y | Drawable's map view |
| `L13030` | Picture-51 phase-8 cue | Newly updated drawable x/y after this step | Drawable's map view |

The raw direct-call population has 13 helper calls, each associated above.
Computed/table helper calls, copied pointer-bearing objects, intervening
position writes and per-caller receiver lifetime are not closed by that census.
— VIDEO-SFX-089

## Ambient input and arithmetic

Ambient uses cell x/y and integer listener
`originX+floor(widthCells/2),originY+floor(heightCells/2)` for ordinary
nonnegative spans. Origin 8/span 15 centres ambient at cell 15; positional
centre is 15.5. It scans inclusive x from `max(originX-widthCells,8)` through
`min(originX+2*widthCells,mapWidth-8)`, and corresponding y bounds.
— VIDEO-SFX-090

River cells have terrain kind `((terrainWord&0x1fff)>>6)` in 8..11.
Firewall cells have flag 8 in their packed-cell lookup entry. Due object
placements use their nonzero byte minus one to select a definition: the
named definition field >=0 adds a bird and ==-2 adds a crow; other negatives
add neither. These source rules establish semantic membership without deriving
metadata from a sample filename. — VIDEO-SFX-090

For valid dimensions and ordinary representable arithmetic, each matched cell
uses deltas from the integer ambient listener:

```
rawPan = clamp(trunc(dx*2000/widthCells), -2000, 2000)
radius = trunc(sqrt(int32(dx*dx + dy*dy)))
q = trunc(radius/8)
A = max(0, trunc(10000 - (exp(q)-1)*100))
weightedPan = trunc(int32(rawPan*A)/10000)
groupPan = trunc(sum(weightedPan)/count)
riverOrFirewallTerm = max(A)-10000
birdOrCrowTerm = min(trunc(count*1000/30),1000)-2000
```

Ambient squared deltas are original 32-bit integer products before sqrt.
River and firewall have separate count, sum and maximum; bird and crow share
count and sum. Counts include matched zero-weight sources. Pan divides by
source count rather than the sum of weights. Ambient does not reuse positional
D. Explicit-CW027f original-instruction controls give term 0 at seven cells
and -172 at eight; the positional eight-cell D is -171. Native bit precision,
overflow and malformed source tables remain Unknown. — VIDEO-SFX-090

The three terminal consumers remain river and firewall repeat-1 requests,
and the bird/crow repeat-0 request. Each uses group pan and its term plus
effects setting, priority 220 and frequency zero. — VIDEO-SFX-090, VIDEO-SFX-085

## Ambient update and stop boundary

Origin movement against its cached x/y or unsigned deadline < clock permits
the scan. Equality is not due. Only movement updates the origin cache;
width/height changes alone do not satisfy that predicate. River/firewall
scanning uses either condition; object/bird/crow scanning requires due.
Nonempty bird/crow count stores `clock+trunc(rand()/2)+10000` after the
request, while zero count leaves the deadline. Native cadence and clock-wrap
behavior remain Unknown. — VIDEO-SFX-091

River/firewall with no playing channel and a nonempty matched population
request their loops. An existing channel with a nonempty population updates
its non-null buffer volume to `clamp(term+effectsSetting,-10000,0)` and pan.
An existing channel with an empty population stops its non-null buffer and
rewinds to position zero. These local paths do not clear the channel or
destroy its sample. Future caller retries, loop restarts and device success
remain Unknown. — VIDEO-SFX-091

## Terminal request tuples

Every row is the reviewed request tuple under VIDEO-SFX-085. Filename family
does not determine volume category: shop `start.wav` uses speech, while school
Command1..3 and Fight1/Fight2 use effects. Hurt/fall and bank replies retain
the ANIM-094 and ANIM-119 selectors shown in their source rules. No separate
fall filename or timing rule is added. — VIDEO-SFX-085

| Original site | Original owner | Source kind | Reviewed sample source | Volume category | Volume term | Pan | Repeat | Priority u8 | Frequency |
|---|---|---|---|---|---|---|---|---|---|
| `L13031` | `R0333` | `registry selector` | `registry[8]: click04` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13032` | `R0509` | `registry selector` | `registry[15]: message` | `effects` | `setting` | `0` | `0` | `100` | `0` |
| `L13033` | `R0509` | `registry selector` | `registry[15]: message` | `effects` | `setting` | `0` | `0` | `100` | `0` |
| `L13034` | `R0509` | `registry selector` | `registry[15]: message` | `effects` | `setting` | `0` | `0` | `100` | `0` |
| `L13021` | `R0509` | `registry selector` | `registry[500+picture]` | `effects` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L13022` | `R0509` | `registry selector` | `registry[500+picture]` | `effects` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L13023` | `R0509` | `registry selector` | `registry[500+object.picture]` | `effects` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L13035` | `R0382` | `registry selector` | `registry[7]: ibook` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13036` | `R0381` | `registry selector` | `registry[7]: ibook` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13037` | `R0383` | `registry selector` | `registry[7]: ibook` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13038` | `R0335` | `registry selector` | `registry[7]: ibook` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13039` | `R0484` | `registry selector` | `registry[50]: ambient/river` | `effects` | `ambient_term+setting` | `positional` | `1` | `220` | `0` |
| `L13040` | `R0484` | `registry selector` | `registry[90]: magic/firewall` | `effects` | `ambient_term+setting` | `positional` | `1` | `220` | `0` |
| `L13041` | `R0484` | `registry selector` | `registry[60..62] or registry[70]: ambient/bird1..3 or crow; slot62 absent` | `effects` | `ambient_term+setting` | `positional` | `0` | `220` | `0` |
| `L13042` | `R0830` | `owner sample field` | `parent+0x84: SFX/ChrGen/+_-.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13043` | `R0831` | `owner sample field` | `parent+0x84: SFX/ChrGen/+_-.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13044` | `R2113` | `owner sample field` | `parent+0x88: SFX/Click_Ok.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13045` | `R2113` | `owner sample field` | `parent+0x8c: SFX/Sbros.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13046` | `R2113` | `owner sample field` | `parent+0x90: SFX/Back.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13047` | `R1552` | `owner sample field` | `this+0x120+4*skill: SFX/ChrGen/Skill/F{Sword,Axe,Club,Pike,Bow}.wav or M{Fire,Water,Air,Earth,Astral}.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13048` | `R2114` | `owner sample field` | `parent+0x8c: SFX/Rename.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13049` | `R2114` | `owner sample field` | `parent+0x90: SFX/Delete.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13050` | `R2115` | `owner sample field` | `parent+0x94: SFX/Undo.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13051` | `R1878` | `owner sample field` | `this+0x194: SFX/ChrGen/Level1.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13052` | `R1878` | `owner sample field` | `this+0x198: SFX/ChrGen/Level2.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13053` | `R1878` | `owner sample field` | `this+0x19c: SFX/ChrGen/Level3.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13054` | `R1878` | `owner sample field` | `this+0x1a0: SFX/ChrGen/Char.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13055` | `R1878` | `owner sample field` | `this+0x1a0: SFX/ChrGen/Char.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13056` | `R1878` | `owner sample field` | `this+0x1a0: SFX/ChrGen/Char.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13057` | `R2116` | `owner sample field` | `this+0x1a4: SFX/ChrGen/Ok.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13058` | `R2117` | `owner sample field` | `this+0x1a8: SFX/ChrGen/Ok.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13059` | `R2118` | `owner sample field` | `this+0x1ac+4*index: SFX/Letter1..3.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L12955` | `L03131` | `registry selector` | `registry[100]: units/sword` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13060` | `L03131` | `voice bank field` | `mf_merc bank+0x8: sfx/mf_merc/select2.wav` | `speech` | `setting` | `0` | `0` | `220` | `0` |
| `L13061` | `R0814` | `owner sample field` | `this+0xc8: SFX/ChrGen/Ok.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L03156` | `R0573` | `registry selector` | `registry[class.Sound[0]]` | `effects` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L13024` | `L13062` | `registry selector` | `registry[500+unit.actionspell]` | `effects` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L02732` | `R0575` | `registry selector` | `registry[class.Sound[1]]; hurt k=0` | `effects` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L02700` | `R0575` | `voice bank or registry selector` | `bank+0x24/+0x28/+0x2c (easy/hard/die.wav) for k=1/2/3, or registry[class.Sound[k+1]]; ANIM-094` | `speech` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L13025` | `R0578` | `voice bank field` | `bank+0xc/+0x10/+0x14 or +0x1c: command1..3 or defend.wav; ANIM-119` | `speech` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L13026` | `R0579` | `voice bank field` | `bank+0x4/+0x8: select1/2.wav; ANIM-119` | `speech` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L13027` | `R0580` | `voice bank field` | `bank+0x1c: defend.wav; ANIM-119` | `speech` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L13028` | `R0581` | `voice bank field` | `bank+0x18: retreat.wav; ANIM-119` | `speech` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L13029` | `R0582` | `voice bank field` | `bank+0x20: idle.wav; ANIM-119` | `speech` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L13030` | `R0558` | `registry selector` | `registry[551]: picture51 phase8 selector` | `effects` | `D+setting` | `positional` | `0` | `trunc((10000-abs(D))/100)&255` | `0` |
| `L13063` | `R1530` | `owner sample field` | `this+0x17c/+0x180: SFX/Point1/2.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13064` | `R1946` | `owner sample field` | `this+0x174: SFX/ScrollUp.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13065` | `R1946` | `owner sample field` | `this+0x178: SFX/ScrollDn.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13066` | `R0701` | `registry selector` | `registry[14]: mcomplet` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13067` | `R0701` | `registry selector` | `registry[16]: mfailed` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13068` | `L13069` | `owner sample field` | `parent+0xa8: SFX/Town/Inn/Helper.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13070` | `L13071` | `owner sample field` | `parent+0xa8: SFX/Town/Inn/Helper.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13072` | `R1807` | `owner sample field` | `parent+0xa8: SFX/Town/Inn/Helper.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09952` | `R1811` | `owner sample field` | `parent+0xb0: SFX/Out.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09953` | `R0703` | `owner sample field` | `parent+0xb4: SFX/Talk.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09954` | `R1902` | `owner sample field` | `inn+0x8c: SFX/Town/Inn/steam.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09955` | `R1902` | `owner sample field` | `inn+0x94: SFX/Town/Inn/chair.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09956` | `R1902` | `owner sample field` | `inn+0xac: SFX/Town/Shop/Breath.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09957` | `R1902` | `owner sample field` | `inn+0x84: SFX/Town/Inn/drink.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09958` | `R1902` | `owner sample field` | `inn+0x88: SFX/Town/Inn/glotok.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09959` | `R1935` | `owner sample field` | `inn+0x90: SFX/Town/Inn/water.wav` | `effects` | `setting` | `0` | `1` | `128` | `0` |
| `L09960` | `R1412` | `owner sample field` | `inn+0xa4: SFX/Town/Inn/enter.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09961` | `R1441` | `owner sample field` | `inn+0x98: SFX/Add.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09962` | `R1441` | `owner sample field` | `inn+0xa0: SFX/Town/Shop/nofit.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09963` | `R1440` | `owner sample field` | `inn+0x9c: SFX/NoAdd.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09964` | `R0702` | `owner sample field` | `parent+0xb4: SFX/Talk.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13073` | `R2119` | `owner sample field` | `this+0xc0: SFX/ChrGen/Ok.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13074` | `R1205` | `registry selector` | `registry[1]: click00` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13075` | `R0332` | `registry selector` | `registry[1]: click00` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13076` | `R0332` | `registry selector` | `registry[1]: click00` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13077` | `R0332` | `registry selector` | `registry[1]: click00` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L13078` | `L09946` | `owner sample field` | `owner=this+0x20ac; owner+0x90: SFX/Town/Shop/nofit.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13079` | `R1749` | `owner sample field` | `this+0x20bc: SFX/Scroll.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13080` | `R1750` | `owner sample field` | `this+0x20bc: SFX/Scroll.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13081` | `L10016` | `owner sample field` | `this+0x20bc: SFX/Scroll.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L10018` | `L10005` | `owner sample field` | `this+0x20bc: SFX/Scroll.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L09944` | `L09858` | `owner field with computed path` | `this+0x20b0: speech/shop/effects/%.2d.wav, books/%.2d.wav or s%.2di%.2dp%d.wav; SHOP-109` | `speech` | `setting` | `0` | `0` | `128` | `0` |
| `L09945` | `L09858` | `owner sample field` | `this+0x20b4: SFX/Put_On.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13082` | `L09903` | `owner sample field` | `this+0x20b8: SFX/Put_Off.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13083` | `R0705` | `owner sample field` | `shop+0xb0: SFX/Town/Shop/start.wav` | `speech` | `setting` | `0` | `0` | `128` | `0` |
| `L13084` | `L09659` | `owner sample field` | `shop+0xb0: SFX/Town/Shop/start.wav` | `speech` | `setting` | `0` | `0` | `128` | `0` |
| `L13085` | `R1701` | `argument sample field` | `*argument1 sample field; ordinary wrapper; ten shop caller fields and literal paths below` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13086` | `R2104` | `argument sample field` | `*argument1 sample field; loop wrapper; L13087 passes shop+0xbc: SFX/Town/Shop/InShop.wav` | `effects` | `setting` | `0` | `1` | `128` | `0` |
| `L13088` | `R1916` | `owner sample field` | `town+0x80/+0x84: SFX/Town/Birds1/2.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13089` | `R1489` | `owner sample field` | `town+0xa8: SFX/Town/Horse1.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13090` | `R1489` | `owner sample field` | `town+0xa0: SFX/Town/Horse2.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13091` | `R1489` | `owner sample field` | `town+0xa0: SFX/Town/Horse2.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13092` | `R1927` | `owner sample field` | `town+0x7c: SFX/Town/Guard2.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13093` | `R1927` | `owner sample field` | `town+0x7c: SFX/Town/Guard1.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L11557` | `R1922` | `owner sample field` | `town+0x90: SFX/Town/Point.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13094` | `R1928` | `owner sample field` | `town+0x8c: SFX/Town/Flag.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13095` | `R1926` | `owner sample field` | `town+0x78: SFX/Town/GateUp.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13096` | `R1926` | `owner sample field` | `town+0x78: SFX/Town/GateDn.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13097` | `R1918` | `owner sample field` | `town+0x9c: SFX/Town/Stars.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13098` | `R1929` | `owner sample field` | `town+0x88: SFX/Town/Flugel.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L11556` | `L11546` | `owner sample field` | `town+0x94: SFX/Town/Shop/enter.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L11561` | `L11546` | `owner sample field` | `town+0x98: SFX/Town/School/Point.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13099` | `L11529` | `owner sample field` | `town+0x74: SFX/Town/Crowd.wav` | `effects` | `setting` | `0` | `1` | `128` | `0` |
| `L13100` | `R0706` | `owner sample field` | `school+0xc0: SFX/Town/School/Enter.wav; L13101 loads this through R2120` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L07697` | `R1477` | `owner sample field` | `school+0xa0/+0xa8: SFX/ChrGen/Skill/MAir/MAstral.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13102` | `R1477` | `owner sample field` | `school+0x94/+0x98: SFX/ChrGen/Skill/FBow/MFire.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13103` | `R1971` | `owner sample field` | `school+0xb4: SFX/Town/School/Command3.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13104` | `R1971` | `owner sample field` | `school+0xac: SFX/Town/School/Command1.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13105` | `R1971` | `owner sample field` | `school+0xb0: SFX/Town/School/Command2.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13106` | `R1971` | `owner sample field` | `school+0xb8: SFX/Town/School/Fight1.wav` | `effects` | `setting` | `0` | `1` | `128` | `0` |
| `L13107` | `R1971` | `owner sample field` | `school+0xb0: SFX/Town/School/Command2.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13108` | `R1971` | `owner sample field` | `school+0xbc: SFX/Town/School/Fight2.wav` | `effects` | `setting` | `0` | `1` | `128` | `0` |
| `L13109` | `L13110` | `owner sample field` | `school+0x80: SFX/Town/School/Rotate.wav at L13111/L13112 or teach.wav at L12289; L13113 constructs argument path` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13114` | `L12172` | `owner sample field` | `parent+0xc4: SFX/Out.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L12288` | `R1548` | `owner field with computed path` | `this+0xa8: Speech/training/npc33s%dl%d.wav or npc34s%dl%d.wav` | `speech` | `setting` | `0` | `0` | `128` | `0` |
| `L12956` | `L03645` | `registry selector` | `registry[2]: click01; preceding slot4/slot1 stores fall through to slot2` | `effects` | `setting` | `0` | `0` | `220` | `0` |
| `L09967` | `R0711` | `owner field with computed path` | `panel+0x70: speech/directory/(sound override or npc%02de%sp%d or %sp%d).wav; DLG-SOUND-028` | `speech` | `setting` | `0` | `0` | `128` | `0` |

## Shop wrapper recipes

The ordinary wrapper has ten recovered encoded callers; the loop wrapper has
one. Each caller passes the sample field and literal path below. The remaining
tuple fields come from its terminal request above. Original caller, owner,
wrapper and terminal addresses remain navigation identifiers. These rows do
not add another eleven terminal sites to the 112-site census.
— VIDEO-SFX-085

| Original caller | Original owner | Original wrapper | Terminal site | Sample field | Literal path | Volume category | Volume term | Pan | Repeat | Priority u8 | Frequency |
|---|---|---|---|---|---|---|---|---|---|---|---|
| `L13115` | `R0705` | `R1701` | `L13085` | `shop+0xac` | `SFX/Town/Shop/enter.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13087` | `R0705` | `R2104` | `L13086` | `shop+0xbc` | `SFX/Town/Shop/InShop.wav` | `effects` | `setting` | `0` | `1` | `128` | `0` |
| `L13116` | `R1737` | `R1701` | `L13085` | `shop+0xa8` | `SFX/Town/sell.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13117` | `R1736` | `R1701` | `L13085` | `shop+0x90` | `SFX/Town/Shop/nofit.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13118` | `R1736` | `R1701` | `L13085` | `shop+0xa4` | `SFX/Town/buy.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13119` | `L09659` | `R1701` | `L13085` | `shop+0xb4` | `SFX/Town/Shop/Povorot1.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13120` | `L09659` | `R1701` | `L13085` | `shop+0x94` | `SFX/Town/Shop/step1.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13121` | `L09659` | `R1701` | `L13085` | `shop+0xb8` | `SFX/Town/Shop/Povorot2.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13122` | `R1523` | `R1701` | `L13085` | `parent+0xa0` | `SFX/Town/Shop/depart.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13123` | `L12179` | `R1701` | `L13085` | `parent+0xc4` | `SFX/Undo.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |
| `L13124` | `L12179` | `R1701` | `L13085` | `parent+0xc0` | `SFX/Out.wav` | `effects` | `setting` | `0` | `0` | `128` | `0` |

## Source and ownership boundaries

Ambient registry slot 62 is absent in the reused installed registry. Horse3
is constructed, but no recovered direct horse-family request chooses it.
This bounded result does not establish global silence. A computed selector
or path is a source rule, not a guarantee that every possible selected member
exists. School entry binds `SFX/Town/School/Enter.wav`; the shared school
action recipe uses `Rotate.wav` at its two recorded callers and `teach.wav`
at its recorded teaching caller. Their original caller identifiers remain in
the terminal source description. — VIDEO-SFX-085

The named startup count is 16 for both the global channel array and each
successfully loaded sample's buffer array. A request requires a device,
successful lazy load and a nonplaying duplicate of its own sample before
global channel competition. A full sample array returns before the allocator,
regardless of priority. Checked file/open or initial buffer-creation failure
ends the request. An admitted sample takes the first empty or nonplaying
global channel, or the first lowest-priority busy channel strictly below its
priority. Equality cannot evict; no eligible channel ends the request.
The inspected refusal branches store no pending request. — VIDEO-SFX-086

Priority eviction stops and rewinds the victim and clears its channel buffer,
sample and priority fields, while retaining the victim sample's buffers.
With a device and a loaded sample, sample destruction stops every playing
duplicate and clears each matching global channel's buffer pointer; channel
sample/priority fields remain until the allocator normalizes the empty
channel. Buffer release frees the sample array and clears its array/loaded
fields. Replacement and town cleanup first stop/rewind one returned playing
channel, then destroy the whole sample. Eviction itself supplies no loop
restart; later requests depend on the caller. — VIDEO-SFX-086

The two inspected global cleanup slices Stop every non-null channel buffer
and clear all three channel fields without rewinding, then unload registry
samples 14 and 16. Their local instruction slices do not establish that every
branch or mission transition reaches them. — VIDEO-SFX-086

Music owns its streaming buffer and repeat flag under the music enable and
volume settings. Movie audio belongs to decoder sound-backend setup and the
game-side open/close controls. These inspected named routes bypass the shared
SFX sample request and priority allocator. Separate playback ownership does
not establish independent physical device resources. — VIDEO-SFX-087

## Unknowns

The terminal table and positional input association record source-field
recipes. Exact live semantic trigger coverage, per-caller sample lifetime,
retry policy and restart cadence outside the named ambient paths remain
Unknown. Exact live bank/registry
selection, dynamic sample existence, routes through computed pointers or
copied pointer-bearing structures, native audibility and latency, device
failures and malformed spatial arithmetic remain Unknown. — VIDEO-SFX-085

Native playback success and saturation behavior, allocation/duplicate-creation
failures, driver HRESULTs, teardown after device loss and arbitrary malformed
state remain Unknown. Unchecked allocation or duplicate creation does not
guarantee graceful refusal. Future caller retries and unobserved loop restart
cadence remain separate from the shared service's refusal/eviction branches.
— VIDEO-SFX-086

Native music/movie/SFX coexistence, physical device competition, audible
transitions and computed bypass routes outside the inspected named bodies
remain Unknown. — VIDEO-SFX-087
