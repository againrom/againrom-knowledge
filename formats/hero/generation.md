# Character generation and identity

[Reference](format.md)

## Integer conventions

City sale reaches full derive only if `R0451` finds a changed truncating
signed16 load/capacity quotient. It uses stored prior load, updates the load
word with wrapping arithmetic, and has no zero-divisor guard (`SAV-CITYSALE-513`).
A successful school purchase instead ends with unconditional Human derive.
Derive rebuilds live fields and clamps current health/mana but does not rebuild
the modifier block from equipment: it consumes retained modifiers, with the
conditional negative-speed reset of `+d8`. Its direct stores leave base and XP
blocks intact; school updates selected skill/base/XP before the derive call.
This bounded producer result does not establish runtime callback chronology or
safe initial values (`SAV-CITYDERIVE-515`).

- `ftol` is `__ftol` (`R0279`). It forces x87 rounding-control to **11 = truncate toward zero**
  before `FISTP qword`, then returns the 64-bit integer in `EDX:EAX`. **It truncates. It does not
  round.** A caller that stores only EAX narrows that result to the low 32 bits.
- `/` between integers is `IDIV` or the `CDQ`/`SUB`/`SAR 1` idiom: also **truncation toward zero**.
- The **only** rounding-to-nearest in this whole area is the `+ 0.5` inside `T(n)` below.
- Every intermediate of the health and mana chains is stored back into a **16-bit** field, so the
  truncations compound; a consumer that carries doubles through will drift.
- `pow(1.1, s)` is `R0874` and `log₁.₁(x)` is `R2109`. The `1.1` is pushed as `pow`'s
  **first** argument, so the stat is the **exponent**.


## Primary attributes

| name on screen (RU) | engine name | actor | record | chargen panel row |
|---|---|---|---|---|
| Сила | `Body` | `+0x84` u16 | `+0x138` u8 | 0 |
| Ловкость | `Reaction` | `+0x86` u16 | `+0x13b` u8 | 1 |
| Разум | `Mind` | `+0x88` u16 | `+0x139` u8 | 2 |
| Дух | `Spirit` | `+0x8a` u16 | `+0x13a` u8 | 3 |

The engine names are `scenario\npc.reg` key literals, read by the pre-create loader in that order.
The record order is **not** the panel order; the panel reads `0x138, 0x13b, 0x139, 0x13a`.

Effect kinds 3 and 4 write Mind at `+0x88` and Reaction at `+0x86`.
The object serializer sends one byte per attribute: stream `+0x28` Body,
`+0x29` Mind, `+0x2a` Spirit, `+0x2b` Reaction; a received zero skips that
attribute's update. The record writer copies actor `+0x84/+0x88/+0x8a/+0x86`
to record `+0x138/+0x139/+0x13a/+0x13b`. Wire storage is byte-wide,
independently of the derive caps. — HERO-STAT-001

Magic consumers of Mind and Spirit are specified in [MAGIC](../magic/format.md). Outside
magic: Mind buys sight and multiplies experience gain; Spirit buys the mana maximum and the five
elemental protections. Neither is read by movement, by the hit resolver or by regeneration.


## Character generation

```
T(n) = ftol( 0.349 * pow(1.15, n - 1) + 0.5 )        the CUMULATIVE cost of one stat at n

start        all four stats = 25, pool = 100
"+" on v     refuse if pool < T(v+1) - T(v)  or  v >= 45
             else  v += 1 ;  pool -= T(v_old+1) - T(v_old)
"-" on v     refuse if v <= 15
             else  v -= 1 ;  pool += T(v_old) - T(v_old-1)
budget       accepted iff  T(body) + T(reaction) + T(mind) + T(spirit) <= 140
             otherwise all four are forced back to 25
identity     pool == 140 - sum T(stat)   (because 4*T(25) + 100 == 140)
```

The refund is **exactly symmetric**: `refund(v) == cost(v-1)` for every `v` in `[16..45]`.
The counter drawn on the creation panel is the **remaining** pool.

