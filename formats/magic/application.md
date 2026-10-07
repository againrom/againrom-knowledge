# Application and ordinary Effect lifetime

[Reference](format.md)

## Effect ownership and credit

An attached ordinary Effect has source state at+44 separate from the target's
credited actor at+40. Default Effect construction clears its source, copy
preserves it, and continuous same-id refresh retains the old source. Template
construction's null-parser arm is different: it does not write+44, and its
ordinary reach/allocation value remains Unknown. — SAV-906

The new-Effect SAV route uses the default factory and omits source+44 from its
44-byte record. An archive back-reference preserves the loaded Effect alias.
The selected post-load paths do not reconstruct its source. In contrast, the
target's raw saved credit key is explicitly remapped on world LOAD, or cleared
on a lookup miss; the no-world document arm clears it. — SAV-907, SAV-908

For nonzero Token8 periodic damage, victim HP changes before source admission.
A null source skips the source callback, negative source HP clears the source,
and zero/positive HP admit virtual+64. A duration-only timer does not make that
periodic call. Later death credit reads target+40 and can name a different actor;
an intervening source+48 callback precedes the source+60 award. Native lifetime,
later repairs and first-consumer chronology remain Unknown. — SAV-909, SAV-910

## PointEffect target identity and caster non-persistence

`PointEffect`'s own+44 is the target Unit, not a source: the live constructor
writes its second argument straight into+44, and that field is the raw dword
`PointEffect::Serialize` stores/repairs (`SAV-CLASSSER-174`). Identity repair
happens once, inside `Serialize`'s own load arm; the class's separate post-load
hook never re-touches+44. SAV-1010's target-registration gloss is retracted:
the constructor's `L05872` call copies target Position into the effect's
Position; it is not list registration. The target store/repair remains valid.
— SAV-1010, SAV-1068

`SpellEffect`'s own+0x3c (caster) is zero-written by every construction,
including the one `Serialize`'s LOAD performs, and is not part of any
`Serialize` body in the class family. A `PointEffect` therefore always has a
null caster immediately after LOAD, independent of whether its own target
lookup hits or misses. — SAV-1011

The archive's own create/factory dispatch for that LOAD-path construction is
now traced rather than assumed: `CArchive::ReadObject` resolves the class
descriptor and calls through its own+0xC field, which, for all four lineage
classes, reaches the identical base constructor that zero-writes+0x3c — this
confirms, from the dispatch itself, that LOAD-path construction is one of the
two constructors already read zero-writing the field, for every class in the
lineage, not only the one candidate structurally guessed at before. The base
`Serialize` every concrete body in the family calls first also touches no
+0x3c at any offset in either its STORE or LOAD arm, closing the one level of
the inheritance chain the per-class `Serialize` reading above did not itself
check. — SAV-1054, SAV-1055

`PointEffect`'s and `AreaEffect`'s own per-tick attribution tails each
dereference the recorded caster's own pointee up to three times; the first,
feeding a "definition" test on the caster's own+0x3c (a different field on
the caster's own class, not this one), runs with no liveness check on the
caster performed first, while every later use of the caster in either tail is
preceded by an explicit caster+0x14 liveness check. The shared Tick driver
below already runs the identical check on this same field before either tail
is dispatched at all, so the unguarded read is reached with a stale caster
only inside the window between the driver's own check for a member and
that member's own Tick call completing; whether a caster can become invalid
inside it is Unknown. Neither tail writes+0x3c anywhere in its own body, and a direct-call
census of the two base constructors that zero-write it finds that every
direct caller is a construction path. No writer beyond the four already
known (the two base constructors, `Spell::Apply`, and the Tick driver's own
conditional clear above) was found in the functions read; no image-wide
write-site census of the field was run. — MAGIC-217, SAV-1056

