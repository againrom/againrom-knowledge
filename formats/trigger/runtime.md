# Compilation, evaluation and persistence

[Reference](format.md)

## State and identity

**There is no evaluator for victory.** Winning and losing are two integers that an ordinary
*action* increments. A mission with no action that increments them can never be won and can only
be lost by the hero dying. Implementing "the engine checks whether the objective is met" produces
a game that never ends.

## Compiled state

The map's authored script is **compiled once at load** into three flat stores on the session
object, and never read from the map again:

```
session+0xc2ac   check list      each element an 80-byte record
session+0xc2b0   instant array   each element a 72-byte record
session+0xc2b4   pattern list    each element a 24-byte record
session+0xbd34   int  slot[100]  the result/variable register file
session+0xbec4   byte latch[1000]  one per pattern, indexed by the trigger's array position
```

Both array sizes are hard: nothing bounds-checks them at build time. A map with more than 100
conditions or more than 1000 triggers corrupts memory. The shipped maxima are 64 and 36.

Slots 90–92 have literal, non-indexed writes only at session construction.
Slot 93 (`session+0xbea8`) also has two runtime writers, each followed by
the check evaluator and then the pattern evaluator. Other slots are written
through script-supplied indices. Installed check indices do not reach 93;
the corresponding bound on pattern indices remains Unknown. The exclusive-
slot 93 literal-write clause of partially retracted `SAV-651` is narrowed to
runtime writes. — SAV-650, SAV-651, SAV-658

## Numbering

The builder walks the map's action and condition lists in file order; the loader appends each
node at the tail. Actions below opcode `0x10002` take instant subscripts 0, 1, … only when built.
Conditions take registers 0, 1, …: an opcode `0x10002` condition always, any other below it
only when built. A trigger slot naming an action that took no subscript runs subscript 0, and a
pair naming an unbuilt condition reads register 0. A condition above `0x10002` takes no register
and its id reads the next condition's. Triggers become patterns in array order, and each
pattern's instants keep slot order. — TRIG-BINDORDER-101

## Evaluation pass

Once per **full tick** — before any actor is walked, ≈ 992 ms at the shipped speed default
(`TRIG-EVAL-001`) — the engine reads, increments and stores slot 93 (`session+0xbea8`). It then
calls the check evaluator, then the pattern evaluator, with no other instruction between any of
the three steps. Two engine sites run exactly this triple. `R0076` runs it, reached
unconditionally from the dedicated-server loop `R0075` and again from a
`server+0x04 % 16 == 6` gate inside `R0147` — the same function that advances the
full-tick counter, on its own separate `% 16 == 15` gate (`SESS-TICK-004`). A second, minimal
routine, `R1614`, runs the identical triple as its entire body; no caller was located for
it in `.text` or as a whole-image dword reference, so whether or how it runs is Unknown. — SAV-650

1. **Every check** is evaluated and writes its own slot. One 22-arm dispatch on the check's
   opcode (1..22); an opcode outside that range, and the two dead arms 11 and 13, write nothing,
   so the slot keeps its previous value.
2. **Every pattern** is evaluated. Up to three triples `(slotA, slotB, code)`; the six codes are
   `0 ==`, `1 !=`, `2 >`, `3 <`, `4 >=`, `5 <=` and anything above 5 is false. The triples are
   **ANDed with short-circuit**; there is no OR.
3. A pattern that passes sets its **latch** and runs its up-to-four instants, in slot order. One
   34-arm dispatch on the instant's opcode (1..34); arm 9 is dead and out-of-range logs
   `"Script: Bad instant %d"`.

## Firing discipline

`trigger+0xb4` in the map file is the once flag.

- **1** — fire at most once for the whole session. The latch gates re-entry.
- **0** — the latch is cleared at the head of every pass, so the pattern fires **again on every
  full tick** its conditions hold.

## Persistent state