Reachability, given the budget: one stat alone reaches **42** with the others left at 25 and **43**
with the others floored at 15; all four together reach **34**. **The click bound 45 cannot be reached
in character generation** (`T(45) = 164`; `HERO-BUDGET-004`).

Nothing in the point-buy depends on class, race, or the other three stats, and no registry key or
class table participates. Character generation also sets exactly **one** skill slot — to `20` when the
global `[L03937] == 2`, otherwise `10` — zeroing slots 1..5 first.

### What character generation puts in his hand

`R0837(slot, value)` zeroes skill slots 1..5, writes the one chosen slot, recomputes, then
a comparison of the value with 0xa (`L03942`) selects one of two sets of five weapons, built from **string literals**
through `R0665`. `value` is 20 only when `[L03937] == 2` **and** `campaign+0x6bc == 2`
(`L04211`…`L04212`); otherwise 10 (`HERO-START-039`).

```
slot  skill      value <= 10 (ordinary)          value > 10
1     Blade      Iron Short Sword                Uncommon Steel Two Handed Sword
2     Axe        Uncommon Bronze Axe             Uncommon Steel Axe
3     Bludgen    Uncommon Bronze Mace            Uncommon Steel Mace
4     Pike       Bronze Pike                     Uncommon Steel Pike
5     Shooting   Uncommon Wood Short Bow         Uncommon Magic Wood Short Bow
mage  (staff)    Wood Staff {castSpell=Fire_Arrow:10}   ... :20
```

An unarmed hero has an all-zero equipment modifier and `active = 0`;
the default path does not construct the `BareHands` Weapons definition.
Both skill terms are skipped and the roll is `[d, 2d]`, where
`d = ftol(1.1^body/20)` is zero below Body 32. — HERO-BARE-037


## Class discriminator

`actor+0x4c` bit 2 is the **fighter/mage** flag, derived when the streamed mana maximum is positive
and also set by spellbook construction. `R0442` tests it. Sex is a **different** test:
`R0841` returns `typeID ∈ {0x22, 0x24}`. In non-zero player-character constructor mode the
overwrite is `typeID = gender + (mage ? 0x23 : 0x21)` — gender in the addend, class in the base.
Zero-mode map Humans retain their Data.bin typeID, so this sex predicate is not universal to the
Humans class (`PARTY-M20-031`). The derive calls only the class test; its two conditional edges are the
health and mana multipliers at steps 1 and 2. Outside it there are fifteen more, all in the effect
dispatch: twelve skill arms in two gated sets of six over the same six fields, plus `manaMax` and
`manaRegeneration`. The fighter predicate is the exact negation of the mage one; the polarity is
fixed by the health multiplier, which doubles when the bit is **clear**. On the UI side the same axis is `record+0x18c` bit 1 (mage) and bit
2 (female), which pick the fighter/mag skill column, the `chrgen1m`/`chrgen1f` tips and the
`FacesMM`/`FacesMF`/`FacesFM`/`FacesFF` list.


## Name (`HERO-NAME-079`, `HERO-TYPED-080`)

`actor+0x80` is an MFC `CString` and it is the actor's name. It is constructed and destroyed by
every actor constructor and by the destructor-shaped sixth, so a monster carries one too; it is
per **instance**, never per class. Four writers, and no shipped surface draws it for a non-hero:

| Writer | What it assigns |
|---|---|
| `R0656` `L04507` | the literal `"Unknown"` — only inside a branch no shipped asset reaches |
| `R0656` `L04509`/`L04510` | a default-name array element, same dead branch |
| `R0823` `L13000` | the `CString` at the accepting request's `+0x18` — the **typed** name |
| `R0151` `L04513` | the `.alm` spawner's own local, on the `Humans` sub-arm only |
| `R0907` `L04514` | the `Name` key of a `Humans.Hero` block that ships in no root |

Limits for a consumer: **10 bytes**, refused at the keystroke by the entry control
(`TEXT-NAMEIN-024`, whose vtable clauses are partially retracted; the field only appends,
`TEXT-075`); charset as `TEXT-COLL-025` bounds it; storage is a `CString`, so there is no
fixed buffer to overflow; and the field is **not inert** — `R0508` skips any actor whose
name is empty before evaluating its sex/class predicate, so an empty name changes behaviour.