The caster this "liveness" check reads is, per its own established naming, a
`Unit`/`Human` object in the general case, and its own `+0x14` is the same
field `Token`'s wire format carries on every class in this family — but at
least one cast path (the building-associated type9 arm) constructs a distinct
`VirtualCaster` object instead and copies its own `+0x14` from the source
Building's own `+0x14` at construction, so the caster's own concrete class is
Unknown in general, not established as `Unit`/`Human` by any cited row. A
runtime writer does zero an already-constructed object's own `+0x14`, inside
a removal/detach routine distinct from every construction path and from every
function the combat-death teardown traces; `+0x14` is therefore not
exclusively a construction-time default, and whether that runtime zero is the
same owning-Player field cleared on removal or a liveness-adjacent meaning is
Unknown — this does not discriminate between the two, and is graded Unknown,
not Medium. Across 70 admitted saves, every saved `Human`/`Unit` owner
reference is resolved, none zero; this is compatible with either meaning of
the constructed/removed zero, not itself discriminating. — MAGIC-219,
SAV-1058 (partially retracted only for the SpellEffect-to-Token constructor
identity; the `+14` stores cited here remain valid)

`PointEffect::Tick`'s attribution tail dereferences+44 at two unconditional
sites, with no null guard before either one: the entry dereference, and a
second one reached when the entry dereference's own virtual call returns
zero. A third read of+44, further into the same tail, is guarded by a null
check on a cached copy; that guard covers only that one downstream use, not
the two unconditional sites. The shared list registrar that both cast-time
`Spell::Apply` and delivery-time `SpellTransport::Tick` use to admit a
`PointEffect` to the ticking population performs no field access of its own.
Both of `SpellTransport::Tick`'s own two calls to that registrar compute the
identical container operand, instruction-for-instruction, that both of
`Spell::Apply`'s own calls compute — read directly, not assumed from the
shared function name. — MAGIC-197
Cast-time attachment is the only place a *null target* is excluded, dropping
construction outright when the spell has no target unit; a separate,
target-blind exclusion also exists at delivery-system dispatch, but it does
not inspect+44. Nothing downstream repeats the target-specific check.
Whether a load-time identity repair miss (leaving+44 null on an already-live
PointEffect) is reachable by any admitted native save/load population, and so
whether either unconditional site can actually fault, is Unknown. —
MAGIC-TARGETID-182, MAGIC-TICKGATE-183 (partially retracted for the Position
registration gloss and delivery-admission wording; the target gate remains valid)

### Direct cast-admission field sources

The selected PointEffect caller writes its own `+0c` from Spell `+08`,
`+0e=2*(Spell+08)+9`, and `+41=1` exactly when cached Spell `+0a` is zero.
Delivery 1 appends this object. Delivery 2 constructs a SpellTransport whose
Position source is caster Position and whose `+44` holds the child; it
appends the transport. The selected transport constructor keeps `+48=0`.
No direct transport caller store copies the child's `+0c/+0e/+40/+41` into
the transport. Its inherited `+0e/+40/+41` stores are `0/0/1`; its own
`+0c` and both objects' `+08` remain without direct assignment in this graph.
— MAGIC-225, SAV-1068, SAV-1069, MAGIC-ATTRGATE-118, MAGIC-215

Other delivery values skip wrapping and registration after child construction
has already happened. MAGIC-TICKGATE-183's former no-construction wording is
retracted for that suffix. The direct registrar links a payload pointer
without dereferencing it. These results cover the two selected cast
constructors and their direct admission windows; allocator behavior, arbitrary
aliases, unexpanded helper callbacks and later writes before first SAVE
remain Unknown. — MAGIC-225, MAGIC-TICKGATE-183

`PointEffect`'s own tail reads `+0x14` five times, not three, and
`AreaEffect`'s own tail three: an early, unrelated read on the object the
attached `Effect` payload's own `+0x44` points at (`PointEffect` only,
clearing that reference when zero); the guarded read proper, on the cached
target copy, at the site the null check on that copy immediately precedes;
the caster's own `+0x14`, twice, gating the credit write and the final notify
call; and, in `PointEffect`'s own tail only, one further, independent read of
the target's own `+0x14` immediately before the final notify call. The
guarded read and this further target read are both the target `Unit`'s own
owning-Player field — the same field and meaning already established for the
actor path generally, not a field on `PointEffect` or on the caster.
`Token`'s own copy constructor is also a located reader of the six classes'
own `+0x14`, copying a source object's own field into the object under
construction. `UNIT-OWNER-009` is partially retracted for treating its value
space as a universal authorship bound; the identification this cites
(`actor+0x14` names the owning `Player`) is the retained
local-ownership-chain reading, not the retracted universal claim. — SAV-1060,
UNIT-OWNER-009

