<a id="ai--how-a-unit-chooses-whom-to-fight-and-what-it-does-when-nobody-is-ordering-it-partial"></a>

# Actor AI, targeting and orders

<!-- machine-page -->

Target acquisition reads actor state, diplomacy, candidate populations and
sight. Group and per-actor state machines select orders; movement executes
those orders. Player input creates command records through separate handlers.
— AI-FILTER-001, AI-GUARD-007, AI-TICK-008, AI-DIFF-016

Saved ownership and list order are inputs to continuation. Do not replace
field-specific restoration with a single generic Group reset.

## State and evaluation order

| Input/state | Rule |
|---|---|
| Actor owner, health, flags and cell | Gate acquisition and candidate eligibility |
| Diplomacy and visibility population | Select eligible targets |
| Group order | Dispatch the group behavior; order 0 delegates to member states |
| Actor order at actor+0x158 | Holds the member action and its subject/target |
| Caster with no group-assigned target | Heal itself or the first kept damaged unit, else may buff one unit of its own Player — AI-784, AI-785, AI-786, AI-787, AI-788, AI-789 |

A computer caster's cast takes the target its spell arm fixes; no count of units on the spell's cells is read (`AI-796`, Medium), and the session-list cast routine that selects a unit on a cell is a separate routine (the list entries are not decoded; that it is not a computer unit's cast is Medium, `AI-797`).

The full-tick group pass dispatches its current order, then applies withdrawal
tests. Player commands can allocate a new order 0 group; map scripts can
replace the load-time guard/aggressive order. Movement executes the selected
route and reservations through its own clock. See the linked state tables
for each gate and for field-specific save restoration. The engage routine's
refusal for a short-reach spellcaster of a human participant is a tail call of the
acquire routine, not a no-op.
— AI-TICK-008, AI-GUARD-007, AI-ORDER-031, AI-886

## Reference map

| Reference | Contents |
|---|---|
| <a id="structure-use-command"></a><a id="player-input-contract-ai-input-121-ai-select-122-ai-panel-123-ai-minimap-124-ai-key-125-ai-cursor-126-ai-input-127"></a><a id="map-mouse"></a><a id="minimap-mouse"></a><a id="command-cells-and-keys"></a><a id="explicit-retreat-execution"></a><a id="quick-slot-state-and-lifecycle"></a><a id="book-identity-and-selection-predicates"></a> [Player input and commands](input.md) | Input edges, command bytes, inventory and quick slots |
| <a id="what-this-covers-and-what-owns-the-rest"></a><a id="the-objects"></a><a id="saved-formation-ownership"></a><a id="saved-group-continuation"></a><a id="scope"></a><a id="object-state"></a> [Saved ownership and Group state](state.md) | Ownership, group lists and restoration |
| <a id="the-decision-in-the-order-the-engine-takes-it"></a><a id="target-selection-sequence"></a><a id="prismatic-spray-is-a-distinct-consumer-of-the-group-population-ai-spray-266-ai-spray-267"></a><a id="the-sight-stamp-complete-ai-groupsee-068-ai-los-087ai-los-091-ai-sight-092ai-sight-094"></a> [Target acquisition and sight](targeting.md) | Candidate filters, scorer and sight budget |
| <a id="the-diplomacy-matrix"></a><a id="diplomacy-matrix"></a><a id="the-six-writers"></a><a id="what-a-change-reaches"></a><a id="what-the-script-can-do-with-it"></a><a id="persistence-and-the-wire"></a> [Diplomacy](diplomacy.md) | Relation lookup and changes |
| <a id="who-runs-it-and-how-often"></a><a id="two-levels-of-behaviour"></a><a id="the-group-engagement-machine--how-a-candidate-becomes-a-target"></a><a id="guard--the-whole-of-motion-without-an-order"></a><a id="the-third-radius"></a><a id="update-cadence"></a><a id="group-and-actor-behavior"></a><a id="group-engagement"></a><a id="guard"></a><a id="patrol"></a><a id="engagement-radius"></a> [Group behavior, guard and patrol](behavior.md) | Group/member dispatch, guard and patrol |
| <a id="the-players-order-vocabulary"></a><a id="formation--a-per-player-mode-and-the-only-thing-that-makes-a-group-move-as-a-group"></a><a id="order-progress-and-how-an-order-becomes-an-act"></a><a id="player-orders"></a><a id="formation"></a><a id="order-execution"></a> [Orders, formation and execution](orders.md) | Player/script orders and formation |
| <a id="where-a-behaviour-is-chosen"></a><a id="withdrawing--what-a-ranged-monster-does-when-you-close-to-melee"></a><a id="behavior-selection"></a><a id="at-load"></a><a id="at-run-time--the-maps-own-script"></a><a id="which-break-off-rule-that-puts-in-front-of-a-player"></a><a id="ranged-withdrawal"></a><a id="difficulty"></a><a id="cost"></a> [Behavior selection and difficulty](policy.md) | Load-time stances, script overrides and retreat |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.

## Heal list membership and Heal at a hostile unit

The Heal candidate list can hold the caster (Medium), which is never the pick at full health. — AI-800

A hero mage's book Heal at a hostile unit has paid mana and, under the prelude's gate, trains the school; the Heal arm then skips the heal (Medium: diplomacy polarity and admission target tests unread). — MAGIC-465

Each of the five `R0020` calls in the six cast-chain windows centres on the actor's own cell (AI-862; AI-796, Medium); the Heal list's cell lookup returns slot `+4`, the actor of movement domain 1 or 2 (AI-866; AI-800, Medium); that list is built by bodies that load none of the six health and state offsets searched (AI-872, Medium).

## Group order arms

The dispatcher maps each group order to one arm routine (AI-908). The Guard arm builds the candidate list, clips it, assigns a target to every member and then walks the members, in a fixed order (AI-905, AI-913, AI-909); arms 2, 4, 5 and 0x11 are read (AI-910). The engagement setter stores state 5 on its engagement path and has 14 direct call sites (AI-907). The follow, acquisition and heal helpers store the pending kinds 4, 6, 0xb, 0 and 8 (AI-911). The AI driver has two direct call sites, and the Guard counter counts Guard-arm passes since the last latch flip, not session ticks (AI-904, AI-TICK-008). The sight radius byte has a second direct store (AI-903) and the Guard arm reads it per member through the geometry routine (AI-906). The one immediate store of 0x17 to an actor state word in the AI module's range is in a routine with no direct caller (AI-912).

A ground (domain 1 or 2) caster placed through the occupy routine is stored at slot `+4` of its base cell, and the block builder centres on that cell, so the list holds a placed ground caster when the plane byte of its cell has bit 5 set (`R0035`, AI-866; the bit's meaning is unread) and the occupy did not stop at a taken cell (AI-697) (AI-923, AI-800; Medium: the Position word and bytes are equal in saved states only). The area-object registration only appends a node (AI-922) and `R0184` holds one `0x3c` store (AI-924).

Cast-entry and temporary-caster reads: the gap cast entries `L14677`, `L18397`, `L14681`, `L15398` and `L14679` have no direct call or jump site and no virtual-address dword (AI-925, AI-570, AI-562); `R0289` is a Position store and the six `R0050` sites lie in four named routines and two movement step functions (AI-926); the tail of `R0184` is decoded linearly from `L18395` to the first return `L18396`, the class-arm table targets (UNIT-167) inside it, and stores nothing at `+0x3c` (AI-927); the progress 2 executor arm holds no call and ends in the tail (AI-928); each is Medium for reach and for unread callees.
