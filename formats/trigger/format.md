<a id="trigger--the-mission-script-runtime--specification"></a>

# Mission trigger runtime

The runtime compiles ALM nodes into checks, actions, trigger conditions and a
shared register file. Every check is evaluated; an action runs only when a
surviving trigger names it. The one-second pass, fire-once state and reference
binder define script progression (`TRIG-CLOSURE-037` is amended; `TRIG-TAKEITEM-038` supersedes its
completeness clause). — TRIG-COND-003, TRIG-ACT-004,
TRIG-CMP-006, TRIG-FIRE-007, TRIG-CLOSURE-037

The reference covers the installed campaign's reachable operation vocabulary.
Dormant arms, catalogue-only subcommand 18 and authored but unreferenced
actions do not inherit that coverage. — TRIG-INSTCENSUS-046

Stored node layouts are in [ALM](../alm/format.md); clock/pool execution is in
[SESSION](../session/format.md); persistent state is in [SAV](../sav/format.md).

## Compiled state and pass order

| Session member | State |
|---|---|
| +0xc2ac / +0xc2b0 / +0xc2b4 | Check, instant and pattern collections (80/72/24-byte records) |
| +0xbd34 | 100 integer result/variable slots |
| +0xbec4 | 1000 byte latches, indexed by pattern position |

The full-tick path increments slot 93, evaluates checks, then evaluates
patterns. A pattern ANDs up to three comparisons; on success it sets its
latch and executes up to four actions in slot order. The once flag retains
the latch or clears it for the next pass. Winning and losing are actions
that update outcome state. — TRIG-EVAL-001, TRIG-COND-003, TRIG-ACT-004,
TRIG-CMP-006, TRIG-FIRE-007

## Reference map

| Reference | Contents |
|---|---|
| <a id="the-one-thing-a-consumer-must-not-get-wrong"></a><a id="the-machine"></a><a id="the-pass"></a><a id="what-must-be-in-a-save"></a><a id="how-a-nodes-ten-authored-slots-reach-an-arm"></a><a id="state-and-identity"></a><a id="compiled-state"></a><a id="evaluation-pass"></a><a id="firing-discipline"></a><a id="persistent-state"></a><a id="mission-end"></a><a id="parameter-binding"></a> [Compilation, evaluation and persistence](runtime.md) | Compiler stores, evaluation, latches and save/load order |
| <a id="the-vocabularies"></a><a id="operation-vocabulary"></a><a id="the-three-item-arms-trig-takeitem-038-trig-xferitem-039-trig-itemtest-040"></a><a id="the-campaign-reachable-closure-arms-trig-closure-037"></a><a id="closure-boundary"></a> [Operation vocabulary](operations.md) | Check/action parameters and operation-specific effects |
| <a id="identifier-bands"></a><a id="the-drop-table--where-the-player-lands"></a><a id="starting-the-mission"></a><a id="ending-it"></a><a id="the-register-file-is-shared"></a><a id="mission-text"></a><a id="worldmissionini"></a> [Mission identifiers and entry/exit](mission.md) | Mission identity, map binding and win/lose transitions |
| <a id="map-presence--instants-16-17-18-32-33-trig-offmap-041trig-mapgroup-043"></a><a id="the-cell-record-instant-29-reaches"></a><a id="action-6-the-three-member-helpers-and-the-gate-on-subcommand-10"></a><a id="instants-19-and-22-the-ownership-move"></a><a id="the-two-population-checks-count-the-living-and-neither-tests-health"></a> [Map presence and group operations](groups.md) | Off-map state, group helpers, ownership and live population counts |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.

The selected Trigger-binding block addresses three condition-ID pairs from
record+0x80, four action IDs from+0x98, three comparisons from+0xa8 and the
once word at+0xb4. It publishes separate compiled records. No direct access
to the opaque+0x40..0x7f span is demonstrated in this block; forwarded key
pointers, unexpanded calls and other aliases prevent a global no-use claim.
— TRIG-BIND-063
