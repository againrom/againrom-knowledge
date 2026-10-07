<a id="ai--how-a-unit-chooses-whom-to-fight-and-what-it-does-when-nobody-is-ordering-it-partial"></a>

# Actor AI, targeting and orders

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

The full-tick group pass dispatches its current order, then applies withdrawal
tests. Player commands can allocate a new order 0 group; map scripts can
replace the load-time guard/aggressive order. Movement executes the selected
route and reservations through its own clock. See the linked state tables
for each gate and for field-specific save restoration.
— AI-TICK-008, AI-GUARD-007, AI-ORDER-031

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
