# Diplomacy

[Reference](format.md)

## Diplomacy matrix

50 × 50 bytes at `session+0xa9c4`, all zero at construction — with nothing written, no unit ever
finds an enemy. Cell `[i][j]` is indexed by `Player+0x04`, stride 50; the engine reads three
bits of it.

| bit | meaning |
|---|---|
| 0 | *i* treats *j* as hostile. This is the only bit acquisition tests. |
| 1 | locked: combat may not turn this pair hostile. Forced on every diagonal. |
| 2 | read (`AND …,0x7`) and unused by any shipped map. |

Column0 stores the low byte of `Player+0x28`, replaced with zero only when
the complete dword equals 2. Row0 marks the slot live. This projection is
distinct from the full-word tests elsewhere: `0x100` produces column0 zero,
whereas `0x102` produces two. — SESS-073

<a id="the-six-writers"></a>

### Named matrix writers

These are distinct entry paths. Registration and the authored map-row writer
have separate callers; native ordering is not inferred from this list.
— SESS-073

| # | when | routine | writes | symmetric? | respects bit 1? |
|---|---|---|---|---|---|
| 1 | map load | `R0128` | the `.alm` type-5 record's sixteen `u16` at file `+0x2c` into columns 1…16 of that player's row (low byte only), then forces the diagonal to 2 | one way | n/a |
| 2 | a player joins | `R0263` | column0 = low byte of control, except full value2 maps to 0; row0 = 1; template selection tests the two column0 bytes for zero; forces diagonal2 — SESS-073 | both | no |
| 3 | a player leaves | `R0264` | row 0 = 0 for that slot | n/a | n/a |
| 4 | a mission join | `R0131` / `R0132` | clones a reference player's row *and* column, allies with it, forces neutrality or 2 between participants | both | no |
| 5 | script action **10** | `R0262` | `matrix[p0][p1] = (v &~ 3) + p2` | **one way** | **no — it clears it** |
| 6 | session command **0x45** | `R0061` | assigns the whole row `matrix[setter][*]` from a `u16` array in the command body, each element `& 7` | **one way** | **no** |
| 7 | a blow landing, or a spell cast | `R0126` | sets bit 0 in both directions, each direction **separately** gated on `(cell & 3) == 0` | both | **yes** |

The template of writer2 is four bytes at `session+0xa9bc`, initialized to
`{1,1,0,0}`. With that template, relations between different zero/nonzero
classes are hostile both ways; two unequal nonzero bytes select the same
class. The diagonal is forced to 2. Registration calls this routine only
when the session global exists, and map load has a separate later relation
writer. — SESS-073

These local template, live-slot, diagonal and leave operations retain
AI-DIPLO-083, partially retracted for its universal owner interpretation
and unequal-byte hostility rule. Those interpretations do not apply. Native join order remains
Unknown.

### What a change reaches

Writers 1–6 are **bare**: they change the byte and nothing else. The one exception is the
spell-cast entry into writer 7, which rebuilds the caster's group candidate list in the same
instruction stream. Everything else propagates only because a consumer re-reads the cell — which
for the 95.6 % of hostile placements under a group order means the next AI tick, since the
candidate list is rebuilt and re-filtered on every evaluation. **An order already issued keeps
running**: making a faction friendly does not stop a unit that is already attacking, it stops that
unit being re-selected.

### What the script can do with it

Action opcode **10** writes `matrix[p0][p1]`, one direction, clearing the low two bits and then
**adding** `p2`. Nothing clamps `p2`, so a value ≥ 4 carries into the bits above and ≥ 256
truncates in the byte store; shipped maps use only `{0, 1, 2}`. Check opcode **10** reads
`matrix[A][B] & 3` into a script slot — **narrower than the action can write**.

Mutual changes require two action nodes, one for each ordered pair;
one action never updates the reverse relation. — AI-DIPLO-005

### Persistence and the wire

The matrix is serialised **verbatim** as part of the 2508-byte (`8 + 2500`) sub-object at
`session+0xa9bc`, in the world half of a save, and read back verbatim. Nothing re-derives it from
the map on a save load. Outbound, a changed row is broadcast as message `0xb9` carrying the whole
row widened to `u16`; inbound, session command `0x45` receives a row in the same width. The
relation is byte-wide in the simulation and **word-wide on the wire in both directions**.

The relation is directional: `matrix[A][B]` and `matrix[B][A]` can
differ, including hostility bit 0. Store ordered pairs independently.
— AI-DIPLO-005

`World\Data\ai.reg` contributes one value on this path: `[Scanning] MinimalGuardRange`, a floor
on a group's notice radius. The code default is 10; **the shipped file sets 8**. The file has four
records in total and its other scalar, `[Tasker] IntelligentCons` = 15, has no literal anywhere in
the image — nothing reads it.
