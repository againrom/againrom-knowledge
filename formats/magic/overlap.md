# Fire Wall and Poison Cloud overlap

[Reference](format.md)

## Fire Wall and Poison Cloud overlap

An area object's clock, its cell-layer ownership and a target's Poison attachment
have separate lifetimes. `MAGIC-CLOUDCLOCK-154` corrects the cloud clause of
`MAGIC-AREAPULSE-037` (amended and partially retracted in the ledger): the registration call paints without pulsing. Subsequent
calls first test the old counter; if positive, they decrement and pulse when the
new value is divisible by 16, including zero. Cleanup occurs on the next call.
With unmodified positive initial counter `V`, pulse entries occur at
`(V-1)%16+1`, then every 16 through `V`; count is `ceil(V/16)`, teardown is at
relative tick `V+1`, and registration through teardown takes `V+2` inclusive
calls. Thus `V=240` has 15 pulse entries. These counts do not count HP events.

Each area remains in the original area list independently. A new same-spell
area overwrites the cell's single layer pointer without cancelling the older
object. A pulse tests whether that spell layer is present, then applies the
ticking object's own payload. It does not require the slot to point back to
that object. Its radius-square scan can therefore reach another area's painted
cell, even outside its own painted pattern. Cleanup requires identity with the
stored pointer; removing its owner clears the gate and does not reveal an older
area pointer. An older object's counter may continue without any applications
at the cleared cell (`MAGIC-CLOUDOWNER-155`).

The pulse visits dx in the outer loop and dy in the inner, both from `-radius`
through `radius`. Each admitted cell reads only occupant `+4`. There is no set
of already-visited target identities on this path. Fire Wall invokes direct
damage separately per reached occupied cell; the same actor pointer in several
cells can receive several applications. Neither Fire Wall nor Poison Cloud
takes the division by target footprint reserved for Fireball id 2. The number
of reached cells still depends on placement and painted layers
(`MAGIC-CLOUDVISIT-156`, `MAGIC-FIREDIV-047` (amended and partially retracted in the ledger), `MAGIC-AREAAPPLY-038` (amended and partially retracted in the ledger)). Fire's
per-application payload and resolver remain `MAGIC-DMG-005` and
`MAGIC-RESIST-006`; overlap does not define one fixed damage total.

Poison takes the ordinary timed-Effect path. The selected valid effect text
conditionally parses to kind 6, continuous mode 2, magnitude −2 and counter
128. Its arm scales magnitude by the established power factor and stamps Token
id 8 while retaining the parsed counter (`MAGIC-EFFECT-015`,
`MAGIC-POISONINPUT-157`). First attachment copies kind, mode, id, magnitude,
counter and source, applies the copy immediately, then appends it. A later
continuous same-id application writes only the retained counter. It neither
replaces magnitude/source nor applies HP immediately in that refresh branch,
whether the incoming area has another power, producer or occupied cell
(`MAGIC-POISONREFRESH-158`).

The named world driver calls areas before actors. The actor processes its
attachments before its later health/action logic, except terminal act `0x10`.
Continuous application tests the old remaining counter modulo 8 before
decrement; it is not an independent global eight-tick timer. A newly attached
or refreshed counter of 128 is eligible in the later actor phase of the same
world tick. First attachment can therefore cause an immediate area-phase HP
event followed by an actor-phase HP event. Staggered area refreshes can change
that phase; unrefreshed application has eight-tick spacing
(`MAGIC-ATTACH-016` (amended, partially retracted and superseded in part in the ledger), `MAGIC-POISONPHASE-159`). Poison's nonzero HP branch uses
the separate resistance arithmetic and notification gate in `ANIM-074`.

Fire Wall removes the Poison layer when it paints a conflicting cell. Poison
painting into Fire removes its own layer. Both placement orders leave Fire
admitted at that cell. The conflict does not cancel either whole area object
and does not revoke Poison already attached to an actor. Such an attachment
continues its own counter after the area layer disappears
(`MAGIC-FIREPOISON-160`). The other spell-pair rules retain
`MAGIC-MAPLAYER-040`'s existing scope and confidence.

The retained attachment source can also differ from later area credit. Subject
to the existing domain/source/owner gates, an area application writes its
producer into target `+0x40` and inner Token id into `+0x48`, even when that
application merely refreshed a Poison whose source `+0x44` remains the first
producer. Later Poison HP application still consults that retained source
(`MAGIC-AREASOURCE-161`). Eventual training and kill attribution retain the
gates and limits of `MAGIC-ITEMTRAIN-116`, `MAGIC-ITEMKILL-117` (amended and partially retracted in the ledger) and
`MAGIC-ATTRGATE-118`.

`MAGIC-CLOUDEND-163` qualifies `MAGIC-AREAEND-041` (amended and partially retracted in the ledger): identity-gated cell removal
recounts and recomputes layers. If the cell record becomes empty, the complete
helper restores its stored terrain and flag bytes before removing the record.
This restoration does not reinstate prior area ownership or revoke a target's
separate Poison attachment.

The local branches and stores above are High. `MAGIC-AREAOVERLAP-162` keeps
the joined event streams Medium: they condition on stationary, nonterminal
targets, valid producers, library substitutes and a controlled Fire resolver.
The selected powers 0 and 30 include a disclosed zero-baseline deviation from
the intended positive-power sample. Native RNG totals, arbitrary geometry,
source destruction, terminal transitions and other scheduling populations
remain Unknown. No native gameplay acceptance follows from these conditional
instruction traces.