`AreaEffect`'s own+44 (a typed inner Effect reference) and `SpellTransport`'s
own+44/+48 (`SAV-CASTCONT-1006`) are consumed through the standard typed
archive-reference mechanism, not `PointEffect`'s raw identity dword needing
`R1108`'s explicit map lookup. That contrast is about the persistence
mechanism, not about guard strength: AreaEffect's own consumer dereference is
itself unconditional in its ordinary branch, and at least one further
AreaEffect vtable method dereferences+44 with no null check and no caller
census run against it. Both siblings also have post-load hooks distinct from
their own Tick that touch the same fields; no image-wide reader census was
run for either. — MAGIC-SIBLINGREF-184

## Ticking population, filter and removal

The shared-list driver is `R0641`. Its container holds head/tail node
pointers at its own+0x4/+0x8; each node holds next/payload at+0x0/+0x8. The
walk uses the standard GetHeadPosition/GetNext idiom, and that idiom advances
its own position past a node before the node's own loop-body processing runs,
so the driver's own same-pass removal — always of the member it is currently
processing — cannot corrupt the walk: the removal call chain is never handed
the walk's own iterator, only the container and the member pointer, and the
container it receives is the identical pointer the walk itself uses. Removal
of some other member during the same pass, for instance by a Tick callee, is
not covered by that negative: the live position is the next node, and nothing
read excludes that node being the one removed. — MAGIC-187

Before dispatching Tick, the driver reads exactly two fields on a member —
its own+0x3c, and its own vtable pointer at+0x00, which the dispatch itself
loads — plus, conditionally, the+0x14 field of whatever object+0x3c points
at, which is not a field on the member. Neither branch the+0x3c-derived
reads control skips the Tick call; every member the walk reaches is
Tick-dispatched unconditionally. This confirms cast-time attachment remains
the only target-specific gate: the driver adds no further filter of its own.
— MAGIC-188

The driver's own body writes exactly one field on a member,+0x3c, cleared to
null under the same two-read condition; no other instruction in its body
touches member memory as a write. It never reads or writes a member's+0x44
at all, so a member whose+0x44 is null — including one left null by a
load-time identity-map miss — is not filtered, skipped, or specially routed
by this function; such a member reaches Tick exactly like any other. Whether
Tick's own callee ever receives such a member is the separate reachability
question the identity-repair section above leaves Unknown. — MAGIC-189,
MAGIC-190

Removal, when triggered by a nonzero+0x40 after Tick returns, happens inside
the driver's own body, not deferred to a later pass: the driver calls a
wrapper on the container, which calls a lookup primitive and, when that
returns non-null, an erase primitive on the same container, and the wrapper
then dispatches the member's own vtable+0x04 slot with an argument requesting
destruction and release. The lookup primitive locates the node holding a
given payload without touching the payload; the erase primitive unlinks that
node from the doubly-linked list, symmetric for both the head/tail and the
interior case, then hands the node — not the payload — to a container
method that frees it by position, unread past that confirmation. Removal
is: locate the node holding that payload; unlink and free that node only
if one is found; then, independently, delete the payload whenever the
payload pointer is non-null. The two steps carry different guards, and a
payload that is not in the container is deleted without being unlinked. A
bulk remove-all routine is a second caller of the removal wrapper,
clearing the whole container one member at a time through the identical
wrapper rather than a separate bulk free path. The+0x04 slot was read for `PointEffect`, `AreaEffect` and
`SpellTransport`: each is the scalar-deleting-destructor shape — real
destructor, argument bit-0 test, conditional operator delete — at three
distinct thunk addresses, closing the earlier gap left assumed for
`AreaEffect`. — MAGIC-187, MAGIC-197, MAGIC-216

