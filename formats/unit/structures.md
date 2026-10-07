# Buildings, interaction and cell casters

## Building-associated type9 object

`UNIT-T9CTOR-110` describes the selected VirtualCaster construction. The
0x44-byte object has vtable L09291, a distinct Position at+10 and a6-byte
payload at+40. The caller copies the building Position and owner, writes
caster+3c from the low byte of the unlinked record tag, and writes payload0..3
from low8(spellRaw), byte2(spellRaw), low8(X), low8(Y). It appends the object
to the collection at builder.this+24 pointee+10. Neither authored A nor a
continuing building object pointer is stored by this construction chain.

Payload constructorL12555 supplies no initializer. Payload4..5 retain
allocation contents through this caller; no zero default or tail-coordinate
meaning is established there. This differs from the fully assigned cell-entry
payload described below. Later writers and native allocator contents remain
Unknown.

`UNIT-T9LIFE-111` supplies a local deletion relation. A nonnull candidate at
L12560 receives virtual+4(1) after the unlink helper, whose return is ignored.
On this caster table the call reaches its destructor, releases payload and
Position, zeroes their pointers and can delete the object allocation. The
selected caster virtual+14/+18 methods are empty; they establish no activation.

The session publishes global L05906 to+14 and+68; its cleanup reaches member+10.
The Building damage resolver has the conditional position-key cleanup already
bounded by `UNIT-STRUCTSTOP-066`. The builder's collection pointer has not been
proven identical to that global instance in the selected chain. First list
consumer, actual unlink, native callback/removal order and complete save/load
lifetime remain Unknown. A class serializer's presence is not a SAV parent.

[Reference](format.md)

## Clickable structures

AI-STRUCTUSE-306 schedules approach before the local use action. In that action,
selectors 28/29 toggle word+42 between zero and one, treating every nonzero
input as the zero-producing arm (UNIT-STRUCTUSE-090). Notification uses the
same client projection boundary as UNIT-STRUCTCONT-078.

Selectors15/16 spend one positive word+42 charge and construct, respectively,
Potion Big Healing or Potion Big Mana. For Human receivers the item goes
through the ordinary potion-use virtual before any unconsumed item is returned
to the container (UNIT-STRUCTUSE-091). ITEM-USE-113 and MAGIC-CONSUME-142 give
the existing application and clamped +100 effects. A full stat does not prevent
charge consumption. UNIT-STRUCTZERO-080 describes the separate local recharge
callback; its scheduler period in elapsed play time remains Unknown.

The action performs nearby-object lookups rather than using the retained order
pointer for the effect. Ambiguous neighbors, non-Human receivers, complete
scheduling and original-runtime completion remain outside these local claims.


## Scope and interaction

### Physical orders against a Building

The received object-target attack route can resolve a Building id and retain its pointer
through the shared order and approach machine; there is no separate structure action in
the read route. The scorer, caster alternatives and normal approach predicates still
apply. Successful eligibility of every structure class is not established
(`UNIT-STRUCTORDER-062`).

The base Building vtable returns token size 1, independently of its registered rectangle.
Approach checks facing and edge distance, while application uses the separately rounded
strike distance. Start computes range-dependent delay without rejecting range; application
does reject it (`UNIT-STRUCTREACH-063`).

For live combat block `A=attacker+0xa6`, Building damage is zero when the block is null,
maximum HP is zero, or `u8(A+0x14)==0`. Otherwise it is
`max(u8(A+0x13) + U[0,u8(A+0x14)] - 5, 0)`. Ordinary physical damage, to-hit, defence,
absorption and elemental-kind selection are not read by this resolver. Consequently an
accepted physical attack need not damage a structure (`UNIT-STRUCTDAMAGE-064`).

Both short- and long-reach physical presentation remain in the actor's countdown, with
distance extra at start. Physical application subtracts from word HP without a zero clamp;
the separate effect consumer clamps to zero. Retained unredirected orders inherit the
shared two boundary ticks as well as charge/recovery and their modifiers
(`UNIT-STRUCTDELIVERY-065`).

