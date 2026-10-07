# Actor presentation and names

[Reference](format.md)

## Information-display state

`R0590` (`CUnit`, one caller) copies **25 stores** off the prototype actor in one fixed
order (`UNIT-PANEL-010`, superseded only in its scope: the copy map and arithmetic stand). In order: the four stats, health and
its maximum, mana and its maximum, toHit, defence, absorption, the damage pair, speed, sight, the
five damage-kind resistances, the five elemental protections, and a derived byte
(`2` when the typeID is `0x49`, else `0`).

**It reads no `units.reg` field**, so `InfoPicture` and `DescText` do not arrive here.

**Four values are computed on the way out.** When the app's `+0x6bc` is 2, `[L00285]` is
non-null and that object's `+0x84` is not 2:

| `+0x84` | what the display is given |
|---|---|
| 1 | `healthMax := ftol(healthMax × 0.66)`, then `health := healthMax` |
| 3 | `toHit += 50`, `defence += 50`, `healthMax := ftol(healthMax × 1.5)`, then `health := healthMax` |
| 2 | untouched |

`0.66` and `1.5` are the doubles at `L12621` / `L12622`. `+0x84` is the setting above, and
this table is the display **mirroring** the spawner's rule for a prototype that never went through
it — not a second, display-only rule.

`UNIT-PANEL-011` bounds a displacement sweep; the numeric card readers are
now identified. `R0877` reads Body/Agility/Mind/Spirit from drawable
`+138/+13b/+139/+13a`. It reads attack/defence from signed words
`+146/+144`, armour from byte `+142`, and min/min-plus-spread from
`+148/+149`. It walks `+14b..+14f` and `+150..+154` in two five-row
loops. The first captions differ for types 23/24. The earlier indexed
`L09202` consumer belongs to a spell-mask loop, not these numeric rows.
Disclosure levels gate rows; global `L01661` can force level 7.
Its runtime cause is Unknown. — SAV-1126

Prototype copy precedes packet overwrites. A non-owned Human outside types
33..63 receives mask `50fb`: current/max health and mana remain available,
but primary/combat/sight/speed/extended values are omitted. Types 9/24 also
skip the prototype-stat copy. The selected client path therefore preserves
existing primary bytes; complete initialization and first-frame contents
remain Unknown. An original SAVE reads actor fields, not those drawable
bytes. — SAV-1125, SAV-1126


## Drawable class (`UNIT-APPEAR-030`)

Its own `typeID`, unchanged, and nothing derives it.

The server sends `word[actor+0xe]` — `R0880`, nine instructions, no arithmetic — under field
mask bit `0x4000`, together with a face byte `actor+0x4b`. The client assigns it to `drawable+0x20`,
which is what the frame selector `R0553` subscripts the `units.reg` class array with. For a
class a map places, `R0590` leaves that field alone: its two outer arms are `>= 0x40` and
`< 0x1a`, and the shipped roster occupies exactly `1..27` and `64..80`, so both arms are total on the
shipped population. The appearance routine `R0551` gates on `drawable+0x18c` bit 0, which no
arm sets for these actors, and returns immediately.

So the drawn class of a non-hero actor is a **stored** property and needs no recomputation, and its
sheet is the class record's own `File` in the ordinary way (`REG-UNITS-049`).

This is the opposite of the hero arm, where the id is discarded on arrival and the drawn class is
recomputed from equipment on every state message — [hero appearance](../hero/appearance.md). **One code path
cannot serve both.** A reimplementation that derives a class for every actor is wrong for every unit
a map places; one that stores a class for every actor is wrong for every player character.

**And it never carries visible equipment.** The equipment sender `R0669` calls
`actor->vt+0x30()` before either of its opcode arms and returns without sending anything when it is
false — the same humanoid predicate whose false arm prints *"Trying to takeoff armor from non
humanoid"* (`R0182` returns 0 on the base actor class, `R0183` returns 1 on both human
classes). So a non-humanoid actor's twelve client-side visible-equipment slots stay null for its
whole life, and there is nothing for the hero arm to derive from even if it were reached
(`UNIT-APPEAR-031`).

Claims: `UNIT-APPEAR-030`, `UNIT-APPEAR-031`, `HERO-APPEAR-040`…`045`.


## Display names (`UNIT-NAME-039`…`UNIT-NAMETAB-041`)

A class's displayed name is **line `ID` of `main\text\unitname.txt`**, not `units.reg`'s `DescText`.
The information display `R0877` reads it at `L12730`: it passes `drawable+0x20` (the typeID) as the argument to `R0668` with the table object `L11124` as the receiver. The file has 81 lines against the class array's 81 slots
(`maxID + 1`); 33 of the 34 shipped classes have a non-empty line and every one of those 33 differs
between the roots, while all 42 empty lines are identical. Six non-empty lines (17, 18, 20, 67, 77,
78) belong to no shipped class. The one class with an empty line is ID 2, `Unarmed Fighter with
Shield`.

`DescText` is a second naming system with no reader in `rom.exe`, byte-identical on both roots, and
it disagrees with the shown name in content — `Death Star`/`Daemon`, `Ghost`/`Spirit`,
`Goblin`/`Goblin Pikeman`. A consumer that renders `DescText` renders the wrong string, in the
wrong language.

An **instance** name exists as well — `actor+0x80`, an MFC `CString` on every actor — but no shipped
surface draws it for a non-hero, so two units of one class are not distinguishable by name.