Each class's own real destructor deletes exactly the fields it owns, never a
sibling class's own offset, read as complete instruction sequences rather
than assumed from the shared thunk shape above. `PointEffect`'s own real
destructor conditionally deletes and clears only+0x48, its own inner Effect
payload; `AreaEffect`'s own conditionally deletes and clears only+0x44, its
own inner Effect payload; neither ever reads the other offset. All three
sibling real destructors call one shared base destructor, which installs no
vtable of its own and does nothing but forward one level further, to a base
that touches only two unrelated offsets — neither+0x44 nor+0x48 anywhere in
that further base's own body — though that base's own last call, an unread
further base destructor on the same object, remains uncensused as a
possible second deleter. `SpellTransport`'s own real destructor is
structurally different: it treats+0x44 and+0x48 as two
independent owning pointers, each put through the identical
dispatch-then-clear block the single-field classes above use, once per
field, in the one body. On the one delivery path traced, Tick's own hand-off
of exactly one field to the shared registrar is immediately followed, still
inside Tick's own body, by an unconditional clear of both fields, so this
destructor's own dispatch never runs against a field the container has
already been given; a `SpellTransport` destroyed before its own delivery
countdown completes still holds a live, unregistered child and this
destructor deletes it correctly. Whether a producer this experiment did not
trace can ever leave both fields simultaneously non-null at the moment the
countdown completes — which would leave the non-primary field cleared by
Tick without being deleted or registered — is Unknown; the only writers of
either field this experiment located are the constructor (unconditional for
+0x44, always zero for+0x48) and the archive's own LOAD arm. —
MAGIC-213, MAGIC-214, MAGIC-215

Neither `PointEffect`'s own post-load hook nor `AreaEffect`'s own contains
any call to the shared registrar or the append primitive, extending the
identical, already-published absence for `SpellTransport`'s own post-load
hook to the other two classes in this family: no post-load hook in this
family re-registers a doubly-referenced child into the shared container:
each hook only null-guards its own field and forwards a repair call to
whatever it already holds. — MAGIC-213, SAV-1050

A save LOAD puts a rebuilt `SpellEffect`-lineage member into a container reached by the identical displacement formula as this list
during the archive's own load arm, not in a later pass and not on the first
Tick. The container-level `Serialize` clears the list, reads a count, then
resolves each element through the typed archive-reference mechanism and
appends it with the identical primitive the Tick/Apply-time registrar above
uses. A separate pass walks the same list again after the whole document
loads and dispatches each present member's own post-load hook, but that pass
inserts nothing; whatever it walks was already there. Whether the container
this load arm reaches and the container the Tick-time registrar reaches are
the same runtime object, not merely the same displacement formula applied to
different base pointers, is not established. — MAGIC-202

The STORE arm of that same container-level `Serialize` writes a count and
then one generic write call per element, carrying only the element pointer;
no instruction in the loop pushes a class descriptor, unlike the LOAD arm's
own per-element resolution. Class identification on a STORE is deferred
entirely to the generic write call itself, which looks up the element's own
runtime class through its vtable rather than being told in advance which
class to expect. — MAGIC-209

## Application

