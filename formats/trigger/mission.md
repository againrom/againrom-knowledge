# Mission identifiers and entry/exit

[Reference](format.md)

## Identifier bands

`Target_Unit` is three id spaces, and a consumer that treats it as one will fail to resolve four
fifths of the shipped references:

```
value <  10001    a unit id from the map's own type-6 records
10001..11000      a hero ordinal, resolved when the script is bound (next section)
                  (and unresolvable while the world flag at `+0xc` is set: map player capacity above 1)
value >  11000    an index into a static name table in the executable
```

## Hero ordinals

A hero ordinal names a role, not a roster position. With `k = value - 10001`, the resolver returns
the first **named** actor, scanning the first player's flat actor list and then the actor registry,
whose typeID test, class bit and face satisfy the predicate of `k`. Health, corpse stage, placement, owner,
the `Hero` flag, hire or join origin and `AddHero` are not inputs. A set world flag `[L00285]+0xc` (map player
capacity above 1), a null primary character (`Player+0x34`) and a scan with no match all give 0. — TRIG-HEROORD-075

The typeID test (called sex here) is the actor type in `{0x22, 0x24}`, class is `actor+0x4c & 4`, a bit a positive
mana column sets, and face is `actor+0x4b & 0x3f`. Whether a shipped companion's typeID and class bit satisfy
the tests that follow is not established.
`k = 0` requires the primary character's own three values. `k >= 1` applies the `Flags` tokens of
registry section `npc(21+k)`:

```
10001  the primary's own sex, class and face
10002  npc22  Hero,Mage,!MySex,Start      a mage of the other sex, default face
10003  npc23  Hero,!Mage,!MySex,Start     a non-mage of the other sex, default face
10004  npc24  Hero,!MyClass,MySex,Start   the other class, the same sex, default face
10005  npc25  Hero,Face,!Female,!Mage     a male non-mage whose face is 1
10006  npc26  Hero,Face,Mage,!Female      a male mage whose face is 4
```

`Mage`, `Female`, `MySex` and `MyClass` require the property (the last two relative to the
primary), the `!` forms require its absence, and `Face` requires the template's face. A template
with no `Face` token requires the actor's face to equal the default face of its own class and sex
(`Face` of `[MaleFighter]` 5, `[FemaleFighter]` 1, `[MaleMage]` 3, `[FemaleMage]` 1). The tokens
`Hero`, `Human`, `Me`, `Start` and `Platoon` are never tested here. A token matches as a case-sensitive
substring of `Flags`. Two ordinals can resolve to one actor. The published node table of the 38 EN maps
uses 10001 to 10006 and nothing above. — TRIG-HEROTPL-076

The resolution happens while a script is bound, once per node reference: at map load, and again when a
saved game is restored, after the players and the actor registry are loaded. No check or instant arm
calls the resolver, and no other routine in the swept image compares against the hero band. — TRIG-HEROBIND-077

No match logs `Can't resolve hero %d.` and the node is not built; the load continues. Values above
10235 index past the 256-entry template table, and an `npc` section with empty `Flags` has no
template record; neither is guarded. — TRIG-HEROFAIL-078

## The drop table — where the player lands

The authored action `0x10002` is never dispatched. At build time its `value[0]` and `value[1]` are
truncated to bytes, packed `(y << 8) | x`, and appended to a `CWordArray` at scriptObj`+0x28` =
`server+0x6c`, where scriptObj is the map sub-object at `server+0x44` and `server` is the
singleton `[L00285]`. **Every installed map carries one** — campaign and skirmish alike, and
on 9 of the 10 loose maps it is the map's only script node.

`R0065` runs on `server`, not on a map object, and three details decide whether a
reimplementation lands the player where the engine does (`MISSION-DROP-002`, whose earlier
`mapObj` base name is superseded; the displacements and behaviour stand):

```
if (server+0x0c != 0 && (u16)player+0x60 != 0)  x,y = low,high byte of player+0x60
else if (array.GetSize() > 0)                   i   = R0861(GetSize()-1)   <- RANDOM index
                                                x,y = low,high byte of array[i]
if (x * y == 0)                                 x   = R0861(0x46) + 0x1e   <- 30..100
                                                y   = R0861(0x46) + 0x1e   <- independently
                                                log "no drop location in .alm - random used"
```