The two default-name arrays — 10 and 7 slots at `L07956` and `L07955`, holding **9** male
and **6** female names with slot 0 empty — are `.rdata` literals built by two static initialisers,
so they do not localise; and the gate that reads them requires a template name beginning `Hero`,
which nothing in either install provides.


## Character-generation controls (`HERO-CHARGEN-082`…`HERO-CHARGEN-085`)

Character generation is two screens and two boundaries. Pre-create owns name, class/sex and
difficulty. Its name field opens at `npcnames.txt` entry 20, EN `Danath`, RU `Данас`
(`TEXT-073`), and a hero press rewrites the name only while it is a default (`TEXT-074`). Its
Forward copies those fields into the main-frame draft, constructs a temporary
archetype record and opens detailed generation. Detailed Play joins or resolves the participant and
only then sends command `0x48`, which creates and installs the live actor.

Reset on the detailed screen writes pool 100 and Body/Reaction/Mind/Spirit 25. It preserves the
selected skill index, name, class/sex, difficulty and appearance. Preservation of the skill selector
does not mean preservation of the preview record: the reset reconstructs a fresh archetype actor,
zeros the five class skills, restores the chosen starting skill, recomputes all derived mirrors and
releases/rebuilds the twelve equipment-object slots (`HERO-CHARGEN-082`).

Detailed Back returns to pre-create and copies name, difficulty and the class/sex bits (`main+0x494 >> 6`) back into it. The
pre-create enter that runs next resets class/sex to the first hero and turns a default name into
`npcnames.txt` entry 20 (`TEXT-073`), so only the difficulty and a typed name reach the screen; the copied class/sex is never read (`HERO-CHARGEN-083`, amended).
Stats, skill and appearance are not written at that moment. The next pre-create Forward nevertheless reloads the
archetype defaults unconditionally, so stat and skill edits do not survive the round trip. Pre-create
Back is the destructive cancel: it releases the temporary record and returns to the participant/lobby
screen (`HERO-CHARGEN-083`).

The name field's UI limit remains ten stored bytes. Final Play adds server gates in this order:
non-empty; not the whole string `Self` or `Computer` under an ASCII-case-insensitive comparison; not
a byte-exact name already held by another participant; and no second character on a returning
participant. The first three failures are localized string-table indices 193, 194 and 195
(`HERO-CHARGEN-084`). The ten-byte limit is UI policy, not a server-storage limit: the participant
and actor both use `CString`. Lifting it changes no shipped asset bytes, but a consumer must also
replace the fixed preview-name copy rather than widening the control alone.

A failed Play is not transactional. It closes the detailed screen, sets the ready bit, resets the
connection and releases the temporary preview before the server rejects the name. It creates neither
a participant nor a live actor. Successful `0x48` is the commit point and copies `Player+0x18` to
`actor+0x80` unchanged (`HERO-CHARGEN-085`). Back/Play close is latched and idempotent while the screen
is inactive; repeated Reset reconstructs the same draft.

## Starting templates and non-item state

The four class/sex combinations select `PC_Danath`, `PC_Naira`, `PC_Fergard`, or `PC_Reniesta`.
Their ten `Humans` equipment cells can construct only Weapon, Shield, and Armor objects because the
cell position selects the class. Character generation adds one further direct weapon. It has no
MagicItems construction arm. No `Humans` equipment cell in either preserved root contains `Quest`.

After those paths, live-hero construction calls `R0993`. When `server+0x0c == 0` and
`[L03937] != 2`, the helper constructs `Quest Documents` by row name and appends it to the hero's
carried container before runtime-id assignment. The campaign construction path establishes the
first gate. The second gate was not measured in a fresh campaign, so actual item presence remains
Unknown (`ITEM-DOC-069`).

The character-generation command carries no purse value. The selected template, stats, appearance,
name, and weapon do not change the owning `Player` purse (`HERO-START-081`).