```
Spell::Apply(caster, target, x, y)                                  R0003
  power, then R0625 reloads and raises spell+0x09 and fills
  spell+0x0e / +0x0f / +0x10 (the spell-power fields); an item cast mutates its weapon-owned Spell too

  if (base + spread) != 0 and the spell is neither 6 (Heal) nor 11 (Drain Life):
        an Effect_DirectDamage (0x60 B) is built; its combat block at +0x48 gets
            block+0x13 = base, block+0x14 = spread, block+0x15 = the Sphere
        so SCHOOL i IS DAMAGE KIND i, and the block carries no physical damage

  ...and then RETURNS TO THE ATTACHMENT TAIL. The per-spell arm below is reached ONLY when the
  scratch pair is 0, or the spell is 6 (Heal) or 11 (Drain Life) -- so the seven spells with
  damage columns never execute an arm at all, and the arm they share (L03046) is dead code.

  the per-spell arm, dispatched through the 29-dword table at L03242 (L03243), which
  collapses the 28 shipped ids onto 17 addresses:

        L03046  fire_arrow fire_ball wall_of_fire acid_stream lightning
                    prismatic_spray meteor_storm            (never entered -- see above)
        L05183  the four Protection spells              L13011  fire_sacrifice
        L05185  haste, slow                             L03244  heal
        L05184  bless, curse                            L03256  drain_life
        L13012  freezing_cloud                          L13013  poison_cloud
        L05462  light                                   L05463  darkness
        L13014  shield                                  L05129  wall_of_earth
        L05759  stone_curse                             L13015  invisibility
        L05128  control_spirit                          L05853  teleport

  Eleven arms build their Effect from the row's own `Effects` column (R1047 ->
  R1000), which supplies the KIND and, where the arm writes no +0x40, the magnitude;
  the parser reads ONE record, so stone_curse's second (`defence=-20`) is never built. Five
  use the plain ctor R1049 instead -- bless, curse, invisibility, wall_of_earth and
  the dead default -- so those spells' `Effects` columns are never parsed at all.

  Selected shapes:
        the four Protection spells share one arm: magnitude = power/2, duration as above
        Bless / Curse                            magnitude = (power*4)/5 + 20, negated for Curse
        Haste / Slow                             magnitude = power/15 + 1,     negated for Slow
        Heal                                     refused across the diplomacy table
        Fire Sacrifice     base = min(health + mana, 255), spread = min(min(health + mana +
                                                 power, 512) - base, 255); leaves 1 health and
                                                 0 mana; stamped FIRE by a hard-coded setter
                                                 rather than by the Sphere table. It also
                                                 builds a protectionFire-100 / 32-tick effect
                                                 and NEVER ATTACHES IT -- no AddEffect, no
                                                 list insert, and the pointer is not read again
        Stone Curse        duration = T(10), then x (100 - target.protectionEarth)/100 with a
                                                 floor of 1 tick -- the only place in the image
                                                 where a resistance shortens an effect
        Bless / Curse      the magnitude is a PROBABILITY: R0265 takes the maximum (23)
                                                 or the minimum (27) of the damage roll when
                                                 magnitude > U[0,100], i.e. 20/101 at power 0
                                                 and 100/101 at power 100
        Prismatic Spray    the cast's id==14 arm calls R0269; its ordered selected-victim
                                                 list receives one complete apply per entry.
                                                 The common item wrapper refuses id 14 after a
                                                 caster-item route has already run this selector
        Drain Life                               moves health from the target to the caster
        Control Spirit                           consumes a nearby stage-2 actor on the dead list
                                                 (health := -10001, stage := 5) and raises a
                                                 "Ghost" built by name from the Data.bin Units
                                                 row of that name, placed at the corpse's cell
                                                 and owned by the CASTER's Player. Of the
                                                 corpse only reaction (halved + 1), Mind,
                                                 Spirit, healthMax (halved), health and the
                                                 first 16-bit word of the +0xa6 and +0xbe
                                                 blocks are copied. The power enters no
                                                 expression in this arm
        Teleport                                 attempts relocation; destination may veto it

  every lasting effect gets +0x3c = the kind, +0x3d |= 1 (duration), +0x40 = the magnitude,
  +0x42 = the ticks, +0x0c = the spell id, +0x0e = spellId*2 + 8 (an art index)

  attachment:  DistributionSystem == 1 -> a PointEffect, which REQUIRES a target unit; with
                                          none it prints "Spell, oops - can't cast point
                                          effect of x,y" and the effect is DROPPED.
                                          PointEffect+0x41 = (Defensive != 1), the inverse of
                                          cached Spell+0x0a = (Defensive == 1).
               otherwise               -> an AreaEffect sized by the Radius column (unscaled),
                                          with effect+0x4c = (AreaEffectDuaration << 4), and
                                          += (power << 4)/10 and +0x08 = 1 when that is
                                          nonzero; DistributionSystem 5 instead sets +0x08 = 2
                                          and +0x4c = 0
               DeliverySystem == 2     -> the whole thing is wrapped in a SpellTransport that
                                          flies at SpellEffectSpeed and delivers on arrival;
                                          Lightning and Prismatic Spray overwrite its counter
                                          with the literal 10
               anything else           -> nothing is attached at all
```