Target HP zero does not cancel the named strike body. The resolver's same-cell cleanup
before a lethal result is not the Building destructor. Actual destruction scheduling,
retained-pointer lifetime, obstruction release and mid-attack save/load remain Unknown
(`UNIT-STRUCTSTOP-066`).

The reached notification builds a Building state message (`0x82`) containing HP,
not a death-specific message. Its admitted client arm updates drawable HP; neither
local arm branches on HP sign. Packet virtuals, send and flush descendants remain
an unresolved execution boundary (`UNIT-STRUCTCONT-078`).

A separate `+0x14` callback increments HP for selector byte 15 or 16 when the
session counter is divisible by 60 and `0 <= HP < maximum`. Zero is admitted;
negative HP is not. Its live scheduling and membership are not established
(`UNIT-STRUCTZERO-080`). A retained cell reference and mask keep the same lookup
and movement interpretation regardless of HP (`UNIT-STRUCTNEXT-079`). The bounded
caller search does not establish whether those inputs survive an actual lethal
application; synchronous/deferred removal and permanent registration remain
alternatives (`UNIT-STRUCTBOUND-081`).

| | owner |
|---|---|
| the `Data.bin` file, its grammar and its schema law | [`formats/databin`](../databin/format.md) |
| the **human** arm — chargen, the derive `R0280`, the damage resolver | [`claims/hero.md`](../../claims/hero.md) |
| the client drawable's frame, phases and corpse stage | [`formats/anim`](../anim/format.md) |
| movement, pathing, tick order | [`formats/move`](../move/format.md) |
| **this file** — everything between a Units row and a monster that can hit you | — |


## Multi-cell Building area targets

The simulation object is Building, not the interface CStructure (`UNIT-CLASS-023`). Its presence
mask registers the same pointer in multiple cell records, in y-then-x order. A collision can leave
an attached prefix; the geometry alone is not the accepted-cell set (`UNIT-STRUCTCELL-070`).

Ring and blast visit cell slots, not unique objects. The ordinary cloud pulse omits the Building
slot. For a direct-damage inner effect, each successful lookup can therefore reach a fresh HP
application (`UNIT-AREAVISIT-071`, `UNIT-AREADIRECT-072`). Building's token-size getter returns 1,
so fireball's size-squared divide does not normalize its rectangle. Direct effect application
clamps negative post-subtraction HP to zero and does not reject current HP zero
(`UNIT-AREAHP-073`).

Destructor cleanup walks the mask again, not the accepted prefix, and clears a found `+0xc`
without an identity check. HP-to-destructor timing, live partial-registration teardown and alias
lifetime across real damage calls remain Unknown (`UNIT-STRUCTDETACH-074`). Installed EN/RU
payload and replay populations are explicitly separate (`UNIT-AREAPOP-075`).


## Mission-10 cell-entry caster

`UNIT-M10CELL-054` identifies the two original authored cells, (22,64) and
(21,63), with spell 13 and power 1. They contain source coordinates, not tower
references. Their type-9 source binding is in the ALM format page.

`UNIT-M10ENTRY-055` establishes the admission event: the square-footprint
attachment loop calls the cell helper for each cell. Movement domains 1/2 can
cast from an existing record with spell byte neither zero nor 26; domain 3
does not. The cast request precedes the occupied-ground-slot rejection. The
indexed spell predicate returns literal 1 in this image, so the target is the
entering actor, not the fallback current-cell target. This local arm checks no
allegiance, range, visibility or tower state.

The helper creates a runtime-id-zero actor at the source coordinate, with
Mind=30, skill[Sphere]=power, charge/recovery=1, order 13 and the entering actor
as target. It follows the ordinary spell-13 damage and direct-client effect
routes (`UNIT-M10CAST-056`). The visual flight value 5 is not a repeat timer.

The cell tail is not consumed or time-gated by admission. Another qualifying
attachment attempt may create another caster; standing still is not itself
this event. Zero spell disables this arm and 26 selects another operation.
The persisted cell tail and transient casting actor have different save
lifetimes. No tower-health dependency is present in this local path, but
post-destruction cleanup elsewhere and mission-10 runtime/save-load outcomes
remain Unknown (`UNIT-M10LIFE-057`).
