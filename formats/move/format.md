<a id="move--unit-movement-and-path-selection--specification-partial"></a>

# Movement, path search and cell transitions

A mover combines its domain, footprint, terrain/cell state and order target.
Search produces a route; reservation, step and refresh rules execute it on
the sub-tick clock. Group movement supplies a separately gated speed term.
— MOVE-SEARCH-001, MOVE-REFRESH-012, MOVE-TICK-013, MOVE-TICK-017,
MOVE-RATE-029, MOVE-GROUP-037

Movement does not inspect Building HP. If lethal damage leaves the Building
cell reference and masks unchanged, the next predicate has the same inputs.
The actual lifetime of those inputs to that next consumer remains Unknown;
ruin appearance alone does not establish passage release.
— UNIT-STRUCTNEXT-079, UNIT-STRUCTBOUND-081

Unknowns include R0058's cell-boundary effect, the complete blocked-cell
verdict table, exact contact-ring entry cell, complete order-layer producers
and Building tick-container ownership.

## Movement inputs and sequence

| Input/state | Use |
|---|---|
| Domain and footprint | Candidate-cell admission |
| Terrain costs and cell masks | Search labels and blocked-cell tests |
| Target, route and reservation | Next-cell selection and ownership |
| Speed and group/formation gates | Whole-number step duration |

Search labels the allowed cells, extracts the route and selects its next
step. Execution reserves and enters cells, advances within-cell position on
sub-ticks and refreshes or releases state through the named transition
helpers. Diagonal and straight costs, stopping conditions and failure paths
are explicit in the linked procedures. — MOVE-SEARCH-001, MOVE-REFRESH-012,
MOVE-TICK-013, MOVE-RATE-029

## Reference map

| Reference | Contents |
|---|---|
| <a id="at-a-glance"></a><a id="the-movement-domain-move-dom-024028"></a><a id="the-search--R0053"></a><a id="the-substitute-goal-move-alt-018022"></a><a id="the-route--R0438-dynamic--R0437-static"></a><a id="structure"></a><a id="movement-domain-move-dom-024028"></a><a id="path-search--R0053"></a><a id="substitute-goals-move-alt-018022"></a><a id="route-extraction--R0438-dynamic--R0437-static"></a><a id="parameters--datamapreg-path-finding"></a> [Domains, path search and route extraction](search.md) | Admission, costs, search, termination and reconstruction |
| <a id="the-tick-order--what-first-means-move-tick-013017-move-id-016"></a><a id="the-step"></a><a id="coordination-between-units"></a><a id="tick-ordering-move-tick-013017-move-id-016"></a><a id="movement-step"></a><a id="what-the-cell-boundary-calls-do-terr-cellrec-146-terr-footprint-147"></a><a id="restored-actor-registration-and-the-reached-turn"></a> [Coordination, ticks and cell transitions](execution.md) | Reservations, refresh and cell entry/exit |
| <a id="the-rate-and-the-clock-it-runs-against-move-rate-029034"></a><a id="what-an-area-effect-can-do-to-movement-move-area-038"></a><a id="what-stops-a-mover-before-the-search-runs-move-gate-039-move-step-040"></a><a id="movement-rate-and-clock-move-rate-029034"></a><a id="when-the-group-term-is-set-and-when-it-is-not-move-gate-035037"></a><a id="area-effects-move-area-038"></a><a id="refresh-policy"></a><a id="pre-search-gates-move-gate-039-move-step-040"></a> [Movement rate and formation gates](rate.md) | Speed formula, timing quantization and group gates |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.