**Control Spirit.** The by-name constructor takes the first Units row whose name equals `Ghost` exactly (case-sensitive, ordinals from `0x1a`, skipping `0x1c`..`0x3e`); the Humans table is not consulted, and the row is identical on EN and RU (`MAGIC-249`). The corpse is found by cell equality against the target cell, not by the target pointer, and is consumed before the new actor is placed at radius 0; a placement failure deletes the new actor and leaves the corpse consumed (`MAGIC-251`). The mission difficulty is not applied in the arm or in the constructor chain (`MAGIC-250`, Medium). The item entry refuses only spell id 14, so an item or scroll cast of id 25 reaches this arm (`MAGIC-251`); no authored weapon or item carries id 25, and a shop Scroll or Book can (`MAGIC-252`, Medium). `MAGIC-SING-019` (c) is narrowed by these claims.

**Which spell is aimed at what.** `Spell Target` (title 5, `getParam(row, 4)` at `R0621`)
routes the *cast*: 1 carries the target unit, anything else carries the point. `Distribution
system` (title 9) shapes the *apply*. They agree on 27 of 28 rows — `Shield` ships `Spell Target 2`
with `Distribution 1`. Nothing in the area collector `R0109` consults the diplomacy table:
**`Heal` is the only arm in the whole apply that does**, and an area spell burns its own party.

**Attachment to an actor** (`R0612`) uses Effect Token+0x0c as identity, including
non-spell Potion id 0, and sets `actor+0x144 |= 1 << id` — the bitmask the
damage resolver reads for Bless and Curse and the cast reads for Invisibility. **Ordinary timed
attachments share one record per id**: an effect of the same id already present is un-applied,
given the new magnitude and duration, and re-applied (an incoming `continuous` one has only its
counter refreshed; `MAGIC-ATTACH-016` (amended, partially retracted and superseded in part in the ledger), `MAGIC-POISONREFRESH-158`). **Bless and Curse
annihilate**: casting either on an actor carrying the other removes that one and applies nothing.
`R0673` ticks `+0x42` down, clears the id bit on expiry and ORs in a fifth `+0x3d` bit
(`[L05210] = 128`) that is not one of the five duration words; a `+0x42` above 9600 never counts
down at all, which no shipped spell can reach.

See [overlap rules](overlap.md) for Fire Wall/Poison Cloud and
[training and weapon casting](training.md) for award producers.

### Non-spell Potion lifetime

`MAGIC-CONSUME-142`, `MAGIC-CONSUME-143` and `MAGIC-CONSUME-144` qualify this attachment mechanism. Only mode bits 1
and 2 are timed; charges bit 4 alone is not. Item use turns mode 0 into singleuse 8.
Current-health/mana potions add 30 or 100, capped at maxima. The four +1 attribute
potions change the live attribute without changing its modifier byte and leave no
timer. The common effect-dispatch tail immediately invokes Human derive, which caps it at 50 plus signed modifier.

The five timed rows are absorption+50 for 480 ticks, health/mana regeneration+100
for 960, and the two regeneration bonuses+250 for 1920. All retain Token id 0.
On repeat, attachment finds the old id 0 even if the new characteristic differs.
Non-continuous replacement un-applies the old effect, overwrites only magnitude
and duration, then re-applies the old characteristic. It neither appends an
independent timer nor changes old kind/mode. Incoming continuous refreshes duration
only. A class-inapplicable mana effect can still consume its Potion.

The actor ticks attachments before the health branch, except terminal act `0x10`.
Expiry un-applies non-continuous effects, clears the id bit and marks removal;
continuous expiry does not unapply. Saves carry the attachment's remaining
counter, kind/mode/id, live attributes and the already-applied modifier block.
These readers do not replay a fresh Potion or reset its timer. Native drink-save-load behavior and town-to-mission preservation remain Unknown.

### Prismatic Spray's selected victims (`MAGIC-SPRAY-134`…`MAGIC-SPRAY-137`)

`R0269` constructs a local output list and calls
`R0109(session,caster,primary,out,cap)`. The cap is
`(u8)min(power/20+2,7)`: book power comes from Sphere 2, while an item/staff uses the signed
kind-`0x29` effect's `+0x42`. It is a final victim count, not a distance.

