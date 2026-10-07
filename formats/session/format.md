<a id="session--the-games-own-lifecycle--specification-partial"></a>

# Session state, clocks and command dispatch

A session owns phase/screen state, two simulation counters and the command
pool. Map load and mission completion cross separate simulation and frontend
paths. — SESS-OBJ-001, SESS-CLOCK-005, SESS-CMD-008, SESS-HERO-014.
SESS-HERO-014's default-name array clause is superseded by SESS-DEFNAME-032
and SESS-NAMEGATE-033; its command and suffix clauses stand.

The [trigger runtime](../trigger/format.md) defines script outcomes, and
[dialogue](../dialogue/format.md) defines the resulting panels. The network
client's complete clock and pre-create button rectangles remain Unknown.

## Defeat mode and stepping

Frontend phase (`+0x6bc`), local authority (`+0x6b8`), actual server mode (`+0x0c`) and
participant entry/join bytes are distinct. New/load campaign allocates with argument 2;
phase-3 authority starts pass 0, and phase-1 host start passes 1. Initialization writes
`server+0x0c = (arg < 2)`, but successful ALM load overwrites it from type-0 payload
`+0x70 > 1`. These stores reload the **global server pointer**: the map sub-object's
`+0x44` base adjustment does not apply to them (`SESS-MAP-010` (its own blanket "every other
displacement is 0x44 low" is retracted), `SESS-DEFEAT-064`).
All 28 maps per campaign root load zero. RU loose Horror.alm also loads zero; all 10 EN
loose maps and the other 5 RU loose maps load one. Save deserialization does not restore
this mode dword. Constructor-only and loose-map-equals-network rules are invalid.

With an active human participant, automatic reporter repair requires actual mode nonzero,
latch at least 2, entry-active byte `player+0x3f != 0` and signed primary HP below -53.
Joined byte `+0x3e` instead gates the primary-fall failure packet. Session opcode 5 clears
entry-active without changing the mode (`MISSION-DEFEAT-045`, `HERO-DEFEAT-136`).

Campaign idle skips stepping when phase is 2 and screen mask intersects `0x4008`.
Outcome panels set bit 8 without clearing the server run dword. Network authority/client
branches occur outside that campaign pause condition. The ordinary stepper's four callers
are paced idle, unpaced idle, phase-3 idle and bootstrap. A separate thread loop can report
too, but its launcher has no discovered incoming direct/stored reference. The optional
five-full-tick watchdog has independent `server+0x148/+0x13c` flags and calls placement,
not XP repair; nondefault runtime activation is Unknown (`SESS-DEFEAT-065`).

## Mission input ownership

The root input object is `campaign+0xcc`. Its children are the map `+d0`, then the fixed-width
right column `+d4`. The column's children are minimap `+d8`, command panel `+dc`, character panel
`+e0`, then the resolution filler `+e4`. Mouse dispatch goes to the capture object when one is
set; otherwise to the first child containing the point, with no bubbling after that hit. Keyboard
dispatch goes first to the focused child. This hierarchy, not independent hit tests, decides
precedence: a column click cannot reach the map; text entry suppresses map shortcuts; an Esc
popup handles Esc before the mission root (SESS-INPUT-037). Its former
inventory-parent clause is partially retracted; the measured inventory-open
parent is the map view (MENU-088).

Only map right down sets the engine capture object, for a right-button pan; right up clears it.
When marquee global `[L00210]` is zero, left down instead calls `R2110`, which reaches
OS `ClipCursor` with a computed map rectangle and establishes the marquee; another down while it
is nonzero is a consumed no-op. Left up closes the marquee only while the global is nonzero and
then calls `R0111`; its raw body adds `0x6c0` to the session pointer and passes that stored
rectangle to `ClipCursor`. It does **not** pass NULL, and neither left edge calls the engine
capture helper. An unmatched left up skips the close/hit/selection path, but the selected-item
handling located after that gate still runs. The command panel and minimap do neither. The image has no mission
`WM_CANCELMODE` or `WM_CAPTURECHANGED` handler, so an implementation must separately choose safe
handling for interruption before the normal left-up rectangle replacement and for external loss
of the right-pan capture, and record both as authored rather than as ROM1 behaviour.

The measured inventory-open route appends campaign `+e8`, the inventory grid, to the map view (MENU-088). The former character-panel parent clause of SESS-INPUT-037 is partially retracted; alternate attachment routes remain Unknown. Its physical left-up hit-tests through vslot
`+0x88`; with live grid content and a selected item at `session+0x3cc`, it passes that result to
vslot `+0xa4` and emits transfer `0x22/0x32` to destination code 2. A `-1` hit merges or appends
rather than cancelling. Physical double-click also maps the point to a cell. A regular item can
enter the character-panel Cast/transfer vmethod; a sentinel gold cell (`item+6 == 0xffff`) with a positive purse attaches the always-created Drop Gold dialog at
`campaign+0x10c`. That 296x168 modal takes focus. Its button left-up or Enter action emits item
opcode `0x23` at the current cell after subtracting the resolver-returned amount; cancel-button left-up or
Esc closes without an order. While attached, the root's focus/last-child rules give it precedence
over map shortcuts (SESS-INPUT-037).

The same grid's left-down handler scrolls one position through viewport-dependent left and right
edge strips. Its mouse-move handler begins a drag only with `MK_LBUTTON`, live content, no existing
selected item and a live-cell hit. Shift clear requests one regular item or 1000 gold; Shift set
requests the source stack's quantity. A successful gold selection debits the purse when the drag
is built. No command is emitted until left-up. Right down, up and double-click return consumed
without changing selection or issuing a command.

Drop Gold opening sets edit text to `0`; the later selection helper separately
tests existing selection and current-actor ownership. Its generic byte edit and
floating parse do not establish safe numeric validation (MENU-090, MENU-092).
MENU-088's explicit Backspace selection-reset clause is partially retracted;
the local branch changes text/caret without endpoint stores (MENU-091).
The modal's own Enter requests action, while a focused Cancel button consumes
Enter as cancel. Successful action changes client selected-object/purse state
before enqueuing `0x23`; modal cancel bypasses it. Full drag refund and native
delivery chronology remain Unknown (MENU-093).

<a id="the-one-thing-a-consumer-must-not-get-wrong"></a>

## State and identity

**There are two clocks, and every published "tick" belongs to one of them.**

- `server+0x04` — the **sub-tick**. Incremented once per `R0193`. Commands, the world
  update and every actor's `vt+0x18` run here. `R0547` measures against it; the actor's
  action deadline `actor+0x138` is in these units.
- `server+0x00` — the **full tick**. Incremented once per **16** sub-ticks. Regeneration
  (`vt+0x14`), corpse decay, the shop restock divisor and the lose watchdog are filtered on it.

A consumer that treats the two as one will run regeneration sixteen times too often or the shop
restock sixteen times too slowly.

<a id="the-rate"></a>

## Tick rate

`campaign+0x3f0` = `1000 / R` milliseconds per sub-tick, `R` selected by `campaign+0x3f4 ∈ [0,8]`,
default **4**. A full tick is 16 sub-ticks.

| speed index | sub-ticks/s | ms/sub-tick | ms/full tick |
|---|---|---|---|
| 0 | 8 | 125 | 2000 |
| 1 | 10 | 100 | 1600 |
| 2 | 12 | 83 | 1328 |
| 3 | 14 | 71 | 1136 |
| **4 (default)** | **16** | **62** | **992** |
| 5 | 20 | 50 | 800 |
| 6 | 24 | 41 | 656 |
| 7 | 28 | 35 | 560 |
| 8 | 32 | 31 | 496 |

The pacing is a **deadline** loop with catch-up, so the figure is a ceiling. Under
`campaign+0x40c != 0` there is no deadline at all.

<a id="the-objects"></a>

## Object ownership

| Object | Where | Size | What it is |
|---|---|---|---|
| session window | `AfxGetThread()->vt+0x7c()`, = `app+0x1c` | `0x6d0` | the MFC main frame; carries the phase, the screen mask, the clock, the campaign state and the chargen inputs |
| server singleton | `[L00285]` | `0x174` | the simulation's own root; carries both counters, the participant flags and the difficulty (`UNIT-GATE-012`); [SAV](../sav/format.md)'s own "world" names this row, not the "world" row below (`SAV-657`) |
| world | `[L04624]` | `0xa4558` | built by `R0118` from the loaded map |
| session/AI | `[L00004]` | `0xc320` | built by `R0077`; carries the 50×50 diplomacy matrix at `+0xa9c4` |

<a id="the-session-windows-fields-as-far-as-read"></a>

## Session-window fields

| Offset | Meaning |
|---|---|
| `+0x1c` | `HWND` (also published to `[L06446]`) |
| `+0x3dc` | screen-state **bitmask**; bit 0 = a map session owns the frame and ticks; the whole word `0` is the town/home state (`SHOP-TOWN-023`) |
| `+0x3e4` | sub-tick phase 0..15 |
| `+0x3ec` | epoch, `timeGetTime` at the last phase-0 |
| `+0x3f0` | ms per sub-tick |
| `+0x3f4` | game-speed index, saved as `[GameOptions] Speed` |
| `+0x40c` | non-zero selects the unpaced loop |
| `+0x494` | hero class bits: **6 = mage, 7 = female** |
| `+0x4b0/4b4/4b8/4bc` | the four point-buy stats |
| `+0x4c0` | fifth chargen byte |
| `+0x4c4` | appearance index (low 6 bits of the wire byte) |
| `+0x548` | the mission record (`SHOP-MISSION-018`) |
| `+0x65c` | difficulty index 0..2 (`UNIT-GATE-012`) |
| `+0x660` | mission number; `% 10 == 0` is a main mission |
| `+0x66c` | accumulated play time, in **full ticks** |
| `+0x6b8` | this process owns the simulation |
| `+0x6bc` | phase, `{0,1,2,3}` |

<a id="the-lifecycle"></a>

## Lifecycle

```
InitInstance              new(0x6d0) -> R0341 -> app->m_pMainWnd
                          campaign+0x65c = 1, campaign+0x3f4 = 4
logo / intro              R2036, R2035
chargen                   R0909 ("music\chrgen.wav")
                            precreate+0x1d0 -> campaign+0x65c   (difficulty 0..2)
                            precreate+0x1cc -> T[] -> <<6 -> campaign+0x494 (mage, female)
                            portrait arrows -> campaign+0x4c4
new campaign              R1382: campaign+0x6bc = 2, +0x494 = 0 | -female | -mage
create the hero           R0822  -> command 0x48 -> R0823
                            cmd+0x0f = appearance | classBits
                            "PC_Danath" / "PC_Naira" / "PC_Fergard" / "PC_Reniesta" + ".m|.f<n>"
                            budget 140; over budget -> all four stats := 25
start a session           R0512(name)
                            mission number := atoi(name before ".alm")   (no ".alm" -> error 3)
                            R0966: World\Mission\<n>.ini
                            server+0x154 ? restore game0000.sav (error 4)
                                         : R0128 (error 5)
load the map              R0128            [base = server+0x44; every
                                                     server+N below is 0x44 low]
                            server+0x20 ? "Scenario\" + name : name   (= server+0x64,
                                                     the mission number)
                            map -> server+0x0c, +0x120 -> players -> WORLD -> type-4
                            -> SESSION -> diplomacy -> type-6 spawn -> ... -> server+0x14c
playing                   R0560 -> one of four idle arms -> R0147 per sub-tick
finish mission            command 0x41d -> R0701 case 0x41d, arm by phase
                            campaign+0x66c += server+0x04 / 16
                            tear down, cutscene, then command 0x3f (shop cap) on the way to town
```

### New-campaign start state

The new-campaign arm constructs the session, resets the campaign record, and loads mission 10 in
that order. Mission 10's text-document entries 1, 2, and 3 therefore exist on the campaign record
before command `0x48` creates the hero and before the first simulation tick. The participant purse
is separate state on the `Player`. Its constructor sets it to zero, then participant factory
`R0913` replaces zero with **100** before hero creation. Join next applies a zero delta and
sends opcode `0x67`, so the first playable client state and the first-sub-tick automatic save both
carry 100 (`PARTY-MONEY-024`, correcting `SESS-START-034`'s former constructor-only reading).

### Authored Player control and join

`Player+0x28` retains the full authored dword when roster transfer copies
`CPlayer+0x30`. Reuse requires server+0x0c zero, roster slot1 and exactly
one manager entry; it selects the first entry and leaves its ordinal+0x08
unchanged. A new entry receives the authored ordinal before registration.
Participant factory instead queries `Self` only in mode zero, then writes
control zero on its reused or new entry. Only a new entry is registered.
Native order and object identity across these transactions remain Unknown.
— SESS-072

Manager registration derives the scan-mask word+0x2c only for control zero
and clears it for nonzero control. Its matrix call is conditional on a
nonnull session global. The matrix stores the control low byte, then forces
that byte to zero only for complete value2. It chooses relations by the
stored byte's zero/nonzero class. Thus `0x100` and `2` yield column0 zero,
while `0x102` yields two; all three are nonzero to the manager's dword test.
Map transfer precedes the new session's construction and the authored
relation-row loop, so a matrix result cannot be inferred solely from the
transfer store. — SESS-073

Message96 emits full-word `!=0` as a bit, and its interface receiver changes
only `CPlayer+0x30` bit0. Player persistence instead writes and reads four
bytes unchanged on the selected buffered path. These projections do not
establish a universal Boolean schema. Full native load/join, UI, save/reload
and aliases outside the selected population remain Unknown. — SESS-074,
SESS-075

<a id="the-map-load-in-order"></a>

## Map-load sequence

`R0128` receives `server+0x44` as its map sub-object. Its
receiver-relative+0x20 test is therefore server+0x64, the mission number;
nonzero selects the campaign Scenario path. Mode stores instead reload the
global server pointer and use actual server+0x0c,+0x120,+0x14c. The blanket
0x44 adjustment in `SESS-MAP-010` is retracted; apply the base belonging to
each access. — SESS-MAP-010, SESS-DEFEAT-064

`R0128` is the only routine that turns a map into a session, and its order is normative:
`map+0xd4` → `server+0x0c`/`+0x120`; `server+0x170 = map+0x94`; players (type-5); **the world**;
the type-4 walker; **the session and its diplomacy matrix**; the type-6 spawner; two further
passes; multiplayer-only `R1603`; `server+0x14c`. The world exists before any placement is
walked, and the diplomacy matrix before any unit is spawned.

The six load failures the engine names for itself: *File not found*, *Not a map file*, *Wrong block
number*, *Map version too new (update loader!)*, *Tiles block not found*, *Altitudes block not
found*. The last two are `ALM-REQ-055`'s required pair, named from the loader's own side.

<a id="where-an-embedded-campaign-map-comes-from"></a>

## Map resource selection

`server+0x20 != 0` prefixes `"Scenario\"`, which is `RES-IDENT-034`'s dispatch to `scenario.res`;
the campaign names its map `"%d.alm"` from the mission number. A loose map keeps its bare name and
falls to the loose-file tier. **The field is `server+0x64`** — see the base warning above; the
routine runs on `server+0x44`, and the field it tests is the mission number itself.

## Commands

Everything the simulation is told is a byte opcode executed inside a sub-tick. `R0061`
dispatches on `cmd+0x04`: `>= 1` selects the **order** space, opcodes `0x14..0x26` over 19 direct
dwords at `L00529` ([AI](../ai/format.md) has the vocabulary); `== 0` the **session** space, opcodes
`0x02..0xBE` through a 189-byte index at `L00564` into a 29-entry table at `L03265`, of which
**28 opcodes are live** and 161 map to the exit. The two spaces are disjoint and share the opcode
byte, so an opcode has no meaning without `cmd+0x04`. `SESS-CMD-008`, `SESS-CMD-015` and
`SESS-CMD-016` carry the full independently read tables and their opcode-for-opcode agreement.

**It is fed by a pool, not dispatched directly by the interface.** `R0061` has one caller, `R0191`,
which is a drain loop: sets the this pointer to the constant L00522, calls the routine at R0513, dispatches the result if it is non-null, and repeats. No
routine dispatches a command it synthesised. `AI-INPUT-121`, `AI-SELECT-122`, `AI-PANEL-123`,
`AI-MINIMAP-124`, `AI-KEY-125`, `AI-CURSOR-126` and `AI-INPUT-127` establish the mission input enqueuers in
[AI](../ai/format.md), including the Drop Gold dialog's `0x23`. Other non-input command producers
are outside that input mapping; no universal producer list is established.

The command's shape, from the instructions that read it: `+0x04` space, `+0x05` a `u16` player,
`+0x09` the opcode, `+0x0a`/`+0x0c` a column and a row (a `u32` parameter id under session `0x46`),
`+0x0e` a `u16` id or slot (a signed `u32` parameter value on all four live arms of session
`0x46`, spanning `+0x0e..+0x11` and overlapping the `+0x10` index below), `+0x10` a `u16`
index, `+0x12` a `u8` member count and `+0x13 + 2i` the `u16` member ids. **That `u8` is a
real ceiling: one order commands at most 255 units.**

Session opcode `0x46` is a generic **set-parameter** with its own sub-switch on the `u32` at
`cmd+0x0a`: 128 ids in range, four live (`1`, `2`, `3`, `0x80`), 124 on a default that prints the
engine's own `"Request to set unknown parameter "`. Id 1 is the withdraw knob (`AI-CLASS-030`);
id 3 writes Player+0x58, mapping 0/1/2 to 100/50/0 and accepting raw 3..100 while retaining
the prior value outside that range — id 3's value is the same `+0x0e` dword above, so a
`+0x10` index fill with `+0x0e = 1` reads above 100 there and stores nothing.
The earlier "parameter 3 is the withdraw knob" wording is corrected. — SESS-PARAM-017
(partially retracted), SAV-726

**Customisation.** 161 empty session opcodes, 4 empty order opcodes and 124 empty parameter ids
are all *absent implementation* and free: no shipped `.alm`, `.res` or save carries a command
opcode, so adding an arm changes the bytes of nothing the original wrote. The ceilings `0xbe`
and `0x26` are compiled constants over bytes that already hold `0..0xff`. What is **not** free is
the command's own layout — it is the multiplayer wire.
