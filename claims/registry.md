# Claim registry

This file indexes the claim ledgers and defines how a claim is written. The
detailed research chronology, confidence-review notes, experiment write-ups,
disassembly and raw evidence live in this private research repository, indexed
by the research `docs/REGISTRY-LOG.md`; they are not part of the
public functional edition.

A claim ID is permanent. `go run ./tools/claim <ID>` prints one claim with its
retraction entries; `-l <ledger>` lists a ledger's headlines and `-k <regexp>`
searches every ledger. Corrections and narrowings are recorded in
[`retracted.md`](retracted.md).

## ROM1 ledgers

| Area | Ledger |
|---|---|
| AI and targeting | [`ai.md`](ai.md) |
| ALM scenario/map data | [`alm.md`](alm.md) |
| Animation | [`anim.md`](anim.md) |
| Data.bin | [`databin.md`](databin.md) |
| Dialogue | [`dialogue.md`](dialogue.md) |
| Hall of fame | [`fame.md`](fame.md) |
| Heroes and character generation | [`hero.md`](hero.md) |
| Inventory | [`inv.md`](inv.md) |
| Items and equipment | [`item.md`](item.md) |
| Magic | [`magic.md`](magic.md) |
| Menus | [`menu.md`](menu.md) |
| Mission flow | [`mission.md`](mission.md) |
| Movement | [`move.md`](move.md) |
| Palettes | [`pal.md`](pal.md) |
| Party | [`party.md`](party.md) |
| Inline registry | [`reg.md`](reg.md) |
| RES/LM container | [`res.md`](res.md) |
| Saves | [`sav.md`](sav.md) |
| Session/campaign state | [`session.md`](session.md) |
| Shops | [`shop.md`](shop.md) |
| 16/16A sprites | [`spr16a.md`](spr16a.md) |
| 256 sprites | [`spr256.md`](spr256.md) |
| Tavern | [`tavern.md`](tavern.md) |
| Terrain and map presentation | [`terrain.md`](terrain.md) |
| Text and encoding | [`text.md`](text.md) |
| Town UI and flow | [`town.md`](town.md) |
| Triggers | [`trigger.md`](trigger.md) |
| Units and combat | [`unit.md`](unit.md) |
| Audio/video | [`video.md`](video.md) |

## ROM2 ledgers

| Area | Ledger |
|---|---|
| Asset layout and identity survey | [`rom2-asset.md`](rom2-asset.md) |
| Engine-level observations | [`rom2-engine.md`](rom2-engine.md) |
| Session/network observations, save and character files | [`rom2-session.md`](rom2-session.md) |

`claims/rom2-asset.md` carries the functional authorities for the six asset-layout
surfaces in the seven-row ROM2 survey; the session-header authority is in
`rom2-session.md`. The original public snapshot `k1` omitted the asset ledger.
Each public snapshot's `SOURCE.md` identifies its contents and source commit;
private experiment links preserve evidence identity without exporting evidence.

## How a claim is written

A ledger opens with its scope and any terms its claims share, then groups
claims under topic headings. Each topic holds an index table and one card per
claim, in the same order.

```markdown
  ## Topic

  | ID | Claim | Confidence | Status | Evidence |
  |---|---|---|---|---|
  | AREA-TOPIC-NNN | One sentence a reader can build on, scope included. | High | ● active | [EXP-NNNN](../experiments/EXP-NNNN-slug/) |

  ### AREA-TOPIC-NNN

  The facts behind the headline: offsets, sizes, addresses, counts and the
  population measured, as short sentences or a list.

  **Confidence.** What rules out the live alternatives; a separate grade for each
  clause that differs from the cell.

  **Unknown.** What the evidence leaves open.

  **Amended.** Which clause a later claim or retraction changed, and where.
```

- **Claim** is one sentence of at most 240 characters. It carries the scope
  when the fact holds only for a searched population, a locale or a code path.
- **Confidence** is a grade only: `High`, `Medium`, `Low` or `Unknown`. When
  clauses differ, each grade that applies appears once, in that order, joined by
  ` / `; the card's **Confidence.** paragraph says which clause has which.
  A retracted claim with no standing grade carries `—`.
  Confidence records the strength and scope of the research evidence. It is not
  a legal clearance, a compatibility guarantee or permission to redistribute
  original game material.
- **Status** is one marker, optionally followed by qualifiers in parentheses,
  separated by commas:

  | marker | meaning |
  |---|---|
  | `✔ promoted` | a format page also states the claim |
  | `● active` | current, not stated by a format page |
  | `✖ retracted` | the whole claim is withdrawn; [`retracted.md`](retracted.md) holds the correction |

  | qualifier | meaning |
  |---|---|
  | `amended` | a later claim extended, corrected or narrowed a clause |
  | `partially retracted` | a clause is withdrawn; the rest stands |
  | `superseded` | a later claim replaces a clause |
  | `contested` | two readings stand; the card names both |
  | `branch candidate` | published from a bounded branch candidate whose confidence review had not passed |

  `amended`, `partially retracted`, `superseded` and `contested` require an
  **Amended.** paragraph in the card, and an **Amended.** paragraph requires one
  of them unless the claim is `✖ retracted`. A claim with an entry in
  [`retracted.md`](retracted.md) carries a qualifier; a `SUPERSEDED` entry
  requires `superseded`, and a `REFUTED` or `RETRACTED` entry requires
  `partially retracted` unless the whole claim is `✖ retracted`.
- **Evidence** links the experiments whose evidence earned the claim.
- A card stays under 8 KB; `tools/claim -check` refuses a longer one.
- The card holds facts, not their history. Git and the experiment record carry
  how a fact was found; another claim is cited by its ID rather than restated.
  **Confidence.**, **Unknown.** and **Amended.** appear only when they have
  content.

## Publication boundary

The public ledgers should state independently expressed functional conclusions,
confidence, scope, status and the identity of the private evidence that earned them.
Instruction listings, decompiler output, reconstructable shipped-content tables and
internal review narrative are publication-review material under the knowledge
repository's `PUBLICATION.md`, not required parts of this index.

## SFX spatial arithmetic amendment

VIDEO-SFX-088 through VIDEO-SFX-091 establish the named positional and ambient
input/arithmetic contracts. The log10 arithmetic label of ANIM-SND-022 is
partially retracted; read its amendment and overturn entry before using the
former wording. Other clauses and their confidence remain unchanged.

## Drawable registration scope

ANIM-REGISTER-083 through ANIM-WALKORDER-088 distinguish ordinary, air and
alternate CUnit storage and dispatch phases. The affected clauses of
TERR-STRUCT-104, TERR-SPR-065, TERR-SPR-137, TERR-SPR-138, TERR-SPR-139,
REG-UNITS-061 and ANIM-047 are amended; read their retraction entries before
using earlier all-unit or common-bound shorthand.