`server+0x0c` is `arg0 < 2` from the server's constructor, and the campaign start passes 2, so
on the campaign path `player+0x60` is never read. `R0861(n)` is
`(rand() * (n+1)) / 32768` — **inclusive of `n`**, and not a modulo. The array is **not** read
at `[0]`.

## Starting the mission

`R0065` never reads the type-6 array. It iterates the **player's own unit list** at
`player+0x20`, placing the hero (`player+0x34`) at the drop cell exactly and everything else
within `ftol(max(5.0, sqrt(nUnits) + K))` of it. Over all 28 campaign maps of both roots, roster
slot 1 owns 19 of 2333 type-6 records, on five maps: `41.alm` 3, `71.alm` 4, `120.alm` 10,
`150.alm` 1 and `151.alm` 1. The spawner installs those records into the player's own
containers, so the walk moves them off their authored cells to the drop cell; the other 23 maps
place nobody for the player. A consumer that builds the party from the map file starts 23 of
the 28 campaign missions empty, and on the other five keeps 19 units at authored positions that
the engine replaces with the drop cell (`MISSION-START-001`; its former clause that a campaign
map places nobody for the player is superseded by this census).

The routine's second half — a by-name lookup of a node called **`"Humans"`** whose children are
`label#id` strings with their own `(x, y, radius)`, or `param[1] == -1` meaning "at the drop
cell, radius 8" — is the `.ini` overlay's entry point below, and **no shipped map has such a
node**, so on shipped data it never runs.

## Ending it

`instant 4` is the only thing that wins, and **every campaign map carries exactly one**
(`MISSION-WIN-003`). None of the ten loose maps carries any: a skirmish map cannot be won,
only lost. A win action may be referenced by more than one trigger — four on `81.alm` — so
bind by node id. `MISSION-WIN-003`'s win-action census stands; its lose-arm map counts are
superseded: instant 5 is authored on 15 of 28 maps (`MISSION-TYP-028`), and check 18 on 7 EN
and 6 RU maps (`MISSION-LOSE-025`).

`check 18` is the "protect this unit" objective and it is authored as a trigger with **no
action at all**: the check writes no slot, so the pattern never passes, and its only effect is
the lose it raises when its unit dies. A check is armed by being **authored**, not by being
referenced — the builder gives every check node a slot and the pass evaluates every check
(`MISSION-VIP-004`; its corpus map count is partially retracted, the arming rule stands).

## The register file is shared

The build assigns check node *i* the slot *i*, constants included, and an authored **variable**
addresses the same `session[0xbd34 + p0*4]` array with its own literal number. Any variable index
below a map's check-node count is overwritten every full tick. `60.alm` ships exactly that
collision — variables 32 and 33 against 64 check nodes (`MISSION-SLOT-008`).

## Mission text

`instant 2` carries a **number**, not a string. `R0701` opens
`main.res::text/battle/m<mission>/event<NN>.txt`; the same block holds `briefing.txt`,
`briefmap.txt`, `title.txt` and `tips<NN>.txt`. The files are markup with `<NPC=n,Part=k,…>`
tags whose `n` is an `npc.reg` section number (`MISSION-TEXT-005`). **A number whose file does
not ship is a silent no-op** — the window is never built and nothing is logged; the window
itself, its lifecycle and the tag vocabulary are [DIALOGUE](../dialogue/format.md)
(`DLG-WIN-001`…`DLG-MARKUP-007`).

## `World\Mission\<n>.ini`

Built and opened at every map load, absent from the install, and its absence is a **silent
no-op** — the open fails, the reader returns immediately, and the one call site discards the
result. It is a sectioned text overlay whose sections the engine looks for by name:
`Humans.Hero` (`Name`, `Bag.`, `Armor`, `Weapon`, `Shield`, `Item`), `Outposts`, `Patrol`,
`StandGround`, `Items`, `Players`, `Mission`, `Monsters`. **Nothing in the trigger machinery reads
it**; its consumers are hero import, patrol paths, outposts and item setup.
