<a id="shop--stock-price-and-trade--specification-partial"></a>

# Shop stock, prices and trade

A shop holds generated stock, a shared five-place trade table and per-shop
price bounds. Campaign state selects the town shop; multiplayer placed shops
use their own refresh path. Stock generation, basket totals and final payment
have distinct rounding and lifecycle rules. The stocked-item label for
Template+0x04 in SHOP-OBJ-002 is retracted: it counts open customers.
— SHOP-CLS-001, SHOP-OBJ-002,
SHOP-LIFE-013, SHOP-ROUND-017

Partial-stack selection and the exact clock/draw position at the first shelf
after mission or save load remain Unknown. The stock clock uses the session
tick counter described in [SESSION](../session/format.md).

## Shared and per-customer state

| Runtime object/member | Meaning |
|---|---|
| Shop+0x6c / +0x70 | Template pointer / item-value cap |
| Template+0x04 / +0x08 | Open-customer count / pending-restock counter |
| Template+0x0c | Four shelves, codes 1,2,4,3 |
| Instance+0x74 / +0x78 | Customer / tray list |

Resolve the campaign static shop or multiplayer placement, apply its cap,
then generate and expose stock under the customer/restock gates. Trade moves
objects onto the shared table, computes basket prices, checks payment and
commits or returns items through their distinct load-update paths.
— SHOP-ENTRY-003, SHOP-CAP-004, SHOP-ENTRY-016

## Reference map

| Reference | Contents |
|---|---|
| <a id="at-a-glance"></a><a id="structure"></a><a id="actor-state-across-trade-boundaries"></a><a id="where-a-shop-comes-from"></a><a id="the-one-per-shop-parameter"></a><a id="the-campaign-chain"></a><a id="when-the-ceiling-is-applied"></a><a id="the-stock-lifecycle"></a><a id="persistence"></a><a id="the-candidate-pool"></a><a id="the-draw"></a><a id="the-enchantment-stage"></a><a id="randomness"></a> [Stock and generation](stock.md) | Providers, cap, candidate pool and enchantment generation |
| <a id="trading"></a><a id="where-the-price-comes-from-and-every-rounding"></a><a id="which-shelf-an-item-returns-to"></a><a id="the-counter--the-table-the-two-buttons-the-move"></a><a id="shop-interior-progression"></a><a id="the-character-panel-0x7c"></a><a id="duplicates"></a> [Trading and shop interface](trade.md) | Prices, tray movement, payment, refusal and return |
| <a id="consumer-notes"></a><a id="the-runtime-class-table"></a><a id="state-requirements"></a><a id="document-access-item-exclusion"></a> [Classes and state requirements](state.md) | Classes, persistence and retained state |

Runtime member offsets identify original in-memory fields. Wire offsets and
byte order are stated separately in the relevant layouts.