The selector first alarms the primary and applies the directional hostility flip. It asks
`R0110` for the caster group's shared candidate lists: sight is cleared once and stamped by
every member; the global actor list is walked head to tail; the first member supplies diplomacy and
the See-invisible exception; health-below-one candidates move from A to B. It then copies A followed
by B into a 100-entry pointer region at stack offset +0x60, with no bound check. Parallel scores are written
to `[session+0xd74+4*i]`; this scratch region's capacity is Unknown. A score is
`((edgeDistance<<8)+turnCost)&0xffff`; B score is
`((((edgeDistance<<8)+turnCost)&0xffff)<<8)`. If A was empty, the builder has already moved all of B
into A, so an only-corpse population uses the A formula. With living A, B remains in its own list;
the bytes do not prove that its transformed score can never pass the selector thresholds because
the reachable edge-distance domain was not established.

On the nonempty-A path, up to `cap` winners are selected by repeated strict-minimum scans. Equal
scores preserve source-list order. Ten winner indices and ten score thresholds begin at 65530. A
winner index below 60000 gates only stamping its score scratch to 65500; append separately requires
both the winner index and saved threshold below 65000. The primary is appended first, selected
secondaries follow in rank order, a secondary equal to the primary is skipped, and the tail is
removed until `count <= cap`. Thus the primary bypasses secondary visibility, diplomacy and corpse admission but
still consumes one place. With nonempty A, cap zero removes it. With empty A, however, an earlier
arm appends the primary and returns before reading the cap. Ordinary book/item power produces 2..7,
so this exception does not exceed the authored cap. Neither pointer-copy loop nor a custom cap above
ten is checked. More than 100 candidates leaves the pointer stack region; a cap above ten leaves the
winner/threshold regions. Whether a candidate score leaves its session-owned scratch region is
Unknown.

The same ordered list is applied, then sent. Opcode `0x8a` carries an ordinary caster; `0x8c`
carries a packed cell for a temporary caster. Both client arms refill source `+0xac/+0xb0` with the
victim ids in order, and picture 36 draws one link for each id. The list is therefore both the
simulation apply order and the visible branch order.

`Spell+0x09` is separate. Max Range plus the power bonus is copied to order `+0x14` and controls
cast approach/admission. It neither limits group sight nor enters `R0109`. Script instant 24
structurally creates the unit-targeted state that can converge here, but neither preserved campaign
root authors spell 14 in instant 24. Instant 21 makes the point-targeted state with a null target;
a custom spell-14 record is predicted to fault at the selector's immediate primary dereference.
That prediction was not run. A temporary caster built for states `0x0d` and `0x0e` has owner `+0x14` and group
`+0x70` zero from its constructors, and no call site of its builder assigns either (`MAGIC-235`); the selector
reads the caster's owner byte and its group list with no null test, so the same prediction covers a spell-14
cast by that caster. The first unguarded read is the owner read inside the helper called at selector entry,
before the group read and the empty-list test; no try block is in the frames read (`MAGIC-241`). Whether a later tick
assigns an owner is Unknown.

## Effect duration

`effect+0x3d` carries the game's own five words, parsed by `R0857`:

```
permanent   0     no counter
duration    1     +0x42 counts remaining ticks
continuous  2     +0x42 counts, and the effect re-applies
charges     4     +0x42 counts uses
singleuse   8     applied once; does NOT raise a stat's cap byte (HERO derived-stat step 0)
```

Any of `1 / 2 / 4` also makes the effect dispatch read the magnitude as a **signed 16-bit** value at
`+0x40`; without them it is a signed 32-bit value.

## Teleport destination refusal

An ordinary mage book cast spends its cost before the later teleport
application. Relocation tests every footprint cell against mover+5 in the
map's plane at+20000, then calls a separate placement predicate. Either early
veto leaves Position unchanged. The teleport arm ignores that return and sends
its0x20 update anyway; these paths contain no mana refund. Cast presentation
was already produced, and the client picture60 route owns the two sprites.
An insufficient-mana cast exits before those later paths. — MAGIC-TELEPORT-174

The conditional proof executes separated original fragments with explicit
predicate, occupancy, visibility and message-service stubs. It does not certify
native reachability of every synthetic actor state, footprint-edge behavior or
exact rendered frames. — MAGIC-TELEPORT-174