`slot[100]`, `latch[1000]` and the three outcome integers are in the **world half** of a save
(`SAV-SESS-031`): present in mid-mission saves and absent between missions. **A mid-mission load
restores this session block first (`L07069`) and rebuilds the map trigger programme later
(`L01841`).** The builder can preset at least an authored-variable result slot, so builder-first
followed by a blanket saved-slot overwrite is not equivalent. A consumer that saves mid-mission
without the latch array re-fires every one-shot trigger on reload. — TRIG-SAVE-008 (load
order amended; former saved-population count withdrawn), SAV-RECON-268 (its
universal saved-row selector for Human definition pointers is withdrawn; the
session-block order stands)

The selected positive overwrite is the authored opcode `0x10002` check arm: it copies node+0x48
into the current compiled register slot and advances that slot. Successfully compiled earlier
normal checks and constants determine its index; a resolution-rejected normal check does not
advance it. Later `R1485` repairs `session+0xa9c0` only. Its adjacency to `L07742` does not prove
a LOAD call to that separate register writer. — SAV-652 (partially retracted), SAV-710,
SAV-711

## Mission end

```
instant 4  ->  session+0xb3ac++     win
instant 5  ->  session+0xb3b4++     lose
check   18 ->  session+0xb3b4++     lose, when its unit is dead; writes no slot
```

At full-tick reporting, script loss `== 1` precedes win `== 1`, but only after the active-human,
latch and primary-fall gates. Latch `player+0x3c >= 2` cannot reach either counter arm; campaign
mode returns without automatic latch reset. Loss-to-win is therefore not enabled merely by
incrementing the loss counter past 1. With a living primary and latch below 2, counter outcomes
send `0xb4`/`0xb5` and write latch 2/1 without a mode predicate (`TRIG-END-009`, amended).

An earlier branch tests the primary pointer `player+0x34` and its stage, not every companion.
It writes latch 2 and sets owned actors' HP to -50. Its failure packet is restricted to joined
participants with actual `server+0x0c == 0`; nonzero-mode recovery has separate entry-active and
HP-below--53 gates. Campaign panels also pause frontend ticking (`MISSION-DEFEAT-045`,
`SESS-DEFEAT-065`).

## Parameter binding

Every arm below is written in terms of `p0`, `p1`, … and of six reference fields. **Those are
not the file's `Par0..Par9`.** The builder loops all ten `(value, type)` slots and dispatches
`type − 2`, bounded by 7, through an 8-entry table (`TRIG-PARAM-030`):

```
type 2  -> group      first -> rec+0x34, a second one -> rec+0x3c
type 3  -> player     first -> rec+0x38, a second one -> rec+0x3c
type 4  -> unit       first -> rec+0x30, a second one -> rec+0x3c
type 8  -> item       rec+0x40, a WORD, and the stored value is value + 0xe18
type 9  -> building   rec+0x44, resolved by NAME
everything else       the next plain-int slot, rec+0x08 + 4*k
```

`k` is a **separate** counter that advances only on the last line, so a reference-typed slot
takes no `p` index while an *unused* slot — the editor's `<None>`, type 0 — does:

```
p_index(i) = i - (number of reference-typed slots before i)
```

The file's own layout is the opposite (`ALM-TRIG-045`: parameters are stored by slot, not
packed), and reading an arm against the file rather than against this rule silently changes
what it does. The standing example is instant 23, whose every shipped node authors `Player`
in `Par0` and `Money` in `Par1`: read by file slot, the arm adds a player index to a purse
and never reads the money at all.

The item slot is an **offset, not a code**. `0xe18` decodes under `ITEM-CODE-029` as class
14, index 24, so an authored `Item` value `V` names class-14 index `24 + V`. The shipped
labels agree: `V = 11` is `"Add item35 …"`, `V = 13` is `"Add item Tooth(37) …"`. `V` above
231 carries into the class nibble and the factory returns null.

**No arm null-tests a reference field**, and none needs to: a node whose reference fails to
resolve is dropped whole at build time and logged (`TRIG-BIND-010`; its slot-0 clause is
partially retracted for a condition above `0x10002`, see Numbering).
